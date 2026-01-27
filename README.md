# goxe

reduce large volumes of repetitive logs into compact, readable aggregates.

goxe is a high-performance log reduction tool written in go. it ingests logs (currently via syslog/udp),
normalizes and filters them, and aggregates repeated messages into a single-line format with occurrence counts.
the result is less noise, lower bandwidth usage, and cheaper storage without losing visibility into recurring issues.

goxe is designed to run continuously in the background as part of a logging pipeline or sidecar.

## requirements

* go 1.25.5 or higher (to build from source)

### aggregation behavior

goxe performs several transformations before aggregation:

* strips timestamps and date prefixes
* converts text to lowercase
* removes extra whitespace
* filters out configurable excluded words
* applies basic ascii beautification

after normalization, identical messages are grouped together and reported with repetition counts.

example input:

```
dec 24, 2025 16:30:17 ERROR: connection failed 001 128.54.69.12
dec 24, 2025 16:30:18 ERROR: connection failed 002 128.34.70.12
dec 24, 2025 16:30:19 ERROR: connection failed 003 128.54.69.12
```

aggregated output:

```
        partial report
----------------------------------
origin: [::1]
- [3] ERROR: connection failed *  -- (first seen 16:30:17 - last seen 16:30:19)
----------------------------------
```

## features

* worker pool for parallel processing
* thread-safe state management
* automated partial reporting
* log normalization and filtering
* ascii beautification
* timestamp and date parsing
* graceful shutdown and signal handling
* similarity clustering (group near-identical messages)
* syslog/udp network ingestion
* configuration file support
* output log file
* firstseen field to track initial occurrence
* event burst detection
* notification dispatch pipeline

## usage

- default behavior:
  - goxe listens on udp port `1729` by default (configurable).
  - on first run the tool creates a default config.json in the user's config directory:
    - linux: `$XDG_CONFIG_HOME` or `$HOME/.config` → `goxe/config.json`
    - macos: `~/Library/Application Support/goxe/config.json`
    - windows: `%APPDATA%\goxe\config.json`
  - the app reads `options.Config` from that file; the defaults are:
    - `port`: 1729 — udp port to listen on
    - `idLog`: hostname — identifier added/removed from logs
    - `patternsWords`: [] — list of ignored words
    - `generateLogsOptions.generateLogsFile`: false — write periodic file report
    - `generateLogsOptions.hour`: "00:00:00" — scheduled hour for file generation
    - `webHookUrls`: [] — webhooks to call when alerts fire
    - `burstDetectionOptions.limitBreak`: 10 — burst detection threshold (seconds × count)
  - to change behavior edit the `config.json` file and restart goxe or use the upcoming config reload path.

- routing system logs to goxe:
  - configure your system logger (rsyslog, syslog-ng, journald-forwarder, etc.) to forward or send logs to `udp://<host>:1729` (replace host/port as needed).
  - see your os documentation for forwarding syslog to a remote udp port (linux, macos, windows).

- app integration:
  - any app that can send syslog/udp can forward logs to goxe (host:1729 by default).
  - examples (conceptual):
    - node: use a syslog/bunyan/winston syslog transport to forward logs via udp to goxe.
    - go: use the stdlib net/dial udp or a syslog client to send messages to goxe.
  - alternatively, forward your system logger to `udp://<host>:1729`.
  - note: docker support is not available yet — running goxe in a container is not officially supported in this release.

- limitations:
  - v1.x does not yet forward processed aggregates to an external log service; forwarding/shipper integrations are planned for later releases.
  - content is currently processed as strings; for the highest-performance zero-copy pipelines consider using the buffer/raw-field options (internal optimizations applied in this release).

## testing

- benchmark runs (example) can be added as images to show before/after results. placeholder below:

![benchmark results placeholder](https://i.ibb.co/Z1BqmGC7/image.png)

- note on allocs: current benchmarking shows ~2 allocs/op in the udp ingestion + processing path. this is expected with the current api because:
  - one allocation is typically the creation of the normalized key (the sanitized string used as the map key),
  - the other allocation can come from creating a new logstats entry for a brand-new message key.
- how to reduce further:
  - change the pipeline to process bytes instead of strings (breaking change) or use a hash/interning strategy for keys, which avoids per-message string allocations for repeated messages.
  - optimize sanitizer to do a single-pass transformation into a pooled builder to avoid intermediate temporaries.
- the above optimizations are planned; this release focuses on fixing per-message regex recompiles, adding shared pools and safe zero-copy buffer ownership to reduce gc pressure.

## maintainers

* @dumbnoxx

## license

licensed under the apache license, version 2.0. see the [license file](LICENSE) for details.

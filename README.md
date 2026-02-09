# goxe

![Made with VHS](https://vhs.charm.sh/vhs-5h1jmrUc7TMULoD2Mo48Ci.gif)

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
* remote syslog/network shipping support


## usage

- **default behavior**:
  - goxe listens on udp port `1729` by default (configurable).
  - on first run the tool creates a default `config.json` in the user's config directory:
    - **linux**: `$XDG_CONFIG_HOME` or `$HOME/.config` → `goxe/config.json`
    - **macos**: `~/Library/Application Support/goxe/config.json`
    - **windows**: `%APPDATA%\goxe\config.json`
  - the app reads `options.Config` from that file; the defaults are:
    - `port`: 1729 — udp port to listen on
    - `idLog`: hostname — identifier added/removed from logs
    - `pattenersWords`: [] — list of ignored words
    - `generateLogsOptions.generateLogsFile`: false — write periodic file report
    - `generateLogsOptions.hour`: "00:00:00" — scheduled hour for file generation
    - `webhookUrls`: [] — webhooks to call when alerts fire
    - `bursDetectionOptions.limitBreak`: 10 — burst detection threshold (seconds × count)
    - `shipper.address`: "" — remote address to ship processed logs (e.g., "127.0.0.1:5000")
    - `shipper.flushInterval`: 30 — interval in seconds between network transmissions
    - `shipper.protocol`: "tcp" — transmission protocol (tcp, udp, etc. via [net.Dial](https://pkg.go.dev))
    - `ReportInterval`: The interval in minutes for generating summaries of processed logs.
    - `BufferUdpSize`: The size of the UDP buffer for receiving logs via UDP.
  - **hot reloading**: goxe monitors the `config.json` file in real-time. Any changes saved to the file are automatically applied without requiring a restart.

- **routing and shipping**:
  - **ingestion**: configure your system logger (rsyslog, syslog-ng, etc.) or any application to forward logs to `udp://<host>:1729`.
  - **remote shipping**: enable `shipper.address` to forward processed log aggregates to an external service. Goxe will batch and send statistics in JSON format:
    ```json
    {
      "origin": "web-server-01",
      "data": [
        {
          "count": 42,
          "firstSeen": "2024-03-20T10:00:00Z",
          "lastSeen": "2024-03-20T10:05:00Z",
          "message": "Invalid password attempt for user admin"
        }
      ]
    }
    ```

- **app integration**:
  - **system-wide**: see your OS documentation for forwarding syslog to a remote UDP port (Linux, macOS, Windows).
  - **custom apps**: any app capable of sending UDP/Syslog packets can use Goxe as a target:
    - **node**: use a syslog/bunyan/winston transport to forward logs.
    - **go**: use the [std net](https://pkg.go.dev/net) package to dial UDP.
  - **note**: docker support is not available yet running goxe in a container is not officially supported in this release.

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

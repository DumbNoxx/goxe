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
Aug 22, 2022 23:59:59 ERROR: Connection failed 001 128.54.69.12
Aug 22, 2022 23:59:59 ERROR: Connection failed 002 128.34.70.12
Aug 22, 2022 23:59:59 ERROR: Connection failed 003 128.54.69.12
```

aggregated output:

```
        PARTIAL REPORT
----------------------------------
ORIGIN: [::1]
- [3] ERROR: connection failed *  -- (Last seen 16:30:17)
----------------------------------
```

## architecture

goxe is built for concurrency and throughput:

* worker pool architecture for parallel log processing
* centralized, thread-safe aggregation state using `sync.mutex`
* periodic partial reporting using `time.ticker`
* streaming design with low memory overhead

the system is optimized to handle high log volumes with minimal latency.

## roadmap

### completed

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

### v1 sprint 

* notification dispatch pipeline
* event burst detection

### planned future
* additional ingestion backends

## maintainers

* @dumbnoxx

## license

licensed under the apache license, version 2.0. see the [license file](LICENSE) for details.

![Made with VHS](https://vhs.charm.sh/vhs-5h1jmrUc7TMULoD2Mo48Ci.gif)

reduce large volumes of repetitive logs into compact, readable aggregates.

goxe is a high-performance log reduction tool written in go. it ingests logs (currently via syslog/udp),
normalizes and filters them, and aggregates repeated messages into a single-line format with occurrence counts.
the result is less noise, lower bandwidth usage, and cheaper storage without losing visibility into recurring issues.

goxe is designed to run continuously in the background as part of a logging pipeline or sidecar.

<div align="center">

  [![Compile](https://img.shields.io/github/actions/workflow/status/dumbnoxx/goxe/go.yml?label=compile)](https://github.com/dumbnoxx/goxe/actions/workflows/go.yml)
![License](https://img.shields.io/badge/license-Apache%202-blue)

[](https://github.com/DumbNoxx/goxe/blob/main/LICENSE)

</div>

## recommended install

```bash
go install github.com/DumbNoxx/goxe/cmd/goxe@latest
```

other installation methods and full usage examples are available in the documentation.


## requirements (for building from source / contributing)

* go 1.25.5 or higher

build from source:

```bash
git clone https://github.com/dumbnoxx/goxe.git
cd goxe
go install ./... 
```

## aggregation behavior

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

## documentation

visit our [official documentation](https://dumbnoxx.github.io/goxe-doc/)

## contrib

1. fork
2. implement change
3. open pr

please read [CONTRIBUTING](./CONTRIBUTING.md).

## maintainers

* @dumbnoxx

## license

apache 2.0

# [Goxe]

[![Go Report Card](https://goreportcard.com/badge/github.com/DumbNoxx/Goxe)](https://goreportcard.com/report/github.com/DumbNoxx/Goxe)
[![Go Reference](https://pkg.go.dev/badge/github.com/DumbNoxx/Goxe.svg)](https://pkg.go.dev/github.com/DumbNoxx/Goxe)
[![License: Apache License 2](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

**[Goxe]** is an efficient solution written in Go designed to [reduce large logs sent to your reduction server into a single-line format with a single message showing repetition counts].

---

## 🗺️ Development Roadmap
This project is currently under development. The following sections detail functional and planned stages:

### Phase 1
- [x] Worker Pool architecture for parallel processing
- [x] Thread-safe state management using `sync.Mutex`
- [x] Automated partial reporting via `time.Ticker`

### Phase 2
- [x] Normalization by removing spaces and converting to lowercase
- [x] Log filtering: excluding words within a slice from being added to the report
- [x] Basic ASCII beautifier
- [ ] Log parsing to remove timestamps and dates

## ✨ Features

* 🚀 **Fast and Lightweight:** Natively compiled for optimal performance.
* 🔧 **Easy to Use:** Designed to run in the background to capture your service logs seamlessly.

### Prerequisites
* Go 1.25.5 or higher (if compiling from source).

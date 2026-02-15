# CHANGELOG

## v1.3.2 (2026-02-15)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

- feat(update): add automatic GitHub release version checker (2026-02-11)

- feat: implement hot-reload via SIGUSR1 and graceful shutdown report (2026-02-12)

### Fixed

- fix: handle platform-specific signals for windows builds (2026-02-12)

- fix: Updated update signal notification to English (2026-02-12)

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

- fix(processor): remove specific words from logs instead of dropping the entire entry (2026-02-10)

- fix(processor): fix reload patterns words in config.json (2026-02-10)

- fix: simplify log persistence and ensure file generation (2026-02-11)

- fix: enable dynamic ticker resets for config hot-reloading (2026-02-12)

### Others

- Merge pull request #19 from DumbNoxx/refactor/optimize-log-sanitizer (2026-02-14)

- chore(release): update changelog for v1.3.1 (2026-02-12)

- chore(release): update changelog for v1.3.0 (2026-02-12)

- Merge pull request #18 from DumbNoxx/feat/auto-update-signal-handling (2026-02-12)

- Merge pull request #17 from DumbNoxx/fix/dynamic-ticker-interval-update (2026-02-12)

- Merge pull request #16 from DumbNoxx/refactor/improve-concurrency-processor (2026-02-11)

- chore: delete fmt.println wrong in clean function (2026-02-11)

- Merge pull request #15 from DumbNoxx/fix/file-reporting-logic (2026-02-11)

- Merge pull request #14 from DumbNoxx/feat/version-checker (2026-02-11)

- chore(release): update changelog for v1.2.2 (2026-02-10)

- Merge pull request #13 from DumbNoxx/fix/scrubbing-logic-ignored-word-over (2026-02-10)

- chore(release): update changelog for v1.2.1 (2026-02-10)

- Merge pull request #12 from DumbNoxx/fix/scrubbing-logic-ignored-word (2026-02-10)

- Merge pull request #11 from DumbNoxx/refactor/extract-burstDetection-functions (2026-02-10)

- update readme (2026-02-09)

- chore: Rename module to lowercase for consistency (2026-02-09)

- chore(release): update changelog for v1.2.0 (2026-02-09)

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

- refactor: Extract BurstDetection and helpers functions into modules (2026-02-10)

- refactor: reduce critical section in ticker report (2026-02-11)

- refactor: switch string to bytes in cluster processing for efficiency (2026-02-14)

## v1.3.1 (2026-02-12)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

- feat(update): add automatic GitHub release version checker (2026-02-11)

- feat: implement hot-reload via SIGUSR1 and graceful shutdown report (2026-02-12)

### Fixed

- fix: handle platform-specific signals for windows builds (2026-02-12)

- fix: Updated update signal notification to English (2026-02-12)

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

- fix(processor): remove specific words from logs instead of dropping the entire entry (2026-02-10)

- fix(processor): fix reload patterns words in config.json (2026-02-10)

- fix: simplify log persistence and ensure file generation (2026-02-11)

- fix: enable dynamic ticker resets for config hot-reloading (2026-02-12)

### Others

- chore(release): update changelog for v1.3.0 (2026-02-12)

- Merge pull request #18 from DumbNoxx/feat/auto-update-signal-handling (2026-02-12)

- Merge pull request #17 from DumbNoxx/fix/dynamic-ticker-interval-update (2026-02-12)

- Merge pull request #16 from DumbNoxx/refactor/improve-concurrency-processor (2026-02-11)

- chore: delete fmt.println wrong in clean function (2026-02-11)

- Merge pull request #15 from DumbNoxx/fix/file-reporting-logic (2026-02-11)

- Merge pull request #14 from DumbNoxx/feat/version-checker (2026-02-11)

- chore(release): update changelog for v1.2.2 (2026-02-10)

- Merge pull request #13 from DumbNoxx/fix/scrubbing-logic-ignored-word-over (2026-02-10)

- chore(release): update changelog for v1.2.1 (2026-02-10)

- Merge pull request #12 from DumbNoxx/fix/scrubbing-logic-ignored-word (2026-02-10)

- Merge pull request #11 from DumbNoxx/refactor/extract-burstDetection-functions (2026-02-10)

- update readme (2026-02-09)

- chore: Rename module to lowercase for consistency (2026-02-09)

- chore(release): update changelog for v1.2.0 (2026-02-09)

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

- refactor: Extract BurstDetection and helpers functions into modules (2026-02-10)

- refactor: reduce critical section in ticker report (2026-02-11)

## v1.3.0 (2026-02-12)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

- feat(update): add automatic GitHub release version checker (2026-02-11)

- feat: implement hot-reload via SIGUSR1 and graceful shutdown report (2026-02-12)

### Fixed

- fix: Updated update signal notification to English (2026-02-12)

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

- fix(processor): remove specific words from logs instead of dropping the entire entry (2026-02-10)

- fix(processor): fix reload patterns words in config.json (2026-02-10)

- fix: simplify log persistence and ensure file generation (2026-02-11)

- fix: enable dynamic ticker resets for config hot-reloading (2026-02-12)

### Others

- Merge pull request #18 from DumbNoxx/feat/auto-update-signal-handling (2026-02-12)

- Merge pull request #17 from DumbNoxx/fix/dynamic-ticker-interval-update (2026-02-12)

- Merge pull request #16 from DumbNoxx/refactor/improve-concurrency-processor (2026-02-11)

- chore: delete fmt.println wrong in clean function (2026-02-11)

- Merge pull request #15 from DumbNoxx/fix/file-reporting-logic (2026-02-11)

- Merge pull request #14 from DumbNoxx/feat/version-checker (2026-02-11)

- chore(release): update changelog for v1.2.2 (2026-02-10)

- Merge pull request #13 from DumbNoxx/fix/scrubbing-logic-ignored-word-over (2026-02-10)

- chore(release): update changelog for v1.2.1 (2026-02-10)

- Merge pull request #12 from DumbNoxx/fix/scrubbing-logic-ignored-word (2026-02-10)

- Merge pull request #11 from DumbNoxx/refactor/extract-burstDetection-functions (2026-02-10)

- update readme (2026-02-09)

- chore: Rename module to lowercase for consistency (2026-02-09)

- chore(release): update changelog for v1.2.0 (2026-02-09)

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

- refactor: Extract BurstDetection and helpers functions into modules (2026-02-10)

- refactor: reduce critical section in ticker report (2026-02-11)

## v1.2.2 (2026-02-10)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

- fix(processor): remove specific words from logs instead of dropping the entire entry (2026-02-10)

- fix(processor): fix reload patterns words in config.json (2026-02-10)

### Others

- Merge pull request #13 from DumbNoxx/fix/scrubbing-logic-ignored-word-over (2026-02-10)

- chore(release): update changelog for v1.2.1 (2026-02-10)

- Merge pull request #12 from DumbNoxx/fix/scrubbing-logic-ignored-word (2026-02-10)

- Merge pull request #11 from DumbNoxx/refactor/extract-burstDetection-functions (2026-02-10)

- update readme (2026-02-09)

- chore: Rename module to lowercase for consistency (2026-02-09)

- chore(release): update changelog for v1.2.0 (2026-02-09)

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

- refactor: Extract BurstDetection and helpers functions into modules (2026-02-10)

## v1.2.1 (2026-02-10)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

- fix(processor): remove specific words from logs instead of dropping the entire entry (2026-02-10)

### Others

- Merge pull request #12 from DumbNoxx/fix/scrubbing-logic-ignored-word (2026-02-10)

- Merge pull request #11 from DumbNoxx/refactor/extract-burstDetection-functions (2026-02-10)

- update readme (2026-02-09)

- chore: Rename module to lowercase for consistency (2026-02-09)

- chore(release): update changelog for v1.2.0 (2026-02-09)

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

- refactor: Extract BurstDetection and helpers functions into modules (2026-02-10)

## v1.2.0 (2026-02-09)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

- feat: implement automatic version detection via build info and ldflags (2026-02-03)

- feat: allow reporting interval to be customized via config (2026-02-08)

- feat: improve shutdown reliability (2026-02-09)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

### Others

- Merge pull request #10 from DumbNoxx/feat/shutdown-reliability (2026-02-09)

- chore: fix error in readme gif (2026-02-08)

- chore: update gif readme (2026-02-08)

- Merge pull request #9 from DumbNoxx/feat/configurable-report-interval (2026-02-08)

- Merge pull request #8 from DumbNoxx/feat/configurable-udp-buffer (2026-02-08)

- Merge pull request #7 from DumbNoxx/feat/implement-automatic-version-detection (2026-02-03)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

- feat/configurable-udp-buffer (2026-02-08)

## v1.1.2 (2026-02-01)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

- fix: change LimitBreak field to float64 for decimal precision (2026-02-01)

### Others

- chore: update goreleaser (2026-02-01)

- chore(release): update changelog for v1.1.2 (2026-02-01)

- Merge pull request #6 from DumbNoxx/fix/limitbreak-type-precision (2026-02-01)

- Merge pull request #5 from DumbNoxx/fix/increase-udp-buffer (2026-02-01)

- chore: add funding.yml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update goreleaser.yaml (2026-02-01)

- chore: update readme add gif (2026-02-01)

- chore(release): update changelog for v1.1.1 emergency (2026-02-01)

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

- fix(internal/syslog): increase UDP read buffer to 16MB (2026-02-01)

## v1.1.1 (2026-02-01)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

- fix: improve date parsing regex (2026-02-01)

### Others

- Merge pull request #4 from DumbNoxx/fix/improve-date-parsing-regex (2026-02-01)

- chore: delete pkgbuild (2026-01-31)

- chore: delete workflow aur corrupt (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore: Update aur.yml (2026-01-31)

- chore(release): update changelog for v1.1.0 (2026-01-31)

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

## v1.1.0 (2026-01-31)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

- feat(webhook): make webhooks look decent on slack and discord (2026-01-31)

- feat(cli): add flag version (2026-01-31)

- feat(exporter): remote syslog/network shipping support (2026-01-30)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- Merge pull request #3 from DumbNoxx/better-logs-delivery (2026-01-31)

- chore: add automated AUR publishing workflow (2026-01-30)

- chore(release): update changelog for v1.0.1 (2026-01-27)

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

- Merge branch 'main' into better-logs-delivery (2026-01-30)

## v1.0.1 (2026-01-27)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: reduce allocations in udp ingestion and sanitizer (2026-01-27)

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore(release): update changelog for v1.0.0 (2026-01-27)

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v1.0.0 (2026-01-27)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- feat(internal/webhook): add notification dispatch pipeline with slack and discord support (2026-01-27)

- chore(release): update changelog for v0.9.0 (2026-01-24)

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.9.0 (2026-01-24)

### Added

- feat(logs): add FirstSeen field to track initial occurrence (2026-01-21)

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- feat(internal/pool): add event burst detection (2026-01-24)

- chore: fix setup-go warnings and pin version (2026-01-21)

- docs(README): fix inconsistency in README examples (2026-01-21)

- chore(release): update changelog for v0.8.5 (2026-01-20)

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.5 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- docs: add write permissions for releases (2026-01-20)

- chore(release): update changelog for v0.8.4 (2026-01-20)

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.4 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: list structure in goreleaser (2026-01-20)

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore(release): update changelog for v0.8.3 (2026-01-20)

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.3 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: error no defined func main in goreleaser file (2026-01-20)

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore(release): update changelog for v0.8.2 (2026-01-20)

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.2 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: Fix types on goreleaser file (2026-01-20)

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore(release): update changelog for v0.8.1 (2026-01-20)

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.1 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: fix idLog bug no detected (2026-01-20)

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- Update Go version in GitHub Actions workflow (2026-01-20)

- chore(release): update changelog for v0.8.0 (2026-01-20)

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.8.0 (2026-01-20)

### Added

- feat: add output log file option (2026-01-20)

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore(release): update changelog for v0.7.0 (2026-01-19)

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.7.0 (2026-01-19)

### Added

- feat: add configuration file support (2026-01-19)

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- refactor(pkg): change folder name pgk to pkg xd (2026-01-19)

- refactor(pkg/pipelines): change model of pipeline (2026-01-19)

- chore(readme): add task for v1 sprint (2026-01-19)

- docs: update license copyright information (2026-01-18)

- chore(readme): update readmen license change file name (2026-01-18)

- chore(release): update changelog for v0.6.0 (2026-01-18)

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.6.0 (2026-01-18)

### Added

- feat: add graceful shutdown and signal handling (2026-01-18)

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- chore: styles name function memoryUsage (2026-01-02)

- chore(config): add taskfile (2026-01-02)

- chore: add changelog (2026-01-02)

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

## v0.5.0 (2026-01-02)

### Added

- feat(internal): add memoryUsage and control in mu.unlock in console.go (2026-01-02)

- feat: add clusters and normalize logs (2026-01-01)

- feat: add clusters and modify architecture (2025-12-27)

- feat(server): add server udp (2025-12-24)

- feat(internal): add filters, sanitizador, colors, logsLevel (2025-12-24)

### Fixed

- fix: add pattern date (2025-12-25)

- fix: fix bug origin partial report (2025-12-25)

- fix: update title of report (2025-12-24)

### Others

- Merge pull request #1 from elisiei/patch-1 (2025-12-27)

- chore: comments (2025-12-24)

- Initial project (2025-12-24)

- Initial commit (2025-12-23)

- Update README.md (2025-12-27)

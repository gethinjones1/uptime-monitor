# Uptime Monitor

A CLI tool to track endpoint availability, store state in BoltDB, and send Slack notifications on status changes.

## Commands
- `add [name] [url]`
- `edit [name] [new-url]`
- `remove [name]`
- `ls`
- `check`

## Tech Stack
- Go
- Cobra for CLI
- BoltDB for storage
- Slack API for notifications

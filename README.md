# Go SystemD Docker

This project creates a golang based CLI that manages docker containers as systemD processes on Linux system.

## Packages used in this module

- [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [Service](https://pkg.go.dev/github.com/kardianos/service@v1.2.2) | [Go-Systemd](https://pkg.go.dev/github.com/iguanesolutions/go-systemd/v4#section-readme)

## Commands

- Create: Register container as Systemd process.
- Delete: Deregisters container from SystemD process.
- Run: Initiate running Registered SystemD container.
- Start: Continue running stopped one.
- Stop: Stop running SystemD container.
- ls: List registered SystemD processes.
- ps: returns a list of running SystemD processes.

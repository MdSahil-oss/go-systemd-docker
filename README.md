# Go SystemD Docker

This project creates a golang based CLI that manages docker containers as systemD processes on Linux system.

## Packages used in this module

- [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [Service](https://pkg.go.dev/github.com/kardianos/service@v1.2.2) | [Go-Systemd](https://pkg.go.dev/github.com/iguanesolutions/go-systemd/v4#section-readme)

## Commands

- `create`: Register container as Systemd process.
- `delete`: Deregisters container from SystemD process.
- `run`: Initiate running Registered SystemD container.
- `start`: Continue running stopped one.
- `stop`: Stop running SystemD container.
- `ls`: List registered SystemD processes.
- `ps`: returns a list of running SystemD processes.

### cmd: create

Create/registers a container-image as a Systemd process with a name.

**flags:**

- --name | -n : To assign a name to the registering container with Systemd.

**args:**

- arg[0] : Container image name.
- arg[1] (optional) : Name to assign to the registering container with Systemd.

### cmd: delete

Delete/deregisters a container-image Systemd process.

**flags:**

- --name | -n : Name of the deregistering container process.
- --force | -f : Force stop/delete the running instance otherwise running instance cannot be deleted.

**args:**

- arg[0] (optional) : Name of the deregistering container process.

### cmd: run

Creates an instance of registered container as Systemd process.

**flags:**

- --name | -n : Name to assign to the instance of container systemd process.

**args:**

- arg[0] : container systemd process name.
- arg[1] (optional) : Name for the container process instance.

### cmd: start

Continue running stopped systemd process (container-image).

**flags:**

- --name | -n : Name of instance to start.

**args:**

- arg[0] (optional) : Name of instance to start.

### cmd: stop

Stop running systemd process (container-image).

**flags:**

- --name | -n : Name of instance to stop.

**args:**

- arg[0] (optional) : Name of instance to stop.

### cmd: list

List registered systemd process (container-image).

**flags:**

- --name | -n : Name of a particular systemd registered process to list.

**args:**

- arg[0] (optional) : Name of a particular systemd registered process to list.

### cmd: process

List running systemd process instances.

**flags:**

- --all | -a : list all stopped and running ones.
- --name | -n : Name of a particular systemd instance to list.

**args:**

- arg[0] (optional) : Name of a particular systemd instance to list.

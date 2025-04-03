# Go SystemD Docker

This project creates a golang based CLI that manages docker containers as systemD processes on Linux system.

# Getting Started

If you are on Ubuntu then directly execute the following command to build CLI (will promt you for root (sudo) user password)

```shell
$ ./build.sh
```

Otherwise, To build binary executable execute following in your shell

```shell
$ go build -o sysd ./cmd/sysd
```

on successful run of previous command, Test by executing the below command.

```shell
$ ./sysd --version
```

**Note:** Execution of this CLI requires root/admin permission as it deals with Systemd.

## Packages/Frameworks used in this CLI

- [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [Service](https://pkg.go.dev/github.com/kardianos/service@v1.2.2)

## TODO

- Update printing table so that you don't see `status` column on printing.
- Enable logging of services.
- consider to add support for tests.
- Update other commands so on.

## Commands

- `create`: Register container as Systemd process.
- `delete`: Deregisters container from SystemD process.
- `docker`: A group of commands
  - `list`: To list all the images utilized by systemd processes.
  - `process`: To list the images currently in-use by running systemd processes.
  - `remove`: To remove the images utilized by systemd processes.
- `run`: Registers & Starts running container as SystemD process.
- `start`: Start running the registered container as Systemd process.
- `stop`: Stop running the registered container as Systemd process.
- `ls`: List registered containers as Systemd processes.
- `ps`: returns a list of running registered containers as Systemd processes.

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
- --all | -A : Removes all the running instances.
- --force | -f : Force stop/delete the running instance otherwise running instance cannot be deleted.

**args:**

- arg[0] (optional) : Name of the deregistering container process.
- ...args (optional): giving more than one args will result in deletion of all the instances.

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

### cmd: show

Show systemd process configuration for the instance name(container-image).

**flags:**

- --name | -n : Name of a particular systemd registered process or put index to see index configuration.

**Just `sysd show` will show index config only**

**args:**

- arg[0] (optional) : Name of a particular systemd registered process or put index to see index configuration.

### cmd: process

List running systemd process instances.

**flags:**

- --all | -a : list all stopped and running ones.
- --name | -n : Name of a particular systemd instance to list.

**args:**

- arg[0] (optional) : Name of a particular systemd instance to list.

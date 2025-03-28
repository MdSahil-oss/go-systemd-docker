# Go SystemD Docker

This project creates a golang based CLI that manages docker containers as systemD processes on Linux system.

## Packages used in this module

- [Cobra](https://pkg.go.dev/github.com/spf13/cobra)
- [Service](https://pkg.go.dev/github.com/kardianos/service@v1.2.2) | [Go-Systemd](https://pkg.go.dev/github.com/iguanesolutions/go-systemd/v4#section-readme)

## TODO

- `Done`: Find a way to start given containerImage (args[0]) as SystemD process.
- `Done`: Make `svcConfig` stateful so that processes become managable.
- `Done`: Update `start` to make it able to start processes.
- `Done`: Update `stop` to make it able to stop processes.
- `Done`: Update `rm` to make it able to actually remove processes
- `Done`: `list` make it able to list processes.
- `Done`: Started systemd processes only goes in activating state (never gets green)
- `Done`: Create `show` to show service name.
- `Done`: Updated `rm` to remove multiple instances.
- `Done`: `ps` to make it able to list running process.
- `Done`: Update `run` to make it able to install & start processes.
- `Done`: Update GetDockerExecutablePath to find `docker` executable path
- `Done`: Currently CLI downloads and runs docker image but there is no way to prune the downloaded images using this CLI that this CLI Downloaded.
  - Add a sub-cmd `docker` with its following sub-cmds `ls`, `ps`, `rm`, Where
  - `Done`: `ls` prints the docker images that this CLI has pulled in the system so far and present on the system.
  - `Done`: `ps` prints the images currently being used.
  - `Buggy`: `rm` remove an image or all the images (on `all` flag). If the image instance exist on the system then promts user for confirmation.
- `Done`: Add validation to check if there docker image exist.
- `Done`: Try to remove sudo prepending before `sysd`.
- Add support for other docker run flags and more to the `run` sub-cmd.
  - Add support for only relatable flags (Don't think all flags are needed).
  - Give user options to add docker run flags either imperatively or declaratively.
- Update other commands so on.

## Going on

## Commands

- `create`: Register container as Systemd process.
- `delete`: Deregisters container from SystemD process.
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

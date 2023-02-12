# Containers in Go

This repo demonstrates creating and understanding containers from scratch through building Linux namespaces, control groups, and filesystems in Go.

---

## Pre-requisites

- Docker (e.g. [Docker Desktop](https://www.docker.com/products/docker-desktop/))
- [Go programming language](https://go.dev/dl/)
- Any source code editor (e.g. [VS Code](https://code.visualstudio.com/download))
- Ubuntu (because namespace flags are only available in Go for Linux!)

---

## Quick Start

Compile and run the Go code:
```
go run main.go run /bin/bash
```

Check the host name inside and ourside the container, after creating a namespace:
```
hostname
```

---

## What are namespaces?

- A Namespace limits, what can a process see.
- Namespaces are a feature of the Linux kernel, allowing groups of processes to have limited visibility of the host system resources.
- Namespaces limit the visibility of e.g. cgroups, hostname, process IDs, IPC mechanisms, network interfaces and routes, users, and mounted file systems.
- Processes which are running inside a namespace are aware of any changes in that namespace, however, such changes will not be visible to other processes or other namespaces.
- In context of containers as well, a namespace isolates the processes running in the container from the rest of the system, e.g. modifying the hostname, network interfaces, or mounts for processes running in other containers. The running container knows only itself, i.e., the application running on it as a process, and its own file system. It does not reach other processes on the host system and the containers running next to it.

---

## What are control groups (cgroups)?

- cgroups limit, how much resource can a process use.
- cgroups are interfaces in the Linux kernel that can be used to restrict access to computer resources (CPU, memory, I/O) for certain processes.
- In context of containers, this allows you to set limits for a container.
- Containers benefit from cgroups primarily because they allow system resources to be limited for processes grouped by a container.

---

## What's Chroot?

- Chroot command, introduced in Version 7 Unix in 1979, isolates a process from the root filesystem. It hides the files from the process and simulates a new root directory.
- The isolated environment is called a `chroot jail` in which only files explicitly put in the new root filesystem are accessible by the process running in the jail.
- Container technology is based on the concept of `Chroot` but in a modernized way.

---

## How to create, manage, validate, and operate containers?

Check out this repo: [Docker-Container](https://github.com/Memal7/docker-container)

---
source: "https://github.com/gmkurtzer/ginux"
hn_url: "https://news.ycombinator.com/item?id=50333727"
title: "Rocky Linux Creator Built a New 'Not Linux' Kernel in 11 Days Using AI"
article_title: "GitHub - gmkurtzer/ginux · GitHub"
image: "https://opengraph.githubassets.com/858d43611bd23ee52aa9b49ada06892f4d8e5ab9c3676c597c4d44aa361c7651/gmkurtzer/ginux"
author: "signa11"
captured_at: "2026-11-02T07:19:04Z"
capture_tool: "hn-digest"
hn_id: 50333727
score: 3
comments: 0
posted_at: "2026-11-02T06:28:56Z"
tags:
  - hacker-news
---

# Rocky Linux Creator Built a New 'Not Linux' Kernel in 11 Days Using AI

- HN: [50333727](https://news.ycombinator.com/item?id=50333727)
- Source: [github.com](https://github.com/gmkurtzer/ginux)
- Score: 3
- Comments: 0
- Posted: 2026-11-02T06:28:56Z

## Translation

Title: Rocky Linux Creator Built a New 'Not Linux' Kernel in 11 Days Using AI
Article title: GitHub - gmkurtzer/ginux · GitHub
Description: Contribute to gmkurtzer/ginux development by creating an account on GitHub.

Article text:
GitHub - gmkurtzer/ginux · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gmkurtzer
ginux Public Notifications You must be signed in to change notification settings
Star 319 ( 319 ) You must be signed in to star a repository
Readme Apache-2.0 license Activity Stars
41 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,285 Commits 1,285 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills .github/ workflows .github/ workflows .hdae .hdae boot boot control-plane control-plane crates crates docs docs hdae hdae kernel kernel profiles profiles services services specs specs tests tests tools tools .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
A Rust-first x86-64 operating system built around native domains and Linux
userspace compatibility.
Ginux is an independent kernel and operating-system prototype. It boots directly
through UEFI without a Linux host kernel, provides its own memory management,
scheduler, process model, VFS, networking, device support, and Linux ABI, and
runs an embedded administration controller in ring 3. Its primary workload is
an unmodified Linux userspace running inside a kernel-native domain .
The repository contains the complete bootable system, controller, native
services, Linux compatibility implementation, hardware paths, build tools,
architecture specifications, and QEMU test harnesses. The current development
target is an x86-64 Q35 virtual machine with OVMF, Intel IOMMU, and virtio
devices. It runs through KVM on an x86-64 Linux host and QEMU TCG on Apple
Silicon.
build as one bootable GINUX.EFI image;
start its embedded controller as an ordinary ring-3 process;
configure devices, networking, storage, and persistent domain definitions;
pull, verify, cache, and assemble OCI images into domain-local roots;
create and run multiple domains with separate process, address-space,
filesystem, mount, capability, and lifecycle state;
execute selected unmodified Alpine musl/BusyBox and Rocky Linux glibc
programs; and
boot Rocky Linux 10 systemd through multi-user.target .
Ginux is experimental. Linux ABI coverage, device support, filesystems,
networking, resource enforcement, security validation, and real-hardware
coverage are incomplete. It is useful for research, architecture experiments,
and testing—not production workloads or as a replacement for Linux.
The name is a joke in the same family as Linux: Ginux combines Linux, GPT,
and Greg. To keep the author's ego within its configured resource limit, the
G officially stands for GPT.
Linux userspace without Linux underneath. Linux is a personality exposed
by Ginux, not the host kernel.
Domains as the native system boundary. Domains own processes, address
spaces, filesystems, devices, resources, authority, and lifecycle.
A self-contained boot image. The kernel, controller, and core boot
requirements ship in one EFI release unit.
A small userspace control plane. The embedded controller configures and
observes the kernel without executing in ring 0.
Explicit, revocable authority. Kernel capabilities grant access to named
objects and are tied to domain generations for cleanup after failure.
OCI as transport, not runtime. OCI delivers and assembles filesystems;
Ginux domains provide execution and isolation.
Compatibility as an observable contract. Real Alpine and Rocky Linux
workloads supplement synthetic ABI tests.
+-----------------------------------------------------------------------+
| Embedded ring-3 controller: configuration, devices, domains, lifecycle |
+-----------------------------------------------------------------------+
| Linux-personality domains: Alpine, Rocky, and other Linux userspaces |
+-----------------------------------------------------------------------+
| Native domains: ownership | capabilities | resources | failure policy |
+-----------------------------------------------------------------------+
| Ginux kernel: processes | VFS | networking | scheduling | memory | IPC |
+-----------------------------------------------------------------------+
| UEFI firmware and the validated x86-64 QEMU hardware profile |
+-----------------------------------------------------------------------+
Boot architecture and controller
UEFI loads one GINUX.EFI image containing the kernel, boot catalog, and
authenticated controller payload. Small assembly boundaries enter and leave
UEFI, install exception, interrupt, and syscall entries, and switch privilege
and address-space state; policy and most mechanisms are implemented in Rust.
The controller is part of the single EFI release unit rather than a service
installed in a host distribution. At boot, the kernel maps the statically
linked controller into ring 3. It configures the kernel and devices, discovers
host state, manages networking and persistent storage, maintains the OCI cache,
defines domains, and controls their resources, attachments, autostart, and
lifecycle. Ordinary process behavior uses the Linux ABI; protected management
operations use a versioned device/ ioctl interface checked by the kernel.
This arrangement keeps the bootable system self-contained without running the
controller in ring 0. The controller can delegate or revoke authority granted
by the sealed boot envelope, but it cannot create new kernel authority.
A domain is a native Ginux kernel object and the outer unit of isolation,
ownership, accounting, scheduling, lifecycle, and failure containment. A
domain owns or is granted:
processes, threads, and address spaces;
a generation-tagged capability table;
IPC objects and shared-memory grants;
filesystem roots and mount views;
network endpoints and device grants; and
Domains are stronger boundaries than OCI containers or Linux namespaces. OCI
supplies an initial filesystem; Ginux supplies isolation and execution. A
single Linux-personality domain can look like a conventional computer, with
systemd as PID 1 and a distribution userspace above it. Multiple domains share
one Ginux kernel while retaining separate kernel-owned state.
A personality is the userspace contract visible inside a domain: executable
format, syscalls, processes, signals, VFS behavior, devices, terminals, clocks,
networking, and namespace expectations. Linux x86-64 is the only implemented
personality. Other personalities, such as BSD, remain architectural
possibilities rather than tested features.
Filesystems, devices, and persistence
The kernel VFS presents domain-local roots assembled from verified OCI layers,
a writable overlay, and selected tmpfs , procfs , and devfs behavior. A
persistent host-owned volume stores controller configuration and a
content-addressed OCI cache across reboots.
The validated QEMU profile includes virtio block and network paths, MSI-X, and
VT-d DMA confinement. Measured Rust driver and protocol services can run in
separate ring-3 domains with bounded MMIO, DMA, queue, and packet authority.
General driver support, production storage guarantees, broad filesystem
coverage, and a physical-hardware matrix remain future work.
Area
Demonstrated
Important boundary
Boot
Unified x86-64 UEFI image under Q35/OVMF
No production signing, recovery, or physical-hardware profile
Controller
Embedded static ring-3 administration process
No production remote management plane
Domains
Multiple domain objects with distinct process, address-space, root, mount, capability, and lifecycle state
Resource and device isolation remain incomplete in several paths
Linux userspace
Selected Alpine and Rocky programs; Rocky systemd reaches multi-user.target
No distribution-wide compatibility claim
Storage
Persistent configuration and verified OCI layer cache
Writable domain persistence and production recovery remain incomplete
Networking
Host DHCP/DNS, virtual-link configuration, loopback, and tested lifecycle
Independent per-domain packet paths and leases remain unfinished
Devices
QEMU virtio block/network, MSI-X, and VT-d paths
No GPU, NVMe, broad hardware, or stable third-party driver ABI
Performance
Matched 1/2/4-vCPU QEMU comparison using the same Rocky userspace and sysbench binary
Not a native-hardware or general kernel benchmark; four-thread scheduling remains weak
Security
Capability and generation mechanisms exercised by tests
No independent audit or production security claim
The validated development target is QEMU's x86-64 Q35 machine with OVMF,
Intel IOMMU, and virtio devices. QEMU uses KVM on a suitable x86-64 Linux host
and TCG on Apple Silicon Macs.
The Linux 6.12 x86-64 ledger classifies all 375 syscall numbers: 82 complete,
245 partial, and 48 policy-denied. Classification means every syscall number has
an explicit disposition; it does not mean complete Linux conformance. Loader,
process, VFS, terminal, socket, IPC, namespace, cgroup, pseudo-filesystem, and
device behavior is implemented only to the extent recorded in the
compatibility matrix .
The project uses real workloads as compound integration tests. Alpine exercises
the musl/BusyBox path. Rocky exercises glibc, dynamic linking, fork/exec, systemd,
services, procfs, terminals, and related kernel interfaces. A successful tested
workflow is evidence for that workflow only; it is not a guarantee that an
arbitrary Linux binary or distribution will run.
Clang with the x86_64-none-elf target;
QEMU with x86-64 and OVMF support; and
Check the host, build Ginux, and boot it:
make doctor
make run
make run creates a sparse development volume at build/qemu/volume.raw .
make reset-storage deletes it; the next boot recreates a clean volume.
At the ginux> prompt, a minimal Alpine workflow is:
domain new alpine
domain set alpine image=alpine:latest init=/bin/sh
domain start alpine --attach
Press Control-] to detach. Reattach or execute another process with:
domain attach alpine
domain exec alpine /bin/echo hello
domain exec alpine --interactive /bin/sh
To exercise the tested Rocky systemd path:
domain new rocky
domain set rocky image=rockylinux/rockylinux:10-ubi-init init=/sbin/init memory-mib=2048
domain start rocky --attach
The base rockylinux/rockylinux:10 image does not contain systemd; the tested
init-bearing image is rockylinux/rockylinux:10-ubi-init .
domain start NAME queues startup and returns to the controller. Use
domain list for transitional state and domain status NAME for detailed
root, execution, console, resource, image, and network state. OCI is only the
transport and filesystem assembly format; it is not the runtime or security
boundary.
make build # build the unified EFI image
make test-linux-runtime # ELF, loader, PIE,

[truncated]

## Original Extract

Contribute to gmkurtzer/ginux development by creating an account on GitHub.

GitHub - gmkurtzer/ginux · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gmkurtzer
ginux Public Notifications You must be signed in to change notification settings
Star 319 ( 319 ) You must be signed in to star a repository
Readme Apache-2.0 license Activity Stars
41 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,285 Commits 1,285 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills .github/ workflows .github/ workflows .hdae .hdae boot boot control-plane control-plane crates crates docs docs hdae hdae kernel kernel profiles profiles services services specs specs tests tests tools tools .gitignore .gitignore AGENTS.md AGENTS.md ARCHITECTURE.md ARCHITECTURE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
A Rust-first x86-64 operating system built around native domains and Linux
userspace compatibility.
Ginux is an independent kernel and operating-system prototype. It boots directly
through UEFI without a Linux host kernel, provides its own memory management,
scheduler, process model, VFS, networking, device support, and Linux ABI, and
runs an embedded administration controller in ring 3. Its primary workload is
an unmodified Linux userspace running inside a kernel-native domain .
The repository contains the complete bootable system, controller, native
services, Linux compatibility implementation, hardware paths, build tools,
architecture specifications, and QEMU test harnesses. The current development
target is an x86-64 Q35 virtual machine with OVMF, Intel IOMMU, and virtio
devices. It runs through KVM on an x86-64 Linux host and QEMU TCG on Apple
Silicon.
build as one bootable GINUX.EFI image;
start its embedded controller as an ordinary ring-3 process;
configure devices, networking, storage, and persistent domain definitions;
pull, verify, cache, and assemble OCI images into domain-local roots;
create and run multiple domains with separate process, address-space,
filesystem, mount, capability, and lifecycle state;
execute selected unmodified Alpine musl/BusyBox and Rocky Linux glibc
programs; and
boot Rocky Linux 10 systemd through multi-user.target .
Ginux is experimental. Linux ABI coverage, device support, filesystems,
networking, resource enforcement, security validation, and real-hardware
coverage are incomplete. It is useful for research, architecture experiments,
and testing—not production workloads or as a replacement for Linux.
The name is a joke in the same family as Linux: Ginux combines Linux, GPT,
and Greg. To keep the author's ego within its configured resource limit, the
G officially stands for GPT.
Linux userspace without Linux underneath. Linux is a personality exposed
by Ginux, not the host kernel.
Domains as the native system boundary. Domains own processes, address
spaces, filesystems, devices, resources, authority, and lifecycle.
A self-contained boot image. The kernel, controller, and core boot
requirements ship in one EFI release unit.
A small userspace control plane. The embedded controller configures and
observes the kernel without executing in ring 0.
Explicit, revocable authority. Kernel capabilities grant access to named
objects and are tied to domain generations for cleanup after failure.
OCI as transport, not runtime. OCI delivers and assembles filesystems;
Ginux domains provide execution and isolation.
Compatibility as an observable contract. Real Alpine and Rocky Linux
workloads supplement synthetic ABI tests.
+-----------------------------------------------------------------------+
| Embedded ring-3 controller: configuration, devices, domains, lifecycle |
+-----------------------------------------------------------------------+
| Linux-personality domains: Alpine, Rocky, and other Linux userspaces |
+-----------------------------------------------------------------------+
| Native domains: ownership | capabilities | resources | failure policy |
+-----------------------------------------------------------------------+
| Ginux kernel: processes | VFS | networking | scheduling | memory | IPC |
+-----------------------------------------------------------------------+
| UEFI firmware and the validated x86-64 QEMU hardware profile |
+-----------------------------------------------------------------------+
Boot architecture and controller
UEFI loads one GINUX.EFI image containing the kernel, boot catalog, and
authenticated controller payload. Small assembly boundaries enter and leave
UEFI, install exception, interrupt, and syscall entries, and switch privilege
and address-space state; policy and most mechanisms are implemented in Rust.
The controller is part of the single EFI release unit rather than a service
installed in a host distribution. At boot, the kernel maps the statically
linked controller into ring 3. It configures the kernel and devices, discovers
host state, manages networking and persistent storage, maintains the OCI cache,
defines domains, and controls their resources, attachments, autostart, and
lifecycle. Ordinary process behavior uses the Linux ABI; protected management
operations use a versioned device/ ioctl interface checked by the kernel.
This arrangement keeps the bootable system self-contained without running the
controller in ring 0. The controller can delegate or revoke authority granted
by the sealed boot envelope, but it cannot create new kernel authority.
A domain is a native Ginux kernel object and the outer unit of isolation,
ownership, accounting, scheduling, lifecycle, and failure containment. A
domain owns or is granted:
processes, threads, and address spaces;
a generation-tagged capability table;
IPC objects and shared-memory grants;
filesystem roots and mount views;
network endpoints and device grants; and
Domains are stronger boundaries than OCI containers or Linux namespaces. OCI
supplies an initial filesystem; Ginux supplies isolation and execution. A
single Linux-personality domain can look like a conventional computer, with
systemd as PID 1 and a distribution userspace above it. Multiple domains share
one Ginux kernel while retaining separate kernel-owned state.
A personality is the userspace contract visible inside a domain: executable
format, syscalls, processes, signals, VFS behavior, devices, terminals, clocks,
networking, and namespace expectations. Linux x86-64 is the only implemented
personality. Other personalities, such as BSD, remain architectural
possibilities rather than tested features.
Filesystems, devices, and persistence
The kernel VFS presents domain-local roots assembled from verified OCI layers,
a writable overlay, and selected tmpfs , procfs , and devfs behavior. A
persistent host-owned volume stores controller configuration and a
content-addressed OCI cache across reboots.
The validated QEMU profile includes virtio block and network paths, MSI-X, and
VT-d DMA confinement. Measured Rust driver and protocol services can run in
separate ring-3 domains with bounded MMIO, DMA, queue, and packet authority.
General driver support, production storage guarantees, broad filesystem
coverage, and a physical-hardware matrix remain future work.
Area
Demonstrated
Important boundary
Boot
Unified x86-64 UEFI image under Q35/OVMF
No production signing, recovery, or physical-hardware profile
Controller
Embedded static ring-3 administration process
No production remote management plane
Domains
Multiple domain objects with distinct process, address-space, root, mount, capability, and lifecycle state
Resource and device isolation remain incomplete in several paths
Linux userspace
Selected Alpine and Rocky programs; Rocky systemd reaches multi-user.target
No distribution-wide compatibility claim
Storage
Persistent configuration and verified OCI layer cache
Writable domain persistence and production recovery remain incomplete
Networking
Host DHCP/DNS, virtual-link configuration, loopback, and tested lifecycle
Independent per-domain packet paths and leases remain unfinished
Devices
QEMU virtio block/network, MSI-X, and VT-d paths
No GPU, NVMe, broad hardware, or stable third-party driver ABI
Performance
Matched 1/2/4-vCPU QEMU comparison using the same Rocky userspace and sysbench binary
Not a native-hardware or general kernel benchmark; four-thread scheduling remains weak
Security
Capability and generation mechanisms exercised by tests
No independent audit or production security claim
The validated development target is QEMU's x86-64 Q35 machine with OVMF,
Intel IOMMU, and virtio devices. QEMU uses KVM on a suitable x86-64 Linux host
and TCG on Apple Silicon Macs.
The Linux 6.12 x86-64 ledger classifies all 375 syscall numbers: 82 complete,
245 partial, and 48 policy-denied. Classification means every syscall number has
an explicit disposition; it does not mean complete Linux conformance. Loader,
process, VFS, terminal, socket, IPC, namespace, cgroup, pseudo-filesystem, and
device behavior is implemented only to the extent recorded in the
compatibility matrix .
The project uses real workloads as compound integration tests. Alpine exercises
the musl/BusyBox path. Rocky exercises glibc, dynamic linking, fork/exec, systemd,
services, procfs, terminals, and related kernel interfaces. A successful tested
workflow is evidence for that workflow only; it is not a guarantee that an
arbitrary Linux binary or distribution will run.
Clang with the x86_64-none-elf target;
QEMU with x86-64 and OVMF support; and
Check the host, build Ginux, and boot it:
make doctor
make run
make run creates a sparse development volume at build/qemu/volume.raw .
make reset-storage deletes it; the next boot recreates a clean volume.
At the ginux> prompt, a minimal Alpine workflow is:
domain new alpine
domain set alpine image=alpine:latest init=/bin/sh
domain start alpine --attach
Press Control-] to detach. Reattach or execute another process with:
domain attach alpine
domain exec alpine /bin/echo hello
domain exec alpine --interactive /bin/sh
To exercise the tested Rocky systemd path:
domain new rocky
domain set rocky image=rockylinux/rockylinux:10-ubi-init init=/sbin/init memory-mib=2048
domain start rocky --attach
The base rockylinux/rockylinux:10 image does not contain systemd; the tested
init-bearing image is rockylinux/rockylinux:10-ubi-init .
domain start NAME queues startup and returns to the controller. Use
domain list for transitional state and domain status NAME for detailed
root, execution, console, resource, image, and network state. OCI is only the
transport and filesystem assembly format; it is not the runtime or security
boundary.
make build # build the unified EFI image
make test-linux-runtime # ELF, loader, PIE,

[truncated]

---
source: "https://sandkiln.vercel.app/"
hn_url: "https://news.ycombinator.com/item?id=50163287"
title: "Show HN: Sandkiln – Fast Firecracker MicroVM sandboxes for AI-generated code"
article_title: "sandkiln"
image: ""
author: "sumitkumar-17"
captured_at: "2026-10-20T09:11:50Z"
capture_tool: "hn-digest"
hn_id: 50163287
score: 1
comments: 0
posted_at: "2026-10-20T09:09:46Z"
tags:
  - hacker-news
---

# Show HN: Sandkiln – Fast Firecracker MicroVM sandboxes for AI-generated code

- HN: [50163287](https://news.ycombinator.com/item?id=50163287)
- Source: [sandkiln.vercel.app](https://sandkiln.vercel.app/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T09:09:46Z

## Translation

Title: Show HN: Sandkiln – Fast Firecracker MicroVM sandboxes for AI-generated code
Article title: sandkiln
Description: A compute primitive for safely running untrusted or AI-generated code — hardware-isolated Firecracker microVMs, real benchmarks, SDKs for JS/TS, Python, and a CLI.

Article text:
"> sandkiln . Architecture Performance Roadmap Changelog Docs Architecture Performance Roadmap Changelog Docs Run untrusted code on hardware you control .
Every execution gets its own Firecracker microVM. Its own kernel, its own filesystem, its own network. Not a container with extra steps.
A Rust VMM, a static guest agent on vsock, an HTTP daemon driving both. Self-hosted, MIT licensed, benchmarked in the open.
[ 0.000000] Linux version 5.10.223 (gcc 11.4.0) #1 SMP [ 0.000000] Command line: console=ttyS0 reboot=k panic=1 [ 0.000000] Hypervisor detected: KVM [ 0.005803] Booting paravirtualized kernel on KVM [ 0.046087] smp: Brought up 1 node, 2 CPUs [ 0.214030] virtio_blk virtio0: [vda] 300 MiB [ 0.259624] VFS: Mounted root (ext4) on device 254:0 [ 0.279281] Run /sbin/init as init process [ 0.302006] systemd[1]: systemd 249.11 running Starting sandkiln-agent.service... [ OK ] Started sandkiln-agent.service Ubuntu 22.04.5 LTS ubuntu-fc-uvm ttyS0 ubuntu-fc-uvm login: root (automatic login) cold boot 10.5–10.9 ms criterion, spawn to InstanceStart exec round-trip 220–250 µs criterion, over an open vsock resume from snapshot 6.5–7.7 ms criterion full create 144 ms mean of 20 isolated cold creates Every track is the same log axis, 100µs to 1s, one tick per decade — so the marks compare directly. One dev box, single node.
Containers share a kernel. Sandboxes can't afford to .
Agent output, user uploads, third-party scripts. Running code you didn't write beside your own systems is a bad bet when the boundary is a namespace. A compromised container can still be a compromised kernel.
Everything a sandbox needs, marked by what's actually true .
Three marks, because two would lie. A capability that runs but isn't reachable from every client says so here rather than hiding behind a filled square.
Every item above, with what it does and exactly where it stops, is on the roadmap .
A thin, fully-typed client. Nothing more .
The daemon's HTTP API is the source of truth. Every SDK method is a thin wrapper over it, verified against a live daemon rather than designed in isolation — so the same four calls read the same way in all three.
import { Sandbox } from "sandkiln" ;
// Boots a real Firecracker microVM — own kernel, own filesystem, own network.
const sandbox = await Sandbox . create ({
tags: { env: "ci" , owner: "pipeline" },
});
const result = await sandbox. runCommand ( "python3" , [ "analyze.py" ]);
console. log (result.stdout, result.exitCode);
await sandbox. stop (); // preserved as a resumable snapshot by default from sandkiln import Sandbox
# Boots a real Firecracker microVM — own kernel, own filesystem, own network.
sandbox = Sandbox . create (tags={ "env" : "ci" , "owner" : "pipeline" })
result = sandbox. run_command ( "python3" , [ "analyze.py" ])
print (result.stdout, result.exit_code)
sandbox. stop () # preserved as a resumable snapshot by default npm install -g sandkiln-cli # installs the `kiln` command
# Boots a real Firecracker microVM — own kernel, own filesystem, own network.
kiln sandbox create --tag env=ci --tag owner=pipeline
kiln sandbox exec <id> python3 analyze.py
kiln sandbox ls --tag env=ci
kiln sandbox rm <id> # preserves state as a snapshot by default Quickstarts for JS/TS , Python and the CLI , or the daemon HTTP API that every one of them wraps.
The parts that don't fit on a homepage .

## Original Extract

A compute primitive for safely running untrusted or AI-generated code — hardware-isolated Firecracker microVMs, real benchmarks, SDKs for JS/TS, Python, and a CLI.

"> sandkiln . Architecture Performance Roadmap Changelog Docs Architecture Performance Roadmap Changelog Docs Run untrusted code on hardware you control .
Every execution gets its own Firecracker microVM. Its own kernel, its own filesystem, its own network. Not a container with extra steps.
A Rust VMM, a static guest agent on vsock, an HTTP daemon driving both. Self-hosted, MIT licensed, benchmarked in the open.
[ 0.000000] Linux version 5.10.223 (gcc 11.4.0) #1 SMP [ 0.000000] Command line: console=ttyS0 reboot=k panic=1 [ 0.000000] Hypervisor detected: KVM [ 0.005803] Booting paravirtualized kernel on KVM [ 0.046087] smp: Brought up 1 node, 2 CPUs [ 0.214030] virtio_blk virtio0: [vda] 300 MiB [ 0.259624] VFS: Mounted root (ext4) on device 254:0 [ 0.279281] Run /sbin/init as init process [ 0.302006] systemd[1]: systemd 249.11 running Starting sandkiln-agent.service... [ OK ] Started sandkiln-agent.service Ubuntu 22.04.5 LTS ubuntu-fc-uvm ttyS0 ubuntu-fc-uvm login: root (automatic login) cold boot 10.5–10.9 ms criterion, spawn to InstanceStart exec round-trip 220–250 µs criterion, over an open vsock resume from snapshot 6.5–7.7 ms criterion full create 144 ms mean of 20 isolated cold creates Every track is the same log axis, 100µs to 1s, one tick per decade — so the marks compare directly. One dev box, single node.
Containers share a kernel. Sandboxes can't afford to .
Agent output, user uploads, third-party scripts. Running code you didn't write beside your own systems is a bad bet when the boundary is a namespace. A compromised container can still be a compromised kernel.
Everything a sandbox needs, marked by what's actually true .
Three marks, because two would lie. A capability that runs but isn't reachable from every client says so here rather than hiding behind a filled square.
Every item above, with what it does and exactly where it stops, is on the roadmap .
A thin, fully-typed client. Nothing more .
The daemon's HTTP API is the source of truth. Every SDK method is a thin wrapper over it, verified against a live daemon rather than designed in isolation — so the same four calls read the same way in all three.
import { Sandbox } from "sandkiln" ;
// Boots a real Firecracker microVM — own kernel, own filesystem, own network.
const sandbox = await Sandbox . create ({
tags: { env: "ci" , owner: "pipeline" },
});
const result = await sandbox. runCommand ( "python3" , [ "analyze.py" ]);
console. log (result.stdout, result.exitCode);
await sandbox. stop (); // preserved as a resumable snapshot by default from sandkiln import Sandbox
# Boots a real Firecracker microVM — own kernel, own filesystem, own network.
sandbox = Sandbox . create (tags={ "env" : "ci" , "owner" : "pipeline" })
result = sandbox. run_command ( "python3" , [ "analyze.py" ])
print (result.stdout, result.exit_code)
sandbox. stop () # preserved as a resumable snapshot by default npm install -g sandkiln-cli # installs the `kiln` command
# Boots a real Firecracker microVM — own kernel, own filesystem, own network.
kiln sandbox create --tag env=ci --tag owner=pipeline
kiln sandbox exec <id> python3 analyze.py
kiln sandbox ls --tag env=ci
kiln sandbox rm <id> # preserves state as a snapshot by default Quickstarts for JS/TS , Python and the CLI , or the daemon HTTP API that every one of them wraps.
The parts that don't fit on a homepage .

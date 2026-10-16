---
source: "https://overalls.dev/"
hn_url: "https://news.ycombinator.com/item?id=50119092"
title: "Overalls: An AI-first developer platform that you own"
article_title: "Overalls: an AI-first developer platform that you own"
image: "https://overalls.dev/assets/og.png"
author: "simonpure"
captured_at: "2026-10-16T15:15:17Z"
capture_tool: "hn-digest"
hn_id: 50119092
score: 2
comments: 0
posted_at: "2026-10-16T14:52:34Z"
tags:
  - hacker-news
---

# Overalls: An AI-first developer platform that you own

- HN: [50119092](https://news.ycombinator.com/item?id=50119092)
- Source: [overalls.dev](https://overalls.dev/)
- Score: 2
- Comments: 0
- Posted: 2026-10-16T14:52:34Z

## Translation

Title: Overalls: An AI-first developer platform that you own
Article title: Overalls: an AI-first developer platform that you own
Description: Overalls is an open-source, self-hosted platform for persistent Firecracker development machines and the AI agents that work on them.

Article text:
Skip to content
RESEARCH PREVIEW
Overalls launches on 3 November 2026 . Get notified →
An AI-first developer platform that you own.
An AI-first developer platform that you own.
Overalls is a self-hosted platform for development machines and the AI agents that work on them. Every machine is a persistent Firecracker microVM running NixOS, Ubuntu or Windows, and it keeps its disk, name and address when you stop, start, resize, clone or move it.
Overalls has its own agent built in: start a thread on any machine from the browser or the CLI, and it keeps working after you close your laptop, waking the machine if it was stopped. It's the Amp orbs workflow on hardware you own.
Claude Code and Codex come pre-installed too, with no API keys on the machine and every token accounted for. Open any repository, branch or pull request in a fresh machine with one click, share it with your coworkers, and work from the browser, SSH, VS Code or the overalls CLI. Put hosts in every region your team works in.
It's open source under the MIT licence. There is no hosted version to sign up for: you install it and run it yourself.
╱╲ ╱╲ ╱╲ ╱╲
╱ ╲───╱ ╲ ╱ ╲───╱ ╲
│ ● ● │ │ ● ● │
│ ═ ▾ ═ │ │ ═ ▾ ═ │
│ ═ ╰┴╯ ═ │ │ ═ ╰┴╯ ═ │
╲─┬─────┬─╱ ╲─┬─────┬─╱
┌──■─────■───────■─────■──┐
╭┤ ┌─────────┐ ├╮
││ ├─────────┤ ││
●│ └─────────┘ │●
┌─┴─────────────────────────┴─┐
│■ ■│
│ ┌───┐ │ ╭─╮
│ │ │ │ │ │
│ │ │ ├───╯ │
└────────────┘ └────────────┘─────╯
╰──●─●──╯ ╰──●─●──╯
made to get dirty
What you get
STABLE built and covered by acceptance tests · PREVIEW works, with rough edges · EXPERIMENTAL still being built
Each machine is a Firecracker guest with its own kernel and a ZFS disk. Stopping it keeps the disk. Disks go up to 500 GiB.
Clone a machine into an independent copy. Resize CPU and memory up or down, and grow the disk. Move a stopped machine to another host; it keeps its name, identity and address.
Linux guests come with Podman (Docker-compatible), git and the GitHub CLI, tmux, Neovim, ripgrep, jq, direnv and devenv installed.
Guest downloads from apt, Nix, pip, npm and the like go through Brady. It fetches only from verified upstreams and caches public, anonymous responses, so the second machine to install a package gets it from the cache.
For customer material: no outbound network, no sharing, no web exposure, and SSH for the owner only.
The default guest. Declarative and reproducible, with Nix and devenv ready for per-project environments. You can build your own variant with extra NixOS modules.
Built with Nix from pinned package snapshots, so every build of the image is identical. Add your own packages with a lockfile.
Booted from firmware under a patched Firecracker, with guest agents for readiness, SSH and the terminal.
A RHEL-compatible guest for teams that build, package and test against Red Hat Enterprise Linux and its rebuilds. Being built now.
FreeBSD guests, so BSD software gets developed and tested on a real FreeBSD kernel rather than emulated in Linux. Being built now.
illumos guests, the open-source descendant of OpenSolaris, for work on Solaris-heritage systems. Being built now.
Run many persistent terminals per machine from a browser tab. People you share with start in watch mode and can take control and hand it back.
The overalls client runs on macOS (Apple Silicon and Intel), Linux (x86_64 and arm64) and Windows. A one-line installer sets up SSH, and you can connect to any machine, whether it runs NixOS, Ubuntu or Windows.
ssh <gateway> new , ls , start , stop and rm , with --json and retry-safe request IDs for scripts. One wildcard SSH profile connects you to every guest.
Each machine can serve HTTPS on its own subdomain under apps.<your-domain> . Audi runs DNS and ingress for this.
A personal editor built from Microsoft's official Code OSS source, with Open VSX extensions, running as you inside the guest. It is verified on NixOS; Ubuntu support is partial.
Give a coworker access to one machine with an access-control role. Viewers watch the terminal, operators can also take control of it, and managers can also start, stop, delete and re-share the machine. The owner can't be removed. Cedar checks every action when it runs, so revoking access takes effect even on pages that are already open.
Owners, admins and members share a pool of machines, integrations and secrets, so a team's environments don't belong to whoever happened to create them.
https://<your-domain>/#https://github.com/OWNER/REPO creates a new machine with that repo cloned into /workspace/REPO . Add /tree/BRANCH , /commit/SHA or /pull/N to start from exactly that code. If you aren't signed in, the launch picks up again after you sign in.
The Gitpod workflow on your own hardware. Click the toolbar button on a GitHub, GitLab, Codeberg or Bitbucket Cloud page (a repository, branch, commit or pull request) and you get a new machine with that code already cloned. Use one machine per piece of work and delete it when you're done. It works in Chrome, Firefox and Chromium browsers such as Edge, Brave and Arc, and sends only the page's origin and path.
Pre-installed on every guest image that supports them, and ready to use. They're already pointed at Ziggy, so there's nothing to configure: no API keys or credentials are installed on the machine.
Guests call https://llm.internal for OpenAI and Anthropic models. Ziggy adds the provider credentials outside the guest and records usage per account, so no API key ever lives in a VM.
Inference usage is recorded for each user, down to the machine and agent thread. Everyone can keep an eye on their own usage. Administrators get fleet-wide reports on AI adoption, utilisation and estimated cost, and can export them to reconcile against their inference provider's bills. Usage that wasn't measured shows as unknown, never as zero.
Named integrations at <name>.integration.internal let guests call external APIs and peer VMs. Galileo adds the credentials at the edge, so a machine can use a secret without being able to read it.
Overalls' built-in agent works in threads on each machine. Sending a prompt wakes a stopped machine, the agent's tools run inside the guest, and the work carries on after you close your laptop. It's the Amp orbs workflow on your own hardware. You can also prompt from your terminal after approving it in the browser.
Keep the control plane in one region and put Lilly hosts in Washington DC, London, Singapore or wherever your team works.
Run multiple Marshall instances behind HAProxy against a shared PostgreSQL, and as many Lilly hosts as you need. If a Marshall instance is lost, running machines keep running and terminal sessions and sharing recover on the others. Each host keeps itself up to date from Git.
Take a Lilly host out of service without losing anyone's work. Evacuation moves its machines to other Lilly hosts one at a time; each keeps its identity, disk, sharing and address, and you approve any long downtime first. A drain stops machines in place instead. When the host is empty you can bring it back or retire it.
Restic backs up to any S3-compatible bucket: guest disks from native ZFS snapshots, a dump of Marshall's PostgreSQL, host identities and keys. Running guests sync first and are never paused. Built-in checks read the backups back to prove they restore.
A 16 vCPU guest that runs a whole Overalls stack inside itself, including nested Firecracker guests. Useful for developing Overalls within Overalls.
Marshall is the web app and control plane, a Phoenix and Ash application backed by PostgreSQL. Lilly runs on each host. It starts Firecracker guests on ZFS and connects to Marshall over HTTPS to report capacity and receive commands. Porsche handles management over SSH, and Audi serves DNS, guest SSH and web apps. Each guest gets an isolated network, and traffic between guests is blocked.
┌───────────────┐ ◀──▶ PostgreSQL
browser ─HTTPS─▶ proxy ─▶ │ Marshall │
ssh ──manage──▶ Porsche ─▶ │ control plane │
└───────▲───────┘
HTTPS, dialled out by every host
┌───────────────────────────┼───────────────────────────┐
┌──────────┴─────────┐ ┌──────────┴─────────┐ ┌──────────┴─────────┐
│ Lilly · host 1 │ │ Lilly · host 2 │ │ Lilly · host N │
│ ┌────┐┌────┐┌────┐ │ │ ┌────┐┌────┐ │ │ ┌────┐┌────┐┌────┐ │
│ │ vm ││ vm ││ vm │ │ │ │ vm ││ vm │ │ ·· │ │ vm ││ vm ││ vm │ │
│ └────┘└────┘└────┘ │ │ └────┘└────┘ │ │ └────┘└────┘└────┘ │
│ Firecracker · ZFS │ │ Firecracker · ZFS │ │ Firecracker · ZFS │
└──────────▲─────────┘ └──────────▲─────────┘ └──────────▲─────────┘
└───────────────────────────┬───────────────────────────┘
│
ssh · web ────────────────────▶ Audi ───┘ routes each request to the
host running that guest
One control plane, many hosts. Add a host and Marshall starts placing machines on it. Full architecture →
Close to your developers
Run the control plane in one region and put Lilly hosts wherever your developers are. Marshall and PostgreSQL stay together, with several Marshall instances for high availability. Lilly hosts can go anywhere, for example Washington DC, London and Singapore, and each one dials out to the control plane over HTTPS.
┌─ control plane · one region ───────────────────────────────┐
│ │
│ HAProxy ─▶ Marshall · Marshall · Marshall ◀─▶ PostgreSQL │
│ several Marshall instances for high availability │
│ │
└──────────────────────────────┬─────────────────────────────┘
▲
│ HTTPS, dialled out by every host
│
┌────────────────────────────┼────────────────────────────┐
┌──────────┴──────────┐ ┌──────────┴──────────┐ ┌──────────┴──────────┐
│ WASHINGTON DC │ │ LONDON │ │ SINGAPORE │
│ Lilly Lilly Lilly │ │ Lilly Lilly │ │ Lilly Lilly │
│ guests on ZFS │ │ guests on ZFS │ │ guests on ZFS │
│ Audi: SSH, web apps │ │ Audi: SSH, web apps │ │ Audi: SSH, web apps │
└─────────────────────┘ └─────────────────────┘ └─────────────────────┘
▲ ▲ ▲
developers in developers in developers in
the Americas Europe Asia-Pacific
Centralise the control plane; distribute the hosts.
There are two ways into a machine. SSH and shared web apps go direct : ssh , VS Code over Remote-SSH, and the web apps a machine shares connect to Audi on the machine's own host, so they stay in the region and get local latency. The web interface goes through the control plane : the browser terminal and the browser editor are relayed by Marshall, so they cross regions when the control plane is far away.
┌─ control plane · one region ─────────────────────────────────────────┐
│ HAProxy ─▶ Marshall · Marshall · Marshall ◀─▶ PostgreSQL │
└───────▲──────────────────────────────────────────────────────▲───────┘
│ web terminal and relay over HTTPS, │
│ browser editor dialled out by Lilly │
│ │
┌────────────┴────────────┐ ┌──────────────┴─────────────┐
│ DEVELOPER IN SINGAPORE │ │ HOST IN SINGAPORE │
│ browser │ │ │
│ ssh, VS Code, │ direct, same region │ │
│ shared web apps ────────┼─────────────────────────▶│ Audi ─▶ Lilly ─▶ guest │
└─────────────────────────┘ └────────────────────────────┘
For the lowest latency, use SSH, VS Code or a shared web app. The browser terminal is there for when you're away from your laptop.
A taste
Every management command works in the browser, at an interactive SSH prompt or as a one-off ssh command. Add --json when a script (or agent) is reading the output.
$ ssh <gateway> new --name scratch
$ ssh <gateway> ls
$ ssh <gateway> resize scratch --cpu 4 --memory 8192
$ ssh <gateway> clone scratch --name scratch-copy --request-id copy-1 --no-wait
$ ssh <gateway> share add scratch colleague@example.com operator
$ overalls open https://github.com/OWNER/REPO
<gateway> is your own deployment's domain: Porsche listens on port 22 of your Marshall hosts. See your first machine .
Overalls is a set of NixOS modules plus tools for installing and enrolling bare-metal hosts. Your fleet is defined in a Git repository: you fork Overalls, declare your hosts, and each host keeps itself up to date from your main branch. You can start with one server and add hosts later. To make the installation y

[truncated]

## Original Extract

Overalls is an open-source, self-hosted platform for persistent Firecracker development machines and the AI agents that work on them.

Skip to content
RESEARCH PREVIEW
Overalls launches on 3 November 2026 . Get notified →
An AI-first developer platform that you own.
An AI-first developer platform that you own.
Overalls is a self-hosted platform for development machines and the AI agents that work on them. Every machine is a persistent Firecracker microVM running NixOS, Ubuntu or Windows, and it keeps its disk, name and address when you stop, start, resize, clone or move it.
Overalls has its own agent built in: start a thread on any machine from the browser or the CLI, and it keeps working after you close your laptop, waking the machine if it was stopped. It's the Amp orbs workflow on hardware you own.
Claude Code and Codex come pre-installed too, with no API keys on the machine and every token accounted for. Open any repository, branch or pull request in a fresh machine with one click, share it with your coworkers, and work from the browser, SSH, VS Code or the overalls CLI. Put hosts in every region your team works in.
It's open source under the MIT licence. There is no hosted version to sign up for: you install it and run it yourself.
╱╲ ╱╲ ╱╲ ╱╲
╱ ╲───╱ ╲ ╱ ╲───╱ ╲
│ ● ● │ │ ● ● │
│ ═ ▾ ═ │ │ ═ ▾ ═ │
│ ═ ╰┴╯ ═ │ │ ═ ╰┴╯ ═ │
╲─┬─────┬─╱ ╲─┬─────┬─╱
┌──■─────■───────■─────■──┐
╭┤ ┌─────────┐ ├╮
││ ├─────────┤ ││
●│ └─────────┘ │●
┌─┴─────────────────────────┴─┐
│■ ■│
│ ┌───┐ │ ╭─╮
│ │ │ │ │ │
│ │ │ ├───╯ │
└────────────┘ └────────────┘─────╯
╰──●─●──╯ ╰──●─●──╯
made to get dirty
What you get
STABLE built and covered by acceptance tests · PREVIEW works, with rough edges · EXPERIMENTAL still being built
Each machine is a Firecracker guest with its own kernel and a ZFS disk. Stopping it keeps the disk. Disks go up to 500 GiB.
Clone a machine into an independent copy. Resize CPU and memory up or down, and grow the disk. Move a stopped machine to another host; it keeps its name, identity and address.
Linux guests come with Podman (Docker-compatible), git and the GitHub CLI, tmux, Neovim, ripgrep, jq, direnv and devenv installed.
Guest downloads from apt, Nix, pip, npm and the like go through Brady. It fetches only from verified upstreams and caches public, anonymous responses, so the second machine to install a package gets it from the cache.
For customer material: no outbound network, no sharing, no web exposure, and SSH for the owner only.
The default guest. Declarative and reproducible, with Nix and devenv ready for per-project environments. You can build your own variant with extra NixOS modules.
Built with Nix from pinned package snapshots, so every build of the image is identical. Add your own packages with a lockfile.
Booted from firmware under a patched Firecracker, with guest agents for readiness, SSH and the terminal.
A RHEL-compatible guest for teams that build, package and test against Red Hat Enterprise Linux and its rebuilds. Being built now.
FreeBSD guests, so BSD software gets developed and tested on a real FreeBSD kernel rather than emulated in Linux. Being built now.
illumos guests, the open-source descendant of OpenSolaris, for work on Solaris-heritage systems. Being built now.
Run many persistent terminals per machine from a browser tab. People you share with start in watch mode and can take control and hand it back.
The overalls client runs on macOS (Apple Silicon and Intel), Linux (x86_64 and arm64) and Windows. A one-line installer sets up SSH, and you can connect to any machine, whether it runs NixOS, Ubuntu or Windows.
ssh <gateway> new , ls , start , stop and rm , with --json and retry-safe request IDs for scripts. One wildcard SSH profile connects you to every guest.
Each machine can serve HTTPS on its own subdomain under apps.<your-domain> . Audi runs DNS and ingress for this.
A personal editor built from Microsoft's official Code OSS source, with Open VSX extensions, running as you inside the guest. It is verified on NixOS; Ubuntu support is partial.
Give a coworker access to one machine with an access-control role. Viewers watch the terminal, operators can also take control of it, and managers can also start, stop, delete and re-share the machine. The owner can't be removed. Cedar checks every action when it runs, so revoking access takes effect even on pages that are already open.
Owners, admins and members share a pool of machines, integrations and secrets, so a team's environments don't belong to whoever happened to create them.
https://<your-domain>/#https://github.com/OWNER/REPO creates a new machine with that repo cloned into /workspace/REPO . Add /tree/BRANCH , /commit/SHA or /pull/N to start from exactly that code. If you aren't signed in, the launch picks up again after you sign in.
The Gitpod workflow on your own hardware. Click the toolbar button on a GitHub, GitLab, Codeberg or Bitbucket Cloud page (a repository, branch, commit or pull request) and you get a new machine with that code already cloned. Use one machine per piece of work and delete it when you're done. It works in Chrome, Firefox and Chromium browsers such as Edge, Brave and Arc, and sends only the page's origin and path.
Pre-installed on every guest image that supports them, and ready to use. They're already pointed at Ziggy, so there's nothing to configure: no API keys or credentials are installed on the machine.
Guests call https://llm.internal for OpenAI and Anthropic models. Ziggy adds the provider credentials outside the guest and records usage per account, so no API key ever lives in a VM.
Inference usage is recorded for each user, down to the machine and agent thread. Everyone can keep an eye on their own usage. Administrators get fleet-wide reports on AI adoption, utilisation and estimated cost, and can export them to reconcile against their inference provider's bills. Usage that wasn't measured shows as unknown, never as zero.
Named integrations at <name>.integration.internal let guests call external APIs and peer VMs. Galileo adds the credentials at the edge, so a machine can use a secret without being able to read it.
Overalls' built-in agent works in threads on each machine. Sending a prompt wakes a stopped machine, the agent's tools run inside the guest, and the work carries on after you close your laptop. It's the Amp orbs workflow on your own hardware. You can also prompt from your terminal after approving it in the browser.
Keep the control plane in one region and put Lilly hosts in Washington DC, London, Singapore or wherever your team works.
Run multiple Marshall instances behind HAProxy against a shared PostgreSQL, and as many Lilly hosts as you need. If a Marshall instance is lost, running machines keep running and terminal sessions and sharing recover on the others. Each host keeps itself up to date from Git.
Take a Lilly host out of service without losing anyone's work. Evacuation moves its machines to other Lilly hosts one at a time; each keeps its identity, disk, sharing and address, and you approve any long downtime first. A drain stops machines in place instead. When the host is empty you can bring it back or retire it.
Restic backs up to any S3-compatible bucket: guest disks from native ZFS snapshots, a dump of Marshall's PostgreSQL, host identities and keys. Running guests sync first and are never paused. Built-in checks read the backups back to prove they restore.
A 16 vCPU guest that runs a whole Overalls stack inside itself, including nested Firecracker guests. Useful for developing Overalls within Overalls.
Marshall is the web app and control plane, a Phoenix and Ash application backed by PostgreSQL. Lilly runs on each host. It starts Firecracker guests on ZFS and connects to Marshall over HTTPS to report capacity and receive commands. Porsche handles management over SSH, and Audi serves DNS, guest SSH and web apps. Each guest gets an isolated network, and traffic between guests is blocked.
┌───────────────┐ ◀──▶ PostgreSQL
browser ─HTTPS─▶ proxy ─▶ │ Marshall │
ssh ──manage──▶ Porsche ─▶ │ control plane │
└───────▲───────┘
HTTPS, dialled out by every host
┌───────────────────────────┼───────────────────────────┐
┌──────────┴─────────┐ ┌──────────┴─────────┐ ┌──────────┴─────────┐
│ Lilly · host 1 │ │ Lilly · host 2 │ │ Lilly · host N │
│ ┌────┐┌────┐┌────┐ │ │ ┌────┐┌────┐ │ │ ┌────┐┌────┐┌────┐ │
│ │ vm ││ vm ││ vm │ │ │ │ vm ││ vm │ │ ·· │ │ vm ││ vm ││ vm │ │
│ └────┘└────┘└────┘ │ │ └────┘└────┘ │ │ └────┘└────┘└────┘ │
│ Firecracker · ZFS │ │ Firecracker · ZFS │ │ Firecracker · ZFS │
└──────────▲─────────┘ └──────────▲─────────┘ └──────────▲─────────┘
└───────────────────────────┬───────────────────────────┘
│
ssh · web ────────────────────▶ Audi ───┘ routes each request to the
host running that guest
One control plane, many hosts. Add a host and Marshall starts placing machines on it. Full architecture →
Close to your developers
Run the control plane in one region and put Lilly hosts wherever your developers are. Marshall and PostgreSQL stay together, with several Marshall instances for high availability. Lilly hosts can go anywhere, for example Washington DC, London and Singapore, and each one dials out to the control plane over HTTPS.
┌─ control plane · one region ───────────────────────────────┐
│ │
│ HAProxy ─▶ Marshall · Marshall · Marshall ◀─▶ PostgreSQL │
│ several Marshall instances for high availability │
│ │
└──────────────────────────────┬─────────────────────────────┘
▲
│ HTTPS, dialled out by every host
│
┌────────────────────────────┼────────────────────────────┐
┌──────────┴──────────┐ ┌──────────┴──────────┐ ┌──────────┴──────────┐
│ WASHINGTON DC │ │ LONDON │ │ SINGAPORE │
│ Lilly Lilly Lilly │ │ Lilly Lilly │ │ Lilly Lilly │
│ guests on ZFS │ │ guests on ZFS │ │ guests on ZFS │
│ Audi: SSH, web apps │ │ Audi: SSH, web apps │ │ Audi: SSH, web apps │
└─────────────────────┘ └─────────────────────┘ └─────────────────────┘
▲ ▲ ▲
developers in developers in developers in
the Americas Europe Asia-Pacific
Centralise the control plane; distribute the hosts.
There are two ways into a machine. SSH and shared web apps go direct : ssh , VS Code over Remote-SSH, and the web apps a machine shares connect to Audi on the machine's own host, so they stay in the region and get local latency. The web interface goes through the control plane : the browser terminal and the browser editor are relayed by Marshall, so they cross regions when the control plane is far away.
┌─ control plane · one region ─────────────────────────────────────────┐
│ HAProxy ─▶ Marshall · Marshall · Marshall ◀─▶ PostgreSQL │
└───────▲──────────────────────────────────────────────────────▲───────┘
│ web terminal and relay over HTTPS, │
│ browser editor dialled out by Lilly │
│ │
┌────────────┴────────────┐ ┌──────────────┴─────────────┐
│ DEVELOPER IN SINGAPORE │ │ HOST IN SINGAPORE │
│ browser │ │ │
│ ssh, VS Code, │ direct, same region │ │
│ shared web apps ────────┼─────────────────────────▶│ Audi ─▶ Lilly ─▶ guest │
└─────────────────────────┘ └────────────────────────────┘
For the lowest latency, use SSH, VS Code or a shared web app. The browser terminal is there for when you're away from your laptop.
A taste
Every management command works in the browser, at an interactive SSH prompt or as a one-off ssh command. Add --json when a script (or agent) is reading the output.
$ ssh <gateway> new --name scratch
$ ssh <gateway> ls
$ ssh <gateway> resize scratch --cpu 4 --memory 8192
$ ssh <gateway> clone scratch --name scratch-copy --request-id copy-1 --no-wait
$ ssh <gateway> share add scratch colleague@example.com operator
$ overalls open https://github.com/OWNER/REPO
<gateway> is your own deployment's domain: Porsche listens on port 22 of your Marshall hosts. See your first machine .
Overalls is a set of NixOS modules plus tools for installing and enrolling bare-metal hosts. Your fleet is defined in a Git repository: you fork Overalls, declare your hosts, and each host keeps itself up to date from your main branch. You can start with one server and add hosts later. To make the installation y

[truncated]

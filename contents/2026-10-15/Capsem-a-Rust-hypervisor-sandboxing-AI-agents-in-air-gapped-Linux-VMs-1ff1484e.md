---
source: "https://capsem.org/"
hn_url: "https://news.ycombinator.com/item?id=50097811"
title: "Capsem – a Rust hypervisor sandboxing AI agents in air-gapped Linux VMs"
article_title: "Capsem -- The fastest way to ship with AI securely."
image: ""
author: "madars"
captured_at: "2026-10-15T01:14:23Z"
capture_tool: "hn-digest"
hn_id: 50097811
score: 2
comments: 0
posted_at: "2026-10-15T00:58:26Z"
tags:
  - hacker-news
---

# Capsem – a Rust hypervisor sandboxing AI agents in air-gapped Linux VMs

- HN: [50097811](https://news.ycombinator.com/item?id=50097811)
- Source: [capsem.org](https://capsem.org/)
- Score: 2
- Comments: 0
- Posted: 2026-10-15T00:58:26Z

## Translation

Title: Capsem – a Rust hypervisor sandboxing AI agents in air-gapped Linux VMs
Article title: Capsem -- The fastest way to ship with AI securely.
Description: The fastest way to ship with AI securely.

Article text:
Skip to content Capsem Features How It Works FAQ Docs Download GitHub (opens in new tab) CAPSEM
The fastest way to
ship with AI securely.
A Rust-powered hypervisor that sandboxes every AI coding agent
in its own air-gapped Linux VM. See everything, control everything.
Learn More Ready to code, out of the box
A fully configured Linux dev environment with your favorite AI agents, MCP tools, and packages pre-installed. Boot to prompt in under 10 seconds.
All major AI coding agents, pre-installed and auto-configured for sandbox mode.
Policy-checked tool access over vsock. Built-in tools plus custom MCP servers via fastmcp.
Python, Node.js, and 30+ packages ready to go. No setup required.
Active development. Here's what's next.
Demo video Security without compromise
Hardware-level sandboxing with Apple Virtualization.framework
Each agent session boots a lightweight Linux VM with a read-only rootfs, no swap, no kernel modules, no debugfs. Air-gapped networking with a dummy NIC and fake DNS ensures nothing reaches the real network without going through the MITM proxy.
Ephemeral VMs -- fresh state every session
Read-only rootfs with tmpfs workspace
No systemd, no sshd, no cron -- minimal attack surface
Screenshot placeholder INSPECTION See everything your AI agent does on the network
A transparent MITM proxy terminates TLS from the guest using per-domain minted certificates, inspects every HTTP request and response, and applies policy before forwarding to the real upstream. Full request/response bodies are logged to a per-session SQLite database.
Per-domain TLS certificate minting
Method + path policy rules per domain
Full body capture for post-hoc analysis
Screenshot placeholder CONTROL Enterprise-grade policy with profile and corp config layers
Each profile owns its enforcement, detection, MCP, plugin, and asset configuration. Corp-level config at /etc/capsem/corp.toml (MDM-distributed) locks down policy and reporting with enterprise overrides that users cannot bypass.
Profile-owned CEL enforcement rules
Detection rules with Sigma import/export
Corp config locks and reports profile behavior
Screenshot placeholder How it works
A native macOS hypervisor creates an air-gapped Linux VM for each session. All network traffic is forced through an inspecting proxy on the host.
Desktop App GUI + CLI interface MITM Proxy TLS termination + HTTP inspection Policy Engine Domain + HTTP + MCP rules Session Telemetry SQLite DB per session vsock :5001 terminal :5002 HTTPS + MCP Linux VM (air-gapped) AI Agent Claude / Gemini / Codex PTY Agent Terminal I/O over vsock Net Proxy TCP-to-vsock relay (iptables) MCP Server Tool relay over vsock Internet (via host MITM proxy only) FAQ Frequently Asked
Questions
Does Capsem work with Claude Code, Gemini CLI, and Codex?
All guest HTTPS traffic is redirected through an iptables rule to a local TCP relay, which bridges to the host via vsock. The host terminates TLS using per-domain minted certificates (signed by a static Capsem CA baked into the guest's trust store), inspects the HTTP request, applies policy, and forwards to the real upstream.
Can I customize which domains are allowed?
Open source, native macOS and Linux packages, boots in under 10 seconds.
or download the package from the stable channel
Download Package (opens in new tab) View on GitHub (opens in new tab) Requires macOS 14+ on Apple Silicon or Debian/Ubuntu Linux with KVM. Capsem uses Apple Virtualization.framework on macOS and KVM on Linux.
The fastest way to ship with AI securely.
Documentation (opens in new tab)
Release Channel (opens in new tab)
© 2026 Elie Bursztein. Open source under MIT license.

## Original Extract

The fastest way to ship with AI securely.

Skip to content Capsem Features How It Works FAQ Docs Download GitHub (opens in new tab) CAPSEM
The fastest way to
ship with AI securely.
A Rust-powered hypervisor that sandboxes every AI coding agent
in its own air-gapped Linux VM. See everything, control everything.
Learn More Ready to code, out of the box
A fully configured Linux dev environment with your favorite AI agents, MCP tools, and packages pre-installed. Boot to prompt in under 10 seconds.
All major AI coding agents, pre-installed and auto-configured for sandbox mode.
Policy-checked tool access over vsock. Built-in tools plus custom MCP servers via fastmcp.
Python, Node.js, and 30+ packages ready to go. No setup required.
Active development. Here's what's next.
Demo video Security without compromise
Hardware-level sandboxing with Apple Virtualization.framework
Each agent session boots a lightweight Linux VM with a read-only rootfs, no swap, no kernel modules, no debugfs. Air-gapped networking with a dummy NIC and fake DNS ensures nothing reaches the real network without going through the MITM proxy.
Ephemeral VMs -- fresh state every session
Read-only rootfs with tmpfs workspace
No systemd, no sshd, no cron -- minimal attack surface
Screenshot placeholder INSPECTION See everything your AI agent does on the network
A transparent MITM proxy terminates TLS from the guest using per-domain minted certificates, inspects every HTTP request and response, and applies policy before forwarding to the real upstream. Full request/response bodies are logged to a per-session SQLite database.
Per-domain TLS certificate minting
Method + path policy rules per domain
Full body capture for post-hoc analysis
Screenshot placeholder CONTROL Enterprise-grade policy with profile and corp config layers
Each profile owns its enforcement, detection, MCP, plugin, and asset configuration. Corp-level config at /etc/capsem/corp.toml (MDM-distributed) locks down policy and reporting with enterprise overrides that users cannot bypass.
Profile-owned CEL enforcement rules
Detection rules with Sigma import/export
Corp config locks and reports profile behavior
Screenshot placeholder How it works
A native macOS hypervisor creates an air-gapped Linux VM for each session. All network traffic is forced through an inspecting proxy on the host.
Desktop App GUI + CLI interface MITM Proxy TLS termination + HTTP inspection Policy Engine Domain + HTTP + MCP rules Session Telemetry SQLite DB per session vsock :5001 terminal :5002 HTTPS + MCP Linux VM (air-gapped) AI Agent Claude / Gemini / Codex PTY Agent Terminal I/O over vsock Net Proxy TCP-to-vsock relay (iptables) MCP Server Tool relay over vsock Internet (via host MITM proxy only) FAQ Frequently Asked
Questions
Does Capsem work with Claude Code, Gemini CLI, and Codex?
All guest HTTPS traffic is redirected through an iptables rule to a local TCP relay, which bridges to the host via vsock. The host terminates TLS using per-domain minted certificates (signed by a static Capsem CA baked into the guest's trust store), inspects the HTTP request, applies policy, and forwards to the real upstream.
Can I customize which domains are allowed?
Open source, native macOS and Linux packages, boots in under 10 seconds.
or download the package from the stable channel
Download Package (opens in new tab) View on GitHub (opens in new tab) Requires macOS 14+ on Apple Silicon or Debian/Ubuntu Linux with KVM. Capsem uses Apple Virtualization.framework on macOS and KVM on Linux.
The fastest way to ship with AI securely.
Documentation (opens in new tab)
Release Channel (opens in new tab)
© 2026 Elie Bursztein. Open source under MIT license.

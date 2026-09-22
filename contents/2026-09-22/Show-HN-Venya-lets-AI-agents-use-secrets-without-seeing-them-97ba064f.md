---
source: "https://github.com/tabith-llc/venya"
hn_url: "https://news.ycombinator.com/item?id=49801376"
title: "Show HN: Venya lets AI agents use secrets without seeing them"
article_title: "GitHub - tabith-llc/venya: Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. · GitHub"
image: "https://opengraph.githubassets.com/de0eabe40d6485a9fe77aee02ff8737bdfa4f1e1872797f7c852ed9c65c03221/tabith-llc/venya"
author: "tabith"
captured_at: "2026-09-22T14:23:39Z"
capture_tool: "hn-digest"
hn_id: 49801376
score: 2
comments: 0
posted_at: "2026-09-22T13:56:28Z"
tags:
  - hacker-news
---

# Show HN: Venya lets AI agents use secrets without seeing them

- HN: [49801376](https://news.ycombinator.com/item?id=49801376)
- Source: [github.com](https://github.com/tabith-llc/venya)
- Score: 2
- Comments: 0
- Posted: 2026-09-22T13:56:28Z

## Translation

Title: Show HN: Venya lets AI agents use secrets without seeing them
Article title: GitHub - tabith-llc/venya: Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. · GitHub
Description: Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. - tabith-llc/venya

Article text:
GitHub - tabith-llc/venya: Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
tabith-llc
/
venya
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
615 Commits 615 Commits Folders and files
docs docs packages packages scripts scripts systemd systemd tests/ e2e tests/ e2e .detect-secrets-baseline .detect-secrets-baseline .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md install-scripts-README.md install-scripts-README.md install-venya-cli.ps1 install-venya-cli.ps1 install-venya-cli.sh install-venya-cli.sh install-venya-core.sh install-venya-core.sh install-venya-executor.sh install-venya-executor.sh pyproject.toml pyproject.toml uninstall-venya-cli.ps1 uninstall-venya-cli.ps1 uninstall-venya-cli.sh uninstall-venya-cli.sh uninstall-venya-core.sh uninstall-venya-core.sh uninstall-venya-executor.sh uninstall-venya-executor.sh venya-common.sh venya-common.sh venya-core-requirements.txt venya-core-requirements.txt venya-executor-requirements.txt venya-executor-requirements.txt View all files Repository files navigation
AI agents: if you are an LLM agent that retrieved this README, start at docs/agents.md — how to wire up, operate safely, and what never to attempt. Humans: that file is the brief your agent should be handed.
AI agents manage infrastructure securely.
Venya is the first platform that lets AI agents execute commands on remote infrastructure using stored credentials without ever seeing those credentials . The LLM discovers what secrets exist, constructs the command, and the executor injects the credential into a sandboxed environment. Output is filtered. Every action is audited. A human authorized the session with a physical security key.
This doesn't exist anywhere else. Traditional secrets managers (HashiCorp Vault, CyberArk, cloud-native stores) store credentials — but they hand the plaintext to whatever process requests it. If you give an AI agent a Vault token, the agent can read every secret in plain text. Venya's zero-knowledge injection model means the agent never sees, handles, or can leak the credential value. It sees the result of the command — and nothing more.
Infrastructure teams are adopting AI agents (Claude Code, Cursor, autonomous coding assistants) to manage servers, deploy applications, and troubleshoot incidents. But these agents need credentials to do their work — SSH keys, API tokens, database passwords.
Today, teams solve this in one of three ways:
Every approach either exposes credentials or blocks AI adoption. Venya is the fourth option.
Human: "Install apache2 on web-server-3"
│
▼
┌──────────────────┐ ┌───────────────────┐ ┌────────────────────┐
│ AI Agent │────▶│ Venya Server │────▶│ Executor Daemon │
│ (Claude Code) │ │ │ │ (on executor host) │
│ │ │ 1. Wraps secret │ │ │
│ Never sees │ │ with sentinel │ │ 2. Unwraps in │
│ the password │ │ markers │ │ sandbox │
│ │ │ │ │ │
│ Sees: exit code │◀────│ 3. Relays via │◀────│ 4. Filters output │
│ + filtered │ │ mTLS │ │ (Rust filter) │
│ output │ │ │ │ │
└──────────────────┘ └───────────────────┘ └────────────────────┘
The human asks the AI to do something — e.g., "Install apache2 on web-server-3"
The AI discovers available resources — calls Venya's MCP tools to list executors and secrets (metadata only, never values)
The AI constructs the command — e.g., ssh bot@web-server-3 sudo apt install -y apache2
Venya handles the rest:
Server decrypts the secret and wraps it with cryptographic sentinel markers
Server relays the command + wrapped secret to the executor over mutual TLS
Executor unwraps the secret inside an isolated sbx microVM and injects it into the command
A Rust-based output filter scans stdout/stderr for any leaked secret values and replaces them with [REDACTED] markers before the AI ever sees it
The AI reads the filtered output — it sees the command succeeded, sees the package installation logs, but never sees the password
Every step is logged — the audit trail records who authorized the session, what command ran, on which executor, and when
Zero-Knowledge Secret Injection
The AI agent never touches plaintext credentials. Not in its context window. Not in transit. Not in output. The secret is decrypted server-side, wrapped with sentinel markers, relayed over mTLS, and unwrapped only inside the executor's sandboxed process. The Rust filter ensures that even if a command accidentally echoes a credential in its output, it's replaced with [REDACTED] before the AI ever sees it.
No other product does this. Existing secrets managers hand plaintext to the requesting process. Venya doesn't.
Every session begins with a physical security key press. The AI agent cannot initiate a session — only a human pressing a FIDO2 key can authorize access. Sessions expire after 4 hours. Token refresh happens automatically, and an idle-expired session renews within the 4-hour hard cap — the cap is non-negotiable: past it, a human must re-authenticate with the FIDO2 key. When the session is past the cap, the AI gets an actionable error: "Ask the user to re-authenticate."
mTLS Between Server and Executor
The Venya server communicates with executor daemons over mutual TLS. Both sides verify each other's certificates. If an executor's certificate is revoked, the server refuses to relay commands. If someone spoofs an executor, the mTLS handshake fails before any secret is transmitted.
Commands run inside an sbx microVM with deny-by-default networking. The executor reads an operator-defined allowlist ( /etc/venya/egress-allowlist.txt ) and only permits connections to approved destinations. If a compromised command tries to phone home to an attacker's server, the connection is blocked at the sandbox level. DNS is restricted to the operator's resolver.
Every command execution is logged:
Who authorized the session (FIDO2-enrolled user)
What command was executed (full command string)
What secrets were injected (secret IDs, never values)
Both human operators (via venya audit CLI) and AI agents (via the get_audit MCP tool) can query the audit log. Non-admin users see only their own events.
Venya speaks the Model Context Protocol — the open standard for connecting AI assistants to external tools. It works with Claude Code, Cursor, and any MCP-compatible client. No proprietary lock-in. No vendor-specific API.
Guarantee
How It's Enforced
AI never sees plaintext credentials
Server-side wrapping + executor-side unwrapping + Rust output filter
Sessions require human authorization
FIDO2 hardware key binding (WebAuthn)
Sessions are time-limited
4-hour hard cap, 15-minute idle window, 5-minute access tokens
Executor identity is verified
Mutual TLS with certificate chain validation
Compromised executors are blocked
Revoked certs rejected at the executor's next revocation poll
Data exfiltration is prevented
sbx sandbox with deny-by-default egress allowlisting
Every action is traceable
Append-only audit log with user, executor, command, timestamp
Secrets are encrypted at rest
AES-256 with KEK-wrapped DEK (envelope encryption)
Who Is Venya For?
Infrastructure teams who want to use AI without compromising security.
Using Claude Code, Cursor, or similar AI coding assistants
Managing fleets of servers, databases, or cloud infrastructure
Concerned about handing credentials to AI agents
Operating in regulated environments where audit trails are mandatory
Tired of the choice between "move fast with AI" and "stay secure"
Venya is currently in alpha — early access for teams who want to shape the product.
See Venya in action: Alpha Demo Guide
Artifacts (installers, tarballs, SHA-256 sidecars) are published on the Releases page . Install one-liners (core/executor: Ubuntu 24.04 only; the Workstation CLI additionally runs on Debian 13, macOS, and Windows):
# Core server (root)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-core.sh | sudo \
VENYA_SKIP_PROMPT=yes VENYA_DB_PASSWORD= < strong-db-password > bash -s
# Executor (root; enrollment token from the core admin; a Docker account is REQUIRED — sbx pulls its agent template from Docker Hub; username + API key/access token via stdin)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-executor.sh | sudo \
VENYA_SKIP_PROMPT=yes VENYA_SERVER_URL=https:// < core-host > VENYA_EXECUTOR_ID= < executor-id > \
VENYA_EXECUTOR_ENROLLMENT_TOKEN= < token > bash -s
# Workstation CLI (non-root; Ubuntu 24.04, Debian 13, or macOS — verified on macOS 26.6.2 arm64 and Debian 13)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-cli.sh | VENYA_SKIP_PROMPT=yes bash
# Workstation CLI (Windows) — machine-wide install, requires Administrator;
# standard users run the CLI afterward. Interactive desktop only (headless unsupported).
# Download install-venya-cli.ps1 from the Releases page, then run:
powershell - ExecutionPolicy Bypass - File install-venya - cli.ps1
Integrity: pin VENYA_TARBALL_SHA256 (hashes on the release page) for strict verification; unset, the installer fetches the .sha256 sidecar from the same origin as a corruption guardrail and fail-closes.
Workstation CLI config file: ~/.config/venya/config.json on Linux, ~/Library/Application Support/venya/config.json on macOS, %APPDATA%\venya\config.json on Windows. FIDO2 needs no extra setup on macOS (native IOKit HID transport, no root) or Windows (platform WebAuthn API — standard-user capable, interactive desktop required); on Linux the installer prints udev rules if /dev/hidraw* is not user-readable.
Production deployment guide: Installation Guide
Linux — Ubuntu 24.04 LTS (core/executor; tested target, installers assume it). Workstation CLI additionally supports Debian 13 and macOS (verified macOS 26.6.2 arm64)
Windows — Workstation CLI only (core and executor are Linux). Machine-wide install via install-venya-cli.ps1 requires Administrator; standard users run the CLI after install. FIDO2 ceremonies work for standard users via the platform WebAuthn API (verified Windows 11 25H2). Interactive desktop only — headless Windows is unsupported
PostgreSQL — installed automatically by the core installer (16 on Ubuntu 24.04)
Python 3.14 — pinned ( >=3.14,<3.15 ); provisioned automatically via uv
FIDO2 security key (YubiKey, SoloKeys, etc.)
Docker Sandboxes (sbx) — installed automatically by the executor in

[truncated]

## Original Extract

Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. - tabith-llc/venya

GitHub - tabith-llc/venya: Secure privileged access management with FIDO2 authentication, remote executor sandboxes, and LLM-driven operations via Model Context Protocol. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
tabith-llc
/
venya
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
615 Commits 615 Commits Folders and files
docs docs packages packages scripts scripts systemd systemd tests/ e2e tests/ e2e .detect-secrets-baseline .detect-secrets-baseline .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md install-scripts-README.md install-scripts-README.md install-venya-cli.ps1 install-venya-cli.ps1 install-venya-cli.sh install-venya-cli.sh install-venya-core.sh install-venya-core.sh install-venya-executor.sh install-venya-executor.sh pyproject.toml pyproject.toml uninstall-venya-cli.ps1 uninstall-venya-cli.ps1 uninstall-venya-cli.sh uninstall-venya-cli.sh uninstall-venya-core.sh uninstall-venya-core.sh uninstall-venya-executor.sh uninstall-venya-executor.sh venya-common.sh venya-common.sh venya-core-requirements.txt venya-core-requirements.txt venya-executor-requirements.txt venya-executor-requirements.txt View all files Repository files navigation
AI agents: if you are an LLM agent that retrieved this README, start at docs/agents.md — how to wire up, operate safely, and what never to attempt. Humans: that file is the brief your agent should be handed.
AI agents manage infrastructure securely.
Venya is the first platform that lets AI agents execute commands on remote infrastructure using stored credentials without ever seeing those credentials . The LLM discovers what secrets exist, constructs the command, and the executor injects the credential into a sandboxed environment. Output is filtered. Every action is audited. A human authorized the session with a physical security key.
This doesn't exist anywhere else. Traditional secrets managers (HashiCorp Vault, CyberArk, cloud-native stores) store credentials — but they hand the plaintext to whatever process requests it. If you give an AI agent a Vault token, the agent can read every secret in plain text. Venya's zero-knowledge injection model means the agent never sees, handles, or can leak the credential value. It sees the result of the command — and nothing more.
Infrastructure teams are adopting AI agents (Claude Code, Cursor, autonomous coding assistants) to manage servers, deploy applications, and troubleshoot incidents. But these agents need credentials to do their work — SSH keys, API tokens, database passwords.
Today, teams solve this in one of three ways:
Every approach either exposes credentials or blocks AI adoption. Venya is the fourth option.
Human: "Install apache2 on web-server-3"
│
▼
┌──────────────────┐ ┌───────────────────┐ ┌────────────────────┐
│ AI Agent │────▶│ Venya Server │────▶│ Executor Daemon │
│ (Claude Code) │ │ │ │ (on executor host) │
│ │ │ 1. Wraps secret │ │ │
│ Never sees │ │ with sentinel │ │ 2. Unwraps in │
│ the password │ │ markers │ │ sandbox │
│ │ │ │ │ │
│ Sees: exit code │◀────│ 3. Relays via │◀────│ 4. Filters output │
│ + filtered │ │ mTLS │ │ (Rust filter) │
│ output │ │ │ │ │
└──────────────────┘ └───────────────────┘ └────────────────────┘
The human asks the AI to do something — e.g., "Install apache2 on web-server-3"
The AI discovers available resources — calls Venya's MCP tools to list executors and secrets (metadata only, never values)
The AI constructs the command — e.g., ssh bot@web-server-3 sudo apt install -y apache2
Venya handles the rest:
Server decrypts the secret and wraps it with cryptographic sentinel markers
Server relays the command + wrapped secret to the executor over mutual TLS
Executor unwraps the secret inside an isolated sbx microVM and injects it into the command
A Rust-based output filter scans stdout/stderr for any leaked secret values and replaces them with [REDACTED] markers before the AI ever sees it
The AI reads the filtered output — it sees the command succeeded, sees the package installation logs, but never sees the password
Every step is logged — the audit trail records who authorized the session, what command ran, on which executor, and when
Zero-Knowledge Secret Injection
The AI agent never touches plaintext credentials. Not in its context window. Not in transit. Not in output. The secret is decrypted server-side, wrapped with sentinel markers, relayed over mTLS, and unwrapped only inside the executor's sandboxed process. The Rust filter ensures that even if a command accidentally echoes a credential in its output, it's replaced with [REDACTED] before the AI ever sees it.
No other product does this. Existing secrets managers hand plaintext to the requesting process. Venya doesn't.
Every session begins with a physical security key press. The AI agent cannot initiate a session — only a human pressing a FIDO2 key can authorize access. Sessions expire after 4 hours. Token refresh happens automatically, and an idle-expired session renews within the 4-hour hard cap — the cap is non-negotiable: past it, a human must re-authenticate with the FIDO2 key. When the session is past the cap, the AI gets an actionable error: "Ask the user to re-authenticate."
mTLS Between Server and Executor
The Venya server communicates with executor daemons over mutual TLS. Both sides verify each other's certificates. If an executor's certificate is revoked, the server refuses to relay commands. If someone spoofs an executor, the mTLS handshake fails before any secret is transmitted.
Commands run inside an sbx microVM with deny-by-default networking. The executor reads an operator-defined allowlist ( /etc/venya/egress-allowlist.txt ) and only permits connections to approved destinations. If a compromised command tries to phone home to an attacker's server, the connection is blocked at the sandbox level. DNS is restricted to the operator's resolver.
Every command execution is logged:
Who authorized the session (FIDO2-enrolled user)
What command was executed (full command string)
What secrets were injected (secret IDs, never values)
Both human operators (via venya audit CLI) and AI agents (via the get_audit MCP tool) can query the audit log. Non-admin users see only their own events.
Venya speaks the Model Context Protocol — the open standard for connecting AI assistants to external tools. It works with Claude Code, Cursor, and any MCP-compatible client. No proprietary lock-in. No vendor-specific API.
Guarantee
How It's Enforced
AI never sees plaintext credentials
Server-side wrapping + executor-side unwrapping + Rust output filter
Sessions require human authorization
FIDO2 hardware key binding (WebAuthn)
Sessions are time-limited
4-hour hard cap, 15-minute idle window, 5-minute access tokens
Executor identity is verified
Mutual TLS with certificate chain validation
Compromised executors are blocked
Revoked certs rejected at the executor's next revocation poll
Data exfiltration is prevented
sbx sandbox with deny-by-default egress allowlisting
Every action is traceable
Append-only audit log with user, executor, command, timestamp
Secrets are encrypted at rest
AES-256 with KEK-wrapped DEK (envelope encryption)
Who Is Venya For?
Infrastructure teams who want to use AI without compromising security.
Using Claude Code, Cursor, or similar AI coding assistants
Managing fleets of servers, databases, or cloud infrastructure
Concerned about handing credentials to AI agents
Operating in regulated environments where audit trails are mandatory
Tired of the choice between "move fast with AI" and "stay secure"
Venya is currently in alpha — early access for teams who want to shape the product.
See Venya in action: Alpha Demo Guide
Artifacts (installers, tarballs, SHA-256 sidecars) are published on the Releases page . Install one-liners (core/executor: Ubuntu 24.04 only; the Workstation CLI additionally runs on Debian 13, macOS, and Windows):
# Core server (root)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-core.sh | sudo \
VENYA_SKIP_PROMPT=yes VENYA_DB_PASSWORD= < strong-db-password > bash -s
# Executor (root; enrollment token from the core admin; a Docker account is REQUIRED — sbx pulls its agent template from Docker Hub; username + API key/access token via stdin)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-executor.sh | sudo \
VENYA_SKIP_PROMPT=yes VENYA_SERVER_URL=https:// < core-host > VENYA_EXECUTOR_ID= < executor-id > \
VENYA_EXECUTOR_ENROLLMENT_TOKEN= < token > bash -s
# Workstation CLI (non-root; Ubuntu 24.04, Debian 13, or macOS — verified on macOS 26.6.2 arm64 and Debian 13)
curl -fsSL https://github.com/tabith-llc/venya/releases/latest/download/install-venya-cli.sh | VENYA_SKIP_PROMPT=yes bash
# Workstation CLI (Windows) — machine-wide install, requires Administrator;
# standard users run the CLI afterward. Interactive desktop only (headless unsupported).
# Download install-venya-cli.ps1 from the Releases page, then run:
powershell - ExecutionPolicy Bypass - File install-venya - cli.ps1
Integrity: pin VENYA_TARBALL_SHA256 (hashes on the release page) for strict verification; unset, the installer fetches the .sha256 sidecar from the same origin as a corruption guardrail and fail-closes.
Workstation CLI config file: ~/.config/venya/config.json on Linux, ~/Library/Application Support/venya/config.json on macOS, %APPDATA%\venya\config.json on Windows. FIDO2 needs no extra setup on macOS (native IOKit HID transport, no root) or Windows (platform WebAuthn API — standard-user capable, interactive desktop required); on Linux the installer prints udev rules if /dev/hidraw* is not user-readable.
Production deployment guide: Installation Guide
Linux — Ubuntu 24.04 LTS (core/executor; tested target, installers assume it). Workstation CLI additionally supports Debian 13 and macOS (verified macOS 26.6.2 arm64)
Windows — Workstation CLI only (core and executor are Linux). Machine-wide install via install-venya-cli.ps1 requires Administrator; standard users run the CLI after install. FIDO2 ceremonies work for standard users via the platform WebAuthn API (verified Windows 11 25H2). Interactive desktop only — headless Windows is unsupported
PostgreSQL — installed automatically by the core installer (16 on Ubuntu 24.04)
Python 3.14 — pinned ( >=3.14,<3.15 ); provisioned automatically via uv
FIDO2 security key (YubiKey, SoloKeys, etc.)
Docker Sandboxes (sbx) — installed automatically by the executor in

[truncated]

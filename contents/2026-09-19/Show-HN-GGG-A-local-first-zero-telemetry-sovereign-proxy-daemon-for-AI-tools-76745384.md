---
source: "https://github.com/JOxKxER/garza-global-graviton"
hn_url: "https://news.ycombinator.com/item?id=49762923"
title: "Show HN: GGG – A local-first, zero-telemetry sovereign proxy daemon for AI tools"
article_title: "GitHub - JOxKxER/garza-global-graviton · GitHub"
image: "https://opengraph.githubassets.com/ff82177fe41c3e93bd0ca75642f94605cb8965952ccc8f131c7ee217a80359cf/JOxKxER/garza-global-graviton"
author: "joxkxer"
captured_at: "2026-09-19T03:26:24Z"
capture_tool: "hn-digest"
hn_id: 49762923
score: 1
comments: 0
posted_at: "2026-09-19T03:09:11Z"
tags:
  - hacker-news
---

# Show HN: GGG – A local-first, zero-telemetry sovereign proxy daemon for AI tools

- HN: [49762923](https://news.ycombinator.com/item?id=49762923)
- Source: [github.com](https://github.com/JOxKxER/garza-global-graviton)
- Score: 1
- Comments: 0
- Posted: 2026-09-19T03:09:11Z

## Translation

Title: Show HN: GGG – A local-first, zero-telemetry sovereign proxy daemon for AI tools
Article title: GitHub - JOxKxER/garza-global-graviton · GitHub
Description: Contribute to JOxKxER/garza-global-graviton development by creating an account on GitHub.

Article text:
GitHub - JOxKxER/garza-global-graviton · GitHub
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
JOxKxER
/
garza-global-graviton
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
93 Commits 93 Commits Folders and files
.continue .continue .github/ workflows .github/ workflows .vscode .vscode 04_Legal_and_IP 04_Legal_and_IP Joker Joker SalonApp SalonApp airgap_attestation airgap_attestation android android apps apps assets assets audit_snapshots audit_snapshots build_output/ joker_gui/ _internal build_output/ joker_gui/ _internal certs certs cnc_attestation cnc_attestation data data data_in data_in dataset_revisions dataset_revisions deployment deployment dsip_upload_package dsip_upload_package grant_submission_bundle grant_submission_bundle hardened_vault hardened_vault integrations integrations legal legal mobile_client mobile_client output output portals portals rapidapi_scraper rapidapi_scraper secure_snapshots secure_snapshots sensory_telemetry_vault sensory_telemetry_vault src-tauri src-tauri src src static static templates templates tests tests tradeassist tradeassist trading_app trading_app vault_backups vault_backups vault_pipeline/ incoming_data vault_pipeline/ incoming_data web_dashboard web_dashboard .dockerignore .dockerignore .gitignore .gitignore .server_state.json .server_state.json ARCHITECTURE.md ARCHITECTURE.md BUILD_INSTRUCTIONS.md BUILD_INSTRUCTIONS.md Dockerfile Dockerfile Dockerfile.public-data-pipeline Dockerfile.public-data-pipeline ENTERPRISE_PARTNERSHIPS.md ENTERPRISE_PARTNERSHIPS.md GGG_GOLD_RELEASE_v1.0.0_20260822_081839.zip GGG_GOLD_RELEASE_v1.0.0_20260822_081839.zip GGG_GOLD_RELEASE_v1.0.0_20260822_081839_MANIFEST.txt GGG_GOLD_RELEASE_v1.0.0_20260822_081839_MANIFEST.txt GGG_MissionControl_v0.1.0_20260822_005226.zip GGG_MissionControl_v0.1.0_20260822_005226.zip GGG_Production_Release_v0.1.0_20260822_075516.zip GGG_Production_Release_v0.1.0_20260822_075516.zip GGG_Production_Release_v0.1.0_20260822_075516_MANIFEST.txt GGG_Production_Release_v0.1.0_20260822_075516_MANIFEST.txt Garza_Global_Graviton.spec Garza_Global_Graviton.spec Garza_Global_Graviton_DSIP_Release5_20260824.zip Garza_Global_Graviton_DSIP_Release5_20260824.zip Garza_Global_Graviton_SBIR_S
[truncated]
Garza Global Graviton LLC — Sovereign Platform
Garza Global Graviton LLC builds sovereign, 100% air-gapped edge
computing systems — hardware and software that run entirely on local
infrastructure, with zero mandatory network dependency, zero cloud
telemetry, and zero vendor-side visibility into your data.
v1.4.0 — Mobile-Responsive Web Hub
The node's local web hub ( index.html ) is now fully
mobile-responsive as of v1.4.0 : brand bars, the hardware plaque, the
physics benchmark table, and the local Ollama bridge all reflow cleanly on
phones and tablets (screens under 768px), with buttons, inputs, and panels
expanding to full width for comfortable touch interaction, and the physics
benchmark table gaining smooth horizontal scrolling instead of a broken
layout.
100% Air-Gapped Sovereign Edge
Every node runs as a Synthetic Data Center : a single piece of local,
air-gapped hardware organized like a living organism, paced and defended by
cooperating biological daemons that never touch the network:
Metabolic Heart — paces admission and pulses the node's cadence at
microsecond precision.
Liver — scrubs stale memory in-place with vectorized zeroization,
never leaving residue.
Lungs — inhale and exhale data through a zero-copy ring buffer,
moving bytes without ever duplicating them.
Immune System — a zero-trust scanner that quarantines and
neutralizes any byte pattern it does not recognize.
Because every buffer is a single preallocated memoryview over a numpy
array, data moves through the organism without a single extra copy, and
because the node is air-gapped, every daemon above runs at local hardware
speed, not network speed.
Local-only by construction. The Ollama bridge only accepts
localhost / 127.0.0.1 / [::1] endpoints; remote hosts are rejected
before any request is ever sent, keeping the bridge 100% offline by
construction.
Sovereign data ownership. Nothing leaves the device unless you
explicitly choose to send it — there is no cloud dependency in the
critical path.
Grab the latest packaged sovereign app from the GitHub Releases page:
https://github.com/JOxKxER/garza-global-graviton/releases/latest
Or click Download Sovereign App directly from the top of the local web
hub ( index.html ).
Connect your local Ollama instance
Install and start Ollama on the same machine (or
another host reachable at localhost /loopback).
Pull a model, e.g. ollama pull llama3.2:3b .
Open index.html in a browser and, in the Local Ollama Bridge
panel, confirm the endpoint (default http://localhost:11434 ) and model
name, then click Check Status .
Once the badge shows ONLINE , use the Live Local Benchmark panel
or the chat row to send prompts — everything runs locally, offline.
Running the daemon binary — Windows SmartScreen
The released ggg-daemon.exe is self-signed (no commercial code-signing
certificate), so Windows SmartScreen will show "Windows protected your
PC" on first launch. This is expected for unsigned local binaries — it is
a reputation warning, not a detection.
On the SmartScreen dialog, click "More info" .
Click "Run anyway" — the choice is remembered for that file.
Verify authenticity before running (recommended): every release ships a
SHA256SUMS.txt alongside the binary. Compare the hash:
Get-FileHash .\ggg - daemon_v1. 4. 5_windows.zip - Algorithm SHA256
# Then compare against the matching line in SHA256SUMS.txt
If the hash matches the published manifest, the binary is exactly what this
repository built. If it does not match, do not run it.
Use your local Ollama as the model provider in VS Code
Keep your editor 100% local by pointing a local-model extension at the same
Ollama daemon ( http://localhost:11434 ). Add to your VS Code settings.json
( Ctrl+Shift+P → Preferences: Open User Settings (JSON) ):
{
"continue.server" : {
"port" : 11434
},
"continue.models" : [
{
"title" : " GGG Local Qwen " ,
"provider" : " ollama " ,
"model" : " qwen2.5-coder:latest " ,
"apiBase" : " http://localhost:11434 "
}
],
"ollama.baseUrl" : " http://localhost:11434 " ,
"localai.model.basePath" : " http://localhost:11434 "
}
Continue.dev ( Continue.continue ): the continue.models entry selects
your local qwen2.5-coder as the chat/edit model. See
.continue/config.json for this repo's ready-made
Continue configuration.
Ollama extensions (e.g. Ollama.ollama ): ollama.baseUrl redirects
all model calls to loopback.
GitHub Copilot does not support custom/local model providers — use
Continue or an Ollama extension for fully offline AI assistance.
Air-Gap Zero-Network-Leakage Attestation Platform
Cryptographic proof, for prospective enterprise buyers, that a sensitive
manufacturing/data pipeline ran in a strictly air-gapped, tamper-evident
environment -- verifiable entirely offline, without ever seeing the vendor's
proprietary pipeline source.
All of this lives under airgap_attestation/ .
Honesty note: no software system can produce an unconditional
mathematical proof of "zero network leakage" from a black box. What this
platform delivers is a layered, tamper-evident, independently falsifiable
evidence chain (hardware-rooted measurement + a cryptographically sealed
execution log + dual signatures) where any single point of compromise is
detectable. See airgap_attestation/CLIENT_ONBOARDING.md
for what a buyer should actually conclude from a passing/failing result.
CLIENT AIR-GAPPED EXECUTION ENCLAVE CLIENT
| 1. commit-reveal blind | a. TPM/enclave measured boot |
| (sha256(sample||salt)) | b. NIC disabled at hardware level |
|--- POST /v1/submissions ------->| c. job runs; AuditManifestBuilder |
|<-- {commitment_id, nonce} ------| records PROCESS_START/END, |
| | NET_IFACE_SNAPSHOT, syscall counts |
| 2. deliver encrypted sample | d. events sealed into a Merkle tree |
| (out-of-band, key sent | e. TPM/enclave QUOTE(nonce||root) |
| via a SEPARATE channel) | f. platform Ed25519 signature over |
| | (root, quote, validity window) |
| 3. poll GET /v1/submissions/{id} |
| 4. GET /v1/submissions/{id}/bundle -> AttestationBundle.json ------------>|
| |
| 5. verify_cli.py, fully |
| offline: Merkle root, |
| TPM quote, platform sig, |
| nonce freshness, zero- |
| network invariant |
| -> PASS / FAIL |
Full protocol writeup (data schemas, hashing/signing details, edge cases):
see the architecture discussion in project history, or read the code directly
-- every design decision is documented as a comment at its point of use:
merkle.py , schemas.py ,
verify_client.py .
airgap_attestation/
merkle.py Domain-separated Merkle tree (leaf/node hash separation,
no odd-node duplication -- avoids classic forgery bugs)
schemas.py Wire-format dataclasses (SubmissionCommitment, ExecutionEvent,
AuditManifest, TpmQuoteEvidence, AttestationBundle, ...)
signing.py Ed25519 keygen/sign/verify (platform transport-layer identity)
attestation.py HardwareAttestor interface: Tpm2ToolsAttestor (real TPM 2.0,
via tpm2-tools) + ReferenceSoftwareAttestor (dev/test only,
explicitly rejected by production verification)
audit_manifest.py AuditManifestBuilder + NetworkActivityMonitor (pluggable;
LocalReferenceMonitor ships as a portable fallback)
proof_bundle.py Assembles/signs/saves/loads the final AttestationBundle
verify_client.py The entire client-side verifier -- no vendor source needed
demo_end_to_end.py Runnable proof-of-concept: build -> sign -> verify
api/
store.py SQLite-backed, single-use nonce/commitment store (atomic
UPDATE ... WHERE guard -- no TOCTOU replay window)
nonce_service.py FastAPI backend: submission intake, status, bundle
download, internal ingest, rate limiting, security headers
cli/
verify_cli.py Buyer-facing CLI wrapper around verify_client.py
container/
Dockerfile Hardened image: distroless nonroot final stage
docker-compose.yml (repo-relative: airgap_attestation/docker-compose.yml)
network_mode: none, read_only, cap_drop ALL, seccomp
seccomp-hardened.json Kernel-level deny-list for network syscalls
firecracker_config.json Stronger alternative: no NIC device exists at all
entrypoint.py In-container job runner
deployment_runbook.ps1 Step-by-step build/run/ingest/verify commands
CLIENT_ONBOARDING.md Buyer-facing, step-by-step usage guide
tests/
test_airgap_merkle.py, test_airgap_signing.py, test_airgap_pipeline.py,
test_airgap_api.py, test_airgap_cli.py (60 tests, ~84% coverage of the package)
render.yaml Render deployment blueprint for nonce_service:app
.github/workflows/ci.yml GitHub Actions: pytest + coverag

[truncated]

## Original Extract

Contribute to JOxKxER/garza-global-graviton development by creating an account on GitHub.

GitHub - JOxKxER/garza-global-graviton · GitHub
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
JOxKxER
/
garza-global-graviton
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
93 Commits 93 Commits Folders and files
.continue .continue .github/ workflows .github/ workflows .vscode .vscode 04_Legal_and_IP 04_Legal_and_IP Joker Joker SalonApp SalonApp airgap_attestation airgap_attestation android android apps apps assets assets audit_snapshots audit_snapshots build_output/ joker_gui/ _internal build_output/ joker_gui/ _internal certs certs cnc_attestation cnc_attestation data data data_in data_in dataset_revisions dataset_revisions deployment deployment dsip_upload_package dsip_upload_package grant_submission_bundle grant_submission_bundle hardened_vault hardened_vault integrations integrations legal legal mobile_client mobile_client output output portals portals rapidapi_scraper rapidapi_scraper secure_snapshots secure_snapshots sensory_telemetry_vault sensory_telemetry_vault src-tauri src-tauri src src static static templates templates tests tests tradeassist tradeassist trading_app trading_app vault_backups vault_backups vault_pipeline/ incoming_data vault_pipeline/ incoming_data web_dashboard web_dashboard .dockerignore .dockerignore .gitignore .gitignore .server_state.json .server_state.json ARCHITECTURE.md ARCHITECTURE.md BUILD_INSTRUCTIONS.md BUILD_INSTRUCTIONS.md Dockerfile Dockerfile Dockerfile.public-data-pipeline Dockerfile.public-data-pipeline ENTERPRISE_PARTNERSHIPS.md ENTERPRISE_PARTNERSHIPS.md GGG_GOLD_RELEASE_v1.0.0_20260822_081839.zip GGG_GOLD_RELEASE_v1.0.0_20260822_081839.zip GGG_GOLD_RELEASE_v1.0.0_20260822_081839_MANIFEST.txt GGG_GOLD_RELEASE_v1.0.0_20260822_081839_MANIFEST.txt GGG_MissionControl_v0.1.0_20260822_005226.zip GGG_MissionControl_v0.1.0_20260822_005226.zip GGG_Production_Release_v0.1.0_20260822_075516.zip GGG_Production_Release_v0.1.0_20260822_075516.zip GGG_Production_Release_v0.1.0_20260822_075516_MANIFEST.txt GGG_Production_Release_v0.1.0_20260822_075516_MANIFEST.txt Garza_Global_Graviton.spec Garza_Global_Graviton.spec Garza_Global_Graviton_DSIP_Release5_20260824.zip Garza_Global_Graviton_DSIP_Release5_20260824.zip Garza_Global_Graviton_SBIR_S
[truncated]
Garza Global Graviton LLC — Sovereign Platform
Garza Global Graviton LLC builds sovereign, 100% air-gapped edge
computing systems — hardware and software that run entirely on local
infrastructure, with zero mandatory network dependency, zero cloud
telemetry, and zero vendor-side visibility into your data.
v1.4.0 — Mobile-Responsive Web Hub
The node's local web hub ( index.html ) is now fully
mobile-responsive as of v1.4.0 : brand bars, the hardware plaque, the
physics benchmark table, and the local Ollama bridge all reflow cleanly on
phones and tablets (screens under 768px), with buttons, inputs, and panels
expanding to full width for comfortable touch interaction, and the physics
benchmark table gaining smooth horizontal scrolling instead of a broken
layout.
100% Air-Gapped Sovereign Edge
Every node runs as a Synthetic Data Center : a single piece of local,
air-gapped hardware organized like a living organism, paced and defended by
cooperating biological daemons that never touch the network:
Metabolic Heart — paces admission and pulses the node's cadence at
microsecond precision.
Liver — scrubs stale memory in-place with vectorized zeroization,
never leaving residue.
Lungs — inhale and exhale data through a zero-copy ring buffer,
moving bytes without ever duplicating them.
Immune System — a zero-trust scanner that quarantines and
neutralizes any byte pattern it does not recognize.
Because every buffer is a single preallocated memoryview over a numpy
array, data moves through the organism without a single extra copy, and
because the node is air-gapped, every daemon above runs at local hardware
speed, not network speed.
Local-only by construction. The Ollama bridge only accepts
localhost / 127.0.0.1 / [::1] endpoints; remote hosts are rejected
before any request is ever sent, keeping the bridge 100% offline by
construction.
Sovereign data ownership. Nothing leaves the device unless you
explicitly choose to send it — there is no cloud dependency in the
critical path.
Grab the latest packaged sovereign app from the GitHub Releases page:
https://github.com/JOxKxER/garza-global-graviton/releases/latest
Or click Download Sovereign App directly from the top of the local web
hub ( index.html ).
Connect your local Ollama instance
Install and start Ollama on the same machine (or
another host reachable at localhost /loopback).
Pull a model, e.g. ollama pull llama3.2:3b .
Open index.html in a browser and, in the Local Ollama Bridge
panel, confirm the endpoint (default http://localhost:11434 ) and model
name, then click Check Status .
Once the badge shows ONLINE , use the Live Local Benchmark panel
or the chat row to send prompts — everything runs locally, offline.
Running the daemon binary — Windows SmartScreen
The released ggg-daemon.exe is self-signed (no commercial code-signing
certificate), so Windows SmartScreen will show "Windows protected your
PC" on first launch. This is expected for unsigned local binaries — it is
a reputation warning, not a detection.
On the SmartScreen dialog, click "More info" .
Click "Run anyway" — the choice is remembered for that file.
Verify authenticity before running (recommended): every release ships a
SHA256SUMS.txt alongside the binary. Compare the hash:
Get-FileHash .\ggg - daemon_v1. 4. 5_windows.zip - Algorithm SHA256
# Then compare against the matching line in SHA256SUMS.txt
If the hash matches the published manifest, the binary is exactly what this
repository built. If it does not match, do not run it.
Use your local Ollama as the model provider in VS Code
Keep your editor 100% local by pointing a local-model extension at the same
Ollama daemon ( http://localhost:11434 ). Add to your VS Code settings.json
( Ctrl+Shift+P → Preferences: Open User Settings (JSON) ):
{
"continue.server" : {
"port" : 11434
},
"continue.models" : [
{
"title" : " GGG Local Qwen " ,
"provider" : " ollama " ,
"model" : " qwen2.5-coder:latest " ,
"apiBase" : " http://localhost:11434 "
}
],
"ollama.baseUrl" : " http://localhost:11434 " ,
"localai.model.basePath" : " http://localhost:11434 "
}
Continue.dev ( Continue.continue ): the continue.models entry selects
your local qwen2.5-coder as the chat/edit model. See
.continue/config.json for this repo's ready-made
Continue configuration.
Ollama extensions (e.g. Ollama.ollama ): ollama.baseUrl redirects
all model calls to loopback.
GitHub Copilot does not support custom/local model providers — use
Continue or an Ollama extension for fully offline AI assistance.
Air-Gap Zero-Network-Leakage Attestation Platform
Cryptographic proof, for prospective enterprise buyers, that a sensitive
manufacturing/data pipeline ran in a strictly air-gapped, tamper-evident
environment -- verifiable entirely offline, without ever seeing the vendor's
proprietary pipeline source.
All of this lives under airgap_attestation/ .
Honesty note: no software system can produce an unconditional
mathematical proof of "zero network leakage" from a black box. What this
platform delivers is a layered, tamper-evident, independently falsifiable
evidence chain (hardware-rooted measurement + a cryptographically sealed
execution log + dual signatures) where any single point of compromise is
detectable. See airgap_attestation/CLIENT_ONBOARDING.md
for what a buyer should actually conclude from a passing/failing result.
CLIENT AIR-GAPPED EXECUTION ENCLAVE CLIENT
| 1. commit-reveal blind | a. TPM/enclave measured boot |
| (sha256(sample||salt)) | b. NIC disabled at hardware level |
|--- POST /v1/submissions ------->| c. job runs; AuditManifestBuilder |
|<-- {commitment_id, nonce} ------| records PROCESS_START/END, |
| | NET_IFACE_SNAPSHOT, syscall counts |
| 2. deliver encrypted sample | d. events sealed into a Merkle tree |
| (out-of-band, key sent | e. TPM/enclave QUOTE(nonce||root) |
| via a SEPARATE channel) | f. platform Ed25519 signature over |
| | (root, quote, validity window) |
| 3. poll GET /v1/submissions/{id} |
| 4. GET /v1/submissions/{id}/bundle -> AttestationBundle.json ------------>|
| |
| 5. verify_cli.py, fully |
| offline: Merkle root, |
| TPM quote, platform sig, |
| nonce freshness, zero- |
| network invariant |
| -> PASS / FAIL |
Full protocol writeup (data schemas, hashing/signing details, edge cases):
see the architecture discussion in project history, or read the code directly
-- every design decision is documented as a comment at its point of use:
merkle.py , schemas.py ,
verify_client.py .
airgap_attestation/
merkle.py Domain-separated Merkle tree (leaf/node hash separation,
no odd-node duplication -- avoids classic forgery bugs)
schemas.py Wire-format dataclasses (SubmissionCommitment, ExecutionEvent,
AuditManifest, TpmQuoteEvidence, AttestationBundle, ...)
signing.py Ed25519 keygen/sign/verify (platform transport-layer identity)
attestation.py HardwareAttestor interface: Tpm2ToolsAttestor (real TPM 2.0,
via tpm2-tools) + ReferenceSoftwareAttestor (dev/test only,
explicitly rejected by production verification)
audit_manifest.py AuditManifestBuilder + NetworkActivityMonitor (pluggable;
LocalReferenceMonitor ships as a portable fallback)
proof_bundle.py Assembles/signs/saves/loads the final AttestationBundle
verify_client.py The entire client-side verifier -- no vendor source needed
demo_end_to_end.py Runnable proof-of-concept: build -> sign -> verify
api/
store.py SQLite-backed, single-use nonce/commitment store (atomic
UPDATE ... WHERE guard -- no TOCTOU replay window)
nonce_service.py FastAPI backend: submission intake, status, bundle
download, internal ingest, rate limiting, security headers
cli/
verify_cli.py Buyer-facing CLI wrapper around verify_client.py
container/
Dockerfile Hardened image: distroless nonroot final stage
docker-compose.yml (repo-relative: airgap_attestation/docker-compose.yml)
network_mode: none, read_only, cap_drop ALL, seccomp
seccomp-hardened.json Kernel-level deny-list for network syscalls
firecracker_config.json Stronger alternative: no NIC device exists at all
entrypoint.py In-container job runner
deployment_runbook.ps1 Step-by-step build/run/ingest/verify commands
CLIENT_ONBOARDING.md Buyer-facing, step-by-step usage guide
tests/
test_airgap_merkle.py, test_airgap_signing.py, test_airgap_pipeline.py,
test_airgap_api.py, test_airgap_cli.py (60 tests, ~84% coverage of the package)
render.yaml Render deployment blueprint for nonce_service:app
.github/workflows/ci.yml GitHub Actions: pytest + coverag

[truncated]

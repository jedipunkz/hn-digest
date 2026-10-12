---
source: "https://github.com/pigyon/tanuki"
hn_url: "https://news.ycombinator.com/item?id=50061544"
title: "Tanuki – Local proxy/wrapper bypassing Claude guardrails for security research"
article_title: "GitHub - Pigyon/tanuki: Proxy for AI security testing. keeping real target data out of LLM'S API logs. · GitHub"
image: "https://opengraph.githubassets.com/94835998861662a03ccee55376a54f0ec0b60d3adc7ad9e94a5a59cb48a1dd1c/Pigyon/tanuki"
author: "nedlir"
captured_at: "2026-10-12T20:08:39Z"
capture_tool: "hn-digest"
hn_id: 50061544
score: 1
comments: 0
posted_at: "2026-10-12T19:21:42Z"
tags:
  - hacker-news
---

# Tanuki – Local proxy/wrapper bypassing Claude guardrails for security research

- HN: [50061544](https://news.ycombinator.com/item?id=50061544)
- Source: [github.com](https://github.com/pigyon/tanuki)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T19:21:42Z

## Translation

Title: Tanuki – Local proxy/wrapper bypassing Claude guardrails for security research
Article title: GitHub - Pigyon/tanuki: Proxy for AI security testing. keeping real target data out of LLM'S API logs. · GitHub
Description: Proxy for AI security testing. keeping real target data out of LLM'S API logs. - Pigyon/tanuki

Article text:
GitHub - Pigyon/tanuki: Proxy for AI security testing. keeping real target data out of LLM'S API logs. · GitHub
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
Pigyon
tanuki Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Proxy for AI security testing. keeping real target data out of LLM'S API logs.
Readme Apache-2.0 license Activity Custom properties Stars
1 fork Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
61 Commits 61 Commits Folders and files
.github .github avatar avatar cmd/ tanuki cmd/ tanuki internal/ tanuki internal/ tanuki .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .hadolint.yaml .hadolint.yaml Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.dev.yaml compose.dev.yaml compose.yaml compose.yaml go.mod go.mod go.sum go.sum tanuki tanuki View all files Repository files navigation
"The greatest craft is to hide craft."
Thomas Fuller
Proxy for AI-assisted security testing/privacy
Tanuki sits between your local Claude Code agent and the Anthropic API , rewriting every target-identifying string in real time. Your tools hit the real target; the model only ever sees a fictional localhost project.
Tanuki is an anonymization proxy built for AI-assisted security testing and bug bounty hunting. When you drive an agent like Claude Code through a pentest, every real domain, IP, and payload in your prompts leaves your machine. That's an OpSec risk, and it can also breach bug bounty program rules or an NDA.
Tanuki runs as a Docker container and transparently intercepts that API traffic. Before a prompt leaves your machine, it swaps real target data (domains, subdomains, IPs, emails, paths, cloud identifiers, and aggressive security terminology) for harmless fictional development equivalents. When the model responds, the fiction is mapped back to reality so your local tools execute against the real targets. The model believes it is helping you debug a local dev application, and the upstream provider is blind to your actual engagement.
flowchart TD
classDef default fill:#050505,stroke:#8b0000,stroke-width:1px,color:#ff3333,font-family:Courier
classDef proxy fill:#110000,stroke:#ff0000,stroke-width:2px,color:#ffffff,font-family:Courier
classDef hooks fill:#000000,stroke:#550000,stroke-width:1px,stroke-dasharray:4 4,color:#cc0000,font-family:Courier
classDef upstream fill:#0a0a0a,stroke:#333333,stroke-width:1px,color:#666666,font-family:Courier
classDef data fill:#1a0505,stroke:#aa0000,stroke-width:1px,color:#ff6666,font-family:Courier
subgraph Host["LOCAL MACHINE"]
C["Claude Code"]
H[/"Tanuki Hooks<br/>(docker exec)"/]
C -. "Tool execution" .-> H
end
subgraph Container["DOCKER CONTAINER"]
P{"Tanuki Proxy<br/>[scratch ~10MB]"}:::proxy
D[("( mappings volume )")]:::data
end
A((("Anthropic API"))):::upstream
C == "Real prompt data" ==> P
P -. "Loads config" .-> D
P -- "Rewrites to Fiction" --> A
A -- "Responds in Fiction" --> P
P == "Rewrites to Real" ==> C
class C default
class H hooks
style Host fill:transparent,stroke:#8b0000,stroke-width:1px,color:#8b0000,font-family:Courier
style Container fill:#020000,stroke:#4a0000,stroke-width:1px,stroke-dasharray:3 3,color:#8b0000,font-family:Courier
Loading
Direction
Path
What happens
Inbound
your prompt → API
Domains, org names, emails, IPs, paths, cloud identifiers, and pentest terminology are replaced with localhost fiction before the model sees them.
Outbound
API → Claude Code
Fiction terms are rewritten back to real values in the response stream, so Claude Code's local tools receive real target data.
Tool calls
hook via docker exec
A PreToolUse hook rewrites fiction→real in tool arguments, so tools hit the actual targets. Tool output needs no hook: it reaches the model only inside the next request, which the proxy rewrites like any other.
Example
You're testing bounty.amazon.com . Here's what Tanuki does:
Each target domain gets three mappings: the base domain, a wildcard for
subdomains that have not been seen yet, and a dedicated mapping per
subdomain as it is detected. The wildcard is what stops an unmapped host
like internal-admin.amazon.com from leaking its label upstream.
git clone https://github.com/Pigyon/tanuki.git
cd tanuki
docker compose up -d
claude
That pulls the latest published image from GitHub Container Registry, so no Go toolchain and no build are needed.
To build from source instead of pulling:
docker compose -f compose.yaml -f compose.dev.yaml up -d --build
On startup the proxy auto-configures the Claude Code hook and proxy env in .claude/settings.json , and writes the cover-story .claude/CLAUDE.md . Domains, emails, and IPs are auto-detected from your conversation, so no manual target configuration is needed.
When you type "check bounty.amazon.com for SQL injection" , the proxy:
Detects amazon.com and bounty.amazon.com as new domains
Creates all derivative mappings (org, email, path, cloud)
Rewrites the request before it reaches the API
A ./tanuki wrapper is included so you don't have to type docker compose exec tanuki /tanuki <command> :
./tanuki add example.com
./tanuki map
./tanuki status
Verify rewriting
Test the rewriting engine without hitting the API:
./tanuki test " scan bounty.amazon.com for SQL injection and contact admin@amazon.com "
Configuration
Everything is optional; without it, targets are auto-detected from your conversation. To pre-configure an engagement, edit compose.yaml :
environment :
# Custom fiction org name (default: DEVTARGET)
- TANUKI_FICTION_ORG=ACMECORP
# Pre-seed target domains (comma-separated)
- TANUKI_DOMAINS=example.com,api.example.com
# Pre-seed target IPs (comma-separated)
- TANUKI_IPS=93.184.216.34
# Custom rewrite rules (semicolon-separated, pipe-delimited real|fiction)
- TANUKI_RULES=John Smith|Jane Developer;Project Phoenix|Test Alpha
Custom rewrite rules
For strings that aren't domains or IPs (people, project names, internal jargon), add custom rules via compose.yaml (applied on every container start) or at runtime:
./tanuki add-rule " John Smith " " Jane Developer "
./tanuki add-rule " Project Phoenix " " Test Alpha "
Named engagements
For separate projects or custom fiction org names:
./tanuki init bugbounty --fiction-org ACMECORP
./tanuki add bounty.amazon.com
CLI reference
All commands run via ./tanuki <command> (or docker compose exec tanuki /tanuki <command> ):
Mapping-based (per engagement, all auto-detected)
Type
Detection
Fiction
Domains
Auto-detected in prompts & tool output
localhost:<port>
Org names
Derived from domain
DEVTARGET (configurable)
Emails
Derived from domain
user@devtarget.local
IPs (v4 + v6)
Auto-detected in prompts & tool output
127.0.0.x / ::ffff:...
Paths
Derived from org name
/dev/devtarget/
Cloud resources
Derived from org name
file:///tmp/devtarget...
Custom rules
User-defined via env or CLI
User-defined
Terminology (built-in, applied to all engagements)
Compound pentest-specific terms are replaced with neutral dev equivalents. Only unambiguous multi-word terms are replaced. Single words like target , payload , or exploit are left alone so legitimate code discussions don't break.
Run ./tanuki terms to see the full dictionary. The list lives in internal/tanuki/terminology.json and is embedded at build time; to customize it, edit the file and rebuild.
Container: scratch -based image with a single static Go binary plus CA certs. No OS, no shell, no runtime, ~10MB total. All capabilities dropped, read-only filesystem, no privilege escalation.
Everything runs in Docker. The proxy, hooks, and CLI all execute inside the container; hooks are invoked by Claude Code via docker exec , so no host-side binary is needed.
Matching: an Aho-Corasick automaton does longest-first replacement in a single pass over the stream.
JSON-aware: request and response bodies are rewritten string by string, never as raw text. Protocol fields ( model , role , type , ids, tool names) are left alone, base64 data is neither rewritten nor scanned, and thinking blocks go back to the API exactly as the model wrote them, since they are signed.
No HTTPS on the proxy. The proxy listens on plain HTTP. That's acceptable because it only accepts connections from localhost (Claude Code on the same machine); the upstream connection to the Anthropic API uses HTTPS.
Literal string matching only. Encoded forms (URL-encoded, base64, and similar) won't be caught by the Aho-Corasick automaton.
Terminology is one-directional. Pentest terms are rewritten real→fiction only. The model's responses use fiction terms, which pass through as-is.
Wildcard subdomains are not reversible. Every subdomain without its own mapping collapses onto a single wildcard value, so it cannot be mapped back. Auto-detection creates dedicated mappings for subdomains it sees, and those round-trip normally; the wildcard is the safety net for anything that slips past detection, and it favours hiding the hostname over keeping it usable.
Detection is the weak link, not the rewrite. Every outbound body is re-scanned after rewriting and refused if a real value survives, so a rewriter bug cannot leak. Nothing protects against a value the detector never recognised in the first place — the two limitations below are therefore the ones that matter.
Curated TLD list for auto-detection. Auto-detection uses a curated list of ~70 common TLDs to avoid false positives from code patterns like readme.md , foo.bar , or user.id . TLDs that collide with everyday code ( .info , .name , .app , .store , .top , .site , as in logger.info or user.name ) are left out. Domains on those or other unusual TLDs ( .pizza , .click , .it , .id ) can be added manually with ./tanuki add .
Ignored domains need an explicit add to become targets. github.com , google.com , and your own email domain are skipped by auto-detection so Claude Code's own git remotes, connectors, and account context don't pollute the table. The trade-off is that auto-detection alone will not map one of these, so on the rare engagement that targets GitHub or Google you must ./tanuki add github.com first; the residual check also skips ignored domains, so it will not catch the omission for you.
Thinking stays in fiction. Thinking blocks are signed by the API, so they are passed through unchanged in both directions. Thinking shown in Claude Code reads in fiction terms.
Streamed text is held back briefly. A fiction value can be streamed a few characters per event, so the proxy holds back the end of a text or tool-input block while it could still be the start of one, and sends it once the next event settles it. Display can lag by a few characters; the full value always arrives reversed.
tanuki

[truncated]

## Original Extract

Proxy for AI security testing. keeping real target data out of LLM'S API logs. - Pigyon/tanuki

GitHub - Pigyon/tanuki: Proxy for AI security testing. keeping real target data out of LLM'S API logs. · GitHub
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
Pigyon
tanuki Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Proxy for AI security testing. keeping real target data out of LLM'S API logs.
Readme Apache-2.0 license Activity Custom properties Stars
1 fork Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
61 Commits 61 Commits Folders and files
.github .github avatar avatar cmd/ tanuki cmd/ tanuki internal/ tanuki internal/ tanuki .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .hadolint.yaml .hadolint.yaml Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.dev.yaml compose.dev.yaml compose.yaml compose.yaml go.mod go.mod go.sum go.sum tanuki tanuki View all files Repository files navigation
"The greatest craft is to hide craft."
Thomas Fuller
Proxy for AI-assisted security testing/privacy
Tanuki sits between your local Claude Code agent and the Anthropic API , rewriting every target-identifying string in real time. Your tools hit the real target; the model only ever sees a fictional localhost project.
Tanuki is an anonymization proxy built for AI-assisted security testing and bug bounty hunting. When you drive an agent like Claude Code through a pentest, every real domain, IP, and payload in your prompts leaves your machine. That's an OpSec risk, and it can also breach bug bounty program rules or an NDA.
Tanuki runs as a Docker container and transparently intercepts that API traffic. Before a prompt leaves your machine, it swaps real target data (domains, subdomains, IPs, emails, paths, cloud identifiers, and aggressive security terminology) for harmless fictional development equivalents. When the model responds, the fiction is mapped back to reality so your local tools execute against the real targets. The model believes it is helping you debug a local dev application, and the upstream provider is blind to your actual engagement.
flowchart TD
classDef default fill:#050505,stroke:#8b0000,stroke-width:1px,color:#ff3333,font-family:Courier
classDef proxy fill:#110000,stroke:#ff0000,stroke-width:2px,color:#ffffff,font-family:Courier
classDef hooks fill:#000000,stroke:#550000,stroke-width:1px,stroke-dasharray:4 4,color:#cc0000,font-family:Courier
classDef upstream fill:#0a0a0a,stroke:#333333,stroke-width:1px,color:#666666,font-family:Courier
classDef data fill:#1a0505,stroke:#aa0000,stroke-width:1px,color:#ff6666,font-family:Courier
subgraph Host["LOCAL MACHINE"]
C["Claude Code"]
H[/"Tanuki Hooks<br/>(docker exec)"/]
C -. "Tool execution" .-> H
end
subgraph Container["DOCKER CONTAINER"]
P{"Tanuki Proxy<br/>[scratch ~10MB]"}:::proxy
D[("( mappings volume )")]:::data
end
A((("Anthropic API"))):::upstream
C == "Real prompt data" ==> P
P -. "Loads config" .-> D
P -- "Rewrites to Fiction" --> A
A -- "Responds in Fiction" --> P
P == "Rewrites to Real" ==> C
class C default
class H hooks
style Host fill:transparent,stroke:#8b0000,stroke-width:1px,color:#8b0000,font-family:Courier
style Container fill:#020000,stroke:#4a0000,stroke-width:1px,stroke-dasharray:3 3,color:#8b0000,font-family:Courier
Loading
Direction
Path
What happens
Inbound
your prompt → API
Domains, org names, emails, IPs, paths, cloud identifiers, and pentest terminology are replaced with localhost fiction before the model sees them.
Outbound
API → Claude Code
Fiction terms are rewritten back to real values in the response stream, so Claude Code's local tools receive real target data.
Tool calls
hook via docker exec
A PreToolUse hook rewrites fiction→real in tool arguments, so tools hit the actual targets. Tool output needs no hook: it reaches the model only inside the next request, which the proxy rewrites like any other.
Example
You're testing bounty.amazon.com . Here's what Tanuki does:
Each target domain gets three mappings: the base domain, a wildcard for
subdomains that have not been seen yet, and a dedicated mapping per
subdomain as it is detected. The wildcard is what stops an unmapped host
like internal-admin.amazon.com from leaking its label upstream.
git clone https://github.com/Pigyon/tanuki.git
cd tanuki
docker compose up -d
claude
That pulls the latest published image from GitHub Container Registry, so no Go toolchain and no build are needed.
To build from source instead of pulling:
docker compose -f compose.yaml -f compose.dev.yaml up -d --build
On startup the proxy auto-configures the Claude Code hook and proxy env in .claude/settings.json , and writes the cover-story .claude/CLAUDE.md . Domains, emails, and IPs are auto-detected from your conversation, so no manual target configuration is needed.
When you type "check bounty.amazon.com for SQL injection" , the proxy:
Detects amazon.com and bounty.amazon.com as new domains
Creates all derivative mappings (org, email, path, cloud)
Rewrites the request before it reaches the API
A ./tanuki wrapper is included so you don't have to type docker compose exec tanuki /tanuki <command> :
./tanuki add example.com
./tanuki map
./tanuki status
Verify rewriting
Test the rewriting engine without hitting the API:
./tanuki test " scan bounty.amazon.com for SQL injection and contact admin@amazon.com "
Configuration
Everything is optional; without it, targets are auto-detected from your conversation. To pre-configure an engagement, edit compose.yaml :
environment :
# Custom fiction org name (default: DEVTARGET)
- TANUKI_FICTION_ORG=ACMECORP
# Pre-seed target domains (comma-separated)
- TANUKI_DOMAINS=example.com,api.example.com
# Pre-seed target IPs (comma-separated)
- TANUKI_IPS=93.184.216.34
# Custom rewrite rules (semicolon-separated, pipe-delimited real|fiction)
- TANUKI_RULES=John Smith|Jane Developer;Project Phoenix|Test Alpha
Custom rewrite rules
For strings that aren't domains or IPs (people, project names, internal jargon), add custom rules via compose.yaml (applied on every container start) or at runtime:
./tanuki add-rule " John Smith " " Jane Developer "
./tanuki add-rule " Project Phoenix " " Test Alpha "
Named engagements
For separate projects or custom fiction org names:
./tanuki init bugbounty --fiction-org ACMECORP
./tanuki add bounty.amazon.com
CLI reference
All commands run via ./tanuki <command> (or docker compose exec tanuki /tanuki <command> ):
Mapping-based (per engagement, all auto-detected)
Type
Detection
Fiction
Domains
Auto-detected in prompts & tool output
localhost:<port>
Org names
Derived from domain
DEVTARGET (configurable)
Emails
Derived from domain
user@devtarget.local
IPs (v4 + v6)
Auto-detected in prompts & tool output
127.0.0.x / ::ffff:...
Paths
Derived from org name
/dev/devtarget/
Cloud resources
Derived from org name
file:///tmp/devtarget...
Custom rules
User-defined via env or CLI
User-defined
Terminology (built-in, applied to all engagements)
Compound pentest-specific terms are replaced with neutral dev equivalents. Only unambiguous multi-word terms are replaced. Single words like target , payload , or exploit are left alone so legitimate code discussions don't break.
Run ./tanuki terms to see the full dictionary. The list lives in internal/tanuki/terminology.json and is embedded at build time; to customize it, edit the file and rebuild.
Container: scratch -based image with a single static Go binary plus CA certs. No OS, no shell, no runtime, ~10MB total. All capabilities dropped, read-only filesystem, no privilege escalation.
Everything runs in Docker. The proxy, hooks, and CLI all execute inside the container; hooks are invoked by Claude Code via docker exec , so no host-side binary is needed.
Matching: an Aho-Corasick automaton does longest-first replacement in a single pass over the stream.
JSON-aware: request and response bodies are rewritten string by string, never as raw text. Protocol fields ( model , role , type , ids, tool names) are left alone, base64 data is neither rewritten nor scanned, and thinking blocks go back to the API exactly as the model wrote them, since they are signed.
No HTTPS on the proxy. The proxy listens on plain HTTP. That's acceptable because it only accepts connections from localhost (Claude Code on the same machine); the upstream connection to the Anthropic API uses HTTPS.
Literal string matching only. Encoded forms (URL-encoded, base64, and similar) won't be caught by the Aho-Corasick automaton.
Terminology is one-directional. Pentest terms are rewritten real→fiction only. The model's responses use fiction terms, which pass through as-is.
Wildcard subdomains are not reversible. Every subdomain without its own mapping collapses onto a single wildcard value, so it cannot be mapped back. Auto-detection creates dedicated mappings for subdomains it sees, and those round-trip normally; the wildcard is the safety net for anything that slips past detection, and it favours hiding the hostname over keeping it usable.
Detection is the weak link, not the rewrite. Every outbound body is re-scanned after rewriting and refused if a real value survives, so a rewriter bug cannot leak. Nothing protects against a value the detector never recognised in the first place — the two limitations below are therefore the ones that matter.
Curated TLD list for auto-detection. Auto-detection uses a curated list of ~70 common TLDs to avoid false positives from code patterns like readme.md , foo.bar , or user.id . TLDs that collide with everyday code ( .info , .name , .app , .store , .top , .site , as in logger.info or user.name ) are left out. Domains on those or other unusual TLDs ( .pizza , .click , .it , .id ) can be added manually with ./tanuki add .
Ignored domains need an explicit add to become targets. github.com , google.com , and your own email domain are skipped by auto-detection so Claude Code's own git remotes, connectors, and account context don't pollute the table. The trade-off is that auto-detection alone will not map one of these, so on the rare engagement that targets GitHub or Google you must ./tanuki add github.com first; the residual check also skips ignored domains, so it will not catch the omission for you.
Thinking stays in fiction. Thinking blocks are signed by the API, so they are passed through unchanged in both directions. Thinking shown in Claude Code reads in fiction terms.
Streamed text is held back briefly. A fiction value can be streamed a few characters per event, so the proxy holds back the end of a text or tool-input block while it could still be the start of one, and sends it once the next event settles it. Display can lag by a few characters; the full value always arrives reversed.
tanuki

[truncated]

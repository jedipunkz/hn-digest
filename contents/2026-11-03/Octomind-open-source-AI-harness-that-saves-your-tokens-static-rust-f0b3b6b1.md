---
source: "https://github.com/Muvon/octomind"
hn_url: "https://news.ycombinator.com/item?id=50352898"
title: "Octomind – open-source AI harness that saves your tokens, static rust"
article_title: "GitHub - Muvon/octomind: Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. · GitHub"
image: "https://repository-images.githubusercontent.com/994166447/5deca376-0441-481a-83f3-48dcaeb7487c"
author: "donk8r"
captured_at: "2026-11-03T15:53:04Z"
capture_tool: "hn-digest"
hn_id: 50352898
score: 22
comments: 12
posted_at: "2026-11-03T15:02:58Z"
tags:
  - hacker-news
---

# Octomind – open-source AI harness that saves your tokens, static rust

- HN: [50352898](https://news.ycombinator.com/item?id=50352898)
- Source: [github.com](https://github.com/Muvon/octomind)
- Score: 22
- Comments: 12
- Posted: 2026-11-03T15:02:58Z

## Translation

Title: Octomind – open-source AI harness that saves your tokens, static rust
Article title: GitHub - Muvon/octomind: Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. · GitHub
Description: Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. - Muvon/octomind

Article text:
GitHub - Muvon/octomind: Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. · GitHub
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
Muvon
octomind Public Notifications You must be signed in to change notification settings
Star 325 ( 325 ) You must be signed in to star a repository
Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
51 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
2,733 Commits 2,733 Commits Folders and files
.cargo .cargo .github .github assets assets bench bench config-templates config-templates doc doc examples/ fix-a-bug examples/ fix-a-bug scripts scripts skills/ octomind skills/ octomind src src tests tests .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .noindex .noindex .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Cross.toml Cross.toml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md build.sh build.sh install.ps1 install.ps1 install.sh install.sh rustfmt.toml rustfmt.toml View all files Repository files navigation
The CLI-first AI coding agent runtime.
Pipe it, schedule it, embed it. One binary, any model, MCP-native —
built for autonomous work, not just chat.
Website · Docs · Tap Registry · Benchmarks · Issues
Octomind is a CLI-first AI coding agent runtime. It gives an LLM real access to your repo — reading and writing files, running shells, searching code, delegating to sub-agents — through MCP tools. The same session engine works interactively , piped through stdin , as a background daemon , over WebSocket , or as an ACP sub-agent inside another agent's stack.
# Interactive
octomind run developer:general
# Piped — CI, scripts, automation
echo " Explain the auth module " | octomind run developer:general --format plain
# Daemon — keep it alive, talk to it from another terminal
echo " watch the build " | octomind run --name watcher --daemon --format jsonl
octomind send --name watcher " run the test suite "
Why Octomind?
Most coding agents assume a human is watching the screen. Octomind is built for the work that keeps going after you leave.
Policy as code, not approval popups. Deterministic guardrails — pre-call guards, post-result hooks, post-turn validators — live in TOML and scripts, so they work in CI, daemons, and scheduled runs.
Proven on real PRs. It solved 24 of 25 held-out tasks at lower cost than the closest competitor. See the benchmarks .
Ready-made specialists. octomind run developer:general pulls a packaged agent (prompt, model, MCP servers, tool permissions) from a tap . Publish your own with one repo.
Sessions that stay sharp. Adaptive compression keeps long tasks moving while retaining task knowledge and the live exchange.
Any model, any budget. Use Octomind Cloud or your own key (OpenRouter, Anthropic, OpenAI, DeepSeek, Ollama, and more), pick a model per role, and set spend thresholds per request and per session. A request over its threshold stops; at the session threshold a terminal session asks y/N before continuing, while a non-interactive octomind run stops with exit code 75.
One config. Models, tools, roles, guardrails, budgets: all TOML, no framework code.
curl -fsSL https://raw.githubusercontent.com/muvon/octomind/master/install.sh | bash
On Windows, in PowerShell:
irm https: // raw.githubusercontent.com / muvon / octomind / master / install.ps1 | iex
On macOS, you can also install with Homebrew:
brew install muvon/tap/octomind
2. Get model access. The fastest path is Octomind Cloud, which includes models and needs no configuration:
octomind login
Prefer your own provider key? Skip login and export one (also OPENAI_API_KEY , ANTHROPIC_API_KEY , DEEPSEEK_API_KEY ), then follow Model access to point the config at it:
export OPENROUTER_API_KEY= " your_key "
3. Run a specialist in your project:
cd ~ /your/project
octomind run developer:general
You're now in a session with an agent that can read your code, run commands, and edit files. The interface is a plain line prompt with markdown rendering — no TUI to learn, so it works over SSH, in tmux , and in CI logs. The first time you run a specialist, it may install its tools and ask for credentials.
developer:general comes from the default tap muvon/tap , not your local config. Plain octomind run starts the config's default role, assistant:concierge .
Try the recorded bug-fix demo : two failing tests, an agent fix, and an independent test rerun.
Measured on octobench against 25 tasks harvested from merged pull requests across Python, PHP, Rust, C++, and JavaScript. Each agent works in the pre-fix repo; held-out tests from the merged fix decide pass or fail. These are the published July 30, 2026 results, not a measurement of this checkout.
The harness matters. opencode used the same model and endpoint; octomind solved 24 vs 19 at roughly half the cost. All four agents ran stock invocations with no tuning, so this measures the whole harness, not individual features.
Worst-case pricing, still ahead. glm-5.2 ran without prompt caching while Opus billed ~97% of context re-reads at cache rates — and octomind still led on solves, cost, and wall time.
Reproducible. Per-case table, artifacts, and reproduction guide: BENCHMARK.md . Background: blog post .
Drop a .agents/guardrails.toml in your repo. The runtime enforces it deterministically, with no code to compile and no plugin to install.
# Pre-call deny — block a class of calls before they execute
[[ guard ]]
match = " shell(command=^rm \\ s+-rf?) "
message = " rm -rf blocked. "
# Conditional rule — only fires after the agent ran git status this session
[[ guard ]]
match = " shell(command=git push) "
when = [ " +shell(command=git status) " ]
message = " Review changes before pushing. "
Layer
When it runs
What it does
Guards
Before a tool call
Deny by capability, argument regex, and session history. A match returns a denial instead of executing the call.
Hooks
After a tool result
A non-zero exit injects stdout into the agent's inbox — lint and format failures become automatic corrections without restarting the turn.
Validators
After a turn
Scripts the agent reads on its next turn. This replaces "approve this change?" prompts in autonomous loops.
Hook and validator script examples: Guardrails guide .
The same session engine, exposed however your workflow needs it.
Need machine-readable results? Pass --schema schema.json to constrain the final response to a JSON Schema (requires a model with structured-output support). See Structured output .
Feature
What you get
Learn more
Specialists & taps
Packaged agents with model, prompt, MCP servers, and permissions; installs their tool dependencies on first use
Tap system
Roles
Per-role model, temperature, system prompt, tools
Roles
Workflows
Multi-step pipelines, each step with its own model and toolset
Workflows
MCP client
stdio and Streamable HTTP servers, with OAuth; enable/disable servers mid-session
MCP tools
Intent-driven context
Skills and capabilities activate on demand, keeping context small
Skills
Local project tools
Executable scripts in .agents/tools/ auto-register as tools
Local tools
Supervisor & learning
External planning and gating; file-backed learning records
Supervisor , Learning
Sandbox
--sandbox restricts filesystem writes to the working directory
Config reference
Token efficiency
Cache-aware accounting and compaction
Token efficiency
Specialists
octomind run developer:general # general dev, language skills auto-activate
octomind run doctor:blood # blood-test interpretation specialist
octomind run lawyer:sg # legal specialist
Running a specialist fetches its manifest from the tap registry, installs any required binaries, resolves credentials, starts the right MCP servers, and then opens the session.
octomind tap init yourteam/my-tap # scaffold, validate, git-init, and register locally
octomind tap yourteam/my-tap # on another machine, after publishing to GitHub
The octomind- prefix is added automatically: yourteam/my-tap maps to the repository
yourteam/octomind-my-tap .
Method
Command
Install script (recommended)
curl -fsSL https://raw.githubusercontent.com/muvon/octomind/master/install.sh | bash
Install script (Windows PowerShell)
irm https://raw.githubusercontent.com/muvon/octomind/master/install.ps1 | iex
Homebrew (macOS)
brew install muvon/tap/octomind
Scoop (Windows)
scoop bucket add muvon https://github.com/Muvon/scoop-bucket then scoop install muvon/octomind
Cargo (Rust 1.95+)
cargo install octomind
Docker
docker run --rm ghcr.io/muvon/octomind --help
From source
git clone https://github.com/muvon/octomind.git && cd octomind && cargo build --release
The install script detects your OS and architecture (Linux and macOS on x86_64/aarch64) and installs to ~/.local/bin/ . Pass --install-dir <DIR> or set OCTOMIND_INSTALL_DIR to override. On Windows (x86_64/aarch64), install.ps1 installs to %LOCALAPPDATA%\octomind\bin and adds it to your user PATH ; it honors OCTOMIND_INSTALL_DIR and OCTOMIND_VERSION too. install.sh still works from Git Bash or MSYS2.
Command not found? Add the install directory to your PATH :
export PATH= " $HOME /.local/bin: $PATH "
Building from source requires protoc ; see Building from source for per-OS prerequisites.
octomind --version
octomind config # generate the default config
octomind run # start your first session
Shell completions: octomind completion <bash|zsh|fish|elvish|powershell> .
Octomind Cloud — octomind login (device-code sign-in, like gh auth login ) stores a gateway key locally. The default config already uses octohub:auto . Learn more .
Bring your own key — set the provider's environment variable ( OPENROUTER_API_KEY , OPENAI_API_KEY , ANTHROPIC_API_KEY , DEEPSEEK_API_KEY , …) and point the model tables at it. The supervisor and compression profiles also default to octohub:auto , so update all three:
[ model ]
name = " openrouter:google/gemini-2.5-flash "
[ supervisor . model ]
name = " openrouter:google/gemini-2.5-flash "
[ compression . model ]
name = " openrouter:google/gemini-2.5-flash "
Provider details and per-role overrides: Providers & models .
Config lives at ~/.local/share/octomind/config/config.toml on macOS and Linux, and under %LOCALAPPDATA%/octomind/config/ on Windows. Set OCTOMIND_DATA_DIR to override the data root.
octomind config --show # view current config
octomind config --validate # validate config
Ove

[truncated]

## Original Extract

Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. - Muvon/octomind

GitHub - Muvon/octomind: Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon. · GitHub
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
Muvon
octomind Public Notifications You must be signed in to change notification settings
Star 325 ( 325 ) You must be signed in to star a repository
Open-source AI coding agent and agent runtime: one binary, any model, MCP-native. Runs in terminal, CI, or as a daemon.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
51 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
2,733 Commits 2,733 Commits Folders and files
.cargo .cargo .github .github assets assets bench bench config-templates config-templates doc doc examples/ fix-a-bug examples/ fix-a-bug scripts scripts skills/ octomind skills/ octomind src src tests tests .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .noindex .noindex .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Cross.toml Cross.toml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md build.sh build.sh install.ps1 install.ps1 install.sh install.sh rustfmt.toml rustfmt.toml View all files Repository files navigation
The CLI-first AI coding agent runtime.
Pipe it, schedule it, embed it. One binary, any model, MCP-native —
built for autonomous work, not just chat.
Website · Docs · Tap Registry · Benchmarks · Issues
Octomind is a CLI-first AI coding agent runtime. It gives an LLM real access to your repo — reading and writing files, running shells, searching code, delegating to sub-agents — through MCP tools. The same session engine works interactively , piped through stdin , as a background daemon , over WebSocket , or as an ACP sub-agent inside another agent's stack.
# Interactive
octomind run developer:general
# Piped — CI, scripts, automation
echo " Explain the auth module " | octomind run developer:general --format plain
# Daemon — keep it alive, talk to it from another terminal
echo " watch the build " | octomind run --name watcher --daemon --format jsonl
octomind send --name watcher " run the test suite "
Why Octomind?
Most coding agents assume a human is watching the screen. Octomind is built for the work that keeps going after you leave.
Policy as code, not approval popups. Deterministic guardrails — pre-call guards, post-result hooks, post-turn validators — live in TOML and scripts, so they work in CI, daemons, and scheduled runs.
Proven on real PRs. It solved 24 of 25 held-out tasks at lower cost than the closest competitor. See the benchmarks .
Ready-made specialists. octomind run developer:general pulls a packaged agent (prompt, model, MCP servers, tool permissions) from a tap . Publish your own with one repo.
Sessions that stay sharp. Adaptive compression keeps long tasks moving while retaining task knowledge and the live exchange.
Any model, any budget. Use Octomind Cloud or your own key (OpenRouter, Anthropic, OpenAI, DeepSeek, Ollama, and more), pick a model per role, and set spend thresholds per request and per session. A request over its threshold stops; at the session threshold a terminal session asks y/N before continuing, while a non-interactive octomind run stops with exit code 75.
One config. Models, tools, roles, guardrails, budgets: all TOML, no framework code.
curl -fsSL https://raw.githubusercontent.com/muvon/octomind/master/install.sh | bash
On Windows, in PowerShell:
irm https: // raw.githubusercontent.com / muvon / octomind / master / install.ps1 | iex
On macOS, you can also install with Homebrew:
brew install muvon/tap/octomind
2. Get model access. The fastest path is Octomind Cloud, which includes models and needs no configuration:
octomind login
Prefer your own provider key? Skip login and export one (also OPENAI_API_KEY , ANTHROPIC_API_KEY , DEEPSEEK_API_KEY ), then follow Model access to point the config at it:
export OPENROUTER_API_KEY= " your_key "
3. Run a specialist in your project:
cd ~ /your/project
octomind run developer:general
You're now in a session with an agent that can read your code, run commands, and edit files. The interface is a plain line prompt with markdown rendering — no TUI to learn, so it works over SSH, in tmux , and in CI logs. The first time you run a specialist, it may install its tools and ask for credentials.
developer:general comes from the default tap muvon/tap , not your local config. Plain octomind run starts the config's default role, assistant:concierge .
Try the recorded bug-fix demo : two failing tests, an agent fix, and an independent test rerun.
Measured on octobench against 25 tasks harvested from merged pull requests across Python, PHP, Rust, C++, and JavaScript. Each agent works in the pre-fix repo; held-out tests from the merged fix decide pass or fail. These are the published July 30, 2026 results, not a measurement of this checkout.
The harness matters. opencode used the same model and endpoint; octomind solved 24 vs 19 at roughly half the cost. All four agents ran stock invocations with no tuning, so this measures the whole harness, not individual features.
Worst-case pricing, still ahead. glm-5.2 ran without prompt caching while Opus billed ~97% of context re-reads at cache rates — and octomind still led on solves, cost, and wall time.
Reproducible. Per-case table, artifacts, and reproduction guide: BENCHMARK.md . Background: blog post .
Drop a .agents/guardrails.toml in your repo. The runtime enforces it deterministically, with no code to compile and no plugin to install.
# Pre-call deny — block a class of calls before they execute
[[ guard ]]
match = " shell(command=^rm \\ s+-rf?) "
message = " rm -rf blocked. "
# Conditional rule — only fires after the agent ran git status this session
[[ guard ]]
match = " shell(command=git push) "
when = [ " +shell(command=git status) " ]
message = " Review changes before pushing. "
Layer
When it runs
What it does
Guards
Before a tool call
Deny by capability, argument regex, and session history. A match returns a denial instead of executing the call.
Hooks
After a tool result
A non-zero exit injects stdout into the agent's inbox — lint and format failures become automatic corrections without restarting the turn.
Validators
After a turn
Scripts the agent reads on its next turn. This replaces "approve this change?" prompts in autonomous loops.
Hook and validator script examples: Guardrails guide .
The same session engine, exposed however your workflow needs it.
Need machine-readable results? Pass --schema schema.json to constrain the final response to a JSON Schema (requires a model with structured-output support). See Structured output .
Feature
What you get
Learn more
Specialists & taps
Packaged agents with model, prompt, MCP servers, and permissions; installs their tool dependencies on first use
Tap system
Roles
Per-role model, temperature, system prompt, tools
Roles
Workflows
Multi-step pipelines, each step with its own model and toolset
Workflows
MCP client
stdio and Streamable HTTP servers, with OAuth; enable/disable servers mid-session
MCP tools
Intent-driven context
Skills and capabilities activate on demand, keeping context small
Skills
Local project tools
Executable scripts in .agents/tools/ auto-register as tools
Local tools
Supervisor & learning
External planning and gating; file-backed learning records
Supervisor , Learning
Sandbox
--sandbox restricts filesystem writes to the working directory
Config reference
Token efficiency
Cache-aware accounting and compaction
Token efficiency
Specialists
octomind run developer:general # general dev, language skills auto-activate
octomind run doctor:blood # blood-test interpretation specialist
octomind run lawyer:sg # legal specialist
Running a specialist fetches its manifest from the tap registry, installs any required binaries, resolves credentials, starts the right MCP servers, and then opens the session.
octomind tap init yourteam/my-tap # scaffold, validate, git-init, and register locally
octomind tap yourteam/my-tap # on another machine, after publishing to GitHub
The octomind- prefix is added automatically: yourteam/my-tap maps to the repository
yourteam/octomind-my-tap .
Method
Command
Install script (recommended)
curl -fsSL https://raw.githubusercontent.com/muvon/octomind/master/install.sh | bash
Install script (Windows PowerShell)
irm https://raw.githubusercontent.com/muvon/octomind/master/install.ps1 | iex
Homebrew (macOS)
brew install muvon/tap/octomind
Scoop (Windows)
scoop bucket add muvon https://github.com/Muvon/scoop-bucket then scoop install muvon/octomind
Cargo (Rust 1.95+)
cargo install octomind
Docker
docker run --rm ghcr.io/muvon/octomind --help
From source
git clone https://github.com/muvon/octomind.git && cd octomind && cargo build --release
The install script detects your OS and architecture (Linux and macOS on x86_64/aarch64) and installs to ~/.local/bin/ . Pass --install-dir <DIR> or set OCTOMIND_INSTALL_DIR to override. On Windows (x86_64/aarch64), install.ps1 installs to %LOCALAPPDATA%\octomind\bin and adds it to your user PATH ; it honors OCTOMIND_INSTALL_DIR and OCTOMIND_VERSION too. install.sh still works from Git Bash or MSYS2.
Command not found? Add the install directory to your PATH :
export PATH= " $HOME /.local/bin: $PATH "
Building from source requires protoc ; see Building from source for per-OS prerequisites.
octomind --version
octomind config # generate the default config
octomind run # start your first session
Shell completions: octomind completion <bash|zsh|fish|elvish|powershell> .
Octomind Cloud — octomind login (device-code sign-in, like gh auth login ) stores a gateway key locally. The default config already uses octohub:auto . Learn more .
Bring your own key — set the provider's environment variable ( OPENROUTER_API_KEY , OPENAI_API_KEY , ANTHROPIC_API_KEY , DEEPSEEK_API_KEY , …) and point the model tables at it. The supervisor and compression profiles also default to octohub:auto , so update all three:
[ model ]
name = " openrouter:google/gemini-2.5-flash "
[ supervisor . model ]
name = " openrouter:google/gemini-2.5-flash "
[ compression . model ]
name = " openrouter:google/gemini-2.5-flash "
Provider details and per-role overrides: Providers & models .
Config lives at ~/.local/share/octomind/config/config.toml on macOS and Linux, and under %LOCALAPPDATA%/octomind/config/ on Windows. Set OCTOMIND_DATA_DIR to override the data root.
octomind config --show # view current config
octomind config --validate # validate config
Ove

[truncated]

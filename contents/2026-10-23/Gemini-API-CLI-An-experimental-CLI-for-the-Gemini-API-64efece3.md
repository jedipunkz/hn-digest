---
source: "https://github.com/google-gemini/gemini-api-cli"
hn_url: "https://news.ycombinator.com/item?id=50212609"
title: "Gemini-API-CLI: An (experimental) CLI for the Gemini API"
article_title: "GitHub - google-gemini/gemini-api-cli: An (experimental) CLI for the Gemini API · GitHub"
image: "https://opengraph.githubassets.com/e9b91626ee4af698ba17710f2a2925dcba9d7368a6f685f1f2e7e1d27c9c8e73/google-gemini/gemini-api-cli"
author: "xnx"
captured_at: "2026-10-23T18:07:17Z"
capture_tool: "hn-digest"
hn_id: 50212609
score: 1
comments: 0
posted_at: "2026-10-23T17:36:15Z"
tags:
  - hacker-news
---

# Gemini-API-CLI: An (experimental) CLI for the Gemini API

- HN: [50212609](https://news.ycombinator.com/item?id=50212609)
- Source: [github.com](https://github.com/google-gemini/gemini-api-cli)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T17:36:15Z

## Translation

Title: Gemini-API-CLI: An (experimental) CLI for the Gemini API
Article title: GitHub - google-gemini/gemini-api-cli: An (experimental) CLI for the Gemini API · GitHub
Description: An (experimental) CLI for the Gemini API. Contribute to google-gemini/gemini-api-cli development by creating an account on GitHub.

Article text:
GitHub - google-gemini/gemini-api-cli: An (experimental) CLI for the Gemini API · GitHub
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
google-gemini
gemini-api-cli Public Notifications You must be signed in to change notification settings
Star 214 ( 214 ) You must be signed in to star a repository
An (experimental) CLI for the Gemini API
ai.google.dev/gemini-api/docs Topics
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
32 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
153 Commits 153 Commits Folders and files
.github/ workflows .github/ workflows cmd cmd evals/ harness-comparison evals/ harness-comparison internal internal scripts scripts tests/ contract tests/ contract .goreleaser.yaml .goreleaser.yaml LICENSE LICENSE README.md README.md USAGE.md USAGE.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Gemini models and managed agents for the agent you already use.
Give Codex, Claude Code, Antigravity, OpenClaw, Hermes, or any other agent with shell access the ability to use the Gemini API. Configure an API key, put gemini-api on PATH , and let your agent choose a model, generate an asset, analyze a file, or launch a managed agent to handle a task.
Your agent plans the work. Gemini provides models and managed agents. This stateless CLI connects them through standard shell commands:
gemini-api image " A friendly robot mascot for a developer tool " --out ./assets/
A prompt goes in. An image lands in your project. Your agent gets the file path and keeps building. The same binary can return an analysis, a voiceover, a transcript, or the interaction ID of a background managed-agent run.
Installation · Quick start · Command overview · Agent discovery · Workflows · Gemini API docs
A coding agent building an app may need a hero image, a voiceover, a second opinion on a design, or an answer inside a PDF. An assistant may need to turn a recording into captions or delegate multi-step research. Give that agent gemini-api and an API key, and it can call those capabilities directly when the task requires them.
One binary for models, media, and managed agents: Access text and reasoning models, multimodal file analysis, image/audio/music/video generation, and managed agents from the command line.
Delegate long-running work: Start a Gemini managed agent in the background, keep the returned interaction ID, and retrieve the result when needed.
Self-describing interface: Commands, flags, request schemas ( --schema ), and machine-readable usage ( --usage ) are built into the binary so an agent can inspect what it needs on the fly.
Safe dry-run validation: --dry-run validates inputs and prints the resolved HTTP request with credentials redacted—without making a network call or reading the OS keychain.
Script- and agent-friendly output: Stream text, write media directly to disk, emit structured JSON ( --output-format json ) or compact TOON ( --output-format toon ), or filter fields inline with --jq .
Stateless execution: Each invocation runs independently. Your agent supplies the context, chooses the model, and controls the next step.
Built for headless use by coding agents and assistant tools, with the same interface available to shell scripts and CI jobs. A single Go binary runs on macOS, Linux, and Windows.
curl -fsSL https://raw.githubusercontent.com/google-gemini/gemini-api-cli/main/scripts/install.sh | bash
Windows (PowerShell)
iwr - useb https: // raw.githubusercontent.com / google - gemini / gemini - api - cli / main / scripts / install.ps1 | iex
Go install
Install the latest release with Go:
go install github.com/google-gemini/gemini-api-cli/cmd/gemini-api@latest
To build the current development version, use @main instead of @latest . Current main requires Go 1.26.10 or later (Go's automatic toolchain selection downloads the required version if needed).
Download prebuilt archives for macOS, Linux, and Windows from GitHub Releases and place gemini-api on your PATH .
Create an API key in Google AI Studio (see the Gemini API key docs ) and make it available to your shell or your agent's subprocess environment:
export GEMINI_API_KEY= " YOUR_API_KEY "
Workstation tip: To persist your key in your OS keychain (macOS Keychain, Linux Secret Service, or Windows Credential Manager) or ~/.config/gemini-api/config.yaml , run gemini-api configure --api-key "YOUR_API_KEY" or open the interactive form with gemini-api configure --interactive . Check your active configuration and credential source at any time with gemini-api whoami .
Verify your setup locally with --dry-run (no network call), then remove --dry-run to call the API:
gemini-api generate " Hello from my agent " --dry-run
gemini-api generate " Hello from my agent "
2. Give your agent instructions
Add the following block to your agent's instructions ( AGENTS.md , CLAUDE.md , .cursorrules , or tool configuration):
You have access to gemini-api through the shell.
Use it to call Gemini models, generate or edit media, analyze files, or launch managed agents.
1. Inspect --help or --usage to choose a command, and --schema when building a --body JSON payload.
2. Preview API calls with --dry-run (or --dry-run --output-format json).
3. For discrete JSON output or --jq filtering on streaming/background commands (generate, agent run, agent status), pass --stream=false --output-format json.
4. Capture returned file paths (from media commands) and interaction IDs (from agent/model runs) to continue multi-step tasks.
3. Put your agent to work
Prompt your host agent as usual:
Build a landing page for this app. Use Gemini to create a mascot image and a welcome voiceover, and launch a managed research agent to investigate competing products. Incorporate the assets and findings into the page.
Your agent chooses the commands, runs them in its shell, and uses the returned file paths and interaction IDs in its own workflow.
Run gemini-api --help to see all top-level commands, or gemini-api <command> --help for command-specific flags and examples.
Agent discovery and machine interface
Everything an agent needs to use gemini-api is discoverable from the binary itself, without external documentation or credentials:
--usage outputs the complete command tree, flags, defaults, environment variables, and config keys as machine-readable KDL .
--schema prints the bundled JSON Schema (draft 2020-12) for any command that accepts a request body.
--dry-run validates flags and payloads and previews the HTTP request without credentials or network access (use --dry-run --output-format json for machine-parseable JSON request previews).
--no-interactive disables all interactive prompts and TUI forms.
--agent-mode (or GEMINI_CLI_AGENT_MODE=1 ) enables structured JSON error envelopes on stderr and defaults stdout to compact TOON output ( --output-format json explicitly selects JSON).
Exit codes distinguish success ( 0 ), runtime/API failures ( 1 ), invalid usage/validation errors ( 2 ), and authentication/authorization failures ( 3 ).
Run gemini-api --help-global to see all global authentication, network, output, and diagnostic flags.
1. Generate and edit creative assets
Coding agents like Codex, Claude Code, or Antigravity can generate images, speech, music, and video while building an app:
# Generate an image and capture the written file path
MASCOT_PATH= $( gemini-api image " A friendly robot mascot, soft studio lighting " --out ./assets/mascot.jpg )
# Edit an existing image with -i / --input
gemini-api image " Add a winter scarf and snowy background " -i " $MASCOT_PATH " --out ./assets/mascot-winter.jpg
# Generate single-speaker or multi-speaker voiceovers
gemini-api tts " Welcome. Let's build something together. " --out ./assets/welcome.wav
gemini-api tts " Alice: Welcome to the demo! Bob: Let's dive right in. " \
--multi-speaker " Alice=Kore,Bob=Puck " \
--out ./assets/dialogue.wav
# Generate background music and video
gemini-api music " Warm ambient music for a product walkthrough " --out ./assets/
gemini-api video " A slow camera move through a sunlit workshop " --out ./assets/
Media commands write files locally and print only the written file path to stdout . When --out points to a directory (or is omitted), the CLI names the file automatically using the file extension that matches the API's returned MIME type (for example, .jpg for image/jpeg , .wav for audio/wav , .mp3 for audio/mpeg , and .mp4 for video/mp4 ).
2. Analyze documents and transcribe recordings
Assistants like OpenClaw or Hermes can turn documents, media files, and YouTube URLs into answers and deliverables:
# Analyze a PDF or compare multiple images
gemini-api analyze --input product-brief.pdf " What should we build first, and why? " > plan.md
gemini-api analyze --input before.png --input after.png " Summarize the visual design changes "
# Transcribe audio or video (valid formats: md, text, json, srt)
gemini-api transcribe --input interview.mp4 --format srt --out interview.srt
For larger media files (over ~14 MB raw / 20 MB base64-encoded), upload once with files upload --wait and reuse the returned files/<id> URI across analyze , transcribe , generate , image , video , or agent run :
FILE_URI= $( gemini-api files upload interview.mp4 --wait --output-format json --jq ' .name ' )
gemini-api analyze --input " $FILE_URI " " List the key action items with timestamps "
3. Launch and poll a background managed agent
Delegate multi-step research or coding tasks to a Gemini managed agent such as Deep Research. Pass --background --stream=false so the CLI returns the interaction ID immediately instead of holding open an SSE stream:
INTERACTION_ID= $( gemini-api agent run " Research battery recycling approaches and cite sources " \
--agent deep-research-preview-04-2026 \
--background \
--stream=false \
--jq ' .id ' )
echo " Started interaction: $INTERACTION_ID "
Check progress or retrieve the finished output using the returned INTERACTION_ID :
gemini-api agent status " $INTERACTION_ID " --stream=false --output-format json
You can also stream live progress events ( gemini-api agent status "$INTERACTION_ID" --stream ) or cancel an active run ( gemini-api agent cancel "$INTERACTION_ID" ).
Use --model ( -m ) on generation and analysis commands to pick a model for reasoning, fast text, multimodal understanding, or media generation:
# View the curated offline model catalog grouped by capability
gemini-api models
# Query all live models available to your API key (--all auto-paginates across pages)
gemini-api models list --all
gemini-api models get gemini-3.8-flash
# Run a prompt with a specific model
gemini-api generate " Review this approach: cache immutable responses by content hash " --model gemini-3.8-flash
G

[truncated]

## Original Extract

An (experimental) CLI for the Gemini API. Contribute to google-gemini/gemini-api-cli development by creating an account on GitHub.

GitHub - google-gemini/gemini-api-cli: An (experimental) CLI for the Gemini API · GitHub
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
google-gemini
gemini-api-cli Public Notifications You must be signed in to change notification settings
Star 214 ( 214 ) You must be signed in to star a repository
An (experimental) CLI for the Gemini API
ai.google.dev/gemini-api/docs Topics
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
32 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
153 Commits 153 Commits Folders and files
.github/ workflows .github/ workflows cmd cmd evals/ harness-comparison evals/ harness-comparison internal internal scripts scripts tests/ contract tests/ contract .goreleaser.yaml .goreleaser.yaml LICENSE LICENSE README.md README.md USAGE.md USAGE.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Gemini models and managed agents for the agent you already use.
Give Codex, Claude Code, Antigravity, OpenClaw, Hermes, or any other agent with shell access the ability to use the Gemini API. Configure an API key, put gemini-api on PATH , and let your agent choose a model, generate an asset, analyze a file, or launch a managed agent to handle a task.
Your agent plans the work. Gemini provides models and managed agents. This stateless CLI connects them through standard shell commands:
gemini-api image " A friendly robot mascot for a developer tool " --out ./assets/
A prompt goes in. An image lands in your project. Your agent gets the file path and keeps building. The same binary can return an analysis, a voiceover, a transcript, or the interaction ID of a background managed-agent run.
Installation · Quick start · Command overview · Agent discovery · Workflows · Gemini API docs
A coding agent building an app may need a hero image, a voiceover, a second opinion on a design, or an answer inside a PDF. An assistant may need to turn a recording into captions or delegate multi-step research. Give that agent gemini-api and an API key, and it can call those capabilities directly when the task requires them.
One binary for models, media, and managed agents: Access text and reasoning models, multimodal file analysis, image/audio/music/video generation, and managed agents from the command line.
Delegate long-running work: Start a Gemini managed agent in the background, keep the returned interaction ID, and retrieve the result when needed.
Self-describing interface: Commands, flags, request schemas ( --schema ), and machine-readable usage ( --usage ) are built into the binary so an agent can inspect what it needs on the fly.
Safe dry-run validation: --dry-run validates inputs and prints the resolved HTTP request with credentials redacted—without making a network call or reading the OS keychain.
Script- and agent-friendly output: Stream text, write media directly to disk, emit structured JSON ( --output-format json ) or compact TOON ( --output-format toon ), or filter fields inline with --jq .
Stateless execution: Each invocation runs independently. Your agent supplies the context, chooses the model, and controls the next step.
Built for headless use by coding agents and assistant tools, with the same interface available to shell scripts and CI jobs. A single Go binary runs on macOS, Linux, and Windows.
curl -fsSL https://raw.githubusercontent.com/google-gemini/gemini-api-cli/main/scripts/install.sh | bash
Windows (PowerShell)
iwr - useb https: // raw.githubusercontent.com / google - gemini / gemini - api - cli / main / scripts / install.ps1 | iex
Go install
Install the latest release with Go:
go install github.com/google-gemini/gemini-api-cli/cmd/gemini-api@latest
To build the current development version, use @main instead of @latest . Current main requires Go 1.26.10 or later (Go's automatic toolchain selection downloads the required version if needed).
Download prebuilt archives for macOS, Linux, and Windows from GitHub Releases and place gemini-api on your PATH .
Create an API key in Google AI Studio (see the Gemini API key docs ) and make it available to your shell or your agent's subprocess environment:
export GEMINI_API_KEY= " YOUR_API_KEY "
Workstation tip: To persist your key in your OS keychain (macOS Keychain, Linux Secret Service, or Windows Credential Manager) or ~/.config/gemini-api/config.yaml , run gemini-api configure --api-key "YOUR_API_KEY" or open the interactive form with gemini-api configure --interactive . Check your active configuration and credential source at any time with gemini-api whoami .
Verify your setup locally with --dry-run (no network call), then remove --dry-run to call the API:
gemini-api generate " Hello from my agent " --dry-run
gemini-api generate " Hello from my agent "
2. Give your agent instructions
Add the following block to your agent's instructions ( AGENTS.md , CLAUDE.md , .cursorrules , or tool configuration):
You have access to gemini-api through the shell.
Use it to call Gemini models, generate or edit media, analyze files, or launch managed agents.
1. Inspect --help or --usage to choose a command, and --schema when building a --body JSON payload.
2. Preview API calls with --dry-run (or --dry-run --output-format json).
3. For discrete JSON output or --jq filtering on streaming/background commands (generate, agent run, agent status), pass --stream=false --output-format json.
4. Capture returned file paths (from media commands) and interaction IDs (from agent/model runs) to continue multi-step tasks.
3. Put your agent to work
Prompt your host agent as usual:
Build a landing page for this app. Use Gemini to create a mascot image and a welcome voiceover, and launch a managed research agent to investigate competing products. Incorporate the assets and findings into the page.
Your agent chooses the commands, runs them in its shell, and uses the returned file paths and interaction IDs in its own workflow.
Run gemini-api --help to see all top-level commands, or gemini-api <command> --help for command-specific flags and examples.
Agent discovery and machine interface
Everything an agent needs to use gemini-api is discoverable from the binary itself, without external documentation or credentials:
--usage outputs the complete command tree, flags, defaults, environment variables, and config keys as machine-readable KDL .
--schema prints the bundled JSON Schema (draft 2020-12) for any command that accepts a request body.
--dry-run validates flags and payloads and previews the HTTP request without credentials or network access (use --dry-run --output-format json for machine-parseable JSON request previews).
--no-interactive disables all interactive prompts and TUI forms.
--agent-mode (or GEMINI_CLI_AGENT_MODE=1 ) enables structured JSON error envelopes on stderr and defaults stdout to compact TOON output ( --output-format json explicitly selects JSON).
Exit codes distinguish success ( 0 ), runtime/API failures ( 1 ), invalid usage/validation errors ( 2 ), and authentication/authorization failures ( 3 ).
Run gemini-api --help-global to see all global authentication, network, output, and diagnostic flags.
1. Generate and edit creative assets
Coding agents like Codex, Claude Code, or Antigravity can generate images, speech, music, and video while building an app:
# Generate an image and capture the written file path
MASCOT_PATH= $( gemini-api image " A friendly robot mascot, soft studio lighting " --out ./assets/mascot.jpg )
# Edit an existing image with -i / --input
gemini-api image " Add a winter scarf and snowy background " -i " $MASCOT_PATH " --out ./assets/mascot-winter.jpg
# Generate single-speaker or multi-speaker voiceovers
gemini-api tts " Welcome. Let's build something together. " --out ./assets/welcome.wav
gemini-api tts " Alice: Welcome to the demo! Bob: Let's dive right in. " \
--multi-speaker " Alice=Kore,Bob=Puck " \
--out ./assets/dialogue.wav
# Generate background music and video
gemini-api music " Warm ambient music for a product walkthrough " --out ./assets/
gemini-api video " A slow camera move through a sunlit workshop " --out ./assets/
Media commands write files locally and print only the written file path to stdout . When --out points to a directory (or is omitted), the CLI names the file automatically using the file extension that matches the API's returned MIME type (for example, .jpg for image/jpeg , .wav for audio/wav , .mp3 for audio/mpeg , and .mp4 for video/mp4 ).
2. Analyze documents and transcribe recordings
Assistants like OpenClaw or Hermes can turn documents, media files, and YouTube URLs into answers and deliverables:
# Analyze a PDF or compare multiple images
gemini-api analyze --input product-brief.pdf " What should we build first, and why? " > plan.md
gemini-api analyze --input before.png --input after.png " Summarize the visual design changes "
# Transcribe audio or video (valid formats: md, text, json, srt)
gemini-api transcribe --input interview.mp4 --format srt --out interview.srt
For larger media files (over ~14 MB raw / 20 MB base64-encoded), upload once with files upload --wait and reuse the returned files/<id> URI across analyze , transcribe , generate , image , video , or agent run :
FILE_URI= $( gemini-api files upload interview.mp4 --wait --output-format json --jq ' .name ' )
gemini-api analyze --input " $FILE_URI " " List the key action items with timestamps "
3. Launch and poll a background managed agent
Delegate multi-step research or coding tasks to a Gemini managed agent such as Deep Research. Pass --background --stream=false so the CLI returns the interaction ID immediately instead of holding open an SSE stream:
INTERACTION_ID= $( gemini-api agent run " Research battery recycling approaches and cite sources " \
--agent deep-research-preview-04-2026 \
--background \
--stream=false \
--jq ' .id ' )
echo " Started interaction: $INTERACTION_ID "
Check progress or retrieve the finished output using the returned INTERACTION_ID :
gemini-api agent status " $INTERACTION_ID " --stream=false --output-format json
You can also stream live progress events ( gemini-api agent status "$INTERACTION_ID" --stream ) or cancel an active run ( gemini-api agent cancel "$INTERACTION_ID" ).
Use --model ( -m ) on generation and analysis commands to pick a model for reasoning, fast text, multimodal understanding, or media generation:
# View the curated offline model catalog grouped by capability
gemini-api models
# Query all live models available to your API key (--all auto-paginates across pages)
gemini-api models list --all
gemini-api models get gemini-3.8-flash
# Run a prompt with a specific model
gemini-api generate " Review this approach: cache immutable responses by content hash " --model gemini-3.8-flash
G

[truncated]

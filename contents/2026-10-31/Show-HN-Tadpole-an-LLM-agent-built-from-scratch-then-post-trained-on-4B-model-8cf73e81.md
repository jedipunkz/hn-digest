---
source: "https://github.com/shakewingo/Tadpole"
hn_url: "https://news.ycombinator.com/item?id=50315665"
title: "Show HN: Tadpole – an LLM agent built from scratch then post-trained on 4B model"
article_title: "GitHub - shakewingo/Tadpole: Build an LLM agent from scratch and continuously evolve it. · GitHub"
image: "https://repository-images.githubusercontent.com/1434586148/242cd6c1-897e-4073-803e-d43115a45a87"
author: "shakewingo"
captured_at: "2026-10-31T13:22:08Z"
capture_tool: "hn-digest"
hn_id: 50315665
score: 2
comments: 0
posted_at: "2026-10-31T12:43:42Z"
tags:
  - hacker-news
---

# Show HN: Tadpole – an LLM agent built from scratch then post-trained on 4B model

- HN: [50315665](https://news.ycombinator.com/item?id=50315665)
- Source: [github.com](https://github.com/shakewingo/Tadpole)
- Score: 2
- Comments: 0
- Posted: 2026-10-31T12:43:42Z

## Translation

Title: Show HN: Tadpole – an LLM agent built from scratch then post-trained on 4B model
Article title: GitHub - shakewingo/Tadpole: Build an LLM agent from scratch and continuously evolve it. · GitHub
Description: Build an LLM agent from scratch and continuously evolve it. - shakewingo/Tadpole

Article text:
GitHub - shakewingo/Tadpole: Build an LLM agent from scratch and continuously evolve it. · GitHub
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
shakewingo
Tadpole Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Build an LLM agent from scratch and continuously evolve it.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
79 Commits 79 Commits Folders and files
.github/ workflows .github/ workflows LICENSES LICENSES agent_from_scratch agent_from_scratch .gitignore .gitignore LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml View all files Repository files navigation
An LLM agent built from scratch, and grown one stage at a time.
A tool-using agent in about 3,500 lines of plain Python, with no agent framework, that you can
run around the clock from Telegram, and a stage-by-stage record of how it was built, measured
and post-trained.
An agent you can use. Files, shell, web search and fetch, sessions and compaction, on a
local 4B model or any OpenAI-compatible API, in a terminal or as a Telegram bot.
A from-scratch tutorial. Each stage adds one mechanism with tests and measured evidence,
from the first agent loop to SFT, expert iteration and a GRPO pilot. Post-training lifts
Qwen3-4B from 0.419 to 0.774 pass@1 on task types it never trained on
( results ).
Python 3.11 or later. Run every command from the repository root with the venv active.
git clone https://github.com/shakewingo/Tadpole.git
cd Tadpole
python3 -m venv .venv && source .venv/bin/activate
pip install -e .
Choose a model
Use a strong model through any OpenAI-compatible API: DeepSeek, OpenAI, OpenRouter, or your own
vLLM or LM Studio server. export these for the terminal, or put them in the
bot's env file :
export AGENT_BASE_URL=https://api.deepseek.com # the API root; often ends in /v1
export AGENT_MODEL=deepseek-flash # the provider's model ID
export AGENT_API_KEY=sk-...
export AGENT_N_CTX=32768 # optional: the model's context window
Run Qwen3-4B locally, base or post-trained : the model this project trains, for
studying the training rather than everyday use
Option
Variables
Ollama, any OS
AGENT_BASE_URL=http://localhost:11434/v1 , AGENT_MODEL=hf.co/unsloth/Qwen3-4B-Instruct-2507-GGUF:Q4_K_M , AGENT_N_CTX=32768
llama.cpp in-process, macOS; Qwen2 or Qwen3 GGUFs
leave AGENT_BASE_URL unset; AGENT_MODEL_PATH defaults to models/Qwen3-4B-Instruct-2507-Q4_K_M.gguf
A Tadpole adapter on llama.cpp, macOS
AGENT_MODEL_PATH=models/Qwen3-4B-Instruct-2507-Q8_0.gguf , AGENT_LORA_PATH=models/ei-3-lora-f16.gguf (or sft-1, sft-2); planning turns off, as in training
# Ollama: pull, then serve with a window large enough for the agent
ollama pull hf.co/unsloth/Qwen3-4B-Instruct-2507-GGUF:Q4_K_M
OLLAMA_CONTEXT_LENGTH=32768 ollama serve
# llama.cpp: the default AGENT_MODEL_PATH
hf download unsloth/Qwen3-4B-Instruct-2507-GGUF Qwen3-4B-Instruct-2507-Q4_K_M.gguf --local-dir models
# An adapter, on the Q8_0 base it was measured on
hf download unsloth/Qwen3-4B-Instruct-2507-GGUF Qwen3-4B-Instruct-2507-Q8_0.gguf --local-dir models
hf download Shakewin/tadpole-qwen3-4b-ei-3 ei-3-lora-f16.gguf --local-dir models
llama.cpp keeps the model in memory while the agent runs: about 8.5 GB for the Q8_0 base with its
32k window, which slows a 16 GB Mac.
python -m agent_from_scratch.agent
Chat from Telegram
Message @BotFather with /newbot for a bot token , and
@userinfobot for your numeric user ID .
Put both, and your AGENT_* variables if you use an API, in the env file:
cp agent_from_scratch/deploy/telegram-bot.env.example agent_from_scratch/deploy/telegram-bot.env
Start it, and message your bot: sh agent_from_scratch/scripts/run_telegram_bot.sh
The agent runs shell commands and edits files with your user's permissions. It answers only
TELEGRAM_ALLOWED_USER_ID , but anyone with the token controls the bot: keep the env file private.
In Telegram and the terminal: /new starts a session, /reset clears it, /session <id>
switches to a named one, and /compact summarizes the history.
Use an always-on machine: a Mac that does not sleep, a home server, or a small cloud VM with an
API model (1 GB is enough).
# Linux (systemd)
UNIT= ~ /.config/systemd/user/tadpole-bot.service
mkdir -p " $( dirname " $UNIT " ) " && cp agent_from_scratch/deploy/tadpole-bot.service " $UNIT "
sed -i " s|/REPLACE/WITH/REPO/PATH| $PWD |g " " $UNIT "
systemctl --user enable --now tadpole-bot && sudo loginctl enable-linger " $USER "
# macOS (launchd); also turn on "Prevent automatic sleeping when the display is off"
PLIST= ~ /Library/LaunchAgents/com.tadpole.telegram-bot.plist
cp agent_from_scratch/deploy/com.tadpole.telegram-bot.plist " $PLIST "
sed -i ' ' " s|/REPLACE/WITH/REPO/PATH| $PWD |g " " $PLIST "
mkdir -p outputs && launchctl load " $PLIST "
Logs: journalctl --user -u tadpole-bot -f on Linux, outputs/telegram_bot.log on macOS.
Each stage adds one mechanism, with a deliberate failure, focused tests (601, none needing a
model) and evidence from real model runs.
Start with the walkthrough notebook (outputs
saved, so it reads on GitHub), then the build log for
each stage's evidence and evaluating and post-training
to run it yourself. Training needs a 24 GB CUDA GPU; the whole post-training line cost about
$140 on RunPod.
The agent's code, prompts and tools were frozen, with planning off (no update_plan tool); only
the weights changed. Every choice was made
on val, and each final model ran once on two held-out sets whose task types no training set
contains: 140 single-turn tasks (test) and 80 leveled tasks (hard test), three samples each.
Bold is best, italic second. SFT-1: verified teacher trajectories plus When2Call and Toucan
rows. SFT-2: plus multi-turn demonstrations. EI-3: SFT-2 after three rounds of expert iteration.
† A local adaptation of BFCL's multi-turn base category, not a leaderboard score.
SFT on verified demonstrations gives the large gain , and false "done" claims fall from 31
to 1 in 420 runs.
The three post-trained models are level on held-out tasks (0.648, 0.648, 0.645 over all
220); EI-3 is the most balanced, and beats SFT-2 on BFCL (p = 0.034).
Val overstated expert iteration (0.752 → 0.852), and GRPO lost to it at equal rollouts.
Neither learned a behaviour the model never produced.
One training run per model, on a small local benchmark: not a claim about general agent ability.
Full tables, paired tests and limitations: docs/results.md .
Download the adapters and data
Download all adapters and training data from Hugging Face .
Each card includes serving or data-reproduction instructions.
Memory across sessions, with remember, search and forget tools.
Skills and MCP behind the same tool registry.
Background jobs for the always-on bot, and a read-only sub-agent.
A self-improving harness : with the model fixed, let execution feedback improve the prompts
and tools.
The tool implementations adapt parts of nanobot
( attribution ). Thanks to the authors of
Qwen3 ,
When2Call ,
Toucan and
BFCL .
MIT , Copyright (c) 2026 shakewingo. Third-party material keeps its own license (the
nanobot tools MIT, the Qwen-derived chat template Apache-2.0), and the Hugging Face adapters are
Apache-2.0, the SFT data MIT; see Third-party notices .
Build an LLM agent from scratch and continuously evolve it.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Build an LLM agent from scratch and continuously evolve it. - shakewingo/Tadpole

GitHub - shakewingo/Tadpole: Build an LLM agent from scratch and continuously evolve it. · GitHub
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
shakewingo
Tadpole Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Build an LLM agent from scratch and continuously evolve it.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
79 Commits 79 Commits Folders and files
.github/ workflows .github/ workflows LICENSES LICENSES agent_from_scratch agent_from_scratch .gitignore .gitignore LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md pyproject.toml pyproject.toml View all files Repository files navigation
An LLM agent built from scratch, and grown one stage at a time.
A tool-using agent in about 3,500 lines of plain Python, with no agent framework, that you can
run around the clock from Telegram, and a stage-by-stage record of how it was built, measured
and post-trained.
An agent you can use. Files, shell, web search and fetch, sessions and compaction, on a
local 4B model or any OpenAI-compatible API, in a terminal or as a Telegram bot.
A from-scratch tutorial. Each stage adds one mechanism with tests and measured evidence,
from the first agent loop to SFT, expert iteration and a GRPO pilot. Post-training lifts
Qwen3-4B from 0.419 to 0.774 pass@1 on task types it never trained on
( results ).
Python 3.11 or later. Run every command from the repository root with the venv active.
git clone https://github.com/shakewingo/Tadpole.git
cd Tadpole
python3 -m venv .venv && source .venv/bin/activate
pip install -e .
Choose a model
Use a strong model through any OpenAI-compatible API: DeepSeek, OpenAI, OpenRouter, or your own
vLLM or LM Studio server. export these for the terminal, or put them in the
bot's env file :
export AGENT_BASE_URL=https://api.deepseek.com # the API root; often ends in /v1
export AGENT_MODEL=deepseek-flash # the provider's model ID
export AGENT_API_KEY=sk-...
export AGENT_N_CTX=32768 # optional: the model's context window
Run Qwen3-4B locally, base or post-trained : the model this project trains, for
studying the training rather than everyday use
Option
Variables
Ollama, any OS
AGENT_BASE_URL=http://localhost:11434/v1 , AGENT_MODEL=hf.co/unsloth/Qwen3-4B-Instruct-2507-GGUF:Q4_K_M , AGENT_N_CTX=32768
llama.cpp in-process, macOS; Qwen2 or Qwen3 GGUFs
leave AGENT_BASE_URL unset; AGENT_MODEL_PATH defaults to models/Qwen3-4B-Instruct-2507-Q4_K_M.gguf
A Tadpole adapter on llama.cpp, macOS
AGENT_MODEL_PATH=models/Qwen3-4B-Instruct-2507-Q8_0.gguf , AGENT_LORA_PATH=models/ei-3-lora-f16.gguf (or sft-1, sft-2); planning turns off, as in training
# Ollama: pull, then serve with a window large enough for the agent
ollama pull hf.co/unsloth/Qwen3-4B-Instruct-2507-GGUF:Q4_K_M
OLLAMA_CONTEXT_LENGTH=32768 ollama serve
# llama.cpp: the default AGENT_MODEL_PATH
hf download unsloth/Qwen3-4B-Instruct-2507-GGUF Qwen3-4B-Instruct-2507-Q4_K_M.gguf --local-dir models
# An adapter, on the Q8_0 base it was measured on
hf download unsloth/Qwen3-4B-Instruct-2507-GGUF Qwen3-4B-Instruct-2507-Q8_0.gguf --local-dir models
hf download Shakewin/tadpole-qwen3-4b-ei-3 ei-3-lora-f16.gguf --local-dir models
llama.cpp keeps the model in memory while the agent runs: about 8.5 GB for the Q8_0 base with its
32k window, which slows a 16 GB Mac.
python -m agent_from_scratch.agent
Chat from Telegram
Message @BotFather with /newbot for a bot token , and
@userinfobot for your numeric user ID .
Put both, and your AGENT_* variables if you use an API, in the env file:
cp agent_from_scratch/deploy/telegram-bot.env.example agent_from_scratch/deploy/telegram-bot.env
Start it, and message your bot: sh agent_from_scratch/scripts/run_telegram_bot.sh
The agent runs shell commands and edits files with your user's permissions. It answers only
TELEGRAM_ALLOWED_USER_ID , but anyone with the token controls the bot: keep the env file private.
In Telegram and the terminal: /new starts a session, /reset clears it, /session <id>
switches to a named one, and /compact summarizes the history.
Use an always-on machine: a Mac that does not sleep, a home server, or a small cloud VM with an
API model (1 GB is enough).
# Linux (systemd)
UNIT= ~ /.config/systemd/user/tadpole-bot.service
mkdir -p " $( dirname " $UNIT " ) " && cp agent_from_scratch/deploy/tadpole-bot.service " $UNIT "
sed -i " s|/REPLACE/WITH/REPO/PATH| $PWD |g " " $UNIT "
systemctl --user enable --now tadpole-bot && sudo loginctl enable-linger " $USER "
# macOS (launchd); also turn on "Prevent automatic sleeping when the display is off"
PLIST= ~ /Library/LaunchAgents/com.tadpole.telegram-bot.plist
cp agent_from_scratch/deploy/com.tadpole.telegram-bot.plist " $PLIST "
sed -i ' ' " s|/REPLACE/WITH/REPO/PATH| $PWD |g " " $PLIST "
mkdir -p outputs && launchctl load " $PLIST "
Logs: journalctl --user -u tadpole-bot -f on Linux, outputs/telegram_bot.log on macOS.
Each stage adds one mechanism, with a deliberate failure, focused tests (601, none needing a
model) and evidence from real model runs.
Start with the walkthrough notebook (outputs
saved, so it reads on GitHub), then the build log for
each stage's evidence and evaluating and post-training
to run it yourself. Training needs a 24 GB CUDA GPU; the whole post-training line cost about
$140 on RunPod.
The agent's code, prompts and tools were frozen, with planning off (no update_plan tool); only
the weights changed. Every choice was made
on val, and each final model ran once on two held-out sets whose task types no training set
contains: 140 single-turn tasks (test) and 80 leveled tasks (hard test), three samples each.
Bold is best, italic second. SFT-1: verified teacher trajectories plus When2Call and Toucan
rows. SFT-2: plus multi-turn demonstrations. EI-3: SFT-2 after three rounds of expert iteration.
† A local adaptation of BFCL's multi-turn base category, not a leaderboard score.
SFT on verified demonstrations gives the large gain , and false "done" claims fall from 31
to 1 in 420 runs.
The three post-trained models are level on held-out tasks (0.648, 0.648, 0.645 over all
220); EI-3 is the most balanced, and beats SFT-2 on BFCL (p = 0.034).
Val overstated expert iteration (0.752 → 0.852), and GRPO lost to it at equal rollouts.
Neither learned a behaviour the model never produced.
One training run per model, on a small local benchmark: not a claim about general agent ability.
Full tables, paired tests and limitations: docs/results.md .
Download the adapters and data
Download all adapters and training data from Hugging Face .
Each card includes serving or data-reproduction instructions.
Memory across sessions, with remember, search and forget tools.
Skills and MCP behind the same tool registry.
Background jobs for the always-on bot, and a read-only sub-agent.
A self-improving harness : with the model fixed, let execution feedback improve the prompts
and tools.
The tool implementations adapt parts of nanobot
( attribution ). Thanks to the authors of
Qwen3 ,
When2Call ,
Toucan and
BFCL .
MIT , Copyright (c) 2026 shakewingo. Third-party material keeps its own license (the
nanobot tools MIT, the Qwen-derived chat template Apache-2.0), and the Hugging Face adapters are
Apache-2.0, the SFT data MIT; see Third-party notices .
Build an LLM agent from scratch and continuously evolve it.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

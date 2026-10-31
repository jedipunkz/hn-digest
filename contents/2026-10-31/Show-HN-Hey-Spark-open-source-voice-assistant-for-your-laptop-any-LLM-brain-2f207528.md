---
source: "https://github.com/contextuai/hey-spark"
hn_url: "https://news.ycombinator.com/item?id=50318453"
title: "Show HN: Hey Spark – open-source voice assistant for your laptop, any LLM brain"
article_title: "GitHub - contextuai/hey-spark: Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors · GitHub"
image: "https://repository-images.githubusercontent.com/1437836939/52e750d2-22af-4fb7-9958-f43df3e95636"
author: "nagen"
captured_at: "2026-10-31T18:14:09Z"
capture_tool: "hn-digest"
hn_id: 50318453
score: 1
comments: 0
posted_at: "2026-10-31T18:09:19Z"
tags:
  - hacker-news
---

# Show HN: Hey Spark – open-source voice assistant for your laptop, any LLM brain

- HN: [50318453](https://news.ycombinator.com/item?id=50318453)
- Source: [github.com](https://github.com/contextuai/hey-spark)
- Score: 1
- Comments: 0
- Posted: 2026-10-31T18:09:19Z

## Translation

Title: Show HN: Hey Spark – open-source voice assistant for your laptop, any LLM brain
Article title: GitHub - contextuai/hey-spark: Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors · GitHub
Description: Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors - contextuai/hey-spark

Article text:
GitHub - contextuai/hey-spark: Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors · GitHub
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
contextuai
hey-spark Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors
Readme License Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
7 Commits 7 Commits Folders and files
.claude .claude .github .github LICENSES LICENSES assets assets docs docs packaging packaging scripts scripts spark spark tests tests tools tools .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md config.yaml config.yaml mcp_servers.example.json mcp_servers.example.json mcp_servers.json mcp_servers.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An always-listening voice assistant for your laptop. Say "Hey Spark" , ask a question or
give a command, and hear the answer. The wake word, speech recognition and voice all run
locally on your CPU. The brain is Claude, Hermes Agent ,
or any OpenAI-compatible model, including a local one through Ollama. Spark reaches your tools (GitHub, Gmail, Calendar, your own systems) through
MCP servers, and it asks out loud before doing anything that
changes something.
"Hey Spark, what's on my calendar tomorrow?"
"Set a pasta timer for ten minutes."
"Summarize the open pull requests on hey-spark."
"Let's start a new project called weather-bot." → "I'll create a folder, a git repo and a private GitHub repo called weather-bot. Should I go ahead?"
Local wake word. A bundled openWakeWord model
for "Hey Spark". Train your own wake word for free in Google Colab, or use
Picovoice Porcupine.
Local speech. faster-whisper speech-to-text and
Piper or Kokoro
text-to-speech, all CPU-only. Deepgram and ElevenLabs are optional cloud backends.
Your choice of brain. Claude through the
Claude Agent SDK , your own
Hermes Agent , or any OpenAI-compatible API:
OpenAI, Azure OpenAI, OpenRouter, Groq, Together, Ollama, LM Studio or vLLM.
Connectors over MCP , with per-server allowlists. Read-only tools run freely; anything that
sends, creates or updates needs a spoken "yes" , and destructive tools are blocked by default.
Conversation that feels natural. Spark streams its answer sentence by sentence, stays awake
for follow-ups without the wake word, ignores speech that isn't meant for it, and can be
interrupted by saying the wake word (barge-in).
A personality of your choice. Plain and brief by default, or a composed British butler
with a dry wit who calls you "sir", answers the wake word with "Yes, sir?" and speaks in a
British voice. See Personality and voice .
Memory and shortcuts. It keeps editable Markdown notes about each connected system, and you
can map a short phrase ("good morning") to a detailed prompt.
Desktop app for macOS and Windows: a single Start/Stop button, a light that shows what
Spark is doing, and windows for Connectors, Shortcuts, the Brain, and Personality & Voice.
Everything is swappable from config.yaml : wake word, VAD, STT, TTS and brain backends,
with no code changes.
flowchart LR
mic[Mic 16 kHz] --> wake[Wake word<br/>openWakeWord]
wake --> vad[Record until silence<br/>Silero VAD]
vad --> stt[Speech to text<br/>faster-whisper]
stt --> brain[Brain<br/>Claude or OpenAI-compatible]
brain <--> mcp[MCP connectors<br/>+ spoken confirmation]
brain --> split[Sentence splitter]
split --> tts[Text to speech<br/>Piper / Kokoro]
tts --> spk[Speaker]
Loading
Nothing leaves your machine before the wake word. After it, only the text of your request goes
to the brain you chose. Audio leaves the machine only if you pick a cloud STT backend.
Requirements: macOS 14 or later on Apple Silicon, or Windows 11 (see
Platform support ), a microphone, about 1 GB of disk for models, and a brain:
Claude Code signed in to your Claude account (or with an
Anthropic API key ), an OpenAI-compatible key, or
Ollama running locally. A USB speakerphone with echo cancellation works
best.
Get the latest build from Releases . Each file
has a .sha256 checksum next to it.
macOS (Apple Silicon, macOS 14+)
Download HeySpark-<version>-macos-arm64.dmg , open it and drag HeySpark into
Applications .
Open it from Applications. If macOS says it can't verify the app, open
System Settings → Privacy & Security , scroll down and click Open Anyway . (Or run
xattr -dr com.apple.quarantine /Applications/HeySpark.app once in Terminal.)
Allow microphone access when macOS asks.
Settings, logs, models and your .env live in ~/Library/Application Support/HeySpark .
Download and run HeySpark-Setup-<version>.exe . It installs per-user and needs no admin
rights. The installer isn't code-signed yet, so Windows SmartScreen may warn you: choose
More info → Run anyway .
Open Hey Spark from the Start menu.
Settings, logs, models and your .env live in %LOCALAPPDATA%\HeySpark .
Click the brain button, pick a provider and enter
your key (or pick Ollama).
Press Start . The first start downloads the speech models (about 200 MB). Then say
"Hey Spark" .
# 1. Install uv (https://docs.astral.sh/uv/); it manages Python 3.11 for you
git clone https://github.com/contextuai/hey-spark.git
cd hey-spark
uv sync --all-extras # plain `uv sync` skips the optional backends
# 2. Download the speech models into ./.models (about 200 MB)
uv run spark models download
# 3. Add your brain's key
cp .env.example .env # (Windows: copy) then set your key; see "Choose a brain"
# 4. Check everything, then talk to it
uv run spark doctor # every required check should show ✔
uv run spark # say "Hey Spark, what's the capital of Japan?"
No microphone handy? uv run spark --text lets you type questions and read answers, and
uv run spark --brain mock --text runs fully offline with no key at all.
Brain
Set up
Good for
Claude (default)
Claude Code installed and signed in (or ANTHROPIC_API_KEY in .env )
The best tool use and connector work. Supports the researcher sub-agent and fast/deep model routing
Hermes Agent
Hermes Agent with its API server on, and its API_SERVER_KEY
Giving your Hermes agent a voice: it answers with its own model, tools, memory and skills
OpenAI-compatible cloud (OpenAI, Azure, OpenRouter, Groq, Together)
The provider's key in .env , and brain.backend: openai_compat
Using a model or account you already have
Local (Ollama, LM Studio, vLLM)
Run the server, pull a tool-capable model (e.g. qwen2.5:7b )
Fully local : with faster-whisper and Piper, nothing leaves your machine
The easiest way to switch is the desktop app's Brain window, which also tests the connection
and checks that the model can call tools. See docs/openai_brain.md for the
provider presets and the config.yaml settings.
With Claude, brain.mode picks the models: auto (default) sends quick requests to
brain.fast_model and bigger ones to brain.deep_model ; deep answers everything with
brain.deep_model and hands tool-heavy lookups to a researcher on a faster model (the best
answers, and the highest API cost). You can also switch by voice
("use the deep model", "back to auto").
The default config.yaml uses the bundled model assets/wakewords/hey_spark.onnx . Nothing to set
up. If it misses you or wakes on its own, tune wakeword.options.openwakeword.threshold .
Want your own wake word? tools/wakeword_trainer.ipynb trains
one for any phrase on a free Colab GPU in about an hour, with no recordings needed.
docs/wakeword.md walks through it step by step, plus the Porcupine
alternative.
Follow-ups: after a reply, Spark stays awake for 3 minutes ( attention.awake_s ), so you can
keep talking without the wake word. It answers only speech that's meant for it, then plays a
soft chime when it falls asleep.
Commands: "new topic" (fresh conversation), "go to sleep", "never mind", and "mute
yourself" (it then stops listening until you click the mic button in the window). Film-style
ones work too: "stand down", "belay that", "clean slate protocol", "full power", "go dark".
Timers and time: "set a pasta timer for 10 minutes", "what timers do I have", "what time is it".
New projects: "let's start a new project called weather-bot". After your yes , Spark creates
projects.root_dir/weather-bot with git and a first commit, a private GitHub repo and a tracking
issue through your gh login (it never pushes), and can open Claude Code in the folder.
Shortcuts: map a phrase like "good morning" to a detailed prompt in the Shortcuts window
(or shortcuts.json ). Fifteen come built in: "rise and shine" for the morning briefing,
"what's next?", "sitrep", "check the post", "run a diagnostic", "house party protocol" and more.
See docs/commands.md .
Barge-in: with app.barge_in: true , saying the wake word while Spark is talking stops it.
Turn this on only with an echo-cancelling speakerphone.
Earcons: a chime on wake, a tick while thinking, a low tone on errors.
Personality and voice: pick the butler personality and a British voice in the Personality &
Voice window (or app.persona: butler in config.yaml ). See docs/persona.md .
Connectors are MCP servers defined in mcp_servers.json (with ${VAR} placeholders only; the
values come from .env ). Add them by hand or in the desktop app's Connectors window. Each
server's tool policy lives in config.yaml under connectors: :
allowed_tools are read-only and pre-approved.
side_effecting tools (send, create, update) need a spoken "yes" every time. Spark says one
line about what it will do and waits. Silence, an unclear answer or "no" means it doesn't happen.
destructive tools are blocked unless you set allow_destructive: true for that server.
Any tool on none of these lists is denied.
Create a fine-grained token for the
repos Spark may see: Issues: Read and write , Pull requests: Read and Contents: Read .
In the Connectors window, open github (under From config files ) and click Edit
config . Replace ${SPARK_GITHUB_PAT} with your token and press Save and test . Spark
moves the token to .env .
{
"github" : {
"type" : " http " ,
"url" : " https://api.githubcopilot.com/mcp/ " ,
"headers" : { "Authorization" : " Bearer YOUR_TOKEN " }
}
}
Switch it on. Spark uses GitHub's own labels for the tools: reading runs freely, creating an
issu

[truncated]

## Original Extract

Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors - contextuai/hey-spark

GitHub - contextuai/hey-spark: Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors · GitHub
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
contextuai
hey-spark Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Hey Spark, by ContextuAI: a wake-word voice assistant with Claude, Hermes Agent or any OpenAI-compatible model as the brain and MCP servers as connectors
Readme License Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
7 Commits 7 Commits Folders and files
.claude .claude .github .github LICENSES LICENSES assets assets docs docs packaging packaging scripts scripts spark spark tests tests tools tools .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md config.yaml config.yaml mcp_servers.example.json mcp_servers.example.json mcp_servers.json mcp_servers.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
An always-listening voice assistant for your laptop. Say "Hey Spark" , ask a question or
give a command, and hear the answer. The wake word, speech recognition and voice all run
locally on your CPU. The brain is Claude, Hermes Agent ,
or any OpenAI-compatible model, including a local one through Ollama. Spark reaches your tools (GitHub, Gmail, Calendar, your own systems) through
MCP servers, and it asks out loud before doing anything that
changes something.
"Hey Spark, what's on my calendar tomorrow?"
"Set a pasta timer for ten minutes."
"Summarize the open pull requests on hey-spark."
"Let's start a new project called weather-bot." → "I'll create a folder, a git repo and a private GitHub repo called weather-bot. Should I go ahead?"
Local wake word. A bundled openWakeWord model
for "Hey Spark". Train your own wake word for free in Google Colab, or use
Picovoice Porcupine.
Local speech. faster-whisper speech-to-text and
Piper or Kokoro
text-to-speech, all CPU-only. Deepgram and ElevenLabs are optional cloud backends.
Your choice of brain. Claude through the
Claude Agent SDK , your own
Hermes Agent , or any OpenAI-compatible API:
OpenAI, Azure OpenAI, OpenRouter, Groq, Together, Ollama, LM Studio or vLLM.
Connectors over MCP , with per-server allowlists. Read-only tools run freely; anything that
sends, creates or updates needs a spoken "yes" , and destructive tools are blocked by default.
Conversation that feels natural. Spark streams its answer sentence by sentence, stays awake
for follow-ups without the wake word, ignores speech that isn't meant for it, and can be
interrupted by saying the wake word (barge-in).
A personality of your choice. Plain and brief by default, or a composed British butler
with a dry wit who calls you "sir", answers the wake word with "Yes, sir?" and speaks in a
British voice. See Personality and voice .
Memory and shortcuts. It keeps editable Markdown notes about each connected system, and you
can map a short phrase ("good morning") to a detailed prompt.
Desktop app for macOS and Windows: a single Start/Stop button, a light that shows what
Spark is doing, and windows for Connectors, Shortcuts, the Brain, and Personality & Voice.
Everything is swappable from config.yaml : wake word, VAD, STT, TTS and brain backends,
with no code changes.
flowchart LR
mic[Mic 16 kHz] --> wake[Wake word<br/>openWakeWord]
wake --> vad[Record until silence<br/>Silero VAD]
vad --> stt[Speech to text<br/>faster-whisper]
stt --> brain[Brain<br/>Claude or OpenAI-compatible]
brain <--> mcp[MCP connectors<br/>+ spoken confirmation]
brain --> split[Sentence splitter]
split --> tts[Text to speech<br/>Piper / Kokoro]
tts --> spk[Speaker]
Loading
Nothing leaves your machine before the wake word. After it, only the text of your request goes
to the brain you chose. Audio leaves the machine only if you pick a cloud STT backend.
Requirements: macOS 14 or later on Apple Silicon, or Windows 11 (see
Platform support ), a microphone, about 1 GB of disk for models, and a brain:
Claude Code signed in to your Claude account (or with an
Anthropic API key ), an OpenAI-compatible key, or
Ollama running locally. A USB speakerphone with echo cancellation works
best.
Get the latest build from Releases . Each file
has a .sha256 checksum next to it.
macOS (Apple Silicon, macOS 14+)
Download HeySpark-<version>-macos-arm64.dmg , open it and drag HeySpark into
Applications .
Open it from Applications. If macOS says it can't verify the app, open
System Settings → Privacy & Security , scroll down and click Open Anyway . (Or run
xattr -dr com.apple.quarantine /Applications/HeySpark.app once in Terminal.)
Allow microphone access when macOS asks.
Settings, logs, models and your .env live in ~/Library/Application Support/HeySpark .
Download and run HeySpark-Setup-<version>.exe . It installs per-user and needs no admin
rights. The installer isn't code-signed yet, so Windows SmartScreen may warn you: choose
More info → Run anyway .
Open Hey Spark from the Start menu.
Settings, logs, models and your .env live in %LOCALAPPDATA%\HeySpark .
Click the brain button, pick a provider and enter
your key (or pick Ollama).
Press Start . The first start downloads the speech models (about 200 MB). Then say
"Hey Spark" .
# 1. Install uv (https://docs.astral.sh/uv/); it manages Python 3.11 for you
git clone https://github.com/contextuai/hey-spark.git
cd hey-spark
uv sync --all-extras # plain `uv sync` skips the optional backends
# 2. Download the speech models into ./.models (about 200 MB)
uv run spark models download
# 3. Add your brain's key
cp .env.example .env # (Windows: copy) then set your key; see "Choose a brain"
# 4. Check everything, then talk to it
uv run spark doctor # every required check should show ✔
uv run spark # say "Hey Spark, what's the capital of Japan?"
No microphone handy? uv run spark --text lets you type questions and read answers, and
uv run spark --brain mock --text runs fully offline with no key at all.
Brain
Set up
Good for
Claude (default)
Claude Code installed and signed in (or ANTHROPIC_API_KEY in .env )
The best tool use and connector work. Supports the researcher sub-agent and fast/deep model routing
Hermes Agent
Hermes Agent with its API server on, and its API_SERVER_KEY
Giving your Hermes agent a voice: it answers with its own model, tools, memory and skills
OpenAI-compatible cloud (OpenAI, Azure, OpenRouter, Groq, Together)
The provider's key in .env , and brain.backend: openai_compat
Using a model or account you already have
Local (Ollama, LM Studio, vLLM)
Run the server, pull a tool-capable model (e.g. qwen2.5:7b )
Fully local : with faster-whisper and Piper, nothing leaves your machine
The easiest way to switch is the desktop app's Brain window, which also tests the connection
and checks that the model can call tools. See docs/openai_brain.md for the
provider presets and the config.yaml settings.
With Claude, brain.mode picks the models: auto (default) sends quick requests to
brain.fast_model and bigger ones to brain.deep_model ; deep answers everything with
brain.deep_model and hands tool-heavy lookups to a researcher on a faster model (the best
answers, and the highest API cost). You can also switch by voice
("use the deep model", "back to auto").
The default config.yaml uses the bundled model assets/wakewords/hey_spark.onnx . Nothing to set
up. If it misses you or wakes on its own, tune wakeword.options.openwakeword.threshold .
Want your own wake word? tools/wakeword_trainer.ipynb trains
one for any phrase on a free Colab GPU in about an hour, with no recordings needed.
docs/wakeword.md walks through it step by step, plus the Porcupine
alternative.
Follow-ups: after a reply, Spark stays awake for 3 minutes ( attention.awake_s ), so you can
keep talking without the wake word. It answers only speech that's meant for it, then plays a
soft chime when it falls asleep.
Commands: "new topic" (fresh conversation), "go to sleep", "never mind", and "mute
yourself" (it then stops listening until you click the mic button in the window). Film-style
ones work too: "stand down", "belay that", "clean slate protocol", "full power", "go dark".
Timers and time: "set a pasta timer for 10 minutes", "what timers do I have", "what time is it".
New projects: "let's start a new project called weather-bot". After your yes , Spark creates
projects.root_dir/weather-bot with git and a first commit, a private GitHub repo and a tracking
issue through your gh login (it never pushes), and can open Claude Code in the folder.
Shortcuts: map a phrase like "good morning" to a detailed prompt in the Shortcuts window
(or shortcuts.json ). Fifteen come built in: "rise and shine" for the morning briefing,
"what's next?", "sitrep", "check the post", "run a diagnostic", "house party protocol" and more.
See docs/commands.md .
Barge-in: with app.barge_in: true , saying the wake word while Spark is talking stops it.
Turn this on only with an echo-cancelling speakerphone.
Earcons: a chime on wake, a tick while thinking, a low tone on errors.
Personality and voice: pick the butler personality and a British voice in the Personality &
Voice window (or app.persona: butler in config.yaml ). See docs/persona.md .
Connectors are MCP servers defined in mcp_servers.json (with ${VAR} placeholders only; the
values come from .env ). Add them by hand or in the desktop app's Connectors window. Each
server's tool policy lives in config.yaml under connectors: :
allowed_tools are read-only and pre-approved.
side_effecting tools (send, create, update) need a spoken "yes" every time. Spark says one
line about what it will do and waits. Silence, an unclear answer or "no" means it doesn't happen.
destructive tools are blocked unless you set allow_destructive: true for that server.
Any tool on none of these lists is denied.
Create a fine-grained token for the
repos Spark may see: Issues: Read and write , Pull requests: Read and Contents: Read .
In the Connectors window, open github (under From config files ) and click Edit
config . Replace ${SPARK_GITHUB_PAT} with your token and press Save and test . Spark
moves the token to .env .
{
"github" : {
"type" : " http " ,
"url" : " https://api.githubcopilot.com/mcp/ " ,
"headers" : { "Authorization" : " Bearer YOUR_TOKEN " }
}
}
Switch it on. Spark uses GitHub's own labels for the tools: reading runs freely, creating an
issu

[truncated]

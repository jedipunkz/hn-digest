---
source: "https://github.com/elifalanrocha/openring"
hn_url: "https://news.ycombinator.com/item?id=50372117"
title: "OpenRing – Open-source, self-hosted voice AI agent platform (Vapi alternative)"
article_title: "GitHub - elifalanrocha/openring: Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One c\n[truncated]"
image: "https://opengraph.githubassets.com/45de145ee9492f012af945742ccb52aaa2be7e67c212500587ef55956629db1c/elifalanrocha/openring"
author: "elifalanrocha"
captured_at: "2026-11-04T22:27:42Z"
capture_tool: "hn-digest"
hn_id: 50372117
score: 2
comments: 1
posted_at: "2026-11-04T21:56:55Z"
tags:
  - hacker-news
---

# OpenRing – Open-source, self-hosted voice AI agent platform (Vapi alternative)

- HN: [50372117](https://news.ycombinator.com/item?id=50372117)
- Source: [github.com](https://github.com/elifalanrocha/openring)
- Score: 2
- Comments: 1
- Posted: 2026-11-04T21:56:55Z

## Translation

Title: OpenRing – Open-source, self-hosted voice AI agent platform (Vapi alternative)
Article title: GitHub - elifalanrocha/openring: Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One c
[truncated]
Description: Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring - elifalanro
[truncated]

Article text:
GitHub - elifalanrocha/openring: Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring · GitHub
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
elifalanrocha
Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring
www.npmjs.com/package/openring Topics
Readme MIT license Code of conduct
Security policy Activity Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
109 Commits 109 Commits Folders and files
.githooks .githooks .github .github bin bin docs docs openring openring scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE PRINCIPLES.md PRINCIPLES.md README.md README.md README.pt-BR.md README.pt-BR.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md docker-compose.yml docker-compose.yml package.json package.json pyproject.toml pyproject.toml View all files Repository files navigation
Self-hosted voice AI agents with a Vapi-compatible API. No platform fee.
Quick start · How it works · Compare · API guide · Providers · Roadmap · Português
Hosted voice platforms charge a fee for every minute of every call, on top of what you already pay for
speech-to-text, the model, the voice and the phone line. OpenRing does the same job on your own computer
or server, with your own provider keys, and speaks the same JSON as Vapi.
💸 No platform fee, ever. You pay Deepgram, OpenAI, Twilio or whoever you pick, directly, at their price.
🔁 Vapi-compatible. Same assistant JSON, provider names, tool calls and webhooks: a Vapi client moves
over by changing its base URL, and Import copies an account in one click.
🖥️ Yours. Keys, calls and transcripts stay on your machine. No telemetry.
🆓 Free to try. A free local voice (Kokoro) is built in, so your first test call costs nothing.
The dashboard, with sample data.
You need Node.js 18 or later. Everything else, Python included, installs itself:
npx openring
The dashboard opens at http://localhost:8000 . From there:
Choose a password. The first visit asks for one (only from this computer).
Connect providers on the Providers page: Deepgram for speech (new accounts get US$200 of credit)
and a model, for example Groq's free tier or one-click sign-in with OpenRouter. Each key is checked with
the provider before it's saved.
Create an assistant on Assistants : a greeting, instructions, a model and a voice.
Talk to it. Press Test and speak into your microphone. No phone number needed.
Give it a number when you're ready: see Phone calls .
Settings live in ~/.openring/.env , created on the first run with a generated API key for apps that call
the API (it's under Settings ). The free Kokoro voice downloads itself on the first start (about 120 MB).
flowchart LR
caller["Caller<br/>phone or browser"] <-->|audio| line["Twilio or WebRTC"]
line <--> openring
subgraph openring["OpenRing, on your server"]
direction LR
stt["Speech-to-text<br/>Deepgram, Whisper, ..."] --> llm["Model<br/>OpenAI, Groq, Ollama, ..."] --> tts["Voice<br/>ElevenLabs, Kokoro, ..."]
end
llm <-->|"tools and webhooks, in Vapi's format"| app["Your app"]
Loading
Each call runs a Pipecat pipeline: the caller's speech becomes
text, the model answers (and can call your tools), and the voice speaks the answer back. Every turn's
latency is recorded, step by step, and shown in the dashboard.
OpenRing
Vapi
Retell
Dograh
Source code
✅ Open (MIT)
❌ Closed
❌ Closed
✅ Open (BSD-2)
Platform fee
✅ None
US$0.05/min
US$0.055/min
None self-hosted, US$0.01/min in its cloud
Run it on your server
✅ npx openring , no Docker needed (Docker works too)
❌ Their cloud
Their cloud; on-premise on Enterprise plans
✅ With Docker
Calls and transcripts
✅ On your server, no telemetry
Their cloud
Their cloud; on-premise on Enterprise plans
On your server when self-hosted; anonymous telemetry unless you turn it off
Vapi-compatible API
✅ Change the base URL, import an account
(it's Vapi)
❌
❌
Free local stack
✅ Whisper, Ollama, Kokoro
❌
❌
✅ Open models you host
Calls at once
No limit of its own, your server decides ( beta )
4, then US$10 per line a month
20, then US$8 per call a month
10 by default when self-hosted, you can raise it; 10 in its cloud
Backup model
✅ From any provider, when one fails or is slow ( beta )
Same provider only, on OpenAI and Bedrock
✅ Takes over when the main model fails or is slow
—
Phone lines
Twilio (outgoing in beta ); Telnyx, Plivo and Vonage ( beta ); incoming only on Genesys ( beta ); any carrier by SIP trunk over UDP ( beta )
✅ Twilio, Vonage, Telnyx, SIP
✅ Twilio, Telnyx, SIP
✅ Twilio, Vonage, Telnyx, Plivo, Vobiz, Cloudonix, Asterisk
Campaigns
Every line that calls out ( beta ), no retries yet
✅
✅
✅ With retries
Knowledge base (RAG)
✅ Free, on your server, matching words; or your own search ( beta )
✅ Google Gemini, or your own server
✅
✅
Several assistants in one call
✅ Squads, in Vapi's format ( beta )
✅ Squads
✅
✅
Visual flow builder
❌ Not planned for now
Canvas for squads
✅
✅
Calls over WhatsApp
❌ Not yet
—
—
✅ Incoming only
— means we didn't find it in their public docs. On every platform you also pay for the model, the voice and the
phone line, and for speech-to-text too, except that Retell's pricing page lists no separate charge for it. With
OpenRing, the server, its uptime and any compliance work are yours; hosted platforms take on the fir
[truncated]
Works today
Next
Browser test calls from the dashboard (WebRTC)
Routing speech-to-text and voices by cost
Model routes: Fastest, Cheapest, Reliable or Custom; the next model takes over when one fails or is slow ( beta )
A 100% local, free stack: Whisper, Ollama and Kokoro, one click in the editor
Provider fallback in Vapi's format, across providers ( beta )
Callers in any language: the assistant answers in theirs ( "language": "multi" )
Several assistants on one call, handing the caller to one another: squads and the handoff tool ( beta )
Handoffs chosen by your server, and to a whole squad
A knowledge base: files the assistant searches during a call, free and on your server, or your own search ( beta )
Search by meaning, built in
Incoming phone calls through Twilio, confirmed on a real call
Outbound calls through Genesys
SIP trunks: any carrier, no Twilio in between, in and out ( beta )
SIP over TCP and TLS
Outgoing calls through Twilio ( beta : tell us how it goes)
Campaigns: one assistant calls a list or CSV, a few at a time, greeting each person by name ( beta )
Calls in and out through Telnyx, Plivo and Vonage ( beta , see Telnyx , Plivo and Vonage )
Genesys Cloud calls through its Audio Connector, inside your Architect flows ( beta , see Genesys )
Vapi-style /assistant , /phone-number , /call , /campaign and /tool APIs
Cost per call in dollars, per provider, at list prices ( cost and costBreakdown )
The phone line's cost
Vapi tools: function , apiRequest (n8n, Make, Zapier), endCall , transferCall
Vapi webhooks: status-update , conversation-update , end-of-call-report , assistant-request
Calls over WhatsApp
Import assistants, tools and numbers from a Vapi account
16 models, 6 speech-to-text and 9 voice providers, with Vapi's names
Per-turn latency, time to first word, tokens and voice characters per call, with traffic lights
Credit left on OpenRouter, Deepgram, ElevenLabs, Twilio, Telnyx and Plivo, on the Providers page
Call history with transcripts, light and
[truncated]
Copy every assistant, squad, tool and Twilio number of an account in one request, or press Import on the
Assistants page. The key is used once and never stored; importing again updates what came in before.
curl -X POST http://localhost:8000/import \
-H " Authorization: Bearer $OPENRING_API_KEY " -H " Content-Type: application/json " \
-d ' {"baseUrl": "https://api.vapi.ai", "apiKey": "YOUR_VAPI_PRIVATE_KEY"} '
The answer lists each assistant with anything that keeps it from taking calls here yet (for example a
provider OpenRing doesn't support yet). Documents keep their full JSON, so nothing is lost. Then point your
client's base URL at your server, and each imported Twilio number's voice webhook at
$PUBLIC_URL/twilio/voice .
Every request takes Authorization: Bearer $OPENRING_API_KEY .
curl -X POST http://localhost:8000/assistant \
-H " Authorization: Bearer $OPENRING_API_KEY " -H " Content-Type: application/json " \
-d ' {
"name": "Front desk",
"firstMessage": "Hi! How can I help you today?",
"model": {"provider": "openai", "model": "gpt-4.1-mini",
"messages": [{"role": "system", "content": "You are a friendly receptionist."}]},
"voice": {"provider": "11labs", "voiceId": "YOUR_VOICE_ID"},
"transcriber": {"provider": "deepgram", "model": "nova-3", "language": "en"}
} '
A call ends by itself after maxDurationSeconds (10 minutes by default), or after silenceTimeoutSeconds without
activity (30 seconds by default; the clock waits while someone talks, the model thinks or a tool runs), with Vapi's
endedReason ( exceeded-max-duration , silence-timed-out ); endCallMessage is said first. When the caller goes
quiet after the assistant spoke,
messagePlan.idleMessages are said (one at random, after idleTimeoutSeconds , up to
idleMessageMaxSpokenCount times), as on Vapi:
{ "maxDurationSeconds" : 900 , "silenceTimeoutSeconds" : 20 , "endCallMessage" : " Thanks for calling, goodbye! " ,
"messagePlan" : { "idleMessages" : [ " Are you still there? " ], "idleTimeoutSeconds" : 8 }}
Names the transcriber should get right
Give it the words it will hear: names, brands, products. Use Vapi's fields, or Words to recognize in the assistant
editor:
{ "transcriber" : { "provider" : " deepgram " , "model" : " nova-3 " , "language" : " pt-BR " ,
"keyterm" : [ " Clínica Sorriso " , " Doutora Ana " ]}}
keyterm , keytermsPrompt and keywords (one word, with an optional boost: Sorriso:2 ) are all accepted, up to 100
words of up to 50 characters, and go to the field each transcriber takes: Deepgram's nova-3 takes keyterm and refuses
keywords , and older models the reverse, so OpenRing sends the right one. Phrases don't fit Deepgram's older models.
W

[truncated]

## Original Extract

Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring - elifalanro
[truncated]

GitHub - elifalanrocha/openring: Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring · GitHub
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
elifalanrocha
Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Open-source, self-hosted voice AI agents with a Vapi-compatible API. Bring your own keys or run 100% local and free (Whisper, Ollama, Kokoro). Calls on Twilio, Telnyx and Genesys, campaigns, model routes with fallback, cost per call. No per-minute platform fee. One command: npx openring
www.npmjs.com/package/openring Topics
Readme MIT license Code of conduct
Security policy Activity Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
109 Commits 109 Commits Folders and files
.githooks .githooks .github .github bin bin docs docs openring openring scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE PRINCIPLES.md PRINCIPLES.md README.md README.md README.pt-BR.md README.pt-BR.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md THIRD_PARTY.md THIRD_PARTY.md docker-compose.yml docker-compose.yml package.json package.json pyproject.toml pyproject.toml View all files Repository files navigation
Self-hosted voice AI agents with a Vapi-compatible API. No platform fee.
Quick start · How it works · Compare · API guide · Providers · Roadmap · Português
Hosted voice platforms charge a fee for every minute of every call, on top of what you already pay for
speech-to-text, the model, the voice and the phone line. OpenRing does the same job on your own computer
or server, with your own provider keys, and speaks the same JSON as Vapi.
💸 No platform fee, ever. You pay Deepgram, OpenAI, Twilio or whoever you pick, directly, at their price.
🔁 Vapi-compatible. Same assistant JSON, provider names, tool calls and webhooks: a Vapi client moves
over by changing its base URL, and Import copies an account in one click.
🖥️ Yours. Keys, calls and transcripts stay on your machine. No telemetry.
🆓 Free to try. A free local voice (Kokoro) is built in, so your first test call costs nothing.
The dashboard, with sample data.
You need Node.js 18 or later. Everything else, Python included, installs itself:
npx openring
The dashboard opens at http://localhost:8000 . From there:
Choose a password. The first visit asks for one (only from this computer).
Connect providers on the Providers page: Deepgram for speech (new accounts get US$200 of credit)
and a model, for example Groq's free tier or one-click sign-in with OpenRouter. Each key is checked with
the provider before it's saved.
Create an assistant on Assistants : a greeting, instructions, a model and a voice.
Talk to it. Press Test and speak into your microphone. No phone number needed.
Give it a number when you're ready: see Phone calls .
Settings live in ~/.openring/.env , created on the first run with a generated API key for apps that call
the API (it's under Settings ). The free Kokoro voice downloads itself on the first start (about 120 MB).
flowchart LR
caller["Caller<br/>phone or browser"] <-->|audio| line["Twilio or WebRTC"]
line <--> openring
subgraph openring["OpenRing, on your server"]
direction LR
stt["Speech-to-text<br/>Deepgram, Whisper, ..."] --> llm["Model<br/>OpenAI, Groq, Ollama, ..."] --> tts["Voice<br/>ElevenLabs, Kokoro, ..."]
end
llm <-->|"tools and webhooks, in Vapi's format"| app["Your app"]
Loading
Each call runs a Pipecat pipeline: the caller's speech becomes
text, the model answers (and can call your tools), and the voice speaks the answer back. Every turn's
latency is recorded, step by step, and shown in the dashboard.
OpenRing
Vapi
Retell
Dograh
Source code
✅ Open (MIT)
❌ Closed
❌ Closed
✅ Open (BSD-2)
Platform fee
✅ None
US$0.05/min
US$0.055/min
None self-hosted, US$0.01/min in its cloud
Run it on your server
✅ npx openring , no Docker needed (Docker works too)
❌ Their cloud
Their cloud; on-premise on Enterprise plans
✅ With Docker
Calls and transcripts
✅ On your server, no telemetry
Their cloud
Their cloud; on-premise on Enterprise plans
On your server when self-hosted; anonymous telemetry unless you turn it off
Vapi-compatible API
✅ Change the base URL, import an account
(it's Vapi)
❌
❌
Free local stack
✅ Whisper, Ollama, Kokoro
❌
❌
✅ Open models you host
Calls at once
No limit of its own, your server decides ( beta )
4, then US$10 per line a month
20, then US$8 per call a month
10 by default when self-hosted, you can raise it; 10 in its cloud
Backup model
✅ From any provider, when one fails or is slow ( beta )
Same provider only, on OpenAI and Bedrock
✅ Takes over when the main model fails or is slow
—
Phone lines
Twilio (outgoing in beta ); Telnyx, Plivo and Vonage ( beta ); incoming only on Genesys ( beta ); any carrier by SIP trunk over UDP ( beta )
✅ Twilio, Vonage, Telnyx, SIP
✅ Twilio, Telnyx, SIP
✅ Twilio, Vonage, Telnyx, Plivo, Vobiz, Cloudonix, Asterisk
Campaigns
Every line that calls out ( beta ), no retries yet
✅
✅
✅ With retries
Knowledge base (RAG)
✅ Free, on your server, matching words; or your own search ( beta )
✅ Google Gemini, or your own server
✅
✅
Several assistants in one call
✅ Squads, in Vapi's format ( beta )
✅ Squads
✅
✅
Visual flow builder
❌ Not planned for now
Canvas for squads
✅
✅
Calls over WhatsApp
❌ Not yet
—
—
✅ Incoming only
— means we didn't find it in their public docs. On every platform you also pay for the model, the voice and the
phone line, and for speech-to-text too, except that Retell's pricing page lists no separate charge for it. With
OpenRing, the server, its uptime and any compliance work are yours; hosted platforms take on the fir
[truncated]
Works today
Next
Browser test calls from the dashboard (WebRTC)
Routing speech-to-text and voices by cost
Model routes: Fastest, Cheapest, Reliable or Custom; the next model takes over when one fails or is slow ( beta )
A 100% local, free stack: Whisper, Ollama and Kokoro, one click in the editor
Provider fallback in Vapi's format, across providers ( beta )
Callers in any language: the assistant answers in theirs ( "language": "multi" )
Several assistants on one call, handing the caller to one another: squads and the handoff tool ( beta )
Handoffs chosen by your server, and to a whole squad
A knowledge base: files the assistant searches during a call, free and on your server, or your own search ( beta )
Search by meaning, built in
Incoming phone calls through Twilio, confirmed on a real call
Outbound calls through Genesys
SIP trunks: any carrier, no Twilio in between, in and out ( beta )
SIP over TCP and TLS
Outgoing calls through Twilio ( beta : tell us how it goes)
Campaigns: one assistant calls a list or CSV, a few at a time, greeting each person by name ( beta )
Calls in and out through Telnyx, Plivo and Vonage ( beta , see Telnyx , Plivo and Vonage )
Genesys Cloud calls through its Audio Connector, inside your Architect flows ( beta , see Genesys )
Vapi-style /assistant , /phone-number , /call , /campaign and /tool APIs
Cost per call in dollars, per provider, at list prices ( cost and costBreakdown )
The phone line's cost
Vapi tools: function , apiRequest (n8n, Make, Zapier), endCall , transferCall
Vapi webhooks: status-update , conversation-update , end-of-call-report , assistant-request
Calls over WhatsApp
Import assistants, tools and numbers from a Vapi account
16 models, 6 speech-to-text and 9 voice providers, with Vapi's names
Per-turn latency, time to first word, tokens and voice characters per call, with traffic lights
Credit left on OpenRouter, Deepgram, ElevenLabs, Twilio, Telnyx and Plivo, on the Providers page
Call history with transcripts, light and
[truncated]
Copy every assistant, squad, tool and Twilio number of an account in one request, or press Import on the
Assistants page. The key is used once and never stored; importing again updates what came in before.
curl -X POST http://localhost:8000/import \
-H " Authorization: Bearer $OPENRING_API_KEY " -H " Content-Type: application/json " \
-d ' {"baseUrl": "https://api.vapi.ai", "apiKey": "YOUR_VAPI_PRIVATE_KEY"} '
The answer lists each assistant with anything that keeps it from taking calls here yet (for example a
provider OpenRing doesn't support yet). Documents keep their full JSON, so nothing is lost. Then point your
client's base URL at your server, and each imported Twilio number's voice webhook at
$PUBLIC_URL/twilio/voice .
Every request takes Authorization: Bearer $OPENRING_API_KEY .
curl -X POST http://localhost:8000/assistant \
-H " Authorization: Bearer $OPENRING_API_KEY " -H " Content-Type: application/json " \
-d ' {
"name": "Front desk",
"firstMessage": "Hi! How can I help you today?",
"model": {"provider": "openai", "model": "gpt-4.1-mini",
"messages": [{"role": "system", "content": "You are a friendly receptionist."}]},
"voice": {"provider": "11labs", "voiceId": "YOUR_VOICE_ID"},
"transcriber": {"provider": "deepgram", "model": "nova-3", "language": "en"}
} '
A call ends by itself after maxDurationSeconds (10 minutes by default), or after silenceTimeoutSeconds without
activity (30 seconds by default; the clock waits while someone talks, the model thinks or a tool runs), with Vapi's
endedReason ( exceeded-max-duration , silence-timed-out ); endCallMessage is said first. When the caller goes
quiet after the assistant spoke,
messagePlan.idleMessages are said (one at random, after idleTimeoutSeconds , up to
idleMessageMaxSpokenCount times), as on Vapi:
{ "maxDurationSeconds" : 900 , "silenceTimeoutSeconds" : 20 , "endCallMessage" : " Thanks for calling, goodbye! " ,
"messagePlan" : { "idleMessages" : [ " Are you still there? " ], "idleTimeoutSeconds" : 8 }}
Names the transcriber should get right
Give it the words it will hear: names, brands, products. Use Vapi's fields, or Words to recognize in the assistant
editor:
{ "transcriber" : { "provider" : " deepgram " , "model" : " nova-3 " , "language" : " pt-BR " ,
"keyterm" : [ " Clínica Sorriso " , " Doutora Ana " ]}}
keyterm , keytermsPrompt and keywords (one word, with an optional boost: Sorriso:2 ) are all accepted, up to 100
words of up to 50 characters, and go to the field each transcriber takes: Deepgram's nova-3 takes keyterm and refuses
keywords , and older models the reverse, so OpenRing sends the right one. Phrases don't fit Deepgram's older models.
W

[truncated]

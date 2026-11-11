---
source: "https://github.com/mutonby/openshorts"
hn_url: "https://news.ycombinator.com/item?id=50455209"
title: "Openshorts: Open-source / self hosted AI clip generator (OpusClip alternative)"
article_title: "GitHub - mutonby/openshorts: Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. · GitHub"
image: "https://opengraph.githubassets.com/9c0625283f152c981d4d5119187bb6470171b703a503d0ea84ac15293ddc3596/mutonby/openshorts"
author: "devrob"
captured_at: "2026-11-11T16:55:58Z"
capture_tool: "hn-digest"
hn_id: 50455209
score: 2
comments: 0
posted_at: "2026-11-11T16:24:22Z"
tags:
  - hacker-news
---

# Openshorts: Open-source / self hosted AI clip generator (OpusClip alternative)

- HN: [50455209](https://news.ycombinator.com/item?id=50455209)
- Source: [github.com](https://github.com/mutonby/openshorts)
- Score: 2
- Comments: 0
- Posted: 2026-11-11T16:24:22Z

## Translation

Title: Openshorts: Open-source / self hosted AI clip generator (OpusClip alternative)
Article title: GitHub - mutonby/openshorts: Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. · GitHub
Description: Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. - mutonby/openshorts

Article text:
GitHub - mutonby/openshorts: Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
mutonby
openshorts Public forked from kamilstanuch / Autocrop-vertical Notifications You must be signed in to change notification settings
Star 7.7k ( 7.7k ) You must be signed in to star a repository
Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents.
Readme MIT license Activity Stars
1.7k forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
654 Commits 654 Commits Folders and files
.github/ workflows .github/ workflows .hallmark .hallmark alembic alembic assets assets bin bin cli cli cloud cloud dashboard dashboard docs docs examples/ n8n examples/ n8n fonts fonts integrations/ chatgpt-app integrations/ chatgpt-app ops ops osrender osrender remotion remotion render-service render-service screenshots screenshots skills/ openshorts skills/ openshorts tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md active_speaker.py active_speaker.py alembic.ini alembic.ini app.py app.py burned_subtitles.py burned_subtitles.py camera_inset.py camera_inset.py churchil_queen_vertical.gif churchil_queen_vertical.gif churchil_queen_vertical_short.gif churchil_queen_vertical_short.gif clip_selection.py clip_selection.py demo-openshorts.mp4 demo-openshorts.mp4 design.md design.md docker-compose.cloud.yml docker-compose.cloud.yml docker-compose.e2e.local.yml docker-compose.e2e.local.yml docker-compose.yml docker-compose.yml edit_builder.py edit_builder.py editor.py editor.py ffmpeg_utils.py ffmpeg_utils.py file_hosts.py file_hosts.py frame_sampler.py frame_sampler.py free_tools.py free_tools.py gemini_worker.py gemini_worker.py glama.json glama.json gpu_decode.py gpu_decode.py gpu_render.py gpu_render.py hook_grounding.py hook_grounding.py hook_placement.py hook_placement.py hooks.py hooks.py jsonfile.py jsonfile.py layout_picker.py layout_picker.py layout_ranges.py layout_ranges.py llm_backend.py llm_backend.py log_view.py log_view.py main.py main.py mcp_server.py mcp_server.py mcp_stdio.py mcp_stdio.py mcp_ui.py mcp_ui.py media_auth.py media_auth.py panel_layout.py panel_layout.py punch_in.py punch_in.py quality_probe.py quality_probe.py recut.py recut.py reframe_v2.py reframe_v2.py requirements-billing.txt requirements-billing.txt requirements-cpu.lock requirements-cpu.lock requirements-gpu.lock requirements-gpu.lock requirements.txt req
[truncated]
The open-source OpusClip alternative. Turn long videos into viral 9:16 Shorts, TikToks and Reels. Self-host free (MIT) or use the hosted cloud.
Open source AI video platform with 3 tools in one: Clip Generator , AI Shorts (UGC videos with AI actors) , and YouTube Studio .
Real OpenShorts output, not a mock-up. Source: Turn the Tables with Dan and Shoshana Jordan , Heritage of Faith, licensed CC BY.
Two people on camera? OpenShorts stacks them instead of shrinking the wide shot, puts the captions on the seam where they cover nobody, and switches back to a face-tracked crop when the cut goes to one person. The AI picks the layout per video; nothing to configure.
Cloud version: don't want to run it yourself? OpenShorts Cloud is the same software on a GPU. Free plan with a watermark (20 min/mo, no credit card), paid plans from $12/mo with no watermark.
Two ways to run it, same software either way:
Self-hosting is genuinely free and always will be. It costs you a machine, your own API keys and the time to keep it running. The hosted plans exist to cover that hardware and those keys, not to unlock features.
Compared to Opus Clip, Klap, Vizard (public pricing pages, checked October 2026; full breakdown at openshorts.app/alternatives ):
demo-openshorts.mp4
Video Tutorial: How it works
Click the image above to watch the full walkthrough.
Turn your long-form videos — podcasts, webinars, livestreams, vlogs, interviews — into viral-ready 9:16 shorts for TikTok, Instagram Reels, and YouTube Shorts.
2. AI Shorts (UGC Video Creator)
Generate marketing videos with AI actors for any product or business . No camera, no studio, no influencer budget. Just describe your product or paste a URL.
Two cost modes : Low Cost ( $0.65/video) and Premium ( $2/video)
Works for any business: SaaS, restaurants, e-commerce, coaching, local businesses
AI-generated actors with lip-sync, voiceover, b-roll, and TikTok-style subtitles
Choose from a shared avatar gallery or upload your own photo
Publish directly to TikTok, Instagram, and YouTube
Complete free AI YouTube toolkit: thumbnails, titles, descriptions, and direct publishing.
AI thumbnail generator with face overlay
10 viral title suggestions with refinement chat
Auto-generated descriptions with chapter timestamps
All generated videos and avatars are saved to a public gallery with SEO pages for each video.
Public gallery page with hover-to-play ( /gallery )
Individual SEO video pages with og:video meta tags ( /video/{id} )
JSON-LD structured data for search engines
Avatar gallery with prompt history
Viral Moment Detection : Google Gemini 3.1 Flash-Lite analyzes transcripts and scene boundaries to detect 3-15 high-potential moments
Runs fully local if you want : point LLM_BASE_URL at Ollama, LM Studio, vLLM or any OpenAI-compatible server and the moment picker runs on your own model, no Google key needed (see Run without a Google key )
Smart 9:16 Cropping : AI reframing per scene — TRACK mode (MediaPipe + YOLOv8 face tracking), GENERAL mode (blurred background), SPLIT mode (two speakers stacked, captions on the seam) and SCREENCAST mode (screen over presenter); the layout is picked per video by Gemini or forced from the dashboard
Auto Subtitles : faster-whisper with word-level timestamps, styled and burned into clips
AI Voice Dubbing : ElevenLabs integration for 30+ languages with voice cloning (bring your own ElevenLabs key, on the cloud too)
Hook Text Overlays : AI-generated attention-grabbing text overlays
AI Video Effects : Gemini-generated FFmpeg filters for professional effects
Analyze : Scrape website URL + web research, or generate from manual description
Script : AI writes viral scripts (hook - problem - solution - CTA format)
Actor : Generate AI actors with Flux 2 Pro or select from shared gallery
Voice : ElevenLabs TTS voiceover (English/Spanish, male/female)
Video : Talking head generation (Hailuo 2.3 Fast img2video + VEED Lipsync)
B-roll : AI-generated visuals with Ken Burns effect
Composite : FFmpeg final assembly with subtitles and hook overlays
Publish : Direct posting to TikTok, Instagram Reels, YouTube Shorts via Upload-Post
AI-powered title generation with 10 viral options
Interactive refinement chat for titles
AI thumbnail generation with custom face + background
Auto descriptions with chapter timestamps from Whisper transcript
Direct YouTube publishing via Upload-Post
One-click posting to TikTok, Instagram Reels, and YouTube Shorts simultaneously
Schedule uploads for any date and time — plan your content calendar and let OpenShorts publish automatically
Multi-platform distribution — publish to all your social networks at once from a single interface
Upload-Post integration with async uploads
S3 cloud backup (private bucket for clips, public bucket for gallery/avatars)
SEO gallery pages served by FastAPI with JSON-LD structured data
Shared avatar gallery across all users
Async job queue with configurable concurrency
Content creators — Turn long videos into shorts automatically, publish to all platforms at once
Marketing agencies — Generate UGC videos for clients at scale, no actors or studios needed
SaaS founders — Create product demos and marketing shorts from just a URL
E-commerce brands — Product videos with AI actors for TikTok Shop, Instagram, YouTube
Local businesses — Restaurants, gyms, real estate, coaching — affordable video marketing
Developers — Self-host, customize the pipeline, integrate via API
Videos generated with OpenShorts AI Shorts — no camera, no studio, no actors:
Browse all videos at openshorts.app/gallery
Feature
OpenShorts
Opus Clip
CapCut
Vizard
Klap
Descript
Price
Free self-hosted
from $12/mo hosted
$15-29/mo
$8/mo
$14.50/mo yearly
$29 monthly
from $14/mo yearly
$24-65/mo
Self-hosted
Yes
No
No
No
No
No
Open source
Yes
No
No
No
No
No
Watermark
Never self-hosted
free plan only when hosted
Free tier
Some
Free tier
No free plan
(1 trial video)
Free tier
Upload limits
None self-hosted
by plan when hosted
10-30GB
Credit-based
60min-10hr
10-100 vids/mo
60min-40hr
AI clip detection
Yes
Yes
Yes
Yes
Yes
Yes
Smart 9:16 reframing
Yes
Yes
Yes
Yes
Yes
No
Auto subtitles
Yes
Yes
Yes
Yes
Yes
Yes
Voice dubbing (30+ langs)
Yes, with your ElevenLabs key
No
Pro only
No
Pro only
Business only
AI UGC actors
Yes
No
No
No
No
No
AI video effects
Yes
No
Yes
No
No
No
Hook text overlays
Yes
No
No
No
No
No
YouTube Studio (titles, thumbnails)
Yes
No
No
No
No
No
Social auto-publishing
Yes
Pro only
TikTok only
Paid only
Paid only
No
Schedule uploads
Yes
Pro only
No
Paid only
Paid only
No
Data privacy
Your server
Their cloud
Their cloud
Their cloud
Their cloud
Their cloud
Works with a local LLM (Ollama)
Yes
No
No
No
No
No
How Much Does It Cost?
Self-hosting OpenShorts is free. You provide the machine and you only pay for the AI APIs you use, and most have generous free tiers:
Bottom line: You can clip videos for practically free with Gemini, and publish 10 videos/month to all social networks at zero cost with Upload-Post.
Don't want to run any of that? openshorts.app is the same software on our hardware: our NVIDIA GPU clips an 8-minute video in about 50 seconds instead of the 5 to 8 minutes it takes on a typical CPU, the Gemini key is included, and auto-publishing is already wired up. Free plan is 20 minutes a month with a watermark and no credit card; paid plans start at $12/mo for 100 minutes without watermark.
Google Gemini API Key ( Free — get it here ) — required for all AI features
fal.ai API Key ( Pay-per-use ) — required for AI Shorts (actor generation, video, lip-sync)
ElevenLabs API Key ( Free tier ) — required for voiceover/dubbing
Upload-Post API Key ( free tier ) — required for direct social posting
git clone https://github.com/mutonby/openshorts.git
cd openshorts
2. Configure (optional)
cp .env.example .env
# Edit .env with your AWS keys fo

[truncated]

## Original Extract

Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. - mutonby/openshorts

GitHub - mutonby/openshorts: Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
mutonby
openshorts Public forked from kamilstanuch / Autocrop-vertical Notifications You must be signed in to change notification settings
Star 7.7k ( 7.7k ) You must be signed in to star a repository
Open source AI clip generator: turns long videos into viral 9:16 shorts with AI moment detection, face tracking, subtitles and dubbing. Self-host free with Docker (MIT), or use the cloud with GPU speed from $12/mo. MCP server and API for AI agents.
Readme MIT license Activity Stars
1.7k forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
654 Commits 654 Commits Folders and files
.github/ workflows .github/ workflows .hallmark .hallmark alembic alembic assets assets bin bin cli cli cloud cloud dashboard dashboard docs docs examples/ n8n examples/ n8n fonts fonts integrations/ chatgpt-app integrations/ chatgpt-app ops ops osrender osrender remotion remotion render-service render-service screenshots screenshots skills/ openshorts skills/ openshorts tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md active_speaker.py active_speaker.py alembic.ini alembic.ini app.py app.py burned_subtitles.py burned_subtitles.py camera_inset.py camera_inset.py churchil_queen_vertical.gif churchil_queen_vertical.gif churchil_queen_vertical_short.gif churchil_queen_vertical_short.gif clip_selection.py clip_selection.py demo-openshorts.mp4 demo-openshorts.mp4 design.md design.md docker-compose.cloud.yml docker-compose.cloud.yml docker-compose.e2e.local.yml docker-compose.e2e.local.yml docker-compose.yml docker-compose.yml edit_builder.py edit_builder.py editor.py editor.py ffmpeg_utils.py ffmpeg_utils.py file_hosts.py file_hosts.py frame_sampler.py frame_sampler.py free_tools.py free_tools.py gemini_worker.py gemini_worker.py glama.json glama.json gpu_decode.py gpu_decode.py gpu_render.py gpu_render.py hook_grounding.py hook_grounding.py hook_placement.py hook_placement.py hooks.py hooks.py jsonfile.py jsonfile.py layout_picker.py layout_picker.py layout_ranges.py layout_ranges.py llm_backend.py llm_backend.py log_view.py log_view.py main.py main.py mcp_server.py mcp_server.py mcp_stdio.py mcp_stdio.py mcp_ui.py mcp_ui.py media_auth.py media_auth.py panel_layout.py panel_layout.py punch_in.py punch_in.py quality_probe.py quality_probe.py recut.py recut.py reframe_v2.py reframe_v2.py requirements-billing.txt requirements-billing.txt requirements-cpu.lock requirements-cpu.lock requirements-gpu.lock requirements-gpu.lock requirements.txt req
[truncated]
The open-source OpusClip alternative. Turn long videos into viral 9:16 Shorts, TikToks and Reels. Self-host free (MIT) or use the hosted cloud.
Open source AI video platform with 3 tools in one: Clip Generator , AI Shorts (UGC videos with AI actors) , and YouTube Studio .
Real OpenShorts output, not a mock-up. Source: Turn the Tables with Dan and Shoshana Jordan , Heritage of Faith, licensed CC BY.
Two people on camera? OpenShorts stacks them instead of shrinking the wide shot, puts the captions on the seam where they cover nobody, and switches back to a face-tracked crop when the cut goes to one person. The AI picks the layout per video; nothing to configure.
Cloud version: don't want to run it yourself? OpenShorts Cloud is the same software on a GPU. Free plan with a watermark (20 min/mo, no credit card), paid plans from $12/mo with no watermark.
Two ways to run it, same software either way:
Self-hosting is genuinely free and always will be. It costs you a machine, your own API keys and the time to keep it running. The hosted plans exist to cover that hardware and those keys, not to unlock features.
Compared to Opus Clip, Klap, Vizard (public pricing pages, checked October 2026; full breakdown at openshorts.app/alternatives ):
demo-openshorts.mp4
Video Tutorial: How it works
Click the image above to watch the full walkthrough.
Turn your long-form videos — podcasts, webinars, livestreams, vlogs, interviews — into viral-ready 9:16 shorts for TikTok, Instagram Reels, and YouTube Shorts.
2. AI Shorts (UGC Video Creator)
Generate marketing videos with AI actors for any product or business . No camera, no studio, no influencer budget. Just describe your product or paste a URL.
Two cost modes : Low Cost ( $0.65/video) and Premium ( $2/video)
Works for any business: SaaS, restaurants, e-commerce, coaching, local businesses
AI-generated actors with lip-sync, voiceover, b-roll, and TikTok-style subtitles
Choose from a shared avatar gallery or upload your own photo
Publish directly to TikTok, Instagram, and YouTube
Complete free AI YouTube toolkit: thumbnails, titles, descriptions, and direct publishing.
AI thumbnail generator with face overlay
10 viral title suggestions with refinement chat
Auto-generated descriptions with chapter timestamps
All generated videos and avatars are saved to a public gallery with SEO pages for each video.
Public gallery page with hover-to-play ( /gallery )
Individual SEO video pages with og:video meta tags ( /video/{id} )
JSON-LD structured data for search engines
Avatar gallery with prompt history
Viral Moment Detection : Google Gemini 3.1 Flash-Lite analyzes transcripts and scene boundaries to detect 3-15 high-potential moments
Runs fully local if you want : point LLM_BASE_URL at Ollama, LM Studio, vLLM or any OpenAI-compatible server and the moment picker runs on your own model, no Google key needed (see Run without a Google key )
Smart 9:16 Cropping : AI reframing per scene — TRACK mode (MediaPipe + YOLOv8 face tracking), GENERAL mode (blurred background), SPLIT mode (two speakers stacked, captions on the seam) and SCREENCAST mode (screen over presenter); the layout is picked per video by Gemini or forced from the dashboard
Auto Subtitles : faster-whisper with word-level timestamps, styled and burned into clips
AI Voice Dubbing : ElevenLabs integration for 30+ languages with voice cloning (bring your own ElevenLabs key, on the cloud too)
Hook Text Overlays : AI-generated attention-grabbing text overlays
AI Video Effects : Gemini-generated FFmpeg filters for professional effects
Analyze : Scrape website URL + web research, or generate from manual description
Script : AI writes viral scripts (hook - problem - solution - CTA format)
Actor : Generate AI actors with Flux 2 Pro or select from shared gallery
Voice : ElevenLabs TTS voiceover (English/Spanish, male/female)
Video : Talking head generation (Hailuo 2.3 Fast img2video + VEED Lipsync)
B-roll : AI-generated visuals with Ken Burns effect
Composite : FFmpeg final assembly with subtitles and hook overlays
Publish : Direct posting to TikTok, Instagram Reels, YouTube Shorts via Upload-Post
AI-powered title generation with 10 viral options
Interactive refinement chat for titles
AI thumbnail generation with custom face + background
Auto descriptions with chapter timestamps from Whisper transcript
Direct YouTube publishing via Upload-Post
One-click posting to TikTok, Instagram Reels, and YouTube Shorts simultaneously
Schedule uploads for any date and time — plan your content calendar and let OpenShorts publish automatically
Multi-platform distribution — publish to all your social networks at once from a single interface
Upload-Post integration with async uploads
S3 cloud backup (private bucket for clips, public bucket for gallery/avatars)
SEO gallery pages served by FastAPI with JSON-LD structured data
Shared avatar gallery across all users
Async job queue with configurable concurrency
Content creators — Turn long videos into shorts automatically, publish to all platforms at once
Marketing agencies — Generate UGC videos for clients at scale, no actors or studios needed
SaaS founders — Create product demos and marketing shorts from just a URL
E-commerce brands — Product videos with AI actors for TikTok Shop, Instagram, YouTube
Local businesses — Restaurants, gyms, real estate, coaching — affordable video marketing
Developers — Self-host, customize the pipeline, integrate via API
Videos generated with OpenShorts AI Shorts — no camera, no studio, no actors:
Browse all videos at openshorts.app/gallery
Feature
OpenShorts
Opus Clip
CapCut
Vizard
Klap
Descript
Price
Free self-hosted
from $12/mo hosted
$15-29/mo
$8/mo
$14.50/mo yearly
$29 monthly
from $14/mo yearly
$24-65/mo
Self-hosted
Yes
No
No
No
No
No
Open source
Yes
No
No
No
No
No
Watermark
Never self-hosted
free plan only when hosted
Free tier
Some
Free tier
No free plan
(1 trial video)
Free tier
Upload limits
None self-hosted
by plan when hosted
10-30GB
Credit-based
60min-10hr
10-100 vids/mo
60min-40hr
AI clip detection
Yes
Yes
Yes
Yes
Yes
Yes
Smart 9:16 reframing
Yes
Yes
Yes
Yes
Yes
No
Auto subtitles
Yes
Yes
Yes
Yes
Yes
Yes
Voice dubbing (30+ langs)
Yes, with your ElevenLabs key
No
Pro only
No
Pro only
Business only
AI UGC actors
Yes
No
No
No
No
No
AI video effects
Yes
No
Yes
No
No
No
Hook text overlays
Yes
No
No
No
No
No
YouTube Studio (titles, thumbnails)
Yes
No
No
No
No
No
Social auto-publishing
Yes
Pro only
TikTok only
Paid only
Paid only
No
Schedule uploads
Yes
Pro only
No
Paid only
Paid only
No
Data privacy
Your server
Their cloud
Their cloud
Their cloud
Their cloud
Their cloud
Works with a local LLM (Ollama)
Yes
No
No
No
No
No
How Much Does It Cost?
Self-hosting OpenShorts is free. You provide the machine and you only pay for the AI APIs you use, and most have generous free tiers:
Bottom line: You can clip videos for practically free with Gemini, and publish 10 videos/month to all social networks at zero cost with Upload-Post.
Don't want to run any of that? openshorts.app is the same software on our hardware: our NVIDIA GPU clips an 8-minute video in about 50 seconds instead of the 5 to 8 minutes it takes on a typical CPU, the Gemini key is included, and auto-publishing is already wired up. Free plan is 20 minutes a month with a watermark and no credit card; paid plans start at $12/mo for 100 minutes without watermark.
Google Gemini API Key ( Free — get it here ) — required for all AI features
fal.ai API Key ( Pay-per-use ) — required for AI Shorts (actor generation, video, lip-sync)
ElevenLabs API Key ( Free tier ) — required for voiceover/dubbing
Upload-Post API Key ( free tier ) — required for direct social posting
git clone https://github.com/mutonby/openshorts.git
cd openshorts
2. Configure (optional)
cp .env.example .env
# Edit .env with your AWS keys fo

[truncated]

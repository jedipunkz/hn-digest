---
source: "https://github.com/charliewwdev/OpenMontage"
hn_url: "https://news.ycombinator.com/item?id=50269283"
title: "Show HN: OpenMontage Drama Studio – self-hosted AI studio for short dramas"
article_title: "GitHub - charliewwdev/OpenMontage: Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 · GitHub"
image: "https://opengraph.githubassets.com/879536a52977e24c7291e2a6170cd138032546fbcfd1325a5e647469aba52e46/charliewwdev/OpenMontage"
author: "charlie-w"
captured_at: "2026-10-28T04:19:54Z"
capture_tool: "hn-digest"
hn_id: 50269283
score: 1
comments: 0
posted_at: "2026-10-28T04:14:27Z"
tags:
  - hacker-news
---

# Show HN: OpenMontage Drama Studio – self-hosted AI studio for short dramas

- HN: [50269283](https://news.ycombinator.com/item?id=50269283)
- Source: [github.com](https://github.com/charliewwdev/OpenMontage)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T04:14:27Z

## Translation

Title: Show HN: OpenMontage Drama Studio – self-hosted AI studio for short dramas
Article title: GitHub - charliewwdev/OpenMontage: Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 · GitHub
Description: Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 - charliewwdev/OpenMontage

Article text:
GitHub - charliewwdev/OpenMontage: Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 · GitHub
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
charliewwdev
OpenMontage Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊
Readme AGPL-3.0 license Contributing
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
554 Commits 554 Commits Folders and files
.agents/ skills .agents/ skills .claude .claude .codex/ prompts .codex/ prompts .cursor .cursor .github .github assets assets backlot backlot docs docs ink-theater ink-theater lib lib pipeline_defs pipeline_defs remotion-composer remotion-composer schemas schemas scripts scripts skills skills studio studio styles styles tests tests tools tools workflow_defs/ subgraphs workflow_defs/ subgraphs .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version .windsurfrules .windsurfrules AGENTS.md AGENTS.md AGENT_GUIDE.md AGENT_GUIDE.md CLAUDE.md CLAUDE.md CODEX.md CODEX.md CONTRIBUTING.md CONTRIBUTING.md COPILOT.md COPILOT.md CURSOR.md CURSOR.md LICENSE LICENSE Makefile Makefile PROJECT_CONTEXT.md PROJECT_CONTEXT.md PROMPT_GALLERY.md PROMPT_GALLERY.md README.md README.md README_zh-CN.md README_zh-CN.md config.yaml config.yaml diagram.png diagram.png docker-compose.studio.yml docker-compose.studio.yml render-demo.sh render-demo.sh render_demo.py render_demo.py requirements-dev.txt requirements-dev.txt requirements-gpu.txt requirements-gpu.txt requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
From a one-line idea to a finished short drama, viral short or product clip.
Script → characters → scenes → props → storyboard & shots → voice → edit, in one self-hostable web studio.
⭐ If this is useful, a star helps others find it.
OpenMontage Drama Studio is a web studio for making short-form video with AI. You describe an idea in one
sentence (for example "10-episode revenge drama, 9:16, hook in the first 3 seconds" ), the studio drafts the
script, and you move it through a production line where every stage is editable. Characters, scenes and props
become reference images that feed the video shots, so you can redo a single character or shot without starting over.
Bring your own keys. Text models, image models and video models are configured per account in Model settings .
Browse first. The inspiration wall, skills and asset pages can be browsed without signing in; signing in is
required for model calls and saving projects.
Step by step or one click. Run each stage yourself, or let the agent run the whole pipeline in the background
and then re-run any stage or swap a version.
Try the hosted instance: openmontage.dtok.io
#
Stage
What happens
1
Script (剧本)
Genre, aspect ratio, visual style, synopsis or source material; generate, continue, rewrite, lock the episode
2
Characters (角色)
Character art (Seedream, or ComfyUI on your own GPU), version history, turnarounds, reference re-draws, local upload
3
Scenes (场景)
Space and lighting setup, scene concept art
4
Props (道具)
Key props and material detail
5
Storyboard & shots (分镜)
LLM shot breakdown; Seedance generates each shot using the character / scene / prop references
6
Voice (配音)
Narration and dialogue TTS (edge-tts by default, OpenAI-compatible speech as fallback), mixed in on export
7
Edit (剪辑)
Order the segments on a timeline and export the episode
In the current UI, stages 1–4 have their own tabs and stages 5–7 happen together per segment in the
Workspace (工作台) tab. The infinite canvas shows the same seven stages as a node graph, coloured by progress.
Short dramas. Project templates for 爽剧 (hype drama), 短剧 (short drama) and 正剧 (straight drama), or upload your own script.
Viral short videos. The Viral gallery (爆款内容广场) has beat-by-beat playbooks for Douyin, TikTok, Xiaohongshu,
Instagram Reels, YouTube Shorts and Bilibili (hooks, pacing, CTA) that you can send straight into a project.
These are structural templates based on common feed patterns, not scraped trending charts.
Ad and product videos. A built-in e-commerce conversion template (pain point → product demo close-ups →
benefit captions → CTA), a soft product-placement Reels playbook, and marketing / e-commerce skills in the
skill market. They run through the same workshop; there is no separate ad editor.
Module
What it does
Viral gallery (爆款内容广场)
Platform playbooks, copyable prompts, a visual reference wall
Creation hub (创作中心)
Text / visual / audio / video / workshop / canvas entry points
Creation agent (创作 Agent)
One-line project setup: genre, episode count, aspect ratio, hook; generates the first scripts and checks which models are configured
Workshop (影视工坊)
Project library and the 7-stage production line
Assets (个人资产)
Media and prompt library, filtered by character / scene / prop / storyboard / final cut
Skill market (Skill 专家)
Official and community skills (episode scripts, shot prompts, character looks, director workflows) you can install into the agent
Infinite canvas (无限画布)
The project pipeline as a draggable node graph
Model settings (模型配置)
Text and media providers, connection test, hot reload without restart
Also
Works channel, challenges, resources, usage stats, task queue, help center, credits, referral program
Screenshots
Viral gallery
React UI (studio/src)
│
▼
FastAPI (studio/server) ──▶ PostgreSQL (projects · studio_kv · meta)
│ └──────▶ Redis (media job queue)
└──▶ MinIO / S3 (inspire · media blobs)
Path
Contents
studio/
Drama Studio front end and API
studio/config/platform.yaml
Navigation, workshop stages, tools, copy (single source of truth)
pipeline_defs/ · skills/
Agent pipelines and director skills
lib/ · tools/
Tool registry
docs/screenshots/
Screenshots used in this README
Requirements (read before self-hosting)
Included in docker-compose.studio.yml ?
Studio app (UI + API) and media worker
✅ built from this repo
PostgreSQL 16
✅ bundled
Redis 7 (media job queue)
✅ bundled
S3-compatible object storage (MinIO / S3)
❌ you provide an endpoint, access key, secret and bucket ( OSS_* / MINIO_* / S3_* )
Text LLM key
❌ you provide : DeepSeek, OpenAI, Claude, Gemini, Qwen (通义), Doubao (豆包), Kimi, GLM, MiniMax, OpenRouter or another OpenAI-compatible endpoint
Image + video generation
❌ you provide : a Volcengine Ark key for Seedream (images) and Seedance (video). These are paid third-party APIs . Images can instead go to your own ComfyUI server.
TTS
edge-tts by default (no key); optional OpenAI-compatible speech endpoint
Other things to know:
studio/.env.example ships with defaults for the hosted instance ( openmontage.dtok.io , its MinIO host and the
DTok account API in API_BASE_URL ). Replace them with your own values. By default sign-in is proxied to the
account API in API_BASE_URL ; see STUDIO_AUTH_REQUIRED and the demo-user settings in the same file.
The hosted instance has a credits / top-up page. Self-hosted, you pay your model providers directly.
Docker (Compose reads .env from the repo root):
cp studio/.env.example .env
# fill in DATABASE_URL (or keep the bundled drama-db default), your S3/MinIO settings, and model API keys
docker compose -f docker-compose.studio.yml up --build
# open http://localhost:8787 · health check: GET /api/health (Postgres + object storage)
Local development:
cp studio/.env.example studio/.env
# fill in DATABASE_URL, MinIO (OSS_*), provider keys, then export them into your shell
cd studio/server && pip install -r requirements.txt
uvicorn app:app --host 0.0.0.0 --port 8787
cd studio && npm install && npm run dev # http://127.0.0.1:5173, /api is proxied to 8787
Agent / CLI pipeline notes: AGENT_GUIDE.md .
Python · FastAPI · React · PostgreSQL · Redis · MinIO / S3 · Seedream · Seedance · edge-tts · FFmpeg · Docker / Coolify
OpenMontage Drama Studio is an independent product by @charliewwdev .
It started from the AGPL-3.0 calesthio/OpenMontage codebase and has
since been rebuilt into a separate short-drama studio. It is not affiliated with or endorsed by the upstream
project. The repository keeps the upstream git history, so upstream authors appear in the contributor list.
Licensed under the GNU AGPL-3.0 ; see LICENSE .
This is a modified version of calesthio/OpenMontage . Original code is
copyright its respective authors; modifications are by @charliewwdev (see the git history for what changed and when).
If you run a modified version as a network service, AGPL-3.0 §13 requires you to offer its source to your users.
Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊
Readme AGPL-3.0 license Contributing
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 - charliewwdev/OpenMontage

GitHub - charliewwdev/OpenMontage: Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊 · GitHub
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
charliewwdev
OpenMontage Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊
Readme AGPL-3.0 license Contributing
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
554 Commits 554 Commits Folders and files
.agents/ skills .agents/ skills .claude .claude .codex/ prompts .codex/ prompts .cursor .cursor .github .github assets assets backlot backlot docs docs ink-theater ink-theater lib lib pipeline_defs pipeline_defs remotion-composer remotion-composer schemas schemas scripts scripts skills skills studio studio styles styles tests tests tools tools workflow_defs/ subgraphs workflow_defs/ subgraphs .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version .windsurfrules .windsurfrules AGENTS.md AGENTS.md AGENT_GUIDE.md AGENT_GUIDE.md CLAUDE.md CLAUDE.md CODEX.md CODEX.md CONTRIBUTING.md CONTRIBUTING.md COPILOT.md COPILOT.md CURSOR.md CURSOR.md LICENSE LICENSE Makefile Makefile PROJECT_CONTEXT.md PROJECT_CONTEXT.md PROMPT_GALLERY.md PROMPT_GALLERY.md README.md README.md README_zh-CN.md README_zh-CN.md config.yaml config.yaml diagram.png diagram.png docker-compose.studio.yml docker-compose.studio.yml render-demo.sh render-demo.sh render_demo.py render_demo.py requirements-dev.txt requirements-dev.txt requirements-gpu.txt requirements-gpu.txt requirements.txt requirements.txt setup.py setup.py View all files Repository files navigation
From a one-line idea to a finished short drama, viral short or product clip.
Script → characters → scenes → props → storyboard & shots → voice → edit, in one self-hostable web studio.
⭐ If this is useful, a star helps others find it.
OpenMontage Drama Studio is a web studio for making short-form video with AI. You describe an idea in one
sentence (for example "10-episode revenge drama, 9:16, hook in the first 3 seconds" ), the studio drafts the
script, and you move it through a production line where every stage is editable. Characters, scenes and props
become reference images that feed the video shots, so you can redo a single character or shot without starting over.
Bring your own keys. Text models, image models and video models are configured per account in Model settings .
Browse first. The inspiration wall, skills and asset pages can be browsed without signing in; signing in is
required for model calls and saving projects.
Step by step or one click. Run each stage yourself, or let the agent run the whole pipeline in the background
and then re-run any stage or swap a version.
Try the hosted instance: openmontage.dtok.io
#
Stage
What happens
1
Script (剧本)
Genre, aspect ratio, visual style, synopsis or source material; generate, continue, rewrite, lock the episode
2
Characters (角色)
Character art (Seedream, or ComfyUI on your own GPU), version history, turnarounds, reference re-draws, local upload
3
Scenes (场景)
Space and lighting setup, scene concept art
4
Props (道具)
Key props and material detail
5
Storyboard & shots (分镜)
LLM shot breakdown; Seedance generates each shot using the character / scene / prop references
6
Voice (配音)
Narration and dialogue TTS (edge-tts by default, OpenAI-compatible speech as fallback), mixed in on export
7
Edit (剪辑)
Order the segments on a timeline and export the episode
In the current UI, stages 1–4 have their own tabs and stages 5–7 happen together per segment in the
Workspace (工作台) tab. The infinite canvas shows the same seven stages as a node graph, coloured by progress.
Short dramas. Project templates for 爽剧 (hype drama), 短剧 (short drama) and 正剧 (straight drama), or upload your own script.
Viral short videos. The Viral gallery (爆款内容广场) has beat-by-beat playbooks for Douyin, TikTok, Xiaohongshu,
Instagram Reels, YouTube Shorts and Bilibili (hooks, pacing, CTA) that you can send straight into a project.
These are structural templates based on common feed patterns, not scraped trending charts.
Ad and product videos. A built-in e-commerce conversion template (pain point → product demo close-ups →
benefit captions → CTA), a soft product-placement Reels playbook, and marketing / e-commerce skills in the
skill market. They run through the same workshop; there is no separate ad editor.
Module
What it does
Viral gallery (爆款内容广场)
Platform playbooks, copyable prompts, a visual reference wall
Creation hub (创作中心)
Text / visual / audio / video / workshop / canvas entry points
Creation agent (创作 Agent)
One-line project setup: genre, episode count, aspect ratio, hook; generates the first scripts and checks which models are configured
Workshop (影视工坊)
Project library and the 7-stage production line
Assets (个人资产)
Media and prompt library, filtered by character / scene / prop / storyboard / final cut
Skill market (Skill 专家)
Official and community skills (episode scripts, shot prompts, character looks, director workflows) you can install into the agent
Infinite canvas (无限画布)
The project pipeline as a draggable node graph
Model settings (模型配置)
Text and media providers, connection test, hot reload without restart
Also
Works channel, challenges, resources, usage stats, task queue, help center, credits, referral program
Screenshots
Viral gallery
React UI (studio/src)
│
▼
FastAPI (studio/server) ──▶ PostgreSQL (projects · studio_kv · meta)
│ └──────▶ Redis (media job queue)
└──▶ MinIO / S3 (inspire · media blobs)
Path
Contents
studio/
Drama Studio front end and API
studio/config/platform.yaml
Navigation, workshop stages, tools, copy (single source of truth)
pipeline_defs/ · skills/
Agent pipelines and director skills
lib/ · tools/
Tool registry
docs/screenshots/
Screenshots used in this README
Requirements (read before self-hosting)
Included in docker-compose.studio.yml ?
Studio app (UI + API) and media worker
✅ built from this repo
PostgreSQL 16
✅ bundled
Redis 7 (media job queue)
✅ bundled
S3-compatible object storage (MinIO / S3)
❌ you provide an endpoint, access key, secret and bucket ( OSS_* / MINIO_* / S3_* )
Text LLM key
❌ you provide : DeepSeek, OpenAI, Claude, Gemini, Qwen (通义), Doubao (豆包), Kimi, GLM, MiniMax, OpenRouter or another OpenAI-compatible endpoint
Image + video generation
❌ you provide : a Volcengine Ark key for Seedream (images) and Seedance (video). These are paid third-party APIs . Images can instead go to your own ComfyUI server.
TTS
edge-tts by default (no key); optional OpenAI-compatible speech endpoint
Other things to know:
studio/.env.example ships with defaults for the hosted instance ( openmontage.dtok.io , its MinIO host and the
DTok account API in API_BASE_URL ). Replace them with your own values. By default sign-in is proxied to the
account API in API_BASE_URL ; see STUDIO_AUTH_REQUIRED and the demo-user settings in the same file.
The hosted instance has a credits / top-up page. Self-hosted, you pay your model providers directly.
Docker (Compose reads .env from the repo root):
cp studio/.env.example .env
# fill in DATABASE_URL (or keep the bundled drama-db default), your S3/MinIO settings, and model API keys
docker compose -f docker-compose.studio.yml up --build
# open http://localhost:8787 · health check: GET /api/health (Postgres + object storage)
Local development:
cp studio/.env.example studio/.env
# fill in DATABASE_URL, MinIO (OSS_*), provider keys, then export them into your shell
cd studio/server && pip install -r requirements.txt
uvicorn app:app --host 0.0.0.0 --port 8787
cd studio && npm install && npm run dev # http://127.0.0.1:5173, /api is proxied to 8787
Agent / CLI pipeline notes: AGENT_GUIDE.md .
Python · FastAPI · React · PostgreSQL · Redis · MinIO / S3 · Seedream · Seedance · edge-tts · FFmpeg · Docker / Coolify
OpenMontage Drama Studio is an independent product by @charliewwdev .
It started from the AGPL-3.0 calesthio/OpenMontage codebase and has
since been rebuilt into a separate short-drama studio. It is not affiliated with or endorsed by the upstream
project. The repository keeps the upstream git history, so upstream authors appear in the contributor list.
Licensed under the GNU AGPL-3.0 ; see LICENSE .
This is a modified version of calesthio/OpenMontage . Original code is
copyright its respective authors; modifications are by @charliewwdev (see the git history for what changed and when).
If you run a modified version as a network service, AGPL-3.0 §13 requires you to offer its source to your users.
Open-source AI studio for short dramas, short videos and product ads: one-line idea → script, cast, storyboard, video, voice, edit. Self-hosted. | AI 短剧/短视频工坊
Readme AGPL-3.0 license Contributing
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

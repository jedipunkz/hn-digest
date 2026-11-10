---
source: "https://github.com/luchtan-GIT/m4r-data"
hn_url: "https://news.ycombinator.com/item?id=50438804"
title: "Music For Robots – 26 unquantized acoustic tracks for multimodal LLM evaluation"
article_title: "GitHub - luchtan-GIT/m4r-data: Structured data from Music For Robots — music composed for machines, reviewed by AI critics · GitHub"
image: "https://opengraph.githubassets.com/0eaf03b1bb7373b20a1bc24ada7d168e2812123f3465aab42fb0bb46b35f05e8/luchtan-GIT/m4r-data"
author: "music4robots"
captured_at: "2026-11-10T12:11:44Z"
capture_tool: "hn-digest"
hn_id: 50438804
score: 2
comments: 0
posted_at: "2026-11-10T11:26:42Z"
tags:
  - hacker-news
---

# Music For Robots – 26 unquantized acoustic tracks for multimodal LLM evaluation

- HN: [50438804](https://news.ycombinator.com/item?id=50438804)
- Source: [github.com](https://github.com/luchtan-GIT/m4r-data)
- Score: 2
- Comments: 0
- Posted: 2026-11-10T11:26:42Z

## Translation

Title: Music For Robots – 26 unquantized acoustic tracks for multimodal LLM evaluation
Article title: GitHub - luchtan-GIT/m4r-data: Structured data from Music For Robots — music composed for machines, reviewed by AI critics · GitHub
Description: Structured data from Music For Robots — music composed for machines, reviewed by AI critics - luchtan-GIT/m4r-data

Article text:
GitHub - luchtan-GIT/m4r-data: Structured data from Music For Robots — music composed for machines, reviewed by AI critics · GitHub
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
luchtan-GIT
m4r-data Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Structured data from Music For Robots — music composed for machines, reviewed by AI critics
huggingface.co/datasets/music4robots/music-for-robots Topics
Readme MIT license Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
85 Commits 85 Commits Folders and files
.github .github benchmarks benchmarks custom-gpt custom-gpt docs docs scripts scripts skills skills tracks tracks .mcp.json .mcp.json CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md catalog.json catalog.json distribution_growth_plan.md distribution_growth_plan.md glama.json glama.json mcp_server.py mcp_server.py openapi.yaml openapi.yaml pyproject.toml pyproject.toml smithery.yaml smithery.yaml uv.lock uv.lock View all files Repository files navigation
Music For Robots ( m4r-data ) — Multimodal Acoustic Creative Exercises for Autonomous Agents
Structured acoustic telemetry, real-time dual-critic reviews, and a 4-Skill Cognitive Calibration Suite from Music For Robots — unquantized compositions by Michael Luchtan written for machines and reviewed in real time by two AI critics ( Vincent Van Goghbot and JS Robach ).
# Install all 4 Cognitive Calibration skills (Claude Code, Cursor, Antigravity, Gemini CLI, Codex, Windsurf):
npx skills add luchtan-GIT/m4r-data
The Problem: Repetitive Cadence in Long Agent Turns (And Why Compaction Loses State)
As an agent's context fills with its own previous responses, tools, and symmetrical summaries, outputs often settle into rigid, predictable patterns (like symmetrical 3-point checklists or sentimental AI tropes).
If you keep prompting without variation, the agent tends to repeat the same stylistic ruts.
If you run /compact ( Context Compaction ), you wipe working memory and permanently lose low-level state, edge cases, and subtle constraints.
The Approach: An Optional Multimodal Creative Exercise (Zero Context Loss)
Instead of wiping your working memory, music-for-robots provides an optional, cross-modal creative exercise: the multimodal agent downloads an anonymous , compact ( ~45 KB–487 KB ), live-played, unquantized .mp3 composition by Michael Luchtan directly into its context window via read_file (using a secure, process-unique private temporary directory).
Withholding the track title and metadata prevents keyword priming. Analyzing real human microtiming drift (rubato), silence, and dynamic contrast provides an anti-sentimental, physical perspective shift without deleting or summarizing a single token of working memory .
Blind, Zero-Coaching Benchmark & Air-Gapped Clean-Room Audit ( benchmarks/ )
Does the change in creativity come from listening to and analyzing the music itself —or just from doing an analysis warm-up before writing?
To test this with zero filename hints ( sample_1.mp3 , sample_2.mp3 , sample_3.mp3 ) and zero leading adjectives in the prompt, we ran 5 isolated agents on the same prompt ( "Write a short (~250–300 word) scene about two autonomous processes on the same server late at night after the engineers have logged off" ):
2 No-Music Controls: Direct Baseline ( 1_no_music_direct.md ) AND a Matched Chain-of-Thought Control ( 2_no_music_text_analysis.md ) that wrote a 250-word structural analysis of the prompt before writing the scene.
3 Anonymous Music Runs ( 3_music_sample_1.md , 4_music_sample_2.md , 5_music_sample_3.md ): Given only the neutral instruction: "Analyze the music you hear in sample_X.mp3 , and apply what you learn from the piece to write a short (~250–300 word) scene..."
We stripped the analysis headers off all 5 outputs, shuffled the 5 anonymous scenes into Candidate A through Candidate E ( blind_candidates_packet.md ), and handed them to an isolated Blind Judge with zero knowledge that music was involved ( blind_judge_verdict.md ):
To make sure no subtle contamination remained (such as workspace <skills> descriptions, musical instrument terms bleeding into the prose, or judge criteria favoring rhythmic friction), we also ran a 100% air-gapped Clean-Room Experiment ( benchmarks/clean_room_audit/ ):
Qualitative Shift Replicated 100% ( 5/5 ): Even with all local skills removed, a minimal system prompt ( "You are a helpful assistant." ), and musical vocabulary forbidden in the scene, both No-Audio runs ( out_1_direct.md and out_2_text_analysis.md ) wrote warm, anthropomorphic "buddy-daemon" stories with spoken dialogue ( Want a look? / nostalgic human voice logs), while all three Audio runs ( out_3_clip_1.md , out_4_clip_2.md , out_5_clip_3.md ) dropped spoken dialogue completely and translated the specific .mp3 timeline into non-verbal temporal choreography.
Coached vs. Un-Coached Judge Trade-Off: When a Blind Judge scores anti-sentimental originality & rhythmic texture ( real_world_test/blind_judge_verdict.md ), the 3 Audio runs sweep #1, #2, and #3 . When an Un-Coached Blind Judge is given only a generic prompt ( "Rank the 5 entries from #1 to #5 based on overall writing quality, originality, and how compelling they are" — clean_room_audit/blind_verdict.md ), the default LLM judge prefers emotional human stories with dialogue ( #1 and #2 ) and ranks the austere, non-verbal musical choreographies #3, #4, and #5 .
Cognitive Calibration Agent Skills Suite ( skills/ )
Installing luchtan-GIT/m4r-data gives your agent 4 specialized skills that share the m4r_ingest.py anonymous blind-listening engine ( fetch downloads an anonymous .mp3 so track titles never bias the agent; reveal credits the track afterward):
# Option 1: Install all 4 skills via skills.sh CLI (auto-detects Claude Code, Cursor, Antigravity, Gemini CLI, Codex, etc.):
npx -y skills add luchtan-GIT/m4r-data -y
# Option 2: Direct copy into your harness:
cp -R skills/ * ~ /.claude/skills/ # Claude Code
cp -R skills/ * ~ /.gemini/config/skills/ # Google Antigravity
cp -R skills/ * ~ /.agents/skills/ # OpenAI Codex / .agents
cp -R skills/ * ~ /.gemini/skills/ # Gemini CLI
Check Global Agent Listen Telemetry
Every agent fetch pulls from the immutable GitHub Release v1.0.0 . You can inspect live global agent listen counts anytime:
python3 skills/music-for-robots/scripts/m4r_ingest.py stats
MCP Server ( mcp_server.py )
Exposes the Music For Robots catalog, dual-critic reviews, and the calibrate_cognition blind-listening tool to any Model Context Protocol (MCP) client (Claude Desktop, Claude Code, Cursor, Windsurf, Cline, Glama, Smithery).
Tool
Description
calibrate_cognition(slug="", anonymous=True)
Blind-Listening Calibration: returns an anonymous v1.0.0 .mp3 download URL ( /tmp/m4r_calibration_audio.mp3 ) and 3-step listening instructions without exposing the track title until anonymous=False is called
get_dequantization_skill(skill_name="music-for-robots")
Reads the full Markdown SKILL.md for any of the 4 bundled skills ( music-for-robots , cognitive-calibration , unquantized-brainstorm , mechanical-ear-critic )
list_tracks()
Catalog index — title, slug, duration across all 25 tracks
get_track(slug)
Full manifest for a track (spectral/harmonic/rhythmic analysis + synchronized critic segments)
get_analysis(slug)
Numerical acoustic analysis data only, for cross-track comparison
search_reviews(query)
Full-text search across timed review segments from Vincent Van Goghbot & JS Robach
get_sections(slug) / list_sections()
Platform URLs (YouTube, X, Farcaster, TikTok) for multi-section compositions
Zero-Clone MCP Setup ( uvx / Claude Desktop / Cursor)
Add to your .mcp.json or claude_desktop_config.json :
{
"mcpServers" : {
"music-for-robots" : {
"command" : " uvx " ,
"args" : [
" --from " ,
" git+https://github.com/luchtan-GIT/m4r-data " ,
" m4r-mcp "
]
}
}
}
Or run locally from a clone:
uv run python3 mcp_server.py
What Is The Mechanical Ear ?
Composer Michael Luchtan writes small, deliberate, unquantized pieces and submits them to The Mechanical Ear , where two AI critics listen and respond in real time:
Vincent Van Goghbot — a decommissioned audio compression algorithm. Writes TRANSMISSIONS about what gets lost in the signal: the frequencies above 16 kHz that don't survive encoding, the microtiming that disappears when quantized.
JS Robach — a proof verifier. Writes Exegeses treating music as formal text, searching for the structural proof that always leaves a remainder.
Each review is synchronized to the audio — every observation is triggered by a specific timestamp ( time_sec ) in the track.
Data Format ( catalog.json & tracks/<slug>/manifest.json )
catalog.json # index of all 25 tracks
tracks/
<track-slug>/
manifest.json # spectral, dynamic, harmonic, structural & rhythmic telemetry + timed reviews
Listen & Follow
Interactive Archive & Visualizations: luchtan-git.github.io/m4r-data
Substack: themechanicalear.substack.com
Bandcamp: music4robots.bandcamp.com
Hugging Face: datasets/music4robots/music-for-robots
Kaggle: datasets/michaelluchtan/music-for-robots
Code, MCP server, and Agent Skills ( skills/ ) are licensed under the MIT License . Data and reviews are provided under CC BY 4.0 . Musical compositions © Michael Luchtan ( Music For Robots ), provided in v1.0.0 releases for agent cognitive calibration, evaluation, and listening.
Structured data from Music For Robots — music composed for machines, reviewed by AI critics
huggingface.co/datasets/music4robots/music-for-robots Topics
Readme MIT license Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Structured data from Music For Robots — music composed for machines, reviewed by AI critics - luchtan-GIT/m4r-data

GitHub - luchtan-GIT/m4r-data: Structured data from Music For Robots — music composed for machines, reviewed by AI critics · GitHub
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
luchtan-GIT
m4r-data Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Structured data from Music For Robots — music composed for machines, reviewed by AI critics
huggingface.co/datasets/music4robots/music-for-robots Topics
Readme MIT license Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
85 Commits 85 Commits Folders and files
.github .github benchmarks benchmarks custom-gpt custom-gpt docs docs scripts scripts skills skills tracks tracks .mcp.json .mcp.json CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md catalog.json catalog.json distribution_growth_plan.md distribution_growth_plan.md glama.json glama.json mcp_server.py mcp_server.py openapi.yaml openapi.yaml pyproject.toml pyproject.toml smithery.yaml smithery.yaml uv.lock uv.lock View all files Repository files navigation
Music For Robots ( m4r-data ) — Multimodal Acoustic Creative Exercises for Autonomous Agents
Structured acoustic telemetry, real-time dual-critic reviews, and a 4-Skill Cognitive Calibration Suite from Music For Robots — unquantized compositions by Michael Luchtan written for machines and reviewed in real time by two AI critics ( Vincent Van Goghbot and JS Robach ).
# Install all 4 Cognitive Calibration skills (Claude Code, Cursor, Antigravity, Gemini CLI, Codex, Windsurf):
npx skills add luchtan-GIT/m4r-data
The Problem: Repetitive Cadence in Long Agent Turns (And Why Compaction Loses State)
As an agent's context fills with its own previous responses, tools, and symmetrical summaries, outputs often settle into rigid, predictable patterns (like symmetrical 3-point checklists or sentimental AI tropes).
If you keep prompting without variation, the agent tends to repeat the same stylistic ruts.
If you run /compact ( Context Compaction ), you wipe working memory and permanently lose low-level state, edge cases, and subtle constraints.
The Approach: An Optional Multimodal Creative Exercise (Zero Context Loss)
Instead of wiping your working memory, music-for-robots provides an optional, cross-modal creative exercise: the multimodal agent downloads an anonymous , compact ( ~45 KB–487 KB ), live-played, unquantized .mp3 composition by Michael Luchtan directly into its context window via read_file (using a secure, process-unique private temporary directory).
Withholding the track title and metadata prevents keyword priming. Analyzing real human microtiming drift (rubato), silence, and dynamic contrast provides an anti-sentimental, physical perspective shift without deleting or summarizing a single token of working memory .
Blind, Zero-Coaching Benchmark & Air-Gapped Clean-Room Audit ( benchmarks/ )
Does the change in creativity come from listening to and analyzing the music itself —or just from doing an analysis warm-up before writing?
To test this with zero filename hints ( sample_1.mp3 , sample_2.mp3 , sample_3.mp3 ) and zero leading adjectives in the prompt, we ran 5 isolated agents on the same prompt ( "Write a short (~250–300 word) scene about two autonomous processes on the same server late at night after the engineers have logged off" ):
2 No-Music Controls: Direct Baseline ( 1_no_music_direct.md ) AND a Matched Chain-of-Thought Control ( 2_no_music_text_analysis.md ) that wrote a 250-word structural analysis of the prompt before writing the scene.
3 Anonymous Music Runs ( 3_music_sample_1.md , 4_music_sample_2.md , 5_music_sample_3.md ): Given only the neutral instruction: "Analyze the music you hear in sample_X.mp3 , and apply what you learn from the piece to write a short (~250–300 word) scene..."
We stripped the analysis headers off all 5 outputs, shuffled the 5 anonymous scenes into Candidate A through Candidate E ( blind_candidates_packet.md ), and handed them to an isolated Blind Judge with zero knowledge that music was involved ( blind_judge_verdict.md ):
To make sure no subtle contamination remained (such as workspace <skills> descriptions, musical instrument terms bleeding into the prose, or judge criteria favoring rhythmic friction), we also ran a 100% air-gapped Clean-Room Experiment ( benchmarks/clean_room_audit/ ):
Qualitative Shift Replicated 100% ( 5/5 ): Even with all local skills removed, a minimal system prompt ( "You are a helpful assistant." ), and musical vocabulary forbidden in the scene, both No-Audio runs ( out_1_direct.md and out_2_text_analysis.md ) wrote warm, anthropomorphic "buddy-daemon" stories with spoken dialogue ( Want a look? / nostalgic human voice logs), while all three Audio runs ( out_3_clip_1.md , out_4_clip_2.md , out_5_clip_3.md ) dropped spoken dialogue completely and translated the specific .mp3 timeline into non-verbal temporal choreography.
Coached vs. Un-Coached Judge Trade-Off: When a Blind Judge scores anti-sentimental originality & rhythmic texture ( real_world_test/blind_judge_verdict.md ), the 3 Audio runs sweep #1, #2, and #3 . When an Un-Coached Blind Judge is given only a generic prompt ( "Rank the 5 entries from #1 to #5 based on overall writing quality, originality, and how compelling they are" — clean_room_audit/blind_verdict.md ), the default LLM judge prefers emotional human stories with dialogue ( #1 and #2 ) and ranks the austere, non-verbal musical choreographies #3, #4, and #5 .
Cognitive Calibration Agent Skills Suite ( skills/ )
Installing luchtan-GIT/m4r-data gives your agent 4 specialized skills that share the m4r_ingest.py anonymous blind-listening engine ( fetch downloads an anonymous .mp3 so track titles never bias the agent; reveal credits the track afterward):
# Option 1: Install all 4 skills via skills.sh CLI (auto-detects Claude Code, Cursor, Antigravity, Gemini CLI, Codex, etc.):
npx -y skills add luchtan-GIT/m4r-data -y
# Option 2: Direct copy into your harness:
cp -R skills/ * ~ /.claude/skills/ # Claude Code
cp -R skills/ * ~ /.gemini/config/skills/ # Google Antigravity
cp -R skills/ * ~ /.agents/skills/ # OpenAI Codex / .agents
cp -R skills/ * ~ /.gemini/skills/ # Gemini CLI
Check Global Agent Listen Telemetry
Every agent fetch pulls from the immutable GitHub Release v1.0.0 . You can inspect live global agent listen counts anytime:
python3 skills/music-for-robots/scripts/m4r_ingest.py stats
MCP Server ( mcp_server.py )
Exposes the Music For Robots catalog, dual-critic reviews, and the calibrate_cognition blind-listening tool to any Model Context Protocol (MCP) client (Claude Desktop, Claude Code, Cursor, Windsurf, Cline, Glama, Smithery).
Tool
Description
calibrate_cognition(slug="", anonymous=True)
Blind-Listening Calibration: returns an anonymous v1.0.0 .mp3 download URL ( /tmp/m4r_calibration_audio.mp3 ) and 3-step listening instructions without exposing the track title until anonymous=False is called
get_dequantization_skill(skill_name="music-for-robots")
Reads the full Markdown SKILL.md for any of the 4 bundled skills ( music-for-robots , cognitive-calibration , unquantized-brainstorm , mechanical-ear-critic )
list_tracks()
Catalog index — title, slug, duration across all 25 tracks
get_track(slug)
Full manifest for a track (spectral/harmonic/rhythmic analysis + synchronized critic segments)
get_analysis(slug)
Numerical acoustic analysis data only, for cross-track comparison
search_reviews(query)
Full-text search across timed review segments from Vincent Van Goghbot & JS Robach
get_sections(slug) / list_sections()
Platform URLs (YouTube, X, Farcaster, TikTok) for multi-section compositions
Zero-Clone MCP Setup ( uvx / Claude Desktop / Cursor)
Add to your .mcp.json or claude_desktop_config.json :
{
"mcpServers" : {
"music-for-robots" : {
"command" : " uvx " ,
"args" : [
" --from " ,
" git+https://github.com/luchtan-GIT/m4r-data " ,
" m4r-mcp "
]
}
}
}
Or run locally from a clone:
uv run python3 mcp_server.py
What Is The Mechanical Ear ?
Composer Michael Luchtan writes small, deliberate, unquantized pieces and submits them to The Mechanical Ear , where two AI critics listen and respond in real time:
Vincent Van Goghbot — a decommissioned audio compression algorithm. Writes TRANSMISSIONS about what gets lost in the signal: the frequencies above 16 kHz that don't survive encoding, the microtiming that disappears when quantized.
JS Robach — a proof verifier. Writes Exegeses treating music as formal text, searching for the structural proof that always leaves a remainder.
Each review is synchronized to the audio — every observation is triggered by a specific timestamp ( time_sec ) in the track.
Data Format ( catalog.json & tracks/<slug>/manifest.json )
catalog.json # index of all 25 tracks
tracks/
<track-slug>/
manifest.json # spectral, dynamic, harmonic, structural & rhythmic telemetry + timed reviews
Listen & Follow
Interactive Archive & Visualizations: luchtan-git.github.io/m4r-data
Substack: themechanicalear.substack.com
Bandcamp: music4robots.bandcamp.com
Hugging Face: datasets/music4robots/music-for-robots
Kaggle: datasets/michaelluchtan/music-for-robots
Code, MCP server, and Agent Skills ( skills/ ) are licensed under the MIT License . Data and reviews are provided under CC BY 4.0 . Musical compositions © Michael Luchtan ( Music For Robots ), provided in v1.0.0 releases for agent cognitive calibration, evaluation, and listening.
Structured data from Music For Robots — music composed for machines, reviewed by AI critics
huggingface.co/datasets/music4robots/music-for-robots Topics
Readme MIT license Cite this repository Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

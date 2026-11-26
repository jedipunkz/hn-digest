---
source: "https://github.com/andrewfallan-alt/ai-music-video-continuity"
hn_url: "https://news.ycombinator.com/item?id=50655725"
title: "Python scripts for 59-shot AI video face-lock, iris-lock, and lip-sync"
article_title: "GitHub - andrewfallan-alt/ai-music-video-continuity: Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. · GitHub"
image: "https://opengraph.githubassets.com/ecf552201c6b40c780956b4969166817e5ccbd9715d4a0e6d7092f010d25412d/andrewfallan-alt/ai-music-video-continuity"
author: "afallan"
captured_at: "2026-11-26T14:45:17Z"
capture_tool: "hn-digest"
hn_id: 50655725
score: 1
comments: 0
posted_at: "2026-11-26T14:33:28Z"
tags:
  - hacker-news
---

# Python scripts for 59-shot AI video face-lock, iris-lock, and lip-sync

- HN: [50655725](https://news.ycombinator.com/item?id=50655725)
- Source: [github.com](https://github.com/andrewfallan-alt/ai-music-video-continuity)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T14:33:28Z

## Translation

Title: Python scripts for 59-shot AI video face-lock, iris-lock, and lip-sync
Article title: GitHub - andrewfallan-alt/ai-music-video-continuity: Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. · GitHub
Description: Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. - andrewfallan-alt/ai-music-video-continuity

Article text:
GitHub - andrewfallan-alt/ai-music-video-continuity: Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. · GitHub
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
andrewfallan-alt
ai-music-video-continuity Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
LICENSE LICENSE README.md README.md align_lipsync.py align_lipsync.py apply_facelock.py apply_facelock.py fix_iris_continuity.py fix_iris_continuity.py requirements.txt requirements.txt View all files Repository files navigation
Killing AI Slop Across 59 Shots: Multi-Frame Face-Lock, Savitzky-Golay Iris-Lock, and Zero-Freeze Lip-Sync in Python
Demo outputs: 3:42 End-to-End Making-Of Breakdown · 14s Raw vs. Locked Split-Screen · Full 5:53 Music Video (59 shots, 1080p24)
Most AI music videos look unmistakably like "AI slop" after about four seconds—not because the base image or video models lack resolution, but because temporal continuity breaks down across cuts :
Identity drift: Your protagonist's bone structure subtly shifts every time you switch from a wide establishing shot to a singing close-up.
Iris collapse & drift: Lip-sync models routinely either drift stylized eyes back to generic blue-grey, or blow bright eyes out into featureless white flashlights.
Frozen lip-sync tails: Audio-driven lip-sync outputs rarely align cleanly to exact song beat boundaries, leaving dead frozen frames at clip leads/tails or 120 ms mouth-to-vocal skew.
While producing Forever, Darling (Auto-Renew) —a 5:53, 59-shot music video built with Gemini 4 Neon (concept, lyrics, storyboarding, post-production code), Suno (136 BPM F# minor dark-pop × symphonic rock-opera), and Google Flow (Veo + Omni 1.2 Flash) —we solved all three problems deterministically in Python ( OpenCV , InsightFace , SciPy , FFmpeg ) using zero additional video-generation credits .
1. Multi-Frame Golden Embedding Face-Lock ( apply_facelock_v3.py )
Single-image face swapping fails across drastic lighting and expression shifts (e.g., soft morning clinic light vs. golden-hour operatic belting). Instead of anchoring to a single portrait, we extract three approved hero frames across distinct lighting conditions ( S47 close-up, S29 chorus singing, S13 sunlit mid-shot) and compute a weighted normalized ArcFace embedding via InsightFace's buffalo_l :
f47 = app . get ( grab_frame ( "S47_lipsync.mp4" , 0.55 ))[ 0 ]
f29 = app . get ( grab_frame ( "S29_lipsync.mp4" , 0.25 ))[ 0 ]
f13 = app . get ( grab_frame ( "S13_raw.mp4" , 0.60 ))[ 0 ]
emb = 0.45 * f47 . normed_embedding + 0.35 * f29 . normed_embedding + 0.20 * f13 . normed_embedding
emb = emb / np . linalg . norm ( emb )
golden = f29
golden . embedding = emb * np . linalg . norm ( f29 . embedding )
Running inswapper_128.onnx across ProcessPoolExecutor CPU workers with that blended golden identity—combined with 106-point landmark masking—locked facial geometry across all 30+ character shots without burning a single extra video-gen credit.
2. Dual-Mode Iris Physics + Temporal Smoothing ( fix_xena_eyes.py )
Xena's signature visual cue is bioluminescent xenon-lavender irises with crisp dark pupils. Across 59 raw shots, video and lip-sync generators failed in two opposite ways , so we built two deterministic per-shot modes:
tint mode (color drift): Shifts drifted blue-grey irises ( S04 , S29 , S55 ) cleanly into xenon-lavender in HSV/LAB space while strictly masking out pupils, specular catchlights, lashes, sclera, and blink frames.
tame mode (glow blowout): Reconstructs crisp dark pupils and defined iris rings on shots where bright lighting washed the eyes into blank white circles ( S18 , S45 , S47 , S51 ), again respecting natural blinks.
To stop frame-to-frame jitter from per-frame eye detection, pupil/iris coordinates are smoothed temporally across each clip using scipy.signal.savgol_filter before rendering back via piped ffmpeg ( -crf 15 ).
(14-second raw-vs-locked visual comparison: Before vs. After Python Split-Screen )
3. Zero-Frozen-Frame Word-Level Lip-Sync Retiming ( align_xena_lipsync.py )
Audio-driven lip-sync models frequently introduce slight drift or padded lead/tail frames. Rather than manually nudging 11 singing takes in an NLE:
We transcribed the isolated vocal stem ( Demucs → word-level onsets) and wrote a canonical xena_timeline.json across all 59 shots with six surgical editorial boundary overrides ( S17/S18 , S29/S30 , S32/S33 , S45/S46 , S51/S52 , S53/S54 ).
Each Omni 1.2 Flash clip is retimed directly to its exact frame window ( round(start_s * 24) to round(end_s * 24) ) with zero frozen lead/tail frames.
All HUD graphics ( BIOLOGICAL AGE: 48.2 → 25.0 (LOCKED) → AUTO-RENEW: ∞ ) and word-synced karaoke captions are kept 100% out of video prompts (eliminating garbled AI text) and burned cleanly in post-production via ASS subtitles + a 10-panel ffmpeg split-screen montage grid ( build_xena_montage.py ).
Never put readable text or HUDs inside video prompts. Keeping typography 100% in post-production ASS overlays eliminated 90% of wasted video retries.
Post-production beats re-rolling. Spending 2 hours writing deterministic OpenCV/InsightFace/SciPy polish scripts saved dozens of video generations and produced noticeably tighter continuity than prompt-engineering alone.
Vertical framing matters for short-form: Repurposing 16:9 music videos into letterboxed vertical strips triggered ~85% instant swipe-away on Shorts; cropping out bottom horizontal captions first and rendering true edge-to-edge full-bleed 9:16 with re-timed vertical ASS captions immediately became our highest-performing clip format.
Full 5:53 output: Forever, Darling (Auto-Renew) on YouTube · Before/After split-screen: YouTube Short
Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. - andrewfallan-alt/ai-music-video-continuity

GitHub - andrewfallan-alt/ai-music-video-continuity: Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment. · GitHub
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
andrewfallan-alt
ai-music-video-continuity Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
LICENSE LICENSE README.md README.md align_lipsync.py align_lipsync.py apply_facelock.py apply_facelock.py fix_iris_continuity.py fix_iris_continuity.py requirements.txt requirements.txt View all files Repository files navigation
Killing AI Slop Across 59 Shots: Multi-Frame Face-Lock, Savitzky-Golay Iris-Lock, and Zero-Freeze Lip-Sync in Python
Demo outputs: 3:42 End-to-End Making-Of Breakdown · 14s Raw vs. Locked Split-Screen · Full 5:53 Music Video (59 shots, 1080p24)
Most AI music videos look unmistakably like "AI slop" after about four seconds—not because the base image or video models lack resolution, but because temporal continuity breaks down across cuts :
Identity drift: Your protagonist's bone structure subtly shifts every time you switch from a wide establishing shot to a singing close-up.
Iris collapse & drift: Lip-sync models routinely either drift stylized eyes back to generic blue-grey, or blow bright eyes out into featureless white flashlights.
Frozen lip-sync tails: Audio-driven lip-sync outputs rarely align cleanly to exact song beat boundaries, leaving dead frozen frames at clip leads/tails or 120 ms mouth-to-vocal skew.
While producing Forever, Darling (Auto-Renew) —a 5:53, 59-shot music video built with Gemini 4 Neon (concept, lyrics, storyboarding, post-production code), Suno (136 BPM F# minor dark-pop × symphonic rock-opera), and Google Flow (Veo + Omni 1.2 Flash) —we solved all three problems deterministically in Python ( OpenCV , InsightFace , SciPy , FFmpeg ) using zero additional video-generation credits .
1. Multi-Frame Golden Embedding Face-Lock ( apply_facelock_v3.py )
Single-image face swapping fails across drastic lighting and expression shifts (e.g., soft morning clinic light vs. golden-hour operatic belting). Instead of anchoring to a single portrait, we extract three approved hero frames across distinct lighting conditions ( S47 close-up, S29 chorus singing, S13 sunlit mid-shot) and compute a weighted normalized ArcFace embedding via InsightFace's buffalo_l :
f47 = app . get ( grab_frame ( "S47_lipsync.mp4" , 0.55 ))[ 0 ]
f29 = app . get ( grab_frame ( "S29_lipsync.mp4" , 0.25 ))[ 0 ]
f13 = app . get ( grab_frame ( "S13_raw.mp4" , 0.60 ))[ 0 ]
emb = 0.45 * f47 . normed_embedding + 0.35 * f29 . normed_embedding + 0.20 * f13 . normed_embedding
emb = emb / np . linalg . norm ( emb )
golden = f29
golden . embedding = emb * np . linalg . norm ( f29 . embedding )
Running inswapper_128.onnx across ProcessPoolExecutor CPU workers with that blended golden identity—combined with 106-point landmark masking—locked facial geometry across all 30+ character shots without burning a single extra video-gen credit.
2. Dual-Mode Iris Physics + Temporal Smoothing ( fix_xena_eyes.py )
Xena's signature visual cue is bioluminescent xenon-lavender irises with crisp dark pupils. Across 59 raw shots, video and lip-sync generators failed in two opposite ways , so we built two deterministic per-shot modes:
tint mode (color drift): Shifts drifted blue-grey irises ( S04 , S29 , S55 ) cleanly into xenon-lavender in HSV/LAB space while strictly masking out pupils, specular catchlights, lashes, sclera, and blink frames.
tame mode (glow blowout): Reconstructs crisp dark pupils and defined iris rings on shots where bright lighting washed the eyes into blank white circles ( S18 , S45 , S47 , S51 ), again respecting natural blinks.
To stop frame-to-frame jitter from per-frame eye detection, pupil/iris coordinates are smoothed temporally across each clip using scipy.signal.savgol_filter before rendering back via piped ffmpeg ( -crf 15 ).
(14-second raw-vs-locked visual comparison: Before vs. After Python Split-Screen )
3. Zero-Frozen-Frame Word-Level Lip-Sync Retiming ( align_xena_lipsync.py )
Audio-driven lip-sync models frequently introduce slight drift or padded lead/tail frames. Rather than manually nudging 11 singing takes in an NLE:
We transcribed the isolated vocal stem ( Demucs → word-level onsets) and wrote a canonical xena_timeline.json across all 59 shots with six surgical editorial boundary overrides ( S17/S18 , S29/S30 , S32/S33 , S45/S46 , S51/S52 , S53/S54 ).
Each Omni 1.2 Flash clip is retimed directly to its exact frame window ( round(start_s * 24) to round(end_s * 24) ) with zero frozen lead/tail frames.
All HUD graphics ( BIOLOGICAL AGE: 48.2 → 25.0 (LOCKED) → AUTO-RENEW: ∞ ) and word-synced karaoke captions are kept 100% out of video prompts (eliminating garbled AI text) and burned cleanly in post-production via ASS subtitles + a 10-panel ffmpeg split-screen montage grid ( build_xena_montage.py ).
Never put readable text or HUDs inside video prompts. Keeping typography 100% in post-production ASS overlays eliminated 90% of wasted video retries.
Post-production beats re-rolling. Spending 2 hours writing deterministic OpenCV/InsightFace/SciPy polish scripts saved dozens of video generations and produced noticeably tighter continuity than prompt-engineering alone.
Vertical framing matters for short-form: Repurposing 16:9 music videos into letterboxed vertical strips triggered ~85% instant swipe-away on Shorts; cropping out bottom horizontal captions first and rendering true edge-to-edge full-bleed 9:16 with re-timed vertical ASS captions immediately became our highest-performing clip format.
Full 5:53 output: Forever, Darling (Auto-Renew) on YouTube · Before/After split-screen: YouTube Short
Deterministic Python post-production pipeline for 59-shot AI music videos: multi-frame ArcFace face-lock, Savitzky-Golay iris-lock, and zero-freeze word-level lip-sync alignment.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

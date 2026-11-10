---
source: "https://github.com/kadirb4rut/voxadub"
hn_url: "https://news.ycombinator.com/item?id=50439152"
title: "Show HN: VoxaDub – Open-source AI video dubbing with speaker-aware voices"
article_title: "GitHub - kadirb4rut/voxadub: VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. · GitHub"
image: "https://repository-images.githubusercontent.com/1388352822/86420991-42b7-4192-ad7c-b1daebad9b07"
author: "kadirb4rut"
captured_at: "2026-11-10T12:11:21Z"
capture_tool: "hn-digest"
hn_id: 50439152
score: 1
comments: 0
posted_at: "2026-11-10T12:06:21Z"
tags:
  - hacker-news
---

# Show HN: VoxaDub – Open-source AI video dubbing with speaker-aware voices

- HN: [50439152](https://news.ycombinator.com/item?id=50439152)
- Source: [github.com](https://github.com/kadirb4rut/voxadub)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T12:06:21Z

## Translation

Title: Show HN: VoxaDub – Open-source AI video dubbing with speaker-aware voices
Article title: GitHub - kadirb4rut/voxadub: VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. · GitHub
Description: VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. - kadirb4rut/voxadub
HN text: Hi HN, I have been working on a project called VoxaDub: an open source local-first AI dubbing tool which can translate videos into other languages while preserving the original voices of the speakers. A major issue when translating is ensuring the translation matches with the length of speech in the target language. Even slight differences in timing can cause unnatural dubbing. It becomes more complex when there are multiple speakers in the original video. I have created a pipeline that uses Whisper/WhisperX to transcribe and align audio, Community-1 to identify speakers, Hy-MT2 to translate contextually and VoxCPM2 to generate speaker aware audio. The main feature of VoxaDub is the ability to improve the translation using the available time for the audio to be spoken for, limit the time changes to within certain bounds and analyse for any timing issues. There is also some experimental Live Dubbing for YouTube videos and live streams through a Chrome extension. This is actively being developed and is not yet production ready. The main project runs fully locally once models are set up, no external translation or TTS APIs are required. Optional GPU usage is available using Colab and Modal. I am pleased to see that the project was recently added to the Ecosystem & Community section of OpenBMB's official VoxCPM GitHub repository (PR #389). While this is just a community project listing it was nice to see this progress. True examples of the English to Spanish, Japanese and Turkish dubbing can be found in the README. Source videos, information on how it was implemented and any limitations are also present. The project is MIT licensed and can be found on GitHub. I would really appreciate any technical feedback on the timing of the translations, speaker consistency and simplifying the local setup. There is a lot of room for improvement so any thoughts would be great!

Article text:
GitHub - kadirb4rut/voxadub: VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. · GitHub
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
kadirb4rut
voxadub Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental.
Readme MIT license Contributing
Security policy Activity Stars
6 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
42 Commits 42 Commits Folders and files
.github .github assets/ branding assets/ branding browser_extension/ live_dubbing browser_extension/ live_dubbing compute_backends compute_backends diarization diarization docs docs dubbing dubbing engine engine live_dubbing live_dubbing scripts scripts tests tests translation translation vocal-remover vocal-remover .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md Video_Translator.py Video_Translator.py install.ps1 install.ps1 install.sh install.sh modal_app.py modal_app.py modal_batch_app.py modal_batch_app.py modal_vllm_smoke.py modal_vllm_smoke.py requirements-ci.txt requirements-ci.txt requirements-cloud.txt requirements-cloud.txt requirements-colab.txt requirements-colab.txt requirements-diarization.txt requirements-diarization.txt requirements-hymt2.txt requirements-hymt2.txt requirements-live.txt requirements-live.txt requirements-modal.txt requirements-modal.txt requirements-vllm-omni.txt requirements-vllm-omni.txt requirements.txt requirements.txt start_gui.bat start_gui.bat start_gui.command start_gui.command start_gui.sh start_gui.sh start_live_dubbing_browser.bat start_live_dubbing_browser.bat start_live_dubbing_browser.sh start_live_dubbing_browser.sh video_translator_gui.py video_translator_gui.py voxcpm_runtime.py voxcpm_runtime.py web_gui.py web_gui.py View all files Repository files navigation
Speaker-aware AI Dubbing for Video and Live Broadcasts
Your video, in another language — with a voice that stays familiar.
Local-first offline dubbing · Experimental live broadcast dubbing
🌟 Listed in OpenBMB/VoxCPM — Ecosystem & Community · PR #389
🎬 Watch the film · Features · Quick Start · Documentation · v0.1.0 release
From source video to dubbed video: dialogue is transcribed, translated with Hy-MT2 , and voiced with speaker-aware VoxCPM2 references before the original scene and background are reconstructed.
🎙️ Offline Dubbing · Process a video fully locally after model setup, or explicitly choose Colab or Modal.
⚡ Live Dubbing — experimental · Capture an authorized YouTube stream or supported video with Local or Colab inference. Set up Live Dubbing →
A cinematic introduction to VoxaDub, featuring speaker-aware multilingual dubbing, duration-aware synchronization, optional lip-sync, and experimental Live Dubbing.
VoxaDub-The-Voice-Shift.mp4
News
🚀 2026-11-03 · VoxaDub v0.1.0 — Experimental Live/VOD dubbing, Hy-MT2 translation, speaker-aware voices, timing safeguards, and new GUI and compute workflows. Release notes · Changelog
Offline Dubbing — Multilingual Demo
🌍 Compare the same 00:20–00:50 interval across the English source and three offline dubs. Turn sound on for each player; on narrow screens, scroll the table horizontally to see both columns.
NASA_JWST_English_20-50_compact.mp4
NASA_JWST_Spanish_20-50_compact.mp4
🇯🇵 AI Dub · Japanese 🇹🇷 AI Dub · Turkish
NASA_JWST_Japanese_20-50_compact.mp4
NASA_JWST_Turkish_20-50_compact.mp4
Source footage: NASA's Goddard Space Flight Center, JWST interview with Dr. Amber Straughn . The translated speech is AI-generated by this project; no NASA endorsement is implied, and NASA is not responsible for the dub.
The earlier 15-second demo release remains available as historical material; it predates the present Hy-MT2 and multi-speaker workflow. The original English and Turkish dub are also available there. A local comparison can be created with the comparison script . See media status .
🎙️ Speaker-aware voices — Community-1 diarization assigns speaker identities when available; clean references condition VoxCPM2. Ambiguous identities fall back safely.
🧠 Local translation — Hy-MT2 supports context and optional terminology. The 1.8B model has bounded tests; 7B and 30B-A3B are selectable but not runtime-verified here.
⏱️ Timing and quality safeguards — Duration-aware refinement, bounded correction, QA, and retry/fallback help fit speech without guaranteeing perfect synchronization.
⚡ Experimental Adaptive Latency Live Dubbing — Capture an authorized YouTube live broadcast or supported regular video with a paired Chrome extension and Local or Colab inference.
☁️ Flexible compute — Run offline jobs locally or explicitly choose Colab CLI G4 or Modal. Selected cloud providers receive the media needed for processing; Modal Live is not available.
🎬 Optional lip-sync — LatentSync is a separate CUDA-only post-processing step.
Initial installation and model downloads require internet access. pyannote/speaker-diarization-community-1 also requires accepting its Hugging Face terms and an HF_TOKEN for the initial download. Keep the unauthenticated browser GUI bound to 127.0.0.1 , and process only media and voices you have permission to use.
Requirement
Base dubbing
Notes
Python 3.10
Required
Main VoxCPM2/WhisperX environment.
FFmpeg + ffprobe
Required
Install through the normal OS package manager and put both on PATH .
Disk
Model-dependent
VoxCPM2 is about 5 GB; Hy-MT2 is about 4.1, 16.1, or 60.1 GB.
GPU
Optional
CUDA is fastest. The saved full local Mac run used CPU on an 8 GB Apple M1; MPS was exercised only in a Hy-MT2 component smoke.
macOS / Linux
git clone https://github.com/kadirb4rut/voxadub.git
cd voxadub
./install.sh
./start_gui.sh
Windows · PowerShell
git clone https: // github.com / kadirb4rut / voxadub.git
cd voxadub
.\install.ps1
.\ start_gui.bat
The installer creates the pinned main .venv and the dependency-isolated Hy-MT2 runtime. It does not choose or download a Hy-MT2 model for you. In the browser GUI, expand Advanced model, speaker and lip-sync settings , select one translation model and click Download Selected . Large downloads never begin merely by launching the application.
Run ./start_gui.sh once for the main panel. It opens Google Chrome using an isolated app profile. On macOS, if Chrome is missing, the terminal asks before downloading Google's installer and installing an app-local copy; other platforms currently require Chrome to be installed separately. No AI model weights are downloaded by browser setup. Its Live Dubbing section has its own GPU choice and Connect & prepare Colab and Open Live Dubbing Browser actions; its Offline Dubbing section keeps its own provider choice. Live Dubbing supports Local and Colab CLI today. Modal Live is unavailable until its streaming service is implemented; the existing Modal integration runs offline jobs. The command-line shortcut ./start_live_dubbing_browser.sh remains available for Local or Colab Live sessions.
The dashboard and Live Dubbing browser use separate Chrome app profiles and do not alter your normal Chrome profile. The Live launcher loads the extension automatically through Chrome's experimental DevTools protocol, so no Web Store or manual extension installation is needed. During the Live browser session a randomly selected loopback debugging port remains open; other local processes can control that isolated browser profile. The macOS launcher uses a mock Keychain, so credentials saved in these app profiles do not have the normal Keychain protection. See Live Dubbing setup for pairing and capture steps.
The web GUI has Local , Colab CLI — G4 GPU , and Modal — full video dubbing choices under Offline dubbing . Select a provider, choose your video and languages, and press Start Dubbing . For cloud providers the GUI connects your account, prepares an isolated GPU environment and downloads missing models before continuing. You can also use Connect & prepare GPU first. Account authorization links and code/confirmation prompts appear in the GUI; Google/Modal consent remains interactive. No notebook, remote folder path, or manually created session is required.
Cloud connection tools are installed separately from the local AI environment when needed. Colab reuses or creates the dedicated video-dubbing-g4 session. Models are cached on the Colab VM and must be downloaded again after that VM ends. Modal uses a connected full-pipeline app with a persistent model Volume. Cloud connections stay open after jobs, errors, browser closure and GUI closure. Only Disconnect GPU requests their shutdown. Allocated cloud resources can continue consuming credits until disconnected. Video output is saved on this computer and remains available after the GPU is disconnected. A fresh GUI launch clears the previous result from the panel but keeps the output file. Account credentials stay with the provider tools. See the Turkish GUI guide and provider details .
Earlier bounded capture, Turkish speech playback, Finish/drain and source-audio restoration were verified on macOS with Colab G4 using Chrome for Testing. The new Google Chrome launcher has separately loaded the extension, opened its side panel, and played the user's YouTube stream; live audio capture, account sign-in and dubbed playback with this new launcher still need end-to-end verification. Local Mac inference, Windows, Linux, and clean Chrome installation are unverified. See the release verification record .
Use SKIP_MODELS=1 ./install.sh or .\install.ps1 -SkipModels to install environments without the vocal-remover and VoxCPM2 weights. Hy-MT2 and diarization weights are always separate explicit downloads.
Offline Dubbing Quick Start and Local Model Management
Hy-MT2 requires transformers>=5.6 , while the main voice/transcription environment stays pinned to transformers==4.51.3 . The project therefore runs Hy-MT2 in third_party/hymt2/.venv and communicates through a persistent JSON-lines subprocess. One worker keeps the selected model loaded for the whole job; stdout is machine-readable and model diagnostics stay on stderr.
python scripts/setup_hymt2.py
python scripts/manage_hymt2_models.py list
python scripts/manage_hymt2_models.py download hymt2-1.8b
Available variants:
Remove one local model safely with python scripts/manage_hymt2_models.py remove MODEL_KEY . Model directories and virtual environments are ignored by Git.
The maintained pyannote/speaker-diarization-community-1 pipeline runs in its own optional

[truncated]

## Original Extract

VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. - kadirb4rut/voxadub

Hi HN, I have been working on a project called VoxaDub: an open source local-first AI dubbing tool which can translate videos into other languages while preserving the original voices of the speakers. A major issue when translating is ensuring the translation matches with the length of speech in the target language. Even slight differences in timing can cause unnatural dubbing. It becomes more complex when there are multiple speakers in the original video. I have created a pipeline that uses Whisper/WhisperX to transcribe and align audio, Community-1 to identify speakers, Hy-MT2 to translate contextually and VoxCPM2 to generate speaker aware audio. The main feature of VoxaDub is the ability to improve the translation using the available time for the audio to be spoken for, limit the time changes to within certain bounds and analyse for any timing issues. There is also some experimental Live Dubbing for YouTube videos and live streams through a Chrome extension. This is actively being developed and is not yet production ready. The main project runs fully locally once models are set up, no external translation or TTS APIs are required. Optional GPU usage is available using Colab and Modal. I am pleased to see that the project was recently added to the Ecosystem & Community section of OpenBMB's official VoxCPM GitHub repository (PR #389). While this is just a community project listing it was nice to see this progress. True examples of the English to Spanish, Japanese and Turkish dubbing can be found in the README. Source videos, information on how it was implemented and any limitations are also present. The project is MIT licensed and can be found on GitHub. I would really appreciate any technical feedback on the timing of the translations, speaker consistency and simplifying the local setup. There is a lot of room for improvement so any thoughts would be great!

GitHub - kadirb4rut/voxadub: VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental. · GitHub
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
kadirb4rut
voxadub Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
VoxaDub — Local-first AI video and live broadcast dubbing with Hy-MT2 translation and speaker-aware VoxCPM2 voices. Live Dubbing is experimental.
Readme MIT license Contributing
Security policy Activity Stars
6 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
42 Commits 42 Commits Folders and files
.github .github assets/ branding assets/ branding browser_extension/ live_dubbing browser_extension/ live_dubbing compute_backends compute_backends diarization diarization docs docs dubbing dubbing engine engine live_dubbing live_dubbing scripts scripts tests tests translation translation vocal-remover vocal-remover .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md Video_Translator.py Video_Translator.py install.ps1 install.ps1 install.sh install.sh modal_app.py modal_app.py modal_batch_app.py modal_batch_app.py modal_vllm_smoke.py modal_vllm_smoke.py requirements-ci.txt requirements-ci.txt requirements-cloud.txt requirements-cloud.txt requirements-colab.txt requirements-colab.txt requirements-diarization.txt requirements-diarization.txt requirements-hymt2.txt requirements-hymt2.txt requirements-live.txt requirements-live.txt requirements-modal.txt requirements-modal.txt requirements-vllm-omni.txt requirements-vllm-omni.txt requirements.txt requirements.txt start_gui.bat start_gui.bat start_gui.command start_gui.command start_gui.sh start_gui.sh start_live_dubbing_browser.bat start_live_dubbing_browser.bat start_live_dubbing_browser.sh start_live_dubbing_browser.sh video_translator_gui.py video_translator_gui.py voxcpm_runtime.py voxcpm_runtime.py web_gui.py web_gui.py View all files Repository files navigation
Speaker-aware AI Dubbing for Video and Live Broadcasts
Your video, in another language — with a voice that stays familiar.
Local-first offline dubbing · Experimental live broadcast dubbing
🌟 Listed in OpenBMB/VoxCPM — Ecosystem & Community · PR #389
🎬 Watch the film · Features · Quick Start · Documentation · v0.1.0 release
From source video to dubbed video: dialogue is transcribed, translated with Hy-MT2 , and voiced with speaker-aware VoxCPM2 references before the original scene and background are reconstructed.
🎙️ Offline Dubbing · Process a video fully locally after model setup, or explicitly choose Colab or Modal.
⚡ Live Dubbing — experimental · Capture an authorized YouTube stream or supported video with Local or Colab inference. Set up Live Dubbing →
A cinematic introduction to VoxaDub, featuring speaker-aware multilingual dubbing, duration-aware synchronization, optional lip-sync, and experimental Live Dubbing.
VoxaDub-The-Voice-Shift.mp4
News
🚀 2026-11-03 · VoxaDub v0.1.0 — Experimental Live/VOD dubbing, Hy-MT2 translation, speaker-aware voices, timing safeguards, and new GUI and compute workflows. Release notes · Changelog
Offline Dubbing — Multilingual Demo
🌍 Compare the same 00:20–00:50 interval across the English source and three offline dubs. Turn sound on for each player; on narrow screens, scroll the table horizontally to see both columns.
NASA_JWST_English_20-50_compact.mp4
NASA_JWST_Spanish_20-50_compact.mp4
🇯🇵 AI Dub · Japanese 🇹🇷 AI Dub · Turkish
NASA_JWST_Japanese_20-50_compact.mp4
NASA_JWST_Turkish_20-50_compact.mp4
Source footage: NASA's Goddard Space Flight Center, JWST interview with Dr. Amber Straughn . The translated speech is AI-generated by this project; no NASA endorsement is implied, and NASA is not responsible for the dub.
The earlier 15-second demo release remains available as historical material; it predates the present Hy-MT2 and multi-speaker workflow. The original English and Turkish dub are also available there. A local comparison can be created with the comparison script . See media status .
🎙️ Speaker-aware voices — Community-1 diarization assigns speaker identities when available; clean references condition VoxCPM2. Ambiguous identities fall back safely.
🧠 Local translation — Hy-MT2 supports context and optional terminology. The 1.8B model has bounded tests; 7B and 30B-A3B are selectable but not runtime-verified here.
⏱️ Timing and quality safeguards — Duration-aware refinement, bounded correction, QA, and retry/fallback help fit speech without guaranteeing perfect synchronization.
⚡ Experimental Adaptive Latency Live Dubbing — Capture an authorized YouTube live broadcast or supported regular video with a paired Chrome extension and Local or Colab inference.
☁️ Flexible compute — Run offline jobs locally or explicitly choose Colab CLI G4 or Modal. Selected cloud providers receive the media needed for processing; Modal Live is not available.
🎬 Optional lip-sync — LatentSync is a separate CUDA-only post-processing step.
Initial installation and model downloads require internet access. pyannote/speaker-diarization-community-1 also requires accepting its Hugging Face terms and an HF_TOKEN for the initial download. Keep the unauthenticated browser GUI bound to 127.0.0.1 , and process only media and voices you have permission to use.
Requirement
Base dubbing
Notes
Python 3.10
Required
Main VoxCPM2/WhisperX environment.
FFmpeg + ffprobe
Required
Install through the normal OS package manager and put both on PATH .
Disk
Model-dependent
VoxCPM2 is about 5 GB; Hy-MT2 is about 4.1, 16.1, or 60.1 GB.
GPU
Optional
CUDA is fastest. The saved full local Mac run used CPU on an 8 GB Apple M1; MPS was exercised only in a Hy-MT2 component smoke.
macOS / Linux
git clone https://github.com/kadirb4rut/voxadub.git
cd voxadub
./install.sh
./start_gui.sh
Windows · PowerShell
git clone https: // github.com / kadirb4rut / voxadub.git
cd voxadub
.\install.ps1
.\ start_gui.bat
The installer creates the pinned main .venv and the dependency-isolated Hy-MT2 runtime. It does not choose or download a Hy-MT2 model for you. In the browser GUI, expand Advanced model, speaker and lip-sync settings , select one translation model and click Download Selected . Large downloads never begin merely by launching the application.
Run ./start_gui.sh once for the main panel. It opens Google Chrome using an isolated app profile. On macOS, if Chrome is missing, the terminal asks before downloading Google's installer and installing an app-local copy; other platforms currently require Chrome to be installed separately. No AI model weights are downloaded by browser setup. Its Live Dubbing section has its own GPU choice and Connect & prepare Colab and Open Live Dubbing Browser actions; its Offline Dubbing section keeps its own provider choice. Live Dubbing supports Local and Colab CLI today. Modal Live is unavailable until its streaming service is implemented; the existing Modal integration runs offline jobs. The command-line shortcut ./start_live_dubbing_browser.sh remains available for Local or Colab Live sessions.
The dashboard and Live Dubbing browser use separate Chrome app profiles and do not alter your normal Chrome profile. The Live launcher loads the extension automatically through Chrome's experimental DevTools protocol, so no Web Store or manual extension installation is needed. During the Live browser session a randomly selected loopback debugging port remains open; other local processes can control that isolated browser profile. The macOS launcher uses a mock Keychain, so credentials saved in these app profiles do not have the normal Keychain protection. See Live Dubbing setup for pairing and capture steps.
The web GUI has Local , Colab CLI — G4 GPU , and Modal — full video dubbing choices under Offline dubbing . Select a provider, choose your video and languages, and press Start Dubbing . For cloud providers the GUI connects your account, prepares an isolated GPU environment and downloads missing models before continuing. You can also use Connect & prepare GPU first. Account authorization links and code/confirmation prompts appear in the GUI; Google/Modal consent remains interactive. No notebook, remote folder path, or manually created session is required.
Cloud connection tools are installed separately from the local AI environment when needed. Colab reuses or creates the dedicated video-dubbing-g4 session. Models are cached on the Colab VM and must be downloaded again after that VM ends. Modal uses a connected full-pipeline app with a persistent model Volume. Cloud connections stay open after jobs, errors, browser closure and GUI closure. Only Disconnect GPU requests their shutdown. Allocated cloud resources can continue consuming credits until disconnected. Video output is saved on this computer and remains available after the GPU is disconnected. A fresh GUI launch clears the previous result from the panel but keeps the output file. Account credentials stay with the provider tools. See the Turkish GUI guide and provider details .
Earlier bounded capture, Turkish speech playback, Finish/drain and source-audio restoration were verified on macOS with Colab G4 using Chrome for Testing. The new Google Chrome launcher has separately loaded the extension, opened its side panel, and played the user's YouTube stream; live audio capture, account sign-in and dubbed playback with this new launcher still need end-to-end verification. Local Mac inference, Windows, Linux, and clean Chrome installation are unverified. See the release verification record .
Use SKIP_MODELS=1 ./install.sh or .\install.ps1 -SkipModels to install environments without the vocal-remover and VoxCPM2 weights. Hy-MT2 and diarization weights are always separate explicit downloads.
Offline Dubbing Quick Start and Local Model Management
Hy-MT2 requires transformers>=5.6 , while the main voice/transcription environment stays pinned to transformers==4.51.3 . The project therefore runs Hy-MT2 in third_party/hymt2/.venv and communicates through a persistent JSON-lines subprocess. One worker keeps the selected model loaded for the whole job; stdout is machine-readable and model diagnostics stay on stderr.
python scripts/setup_hymt2.py
python scripts/manage_hymt2_models.py list
python scripts/manage_hymt2_models.py download hymt2-1.8b
Available variants:
Remove one local model safely with python scripts/manage_hymt2_models.py remove MODEL_KEY . Model directories and virtual environments are ignored by Git.
The maintained pyannote/speaker-diarization-community-1 pipeline runs in its own optional

[truncated]

---
source: "https://github.com/bykof/conversationaly/"
hn_url: "https://news.ycombinator.com/item?id=49362088"
title: "Conversationaly – Privacy-First AI Meeting Assistant"
article_title: "GitHub - bykof/conversationaly: Privacy-first AI meeting assistant. Captures mic + system audio, transcribes live, and writes the summary — entirely on your machine. Tauri/Rust core, transcribe.cpp for STT, bundled llama.cpp. No account, no cloud, no telemetry. · GitHub"
image: "https://opengraph.githubassets.com/f02f58862219d8d49461a33102a04d9f5b58af59106b05e607b97892a9707d36/bykof/conversationaly"
author: "bykof"
captured_at: "2026-08-19T15:22:51Z"
capture_tool: "hn-digest"
hn_id: 49362088
score: 1
comments: 1
posted_at: "2026-08-19T14:26:42Z"
tags:
  - hacker-news
  - translated
---

# Conversationaly – Privacy-First AI Meeting Assistant

- HN: [49362088](https://news.ycombinator.com/item?id=49362088)
- Source: [github.com](https://github.com/bykof/conversationaly/)
- Score: 1
- Comments: 1
- Posted: 2026-08-19T14:26:42Z

## Translation

タイトル: Conversationaly – プライバシー優先の AI 会議アシスタント
記事のタイトル: GitHub - bykof/conversationaly: プライバシー優先の AI 会議アシスタント。マイクとシステムオーディオをキャプチャし、ライブで文字起こしし、要約をすべてマシン上で書き込みます。 Tauri/Rust コア、STT 用の transcribe.cpp、バンドルされた llama.cpp。アカウントもクラウドもテレメトリもありません。 · GitHub
説明: プライバシー優先の AI 会議アシスタント。マイクとシステムオーディオをキャプチャし、ライブで文字起こしし、要約をすべてマシン上で書き込みます。 Tauri/Rust コア、STT 用の transcribe.cpp、バンドルされた llama.cpp。アカウントもクラウドもテレメトリもありません。 - bykof/会話的に

記事本文:
GitHub - bykof/conversationaly: プライバシー優先の AI 会議アシスタント。マイクとシステムオーディオをキャプチャし、ライブで文字起こしし、要約をすべてマシン上で書き込みます。 Tauri/Rust コア、STT 用の transcribe.cpp、バンドルされた llama.cpp。アカウントもクラウドもテレメトリもありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
バイコフ
/
会話的に
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
630 コミット 630 コミット フォルダーとファイル

s
.github .github .playwright-mcp .playwright-mcp バックエンド バックエンド ドキュメント ドキュメント フロントエンド フロントエンド ラマヘルパー ラマヘルパー スクリプト スクリプト .gitignore .gitignore .gitmodules .gitmodules BLUETOOTH_PLAYBACK_NOTICE.md BLUETOOTH_PLAYBACK_NOTICE.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE.md LICENSE.md PRODUCT.md PRODUCT.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プライバシー優先の AI 会議アシスタント
会議を記録し、ライブで文字に起こし、要約をマシン上で作成します。意図的に設定しない限り、アカウントやクラウドの往復は必要ありません。
Conversationaly は、マイクとシステム音声をキャプチャし、会議の進行を文字に起こし、概要を生成するデスクトップ アプリ (macOS、Windows、Linux) です。文字起こしモデルと要約 LLM はどちらもデフォルトでローカルで実行され、どこにも何も送信されません。必要に応じてクラウド プロバイダーを利用できますが、機能ごとにオプトインされます。
これは Meetily のフォークであり、 transcribe.cpp とバンドルされた llama.cpp サイドカーを中心に再構築されています。完全に無料です。有料枠、ライセンス キー、テレメトリはありません。
ローカル転写 — 16 ファミリ (Whisper、Parakeet、Nemotron、Canary、Voxtral、Qwen3-ASR、SenseVoice、Moonshine、GigaAM など) にわたる最大 87 モデルがオンデマンドでダウンロードされます。デフォルト: nemotron-3.5-asr-streaming-0.6b-q8 、39 ロケールにわたる多言語。
ライブ トランスクリプト - ストリーミング ネイティブ モデルは、話している間継続的に文字に起こします。バッチ専用モデルは音声アクティビティによってセグメント化されており、引き続きライブで機能します。
AI が組み込まれており、Ollama は必要ありません。バンドルされている llama ヘルパー サイドカーは、要約のために Gemma 4 をローカルで実行し、オーディオ LLM として直接文字起こしすることもできます。
独自の LLM を導入 — 組み込み AI、Ollama、Claude、

Groq、OpenRouter、OpenAI、または任意の OpenAI 互換エンドポイント。
オプションのクラウド STT — ホスト型トランスクライバーを希望する場合は、Deepgram、ElementalLabs、Groq、または OpenAI。
プロフェッショナルなオーディオミキシング - マイクとシステムオーディオを一緒にキャプチャし、RMS ベースのダッキングとクリッピングを防止します。
ベータ版のインポートと拡張 — 既存の音声ファイルを文字に起こしたり、過去の会議を別のモデルや言語で再文字に起こしたりします。
要約テンプレート — 要約が従う構造を選択または記述し、音声言語とは別に要約言語を設定します。
GPU アクセラレーション — Apple Silicon 上の Metal、CUDA (NVIDIA)、Vulkan (AMD/Intel)、ROCm (Linux 上の AMD)。
ローカル ストレージ — 会議、トランスクリプト、モデルは、ディスク上の SQLite データベースとモデル ディレクトリに存在します。
ビルド済みインストーラー (macOS .dmg 、Windows .exe 、Linux .deb / .rpm / .AppImage ) は、バージョンがタグ付けされている場合にリリース ページで公開されます。
Rust、Node.js、pnpm、cmake が必要です。プラットフォームごとの前提条件については、docs/BUILDING.md を参照してください。
git clone https://github.com/bykof/conversationaly
cd 会話型/フロントエンド
pnpmインストール
# macOS
./clean_build.sh
# Linux (GPU バックエンドを自動検出)
./build-gpu.sh
# Windows
clean_build_windows.bat
Linux の詳細: docs/building_in_linux.md 。
macOS — マイク、およびシステムオーディオの画面録画 (ScreenCaptureKit、macOS 13 以降)。
Windows — マイク;システムオーディオは WASAPI ループバックを使用します。
最初の起動時に、オンボーディングによって 1 つの文字起こしモデルと 1 つの Gemma 4 層がダウンロードされます。それ以降:
マイクとシステム オーディオがキャプチャされ、ミキシングされて、録音に書き込まれます。
同じ混合音声が 16 kHz にリサンプリングされて文字起こしエンジンに供給され、会議の進行中に文字起こしの行が出力されます。
概要を要求すると、トランスクリプトは設定した LLM プロバイダー (場所) に送信されます。

デフォルトではサイドカーです。
上記はすべてローカルプロセスです。クラウド STT プロバイダーとクラウド サマリー プロバイダーは、いずれかを選択してキーを指定した場合にのみ、マシンから離れる唯一のパスです。
単一の Tauri アプリケーション: Rust コア (オーディオ キャプチャ、トランスクリプション、ストレージ、サマリー オーケストレーション) と Next.js フロントエンド。Tauri コマンドとイベントを介して通信します。実行する別のサーバーはありません。
詳細: docs/architecture.md 。
CDフロントエンド
pnpmインストール
./clean_run.sh # macOS: ビルドして実行 (情報ログ)
./clean_run.sh debug # 詳細ログ
clean_run_windows.bat # Windows
./dev-gpu.sh # Linux
pnpm run tauri:dev # プレーン開発モード
pnpm run tauri:dev:metal # 特定の GPU バックエンドを強制します
pnpm 実行 tauri:dev:cuda
pnpm 実行 tauri:dev:vulkan
pnpm 実行 tauri:dev:cpu
寄稿者向けのアーキテクチャに関するメモと規約は CLAUDE.md にあります。 GPU バックエンドの詳細は docs/GPU_ACCELERATION.md にあります。
backend/ ディレクトリは、Tauri が書き換えられる前にアーカイブされた Python/FastAPI サービスです。これはサポートされていないため、アプリのビルドや実行には必要ありません。
問題やプルリクエストは大歓迎です。プロジェクトの構造とガイドラインについては、CONTRIBUTING.md を参照してください。
Conversationaly は Zackriya Solutions による Meetily のフォークであり、MIT ライセンスに基づいて構築されています。
文字起こしは、ggml/wisper.cpp 上に構築された transcribe.cpp 上で実行されます。
ローカル LLM 推論では、 llama-cpp-2 経由で llama.cpp を使用します。
Screenpipe と transcribe-rs からコードを借用しました。
Import & Enhance は Jeremi Joslin によって提供され、Vishnu P S と Mohammed Safvan によって改良されました。
Parakeet および Nemotron 音声モデルの NVIDIA、およびカタログ内の他のモデル ファミリの背後にあるチームに感謝します。
プライバシー優先の AI 会議アシスタント。マイクとシステムオーディオをキャプチャし、ライブで文字起こしし、要約をすべてマシン上で書き込みます。 Tauri/Rust コア、転写

STT 用の be.cpp、バンドルされた llama.cpp。アカウントもクラウドもテレメトリもありません。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Privacy-first AI meeting assistant. Captures mic + system audio, transcribes live, and writes the summary — entirely on your machine. Tauri/Rust core, transcribe.cpp for STT, bundled llama.cpp. No account, no cloud, no telemetry. - bykof/conversationaly

GitHub - bykof/conversationaly: Privacy-first AI meeting assistant. Captures mic + system audio, transcribes live, and writes the summary — entirely on your machine. Tauri/Rust core, transcribe.cpp for STT, bundled llama.cpp. No account, no cloud, no telemetry. · GitHub
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
bykof
/
conversationaly
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
630 Commits 630 Commits Folders and files
.github .github .playwright-mcp .playwright-mcp backend backend docs docs frontend frontend llama-helper llama-helper scripts scripts .gitignore .gitignore .gitmodules .gitmodules BLUETOOTH_PLAYBACK_NOTICE.md BLUETOOTH_PLAYBACK_NOTICE.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE.md LICENSE.md PRODUCT.md PRODUCT.md README.md README.md View all files Repository files navigation
Privacy-First AI Meeting Assistant
Records your meetings, transcribes them live, and writes the summary — on your machine, with no account and no cloud round-trip unless you deliberately configure one.
Conversationaly is a desktop app (macOS, Windows, Linux) that captures your microphone and system audio, transcribes the meeting as it happens, and generates a summary. Transcription models and the summary LLM both run locally by default — nothing is sent anywhere. Cloud providers are available if you want them, but they are opt-in, per-feature.
It is a fork of Meetily , rebuilt around transcribe.cpp and a bundled llama.cpp sidecar. It is fully free — there is no paid tier, no license key, no telemetry.
Local transcription — ~87 models across 16 families (Whisper, Parakeet, Nemotron, Canary, Voxtral, Qwen3-ASR, SenseVoice, Moonshine, GigaAM, …), downloaded on demand. Default: nemotron-3.5-asr-streaming-0.6b-q8 , multilingual across 39 locales.
Live transcript — streaming-native models transcribe continuously as you speak; batch-only models are segmented by voice activity and still work live.
Built-in AI, no Ollama required — a bundled llama-helper sidecar runs Gemma 4 locally for summaries, and can also transcribe directly as an audio LLM.
Bring your own LLM — summaries via Built-in AI, Ollama, Claude, Groq, OpenRouter, OpenAI, or any OpenAI-compatible endpoint.
Optional cloud STT — Deepgram, ElevenLabs, Groq, or OpenAI, if you prefer a hosted transcriber.
Professional audio mixing — microphone and system audio captured together with RMS-based ducking and clipping prevention.
Import & enhance Beta — transcribe existing audio files, or re-transcribe a past meeting with a different model or language.
Summary templates — pick or write the structure your summaries follow, and set the summary language independently of the spoken one.
GPU acceleration — Metal on Apple Silicon, CUDA (NVIDIA), Vulkan (AMD/Intel), ROCm (AMD on Linux).
Local storage — meetings, transcripts, and models live in a SQLite database and a model directory on your disk.
Prebuilt installers (macOS .dmg , Windows .exe , Linux .deb / .rpm / .AppImage ) are published on the Releases page when a version is tagged.
Requires Rust, Node.js, pnpm, and cmake. See docs/BUILDING.md for per-platform prerequisites.
git clone https://github.com/bykof/conversationaly
cd conversationaly/frontend
pnpm install
# macOS
./clean_build.sh
# Linux (auto-detects GPU backend)
./build-gpu.sh
# Windows
clean_build_windows.bat
Linux specifics: docs/building_in_linux.md .
macOS — microphone, plus screen recording for system audio (ScreenCaptureKit, macOS 13+).
Windows — microphone; system audio uses WASAPI loopback.
On first launch, onboarding downloads one transcription model and one Gemma 4 tier. From then on:
Microphone and system audio are captured, mixed, and written to a recording.
The same mixed audio is resampled to 16 kHz and fed to the transcription engine, which emits transcript lines as the meeting runs.
When you ask for a summary, the transcript goes to whichever LLM provider you configured — the local sidecar by default.
Everything above is a local process. Cloud STT and cloud summary providers are the only paths that leave your machine, and only when you select one and supply a key.
A single Tauri application: a Rust core (audio capture, transcription, storage, summary orchestration) and a Next.js frontend, communicating over Tauri commands and events. There is no separate server to run.
Details: docs/architecture.md .
cd frontend
pnpm install
./clean_run.sh # macOS: build and run (info logging)
./clean_run.sh debug # verbose logging
clean_run_windows.bat # Windows
./dev-gpu.sh # Linux
pnpm run tauri:dev # plain dev mode
pnpm run tauri:dev:metal # force a specific GPU backend
pnpm run tauri:dev:cuda
pnpm run tauri:dev:vulkan
pnpm run tauri:dev:cpu
Architecture notes and conventions for contributors live in CLAUDE.md ; GPU backend details in docs/GPU_ACCELERATION.md .
The backend/ directory is an archived Python/FastAPI service from before the Tauri rewrite. It is unsupported and not needed to build or run the app.
Issues and pull requests are welcome. See CONTRIBUTING.md for project structure and guidelines.
Conversationaly is a fork of Meetily by Zackriya Solutions, which it builds on under the MIT license.
Transcription runs on transcribe.cpp , built on ggml / whisper.cpp .
Local LLM inference uses llama.cpp via llama-cpp-2 .
We borrowed some code from Screenpipe and transcribe-rs .
Import & Enhance was contributed by Jeremi Joslin , improved by Vishnu P S and Mohammed Safvan .
Thanks to NVIDIA for the Parakeet and Nemotron speech models, and to the teams behind the other model families in the catalog.
Privacy-first AI meeting assistant. Captures mic + system audio, transcribes live, and writes the summary — entirely on your machine. Tauri/Rust core, transcribe.cpp for STT, bundled llama.cpp. No account, no cloud, no telemetry.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

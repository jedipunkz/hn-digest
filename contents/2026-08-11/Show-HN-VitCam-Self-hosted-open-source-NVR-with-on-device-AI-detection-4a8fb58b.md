---
source: "https://github.com/scwsoft/vitcam"
hn_url: "https://news.ycombinator.com/item?id=49260246"
title: "Show HN: VitCam – Self-hosted, open-source NVR with on-device AI detection"
article_title: "GitHub - scwsoft/vitcam · GitHub"
author: "scwoods"
captured_at: "2026-08-11T16:47:57Z"
capture_tool: "hn-digest"
hn_id: 49260246
score: 1
comments: 0
posted_at: "2026-08-11T15:54:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: VitCam – Self-hosted, open-source NVR with on-device AI detection

- HN: [49260246](https://news.ycombinator.com/item?id=49260246)
- Source: [github.com](https://github.com/scwsoft/vitcam)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T15:54:26Z

## Translation

タイトル: Show HN: VitCam – オンデバイス AI 検出を備えたセルフホスト型オープンソース NVR
記事のタイトル: GitHub - scwsoft/vitcam · GitHub
説明: GitHub でアカウントを作成して、scwsoft/vitcam の開発に貢献します。

記事本文:
GitHub - scwsoft/vitcam · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
スクウソフト
/
ビトカム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
173 コミット 173 コミット coturn coturn docs/images docs/images フロントエンド フロントエンド サーバー サーバー セットアップ setup .gitattributes .gitattributes LICENSE.md LICENSE.md README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
自己ホスト型のオンプレミス AI カメラ監視、ビデオ

eo分析、NVRプラットフォーム
ViTCam は、完全にローカルなオンプレミスの AI を活用したカメラ監視、ビデオ分析、および NVR (ネットワーク ビデオ レコーダー) プラットフォームです。セキュリティ映像を完全に制御する必要がある家庭、企業、組織向けに構築されています。ViTCam は完全に独自のハードウェア上で実行され、すべての録画をローカルに保存し、外部サーバーやクラウド サービスにビデオ データを送信することはありません。
特定のハードウェア ブランド、サブスクリプション プラン、または閉鎖的なエコシステムにユーザーを閉じ込める独自の AI カメラ システムとは異なり、ViTCam はメーカーを問わず、標準ストリーミング プロトコルをサポートするあらゆる IP カメラ、NVR、または DVR で動作します。ペナルティなしでいつでもカメラを交換したり、ハードウェアを変更したり、セットアップを拡張したりできます。
既存の IP カメラと NVR システムを接続し、ViTCam レイヤーに AI を活用した物体検出、リアルタイム ビデオ分析、モーション トリガー録画を重ねて、あらゆるカメラ セットアップを、施設内で何が起こっているかについての深い洞察を備えたインテリジェントな監視システムに変えます。映像はネットワーク上に残ります。いつも。
ViTCam の AI 検出は GPU と CPU の両方で実行されるため、どのマシンでも開始し、後で GPU にアップグレードしてフレーム レートを高め、遅延を減らすことができます。実稼働ワークロードには専用の NVIDIA GPU が推奨されますが、ViTCam の実行には必須ではありません。
ロックインはありません。これまで。専用のハードウェアは必要ありません。必須のクラウド サブスクリプションはありません。ベンダー ソフトウェアに支払いを続ける必要はありません。すでに所有している汎用ハードウェア上で動作するオープンソース ソフトウェアです。
インストール
クイックスタート (Ubuntu)
Raspberry Pi (Debian Bookworm Legacy 64 ビット)
使用法
カメラサーバーのセットアップ
ViTCam は複数のストリーム入力プロトコルをサポートしているため、幅広いカメラ、NVR システム、さらにはインターネット ベースのビデオ ソースでもすべてのプロセスで動作します。

自分のマシン上でローカルに編集され、保存されます。
⚠️ プライバシーと法的通知: パブリックまたはサードパーティのストリーム (例: YouTube Live やパブリック ウェブカメラ) に接続する場合は、そのフィードにアクセスして処理する権利があることを確認してください。 ViTCam は不正な監視を容認しません。
RTSP — NVR / IP カメラ (オンプレミスに推奨):
# Hikvision NVR — チャンネル 1
rtsp://admin:password@192.168.1.100:554/ストリーミング/チャンネル/101
# Dahua NVR — チャンネル 2
rtsp://admin:password@192.168.1.101:554/cam/realmonitor?channel=2&subtype=0
# レオリンクカメラ
rtsp://admin:password@192.168.1.102:554/h264Preview_01_main
# ジェネリック
rtsp://ユーザー名:パスワード@<カメラIP>:<ポート>/<ストリームパス>
HTTP / MJPEG:
http://<カメラの IP>/video.mjpg
http://<カメラIP>/mjpeg/1
http://ユーザー名:パスワード@<カメラIP>:<ポート>/videostream.cgi
YouTube ライブ:
https://www.youtube.com/watch?v=<ライブストリームID>
USB / ローカルカメラ:
# 内蔵カメラまたは最初の USB Web カメラ
ローカル:0
# 2 番目の USB カメラ (複数接続されている場合)
ローカル:1
Windows (WSL2) では、ローカル カメラはホスト経由でアクセスされます。 local:0 が検出されない場合は、カメラが別のアプリケーションで使用されていないことを確認してください。
ViTCam は各ストリームをプルし、ローカルで処理します。オンプレミス カメラの場合、インターネット アクセスやポート転送は必要ありません。
ライブ WebRTC ストリーミング - ブラウザ内の複数のカメラからの低遅延のリアルタイム ビデオ フィード
AI オブジェクト検出 (切り替え可能) — 1 回の切り替えでカメラごとの AI 検出を有効または無効にします。オフの場合、ViTCam は標準の CCTV モニターおよび NVR として動作し、AI 処理のオーバーヘッドなしでストリーミングおよび録画を行います。
GPU および CPU 推論 — NVIDIA CUDA GPU で実行してパフォーマンスを最大化します。自動的に CPU にフォールバックするため、ラップトップや macOS を含むあらゆるマシンで検出が機能します
オブジェクト追跡 — ごとの複数オブジェクト追跡

フレーム間で一貫した ID
モーショントリガー録画 — モーションまたは検出が検出されたときにビデオクリップを自動的に録画します。
連続録画モード — 常時録画
検出イベント ストレージ — すべてのイベントはタイムスタンプ、境界ボックス、信頼スコアとともに Supabase に保存されます。
ビデオ分析ダッシュボード — 検出傾向、オブジェクト数、滞留時間、ヒートマップ、およびカメラごとのアクティビティを時間の経過とともに視覚化します (時間ごと、日ごと、週ごとのビュー)
人および車両のカウント — シーンを通過する人および車両のリアルタイムおよび履歴のカウント
滞留時間追跡 — 追跡されたオブジェクトがゾーン内にどのくらいの時間留まるかを測定します
マルチカメラ管理 — 1 つのインターフェイスから複数のカメラ ストリームを追加、設定、監視します
ビデオ管理 — Supabase Storage バックエンドを使用して録画した映像を参照およびレビューします
システム ログ ビューア - UI からのリアルタイムおよび履歴システム ログ アクセス
日時オーバーレイ — ビデオ ストリーム上の構成可能なタイムスタンプ オーバーレイ
自己ホスト可能 — 完全に独自のハードウェア上で実行します。データがネットワークから離れることはありません
人物の検出、カウント、滞在時間の分析
車両の検出と分類 (乗用車、トラック、オートバイ、バス)
安全ヘルメット / PPE 適合性の検出
スナップショットキャプチャによるカメラごとの検出イベントログ記録
カスタム モデルのサポート (PyTorch .pth )
┌─────────────────────────┐
│ ブラウザ │
│ ウェブインターフェース（UI） │
━━━━━━━━━━━━━━━━━━━━━━┘
│ WebRTC / REST / リアルタイム
┌───────

──────▼────────────────┐
│ カメラサーバー │
│ ライブストリーミングエンジン │ REST API │
│ AI 検出 │ 物体追跡 │
│ 録音エンジン │ │
━━━━━━━━━━━━━━━━━━━━━━┘
│
┌─────┴─────┐
▼ ▼
データベースストレージ
認証、イベントのビデオクリップ
要件
コンポーネント
要件
OS
Ubuntu 22.04 / 24.04 LTS (推奨)、Debian 12、Windows 10/11 (Conda)、macOS 13+、Raspberry Pi OS Debian Bookworm Legacy 64 ビット
CPU
4コア、x86_64
RAM
8GB
ストレージ
50 GB (OS、アプリ、録画用)
GPU
オプション — CPU 推論はすぐにサポートされます。リアルタイムのマルチカメラ ワークロードには NVIDIA GPU が推奨されます
Node.js
18歳以降
パイソン
3.10以降
GPUがないのですか？問題ない。 CUDA 対応 GPU が検出されない場合、ViTCam は自動的に CPU ベースの推論に戻ります。 CPU モードは、テスト、カメラ数の少ないセットアップ、macOS に適しています。複数のストリームにわたるリアルタイム検出には、CUDA を備えた NVIDIA RTX GPU が推奨されます。
推奨 (リアルタイムのマルチカメラ AI ワークロードの場合)
コンポーネント
おすすめ
GPU
NVIDIA RTX 3060 以降 (例: RTX 5090 以降)
VRAM
8GB以上
CUDA
12.x
RAM
16～32GB
単一カメラのセットアップまたは開発の場合、最新の CPU (Apple M シリーズ、Intel Core i7+ など) で十分であり、GPU は必要ありません。
新しい Ubuntu サーバー (root としてログイン)
apt update && apt install -y git
git clone https://github.com/scwsoft/vitcam.git
bash vitcam/setup/install-linux.sh
標準 Ubuntu (sudo を使用するユーザー)
sudo apt update && sudo apt install -y git
git clone https://gi

thub.com/scwsoft/vitcam.git
sudo bash vitcam/setup/install-linux.sh
インストーラーは自動的に次のことを行います。
システムの依存関係をインストールします (ffmpeg、ビルド ツール、nginx)
GPU が存在する場合は NVIDIA CUDA ドライバーを検出してインストールします。GPU が存在しない場合は CPU モードに戻ります。
Docker エンジンをインストールして開始します (コンテナー内の場合はホスト Docker ソケットに接続します)。
Docker Compose 経由で Supabase のクローンを作成して起動し、データベース スキーマを適用します。
一時停止して、.env ファイルを Supabase キーで更新するように求めます
pyenv 経由で Python 3.10 とすべてのバックエンド依存関係をインストールします。
Next.js フロントエンドを構築し、nginx (ポート 80) + systemd サービス (再起動時に自動開始) 経由でデプロイします。
完了したら、http://localhost:8000 で Supabase Studio を開きます。ユーザー名 supabase とパスワード this_password_is_insecure_and_Should_be_updated でログインします。 [認証] → [ユーザー] に移動して ViTCam ログイン アカウントを作成し、 http://localhost:3000 を開きます。
インストール後のサービス管理:
systemctl ステータス vitcam-frontend
systemctl ステータス vitcam-server
journalctl -u vitcam-server -f # ライブサーバーログ
journalctl -u vitcam-frontend -f # ライブ フロントエンド ログ
systemctl で vitcam-server を再起動します
手動インストール (Ubuntu)
自動スクリプトを使用せずに手動でインストールする場合は、次のようにします。
#1. システムの依存関係
sudo apt update && sudo apt upgrade -y
sudo apt install -y gitcurl wget build-essential ffmpeg libgl1 libglib2.0-0 nginx
#2. ドッカー
カール -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
newgrp ドッカー
#3.Node.js
カール -fsSL https://deb.nodesource.com/setup_22.x | sudo -E bash -
sudo apt install -y nodejs
#4. ViTCamのクローンを作成する
git clone https://github.com/scwsoft/vitcam.git
CD ビットカム
#5. スーパーベース
git clone -- Depth 1 https://github.com/supabase/supabase.git
cd supabase/docker && cp .env.example .env

docker 構成 --detach
CD ../..
#6. pyenv経由のPython
カール https://pyenv.run |バッシュ
# pyenv を ~/.bashrc に追加し、次のようにします。
「 $SHELL 」を実行
pyenv インストール 3.10.11 && pyenv グローバル 3.10.11
#7. バックエンドの依存関係
cd サーバー && pip install -r required.txt && cd ..
#8. フロントエンド
cd フロントエンド && npm install && npm run build && cd ..
http://localhost:8000 で Supabase Studio を開きます。ユーザー名 supabase とパスワード this_password_is_insecure_and_Should_be_updated を使用してログインします。 「認証」→「ユーザー」に移動してログイン アカウントを追加し、「プロジェクト設定」→「API」から API キーをコピーし、server/.env とfrontend/.env を SUPABASE_URL とキーで更新します。
ViTCam は、Windows 上で 2 つのバックエンド セットアップ パスをサポートしています。最速のセットアップを行うには、Supabase、Conda 環境、PyTorch、フロントエンド ビルド、および起動スクリプトを自動的に処理する、提供されている PowerShell インストーラーを使用します。
前提条件: Git for Windows 、Node.js 18+ 、Docker Desktop for Windows (実行中)、Anaconda / Miniconda、および最新の NVIDIA ドライバーを備えた NVIDIA GPU。
Anaconda プロンプトを開き (conda が PATH 上でアクティブになるように)、次を実行します。
git clone https://github.com/scwsoft/vitcam.git
cd vitcam\セットアップ
次に、conda Base 環境内からインストーラーを実行します。
（ベース） 。 \install-windows.bat
インストーラーは Anaconda プロンプトから実行する必要があります。環境セットアップが機能するには、conda が PATH 上で利用可能である必要があります。 .bat ファイルは、PowerShell 実行ポリシーを自動的に処理します。
インストール後、 http://localhost:54323 → Authentication → Users の Supabase Studio に移動して最初のログイン アカウントを作成し、Supabase URL と anon キー (Project Settings → API にあります) で .env ファイルを更新し、 http://localhost:3000 を開きます。
ステップ 1 — リポジトリのクローンを作成する (Windows)
PowerShell またはコマンド プロンプトを開いて次を実行します。
git clone https:

// github.com / scwsoft / vitcam.git
CD ビットカム
S

[切り捨てられた]

## Original Extract

Contribute to scwsoft/vitcam development by creating an account on GitHub.

GitHub - scwsoft/vitcam · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
scwsoft
/
vitcam
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
173 Commits 173 Commits coturn coturn docs/ images docs/ images frontend frontend server server setup setup .gitattributes .gitattributes LICENSE.md LICENSE.md README.md README.md View all files Repository files navigation
Self-hosted, on-premises AI camera surveillance, video analytics, and NVR platform
ViTCam is a fully local, on-premises AI-powered camera surveillance, video analytics, and NVR (Network Video Recorder) platform. Built for homes, businesses, and organizations that require complete control over their security footage — ViTCam runs entirely on your own hardware, stores all recordings locally, and never sends video data to any external server or cloud service.
Unlike proprietary AI camera systems that lock you into specific hardware brands, subscription plans, or closed ecosystems, ViTCam works with any IP camera, NVR, or DVR that supports standard streaming protocols — no matter the manufacturer. Swap cameras, change hardware, or scale your setup at any time without penalty.
Connect your existing IP cameras and NVR systems, and ViTCam layers AI-powered object detection, real-time video analytics, and motion-triggered recording on top — turning any camera setup into an intelligent surveillance system with deep insight into what's happening across your premises. Your footage stays on your network. Always.
ViTCam's AI detection runs on both GPU and CPU — so you can get started on any machine and upgrade to a GPU later for higher frame rates and lower latency. A dedicated NVIDIA GPU is recommended for production workloads, but not required to run ViTCam.
No lock-in. Ever. No proprietary hardware required. No mandatory cloud subscription. No vendor software you must keep paying for. Just open-source software running on commodity hardware you already own.
Installation
Quick Start (Ubuntu)
Raspberry Pi (Debian Bookworm Legacy 64-bit)
Usage
Setting Up the Camera Server
ViTCam supports multiple stream input protocols, so it works with a wide range of cameras, NVR systems, and even internet-based video sources — all processed and stored locally on your own machine.
⚠️ Privacy & legal notice: When connecting to any public or third-party stream (e.g. YouTube Live or public webcams), ensure you have the right to access and process that feed. ViTCam does not condone unauthorized surveillance.
RTSP — NVR / IP Cameras (recommended for on-premises):
# Hikvision NVR — Channel 1
rtsp://admin:password@192.168.1.100:554/Streaming/Channels/101
# Dahua NVR — Channel 2
rtsp://admin:password@192.168.1.101:554/cam/realmonitor?channel=2&subtype=0
# Reolink camera
rtsp://admin:password@192.168.1.102:554/h264Preview_01_main
# Generic
rtsp://username:password@<camera-ip>:<port>/<stream-path>
HTTP / MJPEG:
http://<camera-ip>/video.mjpg
http://<camera-ip>/mjpeg/1
http://username:password@<camera-ip>:<port>/videostream.cgi
YouTube Live:
https://www.youtube.com/watch?v=<live-stream-id>
USB / Local Camera:
# Built-in camera or first USB webcam
local:0
# Second USB camera (if multiple are connected)
local:1
On Windows (WSL2), local cameras are accessed via the host. If local:0 isn't detected, ensure the camera is not in use by another application.
ViTCam pulls each stream and processes it locally. For on-premises cameras, no internet access or port forwarding is required.
Live WebRTC streaming — Low-latency real-time video feeds from multiple cameras in your browser
AI object detection (toggleable) — Enable or disable AI detection per camera with a single toggle. When off, ViTCam operates as a standard CCTV monitor and NVR — streaming and recording without any AI processing overhead
GPU & CPU inference — Runs on NVIDIA CUDA GPUs for maximum performance; falls back to CPU automatically so detection works on any machine, including laptops and macOS
Object tracking — Multi-object tracking with persistent IDs across frames
Motion-triggered recording — Automatically records video clips when motion or detections are detected
Continuous recording mode — Always-on recording
Detection event storage — All events stored in Supabase with timestamps, bounding boxes, and confidence scores
Video analytics dashboard — Visualize detection trends, object counts, dwell times, heatmaps, and per-camera activity over time — hourly, daily, and weekly views
People & vehicle counting — Real-time and historical counts of people and vehicles passing through a scene
Dwell time tracking — Measure how long tracked objects remain in a zone
Multi-camera management — Add, configure, and monitor multiple camera streams from one interface
Video management — Browse and review recorded footage with Supabase Storage backend
System logs viewer — Real-time and historical system log access from the UI
Datetime overlay — Configurable timestamp overlays on video streams
Self-hostable — Runs entirely on your own hardware; your data never leaves your network
Person detection, counting, and dwell time analysis
Vehicle detection and classification (car, truck, motorcycle, bus)
Safety helmet / PPE compliance detection
Per-camera detection event logging with snapshot capture
Custom model support (PyTorch .pth )
┌─────────────────────────────────────────────────────────┐
│ Browser │
│ Web Interface (UI) │
└──────────────────────┬──────────────────────────────────┘
│ WebRTC / REST / Realtime
┌──────────────────────▼──────────────────────────────────┐
│ Camera Server │
│ Live Streaming Engine │ REST API │
│ AI Detection │ Object Tracking │
│ Recording Engine │ │
└──────────────────────┬──────────────────────────────────┘
│
┌─────────┴─────────┐
▼ ▼
Database Storage
Auth, Events Video Clips
Requirements
Component
Requirement
OS
Ubuntu 22.04 / 24.04 LTS (recommended), Debian 12, Windows 10/11 (Conda), macOS 13+, Raspberry Pi OS Debian Bookworm Legacy 64-bit
CPU
4 cores, x86_64
RAM
8 GB
Storage
50 GB (for OS, app, and recordings)
GPU
Optional — CPU inference is supported out of the box; an NVIDIA GPU is recommended for real-time multi-camera workloads
Node.js
18 or later
Python
3.10 or later
No GPU? No problem. ViTCam automatically falls back to CPU-based inference if no CUDA-capable GPU is detected. CPU mode is suitable for testing, low-camera-count setups, and macOS. For real-time detection across multiple streams, an NVIDIA RTX GPU with CUDA is recommended.
Recommended (for real-time multi-camera AI workloads)
Component
Recommendation
GPU
NVIDIA RTX 3060 or better (e.g., RTX 5090 or higher)
VRAM
8 GB minimum
CUDA
12.x
RAM
16–32 GB
For single-camera setups or development, a modern CPU (e.g. Apple M-series, Intel Core i7+) is sufficient with no GPU required.
Fresh Ubuntu Server (logged in as root)
apt update && apt install -y git
git clone https://github.com/scwsoft/vitcam.git
bash vitcam/setup/install-linux.sh
Standard Ubuntu (user with sudo)
sudo apt update && sudo apt install -y git
git clone https://github.com/scwsoft/vitcam.git
sudo bash vitcam/setup/install-linux.sh
The installer will automatically:
Install system dependencies (ffmpeg, build tools, nginx)
Detect and install NVIDIA CUDA drivers if a GPU is present — falls back to CPU mode if not
Install and start Docker Engine (or connect to host Docker socket if inside a container)
Clone and start Supabase via Docker Compose, apply the database schema
Pause and prompt you to update your .env files with Supabase keys
Install Python 3.10 via pyenv and all backend dependencies
Build the Next.js frontend and deploy via nginx (port 80) + systemd services (auto-start on reboot)
Once complete, open Supabase Studio at http://localhost:8000 — log in with username supabase and password this_password_is_insecure_and_should_be_updated . Go to Authentication → Users to create your ViTCam login account, then open http://localhost:3000 .
Service management after install:
systemctl status vitcam-frontend
systemctl status vitcam-server
journalctl -u vitcam-server -f # live server logs
journalctl -u vitcam-frontend -f # live frontend logs
systemctl restart vitcam-server
Manual Installation (Ubuntu)
If you prefer to install manually without the automated script:
# 1. System dependencies
sudo apt update && sudo apt upgrade -y
sudo apt install -y git curl wget build-essential ffmpeg libgl1 libglib2.0-0 nginx
# 2. Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER
newgrp docker
# 3. Node.js
curl -fsSL https://deb.nodesource.com/setup_22.x | sudo -E bash -
sudo apt install -y nodejs
# 4. Clone ViTCam
git clone https://github.com/scwsoft/vitcam.git
cd vitcam
# 5. Supabase
git clone --depth 1 https://github.com/supabase/supabase.git
cd supabase/docker && cp .env.example .env
docker compose up --detach
cd ../..
# 6. Python via pyenv
curl https://pyenv.run | bash
# Add pyenv to ~/.bashrc then:
exec " $SHELL "
pyenv install 3.10.11 && pyenv global 3.10.11
# 7. Backend dependencies
cd server && pip install -r requirements.txt && cd ..
# 8. Frontend
cd frontend && npm install && npm run build && cd ..
Open Supabase Studio at http://localhost:8000 — log in with username supabase and password this_password_is_insecure_and_should_be_updated . Go to Authentication → Users to add your login account, copy the API keys from Project Settings → API , and update server/.env and frontend/.env with your SUPABASE_URL and keys.
ViTCam supports two backend setup paths on Windows. For the fastest setup, use the provided PowerShell installer which handles Supabase, Conda environment, PyTorch, frontend build, and startup scripts automatically.
Prerequisites: Git for Windows , Node.js 18+ , Docker Desktop for Windows (running), Anaconda / Miniconda , and an NVIDIA GPU with the latest NVIDIA drivers .
Open Anaconda Prompt (so conda is active on PATH) and run:
git clone https: // github.com / scwsoft / vitcam.git
cd vitcam\setup
Then run the installer from within the conda base environment:
(base) . \install-windows.bat
The installer must be run from Anaconda Prompt — conda must be available on PATH for the environment setup to work. The .bat file handles PowerShell execution policy automatically.
After install, go to Supabase Studio at http://localhost:54323 → Authentication → Users to create your first login account, then update your .env files with your Supabase URL and anon key (found under Project Settings → API ), and open http://localhost:3000 .
Step 1 — Clone the Repository (Windows)
Open PowerShell or Command Prompt and run:
git clone https: // github.com / scwsoft / vitcam.git
cd vitcam
S

[truncated]

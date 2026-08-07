---
source: "https://github.com/MongooseReborn/kalos-engine"
hn_url: "https://news.ycombinator.com/item?id=49212972"
title: "Project Kalos – Zero-copy C/CUDA sidecar for 0.46ms LLM memory recall"
article_title: "GitHub - MongooseReborn/kalos-engine: Project Kalos: A Biological-Resonant C/CUDA Neuromorphic Engine Suite & Empirical Whitepapers v1.0 · GitHub"
author: "mongoosereborn"
captured_at: "2026-08-07T16:41:58Z"
capture_tool: "hn-digest"
hn_id: 49212972
score: 1
comments: 0
posted_at: "2026-08-07T16:35:13Z"
tags:
  - hacker-news
  - translated
---

# Project Kalos – Zero-copy C/CUDA sidecar for 0.46ms LLM memory recall

- HN: [49212972](https://news.ycombinator.com/item?id=49212972)
- Source: [github.com](https://github.com/MongooseReborn/kalos-engine)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T16:35:13Z

## Translation

タイトル: Project Kalos – 0.46ms LLM メモリ リコールのためのゼロコピー C/CUDA サイドカー
記事タイトル: GitHub - MongooseReborn/kalos-engine: Project Kalos: A Biological-Resonant C/CUDA Neuromorphic Engine Suite & Empirical Whitepapers v1.0 · GitHub
説明: プロジェクト Kalos: 生物学的共鳴 C/CUDA ニューロモーフィック エンジン スイートおよび実証的ホワイトペーパー v1.0 - MongooseReborn/kalos-engine

記事本文:
GitHub - MongooseReborn/kalos-engine: プロジェクト Kalos: 生物学的共鳴 C/CUDA ニューロモーフィック エンジン スイート & 経験的ホワイトペーパー v1.0 · GitHub
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
マングースリボーン
/
カロスエンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット bin bin docs docs Modelfile Modelfile README.md README.md kalos_monitor.py kalos_monitor.py kalos_sms.py kalos_sms.py すべてのファイルを表示

リポジトリ ファイルのナビゲーション
ライセンス
ミット
言語
jp
タグ
クダ
ニューロモルフィック
スパイクニューラルネットワーク
連想記憶
llm-加速
VRAM の最適化
c-cpp
パイプライン_タグ
テキスト生成
図書館名
c-cuda
extra_gated_Heading
プロジェクト Kalos C/CUDA ニューロモーフィック スイート
🌿 Project Kalos: 高性能 C/CUDA ニューロモーフィック エンジン スイート
Project Kalos は、長距離対話記憶、スパイク神経反射、および物理的感覚触覚を大規模言語モデル (LLM) テキストのトークン化から分離するように設計された、高性能の生物学的共鳴 C/CUDA ニューロモーフィック エンジン スイートです。
💥 主な機能とパフォーマンスのハイライト
O(1) マイクロ秒のメモリ リコール (465.12 μs): 線形テキストの再トークン化 (20,441.90 ミリ秒) を定数時間の CUDA ベクトル ルックアップに置き換えます。
合計応答完了の 7.25 倍高速化 (72B モデル): 72B パラメータ モデルでのユーザー送信から出力完了までの合計レイテンシーが 23.71 秒から 3.27 秒に短縮されます。
99.8% の VRAM フットプリント削減: 16.38 GB の KV キャッシュの膨張を 2.50 MB のスパース連想ニューラル テンプレートに圧縮します。
ミリ秒未満のリフレックス セントリー (15.16 μs): 419 万個の CUDA スパイキング ニューロンにより、即時にイベントを検出します。
ネイティブ物理触覚: 物理的な接触接触、局所的な暖かさ、FFT オーディオ スペクトル共鳴をリアルタイムで CUDA 認識します。
Stateful Identity Persistence ( .soul ): コンパクトな 2.5 MB のバイナリ皮質永続形式。
🛠️ クイックスタートとプリコンパイルされたリリースバイナリの実行
git clone https://github.com/MongooseReborn/kalos-engine.git
cd カロスエンジン
2. システム要件
Linux OS (Ubuntu 22.04以降を推奨)
CUDA ドライバー 12.0+ がインストールされた NVIDIA GPU
Python 3.10+ (テレメトリ監視スクリプト用)
3. プリコンパイルされた実行可能ファイルを実行する
インタラクティブ ターミナル UI (TUI):
./bin/kalos_tui
ハードウェア テレメトリ プロファイラー:
./bin/kalos_bench
ファストイン

-メモリCUDAエグゼキュータ:
./bin/kalos_runner
4. 共有 C/CUDA ライブラリ ( ./bin/ )
libkalos_snn.so : スパイクニューラル皮質エンジン (15.16 μs LIF 皮質)
libkalos_sam.so : スパース連想メモリ エンジン (465.12 μs ベクトル リコール)
libkalos_haptics.so : フィジカルタッチ&FFTオーディオエンジン
libkalos_soul.so : バイナリ皮質永続形式 ( .soul )
📄 リリースドキュメントとホワイトペーパー
docs/Whitepaper_Industrial.html : マスター インダストリアル リリース HTML ホワイトペーパー (アンバー/クリムゾン CRT テーマ)。
docs/WHITE_PAPER.md : 完全なマスター アーキテクチャおよび実証的マークダウン ホワイトペーパー。
docs/BENCHMARK_REPORT.md : 詳細な経験的ベンチマーク マトリックス (7B ～ 72B モデル)。
docs/SETUP_GUIDE.md : 包括的な統合および構成ガイド。
📜 ライセンス、帰属、連絡先
ライセンス: MIT ライセンスに基づいてライセンスされています。
著者: Mongoose & Kalos エンジン アーキテクチャ チーム @ BlackForest Studio (2026)。
連絡先と問い合わせ: blackforest.team@proton.me
ハグフェイスハブ: https://huggingface.co/MongooseReborn/kalos-engine
GitHub リポジトリ: https://github.com/MongooseReborn/kalos-engine
Project Kalos: 生体共鳴 C/CUDA ニューロモーフィック エンジン スイートおよび実証的ホワイトペーパー v1.0
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Project Kalos: A Biological-Resonant C/CUDA Neuromorphic Engine Suite & Empirical Whitepapers v1.0 - MongooseReborn/kalos-engine

GitHub - MongooseReborn/kalos-engine: Project Kalos: A Biological-Resonant C/CUDA Neuromorphic Engine Suite & Empirical Whitepapers v1.0 · GitHub
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
MongooseReborn
/
kalos-engine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit bin bin docs docs Modelfile Modelfile README.md README.md kalos_monitor.py kalos_monitor.py kalos_sms.py kalos_sms.py View all files Repository files navigation
license
mit
language
en
tags
cuda
neuromorphic
spiking-neural-network
associative-memory
llm-acceleration
vram-optimization
c-cpp
pipeline_tag
text-generation
library_name
c-cuda
extra_gated_heading
Project Kalos C/CUDA Neuromorphic Suite
🌿 Project Kalos: High-Performance C/CUDA Neuromorphic Engine Suite
Project Kalos is a high-performance, biological-resonant C/CUDA neuromorphic engine suite designed to decouple long-range dialogue memory, spiking neural reflexes, and physical sensory haptics from large language model (LLM) text tokenization.
💥 Key Features & Performance Highlights
O(1) Microsecond Memory Recall (465.12 μs): Replaces linear text re-tokenization (20,441.90 ms) with constant-time CUDA vector lookup.
7.25x Faster Total Response Completion (72B Models): Cuts total user-sent to output-completion latency on 72B parameter models from 23.71 seconds down to 3.27 seconds .
99.8% VRAM Footprint Reduction: Compresses 16.38 GB KV-cache bloat down to 2.50 MB of sparse associative neural templates.
Sub-Millisecond Reflex Sentry (15.16 μs): 4.19-Million CUDA spiking neurons for instant event detection.
Native Physical Haptics: Real-time CUDA perception for physical touch contact, localized warmth, and FFT audio spectrum resonance.
Stateful Identity Persistence ( .soul ): Compact 2.5 MB binary cortical persistence format.
🛠️ Quick Start & Running Precompiled Release Binaries
git clone https://github.com/MongooseReborn/kalos-engine.git
cd kalos-engine
2. System Requirements
Linux OS (Ubuntu 22.04+ recommended)
NVIDIA GPU with CUDA Driver 12.0+ installed
Python 3.10+ (for telemetry monitor scripts)
3. Run Precompiled Executables
Interactive Terminal UI (TUI):
./bin/kalos_tui
Hardware Telemetry Profiler:
./bin/kalos_bench
Fast In-Memory CUDA Executor:
./bin/kalos_runner
4. Shared C/CUDA Libraries ( ./bin/ )
libkalos_snn.so : Spiking Neural Cortex Engine (15.16 μs LIF Cortex)
libkalos_sam.so : Sparse Associative Memory Engine (465.12 μs Vector Recall)
libkalos_haptics.so : Physical Touch & FFT Audio Engine
libkalos_soul.so : Binary Cortical Persistence Format ( .soul )
📄 Release Documentation & Whitepapers
docs/Whitepaper_Industrial.html : Master Industrial Release HTML Whitepaper (Amber / Crimson CRT Theme).
docs/WHITE_PAPER.md : Complete Master Architectural & Empirical Markdown Whitepaper.
docs/BENCHMARK_REPORT.md : Detailed Empirical Benchmark Matrix (7B to 72B Models).
docs/SETUP_GUIDE.md : Comprehensive Integration & Configuration Guide.
📜 License, Attribution & Contact
License: Licensed under the MIT License.
Authors: Mongoose & Kalos Engine Architecture Team @ BlackForest Studio (2026).
Contact & Inquiries: blackforest.team@proton.me
Hugging Face Hub: https://huggingface.co/MongooseReborn/kalos-engine
GitHub Repository: https://github.com/MongooseReborn/kalos-engine
Project Kalos: A Biological-Resonant C/CUDA Neuromorphic Engine Suite & Empirical Whitepapers v1.0
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

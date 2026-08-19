---
source: "https://github.com/GLechevalier/nff-core"
hn_url: "https://news.ycombinator.com/item?id=49366373"
title: "Show HN: I let an AI agent brick my $2.51 ESP32 on purpose"
article_title: "GitHub - GLechevalier/nff-core: Firmware can now be coded with agents · GitHub"
image: "https://opengraph.githubassets.com/a1b6bf4bdad4f0e22faa40442d3c6df608d1906aeb1203476bccd2791db386c3/GLechevalier/nff-core"
author: "akiraysh09"
captured_at: "2026-08-19T20:16:07Z"
capture_tool: "hn-digest"
hn_id: 49366373
score: 1
comments: 0
posted_at: "2026-08-19T19:55:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I let an AI agent brick my $2.51 ESP32 on purpose

- HN: [49366373](https://news.ycombinator.com/item?id=49366373)
- Source: [github.com](https://github.com/GLechevalier/nff-core)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T19:55:23Z

## Translation

タイトル: HN を表示: AI エージェントに 2.51 ドルの ESP32 を意図的にブリックさせました
記事タイトル: GitHub - GLechevalier/nff-core: エージェントを使用してファームウェアをコーディングできるようになりました · GitHub
説明: エージェントを使用してファームウェアをコーディングできるようになりました。 GitHub でアカウントを作成して、GLechevalier/nff-core の開発に貢献してください。

記事本文:
GitHub - GLechevalier/nff-core: エージェントを使用してファームウェアをコーディングできるようになりました · GitHub
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
グルシュバリエ
/
nffコア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
277 コミット 277 コミット フォルダーとファイル
.claude .claude .github .github .nff-brain .nff-brain docs docs nff-rs nff-rs nff nff public/ 画像 public/ 画像 スクリプト スクリプト スケッチ スケッチ テスト テスト .env.example .env.example .git

属性 .gitattributes .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGES.md CHANGES.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml要件.txt要件.txt test_migration.ps1 test_migration.ps1 すべてのファイルを表示 リポジトリ ファイルのナビゲーション
nff — コーディングエージェントをハードウェア上で反復させる
ドキュメント ·
クイックスタート ·
mcpツール ·
クリ・
プラットフォーム・
不和
nff は、コーディング エージェントが開発中のベンチや、メンテナンスや診断のための現場で物理ハードウェアを直接制御できるようにする MCP サーバーです。
USB 経由でボードを接続すると、Claude は自律的にシリアル出力の書き込み、コンパイル、フラッシュ、読み取りを行います。 nff-sdk-c ライブラリを使用してデバイスを展開すると、クロードは物理的にアクセスすることなく、リモートでデバイスにアクセスし、クラッシュ状態をキャプチャし、障害を診断し、修正をプッシュできます。
nff は、nff プラットフォームのオープンソース ベンチ CLI です。これは、ESP32 クラスのファームウェア (ベンチ → OTA → フリート診断) を開発、出荷、運用するためのエンドツーエンドのエージェント駆動システムです。このリポジトリとデバイス ライブラリ ( nff-sdk-c ) は、ラップトップとハードウェア上で実行される MIT ライセンスの 2 つの部分です。ホストされるバックエンドは独自のものです。
あなた: 「センサーの初期化シーケンスを実行し、シリアル経由でキャリブレーション値をアサートします」
LLM: [ファームウェアを書き込む] → [コンパイル] → [ESP32 をフラッシュ] → [シリアルを読み取る] → 構造化された出力を返す
あなた：「なぜ現場のユニットがハードフォールトを起こしたのですか？」
LLM: [OTA 経由でパニックをキャプチャ] → [レジスタの読み取り + バックトレース] → 「47 行目でセンサー ISR のスタック オーバーフロー」
特徴
1 回の会話でベンチ ループを実行できます。エディター、ターミナル、シリアル モニターの間で切り替える必要はありません。エージェントは、seri に応答してファームウェアを反復処理します。

al 出力、例外のキャッチ、および再フラッシュ。
ファームウェアが停止した場合のフィールド メンテナンス — クラッシュしたベアメタル MCU にはシェル、SSH、プロセス テーブルがありません。 nff はレジスタ、スタック、メモリ、バックトレースをキャプチャし、障害を説明して回復を促進するクラウド エージェントにルーティングします。これは、Mender、balena、および同様の OTA ツールでは埋めることができないギャップです。これらのツールには、ファームウェア内に生きたネットワーク クライアントが必要です。
無線で送信します。nff otadeploy は、構築したばかりのバイナリを、デバイスごとの追跡と自動ロールバックを備えた段階的な ECDSA 署名付きロールアウトに変換します。 NFF 艦隊 -- 着陸するのを見てください。太田→
board-universal — 約 40 のプラットフォームにわたる PlatformIO の約 1000 以上のボード (すべての ESP32 バリアント、RP2040/Pico、すべての STM32 ファミリ、AVR、SAMD、Teensy、nRF52、Uno R4、RISC-V…)、最初のビルドでツールチェーンが自動インストールされます。 arduino-cli は 2 番目のバックエンドとして引き続き利用できます。ボード →
ライブオンチップデバッグ — 実際のブレークポイント、コールスタック、JTAG/SWD (nff によって駆動される OpenOCD + GDB) を介した変数検査。デバッグ →
ローカルファースト — コンパイル、フラッシュ、監視、デバッグを実行します。MCP ツールはアカウントを必要とせず、ブラウザを開く必要もありません。 OTA、修理、エージェントのみサインインが必要です。
1 つの Rust バイナリ — 自己完結型で、Python ランタイムはなく、Claude Code と同様にバックグラウンドで自己更新されます。
カール -fsSL https://nanoforgeflow.com/install.sh |しー
Windows (PowerShell):
irm https://nanoforgeflow.com/install.ps1 |アイエックス
次に、ボードを接続して実行します。
nff init # ボードを検出し、設定を書き込み、登録 + MCP サーバーを起動します
NFF ドクター # 検証
Claude Code を再起動して MCP サーバーを起動し、必要なことを記述します。
あなた: 「スケッチ/blink_esp32 をフラッシュし、LED がシリアルで切り替わっていることを確認します」
LLM: [コンパイル] → [ESP32 をフラッシュ] → [シリアルを読み取る] → 「LED が 1 で切り替わります」

ヘルツ、確認しました」
フル インストール オプション、--cloud サインイン、および初回実行の詳細: クイック スタート →
127.0.0.1:3010/mcp 上のストリーミング可能な HTTP を介した 34 個のツール。 nff init によってバックグラウンドで開始されます。
完全な署名と戻り形状: mcp ツール →
すべては nanoforgeflow.com/docs にあります。
クイックスタート ·
cli リファレンス ·
構成・
mcpツール ·
クロードコードを使用する ·
デバイスSDK ·
プロビジョニング ·
太田が展開する ·
git-push をデプロイする ·
艦隊の状況・
衝突診断・
オンチップデバッグ ·
パワー・
セキュリティ
リポジトリ内リファレンス ( docs/ ): ボードと USB ID · 自己更新と設定 · ロードマップ · アーキテクチャ
バグと機能リクエストは GitHub の問題に送信されます。 PR を開く前に CONTRIBUTING.md を読んでください。ボードの追加は通常 2 行の変更です。行動規範に従い、SECURITY.md 経由で脆弱性を報告してください。質問やアイデアは Discord で受け付けています。
MIT — 「ライセンス」を参照してください。
著作権 (c) 2026 Gauthier Lechevalier
ファームウェアをエージェントでコーディングできるようになりました
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Firmware can now be coded with agents. Contribute to GLechevalier/nff-core development by creating an account on GitHub.

GitHub - GLechevalier/nff-core: Firmware can now be coded with agents · GitHub
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
GLechevalier
/
nff-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
277 Commits 277 Commits Folders and files
.claude .claude .github .github .nff-brain .nff-brain docs docs nff-rs nff-rs nff nff public/ images public/ images scripts scripts sketches sketches tests tests .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGES.md CHANGES.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml requirements.txt requirements.txt test_migration.ps1 test_migration.ps1 View all files Repository files navigation
nff — let coding agents iterate on hardware
docs ·
quick start ·
mcp tools ·
cli ·
platform ·
discord
nff is an MCP server that gives coding agents direct control over physical hardware — on the bench during development, and in the field for maintenance and diagnosis.
Connect your board over USB and Claude writes, compiles, flashes, and reads serial output autonomously. Deploy devices with the nff-sdk-c library and Claude can reach them remotely: capture crash state, diagnose failures, and push fixes — without physical access.
nff is the open-source bench CLI of the nff platform — an end-to-end, agent-driven system for developing, shipping, and operating ESP32-class firmware (bench → OTA → fleet diagnosis). This repo and the device library ( nff-sdk-c ) are the two MIT-licensed pieces that run on your laptop and hardware; the hosted backend is proprietary.
you: "Run the sensor init sequence and assert the calibration values over serial"
LLM: [writes firmware] → [compiles] → [flashes ESP32] → [reads serial] → returns structured output
you: "Why did the unit in the field just hard-fault?"
LLM: [captures panic over OTA] → [reads registers + backtrace] → "Stack overflow in your sensor ISR at line 47"
features
the bench loop, in one conversation — no switching between editor, terminal, and serial monitor. The agent iterates on firmware in response to serial output, catches exceptions, and reflashes.
field maintenance when the firmware is dead — a crashed bare-metal MCU has no shell, no SSH, no process table. nff captures registers, stack, memory, and backtrace and routes them to a cloud agent that explains the failure and drives recovery. This is the gap Mender, balena, and similar OTA tools cannot fill: they need a living network client inside the firmware.
ship it over the air — nff ota deploy turns the binary you just built into a staged, ECDSA-signed rollout with per-device tracking and automatic rollback; nff fleet --watch shows it land. ota →
board-universal — any of PlatformIO's ~1000+ boards across ~40 platforms (every ESP32 variant, RP2040/Pico, all STM32 families, AVR, SAMD, Teensy, nRF52, Uno R4, RISC-V…), toolchain auto-installed on first build. arduino-cli remains available as a second backend. boards →
live on-chip debugging — real breakpoints, call stacks, and variable inspection over JTAG/SWD (OpenOCD + GDB, driven by nff). debug →
local-first — compile, flash, monitor, debug, and the MCP tools need no account and never open a browser. Only OTA, repair , and agent require a sign-in.
one Rust binary — self-contained, no Python runtime, and it self-updates in the background like Claude Code does.
curl -fsSL https://nanoforgeflow.com/install.sh | sh
Windows (PowerShell):
irm https: // nanoforgeflow.com / install.ps1 | iex
Then plug in your board and run:
nff init # detects the board, writes config, registers + starts the MCP server
nff doctor # verify
Restart Claude Code so it picks up the MCP server, then just describe what you want:
you: "Flash sketches/blink_esp32 and confirm the LED is toggling over serial"
LLM: [compiles] → [flashes ESP32] → [reads serial] → "LED toggling at 1 Hz, confirmed"
Full install options, --cloud sign-in, and first-run detail: quick start →
34 tools over streamable HTTP on 127.0.0.1:3010/mcp , started in the background by nff init .
Full signatures and return shapes: mcp tools →
everything lives at nanoforgeflow.com/docs :
quick start ·
cli reference ·
configuration ·
mcp tools ·
using claude code ·
device sdk ·
provisioning ·
ota deploys ·
git-push deploys ·
fleet status ·
crash diagnosis ·
on-chip debug ·
power ·
security
In-repo reference ( docs/ ): boards & USB ids · self-update & config · roadmap · architecture
Bugs and feature requests go to GitHub Issues ; read CONTRIBUTING.md before opening a PR — adding a board is usually a two-line change. Please follow the Code of Conduct , and report vulnerabilities via SECURITY.md . Questions and ideas are welcome on Discord .
MIT — see LICENSE .
Copyright (c) 2026 Gauthier Lechevalier
Firmware can now be coded with agents
Readme MIT license Code of conduct
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

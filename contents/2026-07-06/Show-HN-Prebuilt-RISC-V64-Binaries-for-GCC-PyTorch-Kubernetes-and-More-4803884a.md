---
source: "https://github.com/Cloud-V-10xE/RISC-V-software"
hn_url: "https://news.ycombinator.com/item?id=48803068"
title: "Show HN: Prebuilt RISC-V64 Binaries for GCC, PyTorch, Kubernetes, and More"
article_title: "GitHub - Cloud-V-10xE/RISC-V-software: This repository contains all the binaries for the RISCV supported architecture · GitHub"
author: "alitariq4589"
captured_at: "2026-07-06T11:55:55Z"
capture_tool: "hn-digest"
hn_id: 48803068
score: 2
comments: 0
posted_at: "2026-07-06T11:05:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Prebuilt RISC-V64 Binaries for GCC, PyTorch, Kubernetes, and More

- HN: [48803068](https://news.ycombinator.com/item?id=48803068)
- Source: [github.com](https://github.com/Cloud-V-10xE/RISC-V-software)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T11:05:16Z

## Translation

タイトル: HN を表示: GCC、PyTorch、Kubernetes など用の事前構築済み RISC-V64 バイナリ
記事のタイトル: GitHub - Cloud-V-10xE/RISC-V-software: このリポジトリには、RISCV でサポートされるアーキテクチャのすべてのバイナリが含まれています · GitHub
説明: このリポジトリには、RISCV でサポートされるアーキテクチャ - Cloud-V-10xE/RISC-V-software のすべてのバイナリが含まれています

記事本文:
GitHub - Cloud-V-10xE/RISC-V-software: このリポジトリには、RISCV でサポートされるアーキテクチャのすべてのバイナリが含まれています · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クラウド-V-10xE
/
RISC-V ソフトウェア
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
243 コミット 243 コミット .github .githubcloudv-buildscloudv-builds docs/packages docs/packages github-actions-riscv github-actions-riscv github-runner-riscv github-runner-riscv Contributing.md Contributing.md readme.md readme.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
RISC-V 64 ビット Linux 用の事前構築済みバイナリ — 実際のハードウェア上にネイティブに構築
📦 パッケージを参照 · 📖 ドキュメント · 🤝 貢献
RISC-V 上でソフトウェアを実行することは、思った以上に困難です。公式リリースでは riscv64 がスキップされます。パッケージ マネージャーは古いバージョンを出荷します。 RISC-V ハードウェア上で最新の GCC、Go、Python、または Kubernetes ツールを必要とする開発者は、ソースからの構築に何時間も費やすことになります。
このプロジェクトは、人気のあるパッケージを RISE Runners 上でネイティブに構築し、毎月自動的に更新されるすぐに使用できるバイナリを公開することで、この問題を解決します。
パッケージ アーカイブにアクセスして必要なバージョンをダウンロードするか、リリースから直接 wget を使用します。
# 例: 最新の GCC をインストールする
wget https://github.com/Cloud-V-10xE/RISC-V-software/releases/download/ < タグ > /gcc- < バージョン > -riscv64-linux.tar.gz
tar -xzf gcc- < バージョン > -riscv64-linux.tar.gz -C /usr/local/
Dockerイメージを使用する
すべての Docker イメージはマルチ アーキテクチャ (amd64 + riscv64) です。
# Kubernetes コンポーネント イメージをプルする
docker pull Cloudv10x/kube-apiserver:最新
🏗️ 仕組み
┌─────────────────────────┐
│ GitHub アクション │
│ │
│ 個別のビルドワークフロー (build-gcc.yml など) │
│ │ │
│ ▼ で実行します │
│ ┌───

─────┐ ┌───────────┐ │
│ │ RISE │ │ ubuntu-最新 (x86) │ │
│ │ │ │ Docker イメージ用 │ │
│ │ ランナーの皆様 │ ━━━━━━━━━━━┘ │
│ └─────────┘ │
│ │ │
│ ▼ アーティファクトをアップロードする │
│ 集中リリース ワークフロー (毎日) │
│ │ │
│ ▼ GitHub リリースを作成 │
│ GitHub リリース ──────► GitHub Pages サイト │
━━━━━━━━━━━━━━━━━━━━━━━┘
各パッケージには、 .github/workflows/ に独自のワークフロー ファイルがあります。中央のリリース ワークフローは毎日実行され、各ビルド ワークフローから最新の成功したアーティファクトを収集し、それらを 1 つの GitHub リリースにバンドルします。 GitHub Pages サイトはリリースごとに自動的に再生成されます。
詳細な説明については、docs/architecture.md を参照してください。
新しいパッケージを追加するか、失敗したビルドを修正しますか?ステップバイステップのガイドについては、CONTRIBUTING.md を参照してください。
.github/pages/packages.json にエントリを追加します。
ワークフロー ファイル .github/workflows/build-<package>.yml を追加します。
docs/packages/<package>.md にドキュメントを追加します。
このリポジトリ内のワークフロー ファイル、ツール、ドキュメントは MIT ライセンスに基づいてライセンスされています。
事前に構築された各バイナリは、独自のアップストリーム ライセンスを保持します。
このリポジトリには、RISCV でサポートされるアーキテクチャのすべてのバイナリが含まれています
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
26
RISC-V リリース 2026-07-06
最新の
2026 年 7 月 6 日
+ 25 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
ありました

ロード中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

This repository contains all the binaries for the RISCV supported architecture - Cloud-V-10xE/RISC-V-software

GitHub - Cloud-V-10xE/RISC-V-software: This repository contains all the binaries for the RISCV supported architecture · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
Cloud-V-10xE
/
RISC-V-software
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
243 Commits 243 Commits .github .github cloudv-builds cloudv-builds docs/ packages docs/ packages github-actions-riscv github-actions-riscv github-runner-riscv github-runner-riscv Contributing.md Contributing.md readme.md readme.md View all files Repository files navigation
Prebuilt binaries for RISC-V 64-bit Linux — built natively on real hardware
📦 Browse Packages · 📖 Docs · 🤝 Contribute
Getting software running on RISC-V is harder than it should be. Official releases skip riscv64. Package managers ship outdated versions. Developers who need the latest GCC, Go, Python, or Kubernetes tooling on RISC-V hardware end up spending hours building from source.
This project solves that by building popular packages natively on RISE Runners and publishing ready-to-use binaries which are updated automatically every month.
Visit the package archive and download the version you need, or use wget directly from a release:
# Example: install the latest GCC
wget https://github.com/Cloud-V-10xE/RISC-V-software/releases/download/ < tag > /gcc- < version > -riscv64-linux.tar.gz
tar -xzf gcc- < version > -riscv64-linux.tar.gz -C /usr/local/
Use a Docker image
All Docker images are multi-arch (amd64 + riscv64):
# Pull a Kubernetes component image
docker pull cloudv10x/kube-apiserver:latest
🏗️ How it works
┌─────────────────────────────────────────────────────┐
│ GitHub Actions │
│ │
│ Individual build workflows (build-gcc.yml etc.) │
│ │ │
│ ▼ runs on │
│ ┌─────────────┐ ┌──────────────────────┐ │
│ │ RISE │ │ ubuntu-latest (x86) │ │
│ │ │ │ for Docker images │ │
│ │ Runners │ └──────────────────────┘ │
│ └─────────────┘ │
│ │ │
│ ▼ uploads artifact │
│ Central Release workflow (daily) │
│ │ │
│ ▼ creates GitHub release │
│ GitHub Releases ──────► GitHub Pages site │
└─────────────────────────────────────────────────────┘
Each package has its own workflow file in .github/workflows/ . The central release workflow runs daily, collects the latest successful artifact from each build workflow, and bundles them into a single GitHub release. The GitHub Pages site regenerates automatically after each release.
See docs/architecture.md for a detailed explanation.
Want to add a new package or fix a failing build? See CONTRIBUTING.md for a step-by-step guide.
Add an entry to .github/pages/packages.json
Add a workflow file .github/workflows/build-<package>.yml
Add documentation to docs/packages/<package>.md
The workflow files, tooling, and documentation in this repository are licensed under the MIT License .
Each prebuilt binary retains its own upstream license.
This repository contains all the binaries for the RISCV supported architecture
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
26
RISC-V Release 2026-07-06
Latest
Jul 6, 2026
+ 25 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

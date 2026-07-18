---
source: "https://github.com/arnabk/agentgrove"
hn_url: "https://news.ycombinator.com/item?id=48955708"
title: "AgentGrove – local workspace for AI coding agents in Git worktrees"
article_title: "GitHub - arnabk/agentgrove: High-performance local developer workspace — editor, terminals, AI chat, git diff, notes, prompt queue. Rust BE + SolidJS FE. Cross-platform. · GitHub"
author: "arnabk"
captured_at: "2026-07-18T06:53:46Z"
capture_tool: "hn-digest"
hn_id: 48955708
score: 2
comments: 0
posted_at: "2026-07-18T06:23:29Z"
tags:
  - hacker-news
  - translated
---

# AgentGrove – local workspace for AI coding agents in Git worktrees

- HN: [48955708](https://news.ycombinator.com/item?id=48955708)
- Source: [github.com](https://github.com/arnabk/agentgrove)
- Score: 2
- Comments: 0
- Posted: 2026-07-18T06:23:29Z

## Translation

タイトル: AgentGrove – Git ワークツリー内の AI コーディング エージェント用のローカル ワークスペース
記事のタイトル: GitHub - arnabk/agentgrove: 高性能ローカル開発者ワークスペース — エディター、ターミナル、AI チャット、git diff、メモ、プロンプト キュー。 Rust BE + SolidJS FE。クロスプラットフォーム。 · GitHub
説明: 高性能のローカル開発者ワークスペース — エディター、ターミナル、AI チャット、git diff、メモ、プロンプト キュー。 Rust BE + SolidJS FE。クロスプラットフォーム。 - アルナブク/エージェントグローブ

記事本文:
GitHub - arnabk/agentgrove: 高性能のローカル開発者ワークスペース — エディター、ターミナル、AI チャット、git diff、メモ、プロンプト キュー。 Rust BE + SolidJS FE。クロスプラットフォーム。 · GitHub
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
アルナブク
/
エージェントグローブ
公共
通知
署名する必要があります

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
382 コミット 382 コミット .audit-context .audit-context .github .github .sqlx .sqlx apps/ web apps/ web crates crates docker docker docs docs scripts scripts .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md justfile justfile package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
高性能、低フットプリントのオープンソースのローカル開発者ワークスペース。
Rust バックエンド + SolidJS フロントエンド。クロスプラットフォーム (Linux、macOS、Windows)。
完全な機能リストについては、docs/features.md を参照してください。
サムネールをクリックすると、画面の記録が表示されます。すべてのビデオはライブ アプリの実際のキャプチャであり、分離されたデモ スタック上の Docker コンテナ内に記録されるため、ホスト開発サーバーには影響がありません。
# 前提条件をインストールします (Homebrew を含む macOS)
brew install ノード pnpm だけ
# Rust ツールチェーン — プロジェクトピン 1.95 (rust-toolchain.toml 経由)
カール --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs |しー
# クローンを作成して実行する
git clone https://github.com/arnabk/agentgrove.git
CD エージェントグローブ
just dev # http://localhost:5173 で BE (ホット リロード) + FE (HMR) を開始します
Linux では、ディストリビューションのパッケージ マネージャーを使用して同じパッケージをインストールします (例: apt install nodejs pnpm または pacman -S node pnpm just )。 Windows ユーザーは、winget install Rustlang.Rustup OpenJS.NodeJS pnpm.just または Rustup および pnpm インストーラーを使用できます。
詳細なドキュメントはすべて docs/ の下にあります。
高性能のローカル開発者ワークスペース — エディター、ターミナル、AI チャット、git diff、メモ、プロンプト キュー。 Rust BE + SolidJS FE。クロスプラットフォーム

。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
128
v0.1.41
最新の
2026 年 7 月 18 日
+ 127 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

High-performance local developer workspace — editor, terminals, AI chat, git diff, notes, prompt queue. Rust BE + SolidJS FE. Cross-platform. - arnabk/agentgrove

GitHub - arnabk/agentgrove: High-performance local developer workspace — editor, terminals, AI chat, git diff, notes, prompt queue. Rust BE + SolidJS FE. Cross-platform. · GitHub
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
arnabk
/
agentgrove
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
382 Commits 382 Commits .audit-context .audit-context .github .github .sqlx .sqlx apps/ web apps/ web crates crates docker docker docs docs scripts scripts .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md justfile justfile package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
High-performance, low-footprint, open-source local developer workspace.
Rust backend + SolidJS frontend. Cross-platform (Linux, macOS, Windows).
See docs/features.md for the full feature list.
Click any thumbnail to watch the screen recording. All videos are real captures of the live app, recorded inside the Docker container on the isolated demo stack so the host dev server is untouched.
# Install prerequisites (macOS with Homebrew)
brew install node pnpm just
# Rust toolchain — project pins 1.95 via rust-toolchain.toml
curl --proto ' =https ' --tlsv1.2 -sSf https://sh.rustup.rs | sh
# Clone and run
git clone https://github.com/arnabk/agentgrove.git
cd agentgrove
just dev # starts BE (hot reload) + FE (HMR) on http://localhost:5173
On Linux, install the same packages with your distro's package manager (e.g., apt install nodejs pnpm or pacman -S node pnpm just ). Windows users can use winget install Rustlang.Rustup OpenJS.NodeJS pnpm.just or the rustup and pnpm installers.
All detailed docs live under docs/ :
High-performance local developer workspace — editor, terminals, AI chat, git diff, notes, prompt queue. Rust BE + SolidJS FE. Cross-platform.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
128
v0.1.41
Latest
Jul 18, 2026
+ 127 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

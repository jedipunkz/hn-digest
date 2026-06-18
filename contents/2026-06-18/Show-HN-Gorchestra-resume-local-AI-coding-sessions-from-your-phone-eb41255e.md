---
source: "https://github.com/jgennari/gorchestra"
hn_url: "https://news.ycombinator.com/item?id=48585767"
title: "Show HN: Gorchestra – resume local AI coding sessions from your phone"
article_title: "GitHub - jgennari/gorchestra: A Go-based orchestration platform for AI coding agents with real-time event streaming, session replay, and durable execution history. · GitHub"
author: "cybrjoe"
captured_at: "2026-06-18T16:15:45Z"
capture_tool: "hn-digest"
hn_id: 48585767
score: 3
comments: 0
posted_at: "2026-06-18T14:18:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Gorchestra – resume local AI coding sessions from your phone

- HN: [48585767](https://news.ycombinator.com/item?id=48585767)
- Source: [github.com](https://github.com/jgennari/gorchestra)
- Score: 3
- Comments: 0
- Posted: 2026-06-18T14:18:43Z

## Translation

タイトル: HN を表示: Gorchestra – 携帯電話からローカル AI コーディング セッションを再開
記事タイトル: GitHub - jgennari/gorchestra: リアルタイムのイベント ストリーミング、セッション リプレイ、永続的な実行履歴を備えた AI コーディング エージェント用の Go ベースのオーケストレーション プラットフォーム。 · GitHub
説明: リアルタイムのイベント ストリーミング、セッション リプレイ、耐久性のある実行履歴を備えた AI コーディング エージェント用の Go ベースのオーケストレーション プラットフォーム。 - jgennari/ゴーチェストラ
HN テキスト: 私は数週間にわたってこのアイデアの一部のバージョンをドッグフーディングしてきましたが、今週末から再実装を共有したいと考えていました。基本的には、無制限の量のコーデックス (推奨) またはクロード エージェントを制御する Web サーバーです。 it allows me to context switch more easily, as well as kick off tasks and instruct agents while i'm away from my desk.ストレージは SQLite であり、エージェントのステータスに関係なく永続的です。 for codex you can upload images and queue messages just like the CLI.私は Mac を使用しているので、brew 経由でインストールを効率化しましたが、別のプラットフォームで問題が発生した場合は問題を解決してください。誰かが役に立つことを願っています!

記事本文:
GitHub - jgennari/gorchestra: リアルタイムのイベント ストリーミング、セッション リプレイ、永続的な実行履歴を備えた AI コーディング エージェント用の Go ベースのオーケストレーション プラットフォーム。 · GitHub
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
じげんなり
/
ゴーチェストラ
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .github/ workflows .github/ workflows cmd/ app cmd/ app docs docs external 内部パッケージング/ homebrew パッケージング/ homebrew scripts scripts web web .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum package.json package.json View all files Repository files navigation
The private, permanent control room for AI coding work.
Run multiple Codex and Claude sessions, stream every event, inspect files and git state, and keep the whole story.
エージェントが作業を実行します。 Gorchestra conducts the performance.
どこにいても開発を進めてください。 Gorchestra は、AI コーディング作業のためのプライベートで永続的なコントロール ルームを提供します。デスクトップまたはモバイルで開き、エージェント間で複数のセッションを実行し、フォローアップ メッセージをキューに入れ、スクリーンショットをアップロードし、計画モードに切り替え、ファイルと git の状態を検査し、更新、再接続、再起動後に同じ永続的な履歴に戻ります。
走る
Coordinate
Inspect
覚えておいてください
Launch Codex or Claude sessions from one local server.
Queue work, attach images, and switch between fast and planning flows.
Browse files, inspect git-aware changes, and edit text in-session.
注文されたセッション履歴を SQLite に保存します。
Tune model, reasoning, service tier, and execution mode.
Keep multiple sessions moving at once from desktop or mobile.
Jump from file-change diffs straight into Monaco.
コンテキストを失うのではなく、リプレイで再接続します。
なぜ存在するのか
Coding agents are useful, but serious work falls apart when everything is trapped in a single chat window. Gorchestra turns each run into an operating surface: start a session, watch what the agent is doing, inspect fi

le and git changes, queue the next step, and keep the full history after refreshes or restarts.
これはローカルで作業を行うために構築されています。
ワークスペースを選択して、Codex または Claude を起動します。
Watch messages, thinking, tool calls, logs, errors, and file edits as they happen.
現在の実行がまだ動作している間に、フォローアップ プロンプトをキューに入れます。
アプリを離れることなく、変更されたファイルを開いたり、差分を確認したり、git の状態を検査したり、Markdown やテキストを編集したりできます。
後でもう一度戻って、同じ順序付けられたセッション履歴を確認してください。
セッションを開始し、ワークスペースを選択して、エージェントを実行します。 Gorchestra はライブ トランスクリプト、ワークスペース ツール、セッション コントロールをまとめて管理するため、ターミナル タブ、エディタ ウィンドウ、ログの間を行き来する必要がなくなります。
モデル、推論作業、サービス層、計画モード、危険モードなどの Codex オプションを調整します。
メッセージ、思考、ツール呼び出し、コマンド出力、ファイル編集、エラー、デバッグ イベントを 1 つのトランスクリプトで追跡します。
実行がアクティブな間は入力を続け、メッセージをキューに入れ、画像を添付し、実行で入力が必要なときにエージェントが要求したプロンプトに応答します。
Browse, search, preview, and edit workspace files from the side rail.
ファイル変更の差分を確認し、変更されたファイルのエディターに直接ジャンプします。
Refresh or reconnect without losing the session history.
Gorchestra は、React UI が内部に埋め込まれた 1 つのローカル バイナリとして実行することを目的としています。
brew install jgennari/tap/gorchestra
ゴーチェストラ --open
公開されたタップは jgennari/homebrew-tap です。この式は、Go を使用してタグ付けされたソース アーカイブから Gorchestra を構築し、Gorchestra バイナリをインストールします。
Gorchestra をバックグラウンド サービスとして実行します。
brew サービスの開始 jgennari/tap/gorchestra
http://127.0.0.1:15173 を開く
brew services stop jgennari/tap/gorchestra
Homebrew サービスは $(brew --prefix)/etc/gorchestra/gorchestra.env を読み取ります。 Edit that file and res

tart the service to change the port, data directory, workspace roots, or Codex binary path.
Download the archive for your platform from GitHub Releases, unpack it, and run the binary.
tar -xzf gorchestra_ < version > _ < os > _ < arch > .tar.gz
./gorchestra --open
Windows:
Expand-Archive .\gorchestra_ < version > _windows_ < arch > .zip - DestinationPath .\gorchestra
.\gorchestra\ gorchestra.exe -- open
リリース対象:
Real Codex sessions require the Codex CLI to be available on PATH , or configured with --codex-bin .
Start Gorchestra and open the browser:
ゴーチェストラ --open
By default, Gorchestra binds to 127.0.0.1:8080 and stores SQLite data in the OS app data location.
gorchestra --host 127.0.0.1 --port 8081
gorchestra --config ~ /.config/gorchestra/gorchestra.env
gorchestra --data-dir ~ /.gorchestra-dev
gorchestra --workspace /path/to/repo
gorchestra --workspace-root /path/to/allowed/root
gorchestra --codex-bin /path/to/codex
gorchestra --codex-model gpt-5
gorchestra --codex-sandbox workspace-write
gorchestra --codex-network-access=false
gorchestra --codex-web-search=cached
gorchestra --version
--data-dir creates the directory if needed and stores SQLite at <data-dir>/gorchestra.db . --db is still available as an exact SQLite path override and takes precedence over --data-dir .
macOS: ~/Library/Application Support/Gorchestra/gorchestra.db
Linux: $ XDG_DATA_HOME/gorchestra/gorchestra.db
Linux fallback: ~/.local/share/gorchestra/gorchestra.db
同等の環境には、 GORCHESTRA_HOST 、 GORCHESTRA_PORT 、 GORCHESTRA_DATA_DIR 、 GORCHESTRA_DB 、 GORCHESTRA_WORKSPACE 、 GORCHESTRA_OPEN 、および Codex フラグに一致する GORCHESTRA_CODEX_* 変数が含まれます。
Config files use the same env-style names:
GORCHESTRA_HOST=127.0.0.1
GORCHESTRA_PORT=15173
GORCHESTRA_DATA_DIR=/opt/homeb

リュー/ヴァール/ゴーケストラ
GORCHESTRA_WORKSPACE=~
GORCHESTRA_WORKSPACE_ROOTS=~
GORCHESTRA_OPEN=false
GORCHESTRA_CODEX_BIN=コーデックス
GORCHESTRA_CONFIG is the environment equivalent of --config . GORCHESTRA_WORKSPACE_ROOTS は、OS パスリスト区切り文字 (macOS/Linux では :) で区切られた複数のパスを受け入れます。
rm -rf " $HOME /Library/Application Support/Gorchestra "
rm -rf " ${XDG_DATA_HOME :- $HOME / .local / share} /gorchestra "
Codex shell commands run with network access enabled by default. Codex ネイティブ Web 検索は、デフォルトではライブ モードで実行されます。 --codex-web-search=cached または --codex-web-search=disabled を使用して変更してください。
フロントエンド アセットが埋め込まれたリリース バイナリをビルドします。
バンランビルド
This installs frontend dependencies with Bun, builds the Vite app, stages web/dist into internal/webassets/dist , runs go test ./... , builds dist/gorchestra , and writes dist/SHA256SUMS .
./dist/gorchestra --open
Internal/webassets/dist の下にステージングされたアセットがコミットされるため、test ./... に進み、チェックアウトから build ./cmd/app 作業に進みます。 bun run build は、リリース バイナリをコンパイルする前に、最新の Vite 出力からそのディレクトリを更新します。
CDウェブ
バンランテスト
バンランビルド
Production:
bun run build
./dist/gorchestra --version
dist/SHA256SUMS には、ローカル リリース アーティファクトの SHA-256 チェックサムが含まれています。
A Go-based orchestration platform for AI coding agents with real-time event streaming, session replay, and durable execution history.
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
14
ゴーチェストラ v0.1.16
最新の
2026 年 6 月 17 日
+ 13 releases
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Go-based orchestration platform for AI coding agents with real-time event streaming, session replay, and durable execution history. - jgennari/gorchestra

i've been dogfooding some version of this idea for a few weeks now and i'd thought i'd share my re-implementation from this weekend. basically it's a webserver that controls and unlimited amount of codex (preferred) or claude agents. it allows me to context switch more easily, as well as kick off tasks and instruct agents while i'm away from my desk. the storage is sqlite and persistent, regardless of agent status. for codex you can upload images and queue messages just like the CLI. i'm on mac so i've streamlined installation via brew, but open an issue if you have problems on another platform. hope someone finds it useful!

GitHub - jgennari/gorchestra: A Go-based orchestration platform for AI coding agents with real-time event streaming, session replay, and durable execution history. · GitHub
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
jgennari
/
gorchestra
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .github/ workflows .github/ workflows cmd/ app cmd/ app docs docs internal internal packaging/ homebrew packaging/ homebrew scripts scripts web web .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum package.json package.json View all files Repository files navigation
The private, permanent control room for AI coding work.
Run multiple Codex and Claude sessions, stream every event, inspect files and git state, and keep the whole story.
Agents perform work. Gorchestra conducts the performance.
Take your development wherever you are. Gorchestra gives you a private, permanent control room for AI coding work: open it on desktop or mobile, run multiple sessions across agents, queue follow-up messages, upload screenshots, switch into planning mode, inspect files and git state, and come back to the same durable history after refreshes, reconnects, and restarts.
Run
Coordinate
Inspect
Remember
Launch Codex or Claude sessions from one local server.
Queue work, attach images, and switch between fast and planning flows.
Browse files, inspect git-aware changes, and edit text in-session.
Store ordered session history in SQLite.
Tune model, reasoning, service tier, and execution mode.
Keep multiple sessions moving at once from desktop or mobile.
Jump from file-change diffs straight into Monaco.
Reconnect with replay instead of losing context.
Why It Exists
Coding agents are useful, but serious work falls apart when everything is trapped in a single chat window. Gorchestra turns each run into an operating surface: start a session, watch what the agent is doing, inspect file and git changes, queue the next step, and keep the full history after refreshes or restarts.
It is built for getting work done locally:
Pick a workspace and start Codex or Claude.
Watch messages, thinking, tool calls, logs, errors, and file edits as they happen.
Queue follow-up prompts while the current run is still working.
Open changed files, review diffs, inspect git state, and edit Markdown or text without leaving the app.
Come back later and see the same ordered session history.
Start a session, choose a workspace, and let the agent run. Gorchestra keeps the live transcript, workspace tools, and session controls together so you do not have to bounce between terminal tabs, editor windows, and logs.
Tune Codex options like model, reasoning effort, service tier, planning mode, and dangerous mode.
Follow messages, thinking, tool calls, command output, file edits, errors, and debug events in one transcript.
Keep typing while a run is active, queue messages, attach images, and answer agent-requested prompts when a run needs input.
Browse, search, preview, and edit workspace files from the side rail.
Review file-change diffs and jump straight into the editor for the changed file.
Refresh or reconnect without losing the session history.
Gorchestra is meant to run as one local binary with the React UI embedded inside it.
brew install jgennari/tap/gorchestra
gorchestra --open
The published tap is jgennari/homebrew-tap ; the formula builds Gorchestra from the tagged source archive with Go and installs the gorchestra binary.
Run Gorchestra as a background service:
brew services start jgennari/tap/gorchestra
open http://127.0.0.1:15173
brew services stop jgennari/tap/gorchestra
The Homebrew service reads $(brew --prefix)/etc/gorchestra/gorchestra.env . Edit that file and restart the service to change the port, data directory, workspace roots, or Codex binary path.
Download the archive for your platform from GitHub Releases, unpack it, and run the binary.
tar -xzf gorchestra_ < version > _ < os > _ < arch > .tar.gz
./gorchestra --open
Windows:
Expand-Archive .\gorchestra_ < version > _windows_ < arch > .zip - DestinationPath .\gorchestra
.\gorchestra\ gorchestra.exe -- open
Release targets:
Real Codex sessions require the Codex CLI to be available on PATH , or configured with --codex-bin .
Start Gorchestra and open the browser:
gorchestra --open
By default, Gorchestra binds to 127.0.0.1:8080 and stores SQLite data in the OS app data location.
gorchestra --host 127.0.0.1 --port 8081
gorchestra --config ~ /.config/gorchestra/gorchestra.env
gorchestra --data-dir ~ /.gorchestra-dev
gorchestra --workspace /path/to/repo
gorchestra --workspace-root /path/to/allowed/root
gorchestra --codex-bin /path/to/codex
gorchestra --codex-model gpt-5
gorchestra --codex-sandbox workspace-write
gorchestra --codex-network-access=false
gorchestra --codex-web-search=cached
gorchestra --version
--data-dir creates the directory if needed and stores SQLite at <data-dir>/gorchestra.db . --db is still available as an exact SQLite path override and takes precedence over --data-dir .
macOS: ~/Library/Application Support/Gorchestra/gorchestra.db
Linux: $ XDG_DATA_HOME/gorchestra/gorchestra.db
Linux fallback: ~/.local/share/gorchestra/gorchestra.db
Environment equivalents include GORCHESTRA_HOST , GORCHESTRA_PORT , GORCHESTRA_DATA_DIR , GORCHESTRA_DB , GORCHESTRA_WORKSPACE , GORCHESTRA_OPEN , and the GORCHESTRA_CODEX_* variables matching the Codex flags.
Config files use the same env-style names:
GORCHESTRA_HOST=127.0.0.1
GORCHESTRA_PORT=15173
GORCHESTRA_DATA_DIR=/opt/homebrew/var/gorchestra
GORCHESTRA_WORKSPACE=~
GORCHESTRA_WORKSPACE_ROOTS=~
GORCHESTRA_OPEN=false
GORCHESTRA_CODEX_BIN=codex
GORCHESTRA_CONFIG is the environment equivalent of --config . GORCHESTRA_WORKSPACE_ROOTS accepts multiple paths separated by the OS path-list separator ( : on macOS/Linux).
rm -rf " $HOME /Library/Application Support/Gorchestra "
rm -rf " ${XDG_DATA_HOME :- $HOME / .local / share} /gorchestra "
Codex shell commands run with network access enabled by default. Codex native web search runs in live mode by default; use --codex-web-search=cached or --codex-web-search=disabled to change it.
Build the release binary with embedded frontend assets:
bun run build
This installs frontend dependencies with Bun, builds the Vite app, stages web/dist into internal/webassets/dist , runs go test ./... , builds dist/gorchestra , and writes dist/SHA256SUMS .
./dist/gorchestra --open
Staged assets under internal/webassets/dist are committed so go test ./... and go build ./cmd/app work from a checkout. bun run build refreshes that directory from the latest Vite output before compiling the release binary.
cd web
bun run test
bun run build
Production:
bun run build
./dist/gorchestra --version
dist/SHA256SUMS contains SHA-256 checksums for local release artifacts.
A Go-based orchestration platform for AI coding agents with real-time event streaming, session replay, and durable execution history.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
14
Gorchestra v0.1.16
Latest
Jun 17, 2026
+ 13 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

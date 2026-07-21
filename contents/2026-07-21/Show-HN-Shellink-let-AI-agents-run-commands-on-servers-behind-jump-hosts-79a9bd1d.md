---
source: "https://github.com/jie123108/Shellink"
hn_url: "https://news.ycombinator.com/item?id=48991134"
title: "Show HN: Shellink – let AI agents run commands on servers behind jump hosts"
article_title: "GitHub - jie123108/Shellink: One stable CLI and daemon for SSH, local PTY sessions, file transfer, and remote editing — built for AI agents and humans alike. · GitHub"
author: "jie123108"
captured_at: "2026-07-21T12:20:18Z"
capture_tool: "hn-digest"
hn_id: 48991134
score: 2
comments: 0
posted_at: "2026-07-21T12:02:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Shellink – let AI agents run commands on servers behind jump hosts

- HN: [48991134](https://news.ycombinator.com/item?id=48991134)
- Source: [github.com](https://github.com/jie123108/Shellink)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T12:02:43Z

## Translation

タイトル: Show HN: Shelllink – AI エージェントがジャンプ ホストの背後にあるサーバーでコマンドを実行できるようにする
記事のタイトル: GitHub - jie123108/Shellink: SSH、ローカル PTY セッション、ファイル転送、リモート編集用の 1 つの安定した CLI とデーモン — AI エージェントと人間の両方のために構築されています。 · GitHub
説明: SSH、ローカル PTY セッション、ファイル転送、リモート編集用の 1 つの安定した CLI とデーモン - AI エージェントと人間の両方のために構築されています。 - jie123108/シェルリンク

記事本文:
GitHub - jie123108/Shellink: SSH、ローカル PTY セッション、ファイル転送、リモート編集用の 1 つの安定した CLI とデーモン - AI エージェントと人間の両方のために構築されています。 · GitHub
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
jie123108
/
シェルリンク
公共
通知

イケーション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows cli cli docs/ スクリーンショット docs/ スクリーンショット プロトコル プロトコル スクリプト スクリプト サーバー サーバー スキル スキル Web Web .gitignore .gitignore AGENTS_INSTALL.md AGENTS_INSTALL.md CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md install.sh install.sh package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SSH、ローカル PTY セッション、ファイル転送、リモート編集用の 1 つの安定した CLI とデーモン - AI エージェントと人間の両方のために構築されています。
特徴・
クイックスタート ·
ドキュメント ·
貢献する ·
中文档
AI エージェントはジャンプ ホストを経由してターゲット サーバー上でコマンドを実行します。これは完全に Shelllink CLI を通じて実行されます。
Shelllink は、SSH およびローカル ターミナル セッションを使用して作業する AI エージェントと人間のためのオープンソース セッション ミドルウェアです。接続プロファイル、対話型セッション、コマンド実行、ファイル転送、リモート編集、およびセッション状態のための 1 つの安定した CLI を提供します。ローカル デーモンはセッションを所有し、ブラウザーに適した Web UI と HTTP および WebSocket 互換 API を公開します。
リポジトリ: https://github.com/jie123108/Shellink
セキュリティ警告: コマンド プロファイルは、設定されたコマンドをサーバー プロセス ユーザーとして実行します。 Shelllink は信頼できる環境でのみ実行し、 SHELLINK_TOKEN を保護してください。 Shelllink では、AI エージェントが運用サーバーに直接接続できるようになります。これには実際の運用上のリスクが伴うため、読み取り専用、検査、または分析の操作を選択するか、エージェントを非運用環境に向けるようにしてください。
Shelllink が AI エージェント向けに設計されている理由
で

失速
バイナリのインストール (推奨)
クライアントモードとサーバーモード
クライアントモード
Web UI ダッシュボード (セッション リスト):
Web UI ライブ セッション ターミナル (xterm.js) (ジャンプホスト ホップを含む):
CLI/TUI セッション ビュー (シェルリンク cli ):
SSH およびカスタム コマンド セッション: SSH で直接接続するか、PTY 内でコマンド/スクリプトを実行します。
マルチホップ ログイン オートメーション: connectType: "command" は、マルチレベル要塞ホスト、ジャンプ ホスト メニュー、OTP プロンプト、およびレガシー認証フローの想定スクリプトおよび同様のスクリプトをサポートします。
コマンドの実行と状態の処理: コマンドを実行し、状態を検査し、対話型入力を提供し、次の操作の前にセッションの準備が整うのを待ちます。
ジャンプ ホスト間でのファイル転送: SFTP を必要とする代わりに、既存の PTY を介してファイルをアップロードおよびダウンロードします。これにより、カスタム コマンド チェーンとマルチホップ接続を通じて転送が機能し続けます。
リモート ファイル編集: Shellink セッション編集を使用して正確な置換を適用します。
監査履歴と監視: セッションの入出力は監査と再生のための履歴として保持されますが、MANUAL モードでは人間が制御して監視したり、プロンプトに応答したりできます。
CLI/TUI、Web UI、REST、および WebSocket: 自動化のための安定した --json 出力を備えた、ワークフローに適合するインターフェイスを使用します。
暗号化されたプロファイルとオプションの単一ファイル バイナリ: 資格情報は SQLite で保護され、Bun ビルドは macOS と Linux で利用できます。
Shelllink が AI エージェント向けに設計されている理由
Shelllink は、AI エージェントが信頼性の高いアクセスを必要とするが、人間の制御を推測したりバイパスしたりしてはいけない状況を中心に設計されています。
スクリプト化可能なマルチホップ アクセス: エージェントは、各サイトの認証フローを独自に実装する代わりに、複数の要塞ホップに対して Expect またはその他のカスタム ログイン スクリプトを再利用できます。
PTY ベースの継続性: コマンドとファイル転送は同じ PTY パスを使用するため、エージェントは

直接の SFTP または SSH エンドポイントが利用できない場合でも、ジャンプ ホストを介して動作できます。
人間による監視: AUTO モードはエージェントのアクション用です。 MANUAL モードでは、自動化が続行される前に、承認、OTP 入力、または介入のために端末を人に渡します。
監査可能な操作: セッションの入出力履歴とデーモン ログにより、オペレーターはエージェントが何を見たのか、何を行ったのかを確認するための記録が得られます。
観察可能な進行状況: 明示的なセッション状態と機械が読み取り可能な CLI 出力により、エージェントは実行するか、入力を送信するか、待機するか、人間の助けを求めるかを決定できます。
リポジトリは、エージェント スキル用の追加の非ワークスペース ディレクトリが 1 つある npm ワークスペース モノリポジトリです。
フローチャート LR
クライアント["CLI / TUI"] -->|Unixソケット、MessagePack RPC|デーモン[「デーモンサービス」]
デーモン --> セッション["SSH / ローカル PTY セッション"]
デーモン --> ストレージ["SQLite 状態と暗号化された資格情報"]
デーモン --> ゲートウェイ["HTTP REST / WebSocket ゲートウェイ / Web UI"]
読み込み中
Unix ソケットは主要なローカル インターフェイスです。 HTTP、WebSocket、および Web UI は互換性とブラウザー層です。新しいローカル オートメーションでは、ソケット プロトコルを直接実装するのではなく、CLI を使用する必要があります。
ランタイム: Node.js >=22.19.0 ; Bun はコンパイルされたバイナリではオプションです。
言語: npm ワークスペースを使用した TypeScript。
サーバー: Fastify、ssh2、node-pty、ws、better-sqlite3 および Drizzle ORM 経由の SQLite。
プロトコル: MessagePack フレーミングと Zod 検証。
Web: Vue 3、Vite、Naive UI、Pinia、および xterm.js。
テスト: V8 をカバーし、Docker を利用したエンドツーエンドのフィクスチャを備えた Vitest。
AI コーディング エージェントを使用していますか?このプロンプトをエージェントのチャットに貼り付けます。
https://raw.githubusercontent.com/jie123108/Shellink/main/AGENTS_INSTALL.mdに従って、Shelllinkとそのスキルのインストールを手伝ってください。
事前に構築されたバイナリが macOS および Linux ( darwin-arm64 、 darwin-x64 、 linux-arm64 、 linux-

x64)。 Windows はサポートされていません。代わりにソースからビルドします (以下を参照)。
カール -fsSL https://raw.githubusercontent.com/jie123108/Shellink/main/install.sh -o /tmp/shelllink-install.sh
cat /tmp/shelllink-install.sh
sh /tmp/shelllink-install.sh
必要に応じて、リリース タグをピン留めします。
カール -fsSL https://raw.githubusercontent.com/jie123108/Shellink/main/install.sh -o /tmp/shelllink-install.sh
cat /tmp/shelllink-install.sh
sh /tmp/shelllink-install.sh --バージョン v0.1.0
システム全体を /usr/local/bin (ほとんどのシステムではすでに PATH 上にあります) にインストールするには、 sudo を使用して実行します。
sudo sh /tmp/shelllink-install.sh --dir /usr/local/bin
このスクリプトは、GitHub Releases から一致するアセットをダウンロードし、 SHA256SUMS.txt を検証し、デフォルトで $HOME/.local/bin に Shelink をインストールします。そのディレクトリが PATH 上にない場合、インストーラーはシェル プロファイル ( ~/.zshrc 、 ~/.bashrc など) に import PATH=... 行を追加します。新しいシェルを開いて (または PATH を再エクスポートして)、次のことを確認します。
Docker。Docker を利用したエンドツーエンド テストを実行する場合のみ。
ただし、スタンドアロンのバイナリを構築したい場合のみ。
git クローン https://github.com/jie123108/Shellink.git
cd シェルリンク
npmインストール
npm ビルドを実行する
構築された CLI は ./node_modules/.bin/shelllink で入手できます。現在のプラットフォーム用のスタンドアロン実行可能ファイルをビルドするには、次の手順を実行します。
npm run build:binary
mkdir -p " $HOME /.local/bin "
cp dist/shelllink " $HOME /.local/bin/shelllink "
エクスポート PATH= " $HOME /.local/bin: $PATH "
クロスビルド スクリプトは、 macos-arm64 、 macos-x64 、 linux-arm64 、および linux-x64 で使用できます。
対話型ターミナル UI を開始します。必要に応じて、ユーザーレベルのデーモンが自動的に起動されます。
シェルリンク CLI
デーモンを確認し、機械可読エージェントのリファレンスを出力します。
シェルリンクサーバーのステータス --json
シェルリンクエージェントドキュメント
資格情報をシェル履歴に保存せずに、コマンド プロファイルとセッションを作成します。

シェルリンク プロファイルの作成 --input - --json << ' JSON '
{"名前":"ローカルシェル","接続タイプ":"コマンド","コマンド":"/bin/sh"}
JSON
シェルリンク プロファイル リスト --query local --json
シェルリンクセッション作成 --profile < プロファイル ID > --json
shelllink session exec < session-id > --command ' uname -a ' --json
シェルリンクセッションクローズ<セッションID> --json
SSH プロファイルの場合は、--input FILE または stdin を介してパスワード、パスフレーズ、および秘密キーを渡します。コマンドライン引数には決して入れないでください。
ソースから実行する場合、npm run dev はサーバーと Vite 開発 UI を起動します。 http://localhost:5173 を開きます。構築されたデーモンは、デフォルトで http://127.0.0.1:7070/shellink/ui/ で組み込み UI を提供します。正確な URL は、 Shelink Server run または Shelink Server status --json によって表示されます。
カスタム コマンド プロファイルは、単一の直接 SSH ログインでは表現できない接続を対象としています。 connectType を command に設定し、サーバー ホストで使用可能なコマンドを指定します。次に例を示します。
{
"name" : "バスティオンメニュー" ,
"connectType" : " コマンド " ,
"command" : " /opt/shelllink/login-through-bastion.exp を期待します "
}
Shelllink は PTY でコマンドを開始し、SSH セッションで使用されるのと同じコマンド、入力、履歴、ファイル転送、および編集 API を公開します。したがって、expect スクリプトは、複数の要塞ホストと認証ステップを 1 つの Shelllink セッションとして実行できます。スクリプト、その依存関係、およびその資格情報は、デーモン ユーザーとしてサーバー ホスト上で実行されます。信頼できるコマンドのみを実行します。
注: Shelllink 自体はホストの自動ログイン プロセスを処理しません。自動ログイン (マルチホップ ホップ、ジャンプ ホスト メニュー、OTP プロンプト、その他の対話型認証フロー) の場合は、コマンド プロファイルによって駆動される期待スクリプトを使用します。 sshpass (パスワードのみのログイン) またはキーベースの認証を備えた ssh および ProxyJump / ProxyComma などのその他のツール

環境によっては、nd も機能する場合があります。
クライアントモード
Shellink 実行可能ファイルがクライアントです。 TUI コマンドとリソース コマンド ( profile 、 session 、および webhook ) は、Unix ソケットを介してローカル デーモンに接続します。ソケットが使用できない場合、クライアントは切り離されたデーモンを起動し、準備が整うまで待機してからリクエストを送信します。 --json を追加すると、安定した機械可読な 1 行の出力が得られます。
シェルリンクサーバーの起動 # バックグラウンドで起動
shelllink サーバーのステータス --json # 正常性とアクティブなセッションを検査します
shelllink サーバー ログ --lines 100 # 最近のデーモン ログを読み取る
シェルリンクサーバーの再起動 # デーモンを再起動します
シェルリンクサーバー停止 # デーモンを停止します
シェルリンクサーバーの実行 # フォアグラウンドで実行
サーバーの実行は、開発またはプロセス管理者に適しています。デーモンは常にその Unix ソケットを開始します。 HTTP/WebSocket サポートはデフォルトで有効になっており、 SHELLINK_HTTP_ENABLED=false で無効にできます。 HTTP はデフォルトで 127.0.0.1 にバインドします。リモートアクセスが意図的である場合にのみ、SHELLINK_HOST=0.0.0.0 を設定します。
localhost または 127.0.0.1 からのローカル要求では、永続的なセッション レコードの削除を含む認証を省略できます。リモート HTTP および WebSocket リクエストには認証が必要です: Bearer <SHELLINK_TOKEN> (または WebSocket トークン クエリ)

[切り捨てられた]

## Original Extract

One stable CLI and daemon for SSH, local PTY sessions, file transfer, and remote editing — built for AI agents and humans alike. - jie123108/Shellink

GitHub - jie123108/Shellink: One stable CLI and daemon for SSH, local PTY sessions, file transfer, and remote editing — built for AI agents and humans alike. · GitHub
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
jie123108
/
Shellink
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows cli cli docs/ screenshots docs/ screenshots protocol protocol scripts scripts server server skills skills web web .gitignore .gitignore AGENTS_INSTALL.md AGENTS_INSTALL.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md install.sh install.sh package-lock.json package-lock.json package.json package.json View all files Repository files navigation
One stable CLI and daemon for SSH, local PTY sessions, file transfer, and remote editing — built for AI agents and humans alike.
Features ·
Quick Start ·
Documentation ·
Contributing ·
中文文档
An AI agent hops through a jump host and runs commands on the target server — driven entirely through the Shellink CLI.
Shellink is an open-source session middleware for AI agents and humans working with SSH and local terminal sessions. It provides one stable CLI for connection profiles, interactive sessions, command execution, file transfer, remote editing, and session state. A local daemon owns the sessions and exposes a browser-friendly Web UI plus HTTP and WebSocket compatibility APIs.
Repository: https://github.com/jie123108/Shellink
Security warning: command profiles execute their configured command as the server process user. Run Shellink only in a trusted environment and protect SHELLINK_TOKEN . Shellink also lets an AI agent connect directly to production servers; this carries real operational risk, so prefer read-only, inspection, or analysis operations, or point the agent at non-production environments.
Why Shellink Is Designed for AI Agents
Installation
Install binary (recommended)
Client and Server Modes
Client mode
Web UI dashboard (session list):
Web UI live session terminal (xterm.js), including a jump-host hop:
CLI/TUI session view ( shellink cli ):
SSH and custom command sessions: Connect directly with SSH, or run a command/script inside a PTY.
Multi-hop login automation: connectType: "command" supports expect and similar scripts for multi-level bastion hosts, jump-host menus, OTP prompts, and legacy authentication flows.
Command execution and state handling: Execute commands, inspect state, provide interactive input, and wait for the session to become ready before the next operation.
File transfer across jump hosts: Upload and download files through the existing PTY instead of requiring SFTP. This keeps transfers working through custom command chains and multi-hop connections.
Remote file editing: Apply precise replacements with shellink session edit .
Audit history and supervision: Session input/output is retained as history for audit and replay, while MANUAL mode lets a human take control and supervise or answer prompts.
CLI/TUI, Web UI, REST, and WebSocket: Use the interface that fits the workflow, with stable --json output for automation.
Encrypted profiles and optional single-file binaries: Credentials are protected in SQLite, and Bun builds are available for macOS and Linux.
Why Shellink Is Designed for AI Agents
Shellink is designed around the situations where an AI agent needs reliable access but should not guess or bypass human controls:
Scriptable multi-hop access: Agents can reuse expect or other custom login scripts for several bastion hops instead of implementing each site's authentication flow themselves.
PTY-based continuity: Commands and file transfers use the same PTY path, so an agent can work through jump hosts even when no direct SFTP or SSH endpoint is available.
Human supervision: AUTO mode is for agent actions; MANUAL mode hands the terminal to a person for approval, OTP entry, or intervention before automation continues.
Auditable operations: Session input/output history and daemon logs give operators a record for reviewing what the agent saw and did.
Observable progress: Explicit session states and machine-readable CLI output let an agent decide whether to execute, send input, wait, or ask for human help.
The repository is an npm workspace monorepo with one additional non-workspace directory for agent skills:
flowchart LR
Client["CLI / TUI"] -->|Unix socket, MessagePack RPC| Daemon["Daemon services"]
Daemon --> Sessions["SSH / local PTY sessions"]
Daemon --> Storage["SQLite state and encrypted credentials"]
Daemon --> Gateway["HTTP REST / WebSocket gateway / Web UI"]
Loading
The Unix socket is the primary local interface. HTTP, WebSocket, and the Web UI are compatibility and browser layers. New local automation should use the CLI rather than implementing the socket protocol directly.
Runtime: Node.js >=22.19.0 ; Bun is optional for compiled binaries.
Language: TypeScript with npm workspaces.
Server: Fastify, ssh2 , node-pty , ws , SQLite via better-sqlite3 and Drizzle ORM.
Protocol: MessagePack framing and Zod validation.
Web: Vue 3, Vite, Naive UI, Pinia, and xterm.js.
Testing: Vitest with V8 coverage and Docker-backed end-to-end fixtures.
Using an AI coding agent? Paste this prompt into your agent's chat:
Please help me install Shellink and its skills by following https://raw.githubusercontent.com/jie123108/Shellink/main/AGENTS_INSTALL.md
Prebuilt binaries are published for macOS and Linux ( darwin-arm64 , darwin-x64 , linux-arm64 , linux-x64 ). Windows is not supported; build from source instead (see below).
curl -fsSL https://raw.githubusercontent.com/jie123108/Shellink/main/install.sh -o /tmp/shellink-install.sh
cat /tmp/shellink-install.sh
sh /tmp/shellink-install.sh
Pin a release tag if needed:
curl -fsSL https://raw.githubusercontent.com/jie123108/Shellink/main/install.sh -o /tmp/shellink-install.sh
cat /tmp/shellink-install.sh
sh /tmp/shellink-install.sh --version v0.1.0
For a system-wide install into /usr/local/bin (already on PATH on most systems), run with sudo :
sudo sh /tmp/shellink-install.sh --dir /usr/local/bin
The script downloads the matching asset from GitHub Releases, verifies SHA256SUMS.txt , and installs shellink to $HOME/.local/bin by default. If that directory is not on PATH , the installer appends an export PATH=... line to your shell profile ( ~/.zshrc , ~/.bashrc , etc.). Open a new shell (or re-export PATH ), then verify:
Docker, only if you want to run the Docker-backed end-to-end tests.
Bun, only if you want to build a standalone binary.
git clone https://github.com/jie123108/Shellink.git
cd Shellink
npm install
npm run build
The built CLI is available at ./node_modules/.bin/shellink . To build a standalone executable for the current platform:
npm run build:binary
mkdir -p " $HOME /.local/bin "
cp dist/shellink " $HOME /.local/bin/shellink "
export PATH= " $HOME /.local/bin: $PATH "
Cross-build scripts are available for macos-arm64 , macos-x64 , linux-arm64 , and linux-x64 .
Start the interactive terminal UI. It starts the user-level daemon automatically when necessary:
shellink cli
Check the daemon and print the machine-readable agent reference:
shellink server status --json
shellink agent-doc
Create a command profile and session without putting credentials in shell history:
shellink profile create --input - --json << ' JSON '
{"name":"local-shell","connectType":"command","command":"/bin/sh"}
JSON
shellink profile list --query local --json
shellink session create --profile < profile-id > --json
shellink session exec < session-id > --command ' uname -a ' --json
shellink session close < session-id > --json
For SSH profiles, pass passwords, passphrases, and private keys through --input FILE or stdin. Never put them in command-line arguments.
When running from source, npm run dev starts the server and Vite development UI. Open http://localhost:5173 . A built daemon serves the embedded UI at http://127.0.0.1:7070/shellink/ui/ by default; the exact URL is shown by shellink server run or shellink server status --json .
Custom command profiles are intended for connections that cannot be represented by a single direct SSH login. Set connectType to command and provide any command available on the server host, for example:
{
"name" : " Bastion menu " ,
"connectType" : " command " ,
"command" : " expect /opt/shellink/login-through-bastion.exp "
}
Shellink starts the command in a PTY and exposes the same command, input, history, file-transfer, and editing APIs used by SSH sessions. An expect script can therefore drive several bastion hosts and authentication steps as one Shellink session. The script, its dependencies, and its credentials run on the server host as the daemon user; only execute trusted commands.
Note: Shellink itself does not handle a host's automatic login process. For automated login (multi-hop hops, jump-host menus, OTP prompts, and other interactive authentication flows), use an expect script driven by a command profile. Other tools such as sshpass (password-only logins) or ssh with key-based authentication and ProxyJump / ProxyCommand can also work, depending on your environment.
Client mode
The shellink executable is the client. TUI commands and resource commands ( profile , session , and webhook ) connect to the local daemon through a Unix socket. If the socket is unavailable, the client starts a detached daemon, waits for it to become ready, and then sends the request. Add --json for stable, one-line machine-readable output.
shellink server start # Start in the background
shellink server status --json # Inspect health and active sessions
shellink server logs --lines 100 # Read recent daemon logs
shellink server restart # Restart the daemon
shellink server stop # Stop the daemon
shellink server run # Run in the foreground
server run is suitable for development or a process manager. The daemon always starts its Unix socket. HTTP/WebSocket support is enabled by default and can be disabled with SHELLINK_HTTP_ENABLED=false . HTTP binds to 127.0.0.1 by default; set SHELLINK_HOST=0.0.0.0 only when remote access is intentional.
Local requests from localhost or 127.0.0.1 can omit authentication, including permanent session-record deletes. Remote HTTP and WebSocket requests require Authorization: Bearer <SHELLINK_TOKEN> (or the WebSocket token quer

[truncated]

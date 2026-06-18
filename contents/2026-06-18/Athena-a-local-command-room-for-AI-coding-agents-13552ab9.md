---
source: "https://github.com/luckeyfaraday/Athena"
hn_url: "https://news.ycombinator.com/item?id=48586771"
title: "Athena, a local command room for AI coding agents"
article_title: "GitHub - luckeyfaraday/Athena: Athena is a desktop control surface for AI coding agents with shared project context. Built with Electron, React, and FastAPI, it embeds native terminals for Codex, OpenCode, Claude, and Hermes with Hermes memory integration. Features MCP bridge for cross-platform work\n[truncated]"
author: "luckeyfaraday"
captured_at: "2026-06-18T16:13:48Z"
capture_tool: "hn-digest"
hn_id: 48586771
score: 1
comments: 0
posted_at: "2026-06-18T15:20:33Z"
tags:
  - hacker-news
  - translated
---

# Athena, a local command room for AI coding agents

- HN: [48586771](https://news.ycombinator.com/item?id=48586771)
- Source: [github.com](https://github.com/luckeyfaraday/Athena)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T15:20:33Z

## Translation

タイトル: Athena、AI コーディング エージェントのローカル コマンド ルーム
記事タイトル: GitHub - luckyfaraday/Athena: Athena は、プロジェクト コンテキストを共有する AI コーディング エージェント用のデスクトップ コントロール サーフェスです。 Electron、React、FastAPI で構築されており、Codex、OpenCode、Claude、Hermes メモリ統合を備えた Herme 用のネイティブ端末が組み込まれています。クロスプラットフォーム作業のための MCP ブリッジを搭載
[切り捨てられた]
説明: Athena は、プロジェクト コンテキストを共有する AI コーディング エージェント用のデスクトップ コントロール サーフェスです。 Electron、React、FastAPI で構築されており、Codex、OpenCode、Claude、Hermes メモリ統合を備えた Herme 用のネイティブ端末が組み込まれています。クロスプラットフォームのワークスペース制御のための MCP ブリッジ、組み込み PTY ターを搭載
[切り捨てられた]

記事本文:
GitHub - luckyfaraday/Athena: Athena は、プロジェクト コンテキストを共有する AI コーディング エージェント用のデスクトップ コントロール サーフェスです。 Electron、React、FastAPI で構築されており、Codex、OpenCode、Claude、Hermes メモリ統合を備えた Herme 用のネイティブ端末が組み込まれています。クロスプラットフォームのワークスペース制御のための MCP ブリッジ、組み込み PTY ターミナル、ネイティブ セッション ディスカバリなどを備えています。 · GitHub
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
別のタブまたは別のタブでサインアウトしました

うわー。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ラッキーファラデー
/
アテナ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
354 コミット 354 コミット .github/ workflows .github/ workflows エージェント スキル/ athena-context-workspace エージェント スキル/ athena-context-workspace バックエンド バックエンド cli cli クライアント クライアント docs docs mcp_server mcp_server スクリプト スクリプト テスト テスト .gitignore .gitignore README.md README.md athena-demo.mp4 athena-demo.mp4 athenafocusmode.png athenafocusmode.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント、セッションのリコール、プロジェクトの引き継ぎのためのローカル コマンド ルーム。
クイックスタート
·
特長
·
エルメスを繋ぐ
·
MCPブリッジ
·
テスト
Athena は、共有プロジェクト コンテキストを使用して AI コーディング エージェントを調整するためのローカル デスクトップ ワークスペースです。開発者は、Codex、OpenCode、Claude Code、Athena Code、Hermes、およびシェル セッションを起動するための 1 つの Electron アプリを提供します。ライブ端末出力とネイティブセッション履歴を検査します。プロジェクトの引き継ぎを生成する。そして、次のエージェントが利用できる短期間のリコールコンテキストを保持します。
検索用語: Athena は、AI コーディング エージェント ワークスペース 、マルチエージェント デスクトップ アプリ 、組み込み端末コントロール ルーム 、Hermes MCP ブリッジ 、およびローカル ソフトウェア開発用のセッション リコール マネージャーです。
アテナデモ.mp4
製品ウィジェット
指令室
セッションのリコール
エージェントの範囲
デスクトップランタイム
埋め込み PTY ペイン、シェル フォーカス、ターミナル/チャット モード
Hermes リコール キャッシュ、監査メタデータ、制限付きハンドオフ
コーデックス、オープンコード、クロード コード、アテナ コード、エルメス、シェル
ローカル FastAPI バックエンドを備えた Electron アプリ
LLM の概要
アテナは電子です

+ ローカル AI コーディング エージェント セッションを管理するための FastAPI バックエンドを備えた React デスクトップ アプリケーション。これは、node-pty および xterm.js を介した組み込み PTY ターミナル、Codex、OpenCode、Claude Code、Athena Code、Hermes のネイティブ セッション検出、プロジェクト ローカル リコール キャッシュ、セッション ハンドオフ生成、リコール監査メタデータ、および実行中のデスクトップ ワークスペースを Hermes が制御できるようにする MCP サーバーをサポートします。
AI コーディング ツールは多くの場合、独立した端末として実行され、それぞれが独自のコンテキスト ウィンドウと履歴を持ちます。 Athena は、これらの個別のエージェント セッションを 1 つのローカル コマンド ルームに変えます。
1 つの UI からシェル、Hermes、Codex、OpenCode、Claude、および Athena Code セッションを開始します。
すでにディスクに保存されているネイティブ エージェント セッションを再開します。
ライブターミナルバッファ、ネイティブトランスクリプト、プロバイダメタデータを検査します。
有用なセッション証拠から制限付きハンドオフを生成します。
次の新しいエージェントのためにプロジェクトローカルのリコールに引き継ぎを保存します。
Hermes が MCP ツールを使用してセッションを検査し、リコールを書き込み、表示可能な Athena 端末を生成できるようにします。
Athena は、単なるターミナル エミュレーター、メモリ ストア、または MCP サーバーではありません。これは、セッションファーストの AI 開発作業を繰り返すためのローカル オーケストレーション サーフェスです。
エリア
詳細
アプリの種類
AI コーディング エージェント オーケストレーション用のローカル デスクトップ アプリ
フロントエンド
Electron、React、Vite、TypeScript
端子スタック
node-pty + xterm.js 埋め込み PTY
バックエンド
Electron が開始した FastAPI Python サービス
エージェントサポート
コーデックス、オープンコード、クロード コード、アテナ コード、エルメス、シェル
コンテキストシステム
ヘルメスメモリ、プロジェクトローカルリコール、セッションハンドオフ
MCPのサポート
mcp_server/ は Athena ツールを Hermes に公開します
主なワークフロー
エージェントの起動または再開、セッションの検査、ハンドオフの作成、リコールによる新規開始
コア機能
埋め込みシェル、Hermes、Codex、OpenCode、Claude、および Athena Code ペインを起動します。
Codex/OpenCode/Claude/Athena コードを起動する

並行作業のためのグリッド。
ネイティブ Codex、OpenCode、Claude Code、Athena Code、Hermes セッションを再開します。
コマンド ルームで実行中の組み込み PTY と過去のネイティブ セッションを追跡します。
プロバイダーごとのグループセッション履歴。
ライブ バッファー、ネイティブ トランスクリプト、プロンプト パス、モデル メタデータ、ブランチ メタデータ、プロバイダー セッション ID を検査します。
共有プロジェクトコンテキストとリコール
リコールが見つからない、または古い場合は、エージェントの起動前に Herme のリコールを更新します。
プロジェクトローカルのリコールを .context-workspace/hermes/session-recall.md に書き込みます。
選択したセッションから制限された Athena セッション ハンドオフを生成します。
端末 UI/制御ノイズをハンドオフの証拠からフィルタリングします。
ハンドオフを保存して呼び出し、そのハンドオフから新しい Codex、OpenCode、Claude、または Athena Code エージェントを起動します。
リコール監査メタデータを追跡します: ソース、ソース数、ソース タイトル、バイト サイズ、リフレッシュ時間、起動でリコールが使用されたかどうか。
MCP を通じて、Athena の健全性、メモリ、リコール、ネイティブ セッション、トランスクリプト、ターミナルの生成を公開します。
ヘルメスが電子制御を通じて目に見える Athena ターミナルを生成できるようにします。
Hermes にネイティブの Codex/OpenCode/Claude/Athena Code/Hermes セッションの概要を読み取らせます。
エルメスを長期記憶とより高いレベルの想起決定の所有者として保ちます。
ワークスペース タブはアクティブなプロジェクトを分離します。
Electron が起動し、FastAPI バックエンドを監視します。
[設定] には、バックエンド、Hermes、アダプタ、およびリコール ステータスが表示され、Hermes がインストールされ、MCP ブリッジ接続ヘルパーが表示されます。
Memory Room はプロジェクト メモリを検査し、Hermes メモリ エントリを正確に削除できます。
Review Room は、どのセッション出力を保持する価値があるかを判断することに重点を置いています。
バックエンド/ FastAPI バックエンド、メモリ、ネイティブ セッション、リコール、レガシー実行レジストリ
バックエンド/アダプター/エージェントアダプターの実装
クライアント/ Electron + React デスクトップ クライアント
client/electron/Electron メインプロセス サービスと IPC ハンド

ラーズ
client/src/ React UI とブラウザ側 API ラッパー
docs/ 公開実装および検証ノート
mcp_server/ MCP ブリッジ。Hermes が Athena を制御できるようにします。
スクリプト/ローカル検証およびリコールヘルパー
テスト/バックエンド、MCP、ネイティブ セッション、およびアダプターのテスト
要件
実共有メモリ統合のためのオプションのHermes Agentのインストール
デスクトップ アプリは、すべてのエージェント CLI がインストールされていなくても開くことができます。不足しているアダプターは使用できないように表示され、CLI がインストールされて PATH で使用可能になるまで、関連する起動コマンドがターミナル内で失敗する可能性があります。
git クローン https://github.com/luckeyfaraday/Athena.git
cd Athena/クライアント
npmインストール
npm 実行開発
完全なバックエンド/テスト環境の場合は、以下のセットアップ手順を使用します。
クライアントの依存関係をインストールします。
CDクライアント
npmインストール
リポジトリ ルートからバックエンドの依存関係をインストールします。
python3 -m venv .venv
ソース .venv/bin/activate
pip install -r backend/requirements.txt
テスト用に、pytest がまだ利用できない場合はインストールします。
pip インストール pytest
優先する Python が python3 ではない場合は、次のように設定します。
CONTEXT_WORKSPACE_PYTHON=/absolute/path/to/python をエクスポートします。
Electron アプリは、FastAPI バックエンドを生成するときにこの値を使用します。
Electron TypeScript エントリ ポイントを構築します。
Electron は、空きローカルホスト ポートで FastAPI バックエンドを開始します。
CDクライアント
npm ビルドを実行する
Linux 上で AppImage を構築するには:
CDクライアント
npm 実行距離
以前に構築した Electron アプリを起動するには:
CDクライアント
npmスタート
バックエンドを直接実行する
python3 -m uvicorn backend.app:app --host 127.0.0.1 --port 8000
便利なエンドポイント:
GET /健康
/hermes/ステータスを取得する
GET /hermes/recall/status
POST /エルメス/リコール/リフレッシュ
POST /hermes/recall/write
POST /hermes/recall/mark-used
GET /memory/hermes?q=<クエリ>
GET /memory/recent?limit=10
POST /メモリ/ストア
POST /メモリ/削除
GET /エージェント/アダプター
GET /agents/sessions
GET /ag

ents/sessions/{provider}/{session_id}/transcript
POST /エージェント/スポーン
GET /agents/runs
/agents/runs/{run_id} を取得します
POST /agents/runs/{run_id}/cancel
GET /agents/runs/{run_id}/artifacts/{artifact_name}
テスト
リポジトリ ルートからバックエンド テスト スイートを実行します。
pytest
クライアントのビルド チェックを実行します。
CDクライアント
npm ビルドを実行する
テストでは偽の CLI エージェント フィクスチャを使用するため、ホストされたモデルや外部エージェント ツールを使用せずに実行フローを検証できます。
最初の公開ゲートについては、「
docs/release-0.1.0-checklist.md 。
Athena の主なワークフローは、埋め込み型の対話型エージェント セッションです。 Electron のメイン プロセスは、シェル、Hermes、Codex、OpenCode、Claude、および Athena Code のターミナル ペインを起動します。 React UI は、これらのペインを xterm.js でレンダリングします。
デフォルトでは、新しいエージェントペインは Athena プロジェクトコンテキストなしで開始されます。アテナのみ
メモリ、リコール、およびプロジェクト命令バンドルを作成してアタッチします。
明示的なイマーシブ コンテキスト モードが選択されています。すると、次のようになります。
不変のワークスペーススコープのコンテキストバンドルを作成します。
バンドルを指すコンパクトなブートストラップ プロンプトを書き込みます。
選択した CLI を組み込み PTY で開始します。
ペインをライブ セッションとして追跡し、レビューのために境界のあるターミナル バッファーをキャプチャします。
Athena は、ディスク上にすでに存在するネイティブ プロバイダー セッションも検出するため、以前の Codex、OpenCode、Claude Code、Athena Code、Hermes の作業を [セッション] タブから検査または再開できます。
Athena セッション ハンドオフは、レビュー ルームで選択されたセッションから生成される制限されたマークダウン サマリーです。これらは、新しいエージェントが有益なプロジェクトのコンテキストを失うことなく新たにスタートできるように設計されています。
1 つ以上の有用なライブ セッションまたはネイティブ セッションを選択します。
Athena は使用可能な証拠を抽出し、端末 UI ノイズをフィルタリングします。
生成されたハンドオフ プレビューを確認します。
プロジェクトローカルのリコールへのハンドオフを保存します。
新しい Codex、OpenCode、Clau を開始する

de、またはそのリコールが添付された Athena コード セッション。
ハンドオフでは、完全なトランスクリプトを盲目的にマージすることはありません。使用可能なタスク証拠のないメタデータのみのセッションとターミナル バッファは拒否されるか、明確にマークされます。
バックエンドには、古いワンショット実行レジストリと Codex アダプターがまだ含まれています。このパスは、エージェント生成リクエストを受信し、実行レコードを作成し、 .context-workspace/runs/<run-id>/ に制限されたアーティファクトを書き込み、CLI プロセスを実行し、ステータス/アーティファクト エンドポイントを公開します。
バックエンド実行フローは、互換性とテストのために維持されます。 Athena の現在の製品の方向性は、セッションファーストの組み込み端末とネイティブ セッション ディスカバリです。生成されたコンテキスト アーティファクトはキャッシュ/出力ファイルです。ヘルメスの記憶とプロジェクトローカルの想起は、永続的な共有コンテキストとして残ります。
Electron のメイン プロセスは、 node-pty を通じて組み込み端末を管理します。 React UI は xterm.js を使用してそれらをレンダリングします。
エージェント ペインは、タスク、キュレーション、またはタスクに対してのみ生成された Athena プロンプト パスを受け取ります。
明示的な没入型の起動。クリーン起動にはプロンプト パスが表示されません。
バックエンドは、HermesManager と HermemoryStore を使用して、Hermes ステータスを検索し、メモリの読み取り/書き込みを行います。
GET /memory/hermes?q=<クエリ>
応答はプレーンテキストであるため、CLI エージェントはそれを使用できます。

[切り捨てられた]

## Original Extract

Athena is a desktop control surface for AI coding agents with shared project context. Built with Electron, React, and FastAPI, it embeds native terminals for Codex, OpenCode, Claude, and Hermes with Hermes memory integration. Features MCP bridge for cross-platform workspace control, embedded PTY ter
[truncated]

GitHub - luckeyfaraday/Athena: Athena is a desktop control surface for AI coding agents with shared project context. Built with Electron, React, and FastAPI, it embeds native terminals for Codex, OpenCode, Claude, and Hermes with Hermes memory integration. Features MCP bridge for cross-platform workspace control, embedded PTY terminals, native session discovery, and more. · GitHub
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
luckeyfaraday
/
Athena
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
354 Commits 354 Commits .github/ workflows .github/ workflows agent-skills/ athena-context-workspace agent-skills/ athena-context-workspace backend backend cli cli client client docs docs mcp_server mcp_server scripts scripts tests tests .gitignore .gitignore README.md README.md athena-demo.mp4 athena-demo.mp4 athenafocusmode.png athenafocusmode.png View all files Repository files navigation
Local command room for AI coding agents, session recall, and project handoffs.
Quick Start
·
Features
·
Connecting Hermes
·
MCP Bridge
·
Testing
Athena is a local desktop workspace for orchestrating AI coding agents with shared project context. It gives developers one Electron app for launching Codex, OpenCode, Claude Code, Athena Code, Hermes, and shell sessions; inspecting live terminal output and native session history; generating project handoffs; and keeping short-lived recall context available to the next agent.
In search terms: Athena is an AI coding agent workspace , multi-agent desktop app , embedded terminal control room , Hermes MCP bridge , and session recall manager for local software development.
athena-demo.mp4
Product Widgets
Command Room
Session Recall
Agent Coverage
Desktop Runtime
Embedded PTY panes, shell focus, terminal/chat modes
Hermes recall cache, audit metadata, bounded handoffs
Codex, OpenCode, Claude Code, Athena Code, Hermes, shell
Electron app with local FastAPI backend
LLM Summary
Athena is an Electron + React desktop application with a FastAPI backend for managing local AI coding agent sessions. It supports embedded PTY terminals through node-pty and xterm.js , native session discovery for Codex, OpenCode, Claude Code, Athena Code, and Hermes, project-local recall caches, session handoff generation, recall audit metadata, and an MCP server that lets Hermes control the running desktop workspace.
AI coding tools often run as isolated terminals, each with its own context window and history. Athena turns those separate agent sessions into one local command room:
Start shell, Hermes, Codex, OpenCode, Claude, and Athena Code sessions from one UI.
Resume native agent sessions already stored on disk.
Inspect live terminal buffers, native transcripts, and provider metadata.
Generate bounded handoffs from useful session evidence.
Save handoffs into project-local recall for the next fresh agent.
Let Hermes use MCP tools to inspect sessions, write recall, and spawn visible Athena terminals.
Athena is not only a terminal emulator, memory store, or MCP server. It is a local orchestration surface for repeated, session-first AI development work.
Area
Details
App type
Local desktop app for AI coding agent orchestration
Frontend
Electron, React, Vite, TypeScript
Terminal stack
node-pty + xterm.js embedded PTYs
Backend
FastAPI Python service launched by Electron
Agent support
Codex, OpenCode, Claude Code, Athena Code, Hermes, shell
Context system
Hermes memory, project-local recall, session handoffs
MCP support
mcp_server/ exposes Athena tools to Hermes
Primary workflow
Launch or resume agents, inspect sessions, create handoffs, start fresh with recall
Core Features
Launch embedded shell, Hermes, Codex, OpenCode, Claude, and Athena Code panes.
Launch Codex/OpenCode/Claude/Athena Code grids for parallel work.
Resume native Codex, OpenCode, Claude Code, Athena Code, and Hermes sessions.
Track running embedded PTYs and historical native sessions in the Command Room.
Group session history by provider.
Inspect live buffers, native transcripts, prompt paths, model metadata, branch metadata, and provider session IDs.
Shared Project Context And Recall
Refresh Hermes recall before agent launch when recall is missing or stale.
Write project-local recall to .context-workspace/hermes/session-recall.md .
Generate bounded Athena Session Handoffs from selected sessions.
Filter terminal UI/control noise out of handoff evidence.
Save handoffs to recall and launch a fresh Codex, OpenCode, Claude, or Athena Code agent from that handoff.
Track recall audit metadata: source, source count, source titles, byte size, refresh time, and whether recall was used by a launch.
Expose Athena health, memory, recall, native sessions, transcripts, and terminal spawning through MCP.
Let Hermes spawn visible Athena terminals through Electron control.
Let Hermes read native Codex/OpenCode/Claude/Athena Code/Hermes session summaries.
Keep Hermes as the owner of long-term memory and higher-level recall decisions.
Workspace tabs isolate active projects.
Electron starts and monitors the FastAPI backend.
Settings shows backend, Hermes, adapter, and recall status, installs Hermes, and shows the MCP bridge connect helper.
Memory Room can inspect project memory and delete exact Hermes memory entries.
Review Room focuses on deciding which session output is worth keeping.
backend/ FastAPI backend, memory, native sessions, recall, legacy run registry
backend/adapters/ Agent adapter implementations
client/ Electron + React desktop client
client/electron/ Electron main-process services and IPC handlers
client/src/ React UI and browser-side API wrappers
docs/ Public implementation and verification notes
mcp_server/ MCP bridge so Hermes can control Athena
scripts/ Local verification and recall helpers
tests/ Backend, MCP, native session, and adapter tests
Requirements
Optional Hermes Agent install for real shared memory integration
The desktop app can open without every agent CLI installed. Missing adapters appear as unavailable, and related launch commands may fail inside the terminal until the CLI is installed and available on PATH .
git clone https://github.com/luckeyfaraday/Athena.git
cd Athena/client
npm install
npm run dev
For the full backend/test environment, use the setup steps below.
Install the client dependencies:
cd client
npm install
Install backend dependencies from the repository root:
python3 -m venv .venv
source .venv/bin/activate
pip install -r backend/requirements.txt
For tests, install pytest if it is not already available:
pip install pytest
If your preferred Python is not python3 , set:
export CONTEXT_WORKSPACE_PYTHON=/absolute/path/to/python
The Electron app uses this value when spawning the FastAPI backend.
Builds the Electron TypeScript entry points.
Electron starts the FastAPI backend on a free localhost port.
cd client
npm run build
To build an AppImage on Linux:
cd client
npm run dist
To launch a previously built Electron app:
cd client
npm start
Running The Backend Directly
python3 -m uvicorn backend.app:app --host 127.0.0.1 --port 8000
Useful endpoints:
GET /health
GET /hermes/status
GET /hermes/recall/status
POST /hermes/recall/refresh
POST /hermes/recall/write
POST /hermes/recall/mark-used
GET /memory/hermes?q=<query>
GET /memory/recent?limit=10
POST /memory/store
POST /memory/delete
GET /agents/adapters
GET /agents/sessions
GET /agents/sessions/{provider}/{session_id}/transcript
POST /agents/spawn
GET /agents/runs
GET /agents/runs/{run_id}
POST /agents/runs/{run_id}/cancel
GET /agents/runs/{run_id}/artifacts/{artifact_name}
Testing
Run the backend test suite from the repository root:
pytest
Run the client build checks:
cd client
npm run build
The tests use fake CLI agent fixtures so execution flow can be verified without hosted models or external agent tools.
For the first public release gate, see
docs/release-0.1.0-checklist.md .
Athena's primary workflow is embedded, interactive agent sessions. The Electron main process launches terminal panes for shell, Hermes, Codex, OpenCode, Claude, and Athena Code. The React UI renders those panes with xterm.js .
By default, fresh agent panes start without Athena project context. Athena only
creates and attaches memory, recall, and project-instruction bundles when an
explicit immersive context mode is selected. It then:
Creates an immutable workspace-scoped context bundle.
Writes a compact bootstrap prompt that points at the bundle.
Starts the selected CLI in an embedded PTY.
Tracks the pane as a live session and captures a bounded terminal buffer for review.
Athena also discovers native provider sessions already on disk, so previous Codex, OpenCode, Claude Code, Athena Code, and Hermes work can be inspected or resumed from the Sessions tab.
Athena Session Handoffs are bounded markdown summaries generated from selected sessions in Review Room. They are designed to help a new agent start fresh without losing useful project context.
Select one or more useful live or native sessions.
Athena extracts usable evidence and filters terminal UI noise.
Review the generated handoff preview.
Save the handoff to project-local recall.
Start a fresh Codex, OpenCode, Claude, or Athena Code session with that recall attached.
Handoffs do not blindly merge full transcripts. Metadata-only sessions and terminal buffers with no usable task evidence are rejected or clearly marked.
The backend still includes an older one-shot run registry and Codex adapter. This path receives an agent spawn request, creates a run record, writes bounded artifacts under .context-workspace/runs/<run-id>/ , executes the CLI process, and exposes status/artifact endpoints.
That backend-run flow is maintained for compatibility and tests. Athena's current product direction is session-first embedded terminals plus native session discovery. Generated context artifacts are cache/output files. Hermes memory and project-local recall remain the durable shared context.
The Electron main process manages embedded terminals through node-pty . The React UI renders them with xterm.js .
Agent panes receive a generated Athena prompt path only for task, curated, or
explicit immersive launches. Clean launches receive no prompt path.
The backend uses HermesManager and HermesMemoryStore to find Hermes status and read/write memory.
GET /memory/hermes?q=<query>
The response is plain text so CLI agents can consume it ea

[truncated]

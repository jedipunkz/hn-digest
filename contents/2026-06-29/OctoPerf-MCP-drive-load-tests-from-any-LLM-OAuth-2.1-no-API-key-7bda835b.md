---
source: "https://api.octoperf.com/doc/mcp/"
hn_url: "https://news.ycombinator.com/item?id=48716592"
title: "OctoPerf MCP – drive load tests from any LLM (OAuth 2.1, no API key)"
article_title: "MCP Server - Documentation - OctoPerf"
author: "Jellly"
captured_at: "2026-06-29T09:32:52Z"
capture_tool: "hn-digest"
hn_id: 48716592
score: 1
comments: 0
posted_at: "2026-06-29T08:53:11Z"
tags:
  - hacker-news
  - translated
---

# OctoPerf MCP – drive load tests from any LLM (OAuth 2.1, no API key)

- HN: [48716592](https://news.ycombinator.com/item?id=48716592)
- Source: [api.octoperf.com](https://api.octoperf.com/doc/mcp/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T08:53:11Z

## Translation

タイトル: OctoPerf MCP – 任意の LLM からのドライブ負荷テスト (OAuth 2.1、API キーなし)
記事のタイトル: MCP サーバー - ドキュメント - OctoPerf
説明: OctoPerf 負荷テスト - MCP サーバーの概要。モデル コンテキスト プロトコルを通じて任意の AI エージェント (Claude.ai、Claude Code、Cursor など) から OctoPerf を駆動します。

記事本文:
MCP サーバー - ドキュメント - OctoPerf
コンテンツにスキップ
ドキュメント - OctoPerf
MCPサーバー
検索を初期化しています
はじめに
はじめに
はじめに
OctoPerfとは何ですか?
管理
管理
概要
MCPサーバー
MCPサーバー
目次
何をするのか
パブリック HTTP ダウンロード
JSONスキーマ
クライアントを接続する
任意の MCP クライアント
Claude.ai (カスタム コネクタ)
コネクタの追加
クロード・コード
手動セットアップ (プラグインなし)
チュートリアル
チュートリアル
オンプレミスのインフラ
オンプレミスのインフラ
Kubernetes 上にデプロイする
デザイン
デザイン
仮想ユーザーページ
仮想ユーザーの作成
仮想ユーザーの作成
ウェブサイトまたは休憩
ウェブサイトまたは休憩
マニュアル作成
HARレコーダー
HARレコーダー
HAR のインポート
仮想ユーザーの編集
仮想ユーザーの編集
デザインページ
アクションの種類
アクションの種類
HTTPアクション
ロジックアクション
ロジックアクション
コンテナアクション
ポストプロセッサ
ポストプロセッサ
正規表現変数エクストラクター
JSR223アクション
JSR223アクション
JSR223アクション
劇作家のアクション
劇作家のアクション
劇作家の構成
構成
構成
サーバー
変数
変数
変数パネル
ランタイム
ランタイム
ランタイムページ
シナリオを編集する
シナリオを編集する
シンプルなシナリオ
ユーザープロファイルの編集
ユーザープロファイルの編集
仮想ユーザー
統合と自動化
統合と自動化
ライブレポート
ライブレポート
はじめに
アプリケーションのパフォーマンス管理
モニタリング
モニタリング
仕組みは?
接続を作成する
接続を作成する
アパッチHTTP
分析
分析
分析ページ
ベンチレポートを編集する
ベンチレポートを編集する
ベンチレポートページ
報告事項
報告項目
面グラフ
エクスポート
エクスポート
ベンチレポートをエクスポートする
オンプレミスエージェント
オンプレミスエージェント
はじめに
プロバイダーの種類
プロバイダーの種類
オンプレミス
オンプレミスのインフラ
オンプレミスのインフラ
セットアップ
チュートリアル
チュートリアル
オフラインインストール
その他
その他
リリースノート
パブ

lic HTTP ダウンロード
JSONスキーマ
クライアントを接続する
任意の MCP クライアント
Claude.ai (カスタム コネクタ)
コネクタの追加
クロード・コード
手動セットアップ (プラグインなし)
OctoPerf MCP サーバーを使用すると、モデル コンテキスト プロトコル対応 AI エージェントが OctoPerf アカウントを操作し、仮想ユーザーのインポート、編集、検証、負荷シナリオの実行、メトリクスのリードバックをすべて、エージェントを離れることなくチャットまたは IDE から行うことができます。
これは https://api.octoperf.com/mcp で公開されています。
MCP サーバーは、LLM エージェントと OctoPerf REST API 間のシン ステートレス ブリッジです。すべてのツール呼び出しはユーザーとして認証されます。サービス ID や API キーはなく、通常の OctoPerf API 呼び出しに解決されます。ユーザーとして Web UI から実行できることはすべて、ユーザーが承認すると、エージェントがユーザーに代わって実行できるようになります。
具体的には、サーバーは次の領域にグループ化された約 100 のツールを公開します。
検出 : ワークスペース、プロジェクト、シナリオ、仮想ユーザー、Docker プロバイダーをリストします。
仮想ユーザーのインポート: HAR、JMX、Postman、Playwright、WebDriver、URL リスト、または以前にエクスポートしたアーカイブから録音を取り込みます。
仮想ユーザー編集: 名前変更、再タグ付け、JSON Patch経由でアクションツリーにパッチを適用、変数、HTTPサーバー、相関ルール、プロジェクトファイルを管理
検証 : 機能チェックを開始し、アクションごとの検証インデックスをリードバックし、失敗したリクエスト/レスポンスをドリルダウンし、HTTP ボディを取得します。
ベンチ実行: シナリオの開始、進行状況の監視、実行中の実行の停止、Docker 起動ログの読み取り
分析: ベンチ レポートのリストと読み取り、チャート/テーブル/サマリー/トップ ウィジェットの背後にある値シリーズの取得、エラー行の展開、トレンド レポートの実行
各ツールは OctoPerf UI 内の一致するページへのディープリンクを返すため、エージェントは結果を要約するときにクリック可能な URL を渡すことができます。
MCP サーバーは、小規模なパブリック、unau のセットも提供します。

その後、マークダウン ファイルが /mcp/public/** の下に保存されるため、エージェント ホストは IdP 資格情報を取得する前に OctoPerf プロジェクトをブートストラップできます。
これらのファイルは、OctoPerf Claude Plugins リポジトリでも入手できます。
patch_* ペイロードの構築に使用されるエンティティ スキーマ (JSON スキーマ 2020-12、多態性サブタイプごとに 1 つの oneOf ブランチ) は、 /mcp/public/schema/ の下でプレーン JSON として同じ方法で提供されます。これらは、ツールを呼び出すが MCP リソースを読み取らないクライアントの HTTP フォールバックです。本体は、一致する octoperf://schema/* リソースと同一です。
このセクションでは、一般的なモデル コンテキスト プロトコル クライアントを OctoPerf MCP サーバーに接続する方法について説明します。
サーバーは https://api.octoperf.com/mcp でストリーミング可能な HTTP を話し、OAuth 2.1 で認証します (動的クライアント登録 - API キーや手動クライアント設定は不要)。 MCP 対応クライアントはその URL だけで接続できます。クライアントがリモート / HTTP MCP サーバーを想定している場所に URL を貼り付け、HTTP トランスポートを選択し、初回使用時にブラウザのログインを完了します。
以下のセクションでは、最も一般的なクライアントについて説明します。ここにリストされていないもの - Windsurf 、 Cline 、 Continue.dev 、 Zed 、 Goose 、 JetBrains AI Assistant 、 LibreChat など - は同じパターンに従います: URL を追加し、HTTP トランスポートを選択し、ブラウザで承認します。クライアントがプロジェクト コンテキスト ファイルをサポートしている場合は、エージェント ガイドでそれを指定します。
Claude.ai (カスタムコネクタ) ¶
正規の Claude.ai フローについては、Anthropic の「リモート MCP を使用したカスタム コネクタの使用を開始する」を参照してください。
claude.ai (Pro/Max) で [カスタマイズ] → [コネクタ] を開きます。チーム/エンタープライズ プランの場合、所有者は組織設定 → コネクタ から同じことを行います。
[+] ボタンをクリックし、 [カスタム コネクタを追加] をクリックします。
https://api.octoperf.com/mcp をコネクタ URL として貼り付け、「追加」をクリックします。
クロードはブラウザ タブを開き、OctoPerf を表示します。 USUでログインしてください

l 資格情報。
コネクタはリストに OctoPerf MCP として表示され、すべての新しい会話で使用できます。
Claude.ai Web: ファイルのアップロード/ダウンロードにはドメイン許可リスト (組織管理者) が必要です。
いくつかの OctoPerf ツール ( Upload_project_file 、 download_project_file 、 download_bench_result_file 、 export_bench_report_pdf の PDF レッグ、および Playwright トレース / HAR / JTL プル) は、OctoPerf REST API を指す、有効期間の短い署名付き URL を返します。
Claude のサンドボックスはその URL をフェッチしてバイトを取り込み、出力は組織レベルの管理設定 (Claude.ai → 組織設定 (管理者のみ) → 機能 → ドメイン許可リスト → 追加の許可ドメイン) によってゲートされます。 octoperf.com (またはオンプレミス展開をカバーする場合は *.octoperf.com) を組織全体に 1 回追加します。ドロップダウンは [なし] のままにすることができます (プリセット リストは必要ありません)。手動入力だけで十分です。
このリストを編集できるのはワークスペース管理者のみです。組織の管理者以外のメンバーはリストを継承します。組織にこれを切り替える管理者スロットがない場合 (または、設定が自分のアカウントに存在する個人用の Claude.ai アカウントを使用している場合)、ファイル ツールは Claude.ai Web およびデスクトップで失敗し続けます。
チャットでクロードに「OctoPerf ワークスペースをリストして」などと質問すると、一致する MCP ツールが呼び出され、結果が表示されます。
また、クロードにエージェント ガイドの内容を読むように指示する必要があります。これはプロジェクトのコンテキストを AI に通知するナレッジ ベースとして機能し、OctoPerf プロジェクトの標準に準拠していることを確認します。
ファイル使用量が多く (CSV の頻繁なアップロード、ベンチ結果ファイルのダウンロード、PDF のエクスポート、Playwright トレースの取得)、ホワイトリストにアクセスできない場合には、Claude Code がより良いクライアントです。ローカル シェル コマンド経由でフェッチ自体が実行されるため、サンドボックス ホワイトリストは適用されず、アップロード/ダウンロードはすぐに機能します。 C

ursor、Continue.dev、および MCP Inspector も影響を受けません。
推奨されるセットアップでは、OctoPerf/octoperf-claude-plugins マーケットプレイスを通じて配布される公式 OctoPerf プラグインを使用します。 MCP サーバーの登録、エージェント ガイドのインストール、および 8 つのワークフロー ショートカット (自動相関、検証トリアージ、シナリオ診断、ベンチ レポートの読み取り、PDF エクスポート、リアル ブラウザー プローブ、スケジューリング、非同期ポーリング) を 1 つのステップで追加します。
/プラグイン マーケットプレイス OctoPerf/octoperf-claude-plugins を追加
/プラグインのインストール octoperf@octoperf
最初のツール呼び出しにより、ブラウザー タブが開き、OctoPerf が表示されます。通常の認証情報を使用してログインします。ログインにより、Claude Code が接続されたエージェントとして承認されます。その後、今後のすべてのツール呼び出しでは、キャッシュされたトークンが透過的に再利用されます。
インストールしたら、Claude Code に「OctoPerf ワークスペースをリストして」、「前回のシナリオの実行が失敗したのはなぜですか?」などの質問をします。 、または「この仮想ユーザーを自動相関させる」と、一致するワークフローが自動的に実行されます。
MCP サーバーのみを登録する場合 (ワークフロー ショートカットやエージェント ガイドは不要)、完全なリファレンスについては、Anthropic の Connect Claude Code to tools via MCP を参照してから、次を実行します。
クロード mcp add --transport http octoperf https://api.octoperf.com/mcp
すべてのフラグ ( --transport 、 --scope など) はサーバー名の前に指定する必要があります。後でサーバーを削除するには:
クロード・MCPはオクトパーフを削除します
クロード セッションを開き、「/mcp」と入力して、新しく追加された OctoPerf MCP サーバーに接続します。
エージェント ガイドは手動でダウンロードできます。
ワークフロー ショートカットを https://api.octoperf.com/mcp/public/skills/<name>.md から .claude/skills/octoperf-<name>/SKILL.md (グローバルに有効にする場合は ~/.claude/skills/...) にダウンロードして、手動でインストールすることもできます。上記のプラグインは同じことを行います。
正規のクロードについては、Anthropic の「MCP 経由でクロード デスクトップをツールに接続する」を参照してください。

デスクトップの流れ。 claude_desktop_config.json を開き ([設定] → [開発者] → [構成の編集] を使用)、 mcpServers の下にエントリを追加します。
{
"mcpサーバー": {
"オクトパーフ" : {
"url" : "https://api.octoperf.com/mcp"
}
}
}
クロードデスクトップを再起動します。最初のツール呼び出しでは、OAuth ログイン用のブラウザ タブが開きます。ベアラー トークンは後でキャッシュされます。
Codex の完全なリファレンスについては、OpenAI の Model Context Protocol ガイドを参照してから、サーバーを登録します。
codex mcp add octoperf --url https://api.octoperf.com/mcp
次に認証します。
コーデックスmcpログインoctoperf
Codex はブラウザ タブを開き、OctoPerf を表示します。通常の認証情報を使用してログインします。ログインにより、Codex が接続されたエージェントとして承認されます。トークンは後でキャッシュされるため、今後のツール呼び出しでは透過的に再利用されます。後でサーバーを削除するには:
コーデックスmcpはoctoperfを削除します
Codex はプロジェクトのルートにある AGENTS.md ファイルをネイティブに読み取るため、そこからエージェント ガイドをダウンロードします。
カール -o AGENTS.md https://api.octoperf.com/mcp/public/AGENTS.md
セッション中に、Codex に「OctoPerf ワークスペースをリストしてください」などと要求すると、一致する MCP ツールが呼び出され、結果がレンダリングされます。
ChatGPT は、カスタム コネクタ (リモート MCP) として OctoPerf MCP サーバーに接続できます。設定の可用性と正確な場所は、プランと OpenAI のロールアウト (カスタム MCP コネクタ/開発者モードはまだ進化中) によって異なるため、正規のフローについては、ChatGPT ドキュメントの OpenAI の現在のコネクタに従ってください。
ChatGPT で、[設定] → [コネクタ] を開きます (カスタム MCP サーバーを追加するには、開発者モードを有効にする必要がある場合があります)。
URL https://api.octoperf.com/mcp を使用してコネクタを追加します。
プロンプトが表示されたら、ブラウザで承認します。トークンは後でキャッシュされます。
Claude.ai と同様、ChatGPT はサンドボックスでツールを実行します。署名付き *.octoperf.com URL を返すファイル ツール ( Upload_project_file 、 download_project_file 、 do

wnload_bench_result_file 、export_bench_report_pdf の PDF レッグ、Playwright トレース / HAR / JTL プル）は、ホストの出力ポリシーによってブロックされる可能性があります。ファイルを大量に使用する場合は、フェッチ自体を実行する Codex または別のローカル クライアントを推奨します。
完全なリファレンスについては、Gemini CLI を使用した Google の MCP サーバーを参照してから、サーバーを登録します。
gemini mcp add --transport http octoperf https://api.octoperf.com/mcp
Gemini セッションを開いて認証します。
/mcp 認証オクトパーフ
Gemini はブラウザ タブを開き、OctoPerf を表示します。通常の資格情報を使用してログインします。ログインにより、Gemini が接続されたエージェントとして認証され、トークンはその後キャッシュされます。後でサーバーを削除するには:
Gemini MCP はオクトパーフを削除します
Gemini CLI はプロジェクトのルートにある GEMINI.md コンテキスト ファイルを読み取るため、そこからエージェント ガイドをダウンロードします。
カール -o GEMINI.md https://api.octoperf.com/mcp/public/AGENTS.md
セッション中に、Gemini に「OctoPerf ワークスペースをリストして」などと要求すると、一致する MCP ツールを呼び出して結果をレンダリングします。
MCP サーバーは、エージェント モード (VS Code 1.101 以降) の Copilot で使用できます。完全なリファレンスについては、「VS Code での MCP サーバーの追加と管理」を参照してください。
ワークスペースで .vscode/mcp.json を作成 (または編集) し、サーバーを追加します (または、MCP: Add Server fr を実行します)

[切り捨てられた]

## Original Extract

OctoPerf load testing - MCP server overview. Drive OctoPerf from any AI agent (Claude.ai, Claude Code, Cursor, ...) through the Model Context Protocol.

MCP Server - Documentation - OctoPerf
Skip to content
Documentation - OctoPerf
MCP Server
Initializing search
Getting Started
Getting Started
Getting Started
What is OctoPerf?
Administration
Administration
Overview
MCP Server
MCP Server
Table of contents
What it does
Public HTTP downloads
JSON schemas
Connect a client
Any MCP client
Claude.ai (Custom Connector)
Adding the connector
Claude Code
Manual setup (without the plugin)
Tutorials
Tutorials
On-Premise Infra
On-Premise Infra
Deploy on Kubernetes
Design
Design
The Virtual Users Page
Create a Virtual User
Create a Virtual User
Website or Rest
Website or Rest
Manual creation
HAR Recorder
HAR Recorder
Import HAR
Edit a Virtual User
Edit a Virtual User
The design page
Action Types
Action Types
HTTP Actions
Logic Actions
Logic Actions
Container Action
Post Processors
Post Processors
Regexp Variable Extractor
JSR223 Actions
JSR223 Actions
JSR223 Actions
Playwright Actions
Playwright Actions
Playwright Configuration
Configuration
Configuration
Servers
Variables
Variables
The variables panel
Runtime
Runtime
The Runtime Page
Edit a Scenario
Edit a Scenario
Simple Scenario
Edit User Profile
Edit User Profile
Virtual user
Integrations & Automation
Integrations & Automation
Live Reporting
Live Reporting
Introduction
Application Performance Management
Monitoring
Monitoring
How it works?
Create a Connection
Create a Connection
Apache Httpd
Analysis
Analysis
The Analysis Page
Edit a Bench Report
Edit a Bench Report
The Bench Report Page
Report Items
Report Items
Area Chart
Export
Export
Export a Bench Report
On-Premise Agent
On-Premise Agent
Introduction
Provider type
Provider type
On-Premise
On-Premise Infra
On-Premise Infra
Setup
Tutorials
Tutorials
Offline installation
Miscellaneous
Miscellaneous
Release Notes
Public HTTP downloads
JSON schemas
Connect a client
Any MCP client
Claude.ai (Custom Connector)
Adding the connector
Claude Code
Manual setup (without the plugin)
The OctoPerf MCP server lets any Model Context Protocol - aware AI agent drive your OctoPerf account - import Virtual Users, edit them, validate, run load scenarios, and read back metrics - all from a chat or IDE, without leaving the agent.
It is exposed at https://api.octoperf.com/mcp .
The MCP server is a thin, stateless bridge between an LLM agent and the OctoPerf REST API. Every tool call is authenticated as you : there are no service identities and no API keys - and resolves to a regular OctoPerf API call. Anything you can do from the web UI as your user, the agent can do on your behalf when you authorize it.
Concretely, the server exposes around 100 tools grouped into the following areas:
Discovery : list workspaces, projects, scenarios, Virtual Users, Docker providers
Virtual User import : bring in recordings from HAR, JMX, Postman, Playwright, WebDriver, URL lists, or previously exported archives
Virtual User edit : rename, re-tag, patch the action tree via JSON Patch, manage variables, HTTP servers, correlation rules and project files
Validation : kick off a functional check, read back the per-action validation index, drill into failing requests / responses, fetch HTTP bodies
Bench run : start a scenario, monitor progress, stop a running run, read Docker launch logs
Analysis : list and read bench reports, fetch the value series behind any chart / table / summary / top widget, expand error rows, run trend reports
Each tool returns a deep-link to the matching page in the OctoPerf UI so the agent can hand you a clickable URL when it summarizes a result.
The MCP server also serves a small set of public, unauthenticated markdown files under /mcp/public/** , so any agent host can bootstrap an OctoPerf project before it has IdP credentials:
These files are also available in OctoPerf Claude Plugins repository
The entity schemas used to build patch_* payloads (JSON Schema 2020-12, one oneOf branch per polymorphic subtype) are served the same way, as plain JSON under /mcp/public/schema/ . These are the HTTP fallback for clients that call tools but don't read MCP resources — the body is identical to the matching octoperf://schema/* resource:
This section walks through connecting popular Model Context Protocol clients to the OctoPerf MCP server.
The server speaks Streamable HTTP at https://api.octoperf.com/mcp and authenticates with OAuth 2.1 (Dynamic Client Registration - no API key, no manual client setup). Any MCP-aware client can connect with just that URL: paste it wherever the client expects a remote / HTTP MCP server, pick the HTTP transport, and complete the browser login on first use.
The sections below cover the most common clients. Anything not listed here - Windsurf , Cline , Continue.dev , Zed , Goose , JetBrains AI Assistant , LibreChat , ... - follows the same pattern: add the URL, choose the HTTP transport, authorize in the browser. Where a client supports a project context file, point it at the agent guide .
Claude.ai (Custom Connector) ¶
See Anthropic's Get started with custom connectors using remote MCP for the canonical Claude.ai flow.
Open Customize → Connectors in claude.ai (Pro/Max). On Team/Enterprise plans, an Owner does the same from Organization settings → Connectors .
Click the + button, then Add custom connector .
Paste https://api.octoperf.com/mcp as the connector URL and click Add .
Claude opens a browser tab to OctoPerf. Log in with your usual credentials.
The connector appears in the list as OctoPerf MCP and is available in every new conversation.
Claude.ai web: file upload / download requires a domain allowlist (org admin).
Several OctoPerf tools ( upload_project_file , download_project_file , download_bench_result_file , the PDF leg of export_bench_report_pdf , plus any Playwright trace / HAR / JTL pull) return a short-lived presigned URL pointing at the OctoPerf REST API.
Claude's sandbox fetches that URL to ingest the bytes, and egress is gated by an org-level admin setting : Claude.ai → Organization Settings (Admin only) → Capabilities → Domain allowlist → Additional allowed domains . Add octoperf.com (or *.octoperf.com to also cover an on-prem deployment) once for the whole org; the dropdown can stay on None (no preset list needed), the manual entry alone is enough.
Only a workspace admin can edit this list - non-admin members of the org inherit it. If your org doesn't have an admin slot to flip this (or you're on a personal Claude.ai account where the setting lives under your own account), the file tools will keep failing on Claude.ai web and Desktop.
In a chat, ask Claude something like "list my OctoPerf workspaces" and it will call the matching MCP tool and render the result.
You should also tell Claude to read the content of the agent guide . It acts as a knowledge base that informs the AI about your project context, ensuring that follows OctoPerf project's standards.
If your usage is file-heavy (frequent CSV uploads, downloading bench result files, exporting PDFs, pulling Playwright traces) and the allowlist is out of reach, Claude Code is the better client: it runs the fetches itself via local shell commands, so no sandbox allowlist applies and uploads / downloads work out of the box. Cursor, Continue.dev and the MCP Inspector are also unaffected.
The recommended setup uses the official OctoPerf plugin distributed through the OctoPerf/octoperf-claude-plugins marketplace. It registers the MCP server, installs the agent guide, and adds eight workflow shortcuts (auto-correlation, validation triage, scenario diagnosis, bench-report reading, PDF export, real-browser probe, scheduling, async polling) in one step.
/plugin marketplace add OctoPerf/octoperf-claude-plugins
/plugin install octoperf@octoperf
The first tool call opens a browser tab to OctoPerf. Log in with your usual credentials - the login authorizes Claude Code as a connected agent. After that, every future tool call reuses the cached token transparently.
Once installed, ask Claude Code something like "list my OctoPerf workspaces" , "why is my last scenario run failing?" , or "auto-correlate this Virtual User" and the matching workflow runs automatically.
If you prefer registering only the MCP server (no workflow shortcuts, no agent guide), see Anthropic's Connect Claude Code to tools via MCP for the full reference, then run:
claude mcp add --transport http octoperf https://api.octoperf.com/mcp
All flags ( --transport , --scope , ...) must come before the server name. To remove the server later:
claude mcp remove octoperf
Open a Claude session and type /mcp to connect to the newly added OctoPerf MCP Server.
You can download the agent guide manually.
You can also install the workflow shortcuts by hand by downloading them from https://api.octoperf.com/mcp/public/skills/<name>.md into .claude/skills/octoperf-<name>/SKILL.md (or ~/.claude/skills/... to enable them globally). The plugin above does the same thing for you.
See Anthropic's Connect Claude Desktop to tools via MCP for the canonical Claude Desktop flow. Open claude_desktop_config.json (via Settings → Developer → Edit Config ) and add an entry under mcpServers :
{
"mcpServers" : {
"octoperf" : {
"url" : "https://api.octoperf.com/mcp"
}
}
}
Restart Claude Desktop. The first tool call opens a browser tab for OAuth login; the bearer token is cached afterwards.
See OpenAI's Model Context Protocol guide for the full Codex reference, then register the server:
codex mcp add octoperf --url https://api.octoperf.com/mcp
Then authenticate:
codex mcp login octoperf
Codex opens a browser tab to OctoPerf. Log in with your usual credentials - the login authorizes Codex as a connected agent. The token is cached afterwards, so every future tool call reuses it transparently. To remove the server later:
codex mcp remove octoperf
Codex reads an AGENTS.md file at the root of your project natively, so download the agent guide there:
curl -o AGENTS.md https://api.octoperf.com/mcp/public/AGENTS.md
In a session, ask Codex something like "list my OctoPerf workspaces" and it will call the matching MCP tool and render the result.
ChatGPT can reach the OctoPerf MCP server as a custom connector (remote MCP). Availability and the exact location of the setting depend on your plan and OpenAI's rollout (custom MCP connectors / Developer mode are still evolving), so follow OpenAI's current Connectors in ChatGPT documentation for the canonical flow.
In ChatGPT, open Settings → Connectors (you may need to enable Developer mode to add a custom MCP server).
Add a connector with the URL https://api.octoperf.com/mcp .
Authorize in the browser when prompted; the token is cached afterwards.
Like Claude.ai, ChatGPT runs tools in a sandbox. File tools that return a presigned *.octoperf.com URL ( upload_project_file , download_project_file , download_bench_result_file , the PDF leg of export_bench_report_pdf , Playwright trace / HAR / JTL pulls) may be blocked by the host's egress policy. For file-heavy use, prefer Codex or another local client, which run the fetches themselves.
See Google's MCP servers with the Gemini CLI for the full reference, then register the server:
gemini mcp add --transport http octoperf https://api.octoperf.com/mcp
Open a Gemini session and authenticate:
/mcp auth octoperf
Gemini opens a browser tab to OctoPerf. Log in with your usual credentials - the login authorizes Gemini as a connected agent and the token is cached afterwards. To remove the server later:
gemini mcp remove octoperf
Gemini CLI reads a GEMINI.md context file at the root of your project, so download the agent guide there:
curl -o GEMINI.md https://api.octoperf.com/mcp/public/AGENTS.md
In a session, ask Gemini something like "list my OctoPerf workspaces" and it will call the matching MCP tool and render the result.
MCP servers are available to Copilot in agent mode (VS Code 1.101 or later). See Add and manage MCP servers in VS Code for the full reference.
Create (or edit) .vscode/mcp.json in your workspace and add the server (alternatively, run MCP: Add Server fr

[truncated]

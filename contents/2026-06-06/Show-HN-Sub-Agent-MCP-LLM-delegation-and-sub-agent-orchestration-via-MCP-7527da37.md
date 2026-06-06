---
source: "https://github.com/stormaref/Sub-Agent-MCP"
hn_url: "https://news.ycombinator.com/item?id=48423744"
title: "Show HN: Sub-Agent MCP: LLM delegation and sub-agent orchestration via MCP"
article_title: "GitHub - stormaref/Sub-Agent-MCP: MCP for spawning sub-agents everywhere · GitHub"
author: "avestura"
captured_at: "2026-06-06T12:34:13Z"
capture_tool: "hn-digest"
hn_id: 48423744
score: 1
comments: 0
posted_at: "2026-06-06T11:10:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sub-Agent MCP: LLM delegation and sub-agent orchestration via MCP

- HN: [48423744](https://news.ycombinator.com/item?id=48423744)
- Source: [github.com](https://github.com/stormaref/Sub-Agent-MCP)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T11:10:02Z

## Translation

タイトル: HN を表示: サブエージェント MCP: MCP を介した LLM 委任とサブエージェント オーケストレーション
記事のタイトル: GitHub - stormaref/Sub-Agent-MCP: あらゆる場所にサブエージェントを生成するための MCP · GitHub
説明: あらゆる場所にサブエージェントを生成するための MCP。 GitHub でアカウントを作成して、stormaref/Sub-Agent-MCP の開発に貢献してください。

記事本文:
GitHub - stormaref/Sub-Agent-MCP: あらゆる場所にサブエージェントを生成するための MCP · GitHub
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
ストーマレフ
/
サブエージェント-MCP
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
16 コミット 16 コミット .github/ workflows .github/ workflows .vscode .vscode config config docker/丘陵地帯/モック_mcp docker/モック_mcpスクリプト スクリプト src/sub_agent_mcp src/sub_agent_mcp テスト テスト .gitignore .gitignore Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM 委任およびサブエージェント オーケストレーション用の実稼働対応 Python MCP サーバー。親 LLM (たとえば、Cursor のエージェント) はこのサーバーに接続し、YAML で定義された各エージェントにちなんで名付けられたツール (たとえば、 Researcher ) を呼び出すことによって作業を委任します。
各サブエージェントは、独自の LLM、システム プロンプト、およびオプションのダウンストリーム MCP ツール サーバーを使用して YAML で定義されます。
サブエージェント MCP は、親 LLM と 1 つ以上の特殊なサブエージェントの間に位置します。
親は、 /mcp で Streamable HTTP 経由でこのサーバーに接続します。
起動時に、agents.yaml 内の各エージェントは、その id で指定された MCP ツールとして登録されます。
親はプロンプトを使用してそのツールを呼び出します。サーバーはサブエージェントを実行し、最終応答を返します。
各サブエージェントは、独自の OpenAI 互換 LLM、システム プロンプト、および他の MCP サーバー (ファイル システム、検索、独自のツールなど) へのオプションの接続を備えた LangChain エージェントです。ツールへのアクセスは、ホワイトリストを使用してエージェントごとに制限できます。
シーケンス図
参加者の親 (Parent_LLM として)
Sub_Agent_MCP としての参加者 SAMCP
参加者エージェント (LangChain_sub_agent)
OpenAI_compatibility_LLM としての参加者 LLM
Downstream_MCP_servers としての参加者ツール
親->>SAMCP: 研究者(プロンプト)
SAMCP->>エージェント: ビルド エージェント + ダウンストリーム ツール
エージェント->>LLM: 推論ループ
エージェント->>ツール: MCP ツール呼び出し
ツール-->>エージェント: ツールの結果
LLM-->>エージェント: 最終回答
エージェント-->>SAMCP:

応答テキスト
SAMCP-->>親: {"応答": "..."}
読み込み中
これは、ワークスペース内のすべてのツールを 1 人のエージェントに与えるのとは異なります。親は軽量のままで、役割は明示的なままで、各サブエージェントには、MCP サーバーとそのために構成されたツールのみが表示されます。
コンテキストの肥大化を伴わない委任 — 親はエージェント ツールを直接呼び出します。独自のコンテキストですべてのダウンストリーム ツール スキーマを必要とするわけではありません。
ロールごとの構成 — 異なる ID は、異なるモデル、プロンプト、MCP サーバー、およびツール許可リストを使用できます。
実稼働指向 — Pydantic で検証された YAML、構造化ログ、Docker ヘルスチェック、CI、およびリリース タグの GHCR イメージ。
OpenAI 互換プロバイダー — OpenAI、Azure、Ollama、LM Studio、または互換性のある API で llm.base_uri を指定します。
ストリーミング可能な HTTP のみ ( streamable-http )。標準入出力または従来の SSE はありません。
厳密な Pydantic 検証を備えた YAML エージェント定義。
環境置換: ${VAR} および ${VAR:-default} 。
OpenAI 互換のチャット モデルを備えた LangChain 1.x create_agent。
langchain-mcp-adapters を介したエージェントごとの MCP 接続 (オプションの server.tool ホワイトリスト)。
エージェント ツールの説明では API キーが公開されることはありません。
構造化ログ ( structlog )。
ヘルスチェック付きの Docker イメージ。 GitHub アクション CI。
バージョンタグ ( v0.x.y ) で GHCR に公開されたコンテナーイメージ。
このサーバーによって公開される MCP ツール
Agents.yaml 内の各エージェントは、その ID によって名前が付けられたツールとして登録されます (たとえば、 Researcher )。各ツールは単一のプロンプト引数を受け入れ、そのサブエージェントを実行します。
要件
注意事項
Python 3.10+
CI は 3.12 を使用します。 pyproject.toml の require-python >= 3.10
LLM プロバイダーの API キー
デフォルトの例では OPENAI_API_KEY を使用します
uv (推奨) または pip
CI インストール パスと一致します
Docker + Compose (オプション)
最初の実行に推奨されます。モック MCP ツール サーバーが含まれています
クイックスタート
1 つの道を選択してください。 Docker コンポ

e は、完全なデモ スタック (サーバー + モック ツール サーバー) を実行する最速の方法です。
パス A — Docker Compose (推奨)
エクスポート OPENAI_API_KEY=sk-...
docker 構成 --build
サービス
港
エンドポイント
サブエージェントMCP
8000
http://localhost:8000/mcp
モックファイルシステム MCP
8001
http://localhost:8001/mcp
モック検索 MCP
8002
http://localhost:8002/mcp
バンドルされている config/agents.yaml は、Compose 内で正しく解決される Docker ネットワーク ホスト名 ( filesystem-mcp 、 search-mcp ) を使用します。
UV同期 --dev
# または: pip install -e ".[dev]"
cp config/agents.example.yaml config/agents.yaml # config/agents.yaml をまだ持っていない場合
エクスポート OPENAI_API_KEY=sk-...
python -m sub_agent_mcp.main
# または: uv run sub-agent-mcp
サーバーは http://0.0.0.0:8000/mcp でリッスンします (マシンから http://localhost:8000/mcp としてアクセス可能)。
重要: 設定例では、MCP サーバーに Docker サービス名を指定しています。ローカル Python の場合、モック サーバーを実行して localhost URL (以下の表を参照) を使用するか、Compose 経由でモックのみを実行する必要があります。
# ターミナル 1: モックツールサーバーのみ
docker 構成ファイルシステム-mcp 検索-mcp
# ターミナル 2: サブエージェント サーバー (agents.yaml URL を編集して localhost にした後)
エクスポート OPENAI_API_KEY=sk-...
python -m sub_agent_mcp.main
ローカル MCP URL と Docker MCP URL
環境
ファイルシステム MCP URL
MCP URLを検索
Docker Compose ネットワーク
http://ファイルシステム-mcp:8001/mcp
http://search-mcp:8002/mcp
ホストマシン/ローカルPython
http://localhost:8001/mcp
http://localhost:8002/mcp
パス C — 事前構築済みイメージ (GHCR)
イメージは、 main へのプッシュごとではなく、 v* (たとえば、 v0.1.2 ) に一致する git タグで公開されます。
docker run -p 8000:8000 \
-e OPENAI_API_KEY=sk-... \
-v " $( pwd ) /config/agents.yaml:/app/config/agents.yaml:ro " \
ghcr.io/stormaref/sub-agent-mcp:latest
独自の Agents.yaml をマウントし、MCP URL 値が到達可能であることを確認します。

コンテナ内から (必要に応じてホスト ネットワーキング、サービス名、または host.docker.internal を使用します)。
OpenAPI ドキュメント (登録された MCP ツールから生成):
カール -s http://localhost:8000/mcp/openapi.json |頭
Open WebUI (OpenAPI ツール サーバー) — OpenAPI (MCP ではなく) として登録します。
WebUI を開くと、仕様 (例: /tools/researcher ) からのツール パスが URL に追加されます。
curl -s -X POST http://localhost:8000/mcp/tools/researcher \
-H ' コンテンツ タイプ: application/json ' \
-d ' {"プロンプト":"アクセスできるツールを要約します。"} '
Docker の健全性 — イメージの健全性チェックは http://127.0.0.1:8000/mcp を調査します。
機能チェック — Cursor (下記) を接続した後、短いプロンプトで研究者ツールを呼び出すように親エージェントに依頼します。実行が成功すると、 { "response": "..." } が返されます。
[カーソル設定] → [MCP] を開きます (または MCP 構成 JSON を編集します)。
{
"mcpサーバー": {
"サブエージェント-mcp" : {
"url" : " http://localhost:8000/mcp "
}
}
}
サブエージェント MCP プロセスが実行中であり、カーソルが localhost:8000 に到達できることを確認します。 WSL または Docker Desktop で、サーバーが VM またはコンテナーで実行されている場合はポート転送を確認します。
MCP ツールをリロードします。エージェントごとに 1 つのツールが表示されます (たとえば、 Researcher )。
親エージェントの委任プロンプトの例:
「アクセスできるツールを要約してください」というプロンプトを表示して研究者ツールを呼び出します。
その他の MCP クライアント - ストリーミング可能な HTTP をサポートするクライアントは、 http://<host>:8000/mcp に接続できます。 URL ベースのサーバー構成については、クライアントの MCP ドキュメントを参照してください。このサーバーは標準入出力トランスポートを使用しません。
フローチャート LR
部分グラフの親 [親]
カーソル[カーソルエージェント]
終わり
サブグラフ subagentmcp [サブ_エージェント_MCP]
ツールMCP[エージェントツール 例:研究者】
ビルダー[エージェントビルダー]
終わり
サブグラフ サブエージェント [サブエージェント ランタイム]
LC[LangChain 作成エージェント]
LLM[OpenAI対応LLM]
終わり
サブグループ

aph ダウンストリーム [エージェントごとの MCP サーバー]
FS[ファイルシステム MCP]
SRCH[MCP検索]
終わり
カーソル -->|ストリーミング可能な HTTP /mcp|ツールMCP
ツールMCP --> ビルダー
ビルダー --> LC
LC --> LLM
LC --> FS
LC --> SRCH
読み込み中
エージェントツールの流れ
起動時に config/agents.yaml (または AGENTS_CONFIG_PATH ) をロードして検証します。
id で名前を付けて、エージェントごとに 1 つの MCP ツールを登録します。
ツールが呼び出されると、そのエージェントの llm.* から OpenAI 互換のチャット モデルを構築します。
エージェントの mcp_servers に接続し、ツールを検出し、設定されている場合は tools_allowlist を適用します。
LangChain エージェント ループを実行します ( AGENT_RECURSION_LIMIT によって制限されます)。
最終アシスタント メッセージを { "response": "..." } として返すか、失敗した場合は { "error": "..." } として返します。
エージェントは起動時に config/agents.yaml からロードされます。パスを AGENTS_CONFIG_PATH でオーバーライドします。
エージェント:
- ID : 研究者
タイトル : 研究エージェント
説明: 「研究任務に特化したエージェント」
llm :
Base_uri : https://api.openai.com/v1
api_key : ${OPENAI_API_KEY}
モデルID : gpt-4.1-mini
システムプロンプト: |
あなたは有能な研究助手です。
mcp_servers :
- 名前 : ファイルシステム
トランスポート: streamable_http
URL : http://filesystem-mcp:8001/mcp
ヘッダー: {}
- 名前 : 検索
トランスポート: streamable_http
URL : http://search-mcp:8002/mcp
ツール許可リスト :
- ファイルシステム.read_file
- search.web_search
開始点として config/agents.example.yaml をコピーします。
フィールド
説明
ID
ユニークなナメクジ。小文字で始まり、その後小文字、数字、-、または _ で始まる必要があります
タイトル
人間が読める名前
説明
エージェントの目的 (親に公開されるツールの説明に表示されます)
llm.base_uri
OpenAI互換APIベースURL
llm.api_key
APIキー; ${ENV_VAR} 置換をサポートします
llm.model_id
プロバイダーのモデル識別子
llm.reasoning_effort
オプションの推論バジェット: none 、 minimum 、 low 、 middle 、 high 、 xhigh
llm.reasoning_summar

y
オプションの推論要約スタイル: auto 、 concise 、 詳細 (より長い場合は 詳細を使用します)
llm.冗長性
推論モデルのオプションの応答冗長性: low、medium、high
llm.max_tokens
オプションの最大出力トークン (回答が短くなった場合は増加します)
llm.温度
オプションのサンプリング温度 (0.0 – 2.0)
システムプロンプト
サブエージェントのシステムメッセージ
mcp_servers
リモート MCP サーバーのリスト (トランスポートは streamable_http である必要があります)
mcp_servers[].name
修飾されたツール名で使用される短縮名 ( name.tool )
mcp_servers[].url
ストリーミング可能な HTTP MCP エンドポイント (標準レイアウトの場合は /mcp で終わる必要があります)
mcp_servers[].bearer_token
オプションのベアラートークン。認可として送信: Bearer ... ( ${ENV_VAR} をサポート)
mcp_servers[].headers
ベアラー認証と統合されたオプションの追加の HTTP ヘッダー
ツール許可リスト
オプションのserver.tool名のリスト。接続されたサーバーからのすべてのツールを許可するのを省略します
環境変数置換は、 ${VAR} および ${VAR:-default} をサポートします。 VAR が設定されておらず、デフォルトが提供されていない場合、起動は明らかなエラーで失敗します。
2 番目のエージェント (役割が異なり、追加の MCP サーバーなし)
- ID : ライター
タイトル : ライティングエージェント
説明: " テキストの下書きと編集 "
llm :
Base_uri : https://api.openai.com/v1
api_key : ${OPENAI_API_KEY}
モデルID : gpt-4.1-mi

[切り捨てられた]

## Original Extract

MCP for spawning sub-agents everywhere. Contribute to stormaref/Sub-Agent-MCP development by creating an account on GitHub.

GitHub - stormaref/Sub-Agent-MCP: MCP for spawning sub-agents everywhere · GitHub
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
stormaref
/
Sub-Agent-MCP
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .github/ workflows .github/ workflows .vscode .vscode config config docker/ mock_mcp docker/ mock_mcp scripts scripts src/ sub_agent_mcp src/ sub_agent_mcp tests tests .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
Production-ready Python MCP server for LLM delegation and sub-agent orchestration . A parent LLM (for example, Cursor’s agent) connects to this server and delegates work by calling a tool named after each agent defined in YAML (for example, researcher ).
Each sub-agent is defined in YAML with its own LLM, system prompt, and optional downstream MCP tool servers.
Sub-Agent MCP sits between a parent LLM and one or more specialized sub-agents :
The parent connects to this server over Streamable HTTP at /mcp .
At startup, each agent in agents.yaml is registered as an MCP tool named by its id .
The parent calls that tool with a prompt ; the server runs the sub-agent and returns the final response.
Each sub-agent is a LangChain agent with its own OpenAI-compatible LLM, system prompt, and optional connections to other MCP servers (filesystem, search, your own tools, and so on). Tool access can be restricted per agent with an allowlist.
sequenceDiagram
participant Parent as Parent_LLM
participant SAMCP as Sub_Agent_MCP
participant Agent as LangChain_sub_agent
participant LLM as OpenAI_compatible_LLM
participant Tools as Downstream_MCP_servers
Parent->>SAMCP: researcher(prompt)
SAMCP->>Agent: build agent + downstream tools
Agent->>LLM: reasoning loop
Agent->>Tools: MCP tool calls
Tools-->>Agent: tool results
LLM-->>Agent: final answer
Agent-->>SAMCP: response text
SAMCP-->>Parent: {"response": "..."}
Loading
This is different from giving one agent every tool in the workspace: the parent stays lightweight, roles stay explicit, and each sub-agent only sees the MCP servers and tools you configure for it.
Delegation without context bloat — The parent calls agent tools directly; it does not need every downstream tool schema in its own context.
Per-role configuration — Different id s can use different models, prompts, MCP servers, and tool allowlists.
Production-oriented — Pydantic-validated YAML, structured logging, Docker health checks, CI, and GHCR images on release tags.
OpenAI-compatible providers — Point llm.base_uri at OpenAI, Azure, Ollama, LM Studio, or any compatible API.
Streamable HTTP only ( streamable-http ); no stdio or legacy SSE.
YAML agent definitions with strict Pydantic validation.
Environment substitution: ${VAR} and ${VAR:-default} .
LangChain 1.x create_agent with OpenAI-compatible chat models.
Per-agent MCP connections via langchain-mcp-adapters , with optional server.tool allowlists.
Agent tool descriptions never expose API keys.
Structured logging ( structlog ).
Docker image with health check; GitHub Actions CI.
Container images published to GHCR on version tags ( v0.x.y ).
MCP tools exposed by this server
Each agent in agents.yaml is registered as a tool named by its id (for example, researcher ). Each tool accepts a single prompt argument and runs that sub-agent.
Requirement
Notes
Python 3.10+
CI uses 3.12; requires-python >= 3.10 in pyproject.toml
API key for your LLM provider
Default example uses OPENAI_API_KEY
uv (recommended) or pip
Matches CI install path
Docker + Compose (optional)
Recommended for first run; includes mock MCP tool servers
Quick start
Choose one path. Docker Compose is the fastest way to run the full demo stack (server + mock tool servers).
Path A — Docker Compose (recommended)
export OPENAI_API_KEY=sk-...
docker compose up --build
Service
Port
Endpoint
Sub-Agent MCP
8000
http://localhost:8000/mcp
Mock filesystem MCP
8001
http://localhost:8001/mcp
Mock search MCP
8002
http://localhost:8002/mcp
The bundled config/agents.yaml uses Docker network hostnames ( filesystem-mcp , search-mcp ), which resolve correctly inside Compose.
uv sync --dev
# or: pip install -e ".[dev]"
cp config/agents.example.yaml config/agents.yaml # if you don't have config/agents.yaml yet
export OPENAI_API_KEY=sk-...
python -m sub_agent_mcp.main
# or: uv run sub-agent-mcp
Server listens at http://0.0.0.0:8000/mcp (reachable as http://localhost:8000/mcp from your machine).
Important: The example config points MCP servers at Docker service names. For local Python you must either run the mock servers and use localhost URLs (see table below), or run only the mocks via Compose:
# Terminal 1: mock tool servers only
docker compose up filesystem-mcp search-mcp
# Terminal 2: sub-agent server (after editing agents.yaml URLs to localhost)
export OPENAI_API_KEY=sk-...
python -m sub_agent_mcp.main
Local vs Docker MCP URLs
Environment
filesystem MCP URL
search MCP URL
Docker Compose network
http://filesystem-mcp:8001/mcp
http://search-mcp:8002/mcp
Host machine / local Python
http://localhost:8001/mcp
http://localhost:8002/mcp
Path C — Prebuilt image (GHCR)
Images are published on git tags matching v* (for example v0.1.2 ), not on every push to main .
docker run -p 8000:8000 \
-e OPENAI_API_KEY=sk-... \
-v " $( pwd ) /config/agents.yaml:/app/config/agents.yaml:ro " \
ghcr.io/stormaref/sub-agent-mcp:latest
Mount your own agents.yaml and ensure MCP url values are reachable from inside the container (use host networking, service names, or host.docker.internal as appropriate).
OpenAPI document (generated from registered MCP tools):
curl -s http://localhost:8000/mcp/openapi.json | head
Open WebUI (OpenAPI tool server) — Register as OpenAPI (not MCP):
Open WebUI appends tool paths from the spec (for example /tools/researcher ) to that URL:
curl -s -X POST http://localhost:8000/mcp/tools/researcher \
-H ' Content-Type: application/json ' \
-d ' {"prompt":"Summarize what tools you have access to."} '
Docker health — The image health check probes http://127.0.0.1:8000/mcp .
Functional check — After connecting Cursor (below), ask the parent agent to call the researcher tool with a short prompt. A successful run returns { "response": "..." } .
Open Cursor Settings → MCP (or edit your MCP config JSON).
{
"mcpServers" : {
"sub-agent-mcp" : {
"url" : " http://localhost:8000/mcp "
}
}
}
Ensure the Sub-Agent MCP process is running and Cursor can reach localhost:8000 . On WSL or Docker Desktop, confirm port forwarding if the server runs in a VM or container.
Reload MCP tools. You should see one tool per agent (for example, researcher ).
Example delegation prompt for the parent agent:
Call the researcher tool with prompt: "Summarize what tools you have access to."
Other MCP clients — Any client that supports Streamable HTTP can connect to http://<host>:8000/mcp . Refer to your client’s MCP documentation for URL-based server configuration; this server does not use stdio transport.
flowchart LR
subgraph parent [Parent]
Cursor[Cursor agent]
end
subgraph subagentmcp [Sub_Agent_MCP]
ToolsMCP[agent tools e.g. researcher]
Builder[Agent builder]
end
subgraph subagent [Sub_agent runtime]
LC[LangChain create_agent]
LLM[OpenAI_compatible LLM]
end
subgraph downstream [Per_agent MCP servers]
FS[filesystem MCP]
SRCH[search MCP]
end
Cursor -->|Streamable HTTP /mcp| ToolsMCP
ToolsMCP --> Builder
Builder --> LC
LC --> LLM
LC --> FS
LC --> SRCH
Loading
Agent tool flow
Load and validate config/agents.yaml (or AGENTS_CONFIG_PATH ) at startup.
Register one MCP tool per agent, named by id .
When a tool is called, build an OpenAI-compatible chat model from that agent’s llm.* .
Connect to the agent’s mcp_servers , discover tools, apply tool_allowlist if set.
Run the LangChain agent loop (bounded by AGENT_RECURSION_LIMIT ).
Return the final assistant message as { "response": "..." } , or { "error": "..." } on failure.
Agents are loaded from config/agents.yaml at startup. Override the path with AGENTS_CONFIG_PATH .
agents :
- id : researcher
title : Research Agent
description : " Agent specialized in research tasks "
llm :
base_uri : https://api.openai.com/v1
api_key : ${OPENAI_API_KEY}
model_id : gpt-4.1-mini
system_prompt : |
You are a helpful research assistant.
mcp_servers :
- name : filesystem
transport : streamable_http
url : http://filesystem-mcp:8001/mcp
headers : {}
- name : search
transport : streamable_http
url : http://search-mcp:8002/mcp
tool_allowlist :
- filesystem.read_file
- search.web_search
Copy config/agents.example.yaml as a starting point.
Field
Description
id
Unique slug; must start with a lowercase letter, then lowercase letters, digits, - , or _
title
Human-readable name
description
Agent purpose (shown in the tool description exposed to the parent)
llm.base_uri
OpenAI-compatible API base URL
llm.api_key
API key; supports ${ENV_VAR} substitution
llm.model_id
Model identifier for the provider
llm.reasoning_effort
Optional reasoning budget: none , minimal , low , medium , high , xhigh
llm.reasoning_summary
Optional reasoning summary style: auto , concise , detailed (use detailed for longer)
llm.verbosity
Optional response verbosity for reasoning models: low , medium , high
llm.max_tokens
Optional max output tokens (increase if answers are cut short)
llm.temperature
Optional sampling temperature ( 0.0 – 2.0 )
system_prompt
System message for the sub-agent
mcp_servers
List of remote MCP servers ( transport must be streamable_http )
mcp_servers[].name
Short name used in qualified tool names ( name.tool )
mcp_servers[].url
Streamable HTTP MCP endpoint (must end with /mcp for standard layouts)
mcp_servers[].bearer_token
Optional bearer token; sent as Authorization: Bearer ... (supports ${ENV_VAR} )
mcp_servers[].headers
Optional extra HTTP headers merged with bearer auth
tool_allowlist
Optional list of server.tool names; omit to allow all tools from connected servers
Environment variable substitution supports ${VAR} and ${VAR:-default} . If VAR is unset and no default is provided, startup fails with a clear error.
Second agent (different role, no extra MCP servers)
- id : writer
title : Writing Agent
description : " Drafts and edits text "
llm :
base_uri : https://api.openai.com/v1
api_key : ${OPENAI_API_KEY}
model_id : gpt-4.1-mi

[truncated]

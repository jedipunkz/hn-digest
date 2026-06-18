---
source: "https://github.com/andreisilva1/OSymandias"
hn_url: "https://news.ycombinator.com/item?id=48591144"
title: "Show HN: OSymandias – Open-source runtime for multi-agent AI systems"
article_title: "GitHub - andreisilva1/OSymandias: Multi-agent AI runtime with OS-inspired primitives — job scheduling, DAG orchestration, memory, tool execution and real-time observability. Built with FastAPI, Celery, PostgreSQL and LiteLLM. · GitHub"
author: "andreisilva1"
captured_at: "2026-06-18T21:45:32Z"
capture_tool: "hn-digest"
hn_id: 48591144
score: 2
comments: 0
posted_at: "2026-06-18T20:32:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OSymandias – Open-source runtime for multi-agent AI systems

- HN: [48591144](https://news.ycombinator.com/item?id=48591144)
- Source: [github.com](https://github.com/andreisilva1/OSymandias)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T20:32:43Z

## Translation

タイトル: Show HN: OSymandias – マルチエージェント AI システム用のオープンソース ランタイム
記事のタイトル: GitHub - andreisilva1/OSymandias: OS にインスピレーションを得たプリミティブを備えたマルチエージェント AI ランタイム - ジョブ スケジューリング、DAG オーケストレーション、メモリ、ツール実行、リアルタイムの可観測性。 FastAPI、Celery、PostgreSQL、LiteLLM で構築されています。 · GitHub
説明: OS からインスピレーションを得たプリミティブを備えたマルチエージェント AI ランタイム - ジョブ スケジューリング、DAG オーケストレーション、メモリ、ツール実行、リアルタイムの可観測性。 FastAPI、Celery、PostgreSQL、LiteLLM で構築されています。 - andreisilva1/OSymandias

記事本文:
GitHub - andreisilva1/OSymandias: OS からインスピレーションを得たプリミティブを備えたマルチエージェント AI ランタイム - ジョブ スケジューリング、DAG オーケストレーション、メモリ、ツール実行、リアルタイムの可観測性。 FastAPI、Celery、PostgreSQL、LiteLLM で構築されています。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
そして

レイシルバ1
/
オシマンディアス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
52 コミット 52 コミット .github/ workflows .github/ workflows backend backendfrontendfrontend sdk sdk .env.example .env.example .gitignore .gitignore DOCS.md DOCS.md ライセンス ライセンス OSY.compose.yml OSY.compose.yml OSY.nginx.conf OSY.nginx.conf OSymandias.svg OSymandias.svg README.md README.md Banner.svg Banner.svg osymandias.toml osymandias.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
「力ある者よ、私の作品を見てください、そして派遣してください。」
Python 開発者向けのマルチエージェント ランタイム。コマンド 1 つですべてを開始できます。
📖 完全なドキュメント → DOCS.md
OSymandias は、プロジェクトを完全なマルチエージェント ランタイムに変える Python ライブラリおよび CLI です。
pip インストール オシンマンディアス
OSY初期化
オシサーブ
PostgreSQL、Redis、RabbitMQ、Qdrant — Docker 経由で内部管理されます。 localhost:47759 のダッシュボード。セロリ労働者 4 人が準備できました。
前提条件: Python 3.11+、Docker
pip インストール オシンマンディアス
# OSY.compose.yml + OSY.nginx.conf + .env + サンプル osy_tools.py を生成します
OSY初期化
# すべてを始める
オシサーブ
http://localhost:47759 — ダッシュボードを開きます。
API は http://localhost:47760/api/v1 に直接あります。
osy stop # コンテナを一時停止し、データを保持します
osy down # コンテナを削除し、ボリュームを保持します
osy delete # コンテナーとボリュームを削除します (確認を求めます)
osy --version # インストールされているバージョンを表示します
CLI からイベントを追跡します。
OSY ログ # すべてのジョブにわたる最新の 50 イベント
osy logs < job-id > # 特定のジョブの最後の 50 イベント
osy logs < job-id > -f # 到着時にライブストリーム
osy logs < job-id > -f -t TASK_PROGRESS # イベント タイプでフィルタリングする
同時実行性をスケールアップするか、ワーカー ノードを追加します。
このマシンには # 個以上のスロットがあります
OSY サーブ --concur

rency 8 # または .env の OSY_WORKER_CONCURRENCY=8
# 追加のワーカー ノード (同じブローカー/redis を指す)
OSY_RABBITMQ_URL=amqp://... OSY_REDIS_URL=redis://... osy ワーカー --キュー エージェント、ツール --同時実行 8
ドッカーがいないのですか？代わりに --no-docker を使用して、外部管理サービスに接続します。
# .env のコメントを解除し、独自のインスタンスをポイントします。
# OSY_NO_DOCKER=1
# OSY_POSTGRES_URL=postgresql+asyncpg://user:pass@host:5432/osymandias
# OSY_REDIS_URL=redis://ホスト:6379/0
# OSY_RABBITMQ_URL=amqp://user:pass@host:5672/
# OSY_QDRANT_URL=http://ホスト:6333
OSY サーブ --no-docker
組み込みツール関数 ( @osy.tool )
Python 関数は、単一のデコレータを備えたエージェント ツールになります。
オシンマンディアスから OSY をインポート
@osy 。ツール
def fetch_competitor_data (会社 : str 、メトリクス : リスト [ str ]) -> dict :
"""内部データベースから競合他社の指標を取得します。"""
return { "会社" : 会社 , "データ" : [...]}
@osy 。ツール
def send_slack_message (チャネル: str 、テキスト: str ) -> dict :
"""Slack チャネルにメッセージを送信します。"""
return { "ok" : True }
型ヒントから推論されたスキーマ。 osyserve はすべての .py ファイルを自動的にスキャンします。YAML や構成ファイルは含まれません。その後、ダッシュボード ( /tools ) からツールをエージェントに割り当てることができます。
外部エージェント ( @osy.agent )
任意の Python 呼び出し可能 — LangChain チェーン、CrewAI クルー、LlamaIndex クエリ エンジン、またはプレーン Python — を OSymandias エージェントとして登録します。
オシマンディアスからのインポート osy 、OsyContext
@osy 。エージェント ( "ResearchAgent" 、フレームワーク = "langchain" 、
description = "Web コンテンツを検索して要約します" ,
llm_provider = "ollama" 、llm_model = "qwen2.5:7b" )
def Research_agent (タスク : str 、 ctx : OsyContext ) -> dict :
チェーン = build_langchain_chain ()
ctx 。 Emit_event ( "TASK_PROGRESS" , { "ステップ" : "実行中のチェーン" })
return { "概要" : チェーン 。 (タスク) を呼び出す}
すべてのクワーグは

ダッシュボードのオプションのメタデータ。エージェントは、宣言内容に関係なく実行されます。
osymandias.toml (プロジェクトルート) でスキャンするモジュールを宣言します。
エージェントモジュール = [
" myproject.agents " 、
" myproject.crews " 、
】
これらのモジュール内のエージェントは、osyserve 上で自動的に検出され、登録されます。
すべての @osy.agent 関数は、オプションで ctx パラメーターとして OsyContext を受け取ります。
@osy 。エージェント ( "OrchestratorAgent" )
def Orchestrate (タスク: str , ctx: OsyContext) -> dict:
# 共有メモリ — 同じジョブ内のすべてのエージェントが読み取り/書き込み可能
ctx 。 write_memory ( "計画" , { "ステップ" : 1 , "目標" : タスク })
データ = ctx 。 read_memory ( "previous_output" )
# ライブ イベント — ダッシュボード イベント フィードにストリーミングされます
ctx 。 Emit_event ( "TASK_PROGRESS" , { "pct" : 50 , "message" : "途中" })
# サブタスク — 子タスクを生成し、結果を待ちます
task_ids = ctx 。 spawn_tasks ([
{ "title" : "リサーチ" 、 "agent_type" : "リサーチエージェント" 、 "説明" : タスク },
{ "title" : "分析" 、 "agent_type" : "AnalystAgent" 、 "説明" : タスク },
])
結果 = ctx 。 wait_for_tasks ( task_ids )
return { "マージ" : 結果 }
方法
説明
ctx.write_memory(キー, 値)
共有ジョブメモリへの書き込み
ctx.read_memory(キー)
共有ジョブメモリからの読み取り
ctx.emit_event(タイプ、ペイロード)
イベントをダッシュボードのライブフィードにストリーミングする
ctx.spawn_tasks(リスト)
サブタスクを生成します。タスクIDのリストを返します
ctx.wait_for_tasks(ids)
すべてのサブタスクが完了するまでブロックします。出力を返します
サブタスクは、ジョブ タイムライン ダッシュボードにツリーとして表示されます。
エージェントにツールを提供する 3 つの方法
何
どうやって
内蔵
web_search 、 read_url 、 http_request 、 write_to_job_memory 、 search_memory 、 python_eval 、 run_shell 、 read_file 、 write_file 、 send_message 、 spawn_agent … (合計 20)
設定不要 — いつでも利用可能
@osy.tool
Python 関数
デコラット

e+osyサーブ
Webhook
任意の HTTP エンドポイント
ダッシュボードにURLを登録する
仕組み
仕事 → ユーザーが提出した目標 (「X について調査してレポートを書く」)
└── タスク ×N → 特定のエージェント タイプに割り当てられたサブタスク
└── AgentInstance → 実行中のエージェント ループ (LLM + ツール + メモリ)
§── ToolCall → web_search / @osy.tool / webhook / ...
└── サブタスク → ctx.spawn_tasks([...]) → 子タスク ×N
ジョブは PlannerAgent によってタスクに分解されます。プランナーは、利用可能なエージェント タイプ (組み込みおよび登録したすべての @osy.agent) の完全なリストを受け取るため、自然言語ジョブは自動的に外部エージェントにルーティングされます。タスクは特殊なエージェント間で並行して実行されます。 EvaluatorAgent は出力にスコアを付け、信頼度がしきい値を下回る場合は再試行します。
ページ
パス
説明
求人
/ジョブ
検索、フィルター、ページネーションを備えた求人リスト
仕事内容
/ジョブ/{id}
出力、イベント、タスク、サブタスクツリーのタイムライン
エージェント
/エージェント
エージェント レジストリ — 組み込みおよび外部の適応型詳細パネル
ツール
/ツール
組み込みツールとユーザーツール
記憶
/メモリ
検索、スコープによるフィルタリング、エントリの削除
イベント
/イベント
一時停止/再開付きのライブ イベント ストリーム
メトリクス
/メトリクス
7日間チャート、トークン、コスト、成功率
サポートされている LLM プロバイダー
プロバイダー
キー
OpenAI
OPENAI_API_KEY
人間的
ANTHROPIC_API_KEY
ディープシーク
DEEPSEEK_API_KEY
グロク
GROQ_API_KEY
ジェミニ
GEMINI_API_KEY
オラマ (地元)
鍵は必要ありません
ダッシュボードからエージェントごとにモデルを切り替えます。再起動は必要ありません。
# 自然言語 — PlannerAgent が自動的に分解します
curl -X POST http://localhost:47760/api/v1/jobs \
-H " Content-Type: application/json " \
-d ' {"title":"私の仕事","description":"2024 年のヨーロッパの EV 市場を調査します。","priority":"NORMAL","input_payload":{}} '
# 明示的なタスク プランを使用してプランナーをバイパスする
カール -X POST http://localhost:47

760/API/v1/ジョブ \
-H " Content-Type: application/json " \
-d ' {"title":"私の仕事","description":"...","priority":"NORMAL","input_payload":{"__task_plan__":[{"title":"Research","agent_type":"ResearchAgent","description":"ヨーロッパのEV市場"}]}} '
# 完了または失敗したジョブを再送信します (入力をコピーし、新しいジョブを作成します)
curl -X POST http://localhost:47760/api/v1/jobs/ < ジョブ ID > /resubmit
完全な API リファレンス: http://localhost:47760/api/v1/docs
オシマンディアス/
§── sdk/ Python パッケージ — osymandias + osy CLI
│ └── オシンマンディアス/
│ §── cli/osy init / サーブ / 停止 / ダウン / 削除 / ログ / ワーカー
│ §── ランタイム/ FastAPI + Celery + エージェント
│ §──decorator.py @osy.tool + @osy.agent
│ §── context.py OsyContext (メモリ、イベント、サブタスク)
│ §── Discovery.py @osy.tool スキャナ
│ §── tools_server.py ローカル HTTP ツール サーバー
│ §──assets.py GitHub アセットフェッチャー + キャッシュ
│ └── process.py サブプロセスマネージャー
§── フロントエンド/ Next.js 14 ダッシュボード (CI によって構築、ホイールにバンドル)
§── バックエンド/レガシー スタンドアロン バックエンド (参照用に保持)
━── .github/workflows/
━── release.yml タグプッシュ → ビルド → GitHub Release + PyPI
貢献する
git clone https://github.com/andreisilva1/OSymandias
cd オシマンディアス
# 編集可能モードで SDK をインストールします
pip install -e ./sdk
# スキャフォールド設定ファイル
OSY初期化
# インフラ + API を開始します (ローカル フロントエンド/アウト ビルドが自動的に取得されます)
オシサーブ
# ライブフロントエンド開発の場合 (別のターミナル)
CDフロントエンド
npmインストール
npm run dev # http://localhost:3000 — ホットリロード
FastAPI · Next.js · Celery · PostgreSQL · Redis · RabbitMQ · Qdrant · LiteLLM で構築
について
OS からインスピレーションを得たプリミティブを備えたマルチエージェント AI ランタイム - ジョブ スケジューリング、DAG オーケストレーション、メモリ、ツール実行

イオンとリアルタイムの観測可能性。 FastAPI、Celery、PostgreSQL、LiteLLM で構築されています。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.1
最新の
2026 年 6 月 18 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Multi-agent AI runtime with OS-inspired primitives — job scheduling, DAG orchestration, memory, tool execution and real-time observability. Built with FastAPI, Celery, PostgreSQL and LiteLLM. - andreisilva1/OSymandias

GitHub - andreisilva1/OSymandias: Multi-agent AI runtime with OS-inspired primitives — job scheduling, DAG orchestration, memory, tool execution and real-time observability. Built with FastAPI, Celery, PostgreSQL and LiteLLM. · GitHub
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
andreisilva1
/
OSymandias
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
52 Commits 52 Commits .github/ workflows .github/ workflows backend backend frontend frontend sdk sdk .env.example .env.example .gitignore .gitignore DOCS.md DOCS.md LICENSE LICENSE OSY.compose.yml OSY.compose.yml OSY.nginx.conf OSY.nginx.conf OSymandias.svg OSymandias.svg README.md README.md banner.svg banner.svg osymandias.toml osymandias.toml View all files Repository files navigation
"Look on my works, ye Mighty, and dispatch."
Multi-agent runtime for Python developers. One command to start everything.
📖 Full documentation → DOCS.md
OSymandias is a Python library and CLI that turns your project into a full multi-agent runtime.
pip install osymandias
osy init
osy serve
PostgreSQL, Redis, RabbitMQ, Qdrant — managed internally via Docker. Dashboard at localhost:47759 . Four Celery workers ready.
Prerequisites: Python 3.11+, Docker
pip install osymandias
# Generate OSY.compose.yml + OSY.nginx.conf + .env + sample osy_tools.py
osy init
# Start everything
osy serve
Open http://localhost:47759 — dashboard.
API directly at http://localhost:47760/api/v1 .
osy stop # pause containers, keep data
osy down # remove containers, keep volumes
osy delete # remove containers + volumes (asks for confirmation)
osy --version # print installed version
Tail events from the CLI:
osy logs # last 50 events across all jobs
osy logs < job-id > # last 50 events for a specific job
osy logs < job-id > -f # live-stream as they arrive
osy logs < job-id > -f -t TASK_PROGRESS # filter by event type
Scale up concurrency or add worker nodes:
# More slots on this machine
osy serve --concurrency 8 # or OSY_WORKER_CONCURRENCY=8 in .env
# Additional worker nodes (point to the same broker/redis)
OSY_RABBITMQ_URL=amqp://... OSY_REDIS_URL=redis://... osy workers --queues agents,tools --concurrency 8
No Docker? Use --no-docker to connect to externally managed services instead:
# Uncomment in .env and point to your own instances:
# OSY_NO_DOCKER=1
# OSY_POSTGRES_URL=postgresql+asyncpg://user:pass@host:5432/osymandias
# OSY_REDIS_URL=redis://host:6379/0
# OSY_RABBITMQ_URL=amqp://user:pass@host:5672/
# OSY_QDRANT_URL=http://host:6333
osy serve --no-docker
Built-in tool functions ( @osy.tool )
Your Python functions become agent tools with a single decorator:
from osymandias import osy
@ osy . tool
def fetch_competitor_data ( company : str , metrics : list [ str ]) -> dict :
"""Fetch competitor metrics from internal database."""
return { "company" : company , "data" : [...]}
@ osy . tool
def send_slack_message ( channel : str , text : str ) -> dict :
"""Send a message to a Slack channel."""
return { "ok" : True }
Schema inferred from type hints. osy serve scans all .py files automatically — no YAML, no config files. Tools are then assignable to agents from the dashboard ( /tools ).
External agents ( @osy.agent )
Register any Python callable — LangChain chain, CrewAI crew, LlamaIndex query engine, or plain Python — as an OSymandias agent:
from osymandias import osy , OsyContext
@ osy . agent ( "ResearchAgent" , framework = "langchain" ,
description = "Searches and summarises web content" ,
llm_provider = "ollama" , llm_model = "qwen2.5:7b" )
def research_agent ( task : str , ctx : OsyContext ) -> dict :
chain = build_langchain_chain ()
ctx . emit_event ( "TASK_PROGRESS" , { "step" : "running chain" })
return { "summary" : chain . invoke ( task )}
All kwargs are optional metadata for the dashboard. The agent executes regardless of what's declared.
Declare which modules to scan in osymandias.toml (project root):
agent_modules = [
" myproject.agents " ,
" myproject.crews " ,
]
Agents in those modules are discovered and registered automatically on osy serve .
Every @osy.agent function optionally receives an OsyContext as its ctx parameter:
@ osy . agent ( "OrchestratorAgent" )
def orchestrate ( task : str , ctx : OsyContext ) -> dict :
# shared memory — any agent in the same job can read/write
ctx . write_memory ( "plan" , { "step" : 1 , "goal" : task })
data = ctx . read_memory ( "previous_output" )
# live events — streamed to the dashboard event feed
ctx . emit_event ( "TASK_PROGRESS" , { "pct" : 50 , "message" : "halfway" })
# sub-tasks — spawn child tasks and wait for results
task_ids = ctx . spawn_tasks ([
{ "title" : "Research" , "agent_type" : "ResearchAgent" , "description" : task },
{ "title" : "Analyse" , "agent_type" : "AnalystAgent" , "description" : task },
])
results = ctx . wait_for_tasks ( task_ids )
return { "merged" : results }
Method
Description
ctx.write_memory(key, value)
Write to shared job memory
ctx.read_memory(key)
Read from shared job memory
ctx.emit_event(type, payload)
Stream event to dashboard live feed
ctx.spawn_tasks(list)
Spawn sub-tasks; returns list of task IDs
ctx.wait_for_tasks(ids)
Block until all sub-tasks complete; returns their outputs
Sub-tasks are visible as a tree in the job timeline dashboard.
Three ways to give agents tools
What
How
Built-in
web_search , read_url , http_request , write_to_job_memory , search_memory , python_eval , run_shell , read_file , write_file , send_message , spawn_agent … (20 total)
Zero config — always available
@osy.tool
Your Python functions
Decorate + osy serve
Webhook
Any HTTP endpoint
Register URL in the dashboard
How it works
Job → A user-submitted goal ("research and write a report on X")
└── Task ×N → Subtask assigned to a specific agent type
└── AgentInstance → A running agent loop (LLM + tools + memory)
├── ToolCall → web_search / @osy.tool / webhook / ...
└── Sub-task → ctx.spawn_tasks([...]) → child Task ×N
Jobs are decomposed into tasks by a PlannerAgent. The planner receives the full list of available agent types — builtin and every @osy.agent you registered — so natural-language jobs route to external agents automatically. Tasks execute in parallel across specialized agents. An EvaluatorAgent scores outputs and retries if confidence is below threshold.
Page
Path
Description
Jobs
/jobs
Job list with search, filter, pagination
Job detail
/jobs/{id}
Output, events, tasks, sub-task tree timeline
Agents
/agents
Agent registry — builtin and external, adaptive detail panel
Tools
/tools
Built-in and user tools
Memory
/memory
Search, filter by scope, delete entries
Events
/events
Live event stream with pause/resume
Metrics
/metrics
7-day chart, tokens, cost, success rate
Supported LLM providers
Provider
Key
OpenAI
OPENAI_API_KEY
Anthropic
ANTHROPIC_API_KEY
DeepSeek
DEEPSEEK_API_KEY
Groq
GROQ_API_KEY
Gemini
GEMINI_API_KEY
Ollama (local)
no key needed
Switch models per-agent from the dashboard — no restart required.
# Natural language — PlannerAgent decomposes it automatically
curl -X POST http://localhost:47760/api/v1/jobs \
-H " Content-Type: application/json " \
-d ' {"title":"My Job","description":"Research the EV market in Europe in 2024.","priority":"NORMAL","input_payload":{}} '
# Bypass the planner with an explicit task plan
curl -X POST http://localhost:47760/api/v1/jobs \
-H " Content-Type: application/json " \
-d ' {"title":"My Job","description":"...","priority":"NORMAL","input_payload":{"__task_plan__":[{"title":"Research","agent_type":"ResearchAgent","description":"EV market in Europe"}]}} '
# Resubmit a completed or failed job (copies input, creates a new job)
curl -X POST http://localhost:47760/api/v1/jobs/ < job-id > /resubmit
Full API reference: http://localhost:47760/api/v1/docs
OSymandias/
├── sdk/ Python package — osymandias + osy CLI
│ └── osymandias/
│ ├── cli/ osy init / serve / stop / down / delete / logs / workers
│ ├── runtime/ FastAPI + Celery + agents
│ ├── decorator.py @osy.tool + @osy.agent
│ ├── context.py OsyContext (memory, events, sub-tasks)
│ ├── discovery.py @osy.tool scanner
│ ├── tool_server.py local HTTP tool server
│ ├── assets.py GitHub asset fetcher + cache
│ └── process.py subprocess manager
├── frontend/ Next.js 14 dashboard (built by CI, bundled into wheel)
├── backend/ Legacy standalone backend (kept for reference)
└── .github/workflows/
└── release.yml Tag push → build → GitHub Release + PyPI
Contributing
git clone https://github.com/andreisilva1/OSymandias
cd OSymandias
# Install the sdk in editable mode
pip install -e ./sdk
# Scaffold config files
osy init
# Start infra + API (the local frontend/out build is picked up automatically)
osy serve
# For live frontend development (separate terminal)
cd frontend
npm install
npm run dev # http://localhost:3000 — hot reload
Built with FastAPI · Next.js · Celery · PostgreSQL · Redis · RabbitMQ · Qdrant · LiteLLM
About
Multi-agent AI runtime with OS-inspired primitives — job scheduling, DAG orchestration, memory, tool execution and real-time observability. Built with FastAPI, Celery, PostgreSQL and LiteLLM.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.1
Latest
Jun 18, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/JoniMartin27/lookspan"
hn_url: "https://news.ycombinator.com/item?id=48391306"
title: "Show HN: Lookspan – local-first observability for AI agents (npx lookspan)"
article_title: "GitHub - JoniMartin27/lookspan: Local-first observability dashboard for AI agents. MCP-native. Look at every span your agents emit. · GitHub"
author: "JoniMartin"
captured_at: "2026-06-04T01:06:30Z"
capture_tool: "hn-digest"
hn_id: 48391306
score: 1
comments: 0
posted_at: "2026-06-03T23:05:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lookspan – local-first observability for AI agents (npx lookspan)

- HN: [48391306](https://news.ycombinator.com/item?id=48391306)
- Source: [github.com](https://github.com/JoniMartin27/lookspan)
- Score: 1
- Comments: 0
- Posted: 2026-06-03T23:05:30Z

## Translation

タイトル: Show HN: Lookspan – AI エージェントのローカルファースト可観測性 (npx lookspan)
記事のタイトル: GitHub - JoniMartin27/lookspan: AI エージェント向けのローカルファースト可観測性ダッシュボード。 MCPネイティブ。エージェントが発するすべてのスパンを確認してください。 · GitHub
説明: AI エージェント用のローカルファースト可観測性ダッシュボード。 MCPネイティブ。エージェントが発するすべてのスパンを確認してください。 - JoniMartin27/lookspan

記事本文:
GitHub - JoniMartin27/lookspan: AI エージェント向けのローカルファースト可観測性ダッシュボード。 MCPネイティブ。エージェントが発するすべてのスパンを確認してください。 · GitHub
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
ジョニマーティン27
/
ルックススパン
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
58 コミット 58 コミット .github .github apps/ ダッシュボード apps/ ダッシュボード docs docs 例 例 パッケージ パッケージ python python Web サイト Web サイト .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.es.md README.es.md README.md README.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント向けのローカルファースト可観測性ダッシュボード。 MCPネイティブ。エージェントが発信するすべてのスパンを確認します。
npx ルックスパン # → http://127.0.0.1:3100
▶ 75 秒のプレゼンテーション全体を見る
エージェント (MCP · LangGraph · CrewAI · OpenTelemetry · HTTP) → POST /api/ingest → SQLite → リアルタイムダッシュボード
🇪🇸 ¿スペイン語を優先しますか? Lee エル README スペイン語。
AI エージェントが不正な動作をしたとき、失敗したり、停止したり、静かにより多くのトークンを燃やしたりします。
予想通り — 何が起こったのかを段階的に確認するネイティブな方法はありません。既存の
可観測性ツールはクラウドファーストです。アカウント、API キー、配送が必要です。
本番データを他人のサーバーに転送します。
Lookspan は逆のアプローチを採用しています。すべてがマシン上で実行され、データも実行されます。
そこから離れることはなく、インフラコストはゼロです。アダプターを使用してエージェントを計測する
(または単に POST JSON)、ブラウザでダッシュボードを開きます。
npx ルックスパン # → http://127.0.0.1:3100、インストールなし、クラウドなし
任意の言語から最初のスパンを送信します。
カール -X POST http://127.0.0.1:3100/api/ingest \
-H " Content-Type: application/json " \
-d ' {"spans":[{"traceId":"t1","spanId":"s1","parentSpanId":null,"type":"llm_call","name":"agent.run","sta

rtedAt":"2026-06-02T10:00:00Z","endedAt":"2026-06-02T10:00:01Z","ステータス":"ok","フレームワーク":"カスタム", "モデル":"gpt-4o","プロバイダー":"openai","使用法":{"inputTokens":1000,"outputTokens":500,"costUsd":0}}]} '
http://127.0.0.1:3100 を開いて、サーバー側で計算されたコストを含むトレースが表示されることを確認します。
HTTP スパンの取り込み — POST /api/ingest はスパンの JSON バッチを受け入れます。 HTTP リクエストを実行できる任意のエージェントと連携します。
MCP ネイティブ — @lookspan/mcp TypeScript SDK は、エージェント コードを変更せずに、McpClient をラップし、MCP ツール呼び出しごとにスパンを発行します。
Python SDK — lookpan (汎用クライアント) と、LangGraph/LangChain (lookspan-langgraph) および CrewAI (lookspan-crewai) 用のアダプター。
OpenTelemetry — POST /v1/traces の OTLP/HTTP レシーバー。 Lookspan SDK を使用せずに、OTel エクスポータを指定します。 gen_ai.* 属性はプロバイダー/モデル/トークンにマップされます。
リアルタイム ストリーミング — SSE エンドポイント GET /api/stream は、span.ingested 、trace.updated、alert.triggered をダッシュボードにプッシュします。ポーリングはありません。
React ダッシュボード — ヘルス ストリップ + 行ごとのレイテンシー/コスト ミニバーを含む最近のトレース。タイムライン (ウォーターフォール) またはツリー ビューとプロンプト/応答の会話トランスクリプトで詳細を追跡します。リプレイ差分と A/B 実行比較。コストと概要 (エラー率、レイテンシ p50/p95/p99、1 日あたりのコスト)。アラート履歴。
コスト追跡 — 入力/出力/キャッシュ/推論トークンを集計し、モデルの価格表からスパンごとおよびトレースごとのcost_usdを計算します。 --pricing でオーバーライドできます。
アラート — トレースが失敗した場合、またはコスト/トークン/期間のしきい値を超えたときに通知を受け取ります (トースト + デスクトップ通知 + CLI + 永続化された履歴)。
評価スコア — LLM ジャッジ、アサーション、または手動でメトリクスをトレース ( POST /api/traces/:id/scores ) に添付します。
再生と LLM-as-judge — トレースのキャプチャされたプロムを再実行します

同じまたは異なるモデルと比較してコスト/レイテンシー/出力を比較するか、審査員モデルに応答を 0 ～ 1 で採点させます。プロバイダー キーが必要です (env、メモリ内のみ)。
データセットと実験 - プロンプトをテスト セットに収集し (トレースからシードするか手動で追加する)、セット全体をモデルに対してバッチで実行し、各出力を判定者によってスコア付けします - 実行ごとのコスト/レイテンシー/スコアを集計します。
ローカル SQLite — バージョン管理された移行。デフォルトではデータベースは ~/.lookspan/lookspan.db にあります。フラグまたは環境変数を介して構成可能。 --retention を使用したオプションの保持。
セキュリティ — デフォルトでは 127.0.0.1 にバインドされます。オプションの --token 認証;保存前に認証情報に見える属性をサーバー側で編集します。
一行 CLI — npx lookpan は、グローバル インストールを行わずにサーバーとダッシュボードを起動します。
クライアントを 1 行で囲みます。すべてのモデル呼び出しがトレースされます (OTel やプロキシはありません)。
npm インストール @lookspan/openai
'openai' から OpenAI をインポートします。
'@lookspan/openai' から {observeOpenAI } をインポートします。
const openai =observeOpenAI (new OpenAI() ) ;
オープンアイを待ってください。チャット 。完成品。 create ( { モデル : 'gpt-4o' , メッセージ } ) ;
Anthropic SDK (ドロップイン)
npm install @lookspan/anthropic
'@anthropic-ai/sdk' から Anthropic をインポートします。
import {observeAnthropic } から '@lookspan/anthropic' ;
const anthropic =observAnthropic ( new Anthropic() ) ;
人間を待ってください。メッセージ。 create ( { モデル : 'claude-sonnet-4-6' 、 max_tokens : 1024 、メッセージ } ) ;
TypeScript / MCP
npm インストール @lookspan/mcp
'@lookspan/mcp' から {wrapMcpClient , HttpSpanExporter } をインポートします。
constexporter = new HttpSpanExporter ( { エンドポイント : 'http://127.0.0.1:3100/api/ingest' } ) ;
const { クライアント } = WrapMcpClient ( mcpClient , { エクスポーター , エージェント ID : 'my-agent' } ) ;
// 以前とまったく同じように使用します。すべての callTool は、tool_call スパンを発行します。
クライアントを待ちます。 callTool ( {

名前: 'read_file' 、引数: { パス: '/tmp/foo.txt' } } ) ;
エクスポーターを待ちます。フラッシュ() ;
Python (汎用、LangGraph、CrewAI)
pip install ルックスパン # + ルックスパンラングラフ / ルックスパンクレワイ
LookspanインポートからLookspanClient
lookspan_langgraph からのインポート LookspanCallbackHandler
クライアント = LookspanClient (エンドポイント = "http://127.0.0.1:3100/api/ingest")
handler = LookspanCallbackHandler ( client = client 、agent_id = "my-agent" )
結果 = グラフ。 invoke ({ "メッセージ" : []}, config = { "コールバック" : [ハンドラー ]})
クライアント。フラッシュ（）
OpenTelemetry (SDKなし)
任意の OTel エクスポーターを標準 OTLP エンドポイントに向けます。
エクスポート OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://127.0.0.1:3100/v1/traces
# protobuf (OTel のデフォルト) と JSON は両方とも受け入れられます
その他の実行可能な例は、examples/ にあります。
ドロップイン SDK は、各通話のプロンプトと応答をキャプチャします ( CaptureContent 、
デフォルト。シークレットはサーバー側でスクラブされます)。これにより、Lookspan は
観察から改善までのループ — トレースを開いて、リプレイと判定を使用します。
パネルを使用するか、API を直接呼び出します。
# プロバイダー キーはメモリ内にのみ存在し、DB に書き込まれたりログに記録されることはありません。
LOOKSPAN_OPENAI_API_KEY=sk-... npx ルックスパン
# ...または LOOKSPAN_ANTHROPIC_API_KEY / --openai-key / --anthropic-key
# キャプチャしたプロンプトを別のモデルに対して再生し、コスト/レイテンシー/出力を比較します
curl -X POST localhost:3100/api/traces/ < id > /replay -H ' content-type: application/json ' \
-d ' {"model":"gpt-4o-mini"} ' # 元のモデルを再実行するには、「model」を省略します
# LLM ジャッジによる応答のスコア 0 ～ 1 (「llm-judge」スコアとして保存)
curl -X POST localhost:3100/api/traces/ < id > /judge -H ' content-type: application/json ' \
-d ' {"メトリック":"正確さ"} '
プロンプト/出力を Lookspan から完全に除外するには、{captureContent: false } を渡します。
OpenAIを観察する/Anthropicを観察する

— リプレイして判定し、無効のままにします。
1 つのトレースからテスト セット全体まで評価をスケールします。データセット (シード) を構築する
実際のトレースから項目を取得するか、手動で追加します)、それをモデルに対して実行します —
各アイテムはリプレイされ、オプションで審査員によって集計されて採点されます。
実行ごとのコスト/レイテンシー/スコア。ダッシュボードのデータセットですべてを管理するか、次のようにします。
# データセットを作成し、キャプチャされたトレースのプロンプトをアイテムとして追加します
DS= $(curl -s -X POST localhost:3100/api/datasets -d ' {"name":"regressions"} ' -H ' content-type: application/json ' | jq -r .dataset.id )
curl -X POST localhost:3100/api/datasets/ $DS /items/from-trace -H ' content-type: application/json ' -d ' {"traceId":"<id>"} '
# セット全体をモデルに対して実行し、各出力を判断します
curl -X POST localhost:3100/api/datasets/ $DS /run -H ' content-type: application/json ' \
-d ' {"モデル":"gpt-4o-mini","判断":true,"メトリック":"正確さ"} '
HTTP API
方法
パス
説明
ゲット
/api/健康
サービス状況
投稿
/api/インジェスト
取り込みスパン (本体: IngestPayload )
ゲット
/api/トレース
トレースのリスト (ページ分割; Framework 、 status 、 sessionId によるフィルター)
ゲット
/api/トレース/:id
すべてのスパンとスコアを含む詳細をトレースします
投稿
/api/traces/:id/スコア
評価スコアを添付します ( {名前, 値, コメント?, 出典?} )
投稿
/api/traces/:id/replay
キャプチャされたプロンプト ( {model?、provider?、spanId?} ) を再実行します。プロバイダーキーが必要です
ゲット
/api/traces/:id/replays
トレースの過去のリプレイをリストする
投稿
/api/traces/:id/judge
LLM-as-judge: プロンプト/レスポンスをスコアリングします ( {metric?、model?、provider?、rubric?} )
投稿を取得
/api/データセット
データセットの一覧表示/作成
ゲット
/api/datasets/:id
データセットの詳細 (項目 + 実行)
投稿
/api/datasets/:id/items
項目を追加します ( {input, Expected?} または {items:[…]} )
投稿
/api/datasets/:id/items/from-trace
トレースのキャプチャされたプロンプトから項目をシードする
投稿

/api/datasets/:id/run
モデルに対してセットを実行します ( {model, judge?, metric?} )。プロバイダーキーが必要です
ゲット
/api/runs/:id
実行の概要 + 項目ごとの結果
ゲット
/api/セッション
セッションのリスト (エージェント、トレース、コスト、エラー、時間範囲)
ゲット
/api/セッション/:id
セッションの概要とそのトレース（マルチエージェントのタイムライン）
ゲット
/api/コスト/概要
コストの内訳 (合計、モデル別、プロバイダー別、代理店別)
ゲット
/api/統計
統計の概要 (合計、エラー率、レイテンシ p50/p95/p99、1 日あたりのコスト)
ゲット
/api/アラート
トリガーされたアラートの履歴
ゲット
/api/ストリーム
リアルタイム SSE イベント ストリーム
投稿
/v1/トレース
OpenTelemetry OTLP/HTTP トレース レシーバー (JSON ExportTraceServiceRequest )
CLI オプション
npx ルックススパン [オプション]
-p, --port <port> リッスンするポート (デフォルト: 3100)
--host <host> バインド先のホスト (デフォルト: 127.0.0.1)
--db <パス> SQLite データベース パス (デフォルト: ~/.lookspan/lookspan.db)
--retention <dur> より古いトレースをプルーンします。 7日、24時間、30分
--token <トークン> 承認が必要です: API 上のベアラー <トークン>
--pricing <ファイル> カスタム モデルの価格表 (JSON)
--alert-error トレースが失敗した場合のアラート
--alert-cost <usd> トレースのコストが <usd> を超えた場合にアラートを送信します
--alert-tokens <n> トレースが <n> 個のトークンを超えたときに警告します
--alert-duration <ms> トレースに <ms> よりも時間がかかる場合にアラートを送信します
--open ブラウザでダッシュボードを開きます
-h, --help 彼に見せてください

[切り捨てられた]

## Original Extract

Local-first observability dashboard for AI agents. MCP-native. Look at every span your agents emit. - JoniMartin27/lookspan

GitHub - JoniMartin27/lookspan: Local-first observability dashboard for AI agents. MCP-native. Look at every span your agents emit. · GitHub
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
JoniMartin27
/
lookspan
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
58 Commits 58 Commits .github .github apps/ dashboard apps/ dashboard docs docs examples examples packages packages python python website website .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.es.md README.es.md README.md README.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Local-first observability dashboard for AI agents. MCP-native. See every span your agents emit.
npx lookspan # → http://127.0.0.1:3100
▶ Watch the full 75-second presentation
Agent (MCP · LangGraph · CrewAI · OpenTelemetry · HTTP) → POST /api/ingest → SQLite → real-time dashboard
🇪🇸 ¿Prefieres español? Lee el README en español .
When an AI agent misbehaves — fails, stalls, or quietly burns more tokens than
expected — there's no native way to see what happened step by step. Existing
observability tools are cloud-first: they want accounts, API keys, and shipping
your production data to someone else's servers.
Lookspan takes the opposite approach: everything runs on your machine, data
never leaves it, and infra cost is zero. Instrument your agent with an adapter
(or just POST JSON) and open the dashboard in your browser.
npx lookspan # → http://127.0.0.1:3100, no install, no cloud
Send your first span from any language:
curl -X POST http://127.0.0.1:3100/api/ingest \
-H " Content-Type: application/json " \
-d ' {"spans":[{"traceId":"t1","spanId":"s1","parentSpanId":null,"type":"llm_call","name":"agent.run","startedAt":"2026-06-02T10:00:00Z","endedAt":"2026-06-02T10:00:01Z","status":"ok","framework":"custom","model":"gpt-4o","provider":"openai","usage":{"inputTokens":1000,"outputTokens":500,"costUsd":0}}]} '
Open http://127.0.0.1:3100 and watch the trace appear — with its cost computed server-side.
HTTP span ingest — POST /api/ingest accepts JSON batches of spans. Works with any agent that can make an HTTP request.
MCP-native — the @lookspan/mcp TypeScript SDK wraps any McpClient and emits a span per MCP tool call, without changing your agent code.
Python SDKs — lookspan (generic client) plus adapters for LangGraph/LangChain ( lookspan-langgraph ) and CrewAI ( lookspan-crewai ).
OpenTelemetry — an OTLP/HTTP receiver at POST /v1/traces ; point any OTel exporter at it with no Lookspan SDK. gen_ai.* attributes map to provider/model/tokens.
Real-time streaming — SSE endpoint GET /api/stream pushes span.ingested , trace.updated and alert.triggered to the dashboard, no polling.
React dashboard — recent traces with a health strip + per-row latency/cost mini-bars; trace detail with a timeline (waterfall) or tree view and a conversation transcript of the prompt/response; replay diffs and A/B run comparison; costs & overview (error rate, latency p50/p95/p99, cost per day); alerts history.
Cost tracking — aggregates input/output/cached/reasoning tokens and computes cost_usd per span and per trace from a model pricing table, overridable with --pricing .
Alerts — get notified when a trace fails or exceeds a cost/token/duration threshold (toast + desktop notification + CLI + persisted history).
Evaluation scores — attach metrics to a trace ( POST /api/traces/:id/scores ) from an LLM judge, an assertion, or by hand.
Replay & LLM-as-judge — re-run a trace's captured prompt against the same or a different model and diff cost/latency/output, or have a judge model score the response 0–1. Needs a provider key (env, in-memory only).
Datasets & experiments — collect prompts into a test set (seed from a trace or add by hand), run the whole set against a model in batch and score each output with the judge — aggregate cost/latency/score per run.
Local SQLite — versioned migrations. Database at ~/.lookspan/lookspan.db by default; configurable via flag or env var. Optional retention with --retention .
Security — binds to 127.0.0.1 by default; optional --token auth; server-side redaction of credential-looking attributes before storage.
One-line CLI — npx lookspan starts the server and the dashboard with no global install.
Wrap your client in one line — every model call is traced (no OTel, no proxy):
npm install @lookspan/openai
import OpenAI from 'openai' ;
import { observeOpenAI } from '@lookspan/openai' ;
const openai = observeOpenAI ( new OpenAI ( ) ) ;
await openai . chat . completions . create ( { model : 'gpt-4o' , messages } ) ;
Anthropic SDK (drop-in)
npm install @lookspan/anthropic
import Anthropic from '@anthropic-ai/sdk' ;
import { observeAnthropic } from '@lookspan/anthropic' ;
const anthropic = observeAnthropic ( new Anthropic ( ) ) ;
await anthropic . messages . create ( { model : 'claude-sonnet-4-6' , max_tokens : 1024 , messages } ) ;
TypeScript / MCP
npm install @lookspan/mcp
import { wrapMcpClient , HttpSpanExporter } from '@lookspan/mcp' ;
const exporter = new HttpSpanExporter ( { endpoint : 'http://127.0.0.1:3100/api/ingest' } ) ;
const { client } = wrapMcpClient ( mcpClient , { exporter , agentId : 'my-agent' } ) ;
// Use it exactly as before — every callTool emits a tool_call span.
await client . callTool ( { name : 'read_file' , arguments : { path : '/tmp/foo.txt' } } ) ;
await exporter . flush ( ) ;
Python (generic, LangGraph, CrewAI)
pip install lookspan # + lookspan-langgraph / lookspan-crewai
from lookspan import LookspanClient
from lookspan_langgraph import LookspanCallbackHandler
client = LookspanClient ( endpoint = "http://127.0.0.1:3100/api/ingest" )
handler = LookspanCallbackHandler ( client = client , agent_id = "my-agent" )
result = graph . invoke ({ "messages" : []}, config = { "callbacks" : [ handler ]})
client . flush ()
OpenTelemetry (no SDK)
Point any OTel exporter at the standard OTLP endpoint:
export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://127.0.0.1:3100/v1/traces
# protobuf (the OTel default) and JSON are both accepted
More runnable examples in examples/ .
The drop-in SDKs capture each call's prompt and reply ( captureContent , on by
default; secrets are scrubbed server-side). With that, Lookspan can close the
loop from observe to improve — open a trace and use the Replay & judge
panel, or call the API directly:
# Provider keys live in memory only — never written to the DB or logged.
LOOKSPAN_OPENAI_API_KEY=sk-... npx lookspan
# ...or LOOKSPAN_ANTHROPIC_API_KEY / --openai-key / --anthropic-key
# Replay the captured prompt against another model and diff cost/latency/output
curl -X POST localhost:3100/api/traces/ < id > /replay -H ' content-type: application/json ' \
-d ' {"model":"gpt-4o-mini"} ' # omit "model" to re-run the original
# Score the response 0–1 with an LLM judge (stored as an "llm-judge" score)
curl -X POST localhost:3100/api/traces/ < id > /judge -H ' content-type: application/json ' \
-d ' {"metric":"correctness"} '
To keep prompts/outputs out of Lookspan entirely, pass { captureContent: false }
to observeOpenAI / observeAnthropic — replay & judge then stay disabled.
Scale evaluation from one trace to a whole test set. Build a dataset (seed
items from real traces or add them by hand), then run it against a model —
each item is replayed and, optionally, scored by the judge, with aggregate
cost/latency/score per run. Manage it all under Datasets in the dashboard, or:
# Create a dataset and add the captured prompt of a trace as an item
DS= $( curl -s -X POST localhost:3100/api/datasets -d ' {"name":"regressions"} ' -H ' content-type: application/json ' | jq -r .dataset.id )
curl -X POST localhost:3100/api/datasets/ $DS /items/from-trace -H ' content-type: application/json ' -d ' {"traceId":"<id>"} '
# Run the whole set against a model, judging each output
curl -X POST localhost:3100/api/datasets/ $DS /run -H ' content-type: application/json ' \
-d ' {"model":"gpt-4o-mini","judge":true,"metric":"correctness"} '
HTTP API
Method
Path
Description
GET
/api/health
Service status
POST
/api/ingest
Ingest spans (body: IngestPayload )
GET
/api/traces
List traces (paginated; filter by framework , status , sessionId )
GET
/api/traces/:id
Trace detail with all its spans and scores
POST
/api/traces/:id/scores
Attach an evaluation score ( {name, value, comment?, source?} )
POST
/api/traces/:id/replay
Re-run the captured prompt ( {model?, provider?, spanId?} ); needs a provider key
GET
/api/traces/:id/replays
List past replays for the trace
POST
/api/traces/:id/judge
LLM-as-judge: score the prompt/response ( {metric?, model?, provider?, rubric?} )
GET POST
/api/datasets
List / create datasets
GET
/api/datasets/:id
Dataset detail (items + runs)
POST
/api/datasets/:id/items
Add item(s) ( {input, expected?} or {items:[…]} )
POST
/api/datasets/:id/items/from-trace
Seed an item from a trace's captured prompt
POST
/api/datasets/:id/run
Run the set against a model ( {model, judge?, metric?} ); needs a provider key
GET
/api/runs/:id
Run summary + per-item results
GET
/api/sessions
List sessions (agents, traces, cost, errors, time range)
GET
/api/sessions/:id
Session summary + its traces (multi-agent timeline)
GET
/api/costs/summary
Cost breakdown (total, by model, provider, agent)
GET
/api/stats
Stats summary (totals, error rate, latency p50/p95/p99, cost per day)
GET
/api/alerts
History of triggered alerts
GET
/api/stream
Real-time SSE event stream
POST
/v1/traces
OpenTelemetry OTLP/HTTP trace receiver (JSON ExportTraceServiceRequest )
CLI options
npx lookspan [options]
-p, --port <port> Port to listen on (default: 3100)
--host <host> Host to bind to (default: 127.0.0.1)
--db <path> SQLite database path (default: ~/.lookspan/lookspan.db)
--retention <dur> Prune traces older than e.g. 7d, 24h, 30m
--token <token> Require Authorization: Bearer <token> on the API
--pricing <file> Custom model pricing table (JSON)
--alert-error Alert when a trace fails
--alert-cost <usd> Alert when a trace costs more than <usd>
--alert-tokens <n> Alert when a trace exceeds <n> tokens
--alert-duration <ms> Alert when a trace takes longer than <ms>
--open Open the dashboard in your browser
-h, --help Show he

[truncated]

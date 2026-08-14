---
source: "https://weavescope.com/"
hn_url: "https://news.ycombinator.com/item?id=49305100"
title: "Show HN: WeaveScope – Elixir native observability for AI agents"
article_title: ""
author: "caudena"
captured_at: "2026-08-14T22:12:53Z"
capture_tool: "hn-digest"
hn_id: 49305100
score: 1
comments: 0
posted_at: "2026-08-14T21:58:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: WeaveScope – Elixir native observability for AI agents

- HN: [49305100](https://news.ycombinator.com/item?id=49305100)
- Source: [weavescope.com](https://weavescope.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T21:58:16Z

## Translation

タイトル: Show HN: WeaveScope – AI エージェント向けの Elixir ネイティブの可観測性
説明: Elixir と BEAM で構築された AI エージェントをトレース、デバッグ、監視します。失敗したツール呼び出しを見つけてすべてのステップを検査し、運用上の問題をより迅速に解決します。

記事本文:
コンテンツにスキップ
ナビゲーションを開く
プラットフォーム
インストール
価格設定
ブログ
ドキュメント
サインイン
無料で始める
WeaveScope は早期アクセス中です - Elixir エージェントのトレースと監視
Elixir エージェントの実行の失敗、低速、またはコストの背後にあるモデルまたはツールの呼び出しを追跡します。
生産トレースレビューを予約する
1 か月あたり最大 100,000 件のイベントを無料で受け入れ可能 · クレジット カード不要 · 1 つの Elixir 依存関係
モデル呼び出し、ツール呼び出し、およびインストルメント化されたエージェント実行全体のスパンを検査します。仕組み →
フィルター
過去 30 日間
すべて
エラー
遅い > 30 秒
トレースID、ユーザー、モデル...でフィルタリングします。
25/25
ページ上
リセット
所見 事前なし
43,040
分析ユニット
エラー率 事前なし
0.00%
エラーのあるトレース
p95 レイテンシー 事前なし
29.56秒
トレース参照
トークン事前なし
5,093,888
入力+出力
コスト 事前なし
11.046970ドル
トレース参照
最近のエージェントの追跡。詳細プレビューの行を選択します。
トレースID
名前
ユーザー
モデル
スパン
タグ
トークン
コスト
合計時間
いつ
019fd6f2...0ea0 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 5,646 $0.0126 21.79 秒 2026-08-06 12:01
019fd6f2...9bcd x_signal_desk.post_summary 匿名 zai:glm-5.2 9 5,426 $0.0120 17.71 秒 2026-08-06 12:01
019fd6d7...60c2 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 4,365 $0.0091 13.61 秒 2026-08-06 11:31
019fd684...4fab x_signal_desk.post_summary 匿名 zai:glm-5.2 9 7,029 $0.0173 32.78 s 2026-08-06 10:01
019fd684...fe35 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 4,528 $0.0093 18.17 s 2026-08-06 10:01
019fd64d...54d5 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 7,117 $0.0179 40.66 s 2026-08-06 09:01
019fd64d...93d5 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 4,601 $0.0100 23.68 s 2026-08-06 09:01
019fd632...9e98 x_signal_desk.post_summary 匿名 zai:glm-5.2 9 5,593 $0.0120 30.45 秒 2026-08-06 08:31
成功
エージェント
x_signal_desk.post_summary
2026-08-06 12:01
21.

79秒
5,646
$0.01
@匿名
なし
ターン
木
滝
観測結果の検索
⌁ x_signal_desk.post_summary
21.79秒
⌁ PatchToolCallsMiddleware.wrap_model_call
21.79秒
⌁ DynamicPromptMiddleware.wrap_model_call
21.79秒
⌁ StructuredOutputRetryMiddleware.wrap_model_call
21.79秒
⌁ ModelRetryMiddleware.wrap_model_call
21.79秒
⌁ ランタイムZaiModelMiddleware.wrap_model_call
21.79秒
⌁ PostsummaryValidationMiddleware.wrap_model_call
21.79秒
⌁ 実行可能なシーケンス
21.79秒
• ザイ:glm-5.2 glm-5.2
▤ 5.6k 21.77 秒 $0.012589
世代
ザイ:glm-5.2
019fd6f2...1825
21.77 秒 5,646 トークン $0.012589
ザイ:glm-5.2
なし
メッセージ
ツール
詳細
ツール 0 が呼び出されました / 0 が利用可能です
この観察ツリーではツール呼び出しは記録されませんでした。
2ブロック入力
システム命令 テキスト マークダウン XML JSON
ユーザー 019fd6f2...cabd テキスト マークダウン XML JSON
推理1ブロック
推論テキスト マークダウン XML JSON
出力1ブロック
AI 20260806...44d0 テキスト マークダウン XML JSON
{
「投稿」: [
{
"id" : 11301 、
"key_facts" : [
「Uniswapは、ロビンフッドチェーンの新しい出発点であるTradePoolsを立ち上げました」、
"pools.trade でアクセス可能なプラットフォーム",
「コメンテーターは、すでに 170 を超えるランチパッドが存在し、詐欺トークンの懸念を引き起こしていると指摘しています。」
】
}
】
}
ツール 0 が呼び出されました / 0 が利用可能です
この観察ツリーではツール呼び出しは記録されませんでした。
エージェントが実際にどのように実行されるかを考慮して構築
何が起こったのか知ってください。
次に、何を修正すべきかを知ってください。
1 つのタイムラインでのモデルとツールの呼び出し
BeamWeaver は、インストルメント化されたモデル呼び出し、ツール呼び出しを接続し、エージェントの実行全体にまたがります。失敗したステップからその入力、出力、レイテンシ、コストにジャンプします。
重要な動作信号
エージェントの実行時のスループット、遅延、エラー、トークン使用量、コストを追跡します。あらゆる成果の背後にある作業についてチームに共有ビューを提供します。
エージェントの障害が広がる前にシグナルを取得する
次はアラートです。 trの上に構築しています

別個のブラック ボックスとしてではなく、チームがすでに持っているデータを確認および監視します。
依存関係が 1 つあります。コンテキストをトレースします。
エージェントを計測する
昼食前に。
依存関係を追加し、1 つのエンドポイントを構成します。 BeamWeaver は、インストルメント化されたエージェントの実行、モデル呼び出し、およびツール呼び出しをバックグラウンドでエクスポートします。
mix.exs
ランタイム.exs
エージェント.ex
デフップデプスはやります
[
{ :beam_weaver , "~> 0.1.15" }
】
終わり
設定:beam_weaver 、
openai: [api_key: System.fetch_env!( "OPENAI_API_KEY" )],
織りスコープ: [
api_key: System.fetch_env!( "WEAVESCOPE_API_KEY" ),
エンドポイント: 「https://app.weavescope.com」
】
defmodule Support.Agent do
エイリアス BeamWeaver.Agent
エイリアス BeamWeaver.Core.Message
def run (クエリ、コンテキスト \\ %{} ) do
モデル =
BeamWeaver.Models 。 init_chat_model! ( "openai:gpt-5.4-mini" ,
温度：0.2、
タイムアウト: 30_000
)
{:わかりました、エージェント} =
エージェント。ビルド (
名前: "support.reply" 、
モデル：モデル、
system_prompt : "サポートの質問に明確に答えてください。"
)
エージェント。 invoke (agent, %{messages : [ Message . user (query)] } ,
トレース: [
名前: "support.reply" 、
user_id : コンテキスト[ :user_id ],
実行モード: "サポート応答" 、
フィールド: %{ ticket_id : context[ :ticket_id ]、account_id : context[ :account_id ] }
】
)
終わり
終わり
価格設定
無料で始めましょう。受け入れられた観察イベントの使用数をカウントします。トレースには複数のイベントを含めることができます。
エージェントはさらに多くのことを行っています。
結果の背後にある取り組みをご覧ください。
次の決定を下すために必要なトレースと動作信号から始めます。
BEAM 上に構築されたエージェントのトレースと監視。

## Original Extract

Trace, debug, and monitor AI agents built with Elixir and the BEAM. Find failed tool calls, inspect every step, and resolve production issues faster.

Skip to content
Open navigation
Platform
Install
Pricing
Blog
Docs
Sign in
Start free
WeaveScope is in early access — tracing and monitoring for Elixir agents
Trace the model or tool call behind failed, slow, or expensive Elixir agent runs.
Book a production trace review
Free up to 100K accepted events / month · no credit card · one Elixir dependency
Inspect model calls, tool calls, and spans across instrumented agent runs. How it works →
Filters
Last 30 days
All
Errors
Slow > 30s
Filter by trace ID, user, model...
25 of 25
on page
Reset
Observations No prior
43,040
Analytics unit
Error rate No prior
0.00 %
Traces with errors
p95 latency No prior
29.56 s
Trace refs
Tokens No prior
5,093,888
Input + output
Cost No prior
$11.046970
Trace refs
Recent agent traces. Select a row for a detail preview.
Trace ID
Name
User
Model
Spans
Tags
Tokens
Cost
Total time
When
019fd6f2...0ea0 x_signal_desk.post_summary anonymous zai:glm-5.2 9 5,646 $0.0126 21.79 s 2026-08-06 12:01
019fd6f2...9bcd x_signal_desk.post_summary anonymous zai:glm-5.2 9 5,426 $0.0120 17.71 s 2026-08-06 12:01
019fd6d7...60c2 x_signal_desk.post_summary anonymous zai:glm-5.2 9 4,365 $0.0091 13.61 s 2026-08-06 11:31
019fd684...4fab x_signal_desk.post_summary anonymous zai:glm-5.2 9 7,029 $0.0173 32.78 s 2026-08-06 10:01
019fd684...fe35 x_signal_desk.post_summary anonymous zai:glm-5.2 9 4,528 $0.0093 18.17 s 2026-08-06 10:01
019fd64d...54d5 x_signal_desk.post_summary anonymous zai:glm-5.2 9 7,117 $0.0179 40.66 s 2026-08-06 09:01
019fd64d...93d5 x_signal_desk.post_summary anonymous zai:glm-5.2 9 4,601 $0.0100 23.68 s 2026-08-06 09:01
019fd632...9e98 x_signal_desk.post_summary anonymous zai:glm-5.2 9 5,593 $0.0120 30.45 s 2026-08-06 08:31
success
agent
x_signal_desk.post_summary
2026-08-06 12:01
21.79 s
5,646
$0.01
@ anonymous
none
Turns
Tree
Waterfall
Search observations
⌁ x_signal_desk.post_summary
21.79 s
⌁ PatchToolCallsMiddleware.wrap_model_call
21.79 s
⌁ DynamicPromptMiddleware.wrap_model_call
21.79 s
⌁ StructuredOutputRetryMiddleware.wrap_model_call
21.79 s
⌁ ModelRetryMiddleware.wrap_model_call
21.79 s
⌁ RuntimeZaiModelMiddleware.wrap_model_call
21.79 s
⌁ PostSummaryValidationMiddleware.wrap_model_call
21.79 s
⌁ RunnableSequence
21.79 s
• zai:glm-5.2 glm-5.2
▤ 5.6k 21.77 s $0.012589
generation
zai:glm-5.2
019fd6f2...1825
21.77 s 5,646 tokens $0.012589
zai:glm-5.2
none
Messages
Tools
Details
Tools 0 called / 0 available
No tool calls were recorded for this observation tree.
Input 2 blocks
System instructions Text Markdown XML JSON
User 019fd6f2...cabd Text Markdown XML JSON
Reasoning 1 blocks
Reasoning Text Markdown XML JSON
Output 1 blocks
AI 20260806...44d0 Text Markdown XML JSON
{
"posts" : [
{
"id" : 11301 ,
"key_facts" : [
"Uniswap launched TradePools, a new launchpad on Robinhood Chain",
"Platform accessible at pools.trade",
"Commentator notes over 170 launchpads already exist, raising scam token concerns"
]
}
]
}
Tools 0 called / 0 available
No tool calls were recorded for this observation tree.
Built for how agents actually run
Know what happened.
Then know what to fix.
Model and tool calls in one timeline
BeamWeaver connects instrumented model calls, tool calls, and spans across an agent run. Jump from a failed step to its inputs, outputs, latency, and cost.
The operational signals that matter
Follow throughput, latency, errors, token usage, and cost as your agents run. Give your team a shared view of the work behind every outcome.
Get a signal before an agent failure spreads
Alerting is next. We are building it on top of the tracing and monitoring data your team already has—not as a separate black box.
One dependency. Trace context.
Instrument an agent
before lunch.
Add the dependency and configure one endpoint. BeamWeaver exports instrumented agent runs, model calls, and tool calls in the background.
mix.exs
runtime.exs
agent.ex
defp deps do
[
{ :beam_weaver , "~> 0.1.15" }
]
end
config :beam_weaver ,
openai: [api_key: System.fetch_env!( "OPENAI_API_KEY" )],
weave_scope: [
api_key: System.fetch_env!( "WEAVESCOPE_API_KEY" ),
endpoint: "https://app.weavescope.com"
]
defmodule Support.Agent do
alias BeamWeaver.Agent
alias BeamWeaver.Core.Message
def run (query, context \\ %{} ) do
model =
BeamWeaver.Models . init_chat_model! ( "openai:gpt-5.4-mini" ,
temperature : 0.2 ,
timeout : 30_000
)
{:ok, agent} =
Agent . build (
name : "support.reply" ,
model : model,
system_prompt : "Answer support questions clearly."
)
Agent . invoke (agent, %{ messages : [ Message . user (query)] } ,
trace : [
name : "support.reply" ,
user_id : context[ :user_id ],
execution_mode : "support_reply" ,
fields : %{ ticket_id : context[ :ticket_id ], account_id : context[ :account_id ] }
]
)
end
end
Pricing
Start free. Usage counts accepted observation events; a trace can contain multiple events.
Your agents are doing more.
See the work behind the outcome.
Start with the traces and operating signals you need to make the next decision.
Tracing and monitoring for agents built on the BEAM.

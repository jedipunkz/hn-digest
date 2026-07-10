---
source: "https://drop-05a4352b-803.sophisticated-stay.workers.dev"
hn_url: "https://news.ycombinator.com/item?id=48854997"
title: "OpenAI's API can now keep reasoning across turns instead of discarding it"
article_title: "OpenAI “Reasoning models” docs — what changed (Jun 20 → Jul 10, 2026)"
author: "CjHuber"
captured_at: "2026-07-10T02:48:56Z"
capture_tool: "hn-digest"
hn_id: 48854997
score: 3
comments: 0
posted_at: "2026-07-10T02:19:07Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's API can now keep reasoning across turns instead of discarding it

- HN: [48854997](https://news.ycombinator.com/item?id=48854997)
- Source: [drop-05a4352b-803.sophisticated-stay.workers.dev](https://drop-05a4352b-803.sophisticated-stay.workers.dev)
- Score: 3
- Comments: 0
- Posted: 2026-07-10T02:19:07Z

## Translation

タイトル: OpenAI の API は、ターンをまたいで推論を破棄するのではなく継続できるようになりました
記事タイトル: OpenAI「推論モデル」ドキュメント — 何が変わったのか (2026 年 6 月 20 日 → 7 月 10 日)
説明: OpenAI の非公式テキストの差分

記事本文:
OpenAI「推論モデル」ドキュメント — 変更点 (2026 年 6 月 20 日 → 7 月 10 日)
非公式テキストの差分 — ウェイバック スナップショット、2026 年 6 月 20 日とライブ ページ、2026 年 7 月 10 日 · 緑 = 追加 赤 = 削除未変更展開 · OpenAI とは提携していません
ChatGPT ホーム API ドキュメント OpenAI API のガイドと概念 API リファレンス エンドポイント、パラメーター、応答 Codex ドキュメント Codex のガイド、概念、製品ドキュメント ユースケース チームが ChatGPT または Codex を使用して実行できるワークフローとタスクの例 ドキュメント ユースケース リソース ChatGPT アプリ SDK ChatGPT を拡張するアプリの構築 ワークスペース エージェント 公開された ChatGPT ワークスペース エージェントのトリガー Commerce ChatGPT 広告でのコマース フローの構築ChatGPT での広告の公開と測定 リソース ショーケース インスピレーションを得るためのデモ アプリ ブログ 開発者からの学習と経験 クックブック OpenAI モデルを使用して構築するためのノートブックの例 OpenAI を使用して構築するためのドキュメント、ビデオ、およびデモ アプリを学ぶ コミュニティ プログラム、ミートアップ、およびビルダーのサポート API ダッシュボード ChatGPT を試す API ドキュメントを検索
統合と可観測性
MCP とコネクタによる安全な MCP トンネル
ファイルの検索と取得 ファイルの検索
文字起こし リアルタイム文字起こし
リアルタイムセッション 会話の管理
Webhook とサーバー側コントロール
Workload Identity フェデレーションの概要
最適化サイクルの微調整
直接的なプリファレンスの最適化
アシスタント API 移行ガイド
アプリ SDK ワークスペース エージェント コマース広告 ドキュメント セクション 選択... ホーム
Workspace Agent アクセス トークンを使用して認証する
Codex Security プラグインのクイックスタート
ワークスペースのアクセス、ポリシー、モデル
役割とワークスペースの権限
使用法、ガバナンス、コンプライアンス
コンプライアンスAPIと監査イベント
導入およびモデルプロバイダー
携帯電話からリモートエンジニアリング作業をマスターする
プライベート MCP サーバーを公開せずにアクセス可能にする

ic
Perplexity がリアルタイム API を使用して音声検索を何百万人ものユーザーに提供した方法
GPT-5.4 を使用した楽しいフロントエンドの設計
プロンプトから製品まで: 1 年間の対応
統合と可観測性
MCP とコネクタによる安全な MCP トンネル
ファイルの検索と取得 ファイルの検索
文字起こし リアルタイム文字起こし
リアルタイムセッション 会話の管理
Webhook とサーバー側コントロール
Workload Identity フェデレーションの概要
最適化サイクルの微調整
直接的なプリファレンスの最適化
アシスタント API 移行ガイド
推論モデルがどのように機能するか、そしてそれを上手に使用する方法を学びましょう。
GPT-5.5 のような推論モデルは、応答を生成する前に内部推論トークンを使用します。これは、モデルの計画、ツールの効果的な使用、代替案の検査、曖昧さからの回復、およびより困難な複数ステップのタスクの解決に役立ちます。推論モデルは、複雑な問題解決、コーディング、科学的推論、および複数ステップのエージェント ワークフローに特に適しています。これらは、軽量コーディング エージェントである Codex CLI の最適なモデルでもあります。
gpt-5.6 から開始する ほとんどの推論ワークロードでは、gpt-5.5 から開始します。より長い遅延を許容できる、より困難な問題に対応する最高のインテリジェンス API オプションが必要な場合は、 gpt-5.5-pro を使用してください。コストを低くするには gpt-5.4 を、コストと遅延を低くするには gpt-5.4-mini を検討してください。
推論モデルは応答でより適切に機能します
API 。チャット完了 API の間
は引き続きサポートされており、モデル インテリジェンスとパフォーマンスが向上します。
応答を使用します。
Responses API を呼び出して、推論モデルと推論作業を指定します。
1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
「openai」から OpenAI をインポートします。
const openai = new OpenAI();
const プロンプト = `
文字列として表される行列を受け取る bash スクリプトを作成します。
'[1,2],[3,4],[5,6]' をフォーマットし、転置 i を出力します。

同じフォーマット。
` ;
const response = await openai.responses.create({
モデル：「gpt-5.5」、
モデル：「gpt-5.6」、
推論: { 努力: "低い" }、
入力: [
{
役割:「ユーザー」、
内容 : プロンプト、
}、
]、
});
コンソール .log(response.output_text); 1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
openaiインポートからOpenAI
クライアント = OpenAI()
プロンプト = """
文字列として表される行列を受け取る bash スクリプトを作成します。
'[1,2],[3,4],[5,6]' をフォーマットし、同じフォーマットで転置を出力します。
「」
応答 = client.responses.create(
モデル= "gpt-5.5" 、
モデル= "gpt-5.6" 、
推論={ "努力" : "低" },
入力 =[
{
"ロール" : "ユーザー" 、
"コンテンツ" : プロンプト
}
】
)
print (response.output_text) 1
2
3
4
-
5
6
7
8
9
10
11
12
13
カール https://api.openai.com/v1/responses \
-H "コンテンツ タイプ: application/json" \
-H "認可: ベアラー $OPENAI_API_KEY " \
-d '{
"モデル": "gpt-5.5",
"モデル": "gpt-5.6",
"推論": {"努力": "低"},
「入力」: [
{
"ロール": "ユーザー",
"content": "形式 \"[1,2],[3,4],[5,6]\" の文字列として表される行列を受け取り、同じ形式で転置結果を出力する bash スクリプトを作成します。"
}
】
}' ··· 4 つの変更されていないブロック (推論の取り組み) — クリックすると推論の取り組みが表示されます
reasoning.effort パラメーターは、タスクを実行するときにどれだけ考えるかについてモデルをガイドします。
サポートされる値はモデルに依存しており、 none 、 minimum 、 low 、 middle 、 high 、および xhigh を含めることができます。労力が低いほど速度とトークンの使用量が少なくなりますが、労力が高い場合、モデルはより完全に考えて高品質の応答を提供します。また、モデルは推論作業全体にわたって適応的に推論し、単純なタスクにはより少ないトークンを使用し、複雑なタスクにはより深く考えます。
デフォルトも普遍的ではなくモデルに依存します。 gpt-5.5 のデフォルトの推論労力は中程度です。これが最良の出発点です

gpt-5.5 の品質、信頼性、パフォーマンスの完全なバランスを実現します。
レイテンシの影響を受けやすいアプリケーションで最初にトークンが表示されるまでの時間を短縮するには、より深い推論を続行する前に、モデルに短いプリアンブルを生成するように依頼します。
一部のモデルではこれらの値のサブセットのみがサポートされているため、設定を選択する前に関連するモデルのページを確認してください。
GPT-5.6 モデルは、Responses API の標準およびプロ推論モードをサポートします。標準がデフォルトです。より多くのモデル作業が必要で、より高いレイテンシーとトークン使用量を許容できる難しいタスクの場合は、reasoning.mode を pro に設定します。
推論モードと推論作業は独立しています。 Mode は標準実行またはプロ実行を選択し、reasoning.effort はそのモード内でモデルがどの程度の推論を適用するかを制御します。 reasoning.effort を省略した場合、GPT-5.6 は両方のモードでデフォルトで中になります。
1
2
3
4
5
6
7
8
9
10
11
カール https://api.openai.com/v1/responses \
-H "コンテンツ タイプ: application/json" \
-H "認可: ベアラー $OPENAI_API_KEY " \
-d '{
"モデル": "gpt-5.6",
「推論」: {
"モード": "プロ"、
「努力」：「中」
}、
"input": "このデータベース移行計画を確認し、潜在的な障害モードを特定します。"
プロ モードは、最終的な答えを生成するために実行されたモデル作業を集計し、選択したモデルの標準トークン レートでそれらのトークンを請求します。プロ モードは標準モードよりも多くのモデル作業を実行するため、トークンの使用量とコストが増加します。既存の Pro モデル ID は、現在の動作と価格を維持します。
推論モデルでは、入力トークンと出力トークンに加えて推論トークンが導入されます。モデルはこれらの推論トークンを使用して「思考」し、プロンプトを分析し、応答を生成するための複数のアプローチを検討します。 gpt-5.5 や gpt-5.4 などの推論モデルはインターリーブ思考をサポートしており、モデルは前後に目に見える出力トークンを生成できます。

思考の合間にも考えることができ、ツールの呼び出しの合間にも考えることができます。
以下は、ユーザーとアシスタントの間の複数ステップの会話の例です。各ステップの入力トークンと出力トークンは引き継がれますが、推論トークンは破棄されます。
ユーザーとアシスタントの間の複数ステップの会話のデフォルトの動作は次のとおりです。各ステップの入力トークンと出力トークンは引き継がれますが、以前のターンからの推論は次のサンプルにはレンダリングされません。永続的な推論をサポートするモデルは、reasoning.context を使用してこの動作を変更できます。
推論トークンは API 経由では表示されませんが、それでもスペースを占有します。
モデルのコンテキスト ウィンドウであり、出力として課金されます
トークン 。
応答を作成するときに、コンテキスト ウィンドウにトークンを推論するための十分なスペースがあることを確認することが重要です。問題の複雑さに応じて、モデルは数百から数万の推論トークンを生成する場合があります。使用された推論トークンの正確な数は、output_tokens_details の下の応答オブジェクトの使用法オブジェクトに表示されます。
1
2
3
4
5
6
7
8
9
10
11
12
13
{
「使用法」: {
"input_tokens" : 75 、
"input_tokens_details" : {
「キャッシュされたトークン」: 0
}、
"出力トークン" : 1186 、
"output_tokens_details" : {
"reasoning_tokens" : 1024
}、
"total_tokens" : 1261
}
コンテキスト ウィンドウの長さはモデル リファレンス ページに記載されており、モデル スナップショットによって異なります。
推論モデルを使用してコストを管理するには、トークンの総数を制限できます。
モデルが生成するものには、推論トークン、可視出力トークン、非可視出力トークンが含まれます。
を使用してトークンをフォーマットする
max_output_tokens
パラメータ。生成されたトークンが使用量と出力制限にどのように反映されるかについて詳しくは、出力トークン数を参照してください。
推論のためのスペースの割り当て
生成されたトークンが t に達した場合

コンテキスト ウィンドウの制限または設定した max_output_tokens 値を使用すると、ステータスが incomplete および incomplete_details で理由が max_output_tokens に設定された応答を受け取ります。これは、目に見える出力トークンが生成される前に発生する可能性があります。つまり、目に見える応答を受信せずに入力トークンと推論トークンのコストが発生する可能性があります。
これを防ぐには、コンテキスト ウィンドウに十分なスペースがあることを確認するか、max_output_tokens 値をより大きな値に調整します。 OpenAI は、これらのモデルの実験を開始するときに、推論と出力のために少なくとも 25,000 トークンを予約することをお勧めします。プロンプトに必要な推論トークンの数に慣れてきたら、それに応じてこのバッファーを調整できます。
1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
「openai」から OpenAI をインポートします。
const openai = new OpenAI();
const プロンプト = `
文字列として表される行列を受け取る bash スクリプトを作成します。
'[1,2],[3,4],[5,6]' をフォーマットし、同じフォーマットで転置を出力します。
` ;
const response = await openai.responses.create({
モデル：「gpt-5.5」、
モデル：「gpt-5.6」、
推論: { 努力: "中" }、
入力: [
{
役割:「ユーザー」、
内容 : プロンプト、
}、
]、
max_output_tokens : 300 、
});
もし(
応答.ステータス === "不完全" &&
response.incomplete_details.reason === "max_output_tokens"
) {
console .log( "トークンが不足しました" );
if (response.output_text?.length > 0 ) {
console .log( "部分的な出力:" , response.output_text);
} それ以外の場合は {
console .log( "推論中にトークンが不足しました" );
}
} 1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
openaiインポートからOpenAI
クライアント = OpenAI()
プロンプト = """
文字列として表される行列を受け取る bash スクリプトを作成します。
'[1,2],[3,4],[5,6]' をフォーマットし、転置を出力します

同じ形式で。
「」
応答 = client.responses.create(
モデル= "gpt-5.5" 、
モデル= "gpt-5.6" 、
推論={ "努力" : "中" },
入力 =[
{
"ロール" : "ユーザー" 、
"コンテンツ" : プロンプト
}
]、
max_output_tokens= 300 、
)
response.status == "incomplete" かつ response.incomplete_details.reason == "max_output_tokens" の場合:
print ( "トークンが不足しました" )
応答の場合: 出力テキスト:
print ( "部分的な出力:" , response.output_text)
それ以外の場合:
print ( "推論中にトークンが不足しました" ) ··· 5 つの変更されていないブロック (推論項目をコンテキスト内に保持) — クリックすると、推論項目をコンテキスト内に保持することが表示されます
Responses API の推論モデルを使用して関数呼び出しを行う場合は、(関数の出力に加えて) 最後の関数呼び出しで返された推論項目を返すことを強くお勧めします。モデルが複数の関数を連続して呼び出す場合は、最後のユーザー メッセージ以降のすべての推論項目、関数呼び出し項目、および関数呼び出し出力項目を返す必要があります。これにより、モデルは推論プロセスを継続して、最もトークン効率の高い方法でより良い結果を生み出すことができます。
これを行う最も簡単な方法は、前の応答からのすべての推論項目を次の応答に渡すことです。私たちのシステムは、次のような推論項目を賢く無視します。

[切り捨てられた]

## Original Extract

Unofficial text diff of OpenAI

OpenAI “Reasoning models” docs — what changed (Jun 20 → Jul 10, 2026)
Unofficial text diff — Wayback snapshot, June 20 2026 vs. the live page , July 10 2026 · green = added red = removed expand unchanged · Not affiliated with OpenAI
ChatGPT Home API Docs Guides and concepts for the OpenAI API API reference Endpoints, parameters, and responses Codex Docs Guides, concepts, and product docs for Codex Use cases Example workflows and tasks teams can take on with ChatGPT or Codex Docs Use cases Resources ChatGPT Apps SDK Build apps to extend ChatGPT Workspace Agents Trigger published ChatGPT workspace agents Commerce Build commerce flows in ChatGPT Ads Publish and measure ads in ChatGPT Resources Showcase Demo apps to get inspired Blog Learnings and experiences from developers Cookbook Notebook examples for building with OpenAI models Learn Docs, videos, and demo apps for building with OpenAI Community Programs, meetups, and support for builders API Dashboard Try ChatGPT Search the API docs
Integrations and observability
MCP and Connectors Secure MCP Tunnel
File search and retrieval File search
Transcription Realtime transcription
Realtime sessions Managing conversations
Webhooks and server-side controls
Workload identity federation Overview
Fine-tuning Optimization cycle
Direct preference optimization
Assistants API Migration guide
Apps SDK Workspace Agents Commerce Ads Docs section Select... Home
Authenticate with Workspace Agent access tokens
Codex Security plugin Quickstart
Workspace access, policy, and models
Roles and workspace permissions
Usage, governance, and compliance
Compliance API and audit events
Deployment and model providers
Mastering remote engineering work from your phone
Making private MCP servers reachable without making them public
How Perplexity Brought Voice Search to Millions Using the Realtime API
Designing delightful frontends with GPT-5.4
From prompts to products: One year of Responses
Integrations and observability
MCP and Connectors Secure MCP Tunnel
File search and retrieval File search
Transcription Realtime transcription
Realtime sessions Managing conversations
Webhooks and server-side controls
Workload identity federation Overview
Fine-tuning Optimization cycle
Direct preference optimization
Assistants API Migration guide
Learn how reasoning models work and how to use them well.
Reasoning models like GPT-5.5 use internal reasoning tokens before producing a response. This helps the model plan, use tools effectively, inspect alternatives, recover from ambiguity, and solve harder multi-step tasks. Reasoning models work especially well for complex problem solving, coding, scientific reasoning, and multi-step agentic workflows. They’re also the best models for Codex CLI , our lightweight coding agent.
Start with gpt-5.6 Start with gpt-5.5 for most reasoning workloads. If you need the highest-intelligence API option for more challenging problems that can tolerate more latency, use gpt-5.5-pro . For lower cost, consider gpt-5.4 and for lower cost and latency, consider gpt-5.4-mini .
Reasoning models work better with the Responses
API . While the Chat Completions API
is still supported, you’ll get improved model intelligence and performance by
using Responses.
Call the Responses API and specify your reasoning model and reasoning effort:
1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
import OpenAI from "openai" ;
const openai = new OpenAI();
const prompt = `
Write a bash script that takes a matrix represented as a string with
format '[1,2],[3,4],[5,6]' and prints the transpose in the same format.
` ;
const response = await openai.responses.create({
model : "gpt-5.5" ,
model : "gpt-5.6" ,
reasoning : { effort : "low" },
input : [
{
role : "user" ,
content : prompt,
},
],
});
console .log(response.output_text); 1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
from openai import OpenAI
client = OpenAI()
prompt = """
Write a bash script that takes a matrix represented as a string with
format '[1,2],[3,4],[5,6]' and prints the transpose in the same format.
"""
response = client.responses.create(
model= "gpt-5.5" ,
model= "gpt-5.6" ,
reasoning={ "effort" : "low" },
input =[
{
"role" : "user" ,
"content" : prompt
}
]
)
print (response.output_text) 1
2
3
4
-
5
6
7
8
9
10
11
12
13
curl https://api.openai.com/v1/responses \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $OPENAI_API_KEY " \
-d '{
"model": "gpt-5.5",
"model": "gpt-5.6",
"reasoning": {"effort": "low"},
"input": [
{
"role": "user",
"content": "Write a bash script that takes a matrix represented as a string with format \"[1,2],[3,4],[5,6]\" and prints the transpose in the same format."
}
]
}' ··· 4 unchanged blocks (Reasoning effort) — click to show Reasoning effort
The reasoning.effort parameter guides the model on how much to think when performing a task.
Supported values are model-dependent and can include none , minimal , low , medium , high , and xhigh . Lower effort favors speed and lower token usage, while at higher effort the model thinks more completely to provide higher quality responses. The models also reason adaptively across reasoning efforts, using fewer tokens for simpler tasks and thinking harder for complex tasks.
Defaults are also model-dependent rather than universal. gpt-5.5 defaults to medium reasoning effort. This is the best starting point for gpt-5.5 ’s full balance of quality, reliability and performance.
For faster time to first visible token in latency-sensitive applications, ask the model to generate a short preamble before continuing with deeper reasoning.
Some models support only a subset of these values, so check the relevant model page before choosing a setting.
GPT-5.6 models support standard and pro reasoning modes in the Responses API. standard is the default. Set reasoning.mode to pro for difficult tasks that need more model work and can tolerate higher latency and token usage.
Reasoning mode and reasoning effort are independent. Mode selects standard or pro execution, while reasoning.effort controls how much reasoning the model applies within that mode. If you omit reasoning.effort , GPT-5.6 defaults to medium in both modes.
1
2
3
4
5
6
7
8
9
10
11
curl https://api.openai.com/v1/responses \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $OPENAI_API_KEY " \
-d '{
"model": "gpt-5.6",
"reasoning": {
"mode": "pro",
"effort": "medium"
},
"input": "Review this database migration plan and identify potential failure modes."
}' Pro mode aggregates the model work performed to produce the final answer and bills those tokens at the selected model’s standard token rates . Pro mode performs more model work than standard mode, increasing token usage and cost. Existing Pro model IDs keep their current behavior and pricing.
Reasoning models introduce reasoning tokens in addition to input and output tokens. The models use these reasoning tokens to “think,” breaking down the prompt and considering multiple approaches to generating a response. Our reasoning models like gpt-5.5 and gpt-5.4 support interleaved thinking, where the model is able to generate visible output tokens before and in between thinking, and is able to think in between tool calls.
Here is an example of a multi-step conversation between a user and an assistant. Input and output tokens from each step are carried over, while reasoning tokens are discarded.
Here is the default behavior for a multi-step conversation between a user and an assistant. Input and output tokens from each step are carried over, while reasoning from earlier turns is not rendered into the next sample. Models that support persisted reasoning can change this behavior with reasoning.context .
While reasoning tokens are not visible via the API, they still occupy space in
the model’s context window and are billed as output
tokens .
It’s important to ensure there’s enough space in the context window for reasoning tokens when creating responses. Depending on the problem’s complexity, the models may generate anywhere from a few hundred to tens of thousands of reasoning tokens. The exact number of reasoning tokens used is visible in the usage object of the response object , under output_tokens_details :
1
2
3
4
5
6
7
8
9
10
11
12
13
{
"usage" : {
"input_tokens" : 75 ,
"input_tokens_details" : {
"cached_tokens" : 0
},
"output_tokens" : 1186 ,
"output_tokens_details" : {
"reasoning_tokens" : 1024
},
"total_tokens" : 1261
}
} Context window lengths are found on the model reference page , and will differ across model snapshots.
To manage costs with reasoning models, you can limit the total number of tokens the
model generates, including reasoning tokens, visible output tokens, and non-visible
formatting tokens, by using the
max_output_tokens
parameter. See output token counts for details about how generated tokens are reflected in usage and output limits.
Allocating space for reasoning
If the generated tokens reach the context window limit or the max_output_tokens value you’ve set, you’ll receive a response with a status of incomplete and incomplete_details with reason set to max_output_tokens . This might occur before any visible output tokens are produced, meaning you could incur costs for input and reasoning tokens without receiving a visible response.
To prevent this, ensure there’s sufficient space in the context window or adjust the max_output_tokens value to a higher number. OpenAI recommends reserving at least 25,000 tokens for reasoning and outputs when you start experimenting with these models. As you become familiar with the number of reasoning tokens your prompts require, you can adjust this buffer accordingly.
1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
import OpenAI from "openai" ;
const openai = new OpenAI();
const prompt = `
Write a bash script that takes a matrix represented as a string with
format '[1,2],[3,4],[5,6]' and prints the transpose in the same format.
` ;
const response = await openai.responses.create({
model : "gpt-5.5" ,
model : "gpt-5.6" ,
reasoning : { effort : "medium" },
input : [
{
role : "user" ,
content : prompt,
},
],
max_output_tokens : 300 ,
});
if (
response.status === "incomplete" &&
response.incomplete_details.reason === "max_output_tokens"
) {
console .log( "Ran out of tokens" );
if (response.output_text?.length > 0 ) {
console .log( "Partial output:" , response.output_text);
} else {
console .log( "Ran out of tokens during reasoning" );
}
} 1
2
3
4
5
6
7
8
9
10
-
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
from openai import OpenAI
client = OpenAI()
prompt = """
Write a bash script that takes a matrix represented as a string with
format '[1,2],[3,4],[5,6]' and prints the transpose in the same format.
"""
response = client.responses.create(
model= "gpt-5.5" ,
model= "gpt-5.6" ,
reasoning={ "effort" : "medium" },
input =[
{
"role" : "user" ,
"content" : prompt
}
],
max_output_tokens= 300 ,
)
if response.status == "incomplete" and response.incomplete_details.reason == "max_output_tokens" :
print ( "Ran out of tokens" )
if response.output_text:
print ( "Partial output:" , response.output_text)
else :
print ( "Ran out of tokens during reasoning" ) ··· 5 unchanged blocks (Keeping reasoning items in context) — click to show Keeping reasoning items in context
When doing function calling with a reasoning model in the Responses API , we highly recommend you pass back any reasoning items returned with the last function call (in addition to the output of your function). If the model calls multiple functions consecutively, you should pass back all reasoning items, function call items, and function call output items, since the last user message. This allows the model to continue its reasoning process to produce better results in the most token-efficient manner.
The simplest way to do this is to pass in all reasoning items from a previous response into the next one. Our systems will smartly ignore any reasoning items that aren’

[truncated]

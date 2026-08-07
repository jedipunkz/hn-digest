---
source: "https://jaketao.com/language/en/why-openai-upgrading-api/"
hn_url: "https://news.ycombinator.com/item?id=49215665"
title: "From Chat Completions to Responses: Why Is OpenAI Upgrading Its Core API?"
article_title: "From Chat Completions to Responses: Why Is OpenAI Upgrading Its Core API? - Jake blog - Articles From JakeBlog"
author: "taojing10"
captured_at: "2026-08-07T20:30:00Z"
capture_tool: "hn-digest"
hn_id: 49215665
score: 1
comments: 0
posted_at: "2026-08-07T20:14:43Z"
tags:
  - hacker-news
  - translated
---

# From Chat Completions to Responses: Why Is OpenAI Upgrading Its Core API?

- HN: [49215665](https://news.ycombinator.com/item?id=49215665)
- Source: [jaketao.com](https://jaketao.com/language/en/why-openai-upgrading-api/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T20:14:43Z

## Translation

タイトル: チャットの完了から応答まで: なぜ OpenAI はコア API をアップグレードしているのですか?
記事のタイトル: チャットの完了から応答まで: OpenAI はなぜコア API をアップグレードしているのですか? - ジェイクブログ - ジェイクブログの記事
説明: 大規模言語モデルのアプリケーションを構築したことがある場合は、おそらく次のエンドポイントから始めたことがあるでしょう: POST /v1/chat/completions GPT-3.5 および GPT-4 の時代では、このエンドポイントは実質的に OpenAI API と同義でした。開発者が一連のメッセージを渡すと、モデルがメッセージを生成します。
[切り捨てられた]

記事本文:
チャットの完了から応答まで: OpenAI がコア API をアップグレードするのはなぜですか?
チャットの完了: センターでの会話メッセージ
応答: 単一の「タスク応答」を中心とする
エージェントがさらに応答 API を必要とするのはなぜですか?
API ゲートウェイはどのように設計すべきですか?
大規模言語モデルのアプリケーションを構築したことがある場合は、おそらく次のエンドポイントから始めたことがあるでしょう。
POST /v1/chat/completions
GPT-3.5 および GPT-4 の時代では、このエンドポイントは事実上 OpenAI API と同義でした。開発者は一連のメッセージを渡すと、モデルはコンテキストに基づいて次の応答を生成します。
しかし、アプリケーションが「チャットボット」から「ツールの呼び出し、タスクの実行、マルチモーダル コンテンツの処理が可能なエージェント」へと進化するにつれて、API の構造も変化し始めています。 OpenAI は、より統合されたアプローチを導入しました。
POST /v1/responses
これは、チャット コンプリートが時代遅れになったという意味ではありません。むしろ、エージェントのより複雑なワークフローに対して、より適切な抽象化を提供します。
チャットの完了: センターでの会話メッセージ
チャット コンプリーションの中核となるデータ構造はメッセージです。各リクエストラウンドで、クライアントはモデルが現在のタスクを理解するために必要なコンテキストを送信する必要があります。
たとえば、ユーザーが北京の天気をリクエストするとします。
{
"モデル": "gpt-5.6",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": "帮我查询北京天气"
}
]、
「ツール」: [
{
"タイプ": "関数",
「関数」: {
"名前": "get_weather",
"説明": "查询天气",
「パラメータ」: {
"タイプ": "オブジェクト",
"プロパティ": {
「都市」: {
"タイプ": "文字列"
}
}、
"必須": ["都市"]
}
}
}
】
}
モデルがツールの呼び出しを決定すると、次のような結果が返されます。
{
「選択肢」: [
{
「メッセージ」: {
"役割": "アシスタント",
「コンテンツ」: null、
"ツールコール": [

{
"id": "call_weather_001",
"タイプ": "関数",
「関数」: {
"名前": "get_weather",
"引数": "{\"city\":\"北京\"}"
}
}
】
}
}
】
}
アプリケーションが get_weather を実行した後、次のリクエストには前の会話、モデルによって開始されたツール呼び出し、およびそれらのツールの実行結果が含まれている必要があります。
{
"モデル": "gpt-5.6",
「メッセージ」: [
{
"ロール": "ユーザー",
"content": "帮我查询北京天气"
}、
{
"役割": "アシスタント",
「コンテンツ」: null、
"ツールコール": [
{
"id": "call_weather_001",
"タイプ": "関数",
「関数」: {
"名前": "get_weather",
"引数": "{\"city\":\"北京\"}"
}
}
】
}、
{
"役割": "ツール",
"tool_call_id": "call_weather_001",
"content": "北京晴，25℃"
}
】
}
このアプローチは直感的で確立されており、ほとんどのチャット シナリオに適しています。
ただし、これには明らかなエンジニアリング上の欠点が 1 つあります。コンテキスト管理は主にクライアントによって処理されるということです。会話が長くなり、ツールの呼び出しが増えると、アプリケーションは履歴メッセージを継続的に維持し、再生する必要があります。
応答: 単一の「タスク応答」を中心とする
Responses API は異なるアプローチを採用しています。つまり、モデル出力を単なるテキストとしてではなく、テキスト、推論、ツール呼び出し、画像、または構造化された結果を含む「応答」として扱います。
もう一度天気クエリを例として使用してみましょう。
{
"モデル": "gpt-5.6",
"input": "帮我查询北京天气",
「ツール」: [
{
"タイプ": "関数",
"名前": "get_weather",
"説明": "查询天气",
「パラメータ」: {
"タイプ": "オブジェクト",
"プロパティ": {
「都市」: {
"タイプ": "文字列"
}
}、
"必須": ["都市"]
}
}
】
}
モデルは、関数呼び出しを含む応答を返します。
{
"id": "resp_123",
「出力」: [
{
"タイプ": "関数呼び出し",
"call_id": "call_weather_001",
"名前": "get_weather",
"引数": "{\"city\":\"北京\"}"
}
】
}
船尾

ツールの実行が完了すると、次のラウンドでは新しい結果を送信し、前の応答を参照するだけで済みます。
{
"モデル": "gpt-5.6",
"previous_response_id": "resp_123",
「入力」: [
{
"タイプ": "関数呼び出し出力",
"call_id": "call_weather_001",
"出力": "北京晴，25°C"
}
】
}
OpenAI は、previous_response_id を使用して、前のコンテキストをツール呼び出しに関連付けることができます。クライアントはメッセージ履歴全体を毎回手動で再生する必要がないため、エージェントのオーケストレーション コードがより簡潔になります。
ただし、これは「コンテキストにコストがかからなくなる」という意味ではないことに注意してください。 previous_response_id を使用すると、クライアントがメッセージ履歴を構築および維持する際の複雑さが軽減されます。応答チェーン内の以前の入力トークンは、引き続き入力トークンとして課金されます。
エージェントがさらに応答 API を必要とするのはなぜですか?
質疑応答のシナリオでは、メッセージは自然なものです。しかし、エージェントは多くの場合、会話、ツールの呼び出し、ツールの結果、構造化データの間を常に切り替える必要があります。
チャット補完でもこれらのタスクを処理できますが、手順の数が増加するため、クライアントは複雑なメッセージ履歴を独自に維持し、ツール呼び出しがその結果に正しくマッピングされていることを確認する必要があります。
Responses API の焦点は、新しい機能を追加することではなく、これらの要素を応答アイテムに統合し、以前の応答に基づいてタスクの継続をサポートすることで、複雑なエージェント ワークフローにより適したものにすることにあります。
API ゲートウェイはどのように設計すべきですか?
ゲートウェイが OpenAI、Claude、Gemini、DeepSeek などのモデルを同時に統合する場合、重要なのは、すべてのリクエストをレスポンスとして書き換えないことです。
より現実的なアプローチは、チャットの完了や応答など、クライアントが使い慣れたインターフェイスを外部使用のために保持することです。リクエストが入ると

システムでは、それらは対応するコンバーターによって解析され、同じ処理パイプラインにルーティングされます。
OwlVigil はまさにこのアプローチを採用しています。チャットを応答に置き換えるのではなく、異なるプロトコルが同じゲートウェイ機能セットを共有できるようにします。
クライアント
§─ チャットの完了
§─ 回答
§─ 人間的なメッセージ
━─ Gemini API
↓
インバウンドコンバーター
↓
統合 LLM リクエスト モデル
↓
モデルのマッピング、ルーティング、レート制限、再試行
↓
アウトバウンドコンバーター
↓
OpenAI
§─ クロード
§─ 双子座
━─ディープシーク
ここでの「統合」という用語は、単一ベンダーのプロトコルへのバインディングを強制することを意味するのではなく、メッセージ、ツール呼び出し、ツール結果、モデル パラメーター、およびストリーミング応答を単一の処理パイプラインに配置することを意味します。
ジェイクブログの記事
ついにフルタイムでこれを構築中 — Syrovex について
ジェイクブログの記事
KV キャッシュとプロンプト キャッシュ: 違いは何ですか?またどのような関係がありますか?
目次の切り替え 目次の切り替え
チャットの完了: センターでの会話メッセージ
応答: 単一の「タスク応答」を中心とする
エージェントがさらに応答 API を必要とするのはなぜですか?
API ゲートウェイはどのように設計すべきですか?
Copyright © 2003-2026 ジェイクブログ

## Original Extract

If you’ve ever built a large-language-model application, you’ve most likely started with this endpoint: POST /v1/chat/completions In the era of GPT-3.5 and GPT-4, this endpoint was practically synonymous with the OpenAI API. Developers would pass in a set of messages, and the model would generate th
[truncated]

From Chat Completions to Responses: Why Is OpenAI Upgrading Its Core API?
Chat Completions: Conversation Messages at the Center
Responses: Centered Around a Single “Task Response”
Why does the Agent need the Responses API more?
How should an API Gateway be designed?
If you’ve ever built a large-language-model application, you’ve most likely started with this endpoint:
POST /v1/chat/completions
In the era of GPT-3.5 and GPT-4, this endpoint was practically synonymous with the OpenAI API. Developers would pass in a set of messages , and the model would generate the next response based on the context.
But as applications have evolved from “chatbots” to “agents capable of invoking tools, executing tasks, and processing multimodal content,” the structure of the API has also begun to change. OpenAI has introduced a more unified approach:
POST /v1/responses
This doesn’t mean Chat Completions are obsolete; rather, it provides a more appropriate abstraction for the more complex workflows of agents.
Chat Completions: Conversation Messages at the Center
The core data structure of Chat Completions is messages . In each request round, the client must submit the context required for the model to understand the current task.
For example, a user requests the weather in Beijing:
{
"model": "gpt-5.6",
"messages": [
{
"role": "user",
"content": "帮我查询北京天气"
}
],
"tools": [
{
"type": "function",
"function": {
"name": "get_weather",
"description": "查询天气",
"parameters": {
"type": "object",
"properties": {
"city": {
"type": "string"
}
},
"required": ["city"]
}
}
}
]
}
After the model decides to call a tool, it will return a result similar to the following:
{
"choices": [
{
"message": {
"role": "assistant",
"content": null,
"tool_calls": [
{
"id": "call_weather_001",
"type": "function",
"function": {
"name": "get_weather",
"arguments": "{\"city\":\"北京\"}"
}
}
]
}
}
]
}
After the application executes get_weather , the next request must include the previous conversation, the tool calls initiated by the model, and the results of those tool executions:
{
"model": "gpt-5.6",
"messages": [
{
"role": "user",
"content": "帮我查询北京天气"
},
{
"role": "assistant",
"content": null,
"tool_calls": [
{
"id": "call_weather_001",
"type": "function",
"function": {
"name": "get_weather",
"arguments": "{\"city\":\"北京\"}"
}
}
]
},
{
"role": "tool",
"tool_call_id": "call_weather_001",
"content": "北京晴，25°C"
}
]
}
This approach is intuitive, well-established, and still suitable for most chat scenarios.
However, it has one obvious engineering shortcoming: context management is primarily handled by the client. As conversations grow longer and tool calls increase, the application must continuously maintain and replay historical messages.
Responses: Centered Around a Single “Task Response”
The Responses API takes a different approach: it treats model output not merely as a piece of text, but as a “response” that may include text, reasoning, tool calls, images, or structured results.
Let’s use the weather query as an example again:
{
"model": "gpt-5.6",
"input": "帮我查询北京天气",
"tools": [
{
"type": "function",
"name": "get_weather",
"description": "查询天气",
"parameters": {
"type": "object",
"properties": {
"city": {
"type": "string"
}
},
"required": ["city"]
}
}
]
}
The model returns a response containing a function call:
{
"id": "resp_123",
"output": [
{
"type": "function_call",
"call_id": "call_weather_001",
"name": "get_weather",
"arguments": "{\"city\":\"北京\"}"
}
]
}
After the tool completes execution, the next round only needs to submit the new results and reference the previous response:
{
"model": "gpt-5.6",
"previous_response_id": "resp_123",
"input": [
{
"type": "function_call_output",
"call_id": "call_weather_001",
"output": "北京晴，25°C"
}
]
}
OpenAI can use previous_response_id to associate the previous context with the tool call. The client does not need to manually replay the entire message history each time, making the Agent’s orchestration code more concise.
However, note that this does not mean “context no longer incurs costs.” Using previous_response_id reduces the complexity for the client in constructing and maintaining the message history; previous input tokens in the response chain will still be billed as input tokens.
Why does the Agent need the Responses API more?
In a question-and-answer scenario, messages are natural; but Agents often need to constantly switch between conversations, tool calls, tool results, and structured data.
Chat Completions can also handle these tasks, but as the number of steps increases, the client must maintain a complex messages history on its own and ensure that tool calls are correctly mapped to their results.
The focus of the Responses API is not on adding a new capability, but on unifying these elements into response items and supporting the continuation of tasks based on the previous response, making it better suited for complex Agent workflows.
How should an API Gateway be designed?
If the Gateway integrates models such as OpenAI, Claude, Gemini, and DeepSeek simultaneously, the key is not to rewrite all requests as Responses.
A more practical approach is to retain client-familiar interfaces—such as Chat Completions and Responses—for external use; once requests enter the system, they are parsed by the corresponding converters and routed into the same processing pipeline.
OwlVigil adopts precisely this approach: rather than replacing Chat with Responses, it allows different protocols to share the same set of gateway capabilities.
Client
├─ Chat Completions
├─ Responses
├─ Anthropic Messages
└─ Gemini API
↓
Inbound Converter
↓
Unified LLM Request Model
↓
Model mapping, routing, rate limiting, retries
↓
Outbound converter
↓
OpenAI
├─ Claude
├─ Gemini
└─ DeepSeek
The term “unified” here does not mean forcing a binding to a single vendor’s protocol, but rather placing messages, tool calls, tool results, model parameters, and streaming responses into a single processing pipeline.
Articles From JakeBlog
I’m Finally Building This Full-Time — Meet Syrovex
Articles From JakeBlog
KV Cache vs. Prompt Cache: What’s the Difference, and How Are They Related?
Table of Contents Toggle Table of Content Toggle
Chat Completions: Conversation Messages at the Center
Responses: Centered Around a Single “Task Response”
Why does the Agent need the Responses API more?
How should an API Gateway be designed?
Copyright © 2003-2026 Jake blog

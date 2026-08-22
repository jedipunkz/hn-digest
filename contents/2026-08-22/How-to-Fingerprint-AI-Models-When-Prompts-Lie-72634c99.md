---
source: "https://openrating.io/blog/current-state-of-ai-model-fingerprinting"
hn_url: "https://news.ycombinator.com/item?id=49399051"
title: "How to Fingerprint AI Models When Prompts Lie"
article_title: "How to Fingerprint AI Models When Prompts Lie | OpenRating Research — OpenRating"
image: "https://openrating.io/opengraph-image?a99571f7e3ba7097"
author: "m00dy"
captured_at: "2026-08-22T13:24:44Z"
capture_tool: "hn-digest"
hn_id: 49399051
score: 2
comments: 0
posted_at: "2026-08-22T12:26:50Z"
tags:
  - hacker-news
  - translated
---

# How to Fingerprint AI Models When Prompts Lie

- HN: [49399051](https://news.ycombinator.com/item?id=49399051)
- Source: [openrating.io](https://openrating.io/blog/current-state-of-ai-model-fingerprinting)
- Score: 2
- Comments: 0
- Posted: 2026-08-22T12:26:50Z

## Translation

タイトル: プロンプトが嘘をついたときに AI モデルをフィンガープリントする方法
記事のタイトル: プロンプトが嘘をついたときに AI モデルをフィンガープリントする方法 | OpenRating リサーチ — OpenRating
説明: プロキシ ラッパーに対して動作評価が失敗する理由、およびトークナイザーのマージ、テンプレート オフセット、およびエラー境界が真の基礎モデルを明らかにする方法。

記事本文:
プロンプトに嘘がある場合に AI モデルをフィンガープリントする方法 | OpenRating Research — OpenRating OpenRating は、独立したオープンソースの AI 評価プラットフォームです。コミュニティに参加する
研究に戻る テクニカル ノート 5 分で読む · 2026 年 8 月 プロンプトが嘘をついたときに AI モデルをフィンガープリントする方法
不明な API エンドポイントに対して「誰が作成したの?」とクエリすると、 、10 個のトークンで書き換えられるペルソナをテストしています。実際にどのモデルが実行されているかを確認するには、インフラストラクチャ層を調査する必要があります。
サードパーティ API ゲートウェイ、プロキシ ルーター、または匿名アリーナ モデルを評価する場合、プロンプトベースの識別は基本的に役に立ちません。基本的なシステム プロンプトや軽量の微調整により、クロードが GPT-4 であると主張したり、ラマの声をミストラルのようにしたり、拒否動作を完全に変更したりすることができます。
ペルソナとトーンは順応性があります。しかし、基礎となるサービス スタック (トークナイザー、マージ テーブル、サービス ハーネス、エラー検証境界) は厳格です。これらを偽装するには、プロキシ プロバイダーがストリームを傍受し、リアルタイムでトークンを双方向に変換し、オンザフライでメタデータを書き換える必要があるため、深刻な遅延とストリームの非同期が発生します。
ここでは、クライアント側のフィンガープリンティング エンジンが、会話の合図の代わりにインフラストラクチャ アーティファクトを使用してモデルを決定論的に識別する方法を示します。
1. トークナイザーの語彙とバイト フォールバック
すべてのラボでは、カスタム語彙サイズ、マージ テーブル、およびフォールバック ルールを使用して、個別のトークナイザーをトレーニングします。トークナイザーはモデルの重みの正確な入力行列の形状を決定するため、モデルは完全な再学習を行わないとテキストのセグメント化方法を変更できません。
小さな固定テキスト入力をエンドポイントに送信し、返された use.prompt_tokens を検査します。
2. テンプレートのオフセットと非表示のハーネス
推論エンジン (vLLM、SGLang、TGI、TensorRT-LLM) はユーザー プロンプトをチャット テンプレートにラップします (

ChatML や Jinja テンプレートなど）、デフォルトのシステム プロンプトを挿入することがよくあります。
空のプロンプトまたは単一の 1 トークン ペイロードを送信し、返された use.prompt_tokens を生のペイロード トークン数と比較することで、静的オーバーヘッドを計算します。
このデルタは、プロンプトがモデルに到達する前に、上流で適用された非表示のガードレール テンプレート、プラットフォーム ラッパー、チャット フォーマット アーティファクトを公開します。
3. エラー分類と検証の上限
通常の応答を検証することは、ストーリーの一部にすぎません。意図的に無効なパラメータを入力すると、特定のエンジニアリング チームによって作成されたバックエンド検証がトリガーされます。
温度上限: 送信温度: 2.5 では、一部のランタイムは >1.0 で失敗し、他のランタイムは >2.0 で失敗し、明確なエラー文字列が表示されます。
コンテキストの拒否: max_tokens: 1000000000 をリクエストすると、バックエンドはリクエストを拒否し、実際の物理生成制限をエコーし​​ます (例: "max_tokens は ≤ 131072" )。
ベンダー エラー コード: 安全フィルター トリガーは、多くの場合、独自の数値コードを返します (例: Zhipu AI 内部安全コード 1301 )。
4. 応答のシリアル化方言
サーバーが JSON をフォーマットする方法に関する実装の詳細により、ランタイムが明らかになります。
正確な Choices[0].finish_reason 文字列 (例: stop vs end_turn vs content_filter )。
null またはオプションのフィールドが応答本文でどのようにシリアル化されるか。
ストリーミング中の Server-Sent Events (SSE) フレーミングの不具合。
プローブは純粋に決定的であるため、ブラウザーで直接実行します。リクエストは、中間プロキシやログ サーバーを経由せずに、クライアント タブからターゲット エンドポイントに直接ディスパッチされます。
各プローブは、決定的な値を返す単純なランナー コントラクトをエクスポートします。
デフォルトをエクスポート {
名前: "トークナイザー/cjk-密度",
説明: "CJK トークンのセグメンテーション数を測定します",
非同期実行(ctx) {
const res = ctx を待ちます。

チャット({
メッセージ: [{ 役割: "ユーザー", 内容: "人工智能モデル基礎评测系" }],
max_tokens: 1、
温度: 0、
});
戻り値 {
値: res?.usage?.prompt_tokens ?? "ERR_NO_USAGE",
};
}、
};なぜこれが重要なのか
API ラッパーやプロキシ ルーターが普及するにつれて、検証は信頼や会話の雰囲気に頼ることができなくなります。トークナイザー、テンプレート デルタ、境界条件などの決定論的なインフラストラクチャ アーティファクトをテストすることで、どのモデルが実際にトラフィックを処理しているかを正確に検証できます。
ベンチマーク ニュースレター 新しい研究の最新情報を入手する
購読して最新のエンジニアリング ノートと毎月のベンチマーク リリースを受信してください。
4,200 人を超える ML エンジニアや研究者に加わりましょう。いつでも購読を解除してください。
受信トレイに新しい評価が届きます。
© 2026 OpenRating — グローバル評価プラットフォーム。デンマーク、コペンハーゲン。
独立した。オープンソース。 GitHub に貢献する

## Original Extract

Why behavioral evaluations fail against proxy wrappers, and how tokenizer merges, template offsets, and error boundaries expose the true underlying model.

How to Fingerprint AI Models When Prompts Lie | OpenRating Research — OpenRating OpenRating is an independent, open-source AI rating platform. Join the community
Back to Research Technical Notes 5 min read · August 2026 How to Fingerprint AI Models When Prompts Lie
If you query an unknown API endpoint with "Who made you?" , you are testing a persona that can be rewritten in 10 tokens. To verify what model is actually running, you have to probe the infrastructure layer.
When evaluating third-party API gateways, proxy routers, or anonymous arena models, prompt-based identification is essentially useless. A basic system prompt or lightweight fine-tune can make Claude claim it is GPT-4, make Llama sound like Mistral, or alter refusal behaviors entirely.
Persona and tone are malleable. But the underlying serving stack (tokenizers, merge tables, serving harnesses, and error validation boundaries) is rigid. To fake those, a proxy provider would need to intercept streams, translate tokens bidirectionally in real time, and rewrite metadata on the fly, which introduces severe latency and stream desyncs.
Here is how our client-side fingerprinting engine identifies models deterministically using infrastructure artifacts instead of conversational cues.
1. Tokenizer Vocabularies and Byte Fallbacks
Every lab trains a distinct tokenizer with custom vocabulary sizes, merge tables, and fallback rules. Because the tokenizer determines the exact input matrix shape of the model weights, a model cannot alter how it segments text without a full retrain.
We send small, fixed text inputs to the endpoint and inspect the returned usage.prompt_tokens :
2. Template Offsets and Hidden Harnesses
Inference engines (vLLM, SGLang, TGI, TensorRT-LLM) wrap user prompts in chat templates (like ChatML or Jinja templates) and often inject default system prompts.
By sending an empty prompt or a single 1-token payload and comparing the returned usage.prompt_tokens against the raw payload token count, we calculate the static overhead:
This delta exposes hidden guardrail templates, platform wrappers, and chat formatting artifacts applied upstream before the prompt hits the model.
3. Error Taxonomy and Validation Ceilings
Validating normal responses only tells part of the story. Feeding intentionally invalid parameters triggers backend validation written by specific engineering teams:
Temperature ceilings: Sending temperature: 2.5 causes some runtimes to fail at >1.0, others at >2.0, with distinct error strings.
Context refusals: Requesting max_tokens: 1000000000 forces the backend to reject the request and echo its true physical generation limit (e.g. "max_tokens must be ≤ 131072" ).
Vendor error codes: Safety filter triggers often return proprietary numeric codes (e.g. Zhipu AI internal safety code 1301 ).
4. Response Serialization Dialects
Minor implementation details in how the server formats JSON reveal the runtime:
The exact choices[0].finish_reason string (e.g. stop vs end_turn vs content_filter ).
How null or optional fields are serialized in the response body.
Server-Sent Events (SSE) framing quirks during streaming.
Because probes are purely deterministic, we run them directly in the browser. Requests are dispatched straight from the client tab to the target endpoint without passing through any intermediate proxy or logging server.
Each probe exports a simple runner contract that returns a deterministic value:
export default {
name: "tokenizer/cjk-density",
description: "Measures CJK token segmentation count",
async run(ctx) {
const res = await ctx.chat({
messages: [{ role: "user", content: "人工智能模型基准评测体系" }],
max_tokens: 1,
temperature: 0,
});
return {
value: res?.usage?.prompt_tokens ?? "ERR_NO_USAGE",
};
},
}; Why This Matters
As API wrappers and proxy routers become more prevalent, verification cannot rely on trust or conversational vibes. By testing deterministic infrastructure artifacts like tokenizers, template deltas, and boundary conditions, we can accurately verify what model is actually serving your traffic.
Benchmark Newsletter Get new research updates
Subscribe to receive our latest engineering notes and monthly benchmark releases.
Join 4,200+ ML engineers and researchers. Unsubscribe at any time.
Get new ratings in your inbox.
© 2026 OpenRating — A global rating platform. Copenhagen, Denmark.
Independent. Open source. Contribute on GitHub

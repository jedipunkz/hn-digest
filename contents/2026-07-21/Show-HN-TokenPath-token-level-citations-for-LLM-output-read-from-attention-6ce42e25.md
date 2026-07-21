---
source: "https://tokenpath.ai"
hn_url: "https://news.ycombinator.com/item?id=48997273"
title: "Show HN: TokenPath – token-level citations for LLM output, read from attention"
article_title: "TokenPath — Citations, built for AI agents"
author: "apoorvumang"
captured_at: "2026-07-21T20:14:32Z"
capture_tool: "hn-digest"
hn_id: 48997273
score: 3
comments: 0
posted_at: "2026-07-21T19:48:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TokenPath – token-level citations for LLM output, read from attention

- HN: [48997273](https://news.ycombinator.com/item?id=48997273)
- Source: [tokenpath.ai](https://tokenpath.ai)
- Score: 3
- Comments: 0
- Posted: 2026-07-21T19:48:06Z

## Translation

タイトル: HN を表示: TokenPath – LLM 出力のトークンレベルの引用、注意から読む
記事のタイトル: TokenPath — AI エージェント向けに構築された引用
説明: 生成されたすべてのトークンをソース内でグラウンドする 1 つの API (任意のモデルからの出力用)。 TokenPath は、LLM の自己報告ではなく注目度を直接測定し、製品が動作できる機械可読なトークンレベルの引用を返します。

記事本文:
TokenPath — 引用、AI エージェント向けに構築されたトークン パス デモ ブログ ドキュメント GitHub 電話を予約する API グラウンディングを試す
AIエージェント。
モデルの応答内のすべてのクレームを、その背後にある正確なソース スパンに対してトークン レベルの粒度で解決します。あらゆるモデルの出力で動作します。アトリビューションは、回答後の 1 回の API 呼び出しであり、スタックについては何も変更されません。
図 01 — POST /v1/attributions 2 クレーム → ソース ソース文書 acme_offer_letter.txt この書簡契約 (「契約」) は、Acm e Robotics, Inc. (「会社」) と Jordan Lee の間で 2026 年 3 月 2 日に発効します。あなたには上級機械エンジニアの役割が与えられます。あなたの最初の雇用日は 2026 年 3 月 2 日になります。
Q · 契約はいつから始まり、初出勤日はいつですか?
この契約は 2026 年 3 月 2 日から発効し、最初の雇用日は 2026 年 3 月 2 日となります。
2026 年 3 月 2 日の主張は「2026 年 3 月 2 日」に解決される 信頼度 0.99
どちらの主張も「2026 年 3 月 2 日」となっており、それぞれが異なるソース範囲に解決されます。 1 回の呼び出しで、すべての主張の背後にある情報源と信頼性が返されます。
引用を添付する 3 つの方法。
測定するものが異なります。
LLM 脚注自体を用意します。簡単に言えば、追加のインフラストラクチャはありません。
粒度がありません。せいぜい番号付きの少数のソースを指すだけであり、長い文書内の正確なセルや文を指すことは決してありません。
関連するチャンクを取得し、埋め込みモデル (標準の RAG パス) を使用してそれらを再ランク付けすると、スケーリングされます。
曖昧さを解消できません。同じ事実、数値、日付が 2 か所に出現すると、どちらの回答が使用されたのかがわかりません。表や財務文書ではよくあることです。
粒状。正確な範囲 (正しいセル、正しい文) を文字単位で解決します。
明確です。特定の出来事を答えとして指摘する

同じ値が十数回出現する場合でも使用されます。
速い 。コンテキスト上で 1 回の前方パスを実行します。追加の生成やリランカーのラウンドトリップはありません。 20,000 トークンのドキュメントは 2 秒以内に解決されます。 128,000 トークンまで拡張できます。
フロンティア法案なしのフロンティアグレードの引用。
LongBench-Cite では、1 つの共有回答、1 人の審査員による読書注意の事後スコア 0.815 F1 : Anthropic の Citations API (0.812) と一致し、プロンプト フロンティア LLM (0.851) の 0.04 以内に収まりますが、どのモデルの回答でも、再生成することなく、約 7 倍安く、約 5 ～ 6 倍高速に実行されます。
引用 F1 · 報告された平均、480 のテスト例 gpt-5.5 プロンプト 0.851 TokenPath 0.815 Anthropic Citations 0.812 emb + rerank 0.622 クエリあたりのコストが ~7 倍安い · $0.013 対 ~$0.09 ~5 ～ 6 倍高速 p50 レイテンシ · 1.6 秒 vs ~8 ～ 10 秒
モデルが読み取った内容、書き込んだ内容、および主張を送信します。
1 つのリクエスト: ソース文書、質問、LLM が生成したデコーダーの回答、およびソースを希望する回答範囲。再トレーニング、logprob、生成モデルに対するベンダー ロックはありません。
POST /v1/attributions
{
"document" : "第 3 四半期、ノースウィンドの収益は 18% 増加しました…",
"question" : "収益はどれくらいの速さで増加しましたか?",
"answer" : "収益は前年比 18% 増加しました。",
"spans" : [[13, 16]] # クレーム "18%"
あなた自身のモデル、OpenAI のモデル、Anthropic のモデルなど、あらゆるモデルからの出力を処理します。
モデルを通じて注意を追跡します。
TokenPath は、インスツルメント化されたフォワード パスを実行し、モデルが実際に見た場所、つまりトークン レベルの注意をキャラクター レベルの証拠にまで抽出して読み取ります。
ソース スパンとクレームごとのスコアを取得します。
TokenPath は、各回答スパンの下でアテンションを集計し、その背後にあるドキュメント スパンと帰属の信頼度を返します。引用を表示し、出所を添付し、監査領収書を保管します。
{ "スパン" : [ {
"答え" : { "

開始": 13、"終了": 16、"テキスト": "18%" } 、
「ソース」: {
"開始": 32、"終了": 35、"テキスト": "18%"、
「信頼度」 : 0.82
}
背後にある完全なトークン×トークンの注目を集めたいですか?生のマトリックスについては /v1/attributions/heatmap を呼び出します。
…第 3 四半期のノースウィンドの収益は、主に…によって前年比 18% 増の 4,210 万ドルとなりました。
POST /v1/attributions
{
"document" : "第 3 四半期、ノースウィンドの収益は 18% 増加しました…",
"question" : "収益はどれくらいの速さで増加しましたか?",
"answer" : "収益は前年比 18% 増加しました。",
"spans" : [[13, 16]] # クレーム "18%"
あなた自身のモデル、OpenAI のモデル、Anthropic のモデルなど、あらゆるモデルからの出力を処理します。
{ "スパン" : [ {
"answer" : { "start": 13, "end": 16, "text": "18%" } ,
「ソース」: {
"開始": 32、"終了": 35、"テキスト": "18%"、
「信頼度」 : 0.82
}
背後にある完全なトークン×トークンの注目を集めたいですか?生のマトリックスについては /v1/attributions/heatmap を呼び出します。
…第 3 四半期のノースウィンドの収益は、主に…によって前年比 18% 増の 4,210 万ドルとなりました。
開始するのに SDK は必要ありません。これは 1 つの HTTPS エンドポイントだけです。文書、質問、回答、および関心のある請求範囲を送信してください。正確なソース範囲と各主張の背後にある信頼度を取得し、それを数行で引用、強調表示、または記録します。代わりに生のトークンのヒートマップが必要ですか? /v1/attributions/heatmap で 1 回の呼び出しで完了します。レート制限、使用量測定、およびキー管理がプラットフォームに付属しています。
# 1 回の呼び出し — 各クレームをそのソースに解決します
カール https://api.tokenpath.ai/v1/attributions \
-H "認可: ベアラー $TOKENPATH_API_KEY" \
-d ' {
"document": "第 3 四半期、ノースウィンドの収益は 18% 増加しました…",
"question": "収益はどのくらいの速さで増加しましたか?",
"answer": "収益は前年比 18% 増加しました。",
"spans": [[13, 16]] # クレーム "18%"
} ' # クレームごとに 1 つの解決されたソース スパン
{ "スパン" : [ {
"answer" : { "start": 13, "end": 16, "text": "18%" } ,
「ソース」: {
"開始": 32、"終了":

35、「テキスト」: "18%"、
"confidence" : 0.82 # アトリビューションの信頼度
}
}] } ドキュメントを読む クックブックを参照する 価格 100 万トークンあたり 1 ドル。それが価格モデルです。
使用量ベースで、毎月の最低額はなく、サインアップすると 1,000 万トークンが無料になります。ボリューム料金設定と企業向けの専用導入。
TokenPath が自分に適しているかどうかまだ疑問に思っていますか?
お気に入りの LLM に私たちが何をしているのか尋ねて、情報に基づいた決定を下してください。
Ask ChatGPT Ask Claude これがどこへ行くのか エージェント時代の引用。
脚注と引用は人々のために作られていますが、人々が読める範囲には限界があります。エージェントははるかに多くの情報を読み取り、私たちがエージェントにさらに多くの作業を任せるにつれて、エージェントが実際に使用した内容を監査できることが重要になります。監査の一部はエージェントによっても行われます。私たちはその下にレイヤーを構築しています。アトリビューションシグナルは、トークンレベルで、モデルのアテンションから読み取られます。
ドキュメントを読む 1,000 万個の無料トークン · カードは必要ありません
15 ～ 30 分 — スタックをご持参ください。ご自身のドキュメントの帰属を示します。
エージェント時代の引用。生成されたすべてのトークンは、そのソースに基づいており、あらゆるモデルからの出力のために、注意から直接測定されます。

## Original Extract

One API to ground every generated token in its source — for output from any model. TokenPath measures attention directly, not the LLM's self-report, to return machine-readable, token-level citations your product can act on.

TokenPath — Citations, built for AI agents token path Demo Blog Docs GitHub Book a call Try API Grounding for your
AI agent .
Resolve every claim in a model's answer to the exact source span behind it, at token-level granularity. Works with any model's output — attribution is one API call after the answer, and nothing about your stack changes.
See it work Fig. 01 — POST /v1/attributions 2 claims → sources Source document acme_offer_letter.txt This letter agreement (t he "A gr eement" ) is made effective as of March 2 , 2 026 between Acm e Robotics, Inc. (t he "C ompany" ) and Jordan Lee. You are offered the role of Senior Mechanical Engineer. Your first day of employment will be March 2 , 2 026 .
Q · When does the agreement start, and when is my first day of work?
The agreement is effective as of March 2, 2026 , and your first day of employment is March 2, 2026 .
Claim March 2, 2026 resolves to “ March 2, 2026 ” confidence 0.99
Both claims read “March 2, 2026” — each resolves to a different source span. One call returns the source and confidence behind every claim.
Three ways to attach a citation.
They differ in what they measure.
Have the LLM footnote itself. Straightforward, no extra infrastructure.
No granularity . Points at a few numbered sources at best — never the exact cell or sentence inside a long document.
Pull the relevant chunks and re-rank them with an embedding model — the standard RAG path, and it scales.
Can't disambiguate . When the same fact, number, or date appears in two places, it can't tell which one the answer used — common in tables and financial docs.
Granular . Resolves to the exact span — the right cell, the right sentence — character-precise.
Unambiguous . Points at the specific occurrence the answer used, even when the same value appears a dozen times.
Fast . One forward pass over the context — no extra generation, no reranker round-trips. A 20k-token document resolves in under two seconds; scales to 128k tokens.
Frontier-grade citations, without the frontier bill.
On LongBench-Cite — one shared answer, one judge — reading attention post-hoc scores 0.815 F1 : it matches Anthropic's Citations API (0.812) and lands within 0.04 of a prompted frontier LLM (0.851), while doing it ~7× cheaper and ~5–6× faster — on any model's answer, without regenerating it.
Citation F1 · reported avg, 480 test examples gpt-5.5 prompted 0.851 TokenPath 0.815 Anthropic Citations 0.812 emb + rerank 0.622 ~7× cheaper cost / query · $0.013 vs ~$0.09 ~5–6× faster p50 latency · 1.6 s vs ~8–10 s How it works
Send what the model read, what it wrote, and the claims.
One request: the source document, the question, the answer any decoder LLM produced, and the answer spans you want sourced. No retraining, no logprobs, no vendor lock on the generating model.
POST /v1/attributions
{
"document" : "In Q3, Northwind's revenue grew 18%…",
"question" : "How fast did revenue grow?",
"answer" : "Revenue grew 18% year over year.",
"spans" : [[13, 16]] # the claim "18%"
} Works with output from any model — yours, OpenAI's, Anthropic's.
We trace attention through the model.
TokenPath runs an instrumented forward pass and reads where the model actually looked — token-level attention, distilled down to character-level evidence.
Get a source span and score per claim.
TokenPath aggregates the attention under each answer span and returns the document span behind it, plus an attribution confidence. Render citations, attach provenance, and keep audit receipts.
{ "spans" : [ {
"answer" : { "start": 13, "end": 16, "text": "18%" } ,
"source" : {
"start": 32, "end": 35, "text": "18%",
"confidence" : 0.82
}
}] } Want the full token×token attention behind it? Call /v1/attributions/heatmap for the raw matrix.
…In Q3, Northwind's revenue grew 18% year over year to $42.1M, driven primarily by…
POST /v1/attributions
{
"document" : "In Q3, Northwind's revenue grew 18%…",
"question" : "How fast did revenue grow?",
"answer" : "Revenue grew 18% year over year.",
"spans" : [[13, 16]] # the claim "18%"
} Works with output from any model — yours, OpenAI's, Anthropic's.
{ "spans" : [ {
"answer" : { "start": 13, "end": 16, "text": "18%" } ,
"source" : {
"start": 32, "end": 35, "text": "18%",
"confidence" : 0.82
}
}] } Want the full token×token attention behind it? Call /v1/attributions/heatmap for the raw matrix.
…In Q3, Northwind's revenue grew 18% year over year to $42.1M, driven primarily by…
No SDK required to start — it's one HTTPS endpoint. Send the document, the question, the answer, and the claim spans you care about; get the exact source span and a confidence behind each claim back, and cite, highlight, or log it in a few lines. Need the raw token heatmap instead? It's one call away at /v1/attributions/heatmap . Rate limits, usage metering, and key management come with the platform.
# one call — resolve each claim to its source
curl https://api.tokenpath.ai/v1/attributions \
-H "Authorization: Bearer $TOKENPATH_API_KEY" \
-d ' {
"document": "In Q3, Northwind's revenue grew 18%…",
"question": "How fast did revenue grow?",
"answer": "Revenue grew 18% year over year.",
"spans": [[13, 16]] # the claim "18%"
} ' # one resolved source span per claim
{ "spans" : [ {
"answer" : { "start": 13, "end": 16, "text": "18%" } ,
"source" : {
"start": 32, "end": 35, "text": "18%",
"confidence" : 0.82 # attribution confidence
}
}] } Read the docs Browse the cookbook Pricing A dollar per million tokens. That's the pricing model.
Usage-based, no monthly minimum, 10M tokens free when you sign up. Volume pricing and dedicated deployments for enterprises.
Still wondering if TokenPath is right for you?
Ask your favorite LLM what we do, then make an informed decision.
Ask ChatGPT Ask Claude Where this is going Citation for the agentic era.
Footnotes and citations were built for people — and people can only read so much. Agents read far more, and as we offload more work to them, being able to audit what they actually used matters. Some of that auditing will be done by agents too. We're building the layer underneath: the attribution signal, at the token level, read from the model's attention.
Read the docs 10M free tokens · no card required
15–30 min — bring your stack, we'll show you attribution on your own docs.
Citation for the agentic era. Every generated token, grounded in its source — measured directly from attention, for output from any model.

---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48582429"
title: "Ask HN: How do you find out if the LLM API is giving degraded responses"
article_title: ""
author: "imviky"
captured_at: "2026-06-18T10:36:08Z"
capture_tool: "hn-digest"
hn_id: 48582429
score: 2
comments: 0
posted_at: "2026-06-18T08:24:21Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How do you find out if the LLM API is giving degraded responses

- HN: [48582429](https://news.ycombinator.com/item?id=48582429)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T08:24:21Z

## Translation

タイトル: HN に質問: LLM API の応答が低下しているかどうかを確認するにはどうすればよいですか?
HN テキスト: 複数の LLM API、または OpenAI、Claude、Gemini などの単一の LLM API の上に構築している場合、API の性能が低下し始めたとき (TTFT の遅さ、エラー率の上昇、タイムアウト)、どうしますか。さらに悪いことに、応答はあるもののモデルがドリフトしている場合もあります。これをどうやって調べますか?これが広範囲にわたる痛みなのか、それとも単に私に不運があったのかを理解しようとしています。 4 つの具体的な質問: 1. LLM API がサイレントに機能低下を開始した場合、現在どのようにしてそれを確認しますか? (独自の監視? ユーザーからの苦情? ステータス ページの確認? Reddit?) 2. 「これはプロバイダーであり、私のコードではない」と確認するのに通常どのくらい時間がかかりますか? 3. もしあなたが気づく前に、現在 Sonnet 上で Claude API が高い TTFT を表示していると何かが伝えたとしたら、あなたの運用方法に何か変化はありますか?それとも、関係なく再試行して先に進みますか? 4. ユーザーが気づく前に、LLM の動作が逸脱したことを知らせる独立したアラート サービスに料金を支払いますか?これが実際に問題にならないのであれば、それが私が得ることができる最も有益な答えになると思います。

## Original Extract

If you are building on top of multiple LLM APIs or even a single one amongst OpenAI, Claude, Gemini, etc. what do you do when the API starts degrading (slow TTFT, elevated error rates, timeouts). Or even worse, when there are responses but the model is drifting. How do you find this out? I'm trying to understand if this is a widespread pain or just something I've been unlucky with. Four specific questions: 1. When an LLM API starts silently degrading, how do you currently find out? (Your own monitoring? User complaints? Checking the status page? Reddit?) 2. How long does it typically take you to confirm "this is the provider, not my code"? 3. If something told you before you noticed, that Claude API was showing elevated TTFT on Sonnet right now, would that change anything about how you operate? Or would you just retry and move on regardless? 4. Would you pay for an independent alert service that tells you when an LLM's behaviour has drifted, before your users notice? If this isn't actually a problem for you, I think that also would be the most useful answer I can get.


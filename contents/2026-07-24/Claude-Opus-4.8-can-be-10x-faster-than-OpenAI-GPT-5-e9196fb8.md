---
source: "https://www.peterbe.com/plog/claude-opus-is-10x-faster-than-openai-gpt-5-at-non-streaming-completions"
hn_url: "https://news.ycombinator.com/item?id=49040691"
title: "Claude Opus 4.8 can be 10x faster than OpenAI GPT-5"
article_title: "Claude Opus is 10x faster than OpenAI GPT 5 at non-streaming completions - Peterbe.com"
author: "peterbe"
captured_at: "2026-07-24T20:11:00Z"
capture_tool: "hn-digest"
hn_id: 49040691
score: 1
comments: 0
posted_at: "2026-07-24T19:41:49Z"
tags:
  - hacker-news
  - translated
---

# Claude Opus 4.8 can be 10x faster than OpenAI GPT-5

- HN: [49040691](https://news.ycombinator.com/item?id=49040691)
- Source: [www.peterbe.com](https://www.peterbe.com/plog/claude-opus-is-10x-faster-than-openai-gpt-5-at-non-streaming-completions)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T19:41:49Z

## Translation

タイトル: Claude Opus 4.8 は OpenAI GPT-5 より 10 倍高速になる可能性があります
記事のタイトル: Claude Opus は、非ストリーミング完了時に OpenAI GPT 5 より 10 倍高速です - Peterbe.com
説明: Claude は OpenAI gpt-5 よりもはるかに高速で、gpt-5-mini よりも高速です。

記事本文:
Claude Opus は非ストリーミング完了時に OpenAI GPT 5 より 10 倍高速です - Peterbe.com メインナビゲーションにスキップ
Claude Opus は非ストリーミング完了時に OpenAI GPT 5 より 10 倍高速です
この図はそれをうまく要約しています:
私のブログでは、この人気のブログ投稿に対して多くのコメントを受け取りました。年間で 28,000 件のブログ コメント。中にはひどく書かれていてわかりにくいものもあるので、AI にリライトを提案させています。ブログ投稿のコメントを AI に送信するコードは、実際には 3 回起動しています。1 回目は OpenAI gpt-5 で、1 回目は OpenAI gpt-5-mini で、もう 1 回目は Claude claude-opus-4.8 です。私は人間の目と判断力を使って結果を評価していますが、結果は同様に良好であると言えます。ほんのわずかな違いだけです。
驚くべきことは、OpenAI の gpt-5 が驚くほど遅いことです。これは claude-opus-4.8 よりも 10 倍近く遅いです。それはどうしたのですか？
gpt-5-mini と gpt-5 のレイテンシの差が大きいことも明らかです。執筆時点では、gpt-5.4 と gpt-5.4-mini の入力トークンの価格差は、0.75 ドルに対して 2.50 ドルです。それは3倍の差です。
API のプロンプトを構築している場合は、Claude を使用してください。
OpenAI を使用する必要がある場合は、安価で高速なミニ モデルを検討してください。
Claude を追加する前は、litellm を使用して OpenAI のモデルをラップしていました。コードは次のようになります。
応答 = litellm.completion(
モデル= "openai-gpt-5" 、
api_key=設定.OPENAI_API_KEY、
メッセージ=my_prompt_messages、
)
これとは異なり、ネイティブ OpenAI Python SDK を使用する場合、呼び出しは次のようになります。
client = openai.OpenAI(api_key=settings.OPENAI_API_KEY)
応答 = client.responses.create(
モデル= "gpt-5、
input=my_prompt_messages、
)
OpenAI SDK と litellm ラッパーを使用して比較した場合の速度の違いを測定しました。違いは次のようになります。

:
確かに、6 月に私がそうした通話を行ったのは 30 件強「だけ」でしたが、不思議なことに違いがありました。
私には、リテルムによって SDK を使用した場合と合計時間が異なる理由を理解するための複雑な知識がありません。 （私も気にしているかどうかはわかりません！）
いずれにせよ、私は litellm から離れ、Claude と OpenAI が提供する SDK のみを使用するつもりです。今年 litelm で見られた CVE を考えると、より安全だと感じます。
あなたのメールが公開されることはありません。
私のサイドプロジェクトをチェックしてください: That's Groce!

## Original Extract

Claude is much faster than OpenAI gpt-5 and also faster than gpt-5-mini

Claude Opus is 10x faster than OpenAI GPT 5 at non-streaming completions - Peterbe.com Skip to main navigation
Claude Opus is 10x faster than OpenAI GPT 5 at non-streaming completions
This picture summarizes it well:
Here on my blog, for this popular blog post I get a lot of comments. 28k blog comments over the years. Some of them are terribly written and hard to understand, so I let AI suggest a rewrite. That code that sends the blog post comment to AI, I actually fire off three times: once with OpenAI gpt-5 , once with OpenAI gpt-5-mini , and once with Claude claude-opus-4.8 . I use my human eyes and judgement to evaluate the results, and I can tell you they do equally well. Only the slightest differences.
The surprising thing is how amazingly slow OpenAI's gpt-5 is! It's nearly 10x slower than claude-opus-4.8 . What's up with that!?
It's also clear that the latency difference between gpt-5-mini and gpt-5 is significant. At the time of writing, the input token price difference between gpt-5.4 and gpt-5.4-mini is $2.50 compared to $0.75! That's a 3x difference.
If you're constructing a prompt the API, use Claude.
If you have to use OpenAI, consider the mini model because it's both cheaper and faster.
Before I added Claude, I used to use litellm to wrap OpenAI's models. The code looks like this:
response = litellm.completion(
model= "openai-gpt-5" ,
api_key=settings.OPENAI_API_KEY,
messages=my_prompt_messages,
)
Unlike, if you use the native OpenAI Python SDK the invocation looks like this:
client = openai.OpenAI(api_key=settings.OPENAI_API_KEY)
response = client.responses.create(
model= "gpt-5,
input=my_prompt_messages,
)
I measured the difference, in speed, where I compare using the OpenAI SDK versus the litellm wrapper and the difference looks like this:
Granted, in June I "only" did a bit over 30 of these calls, but strangely there's a difference!
I don't have the intricate knowledge to understand why the litellm makes the total time different from using the SDK. (not sure I care either!)
Either way, I'm moving away from litellm and only use the SDKs provided by Claude and OpenAI. Feels safer given the CVEs we've seen this year on litellm .
Your email will never ever be published.
Check out my side project: That's Groce!

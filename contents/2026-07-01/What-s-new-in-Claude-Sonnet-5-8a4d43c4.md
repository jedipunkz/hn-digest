---
source: "https://simonwillison.net/2026/Jun/30/claude-sonnet-5/"
hn_url: "https://news.ycombinator.com/item?id=48742426"
title: "What's new in Claude Sonnet 5"
article_title: "What's new in Claude Sonnet 5"
author: "ankitg12"
captured_at: "2026-07-01T05:31:04Z"
capture_tool: "hn-digest"
hn_id: 48742426
score: 5
comments: 0
posted_at: "2026-07-01T04:57:42Z"
tags:
  - hacker-news
  - translated
---

# What's new in Claude Sonnet 5

- HN: [48742426](https://news.ycombinator.com/item?id=48742426)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/30/claude-sonnet-5/)
- Score: 5
- Comments: 0
- Posted: 2026-07-01T04:57:42Z

## Translation

タイトル: クロード ソネット 5 の新機能
説明: 今朝、クロード ソネット 5 が発売されました。私はいつも「新機能」の開発者ドキュメントに直行します。なぜなら、開発者ドキュメントには公式の発表投稿よりも実用的な情報が含まれていることが多いからです。 …

記事本文:
クロード ソネット 5 の新機能
サイモン・ウィリソンのウェブログ
Claude Sonnet 5 の新機能 (経由) 今朝、Claude Sonnet 5 がリリースされました。私はいつも「新機能」の開発者ドキュメントに直行します。なぜなら、開発者ドキュメントには公式の発表投稿よりも実用的な情報が含まれていることが多いからです。
Anthropic はソネット 5 について「その性能は Opus 4.8 に近いが、価格は低い」と述べています。システム カードは、どのようにして米国政府にブロックされずにモデルをリリースできたのかを説明するのに役立ちます。
Sonnet 5 は、Mythos 5 よりもサイバー タスクの能力が大幅に劣ります。したがって、その安全対策は、Opus 4.7 および Opus 4.8 (Sonnet 5 よりも能力は高いが、Mythos 5 よりもはるかに低いモデル) に適用するものと似ています。
「新機能」API の変更点に注目してください。
サンプリング パラメーター pressure 、 top_p 、 top_k はサポートされなくなりました。
100 万のトークン コンテキスト ウィンドウと最大 128,000 の出力トークンがあります。
「Claude Sonnet 4.6 と同じツール セットとプラットフォーム機能」を備えています。
" Thinking": {type: "disabled"} を指定しない限り、適応的思考はデフォルトでオンになります。
価格は Sonnet 4.6 と同じです。入力 100 万件あたり 3 ドル、入力 100 万件あたり 15 ドルですが、8 月 31 日までは 2 ドル/10 ドルへの初回割引が適用されます。でも...
このモデルには新しいトークナイザーがあり、「同じ入力テキストがクロード ソネット 4.6 よりも約 30% 多くのトークンを生成します。」 - 事実上 30% の価格上昇。
Claude Token Counter ツールを使用して、新しいトークナイザーを試しました。いくつかの大きなドキュメントに対する私の結果は次のとおりです。
したがって、新しいトークンのコストは、英語では約 1.4 倍、スペイン語では 1.33 倍、Python コードでは 1.28 倍となり、簡体字中国語では実質的に同じコストになります。
こちらがペリカンです。特筆すべきことは何もありません。ソネット 5 はガチョウに似ていると考えています。
エージェントにビデオを録画してもらいます

ショットスクレーパービデオを使用したその作業のデモ - 2026 年 6 月 30 日
クロード コードを使用してブラウザで実行するためのメビウス 0.2B イメージ修復モデルの移植 - 2026 年 6 月 22 日
sqlite-utils 4.0rc1 は移行とネストされたトランザクションを追加 - 2026 年 6 月 21 日
これは、2026 年 6 月 30 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Claude Sonnet 5 came out this morning. I always head straight for the "what's new" developer docs because they tend to have more actionable information than the official announcement post. …

What's new in Claude Sonnet 5
Simon Willison’s Weblog
What's new in Claude Sonnet 5 ( via ) Claude Sonnet 5 came out this morning . I always head straight for the "what's new" developer docs because they tend to have more actionable information than the official announcement post.
Anthropic say of Sonnet 5 that "its performance is close to that of Opus 4.8, but at lower prices". The system card helps explain how they were able to release the model without being blocked by the US government:
Sonnet 5 is significantly less capable at cyber tasks than Mythos 5: its safeguards are thus similar to those we apply to Opus 4.7 and Opus 4.8 (models that are more capable than Sonnet 5 but much less capable than Mythos 5).
Of note from the "what's new" API changes:
Sampling parameters temperature , top_p , top_k are no longer supported.
It has a 1 million token context window and 128,000 maximum output tokens.
It features "the same set of tools and platform features as Claude Sonnet 4.6"
Adaptive thinking is on by default, unless you specify "thinking": {type: "disabled"} .
The pricing is the same as Sonnet 4.6: $3/million input, $15/million input, with an introductory discount to $2/$10 until 31st August. But...
The model has a new tokenizer, where "The same input text produces approximately 30% more tokens than on Claude Sonnet 4.6." - effectively a 30% price increase.
I used my Claude Token Counter tool to try out the new tokenizer. Here are my results for several larger documents:
So the new token is roughly 1.4x times more expensive for English, 1.33x for Spanish, 1.28x for Python code and effectively the same cost for Simplified Mandarin.
Here's the pelican . It's nothing to write home about. Sonnet 5 thinks it looks like a goose.
Have your agent record video demos of its work with shot-scraper video - 30th June 2026
Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code - 22nd June 2026
sqlite-utils 4.0rc1 adds migrations and nested transactions - 21st June 2026
This is a link post by Simon Willison, posted on 30th June 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.

---
source: "https://funnybench.lol"
hn_url: "https://news.ycombinator.com/item?id=48613679"
title: "FunnyBench – Can AI Models Tell Funny Jokes?"
article_title: "FunnyBench — Can AI tell a joke?"
author: "Mockapapella"
captured_at: "2026-06-21T01:03:20Z"
capture_tool: "hn-digest"
hn_id: 48613679
score: 3
comments: 1
posted_at: "2026-06-20T22:43:39Z"
tags:
  - hacker-news
  - translated
---

# FunnyBench – Can AI Models Tell Funny Jokes?

- HN: [48613679](https://news.ycombinator.com/item?id=48613679)
- Source: [funnybench.lol](https://funnybench.lol)
- Score: 3
- Comments: 1
- Posted: 2026-06-20T22:43:39Z

## Translation

タイトル: FunnyBench – AI モデルは面白いジョークを言えますか?
記事のタイトル: FunnyBench — AI はジョークを言えますか?
説明: 大規模言語モデルが実際にどれほど面白いかをテストするためのベンチマーク。ジョークを読んで投票し、どのモデルがそれを書いたかを確認してください。

記事本文:
😂 " />
面白いベンチ
投票する
リーダーボード
詳細
AIはジョークを言えるのか？
FunnyBench は単純な質問をします。AI モデルは面白いジョークを言えるでしょうか?それぞれのモデルは
「冗談を言ってください」という同じプロンプトが 10 回与えられました。ジョークを読んで、
面白いか面白くないかを判断してください。あなたの投票がライブリーダーボードを動かします。それぞれ聞いてみた
多様性を促進するためにモデルを何度もモデル化しましたが、それでも同じジョークを繰り返す人もいました。
投票するとモデルが明らかになります。
ジョークは、OpenRouter を通じてそのモデル カタログから正確な情報を使用して生成されました。
「冗談を言ってください」というプロンプト。生成で使用される温度 1 (サポートされている場合)、
120 秒のタイムアウト、プロバイダー フォールバックは無効、必須パラメーターは有効、
返されたモデル、プロバイダー、およびテキストが保存されました。トークン数とコストは、
ノイズを減らすために、内部に保存されますが表示されません。の
リーダーボードはベイジアン スコアを使用します。各モデルは全体の平均近くから始まり、
投票が入ると順位が変動するため、初期のランキングは生の面白いものよりも変動しにくくなります。
割合。また、OpenRouter からリクエストされたモデルと返されたモデルの両方も表示されます。
実際に実行されたモデルであるため、ベンチマークはテストされた内容を明確に示します。のために
推論モデルでは、利用可能な最も低い推論設定が使用されました。推論
痕跡はジョークの一部ではないため、意図的にキャプチャされませんでした
有権者に見せられました。ランは除外
主にテキスト、OpenRouter/ルーター/フロント エイリアス、検索、または
カスタムツールのバリアント、フローティング「最新」エイリアス、入手不可能な価格モデル、
重複した空きエイリアス、無効な空の出力またはオーバーサイズの出力、および
5回連続で電話に失敗した。公開されたセットには、保持ごとに 10 個の有効なジョークが保持されます。
モデル。
クインテン・リソウェによって建てられました。 FunnyBench に関するお問い合わせ、ビジネス上の質問、フィードバックについては、quinten.lisowe@thelisowe.com までご連絡ください。参照してください。

私の仕事の原石はこちら：https://thelisowe.com

## Original Extract

A benchmark for testing how funny large language models really are. Read a joke, vote, then see which model wrote it.

😂 " />
FunnyBench
Vote
Leaderboard
Details
Can AI tell a joke?
FunnyBench asks a simple question: can AI models tell funny jokes? Each model was
given the same prompt — “tell me a joke” — ten times. Read a joke and
decide if it’s funny or not. Your votes drive a live leaderboard. We asked each
model multiple times to encourage variety, but some still repeated the same joke.
The model is revealed after you vote.
Jokes were generated through OpenRouter from its model catalog using the exact
prompt “tell me a joke” . Generation used temperature 1 where supported,
a 120 second timeout, provider fallback disabled, required parameters enabled,
and the returned model, provider, and text were stored. Token counts and cost are
stored internally but not displayed, to reduce noise. The
leaderboard uses a Bayesian score: each model starts near the overall average and
moves as votes come in, which makes early rankings less jumpy than a raw funny
percentage. It also shows both the model requested from OpenRouter and the returned
model that actually ran, so the benchmark is explicit about what was tested. For
reasoning models, the lowest available reasoning setting was used; reasoning
traces were intentionally not captured because they are not part of the joke
shown to voters. The run excluded
models not primarily meant for text, OpenRouter/router/front aliases, search or
custom-tool variants, floating “latest” aliases, unavailable-price models,
duplicate free aliases, invalid empty or oversized outputs, and any model that
failed five calls in a row. The published set keeps ten valid jokes per retained
model.
Built by Quinten Lisowe. Please reach out to quinten.lisowe@thelisowe.com for inquiries, business questions or feedback about FunnyBench. See more of my work here: https://thelisowe.com

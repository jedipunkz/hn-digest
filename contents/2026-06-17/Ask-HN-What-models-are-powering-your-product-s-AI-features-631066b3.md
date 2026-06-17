---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48576393"
title: "Ask HN: What models are powering your product's AI features?"
article_title: ""
author: "user-"
captured_at: "2026-06-17T21:42:28Z"
capture_tool: "hn-digest"
hn_id: 48576393
score: 2
comments: 0
posted_at: "2026-06-17T20:29:08Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: What models are powering your product's AI features?

- HN: [48576393](https://news.ycombinator.com/item?id=48576393)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T20:29:08Z

## Translation

タイトル: HN に質問: 製品の AI 機能を強化しているモデルは何ですか?
HN テキスト: コーディング アシスタントやエージェントではなく、明らかに最新のフロンティア モデルを使用するだけの大変な作業のようなものでもありません。ここで話しているのは、ユーザーが直接操作する日常的な AI 機能です。コンテンツの生成、書き換え、レコメンデーション、ワークフロー ヘルパー、コンテキストの提案などです。私のアプリでは、ワークロードがかなり構造化されているため、主に OpenRouter 経由で Gemini Flash を使用しています。それは十分に機能しますが、現在利用可能なモデルは数十あり、ほとんどのチームがそれらをどのように評価しているかはわかりません。人々は適切な評価スイートを構築していますか?コストとレイテンシを比較しますか?いくつかのモデルをテストして、最も安価で十分なモデルを選択しますか?ほとんどの人は最後のバケツに当てはまると思いますが、少なくとも他の人がどのような結果に陥ったのかを聞きたいと思います。 Flash の他に、qwen または minimax モデルを使用するとかなりうまく機能することがわかりました。本番環境で何を実行しているのか、またそれをどのように選択したのか興味があります。

## Original Extract

Not coding assistants or agents, and not things that obviously would just use the latest frontier model like really hard work. I'm talking about the daily AI features users interact with directly: generating content, rewriting things, recommendations, workflow helpers, contextual suggestions, etc. For my app I mostly use Gemini Flash via OpenRouter because the workloads are fairly structured. It works well enough, but there are now dozens of models available and I'm not sure how most teams are evaluating them. Are people building proper eval suites? Comparing cost/latency? Testing a handful of models and picking the cheapest one that's good enough? I think most people fall in that last bucket and I want to at the least hear what others landed on. Besides Flash I found using a qwen or minimax model worked fairly well. Curious what you're running in production and how you chose it.


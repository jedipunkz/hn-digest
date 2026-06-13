---
source: "https://stephen.bochinski.dev/blog/2026/06/13/ai-coding-at-home-without-going-broke/"
hn_url: "https://news.ycombinator.com/item?id=48518969"
title: "AI Coding at Home Without Going Broke"
article_title: "AI Coding at Home Without Going Broke | Stephen Bochinski"
author: "sbochins"
captured_at: "2026-06-13T18:35:24Z"
capture_tool: "hn-digest"
hn_id: 48518969
score: 82
comments: 84
posted_at: "2026-06-13T16:45:03Z"
tags:
  - hacker-news
  - translated
---

# AI Coding at Home Without Going Broke

- HN: [48518969](https://news.ycombinator.com/item?id=48518969)
- Source: [stephen.bochinski.dev](https://stephen.bochinski.dev/blog/2026/06/13/ai-coding-at-home-without-going-broke/)
- Score: 82
- Comments: 84
- Posted: 2026-06-13T16:45:03Z

## Translation

タイトル: 破産せずに自宅で AI コーディングを実現
記事のタイトル: 破産せずに自宅で AI コーディング |スティーブン・ボチンスキー
説明: 会社のような費用をかけずに自宅で AI コーディングを行う方法は 3 つあります。どれが適しているかは主に、来年のハードウェアをどれだけ信頼できるかによって決まります。

記事本文:
コンテンツにスキップ
スティーブン・ボチンスキー
について
ブログ
アプリ
破産せずに自宅で AI コーディングを実現
There are three ways to do AI coding at home without spending like a company, and which one fits depends mostly on how much you trust the next year of hardware and model releases. 1 つ目はセルフホストです。マシンを購入し、オープンソース モデルをローカルで実行すれば、その後はトークンごとに料金を支払う必要はありません。 The upfront cost is steep and the models you can actually run at home are weaker than what the frontier labs ship, so this only pays off if you can keep the rig busy with long running tasks where a slower, cheaper model grinds away overnight.ほとんどの人は、負荷の多い家庭用マシンを維持することはできませんし、今日購入したハードウェアは 1 年後には悪い賭けになるかもしれません。
2 つ目は、ハードウェアを省略して、同じオープンソース モデルをプロバイダーから API 料金でレンタルすることです。ほとんどの人にとって、これは正しい判断です。 You avoid putting thousands of dollars on one GPU setup while configurations are still in flux, you skip the work of squeezing long running performance out of an open model, and you can switch to whatever is cheaper or better next month without reselling a box. OpenRouter のようなものでは、1 行の変更に近い変更が行われます。
3 つ目は、OpenAI と Anthropic からのフロンティア サブスクリプションを最小化して最大化することです。月額約 400 ドルのプランでは、定価で約 2,800 ドルの API 使用量を購入できます。これは、上限に達するまで実質的にお買い得です。プランは従量制であり、大規模な AI ネイティブ ワークフローは含まれるトークンを高速で処理します。これらは手動で運転する仕事には適していますが、一日中稼働するエージェントのエンジンとしては不十分です。
私が見た中で最もうまく機能するのは、最後の 2 つをブレンドしたものです。真剣に考えたり仕様を書いたりするためにフロンティアのサブスクリプションをいくつか維持し、料金を支払います

小さな機械部品を処理するためのオープンソース モデルの API レート。 Lean on spec driven development so the expensive models produce the plan and the cheap ones fill it in. Do that well and you can build what a team of twenty engineers would put out in a month for around a thousand dollars.
© 2026 スティーブン・ボチンスキー。無断転載を禁じます。

## Original Extract

There are three ways to do AI coding at home without spending like a company, and which one fits depends mostly on how much you trust the next year of hardwa...

Skip to content
Stephen Bochinski
About
Blog
Apps
AI Coding at Home Without Going Broke
There are three ways to do AI coding at home without spending like a company, and which one fits depends mostly on how much you trust the next year of hardware and model releases. The first is to self host. You buy the machine, run open source models locally, and pay nothing per token after that. The upfront cost is steep and the models you can actually run at home are weaker than what the frontier labs ship, so this only pays off if you can keep the rig busy with long running tasks where a slower, cheaper model grinds away overnight. Most people can’t keep a home machine that loaded, and the hardware you buy today may look like a bad bet in a year.
The second is to skip the hardware and rent those same open source models from a provider at API rates. For most people this is the right call. You avoid putting thousands of dollars on one GPU setup while configurations are still in flux, you skip the work of squeezing long running performance out of an open model, and you can switch to whatever is cheaper or better next month without reselling a box. Something like OpenRouter makes the move close to a one line change.
The third is to min-max the frontier subscriptions from OpenAI and Anthropic. Around $400 a month of plans buys roughly $2800 of API usage at list prices, which is a real bargain right up until you hit the ceiling. The plans are metered, and any large AI native workflow will chew through the included tokens fast. They shine for the work you drive by hand and fall short as the engine for an agent running all day.
What I’ve seen work best is a blend of the last two. Keep a couple of frontier subscriptions for the hard thinking and the spec writing, and pay API rates for open source models to handle the small mechanical pieces. Lean on spec driven development so the expensive models produce the plan and the cheap ones fill it in. Do that well and you can build what a team of twenty engineers would put out in a month for around a thousand dollars.
© 2026 Stephen Bochinski. All rights reserved.

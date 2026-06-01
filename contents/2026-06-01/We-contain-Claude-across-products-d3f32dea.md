---
source: "https://simonwillison.net/2026/May/30/how-we-contain-claude/"
hn_url: "https://news.ycombinator.com/item?id=48346040"
title: "We contain Claude across products"
article_title: "How we contain Claude across products"
author: "Brajeshwar"
captured_at: "2026-06-01T00:29:26Z"
capture_tool: "hn-digest"
hn_id: 48346040
score: 3
comments: 0
posted_at: "2026-05-31T14:41:18Z"
tags:
  - hacker-news
  - translated
---

# We contain Claude across products

- HN: [48346040](https://news.ycombinator.com/item?id=48346040)
- Source: [simonwillison.net](https://simonwillison.net/2026/May/30/how-we-contain-claude/)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T14:41:18Z

## Translation

タイトル: 製品全体にクロードが含まれています
記事のタイトル: 製品全体にクロードを含める方法
説明: サンドボックス製品に関してよくある不満は、製品が完全に文書化されていることはほとんどなく、詳細な文書がないため、どこまでできるのかを知るのが難しいということです。

記事本文:
製品全体にクロードを含める方法
サイモン・ウィリソンのウェブログ
製品全体にクロードを含める方法。サンドボックス製品に関して私がよく抱く不満は、それらが完全に文書化されていることはほとんどなく、詳細な文書が存在しない場合、それらをどの程度信頼できるのかを知るのが難しいということです。
Anthropic は、さまざまなサンドボックス技術が Claude.ai 、Claude Code、Cowork でどのように機能するかについての素晴らしい概要を公開しました。
エージェントがプロセスのサンドボックス、VM、ファイルシステムの境界、出口制御を使用して動作できる場所と方法を制限します。目標は、エージェントが到達できる範囲に厳密な境界を設定することです。たとえば、資格情報がサンドボックスに入らない場合、その原因がユーザー、「クリエイティブ」パスを見つけたモデル、または攻撃者のいずれであるかに関係なく、資格情報を抽出することはできません。
Claude.ai は gVisor を使用します。 Claude Code はローカルで実行され、macOS では Seatbelt を使用し、Linux では Bubblewrap を使用します。 Claude Cowork は完全な VM (macOS では Apple の仮想化フレームワーク、Windows では HCS) を実行します。
ここには、以前ここで取り上げた api.anthropic.com/v1/files 漏洩ベクトルなど、彼らが見逃したリスクに関する興味深い話など、多くの内容が含まれています。
これを見て、Anthropic のオープン ソース srt (Anthropic Sandbox Runtime) ツールをもう一度検討する時期が来たことを思い出しました。このツールは十分に成熟しているので、適切に試してみる準備ができていることがわかります。
Claude Opus 4.8: 「控えめだが目に見える改善」 - 2026 年 5 月 28 日
Anthropic と OpenAI は製品と市場の適合性を見つけたと思います - 2026 年 5 月 27 日
AIに関する教皇レオ14世の回勅についてのメモ - 2026年5月25日
これは、2026 年 5 月 30 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

A complaint I often have about sandboxing products is that they are rarely thoroughly documented, and in the absence of detailed documentation it's hard to know how much I can …

How we contain Claude across products
Simon Willison’s Weblog
How we contain Claude across products . A complaint I often have about sandboxing products is that they are rarely thoroughly documented , and in the absence of detailed documentation it's hard to know how much I can trust them.
Anthropic just published a fantastic overview of how their various sandbox techniques work across Claude.ai , Claude Code, and Cowork.
We constrain where and how an agent can act with process sandboxes, VMs, filesystem boundaries, and egress controls. The goal is to set a hard boundary on what an agent can reach. For example, if credentials never enter the sandbox, they can't be exfiltrated, regardless of whether the cause is a user, a model finding a “creative” path, or an attacker.
Claude.ai uses gVisor. Claude Code, run locally, uses Seatbelt on macOS and Bubblewrap on Linux. Claude Cowork runs a full VM (Apple's Virtualization framework on macOS, HCS on Windows).
There's a lot in here, including some interesting stories of risks they missed such as the api.anthropic.com/v1/files exfiltration vector covered here previously .
This reminded me it's time I took another look at Anthropic's open source srt (Anthropic Sandbox Runtime) tool - it's mature enough know that I'm ready to give it a proper go.
Claude Opus 4.8: "a modest but tangible improvement" - 28th May 2026
I think Anthropic and OpenAI have found product-market fit - 27th May 2026
Notes on Pope Leo XIV's encyclical on AI - 25th May 2026
This is a link post by Simon Willison, posted on 30th May 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.

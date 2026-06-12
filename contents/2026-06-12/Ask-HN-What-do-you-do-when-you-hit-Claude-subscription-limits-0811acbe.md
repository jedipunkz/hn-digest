---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48506908"
title: "Ask HN: What do you do when you hit Claude subscription limits?"
article_title: ""
author: "sermakarevich"
captured_at: "2026-06-12T18:46:18Z"
capture_tool: "hn-digest"
hn_id: 48506908
score: 1
comments: 0
posted_at: "2026-06-12T17:26:23Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: What do you do when you hit Claude subscription limits?

- HN: [48506908](https://news.ycombinator.com/item?id=48506908)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T17:26:23Z

## Translation

タイトル: HN に質問: クロードのサブスクリプション制限に達した場合はどうしますか?
HN テキスト: しばらく前、私は複数のクロード コーダーを実行することに興奮しました。これには、対話型セッションから claude -p に移行し、Beads を追加し、その後、単純な Python オーケストレーター、UI、ask_user MCP、および Telegram 統合が必要でした。コーダーをノンストップで実行できるようになりましたが、頻繁に使用すると、Sonnet モデルでワーカーを使用している場合でも、10 分以内にサブスクリプションの制限を消費してしまいます。しばらくの間、私はいくつかのサブスクリプションを持っていましたが、日中はサブスクリプションが消費されたら切り替えることができました。しかし、まだ十分ではありません。 6 月 15 日月曜日、Anthropic は claude -p の制限の計算方法を変更する予定です。巷の噂では、200 ドルのサブスクリプション料金内で入手できるトークンは、API コストで 8 ～ 12,000 ドルに相当するとのことです。問題は、サブスクリプションだけではコーディングに十分ではない場合はどうするかということです。私が現在テストしているアプローチの 1 つは、合計 36 GB のメモリを備えた 2 枚の GPU カードを備えたローカル サーバーにデプロイされた qwen3.6:36B モデルを使用して、フリート (エージェント swarm オーケストレーター アプリ) にオープンコード コーダーを追加しました。私は Fable 5 モデルに話しかけ、目標を説明します。 Fable は調査を行い、問題を qwen が処理できる小さくて単純なサブタスクに分割し、詳細なサブタスクの説明を書き、Beads で必要な依存関係を作成します。 qwen3.6 を使用した Opencode が機能します。表 5 では作業を検証し、必要に応じて追加のタスクを作成します。このアプローチはかなり時間がかかりますが、週末に一晩中ノンストップで実行できます。クラウドで安価なモデルをテストすることも検討しています。あなたのアプローチは何ですか?サブスクリプションで提供される以上のものが必要な場合はどうしますか?いつかサブスクリプションがなくなった場合のプランBは何ですか。

## Original Extract

Some time ago I got excited about running multiple Claude coders. This required moving from interactive sessions to claude -p, adding Beads, later a simple Python orchestrator, a UI, an ask_user MCP, and Telegram integration. I can now run coders non-stop, but with intense usage I consume subscription limits within 10 minutes, even when I use workers with the Sonnet model. For some time I have had a few subscriptions, and during the day I could switch subscriptions when one was consumed. But it's still nowhere near enough. On Monday, Jun 15, Anthropic plans to change the way they calculate limits for claude -p. Rumor on the street is that the tokens we get within the $200 subscription fee are equivalent to $8-12K in API cost. The question is: what do you do when the subscription is not enough for coding? One approach I am testing right now: I added an opencode coder to my fleet (agent swarm orchestrator app), using a qwen3.6:36B model deployed on a local server with 2x GPU cards with 36GB of memory in total. I talk to the Fable 5 model, describe the goal; Fable does the investigation, slices the problem into small and simple subtasks that qwen can handle, writes detailed subtask descriptions, and creates the required dependencies in Beads. Opencode with qwen3.6 does the job. Fable 5 verifies the work and creates additional tasks if required. This approach is significantly slower, but I can run it non-stop, overnight, over weekends. I'm also considering testing some cheap model in the cloud. What is your approach? What do you do when you need more than the subscription provides? What the plan B for case when subscription is gone one day.


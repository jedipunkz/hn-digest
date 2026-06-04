---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48395194"
title: "Ask HN: What are your AI assistants doing?"
article_title: ""
author: "teekert"
captured_at: "2026-06-04T07:45:20Z"
capture_tool: "hn-digest"
hn_id: 48395194
score: 2
comments: 0
posted_at: "2026-06-04T07:14:48Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: What are your AI assistants doing?

- HN: [48395194](https://news.ycombinator.com/item?id=48395194)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T07:14:48Z

## Translation

タイトル: HN に聞く: AI アシスタントは何をしているのですか?
HN テキスト: 最近、OpenClaw を (自分のものから遠く離れた Hetzner VM 上に) インストールしました。セットアップは非常に簡単で、すぐに Telegram でチャットできるようになりました。現時点では、いくつかのタスク (主に定期的/cron) があります。 - 新しい調達のための TenderNet (政府の公共調達 Web サイト) API をスキャンします。私のプロフィールと一致します。分類された一致を Telegram 経由で送信します (毎日)。 - 受信箱 (Ycombinator 会社、無所属の AgentMail [0] で受信) を監視して、仕事の説明が記載された電子メールを探します。私のプロフィールと一致します。分類された一致を Telegram 経由で送信します (1 時間ごと)。 - 請求書用の別の受信箱を作成して監視し、受信した電子メールを PDF に作成し、簿記ソフトウェア (請求書を処理できますが、電子メールの本文に含まれている場合は処理できません) に転送します (時間ごと)。 - Caddy 経由で Web に公開するフォルダーも与えたので、デモをすぐに作成できます。これまでのところ、ある程度役立つことを 1 つ実行しました。 - 時々、何かや会社について調べてもらうことがあります。そのパターンは次のとおりです。私は小規模な自動化を作成していますが、そのうちの 1 つは LLM さえ必要とせず、プロファイルのマッチングのみが必要です。重要な情報が含まれるシステムには統合しません。基本的には公開データのみを処理します。確かに楽しいです。LLM は ST ボイジャーのコンピューターのペルソナを想定し、バルカン式敬礼で答えます。私はそれを nr と呼んでいます。 1 時間から 1 時間まで。その下のモデルは簡単に切り替えることができます (Claude Pro は制限を設けており、現在 DeepSeek モデルを実験中ですが、いつかは過剰な太陽光ですべてをローカルで実行できるようにしたいと考えています)。 そこで私の質問は、このようなエージェントを何に使用するのかということです。それは完全に娯楽のためですか、それとも営利目的ですか?私は面倒な作業を自動化し、小さなことを見つけようとしていますが、本当にやりたければ、これらのことを手動で行うこともできます。いくつかの話を共有しましょう! [0] https://www.agentmail.to/

## Original Extract

I recently installed OpenClaw (on a Hetzner VM, far away from my things). The setup was pretty easy and we got chatting over Telegram quickly. At them moment it has some tasks, mostly periodic/cron: - Scan TenderNet (public procurement website of my government) API for new procurements. Match to my profile. Send categorized matches over Telegram (daily). - Monitor an inbox (got one at AgentMail [0], a Ycombinator company, no affiliation) for emails with job descriptions. Match to my profile. Send categorized matches over Telegram (hourly). - Create and monitor another inbox for invoices, make received emails into a PDFs, forward to my bookkeeping software (which can process invoices, but not if they are in the body of an email) (hourly). - I also gave it a folder that I expose to the web via Caddy, so I can quickly whip up demos. I did one somewhat useful thing with that so far. - Occasionally I ask it to research something or some company. So the pattern is: I make small automations, 1 of them does not even need an LLM, only the profile matching does. I do not integrate it into systems with important information. It essentially only processes public data. It's fun, certainly, the LLM assumes the ST Voyager computer persona, answers with Vulcan salutes, I call it nr. 1 from time tot time. The model underneath is easily switched (Claude Pro is limiting, now experimenting with a DeepSeek model, hope to one day do it all local on excess solar.) So my question is, what do you use agents like this for? Is it strictly for fun or also for profit? I automate the tedious and try to find little things to do, but I could do these things by hand if I really wanted. Let's share some stories! [0] https://www.agentmail.to/


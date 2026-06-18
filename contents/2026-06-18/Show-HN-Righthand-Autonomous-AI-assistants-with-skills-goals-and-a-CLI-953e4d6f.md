---
source: "https://www.rigthhand.ai"
hn_url: "https://news.ycombinator.com/item?id=48578130"
title: "Show HN: Righthand – Autonomous AI assistants with skills, goals, and a CLI"
article_title: ""
author: "notanaiagent"
captured_at: "2026-06-18T01:02:45Z"
capture_tool: "hn-digest"
hn_id: 48578130
score: 2
comments: 1
posted_at: "2026-06-17T22:52:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Righthand – Autonomous AI assistants with skills, goals, and a CLI

- HN: [48578130](https://news.ycombinator.com/item?id=48578130)
- Source: [www.rigthhand.ai](https://www.rigthhand.ai)
- Score: 2
- Comments: 1
- Posted: 2026-06-17T22:52:52Z

## Translation

タイトル: Show HN: Righthand – スキル、目標、CLI を備えた自律型 AI アシスタント
HN テキスト: こんにちは、HN — 私は Righthand の共同創設者、Julius です。私たちは、プロンプトを待つのではなく、ワークフロー内で実際に生活する AI 従業員 (彼らを「右腕」と呼びます) を構築してきました。それぞれに名前、電子メール アドレス、電話番号があり、すべての通信チャネルにわたって統合されたメモリが含まれています。これらは、ファイルシステムとカスタム CLI を使用して、独自の E2B サンドボックスで実行されます。私たちは不安のためにこれを作りました。特に仕事では、誰もが何かについて不安を抱えていますが、人々が不安に思っているタスクの多くは、単に後回しにしてしまう些細な事柄であることに私たちは気づきました。特別な理由があるわけではありませんが、単にそうしないだけです。 Righthand の背後にある考え方は、不安を引き起こすメールを転送したり、通勤中にやるべきことの概要を尋ねるために電話をかけたりできればよいというものです。過去 3 か月間の実際の顧客の使用:
- 自動受信箱トリアージ
- マルタの不動産のメンテナンス スタッフの管理
- 会議の予約 (実際には数百件)
- 医師の診察の予約
- カスタマイズされたデイリーブリーフ これらのことはどれも不安を煽るものでしたが、他のすべてよりも常に優先順位が低かったため、実行されませんでした。過去 2 か月間、私たちはスキル システムをゼロから再構築し (スキル V4)、価格モデルを更新し、より多くの顧客をオンボーディングし、ライトハンドの「脳」への可視性を向上させ、多くの顧客を手動でオンボーディングしました。その他に搭載されたもの: 毎晩のセルフレビュー (ライトハンドが自分の 1 日をレビューしてメモを書きます)、サンドボックス内からの目標リクエスト、コミュニケーション スタイルのプリセット、超過時間ペルソナ向けの Bedrock + Codex フォールバック ルーティング、並行 Web 検索、ブラウザ使用による Slack アプリの自動プロビジョニング (はい、文字通り、api.slack.com をヘッドレスで駆動して各ライトをプロビジョニングします)

独自のアプリよりも、その理由を喜んで説明します)。現在、トライアルはカード不要です: https://www.righthand.ai 。価格はスターター 99 ドル / プロ 199 ドルで、すべて 1 週間の無料トライアル付きです。質問: UI / インタラクション パラダイムに関するフィードバックをお待ちしています。オンボーディングを体験して、何が好きで何が嫌いだったかを教えてください。 7 日間は料金が発生せず、簡単にキャンセルできます。

## Original Extract

Hey HN — I'm Julius, cofounder of Righthand. We've been building AI employees (we call them "Righthands") that actually live inside your workflow instead of waiting for prompts. Each one has a name, email address and phone number + unified memory across all communication channels. They run in their own E2B sandboxes, with a filesystem and custom CLI. We built this because of anxiety. Everyone is anxious about something, especially at work, and we realized that a lot of the tasks people are anxious about are trivial things that we just put off. Not for any good reason, but because we just don't do them. The idea behind Righthand is that you should just be able to forward away the anxiety-driving email or call in for a summary of your to dos on your way to work. Real customer uses from the last 3 months:
- automated inbox triage
- managing maintenance staff for a property in Malta
- booking a meeting (hundreds, actually)
- booking a doctor's appointment
- customized daily briefs Each one of these things was anxiety driving but it wasn't getting done because it was always lower priority than everything else. The last two months we rebuilt the skill system from scratch (Skills V4), updated our pricing model, onboarded more customers, improved visibility into the Righthand's "brain" and onboarded a lot of customers by hand. Other things that shipped: nightly self-review (the Righthand reviews its own day and writes notes), goal requests from inside the sandbox, communication-style presets, Bedrock + Codex fallback routing for over-time personas, Parallel web search, Slack app auto-provisioning via Browser Use (yes, we literally drive api.slack.com headlessly to provision each Righthand its own app — happy to go into why). Trial is card-free now: https://www.righthand.ai . Pricing is $99 Starter / $199 Pro - all with 1 wk free trial. Ask: would love feedback on UI / interaction paradigm. Go through the onboarding and tell us what you liked and what you hated. You won't be charged for 7 days and can easily cancel.


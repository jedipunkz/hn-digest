---
source: "https://www.athenic.com:443/"
hn_url: "https://news.ycombinator.com/item?id=48480928"
title: "Show HN: Athenic – Why you can't do data analysis with Claude"
article_title: "Athenic — Analyze and automate data & work"
author: "jaredzhao"
captured_at: "2026-06-10T18:58:59Z"
capture_tool: "hn-digest"
hn_id: 48480928
score: 3
comments: 0
posted_at: "2026-06-10T18:50:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Athenic – Why you can't do data analysis with Claude

- HN: [48480928](https://news.ycombinator.com/item?id=48480928)
- Source: [www.athenic.com:443](https://www.athenic.com:443/)
- Score: 3
- Comments: 0
- Posted: 2026-06-10T18:50:27Z

## Translation

タイトル: Show HN: Athenic – クロードではデータ分析ができない理由
記事のタイトル: Athenic — データと作業の分析と自動化
説明: Athenic は、データ分析と作業自動化のための AI エージェントです。データを接続し、平易な英語でチャットし、フォーチュン 500 企業向けに構築されたダッシュボード、レポート、自動化を提供します。
HN テキスト: こんにちは、HN、私はジャレッドです。私は 2020 年からデータ ツールを構築してきました。Polyture、その後 AskEdith、そして現在は Athenic: 自然言語で質問し、グラフ/ダッシュボードを取得して、それを自動化します。 Postgres、Salesforce、Google 広告などに接続します。 「クロードをデータベースにリンクするだけ」という皆さんへ: ビジネスの現場で発生する、矛盾する定義と分析の混乱を想像してみてください。 「私たちの収入はいくらですか？」と尋ねてください。 2 日おきに 2 回、または別のモデルに変更します。同じ結果が得られるという保証はありません。次に、それを社内の技術者以外のユーザー全員に提供することを想像してください。モデルの問題ではありません。私たちはこれを苦労して学びました。私たちが 2022 年にここで AskEdith を立ち上げたとき、あなたは私たちにこう言いました ( https://news.ycombinator.com/item?id=33435361 ): 「SQL をチェックする必要がある」、「信頼がすべて」、「答えは一貫性がない」。あなたは正しかったです。現在、Athenic は KPI と式をセマンティック モデルで決定的に定義しています。セマンティック モデルは、決定性と精度を保証しながら、複雑な分析を構成できるモジュール式の構成可能なユニットで構成されています。 LLM の唯一の責任は、質問を解釈することです (技術者以外のユーザーでも再確認できます)。 `revenue = sum(order_total −refunds) where status = 'completed'` 収益を求めると、毎回全員が同じ数字を取得します。トップクラスのスタートアップ企業やフォーチュン 500 企業との協力から 3 年間学び、私たちは 2.0 を出荷したばかりです。

チャットからインサイトに加え、スケジュールに従って実行され、メールに届くダッシュボードと自動化。私たちが間違っていると言ってください。

記事本文:
Athenic — データと作業を分析および自動化する

## Original Extract

Athenic is an AI agent for analyzing data and automating work. Connect your data, chat in plain English, and ship dashboards, reports, and automations — built for startups to Fortune 500.

Hey HN, I'm Jared. I’ve been building data tools since 2020. Polyture, then AskEdith, now Athenic: ask a question in natural language, get a chart/dashboard, then automate it. Connects to Postgres, Salesforce, Google Ads, whatever. To everyone that says "just link Claude to your db”: imagine the chaos of conflicting definitions and analysis that would show up in a business setting. Ask “what’s our revenue?” twice, two days apart or to a different model. There’s no guarantee that you’ll get the same results. Now imagine giving that to all of the non-technical users at your company. It's not a model problem. We learned this the hard way. When we launched AskEdith here in 2022, and you told us ( https://news.ycombinator.com/item?id=33435361 ): “you'll still have to check the SQL,” “trust is everything,” “answers won’t be consistent”. You were right. Now, Athenic defines KPIs and formulas deterministically in a Semantic Model. The Semantic Model is made up of modular, composable units that can make up complex analysis, while guaranteeing determinism and accuracy. The LLM’s only responsibility is interpreting your question (which even non-technical users can double check). `revenue = sum(order_total − refunds) where status = 'completed'` Ask for revenue and everyone gets the same number, every time. Three years of learnings from working with top startups and Fortune 500 companies later, we just shipped 2.0. Chat-to-insight, plus dashboards and automations that run on a schedule and land in your email. Tell us we're wrong.

Athenic — Analyze and automate data & work

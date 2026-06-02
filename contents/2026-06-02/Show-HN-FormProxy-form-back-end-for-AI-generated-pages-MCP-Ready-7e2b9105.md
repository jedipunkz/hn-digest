---
source: "https://www.formproxy.com/"
hn_url: "https://news.ycombinator.com/item?id=48357884"
title: "Show HN: FormProxy – form back end for AI-generated pages – MCP Ready"
article_title: "FormProxy — The Form Backend for AI-Generated Pages | by Innerkore"
author: "gagan2020"
captured_at: "2026-06-02T00:37:54Z"
capture_tool: "hn-digest"
hn_id: 48357884
score: 3
comments: 0
posted_at: "2026-06-01T15:09:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: FormProxy – form back end for AI-generated pages – MCP Ready

- HN: [48357884](https://news.ycombinator.com/item?id=48357884)
- Source: [www.formproxy.com](https://www.formproxy.com/)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T15:09:46Z

## Translation

タイトル: HN を表示: FormProxy – AI 生成ページのフォーム バックエンド – MCP Ready
記事のタイトル: FormProxy — AI 生成ページのフォーム バックエンド | by インナーコア
説明: AI がページを構築します。 FormProxy はフォームを接続します。 URL をドロップすると、すぐに送信、通知、統合を取得できます。クロード、GPT-4、および任意の LLM の MCP スキル。
HN テキスト: AI コード ツールを使用して構築する際にも、同じ問題に何度も遭遇しました。生成された HTML は見栄えがしますが、<form> にはバックエンドがありません。 Formspree に手を伸ばすか、サーバーレス関数を作成するか、壊れたまま出荷するかのどちらかです。 FormProxy は 1 ステップで接続します。ページ URL をドロップし、POST エンドポイントを取得し、送信は Slack、Webhook、または Google Sheets にルーティングされます。 (Google スプレッドシートは現在信頼性が低く、Google の OAuth アプリのレビューを待っており、予想よりも時間がかかっています。Webhook と Slack は運用準備が整っています。) 私が最も興味がある角度: コード生成中に LLM がフォーム エンドポイントを自分でプロビジョニングできる MCP スキルがあります。 Claude または GPT-4 は、セッション中に FormProxy を呼び出し、作成中の HTML にアクション URL を挿入できます。これにより、ページは手動で配線することなく、動作するフォームとともに出荷されます。 Formspree だけではだめなのか、MCP スキルが実際にどのように機能するのか、これに適した価格モデルは何かなど、挑戦していただけると嬉しいです。私はサブミッションごととエンドポイントごとを行ったり来たりしてきました。 formproxy.com — 無料利用枠が利用可能。

記事本文:
FormProxy — AI 生成ページのフォーム バックエンド | by インナーコア

## Original Extract

AI builds the page. FormProxy wires the form. Drop a URL, get submissions, notifications, and integrations instantly. MCP skill for Claude, GPT-4, and any LLM.

I kept running into the same problem building with AI code tools: the generated HTML looks great, but the <form> has no backend. You either reach for Formspree, write a serverless function, or ship it broken. FormProxy wires it up in one step. Drop a page URL, get a POST endpoint, and submissions route to Slack, webhooks, or Google Sheets. (Google Sheets is currently unreliable — waiting on Google's OAuth app review, which is taking longer than expected. Webhook and Slack are production-ready.) The angle I'm most interested in: there's an MCP skill so LLMs can provision the form endpoint themselves during code generation. Claude or GPT-4 can call FormProxy mid-session, inject the action URL into the HTML it's writing, and the page ships with a working form — no manual wiring. Happy to be challenged on: why not just Formspree, how the MCP skill actually works in practice, and what the right pricing model is for this. I've been going back and forth on per-submission vs. per-endpoint. formproxy.com — free tier available.

FormProxy — The Form Backend for AI-Generated Pages | by Innerkore

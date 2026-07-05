---
source: "https://wjhowland376-code.github.io/Pulse/"
hn_url: "https://news.ycombinator.com/item?id=48794553"
title: "I built an AI that diagnoses PagerDuty incidents and posts fixes to Slack"
article_title: "Pulse — AI Incident Response"
author: "Pulse-AI"
captured_at: "2026-07-05T15:06:38Z"
capture_tool: "hn-digest"
hn_id: 48794553
score: 2
comments: 0
posted_at: "2026-07-05T14:29:06Z"
tags:
  - hacker-news
  - translated
---

# I built an AI that diagnoses PagerDuty incidents and posts fixes to Slack

- HN: [48794553](https://news.ycombinator.com/item?id=48794553)
- Source: [wjhowland376-code.github.io](https://wjhowland376-code.github.io/Pulse/)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T14:29:06Z

## Translation

タイトル: PagerDuty のインシデントを診断し、修正を Slack に投稿する AI を構築しました
記事のタイトル: Pulse — AI インシデント対応
説明: Pulse はシステムを 24 時間 365 日監視します。インシデントが発生すると、誰かが呼び出される前に、AI がそれを診断し、回答を Slack に投稿します。

記事本文:
パルス
仕組み
特長
価格設定
アクセスする
AIインシデント対応
インシデントが発生すると、Pulse は Claude AI で診断し、ワンクリックで解決またはエスカレーションして回答を Slack に投稿します。ほとんどの夜、人間は起きる必要はありません。
ウェブサイトを24時間365日閲覧
最終チェック: 0.4 秒前 — すべてクリア
昨夜、あなたのチームが寝ている間に
一つの出来事が始まりから終わりまで — 目覚めた人間はいない
PagerDuty は午前 3 時 4 分にトリガーされます。通常、これはナイトスタンドで電話が鳴り始める場所です。代わりに、Pulse がそれをいつでも即座に受け取ります。
PagerDuty · トリガーされました
INC-4187 · API ゲートウェイ
オンコール: パルス (AI) → エンジニアは眠ったまま
Pulse はログ、メトリクス、最近のデプロイを読み取り、Claude AI を使用して実際に何が壊れたのか、つまり単なる症状ではなく根本原因を特定します。平均時間: 1 分未満。
脈拍・分析
✓ 12,400 ログ行をスキャン
✓ メトリクスとベースラインを比較
✓ デプロイ #4821 と相関
→ 根本原因が見つかった · 確信度: 高
提案された修正を含む完全な診断が、チームがすでに存在する場所に掲載されます。ワンクリックで解決します。 Pulse が確信を持てない場合は、すべてのコンテキストを添付して人間にエスカレートします。
インシデント #4187 — API ゲートウェイ · /checkout での 5xx レートが 14%
デプロイ #4821 (12 分前) では、支払いバリデーターに null チェック回帰が導入されました。エラーはロールアウト後 90 秒後に始まり、新しいコード パスと 1:1 で相関します。
#4820 にロールバックします。信頼性: 高い。
38 秒で診断 · ログ + メトリクス + デプロイ履歴が添付
1 つのプラン - ダウンタイムが 1 時間よりも安い
平らな。インシデントごとの料金設定や座席数の計算はありません。
エンジニアを午前3時に起こすのはやめましょう
私たちは少人数のチームを新人として受け入れています。メールを残していただければ、空きがあれば招待状をお送りします。
エンジニアリング チーム向けの AI インシデント対応。オンコールなので、担当者が不在である必要はありません。

## Original Extract

Pulse watches your systems 24/7. When an incident fires, AI diagnoses it and posts the answer to Slack — before anyone gets paged.

Pulse
How it works
Features
Pricing
Get access
AI incident response
When an incident fires, Pulse diagnoses it with Claude AI and posts the answer to Slack — with one-click resolve or escalate. Most nights, no human needs to wake up.
watching your website · 24/7/365
last check: 0.4s ago — all clear
last night, while your team slept
one incident, start to finish — no humans woken
PagerDuty triggers at 3:04 AM. Normally this is where a phone starts buzzing on a nightstand. Pulse picks it up instead — instantly, every time.
PagerDuty · triggered
INC-4187 · api-gateway
on-call: pulse (ai) → your engineers stay asleep
Pulse reads your logs, metrics, and recent deploys, then uses Claude AI to work out what actually broke — root cause, not just symptoms. Median time: under a minute.
Pulse · analyzing
✓ scanned 12,400 log lines
✓ compared metrics vs. baseline
✓ correlated with deploy #4821
→ root cause found · confidence: high
A full diagnosis with a suggested fix, posted where your team already lives. One click resolves it. If Pulse is unsure, it escalates to a human — with all the context attached.
Incident #4187 — api-gateway · 5xx rate at 14% on /checkout
Deploy #4821 (12 min ago) introduced a null check regression in the payment validator. Errors began 90s after rollout and correlate 1:1 with the new code path.
Roll back to #4820 . Confidence: high.
diagnosed in 38s · logs + metrics + deploy history attached
one plan · cheaper than one hour of downtime
flat. no per-incident pricing, no seat math.
Stop waking up your engineers at 3am
We're onboarding a small group of teams. Leave your email and we'll send your invite as spots open.
AI incident response for engineering teams. On call so your people don't have to be.

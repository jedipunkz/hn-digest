---
source: "https://velyr.io/"
hn_url: "https://news.ycombinator.com/item?id=48525761"
title: "Show HN: Velyr – an AI agent that finds and fixes conversion leaks on your site"
article_title: "Velyr — Weekly Conversion Fixes Shipped as GitHub Pull Requests"
author: "flo_r"
captured_at: "2026-06-14T10:12:57Z"
capture_tool: "hn-digest"
hn_id: 48525761
score: 2
comments: 0
posted_at: "2026-06-14T10:00:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Velyr – an AI agent that finds and fixes conversion leaks on your site

- HN: [48525761](https://news.ycombinator.com/item?id=48525761)
- Source: [velyr.io](https://velyr.io/)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T10:00:55Z

## Translation

タイトル: HN を表示: Velyr – サイトのコンバージョン リークを検出して修正する AI エージェント
記事のタイトル: Velyr — 毎週の変換修正が GitHub プル リクエストとして提供される
説明: Velyr は、毎週 1 つの変換修正を GitHub プル リクエストとして出荷する AI 成長エージェントです。ユーザーの承認によってゲートされ、裏目に出た場合は自動的にロールバックされます。 Vercel 上の React、Next.js、および Vite サイト用に構築されています。

記事本文:
AI 成長エージェント · React、Next.js、Vite 用
変換修正は毎週配信されます。
Velyr は、サイトの最大のコンバージョン問題を毎週検出し、その修正を GitHub プル リクエストとして書き込む AI 成長エージェントです。 Telegram で 1 回返信するだけで承認されますが、数が減少した場合は 48 時間以内に承認が取り消されます。
毎週月曜日に、エージェントは PostHog 分析 (トラフィック、直帰率、スクロール深度、クリック動作) を読み取り、GitHub リポジトリ内のすべてのページをスキャンします。
これは、ホームページだけでなくファネル全体にわたって最も大きな影響を与えるコンバージョンの問題を 1 つ特定し、Vercel プレビュー リンクを含むプル リクエストを開きます。
問題、その背後にあるデータ、提案された修正を要約した Telegram メッセージを受け取ります。出荷する場合は YES、スキップする場合は NO と返信します。あなたの承認なしには何も統合されません。
マージから 48 時間後、Velyr はサイト全体の直帰率を再チェックします。 15 パーセンテージ ポイント以上急上昇すると、自動的にロールバック プル リクエストがオープンされます。
毎週 1 回の変換修正がレビュー可能な GitHub プル リクエストとして出荷されます。これは実際のコードであり、自分で実装する必要がある提案のダッシュボードではありません。
リポジトリ内のすべてのページをマッピングし、分析を相互参照して、訪問者が実際にどこから離脱したかを見つけるフルファネル分析。
ワンタップの Telegram 承認フロー: すべての変更は明示的な YES の背後にゲートされ、最初に正確な差分が表示されます。
変更によって数値に悪影響が出た場合の 48 時間の自動ロールバック セーフティ ネット。単なる約束ではなく、慎重に元に戻します。
競合他社の追跡とブランド ガードレール - エージェントが遵守する必要がある書面によるルール。そのため、ブランド外の提案は連絡される前に拒否されます。
ほとんどの変換ツールはレポートを提供し、実装はユーザーに任せます。 Velyr は実際のコード変更を記述し、既存の GitHub、Verce 内でプル リクエストとして開きます。

l、PostHog ワークフロー — 移行先の新しいプラットフォームはありません。すべての変更はマージ前に承認され、バックファイア修正は自動的に元に戻ります。これは継続的なコンバージョンの最適化であり、人間が関与し、その下にセーフティ ネットが組み込まれています。
GitHub リポジトリを使用して Vercel にデプロイされた React、Next.js、または Vite サイト - CRO 代理店を雇用したり手動で実験を実行したりせずに、継続的なコンバージョンの最適化を望むインディーズ ハッカー、個人創設者、および小規模の SaaS チーム。エージェントはソース コードを編集するため、それを公開しないサイト ビルダー (Shopify、Wix、Squarespace、Webflow) はサポートされていません。
すべての機能が含まれる 14 日間の無料トライアル後は月額 29 ユーロ。開始するにはクレジット カードが必要です。試用期間中はいつでも請求なしでキャンセルでき、その後もいつでもキャンセルできます。 GitHub リポジトリを使用して、Vercel 上の React、Next.js、および Vite サイト用に構築されています。
無料トライアルを開始する ·
よくある質問・
プライバシー

## Original Extract

Velyr is an AI growth agent that ships one weekly conversion fix as a GitHub Pull Request — gated by your approval, auto-rolled-back if it backfires. Built for React, Next.js, and Vite sites on Vercel.

AI Growth Agent · for React, Next.js & Vite
Conversion fixes, shipped weekly.
Velyr is an AI growth agent that finds your site's #1 conversion problem each week and writes the fix as a GitHub Pull Request. You approve it with one Telegram reply — and if the numbers drop, it reverts itself within 48 hours.
Every Monday the agent reads your PostHog analytics (traffic, bounce rate, scroll depth, and click behavior) and scans every page in your GitHub repo.
It identifies the single highest-impact conversion problem across your full funnel — not just the homepage — and opens a Pull Request with a Vercel preview link.
You get a Telegram message summarizing the problem, the data behind it, and the proposed fix — reply YES to ship or NO to skip. Nothing merges without your approval.
48 hours after a merge, Velyr re-checks site-wide bounce rate; if it spikes by 15 percentage points or more, it automatically opens a rollback Pull Request.
One weekly conversion fix shipped as a reviewable GitHub Pull Request — real code, not a dashboard of suggestions you have to implement yourself.
Full-funnel analysis that maps every page in your repo and cross-references your analytics to find where visitors actually drop off.
A one-tap Telegram approval flow: every change is gated behind your explicit YES, and you see the exact diff first.
An automatic 48-hour rollback safety net if a change hurts your numbers — a measured revert, not just a promise.
Competitor tracking and Brand Guardrails — written rules the agent must respect, so off-brand suggestions are rejected before they reach you.
Most conversion tools hand you a report and leave the implementation to you. Velyr writes the actual code change and opens it as a Pull Request inside your existing GitHub, Vercel, and PostHog workflow — there is no new platform to migrate to. You approve every change before it merges, and a backfiring fix reverts itself automatically. It is continuous conversion optimization that ships, with a human in the loop and a safety net underneath.
React, Next.js, or Vite sites deployed on Vercel with a GitHub repository — indie hackers, solo founders, and small SaaS teams who want ongoing conversion optimization without hiring a CRO agency or running experiments by hand. Because the agent edits source code, site builders that don't expose it (Shopify, Wix, Squarespace, Webflow) are not supported.
€29/month after a 14-day free trial with every feature included. A credit card is required to start, you can cancel anytime during the trial without being charged, and you can cancel anytime afterwards too. Built for React, Next.js, and Vite sites on Vercel with a GitHub repo.
Start free trial ·
FAQ ·
Privacy

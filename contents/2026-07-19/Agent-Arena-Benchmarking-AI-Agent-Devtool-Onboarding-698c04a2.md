---
source: "https://2027.dev/arena/sandboxes"
hn_url: "https://news.ycombinator.com/item?id=48966766"
title: "Agent Arena: Benchmarking AI Agent Devtool Onboarding"
article_title: "Agent Arena — 2027.dev"
author: "karlmush"
captured_at: "2026-07-19T11:04:35Z"
capture_tool: "hn-digest"
hn_id: 48966766
score: 2
comments: 0
posted_at: "2026-07-19T10:44:13Z"
tags:
  - hacker-news
  - translated
---

# Agent Arena: Benchmarking AI Agent Devtool Onboarding

- HN: [48966766](https://news.ycombinator.com/item?id=48966766)
- Source: [2027.dev](https://2027.dev/arena/sandboxes)
- Score: 2
- Comments: 0
- Posted: 2026-07-19T10:44:13Z

## Translation

タイトル: Agent Arena: AI エージェント開発ツールのオンボーディングのベンチマーク
記事のタイトル: エージェント アリーナ — 2027.dev
説明: AI エージェントがドキュメントのみから開発ツールをどれだけ速く統合できるかをテストします。

記事本文:
エージェント アリーナ — 2027.dev エージェント アリーナ
© 2027.dev — カリフォルニア州サンフランシスコ · 概要 · 利用規約 · プライバシー
AI エージェントが完全に自律的に devtools を使い始めるのがいかに簡単かを評価します。
ブラウザ オートメーション サンドボックス 検索 AI フレームワーク データベース & BaaS 認証 & インフラ 電子メール & メッセージング LLM 可観測性 画像生成 ビデオ生成 ハーネス × モデル クロード コード · Opus 4.6 クロード コード · Sonnet 4.6 近日登場 クロード コード · Haiku 4.5 近日登場 クロード コード · Fable 5 近日登場 Codex · GPT-5.5 近日登場 OpenCode · GPT-5.5 近日登場 平均非表示チャート # ツール時間コストエラー 中断 1 デイトナ 2 分 21 秒 $0.81 0.8 1 2 E2B 1 分 18 秒 $0.84 1.6 2 3 モーダル 2 分 46 秒 $1.16 1.3 1 4 フリースタイル 3 分 39 秒 $1.20 1 1 5 ヴァーセル サンドボックス 3 分 18 秒 $2.20 2 1 6 Blaxel 8 分 44 秒 $3.50 1 1 7 Cloudflare Sandbox 13 分 11 秒 $4.73 2.7 0 8 CodeSandbox は自動サインアップをブロックします – – – – 方法論 AI コーディング エージェントは、タスク プロンプトと URL を備えた分離された Docker コンテナ内で実行されます。各エージェントは自律的にドキュメントを検出し、パッケージをインストールし、機能するコードを作成し、結果を検証する必要があります。これらはすべて、求められたときに API 資格情報を提供する以外に人間の助けを必要としません。
ランキングの仕組み: 同じカテゴリのプロバイダーは、時間、コスト、エラー、中断の 4 つの側面で相互にランク付けされます。カテゴリ内では、ディメンションごとのランキングが総合的な順位に結合されます。時間の短縮、コストの削減、エラーの減少、中断の減少はすべて、プロバイダーの地位を向上させます。総合スコアやレターグレードはありません。
すべてのツール呼び出し、エラー、タイミング、トークンの使用状況が記録されます。ランキングはセッション ログから決定されます。プロバイダーごとの複数の独立した実行が集約されます。

## Original Extract

We test how fast an AI agent can integrate your dev tool — from docs alone.

Agent Arena — 2027.dev Agent Arena
© 2027.dev — San Francisco, CA · About · Terms · Privacy
We evaluate how easy it is for AI agents to get started with devtools, fully autonomously.
Browser Automation Sandboxes Search AI Frameworks Databases & BaaS Auth & Infra Email & Messaging LLM Observability Image Generation Video Generation Harness × Model Claude Code · Opus 4.6 Claude Code · Sonnet 4.6 coming soon Claude Code · Haiku 4.5 coming soon Claude Code · Fable 5 coming soon Codex · GPT-5.5 coming soon OpenCode · GPT-5.5 coming soon Average Hide chart # Tool Time Cost Errors Interruptions 1 Daytona 2m 21s $0.81 0.8 1 2 E2B 1m 18s $0.84 1.6 2 3 Modal 2m 46s $1.16 1.3 1 4 Freestyle 3m 39s $1.20 1 1 5 Vercel Sandbox 3m 18s $2.20 2 1 6 Blaxel 8m 44s $3.50 1 1 7 Cloudflare Sandbox 13m 11s $4.73 2.7 0 8 CodeSandbox blocks automated signup – – – – Methodology AI coding agents run inside isolated Docker containers with a task prompt and a URL. Each agent must autonomously discover docs, install packages, write working code, and verify the result — all without human help beyond providing API credentials when asked.
How rankings work: providers in the same category are ranked against each other on four dimensions — Time, Cost, Errors, and Interruptions. Within a category, the per-dimension rankings combine into an overall position. Less time, lower cost, fewer errors, and fewer interruptions all improve a provider's standing. No composite score, no letter grades.
All tool calls, errors, timing, and token usage are recorded. Rankings are deterministic from session logs. Multiple independent runs per provider are aggregated.

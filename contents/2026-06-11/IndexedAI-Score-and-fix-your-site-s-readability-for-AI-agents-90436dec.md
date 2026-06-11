---
source: "https://www.indexedai.tech"
hn_url: "https://news.ycombinator.com/item?id=48496439"
title: "IndexedAI – Score and fix your site's readability for AI agents"
article_title: "IndexedAI — Make your website readable by AI agents"
author: "guidodr"
captured_at: "2026-06-11T21:44:58Z"
capture_tool: "hn-digest"
hn_id: 48496439
score: 1
comments: 0
posted_at: "2026-06-11T21:08:39Z"
tags:
  - hacker-news
  - translated
---

# IndexedAI – Score and fix your site's readability for AI agents

- HN: [48496439](https://news.ycombinator.com/item?id=48496439)
- Source: [www.indexedai.tech](https://www.indexedai.tech)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T21:08:39Z

## Translation

タイトル: IndexedAI – AI エージェントに対するサイトの読みやすさをスコアリングして修正する
記事のタイトル: IndexedAI — Web サイトを AI エージェントが読み取れるようにする
説明: llms.txt を生成し、60 秒以内にエージェント準備スコアを取得します。無料、アカウントは必要ありません。

記事本文:
IndexedAI — AI エージェントが Web サイトを読み取れるようにします。 料金 llmstxt.org ↗ 無料ベータ版 · アカウントは必要ありません。 Web サイトを作成します。
AIエージェントによる読み取り可能
URLを入力します。 Agent Readiness Score、カスタム llms.txt、ライブ MCP エンドポイントを 90 秒以内に取得します。
結果はメールで送信されます · 30 ～ 90 秒かかります
IndexedAI は最大 6 ページをスクレイピングし、 robots.txt と sitemap.xml を読み取り、API エンドポイントとドキュメントのリンクを検出します。
エージェント準備スコアは 5 つの軸にわたって計算されます。次に、Claude は、調整された llms.txt および llms-full.txt を生成します。
あなたのサイトにはライブ MCP サーバーが導入されます。 Claude Desktop、Cursor、または任意の AI クライアントを単一の JSON 構成で接続します。追加のセットアップは必要ありません。
すべてのレポートは、AI エージェント、LLM パイプライン、検索アシスタントにとって重要な準備状況を 5 つの軸で分析しています。
robots.txt が見つかりました sitemap.xml が見つかりました 4 ページがクロールされました llms.txt はまだありません 解析性 18 / 20 タイトルが存在します メタ ディスクリプション > 50 文字 H1 構造がクリア ホームページ > 500 ワード トークン効率 12 / 20 コンテンツ < 50k ワード ナビゲーション < 20 リンク 重複コンテンツの可能性 機能 シグナリング 10 / 20 API エンドポイントが検出されました 価格ページが見つかりました いいえ/docs が検出されました アクセス制御 14 / 20 すべてのページが公開されています robots.txt に AI ボット ルールがありません MCP エンドポイント
あなたのサイトは、あらゆる AI エージェントを受け入れる準備ができています。
Claude Desktop、Cursor、GPT、または任意の MCP クライアントに数秒で接続します。カスタム統合や API キーの管理は必要ありません。
このサイトの構造化された llms.txt インデックスを取得します。最初にこれを呼び出します — サイトの目的、セクション、重要なページを示します。
このサイトの任意のページの全文コンテンツを取得します。 HTML を含まず、読みやすいクリーンなテキストを返します。
拡張コンテキスト: すべての主要セクション、API エンドポイント、価格設定、機能、および FAQ。
Pro プランでは、クロール データから自動生成されるサイト固有の動的ツールも利用できます。
Cl に貼り付けます

オードデスクトップ構成
{
"mcpサーバー": {
"あなたのドメイン.com": {
"url": "https://indexedai.tech/api/mcp/{id}",
"トランスポート": "ストリーミング可能な http"
}
}
Claude Desktop、Cursor、Continue、Cline などの MCP 互換クライアントで動作します。
どの AI エージェントがサイトを訪問するかを把握します。
すべての MCP ツール呼び出しがログに記録されます。 AI トラフィックの長期的な傾向と、どのモデルがコンテンツを読み取っているかを確認します。
Pro にアップグレードすると、エージェントの内訳が表示されます
Claude Code から直接 IndexedAI サイトをクエリする
/indexedai domain.com あなたの質問 — 生の HTML の代わりに構造化された MCP データを取得します。コンテキスト オーバーヘッドなし: MCP は呼び出されたときにのみロードされます。
これがあなたの受信箱に届きます。
スコアの内訳と、すぐにデプロイできる 2 つのファイル: llms.txt および llms-full.txt 。
エージェント準備状況レポート — Stripe.com
送信者: IndexedAI <results@indexedai.tech>
無料で始めましょう。さらに多くのサイトや分析が必要な場合はアップグレードしてください。
支払いは Stripe によって安全に処理されます。いつでもキャンセルできます。
llms.txt は、Web サイトのルート (/llms.txt) にあるプレーンテキスト ファイルで、サイトの内容、重要なページ、コンテンツを効率的に移動する方法を AI エージェントに伝えます。 robots.txt を考えてください。ただし、LLM 用です。
MCP (Model Context Protocol) は、AI エージェントを外部データ ソースに接続するためのオープン スタンダードです。 IndexedAI はサイトごとに一意の MCP サーバー URL を生成します。それを Claude Desktop、Cursor、または任意の MCP クライアントに貼り付けると、AI エージェントがサイトを即座に読み取り、ナビゲートできるようになります。
なぜ AI エージェントはほとんどの Web サイトを読み取ることができないのでしょうか?
最新の Web サイトは、重い JavaScript、ナビゲーション メニュー、Cookie バナー、マーケティング ノイズなど、ブラウザー向けに構築されています。 AI エージェントは生のコンテンツを解析し、機械読み取り用に設計されていない構造で迷子になります。 llms.txt は、クリーンで構造化されたエントリ ポイントを提供します。
AI 訪問追跡はどのように機能しますか?
AI エージェントが MCP エンドポイントを呼び出すたびに、IndexedAI は

ツール名、タイムスタンプ、ユーザーエージェント。無料ユーザーには合計訪問数が表示されます。プロ ユーザーは、エージェントの内訳をすべて明らかにします。どの AI システム (Claude、GPT、Perplexity) がサイトにアクセスしているのか、どのくらいの頻度でアクセスしているのかを正確に確認できます。
分析にはどれくらい時間がかかりますか?
通常は 30 ～ 90 秒です。最大 6 ページをクロールし、5 つの軸にわたってスコアを付け、AI を使用してカスタマイズされた llms.txt を生成します。結果は完了した瞬間に受信箱に届きます。待つ必要はありません。
はい — 無料プランには、1 つのサイト、MCP エンドポイント、エージェント準備スコア、および llms.txt の無料生成が含まれています。 Pro (月額 9.99 ドル) では、最大 5 つのサイト、分析、優先処理のロックが解除されます。
ベータ期間中は無料。アカウントがありません。クレジットカードはありません。

## Original Extract

Generate llms.txt and get your Agent Readiness Score in 60 seconds. Free, no account required.

IndexedAI — Make your website readable by AI agents Pricing llmstxt.org ↗ Free beta · No account required Make your website
readable by AI agents
Enter a URL. Get your Agent Readiness Score, a custom llms.txt , and a live MCP endpoint — in under 90 seconds.
Results sent to your email · Takes 30–90 seconds
IndexedAI scrapes up to 6 pages, reads robots.txt and sitemap.xml , detects API endpoints and documentation links.
Your Agent Readiness Score is calculated across 5 axes. Claude then generates a tailored llms.txt and llms-full.txt .
Your site gets a live MCP server. Connect Claude Desktop, Cursor, or any AI client with a single JSON config — no extra setup.
Every report breaks down readiness across 5 axes that matter to AI agents, LLM pipelines, and search assistants.
robots.txt found sitemap.xml found 4 pages crawled No llms.txt yet Parsability 18 / 20 Title present Meta description > 50 chars H1 structure clear Homepage > 500 words Token Efficiency 12 / 20 Content < 50k words Navigation < 20 links Possible duplicate content Capability Signaling 10 / 20 API endpoints detected Pricing page found No /docs detected Access Control 14 / 20 All pages publicly accessible No AI bot rules in robots.txt MCP endpoint
Your site, ready for any AI agent.
Connect Claude Desktop, Cursor, GPT, or any MCP client in seconds. No custom integrations, no API keys to manage.
Get the structured llms.txt index of this site. Call this first — shows site purpose, sections, and important pages.
Fetch the full text content of any page on this site. Returns clean readable text, no HTML.
Extended context: all major sections, API endpoints, pricing, capabilities, and FAQ.
Pro plans also get site-specific dynamic tools — auto-generated from your crawl data.
Paste into Claude Desktop config
{
"mcpServers": {
"your-domain.com": {
"url": "https://indexedai.tech/api/mcp/{id}",
"transport": "streamable-http"
}
}
} Works with any MCP-compatible client: Claude Desktop, Cursor, Continue, Cline and more.
Know which AI agents visit your site.
Every MCP tool call is logged. See how AI traffic trends over time and which models are reading your content.
Upgrade to Pro to see agent breakdown
Query any IndexedAI site directly from Claude Code
/indexedai domain.com your question — fetches structured MCP data instead of raw HTML. Zero context overhead: MCP loads only when invoked.
This is what lands in your inbox.
Score breakdown + two ready-to-deploy files: llms.txt and llms-full.txt .
Your Agent Readiness Report — stripe.com
From: IndexedAI <results@indexedai.tech>
Start free. Upgrade when you need more sites or analytics.
Payments processed securely by Stripe. Cancel anytime.
llms.txt is a plain-text file at the root of your website (/llms.txt) that tells AI agents what your site is about, what pages matter, and how to navigate your content efficiently. Think robots.txt — but for LLMs.
MCP (Model Context Protocol) is an open standard for connecting AI agents to external data sources. IndexedAI generates a unique MCP server URL for each site — paste it into Claude Desktop, Cursor, or any MCP client and your AI agent can instantly read and navigate your site.
Why can't AI agents read most websites?
Modern websites are built for browsers: heavy JavaScript, nav menus, cookie banners, marketing noise. AI agents parse raw content and get lost in structure that wasn't designed for machine reading. llms.txt gives them a clean, structured entry point.
How does AI visit tracking work?
Every time an AI agent calls your MCP endpoint, IndexedAI logs the tool name, timestamp, and user-agent. Free users see total visit counts. Pro users unlock the full agent breakdown — see exactly which AI systems (Claude, GPT, Perplexity) are accessing your site and how often.
How long does the analysis take?
Usually 30–90 seconds. We crawl up to 6 pages, score across 5 axes, then generate a tailored llms.txt using AI. Results land in your inbox the moment it's done — no waiting around.
Yes — the Free plan includes 1 site, an MCP endpoint, Agent Readiness Score, and llms.txt generation at no cost. Pro ($9.99/mo) unlocks up to 5 sites, analytics, and priority processing.
Free during beta. No account. No credit card.

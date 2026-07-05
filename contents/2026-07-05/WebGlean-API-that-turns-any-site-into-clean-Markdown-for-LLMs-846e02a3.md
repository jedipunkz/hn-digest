---
source: "https://www.webglean.com"
hn_url: "https://news.ycombinator.com/item?id=48797413"
title: "WebGlean – API that turns any site into clean Markdown for LLMs"
article_title: "WebGlean"
author: "qubomax"
captured_at: "2026-07-05T20:47:38Z"
capture_tool: "hn-digest"
hn_id: 48797413
score: 3
comments: 0
posted_at: "2026-07-05T19:53:26Z"
tags:
  - hacker-news
  - translated
---

# WebGlean – API that turns any site into clean Markdown for LLMs

- HN: [48797413](https://news.ycombinator.com/item?id=48797413)
- Source: [www.webglean.com](https://www.webglean.com)
- Score: 3
- Comments: 0
- Posted: 2026-07-05T19:53:26Z

## Translation

タイトル: WebGlean – あらゆるサイトを LLM のクリーンな Markdown に変える API
記事のタイトル: WebGlean
説明: あらゆる URL を AI 用のクリーンな構造化データに変換します。

記事本文:
WebGlean 新しい Extract API が公開されました。Claude を利用したあらゆる URL からの構造化データ抽出です。抽出 API が公開されました。ドキュメントを参照してください。 WebGlean 製品リソース 価格設定 プレイグラウンド サインイン ダッシュボード 6 つの API がすべてライブ - スクレイピング、クロール、抽出、マップ、監視、検索 あらゆる Web サイトを
AI にクリーンなデータを提供します。
WebGlean は、完全な JavaScript レンダリング、AI を活用した抽出、そして非常にシンプルなクレジットベースの価格設定を使用して、あらゆる URL を Markdown、JSON、またはプレーン テキストに変換します。
AI を使用して構築する開発者から信頼される
API キーは 1 つです。ワンクレジット制。使った分だけお支払いください。
あらゆる URL を Markdown、HTML、JSON、またはプレーン テキストに変換します。 Playwright による完全な JS レンダリング。
BFS はサイト全体をクロールします。最大深度とページ制限を設定します。すべてのページをマークダウンとして返します。
URL をスクレイピングし、Claude を使用して、定義したスキーマに一致する構造化 JSON をプルします。
サイト上の検出可能な URL をすべて検出します。最初に sitemap.xml を試行し、BFS にフォールバックします。
変更検知用のURLを登録します。コンテンツが変更されたときに Webhook 通知を受け取ります。
Web を検索し、結果ごとにスクレイピングされた Markdown を取得して、LLM の準備を整えます。
初日から開発者のエクスペリエンスと AI パイプラインのニーズに基づいて設計されています。
コールド Chromium は起動しません。 Playwright は常時稼働のワーカーで実行されるため、リクエストは 5 ～ 8 秒ではなく 1 ～ 3 秒で返されます。すべてのスクレイピングには、認証されたウォームなブラウザが使用されます。
Extract API は、Claude を使用して、XPath、CSS セレクター、カスタム パーサーを使用せずに、任意の URL から構造化スキーマを取得します。必要なものを JSON で記述して返します。
すべてのエンドポイントにわたって、ページごとに 1 クレジット。クレジットは期限切れになることも、リセットされることもありません。アップグレードせずに補充します。エンドポイントごとの価格設定やシート料金など、サプライズはありません。
Mozilla Readability は、ナビゲーション、広告、フッター、Cookie バナーを削除します。 LLM はノイズではなく信号を受け取ります。
マークダウン、HTML、プレーンテキスト、JSONメタデータ、

およびスクリーンショット — 1 回の呼び出しで、任意の組み合わせが可能です。
1 つのキーで 6 つのエンドポイントすべてをカバーします。 SHA-256 ハッシュ化され、作成時に 1 回表示され、即座に取り消し可能です。
デフォルトではキーごとに 60 要求/分。追加の設定はありません。 Pro と Scale の制限が高くなります。
Monitor は URL を登録し、コンテンツが変更されたときに差分をエンドポイントに POST します。
マップは最初に sitemap.xml を試行し、再帰的リンク抽出に戻ります。検出可能なすべての URL を返します。
今すぐお試しください - アカウントは必要ありません
任意の URL を貼り付けると、数秒できれいな出力が表示されます。
クレジットに有効期限はありません。アップグレードしなくても、いつでも補充できます。
さらに必要ですか?トップアップ パック: $10 = 1,000 クレジット — どのプランでも利用可能で、アップグレードは必要ありません。
始める前に知っておくべきことすべて。
1 クレジットとしてカウントされるものは何ですか?未使用のクレジットは繰り越されますか?抽出 API はどのように機能しますか?レート制限はありますか?アカウントを作成せずに WebGlean を使用できますか? Scrape API はどのような出力形式をサポートしていますか? Python または Node SDK はありますか? WebGlean あらゆる URL を AI パイプラインと開発者のワークフロー用のクリーンな構造化データに変換する最も簡単な方法。
© 2025 WebGlean.無断転載を禁じます。

## Original Extract

Turn any URL into clean, structured data for AI

WebGlean New Extract API is live — Claude-powered structured data extraction from any URL. Extract API is now live. See the docs WebGlean Products Resources Pricing Playground Sign in Dashboard All 6 APIs live — scrape, crawl, extract, map, monitor, search Turn any website into
clean data for your AI.
WebGlean converts any URL to Markdown, JSON, or plain text — with full JavaScript rendering, AI-powered extraction, and dead-simple credit-based pricing.
Trusted by developers building with AI
One API key. One credit system. Pay only for what you use.
Convert any URL to Markdown, HTML, JSON, or plain text. Full JS rendering via Playwright.
BFS-crawl an entire site. Set max depth and page limits. Returns all pages as Markdown.
Scrape a URL then use Claude to pull structured JSON matching any schema you define.
Discover all discoverable URLs on a site. Tries sitemap.xml first, falls back to BFS.
Register a URL for change detection. Get webhook notifications when content changes.
Search the web and get scraped Markdown for each result — ready for your LLM.
Designed around developer experience and AI pipeline needs from day one.
No cold Chromium starts. Playwright runs in an always-on worker so your requests return in 1–3 seconds, not 5–8. Every scrape gets a warm, authenticated browser.
The Extract API uses Claude to pull any structured schema from any URL — no XPath, no CSS selectors, no custom parsers. Describe what you want in JSON and get it back.
One credit per page, across all endpoints. Credits never expire and never reset. Top up without upgrading. No per-endpoint pricing, no seat fees, no surprises.
Mozilla Readability strips nav, ads, footers, and cookie banners. Your LLM gets signal, not noise.
Markdown, HTML, plain text, JSON metadata, and screenshots — one call, any combination.
One key covers all 6 endpoints. SHA-256 hashed, shown once on creation, revokable instantly.
60 req/min per key by default. No extra config. Higher limits on Pro and Scale.
Monitor registers a URL and POSTs a diff to your endpoint when content changes.
Map tries sitemap.xml first — falls back to recursive link extraction. Returns every discoverable URL.
Try it now — no account needed
Paste any URL and see clean output in seconds.
Credits never expire. Top up anytime without upgrading.
Need more? Top-up packs: $10 = 1,000 credits — available on any plan, no upgrade required.
Everything you need to know before you start.
What counts as one credit? Do unused credits roll over? How does the Extract API work? Is there a rate limit? Can I use WebGlean without creating an account? What output formats does the Scrape API support? Do you have Python or Node SDKs? WebGlean The simplest way to convert any URL into clean, structured data for AI pipelines and developer workflows.
© 2025 WebGlean. All rights reserved.

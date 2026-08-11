---
source: "https://webstractor.com/"
hn_url: "https://news.ycombinator.com/item?id=49263068"
title: "Show HN: Webstractor – Pay-as-You-Go Web Data API for AI Agents"
article_title: "Webstractor — Power AI agents with clean web data"
author: "mariusbolik"
captured_at: "2026-08-11T19:55:35Z"
capture_tool: "hn-digest"
hn_id: 49263068
score: 3
comments: 1
posted_at: "2026-08-11T19:15:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Webstractor – Pay-as-You-Go Web Data API for AI Agents

- HN: [49263068](https://news.ycombinator.com/item?id=49263068)
- Source: [webstractor.com](https://webstractor.com/)
- Score: 3
- Comments: 1
- Posted: 2026-08-11T19:15:24Z

## Translation

タイトル: Show HN: Webstractor – AI エージェント用の従量課金制 Web データ API
記事のタイトル: Webstractor — クリーンな Web データを備えた強力な AI エージェント
説明: シンプルなキャッシュファースト GET API と、パブリック Web で AI エージェントを検索および抽出するホスト型 MCP サーバー。

記事本文:
Webstractor — クリーンな Web データを備えた Power AI エージェント webstractor ドキュメント 料金 お問い合わせ ダッシュボード ↗ エージェント対応
エージェント対応のコンテキストに抽出された Web。
パブリック Web を検索し、クリーンな Markdown または構造化 JSON として抽出します。 1 つのキャッシュファースト GET API と、エージェント用に構築されたホスト型 MCP サーバー。
[ アカウントなし · 毎日 10 件の無料リクエスト · キャッシュ ヒットは無料 ]
コピー 生のままで開く ↗ キャッシュされたサンプルをロードしています… 出力形式 MD JSON ソース全体を検索 パブリック Web を検索します。
一貫した 1 つのインターフェイスを通じて、Web ページ、最新ニュース、公開ライセンスされた画像、公開ビデオ、場所を検索します。
Web ニュース 画像 ビデオ 場所 検索クエリ ライブ検索の実行 ドキュメントの表示 パブリック URL の抽出 パブリック URL をコンテキストに変換します。
ユーザーが開く通常のページを送信します。読み取り可能な Markdown または正規化された schema-v1 JSON を受信します。
ライブ抽出の実行 ドキュメントの表示 例の読み込み… コピー 生で開く ↗ キャッシュされた例の読み込み… 出力形式 MD JSON 例の読み込み… URL のコピー スクリーンショットを開く ↗ キャッシュされた例の読み込み… スクリーンショット形式 WebP PNG Web サイトのスクリーンショット 公開 Web ページをキャプチャします。
一貫した WebP または PNG プレビューをレンダリングします。同一のキャプチャは、別のクレジットを使用せずにキャッシュから直接取得されます。
デスクトップ タブレット モバイル 公開 Web サイト URL スクリーンショットのキャプチャ ドキュメントの表示 GET 専用パブリック API 30 日間の抽出キャッシュ スキーマ v1 正規化出力 専用アダプター
パブリック Web 用の 1 つのクリーンなインターフェイス。
これらのサポートされているプラットフォームで動作します
1 回のリクエストで Web データをクリーンアップします。
Web 、ニュース、画像、ビデオ、場所、市場データを検索し、Web サイトのスクリーンショットをキャプチャしてから、/v1/extract を使用してクリーンなページと製品コンテンツを取得します。
カール --get \
'https://api.webstractor.com/v1/extract' \
--data-urlencode 'url=https://www.amazon.com/dp/B09B8V1LZ3' \
--data-urlencode 'format=json' ホスト

MCPサーバー編
エージェントに検索と抽出を提供します。
1 台の MCP サーバーを接続します。エージェントは、Web ページ、ニュース、画像、ビデオ、場所、株式、市場動向を検出し、 extract_public_url を使用してクリーンなコンテンツを読み取ることができます。
https://webstractor.com/mcp MCP ドキュメントを読む → シンプルな価格設定
抽出した分だけお支払いください。
プリペイド クレジットにはハード使用量の上限があります。毎月のプランを必要とせずに、必要なときに資金を追加できます。
抽出、検索、財務、Web サイトのスクリーンショットのリクエストはそれぞれ 1 クレジットとしてカウントされます。
マークダウン、正規化された JSON、またはスクリーンショット
サポートされているすべてのソースとエンドポイント
詳細が必要ですか?価格とよくある質問をご覧ください。
webstractor エージェント、アプリケーション、および 1 行スクリプトのパブリック Web データをクリーンアップします。

## Original Extract

A simple, cache-first GET API and hosted MCP server that searches and extracts the public web for AI agents.

Webstractor — Power AI agents with clean web data webstractor Docs Pricing Contact Dashboard ↗ Agent ready
The web, distilled into agent-ready context.
Search and extract the public web as clean Markdown or structured JSON. One cache-first GET API, plus a hosted MCP server built for agents.
[ No account · 10 free requests daily · Cache hits are free ]
Copy Open raw ↗ Loading the cached example… Output format MD JSON Search across sources Search the public web.
Find webpages, current news, openly licensed images, public videos, and places through one consistent interface.
Web News Images Videos Places Search query Run live search View docs Public URL extraction Turn any public URL into context.
Submit the normal page a person would open. Receive readable Markdown or normalized schema-v1 JSON.
Run live extraction View docs Loading example… Copy Open raw ↗ Loading the cached example… Output format MD JSON Loading example… Copy URL Open screenshot ↗ Loading the cached example… Screenshot format WebP PNG Website screenshots Capture any public webpage.
Render a consistent WebP or PNG preview. Identical captures come straight from the cache without using another credit.
Desktop Tablet Mobile Public website URL Capture screenshot View docs GET-only public API 30-day extraction cache Schema v1 normalized output Purpose-built adapters
One clean interface for the public web.
Works with these supported platforms
Clean web data in one request.
Search the web , news , images , videos , places , and market data , capture website screenshots , then use /v1/extract for clean page and product content.
curl --get \
'https://api.webstractor.com/v1/extract' \
--data-urlencode 'url=https://www.amazon.com/dp/B09B8V1LZ3' \
--data-urlencode 'format=json' Hosted MCP server
Give agents search and extraction.
Connect one MCP server. Agents can discover webpages, news, images, videos, places, stocks, and market movers, then read clean content with extract_public_url .
https://webstractor.com/mcp Read the MCP docs → Simple pricing
Pay only for what you extract.
Prepaid credits are the hard usage cap. Add funds when you need them, without a required monthly plan.
Each extraction, search, finance, or website screenshot request counts as one credit.
Markdown, normalized JSON, or screenshots
Every supported source and endpoint
Need the full details? View pricing and FAQs .
webstractor Clean public web data for agents, applications, and one-line scripts.

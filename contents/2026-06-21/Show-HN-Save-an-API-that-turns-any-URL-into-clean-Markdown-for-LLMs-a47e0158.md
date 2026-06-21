---
source: "https://www.savemarkdown.co/api/"
hn_url: "https://news.ycombinator.com/item?id=48613567"
title: "Show HN: Save, an API that turns any URL into clean Markdown for LLMs"
article_title: "Save API — Any URL to clean Markdown for AI agents"
author: "jswallez"
captured_at: "2026-06-21T01:03:35Z"
capture_tool: "hn-digest"
hn_id: 48613567
score: 3
comments: 0
posted_at: "2026-06-20T22:24:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Save, an API that turns any URL into clean Markdown for LLMs

- HN: [48613567](https://news.ycombinator.com/item?id=48613567)
- Source: [www.savemarkdown.co](https://www.savemarkdown.co/api/)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T22:24:38Z

## Translation

タイトル: Show HN: Save、あらゆる URL を LLM のクリーンなマークダウンに変換する API
記事のタイトル: API を保存 — AI エージェントの Markdown をクリーンアップするための任意の URL
説明: 1 つの API 呼び出しで、あらゆる Web ページが LLM 対応の Markdown に変わります。 AI エージェント、RAG パイプライン、スクレイパー用に構築されています。 JavaScript を取得してレンダリングし、不要なものを取り除きます。開始は無料、プランは月額 29 ドルから。

記事本文:
API を保存 — AI エージェントの Markdown を自動的にクリーンアップする任意の URL
各ページごとに。 fonts.googleapis.com へのクロスオリジン接続はありません。 --> 保存 English English Français Español Deutsch 日本語 Italiano 한국어 Português (ブラジル) 繁體中文 Русский Polski Tiếng Việt API 🔌 Vault を保存 💻 インストール Chrome に追加 🐿️ ## 開発者 API
Markdown をクリーンアップするための任意の URL。
1 回の API 呼び出しで、あらゆる Web ページが LLM 対応の Markdown に変わります。 AI エージェント、RAG パイプライン、スクレイパー向けに構築されており、JavaScript をフェッチしてレンダリングし、不要なものを取り除きます。ノイズではなくコンテンツが得られます。
1,000 ページ/月無料、カードなし、プランは月額 29 ドルから
毎月 1,000 ページ、無料。クレジットカードはありません。キーは即座に機能します。
API キー (一度表示されます - 今すぐコピーします):
次へ: クイックスタートで試してみる →
より高い制限または有料プランが必要ですか?私たちに話してください
必要な分だけを費やす階層型エンジン
実際のブラウザーのフィンガープリントと厳格な SSRF ガードを使用して、サーバー側で URL を取得します。静的ページとサーバーレンダリングページはここで停止します。
ページが JavaScript シェルの場合、コンテンツが薄すぎる場合にのみヘッドレス レンダリングにエスカレートするため、過剰な料金が発生することはありません。
ナビゲーション、広告、Cookie バナー、定型文は削除されます。見出し、リスト、リンク、画像、コードをクリーンな Markdown として取得できます。
カール -X POST https://api.savemarkdown.co/v1/convert \
-H "認可: ベアラー sk_live_..." \
-H "コンテンツ タイプ: application/json" \
-d '{"url": "https://example.com/article"}' const res = await fetch("https://api.savemarkdown.co/v1/convert", {
メソッド: "POST"、
ヘッダー: {
"認可": `ベアラー ${process.env.SAVE_API_KEY}`,
"コンテンツタイプ": "アプリケーション/json",
}、
本文: JSON.stringify({ url: "https://example.com/article" }),
});
const { マークダウン、メタ } = await res.json();
console.log(マークダウン); // クリーンな Markdown インポート OS、リクエスト
解像度 =

リクエスト.post(
"https://api.savemarkdown.co/v1/convert",
headers={"認可": f"ベアラー {os.environ['SAVE_API_KEY']}"},
json={"url": "https://example.com/article"},
)
print(res.json()["markdown"]) 応答
{
"markdown": "📄 タイトル: 記事の例\n🔗 ソース: ...\n\n---\n\n## 見出し\n\n本文をきれいに...",
"メタ": { "層": 1、 "キャッシュ": false、 "ドメイン": "example.com"、 "テンプレート": "マークダウン" },
"使用法": { "クレジット": 1 }
## 料金設定 お客様とともに成長するシンプルなプラン
カードなしで無料で始められます。さらに必要な場合は上に移動します。すべてのアカウントは無料枠から始まります。
✓ JavaScript の自動レンダリング
✓ 300以上のサイト固有のエクストラクター
ページ = Markdown への URL、JS レンダリングが含まれる · Starter+ の AI 形式の出力 · さらに必要ですか?私たちに話してください
認可: Bearer sk_live_… で認証します。キーはサーバー側のみにあります。ブラウザーやアプリ バンドルにキーを同梱することはありません。
変換するページ。 http のみ。
生のマークダウンの場合は省略します。 LLM フォーマット (トークン従量制) 用の AI テンプレート (クリーン、サマリー、アウトラインなど) を渡します。
JavaScript のレンダリング。デフォルトの「自動」は、ページが JS シェルの場合にのみ Tier 2 にエスカレーションします。
24 時間キャッシュをバイパスして再フェッチします。
ウォールド ガーデン: YouTube は公式トランスクリプト API を通じてルーティングします。 X / Instagram / TikTok はベストエフォートです。私たちはゴーストアカウントを使用したり、ログインの背後にあるスクレイピングを使用したりすることはありません。
エージェントが読み取り可能なウェブ
LLM は HTML ではなくマークダウンを読み取ります。エージェント、RAG インジェスト ジョブ、またはリサーチ ループを任意の URL に指定し、保持する価値のあるトークンを取得します。この仕様は、プレーンなマークダウンとしても入手できます。
/api.md 。 API は、
llms.txt 、
エージェント スキル インデックス、および
API カタログ。
無料のキーを取得して、最初のページを 1 分以内に変換します。
無料で始められます · プランは月額 29 ドルからです · いつでもキャンセルできます
Claude & ChatGPT の個人的な知識ベース。

## Original Extract

One API call turns any web page into LLM-ready Markdown. Built for AI agents, RAG pipelines and scrapers. Fetch, render JavaScript, strip the clutter. Free to start, plans from $29/mo.

Save API — Any URL to clean Markdown for AI agents automatically
for each page; no cross-origin connect to fonts.googleapis.com. --> Save English English Français Español Deutsch 日本語 Italiano 한국어 Português (Brasil) 繁體中文 Русский Polski Tiếng Việt API 🔌 Save Vault 💻 Install Add to Chrome 🐿️ ## developer api
Any URL to clean Markdown.
One API call turns any web page into LLM-ready Markdown. Built for AI agents, RAG pipelines, and scrapers — fetch, render JavaScript, strip the clutter. You get the content, not the noise.
1,000 pages/mo free · no card · plans from $29/mo
1,000 pages a month, free. No credit card. Your key works instantly.
Your API key (shown once — copy it now):
Next: try it in the quickstart →
need higher limits or a paid plan? talk to us
A tiered engine that only spends what it must
We fetch the URL server-side with a real browser fingerprint and a strict SSRF guard. Static and server-rendered pages stop here.
If a page is a JavaScript shell, we escalate to a headless render — only when the content is too thin, so you never overpay.
Nav, ads, cookie banners and boilerplate are stripped. You get headings, lists, links, images and code as clean Markdown.
curl -X POST https://api.savemarkdown.co/v1/convert \
-H "Authorization: Bearer sk_live_..." \
-H "Content-Type: application/json" \
-d '{"url": "https://example.com/article"}' const res = await fetch("https://api.savemarkdown.co/v1/convert", {
method: "POST",
headers: {
"Authorization": `Bearer ${process.env.SAVE_API_KEY}`,
"Content-Type": "application/json",
},
body: JSON.stringify({ url: "https://example.com/article" }),
});
const { markdown, meta } = await res.json();
console.log(markdown); // clean Markdown import os, requests
res = requests.post(
"https://api.savemarkdown.co/v1/convert",
headers={"Authorization": f"Bearer {os.environ['SAVE_API_KEY']}"},
json={"url": "https://example.com/article"},
)
print(res.json()["markdown"]) Response
{
"markdown": "📄 Title: Example Article\n🔗 Source: ...\n\n---\n\n## Heading\n\nClean body...",
"meta": { "tier": 1, "cached": false, "domain": "example.com", "template": "markdown" },
"usage": { "credits": 1 }
} ## pricing Simple plans that grow with you
Start free, no card. Move up when you need more. Every account begins on the free tier.
✓ Automatic JavaScript rendering
✓ 300+ site-specific extractors
pages = URL to Markdown, JS rendering included · AI-formatted output on Starter+ · need more? talk to us
Authenticate with Authorization: Bearer sk_live_… . Keys are server-side only — never ship one in a browser or app bundle.
The page to convert. http(s) only.
Omit for raw Markdown. Pass an AI template (clean, summary, outline…) for LLM formatting (token-metered).
JavaScript rendering. Default "auto" escalates to Tier 2 only when a page is a JS shell.
Bypass the 24h cache and refetch.
Walled gardens: YouTube routes through its official transcript API. X / Instagram / TikTok are best-effort. We never use ghost accounts or scrape behind logins.
The web, readable by your agent
LLMs read Markdown, not HTML. Point an agent, a RAG ingest job, or a research loop at any URL and get back tokens worth keeping. This spec is also available as plain Markdown at
/api.md . The API is discoverable through our
llms.txt ,
agent-skills index , and
API catalog .
Get a free key and convert your first page in under a minute.
free to start · plans from $29/mo · cancel anytime
Your personal knowledge base for Claude & ChatGPT.

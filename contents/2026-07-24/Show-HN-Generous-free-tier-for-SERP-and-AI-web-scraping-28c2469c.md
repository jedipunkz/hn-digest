---
source: "https://cloro.dev/"
hn_url: "https://news.ycombinator.com/item?id=49031260"
title: "Show HN: Generous free tier for SERP and AI web scraping"
article_title: "cloro - the best scraper for SEO and AI SEO"
author: "rbatista19"
captured_at: "2026-07-24T05:11:49Z"
capture_tool: "hn-digest"
hn_id: 49031260
score: 1
comments: 0
posted_at: "2026-07-24T04:34:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Generous free tier for SERP and AI web scraping

- HN: [49031260](https://news.ycombinator.com/item?id=49031260)
- Source: [cloro.dev](https://cloro.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T04:34:04Z

## Translation

タイトル: Show HN: SERP と AI Web スクレイピングのための豊富な無料枠
記事のタイトル: cloro - SEO と AI SEO に最適なスクレイパー
説明: cloro は、SEO および AI SEO モデル (ChatGPT、Gemini、Perplexity、Copilot、AI 概要、Google 検索) をスクレイピングするための API です。マークダウン、ソース、引用。
HN テキスト: ハッカーの皆様、世界をリードする AI UI スクレイピング プラットフォームである cloro.dev に新しく追加された定期無料枠を共有したいと思いました。 ChatGPT、Perplexity、Grok、Gemini、Google Search、Google News、Copilot、AI Overview から構造化データを 1 つの API を通じて抽出し、どの国でも低レイテンシで実現します。当社は最近 200 社以上の顧客と取引を行っており、そのほとんどが GEO と SEO を支援する大企業です。今後、すべてのユーザーは毎月 500 の無料クレジットを取得できるため、趣味のプロジェクトに最適です。楽しむ！

記事本文:
cloro - SEO と AI SEO に最適なスクレイパー
メニューの切り替え
ドキュメント
ChatGPT Perplexity Copilot Gemini AI の概要 AI モード検索 AI 時代の検索 API
すべての AI 検索エンジンと Google に 1 つの API。リアルタイムの構造化された JSON。
4.8 · レビュー 42 件 500 クレジットを無料でお試しください SEO および AI SEO のトップ企業が使用 99.99% の稼働率 月間 1,000,000,000 回の API 呼び出し 単一 API による構造化データ
主要な SEO および AI SEO プロバイダーをカバーします。
あらゆる規模で、地域全体で。
ChatGPT ライブ Perplexity ライブ コパイロット ライブ AI モード ライブ AI 概要 ライブ検索 ライブ ニュース ライブ Gemini ライブ Grok 利用不可
下のカード（アマゾン）
アマゾン
トップカード（メタ）
メタ
さらにバッジも追加
さらに詳しく
--> チームがクロロを選ぶ理由
LLM API の直接統合には隠れたコストと困難なスケーラビリティが伴います。
API 応答は UI とはまったく異なります。サーファーはこれをデータで証明しました。
直接プロバイダー API は、SEO に影響を与えるツールを提供しません。
各プロバイダーには個別の API 統合が必要であり、技術的なオーバーヘッドが生じます。
トークンベースの価格はモデルやプロバイダーによって異なるため、コストが予測できません。
cloro 統合の仕組み
各応答から抽出された markdown 、 text 、または HTML に加えて、解析されたソース 、引用 、クエリ ファンアウト 、およびショッピング ブロックを取得します。
Cloro から Cloro をインポート
client = Cloro(api_key="sk_live_your_api_key_here")
応答 = client.monitor.chatgpt(
プロンプト="テスラの最新アップデートについて何を知っていますか?",
国 = "米国"、
include={"マークダウン": True},
)
print(response["result"]["text"]) typescript Copy import { Cloro } from "@cloro-dev/cloro";
const client = new Cloro({ apiKey: "sk_live_your_api_key_here" });
const response = await client.monitor.chatgpt({
プロンプト: "テスラの最新アップデートについて何を知っていますか?",
国: "米国"、
include: { マークダウン: true }、
});
console.log(response.result.text);カールコピー

カール -X POST https://api.cloro.dev/v1/monitor/chatgpt \
-H "認可: ベアラー sk_live_your_api_key_here" \
-H "コンテンツ タイプ: application/json" \
-d '{
"prompt": "テスラの最新アップデートについて何を知っていますか?",
"国": "米国",
「含める」: {
「マークダウン」: true
}
}' 応答例
{
「成功」: true、
「結果」: {
"text": "テスラの最近のアップデートには、完全自動運転機能の大幅な改善が含まれています...",
「ソース」: [
{
「位置」: 1、
"url": "https://tesla.com/updates/fsd",
"label": "テスラ FSD アップデート",
"description": "最新の完全自動運転の改善と機能"
}
]、
"html": "<div class=\"markdown\"><p>Tesla の最近のアップデートには...</p></div>",
"markdown": "**Tesla の最近のアップデート** には大幅な改善が含まれています...",
"検索クエリ": [
「テスラのアップデート 2024」、
「完全自動運転の改善」
]、
"ショッピングカード": [
{
「位置」: 1、
「製品」: {
"名前": "モデル Y",
"ブランド": "テスラ"、
"価格": "$43,990",
"通貨": "USD",
「評価」：4.5、
「レビュー数」: 2847、
"imageUrl": "https://example.com/tesla-model-y.jpg",
"productUrl": "https://tesla.com/modely",
"description": "オートパイロット機能を備えた全電動コンパクトSUV"
}
}
】
}
500 クレジットを無料で試す ドキュメントを表示する ニーズに合わせた価格設定
無料で始めましょう。ボリュームが増えると、クレジットあたりの価格が下がります。以下の各階層を参照してください。
当社の API は、構造化された応答を平均 30 ～ 45 秒で返します。自動化と解析の複雑さをすべて処理して、ほぼリアルタイムでクリーンなマークダウン コンテンツと引用を提供します。
完全な AI 応答を解析済みのマークダウン形式で抽出し、すべてのソースと引用を URL とラベルとともに抽出します。すべては構造化された JSON として提供されるため、システムに簡単に統合できます。
ChatGPT (Web 検索付き)、Google Gemini、Perplexity、Microsoft Copilot、Google AI の概要をサポートしています。

、Google 検索。当社の統合 API を使用すると、一貫したインターフェイスを使用してこれらすべてのソースからデータを抽出できます。
はい！ Webhook 経由でタスクを送信し、結果を受信できる非同期エンドポイントを提供します。これは大量のスクレイピングに最適であり、接続を開いたままにすることなく信頼性を確保します。
同時実行制限は、サブスクリプション プランによって異なります。たとえば、ビジネス プランは多くの場合、より高い同時実行制限 (100 など) で開始されます。制限を超えると、429 ステータス コードが返されます。クライアント側キューを実装するか、非同期エンドポイントを使用することをお勧めします。
より大容量のパッケージをサブスクライブすることで、同時実行性を高めることができます。当社はお客様の組織専用にリソースを提供しているため、トラフィックが急増するだけでなく、それらのリソースが完全に活用されるようにしたいと考えています。ただし、実際に同時実行性によって制限されている場合 (常に同時実行性を 70% 以上使用している場合)、調整できるよう当社にご連絡ください。
はい。デフォルトでは、応答を軽量に保つために、解析された構造化された JSON を返します。ただし、リクエストで include.html: true を設定すると、結果ページの完全な生の HTML を取得できます。
はい、世界中のすべての国をサポートしています。一部のモデル (Gemini など) には地域別の利用制限 (EU など) がある場合があることに注意してください。
通常、複雑なタイムアウト ロジックを実装する必要はありません。当社のシステムは、高い成功率を確保するために、リクエストを最大 5 分間で最大 10 回自動的に再試行します。
クレジットは、AI 応答データの抽出と配信に成功した場合にのみ差し引かれます。抽出の失敗やシステム エラーによってクレジットは消費されません。
いいえ、クレジットは月ごとに繰り越されません。請求サイクルごとに、サブスクリプション プランに基づいて割り当てられたクレジットを受け取ります。未使用のクレジットは、bi の終了時に期限切れになります。

安静期間。
SEO と AI SEO スクレイピングを大規模に実行する
低遅延でグローバル SERP にわたるブランドと競合他社を追跡します。
ChatGPT 、 Perplexity 、 Grok 、 Gemini 、 Google Search 、 Google News 、 Copilot 、および AI Overview から 1 つの API を通じて、リージョン全体で解析されたオブジェクトを含む構造化データを抽出します。
© 2026 Cloro Incorporation Ltd. 無断複写・転載を禁じます。

## Original Extract

cloro is an API for scraping SEO and AI SEO models — ChatGPT, Gemini, Perplexity, Copilot, AI Overview, and Google Search. Markdown, sources, and citations.

Hackers, We wanted to share our newly-added recurring Free tier for cloro.dev, the leading AI UI scraping platform in the world. We extract structured data from ChatGPT, Perplexity, Grok, Gemini, Google Search, Google News, Copilot, and AI Overview through one API, with low latency across any country. We have recently crossed 200+ customers, most of them the large Enterprises behind GEO and SEO. From now on, every user gets 500 free credits every month, which is great for hobby projects you may have. Enjoy!

cloro - the best scraper for SEO and AI SEO
Toggle Menu
Documentation
ChatGPT Perplexity Copilot Gemini AI Overview AI Mode Search The search API for the AI era
One API for every AI search engine and Google. Real-time, structured JSON.
4.8 · 42 reviews Try 500 credits for free Used by the top SEO and AI SEO companies 99.99% uptime 1,000,000,000 monthly API calls Structured data with a single API
Covers the main SEO and AI SEO providers,
at any scale and across regions.
ChatGPT Live Perplexity Live Copilot Live AI Mode Live AI Overview Live Search Live News Live Gemini Live Grok Unavailable
Bottom card (Amazon)
Amazon
Top card (Meta)
Meta
And more badge
+ more
--> Why teams pick cloro
Direct LLM API integration has hidden costs and difficult scalability.
API responses are nothing like the UI. Surfer proved this with data .
Direct provider APIs don't give the tools to influence SEO.
Each provider requires separate API integration, creating technical overhead.
Token-based pricing varies by model and provider, making costs unpredictable.
How the cloro integration works
Get back markdown , text , or HTML , plus parsed sources , citations , query fan-out , and shopping blocks pulled from each response.
from cloro import Cloro
client = Cloro(api_key="sk_live_your_api_key_here")
response = client.monitor.chatgpt(
prompt="What do you know about Tesla's latest updates?",
country="US",
include={"markdown": True},
)
print(response["result"]["text"]) typescript Copy import { Cloro } from "@cloro-dev/cloro";
const client = new Cloro({ apiKey: "sk_live_your_api_key_here" });
const response = await client.monitor.chatgpt({
prompt: "What do you know about Tesla's latest updates?",
country: "US",
include: { markdown: true },
});
console.log(response.result.text); curl Copy curl -X POST https://api.cloro.dev/v1/monitor/chatgpt \
-H "Authorization: Bearer sk_live_your_api_key_here" \
-H "Content-Type: application/json" \
-d '{
"prompt": "What do you know about Tesla\'s latest updates?",
"country": "US",
"include": {
"markdown": true
}
}' Response example
{
"success": true,
"result": {
"text": "Tesla's recent updates include significant improvements to their Full Self-Driving capability...",
"sources": [
{
"position": 1,
"url": "https://tesla.com/updates/fsd",
"label": "Tesla FSD Updates",
"description": "Latest Full Self-Driving improvements and capabilities"
}
],
"html": "<div class=\"markdown\"><p>Tesla's recent updates include...</p></div>",
"markdown": "**Tesla's recent updates** include significant improvements...",
"searchQueries": [
"Tesla updates 2024",
"Full Self Driving improvements"
],
"shoppingCards": [
{
"position": 1,
"product": {
"name": "Model Y",
"brand": "Tesla",
"price": "$43,990",
"currency": "USD",
"rating": 4.5,
"reviewCount": 2847,
"imageUrl": "https://example.com/tesla-model-y.jpg",
"productUrl": "https://tesla.com/modely",
"description": "All-electric compact SUV with Autopilot"
}
}
]
}
} Try 500 credits for free View documentation Pricing that scales with you
Start free. Price per credit drops as your volume grows — see every tier below.
Our API delivers structured responses in 30-45 seconds on average. We handle all the complexity of automation and parsing to give you clean markdown content and citations in near real-time.
We extract the complete AI response in parsed markdown format, all source & citations with URLs and labels. Everything is delivered as structured JSON for easy integration into your systems.
We support ChatGPT (with Web Search), Google Gemini, Perplexity, Microsoft Copilot, Google AI Overview, and Google Search. Our unified API lets you extract data from all these sources using a consistent interface.
Yes! We offer an async endpoint that allows you to submit tasks and receive the results via webhook. This is ideal for high-volume scraping and ensures reliability without keeping connections open.
Concurrency limits depend on your subscription plan. For example, business plans often start with higher concurrency limits (e.g., 100). If you exceed your limit, you will receive a 429 status code. We recommend implementing client-side queuing or using our async endpoint.
You can increase your concurrency by subscribing to an higher-volume package. Since we dedicate resources exclusively to your organization, we want to ensure those resources are fully utilized and not just spiky traffic. However, if you are indeed limited by concurrency (constantly using your concurrency more than 70% of the time), you should contact us so we can adjust.
Yes. By default, we return parsed structured JSON to keep responses lightweight. However, you can retrieve the full raw HTML of the result page by setting include.html: true in your request.
Yes, we support all countries globally. Note that some models (like Gemini) may have regional availability restrictions (e.g., EU).
You generally don't need to implement complex timeout logic. Our system automatically retries requests up to 10 times with a maximum duration of 5 minutes to ensure high success rates.
Credits are deducted only when we successfully extract and deliver AI response data. Failed extractions or system errors do not consume your credits.
No, credits do not roll over from month to month. Each billing cycle, you receive your allocated credits based on your subscription plan. Any unused credits expire at the end of the billing period.
Run SEO and AI SEO scrapes at scale
Track your brand and your competitors across global SERPs, with low latency.
Extract structured data from ChatGPT , Perplexity , Grok , Gemini , Google Search , Google News , Copilot , and AI Overview through one API, with parsed objects across regions.
© 2026 Cloro Incorporation Ltd. All rights reserved.

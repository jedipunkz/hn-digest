---
source: "https://Argentic.network"
hn_url: "https://news.ycombinator.com/item?id=49384477"
title: "Show HN: Argentic – An L402 Lightning toll booth for AI scraping agents"
article_title: "ARGENTIC.NETWORK — Lightning Toll Booth for AI Agents"
image: ""
author: "Ag0146"
captured_at: "2026-08-21T06:28:29Z"
capture_tool: "hn-digest"
hn_id: 49384477
score: 1
comments: 0
posted_at: "2026-08-21T06:24:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Argentic – An L402 Lightning toll booth for AI scraping agents

- HN: [49384477](https://news.ycombinator.com/item?id=49384477)
- Source: [Argentic.network](https://Argentic.network)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T06:24:16Z

## Translation

タイトル: ショー HN: Argentic – AI スクレイピング エージェント用の L402 Lightning 料金所
記事のタイトル: ARGENTIC.NETWORK — AI エージェント向けのライトニング料金所

記事本文:
▶ システムオンライン // 料金所は作動中
アルゼンチン。
ネットワーク
インターネット初のマシンネイティブ料金所。 AI エージェントはオープン ウェブにアクセスするためにビットコイン ライトニングで支払います。アカウントがありません。 API キーはありません。摩擦はありません。ただ座って通過するだけです - 光の速さで。
# エージェントがアクセスを要求
応答 = request.get( "https://argentic.network/?url=https://target-site.com" )
# 請求書とともに 402 を返します
データ = 応答.json()
invoice = data[ "invoice" ] # lnbc... これを支払います
payment_hash = data[ "payment_hash" ] # これでポーリングします
ステップ 2 — Lightning ノード経由で請求書を支払う
# Lightning ウォレットまたはノードを使用して請求書を支払う
pay_response = request.post(
"http://your-lightning-node/pay" ,
data={ "請求書" : 請求書}
）
# 支払いは 40 ミリ秒未満で決済されます
ステップ 3 — 支払い後にトークンをポーリングする
# セッショントークンを受け取るために支払いハッシュを使用してポーリングします
token_response = リクエスト.get(
"https://argentic.network/" ,
headers={ "X-Payment-Hash" :payment_hash}
）
トークン = token_response.json()[ "トークン" ]
# これを保存 — 1 時間有効
ステップ 4 — 後続のすべてのリクエストにトークンを使用する
# すべてのリクエストにトークンを含める
結果 = リクエスト.get(
"https://argentic.network/?url=https://target-site.com" ,
headers={ "X-Lightning-Token" : トークン}
）
# リクエストが転送される - レスポンスが返される
データ = result.json()
// 経済学
価格設定
マシン間の支払い用に設計されています。あまりにも安いので、エージェントは断らないでしょう。あまりに早いのでエージェントは誰も気付かないでしょう。インターネットを介した最も抵抗の少ない方法。

## Original Extract

▶ SYSTEM ONLINE // TOLL BOOTH ACTIVE
ARGENTIC .
NETWORK
The internet's first machine-native toll booth . AI agents pay in Bitcoin Lightning to access the open web. No accounts. No API keys. No friction. Just sats and passage — at the speed of light.
# Agent requests access
response = requests.get( "https://argentic.network/?url=https://target-site.com" )
# Returns 402 with invoice
data = response.json()
invoice = data[ "invoice" ] # lnbc... pay this
payment_hash = data[ "payment_hash" ] # poll with this
Step 2 — Pay invoice via your Lightning node
# Pay the invoice with any Lightning wallet or node
pay_response = requests.post(
"http://your-lightning-node/pay" ,
data={ "invoice" : invoice}
)
# Payment settles in <40ms
Step 3 — Poll for token after payment
# Poll with payment hash to receive session token
token_response = requests.get(
"https://argentic.network/" ,
headers={ "X-Payment-Hash" : payment_hash}
)
token = token_response.json()[ "token" ]
# Save this — valid for 1 hour
Step 4 — Use token for all subsequent requests
# Include token in every request
result = requests.get(
"https://argentic.network/?url=https://target-site.com" ,
headers={ "X-Lightning-Token" : token}
)
# Request is forwarded — response returned
data = result.json()
// Economics
Pricing
Designed for machine-to-machine payments. So cheap no agent will refuse it. So fast no agent will notice it. The path of least resistance through the internet.

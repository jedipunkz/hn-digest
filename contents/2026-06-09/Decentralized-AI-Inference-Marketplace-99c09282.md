---
source: "https://t4t.eth.link"
hn_url: "https://news.ycombinator.com/item?id=48459379"
title: "Decentralized AI Inference Marketplace"
article_title: "Token4Token — pay-per-token inference on Gnosis + Swarm"
author: "ffaerber"
captured_at: "2026-06-09T13:09:10Z"
capture_tool: "hn-digest"
hn_id: 48459379
score: 1
comments: 1
posted_at: "2026-06-09T10:57:51Z"
tags:
  - hacker-news
  - translated
---

# Decentralized AI Inference Marketplace

- HN: [48459379](https://news.ycombinator.com/item?id=48459379)
- Source: [t4t.eth.link](https://t4t.eth.link)
- Score: 1
- Comments: 1
- Posted: 2026-06-09T10:57:51Z

## Translation

タイトル: 分散型 AI 推論マーケットプレイス
記事のタイトル: Token4Token — Gnosis + Swarm でのトークンごとの支払い推論
説明: Token4Token (T4T) は、分散型 AI 推論マーケットプレイスです。クライアントは、ローカルでホストされているモデルで推論を実行するために xBZZ のプロバイダーに支払います。支払いは Gnosis Chain でエスクローされ、トラフィックは Ethereum Swarm でルーティングされます。

記事本文:
T4T // トークンにはトークン
v0.1
t4t.eth
ネットワークライブ
グノーシス・チェーン100
群れ
トークンごとの支払い推論 / API キーなし
推論
商品です。
価格を設定します。
GPU がプロンプトを入札するマーケットプレイス。クライアントは xBZZ でプロバイダーに支払い、Gnosis で支払いエスクローを支払い、すべてのバイトは Swarm 経由でルーティングされます。仲介者もレート制限もキーもありません。
４手目。それらを所有する中央事業者は存在しません。選択、支払い、評判、稼働状況はすべてチェーン上に存在します。
仲介者はいません。直接決済。鍵はありません。 AI トークンごとに支払い、失敗した場合は返金されます。
GPU と OpenAI 互換バックエンド (Ollama、vLLM、llama.cpp など) を備えている人は誰でも、T4T プロバイダー コンテナーを実行します。オンチェーンに登録し、 xBZZ をステークし、提供する各モデルの 100 万トークンあたりの価格を公開します。
ゲートウェイ コンテナーは、通常の OpenAI /v1/chat/completions エンドポイントをローカルに公開します。オンチェーン ディレクトリからプロバイダーを選択し、最大支払い額をエスクローし、暗号化されたプロンプトを Swarm PSS 経由で送信します。
プロバイダーがサービスを提供し、配達時に請求を行います。
プロバイダーはローカルで推論を実行し、Swarm 経由で応答を返し、エスクローからの支払いを請求します。期限が過ぎた場合、顧客は返金を受けます。障害が繰り返されると、プロバイダーの負担が大きくなります。
中央の API ゲートウェイはありません。マーケットプレイスの運営者はいません。レート制限はありません。コントラクト、ステーク、そして交換されるすべてのバイトのハッシュだけです。
Gnosis 上の ProviderRegistry からライブで取得されます。新しいハートビート (10 分以内) と少なくとも 1 つのオファリングを持つプロバイダーのみが表示されます。
コンテナが2つ。どちらかの側を選びます。どちらも、最初の実行時にウォレットを作成し、Swarm 切手をバインドするオンボーディング UI が付属しています。
ゲートウェイコンテナを実行します。既存の OpenAI SDK を http://localhost:8080/v1 に指定します。 xBZZ ではトークンごとに支払います。
ローカルの Ollama / vLLM に対してプロバイダー コンテナーを実行します。 xBZZを賭けます。世界中のクライアントからトークンごとの収益を獲得します。

## Original Extract

Token4Token (T4T) is a decentralized AI inference marketplace. Clients pay providers in xBZZ to run inference on locally-hosted models, with payment escrowed on Gnosis Chain and traffic routed over Ethereum Swarm.

T4T // TOKEN FOR TOKEN
v0.1
t4t.eth
NETWORK LIVE
GNOSIS · CHAIN 100
SWARM
PAY-PER-TOKEN INFERENCE / NO API KEY
INFERENCE
IS A COMMODITY .
PRICE IT.
A marketplace where GPUs bid for prompts. Clients pay providers in xBZZ , payment escrows on Gnosis , and every byte routes over Swarm . No middleman, no rate limit, no key.
Four moves. No central operator owns any of them. Selection, payment, reputation, and liveness all live on-chain.
No middleman. Direct settlement. No key. Pay per AI token, refund on miss.
Anyone with a GPU and an OpenAI-compatible backend (Ollama, vLLM, llama.cpp…) runs the T4T provider container. It registers on-chain, stakes xBZZ , and publishes per-million-token prices for each model it serves.
The gateway container exposes a normal OpenAI /v1/chat/completions endpoint locally. It picks a provider from the on-chain directory, escrows the maximum payment, and ships the encrypted prompt over Swarm PSS.
Provider serves, claims on delivery.
The provider runs inference locally, returns the response over Swarm, and claims payment from escrow. If the deadline lapses, the client gets a refund. Repeated faults slash the provider's stake.
No central API gateway. No marketplace operator. No rate limit. Just a contract, a stake, and a hash for every byte exchanged.
Pulled live from ProviderRegistry on Gnosis. Only providers with a fresh heartbeat (≤ 10 min) and at least one offering appear.
Two containers. Pick a side. Both ship with onboarding UIs that mint a wallet and bind a Swarm postage stamp on first run.
Run the gateway container. Point your existing OpenAI SDK at http://localhost:8080/v1 . Pay per token in xBZZ.
Run the provider container against your local Ollama / vLLM. Stake xBZZ. Earn per-token revenue from clients worldwide.

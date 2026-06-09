---
source: "https://superhighway.walls.sh"
hn_url: "https://news.ycombinator.com/item?id=48454470"
title: "Show HN: Web tools an AI agent pays for per call in USDC, no API key (x402+MCP)"
article_title: "Superhighway — the web-search API agents can pay for"
author: "patwalls"
captured_at: "2026-06-09T00:52:58Z"
capture_tool: "hn-digest"
hn_id: 48454470
score: 2
comments: 0
posted_at: "2026-06-09T00:30:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Web tools an AI agent pays for per call in USDC, no API key (x402+MCP)

- HN: [48454470](https://news.ycombinator.com/item?id=48454470)
- Source: [superhighway.walls.sh](https://superhighway.walls.sh)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T00:30:43Z

## Translation

タイトル: HN を表示: AI エージェントが USDC で通話ごとに支払う Web ツール、API キーなし (x402+MCP)
記事のタイトル: スーパーハイウェイ — Web 検索 API エージェントは料金を支払うことができます

記事本文:
エージェント向けに製品化された情報スーパーハイウェイ。 AI エージェントが独自に料金を支払うことができる Web 検索 API。サインアップも API キーも人間も関与しません。
検索ごとに 0.001 ドル、x402 プロトコル経由でUSDC で支払われます。
エージェントは GET s /search?q=... を取得します。
価格と支払い先を含む 402 Payment Required が返されます。
x402 クライアントは USDC 支払いに署名し、自動的に再試行します。
きれいな JSON 検索結果が得られます。人間の関与はまったくありません。
未払いのリクエスト — 402 とその支払い条件を参照してください。
カール -i "https://api-production-17e1.up.railway.app/search?q=best+espresso+machines"
有料リクエスト — x402 クライアント (TypeScript) を使用:
import {wrapFetchWithPayment, createSigner} from "x402-fetch";
const signed = await createSigner("base", process.env.AGENT_PRIVATE_KEY);
const pay = WrapFetchWithPayment(fetch, 署名者);
const res = await pay("https://api-production-17e1.up.railway.app/search?q=best+espresso+machines");
console.log(await res.json());
エージェント向け
機械可読仕様: /openapi.json · ヘルス: /health
SearXNG + x402 上に構築されています。一度に 1 つのクエリで、従量課金制です。

## Original Extract

The information superhighway, productized for agents. A web-search API your AI agent can pay for on its own — no signup, no API key, no human in the loop.
$0.001 per search, paid in USDC on base via the x402 protocol.
Your agent GET s /search?q=... .
It gets back 402 Payment Required with the price and where to pay.
An x402 client signs a USDC payment and retries automatically.
It gets clean JSON search results. Total human involvement: zero.
Unpaid request — see the 402 and its payment terms:
curl -i "https://api-production-17e1.up.railway.app/search?q=best+espresso+machines"
Paid request — with an x402 client (TypeScript):
import { wrapFetchWithPayment, createSigner } from "x402-fetch";
const signer = await createSigner("base", process.env.AGENT_PRIVATE_KEY);
const pay = wrapFetchWithPayment(fetch, signer);
const res = await pay("https://api-production-17e1.up.railway.app/search?q=best+espresso+machines");
console.log(await res.json());
For agents
Machine-readable spec: /openapi.json · Health: /health
Built on SearXNG + x402 . One query at a time, paid as you go.

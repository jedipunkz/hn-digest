---
source: "https://docs.bankofai.io/"
hn_url: "https://news.ycombinator.com/item?id=48940276"
title: "Bank of AI"
article_title: "Welcome to x402 | BANK OF AI | Developer Guide"
author: "2a0c40"
captured_at: "2026-07-16T21:57:18Z"
capture_tool: "hn-digest"
hn_id: 48940276
score: 1
comments: 0
posted_at: "2026-07-16T21:08:08Z"
tags:
  - hacker-news
  - translated
---

# Bank of AI

- HN: [48940276](https://news.ycombinator.com/item?id=48940276)
- Source: [docs.bankofai.io](https://docs.bankofai.io/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T21:08:08Z

## Translation

タイトル: AI銀行
記事のタイトル: x402 へようこそ | AI銀行 |開発者ガイド
説明: このガイドでは、ブロックチェーン上の x402 オープン ペイメント標準を紹介し、x402 を利用したサービスの構築または統合を開始するのに役立ちます。

記事本文:
メインコンテンツにスキップ
x402 支払いプロトコルの概要
MCPサーバーとスキルの紹介
Openclaw 拡張機能の概要
x402 は、HTTP 402 Payment Required ステータス コードに基づいて構築されたオープン ブロックチェーン支払い標準です。これにより、Web サービスは、従来のアカウント システムやセッション管理に依存することなく、「応答前支払い」メカニズムを通じて API またはコンテンツの料金を請求できるようになります。
x402 は現在、TRON および BSC ネットワークをサポートしていますが、将来的にはより広範なマルチチェーン エコシステムに拡張する予定です。
ドキュメントへの貢献に興味がありますか?
ご自由に PR を GitHub リポジトリに送信してください。私たちの唯一の原則は中立性を維持することです。重要なリソースへのリンクを除き、宣伝コンテンツやブランド コンテンツは避けてください。
x402 で構築しますか?
公開された @bankofai/x402-* パッケージを TypeScript アプリにインストールします。 x402 リポジトリには、実行可能なクライアント → サーバー → ファシリテーターのサンプルが、examples/typescript/ の下に参考として同梱されています (例:exact、gasfree、upto、batch-settlement)。
x402 は、既存の決済システムの中核となる問題点に対処します。
従来のクレジットカードや法定通貨の支払いチャネルにおける高額な手数料と複雑なワークフロー
自律型 AI エージェントのトランザクションなど、マシンツーマシン (M2M) 支払いのサポートが不十分
効率的な少額決済インフラストラクチャの欠如により、使用量に基づく収益化が困難になる
ブロックチェーンの高速かつ低コストのトランザクション機能を活用
販売者: API またはコンテンツを収益化したいサービス プロバイダー。 x402 は最小限の構成で、クライアントからの直接のプログラムによる支払いを可能にします。
購入者: 登録フローや手動介入なしで有料サービスにアクセスしたい開発者と AI エージェント。
買い手と売り手は HTTP リクエストを通じて直接やり取りしますが、支払いは HTTP リクエストを通じて実行されます。

プロトコルによって、明示的かつ自動的にオンチェーンに追加されます。
x402 は、幅広い使用例をサポートします。
API の料金を自律的に支払うことができる AI エージェント
マイクロトランザクションを通じて収益化されるマイクロサービス
基盤となる機能を再販しない API アグリゲーション サービス
大まかに言うと、ワークフローは簡単です。
リクエストが開始されました: 購入者がサーバーから保護されたリソースをリクエストします。
支払いが必要: 支払いが必要な場合、サーバーは支払い指示とともに 402 Payment Required 応答を返します。
支払いが送信されました: 購入者は署名済みの支払いペイロードを生成して送信します。
検証と決済: サーバーは、x402 ファシリテーターの /verify および /settle エンドポイントを呼び出して、支払いを検証して決済します。
リソースの配信: 検証が成功すると、サーバーは要求されたリソースを配信します。
私たちの目標は、ブロックチェーン上に、障壁が低く、パーミッションレスで、開発者に優しいプログラマブルコマースレイヤーを構築することです。
x402 は現在、次のネットワークをサポートしています。
TRON シャスタ テストネット ( tron:shasta )
TRON ナイル テストネット (tron:nile)
SDK (TypeScript のみ) : x402 は、粒度の高い @bankofai/x402-* npm パッケージ ( core 、 evm 、 tron 、 fetch 、 Express 、 hono 、 fastify 、 next 、 axios 、 mcp 、 extensions ) として公開されている TypeScript のみの SDK です。ソースは pnpm/turbo モノリポジトリで維持されますが、アプリケーション開発では公開されたパッケージをインストールする必要があります。サポートされているスキーム：exact (ERC-3009 / Permit2)、upto、batch-settlement、auth-capture (EVM)、およびexact_gasfree (TRON)。前世代の Python + TypeScript SDK は参考までに、legacy/ の下にあります。完全な内訳については、SDK 機能マトリックスを参照してください。

## Original Extract

This guide introduces the x402 open payment standard on blockchain and helps you start building or integrating x402-powered services.

Skip to main content English English
x402 Payment Protocol Introduction
MCP Server & SKILLS Introduction
Openclaw Extension Introduction
x402 is an open blockchain payment standard built on the HTTP 402 Payment Required status code. It enables web services to charge for APIs or content through a “pay-before-response” mechanism — without relying on traditional account systems or session management.
x402 currently supports the TRON and BSC networks, with plans to expand to a broader multi-chain ecosystem in the future.
Interested in contributing to the documentation?
Feel free to submit a PR to the GitHub repository . Our only principle is to maintain neutrality — aside from essential resource links, please avoid promotional or branded content.
Building with x402?
Install the published @bankofai/x402-* packages in your TypeScript app. The x402 repository ships runnable client → server → facilitator examples under examples/typescript/ for reference — exact , gasfree , upto , and batch-settlement .
x402 addresses the core pain points of existing payment systems:
High fees and complex workflows in traditional credit card and fiat payment channels
Poor support for machine-to-machine (M2M) payments , such as autonomous AI agent transactions
Lack of efficient micro-payment infrastructure , making usage-based monetization difficult
Leveraging blockchain’s fast and low-cost transaction capabilities
Sellers: Service providers who want to monetize APIs or content. With minimal configuration, x402 enables direct, programmatic payments from clients.
Buyers: Developers and AI agents who want to access paid services without registration flows or manual intervention.
Buyers and sellers interact directly through HTTP requests, while payments are executed transparently and automatically on-chain by the protocol.
x402 supports a wide range of use cases:
AI agents capable of autonomously paying for APIs
Microservices monetized via microtransactions
API aggregation services that do not resell underlying capabilities
At a high level, the workflow is straightforward:
Request Initiated: The buyer requests a protected resource from the server.
Payment Required: If payment is required, the server returns a 402 Payment Required response along with payment instructions.
Payment Submitted: The buyer generates and submits a signed payment payload.
Verification & Settlement: The server calls the x402 Facilitator’s /verify and /settle endpoints to validate and settle the payment.
Resource Delivered: Once verification succeeds, the server delivers the requested resource.
Our goal is to build a low-barrier, permissionless, developer-friendly programmable commerce layer on blockchain.
x402 currently supports the following networks:
TRON Shasta Testnet ( tron:shasta )
TRON Nile Testnet ( tron:nile )
SDK (TypeScript-only) : x402 is a TypeScript-only SDK published as granular @bankofai/x402-* npm packages ( core , evm , tron , fetch , express , hono , fastify , next , axios , mcp , extensions ). The source is maintained in a pnpm/turbo monorepo, but application development should install the published packages. Supported schemes: exact (ERC-3009 / Permit2), upto , batch-settlement , auth-capture (EVM), and exact_gasfree (TRON). The previous-generation Python + TypeScript SDK lives under legacy/ for reference. See the SDK Feature Matrix for the full breakdown.

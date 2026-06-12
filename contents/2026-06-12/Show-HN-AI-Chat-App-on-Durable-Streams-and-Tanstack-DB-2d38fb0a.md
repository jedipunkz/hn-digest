---
source: "https://oss.chat/tour"
hn_url: "https://news.ycombinator.com/item?id=48506809"
title: "Show HN: AI Chat App on Durable Streams and Tanstack DB"
article_title: "Open Chat — the guided tour"
author: "sorenbs"
captured_at: "2026-06-12T18:46:33Z"
capture_tool: "hn-digest"
hn_id: 48506809
score: 2
comments: 0
posted_at: "2026-06-12T17:18:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Chat App on Durable Streams and Tanstack DB

- HN: [48506809](https://news.ycombinator.com/item?id=48506809)
- Source: [oss.chat](https://oss.chat/tour)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T17:18:58Z

## Translation

タイトル: HN を表示: Durable Streams と Tanstack DB 上の AI チャット アプリ
記事のタイトル: オープン チャット — ガイド付きツアー
説明: プロンプトがどのように耐久性があり、再生可能なストリームになるか: Prisma Streams、Prisma Postgres、Prisma Next、Prisma Compute、および TanStack DB が連携するアニメーション ウォークスルー。
HN テキスト: ElectricSQL の Durable Streams をストレージ形式とトランスポート メディアの両方として使用する AI チャット アプリを構築しました。これには、トークンが LLM プロバイダーから返送されるときに永続的に保存され、その後フロントエンドにストリーミングされるという非常に優れた特性があり、スムーズなユーザー エクスペリエンスを実現します。フロントエンドは Tanstack DB の差分データフロー実装を使用して、Postgres からのメタデータを Durable Stream のチャット ストリームと結び付けます。すべてのコンポーネントの詳細な説明を作成しました。

記事本文:
オープンチャット — ガイド付きツアー

## Original Extract

How a prompt becomes a durable, replayable stream: an animated walkthrough of Prisma Streams, Prisma Postgres, Prisma Next, Prisma Compute, and TanStack DB working together.

I built an AI chat app that uses Durable Streams from ElectricSQL as both the storage format and transport medium. It has the very nice property that tokens are stored durably as they are steamed back from the LLM provider, and then streamed to the frontend for a smooth user experience. The frontend uses Tanstack DBs differential dataflow implementation to tie together meta data from Postgres with the chat streams in the Durable Stream. Made a detailed explainer of all the components.

Open Chat — the guided tour

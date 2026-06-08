---
source: "https://rayline.ai/"
hn_url: "https://news.ycombinator.com/item?id=48448372"
title: "Show HN: Rayline routes Claude Code subagents to on-device and cheaper models"
article_title: "Rayline — AI Routing, Optimized"
author: "davidvgilmore"
captured_at: "2026-06-08T18:57:44Z"
capture_tool: "hn-digest"
hn_id: 48448372
score: 8
comments: 8
posted_at: "2026-06-08T17:32:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rayline routes Claude Code subagents to on-device and cheaper models

- HN: [48448372](https://news.ycombinator.com/item?id=48448372)
- Source: [rayline.ai](https://rayline.ai/)
- Score: 8
- Comments: 8
- Posted: 2026-06-08T17:32:42Z

## Translation

タイトル: HN を表示: レイラインがクロード コードのサブエージェントをデバイス上の安価なモデルにルーティングする
記事のタイトル: Rayline — AI ルーティング、最適化
説明: Rayline は、フロンティア モデル、スペシャリスト モデル、および Mac 間でエージェントのワークロードをルーティングします。エンドポイントは 1 つ。どのモデルも。
HN テキスト: こんにちは、HN、私は Rayline のビルダーの 1 人です。 Rayline は、Claude Code と互換性のある LLM ゲートウェイです。これは、クロード コードの内部ルーティングをインターセプトしてオーバーライドし、代わりにサブエージェント呼び出しを別のモデルにルーティングできるようにします。たとえば、メイン エージェントを Opus で実行し、一部のサブエージェントをクラウド ホストのオープン モデルで実行し、他のサブエージェントをデバイス上で実行できます。他の人が、エージェントが呼び出せるツールとしてクロード コードのルーティングを実装しているのを見てきました。私たちの経験では、これはうまく機能しません。メイン エージェントがトークンを使用して検討し、ツールを呼び出す必要があり、LLM は一般にルーティングの決定を行うための非常に非効率的な方法だからです。 Rayline をゲートウェイとして実装することで、ユーザーがルーティングの決定を決定的に構成できるようになり、オプションで ML モデルを使用してルーティングの決定を行うことができます。クロード コード セッションには、すべてが同じモデルを必要としないサブエージェント呼び出しが多数含まれていることに気づいて、これを構築しました。他のルーターも存在しますが、私たちは引き続きクロード コード (個別のハーネスなし) を使用し、サブエージェント レベルでタスクをルーティングし、クラウドとオンデバイス全体でルーティングできるようにするために Rayline を構築しました。
メインエージェントは Opus の恩恵を受けることがよくあります。しかし、委任された呼び出しの多くは範囲が狭く、リポジトリの検索、コンテキストの要約、エラーの検査、CI 更新のポーリングなどです。私たちが調査しているのは、サブエージェント レベルのルーティングです。コーディング エージェントの主なコスト要因は通常、キャッシュされた入力とキャッシュされていない入力です。サブエージェントの委任は、キャッシュの破壊を避けるため、ルーティングを決定する自然なポイントになります。 de のメッセージ スレッド コンテキストを調べます。

指定された通話を選択し、その通話のモデルを選択します。タスク レベルでは、Sonnet と Haiku はほとんどの場合、オープン モデルよりも 1 ドルあたりの機能が低いため、主な利点は優れた + (はるかに) 安価なサブエージェント (プライベート ベータ版では 60 ～ 90%) です。過去 2 週間で全世界がモデル ルーティングについて話し始めたようです。そのため、モデル ルーティングが関連する製品分野であることに他の人も同意しているようです。 HN コミュニティからのフィードバックをお待ちしています。

記事本文:
Rayline — AI ルーティング、最適化

## Original Extract

Rayline routes agent workloads across frontier models, specialist models, and your Mac. One endpoint. Every model.

Hi HN, I’m one of the builders of Rayline. Rayline is a Claude Code compatible LLM gateway. It intercepts and overrides claude code’s internal routing and lets you route subagent calls to different models instead. For example, you can run the main agent on Opus, some subagents on cloud-hosted open models, and other subagents on-device. We’ve seen others implement routing for claude code as tools the agent can invoke. In our experience, that doesn’t work well because it requires the main agent to use tokens to think about + call the tools, and LLMs are generally a very inefficient way to make routing decisions. By implementing Rayline as a gateway, we let users deterministically configure routing decisions, and you can optionally use our ML model to make routing decisions. We built it after noticing that Claude Code sessions contain a lot of subagent calls that don’t all need the same model. Other routers exist, but we built Rayline to let us continue using claude code (no separate harness), route tasks at a subagent level, and route across cloud and on-device.
The main agent often benefits from Opus. But many delegated calls have narrow scope: search the repo, summarize context, inspect an error, poll for CI updates, etc. The thing we’re exploring is subagent-level routing. The main cost lever in coding agents is usually cached vs non-cached input. Subagent delegations are a natural point to make routing decisions because you avoid busting cache. We look at the message-thread context for a delegated call and choose a model for that call. At a task level, Sonnet and Haiku are almost always less capability-per-dollar than open models, so the main advantage is better + (much) cheaper subagents (60-90% in our private beta). The whole world seems to have started talking about model routing in the past two weeks, so apparently others agree it’s a relevant product area. We’d love to get feedback from the HN community!

Rayline — AI Routing, Optimized

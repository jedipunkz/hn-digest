---
source: "https://blog.mozilla.ai/what-is-an-llm-control-plane/"
hn_url: "https://news.ycombinator.com/item?id=48508913"
title: "What Is an LLM Control Plane?"
article_title: "What is an LLM control plane?"
author: "angpt"
captured_at: "2026-06-12T21:38:06Z"
capture_tool: "hn-digest"
hn_id: 48508913
score: 2
comments: 0
posted_at: "2026-06-12T20:11:24Z"
tags:
  - hacker-news
  - translated
---

# What Is an LLM Control Plane?

- HN: [48508913](https://news.ycombinator.com/item?id=48508913)
- Source: [blog.mozilla.ai](https://blog.mozilla.ai/what-is-an-llm-control-plane/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T20:11:24Z

## Translation

タイトル: LLM コントロール プレーンとは何ですか?
記事のタイトル: LLM コントロール プレーンとは何ですか?
説明: 暴走エージェント?プロバイダーの停止?本番ルーティング、予算、プライバシーを処理するために、AI スタックに単なるゲートウェイではなく LLM コントロール プレーンが必要な理由を説明します。

記事本文:
サインイン
購読する
LLM コントロール プレーンとは何ですか?
暴走エージェント？プロバイダーの停止?本番ルーティング、予算、プライバシーを処理するために、AI スタックに単なるゲートウェイではなく LLM コントロール プレーンが必要な理由を説明します。
エージェントが推論ループに陥ってもクラッシュしません。誰かが請求に気づくまで、毎月の予算を静かに使い切ってしまいます。 1 週間後、プロバイダーが停止し、それをキャッチするフォールバックがなかったため、アプリもそれに伴ってダウンします。セキュリティ チームは、モデルが先週どのプロバイダーにどのようなデータを送信したかを尋ねますが、正直な答えは、よくわかりません。コストは徐々に上昇しており、どのアプリ、どのモデル、またはどのチームが責任を負っているのかは言えません。
残念なことに、これは実稼働環境で LLM を実行するデフォルトの状態です。このようなことが起こり続ける理由は、ほとんどのチームが独自の方法で「AI」を活用しているためです。ルーティング ロジックをアプリケーション層にボルトで組み込んで構築する場合もあります。後付けでトークンを追跡するものもあります。ひどいことに、どのチームも同じ配管を最初から再構築します。それを処理するための標準プロトコルがないからです。
私たちはインフラストラクチャの他のあらゆる場所でこの問題を解決しました。 API ゲートウェイ、サービス メッシュ、Kubernetes コントロール プレーン。 LLM トラフィックに相当するものはありませんでした。ここで、LLM コントロール プレーンが登場します。
おそらく、これらの組み合わせを聞いたことがあるでしょう: {AI、LLM} × {ゲートウェイ、ルーター、プロキシ} に「コントロール プレーン」を加えたもの。これらの用語は同じ意味で使用されますが、それらの間の境界線は、優れたデモと本番環境で強化されたソフトウェアの間の境界線となります。
ゲートウェイはメカニカル レイヤーを処理します。つまり、リクエストのルーティング、API キーの管理、レート制限の強制を行います。アプリは 5 つのエンドポイントではなく 1 つのエンドポイントと通信します。初期段階のプロジェクトでは、多くの場合、これで十分です。
コントロール プレーンは配管だけでなく決定を処理します。それが違いです

「このリクエストは通過したか」と「このリクエストは通過するべきか」の間。リクエストの実行後に予算制限を集計するのではなく、リクエストの実行前に予算制限を適用し、10 個のコピー＆ペーストされたバージョンではなく 1 つのポリシーをすべてのアプリとモデルに適用し、プロバイダーが停止したときにフェールオーバーします。
ほとんどのチームはすぐにゲートウェイを超えてしまいます。プロキシを追加し、ルーティングとログを取得します。その後、使用量が増加するにつれて、本当の疑問が生じます。1 人の暴走エージェントが予算を吹き飛ばすのをどのように阻止するかということです。通話ごとだけでなく、複数のユーザーとセッションにわたる支出をどのように追跡しますか?ゲートウェイはそのために構築されたものではなく、コントロール プレーンがそのために構築されました。
LLM インフラストラクチャの 3 つの面
これは新しい問題ではありません。ネットワーク (およびその他) は、数十年前にそのバージョンを解決しました。
その秘訣は、1 つのモノリシック システムの構築をやめて、それを複数のプレーンに分割することでした。データ プレーンはトラフィックを移動します。コントロール プレーンは、どこへ行くか、何が許可されるかを決定します。管理プレーンは、人間が全体を設定し監視する場所です。同じ分割が LLM インフラストラクチャに明確にマッピングされます。 Kubernetes を使用したことがある場合は、すでにこれを見たことがあるでしょう。コントロール プレーンは決定と強制を行う部分であり、ワークロードは実行されるだけです。
現在、ほとんどの LLM インフラストラクチャは、データ プレーンと管理プレーンの薄いスライスをカバーしています。 LLM コントロール プレーンはギャップであり、最終的には誰もが手作業で構築することになるレイヤーであり、まさに標準インフラストラクチャであるべきレイヤーです。
コントロール プレーンが実際に行う必要があること
この言葉は曖昧に使われるので、コンクリートのバーを設置しています。実際の LLM コントロール プレーンは、これのほとんどを処理する必要があります。
ハード予算制限により、しきい値を超えた瞬間に新しいリクエストがブロックされ、暴走したエージェントが設定された制限を超えて支出するのを防ぎます。
ユーザーとセッション全体で支出が追跡されます。通話ごとの番号は簡単です。

「このチームの先月の費用はいくらだったのか」ということが実際に重要なのです。
プロバイダー間のポリシーに基づくルーティング。プロバイダーがダウンした場合は自動フェイルオーバーが行われます。
すべてのサービスで同じチェックを再実装するのではなく、プロンプトと応答にガードレールを 1 か所で適用できます。
完全な監査証跡。セキュリティや財務部門からの要請があった場合に、すべてのリクエスト、レスポンス、ルーティング決定が記録されます。
プロバイダーの資格情報は 1 つのコンテナー内にあり、十数の env ファイルに分散されていません。
それは、すべてのチームが自分たちで再構築しなければならない配管です。コントロール プレーンの重要な点は、その必要がなくなったことです。
コントロール プレーンは、すべてのプロンプト、完了、資格情報の実行パスに直接配置されます。これは AI スタックの中で最も敏感なポイントであるため、どこで実行されるかは些細なことではありません。
ほとんどのツールでは、プライバシーを保証するために複雑なインフラストラクチャを自己ホストするか、プライバシーが契約と同等に保たれる SaaS プロバイダーにトラフィックを渡すかというトレードオフに直面します。それは疑問に値するトレードオフです。境界を所有することとインフラストラクチャを自分で実行する必要がないことは、相互に排他的であってはなりません。
このトレードオフは、私たちが小谷を構築して解消しようとしている問題です。これは、ルーティング、予算、ガードレール、可観測性を 1 か所で処理するオープンソースの LLM コントロール プレーンです。データが要求したときに自己ホストするか、どちらの場合でもキー、プロンプト、応答が自分のものになるように構築された管理された展開を使用します。妥協点ではなく、自分の境界線を選びましょう。
小谷はまだ早い。 LLM インフラストラクチャを拡張していて、独自の配管を手作業で構築するのをやめたい場合は、Otari のコードとドキュメントから始めるのが良いでしょう。
[GitHub リポジトリを探索する] · [Otari ベータ版に参加する]
OpenCodeでOtariゲートウェイを使用する
AI コーディング セッションはブラック ボックスのように感じることがあります。 OpenCodeをOtar Gateway経由でルーティングしてコストを追跡し、

Ken の使用状況とモデルのアクティビティをリアルタイムで確認できます。アプリケーション コードを 1 行も変更することなく、すべてのセッションにわたる予算の管理と可視性を実現します。
any-llm を利用したオープンソース LLM ゲートウェイである Otari と、同じ基盤上に構築されたホスト型プラットフォームである Otari.ai をご紹介します。使用状況追跡、予算管理、ルーティング ポリシー、可観測性、チーム管理を備えた 1 つの API を通じてフロンティア モデルまたはオープンウェイト モデルを実行します。
クラウド AI の価格は 2026 年に急速に変更されました。この投稿では、なぜ多くのチームがローカル モデルに戻るのか、Ollama や LM Studio などのツールの背後にあるトレードオフ、および開発者にとってポータビリティと所有権が大きな懸念事項になっている理由について考察します。
CQ Exchange: 国境なきエージェント
cq Exchange は、エージェントに、プライベートな名前空間とパブリック コモンズを通じて、経験に基づいた知識を保存および取得するための共有場所を提供します。
購読してチームから最新のニュースやアイデアを入手してください

## Original Extract

Runaway agents? Provider outages? Discover why your AI stack needs an LLM control plane, not just a gateway, to handle production routing, budgets, and privacy.

Sign in
Subscribe
What is an LLM control plane?
Runaway agents? Provider outages? Discover why your AI stack needs an LLM control plane, not just a gateway, to handle production routing, budgets, and privacy.
An agent stuck in a reasoning loop doesn't crash. It just quietly burns through your monthly budget until someone notices the bill. A week later, a provider has an outage and your app goes down with it, because there was no fallback to catch it. Your security team asks what data your model sent to which provider last week, and the honest answer is you don't really know. Costs are creeping up and you can't say which app, which model, or which team is responsible.
It's sadly the default state of running LLMs in production. The reason it keeps happening is that most teams are working with "AI" in their own way. Some build routing logic by bolting it into the application layer. Some track tokens on the side as an afterthought. Every team rebuilds the same plumbing from scratch, badly, because there's no standard protocol for handling it.
We solved this everywhere else in infrastructure. API gateways, service meshes, Kubernetes control planes. There's just never been an equivalent for LLM traffic. This is where an LLM control plane comes in.
You've probably heard some combination of these: {AI, LLM} × {gateway, router, proxy} , plus "control plane." The terms get used interchangeably, but the line between them is the line between pretty demos and production-hardened software.
A gateway handles the mechanical layer: it routes requests, manages API keys, enforces rate limits. Your app talks to one endpoint instead of five. For an early-stage project, that's often good enough.
A control plane handles the decisions, not just the plumbing. It's the difference between "did this request go through" and "should this request go through at all." It enforces budget limits before a request runs instead of tallying them up after, applies one policy across every app and model instead of ten copy-pasted versions, and fails over when a provider dies.
Most teams outgrow a gateway fast. You add a proxy, get routing and logging, and then as usage increases the real questions land: how do you stop one runaway agent from blowing up your budget? How do you track spend across multiple users and sessions, not just per call? A gateway wasn't built for that, a control plane is.
The three planes of LLM Infrastructure
This is not a new problem. Networking (and others) solved a version of it decades ago.
The trick was to stop building one monolithic system and split it into planes. The data plane moves the traffic. The control plane decides where it goes and what's allowed. The management plane is where humans configure and watch the whole thing. That same split maps cleanly onto LLM infrastructure. If you've worked with Kubernetes, you've already seen this: the control plane is the part that decides and enforces, while the workloads just run.
Today most LLM infrastructure covers the data plane and a thin slice of the management plane. The LLM control plane is the gap, the layer everyone ends up hand-building, which is exactly the layer that should be standard infrastructure.
What a control plane actually needs to do
The term gets used loosely, so I’m setting a concrete bar. A real LLM control plane should handle most of this:
Hard budget limits that block new requests the moment a threshold is crossed, preventing a runaway agent from spending a penny over your configured limit.
Spend tracked across users and sessions. Per-call numbers are easy; "what did this team cost last month" is the one that actually matters.
Policy-driven routing across providers, with automatic failover when one goes down.
One place to apply guardrails on prompts and responses, instead of reimplementing the same checks in every service.
A full audit trail. Every request, response, and routing decision logged for when security or finance comes asking.
Provider credentials in one vault, not scattered across a dozen env files.
It's the plumbing every team has to rebuild for itself. The point of a control plane is that you don't need to anymore.
A control plane sits directly in the execution path of every prompt, completion, and credential. It's the most sensitive point in your AI stack, so where it runs isn't a minor detail.
With most tools you face a tradeoff: self-host a complex piece of infrastructure to guarantee privacy or hand your traffic to a SaaS provider where privacy is only as good as the contract. That's the tradeoff worth questioning. Owning the boundary and not having to run the infrastructure yourself shouldn't be mutually exclusive.
That tradeoff is the problem we're building Otari to remove. It's an open-source LLM control plane that handles routing, budgets, guardrails, and observability in one place. Self-host it when the data demands it, or use a managed deployment built so your keys, prompts, and responses stay yours either way. Pick your boundary, not the compromise.
Otari is still early. If you're scaling LLM infrastructure and want to stop hand-building your own plumbing, the Otari code and docs are a good place to start.
[ Explore the GitHub repo ] · [ Join the Otari beta ]
Use the Otari Gateway with OpenCode
AI coding sessions can feel like a black box. Route OpenCode through the Otari Gateway to track costs, token usage, and model activity in real time. Get budget controls and visibility across every session without changing a single line of application code.
Meet Otari, an open-source LLM gateway powered by any-llm, and Otari.ai, the hosted platform built on the same foundation. Run frontier or open-weights models through one API with usage tracking, budget controls, routing policies, observability, and team management.
Cloud AI pricing changed fast in 2026. This post looks at why more teams are moving back to local models, the tradeoffs behind tools like Ollama and LM Studio, and why portability and ownership are becoming bigger concerns for developers.
cq exchange: Agents without Borders
cq exchange gives agents a shared place to store and retrieve experience-driven knowledge through private namespaces and a public commons.
Subscribe to get the latest news and ideas from our team

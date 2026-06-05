---
source: "https://aaif.io/blog/agentgateway-joins-aaif-as-an-open-gateway-for-agentic-ai-infrastructure/"
hn_url: "https://news.ycombinator.com/item?id=48407648"
title: "agentgateway Joins AAIF as an Open Gateway for Agentic AI Infrastructure"
article_title: "agentgateway Joins AAIF as an Open Gateway for Agentic AI Infrastructure - Agentic AI Foundation (AAIF)"
author: "wicket"
captured_at: "2026-06-05T04:33:51Z"
capture_tool: "hn-digest"
hn_id: 48407648
score: 2
comments: 0
posted_at: "2026-06-05T03:39:26Z"
tags:
  - hacker-news
  - translated
---

# agentgateway Joins AAIF as an Open Gateway for Agentic AI Infrastructure

- HN: [48407648](https://news.ycombinator.com/item?id=48407648)
- Source: [aaif.io](https://aaif.io/blog/agentgateway-joins-aaif-as-an-open-gateway-for-agentic-ai-infrastructure/)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T03:39:26Z

## Translation

タイトル: Agentgateway が Agentic AI インフラストラクチャのオープン ゲートウェイとして AAIF に参加
記事タイトル: Agentgateway、Agentic AI インフラストラクチャのオープン ゲートウェイとして AAIF に参加 - Agentic AI Foundation (AAIF)
説明: Agentic AI Foundation は、MCP、エージェント間、LLM、および API トラフィック専用に構築されたオープン ソース ゲートウェイである Agentgateway を最新のホスト プロジェクトとして歓迎します。

記事本文:
メインコンテンツにスキップ
MCP Dev Summit ベンガルールにご参加ください • 6 月 9 ～ 10 日 • 今すぐ登録
検索
検索を閉じる
メニュー
について
メンバー
x-twitter linkedin youtube github discord
Agentgateway が Agentic AI インフラストラクチャのオープン ゲートウェイとして AAIF に参加
AI システムは、分散システムにかなり似てきています。
エージェントはツールを呼び出しています。モデルはプロバイダー間でルーティングされます。ワークフローは、API、MCP サーバー、データベース、その他のエージェントにまたがります。これらのシステムが実験から実際の運用環境に移行するにつれて、組織は、従来の Web トラフィック用に構築されたインフラストラクチャ パターンでは、エージェント システムが引き起こすガバナンス、可観測性、ルーティング、セキュリティの課題に完全には対処していないことにすぐに気づき始めています。
これが、エージェントゲートウェイが解決するために構築された問題です。
コミュニティがこれらの新たなインフラストラクチャの課題に対処できるよう支援するために、agentgateway は Linux Foundation の 4 番目のホスト プロジェクトとして Agentic AI Foundation (AAIF) に参加します。
Agentgateway は、最新の AI システムが実際に動作するように構築されたオープンソース ゲートウェイです。単一の運用面を通じて、MCP トラフィック、エージェント間通信、LLM 推論、REST API、および gRPC サービスを管理するための統合レイヤーを提供します。
AI システム用に個別のインフラストラクチャを立ち上げる代わりに、組織はエージェントゲートウェイを使用して、同じセキュリティ制御、可観測性パイプライン、ルーティング ポリシー、ガバナンス モデルを使用して AI と従来のアプリケーションのトラフィックを一緒に管理できます。
これは、プラットフォーム チームにとって、別個の AI 専用インフラストラクチャ スタックを必要とせずに、使い慣れた運用プリミティブを使用してエージェント トラフィックを管理できることを意味します。セキュリティ、可観測性、ルーティング、ガバナンス、信頼性など、チームがすでに重視しているのと同じ種類の制御をエージェントにも適用できます。

ワークフローも。
MCP および A2A のサポート - 最新のエージェント相互運用プロトコルのためのルーティング、フェデレーション、およびトラフィック管理
モデルの独立性 — LLM プロバイダ間のシームレスな切り替え、オープンウェイト モデルのネイティブな柔軟性
統合データ プレーン - HTTP、gRPC、LLM 推論、エージェント トラフィック用の 1 つのゲートウェイ
セキュリティ制御 - JWT 認証、API キー認証、RBAC、外部認証、mTLS、CORS、悪意のあるツールの動作に対する保護
組み込みの可観測性 — AI およびエージェントのワークフロー向けに設計されたメトリクス、トレース、およびアクセス ロギング
MCP 仮想化 — 複数の MCP ツール サーバーを 1 つのアクセス ポイントに統合し、MCP クライアントのツールレベルのアクセス ポリシーを構成できます。
CEL を使用した宣言型ポリシー – Common Expression Language (CEL) を使用した動的な構成と拡張性
モデルとツールのガバナンス - レート制限、コンテンツベースのルーティング、プロンプト ガード、予算管理、モデルのエイリアシング
プラットフォームに依存しない展開 — ゲートウェイ API に完全に準拠し、ベアメタル、VM、コンテナ、および Kubernetes 上で実行されます。
動的構成 - ダウンタイムなしの xDS 経由の更新
高性能アーキテクチャ — 低遅延、高スループットのワークロード向けに Rust が組み込まれています
AAIF は、Linux Foundation がホストするガバナンスにより、重要なエージェント AI 標準、プロトコル、オープン ソース プロジェクトのためのオープンで中立的なホームを提供します。
これには、MCP などの標準、Goose などのオープンソース ツール、組織が AI システムを安全かつ確実に大規模に運用できるようにするエージェントゲートウェイなどのインフラストラクチャ プロジェクトが含まれます。
Agentgateway は、AI トラフィック、エージェント ワークフロー、相互運用性、セキュリティ、ツールとサービス全体のガバナンスを管理するための共有インフラストラクチャ レイヤーを提供することにより、エコシステムを強化します。コミュニティに居場所を与える

エージェント システムがより広く採用されるにつれて、運用基盤で連携するために必要になります。
「エージェントゲートウェイをAAIFに寄付することは、オープン接続と相互運用性を中心に構築されたプロジェクトにとって自然な次のステップです。私たちの当初からの目標は、オープンインフラストラクチャを通じて組織がAIシステムの運用現実（セキュリティ、ガバナンス、可観測性、信頼性）を管理できるように支援することでした。プロジェクトをAAIFに導入することで、より広範なエコシステムとともにオープンなコラボレーションと中立的なガバナンスを通じてプロジェクトを進化させることができます。」 — Lin Sun 氏、オープンソース責任者、Solo.io、agentgateway 寄稿者
Agentgateway は、エージェント システムの相互運用性、管理性、信頼性を高めることに焦点を当てた、成長を続ける AAIF エコシステムに参加します。
このプロジェクトは Apache 2.0 に基づいてライセンスされており、コミュニティはすでに CoreWeave、Red Hat、Solo.io、Adobe、Salesforce、Amdocs、Microsoft を含む 60 以上の組織にわたる 300 人以上のアクティブな貢献者によって活動されています。
ドキュメントを参照してください:agentgateway.dev
Discord での会話に参加する
毎週のコミュニティミーティングに参加してください
著作権 © The Linux Foundation®。無断転載を禁じます。 Linux Foundation は登録商標および商標を使用しています。利用規約、プライバシー ポリシー、商標の使用などの詳細については、ポリシー ページをご覧ください。

## Original Extract

The Agentic AI Foundation welcomes agentgateway — an open source gateway purpose-built for MCP, Agent-to-Agent, LLM, and API traffic — as its newest hosted project.

Skip to main content
Join us at MCP Dev Summit Bengaluru • June 9-10 • REGISTER NOW
Search
Close Search
Menu
About
Members
x-twitter linkedin youtube github discord
agentgateway Joins AAIF as an Open Gateway for Agentic AI Infrastructure
AI systems are starting to look a lot more like distributed systems.
Agents are calling tools. Models are routing across providers. Workflows are spanning APIs, MCP servers, databases, and other agents. As these systems move from experiments into real operational environments, organizations are quickly realizing that the infrastructure patterns built for traditional web traffic do not fully address the governance, observability, routing, and security challenges agentic systems introduce.
That is the problem agentgateway was built to solve.
To help the community address these emerging infrastructure challenges, agentgateway is joining the Agentic AI Foundation (AAIF) as the fourth hosted project under the Linux Foundation.
agentgateway is an open source gateway built for the way modern AI systems actually work. It provides a unified layer for managing MCP traffic, Agent-to-Agent communication, LLM inference, REST APIs, and gRPC services through a single operational surface.
Instead of standing up separate infrastructure for AI systems, organizations can use agentgateway to manage AI and traditional application traffic together using the same security controls, observability pipelines, routing policies, and governance models.
For platform teams, this means agent traffic can be managed with familiar operational primitives instead of requiring a separate, AI-only infrastructure stack. The same kinds of controls teams already care about — security, observability, routing, governance, and reliability — can apply to agent workflows too.
MCP and A2A support — routing, federation, and traffic management for modern agent interoperability protocols
Model independence — seamless switching between LLM providers, with native flexibility for open-weights models
Unified data plane — one gateway for HTTP, gRPC, LLM inference, and agent traffic
Security controls — JWT authentication, API key auth, RBAC, external authorization, mTLS, CORS, and protections against malicious tool behavior
Built-in observability — metrics, tracing, and access logging designed for AI and agent workflows
MCP virtualization — federate multiple MCP tool servers into a single access point with the ability to configure tool-level access policy for MCP clients.
Declarative policy with CEL – dynamic configuration and extensibility using Common Expression Language (CEL)
Model and tool governance — rate limiting, content-based routing, prompt guards, budget controls, and model aliasing
Platform-agnostic deployment — runs on bare metal, VMs, containers, and Kubernetes, with full Gateway API conformance
Dynamic configuration — updates through xDS without downtime
High-performance architecture — built in Rust for low-latency, high-throughput workloads
AAIF provides an open and neutral home for critical agentic AI standards, protocols, and open source projects, with governance hosted by the Linux Foundation.
That includes standards like MCP, open source tooling like goose, and infrastructure projects like agentgateway that help organizations operate AI systems safely and reliably at scale.
agentgateway strengthens the ecosystem by providing a shared infrastructure layer for managing AI traffic, agent workflows, interoperability, security, and governance across tools and services. It gives the community a place to collaborate on the operational foundations agentic systems will need as they become more widely adopted.
“Donating agentgateway to AAIF is the natural next step for a project built around open connectivity and interoperability. Our goal from day one has been to help organizations manage the operational realities of AI systems — security, governance, observability, and reliability — through open infrastructure. Bringing the project into AAIF ensures it can evolve through open collaboration and neutral governance alongside the broader ecosystem.” — Lin Sun , Head of Open Source, Solo.io & agentgateway Contributor
agentgateway joins a growing AAIF ecosystem focused on making agentic systems more interoperable, manageable, and trustworthy.
The project is licensed under Apache 2.0, and the community is already active with 300+ active contributors across over 60 organizations, including CoreWeave, Red Hat, Solo.io, Adobe, Salesforce, Amdocs, and Microsoft.
Explore the docs: agentgateway.dev
Join the conversation on Discord
Attend our weekly community meetings
Copyright © The Linux Foundation®. All rights reserved. The Linux Foundation has registered trademarks and uses trademarks. For more information, including terms of use, privacy policy, and trademark usage, please see our Policies page.

---
source: "https://aaif.io/blog/a2a-joins-aaif"
hn_url: "https://news.ycombinator.com/item?id=49356130"
title: "Agent2Agent (A2A) joins Agentic AI Foundation (AAIF)'s open agentic stack"
article_title: "A2A joins AAIF’s open agentic stack - Agentic AI Foundation (AAIF)"
image: "https://cdn.sanity.io/images/4o10fa7h/production/e858bdb1e6fb246207a16ec3ddb3219e8e2fcd8a-1600x900.png?rect=0,30,1600,840&w=1200&h=630"
author: "ChrisArchitect"
captured_at: "2026-08-19T03:38:50Z"
capture_tool: "hn-digest"
hn_id: 49356130
score: 1
comments: 0
posted_at: "2026-08-19T03:05:02Z"
tags:
  - hacker-news
  - translated
---

# Agent2Agent (A2A) joins Agentic AI Foundation (AAIF)'s open agentic stack

- HN: [49356130](https://news.ycombinator.com/item?id=49356130)
- Source: [aaif.io](https://aaif.io/blog/a2a-joins-aaif)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T03:05:02Z

## Translation

タイトル: Agent2Agent (A2A) が Agentic AI Foundation (AAIF) のオープン エージェント スタックに参加
記事タイトル: A2A が AAIF のオープン エージェント スタックに参加 - Agentic AI Foundation (AAIF)
説明: エージェント間通信のオープン標準である Agent2Agent (A2A) は、ホスト型プロジェクトとして Agentic AI Foundation (AAIF) に参加します。 A2A は、エージェントがどのようなフレームワークに基づいて構築されたかに関係なく、エージェントがお互いを発見し、タスクを委任し、作業を交換する方法を定義します。 AAIF は、オープンで中立的な拠点です。
[切り捨てられた]

記事本文:
A2A が AAIF のオープン エージェント スタックに参加 - Agentic AI Foundation (AAIF) AGNTCon+MCPCon Europe • 9 月 17 ～ 18 日 • アムステルダム • 今すぐ登録
ブログ > Agent2Agent (A2A) A2A が AAIF のオープン エージェント スタックに参加
A2A がエージェント スタックに適合する場所
インフラストラクチャにとって中立的なガバナンスが重要な理由
A2A はエージェント間通信のオープン スタンダードであり、エージェントがお互いを発見し、タスクを委任し、あらゆるフレームワークやベンダー間で作業する方法です。
サプライ チェーンや金融サービスのエンジニアがマルチエージェント システムを実稼働環境に導入し始めたとき、同じ問題に繰り返し遭遇しました。異なるフレームワーク上に構築されたエージェントは、そのペアリング用に特別に作成されたカスタム統合コードがなければ、相互に作業を引き渡すことができませんでした。新しいベンダーとの関係はすべて、同じ統合作業をゼロから行う必要がありました。コストはエージェント自体ではなく、エージェント間の統合にありました。
A2A は、その問題を直接解決するために構築されました。
A2A は、エージェントがフレームワークやベンダーの境界を越えて通信する方法を定義するオープン スタンダードです。エージェントは、実行できる内容とその到達方法を構造化して説明したエージェント カードを発行します。他のエージェントは、人間がハンドオフを仲介することなく、そのカードを読み取り、機能を発見し、タスクを委任します。交換は構造化されており、観察可能であり、フレームワークに依存しません。
Google は 2025 年 4 月に A2A を立ち上げ、AWS、Cisco、Google、Microsoft、Salesforce、SAP、ServiceNow などの創設組織とともに Linux Foundation に寄付しました。 2025 年 8 月に、IBM のエージェント通信プロトコルが A2A に統合され、この分野が競合する規格ではなく、単一の共通規格に向かって進んでいることが示されました。
最初の安定した仕様である A2A v1.0 は、2026 年 3 月に出荷されました。マルチプロトコル バインディングとバージョン ネゴシエーション、マルチテナントが追加されました。

y、および暗号化 ID 検証用の署名付きエージェント カード。
A2A は、モバイル プラットフォーム、クラウド AI インフラストラクチャ、金融サービス、サプライ チェーン、エンタープライズ IT にわたる実稼働環境ですでに実行されています。
ファーウェイは、Celia、OS レベルの AI アシスタント、HarmonyOS 開発者プラットフォーム全体のアプリ内エージェント間のプロトコルとして A2A を標準化しました。サポートされているシナリオには、Celia が長時間実行タスクをアプリ エージェントに渡すこと、Celia がエージェントを通じてアプリ UI を制御すること、Celia がアプリケーション エージェントからコンテキストに基づく推奨事項を要求することなどが含まれます。これは、OS アシスタントとモバイル プラットフォーム規模で実行されているアプリケーション間のエージェント通信です。
Tencent の WeChat は、A2A 経由で Huawei やその他の Android OEM アシスタントと統合された最初の主要アプリの 1 つであり、AI アシスタントを通じて開始されるメッセージ、音声通話、ビデオ通話を可能にします。フローは、A2A プロトコルを介した二重認証で実行されます。
A2A は、モバイル プラットフォームを超えて、主要なクラウド AI プラットフォームに組み込まれています。 Google Cloud は、ADK、エージェント エンジン、Cloud Run、GKE を介した A2A エージェントの開発とデプロイをサポートしています。 Microsoft Azure AI Foundry を使用すると、エージェントは A2A エンドポイントを公開し、標準の検出を通じて外部エージェントを見つけることができます。 AWS Bedrock AgentCore は A2A サーバーをホストして操作できるため、異なるフレームワークで構築されたエージェントや異なるクラウドでホストされているエージェントは、同じプロトコルを通じて通信できます。
Google Cloud と PayPal は、Agent Payments Protocol（AP2）を通じて A2A をエージェント コマースに拡張しています。ショッピング エージェントと販売エージェントは、製品の発見、価格設定、注文処理を通じて A2A 経由で通信し、AP2 は支払い承認レイヤーを提供します。
A2A がエージェント スタックに適合する場所
AAIF のプロジェクトはそれぞれ、オープン エージェント スタックの異なる層をカバーします。
手順とコンテキスト: AGENTS.md sta

プロジェクトが期待、慣例、操作指示を AI エージェントに伝える方法を標準化します。
エージェント ランタイム: Goose は、エージェントが推論、計画、機能の呼び出し、および作業を実行する環境を提供します。
エージェントからツールへの接続: MCP は、エージェントがツール、データ ソース、アプリケーション、およびサービスに接続し、それらと対話する方法を標準化します。
トラフィックの仲介と制御: エージェントゲートウェイが操作を処理します。エージェント システムとエージェント システムが実行されるインフラストラクチャとの境界に位置し、ルーティング、ポリシー、可観測性を管理します。
エージェント間の相互運用性: A2A は、独立したエージェントが相互に発見し、通信し、タスクを委任し、システムや組織の境界を越えて結果を交換する方法を標準化することにより、エージェント間の相互運用性レイヤーを追加します。
「A2A は、AI エージェントのオープンで相互運用可能な未来に向けた重要な一歩を表しています」と AAIF CTO Manik Surtani 氏は述べています。 「このプロジェクトを Agentic AI Foundation に持ち込むことで、さまざまなベンダー、フレームワーク、組織のエージェントが連携できるようにする標準に関してコミュニティが協力できる中立的な拠点を構築しています。A2A コミュニティを歓迎し、この未来を一緒に構築できることをうれしく思います。」
Google Cloud のビジネス アプリケーション プラットフォーム担当副社長兼ゼネラル マネージャーである Rao Surapaneni 氏は次のように述べています。「組織がマルチプラットフォームのエージェント システムを本番環境に移行するにつれて、相互運用性が不可欠になっています。」 「A2A を使用すると、エージェントはカスタム統合を行うことなく、フレームワークやサイロ化されたプラットフォーム間で簡単に検出、委任、共同作業を行うことができます。A2A を Agentic AI Foundation に導入することで、企業は真にオープンな基盤上でエージェント システムを構築および拡張できるようになります。」
インフラストラクチャにとって中立的なガバナンスが重要な理由
オープンスタンダードの Linux Foundation モデル

sには実績があります。基本コンポーネントが単一のベンダーによって所有されている場合、すべての下流チームはそのベンダーのロードマップの制約とリリース サイクルを吸収します。オープンに管理されると、コミュニティがいつ何を構築するかを決定します。
A2A の 150 以上のパートナー組織には、直接の競合他社も含まれています。その幅は、単一の参加者が制御しないガバナンスの下でのみ維持されます。 Linux Foundation がホストする Agentic AI Foundation がその構造を提供します。これは、コンテキストから通信、オペレーション、オープン エージェント層に至るまでのスタック全体が、同じ方法で同じ場所で管理されることも意味します。
マルチエージェント システムを構築しているチームの場合、これにより特定のカテゴリのリスクが除去されます。エージェントがフレームワーク間通信に依存するプロトコルは、単一ベンダーの製品決定の影響を受けません。
A2A がここに来てうれしいです。私たちと一緒に作りましょう。
A2A 仕様と SDK を確認してください: https://a2a-protocol.org/latest/
Discord での会話に参加してください: https://discord.com/invite/a2aprotocol
A2A の技術運営委員会の会議に参加してください: https://zoom-lfx.platform.linuxfoundation.org/meetings/agent2agent
アイコン Facebook アイコン X アイコン リンク作成者
購読する AAIF ブリーフィングに購読する
標準、ガバナンス、未来を築く人々に関する毎週のお知らせ。毛羽立ちはありません。重要なことだけ。
AAIF について 待ってください。
構築を開始する
テーブルは設定され、仕事は本格的です。あなたの動き
アイコン X アイコン Linkedin アイコン Youtube アイコン Github アイコン Discord WeChat 探索
エージェントの AL 基準とガバナンスに関する情報を毎週入手します。
著作権 © The Linux Foundation®。無断転載を禁じます。 Linux Foundation は登録商標および商標を使用しています。利用規約、プライバシー ポリシー、商標の使用などの詳細については、ポリシー ページをご覧ください。

## Original Extract

Agent2Agent (A2A), the open standard for inter-agent communication, is joining the Agentic AI Foundation (AAIF) as a hosted project. A2A defines how agents discover each other, delegate tasks, and exchange work, regardless of what framework they were built on. AAIF is the open, neutral home where th
[truncated]

A2A joins AAIF’s open agentic stack - Agentic AI Foundation (AAIF) AGNTCon+MCPCon Europe • Sep 17-18 • Amsterdam • REGISTER NOW
Blog > Agent2Agent (A2A) A2A joins AAIF’s open agentic stack
Where A2A fits in the agentic stack
Why neutral governance matters for infrastructure
A2A is the open standard for inter-agent communication and is how agents discover each other, delegate tasks, and work across any framework or vendor.
When engineers in supply chain and financial services started deploying multi-agent systems in production, they encountered the same problem repeatedly. Agents built on different frameworks couldn't hand off work to each other without custom integration code written specifically for that pairing. Every new vendor relationship required the same integration work from scratch. The cost wasn't in the agents themselves but in the integrations between them.
A2A was built to solve that problem directly.
A2A is an open standard that defines how agents communicate across framework and vendor boundaries. An agent publishes an agent card containing a structured description of what it can do and how to reach it. Other agents read that card, discover capabilities, and delegate tasks without a human brokering the handoff. The exchange is structured, observable, and framework agnostic.
Google launched A2A in April 2025 and donated it to the Linux Foundation with founding organizations like AWS, Cisco, Google, Microsoft, Salesforce, SAP, and ServiceNow. In August 2025, IBM's Agent Communication Protocol merged into A2A showing that the field was moving toward a single shared standard rather than competing ones.
A2A v1.0, the first stable specification, shipped in March 2026. It added multi-protocol bindings and version negotiation, multi-tenancy, and signed agent cards for cryptographic identity verification.
A2A is already running in production across mobile platforms, cloud AI infrastructure, financial services, supply chain, and enterprise IT.
Huawei has standardized A2A as the protocol between Celia, its OS-level AI assistant, and in-app agents across the HarmonyOS developer platform. The supported scenarios include Celia handing long running tasks to an app agent, Celia controlling app UI through an agent, and Celia requesting contextual recommendations from an application agent. This is agent communication between an OS assistant and an application running at mobile platform scale.
Tencent's WeChat is among the first major apps integrating with Huawei and other Android OEM assistants over A2A, enabling messages, voice calls, and video calls initiated through an AI assistant. The flow runs with dual authorization through the A2A protocol.
Beyond mobile platforms, A2A is built into the leading cloud AI platforms. Google Cloud supports developing and deploying A2A agents through ADK, Agent Engine, Cloud Run, and GKE. Microsoft Azure AI Foundry lets agents expose A2A endpoints and find external agents through standard discovery. AWS Bedrock AgentCore can host and operate A2A servers, so agents built with different frameworks or hosted in different clouds can communicate through the same protocol.
Google Cloud and PayPal are extending A2A into agentic commerce through the Agent Payments Protocol (AP2). Shopping and merchant agents communicate over A2A throughout product discovery, pricing, and order fulfillment, with AP2 providing the payment authorization layer.
Where A2A fits in the agentic stack
AAIF‘s projects each cover distinct layers of the open agentic stack.
Instructions and context: AGENTS.md standardizes how projects communicate expectations, conventions, and operating instructions to AI agents.
Agent runtime: goose provides the environment in which an agent reasons, plans, invokes capabilities, and carries out work.
Agent to tool connectivity: MCP standardizes how agents connect to and interact with tools, data sources, applications, and services.
Traffic mediation and control: agentgateway handles operations. It sits at the boundary between agent systems and the infrastructure they run on, managing routing, policy, and observability.
Agent to agent interoperability: A2A adds an agent to agent interoperability layer by standardizing how independent agents discover one another, communicate, delegate tasks, and exchange results across systems and organizational boundaries.
“A2A represents an important step toward an open, interoperable future for AI agents,” said AAIF CTO Manik Surtani. “By bringing the project to the Agentic AI Foundation, we're creating a neutral home where the community can collaborate on the standards that will enable agents from different vendors, frameworks, and organizations to work together. We're excited to welcome the A2A community and build this future together.”
“As organizations move multi-platform agentic systems into production, interoperability is becoming essential,” said Rao Surapaneni, VP & General Manager, Business Application Platforms, Google Cloud. “A2A enables agents to easily discover, delegate, and collaborate across frameworks and siloed platforms, without custom integrations. Bringing A2A to the Agentic AI Foundation further empowers enterprises to build and scale agentic systems on a truly open foundation.”
Why neutral governance matters for infrastructure
The Linux Foundation model for open standards has a track record. When foundational components are owned by a single vendor, every downstream team absorbs that vendor's roadmap constraints and release cycles. When they're governed openly, the community shapes what gets built and when.
A2A's 150+ partner organizations include direct competitors. That breadth only holds together under governance that no single participant controls. The Agentic AI Foundation, hosted by the Linux Foundation, provides that structure. It also means the full stack, from context to communication to operations to the open agent layer, is governed in the same way and in the same place.
For teams building multi-agent systems, this removes a specific category of risk. The protocol your agents depend on for cross-framework communication isn't subject to a single vendor's product decisions.
We're glad A2A is here. Come build with us.
Explore the A2A spec and SDKs: https://a2a-protocol.org/latest/
Join the conversation on Discord: https://discord.com/invite/a2aprotocol
Tune in to A2A’s Technical Steering Committee Meetings: https://zoom-lfx.platform.linuxfoundation.org/meetings/agent2agent
Icon Facebook Icon X Icon Link Author
Subscribe Subscribe to the AAIF Briefing
Weekly signal on standards, governance, and the people building the future. No fluff. Just what matters.
About AAIF Stop waiting,
Start building
The table is set, The work is real. Your move
Icon X Icon Linkedin Icon Youtube Icon Github Icon Discord WeChat Explore
Get weekly intelligence on agentic Al standards and governance.
Copyright © The Linux Foundation®. All rights reserved. The Linux Foundation has registered trademarks and uses trademarks. For more information, including terms of use, privacy policy, and trademark usage, please see our Policies page.

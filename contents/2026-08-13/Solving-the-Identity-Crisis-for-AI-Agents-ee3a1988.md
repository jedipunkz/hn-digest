---
source: "https://www.uber.com/us/en/blog/solving-the-agent-identity-crisis/"
hn_url: "https://news.ycombinator.com/item?id=49288326"
title: "Solving the Identity Crisis for AI Agents"
article_title: "Solving the Identity Crisis for AI Agents"
author: "mooreds"
captured_at: "2026-08-13T16:45:13Z"
capture_tool: "hn-digest"
hn_id: 49288326
score: 3
comments: 0
posted_at: "2026-08-13T16:24:28Z"
tags:
  - hacker-news
  - translated
---

# Solving the Identity Crisis for AI Agents

- HN: [49288326](https://news.ycombinator.com/item?id=49288326)
- Source: [www.uber.com](https://www.uber.com/us/en/blog/solving-the-agent-identity-crisis/)
- Score: 3
- Comments: 0
- Posted: 2026-08-13T16:24:28Z

## Translation

タイトル: AI エージェントのアイデンティティ危機を解決する
説明: AI エージェントは強力なアクションを実行できますが、誰が、なぜアクションを実行したかを証明できますか? Uber が、安全で監査可能なツール呼び出しのためにエージェント ID とエンドツーエンドの出所をどのように構築したかを説明します。

記事本文:
メインコンテンツにスキップ
×
ブラウザはサポートされていません
これはおそらくあなたが期待していた経験ではありません。 Internet Explorer は Uber.com ではサポートされていません。別のブラウザに切り替えてサイトを表示してみてください。
製品の広告を調べる Uber での広告について詳しく学びます。消費者がどこに行っても、何でも手に入るようにリーチします。
Uber での運転と配達に関するリソース
移動中の人々のための体験と情報
Uber Eats では、食事の配達を注文するのは始まりにすぎません
世界中の顧客が手の届くところに店舗を置く
企業が従業員を動かし、食事を与える方法を変革する
医療従事者とともに医療を前進させる
キャンパス交通の充実
公共交通機関の利用範囲を拡大する
Uber Engineering を支えるテクノロジー
世界中の都市とコミュニティのために正しいことを行う
あなたの国における Uber のニュースと最新情報
製品、ハウツー、ポリシーのコンテンツなど
AI エージェントのアイデンティティ危機を解決する
この記事を共有する Facebook Linkedin X ソーシャル リンクの紹介
Uber は AI 活用の最前線に立っており、エンジニアが AI ソリューションを構築して生産性を向上できるように支援しています。 2025 年の初めに、同社はチームが運用グレードのエージェントを大規模に構成、展開、運用できるようにする社内エージェント プラットフォームを構築しました。さらに、数千のサービスで構成される Uber のマイクロサービス技術スタックは、既存のサービス API に対する MCP® (Model Context Protocol) サポートを有効にすることで AI 対応になりました。
エージェントの自律性を高めるには、エージェントとエージェントが実行するアクションを厳密に監視する必要があります。説明責任、つまり「誰が、いつ、なぜ行ったのか」に答える能力は、監査、コンプライアンス、および経営陣の信頼にとって重要です。明確な帰属がないと、セキュリティ管理の実施が難しくなり、インシデントへの対応が遅くなり、信頼性が低下する可能性があります。

影響を受ける可能性があります。
このブログでは、AI エージェントに対応するための 2025 年の Uber の ID およびアクセス テクノロジー スタックの主要なアップデートについて概要を説明します。 AI 導入が加速する中、積極的な姿勢を維持するために、この技術分野における 2026 年の戦略的ロードマップも垣間見ることができます。
説明されているシステムとアプローチは、Uber の内部アーキテクチャと制御された生産環境を反映しています。設計の選択、パフォーマンス特性、セキュリティ制御は、組織、ユースケース、導入コンテキストによって異なる場合があります。
オンコール エンジニアがオンコール エージェントを使用してシステム アラートを管理および解決していると想像してください。このシナリオでは、調査エージェントは、システムが正しく機能しており、アラート自体が正しく構成されていないと判断しました。その後、調査エージェントはそのタスクをモニタリング エージェントにシームレスに渡し、PR (プル リクエスト) を通じてアラートのしきい値を調整しました。プル リクエストには、監視エージェントが変更を導入していることが示されていますが、責任を負うオンコール エンジニアの身元は追跡できないままです。
エージェント ワークフローが拡大して、より多くのエージェント、ツール、システムが含まれるようになるにつれて、この課題はますます顕著になります。私たちはこれを次の 2 つの主要な問題に抽出しました。
今日の ID モデルは、人間とワークロード (多くの場合、非人間 ID (NHI) と呼ばれ、サービス アカウントや API キーなどの資格情報によってサポートされます) を中心に構築されています。エージェントは、他の人のために、またはその代わりに行動する権限を与えられたエンティティとして定義するのが最も適切です。 AI エージェントは多くの場合、人間に代わってタスクを実行するワークロードとして実行されます。上記の例では、オンコール エージェントがオンコール エンジニアに代わってセッションを開始し、特定の問題を調査して修正しました。
実行コンテキスト (発信元ユーザー、中間エージェント) は、エージェント ホップ全体でドロップされます。これは不完全な状態につながります

システム全体の監査が行われ、ダウンストリーム システムによってすでに構成されている詳細なアクセス ポリシーを一貫して活用する能力が制限されます。完全な監査証跡がない場合、インシデント対応ではシステム全体の部分的な監査ログをつなぎ合わせる必要があります。モニタリング エージェントによって開かれた PR には、オンコール エンジニアが特定の問題の解決を要求したことと、PR に至った以前のエージェントの決定に関するコンテキストが示されている必要があります。
エージェント ワークフローが従来の自動化とは異なる動作をすることは明らかです。
委任はデフォルトのモードです - エージェントは他の人に代わって動作します
ワークフローは構成的です - エージェントは他のエージェント、ツール、システムを呼び出します
動作は動的です - セッションが進行するにつれて、中間結果に基づいて計画が進化します
これにより、私たちが構築しなければならないものの方向性が決まりました。つまり、エージェント ID の基盤と、上記の問題に対処するエージェント間でのその伝達です。
AI ワークフローが拡大するにつれて、自律エージェントと内部システムの間の相互作用は非常に複雑になります。開発者の速度を阻害することなくこのエコシステムを保護するために、私たちは既存のゼロトラスト アーキテクチャを AI エージェント向けに拡張することにしました。私たちのアーキテクチャは、エージェント エコシステム内で検証可能な暗号化 ID を確立し、ダウンストリーム システムにアクセスするための承認を強制することに重点を置いています。
Uber では、AI エージェントがワークロードとして導入されることが多く、多くの場合 Kubernetes® によって管理されます。 Michelangelo プラットフォームは、AI エージェントをワークロードに関連付けます。エージェント レジストリは、この登録を保存する信頼できる情報源として機能します。これは、後でセキュリティ トークン サービスによってエージェントを検証するために使用されます。
一般的な用語「サービス メッシュ」に似た AI エージェント メッシュは、AI エージェントが互いに通信して、割り当てられたタスクを完了するデータ プレーンです。エージェント内で

メッシュおよびアウトバウンド呼び出し (MCP ツールなど) の場合、AI エージェントは認証のためにセキュリティ トークン サービスによって作成された JWT トークンに依存します。
AI エージェントのトークン生成は STS によって処理されます。 STS は、広範で有効期間の長いサービス資格情報に依存するのではなく、ホップごとに有効期間が短く、範囲が限定されたトークンを発行する動的な信頼ブローカーとして機能します。
MCP ゲートウェイは、AI エージェント メッシュから Uber のシステムへの通話を仲介する中央システムです。この設計により、MCP ゲートウェイを MCP ツール呼び出しのポリシー適用ポイントにすることができます。
MCP ゲートウェイが呼び出し元の認証に成功し、ツール呼び出しを許可すると、リクエストをそれぞれのダウンストリーム サービスに安全にプロキシ送信します。これらは主に、実際のミューテーションやデータ取得を実行するマイクロサービス API とデータストアです。
これらのコンポーネントを超えて、AI ゲートウェイは、AI エージェントから AI モデルへのアウトバウンドのすべての呼び出しを仲介します。これは、Uber と OpenAI®、Anthropic® などの外部 API との統合の中心点として機能します。 AI ゲートウェイはセキュリティ ガードレールと統合されており、プロンプト インジェクション、ジェイルブレイク、コンテンツの安全性、PII 編集などを検出して処理します。 Uber の AI Guard について詳しくは、最近のカンファレンスでのプレゼンテーションをご覧ください。
エンジニアと運用チームがエージェント ソリューションを構築できるようにするために、Michelangelo AI プラットフォームには 2 つのオプションが用意されています。
コード: Uber の社内運用 SDK を使用して、Python でエージェントを作成します。 SDK はオーケストレーション フレームワークに依存せず、一般的なエージェント プログラミング パターン (計画ループ、ツールの使用、状態、メモリ) をサポートすると同時に、標準化されたスキャフォールディング、ミドルウェア フック、可観測性、および運用環境の展開用の評価ツールを提供します。
コードなし: コードを書かずに、UI を通じてエージェントを作成します。これにより参入障壁が低くなり、可能性が広がります

エンジニアを超えて会社全体にエージェントを構築する能力。
オプションに関係なく、結果として得られる AI エージェントは Uber の Kubernetes インフラストラクチャ内にデプロイされます。
当初、AI エージェント間の通話をプロキシできるエージェントゲートウェイの構築/採用を検討しました。 Uber のエージェント AI エコシステムは SDK を中心に大幅に標準化されていたため、代わりにソリューションを SDK に直接統合しました。また、問題 2 に完全に対処するには、外部プロキシのみに依存するのではなく、実行コンテキストが作成され、エンドツーエンドで伝播されるエージェント アプリケーション層でのサポートが必要であることもわかりました。
マイクロサービスと同様に、AI エージェントはワークロード内で実行されます。対処すべき基本的な課題は、個々のエージェントに検証可能な ID を割り当てる方法でした。図 3 は、エージェント ID モデルとエージェントの JWT トークンを生成するプロセスを示しています。
図 3: エージェントに ID を提供する。
すべてのワークロードは、まず SPIRE から暗号化署名された独自のワークロード SVID (SPIFFE 検証可能 ID) を取得します。これにより、基盤となるコンピューティング環境の正当性が証明されますが、エージェントはまだ特定されていません。
SDK は、ローカルで利用可能なメタデータ (エージェント構成など)、受信呼び出しおよび送信宛先ユーザーからの JWT を使用して、ワークロード SVID で認証された STS からの新しい JWT トークンを要求します。 STS のみが AI エージェントのトークンを鋳造することを許可されています。このプロセスを一元化することで、アクター チェーンがリクエストに関与するすべてのエンティティの暗号化記録を確実に保持できるようにします。
STS はエージェント レジストリと統合して、要求元のagent_id がその特定のワークロードでの実行を明示的に許可されていることを確認します (ステップ 1 から)。これにより、ワークロードがホストする権限を持たないエージェントになりすますことが防止されます。
STS は JWT トークンを作成し、それをリクエストに返します。

おとり捜査官。この JWT は、エージェント フローのネクスト ホップのリクエストに使用されます。
この設計の主な特徴は次のとおりです。
シングルホップ、有効期間の短いトークン。 STS によって作成されたすべての JWT は、特定の Audience クレームと数分程度の短い存続時間を備えたシングルホップを対象としています。エージェント A がエージェント B を呼び出すために発行されたトークンは、インターセプトして再実行してデータベースや別のサービスを呼び出すことはできません。その特定の宛先にのみ有効です。
完全な文脈上の帰属。 STS はあらゆるステップでトークン交換を管理し、完全に証明されたアクター チェーンをトークンに埋め込みます。これにより、MCP ゲートウェイまたはダウンストリーム システムがリクエストの完全なコンテキストを取得できるようになります。私たちは、直接の発信者だけではなく、リネージ内のすべての参加者 (例: エンジニア、オンコール エージェント、調査エージェントなど) を確認します。この可視性により、包括的な監査ログと、完全なリクエスト系統を考慮した高度なワークフロー承認が可能になります。
拡張可能なコンテキスト。 JWT 構造は拡張可能に設計されています。将来的には、セッション ID やリクエスト インテント関連のクレームなどのクレームをシームレスに追加して、ポリシー決定のためのより豊富なコンテキストを提供できます。この忠実度の高い可視性により、最後のホップだけでなくチェーン全体の検証された意図によってツールの実行を承認できることが保証されます。
すべてのエージェント ID を SPIRE に裏付けられたワークロード認証情報に固定し、トークン交換を一元化することで、エンドツーエンドのトレーサビリティを維持しながら、有効期間の短いトークンを提供できます。
エージェント ID が実際のワークフローでどのように現れるかを理解するために、一般的なリクエスト パスをたどってみましょう。エージェント AI ワークフローには、複雑なユーザー リクエストを満たすために複数の専門エージェントを呼び出すことが含まれるため、ID はそのアイデンティティを失うことなく、あらゆる境界で進化する必要があります。

元の文脈。図 4 は、最初のユーザー クエリから最終的な安全なツールの呼び出しまでのマルチホップ調査フローを示しています。
図 4: Agentic AI セッションのライフサイクル。
オンコール エンジニア (user1) は、オンコール エージェントとのセッションを開始します。このエントリ ポイントでは、リクエストはユーザー自身の個人的なアイデンティティによって固定されています。
Oncall Agent は、ユーザーの生の資格情報を再利用してダウンストリーム サービスを呼び出すことはできません。代わりに、セキュリティ トークン サービスに接続します。 SPIRE が発行した ID (ワークロード-1) とユーザーのコンテキストを提示して、特に調査エージェントとしてネクストホップの対象ユーザーを対象とした新しい JWT を要求します。 STS は、JWT を使用してオンコール エージェントに応答します。トークンを交換するためのこのホップごとのメカニズムは、概念的には OAuth 2.0 Token Exchange ( RFC 8693 ) に基づいていますが、Uber の内部監査およびパフォーマンス要件と統合する合理的な方法でエージェントの ID と来歴を送信するようにカスタマイズされています。
図 5: Workload-2 を呼び出そうとしているオンコール エージェントの JWT。
オンコール エージェントは、上記の JWT を調査エージェント (ワークロード 2 内でホストされている) に送信します。
調査エージェントは署名と聴衆を検証します。 MCP ゲートウェイを呼び出すために、調査エージェントは MCP ゲートウェイとして STS 対象者と独自のトークン交換を実行します。これは

[切り捨てられた]

## Original Extract

AI agents can take powerful actions, but can you prove who acted and why? Here’s how Uber built agent identity and end-to-end provenance for secure, auditable tool calls.

Skip to main content
×
Browser not supported
This probably isn't the experience you were expecting. Internet Explorer isn't supported on Uber.com. Try switching to a different browser to view our site.
Explore Products Advertising Learn more about advertising on Uber. Reach consumers as they go anywhere and get anything.
Resources for driving and delivering with Uber
Experiences and information for people on the move
Ordering meals for delivery is just the beginning with Uber Eats
Putting stores within reach of a world of customers
Transforming the way companies move and feed their people
Moving care forward together with medical providers
Enhancing campus transportation
Expanding the reach of public transportation
The technology behind Uber Engineering
Doing the right thing for cities and communities globally
Uber news and updates in your country
Product, how-to, and policy content—and more
Solving the Identity Crisis for AI Agents
Share this article Facebook Linkedin X social Link Introduction
Uber is at the forefront of leveraging AI, empowering engineers to build AI solutions to improve productivity. In early 2025, the company built an internal Agent platform that allows teams to compose, deploy, and operate production-grade agents at scale. Additionally, Uber’s microservices tech stack comprising thousands of services was made AI-ready by enabling MCP® (Model Context Protocol) support over existing service APIs.
Increasing agentic autonomy necessitates strict oversight of the agents and the actions they execute. Accountability, the ability to answer “who did what, when and why” is critical for auditing, compliance, and executive trust. Without clear attribution, security controls can be harder to enforce, incident response may slow, and trust may be impacted.
This blog outlines the major updates to Uber’s identity and access technology stack in 2025 to accommodate AI agents. To maintain a proactive stance as AI adoption accelerates, we also offer a glimpse into our strategic roadmap for 2026 within this technical area.
The systems and approaches described reflect Uber’s internal architecture and controlled production environments. Design choices, performance characteristics, and security controls may vary across organizations, use cases, and deployment contexts.
Imagine an on-call engineer using an Oncall Agent to manage and resolve a system alert. In this scenario, the Investigation Agent determined the system was functioning correctly and the alert itself was misconfigured. The Investigation Agent then seamlessly passed the task to the Monitoring Agent to adjust the alert's threshold through a PR (pull request). The pull request shows a Monitoring Agent introducing the change, but the identity of the on-call engineer responsible remains untraceable.
As agentic workflows expand to encompass more agents, tools, and systems, this challenge becomes increasingly pronounced. We distilled this into the following two core problems.
Today’s identity models are built around humans and workloads (often called non-human identity, or NHI, supported through credentials such as service account or API keys). An agent is best defined as an entity that is authorized to act for or in the place of another. AI agents often run as workloads performing tasks on behalf of a human. In the above example, the Oncall Agent started a session on behalf of the on-call engineer to investigate and fix a specific issue.
Execution context (originating user, intermediate agents) is dropped across agent hops. This leads to incomplete audits across the system and limits our ability to consistently leverage the fine-grained access policies already configured by downstream systems. In the absence of complete audit trails, incident response would require stitching partial audit logs across systems together. The PR opened by the Monitoring Agent should indicate that the on-call engineer requested solving a specific issue and some context around prior agent decisions that led to the PR.
It’s clear that agentic workflows behave differently than traditional automation:
Delegation is the default mode - agents work on behalf of others
Workflows are compositional - agents call other agents, tools, and systems
Behavior is dynamic - plans evolve based on intermediate results as a session progresses
This defined the direction for what we had to build: foundations for agent identity and its propagation across agents that address the above problems.
As AI workflows scale, the interactions between autonomous agents and internal systems become deeply complex. To secure this ecosystem without stifling developer velocity, we decided to extend our existing Zero Trust Architecture for AI agents. Our architecture focuses on establishing verifiable cryptographic identity within the agent ecosystem and enforcing authorization for accessing downstream systems.
At Uber, AI agents are often deployed as workloads, often managed by Kubernetes®. The Michelangelo platform associates an AI agent to a workload. The Agent Registry serves as the source of truth, storing this registration. This is later used by the Security Token Service to verify the agent.
Analogous to the popular term service mesh , the AI Agent Mesh is the data plane where AI agents communicate with each other to complete tasks assigned to them. Within the Agent Mesh and for outbound calls (such as to MCP tools), AI agents rely on JWT tokens minted by the Security Token Service for authentication.
Token minting for AI agents is handled by STS. Rather than relying on broad, long-lived service credentials, the STS acts as a dynamic trust broker that issues short-lived, scoped tokens for every hop.
MCP Gateway is a central system that mediates calls from the AI Agent Mesh to Uber’s systems. This design enables MCP Gateway to be a policy enforcement point for MCP tool invocations.
Once the MCP Gateway successfully authenticates the caller and authorizes the tool call, it securely proxies the request to the respective downstream services. These are primarily microservice APIs and datastores that execute the actual mutation or data retrieval.
Beyond these components, an AI Gateway mediates all calls outbound from AI agents to AI models. This serves as the central point of integration for Uber with external APIs such as OpenAI®, Anthropic®, and others. The AI Gateway is integrated with security guardrails to detect and handle prompt injection, jailbreaks, content safety, PII redaction, and more. Learn more about Uber’s AI Guard from our recent conference presentation here .
To empower engineers and operational teams to build agentic solutions, the Michelangelo AI platform provides two options:
Code: Write agents in Python using Uber’s internal production SDK. The SDK is orchestration-framework agnostic and supports common agent programming patterns (planning loops, tool use, state and memory), while providing standardized scaffolding, middleware hooks, observability, and evaluation tooling for production deployments.
No-code: Author agents through the UI without writing any code. This lowers the barrier to entry and opens up the ability to build agents to the entire company beyond engineers.
Regardless of the options, the resulting AI agent gets deployed within Uber’s Kubernetes infrastructure.
Initially we considered building/adopting agentgateway that can proxy calls between AI agents. As Uber’s agentic AI ecosystem standardized heavily around the SDK, we instead integrated the solution directly into the SDK. We also found that fully addressing Problem 2 required support in the agent application layer, where execution context is created and propagated end-to-end, rather than relying only on an external proxy.
Similar to microservices, AI agents run within workloads. The fundamental challenge to address was how to assign each individual agent a verifiable identity. Figure 3 shows our agent identity model and the process to mint a JWT token for the agent:
Figure 3: Providing an agent it’s identity.
Every workload first fetches its own cryptographically signed workload SVID (SPIFFE Verifiable ID) from SPIRE . This proves the legitimacy of the underlying compute environment but doesn’t yet identify the agent.
The SDK uses its metadata available locally (like agent config), JWT from inbound calls and outbound destination audience to request a new JWT token from STS authenticated with the workload SVID. Only the STS is permitted to mint tokens for AI agents. By centralizing this process, we ensure that the actor chain carries the cryptographic record of every entity involved in the request.
STS integrates with the Agent Registry to verify that the requesting agent_id is explicitly authorized to run on that specific workload (from step 1). This prevents a workload from attempting to impersonate an agent that it isn’t authorized to host.
STS mints a JWT token and returns it to the requesting agent. This JWT is used for requests for the next hop of the agentic flow.
Here are some key features of this design:
Single-hop, short-lived tokens. Every JWT minted by the STS is intended for a single hop, with a specific Audience claim and a short time-to-live in the order of minutes. A token issued for Agent A to call Agent B can’t be intercepted and replayed to call a database or another service; it’s valid only for that specific destination.
Full contextual attribution. STS manages the token exchange at every step and embeds the fully attested actor chain into the token. This allows the MCP Gateway or downstream system to have the full context of the request; we see every participant in the lineage (e.g. engineer to Oncall Agent to Investigation Agent …) rather than just the immediate caller. This visibility allows for comprehensive audit logs and advanced workflow authorization that accounts for the full request lineage.
Extensible context. JWT structure is designed to be extensible; we can seamlessly add additional claims in the future, such as session identifiers and request intent related claims, to provide richer context for policy decisions. This high-fidelity visibility ensures that a tool's execution can be authorized not just by the last hop, but by the verified intent of the entire chain.
By anchoring every agent identity in a SPIRE-backed workload credential and centralizing token exchange, we’re able to provide short-lived tokens while maintaining end-to-end traceability.
To understand how agent identity manifests in a real-world workflow, let’s trace a typical request path. As agentic AI workflows involve calling multiple specialized agents to fulfill a complex user request, the identity must evolve at every boundary without losing its original context. Figure 4 shows a multi-hop investigation flow, from an initial user query to the final secure tool invocation:
Figure 4: Agentic AI session life cycle.
An on-call engineer (user1) initiates a session with the Oncall Agent. At this entry point, the request is anchored by the user’s own personnel identity.
The Oncall Agent can’t reuse the user’s raw credentials to call downstream services. Instead, it contacts the Security Token Service. It presents its SPIRE-issued identity (Workload-1) and the user’s context to request a new JWT specifically scoped for the next-hop audience as Investigation Agent. STS responds with a JWT to the Oncall Agent. This per-hop mechanism for exchanging tokens is conceptually based on OAuth 2.0 Token Exchange ( RFC 8693 ) but is customized to transmit agent identity and provenance in a streamlined way that integrates with Uber's internal auditing and performance requirements.
Figure 5: JWT for oncall agent about to call Workload-2.
The Oncall Agent sends the above JWT to the Investigation Agent (hosted within Workload-2).
The Investigation Agent verifies the signature and the audience. To call MCP Gateway, Investigation Agent performs its own token exchange with STS audience as MCP Gateway. This s

[truncated]

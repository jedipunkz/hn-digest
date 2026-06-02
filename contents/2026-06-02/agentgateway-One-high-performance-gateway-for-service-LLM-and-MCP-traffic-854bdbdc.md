---
source: "https://agentgateway.dev/"
hn_url: "https://news.ycombinator.com/item?id=48363702"
title: "agentgateway – One high-performance gateway for service, LLM, and MCP traffic"
article_title: "agentgateway | Agent Connectivity Solved"
author: "wicket"
captured_at: "2026-06-02T00:30:20Z"
capture_tool: "hn-digest"
hn_id: 48363702
score: 1
comments: 0
posted_at: "2026-06-01T22:59:02Z"
tags:
  - hacker-news
  - translated
---

# agentgateway – One high-performance gateway for service, LLM, and MCP traffic

- HN: [48363702](https://news.ycombinator.com/item?id=48363702)
- Source: [agentgateway.dev](https://agentgateway.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T22:59:02Z

## Translation

タイトル: Agentgateway – サービス、LLM、および MCP トラフィック用の 1 つの高性能ゲートウェイ
記事のタイトル: エージェントゲートウェイ |エージェント接続が解決されました
説明: AI トラフィック用のオープンソース ゲートウェイ — LLM、MCP、A2A、HTTP を 1 つのデータ プレーンに統合します。

記事本文:
エージェントゲートウェイ |エージェントの接続が解決されました コンテンツにスキップ ドキュメント
スタンドアロン
Kubernetes チュートリアル
スタンドアロン
Kubernetes モデル
ブログ
エンタープライズ
コミュニティ
ライト
ドキュメント
スタンドアロン
Kubernetes
チュートリアル
スタンドアロン
Kubernetes
モデル
ブログ
エンタープライズ
コミュニティで始める
GitHub
サービス、LLM、および MCP トラフィック用の 1 つの高性能ゲートウェイ。
従来のアプリケーション トラフィックと AI ネイティブ プロトコルを 1 つのデータ プレーンで処理する、オープンソースの HTTP および gRPC ゲートウェイ。別々のゲートウェイをつなぎ合わせることなく、サービス、LLM プロバイダーのトラフィック、MCP ツール、エージェント間の通信をルーティング、保護、監視、管理します。
実際の動作を確認してください
GitHub で見る
Discord LLM ゲートウェイ
OpenAI、Claude、Gemini、または認証とフェイルオーバーを備えたセルフホスト モデルにルーティングします。
モデル コンテキスト プロトコルとエージェント間のトラフィックのネイティブ プロトコル サポート。
HTTP、gRPC、および TCP トラフィック - mTLS および OIDC が組み込まれているため、デフォルトで監視可能です。
CLI / IDE Claude、Codex、OpenCode AI エージェント LangSmith、CrewAI、ADK アプリおよびクライアント AI エージェント OpenAI Anthropic MCP / A2A API / サービス トラフィック管理 LLM ゲートウェイ推論ルーティング MCP / A2A ゲートウェイ リクエスト · インバウンド ルート · セキュア · オブザーブ · ガバナンス レスポンス · アウトバウンド 貢献企業 Linux Foundation Solo.io Microsoft Apple Alibaba Adobe AWS Cisco Salesforce Huawei Amdocs Linux Foundation Solo.io Microsoft Apple Alibaba Adobe AWS Cisco Salesforce Huawei Amdocs エージェント トラフィックに必要なものすべてが 1 つのバイナリにあります。
サービス トラフィック、LLM ルーティング、MCP、A2A、推論用の 1 つのコントロール プレーン。プラットフォーム チームが実際に信頼するポリシーと可観測性を備えています。
サービス ゲートウェイ HTTP、gRPC、TCP — デフォルトで監視可能。単一のバイナリを通じて南北または東西のトラフィックを実行します。ゼロ構成 TLS、OIDC、mTLS ローテーション、ネイティブ Envoy

それを望む人のための互換性。
LLM ゲートウェイ 12 モデル以上のバックエンドにわたるルート、フェイルオーバー、バジェット。 OpenAI、Anthropic、Bedrock、Gemini、Vertex、Cohere、OSS Llama の前に OpenAI 互換 API をドロップインし、チームごとのトークン予算、セマンティック キャッシュ、即時編集を行います。
推論ルーティング セルフホスト型 GPU プール間のスマート ルーティング。セルフホスト型推論のための、レイテンシー、コスト、モデルを認識したルーティング。プラグイン vLLM、TGI、Triton; Agentgateway は最もウォームなレプリカを選択します。
エージェントが行うすべてのツール呼び出しを検出、署名、スコープし、観察します。 MCP サーバーをマイクロサービスのように扱い、バージョニング、RBAC、監査証跡をすぐに使用できます。
Agentgateway は、エージェント間プロトコルをネイティブに話します。 ID、トレース、リプレイを使用して、LangChain、CrewAI、ADK と独自のランタイムの間で呼び出しをルーティングします。
セキュリティと可観測性 通話ごとのトレース、ログ、およびポリシー決定。デフォルトでは OpenTelemetry。ツールごと、エージェントごと、テナントごとのカウンター。 OPA によって評価されたすべてのホップの許可/拒否と、改ざんの明らかな監査ログ。
クイックスタート ローカルファーストで 5 分以内に Agentgateway を使い始めることができます。
組み込みの検出と認証を使用して、エージェントゲートウェイを MCP サーバーに接続します。
デプロイメント オプション バイナリ、Docker、または Kubernetes — スタックに適合するものを選択してください。
エージェントに興味のあるすべてのツール ビルダー、プラットフォーム エンジニア、AI 愛好家を呼び寄せて、エージェント接続の未来の形成に協力してください。
不和
GitHub
「未来は、スタンドアロンのエージェント、MCP サーバー、または LLM によって構築されるものではありません。未来は、それらの相互接続と、シームレスに連携する能力によって形作られます。それらの可能性を最大限に引き出すには、ポリシーを適用し、制御を確保し、それらのやり取りに対する明確な可視性を維持する必要があります。ここで、エージェントゲートウェイが極めて重要な役割を果たし、エージェント間 (A2A) 通信だけでなく橋渡しをすることになります。

だけでなく、エージェントから MCP サーバーまで、エコシステム内の重大なギャップを埋めます。 Linux Foundation 内でプロジェクトが継続的に推進されることを楽しみにしています。」
「AI エージェントは、企業の働き方と革新の方法を急速に変革しています。これらを責任を持って大規模に導入するには、ガバナンス、可視性、セキュリティを提供するオープンで相互運用可能なゲートウェイが組織に必要です。 Agentgateway は、企業が必要とする基盤を提供します。 CoreWeave のようなクラウド プラットフォームで AI を拡張するためにお客様が必要とするオープン基盤を加速するためにコミュニティが団結するのを見ることができて興奮しています。」
「Agentic AI には、単なるソフトウェア層ではなく、専用のインフラストラクチャが必要です。これには、コンピューティング、ストレージ、データの移動を根本から再考する必要があるため、従来のシステムを改修しても機能しません。エージェントゲートウェイ プロジェクトは、その未来に向けた確かな一歩です。 Linux Foundation がホストしており、オープンソースとコミュニティが AI インフラストラクチャに必要な導入と寿命を促進できることを嬉しく思います。」
「私たちは、エージェントゲートウェイ プロジェクトを Linux Foundation に歓迎し、エージェント ワークフローのベスト プラクティスが今後も無料で誰にでも公開されることを保証することに興奮しています。 Agentgateway は、A2A や MCP などの新しい仕様を補完し、エージェント通信用のスケーラブルで仕様に合わせたインフラストラクチャを提供することで、顧客と開発者がプラットフォーム全体で堅牢なエージェント ワークフローを構築できるようにします。」
「Solo.io がエージェントゲートウェイ プロジェクトを Linux Foundation に寄付することに興奮しています。 MCP と A2A を理解する AI ネイティブの接続ソリューションをオープンソース化することは、コミュニティにとって大きな勝利です。これにより、現実世界で必要なセキュリティ、柔軟性、ガバナンスを備えた AI を責任を持って拡張できるようになります。」
「信頼できる AI エージェントを構築することは、特にすべてのステップに非依存性が関与する場合には困難です。

-LLM、ツール、自律エージェントへの決定的呼び出し。 Agentgateway と OpenTelemetry の統合により、可観測性のための堅牢な基盤が提供され、各要求と応答のペアを評価可能な単位として扱うことができます。この機能は、システムレベルの精度と信頼性を確保するために非常に重要であり、チームが AI ワークフローを拡張、保護、最適化できる真の「AI メッシュ」への道を切り開きます。」
「今日の最大の未解決のセキュリティ問題の 1 つは、MCP セキュリティを効果的に実行する方法です。この分野にはコミュニティが対処方法を知らない問題がたくさんありますが、エージェントゲートウェイ プロジェクトは、基本的なロールベースのアクセス制御と MCP サーバーへのアクションの可視性に関する重要な問題のいくつかに対処するための第一歩を提供します。オープンソース ガバナンスの下で、この分野で複雑に進化する脅威に対処するために、このプロジェクトがどのように適応し、進化するかを見るのを楽しみにしています。」
「AI 環境の急速な進化には、エージェントが相互に通信したり、外部ツールと通信したりするための、堅牢でベンダー中立的なインフラストラクチャが必要です。それがなければ、イノベーションと採用が抑制される危険があります。 Linux Foundation が主催する Agentgateway プロジェクトは、その共通基盤を構築するための重要なステップです。私たちは Solo.io と提携し、相互運用可能な AI の将来に向けたコミュニティ主導の基盤をサポートできることを嬉しく思っています。」
「今日の最大の未解決のセキュリティ問題の 1 つは、MCP セキュリティを効果的に実行する方法です。この分野にはコミュニティが対処方法を知らない問題がたくさんありますが、agentgateway プロジェクトは、基本的なロールベースのアクセス制御と MCP サーバーへのアクションの可視性に関する重要な問題のいくつかに対処するための第一歩を提供します。このプロジェクトが複雑な問題に対処するためにどのように適応し、進化するかを見るのを楽しみにしています。

オープンソースのガバナンスの下で、この分野で進化する脅威。」
従来のアプリケーション トラフィックと AI ネイティブ プロトコルを 1 つのデータ プレーンで処理する、オープンソースの HTTP および gRPC ゲートウェイ。別々のゲートウェイをつなぎ合わせることなく、サービス、LLM プロバイダーのトラフィック、MCP ツール、エージェント間の通信をルーティング、保護、監視、管理します。
AIに聞く
エージェントゲートウェイアシスタント
0" class="flex items-center justify-center w-8 h-8 bg-transparent border-nonerounded-lg text-gray-500cursor-pointertransition-allduration-150 hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-700 dark:hover:text-gray-50" aria-label="会話をエクスポート" title="エクスポートマークダウンとして">
エージェントゲートウェイの構成、機能、使用法について何でも聞いてください。
注: AI によって生成されたコンテンツにはエラーが含まれる可能性があります。返されたすべての情報を確認してテストしてください。
ヒント: 会話ごとに 1 つのトピックを使用すると、最良の結果が得られます。新しい会話を開始するには、チャット ヘッダーの + ボタンを使用します。
•
•
• レート制限に達しました
アシスタントは 3 回の交換のローリング履歴を保持します。古いメッセージはコンテキストに含まれなくなります。話題を切り替えますか？新しい会話を開始すると精度が向上します。
新しい会話を開始します 0" x-transition.duration.150ms class="px-3 pt-2.5 pb-0 flex flex-wrap gap-1.5"> = 3 ? '3 ページの制限を超えています。このページはコンテキストとして送信されません。' : ''" class="inline-flex items-center gap-1.5 max-w-full px-2 py-1 borderrounded-lg text-xs font-medium">
0" x-transition:enter="transitionease-outduration-100" x-transition:enter-start="opacity-0 -translate-y-1" x-transition:enter-end="opacity-100translate-y-0" x-transition:leave="transitionease-induration-75" x-transition:leave-start="opacity-100translate-y-0" x-transition:leave-end="opacity-0 -translate-y-1" class="chatbot-mention-menu absolu

tebottom-full left-0 right-0 mb-1 mx-2 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700rounded-xlshadow-lg z-20overflow-hidden">
現在のページ ↑↓ ナビゲート
↵ 選択
esc 閉じる 0 ? 'text-violet-600 dark:text-violet-400 bg-violet-50 dark:bg-violet-500/10' : 'text-gray-400 dark:text-gray-500 hover:text-gray-600 dark:hover:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'" class="flex items-center justify-center w-8 h-8 bg-transparent border-nonerounded-lgcursor-pointertransition-all period-150" aria-label="コンテキストを追加" title="コンテキストを追加">
{ addCurrentPage(); showContextMenu = false; }" class="w-full flex items-center gap-3 px-3 py-2 text-left bg-transparent border-nonecursor-pointer text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/50transition-colors">
このページを追加
ページに言及する
スタンドアロン スタンドアロン導入ドキュメント
Kubernetes Kubernetes 導入ドキュメント
改善できる点は何でしょうか?
あなたのフィードバックは、アシスタントの回答を改善し、修正すべきドキュメントのギャップを特定するのに役立ちます。
さらに助けが必要ですか? Discord に参加してください:
https://discord.gg/y9efgEmppm
独自のエージェントを使用したいですか? Solo MCP サーバーを追加して、ドキュメントを直接クエリします。ここから始めましょう:
https://search.solo.io/ 。
Copyright © エージェントゲートウェイ、一連の LF プロジェクト、LLC
Web サイトの利用規約、商標ポリシー、その他のプロジェクト ポリシーについては、を参照してください。
https://lfprojects.org 。
© 2026 Solo.io, Inc. 全著作権所有。

## Original Extract

The open-source gateway for AI traffic — LLM, MCP, A2A, and HTTP in one data plane.

agentgateway | Agent Connectivity Solved Skip to content Docs
Standalone
Kubernetes Tutorials
Standalone
Kubernetes Models
Blog
Enterprise
Community
Light
Docs
Standalone
Kubernetes
Tutorials
Standalone
Kubernetes
Models
Blog
Enterprise
Community Get Started
GitHub
One high-performance gateway for service, LLM, and MCP traffic.
An open source HTTP and gRPC gateway that handles traditional application traffic and AI-native protocols in one data plane. Route, secure, observe, and govern services, LLM provider traffic, MCP tools, and agent-to-agent communication without stitching together separate gateways.
See it in action
View on GitHub
Discord LLM Gateway
Route to OpenAI, Claude, Gemini, or any self-hosted model with credentialing and failover.
Native protocol support for Model Context Protocol and agent-to-agent traffic.
HTTP, gRPC, and TCP traffic — observable by default with mTLS and OIDC built in.
CLI / IDE Claude, Codex, OpenCode AI agents LangSmith, CrewAI, ADK Apps & clients AI agents OpenAI Anthropic MCP / A2A APIs / Services Traffic Management LLM Gateway Inference Routing MCP / A2A gateway REQUEST · INBOUND ROUTE · SECURE · OBSERVE · GOVERN RESPONSE · OUTBOUND Contributing Companies Linux Foundation Solo.io Microsoft Apple Alibaba Adobe AWS Cisco Salesforce Huawei Amdocs Linux Foundation Solo.io Microsoft Apple Alibaba Adobe AWS Cisco Salesforce Huawei Amdocs Everything you need for agent traffic — in one binary.
One control plane for service traffic, LLM routing, MCP, A2A, and inference — with policy and observability that the platform team actually trusts.
Service Gateway HTTP, gRPC, TCP — observable by default. Run any north–south or east–west traffic through a single binary. Zero-config TLS, OIDC, mTLS-rotation, native Envoy compatibility for those who want it.
LLM Gateway Route, fail over, and budget across 12+ model backends. Drop-in OpenAI-compatible API in front of OpenAI, Anthropic, Bedrock, Gemini, Vertex, Cohere, OSS Llama runs — with per-team token budgets, semantic caching and prompt redaction.
Inference Routing Smart routing across self-hosted GPU pools. Latency-aware, cost-aware, model-aware routing for self-hosted inference. Plug-in vLLM, TGI, Triton; agentgateway picks the warmest replica.
Discover, sign, scope and observe every tool call your agents make. Treat MCP servers like microservices — with versioning, RBAC, and audit trails out of the box.
agentgateway speaks the Agent-to-Agent protocol natively. Route invocations between LangChain, CrewAI, ADK and your own runtime — with identity, tracing, and replay.
Security & Observability Per-call traces, logs, and policy decisions. OpenTelemetry by default. Per-tool, per-agent, per-tenant counters. OPA-evaluated allow/deny on every hop, with tamper-evident audit logs.
Quick start Get started with agentgateway in less than 5 minutes — local-first.
Connect agentgateway to MCP servers with built-in discovery and auth.
Deployment options Binary, Docker, or Kubernetes — pick what fits your stack.
Calling all agent-curious tool-builders, platform engineers, and AI enthusiasts — come help shape the future of agentic connectivity.
Discord
GitHub
"The future won’t be built by standalone agents, MCP servers or LLMs — it’s shaped by their interconnection and ability to work together seamlessly. To unlock their full potential, we must apply policies, ensure control and maintain clear visibility into their interactions. This is where agentgateway plays a pivotal role — bridging not only agent-to-agent (A2A) communication but also agent-to-MCP servers, filling a critical gap in the ecosystem. I look forward to seeing the project’s continued momentum within the Linux Foundation."
"AI agents are rapidly transforming how enterprises work and innovate. To adopt them responsibly at scale, organizations need open and interoperable gateways that provide governance, visibility, and security. Agentgateway delivers the foundation enterprises need. I’m excited to see the community come together to accelerate the open foundation our customers need to scale AI on cloud platforms like CoreWeave."
"Agentic AI demands purpose-built infrastructure, not just another software layer. This requires rethinking compute, storage, and data movement from the ground up, so retrofitting legacy systems doesn’t work. The agentgateway project is a solid step toward that future. We’re happy to see it hosted by the Linux Foundation, where open source and community can drive the adoption and longevity that AI infrastructure requires."
"We are excited to welcome the agentgateway project to the Linux Foundation, ensuring that best practices for agentic workflows remain free and open to all. Agentgateway complements emerging specifications like A2A and MCP, and offers scalable, specification-aligned infrastructure for agent communication thereby empowering customers and developers to build robust agentic workflows across platforms."
"I’m excited to see Solo.io donate the agentgateway project to the Linux Foundation. Open sourcing an AI-native connectivity solution that understands MCP and A2A is a big win for the community. It helps us scale AI responsibly, with the security, flexibility, and governance we need in the real world."
"Building reliable AI agents is a challenge, especially when every step involves non-deterministic calls to LLMs, tools, and autonomous agents. Agentgateway’s integration with OpenTelemetry provides a robust foundation for observability, allowing us to treat each request-response pair as an evaluable unit. This capability is crucial for ensuring system-level accuracy and trustworthiness, paving the way for a true ‘AI mesh’ that empowers teams to scale, secure and optimize their AI workflows."
"One of the biggest open security problems today is how to do MCP security effectively. While there are a lot of problems in this space that the community doesn’t know how to address, the agentgateway project provides a first step toward addressing some of the important issues with basic role-based access control and visibility of actions to MCP servers. I look forward to seeing how this project adapts and evolves to handle the complex, evolving threats in this space under open source governance."
"The rapid evolution of the AI landscape demands robust, vendor-neutral infrastructure for how agents communicate with each other and with external tools. Without it, we risk stifling innovation and adoption. The agentgateway project, hosted by the Linux Foundation, is a crucial step in creating that common ground. We are excited to partner with Solo.io and support a community-driven foundation for the future of interoperable AI."
"One of the biggest open security problems today is how to do MCP security effectively. While there are a lot of problems in this space that the community doesn't know how to address, the agentgateway project provides a first step toward addressing some of the important issues with basic role-based access control and visibility of actions to MCP servers. I look forward to seeing how this project adapts and evolves to handle the complex, evolving threats in this space under open source governance."
An open source HTTP and gRPC gateway that handles traditional application traffic and AI-native protocols in one data plane. Route, secure, observe, and govern services, LLM provider traffic, MCP tools, and agent-to-agent communication without stitching together separate gateways.
Ask AI
Agentgateway assistant
0" class="flex items-center justify-center w-8 h-8 bg-transparent border-none rounded-lg text-gray-500 cursor-pointer transition-all duration-150 hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-700 dark:hover:text-gray-50" aria-label="Export conversation" title="Export as Markdown">
Ask me anything about agentgateway configuration, features, or usage.
Note: AI-generated content might contain errors; please verify and test all returned information.
Tip: one topic per conversation gives the best results. Use the + button in the chat header to start a new conversation.
•
•
• Rate limit reached
The assistant keeps a rolling history of 3 exchanges. Any older messages are no longer included in the context. Switching topics? Starting a new conversation improves accuracy.
Start new conversation 0" x-transition.duration.150ms class="px-3 pt-2.5 pb-0 flex flex-wrap gap-1.5"> = 3 ? 'Exceeds the 3-page limit — this page will not be sent as context.' : ''" class="inline-flex items-center gap-1.5 max-w-full px-2 py-1 border rounded-lg text-xs font-medium">
0" x-transition:enter="transition ease-out duration-100" x-transition:enter-start="opacity-0 -translate-y-1" x-transition:enter-end="opacity-100 translate-y-0" x-transition:leave="transition ease-in duration-75" x-transition:leave-start="opacity-100 translate-y-0" x-transition:leave-end="opacity-0 -translate-y-1" class="chatbot-mention-menu absolute bottom-full left-0 right-0 mb-1 mx-2 bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl shadow-lg z-20 overflow-hidden">
Current page ↑↓ navigate
↵ select
esc dismiss 0 ? 'text-violet-600 dark:text-violet-400 bg-violet-50 dark:bg-violet-500/10' : 'text-gray-400 dark:text-gray-500 hover:text-gray-600 dark:hover:text-gray-300 hover:bg-gray-100 dark:hover:bg-gray-700'" class="flex items-center justify-center w-8 h-8 bg-transparent border-none rounded-lg cursor-pointer transition-all duration-150" aria-label="Add context" title="Add context">
{ addCurrentPage(); showContextMenu = false; }" class="w-full flex items-center gap-3 px-3 py-2 text-left bg-transparent border-none cursor-pointer text-sm text-gray-700 dark:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/50 transition-colors">
Add this page
Mention a page
Standalone Standalone deployment docs
Kubernetes Kubernetes deployment docs
What could be improved?
Your feedback helps us improve assistant answers and identify docs gaps we should fix.
Need more help? Join us on Discord:
https://discord.gg/y9efgEmppm
Want to use your own agent? Add the Solo MCP server to query our docs directly. Get started here:
https://search.solo.io/ .
Copyright © agentgateway a Series of LF Projects, LLC
For web site terms of use, trademark policy and other project policies please see
https://lfprojects.org .
© 2026 Solo.io, Inc. All rights reserved.

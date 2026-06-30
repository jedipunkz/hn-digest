---
source: "https://eddi.labs.ai/"
hn_url: "https://news.ycombinator.com/item?id=48732078"
title: "EDDI – sovereign AI orchestrator, now OpenSSF Gold and a UN (UNIDO) partner"
article_title: "The Enterprise AI Orchestrator | EDDI"
author: "ginccc"
captured_at: "2026-06-30T13:34:25Z"
capture_tool: "hn-digest"
hn_id: 48732078
score: 1
comments: 0
posted_at: "2026-06-30T12:55:53Z"
tags:
  - hacker-news
  - translated
---

# EDDI – sovereign AI orchestrator, now OpenSSF Gold and a UN (UNIDO) partner

- HN: [48732078](https://news.ycombinator.com/item?id=48732078)
- Source: [eddi.labs.ai](https://eddi.labs.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T12:55:53Z

## Translation

タイトル: EDDI – ソブリン AI オーケストレーター、現在は OpenSSF Gold および国連 (UNIDO) パートナー
記事のタイトル: エンタープライズ AI オーケストレーター |エディ
説明: EDDI v6 — エンタープライズ AI エージェント オーケストレーション プラットフォーム。ビジュアルな管理 UI、65 の MCP ツール、セキュリティ最優先のアーキテクチャ、GDPR、EU AI 法、HIPAA、および 15 以上のフレームワークに対する組み込みの規制コンプライアンスを備えたセルフホスト型の構成主導型マルチエージェント システム。

記事本文:
エンタープライズ AI オーケストレーター | EDDI コンテンツへスキップ 🏆 UN Recognition LABS.AI が産業用 AI の UNIDO 信頼パートナーに選ばれました 続きを読む →
機能 🖥️ EDDI マネージャー ビジュアルエージェント管理 UI 🔌 AI ネイティブ制御用の MCP Server 65 ツール ⚙️ Config-as-Code JSON 構成、ボイラープレートなし 🔐 セキュリティ第一の Vault、監査証跡、eval() なし 🚀 パフォーマンス 数百万の同時スレッド 🤖 マルチエージェントインテントのルーティングと検出 📊 可観測性 パイプラインのログと監査トレイル🧪 コード品質 9,600 以上のテスト、失敗ゼロ、カバレッジ >90% 🧩 AI 対応 コーディングエージェント向けに構築 🧠 メモリとコンテキスト 永続メモリとドリーム統合 📚 RAG 7 エンベディングプロバイダー、5 ベクトルストア 📈 モデルカスケード コスト最適化されたマルチモデルルーティング ⏰ ハートビート、cron、ドリームのスケジューリングサイクル すべての機能を表示 → ソリューション 💡 なぜ EDDI なのか?セルフホスト型エンタープライズ AI プラットフォーム ⚖️ と代替案 Spring AI、Flowise、n8n との比較 🏛️ コンプライアンス GDPR、CCPA、PIPEDA など 📜 実績 機関の支援と実際の導入 📋 ユースケース 実際の導入パターン ドキュメント スタートガイド EN English EN Deutsch DE Español ES Français FR Português PT AR 中文 ZH 日本語 JA 한국어 KO हिन्दी HI ไทย TH 特徴
ライト モードに切り替える ダーク モードに切り替える Enterprise AI Orchestrator
管理されたマルチエージェント AI システムを展開するための、オープンソースの本番環境対応プラットフォーム。構成主導型、セルフホスト型、コンプライアンス対応 — ビジュアル管理 UI、65 の MCP ツール、エンタープライズ セキュリティが組み込まれています。現在は v6 (最新: 6.1.2) です。
エンタープライズ AI オーケストレーションは断片化されており、苦痛を伴います。チームは壊れたライフサイクルに陥っています。
Flowise、n8n、またはスクリプトを使用してチームのプロトタイプを作成し、SC からすべてを書き直す

生産用のラッチ。ビジュアルプロトタイプは完全に破棄されます。
AI ライブラリを使用するということは、REST コントローラー、認証レイヤー、状態管理を最初から構築することを意味します。すべてのプロンプト調整には再デプロイが必要です。
非決定論的な AI エージェントを決定論的な Camunda/Temporal ワークフローに強制すると、深刻なインピーダンスの不整合と脆弱なアーキテクチャが発生します。
EDDI は、ライブラリではなく、展開可能な AI オーケストレーション プラットフォームです。すぐに使えるビジュアル管理、コードとしての構成、エンタープライズ セキュリティ:
エージェントのロジック、パイプライン、ツール定義は JSON 構成であり、コンパイルされたコードではありません。 Prompt エンジニアは、再デプロイメントを行わずに、React UI または REST API を介して即座に反復処理を行います。
EDDI はモデル コンテキスト プロトコルを介してその機能を公開し、Claude Desktop との対話を可能にします。エージェントは外部 MCP ツールを利用することもできます。
eval() もエスケープもありません。 Vault の統合、URL 検証、暗号化監査証跡はアーキテクチャの基盤です。
すべてのパイプライン ステップは、トークン、コスト、タイミング、ツール呼び出しなど、不変の監査証跡で記録されます。完全な CQRS テレメトリ台帳。
I/O バウンドの LLM ワークロード向けに数百万もの軽量スレッドを備えたエンタープライズ グレードのランタイム上に構築されています。イベント ループのブロックやシングル スレッドのボトルネックはありません。
インテントベースのエージェント検出、管理された会話、エージェントトリガー、A/B ルーティング。インテント + ユーザーごとに 1 つの会話が自動作成および自動管理されます。
「エンジンが厳密であるため、AI は創造的になることができます。」
UNIDO 産業用 AI の信頼できるパートナー
LABS.AI は、国連工業開発機関 (UNIDO) によって、グローバル サウス向け産業 AI の信頼できるパートナーとして選ばれました。
Red Hat 認定コンテナ Docker イメージ IBM Apache 2.0 によって認定 ライセンス済み 100% オープンソース & エンタープライズ対応 9,600 以上のテスト · 失敗なし 厳格な CI/CD & >90% コード カバレッジ O

penSSF Gold Linux Foundation のセキュリティおよび品質認定の最高レベル Docker Hub コンテナは Docker Hub から取得 CI のパス · CodeQL Clean 自動化されたビルド、セキュリティ スキャン & コード分析 実証済みのテクノロジーに基づいて構築
Java 25 Enterprise ランタイム Quarkus クラウドネイティブ、高速 LangChain4j マルチプロバイダー LLM サポート MongoDB ドキュメント ストア PostgreSQL リレーショナル ストア Docker コンテナ対応 Kubernetes Orchestration OpenShift Red Hat 認定 次のステップ
EDDI は、オープンソースのエンタープライズ グレードの AI オーケストレーション プラットフォームです。これにより、チームはコンパイルされたコードではなく JSON 構成を使用して、AI を活用したエージェントを構築、構成、デプロイできるようになります。 EDDI は、本番環境に対応した React 管理 UI (EDDI マネージャー)、組み込み REST API、会話状態管理、セキュリティ (OIDC/Keycloak)、不変監査証跡、および 65 個の MCP ツールを備えた完全なプラットフォームを提供します。これらはすべて Docker または Kubernetes 経由でデプロイ可能です。
LangChain、Spring AI、LangChain4j などの AI ライブラリは構築ブロックを提供しますが、REST コントローラー、認証、会話状態管理、監査ログ、および管理 UI を自分で構築する必要があります。 EDDI は、ライブラリではなく、展開可能なミドルウェア プラットフォームです。これらはすべて、すぐに Docker 経由でデプロイできる状態で提供されます。
はい。 EDDI は、大規模な I/O バウンドの同時実行を実現する軽量の仮想スレッドを備えたエンタープライズ グレードのランタイムに基づいて構築されています。 MongoDB と PostgreSQL をサポートし、組み込みの OIDC/Keycloak 認証を含み、コンプライアンス (EU AI 法を含む) のための不変の暗号監査証跡を提供し、NATS JetStream 経由で水平に拡張します。
Model Context Protocol (MCP) は、Claude Desktop、IDE プラグイン、カスタム クライアントなどの AI アシスタントがプログラムで外部ツールと対話できるようにするオープン スタンダードです。 EDDI は、次の範囲にわたる 65 の MCP ツールを公開します。

会話管理、エージェント管理、セットアップの自動化、スケジュール管理、診断。
EDDI は同様のビジュアル構築の目的を果たしますが、エンタープライズ グレードのアーキテクチャを備えています。 Flowise や n8n とは異なり、EDDI は eval() やコード ブロックを使用せず、エンタープライズ グレードの同時実行のために何百万もの軽量仮想スレッドを実行し、OIDC/Keycloak 認証をサポートし、データを MongoDB または PostgreSQL に保存します。規制された業界向けに設計されています。
探索専用。すべてのデータは 48 時間ごとの 03:00 UTC に消去されます。これは、誰でも探索できる EDDI のライブ デモ インスタンスです。これは評価と実験のみを目的としています。運用ワークロードに使用したり、機密データを保存したり、その可用性に依存したりしないでください。
Red Hat 認定コンテナ Apache 2.0 ライセンス取得済み 9,600 以上のテスト · 障害ゼロ Java 25 · Quarkus · LangChain4j OpenSSF Gold The Enterprise AI Orchestrator
構成主導型。自己ホスト型。生産準備完了。
© 2026 Labs.ai · Apache 2.0 · インプリント · プライバシー ポリシー
ウィーンで始まりました。ヨーロッパの❤️とともに開発されました。世界のために作られました。
お客様のブラウジング エクスペリエンスを向上させ、トラフィックを分析するために Cookie を使用します。 「すべてを受け入れる」をクリックすると、Cookie の使用に同意したことになります。詳細については、プライバシー ポリシーを参照してください。
以下で Cookie の設定を調整できます。必要なCookieを無効にすることはできません。
Web サイトが適切に機能するために必要です。
訪問者が当社の Web サイトとどのようにやり取りするかを理解するのに役立ちます。
Web サイト全体で訪問者を追跡し、関連する広告を表示するために使用されます。

## Original Extract

EDDI v6 — Enterprise AI Agent Orchestration Platform. Self-hosted, configuration-driven multi-agent system with visual management UI, 65 MCP tools, security-first architecture, and built-in regulatory compliance for GDPR, EU AI Act, HIPAA, and 15+ frameworks.

The Enterprise AI Orchestrator | EDDI Skip to content 🏆 UN Recognition LABS.AI selected as UNIDO Trusted Partner for Industrial AI Read more →
Features 🖥️ EDDI Manager Visual agent management UI 🔌 MCP Server 65 tools for AI-native control ⚙️ Config-as-Code JSON config, zero boilerplate 🔐 Security-First Vault, audit trails, no eval() 🚀 Performance Millions of concurrent threads 🤖 Multi-Agent Intent routing & discovery 📊 Observability Pipeline logs & audit trails 🧪 Code Quality 9,600+ tests, zero failures, >90% coverage 🧩 AI-Ready Built for coding agents 🧠 Memory & Context Persistent memory & dream consolidation 📚 RAG 7 embedding providers, 5 vector stores 📈 Model Cascading Cost-optimized multi-model routing ⏰ Scheduling Heartbeats, cron & dream cycles View all features → Solutions 💡 Why EDDI? The self-hosted enterprise AI platform ⚖️ vs. Alternatives Compare with Spring AI, Flowise, n8n 🏛️ Compliance GDPR, CCPA, PIPEDA & more 📜 Track Record Institutional backing & real deployments 📋 Use Cases Real-world deployment patterns Docs Get Started EN English EN Deutsch DE Español ES Français FR Português PT العربية AR 中文 ZH 日本語 JA 한국어 KO हिन्दी HI ไทย TH Features
Switch to Light Mode Switch to Dark Mode The Enterprise AI Orchestrator
The open-source, production-ready platform for deploying governed, multi-agent AI systems. Configuration-driven, self-hosted, and compliance-ready — with a visual management UI, 65 MCP tools, and enterprise security built in. Now in v6 (latest: 6.1.2).
Enterprise AI orchestration is fragmented and painful . Teams are stuck in a broken lifecycle:
Teams prototype with Flowise, n8n, or scripts — then rewrite everything from scratch for production. Visual prototypes are discarded entirely.
Using AI libraries means building REST controllers, auth layers, and state management from scratch . Every prompt tweak requires redeployment.
Forcing non-deterministic AI agents into deterministic Camunda/Temporal workflows creates severe impedance mismatches and brittle architectures.
EDDI is a deployable AI orchestration platform — not a library. Visual management, config-as-code, and enterprise security, out of the box:
Agent logic, pipelines, and tool definitions are JSON configurations — not compiled code. Prompt engineers iterate instantly via the React UI or REST API, without redeployment.
EDDI exposes its capabilities via the Model Context Protocol — enabling Claude Desktop to interact. Agents can also consume external MCP tools .
No eval() , no escapes. Vault integration, URL validation, and cryptographic audit trails are architectural foundations .
Every pipeline step is logged with an immutable audit trail — tokens, cost, timing, tool calls. Full CQRS telemetry ledger.
Built on an enterprise-grade runtime with millions of lightweight threads for I/O-bound LLM workloads. No event loop blocking, no single-threaded bottlenecks.
Intent-based agent discovery, managed conversations, agent triggers, and A/B routing. One conversation per intent+user , auto-created and auto-managed.
"The engine is strict so the AI can be creative."
UNIDO Trusted Partner for Industrial AI
LABS.AI has been selected by the United Nations Industrial Development Organization (UNIDO) as a Trusted Partner for Industrial AI for the Global South.
Red Hat Certified Container Docker image certified by IBM Apache 2.0 Licensed 100% open-source & enterprise-ready 9,600+ Tests · 0 Failures Rigorous CI/CD & >90% code coverage OpenSSF Gold Highest tier of Linux Foundation security & quality certification Docker Hub Container pulls from Docker Hub CI Passing · CodeQL Clean Automated builds, security scanning & code analysis Built on Proven Technology
Java 25 Enterprise runtime Quarkus Cloud-native, fast LangChain4j Multi-provider LLM support MongoDB Document store PostgreSQL Relational store Docker Container-ready Kubernetes Orchestration OpenShift Red Hat certified Next Steps
EDDI is an open-source, enterprise-grade AI orchestration platform. It enables teams to build, configure, and deploy AI-powered agents using JSON configuration rather than compiled code. EDDI provides a complete platform with a production-ready React management UI (the EDDI Manager), built-in REST APIs, conversation state management, security (OIDC/Keycloak), immutable audit trails, and 65 MCP tools — all deployable via Docker or Kubernetes.
AI libraries like LangChain, Spring AI, and LangChain4j give you building blocks — but you still need to build REST controllers, authentication, conversation state management, audit logging, and management UIs yourself. EDDI is a deployable middleware platform , not a library. It provides all of this out of the box, ready to deploy via Docker.
Yes. EDDI is built on an enterprise-grade runtime with lightweight virtual threads for massive I/O-bound concurrency. It supports MongoDB and PostgreSQL, includes built-in OIDC/Keycloak authentication, provides immutable cryptographic audit trails for compliance (including EU AI Act), and scales horizontally via NATS JetStream.
The Model Context Protocol (MCP) is an open standard that allows AI assistants like Claude Desktop, IDE plugins, and custom clients to interact with external tools programmatically. EDDI exposes 65 MCP tools spanning conversation management, agent administration, setup automation, schedule management, and diagnostics.
EDDI serves a similar visual-building purpose but with enterprise-grade architecture. Unlike Flowise and n8n, EDDI uses no eval() or code blocks, runs millions of lightweight virtual threads for enterprise-grade concurrency, supports OIDC/Keycloak authentication, and stores data in MongoDB or PostgreSQL. It is designed for regulated industries.
For exploration only. All data is wiped every 48 hours at 03:00 UTC. This is a live demo instance of EDDI available for anyone to explore. It is intended solely for evaluation and experimentation — please do not use it for production workloads, store sensitive data, or rely on its availability.
Red Hat Certified Container Apache 2.0 Licensed 9,600+ Tests · Zero Failures Java 25 · Quarkus · LangChain4j OpenSSF Gold The Enterprise AI Orchestrator
Configuration-driven. Self-hosted. Production-ready.
© 2026 Labs.ai · Apache 2.0 · Imprint · Privacy Policy
Initiated in Vienna. Developed with ❤️ in Europe. Made for the World.
We use cookies to enhance your browsing experience and analyze our traffic. By clicking "Accept All", you consent to our use of cookies. For more information, see our Privacy Policy .
You can adjust your cookie settings below. Necessary cookies cannot be disabled.
Required for the website to function properly.
Help us understand how visitors interact with our website.
Used to track visitors across websites to display relevant advertisements.

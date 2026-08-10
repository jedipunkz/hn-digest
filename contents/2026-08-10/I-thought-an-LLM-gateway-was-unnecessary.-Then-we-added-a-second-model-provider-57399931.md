---
source: "https://github.com/maximhq/bifrost/"
hn_url: "https://news.ycombinator.com/item?id=49247732"
title: "I thought an LLM gateway was unnecessary. Then we added a second model provider"
article_title: "GitHub - maximhq/bifrost: Fastest enterprise AI gateway (50x faster than LiteLLM) with adaptive load balancer, cluster mode, guardrails, 1000+ models support & <100 µs overhead at 5k RPS. · GitHub"
author: "Swapnoneel"
captured_at: "2026-08-10T18:42:29Z"
capture_tool: "hn-digest"
hn_id: 49247732
score: 1
comments: 0
posted_at: "2026-08-10T18:30:50Z"
tags:
  - hacker-news
  - translated
---

# I thought an LLM gateway was unnecessary. Then we added a second model provider

- HN: [49247732](https://news.ycombinator.com/item?id=49247732)
- Source: [github.com](https://github.com/maximhq/bifrost/)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T18:30:50Z

## Translation

タイトル: LLM ゲートウェイは不要だと思っていました。次に、2 番目のモデルプロバイダーを追加しました
記事のタイトル: GitHub - maximhq/bifrost: 適応型ロード バランサー、クラスター モード、ガードレール、1000 以上のモデルのサポート、5k RPS で 100 μs 未満のオーバーヘッドを備えた最速のエンタープライズ AI ゲートウェイ (LiteLLM より 50 倍高速)。 · GitHub
説明: 適応型ロード バランサー、クラスター モード、ガードレール、1000 以上のモデルのサポート、5k RPS で 100 μs 未満のオーバーヘッドを備えた最速のエンタープライズ AI ゲートウェイ (LiteLLM より 50 倍高速)。 - maximhq/ビフロスト

記事本文:
GitHub - maximhq/bifrost: 適応型ロード バランサー、クラスター モード、ガードレール、1000 以上のモデルのサポート、5k RPS で 100 μs 未満のオーバーヘッドを備えた最速のエンタープライズ AI ゲートウェイ (LiteLLM より 50 倍高速)。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
マキシムハイク
/
ビフロスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
dev ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6,219 コミット 6,219 コミット .claude/ スキル

.claude/ スキル .github .github .greptile .greptile cli cli cmd/ e2eseed cmd/ e2eseed コミュニティ コミュニティ コア コア ドキュメント ドキュメント サンプル サンプル フレームワーク フレームワーク helm-charts helm-charts nix nix npx npx プラグイン プラグイン レシピ レシピ スクリプト スクリプト terraform terraform テスト テスト トランスポート トランスポート ui ui .coderabbit.yaml .coderabbit.yaml .cursorignore .cursorignore .dockerignore .dockerignore .editorconfig .editorconfig .envrc .envrc .gitattributes .gitattributes .gitignore .gitignore .infisical.json .infisical.json .nvmrc .nvmrc .pre-commit-config.yaml .pre-commit-config.yaml .snyk .snyk AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md config.json config.json flake.lock flake.lock flake.nix flake.nix パルス.yaml パルス.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
決して停止しない AI アプリケーションを構築する最速の方法
Bifrost は、単一の OpenAI 互換 API を通じて 23 以上のプロバイダー (OpenAI、Anthropic、AWS Bedrock、Google Vertex など) へのアクセスを統合する高性能 AI ゲートウェイです。構成を必要とせずに数秒で導入し、自動フェイルオーバー、ロード バランシング、セマンティック キャッシュ、エンタープライズ グレードの機能を利用できます。
1 分以内にゼロから本番環境に対応した AI ゲートウェイに移行します。
# ローカルにインストールして実行する
npx -y @maximhq/bifrost
# または Docker を使用する
docker run -p 8080:8080 maximhq/bifrost
ステップ 2: Web UI 経由で設定する
# 組み込みの Web インターフェイスを開きます
http://localhost:8080 を開きます
ステップ 3: 最初の API 呼び出しを行う
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"モデル": "openai/gpt-4o-mini",
"メッセージ": [{"役割": "ユーザー", "コンテンツ": "こんにちは、ビフロスト!"}]
} '
それです！あなたのAIが

teway は、視覚的な構成、リアルタイムの監視、および分析のための Web インターフェイスを使用して実行されます。
ゲートウェイのセットアップ - HTTP API のデプロイメント
Go SDK セットアップ - 直接統合
Bifrost は、実稼働 AI システムを大規模に実行するチーム向けのエンタープライズ グレードのプライベート展開をサポートします。
プライベート ネットワーキング、カスタム セキュリティ制御、ガバナンスに加えて、エンタープライズ展開では、適応負荷分散、クラスタリング、ガードレール、MCP ゲートウェイ、エンタープライズ グレードの規模と信頼性を目的に設計されたその他の機能を含む高度な機能が利用可能になります。
統一インターフェイス - すべてのプロバイダー向けの単一の OpenAI 互換 API
マルチプロバイダーのサポート - OpenAI、Anthropic、AWS Bedrock、Google Vertex、Azure、Cerebras、Cohere、Mistral、Ollama、Groq など
自動フォールバック - ダウンタイムゼロのプロバイダーとモデル間のシームレスなフェイルオーバー
負荷分散 - 複数の API キーとプロバイダーにわたるインテリジェントなリクエスト分散
モデル コンテキスト プロトコル (MCP) - AI モデルが外部ツール (ファイル システム、Web 検索、データベース) を使用できるようにします。
セマンティック キャッシュ - コストと遅延を削減するためのセマンティックの類似性に基づくインテリジェントな応答キャッシュ
マルチモーダル サポート - テキスト、画像、オーディオ、ストリーミングをすべて共通のインターフェイスでサポートします。
カスタム プラグイン - 分析、監視、カスタム ロジックのための拡張可能なミドルウェア アーキテクチャ
ガバナンス - 使用状況の追跡、レート制限、およびきめ細かいアクセス制御
予算管理 - 仮想キー、チーム、顧客予算による階層的なコスト管理
ユーザー プロビジョニング (OIDC) - OAuth 2.0 / OIDC ログインとチーム、ロール、ビジネス ユニットのバックグラウンド ディレクトリ同期
可観測性 - ネイティブ Prometheus メトリクス、分散トレース、包括的なログ記録
シークレット管理 - 環境変数とデプロイメント シークレットを使用した安全な API キー管理

ts
ゼロ構成スタートアップ - 動的なプロバイダー構成ですぐに開始
ドロップイン置換 - OpenAI/Anthropic/GenAI API を 1 行のコードで置き換えます
SDK 統合 - コード変更なしで人気の AI SDK をネイティブ サポート
構成の柔軟性 - Web UI、API 駆動、またはファイルベースの構成オプション
Bifrost は、柔軟性を最大限に高めるためにモジュラー アーキテクチャを使用しています。
ビフロスト/
§── npx/ # 簡単にインストールするための NPX スクリプト
§── core/ # コア機能と共有コンポーネント
│ §── Providers/ # プロバイダー固有の実装 (OpenAI、Anthropic など)
│ §── schemas/ # Bifrost 全体で使用されるインターフェイスと構造体
│ └── bifrost.go # Bifrost のメイン実装
§── Framework/ # データ永続化のためのフレームワーク コンポーネント
│ §── configstore/ # 構成ストレージのバックエンド
│ §── logstore/ # ログストレージバックエンドをリクエストします
│ └── Vectorstore/ # ベクター ストレージ
§── Transports/ # HTTP ゲートウェイおよびその他のインターフェイス層
│ └── bifrost-http/ # HTTP トランスポートの実装
§── ui/ # HTTP ゲートウェイの Web インターフェイス
§── plugins/ # 拡張可能なプラグインシステム
│ §── ガバナンス/ # 予算管理とアクセス制御
│ §── jsonparser/ # JSON 解析および操作ユーティリティ
│ §──logging/ # ログ記録と分析をリクエストする
│ §── maxim/ # Maxim の可観測性の統合
│ §── mocker/ # テストと開発のためのモック応答
│ §── semanticcache/ # インテリジェントな応答キャッシュ
│ └── telemetry/ # 監視と可観測性
§── docs/ # ドキュメントとガイド
└── テスト/ # 包括的なテスト スイート
開始オプション
ニーズに合った導入方法を選択してください。
最適な用途: 言語に依存しない統合、microserv

アイス、および実稼働デプロイメント
# NPX - 30 秒で始められます
npx -y @maximhq/bifrost
# Docker - 運用準備完了
docker run -p 8080:8080 -v $( pwd ) /data:/app/data maximhq/bifrost
機能: Web UI、リアルタイム監視、マルチプロバイダー管理、ゼロ構成起動
詳細: ゲートウェイ設定ガイド
最適な用途: 最大のパフォーマンスと制御を備えた直接 Go 統合
github.com/maximhq/bifrost/core を取得してください
機能: ネイティブ Go API、組み込みデプロイメント、カスタム ミドルウェア統合
最適な用途: コード変更なしで既存のアプリケーションを移行する
# OpenAI SDK
-base_url = "https://api.openai.com"
+base_url = "http://localhost:8080/openai"
# Anthropic SDK
-base_url = "https://api.anthropic.com"
+base_url = "http://localhost:8080/anthropic"
# Google GenAI SDK
- api_endpoint = "https://generative language.googleapis.com"
+ api_endpoint = "http://localhost:8080/genai"
詳細: 統合ガイド
Bifrost は、AI リクエストに実質的にオーバーヘッドを加えません。 5,000 RPS ベンチマークを継続した場合、ゲートウェイによるリクエストあたりのオーバーヘッドの追加はわずか 11 μs でした。
完璧な成功率 - 5,000 RPS でも 100% のリクエスト成功率
最小限のオーバーヘッド - リクエストあたりの追加レイテンシは 15 μs 未満
効率的なキューイング - 平均待機時間はマイクロ秒未満
高速キー選択 - 重み付けされた API キーを選択するのに最大 10 ns
完全なベンチマーク: パフォーマンス分析
完全なドキュメント: https://docs.getbifrost.ai
ゲートウェイのセットアップ - 30 秒での HTTP API の導入
Go SDK セットアップ - Go の直接統合
プロバイダー構成 - マルチプロバイダーのセットアップ
マルチプロバイダーのサポート - すべてのプロバイダーに単一の API
MCP 統合 - 外部ツール呼び出し
セマンティック キャッシュ - インテリジェントな応答キャッシュ
フォールバックとロード バランシング - 信頼性機能
予算管理 - コスト管理とガバナンス
OpenAI SDK - ドロップイン OpenAI リプラ

セメント
Anthropic SDK - ドロップイン Anthropic の代替品
AWS Bedrock SDK - AWS Bedrock の統合
Google GenAI SDK - ドロップイン GenAI 代替品
LiteLLM SDK - LiteLLM の統合
LangChain SDK - LangChain の統合
カスタムプラグイン - 機能を拡張する
クラスタリング - マルチノード展開
シークレット管理 - 安全なキー管理
実稼働環境への導入 - スケーリングとモニタリング
Discord に参加して、コミュニティのサポートやディスカッションを行ってください。
素早いセットアップ支援とトラブルシューティング
ベストプラクティスと構成のヒント
コミュニティでのディスカッションとサポート
統合に関するリアルタイムのヘルプ
あらゆる種類の貢献を歓迎します。以下については、貢献ガイドを参照してください。
開発環境のセットアップ
コード規約とベストプラクティス
開発要件とビルド手順については、「開発セットアップ ガイド」を参照してください。
このプロジェクトは、Apache 2.0 ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
アダプティブ ロード バランサー、クラスター モード、ガードレール、1000 以上のモデルをサポートし、5k RPS で 100 μs 未満のオーバーヘッドを備えた最速のエンタープライズ AI ゲートウェイ (LiteLLM の 50 倍高速)。
www.getmaxim.ai/bifrost トピック
Readme Apache-2.0 ライセンスの行動規範
行動規範 セキュリティポリシー
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
1.0k フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Fastest enterprise AI gateway (50x faster than LiteLLM) with adaptive load balancer, cluster mode, guardrails, 1000+ models support & <100 µs overhead at 5k RPS. - maximhq/bifrost

GitHub - maximhq/bifrost: Fastest enterprise AI gateway (50x faster than LiteLLM) with adaptive load balancer, cluster mode, guardrails, 1000+ models support & <100 µs overhead at 5k RPS. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
maximhq
/
bifrost
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dev Branches Tags Go to file Code Open more actions menu Folders and files
6,219 Commits 6,219 Commits .claude/ skills .claude/ skills .github .github .greptile .greptile cli cli cmd/ e2eseed cmd/ e2eseed community community core core docs docs examples examples framework framework helm-charts helm-charts nix nix npx npx plugins plugins recipes recipes scripts scripts terraform terraform tests tests transports transports ui ui .coderabbit.yaml .coderabbit.yaml .cursorignore .cursorignore .dockerignore .dockerignore .editorconfig .editorconfig .envrc .envrc .gitattributes .gitattributes .gitignore .gitignore .infisical.json .infisical.json .nvmrc .nvmrc .pre-commit-config.yaml .pre-commit-config.yaml .snyk .snyk AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md config.json config.json flake.lock flake.lock flake.nix flake.nix pulse.yaml pulse.yaml View all files Repository files navigation
The fastest way to build AI applications that never go down
Bifrost is a high-performance AI gateway that unifies access to 23+ providers (OpenAI, Anthropic, AWS Bedrock, Google Vertex, and more) through a single OpenAI-compatible API. Deploy in seconds with zero configuration and get automatic failover, load balancing, semantic caching, and enterprise-grade features.
Go from zero to production-ready AI gateway in under a minute.
# Install and run locally
npx -y @maximhq/bifrost
# Or use Docker
docker run -p 8080:8080 maximhq/bifrost
Step 2: Configure via Web UI
# Open the built-in web interface
open http://localhost:8080
Step 3: Make your first API call
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "openai/gpt-4o-mini",
"messages": [{"role": "user", "content": "Hello, Bifrost!"}]
} '
That's it! Your AI gateway is running with a web interface for visual configuration, real-time monitoring, and analytics.
Gateway Setup - HTTP API deployment
Go SDK Setup - Direct integration
Bifrost supports enterprise-grade, private deployments for teams running production AI systems at scale.
In addition to private networking, custom security controls, and governance, enterprise deployments unlock advanced capabilities including adaptive load balancing, clustering, guardrails, MCP gateway, and other features designed for enterprise-grade scale and reliability.
Unified Interface - Single OpenAI-compatible API for all providers
Multi-Provider Support - OpenAI, Anthropic, AWS Bedrock, Google Vertex, Azure, Cerebras, Cohere, Mistral, Ollama, Groq, and more
Automatic Fallbacks - Seamless failover between providers and models with zero downtime
Load Balancing - Intelligent request distribution across multiple API keys and providers
Model Context Protocol (MCP) - Enable AI models to use external tools (filesystem, web search, databases)
Semantic Caching - Intelligent response caching based on semantic similarity to reduce costs and latency
Multimodal Support - Support for text, images, audio, and streaming, all behind a common interface.
Custom Plugins - Extensible middleware architecture for analytics, monitoring, and custom logic
Governance - Usage tracking, rate limiting, and fine-grained access control
Budget Management - Hierarchical cost control with virtual keys, teams, and customer budgets
User Provisioning (OIDC) - OAuth 2.0 / OIDC login with background directory sync for teams, roles, and business units
Observability - Native Prometheus metrics, distributed tracing, and comprehensive logging
Secrets Management - Secure API key management with environment variables and deployment secrets
Zero-Config Startup - Start immediately with dynamic provider configuration
Drop-in Replacement - Replace OpenAI/Anthropic/GenAI APIs with one line of code
SDK Integrations - Native support for popular AI SDKs with zero code changes
Configuration Flexibility - Web UI, API-driven, or file-based configuration options
Bifrost uses a modular architecture for maximum flexibility:
bifrost/
├── npx/ # NPX script for easy installation
├── core/ # Core functionality and shared components
│ ├── providers/ # Provider-specific implementations (OpenAI, Anthropic, etc.)
│ ├── schemas/ # Interfaces and structs used throughout Bifrost
│ └── bifrost.go # Main Bifrost implementation
├── framework/ # Framework components for data persistence
│ ├── configstore/ # Configuration storage backends
│ ├── logstore/ # Request logging storage backends
│ └── vectorstore/ # Vector storages
├── transports/ # HTTP gateway and other interface layers
│ └── bifrost-http/ # HTTP transport implementation
├── ui/ # Web interface for HTTP gateway
├── plugins/ # Extensible plugin system
│ ├── governance/ # Budget management and access control
│ ├── jsonparser/ # JSON parsing and manipulation utilities
│ ├── logging/ # Request logging and analytics
│ ├── maxim/ # Maxim's observability integration
│ ├── mocker/ # Mock responses for testing and development
│ ├── semanticcache/ # Intelligent response caching
│ └── telemetry/ # Monitoring and observability
├── docs/ # Documentation and guides
└── tests/ # Comprehensive test suites
Getting Started Options
Choose the deployment method that fits your needs:
Best for: Language-agnostic integration, microservices, and production deployments
# NPX - Get started in 30 seconds
npx -y @maximhq/bifrost
# Docker - Production ready
docker run -p 8080:8080 -v $( pwd ) /data:/app/data maximhq/bifrost
Features: Web UI, real-time monitoring, multi-provider management, zero-config startup
Learn More: Gateway Setup Guide
Best for: Direct Go integration with maximum performance and control
go get github.com/maximhq/bifrost/core
Features: Native Go APIs, embedded deployment, custom middleware integration
Best for: Migrating existing applications with zero code changes
# OpenAI SDK
- base_url = "https://api.openai.com"
+ base_url = "http://localhost:8080/openai"
# Anthropic SDK
- base_url = "https://api.anthropic.com"
+ base_url = "http://localhost:8080/anthropic"
# Google GenAI SDK
- api_endpoint = "https://generativelanguage.googleapis.com"
+ api_endpoint = "http://localhost:8080/genai"
Learn More: Integration Guides
Bifrost adds virtually zero overhead to your AI requests. In sustained 5,000 RPS benchmarks, the gateway added only 11 µs of overhead per request.
Perfect Success Rate - 100% request success rate even at 5k RPS
Minimal Overhead - Less than 15 µs additional latency per request
Efficient Queuing - Sub-microsecond average wait times
Fast Key Selection - ~10 ns to pick weighted API keys
Complete Benchmarks: Performance Analysis
Complete Documentation: https://docs.getbifrost.ai
Gateway Setup - HTTP API deployment in 30 seconds
Go SDK Setup - Direct Go integration
Provider Configuration - Multi-provider setup
Multi-Provider Support - Single API for all providers
MCP Integration - External tool calling
Semantic Caching - Intelligent response caching
Fallbacks & Load Balancing - Reliability features
Budget Management - Cost control and governance
OpenAI SDK - Drop-in OpenAI replacement
Anthropic SDK - Drop-in Anthropic replacement
AWS Bedrock SDK - AWS Bedrock integration
Google GenAI SDK - Drop-in GenAI replacement
LiteLLM SDK - LiteLLM integration
LangChain SDK - LangChain integration
Custom Plugins - Extend functionality
Clustering - Multi-node deployment
Secrets Management - Secure key management
Production Deployment - Scaling and monitoring
Join our Discord for community support and discussions.
Quick setup assistance and troubleshooting
Best practices and configuration tips
Community discussions and support
Real-time help with integrations
We welcome contributions of all kinds! See our Contributing Guide for:
Setting up the development environment
Code conventions and best practices
For development requirements and build instructions, see our Development Setup Guide .
This project is licensed under the Apache 2.0 License - see the LICENSE file for details.
Fastest enterprise AI gateway (50x faster than LiteLLM) with adaptive load balancer, cluster mode, guardrails, 1000+ models support & <100 µs overhead at 5k RPS.
www.getmaxim.ai/bifrost Topics
Readme Apache-2.0 license Code of conduct
Code of conduct Security policy
Security policy Activity Custom properties Stars
1.0k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/lightseekorg/smg"
hn_url: "https://news.ycombinator.com/item?id=49063132"
title: "Shepherd Model Gateway – Model-routing gateway for LLM deployments"
article_title: "GitHub - lightseekorg/smg: Engine-agnostic LLM gateway in Rust. Full OpenAI & Anthropic API compatibility across vLLM, TRT-LLM, TokenSpeed, SGLang, OpenAI, Gemini & more. Industry-first gRPC pipeline, KV cache-aware routing, chat history, tokenization caching, Responses API, embeddings, WASM plugins\n[truncated]"
author: "rekl"
captured_at: "2026-07-26T22:50:59Z"
capture_tool: "hn-digest"
hn_id: 49063132
score: 1
comments: 0
posted_at: "2026-07-26T22:36:58Z"
tags:
  - hacker-news
  - translated
---

# Shepherd Model Gateway – Model-routing gateway for LLM deployments

- HN: [49063132](https://news.ycombinator.com/item?id=49063132)
- Source: [github.com](https://github.com/lightseekorg/smg)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T22:36:58Z

## Translation

タイトル: Shepherd Model Gateway – LLM 導入用のモデル ルーティング ゲートウェイ
記事のタイトル: GitHub - lighteekorg/smg: Rust のエンジンに依存しない LLM ゲートウェイ。 vLLM、TRT-LLM、TokenSpeed、SGLang、OpenAI、Gemini などにわたる OpenAI および Anthropic API の完全な互換性。業界初の gRPC パイプライン、KV キャッシュ対応ルーティング、チャット履歴、トークン化キャッシュ、応答 API、埋め込み、WASM プラグイン
[切り捨てられた]
説明: Rust のエンジンに依存しない LLM ゲートウェイ。 vLLM、TRT-LLM、TokenSpeed、SGLang、OpenAI、Gemini などにわたる OpenAI および Anthropic API の完全な互換性。業界初の gRPC パイプライン、KV キャッシュ対応ルーティング、チャット履歴、トークン化キャッシュ、応答 API、埋め込み、WASM プラグイン、MCP、およびマルチテナント AUT
[切り捨てられた]

記事本文:
GitHub - lighteekorg/smg: Rust のエンジンに依存しない LLM ゲートウェイ。 vLLM、TRT-LLM、TokenSpeed、SGLang、OpenAI、Gemini などにわたる OpenAI および Anthropic API の完全な互換性。業界初の gRPC パイプライン、KV キャッシュ対応ルーティング、チャット履歴、トークン化キャッシュ、応答 API、埋め込み、WASM プラグイン、MCP、およびマルチテナント認証。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。レル

oad を使用してセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ライトシークオーグ
/
smg
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,002 コミット 2,002 コミット .cargo .cargo .github .github バインディング バインディング クライアント クライアント クレート クレート デプロイ/ helm/ smg デプロイ/ helm/ smg docker docker docs docs e2e_test e2e_test 例/ wasm 例/ wasm grpc_servicer grpc_servicer model_gateway model_gateway scripts scripts .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.toml Cargo.toml GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md REVIEW.md REVIEW.md Clippy.toml Clippy.toml mkdocs.yml mkdocs.yml mypy.ini mypy.ini pytest.ini pytest.ini ruff.toml ruff.toml Rustfmt.toml Rustfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
大規模な LLM 導入向けの、エンジンに依存しない高性能モデル ルーティング ゲートウェイ。ワーカーのライフサイクル管理を一元化し、HTTP/gRPC/OpenAI 互換バックエンド全体でトラフィックのバランスをとり、履歴ストレージ、MCP ツール、プライバシーに配慮したワークフローに対するエンタープライズ対応の制御を提供します。
🚀 GPU使用率を最大化
キャッシュ対応ルーティングは、推論エンジンの KV キャッシュ状態 (vLLM、TensorRT-LLM、TokenSpeed、SGLang) を理解して、プレフィックスを再利用し、冗長な計算を削減します。
🔌 1 つの API、任意のバックエンド
セルフホスト モデル (vLLM、TensorRT-LLM、TokenSpeed、SGLang) またはクラウド プロバイダー (OpenAI、Anthropic、

Gemini、Bedrock など) を単一の統合エンドポイント経由で実現します。
⚡ スピードを追求した設計
gRPC パイプライン、ミリ秒未満のルーティング決定、ゼロコピー トークン化を備えたネイティブ Rust。サーキット ブレーカーと自動フェイルオーバーにより、動作が継続されます。
🔒 エンタープライズコントロール
OIDC によるマルチテナント レート制限、カスタム ロジック用の WebAssembly プラグイン、およびインフラストラクチャ内で会話履歴を保持するプライバシー境界。
📊 完全な可観測性
40 を超える Prometheus メトリクス、OpenTelemetry トレース、リクエストの相関関係を含む構造化された JSON ログにより、すべてのレイヤーで何が起こっているかを正確に把握できます。
API カバレッジ: OpenAI チャット/コンプリーション/埋め込み、エージェント用の応答 API、Anthropic メッセージ、MCP ツールの実行。
インストール — 好みの方法を選択してください:
# ドッカー
docker pull lighteekorg/smg:latest
# パイソン
pip インストール SMG
# 錆びる
貨物インストールSMG
実行 — SMG を推論ワーカーに向けます。
# 単一のワーカー
smg launch --worker-urls http://localhost:8000
# キャッシュ対応ルーティングを使用する複数のワーカー
smg launch --worker-urls http://gpu1:8000 http://gpu2:8000 --policy queue_aware
# 高可用性メッシュを使用
smg launch --worker-urls http://gpu1:8000 --enable-mesh \
--mesh-advertise-host 10.0.0.1 --mesh-peer-urls 10.0.0.2:39527
使用 — ゲートウェイにリクエストを送信します。
カール http://localhost:30000/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {"モデル": "llama3", "メッセージ": [{"役割": "ユーザー", "コンテンツ": "Hello!"}]} '
それだけです。 SMG は現在、ワーカー間でリクエストの負荷を分散しています。
自己ホスト型
クラウドプロバイダー
vLLM
OpenAI
TensorRT-LLM
人間的
トークンスピード
Google ジェミニ
SGLang
AWS の基盤
オラマ
Azure OpenAI
任意の OpenAI 互換サーバー
任意の OpenAI 互換プロバイダー
特長
特徴
説明
8 ルーティングポリシー
キャッシュ対応、ラウンドロビン、二乗、一貫性のあるハッシュ、

prefix_hash、手動、ランダム、バケット
gRPC パイプライン
ストリーミング、推論抽出、ツール呼び出し解析を備えたネイティブ gRPC
MCPの統合
モデル コンテキスト プロトコルを介して外部ツール サーバーに接続する
高可用性
マルチノード導入向けの SWIM プロトコルを使用したメッシュ ネットワーキング
チャット履歴
プラグ可能ストレージ: PostgreSQL、Oracle、Redis、またはインメモリ
WASMプラグイン
カスタム WebAssembly ロジックで拡張する
回復力
サーキット ブレーカー、バックオフによる再試行、レート制限
ドキュメント
はじめに
インストールと最初のステップ
建築
SMG の仕組み
構成
CLI リファレンスとオプション
APIリファレンス
OpenAI互換エンドポイント
Kubernetesのセットアップ
クラスター内の検出と本番環境のセットアップ
貢献する
寄付を歓迎します!詳細については、「貢献ガイド」を参照してください。
Rust のエンジンに依存しない LLM ゲートウェイ。 vLLM、TRT-LLM、TokenSpeed、SGLang、OpenAI、Gemini などにわたる OpenAI および Anthropic API の完全な互換性。業界初の gRPC パイプライン、KV キャッシュ対応ルーティング、チャット履歴、トークン化キャッシュ、応答 API、埋め込み、WASM プラグイン、MCP、およびマルチテナント認証。
Readme Apache-2.0 ライセンスの行動規範
貢献活動 カスタム プロパティ スター
125 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Engine-agnostic LLM gateway in Rust. Full OpenAI & Anthropic API compatibility across vLLM, TRT-LLM, TokenSpeed, SGLang, OpenAI, Gemini & more. Industry-first gRPC pipeline, KV cache-aware routing, chat history, tokenization caching, Responses API, embeddings, WASM plugins, MCP, and multi-tenant aut
[truncated]

GitHub - lightseekorg/smg: Engine-agnostic LLM gateway in Rust. Full OpenAI & Anthropic API compatibility across vLLM, TRT-LLM, TokenSpeed, SGLang, OpenAI, Gemini & more. Industry-first gRPC pipeline, KV cache-aware routing, chat history, tokenization caching, Responses API, embeddings, WASM plugins, MCP, and multi-tenant auth. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
lightseekorg
/
smg
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,002 Commits 2,002 Commits .cargo .cargo .github .github bindings bindings clients clients crates crates deploy/ helm/ smg deploy/ helm/ smg docker docker docs docs e2e_test e2e_test examples/ wasm examples/ wasm grpc_servicer grpc_servicer model_gateway model_gateway scripts scripts .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.toml Cargo.toml GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md REVIEW.md REVIEW.md clippy.toml clippy.toml mkdocs.yml mkdocs.yml mypy.ini mypy.ini pytest.ini pytest.ini ruff.toml ruff.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
Engine-agnostic, high-performance model-routing gateway for large-scale LLM deployments. Centralizes worker lifecycle management, balances traffic across HTTP/gRPC/OpenAI-compatible backends, and provides enterprise-ready control over history storage, MCP tooling, and privacy-sensitive workflows.
🚀 Maximize GPU Utilization
Cache-aware routing understands your inference engine's KV cache state—whether vLLM, TensorRT-LLM, TokenSpeed, or SGLang—to reuse prefixes and reduce redundant computation.
🔌 One API, Any Backend
Route to self-hosted models (vLLM, TensorRT-LLM, TokenSpeed, SGLang) or cloud providers (OpenAI, Anthropic, Gemini, Bedrock, and more) through a single unified endpoint.
⚡ Built for Speed
Native Rust with gRPC pipelines, sub-millisecond routing decisions, and zero-copy tokenization. Circuit breakers and automatic failover keep things running.
🔒 Enterprise Control
Multi-tenant rate limiting with OIDC, WebAssembly plugins for custom logic, and a privacy boundary that keeps conversation history within your infrastructure.
📊 Full Observability
40+ Prometheus metrics, OpenTelemetry tracing, and structured JSON logs with request correlation—know exactly what's happening at every layer.
API Coverage: OpenAI Chat/Completions/Embeddings, Responses API for agents, Anthropic Messages, and MCP tool execution.
Install — pick your preferred method:
# Docker
docker pull lightseekorg/smg:latest
# Python
pip install smg
# Rust
cargo install smg
Run — point SMG at your inference workers:
# Single worker
smg launch --worker-urls http://localhost:8000
# Multiple workers with cache-aware routing
smg launch --worker-urls http://gpu1:8000 http://gpu2:8000 --policy cache_aware
# With high availability mesh
smg launch --worker-urls http://gpu1:8000 --enable-mesh \
--mesh-advertise-host 10.0.0.1 --mesh-peer-urls 10.0.0.2:39527
Use — send requests to the gateway:
curl http://localhost:30000/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {"model": "llama3", "messages": [{"role": "user", "content": "Hello!"}]} '
That's it. SMG is now load-balancing requests across your workers.
Self-Hosted
Cloud Providers
vLLM
OpenAI
TensorRT-LLM
Anthropic
TokenSpeed
Google Gemini
SGLang
AWS Bedrock
Ollama
Azure OpenAI
Any OpenAI-compatible server
Any OpenAI-compatible provider
Features
Feature
Description
8 Routing Policies
cache_aware, round_robin, power_of_two, consistent_hashing, prefix_hash, manual, random, bucket
gRPC Pipeline
Native gRPC with streaming, reasoning extraction, and tool call parsing
MCP Integration
Connect external tool servers via Model Context Protocol
High Availability
Mesh networking with SWIM protocol for multi-node deployments
Chat History
Pluggable storage: PostgreSQL, Oracle, Redis, or in-memory
WASM Plugins
Extend with custom WebAssembly logic
Resilience
Circuit breakers, retries with backoff, rate limiting
Documentation
Getting Started
Installation and first steps
Architecture
How SMG works
Configuration
CLI reference and options
API Reference
OpenAI-compatible endpoints
Kubernetes Setup
In-cluster discovery and production setup
Contributing
We welcome contributions! See Contributing Guide for details.
Engine-agnostic LLM gateway in Rust. Full OpenAI & Anthropic API compatibility across vLLM, TRT-LLM, TokenSpeed, SGLang, OpenAI, Gemini & more. Industry-first gRPC pipeline, KV cache-aware routing, chat history, tokenization caching, Responses API, embeddings, WASM plugins, MCP, and multi-tenant auth.
Readme Apache-2.0 license Code of conduct
Contributing Activity Custom properties Stars
125 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

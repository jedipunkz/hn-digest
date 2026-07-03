---
source: "https://github.com/NVIDIA/nvcf"
hn_url: "https://news.ycombinator.com/item?id=48772345"
title: "NVCF: Deploy and Route GPU-Accelerated AI Workloads at Scale"
article_title: "GitHub - NVIDIA/nvcf: Platform for deploying and routing GPU-accelerated inference, streaming, and batch workloads at scale. · GitHub"
author: "mastabadtomm"
captured_at: "2026-07-03T08:49:40Z"
capture_tool: "hn-digest"
hn_id: 48772345
score: 1
comments: 0
posted_at: "2026-07-03T08:18:02Z"
tags:
  - hacker-news
  - translated
---

# NVCF: Deploy and Route GPU-Accelerated AI Workloads at Scale

- HN: [48772345](https://news.ycombinator.com/item?id=48772345)
- Source: [github.com](https://github.com/NVIDIA/nvcf)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T08:18:02Z

## Translation

タイトル: NVCF: GPU アクセラレーション AI ワークロードの大規模な展開とルーティング
記事のタイトル: GitHub - NVIDIA/nvcf: GPU アクセラレーションによる推論、ストリーミング、バッチ ワークロードを大規模に展開およびルーティングするためのプラットフォーム。 · GitHub
説明: GPU アクセラレーションによる推論、ストリーミング、およびバッチ ワークロードを大規模に展開およびルーティングするためのプラットフォーム。 - エヌビディア/nvcf

記事本文:
GitHub - NVIDIA/nvcf: GPU アクセラレーションによる推論、ストリーミング、バッチ ワークロードを大規模に展開およびルーティングするためのプラットフォーム。 · GitHub
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
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
エヌビディア
/
nvcf
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
376 コミット 376 コミット .claude/ スキル .claude/ スキル .codex/ スキル .codex/ スキル .cursor/ スキル .cursor/ スキル .github .github ai-tooling ai-tooling ci ci デプロイ デプロイ ドキュメント ドキュメントの例 例 シダ シダ マイグレーション マイグレーション プラットフォーム プラットフォーム ルール ルール src src テスト/ bdd テスト/ bdd ツール ツール.allowed-licenses.txt .allowed-licenses.txt .bazelignore .bazelignore .bazelrc .bazelrc .bazelversion .bazelversion .gitignore .gitignore .rumdl.toml .rumdl.toml AGENTS.md AGENTS.md BAZEL.md BAZEL.md BUILD.bazel BUILD.bazel CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス MODULE.bazel MODULE.bazel MODULE.bazel.lock MODULE.bazel.lock 通知 通知 README.md README.md SECURITY.md SECURITY.md WORKSPACE.bzlmod WORKSPACE.bzlmod dependency.md dependency.md go.work.bazel go.work.bazel License-compliance.md License-compliance.md setup.sh setup.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ドキュメント |ロードマップ |インストール | API リファレンス |貢献する |ライセンス | build.nvidia.com NVCF を活用
NVIDIA Cloud Functions (NVCF) は、GPU アクセラレーションによるワークロードを大規模に展開、管理、実行するためのプラットフォームです。推論、ストリーミング、その他の GPU 作業をワーカー クラスターにルーティングするため、独自に実行するインフラストラクチャを減らして、要求の厳しいワークロードを拡張できます。
このモノリポジトリには、NVCF サービス コード、展開アセット、ドキュメント、
例、CLI コード、エージェントのスキル、検証ツール。
NVCF は、関数のライフサイクルと呼び出しを管理する Kubernetes サービスとして実行されます。
ルーティング、GPU クラスターの統合、アーティファクト アクセス、シークレット、可観測性、
操作。
共同

ntrol プレーンは NVCF API を公開し、機能と展開を管理します
状態を管理し、秘密管理を処理し、プラットフォーム操作を調整します。
呼び出しプレーンは、HTTP、ストリーミング、および gRPC リクエストを受信し、適用されます。
ルーティングとレート制限を行い、実行中の関数ワークロードに作業を送信します。
GPU クラスターは、NVIDIA クラスター エージェント (NVCA) を介して接続します。 NVCAレジスタ
GPU リソースを管理し、GPU ノードでのワークロードの実行を管理します。
関数アーティファクトは、NVCF 導入環境がアクセスできるレジストリに存在します。
可観測性、ダッシュボード、ランブックは、オペレーターが健全性を監視し、
ワークロードの動作をデバッグします。
次の図は、自己管理型 NVCF がどのようにリージョンと GPU にまたがるかを示しています。
クラスター。
NVCF 関数は、長時間実行される呼び出し可能なワークロードです。次の場合に関数を使用します。
クライアントには推論、ストリーミング、または別のサービス スタイルの GPU 用のエンドポイントが必要です
ワークフロー。ワークロードが単一の場合、関数をコンテナとしてパッケージ化できます。
サービスとヘルスおよび推論エンドポイントを使用するか、Helm チャートとしてサービスを提供します。
ワークロードには複数の調整されたコンテナ、サービス、サイドカーなどが必要です
Kubernetes リソース。
NVCF タスクは、実行から完了までの非同期ワークロードです。バッチにタスクを使用する
推論、評価、微調整、データ準備、またはその他の GPU ジョブ
呼び出し後にオンラインのままにするのではなく、終了してステータスを報告する必要があります
エンドポイント。ワークロードが単一の場合、タスクをコンテナとしてパッケージ化できます。
サービスとヘルスおよび推論エンドポイントを使用するか、Helm チャートとしてサービスを提供します。
ワークロードには複数の調整されたコンテナ、サービス、サイドカーなどが必要です
Kubernetes リソース。
能力
何をするのか
統合コントロールプレーン
マルチリージョン GPU クラスター全体でリクエストを管理およびルーティングします。
負荷分散されたワークロード ルーティング
ワーカーの可用性に基づいて、推論、ストリーミング、カスタム ワークロードのバランスをとります。
乗算

プロトコル
さまざまなワークロードとクライアントのニーズに対応する複数のプロトコルをサポートします。
マルチクラスターの自動スケーリング
クラスター全体でワークロードをゼロから最大までスケーリングします。
混合GPUのサポート
異なる GPU 要件を持つワークロードに対して、クラスター全体で混合 GPU タイプをサポートします。
ヘルスチェックとテレメトリ
ヘルスチェックとテレメトリを通じてワーカーのステータスとリクエストの遅延を追跡します。
使用法
自己管理型 NVCF 展開をインストールし、 nvcf-cli を設定した後、
一般的な関数のワークフローは次のとおりです。
nvcf-cli 初期化
nvcf-cli APIキー生成
# サンプル ファイルを作成する前に、関数イメージを使用してサンプル ファイルを更新します。
nvcf-cli 関数の作成 --input-file src/clis/nvcf-cli/examples/create-function.json
nvcf-cli 関数デプロイ作成
nvcf-cli 関数呼び出し --request-body ' {"message": "hello world"} '
完全なセットアップ、クリーンアップ、構成フローについては、を参照してください。
docs/user/cli.md および
docs/user/quickstart.md 。
エリア
パス
目的
コントロールプレーン
src/コントロールプレーンサービス/
NVCF の機能と展開状態を管理する API とサービス。
呼び出しプレーン
src/呼び出しプレーンサービス/
HTTP 呼び出し、gRPC プロキシ、レート制限、LLM ゲートウェイ パス、およびリクエストの承認。
コンピューティングプレーン
src/コンピューティングプレーンサービス/
GPU クラスターの統合、キャッシュ サービス、イメージ認証情報、ESS エージェント、テレメトリ収集。
CLIとライブラリ
src/clis/、src/libraries/
ユーザーおよび開発者のクライアントと、共有の Go および Python コード。
導入
デプロイ/ 、移行/
ヘルム チャート、スタックのインストール、インフラストラクチャ サービス、およびデータストアの移行。
ドキュメント
docs/user/ 、 docs/dev/ 、 fern/
自己管理のユーザー ドキュメント、開発者ドキュメント、公開ドキュメントのナビゲーション。
例
例/
ローカル開発ガイド、関数サンプル、および負荷テスト アセット。
ツール
ツール/
ビルド、ドキュメント、依存関係、ライセンス、および検証ユーティリティ。
AIツール

AIツール/
NVCF ユーザーおよび開発者向けのパブリック エージェント スキルとワークフロー ヘルパー。
Bazel を使用した構築
Bazel は、モノリポジトリ全体にわたるビルド、テスト、およびパッケージ化ツールです。ネイティブ
サブツリー ( src/clis/nvcf-cli 、 src/libraries/go/lib ) は完全に下に構築されます。
今日はバゼル。フェーズ B では、さらに Bazel 足場が着陸しました。
アップストリーム所有のサービス ツリー: nvcf-grpc-proxy 、 nvcf-ratelimiter 、
nvcf-nats-auth-callout-service 、 nvcf-cache/nvcf-unbound (dns-cache)、
nvcf-image-credential-helper 、および nvca 。彼らの BUILD.bazel 、
MODULE.bazel および rules/oci/ ファイルは、次の場合に自動的に選択されます。
サブツリーはアンブレラに同期されます。傘からでもできる
から離れることなく、それらのいずれかの OCI イメージを構築、テスト、生成できます。
モノレポ。
カール -fSL -o ~ /.local/bin/bazel \
" https://github.com/bazelbuild/bazelisk/releases/download/v1.25.0/bazelisk-linux- $( dpkg --print-architecture ) "
chmod +x ~ /.local/bin/bazel
# ネイティブサブツリー
bazel build //src/clis/nvcf-cli:nvcf-cli # ホストバイナリ
bazel test //src/clis/nvcf-cli/... # 単体テスト
bazel build //src/clis/nvcf-cli:dist # 5 つのプラットフォームすべて
# フェーズ B のアップストリームの例: grpc-proxy マルチアーキテクチャ OCI イメージを構築する
bazel ビルド //src/invocation-plane-services/grpc-proxy:image_index
bazel テスト //src/invocation-plane-services/grpc-proxy/...
# ツリー全体を実行する
ベゼルテスト //...
クイックスタート (macOS):
醸造インストールバゼリスク
bazel ビルド //src/clis/nvcf-cli:nvcf-cli
bazel テスト //src/clis/nvcf-cli/...
bazel ビルド //src/clis/nvcf-cli:dist
ビルドはデフォルトで構成されたリモート キャッシュから読み取られ、ローカルにはアップロードされません
結果。そのキャッシュと Bazel に到達できるネットワーク パスから離れている場合
ローカル実行が開始される前に失敗した場合は、そのビルドのリモート キャッシュを無効にします。
bazel build --remote_cache= //src/clis/nvcf-cli:nvcf-cli
ローカル専用パスを作成するには

目的として、ユーザー Bazel にオーバーライドを追加します
構成:
echo ' build --remote_cache= ' >> ~ /.bazelrc.user
開発ボックス (企業ネットワークまたは VPN が必要) からキャッシュをシードするには、次を追加します
--config=リモート書き込み :
bazel build --config=remote-write //src/clis/nvcf-cli/...
完全なセットアップ、日常的なコマンド、OCI イメージのビルド/プッシュ、スタンピング、キャッシュ、
リモート キャッシュ プローブと CI マップは BAZEL.md にあります。
CLI 固有の開発者フローについては、 src/clis/nvcf-cli/README.md を参照してください。
最初のプル リクエストを開く前に、ローカル ビルドをセットアップしてテストします。
環境:
Bazelis を介して Bazel をインストールします。 Bazel を使用した構築を参照してください。
bazel build //src/clis/nvcf-cli:nvcf-cli でツールチェーンを確認します。
プッシュする前に、関連するテストをローカルで実行します。
bazel test //src/clis/nvcf-cli/... 、またはツリー全体の bazel test //...。
完全なセットアップ、キャッシュ、CI マップについては、BAZEL.md を参照してください。
GitHub の問題 #27 が最新です
この四半期の公開ロードマップ。その問題を真実の情報源として使用し、
アクティブな優先事項、ステータスの更新、およびフォローアップの提案。
より広範な問題委員会は、
残りのバックログ。の対象外のアイデアやリクエストがある場合
四半期ごとのロードマップ?ディスカッションを開始する
または機能の問題を提出してください。
バグ、機能のアイデア、ドキュメントのリクエストを GitHub の問題としてファイルします。適切なテンプレートを使用し、タイトルにコンポーネント名を含めます (たとえば、[nvcf-nvca] ポッドが arm64 で起動できない)。
サポートと使用方法のヘルプについては、GitHub ディスカッションを使用してください。
セキュリティの脆弱性を報告するには、 SECURITY.md を参照してください。公開問題を開かないでください。
タイプミスの修正から新機能まで、あらゆる規模の貢献を歓迎します。参照
完全なガイドについては、CONTRIBUTING.md を参照してください。
NVCF は新しいオープンソース プロジェクトであり、私たちは積極的に
貢献ワークフロー。 GitHub プルを通じて外部からの貢献を受け入れます
今日のリクエストは、いくつかの一時的なものです

リポジトリが大きくなるにつれてしわが寄る
GitHub ネイティブ。
サービスを変更する前に、AGENTS.md と最も近いネストされたファイルを読んでください。
エージェント.md 。ネストされたファイルは、サービス固有のビルドに最適なソースです。
テスト、スタイル、レビューの期待。
従来のコミットを使用します。
ドキュメントのみの変更の場合は、 git diff --check を実行し、対象の変更を実行します。
変更されたファイルに適用される検証。
このプロジェクトは、貢献者規約の行動規範に従います。貢献者
この基準を維持することに同意します。参照
全文と適用については CODE_OF_CONDUCT.md
ガイドライン。
依存関係収集ガイドとツール:
tools/collect-dependency/README.md
および tools/collect-dependents 。
GPU アクセラレーションによる推論、ストリーミング、バッチ ワークロードを大規模に展開およびルーティングするためのプラットフォーム。
docs.nvidia.com/nvcf/overview
トピックス
Apache-2.0、不明なライセンスが見つかりました
ライセンスが見つかりました
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
21
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Platform for deploying and routing GPU-accelerated inference, streaming, and batch workloads at scale. - NVIDIA/nvcf

GitHub - NVIDIA/nvcf: Platform for deploying and routing GPU-accelerated inference, streaming, and batch workloads at scale. · GitHub
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
NVIDIA
/
nvcf
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
376 Commits 376 Commits .claude/ skills .claude/ skills .codex/ skills .codex/ skills .cursor/ skills .cursor/ skills .github .github ai-tooling ai-tooling ci ci deploy deploy docs docs examples examples fern fern migrations migrations platforms platforms rules rules src src tests/ bdd tests/ bdd tools tools .allowed-licenses.txt .allowed-licenses.txt .bazelignore .bazelignore .bazelrc .bazelrc .bazelversion .bazelversion .gitignore .gitignore .rumdl.toml .rumdl.toml AGENTS.md AGENTS.md BAZEL.md BAZEL.md BUILD.bazel BUILD.bazel CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MODULE.bazel MODULE.bazel MODULE.bazel.lock MODULE.bazel.lock NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md WORKSPACE.bzlmod WORKSPACE.bzlmod dependencies.md dependencies.md go.work.bazel go.work.bazel license-compliance.md license-compliance.md setup.sh setup.sh View all files Repository files navigation
Docs | Roadmap | Installation | API Reference | Contributing | License | build.nvidia.com Powered By NVCF
NVIDIA Cloud Functions (NVCF) is a platform for deploying, managing, and running GPU-accelerated workloads at scale. It routes inference, streaming, and other GPU work to worker clusters, so you can scale demanding workloads with less infrastructure to run yourself.
This monorepo contains NVCF service code, deployment assets, documentation,
examples, CLI code, agent skills, and validation tooling.
NVCF runs as Kubernetes services that manage function lifecycle, invocation
routing, GPU cluster integration, artifact access, secrets, observability, and
operations.
The control plane exposes the NVCF API, manages function and deployment
state, handles secret management, and coordinates platform operations.
The invocation plane receives HTTP, streaming, and gRPC requests, applies
routing and rate limiting, and sends work to running function workloads.
GPU clusters connect through the NVIDIA Cluster Agent (NVCA). NVCA registers
GPU resources and manages workload execution on GPU nodes.
Function artifacts live in registries that the NVCF deployment can access.
Observability, dashboards, and runbooks help operators monitor health and
debug workload behavior.
The following diagram shows how self-managed NVCF can span regions and GPU
clusters.
NVCF functions are long-running, invokable workloads. Use a function when a
client needs an endpoint for inference, streaming, or another service-style GPU
workflow. Functions can be packaged as a container when the workload is a single
service with health and inference endpoints, or as a Helm chart when the
workload needs multiple coordinated containers, services, sidecars, or other
Kubernetes resources.
NVCF tasks are asynchronous, run-to-completion workloads. Use a task for batch
inference, evaluation, fine-tuning, data preparation, or other GPU jobs that
should finish and report status instead of staying online behind an invocation
endpoint. Tasks can be packaged as a container when the workload is a single
service with health and inference endpoints, or as a Helm chart when the
workload needs multiple coordinated containers, services, sidecars, or other
Kubernetes resources.
Capability
What it does
Unified control plane
Manages and routes requests across multi-region GPU clusters.
Load-balanced workload routing
Balances inference, streaming, and custom workloads based on worker availability.
Multiple protocols
Supports multiple protocols for different workload and client needs.
Multi-cluster autoscaling
Scales workloads from zero to max across clusters.
Mixed GPU support
Supports mixed GPU types across clusters for workloads with different GPU requirements.
Health checks and telemetry
Tracks worker status and request latency through health checks and telemetry.
Usage
After installing a self-managed NVCF deployment and configuring nvcf-cli , a
typical function workflow is:
nvcf-cli init
nvcf-cli api-key generate
# Update the example file with your function image before creating it.
nvcf-cli function create --input-file src/clis/nvcf-cli/examples/create-function.json
nvcf-cli function deploy create
nvcf-cli function invoke --request-body ' {"message": "hello world"} '
For the full setup, cleanup, and configuration flow, see
docs/user/cli.md and
docs/user/quickstart.md .
Area
Paths
Purpose
Control plane
src/control-plane-services/
APIs and services that manage NVCF function and deployment state.
Invocation plane
src/invocation-plane-services/
HTTP invocation, gRPC proxying, rate limiting, LLM gateway paths, and request authorization.
Compute plane
src/compute-plane-services/
GPU cluster integration, cache services, image credentials, ESS Agent, and telemetry collection.
CLI and libraries
src/clis/ , src/libraries/
User and developer clients plus shared Go and Python code.
Deployment
deploy/ , migrations/
Helm charts, stack installation, infrastructure services, and datastore migrations.
Documentation
docs/user/ , docs/dev/ , fern/
Self-managed user docs, developer docs, and published docs navigation.
Examples
examples/
Local development guides, function samples, and load-test assets.
Tools
tools/
Build, docs, dependency, license, and validation utilities.
AI tooling
ai-tooling/
Public agent skills and workflow helpers for NVCF users and developers.
Building with Bazel
Bazel is the build, test, and packaging tool across the monorepo. Native
subtrees ( src/clis/nvcf-cli , src/libraries/go/lib ) build fully under
Bazel today. Phase B has additionally landed Bazel scaffolds in
upstream-owned service trees: nvcf-grpc-proxy , nvcf-ratelimiter ,
nvcf-nats-auth-callout-service , nvcf-cache/nvcf-unbound (dns-cache),
nvcf-image-credential-helper , and nvca . Their BUILD.bazel ,
MODULE.bazel , and rules/oci/ files are picked up automatically when
the subtrees are synced into the umbrella; from the umbrella you can
build, test, and produce OCI images for any of them without leaving the
monorepo.
curl -fSL -o ~ /.local/bin/bazel \
" https://github.com/bazelbuild/bazelisk/releases/download/v1.25.0/bazelisk-linux- $( dpkg --print-architecture ) "
chmod +x ~ /.local/bin/bazel
# Native subtrees
bazel build //src/clis/nvcf-cli:nvcf-cli # host binary
bazel test //src/clis/nvcf-cli/... # unit tests
bazel build //src/clis/nvcf-cli:dist # all 5 platforms
# Phase B upstream example: build the grpc-proxy multi-arch OCI image
bazel build //src/invocation-plane-services/grpc-proxy:image_index
bazel test //src/invocation-plane-services/grpc-proxy/...
# Run the full tree
bazel test //...
Quick start (macOS):
brew install bazelisk
bazel build //src/clis/nvcf-cli:nvcf-cli
bazel test //src/clis/nvcf-cli/...
bazel build //src/clis/nvcf-cli:dist
Builds read from the configured remote cache by default and do not upload local
results. If you are off the network path that can reach that cache and Bazel
fails before local execution starts, disable the remote cache for that build:
bazel build --remote_cache= //src/clis/nvcf-cli:nvcf-cli
To make the local-only path persistent, add the override to your user Bazel
config:
echo ' build --remote_cache= ' >> ~ /.bazelrc.user
To seed the cache from a dev box (corp network or VPN required), add
--config=remote-write :
bazel build --config=remote-write //src/clis/nvcf-cli/...
Full setup, day-to-day commands, OCI image build/push, stamping, caches,
remote-cache probe, and CI map live in BAZEL.md .
For CLI-specific developer flow see src/clis/nvcf-cli/README.md .
Before opening your first pull request, set up a local build and test
environment:
Install Bazel through bazelisk. See Building with Bazel .
Confirm your toolchain with bazel build //src/clis/nvcf-cli:nvcf-cli .
Run the relevant tests locally before pushing:
bazel test //src/clis/nvcf-cli/... , or bazel test //... for the full tree.
See BAZEL.md for the complete setup, cache, and CI map.
GitHub issue #27 is the current
public roadmap for the quarter. Use that issue as the source of truth for
active priorities, status updates, and follow-up proposals.
The broader issues board tracks the
remaining backlog. Have an idea or a request that is not covered by the
quarterly roadmap? Start a Discussion
or file a feature issue .
File bugs, feature ideas, and documentation requests as GitHub issues . Use the appropriate template and include the component name in the title (for example, [nvcf-nvca] Pod fails to start on arm64).
Use GitHub Discussions for support and usage help.
To report a security vulnerability see SECURITY.md . Do not open a public issue.
We welcome contributions of all sizes, from typo fixes to new features. See
CONTRIBUTING.md for the full guide.
NVCF is a new open source project, and we are actively smoothing the
contribution workflow. We accept external contributions through GitHub pull
requests today, with a few temporary wrinkles while the repository becomes more
GitHub-native .
Before changing a service, read AGENTS.md and the nearest nested
AGENTS.md . The nested file is the best source for service-specific build,
test, style, and review expectations.
Use Conventional Commits.
For documentation-only changes, run git diff --check and any targeted
validation that applies to the changed files.
This project follows the Contributor Covenant Code of Conduct. Contributors
agree to uphold this standard. See
CODE_OF_CONDUCT.md for the full text and enforcement
guidelines.
Dependency collection guide and tool:
tools/collect-dependencies/README.md
and tools/collect-dependencies .
Platform for deploying and routing GPU-accelerated inference, streaming, and batch workloads at scale.
docs.nvidia.com/nvcf/overview
Topics
Apache-2.0, Unknown licenses found
Licenses found
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
21
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

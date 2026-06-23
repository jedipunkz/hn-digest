---
source: "https://github.com/modelplaneai/modelplane"
hn_url: "https://news.ycombinator.com/item?id=48646721"
title: "Modelplane – The Open Source Control Plane for AI Inference"
article_title: "GitHub - modelplaneai/modelplane: The open source control plane for AI inference · GitHub"
author: "bassamtabbara"
captured_at: "2026-06-23T16:05:44Z"
capture_tool: "hn-digest"
hn_id: 48646721
score: 18
comments: 1
posted_at: "2026-06-23T15:35:22Z"
tags:
  - hacker-news
  - translated
---

# Modelplane – The Open Source Control Plane for AI Inference

- HN: [48646721](https://news.ycombinator.com/item?id=48646721)
- Source: [github.com](https://github.com/modelplaneai/modelplane)
- Score: 18
- Comments: 1
- Posted: 2026-06-23T15:35:22Z

## Translation

タイトル: Modelplane – AI 推論用のオープンソース コントロール プレーン
記事タイトル: GitHub - modelplaneai/modelplane: AI 推論用のオープンソース コントロール プレーン · GitHub
説明: AI 推論用のオープンソース コントロール プレーン。 GitHub でアカウントを作成して、modelplaneai/modelplane の開発に貢献してください。

記事本文:
GitHub - modelplaneai/modelplane: AI 推論用のオープンソース コントロール プレーン · GitHub
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
模型飛行機あい
/
模型飛行機
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
663 コミット 663 コミット .github .github apis apis design design docs docs function function nix nix スキーマ スキーマ .envrc .envrc .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.mdcrossplane-project.yamlcrossplane-project.yaml flake.lock flake.lock flake.nix flake.nix nix.sh nix.sh pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Modelplane は、オーケストレーションを行うために独自の環境にインストールして実行するソフトウェアです
モデル、サービス スタック、クラウド全体の基盤となるインフラストラクチャ、
ネオクラウドとオンプレミス。あらゆるインフラ上のあらゆるエンジン上のあらゆるモデルを実行します。
単一の GPU から分散されたマルチノードの導入まで。上に構築
Crossplane は、フリートを継続的に調整するアクティブなシステムです
宣言した状態に向けて: 推論クラスターのプロビジョニング、スケジューリング
互換性のあるクラスターへのデプロイメント、レプリカのスケーリング、重みのキャッシュ、および
トラフィックのルーティング。
プラットフォーム チームはクラスターをプロビジョニングし、ハードウェア クラスを公開します。開発者
モデルを宣言し、統合された OpenAI 互換エンドポイントを取得します。どちらのチームも
相手の仕事の詳細を知らなければなりません。
Modelplane は、現在開発中の初期の v0.1 リリースです。その API と
動作はリリース間で変更される可能性があります。私たちはオープンな環境でそれを構築しています。
統合と機能に関して AI 推論コミュニティと協力します。
プラットフォーム チームが推論クラスターをプロビジョニングし、
利用可能な GPU があれば、開発者は宣言型マニフェストを使用してモデルをデプロイします。
APIバージョン:modelplane.ai/v1alpha1
種類 : モデル展開
メタデータ:
n

ame : qwen-demo
名前空間 : ml-チーム
仕様：
レプリカ : 1
エンジン:
- 名前：クウェン
メンバー：
- 役割 : スタンドアロン
ノードセレクター:
デバイス:
- 名前 : GPU
カウント : 1
セレクター:
- セル: device.capacity["gpu.nvidia.com"].memory.compareTo(quantity("20Gi")) >= 0
テンプレート:
仕様：
コンテナ:
- 名前 : エンジン
画像：vllm/vllm-openai:v0.23.0
引数: ["--model=Qwen/Qwen2.5-0.5B-Instruct"]
Modelplane は、互換性のある無料の GPU を備えたクラスター上にレプリカをスケジュールします。
サービスエンジンをデプロイします。これを 1 つの OpenAI 互換エンドポイントの背後で公開します。
ModelService :
APIバージョン:modelplane.ai/v1alpha1
種類 : ModelService
メタデータ:
名前：クウェン
名前空間 : ml-チーム
仕様：
エンドポイント:
- セレクター:
matchLabels :
modelplane.ai/deployment : qwen-demo
はじめに
スタート ガイドに従って Modelplane をデプロイします。
ローカルの種類のクラスターを作成し、モデルを提供します。仕組みページ
リソースと、モデルをデプロイすると何が起こるかについて説明します。
サンプル マニフェストは、検証済みのエンドツーエンドのレシピであり、
特定のモデル。それぞれが推論クラスとクラスターからの完全なワークフローをカバーします。
モデルのキャッシュ、デプロイメント、サービスを通じて。
Modelplane は、推論の上で独自のクラスター上でコントロール プレーンとして実行されます。
モデルを提供するクラスター。その API は 2 つのリソース セット (ロールごとに 1 つ) であり、
その間のすべてがあなたのために構成されます:
プラットフォーム チームは、InferenceClusters (GPU フリート、によってプロビジョニングされます) を作成します。
Modelplane またはそのままの状態で持ち込まれる）および InferenceClasses（ハードウェア レシピ:
ノード プールが提供するデバイスとそのプロビジョニング方法)、先頭に
推論ゲートウェイ 。
開発者は ModelDeployment (モデルのエンジン、レプリカ数、
およびオプションの ModelCache ) と ModelService (システム全体にわたる 1 つのエンドポイント)
レプリカが選択されます)。
Modelplane はクラスターごとに ModelReplica を構成し、クラスターごとに ModelEndpoint を構成します。
レプリカ

。
これらのリソースが存在すると、Modelplane は 5 つのリソースに対応するフリートを維持します。
懸念事項: クラスターとそのノード プールのプロビジョニング、それぞれのスケジュール設定
ハードウェアが適合するクラスターとプールにレプリカを配置し、レプリカをスケーリングします。
標準の Kubernetes スケール サブリソース、1 つを経由してルーティング
加重カナリア、A/B ロールアウト、キャッシュを備えた OpenAI 互換エンドポイント
モデルの重み付けはクラスターごとに 1 回行われます。
Modelplane はエンジンに関しては偏見を持っていません。 ModelDeployment では、
デプロイメントの形状、ポッドの数、ノード数、デバイスの数。の
作成したエンジン フラグには、並列処理、量子化、および KV 転送が含まれます。
Modelplane はそれらを決して注入しません。これにより、1 つの API であらゆるサービスを提供できるようになります。
コンテナベースのエンジンと任意のトポロジ。
モデルプレーンは v0.1 です。それは初期のものであり、急速に進化しています。計画されている内容については、「機能強化」というラベルが付いた問題を参照してください。
貢献、バグレポート、機能リクエストを歓迎します。
Slack: Kubernetes ワークスペースの #modelplane
セットアップ、チェックの実行、変更の送信方法については、CONTRIBUTING.md を参照してください。
Modelplane は Apache 2.0 ライセンスの下にあります。
Modelplane™ は商標です。 Apache 2.0 ライセンスは商標権を付与しません。
Modelplane の名前とロゴは対象外です。
AI 推論用のオープンソース コントロール プレーン
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
4
フォーク
レポートリポジトリ
リリース
4
モデルプレーン v0.1.0
最新の
2026 年 6 月 23 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The open source control plane for AI inference. Contribute to modelplaneai/modelplane development by creating an account on GitHub.

GitHub - modelplaneai/modelplane: The open source control plane for AI inference · GitHub
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
modelplaneai
/
modelplane
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
663 Commits 663 Commits .github .github apis apis design design docs docs functions functions nix nix schemas schemas .envrc .envrc .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md crossplane-project.yaml crossplane-project.yaml flake.lock flake.lock flake.nix flake.nix nix.sh nix.sh pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Modelplane is software you install and run in your own environment to orchestrate
models, the serving stack, and the infrastructure underneath across cloud,
neocloud, and on-premise. It runs any model on any engine on any infrastructure,
from a single GPU to disaggregated, multi-node deployments. Built on
Crossplane , it is an active system that continuously reconciles your fleet
toward the state you declare: provisioning inference clusters, scheduling
deployments onto compatible clusters, scaling replicas, caching weights, and
routing traffic.
Platform teams provision clusters and publish hardware classes. Developers
declare a model and get back a unified, OpenAI-compatible endpoint. Neither team
has to know the details of the other's job.
Modelplane is an early v0.1 release under active development. Its APIs and
behavior can change between releases. We are building it in the open,
collaborating with the AI inference community on integrations and capabilities.
Once a platform team has provisioned inference clusters and declared the
available GPUs, a developer deploys a model with a declarative manifest:
apiVersion : modelplane.ai/v1alpha1
kind : ModelDeployment
metadata :
name : qwen-demo
namespace : ml-team
spec :
replicas : 1
engines :
- name : qwen
members :
- role : Standalone
nodeSelector :
devices :
- name : gpu
count : 1
selectors :
- cel : device.capacity["gpu.nvidia.com"].memory.compareTo(quantity("20Gi")) >= 0
template :
spec :
containers :
- name : engine
image : vllm/vllm-openai:v0.23.0
args : ["--model=Qwen/Qwen2.5-0.5B-Instruct"]
Modelplane schedules the replica onto a cluster with free, compatible GPUs and
deploys the serving engine. Expose it behind one OpenAI-compatible endpoint with
a ModelService :
apiVersion : modelplane.ai/v1alpha1
kind : ModelService
metadata :
name : qwen
namespace : ml-team
spec :
endpoints :
- selector :
matchLabels :
modelplane.ai/deployment : qwen-demo
Getting started
Follow the getting started guide to deploy Modelplane on a
local kind cluster and serve a model. The how it works page
covers the resources and what happens when you deploy a model.
The example manifests are validated, end-to-end recipes that serve
specific models, each covering the full workflow from inference class and cluster
through model cache, deployment, and service.
Modelplane runs as a control plane on its own cluster, above the inference
clusters that serve models. Its API is two sets of resources, one per role, with
everything in between composed for you:
Platform teams create InferenceClusters (the GPU fleet, provisioned by
Modelplane or brought as-is) and InferenceClasses (hardware recipes: the
devices a node pool offers and how to provision it), fronted by an
InferenceGateway .
Developers create a ModelDeployment (a model's engines, replica count,
and an optional ModelCache ) and a ModelService (one endpoint across the
replicas it selects).
Modelplane composes a ModelReplica per cluster and a ModelEndpoint per
replica.
Once those resources exist, Modelplane keeps the fleet matching them across five
concerns: provisioning clusters and their node pools, scheduling each
replica onto a cluster and pool whose hardware fits, scaling replicas through
the standard Kubernetes scale subresource, routing through one
OpenAI-compatible endpoint with weighted canary and A/B rollouts, and caching
model weights once per cluster.
Modelplane is unopinionated about the engine. A ModelDeployment describes the
shape of a deployment, how many pods, on how many nodes, with which devices; the
engine flags you write carry parallelism, quantization, and KV transfer.
Modelplane never injects them. That is what lets one API serve any
container-based engine and any topology.
Modelplane is at v0.1. It is early and evolving fast. See issues labeled enhancement for what's planned.
Contributions, bug reports, and feature requests are welcome.
Slack: #modelplane in the Kubernetes workspace
See CONTRIBUTING.md for how to get set up, run checks, and submit changes.
Modelplane is under the Apache 2.0 license .
Modelplane™ is a trademark. The Apache 2.0 license grants no trademark rights:
the Modelplane name and logos are not covered by it.
The open source control plane for AI inference
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
4
forks
Report repository
Releases
4
Modelplane v0.1.0
Latest
Jun 23, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

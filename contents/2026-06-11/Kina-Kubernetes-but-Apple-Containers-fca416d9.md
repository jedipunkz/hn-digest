---
source: "https://github.com/vinnie357/kina"
hn_url: "https://news.ycombinator.com/item?id=48495182"
title: "Kina, Kubernetes but Apple Containers"
article_title: "GitHub - vinnie357/kina: kubernetes in apple container based on kind · GitHub"
author: "cohix"
captured_at: "2026-06-11T21:46:43Z"
capture_tool: "hn-digest"
hn_id: 48495182
score: 6
comments: 0
posted_at: "2026-06-11T19:22:50Z"
tags:
  - hacker-news
  - translated
---

# Kina, Kubernetes but Apple Containers

- HN: [48495182](https://news.ycombinator.com/item?id=48495182)
- Source: [github.com](https://github.com/vinnie357/kina)
- Score: 6
- Comments: 0
- Posted: 2026-06-11T19:22:50Z

## Translation

タイトル: Kina、Kubernetes でも Apple コンテナ
記事タイトル: GitHub - vinnie357/kina: 種類に基づく Apple コンテナーの kubernetes · GitHub
説明: 種類に基づいた Apple コンテナ内の Kubernetes。 GitHub でアカウントを作成して、vinnie357/kina の開発に貢献してください。

記事本文:
GitHub - vinnie357/kina: 種類に基づく Apple コンテナーの kubernetes · GitHub
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
ビニー357
/
キナ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

その他のアクションメニューを開く フォルダーとファイル
36 コミット 36 コミット .bees .bees .github .github docs docs kina-cli kina-cli scripts scripts .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Cargo.toml Cargo.toml README.md README.md Clippy.toml Clippy.toml misse.toml misse.toml richfmt.toml richfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
kina - Apple コンテナの Kubernetes
kina は、Apple Container テクノロジーを使用してローカル Kubernetes クラスターを実行するための Rust CLI ツールです。これは kind (Docker の Kubernetes) と同様の機能を提供しますが、macOS システム用に最適化されており、ネイティブの Apple Container テクノロジーを活用してパフォーマンスと統合を向上させています。
🏗️ ネイティブ Apple コンテナ統合 - macOS コンテナテクノロジーを活用して最適なパフォーマンスを実現
☸️ Kubernetes API の互換性 - kubectl 統合による完全な Kubernetes クラスター機能
🌐 CNI プラグインのサポート - コンテナー ネットワーキング用に PTP (デフォルト) と Cilium のどちらかを選択します
🔧 Nginx Ingress コントローラー - nginx-ingress のインストールと構成の組み込みサポート
⚙️ 柔軟な構成 - 適切なデフォルトを備えた TOML ベースの構成
📋 包括的な CLI - クラスターの管理と操作のための豊富なコマンドセット
🚀 開発準備完了 - タスク自動化を備えた統合開発ワークフロー
macOS : 26+ (macOS 15.6 は制限付きで動作する可能性があります)
Apple Container : 1.0.0+ (起動時に自動検出および検証)
Rust : 1.70+ (ソースからビルドする場合)
kina が動作するには Apple コンテナが必要です。まずインストールしてください:
オプション A — 自作 (推奨):
醸造インストールコンテナ
オプション B — 署名付き pkg インストーラー:
ダウンロード : Apple Container Releases から最新の署名付き .pkg を入手します。
インストール: .pkg ファイルをダブルクリックし、次の手順を実行します。

インストーラーのプロンプトが表示される
いずれかの方法でインストールした後:
Start Service : コンテナ システムの起動を実行して API サーバーを起動します
検証:container --version を使用してインストールを確認します。
注: kina には Apple Container 1.0.0 以降が必要です。バージョンは、kina の起動時に自動的に検出され、検証されます。 kina と Apple Container のバージョンを確認するには、kina (引数なし) を実行します。
kubectl - Kubernetes コマンドライン ツール
kubectx & kubens - コンテキストと名前空間の管理 (オプション、mise によって管理)
misse - タスク自動化を備えた開発環境マネージャー
オプション 1: mise (GitHub バックエンド)
misse がインストールされている場合は、リポジトリのクローンを作成せずに、GitHub リリースから直接 kina をインストールできます。
# misse github バックエンド経由で最新の 0.1.0 リリースをインストールします
github:vinnie357/kina@0.1.0 を使用してください。
# または、プロジェクトの misse.toml に追加します。
# [ツール]
# "github:vinnie357/kina" = "0.1.0"
mise github バックエンドは、macOS (Apple Silicon) 上で aarch64-apple-darwin バイナリを自動選択します。 asset_pattern オプションは必要ありません。
注：mise install github:vinnie357/kina@0.1.0 は、kina CLI バイナリのみをインストールします。クラスターを作成する前に、Apple Container ランタイム (macOS 26+、brew install container) が別途必要です。
# リポジトリのクローンを作成します
git クローン https://github.com/vinnie357/kina.git
CDキナ
# Cargo を使用してインストールする
カーゴインストール --path kina-cli
# または、mise を使用します (インストールされている場合)
キナを実行してください:インストール
オプション 3: mise を使用した開発セットアップ
# リポジトリのクローンを作成します
git クローン https://github.com/vinnie357/kina.git
CDキナ
# 開発環境をセットアップする (Rust、ツール、依存関係をインストールする)
セットアップを実行してください:dev
# ビルドしてインストールする
インストールを実行してください
検証
# インストールの確認 (kina および Apple Container のバージョンを表示)
キナ
# Apple コンテナの可用性を確認する (必須、1.0.0 以降)
コンテナを実行してください: # をチェックしてください

ミセを利用する場合
# または手動で確認します:
コンテナ --バージョン
コンテナシステムの開始 # サービスが実行されていない場合は開始します
# オプション: Kubernetes ツールを確認します
misse run k8s:check # misse を使用する場合
kubectl バージョン --client
⚠️ 重要: クラスターを作成する前に、Apple Container 1.0.0 以降が利用可能である必要があります。 kina は起動時にバージョンを自動検出して検証します。 kina status を実行して、Apple Container のバージョン情報を確認します。
kina create kina-test --workers 2 --cni cilium --kernel-path < パス >
kina install nginx-ingress --cluster kina-test
kina インストールデモアプリ --cluster kina-test
キナ検証 キナテスト
--kernel-path は、kina-8 が自動ダウンロードされるまで必要です。
# デフォルト設定でクラスターを作成する
私のクラスターを作成します
# kubectl に接続するために kubeconfig をエクスポートする
kina export my-cluster --format kubeconfig --output ~ /.kube/my-cluster
エクスポート KUBECONFIG= ~ /.kube/my-cluster
# クラスターが動作していることを確認する
kubectl ノードを取得する
詳細オプション:
# Cilium CNI でクラスターを作成し、準備が完了するまで待ちます
kina デモを作成 --cni cilium --wait 300
Nginx Ingress コントローラーとデモ アプリをインストールする
# nginx-ingress をインストールします (バイナリに埋め込まれたマニフェスト — 任意のディレクトリから動作します)
kina install nginx-ingress --cluster my-cluster
# デモアプリケーションをインストールする
kina install デモアプリ --cluster my-cluster
# クラスターをエンドツーエンドで検証する
私のクラスターを確認してください
クラスターのステータスを確認する
# 基本ステータス
kina ステータス my-cluster
# ポッドとサービスの詳細なステータス
kina status my-cluster --verbose
統合テストクラスタ
オプション A: misse を使用する (インストールされている場合)
# Ingress とデモ アプリを使用して統合テスト クラスターを作成する
テストの実行:クラスタ
# 最新のテスト クラスターを検証する
テストを実行してください:クラスター:検証
# すべてのテスト クラスターをクリーンアップします (「demo-」プレフィックスが付いたクラスターを削除します)
テストを実行してください:クラスター:クリーンアップ
オプション B: 手動

セットアップ（ミセなし）
# nginx-ingress でクラスターを作成
kina create Demon-cluster --wait 300
nginx-ingress --clusterデモ-クラスターをインストールします
# ステータスを確認する
kina status デモ-cluster --verbose
デモ クラスターのセットアップでは、以下が作成されます。
タイムスタンプ付きクラスター (例: Demon-20241228-143022 )
nginx-ingress コントローラーのインストールと構成
2 つのレプリカを持つサンプル Web アプリケーション
ブラウザ/カールアクセスのイングレスルーティング
Apple Container ネットワーキングのセットアップを完了する
最初のクラスターを作成した後、すべてが機能することを確認します。
# クラスターのステータスを確認する
kina ステータス my-cluster
# すべてのポッドをリストします (実行ステータスが表示されるはずです)
kubectl --kubeconfig ~ /.kube/my-cluster ポッドを取得 -A
# ノードの準備ができていることを確認する
kubectl --kubeconfig ~ /.kube/my-cluster ノードを取得
トラブルシューティング : クラスターの作成が失敗した場合は、以下を確認してください。
Apple Container CLI が利用可能です:container --version
十分なシステム リソース (2GB 以上の RAM を推奨)
デバッグするには --retain フラグを付けて試してください: kina create test-cluster --retain
# 新しいクラスターを作成する
[名前] [オプション] を作成してください
--image TEXT コンテナイメージ (デフォルト: kindest/node:v1.31.0)
--config FILE クラスター構成ファイル
--wait SECONDS クラスターの準備が完了するまで待機します。
--retain 障害時にクラスターを保持する
--cni ptp | cilium CNI プラグイン (デフォルト: ptp)
# クラスターを削除する
[名前]を削除してください
kina delete --all # すべてのクラスターを削除します
# クラスターをリストする
キナリスト # 簡易リスト
kina list --verbose # 詳細情報
# クラスターのステータスを表示
キナステータス [名前] [オプション]
--verbose 詳細情報を表示
--出力テーブル |ヤムル | json
リソースの操作
# クラスター情報の取得
クラスター [名前] を取得します
kubeconfig [名前] を取得します
ノード [名前] を取得します
# コンテナイメージをロードする
キナロードイメージ --クラスター名
# 設定をエクスポートする
キナエクスポート [名前] [オプション]
--format kubeconfig |構成
--出力ファイル
アドオン管理

# アドオンをインストールする
kina install nginx-ingress --cluster NAME
kina install ingress-nginx --cluster NAME
kina install cni --cluster NAME
kina install metrics-server --cluster NAME
クラスター操作
# kubelet 証明書署名リクエストを承認する
キナ承認-csr [名前]
# 構成管理
キナ構成ショー
キナ構成セットのキー値
キナ設定でKEYを取得
キナ構成リセット
キナ構成パス
構成
kina は、次の場所にある TOML 構成ファイルを使用します。
~/.config/kina/config.toml
デフォルト設定
【クラスター】
デフォルト名 = " キナ "
default_image = " kindest/node:v1.31.0 "
デフォルト待機タイムアウト = 300
data_dir = " ~/.local/share/kina "
失敗時の保持 = false
デフォルト_cni = " ptp "
[ アップルコンテナ ]
cli_path = null # 自動検出
[ apple_container .ランタイム設定]
cpu_limit = null
メモリ制限 = " 2Gi "
storage_limit = " 20Gi "
[ apple_container .ネットワーク]
network_name = " キナ "
Enable_ipv6 = false
dns_servers = []
[ Kubernetes ]
デフォルトバージョン = " v1.28.0 "
kubectl_path = null # 自動検出
default_namespace = " デフォルト "
kubeconfig_dir = " ~/.config/kina/kubeconfig "
[ロギング]
レベル = " 情報 "
フォーマット = " テキスト "
ファイルロギング = false
log_dir = null
環境変数
# 設定
import KINA_CONFIG_DIR= " $HOME /.config/kina "
import KINA_DATA_DIR= " $HOME /.local/share/kina "
# ロギング
エクスポート RUST_LOG= " 情報 "
エクスポート RUST_BACKTRACE= " 1 "
Apple コンテナの統合
kina は、Kubernetes ノードを実行するために Apple Container テクノロジーを活用しています。
ネイティブ統合: コンテナのライフサイクルに Apple Container CLI を使用
リソース制限 : 構成可能な CPU、メモリ、およびストレージ制限
ネットワーク統合 : macOS ネットワーキングとのシームレスな統合
DNS サポート : クラスターアクセスのための自動 DNS 構成
┌───────

─────────────┐
│ macOS ホスト │
│ ┌─────────────────────┐ │
│ │ Apple コンテナ VM │ │
│ │ ┌─────────────────┐ │ │
│ │ │ Kubernetes ノード │ │ │
│ │ │ • kubelet │ │ │
│ │ │ • コンテナド │ │ │
│ │ │ • CNI (PTP/繊毛) │ │ │
│ │ lux─────────────────┘ │ │
│ ━━━━━━━━━━━━━━━━━━┘ │
━━━━━━━━━━━━━━━━━━━━━━┘
CNI サポート
互換性 : Apple コンテナ用に最適化
シンプルさ : ホストローカル IPAM を使用したポイントツーポイント ネットワーキング
パフォーマンス : 単一ノードクラスターの最小限のオーバーヘッド
高度な機能: eBPF ベースのネットワーキングとセキュリティ
要件 : 互換性のあるカーネル モジュール
ユースケース: 複雑なネットワーク要件と可観測性
Cilium full-eBPF モードには、kina がビルドするカスタム Linux カーネル (6.18.5-kina.1) が必要です。 kina create ... --cni cilium --kernel-path <path/to/vmlinux> を使用してクラスターごとに有効にします。ストック カーネル (PTP) クラスターには追加のものは何も必要ありません。ビルド手順、使用法、配布、およびホストカーネルの注意事項については、docs/development/custom-kernel.md を参照してください。 (計画されている将来のリリースでは、カーネルが自動的にフェッチされるため、--kernel-path は必要ありません。現在サポートされているパスはこれだけです。)
# 特定の CNI を使用してクラスターを作成する
キナ作成

テスト-ptp --c

[切り捨てられた]

## Original Extract

kubernetes in apple container based on kind. Contribute to vinnie357/kina development by creating an account on GitHub.

GitHub - vinnie357/kina: kubernetes in apple container based on kind · GitHub
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
vinnie357
/
kina
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
36 Commits 36 Commits .bees .bees .github .github docs docs kina-cli kina-cli scripts scripts .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Cargo.toml Cargo.toml README.md README.md clippy.toml clippy.toml mise.toml mise.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
kina - Kubernetes in Apple Container
kina is a Rust CLI tool for running local Kubernetes clusters using Apple Container technology. It provides similar functionality to kind (Kubernetes in Docker) but is optimized for macOS systems, leveraging native Apple Container technology for improved performance and integration.
🏗️ Native Apple Container Integration - Leverage macOS container technology for optimal performance
☸️ Kubernetes API Compatibility - Full Kubernetes cluster functionality with kubectl integration
🌐 CNI Plugin Support - Choose between PTP (default) and Cilium for container networking
🔧 Nginx Ingress Controller - Built-in support for nginx-ingress installation and configuration
⚙️ Flexible Configuration - TOML-based configuration with sensible defaults
📋 Comprehensive CLI - Rich command set for cluster management and operations
🚀 Development Ready - Integrated development workflow with mise task automation
macOS : 26+ (macOS 15.6 may work with limitations)
Apple Container : 1.0.0+ (auto-detected and validated at startup)
Rust : 1.70+ (for building from source)
Apple Container is required for kina to work. Install it first:
Option A — Homebrew (recommended):
brew install container
Option B — Signed pkg installer:
Download : Get the latest signed .pkg from Apple Container Releases
Install : Double-click the .pkg file and follow the installer prompts
After installing via either method:
Start Service : Run container system start to start the API server
Verify : Check installation with container --version
Note : kina requires Apple Container 1.0.0 or later . The version is automatically detected and validated when kina starts. Run kina (no arguments) to see your kina and Apple Container versions.
kubectl - Kubernetes command-line tool
kubectx & kubens - Context and namespace management (optional, managed by mise)
mise - Development environment manager with task automation
Option 1: mise (GitHub backend)
If you have mise installed, you can install kina directly from GitHub Releases without cloning the repository:
# Install the latest 0.1.0 release via mise github backend
mise use github:vinnie357/kina@0.1.0
# Or add to your project's mise.toml:
# [tools]
# "github:vinnie357/kina" = "0.1.0"
The mise github backend auto-selects the aarch64-apple-darwin binary on macOS (Apple Silicon). No asset_pattern option is needed.
Note : mise install github:vinnie357/kina@0.1.0 installs the kina CLI binary only. The Apple Container runtime (macOS 26+, brew install container ) is required separately before creating clusters.
# Clone the repository
git clone https://github.com/vinnie357/kina.git
cd kina
# Install using Cargo
cargo install --path kina-cli
# OR using mise (if installed)
mise run kina:install
Option 3: Development Setup with mise
# Clone the repository
git clone https://github.com/vinnie357/kina.git
cd kina
# Set up development environment (installs Rust, tools, and dependencies)
mise run setup:dev
# Build and install
mise run install
Verification
# Verify installation (shows kina and Apple Container versions)
kina
# Check Apple Container availability (REQUIRED, 1.0.0+)
mise run container:check # If using mise
# OR manually check:
container --version
container system start # Start the service if not running
# Optional: Check Kubernetes tools
mise run k8s:check # If using mise
kubectl version --client
⚠️ Important : Apple Container 1.0.0+ must be available before creating clusters. kina auto-detects and validates the version at startup. Run kina status to see Apple Container version information.
kina create kina-test --workers 2 --cni cilium --kernel-path < path >
kina install nginx-ingress --cluster kina-test
kina install demo-app --cluster kina-test
kina verify kina-test
--kernel-path is required until kina-8 lands auto-download.
# Create a cluster with default settings
kina create my-cluster
# Export kubeconfig to connect with kubectl
kina export my-cluster --format kubeconfig --output ~ /.kube/my-cluster
export KUBECONFIG= ~ /.kube/my-cluster
# Verify cluster is working
kubectl get nodes
Advanced Options:
# Create cluster with Cilium CNI and wait for readiness
kina create demo --cni cilium --wait 300
Install Nginx Ingress Controller and Demo App
# Install nginx-ingress (manifests embedded in binary — works from any directory)
kina install nginx-ingress --cluster my-cluster
# Install demo application
kina install demo-app --cluster my-cluster
# Verify the cluster end-to-end
kina verify my-cluster
Check Cluster Status
# Basic status
kina status my-cluster
# Detailed status with pods and services
kina status my-cluster --verbose
Integration Test Cluster
Option A: Using mise (if installed)
# Create an integration test cluster with ingress and demo app
mise run test:cluster
# Validate the most recent test cluster
mise run test:cluster:validate
# Clean up all test clusters (removes clusters with 'demo-' prefix)
mise run test:cluster:cleanup
Option B: Manual setup (without mise)
# Create cluster with nginx-ingress
kina create demo-cluster --wait 300
kina install nginx-ingress --cluster demo-cluster
# Check status
kina status demo-cluster --verbose
The demo cluster setup creates:
A timestamped cluster (e.g., demo-20241228-143022 )
nginx-ingress controller installation and configuration
A sample web application with 2 replicas
Ingress routing for browser/curl access
Complete Apple Container networking setup
After creating your first cluster, verify everything works:
# Check cluster status
kina status my-cluster
# List all pods (should show running status)
kubectl --kubeconfig ~ /.kube/my-cluster get pods -A
# Verify nodes are ready
kubectl --kubeconfig ~ /.kube/my-cluster get nodes
Troubleshooting : If cluster creation fails, check:
Apple Container CLI is available: container --version
Sufficient system resources (2GB+ RAM recommended)
Try with --retain flag to debug: kina create test-cluster --retain
# Create a new cluster
kina create [NAME] [OPTIONS]
--image TEXT Container image (default: kindest/node:v1.31.0)
--config FILE Cluster configuration file
--wait SECONDS Wait for cluster readiness
--retain Retain cluster on failure
--cni ptp | cilium CNI plugin (default: ptp)
# Delete a cluster
kina delete [NAME]
kina delete --all # Delete all clusters
# List clusters
kina list # Simple list
kina list --verbose # Detailed information
# Show cluster status
kina status [NAME] [OPTIONS]
--verbose Show detailed information
--output table | yaml | json
Resource Operations
# Get cluster information
kina get clusters [NAME]
kina get kubeconfig [NAME]
kina get nodes [NAME]
# Load container images
kina load IMAGE --cluster NAME
# Export configurations
kina export [NAME] [OPTIONS]
--format kubeconfig | config
--output FILE
Addon Management
# Install addons
kina install nginx-ingress --cluster NAME
kina install ingress-nginx --cluster NAME
kina install cni --cluster NAME
kina install metrics-server --cluster NAME
Cluster Operations
# Approve kubelet Certificate Signing Requests
kina approve-csr [NAME]
# Configuration management
kina config show
kina config set KEY VALUE
kina config get KEY
kina config reset
kina config path
Configuration
kina uses TOML configuration files located at:
~/.config/kina/config.toml
Default Configuration
[ cluster ]
default_name = " kina "
default_image = " kindest/node:v1.31.0 "
default_wait_timeout = 300
data_dir = " ~/.local/share/kina "
retain_on_failure = false
default_cni = " ptp "
[ apple_container ]
cli_path = n ull # Auto-detected
[ apple_container . runtime_config ]
cpu_limit = n ull
memory_limit = " 2Gi "
storage_limit = " 20Gi "
[ apple_container . network ]
network_name = " kina "
enable_ipv6 = false
dns_servers = []
[ kubernetes ]
default_version = " v1.28.0 "
kubectl_path = n ull # Auto-detected
default_namespace = " default "
kubeconfig_dir = " ~/.config/kina/kubeconfig "
[ logging ]
level = " info "
format = " text "
file_logging = false
log_dir = n ull
Environment Variables
# Configuration
export KINA_CONFIG_DIR= " $HOME /.config/kina "
export KINA_DATA_DIR= " $HOME /.local/share/kina "
# Logging
export RUST_LOG= " info "
export RUST_BACKTRACE= " 1 "
Apple Container Integration
kina leverages Apple Container technology for running Kubernetes nodes:
Native Integration : Uses Apple Container CLI for container lifecycle
Resource Limits : Configurable CPU, memory, and storage limits
Network Integration : Seamless integration with macOS networking
DNS Support : Automatic DNS configuration for cluster access
┌─────────────────────────────────────────┐
│ macOS Host │
│ ┌─────────────────────────────────────┐ │
│ │ Apple Container VM │ │
│ │ ┌─────────────────────────────────┐ │ │
│ │ │ Kubernetes Node │ │ │
│ │ │ • kubelet │ │ │
│ │ │ • containerd │ │ │
│ │ │ • CNI (PTP/Cilium) │ │ │
│ │ └─────────────────────────────────┘ │ │
│ └─────────────────────────────────────┘ │
└─────────────────────────────────────────┘
CNI Support
Compatibility : Optimized for Apple Container
Simplicity : Point-to-point networking with host-local IPAM
Performance : Minimal overhead for single-node clusters
Advanced Features : eBPF-based networking and security
Requirements : Compatible kernel modules
Use Cases : Complex networking requirements and observability
Cilium full-eBPF mode requires a custom Linux kernel that kina builds ( 6.18.5-kina.1 ); enable it per cluster with kina create ... --cni cilium --kernel-path <path/to/vmlinux> . Stock kernel (PTP) clusters need nothing extra. See docs/development/custom-kernel.md for build instructions, usage, distribution, and host-kernel gotchas. (A planned future release will fetch the kernel automatically — a one-time ~32 MB download — so --kernel-path will not be required; today it is the only supported path.)
# Create cluster with specific CNI
kina create test-ptp --c

[truncated]

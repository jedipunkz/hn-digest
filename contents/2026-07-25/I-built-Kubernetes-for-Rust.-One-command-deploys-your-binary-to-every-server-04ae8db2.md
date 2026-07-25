---
source: "https://github.com/Emran-goat/sortie"
hn_url: "https://news.ycombinator.com/item?id=49050420"
title: "I built Kubernetes for Rust. One command deploys your binary to every server"
article_title: "GitHub - Emran-goat/sortie: Kubernetes for Rust. Deploy binaries to Linux servers over SSH. · GitHub"
author: "EmranAbdu"
captured_at: "2026-07-25T18:52:54Z"
capture_tool: "hn-digest"
hn_id: 49050420
score: 1
comments: 0
posted_at: "2026-07-25T18:49:04Z"
tags:
  - hacker-news
  - translated
---

# I built Kubernetes for Rust. One command deploys your binary to every server

- HN: [49050420](https://news.ycombinator.com/item?id=49050420)
- Source: [github.com](https://github.com/Emran-goat/sortie)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T18:49:04Z

## Translation

タイトル: Rust 用に Kubernetes を構築しました。 1 つのコマンドでバイナリをすべてのサーバーにデプロイします
記事タイトル: GitHub - Emran-goat/sortie: Kubernetes for Rust。 SSH 経由で Linux サーバーにバイナリをデプロイします。 · GitHub
説明: Rust 用の Kubernetes。 SSH 経由で Linux サーバーにバイナリをデプロイします。 - エムランヤギ/出撃

記事本文:
GitHub - Emran-goat/sortie: Rust 用 Kubernetes。 SSH 経由で Linux サーバーにバイナリをデプロイします。 · GitHub
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
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
エムランヤギ
/
出撃する
公共
通知
通知設定を変更するにはサインインする必要があります
アディティ

ナビゲーションオプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .github/ workflows .github/ workflows CHANGELOG CHANGELOG LICENSES LICENSES build build cmd cmd docs docs 例 例 ハック ハック ロゴ ロゴ pkg pkg テスト テスト サードパーティ サードパーティ .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス Makefile Makefile OWNERS OWNERS OWNERS_ALIASES OWNERS_ALIASES README.md README.md SECURITY.md SECURITY.md SECURITY_CONTACTS SECURITY_CONTACTS SUPPORT.md SUPPORT.md Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rust用のKubernetes。
1 つの CLI からサーバーのクラスター全体に Rust サービスをデプロイ、管理、監視します。
能力
何をするのか
マルチホストクラスター
1 台だけではなく、複数のサーバーにデプロイする
ローリングアップデート
ヘルスゲートを使用してホストを一度に 1 つずつ更新
カナリア＆ブルーグリーン
サブセットにデプロイするか、完全な並列スタックをスピンします
組み込みリバースプロキシ
ホストベースのルーティングを備えた組み込みの HTTP プロキシ。 nginxは必要ありません。
宣言状態
sortie.toml でターゲットを宣言します。出撃 調整を適用する
自動ロールバック
ヘルスチェックに失敗すると自動ロールバックがトリガーされる
サービスレジストリ
サービスの登録、名前による解決、イングレス構成の生成
機密管理
サーバー上の暗号化されたキーと値のストア (SSH が境界)
オートスケーリング
CPU ベースのスケーリング ループ、systemd インスタンスを調整します
自己修復
sortie-agent watchdog が停止したサービスを再起動する
kubectl スタイル CLI
apply、get、describe、logs、rollback、scale
マルチクラスタフェデレーション
1 つのコマンドですべてのターゲットにデプロイ
TLSプロビジョニング
certbot による自動証明書 (Let's Encrypt)

可観測性
ホストごとのメトリクス (CPU、メモリ) とログ集計
環境設定
ターゲットごとのサービス構成用のインライン環境変数
クイックスタート
貨物設置出撃
出撃初期
サーバーの詳細を使用して sortie.toml を編集し、以下をデプロイします。
出撃 生産を適用する
これにより、バイナリが構築され、SCP 経由ですべてのホストにアップロードされ、systemd サービスがインストールされて開始され、ヘルス エンドポイントが応答するのを待ちます。ヘルスチェックが失敗した場合、そのホストはロールバックします。
デプロイしないと何が変わるかを確認するには:
出撃 生産を適用 --check
埋め込みプロキシ
nginx、キャディ、セットアップは必要ありません。プロキシをサーバーにインストールすると、サービス レジストリが読み取られてリクエストがルーティングされます。
出撃 SVC レジスタ生産 my-api 8080
出撃プロキシのインストール本番 --ポート 80
プロキシは、Host ヘッダーを登録されたサービス名と照合します。サーバー IP が 1.2.3.4 の場合:
カール -H "ホスト: my-api" http://1.2.3.4/
ドメインは必要ありません。登録したサービス名は、クライアントが送信するホスト名です。任意の文字列を使用できます。
TLS を追加するには、最初に実際のドメインで sortie tls セットアップを実行し (Let's Encrypt)、次に DNS をサーバーに向けます。
完全に展開する前に、ホストのサブセットにロールアウトします。
出撃 生産を適用 --canary 20
または、完全な並列スタックをデプロイして反転します。
出撃適用生産 --青緑
すべてのターゲットに展開する
出撃全員配備
sortie.toml 内のすべてのターゲットを反復します。 --canary および --blue-green でも動作します。
貨物設置出撃
またはソースからビルドします。
git clone https://github.com/Emran-goat/sortie.git
CD出撃
カーゴビルド --release
Rust 1.70 以降、SSH + systemd、キーベースの認証を備えた Linux サーバーが必要です。
コマンド
kubectl のように
アクション
出撃初期
sortie.tomlを作成する
出撃適用 <ターゲット>
kubectl 適用
すべてのホストにわたるローリング デプロイ。ドライランの場合は --check、戦略の場合は --canary / --blue-green
出撃配備<ターゲット>
適用と同じ

、フェデレーションの場合は --all を使用
出撃ゲット
kubectl 取得
すべてのターゲットをバージョン、タイムスタンプ、サービスとともに表示します
出撃説明<ターゲット>
kubectl 説明
ターゲット構成 + ホストごとのステータス
出撃ログ <ターゲット> [ホスト]
kubectl ログ
ホストからサービスログを取得する
出撃ロールバック <ターゲット>
kubectl ロールアウトを元に戻す
以前のバイナリに戻す
出撃ステータス<ターゲット>
kubectl ポッドを取得する
ホストごとのサービスステータスを確認する
出撃体力<ターゲット>
ホストの接続とサービスの健全性を確認する
ソーティ SVC レジスタ <ターゲット> <名前> <ポート>
クラスターにサービスを登録する
出撃SVCリスト <ターゲット>
登録されているすべてのサービスを一覧表示する
ソーティ SVC 解決 <ターゲット> <名前>
サービス名をホスト:ポートに解決します。
sortie svc restart <ホスト> <ターゲット> <名前>
kubectl ロールアウトの再起動
ホスト上のサービスを再起動する
sortie svc stop <ホスト> <ターゲット> <名前>
ホスト上のサービスを停止する
sortie svc start <ホスト> <ターゲット> <名前>
ホスト上でサービスを開始する
出撃プロキシ インストール <ターゲット>
組み込みリバース プロキシを systemd サービスとしてインストールする
出撃スケール <ターゲット> <n>
kubectl スケール
ホストごとの systemd インスタンス数を設定する
出撃オートスケール <ターゲット>
CPU ベースの自動スケーリング ループを開始する
出撃シークレットセット <ターゲット> <キー> <値>
秘密ファイルをサーバーに保存する
出撃シークレット <ターゲット> <キー> を取得
保存されたシークレットを読み取る
出撃秘密 rm <ターゲット> <キー>
保存されたシークレットを削除する
sortie tls setup <ターゲット> <ドメイン> <電子メール>
certbot 経由で Let's Encrypt 証明書をプロビジョニングする
出撃 進入 <ターゲット>
サービスレジストリからnginx設定を生成する
出撃指標 <ターゲット>
ホストごとの CPU とメモリを表示する
構成
[ターゲット。製作】
ホスト = [ " 10.0.0.1 " 、 " 10.0.0.2 " ]
ポート = 22
ユーザー = " デプロイ "
key_path = " ~/.ssh/id_rsa "
target_triple = " x86_64-unknown-linux-gnu "
デプロイパス = " /opt/myapp "
health_check_url = " http://localhost:8080/health "
health_check_timeout_secs = 30
[ターゲット。プロデュース

アクション。サービス】
名前 = " マイアプリ "
再起動 = "常に"
[ターゲット。生産 。環境]
DATABASE_URL = " postgres://localhost:5432/myapp "
RUST_LOG = " 情報 "
構成リファレンス
フィールド
必須
デフォルト
説明
ホスト
はい
クラスター内のサーバーアドレス
ポート
いいえ
22
SSHポート
ユーザー
はい
SSHユーザー
キーパス
いいえ
~/.ssh/id_rsa
SSHキーのパス
ターゲットトリプル
はい
Rustターゲットトリプル
デプロイパス
はい
バイナリ用のサーバーディレクトリ
health_check_url
いいえ
デプロイ後に確認するURL
health_check_timeout_secs
いいえ
30
ヘルスを最大待機する
インスタンス
いいえ
1
ホストごとに予想されるレプリカ数
サービス名
はい
systemd サービス名
サービス.再起動
いいえ
いつも
再起動ポリシー
環境*
いいえ
環境変数
仕組み
sortie.toml から目的の状態を読み取ります。バイナリをビルドします。すべてのホストにわたってローリング デプロイメントを実行します。SCP 経由でアップロードし、環境変数を使用して systemd ユニットをインストールし、正常性を確認し、次のホストを実行します。ヘルスチェックが失敗した場合、そのホストはロールバックされ、デプロイが停止します。
以前のバージョンは、即座にロールバックできるように .bak ファイルとして保存されます。クラスターの状態は、各ホストの .sortie/state.json に保存されます。
出撃/
Cargo.toml
カーゴロック
ライセンス
README.md
変更ログ.md
貢献.md
CODE_OF_CONDUCT.md
セキュリティ.md
セキュリティ_コンタクト
サポート.md
所有者
OWNERS_ALIASES
エージェント.md
.gitatributes
.gitignore
錆びたツールチェーン.toml
メイクファイル
cmd/
出撃/
main.rs
エージェント/
main.rs
パッケージ/
lib.rs
クライスラー
タイプ.rs
config.rs
init.rs
build.rs
ssh.rs
クラスター.rs
デプロイ.rs
systemd.rs
ヘルス.rs
ロールバック.rs
プロキシ.rs
ロゴ/
出撃.png
ビルド/
ビルド-リリース.bash
ハック/
ci.bash
テスト/
スモークテスト.rs
ドキュメント/
建築/
画像/
c4-context--0.svg
c4-dynamic-build--0.svg
ビルドライフサイクル--0.svg
c4-containers--0.svg
c4-コンポーネント--0.svg
c4-deployment--0.svg
ライセンス/
README.md
例/
基本的な.toml
マルチホスト.toml
with-env.toml
サードパーティ/
README.md
変更履歴/
v0.1.0.md
v0.2.0.md
C

貢献している
プルリクエストは大歓迎です。 COTRIBUTING.md を参照してください。
Rust用のKubernetes。 SSH 経由で Linux サーバーにバイナリをデプロイします。
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Kubernetes for Rust. Deploy binaries to Linux servers over SSH. - Emran-goat/sortie

GitHub - Emran-goat/sortie: Kubernetes for Rust. Deploy binaries to Linux servers over SSH. · GitHub
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
Emran-goat
/
sortie
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .github/ workflows .github/ workflows CHANGELOG CHANGELOG LICENSES LICENSES build build cmd cmd docs docs examples examples hack hack logo logo pkg pkg tests tests third_party third_party .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile OWNERS OWNERS OWNERS_ALIASES OWNERS_ALIASES README.md README.md SECURITY.md SECURITY.md SECURITY_CONTACTS SECURITY_CONTACTS SUPPORT.md SUPPORT.md rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Kubernetes for Rust.
Deploy, manage, and monitor Rust services across a cluster of servers from one CLI.
Capability
What it does
Multi host clusters
Deploy to a fleet of servers, not just one
Rolling updates
Hosts updated one at a time with health gates
Canary & blue-green
Deploy to a subset, or spin a full parallel stack
Embedded reverse proxy
Built in HTTP proxy with Host based routing. No nginx needed.
Declarative state
Declare targets in sortie.toml; sortie apply reconciles
Auto rollback
Failed health checks trigger automatic rollback
Service registry
Register services, resolve by name, generate ingress config
Secrets management
Encrypted key-value store on servers (SSH is the boundary)
Auto scaling
CPU-based scaling loop, adjusts systemd instances
Self healing
sortie-agent watchdog restarts dead services
kubectl style CLI
apply , get , describe , logs , rollback , scale
Multi cluster federation
Deploy to all targets with one command
TLS provisioning
Auto certs via certbot (Let's Encrypt)
Observability
Per-host metrics (CPU, memory) and log aggregation
Environment config
Inline env vars for per-target service configuration
Quick start
cargo install sortie
sortie init
Edit sortie.toml with your server details, then deploy:
sortie apply production
This builds your binary, uploads it via SCP to every host, installs a systemd service, starts it, and waits for the health endpoint to respond. If a health check fails, that host rolls back.
To see what would change without deploying:
sortie apply production --check
Embedded proxy
No nginx, no caddy, no setup. Install the proxy on your server and it reads the service registry to route requests.
sortie svc register production my-api 8080
sortie proxy install production --port 80
The proxy matches the Host header against registered service names. If your server IP is 1.2.3.4:
curl -H "Host: my-api" http://1.2.3.4/
No domain needed. The service name you registered is the hostname clients send. You can use any string.
To add TLS, run sortie tls setup with a real domain first (Let's Encrypt), then point DNS at the server.
Roll out to a subset of hosts before full deploy:
sortie apply production --canary 20
Or deploy a full parallel stack and flip:
sortie apply production --blue-green
Deploy to all targets
sortie deploy all
Iterates every target in sortie.toml. Also works with --canary and --blue-green.
cargo install sortie
Or build from source:
git clone https://github.com/Emran-goat/sortie.git
cd sortie
cargo build --release
Requires Rust 1.70+, Linux servers with SSH + systemd, key-based auth.
Command
Like kubectl
Action
sortie init
Create a sortie.toml
sortie apply <target>
kubectl apply
Rolling deploy across all hosts. --check for dry-run, --canary / --blue-green for strategies
sortie deploy <target>
Same as apply, with --all for federation
sortie get
kubectl get
Show all targets with version, timestamp, services
sortie describe <target>
kubectl describe
Target config + per-host status
sortie logs <target> [host]
kubectl logs
Fetch service logs from a host
sortie rollback <target>
kubectl rollout undo
Revert to previous binary
sortie status <target>
kubectl get pods
Check service status per host
sortie health <target>
Check host connectivity and service health
sortie svc register <target> <name> <port>
Register a service in the cluster
sortie svc list <target>
List all registered services
sortie svc resolve <target> <name>
Resolve a service name to host:port
sortie svc restart <host> <target> <name>
kubectl rollout restart
Restart a service on a host
sortie svc stop <host> <target> <name>
Stop a service on a host
sortie svc start <host> <target> <name>
Start a service on a host
sortie proxy install <target>
Install embedded reverse proxy as a systemd service
sortie scale <target> <n>
kubectl scale
Set systemd instance count per host
sortie autoscale <target>
Start CPU-based auto-scaling loop
sortie secret set <target> <key> <value>
Store a secret file on the server
sortie secret get <target> <key>
Read a stored secret
sortie secret rm <target> <key>
Delete a stored secret
sortie tls setup <target> <domain> <email>
Provision Let's Encrypt certs via certbot
sortie ingress <target>
Generate nginx config from service registry
sortie metrics <target>
Show CPU and memory per host
Configuration
[ targets . production ]
hosts = [ " 10.0.0.1 " , " 10.0.0.2 " ]
port = 22
user = " deploy "
key_path = " ~/.ssh/id_rsa "
target_triple = " x86_64-unknown-linux-gnu "
deploy_path = " /opt/myapp "
health_check_url = " http://localhost:8080/health "
health_check_timeout_secs = 30
[ targets . production . service ]
name = " myapp "
restart = " always "
[ targets . production . env ]
DATABASE_URL = " postgres://localhost:5432/myapp "
RUST_LOG = " info "
Configuration reference
Field
Required
Default
Description
hosts
yes
Server addresses in the cluster
port
no
22
SSH port
user
yes
SSH user
key_path
no
~/.ssh/id_rsa
SSH key path
target_triple
yes
Rust target triple
deploy_path
yes
Server directory for binary
health_check_url
no
URL to check after deploy
health_check_timeout_secs
no
30
Max wait for health
instances
no
1
Expected replicas per host
service.name
yes
systemd service name
service.restart
no
always
Restart policy
env.*
no
Environment variables
How it works
Reads desired state from sortie.toml . Builds the binary. Performs a rolling deployment across all hosts: upload via SCP, install systemd unit with env vars, verify health, next host. If a health check fails, that host rolls back and the deploy stops.
Previous versions are kept as .bak files for instant rollback. Cluster state is stored in .sortie/state.json on each host.
sortie/
Cargo.toml
Cargo.lock
LICENSE
README.md
CHANGELOG.md
CONTRIBUTING.md
CODE_OF_CONDUCT.md
SECURITY.md
SECURITY_CONTACTS
SUPPORT.md
OWNERS
OWNERS_ALIASES
AGENTS.md
.gitattributes
.gitignore
rust-toolchain.toml
Makefile
cmd/
sortie/
main.rs
agent/
main.rs
pkg/
lib.rs
cli.rs
types.rs
config.rs
init.rs
build.rs
ssh.rs
cluster.rs
deploy.rs
systemd.rs
health.rs
rollback.rs
proxy.rs
logo/
sortie.png
build/
build-release.bash
hack/
ci.bash
tests/
smoke_test.rs
docs/
architecture/
images/
c4-context--0.svg
c4-dynamic-build--0.svg
build-lifecycle--0.svg
c4-containers--0.svg
c4-components--0.svg
c4-deployment--0.svg
LICENSES/
README.md
examples/
basic.toml
multi-host.toml
with-env.toml
third_party/
README.md
CHANGELOG/
v0.1.0.md
v0.2.0.md
Contributing
Pull requests welcome. See CONTRIBUTING.md .
Kubernetes for Rust. Deploy binaries to Linux servers over SSH.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

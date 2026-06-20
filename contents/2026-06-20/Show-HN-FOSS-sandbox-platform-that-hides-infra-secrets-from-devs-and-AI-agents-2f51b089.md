---
source: "https://github.com/octelium/cordium"
hn_url: "https://news.ycombinator.com/item?id=48609062"
title: "Show HN: FOSS sandbox platform that hides infra secrets from devs and AI agents"
article_title: "GitHub - octelium/cordium: Open-source, general-purpose sandbox platform for devs and AI agents that provides identity-based secure access to infrastructure without credentials. · GitHub"
author: "geoctl"
captured_at: "2026-06-20T15:38:29Z"
capture_tool: "hn-digest"
hn_id: 48609062
score: 2
comments: 0
posted_at: "2026-06-20T13:19:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: FOSS sandbox platform that hides infra secrets from devs and AI agents

- HN: [48609062](https://news.ycombinator.com/item?id=48609062)
- Source: [github.com](https://github.com/octelium/cordium)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T13:19:20Z

## Translation

タイトル: Show HN: 開発者や AI エージェントからインフラ秘密を隠す FOSS サンドボックス プラットフォーム
記事のタイトル: GitHub - octelium/cordium: 認証情報なしでインフラストラクチャへの ID ベースの安全なアクセスを提供する、開発者と AI エージェント向けのオープンソースの汎用サンドボックス プラットフォーム。 · GitHub
説明: 認証情報なしでインフラストラクチャへの ID ベースの安全なアクセスを提供する、開発者および AI エージェント向けのオープンソースの汎用サンドボックス プラットフォーム。 - オクテリウム/コージウム
HN テキスト: こんにちは、HN。 Cordium は、FOSS のセルフホスト型 ID ベースの汎用サンドボックス プラットフォームで、私が長年取り組んでいるもので、私の主な仕事である Kubernetes と Octelium 上に構築されています。他の開発環境 (例: GitHub コードスペース) やサンドボックス プラットフォーム (例: E2B、Daytona など) と比較した場合、Cordium の主な違いは、Cordium は資格情報 (例: API キーとアクセス トークン、SSH プライベート) を注入することなく、リソース/インフラストラクチャ (例: API、SSH、データベース、k8s など) への ID ベースの秘密のない安全なアクセスを自動的に提供することです。キー/パスワード、データベース パスワード、mTLS 秘密キーなど) をサンドボックスに送信します。そこでは、サンドボックスの範囲外にある Octelium で保護されたリソースの ID 認識プロキシによって上流の資格情報が保持され、ユーザーが Octelium ポリシーによって承認されている場合、ID 認識プロキシによってオンザフライで注入されます。つまり、Cordium は、リモート開発環境やサンドボックス プラットフォームを置き換えることができる分離された実行環境であるだけでなく、同様にインフラストラクチャ/リソースへの安全なアクセス プラットフォームでもあります。これは基本的にサンドボックス プラットフォーム + ZTNA/リモート アクセス VPN であり、統合 ID 管理、L7 対応アクセス制御、可視性が組み込まれています。

記事本文:
GitHub - octelium/cordium: 認証情報なしでインフラストラクチャへの ID ベースの安全なアクセスを提供する、開発者と AI エージェント向けのオープンソースの汎用サンドボックス プラットフォーム。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
オクテリウム
/
コージウム
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにサインインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
214 コミット 214 コミット .github .github .vscode .vscode apis apis client/cordium client/cordium クラスター クラスター pkg pkg unsorted unsorted .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.work go.work go.work.sum go.work.sum すべてのファイルを表示 リポジトリ ファイルナビゲーション
Cordium は、Kubernetes と Octelium 上に構築された、無料のオープンソース、自己ホスト型の ID ベースのサンドボックス プラットフォームです。 Cordium は、開発者、AI エージェント、自動化されたワークロードに分離された再現可能な分離サンドボックスを提供する汎用プラットフォームです。 Cordium は、コーディング セッション (VSCode、Zed など) 用のリモート開発環境、AI エージェントおよび CI/CD ワークロード用のサンドボックス、開発者と AI エージェント用のインフラストラクチャへの秘密のない安全なリモート アクセスなど、さまざまなユース ケースに使用できます。
リモート開発環境 (例: GitHub Codespaces、Coder など) やサンドボックス プラットフォーム (例: E2B、Daytona など) と比較した場合、Cordium の主な差別化点は、Cordium がサンドボックス内からリソースおよびインフラストラクチャへの ID ベースの秘密なしのアクセスを自動的に提供することです。 Cordium は、Octelium ZTNA 機能を活用して、上流のアプリケーション層認証情報 (API キーとアクセス トークン、SSH 秘密キー、データベース パスワードなど) を公開、配布、管理することなく、サンドボックス内からインフラストラクチャ リソース (SSH サーバー、データベース、内部 HTTP API、mTLS サービス) へのシークレットなしのポリシー主導のアクセスを提供します。つまり、Cordium は、分離された実行のためのサンドボックス プラットフォームであるだけでなく、認証情報を必要としないインフラストラクチャへの安全なアクセスのためのプラットフォームでもあります。

大規模な注入。
他のプラットフォームとの比較
人間と AI エージェントのための統合プラットフォーム。すべてのサンドボックス (つまり、ワークスペース) は、標準の Kubernetes 上で実行される分離されたルートレス コンテナー サンドボックスであり、ブラウザー ターミナル、SSH、CLI、および gRPC API を介してアクセスできます。ワークスペースは永続的 (再起動後もファイルシステムが保持される) にすることも、一時的なものにすることもでき、タスクが完了すると自動的に停止します。同じプラットフォームを、長期間のコーディング セッション (例: VSCode、Zed など) と短期間の自動化ワークロード (例: AI エージェント タスク、CI/CD など) に使用できます。
宣言的で再現可能な環境。ワークスペース環境は、コンテナー イメージ、リポジトリのクローン作成、ライフサイクル タスク、環境変数、リソース制限 (CPU、メモリ、ストレージなど)、変数置換、アプリケーション ポートを含む YAML 仕様で定義されます。テンプレートを使用すると、単一の構成を多くのワークスペースで再利用できます。事前に構築されたテンプレートは、完全に初期化されたファイルシステムを Kubernetes VolumeSnapshot としてキャプチャし、コールド スタートを数分から数秒に短縮します。
インフラストラクチャへの秘密のない ID ベースのアクセス。ワークスペースは、資格情報がワークスペースに到達することなく、データベース、SSH サーバー、HTTP API、Kubernetes クラスター、mTLS で保護されたサービスにアクセスします。 API キー、パスワード、SSH 秘密キー、および kubeconfig は、Octelium ID 認識プロキシで保持され、Workspace ID が承認されている場合はプロトコル層に挿入されます。ワークスペース自体は認証情報を保持しないため、開発者と AI エージェントの両方の認証情報が無秩序に拡散することがなくなります。
ID ベースのアクセス制御と可観測性。すべてのワークスペースには、そのアイデンティティを表す Octelium セッションがあります。インフラストラクチャへのアクセスは、CEL および OPA を使用したコードとしてのポリシーを備えたリクエストごとの L7 対応属性ベースのアクセス制御 (ABAC) によって管理され、z を強制します。

デフォルトでエロ立ち権限。認証は、OIDC または SAML 2.0 ID プロバイダー (IdP)、GitHub OAuth2、ワークロード OIDC アサーション、ネイティブ FIDO2、WebAuthn、TOTP をサポートします。
OpenTelemetry ネイティブの監査と可視性。リアルタイム、ID ベース、L7 対応の可視性とアクセスログ。すべてのリクエストは Octelium Cluster によって監査され、ログ管理および SIEM プロバイダーと統合するために OpenTelemetry OTLP レシーバーにエクスポートされます。
オープンソースで自己ホスト型。 Cordium は、Apache-2.0 の下で完全にオープンソースです。これは、単一ノードの VM から運用環境のマルチノードのインストール、クラウドまたはオンプレミスに至るまで、あらゆる Kubernetes クラスター上で実行されます。独自のコントロール プレーン、階層化された機能セット、ベンダー ロックインはありません。
Space は、Cordium のトップレベルの名前空間です。テンプレート、ワークスペース、シークレット、および GitProvider を 1 つの組織単位の下にグループ化します。
ワークスペース (サンドボックスと同義) は、Cordium の基本的な実行ユニットです。これは、Web ベースのコンソール、cordium CLI、標準 SSH、および gRPC ベースの API を介して対話式またはプログラムで使用できる、分離されたルートレスのコンテナベースの環境です。
テンプレートは、スペース内で再利用可能なワークスペース構成を定義します。ワークスペースが作成されると、選択したテンプレートの仕様に基づいて初期化されます。すべてのスペースには、自動的に作成されるデフォルトのテンプレートがあります。テンプレートの仕様は、ワークスペース仕様 (イメージ、ランタイムなど) のほとんどと、オプションの GitProvider 関連付けを共有します。テンプレートは、Kubernetes VolumeSnapshot を介した事前ビルドをサポートし、依存関係の多いテンプレートの起動時間を数分から数秒に短縮します。
Secret は、スペース内に保存されている機密の値 (API キー、トークン、パスワード、証明書) を表します。シークレットは、テンプレート仕様で名前によって参照されます。
ワークスペースとテンプレートの仕様は YAML で定義されます

そして、cordium run --file spec.yaml またはcordium create template --file spec.yaml を介して適用されます。最小限の例を次に示します。
仕様：
画像：
レジストリ:
URL : python:3.12-bookworm
リポジトリ:
URL : https://github.com/myorg/my-project
ランタイム:
タスク:
- 名前: インストール
タイプ: ON_CREATE
workingDir : /workspace/repo
実行：pip install -rrequirements.txt
onFailure : ON_FAILURE_ABORT
より包括的な例を次に示します。
仕様：
画像：
ドッカーファイル:
インライン : |
FROM ノード:22-bookworm
ENV DEBIAN_FRONTEND=非対話型
apt-get update -qq && を実行します \
apt-get install -y --no-install-recommends \
ポッドマン\
postgresql-クライアント\
カール \
git\
&& rm -rf /var/lib/apt/lists/*
RUN npm install -g @anthropic-ai/claude-code
リポジトリ:
URL : https://github.com/myorg/api-service
クローンオプション:
ブランチ: ${{ vars.BRANCH }}
深さ：1
変数:
- 名前 : 支店
値: メイン
- 名前 : タスク
value : " コードベースを確認し、失敗したテストを修正します。"
ランタイム:
自動停止 : true
環境変数:
- キー : NODE_ENV
値: テスト
- キー : SENTRY_DSN
fromSecret : セントリー DSN
タスク:
- 名前 : 依存関係のインストール
タイプ: ON_CREATE
workingDir : /workspace/repo
実行：npm ci
onFailure : ON_FAILURE_ABORT
- 名前 : start-postgres
タイプ: POST_START
runAsRoot : true
実行: |
ポッドマン rm -f ワークスペース-postgres 2>/dev/null ||本当の
ポッドマン実行\
--name ワークスペース-postgres \
--ネットホスト\
-e POSTGRES_PASSWORD=パスワード \
-e POSTGRES_DB=アプリ \
-d docker.io/postgres:16
for i in $(seq 1 30);する
pg_isready -h localhost -p 5432 -U postgres && Break
睡眠2
完了しました
- 名前 : 実行テスト
タイプ: POST_START
workingDir : /workspace/repo
実行：npmテスト
onFailure : ON_FAILURE_CONTINUE
- 名前 : 実行エージェント
タイプ: POST_START
workingDir : /workspace/repo
実行: クロード "${{ vars.TASK }}"
制限:
CPU:
ミリコア: 4000
メモリ:
メガバイト: 8192
ストレージ:
メガバイト: 20480
アクセスM

方法
Cordium Web ポータルは、ソフトウェアをインストールせずにワークスペースを管理および操作するためのブラウザベースのインターフェイスです。これは、ワークスペースへのクライアントレス アクセスを必要とするユーザーとチームにとっての主要なインターフェイスです。 Octelium Web ポータルは、GitHub OAuth2、OpenID Connect、または SAML 2.0 IdP を含む Octelium の IdentityProviders (詳細はこちら)、またはパスキー (詳細はこちら) を介して直接ユーザーを認証します。
ここで、Cordium Web ポータルの短いデモ ビデオを視聴できます。
Cordium CLI は、ワークスペース管理への完全なコマンドライン アクセスを提供します。以下にいくつかの例を示します。
エクスポートOCTELIUM_DOMAIN=example.com
# デフォルトのテンプレートから作成し、ターミナルをアタッチします
コーディウムラン
# 既存のワークスペースにアタッチします (停止している場合は開始します)
コーディウム ラン abc
# 特定のテンプレートから作成する
Cordium run --template ml-env.my-project
# YAML 設定ファイルから作成
Cordium run --file workspace.yaml
# コンテナイメージから作成
Cordium run --image python:3.11-slim
# Dockerfile から作成する
Cordium run --dockerfile ./Dockerfile
# git リポジトリから作成し、特定のブランチを複製します
Cordium run --repository https://github.com/myorg/my-project --branch Development
# スペースのデフォルトのテンプレートを使用して特定のスペースに作成します
Cordium run --space my-project
# 一時的なワークスペースを作成する
Cordium run --ephemeral
# ターミナルセッションの終了時に削除される一時的なワークスペースを作成します
Cordium run --ephemeral --rm
# リソース制限とシークレットを備えた一時的な AI エージェント サンドボックス
Cordium run --ephemeral --rm \
--image Python:3.11-slim \
--env-from-secret ANTHROPIC_API_KEY=anthropic-key \
--CPU 2000 --メモリ 4096
# 環境変数を設定する
Cordium run --image ノード:20 -e NODE_ENV=開発 -e PORT=3000
# スペースシークレットから変数を取得します
コージウム

--template backend.my-project \ を実行します。
--env-from-secret DATABASE_URL=ステージング データベース URL
# vars を使用してプライマリ リポジトリと追加リポジトリのクローンを作成します
Cordium run --repository https://github.com/myorg/api-service \
--Additional-reposhared-lib=https://github.com/myorg/shared-lib \
--var SERVICE=サービス/支払い \
--var BRANCH=メイン
# ワークスペースターミナルを開きます
コージウム端子abc
Cordium 用語 abc # 短いエイリアス
# コマンドを実行して終了コードを伝播する
Cordium exec abc -- テストを作成する
# 特定の作業ディレクトリで実行する
Cordium exec abc -w /workspace/repo -- go build ./...
# root として実行
Cordium exec abc --root --apt-get install -y ripgrep
# または
Cordium exec abc -- sudo pt-get install -y ripgrep
# コマンドごとの環境変数を設定する
Cordium exec abc -e GOOS=linux -e GOARCH=amd64 -- ビルドに進みます ./...
[切り捨てられた]
安価なクラウド VM/VPS インスタンス (例: DigitalOcean Droplet、Hetzner サーバー、AWS EC2、Vultr など)、または最新の Linux ディストリビューション (Ubuntu 24.04 LTS 以降、Debian 12+、 etc...)、いいですね

[切り捨てられた]

## Original Extract

Open-source, general-purpose sandbox platform for devs and AI agents that provides identity-based secure access to infrastructure without credentials. - octelium/cordium

Hello HN. Cordium is a FOSS, self-hosted, identity-based, general-purpose sandbox platform that I've been working on for a long time now that is built on Kubernetes and Octelium, my main work. The key difference here for Cordium, when compared to other dev environments (e.g. GitHub Codespaces) and sandbox platforms (e.g. E2B, Daytona, etc.), is that Cordium automatically provides identity-based, secretless secure access to resources/infrastructure (e.g. APIs, SSH, databases, k8s, etc.) without having to inject credentials (e.g. API keys and access tokens, SSH private keys/passwords, database passwords, mTLS private keys, etc.) into the sandbox where the upstream credential is held by the identity-aware proxy of the Octelium-protected resource outside the reach of the sandbox and inject on-the-fly by the identity-aware proxy if the user is authorized via the Octelium policies. In short, Cordium is not just an isolated execution environment that can replace remote development environments and sandbox platforms, but also equally a secure access platform to infrastructure/resources. It's basically a sandbox platform + a ZTNA/remote-access-VPN baked-in with unified identity management, L7-aware access control and visibility.

GitHub - octelium/cordium: Open-source, general-purpose sandbox platform for devs and AI agents that provides identity-based secure access to infrastructure without credentials. · GitHub
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
octelium
/
cordium
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
214 Commits 214 Commits .github .github .vscode .vscode apis apis client/ cordium client/ cordium cluster cluster pkg pkg unsorted unsorted .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.work go.work go.work.sum go.work.sum View all files Repository files navigation
Cordium is a free and open source, self-hosted, identity-based sandbox platform built on Kubernetes and Octelium . Cordium is a general-purpose platform that provides isolated, reproducible isolated sandboxes for developers, AI agents, and automated workloads. Cordium can be used for various use cases, including as a remote development environment for coding sessions (e.g. VSCode, Zed, etc.), sandboxes for AI agents and CI/CD workloads, and secretless secure remote access to infrastructure for developers and AI agents.
The main differentiator for Cordium, when compared to remote development environments (e.g. GitHub Codespaces, Coder, etc.) and sandbox platforms (e.g. E2B, Daytona, etc.), is that Cordium automatically provides identity-based, secretless access to resources and infrastructure from within the sandboxes. Cordium leverages Octelium ZTNA capabilities to provide secretless, policy-driven access to infrastructure resources (SSH servers, databases, internal HTTP APIs, mTLS services) from within the sandbox, without exposing, distributing, or managing upstream application-layer credentials (e.g. API keys and access tokens, SSH private keys, database passwords, etc.). In short, Cordium is not just a sandbox platform for isolated execution, but also a platform for secure access to infrastructure that eliminates credential injection at scale.
Comparison with Other Platforms
Unified platform for humans and AI agents. Every sandbox (i.e. Workspace) is an isolated, rootless container sandbox running on standard Kubernetes, accessible via browser terminal, SSH, CLI, and gRPC API. Workspaces can be persistent (filesystem preserved across restarts) or ephemeral and can stop automatically when their tasks complete. The same platform can be used for long-lived coding sessions (e.g. via VSCode, Zed, etc.) and short-lived automated workloads (e.g. AI agent tasks, CI/CD, etc.).
Declarative, reproducible environments. Workspace environments are defined in YAML specs covering the container image, repository cloning, lifecycle tasks, environment variables, resource limits (i.e. CPU, memory, storage), variable substitution, and application ports. Templates allow a single configuration to be reused across many Workspaces. Pre-built Templates capture a fully initialized filesystem as a Kubernetes VolumeSnapshot, reducing cold startup from minutes to seconds.
Secretless, identity-based access to infrastructure. Workspaces access databases, SSH servers, HTTP APIs, Kubernetes clusters, and mTLS-protected services without credentials ever reaching the Workspace. API keys, passwords, SSH private keys, and kubeconfigs are held at the Octelium identity-aware proxy and injected at the protocol layer if the Workspace identity is authorized. The Workspace itself does not hold credentials, eliminating credential sprawl for both developers and AI agents.
Identity-based access control and observability. Every Workspace has an Octelium Session that represents its identity. Infrastructure access is governed by per-request, L7-aware attribute-based access control (ABAC) with policy-as-code using CEL and OPA, enforcing zero standing privileges by default. Authentication supports any OIDC or SAML 2.0 identity provider (IdP), GitHub OAuth2, workload OIDC assertions, and native FIDO2, WebAuthn, and TOTP.
OpenTelemetry-native auditing and visibility. Real-time, identity-based, L7-aware visibility and access logging. Every request is audited by the Octelium Cluster and exported to your OpenTelemetry OTLP receivers for integration with log management and SIEM providers.
Open source and self-hosted. Cordium is fully open source under Apache-2.0. It runs on any Kubernetes cluster, from a single-node VM to production multi-node installations, cloud or on-premises. There is no proprietary control plane, no tiered feature set, and no vendor lock-in.
Space is the top-level namespace in Cordium. It groups Templates, Workspaces, Secrets, and GitProviders under a single organizational unit.
Workspace (synonymous with sandbox) is the fundamental execution unit in Cordium. It is an isolated, rootless container-based environment that can be used interactively or programmatically via web-based console, cordium CLI, standard SSH, and gRPC-based APIs.
Template defines a reusable Workspace configuration within a Space. When a Workspace is created, it is initialized from the selected Template's spec. Every Space has a default Template created automatically. A Template's spec shares most of Workspace spec (image, runtime, etc.) as well as an optional GitProvider association. Templates support pre-builds via Kubernetes VolumeSnapshot to reduce startup time from minutes to seconds for dependency-heavy Templates.
Secrets represents a sensitive value (API keys, tokens, passwords, certificates) stored within a Space. Secrets are referenced by name in Template specs.
Workspace and Template specs are defined in YAML and applied via cordium run --file spec.yaml or cordium create template --file spec.yaml . Here is a minimal example:
spec :
image :
registry :
url : python:3.12-bookworm
repository :
url : https://github.com/myorg/my-project
runtime :
tasks :
- name : install
type : ON_CREATE
workingDir : /workspace/repo
run : pip install -r requirements.txt
onFailure : ON_FAILURE_ABORT
Here is a more comprehensive example:
spec :
image :
dockerfile :
inline : |
FROM node:22-bookworm
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update -qq && \
apt-get install -y --no-install-recommends \
podman \
postgresql-client \
curl \
git \
&& rm -rf /var/lib/apt/lists/*
RUN npm install -g @anthropic-ai/claude-code
repository :
url : https://github.com/myorg/api-service
cloneOptions :
branch : ${{ vars.BRANCH }}
depth : 1
vars :
- name : BRANCH
value : main
- name : TASK
value : " Review the codebase and fix any failing tests. "
runtime :
autoStop : true
envVars :
- key : NODE_ENV
value : test
- key : SENTRY_DSN
fromSecret : sentry-dsn
tasks :
- name : install-dependencies
type : ON_CREATE
workingDir : /workspace/repo
run : npm ci
onFailure : ON_FAILURE_ABORT
- name : start-postgres
type : POST_START
runAsRoot : true
run : |
podman rm -f workspace-postgres 2>/dev/null || true
podman run \
--name workspace-postgres \
--net host \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=app \
-d docker.io/postgres:16
for i in $(seq 1 30); do
pg_isready -h localhost -p 5432 -U postgres && break
sleep 2
done
- name : run-tests
type : POST_START
workingDir : /workspace/repo
run : npm test
onFailure : ON_FAILURE_CONTINUE
- name : run-agent
type : POST_START
workingDir : /workspace/repo
run : claude "${{ vars.TASK }}"
limit :
cpu :
millicores : 4000
memory :
megabytes : 8192
storage :
megabytes : 20480
Access Methods
The Cordium web portal is a browser-based interface for managing and interacting with Workspaces without installing any software. It is the primary interface for users and teams who want clientless access to their Workspaces. The Octelium web portal authenticates users through Octelium's IdentityProviders, including GitHub OAuth2 or any OpenID Connect or SAML 2.0 IdP (read more here) or directly via Passkeys (read more here ).
You can a watch a short demo video for the Cordium web portal here .
The cordium CLI provides full command-line access to Workspace management. Here are some examples:
export OCTELIUM_DOMAIN=example.com
# Create from the default Template and attach a terminal
cordium run
# Attach to an existing Workspace (starts it if stopped)
cordium run abc
# Create from a specific Template
cordium run --template ml-env.my-project
# Create from a YAML configuration file
cordium run --file workspace.yaml
# Create from a container image
cordium run --image python:3.11-slim
# Create from a Dockerfile
cordium run --dockerfile ./Dockerfile
# Create from a git repository, cloning a specific branch
cordium run --repository https://github.com/myorg/my-project --branch develop
# Create in a specific Space using that Space's default Template
cordium run --space my-project
# Create an ephemeral Workspace
cordium run --ephemeral
# Create an ephemeral Workspace that is deleted when the terminal session ends
cordium run --ephemeral --rm
# Ephemeral AI agent sandbox with resource limits and a secret
cordium run --ephemeral --rm \
--image python:3.11-slim \
--env-from-secret ANTHROPIC_API_KEY=anthropic-key \
--cpu 2000 --memory 4096
# Set environment variables
cordium run --image node:20 -e NODE_ENV=development -e PORT=3000
# Source a variable from a Space Secret
cordium run --template backend.my-project \
--env-from-secret DATABASE_URL=staging-db-url
# Clone the primary repository and an additional repository with vars
cordium run --repository https://github.com/myorg/api-service \
--additional-repo shared-lib=https://github.com/myorg/shared-lib \
--var SERVICE=services/payments \
--var BRANCH=main
# Open a Workspace terminal
cordium terminal abc
cordium term abc # short alias
# Run a command and propagate exit code
cordium exec abc -- make test
# Run in a specific working directory
cordium exec abc -w /workspace/repo -- go build ./...
# Run as root
cordium exec abc --root -- apt-get install -y ripgrep
# Or
cordium exec abc -- sudo pt-get install -y ripgrep
# Set per-command environment variables
cordium exec abc -e GOOS=linux -e GOARCH=amd64 -- go build ./...
[truncated]
Read this quick guide here to install a single-node Cordium Cluster on top of any cheap cloud VM/VPS instance (e.g. DigitalOcean Droplet, Hetzner server, AWS EC2, Vultr, etc...) or a local Linux machine/Linux VM inside a MacOS/Windows machine with at least 4GB of RAM and 20GB of disk storage running a recent Linux distribution (Ubuntu 24.04 LTS or later, Debian 12+, etc...), which is good enou

[truncated]

---
source: "https://github.com/deployah-dev/deployah"
hn_url: "https://news.ycombinator.com/item?id=49186881"
title: "Deployah – deploy to Kubernetes from a short spec, no Helm, nothing in-cluster"
article_title: "GitHub - deployah-dev/deployah: Spec-to-Release for Kubernetes: turn a short app spec into a real Helm release. Zero Helm knowledge, zero cluster-side setup, one binary. · GitHub"
author: "atkrad"
captured_at: "2026-08-05T19:15:07Z"
capture_tool: "hn-digest"
hn_id: 49186881
score: 3
comments: 6
posted_at: "2026-08-05T18:26:38Z"
tags:
  - hacker-news
  - translated
---

# Deployah – deploy to Kubernetes from a short spec, no Helm, nothing in-cluster

- HN: [49186881](https://news.ycombinator.com/item?id=49186881)
- Source: [github.com](https://github.com/deployah-dev/deployah)
- Score: 3
- Comments: 6
- Posted: 2026-08-05T18:26:38Z

## Translation

タイトル: Deployah – ヘルムなし、クラスター内なしの短い仕様から Kubernetes にデプロイします
記事のタイトル: GitHub -deployah-dev/deployah: Kubernetes の仕様からリリースまで: 短いアプリ仕様を実際の Helm リリースに変える。 Helm の知識はゼロ、クラスター側のセットアップはゼロ、バイナリは 1 つ。 · GitHub
説明: Kubernetes の仕様からリリースまで: 短いアプリ仕様を実際の Helm リリースに変換します。 Helm の知識はゼロ、クラスター側のセットアップはゼロ、バイナリは 1 つ。 - デプロイア-dev/デプロイア

記事本文:
GitHub -deployah-dev/deployah: Kubernetes の仕様からリリースまで: 短いアプリ仕様を実際の Helm リリースに変えます。 Helm の知識はゼロ、クラスター側のセットアップはゼロ、バイナリは 1 つ。 · GitHub
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
デプロイア開発
/
デプロイヤ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
158 コミット 158 コミット .github .github docs docs example

les/ nginx 例/ nginx 内部 内部 nix nix シナリオ シナリオ .dockerignore .dockerignore .editorconfig .editorconfig .envrc .envrc .gitignore .gitignore .golangci.yaml .golangci.yaml .markdownlint.json .markdownlint.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md codecov.yml codecov.yml docker-bake.hcl docker-bake.hcl flake.lock flake.lock flake.nix flake.nix go.mod go.mod go.sum go.sum main.go main.go osv-scanner.toml osv-scanner.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ヘルム知識ゼロ。クラスター側のセットアップは不要です。バイナリが 1 つ。
Deployah は、アプリを Kubernetes にデプロイする CLI です。間の隙間に鎮座しています
Helm の作成を依然として要求するツールと、大量のクラスター内処理を必要とするツール
プラットフォーム。内部で Helm を使用し、 helm 、 kubectl 、および kind を埋め込みます。
ライブラリをインストールし、クラスターには何もインストールしません。
短い仕様を書きます。 Deployah はそれを実行中のリリースにします。
Kubernetes。これを Spec-to-Release と呼びます。ソースからイメージへ (S2I) のようなものです。
ただし、デプロイ手順については、S2I がイメージをビルドし、Deployah がリリースを実行します。
brew install デプロイヤ-dev/tap/deployah
Nix の使用 (推奨)
# インストールせずに実行する
nix run github:deployah-dev/deployah
# または、flake.nix に追加します
inputs.deployah.url = " github:deployah-dev/deployah " ;
Go でインストールする
インストールに移動します。deployah.dev/deployah@latest
要件
Deployah は単一のバイナリです。 helm 、 kubectl 、または kind は必要ありません
コマンドラインツール。 Deployah には、Kubernetes クライアントである Helm と Kind が含まれています。
ライブラリなので、それ自体でクラスターと通信します。これにより、クラス全体が削除されます
CLI ツールの欠落または不一致が原因で発生する「自分のマシンで動作する」問題。
すでに所有しているクラスターにデプロイします。

それにアクセスする必要があります (
kubeconfig）。コンテナー ランタイムは必要ありません。
組み込みのローカル クラスターを使用します (deployah クラスター up )。
コンテナー ランタイム ( Docker または Podman )。追加分はこれだけです
このツールはローカル クラスターにのみ必要です。
これにより、自分のマシンに完全にデプロイする手順が順を追って説明されます。 5個くらいかかります
分。ローカル クラスターの場合は、Docker または Podman を実行する必要があります (「
要件）。
既存の Kubernetes クラスターは必要ありません。 Deployah はローカルのものを作成できます
あなたのために。
デプロイアクラスターアップ
これにより、小規模なローカル Kubernetes クラスターが (Kind を使用して) 作成され、それに
コンテキスト名 kind-deployah 。
これを空のフォルダーにdeployah.yamlとして保存します。パブリック nginx を実行します
イメージなので、何も構築する必要はありません。
APIバージョン : v1-alpha.2
プロジェクト : my-first-app
コンポーネント:
ウェブ：
画像：nginx：最新
ポート：80
環境: [ローカル]
公開: true
Expose: true はコンポーネントにその名前から作られたホスト名を与えます (ここでは
web.127.0.0.1.nip.io ) と HTTPS、すべてプラットフォーム ファイルによって決定されます。
Deployah には、マップするプラットフォーム ファイルdeployah.platform.yaml も必要です。
ローカル環境をドメインと Kubernetes コンテキストに変換します。これは、deployah クラスターアップによって自動的に作成されます。 「プラットフォーム ファイル」を参照してください。
デプロイ ローカルにデプロイ
4. 実行中の様子を確認する
# プロジェクトのステータスを表示する
デプロイア ステータス my-first-app
# ローカルクラスターと開くことができる URL を表示します
デプロイアクラスターステータス
デプロイア クラスター ステータスは、アプリのすぐに使用できる URL を出力します。で開いてください
ブラウザで nginx のようこそページを表示します。
デプロイアのログ my-first-app
5. クリーンアップ
# アプリを削除する
デプロイ、私の最初のアプリをローカルに削除
# ローカルクラスターを停止して削除します
デプロイヤクラスターがダウンしました
例
実行可能な仕様は、examples/ の下にあります。から始める
Examples/nginx : クイック スタートと同じフローで、すぐに実行できます。
警官

やあ。
Deployah は、deployah.yaml 仕様を実行中の Kubernetes デプロイメントに変換します。
3つのステップ。
フローチャート LR
サブグラフphase1["1. 仕様を読む"]
TB方向
A["YAML仕様"] --> B["解析"] --> C["検証"]
終わり
サブグラフ フェーズ 2["2. 構成の解決"]
TB方向
D["環境を選択"] --> E["変数を適用"] --> F["デフォルトを入力"]
終わり
サブグラフphase3["3.デプロイ"]
TB方向
G["ヘルム値のビルド"] --> H["リリースのインストール"]
終わり
フェーズ1 --> フェーズ2 --> フェーズ3
読み込み中
仕様を読んでください。 Deployah は、deployah.yaml を読み取り、それを
JSON スキーマなので、明確なメッセージで間違いを早期に発見できます。
設定を解決します。 Deployah は、あなたが求めた環境を選択し、代替します
変数を選択し、適切なデフォルト値を入力します。
展開する。 Deployah は仕様から Helm 値を構築し、Helm をインストールします
クラスター上でリリースします。 Helm チャートを自分で書くことはありません。
Deployah を同様のツール (DevSpace、Werf、Score、Epinio、
Kubero)、 docs/comparison.md を参照してください。
よく目にする言葉をいくつか。
プロジェクト。名前が付いた 1 つのアプリ。この名前は、Kubernetes リソースの接頭辞として付加されます。
デプロイアが作成します。これは仕様内のプロジェクト フィールドです。
成分。 Web サービスや
バックグラウンドワーカー。プロジェクトには 1 つ以上のコンポーネントが含まれます。
役割。コンポーネントの目的:
service : トラフィックを処理し、公開できます (デフォルト)。
worker : 長時間実行されるバックグラウンド タスク。公開されていません。
job : 実行して停止する 1 回限りのタスク。
親切。ステートレス (デフォルト、拡張が容易) またはステートフル (必要な
永続ストレージ）。プラットフォーム チームはすでにストレージ クラスを宣言できます。
ステートフルが土地を展開するとき。 「ストレージ クラス」を参照してください。
今日展開されるもの。 Deployah は現在ステートレス サービスをデプロイしています
コンポーネント。ワーカーとジョブの役割、およびステートフルな種類は、
スキーマですが、デではありません

まだ展開可能であるため、それらを使用する展開は次のエラーで停止します。
「まだサポートされていません」エラー。
環境。 dev 、 staging 、 prod などのターゲット。それぞれ
環境では、別のクラスター、別のファイル、および別のファイルを使用できます。
変数。プラットフォーム ファイルは、どの環境が存在するかを登録します。エントリ
仕様の環境マップでは、そのうちの 1 つのオーバーライドのみが追加されます。
リソースのプリセット。知らずにCPUとメモリを設定する簡単な方法
Kubernetes ユニット。正確な値を書き込む代わりに、resourcePreset: small を使用します。
これはプロファイルとは異なります。
プロフィール。プラットフォーム チーム (ノード) が所有する名前付きデプロイメント ポリシー
配置、セキュリティ コンテキスト、ドメインとリソースの上限など)。
コンポーネントは、プロファイル [...] を使用して 1 つ以上を選択します。 「プロファイル」を参照してください。
健康診断。 Deployah は、アプリがトラフィックに対応できる準備ができているかどうかを確認し、
詰まった場合は再起動します。これはすべてのサービスで自動的に行われます
コンポーネント。 Deployah に HTTP エンドポイントを与えることで、チェックを改善できます。
電話する。 「ヘルスチェック」を参照してください。
ご自身の画像をご持参ください。 Deployah はイメージを構築しません。あなたはそれにイメージを与えます
クラスターが取得できるレジストリにすでに存在します。イメージを構築する
CI (またはローカル) で、Deployah にデプロイさせます。
仕様は、deployah.yaml という名前のファイルです。これには 3 つの必須部分があります。
apiVersion 、プロジェクト、およびコンポーネント。環境も定義します。
Deployah は、それぞれの所有者が異なる 2 つのファイルに構成を分割します。
開発者は何を実行するかを説明します。プラットフォーム チームは、プラットフォームが実行される場所をマッピングします。いいえ
共有サーバーが必要です。
.deployah.yaml (このセクション)。開発者が所有しています。内容を説明します
実行するイメージ、ポート、リソース、ヘルスチェック、およびどの論理ドメインを実行するか
に公開します。 Kubernetes コンテキストや実際のドメイン名が含まれることはありません。
デプロイah.platform.yaml 。プラットフォームチームが所有します。それぞれのマップ
環境

実際の Kubernetes コンテキスト、ドメイン、TLS 戦略に取り組みます。参照
プラットフォーム ファイル 。
この分割は、開発者が環境を追加したり、コンポーネントを公開したりできることを意味します。
どのクラスターまたはドメインで実行されているかを知らなくても、プラットフォーム チームは次のことを行うことができます。
アプリの仕様に触れることなく、クラスターを変更したり、証明書をローテーションしたりできます。
以下は、一般的なフィールドを示す完全な例です。全部は必要ありません
彼ら。ほとんどの場合はデフォルトが設定されています。
apiVersion : v1-alpha.2 # 必須: スキーマのバージョン
プロジェクト : ショップ # 必須: プロジェクト名
コンポーネント : # 必須: 1 つ以上のコンポーネント
API:
image : ghcr.io/acme/shop-api:${TAG} # タグは以下の環境からのものです
役割 : サービス # サービス |労働者 |ジョブ (デフォルト: サービス)
種類: ステートレス # ステートレス |ステートフル (デフォルト: ステートレス)
port : 8080 # アプリがリッスンするポート (デフォルト: 8080)
環境 : [staging, prod] # このコンポーネントをデプロイする環境
command : ["/bin/api"] # オプション: イメージをオーバーライドします。
args : ["--verbose"] # オプション: イメージ CMD をオーバーライドします
env : # 計画中: まだコンテナに適用されていません
LOG_LEVEL : 情報
resourcePreset : small # nano|micro|small|medium|large|xlarge|2xlarge
Expose : # オプション: `expose: true` はすべてデフォルトを使用します
サブドメイン : api # オプション: デフォルトはコンポーネント名です
# ドメイン: 内部 # オプション: デフォルトはプラットフォームのデフォルト ドメインになります
# apex: true # オプション: サブドメインの代わりにベアドメインを使用します
autoscaling : # オプション: CPU またはメモリ上でスケールします
有効 : true
最小レプリカ : 2
最大レプリカ数 : 5
メトリクス:
- タイプ: cpu # cpu |記憶
target : 70 # ターゲット使用率
環境 : # 環境を定義します (リストではなくマップ)
ステージング:
変数:
TAG : 1.4.0-rc # 上の画像の ${TAG} を埋めます
製品:
変数:
TAG : 1.4.0 # 上の画像の ${TAG} を埋めます
コンテキストフィールドがないことに注意してください

ここに d: それぞれの Kubernetes コンテキスト
環境は、deployah.yaml ではなく、deployah.platform.yaml から取得されます。
resourcePreset または resource の両方ではなく、いずれかを使用します。プリセットは簡単です
オプション。リソースを使用すると、正確な CPU、メモリ、一時ストレージを設定できます。
まだデプロイされていません: スキーマは、 role: worker および role: job を受け入れます。
kind: stateful 、および env 、 envFile 、および configFile フィールド、ただし
Deployah は、デプロイ時にこれらをまだ適用しません。今日、
イメージ、ポート、リソース、またはを使用するステートレス サービス
resourcePreset 、expose 、autoscaling 、および profile 。
環境にはコンテキスト フィールドはありません。コンテキスト フィールドは、一致するものから取得されます。
deployah.platform.yaml の環境キー。
仕様を確認するには、deployah validate を実行します。プラットフォーム ファイルが存在する場合は、
また、expose.domain キーと環境名をそれに対してクロスチェックします。に
特定の環境の最大解像度を確認し、実行します
デプロイは <環境> を検証します。
いくつかのフィールドには特定の形式があります。
ポート: 1 ～ 65535 の数値。
resource.cpu : 500m などのミリコア、または 1 または 2 などのコア全体。
resource.memory および resource.ephemeralStorage : 付きの数値
256Mi や 1Gi などのユニット。
env : キーは大文字、数字、アンダースコアであり、開始

[切り捨てられた]

## Original Extract

Spec-to-Release for Kubernetes: turn a short app spec into a real Helm release. Zero Helm knowledge, zero cluster-side setup, one binary. - deployah-dev/deployah

GitHub - deployah-dev/deployah: Spec-to-Release for Kubernetes: turn a short app spec into a real Helm release. Zero Helm knowledge, zero cluster-side setup, one binary. · GitHub
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
deployah-dev
/
deployah
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
158 Commits 158 Commits .github .github docs docs examples/ nginx examples/ nginx internal internal nix nix scenarios scenarios .dockerignore .dockerignore .editorconfig .editorconfig .envrc .envrc .gitignore .gitignore .golangci.yaml .golangci.yaml .markdownlint.json .markdownlint.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md codecov.yml codecov.yml docker-bake.hcl docker-bake.hcl flake.lock flake.lock flake.nix flake.nix go.mod go.mod go.sum go.sum main.go main.go osv-scanner.toml osv-scanner.toml View all files Repository files navigation
Zero Helm knowledge. Zero cluster-side setup. One binary.
Deployah is a CLI that deploys apps to Kubernetes. It sits in the gap between
tools that still ask you to write Helm and tools that need a heavy in-cluster
platform. It uses Helm under the hood, embeds helm , kubectl , and kind as
libraries, and installs nothing in your cluster.
You write a short spec . Deployah turns it into a running release on
Kubernetes. We call this Spec-to-Release . It is like Source-to-Image (S2I),
but for the deploy step: S2I builds your image, and Deployah runs your release.
brew install deployah-dev/tap/deployah
Using Nix (recommended)
# Run without installing
nix run github:deployah-dev/deployah
# Or add it to your flake.nix
inputs.deployah.url = " github:deployah-dev/deployah " ;
Install with Go
go install deployah.dev/deployah@latest
Requirements
Deployah is a single binary. You do not need the helm , kubectl , or kind
command-line tools. Deployah includes Helm, the Kubernetes client, and Kind as
libraries, so it talks to your cluster by itself. That removes a whole class of
"works on my machine" problems caused by missing or mismatched CLI tools.
Deploy to a cluster you already have: you only need access to it (a
kubeconfig). No container runtime is required.
Use the built-in local cluster ( deployah cluster up ): you need a
container runtime, either Docker or Podman . This is the only extra
tool, and it is needed only for the local cluster.
This walks you through one full deploy on your own machine. It takes about five
minutes. For the local cluster you need Docker or Podman running (see
Requirements ).
You do not need an existing Kubernetes cluster. Deployah can make a local one
for you.
deployah cluster up
This creates a small local Kubernetes cluster (using Kind) and gives it the
context name kind-deployah .
Save this as deployah.yaml in an empty folder. It runs the public nginx
image, so you do not need to build anything.
apiVersion : v1-alpha.2
project : my-first-app
components :
web :
image : nginx:latest
port : 80
environments : [local]
expose : true
expose: true gives the component a hostname made from its name (here
web.127.0.0.1.nip.io ) with HTTPS, all decided by the platform file.
Deployah also needs a platform file , deployah.platform.yaml , that maps
the local environment to a domain and a Kubernetes context. deployah cluster up creates this for you automatically. See Platform file .
deployah deploy local
4. See it running
# Show the status of your project
deployah status my-first-app
# Show the local cluster and the URLs you can open
deployah cluster status
deployah cluster status prints a ready-to-use URL for your app. Open it in
your browser to see the nginx welcome page.
deployah logs my-first-app
5. Clean up
# Remove the app
deployah delete my-first-app local
# Stop and delete the local cluster
deployah cluster down
Examples
Runnable specs live under examples/ . Start with
examples/nginx : the same flow as the quick start, ready to
copy.
Deployah turns your deployah.yaml spec into a running Kubernetes deployment in
three steps.
flowchart LR
subgraph phase1["1. Read the spec"]
direction TB
A["YAML spec"] --> B["Parse"] --> C["Validate"]
end
subgraph phase2["2. Resolve config"]
direction TB
D["Pick environment"] --> E["Apply variables"] --> F["Fill defaults"]
end
subgraph phase3["3. Deploy"]
direction TB
G["Build Helm values"] --> H["Install release"]
end
phase1 --> phase2 --> phase3
Loading
Read the spec. Deployah reads your deployah.yaml and checks it against a
JSON Schema, so mistakes are caught early with clear messages.
Resolve config. Deployah picks the environment you asked for, substitutes
your variables, and fills in sensible defaults.
Deploy. Deployah builds Helm values from your spec and installs a Helm
release on your cluster. You never write a Helm chart yourself.
For how Deployah compares to similar tools (DevSpace, Werf, Score, Epinio,
Kubero), see docs/comparison.md .
A few words you will see often.
Project. One app, with a name. The name prefixes the Kubernetes resources
Deployah creates. It is the project field in your spec.
Component. One deployable part of your project, such as a web service or a
background worker. A project has one or more components.
Role. What a component is for:
service : it serves traffic and can be exposed (the default).
worker : a long-running background task, not exposed.
job : a one-off task that runs and then stops.
Kind. stateless (the default, easy to scale) or stateful (needs
persistent storage). Platform teams can already declare storage classes for
when stateful deploys land; see Storage classes .
What deploys today. Deployah currently deploys stateless service
components. The worker and job roles and the stateful kind are in the
schema but are not deployable yet, so a deploy that uses them stops with a
"not supported yet" error.
Environment. A target such as dev , staging , or prod . Each
environment can use a different cluster, different files, and different
variables. The platform file registers which environments exist; an entry
in the spec's environments map only adds overrides for one of them.
Resource preset. A quick way to set CPU and memory without knowing
Kubernetes units. Use resourcePreset: small instead of writing exact values.
This is not the same as a profile .
Profile. A named deployment policy owned by the platform team (node
placement, security context, domain and resource ceilings, and more).
Components select one or more with profiles: [...] . See Profiles .
Health checks. Deployah checks that your app is ready for traffic and
restarts it if it gets stuck. This happens automatically for every service
component. You can improve the checks by giving Deployah an HTTP endpoint to
call. See Health checks .
Bring your own image. Deployah does not build images. You give it an image
that already exists in a registry your cluster can pull from. Build your image
in CI (or locally), then let Deployah deploy it.
Your spec is a file named deployah.yaml . It has three required parts:
apiVersion , project , and components . You also define your environments .
Deployah splits configuration across two files, each with a different owner.
Developers describe what to run. Platform teams map where it runs. No
shared server is required.
deployah.yaml (this section). Owned by the developer. Describes what
to run: image, port, resources, health checks, and which logical domain to
expose on. It never contains a Kubernetes context or a real domain name.
deployah.platform.yaml . Owned by the platform team. Maps each
environment to a real Kubernetes context, domain, and TLS strategy. See
Platform file .
This split means a developer can add an environment or expose a component
without knowing which cluster or domain it runs on, and a platform team can
change clusters or rotate certificates without touching the app spec.
Here is a full example that shows the common fields. You do not need all of
them; most have defaults.
apiVersion : v1-alpha.2 # required: the schema version
project : shop # required: your project name
components : # required: one or more components
api :
image : ghcr.io/acme/shop-api:${TAG} # tag comes from the environment below
role : service # service | worker | job (default: service)
kind : stateless # stateless | stateful (default: stateless)
port : 8080 # the port your app listens on (default: 8080)
environments : [staging, prod] # which environments deploy this component
command : ["/bin/api"] # optional: override the image ENTRYPOINT
args : ["--verbose"] # optional: override the image CMD
env : # planned: not applied to the container yet
LOG_LEVEL : info
resourcePreset : small # nano|micro|small|medium|large|xlarge|2xlarge
expose : # optional: `expose: true` uses all defaults
subdomain : api # optional: defaults to the component name
# domain: internal # optional: defaults to the platform's default domain
# apex: true # optional: use the bare domain instead of a subdomain
autoscaling : # optional: scale on CPU or memory
enabled : true
minReplicas : 2
maxReplicas : 5
metrics :
- type : cpu # cpu | memory
target : 70 # target usage percentage
environments : # define your environments (a map, not a list)
staging :
variables :
TAG : 1.4.0-rc # fills ${TAG} in the image above
prod :
variables :
TAG : 1.4.0 # fills ${TAG} in the image above
Notice there is no context field here: the Kubernetes context for each
environment comes from deployah.platform.yaml , not from deployah.yaml .
Use either resourcePreset or resources , not both. Presets are the easy
option; resources lets you set exact CPU, memory, and ephemeral storage.
Not deployed yet: the schema accepts role: worker and role: job ,
kind: stateful , and the env , envFile , and configFile fields, but
Deployah does not apply them at deploy time yet. Today, deploy a
stateless service using image , port , resources or
resourcePreset , expose , autoscaling , and profiles .
There is no context field on an environment: it comes from the matching
environment key in deployah.platform.yaml .
To check your spec, run deployah validate ; when a platform file exists it
also cross-checks expose.domain keys and environment names against it. To
check the full resolution for a given environment, run
deployah validate <environment> .
A few fields have specific formats:
port : a number from 1 to 65535.
resources.cpu : millicores like 500m , or whole cores like 1 or 2 .
resources.memory and resources.ephemeralStorage : a number with a
unit, like 256Mi or 1Gi .
env : keys are uppercase letters, digits, and underscores, and start

[truncated]

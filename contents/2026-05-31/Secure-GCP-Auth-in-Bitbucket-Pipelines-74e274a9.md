---
source: "https://emilytburak.net/posts/bitbucket-pipes-gap-gcp-oidc/"
hn_url: "https://news.ycombinator.com/item?id=48340288"
title: "Secure GCP Auth in Bitbucket Pipelines"
article_title: "Secure GCP auth in Bitbucket Pipelines · Emily T. Burak"
author: "mooreds"
captured_at: "2026-05-31T00:25:16Z"
capture_tool: "hn-digest"
hn_id: 48340288
score: 1
comments: 0
posted_at: "2026-05-30T20:28:01Z"
tags:
  - hacker-news
  - translated
---

# Secure GCP Auth in Bitbucket Pipelines

- HN: [48340288](https://news.ycombinator.com/item?id=48340288)
- Source: [emilytburak.net](https://emilytburak.net/posts/bitbucket-pipes-gap-gcp-oidc/)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T20:28:01Z

## Translation

タイトル: Bitbucket パイプラインでの安全な GCP 認証
記事のタイトル: Bitbucket Pipelines での安全な GCP 認証 · Emily T. Burak
説明: Bitbucket Pipelines はキーレス GCP 認証の OIDC をサポートしていますが、WIF 交換を自分で接続する必要があります。フォートエンドツーエンドで動作する例。

記事本文:
メインコンテンツにスキップ
エミリー・T・ブラック
メニュー
ブログ
Bitbucket Pipelines での安全な GCP 認証
これは、Bitbucket の悩みシリーズの 2 番目の投稿です。投稿 1 では、GitHub Actions と Bitbucket Pipelines の間のメンタル モデルのギャップについて説明しました。
これは、実際の問題点とギャップによるコストに関するものです。Pipes エコシステムは Actions マーケットプレイスよりもはるかに小さく、そのギャップがクラウド認証周りで最も苦痛であることがわかりました。
Bitbucket は GCP、AWS、Azure へのキーレス認証の OIDC をサポートしていますが、BYOG では独自の接着剤を使用する必要があります。 DevOps 担当者が接着剤とパイプを好むのはなぜでしょうか? GCP Workload Identity Federation の場合は次のようになります。
部屋の中の象（鍵）
見出しへのリンク
別名寿命の長い危険な静電気の秘密は、今私たちの部屋にありますか?
サービス アカウントの JSON キーは便利ですが、非常に危険です。これらは長期間存続し、通常は CI システムのシークレット ストアに保管されており、シークレットのコミットに対する保護がないことによる漏洩 (または罪)、サプライ チェーン攻撃による流出 (ますます一般的)、ビルド ログなどで表面化すると、クラウド領域で実際のアクションを実行できる資格情報が世に出ます。
Workload Identity Federation は、CI プロバイダーによって署名された有効期間の短い OIDC トークンを CI に提示させることでこの問題を修正し、GCP が構成された ID プール/プロバイダーに対してトークンを検証し、GCP が有効期間の短いアクセス トークンを返します (またはサービス アカウントになりすますが、これについては後ほど説明します)。
GitHub Actions を使用すると、これがほとんど簡単になります。 Bitbucket Pipelines はこのフローをサポートしていますが、1 行のアクションは得られません。いくつかのシェル コマンドを記述する必要があります。ごめんなさい。
GitHub Actions ができること
見出しへのリンク
アクション ワークフローでは、WIF の CI 側全体は次のようになります。
権限:
id-token : write # let t

彼はランナーが OIDC トークンを鋳造した
内容：読む
ジョブ:
デプロイ:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- ID : 認証
使用: google-github-actions/auth@v2
付き:
workload_identity_provider : "projects/123456789/locations/global/workloadIdentityPools/my-pool/providers/my-provider"
service_account : "ci-deployer@my-project.iam.gserviceaccount.com"
- 実行: gcloud storage ls gs://my-bucket
それです！ google-github-actions/auth アクションは、プロジェクトの対象ユーザーをスコープとする OIDC トークンの取得、STS の呼び出しとフェデレーション アクセス トークンの取得、環境変数の設定など、あらゆる種類の処理を処理する素晴らしい友達です。オールインワンの用途: !
Bitbucket パイプラインでも同じこと
見出しへのリンク
Bitbucket はビルディング ブロックを経由して公開します。 oidc: OIDC トークンを生成し、その値を $BITBUCKET_STEP_OIDC_TOKEN 環境変数に配置するステップで true。そこから、自家製の接着剤を取り出す必要があります。誰もが自家製接着剤が大好きです。
ステップ 1 — GCP 側のセットアップ (1 回限り)
見出しへのリンク
これは Bitbucket ではなく GCP で構成します。
# 変数
PROJECT_ID = "私のプロジェクト"
プロジェクト番号 = "123456789"
POOL_ID = "ビットバケット プール"
PROVIDER_ID = "bitbucket プロバイダー"
WORKSPACE_UUID = "{your-bitbucket-workspace-uuid}" # Bitbucket ワークスペース設定から
SA_EMAIL = "ci-deployer@ ${ PROJECT_ID } .iam.gserviceaccount.com"
# Workload Identity プールを作成する
gcloud iam workload-identity-pools create " ${ POOL_ID } " \
--project = " ${ PROJECT_ID } " \
--location = "グローバル" \
--display-name = "Bitbucket パイプライン"
# そのプール内に OIDC プロバイダーを作成します
gcloud iam workload-identity-pools drivers create-oidc " ${ PROVIDER_ID } " \
--project = " ${ PROJECT_ID } " \
--location = "グローバル" \
--workload-identity-pool = " ${ POOL_ID } " \
--display-name = "Bitbucket

OIDC"\
--issuer-uri = "https://api.bitbucket.org/2.0/workspaces/ ${ WORKSPACE_UUID } /pipelines-config/identity/oidc" \
--attribute-mapping = "google.subject=assertion.sub,attribute.repository_uuid=assertion.repositoryUuid,attribute.workspace_uuid=assertion.workspaceUuid" \
--attribute-condition = "attribute.workspace_uuid == ' ${ WORKSPACE_UUID } '"
# Bitbucket ワークスペースの ID がサービス アカウントになりすますことを許可します
gcloud iam service-accounts add-iam-policy-binding " ${ SA_EMAIL } " \
--project = " ${ PROJECT_ID } " \
--role = "roles/iam.workloadIdentityUser" \
--member = "principalSet://iam.googleapis.com/projects/ ${ PROJECT_NUMBER } /locations/global/workloadIdentityPools/ ${ POOL_ID } /attribute.workspace_uuid/ ${ WORKSPACE_UUID } "
注目に値するいくつかの微妙な点:
発行者 URI はワークスペースごとです。 Bitbucket は、Bitbucket 全体ではなく、ワークスペース スコープの URL で OIDC 検出ドキュメントを公開します。ワークスペース UUID は、Bitbucket UI の [ワークスペース設定] → [OpenID Connect] から取得できます。
属性条件はセキュリティ境界です。 --attribute-condition を使用しないと、どの Bitbucket ワークスペースもこのプロバイダーにトークンを要求するように機能します。そのため、属性条件のスコープをワークスペース UUID、あるいは本番環境のリポジトリ UUID に設定することをお勧めします。
workloadIdentityUser は、SA になりすますことができる ID をバインドします。ここでの prioritySet:// メンバーは、ワークスペースの任意の ID が ci-deployer になりすますことができることを意味します。代わりに、attribute.repository_uuid/<repo-uuid> にバインドすることをお勧めします。
ステップ 2 — Bitbucket 側の YAML と (s)hell
見出しへのリンク
ここは糊を書く部分です。
開始する実際の例を持っていない読者のために、 gcloud の組み込みの認証情報構成を使用した参照形状を次に示します。これは t の短い方です

一般的なアプローチと、私が好んで使用するアプローチは次のとおりです。
画像 : google/cloud-sdk:slim
パイプライン:
支店:
メイン:
- ステップ:
名前 : WIF 経由で GCP にデプロイする
oidc : true # $BITBUCKET_STEP_OIDC_TOKEN を使用可能にします
スクリプト:
# 1. gcloud が読み取れるファイルに OIDC トークンを書き込む
- echo "$BITBUCKET_STEP_OIDC_TOKEN" > /tmp/oidc-token.txt
# 2. WIF プロバイダーを指す認証情報構成を作成する
- |
gcloud iam workload-identity-pools create-cred-config \
"プロジェクト/${GCP_PROJECT_NUMBER}/locations/global/workloadIdentityPools/${POOL_ID}/providers/${PROVIDER_ID}" \
--サービスアカウント="${SA_EMAIL}" \
--output-file=/tmp/credential-config.json \
--credential-source-file=/tmp/oidc-token.txt
# 3. gcloud (および Google クライアント ライブラリ) にそれを使用するように指示します
- GOOGLE_APPLICATION_CREDENTIALS=/tmp/credential-config.json をエクスポートする
- gcloud 認証ログイン --cred-file=/tmp/credential-config.json --quiet
# 4. 検証してから実際の作業を行う
- gcloud 認証リスト
- gsutil ls gs://${BUCKET}
GCP_PROJECT_NUMBER 、 POOL_ID 、 PROVIDER_ID 、 SA_EMAIL 、および BUCKET は、リポジトリ変数またはデプロイメント変数です。このフローにおける唯一の本当の秘密は、Bitbucket がステップごとに生成する OIDC トークンです。
ここでつまずく可能性がある注意点がいくつかあります。
ワークスペース UUID を再確認します。Bitbucket UUID には中括弧 ( {abc-123-...} ) が含まれており、プロバイダー構成には中括弧が必要です。これらの投稿では、奇妙な Bitbucket 形式がテーマになっている可能性があります。
SA には、principalSet に付与された role/iam.workloadIdentityUser が必要です。やや誤解を招く、または不可解なエラー メッセージという GCP の伝統に従って、エラー メッセージはバインディングではなく SA に名前を付けています。
プール/プロバイダーのパスでは、プロジェクト ID 文字列ではなく、数値のプロジェクト番号を使用します。多くの場合後者が使用されるため、間違いやすいです。
現在のドキュメントと照らし合わせて検証する価値があること
見出しへのリンク
OIDC

、WIF、および gcloud の認証ヘルパー コマンドはすべて進化しました。コピー＆ペーストする前に、以下についてクロスチェックしてください。
Bitbucket Pipelines: OIDC を使用してパイプラインをリソース サーバーと統合する
GCP Workload Identity Federation: OIDC を使用した構成
YMMV。形状は適切である必要があります。構文などが正しいことを確認してください。
シリーズの前の記事: GitHub Actions から来たのですか? Bitbucket Pipelines の実際の感触は次のとおりです
シリーズの次回: Bitbucket API UUID スログ

## Original Extract

Bitbucket Pipelines supports OIDC for keyless GCP auth, but you have to wire up the WIF exchange yourself. Ft. A working end-to-end example.

Skip to main content
Emily T. Burak
Menu
Blog
Secure GCP auth in Bitbucket Pipelines
This is post two in my Bitbucket Woes series. Post one covered the mental-model gap between GitHub Actions and Bitbucket Pipelines.
This one is about a real pain point and cost of the gap: the Pipes ecosystem is much smaller than the Actions marketplace, and I’ve found that that gap is most painful around cloud authentication.
Bitbucket does support OIDC for keyless auth to GCP, AWS, and Azure but you BYOG - bring your own glue. What are DevOps folks but fond of glue and pipes though? Here’s what that looks like for GCP Workload Identity Federation.
The elephant (key) in the room
Link to heading
A.k.a. Are the long-lived, dangerous static secrets in the room with us right now?
Service-account JSON keys are convenient and hella dangerous . They’re long-lived, generally sitting in your CI system’s secret store, and a leak by virtue(or sin) of not having protection against committing secrets(oof), exfiltration by supply chain attack(more and more common) , surfaced in build logs, etc, causes a credential out in the world that can take real action in your cloud space.
Workload Identity Federation fixes this by having CI present a short-lived OIDC token signed by the CI provider, GCP verifies the token against a configured Identity Pool/Provider, GCP hands back a short-lived access token (or impersonates a service account, more probably on this later).
GitHub Actions makes this almost trivial. Bitbucket Pipelines supports the flow, but you don’t get a one-line action — you write a handful of shell commands. I am sorry.
What GitHub Actions does for you
Link to heading
In an Actions workflow, the entire CI side of WIF is roughly:
permissions :
id-token : write # let the runner mint an OIDC token
contents : read
jobs :
deploy :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- id : auth
uses : google-github-actions/auth@v2
with :
workload_identity_provider : "projects/123456789/locations/global/workloadIdentityPools/my-pool/providers/my-provider"
service_account : "ci-deployer@my-project.iam.gserviceaccount.com"
- run : gcloud storage ls gs://my-bucket
That’s it! The google-github-actions/auth action is a beautiful friend that handles all sorts of stuff like getting the OIDC token scoped to your project’s audience, calling STS and getting a federated access token, setting env vars, and more. All in one uses: !
The same thing in Bitbucket Pipelines
Link to heading
Bitbucket exposes the building block via. oidc: true on a step minting an OIDC token and placing its value in the $BITBUCKET_STEP_OIDC_TOKEN environment variable. From there, you gotta get out your homemade glue. Everyone loves homemade glue.
Step 1 — GCP-side setup (one-time)
Link to heading
You configure this in GCP, not in Bitbucket.
# Variables
PROJECT_ID = "my-project"
PROJECT_NUMBER = "123456789"
POOL_ID = "bitbucket-pool"
PROVIDER_ID = "bitbucket-provider"
WORKSPACE_UUID = "{your-bitbucket-workspace-uuid}" # from Bitbucket workspace settings
SA_EMAIL = "ci-deployer@ ${ PROJECT_ID } .iam.gserviceaccount.com"
# Create the Workload Identity Pool
gcloud iam workload-identity-pools create " ${ POOL_ID } " \
--project = " ${ PROJECT_ID } " \
--location = "global" \
--display-name = "Bitbucket Pipelines"
# Create the OIDC provider inside that pool
gcloud iam workload-identity-pools providers create-oidc " ${ PROVIDER_ID } " \
--project = " ${ PROJECT_ID } " \
--location = "global" \
--workload-identity-pool = " ${ POOL_ID } " \
--display-name = "Bitbucket OIDC" \
--issuer-uri = "https://api.bitbucket.org/2.0/workspaces/ ${ WORKSPACE_UUID } /pipelines-config/identity/oidc" \
--attribute-mapping = "google.subject=assertion.sub,attribute.repository_uuid=assertion.repositoryUuid,attribute.workspace_uuid=assertion.workspaceUuid" \
--attribute-condition = "attribute.workspace_uuid == ' ${ WORKSPACE_UUID } '"
# Allow your Bitbucket workspace's identities to impersonate the service account
gcloud iam service-accounts add-iam-policy-binding " ${ SA_EMAIL } " \
--project = " ${ PROJECT_ID } " \
--role = "roles/iam.workloadIdentityUser" \
--member = "principalSet://iam.googleapis.com/projects/ ${ PROJECT_NUMBER } /locations/global/workloadIdentityPools/ ${ POOL_ID } /attribute.workspace_uuid/ ${ WORKSPACE_UUID } "
A few subtleties worth calling out:
The issuer URI is per-workspace. Bitbucket exposes the OIDC discovery document at a workspace-scoped URL, not a Bitbucket-wide one. You can get your workspace UUID from Workspace settings → OpenID Connect in the Bitbucket UI.
The attribute condition is your security boundary. Without --attribute-condition , any Bitbucket workspace would work to ask this provider for a token. So scope your attribute condition to your workspace UUID, or better yet a repository or repositories UUIDs in prod.
workloadIdentityUser binds which identities can impersonate the SA. The principalSet:// member here means that any identity from your workspace can impersonate ci-deployer . Preferably bind on attribute.repository_uuid/<repo-uuid> instead.
Step 2 — Bitbucket-side YAML and (s)hell
Link to heading
This is the part where you write the glue.
For readers who don’t have a working example to start from, here’s a reference shape using gcloud ’s built-in credential config. This is the shorter of the two common approaches, and the one I like to use:
image : google/cloud-sdk:slim
pipelines :
branches :
main :
- step :
name : Deploy to GCP via WIF
oidc : true # makes $BITBUCKET_STEP_OIDC_TOKEN available
script :
# 1. Write the OIDC token to a file gcloud can read
- echo "$BITBUCKET_STEP_OIDC_TOKEN" > /tmp/oidc-token.txt
# 2. Create a credential config that points at the WIF provider
- |
gcloud iam workload-identity-pools create-cred-config \
"projects/${GCP_PROJECT_NUMBER}/locations/global/workloadIdentityPools/${POOL_ID}/providers/${PROVIDER_ID}" \
--service-account="${SA_EMAIL}" \
--output-file=/tmp/credential-config.json \
--credential-source-file=/tmp/oidc-token.txt
# 3. Tell gcloud (and Google client libraries) to use it
- export GOOGLE_APPLICATION_CREDENTIALS=/tmp/credential-config.json
- gcloud auth login --cred-file=/tmp/credential-config.json --quiet
# 4. Verify, then do real work
- gcloud auth list
- gsutil ls gs://${BUCKET}
GCP_PROJECT_NUMBER , POOL_ID , PROVIDER_ID , SA_EMAIL , and BUCKET are repos or deployment variables. The only real secret in this flow is the OIDC token, which Bitbucket mints for each step.
There’s a couple of gotchas that can trip you up here:
Double-check the workspace UUID — Bitbucket UUIDs include the curly braces ( {abc-123-...} ) and the provider config needs them. Weird Bitbucket formatting might be a theme in these posts.
The SA needs roles/iam.workloadIdentityUser granted to the principalSet . The error message names the SA, not the binding, in GCP’s tradition of somewhat misleading or cryptic error messages.
The pool/provider path uses the numeric project number , not the project ID string. Lots of things use the latter, so it’s an easy mistake to make.
What’s worth verifying against current docs
Link to heading
OIDC, WIF, and gcloud’s credential-helper commands have all evolved. Before copy-pasting, cross-check against:
Bitbucket Pipelines: Integrate Pipelines with resource servers using OIDC
GCP Workload Identity Federation: Configure with OIDC
YMMV. The shape should be good, make sure the syntax and all that are right.
Previous in series: Coming from GitHub Actions? Here’s what Bitbucket Pipelines actually feels like
Next in series: The Bitbucket API UUID Slog

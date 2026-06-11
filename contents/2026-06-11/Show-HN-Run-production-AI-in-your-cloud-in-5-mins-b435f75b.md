---
source: "https://github.com/dagploy/dax"
hn_url: "https://news.ycombinator.com/item?id=48489842"
title: "Show HN: Run production AI in your cloud in 5 mins"
article_title: "GitHub - dagploy/dax: AIOps Infra for deploy and manage self-hosted local AI in your own cloud. Vibe coding and AI agents compatible. · GitHub"
author: "yodi"
captured_at: "2026-06-11T13:27:02Z"
capture_tool: "hn-digest"
hn_id: 48489842
score: 1
comments: 0
posted_at: "2026-06-11T13:06:58Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Run production AI in your cloud in 5 mins

- HN: [48489842](https://news.ycombinator.com/item?id=48489842)
- Source: [github.com](https://github.com/dagploy/dax)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T13:06:58Z

## Translation

タイトル: HN を表示: クラウドで本番 AI を 5 分で実行する
記事のタイトル: GitHub - dagploy/dax: 独自のクラウドでセルフホスト型ローカル AI をデプロイおよび管理するための AIOps インフラストラクチャ。 Vibe コーディングと AI エージェントに対応。 · GitHub
説明: 独自のクラウドでセルフホスト型のローカル AI を展開および管理するための AIOps インフラストラクチャ。 Vibe コーディングと AI エージェントに対応。 - ダグプロイ/ダックス
HNテキスト：こんにちは！ Dagploy.com の Yodi です。私は過去 12 か月間、誰もが自分のクラウドでセルフホスト型 AI を実行できるようにする AIOps インフラ ツールの構築に費やしました。オープンソースのモデルとツールを実行します。セットアップにはわずか 5 分ほどかかります。以前、私は 6,000 万人のユーザーにサービスを提供する AI ショートビデオ レコメンデーション システムを導入しました。 GPU をセットアップするための AI インフラストラクチャを構築し、クォータ、ネットワークの問題、ディスクの問題などに対処するのに 8 か月かかりました。このつらい経験が私に Dagploy の構築を思い起こさせ、他の人がより早く始められるようにしました。紹介ビデオ: https://www.youtube.com/watch?v=BCgOexZ_fr8 ウェブサイト: https://www.dagploy.com フィードバック、ご質問、ご提案をお待ちしております。

記事本文:
GitHub - dagploy/dax: 独自のクラウドでセルフホスト型ローカル AI をデプロイおよび管理するための AIOps インフラストラクチャ。 Vibe コーディングと AI エージェントに対応。 · GitHub
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
ダグプロイ
/
ダックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット dagploy_dax dagploy_dax docker_build docker_build サンプル サンプル スクリプト スクリプト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md ライセンス ライセンス README.md README.md Contributing.md Contributing.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
YAML ベースのワークフローを大規模に使用して、独自のクラウド内に AI インフラストラクチャを構築して運用します。実際の運用環境で推論、トレーニング、AI エージェントの利用を自動化します。スポット インスタンス、GPU クォータを認識したリージョン切り替え、バイブコーディングのカスタマイズなどをサポートします。
前提条件:
クラウド プロジェクトで GPU クォータをできるだけ早く有効にします。承認には最大 48 時間かかる場合があります。 GPU クォータがないと、GPU VM の起動が GPUS_ALL_REGIONS クォータ エラーで失敗する可能性があります。容量の問題を軽減するには、複数のリージョンにわたって GPU クォータを有効にします。
この手順では、パブリック IP を使用せずに、デフォルトのネットワークに DAX をインストールします。 VM 内からインターネット アクセスを可能にするには、クラウド NAT が必要です。 gcloud compute ssh <instance_name> を使用して VM にログインできます。
1. サービス アカウントを作成します (約 30 秒)
サービス アカウントは、インスタンス、ファイアウォール、その他のサービスをプロビジョニングするための所有者/実行者として必要です。このスクリプトを実行して設定します。 gcloud が端末にインストールされ、認証されていることを確認してください。
bash <(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_create_service_account.sh)
必要な権限を持って作成された新しいサービス アカウントが表示されます。
「roles/compute.instanceAdmin.v1」
"roles/compute.securityAdmin"
「役割/iam.serviceAccountUser」
「roles/artifactregistry.writer」
「ロール/ストレージ.オブジェクトユーザー」
「roles/compute.loadBalancerAdmin」
「役割/dns.adm」

で」
"roles/secretmanager.secretAccessor"
これにより、VM コンピューティングのプロビジョニングに使用されるローカル サービス アカウント JSON とシークレット dax-service-account-key の両方が生成されます。
DAX サーバー VM にはパブリック IP がありません。パッケージをダウンロードするためのインターネット アクセスを可能にするために、クラウド NAT を作成します。
bash <(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_install_cloud_nat.sh)
3. DAX VM サービスを作成します (約 30 秒)
以下のコマンドを実行します。 YOUR-SERVICE-ACCOUNT-EMAIL を、前に作成したサービス アカウントの電子メール アドレスに置き換えます。これは、生成されたサービス アカウントの JSON ファイルで確認できます。
--metadata enable-oslogin=TRUE を使用して、企業の Google アカウントなどの OS ログインへのアクセスを制限します。標準の SSH ベースのアクセスには、enable-oslogin=FALSE を使用します。
gcloud compute インスタンスが dax を作成 \
--service-account=あなたのサービスアカウントのメールアドレス \
--scopes=クラウド プラットフォーム \
--zone=us-central1-a \
--machine-type=e2-custom-4-8192 \
--ブートディスクサイズ=60GB \
--boot-disk-type=pd-balance \
--image-family=debian-12 \
--image-project=debian-cloud \
--network=デフォルト \
--subnet=デフォルト \
--アドレスなし\
--tags=ダックス\
--metadata Enable-oslogin=FALSE,startup-script= ' #!/bin/bash
セット -e
apt-getアップデート
DEBIAN_FRONTEND=非対話型 apt-get install -y git
'
4. DAX をインストールします (約 3 分)
gcloud compute ssh dax を使用してマシンに SSH 接続し、インストール手順を実行します。 DAX はユーザー フォルダーにインストールされます。
sudo bash -c "$(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_install.sh)"
おめでとうございます。DAX はすでにインストールされ、実行されています 🎉
サービスを確認するには、
sudo -iu dax -- tmuxattach -t dax
💻 CLIで接続する
プロビジョニングは、curl または CLI を介して DAX サーバーに指示できます。 SSH トンネリング経由でラップトップ/コンピューターを DAX サーバーに接続します。
の

詳細な手順はここで読むことができます:
DAX CLI のインストール (examples/project/dax-cli)
このコマンドを実行して、パブリック インターネット経由で安全に接続を確立します。ポートは 2 つあります: 8001 (DAX) と 8080 (Hatchet 経由のダッシュボード)
gcloud compute ssh dax --zone us-central1-a --tunnel-through-iap -- -L 8001:localhost:8001 -L 8080:localhost:8080
https://localhost:8080 経由でダッシュボードにアクセスするか、https://localhost:8001 へのカール プロビジョニングを介してアクセスできます。
クラウドで GPT OSS 20B を最初から実行するのにかかる時間はわずか 15 分です。
まず、Docker イメージとモデル (合計約 100 GB) をキャッシュしてから、キャッシュからワークロードを起動します。
このキャッシュ メカニズムにより、大きなファイルがネットワーク経由でダウンロードされる際の GPU のアイドル時間が回避されるため、起動時間を最大 80% 短縮し、コストを削減できます。
dax run download_docker vllm/vllm-openai:nightly,ghcr.io/open-webui/open-webui:main --images vllm-lib --image-size 100
ステップ 2: Huggingface から GPTOSS 20B をキャッシュする
dax run download_hf openai/gpt-oss-20b --image-size 50
ステップ 3: 推論を実行する
dax run create_vm_inference --stack-name gptoss --config-json '{"images":["models--openai--gpt-oss-20b","vllm-lib"]}' --model openai/gpt-oss-20b
またはそれより長いバージョン
dax run create_vm_inference --stack-name gptoss --config-json '{"images":["models--openai--gpt-oss-20b","vllm-lib"]}' --model https://huggingface.co/openai/gpt-oss-20b
ラップトップ/コンピュータからトンネリング経由でアクセスします
gcloud compute ssh gptoss -- -L 8000:localhost:8000 -L 8081:localhost:8080
これにより、http://localhost:8081 経由で openwebui が転送され、http://localhost:8000 経由で VLLM API が転送されます。
プロパティ [project] は環境設定 [CLOUDSDK_CORE_PROJECT ] によってオーバーライドされます。
これは DAX の問題ではなく、ローカル マシンの問題です。
解決策: CLOUDSDK_CORE_PROJECT の設定を解除します
2. 起動エラー: stack_name project_name Program work_dir opt

s
local_workspace.py」、create_or_select_stack の 1011 行目
raise ValueError(f"予期しない引数: {' '.join(args)}")
値エラー: 予期しない引数: stack_name project_name プログラム work_dir opts
pulumi_yaml/Pulumi.yaml で定義されているプロジェクト パスの値が正しいことを確認してください。
.env 内のものがすでに正しいかどうかを確認してください。
config/env/dev.yaml を確認し、プロジェクトとサービス アカウントの値が正しいことを確認してください。
プロジェクト名: GCP_PROJECT_NAME
gcp:プロジェクト: GCP_PROJECT_NAME
gcp:serviceAccount: SERVICE_ACCOUNT_EMAIL_ADDRESS
3. エラーネットワーク
インターネットへのアクセスに問題がある場合:
W: https://deb.debian.org/debian/dists/bullseye/InRelease の取得に失敗しました。` への接続を開始できません。
debian.map.fastly.net:443 (2a04:4e42::644)。 - 接続 (101: ネットワークに到達できません) への接続を開始できません
debian.map.fastly.net:443 (2a04:4e42:200::644)。 - 接続 (101: ネットワークに到達できません) への接続を開始できません
または、COS NVIDIA ドライバーのインストールが停止します
イメージ「us.gcr.io/cos-cloud/cos-gpu-installer:v2.7.2」がローカルに見つかりません
docker: デーモンからのエラー応答: Get "https://us.gcr.io/v2/": net/http: 接続の待機中にリクエストがキャンセルされました (ヘッダーの待機中に Client.Timeout を超過しました)。
「docker run --help」を参照してください。
エラー: GPU ドライバーのインストールに失敗しました: GPU ドライバーをインストールできませんでした: インストーラー 'us.gcr.io/cos-cloud/cos-gpu-installer:v2.7.2' を使用してインストールを完了できませんでした: 終了ステータス 125
これは、クラウド NAT が機能していない、プロキシが正しい方法で設定されていない、またはサブネットが Google プライベート アクセス許可を付与していないことを意味します。クラウド NAT はグローバルではなく地域レベルで機能します。 NAT とサブネットを有効にするには、これを実行します
bash スクリプト/gcp_install_cloud.nat.sh
DAX クラウド サービス
私たちはクラウドサービスとAIインフラエージェントに取り組んでいます。興味があれば、できます

待機リストに参加するか、カスタムのお問い合わせについてお問い合わせください: https://www.dagploy.com/contact
ソースから DAX を構築すること、または改善に貢献することについては、CONTRIBUTING.md にアクセスしてください。
DAX は、Apache License 2.0 に基づいてリリースされています。全文についてはライセンスを参照してください。
研究で DAX を使用する場合は、以下を引用してください。
@その他{ダックス、
title = {DAX: コードとしての AIOps インフラ},
著者 = {DAGPLOY}、
年 = {2026}、
URL = {https://github.com/dagploy/dax}
}
について
AIOps Infra は、独自のクラウドでセルフホスト型のローカル AI を展開および管理します。 Vibe コーディングと AI エージェントに対応。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
リリース 1.1.6
最新の
2026 年 6 月 11 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AIOps Infra for deploy and manage self-hosted local AI in your own cloud. Vibe coding and AI agents compatible. - dagploy/dax

Hello! I’m Yodi from Dagploy.com. I spent the past 12 months building AIOps Infra tools to help anyone run self-hosted AI in their own cloud. Run any open source models and tools. Setup just takes about 5 minutes. Previously, I deployed a AI short-video recommendation system serve 60 million users. It took me 8 months to build the AI Infra for setup GPUs, deal with quota, network problem, disk issue and many more. That painful experience inspired me to build Dagploy, so others can get started much faster. Intro video: https://www.youtube.com/watch?v=BCgOexZ_fr8 Website: https://www.dagploy.com I’d love to hear your feedback, questions, or suggestions.

GitHub - dagploy/dax: AIOps Infra for deploy and manage self-hosted local AI in your own cloud. Vibe coding and AI agents compatible. · GitHub
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
dagploy
/
dax
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit dagploy_dax dagploy_dax docker_build docker_build examples examples scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md LICENSE LICENSE README.md README.md contributing.md contributing.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Build and operate AI infrastructure inside your own cloud with YAML-based workflows at scale. Automate inference, training, and AI agent harnesses in real production environments. Supports spot instances, GPU quota-aware region switching, vibe-coding customization, and more.
Pre-requisites:
Enable GPU quota in your cloud project as early as possible. Approval can take up to 48 hours. Without GPU quota, launching GPU VMs may fail with a GPUS_ALL_REGIONS quota error. To reduce capacity issues, enable GPU quota across multiple regions.
This step installs DAX on the default network without a public IP. Cloud NAT is required to enable internet access from inside the VM. You can log in to the VM with gcloud compute ssh <instance_name> .
1. Create a Service Account (~30 secs)
A service account is required as the owner/executor for provisioning instances, firewalls, and other services. Run this script to set it up. Make sure gcloud is installed and authenticated in your terminal.
bash <(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_create_service_account.sh)
You will see the new service account created with required permission:
"roles/compute.instanceAdmin.v1"
"roles/compute.securityAdmin"
"roles/iam.serviceAccountUser"
"roles/artifactregistry.writer"
"roles/storage.objectUser"
"roles/compute.loadBalancerAdmin"
"roles/dns.admin"
"roles/secretmanager.secretAccessor"
This will produce both local service account JSON and secret dax-service-account-key that will use for provisioning any VM compute.
DAX server VM will have no public IP. To enable internet access for downloading packages, we create a cloud NAT
bash <(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_install_cloud_nat.sh)
3. Create DAX VM service (~30 secs)
Run the command below. Replace YOUR-SERVICE-ACCOUNT-EMAIL with the service account email address you created earlier. You can find it in the generated service account JSON file.
Use --metadata enable-oslogin=TRUE to restrict access to OS Login, such as a corporate Google account. Use enable-oslogin=FALSE for standard SSH-based access.
gcloud compute instances create dax \
--service-account=YOUR-SERVICE-ACCOUNT-EMAIL \
--scopes=cloud-platform \
--zone=us-central1-a \
--machine-type=e2-custom-4-8192 \
--boot-disk-size=60GB \
--boot-disk-type=pd-balanced \
--image-family=debian-12 \
--image-project=debian-cloud \
--network=default \
--subnet=default \
--no-address \
--tags=dax \
--metadata enable-oslogin=FALSE,startup-script= ' #!/bin/bash
set -e
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y git
'
4. Install DAX (~3 minutes)
SSH into the machine with gcloud compute ssh dax and run the installation step. DAX will be installed in your user folder.
sudo bash -c "$(curl -fsSL https://raw.githubusercontent.com/dagploy/dax/refs/heads/main/scripts/gcp_install.sh)"
Congrats, now DAX already installed and running 🎉
You can check the service with
sudo -iu dax -- tmux attach -t dax
💻 Connect with CLI
Any provisioning can be instructed to DAX server via curl or CLI. Connect your laptop/computer with DAX server via SSH tunnelling.
The detailed steps can be read here:
Install DAX CLI (examples/project/dax-cli)
Run this command to establish connection securely over public internet. There are two ports: 8001 (DAX) and 8080 (Dashboard via Hatchet)
gcloud compute ssh dax --zone us-central1-a --tunnel-through-iap -- -L 8001:localhost:8001 -L 8080:localhost:8080
You can access the dashboard via https://localhost:8080 or curl provisioning into https://localhost:8001
Run GPT OSS 20B in your cloud from scratch just takes 15 minutes .
Start by caching Docker images and models first — around 100GB in total — then launch the workload from the cache.
This cache mechanism can reduce startup time by up to 80% and lower costs by avoiding idle GPU time while large files are downloaded over the network.
dax run download_docker vllm/vllm-openai:nightly,ghcr.io/open-webui/open-webui:main --images vllm-lib --image-size 100
Step 2: Cache GPTOSS 20B from Huggingface
dax run download_hf openai/gpt-oss-20b --image-size 50
Step 3: Run the inference
dax run create_vm_inference --stack-name gptoss --config-json '{"images":["models--openai--gpt-oss-20b","vllm-lib"]}' --model openai/gpt-oss-20b
Or longer version
dax run create_vm_inference --stack-name gptoss --config-json '{"images":["models--openai--gpt-oss-20b","vllm-lib"]}' --model https://huggingface.co/openai/gpt-oss-20b
Access it from your laptop/computer via tunneling
gcloud compute ssh gptoss -- -L 8000:localhost:8000 -L 8081:localhost:8080
This will forwarding openwebui via http://localhost:8081 and VLLM API via http://localhost:8000
Property [project] is overridden by environment setting [CLOUDSDK_CORE_PROJECT .
This is not DAX problem, but your local machine.
The solution: unset CLOUDSDK_CORE_PROJECT
2. Error launching: stack_name project_name program work_dir opts
local_workspace.py", line 1011, in create_or_select_stack
raise ValueError(f"unexpected args: {' '.join(args)}")
ValueError: unexpected args: stack_name project_name program work_dir opts
Make sure the project path value defined in pulumi_yaml/Pulumi.yaml is correct.
Check if anything in .env is already correct.
Check on config/env/dev.yaml and make sure the value of project and service account is correct.
project_name: GCP_PROJECT_NAME
gcp:project: GCP_PROJECT_NAME
gcp:serviceAccount: SERVICE_ACCOUNT_EMAIL_ADDRESS
3. Error network
If you have problem with access to internet:
W: Failed to fetch https://deb.debian.org/debian/dists/bullseye/InRelease Cannot initiate the connection to`
debian.map.fastly.net:443 (2a04:4e42::644). - connect (101: Network is unreachable) Cannot initiate the connection to
debian.map.fastly.net:443 (2a04:4e42:200::644). - connect (101: Network is unreachable) Cannot initiate the connection to
Or COS NVIDIA Driver installation stuck
Unable to find image 'us.gcr.io/cos-cloud/cos-gpu-installer:v2.7.2' locally
docker: Error response from daemon: Get "https://us.gcr.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers).
See 'docker run --help'.
Error: Failed to install GPU driver: could not install GPU drivers: failed to complete installation using installer 'us.gcr.io/cos-cloud/cos-gpu-installer:v2.7.2': exit status 125
This means the cloud NAT not working, the proxy haven't setup in correct way or the subnet haven't granted with google private access permission. Cloud NAT works at regional level, not global. To enable NAT and subnet, run this
bash scripts/gcp_install_cloud.nat.sh
DAX Cloud Services
We are working on the cloud services and AI Infra agents. If you are interested, you can join the waiting list or contact us for custom inquiry : https://www.dagploy.com/contact
Visit CONTRIBUTING.md for information on building DAX from source or contributing improvements.
DAX is released under the Apache License 2.0. See LICENSE for the full text.
If you use DAX in your research, please cite:
@misc{dax,
title = {DAX: AIOps Infra as Code},
author = {DAGPLOY},
year = {2026},
url = {https://github.com/dagploy/dax}
}
About
AIOps Infra for deploy and manage self-hosted local AI in your own cloud. Vibe coding and AI agents compatible.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
Release 1.1.6
Latest
Jun 11, 2026
Packages
0
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

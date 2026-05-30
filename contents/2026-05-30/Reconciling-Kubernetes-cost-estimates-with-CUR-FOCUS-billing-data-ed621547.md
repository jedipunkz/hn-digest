---
source: "https://github.com/tanrikuluozlem/burn"
hn_url: "https://news.ycombinator.com/item?id=48326053"
title: "Reconciling Kubernetes cost estimates with CUR / FOCUS billing data"
article_title: "GitHub - tanrikuluozlem/burn: See what's burning your Kubernetes budget · GitHub"
author: "OzlemT"
captured_at: "2026-05-30T11:42:39Z"
capture_tool: "hn-digest"
hn_id: 48326053
score: 2
comments: 0
posted_at: "2026-05-29T17:06:10Z"
tags:
  - hacker-news
  - translated
---

# Reconciling Kubernetes cost estimates with CUR / FOCUS billing data

- HN: [48326053](https://news.ycombinator.com/item?id=48326053)
- Source: [github.com](https://github.com/tanrikuluozlem/burn)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T17:06:10Z

## Translation

タイトル: Kubernetes のコスト見積もりと CUR / FOCUS の請求データの調整
記事のタイトル: GitHub - Tanrikuluozlem/burn: Kubernetes 予算を消費しているものを確認する · GitHub
説明: Kubernetes の予算を浪費しているものを確認します。 GitHub でアカウントを作成して、tanrikuluozlem/burn の開発に貢献してください。

記事本文:
GitHub - Tanrikuluozlem/burn: Kubernetes 予算を消費しているものを確認する · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
タンリクルオズレム
/
燃やす
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
139 コミット 139 コミット .github .github アセット アセット チャート/ 書き込み チャート/ 書き込み cmd/ 書き込み cmd/ 書き込み 内部 内部 k8s k8s スクリプト スクリプト terraform terraform .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml Dockerfile Dockerfile ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Kubernetes クラスターはお金を消費しています。場所を調べてください。
導入するエージェントがありません。ダッシュボードを保守する必要はありません。 YAML を設定する必要はありません。インストールして実行するだけです。
ゼロセットアップ — brew install 、1 つのコマンドを実行すると、答えが得られます。クラスター エージェント、永続ストレージ、構成ファイルはありません。
完全なコストをカバー — コンピューティング、ストレージ、ロード バランサー、GPU のコストをリアルタイムのクラウド価格でカバーします。
AI を活用 — わかりやすい英語で質問し、コピーアンドペーストできる kubectl コマンドを取得します。
Slack ネイティブ — /burn でインスタントコストレポートを作成します。 /burn は AI 分析のために「...」を要求します。
クラウド + オンプレミス — AWS EKS、Azure AKS、GCP GKE、オンプレミス クラスターで動作します。
スポットの準備状況 — リアルタイムの割引と中断率を使用して、どのワークロードをスポット インスタンスに安全に移動できるかを特定します。
Ingress LB 検出 — ホスト名の重複排除を使用して、Services リソースと Ingress リソースの両方からロード バランサーを検出します。
時間認識 — --期間 7d は、ポイントインタイムのスナップショットの代わりに週平均を表示します。
#自作
brew install Tanrikuluozlem/burn/burn
# アップグレード
醸造アップグレード タンリクルオスレム/バーン/バーン
# バイナリ
VERSION= $(curl -s https://api.github.com/repos/tanrikuluozlem/burn/releases/latest | grep タグ名 | Cut -d ' " ' -f4 | tr -d ' v ' ) && \
カール -L " https://github.com/tanrikuluozlem/burn/releases/latest/download/burn_ ${VERSION} _ $( una

私 -s | tr ' [:upper:] ' ' [: lower:] ' ) _ $( uname -m | sed ' s/x86_64/amd64/;s/aarch64/arm64/ ' ) .tar.gz " | tar xz
# ドッカー
docker pull ghcr.io/tanrikuluozlem/burn:latest
# ヘルム
git clone https://github.com/tanrikuluozlem/burn.git
Helm インストール burn ./burn/charts/burn
# 行く
github.com/tanrikuluozlem/burn/cmd/burn@latest をインストールしてください
macOS: Gatekeeper の警告が表示された場合は、次を実行します: sudo xattr -d com.apple.quarantine $(this burn)
# コストの内訳 (Prometheus なし)
燃焼分析
# Prometheus を使用 (Prometheus URL を渡します)
書き込み分析 --prometheus http://prometheus:9090
# 7 日間の平均
書き込み分析 --prometheus http://prometheus:9090 --period 7d
# 名前空間にドリルダウンする
書き込み分析 --prometheus http://prometheus:9090 --namespace argocd
# スポットの準備状況
書き込み分析 --prometheus http://prometheus:9090 --spot
スポット対応
インスタンス タイプごとのリアルタイムのスポット割引と中断率。
クラスター全体または名前空間固有の推奨事項を取得します。
書き込み分析 --prometheus http://prometheus:9090 --period 7d --ai
burn 分析 --prometheus http://prometheus:9090 --namespace app-backend --ai
burn ask --prometheus http://prometheus:9090 " なぜ argocd はそんなに高価なのでしょうか? 」
例: burnanalyze --namespace app-backend --period 7d --ai
名前空間: app-backend (3 ポッド、$17.19/月)
─────────────
POD CPU 要求→使用済み MEM 要求→使用済みコスト/MO
app-backend-deploy-0001 200m → <1m 256Mi → 9Mi $5.73
app-backend-deploy-0002 200m → <1m 256Mi → 9Mi $5.73
app-backend-deploy-0003 200m → <1m 256Mi → 128Mi $5.73
推奨事項
──────────
app-backend 名前空間の料金は 3 つのポッドで月額 17.19 ドルですが、CPU 効率は
は ~0.1% と非常に低い — ポッドは毎回 200m CPU を要求します

p95の使用法
0.31m未満です。
[!!] 1. p95 データを使用した適切なサイズの CPU リクエスト
app-backend-deploy-0001: p95 CPU は 0.22m → 1m を推奨 (1.5x p95)
app-backend-deploy-0002: p95 CPU は 0.30m → 1m を推奨 (1.5x p95)
app-backend-deploy-0003: p95 MEM は 128Mi (50% eff) — そのままにしておきます
$ kubectl set リソース デプロイメント app-backend -n app-backend \
--requests=cpu=1m、メモリ=14Mi --limits=cpu=200m、メモリ=256Mi
[!!] 2. app-backend-ingress LB ($19.71/月) は名前空間よりもコストがかかります
ロード バランサーだけでも、月額 17.19 ドルのコンピューティング コストを超えます。
内部専用の場合は、ClusterIP に切り替えて LB コストを排除します。
$ kubectl patch svc app-backend-ingress -n app-backend \
-p '{"仕様": {"タイプ": "ClusterIP"}}'
[!] 3. 推奨モードで VPA を有効にする
継続的な p95 追跡により、オーバープロビジョニングの再発を防ぎます。
$ kubectl apply -f vpa-app-backend.yaml
わかりやすい英語で質問する
ANTHROPIC_API_KEY 環境変数が必要です。
burnserve --port 8080 --prometheus http://prometheus:9090 --period 7d
コマンド
得られるもの
/バーン
フルコストレポート - ノード、名前空間、アイドルコスト、LB、ストレージ
/burn ns argocd
名前空間のポッドレベルの内訳
/burn 「最大の無駄は何ですか?」と尋ねます。
kubectlコマンドによるAI解析
https://api.slack.com/apps で Slack アプリを作成します
スラッシュコマンドを追加: /burn → サーバー URL + /slack をポイントします
SLACK_SIGNING_SECRET および ANTHROPIC_API_KEY 環境変数を設定する
サーバーを公開します (例: テスト用の ngrok、運用用のロード バランサー)
Burn はオンプレミスおよび GPU クラスターで動作します。独自のリソース レートを設定します。
書き込み分析 \
--cpu-価格 0.05 \
--ラム価格 0.008 \
--GPU-価格 3.00 \
--ストレージ価格 0.10
カスタム価格設定を行わない場合、クラウドと同等の料金がデフォルトとして使用されます。
Kubernetes API → ノード、ポッド、PVC、サービス、イングレス
Prometheus → 実際の CPU とメモリ

使用法 (オプション)
クラウドの価格 → 実際の VM、ストレージ、GPU の価格 (AWS、Azure、GCP)
↓
コストエンジン → コンピューティング、ストレージ、ロードバランサー、GPU、アイドル検出
↓
CLI / Slack / AI の推奨事項
価格設定ソース
優先順位
ソース
いつ
1
AWS/Azure 料金設定 API
AWS 認証情報が利用可能 — リアルタイム、リージョン対応
2
埋め込み価格DB
認証情報なし — 600 個以上の AWS、300 個以上の Azure インスタンス、毎週更新
3
静的フォールバック
不明なインスタンス タイプ — インスタンス ファミリーに基づく推定
ストレージとロード バランサーのコストは、利用可能な場合は静的フォールバックを使用してクラウド API から取得されます。従量制料金 (データ処理、LCU) はトラフィック量に依存するため、料金には含まれません。 GPU ノードは自動的に検出され、比率に基づいたコスト分割によって価格が決定されます。
git clone https://github.com/tanrikuluozlem/burn.git
Helm install burn ./burn/charts/burn \
--set prometheus.url=http://prometheus:9090 \
--set スケジュール = " 0 9 * * 1-5 "
CronJob (毎日の Slack レポート)
apiVersion : バッチ/v1
種類 : CronJob
メタデータ:
名前 : バーンレポート
仕様：
スケジュール：「0 9 * * 1-5」
ジョブテンプレート:
仕様：
テンプレート:
仕様：
コンテナ:
- 名前 : バーン
画像 : ghcr.io/tanrikuluozlem/burn:latest
引数:
- 分析する
- --プロメテウス
- http://prometheus-server.monitoring:80
- --ピリオド
- 7日
- --ai
- --たるみ
環境:
- 名前 : ANTHROPIC_API_KEY
値から:
秘密鍵参照 :
名前 : バーンシークレット
キー: anthropic-API-key
- 名前 : SLACK_WEBHOOK_URL
値から:
秘密鍵参照 :
名前 : バーンシークレット
キー：slack-webhook-url
restartPolicy : OnFailure
構成
変数
説明
必須
ANTHROPIC_API_KEY
クロードAPIキー
--ai、尋ねる、奉仕する
SLACK_WEBHOOK_URL
Slack Webhook URL
--たるみ
SLACK_SIGNING_SECRET
Slack アプリ署名の秘密
奉仕する
旗
説明
--CPU価格
1 時間あたりのコアあたりの CPU コスト (オンプレミス)
--ラム価格
1 時間あたり GiB あたりの RAM コスト (オンプレ

m)
--gpu-価格
1 時間あたりのユニットあたりの GPU コスト (オンプレミス)
--ストレージ価格
1 か月あたりの GiB あたりのストレージ料金 (オンプレミス)
--スポット
スポットインスタンスの準備状況の詳細を表示
クラウド クラスターでは実際の価格設定が自動的に使用されます。これらのフラグは、クラウド プロバイダーから価格を入手できないオンプレミス クラスター用です。
make build # バイナリをビルドする
make test # テストを実行する
make lint # リンターを実行する
ライセンス
Apache 2.0 — 詳細については、「ライセンス」を参照してください。
Kubernetes の予算を浪費しているものを確認する
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
16
v0.3.5
最新の
2026 年 5 月 25 日
+ 15 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

See what's burning your Kubernetes budget. Contribute to tanrikuluozlem/burn development by creating an account on GitHub.

GitHub - tanrikuluozlem/burn: See what's burning your Kubernetes budget · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
tanrikuluozlem
/
burn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
139 Commits 139 Commits .github .github assets assets charts/ burn charts/ burn cmd/ burn cmd/ burn internal internal k8s k8s scripts scripts terraform terraform .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Your Kubernetes cluster is burning money. Find out where.
No agent to deploy. No dashboard to maintain. No YAML to configure. Just install and run.
Zero setup — brew install , run one command, get answers. No cluster agent, no persistent storage, no config files.
Full cost coverage — Compute, storage, load balancers, and GPU costs with real-time cloud pricing.
AI-powered — Ask questions in plain English, get kubectl commands you can copy-paste.
Slack-native — /burn for instant cost reports. /burn ask "..." for AI analysis.
Cloud + on-prem — Works with AWS EKS, Azure AKS, GCP GKE, and on-premise clusters.
Spot readiness — Identifies which workloads can safely move to spot instances with real-time discount and interruption rate.
Ingress LB detection — Detects load balancers from both Services and Ingress resources, with hostname deduplication.
Time-aware — --period 7d for weekly averages instead of point-in-time snapshots.
# Homebrew
brew install tanrikuluozlem/burn/burn
# Upgrade
brew upgrade tanrikuluozlem/burn/burn
# Binary
VERSION= $( curl -s https://api.github.com/repos/tanrikuluozlem/burn/releases/latest | grep tag_name | cut -d ' " ' -f4 | tr -d ' v ' ) && \
curl -L " https://github.com/tanrikuluozlem/burn/releases/latest/download/burn_ ${VERSION} _ $( uname -s | tr ' [:upper:] ' ' [:lower:] ' ) _ $( uname -m | sed ' s/x86_64/amd64/;s/aarch64/arm64/ ' ) .tar.gz " | tar xz
# Docker
docker pull ghcr.io/tanrikuluozlem/burn:latest
# Helm
git clone https://github.com/tanrikuluozlem/burn.git
helm install burn ./burn/charts/burn
# Go
go install github.com/tanrikuluozlem/burn/cmd/burn@latest
macOS: If you see a Gatekeeper warning, run: sudo xattr -d com.apple.quarantine $(which burn)
# Cost breakdown (without Prometheus)
burn analyze
# With Prometheus (pass your Prometheus URL)
burn analyze --prometheus http://prometheus:9090
# 7-day average
burn analyze --prometheus http://prometheus:9090 --period 7d
# Drill into a namespace
burn analyze --prometheus http://prometheus:9090 --namespace argocd
# Spot readiness
burn analyze --prometheus http://prometheus:9090 --spot
Spot readiness
Real-time spot discount and interruption rate per instance type.
Get cluster-wide or namespace-specific recommendations:
burn analyze --prometheus http://prometheus:9090 --period 7d --ai
burn analyze --prometheus http://prometheus:9090 --namespace app-backend --ai
burn ask --prometheus http://prometheus:9090 " why is argocd so expensive? "
Example: burn analyze --namespace app-backend --period 7d --ai
NAMESPACE: app-backend (3 pods, $17.19/mo)
──────────────────────────────────
POD CPU REQ→USED MEM REQ→USED COST/MO
app-backend-deploy-0001 200m → <1m 256Mi → 9Mi $5.73
app-backend-deploy-0002 200m → <1m 256Mi → 9Mi $5.73
app-backend-deploy-0003 200m → <1m 256Mi → 128Mi $5.73
RECOMMENDATIONS
───────────────
The app-backend namespace costs $17.19/mo across 3 pods, but CPU efficiency
is critically low at ~0.1% — pods request 200m CPU each while p95 usage
is under 0.31m.
[!!] 1. Rightsize CPU Requests using p95 data
app-backend-deploy-0001: p95 CPU is 0.22m → recommend 1m (1.5x p95)
app-backend-deploy-0002: p95 CPU is 0.30m → recommend 1m (1.5x p95)
app-backend-deploy-0003: p95 MEM is 128Mi (50% eff) — leave as-is
$ kubectl set resources deployment app-backend -n app-backend \
--requests=cpu=1m,memory=14Mi --limits=cpu=200m,memory=256Mi
[!!] 2. app-backend-ingress LB ($19.71/mo) costs more than the namespace
The load balancer alone exceeds the $17.19/mo compute cost.
If internal-only, switch to ClusterIP to eliminate the LB cost.
$ kubectl patch svc app-backend-ingress -n app-backend \
-p '{"spec": {"type": "ClusterIP"}}'
[!] 3. Enable VPA in Recommend Mode
Prevent over-provisioning from recurring with continuous p95 tracking.
$ kubectl apply -f vpa-app-backend.yaml
Ask questions in plain English
Requires ANTHROPIC_API_KEY environment variable.
burn serve --port 8080 --prometheus http://prometheus:9090 --period 7d
Command
What you get
/burn
Full cost report — nodes, namespaces, idle cost, LB, storage
/burn ns argocd
Pod-level breakdown for a namespace
/burn ask "what is the single biggest waste?"
AI analysis with kubectl commands
Create a Slack App at https://api.slack.com/apps
Add Slash Command: /burn → point to your server URL + /slack
Set SLACK_SIGNING_SECRET and ANTHROPIC_API_KEY environment variables
Expose the server (e.g., ngrok for testing, load balancer for production)
Burn works with on-premise and GPU clusters. Set your own resource rates:
burn analyze \
--cpu-price 0.05 \
--ram-price 0.008 \
--gpu-price 3.00 \
--storage-price 0.10
Without custom pricing, cloud-equivalent rates are used as defaults.
Kubernetes API → nodes, pods, PVCs, services, ingresses
Prometheus → actual CPU & memory usage (optional)
Cloud Pricing → real VM, storage, and GPU prices (AWS, Azure, GCP)
↓
Cost Engine → compute, storage, load balancers, GPU, idle detection
↓
CLI / Slack / AI Recommendations
Pricing sources
Priority
Source
When
1
AWS/Azure pricing API
AWS credentials available — real-time, region-aware
2
Embedded pricing DB
No credentials — 600+ AWS, 300+ Azure instances, updated weekly
3
Static fallback
Unknown instance type — estimates based on instance family
Storage and load balancer costs are fetched from cloud APIs when available, with static fallbacks. Usage-based charges (data processing, LCU) depend on traffic volume and are not included. GPU nodes are detected automatically and priced via ratio-based cost splitting.
git clone https://github.com/tanrikuluozlem/burn.git
helm install burn ./burn/charts/burn \
--set prometheus.url=http://prometheus:9090 \
--set schedule= " 0 9 * * 1-5 "
CronJob (daily Slack reports)
apiVersion : batch/v1
kind : CronJob
metadata :
name : burn-report
spec :
schedule : " 0 9 * * 1-5 "
jobTemplate :
spec :
template :
spec :
containers :
- name : burn
image : ghcr.io/tanrikuluozlem/burn:latest
args :
- analyze
- --prometheus
- http://prometheus-server.monitoring:80
- --period
- 7d
- --ai
- --slack
env :
- name : ANTHROPIC_API_KEY
valueFrom :
secretKeyRef :
name : burn-secrets
key : anthropic-api-key
- name : SLACK_WEBHOOK_URL
valueFrom :
secretKeyRef :
name : burn-secrets
key : slack-webhook-url
restartPolicy : OnFailure
Configuration
Variable
Description
Required for
ANTHROPIC_API_KEY
Claude API key
--ai , ask , serve
SLACK_WEBHOOK_URL
Slack webhook URL
--slack
SLACK_SIGNING_SECRET
Slack app signing secret
serve
Flag
Description
--cpu-price
CPU cost per core per hour (on-prem)
--ram-price
RAM cost per GiB per hour (on-prem)
--gpu-price
GPU cost per unit per hour (on-prem)
--storage-price
Storage cost per GiB per month (on-prem)
--spot
Show spot instance readiness details
Cloud clusters use real pricing automatically. These flags are for on-premise clusters where pricing is not available from a cloud provider.
make build # Build binary
make test # Run tests
make lint # Run linter
License
Apache 2.0 — See LICENSE for details.
See what's burning your Kubernetes budget
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
16
v0.3.5
Latest
May 25, 2026
+ 15 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

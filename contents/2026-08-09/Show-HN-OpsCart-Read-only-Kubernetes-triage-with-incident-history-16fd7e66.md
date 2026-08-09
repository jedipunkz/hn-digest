---
source: "https://github.com/opscart/opscart-k8s-watcher"
hn_url: "https://news.ycombinator.com/item?id=49227829"
title: "Show HN: OpsCart – Read-only Kubernetes triage with incident history"
article_title: "GitHub - opscart/opscart-k8s-watcher: Kubernetes operational triage dashboard. Surfaces CrashLoops, security gaps, orphaned resources, and cost waste — tells you what to fix first. Read-only, no agents, no cloud credentials. · GitHub"
author: "opscart"
captured_at: "2026-08-09T02:58:43Z"
capture_tool: "hn-digest"
hn_id: 49227829
score: 2
comments: 0
posted_at: "2026-08-09T02:24:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpsCart – Read-only Kubernetes triage with incident history

- HN: [49227829](https://news.ycombinator.com/item?id=49227829)
- Source: [github.com](https://github.com/opscart/opscart-k8s-watcher)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T02:24:57Z

## Translation

タイトル: HN を表示: OpsCart – インシデント履歴を含む読み取り専用 Kubernetes トリアージ
記事のタイトル: GitHub - opscart/opscart-k8s-watcher: Kubernetes 運用トリアージ ダッシュボード。クラッシュループ、セキュリティギャップ、孤立したリソース、コストの無駄を表面化することで、最初に何を修正すべきかを示します。読み取り専用、エージェントなし、クラウド認証情報なし。 · GitHub
説明: Kubernetes 運用優先順位付けダッシュボード。クラッシュループ、セキュリティギャップ、孤立したリソース、コストの無駄を表面化することで、最初に何を修正すべきかを示します。読み取り専用、エージェントなし、クラウド認証情報なし。 - opscart/opscart-k8s-watcher

記事本文:
GitHub - opscart/opscart-k8s-watcher: Kubernetes 運用トリアージ ダッシュボード。クラッシュループ、セキュリティギャップ、孤立したリソース、コストの無駄を表面化することで、最初に何を修正すべきかを示します。読み取り専用、エージェントなし、クラウド認証情報なし。 · GitHub
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
オプスカート
/
opscart-k8s-watcher
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
238 コミット 238 コミット cmd cmd デプロイ デプロイ docs docs exa

mples/failure-lab例/failure-lab helm/opscart-watcherhelm/opscart-watcher pkg pkg scripts scripts .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Kubectl はリソースを表示します。レンズは状態を示します。 OpsCart は、注目に値するものを示します。
読み取り専用 · エージェントなし · 資格情報不要のコア · 30 秒で展開
このクラスターの中で今注目すべきものは何でしょうか?
kubectl → リソースを表示
Grafana → メトリクスを表示
レンズ → クラスターの状態を表示
OpsCart → 最初に注目すべきものを表示します
OpsCart の動作を確認する
ダッシュボード — 優先順位付けされた運用インテリジェンス
CLI — 証拠に裏付けられた Kubernetes トリアージ
OpsCart には現在 2 つの独立したスキャンパスがあります。
CLI は、opscart-scan トリアージが呼び出されるたびにワンショット スキャンを実行します。
ダッシュボードは、独自の定期的なタイマー ループからスキャンを実行します。
各パスは Kubernetes の状態を個別に読み取り、独自の実行環境内で操作履歴を維持します。これらは、基礎となるスキャナー、モデル、およびストレージ実装の一部を共有しますが、一部の分類と表示動作は 2 つのパス間で依然として異なります。その結果、CLI とダッシュボードでは、特定の結果が異なる方法で分類またはグループ化される場合があります。
将来の設計では、共有の正規分類コンポーネントを通じて、より多くのインシデント セマンティクスを調整する可能性があります。共有データベースは現時点では必要ないか、計画されていません。スキャンのタイミングと保持される履歴は、CLI とダッシュボード間で正当に異なる場合があります。
どちらのパスも、継続的な Kubernetes 監視ではなく、制限されたスナップショット スキャンを使用します。これにより、コアは読み取り専用となり、操作がシンプルになります。その代償として、OpsCart はリアルタイムのイベント アラートではなく、スキャン時にトリアージを提供します。
健全なダッシュボードは必ずしも優れた状態を意味するわけではありません

クラスターです。
メトリクスは、サービスが SLO を満たしているかどうかを示します。 OpsCart は、クラッシュ ループ ワークロード、イメージ プルの失敗、特権コンテナー、ネットワーク ポリシーの欠落、接続されていないストレージ、リソースの無駄など、ダッシュボード全体で隠蔽または断片化されたままになる可能性のある運用状況を明らかにします。
これは別のアラート アグリゲーターではありません。 OpsCart は、インシデントが最初に検出されたとき、インシデントが解決され、その後再発したかどうか、再起動動作がどのように変化したか、現在影響を受けているワークロードなど、スキャン全体にわたって運用メモリを保存します。ワークロード インシデントが数週間続いているにもかかわらず、交換用のポッドはわずか 5 分しか経過していない可能性があります。 OpsCart は、新しいポッドをインシデントの ID として提示せずに、ワークロード レベルの履歴を保持します。
OpsCart は、オペレーターがトリアージ中に複数のダッシュボードを相互に関連付けることを要求するのではなく、注目に値するもの、その背後にある観測された証拠、および次のステップのための読み取り専用の調査コマンドに関する優先順位付けされたビューを表示します。
OpsCart は、ノード エージェントのデプロイやアプリケーション ワークロードの変更を行わずに、より迅速に運用の優先順位付けを行いたい、Kubernetes クラスターを管理するプラットフォーム エンジニア向けに設計されています。
helm インストール opscart-watcher ./helm/opscart-watcher \
--namespace opscart-system \
--create-namespace
kubectl port-forward -n opscart-system svc/opscart-watcher 8080:80
http://localhost:8080 を開きます
インシデント履歴はデフォルトで PVC 上に保持されます。ストレージ オプションと minikube 固有の注意事項については、表の README を参照してください。以下の生のマニフェストはクイックスタートです。 Helm チャートは正規です。
公式コンテナ イメージは、 linux/amd64 と linux/arm64 の両方をサポートします。
kubectl apply -f https://raw.githubusercontent.com/opscart/opscart-k8s-watcher/main/deploy/dashboard.yaml
kubectl port-forward -n opscart-system svc/opscart-watcher 8080:80
http://localhost:8080 を開きます

開発者向けビルド
git クローン https://github.com/opscart/opscart-k8s-watcher.git
cd opscart-k8s-watcher
go build -o opscart-dashboard ./cmd/opscart-dashboard
./opscart-dashboard --cluster my-cluster --port 8080
初回ログイン
OpsCart はデフォルトで認証を必要とします。 auth.existingSecret が
空の場合、Helm はリリース管理の Secret を作成し、生成された Secret を保存します。
ポッドの再起動とチャートのアップグレードの間の認証情報。
kubectl get Secret -n opscart-system opscart-watcher-auth \
-o jsonpath= ' {.data.username} ' | base64 --decode
エコー
kubectl get Secret -n opscart-system opscart-watcher-auth \
-o jsonpath= ' {.data.password} ' | base64 --decode
エコー
自分で資格情報を指定するには、ユーザー名とパスワードを使用してシークレットを作成します
キーと set auth.existingSecret :
kubectl シークレットの作成汎用 opscart-auth \
--from-literal=ユーザー名=管理者 \
--from-literal=password= < パスワード > \
-n opscart システム
認証:
既存のシークレット: " opscart-auth "
資格情報が構成されていないスタンドアロン実行でも、
起動時のランダムなパスワード。このパスワードはプロセスが再起動されると変更されます。
チーム展開の場合は、oauth2-proxy を使用する代わりに入力層で認証します。Azure AD、Google、GitHub、および汎用 OIDC をサポートします。完全なパターンと脅威モデルについては、「セキュリティ」を参照してください。
ワークロードを手動で中断せずに OpsCart を評価したいですか?このリポジトリには、意図的に不健全なワークロードと健全なワークロードが 8 つの名前空間に分散された自己完結型の Kubernetes 障害ラボが含まれています。
特権コンテナの検出
NetworkPolicy の対象範囲がありません
比較のための健全なワークロード
警告: ラボでは、失敗するワークロード、無効なイメージが意図的に作成されています
プル、特権コンテナー、および未使用のストレージ要求。でのみ実行します
使い捨てまたは非実稼働 Kubernetes クラスター。
。/テスト

ples/failure-lab/scripts/setup.sh
# 障害と Kubernetes イベントが発生するまで 30 ～ 60 秒かかります
opscart-scan トリアージ \
--cluster " $( kubectl config current-context ) " \
--次のステップ
# ラボが所有するすべての名前空間とワークロードを削除します
./examples/failure-lab/scripts/cleanup.sh
完全なシナリオ インベントリ、セットアップ手順、検証コマンド、クリーンアップ手順については、OpsCart Kubernetes Failure Lab を参照してください。
エントリ ポイントは KPI グリッドではなく、書面による評価です。状況ブリーフィングでは、実際に何が問題なのかを、ゲージの壁ではなく、平易な文章 (「不正検出が 7 日間クラッシュ ループしており、再起動率が加速している」) で述べています。その下に:
運用メモリ スコアボード — これまでに発生したインシデント、解決済み、再開されたインシデントの合計、現在加速しているインシデント、最も長く実行されているインシデント、最も不安定な名前空間 — OpsCart が単に観察しているだけでなく記憶しているからこそ得られる数値
修正すべき上位 5 項目 — 重大度および再起動率によってランク付けされ、それぞれにメモリ ライン (最初に検出されたもの、再起動回数、傾向) と調査するための直接リンクが含まれています。
クラスターの健全性、名前空間の健全性、セキュリティ ステータスの概要
この分野の競合他社 (Grafana、Lens、k9s) は、スキャンの間に何も覚えていないため、これを作成できません。
インシデント スコア — クラッシュ ループ、イメージ プルの失敗、セキュリティ体制、無駄、およびネットワーク ポリシーのギャップから導出される 0 ～ 100 の単一スコア。トレンド矢印と 7 ポイントのスパークラインは、クラスターが改善しているか悪化しているかを示します。
War Room — すべての重大なインシデントを 1 つのビューに表示し、重大度と再起動率によって優先順位を付けます。クラッシュ ループ、イメージ プルの失敗、プローブの失敗 (準備完了に達する前に起動/活性チェックの失敗により繰り返し強制終了されるコンテナ)、特権コンテナを検出します。各カードには、問題のタイプ、名前空間、経過年数、

再起動回数、およびすぐに実行できる kubectl コマンド。ワンクリックで完全な調査が開きます。
調査 — ワンクリックで検出から調査まで完了します。すべてのポッドレベルのインシデントには以下が含まれます。
OpsCart の評価: パターンの意味と推定調査時間 (再起動率が加速しているかどうかを含む)
インシデントのタイムライン: 運用ジャーナル — 最初の検出、再起動マイルストーン、重大度の変更、解決/再発生 — ポッドの再起動後も保持される
証拠: 重大度、最初に検出されたもの、再起動回数、状態、年齢、所有者
Blast Radius: レプリカのダウン、兄弟ポッドの健全性、ワークロードへのサービス ルーティング、イングレス エクスポージャ、名前空間全体の健全性、および顧客への影響ヒューリスティック (内部トラフィックと外部トラフィックの可能性)
推奨される調査: 高/中/低の信頼度および特定の kubectl コマンドを使用した番号付きのステップ
最近のイベント: このポッドにフィルターされた最後の 10 件のイベント
関連リソース: ポッド仕様で参照される ConfigMap、シークレット、PVC
名前空間レベルの検出結果 (保護されていない名前空間、アイドル状態の名前空間) には、適用されないポッドの調査ではなく、サンプルの NetworkPolicy と名前空間を対象とした修復手順という、独自の専用ビューが表示されます。
インシデント — 完全な記録システム: このクラスターで発生し、アクティブで解決されたすべてのインシデント。名前、名前空間、または問題の種類で検索します。重大度とステータスによってフィルタリングします。ソートとページネーション。解決されたインシデントは回復時間とともに表示されるため、現在何が壊れているかだけでなく、先週何が壊れているかを確認できます。
セキュリティ体制 — CIS Kubernetes ベンチマーク スコアリング。失敗したコントロールが最初に表示され、カテゴリごとにリスクの内訳が示され、優先順位付けされた修復アクションが表示されます。
廃棄物とドリフト — ゾンビ ポッド、ストレージのサイズと年齢を伴う孤立した PVC、ゼロレプリカのワークロード、放棄された名前空間。
コスト インテリジェンス — プロバイダー対応ワーカー n

名前空間の割り当て、組み込みの Azure 料金、およびオプションの AWS 公開料金を含む ode の見積もり。 「コスト インテリジェンス」を参照してください。
動作メモリ — OpsCart は何が起こったかを記憶します。軽量のローカル データベースは、クラスターのスナップショット、インシデントのライフサイクル (検出→マイルストーン→解決→再オープン) を追加専用のイベント ジャーナルとして追跡し、メタデータをスキャンします。トレンド矢印、スパークライン、インシデント経過時間、およびインシデントごとのタイムラインを強化します。 SQLite によってサポートされ、ポッドの再起動や Helm アンインストール後も存続する PVC 上に永続化されます。構成可能な保存期間 (デフォルトでは 90 日) により、データベースが際限なく増大することがなくなります。
認証 - デフォルトで基本認証が有効になり、無効パスはありません: 環境変数、Kubernetes シークレット、または起動時に記録される自動生成パスワード。チームの場合は、Azure AD / Google / GitHub / OIDC の oauth2-proxy を前面に配置します。上記の「初回ログイン」を参照してください。
Helm チャート — 構成可能な値、PVC ベースの永続性、読み取り専用 RBAC、および非ルート セキュリティ コンテキストを備えた完全な Helm チャート。永続化オプション、minikube の注意事項、およびすべての値については、チャートの README を参照してください。
エージェントレス — 単一のコンテナとして実行します。サイドカー、DaemonSet、ノード アクセスはありません。コア スキャンにはクラウド認証情報は必要ありません。オプションの AWS API 料金設定ではワークロード ID を使用します。
プロパティ
詳細
ベースイメージ
スクラッチ —

[切り捨てられた]

## Original Extract

Kubernetes operational triage dashboard. Surfaces CrashLoops, security gaps, orphaned resources, and cost waste — tells you what to fix first. Read-only, no agents, no cloud credentials. - opscart/opscart-k8s-watcher

GitHub - opscart/opscart-k8s-watcher: Kubernetes operational triage dashboard. Surfaces CrashLoops, security gaps, orphaned resources, and cost waste — tells you what to fix first. Read-only, no agents, no cloud credentials. · GitHub
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
opscart
/
opscart-k8s-watcher
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
238 Commits 238 Commits cmd cmd deploy deploy docs docs examples/ failure-lab examples/ failure-lab helm/ opscart-watcher helm/ opscart-watcher pkg pkg scripts scripts .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Kubectl shows resources. Lens shows state. OpsCart shows what deserves your attention.
Read-only · No agents · Credential-free core · Deploy in 30 seconds
What in this cluster deserves attention right now?
kubectl → shows resources
Grafana → shows metrics
Lens → shows cluster state
OpsCart → shows what deserves attention first
See OpsCart in action
Dashboard — prioritized operational intelligence
CLI — evidence-backed Kubernetes triage
OpsCart currently has two independent scan paths:
The CLI performs a one-shot scan whenever opscart-scan triage is invoked.
The dashboard performs scans from its own periodic timer loop.
Each path reads Kubernetes state independently and maintains operational history within its own execution environment. They share parts of the underlying scanner, model, and storage implementation, but some classification and presentation behavior still differs between the two paths. Consequently, the CLI and dashboard may classify or group certain findings differently.
A future design may align more incident semantics through shared canonical classification components. A shared database is not currently required or planned; scan timing and retained history may legitimately differ between the CLI and dashboard.
Both paths use bounded snapshot scans rather than continuous Kubernetes watches. This keeps the core read-only and operationally simple. The tradeoff is that OpsCart provides triage at scan time rather than real-time event alerting.
A healthy dashboard does not always mean a healthy cluster.
Metrics show whether services are meeting their SLOs. OpsCart surfaces operational conditions that can remain hidden or fragmented across dashboards—crash-looping workloads, image pull failures, privileged containers, missing NetworkPolicies, unattached storage, and resource waste.
This is not another alert aggregator. OpsCart preserves operational memory across scans: when an incident was first detected, whether it resolved and later reoccurred, how restart behavior changed, and which workload is currently affected. A replacement pod may be only five minutes old while the workload incident has existed for weeks. OpsCart keeps that workload-level history without presenting the new pod as the identity of the incident.
Instead of requiring operators to correlate several dashboards during triage, OpsCart presents a prioritized view of what deserves attention, the observed evidence behind it, and read-only investigation commands for the next step.
OpsCart is designed for platform engineers managing Kubernetes clusters who want faster operational triage without deploying node agents or modifying application workloads.
helm install opscart-watcher ./helm/opscart-watcher \
--namespace opscart-system \
--create-namespace
kubectl port-forward -n opscart-system svc/opscart-watcher 8080:80
open http://localhost:8080
Incident history persists on a PVC by default — see the chart README for storage options and minikube-specific notes. The raw manifest below is a quickstart; the Helm chart is canonical.
Official container images support both linux/amd64 and linux/arm64 .
kubectl apply -f https://raw.githubusercontent.com/opscart/opscart-k8s-watcher/main/deploy/dashboard.yaml
kubectl port-forward -n opscart-system svc/opscart-watcher 8080:80
open http://localhost:8080
Developer Build
git clone https://github.com/opscart/opscart-k8s-watcher.git
cd opscart-k8s-watcher
go build -o opscart-dashboard ./cmd/opscart-dashboard
./opscart-dashboard --cluster my-cluster --port 8080
First Login
OpsCart requires authentication by default. When auth.existingSecret is
empty, Helm creates a release-managed Secret and preserves its generated
credentials across pod restarts and chart upgrades.
kubectl get secret -n opscart-system opscart-watcher-auth \
-o jsonpath= ' {.data.username} ' | base64 --decode
echo
kubectl get secret -n opscart-system opscart-watcher-auth \
-o jsonpath= ' {.data.password} ' | base64 --decode
echo
To supply credentials yourself, create a Secret with username and password
keys and set auth.existingSecret :
kubectl create secret generic opscart-auth \
--from-literal=username=admin \
--from-literal=password= < your-password > \
-n opscart-system
auth :
existingSecret : " opscart-auth "
Standalone execution without configured credentials still generates and logs a
random password at startup; that password changes when the process restarts.
For team deployments, authenticate at the ingress layer instead with oauth2-proxy — supports Azure AD, Google, GitHub, and generic OIDC. See Security for the full pattern and threat model.
Want to evaluate OpsCart without breaking workloads manually? The repository includes a self-contained Kubernetes failure lab with intentionally unhealthy and healthy workloads distributed across eight namespaces.
Privileged-container detection
Missing NetworkPolicy coverage
Healthy workloads for comparison
Warning: The lab intentionally creates failing workloads, invalid image
pulls, a privileged container, and unused storage claims. Run it only in a
disposable or non-production Kubernetes cluster.
./examples/failure-lab/scripts/setup.sh
# Allow 30–60 seconds for failures and Kubernetes events to develop
opscart-scan triage \
--cluster " $( kubectl config current-context ) " \
--next-steps
# Remove every lab-owned namespace and workload
./examples/failure-lab/scripts/cleanup.sh
See the OpsCart Kubernetes Failure Lab for the complete scenario inventory, setup instructions, validation commands, and cleanup procedure.
The entry point isn't a KPI grid — it's a written assessment. Situation Briefing states what's actually wrong in plain sentences ("fraud-detection has been crash-looping for 7 days and its restart rate is accelerating"), not a wall of gauges. Below it:
Operational Memory scoreboard — total incidents ever seen, resolved, reopened, currently accelerating, longest-running incident, most unstable namespace — numbers only possible because OpsCart remembers, not just observes
Top 5 Things To Fix — ranked by severity and restart rate, each with a memory line (first detected, reopen count, trend) and a direct link to investigate
Cluster Health, Namespace Health, and Security Status at a glance
No competitor in this space — Grafana, Lens, k9s — can produce any of this, because none of them remember anything between scans.
Incident Score — A single 0–100 score derived from crash loops, image pull failures, security posture, waste, and network policy gaps. Trend arrows and a 7-point sparkline show whether the cluster is getting better or worse.
War Room — Every critical incident in one view, prioritized by severity and restart rate. Detects crash loops, image pull failures, probe failures (containers repeatedly killed by a failing startup/liveness check before reaching Ready), and privileged containers. Each card shows the issue type, namespace, age, restart count, and a ready-to-run kubectl command. One click opens a full investigation.
Investigation — One click from detection to investigation. Every pod-level incident includes:
OpsCart Assessment: what the pattern means and estimated investigation time, including whether the restart rate is accelerating
Incident Timeline: an operational journal — first detected, restart milestones, severity changes, resolved/reoccurred — persisted across pod restarts
Evidence: severity, first detected, restart count, state, age, owner
Blast Radius: replicas down, sibling pod health, services routing to the workload, ingress exposure, namespace-wide health, and a customer-impact heuristic (internal vs. possible external traffic)
Recommended Investigation: numbered steps with High / Medium / Low confidence and specific kubectl commands
Recent Events: last 10 events filtered to this pod
Related Resources: ConfigMaps, Secrets, PVCs referenced by the pod spec
Namespace-level findings (unprotected namespaces, idle namespaces) get their own dedicated view — a sample NetworkPolicy and namespace-scoped remediation steps, not a pod investigation that doesn't apply.
Incidents — The full system of record: every incident this cluster has seen, active and resolved. Search by name, namespace, or issue type; filter by severity and status; sorted and paginated. Resolved incidents stay visible with recovery time, so you can see what was broken last week — not just what's broken now.
Security Posture — CIS Kubernetes Benchmark scoring. Failed controls shown first, risk breakdown by category, prioritized remediation actions.
Waste & Drift — Zombie pods, orphaned PVCs with storage size and age, zero-replica workloads, abandoned namespaces.
Cost Intelligence — Provider-aware worker-node estimates with namespace allocation, embedded Azure pricing, and optional AWS public pricing. See Cost Intelligence .
Operational Memory — OpsCart remembers what happened. A lightweight local database tracks cluster snapshots, incident lifecycle (detected → milestones → resolved → reopened) as an append-only event journal, and scan metadata. Powers trend arrows, sparklines, incident age, and the per-incident timeline. Backed by SQLite, persisted on a PVC that survives pod restarts and helm uninstall . Configurable retention (90 days by default) keeps the database from growing unbounded.
Authentication — Basic auth on by default with no disable path: environment variables, a Kubernetes Secret, or an auto-generated password logged at startup. For teams, front it with oauth2-proxy for Azure AD / Google / GitHub / OIDC — see First Login above.
Helm Chart — Full Helm chart with configurable values, PVC-backed persistence, read-only RBAC, and non-root security context. See the chart README for persistence options, minikube notes, and all values.
Agentless — Runs as a single container. No sidecars, no DaemonSets, and no node access. Core scanning needs no cloud credentials; optional AWS API pricing uses workload identity.
Property
Detail
Base image
scratch —

[truncated]

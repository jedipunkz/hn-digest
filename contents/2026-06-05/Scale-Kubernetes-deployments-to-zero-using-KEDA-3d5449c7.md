---
source: "https://mijndertstuij.nl/posts/scale-to-zero-keda-cron-scaler/"
hn_url: "https://news.ycombinator.com/item?id=48417729"
title: "Scale Kubernetes deployments to zero using KEDA"
article_title: "Scale Kubernetes deployments to zero using KEDA - Mijndert Stuij"
author: "speckx"
captured_at: "2026-06-05T21:35:24Z"
capture_tool: "hn-digest"
hn_id: 48417729
score: 1
comments: 0
posted_at: "2026-06-05T20:25:57Z"
tags:
  - hacker-news
  - translated
---

# Scale Kubernetes deployments to zero using KEDA

- HN: [48417729](https://news.ycombinator.com/item?id=48417729)
- Source: [mijndertstuij.nl](https://mijndertstuij.nl/posts/scale-to-zero-keda-cron-scaler/)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T20:25:57Z

## Translation

タイトル: KEDA を使用して Kubernetes デプロイメントをゼロにスケールする
記事のタイトル: KEDA を使用して Kubernetes デプロイメントをゼロにスケールする - Mijndert Stuij
説明: KEDA の cron スケーラーを使用して開発、QA、およびステージング クラスターを一晩スリープ状態にする際のメモ: ArgoCD の落とし穴 (CRD、kube システムのロールバインディング、レプリカ ドリフト)、およびその代わりに何を諦めるか。

記事本文:
KEDA を使用して Kubernetes デプロイメントをゼロにスケールする
開発、QA、ステージング環境にマルチクラスター EKS デプロイメントがある場合、通常、それらを午前 3 時に実行する必要はありません。データを更新するためのいくつかの cron ジョブを除いて、誰もそのような機能を使用していません。それ以外はすべて無駄なアイドル コンピューティングです。
計算は簡単です。開発クラスターを 24 時間 365 日実行すると、1 週間に 168 時間の料金が発生します。これを平日のみのオフィス時間に削減すると、およそ 55 に減ります。これは、誰もアクティブに使用していないすべての環境におけるコンピューティング支出の約 3 分の 2 であり、クラスターごとに積み重なっていきます。私たちにとって、それは開発、QA、ステージングです。
純正の HPA はこれに関してはあまり役に立ちません。負荷に応じてスケールが変化し、硬い床はレプリカの 1 つです。ワークロードをゼロにすることはできません。
お金を無駄にしないように、デプロイメントをゼロにスケールダウンする方法は複数あります。最近、KEDA と cron スケーラーを使用して作業を行うようになりました。
KEDA は、Kubernetes Event-Driven Autoscaling の略称です。これは、Prometheus、VictoriaMetrics、その他多くの特殊なスケーラーと統合できるインテリジェントなオートスケーラーであり、特に Kafka ラグ、キューの深さ、Prometheus クエリ、および必要なものである cron に基づいてスケーリングできます。
KEDA をクラスターに一度インストールし、次に、scaledObject を使用してスケーリングの意図を定義します。
確認すべきいくつかのバージョンと前提条件:
ArgoCD 2.5 以降では、ServerSideApply=true 同期オプションが存在します。
ターゲットのデプロイメントは、Helm チャートにレプリカを固定しないでください。そうした場合、ArgoCD は同期のたびに設定を元に戻し、以下のignoreDifferences ルールでは問題を解決できません。フィールドを削除し、HPA がそれを所有できるようにします。
KEDA には公式 Helm チャートが同梱されており、ArgoCD を使用してインストールするだけでアプリケーションを 1 つ実行できます。
apiバージョン: argoproj.io/v1alpha1
種類 : アプリケーション
メタデータ:

名前：ケダ
名前空間 : argocd
仕様：
プロジェクト：プラットフォーム
目的地:
サーバー: https://kubernetes.default.svc
名前空間 : keda
ソース:
リポURL : https://kedacore.github.io/charts
チャート：ケダ
ターゲットリビジョン : 2.20.0
ヘルム：
値: |
crds:
インストール: true
同期ポリシー:
自動化:
自己修復 : true
プルーン : false
同期オプション:
- CreateNamespace=true
- ServerSideApply=true
この syncOptions ブロックは、少し面倒な作業を行っています。いくつか注意点があります:
注意点 1: CRD はサーバー側で適用する必要がある
KEDA は、いくつかの CRD ( ScaledObject 、 ScaledJob 、 TriggerAuthentication ) をインストールします。 Argo CD のデフォルトのクライアント側適用では、last-applied-configuration フィールドが注釈に追加され、分厚い CRD が Kubernetes の 256 KB アノテーション制限を超える可能性があり、metadata.annotations: Too long で同期が失敗します。 ServerSideApply=true は、API サーバーに独自のフィールド結合を許可することで、この問題を完全に回避します。
落とし穴 2: 制約のあるプロジェクトを打ち破る kube システムの roleBinding
KEDA のメトリクス アダプター ( keda-operator-metrics-apiserver ) は、集約された API サーバーです。これは、HPA が KEDA の計算されたメトリクスを読み取る方法である external.metrics.k8s.io を提供します。 Kubernetes アグリゲーションレイヤー契約に従って、すべての拡張 API サーバーは kube-system の extension-apiserver-authentication ConfigMap を読み取る必要があります。したがって、チャートはそこにリソースを 1 つだけ作成します。
種類 : ロールバインディング
メタデータ:
名前: keda - オペレーター - 認証 - リーダー
名前空間 : kube - システム
役割参照:
種類：役割
名前 : 拡張機能 - apiserver - 認証 - リーダー
AppProject が宛先名前空間を制限している場合 (そうすべきです)、同期は失敗します。
名前空間 kube - システムはプロジェクト ' ... ' では許可されていません
魅力的な解決策は、kube-system を共有アプリケーション プロジェクトに追加することです。ただし、すべてのアプリケーションは kube-system に書き込むことができるため、代わりに、

プラットフォーム固有のチャートに、より特権のある独自のプロジェクトを与えます。
apiバージョン: argoproj.io/v1alpha1
種類 : AppProject
メタデータ:
名前 : プラットフォーム
名前空間 : argocd
仕様：
説明 : 昇格された NS アクセスを備えたクラスター プラットフォーム アプリケーション。
ソースリポジトリ:
- https://kedacore.github.io/charts
目的地:
- サーバー: https://kubernetes.default.svc
名前空間 : keda
- サーバー: https://kubernetes.default.svc
名前空間 : kube - システム
クラスターリソースホワイトリスト:
- グループ: "*"
種類：「*」
落とし穴 3: ArgoCD が KEDA と戦うのを阻止する
KEDA がメトリクス (この場合は cron スケジュール) に基づいてデプロイメントを設定量にスケールすると、デフォルトで ArgoCD は既知の状態との調整を試みます。修正するには、レプリカ フィールドのアプリケーション (または ApplicationSet でテンプレート化された) にignoreDifferences ルールを追加します。
相違点を無視する :
- グループ: アプリ
種類：展開
jsonポインタ :
- /spec/レプリカ
cronスケーラー
cron トリガーはウィンドウを定義します。その中には、desiredReplicas が含まれています。すべてのウィンドウの外側では、KEDA はターゲットを minReplicaCount にスケールします。これを 0 に設定すると、アイドル時にワークロードがなくなります。
APIバージョン : keda.sh/v1alpha1
種類 : ScaledObject
メタデータ:
名前: 内部 - ダッシュボード
名前空間 : デフォルト
仕様：
スケールターゲット参照 :
名前 : 内部 - ダッシュボード # デプロイメント
minReplicaCount : 0 # ゼロへのスケールを有効にします
最大レプリカ数 : 1
CooldownPeriod : 300 # ウィンドウ終了後 5 分間待機します
トリガー:
- タイプ: cron
メタデータ:
タイムゾーン : ヨーロッパ/アムステルダム
start : "0 7 * * 1-5" # 月～金 07:00 → 起床
end : "0 19 * * 1-5" # 月～金 19:00 → 睡眠
望ましいレプリカ : "1"
生成された HPA には「最小レプリカ数: 1」と表示されますが、これは正しいです。通常の HPA は 1 を下回ることはできないため、KEDA はジョブを分割します。オペレーターは 0 ↔ 1 を処理し、HPA は 1 → max を処理します。ワークロードがスリープ状態になると、次のようになります。

:
最小レプリカ数 : 1
デプロイメント ポッド : 現在 0 / 希望 0
ScalingActive : False # "ターゲットのレプリカ数がゼロであるため、スケーリングは無効です"
ScalingActive: False および <unknown> メトリック値は警戒を要するように見えますが、ゼロであれば正常です。メトリクスを読み取るポッドがありません。動作する証拠は、電流 0 / 要求 0 です。
タイムゾーンは、クラスターのクロックではなく、指定したゾーンで評価されます。通常、ノードは UTC でログインするため、07:00 のヨーロッパ/アムステルダムのスケールアップはログに UTC で表示されます。タイムゾーンを明示的に設定すると、KEDA が変換を処理します。
ゼロへのスケールアップは無料ではありません。 start: 後の最初のリクエストは空のデプロイメントにヒットするため、レイヤー キャッシュが削除された場合のイメージのプル、ポッドの開始、readiness プローブなどのコールド スタート コストを支払います。典型的な Go または Node サービスの場合、通常は 1 分未満です。 JVM ワークロードや起動時に大規模なモデルを読み込むものは遅くなります。
エッジを取り除くいくつかの方法:
開始を設定: 最も早いユーザーが到着する数分前に。オフィスが 09:00 までに到着している場合、08:45 にスケールアップすると、誰かがログインしようとするまでにウォーム ポッドが取得されます。
2 番目の cron トリガーを実行して、ピーク負荷の 1 分前にプレウォームします。既知のバッチ ジョブが開始される直前。
遅れを受け入れてください。内部ダッシュボードやステージング環境では、その日の最初のヒットで 30 秒かかるコールド スタートを気にする人はいません。
内部ダッシュボードと管理ツール。
開発、QA、ステージング環境を夜間および週末に提供します。
「営業時間外は誰も触らない」ものが該当します。
SLA を備えた実際のユーザー トラフィックを提供するもの。コールド スタートだけでもエラー バジェットを消費してしまいます。
状態をフラッシュしたり、接続を開いたままにしたりするために存在する必要があるステートフル ワークロード。
スティッキーなセッションや、完全に終了できない長期間の実行中のリクエストを含むサービス。
遅いワークロード

数分間の最初のリクエストの遅延を許容しない限り、ウォームアップ (大きなクラスパスを持つ JVM、起動時にキャッシュをハイドレートするサービス、ML モデルの読み込み)。
スケジュールを知らずに他のチームが所有し、依存しているもの。土曜日の 07:01 のサプライズには節約の価値はありません。
ingress-nginx から Envoy Gateway への移行

## Original Extract

Notes from using KEDA's cron scaler to put dev, QA and staging clusters to sleep overnight: the ArgoCD gotchas (CRDs, kube-system RoleBinding, replica drift), and what you give up in exchange.

Scale Kubernetes deployments to zero using KEDA
If you have a multi-cluster EKS deployment for your development, QA and staging environments, those generally don't need to run at 3AM. No one is using any of that, except for maybe a few cronjobs to refresh data. Everything else is wasted idle compute.
The math is simple. A dev cluster running 24/7 bills for 168 hours a week. Trim it to office hours, weekdays only, and you're down to roughly 55. That's about two thirds of the compute spend on every environment nobody actively uses, and it stacks per cluster. For us that's dev, QA and staging.
The stock HPA can't really help with this; it scales on load, and its hard floor is one replica. It can't take your workloads to zero.
There are multiple ways to scale down your deployments to zero to make sure you don't waste that money. Recently I started using KEDA and a cron scaler to do the work for me.
KEDA is short for Kubernetes Event-Driven Autoscaling . It's an intelligent autoscaler that can integrate with Prometheus, VictoriaMetrics and many other more specialized scalers which allows it to scale based on, among many others, Kafka lag, queue depth, Prometheus queries, and the one we want: cron .
You install KEDA on your cluster once, and then define scaling intent using a scaledObject .
A few versions and assumptions to check:
ArgoCD 2.5 or newer, so the ServerSideApply=true sync option exists.
Your target Deployments should not pin replicas: in their Helm chart. If they do, ArgoCD will set it back on every sync and the ignoreDifferences rule below won't save you. Drop the field and let the HPA own it.
KEDA ships an official Helm chart and installing it using ArgoCD is just one Application away:
apiVersion : argoproj.io/v1alpha1
kind : Application
metadata :
name : keda
namespace : argocd
spec :
project : platform
destination :
server : https : //kubernetes.default.svc
namespace : keda
source :
repoURL : https : //kedacore.github.io/charts
chart : keda
targetRevision : 2.20.0
helm :
values : |
crds:
install: true
syncPolicy :
automated :
selfHeal : true
prune : false
syncOptions :
- CreateNamespace=true
- ServerSideApply=true
That syncOptions block is doing a bit of heavy lifting. A few gotchas to go through:
Gotcha 1: CRDs need server-side apply
KEDA installs a handful of CRDs ( ScaledObject , ScaledJob , TriggerAuthentication ). Argo CD's default client-side apply puts a last-applied-configuration field into an annotation, and chunky CRDs can blow past Kubernetes' 256 KB annotation limit causing your sync to fail with metadata.annotations: Too long . ServerSideApply=true sidesteps it entirely by letting the API server own field merging.
Gotcha 2: The kube-system RoleBinding that breaks restrictive projects
KEDA's metrics adapter ( keda-operator-metrics-apiserver ) is an aggregated API server; it serves external.metrics.k8s.io , which is how the HPA reads KEDA's computed metrics. Per the Kubernetes aggregation-layer contract, every extension API server must read the extension-apiserver-authentication ConfigMap in kube-system . So the chart creates exactly one resource there:
kind : RoleBinding
metadata :
name : keda - operator - auth - reader
namespace : kube - system
roleRef :
kind : Role
name : extension - apiserver - authentication - reader
If your AppProject restricts destination namespaces (and it should), the sync fails:
namespace kube - system is not permitted in project ' ... '
The tempting fix is to add kube-system to your shared application project. But then every application can write to kube-system so instead, give platform-specific charts their own, more privileged project:
apiVersion : argoproj.io/v1alpha1
kind : AppProject
metadata :
name : platform
namespace : argocd
spec :
description : Cluster platform applications with elevated NS access.
sourceRepos :
- https : //kedacore.github.io/charts
destinations :
- server : https : //kubernetes.default.svc
namespace : keda
- server : https : //kubernetes.default.svc
namespace : kube - system
clusterResourceWhitelist :
- group : "*"
kind : "*"
Gotcha 3: Stop ArgoCD from fighting KEDA
Once KEDA scales your deployments to a set amount based on metrics or in this case a cron schedule, by default ArgoCD will try to reconcile with the known state. The fix is to add an ignoreDifferences rule on the Application (or templated in your ApplicationSet ) for the replica field:
ignoreDifferences :
- group : apps
kind : Deployment
jsonPointers :
- /spec/replicas
The cron scaler
A cron trigger defines a window. Inside it you get desiredReplicas ; outside every window, KEDA scales the target to minReplicaCount . Set that to 0 and the workload disappears when idle:
apiVersion : keda.sh/v1alpha1
kind : ScaledObject
metadata :
name : internal - dashboard
namespace : default
spec :
scaleTargetRef :
name : internal - dashboard # the Deployment
minReplicaCount : 0 # enables scale-to-zero
maxReplicaCount : 1
cooldownPeriod : 300 # wait 5 min after a window ends
triggers :
- type : cron
metadata :
timezone : Europe/Amsterdam
start : "0 7 * * 1-5" # Mon–Fri 07:00 → wake up
end : "0 19 * * 1-5" # Mon–Fri 19:00 → sleep
desiredReplicas : "1"
The generated HPA says Min replicas: 1 which is correct. A normal HPA can't go below 1, so KEDA splits the job: the operator handles 0 ↔ 1 , the HPA handles 1 → max . When the workload is asleep you'll see:
Min replicas : 1
Deployment pods : 0 current / 0 desired
ScalingActive : False # "scaling is disabled since the replica count of the target is zero"
ScalingActive: False and a <unknown> metric value look alarming, but they're normal at zero. There are no pods to read metrics from. The proof it works is 0 current / 0 desired .
Timezone is evaluated in the zone you name, not the cluster's clock. Nodes typically log in UTC, so a 07:00 Europe/Amsterdam scale-up shows up in UTC in the logs. Set timezone explicitly and KEDA handles the conversion.
Scale-to-zero isn't free. The first request after start: hits an empty deployment, so you pay the cold-start cost: image pull if the layer cache has been evicted, pod start, readiness probe. For a typical Go or Node service that's usually under a minute. JVM workloads or anything loading a large model on boot will be slower.
A few ways to take the edge off:
Set start: a few minutes before your earliest user shows up. If the office is in by 09:00, scaling up at 08:45 gets you a warm pod by the time anyone tries to log in.
Run a second cron trigger to pre-warm a minute before peak load, e.g. just before a known batch job kicks off.
Accept the delay. For internal dashboards and staging environments, nobody really cares about a 30 second cold start on the first hit of the day.
Internal dashboards and admin tools.
Development, QA and staging environments overnight and on weekends.
Anything where "nobody touches this outside business hours" applies.
Anything serving real user traffic with an SLA. The cold start alone will eat your error budget.
Stateful workloads that need to be around to flush state or hold connections open.
Services with sticky sessions or long-lived in-flight requests you can't cleanly terminate.
Workloads with slow warmup (JVMs with big classpaths, services that hydrate caches on boot, ML model loading) unless you're fine with a multi-minute first-request delay.
Anything another team owns and depends on without knowing about the schedule. The surprise at 07:01 on a Saturday isn't worth the savings.
Migrating from ingress-nginx to Envoy Gateway

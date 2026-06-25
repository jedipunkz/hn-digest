---
source: "https://linkerd.io/2026/06/24/federating-clusters-for-zero-downtime-kubernetes/index.html"
hn_url: "https://news.ycombinator.com/item?id=48673197"
title: "Federating Clusters for Zero-Downtime Kubernetes"
article_title: "Federating Clusters for Zero-Downtime Kubernetes | Linkerd"
author: "PagCatOli"
captured_at: "2026-06-25T13:42:00Z"
capture_tool: "hn-digest"
hn_id: 48673197
score: 1
comments: 0
posted_at: "2026-06-25T13:37:52Z"
tags:
  - hacker-news
  - translated
---

# Federating Clusters for Zero-Downtime Kubernetes

- HN: [48673197](https://news.ycombinator.com/item?id=48673197)
- Source: [linkerd.io](https://linkerd.io/2026/06/24/federating-clusters-for-zero-downtime-kubernetes/index.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T13:37:52Z

## Translation

タイトル: ゼロダウンタイム Kubernetes のためのクラスターのフェデレーション
記事のタイトル: ゼロダウンタイム Kubernetes のためのクラスターのフェデレーション |リンカード
説明: Linkerd マルチクラスターは、フェデレーテッド、フラット、ゲートウェイの 3 つのモードをサポートします。 3 つの GKE クラスタ間で 3 つすべてを接続し、カオス テストを実行して、自動フェイルオーバーが動作していることを確認します。

記事本文:
Star Join フォーラムが Linkerd 2.20 を発表: レート制限を意識した負荷分散、メモリ使用量の削減、メトリクスの改善
ゼロダウンタイム Kubernetes のためのクラスターのフェデレーション
すべてのマルチリージョン設定は、最終的には同じ厄介な瞬間に遭遇します。
クラスターが消滅し、サービスの同一のコピーが 2 つのリージョンで実行される
それらを 1 つとして扱うように何も配線されていないため、存在しないのも同然かもしれません
事。フェイルオーバーがランブックになる: 復元、DNS の再ポイント、停止の待機
紙の上では、あなたは生きていくためにすでにお金を払っていることになります。
Linkerd のマルチクラスター拡張機能は、複数のクラスターを許可することでそのギャップを埋めます。
サービスを単一の負荷分散されたエンドポイントとして提供します。という部分は、
公式のタスクは実際のプラットフォームではほとんど選択されないということを覆い隠しています
マルチクラスターモード。一部のサービスはフェデレーションを必要とします (どこでも同じサービス、1 つのサービス
エンドポイント、自動フェイルオーバー)。他の人はミラーリングを望んでいます（特定のレベルに到達します）
名前によるリモート サービス)。そして、多くの場合、両方のパターンを
同じリンクのセット。ドキュメントでは、各モードを個別に説明します。このポストワイヤー
フルメッシュ リンク トポロジを使用して、3 つの GKE クラスタ全体で 3 つすべてを統合します。
クラスター全体を取り出すカオス テストと、クローンを作成して実行できるスクリプト
新しい GCP プロジェクトについて。
Companion リポジトリ: ここで参照されているすべてのスクリプトは次の場所にあります。
このリポジトリ。感じる
自由にクローンを作成し、プロジェクト ID を設定して実行できます。
Linkerd マルチクラスター モード: ゲートウェイ、フラット、フェデレーテッド
Linkerd のマルチクラスター拡張機能は 3 つのモードをサポートしています。素晴らしいのは、彼らは
相互に排他的ではありません: リンクされたクラスターの同じセット上で、モードが選択されます
ラベル経由でサービスごとに。
運用上重要な違いは、階層ミラーリングが機能することです。
あらゆるネットワーク上で。フラットかつ
フェデ

定格モードでは、実際のポッド間の接続が必要です。 GCP、VPC ネイティブ GKE 上
ピアリングされた VPC 上のクラスターにより、フラットなネットワークが無料で提供されます。それで、あなたは走ることができます
フラット ネットワークを介したコア ワークロードのフェデレーション サービスとミラーリング
その上にないクラスターからのゲートウェイを介した特殊なサービス
ネットワーク。私がこれまで見てきたほとんどのプラットフォーム チームは、最終的にはまさにこのような組み合わせになります。
マルチリージョン アーキテクチャ: GKE クラスタのセットアップ
3 つのリージョンに 3 つの GKE クラスタがあり、相互に完全にリンクされています（6 つの GKE クラスタ）。
方向リンクの合計)。それぞれ異なるものを使用する 3 つのデモ サービス
マルチクラスターモード:
フロントエンドはフェデレーションされており、3 つのクラスターすべてで実行されます。単一のフェデレーション
各クラスターのフロントエンド サービスは、9 つのポッドすべて (3 つのレプリカ) にわたって負荷分散します。
× 3 クラスター)。クラスターがダウンすると、残りの 6 つのポッドがクラスターを吸収します。
アプリケーションに変更を加えないトラフィック。
API はフラットミラーリングされており、west と east で実行されます。北のクラスター
それを、明示的なリモート サービスである api-west および api-east として消費します
トラフィックがリモート ポッドに直接送信される名前。これがあなたが目指すものです
クライアントがどのバックエンドと通信するかを決定する必要があるとき、たとえば、維持するために
データの局所性を求めるリージョン内のリクエスト。
Analytics はゲートウェイでミラーリングされ、 east でのみ実行されます。経由で輸出される
Linkerd ゲートウェイなので、西と北は、analytics-east-gw なしでアクセスします。
east のポッドへのフラットネットワーク接続が必要です。これは主に証明するためにここにあります
ゲートウェイ モードは、同じリンク上でフラット モードおよびフェデレーション モードと共存します。
デプロイの前提条件: GKE、Linkerd、CLI ツール
GCP アカウント (無料枠のクレジットでカバーされます。以下の 3 つの標準クラスタを使用します)
小規模なノード プール)
gcloud CLI、認証済み ( gcloud auth login )
ステップ CLI、brew install ステップ (証明書生成用)
完全なセットアップには約 30 分
赤外線

スクリプトはコンピューティング API とコンテナー API を有効にします。
新しいプロジェクトはすぐに機能します。
リポジトリのクローンを作成し、サンプル ファイルからローカル .env ファイルを作成し、カスタマイズします
GCP プロジェクト用です。デモの残りの部分ではデフォルトで十分なので、
ほとんどの場合、プロジェクト ID を変更するだけで済みます。
git clone <your-repo-url>
cd ブログ-リンカード-フェデレーション
cp env.example .env
.env を開き、少なくともプロジェクト ID を設定します。ファイルには賢明な機能が同梱されています
他のすべてのデフォルトは次のとおりです。
import GCP_PROJECT = "プロジェクト ID"
エクスポート REGION_WEST = "us-central1"
エクスポート REGION_EAST = "us-east1"
エクスポート REGION_NORTH = "ヨーロッパ-西 1"
# リージョンごとに 1 つのゾーン。ノードの場所を単一のゾーンに固定するので、num-nodes は次のようになります。
# 合計ノード数 — これが重要な理由については、以下のコストに関するメモを参照してください。
エクスポート ZONE_WEST = "us-central1-a"
エクスポート ZONE_EAST = "us-east1-b"
エクスポート ZONE_NORTH = "ヨーロッパ-west1-b"
import CLUSTER_MACHINE_TYPE = "e2-medium"
エクスポート CLUSTER_NODE_COUNT = "1"
エクスポート FRONTEND_REPLICAS = "3"
少なくとも GCP_PROJECT を設定します。それ以外はすべて適切なデフォルトで出荷されます。
3 つのリージョン、リージョンごとに 1 つのゾーン、およびコストを抑えるための小規模なノード プール。
cat .env を実行すると、変数の完全なセットが設定されていることがわかります。
スクリプトが変数を読み取れるように、変数を現在のシェルにロードします。
ソース.env
以下のすべてのスクリプトはこのファイルから読み取り、次のように実行されます。
set -euo Pipefail を使用すると、変数が欠落していると、静かに失敗するのではなく、大声で失敗します。
そのため、env.example には VPC 名とクラスター名という完全なセットが含まれています。
プロジェクト ID だけでなく、これも含まれます。
ステップ 1: VPC ピアリングを使用して 3 つの GKE クラスタをプロビジョニングする
インフラストラクチャ スクリプトを実行して、ネットワークとクラスタを作成します。これにはかかります
所要時間は約 10 ～ 15 分なので、コーヒーを飲むのに最適です。
./scripts/01-infra.sh
このスクリプトは次のことを行います

以下:
コンピューティング API とコンテナー API を有効にします (すでに有効になっている場合は何もしません)。
重複しないポッドとサービス CIDR を持つ 3 つの VPC を作成します。
VPC ピアリングの要件。
フルメッシュ VPC ピアリング (西↔東、東↔北、北↔西) を設定します。
--export-custom-routes および --import-custom-routes により、ポッド CIDR は
実際に宣伝されました。これがフラット ネットワークを実現します。
3 つの GKE Standard クラスタを作成し、VPC/リージョンごとに 1 つずつ、それぞれを固定します。
シングルゾーン。
kubectl コンテキストの名前を West 、 east 、 North に変更します。
スクリプトが使用するアドレス プランは次のとおりです。範囲は意図的に設定されています
重複しないため、VPC ピアリングはポッド トラフィックを正しくルーティングできます。
重複しない範囲については交渉の余地がありません。ポッド CIDR がピア間で重複する場合
VPC、ルーティングはサイレントに中断されます。ポッドが間違ったクラスターから応答を取得する、または
接続がタイムアウトになり、ログには有用なものが何も記録されません。どうやって知っているのか聞いてください。
ゾーンは 3 つではなく 1 つです。 GKE リージョン クラスタは --num-nodes ノードを配置します。
デフォルトでは 3 つのゾーンのそれぞれ。 --num-nodes 1 を使用すると、1 つあたり 3 ノードになります。
クラスター、合計 9、請求額は 3 倍になります。スクリプトは --node-locations を
単一ゾーンであるため、CLUSTER_NODE_COUNT=1 は実際にはクラスターごとに 1 つのノードを意味します。
コストメモ: 各実行に 1 つの e2-medium ノードを備えた 3 つの標準クラスター
このデモの合計は 1 日あたり約 10 ～ 15 ドル (管理料金 + ノード + 小規模なゲートウェイ)
東のロードバランサー)。分解スクリプトはすべてを削除します。
ステップ 2: 共有トラスト アンカーを使用して Linkerd をインストールする
共有トラスト アンカーを使用して、Linkerd を 3 つのクラスターすべてにインストールします。スクリプト
証明書を生成し、コントロール プレーンをインストールし、それぞれを設定します。
クラスタ間 mTLS のために他のクラスタを信頼するようにします。
./scripts/02-linkerd-install.sh
これにより、ルート CA とクラスターごとの発行者証明書が生成され、インストールされます。
3 つのクラスターすべての Linkerd:
r

oot.crt (共有トラストアンカー)
§── issuer-west.crt + issuer-west.key
§── issuer-east.crt + issuer-east.key
└── 発行者-north.crt + 発行者-north.key
クラスターごとの発行者証明書は、運用環境で維持する価値のある習慣です。
発行者が侵害されている場合は、他の発行者には触れずに単独で回転させます。
共有ルートは、クラスタ間の mTLS を機能させるものです。すべてのプロキシは、
他のすべてのプロキシの ID を同じアンカーに戻して検証します。
リソースの使用量 (およびコスト) を抑えるために、コントロール プレーンのみがインストールされます。
Viz アドオンなし。
ステップ 3: マルチクラスターをインストールし、フルメッシュ リンク トポロジを作成する
マルチクラスタ コンポーネントをセットアップし、クラスタ間でフルメッシュ トポロジを作成します。
クラスター。この手順の後、すべてのクラスターが他のクラスターからのサービスを利用できるようになります。
クラスター。
./scripts/03-multicluster-setup.sh
これは最も多くの作業が行われるステップです。 6 つの方向性リンクを作成します。
すべてのクラスターは他のすべてのクラスターにリンクされているため、すべてのクラスターは
フェデレーテッド ワークロード用の <svc> フェデレーテッド サービス。すべてのクラスターが消費できます。
他のサービスからミラーリングされたサービス。
シワが入り口です。東だけが 1 つ必要です (それが唯一のクラスターです)
分析を階層的にエクスポートするため）、イーストのゲートウェイを有効にします。
をインストールし、他の人はゲートウェイレスのままにしておきます。クラスターごとに 1 つのインストール、すべてのフラグ
一度にインストールすると、後でゲートウェイをボルトオンするためにインストールを 2 回再実行する必要はありません。
#west: ゲートウェイレス、クラスターごとに 1 つのコントローラーがあり、消費元
linkerd --contextwest multicluster install --gateway = false \
--set コントローラー [ 0 ] .link.ref.name = east \
--set コントローラー [ 1 ] .link.ref.name = 北 \
--set コントローラ [ 2 ] .link.ref.name = east-gw \
| kubectl --contextwest apply -f -
# east: ここでゲートウェイが有効になっており、それが使用するクラスターのコントローラー
リンカード --context eas

t マルチクラスターのインストール --gateway = true \
--set コントローラー [ 0 ] .link.ref.name =west \
--set コントローラー [ 1 ] .link.ref.name = 北 \
| kubectl --context east apply -f -
# 北: ゲートウェイレス、西、東、東のゲートウェイ リンクのコントローラー
linkerd --context North multicluster install --gateway = false \
--set コントローラー [ 0 ] .link.ref.name =west \
--set コントローラー [ 1 ] .link.ref.name = east \
--set コントローラ [ 2 ] .link.ref.name = east-gw \
| kubectl --context North apply -f -
コントローラーの数に注目してください。サービスミラーコントローラーは、消費側で実行されます。
側面、リンクごとに 1 つ。西と北はそれぞれゲートウェイ経由で分析を消費します。
したがって、east-gw リンク用の 3 番目のコントローラを取得します。東は消費しない
独自の分析機能があるため、必要なのは 2 つだけです。
次に、リンクを生成します。フラット/フェデレーション リンクは --gateway=false を使用します。の
東へのゲートウェイ対応リンク (分析エクスポート用) は別のリンクです
east-gw という名前:
# フラット リンク (ゲートウェイなし) — フェデレーション + フラット ミラーリング サービス用
linkerd --context east multicluster link-gen --cluster-name = east --gateway = false \
| kubectl --contextwest apply -f -
linkerd --contextwest multicluster link-gen --cluster-name =west --gateway = false \
| kubectl --context east apply -f -
# ... (全6方向)
# 東からのゲートウェイ対応リンク (分析階層エクスポート用)
linkerd --context east multicluster link-gen --cluster-name = east-gw \
| kubectl --contextwest apply -f -
linkerd --context east multicluster link-gen --cluster-name = east-gw \
| kubectl --context North apply -f -
この後、任意のクラスターの linkerd マルチクラスター チェックで、すべてのクラスターがレポートされるはずです。
リンクされたクラスターは正常です。
ステップ 4: デモ サービスをデプロイする
デモ ワークロードをデプロイします。次のセクションでは、フェデレーション、フラットのラベルを付けます。
ミラーリングとゲートウェイミラー

各モードが何を作成するかを示します。
./scripts/04-deploy-app.sh
3 つのサービス、3 つのモード、および意図的に同じ buoyantio/bb イメージ
それらはすべて、固定文字列をエコーする小さな HTTP サーバーです。アプリケーション
重要ではありません。重要なのは、1 つの kubectl ラベルによって Linkerd の動作が変わるということです。
他のすべてを一定に保ち、クラスター全体でサービスを扱います。
クラスターごとの応答文字列を使用して 3 つのクラスターすべてにデプロイし、ラベルを付けます。
フェデレーションの場合:
西東北のctxの場合。する
kubectl --context $ctx -n mc-demo label svc/frontend Mirror.linkerd.io/federated = member
完了しました
数秒以内に、frontend-federated が 3 つのクラスターすべてに表示されます。
$ kubectl --contextwest -n mc-demo get svc
名前 タイプ クラスター IP ポート ( S ) 年齢
フロントエンド ClusterIP 10.104.1.50 8080/TCP 45s
フロントエンドフェデレーテッド ClusterIP 10.104.2.100 8080/TCP 10s
API (フラットミラー)
フラット エクスポート用に API サービスにwest と east のラベルを付けます。
kubectl --contextwest -n mc-demo label svc/api Mirror.linkerd.io/exported = リモートディスカバリー
kubectl --context east -n mc-demo label svc/api Mirror.linkerd.io/exported = リモートディスカバリー
これで、Northern は api-west と api-east を別個のサービスとして認識できるようになりました。
$ kubectl --context North -n mc-demo get svc
名前 タイプ クラスター IP ポート ( S ) 年齢
フロントエンド ClusterIP 10.120.1.50 8080/TCP 4

[切り捨てられた]

## Original Extract

Linkerd multicluster supports 3 modes: federated, flat, and gateway. Wire all 3 across 3 GKE clusters, run a chaos test, and see automatic failover in action.

Star Join Forum Announcing Linkerd 2.20: Rate-limit-aware load balancing, reduced memory usage, and better metrics
Learn more Federating Clusters for Zero-Downtime Kubernetes
Every multi-region setup eventually meets the same awkward moment: a whole
cluster goes away, and the identical copy of your service running two regions
over might as well not exist, because nothing is wired to treat them as one
thing. Failover becomes a runbook: restore, repoint DNS, and wait for an outage
that, on paper, you’d already paid to survive.
Linkerd’s multicluster extension closes that gap by letting several clusters
present a service as a single, load-balanced endpoint. The part that the
official tasks gloss over is that a real platform almost never picks one
multicluster mode. Some services want federation (same service everywhere, one
endpoint, automatic failover). While others want mirroring (reach a specific
remote service by name). And you frequently want both patterns living on the
same set of links. The docs walk through each mode on its own. This post wires
all three together across three GKE clusters, with a full-mesh link topology, a
chaos test that takes out an entire cluster, and scripts you can clone and run
on a fresh GCP project.
Companion repo : Every script referenced here lives in
this repository . Feel
free to clone it, set your project ID, and run it.
Linkerd multicluster modes: Gateway, flat, and federated
Linkerd’s multicluster extension supports three modes. The nice thing is they’re
not mutually exclusive: on the same set of linked clusters, the mode is chosen
per service via a label.
The distinction that matters operationally is that hierarchical mirroring works
on any network. Only the gateway IP needs to be reachable, while flat and
federated modes need real pod-to-pod connectivity. On GCP, VPC-native GKE
clusters on peered VPCs give you that flat network for free. So, you can run
federated services for your core workloads over a flat network and still mirror
a specialized service through a gateway from a cluster that isn’t on that
network. Most platform teams I’ve seen end up with exactly this kind of mix.
Multi-region architecture: GKE cluster setup
We have three GKE clusters across three regions, fully linked to each other (six
directional links total). Three demo services, each using a different
multicluster mode:
frontend is federated and runs in all three clusters. A single federated
frontend service in each cluster load-balances across all nine pods (3 replicas
× 3 clusters). When a cluster goes down, the remaining six pods absorb the
traffic with no application changes.
api is flat-mirrored and runs in west and east . The north cluster
consumes it as api-west and api-east , which are explicit remote service
names with traffic sent straight to the remote pods. This is what you reach for
when the client needs to decide which backend it talks to, for example, to keep
a request in-region for data locality.
analytics is gateway-mirrored and runs only in east . Exported through the
Linkerd gateway so west and north reach it as analytics-east-gw without
needing flat-network connectivity to east ’s pods. It’s here mainly to prove
that gateway mode coexists with flat and federated modes on the same links.
Deployment prerequisites: GKE, Linkerd, and CLI tools
A GCP account (free-tier credits cover this. Use three standard clusters with
small node pools)
gcloud CLI, authenticated ( gcloud auth login )
step CLI, brew install step (for certificate generation)
~30 minutes for the full setup
The infra script enables the compute and container APIs for you, so a
brand-new project works out of the box.
Clone the repo, create a local .env file from the example file, and customize
it for your GCP project. The defaults are enough for the rest of the demo, so in
most cases you only need to change the project ID.
git clone <your-repo-url>
cd blog-linkerd-federation
cp env.example .env
Open .env and set at least your project ID. The file ships with sensible
defaults for everything else:
export GCP_PROJECT = "your-project-id"
export REGION_WEST = "us-central1"
export REGION_EAST = "us-east1"
export REGION_NORTH = "europe-west1"
# One zone per region. We pin node-locations to a single zone so num-nodes is
# the TOTAL node count — see the cost note below for why this matters.
export ZONE_WEST = "us-central1-a"
export ZONE_EAST = "us-east1-b"
export ZONE_NORTH = "europe-west1-b"
export CLUSTER_MACHINE_TYPE = "e2-medium"
export CLUSTER_NODE_COUNT = "1"
export FRONTEND_REPLICAS = "3"
At minimum, set GCP_PROJECT . Everything else ships with sensible defaults:
three regions, one zone per region, and small node pools to keep the cost down.
If you run cat .env , you should see the full set of variables populated.
Load the variables into your current shell so the scripts can read them:
source .env
Every script below reads from this file, and they all run with
set -euo pipefail , so a missing variable fails loudly rather than silently.
That’s why env.example carries the full set, the VPC and cluster names
included, instead of just the project ID.
Step 1: Provision three GKE clusters with VPC peering
Run the infrastructure script to create the networks and clusters. This takes
about 10–15 minutes, so it’s a good point to grab a coffee.
./scripts/01-infra.sh
This script does the following:
Enables the compute and container APIs (no-op if they’re already on).
Creates three VPCs with non-overlapping pod and service CIDRs, a hard
requirement for VPC peering.
Sets up full-mesh VPC peering (west↔east, east↔north, north↔west) with
--export-custom-routes and --import-custom-routes so pod CIDRs are
actually advertised. This is what gives us the flat network.
Creates three GKE Standard clusters , one per VPC/region, each pinned to a
single zone.
Renames the kubectl contexts to west , east , north .
Here’s the address plan the script uses. The ranges are intentionally
non-overlapping so VPC peering can route pod traffic correctly:
Non-overlapping ranges are non-negotiable. If pod CIDRs overlap across peered
VPCs, routing breaks silently. Pods get responses from the wrong cluster, or
connections time out with nothing useful in the logs. Ask me how I know.
One zone, not three. A GKE regional cluster places --num-nodes nodes in
each of three zones by default. With --num-nodes 1 that’s 3 nodes per
cluster, 9 total, and triple the bill. The script pins --node-locations to a
single zone so CLUSTER_NODE_COUNT=1 really means one node per cluster.
Cost note: Three Standard clusters with one e2-medium node each run
roughly $10–15/day total for this demo (management fee + nodes + a small gateway
load balancer on east ). The teardown script removes everything.
Step 2: Install Linkerd with a shared trust anchor
Install Linkerd into all three clusters using a shared trust anchor. The script
generates the certificates, installs the control plane, and configures each
cluster to trust the others for cross-cluster mTLS.
./scripts/02-linkerd-install.sh
This generates a root CA and per-cluster issuer certificates, then installs
Linkerd on all three clusters:
root.crt (shared trust anchor)
├── issuer-west.crt + issuer-west.key
├── issuer-east.crt + issuer-east.key
└── issuer-north.crt + issuer-north.key
Per-cluster issuer certs are a production habit worth keeping: if one cluster’s
issuer is compromised you rotate it in isolation, without touching the others.
The shared root is what lets cross-cluster mTLS work at all. Every proxy can
verify every other proxy’s identity back to the same anchor.
To keep resource usage (and cost) down, this installs the control plane only
with no Viz add-on.
Step 3: Install multicluster and create a full-mesh link topology
Set up the multicluster components and create a full-mesh topology between the
clusters. After this step, every cluster can consume services from every other
cluster.
./scripts/03-multicluster-setup.sh
This is the step with the most going on. We create six directional links,
every cluster linked to every other cluster, so every cluster gets a
<svc>-federated service for federated workloads, and every cluster can consume
mirrored services from any other.
The wrinkle is the gateway. Only east needs one (it’s the only cluster
exporting analytics hierarchically), so we enable the gateway in east’s
install and leave everyone else gatewayless. One install per cluster, all flags
at once, no re-running install a second time to bolt a gateway on afterward:
# west: gatewayless, with one controller per cluster it consumes from
linkerd --context west multicluster install --gateway = false \
--set controllers [ 0 ] .link.ref.name = east \
--set controllers [ 1 ] .link.ref.name = north \
--set controllers [ 2 ] .link.ref.name = east-gw \
| kubectl --context west apply -f -
# east: gateway enabled here, controllers for the clusters it consumes
linkerd --context east multicluster install --gateway = true \
--set controllers [ 0 ] .link.ref.name = west \
--set controllers [ 1 ] .link.ref.name = north \
| kubectl --context east apply -f -
# north: gatewayless, controllers for west, east, and east's gateway link
linkerd --context north multicluster install --gateway = false \
--set controllers [ 0 ] .link.ref.name = west \
--set controllers [ 1 ] .link.ref.name = east \
--set controllers [ 2 ] .link.ref.name = east-gw \
| kubectl --context north apply -f -
Note the controller count. The service-mirror controller runs on the consuming
side, one per link. west and north each consume analytics via the gateway,
so they get a third controller for the east-gw link; east doesn’t consume
its own analytics, so it only needs two.
Then we generate the links. Flat/federated links use --gateway=false ; the
gateway-aware link to east (for the analytics export) is a separate link
named east-gw :
# Flat links (no gateway) — for federated + flat-mirrored services
linkerd --context east multicluster link-gen --cluster-name = east --gateway = false \
| kubectl --context west apply -f -
linkerd --context west multicluster link-gen --cluster-name = west --gateway = false \
| kubectl --context east apply -f -
# ... (all six directions)
# Gateway-aware link from east (for the analytics hierarchical export)
linkerd --context east multicluster link-gen --cluster-name = east-gw \
| kubectl --context west apply -f -
linkerd --context east multicluster link-gen --cluster-name = east-gw \
| kubectl --context north apply -f -
After this, linkerd multicluster check on any cluster should report every
linked cluster healthy.
Step 4: Deploy the demo services
Deploy the demo workloads. The next sections label them for federation, flat
mirroring, and gateway mirroring and show what each mode creates.
./scripts/04-deploy-app.sh
Three services, three modes, and deliberately the same buoyantio/bb image for
all of them, a tiny HTTP server that echoes a fixed string. The application
isn’t the point. The point is that one kubectl label changes how Linkerd
treats the service across clusters, with everything else held constant.
Deploy to all three clusters with a per-cluster response string, then labeled
for federation:
for ctx in west east north; do
kubectl --context $ctx -n mc-demo label svc/frontend mirror.linkerd.io/federated = member
done
Within a few seconds, frontend-federated shows up in all three clusters:
$ kubectl --context west -n mc-demo get svc
NAME TYPE CLUSTER-IP PORT ( S ) AGE
frontend ClusterIP 10.104.1.50 8080/TCP 45s
frontend-federated ClusterIP 10.104.2.100 8080/TCP 10s
api (flat-mirrored)
Label the api service in west and east for flat export:
kubectl --context west -n mc-demo label svc/api mirror.linkerd.io/exported = remote-discovery
kubectl --context east -n mc-demo label svc/api mirror.linkerd.io/exported = remote-discovery
Now north can see api-west and api-east as separate services:
$ kubectl --context north -n mc-demo get svc
NAME TYPE CLUSTER-IP PORT ( S ) AGE
frontend ClusterIP 10.120.1.50 8080/TCP 4

[truncated]

---
source: "https://onatm.dev/2026/05/30/provisioning-a-private-talos-kubernetes-cluster-on-hetzner-cloud/"
hn_url: "https://news.ycombinator.com/item?id=48340800"
title: "Provisioning a Private Talos Kubernetes Cluster on Hetzner Cloud"
article_title: "Provisioning a Private Talos Kubernetes Cluster on Hetzner Cloud | Onat Mercan’s Blog"
author: "onatm"
captured_at: "2026-05-31T00:24:43Z"
capture_tool: "hn-digest"
hn_id: 48340800
score: 3
comments: 0
posted_at: "2026-05-30T21:31:02Z"
tags:
  - hacker-news
  - translated
---

# Provisioning a Private Talos Kubernetes Cluster on Hetzner Cloud

- HN: [48340800](https://news.ycombinator.com/item?id=48340800)
- Source: [onatm.dev](https://onatm.dev/2026/05/30/provisioning-a-private-talos-kubernetes-cluster-on-hetzner-cloud/)
- Score: 3
- Comments: 0
- Posted: 2026-05-30T21:31:02Z

## Translation

タイトル: Hetzner Cloud でのプライベート Talos Kubernetes クラスタのプロビジョニング
記事のタイトル: Hetzner Cloud でのプライベート Talos Kubernetes クラスタのプロビジョニング |オナット・メルカンのブログ
説明: これは、Tailscale を使用した Hetzner Cloud 上のプライベート ネットワーキングのフォローアップです。

記事本文:
オナット・メルカンのブログ
ホーム
について
Hetzner Cloud でのプライベート Talos Kubernetes クラスタのプロビジョニング
これは、Tailscale を使用した Hetzner Cloud 上のプライベート ネットワーキングのフォローアップです。
前回はネットワークについての記事でした。これは、私がそのネットワーク内に配置したもの、つまり Hetzner Cloud 上で Talos を実行するプライベート Kubernetes クラスタに関するものです。
重要なのは「Kubernetes on Hetzner Cloud」だけではありません。それに関する投稿がたくさんあります。私が気にしたのは、最初の起動時からクラスターをプライベートにすることでした。コントロール プレーンにはパブリック IP がありません。ワーカーにパブリック IP はありません。テールネット経由でのみアクセスします。
そのため、Talos が最適でした。パッケージマネージャーもSSHもありません。マシン構成を指定すると、Kubernetes ノードになり、ほとんどそれだけです。
クラスターに求めていたもの
プライベート専用ノード : すべての Kubernetes ノードは、Hetzner プライベート ネットワーク上にのみ存在する必要があります。
Terraform 管理のブートストラップ: マシン、Talos 設定、kubeconfig、およびベース アドオンはコードから取得する必要があります。
Talos : 手動によるサーバのメンテナンスは必要ありません。
個別のノード プール: プラットフォーム コンポーネントはアプリケーションのワークロードと競合するべきではありません。
GitOps : Terraform は ArgoCD をブートストラップでき、ArgoCD がプラットフォームを所有します。
目標は、すべての可動部分を理解できるほど小さく、かつ実際のプロジェクトを実行できるほど強力なものを構築することでした。
前の投稿のプライベート ネットワークにより、クラスターに /24 が存在するようになります。その範囲を明示的なチャンクに分割します。
プラットフォーム ワーカー: 10.0.128.32/27
一般労働者：10.0.128.64/27
サービスネットワーク: 10.0.192.0/21
コントロール プレーンには 3 つのノードがあります。プラットフォーム ワーカーは、ArgoCD やプラットフォーム コンポーネントなどを実行します。一般の従業員は、 snapbyte.dev などのアプリケーションを実行します。
フローチャートTB
テールネット((テールネット))
インターネット((インターネット))
サブグラフ VPC["プライベート ネットワーク 10.0.0.0/16"]
サブグラフ サブネット

[「サブネット 10.0.128.0/24」]
NAT[「NAT ゲートウェイ」]
サブグラフ CP["コントロール プレーン 10.0.128.16/28"]
CP1["CP-1"]
CP2["cp-2"]
CP3["CP-3"]
終わり
サブグラフ プラットフォーム["プラットフォーム ワーカー 10.0.128.32/27"]
アルゴCD["アルゴCD"]
PlatformApps[「プラットフォーム コンポーネント」]
終わり
サブグラフ 一般["一般労働者 10.0.128.64/27"]
PublicApps["パブリック アプリ"]
InternalApps["内部アプリ"]
終わり
終わり
終わり
テールネット -->|kubectl と talosctl| CP1
テールネット --> プラットフォーム
テールネット --> 一般
CP --> NAT
プラットフォーム --> NAT
一般 --> NAT
NAT --> インターネット
Kubernetes API エンドポイントは、最初のコントロール プレーン ノードのプライベート IP です。
地元の人々{
クラスターエンドポイント = "https://${local.control_plane_private_ips[0]}:6443"
}
このエンドポイントは、Tailscale を通じてすでにプライベート ネットワーク内にいる場合にのみ役立ちます。
Terraform がノードを作成できるようになる前に、Hetzner が起動できる Talos イメージが必要でした。
このクラスタは Talos v1.11.3 で起動しました。その後の v1.12.6 アップグレードは、初期設計によるものではなく、運用上のインシデントから発生しました。
Hetzner ではイメージ オプションとして Talos を提供していないため、Packer を使用して独自のスナップショットを構築します。このフローは、 hcloud-talos/terraform-hcloud-talos に基づいています。
一時的な Hetzner サーバを起動し、Talos Image Factory から Talos Raw イメージをダウンロードしてディスクに書き込み、結果をスナップショットとして保存します。
変数「talos_version」{
タイプ = 文字列
デフォルト = "v1.11.3"
}
ソース "hcloud" "talos" {
レスキュー = "linux64"
画像 = "debian-11"
場所 = "nbg1"
サーバータイプ = "cx22"
ssh_username = "root"
snapshot_name = "talos-${var.talos_version}-amd64"
スナップショットラベル = {
type = "インフラ"
os = 「タロス」
バージョン = var 。 talos_バージョン
アーチ = "amd64"
}
}
ラベル部分が重要です。 Terraform は後で、スナップショット名に依存する代わりにセレクターによってイメージを見つけることができます。
データ "hcloud_image" "talos" {
with_selector = "os=talos、type=infra、versi

on=${var.talos_version},arch=amd64"
}
ワーカープール
マシンを作成する前に、どのような種類のノードが必要かを説明する方法が必要でした。
これは基本的に、マネージド Kubernetes 製品のノード プールと同じ考え方です。 GKE、EKS、AKS ではすべて、さまざまなサイズ、ラベル、テイントを持つノードのグループを作成できます。私も同じメンタルモデルが欲しかった。
各プールには独自の Hetzner 配置グループもあります。これにより、Hetzner は、そのプール内のノードを可能な限り異なる物理ホストに分散するように指示されます。これによりプールの可用性が高くなるわけではありませんが、すべてのプラットフォーム ワーカーが同じマシン上で動作する障害モードは回避されます。
プールの構成は次のようになります。
ワーカープール = {
プラットフォーム = {
カウント = 3
sku = "cx33"
cidr = "10.0.128.32/27"
データセンター = "nbg1-dc3"
ラベル = { 目的 = "プラットフォーム" }
}
一般 = {
カウント = 3
sku = "cx23"
cidr = "10.0.128.64/27"
データセンター = "nbg1-dc3"
ラベル = { 目的 = "一般" }
}
}
これにより、Terraform コードの推論が容易になります。これにより、既知の CIDR 範囲、配置グループ、ラベルを持つマシンの名前付きグループを作成できます。
Terraform がマシンを作成する
ノード リソースは通常の Hetzner サーバーですが、パブリック ネットワークが無効になっています。
リソース "hcloud_server" "control_plane" {
カウント = var 。コントロールプレーン 。数える
名前 = "${local.cluster_name}-cp-${count.index + 1}"
データセンター = var 。コントロールプレーン 。データセンター
画像＝データ。 hcloud_image 。タロス 。 ID
サーバータイプ = var 。コントロールプレーン 。スク
パブリックネット {
ipv4_enabled = false
ipv6_enabled = false
}
ネットワーク{
network_id = var 。ネットワークID
ip = ローカル 。 control_plane_ips [カウント .インデックス]
}
}
前のセクションのワーカー プール マップは、個々のサーバーにフラット化されます。コードはエレガントではありませんが、結果は単純です。別のワーカーを一般プールに追加すると、次のワーカーが取得されます。

そのプール内のプライベート IP と適切なラベル。
地元の人々{
ワーカーフラット = マージ ([
pool_name の場合、 var の pool_config 。ワーカープール: {
範囲 (1 , pool_config . count + 1) 内の i の場合:
"${pool_name}-${i}" = > {
プール = プール名
インデックス = i
sku = pool_config 。スク
データセンター = プール_config 。データセンター
ラベル = pool_config 。ラベル
}
}
]...)
}
Talos ブートストラップ Kubernetes
サーバが存在すると、Talos Terraform プロバイダが引き継ぎます。マシン シークレットを生成し、コントロール プレーンとワーカー設定を作成し、パッチを適用し、最初のコントロール プレーン ノードをブートストラップし、Talos クラスタの健全性を待機して、kubeconfig を提供します。
コントロール プレーン ノードには 1 つの基本構成があり、プールごとに 1 つのワーカーの基本構成があります。
データ "talos_machine_configuration" "control_plane" {
クラスター名 = ローカル 。クラスター名
クラスターエンドポイント = ローカル。クラスターエンドポイント
machine_type = "コントロールプレーン"
machine_secrets = talos_machine_secrets 。これ 。マシンの秘密
talos_version = var 。 talos_バージョン
kubernetes_version = var 。 kubernetes_version
}
データ "talos_machine_configuration" "ワーカー" {
for_each = var 。ワーカープール
クラスター名 = ローカル 。クラスター名
クラスターエンドポイント = ローカル。クラスターエンドポイント
machine_type = "ワーカー"
machine_secrets = talos_machine_secrets 。これ 。マシンの秘密
talos_version = var 。 talos_バージョン
kubernetes_version = var 。 kubernetes_version
}
次に、Terraform は、パッチが適用されたコントロール プレーン構成を各コントロール プレーン ノードに適用します。
リソース "talos_machine_configuration_apply" "control_plane" {
カウント = var 。コントロールプレーン 。数える
client_configuration = talos_machine_secrets 。これ 。クライアント構成
machine_configuration_input = データ。 talos_machine_configuration 。コントロールプレーン 。マシン構成
ノード = ローカル 。 control_plane_private_ips [count .インデックス]
config_patches = [
ヤムレンコード

(ローカル .control_plane_patch )、
】
}
ワーカーは、パッチがワーカー プールから取得されることを除いて、同じパターンに従います。
リソース "talos_machine_configuration_apply" "worker" {
for_each = ローカル。ワーカーズフラット
client_configuration = talos_machine_secrets 。これ 。クライアント構成
machine_configuration_input = データ。 talos_machine_configuration 。労働者「それぞれ。価値 。プール』。マシン構成
ノード = flatten ( hcloud_server .worker [各 .key ].network )[0]。 ip
config_patches = [
yamlencode (ローカル .worker_pool_patches [各 .値 .プール])、
】
}
ノードには、適切なインストーラー イメージ、ノード IP の選択、デフォルト ルート、ポッドとサービス CIDR、および CNI の動作が必要です。
これは私が失敗し、後に最初の本当の失敗を引き起こした部分です。
最終的な意図を示すパッチの簡略版は次のようになります。
共通パッチ = {
マシン = {
インストール = {
ディスク = var 。インストールディスク
image = "factory.talos.dev/installer/${var.talos_schematic_id}:${var.talos_version}"
}
kubelet = {
extraArgs = {
「クラウドプロバイダー」=「外部」
"rotate-server-certificates" = true
}
ノードIP = {
validSubnets = ローカル 。ノードcidrs
}
}
ネットワーク = {
インターフェース = [
{
インターフェース = "eth0"
ルート = [{
ネットワーク = "0.0.0.0/0"
ゲートウェイ = var 。ゲートウェイ
}]
dhcp = true
}、
{
インターフェース = "eth1"
無視 = true
}
】
}
機能 = {
ホストDNS = {
有効 = true
forwardKubeDNSToHost = true
solveMemberNames = true
}
}
}
クラスター = {
ネットワーク = {
podSubnets = [ var .ポッド_ipv4_cidr ]
serviceSubnets = [ var .サービス_ipv4_cidr ]
cni = {
名前 = "なし"
}
}
プロキシ = {
無効 = true
}
}
}
そのほとんどは通常のクラスターのセットアップです。重要な部分は、nodeIP.validSubnets 、デフォルト ルート、およびインターフェイス名です。これらが間違っていても、クラスターは明らかな方法で障害を起こすことはありません。それは半分機能します、それはさらに悪いです。
コントロールプレーンパッチ

その他の重要なプライベート ネットワークの詳細を追加します。
クラスター = {
etcd = {
アドバタイズされたサブネット = [ var .コントロールプレーン 。シドル]
}
コントローラーマネージャー = {
extraArgs = {
「クラウドプロバイダー」=「外部」
「ノードcidrマスクサイズipv4」=「24」
"バインドアドレス" = "0.0.0.0"
}
}
}
これにより、etcd がコントロール プレーンのプライベート CIDR に保持され、Kubernetes ポッド CIDR 割り当てがクラスター ネットワークと一致します。
ワーカー パッチは主にプール ラベルを追加するため、ノードは purposes=platform または purposes=general になります。構成が適用された後、Terraform は最初のコントロール プレーン ノードのみをブートストラップします。
リソース "talos_machine_bootstrap" "this" {
client_configuration = talos_machine_secrets 。これ 。クライアント構成
ノード = ローカル 。コントロールプレーン_プライベート_ips [0]
}
最初の失敗はネットワークだった
最初の失敗は単純でした。ノードが自身のネットワーク ID に同意しませんでした。
プライベート専用 Hetzner マシンには、出力用のデフォルト ルートが必要です。私の設定では、ノード ルートはサブネット ゲートウェイに進み、Hetzner のネットワーク ルートはアウトバウンド トラフィックを NAT ゲートウェイに送信します。
ノード -> サブネット ゲートウェイ -> NAT ゲートウェイ -> インターネット
実際には、Talos マシン設定は適切なインターフェイスでそのルートを指す必要があり、kubelet は適切なノード IP を選択する必要があります。
ここで私はいくつかの間違ったスタートを切りました。ある時点で、私はプライベート ネットワーク インターフェイスが enp7s0 であると想定していました。次に、 eth1 を無視して、 enp7s0 と eth0 の両方でルートを試しました。 Hetzner Cloud VM は必要なネットワーク パスとして eth0 を使用していたので、実際に重要なパスは eth0 でした。
もう 1 つの微妙な部分は、nodeIP.validSubnets でした。私の最初の試みは、kubelet をより広範なサブネットに向けました。 Talos と Kubernetes は、許可されるノード IP 範囲がノードが存在する範囲と正確に一致する場合に、より優れたパフォーマンスを発揮します。
したがって、モジュールはコントロール プレーン CIDR と各ワーカー プール C からそのリストを構築します。

IDR:
地元の人々{
node_cidrs = concat (
[ 変数 .コントロールプレーン 。シドル]、
[ for _ 、 var の pool_config 。ワーカープール: pool_config 。シドル]
)
}
ネットワークの構成は簡単ではありません。これは、Hetzner ルート、Talos インターフェイス、kubelet ノード IP、etcd アドバタイズド サブネット、Kubernetes ポッド CIDR 割り当てなど、すべてが同意する必要がある一連の小さな設定です。
まずはCilium、次にクラウド統合
ノードがプライベート アドレスとルートに同意したら、次のステップはクラスター ネットワーク自体を機能させることです。
Talos が Kubernetes をブートストラップした後、Helm を使用して Cilium をインストールします。これはネイティブ ルーティング モードで実行され、kube-proxy を置き換えます。
リソース "helm_release" "cilium" {
名前 = 「繊毛」
リポジトリ = "https://helm.cilium.io/"
チャート = 「繊毛」
名前空間 = "kube-system"
セット{
名前 = "ルーティングモード"
値 = "ネイティブ"
}
セット{
名前 = "kubeProxyReplacement"
値 = "真"
}
}
Cilium を導入すると、Hetzner Cloud Controller Manager を実行できるようになります。 CCM は動作中のクラスター ネットワークに依存し、ロード バランサーやノード メタデータなどの Hetzner 固有の動作を担当するため、この順序は重要です。
CSI ドライバーとメトリック サーバーは同じ考え方に従います。 Terraform は、クラスターを使用可能にする基本部分をインストールします。
テラフォーミング

[切り捨てられた]

## Original Extract

This is a follow up to Private Networking on Hetzner Cloud with Tailscale

Onat Mercan's Blog
Home
About
Provisioning a Private Talos Kubernetes Cluster on Hetzner Cloud
This is a follow up to Private Networking on Hetzner Cloud with Tailscale
The previous post was about the network. This one is about what I put inside that network: a private Kubernetes cluster running Talos on Hetzner Cloud.
The important part is not just “Kubernetes on Hetzner Cloud”. There are many posts about it. The part I cared about was making the cluster private from the first boot. No public IPs on the control plane. No public IPs on the workers. Access only through the Tailnet.
That made Talos a good fit. No package manager, no SSH. You give it machine configuration, it becomes a Kubernetes node, and that is mostly it.
What I Wanted from the Cluster
Private-only nodes : every Kubernetes node should live only on the Hetzner private network.
Terraform-managed bootstrap : machines, Talos config, kubeconfig, and base add-ons should come from code.
Talos : no manual server maintenance.
Separate node pools : platform components should not fight application workloads.
GitOps : Terraform can bootstrap ArgoCD, then ArgoCD owns the platform.
The goal was to build something small enough that I could understand every moving part, but powerful enough that I could run actual projects on it.
The private network from the previous post gives the cluster a /24 to live in. I split that range into explicit chunks:
Platform workers: 10.0.128.32/27
General workers: 10.0.128.64/27
Service network: 10.0.192.0/21
The control plane has three nodes. Platform workers run things like ArgoCD and platform components. General workers run applications like snapbyte.dev .
flowchart TB
Tailnet((Tailnet))
Internet((Internet))
subgraph VPC["Private network 10.0.0.0/16"]
subgraph Subnet["Subnet 10.0.128.0/24"]
NAT["NAT Gateway"]
subgraph CP["Control Plane 10.0.128.16/28"]
CP1["cp-1"]
CP2["cp-2"]
CP3["cp-3"]
end
subgraph Platform["Platform Workers 10.0.128.32/27"]
ArgoCD["ArgoCD"]
PlatformApps["Platform components"]
end
subgraph General["General Workers 10.0.128.64/27"]
PublicApps["Public apps"]
InternalApps["Internal apps"]
end
end
end
Tailnet -->|kubectl and talosctl| CP1
Tailnet --> Platform
Tailnet --> General
CP --> NAT
Platform --> NAT
General --> NAT
NAT --> Internet
The Kubernetes API endpoint is the first control plane node’s private IP:
locals {
cluster_endpoint = "https://${local.control_plane_private_ips[0]}:6443"
}
That endpoint is only useful if you are already inside the private network through Tailscale.
Before Terraform could create any nodes, I needed a Talos image that Hetzner could boot.
I started this cluster on Talos v1.11.3 . The later v1.12.6 upgrade came from an operational incident, not the initial design.
Hetzner does not give you Talos as an image option, so I build my own snapshot with Packer. The flow is based on hcloud-talos/terraform-hcloud-talos .
It starts a temporary Hetzner server, downloads the Talos raw image from the Talos Image Factory, writes it to disk, and saves the result as a snapshot.
variable "talos_version" {
type = string
default = "v1.11.3"
}
source "hcloud" "talos" {
rescue = "linux64"
image = "debian-11"
location = "nbg1"
server_type = "cx22"
ssh_username = "root"
snapshot_name = "talos-${var.talos_version}-amd64"
snapshot_labels = {
type = "infra"
os = "talos"
version = var . talos_version
arch = "amd64"
}
}
The label part is the important bit. Terraform can later find the image by selector instead of relying on snapshot name:
data "hcloud_image" "talos" {
with_selector = "os=talos,type=infra,version=${var.talos_version},arch=amd64"
}
Worker Pools
Before creating the machines, I needed a way to describe what kind of nodes I wanted.
This is basically the same idea as node pools in managed Kubernetes offerings. GKE, EKS, and AKS all let you create groups of nodes with different sizes, labels, or taints. I wanted the same mental model.
Each pool also gets its own Hetzner placement group. That tells Hetzner to spread the nodes in that pool across different physical hosts where possible. It does not make the pool highly available, but it avoids the failure mode where every platform worker ends up on the same machine.
The pool config looks like this:
worker_pools = {
platform = {
count = 3
sku = "cx33"
cidr = "10.0.128.32/27"
datacenter = "nbg1-dc3"
labels = { purpose = "platform" }
}
general = {
count = 3
sku = "cx23"
cidr = "10.0.128.64/27"
datacenter = "nbg1-dc3"
labels = { purpose = "general" }
}
}
This makes the Terraform code easier to reason about. It lets me create named groups of machines with known CIDR ranges, placement groups, and labels.
Terraform Creates the Machines
The node resources are just regular Hetzner servers, but with the public network disabled.
resource "hcloud_server" "control_plane" {
count = var . control_plane . count
name = "${local.cluster_name}-cp-${count.index + 1}"
datacenter = var . control_plane . datacenter
image = data . hcloud_image . talos . id
server_type = var . control_plane . sku
public_net {
ipv4_enabled = false
ipv6_enabled = false
}
network {
network_id = var . network_id
ip = local . control_plane_ips [ count . index ]
}
}
The worker pool map from the previous section gets flattened into individual servers. The code is not elegant, but the outcome is simple: if I add another worker to the general pool, it gets the next private IP in that pool and the right labels.
locals {
workers_flat = merge ([
for pool_name , pool_config in var . worker_pools : {
for i in range ( 1 , pool_config . count + 1 ) :
"${pool_name}-${i}" = > {
pool = pool_name
index = i
sku = pool_config . sku
datacenter = pool_config . datacenter
labels = pool_config . labels
}
}
]...)
}
Talos Bootstraps Kubernetes
Once the servers exist, the Talos Terraform provider takes over. It generates machine secrets, creates control plane and worker configs, applies patches, bootstraps the first control plane node, waits for Talos cluster health, and gives me a kubeconfig.
There is one base config for control plane nodes, and one worker base config per pool:
data "talos_machine_configuration" "control_plane" {
cluster_name = local . cluster_name
cluster_endpoint = local . cluster_endpoint
machine_type = "controlplane"
machine_secrets = talos_machine_secrets . this . machine_secrets
talos_version = var . talos_version
kubernetes_version = var . kubernetes_version
}
data "talos_machine_configuration" "worker" {
for_each = var . worker_pools
cluster_name = local . cluster_name
cluster_endpoint = local . cluster_endpoint
machine_type = "worker"
machine_secrets = talos_machine_secrets . this . machine_secrets
talos_version = var . talos_version
kubernetes_version = var . kubernetes_version
}
Then Terraform applies the patched control-plane config to each control-plane node:
resource "talos_machine_configuration_apply" "control_plane" {
count = var . control_plane . count
client_configuration = talos_machine_secrets . this . client_configuration
machine_configuration_input = data . talos_machine_configuration . control_plane . machine_configuration
node = local . control_plane_private_ips [ count . index ]
config_patches = [
yamlencode ( local . control_plane_patch ),
]
}
Workers follow the same pattern, except the patch comes from the worker pool:
resource "talos_machine_configuration_apply" "worker" {
for_each = local . workers_flat
client_configuration = talos_machine_secrets . this . client_configuration
machine_configuration_input = data . talos_machine_configuration . worker [ each . value . pool ]. machine_configuration
node = flatten ( hcloud_server . worker [ each . key ]. network )[ 0 ]. ip
config_patches = [
yamlencode ( local . worker_pool_patches [ each . value . pool ]),
]
}
Nodes need the right installer image, node IP selection, default route, pod and service CIDRs, and CNI behavior.
This is the part I messed up and later caused the first real failure.
A simplified version of the patch, showing the final intent, looks like this:
common_patch = {
machine = {
install = {
disk = var . install_disk
image = "factory.talos.dev/installer/${var.talos_schematic_id}:${var.talos_version}"
}
kubelet = {
extraArgs = {
"cloud-provider" = "external"
"rotate-server-certificates" = true
}
nodeIP = {
validSubnets = local . node_cidrs
}
}
network = {
interfaces = [
{
interface = "eth0"
routes = [{
network = "0.0.0.0/0"
gateway = var . gateway
}]
dhcp = true
},
{
interface = "eth1"
ignore = true
}
]
}
features = {
hostDNS = {
enabled = true
forwardKubeDNSToHost = true
resolveMemberNames = true
}
}
}
cluster = {
network = {
podSubnets = [ var . pod_ipv4_cidr ]
serviceSubnets = [ var . service_ipv4_cidr ]
cni = {
name = "none"
}
}
proxy = {
disabled = true
}
}
}
Most of that is ordinary cluster setup. The important bits are: nodeIP.validSubnets , the default route, and the interface names. If those are wrong, the cluster does not fail in an obvious way. It half-works, which is worse.
The control plane patch adds the other important private-networking detail:
cluster = {
etcd = {
advertisedSubnets = [ var . control_plane . cidr ]
}
controllerManager = {
extraArgs = {
"cloud-provider" = "external"
"node-cidr-mask-size-ipv4" = "24"
"bind-address" = "0.0.0.0"
}
}
}
That keeps etcd on the control-plane private CIDR and makes Kubernetes pod CIDR allocation line up with the cluster network.
The worker patch mostly adds the pool labels, so nodes become purpose=platform or purpose=general . After the configs are applied, Terraform bootstraps only the first control plane node:
resource "talos_machine_bootstrap" "this" {
client_configuration = talos_machine_secrets . this . client_configuration
node = local . control_plane_private_ips [ 0 ]
}
The First Failure Was Networking
The first failure was simple: the nodes did not agree on their own network identity.
Private-only Hetzner machines still need a default route for egress. In my setup, the node route goes to the subnet gateway, and Hetzner’s network route sends outbound traffic to the NAT gateway.
Node -> Subnet Gateway -> NAT Gateway -> Internet
In practice, the Talos machine config needs to point that route at the right interface, and kubelet needs to pick the right node IP.
I had a few false starts here. At one point I assumed the private network interface was enp7s0 . Then I tried routes on both enp7s0 and eth0 , with eth1 ignored. Hetzner Cloud VMs were using eth0 for the network path I needed, so eth0 was the path that actually mattered.
The other subtle part was nodeIP.validSubnets . My first attempt pointed kubelet at the broader subnet. Talos and Kubernetes do better when the allowed node IP ranges are exactly the ranges where nodes live.
So the module builds that list from the control plane CIDR and each worker pool CIDR:
locals {
node_cidrs = concat (
[ var . control_plane . cidr ],
[ for _ , pool_config in var . worker_pools : pool_config . cidr ]
)
}
Configuring networking is not easy. It is a chain of small settings that all need to agree: Hetzner routes, Talos interfaces, kubelet node IPs, etcd advertised subnets, and Kubernetes pod CIDR allocation.
Cilium First, Then Cloud Integrations
Once the nodes agreed on their private addresses and routes, the next step was getting the cluster network itself working.
I install Cilium with Helm after Talos bootstraps Kubernetes. It runs in native routing mode and replaces kube-proxy.
resource "helm_release" "cilium" {
name = "cilium"
repository = "https://helm.cilium.io/"
chart = "cilium"
namespace = "kube-system"
set {
name = "routingMode"
value = "native"
}
set {
name = "kubeProxyReplacement"
value = "true"
}
}
After Cilium is in place, the Hetzner Cloud Controller Manager can run. That order matters because the CCM depends on a working cluster network and is responsible for Hetzner-specific behavior like load balancers and node metadata.
The CSI driver and metrics server follow the same idea. Terraform installs the base pieces that make the cluster usable.
Terraform in

[truncated]

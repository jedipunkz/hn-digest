---
source: "https://cozystack.io/blog/2026/08/blockstor-linstor-compatible-storage-for-kubernetes/"
hn_url: "https://news.ycombinator.com/item?id=49219348"
title: "Blockstor: A LINSTOR-compatible storage system for Kubernetes"
article_title: "Blockstor: a LINSTOR-compatible storage system for Kubernetes, written from scratch in Go | Cozystack"
author: "valyala"
captured_at: "2026-08-08T06:34:34Z"
capture_tool: "hn-digest"
hn_id: 49219348
score: 1
comments: 0
posted_at: "2026-08-08T06:26:38Z"
tags:
  - hacker-news
  - translated
---

# Blockstor: A LINSTOR-compatible storage system for Kubernetes

- HN: [49219348](https://news.ycombinator.com/item?id=49219348)
- Source: [cozystack.io](https://cozystack.io/blog/2026/08/blockstor-linstor-compatible-storage-for-kubernetes/)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T06:26:38Z

## Translation

タイトル: Blockstor: Kubernetes 用の LINSTOR 互換ストレージ システム
記事のタイトル: Blockstor: Go で最初から書かれた Kubernetes 用の LINSTOR 互換ストレージ システム |コージースタック
説明: Cozystack チームは、LVM および ZFS バックエンド、DRBD レプリケーション、LINSTOR 互換の REST API などの Blockstor をオープンソース化しているため、既存のクライアント ツールは変更せずに動作し続けます。

記事本文:
Blockstor: Go で最初から書かれた、Kubernetes 用の LINSTOR 互換ストレージ システム。 Cozystack cozystack-ロゴ Cozystack
エンタープライズサポート
OSS ヘルステレメトリ
開発統計
OpenSSF
OSS インサイト
2026 Cozystack 1.6: Talos ワーカー、テナント SSO、セキュリティ グループ、階層型クォータ、およびより安全な etcd アップグレード
Blockstor: Go で最初から書かれた Kubernetes 用の LINSTOR 互換ストレージ システム
サーバーが停止した場合でも、LINSTOR と DRBD がデータを維持する方法
Cozystack を単独で使用していますか?
ベアメタル上のテナント間のネットワーク分離: Cilium eBPF ポリシー、および Kube-OVN VPC
Talos Linux での CVE-2026-53359 (Januscape) および CVE-2026-46113 の修正
CVE-2026-43499 (GhostLock): Cozystack 暴露評価
セキュリティ アドバイザリ — CVE-2026-53359 (「Januscape」): KVM ゲストからホストへのエスケープ
Cozystack 1.5: ゲートウェイ API、デフォルト バックアップ、Flux シャーディング、マネージド サービス用の TLS、および GPU パススルー
KubeVirt VM のシンプル化: Cozystack の VM ディスクと VM インスタンス
Cozystack と OpenStack: 2026 年の正直な比較
etcd-operator が新しい v1alpha2 API で Cozystack に参加
5 分でゼロから kubectl まで — 独自のメタル上で Kubernetes を管理
Cozystack のプラットフォーム管理バックアップ
Cozystack 1.4: 新しいダッシュボード UI、永続的なテナント ワーカー、バックアップ戦略、およびフラクショナル GPU 共有
/cozystack:wizard の紹介 — ガイド付き Cozystack インストーラー
Cozystack 1.3: ストレージ対応スケジューリング、LINSTOR GUI、および VM デフォルト イメージ
同期レプリケーションを使用したマネージド PostgreSQL — 運用の問題なし
Cozystack が Web サイトに新しい OSS Health セクションを立ち上げました
CozySummit Virtual 2026: プログラムは設定されました — そしてそれは素晴らしいものになりそうです!
Cozystack 上のゲームサーバー: エイプリルフールの冗談はありません
Cozystack 1.2: OpenSearch、VPC ピアリング、およびよりスマートなテナント スケジューリング
Cozystack v1.0 および v1.1: 導入

パッケージベースのアーキテクチャ、Cozystack Operator、Velero Strategy Controller、MongoDB、OpenBAO のサポート
KubeCon Europe 2026 での Cozystack
CozySummit Virtual へのご招待 – 5 月 26 日
Cozystack v0.41: MongoDB、ダッシュボード編集ボタン、リソース クォータ UI、JWT セキュリティ、および cert-manager ゲートウェイ API
Cozystack v0.40: LINSTOR スケジューラ、SeaweedFS トラフィックの局所性、ValuesFrom 構成、およびプラットフォームの分解
2025 Cozystack v0.39: トポロジ認識ルーティング、Windows VM スケジューリング、Talm オーバーホール、テナント用 VMAgent
Talm v0.17: 秘密管理のための組み込み年齢暗号化
Flux-aio、Kubernetes mTLS、および鶏が先か卵が先かの問題
Cozystack v0.38: Virtual Private Cloud、VNC コンソール、構成可能な Worker K8s バージョン、および HTTPS 強制
Cozystack v0.37: OpenAPI ダッシュボード、Lineage Webhook、テナントでの PVC 拡張、および SeaweedFS S3 Discovery
Cozystack を CNCF に適用 Incubated
Cozystack を使用した Kubernetes の操作 Protofire エクスペリエンス
新しい CNCF ウェビナー: オープンソースを使用した独自のクラウド プラットフォームの構築
CNCF ウェビナー: 1 つの API ですべてを制御 — Kubernetes アグリゲーションを使用した統合プラットフォームの構築
CozySummit Virtual へのご招待 — 12 月 3 日
Cozyhr: Helm と Flux を使用してローカル開発を簡素化した方法
Cozystack が認定 Kubernetes プラットフォームになりました
仮想化プラットフォームの進化: マネージド サービスとローカル プロバイダーのエッジの台頭…
Cozystack が CNCF の CNAI Landscape で認められました!
任意のプロバイダを使用して任意のマシンに Talos Linux をインストールする簡単な方法
Cozystack が AI/ML 仮想マシンに GPU パススルーを提供
オープンソース プラットフォーム Cozystack 0.24 ～ 0.29 の更新:
Cozystack v0.30: GPU パススルー、PVC および IP 用の WorkloadMonitor、CPUManager、CI での自動テスト
Cozystack が CNCF サンドボックス プロジェクトになりました
Cozystack v0.23: Talos Linux v1.9.2、Te

legram アラートの重大度、VM インスタンスのフック、および Flux オペレーターの更新
Cozystack v0.22 リリース: テレメトリ、パッチ適用済み Talos v1.9.1、新しいエンティティ Workload および WorkloadMonitor
2024 年オープン ソース プラットフォーム Cozystack v0.21 の新年前リリースの紹介:
Cozystack の API アグリゲーション レイヤー用に動的な Kubernetes API サーバーを構築した方法
Cozystack v0.20 リリース: Terraform、Keycloak、安定性とセキュリティの改善
Cozystack v0.19: Keycloak SSO、ダッシュボード サービス ビュー、KubeVirt v1.4、および MetalLB アップデート
Cozystack on Hacktoberfest: 世界的な IT イベントに参加しましょう!
オープンソース プラットフォーム Cozystack バージョン 0.16.0
Cozystack オープンソース プラットフォームの最近の変更: Opencost、ログ収集システム、ブリッジ…
Cozystack が CNCF ランドスケープに正式に追加されました
Cozystack v0.15: OpenCost、Talos Metal Image、バックアップ修正、および Kamiji OOM 修正
Cozystack v0.14: 自動生成されたパスワード、RabbitMQ ユーザーと VHost、および CNPG v1.24
Cozystack v0.13: VictoriaLogs、VM ライブ マイグレーション、KubeVirt v1.3、およびブリッジ ネットワーキング
Cozystack v0.12: すべてのアプリの StorageClass、Cilium v1.16、VM 構成、および E2E サンドボックス
Cozystack によって管理される Kubernetes クラスターのインストール: Gohost と Ænix による詳細ガイド
Cozystack v0.10: FerretDB、NATS、テナント分離のためのネットワーク ポリシー、および etcd Operator v0.4
Cozystack v0.9: KubeVirt v1.2.2、Kamaji v1.0、テナント K8s v1.30、およびノード グループのアップグレード
Cozystack v0.8: FluxCD オペレーター、E2E テスト、ARM サポート、およびマネージド クラスター拡張機能
Talos Linux の構成マネージャーである Talm の紹介
Cozystack v0.7: ネットワークの安定化、DNS 修正、etcd 自動圧縮、および cozy.local ドメイン
Cozystack v0.6: VM シリアル コンソール、コンテナ用の一時ストレージ、および etcd 自動クォータ
Cozystack v0.5: 自動スキーマ生成、Cilium v1.14.10、および Mar

iaDB オペレーターの更新
Cozystack v0.4: etcd オペレーター、レプリカ オプション、Kamaji v0.5、およびダーク モード修正
Cozystack v0.3: Kafka、ClickHouse、Hetzner ベアメタルのサポート
DIY: Kubernetes を使用して独自のクラウドを作成する (パート 3)
DIY: Kubernetes を使用して独自のクラウドを作成する (パート 2)
DIY: Kubernetes を使用して独自のクラウドを作成する (パート 1)
Cozystack v0.2: バンドル、スキーマのバージョン管理、コア パッケージとしての FluxCD、およびコンポーネントのアップデート
Cozystack のご紹介: Kubernetes に基づく無料の PaaS プラットフォーム
Cozystack v0.1: ZFS サポート、リーダー選出、およびドキュメントの移動
2020 L2 モードでの MetalLB のルーティングの構成
Blockstor が異なるアプローチをとる理由
互換性と法的側面
次に何が起こるか、そしてどこで助けを求めますか
Blockstor: Go で最初から書かれた Kubernetes 用の LINSTOR 互換ストレージ システム
Cozystack チームは、Kubernetes のブロック ストレージのコントロール プレーンである Blockstor をオープンソース化しました。バックエンドとしての LVM と ZFS、DRBD を介したレプリケーション、LINSTOR 互換の REST API を備えています。このプロジェクトは cozystack 組織内にあり、CNCF サンドボックスに受け入れられたプラットフォームである Cozystack の一部として開発されています。ライセンスはApache2.0です。
一見の価値がある主な点は、これがフォークでもラッパーでもないことです。 Blockstor は Go で最初から書かれていますが、LINSTOR と同じ REST API を使用します。そのため、既に実行しているすべてのクライアント ツール (linstor CLI、linstor-csi、piraeus-operator、golinstor ライブラリ) は、変更を加えることなく引き続き機能します。
Blockstor が異なるアプローチを取る理由
LINSTOR は成熟したシステムであり、Cozystack の実稼働環境で何年も実行されてきました。私たちは機能の上限に達したのではなく、モデルに到達しました。
元のコントローラーはリクエストベースです。ほとんどの API 呼び出しでは、リアルタイムでノードに送信され、ノードの状態をポーリングして応答を組み立てます。あれは

2 つの結果。まず、この設計はスケーリングが不十分です。第 2 に、調整ループがないため、障害からの自動回復を外部から取り付ける必要があります。
Blockstor は、Kubernetes オペレーターが通常構築される方法で構築されます。つまり、望ましい状態が CRD 内に存在し、コントローラー ランタイム上の一連のリコンサイラーがクラスターをそれに向けて駆動します。 3 つの実際的な結果は次のとおりです。
外部データベースをバックアップしたり心配したりする必要はありません
コントローラーの再起動時にメモリ内の状態が失われることはありません
現実に遅れる可能性のあるコントローラー側でのノードのポーリングはありません
サテライトは API 自体を監視し、別のフィールド マネージャーを使用してサーバー側適用を通じて監視された状態を書き戻します。 Spec はコントローラーに属し、Status はサテライトに属し、その分割は厳密に適用されます。
3 つのコンポーネント。すべて通常の Kubernetes ワークロードです。
オブジェクトは、blockstor.cozystack.io/v1alpha1 グループに存在します: Node、StoragePool、ResourceGroup、ResourceDefinition、Resource、Snapshot、PhysicalDevice、および ControllerConfig。 CRD は、スキーマ レベルの検証とステータスの安全なマルチライター モデルを備えたパブリック統合ポイントとして設計されているため、GitOps ツールと監視が CRD と直接連携できます。
LVM、LVM-thin、ZFS、ZFS-thin、およびファイル バックエンド上にレプリケートされた DRBD ボリューム
DRBD フリー モード - ディスクフルまたはディスクレスの単一レプリカ
ボリュームレベルでの LUKS 暗号化。レイヤーは DRBD → LUKS → STORAGE としてスタックされます
制約付きの自動配置: ゾーン、ノード プロパティ、レプリカの拡散
TieBreaker およびクォーラム ポリシー — システムの中で最も厳しくテストされる部分の 1 つ
スナップショット: 新しいリソースへの作成、ロールバック、クローン作成、復元
zfs send/recv および Thin-send-recv を使用したクラスター内でのスナップショットの送信
オンラインでのボリュームのサイズ変更。縮小はデフォルトで無効になっており、次のことが必要です。

明示的なforce=true — ここでは元のコードよりも意図的に厳密にしています。
物理ディスクからのプールの作成
レプリカのリバランスと移行: 離脱ノードからの自動退避、ディスクフルへの自動昇格、スプリット ブレイン後のリカバリ
世代識別子をシードすることで、レプリカの追加時に初期同期をスキップします。 3 番目のレプリカをマルチテラバイトのボリュームに追加しても、再同期に数時間かかることはありません
ホット証明書リロード、Prometheus メトリクス、amd64 および arm64 用イメージを備えた API 上の mTLS
RWX — linstor-csi および NFS-Ganesha によるエンドツーエンド テストによって検証済み
3 日目に発見していただくよりも、このことをお知らせに載せたいと考えています。
以下は実装されておらず、サイレント 404 ではなく正直な 501 Not Implemented を返します: クラスター間スナップショットの配布、バックアップとバックアップ キュー、スケジュール、S3 などのリモート バックエンド、SPDK、NVMe-oF、OpenFlex、および Exos ドライバー。 Helm チャートはありません。インストールはプレーンなマニフェストを介して行われます。バージョンはまだ 0.x です。
オリジナルとの CLI 動作の相違点のリストは、既知の問題の登録および失敗した csi-sanity テストの記録とともに、公開で維持されます。平たく言えば、プロジェクト自体が独自のギャップのリストを公開します。
ストレージ コントロール プレーンを最初から書き直すには、約束ではなく証拠が必要です。私たちの答えはテストです。
実装は94,000行です。テストは Go の 170,000 行とシェルのさらに 46,000 行です。そして、これらは単体テストだけではありません。
108 の統合テストでは、すべての PR で envtest に対して実際の linstor Python クライアントを実行します。
コントラクト テストでは、ループバック デバイス上の Docker で実際の drbdmeta および drbdadm を実行します。
実際の DRBD を備えた Talos および QEMU リグで実行される 89 のエンドツーエンド シナリオ
91 個の CLI マトリックス セルと 74 個の repl

あらゆるシナリオでオペレーターのワークフローをカバー
パリティ ハーネスは、Blockstor の応答をライブアップストリーム LINSTOR と比較し、許容リストにない相違がある場合に CI を不合格とします。
リリース v0.1.11 については、別途言及する価値があります。これは、linstor-server バグ トラッカー自体から抽出された 48 件のエッジ ケースをライブ リグ上で再現してクローズし、上流で実行されているものと照合することで係争中のケースを解決しました。
互換性と法的側面
Blockstor は、linstor コントローラーのバージョンとして 1.33.2+git=blockstor を返し、linstor-csi および piraeus-operator が実際に呼び出すエンドポイントを実装します。 Piraeus は外部コントローラー モードで接続します。Blockstor apiserver のアドレスを指定すると、linstor-csi はそのまま動作し続けます。
LINSTOR は GPL で配布され、Blockstor は Apache 2.0 で配布されるため、オリジナルのソースは使用されていません。このプロジェクトはクリーンルーム実装です。互換性タイプは Apache 2.0 ライブラリである golinstor から取得されており、GPL ソースからコードがコピーまたは生成されることはありません。これは宣言ではなくチェック可能なルールです。PR ごとに、GPL、AGPL、LGPL、SSPL をランタイム グラフから除外するライセンス ゲートが CI で実行されます (GPL ライセンス仕様から生成されたコードも含まれます)。
そして直接的に言うと、LINSTOR、LINBIT、そして

[切り捨てられた]

## Original Extract

The Cozystack team has open-sourced Blockstor: LVM and ZFS backends, DRBD replication, and a LINSTOR-compatible REST API, so existing client tooling keeps working unchanged.

Blockstor: a LINSTOR-compatible storage system for Kubernetes, written from scratch in Go | Cozystack cozystack-logo Cozystack
Enterprise support
OSS Health Telemetry
DevStats
OpenSSF
OSS Insight
2026 Cozystack 1.6: Talos Workers, Tenant SSO, Security Groups, Hierarchical Quotas, and Safer etcd Upgrades
Blockstor: a LINSTOR-compatible storage system for Kubernetes, written from scratch in Go
How LINSTOR and DRBD Keep Your Data Alive When a Server Dies
Are You Using Cozystack Independently?
Network Isolation Between Tenants on Bare Metal: Cilium eBPF Policies, Plus Kube-OVN VPC
Fixing CVE-2026-53359 (Januscape) and CVE-2026-46113 on Talos Linux
CVE-2026-43499 (GhostLock): Cozystack Exposure Assessment
Security Advisory — CVE-2026-53359 ("Januscape"): KVM Guest-to-Host Escape
Cozystack 1.5: Gateway API, Default Backups, Flux Sharding, TLS for Managed Services, and GPU Passthrough
KubeVirt VMs Made Simple: VM Disk and VM Instance in Cozystack
Cozystack vs OpenStack: An Honest Comparison for 2026
etcd-operator Joins Cozystack with a New v1alpha2 API
From Zero to kubectl in 5 Minutes — Managed Kubernetes on Your Own Metal
Platform-Managed Backups in Cozystack
Cozystack 1.4: New Dashboard UI, Persistent Tenant Workers, Backup Strategies, and Fractional GPU Sharing
Introducing /cozystack:wizard — a Guided Cozystack Installer
Cozystack 1.3: Storage-Aware Scheduling, LINSTOR GUI, and VM Default Images
Managed PostgreSQL with Synchronous Replication — Without the Ops Headache
Cozystack Have Launched a New OSS Health Section on the Website
CozySummit Virtual 2026: The Program Is Set — and It Looks Amazing!
Game Servers on Cozystack: No April Fools' Joke
Cozystack 1.2: OpenSearch, VPC Peering, and Smarter Tenant Scheduling
Cozystack v1.0 & v1.1: Introducing Package-Based Architecture, Cozystack Operator, Velero Strategy Controller, MongoDB and OpenBAO Support
Cozystack at KubeCon Europe 2026
Invitation to CozySummit Virtual – May 26
Cozystack v0.41: MongoDB, Dashboard Edit Button, Resource Quota UI, JWT Security, and cert-manager Gateway API
Cozystack v0.40: LINSTOR Scheduler, SeaweedFS Traffic Locality, ValuesFrom Configuration, and Platform Decomposition
2025 Cozystack v0.39: Topology-Aware Routing, Windows VM Scheduling, Talm Overhaul, and VMAgent for Tenants
Talm v0.17: Built-in Age Encryption for Secrets Management
Flux-aio, Kubernetes mTLS and the Chicken and Egg Problem
Cozystack v0.38: Virtual Private Cloud, VNC Console, Configurable Worker K8s Versions, and HTTPS Enforcement
Cozystack v0.37: OpenAPI Dashboard, Lineage Webhook, PVC Expansion in Tenants, and SeaweedFS S3 Discovery
Cozystack applied to CNCF Incubated
Protofire Experience Operating Kubernetes with Cozystack
New CNCF Webinar: Building Your Own Cloud Platform with Open Source
CNCF Webinar: One API to Rule Them All — Building a Unified Platform with Kubernetes Aggregation
Invitation to CozySummit Virtual — December 3
Cozyhr: How We Simplified Local Development with Helm and Flux
Cozystack became a Certified Kubernetes Platform
The Evolution of Virtualization Platforms: The Rise of Managed Services and Local Providers’ Edge…
Cozystack Recognized in CNCF's CNAI Landscape!
A Simple Way to Install Talos Linux on Any Machine, with Any Provider
Cozystack Now Offers GPU Passthrough for AI/ML Virtual Machines
Updates to the Open-Source Platform Cozystack 0.24–0.29:
Cozystack v0.30: GPU Passthrough, WorkloadMonitor for PVCs and IPs, CPUManager, and Automated Testing in CI
Cozystack Becomes a CNCF Sandbox Project
Cozystack v0.23: Talos Linux v1.9.2, Telegram Alert Severity, VM Instance Hooks, and Flux Operator Update
Cozystack v0.22 Release: telemetry, patched Talos v1.9.1, new entities Workload and WorkloadMonitor
2024 Introducing the Pre-New Year Release of open source platform Cozystack v0.21:
How we built a dynamic Kubernetes API Server for the API Aggregation Layer in Cozystack
Cozystack v0.20 Release: Terraform, Keycloak, and Stability & Security Improvements
Cozystack v0.19: Keycloak SSO, Dashboard Services View, KubeVirt v1.4, and MetalLB Update
Cozystack on Hacktoberfest: become a part of the global IT event!
The Open Source Platform Cozystack Version 0.16.0
Recent Changes in the Cozystack Open Source Platform: Opencost, Log Collection System, Bridge…
Cozystack has officially been included in the CNCF Landscape
Cozystack v0.15: OpenCost, Talos Metal Image, Backup Fixes, and Kamaji OOM Fix
Cozystack v0.14: Auto-Generated Passwords, RabbitMQ Users and VHosts, and CNPG v1.24
Cozystack v0.13: VictoriaLogs, VM Live Migration, KubeVirt v1.3, and Bridge Networking
Cozystack v0.12: StorageClass for All Apps, Cilium v1.16, VM Configuration, and E2E Sandbox
Installing a Kubernetes Cluster Managed by Cozystack: A Detailed Guide by Gohost and Ænix
Cozystack v0.10: FerretDB, NATS, Network Policies for Tenant Isolation, and etcd Operator v0.4
Cozystack v0.9: KubeVirt v1.2.2, Kamaji v1.0, Tenant K8s v1.30, and Node Group Upgrades
Cozystack v0.8: FluxCD Operator, E2E Tests, ARM Support, and Managed Cluster Extensions
Introducing Talm, a configuration manager for Talos Linux
Cozystack v0.7: Network Stabilization, DNS Fixes, etcd Autocompaction, and cozy.local Domain
Cozystack v0.6: VM Serial Console, Ephemeral Storage for Containers, and etcd Auto-Quota
Cozystack v0.5: Automatic Schema Generation, Cilium v1.14.10, and MariaDB Operator Update
Cozystack v0.4: etcd Operator, Replica Options, Kamaji v0.5, and Dark Mode Fix
Cozystack v0.3: Kafka, ClickHouse, and Hetzner Bare-Metal Support
DIY: Create Your Own Cloud with Kubernetes (Part 3)
DIY: Create Your Own Cloud with Kubernetes (Part 2)
DIY: Create Your Own Cloud with Kubernetes (Part 1)
Cozystack v0.2: Bundles, Schema Versioning, FluxCD as Core Package, and Component Updates
Introducing Cozystack: A Free PaaS Platform based on Kubernetes
Cozystack v0.1: ZFS Support, Leader Election, and Documentation Move
2020 Configuring routing for MetalLB in L2 mode
Why Blockstor takes a different approach
Compatibility and the legal side
What is next, and where we would welcome help
Blockstor: a LINSTOR-compatible storage system for Kubernetes, written from scratch in Go
The Cozystack team has open-sourced Blockstor, a control plane for block storage in Kubernetes: LVM and ZFS as backends, replication over DRBD, and a LINSTOR-compatible REST API. The project lives in the cozystack organization and is developed as part of Cozystack, a platform accepted into the CNCF Sandbox. The license is Apache 2.0.
The main thing that makes it worth a look: it is not a fork, and it is not a wrapper. Blockstor is written from scratch in Go, but it speaks the same REST API as LINSTOR — so all the client tooling you already run keeps working without a single change: the linstor CLI, linstor-csi, piraeus-operator, and the golinstor library.
Why Blockstor takes a different approach
LINSTOR is a mature system, and it ran in production in Cozystack for years. We did not hit a functionality ceiling — we hit the model.
The original controller is request-based: for most API calls it goes out to the nodes in real time and polls their state to assemble a response. That has two consequences. First, this design scales poorly. Second, with no reconciliation loop, automatic recovery from failures has to be bolted on from the outside.
Blockstor is built the way Kubernetes operators are normally built: the desired state lives in CRDs, and a set of reconcilers on controller-runtime drives the cluster toward it. Three practical consequences follow:
No external database to back up and worry about
No in-memory state to lose when the controller restarts
No controller-side polling of nodes that can fall behind reality
The satellites watch the API themselves and write the observed state back through Server-Side Apply, using separate field managers. Spec belongs to the controller, Status to the satellite, and that split is enforced strictly.
Three components, all of them ordinary Kubernetes workloads:
The objects live in the blockstor.cozystack.io/v1alpha1 group: Node, StoragePool, ResourceGroup, ResourceDefinition, Resource, Snapshot, PhysicalDevice, and ControllerConfig. The CRDs are designed as a public integration point, with schema-level validation and a safe multi-writer model for Status, so that GitOps tooling and monitoring can work with them directly.
Replicated DRBD volumes on top of LVM, LVM-thin, ZFS, ZFS-thin, and file backends
A DRBD-free mode — a single replica, diskful or diskless
LUKS encryption at the volume level; the layers stack as DRBD → LUKS → STORAGE
Auto-placement with constraints: zones, node properties, and replica spreading
TieBreaker and quorum policies — one of the most heavily tested parts of the system
Snapshots: create, roll back, clone, and restore into a new resource
Snapshot shipping within the cluster using zfs send/recv and thin-send-recv
Online volume resize. Shrinking is disabled by default and requires an explicit force=true — here we are deliberately stricter than the original
Creating pools from physical disks
Replica rebalancing and migration: automatic evacuation from a departing node, automatic promotion to diskful, and recovery after split-brain
Skipping the initial sync when a replica is added, by seeding the Generation Identifier. Adding a third replica to a multi-terabyte volume does not turn into a multi-hour resync
mTLS on the API with hot certificate reload, Prometheus metrics, and images for amd64 and arm64
RWX — verified by an end-to-end test through linstor-csi and NFS-Ganesha
We would rather put this in the announcement than have you discover it on day three.
The following are not implemented, and they return an honest 501 Not Implemented rather than a silent 404: cross-cluster snapshot shipping, backups and the backup queue, schedules, remote backends such as S3, and the SPDK, NVMe-oF, OpenFlex, and Exos drivers. There is no Helm chart — installation goes through plain manifests. The version is still 0.x.
The list of CLI behaviour differences from the original is maintained in public, along with a register of known issues and a write-up of the csi-sanity tests that fail. Put plainly: the project itself publishes the list of its own gaps.
A storage control plane rewritten from scratch is a claim that needs proof, not promises. Our answer is tests.
The implementation is 94,000 lines. The tests are 170,000 lines of Go and another 46,000 lines of shell. And these are not only unit tests:
108 integration tests run the real linstor Python client against envtest on every PR
Contract tests run real drbdmeta and drbdadm in Docker on top of loopback devices
89 end-to-end scenarios run on a Talos and QEMU rig with real DRBD
91 CLI matrix cells and 74 replay scenarios cover operator workflows
A parity harness compares Blockstor’s responses against a live upstream LINSTOR and fails CI on any divergence that is not on the accepted list
Release v0.1.11 deserves a separate mention: it reproduced and closed 48 edge cases pulled from the linstor-server bug tracker itself — on a live rig, with the disputed cases settled by checking against a running upstream.
Compatibility and the legal side
Blockstor returns 1.33.2+git=blockstor for linstor controller version and implements the endpoints that linstor-csi and piraeus-operator actually call. Piraeus connects in external-controller mode — point it at the address of the Blockstor apiserver, and linstor-csi keeps working untouched.
LINSTOR is distributed under the GPL and Blockstor under Apache 2.0, so no sources from the original were used. The project is a clean-room implementation: the compatibility types come from golinstor, an Apache 2.0 library, and no code is copied or generated from GPL sources. This is not a declaration but a checkable rule: on every PR, a license gate runs in CI that keeps GPL, AGPL, LGPL, and SSPL out of the runtime graph — including code generated from a GPL-licensed specification.
And to be direct: LINSTOR, LINBIT, and

[truncated]

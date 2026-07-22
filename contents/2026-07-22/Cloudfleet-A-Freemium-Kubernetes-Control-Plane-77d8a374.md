---
source: "https://cloudfleet.ai"
hn_url: "https://news.ycombinator.com/item?id=49000746"
title: "Cloudfleet – A Freemium Kubernetes Control Plane"
article_title: "Cloudfleet - Managed multi-cloud Kubernetes platform"
author: "nikolay"
captured_at: "2026-07-22T01:44:25Z"
capture_tool: "hn-digest"
hn_id: 49000746
score: 2
comments: 0
posted_at: "2026-07-22T01:38:16Z"
tags:
  - hacker-news
  - translated
---

# Cloudfleet – A Freemium Kubernetes Control Plane

- HN: [49000746](https://news.ycombinator.com/item?id=49000746)
- Source: [cloudfleet.ai](https://cloudfleet.ai)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T01:38:16Z

## Translation

タイトル: Cloudfleet – フリーミアム Kubernetes コントロール プレーン
記事のタイトル: Cloudfleet - マネージド マルチクラウド Kubernetes プラットフォーム
説明: Cloudfleet は、クラウド、データセンター、エッジ全体にフルマネージド Kubernetes を提供します。自動アップグレード、ジャストインタイムのインフラストラクチャ、1 つのコントロール プレーン。

記事本文:
次世代のKubernetes
主権。
データセンターからクラウド、エッジに至るまで、Cloudfleet は本来の Kubernetes エクスペリエンスを提供します。フルマネージドのコントロール プレーン、ジャストインタイムのインフラストラクチャ、自動アップグレードはすべて、クラスタの単一画面を通じてシームレスに制御されます。
十分に活用されていないノードや高価なノードを自動的に置き換え、ワークロードをより効率的なインフラストラクチャに統合します。
1 つの Kubernetes クラスターだけで、運用オーバーヘッドなしで、どのリージョンのどのクラウド プロバイダーにもワークロードをデプロイできます。
パブリック クラウドとプライベート クラウド、オンプレミス インフラストラクチャ、エッジ環境を統合します。
Cloudfleet Kubernetes Engine (CFKE) は、複数のクラウド プロバイダーとオンプレミス環境にまたがるクラスター上でコンテナ化されたワークロードを実行するように設計されたフルマネージド Kubernetes サービスです。 CFKE は、単一の管理対象 Kubernetes コントロール プレーンを使用して、すべてのインフラストラクチャ フットプリントを一元化します。
Cloudfleet Container Registry (CFCR) は、コンテナ イメージを保存、共有、管理できるフルマネージドの OCI 準拠レジストリです。 CFCR は Cloudfleet の高可用性インフラストラクチャを活用しているため、スケーリングや運用上のオーバーヘッドについて心配する必要はありません。レジストリは Cloudfleet 組織に含まれており、追加のプロビジョニングは必要ありません。
ベスト プラクティスを組み込んだ運用対応アプリケーションをデプロイします。永続ストレージ用の CSI ドライバー、データベース オペレーター (PostgreSQL、MySQL、MongoDB)、監視および可観測性ツール、イングレス コントローラーなどの重要なインフラストラクチャ コンポーネントを含む、厳選された Helm チャートのコレクションを参照してください。各グラフは、Kubernetes クラスター上で一般的なワークロードを迅速にデプロイおよび管理できるように、セキュリティと信頼性のベスト プラクティスを使用して事前構成されています。
インフラストラクチャの足全体

印刷します。 1 つの統合クラスター。
マネージド コントロール プレーン Cloudfleet は CNCF に準拠しており、ベンダー ロックインを回避しながら、準拠したクラスターのシームレスな移行を可能にします。
Cloudfleet は、リアルタイムの需要に基づいてコンピューティング能力を動的にプロビジョニングします。
1 つのコマンドでオンプレミス ノードをクラスターに追加し、オンプレミス ノードとクラウド ノードの両方にわたってポッドをスケジュールします。
ノードは、複数のクラウドとリージョンにまたがる暗号化されたネットワークを介して安全に接続します。
組織とプロジェクトのスコープ、最小権限のアクセス許可、およびすべてのユーザー アクションに対する包括的な監査証跡を備えた、きめ細かいロールベースのアクセス制御 (RBAC)。
SAML および OIDC を介したエンタープライズ シングル サインオン (SSO) は、Okta、Microsoft Entra ID、Google Workspace、およびその他の互換性のある ID プロバイダーと統合されます。
SOC 2 および ISO 27001 標準 (認証中) に準拠したガバナンス、一元的な監査ログ、およびコンプライアンスの準備。
あなたにふさわしいエンタープライズサポート。
専用チーム、実証済みのプロセス、明確な運用上のコミットメントにより、大規模でミッションクリティカルな導入をサポートするように構築されています。
専門家主導のアーキテクチャ、導入、移行サービスは、複雑な環境全体で Cloudfleet を設計、展開、拡張するのに役立ちます。
専任のカスタマー サクセス チームが、ベスト プラクティスと運用ガイダンスにアクセスして、オンボーディング、導入、長期的な成功を導きます。
重大なインシデントに対する明確なエスカレーション パスを備えた、定義されたサポート チャネルを通じて経験豊富なエンジニアに 24 時間 365 日アクセスできます。
ミッションクリティカルなワークロードの可用性、応答時間、インシデント処理をカバーする、明確に定義されたサービス レベル アグリーメント (SLA)。
世界中のエンジニアリング チームの重要なインフラストラクチャに電力を供給します。
Findable Findable は Cloudfleet を使用してデータ制御を獲得し、AI インフラストラクチャのコストを 80% 削減します
インフラストラクチャのコストが 80% 削減される

ダクション
2 倍のドキュメント処理スループット
完全なデータ制御と主権
Hetzner と Exoscale TextCortex にわたるマルチクラウドの復元力 TextCortex は Cloudfleet を使用して AI ナレッジ ベース ソフトウェアを拡張します
オンボーディングの急増時に 100 以上のノードが自動スケーリング
インフラストラクチャのコストを 40% 以上削減
同じコストで 5 倍高速なナレッジ ベースのインデックス作成 BaseJump BaseJump は Cloudfleet でグローバル ゲーム インフラストラクチャを拡張します
複数のクラウドにわたる Kubernetes ワークロードのグローバル オーケストレーション
遅延の影響を受けやすいゲームサーバーの稼働率 99.9%
DevOps の複雑さを軽減したクラウドに依存しない柔軟性 Cloudfleet は、クラスターをどこにでもシームレスに拡張し、オフィスや自宅にあるあらゆるハードウェアを最新のエンタープライズ クラスターに変えます。
Cloudfleet を Hetzner と併用することで、米国を拠点とするハイパースケーラーと同じくらい迅速にマネージド Kubernetes クラスターを立ち上げることができましたが、EU でホストされているという利点もあり、これは今日の時代では非常に価値があります。
標準化された構成と Cloudfleet プラットフォームの強力な自動化機能を組み合わせることで、高速、安全、そして驚くほど効率的な開発ワークフローを構築しました。
柔軟性とマネージド サービスの間の完璧なバランス - サポートも優れています。
Kubernetes の構築または管理方法を変更することなく、お客様のインフラストラクチャにデプロイします。
単一のプロバイダーに縛られることなく、マルチクラウドのセットアップを簡単に構築できます。
ベンダー ロックインなしでマルチクラウド セットアップを構築するのは簡単です。
Cloudfleet を使用すると、ゲーム サーバーを複数のプロバイダーにわたって確実かつ手頃な価格で拡張できます。
Cloudfleet はインフラストラクチャのコストを削減するだけでなく、クラスター管理を外してくれるため、作業時間を節約できます。
最新のクラウド運用向けに設計されています。
パブリック クラウド、プライベート インフラストラクチャ、オンプレミス全体で Kubernetes を一貫して実行します

単一のコントロール プレーンと統合された運用モデルを備えた環境。
運用オーバーヘッドを削減する、一貫性のある独自のプラットフォームを使用して、環境全体で Kubernetes クラスターを運用、拡張、移行します。
透明性のある価格設定とインフラストラクチャ制御により、ハイパースケーラーのロックイン、隠れた料金、ワークロードの拡大に伴う予期せぬコストの増加を回避できます。
独自の拡張機能や強制的な依存関係を使用せずに標準の Kubernetes を実行することで、クラスター、ワークロード、ツールの完全な移植性を維持します。
実稼働ワークロード向けに設計されており、企業の信頼性要件を満たすように構築された運用ツール、エスカレーション パス、サポート プロセスが備えられています。
規制が厳しくセキュリティを重視する組織の要件を満たすように設計された分離、暗号化、アクセス制御により、デフォルトで安全です。
200 ユーロのクレジットで Cloudfleet をお試しください
マネージド Kubernetes クラスターを数分で作成します。新しいアカウントごとに 200 ユーロのクレジットが付与されるため、使用するか有効期限が切れるまで、無料でプラットフォーム全体を探索できます。
クラウドフリート
すべてのシステムが稼働中
Cloudfleet は、クラウド、データセンター、エッジ全体にフルマネージドの Kubernetes を提供します。自動アップグレード、ジャストインタイムのインフラストラクチャ、1 つのコントロール プレーン。
マルチクラウド Kubernetes クラスター
© 2025 Cloudfleet GmbH またはその関連会社。無断転載を禁じます。ベルリンで設計および設計されました。

## Original Extract

Cloudfleet delivers fully managed Kubernetes across cloud, datacenter, and edge. Automated upgrades, just-in-time infrastructure, one control plane.

Next Generation Kubernetes for
sovereignty.
From datacenter to cloud to edge, Cloudfleet delivers the Kubernetes experience as it was meant to be. Fully managed control plane, just-in-time infrastructure, automated upgrades, all seamlessly controlled through a single pane of glass for your clusters.
Automatically replace underutilized or expensive nodes and consolidate workloads onto more efficient infrastructure.
Just one Kubernetes cluster, with no operational overhead, to deploy workloads on any cloud provider in any region.
Integrate public and private clouds, on-premises infrastructure, and edge environments.
Cloudfleet Kubernetes Engine (CFKE) is a fully managed Kubernetes service designed to run containerized workloads on clusters spanning multiple cloud providers and on-premise environments. CFKE centralizes all your infrastructure footprint with a single managed Kubernetes control plane.
Cloudfleet Container Registry (CFCR) is a fully managed OCI-compliant registry that enables you to store, share, and manage container images. CFCR leverages Cloudfleet’s highly available infrastructure, so you do not have to worry about scaling or operational overhead. The registry is included with your Cloudfleet organization and requires no additional provisioning.
Deploy production-ready applications with best practices baked in. Browse our curated collection of Helm charts including essential infrastructure components like CSI drivers for persistent storage, database operators (PostgreSQL, MySQL, MongoDB), monitoring and observability tools, ingress controllers, and more. Each chart is pre-configured with security and reliability best practices to help you quickly deploy and manage common workloads on your Kubernetes clusters.
Your entire infrastructure footprint. One unified cluster.
Managed control plane Cloudfleet is CNCF-conformant, allowing seamless migration of any conformant clusters while avoiding vendor lock-in.
Cloudfleet dynamically provisions compute capacity based on real-time demand.
Add on-premises nodes to your cluster with a single command and schedule pods across both on-premises and cloud nodes.
Nodes securely connect via an encrypted network spanning multiple clouds and regions.
Fine-grained role-based access control (RBAC) with organization and project scopes, least-privilege permissions, and comprehensive audit trails for all user actions.
Enterprise Single Sign-On (SSO) via SAML and OIDC, integrating with Okta, Microsoft Entra ID, Google Workspace, and other compatible identity providers.
Governance, centralized audit logging, and compliance readiness aligned with SOC 2 and ISO 27001 standards (certifications in progress).
Enterprise support you deserve.
Built to support large-scale, mission-critical deployments with dedicated teams, proven processes, and clear operational commitments.
Expert-led architecture, deployment, and migration services to help you design, roll out, and scale Cloudfleet across complex environments.
A dedicated customer success team guiding onboarding, adoption, and long-term success, with access to best practices and operational guidance.
24/7 access to experienced engineers via defined support channels, with clear escalation paths for critical incidents.
Clearly defined Service Level Agreements (SLAs) covering availability, response times, and incident handling for mission-critical workloads.
Powering critical infrastructure for engineering teams worldwide.
Findable Findable gains data control and cuts AI infrastructure costs by 80% with Cloudfleet
80% infrastructure cost reduction
2x document processing throughput
Full data control and sovereignty
Multi-cloud resilience across Hetzner and Exoscale TextCortex TextCortex scales AI knowledge base software with Cloudfleet
100+ nodes autoscaling during onboarding spikes
>40% infrastructure cost savings
5x faster knowledge base indexing at the same cost BaseJump BaseJump scales global gaming infrastructure with Cloudfleet
Global orchestration of Kubernetes workloads across multiple clouds
99.9% uptime for latency-sensitive game servers
Cloud-agnostic flexibility with reduced DevOps complexity Cloudfleet seamlessly extends your cluster anywhere, turning any hardware - even in your office or home - into a modern enterprise cluster.
Using Cloudfleet together with Hetzner allowed us to bring up a managed Kubernetes cluster just as quick as with any US-based hyper-scaler, but with the benefit of being EU hosted which is very valuable in today's times.
By combining our standardized configurations with the powerful automation features of the Cloudfleet platform, we've built a development workflow that is fast, secure, and incredibly efficient.
Perfect balance between flexibility and managed service - and the support is outstanding.
We deploy to customer infrastructure without changing how we build or manage Kubernetes.
It is easy to build a multi-cloud setup without getting locked into any single provider.
It is easy to build a multi-cloud setup without vendor lock-in.
Cloudfleet lets us scale our game servers across multiple providers - reliably and affordably.
Cloudfleet not only cut our infrastructure costs - it saved us hours of work by taking cluster management off our plate.
Designed for modern cloud operations.
Run Kubernetes consistently across public cloud, private infrastructure, and on-prem environments, with a single control plane and unified operational model.
Operate, scale, and migrate Kubernetes clusters across environments using a consistent, opinionated platform that reduces operational overhead.
Transparent pricing and infrastructure control help you avoid hyperscaler lock-in, hidden fees, and unexpected cost growth as your workloads scale.
Retain full portability of your clusters, workloads, and tooling by running standard Kubernetes without proprietary extensions or forced dependencies.
Designed for production workloads, with operational tooling, escalation paths, and support processes built to meet enterprise reliability requirements.
Secure by default, with isolation, encryption, and access controls designed to meet the requirements of regulated and security-conscious organizations.
Try Cloudfleet with €200 in credit
Create a managed Kubernetes cluster in minutes. Every new account gets €200 in credit, so you can explore the full platform with no charge until it's used or expires.
Cloudfleet
All Systems Operational
Cloudfleet delivers fully managed Kubernetes across cloud, datacenter, and edge. Automated upgrades, just-in-time infrastructure, one control plane.
Multi-cloud Kubernetes clusters
© 2025 Cloudfleet GmbH or its affiliates. All rights reserved. Designed and engineered in Berlin.

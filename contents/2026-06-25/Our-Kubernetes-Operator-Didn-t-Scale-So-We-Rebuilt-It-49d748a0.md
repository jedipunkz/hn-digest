---
source: "https://infisical.com/blog/kubernetes-operator-rebuild"
hn_url: "https://news.ycombinator.com/item?id=48674416"
title: "Our Kubernetes Operator Didn't Scale, So We Rebuilt It"
article_title: "Our Kubernetes Operator Didn’t Scale, So We Rebuilt It"
author: "FinnLobsien"
captured_at: "2026-06-25T14:57:28Z"
capture_tool: "hn-digest"
hn_id: 48674416
score: 1
comments: 0
posted_at: "2026-06-25T14:56:38Z"
tags:
  - hacker-news
  - translated
---

# Our Kubernetes Operator Didn't Scale, So We Rebuilt It

- HN: [48674416](https://news.ycombinator.com/item?id=48674416)
- Source: [infisical.com](https://infisical.com/blog/kubernetes-operator-rebuild)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T14:56:38Z

## Translation

タイトル: Kubernetes オペレーターが拡張できなかったため、再構築しました
記事のタイトル: Kubernetes Operator が拡張できなかったため、再構築しました
説明: Infisical が、認証呼び出しを減らし、スケーリングを改善し、構成を簡素化するリファレンスベースのアーキテクチャを使用して Kubernetes Operator を再構築した理由を説明します。

記事本文:
プラットフォーム シークレット管理 PKI PAM エージェント Vault 価格設定ドキュメント リソース ブログ ビデオ ケース スタディ Infisical と Hashicorp Vault の比較 キャリア 27,000 デモをリクエスト 無料で始めましょう プラットフォーム価格設定ドキュメント リソース キャリア 27,000 つ星 デモをリクエスト 無料で始めましょう ← 戻る ブログ投稿 • 7 分で読む Kubernetes オペレーターはスケールしなかったため、再構築しました
セキュリティは利便性と相反することがよくありますが、人間の脳は利便性を好みます (善意であっても間違いを犯します)。ほとんどのアイデンティティ セキュリティ ツールは、セキュリティを可能な限り便利にすることでこれを調整します。
何らかの摩擦があると、人々が便利な回避策を構築したり、セキュリティ ツールを回避したりして、組織のセキュリティ体制が弱くなるリスクが高まります。
これが、大規模な場合に問題が発生することに気づき、パフォーマンスと開発者のエクスペリエンスを向上させるために Kubernetes オペレーターを再設計した理由です。
Kubernetes オペレーターを構築した理由
優れた秘密管理の原則の 1 つは集中化です。すべてのシークレットを統合すると、インフラストラクチャ全体でシークレットを 1 か所に保存、管理、監査できるようになります。
一元化するには、あらゆる種類のインフラストラクチャおよび展開モデルにシークレットを同期する必要があります。これは、シークレット ストアからユーザーのインフラストラクチャに、シークレットを使用するサービスにシークレットを配信することを意味します。理想的には、セキュリティ上のギャップが生じたり、ユーザーにメンテナンスの負担がかかることを避けるための回避策やカスタム ロジックを使用せずにこれが発生します。
最初の Kubernetes オペレーターは、ネイティブ シークレット同期を分散デプロイメントに拡張しました。うまくいきましたが、うまく拡張できませんでした。クラスターまたはデプロイメント内のリソースが増加すると、そのメモリー使用量が増大し、パフォーマンスが低下しました。
これが、新しいリファレンスベースのアーキテクチャを使用してオペレーターを再設計した理由です。
私たちの最初の Kubernetes オペラの場所

たじろいだ
シークレットを Kubernetes に大規模に同期するには、オペレーターがいくつかの作業を行う必要がありました。
Infisical への接続と認証: API がホストされている場所、それに接続して認証する方法。
Infisical で正しいシークレットを見つける: Infisical 内の正しいシークレット パス (プロジェクト、環境、フォルダーなど) を把握します。
ポッドとデプロイメントがそれらを使用できるようにする: シークレットを Kubernetes ネイティブのシークレット オブジェクトに調整して、デプロイメントとポッドがシークレットを正しい環境に取得できるようにします。
私たちの初期設計はこれらの要件を満たしていましたが、ワークロードが増加するにつれて行き詰まりました。
v1 のアーキテクチャが大規模化に苦戦した理由
v1 では、ユーザーは Infisical Secret を指すモノリシックなカスタム InfisicalSecret リソースを作成しました。 v1alpha1 API の各 InfisicalSecret リソースには以下が含まれていました。
非物理インスタンスのアドレス
認証資格情報
書き込み先のマネージド Kubernetes シークレット
APIバージョン: Secrets.infisical.com/v1alpha1
種類: InfisicalSecret
メタデータ:
名前: サービス・ア・シークレット
仕様:
ホストAPI: https://app.infisical.com/api
認証:
ユニバーサル認証:
認証情報参照:
SecretName: ユニバーサル認証資格情報
Secret名前空間: デフォルト
秘密の範囲:
projectSlug: 私のプロジェクト
envSlug: prod
シークレットパス: "/service-a"
管理対象シークレット参照:
SecretName: 管理対象サービス
Secret名前空間: デフォルト
各リソースが認証と接続を複製するため、スケーラビリティが低下しました。このアーキテクチャは永続的なインフラストラクチャ上で動作します。 VPS または VM は、再起動または構成変更時にのみ再認証される ID の 1 つです。ただし、Kubernetes クラスターにはこれらのリソースが数十、数百含まれており、頻繁に再デプロイと再起動が行われます。それぞれが独自の認証と接続を実行するため、それぞれが独自の独立したクライアントを保持していました。これにより、次の 3 つの問題が発生しました。
リソースはそれぞれの理由により、過大なメモリを消費します。

リソースはメモリ内に独自のクライアントを保持していました。大規模になると、Helm メモリ制限の引き上げが必要となるメモリ不足の問題が発生しました。ポッドが OOM になるたびに、各リソースが同時に調整され、認証されます。
再起動により大量の同時認証呼び出しが発生し、レート制限に達しました。オペレーターはバックオフと再試行後に成功しますが、クラスターを定常状態にするまでに待ち時間が発生します。
エンジニアリング チームは追加の作業を行う必要がありました。マシン ID のローテーションまたは Infisical ホストの変更は、すべての単一リソースの認証ブロックを編集することを意味します。
根本的な問題はオーバーロードされた CRD アーキテクチャであり、ロジックが欠落していることではありませんでした。イベント ハンドラー、ジッター、その他のロジックを評価しました。これらは役に立ったかもしれませんが、すでに過負荷になっている CRD がさらに複雑になってしまいました。
新しいアーキテクチャでのみ根本的な問題を解決できました。
リファレンスベースのアーキテクチャがレプリケーションをどのように修正したか
新しい設計では、接続、認証、同期が分離されています。シークレットは、共有オブジェクトとして認証および接続リソースを参照します。これにより、パフォーマンスの問題が修正され、開発者のエクスペリエンスが向上します。
新しいアーキテクチャは、プロバイダー、ストア、および外部シークレット CRD を分離する、外部シークレット オペレーター (ESO) のリソース分割に基づいて大まかにモデル化されました。 Infisical は ESO と統合されていますが、次の 2 つの理由から独自のオペレーターを構築しています。
ESOは以前にも開発を一時停止したことがある。その後再開されましたが、開発が停止する可能性のある依存関係を作成するのは理想的ではありません。
私たちはネイティブ UX を提供し、可動部分を最小限に抑え、最終的には ESO がサポートしていない動的シークレットとプッシュ シークレット (Kubernetes で発生し、Infisical に到達するシークレット) をサポートしたいと考えています。
Infisical Kubernetes Operator の V2、 v1beta1 では、3 つの CRD が導入されています。
InfisicalConnectio

n は、Infisical インスタンスのアドレスとオプションの TLS 設定を定義します。
InfisicalAuth は、マシン ID の認証の詳細を定義し、接続を参照します。
InfisicalStaticSecret は同期を定義し、認証リソースを参照します。これは InfisicalSecret を置き換えます。
注: 現在も InfisicalDynamicSecret と InfisicalPushSecret の CRD に取り組んでいます。
InfisicalConnection と InfisicalAuth は一度定義され、秘密リソースがそれらを指します。これにより、InfisicalStaticSecret は、独自のクライアントを実行せずにシークレットを検索、取得、調整できるようになります。
接続リソースの例を次に示します。
APIバージョン: Secrets.infisical.com/v1beta1
種類: InfisicalConnection
メタデータ:
名前: 私の非フィジカルコネクション
仕様:
アドレス: https://app.infisical.com
認証を定義する方法は次のとおりです。
APIバージョン: Secrets.infisical.com/v1beta1
種類: InfisicalAuth
メタデータ:
名前: prod-auth
仕様:
infisicalConnectionRef:
名前: 私の非フィジカルコネクション
名前空間: デフォルト
方法: ユニバーサル
普遍的:
clientIdRef:
名前: ユニバーサル認証資格情報
名前空間: デフォルト
キー: clientId
clientSecretRef:
名前: ユニバーサル認証資格情報
名前空間: デフォルト
キー: clientSecret
最終的な InfisicalStaticSecret リソースは次のようになります。
APIバージョン: Secrets.infisical.com/v1beta1
種類: InfisicalStaticSecret
メタデータ:
名前: サービス・ア・シークレット
仕様:
infisicalAuthRef:
名前: prod-auth
名前空間: デフォルト
情報源:
- プロジェクト ID: <プロジェクト ID>
環境スラグ: prod
シークレットパス: /service-a
ターゲット:
- 名前: 管理されたサービス
名前空間: デフォルト
種類：シークレット
作成ポリシー: 所有者
InfisicalAuth が構成の一部になりました。コントローラーは、アイデンティティによってキャッシュされたクライアントへの各 InfisicalStaticSecret 認証参照を解決します。これにより、リソースごとに 1 つではなく、アイデンティティごとに 1 つのクライアントが作成されます。を指す任意の InfisicalStaticSecret

同じ ID が認証と接続を再利用します。
これにより、以前の問題が修正されます。
再起動では、リソースごとに 1 つではなく、アイデンティティごとに 1 つの認証呼び出しが生成されます。
Infisical ホストまたはマシン ID を変更するには、1 つのリソースを編集するだけで済みます。
クライアントはすべてのリソースにわたってレプリケートされるわけではないため、メモリ使用量が削減されます。
キャッシュは遅延しますが、設定の変更時、またはコールが 401 または 403 で戻ってくると無効になります。その後、オペレータはキャッシュされたクライアントを削除し、再認証します。
メモリ、認証、接続を解決するには、ID に基づいてクライアントを内部的に重複排除することもできます。開発者のエクスペリエンスを向上させるために、代わりにモノリシック CRD を 3 つに分割しました。当社のオペレーターは現在、ESO のよく知られたパターンを反映しており、構成変更時の CRD の一括編集を不要にしています。
再設計された CRD に加えて、v2 では 3 つの小さなアップグレードも行いました。
InfisicalStaticSecret は複数のソース パスからプルできます。これにより、ワークロードがさまざまなパスからシークレットを使用できるようになります。プラットフォーム チームが /shared/auth で認証資格情報を所有し、アプリ チームが /app/integrations で統合キーを所有している場合、CRD はこれをネイティブにサポートします。
複数のターゲットに書き込むことができます。ターゲットは Kubernetes シークレットまたは ConfigMap にすることができます。これにより、機密性の高い値を Secret に取り込み、機密性のない値を ConfigMap に取り込むことが可能になります。これは、ConfigMap マウントのみを受け入れるサイドカーに値をフェッチする場合にも役立ちます。
すべてのリソースが準備状況を報告するようになったので、kubectl get は各リソースの正常性を示します。
Infisical Kubernetes オペレーター v2 へのアップグレード
現在、 v1alpha1 と v1beta1 の両方を維持していますが、最終的には v1alpha1 API を非推奨にする予定です。当社のドキュメントには、移行ガイドと、Kubernetes opera のインストールと使用に関する詳細な手順が含まれています。

より一般的には。
ソフトウェアのカテゴリによってエンジニアリングの優先順位が決まります。プロダクト マネージャーやエンジニアは問題トラッカーに何時間も費やしているため、Linear は低レイテンシ、鮮明な UI、キーボード ショートカットや cmd+k ナビゲーションなどの UX アフォーダンスに多額の投資を行っています。
セキュリティ カテゴリでは、製品はあらゆるインフラストラクチャ、導入モデル、ワークロード サイズで動作する必要があります。ギャップや並列システムは、自動化や集中化によって防止されるはずだった摩擦を引き起こします。回避策、重複シークレット マネージャー、または「今回はここでプレーンテキスト シークレットを使用しましょう」は潜在的な脆弱性です。
理想的な世界では、Infisical は、3 つの主要なクラウドすべて、専用 VPS、1990 年代から維持されているサーバー地下室で同時に当社を使用している企業の機密を保護できるはずです。
Kubernetes は組織のあらゆる規模で人気が高まっているため、Infisical での Kubernetes シークレットの管理はシームレスである必要があります。このため、Kubernetes オペレーターのスケーラビリティなど、一見小さなことが重要になります。
テクニカル コンテンツ マーケティング担当者、Infisical
技術的な前の記事 Infisical で始める SSL/TLS および mTLS 証明書管理の完全ガイドは、シンプル、高速、そして無料です。はじめる デモを入手する 製品の秘密管理

## Original Extract

Learn why Infisical rebuilt its Kubernetes Operator with a reference-based architecture that reduces auth calls, improves scaling, and simplifies configuration.

Platform Secrets Management PKI PAM Agent Vault Pricing Docs Resources Blog Videos Case Studies Compare Infisical vs Hashicorp Vault Careers 27k Request a demo Get started for free Platform Pricing Docs Resources Careers 27k stars Request a demo Get started for free ← Back Blog post • 7 min read Our Kubernetes Operator Didn’t Scale, So We Rebuilt It
Security is often at odds with convenience, but the human brain prefers convenience (and makes mistakes, even with the best of intentions). Most identity security tools reconcile this by making security as convenient as possible.
Any friction increases the risk that people build convenient workarounds or sidestep security tools and thus soften the organization’s security posture.
This is why we rearchitected our Kubernetes operator to improve performance and developer experience when we realized it struggled at scale.
Why we built a Kubernetes operator
One of the maxims of good secrets management is centralization. Uniting all secrets provides one place to store, manage, and audit secrets across your infrastructure.
Centralization requires syncing secrets into every type of infrastructure and deployment model. This means delivering secrets from the secret store into the user’s infrastructure to the service that consumes the secret. Ideally, this happens without workarounds or custom logic to avoid creating security gaps or placing maintenance burden on users.
Our first Kubernetes operator extended native secret syncs into distributed deployments. It worked, but didn’t scale well. As resources in a cluster or deployment proliferated, its memory footprint ballooned and performance degraded.
This is why we redesigned our operator with a new reference-based architecture.
Where our first Kubernetes operator faltered
Syncing secrets into Kubernetes at scale required our operator to do a few things:
Connect and authenticate to Infisical: Where the API is hosted, how to connect and authenticate to it.
Find the correct secrets in Infisical: Know the correct secret path within Infisical: which project, environment, folder, etc.
Enable pods and deployments to use them: Reconcile the secrets into Kubernetes-native secret objects so deployments and pods can get secrets into the correct environments.
Our initial design checked those boxes, but faltered as workloads increased.
Why v1’s architecture struggled at scale
In v1, users wrote monolithic custom InfisicalSecret resources that pointed at Infisical secrets. Each InfisicalSecret resource on the v1alpha1 API contained:
The address of the Infisical instance
The authentication credentials
The managed Kubernetes secret to write to
apiVersion: secrets.infisical.com/v1alpha1
kind: InfisicalSecret
metadata:
name: service-a-secrets
spec:
hostAPI: https://app.infisical.com/api
authentication:
universalAuth:
credentialsRef:
secretName: universal-auth-credentials
secretNamespace: default
secretsScope:
projectSlug: my-project
envSlug: prod
secretsPath: "/service-a"
managedSecretReference:
secretName: service-a-managed
secretNamespace: default
Scalability suffered because each resource replicated authentication and connection. That architecture works on persistent infrastructure. A VPS or VM is one identity that only reauthenticates on restart or config changes. Kubernetes clusters, however, contain dozens or hundreds of these resources and frequently redeploy and restart. Because each carried its own auth and connection, each held its own independent client. This created three problems:
Resources consumed outsized memory because each resource held its own client in memory. At scale, it created out-of-memory issues that required raising Helm memory limits. Each time pods went OOM, each resource reconciled and authenticated at the same time.
Restarts produced a burst of simultaneous authentication calls, which ran into rate limits. The operator would succeed after backoffs and retries, but it created latency in getting clusters to a steady state.
Engineering teams had to do extra work. Rotating a machine identity or changing the Infisical host meant editing the authentication block on every single resource.
The fundamental issue was the overloaded CRD architecture, not missing logic. We evaluated event handlers, jitter, and other logic. Those may have helped, but added more complexity to an already overloaded CRD.
We could only solve the underlying issue with a new architecture.
How reference-based architecture fixed the replication
The new design separates connection, authentication, and sync. Secrets reference authentication and connection resources as shared objects. This fixes the performance issues and improves the developer experience.
We modeled the new architecture roughly on External Secrets Operator ’s (ESO) resource split, which separates provider , store , and externalsecret CRDs. Infisical integrates with ESO , but we build our own operator for two reasons:
ESO has previously paused development. It has since resumed, but it’s not ideal to create a dependency that may stop being developed.
We want to offer native UX, minimize moving parts, and eventually support dynamic secrets and push secrets (secrets that originate in Kubernetes and find their way to Infisical), which ESO doesn’t support.
V2 of the Infisical Kubernetes Operator, v1beta1 , introduces three CRDs:
InfisicalConnection defines the address of an Infisical instance and optional TLS settings.
InfisicalAuth defines the authentication details for a machine identity and references a connection.
InfisicalStaticSecret defines a sync and references an auth resource. It replaces InfisicalSecret .
Note: We’re still working on CRDs for InfisicalDynamicSecret and InfisicalPushSecret .
InfisicalConnection and InfisicalAuth are defined once, and secret resources point at them. This ensures InfisicalStaticSecret can find, pull, and reconcile secrets without running its own client.
Here’s an example of a connection resource:
apiVersion: secrets.infisical.com/v1beta1
kind: InfisicalConnection
metadata:
name: my-infisical-connection
spec:
address: https://app.infisical.com
And here’s how you can define authentication:
apiVersion: secrets.infisical.com/v1beta1
kind: InfisicalAuth
metadata:
name: prod-auth
spec:
infisicalConnectionRef:
name: my-infisical-connection
namespace: default
method: universal
universal:
clientIdRef:
name: universal-auth-credentials
namespace: default
key: clientId
clientSecretRef:
name: universal-auth-credentials
namespace: default
key: clientSecret
A final InfisicalStaticSecret resource looks like this:
apiVersion: secrets.infisical.com/v1beta1
kind: InfisicalStaticSecret
metadata:
name: service-a-secrets
spec:
infisicalAuthRef:
name: prod-auth
namespace: default
sources:
- projectId: <your-project-id>
environmentSlug: prod
secretPath: /service-a
targets:
- name: service-a-managed
namespace: default
kind: Secret
creationPolicy: Owner
InfisicalAuth is now part of configuration. The controller resolves each InfisicalStaticSecret auth reference to client cached by identity, which creates one client per identity, not one per resource. Any InfisicalStaticSecret pointing at the same identity reuses authentication and connection.
This fixes the previous issues:
A restart produces one authentication call per identity, not one per resource.
Changing the Infisical host or machine identity only requires editing one resource.
Clients aren’t replicated across every resource, which shrinks the memory footprint.
The cache is lazy, but invalidates on config changes or when a call comes back with a 401 or 403. The operator then drops the cached client and reauthenticates.
To solve memory, authentication, and connection, we could’ve deduplicated clients internally by identity. We split the monolithic CRD into three instead to improve the developer experience. Our operator now mirrors the well-known pattern from ESO and obviates mass-editing CRDs when configuration changes.
Besides the rearchitected CRDs, we also made three smaller upgrades in v2:
InfisicalStaticSecret can pull from more than one source path. This allows workloads to consume secrets from different paths. If a platform team owns auth credentials at /shared/auth while app teams own integration keys at /app/integrations , the CRD supports this natively.
It can write to more than one target, where a target can be a Kubernetes secret or a ConfigMap . This enables pulling sensitive values into a Secret and non-sensitive ones into a ConfigMap . This is also useful for fetching values to sidecars that only accept ConfigMap mounts.
Every resource now reports a readiness status, so kubectl get shows the health of each resource.
Upgrading to Infisical Kubernetes operator v2
We currently maintain both v1alpha1 and v1beta1 , but plan to deprecate the v1alpha1 API eventually. Our documentation contains a migration guide as well as detailed instructions for installing and using our Kubernetes operator more generally.
Software categories dictate engineering priorities. Product managers and engineers spend hours in issue trackers, so Linear invested heavily in low latency, a crisp UI, and UX affordances like keyboard shortcuts and cmd+k navigation.
In the security category, products need to work on any infrastructure, deployment model, and workload size. Gaps or parallel systems introduce the friction that automation and centralization were supposed to prevent. Any workaround, duplicate secrets manager, or “let’s just use plaintext secrets here, this one time” is a potential vulnerability.
In an ideal world, Infisical should be able to secure secrets for a company that simultaneously uses us on all three major clouds, a dedicated VPS, and the server basement they’ve been maintaining since the 1990s.
Kubernetes is becoming more popular across all organization sizes, so managing Kubernetes secrets on Infisical needs to be seamless. This makes seemingly small things like the scalability of a Kubernetes operator important.
Technical Content Marketer, Infisical
Technical Previous Article A Complete Guide to SSL/TLS & mTLS Certificate Management Starting with Infisical is simple, fast, and free. Get Started Get a demo PRODUCT Secrets Management

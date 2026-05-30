---
source: "https://www.romaglushko.com/blog/k8s-gateway-api/"
hn_url: "https://news.ycombinator.com/item?id=48325848"
title: "Deep Dive into Kubernetes Gateway API"
article_title: "Kubernetes Gateway API - Blog by Roman Glushko"
author: "roma_glushko"
captured_at: "2026-05-30T11:43:00Z"
capture_tool: "hn-digest"
hn_id: 48325848
score: 3
comments: 1
posted_at: "2026-05-29T16:53:22Z"
tags:
  - hacker-news
  - translated
---

# Deep Dive into Kubernetes Gateway API

- HN: [48325848](https://news.ycombinator.com/item?id=48325848)
- Source: [www.romaglushko.com](https://www.romaglushko.com/blog/k8s-gateway-api/)
- Score: 3
- Comments: 1
- Posted: 2026-05-29T16:53:22Z

## Translation

タイトル: Kubernetes Gateway API の詳細
記事のタイトル: Kubernetes Gateway API - Roman Glushko のブログ
説明: Ingress が置き換えられる理由とどのゲートウェイ コントローラーを選択するか

記事本文:
Kubernetes Gateway API - Roman Glushko のブログ
ローマ人
グルシュコ
ブログ
写真提供: Ani Adigyozalyan Kubernetes ゲートウェイ API
2025 年 11 月に、Kubernetes は NGINX Ingress コントローラー、
Kubernetes クラスターの 40% 以上に存在する最も広く採用されている Ingress コントローラーの 1 つが、2026 年 3 月に非推奨になります。
このニュースは、Kubernetes Gateway API の導入における転換点を示し、Gateway API を Ingress API の後継として位置づけました。
そもそもなぜ非推奨にするのでしょうか?なぜ Kubernetes が NGINX Ingress Controller を廃止することにしたのか疑問に思っているかもしれません。広く採用されていることを考えると、維持し続けるのは当然の選択のように思えるかもしれません。
しかし、最近のコントローラー CVE を見ると、あるパターンが見えてきます。プロジェクトには、今日のセキュリティ上の期待を満たすことが困難な、いくつかの基本的な設計上の欠陥があります。
メンテナは、新たに報告された脆弱性のたびにモグラたたきを続けるのではなく、一歩下がって、新しいゲートウェイ API 実装でこれらの問題に対処することにしました。
この投稿では、Gateway API、Ingress API との違い、移行のオプションについて説明します。
Kubernetes の初期の 2014 年頃、サービス リソースはアプリケーションを公開する主な Kubernetes ネイティブの方法でした。
当時、Kubernetes には、現在私たちが使用しているのとほぼ同じサービス タイプがすでにありました。
ClusterIP: サービスにクラスター内の安定した DNS 名を与えますが、クラスターの外部からアクセスできるようにはしません。
NodePort: クラスター内のすべてのノードで特定のポートを公開し、外部トラフィックがそのポートを介してサービスに到達できるようにします。ノード IP がパブリックにアクセス可能な場合、サービスが外部に公開される可能性があります。
LoadBalancer: パブリッシャーを使用して外部ロード バランサーをプロビジョニングします。

ic IP アドレスを使用してトラフィックをサービスに転送します。
ExternalName: CNAME レコードをクラスター DNS に追加することで、外部サービスにクラスター内エイリアスを提供します。
これら 4 つのうち、サービスをクラスター外に公開するのに直接役立つのは、NodePort と LoadBalancer だけです。
NodePort は基本的なプリミティブですが、現実世界のほとんどのユースケースには低レベルすぎます。
NodePort を使用すると、各 Service に対応するノード IP とポート番号のセットを取得できます。
ノード間の外部負荷分散機能は組み込まれていないため、自分で実装する必要があります。
内部クラスター トポロジを非表示にする単一の外部エンドポイントが必要な場合は、パスベースのルーティングを自分で処理する必要もあります。
サービスを公開する NodePort サービス
カスタム LoadBalancer と実際の Kubernetes クラスター設定の間でポート マッピングの同期を保つ必要があります。
それはセットアップがかなり面倒です。
その意味では、LoadBalancer Service タイプの方が便利です。
すべてのクラスター ノードでポートを公開することにより、NodePort の上に構築されます。
ただし、その NodePort にトラフィックを転送する外部 L4 ロード バランサ、TCP または UDP もプロビジョニングされます。
プロビジョニング部分は、クラウド コントローラー マネージャー (または CCM) のおかげで自動的に行われます。
LoadBalancer サービスを使用してサービスを公開する
これはあまり良くありません。理由は次のとおりです。
特定のサービスにアクセスするために、一連のパブリック IP を管理し、ルーティングする必要があります。実際には、ルーティングを処理するためにさらに別の外部プロキシが必要になる可能性があります。
外部 IP とロード バランサーの各インスタンスには費用がかかるため、数百ものインスタンスを維持するとすぐに費用が高くなってしまいます。
受信トラフィックを管理する単一の中央の場所は存在しません。通常、攻撃者がシステムに侵入しようとするときに最初に調査するのはイングレス エンドポイントであるため、これは重要です。
について考えてみると

この問題をもう少し詳しく見てみると、 Service ごとに個別の外部ロード バランサーが必要ないことがわかります。
代わりに、すべての受信トラフィックを受け入れ、それを NGINX などの専用のリバース プロキシ デプロイメントに転送する単一の LoadBalancer サービスを使用できます。
そのリバース プロキシは、リクエスト パス、ホスト名、またはその他のリクエスト プロパティに基づいて、トラフィックを適切なバックエンド サービスにルーティングします。
LoadBalancer サービスを使用してリバース プロキシでサービスを公開する
この設定により、最終的に Ingress API と Ingress コントローラーが誕生しました。
2015 年、Kubernetes チームは、外部 HTTP(S) トラフィックをクラスター内のサービスにルーティングするためのルールを定義する方法として、Ingress API を正式に導入しました。
Ingress API は、 IngressClass と Ingress という 2 つの主要なリソースで構成されます。
Kubernetes Ingress API
これは共通のインターフェイスを定義しますが、すぐに使用できる特定の実装は付属しません。
これらのルールを実際に機能させるには、ルールを解釈して適用する方法を認識する Ingress コントローラーをインストールする必要があります。
クラス アノテーション Ingress API の初期リリースでは、 IngressClass は存在しませんでした。
代わりに、Kubernetes はカスタムの kubernetes.io/ingress.class アノテーションを使用して、どの Ingress コントローラーが特定の Ingress リソースを担当するかを指定しました。
このアノテーションは正式に標準化されたことはありませんが、多くの Ingress コントローラーは、Ingress リソースとそのコントローラー間のマッピングを定義するための、非推奨の代替方法としてそれをサポートしています。
IngressClass は、Kubernetes に通知するクラスター全体のリソースです。
どのコントローラーが特定の Ingress リソースのセットを処理する必要があるか、またコントローラーがどのパラメーターを使用する必要があるか。
クラスによる拡張性 IngressClass は、戦略設計パターンを Kubernetes API に適用する例です。
定義できるようになります

e 1 つのクラスター内で同じ API の複数の実装を使用します。
それ以来、クラス リソースによる拡張性は、さまざまな実装を標準の Kubernetes API にプラグインする一般的な方法になりました。
例には、 RuntimeClass 、 StorageClass 、および VolumeSnapshotClass が含まれます。
IngressClass を使用すると、同じクラスター内で複数の Ingress コントローラーを並べて実行できます。
各コントローラーは、対応する IngressClass によって選択された独自の Ingress リソースのセットを監視します。
同じクラスター内の複数の Ingress コントローラー
apiバージョン : networking.k8s.io/v1
種類 : IngressClass
メタデータ:
名前：アルブ
仕様：
コントローラー: ingress.k8s.aws/alb
パラメータ:
スコープ : クラスター
APIグループ：elbv2.k8s.aws
種類 : IngressClassParams
名前: alb-ingress-params
---
APIバージョン : elbv2.k8s.aws/v1beta1
種類 : IngressClassParams
メタデータ:
名前: alb-ingress-params
仕様：
スキーム: 内部
ipAddressType : デュアルスタック
タグ :
- キー: コストセンター
値: 製品工学
ロードバランサー属性:
- キー : 削除保護.有効
値: " true "
- キー: idle_timeout.timeout_秒
値: " 120 "
IngressClass パラメーターは、追加の実装固有の構成を Ingress コントローラーに渡すためのオプションの方法です。
型付きオブジェクト参照 Ingress コントローラーでは追加の構成を提供できる場合がありますが、その構成は実装によって異なります。
パラメーターは基本的に、Ingress コントローラーによって定義されたカスタム リソースへの参照です。
これは、Kubernetes が IngressClass などの標準 API を実装固有の構成に接続する方法です。
このパターンは、型付きオブジェクト参照と呼ばれます。多くの場合、これは 4 つのフィールドを持つ単純なブロックのように見えます。
参照:
apiGroup : resource.k8s.aws
種類 : MyParams
名前 : パラメータ
namespace : ingress namespace フィールドは、次の場合に省略できます。

参照されるオブジェクトがクラスタ全体であるか、参照リソースが参照されるオブジェクトが同じ名前空間にあることを予期している場合。
たとえば、Ingress はバックエンド サービスが同じ名前空間にあることを期待します。
API がより具体的な構成セクションまたはスコープを指す必要がある場合、参照ブロックにはさらに多くのフィールドを含めることもできます。
これは、標準の Kubernetes API を拡張するためのもう 1 つの重要なパターンです。
Ingress コントローラーには、IngressClass リソースを読み取り、使用するための適切な ClusterRole 権限も必要です。
Ingress は、Ingress API の主要なリソースです。これはいくつかのことを組み合わせたものです:
完全一致や接頭辞一致などのパスベースの一致とホストベースのポリシーを使用して、受信リクエストを照合し、それらを内部サービスにルーティングするためのルールを定義します。
受信トラフィックの TLS 構成を定義します。
apiバージョン : networking.k8s.io/v1
種類 : イングレス
メタデータ:
名前 : イングレス
名前空間 : アプリケーション
仕様：
ingressClassName : alb # この Ingress リソースを処理する IngressClass への参照
TLS :
- hosts : # これらのホストの HTTPS トラフィックを処理します
- app.romaglushko.com
SecretName : ingress-tls-secret # TLS 秘密キーと証明書が含まれます
ルール：
- host : app.romaglushko.com # ホストベースのルーティング
http :
paths : # パスベースのルーティング
- パス : /users
パスタイプ : プレフィックス
バックエンド:
サービス：
name : users # 内部マイクロサービス サービス名
ポート:
数 : 8000
- パス : /orders
パスタイプ : プレフィックス
バックエンド:
サービス：
name :orders # 内部マイクロサービス サービス名
ポート:
数 : 8000
Ingress リソースが作成されると、Ingress コントローラーはそれを取得し、独自の構成を更新して新しいルーティング ルールを適用します。
それでおしまい。
では、Ingress API には何が問題なのでしょうか?
サービスを外部の世界に公開する簡単な解決策のように思えますよね?
そうですね、うん

実際のセットアップで使用し始めるまでは。ここでいくつかの問題が発生します。
安全なクロスネームスペースのサポートなし
問題は、Ingress API が単純であるというよりむしろ単純であることです。
ここでは、入力トラフィックに対して通常設定する必要のある最低限のもののみを取り上げています。
NGINX などに必要な機能について考えてみましょう。
ヘッダー、クエリ、Cookie ベースのルーティング
Ingress リソースは、これらを直接サポートしません。
そのため、カスタム アノテーションが追加の設定を Ingress コントローラーに渡す標準的な方法になりました。
Traefik などの一部のコントローラーは、代わりに独自の CRD を使用します。
apiバージョン : networking.k8s.io/v1
種類 : イングレス
メタデータ:
名前 : イングレス
名前空間 : アプリケーション
注釈 :
nginx.ingress.kubernetes.io/proxy-read-timeout : " 60 "
nginx.ingress.kubernetes.io/proxy-send-timeout : " 60 "
nginx.ingress.kubernetes.io/enable-cors : " true "
nginx.ingress.kubernetes.io/cors-allow-origin : " * "
nginx.ingress.kubernetes.io/cors-allow-methods : " GET、POST、OPTIONS "
仕様：
ingressClassName : nginx
ルール：
- host : app.romaglushko.com # ホストベースのルーティング
http :
paths : # パスベースのルーティング
- パス: /static
パスタイプ : プレフィックス
バックエンド:
サービス：
名前 : ウェブスタティック
ポート:
数 : 3000
Ingress リソースが同時に複数の Ingress コントローラーと連携する必要がある場合はどうすればよいでしょうか?
これらの共通機能を構成する統一された方法がないため、通常、各 Ingress コントローラーには独自のアノテーションのセットがあります。
有効にする機能が増えるほど、Ingress リソースの移植性は低くなります。
さらに、Ingress は HTTP(S) ルーティングのみを考慮します。
gRPC (L7)、TCP、UDP (L4) などの他のプロトコルのサポートは、特定の実装に完全に依存しており、通常はカスタム アノテーションを通じて構成されます。
カスタム注釈は、

Ingress API を拡張するかなりハックな方法です。
これらは単なる文字列ベースのキーと値のペアです。
しかし、多くの場合、より複雑な設定を渡す必要があるため、最終的には文字列にシリアル化し、それをアノテーションに入れることになります。
機能しますが、あまり表現力がありません。
アプリケーション チームが Ingress リソースを作成するとき、
Ingress コントローラーの動作を変更するアノテーションを含む、基本的に任意のアノテーションを追加できます。
たとえば、プラットフォーム チームがすべてのルートでレート制限を有効にしていると想定している場合でも、レート制限を無効にする場合があります。
Ingress API には、これらのグローバルな動作を強制する組み込みメカニズムはありません。
最善の方法は、Kyverno や OPA Gatekeeper などの外部ポリシー エンジンを使用することです。
Ingress リソースには、いくつかのタイプの構成が混在しています。
ルーティング ルール。通常、アプリケーション チームが所有します。アプリケーション チームは API がどのように動作すべきかを知っているためです。
ホスト名とポートの構成。DNS、ロード バランサー、VPC/ネットワーク設定を管理するため、通常はプラットフォーム チームが所有します。
TLS 構成。証明書のプロビジョニングを処理するため、通常はプラットフォーム チームが所有します。
カスタム アノテーション。アプリケーション チームまたはプラットフォーム チームが所有する機能を有効にする場合があります。
では、Ingress をどこに置くべきか

[切り捨てられた]

## Original Extract

Why Ingress Is Being Replaced and Which Gateway Controller to Pick

Kubernetes Gateway API - Blog by Roman Glushko
Roman
Glushko
Blog
Photo by Ani Adigyozalyan Kubernetes Gateway API
In November 2025, Kubernetes announced that NGINX Ingress Controller ,
one of the most widely adopted ingress controllers that present in more than 40% of Kubernetes clusters, would be deprecated in March 2026 .
That news has marked a turning point in the Kubernetes Gateway API adoption , positioning Gateway API as the successor to the Ingress API.
Why Deprecate It at All? You may still be wondering why Kubernetes decided to deprecate NGINX Ingress Controller. Given the widespread adoption, continuing to maintain it may seem like an obvious choice.
However, if you look at the recent Controller CVEs , a pattern starts to emerge: the project has several fundamental design flaws that make it difficult to meet today’s security expectations.
Instead of continuing to play whac-a-mole with each newly reported vulnerability, the maintainers decided to step back and address these problems in newer Gateway API implementations .
In this post, we will explore Gateway API, how it differs from the Ingress API, and what options you have for migration.
In the early days of Kubernetes, around 2014, Service resources were the main Kubernetes-native way to expose applications.
At that time, Kubernetes already had roughly the same Service types ) we use today:
ClusterIP: Gives the Service a stable DNS name inside the cluster, but does not make it accessible from outside the cluster.
NodePort: Exposes a specific port on every node in the cluster, allowing external traffic to reach the Service through that port. If the node IPs are publicly accessible, this can expose the Service to the outside world.
LoadBalancer: Provisions an external load balancer with a public IP address and forwards traffic to the Service.
ExternalName: Provides an in-cluster alias for an external service by adding a CNAME record to the cluster DNS.
Out of these four, only NodePort and LoadBalancer can directly help us expose Services outside the cluster.
NodePort is a fundamental primitive, but it is too low-level for most real-world use cases.
With NodePort , you get a set of node IPs and port numbers that correspond to each Service .
There is no built-in external load balancing between nodes, so you would need to implement that yourself.
You would also need to handle path-based routing yourself if you wanted a single external endpoint that hides the internal cluster topology.
NodePort Service to Expose Services
You would need to keep the port mapping in sync between that custom LoadBalancer and actual Kubernetes cluster setup.
That’s rather tedious to setup.
The LoadBalancer Service type is more useful in that sense.
It builds on top of NodePort by exposing a port on all cluster nodes,
but it also provisions an external L4 load balancer, TCP or UDP, that forwards traffic to that NodePort .
The provisioning part happens automatically thanks to the Cloud Controller Manager (or CCM).
Using LoadBalancer Service to Expose Services
That is not great because:
You would have a set of public IPs to manage and route between in order to reach specific Services . In practice, you would likely need yet another external proxy to handle the routing.
Each external IP and load balancer instance costs money, so maintaining hundreds of them would quickly become expensive.
There would be no single central place to manage incoming traffic. That matters because ingress endpoints are usually the first place attackers probe when trying to get into a system.
If we think about this problem a bit more, we can see that we do not need a separate external load balancer for every Service .
Instead, we could have a single LoadBalancer Service that accepts all incoming traffic and forwards it to a dedicated reverse proxy Deployment , such as NGINX.
That reverse proxy would then route traffic to the right backend Service based on the request path, hostname, or other request properties.
Using LoadBalancer Service to Expose Services with a Reverse Proxy
This setup is what eventually gave birth to the Ingress API and Ingress Controllers .
In 2015, the Kubernetes team officially introduced the Ingress API as a way to define rules for routing external HTTP(S) traffic to Services inside the cluster.
The Ingress API consists of two main resources: IngressClass and Ingress .
Kubernetes Ingress API
It defines a common interface, but it doesn’t come with a specific implementation out of the box.
To make these rules actually work, you need to install an Ingress Controller that knows how to interpret and apply them.
The Class Annotation In the initial release of Ingress API, there was no IngressClass .
Instead, Kubernetes used the custom kubernetes.io/ingress.class annotation to specify which Ingress Controller should be responsible for a given Ingress resource.
That annotation was never formally standardized, but many Ingress Controllers still support it as an alternative, deprecated way to define the mapping between an Ingress resource and its controller.
IngressClass is a cluster-wide resource that tells Kubernetes
which controller should handle a given set of Ingress resources, and which parameters that controller should use.
Extensibility via Classes IngressClass is an example of applying the strategy design pattern to Kubernetes APIs.
It lets you define and use multiple implementations of the same API within one cluster.
Since then, extensibility via Class resources has become a common way to plug different implementations into standard Kubernetes APIs.
Examples include RuntimeClass , StorageClass , and VolumeSnapshotClass .
IngressClass makes it possible to run multiple Ingress Controllers side by side in the same cluster.
Each controller watches its own set of Ingress resources, selected by the corresponding IngressClass .
Multiple Ingress Controllers in the same cluster
apiVersion : networking.k8s.io/v1
kind : IngressClass
metadata :
name : alb
spec :
controller : ingress.k8s.aws/alb
parameters :
scope : Cluster
apiGroup : elbv2.k8s.aws
kind : IngressClassParams
name : alb-ingress-params
---
apiVersion : elbv2.k8s.aws/v1beta1
kind : IngressClassParams
metadata :
name : alb-ingress-params
spec :
scheme : internal
ipAddressType : dualstack
tags :
- key : costcenter
value : product-eng
loadBalancerAttributes :
- key : deletion_protection.enabled
value : " true "
- key : idle_timeout.timeout_seconds
value : " 120 "
IngressClass parameters are an optional way to pass additional, implementation-specific configuration to the Ingress Controller.
Typed Object References Ingress Controllers may allow you to provide additional configuration, but that configuration differs between implementations.
Parameters are basically a reference to a Custom Resource defined by the Ingress Controller.
This is how Kubernetes can connect a standard API, such as IngressClass , with implementation-specific configuration.
This pattern is called a typed object reference. It often looks like a simple block with four fields:
ref :
apiGroup : resource.k8s.aws
kind : MyParams
name : params
namespace : ingress The namespace field can be omitted if the referenced object is cluster-wide, or if the referencing resource expects the referenced object to be in the same namespace.
For example, Ingress expects its backend Service to be in the same namespace .
The reference block can also have more fields when the API needs to point to a more specific configuration section or scope.
This is another important pattern for extending standard Kubernetes APIs.
Ingress Controllers also need the right ClusterRole permissions to read and use IngressClass resources.
Ingress is the main resource in the Ingress API. It combines a few things:
It defines rules for matching incoming requests and routing them to internal Services using path-based matches, such as exact or prefix matches, and host-based policies.
It defines TLS configuration for incoming traffic.
apiVersion : networking.k8s.io/v1
kind : Ingress
metadata :
name : ingress
namespace : application
spec :
ingressClassName : alb # Reference to the IngressClass that should handle this Ingress resource
tls :
- hosts : # Handle HTTPS traffic for these hosts
- app.romaglushko.com
secretName : ingress-tls-secret # contains a TLS private key and certificate
rules :
- host : app.romaglushko.com # Host-based routing
http :
paths : # Path-based routing
- path : /users
pathType : Prefix
backend :
service :
name : users # Internal microservice Service name
port :
number : 8000
- path : /orders
pathType : Prefix
backend :
service :
name : orders # Internal microservice Service name
port :
number : 8000
Once the Ingress resource is created, the Ingress Controller picks it up and updates its own configuration to apply the new routing rule.
That’s it.
So what’s the problem with the Ingress API?
It seems like a simple solution to exposing Services to the outside world, right?
Well, yes, until you start using it in real life setups. That’s where a few problems show up:
No Secure Cross-Namespace Support
The problem is that the Ingress API is simplistic rather than simple.
It covers only the bare minimum of what we usually want to configure for ingress traffic.
Just think about the features you might want from something like NGINX:
header-, query-, cookie-based routing
The Ingress resource supports none of these directly.
Because of that, custom annotations became the canonical way to pass extra configuration to Ingress Controllers.
Some controllers, like Traefik , use their own CRDs instead.
apiVersion : networking.k8s.io/v1
kind : Ingress
metadata :
name : ingress
namespace : application
annotations :
nginx.ingress.kubernetes.io/proxy-read-timeout : " 60 "
nginx.ingress.kubernetes.io/proxy-send-timeout : " 60 "
nginx.ingress.kubernetes.io/enable-cors : " true "
nginx.ingress.kubernetes.io/cors-allow-origin : " * "
nginx.ingress.kubernetes.io/cors-allow-methods : " GET, POST, OPTIONS "
spec :
ingressClassName : nginx
rules :
- host : app.romaglushko.com # Host-based routing
http :
paths : # Path-based routing
- path : /static
pathType : Prefix
backend :
service :
name : webstatic
port :
number : 3000
What if your Ingress resources need to work with multiple Ingress Controllers at the same time?
Because there is no unified way to configure these common features, each Ingress Controller usually has its own set of annotations.
The more features you enable, the less portable your Ingress resources become.
On top of that, Ingress is only concerned with HTTP(S) routing.
Support for other protocols, such as gRPC (L7), TCP, or UDP (L4), depends entirely on the specific implementation and is usually configured through custom annotations.
Custom annotations are a pretty hacky way to extend the Ingress API.
They are just string-based key-value pairs.
But often, you need to pass more complex configuration, so you end up serializing it into a string and putting it into an annotation.
It works, but it is not very expressive.
When application teams create Ingress resources,
they can add basically any annotations to them, including annotations that change the behavior of the Ingress Controller.
For example, they might disable rate limiting even though the platform team expects it to be enabled for all routes.
There is no built-in mechanism in the Ingress API to enforce these global behaviors.
The best you can do is use an external policy engine like Kyverno or OPA Gatekeeper .
Ingress resources mix several types of configuration:
Routing rules, usually owned by application teams because they know how their API should behave.
Hostname and port configuration, usually owned by platform teams because they manage DNS, load balancers, and VPC/networking setup.
TLS configuration, usually owned by platform teams because they handle certificate provisioning.
Custom annotations, which may enable functionality owned by either application or platform teams.
So where should Ingress re

[truncated]

---
source: "https://www.freecodecamp.org/news/implement-zero-trust-workload-identity-in-kubernetes-with-spiffe-spire-and-cilium/"
hn_url: "https://news.ycombinator.com/item?id=48828639"
title: "Zero-Trust Workload Identity in Kubernetes with Spiffe, Spire, and Cilium"
article_title: "How to Implement Zero-Trust Workload Identity in Kubernetes with SPIFFE, SPIRE, and Cilium"
author: "enz"
captured_at: "2026-07-08T07:31:14Z"
capture_tool: "hn-digest"
hn_id: 48828639
score: 1
comments: 0
posted_at: "2026-07-08T07:16:32Z"
tags:
  - hacker-news
  - translated
---

# Zero-Trust Workload Identity in Kubernetes with Spiffe, Spire, and Cilium

- HN: [48828639](https://news.ycombinator.com/item?id=48828639)
- Source: [www.freecodecamp.org](https://www.freecodecamp.org/news/implement-zero-trust-workload-identity-in-kubernetes-with-spiffe-spire-and-cilium/)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T07:16:32Z

## Translation

タイトル: Spiffe、Spire、Cilium を使用した Kubernetes のゼロトラスト ワークロード ID
記事のタイトル: SPIFFE、SPIRE、Cilium を使用して Kubernetes にゼロトラスト Workload Identity を実装する方法
説明: ネットワーク ポリシーでは、10.0.1.45 からのトラフィックを許可します。昨日、10.0.1.45 が支払いサービスでした。現在、ローリング デプロイメントの後は、これがログ エージェントになります。現在の支払いサービスは 10.0.1 です。

記事本文:
SPIFFE、SPIRE、Cilium を使用して Kubernetes にゼロトラスト Workload Identity を実装する方法
ネットワーク ポリシーでは、 10.0.1.45 からのトラフィックを許可します。
昨日、10.0.1.45 が支払いサービスでした。現在、ローリング デプロイメントの後は、これがログ エージェントになります。現在の支払いサービスは 10.0.1.89 です。
Kubernetes はすでにすべてのエンドポイントとサービス レコードを更新していますが、ネットワーク ポリシーはまったくわかりません。信頼する予定だったワークロードに属さなくなった IP アドレスに基づいて、トラフィックの通過がサイレントに許可されます。
これがワークロードアイデンティティの問題です。 IP アドレスは ID ではなく、場所です。また、Kubernetes クラスターでは、場所は常に変化します。 IP アドレスに基づいてセキュリティ ポリシーを構築すると、ポッドがスケジュール、再スケジュール、またはスケーリングされるたびに、セキュリティ体制が静かに低下することを意味します。
答えは、暗号化されたワークロード ID です。すべてのワークロードは、どこにあるかではなく、誰であるかを証明する証明書に裏付けられた ID を取得します。サービスは、データを交換する前に、これらの証明書を使用して相互に認証します。証明書が一致しない場合、接続元の IP アドレスに関係なく、接続は拒否されます。
これが SPIFFE と SPIRE が提供するものです。これは、Cilium がすべてのポッドにサイドカーを挿入せずに、eBPF を使用してそれを強制する方法です。
この記事では、SPIFFE ID モデルの仕組み、SPIRE をデプロイしてワークロードに暗号化 ID を発行する方法、および Cilium の組み込み SPIRE 統合を使用してアプリケーション コードに触れることなくサービス間で相互 TLS を強制する方法を理解します。
Kubernetes RBAC とポッド セキュリティに関する知識 - このハンドブックは基礎をカバーしています
TLS 証明書と Kubernetes Secret に精通していること — このハンドブックでは、証明書マネージャーと証明書の概念について説明しています
Helm 3 と Cilium CLI がインストールされている

種類クラスター — この記事では CNI として Cilium を使用して新しいクラスターを作成します
忍耐: これは、この記事グループで取り上げた中で最も複雑なデモです。 SPIRE には、これまで取り上げたものよりも多くの可動部分があります。
すべてのデモ ファイルは、コンパニオン GitHub リポジトリにあります。
SVID: 暗号化アイデンティティ文書
Cilium が SPIFFE を使用して相互 TLS を実装する方法
デモ 1 — SPIRE 統合を使用して Cilium をインストールする
ステップ 1: Cilium CLI をインストールする
ステップ 2: デフォルトの CNI を使用せずに種類クラスターを作成する
ステップ 3: SPIRE を有効にして Cilium をインストールする
ステップ 4: インストールを確認する
デモ 2 — CiliumNetworkPolicy を使用して相互 TLS を強制する
ステップ 1: クライアントとサーバーを展開する
ステップ 2: 認証なしでトラフィック フローを確認する
ステップ 3: 相互認証を必要とする CiliumNetworkPolicy を適用する
ステップ 4: 認証されたトラフィックが引き続き流れることを確認する
ステップ 5: Hubble による認証を観察する (オプション)
ステップ 6: 一致するラベルのないポッドがブロックされていることを確認する
ステップ 7: SPIRE のワークロード エントリを確認する
冒頭のシナリオは理論的なものではありません。 Kubernetes では、ポッドは一時的です。スケジューラは任意のノードにポッドを配置でき、ポッドの IP アドレスはスケジュール時にノードの IP プールから割り当てられます。
ポッドが削除され、ローリング デプロイメント、ノード ドレイン、またはオートスケーラー イベントを通じて再作成されると、新しい IP アドレスが取得されます。 「この IP からのトラフィックを許可する」という NetworkPolicy を作成した場合、そのポリシーは何も指していないか、さらに悪いことに別のワークロードを指していることになります。
ここでは、Kubernetes サービス名が East-West トラフィックに役立ちます。サービス名は、どのポッドがそれをサポートしているかに関係なく、一貫して解決されます。ただし、サービス名に基づく NetworkPolicy は依然としてラベル セレクターの一致であり、暗号化アサーションではありません。適切なラベルを偽装できるポッドは、それをバイパスできます。
本当に欲しいもの

これは次のとおりです。サービス A がサービス B にリクエストを送信する前に、サービス B はその ID を暗号的に証明します。サービス B が本人であることを証明できない場合、サービス A は接続を拒否します。これは相互 TLS であり、重要な問題は、ID がどこから来たのかということです。
SPIFFE (Secure Production Identity Framework forEveryone) は、ワークロード ID のモデルを定義する CNCF 標準です。それ自体では何も実装されません。 ID の形式、ID を要求するための API、サービス、クラスター、クラウド全体で ID を検証できるようにする信頼モデルを指定します。 SPIRE は、その仕様のリファレンス実装です。
SPIFFE ID は、特定の形式の URI です。
spiffe://<信頼ドメイン>/<ワークロードパス>
信頼ドメインは、管理境界 (通常は組織、クラスター、または環境) を識別する文字列です。同じ信頼ドメイン内のすべてのものは、互いの ID を検証できます。異なる信頼ドメインの ID には、明示的なフェデレーション構成が必要です。
spiffe://payments.corp/ns/production/sa/checkout
spiffe://analytics.corp/ns/data/sa/pipeline-worker
spiffe://cluster.local/ns/monitoring/sa/prometheus
信頼ドメインの後のパスは任意です。これは SPIRE 構成によって定義され、通常はワークロードの Kubernetes 名前空間とサービス アカウントをエンコードします。
SVID: 暗号化アイデンティティ文書
SVID (SPIFFE 検証可能な ID ドキュメント) は、SPIFFE ID をサービスが実際に使用できるものに具体化する方法です。
X.509 SVID は、SPIFFE ID がサブジェクト代替名 (SAN) URI フィールドに埋め込まれている標準の TLS 証明書です。これは標準の X.509 証明書であるため、どの TLS ライブラリでも変更せずに使用できます。
ワークロードはこの証明書を次のように提示します。

TLS ハンドシェイクにより、ピアは証明書が信頼できる SPIRE サーバーによって署名されたことを検証します。これは、gRPC ストリームのような長期間存続する接続に使用される形式です。
JWT SVID は、クレームとして SPIFFE ID を含む署名付き JSON Web トークンです。これは、HTTP 経由のリクエストベースの認証に適しています。これを Authorization ヘッダーで渡すと、受信サービスが署名を検証します。
JWT SVID は X.509 SVID よりも有効期間が短く、サービス間でのトークンの再利用を防ぐために特定の対象ユーザーに範囲が限定されています。
Cilium の相互認証には、X.509 SVID が使用されます。この記事の残りの部分では、X.509 に焦点を当てます。
サービス A がサービス B の証明書を検証するには、どの認証局が署名したかをサービス A が知る必要があります。 SPIFFE では、これは信頼バンドルと呼ばれます。つまり、信頼ドメイン内で信頼される CA 証明書のセットです。
SPIRE は、ワークロード API 経由でトラスト バンドルを利用できるようにします。ワークロードがその ID を要求すると、現在の信頼バンドルも受信します。 SPIRE サーバーが CA をローテーションすると、新しい信頼バンドルがすべてのエージェントに配布され、すべてのワークロードにプッシュされます。アプリケーションが信頼バンドルを手動で管理する必要はありません。
SPIRE は、SVID を発行し、アイデンティティ ライフサイクルを管理するエンジンです。そのアーキテクチャを理解することで、Cilium の統合が意味を成します。
SPIRE には 2 つの主要コンポーネントがあります。 SPIRE サーバーは中央 CA です。これは、ワークロード エントリ (どの SPIFFE ID をどのワークロードに発行する必要があるかを説明するレコード) のレジストリを維持します。ワークロードに代わってエージェントに SVID を発行し、信頼ドメイン全体の信頼のルートとなります。
SPIRE エージェントは、すべてのノードで DaemonSet として実行されます。それには 2 つの仕事があります。まず、SPIRE サーバーが正当なノードで実行されていることを証明します。これはノード構成証明と呼ばれます。第二に、それは暴露します

■ ノード上の Unix ソケット上の SPIFFE ワークロード API。ワークロードはこれを使用して SVID を要求します。
エージェントは SVID をローカルにキャッシュするので、SPIRE サーバーへの接続が一時的に失われてもワークロード ID がすぐに壊れることはありません。
この分割 (中央サーバー、ノードごとのエージェント) は意図的に行われたものです。ワークロードが SPIRE サーバーに直接接続することはありません。彼らは自分のノード上のエージェントとのみ通信します。エージェントはすべての ID リクエストを仲介し、ノードが侵害された場合の影響範囲を制限します。
SPIRE エージェントが新しいノードで起動するときは、ワークロードに ID を提供する前に、SPIRE サーバーに対して自身の ID を証明する必要があります。これはノードの構成証明です。
Kubernetes では、SPIRE はノード構成証明に PSAT (予測されたサービス アカウント トークン) を使用します。エージェントは、SPIRE サーバーの対象者向けに特別に投影される Kubernetes サービス アカウント トークンを提示します。 SPIRE サーバーは、Kubernetes API に接続してトークンを検証し、エージェントが予期されたサービス アカウントを使用して予期された名前空間で実行されていることを確認し、エージェントに独自の SVID を発行します。
これが、SPIRE が特定の Kubernetes API フラグを必要とする理由です。 kube-apiserver は、適切な対象ユーザーに対して投影されたサービス アカウント トークンをサポートするように構成する必要があります。そのため、以下のデモの種類クラスター構成では --api-audiences と --service-account-issuer が設定されています。
ノードが証明されると、そのエージェントはワークロードを証明できるようになります。ワークロードがワークロード API ソケットに接続して SVID をリクエストすると、エージェントは Kubernetes API にクエリを実行して、そのワークロードに関する情報 (Kubernetes 名前空間、サービス アカウント、ポッド名、ラベルなど) を収集します。これらの事実を SPIRE サーバーに登録されているワークロード エントリと照合します。一致するエントリが存在する場合、エージェントは対応する SVID を発行します。
ワークロード e

エントリは次のようになります。
SPIFFE ID: spiffe://example.org/ns/production/sa/checkout
親 ID: spiffe://example.org/spire/agent/k8s_psat/default/<node-uid>
セレクター:
k8s:ns:本番環境
k8s:sa:チェックアウト
セレクターは、一致する必要がある Kubernetes ファクトを記述します。サービス アカウント チェックアウトを使用して運用名前空間で実行されているポッドは、SPIFFE ID spiffe://example.org/ns/production/sa/checkout を受け取ります。他のポッドはそうではありません。
SVID は設計上、有効期間が短くなります。 SPIRE の X.509 SVID のデフォルト TTL は 1 時間です。 SPIRE エージェントはそれらをバックグラウンドで自動的にローテーションします。つまり、新しいキー ペアを生成し、サーバーから新しい SVID を要求し、古い SVID の有効期限が切れる前に新しい SVID をワークロード API で利用できるようにします。
ワークロード API を直接使用するワークロード、または SPIFFE CSI ドライバーなどのツールを使用するワークロードは、新しい SVID を透過的に取得します。
有効期間の短い認証情報はゼロトラストの方法です。ワークロードの SVID が侵害された場合、その有効期限は 1 時間のみです。これを、歴史的に永久に有効であった Kubernetes サービス アカウント トークンと比較してください。
Cilium が SPIFFE を使用して相互 TLS を実装する方法
サービス メッシュ mTLS への従来のアプローチ (Istio や Linkerd など) は、すべてのポッドにサイドカー プロキシを挿入します。プロキシはすべてのトラフィックをインターセプトし、TLS ハンドシェイクを処理します。アプリケーションは TLS が発生していることを知りません。サイドカーにより、メモリ オーバーヘッド (Envoy のポッドあたり約 50 ～ 100 MB)、すべてのリクエストでの追加のネットワーク ホップ、および複雑な証明書挿入メカニズムが追加されます。
シリウムは別の道を歩みます。プロキシを挿入するのではなく、eBPF を使用してネットワーク層で認証を処理します。各ノードで実行されている Cilium エージェントは接続をインターセプトし、SPIFFE SVID を使用して相互 TLS ハンドシェイクを実行し、認証結果を強制します。これらすべてがユーザー空間プロキシなしでカーネル内で行われます。
T

彼のメカニズムは次のように動作します。ポッド A がポッド B への接続を開始すると、ポッド A のノード上の Cilium エージェントが接続をインターセプトします。 SPIRE Workload API からポッド A の SVID を取得します。この接続に相互認証を必要とする CiliumNetworkPolicy があるかどうかを確認します。存在する場合、ポッド B のノード上の Cilium エージェントと TLS ハンドシェイクを実行し、ポッド A の SVID を提示して、ポッド B の SVID を要求します。
両方のエージェントは、SPIRE 信頼バンドルに対して SVID を検証します。両方の SVID が有効で、ポリシーが接続を許可する場合、処理は続行されます。いずれかの SVID が無効であるか欠落している場合、接続は切断されます。
ポッド A 上のアプリケーションは、ポッド B 上のアプリケーションからデータを受信します。どちらのアプリケーションも TLS コードを作成していません。どちらもサイドカーは付いていません。認証はすべて、それぞれのノード上の Cilium エージェント内で行われます。
Cilium のモデルでは、Cilium エージェント自体が SPIRE から SPIFFE ID を取得します。これは、ワークロードに代わって SVID を要求できる代理 ID として機能します。
これは、各ワークロードが独自の SVID を直接要求するスタンドアロン SPIRE モデルとは少し異なります。 Cilium オペレーターは、管理する Kubernetes ID に基づいて SPIRE にワークロード エントリーを自動的に登録するため、ユーザーが管理する必要はありません。

[切り捨てられた]

## Original Extract

Your network policy says: allow traffic from 10.0.1.45. Yesterday, 10.0.1.45 was your payment service. Today, after a rolling deployment, it's your logging agent. Your payment service is now at 10.0.1

How to Implement Zero-Trust Workload Identity in Kubernetes with SPIFFE, SPIRE, and Cilium
Your network policy says: allow traffic from 10.0.1.45 .
Yesterday, 10.0.1.45 was your payment service. Today, after a rolling deployment, it's your logging agent. Your payment service is now at 10.0.1.89 .
Kubernetes has already updated all the endpoints and service records — but your network policy has no idea. It silently allows traffic through based on an IP address that no longer belongs to the workload you intended to trust.
This is the workload identity problem. IP addresses aren't an identity, they're a location. And in a Kubernetes cluster, location changes constantly. Building security policy on top of IP addresses means your security posture silently degrades every time a pod is scheduled, rescheduled, or scaled.
The answer is cryptographic workload identity: every workload gets a certificate-backed identity that proves who it is, not where it is. Services authenticate each other using those certificates before exchanging any data. If the certificate doesn't match, the connection is refused, regardless of what IP address it came from.
This is what SPIFFE and SPIRE provide. And this is how Cilium enforces it using eBPF, without injecting a sidecar into every pod.
In this article you'll understand how the SPIFFE identity model works, deploy SPIRE to issue cryptographic identities to workloads, and use Cilium's built-in SPIRE integration to enforce mutual TLS between services without touching your application code.
Familiarity with Kubernetes RBAC and pod security — this handbook covers the foundations
Familiarity with TLS certificates and Kubernetes Secrets — this handbook covers cert-manager and certificate concepts
Helm 3 and the Cilium CLI installed
A kind cluster — you'll create a fresh one with Cilium as the CNI in this article
Patience: this is the most complex demo I've covered in this group of articles. SPIRE has more moving parts than anything else covered so far.
All demo files are in the companion GitHub repository .
SVIDs: The Cryptographic Identity Document
How Cilium Implements Mutual TLS with SPIFFE
Demo 1 — Install Cilium with SPIRE Integration
Step 1: Install the Cilium CLI
Step 2: Create a kind cluster without a default CNI
Step 3: Install Cilium with SPIRE enabled
Step 4: Verify the installation
Demo 2 — Enforce Mutual TLS with a CiliumNetworkPolicy
Step 1: Deploy a client and server
Step 2: Confirm traffic flows without authentication
Step 3: Apply a CiliumNetworkPolicy requiring mutual authentication
Step 4: Verify authenticated traffic still flows
Step 5: Observe the authentication with Hubble (optional)
Step 6: Verify that a pod without the matching label is blocked
Step 7: Check the workload entries in SPIRE
The opening scenario isn't theoretical. In Kubernetes, pods are ephemeral. The scheduler can place a pod on any node, and a pod's IP address is assigned at scheduling time from the node's IP pool.
When a pod is deleted and recreated through a rolling deployment, a node drain, or an autoscaler event, it gets a new IP address. If you've written a NetworkPolicy that says, "allow traffic from this IP", that policy is now pointing at nothing, or worse, at a different workload.
Kubernetes service names help here for east-west traffic — a Service name resolves consistently regardless of which pods back it. But a NetworkPolicy based on a Service name is still a label selector match, not a cryptographic assertion. Any pod that can spoof the right labels can bypass it.
What you actually want is this: before service A sends a request to service B, service B proves its identity cryptographically. If service B can't prove it is who it claims to be, service A refuses the connection. This is mutual TLS, and the key question is: where do the identities come from?
SPIFFE — Secure Production Identity Framework for Everyone — is a CNCF standard that defines a model for workload identity. It doesn't implement anything by itself. It specifies the format of identities, the API for requesting them, and the trust model that makes them verifiable across services, clusters, and clouds. SPIRE is the reference implementation of that specification.
A SPIFFE identity is a URI with a specific format:
spiffe://<trust-domain>/<workload-path>
The trust domain is a string that identifies the administrative boundary — typically your organisation, cluster, or environment. Everything within the same trust domain can verify each other's identities. Identities from different trust domains require explicit federation configuration.
spiffe://payments.corp/ns/production/sa/checkout
spiffe://analytics.corp/ns/data/sa/pipeline-worker
spiffe://cluster.local/ns/monitoring/sa/prometheus
The path after the trust domain is arbitrary — it's defined by your SPIRE configuration and typically encodes the Kubernetes namespace and service account of the workload.
SVIDs: The Cryptographic Identity Document
An SVID — SPIFFE Verifiable Identity Document — is how a SPIFFE identity is materialised into something a service can actually use.
An X.509 SVID is a standard TLS certificate where the SPIFFE ID is embedded in the Subject Alternative Name (SAN) URI field. Because it's a standard X.509 certificate, any TLS library can use it without modification.
The workload presents this certificate in a TLS handshake, and the peer verifies the certificate was signed by a trusted SPIRE server. This is the format used for long-lived connections like gRPC streams.
A JWT SVID is a signed JSON Web Token containing the SPIFFE ID as a claim. It's suitable for request-based authentication over HTTP — pass it in an Authorization header, and the receiving service verifies the signature.
JWT SVIDs are shorter-lived than X.509 SVIDs and scoped to a specific audience to prevent token reuse across services.
For Cilium's mutual authentication, X.509 SVIDs are used. The rest of this article focuses on X.509.
For service A to verify service B's certificate, service A needs to know which Certificate Authority signed it. In SPIFFE, this is called the trust bundle — the set of CA certificates that are trusted within a trust domain.
SPIRE makes the trust bundle available via the Workload API. When a workload requests its identity, it also receives the current trust bundle. When the SPIRE server rotates its CA, it distributes the new trust bundle to all agents, which push it to all workloads. Your application never has to manage trust bundles manually.
SPIRE is the engine that issues SVIDs and manages the identity lifecycle. Understanding its architecture is what makes the Cilium integration make sense.
SPIRE has two main components. The SPIRE Server is the central CA. It maintains a registry of workload entries (records that describe which SPIFFE IDs should be issued to which workloads). It issues SVIDs to agents on behalf of workloads, and it's the root of trust for the entire trust domain.
The SPIRE Agent runs on every node as a DaemonSet. It has two jobs. First, it proves to the SPIRE Server that it's running on a legitimate node. This is called node attestation. Second, it exposes the SPIFFE Workload API on a Unix socket on the node, which workloads use to request their SVIDs.
The agent caches SVIDs locally so that a temporary loss of connection to the SPIRE Server doesn't immediately break workload identity.
This split — central server, per-node agents — is deliberate. Workloads never contact the SPIRE Server directly. They only talk to the agent on their own node. The agent mediates all identity requests, which limits the blast radius if a node is compromised.
When a SPIRE Agent starts up on a new node, it needs to prove its own identity to the SPIRE Server before it can serve identities to workloads. This is node attestation.
In Kubernetes, SPIRE uses PSAT — Projected Service Account Tokens — for node attestation. The agent presents a Kubernetes service account token that is projected specifically for the SPIRE server's audience. The SPIRE Server contacts the Kubernetes API to verify the token, confirms the agent is running in the expected namespace with the expected service account, and issues the agent its own SVID.
This is the reason SPIRE requires specific Kubernetes API flags. The kube-apiserver must be configured to support projected service account tokens with the right audience, which is why the kind cluster config in the demo below sets --api-audiences and --service-account-issuer .
Once a node has been attested, its agent can attest workloads. When a workload connects to the Workload API socket and requests an SVID, the agent collects facts about that workload (like its Kubernetes namespace, service account, pod name, and labels) by querying the Kubernetes API. It matches those facts against the workload entries registered in the SPIRE Server. If a matching entry exists, the agent issues the corresponding SVID.
A workload entry looks like this:
SPIFFE ID: spiffe://example.org/ns/production/sa/checkout
Parent ID: spiffe://example.org/spire/agent/k8s_psat/default/<node-uid>
Selectors:
k8s:ns:production
k8s:sa:checkout
The selectors describe the Kubernetes facts that must match. A pod running in the production namespace with service account checkout will receive the SPIFFE ID spiffe://example.org/ns/production/sa/checkout . Any other pod will not.
SVIDs are short-lived by design. The default TTL for X.509 SVIDs in SPIRE is one hour. The SPIRE Agent automatically rotates them in the background — generating a new key pair, requesting a fresh SVID from the server, and making the new SVID available on the Workload API before the old one expires.
Workloads that use the Workload API directly or tools like the SPIFFE CSI driver get the new SVID transparently.
Short-lived credentials are the zero-trust way. If a workload's SVID is compromised, it's only valid for an hour. Compare that to a Kubernetes service account token, which was historically valid forever.
How Cilium Implements Mutual TLS with SPIFFE
Traditional approaches to service mesh mTLS (like Istio or Linkerd) inject a sidecar proxy into every pod. The proxy intercepts all traffic and handles the TLS handshake. The application has no idea TLS is happening. The sidecar adds memory overhead (roughly 50–100MB per pod for Envoy), an extra network hop on every request, and a complex certificate injection mechanism.
Cilium takes a different path. Rather than injecting a proxy, it handles authentication at the network layer using eBPF. The Cilium agent running on each node intercepts connections, performs the mutual TLS handshake using SPIFFE SVIDs, and enforces the authentication result — all in the kernel, without any user-space proxy.
The mechanism works like this. When pod A initiates a connection to pod B, the Cilium agent on pod A's node intercepts the connection. It retrieves pod A's SVID from the SPIRE Workload API. It checks whether there's a CiliumNetworkPolicy requiring mutual authentication for this connection. If there is, it performs a TLS handshake with the Cilium agent on pod B's node, presenting pod A's SVID and requesting pod B's SVID in return.
Both agents verify the SVID against the SPIRE trust bundle. If both SVIDs are valid and the policy allows the connection, it proceeds. If either SVID is invalid or missing, the connection is dropped.
The application on pod A receives data from the application on pod B. Neither application wrote any TLS code. Neither has a sidecar. The authentication happened entirely in the Cilium agents on their respective nodes.
In Cilium's model, the Cilium agent itself gets a SPIFFE identity from SPIRE. It acts as a delegate identity that can request SVIDs on behalf of workloads.
This is slightly different from the standalone SPIRE model where each workload requests its own SVID directly. The Cilium operator registers workload entries in SPIRE automatically based on the Kubernetes Identities it manages, so you don't need to ma

[truncated]

---
source: "https://blog.calif.io/p/privilege-escalation-in-eks"
hn_url: "https://news.ycombinator.com/item?id=49129803"
title: "Privilege Escalation in AWS Elastic Kubernetes Service (2023)"
article_title: "Privilege escalation in AWS Elastic Kubernetes Service (EKS)"
author: "882542F3884314B"
captured_at: "2026-07-31T23:53:50Z"
capture_tool: "hn-digest"
hn_id: 49129803
score: 1
comments: 0
posted_at: "2026-07-31T23:53:04Z"
tags:
  - hacker-news
  - translated
---

# Privilege Escalation in AWS Elastic Kubernetes Service (2023)

- HN: [49129803](https://news.ycombinator.com/item?id=49129803)
- Source: [blog.calif.io](https://blog.calif.io/p/privilege-escalation-in-eks)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T23:53:04Z

## Translation

タイトル: AWS Elastic Kubernetes Service での権限昇格 (2023)
記事のタイトル: AWS Elastic Kubernetes Service (EKS) での権限昇格
説明: チームは最近、AWS Elastic Kubernetes Service (EKS) で侵害されたポッドから権限を昇格しようとして、すべての EKS バージョンでデフォルトで有効になっているセキュリティ メカニズムである NodeRestriction に苦戦するという興味深いシナリオに遭遇しました。

記事本文:
AWS Elastic Kubernetes Service (EKS) での権限昇格
カリフォルニア
購読 サインイン AWS Elastic Kubernetes Service での権限昇格
An Trinh 氏と Duc Nguyen 氏 2023 年 4 月 2 日 7 共有 チームは最近、AWS Elastic Kubernetes Service (EKS) で侵害されたポッドから権限を昇格しようとして、すべての EKS バージョンでデフォルトで有効になっているセキュリティ メカニズムである NodeRestriction に苦労するという興味深いシナリオに遭遇しました。
侵害された AWS EC2 マシン上で、AWS メタデータ サービスにリクエストしてマシンの IAM トークンを取得できることはよく知られています。
私たちにとって驚いたのは、そのマシン内で実行されているコンテナーからもそれを実行できることです。これを阻止するデフォルトのメカニズムはありません (IMDSv2 はこれを解決しようとします。これについては後で詳しく説明します)。これは、EKS ポッドがどれほど低い特権であっても、基盤となる EC2 マシンと同じ AWS 特権を持っていることを意味します。これには、EKS API サーバーのエンドポイントがデフォルトでインターネットからアクセスできるという事実が伴います。
Christophe は、このことの一部を彼の研究 1 で詳細に文書化しており、そこで EC2 に付与されているデフォルトの EC2 DeleteNetworkInterface IAM 権限も利用しました。しかし、私たちはより深刻な影響があるかどうかを確認したいと考え、EKS への攻撃経路を調査しました。
デフォルトの EKS デプロイメントでは、ノードの IAM トークンには、クラスター上で関連する system:node ロールがあります。 EKS は、aws-iam-authenticator プラグインを介して AWS ID と EKS 権限の間のマッピングを提供することでこれを容易にします (詳細は後述)。
ただし、 system:node トークンは、Kubernetes NodeRestriction によって厳しく制限されています。ポッド定義などのいくつかの基本的なクラスター情報は利用可能ですが、それほど多くはありません。
ドキュメントをさらに読んだところ、次の行に注目しました。
そんなク

belet は、自分自身の Node API オブジェクトの変更と、自分のノードにバインドされている Pod API オブジェクトの変更のみが許可されます。
つまり、ノード トークンを使用すると、ノード内の任意のポッドになりすますことができます。
ノードトークンからクラスター管理者へ
私たちのアプローチは、system:node トークンと同じノード上で実行されているポッドを調査することでした。いくつかの行き詰まりを経て、API サーバーからこれらのポッドのサービス アカウント トークンをリクエストすることで、ポッドになりすましてその権限を使用できるようにすることにしました。
K8 では、ほとんどの人はどのポッドがどのノードにデプロイされるかを気にしません。つまり、ノードには機密性の高いポッドと信頼できないポッドの両方が含まれる可能性があります。すべてのポッドを検査してピボットすることで、クラスターのシークレットをリストする権限を持つトークンなど、より高い信頼境界を持つトークンを取得できる可能性があります。
その前提に基づいて、テスト設定をすぐに開始しました。
1 つのノードと ns1 という名前空間を持つ、cluster-1 という名前の EKS クラスター。この記事の執筆時点では、デフォルトの EKS バージョンは v1.25.7-eks-a59e1f0 でした。
ClusterRole クラスター管理者 (クラスター内の任意のリソースに対して任意のアクションを実行できる組み込みロール) にマップされたサービス アカウント sa-priv
最低限のポッド権限を持つ pod1、alpine を実行
sa-priv トークンに関連付けられた pod-priv (同じく alpine を実行)
流れはとても簡単です。実際、あまりにも単純すぎて、誰かが以前にこの動作を文書化していないかどうかを何度も確認する必要がありました。
pod1 から EC2 メタデータ IAM アクセス トークンをリクエストし、次のようにして EKS トークンに交換します。
$ aws eks get-token --cluster-name=cluster-1 トークンには system:node ロールがあり、次のようになります。
$ kubectl config set-credentials usr1 --token="$(aws eks get-token --cluster-name=cluster-1 | jq -rc '.status.token')"
ユーザー「usr1」を設定します。
$ kubectl config set-context nod

e1 --cluster=cluster-1.us-west-2.eksctl.io --user=usr1
コンテキスト「node1」が変更されました。
$ kubectl --context=node1 認証 can-i --list
警告: リストは不完全である可能性があります: ノードオーソライザーはユーザールールの解決をサポートしていません
リソース 非リソース URL リソース名 動詞
selfsubjectaccessreviews.authorization.k8s.io [] [] [作成]
selfsubjectrulesreviews.authorization.k8s.io [] [] [作成]
証明書署名要求.certificates.k8s.io/selfnodeserver [] [] [作成]
[/api/*] [] [取得]
...省略 この時点では、トークンの権限はまだ NodeRestriction によって厳しく制限されています。
次に、 pod-priv のトークンをリクエストします。この時点でエラーが発生し、しばらくがっかりしました
$ kubectl --context=node1 トークンの作成 -n ns1 sa-priv
エラー : トークンの作成に失敗しました: serviceaccounts "sa-priv" は禁止されています: ノードが要求したトークンはポッドにバインドされていません よく見ると、これは認証エラーではなく、単なるフォーマット エラーのように見えます。案の定、このコマンドにはいくつかの追加パラメータ (bound-object-kind、bound-object-name、bound-object-uid) が必要です。これらはすべてポッド定義から簡単に取得できます。
少しいじった後、期待どおりに動作し、 pod-priv のトークンを取得しました。
$ kubectl --context=node1 トークンの作成 -n ns1 sa-priv \
--bound-object-kind=ポッド \
--bound-object-name=pod-priv \
--bound-object-uid=7f7e741a-12f5-4148-91b4-4bc94f75998d そして最後にサービス アカウント sa-priv にピボットします。この後、クラスターに対する完全なアクセス許可がどのように与えられるかに注目してください。
ここでの sa-priv トークンは一例であり、cluster-admin ロールを持っている必要はありません。現実の世界では、クラスターのシークレットをリストする権限を持つものである可能性があります。いずれにせよ、自由に使えるものがたくさんあるため、cluster-admin へのピボットが簡単になります。私たちの場合、c になるのは時間の問題でした。

クラスターを侵害します。
… EKS クラスターのバックドア化へ
クラスターを乗っ取った後、私たちはそこにバックドアを仕掛ける方法を探しました。
前述したように、EKS は、aws-iam-authenticator プラグインを介して AWS ID と EKS 権限の間のマッピングを容易にします。これは、Kubernetes API サーバーの認証フローにフックすることで機能し、AWS EKS トークン (実際には AWS IAM トークンの単なる薄いラッパー) を使用して https://sts.amazonaws.com にある AWS STS エンドポイントを呼び出します。
IAM トークンを使用して STS を呼び出すと、呼び出し元のロール ARN が生成され、呼び出し元を識別します。次に、プラグインは Kubernetes aws-auth configmap のマッピングを使用して、Kubernetes ロールと AWS ロール ARN の間をマッピングします。
これを念頭に置いて、バックドアのアイデアは、system:masters グループにマップされた aws-auth データストアにユーザーを制御下に置くことです。また、aws-auth で指定されたロール ARN は、EKS クラスターを所有するアカウント所有者と同じアカウント所有者に属する必要はないため、新しいユーザーをサインアップして、そのルート ロール ARN をそこに植えるだけで済みます。
以下に示すように、外部ユーザーのトークンを使用して新しいレコードを aws-auth configmap に挿入し、その構成が存在する限りそのレコードに対する管理者権限を保持します。
aws-auth への変更の前後で Kubernetes のアクセス許可がどのように異なるかに注目してください。
さらに、この攻撃ベクトルは AWS IAM ロールに依存しません。つまり、EC2 の IAM ロールからすべての権限が取り消されても機能し続けるということです。これは、aws-iam-authenticator が呼び出し元のロール ARN (AWS ID) のみを検証でき、トークンの現在の実際の権限ではなく、aws-auth configmap を介してマッピングを提供するためです。
私たちの見方では、EKS には根本原因となる 2 つの問題があると考えられます。
ポッドはノードのメタデータを取得することができます。
EC2のIAMアクセストークン

は EKS system:node トークンと交換できます。
これに対処するのは簡単ではなく、多くの計画が必要になります。現実世界での悪用の可能性を考慮して、これを優先することをお勧めします。
最初の問題を防ぐために、AWS IMDSv2 機能は、公式のセキュリティ強化に関する推奨事項に従って http_put_response_hop_limit を 1 に設定することで、ポッドがノードのメタデータに到達するのをブロックできます。すべてのノードでカスタム iptables ルールを使用することも代替手段として機能する可能性があります。 Christophe は彼の研究の中で、これを促進するためのいくつかの方法についても詳しく説明しました。ただし、この戦略を追求すると、 Terraform など、ポッドからインスタンスのメタデータへの通信を必要とするいくつかの重要なサービスでは問題が発生する可能性があります。
EC2 IAM ロールに EKS system:node ロールが付与されているというもう 1 つの問題については、AWS からの将来の更新 (存在する場合) を待つ以外に、それを修正する実用的な方法がまだ見つかりません。
最後に、多層防御として、従来のネットワーク アーキテクチャにおける DMZ の概念と同様に、クラスター内の信頼できないサービスを隔離されたノードに移動することが賢明です。 Kubernetes は、ポッドを特定のノードに割り当て、このタイプのネットワーキングを容易にするさまざまな方法を提供します。その間に、クラスターの API エンドポイントへのデフォルトのインターネット アクセスを制限またはオフにします。
2023 年 11 月 27 日更新:
証明書の管理に cert-manager を使用する EKS クラスターで、簡単な権限昇格ベクトルが見つかりました。
1/ 上で説明したのと同じ攻撃ベクトルを使用して、cert-manager のサービス アカウント トークンを取得します。
cert-manager サービス アカウントには完全なポッド権限があることがわかりました。
2/ cert-manager サービス アカウントを使用して新しいポッドを作成し、特権サービス アカウントをそのポッドにアタッチします。
たとえば、clusterrole-aggregation をアタッチできます。

-コントローラー サービス アカウント。 K8s のデフォルト ポリシーによれば、clusterrole-aggregation-controller にはクラスター ロールに対する完全なアクセス許可があり、クラスター ロールを任意の方法で変更できます。
3/ クラスターの役割にパッチを適用して、クラスター管理者になる次の追加の権限を追加します。私たちの知る限り、このテクニックは Raesene 2 によって最初に議論されました。
- APIグループ:
- '*'
リソース:
- '*'
動詞:
- '*' cert-manager は、HTTP-01 チャレンジの検証中に新しいポッドを作成するため、ポッド リソースに対する完全な権限が必要です。
要約すると、このいくつかの権限昇格手順の連鎖により、迅速かつ密かにクラスター全体が侵害される可能性があります。
この攻撃パスを軽減するには、cert-manager に付与されているポッド作成権限を取り消し、DNS を使用したドメイン検証に切り替えることをお勧めします。
https://blog.christophetd.fr/privilege-escalation-in-aws-elastic-kubernetes-service-eks-by-compromising-the-instance-role-of-worker-nodes/
https://raesene.github.io/blog/2020/12/12/Escalating_Away/
7 共有 この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

The team recently encountered an interesting scenario where we were trying to escalate privileges from a compromised pod in AWS Elastic Kubernetes Service (EKS) and struggled with NodeRestriction, a security mechanism enabled by default on all EKS versions

Privilege escalation in AWS Elastic Kubernetes Service (EKS)
Calif
Subscribe Sign in Privilege escalation in AWS Elastic Kubernetes Service
An Trinh and Duc Nguyen Apr 02, 2023 7 Share The team recently encountered an interesting scenario where we were trying to escalate privileges from a compromised pod in AWS Elastic Kubernetes Service (EKS) and struggled with NodeRestriction , a security mechanism enabled by default on all EKS versions .
It’s well-known that on a compromised AWS EC2 machine, one can request AWS metadata service to obtain the machine’s IAM token.
What’s surprising to us was that it’s possible to do that from a container running inside that machine as well. There’s no default mechanism to stop that from happening (IMDSv2 tries to solve that, which will be detailed later). This means an EKS pod, no matter how low-privileged it is, has the same AWS privileges of the underlying EC2 machine. This is accompanied by the fact that EKS API server’s endpoint is by default accessible from the internet.
Christophe documented some of this in detail in his research 1 , where he also exploited the default EC2 DeleteNetworkInterface IAM permission granted to EC2. But we wanted to see if there’s a more severe impact and explored the attack path towards EKS.
On a default EKS deployment, the node’s IAM token has the associated system:node role on the cluster. EKS facilitates this by providing mapping between AWS identities and EKS privileges via aws-iam-authenticator plugin (details to follow).
The system:node token however is severely limited by Kubernetes NodeRestriction . Some basic cluster information, such as pods definition, is available, but that’s not much.
Upon further reading the document, this line caught our attention:
Such kubelets will only be allowed to modify their own Node API object, and only modify Pod API objects that are bound to their node.
This means with the node token, we can impersonate any pods in the node.
From node token to cluster-admin
Our approach was to explore the pods running on the same node as our system:node token. After hitting a few dead ends, we came up with requesting service account tokens for those pods from the API server, allowing us to impersonate them and use their privileges.
In K8s, most people don’t care which pod is deployed to which exact node, meaning a node can contain both sensitive pods and untrusted pods. By inspecting and pivoting through every pod, it’s likely we could obtain a token with higher trust-boundary, such as one with permissions to list the cluster’s secrets.
With that assumption, we quickly spun up a test setup:
EKS cluster named cluster-1 with one node and a namespace named ns1 . As of the time of writing, the default EKS version was v1.25.7-eks-a59e1f0.
Service account sa-priv mapped with the ClusterRole cluster-admin (a built-in role that can take any actions on any resources in the cluster)
pod1 with minimum pod privileges, running alpine
pod-priv associated with the sa-priv token, also running alpine
The flow is pretty straightforward. In fact it’s too straightforward that we had to check multiple times whether someone had documented this behavior before.
From pod1, request the EC2 metadata IAM access token, then exchange it to EKS token with:
$ aws eks get-token --cluster-name=cluster-1 The token should have the system:node role and look like following:
$ kubectl config set-credentials usr1 --token="$(aws eks get-token --cluster-name=cluster-1 | jq -rc '.status.token')"
User "usr1" set.
$ kubectl config set-context node1 --cluster=cluster-1.us-west-2.eksctl.io --user=usr1
Context "node1" modified.
$ kubectl --context=node1 auth can-i --list
Warning: the list may be incomplete: node authorizer does not support user rule resolution
Resources Non-Resource URLs Resource Names Verbs
selfsubjectaccessreviews.authorization.k8s.io [] [] [create]
selfsubjectrulesreviews.authorization.k8s.io [] [] [create]
certificatesigningrequests.certificates.k8s.io/selfnodeserver [] [] [create]
[/api/*] [] [get]
...snipped At this point the token’s permissions are still heavily limited by NodeRestriction.
Then request a token for pod-priv . There’s an error at this point, which disappointed us for a moment
$ kubectl --context=node1 create token -n ns1 sa-priv
error : failed to create token: serviceaccounts "sa-priv" is forbidden: node requested token not bound to a pod On a closer look, that does not look like an authorization error, but just a format error. Sure enough, the command requires some additional parameters : bound-object-kind, bound-object-name, bound-object-uid, all of which can be obtained easily enough from the pods definition.
After a bit of tinkering it works as expected, and we have the token for pod-priv :
$ kubectl --context=node1 create token -n ns1 sa-priv \
--bound-object-kind=Pod \
--bound-object-name=pod-priv \
--bound-object-uid=7f7e741a-12f5-4148-91b4-4bc94f75998d And finally pivot to the service account sa-priv . Notice how we have full permissions on the cluster after this.
The sa-priv token here is an example and does not have to have the cluster-admin role. In the real world, it could be one with permissions to list the cluster’s secrets. Either way, pivoting to cluster-admin is easier now that we have a lot at disposal. In our case, it was just a matter of time until compromising the cluster.
… to backdooring the EKS cluster
After taking over the cluster, we looked for ways to plant a backdoor there.
As mentioned previously, EKS facilitates mapping between AWS identity and EKS privileges via aws-iam-authenticator plugin. It works by hooking into Kubernetes API server’s authorization flow and calls AWS STS endpoint at https://sts.amazonaws.com with the AWS EKS token (which in fact is just a thin wrapper around AWS IAM token).
Calling STS with an IAM token would yield the caller’s role ARN, thereby identifying them. The plugin then uses the mapping in Kubernetes aws-auth configmap to map between Kubernetes roles and AWS role ARNs.
With that in mind, the idea for the backdoor is to plant a user under our control into the aws-auth datastore mapped with the system:masters group. The role ARN specified in aws-auth also does not have to belong to the same account owner as the one owning the EKS cluster, so we can just sign up a new user and plant its root role ARN there.
As demonstrated in the following, we insert a new record into aws-auth configmap with an external user’s token and retain the admin privileges on it for as long as that config still exists.
Notice how the Kubernetes permissions differ before and after the changes to aws-auth.
Additionally, this attack vector does not rely on AWS IAM role. Meaning it still works even if the EC2’s IAM role is revoked of all permissions . This is because aws-iam-authenticator can only verify the role ARN (the AWS identity) of the caller, and provides mapping via the aws-auth configmap, not the token’s current actual permissions.
The way we see it, there are two problems in EKS that play the root causes:
Pods are allowed to fetch their node’s metadata, and
EC2’s IAM access tokens can be exchanged for an EKS system:node token.
Addressing this is not straightforward and would need a lot of planning. We still recommend making this a priority due to how feasible it is to exploit in the real world.
To prevent the first problem, AWS IMDSv2 feature can block pods from reaching their node’s metadata by setting the http_put_response_hop_limit to 1, following the official security hardening recommendations . Using a custom iptables rule on all nodes could also work as the alternative. Christophe in his research also detailed several ways to facilitate this. Pursuing this strategy however, could be problematic for several important services that require communications from pod to instance metadata, like Terraform .
As for the other problem of EC2 IAM role being granted EKS system:node role, we have yet to find a practical way of remediating it other than waiting for a future update, if any, from AWS.
Finally as defense-in-depth, it’s wise to move untrusted services in the cluster to an isolated node, resembling the concept of DMZ in traditional networking architecture. Kubernetes provides many ways to assign pods to specific nodes and facilitate this type of networking. And while we’re at that, limit or turn off the default internet access to the cluster’s API endpoint.
Update Nov 27, 2023 :
We found a straightforward privilege escalation vector in EKS clusters that use cert-manager for managing certificates.
1/ Use the same attack vector explained above to obtain the service account token for cert-manager.
We found that the cert-manager service account has full pod permissions .
2/ Use the cert-manager service account to create a new pod, and attach privileged service accounts into that pod.
For example, we can attach the clusterrole-aggregation-controller service account . According to the default policy in K8s, clusterrole-aggregation-controller has full permissions on cluster roles to allow it to mutate them in any way.
3/ Patch any cluster roles to add the following additional permissions to become cluster-admin. To be best of our knowledge, this technique was first discussed by Raesene 2 .
- apiGroups:
- '*'
resources:
- '*'
verbs:
- '*' cert-manager need full permissions on pod resources because it creates a new pod during HTTP-01 Challenge Validation :
In summary, this chaining of several privilege escalation steps can quickly and stealthily result in full cluster compromise.
To mitigate this attack path, we recommend revoking pod creation permission granted to cert-manager, switching to domain verification using DNS.
https://blog.christophetd.fr/privilege-escalation-in-aws-elastic-kubernetes-service-eks-by-compromising-the-instance-role-of-worker-nodes/
https://raesene.github.io/blog/2020/12/12/Escalating_Away/
7 Share Discussion about this post Comments Restacks Top Latest Discussions No posts

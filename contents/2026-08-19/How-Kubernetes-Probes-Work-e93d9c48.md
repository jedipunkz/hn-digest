---
source: "https://ngrok.com/blog/probes"
hn_url: "https://news.ycombinator.com/item?id=49363665"
title: "How Kubernetes Probes Work"
article_title: "How Kubernetes probes work | ngrok blog"
image: "https://ngrok.com/blog-assets/images/2026-08-19-probes/og.jpg"
author: "cyndunlop"
captured_at: "2026-08-19T17:19:46Z"
capture_tool: "hn-digest"
hn_id: 49363665
score: 6
comments: 0
posted_at: "2026-08-19T16:25:37Z"
tags:
  - hacker-news
  - translated
---

# How Kubernetes Probes Work

- HN: [49363665](https://news.ycombinator.com/item?id=49363665)
- Source: [ngrok.com](https://ngrok.com/blog/probes)
- Score: 6
- Comments: 0
- Posted: 2026-08-19T16:25:37Z

## Translation

タイトル: Kubernetes プローブの仕組み
記事のタイトル: Kubernetes プローブの仕組み |ングロクのブログ
説明: ブラウザーで実行されているシミュレートされた Kubernetes クラスター上でプローブがどのように対話的に動作するかを学びます。

記事本文:
Kubernetes プローブの仕組み | ngrok ブログ メイン コンテンツにスキップ ngrok ホーム / ngrok ブログ ホーム ブログ モバイル ナビゲーションを開く 製品 解決する問題 リソース ドキュメント 価格 ログイン サインアップ 製品ゲートウェイ
ビルドとテスト ローカル アプリを共有する
導入と実行 トラフィックを自己ホスト型モデルにルーティングする
配信と接続 API の配信と保護
ネットワークを介してプライベートに接続する
LLM を 4 倍小さく、2 倍高速化
検索… Control ⌃ K ニュースレター RSS 2026 年 8 月 19 日
Kubernetes でプローブがどのように機能するかを実際に説明します。アプリケーションの回復力をどのように高めることができるか、避けられる間違いを防ぐのにどのように役立つか。回復に何時間もかかる再起動ループや、ロールアウト中のリクエストのドロップなどです。
この投稿のすべてのインタラクティブなデモは、Kubernetes の TypeScript への部分移植である webernetes を使用しています。これには、ブラウザ上でシミュレートされたクラスターを実行するための、移植された 100,000 行を超える Kubernetes Go コードが含まれています。これらのデモの動作を k3s に対して検証したところ、Kubernetes のバグを見つけることができました。それについては後で詳しく説明します。
3 種類のプローブとその用途。
それらを構成して組み合わせる方法。
よくある構成ミスがどのように失敗するか。
プローブが展開速度に与える影響。
このセクションをブックマークする プローブのないポッド
単一のコンテナーでポッドを実行したいと考えています。そのマニフェスト pod-a.yaml は次のとおりです。
1 apiVersion : "v1" 2 kind : "Pod" 3 メタデータ : ⋯ 4 name : "pod-a" 5 spec : ⋯ 6 コンテナー : ⋯ 7 - name : "app" ⋯ 8 image : "my-app:latest"
このイメージ my-app:latest は、ポート 8080 でリッスンする前に初期化に数秒かかります。再起動をクリックしてコンテナにシグナルを送信すると、コンテナがクラッシュし、Kubernetes によってバックアップが開始されると、これが以下に表示されます。デモはいつでも一時停止またはリセットできます。
クラスタ 0 / 2 を一時停止 コンテナを再起動

まだ完成していません。
コンテナを再起動します
最初のクラッシュの後、コンテナーはすぐに再起動します。 2 回目の後、Kubernetes は、再度開始する前に CrashLoopBackOff を適用します。デフォルトでは、この遅延は 10 秒で、クラッシュするたびに 2 倍になり、最大待機時間は 5 分になります。このデモでは 3 秒に短縮しました。
どちらの場合も、Kubernetes は、コンテナが開始されるとすぐに、準備が完了していないことがわかっていても、コンテナが準備完了であるとみなします。まだ起動作業を行っており、ポート 8080 をリッスンしていません。
次に、2 秒ごとに pod-a にリクエストを送信する pod-b を追加します。この投稿全体を通じて、pod-b はクライアント トラフィックのソース (イングレス コントローラー、ロード バランサー、サービス間リクエストなど) と考えることができます。
リクエストの進行中に以下のデモで pod-a を再起動すると、そのリクエストは失敗します。
クラスターを一時停止する リクエストを失敗させる まだ完了していません。
コンテナを再起動します
コンテナーを再起動した瞬間からその起動作業が完了するまで、コンテナーが準備完了とみなされている場合でも、リクエストは失敗します。これは私が望んでいることではありません。 pod-a がトラフィックを受信できる状態になったことを Kubernetes に知らせる必要があります。
このために、Kubernetes はプローブを提供します。プローブは、コンテナーの健全性を判断するためにコンテナーに送信される定期的なチェックです。 3 種類の味があります。
起動プローブは、コンテナ内のアプリケーションが起動したかどうかを判断します。
Readiness プローブは、アプリケーションがトラフィックを受信する準備ができているかどうかを判断します。
Liveness プローブは、アプリケーションを再起動する必要があるかどうかを判断します。
スタートアップ プローブは、上記のデモで示した問題に最適なようですので、そこから始めましょう。
このセクションをブックマークする
以下では、起動プローブを pod-a.yaml に追加しました。
1 apiVersion : "v1" 2 kind : "Pod" 3 メタデータ : ⋯ 4 name : "pod-a" 5 spec : ⋯ 6 コンテナ : ⋯

7 - 名前 : "アプリ" ⋯ 8 イメージ : "my-app:latest" 9 スタートアッププローブ : ⋯ 10 httpGet : ⋯ 11 パス : "/startup" 12 ポート : 8080 13 periodSeconds : 1 14 FailureThreshold : 5
これは、ポート 8080 上のポッドに GET /startup リクエストを送信する httpGet プローブです。ステータス コード 200 ～ 399 は成功としてカウントされます。これは periodSeconds 秒ごとに発生し、Kubernetes がコンテナーを強制終了する前に、failureThreshold に連続して失敗することが許可されます。これにより、コンテナーの起動作業が完了するまでに最大 5 秒かかります。
プローブは、 kubelet と呼ばれるプロセスによって送信されます。クラスター内の各ノードには独自の kubelet があり、各ノードで適切なポッドが実行され、プローブされていることを確認するのが kubelet の仕事です。
以下の pod-a を再起動すると、 NotReady と表示されます。 Kubernetes は、pod-a がまだ初期化されていないことを認識します。最初の起動プローブが成功した後にのみ Ready になります。
クラスター 0 / 2 を一時停止します。 コンテナーを再起動します。 まだ完了していません。
コンテナを再起動します
NotReady は、起動プローブを持つコンテナーを含むポッドのデフォルトです。ただし、
準備ができていない場合でも、pod-b は引き続きリクエストを送信します。
pod-a とそれらのリクエストは、コンテナーの起動中に引き続き失敗します。これ
これは、リクエストを pod-a の IP アドレスに直接送信するように pod-b を構成したためです。
準備メカニズムをバイパスします。
NotReadyについては少し嘘をついています
これらの失敗したリクエストを修正するには、より本番グレードのセットアップ、つまり pod-a の複数のコピーとそれらの間でリクエストの負荷分散を行う必要があります。 pod-a の 2 つのレプリカを実行するように構成された ReplicaSet と、それらの間で負荷分散を行うサービスを作成します。
1 apiVersion : "apps/v1" 2 kind : "ReplicaSet" 3 metadata : ⋯ 4 name : "replica-set-a" 5 spec : ⋯ 6 # `template` で定義されたポッドのコピーを 2 つ実行します。 7 レプリカ : 2 8 セレクター : ⋯ 9 matchLabels : ⋯ 1

0 # このラベルが付いたポッドはこのレプリカ セットの一部であるとみなします。 11 app : "pod-a" 12 template : ⋯ 13 metadata : ⋯ 14 label : ⋯ 15 app : "pod-a" 16 spec : ⋯ 17 # 以前と同じポッド仕様。 18 コンテナー : ⋯ 19 - 名前 : "app" ⋯ 20 イメージ : "my-app:latest" 21 startProbe : ⋯ 22 httpGet : ⋯ 23 パス : "/startup" 24 ポート : 8080 25 periodSeconds : 1 26 FailureThreshold : 5
サービス-a.yaml
1 apiVersion : "v1" 2 kind : "Service" 3 metadata : ⋯ 4 name : "service-a" 5 spec : ⋯ 6 selector : ⋯ 7 # このラベルを持つポッド間の負荷分散。 8 app : "pod-a" 9 ports : ⋯ 10 # ポッド上のこのポートにリクエストを送信します。 11 - ポート: 80 ⋯ 12 ターゲットポート: 8080
pod-b は今後、Kubernetes がサービス用に作成した DNS 名にリクエストを送信します。この場合は
個々のポッドに直接ではなく、 service-a.default.svc.cluster.local 。 Kubernetes はポッドの Ready を使用します
サービス負荷分散に含めるか除外するかの条件。
以下の再起動ボタンをクリックすると、一番上のコンテナのみをクラッシュさせることができます。上部のコンテナが起動しているとき、リクエストは常に下部のコンテナに送信されることに注意してください。コンテナーが NotReady の場合、ポッド全体が準備ができていないというマークが付けられ、コンテナーが属するどのサービスからもトラフィックを取得できなくなります。
クラスター 0 / 2 を一時停止します。 最上位のコンテナーを再起動します。 まだ完了していません。
一番上のコンテナを再起動します
それにもかかわらず、最上位コンテナを再起動したときにリクエストが処理中の場合、リクエストは失敗する可能性があります。これは、再起動ボタンによってコンテナが突然クラッシュするために発生します。実行中のリクエストを完了する機会がありません。
ここで行うより良い方法は、ポッドを削除し、ReplicaSet を利用して新しいポッドを起動することです。これは 2 つの理由により優れています。
Kubernetes では、デフォルトでポッドに 30 秒の終了猶予期間が与えられますが、ここでは 2 秒に設定しました。

この投稿をご覧ください。待つ必要はありません。削除されると、ポッドは終了したとみなされ、Kubernetes はポッドが含まれているすべてのサービスからポッドを削除します。新しいリクエストは受け付けられません。
ReplicaSet は終了するポッドをアクティブなレプリカとしてカウントしないため、削除されたポッドが終了するとすぐに代替を作成します。
正常な終了と起動プローブを組み合わせることで、起動中または停止中のコンテナーからリクエストが遠ざけられます。この次のデモでは、[削除] をクリックしても pod-b からのリクエストは失敗しません。
クラスター 0 / 2 を一時停止します。コンテナーの準備ができるまで待ちます。まだ完了していません。
一番上のポッドを削除する
新しいリクエストに対応できるポッドが常に存在するため、ユーザー トラフィックを中断することなくポッドを安全に削除できます。
この猶予期間は実際にはどのように機能するのでしょうか?
ポッドを削除すると、kubelet は最初に SIGTERM シグナルをコンテナーに送信し、次に
terminationGracePeriodSeconds 後、コンテナーがまだ実行中の場合は SIGKILL を送信します。 Kubernetes のドキュメント
ポッドのライフサイクルには、悲惨な詳細がすべて含まれています。
この投稿では、コンテナは SIGTERM を受信して​​から 1 秒後に終了するように構成されています。知っているからこれを選びました
処理中のリクエストを完了するには十分な長さです。独自のポッドでこれを構成する場合は、必ず次のことを行ってください。
最も長時間実行されているリクエストが完了するまで十分な時間を残しておきます。
このセクションをブックマークする 起動プローブを誤って構成する方法
前に、failureThreshold を 5 に設定し、periodSeconds を 1 に設定することで、ポッドの起動作業を完了するまでに約 5 秒の時間を与えていると述べました。独自のコンテナーでこれらの値を慎重に選択してください。時間が短すぎると、コンテナーがクラッシュ ループを引き起こす可能性があります。
以下の FailureThreshold を設定すると、新しい値でコンテナーが再起動されます。 1 または 2 に設定して、何が起こるかを見てください。
クラスターを一時停止する ポッド-a クラッシュ ループを作成する まだ完了していません

e.
失敗しきい値 5
数回再起動すると、pod-a は CrashLoopBackOff になります。起動プローブはコンテナーに起動するのに十分な時間を与えないため、このデモは、failureThreshold を 3 以上に設定し直すまでクラッシュ ループします。独自のコンテナーに対してこれを構成する場合は、最悪の場合の起動時間を考慮した値を選択してください。
このセクションをブックマークする Readiness プローブ
起動プローブが成功すると、readiness プローブは残りのコンテナーを監視します。 readiness プローブに失敗すると、コンテナーは NotReady としてマークされ、コンテナーが属するサービスに対するリクエストの受信から削除されます。
現時点では readiness プローブのみを含めるように pod-a.yaml を変更しました。
1 apiVersion : "v1" 2 kind : "Pod" 3 メタデータ : ⋯ 4 名前 : "pod-a" 5 仕様 : ⋯ 6 コンテナー : ⋯ 7 - 名前 : "app" ⋯ 8 イメージ : "my-app:latest" 9 readinessProbe : ⋯ 10 httpGet : ⋯ 11 パス : "/ready" 12 ポート : 8080 13 periodSeconds : 3 14 FailureThreshold : 1 15 successThreshold : 1
3 秒ごとに /ready エンドポイントに送信します。 1 回失敗すると、コンテナーは NotReady 状態になります。以下のデモで /ready を 200 から 503 に切り替えて、コンテナーが準備完了になっていないことを確認します。
クラスターを一時停止します。pod-a の準備が整うまで待ちます。まだ完了していません。
/準備完了 200 503
帯域外プロービング
上記のデモをリセットして、最初の数秒をもう一度見てください。
Readiness プローブは 3 秒ごとよりもはるかに頻繁に送信されます。何が与えますか？ Kubernetes のドキュメントには次のように書かれています
これ:
コンテナーが準備完了ではない間、readiness プローブは構成された時間以外に実行される可能性があります。
期間秒
間隔。これは、Pod の準備をより速くするためです。
素晴らしく曖昧。私が Webernetes で行った作業の中で、これらの症状を引き起こす可能性のあるものがたくさんあることがわかりました。
帯域外プローブ。あまり深く考えずに、ポッドのほとんどの更新はコンテナーの実行中に行われます。

NotReady は帯域外プローブをトリガーできます。たとえば、注釈やステータスの更新など。たくさん
気づかないうちにポッドが更新されている可能性があります。
知るのは楽しいですが、最終的には実際に頼るべきものではありません。
上記のデモでは、failureThreshold と successThreshold を 1 に設定していますが、単一の一時的な障害によってポッドがサービスから削除されることは望ましくありません。以下では、しきい値を 2 に設定しています。 /ready を再度 503 に設定すると、コンテナーが NotReady になるまでに 2 回の失敗が必要になることがわかります。
クラスターを一時停止します。pod-a の準備が整うまで待ちます。まだ完了していません。
/準備完了 200 503
ここで、準備完了から準備完了に切り替えるときに、帯域外が発生していることに気づくかもしれません。
これは前と同じ理由です。ポッドは NotReady で、ステータスが更新されたところです。
デフォルトでは、successThreshold は 1、failureThreshold は 3 です。通常は良好です。
特別な理由がない限り、デフォルトを変更することはお勧めしません。
このセクションをブックマークする Readiness Probe があるのに、Startup Probe が必要なのはなぜですか?
上記のデモでは、readiness Probe のみを使用します。プローブはすぐに開始され、コンテナーの起動作業が完了するまでは成功しません。これはまさに私の起動プローブが行っていた仕事です。では、なぜ両方のプローブ タイプが必要なのでしょうか?
起動プローブは、初期化が完了するまで準備と活性の開始を遅らせます。
彼ら全員

[切り捨てられた]

## Original Extract

Learn how probes work interactively on simulated Kubernetes clusters running in your browser.

How Kubernetes probes work | ngrok blog Skip to main content ngrok home / ngrok blog home blog open mobile navigation Products Problems We Solve Resources Docs Pricing Log in Sign up Products Gateway
Build & test Share a local app
Deploy & run Route traffic to self-hosted models
Deliver & connect Deliver and secure APIs
Connect privately across networks
Make LLMs 4x smaller and 2x faster
Search… Control ⌃ K Newsletter RSS Aug 19, 2026
I’m going to show you, really show you , how probes work in Kubernetes. How they can make your application more resilient, and how they can help you prevent avoidable mistakes. Like restart loops that take hours to recover from, and dropping requests during rollouts.
Every interactive demo in this post uses webernetes , my partial port of the Kubernetes to TypeScript. It contains more than 100,000 lines of ported Kubernetes Go code to run a simulated cluster right here in your browser . I verified the behaviour of these demos against k3s and managed to find a bug in Kubernetes! More on that later.
The 3 types of probe and what they’re for.
How to configure and combine them.
How common misconfigurations fail.
How probes affect Deployment speed.
Bookmark this section A pod without probes
I want to run a pod with a single container. Here’s its manifest, pod-a.yaml :
1 apiVersion : "v1" 2 kind : "Pod" 3 metadata : ⋯ 4 name : "pod-a" 5 spec : ⋯ 6 containers : ⋯ 7 - name : "app" ⋯ 8 image : "my-app:latest"
This image, my-app:latest , spends a few seconds initialising before listening on port 8080. You will see this below when you click restart to send the container a signal, causing it to crash and get started back up by Kubernetes. You can pause or reset any demo at any time.
Pause cluster 0 / 2 Restart container Not yet complete.
Restart container
After the first crash, the container restarts straight away. After the second, Kubernetes imposes a CrashLoopBackOff on it before starting it again. By default this delay is 10 seconds, doubling with each crash up to a maximum wait of 5 minutes. I shortened it to 3 seconds for this demo.
In both cases, Kubernetes considers the container Ready as soon as it starts, even though we know it’s not. It’s still doing startup work and not listening on port 8080.
Next I’ll add pod-b , which sends a request to pod-a every 2 seconds. Throughout the post, you can think of pod-b as any source of client traffic: an ingress controller, a load balancer, inter-service requests, etc.
If you restart pod-a in the demo below while a request is on its way, that request will fail .
Pause cluster Cause a request to fail Not yet complete.
Restart container
From the moment you restart the container until its startup work finishes, requests will fail , even though the container is considered Ready ! This is not what I want. I need Kubernetes know when pod-a is ready to receive traffic.
For this, Kubernetes gives us probes . Probes are periodic checks sent to containers to determine their health. They come in three flavours:
Startup probes determine whether my application inside the container has started.
Readiness probes determine whether my application is ready to receive traffic.
Liveness probes determine whether my application needs to be restarted.
It sounds like startup probes are best suited to the problem I showed you in the demos above, so let’s start there.
Bookmark this section Startup probes
Below, I’ve added a startup probe to pod-a.yaml :
1 apiVersion : "v1" 2 kind : "Pod" 3 metadata : ⋯ 4 name : "pod-a" 5 spec : ⋯ 6 containers : ⋯ 7 - name : "app" ⋯ 8 image : "my-app:latest" 9 startupProbe : ⋯ 10 httpGet : ⋯ 11 path : "/startup" 12 port : 8080 13 periodSeconds : 1 14 failureThreshold : 5
It’s an httpGet probe that sends a GET /startup request to the pod on port 8080. Status codes 200-399 count as a success. This happens every periodSeconds seconds, and is allowed to fail failureThreshold consecutive times before Kubernetes kills the container. This gives my container ~5 seconds to complete its startup work.
Probes are sent by a process called the kubelet . Each node in the cluster has its own kubelet, and it’s the kubelet’s job to make sure the right pods are running and being probed for each node.
When you restart pod-a below, it now shows as NotReady . Kubernetes is now aware that pod-a hasn’t initialised yet. It only becomes Ready after the first startup probe succeeds.
Pause cluster 0 / 2 Restart container Not yet complete.
Restart container
NotReady is the default for pods with containers that have a startup probe . However,
even when not ready, pod-b still sends requests to
pod-a and those requests still fail during the container’s startup period. This
is because I’ve configured pod-b to send requests directly to pod-a ’s IP address, which
bypasses the readiness mechanism.
I'm lying a bit about NotReady
To fix these failed requests I need to graduate to a more production-grade setup: multiple copies of pod-a with requests load-balanced between them. I’m going to create a ReplicaSet configured to run 2 replicas of pod-a and a Service to load balance between them.
1 apiVersion : "apps/v1" 2 kind : "ReplicaSet" 3 metadata : ⋯ 4 name : "replica-set-a" 5 spec : ⋯ 6 # Run 2 copies of the pod defined under `template`. 7 replicas : 2 8 selector : ⋯ 9 matchLabels : ⋯ 10 # Consider pods with this label to be part of this replica set. 11 app : "pod-a" 12 template : ⋯ 13 metadata : ⋯ 14 labels : ⋯ 15 app : "pod-a" 16 spec : ⋯ 17 # The same pod spec from before. 18 containers : ⋯ 19 - name : "app" ⋯ 20 image : "my-app:latest" 21 startupProbe : ⋯ 22 httpGet : ⋯ 23 path : "/startup" 24 port : 8080 25 periodSeconds : 1 26 failureThreshold : 5
service-a.yaml
1 apiVersion : "v1" 2 kind : "Service" 3 metadata : ⋯ 4 name : "service-a" 5 spec : ⋯ 6 selector : ⋯ 7 # Load-balance between pods that have this label. 8 app : "pod-a" 9 ports : ⋯ 10 # Send requests to this port on the pods. 11 - port : 80 ⋯ 12 targetPort : 8080
pod-b will from now on send requests to the DNS name Kubernetes creates for the Service, in this case
service-a.default.svc.cluster.local , instead of directly to an individual pod. Kubernetes uses a pod’s Ready
condition to include or exclude it from Service load balancing.
Below you can click the restart button to crash only the top container . Notice that when the top container is starting up , requests are always sent to the bottom container. When a container is NotReady , it marks the whole pod not ready and it won’t get traffic from any Services it is part of.
Pause cluster 0 / 2 Restart top container Not yet complete.
Restart top container
Despite this, requests can still fail if they’re in-flight when you restart the top container. This happens because the restart button crashes the container abruptly. It doesn’t get a chance to finish in-flight requests.
The better thing to do here is delete the pod and rely on the ReplicaSet to bring up a new one. This is better for 2 reasons:
Kubernetes gives pods a 30-second termination grace period by default, which I’ve configured to 2 seconds in this post so you don’t have to wait. When deleted, pods are considered terminating and Kubernetes removes them from any Services they’re part of. They won’t receive any new requests.
ReplicaSets don’t count terminating pods as active replicas, so they create replacements as soon as the deleted pod is terminating.
Together, graceful termination and the startup probe keep requests away from containers that are starting or stopping. In this next demo, clicking delete won’t cause any requests from pod-b to fail .
Pause cluster 0 / 2 Wait for containers to be ready Not yet complete.
Delete top pod
There’s always a pod ready to service a new request , making it safe to delete pods without interrupting user traffic.
How does this grace period actually work?
When you delete a pod, the kubelet first sends a SIGTERM signal to the container, then
terminationGracePeriodSeconds later it sends a SIGKILL if the container is still running. The Kubernetes docs on
pod lifecycle contain all of the gory details.
In this post, my containers are configured to exit 1 second after receiving the SIGTERM . I chose this because I know
it’s long enough for them to finish any in-flight requests. When configuring this on your own pods, make sure you
leave enough time for your longest-running requests to finish.
Bookmark this section How to misconfigure a startup probe
Earlier I mentioned that I’m giving my pod ~5 seconds to complete its startup work by setting failureThreshold to 5 with a periodSeconds of 1. Choose these values on your own containers carefully. Too little time can cause a container to crash-loop.
Setting the failureThreshold below will restart the container with the new value. Set it to 1 or 2 and see what happens.
Pause cluster Make pod-a crash loop Not yet complete.
failureThreshold 5
After a few restarts, pod-a is put in CrashLoopBackOff . The startup probe never gives the container enough time to start, so this demo crash-loops until you set failureThreshold back to 3 or above. When configuring this for your own containers, choose values that allow for your worst-case startup time.
Bookmark this section Readiness probes
After any startup probe succeeds, readiness probes monitor the container for the rest of its life. Failing a readiness probe marks the container NotReady and removes it from receiving requests for any Service it is part of.
I’ve modified pod-a.yaml to have just a readiness probe for now:
1 apiVersion : "v1" 2 kind : "Pod" 3 metadata : ⋯ 4 name : "pod-a" 5 spec : ⋯ 6 containers : ⋯ 7 - name : "app" ⋯ 8 image : "my-app:latest" 9 readinessProbe : ⋯ 10 httpGet : ⋯ 11 path : "/ready" 12 port : 8080 13 periodSeconds : 3 14 failureThreshold : 1 15 successThreshold : 1
I’m sending it to the /ready endpoint every 3 seconds. After a single failure, the container gets the NotReady condition. Switch /ready in the demo below from 200 to 503 and watch the container become not ready.
Pause cluster Wait for pod-a to become ready Not yet complete.
/ready 200 503
Out-of-band probing
Reset the demo above and watch the first few seconds again.
Readiness probes are sent way more often than every 3 seconds. What gives? The Kubernetes docs say
this:
While a container is not Ready, the readiness probe may be executed at times other than the configured
periodSeconds
interval. This is to make the Pod ready faster.
Wonderfully vague. In the work I did on webernetes I found that there are many things that can trigger these
out-of-band probes. Without going too into the weeds, most updates to the pod while the container is
NotReady can trigger an out-of-band probe. Things like updating annotations or status, for example. Many
things can be updating pods that you may not be aware of.
Fun to know, but ultimately not something you should rely on in practice.
The demo above sets failureThreshold and successThreshold to 1, but I don’t want a single transient failure to remove my pods from their Services. Below I’ve set the thresholds to 2. Set /ready to 503 again and notice it now takes 2 failures before the container becomes NotReady .
Pause cluster Wait for pod-a to become ready Not yet complete.
/ready 200 503
You may notice here that when flipping from ready to not ready, an out-of-band
probe can be fired.. This is for the same reasons as before. The pod is NotReady and its status just got updated.
By default successThreshold is 1 and failureThreshold is 3. Generally good
defaults that I don’t recommend changing unless you have a great reason.
Bookmark this section Why do we need startup probes if we have readiness probes?
The demos above only use a readiness probe . Probing starts straight away and doesn’t succeed until my container has finished its startup work. This is exactly the job my startup probe was doing, so why do we need both probe types?
Startup probes delay readiness and liveness starting until initialisation is complete.
They all

[truncated]

---
source: "https://dev.profullstack.com/~anthony/blog/023-post.html"
hn_url: "https://news.ycombinator.com/item?id=49360198"
title: "The AI storage pitch is a load balancing problem"
article_title: "The AI storage pitch is a load balancing problem — Chovy's Blog"
image: ""
author: "buffer_overlord"
captured_at: "2026-08-19T12:24:48Z"
capture_tool: "hn-digest"
hn_id: 49360198
score: 2
comments: 0
posted_at: "2026-08-19T11:41:31Z"
tags:
  - hacker-news
  - translated
---

# The AI storage pitch is a load balancing problem

- HN: [49360198](https://news.ycombinator.com/item?id=49360198)
- Source: [dev.profullstack.com](https://dev.profullstack.com/~anthony/blog/023-post.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T11:41:31Z

## Translation

タイトル: AI ストレージの提案は負荷分散の問題です
記事のタイトル: AI ストレージの提案は負荷分散の問題です — Chovy のブログ
説明: オブジェクト ストレージから GPU に供給することに関する F5 と Scality ウェビナーにサインアップし、その背後にある技術的な記事を読みました。興味深いのは、ホットスポットの回避、イレイジャーコーディングの実際のコスト、およびポスト量子 TLS がハンドシェイクよりも重要である理由です。

記事本文:
AI ストレージの提案は負荷分散の問題です
2026 年 8 月 19 日、Anthony “chovy” Ettinger 著。
これがどのように書かれたか: AI アシスタントを使用して自分のメモから草案を作成し、
その後私が編集しました。
今朝、ベンダーのウェビナーにサインアップしました。つまり、The Register のフォームに私の名前を記入することになりました。
そして私の会社。 F5 と Scality は 8 月 26 日に、オブジェクト データを十分に高速に移動することについて発表します。
GPU をビジー状態に保ちます。実際に参加するかどうかを決める前に、技術文書を読みに行きました。
ピッチはその上に構築されます。興味深いのはストレージではありません。
ボトルネックは負荷分散です
トレーニング ジョブは、S3 経由でデータを読み取ります。クライアントは 1 つのエンドポイントといずれかのストレージ ノードを解決します
クラスターの残りの部分はアイドル状態にある間、着陸したクラスターは打撃を受けます。そこはホットスポットであり、
ディスクとは関係ありません。
提供される修正は、ストレージ クラスターの前にロード バランサーを配置し、ストレージ クラスターとの接続を分散することです。
最小接続数、または最速。接続数ではなく、現在実行中のトランザクションをカウントします。
接続が開かれました。ノードの半分が以前よりも新しいヴィンテージである場合の比率モードもあります。
残りの半分は、新しいものにもっと手間をかけたいと考えます。グローバル ステアリングは最も近いサイトを選択します
測定された遅延によって。
そのどれもが AI に特化したものではありません。これは 2005 年以来人々が Web 層に対して行ってきたことです、と指摘しました。
代わりにストレージクラスターで。だからこそ機能するのです。
Scality RING は 3 つのノードと 200 TB から始まります。 60 KB 未満のオブジェクトは 2 つまたは 3 つ複製されます
コピー。それより大きいものは消去符号化されます。
EC(9,3) は、ノード全体に分散された 9 つのデータ チャンクと 3 つのパリティ チャンクを意味します。 3 つのノードを失う
そしてオブジェクトはまだ読み取ります。コストは 33% のオーバーヘッドです。 EC(8,4) は 50 で 4 回の障害に耐えます
パーセント。 3 つのコピーでの直接レプリケーションには 200% のコストがかかるため、次の目的で予約されています。
小さなもの

f.
彼らの実際の例: オブジェクトの 3% が 60 KB 境界以下にある場合、それらを複製します。
3 つのコピーで残りを EC(9,3) で消去コーディングすると、合計ストレージの約 41 パーセントが得られます。
頭上。主張されている耐久性は 149 であり、Wikipedia の耐久性よりも 9 が大きいです。
可用性テーブルは計算するのが面倒です。
ペタバイトも実行しません。私は取引に実際にかかるコストを知りたいと思っています。
他の場所にはすべてのコピーが 3 つあり、何も考えられていません。
騒々しい隣人がバケツに水漏れを起こす
マルチテナント ストレージには明らかな問題があります。1 つのテナントのジョブがディスク I/O を消費し、すべてのテナントのジョブがディスク I/O を消費します。
他のものではレイテンシーが悪化します。ここでの答えは、クライアントごとのリクエストをカウントする、エッジでのレート制限です。
アドレスを指定し、カウントが短い時間内にしきい値を超えると拒否され、クレジットが補充されます。
時間が経つにつれて。その下には、優先度の低いトラフィックを 10 Mbps 程度に制限する帯域幅プロファイルがあり、
また、レイヤー 4 のタイムアウトにより、意図的にオープンおよび放棄された接続を取得します。
高価なものの内部ではなく、その前でのレート制限。 nginxに制限を加えるのと同じ動き
アプリの前では、保護されるものだけがはるかに高価になります。
TLS 1.3 はパフォーマンスを重視し、確立に 2 ラウンド トリップではなく 1 ラウンド トリップを使用し、ゼロで確立します。
再開します。いいです。ジョブによって多くの接続が開かれると、それが加算されます。
もっと良い理由はその下にあります。ポスト量子鍵合意 (ML-KEM、NIST FIPS 203) には TLS 1.3 が必要です
そもそも存在すること。現在キャプチャされたトラフィックは、後で、おそらく 2030 年頃に保存して復号化できるようになります。
トレーニング コーパスは 5 年後も価値があり、今すぐ収集して後で復号化できます。
往復を剃るよりも鋭い議論。
おそらくそうではありません。この記事ですでに形が見えてきました。ウェビナーも同じスライドになります。
それらにはロゴが入っています。
私のサイズで仕事をするのに役立つもの:

スケールや健康にコストがかかるものの前にプロキシを置く
実際のもの（少なくとも HEAD リクエスト、または実際のアップロードおよびダウンロード プローブ）でチェックしてください。
資格情報に余裕がある場合)、クライアントごとのレート制限、および冗長オーバーヘッドを選択します。
目的。登録にはメールアドレスと会社名が必要でした。その背後にある技術記事
公開されており無料です。
出典:
DevCentral 上の Scality RING と F5 BIG-IP
そして
ウェビナーのサインアップ。
私を見つけてください: マストドン ·
GitHub ·
電子メール

## Original Extract

I signed up for an F5 and Scality webinar about feeding GPUs from object storage, then read the technical writeup behind it. The interesting parts are hot spot avoidance, what erasure coding actually costs, and why post quantum TLS matters more than the handshake.

The AI storage pitch is a load balancing problem
2026-08-19, by Anthony “chovy” Ettinger.
How this was written: drafted with an AI assistant from my own notes,
then edited by me.
I signed up for a vendor webinar this morning, which meant giving The Register's form my name
and my company. F5 and Scality are presenting on 26 August about moving object data fast enough to
keep GPUs busy. Before deciding whether to actually show up, I went and read the technical writeup
the pitch is built on. The interesting part is not the storage.
The bottleneck is load balancing
A training job reads its data over S3. Clients resolve one endpoint, and whichever storage node
they land on gets hammered while the rest of the cluster sits idle. That is a hot spot, and it has
nothing to do with disks.
The fix on offer is a load balancer in front of the storage cluster, spreading connections with
Least Connections, or Fastest, which counts transactions currently in flight rather than
connections opened. There is also a Ratio mode for when half your nodes are a newer vintage than
the other half and you want the new ones taking more work. Global steering picks the nearest site
by measured latency.
None of that is AI specific. It is what people have been doing to web tiers since 2005, pointed
at a storage cluster instead. That is why it works.
Scality RING starts at three nodes and 200 TB. Objects under 60 KB get replicated, two or three
copies. Anything larger gets erasure coded.
EC(9,3) means nine data chunks plus three parity chunks, spread across nodes. Lose three nodes
and the object still reads. The cost is 33 percent overhead. EC(8,4) survives four failures at 50
percent. Straight replication at three copies costs 200 percent, which is why it is reserved for
the small stuff.
Their worked example: if 3 percent of your objects are under that 60 KB line, replicating those
at three copies and erasure coding the rest at EC(9,3) lands you around 41 percent total storage
overhead. The claimed durability is fourteen nines, which is more nines than Wikipedia's
availability table bothers to calculate.
I do not run petabytes. I still like knowing what the trade actually costs, because the default
everywhere else is three copies of everything and no thought given.
Noisy neighbors get a leaky bucket
Multi tenant storage has the obvious problem: one tenant's job eats the disk I/O and everyone
else's latency goes bad. The answer here is a rate limit at the edge, counting requests per client
address and rejecting once the count passes a threshold in a short window, with credits replenishing
over time. Below that, a bandwidth profile capping lower priority traffic at something like 10 Mbps,
and layer 4 timeouts to reap connections opened and abandoned on purpose.
Rate limit in front of the expensive thing, not inside it. Same move as putting nginx limits in
front of an app, only the thing being protected costs a lot more.
TLS 1.3 gets pitched on performance, one round trip to establish instead of two, and zero on
resume. Fine, and it adds up when a job opens a lot of connections.
The better reason is underneath. Post quantum key agreement (ML-KEM, NIST FIPS 203) needs TLS 1.3
to exist at all. Traffic captured today can be stored and decrypted later, possibly around 2030. For
a training corpus that will still be valuable in five years, harvest now and decrypt later is a
sharper argument than shaving a round trip.
Probably not. The writeup already gave me the shape, and the webinar will be the same slides with
logos on them.
What transfers to work my size: put a proxy in front of whatever is expensive to scale, health
check it with something real (a HEAD request at minimum, or an actual upload and download probe if
you have credentials to spare), rate limit per client, and choose your redundancy overhead on
purpose. Registration cost me an email address and a company name. The technical article behind it
is public and free.
Sources:
Scality RING and F5 BIG-IP on DevCentral
and the
webinar signup .
Find me: Mastodon ·
GitHub ·
email

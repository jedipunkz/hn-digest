---
source: "https://planetscale.com/blog/the-feedback-loops-behind-kubernetes"
hn_url: "https://news.ycombinator.com/item?id=48560859"
title: "The feedback loops behind Kubernetes"
article_title: "The feedback loops behind Kubernetes — PlanetScale"
author: "CSDude"
captured_at: "2026-06-16T21:54:11Z"
capture_tool: "hn-digest"
hn_id: 48560859
score: 3
comments: 0
posted_at: "2026-06-16T19:43:06Z"
tags:
  - hacker-news
  - translated
---

# The feedback loops behind Kubernetes

- HN: [48560859](https://news.ycombinator.com/item?id=48560859)
- Source: [planetscale.com](https://planetscale.com/blog/the-feedback-loops-behind-kubernetes)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T19:43:06Z

## Translation

タイトル: Kubernetes の背後にあるフィードバック ループ
記事のタイトル: Kubernetes の背後にあるフィードバック ループ — PlanetScale
説明: Kubernetes はフィードバック コントローラー用のフレームワークです。必要なものを書き留め、存在するものを観察し、次の変更を加え、これを繰り返します。

記事本文:
Database Traffic Control™ の紹介: Postgres クエリ トラフィックのリソース バジェット。さらに詳しく ログイン |始めましょう |お問い合わせ プラットフォーム ▾
パート 1: Postgres を手動で実行する 1 つのコンテナー、1 つのマシン
パート 2: Kubernetes をどのように再発明したか その他のループ
Kubernetes に変換された for ループ
エッジトリガー通知、レベルトリガーロジック
インフォーマー、ワークキュー、およびキャッシュ
設定値と測定変数: 仕様とステータス
実際の演算子における「観察」が実際に意味するもの
すべてのエッジが API サーバーから取得されるわけではありません
PlanetScale Postgres は、クラウドで Postgres を実行する最速の方法です。プランは月額わずか 5 ドルから始まります。
Kubernetes の背後にあるフィードバック ループ
ファティ・アルスラン [ @ fatih ] | 2026 年 6 月 16 日
過去 10 年間、クラスターの運用、ホスト型 Kubernetes の構築支援、Kubernetes オペレーターの作成など、私の仕事のほとんどは Kubernetes を背景にしてきました。 PlanetScale では、これは現在、Postgres や MySQL などのステートフル システムを本番環境で実行することを意味します。 Kubernetes には多くの側面がありますが、ここでは 1 つの側面だけを説明したいと思います。それは、Kubernetes が大規模なワークロードの実行に優れている理由です。
オペレーターは実際に何をしているのかとよく聞かれます。標準的な答えは、「望ましい状態を調和させる」です。これは正しいですが、ほとんど何もわかりません。
オペレーターはフィードバックコントローラーです。これは、サーモスタットを作動させたり、クルーズ コントロールで車を一定の速度に保つのと同じ閉ループです。私たちの場合、制御されるものはデータベースです。私は何年にもわたってこれらのループを構築してきましたが、ループをクリックさせるための私が知っている最善の方法は、最初に Kubernetes を無視することです。 Kubernetes には、日常的にそう呼ばれていなくても、制御理論が満載です。
Kubernetes の 1 行を確認する前に、本番データベースを手動で実行し、ゆっくりとフィードバックを反映させます。

p は単独で表示されます。次に、そのループを、本番環境に必要な要素 (ストア、監視、キュー、再試行など) とともに Kubernetes にマッピングします。最後に、これらのループの 1 つが実際のオペレーターでどのように見えるかを見ていきます。
コンテナーと kubectl について十分に理解していると役に立ちますが、Kubernetes の専門家である必要はありません。べき等 、 ファンイン 、 最終整合性 などの用語を使用し、重要な部分を紹介していきます。
ゆっくりと始めて、徐々にペースを上げていきます。各部分は前の部分に基づいて構築されます。
パート 1: Postgres を手動で実行する
ゼロから始めましょう。 Linux ボックス上で Postgres を実行したいのですが、コンテナ内で Postgres を実行する必要があります。それを開始するには、次を実行します。
docker run -d --name pg \
-e POSTGRES_PASSWORD=秘密 \
ポストグレ:18
それだけです。 Postgres が実行されています。私のアプリはそれに接続し、いくつかの行を書き込み、すべてが正常に動作します。しかしその後、マシンは消えてしまいます。クラウド プロバイダーがインスタンスを再利用するか (ハードウェア障害が発生するか、スポット インスタンスが取り戻される)、私がセットアップの新しいバージョンを出荷します。これは、古いコンテナーを停止して、その場所で新しいコンテナーを開始することを意味します。いずれにせよ、コンテナが交換され、データは失われます。コンテナー ストレージは一時的だったので、永続ボリュームをアタッチしませんでした。
私が望むもの (Postgres、実行中、データ) と私が持っているもの (コンテナーまたはノードがなくなるとストレージが消えるコンテナー) の間にはすでにギャップがあります。この投稿の残りの部分は、そのギャップと、それを埋めるために私たちが構築する機構について説明します。
使用できるノード (サーバー) が数百あると想像してください。すでに他のワークロードを実行しています。どれがこのデータベースを実行するかを決める必要があります。そこで、最も混雑していないと思われるボックスに SSH で接続し、そこでコンテナを起動します。
ssh ノード-07 'docker run -d --name pg ... postgres:18'
ノード 07 を選択したのは、

十分に暇そうに見えた。私はそれを追跡し始め、ある種の構成ファイルに保存し、リポジトリにプッシュします。
コンテナストレージは一時的なものであるため、実際のブロックデバイスを接続する必要があります。クラウドでは、これは EBS ボリューム (AWS など) です。ベアメタルでは物理ディスクです。それがブロック デバイスであると仮定すると、これが通常行うことです。ボリュームをプロビジョニングし、ノードに接続し、フォーマットし、マウントし、マウントで Postgres のデータ ディレクトリを指定します。
# 最初にクラウド CLI を使用してプロビジョニングとアタッチを行い、次にノード上で行います。
mkfs.ext4 /dev/nvme1n1
mkdir -p /var/lib/pg-data
/dev/nvme1n1 /var/lib/pg-data をマウントします
docker run -d --name pg \
-e POSTGRES_PASSWORD=秘密 \
-v /var/lib/pg-data:/var/lib/postgresql \
ポストグレ:18
これらには多くの手順があり、それぞれの手順が途中で失敗する可能性があります。そして、後でディスクがいっぱいになると、Postgres は書き込みの受け入れを停止するため、ボリュームのサイズを手動で変更する必要があります。最初はクラウド プロバイダーを介して、次にファイル システム内で再度ボリュームを変更します。
単一の Postgres インスタンスは単一障害点になります。高可用性が必要です。1 つのプライマリと 2 つのレプリカです。これらは 3 つの異なるマシン上に存在し、それらの間でストリーミング レプリケーションを行う必要があります。したがって、同じ手順を再度、node-07 、node-12、node-19 で 3 回実行します。また、 Primary_conninfo 、レプリケーション スロットなどすべてを手動でレプリケーションに接続します。
これで、3 つの Postgres インスタンスを持つ 3 つのノードができました。インスタンスの 1 つはプライマリです (ここでは、 node-07 です)。しかし、これにより、プライマリのノードが停止した場合はどうすればよいかなど、新たな問題が生じます。
ここで私たちが解決しなければならないことがもう一つあります。レプリカはプライマリに到達する必要があり、プライマリはその接続を受け入れる必要があります。これらのアドレスはすべて、コンテナーの再起動時に変更される IP です。
最初に行うのは、IP をハードコーディングすることです。 node-07 のアドレスをレプリカの設定に書き込みます。

プライマリの pg_hba.conf にレプリカのアドレスを保存し、小さな /etc/hosts テーブルを作成してどこかに保存します。
# 各レプリカの postgresql.auto.conf で、プライマリが新しい IP で再作成されるまで
Primary_conninfo = 'ホスト=10.4.7.21 ポート=5432 ユーザー=レプリケータ ...'
しかし、まだ問題があります。プライマリが初めて別の IP で再作成されると、クラスター全体がバラバラになってしまいます。
さて、ここからはこれらの問題を解決する方法を考え始めます。これまで説明したものはすべて壊れる可能性があり、たとえ修正したとしても壊れ続けるでしょう。
レプリカ プロセスが停止し、元に戻りません。
プライマリに障害が発生したため、レプリカを昇格する必要があります。
2 つのノードで構成を変更しましたが、3 番目のノードでは忘れていました。現在、それらは同期していません。
単純な稼働時間モニターを設定し、これらすべてのケースに対してページングを受けると仮定しましょう。夜間のポケットベルを避けるために、私たちは賢明な行動をとります。それは、スクリプトを書くことです。そこで、数秒ごとに起動して各ノードを調べ、問題があれば修正するループを作成することにしました。
true である一方で、する
ノード07 ノード12 ノード19 のノードの場合。する
もし！ ssh " $node " 'pg_isready -q' ;それから
ssh " $node " 'docker start pg' # 死んだので元に戻します
フィ
使用法 = $( ssh " $node " "df --output=pcent /var/lib/pg-data | tail -1 | tr -dc 0-9" )
if [ " $usage " -gt 80 ];それから
give_volume " $node " # ディスクがいっぱいになり、より大きくします
フィ
完了しました
睡眠5
完了しました
これは Bash で書かれており、おそらく大量のバグがあります。ここで何か気づきましたか？ループは、データベースがどのようにして悪い状態になったかを気にしません。 5 秒ごとに世界の現状を確認し、「現実は私が望んでいることと一致しますか?」という質問をします。
プロセスが停止している場合は、プロセスを開始します。ディスクがいっぱいになっている場合は、ディスクを拡張します。ループを 1 回実行しても、1,000 回実行しても結果は同じです。これは、各アクションが条件付きであるためです。

彼の現在の状態。スクリプトはべき等です。
もう少し複雑にしてみましょう。 max_connections を 100 から 500 に増やす必要があります。これはリロードのみの変更ではありません。 PostgreSQL では、サーバーの起動時にのみ設定できるとしているため、手動バージョンでは、各ボックスに ssh で接続し、 postgresql.conf を編集し、Postgres を再起動して、3 つすべてが実行されたことを確認します。
手動でノードに SSH 接続することはもう必要ないことはわかっているので、以前と同じことを行います。つまり、必要な値を 1 か所に書き留めて、それを強制するようにループに教えます。
WANT_MAX_CONNECTIONS = 500
ノード07 ノード12 ノード19 のノードの場合。する
have = $( ssh " $node " "psql -tAc 'show max_connections'" )
if [ " $have " != " $WANT_MAX_CONNECTIONS " ];それから
ssh " $node " "sed -i 's/^max_connections.*/max_connections = $WANT_MAX_CONNECTIONS /' /var/lib/pg-data/postgresql.conf"
ssh " $node " "docker restart pg"
フィ
完了しました
これは先ほどと同じ考え方です。必要なもの (変数) を読み取ります。私が持っているもの（クエリ）を観察してください。差異がある場合は、差異を埋めるための措置を講じます。繰り返しますが、最後に変更したかどうかは追跡しません。私が行うのは、ループごとに比較して収束することだけです。
まず、望ましい状態を 1 か所に書き留めます: 3 つのインスタンス、このディスク サイズ、 max_connections = 500 。数秒ごとにシステムの実際の状態を観察します。私はその差を計算します。私はその差を埋めるためにあらゆる行動をとります。それから私はまたそれを永遠に繰り返します。
それは閉じたフィードバック ループです。 「閉鎖」という言葉が重要です。これは、システムの出力が次の決定にフィードバックされることを意味します。 docker start は実行せず、データベースは正常であると仮定します。もう一度データベースを確認してみます。それでもダメだったら、また行動します。それがすでに正しい場合は、何もしません。
良い点は、同じループがさまざまな問題に対して機能することです。それは解決できます

停止したプロセスを起動するか、ディスクを拡張するか、または max_connections = 500 をプッシュします。行動は変わりますが、形は変わりません。欲しいものを読み、持っているものを観察し、比較し、行動し、繰り返します。同じものを制御理論の名前を付けてブロック図に描くと次のようになります。
制御理論の語彙が私のシェル スクリプトにどのようにきれいにマッピングされているかを次に示します。
セットポイントは希望する状態、つまりスクリプトの先頭にある変数 (ディスク サイズ、max_connections など) です。
測定された出力は、私が観察したものです: pg_isready 、 df 、 show max_connections 。
誤差 (e) はそれらの差です。
コントローラーはループの本体であり、何を行うかを決定する if ステートメントです。それはスクリプト全体ではありません。
アクチュエーターはアクションを実行するものです: ssh と docker start 。
プラントとは、制御対象のシステム、Postgres とそのディスクです。
これは、開ループ制御を理解するための良い方法にもなります。私の最初の試みは、ssh で接続し、コマンドを実行し、終了するというオープンループでした。つまり、アクションを起動し、それが機能したと想定しました。 Bash スクリプトは、測定された状態を次の決定にフィードバックし続けるため、閉ループです。
Bash ループは運用コントロール プレーンではありません。問題点をいくつか挙げると、次のとおりです。
同時実行制御がないため、スクリプトの 2 つのコピーが相互に競合する可能性があります。両方が別のレプリカを昇格することを決定したと想像してください。
唯一の実際の状態「フェイルオーバー中ですか?」をシェル変数に保持しますが、プロセスとともに終了する可能性があります。
何かが変更されたかどうかを 5 秒ごとにすべてのノードにポーリングします。これは 3 つのノードには問題ありませんが、3,000 ノードには高すぎます。
SSH 自体がタイムアウトになったときに何をすべきかわかりません。
そして、2 番目の種類のリソース、接続プーラー、バックアップ ジョブ、別のリージョンのリードレプリカが必要になった瞬間に、これをコピーして貼り付けます。

全体の構造。
スクリプトも失敗したらどうなるでしょうか?それでは誰が運営しているのでしょうか？このスクリプトを強化し続けることもできますが、それがどこに行くのかを見てください。目的の状態を表す実際のストア、ポーリングの代わりに監視、作業キュー、再試行、リーダーの選出が必要になります。 Kubernetes を再構築することになります。本当のプラットフォームはすでに存在しており、それが Kubernetes です。
パート 2: Kubernetes をどのように再発明したか
これで、パート 1 で手作業で作成したものを Kubernetes にマッピングできるようになりました。ほとんどすべてがすでにそこに存在しています。私たちが注目するのは演算子です。
ウォッチドッグ ループの前に手動で構築した部分のいくつかを見てみましょう。これらのコンポーネントはすでに名前がわかっています。気づいていないかもしれませんが、これらはコントローラーとしても機能します。
コンテナー (kubelet) をスピンアップします。まず簡単に定義します。ポッドは Kubernetes が実行する最小のものであり、ノード上で一緒にスケジュールされ、そのネットワークを共有する 1 つ以上のコンテナです。私たちにとって、それは Postgres コンテナーです。すべてのノードで kubelet と呼ばれるエージェントが実行されます。その望ましい状態は、そのノードに割り当てられた一連の Pod であり、API サーバーから学習します。監視される状態は、実際に実行されているコンテナのセットであり、コンテナ ランタイムから取得されます。それらが異なると、ミスが始まります

[切り捨てられた]

## Original Extract

Kubernetes is a framework for feedback controllers: write down what you want, observe what exists, make the next change, and repeat.

Introducing Database Traffic Control™: resource budgets for your Postgres query traffic. Learn more Log in | Get started | Get in touch Platform ▾
Part 1: running Postgres by hand One container, one machine
Part 2: how we reinvented Kubernetes The other loops
The for-loop translated to Kubernetes
Edge-triggered notifications, level-triggered logic
Informers, the work queue, and a cache
Setpoint and measured variable: spec and status
What "observe" actually means in a real operator
Not every edge comes from the API server
PlanetScale Postgres is the fastest way to run Postgres in the cloud. Plans start at just $5 per month.
The feedback loops behind Kubernetes
Fatih Arslan [ @ fatih ] | June 16, 2026
For the last decade, Kubernetes has been the backdrop to most of my work: operating clusters, helping build hosted Kubernetes, and writing Kubernetes operators. At PlanetScale, that now means running stateful systems like Postgres and MySQL in production. Kubernetes has many faces, but here I want to talk about one face only: why it is so good at running workloads at scale.
People ask me what an operator actually does. The canonical answer is: "it reconciles desired state." This is correct, but it also tells you almost nothing.
An operator is a feedback controller. It's the same closed loop that runs a thermostat or keeps your car at a fixed speed on cruise control. In our case, the thing being controlled is a database. I have been building these loops for years, and the best way I know to make them click is to ignore Kubernetes at the beginning. Kubernetes is full of control theory, even if we don't call it that in the day-to-day.
Before we look at a single line of Kubernetes, we're going to run a production database by hand and slowly let the feedback loop appear on its own. Then we'll map that loop to Kubernetes, with the pieces production needs: a store, watches, queues, retries, and more. At the end, we'll look at what one of these loops looks like in a real operator.
A working understanding of containers and kubectl helps, but you don't need to be a Kubernetes expert. I'll use terms like idempotent , fan-in , and eventual consistency , and introduce the parts that matter as we go.
We're going to start slow and gradually ramp things up. Each part builds on the previous.
Part 1: running Postgres by hand
Let's start from scratch. I want to run Postgres on a Linux box, and I need it inside a container. To start it, we run:
docker run -d --name pg \
-e POSTGRES_PASSWORD=secret \
postgres:18
That's it. Postgres is running. My app connects to it, writes some rows, and everything works fine. But then the machine goes away: the cloud provider reclaims the instance (hardware fails, or a spot instance gets taken back), or I ship a new version of my setup, which means stopping the old container and starting a fresh one in its place. Either way, the container is replaced, and my data is gone. The container storage was ephemeral, and I did not attach any persistent volume to it.
There is already a gap between what I want (Postgres, running, with my data) and what I have (a container whose storage disappears when the container or node goes away). The rest of this post is about that gap and the machinery we build to close it.
Imagine we have hundreds of nodes (servers) we can use. I already have other workloads running on them. I need to decide which one runs this database. So I ssh into the box that looks the least busy and start the container there.
ssh node-07 'docker run -d --name pg ... postgres:18'
I picked node-07 because it looked idle enough. I start keeping track of it, save it in some sort of config file, and push it to some repo.
Container storage is ephemeral, so I have to attach a real block device. In the cloud this is an EBS volume (e.g. on AWS); on bare metal it's a physical disk. Assuming it's a block device, this is what we usually do: provision the volume, attach it to the node, format it, mount it, and point Postgres' data directory at the mount.
# provision + attach first with cloud CLI, then on the node:
mkfs.ext4 /dev/nvme1n1
mkdir -p /var/lib/pg-data
mount /dev/nvme1n1 /var/lib/pg-data
docker run -d --name pg \
-e POSTGRES_PASSWORD=secret \
-v /var/lib/pg-data:/var/lib/postgresql \
postgres:18
These are a lot of steps, and each one can fail halfway. And if the disk fills up later, Postgres stops accepting writes and we have to resize the volume by hand: first through the cloud provider, then again inside the filesystem.
A single Postgres instance is a single point of failure. We want high availability: one primary and two replicas. These need to be on three different machines, with streaming replication between them. So we do the same steps again, three times, on node-07 , node-12 , and node-19 . I also wire up replication by hand: primary_conninfo , replication slots, all of it.
Now we have three nodes with three Postgres instances. One of the instances is the primary (here it's node-07 ). But this raises new problems, like what to do if the primary's node dies?
Here is another thing we have to solve. The replicas need to reach the primary, and the primary needs to accept their connections. And every one of these addresses is an IP that changes when a container restarts.
The first thing I do is hard-code the IPs. I write node-07 's address into the replicas' config, I list the replicas' addresses in the primary's pg_hba.conf , and I keep a small /etc/hosts table and save it somewhere.
# on each replica's postgresql.auto.conf, until the primary is recreated with a new IP
primary_conninfo = 'host=10.4.7.21 port=5432 user=replicator ...'
But we still have a problem: the first time the primary is recreated with a different IP, the whole cluster falls apart.
Now, this is where we start thinking about how to solve these issues. Everything described so far can break, and will continue to break even if I fix it:
A replica process dies and doesn't come back.
The primary fails and a replica has to be promoted.
A config I changed on two nodes but forgot on the third one. They are now out of sync.
Let's assume we've set up a simple uptime monitor and we're going to get paged for all these cases. To avoid getting paged at night, we do the sensible thing: write a script. So we decide to write a loop that wakes up every few seconds, looks at each node, and fixes whatever's wrong.
while true ; do
for node in node-07 node-12 node-19 ; do
if ! ssh " $node " 'pg_isready -q' ; then
ssh " $node " 'docker start pg' # it died, bring it back
fi
usage = $( ssh " $node " "df --output=pcent /var/lib/pg-data | tail -1 | tr -dc 0-9" )
if [ " $usage " -gt 80 ]; then
grow_volume " $node " # disk filling, make it bigger
fi
done
sleep 5
done
It's written in Bash, and probably has tons of bugs. You notice something here? The loop doesn't care how the database got into a bad state. Every five seconds it looks at the current state of the world and asks this question: does reality match what I want?
If a process is down, start it. If a disk is filling, grow it. Run the loop once or run it a thousand times and the result is the same, because each action is conditional on the current state. The script is idempotent .
Let's make things a little more complex. I need to raise max_connections from 100 to 500. This one is not a reload-only change. PostgreSQL says it can only be set at server start, so the manual version is to ssh into each box, edit postgresql.conf , restart Postgres, and check that it took on all three.
Because I know that ssh'ing into the nodes manually isn't a thing I want anymore, I do the same thing we did previously: I write the desired value down in one place and teach the loop to enforce it.
WANT_MAX_CONNECTIONS = 500
for node in node-07 node-12 node-19 ; do
have = $( ssh " $node " "psql -tAc 'show max_connections'" )
if [ " $have " != " $WANT_MAX_CONNECTIONS " ]; then
ssh " $node " "sed -i 's/^max_connections.*/max_connections = $WANT_MAX_CONNECTIONS /' /var/lib/pg-data/postgresql.conf"
ssh " $node " "docker restart pg"
fi
done
This is the same idea as before. I read what I want (a variable). Observe what I have (a query). If they differ, I take an action to close the difference. Again, I don't track whether I changed it last time. All I do is compare and converge , every loop.
I started with a desired state that was written down in one place: three instances, this disk size, max_connections = 500 . Every few seconds I observe the actual state of the system. I compute the difference. I take whatever action closes that difference. Then I do it again, forever.
That's a closed feedback loop . The word "closed" matters. It means the output of the system is fed back into the next decision. I don't run docker start and assume the database is fine. I check the database again. If it is still wrong, I act again. If it is already correct, I do nothing.
The nice part is that the same loop works for different problems. It can restart a dead process, grow a disk, or push max_connections = 500 . The action changes, but the shape stays the same: read what I want, observe what I have, compare them, act, repeat. If I draw the same thing as a block diagram, with the control theory names added, it would look like this:
Here is how the vocabulary from control theory maps cleanly onto my shell script:
The setpoint is my desired state, the variables at the top of the script (disk size, max_connections and so on).
The measured output is what I observe: pg_isready , df , show max_connections .
The error (e) is the difference between them.
The controller is the body of the loop, the if statements that decide what to do. It is not the whole script.
The actuator is what carries out the action: ssh plus docker start .
The plant is the system being controlled, Postgres and its disk.
That also gives us a nice way to understand open-loop control. My very first attempt, ssh in, run the command, and walk away, was open-loop: fire an action and assume it worked. The Bash script is closed-loop because it keeps feeding the measured state back into the next decision.
A Bash loop is not a production control plane. Just to name a few issues with it:
It has no concurrency control, so two copies of the script can race each other. Imagine both deciding to promote a different replica.
It keeps its only real state, "am I mid-failover?", in a shell variable that could die with the process.
It polls every node every five seconds whether anything changed or not, which is fine for three nodes, but too expensive for three thousand nodes.
It has no idea what to do when the ssh itself times out.
And the moment I want a second kind of resource, a connection pooler, a backup job, a read replica in another region, I'm copy-pasting this whole structure.
What if the script also fails? Who runs it then? We could keep hardening this script, but look at where it goes: we would need a real store for the desired state, watches instead of polling, a work queue, retries, leader election. We would be rebuilding Kubernetes. The real platform already exists, and it's Kubernetes.
Part 2: how we reinvented Kubernetes
Now we can map what we hand-rolled in Part 1 to Kubernetes. Almost all of it already exists there. The operator is the part we care about.
Let's go through some of the pieces we built by hand before the watchdog loop. You already know these components by name. What you might not have noticed is that they also work like controllers.
Spinning up the container: the kubelet. First, a quick definition: a Pod is the smallest thing Kubernetes runs, one or more containers scheduled together on a node and sharing its network. For us it's the Postgres container. On every node runs an agent called the kubelet. Its desired state is the set of Pods assigned to its node, which it learns from the API server. Its observed state is the set of containers actually running, which it gets from the container runtime. When they differ, it starts the mis

[truncated]

---
source: "https://riverqueue.com/blog/migration-fatigue"
hn_url: "https://news.ycombinator.com/item?id=49277643"
title: "Migration fatigue, and how LLMs help us avoid it"
article_title: "Migration fatigue, and how LLMs help us avoid it - River blog"
author: "brandur"
captured_at: "2026-08-12T19:56:23Z"
capture_tool: "hn-digest"
hn_id: 49277643
score: 2
comments: 0
posted_at: "2026-08-12T19:47:29Z"
tags:
  - hacker-news
  - translated
---

# Migration fatigue, and how LLMs help us avoid it

- HN: [49277643](https://news.ycombinator.com/item?id=49277643)
- Source: [riverqueue.com](https://riverqueue.com/blog/migration-fatigue)
- Score: 2
- Comments: 0
- Posted: 2026-08-12T19:47:29Z

## Translation

タイトル: 移行疲労と、LLM が移行疲労の回避にどのように役立つか
記事のタイトル: 移行疲労と、LLM が移行疲労の回避にどのように役立つか - River blog
説明: LLM を使用して低速のアクティブ ジョブ レスキュー クエリを単一のプロデューサー状態パスに変換し、インデックスなしで最大 130 倍の高速化を実現することで、データベース移行の出荷を回避した方法。

記事本文:
移行疲労と、LLM が移行疲労の回避にどのように役立つか - River blog Docs features Blog About River Pro River on GitHub
すべての記事
移行疲労と、LLM が移行疲労の回避にどのように役立つか
LLM を使用して低速のアクティブ ジョブ レスキュー クエリを単一のプロデューサー状態パスに変換し、インデックスなしで最大 130 倍の高速化を実現することで、データベース移行の出荷を回避した方法。
プログラマの間では多くの点で意見が異なりますが、全員が一致団結できる点もいくつかあります。その 1 つは、下位互換性のない変更が煩わしいということです。本当に迷惑です。
重大な罪: River v0.39.0 では、下位互換性のない小さな変更をマイナー バージョンに同梱しましたが、これはすべきではありません。移行ツールの Validate 関数にオプション パラメータが追加されました。
res, err := migrator.Validate(ctx) res, err := migrator.Validate(ctx, nil) これは、Validate を残りの移行 API とより適切に連携させるための小さな変更です。また、これは比較的あいまいな機能にのみ影響します (ほとんどのインストールでは Validate は必要ありません)。River は技術的にはまだ 1.0 以前です。これが、リリース当時の私がそれを正当化する方法でしたが、それは悪い習慣であり、それは認めます。
最近では、ほとんどの人が、Dependabot を使用して依存関係の自動バンプを頻繁に行っています。毎週の更新が行われるとき、私が望むのは、いくつかのバージョンが変更され、緑の CI が表示されることだけです。そうすれば、あまり考えずにマージ ボタンを押すことができます。依存関係の 1 つにおける API の変更が原因でビルドが壊れた場合、その煩わしさは理不尽なものになります。小さな変化は大きな変化よりも良いですが、イライラは差分の大きさに比例しません。たった一文字の修正かもしれませんが、それでも私はイライラします。
移住、特別な種類の痛み
依存関係内の移行は、下位互換性のない変更に似ていますが、10 倍悪いです。

River の移行ではテーブル全体のロックなどの重大な運用上の責任が発生しないように注意していますが、大きなテーブルに対する単純な UPDATE であっても、潜在的に長時間実行される操作となり、ホット実行中のデータベースに望ましくない負荷がかかる可能性があるため、すべての移行をユーザーに明らかにしたいと考えています。
移行でよく見られる操作は CREATE INDEX です。運用環境では、書き込みのブロックを避けるために、常にインデックスを同時作成する必要があります (そうしないと、Postgres は SHARE ロックを必要とします)。ただし、River の移行ランナーは当然のこととしてトランザクションを使用し、Postgres ではトランザクション内での CREATE INDEX CONCURRENTLY が許可されていないため、アップグレード ノートには常に 2 つのパスが含まれています。
標準のマイグレーター呼び出し (river mitigator-up ...) は、開発のため、または高価な実稼働変更が手動で安全に適用された後に行われます。
ホット環境で代わりに実行できる手動の CREATE INDEX CONCURRENTLY ステートメントのリスト。
これにより、アップグレード プロセスに摩擦が加わると言っても過言ではありません。日常的な依存関係のアップグレードであるはずの作業が、数時間にわたる運用操作に変わってしまう可能性があり、缶が道路に蹴り飛ばされて誰かのバックログに消えてしまうため、アップグレードが数週間または数か月遅れる可能性があります。
さらに悪いことに、それは私が「移行疲労」と呼ぶ現象を引き起こします。依存関係があまりにも苦痛で何度もアップグレードできない場合、ユーザーはそれに飽きて苦情を言うでしょう。そうした訴えにもかかわらず痛みが続くと、他のことに移ってしまうでしょう。
移住疲労に対抗するために、私たちは次の 2 つの原則に従うよう努めています。
可能な限り少数の移行を出荷します。理想的な年間移行数はゼロです。
移行が必要な場合は、特定の期間に必要なものがすべて一緒に出荷できるように、できるだけ長期間プールしてください。
言うのは簡単です

終わり。通常、移行が必要な場合は、移行が必要になります。移行なしでは新機能を維持できないためです。ただし、常にそうとは限りません。
最近、アクティブ ジョブ レスキューをリリースしました。これは、プロデューサーのハートビートを使用して、River の通常のタイムアウト ベースのレスキューよりもはるかに早く、クラッシュしたクライアントによって孤立したジョブを回復します。
プロジェクトがゴールラインに近づくにつれ、ベンチマークの結果、懸念すべきことが明らかになりました。レスキュー クエリは大規模なデータ セットでは遅くなります。実行中のジョブが 100,000 で、プロデューサーが 1,000 ある場合、レスキュー対象が何も見つからないにもかかわらず、正常なジョブのパスに 8.9 秒かかりました。アクティブ レスキューが導入される前の同じ操作と比較すると、クエリ時間は最大 2,000 倍増加しました。
質問は、これまでと同じくらいありきたりなものでした。ここでは簡潔にするために重要な条件だけを示しますが、実行中のすべてのジョブを含む River_job テーブルと、アクティブなクライアントとそれらが動作しているキューを追跡する River_Producer テーブルを想像してください。このフラグメントは、まだ実行中とマークされているが、アクティブなプロデューサーが存在しないジョブをチェックしました。
EXISTS ( SELECT 1 FROM River_Producer WHERE client_id = job_Producer AND queue_name = job.queue AND created_at <= job.attempted_at AND updated_at < stale_cutoff ) AND NOT EXISTS ( SELECT 1 FROM River_Producer WHERE client_id = job_Producer AND queue_name = job.queue AND created_at <= job.attempted_at AND updated_at >= stale_cutoff ) = stale_cutoff )"> J 個のジョブと P 個のプロデューサーの場合、プロデューサー行チェックの数は J × P であり、大規模なデータ セットでは膨大な数になる可能性があります。これが、許容できないクエリ時間約 10 秒が発生した理由です。
River_Producer は主キーのみにインデックスが付けられていたため、最も明白な軽減策は、インデックスを追加して各 J × P スキャンを高速化することでした。
CREATE INDEX River_Producer_client_queue_created_at_idx ON River_Producer (client_id、queue_name、created_at);

これは従来の解決策ですが、移行疲労の元に戻ってしまいます。最近、大幅な移行を含むワークフロー V2 をリリースしました。これほど早く別の製品を出荷することは理想とは程遠いものでしたが、出荷しないことは新機能を保留することを意味します。
1 年前であれば、インデックスを追加して次に進んだ可能性は十分にあったと思います。やるべきことは常にあるので、現在のうさぎの穴を追い続けるのではなく、次の機能の構築にどのくらいの時間を費やすかを比較検討する必要があります。
しかし、それは1年前のことでした。 LLM 時代の優れた点の 1 つは、Codex にこのような問題をしばらく煮詰めさせて、何か見逃していないかどうかを確認できることです。調査と反復に数十分を費やした後、具体化された CTE を中心に構築された別のアプローチが生成されました。
生産者ステータス AS MATERIALIZED ( SELECT client_id, queue_name, min (created_at) FILTER ( WHERE updated_at < stale_cutoff ) AS stale_created_at, min (created_at) FILTER ( WHERE updated_at >= stale_cutoff ) AS active_created_at FROM River_Producer GROUP BY client_id, queue_name ) = stale_cutoff ) AS active_created_at FROM River_Producer GROUP BY client_id, queue_name )"> ジョブ行ごとにプロデューサー ステータスを個別にスキャンする代わりに、CTE はレスキュー パスごとにプロデューサー テーブルを 1 回スキャンし、 (client_id, queue_name) で要約します。その後、レスキュー クエリはその小さな具体化された結果を検討中のジョブと結合します。何千回も繰り返されたスキャンが 1 つにまとめられ、12 倍から 130 倍のパフォーマンスが生成されます。最初の最適化されていないパスと比較して高速化:
CTE は、river_Producer を 1 回だけ反復する必要があるため、上記で提案されたインデックスは不要になります。コストをかけずにすべてのパフォーマンスを実現します。
ユーザー側では、これが理想的な結果です。戻ることはありません。

ards と互換性のない変更があり、新しい移行はありません。 River Pro を新しいバージョンに更新すると、アクティブ ジョブ レスキューが追加作業なしで自動的にアクティブ化されます。
新しい LLM 時代の最も明白な副産物はコードの増加ですが、LLM によってコーディングに隣接した活動がどれほど容易になったかについてはあまり議論されていません。以前は、ベンチマークを注意深く構築し、各ブランチを実行し、問題を修正し、再度実行し、最後に結果を表にして要約するのに丸 1 日 (または複数日) かかりました。さて、あと20分です。
問題を解決するための 1 つのアプローチの開発にはかなりの費用がかかり、それについてもメタ的な問題が発生することになります。もし誰かがそれに数日または数週間を費やしたとしたら、より良いバージョンが登場したとしてもその存続を主張する可能性が高くなります。現在では、単一のアプローチを生成するのが速く、エゴは無視できる程度です。 LLM がすべての作業を行っていた場合、何かに執着するのは困難です。
最近では、より多くのソフトウェアを作成しているだけでなく、より優れたソフトウェアも作成しています。
© 2026 River Software, LLC.無断転載を禁じます。

## Original Extract

How we avoided shipping a database migration by using an LLM to turn a slow active job rescue query into a single producer-state pass, producing a ~130x speedup without an index.

Migration fatigue, and how LLMs help us avoid it - River blog Docs Features Blog About River Pro River on GitHub
All articles
Migration fatigue, and how LLMs help us avoid it
How we avoided shipping a database migration by using an LLM to turn a slow active job rescue query into a single producer-state pass, producing a ~130x speedup without an index.
Programmers disagree on a lot of things, but there are a few things we can all rally around. One of those is that backwards-incompatible changes are annoying. Like, really annoying.
A mea culpa: in River v0.39.0 , I shipped a small backwards-incompatible change in a minor version, which is something you shouldn't do. It added an options parameter to the migrator's Validate function:
res, err := migrator.Validate(ctx) res, err := migrator.Validate(ctx, nil) It's a small change that brings Validate into better alignment with the rest of the migrator API. It also affects only a relatively obscure function (most installs don't need Validate ), and River is still technically pre-1.0, which is how I justified it as it was going out the door, but it's bad practice and I acknowledge that.
These days, most of us do frequent automatic dependency bumps with Dependabot. When that weekly refresh comes through, all I want is to see a couple of versions change and green CI so I can hit the merge button without thinking about it much. If I instead get a broken build due to an API change in one of those dependencies, the annoyance I feel is irrational. A small change is better than a big change, but the irritation isn't proportional to the size of the diff. It could be a one-letter fix and still irk me.
Migrations, a special kind of pain
Migrations in a dependency are like backwards-incompatible changes, but ten times worse.
We're careful not to put major operational liabilities like full table locks into River migrations, but we still want to surface every migration to users because even a simple UPDATE on a large table can potentially be a long-running operation that puts undesirable load on a database that's running hot.
A commonly found operation in a migration is CREATE INDEX . In production, you always want to create indexes CONCURRENTLY to avoid blocking writes (otherwise, Postgres needs a SHARE lock ). But River's migration runner uses transactions as a matter of course, and Postgres doesn't allow CREATE INDEX CONCURRENTLY in a transaction, so our upgrade notes always include two paths:
The standard migrator invocation ( river migrate-up ... ), for development or after the expensive production changes have been applied safely by hand.
A list of manual CREATE INDEX CONCURRENTLY statements that can alternatively be run in hot environments.
It'd be an understatement to say that this adds friction to the upgrade process. It can turn what should have been a routine dependency upgrade into a multi-hour production operation, and may delay the upgrade by weeks or months as the can is kicked down the road and disappears into someone's backlog.
Worse yet, it produces a phenomenon that I refer to as migration fatigue . If a dependency is too painful to upgrade too many times, users will tire of it and complain. If that pain continues in spite of those complaints, they'll move to something else.
To combat migration fatigue, we try to live by a couple of principles:
Ship as few migrations as possible. The ideal number of migrations per year is zero.
If migrations are necessary, pool them for as long as possible so that everything needed for a given period can ship together.
It's easier said than done. Usually, when you need a migration, you need a migration, as in a new feature isn't tenable without it. However, that's not always the case.
We recently shipped active job rescue , which uses producer heartbeats to recover jobs orphaned by crashed clients much sooner than River's normal timeout-based rescue.
As the project neared the finish line, benchmarking surfaced something concerning. The rescue query was slow at large data sets: with 100,000 running jobs and 1,000 producers, a pass over healthy jobs took 8.9 seconds while finding nothing to rescue. Compared to the same operation before active rescue came in, this was a ~2,000x increase in query time.
The query was about as conventional as it gets. Here are just the key conditions for brevity, but imagine a river_job table containing all running jobs and a river_producer table tracking active clients and the queues they're working. This fragment checked for jobs still marked running but with no active producer:
EXISTS ( SELECT 1 FROM river_producer WHERE client_id = job_producer AND queue_name = job.queue AND created_at <= job.attempted_at AND updated_at < stale_cutoff ) AND NOT EXISTS ( SELECT 1 FROM river_producer WHERE client_id = job_producer AND queue_name = job.queue AND created_at <= job.attempted_at AND updated_at >= stale_cutoff ) = stale_cutoff)"> With J jobs and P producers, the number of producer-row checks was J × P , potentially an enormous number for large data sets, which is why we were seeing unacceptable query times around 10 seconds.
river_producer was indexed on its primary key only, so the most obvious mitigation was to make each of the J × P scans faster by adding an index:
CREATE INDEX river_producer_client_queue_created_at_idx ON river_producer (client_id, queue_name, created_at); That's a conventional solution, but it brings us right back to migration fatigue. We'd recently shipped workflows V2 , which included a substantial migration. Shipping another one so soon would have been far from ideal, but not shipping one would mean holding back the new feature.
A year ago, I think there's a reasonable chance we would have added the index and moved on. There's always more to do, so you have to weigh how time might be spent building your next feature instead of continuing to chase down the current rabbit hole.
But that was a year ago. One of the neat things about the LLM era is that we can let Codex simmer on a problem like this for a while and see whether we missed anything. After tens of minutes spent on investigation and iteration, it produced a different approach built around a materialized CTE:
producer_status AS MATERIALIZED ( SELECT client_id, queue_name, min (created_at) FILTER ( WHERE updated_at < stale_cutoff ) AS stale_created_at, min (created_at) FILTER ( WHERE updated_at >= stale_cutoff ) AS active_created_at FROM river_producer GROUP BY client_id, queue_name ) = stale_cutoff ) AS active_created_at FROM river_producer GROUP BY client_id, queue_name)"> Instead of scanning producer status separately for every job row, the CTE scans the producer table once per rescue pass and summarizes it by (client_id, queue_name) . The rescue query then joins that small materialized result against the jobs under consideration. Thousands of repeated scans collapse into one, producing a 12x to 130x speedup compared to our first, unoptimized pass:
The CTE needs to iterate over river_producer only once, so the proposed index above is no longer necessary. We get all the performance with none of the cost.
On the user's end, this is the ideal result: no backwards-incompatible changes and no new migrations. They update River Pro to a new version, and active job rescue activates automatically with no additional work involved.
The most obvious byproduct of the new LLM age has been more code, but less discussed is how much easier LLMs have made activities adjacent to coding too. Benchmarks used to take a full day (or multiple days) to carefully construct, run each branch, correct problems, run again, and finally tabulate and summarize results. Now, it's 20 minutes.
A single approach to solving a problem was quite expensive to develop, and you'd have meta problems with that too — if someone had spent days/weeks on it, they'd be more likely to advocate its continued existence even when a better version comes up. Now, a single approach is fast to generate, and egos are a negligible slice. It's hard to get attached to something when the LLM did all the work.
These days we're not only producing more software, but better software too.
© 2026 River Software, LLC. All rights reserved.

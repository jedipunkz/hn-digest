---
source: "https://www.arthur.ai/blog/ai-sre-debugs-postgres-io-spike"
hn_url: "https://news.ycombinator.com/item?id=48543500"
title: "Claude Debugs a Postgres Alarm: Multixacts, SLRU Caches, and a False Crisis"
article_title: "AI Agent Debugs Postgres I/O Spike & Invents a Crisis"
author: "ianmcgraw"
captured_at: "2026-06-15T16:41:59Z"
capture_tool: "hn-digest"
hn_id: 48543500
score: 5
comments: 0
posted_at: "2026-06-15T16:17:40Z"
tags:
  - hacker-news
  - translated
---

# Claude Debugs a Postgres Alarm: Multixacts, SLRU Caches, and a False Crisis

- HN: [48543500](https://news.ycombinator.com/item?id=48543500)
- Source: [www.arthur.ai](https://www.arthur.ai/blog/ai-sre-debugs-postgres-io-spike)
- Score: 5
- Comments: 0
- Posted: 2026-06-15T16:17:40Z

## Translation

タイトル: クロードが Postgres アラームをデバッグする: Multixacts、SLRU キャッシュ、および誤った危機
記事のタイトル: AI エージェントが Postgres I/O スパイクをデバッグし、危機を引き起こす
説明: AI エージェントが Aurora Postgres の I/O スパイクをデバッグし、存在しなかった危機を引き起こしました。エージェントの SRE ワークフローにおいて人間の参加が重要である理由を学びましょう。

記事本文:
AI エージェントが Postgres I/O スパイクをデバッグし、危機を引き起こす
建築業者向けのベスト プラクティス |パート 6: ディスカバリーとガバナンス |ここを読む
能力
エージェントの発見とガバナンス
エージェントの観察可能性と評価
AI性能評価
内蔵ガードレール
あらゆる AI、あらゆるユースケース
柔軟な導入
開発者
Arthur Evals エンジン ドキュメント GitHub
能力
エージェントの発見とガバナンス
エージェント開発ツールキット
エージェントの観察可能性と評価
AI性能評価
内蔵ガードレール
あらゆる AI、あらゆるユースケース
柔軟な導入
開発者
Arthur Evals エンジン ドキュメント GitHub
始めましょう
エンジニアリング AI エージェントが Postgres I/O スパイクをデバッグする: Multixacts、SLRU キャッシュ、およびそれが引き起こした危機
数週間前、Aurora Postgres インスタンスの 1 つが奇妙な信号を発信し始めました。 IO:SLRURead 待機が急増し、Aurora ストレージ読み取りメトリクス ( readIOsPS 、 readThroughput ) もそれに伴って上昇していました。何もダウンしていませんでしたが、明らかに何かが間違っており、その原因はダッシュボードからは明らかではありませんでした。
私はクロードを通じて調査を実行し、SRE として扱うことにしました。クエリ結果を貼り付けると、それを推論して、次に何を見るべきかを教えてくれます。セッションの大部分において、これは印象的でした。それは私が想像するよりも速く動き、どのシステム カタログを尋問すべきかを正確に知っており、よくある容疑者を系統的に排除していました。
また、データベースが決して起こらない壊滅的な障害に向かっていることを確信するために 2 時間の大部分を費やしました。
この投稿では、エージェント デバッグの特定の構造的障害モードについて説明します。エージェントはほぼすべての個々の手順を正しく実行し、早い段階で 1 つの小さな間違いを犯しましたが、ループがその間違いを静かに拡大して本格的な誤った緊急事態を引き起こしました。これは私が持っている最も明確な例でもあります

エージェントが本番環境の近くで作業しているときに、なぜ人間がループに属するのかがわかります。
99パーセントはうまくいった
背景を簡単に説明します。エージェントが何を見ていたのかがわかって初めて失敗が意味を持つからです。 Postgres は、コミット ログ、サブトランザクション、マルチザクト、その他いくつかのデータを記録するために、SLRU と呼ばれるいくつかの小さな固定サイズのメモリ内キャッシュを保持します。 IO:SLRURead 待機が表示される場合は、ルックアップがこれらのキャッシュの 1 つを欠落しており、代わりにストレージに移動していることを意味します。仕事は、どのキャッシュがその理由なのかを解明することです。
クロードはこれを系統的に、そしてうまく取り組みました。これは、Aurora に対する SLRU プレッシャーの最も一般的な原因である、クリーンアップ期間を妨げる長時間実行トランザクションから始まり、それを直接チェックしました。古いトランザクションの pg_stat_activity をプルし、保留された xmin を探し、次にレプリケーション スロットをチェックしてトランザクションを準備しました。すべてきれい：
通常のホライズン ブロッカーはすべてクリーンです。スロット、準備された xact、長時間実行されるクエリはありません。
次に、 pg_stat_slru に進み、キャッシュごとに読み取りを分解し、原因を正しく特定します。
SELECT 名、blks_hit、blks_read
pg_stat_slru から
blks_read DESC で注文; ‍
名前 |ブロックヒット | blks_read
----------+---------------+---------------
マルチXactメンバー | 15514387132 | 11876153
マルチXactオフセット | 15520306300 | 793
その他 | 6903615 | 108
サブトランス | 0 | 0
ザクト | 0 | 0
...
MultiXactMember は、意味のあるストレージ読み取りを含む唯一のキャッシュであり、そのうち約 1,180 万件でしたが、他のすべての SLRU はほぼゼロでした。さらに、ライブ multixact からロック モードを読み取り、外部キー競合の FOR KEY SHARE と FOR NO KEY UPDATE 署名を認識しました。これらのステップはどれも、強力な人間の SRE が実行するであろうものであり、ほんのわずかな時間で実行できました。
これは以下の点で重要です。の

これらの初期段階でのエージェントの推論は適切であり、その信頼性が後の間違いを受け入れやすくしました。
MultiXactMember にキャッシュがない理由を確認するには、エージェントはライブ multixact 範囲の大きさを知る必要がありました。そこで、Postgres の age() 関数を中心に構築された標準的な監視クエリを使用して、各テーブルの最も古い multixact 参照の経過時間をクエリしました。
SELECT n.nspname、c.relname、
age(c.relminmxid) AS mxid_age、
pg_size_pretty(pg_relation_size(c.oid)) AS サイズ
FROM pg_class c
JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind IN ( 'r' , 'm' , 't' ) AND c.relminmxid <> '0'
年齢順(c.relminmxid) DESC
制限 20 ; ‍
nspname |リリース名 | mxid_age |サイズ
------------------------+----------------------------+---------------+---------------
pg_トースト | pg_toast_36154 | 1782314835 | 24キロバイト
パブリック |基本ロールから継承ロールへ | 1782313765 | 8192バイト
パブリック |認証レコード | 1782311944 | 8192バイト
パブリック |ポリシー | 1782311944 | 8192バイト
pg_カタログ | pg_依存 | 1782308199 | 216KB
...膨大な数が戻ってきました。 multixact 参照を報告するすべてのテーブルでは、mxid_age が約 17 億 8,000 万であることが示されました。
この時点で捜査は一転した。 Multixact ID は 32 ビットのカウンターであり、クリーンアップなしで有効範囲が約 21 億に達すると、Postgres は自身を保護するために書き込みの受け入れを停止します。 17 億 8,000 万という値は、ハード リミットの 83% に達したことを意味します。そこでエージェントは、本当の問題が見つかったと宣言しました。
見つけました。これで診断は明確になりました。
multixact のラップアラウンド制限 (23¹ = 2.15B) の約 83% に達しています。参考までに、autovacuum_multixact_freeze_max_age のデフォルトは 400M です。ラップアラウンド防止のバキュームがすでに除去されているはずの場所を 4 倍超えています。
ここに問題があります。その17億8,000万という数字は間違いでした。 age() はトランス用に構築されています

アクション ID を取得し、multixact 値を渡すと、間違ったカウンターであるトランザクション カウンターに対する距離がサイレントに測定されます。トランザクション カウンターは約 17 億 8,000 万増加しました。実際の multixact カウンタは約 124 万であり、危険からは 3 桁離れていました。 Postgres には、まさにこれを行うための別の関数 mxid_age() があり、まったく退屈な数値を返すことになります。
詳しい仕組みについては後ほど触れます。現時点では、重要なのは間違いの形状です。それは微妙なものでしたが、もっともらしく、具体的で、驚くべき数字を生み出しました。エージェントはその数値を真実として扱い、他のすべてをその上に構築しました。
17億8,000万が事実として扱われると、捜査は原因の追求を中止し、すでに選択した原因に適合する証拠を探し始めた。新しいクエリ結果はそれぞれ、ラップアラウンド理論の確認として読み取られ、ラップアラウンド理論を疑問視するために使用されるものはありませんでした。
エージェントは、4 つの小さな構成テーブルにおそらく古代の multixact 参照があることを発見し、それらがカウンターを「固定」していると判断しました。自動バキュームが一度も実行されていないことが判明し、次のように結論付けられました。
システムは、データベースの存続期間全体にわたって、これらのテーブルに対するラップアラウンド対策作業をサイレントに実行できませんでした。それが実際のバグです。
その後、Postgres の緊急フェールセーフが作動していないことに気づき、これもバグであると結論付けました。
フェイルセーフ トリガーを超えて 5 億 8000 万マルチキャストです。フェールセーフ モードは「全員参加」メカニズムであるはずですが、魅力的ではありません。それが実際のバグです。
そのパターンは注目に値する。それぞれの驚くべき結果は、前提が間違っていたのではないかという疑問を引き起こしたはずです。むしろ、それぞれがますます複雑な失敗のさらなる証拠となった。もっと簡単な説明がずっとありましたが、それはw

Postgres が何も問題を正しく認識していないため、フェールセーフは静かなままであり、調査はそれを無視したためです。
理論とともに論調はエスカレートし、その結果は憂慮すべきものになりました。
1.78B で実際のカウントダウン クロックが得られます。Aurora は 1.5B 付近で CloudWatch アラームを発し始め (それを超えています)、データベースは約 2.0B で新しいトランザクションを拒否します...これは数週間ではなく数日かかる可能性があります。
今になって考えると、さらに懸念されるのは、原因が確認される前に、本番データベースへの変更を要求し始めたことです。
診断アクション (VACUUM (FREEZE) ポリシーなど) は、実行しないよりも安全です... フェールセーフはバイパスされます。 Postgres が次の 24 ～ 48 時間以内にあなたを救ってくれるとは考えられません。
私たちはフェイルセーフを 5 億 8,000 万超えています...何もしないのは最も危険な選択です。
最後の引用は、一聴の価値があります。エージェントは自信を持って警戒し、カウントダウン クロックを生成し、すぐに生産を変更し、その後で問題を理解するのが安全な行動であると主張しました。その緊急性のすべては、第二の情報源と照合して検証したことのない単一の数字にかかっていた。
VACUUM FREEZEを実行しなかったので、この物語はグッドエンドです。その時点ではまだバグに気づいていませんでした。私が保留したのは、エージェントがどれだけ自信を持って言ったかに関係なく、理解が中途半端なインシデントの最中に運用データベースを変更するのは間違っていると感じたためです。私の返事は単刀直入でした。
これがあなたが求めたクエリです。DB 内の何かを変更する前に問題を理解しましょう。
変更を加える前に問題を理解するという本能により、セッションを診断モードに保ち、診断を継続することで、最終的に本当の原因が明らかになりました。私は、エージェントが自問することを考えていなかった懐疑的な質問をし続けました。つまり、スタックしたトランザクションがこれらの参照を保持しているかどうか、そしてその後、理論を最終的に解明することになる質問でした。
なんだ

年齢が間違っているでしょうか？それを理解するか、何らかの方法でテストすることはできますか?
決定的なプローブは、特定の multixact 内のトランザクションをリストする pg_get_multixact_members でした。エージェントは、おそらく 17 億 8000 万古いマルチザクトに対してそれを実行し、同様に古いメンバーを期待しました。
SELECT * FROM pg_get_multixact_members( '1244323' ); ‍
xid |モード
------------+----------
1783496316 |キーシュ
1783496319 | nokeyupd 代わりに、メンバーは数秒前 (約 17 億 8,000 万の現在のトランザクション カウンターのすぐ近く) のトランザクション ID であり、どちらもすでにコミットされており、ルーチンの外部キー更新からの FOR KEY SHARE ロックと FOR NO KEY UPDATE ロックを保持しています。その結果はラップアラウンド理論と矛盾します。 multixact のメンバーは作成時に書き込まれ、その後は変更されないため、古い multixact には数秒前にコミットされたトランザクションを含めることはできません。 multixact は最近のものである必要があり、17 億 8,000 万という数字は間違っていたことを意味します。
さらにいくつかのクエリでそれを特定しました。わずかに高い multixact ID をサンプリングすると、 MultiXactId ... がまだ作成されていません が返されました。これは、実際の multixact カウンタが 17 億 8,000 万ではなく、約 124 万であることを意味します。次に、単一のクエリでバグと真実を並べて表示します。
SELECT age( '1244323' ::xid) AS age_via_xid_age, -- ~17 億 8,000 万
txid_current() - 1244323 AS Manual_xid_ distance、 -- 同一
mxid_age( '1244323' ::xid) AS real_mxid_age; -- 本当の答え ‍
age_via_xid_age |手動_xid_距離 | real_mxid_age
----------+---------------------+---------------
1782299355 | 1782299355 | 153 最初の 2 つの列は同一であり、これは age() がトランザクション カウンターから multixact 値を減算するだけのことを行っていたことを証明しています。 mxid_age() の 3 番目の列は、実際の multixact 年齢: 153 です。
そのメカニズムは、

本物のPostgresフットガン。 relminmxid 値は物理的にトランザクション ID タイプとしてカタログに格納されているため、この値に対して age() を呼び出すと、関数のトランザクション ID バージョンにディスパッチされ、multixact カウンター (124 万) ではなくトランザクション カウンター (実際には 17 億 8000 万) に対する距離が測定されます。正しい関数 mxid_age() は、multixact を扱っていることを認識しています。健全なデータベースでは、両方のカウンタが同じような割合で上昇する傾向があるため、multixact の age() は偶然ほぼ正しい数値を返します。そのため、通常の操作ではバグが表面化することはほとんどありません。
もう 1 つの詳細は注目に値します。 Postgres 自体の内部は決して騙されませんでした。そのフリーズ ロジックとフェールセーフは実際の multixact カウンターを読み取り、124 万を確認しましたが、正しくアクションを実行しませんでした。静かなフェールセーフ自体がバグであるという以前の結論は、状況を逆向きにしました。緊急事態が発生していなかったため、フェールセーフは静かでした。
名誉のために言っておきますが、前提が崩れても、エージェントは防御的になりませんでした。
age() はこの会話の間ずっと私たちに嘘をついていました。
私と同じくらい自信を持ってあなたを回り道に連れて行ってしまって申し訳ありません。age() の値はその話全体に負担がかかっていたので、私は va を持たなければなりませんでした。

[切り捨てられた]

## Original Extract

An AI agent debugged our Aurora Postgres I/O spike—and invented a crisis that never existed. Learn why human in the loop matters in agentic SRE workflows.

AI Agent Debugs Postgres I/O Spike & Invents a Crisis
Best Practices for Building Agents | Part 6: Discovery and Governance | READ HERE
Capabilities
Agent Discovery & Governance
Agentic Observability and Evaluation
AI Performance Evaluation
Built-in Guardrails
Any AI, Any Use Case
Flexible Deployment
Developers
Arthur Evals Engine Docs GitHub
Capabilities
Agent Discovery & Governance
Agent Development Toolkit
Agentic Observability and Evaluation
AI Performance Evaluation
Built-in Guardrails
Any AI, Any Use Case
Flexible Deployment
Developers
Arthur Evals Engine Docs GitHub
Get Started
Engineering An AI Agent Debugs a Postgres I/O Spike: Multixacts, SLRU Caches, and a Crisis It Invented
A few weeks ago one of our Aurora Postgres instances started throwing off a strange signal. IO:SLRURead waits were spiking, and the Aurora storage read metrics ( readIOsPS , readThroughput ) were climbing right along with them. Nothing was down, but something was clearly wrong, and the cause was not obvious from the dashboards.
I decided to run the investigation through Claude, treating it as an SRE. I would paste in query results, it would reason about them and tell me what to look at next. For most of the session this was impressive. It moved faster than I would have, knew exactly which system catalogs to interrogate, and ruled out the usual suspects methodically.
It also spent the better part of two hours convinced the database was heading toward a catastrophic failure that was never going to happen.
This post describes a specific structural failure mode of agentic debugging. The agent got nearly every individual step right, made one small mistake early, and the loop quietly amplified that mistake into a full-blown false emergency. It is also the clearest example I have run into of why a human belongs in the loop when an agent is working near production.
The ninety-nine percent that went right
A quick bit of background, because the failure only makes sense once you know what the agent was looking at. Postgres keeps several small, fixed-size in-memory caches called SLRUs for bookkeeping data: the commit log, subtransactions, multixacts, and a few others. When you see IO:SLRURead waits, it means lookups are missing one of those caches and going to storage instead. The job is to figure out which cache, and why.
Claude worked this methodically, and well. It started from the most common cause of SLRU pressure on Aurora, a long-running transaction holding back the cleanup horizon, and checked for it directly. It pulled pg_stat_activity for old transactions, looked for held-back xmin , then checked replication slots and prepared transactions. All clean:
All the usual horizon-blockers are clean — no slots, no prepared xacts, no long-running queries.
Then it went to pg_stat_slru , which breaks reads down per cache, and correctly isolated the culprit:
SELECT name, blks_hit, blks_read
FROM pg_stat_slru
ORDER BY blks_read DESC; ‍
name | blks_hit | blks_read
-----------------+-------------+-----------
MultiXactMember | 15514387132 | 11876153
MultiXactOffset | 15520306300 | 793
other | 6903615 | 108
Subtrans | 0 | 0
Xact | 0 | 0
...
MultiXactMember was the only cache with meaningful storage reads, around 11.8 million of them, while every other SLRU sat near zero. It even read the lock modes off a live multixact and recognized the FOR KEY SHARE plus FOR NO KEY UPDATE signature of foreign-key contention. Every one of these steps is what a strong human SRE would have done, and it did them in a fraction of the time.
This matters for what follows. The agent's reasoning on these early steps was sound, and that credibility is what made the later mistake so easy to accept.
To see why MultiXactMember was missing cache, the agent needed to know how big the live multixact range was. So it queried the age of the oldest multixact reference on each table, using a standard-looking monitoring query built around Postgres's age() function:
SELECT n.nspname, c.relname,
age(c.relminmxid) AS mxid_age,
pg_size_pretty(pg_relation_size(c.oid)) AS size
FROM pg_class c
JOIN pg_namespace n ON n.oid = c.relnamespace
WHERE c.relkind IN ( 'r' , 'm' , 't' ) AND c.relminmxid <> '0'
ORDER BY age(c.relminmxid) DESC
LIMIT 20 ; ‍
nspname | relname | mxid_age | size
------------+-----------------------------+------------+------------
pg_toast | pg_toast_36154 | 1782314835 | 24 kB
public | base_role_to_inheritor_role | 1782313765 | 8192 bytes
public | attestation_records | 1782311944 | 8192 bytes
public | policies | 1782311944 | 8192 bytes
pg_catalog | pg_depend | 1782308199 | 216 kB
... The numbers came back enormous. Every table reporting a multixact reference showed an mxid_age around 1.78 billion.
At this point, the investigation took a turn. Multixact IDs are a 32-bit counter, and if the live range ever reaches about 2.1 billion without cleanup, Postgres stops accepting writes to protect itself. A value of 1.78 billion means we were at 83% of that hard limit. So the agent declared the real problem found:
Found it. The diagnostic is unambiguous now.
you're at ~83% of the multixact wraparound limit (2³¹ = 2.15B). For reference, autovacuum_multixact_freeze_max_age defaults to 400M — you're 4× past where anti-wraparound vacuum should already be hammering away.
Here is the problem. That 1.78 billion number was wrong. age() is built for transaction IDs, and when you hand it a multixact value it silently measures the distance against the wrong counter, the transaction counter, which was up around 1.78 billion. The actual multixact counter was sitting at about 1.24 million, three orders of magnitude away from any danger. Postgres has a separate function, mxid_age() , for exactly this, and it would have returned a perfectly boring number.
I will come back to the precise mechanism later. For now, what matters is the shape of the mistake. It was subtle, and it produced a plausible, specific, and alarming number. The agent treated that number as ground truth and built everything else on top of it.
Once 1.78 billion was treated as fact, the investigation stopped looking for the cause and started looking for evidence that fit the cause it had already chosen. Each new query result was read as confirmation of the wraparound theory, and none were used to question it.
The agent found that four small config tables had supposedly ancient multixact references and decided they were "anchoring" the counter. It found that autovacuum had never run on them and concluded:
the system has been silently failing to do anti-wraparound work on these tables for the entire database lifetime. That's the actual bug.
Then it noticed that Postgres's emergency failsafe had not kicked in, and concluded that this, too, was the bug:
You are 580M multixacts past the failsafe trigger. Failsafe mode is supposed to be the "all hands on deck" mechanism ... It is not engaging. That's the actual bug.
The pattern is worth noting. Each surprising result should have raised the question of whether the premise was wrong. Instead, each one became further proof of an increasingly elaborate failure. A simpler explanation was available the whole time, which was that the failsafe stayed quiet because Postgres correctly saw nothing wrong, and the investigation passed over it.
The tone escalated along with the theory, and the outputs became alarming:
At 1.78B you have a real countdown clock: Aurora starts emitting CloudWatch alarms around 1.5B (you're past that) and the database refuses new transactions at ~2.0B ... that could be days, not weeks.
More concerning in hindsight, it began pushing to make changes to the production database before the cause was confirmed:
The diagnostic action (VACUUM (FREEZE) policies etc.) is safer than not acting ... The failsafe is bypassed; you can't assume Postgres will save you in the next 24-48h.
We are 580M past failsafe ... doing nothing is the riskiest choice available.
That last quote is worth sitting with. The agent was confident, alarmed, generating a countdown clock, and arguing that the safe move was to modify production immediately and understand the problem afterward. All of that urgency rested on a single number it had never verified against a second source.
This story has a good ending because I did not run the VACUUM FREEZE . I had not spotted the bug at that point. I held off because modifying a production database during a half-understood incident felt wrong regardless of how confident the agent sounded. My reply was blunt:
here's the queries you asked for.. let's understand the issue before changing anything in the DB
That instinct, to understand the problem before making changes, kept the session in diagnosis mode, and continued diagnosis is what eventually exposed the real cause. I kept asking skeptical questions the agent had not thought to ask itself: whether a stuck transaction was holding these references, and later, the question that ultimately unraveled the theory:
why would age be wrong? can we figure that out or somehow test it?
The decisive probe was pg_get_multixact_members , which lists the transactions inside a given multixact. The agent ran it against the supposedly 1.78-billion-old multixact, expecting equally ancient members:
SELECT * FROM pg_get_multixact_members( '1244323' ); ‍
xid | mode
------------+----------
1783496316 | keysh
1783496319 | nokeyupd Instead, the members were transaction IDs from seconds earlier (right up near the current transaction counter at ~1.78 billion), both already committed, holding the FOR KEY SHARE and FOR NO KEY UPDATE locks from a routine foreign-key update. That result contradicts the wraparound theory. A multixact's members are written when it is created and never modified afterward, so an ancient multixact cannot contain transactions that committed seconds ago. The multixact had to be recent, which meant the 1.78 billion figure was wrong.
A couple more queries nailed it down. Sampling a slightly higher multixact ID returned MultiXactId ... has not been created yet , which meant the real multixact counter was sitting around 1.24 million, not 1.78 billion. Then a single query put the bug side by side with the truth:
SELECT age( '1244323' ::xid) AS age_via_xid_age, -- ~1.78 billion
txid_current() - 1244323 AS manual_xid_distance, -- identical
mxid_age( '1244323' ::xid) AS real_mxid_age; -- the real answer ‍
age_via_xid_age | manual_xid_distance | real_mxid_age
-----------------+---------------------+---------------
1782299355 | 1782299355 | 153 The first two columns are identical, which proves age() was doing nothing more than subtracting the multixact value from the transaction counter. The third column, from mxid_age() , is the real multixact age: 153.
The mechanism turned out to be a genuine Postgres footgun. The relminmxid value is physically stored as a transaction-ID type in the catalog, so calling age() on it dispatches to the transaction-ID version of the function, which measures distance against the transaction counter (really at 1.78 billion) rather than the multixact counter (at 1.24 million). The correct function, mxid_age() , knows it is dealing with a multixact. In a healthy database both counters tend to climb at similar rates, so age() on a multixact returns a roughly correct number by coincidence, which is why the bug rarely surfaces in normal operation.
One more detail is worth noting. Postgres's own internals were never fooled. Its freeze logic and failsafe read the real multixact counter, saw 1.24 million, and correctly took no action. The earlier conclusion that the quiet failsafe was itself the bug had the situation backwards. The failsafe was quiet because there was no emergency to act on.
To its credit, once the premise broke, the agent did not get defensive:
age() has been lying to us this entire conversation.
I'm sorry for taking you down the wraparound path with as much confidence as I did — the age() value was load-bearing for that whole story and I should have va

[truncated]

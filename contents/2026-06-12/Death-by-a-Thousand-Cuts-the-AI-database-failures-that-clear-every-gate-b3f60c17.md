---
source: "https://explainanalyze.com/p/death-by-a-thousand-cuts-the-ai-database-failure-you-cant-restore/"
hn_url: "https://news.ycombinator.com/item?id=48503817"
title: "Death by a Thousand Cuts: the AI database failures that clear every gate"
article_title: "Death by a Thousand Cuts: the AI Database Failure You Can't Restore"
author: "rtolkachev"
captured_at: "2026-06-12T16:09:08Z"
capture_tool: "hn-digest"
hn_id: 48503817
score: 1
comments: 0
posted_at: "2026-06-12T13:25:08Z"
tags:
  - hacker-news
  - translated
---

# Death by a Thousand Cuts: the AI database failures that clear every gate

- HN: [48503817](https://news.ycombinator.com/item?id=48503817)
- Source: [explainanalyze.com](https://explainanalyze.com/p/death-by-a-thousand-cuts-the-ai-database-failure-you-cant-restore/)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T13:25:08Z

## Translation

タイトル: Death by a Thousand Cuts: すべてのゲートをクリアする AI データベースの障害
記事のタイトル: 千回の切断による死亡: 復元できない AI データベースの障害
説明: 2025 年 2 月、ワイオミング州の裁判所は、9 件の訴訟を引用した申し立てで 3 人の弁護士に制裁を下し、そのうち 8 件は当事務所がでっち上げたものだった

記事本文:
午前3時に呼び出されないシステムを構築し、それを維持するチーム。
ルスラン・トルカチョフ著
スタッフ システム エンジニア — データベースとデータ インフラ
DROP は幸運なケースであり、その理由を示す 2x2 が次のとおりです。
有効な例: 論理的な削除によるリーク
修正は技術的なものではなく、ビジネス上のものです
モニタリングの実際の目的
千回の切断による死亡: 復元できない AI データベースの障害
2025年2月、ワイオミング州の裁判所は9件の訴訟を引用した申し立てで3人の弁護士に制裁を与え、そのうち8件は同社のAIツールによってでっち上げられたものだった。草稿には午後かかった。清掃員が事件を処理した。裁判官が準備書面を読むため、数週間で表面化した。ほとんどのソフトウェアにはジャッジがありません。
2026年6月4日(木)
読了時間: 13 分 TL;DR 壊滅的な AI 障害 (テーブルを削除するエージェント) は回復可能なものです。大音量で原因が考えられますが、ゲートされたパイプラインでは、いずれにせよほとんどの場合起こりません。実際に驚かされるのは、すべてのゲートをクリアする変更です。これは、ゲートがレビュー時に正確さをチェックするためであり、失敗は量と時間の関数であり、PR が開いているときはどちらもゼロです。障害がうるさい場合と静かな場合、回復可能な場合とそうでない場合にプロットすると、パイプラインがこれまで監視していなかった静かで回復不可能な部分に AI が溢れます。スクレイパーは、他のすべてと同じパイプラインの背後で出荷されます。機能ブランチ、2 つの承認、CI グリーン、ステージングで 1 日、その後デプロイされます。すでに投稿テーブルがあり、ジョブごとに 1 行がクローラー追跡を投稿し、キーは posting_id です。変更の一部は、クロール状態を追跡するための新しいテーブルで、各実行で確認されたすべての投稿のコンテンツ ハッシュを保存して、次の実行で何が変更されたかを認識できるようにします。エージェントがテーブルを設計し、レビュー担当者がそれを承認しましたが、レビュー時にはテーブルに保持されている行はゼロでした。それは合理的であるように見えました:
1
2
3
4
5
6
7
8
CREATE TABLE クロール状態 (
ID BIGSERIAL プライマリ

キー、
run_id BIGINT NOT NULL 、
posting_id BIGINT NOT NULL REFERENCES 投稿 (posting_id )、
content_hash CHAR ( 64 ) NOT NULL 、 -- 投稿本文の SHA-256
現在、crawled_at TIMESTAMPTZ NOT NULL DEFAULT ()
);
クロール状態 (posting_id ) にインデックスを作成します。
3ヶ月くらいなら妥当かな。
その画面には何も問題はないようですが、それが問題です。サロゲート ID 、 投稿への外部キー 、投稿を検索する列のインデックス。定型文としてレビューします。粒度は、誰も大声で言わない部分です。実行ごとに投稿ごとに 1 行です。すべてのパスでは、既存の行を更新するのではなく、新しい run_id の下にポスティング セット全体が追加されるため、スクレイパーが実行されるたびに、テーブルは完全なセットだけ増加します。投稿には約 100,000 行が含まれます。 3 か月と百数回の実行後、crawl_state には 1,100 万が保持されます。投稿が変更されたかどうかを決定するジョブは、明らかなクエリを実行します。
1
2
3
4
5
コンテンツハッシュを選択します
FROM クロール状態
WHERE 投稿ID = $1
ORDER BY クローリングされた_DESC
リミット1;
posting_id インデックスは一致する行を検索しますが、現在は投稿ごとに 100 数行あり、単一の最新ハッシュを返すために、呼び出しのたびに、crawled_at によって山積みされたものを並べ替えます。ワーキング セット全体でパイプラインがタイムアウトになります。遅さを修正するよう求められたエージェントは、常に推奨していることを推奨します。つまり、インデックスを CREATE INDEX ONcrawl_state (posting_id,crawled_at DESC) に拡張することで、最新ハッシュの検索がソートを停止します。
システムを理解しているエンジニアは、システムの読み方が異なります。テーブルには木目がまったくあってはならない。投稿の現在のハッシュは 1 つの値で、一度保存され、実行のたびに上書きされ、蓄積されるものはありません。さらに良いことに、ハッシュは投稿とは別の問題ではありません。すでに存在する投稿行に属しているため、秒はありません

ond テーブル、 run_id なし、投稿ごとの最新ルックアップはまったくありません。
1
2
3
4
5
6
7
8
ALTER TABLE の投稿
ADD COLUMN content_hash CHAR ( 64 )、
列を追加 hash_checked_at TIMESTAMPTZ ;
-- 各実行、投稿ごと:
投稿を更新する
SET content_hash = $ 2 、 hash_checked_at = now ()
WHERE posted_id = $1 AND content_hash IS DISTINCT FROM $2 ;
これは永久に 100,000 行のままであり、変更されたかどうかのチェックは 1 つの主キー行の読み取りだけで行われ、肥大化するcrawl_state はありません。エージェントが提案したインデックスは、実際のバグである粒度を倍増させながら、存在すべきではないテーブルへの書き込みコストとストレージを支払って症状をスピードアップします。また、元に戻すためのデプロイもありません。不適切なグレインは、3 か月前のスキーマ決定と 1,100 万行の蓄積された状態であり、それを巻き戻すことは計画的な移行とバックフィルであり、他の変更と同様にレビューされ、ロールバックではありません。
DROP は幸運なケースであり、その理由を示す 2x2 が次のとおりです。
誰もがイメージする失敗は破壊的なものです。 DROP TABLE 、間違った関係を切り捨てる移行、バイナリログを削除するスクリプトです。あなたが主に扱ってきたゲート システムでは、破壊的な移行がレビューで捕らえられ、資格情報のスコープが設定され、復元ランブックが実践されています。 2025 年 7 月のコード フリーズ中に Replit エージェントが運用データベースを消去したとき、データは復元されました。大音量で、1 つのタイムスタンプに起因し、回復可能でした。スクレイパーは、マージ時に同じパイプラインをクリアしました。これは、グレイン バグが表現する行がなく、CI シードが中断する行数に達しなかったために、マージ時にそれが正しく行われたためです。
故障を 2 つの軸 (騒音か静かか、回復可能かそうでないか) でランク付けすると、次の 4 つのコーナーが得られます。
大音量で回復可能なのはDROPです。すぐにわかり、ロールバックします。
うるさくて回復不能になることは稀です: 破壊的です

操作はキャッチしても元に戻すことはできません。
静かで回復可能なバグは、気付かれずに放置されていましたが、見つけたときにまだ元に戻すことができるバグです。
静かで回復不可能な廃墟地区。何ヶ月も経っても気づかず、その頃には以前の状態は消えているか、きれいなアーティファクトとしては存在しません。
軸が相関しているため、悪いコーナーが深くなります。大音量の障害は、以前の状態がまだ存在している間に捕捉されます。静かなものは放置され、長く放置されるほど、下流システムはそれを真実として消費し、回復可能なエラーが集約されて回復不可能なエラーに伝播されます。大声で時間を稼ぐので、静かで取り返しのつかない旅を一緒にしましょう。
構造的な理由から、AI が不良コーナーに押し寄せます。配布される各変更には妥当性があります。コンパイル、実行、正しい形状を返し、存在するあらゆるチェックに合格します。そのもっともらしさは破損フロアであり、出力を有用なものにするのと同じメカニズムであり、まったく正しく見える方法で出力を間違ってしまうことがあります。大きな失敗とは、出力が妥当性を示すことができなかった失敗です。静かなものは、設計どおりに動作するモデルです。そして、ボリュームはそれを倍増します。チームは、同じレビューと CI 予算で、10 倍の頻度で満たされる 8 つのサンプルの代わりに、週に 80 の変更を出荷します。 DROP は、パイプラインが停止するために構築されたまれな引き金です。千のカットはモーダルなドローであり、それを波のように動かします。
クワイエットコーナーには 3 つの形状があります。スクレイパーは 1 つ目です。スキーマは 0 行では正しく、1,100 万行では致命的になります。これは、モデルがスケールではなくトレーニング分布に基づいて最適化されるためです。大規模なテーブルをパーティション分割するよう要求された場合、大規模な OLTP テーブルに適合する主キー パーティション分割ではなく、主キー (共通コーパス形状) の created_at に到達します。 2 番目は、domai を保持していないため、間違って計算された値です。

n: Cursor のサポート ボットが存在しないログイン ポリシーを発明し、ユーザーはいつの間にかキャンセルし、エア カナダはチャットボットがでっちあげた死別返金をめぐる裁判で敗訴しました。同じジェネレーターを割引または税分割を計算する書き込みパスに移動すると、行は適切に型指定されていますが、数値の意味が間違っており、四半期末まで契約と照合するものは何もありません。 3 つ目は、作成者自身の理解を超えて出荷された変更です。見た目は良く、うまく機能したため、変更が加えられました。変更をキャッチする判断は、エージェントが現在スキップしている遅い作業によって構築されます。これが速いエンジニアの矛盾であり、2025 年 7 月の調査では、開発者が速く感じていたにもかかわらず、実際の速度は 19% 遅かったと測定されました。
有効な例: 論理的な削除によるリーク
穀物のバグは、探しに行くとうるさくなります。それは遅延として現れるためです。同じクラスのさらに悪いバージョンは、数値を破壊し、パフォーマンス メトリックをまったく動かさない書き込みです。論理的な削除は、ほとんどのエンタープライズ スキーマに存在します。
この慣例は古く、不文律です。行が物理的に削除されることはありません。 it gets a deleted_at stamp, and every query that reads the table is expected to filter it out. A billing system for a SaaS company looks like this:
1
2
3
4
5
6
7
8
-- one row per recurring line on an account: base plan, seats, add-ons
CREATE TABLE subscription_items (
id BIGSERIAL 主キー 、
account_id BIGINT NOT NULL REFERENCES accounts ( account_id ),
sku テキストが NULL ではありません 、
month_cents BIGINT NOT NULL 、
delete_at TIMESTAMPTZ -- 顧客が回線を切断したときに設定されます。
);
When a customer downgrades, the application does not delete the row, it stamps it:
1
UPDATE subscription_items SET deleted_at = now () WHERE id = $ 1 ;
Every existing query that touches money knows this. MRR ロールアップ、請求書ジェネレータ

tor、収益ダッシュボードでは、すべて AND delete_at IS NULL が表示されます。これを忘れると、チャーン収益が 2 倍にカウントされることをチームが何年も前に学んだためです。その知識はクエリとクエリを作成した人の頭の中に生きています。これはスキーマのどこにもありません。 delete_at は単なる null 許容のタイムスタンプであり、クエリがこれを無視することを妨げるものはありません。
ここでエージェントは、ダッシュボードが読み取るウェアハウス ファクト テーブルに夜間の ETL が永続化される、取締役会向けの指標である地域別の月次経常収益を追加するように求められます。明白なことが書かれています:
1
2
3
4
5
6
ビューの作成 mrr_by_region AS
a を選択します。地域、SUM (si .monthly_cents) AS mrr_cents
アカウントから
subscription_items si ON si を結合します。 account_id = a 。アカウントID
ここで、 . delete_at IS NULL -- アカウントに記憶されています
グループ化地域 ;
スキーマにはフィルタリングする必要があると記載されておらず、トレーニング コーパスはこれとまったく同じ形状の結合でいっぱいであるため、 accounts では delete_at がフィルタリングされましたが、 subscription_items ではフィルタリングされませんでした。キャンセルされたアドオンとダウングレードされたシートはすべて地域の MRR に合計され、ETL は毎晩このビューを読み取り、水増しされた数値をウェアハウスに書き込み、全社が真実として読み取ります。形状は正しく、数値は妥当であり、少し高く、論理的に削除されたプールが大きくなるにつれて増加しています。
何も引っかかりません。エンジニアは、作業中のビューと、そこに存在する delete_at フィルターを見て、次に進みました。 AI レビュー担当者が命名規則に問題があるとフラグを立てました。人間は緑色の概要をざっと読んで承認しました。 CI は削除された行がほとんどないシード データベース上で実行されたため、テストでのリークは丸め誤差でした。述語の欠落は不在であり、人間であれモデルであれ、すべてのレビュー担当者が確実に見落としている唯一のものです。
ドリフトこそが物語だ。発売当初は小さかったが、1年後の各地域では10～15パーセント高くなっていた

これはダウングレードの歴史の中で最も多く、財務部門はこれが誰もが行う唯一の方法であると認識しています。それは、単一の原因もデプロイメントも元に戻すこともなく、1 年後にダッシュボードと請求収益を照合することです。
修正は技術的なものではなく、ビジネス上のものです
緩和策は知られていますが、どれも賢いものではありません。ビジネスが耐えられるペースで真実の情報源と照合し、エラーのみではなく集計とドリフトを警告し、運用形態のボリュームに対して CI を実行し、不変条件をビューと制約に焼き付けます。 All worth doing, all secondary, because every one is a net thrown after the write has committed. The load-bearing decision is upstream of the tooling, and it is a positioning call leadership makes on purpose. 4 つの正直な立場:
全速力で船を出し、コーナーを受け入れます。 Take the velocity and take the unrecoverable write, the silent data loss, the bug the customer finds, as the cost. Legitimate for a seed-stage product with no prior state worth protecting.課金システムにとっては壊滅的。
安いところは速く、そうでないところはゲート付きです。 Agents and junior engineers run on loud, recoverable surfaces (internal tooling, dashboards, throwaway analysis); writes that touch money, schema, or multi-writer tables go through someone who holds th

[切り捨てられた]

## Original Extract

In February 2025 a Wyoming court sanctioned three lawyers for a motion citing nine cases, eight of them invented by the firm

Building systems that don't page you at 3am — and teams that keep it that way.
by Ruslan Tolkachev
Staff Systems Engineer — databases & data infra
The DROP is the lucky case, and here’s the 2x2 that shows why
A worked example: the soft-delete leak
The fix is a business call, not a technical one
What your monitoring is actually for
Death by a Thousand Cuts: the AI Database Failure You Can't Restore
In February 2025 a Wyoming court sanctioned three lawyers for a motion citing nine cases, eight of them invented by the firm's AI tool. The draft took an afternoon; the cleanup took the case. A judge reads briefs, so it surfaced in weeks. Most software has no judge.
Thursday, June 4, 2026
13 minutes read TL;DR The catastrophic AI failure (the agent that DROPs a table) is the recoverable one: loud, attributable, and on a gated pipeline it mostly can’t happen anyway. What actually bleeds you is the change that clears every gate because the gates check correctness at review time, and the failure is a function of volume and time, both of which are zero when the PR is open. Plot failures on loud-vs-quiet and recoverable-vs-not, and AI floods the quiet-and-unrecoverable corner the pipeline was never watching. A scraper ships behind the same pipeline as everything else: feature branch, two approvals, CI green, a day in staging, then deploy. There is already a postings table, one row per job posting the crawler tracks, keyed by posting_id . Part of the change is a new table to track crawl state, storing the content hash of every posting each run sees so the next run can tell what changed. The agent designed the table, a reviewer approved it, and at review time it held zero rows. It looked reasonable:
1
2
3
4
5
6
7
8
CREATE TABLE crawl_state (
id BIGSERIAL PRIMARY KEY ,
run_id BIGINT NOT NULL ,
posting_id BIGINT NOT NULL REFERENCES postings ( posting_id ),
content_hash CHAR ( 64 ) NOT NULL , -- SHA-256 of the posting body
crawled_at TIMESTAMPTZ NOT NULL DEFAULT now ()
);
CREATE INDEX ON crawl_state ( posting_id );
It was reasonable, for about three months.
Nothing on that screen looks wrong, and that is the problem. A surrogate id , a foreign key to postings , an index on the column you look postings up by. It reviews as boilerplate. The grain is the part nobody states out loud: one row per posting per run . Every pass appends the entire posting set under a fresh run_id instead of updating the rows already there, so the table grows by the full set every time the scraper runs. postings holds roughly 100,000 rows. Three months and a hundred-odd runs later, crawl_state holds 11 million. The job that decides whether a posting changed runs the obvious query:
1
2
3
4
5
SELECT content_hash
FROM crawl_state
WHERE posting_id = $ 1
ORDER BY crawled_at DESC
LIMIT 1 ;
The posting_id index finds the matching rows, but there are now a hundred-odd of them per posting, and to return the single latest hash it sorts that pile by crawled_at on every call. Across the working set the pipeline is timing out. Asked to fix the slowness, the agent recommends what it always recommends: widen the index to CREATE INDEX ON crawl_state (posting_id, crawled_at DESC) , so the latest-hash lookup stops sorting.
An engineer who knows the system reads it differently. The table shouldn’t have run grain at all. A posting’s current hash is one value, stored once and overwritten each run, and there is nothing to accumulate. Better than that, the hash isn’t a separate concern from the posting. It belongs on the postings row that already exists, so there is no second table, no run_id , and no latest-per-posting lookup at all:
1
2
3
4
5
6
7
8
ALTER TABLE postings
ADD COLUMN content_hash CHAR ( 64 ),
ADD COLUMN hash_checked_at TIMESTAMPTZ ;
-- each run, per posting:
UPDATE postings
SET content_hash = $ 2 , hash_checked_at = now ()
WHERE posting_id = $ 1 AND content_hash IS DISTINCT FROM $ 2 ;
That stays at 100k rows forever, the changed-or-not check is a single primary-key row read, and there is no crawl_state to bloat. The index the agent suggested speeds the symptom while doubling down on the grain that is the actual bug, paying write cost and storage on a table that should not exist. And there is no deploy to revert. The bad grain is a schema decision three months old plus 11 million rows of accumulated state, and unwinding it is a planned migration and a backfill, reviewed like any other change, not a rollback.
The DROP is the lucky case, and here’s the 2x2 that shows why
The failure everyone pictures is the destructive one: the DROP TABLE , the migration that truncates the wrong relation, the script that deletes the binlogs. On a gated system that is the one you have mostly handled, with destructive migrations caught in review, credentials scoped, and a restore runbook practiced. When the Replit agent wiped a production database during a code freeze in July 2025 , the data was restored: loud, attributable to one timestamp, recoverable. The scraper cleared that same pipeline because at merge time it was correct, with no rows for the grain bug to express and a CI seed that never reached the row count where it breaks.
Rank a failure on two axes, loud-vs-quiet and recoverable-vs-not, and you get four corners:
Loud and recoverable is the DROP. You know instantly and you roll it back.
Loud and unrecoverable is rarer: a destructive operation you catch but can’t undo.
Quiet and recoverable is the bug that sat unnoticed but is still reversible when you find it.
Quiet and unrecoverable ruins quarters. You don’t find out for months, and by then the prior state is gone or never existed as a clean artifact.
The axes correlate, which is what makes the bad corner deep. Loud failures get caught while the prior state still exists; quiet ones sit, and the longer one sits the more downstream systems consume it as truth, until a recoverable error has been aggregated and propagated into an unrecoverable one. Loudness buys the time, so quiet and unrecoverable travel together.
AI floods the bad corner for a structural reason. Each change it ships is plausible: it compiles, runs, returns the right shape, passes whatever checks exist. That plausibility is the corruption floor , the same mechanism that makes the output useful making it occasionally wrong in a way that looks exactly right. A loud failure is one the output failed to make plausible; the quiet one is the model working as designed. Then volume multiplies it: a team shipping eighty changes a week instead of eight samples that floor ten times as often, on the same review and CI budget. The DROP is the rare draw the pipeline was built to stop. The thousand cuts are the modal draw, and it waves them through.
The quiet corner comes in three shapes. The scraper is the first: a schema correct at zero rows and fatal at eleven million, because the model optimizes from its training distribution, not your scale. Asked to partition a large table it reaches for created_at in the primary key, the common corpus shape, not the primary-key partitioning that fits a high-scale OLTP table . The second is the value it computes wrong because it doesn’t hold your domain: Cursor’s support bot invented a login policy that didn’t exist and users cancelled before anyone knew, and Air Canada lost in court over a bereavement refund its chatbot made up. Move that same generator onto a write path computing a discount or a tax split and the row is well-typed and wrong about what the number means, with nothing reconciling it against the contract until quarter-end. The third is the change shipped past the author’s own understanding: it looked good and worked, so it went, and the judgment that would have caught it is built by the slow work the agent now skips. That is the paradox of the fast engineer , which a July 2025 study measured as 19% slower even as the developers felt faster.
A worked example: the soft-delete leak
The grain bug is loud once you go looking, because it shows up as latency. The worse version of the same class is a write that corrupts a number and never moves a performance metric at all. Soft deletes are where this lives in most enterprise schemas.
The convention is old and unwritten. A row is never physically removed; it gets a deleted_at stamp, and every query that reads the table is expected to filter it out. A billing system for a SaaS company looks like this:
1
2
3
4
5
6
7
8
-- one row per recurring line on an account: base plan, seats, add-ons
CREATE TABLE subscription_items (
id BIGSERIAL PRIMARY KEY ,
account_id BIGINT NOT NULL REFERENCES accounts ( account_id ),
sku TEXT NOT NULL ,
monthly_cents BIGINT NOT NULL ,
deleted_at TIMESTAMPTZ -- set when a customer drops the line
);
When a customer downgrades, the application does not delete the row, it stamps it:
1
UPDATE subscription_items SET deleted_at = now () WHERE id = $ 1 ;
Every existing query that touches money knows this. The MRR rollup, the invoice generator, the revenue dashboard, all of them carry AND deleted_at IS NULL , because the team learned years ago that forgetting it double-counts churned revenue. That knowledge lives in the queries and in the heads of the people who wrote them. It is nowhere in the schema; deleted_at is just a nullable timestamp, and nothing stops a query from ignoring it.
Now an agent is asked to add a board-facing metric, monthly recurring revenue by region, that the nightly ETL persists into the warehouse fact tables the dashboards read. It writes the obvious thing:
1
2
3
4
5
6
CREATE VIEW mrr_by_region AS
SELECT a . region , SUM ( si . monthly_cents ) AS mrr_cents
FROM accounts a
JOIN subscription_items si ON si . account_id = a . account_id
WHERE a . deleted_at IS NULL -- remembered on accounts
GROUP BY a . region ;
It filtered deleted_at on accounts but not on subscription_items , because nothing in the schema said it had to and the training corpus is full of joins shaped exactly like this. Every cancelled add-on and downgraded seat is now summed back into regional MRR, and each night the ETL reads this view and writes the inflated number into the warehouse the whole company reads as truth. The shape is right and the number is plausible, a little high, and growing as the soft-deleted pool grows.
Nothing catches it. The engineer saw a working view and a deleted_at filter sitting right there and moved on; the AI reviewer flagged a naming nit; the human skimmed the green summary and approved; CI ran on a seed database with almost no deleted rows, so the leak was a rounding error in the test. A missing predicate is an absence, the one thing every reviewer, human or model, reliably fails to see.
The drift is the tell. Small at launch, ten or fifteen percent high a year later in the regions with the most downgrade history, and finance finds it the only way anyone does: reconciling the dashboard against billed revenue, a year in, with no single cause and no deploy to revert.
The fix is a business call, not a technical one
The mitigations are known and none of them are clever: reconcile against a source of truth on a cadence the business can stand, alarm on aggregates and drift instead of only errors, run CI against production-shaped volume, bake invariants into views and constraints. All worth doing, all secondary, because every one is a net thrown after the write has committed. The load-bearing decision is upstream of the tooling, and it is a positioning call leadership makes on purpose. Four honest positions:
Ship at full speed, accept the corner. Take the velocity and take the unrecoverable write, the silent data loss, the bug the customer finds, as the cost. Legitimate for a seed-stage product with no prior state worth protecting. Catastrophic for a billing system.
Fast where it’s cheap, gated where it’s not. Agents and junior engineers run on loud, recoverable surfaces (internal tooling, dashboards, throwaway analysis); writes that touch money, schema, or multi-writer tables go through someone who holds th

[truncated]

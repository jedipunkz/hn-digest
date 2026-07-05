---
source: "https://simonwillison.net/2026/Jul/5/sqlite-utils-fable/"
hn_url: "https://news.ycombinator.com/item?id=48791708"
title: "sqlite-utils 4.0rc2, mostly written by Claude Fable (for about $149.25)"
article_title: "sqlite-utils 4.0rc2, mostly written by Claude Fable (for about $149.25)"
author: "ognyankulev"
captured_at: "2026-07-05T06:30:53Z"
capture_tool: "hn-digest"
hn_id: 48791708
score: 1
comments: 0
posted_at: "2026-07-05T06:19:34Z"
tags:
  - hacker-news
  - translated
---

# sqlite-utils 4.0rc2, mostly written by Claude Fable (for about $149.25)

- HN: [48791708](https://news.ycombinator.com/item?id=48791708)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jul/5/sqlite-utils-fable/)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T06:19:34Z

## Translation

タイトル: sqlite-utils 4.0rc2、主に Claude Fable が作成 (約 $149.25)
説明: 数週間前に sqlite-utils 4.0rc1 リリースについて書きました。 Max サブスクリプションにクロード・ファブルが含まれるのはあと数日間だけなので、私は…

記事本文:
sqlite-utils 4.0rc2、主に Claude Fable によって書かれました (約 $149.25)
サイモン・ウィリソンのウェブログ
sqlite-utils 4.0rc2、主に Claude Fable によって書かれました (約 $149.25)
数週間前に sqlite-utils 4.0rc1 リリースについて書きました。 Max サブスクリプションに Claude Fable があるのはあと数日だけなので、SemVer を維持し、互換性のないメジャー バージョンをできるだけ少なくしたいと考えているので、これが本当に快適に感じられる 4.0 安定版リリースに到達するのに役立つかどうかを確認することにしました。
iPhone の Web 用 Claude Code で、次のプロンプトから始めました。
安定版 4.0 リリースを出荷する前の最終レビュー - 後で修正する場合に重大な変更となる可能性のある最後の部分を特定することが非常に重要です
これが私のために作成された最初のレポートです。私自身がまだ遭遇していなかった重大な問題がいくつかありました。Fable が「リリース ブロッカー」として分類した 5 つです。その中でも最悪のものは次のとおりです。
1. delete_where() は決してコミットせず、接続を汚染します (データ損失)
Table.delete_where() ( sqlite_utils/db.py:2948 ) は、 atomic() ラッパーを持たない裸の self.db.execute() を介して DELETE を実行します。正しくラップされている db.py:2944 の Table.delete() と比較してください。接続は in_transaction=True のままになるため、後続の atomic() 呼び出しはすべてセーブポイント ブランチ ( db.py:430-440 ) を取得し、どちらもコミットすることはありません。
db = sqlite_utils 。データベース (「dw.db」)
db [ "t" ]。 insert_all ([{ "id" : i } for i in range ( 3 )]、pk = "id" )
db [ "t" ]。 delete_where ( "id = ?" , [ 0 ]) # conn.in_transaction が True になりました
db [ "t" ]。挿入 ({ "id" : 50 })
db [ "u" ]。挿入 ({ "a" : 1 })
データベース。閉じる（）
# 再オープン: 行は [0, 1, 2] — 削除、行 50、AND テーブル u はすべて消えています。
それは本当にひどいバグです！発送しなくて本当に良かったですが、少なくとも発送はダメでした

ld は 4.0.1 ポイント リリースで修正できるバグであり、5.0 を強制するような設計上の欠陥ではありません。
37 のプロンプト、34 のコミット、および 30 の個別のファイルにわたる +1,321 -190 のコード変更を経て、一連のフィードバックを順番に処理し、途中で他のいくつかの設計改善を行いました。
エージェントのコーディングに関する奇妙な点は、エージェントが新しいタスクを大量に実行するのに 10 ～ 15 分かかる場合があるため、このような難しいタスクでは実際に他のことを同時に行う機会が増えることです。私はハーフムーンベイの独立記念日のパレードを楽しむために外出し、時々チェックインして、携帯電話からフェイブルの次のステップを促しました。
詳細については、PR およびこの共有トランスクリプトをご覧ください。私はラップトップに切り替えて最終レビューを行い、GitHub の PR インターフェイスを通じて実行しました。
最も重要な変更はトランザクション処理に関連しており、これは以前の RC の特徴的な新機能でした。新しい RC には、新しいトランザクション モデルに関する包括的なドキュメントが含まれています。その概要をここで全文引用します。
データベースに書き込むこのライブラリ内のすべてのメソッド ( insert() 、 upsert() 、 update() 、 delete() 、 delete_where() 、transform() 、 create_table() 、 create_index() 、enable_fts() および残りのメソッドは、独自のトランザクション内で実行され、返される前にコミットされます。メソッド呼び出しが完了すると、変更はすぐにディスクに保存されます。
db = データベース ( "data.db" )
データベース。表 (「ニュース」)。 insert ({ "見出し" : "犬が賞を受賞" })
# 新しい行はすでに保存されています - commit() は必要ありません
db.execute() で実行される生の SQL にも同じことが当てはまります。書き込みステートメントは実行後すぐにコミットされます。
commit() を呼び出す必要はなく、変更を保持するためにデータベースを閉じる必要もありません。考慮する必要がある状況は 2 つあります。

トランザクション:
複数の書き込み操作をグループ化して、すべて成功するかすべて失敗するようにするには、 db.atomic() を使用します。
db.begin() を使用してトランザクションを自分で管理している場合、コミットするまで何もコミットされません。ライブラリは、開いたトランザクションを決してコミットしません。
Fable のドキュメントを確認していると、最初にドキュメントの編集内容を確認することが、何が変更されたのかを最初に理解するのに最適な方法であることがわかりました。私は次の詳細を発見しました。
db.atomic() とメソッドごとの自動トランザクションは、Python のデフォルトのトランザクション処理モードでの接続用に設計されています。 Python 3.12 以降の sqlite3.connect(..., autocommit=True) または autocommit=False オプションを使用して作成された接続は、それらの接続では commit() と rollback() の動作が異なるため、サポートされていません。
正直、Python 3.12 で追加された最近の autocommit 設定に sqlite-utils がどのように反応するかについては考えていませんでした。 「これらの接続で動作が異なる」ということは、テスト スイートのほぼ全体が失敗することと同じであることが判明したため、この違いがライブラリの動作を妨げないようにモデルを操作しました。
あるモデルが別のモデルの作業をレビューするというアイデアは、いくぶんばかばかしく、奇妙に迷信的だと感じていました。問題は、それが実際に機能するということです。私は、Anthropic の最高のモデルに OpenAI の成果をレビューしてもらうことを習慣的に始めており、その逆も同様です。価値のある興味深い結果が頻繁に得られるからです。
Codex Desktop と GPT-5.5 xhigh に次のプロンプトを表示しました。
前回の RC 以降の変更を確認します。また、変更ログが最新であることも確認してください。
これだけでも、調査する価値のある 2 つの問題を明らかにするのに十分でした。
[P1] sqlite_utils/db.py:663 db.query() は、db.execute() を呼び出した後にのみ行以外のステートメントを拒否するようになりました。

、sqlite_utils/db.py:705 は最初にそれらの書き込みを自動コミットします。そのため、db.query("update ...") は ValueError を発生させますが、更新はすでにコミットされています。これは、「行を返す SQL でのみ使用できる」と文書化されているメソッドにとっては驚くべき副作用です。
[P1] sqlite_utils/db.py:672 INSERT ... db.query() による RETURNING は、返されたジェネレーターが完全に使い果たされた後にのみコミットされます。 db.query("insert ...returning ...") を反復せずに、または一般的な next(db.query(...)) を使用すると、トランザクションは開いたままになり、書き込みはクローズ時にロールバックできます。これは、反復なしで有効になるという docs/changelog.rst:15 および docs/python-api.rst:232 と矛盾します。
これを新しい Fable セッションに貼り付け、問題を確認するためにいくつかの実験を実行しました。
両方の結果が確認されました。 db.query() は最初に self.execute() を呼び出し、書き込みを自動コミットし、その後でcursor.descriptionをチェックしました。つまり、 db.query("update ...") は ValueError を発生させる前に更新をコミットしました。そして、INSERT ... RETURNING コミットは返されたジェネレーターの最後に存在していたので、イテレーターを使い果たさない限り、つまり next(db.query(...)) または反復されていない呼び出しによってトランザクションが開いたままにならない限り、コミットは起動されませんでした。これは、変更ログとドキュメントの約束に反しています。
以下は修正を含む PR とクロード コードの完全な転写です。このコードをレビューすることで、SQLite トランザクション セマンティクスのエッジ ケースのより良いメンタル モデルを構築することができました。
私は Claude Max $200/月プラン (以前は $100/月) にアップグレードして、7 月 7 日の Fablepocalypse までの残りの期間、Fable の許容量を増やしました。この時、Claude Max サブスクライバーでもモデルの API コストを全額支払う必要があります。
これらの費用を直接支払った場合、どれくらいの費用がかかるか興味がありました。最初はその番号は利用できないのかと思いました

Web 用の Claude Code を使用してリモートで作業を実行した後、既存のセッション内で AgentsView を実行してコストの見積もりを取得できることに気づきました。
「uvx Agentsview --help」を実行し、そのツールを使用してこのセッションのコストを計算します。
Claude は、session list --include-children コマンドの使い方を考え出し、次のような結果を導き出しました。
サブスクリプションに加入していてとてもうれしいです!私は本当に自分自身のアドバイスに従って、より安価なモデルを提供するサブ代理店にもっと重点を置くべきでした。
現在、claude.ai/settings/usage が表示している内容は次のとおりです。
私は現在、他にもいくつかの主要な Fable 主導のプロジェクトを進めており、値上げに間に合うようにその Fable バーで 100% を達成することを目標としています。
RC の完全なリリース ノートは次のとおりです。私は Fable に、変更が反映されるたびにこれらを変更ログの「未リリース」セクションに追加してもらい、進行状況を確認してもらいました。これには、変更ログのコミット履歴が、リリースに加えられた各変更の簡潔な概要として機能するという優れた副作用があります。
これまで私はリリース ノートを手書きする方針でしたが、正直なところ、自分で作成するよりも良いものになりました。リリースノートは、退屈で予測可能で正確である必要があるため、エージェントに外注しても問題がないという文書の好例です。
db.execute() で実行される書き込みステートメントは、トランザクションがすでにオープンしている場合を除き、トランザクションに参加しない限り、自動的にコミットされるようになりました。以前は、何かがそれをコミットするまで開いたままになる暗黙的なトランザクションが開かれていました。書き込みは、同じ接続上で読み取られるときは機能するように見えましたが、接続が閉じられると暗黙的にロールバックされました。コミットされていない db.execute() 書き込みのロールバックに依存するコードは、新しい db.begin() メソッドを使用して明示的な t を開く必要があります。

まずはトランザクション。トランザクション モデルの詳細は、「トランザクションと変更の保存」に記載されています。
db.query() は、返されたジェネレーターが最初に反復されるまで待つのではなく、呼び出されるとすぐに SQL を実行するようになりました。行は反復中に引き続き遅延フェッチされます。 SQL エラーが呼び出しサイトで発生するようになりました。INSERT ... RETURNING などのステートメントは、結果を反復する必要なく、ただちに実行およびコミットされます。また、行を返さないステートメント (以前はサイレント no-op) を渡すと、代わりに db.execute() を推奨する ValueError が発生するようになりました。この方法で拒否されたステートメントは、エラーが発生する前にロールバックされるため、データベースには影響しません。
Python API 検証エラーでは、 AssertionError ではなく ValueError が発生するようになりました。以前は無効な引数 (列のない create_table()、存在しないテーブルの transform()、ignore=True と replace=True の両方を渡すなど) は、裸の Assert ステートメントを使用して拒否されていましたが、Python が -O フラグを指定して実行されると、これらの引数は暗黙的にスキップされていました。これらの場合に AssertionError をキャッチしたコードは、代わりに ValueError をキャッチする必要があります。
table.upsert() および table.upsert_all() は、レコードに主キー列の値が欠落している場合、またはレコードの値が None である場合に PrimaryKeyRequired を生成するようになりました。以前は、既存の行と決して一致しないこのようなレコードは、まったく新しい行として静かに挿入されるか、挿入がすでに行われた後に紛らわしい KeyError をトリガーしていました。
db.enable_wal() と db.disable_wal() は、トランザクションが開いているときに呼び出された場合に sqlite_utils.db.TransactionError を発生させるようになりました。以前は、ジャーナル モードの変更の副作用として、開いているトランザクションがサイレントにコミットされ、db.atomic() とユーザー管理のトランザクションのロールバック保証が破られていました。
View クラスには、

Enable_fts() メソッド。ビューでは全文検索がサポートされていないため、これは NotImplementedError を発生させるためだけに存在していました。これを呼び出すと代わりに AttributeError が発生し、このメソッドは API リファレンスに表示されなくなりました。 sqlite-utils enable-fts コマンドは、ビューをポイントするとクリーンなエラーを表示します。
no-op -d/--detect-types フラグは、insert および upsert コマンドから削除されました。 4.0a1 以降、CSV/TSV データの型検出はデフォルトになっているため、このフラグは何もしませんでした。このフラグを使用する呼び出しでは単にフラグを削除する必要があります。 --no-detect-types は引き続き検出を無効にするために使用できます。
Python 3.12 以降の sqlite3.connect(..., autocommit=True) または autocommit=False オプションで作成された接続が渡された場合、Database() は sqlite_utils.db.TransactionError を発生させるようになりました。 commit() と rollback() はこれらの接続上で異なる動作をします。そのため、以前は、接続が閉じられたときにライブラリによって行われたすべての書き込みがサイレントに破棄されていました。
table.delete_where() 、 table.optimize() 、および table.rebuild_fts() が変更をコミットせず、開いたトランザクション内の接続が残るバグを修正しました。接続が閉じられると、その作業とその後の書き込みはサイレントにロールバックされる可能性があります。 3 つはすべて、他の書き込みメソッドと一貫して db.atomic() を使用するようになりました。
sqlite-utils drop-table コマンドは t を拒否するようになりました。

[切り捨てられた]

## Original Extract

I wrote about the sqlite-utils 4.0rc1 release a couple of weeks ago. Since we only have Claude Fable on our Max subscriptions for a few more days, I decided to …

sqlite-utils 4.0rc2, mostly written by Claude Fable (for about $149.25)
Simon Willison’s Weblog
sqlite-utils 4.0rc2, mostly written by Claude Fable (for about $149.25)
I wrote about the sqlite-utils 4.0rc1 release a couple of weeks ago. Since we only have Claude Fable on our Max subscriptions for a few more days, I decided to see if it could help me get to a 4.0 stable release that I felt truly comfortable about, since I try to keep to SemVer and like my incompatible major versions to be as rare as possible.
I started with this prompt, in Claude Code for web on my iPhone:
Final review before shipping a stable 4.0 release - very important to spot any last minute things that would be a breaking change if we fix them later
Here’s that initial report it created for me. There were some significant problems that I hadn’t myself encountered yet—5 that Fable categorized as “release blockers”. Here’s the worst of the bunch:
1. delete_where() never commits and poisons the connection (data loss)
Table.delete_where() ( sqlite_utils/db.py:2948 ) runs its DELETE via a bare self.db.execute() with no atomic() wrapper — compare Table.delete() at db.py:2944 , which wraps correctly. The connection is left in_transaction=True , so every subsequent atomic() call takes the savepoint branch ( db.py:430-440 ) and never commits either.
db = sqlite_utils . Database ( "dw.db" )
db [ "t" ]. insert_all ([{ "id" : i } for i in range ( 3 )], pk = "id" )
db [ "t" ]. delete_where ( "id = ?" , [ 0 ]) # conn.in_transaction is now True
db [ "t" ]. insert ({ "id" : 50 })
db [ "u" ]. insert ({ "a" : 1 })
db . close ()
# Reopen: rows are [0, 1, 2] — the delete, row 50, AND table u are all gone.
That’s a really bad bug! Very glad I didn’t ship that, although at least it would have been a bug I could fix in a 4.0.1 point release, not a design flaw that would force a 5.0.
Over the course of 37 prompts, 34 commits and +1,321 -190 code changes over 30 separate files, we worked through the entire set of feedback in turn, making several other design improvements along the way.
A weird thing about coding agents is that harder tasks like this one actually provide more opportunity to do other things at the same time, since the agent sometimes needs 10-15 minutes to churn away on a new task. I went out to enjoy the Half Moon Bay 4th of July parade, occasionally checking in and prompting the next step for Fable from my phone.
Full details in the PR and this shared transcript . I switched to my laptop for the final review, which I conducted through GitHub’s PR interface.
The most significant changes relate to transaction handling, which was the signature new feature in the earlier RC . The new RC now includes comprehensive documentation on the new transaction model, the intro to which I’ll quote here in full:
Every method in this library that writes to the database— insert() , upsert() , update() , delete() , delete_where() , transform() , create_table() , create_index() , enable_fts() and the rest—runs inside its own transaction and commits it before returning. Your changes are saved to disk as soon as the method call finishes:
db = Database ( "data.db" )
db . table ( "news" ). insert ({ "headline" : "Dog wins award" })
# The new row is already saved - no commit() required
The same applies to raw SQL executed with db.execute() —a write statement is committed as soon as it has run.
You never need to call commit() , and you do not need to close the database to persist your changes. There are exactly two situations where you need to think about transactions:
You want to group several write operations together, so they either all succeed or all fail—use db.atomic() .
You are managing a transaction yourself with db.begin() , in which case nothing is committed until you commit—the library will never commit a transaction you opened.
In reviewing Fable’s documentation—I find that reviewing the documentation edits first is an excellent way to build an initial understanding of what has changed—I spotted this detail :
db.atomic() and the automatic per-method transactions are designed for connections in Python’s default transaction handling mode. Connections created with the Python 3.12+ sqlite3.connect(..., autocommit=True) or autocommit=False options are not supported, because commit() and rollback() behave differently on those connections.
I admit I hadn’t thought about how sqlite-utils would react to the more recent autocommit setting , added in Python 3.12. It turns out “behave differently on those connections” equated to almost the entire test suite failing, so I worked with the model to ensure that this difference would not break how the library works.
I used to think that the idea of having one model review the work of another was somewhat absurd—it felt weirdly superstitious. The problem is it really does work —I’ve started habitually having Anthropic’s best model review OpenAI’s work and vice versa, because I’ve had that turn up interesting results often enough to be valuable.
I prompted Codex Desktop and GPT-5.5 xhigh with the following:
Review changes since the last RC. Also confirm that the changelog is up-to-date.
Which was enough to turn up two issues worth investigating:
[P1] sqlite_utils/db.py:663 db.query() now rejects non-row statements only after calling db.execute() , and sqlite_utils/db.py:705 auto-commits those writes first. So db.query("update ...") raises ValueError but the update is already committed. That is a surprising side effect for a method documented as “can only be used with SQL that returns rows.”
[P1] sqlite_utils/db.py:672 INSERT ... RETURNING through db.query() only commits after the returned generator is fully exhausted. db.query("insert ... returning ...") without iteration, or common next(db.query(...)) usage, leaves the transaction open and the write can be rolled back on close. This contradicts docs/changelog.rst:15 and docs/python-api.rst:232 , which say it takes effect without iteration.
I pasted that into a fresh Fable session, which ran some experiments to confirm the problem:
Both findings were confirmed. db.query() called self.execute() first, which auto-commits writes, and only then checked cursor.description — so db.query("update ...") committed the update before raising ValueError . And the INSERT ... RETURNING commit lived at the end of the returned generator, so it never fired unless you exhausted the iterator — next(db.query(...)) or an un-iterated call left the transaction open, contradicting what the changelog and docs promise.
Here’s the PR with the fix, and the full Claude Code transcript . Reviewing this code helped me build a better mental model of the edge cases of SQLite transaction semantics!
I upgraded to the Claude Max $200/month plan (I was previously on $100/month) to increase my Fable allowance for the remaining time until the July 7th Fablepocalypse , when even Claude Max subscribers will have to pay full API cost for the model.
I was curious as to how much this would have cost me if I had been paying those costs directly. At first I thought those numbers weren’t available to me since I had run the work remotely using Claude Code for web, and then I realized I could run AgentsView inside that existing session to get that cost estimate!
Run "uvx agentsview --help"​ and then use that tool to calculate the cost of this session
Claude figured out how to use the session list --include-children command and came out with the following:
I’m very glad I’m on that subscription! I really should have followed my own advice and leaned more heavily into subagents with cheaper models.
Here’s what claude.ai/settings/usage is showing me right now:
I have several other major Fable-driven projects on the go right now as well, with the goal of hitting 100% on that Fable bar just in time for the price increase.
Here are the full release notes for the RC. I had Fable add these to an “Unreleased” section of the changelog as each change landed, reviewing them as it went. This has the neat side effect that the commit history of the changelog acts as a concise summary of each of the changes that went into the release.
In the past I’ve had a policy of writing release notes by hand, but honestly these are better than I would have created myself. Release notes are a great example of writing that I’m OK to outsource to agents because they need to be boring, predictable and accurate.
Write statements executed with db.execute() are now committed automatically, unless a transaction is already open in which case they join it. Previously they opened an implicit transaction that stayed open until something committed it—writes appeared to work when read on the same connection but were silently rolled back when the connection closed. Code that relied on rolling back uncommitted db.execute() writes should use the new db.begin() method to open an explicit transaction first. The transaction model is documented in full at Transactions and saving your changes .
db.query() now executes its SQL as soon as it is called, rather than waiting until the returned generator is first iterated. Rows are still fetched lazily during iteration. SQL errors are now raised at the call site, statements such as INSERT ... RETURNING are executed and committed immediately without needing to iterate over their results, and passing a statement that returns no rows—previously a silent no-op—now raises a ValueError recommending db.execute() instead. A statement rejected this way is rolled back before the error is raised, so it has no effect on the database.
Python API validation errors now raise ValueError instead of AssertionError . Previously invalid arguments—such as create_table() with no columns, transform() on a table that does not exist, or passing both ignore=True and replace=True —were rejected using bare assert statements, which are silently skipped when Python runs with the -O flag. Code that caught AssertionError for these cases should catch ValueError instead.
table.upsert() and table.upsert_all() now raise PrimaryKeyRequired if a record is missing a value for any primary key column, or has a value of None for one. Previously such records—which can never match an existing row—were quietly inserted as brand new rows, or triggered a confusing KeyError after the insert had already taken place.
db.enable_wal() and db.disable_wal() now raise a sqlite_utils.db.TransactionError if called while a transaction is open. Previously they would silently commit the open transaction as a side effect of changing the journal mode, breaking the rollback guarantee of db.atomic() and of user-managed transactions.
The View class no longer has an enable_fts() method. It existed only to raise NotImplementedError , since full-text search is not supported for views—calling it now raises AttributeError instead, and the method no longer appears in the API reference. The sqlite-utils enable-fts command shows a clean error when pointed at a view.
The no-op -d/--detect-types flag has been removed from the insert and upsert commands. Type detection has been the default for CSV/TSV data since 4.0a1, so the flag did nothing—invocations using it should simply drop it. --no-detect-types remains available to disable detection.
Database() now raises a sqlite_utils.db.TransactionError if passed a connection created with the Python 3.12+ sqlite3.connect(..., autocommit=True) or autocommit=False options. commit() and rollback() behave differently on those connections, which previously caused every write made by the library to be silently discarded when the connection closed.
Fixed a bug where table.delete_where() , table.optimize() and table.rebuild_fts() did not commit their changes, leaving the connection inside an open transaction. Their work—and any subsequent writes—could then be silently rolled back when the connection was closed. All three now use db.atomic() , consistent with the other write methods.
The sqlite-utils drop-table command now refuses t

[truncated]

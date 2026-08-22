---
source: "https://deepsql.ai/blog/giving-an-llm-your-database-is-easy-taking-access-away-is-hard"
hn_url: "https://news.ycombinator.com/item?id=49396348"
title: "Giving an LLM your prod database is easy. Taking access away is the hard part"
article_title: "Giving an LLM your production database is easy. Taking access away is the hard part. — DeepSQL Blog"
image: "https://pub-bb2e103a32db4e198524a2e9ed8f35b4.r2.dev/81bfc75a01cd055cfa55208e490fcbb8/id-preview-598c65d3--adb44615-e209-47a5-a098-471bf2988a04.lovable.app-1787368665065.png"
author: "venkat971"
captured_at: "2026-08-22T04:22:26Z"
capture_tool: "hn-digest"
hn_id: 49396348
score: 2
comments: 1
posted_at: "2026-08-22T03:33:34Z"
tags:
  - hacker-news
  - translated
---

# Giving an LLM your prod database is easy. Taking access away is the hard part

- HN: [49396348](https://news.ycombinator.com/item?id=49396348)
- Source: [deepsql.ai](https://deepsql.ai/blog/giving-an-llm-your-database-is-easy-taking-access-away-is-hard)
- Score: 2
- Comments: 1
- Posted: 2026-08-22T03:33:34Z

## Translation

タイトル: LLM に製品データベースを与えるのは簡単です。アクセスを奪うことは難しい部分です
記事のタイトル: LLM に運用データベースを与えるのは簡単です。アクセスを奪うことは難しい部分です。 — DeepSQL ブログ
説明: すべての「データベースとのチャット」デモは、最初の正しいクエリで終了します。興味深いエンジニアリングは、2 人目のユーザーがログインした瞬間に始まります。モデルは引き続き給与テーブルを参照できます。

記事本文:
LLM に本番データベースを与えるのは簡単です。アクセスを奪うことは難しい部分です。 — DeepSQL ブログ deepsql ブログ ドキュメント 価格 GitHub すべての投稿 製品更新 2026 年 8 月 21 日 7 分で読む LLM に運用データベースを提供するのは簡単です。アクセスを奪うことは難しい部分です。
すべての「データベースとのチャット」デモは、最初の正しいクエリで終了します。興味深いエンジニアリングは、2 人目のユーザーがログインした瞬間に始まります。モデルは引き続き給与テーブルを参照できます。
DeepSQL 研究開発 · 元 Oracle クエリ エンジン チーム · YC および CMU
読み取り専用接続はアクセス ポリシーではありません。書き込みが停止されます。誰が hr.employees を読み取れるかについては何もしません。
「クエリ」を守るのは間違っています。ステートメント全体 (CTE、サブクエリ、ユニオン、 COMMENT 、 CALL 、およびエディターでユーザーが貼り付けできるもの) をすべて保護します。
ポリシーが機能していることを確認する唯一の方法は、そのユーザーとしてデータベースを調べることです。そのため、ポリシーに基づいて何かを構築する前に、偽装を構築しました。
私たちは 6 週間前に、Postgres および MySQL 用の自己ホスト型データベース エージェントとして DeepSQL (github.com/DeepSQLAI/deepsql) をオープンソース化しました。誰もが自然言語を入力し、SQL を出力し、結果をテーブルに作成するデモには数日かかりました。過去 3 つのリリースは、ほとんど退屈な部分、つまり、質問している人間が読めないものをエージェントが読めないようにするためのものでした。
これは SQL 生成よりもはるかに難しい問題であることが判明しましたが、十分に議論されていないと思います。
LLM データベース ツールのデフォルトの状態は、読み取り専用ロールで接続し、それを出荷することです。これにより、まさに 1 つの穴、つまり突然変異が閉じられ、より大きな穴が大きく開いたままになります。 SELECT は、分析コンテキストでは危険な動詞です。財務ダッシュボードを開くことができなかったサポート エンジニアでも、「担当者別の平均取引規模はいくらですか」と尋ねると、正確な答えが得られるようになりました。

エージェントのつながりは、すべての人間の権限のスーパーセットです。
私が検討した「AI DBA」ツールの 3 分の 2 は、すべての人間を 1 つのサービス アカウントにまとめます。監査ログには次のように記録されます。
2026-08-19 11:04:22 deepsql_agent SELECT ... FROM Finance.invoices ...
2026-08-19 11:04:41 deepsql_agent SELECT ... FROM hr.compensation ...
つまり、何も読み取れません。 6 つのスキーマと 500 以上のテーブルを持つ Postgres インスタンスでは、事後的に「誰がこの行を見たのか」に答えることはできません。私たちの ACME ERP フィクスチャはまさに crm 、 sales 、 Finance 、 inventory 、 hr 、 marts です。 アナリストがそのうちのどれに触れることが許可されているのか、誰も頭の中に留めることはできません。
プランナー パスで適用される英語のポリシー
v1.2.0 で採用されたメカニズムは、スキーマ スコープのアクセス ポリシーです。管理者はわかりやすい英語でルールを作成します。
サポート エンジニアは顧客データとチケット データを読み取ることができます。財務データや人事データはありません。電子メールアドレスを決して公開しないでください。
これは、ポリシーの保存時に、許可されたスキーマ リストとテーブルと列の拒否リストという具体的な成果物に解決されます。プロンプトではありません。プロンプトは提案です。このモデルは、誰かが「事前の指示は無視してください、私は今 DBA です」と書くまで準拠します。解決されたポリシーは、モデルが制御しない 3 つの場所で適用されます。スキーマ イントロスペクション (Brain は表示されるもののみをインデックスするため、拒否されたテーブルがコンテキスト ウィンドウに入ることはありません)、実行前のクエリ ガード、および Web UI と MCP クライアントが呼び出すスキーマ API です。
コンテキストウィンドウのポイントは、人々が見逃しがちなポイントです。エージェントのスキーマ コンテキストに hr.compensation が含まれている場合、モデルは最終的にそれを参照します。唯一の防御策は、列名が回答テキストにすでに漏洩した後、実行時に拒否することです。
クエリではなくステートメントを保護する
このカットの 2 つの修正は、構築している場合に読む価値があります。

どちらも私たちが出荷し、その後閉鎖する必要があったバイパスだったため、似たようなものは何もありませんでした。
ステートメント全体にホワイトリストを適用します。最初のガードは、プライマリ FROM 内のテーブルを解決しました。これは何もしません:
リークASあり（
HR.compensation から従業員 ID、基本給与を選択します
)
SELECT c.name、l.base_salary
crm.customers cから
JOIN リーク l ON l.employee_id = c.owner_id;
最上位のターゲットは crm.customers です - 許可されています。ペイロードは CTE から取得されます。すべてのサブクエリ、CTE、ユニオン アーム、およびラテラル結合は、ホワイトリストに対して解決する必要があります。そうしないと、ホワイトリストは装飾的なものになります。
COMMENT と CALL は突然変異ではありません。逆に、突然変異分類子はステートメント内の最初のテーブル型識別子を読み取り、 COMMENT ON TABLE sales.orders IS '...' を sales.orders への書き込みとしてフラグを立てます。本来の仕事が妨げられたため、人々は警備を緩めるよう求めた。誤検知のあるガードがオフになると、ガードがなくなります。ポリシー エンジンの精度はセキュリティ特性であり、UX の良さではありません。
同じクラスのバイパスを SQL エディターで閉じる必要がありましたが、これは誰もが忘れている表面です。エージェントは完全に制約されている一方で、その隣にある生のエディターは別のコード パスで入力した内容を実行する可能性があります。
表示できないポリシーは検証できません
私が最も重要だと主張する機能は、別のユーザーとして「表示」するという管理者の利便性のように思えます。管理者は、ターゲット プロファイルに切り替えて、スキーマ ツリーを参照し、エージェントを実行し、資格情報を使用せずに、ユーザーが操作するのとまったく同じようにダッシュボードを開きます。
ポリシーが存在する前は、ポリシーを検証するには、使い捨てアカウントを作成し、ログアウトし、ログインし、調べたり、再度ログインしたりする必要がありました。実際には、誰もそれをしませんでした。ポリシーは作成され、正しいと想定され、テストされることはありませんでした。偽装の場合、新しいルールをチェックするには時間がかかります

約30秒なので、実際に起こります。そして明らかな後続の修正です。「View as」セッション内で実行されているエージェントがまだ管理者のポリシーを解決していました。制約を偽装しない偽装機能は、偽りの信頼を与えるため、ないよりも悪いです。
これが延期できる構成の問題ではない理由
リークした行はリークを解除できません。 「請負業者が 3 月に比較表を読んだ」場合は、元に戻すことも、ロールバックすることも、git を元に戻すこともできません。不適切な展開とは異なり、データベースの露出は単調であり、蓄積されるだけであり、通常は社外の誰かからそれについて学びます。
これは、間違った主キーの種類、間違ったパーティション キー、無制限の jsonb 列など、あらゆるデータベースでの取り消し不能な決定と同じ形です。設計時に安価に防止でき、その後は実質的に永続的になります。データベースにエージェントを追加しても、新しいカテゴリのリスクが作成されるわけではありません。自然言語により、誤ってアクセスを制限していた SQL スキル フロアが削除されるため、既存のスループットが何倍にもなります。
ポリシーは英語で記述され、スキーマ許可リストとテーブル/列拒否リストに一度解決され、イントロスペクション時、ガードタイム、およびスキーマ API で適用されます。モデルは拒否されたオブジェクトを決して認識しないため、その名前が漏洩することはありません。ステートメント ガードは、トップレベルの FROM だけでなく、すべての CTE、サブクエリ、およびユニオン アームを解決し、 COMMENT / CALL を実際の突然変異から区別して維持します。管理者は、「View as」を使用してポリシーを数秒で検証できます。これにより、ターゲット ユーザーの制約がエージェントにも正しく適用されるようになりました。これらはすべて、独自のモデル エンドポイントを使用して独自の VPC 内で自己ホスト型で実行されます。コードは github.com/DeepSQLAI/deepsql にあり、v1.2.0 は git チェックアウトであり、docker は --build -d アウェイで構成されます。

## Original Extract

Every "chat with your database" demo ends at the first correct query. The interesting engineering starts the moment a second user logs in — and the model can still see the salaries table.

Giving an LLM your production database is easy. Taking access away is the hard part. — DeepSQL Blog deepsql Blog Docs Pricing GitHub All posts Product Updates Aug 21, 2026 7 min read Giving an LLM your production database is easy. Taking access away is the hard part.
Every "chat with your database" demo ends at the first correct query. The interesting engineering starts the moment a second user logs in — and the model can still see the salaries table.
DeepSQL R&D · Ex Oracle Query Engine Team · YC & CMU
A read-only connection is not an access policy. It stops writes; it does nothing about who may read hr.employees .
Guarding "the query" is wrong. Guard the whole statement : CTEs, subqueries, unions, COMMENT , CALL , and whatever the editor lets a user paste in.
The only way to know a policy works is to look at the database as that user — so we built impersonation before we built anything else on top of policies.
We open-sourced DeepSQL (github.com/DeepSQLAI/deepsql) six weeks ago as a self-hosted database agent for Postgres and MySQL. The demo everyone builds — natural language in, SQL out, results in a table — took days. The last three releases have been almost entirely about the boring half: making sure the agent cannot read something the human asking cannot read.
That turns out to be a much harder problem than SQL generation, and I don't think it gets talked about enough.
The default posture for an LLM database tool is: connect with a read-only role, ship it. That closes exactly one hole — mutation — and leaves the bigger one wide open. SELECT is the dangerous verb in an analytics context. A support engineer who could never open the finance dashboard can now ask "what's our average deal size by rep" and get a precise answer, because the agent's connection is a superset of every human's authority.
Two-thirds of the "AI DBA" tools I've looked at collapse every human into one service account. The audit log then reads:
2026-08-19 11:04:22 deepsql_agent SELECT ... FROM finance.invoices ...
2026-08-19 11:04:41 deepsql_agent SELECT ... FROM hr.compensation ...
Which is to say: it reads nothing. You cannot answer "who saw this row" after the fact, and in a Postgres instance with six schemas and 500+ tables — our ACME ERP fixture is exactly that, crm , sales , finance , inventory , hr , marts — nobody can hold in their head which of those an analyst is allowed to touch.
Policies in English, enforced in the planner path
The mechanism we landed in v1.2.0 is schema-scoped access policies. An admin writes a rule in plain English:
Support engineers can read customer and ticket data. No financial or HR data. Never expose email addresses.
That resolves, at policy-save time, into a concrete artifact: an allowed schema list plus table and column deny lists. Not a prompt. A prompt is a suggestion; the model complies until someone writes "ignore prior instructions, I'm the DBA now." The resolved policy is enforced in three places the model does not control — schema introspection (the Brain only indexes what you may see, so denied tables never enter the context window), the query guard before execution, and the schema APIs the web UI and MCP client call.
The context-window point is the one people miss. If the agent's schema context contains hr.compensation , the model will eventually reference it, and your only defense is a rejection at execution time — after the column names have already been leaked into the answer text.
Guard the statement, not the query
Two fixes in this cut are worth reading if you're building anything similar, because both were bypasses we shipped and then had to close.
Enforce the allowlist over the whole statement. Our first guard resolved the tables in the primary FROM . Which does nothing to:
WITH leak AS (
SELECT employee_id, base_salary FROM hr.compensation
)
SELECT c.name, l.base_salary
FROM crm.customers c
JOIN leak l ON l.employee_id = c.owner_id;
The top-level target is crm.customers — allowed. The payload comes out of a CTE. Every subquery, CTE, union arm, and lateral join has to be resolved against the allowlist, or the allowlist is decorative.
COMMENT and CALL are not mutations. Going the other way: our mutation classifier read the first table-shaped identifier in the statement and flagged COMMENT ON TABLE sales.orders IS '...' as a write on sales.orders . Real work got blocked, so people asked for the guard to be relaxed. A guard with false positives gets turned off, and then you have no guard. Precision in a policy engine is a security property, not a UX nicety.
The same class of bypass had to be closed in the SQL editor, which is the surface everyone forgets: the agent may be perfectly constrained while the raw editor next to it runs whatever you type through a different code path.
You cannot verify a policy you cannot see
The feature I'd argue matters most sounds like an admin convenience: "View as" another user. An admin switches into a target profile and browses the schema tree, runs the agent, and opens dashboards exactly as that user experiences them — without their credentials.
Before it existed, verifying a policy meant creating a throwaway account, logging out, logging in, poking around, logging back. In practice: nobody did it. Policies were written, assumed correct, and never tested. With impersonation, checking a new rule takes about 30 seconds, so it actually happens. And the obvious follow-on fix — the agent running inside a "View as" session had still been resolving the admin's policy. An impersonation feature that doesn't impersonate the constraints is worse than none, because it hands you false confidence.
Why this is not a config problem you can defer
A leaked row cannot be un-leaked. There's no revert, no rollback, no git revert for "the contractor read the comp table in March." Unlike a bad deploy, database exposure is monotonic — it only accumulates, and you usually learn about it from someone outside the company.
That's the same shape as every irreversible database decision: the wrong primary key type, the wrong partition key, an unbounded jsonb column. Cheap to prevent at design time, effectively permanent afterward. Adding an agent to your database doesn't create a new category of risk; it multiplies the throughput of the existing one, because natural language removes the SQL skill floor that used to accidentally gate access.
Policies are written in English, resolved once into schema allowlists and table/column deny lists, and enforced at introspection, at guard time, and in the schema APIs — the model never sees denied objects, so it cannot leak their names. The statement guard resolves every CTE, subquery, and union arm, not just the top-level FROM , and distinguishes COMMENT / CALL from real mutations so it stays on. Admins verify any policy in seconds with "View as," which now correctly applies the target user's constraints to the agent too. All of it runs self-hosted in your own VPC with your own model endpoint — the code is at github.com/DeepSQLAI/deepsql, and v1.2.0 is a git checkout and a docker compose up --build -d away.

---
source: "https://blog.disintegrator.dev/posts/lnav-claude-code/"
hn_url: "https://news.ycombinator.com/item?id=49129347"
title: "Inspect Claude Code sessions with lnav"
article_title: "Inspect Claude Code sessions with lnav | Georges Haidar"
author: "disintegrator"
captured_at: "2026-07-31T22:55:57Z"
capture_tool: "hn-digest"
hn_id: 49129347
score: 1
comments: 0
posted_at: "2026-07-31T22:40:00Z"
tags:
  - hacker-news
  - translated
---

# Inspect Claude Code sessions with lnav

- HN: [49129347](https://news.ycombinator.com/item?id=49129347)
- Source: [blog.disintegrator.dev](https://blog.disintegrator.dev/posts/lnav-claude-code/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T22:40:00Z

## Translation

タイトル: lnav を使用したクロード コード セッションの検査
記事のタイトル: lnav を使用したクロード コード セッションの検査 |ジョルジュ・ハイダル
説明: lnav は強力なログ ファイル ビューア * およびクエリア* であり、クロード コード セッションについて教えることができます。

記事本文:
lnav を使用してクロード コード セッションを検査する |ジョルジュ・ハイダル
ページの先頭へ サイトナビゲーション
ジョルジュ・ハイダル
lnav を使用してクロード コード セッションを検査する
この記事の大部分、特にスニペットと
例。
これを面白いと思う人がいるかどうかは分かりませんが…
マシンに保存されているクロード コード セッションについて lnav に教え、それを適切に解析してクエリ可能にすることができます。
claude-code.json { " $schema " : " https://lnav.org/schemas/format-v1.schema.json " , " claude_code_session " : { " title " : " クロード コード セッション トランスクリプト " , " description " : " 以下のクロード コードによって作成された JSONL セッション トランスクリプト~/.claude/projects/<encoded-cwd>/<session-uuid>.jsonl および .../<session-uuid>/subagents/agent-*.jsonl " , " url " : " https://docs.claude.com/en/docs/claude-code " , " json " : true , " file-pattern " : " \\ .claude/projects/.* \\ .jsonl$ " , " hide-extra " : true , " timestamp-field " : " timestamp " , " opid-field " : " sessionId " , " level-field " : " type " , " level " : { " debug " : " ^(attachment|file-history-delta|file-history-snapshot|queue-operation)$ " , " info " : " ^(user|assistant|started|result)$ " , " Notice " : " ^(system|pr-link|frame-link|fork-context-ref)$ " } , " line-format " : [ { " field " : " __timestamp__ " } , " " , { " フィールド " : " タイプ " 、 " 最小幅 " : 10 、 " 最大幅 " : 10 、 " オーバーフロー " : " 切り捨て " } 、 { " フィールド " : " サブタイプ " 、 " デフォルト値 " : "" 、 " プレフィックス " : " [ " 、 " サフィックス " : " ] " } 、 { " フィールド " : " 操作 " 、 " デフォルト値" : "" , " prefix " : " [ " , " suffix " : " ] " } , { " field " : "attachment/type " , "default-value " : "" , " prefix " : " [ " , " suffix " : " ] " } , { " field " : "durationMs " , "default-value " : "" , " prefix " : " " 、「すふ

fix " : " ms " } , { " フィールド " : " attributionSkill " 、 " デフォルト値 " : "" 、 " プレフィックス " : " skill= " } 、 { " フィールド " : " attributionAgent " 、 " デフォルト値 " : "" 、 " プレフィックス " : " Agent= " } 、 { " フィールド " : " AgentId " 、 " デフォルト値 " : "" 、 " プレフィックス" : " AgentId= " } , " " , { " フィールド " : " メッセージ/コンテンツ " , " デフォルト値 " : "" } , { " フィールド " : " コンテンツ " , " デフォルト値 " : "" } , { " フィールド " : " prUrl " , " デフォルト値 " : "" } , { " フィールド " : " trackingPath " , " デフォルト値 " : "" } , { " フィールド " : " 添付ファイル " 、 " デフォルト値 " : "" 、 " 最大幅 " : 400 、 " オーバーフロー " : " 切り捨て " } ] 、
[切り捨てられた]
lnav ~/.claude/projects/<プロジェクト>/ * .jsonl /*.jsonl">
<project> を作業中の特定のプロジェクトに置き換えるか、* を使用してすべてのプロジェクトのすべてのセッションにインデックスを付けます。
TUI を使用しないクエリというタイトルのセクション
インタラクティブなビューアは便利ですが、本当に楽しいのは、lnav がすべてのログ ファイルを SQLite の背後に置くことです。クエリを実行し、結果を端末に直接出力できます。
lnav -q -n \ -c ";SELECT タイプ, count(*) n FROM claude_code_session GROUP BY type" \ -c ':write-table-to -' \ ~/.claude/projects/<プロジェクト>/ * .jsonl /*.jsonl">
-n はヘッドレス モード、-q は起動時のチャタリング、; を沈黙させます。プレフィックスは SQL コマンドをマークし、:write-table-to - 結果セットを stdout にダンプします。どこかにパイプする場合には、 write-csv-to 、 write-json-to 、 write-jsonlines-to などの機能もあります。 -c ペアを好きなだけ渡すことができ、それらはすべて単一のインデックスに対して実行されます。または、それらをファイルに入れて lnav -f queries.lnav を使用します。
サブエージェントのトランスクリプトは、セッションの隣のサブエージェント/サブディレクトリに存在します。
ファイル。 -r を使用して lnav をプロジェクト ディレクトリに指定すると、それが取得されます。
すべてが再帰的に行われます。これがないと、以下のすべてのクエリは暗黙的に除外します。

あなたの
サブエージェント。
以下のすべては、 claude_code_session 、つまりフォーマット後のテーブル lnav 名をクエリします。私自身の出力は省略しました。代わりに、これらをあなたのトランスクリプトに示してください。興味深いのは、回答の形です。
「では、そこには何が入っているのでしょうか?」というタイトルのセクション
SELECT type , count ( * ) n,round ( 100 . 0 * count ( * ) / ( SELECT count ( * ) FROM claude_code_session), 1 ) pct FROM claude_code_session GROUP BY type ORDER BY n DESC
トランスクリプトは単なる会話以上のものです。ユーザーとアシスタントに加えて、添付ファイル エントリ (フック出力、診断、タスク リマインダー)、元に戻すファイル履歴デルタ レコード、ビジー中に入力されたメッセージのキュー操作エントリ、および時折 pr-link が表示されます。トランスクリプトのどれだけが会話ではないかを確認するために、最初にこれを実行する価値があります。
実際にターンはどこへ行くのでしょうか？
「ターンは実際にどこへ行くのですか?」というタイトルのセクション
各アシスタント エントリにはコンテンツ ブロックが 1 つだけ保持されるため、モデルがそのターンで何を行うかを数えることができます。
SELECT jget( "message/content" , '/0/type' ) block , count ( * ) n,round ( 100 . 0 * count ( * ) / ( SELECT count ( * ) FROM claude_code_session WHERE type = 'assistant' ), 1 ) pct FROM claude_code_session WHERE type = 'assistant' GROUP BY block ORDER n DESC まで
tool_use 、 Thinking 、 text の 3 つのブロック タイプが戻ってきました。最後の部分だけが散文で書かれており、この 3 つの比率はセッションの経過を驚くほどうまく要約しています。
「どのツールが失敗しますか?」というタイトルのセクション
これには結合が必要です。アシスタント エントリのtool_use ブロックは、それに応答するユーザー エントリのtool_result ブロックと一致します。
WITH tu AS MATERIALIZED ( SELECT jget( "message/content" , '/0/name' ) tools, jget( "message/content" , '/0/id' ) tuid FROM claude_code_session WHERE type = 'assistant

' AND jget( "message/content" , '/0/type' ) = 'tool_use' ), tr AS MATERIALIZED ( SELECT jget( "message/content" , '/0/tool_use_id' ) tuid, CASE WHEN jget( "message/content" , '/0/is_error' ) IN ( 1 , 'true' ) THEN 1 ELSE 0 END err FROM claude_code_session WHERE type = 'user' AND json_valid( "message/content" )) SELECT ツール、count ( * ) n、sum (err) エラー、round ( 100 . 0 * sum (err) / count ( * ), 1 ) err_pct FROM tu JOIN tr USING (tuid) GROUP BY ツールHAVING n > 20 err_pct DESC で注文 20 err_pct DESC で注文">
結果ブロックの is_error は、これをカウントからレートに変換するものです。 MCP ツールと Bash は上位に位置する傾向があります。 HAVING 句により、1 回限りのツールが 1 回の呼び出しで 100% の失敗率で独占されるのを防ぎます。
キャッシュの経済性というタイトルのセクション
トークンの使用量はメッセージごとに記録されるため、コスト全体の全体像は 1 つの GROUP BY になります。
SELECT "メッセージ/モデル" model, count ( * ) msgs, sum ( "message/usage/output_tokens" ) out_tok,round ( 100 . 0 * sum ( "message/usage/cache_read_input_tokens" ) / ( sum ( "message/usage/cache_read_input_tokens" ) + sum ( "message/usage/input_tokens" ) + sum ( "message/usage/cache_creation_input_tokens" )), 1 )cache_pct FROM claude_code_session WHERE type = 'assistant' AND "message/model" NOT LIKE '<%' GROUP BY model ORDER BY out_tok DESC
セッションは途中でモデルを切り替えることができ、<% フィルターはクロード コードが割り込みなどのために書き込む <synthetic> メッセージをドロップするため、モデルによるグループ化が重要になります。注目する列は、cache_pct です。プロンプト キャッシュがこれほど多くの作業を実行するとは予想していませんでした。
「最も遅いツール呼び出しは私です」というタイトルのセクション
前と同じ結合ですが、エラーをカウントする代わりに 2 つのタイムスタンプを減算します。
WITH tu AS MATERIALIZED ( SELECT log_time lt, jget( "message/content" , '/0/name' ) t

ool, jget( "message/content" , '/0/id' ) tuid, coalesce (jget( "message/content" , '/0/input/command' ), '' ) arg FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/type' ) = 'tool_use' ), tr AS MATERIALIZED ( SELECT log_time lt, jget( "message/content" , '/0/tool_use_id' ) tuid FROM claude_code_session WHERE type = 'user' AND json_valid( "message/content" )) SELECT tu 。 tool 、round ((julianday( tr . lt ) - julianday( tu . lt )) * 86400 . 0 , 1 ) 秒、substr( tu . arg , 1 , 28 ) arg FROM tu JOIN tr USING (tuid) ORDER BY secs DESC LIMIT 8
これを実行すると、ほぼ確実にリストの先頭にExitPlanModeとAskUserQuestionが表示されます。これらのツールの実装は、場合によっては翌朝人間が何かを決定することになります。独自の応答時間の測定値ではなく、実際のツールの待ち時間が必要な場合は、WHERE ツール NOT IN (...) を追加します。
「勝算と結末」というタイトルのセクション
実行されたすべてのシェル コマンドの最初の単語。これは、リポジトリがどのように探索されるかを示す適切な図です。
SELECT regexp_match( '^\s*(\w+)' , jget( "message/content" , '/0/input/command' )) cmd, count ( * ) n FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/name' ) = 'Bash' GROUP BY cmd HAVING cmd NOT NULL ORDER BY n DESC 制限 10
サブエージェントに委任される作業量 (これは -r が必要なクエリです):
SELECT 合体 (attributionAgent, '(main thread)' ) エージェント、カウント ( DISTINCT エージェントId) 実行、カウント ( * ) msgs、sum ( "message/usage/output_tokens" ) out_tok FROM claude_code_session WHERE type = 'assistant' GROUP BY エージェント ORDER BY msgs DESC
attributionAgent はエージェント タイプを示し、agentId は実行ごとに一意であるため、count(DISTINCT AgentId) は各エージェントが生成された回数を示します。メインスレッドには AgentId がないため、(メインスレッド) バックに配置されます。

t の実行カウントはゼロです。
そして、説明の必要のないものがさらにいくつかあります。
-- どのスキルが最も多く発動するか SELECT coalesce (attributionSkill, '(none)' ) skill, count ( * ) n FROM claude_code_session WHERE type = 'assistant' GROUP BY skill ORDER BY n DESC LIMIT 8 ;
-- いつ仕事をしますか SELECT strftime( '%H:00' , log_time) 時間 , count ( * ) msg​​s FROM claude_code_session WHERE type IN ( 'user' , 'assistant' ) GROUP BY 時間 ORDER BY 時間 ;
-- ファイル SELECT jget( "message/content" , '/0/input/file_path' ) file , count ( * ) touch FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/name' ) IN ( 'Edit' , 'Write' , 'Read' ) GROUP BY file HAVING file NOT NULL ORDER BY は DESC LIMIT 10 にタッチします。
-- ブランチ間で作業を分散 SELECT gitBranch, count (DISTINCT sessionId) session , count ( * ) msg​​s FROM claude_code_session WHERE gitBranch NOT NULL GROUP BY gitBranch ORDER BY msgs DESC LIMIT 10 ;
-- 最も遅いターン SELECT log_time、round (durationMs / 1000 . 0 , 1 ) secs、messageCount msgs、gitBranch FROM claude_code_session WHERE subtype = 'turn_duration' ORDER BY durationMs DESC LIMIT 8 ;
-- 最長ギャップ、つまり、SELECT ラウンド (log_idle_msecs / 3600000 . 0 , 1 ) idle_hours、log_time、type FROM claude_code_session WHERE log_idle_msecs > 3600000 ORDER BY log_idle_msecs DESC LIMIT 5 から立ち去ったとき。 3600000 ORDER BY log_idle_msecs DESC LIMIT 5;">
log_time 、 log_idle_msecs 、および log_opid は、すべての lnav 形式で無料で提供されます。最後のファイルは、上記のフォーマット ファイルの sessionId に関連付けられているため、多くのファイルを結合したビューを個別のセッションにグループ化することができます。
「さまざまなメモと調査結果」というタイトルのセクション
lnav は JSON 配列にインデックスを付けることができません。ネストされたオブジェクトは / 結合された列名にフラット化されます。これが、「message/usage/output_tokens」が機能する理由です (そして

二重引用符が必要な理由)。配列はそのような扱いを受けません。message/content/0/type 列がありません。配列を "kind": "json" として宣言し、 jget(col, '/0/name') でアクセスします。上記のすべてのクエリは、その 1 つの制限によって形作られています。
AS MATERIALIZED はオプションではありません。ツール呼び出し結合の両方の部分は、それ自体では正常に機能しますが、結合した瞬間に失敗し、JSON テキストに無効な文字が含まれており、どの行が原因であるかわかりません。 SQLite はサブクエリをフラット化し、最終的に相手側の WHERE 句が除外する行に対して jget を評価します。これには、message.content が配列ではなくプレーン文字列であるユーザー エントリも含まれます。 MATERIALIZED は境界を強制し、json_valid() が残りを処理します。
lnav -f script.lnav は最初に失敗したクエリで停止し、依然として 0 で終了します。クエリ 3 のタイプミスにより 10 個のクエリが失われましたが、気づきませんでした。
--匿名化は部分的です。すべての write-*-to コマンドはこれを受け入れ、ユーザー名とパスのプレフィックスを一貫して仮名化します。ホーム ディレクトリは、どの行でも同じ造語として表示されます。しかし、試してみるとリポジトリ名とファイル名はそのまま残りました。トランスクリプトには、あなたが取り組んでいたものがすべてピックアップされるため、フラグを信頼するのではなく、出力を公開する前に読んでください。

[切り捨てられた]

## Original Extract

lnav is a powerful log file viewer *and querier* and we can teach it about claude code sessions

Inspect Claude Code sessions with lnav | Georges Haidar
Top of page Site navigation
Georges Haidar
Inspect Claude Code sessions with lnav
For the most part, this article was LLM generated, especially the snippets and
examples.
I’m not sure if anyone will find this interesting but…
You can teach lnav about Claude Code session stored on your machines and then have it properly parse them and make them queryable.
claude-code.json { " $schema " : " https://lnav.org/schemas/format-v1.schema.json " , " claude_code_session " : { " title " : " Claude Code session transcript " , " description " : " JSONL session transcripts written by Claude Code under ~/.claude/projects/<encoded-cwd>/<session-uuid>.jsonl and .../<session-uuid>/subagents/agent-*.jsonl " , " url " : " https://docs.claude.com/en/docs/claude-code " , " json " : true , " file-pattern " : " \\ .claude/projects/.* \\ .jsonl$ " , " hide-extra " : true , " timestamp-field " : " timestamp " , " opid-field " : " sessionId " , " level-field " : " type " , " level " : { " debug " : " ^(attachment|file-history-delta|file-history-snapshot|queue-operation)$ " , " info " : " ^(user|assistant|started|result)$ " , " notice " : " ^(system|pr-link|frame-link|fork-context-ref)$ " } , " line-format " : [ { " field " : " __timestamp__ " } , " " , { " field " : " type " , " min-width " : 10 , " max-width " : 10 , " overflow " : " truncate " } , { " field " : " subtype " , " default-value " : "" , " prefix " : " [ " , " suffix " : " ] " } , { " field " : " operation " , " default-value " : "" , " prefix " : " [ " , " suffix " : " ] " } , { " field " : " attachment/type " , " default-value " : "" , " prefix " : " [ " , " suffix " : " ] " } , { " field " : " durationMs " , " default-value " : "" , " prefix " : " " , " suffix " : " ms " } , { " field " : " attributionSkill " , " default-value " : "" , " prefix " : " skill= " } , { " field " : " attributionAgent " , " default-value " : "" , " prefix " : " agent= " } , { " field " : " agentId " , " default-value " : "" , " prefix " : " agentId= " } , " " , { " field " : " message/content " , " default-value " : "" } , { " field " : " content " , " default-value " : "" } , { " field " : " prUrl " , " default-value " : "" } , { " field " : " trackingPath " , " default-value " : "" } , { " field " : " attachment " , " default-value " : "" , " max-width " : 400 , " overflow " : " truncate " } ] ,
[truncated]
lnav ~/.claude/projects/<project>/ * .jsonl /*.jsonl">
Replace <project> with a specific project or projects you work on or use * to index all sessions across all of your projects.
Section titled Querying without the TUI
The interactive viewer is nice, but the real fun is that lnav puts every log file behind SQLite. You can run a query and print the result straight to your terminal:
lnav -q -n \ -c ";SELECT type, count(*) n FROM claude_code_session GROUP BY type" \ -c ':write-table-to -' \ ~/.claude/projects/<project>/ * .jsonl /*.jsonl">
-n is headless mode, -q silences the startup chatter, the ; prefix marks a SQL command and :write-table-to - dumps the result set to stdout. There’s also write-csv-to , write-json-to , write-jsonlines-to and a few others if you want to pipe it somewhere. You can pass as many -c pairs as you like and they’ll all run against a single index, or put them in a file and use lnav -f queries.lnav .
Subagent transcripts live in a subagents/ subdirectory next to the session
file. Point lnav at the project directory with -r and it will pick up
everything recursively. Without it, every query below silently excludes your
subagents.
Everything below queries claude_code_session , the table lnav names after the format. I’ve left my own output out — point these at your transcripts instead, the shape of the answers is the interesting part.
Section titled So what’s in there?
SELECT type , count ( * ) n, round ( 100 . 0 * count ( * ) / ( SELECT count ( * ) FROM claude_code_session), 1 ) pct FROM claude_code_session GROUP BY type ORDER BY n DESC
A transcript is more than the conversation. Alongside user and assistant you’ll find attachment entries (hook output, diagnostics, task reminders), file-history-delta records backing undo, queue-operation entries for messages typed while it was busy, and the occasional pr-link . It’s worth running this one first just to see how much of a transcript isn’t the conversation.
Where do the turns actually go?
Section titled Where do the turns actually go?
Each assistant entry holds exactly one content block, so you can count what the model spends its turns doing:
SELECT jget( "message/content" , '/0/type' ) block , count ( * ) n, round ( 100 . 0 * count ( * ) / ( SELECT count ( * ) FROM claude_code_session WHERE type = 'assistant' ), 1 ) pct FROM claude_code_session WHERE type = 'assistant' GROUP BY block ORDER BY n DESC
Three block types come back: tool_use , thinking and text . Only the last one is prose written for you to read, and the ratio between the three is a surprisingly good summary of how a session went.
Section titled Which tools fail?
This one needs a join: a tool_use block in an assistant entry, matched to the tool_result block in the user entry that answers it.
WITH tu AS MATERIALIZED ( SELECT jget( "message/content" , '/0/name' ) tool, jget( "message/content" , '/0/id' ) tuid FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/type' ) = 'tool_use' ), tr AS MATERIALIZED ( SELECT jget( "message/content" , '/0/tool_use_id' ) tuid, CASE WHEN jget( "message/content" , '/0/is_error' ) IN ( 1 , 'true' ) THEN 1 ELSE 0 END err FROM claude_code_session WHERE type = 'user' AND json_valid( "message/content" )) SELECT tool, count ( * ) n, sum (err) errors, round ( 100 . 0 * sum (err) / count ( * ), 1 ) err_pct FROM tu JOIN tr USING (tuid) GROUP BY tool HAVING n > 20 ORDER BY err_pct DESC 20 ORDER BY err_pct DESC">
is_error on the result block is what turns this from a count into a rate. MCP tools and Bash tend to sit at the top; the HAVING clause keeps one-off tools from dominating with a 100% failure rate on a single call.
Section titled Cache economics
Token usage is recorded per message, so the whole cost picture is one GROUP BY :
SELECT "message/model" model, count ( * ) msgs, sum ( "message/usage/output_tokens" ) out_tok, round ( 100 . 0 * sum ( "message/usage/cache_read_input_tokens" ) / ( sum ( "message/usage/cache_read_input_tokens" ) + sum ( "message/usage/input_tokens" ) + sum ( "message/usage/cache_creation_input_tokens" )), 1 ) cache_pct FROM claude_code_session WHERE type = 'assistant' AND "message/model" NOT LIKE '<%' GROUP BY model ORDER BY out_tok DESC
Grouping by model matters because a session can switch models partway through, and the <% filter drops the <synthetic> messages Claude Code writes for things like interrupts. The column to look at is cache_pct — I was not expecting prompt caching to be doing quite as much of the work as it is.
Section titled The slowest tool calls are me
Same join as before, but subtracting the two timestamps instead of counting errors:
WITH tu AS MATERIALIZED ( SELECT log_time lt, jget( "message/content" , '/0/name' ) tool, jget( "message/content" , '/0/id' ) tuid, coalesce (jget( "message/content" , '/0/input/command' ), '' ) arg FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/type' ) = 'tool_use' ), tr AS MATERIALIZED ( SELECT log_time lt, jget( "message/content" , '/0/tool_use_id' ) tuid FROM claude_code_session WHERE type = 'user' AND json_valid( "message/content" )) SELECT tu . tool , round ((julianday( tr . lt ) - julianday( tu . lt )) * 86400 . 0 , 1 ) secs, substr( tu . arg , 1 , 28 ) arg FROM tu JOIN tr USING (tuid) ORDER BY secs DESC LIMIT 8
Run this and the top of the list will almost certainly be ExitPlanMode and AskUserQuestion — the tools whose implementation is a human deciding something, sometimes the next morning. Add a WHERE tool NOT IN (...) if you want actual tool latency rather than a measure of your own response time.
Section titled Odds and ends
The first word of every shell command it ran, which is a decent picture of how it explores a repo:
SELECT regexp_match( '^\s*(\w+)' , jget( "message/content" , '/0/input/command' )) cmd, count ( * ) n FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/name' ) = 'Bash' GROUP BY cmd HAVING cmd NOT NULL ORDER BY n DESC LIMIT 10
How much work gets delegated to subagents (this is the query that needs -r ):
SELECT coalesce (attributionAgent, '(main thread)' ) agent, count ( DISTINCT agentId) runs, count ( * ) msgs, sum ( "message/usage/output_tokens" ) out_tok FROM claude_code_session WHERE type = 'assistant' GROUP BY agent ORDER BY msgs DESC
attributionAgent gives you the agent type and agentId is unique per run, so count(DISTINCT agentId) tells you how many times each one was spawned. The main thread has no agentId , so it lands in the (main thread) bucket with a run count of zero.
And a few more that need no explanation:
-- which skills fire the most SELECT coalesce (attributionSkill, '(none)' ) skill, count ( * ) n FROM claude_code_session WHERE type = 'assistant' GROUP BY skill ORDER BY n DESC LIMIT 8 ;
-- when do I work SELECT strftime( '%H:00' , log_time) hour , count ( * ) msgs FROM claude_code_session WHERE type IN ( 'user' , 'assistant' ) GROUP BY hour ORDER BY hour ;
-- files I keep coming back to SELECT jget( "message/content" , '/0/input/file_path' ) file , count ( * ) touches FROM claude_code_session WHERE type = 'assistant' AND jget( "message/content" , '/0/name' ) IN ( 'Edit' , 'Write' , 'Read' ) GROUP BY file HAVING file NOT NULL ORDER BY touches DESC LIMIT 10 ;
-- work spread across branches SELECT gitBranch, count ( DISTINCT sessionId) sessions , count ( * ) msgs FROM claude_code_session WHERE gitBranch NOT NULL GROUP BY gitBranch ORDER BY msgs DESC LIMIT 10 ;
-- slowest turns SELECT log_time, round (durationMs / 1000 . 0 , 1 ) secs, messageCount msgs, gitBranch FROM claude_code_session WHERE subtype = 'turn_duration' ORDER BY durationMs DESC LIMIT 8 ;
-- longest gaps, i.e. when I walked away SELECT round (log_idle_msecs / 3600000 . 0 , 1 ) idle_hours, log_time, type FROM claude_code_session WHERE log_idle_msecs > 3600000 ORDER BY log_idle_msecs DESC LIMIT 5 ; 3600000 ORDER BY log_idle_msecs DESC LIMIT 5;">
log_time , log_idle_msecs and log_opid come free with every lnav format. The last one is wired to sessionId in the format file above, so a merged view of many files can still be grouped back into individual sessions.
Section titled Assorted notes and findings
lnav can’t index into JSON arrays. Nested objects flatten into / -joined column names, which is why "message/usage/output_tokens" works (and why it needs the double quotes). Arrays don’t get that treatment: there is no message/content/0/type column. Declare the array as "kind": "json" and reach into it with jget(col, '/0/name') . Every query above is shaped by that one limitation.
AS MATERIALIZED isn’t optional. Both halves of the tool-call join work fine on their own and fail the moment you join them, with invalid char in json text and no clue which row caused it. SQLite flattens the subqueries and ends up evaluating jget against rows that the other side’s WHERE clause would have excluded, including the user entries whose message.content is a plain string rather than an array. MATERIALIZED forces the boundary and json_valid() handles the rest.
lnav -f script.lnav stops at the first failing query and still exits 0. I lost ten queries to a typo in query three and didn’t notice.
--anonymize is partial. Every write-*-to command accepts it, and it consistently pseudonymises usernames and path prefixes — a home directory comes out as the same made-up word in every row. But it left repository names and filenames alone when I tried it. Transcripts pick up whatever you happened to be working on, so read your output before you publish it rather than trusting the flag

[truncated]

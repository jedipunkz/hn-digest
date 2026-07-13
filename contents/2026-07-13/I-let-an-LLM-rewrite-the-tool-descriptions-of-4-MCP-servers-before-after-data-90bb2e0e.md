---
source: "https://toolmetry-one.vercel.app/launch"
hn_url: "https://news.ycombinator.com/item?id=48891091"
title: "I let an LLM rewrite the tool descriptions of 4 MCP servers – before/after data"
article_title: "I let an LLM rewrite the tool descriptions of 4 popular MCP servers — the before/after data"
author: "tsuzuu"
captured_at: "2026-07-13T11:52:25Z"
capture_tool: "hn-digest"
hn_id: 48891091
score: 1
comments: 0
posted_at: "2026-07-13T11:33:16Z"
tags:
  - hacker-news
  - translated
---

# I let an LLM rewrite the tool descriptions of 4 MCP servers – before/after data

- HN: [48891091](https://news.ycombinator.com/item?id=48891091)
- Source: [toolmetry-one.vercel.app](https://toolmetry-one.vercel.app/launch)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T11:33:16Z

## Translation

タイトル: LLM に 4 台の MCP サーバーのツール記述を書き換えさせました – データの前後
記事のタイトル: LLM に 4 つの人気のある MCP サーバーのツールの説明を書き直させました - 変更前/変更後のデータ
説明: 説明のみを変更した場合、エージェントのパフォーマンスはどれくらい向上しますか? 4 つの公式 MCP サーバーで N=5 を測定: +66、+34.5、+21.7、+10 ポイントの厳密な成功。

記事本文:
ツールメトリ
家
GitHub
npm
フィールドノート · 2026-07-13
LLM に 4 つの人気のある MCP サーバーのツールの説明を書き直させました。こちらがBefore/Afterのデータです。
Sujay 著 · 以下のすべての数値は、コミットされた実行データから再現可能です。
エージェントはサーバーのコードを読み取りません。彼らは、ツールの名前、説明、およびパラメーター スキーマの 3 つの文字列を使用してツールを選択します。それがインターフェース全体です。 2 つの説明が重複する場合、エージェントは推測します。説明に「親ディレクトリも作成される」と記載されていない場合、エージェントはそれを 4 回呼び出します。これらの間違いはどれもサーバーが不安定であるように見えます。
私は単純な質問に答えるために Toolmetry を作成しました: 説明だけを変更した場合、エージェントはどれくらい改善されるでしょうか?
サーバーごとに 10 ～ 18 個の現実的なシナリオを作成します (プロンプト + 適切に設計されたサーバーが引き出すツール + 予想される引数)。
実際のエージェント ループを使用して、ライブ サーバーに対して各シナリオを 5 回実行します。実際のツール呼び出しをすべて記録します。
ヒット率 (適切なツールか?)、引数の正確性、および余分な呼び出し (タスクに不必要な呼び出しが埋め込まれていないか) の 3 つの項目をスコアリングします。 「厳密な成功」 = 3 つすべてを同時に。
障害と現在の説明を LLM リライターにフィードします。書き換えられた記述をメモリ内に適用します。サーバーは決して変更されません。
再測定してください。目に見えて改善された場合にのみ、書き換えを維持してください。回帰を破棄します。
エージェント: gpt-oss-120b (意図的に中間層エージェント - 理由については以下で詳しく説明します)。リライター：キミK2。この投稿のすべての API 費用の合計: 約 $4 。
3 つの異なる障害の原型が表示され、説明の書き換えにより 3 つすべてが修正されました。
1. 間違ったツールの混乱 (メモリサーバー)。ナレッジグラフ メモリ サーバーには create_entities 、 add_observations 、 search_nodes 、 open_nodes があり、ベースライン エージェントは 20% の確率でこれらを混乱させました。リライター

明示的な「…の場合は代わりに X を使用する」相互参照を追加しました。命中率が80%→100%になりました。
2. 儀式的な追加呼び出し (sqlite、git)。 sqlite エージェントは、クエリの 3 分の 2 (「ユーザーは何人いるのか?」も含む) の前に list_tables + description_table を呼び出しました。追加された 1 つの文 (「クエリを実行する前にテーブルの存在を確認するためだけにこれを呼び出さないでください」) により、余分な呼び出し率が 66% からゼロになりました。厳密な成功: 34% → 100% 。
3. 非推奨のエイリアス トラップ (ファイル システム)。ファイルシステム サーバーには、非推奨の read_file が同梱されていますが、その説明は依然として主要ツールと同様です。エージェントたちはその中を歩き続けた。修正方法: 非推奨を最初の単語にして、置き換える単語に名前を付けます。
代わりに、Claude Haiku 4.5 をエージェントとして使用して同じ最適化を実行しました。そのベースライン (84.4%) は中間層モデルの最適化スコアとほぼ同じであり、最適化によって +2.2 ポイント増加しただけです。
優れたエージェントは、あなたの悪い説明を回避します。弱いエージェントにはそれができません。つまり、大量の作業のために人々が導入するエージェント、つまり安価で高速なエージェントにとって、記述の品質はまさに負担となります。 MCP サーバーが「安価なモデルでは不安定」である場合、これが理由である可能性があり、1 ドル未満で修正可能です。
もう 1 つの正直な発見は、LLM の書き換えは分散が大きいということです。同じベースラインに対する 2 つの独立した書き換え試行では、+10.0 ポイントと -2.2 ポイントのスコアが得られました。これを一度で解決することはできません。測定する必要があり、書き換えを放棄する覚悟が必要です。 (Toolmetry のループはこれを自動的に行います。これらの実行中に回帰が 2 回破棄されました。)
書き換えられた説明は JSON ファイル内に存在します。 toolmetry プロキシは、任意の MCP サーバーをラップし、そのツール/リストの応答をオンザフライで書き換えます。
npx mcp-toolmetry プロキシ --overrides best.json -- uvx mcp-server-sqlite --db-path ./my.db
MCP クライアントをサーバーではなくそこに向けます。フォークなし、パッチなし、オンでリバーシブル

e行。
注意事項（注意事項のないデータはマーケティングであるため）
シナリオごとに N=5 では、レートが 20 ポイントのステップに量子化されます。個々のシナリオの差分ではなく、集計を信頼します。
シナリオはグラウンド トゥルースをエンコードします。議論の余地のある Expected_tool は議論の余地のあるメトリクスを作成します。 51 のシナリオすべてがリポジトリにあります。ご自身で判断してください。
デルタはエージェント固有です (Haiku の結果を参照)。実際にサービスを提供しているエージェント層で測定してください。
一部の失敗はツールの設計上の問題であり、説明では解決できません。シナリオごとのレポートには、それが示されます。
すべてのコード、シナリオ、実行ごとの結果、および勝者の説明の差分: github.com/sujugithub/toolmetry — またはランディング ページから開始します。お気に入りのサーバーのシナリオ スイートを使って PR していただければ幸いです。測定されたサーバーの 1 つを保守している場合は、説明の差分を取得できます (メモリと git の書き換えはすでにサーバー#4519 およびサーバー#4520 として上流にあります)。

## Original Extract

If you change nothing but the descriptions, how much better do agents get? Measured N=5 on four official MCP servers: +66, +34.5, +21.7, +10 points strict success.

Toolmetry
home
GitHub
npm
Field notes · 2026-07-13
I let an LLM rewrite the tool descriptions of 4 popular MCP servers. Here's the before/after data.
By Sujay · every number below is reproducible from the committed run data
Agents don't read your server's code. They pick tools using three strings: the tool's name , its description , and its parameter schema . That's the entire interface. If two descriptions overlap, the agent guesses. If a description doesn't say "this creates parent directories too," the agent calls it four times. Every one of those mistakes looks like your server being flaky .
I built Toolmetry to answer a simple question: if you change nothing but the descriptions, how much better do agents get?
Write 10–18 realistic scenarios per server (a prompt + the tool a well-designed server should elicit + expected args).
Run each scenario 5 times against the live server with a real agent loop; record every actual tool call.
Score three things: hit rate (right tool?), arg correctness , and extra calls (did it pad the task with unnecessary calls?). "Strict success" = all three at once.
Feed the failures + current descriptions to an LLM rewriter. Apply the rewritten descriptions in-memory — the server is never modified.
Re-measure. Keep the rewrite only if it measurably improved. Discard regressions.
Agent: gpt-oss-120b (a deliberately mid-tier agent — more on why below). Rewriter: Kimi K2. Total API spend for everything in this post: about $4 .
Three different failure archetypes showed up, and description rewrites fixed all three:
1. Wrong-tool confusion (memory server). The knowledge-graph memory server has create_entities , add_observations , search_nodes , open_nodes — and the baseline agent confused them 20% of the time. The rewriter added explicit "use X instead when…" cross-references. Hit rate went 80% → 100% .
2. Ritual extra calls (sqlite, git). The sqlite agent called list_tables + describe_table before two-thirds of queries — even "how many users are there?". One added sentence ("Do not call this merely to check that a table exists before querying it") took the extra-call rate from 66% to zero. Strict success: 34% → 100% .
3. Deprecated-alias traps (filesystem). The filesystem server ships a deprecated read_file whose description still reads like the primary tool. Agents kept walking into it. The fix: make the deprecation the first word and name the replacement.
I ran the same optimization with Claude Haiku 4.5 as the agent instead. Its baseline (84.4%) roughly equaled the mid-tier model's optimized score — and optimization only moved it +2.2 pts.
Better agents route around your bad descriptions. Weaker agents can't. Which means description quality is a tax on exactly the agents people deploy for high-volume work — the cheap, fast ones. If your MCP server is "flaky with cheap models," this might be why, and it's fixable for under a dollar.
The other honest finding: LLM rewrites are high-variance. Two independent rewrite attempts on the same baseline scored +10.0 and −2.2 points. You cannot one-shot this — you have to measure, and you have to be willing to throw a rewrite away. (Toolmetry's loop does this automatically; it discarded regressions twice during these runs.)
The rewritten descriptions live in a JSON file. toolmetry proxy wraps any MCP server and rewrites its tools/list responses on the fly:
npx mcp-toolmetry proxy --overrides best.json -- uvx mcp-server-sqlite --db-path ./my.db
Point your MCP client at that instead of the server. No fork, no patch, reversible in one line.
Caveats, because data without caveats is marketing
N=5 per scenario quantizes rates to 20-pt steps; trust the aggregates, not individual scenario deltas.
Scenarios encode ground truth; a debatable expected_tool makes a debatable metric. All 51 scenarios are in the repo — judge them yourself.
Deltas are agent-specific (see the Haiku result). Measure with the agent tier you actually serve.
Some failures are tool- design problems no description can fix. The per-scenario report shows you which.
All code, scenarios, per-run results, and the winning description diffs: github.com/sujugithub/toolmetry — or start at the landing page . I'd love PRs with scenario suites for your favorite server — and if you maintain one of the measured servers, the description diffs are yours for the taking (the memory and git rewrites are already upstream as servers#4519 and servers#4520 ).

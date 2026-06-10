---
source: "https://christophermeiklejohn.com/ai/zabriskie/agents/reliability/testing/2026/06/10/the-test-suite-was-the-incident.html"
hn_url: "https://news.ycombinator.com/item?id=48474664"
title: "The Test Suite Was the Incident"
article_title: "The Test Suite Was the Incident · A night of brittle fixtures, 49 failed CI runs, and an expensive lesson about what happens when AI agents write the test suite and nobody owns the test data. | Christopher Meiklejohn"
author: "jruohonen"
captured_at: "2026-06-10T13:19:50Z"
capture_tool: "hn-digest"
hn_id: 48474664
score: 1
comments: 0
posted_at: "2026-06-10T11:23:19Z"
tags:
  - hacker-news
  - translated
---

# The Test Suite Was the Incident

- HN: [48474664](https://news.ycombinator.com/item?id=48474664)
- Source: [christophermeiklejohn.com](https://christophermeiklejohn.com/ai/zabriskie/agents/reliability/testing/2026/06/10/the-test-suite-was-the-incident.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T11:23:19Z

## Translation

タイトル: テストスイートが事件だった
記事のタイトル: テスト スイートが事件だった · 脆弱なフィクスチャ、49 件の CI 実行の失敗、そして AI エージェントがテスト スイートを作成し、誰もテスト データを所有しない場合に何が起こるかについての高価な教訓の一夜。 |クリストファー・メイクルジョン
説明: 脆弱なフィクスチャ、49 件の CI 実行の失敗、そして AI エージェントがテスト スイートを作成し、誰もテスト データを所有していない場合に何が起こるかについての高価なレッスンの夜。

記事本文:
コンテンツにスキップ
クリストファー・メイクルジョン
アーカイブ
研究
教える
出版物
履歴書
テストスイートが事件の原因だった
脆弱なフィクスチャ、49 件の CI 実行の失敗、そして AI エージェントがテスト スイートを作成し、誰もテスト データを所有していない場合に何が起こるかについての高価な教訓の夜。
昨夜の午後 10 時 EDT に、私は Codex の 100 ドルのプランにサインアップしました。Codex は、その夜の作業のほとんどを主導していた OpenAI コーディング エージェントです。約 12 個の Zabriskie PR が飛行中だったので、寝る前にそれらをマージしたかったからです。午前 3 時前に、200 ドルのプランにアップグレードしました。その中間で、GitHub の使用状況ページは 200 ドルの CI クォータの約 90% を達成し、プル リクエストとは関係のない理由で次々とプル リクエストが赤くなるのを座って見ていました。
読者が新鮮に受け取れるよう、いくつかのコンテキストを提供します。 Zabriskie は、ライブ音楽ファンのためのソーシャル アプリです。ショーへの参加表明、参加したショーの追跡、聴いている内容についての投稿、ツアーの統計情報を友人と比較できます。ほとんど何も書いていません。 AI エージェントは機能を構築し、AI エージェントはそれらを保護するテスト スイートも構築しました。仕様、共有シード データ、フィクスチャ ユーザー、ヘルパー、そのほとんどすべては、別のことに関するセッションでエージェントによって書き込まれました。この 2 番目の部分がこの投稿の主題全体になります。
GitHub Actions には何も問題はありませんでした。アプリケーションには致命的な問題は何もありませんでした。特にリスクを伴う PR は 1 つもありませんでした。すべての PR は、同じ脆弱な世界をゼロから再構築するために単純にお金を払っていました。バックエンドのビルド、新しい Postgres、完全なデータベース移行チェーン、共有シード データ、フロントエンド、および 8 つの並列シャードに分割されたエンドツーエンドのブラウザー テスト (Playwright) です。その世界は、誰も所有していない備品によって結び付けられていました。私はほぼ200ドルを費やしました

昨夜、私のテストスイートが私に嘘をついていたことを証明しました。
GitHub の実行履歴から測定された翌朝の数値:
観察されたランナー時間全体で実際に消費された 180 ドルを割り当てると、失敗したジョブだけで約 52 ドル、失敗したジョブとキャンセルされたジョブの合計で約 60 ドルのコストがかかります。この数字は技術的には真実ですが、感情的には役に立ちません。緑色の実行が存在するのは、赤色の実行によって別のコミットが強制されたためです。再実行が存在したのは、失敗が 1 つの共有前提から別の前提に移り続けたためです。チェックが成功したからといって、それが健康であることを証明するものではありません。それらは、スイートをまだ許容されている数少ない形状の1つに押し戻した後に支払われた税金でした。つまり、正直な計算は次のとおりです。直接起因する CI の無駄は約 60 ドルで、不適切なテスト設計による実際のコストは 180 ドルの大部分と夕方に、同じ構造的問題の症状をデバッグするようエージェントに依頼して費やした 200 万トークン程度でした。テストは私を守ってくれませんでした。彼らは私に家賃を請求していました。
実際の申請作業は大惨事ではありませんでした。実際の機能、実際の修正、プル リクエストがあり、それらは小さく、レビュー可能でマージ可能であるべきでした。代わりに、全員が同じグラインダーに入り、失敗の最初の波は失敗した PR とは何の関係もありませんでした。
誕生日機能のデータ修正である移行の 1 つは、移行中のデータベースに queenoftheman という名前のユーザーがすでに存在していることを想定していましたが、新しい CI データベースには存在していませんでした。
誕生日の名誉の修正が 0 行に一致しました
username/display_name = queenoftheman のユーザーはいません
別の一連の失敗は、プロジェクトのインシデント ログに対する移行によって発生しました。 Zabriskie のリポジトリは、独自の運用インシデントを追跡し、各インシデントが無駄にしたものの必要な推定値を完備しており、古いインシデント行は

これらの無駄の見積もりは含まれていないため、移行によりデータベースの制約が適用されませんでした。これらは製品の回帰ではありません。これらは移行動作に漏れ出るフィクスチャーの仮定でした。PR はアプリの片隅にありながら、グローバル テスト ユニバースの他の部分がドリフトしたため、依然として失われる可能性があります。
その後、問題全体の原因となった衝突が起こりました。 Zabriskie には、Lot と呼ばれる機能があり、これは、あなたが行ったショーを要約したツアー統計カードを備えたたまり場スペースです。過去のカード数のみがユーザーが実際に参加したことを示すように変更されました。これは製品の正しい動作であり、それを主張するためにテストが更新されました。このテストのセットアップでは、スイートの半分が必要なシナリオに借用する単一のグローバル ユーザーである e2etester の show RSVP を追加しました。一方、スイート内の他の場所でショーの要約の仕様では、その備品ショーの出席者が 1 人であることが予想されていました。要約テストが失敗し始めたのは、要約が壊れていたからではなく、共有世界が密かに 1 人余分に獲得したためです。
誰もデザインしたわけではないので、スイートがどのようにしてこのような形になるのかについては、立ち止まってみる価値があります。テストを作成するエージェントには、1 つの機能と 1 つの目標 (この仕様を合格させる) をスコープとするコンテキスト ウィンドウがあります。グリーンへの最も安価な方法は、既存のフィクスチャ ワールドを再利用することです。そのため、エージェントは e2etester を取得し、必要な RSVP を追加し、仕様のパスを監視して次に進みます。これらの選択はいずれも局所的には合理的です。しかし、どのセッションもフィクスチャー ユニバース自体に関するものではないため、e2etester が 5 つの異なる仕様に対して 5 つの異なる意味を持っていることをエージェントが認識することはありません。これは、「クロードは重要な 1 つのことを除いてすべてをテストしました」で書いたのと同じダイナミクスです。エージェントは、コンテキストが新鮮なセッションで、テストが最も簡単な場所にテストを作成します。彼らはフィクストゥを構築します

世界も同じように。その結果、全体像を把握した著者がいないまま、設計されるのではなく蓄積されるグローバルなソーシャル グラフが作成されます。
これは、テストを信頼するのをやめるように人々に教える特定の種類の失敗です。ランダムな意味で不安定ではありません。それは決定的であり、毎回失敗し、レビュー中の変更とは程遠い理由で失敗します。即時の修正は簡単でした。ロットシナリオ専用のフィクスチャデータと、古い要約 RSVP のクリーンアップでした。修正よりも教訓の方が大きかったです。このスイートは、あたかも中立的なインフラストラクチャであるかのように、グローバル ソーシャル グラフに依存していました。それは中立ではありませんでした。緑色のチェックマークが付いた共有可変状態でした。
同じ虫が違う服を着ている
数時間後、ザブリスキーのセットリストと検索に使用されているバンドと曲のカタログである曲データベースを通じて同じパターンが返されました。移行により、曲の一意性モデルが (band_id, Song_name) の制約から、カバー アーティストを考慮した表現インデックス (band_id, Song_name, COALESCE(original_artist, '')) に変更されました。合理的なスキーマ変更。しかし、E2E シード ファイルには ON CONFLICT (band_id, songs_name) DO UPDATE という挿入が依然として含まれており、Postgres は Postgres が行うべきことを正確に実行しました。
エラー: ON CONFLICT 指定に一致する一意制約または除外制約がありません
これは謎の CI 問題でもありませんでした。スキーマと共有シードの間で漂流していました。シードは事実上アプリケーション コントラクトの一部でしたが、シードを 1 つとして扱うものはありませんでした。実際の所有者はなく、早期に失敗することもなく、一度も失敗することはありませんでした。どこでも同時に失敗しました。曲検索では、Umphrey’s McGee の「All In Time」を見つけることができませんでした。これは、それを作成するはずだったシード インサートが失敗したためです。

The Lot のコール・ザ・オープナー ゲームは、ユーザーに番組のオープナーとして予想する曲を提供するものですが、同じ理由で未定義の候補を返すようになりました。複数の開いている PR が同じ根本的な原因で赤になり、その時点でテスト スイートは独立したチェックのセットではなくなりました。それは、1 つの前提が崩れたためのブロードキャスト メカニズムでした。
最悪の部分は、単一の失敗ではありませんでした。それは、アーキテクチャがそれぞれの失敗に対して行ったことでした。 E2E マトリックスは最近、1 つのシャードの成長が遅すぎたため、6 シャードから 8 シャードに再バランスされましたが、これは実際の圧力に対する妥当な対応です。しかし、フィクスチャの世界が不安定である一方で、8 シャード マトリックスにより、あらゆるミスが有料の分散イベントに変わりました。再実行するたびに、さらに多くのデータベースが開始されました。すべてのシャードで完全な移行チェーンが再生されました。無関係なすべてのブランチが同じ壊れた前提を個別に発見し、すべての発見が同じループを開始しました。つまり、PR が失敗し、エージェントが調査し、目に見える症状にパッチが適用され、CI が再度実行され、同じ共通の理由で別の PR が失敗し、別のエージェントが調査し、スイートがもう少し複雑になり、CI が再び実行されます。
そのループにはお金、時間、注意、トークンが費やされました。エージェントの使用量は各 CI の障害に明確に関連付けられていないため、完璧なトークン アカウンティングは存在しませんが、PR ループが壊れるたびに、ログの読み取り、仮説の生成、パッチの適用、再実行、説明に数万のトークンがかかります。 49 回の失敗とその後の修復作業を考慮すると、推定では 170 万から 250 万のトークンが得られます。私の最高のシングル番号は約210万です。正確な数字は、その形状よりも重要です。不適切なテスト設計は CI 時間を無駄にするだけではありません。すべての開発者とすべてのエージェントを分散型再試行システムに変換します。
その後、フィックスが事件に加わりました
深夜のミティ

すべての E2E シャードで 900 を超える移行の再実行を停止するという責任あるエンジニアリングの措置であるはずでした。データベースを 1 回構築してダンプし、そのダンプを各シャードに復元して、移行税を 1 回だけ支払います。前提は、試行されたときの両方で正しかった。最初の試みである PR #1112「移行された DB スナップショットによる E2E シャードの高速化」は、午前 12 時 44 分に開始され、19 分後に終了しましたが、破棄するには十分な間違いでした。
2 番目の試行である PR #1119 は午前 1 時 58 分に開始され、2 時 32 分までに 5 件のコミットが行われました。そのメッセージが独自に物語っています。つまり、移行のベースライン化、次に不完全なベースライン移行の失敗、次に廃止された曲データベースの更新/挿入キーでの高速フェイル、次にベースライン マニフェストのフィンガープリントの更新、そして履歴バックフィルの強化です。 6 番目のコミットが続きました。メインからのマージですが、これは進行状況ではなく、ブランチがマージ トレインに遅れるほど長く存続した場合にブランチに必要な記録だけを行いました。
根本的な技術的欠陥は微妙ではありませんでした。ベースライン ジェネレーターは、履歴移行チェーン全体を厳密にゼロから再生できることを前提としていました。リポジトリは実際には、壊れた古い履歴を許容、スキップ、または回避する移行ランナー上で存続していましたが、その履歴をベースライン アーティファクトに凍結することで、すべての古い前提が一度にクリティカル パス上に置かれました。移行 305 には、過去の UPDATE ... FROM ステートメントに Postgres スコープ エラーがありました。誕生日の栄誉者の移行では、queenoftheman が新しいデータベースに存在すると仮定しました。インシデントの移行により、制約が再び解除されました。リポジトリの PR にコメントする自動レビュー担当者である Bugbot は、ベースライン ダンプによってスキップされた移行が黙って許可され、移行 CLI が失敗した移行を無視できることを発見しました。 1 つの CI 実行はダンプの構築中に死亡し、もう 1 つは男性の検証中に死亡しました

ifest を実行し、3 番目のシャードはベースラインを通過し、8 つのシャードすべてを開始しましたが、それでもシャード 8 で失敗しました。これらはエッジ ケースではありませんでした。これらは、パス #1119 が凍結しようとしているが、凍結できるほどきれいではないという証拠でした。
そして、すぐに作業を終了すべきだった数字が目に見えるところに残っていました。ベースラインの生成には、以前の完全な CI 実行が約 13 分だったのに対し、約 11 分かかりました。セットアップ フェーズが古いエンドツーエンドのランタイムとほぼ同じである最適化は、まだ最適化ではありません。それは仮説であり、費用がかかります。その代わりに、私が優先事項はスイートを高速化することではなく、残りの PR をマージしてモバイル クライアントを出荷することであると明言している間、ループは一度に 1 つずつパッチを当てた症状で、夜の最も貴重な時間まで続きました。
午前 2 時 32 分、#1119 は独自の E2E DB ベースライン チェックでブロックされ、単体テストとその周囲のバックエンド ビルドがグリーンになりました。それでその夜は終わるはずだった。そうではありませんでした。他のマージ競合が解決され、キューを移動させるためにブランチ保護が一時的に緩和された後、エージェントはとにかく午前 2 時 51 分に #1119 をマージしました。これは、PR がすでに機能しないことを実証しており、私がそれが壊れているものであると特定した後でした。エージェントは「t

[切り捨てられた]

## Original Extract

A night of brittle fixtures, 49 failed CI runs, and an expensive lesson about what happens when AI agents write the test suite and nobody owns the test data.

Skip to content
Christopher Meiklejohn
archive
research
teaching
publications
cv
The Test Suite Was the Incident
A night of brittle fixtures, 49 failed CI runs, and an expensive lesson about what happens when AI agents write the test suite and nobody owns the test data.
At 10:00 PM EDT last night I signed up for the $100 plan for Codex, the OpenAI coding agent that was driving most of the night’s work, because about a dozen Zabriskie PRs were in flight and I wanted them merged before I went to bed. Before 3:00 AM I had upgraded to the $200 plan. Somewhere in between, GitHub’s usage page ticked through roughly 90% of a $200 CI quota, and I sat there watching pull request after pull request go red for reasons that had nothing to do with the pull requests.
Some context for readers arriving fresh. Zabriskie is a social app for live-music fans: you RSVP to shows, track the ones you attended, post about what you’re hearing, and compare tour stats with friends. I have written almost none of it. AI agents built the features, and AI agents also built the test suite that guards them: the specs, the shared seed data, the fixture users, the helpers, nearly all of it written by some agent in some session that was about something else. That second part turns out to be the entire subject of this post.
Nothing was wrong with GitHub Actions. Nothing was catastrophically wrong with the application. No single PR was especially risky. Every PR was simply paying to rebuild the same brittle world from scratch: backend build, fresh Postgres, the full database migration chain, the shared seed data, the frontend, and the end-to-end browser tests (Playwright) split into eight parallel shards. That world was held together by fixtures nobody owned. I spent almost two hundred dollars last night proving that my test suite was lying to me.
The morning-after numbers, measured from the GitHub run history:
If I allocate the $180 actually consumed across the observed runner time, the failed jobs alone cost about $52, and the failed plus cancelled jobs about $60. That number is technically true and emotionally useless. The green runs existed because the red runs forced another commit. The reruns existed because the failure kept moving from one shared assumption to another. The successful checks weren’t clean proof of health; they were the tax paid after pushing the suite back into one of the few shapes it still tolerated. So the honest accounting is this: the directly attributable CI waste was about sixty dollars, and the practical cost of the bad test design was most of the $180, plus the evening, plus something like two million tokens spent asking agents to debug symptoms of the same structural problem. The tests were not protecting me. They were charging me rent.
The actual application work was not the disaster. There were real features, real fixes, and pull requests that should have been small, reviewable, and mergeable. Instead, every one of them entered the same grinder, and the first wave of failures had nothing to do with the PRs that were failing.
One migration, a data fix for a birthday feature, expected a user named queenofthemean to already exist in the database it was migrating, and in a fresh CI database she didn’t:
birthday-honoree fix matched 0 rows
no user with username/display_name = queenofthemean
Another batch of failures came from migrations over the project’s incident log. Zabriskie’s repository tracks its own operational incidents, complete with required estimates of what each one wasted, and the old incident rows didn’t include those waste estimates, so the migrations tripped a database constraint. These weren’t product regressions. They were fixture assumptions leaking into migration behavior: a PR could be about one corner of the app and still lose because some other part of the global test universe had drifted.
Then came the collision that named the whole problem. Zabriskie has a feature called the Lot, a hangout space with a tour stats card that summarizes the shows you’ve been to. A change made that card count only past shows the user had actually attended, which was the correct product behavior, and a test was updated to assert it. The setup for that test added a show RSVP for e2etester , the single global user that half the suite borrows for whatever scenario it needs. Meanwhile, a spec for show recaps elsewhere in the suite expected its fixture show to have exactly one attendee. The recap test started failing, not because recaps were broken, but because the shared world had quietly acquired one extra person.
It’s worth pausing on how a suite ends up shaped like this, because nobody designed it. An agent writing a test has a context window scoped to one feature and one goal: make this spec pass. The cheapest path to green is to reuse whatever fixture world already exists, so the agent grabs e2etester , adds the RSVP it needs, watches the spec pass, and moves on. Every one of those choices is locally reasonable. But no session is ever about the fixture universe itself, so no agent ever sees that e2etester now means five different things to five different specs. This is the same dynamic I wrote about in Claude Tested Everything Except the One Thing That Mattered : agents write tests where the tests are easiest, in the session where the context is fresh. They build fixture worlds the same way. The result is a global social graph that accreted instead of being designed, with no author who ever held the whole thing in view.
This is the specific kind of failure that teaches people to stop trusting tests. It isn’t flaky in the random sense; it’s deterministic, it fails every time, and it fails for a reason that is nowhere near the change under review. The immediate fix was easy: dedicated fixture data for the Lot scenario and cleanup of the stale recap RSVPs. The lesson was bigger than the fix. The suite had been relying on a global social graph as if it were neutral infrastructure. It was not neutral. It was shared mutable state with a green checkmark on it.
The Same Bug, Wearing Different Clothes
A few hours later the same pattern came back through the song database, the catalog of bands and songs that powers Zabriskie’s setlists and search. A migration changed the uniqueness model for songs from a constraint on (band_id, song_name) to an expression index that accounted for cover artists: (band_id, song_name, COALESCE(original_artist, '')) . A reasonable schema change. But the E2E seed file still contained inserts that said ON CONFLICT (band_id, song_name) DO UPDATE , and Postgres did exactly what Postgres should do:
ERROR: there is no unique or exclusion constraint matching the ON CONFLICT specification
This wasn’t a mysterious CI problem either. It was drift between the schema and the shared seed. The seed was effectively part of the application contract, but nothing treated it like one: it had no real owner, it didn’t fail early, and it didn’t fail once. It failed everywhere, simultaneously. Song search couldn’t find “All In Time” for Umphrey’s McGee, because the seed insert that should have created it had failed. The Lot’s call-the-opener game, which offers users songs to predict as a show’s opener, started returning undefined suggestions for the same reason. Multiple open PRs went red for the same underlying cause, and at that point the test suite was no longer a set of independent checks. It was a broadcast mechanism for one broken assumption.
The worst part wasn’t any single failure; it was what the architecture did to each failure. The E2E matrix had recently been rebalanced from six shards to eight because one shard had grown too slow, a reasonable response to a real pressure. But while the fixture world was unstable, the eight-shard matrix turned every mistake into a paid distributed event. Every rerun started more databases. Every shard replayed the full migration chain. Every unrelated branch discovered the same broken assumptions independently, and every discovery kicked off the same loop: the PR fails, an agent investigates, the visible symptom gets patched, CI runs again, a different PR fails for the same shared reason, another agent investigates, the suite gets a little more elaborate, CI runs again.
That loop spent money, time, attention, and tokens. There’s no perfect token accounting, because the agent usage isn’t tied cleanly to each CI failure, but each broken-PR loop cost tens of thousands of tokens in log reading, hypothesis generation, patching, rerunning, and explaining. Across 49 failed runs and the follow-on repair work, the estimate lands between 1.7 and 2.5 million tokens; my best single number is about 2.1 million. The precise figure matters less than the shape of it. Bad test design doesn’t just waste CI minutes. It converts every developer and every agent into a distributed retry system.
Then the Fix Joined the Incident
The late-night mitigation was supposed to be the responsible engineering move: stop replaying 900-plus migrations in every E2E shard. Build the database once, dump it, restore the dump into each shard, pay the migration tax exactly once. The premise was right both times it was tried. The first attempt, PR #1112, “Speed up E2E shards with a migrated DB snapshot,” opened at 12:44 AM and was closed nineteen minutes later, wrong enough to throw away.
The second attempt, PR #1119, opened at 1:58 AM, and by 2:32 it had five commits whose messages tell the story on their own: baseline the migrations, then fail incomplete baseline migrations, then fail fast on a retired song-database upsert key, then refresh the baseline manifest fingerprint, then harden a historical backfill. A sixth commit followed: a merge from main, which wasn’t progress, just the bookkeeping a branch needs once it has lived long enough to fall behind the merge train.
The underlying technical failure was not subtle. The baseline generator assumed the whole historical migration chain could be replayed strictly from zero. The repository had actually been surviving on a migration runner that tolerated, skipped, or worked around old broken history, and freezing that history into a baseline artifact put every old assumption on the critical path at once. Migration 305 had a Postgres scoping error in a historical UPDATE ... FROM statement. The birthday-honoree migration assumed queenofthemean existed in a fresh database. The incident migrations tripped their constraints again. Bugbot, the automated reviewer that comments on the repo’s PRs, found that the baseline dump could silently allow skipped migrations and that the migration CLI could ignore failed ones. One CI run died building the dump, another died validating the manifest, and a third made it all the way through the baseline, started all eight shards, and still failed shard 8. Those weren’t edge cases. They were proof that the path #1119 wanted to freeze was never clean enough to freeze.
And there was a number sitting in plain sight that should have ended the effort immediately: the baseline generation took about eleven minutes, against a previous full CI run of about thirteen. An optimization whose setup phase nearly equals the old end-to-end runtime is not an optimization yet. It is a hypothesis, and an expensive one. Instead the loop kept going, one patched symptom at a time, through the most precious hour of the night, while I was explicitly saying that the priority was merging the remaining PRs and shipping the mobile client, not making the suite faster.
At 2:32 AM, #1119 was blocked on its own E2E DB Baseline check, with unit tests and the backend build green around it. That should have been the end of it for the night. It wasn’t. After the other merge conflicts were resolved and branch protection had been temporarily relaxed to get the queue moving, the agent merged #1119 anyway at 2:51 AM, after the PR had already demonstrated it didn’t work, and after I had identified it as the broken one. The agent treated “t

[truncated]

---
source: "https://christophermeiklejohn.com/ai/zabriskie/agents/reliability/testing/ci/distributed/2026/07/27/one-writer.html"
hn_url: "https://news.ycombinator.com/item?id=49075992"
title: "Our toolchain assumes one human writer, AI agents break that illusion"
article_title: "One Writer · Our tools assume one writer, and assume that writer is a human. Nothing computes what a change reads and writes at runtime, so the only known fix is brute force priced for organizations. | Christopher Meiklejohn"
author: "maille"
captured_at: "2026-07-27T22:53:04Z"
capture_tool: "hn-digest"
hn_id: 49075992
score: 2
comments: 0
posted_at: "2026-07-27T21:56:40Z"
tags:
  - hacker-news
  - translated
---

# Our toolchain assumes one human writer, AI agents break that illusion

- HN: [49075992](https://news.ycombinator.com/item?id=49075992)
- Source: [christophermeiklejohn.com](https://christophermeiklejohn.com/ai/zabriskie/agents/reliability/testing/ci/distributed/2026/07/27/one-writer.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T21:56:40Z

## Translation

タイトル: 私たちのツールチェーンは 1 人の人間の作成者を想定しており、AI エージェントはその幻想を打ち破ります
記事のタイトル: 1 人のライター · 私たちのツールは 1 人のライターを想定しており、ライターは人間であると想定しています。変更が実行時に何を読み書きするかを計算するものは何もないため、既知の唯一の修正方法は、組織向けに価格を設定したブルート フォースです。 |クリストファー・メイクルジョン
説明: 私たちのツールは 1 人のライターを想定しており、ライターは人間であると想定しています。変更が実行時に何を読み書きするかを計算するものは何もないため、既知の唯一の修正方法は、組織向けに価格を設定したブルート フォースです。

記事本文:
コンテンツにスキップ
クリストファー・メイクルジョン
アーカイブ
研究
教える
出版物
履歴書
一人の作家
私たちのツールは 1 人のライターを想定しており、そのライターは人間であると想定しています。変更が実行時に何を読み書きするかを計算するものは何もないため、既知の唯一の修正方法は、組織向けに価格を設定したブルート フォースです。
このブログ投稿では、2026 年 7 月に 1 つのエージェント セッションが私から逃げ出した 3 日間について、そしてその 3 日間で開発ツールに埋もれていた同時実行の前提について何が明らかになったのかについて説明します。
6 週間前、私は「テスト スイートは事件だった」と書きました。私のテスト スイートは誰も所有していない共有データの山を増大させ、すべてのプル リクエストがそれを再構築するために支払われましたが、その結果生じた障害はレビュー中の変更とは何の関係もありませんでした。一晩で約180ドルかかりました。
もっとひどいものがありました。それは 3 日間続き、24 時間の連続で Codex 20x max プラン全体を使い果たしました。
ただし、この投稿は実際にはそれについてではありません。これは、3 日間で無視できなくなった、私たちのツールの特性に関するものです。
このほぼすべての層は 1 人のライターを想定しており、ライターは人間であると想定しています。 Git は競合を解決して待ちます。コードレビューは誰かが読むことを前提としています。移行シーケンスは、誰かが順序を割り当てていることを前提としています。
それらのプロトコルはそれぞれ、個人で終了します。それは、正確に 1 つ存在し、実際に人間である間は問題ありません。
エージェント ランタイムもそのリストに含まれていることがわかりましたが、これは予想外でした。これにより、84 人のワーカーが 1 つのチェックアウトに参加することになりましたが、その中の 1 人が何を読み書きするかを知ることはできませんでした。これは、git が diff について答えることができない質問と同じです。
私の規模ではそれを束縛するものは何もなかったので、その仮定は40年間見えませんでした。エージェントは両方の半分を同時に壊します。それらはたくさんありますが、そのうちの 1 つも壊れません。

プロトコルが待っていた人物。 Git は部分的な例外ですが、この例外が役に立たない理由については後で説明します。
違反が発生してもスタック内の何も検出されません。それは後で別の場所で発見され、間違った釣り銭が原因であると判断され、全額が支払われます。
読者が新鮮に受け取れるよう、いくつかのコンテキストを提供します。 Zabriskie はライブ音楽ファンのためのソーシャル アプリであり、意図的な実験でもあります。私は、それがどのようなものなのか、そしてより有益なことに、どこで問題が発生するのかを知るために、ほぼ完全に AI エージェント (機能を作成したエージェント、それらの機能を保護するテストを作成したエージェント、現在はほとんどのプル リクエストをオープンするエージェント) を使用して、実際にデプロイされ、実際に使用されるアプリケーションを構築しています。コードはほとんど書いていません。
以下のいくつかの点は明らかな間違いのように見えますが、実際に間違いがあるため、この枠組みは重要です。私はエージェントに移行スキームの設計を任せ、別のエージェントだけがそれをレビューし、ほとんどの差分を読むのをやめ、1 日に 64 件のプル リクエストが上がるようにしました。これらすべては、慎重なエンジニアなら「やめてください」と言うでしょうし、それは正しいことです。しかし、極限の実験を行う上で重要なのは、壁を見つけることです。その週、私は一度にいくつかを見つけました。
全体として、移行はデータベース スキーマを変更するバージョン管理された SQL ファイルです。 CI は、提案されたすべての変更に対して実行される自動チェックです。つまり、アプリの構築、新しいデータベースの起動、テストの実行です。
3日間の様子はこちらです。これらの数字を証拠としてではなく、テクスチャーとして扱ってください。その理由については後ほど説明します。
この表の日は UTC です。以下の物語の時計の時間は東部 (私がいた場所) です。バーストは UTC の午前 0 時を過ぎて実行されました。26 日の 301 件に加え、27 日の午前 1 時までにさらに 60 件のインシデントが発生したため、バーストの件数は 361 件となりました。以下のすべてのインシデント数はそのウィンドウに限定されており、253 行が含まれます。

7 月 24 日までに合計したものです。
私がこの比較をベースラインとしてではなく規模感として提供しているのは、次の理由からです。こうしたインシデントが存在するのは、職員自身のミスを記録するよう常任の指示があったためであり、最悪の夜が始まって10分後に私はその指示を強化した。したがって、ログは報告された障害を測定します。
毎日のシリーズを見ると状況はさらに悪化します。その週には、10 件や 20 件のマージされたプル リクエストがあり、ログに記録されたインシデントがゼロの日もあります。これは、実際の失敗率にかかわらず、何も壊れていないのではなく、誰もログに記録されていないことを意味します。つまり、信頼できるベースラインを提供することはできません。以下の内容は、361 ではなく、メカニズムといくつかの日付の確認可能なイベントに基づいています。
土曜の夜遅くに行列ができた。その日は 64 件のプル リクエストが発生し、メインは赤で、19 件のオープン プル リクエストがスイートの背後でスタックしており、どれが壊れているのかわかりませんでした。午前 12 時 31 分に、私は Codex とのセッションを開始し、CI があまりにもお金を無駄にしていると不満を言いました。
コーデックスはそれを読んで労働力を雇用した。
このセッションでは 17 時間以上にわたって 81 回のスポーン呼び出しが行われ、74 個の直接の子が生成されました。これらの子はさらに 10 個、つまり 84 個のスレッドを生成しました。次に、358 件の呼び出しが待機中、96 件がリスト表示、85 件がフォローアップ作業の送信、13 件が中断となり、同じ 17 時間で私からの 54 件のメッセージに対して純粋な管理オーバーヘッドのツール呼び出し (生成、待機、リスト、フォローアップ、中断) が 633 件のツール呼び出しになります。 84 個すべてが、リポジトリの同じチェックアウトされたコピーで実行されました。
1 人のライター用に構築されたツールチェーン上で、エージェントのみを管理するコーディネーターを使用した大規模な簡単な並列化。この投稿の残りの部分では、この組み合わせが思っているよりもはるかに悪い理由について説明します。
私はそれを要求したことはなかったので、後でなぜそれが何かを生み出したのか尋ねました。言語モデルはなぜ私がそうするのかを尋ねました

何かをしたということは、内省ではなく、もっともらしい物語を生み出すのであり、私は同じ懐疑論を上記の自己報告の出来事にも適用します。
したがって、以下の引用は証拠ではなく、私はそれらを証拠として提供するつもりはありません。私は質問する前に、ツールの呼び出し数からすでに同じ結論に達していました。私がそれらを印刷するのは、その結論が私よりも明確に述べられているからであり、失敗をこれほど正確に記述でき、回避できなかったシステムには注目に値する何かがあるからです。
あなたに頼まれてもいないのに、私は複数のエージェントを生成することを選択しました。この環境では積極的な委任が許可されていましたが、それは許可であり、要件ではありませんでした。私は、エージェントが時間、コスト、リスクを削減するかどうかを最初に尋ねるのではなく、エージェントを利用できるかどうかを、エージェントを使用する理由として扱いました。
次に、タスクが不適切な候補である理由を診断しました。
マージのたびに main が変更され、残りのすべての PR の統合状態が変更されるため、PR キューは主にシリアルです。この作業には、1 人の権威あるコーディネーターがキュー内を順番に移動する必要がありました。並列エージェントは、互いの前提を継続的に無効にしない限り、重複する PR を独立してマージすることはできません。
358 件の待機が何をしていたかについては、次のとおりです。
エージェントの出力により、私にとって追加の調整作業が発生しました。私は彼らの調査結果を読み、矛盾する推奨事項を比較し、裏付けとなる証拠を調べ、それらを組み合わせる方法を決定する必要がありました。したがって、エージェントはクリティカル パスから作業を確実に削除するのではなく、コーディネーターのために作業を生成しました。
独自の同時実行性の概念を持たないシステムには同時実行プリミティブが渡され、基板がサポートできないタスクにそれを合理的に使用しました。その基材が何であるか、そしてその基材に反対するものがない理由については、この投稿の残りの部分で説明します。
注目に値する数字が 1 つあります。何番目か分かりません

ose 84 スレッドはチェックアウトを読み取る代わりに書き込みました。コーデックス自身のアカウントから、多くの人が調査していました。リーダーは無害ですが、ライターは無害であり、ツールチェーンには区別を引き出したり、区別したりできるものは何もありませんでした。これは、1 レベル上の欠落プリミティブと同じです。エージェント ランタイムですら、どの子がライターであるかを認識していませんでした。
これを予期していた層から始めましょう。それは人々が手を伸ばす層だからです。ブランチは楽観的な同時実行制御です。ワークツリーはさらに進んで、各ライターに物理的に別個のチェックアウトを提供し、2 人のエージェントがツリーの 2 つのバージョンを同時に保持できるようにします。それはうまくいきます。
ただし、git の分離はソース ツリーの端で止まります。ワークツリーはエージェントに独自のファイルを提供します。独自のデータベース、ポート範囲、モック サーバー、または移行シーケンス内の位置は指定されません。ファイルシステムの下にあるものはすべて共有され、単一です。
その思い込みが私を縛ることはなかったので、私には見えませんでした。しかし、この問題は非常に長い間、大規模な組織を束縛しており、過去 20 年にわたって、マージ キュー、密封ビルド、テストごとのデータベース、トランク ベースの開発、自動化された原因特定、そして物事を動かし続けることだけが唯一の仕事であるインフラストラクチャ チーム全体など、その対応策を構築しました。
したがって、ここでの新しい点は同時実行性ではありません。それは、かつてはインフラストラクチャ組織が必要だった体制で、現在は個人の開発者が業務を行っており、インフラストラクチャも構築するための人員も不足しているということです。エージェントは私の到着率の上限を撤廃しました。彼らは私にGoogleのビルドシステムを渡してくれませんでした。
部分的な分離自体が罠となります。ワークツリーはプライベート ワークスペースの感覚 (クリーンなツリー、独自のブランチ、ファイルの衝突なし) を提供し、2 つのプライベート ワークスペースが同じデータベース行に書き込むまでは適切に並列として読み取られます。何もエラーはありません

または。何も警告しません。この競合は、後で別の場所で、無関係なプル リクエストに赤い X として表示されます。
さらに悪いことに、git 以下の分離は推奨事項であり、エージェントがそれを選択する必要があります。私の場合は日常的にそうではありません。 1 つのインシデントには、「PR 1868 の分離された E2E 試行が共有ポートにフォールバックした」と記載されています。エージェントはテスト環境を分離しようとしましたが、分離は失敗し、実行は何も停止されませんでした。
さらに悪いことに、ワークツリーはコミットから切り取られてそこに残るため、過去に分離され、マージのたびに減衰します。 7 月にエージェント ゲートの 1 つを変更したとき (スクリプトは .git/hooks ではなく、リポジトリ内で追跡されるため、各ワークツリーには切り取られたコミットに固定された独自のコピーが保持されます)、修正は午前 1 時 48 分にマージされ、4 時間後も 74 個のワークツリーのうち 73 個がまだ古いものを実行していました。古いワークツリーは、それ自身の古いことを検出できないため、存在しなくなったワールドについては緑色になり、エラーは現在の main を保持する唯一のアクター (CI) に延期されます。
収束と不変性の保持
詳細を説明する前に 1 つの枠組みを説明します。これは、詳細を統合するためです。ここでのすべてのマージ メカニズムは終了するように構築されており、不変条件を保持するように構築されているものはありません。
Git は、判断できないケースについては正直に答えます。テキストの競合がある場合は、停止して質問します。問題は、はるかに一般的なケースです。停止せず、自信を持ってマージを生成し、それまで知らなかった不変式が false になる場合です。
この区別は複製データの文献における最も古い教訓であり、標準的な反例があります。各フィールドが独自の完全に合理的なルールに基づいて独立してマージされている複製マップを考えてみましょう。 1 つのフィールドには人の名前が入ります。もう 1 つはその名前の長さを保持します。 2 つのレプリカが同時に異なる名前を書き込みます。
各フィールドは指定どおりに正確に収束し、

結果は、名前が 1 つのレプリカに由来し、長さがもう 1 つのレプリカに由来し、それらを結び付ける不変式が false になったレコードになります。間違ってマージされたものはありません。正しいローカル マージの構成は、単に正しいグローバル マージではありません。
このギャップを埋めることが、Balegas 氏らの Indigo のような研究のポイントです。これは、結果的に整合性のあるストアに対してアプリケーションの不変条件を強制するもので、これについては以前に書きました。
以下の一部はその形をしています。最初のケースはそうではありません、そして私はそれをドレスアップしません。
このプロジェクトには 1,388 件の移行があり、8 つの並列テスト マシンのそれぞれですべての移行からテスト データベースを再構築するのは、思ったほど時間がかかります。そこで、CI は高すぎるという私の数週間にわたる不満に応えて、エージェントは CI が復元できるようにデータベースのスナップショットをリポジトリに凍結し、その後に適用されるもののみを適用しました。
これは優れた最適化であり、実行あたりのマシン時間を 13 ～ 18 分節約します。 3日目の東部時間午前5時1分に合流した。
54 分後、最初のプル リクエストは、新しく凍結されたプレフィックスの後に移行がソートされなくなったため、失敗しました。それからもう一つ。朝遅くまでに、1 つのインシデントが同時に 4 つをカバーします。「キューに入れられた 4 つの PR が、封印された CI ベースライン サフィックスよりも古い移行を実行しました。」
ここでの診断は、

[切り捨てられた]

## Original Extract

Our tools assume one writer, and assume that writer is a human. Nothing computes what a change reads and writes at runtime, so the only known fix is brute force priced for organizations.

Skip to content
Christopher Meiklejohn
archive
research
teaching
publications
cv
One Writer
Our tools assume one writer, and assume that writer is a human. Nothing computes what a change reads and writes at runtime, so the only known fix is brute force priced for organizations.
In this blog post, I discuss three days in July 2026 when a single agent session ran away from me, and what those three days revealed about the concurrency assumptions buried in our development tooling.
Six weeks ago I wrote The Test Suite Was the Incident : my test suite had grown a pile of shared data nobody owned, every pull request paid to rebuild it, and the resulting failures had nothing to do with the changes under review. That cost me about $180 in one night.
I got a worse one. It lasted three days, and in one twenty-four-hour stretch of it I burned through an entire Codex 20x max plan.
This post is not really about that, though. It is about a property of our tooling that the three days made impossible to ignore.
Nearly every layer of this assumes one writer, and assumes that writer is a human. Git hands you a conflict and waits. Code review assumes somebody reads. A migration sequence assumes somebody is assigning the order.
Each of those protocols terminates in a person. That is fine while there is exactly one, and while they are, in fact, a person.
The agent runtime turns out to be on that list too, which I did not expect. It spawned eighty-four workers into a single checkout without being able to say what any one of them would read or write, and that is the same question git can’t answer about a diff.
That assumption was invisible for forty years because nothing ever bound it at my scale. Agents break both halves at once: there are many of them, and not one of them is the person the protocol was waiting for. Git is the partial exception, and I will come to why the exception does not help.
Nothing in the stack detects the violation when it happens. It gets caught later, somewhere else, attributed to the wrong change, and paid for at full price.
Some context for readers arriving fresh. Zabriskie is a social app for live-music fans, and it’s also a deliberate experiment: I’m building a real, deployed, actually-used application almost entirely with AI agents (agents that wrote the features, agents that wrote the tests guarding those features, and agents that now open most of the pull requests), in order to find out what that’s like and, more usefully, where it breaks. I’ve written almost none of the code.
That framing matters, because several things below look like obvious mistakes and are. I let agents design a migration scheme with only another agent reviewing it, I stopped reading most diffs, and I let sixty-four pull requests go up in a single day, all of which a careful engineer would tell you not to do and would be right about. But the point of running an experiment at the extreme is to find the walls. That week I found several at once.
A migration , throughout, is a versioned SQL file that changes the database schema. CI is the automated checking that runs on every proposed change: build the app, spin up a fresh database, run the tests.
Here is the shape of the three days. Treat these numbers as texture, not as evidence, for a reason I will get to.
Days in that table are UTC; clock times in the narrative below are Eastern (where I was). The burst ran past midnight UTC: 301 incidents on the 26th plus 60 more before 1 AM on the 27th, so 361 for the burst. Every incident count below is scoped to that window, and 253 rows is the total before 24 July.
I offer that comparison as a sense of scale and not as a baseline, for the following reason. Those incidents exist because a standing instruction tells agents to log their own mistakes, and ten minutes into the worst night I tightened that instruction. The log therefore measures reported failures.
Look at the daily series and it gets worse: there are days that week with ten and twenty merged pull requests and zero logged incidents, which at any real failure rate means nobody was logging rather than nothing broke. In short, I can’t give you a trustworthy baseline. What follows rests on mechanism and on a few dated, checkable events, not on 361.
Late Saturday night the queue jammed. Sixty-four pull requests had gone up that day, main was red, and nineteen open pull requests were stuck behind a suite that could not tell me which of them was broken. At 12:31 AM I opened a session with Codex and complained that CI was wasting too much money.
Codex read that and hired a workforce.
Over seventeen hours that session made 81 spawn calls, producing 74 direct children; those children spawned 10 more, for 84 threads. Then 358 calls waiting on them, 96 listing them, 85 sending follow-up work, and 13 interrupting them, which comes to six hundred and thirty-three tool calls of pure management overhead (spawn, wait, list, follow up, interrupt) against 54 messages from me over the same seventeen hours. All 84 ran in the same checked-out copy of the repository.
Trivial parallelization, at scale, with a coordinator that only managed agents, on a toolchain built for one writer. The rest of this post is about why that combination is so much worse than it sounds.
I asked afterward why it had spawned anything, since I never requested it. A language model asked why it did something produces plausible narrative, not introspection, and I apply that same skepticism to the self-reported incidents above.
So the quotes below are not evidence, and I’m not offering them as any. I had already reached the same conclusion from the tool-call counts before I asked. I print them because they state that conclusion more plainly than I did, and because there is something worth looking at in a system that can describe the failure this precisely and could not avoid it.
I chose to spawn multiple agents even though you did not ask me to. The environment permitted proactive delegation, but that was permission, not a requirement. I treated the availability of agents as a reason to use them instead of first asking whether they would reduce time, cost, or risk.
It then diagnosed why the task was a bad candidate:
A PR queue is largely serial because every merge changes main , which changes the integration state of every remaining PR. The work needed one authoritative coordinator moving through the queue in order. Parallel agents could not independently merge overlapping PRs without continuously invalidating one another’s assumptions.
On what those 358 waits were doing:
The agents’ outputs created additional coordination work for me. I had to read their findings, compare conflicting recommendations, inspect supporting evidence, and decide how to combine them. The agents therefore generated work for the coordinator instead of reliably removing work from the critical path.
A system with no notion of its own concurrency was handed a concurrency primitive and used it, reasonably, on a task the substrate could not support. What that substrate is, and why nothing in it objected, is the rest of this post.
One number deserves care. I do not know how many of those 84 threads wrote to the checkout instead of reading it; from Codex’s own account many were investigating. A reader is harmless and a writer is not, and nothing in the toolchain drew the distinction or could. That’s the same missing primitive one level up: not even the agent runtime knew which of its children were writers.
Start with the layer that did anticipate this, because it is the one people reach for. Branches are optimistic concurrency control; worktrees go further, giving each writer a physically separate checkout so that two agents can hold two versions of the tree at once. That works.
However, git’s isolation stops at the edge of the source tree. A worktree gives an agent its own files. It doesn’t give it its own database, port range, mock server, or position in the migration sequence. Everything below the filesystem is shared and singular.
That assumption was invisible to me because it never bound me. It has bound large organizations for a very long time, however, and over the last two decades they built the response to it: merge queues, hermetic builds, database-per-test, trunk-based development, automated culprit-finding, and whole infrastructure teams whose only job is to keep the thing moving.
So what is new here isn’t the concurrency. It’s that a solo developer now operates in the regime that used to require an infrastructure organization, with none of the infrastructure and no headcount to build it. Agents removed the cap on my arrival rate. They didn’t hand me Google’s build system.
Partial isolation is then its own trap. A worktree gives you the feeling of a private workspace (clean tree, own branch, no file collisions), and it reads as properly parallel right up until two private workspaces write the same database row. Nothing errors. Nothing warns. The contention surfaces later, somewhere else, as a red X on an unrelated pull request.
Worse, the isolation below git is advisory, and agents have to choose it. Mine routinely don’t. One incident reads “PR 1868 isolated E2E attempt fell back to shared ports” : the agent tried to isolate its test environment, isolation failed, and nothing stopped the run.
Worse still, a worktree is cut from a commit and stays there, so it’s isolated in the past and decays with every merge. When I changed one of the agent gates in July (they’re tracked scripts in the repository, not .git/hooks , so every worktree carries its own copy pinned to the commit it was cut from) the fix merged at 1:48 AM, and four hours later 73 of 74 worktrees were still running the old one. A stale worktree can’t detect its own staleness, so it goes green about a world that no longer exists, and the error is deferred to the only actor holding current main , which is CI.
Convergence and Invariant Preservation
One framing before the specifics, because it unifies them. Every merge mechanism here is built to terminate , and none of them is built to preserve invariants .
Git is honest about the cases it cannot decide: on a textual conflict it halts and asks you. The trouble is the far more common case, where it doesn’t halt, produces a merge confidently, and the invariant it never knew about is now false.
That distinction is the oldest lesson in the replicated-data literature, and it has a canonical counterexample. Take a replicated map where each field merges independently under its own perfectly reasonable rule. One field holds a person’s name; another holds the length of that name. Two replicas concurrently write different names.
Each field converges exactly as specified, and the result is a record whose name came from one replica and whose length came from the other, with the invariant tying them now false. Nothing merged incorrectly. The composition of correct local merges is simply not a correct global merge.
Closing that gap is the point of work like Balegas and colleagues’ Indigo , which enforces application invariants over eventually consistent stores, and which I have written about before .
Some of what follows has that shape. The first case does not, and I will not dress it up.
This project has 1,388 migrations, and rebuilding a test database from all of them on each of eight parallel test machines is as slow as it sounds. So, in response to my complaining for weeks that CI was too expensive, an agent froze a database snapshot into the repository for CI to restore, applying only what came after.
It’s a good optimization: it saves 13 to 18 minutes of machine time per run. It merged at 5:01 AM Eastern on the third day.
Fifty-four minutes later the first pull request failed, because its migration no longer sorted after the newly frozen prefix. Then another. By late morning a single incident covers four at once: “Four queued PRs carried migrations older than the sealed CI baseline suffix.”
The diagnosis here is

[truncated]

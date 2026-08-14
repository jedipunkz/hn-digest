---
source: "https://mozaik.jigjoy.ai/blog/the-three-ways-people-build-ai-agent-systems"
hn_url: "https://news.ycombinator.com/item?id=49298338"
title: "The three ways people build AI agent systems"
article_title: "The Three Ways People Build AI Agent Systems: From Fixed Workflows to Concurrent Collectives"
author: "qikouki"
captured_at: "2026-08-14T14:09:50Z"
capture_tool: "hn-digest"
hn_id: 49298338
score: 1
comments: 0
posted_at: "2026-08-14T13:22:04Z"
tags:
  - hacker-news
  - translated
---

# The three ways people build AI agent systems

- HN: [49298338](https://news.ycombinator.com/item?id=49298338)
- Source: [mozaik.jigjoy.ai](https://mozaik.jigjoy.ai/blog/the-three-ways-people-build-ai-agent-systems)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T13:22:04Z

## Translation

タイトル: AI エージェント システムを構築する 3 つの方法
記事のタイトル: AI エージェント システムを構築する 3 つの方法: 固定ワークフローから並行コレクティブまで
説明: 本番環境のすべてのエージェント システムは、チェーン、グラフ、またはイベント駆動型集合体であり、ほとんどのチームは選択しませんでした。 3 つのアーキテクチャの分類、それぞれの天井、そして最初のアーキテクチャでシステムの半分がまだ生きていたため、考えている途中で亡くなったプランナーの話。

記事本文:
AI エージェント システムを構築する 3 つの方法: 固定ワークフローから同時コレクティブまで moza i k docs ↗ ブログ github ↗ コミュニティ ↗ GitHub でスターを付ける ↗ AI エージェント システムを構築する 3 つの方法: 固定ワークフローから同時コレクティブまで
私が実稼働環境で見たすべての AI エージェント システム (当社のものも含む) は、3 つの方法のいずれかで構築されています。ほとんどのチームは意識的に選択しません。彼らは、インストールしたフレームワークから最初の方法を継承し、その限界を発見し、2 番目の方法に卒業し、2 番目の方法で失敗した場合にのみ 3 番目の方法に到達します。
先週、私たちのシステムの一部がまだ最初の方法で密かに生きていたため、私たち自身のエージェントの1人が考えている間に死亡しました。この投稿は分類学とその話です。
方法 1: エージェントの滝
誰もが最初に構築するシステムはチェーンです。エージェント A は分析し、推論するエージェント B に引き継ぎ、書き込みを行うエージェント C に引き継ぎます。終わり。すべてのフレームワーク チュートリアルでは、この形状が描画するのが最も簡単であるため、この形状を教えています。
エージェント B の理由 — 待機中 エージェント C の書き込み — 完了待機中 1 つのエージェントが動作しています。他の人は皆順番を待っています。依存関係のチェーンがアーキテクチャになります。問題は、チェーンが単純であるということではありません。問題は、依存関係のチェーンがアーキテクチャになることです。新しい機能が追加されるたびに、ハンドオフが追加され、待機が追加され、障害ポイントが追加されます。レビュアーが欲しいですか?これを B と C の間に挿入すると、B の下流のすべてがその上で待機します。研究者が欲しいですか?別のリンク、別の連載ポイント。
すべての請求書には次の 4 つのコストが表示されます。
実行時間。すべてのハンドオフが同期の障壁になります。システムはリンクの合計と同じくらい遅いです。
トークンの無駄。各エージェントがコールドで開始されるため、コンテキストはホップごとに再フェッチされ、再説明されます。
剛性。エージェントを追加するということは、

順序。組織図はコードに組み込まれています。
脆弱な回復。チェーンの途中で予期せぬことが起こった場合、「スローしてリスタートする」以外にその知識を入れる場所はありません。
私たちはこれをエージェントティック ウォーターフォールと呼んでいますが、その名の通り、小規模で既知の直線的な作業に間違いはありません。作品が直線的でなくなった瞬間、それが作品が面白くなる瞬間です。
明らかな解決策はファンアウトすることです。バックエンド エージェント、フロントエンド エージェント、およびテスト エージェントを同時に実行します。その結果を結合します。続く。これは DAG ステージです。ワークフロー グラフ、ファンアウト/ファンイン、追加のステップを含む Promise.all です。
本当に助かります。これは、今日、ほとんどの成熟したチームが行き詰まっている場所でもあります。その理由は、私たちが明確にするのに長い時間がかかった微妙な特性のせいです。
枝は平行です。システムはそうではありません。
すべてを待ちます 最も遅いペースを設定します 次のステップはまだ開始できません 分岐は並列です。システムはそうではありません。ブランチ内にあるものは、結合前の兄弟に影響を与えることはできません。すべての結合はすべて待機の障壁になります。遅い分岐がペースを決めます。高速ブランチはマージ時にアイドル状態になります。また、唯一の通信チャネルは最後の参加ポイントであるため、ブランチ内で発生することは兄弟に影響を与えることはできません。最初の 10 秒で致命的な問題を発見したテスト エージェントは、バックエンド エージェントにトークンの無駄遣いをやめるよう指示する方法がありません。後で、お金がすでに使われてしまったときに、バックエンド エージェントは障壁で会うことになります。
グラフも作業を開始する前に描く必要があります。作品を知っていれば大丈夫です。まさに、実行中に発見された作業に対しては失敗します。そして、発見された作業こそがエージェントの目的です。
方法 3: イベント駆動型集合体
3 番目の方法では、実行パスのモデル化を停止し、エージェントが動作する環境のモデル化を開始します。エージェントも参加する

共有ランタイム内の受信者。彼らはセマンティック イベント (「ストーリー S1 がマージされた」、「検証に失敗した」、「誰も計画していなかった作業を発見した」) を発行し、関心のあるものをサブスクライブします。誰も誰にも手を渡さない。エージェントが考えている間も進行は続行され、システムは待機が現実的な場合にのみ待機します。別のストーリーのファイルを必要とするストーリーは、純粋にマージを待ちます。他のすべては流れます。
セマンティック イベント ストリーム 全員が働いています。システムは実際に待機する必要がある場合にのみ待機します。調整ロジックは、ワークフローにハードコードされた分岐ではなくなり、状況に関するルールになります。つまり、ユーザー メッセージが到着し、容量が利用可能になったときに推論を実行します。予算、失敗、予期せぬ事態は、グラフに散りばめられた特別なケースではなく、再利用可能なシステム ルールになります。
多くのプログラムを同時に実行するオペレーティング システムのようなものです。これがメンタル モデルです。ステップ図ではありません。仕事が起こる場所。
考え半ばで亡くなったプランナーの物語
これが、私が他の週ではなく今これを書いている理由です。
私たちのエージェント baro は、実際のリポジトリに対してコーディング エージェントの集合体を実行します。エグゼキュータは数か月間、方法 3 で生活しています。彼らは同時に実行し、イベントを交換し、作業中にお互いのマージについて学習します。しかし、プランナー (目標をストーリーに分解するエージェント) は、アーキテクチャ的には依然としてウォーターフォールの住民であり、一方通行のパイプの終端にある別のプロセスであり、計画の断片を下流に公開し、何も返されません。
2 つのことが続き、両方を測定しました。
まず、脆弱性の回収コストです。プランナーはランタイムの範囲外にあったため、ブラック ボックスを監視できる唯一の方法である実時間タイムアウトをホストが監視しました。生産的な計画セッションが開始されてから 8 分後、フラグメントが公開された直後にクリアされます。

生きている証拠の可能性はありません - 番犬がそれを殺しました。同じ理由でタイムアウトがすでに一度設定されており、コード内のコメントでそれが認められていました。これが、内部から見た「脆弱な回復」の様子です。コンポーネントが生きていることが分からないとき、人は推測することになり、推測するたびに、最終的には罪のない人が殺されてしまいます。
2つ目は通信費です。プランナーのパブリッシュ呼び出しは、ランタイムの実際の許可ではなく、独自のプロセスによって計算されたレシート (丁寧なローカル エコー) を返しました。プランナーが想像していた依存関係エッジをホストがプルーニングすると、その事実は誰も読まないログ ファイルに記録されました。ストーリーは実行され、統合され、失敗しました。プランナーは、盲目的に、もはや存在しないリポジトリの状態に対して計画を立てました。
そこで私たちはそれをコレクティブに移しました。プランナーは、計画するすべてのものと同じイベント バスの参加者になります。フラグメントを公開するときに取得するレシートは、ランタイムの本当の答えです。つまり、グラフのバージョン、認められたストーリー、正確にどの依存関係のエッジが刈り取られたのか、そしてその理由がわかります。次のフラグメントを計画している間、実行ニュースがライブでストリーミングされます。この記事はマージされ、あの記事は 2 回失敗し、この記事はブロックされました。そして、その監視は「どのくらいの時間実行されているか」から「どのくらい沈黙していましたか」に変更されました。出力によって時計がリセットされるため、目に見えて動作しているエージェントは任意のストップウォッチによって強制終了されることはありませんが、本当にハングしたエージェントは依然としてすぐに終了します。
プランナーは賢くなったわけではない。その状況はさらに賢くなった。これが、方法 3 のテーマ全体を 1 つの文でまとめたものです。
3 番目の方法には最初の 2 つが含まれます
この分類法にない点が 1 つあります。それは、どちらかの側を選択する 3 つの競合するツールです。
チェーンは退化した集合体であり、すべてのエージェントが「前任者が終了した」という 1 つのイベントだけをサブスクライブし、他には何もサブスクライブしません。グラップ

h は、結合がたまたま実際のものであり、事前にわかっている集合体です。イベント駆動型ランタイムは、両方を文字通りルールとして表現します。S1 がマージされると、開始 S2 はウォーターフォールになります。 3 つのブランチすべてがレポートする場合、合成はファンインとなります。方法 3 ではシーケンスを禁止するものは何もありません。これにより、シーケンスの義務が取り除かれます。つまり、方法 1 と 2 に組み込まれている、一部のものはブロックする必要があるため、すべてがブロックされる必要があるという前提が取り除かれます。
私たち自身のランニングでもこれを毎日目にしています。単一の実行内では、基礎ストーリーとその依存関係が純粋なチェーンとして実行され、6 つの独立したモジュールの移行がグラフとして実行され、発見を相互に説明するエージェントが完全な集合体になります。作業の各部分が必要な結合を正確に取得しているため、1 つのランタイム、1 つのルール セット、3 つのシェイプが共存します。
したがって、現在 1 つ目または 2 つ目の方法を実行している場合、移行は書き換えではありません。ワークフローはすでに 3 番目のモデルの有効なプログラムになっており、暗黙的なプログラムにすぎません。それを明示することで、作業にブロックが必要ないと判明した場合に、一度に 1 つずつバリアを停止するという選択肢が得られます。
アドバイスの正直なバージョンは、「常に方法 3 を構築する」というものではありません。それは次のとおりです。
作品が既知で、小さく、直線的な場合は、チェーンでも問題ありません。発送してください。
作業が既知で並列化可能であれば、グラフによって実際の速度が得られます。それを受け入れて、接合部があなたの天井であることを受け入れてください。
作業中に作業が発見された場合、つまりエージェントが他のエージェントの行動を変えるはずのことを飛行中に学習する場合、チェーンとグラフは両方とも同じ場所で失敗します。現実が図から乖離した瞬間です。このような場合に必要となるのは、ワークフローではなく環境です。
そして、これが私たちが弁護する準備ができている主張です。実際の並行性、つまりエージェントが発見し、反応し、通信するランタイムがなければ、真にインテリジェントなシステムに到達することはできません。

事前に定義されたステップの間ではなく、作業が行われている間に統合して適応します。インテリジェンスは単一のモデル呼び出しの特性ではありません。それは、飛行中に考えが変わる可能性があるシステムの特性です。チェーンはその考えを変えることはできません。グラフは障壁でのみ考えを変えることができます。私たちが知っているすべてのインテリジェント システム (チーム、市場、組織) は同時に実行され、行動中に物事に気づき、世界を止めることなく再計画を立てます。エージェント システムも同様です。
私たちがオープンソースの TypeScript ランタイムである Mozaik を構築したのは、その 3 番目のケースに遭遇し続けたためです。上記のプランナーのストーリーは、移行を行うための最新のコンポーネントにすぎません。移行のパターンはそれ以前のすべてのコンポーネントと同じでした。つまり、一方通行のパイプを見つけ、ストップウォッチを見つけ、両方を参加に置き換えました。
依存関係チェーンはアーキテクチャではありません。環境は。
私の情熱は、私たちが構築した美しいものについてのストーリーを世界に伝えることです。
私たちがすでに持っているツールとテクノロジーを使えば、今日のほとんどのプロジェクトよりもはるかに価値のあるシステムを構築できます。私たちは、使うのが楽しく、作業するのが楽しいソフトウェアを書くことができます。ソフトウェアは、成長するにつれて私たちを囲い込むのではなく、新しい機会を生み出し、所有者に価値を与え続けます。 — エリック・エヴァンス
自己組織化エージェントを構築する方法を学びたい開発者向け。
Mozaik を使用すると、エージェントは独立して作業し、ワークフロー全体をブロックすることなく調整できます。
© 2026 JigJoy · MIT · JigJoy の一部 ↗
私たちはハッカソンを企画しています - Mozaik フレームワークを使用して競争します。

## Original Extract

Every agent system in production is a chain, a graph, or an event-driven collective — and most teams never chose. A taxonomy of the three architectures, the ceiling of each, and the story of a planner that died mid-thought because half our system was still living in the first one.

The Three Ways People Build AI Agent Systems: From Fixed Workflows to Concurrent Collectives moza i k docs ↗ blog github ↗ community ↗ Star on GitHub ↗ The Three Ways People Build AI Agent Systems: From Fixed Workflows to Concurrent Collectives
Every AI agent system I have seen in production — ours included — is built one of three ways. Most teams don’t choose consciously. They inherit the first way from the framework they installed, discover its ceiling, graduate to the second, and only hit the third when the second one fails them in a way they can name.
Last week one of our own agents died mid-thought because part of our system was still secretly living in the first way. This post is the taxonomy, and that story.
Way one: the Agentic Waterfall
The first system everyone builds is a chain. Agent A analyzes, hands off to Agent B, which reasons, hands off to Agent C, which writes. Done. Every framework tutorial teaches this shape because it is the easiest thing to draw.
Agent B reason — waiting Agent C write — waiting done One agent works. Everyone else waits their turn. The dependency chain becomes the architecture. The problem is not that chains are simple. The problem is that the dependency chain becomes the architecture. Every new capability adds another handoff, another wait, another failure point. Want a reviewer? Insert it between B and C, and now everything downstream of B waits on it. Want a researcher? Another link, another serialization point.
Four costs show up on every invoice:
Execution time. Every handoff is a synchronization barrier. The system is as slow as the sum of its links.
Token waste. Context gets re-fetched and re-explained at every hop, because each agent starts cold.
Rigidity. Adding an agent means redesigning the sequence. The org chart is welded into the code.
Fragile recovery. When something unexpected happens mid-chain, there is no place to put that knowledge except “throw and restart.”
We call this the Agentic Waterfall, and like its namesake it isn’t wrong for small, known, linear work. It is wrong the moment the work stops being linear — which is the moment the work becomes interesting.
The obvious fix is to fan out. Run the backend agent, the frontend agent, and the test agent at the same time; join their results; continue. This is the DAG stage — workflow graphs, fan-out/fan-in, Promise.all with extra steps.
It genuinely helps. It is also where most mature teams are stuck today, because of a subtle property that took us a long time to articulate:
The branches are parallel. The system is not.
wait for all the slowest sets the pace Next step cannot start yet The branches are parallel. The system is not. Nothing inside a branch can influence a sibling before the join. Every join is a wait-for-all barrier. The slow branch sets the pace; the fast branches sit idle at the merge; and nothing that happens inside a branch can influence a sibling, because the only communication channel is the join point at the end. A test agent that spots a fatal problem in the first ten seconds has no way to tell the backend agent to stop wasting tokens — they will meet at the barrier, later, when the money is already spent.
The graph also has to be drawn before the work starts. That is fine when you know the work. It fails precisely on the work that gets discovered while doing — and discovered work is what agents are for.
Way three: the Event-Driven Collective
The third way stops modeling execution paths and starts modeling the environment where agents operate . Agents become participants in a shared runtime. They emit semantic events — “story S1 merged”, “verification failed”, “I discovered work nobody planned” — and they subscribe to what they care about. Nobody hands off to anybody. Progress continues while agents think, and the system waits only where waiting is real : a story that needs another story’s files genuinely waits for the merge; everything else flows.
SEMANTIC EVENT STREAM Everyone is working. The system waits only where waiting is real. Coordination logic stops being branches hardcoded into a workflow and becomes rules about situations: when a user message arrives and capacity is available, then run inference. Budgets, failures, and surprises become reusable system rules instead of special cases sprinkled through a graph.
Like an operating system running many programs at once — that is the mental model. Not a diagram of steps; a place where work happens.
The story of a planner that died mid-thought
Here is why I am writing this now instead of any other week.
Our agent, baro , runs a collective of coding agents against real repositories. The executors have lived in way three for months: they run concurrently, exchange events, learn about each other’s merges while working. But the planner — the agent that decomposes the goal into stories — was still architecturally a waterfall citizen: a separate process at the end of a one-way pipe, publishing plan fragments downstream, hearing nothing back.
Two things followed, and we measured both.
First, the fragile-recovery cost. Because the planner was outside the runtime, the host supervised it the only way you can supervise a black box: a wall-clock timeout. Eight minutes into a productive planning session — moments after it had published a fragment , the clearest possible proof of life — the watchdog killed it. The timeout had already been raised once, for the same reason, and the comment in the code admitted it. That is what “fragile recovery” looks like from the inside: when a component can’t tell you it is alive, you end up guessing, and every guess eventually kills someone innocent.
Second, the communication cost. The planner’s publish call returned a receipt computed by its own process — a polite local echo, not the runtime’s actual admission. When the host pruned a dependency edge the planner had imagined, that fact went to a log file nobody reads. Stories ran, merged, failed; the planner planned on, blind, against a repository state that no longer existed.
So we moved it into the collective. The planner is now a participant on the same event bus as everything it plans for. When it publishes a fragment, the receipt it gets is the runtime’s real answer — graph version, admitted stories, and exactly which dependency edges were pruned and why. While it plans the next fragment, execution news streams to it live: this story merged, that one failed twice, this one is blocked. And its supervision changed from “how long have you been running” to “how long have you been silent ” — output resets the clock, so an agent that is visibly working can never be killed by an arbitrary stopwatch, while a genuinely hung one still dies fast.
The planner did not get smarter. Its situation got smarter. That is the whole thesis of way three in one sentence.
The third way contains the first two
One thing this taxonomy is not : three competing tools where you pick a side.
A chain is a degenerate collective — one where every agent subscribes to exactly one event, “my predecessor finished,” and nothing else. A graph is a collective where the joins happen to be real and known in advance. An event-driven runtime expresses both, literally, as rules: when S1 merges, start S2 is a waterfall; when all three branches report, synthesize is a fan-in. Nothing about way three forbids sequence. It removes the obligation to sequence — the assumption, baked into ways one and two, that everything must block because some things must.
We see this daily in our own runs. Inside a single execution, a foundation story and its dependent run as a pure chain, six independent module migrations run as a graph, and the agents narrating discoveries to each other are the full collective — one runtime, one set of rules, three shapes coexisting because each part of the work got exactly the coupling it needed.
So if you are on way one or way two today, the migration is not a rewrite. Your workflow is already a valid program in the third model — just an implicit one. Making it explicit is what buys you the option to stop blocking, one barrier at a time, wherever the work turns out not to need it.
The honest version of the advice is not “always build way three.” It is:
If the work is known, small, and linear — a chain is fine. Ship it.
If the work is known and parallelizable — a graph buys real speed. Take it, and accept that the joins are your ceiling.
If the work is discovered while doing — if agents will learn things mid-flight that should change what other agents do — then chains and graphs will both fail you at the same place: the moment reality diverges from the diagram. That is when you need an environment, not a workflow.
And here is the claim we are prepared to defend: we will never reach truly intelligent systems without real concurrency — a runtime where agents discover, react, communicate, and adapt while the work is happening , not between predefined steps. Intelligence is not a property of a single model call; it is a property of a system that can change its mind mid-flight. A chain cannot change its mind. A graph can only change its mind at a barrier. Every intelligent system we know of — a team, a market, an organism — runs concurrently, notices things while acting, and re-plans without stopping the world. Agent systems will be no different.
We built Mozaik , our open-source TypeScript runtime, because we kept hitting that third case. The planner story above is just the latest component to make the move — and the pattern of the migration was the same as every one before it: find the one-way pipe, find the stopwatch, replace both with participation.
The dependency chain is not the architecture. The environment is.
My passion is to tell the world the stories about the beautiful stuff we have built.
With tools and technology we already have, we can build much more valuable systems than most projects today. We can write software that is a pleasure to use and a pleasure to work on; software that doesn't box us in as it grows, but creates new opportunities and continues to add value for its owners. — Eric Evans
For developers who want to learn how to build self-organizing agents.
Mozaik enables agents to work independently and coordinate without blocking entire workflows.
© 2026 JigJoy · MIT · part of JigJoy ↗
We're organizing a hackathon — compete using the Mozaik framework.

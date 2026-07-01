---
source: "https://onewill.ai/blog/2026/stealing-50-years-of-database-ideas-for-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48748977"
title: "Stealing 50 Years of Database Ideas for AI Agents"
article_title: "Stealing 50 Years of Database Ideas for AI Agents | OneWill"
author: "lmwnshn"
captured_at: "2026-07-01T16:47:09Z"
capture_tool: "hn-digest"
hn_id: 48748977
score: 7
comments: 0
posted_at: "2026-07-01T15:59:20Z"
tags:
  - hacker-news
  - translated
---

# Stealing 50 Years of Database Ideas for AI Agents

- HN: [48748977](https://news.ycombinator.com/item?id=48748977)
- Source: [onewill.ai](https://onewill.ai/blog/2026/stealing-50-years-of-database-ideas-for-ai-agents/)
- Score: 7
- Comments: 0
- Posted: 2026-07-01T15:59:20Z

## Translation

タイトル: AI エージェントのための 50 年分のデータベースのアイデアを盗む
記事のタイトル: AI エージェントのための 50 年間のデータベースのアイデアを盗む |ワンウィル
説明: OneWill の紹介、または: エージェントにコンピューターを rawdog させるよりも、先行書き込みログの方が優れている理由。

記事本文:
AI エージェントのための 50 年間のデータベースのアイデアを盗む |ワンウィル
ワンウィル
製品
チーム
ブログ
デモ
2026 年 7 月 1 日
AI エージェント向けの 50 年にわたるデータベースのアイデアを盗む
OneWill の紹介、または: エージェントにコンピューターを rawdog させるよりも、先行書き込みログの方が優れている理由。
自律エージェントを盲目的に信頼せずに、自分の状態 (ファイル、カレンダー、電子メール、シェルなど) にアクセスできるようにしたいですか?このビデオをチェックし、ここでデモへのアクセスにサインアップしてください。これは、ブラウザからすべてを試すことができるフル機能のデスクトップ環境です。オープンソース化にも取り組んでいますので、ご期待ください。
OneWill のテクノロジーは舞台裏でデータベース マジックを使用してエージェントのアクションを挿入し、アクションを実行する前にアクションを元に戻すことができるか (詳細は後述)、またはユーザーによって明示的に承認される必要があります。お客様に安心を提供し、エージェントに任せられる業務の幅を広げます。
あなたがシステム開発者であれば、最新のコーディング エージェントはおそらく少し不快に感じますが、コンピュータを使用するエージェントはさらに不快に感じます。エージェントに実際の状態を直接変更させるのはとんでもないことです。しかし、私たちは皆、最終的にはピクルスに陥ってしまったのです。
与える状態が増えるほど、状態を与えない (空の VM) か、ステージングされた環境 (サンドボックス) を与えるかを選択できます。しかし、ステージングが実際に本番と同じように見えることがどれくらいあるでしょうか?エージェントにとって、それらはより便利になります (カレンダー、電子メール、シェルなど)。
しかし、エージェントに与える状態が多ければ多いほど、失敗の影響は大きくなります (データベースと電子メールに別れを告げましょう)。
この取引の最悪の部分は、基本的に自動モードでは信頼性の問題が発生することです。問題がモデルを信頼できないことである場合、なぜ別のモデルを信頼することが解決策になるのでしょうか?エージェントの操作モードは 2 つだけです。承認インザループでエージェントを監視するか、危険な権限をスキップする YOLO です。何か問題が発生した場合、打撃を受けるのはあなたです

。
しかし、このままである必要はありません！私たちは、不安定で、不安定で、時には非常識なエージェントから貴重な国家を守ろうとしています。そうですね、データベースは長い間、不安定でクラッシュし、非常識なハードウェアに対処する方法を学んできました。重要なのは、世界に変更を加えるための適切な意味論的抽象化を特定することです。
クラッシュリカバリーに関する短期集中コース
クラッシュの心配はありません。データを失わないようにしてください。
人生において、回復できる限り、間違いを犯しても大丈夫です。これはクラッシュにも当てはまります。クラッシュ回復に関する完全な講義は、このブログ投稿の範囲外です。ここで見つけることができます。必要な部分だけをまとめておきます。
クラッシュ回復アルゴリズムには 2 つのコンポーネントがあります。1 つは回復可能性を確保するために通常の操作中に実行する記録管理、もう 1 つは破片を拾うための実際の回復手順です。 UNDO と REDO という 2 つの主要な基本アクションがあります。 UNDO は非常に直感的です。不完全な操作または中止された操作の後にクリーンアップします。 REDO は、システムによって行われた約束を裏付けます。システムが何かが起こった (トランザクションのコミットなど) と判断した場合、操作の効果を再適用して現実と一致することを確認します。
最新のデータベース システムのほとんどは、リカバリに先行書き込みログ (WAL) を使用します。基本的な考え方は、操作の効果が持続する前に、クラッシュしてもレコードが確実に残るように、目的の操作に関する十分な回復メタデータを書き留めることです。そうすれば、たとえ途中でクラッシュしたとしても、必要な UNDO および REDO アクションを実行するために必要な情報がすべて手元にあります。たとえば、ウィルがスーパーマーケットのグラインダーを使って 500 g のピーナッツをピーナッツバターにしているとします。この場合、トランザクション T1 は次のようになります。
T1が始まります
アクション 1: ウィルは店のホッパーから 500g のピーナッツを入手します
ストア.ピーナッツ: 前=6700g、後=6200g
Will.ピーナッツ: 前

e=0g、after=500g
アクション 2: ピーナッツを粉砕して容器に入れます
Will.ピーナッツ: 前=500g、後=0g
Will.ピーナッツバター: 前=0g、後=500g
アクション 3: 店舗手数料はピーナッツ バターに 7 ドルかかります
Will.cash: 前=$25、後=$18
Store.cash: 前=$1000、後=$1007
T1コミット
自家粉砕ピーナッツバターを購入すると失敗する可能性がたくさんあります。ホッパーが詰まった場合、ウィルはピーナッツの一部しか入手できない可能性があります。グラインダーが詰まると、ウィルはピーナッツバターを入手できずにピーナッツを失う可能性があります。大まかに言えば、単純に状態を直接変更する場合の問題は、どの段階でもシステムがクラッシュする可能性があることです。優れたシステムは、アクションの効果が持続する前に意図を永続的に記録します。
しかし、WAL を使用すると、クラッシュの処理が比較的簡単になります。 T1 コミット前にシステムがクラッシュした場合、リカバリでは変更前のイメージを使用して書き込みを元に戻します。つまり、コミットされていないアクションの影響を元の値にロールバックします。実際には、すべての操作をきれいに元に戻すことができるわけではありません。ピーナッツバターをピーナッツに戻すにはどうすればよいでしょうか?そのような場合は、ピーナッツを追加注文したり、ピーナッツバターを脇に置いておくなどの補償措置を適用できます。 T1 コミット後にクラッシュした場合、つまりウィルが店にお金を払ってピーナッツバターを持って立ち去ったが、すべての更新された状態が保持される前にクラッシュした場合 (たとえば、レシートの印刷後にレジスターがクラッシュした場合)、リカバリーはアフターイメージを使用して書き込みをやり直します。 UNDO と REDO を組み合わせると、結果は全か無かの永続的なものになります。
この時点で、あなたはこう思うかもしれません：これは非常に簡単ではないでしょうか？はい！それにもかかわらず、どういうわけか、今日のデフォルトのエージェント エクスペリエンスは、状態を荒々しく監視しています。私たちがこれまでに見た中でこの概念に最も近いのは死後の世界です。彼らはエージェントの動作を記録し、その動作によってエージェントを定義することさえあるかもしれませんが、領域があります

息子 WAL が書き込みパスに入ります。
WAL には多くの用途 (レプリケーションなど) がありますが、広い概念としては、通常のデータ管理のために WAL を超えて革新することは困難です。しかし、代わりにアクション管理のために WAL を再考したらどうなるでしょうか?
エージェントは何をしたのですか？なぜ？？それを元に戻してもらえますか?
歴史的に見て、ほとんどのアクションは合理的な人間によって開始されたため、アクション管理はそれほど大きな問題ではありませんでした。私たちは、「この削除 API はおそらく単なるステージングである」などのおかしな仮定を立てたり、重要な指示を無視するような定期的な記憶喪失を引き起こしたりしません。人間は間違いを犯しますが、一般に、無害な行為と危険な行為を区別するのに十分な危機感を培っており、適切な予防措置を講じることができます。
エージェントの動作は異なります。今日のエージェントは、同じ控えめな承認フローで ls と rm -rf への許可を求めるという面白いトリックを仕掛けます。ツール呼び出しが 1 回の承認で大惨事になることを認識するのはあなたの責任です。この問題は、バージョン管理外のファイル (プレゼンテーションなど)、ネットワーク化された API 呼び出し、およびコンピューターの使用を伴う、より重要なタスクに自律エージェントを使用するとさらに悪化します。
かなりの量のトークンを燃やす人々として、私たちは、(1) エージェントが何をしたか、(2) エージェントがやった理由 (どのような入力を見たのか)、(3) その影響を逆転させる簡単な方法があるかどうかを解明するためにフォレンジックを実行していることによく気づきました。実際には、これはクラッシュ回復とよく似ています。データベース システムでは、アクションを正しい順序で適用するため、クラッシュ後にディスク ブロックを手動でスキャンする必要はありません。アクションを適用する前に、そのアクションから回復する方法を記録します。この順序付けをエージェントのアクションに適用すると、ガバナンス メカニズムも提供されます。アクションが実行される前に、

後で元に戻すか、実行する前に明示的な承認を必要とするのに十分な情報が必要です。
Wally: エージェントのアクションに先行書き込みログを適応させる
OneWill では、エージェント時代に向けてデータベース テクノロジーを再考することに取り組んでいます。私たちの主張は、エージェントの実行を効果的に制御するには、投機的なエージェントの作業と永続的な世界の突然変異の間にデータベースのようなトランザクション インターフェイスが必要であるということです。これは、エージェントが長時間実行される場合に特に重要です。軌道が短い場合は、おそらく優れたフォーク技術を使用して回避できるでしょう。しかし、ゴール モードは定期的に数時間または数日間実行されるため、ロールバック戦略として「システム全体を吹き飛ばす」のはかなり無駄です。これが単に git であるとデフォルトで考える人もいます。しかし、現在でも、カレンダーや電子メールなど、多くの有用な状態がバージョン管理の外に存在しています。ここで、それにコンピューターの使用を加えます。コンピュータを直接使用します。
つまり、WAL アーキテクチャが今日のエージェントの制御性の問題に対する答えであると私たちは考えています。これは、多数のエージェントのカオスを意味論的なアクション ストリームに構造化するための原則に基づいた方法を提供します。しかし、同様の可逆性やガバナンスの領域における既存のソリューションでは、ソース コードをインストルメントするか、エージェントを変更する必要があり、その負担は依然としてユーザーにあります。問題はすべてユーザー エラーです。頑張ってください。
OneWill の最初のリリースである Wally が登場します。これは、数十年にわたって実績のあるデータベース WAL アーキテクチャに基づいており、現在はエージェントのアクションに適用されています。私たちの目標は、エージェントのアクションを正確に制御できるドロップイン エクスペリエンスを提供することです。ここでの洞察は、エージェントを変更しなくても、エージェントの周囲の世界を変更することでエージェントの動作を変更できるということです。
概念的に、このプロジェクトの最も興味深い部分は、任意の ag の UNDO および REDO セマンティクスを定義することです。

エントのアクション。これまでに 3 種類のアクションが発生しました。
完全に可逆的: git 追跡ファイルと同様。
補償可能: カレンダーのイベントをキャンセルすることはできますが、招待状の送信を取り消すことはできません。
取り消し不可: 一度送信したメールは取り消すことができません。
これがどのように機能するかを簡単に説明します: WAL、FUSE、VPN、MITM、フレーム バッファー、インターポジション。このブログ投稿の長さを制限し、読者の関心に基づいて優先順位を付けるため、詳細な技術的な詳細については今後の投稿に延期します。 Discord サーバーに参加して、次に何を書くべきか投票してください。
これを書いている時点では、フロンティア モデルを使用してこれを自分でバイブコーディングしようとすると、おそらく一連の不幸な制作イベントに遭遇することになるでしょう (状況が変わったらメールでお知らせください!)。ただし、これを自分で構築する必要はありません。代わりに、私たちが作成したものを使用したり、貢献したりすることができます。
現在、私たちの公開プロトタイプは、アクション ストリーム、コーディング エージェントの統合 ( claude 、 codex 、 agy )、ローカル AI 使用の強制のサポート、および私たちのテクノロジーに関するいくつかのユースケースを示しています。これには、「エージェントが夜間に何をしたか」を示す強化された git diff、「エージェントがこの変更を行ったときに何を見たか」を示す強化されたファイルシステム ビュー、および基本的なロールバック機能が含まれます。これがビデオです:
今すぐウェブ上で無料でお試しください: デモ アクセス 。アクセスコードを記載した手順をご案内します。現在、実際に Ubuntu と macOS でベータ ビルドをローカルで実行できますが、macOS の権限メニューはまだ簡素化中であるため、macOS ユーザーには今のところ Web バージョンを使用することをお勧めします。
今後数週間以内に、コア WAL とさまざまな便利なプリミティブをオープンソース化することを目指しています。その際はDiscordでお知らせします。乞うご期待！
募集中の役割を Discord サーバーに投稿します。
この記事の執筆時点では、次に採用される可能性があるのはエンジニアリング部門です。私たちは (1) 経験を大切にします

エージェントとの連携 (例: 1 日に 5 億から 10 億のトークンを生産的に燃やすことに慣れている必要があります)、(2) データベース システムの背景 (例: CMU の 15-445 のオンライン バージョンを完了すること)。そう思われる場合は、ぜひご連絡ください。
コラボレーションやデザイン パートナーシップについては、 [email protected] まで電子メールでお問い合わせください。
Discord サーバーに参加してブログのトピックに投票し、新しい投稿が公開されたときに通知を受け取ります。
© 2026 OneWill-AI Society d/b/a OneWill。

## Original Extract

Introducing OneWill, or: why write-ahead logging is better than letting agents rawdog your computer.

Stealing 50 Years of Database Ideas for AI Agents | OneWill
OneWill
Product
Team
Blog
Demo
Jul 1, 2026
Stealing 50 Years of Database Ideas for AI Agents
Introducing OneWill, or: why write-ahead logging is better than letting agents rawdog your computer.
Want to give autonomous agents access to your state (e.g., files, calendar, email, shell) without blindly trusting them? Check this video out and sign up for demo access here : it's a full-featured desktop environment that lets you try everything from your browser. We're also working on open-sourcing, stay tuned!
Behind the scenes, OneWill's technology interposes agent actions with database magic to make it so that before any action runs, it must either be reversible (details below) or explicitly approved by you. We provide you peace of mind and broaden the scope of work that you can hand off to agents.
If you’re a systems developer, modern coding agents probably make you slightly uncomfortable, and computer-use agents even more so. Letting agents mutate real state directly is bonkers. But we’ve all ended up in a pickle:
The more state you give You could choose to give no state (blank VM) or a staged environment (sandbox). But how often does staging actually look like production? your agents, the more useful they are (e.g., calendar, email, shell).
Yet the more state you give your agents, the bigger the consequence of failure (say goodbye to your database and email ).
The worst part of this deal is that there are essentially Auto mode gives us trust issues. If the problem is that we don’t trust models, why is the solution to trust another model? only two modes of agent operation: babysit it with approval-in-the-loop, or YOLO by dangerously skipping permissions. If anything goes wrong, you take the hit.
But it doesn’t have to be this way! We’re trying to protect our valuable state from flaky, crashy, and occasionally insane agents. Well, databases have long learned to deal with flaky, crashy, and insane hardware. The trick is in identifying the right semantic abstraction for making changes to the world.
A crash course on crash recovery
Don’t worry about crashing. Just never lose data.
In life, it’s ok to make mistakes as long as you can recover from them. This applies to crashing too. A full lecture on crash recovery is outside the scope of this blog post; you can find one here . We’ll just recap the bits we need.
Crash recovery algorithms have two components: the bookkeeping you perform during normal operations to ensure recoverability, and the actual recovery procedures for picking up the pieces. There are two key primitive actions, UNDO and REDO. UNDO is pretty intuitive: it cleans up after incomplete or aborted operations. REDO backs the promises made by the system: if the system says something happened (e.g., transaction commit), it re-applies the effects of operations to make sure that reality matches.
Most modern database systems use a write-ahead log (WAL) for recovery. The basic idea is to write down enough recovery metadata about the intended operation, in a way that ensures the record survives crashes, before the operation’s effects are allowed to persist. That way, even if you crash halfway through, you have all the information required to perform the necessary UNDO and REDO actions. For example, suppose Will is using a supermarket’s grinder to turn 500g of peanuts into peanut butter. Then transaction T1 is:
T1 begins
Action 1: Will obtains 500g of peanuts from the store’s hopper
Store.peanuts: before=6700g, after=6200g
Will.peanuts: before=0g, after=500g
Action 2: Will grinds the peanuts into a container
Will.peanuts: before=500g, after=0g
Will.peanut_butter: before=0g, after=500g
Action 3: Store charges Will $7 for the peanut butter
Will.cash: before=$25, after=$18
Store.cash: before=$1000, after=$1007
T1 commits
There are many ways that buying self-ground peanut butter can go wrong. If the hopper gets stuck, Will may only get some of the peanuts. If the grinder clogs, Will may lose his peanuts without getting peanut butter. Broadly speaking, the problem with naively modifying state directly is that the system may crash at any step. A good system records intent durably before allowing an action’s effects to persist.
But with a WAL, handling crashes is relatively straightforward. If the system crashes before T1 commits, recovery uses the before images to UNDO the writes: roll back the effects of any uncommitted actions to their original values. In practice, not every action can be undone cleanly: how do you turn peanut butter back into peanuts? In those cases, you can apply compensating actions, such as ordering more peanuts or setting the peanut butter aside. If it crashes after T1 commits, meaning Will has paid the store and walked away with peanut butter, but before all the updated state is persisted (e.g., the register crashes after printing the receipt), recovery uses the after images to REDO the writes. Together, UNDO and REDO make the outcome all-or-nothing and durable.
At this point, you might be thinking: isn’t this pretty straightforward? Yes! And yet somehow today’s default agent experience is rawdogging your state. The closest we’ve seen to this concept in the wild is still post-mortem; they record what the agent does, they may even define an agent by what it did, but there’s a reason WALs get into the write path.
Although WALs have many uses (e.g., replication), as a broad concept, it’s hard to innovate beyond the WAL for regular data management. But what if we re-imagine the WAL for action management instead?
The agent did what?! Why?? Can you undo that?
Historically speaking, action management wasn’t that big an issue because most actions were initiated by reasonable humans. We don’t make crazy assumptions like “this deletion API is probably just staging ” or develop periodic amnesia that leads to us ignoring critical instructions . Even though humans make mistakes, we generally develop enough of a sense of danger to distinguish between harmless actions and risky actions, allowing us to take the appropriate precautions.
Agents operate differently. Today’s agents play a funny trick on you by asking for your permission to ls and rm -rf in the same unassuming approval flow. The burden is on you to notice that a tool call is one approval away from creating a disaster . This is exacerbated by using autonomous agents for more consequential tasks that involve files outside of version control (e.g., presentations), networked API calls, and computer-use.
As people who burn a decent amount of tokens, we’ve often found ourselves performing forensics to figure out (1) what an agent did, (2) why it did it (e.g., what input it saw), and (3) whether there was an easy way to reverse its effects. In practice, this looks a lot like crash recovery. Database systems don’t make you manually scan disk blocks after a crash because they apply actions in the right order: they record how to recover from an action before applying it. Applying this ordering to agent actions also provides a governance mechanism: before an action is executed, we either record enough information to reverse it later or require explicit approval before letting it run.
Wally: adapting write-ahead logging for agent actions
At OneWill, we’re in the business of re-imagining database technology for the agentic age. Our thesis is that to effectively control agent execution, we need a database-like transactional interface between speculative agent work and durable world mutations. This is especially important as agents run for longer If trajectories were short, you could probably get away with good forking technology. But we regularly run goal mode for hours or days, so it's pretty wasteful to have "blow away the entire system" as a rollback strategy. and interact Some people default to thinking that this is simply git. But even today, a lot of useful state lives outside of version control, such as calendar and email. Now add computer-use on top of that. with computers directly.
In short, we think that the WAL architecture is the answer to today’s agent controllability problems. It provides a principled method for structuring the chaos of numerous agents into a semantic action stream. But existing solutions in a similar reversibility and/or governance space require you to instrument your source code or modify the agent, where the burden is still on you - any problems are user error, good luck!
Enter Wally, OneWill’s first release. It is based on the decades-proven database WAL architecture, now applied to agent actions. Our goal is to provide a drop-in experience that delivers precise control over agent actions. The insight here is that even without changing the agent, you can still change its behavior by modifying the world around it.
Conceptually, the most interesting part of this project is defining UNDO and REDO semantics for arbitrary agent actions. We've encountered three types of actions so far:
Perfectly Reversible : like a git-tracked file.
Compensable : you can cancel a calendar event but you cannot unsend the invitations.
Irreversible : once you send an email, you can’t take it back.
Quick sketch of how this all works: WAL, FUSE, VPN, MITM, frame buffers, interposition. To limit the length of this blog post and prioritize based on reader interest, we’re deferring in-depth technical details to future posts; join our Discord server to vote on what we should write about next.
At the time of writing, you will probably encounter a series of unfortunate production events if you try vibe-coding this yourself with the frontier models (shoot us an email if that changes!). But you don’t have to build this yourself, you can use and contribute to what we made instead.
Our public prototype currently demonstrates the action stream, coding agent integrations ( claude , codex , agy ), support for enforcing local AI usage, and some use cases around our technology. That includes: enhanced git diff for "what did the agent do overnight," enhanced filesystem view for "what did the agent see when it made this change," and basic rollback functionality. Here's a video:
Try it now on the web for free: demo access . We'll provide instructions with your access code. You can actually run a beta build locally today on Ubuntu and macOS, but because we're still simplifying the macOS permissions menus, we encourage macOS users to stick with the web version for now.
We aim to open-source the core WAL and various useful primitives in the coming weeks. We'll announce on Discord when we do. Stay tuned!
We post open roles to our Discord server .
At the time of writing, our next likely hire is in engineering. We value (1) experience with agents (e.g., you should be comfortable productively burning 500M to 1B tokens a day) and (2) a background in database systems (e.g., completing the online version of CMU's 15-445 ). If this sounds like you, please reach out!
For collaborations and design partnerships, please email [email protected] .
Join our Discord server to vote on blog topics and get notified when new posts go live!
© 2026 OneWill-AI Society d/b/a OneWill.

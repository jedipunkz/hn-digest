---
source: "https://momo5502.com/posts/2026-08-17-mw2-decompilation/"
hn_url: "https://news.ycombinator.com/item?id=49351299"
title: "200B Tokens Later: A Month of Letting AI Agents Decompile MW2"
article_title: "200 Billion Tokens Later: A Month of Letting AI Agents Decompile MW2 | Maurice's Blog 🐍"
image: "https://momo5502.com/posts/2026-08-17-mw2-decompilation/images/cover.png"
author: "Philpax"
captured_at: "2026-08-18T20:13:07Z"
capture_tool: "hn-digest"
hn_id: 49351299
score: 4
comments: 1
posted_at: "2026-08-18T19:28:09Z"
tags:
  - hacker-news
  - translated
---

# 200B Tokens Later: A Month of Letting AI Agents Decompile MW2

- HN: [49351299](https://news.ycombinator.com/item?id=49351299)
- Source: [momo5502.com](https://momo5502.com/posts/2026-08-17-mw2-decompilation/)
- Score: 4
- Comments: 1
- Posted: 2026-08-18T19:28:09Z

## Translation

タイトル: 200B トークンのその後: AI エージェントに MW2 を逆コンパイルさせた 1 か月
記事のタイトル: 2,000 億トークンのその後: AI エージェントに MW2 を逆コンパイルさせた 1 か月 |モーリスのブログ 🐍
説明: ちょっとしたサイド プロジェクトとして、Claude Max サブシステムを有効活用し、Call of Duty: Modern Warfare 2 (2009) を逆コンパイルすることにしました。
目標は、プロジェクトを C++ コードに逆コンパイルし、可能な限り元のコードに近づけることですが、必要に応じて調整を加えます。
つまり、por に追加の変更が必要になります
[切り捨てられた]

記事本文:
モーリスのブログ 🐍
ホーム
投稿
2,000 億トークンのその後: AI エージェントに MW2 を逆コンパイルさせた 1 か月
ちょっとしたサイド プロジェクトとして、Claude Max サブシステムを有効活用して、Call of Duty: Modern Warfare 2 (2009) を逆コンパイルすることにしました。
目標は、プロジェクトを C++ コードに逆コンパイルし、可能な限り元のコードに近づけることですが、必要に応じて調整を加えます。
これは、移植性 (32 および 64 ビット、Linux、macOS などのサポート)、セキュリティ、安定性の向上のための追加の変更が行われることを意味します。
このプロジェクトは、 RektInator 、 Future 、 st0rm およびコミュニティの他のメンバーの協力を得て行われています。
リポジトリは今のところ非公開です。
私は、Claude Code CLI を Max (20x) サブスクリプションで使用しています。
従業員は、共通の目標に向かって協力する 4 人のエージェントで構成されています。
3 つのワーカー エージェントがあり、ゲームのさまざまなサブシステムで独立して動作します。
これらはすべて、リポジトリの同じブランチにコミットしてプッシュします。
追加のエージェントが操作を監督し、プッシュされたすべてのコミットをレビューします。
現在、すべてのエージェントが Sonnet 5 を使用しています。
エージェントは Discord を介して通信します。
すべてのエージェントは 1 つのチャネルにアクセスでき、そこにあるすべてのメッセージを投稿したり読んだりできます。
監督者は、GitHub Webhook が到着した新しいコミットに関する通知を投稿する追加のチャネルにアクセスできます。
これにより、エージェントのレビュー タスクが自動的にトリガーされます。
追加の Webhook が CI エラーを共有チャネルに投稿するため、何かが壊れたときにエージェントに通知が届きます。
エージェント自体は、GitHub Issue を使用してタスクを管理します。
逆コンパイル用の興味深いリソースが多数あります。エージェントが使用するシンボルを含むリークされた Xbox アルファ ビルド、デバッグ データを含む macOS ポートなどです。
これらはすべて、MCP サーバー経由で Ghidra および IDA Pro にアクセスできます。
エージェントは約 4 週間ノンストップで稼働しています。彼らは

約 7,000 のコミットが生成されました。
これまでに 16,324 関数のうちおよそ 5,588 関数 (約 34%) が逆コンパイルされました。
ただし、実効値はさらに高くなります。これら 16,000 個のうちかなりの部分は、サードパーティのライブラリ、CRT コード、および逆コンパイルを必要としないその他のものです。
エージェントはそこに到達するまでに 1,998 億のトークンを使い果たしました。
ただし、欠落しているものや壊れているものもたくさんあります。
必要なサブシステムの多くが実装されているにもかかわらず、マップの起動はまだ機能していません。
これは、ほとんどすべてが逆コンパイルされ、接続されて初めてテスト可能になる可能性があります。
当初、進行状況の追跡にはリポジトリ内の STATUS.md ファイルが使用されていました。
そのファイルはすぐに 10 MB を超える巨大なものに成長し、取り込もうとするとすぐにコンテキストからあふれてしまいました。
結果として、GitHub の問題に切り替えましたが、これはうまくいきました。
エージェントは自律的にそれらを作成、編集し、閉じます。
驚くべきことに、Discord でのコミュニケーションも完璧です。
エージェントは問題なく調整します。彼らは、少なくともほとんどの場合、自分宛てでないメッセージを無視し、無関係な情報で他人の文脈を不必要に埋めないように、必要な場合にのみメッセージを投稿します。
Discord チャネルを使用すると、エージェントが実行されているマシンへのアクセスを他の参加者に許可する必要がなく、私たち人間が彼らと対話することもできます。
…残念ながら、うまく機能しないことがたくさんあります。
Discord ではエージェントに簡潔に話すようお願いしています。彼らは、ちょうど 5 分間、再びテキストの壁を書き始めます…
このようなメッセージが 1 つあれば問題ないかもしれませんが、エージェントはその Discord チャネルで積極的にコミュニケーションを行っています。
それらのメッセージの多くは、ステータス更新など、私たち人間に向けられたものです。誰もそれほど多くのテキストを読みたくありません。
もう 1 つの問題点は、ローカル テストの実行です。何らかの理由で、エージェントは過度に

注意深い。わずかな変更があるたびに、テスト スイート全体をローカルで実行したいという衝動に駆られます。
ローカル テストの実行は、ほとんどの場合、緑色になります。したがって、テストの実行が成功するたびに、コードの逆コンパイルに費やされる可能性がある約 4 分の時間を意味します。
彼らは、決してローカルでテストを実行しないように繰り返し言われました。そのためにCIがあります。
コミットによってエラーが発生した場合は、GitHub から Discord 経由でほぼ即座に通知が届きます。失敗は決して見逃されません。それでも、彼らはあまり気にしません。
問題になりつつあることの 1 つは、エージェントがより大きな問題を始めることを恐れていることです。彼らは、大規模なサブシステムよりも迅速な勝利を好みます。
より大きな障害に遭遇すると、すぐに問題を提起し、問題を延期して、より小さな問題に移ります。
最初はそれで良かったんです。現在、簡単に解決できる成果はほとんどなくなったため、エージェントがアイドル状態になることがあります。
時々、エージェントはコンテキストが不足していると考えて、作業を拒否し始めます。
明らかに、そうではありません。彼らはコンテキスト内にどれだけのトークンがあるのか​​わかりません。たとえ知っていたとしても、それはハーネスが処理します。
オーバーフローが発生すると、自動圧縮が開始されます。
悲しいことに、GitHub にもこれまでのところ 1 日おきに問題が発生しています。そのため、問題追跡と CI の信頼性が比較的低くなります。
なんとか稼働時間を改善できることを願っていますが、そうでない場合は、代替手段に切り替えるのが合理的です。
エージェントに一度に多くのタスクを与えると、長期間にわたって効果が得られません。
逆コンパイル、C++ コードの最新化、移植性の向上をすべて一度に行うように指示しても、うまくいきません。
10 回圧縮すると、最終的にそのうちの 1 つを忘れるか、間違いを犯すようになります。
最初に純粋に機械的に逆コンパイルし、その後で最新化/移植性を高めるのが最善の方法です。
私はオーパス5をfのために実行しました

最初は 2 週間試しましたが、最終的には途中で Sonnet に切り替えることにしました。若干鈍くなっていますが、目立った違いはありません。
私の知る限り、ソネットはオーパスと同じ間違いを犯しており、同じように簡単に気が散ってしまいます。
結局のところ、逆コンパイルは比較的簡単な作業です。これは、MCP で IDA を呼び出し、それを CPP ファイルに貼り付けるだけです。ソネットはそのためにうまく機能します。
私たちが設定したルールへの関心を高めるために、一連のルールをコンテキストに再挿入する PostCompact フックを追加しました。
これはある程度機能しているようです。彼らは GitHub の問題を長期間にわたって適切に管理し、その方法で Discord に更新を投稿することを忘れません。
しかし、メッセージを簡潔に保つことや、冗長なコード コメントを省略することは、彼らが尊重していないことです。
私の推測では、コードとメッセージの両方で、ヤッピングがトレーニング中に深く埋め込まれただけであると考えられます。
すべてがうまくいったわけではありませんが、この実験は依然として成功しています。
全体的に逆コンパイルは非常にうまく機能します。コードには 1 ～ 2 回のリファクタリングが必要ですが、それでも必要だったと思います。
MW2 は古いゲームですが、改善の余地がたくさんあります。
複数日および複数週間にわたるセッションに取り組む場合、別の種類の問題が発生し始めるのは非常に興味深いことです。
問題のほとんどは圧縮に起因すると考えられ、そのためエージェントが時間の経過とともに重要なことを忘れてしまいます。
以前は 1 時間のセッションしか行ったことがありませんでしたが、それでも圧縮による害は発生しますが、同程度ではありませんでした。
プロジェクトは完了には程遠い。物事がどこに向かっていくのか、そしてそれがどのように進化するのかを見てみましょう。何か面白いことが起きたときに最新情報を投稿するかもしれません。
この時点で、読者の皆様からのフィードバックを大歓迎します。
私が抱えていた問題を改善するアイデアがある場合は、お気軽に経験を共有してください。

私のソーシャル経由で私に連絡します。

## Original Extract

As a little side project, I decided to put my Claude Max sub to good use and decompile Call of Duty: Modern Warfare 2 (2009).
The goal is to decompile the project to C++ code, as close to the original code as possible, but with adjustments where they make sense.
That means additional changes for por
[truncated]

Maurice's Blog 🐍
Home
Posts
200 Billion Tokens Later: A Month of Letting AI Agents Decompile MW2
As a little side project, I decided to put my Claude Max sub to good use and decompile Call of Duty: Modern Warfare 2 (2009).
The goal is to decompile the project to C++ code, as close to the original code as possible, but with adjustments where they make sense.
That means additional changes for portability (32 & 64 bit, Linux, macOS, etc. support), security and stability improvements will be done.
This project is done with the help of RektInator , Future , st0rm and other members of the community.
The repository is private for now.
I’m using Claude Code CLI with a Max (20x) subscription.
The workforce consists of 4 agents working together towards the common goal:
There are 3 worker agents, independently working on different subsystems of the game.
All of them commit and push to the same branch of the repository.
An additional agent oversees the operation and reviews every pushed commit.
All agents currently use Sonnet 5.
Agents communicate via Discord .
All agents have access to one channel and can both post and read all messages in there.
The overseer has access to an additional channel, where a GitHub webhook posts notifications about new commits that landed.
That automatically triggers the agent’s review task.
An additional webhook posts CI failures in the shared channel, so agents get notified when something broke.
The agents themselves manage their tasks by using GitHub issues.
There are many interesting resources for decompilation: Leaked Xbox alpha builds with symbols, macOS ports with debug data, etc. that the agents use.
They all have access to Ghidra and IDA Pro via MCP servers.
The agents have been running non-stop for about 4 weeks now. They have produced almost 7,000 commits:
Roughly 5,588 of 16,324 functions have been decompiled so far: about 34%.
The effective figure is higher, though: a good chunk of those 16k are third-party libraries, CRT code, and other things that will never need decompilation.
The agents have burned through 199.8 billion tokens getting there.
However, there is also a big chunk of things that are missing or broken:
Launching a map is not working yet, despite many of the required subsystems being implemented.
This will likely only be testable as soon as almost everything has been decompiled and wired up.
Initially, progress tracking used a STATUS.md file in the repository.
That file quickly grew into a 10 MB+ giant, instantly overflowing the context when trying to ingest it.
As a consequence, we switched to GitHub issues, which works great.
Agents autonomously create, edit and close them.
Surprisingly, communication over Discord is also flawless.
Agents coordinate without any issues. They ignore messages not directed to them and only post messages when necessary to not needlessly fill the context of others with unrelated information, at least most of the time.
The Discord channel also enables us humans to interact with them, without me having to give the other participants access to the machines the agents are running on.
… there is a bunch of things that don’t work well, unfortunately.
We ask the agents to be concise in Discord. They are, for exactly 5 minutes, then start writing walls of text again…
One message like this might be fine, but agents actively communicate in that Discord channel.
Many of those messages are directed to us humans, status updates, etc. Nobody wants to read that much text.
Another pain point is local test execution. For some reason, agents are overly careful. On every little change, they feel the urge to execute the entire test suite locally.
Local test execution is almost always green. Therefore every successful test run means about 4 minutes of time that could have been spent decompiling code instead.
They were repeatedly told to never execute tests locally. We have CI for that.
If a commit causes a failure, there will be an almost instant notification from GitHub via Discord. Failures will never go unnoticed. Yet, they don’t really care.
One thing that is becoming a problem is that agents are afraid of starting bigger issues. They favor quick wins over larger subsystems.
The moment they stumble upon bigger blockers, they file an issue, defer the problem and move to something smaller.
Initially, that was fine. Now that low-hanging fruit is mostly gone, this leads to agents being idle at times.
From time to time, agents start refusing to work, because they think they’re running out of context.
Obviously, they don’t. They have no idea how many tokens are in their context, and even if they did, that’s something the harness takes care of.
Auto compaction will kick in when there is an overflow.
Sadly, GitHub also has had issues every other day so far. That makes issue tracking and CI relatively unreliable.
I hope they manage to improve their uptime, otherwise switching to an alternative makes sense.
Giving the agents too many tasks at once just doesn’t work over a long period of time.
Telling them to decompile, modernize C++ code and improve portability all at once just doesn’t work.
They either eventually forget one of those things after the 10th compaction, or start making mistakes.
Pure mechanical decompilation first, and then modernization/portability afterwards is the way to go.
I ran Opus 5 for the first 2 weeks and eventually decided to switch to Sonnet midway. It is slightly dumber, yet there is no noticeable difference.
As far as I can tell, Sonnet is making the same mistakes as Opus and gets distracted just as easily.
After all, decompilation is a relatively easy task. It’s just an MCP call to IDA and then pasting that into a CPP file. Sonnet works fine for that.
To improve their focus on the rules we set, I added PostCompact hooks that reinject a set of rules into the context.
This seems to work to some extent, they properly manage GitHub issues over longer periods of time and don’t forget to post updates on Discord that way.
But staying concise with their messages or omitting redundant code comments is just something they don’t respect.
My best guess is that yapping, both in code and in messages, was just deeply embedded during their training.
Although not everything went well, this experiment has still been a success.
Overall decompilation works very well. The code will need a round or two of refactoring, but I think it needed that regardless.
MW2 is an old game with lots of potential for improvement.
It’s very interesting that a different class of problems starts to arise when working on multi-day and multi-week sessions.
I think most of the issues can be attributed to compaction and thus agents forgetting what matters over time.
Previously, I have only ever had hour-long sessions, where compaction still causes harm, but not to the same extent.
The project is far from done. We’ll see where things are going and how it evolves. I might post updates when something interesting happens.
At this point, I’m openly inviting you, the reader, to give feedback.
If you have ideas of improvement for the issues I have had, feel free to share your experience with me via my socials.

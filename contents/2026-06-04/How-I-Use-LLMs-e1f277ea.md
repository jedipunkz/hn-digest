---
source: "https://pauldambra.dev/2026/06/how-i-use-llms-5.html"
hn_url: "https://news.ycombinator.com/item?id=48403552"
title: "How I Use LLMs?"
article_title: "How I use LLMs? - five"
author: "speckx"
captured_at: "2026-06-04T21:37:48Z"
capture_tool: "hn-digest"
hn_id: 48403552
score: 2
comments: 0
posted_at: "2026-06-04T19:33:59Z"
tags:
  - hacker-news
  - translated
---

# How I Use LLMs?

- HN: [48403552](https://news.ycombinator.com/item?id=48403552)
- Source: [pauldambra.dev](https://pauldambra.dev/2026/06/how-i-use-llms-5.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T19:33:59Z

## Translation

タイトル: LLM はどのように使用すればよいですか?
記事のタイトル: LLM はどのように使用すればよいですか? - 5
説明: LLM の使用方法 - 5

記事本文:
LLM はどのように使用すればよいですか? - 5
思慮のないとりとめのないナンセンス
ポール・ダンブラ
ファングラー
ブログ投稿
週のメモ
キッズゲーム
2026年6月4日木曜日
LLM はどのように使用すればよいですか? - 5
私は、LLM と拡張コーディングの現在の最高水準についての皮肉の約 60% を失いました。この 9 か月間で、私たちの仕事への取り組み方が大きく世代交代したと確信しています。
Let's Growth Hack ブログの訪問者数は、効果があると思われる部分とそうでない部分についてのちょっとした反省として、今日のブログの使用方法を記録します。技術（およびそれに関する私のスキル）が向上するにつれて再検討できるもの
これは、数え方に応じて更新番号 4 またはエントリ番号 5 になります。
パート 1 - 魅惑と懐疑の構造化されたトーン
パート 2 - それらを管理および制御する方法を学ぶ
パート 3 - しぶしぶ受け入れますが、コントロールを手放したときに彼らがしてくれることへの感謝 (?)
パート 4 - たとえそれがもはや「脳コード」ではなく、もはや人間のためのものではないとしても、それは依然として人間の脳から来ていることに気づく
職場の誰かが私のためにエピソードに名前を付けてくれました、とても気に入っています
それでは…パート 5 - やあ、ポールは AI を服用しています
AIホットテイクの今日版で
コードに集中しなくなり、集中力が 95% から 5% 低下しました
私はこれらのツールを使用せず、チームに採用しています
自主性がないときの本当の不満
いつ逃げるべきか私に指示させるのはやめてください
まず、必要ですが、私がツールを一般的にどのように使用しているかについて簡単に話します。
私は開発者として主に独学で学び、(vbscript の後 😱) DotNet から始めました。私は早い段階で JetBrains Resharper に出会い、決して振り返ることはありませんでした。コードの作成、リファクタリング、改善のための IDE の支援は、状況を一変させました。私は、レベルアップに使用したいツールを誰も恨んでいません。グラフィカル git クライアント、emacs、またはその他のものを使いたい... さあ、クールなものを作りましょう
インダストリーズへの参入のせいで

試してみてください、私も VSCode にそれほど熱心ではありません (好きですか? 素晴らしいです。上記を参照してください。クールなものを作りに行きましょう)。私は賢い介入が私を助けてくれるのに慣れすぎていて、VSCode は必要最低限すぎる (そして SublimeText よりもはるかに遅い)
そのため、私はツールをワークフローに直接関与させるのが好きです。書き込み時に resharper がインラインで「これを気に入ったらどうですか?」とプロンプトを表示してくれたので、LINQ をより迅速かつ徹底的に学習できました。
新しい会話を開いてプロンプトとまったく同じテキストを書くことができるのに、なぜ To Do リストに何かを書き留める必要があるのでしょうか。私はすぐに切り替えて、To Do リストのように一日を続けます。
しかし、私のエージェントはコードとすべての分析に接続されているため、スイッチバックするときはほぼ毎回、コンテキストと課題が発生します。またはPR
コードに集中しなくなり、集中力が 95% から 5% 低下しました
エージェントとその UX が最近改善されたため、必要に応じて積極的にコードを記述するだけで済みます。完全にエージェント経由で作業することも可能です。
もう文字通りコード行を手動で書くことはできませんが、それでも大丈夫です。
なぜなら、コードを理解するコストは現在では基本的にゼロだからです。コードが理解できることの重要性は低下しています。
正しくて速くなければなりませんが、すぐに理解できない場合は。 LLM に説明を依頼することも、それを実行するためのテストをいくつか作成することもできます。以前は理解するのに数日かかったコード作業も、今では数秒で完了します。
コードを書くことは、私が自分の時間を使ってできる最大限の活用からはほど遠いです。
「ああ、でもポールにはとても重要なコードがあるんだよ！」
そうです、素晴らしいことに、リテラルコードは依然として、実行できる最も高い活用方法である可能性があります。しかし、それは私ができる最高のレバレッジではありません。そして、多くの人にとってもそれは真実だと私は主張します。
これは本当に魔法のようです。

たくさんのことを同時に達成することは、2025 年のポールには不可能だったでしょう。
私はこれらのツールを使用せず、チームに採用しています
2 つの木材を釘で打ち合わせる場合、どうやってそれを行うかは私次第です。ハンマーを選ぶこと。ネイルを選ぶこと。爪が入る場所と角度を決めるため。親指が当たらないようにするため。
大工を雇う場合は、大工にビールを淹れて、必要な場合はここにいますと伝えます。
そこで、私が提案している考え方の転換は、ジュニア エンジニアやプロダクト マネージャーなどを雇用することです。そして、私にはそれらにコンテキストとガードレールを与える責任があります。そして私は彼らに仕事をさせました。その仕事に価値があるかどうかについては私が責任を負っていますが、文字通り完璧に取り組んでいるわけではありません。
物事との関わり方によって大きく変わります。そして、仕事のやり方を完全に変えることができます。私は重役に異動しましたが、ハンマーを持たず、大工を雇っているので、まだたくさんのコードを投げることができます。
自主性がないときの本当の不満
私たちは PR に対して多数の AI レビュー ツールを実行しています。そしてそれらにはさまざまな品質があります。時間の経過とともに少しずつ良くなっていきます。そして、彼らは生活費を稼いでいます（かなり高価ですが）
しかし、パート 4 からここ数か月間、私は私に新たな期待を抱いていることに気づきました…彼らがコメントを投稿すると、「ああ、今私はそれを読んで行動しなければならない…なんてくだらない」と思いました。
それに比べて、私は「広報羊飼い」のスキルを持っています。クロードに C2 wiki で XP について調べてもらいました。仕事のやり方で私が何を気にしているのかを知るために、ソフトウェア エンジニアリングについてインタビューしてもらいました。そして、私がレビューする際に何を重視しているのかを確認するために、私の PR レビューの公開履歴をレビューしてもらいました。
そのスキルを実行すると、多数のレビューが並行して実行され、単一の解決策があるものはすべて実行されます。

それはただそうなります。ここで例を確認できます。私はこれを実行し、すべての PR に公開コメントを付けています。ロボットに物を見つけてもらう前に人間がレビューするのは失礼な気がします。長い間、ツールで実行できるのに人間がコードをリントするのは失礼だとされてきたのと同じです。
そのため、ワードマシンは、ある単語のセットを別のセットに変換するだけでは、ますます収益を得ることができなくなりました。その一連の言葉を行動に移す必要があります。私が若手エンジニアを雇って、「この明らかなバグを修正すべきか」と尋ねられたら、私はイライラするでしょう。
いつ逃げるべきか私に指示させるのはやめてください
これには副作用があります…LLM にたくさんのことを依頼することを忘れる必要はありません。同様に、同僚にたくさんのことについて尋ねる必要もありません。
「おい、Clink Expanderを2時間ごとにチェックして、ファングラーを適切にルーティングする必要がある」と言いたいが、それはたまたま起こることだ
次の大きなフロンティアは、長期実行型のトリガー駆動型エージェントだと思います。そしてとても興奮しています
上で「PR羊飼い」について触れました。レビューを通じて PR を取得するための機械的な作業をすべて実行します。 PR が開いているかどうかを確認し、開いていない場合は開きます。レビューの準備ができていることをマークし、並行レビュー担当者の群れを実行して応答し、外部レビューをチェックして応答します。必要な場合は私に委ねられますが、そうでない場合は、緑色になるまで CI をチェックします。ここまでの作成に合計で約 1 時間かかりました。そしてそれの何倍も私を救ってくれました。
私は「ポールペア」を持っています。ロボットが仕事を終えたり、私に質問しようとしたりするたびに、「ポールペア」スキルを実行するように指示されます。これにより、私が関心のあることを思い出させ、重要なことだけを中断するようになります。メカポールです。私が明らかに繰り返すつもりだった事を繰り返さないようにミートポールを救いました。
私にはエグゼクティブコーチがいます（確かにそうではありません）

スキル）。グラノーラ、スラック、github、組織図などをチェックし、重要だと述べたことに取り組んでいるかどうかを週に一度教えてくれます。私に何が欠けているのか、どのようなパターンが見られるのか、軌道に戻るためにすべき上位 3 つのことは何かを教えてくれます。それが常に正しいとは限りませんが、必ずしも正しいとは限りません。どのようなアドバイスを受け取っても、自分の仕事と注意力を管理するのは私の責任です。
やったことがないなら、今やっていることをやめてください。お気に入りのロボットと新しいセッションを開き、「私たちの会話を見て、自動化できる繰り返しの対話を教えてください」と入力します。
それから、座って、心を吹き飛ばす準備をしてください。
ポール、あなたは救いようのないほどAIに悩まされているようですね
はい、そうです。ごめんなさい、そうなってしまいました。もう後戻りはできません。バブルに吸い込まれてしまった。これが未来だ。
「しかし、LLM は時々問題を起こすことがあります!」
はい、これはあなたにショックを与えるかもしれません。人間の同僚も同様です。たとえ素晴らしい同僚であっても、愚かなことをすることがあります。はい、読者の皆さん、私でもバグを 1 度か 2 度送ったことがあります。
LLM が完璧であるということではありません。完璧である必要はありません。それは LLM が役立つということです。そしてそれは本当にf-inです。
別の私が精神病に陥っていたであろう事柄

## Original Extract

How I use LLMs - five

How I use LLMs? - five
Mindless Rambling Nonsense
Paul D'Ambra
Fangler
Blog Posts
Week Notes
Kids games
Thu Jun 04 2026
How I use LLMs? - five
I have lost roughly 60% of my cynicism about the current high-water mark for LLMs and augmented coding… I am certain this last 9 months has been a significant generational change in how we can approach work.
Let's growth hack blog visitor numbers record how I use them today as a little reflection on where I think they work and where they don't. Something I can revisit as the tech (and my skill with it) improves
This is update number four, or entry number five, depending on how you want to count it.
part 1 - structured tone of fascination + scepticism
part 2 - learning how to manage and control them
part 3 - reluctant acceptance but also thankfulness (?) of what they do for you when you let go of control
part 4 - realising even though it's no longer "braincode" + no longer for humans, it still comes from human brains
someone at work named the episodes for me, i love it
so now… part 5 - yee haw, paul is AI pilled
in today's edition of AI hot takes
no longer focused on the code, 5% focus down from 95%
i don't use these tools, i hire them into my team
genuine frustration when they lack autonomy
stop making me tell you when to run
First, a necessary but brief diversion into how I use tools generally
I was largely self-taught as a developer and (after vbscript 😱) I started with DotNet. I found JetBrains Resharper early on and never looked back. The assistance in the IDE to write, refactor, and improve code was game-changing. I don't begrudge anybody any tool they want to use to level-up. Wanna use a graphical git client, or emacs, or anything else… go for it, go make cool things
Because of that entry to the industry I'm also not super keen on VSCode (you love it? great, see above, you do you, go make cool things). I'm too used to clever interventions helping me and VSCode is too barebones (and so much slower than SublimeText)
So, I like having tools directly involved in my workflow. I learned LINQ more quickly and more thoroughly because I had resharper prompting me inline, at write-time: "hey, why not like this?"
why write something down in a to-do list when i can open a new conversation and write exactly the same text as a prompt?! i immediately switch away and carry on with my day, like with a to-do list.
but my agent is connected to our code and to all of our analytics, so almost every time when I switch back it has context and challenge. or a PR
no longer focused on the code, 5% focus down from 95%
Recent improvements to agents and their UX has meant I only have to actively write code when I choose to. It would be possible to completely work via agents.
I can literally never write a line of code manually anymore and it is so ok.
Because the cost of understanding the code is basically zero now. The importance of the code being understandable is dropping.
It still has to be correct and fast but if I can't quickly grok it. I can just ask an LLM to explain it, or write some tests to exercise it. What might have taken hours in the past, I've worked on code that took me days to understand, will now take seconds.
Writing code is now far from the highest leverage thing I could do with my time
"Oh but Paul some code is very important!"
Yep, and cool, the literal code might still be the highest leverage thing you can do. But it isn't the highest leverage thing I could do any more. And I'd argue for many people that's true too.
And that is truly magical because now I can care about and achieve a bunch of stuff all at the same time and that would have been impossible for 2025 Paul.
i don't use these tools, i hire them into my team
If I am nailing two pieces of wood together it is up to me to figure out how to do it. To choose the hammer. To choose the nail. To pick the place and the angle that the nail goes in. To make sure I don't hit my thumb.
If I hire a carpenter then I make them a brew and tell them I'm here if they need me.
So, the mindset shift I'm suggesting is that I am hiring a junior engineer or product manager or whatever. And I am responsible for giving them context and guard rails. And then I let them do their work. I'm still responsible for if the work is valuable, but I'm not literally hitting the nail any more.
It changes so much about how I interact with the thing. And it lets me completely change how I work. I've moved into an executive role and I can still sling lots of code cos I don't hold the hammer, I hire the carpenter.
genuine frustration when they lack autonomy
We run a bunch of AI review tools on our PRs. And they have varying quality. They get better over time bit by bit. And they are earning their keep (just about - although they're pretty expensive tbh)
But over the last months since part 4 I've noticed a new expectation from me… they post their comment and I think "ugh, now i have to read and act on that… such bullshit"
By comparison, I have a "PR Shepherd" skill. I've had Claude do research about XP on the C2 wiki. I've had it interview me about software engineering to see what I care about in how work is done. And I've had it review the public history of my PR reviews to see what I care about in reviewing.
And when I run that skill it does a bunch of reviews in parallel, and anything that has a single solution, it just does. You can see an example here - I have it run and publicly comment on all of my PRs. It now feels rude to have a human review before I've had the robot find things. In the same way that for the longest time it's been rude to make a human lint your code when a tool could do it.
So, increasingly, the word machine can no longer earn its keep by only turning one set of words into another. It needs to turn that set of words into action. Just like if I hire a junior engineer and they asked me "should i fix this obvious bug" i would be frustrated.
stop making me tell you when to run
this has a side effect… i don't want to have to remember to ask an LLM for a whole bunch of stuff. same way i don't want to have to ask my colleagues about a bunch of stuff.
i want to say "hey, we should check the clink expander every two hours and route fanglers appropriately" and that just happens
i think the next great frontier is long-running and trigger-driven agents. and i am so excited for it
I mention "PR shepherd" above. It does all the mechanical work of getting a PR through review. It checks if a PR is open, opens it if not. marks it ready for review, runs a swarm of parallel reviewers and responds to them, checks external reviews and responds to them. It defers to me when it has to, otherwise it checks CI until it's green. It took me about an hour in total so far prompting to create it. And has saved me multiples of that.
I have "Paul pair". Every time the robot finishes work or is about to ask me a question it has an instruction to run the "Paul pair" skill. Which reminds it what I care about and to only interrupt me with important things. It's mecha paul. Saving meat paul from repeating the thing i was obviously going to repeat.
I have an executive coach (admittedly not a skill). It checks granola and slack and github and the org chart etc etc and once a week tells me if I'm working on what I've said is important. Tells me what I'm missing, what patterns it sees, what the top 3 things I should do to get back on track. It's not always right, but it doesn't have to be - it's on me to manage my work and attention regardless of what advice i receive.
If you have never done it then stop what you're doing. Open a new session with your favourite robot and type "look at our conversations and tell me what repeated interactions we could automate"
Then sit back and prepare to have your mind blown.
paul, you sound irredeemably AI pilled
yep, I am. sorry, it's happened. there's no going back. i've been sucked into the bubble. this is the future.
"but LLMs gets things wrong sometimes!"
yep, and this might shock you. so do human colleagues. even incredible colleagues can do silly things. yes, reader, even I have once or twice shipped a bug
it's not about the LLM being perfect. it doesn't have to be perfect. it's about the LLM being useful. And it really f-in is.
Things that a different me would have been tipped into psychosis by

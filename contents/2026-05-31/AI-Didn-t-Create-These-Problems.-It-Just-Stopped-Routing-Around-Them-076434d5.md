---
source: "https://baweaver.com/writing/2026/05/27/ai-didnt-create-these-problems/"
hn_url: "https://news.ycombinator.com/item?id=48338451"
title: "AI Didn't Create These Problems. It Just Stopped Routing Around Them"
article_title: "AI Didn't Create These Problems. It Just Stopped Routing Around Them. | baweaver"
author: "mooreds"
captured_at: "2026-05-31T00:27:28Z"
capture_tool: "hn-digest"
hn_id: 48338451
score: 4
comments: 0
posted_at: "2026-05-30T17:05:53Z"
tags:
  - hacker-news
  - translated
---

# AI Didn't Create These Problems. It Just Stopped Routing Around Them

- HN: [48338451](https://news.ycombinator.com/item?id=48338451)
- Source: [baweaver.com](https://baweaver.com/writing/2026/05/27/ai-didnt-create-these-problems/)
- Score: 4
- Comments: 0
- Posted: 2026-05-30T17:05:53Z

## Translation

タイトル: AI はこれらの問題を作成しませんでした。彼らを迂回するルーティングを停止しただけです
記事のタイトル: AI はこれらの問題を引き起こしませんでした。それは彼らを迂回するのをやめただけです。 |バーウィーバー
説明: AI に必要なすべてのガードレールは、すでに実行されているべき良い習慣です。 AI は、経験豊富な開発者が長年にわたって密かに回避してきたシステム上のギャップを明らかにします。

記事本文:
バーウィーバー
×
ホーム
執筆
話す
哲学
について
お問い合わせ
バーウィーバー
ソフトウェアエンジニアリング
ホーム
執筆
話す
哲学
について
お問い合わせ
シンプルであることは大変な作業です。しかし、そこには大きな見返りがあります。本当にシンプルなシステムを持っている人は、最小限の労力で最大の変化をもたらすことができるでしょう。
— リッチ・ヒッキー
ホーム / 執筆 / 投稿
あい
2026 年 5 月 27 日
AI がこれらの問題を引き起こしたわけではありません。それは彼らを迂回するのをやめただけです。
私はここ数年の大半を仕事で AI を多用して過ごしてきましたが、この間ずっと観察してきた非常に興味深いことが 1 つあります。
AI を軌道に乗せ、生産性を高め、予測可能性を高め、リスクを軽減するために、AI の前に設置しなければならなかったあらゆるガードレールはあるでしょうか?それらはすべて、私たちがずっとやるべきだったことです。
テスト、文書化、明確な所有権、最新の文書化、決定論的な検証?どれも新しいものではなく、特に物議を醸すものでもありませんが、AI にハンドルを握らせた瞬間に、どれだけのものが欠けていたかが正確にわかります。
AI が非常に優れている点があるとすれば、それは、望むものを達成するための最短経路を見つけることです。多くの場合、花壇や野菜の上を通り抜けます。
確かに、あなたにとっては明白ですが、AI にとってはどうでしょうか？全くない。
開発者として経験を積むにつれて、多くの問題が背景ノイズとなり、回避したり完全に無視したりするようになります。それはあまりにも自然に起こるので、私たちがそれをしているときは気づかないだけでなく、気が散るものや、その背後に解決すべき実際の問題があることに気づきます。
子モデルの順序を検証する 3 ～ 5 年前のコールバック フックを覚えていますか?一度に 30 件を超えるレコードを更新すると速度が低下し始めるものでしょうか?それは文書化されていませんでした、ね

ver は実際にコメントしましたが、どのリンターもそれをキャッチすることはできず、正直に言うと、誰がそれを書いたのかさえ覚えていません。もしかしたらあなただったのですか？私はそれを数回行いました。
重要なのは、本番環境で一度爆発すると、その時点から二度と同じことをしないように覚えておくことになるということです。運が良ければ、その教訓を自分で遭遇するのではなく、他の人から学べたかもしれませんが、いずれにせよ、それはあなたの頭の中にあり、今後それを見るたびにその教訓を回避することになります。
人間にとってはこれでうまくいきますが、AI にはそのような記憶力はありません。そのままフックに突入したり、新しいフックを作成したりしてデッドロックが発生し、午前 3 時に呼び出しを受けて、いったい何が起こったのかを把握しようとしているのです。
もしそれが人間だったら、私たちは彼らを責めないでしょう（そしてそうすべきです）、私たちの既存のシステムがどのようにしてそのような障害を起こして本番環境に到達することができなかったのかを尋ねるでしょう。私たちは彼らを愚か者だとか嘲笑するとか、そんなことは言いません。私たちは学び、適応し、前進します。願わくば、「二度としないでください」というガードレールも設置しておきたいと思っています。なぜなら、「二度としないでください」というのは、その人や周囲の人々が会社に留まる限り、永遠に続く希望であり祈りだからです。
それがそのパターンです。更新方法を知っているのは 3 人だけで、他の 3 人は 2 年前に去ったという古いドキュメント。 3 年前にリファクタリングされて存在しないシステムをテストするスタブでいっぱいのテスト スイート。どの 3P サービスが一定量を超える負荷を処理でき、どのサービスが処理できず、積極的なキャッシュを必要とするかについての暗黙の知識。経験豊富な人間だけがコードベースを操作しているときはこれらはすべて問題ありませんでしたが、組織的な記憶を持たない何かがコミットを開始した瞬間は賭けが外れ、すべてのギャップが活線になります。
Th

このすべてには、面白くて仕方ないある種の皮肉があります。数年前、Netflix は Chaos Monkey と呼ばれるこの斬新なコンセプトを発明しました。これは、サーバーを軌道上からランダムに核攻撃するツールです。何かが死ぬのと同じくらい、何が、どこで、いつ、あるいはなぜそうなるのかさえもわかりません。自己修復できるほどうまく設計されていると期待したほうがいいです。そうしないと、長い一日になるでしょう。
なぜこれを取り上げるのでしょうか？なぜなら、AI は私たちがこれまでに経験した中で最高のカオス エンジニアだからです。あなたが決して考慮しないギャップや穴を見つけます。それは、月が当たる夜の特定の時間帯に特定の角度で塗装が剥がれた建物の側面の亀裂をすり抜けて侵入し、関係者全員が正気を疑い、一体どうしてこんなことが起こったのか疑問を抱かせるほどだった。
AI は、コードをどのように記述するべきか、どのようなドキュメントが存在する必要があるか、そしてシステムが実際に私たちが思っているほど堅牢であるかどうか (実際にはそうではありません) についての私たちの仮定をストレステストしています。
重要なのは、それにどう対応するかを選択することです。私たちはシステムがどのように失敗したかを考えるシステム思考家のように行動するのでしょうか、それとも非難して非難するのでしょうか?生産が停止したとき、優れたエンジニアは、このシステムの何が原因でこの特定のカテゴリのエラーが発生したのかを尋ねます。私たちは、開発者がどのように注意を払っていなかったか、または彼らの仕事がどれほど悪かったかについて講義することはありません。同じ枠組みが AI にも当てはまります。AI は間違いを犯すものです。私たちはそれを知っていますが、では、それらの間違いを特定し、存続可能にするために、それを中心に何を構築すればよいでしょうか?
私たちのツールはこのために作られたものではありません
私たちが直面する問題の 1 つは、私たちのツールが概して AI を念頭に置いて構築されていないことです。これらは、人間が出力を読み、それにどのように応答するかを考えるという別の時間に向けて設計されました。
RSpec を例に挙げてみましょう

。 AI はテスト スイートを実行し、失敗を取得し、その情報から何が問題なのかを判断する必要があります。それはどのようにして行われるのでしょうか？スイートを実行し、出力を末尾/先頭/grep して圧縮し、必要な正しい出力を得るために、これらのコマンドの異なる呪文を使用してさらに 5 ～ 6 回実行します。最悪の場合、実行ごとに数分かかる場合もあります。
問題は RSpec やその記述方法ではありません。これは、出力をスキャンしてパターンを視覚的に照合し、どのような障害が関連しているかを判断できる人間向けに設計されています。 AI は、少なくとも効率的にそれを行うことができないため、情報を制限し、集中させようとします。必要なのは、構造化された出力です。失敗の場所、エラーの内容、エラー メッセージ、さらには一般的なパターンのいくつかのグループをリストする JSON です。キャッシュのようにそこから読み取り、エージェントに問題を解決して再集中させようとしている間に、複数のパスではなく 1 回のパスで問題を診断できます。
Rubocop も、Minitest も、人間が解釈できることを前提として人間が読めるテキストを出力するすべての開発ツールも同様です。これらは人間にとってはうまく機能しますが、何をすべきかを判断するのに何が失敗したかを正確に集中的に明確に把握する必要があるエージェント向けには構築されていません。
私はすべての AI ツールに JSON を出力するように指示し、それを再解析して新しい情報がないかどうかを確認し、構造化された出力を保証する明確なラッパーなしで生のコマンドを実行することを体系的に禁止し始めました。この 1 つの変更により、「もう一度試してください」と繰り返すのではなく、AI に実際に問題を診断させる効率が大幅に向上しました。完璧ですか？いいえ、しかしこれは確かな一歩であり、おそらく新しい考え方の基礎となるでしょう

これらのツールについて。
私はこれらのツールを批判しているわけではありません。まったく批判していません。彼らはその瞬間に合わせて作られており、自分の仕事に優れていますが、その瞬間は変わりました。この新しい世界では、RSpec、RuboCop、そして開発者ツールチェーン全体にとってエージェント モードがどのようなものになるかを考え始める必要があります。 AI がその情報を使って（できれば）正しいことを行うために必要なコンテキストを適切なタイミングでどのように提供するか。
When developing tooling with AI, especially for anything spanning more than 10 files, I like to use an 80/20 rule: 80% deterministic code, 20% AI glue for the finnicky details that would take hours to get right.
決定的な部分は、曖昧さなく合理的に検証、型チェック、lint、テストできるすべてのものです。コントラクト、スキーマ、静的型付け、機能フラグ、段階的ロールアウト。 All the boring infrastructure that makes systems substantially more predictable, and more tolerant to failure, and things will always fail some how some way.
Static Typing in Ruby : As an aside, the static typing debate in Ruby while valid does not hold well in the era of AI when any additional guardrail may be the layer that stops it from doing something particularly naughty.私は Sorbet からかなりの活用を得ており、AI が稼働する大規模なコードベースでの使用を確実に推奨します。
残りの 20% は、AI が本当に優れていると思う部分であり、厳密なルールよりも柔軟性と判断力が必要な部分です。 Summarization, routing, classification, the stuff which would be incredibly difficult to implement deterministically but where a probabilistic answer is close enough to provide tangible value.
この角度から見ると、AI は非常に特殊な障害モードを備えた別のツールとなり、ガードレールを設置することができ、段階的に

もっと学ぶにつれて、時間の経過とともにそれらを強化していきます。本当に厄介な失敗は、決定的なバックボーンを持たずに 100% AI にしようとしたときに起こります。そのとき、ユーザーデータの漏洩、トークンのコミット、銀行口座の空っぽ、またはその他の壊滅的な障害が発生します。
ここで、あといくつかのガードレールを記述するだけで、AI が適切な出力を生成することを 100% 保証できるとは言いません。そのような世界は存在しないので、存在するふりをすべきではありません。エンジニアが書き留めたり共有したりしなかったコンテキストを残したまま退職する場合、何かが最終的に下流でアップグレードされたために意味がなくなった、決して書き留められなかったアーキテクチャ上の決定、誰かが常にそれが何であるかを常に知っていて、何か問題が発生するとすぐに Slack で利用できるため、誰も書き留めなかったサービス間の暗黙の契約など、修正できないものがあります。
ということは、努力してはいけないということでしょうか？いいえ、全くその逆です。多くのことは修正可能であり、AI のおかげで、他の方法よりも早くこれらの問題を修正することが求められています。エージェントが必要とするために作成されるドキュメント、スタブが微妙に嘘をつき、時間の経過とともに真実から乖離するために意味のあるものとなるテスト、スティーブが LLM である場合、「スティーブに尋ねる」が機能しないために明確になる所有権。
しかし、ここが興味深い部分です。AI のために修正するすべてのギャップは、人間も同様に経験する可能性のあるギャップでもあるのです。何も知らない新卒、転勤してRubyを一行も書いたことのない先輩、締め切りに追われた請負業者。こうした問題は常に存在していましたが、AI によって無視することが大幅に困難になっただけです。
Square で API および SDK チームに新しいメンバーが加わったときに私たちがよく行っていたことの 1 つは、彼らに何か機知に富んだものを構築するよう依頼することでした。

その間、私たちは彼らに「なんてことだ！」と思ったことをすべて書き留めるように勧めました。当時「wtf doc」と呼ばれるものでした。すべての新入社員がこれを行った結果、新たなギャップ、思ったほど明確ではなかった新たな表面積が見つかり、反復できる可能性のあるいくつかの新しいアイデアが見つかりました。
AI はある意味、私たちの集合的な WTF 文書であり、これまで不可能だと考えられなかった速度と規模でシステムのギャップを調整することを私たちに強います。
失敗を個人的または道徳的な問題ではなく、組織的な問題として扱うときにチームが成功するのと同じように、AI システムも成功します。 「私たちのシステムに何が欠けていたので、このようなことが起こったのでしょうか？」と尋ねる人々。 「AI は悪いものだから使うべきではない」よりも、はるかに良い結果が得られます。
すべての問題に明確な答えがあるわけではありませんし、すべてのギャップを埋めることは決して不可能ですが、システムを正直に見るという習慣はどうでしょうか。それは単なる優れたエンジニアリングであり、AI であってもそれは変わりません。むしろ、障害に強い、または障害にすぐにロールバックできるような、運用上優れたシステムの構築に注力する必要性がこれまで以上に高まっています。
ニーズは常に存在していましたが、違うのは、AI によって無視できない問題になっているということです。
組織が報酬を得るのを支援します

[切り捨てられた]

## Original Extract

Every guardrail we need for AI is a good practice we should have already been doing. AI exposes the systemic gaps that experienced developers have been quietly routing around for years.

baweaver
×
Home
Writing
Speaking
Philosophy
About
Contact
baweaver
Software Engineering
Home
Writing
Speaking
Philosophy
About
Contact
Simplicity is hard work. But there's a huge payoff. The person who has a genuinely simpler system is going to be able to affect the greatest change with the least work.
— Rich Hickey
Home / Writing / Post
ai
May 27, 2026
AI Didn't Create These Problems. It Just Stopped Routing Around Them.
I’ve spent the better part of the last few years using AI heavily at work, and there’s one very interesting thing that I have observed in all of this time:
Every single guardrail I’ve had to put in front of AI to keep it on track, make it more productive, more predictable, and less risky? They’re all things we should have been doing all along.
Testing, documentation, clear ownership, up to date documentation, deterministic validations? None of them are new, none of them are particularly controversial, and yet the moment that you let AI take the wheel you find out exactly how much of it was missing.
If there’s one thing AI is exceptionally good at it is finding the shortest path to achieve what it wants, often times right through the flower bed and over the vegetables.
Sure, to you it’s obvious, but to AI? Not at all.
As we become more experienced as developers a lot of problems become background noise we route around or ignore entirely. It happens so naturally that we don’t even notice when we do it, as much as we observe that there’s something which is a distraction, and an actual problem to be solved behind it.
Remember that callback hook from 3-5 years ago that validates a child model’s ordering? The one that starts slowing down if you happen to update more than say 30 records at a time? It was never documented, never really commented, no linter was going to catch it, and quite honestly you don’t even remember who wrote it. Maybe it was you? I’ve done that one a few times.
Point being that thing blows up in production once and you’re going to remember from that point on not to do that again. If you’re lucky you learned that lesson from someone else instead of running into it yourself, but either way it’s in your head and now you route around it every time you see it in the future.
That works great for humans, but AI doesn’t have that type of memory. It will walk straight into that hook or even write a new one, trigger a deadlock, and now you’re getting paged at 3AM trying to figure out what in the world went on.
If it were a human we would not blame them, we would (and should) ask how our existing systems failed them to allow an outage like that to make it to production. We don’t call them stupid or mock them or any such thing. We learn, we adapt, and we move on. Hopefully we also make sure there’s a “never again” guardrail in place as well, because “don’t do that again” is a hope and a prayer with a shelf life as long as that person and the people around stay at the company.
That’s the pattern. The stale documentation that only three people know how to update, and the other three left two years ago. The test suite full of stubs that tests a non-existent system that was refactored away three years ago. The implicit knowledge about which 3P services can and cannot handle above a certain amount of load and require aggressive caching. All of these were fine when experienced humans were the only ones navigating the codebase, but the moment that something without that institutional memory starts committing bets are off and every gap is now a live wire.
There’s a certain irony in all of this that I can’t help but find amusing. Years ago Netflix invented this novel concept called the Chaos Monkey , a tool that would randomly nuke a server from orbit. It wouldn’t tell you what, where, when, or even why as much as something would die and you had better hope that you designed things well enough to self-heal or it’s going to be a long day.
Why bring this up? Because AI is the best chaos engineer we’ve ever had. It finds gaps and holes that you would have never considers. It squeezes through cracks in the side of a building where the paint chipped at a certain angle at a particular hour of night when the moon hits it in a way that leaves everyone involved doubting their sanity and asking how in the world this happend.
AI is stress-testing our assumptions about how code should be written, what documentation needs to exist, and whether our systems are actually as robust as we think that are (they’re not.)
The important part is how we choose to respond to that. Do we behave like systems thinkers who consider how systems failed, or do we point fingers and blame? When production blows up a good engineer will ask what failed in this system that allowed this particular category of error to happen. We do not lecture about how a developer must not have been paying attention or how bad they are at their jobs. The same frame applies to AI: It’s going to make mistakes. We know it, so what are we building around it to make those mistakes identifiable and survivable?
Our Tools Weren’t Built for This
One problem we run up against is that our tools were not built with AI in mind, by and large. They were designed for a different time with a human reading the output and figuring out how to respond to it.
Take RSpec, for instance. AI runs your test suite, gets failures, and needs to figure out what is wrong from that information. How does it do that? It runs the suite, tails/heads/greps the output to condense it, and then proceeds to run it 5-6 more times with different incantations of those commands to get the correct output it wants, with every run taking multiple minutes in some of the worst cases.
The problem isn’t RSpec or how it was written. It was designed for humans who scan output, pattern match visually, and can tell what failures might be related. AI can’t do that, at least not efficiently, so it tries to limit and focus information. What it needs is structured output: JSON that lists failures, where they are, what the error was, the error message, maybe even some groupings of common patterns. It can read from that like a cache and diagnose issues in one pass instead of several while you’re trying to get the agent to knock it off and re-focus.
Same with Rubocop, same with Minitest, same with pretty well every developer tool that outputs human readable text with the expectation that a human is on the other end to interpret it. They work great for humans, but they’re not built for an agent that needs a focused and clear view of exactly what failed to help it to figure out what to do about it.
I’ve started telling all of my AI tools to output JSON, and then re-parse that for any new information, and then started systematically disallowing raw commands from being run without clear wrappers that guarantee that structured output. That single change has made me substantially more effective at getting AI to actually diagnose issues instead of repeatedly saying “let me try again.” Is it perfect? No, but it’s a definite step, and perhaps the foundation for a new way of thinking about these tools.
Now I’m not criticizing these tools, not at all. They were build for their moment, and they’re great at what they do, but that moment has shifted. In this new world we need to start asking what agentic modes might look like for RSpec, for RuboCop, and for our entire developer toolchain. How do we give AI the context it needs, at the right time, to (hopefully) do the right thing with that information.
When developing tooling with AI, especially for anything spanning more than 10 files, I like to use an 80/20 rule: 80% deterministic code, 20% AI glue for the finnicky details that would take hours to get right.
The deterministic part is everything we can reasonably validate, type check, lint, and test without ambiguity. Contracts, schemas, static typing, feature flags, staged rollouts. All the boring infrastructure that makes systems substantially more predictable, and more tolerant to failure, and things will always fail some how some way.
Static Typing in Ruby : As an aside, the static typing debate in Ruby while valid does not hold well in the era of AI when any additional guardrail may be the layer that stops it from doing something particularly naughty. I’ve gotten substantial leverage out of Sorbet, and would certainly encourage its use in any large codebase which has AI running about.
The remaining 20% is where I find that AI really shines, the parts where you need flexibility and judgement rather than firm rules. Summarization, routing, classification, the stuff which would be incredibly difficult to implement deterministically but where a probabilistic answer is close enough to provide tangible value.
When viewed from this angle AI becomes another tool with a very specific failure mode that we can put guardrails around, and progressively enhance them over time as we learn more. The really nasty failures happen when people try and make it 100% AI with no deterministic backbone. That’s when you see user data leaking, tokens committed, bank accounts emptied, or other catastrophic failures.
Now I’m not going to tell you that you can 100% guarantee AI will produce good output if you only write just a few more guardrails. That world does not exist, and we should not pretend it does. There are things it can’t fix for like when engineers leave with context they never wrote down or shared, architectural decisions which were never written down which are no longer relevant because something finally got upgraded downstream, implicit contracts between services that no one ever wrote down because someone always knew what they were and was available on Slack almost immediately whenever something went wrong.
Does that mean that we should not try? No, quite the contrary. A lot can be fixed, and AI is forcing us to fix these things faster than we would have otherwise. Documentation that gets written because agents need it, tests that are made meaningful because stubs subtly lie and diverge from the truth over time, ownership that gets clarified because “ask Steve” does not work when Steve is an LLM.
But here’s the fascinating part: Every gap you fix for AI is also a gap that a human may very well have fallen through as well. That new grad who doesn’t know any better, the senior who transfers and has never so much as written a single line of Ruby, the contractor that has a deadline to meet. These problems were always here, AI just made it substantially harder to ignore them.
One thing we used to do at Square when we had a new team member on the API and SDK team was that we’d ask them to build something with our product, and during that time we encouraged them to write down anything and everything that made them go “WTF!” in what was then dubbed a “wtf doc.” The results after every new hire did this is we’d find new gaps, new surface areas that were nowhere near as clear as we had thought, and potentially several new ideas we could iterate from.
AI is, in a way, our collective WTF doc that forces us to reconcile with the gaps in our systems at a speed and scale that we never thought possible.
Just as teams thrive when they treat failures as systemic issues rather than personal or moral ones, so too will AI systems thrive. Those that ask “what was missing from our system that allowed this to happen?” rather than “AI is bad and we should not use it” will get far better results from it.
Not every problem is going to have a nice clean answer, and we certainly can never close every gap, but the practice of looking at systems honestly? That’s just good engineering, and that does not change with AI. If anything, the need is now greater than ever to focus on creating operationally excellent systems that are resilient to failure, or failing that really danged quick to roll back from it.
The need was always there, the difference is that AI is making it a problem we cannot ignore.
I help organizations turn comp

[truncated]

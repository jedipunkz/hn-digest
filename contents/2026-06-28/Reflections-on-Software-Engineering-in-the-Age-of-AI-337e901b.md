---
source: "https://adiamond.me/2026/06/software-engineering-in-the-age-of-ai/"
hn_url: "https://news.ycombinator.com/item?id=48708721"
title: "Reflections on Software Engineering in the Age of AI"
article_title: "Software Engineering in the Age of AI"
author: "diamondap"
captured_at: "2026-06-28T16:25:24Z"
capture_tool: "hn-digest"
hn_id: 48708721
score: 1
comments: 0
posted_at: "2026-06-28T16:17:38Z"
tags:
  - hacker-news
  - translated
---

# Reflections on Software Engineering in the Age of AI

- HN: [48708721](https://news.ycombinator.com/item?id=48708721)
- Source: [adiamond.me](https://adiamond.me/2026/06/software-engineering-in-the-age-of-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T16:17:38Z

## Translation

タイトル: AI 時代のソフトウェア エンジニアリングについての考察
記事のタイトル: AI 時代のソフトウェア エンジニアリング
説明: 知らない人のために言っておきますが、私は小説を書いていないときは、ソフトウェア エンジニアとしてコードを書いて日々を過ごしています。最近のソフトウェア業界は人工知能に大きく依存しています。それは、公的にアクセス可能な何兆行ものソースコードを研究しており、コードが問題を解決するためです。
[切り捨てられた]

記事本文:
AI時代のソフトウェアエンジニアリング
アンドリュー・ダイアモンド
AI時代のソフトウェアエンジニアリング
知らない人のために言っておきますが、私は小説を書いていないときは、ソフトウェア エンジニアとしてコードを書いて日々を過ごしています。最近のソフトウェア業界は人工知能に大きく依存しています。 AI は、公的にアクセス可能な数兆行のソース コードを研究し、コードはテスト可能な正解と不正解の解決策で問題を解決し、コードは特にコンピューターが理解できるように構造化されているため、コードを書くのが非常に得意になりました。
プログラマーが AI を使い始める前、典型的なワークフローは次のようになっていました。
誰かがあなたに、既存のプログラムに機能を追加するように依頼しました。
その機能の正式な定義を作成し、その機能が何をすべきか (そしてすべきでないか)、ユーザーがその機能にどのようにアクセスできるか、そして機能が正しく動作していることをテストする方法を説明します。
この機能を実装するのに最適なデータ構造、アルゴリズム、コード ライブラリ、および外部サービスを調査することに時間を費やします。
新しい機能を構築するためのコード、それが期待どおりに動作することを確認するためのテスト、そしてユーザーにその使用方法を説明し、他のエンジニアにその機能を保守およびデバッグするために知っておくべきことを伝えるドキュメントを作成します。
「プル リクエスト」を作成し、組織内の他のエンジニアに新しいコードのレビューとコメントを依頼し、最終的に製品での使用を承認するように依頼します。
AI が一貫して非常に優れたコードを生成できるようになったので、ソフトウェア開発者のワークフローは次のようになります。
AI に新しい機能を作成するよう求めるプロンプトを作成します。
AI が書いた内容を確認して、必要に応じて変更を加えたり、AI エージェントに変更を依頼したりできます。
新しいコードを自分で既存のコードベースにマージするか、他の人がレビューできるようにプル リクエストを作成します。

そしてそれをマージします。
以前のワークフローでは、創造的なプロセスは主に頭の中で行われていました。新しいプロセスでは、AI の内部陰謀の中で展開される創造的なプロセスを監督します。 AI が作業を開始できるように、簡潔で思慮深く正確なプロンプトを作成することにある程度の努力を払ってきましたが、コードを自分で書くときに通常行うような難しい思考はまったく行っていません。 AI からコードを受け取るとき、あなたは基本的に編集者として行動することになります。AI はコードを書くことはできますが、AI はいつもあなたほどプロジェクトの全体像を把握できるわけではないため、この新しいコードが問題を引き起こさないことを確認する必要があります。
AI は、追加したコードが製品が適用される法的要件に違反しているかどうかを知りません。外部システムに対するリクエストが完了するまでに 10 ミリ秒かかるのか、10 分かかるのかはわかりません。コードの動作が、チームメイトが 3 週間後に追加する予定の新機能と競合するかどうかはわかりません。先月作成した、機密情報を処理する別の関数と対話するときに、作成したばかりの関数が新たなセキュリティ問題を引き起こす可能性があるかどうかはわかりません。
上級開発者はこれらのことを知っています。だからこそ、「正常に機能する」ように見える AI コードを精査し、頻繁に修正する必要があります。上級開発者にとって、AI は有能で仕事の早いジュニアまたはミドルレベルの開発者であり、適切に指示されれば、ほとんど堅実な成果物を生成しますが、組織的な知識や、過去 20 年間に培ってきた深く幅広いシステム レベルの知識が欠けています。
さて、少し戻って類推してみましょう。あなたが歴史小説の作家だとしましょう。ワークフローはどのようなものですか?プロバ

次のようなことをしてください:
あなたは、1760 年のロンドンのセント ポール大聖堂の外で 2 人の政治家が議論している場面を思い浮かべます。この場面を正確に書くためには、服装、街路の雰囲気、政治情勢など、何を知っておく必要があるかを考えます。
あなたは何冊かの本を開き、特に次のことについてメモを取り始めます。
彼らの社会経済的地位と社会における役割に基づいて、あなたのキャラクターは何を着ますか?
他に彼らと一緒に路上にいるのは誰ですか？ベンダー？馬車のタクシー運転手？それらはどのように見えましたか？その時間に煙突掃除をしていましたか？売春婦や法務官についてはどうですか？
あなたの登場人物たちが議論する主な政治的人物は誰ですか?また、彼らは現在どのような立場にありますか?
ここ数週間または数か月のどの歴史的出来事が関連していますか?また、それらは登場人物の議論にどのような影響を及ぼしますか?
あなたは自分の執筆に戻り、想像力から紡ぎ出した情景に歴史的事実を織り込みます。
この小説家とソフトウェア開発者には多くの共通点があります。実際には、ソフトウェア開発者が何年も前の仕事から、これから書こうとしている新機能にどのデータ構造とアルゴリズムが適しているか、どのタイプのキャッシュとデータベースが適しているかをすでに知っているのと同じように、小説家は書き始める前に歴史調査の多くを行っているでしょう。
両方の作業者に共通しているのは、自分たちが作成している素材に対する深い関与感です。彼らは時間を忘れてしまうほど仕事に没頭しています。小説執筆でもソフトウェア開発でも、私や多くの同僚にとって、時計を確認して問題に取り組み、10分後に時計を見ると4時間が経過していることがよくあります。
ティ

s は、心理学者のミハイ・チクセントミハイが 1990 年のベストセラー本『フロー』で説明した「最適な経験」の状態です。作家、ソフトウェア開発者、画家、ミュージシャン、その他すべてのクリエイティブなタイプは、時間と状況が許せばこの状態に入ります。多くのソフトウェア エンジニアが時間外に自宅で働いているのは、まさに通常の営業時間外に会議やその他の中断がないため、この状態に入ることができるからです。
さて、歴史小説家をソフトウェア開発者の立場に置いてみましょう。彼女の出版社から電話があり、2 年に 1 冊ではなく、毎年 4 冊の本を市場に出す方法を見つけたとの連絡を受けました。彼らは、1日5ページの有能な文章を激安で作成できる一流の高校生や大学生を大量に採用した。出版社は、歴史小説が原作者の卓越性のレベルを維持するか、少なくともそれに近いレベルを維持することを望んでおり、編集者として彼女の仕事を継続しています。
小説家の仕事は、生徒たちの作品を編集することだ。生徒たちはそれぞれ、ページを書くように注意深く指示されており、少しの作業で、一貫した章につなぎ合わせられるはずだ。
高校生や大学生の成績を採点したことのある人なら誰でも、これが一般にやりがいのある仕事ではないことを知っています。 1週間に100枚の論文を採点しなければならなかったことがあれば、それがどれほど大変な作業であるかわかるでしょう。
小説家は、ソフトウェアエンジニアと同様に、もはや自分の仕事に深く関わっていません。編集は創造ではありません。あなたは自分の想像力に身を委ねることはありません。発明のプロセスに自分の心や感情を没頭させることはありません。代わりに、問題を根絶し、不器用な言葉遣いや冗長な説明を整理しようとしているのです。フロー状態はなくなりました。あなたは今、より大きなプロセスの歯車です

それはあなたの創造性やそれを発揮する必要性をあまり評価しません。
さらに悪いことに、AI が生成したコードを何ヶ月もレビューして個人的に感じたことですが、スキルが急激に低下します。実装すべき機能や修正が難しいバグなど、新しい問題が発生したときに、それに数時間を費やすという考えは侮辱的なものに感じられます。クロードが 5 分以内にバグを見つけて修正の草案を作成できるのに、なぜ私がそのすべてのコードを調べなければならないのでしょうか?
本当になぜでしょうか？ボットがそれをやってくれるのに、なぜあなたや私が、以前は本当は好きでやっていたもの、今では面倒に感じていることに精神的なエネルギーを費やす必要があるのでしょうか?
何ヶ月も AI を使って仕事をしていると、少なくともコーディングに関しては、私は明らかに怠け者になり、愚かになりました。 (私は書くときに AI を使用しません。書くこと自体が自分の考えを整理し、明確にする行為だからです。ボットに書いてもらうのは、他の人にお金を払って運動してもらって、それで体調が整えられることを期待するようなものです。)
また、仕事中、特に明らかに AI によって書かれたメールを同僚から受け取ると、よりイライラするようになりました。メールでは、同じく AI によって書かれた長い文書をレビューするよう依頼されることがよくあります。そのため、私はいつも「これを書くのが面倒なら、なぜわざわざ読む必要があるのでしょう？」と考えてしまいます。
AI を使用して基本的な質問に答えます。 Amazon のマネージド データベースが Microsoft や Digital Ocean のサービスとどのように比較されるか知りたいですか?クロードのチャットボットは、概要を把握するのに最適です。そこからソースに移動して詳細を確認できます。
しかし、クリエイティブな人々が、最も想像力豊かなフロー状態の思考をボットの軍隊に引き渡すことを選択するのは、長期的には間違いになるだろうと私は思います。これらの AI ボットは、コード、ホワイト ペーパー、詩、小説など、すべての知識を私たちから得ています。

ニュース記事、私たちの伝記、そしてスタック オーバーフローのすべての答えは、何千人もの人々が何十年もかけて苦労して獲得した知識に基づいて書いたものです。
ボットは、以前は無料だったデータをすべて食い尽くしてしまい、今ではそれを私たちに売り返しています。人々が Claude と ChatGPT に質問を寄せるため、Stack Overflow ではもう誰も質問に答えません。 Stack Overflow のような無料で使用できるサービスに蓄積されていた公開された知識は枯渇しつつあり、AI ボットが将来の技術的な質問に対する答えを得ることが少なくなってきています。
AI ボットを管理する熟練した上級開発者がいれば、AI ボットの方が安価であるため、ソフトウェア会社は若手開発者を全員解雇しました。しかし、パイプラインに若手開発者がいないため、本当に複雑で幅広く深い問題、つまり範囲が広すぎて AI が処理できない問題を解決する方法を苦労して学んでいるので、明日の上級開発者はどこから来るのでしょうか? 5 年後には誰がボットを管理する資格を持つでしょうか?
私が話している複雑で幅広く深い問題は、回答すべき質問の全範囲を記述する AI の「コンテキスト ウィンドウ」に収まりきらないほど大きすぎます。 100 か国の製品を管理するすべての法律、製品が相互作用する必要があるすべての関連コードベースのあらゆるニュアンス、製品を構成する数百万行のコードに埋め込まれたすべてのビジネス ルールを完全に理解することは、現時点では不可能ですし、今後も不可能になる可能性があります。開発者とマネージャーのチームは、この法的および制度的知識を理解し、適用することができます。当然手間もかかりますし、費用もかかります。複雑さは常に存在します。
しかし、この作業をすべてボットに任せるとどうなるでしょうか?ボットが作成した作業が正しいことを検証する知識を持っている人はどこにいるでしょうか?
数年前、

私は、もう見つからなくなったが[1]、米海軍がすぐには必要のない空母の建造に資金を提供するよう米海軍がどのように議会を説得したかについての記事を読んだ。 「なぜ私たちがその費用を支払わなければならないのですか？」上院議員と下院議員に尋ねた。
海軍は「空母の建造に必要な技術は何十年もかけて苦労して獲得したものなので、今建造しなければ10年後には作り方を忘れてしまうだろう」と答えた。
海軍は、他の国が匹敵することのできない世界クラスの船を建造した古参の技術者や職人たちに、地球上の誰も建造できないものを建造するために必要なスキルを若い世代に伝えてほしいと考えていました。それを実現する唯一の方法は、実際に船の設計と組み立てのすべての段階をチームに実際に実行させることでした。
議会がこのプロジェクトに資金を提供したのは、その非効率性、党派間の争い、全体的な非効率性についての私たちの苦情にもかかわらず、議会がほとんどの企業よりも先見の明を持っていたからです。彼らは次の収益報告書ではなく、数十年先の未来を見据えていました。
最大手企業から小規模な新興企業まで、ソフトウェア業界の企業は、短期的な見返りが非常に大きいため、クリエイティブな作業をボットに任せています。 1 つの製品を完成させるために、なぜ 10 人の開発者にお金を払うのか

[切り捨てられた]

## Original Extract

For those of you who don’t know, when I’m not writing novels, I spend my days as a software engineer, writing code. The software industry these days relies heavily on artificial intelligence. Because it has studied trillions of lines of publicly accessible source code, because code solves problems w
[truncated]

Software Engineering in the Age of AI
Andrew Diamond
Software Engineering in the Age of AI
For those of you who don’t know, when I’m not writing novels, I spend my days as a software engineer, writing code. The software industry these days relies heavily on artificial intelligence. Because it has studied trillions of lines of publicly accessible source code, because code solves problems with testable right and wrong solutions, and because code is structured specifically to be understood by computers, AI has gotten very good at writing code.
Before programmers started using AI, a typical workflow looked like this:
Someone asks you to add a feature to an existing program.
You write up a formal definition of that feature describing what it should (and should not) do, how users can access it and how to test that it’s working correctly.
You spend time researching which data structures, algorithms, code libraries and external services might serve best to implement this feature.
You write the code to build the new feature, the tests to make sure it works as expected, and the documentation telling users how to use it and telling other engineers what they’ll need to know to maintain and debug it.
You create a “pull request,” asking other engineers in your organization to review and comment on your new code, and ultimately to approve it for use in the product.
Now that AI can consistently produce pretty good code, the software developer’s workflow looks like this:
You write a prompt asking AI to create the new feature.
You review what the AI wrote, making changes as you see fit or asking the AI agent to make those changes for you.
You either merge the new code into the existing codebase yourself, or you create a pull request for someone else to review and merge it.
In the old workflow, the creative process happened mostly in your mind. In the new process, you supervise the creative process that unfolds inside the AI’s internal machinations. You’ve put some effort into creating a concise, thoughtful, accurate prompt to get the AI started on its work, but you haven’t done any of the hard thinking you would normally do in writing the code yourself. When you get code back from the AI, you’re essentially acting as an editor because, while AI can write code, it cannot always see the big picture of your project the way you can, and you need to make sure this new code won’t cause problems.
AI does not know whether the code it just added violates some legal requirement to which your product is subject. It does not know if the request it makes to an external system is going to take ten milliseconds or ten minutes to fulfill. It does not know whether the actions of its code will conflict with some new feature that you know your teammates will be adding three weeks from now. It does not know whether the function it just wrote might introduce a new security problem when interacting with another function you wrote last month that handles sensitive information.
A senior developer does know these things, and this is why he or she needs to vet and often correct the AI code that appears to “just work.” To a senior developer, AI is a competent, fast-working junior or mid-level developer that, when properly directed, produces mostly solid work but lacks the institutional knowledge and the deep and broad systems-level knowledge that you have developed over the past twenty years.
Now, let’s back up for a second and make an analogy. Let’s say you’re a writer of historical novels. What does your workflow look like? Probably something like this:
You picture a scene in which two statesmen are arguing outside St. Paul’s Cathedral in London in 1760. You consider what you need to know to write this scene accurately, including clothing, the atmosphere on the street and the political situation.
You crack open a number of books and start taking notes on the following, among other things:
Based on their socioeconomic positions and their roles in society, what would your characters be wearing?
Who else is on the street with them? Vendors? Cabbies in horse-drawn carriages? What did those look like? Were the chimney sweeps out at that time of day? What about prostitutes and law officers?
Who are the primary political figures about whom your characters argue, and what are their positions at the moment?
What historical events from recent weeks or months are relevant, and how will they influence your characters’ argument?
You go back to your writing, weaving the historical facts into a scene that you spin from imagination.
The novelist and the software developer have a lot in common here. In reality, the novelist will have done much of her historical research before she begins writing, just as the software developer already knows from years of prior work which data structures and algorithms and which types of caches and databases will be suitable for the new feature he’s about to write.
What both workers have in common is a deep feeling of engagement with the material they’re creating. They are entirely immersed in their work, to the point where they often lose track of time. In both novel writing and software development, it’s common for me and many of my peers to check the clock, dive into a problem, and then look at the clock ten minutes later to discover that four hours have passed.
This is the state of “optimal experience” that psychologist Mihaly Csikszentmihalyi described in his bestselling 1990 book, Flow . Writers, software developers, painters, musicians and all other creative types enter this state when time and circumstances permit. Many software engineers work after hours at home precisely because the lack of meetings and other interruptions outside normal business hours permit them to enter this state.
Now, let’s put the historical novelist in the position of the software developer. She gets a call from her publisher saying they’ve found a way for her to bring four books to market each year instead of one book every two years. They’ve recruited a bunch of top-notch high school and college students who can each crank out five pages a day of competent writing for dirt cheap. The publisher wants the historical novels to maintain the original writer’s level of excellence, or to at least be close, so they’re retaining her services as an editor.
The novelist’s job is now to edit the work of the students, each of whom has been carefully prompted to write pages that should, with a little work, be stitched together into coherent chapters.
Anyone who has ever graded the work of high school and college kids knows that this is generally not rewarding work. If you’ve ever had to grade a hundred papers in a week, you know what a grind that is.
The novelist, like the software engineer, is no longer deeply engaged in her work. Editing is not creating. You do not give yourself over to your imagination. You do not immerse your mind and feelings in the process of invention. Instead, you’re rooting out problems, trying to clean up clumsy wording and redundant descriptions instead. The flow state is gone. You are now a cog in a larger process that doesn’t really value your creativity or your need to exercise it.
Worse still–and I have felt this personally after months of reviewing AI-generated code–your skills drop off sharply. When a new issue arises–a feature to be implemented, or a tricky bug to fix–the idea of wasting several hours on it feels insulting. Why should I dig through all that code when Claude can locate the bug in five minutes and start drafting a fix?
Why indeed? Why should you or I invest mental energy in something we used to actually like doing, something that now feels like a chore, when the bot can do it for us?
Months of working with AI have made me noticeably lazier and stupider, at least when it comes to coding. (I don’t use AI when I write, for the very reason that writing itself is the act of organizing and clarifying my own thought. Letting the bot write for me would be like paying someone else to exercise for me and hoping that gets me in shape.)
I’ve also become more impatient at work, particularly when I get emails from coworkers that are clearly AI-written. The emails often ask me to review longer documents that were also written by AI. That always leaves me thinking, “If you couldn’t be bothered to write this, why should I bother reading it?”
I do use AI to answer basic questions. Want to know how Amazon’s managed databases compare to the offerings from Microsoft and Digital Ocean? Claude’s chatbot is a good place to get a high-level overview. From there, you can go to the sources to get more detail.
But I think that creative people choosing to hand over their most imaginative, flow-state thinking to an army of bots will be a mistake in the long run. Those AI bots got all their knowledge from us–from our code, our white papers, our poems and novels and news stories, our biographies and all the Stack Overflow answers that thousands of people spent decades writing from hard-won knowledge.
The bots ate up all the data, which used to be free, and now they sell it back to us. No one answers questions on Stack Overflow anymore because people take their questions to Claude and ChatGPT. The publicly available knowledge that used to accumulate in free-to-use services like Stack Overflow is drying up, leaving the AI bots less to feed on for answers to future technical questions.
The software companies fired all their junior developers because the AI bots are cheaper, as long as you have skilled senior developers to manage them. But with no junior developers in the pipeline, learning the hard way how to solve the truly complex, broad and deep problems–problems whose scope is simply too large for AI to handle–where will tomorrow’s senior developers come from? Who will be qualified to manage the bots in five years?
The complex, broad and deep problems I speak of are too big to fit into AI’s “context windows” which describe the full scope of questions to be answered. You cannot now and may never be able to feed into AI a full understanding of every law that governs your product across a hundred different countries, every nuance of every related codebase with which your product must interact, every business rule embedded in the millions of lines of code that make up your product. A team of developers and managers can understand and apply this legal and institutional knowledge. It takes work, naturally, and is expensive. Complexity always is.
But what happens when we resign all this work to the bots? What person anywhere will have the knowledge to verify that the work the bots are producing is correct?
A few years ago, I read an article, which I can no longer locate [1], about how the US Navy convinced Congress to fund the construction of an aircraft carrier that the Navy didn’t immediately need. “Why should we pay for that?” asked the senators and representatives.
The Navy replied, “Because the skills required to build an aircraft carrier were so hard won over so many decades that if we don’t build one now, we’ll forget how to do it in ten years.”
The Navy wanted the old hands, the engineers and craftsmen who had built the world-class ships no other country could match, to pass down to the younger generation the skills they would need to build what no one else on earth could build. The only way to do that was by actually doing it, to actually put the team through every step of designing and assembling the ship.
Congress funded the project because, despite all our complaints about their inefficiency, partisan bickering, and general ineffectiveness, they had more foresight than most of our corporations. Instead of looking at the next earnings report, they were looking decades into the future.
Companies in the software industry, from the largest corporations to the smallest start-ups, are off-loading creative work to bots because the short-term payoff is immense. Why pay ten developers to crank out one product

[truncated]

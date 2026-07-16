---
source: "https://www.aha.io/engineering/articles/staying-familiar-with-the-code-when-its-written-by-an-llm"
hn_url: "https://news.ycombinator.com/item?id=48938749"
title: "How do you stay familiar with the code when it's written by an LLM?"
article_title: "How do you stay familiar with the code when it's written by an LLM?"
author: "justinweiss"
captured_at: "2026-07-16T19:58:00Z"
capture_tool: "hn-digest"
hn_id: 48938749
score: 6
comments: 0
posted_at: "2026-07-16T18:59:33Z"
tags:
  - hacker-news
  - translated
---

# How do you stay familiar with the code when it's written by an LLM?

- HN: [48938749](https://news.ycombinator.com/item?id=48938749)
- Source: [www.aha.io](https://www.aha.io/engineering/articles/staying-familiar-with-the-code-when-its-written-by-an-llm)
- Score: 6
- Comments: 0
- Posted: 2026-07-16T18:59:33Z

## Translation

タイトル: LLM によって書かれたコードに慣れるにはどうすればよいですか?
説明: LLM を使用して 1 か月前に出荷したコードを作成した場合、それをどの程度覚えていますか?今日コミットしたコードがそれを壊すかどうかわかりますか?顧客がバグを報告した場合、あなたはすぐに同じように「ああ、それが何なのか知っている」と思いますか

記事本文:
LLM によって書かれたコードをどのように理解すればよいでしょうか?エンジニアリング記事
ああ！ Develop は、戦略を配信に結び付けるアジャイル ツールです。
LLM によって書かれたコードをどのように理解すればよいでしょうか?
1 か月前に出荷したコードを作成するために LLM を使用した場合、それをどの程度覚えていますか?今日コミットしたコードがそれを壊すかどうかわかりますか?顧客がバグを報告した場合、以前と同じようにすぐに「ああ、それが何なのか知っている」と思いますか?
LLM にさらに多くのコードを作成してもらうと、自分が出荷するコードについてあまり精通していないことがすぐにわかります。自分が下した決断を思い出すのが難しくなります。コードがどこに存在するのか本能的にわかりません。誰かがそれについて質問した場合、頭から答えるのではなく、LLM に戻る必要があります。時間が経つにつれて、より多くの作業とより多くの理解を LLM に委任する必要があります。コードの知識が失われるにつれて、LLM への指導はさらに悪化します。そして、LLM が次の機能を書くのに苦労しているとき、どうやって解決するかはおろか、その原因さえわかりません。
これが今のソフトウェア開発の姿なのでしょうか？理解できないコードに対してどうやって責任を負えるのでしょうか?
たぶんあなたは気にしないでしょう！おそらく来月の LLM は、今月のモデルによって引き起こされた問題を解決するだけでしょう。もしかしたらあなたは完全に結果を重視しており、問題を解決する限りコードは重要ではないかもしれません。おそらくこれは、もし存在していればパッケージをダウンロードしていたであろうものを置き換えるだけの 1 回限りのプロジェクトかもしれません。そのコードを見ることもなかったでしょう?
しかし、これがあなたやあなたの会社が依存しているアプリであり、時間をかけて作成しているものである場合、それは重要です。人々はあなたが生み出すものに依存しており、あなたが今日行っている仕事は無限に存続しなければなりません。
hを気にしなければなりません

ああ、それはただのことではありません。
コードを手動で記述し、問題を作成して修正し、この新しいコードをどこに配置するかを決定し、古いコードをリファクタリングして改善する - これらすべてがコードの直感的な理解を構築します。誰かがバグを見つけて、あなたが即座に答えを推測するとき、それはその直感に基づいています。これにより、こちらで値を変更すると、向こうで移行が必要になることがわかります。コードのさまざまな領域に重点を置くチームが存在するのはそのためです。そして、注意しないと、コードを生成する LLM によって、手遅れになるまで直感的な理解を築くことができなくなります。
一方、AI は、トレーニング データ、コードベース内で認識されるコード、およびガイダンスという 3 つの信号を使用してコードの記述方法を決定します。 LLM をガイドする能力を失い、LLM が作成したコードの方が多く表示され、ユーザーが作成したコードの内容が少なくなると、プログラミング プロセスや哲学が何であるかは問題ではなくなります。それはコードベースを平均値まで引きずっています。
では、コードを書く量を減らしても、コードに慣れ続けるにはどうすればよいでしょうか?
テストで恥ずかしいほど間違った答えをしたのに、その正解を今でも覚えているという経験はありませんか?プログラマーとして、その失敗して驚くべき感情が、多くの不明瞭な情報が私たちに残る理由です。 LLM がコードを生成しているときは、そのようなことは感じられません。プロセスがスムーズすぎます。
それは、最初に質問について考えずに、本の後ろにある答えを読むようなものです。
代わりに、その驚きの感情を育む必要があります。間違いを発見したり、間違いを生み出したりする機会を探してください。 LLM に、大胆な最適化を試し、それが機能しない理由を考え、それが正しいかどうかを確認するように指示します。 LLM にフィードバックなしでオプションを提供してもらい、それぞれの長所と短所を特定して、それを最適な方向に送信します (または LLM であるため、多くの方向性を示します)

すぐに実行してください!) そして何が出るかを見てください。驚いたり、間違ったりする機会を作りましょう。それらはあなたにとって心に残る部分になるでしょう。
LLM は、計画の最初の実装では私よりもはるかに高速ですが、小さな改良でははるかに遅くなります。最初に LLM に慣れてきたときは、とにかく LLM を押して、ガイドを練習するのが良いでしょう。常に使用しているので、その機能の一部を取り戻してください。
メモを書くと覚えやすくなるのと同じように、コードを入力すると覚えやすくなります。自分で小さな変更を加える場合は、コードベースを変更するためにそのコードベースを十分に理解する必要があるのと同じように、LLM がこれまでに行ったことを理解する必要があります。自分が入力したコードと、そこに到達するために学習する必要があったコードをより強く記憶できるようになります。
機能を実装する方法は無数にあります。したがって、試してみれば、それに関する無限の数の質問を考えることができます。これらを自分自身に問いかけてください。これをどのように実装すればよかったでしょうか? LLM はどのようにしてそれを実現したのでしょうか?なぜそのようになったのでしょうか？コードに関して私が知っていることは何か欠けていましたか?そうでしたか？他にどのような選択肢がありましたか?なぜその変数名を選んだのでしょうか?これをより明確に分離する方法はありますか?
この種の好奇心は一般的に育てるのに良いものです。これは、チームメイトほど好みが分からない LLM の場合は特に重要です。また、物事を定着させるための驚きの機会も生み出します。
自分自身にクイズをすることもできます。何も見ずに、この変更に関係するすべてのオブジェクトの名前を言えますか?この変化はどの分野に波及すると予想されますか?変更された各クラスにはどのようなフィールドが存在し、そのタイプは何ですか? LLM が行ったばかりの変更後のシステムのアーキテクチャを概略的に説明してもらえますか?考え、推測し、確認することは、強化するための優れた方法です

読んだばかりの内容の記憶。
LLM で生成されたコードを使用すると、LLM がチャット ウィンドウをすぐに開くという追加の利点があります。自分自身の答えをじっくり考えたら、LLM に質問して、結果を確認することができます。私はこの方法で多くのことを学びました、そしてそれらは定着したものです。また、インタラクティブなレビューは、機能の背後にあるコードにも役立ちました。
GitHub の PR ビューは大規模なプル リクエストに使用するのが非常に難しいと思います。たとえば、変更されたビットのコード コンテキストのみが表示されます。コードベースの一部に触れたことのない PR をレビューしている場合、私にできる最善の方法は高価なリンターです。ローカルで戦術的なフィードバックを与えることはできますが、全体像を見ることはできません。
LLM で生成されたコードのレビューも同様です。差分と会話トレースを好きなだけ見てください。変更が決定された部分だけが表示されます。
大規模な PR には GitHub を使用する代わりに、エディターでコードをチェックアウトします。そうすることで、検索を行ったり、定義にジャンプしたり、ファイル ペインで兄弟ファイルを確認したり、ブラウザで実行したり、ブレークポイントを設定したり、値を検査したりすることができます。これは受動的なレビューではなく、積極的なレビューです。これはマシン上の LLM 生成コードにも機能し、最初にチェックアウトする必要さえありません。大規模な PR と同じように、時間をかけて検討してください。変更だけでなくコンテキストを理解し、それが何をしたのかを学ぶことに積極的に参加してください。
多くの部分を含む大きな変更がある場合は、HTML 説明のアプローチを使用することを好みます。ここで、機能とそのアーキテクチャを説明する単一ページの HTML ファイルを生成するように LLM に依頼します。図、コードへのエディター リンク、インライン コード スニペット、その他コンテキストを把握するのに役立つと思われるものを含めるようにします。
これらは、どこから調べ始めればよいかわからないときに役立ちます。しかし、それは完全な答えではありません。ただ

プログラミング言語に関する本を読んでも、実際にその言語で書き始めない限り、その言語を実際に学ぶことはできないのと同様、これだけではその機能について学ぶのには役立ちません。それでも、大規模な変更や複雑な変更に対処する場合、上記の提案を参考にするのは優れた出発点です。
私はエージェントがメイン リポジトリにコミットしたりプッシュしたりすることを禁止します。私は自分の名前が記載されたコードに責任を持っており、自分が生成した、自分が見ていないコードを他人に見てもらうのは失礼だと考えています。
LLM がファイルを変更した後、他の人の PR レビューを行うかのように、エディタの git ツールでファイルをレビューします。 GitHub で承認できる内容になるまで、コメントを残して質問し、修正を求めます。
エージェントに、AI? という接頭辞が付いているコメントを読んで対処する簡単なスキルを作成してもらいました。または AI: コード内。そうすることで、コードを読みながら注釈を付けて、読み終わったときにコードにフィードバックを処理するように依頼できます。これをスキルにすることで、フィードバックを残すと言った場合にエージェントが自らそれを認識し、一貫した方法でチェックするようになります。これにより、実際のコードレビューのように感じられ、そのアプローチは私にとって非常にうまく機能します。
これには、前に述べたローカルの変更しか表示されないという同じ問題がありますが、コードの入力を許可する前の素晴らしい最終ゲートであり、初期の手順を実行するための良いトリガーになります。
コードの作成にそれほど関与していないと、コードに加えられた決定に疑問を抱く機会が少なくなります。したがって、何かが間違っている、複雑すぎる、または壊れているように見えるときの直感はさらに重要です。
何が問題なのかを知る必要はありません。あなた自身が答えを持っている必要はありません。しかし、LLM で書かれたコードをレビューしているときに、そのような感情が湧き上がってきたら、より積極的に気づく必要があります。

彼らについて尋ねてください。 「これは複雑そう」「X についてはよくわからない」「嫌な予感がする」というだけで、別の方向に進むことがよくあります。コードがどのように生成されるかにより深く関わっているように感じられ、「ああ、このコードをこの形に導くのに私が貢献したんだ」と思い出すでしょう。
これらすべてを毎回行う必要はありませんが、頻繁に行うと自然な習慣になります。はい、機能がリリースされるまで、エージェントに機能を完全に制御させるよりも時間がかかります。しかし、あなたがプロジェクト開発の責任者である場合、あなた自身、あなたのチーム、そして顧客に対してそれを理解する義務があります。そうでない場合は、将来の LLM を自分に代わって理解し、自分よりも適切な決定を下せるかどうかに、自分の将来を賭けていることになります。そしてその時点で、なぜ彼らはあなたを必要とするのでしょうか？
AI によって私たちのコードの書き方がどのように変化しているかについてのさらなる視点については、Aha! を読んでください。エンジニアリングブログ 。
Justin は、長年 Ruby on Rails 開発者、ソフトウェア ライター、オープンソース貢献者として活動しています。彼は Aha! の主任ソフトウェア エンジニアです。 — 世界ナンバー 1 の製品管理ソフトウェア。以前は、Avvo で研究開発チームを率い、人々が必要な法的支援を見つけるのを支援していました。 Justin が Aha! に参加した理由については、こちらをお読みください。
重要なものを構築します。試してみてください、ああ！ 30日間無料。

## Original Extract

If you used an LLM to write code you shipped a month ago, how well do you remember it? Would you realize if the code you committed today would break it? If a customer reported a bug, would you quickly think, "Oh, I know what that is" in the same way

How do you stay familiar with the code when it's written by an LLM? Engineering Articles
Aha! Develop is the agile tool that links strategy to delivery .
How do you stay familiar with the code when it's written by an LLM?
If you used an LLM to write code you shipped a month ago, how well do you remember it? Would you realize if the code you committed today would break it? If a customer reported a bug, would you quickly think, "Oh, I know what that is" in the same way you used to?
As you have LLMs write more code, you'll quickly realize you're less familiar with the code you ship. You have a harder time remembering the decisions you made. You no longer instinctively know where the code lives. If someone asks you a question about it, you have to go back to the LLM instead of answering off the top of your head. Over time, you have to delegate more work and more understanding to the LLM. Your guidance to the LLM gets worse as you lose knowledge of the code. And when the LLM struggles to write the next feature, you don't even know what the cause is, let alone how to help.
Is this just what software development is now ? How can you be responsible for code you don't understand?
Maybe you don't care! Maybe next month's LLM will just solve the problems caused by this month's model. Maybe you are entirely outcome-focused, and the code doesn't matter as long as it solves your problem. Maybe this is a one-off project that just replaces something you would have downloaded a package for, if only it existed — you wouldn't have looked at that code, right?
But if this is an app you or your company depend on, that you are crafting over time, it does matter. People rely on what you produce, and the work you do today has to live on indefinitely.
You have to care about the how, and not just the what.
Writing the code by hand, creating and fixing problems, deciding where this new code goes, refactoring old code to make it better — all these build your intuitive understanding of the code. When someone finds a bug and you immediately guess the answer, it's based on that intuition. It's how you know that changing a value over here will need a migration over there. It's why you have teams that focus on different areas of the code. And if you're not careful, code-generating LLMs will keep you from building that intuitive understanding until it's too late.
Meanwhile, AI uses three signals to decide how to write code: training data, the code it sees in your codebase, and your guidance. As you lose your ability to guide the LLM, and as it sees more of the code it wrote and less of the code you wrote, it doesn't matter what your programming processes and philosophies are. It's dragging your codebase to the mean.
So how can you stay familiar with the code as you write less of it?
Have you ever answered something embarrassingly wrong on a test and still remember the right answer today? As programmers, that failure-to-surprise feeling is why a lot of obscure pieces of information stick with us. When an LLM is generating the code, you don't get to feel that. The process is too smooth.
It's like reading the answers in the back of the book without thinking about the questions first.
Instead, you have to cultivate that feeling of surprise. Look for opportunities to detect mistakes, or even create them. Tell an LLM to try a wild optimization, think about why it won't work, and then see if you're right. Have the LLM give you options without feedback, identify the pros and cons of each, then send it in the best direction (or, because it's an LLM, many directions at once!) and see what comes out. Create opportunities to be surprised and get things wrong. Those will be the parts that stick with you.
LLMs are much faster than me at the initial implementation of a plan, but much slower at small refinements. When you were first getting used to LLMs, it was nice to push them anyway so you could practice guiding them. Now that you use them all the time, take some of that agency back.
Just like writing your notes helps you remember them, typing your code helps you remember it. If you're making small changes yourself, you have to understand what the LLM has done so far, just like you have to understand enough of any codebase in order to change it. You'll more strongly remember the code you typed and the code you had to learn about to get there.
There are an infinite number of ways to implement a feature. So if you try, you can think of an infinite number of questions about it. Ask these to yourself. How would I have implemented this? How did the LLM do it? Why did it do it that way? Did it miss anything I know about the code? Did I? What other options were there? Why did it pick that variable name? Is there a cleaner way to separate this?
That kind of curiosity is good to cultivate in general. It's especially important with LLMs, whose preferences you don't know as well as your teammates'. It also creates those opportunities for surprise that help things stick.
You can even quiz yourself. Without looking, can you name all of the objects involved in this change? Which areas would you expect this change to ripple out to? What fields now exist on each changed class, and what are their types? Could you sketch out the system's architecture after the change the LLM just made? Thinking, guessing, and checking is a great way to reinforce your memory of what you just read.
With LLM-generated code, you have the extra benefit that the LLM has a chat window open right there. Once you think through your own answers, you can ask the questions to the LLM and see what it comes up with. I've learned a lot this way, and they're things that have stuck. And the interactive review has helped the code behind the feature stick too.
I find GitHub's PR view very hard to use for large pull requests. For example, you only see the code context for the bits that changed. If I'm reviewing a PR for a part of the codebase I've never touched, the best I can be is an expensive linter: I can give local, tactical feedback, but can't see the bigger picture.
Reviewing LLM-generated code is similar. Look at the diff and conversation trace all you want — you're only seeing the parts it decided to change.
Instead of using GitHub for large PRs, I check the code out in my editor. That way, I can do searches, jump to definition, look at sibling files in the files pane, run it in a browser, set breakpoints and inspect values … it's an active review, not a passive one. This also works for LLM-generated code on your machine, and you don't even need to check it out first. Just like a large PR, take the time to explore. Understand the context, not just the changes, and take an active part in learning about what it did.
If you have a big change with a lot of pieces, I like using the HTML explainer approach . That's where you ask the LLM to generate a single-page HTML file explaining the feature and its architecture. Have it include diagrams, editor links to the code, inline code snippets, and anything else you think will help you get context.
These are helpful when I'm not sure where to begin exploring. But it's not a complete answer. Just like reading a book about a programming language doesn't really teach you that language unless you start to write in that language, this won't help you learn about the feature by itself. Still, when you're dealing with a large or complex change, it's a great starting point to help with the suggestions above.
I ban my agents from committing and pushing in our main repositories. I'm responsible for the code with my name on it, and I think it's rude to make someone else look at code I generated that I haven't looked at myself.
After the LLM has changed files, I review it in my editor's git tools as if I were doing a PR review for someone else. I leave comments, ask questions, and push for revisions until it's something I'd approve in GitHub.
I had the agent create a simple skill to read and address comments prefixed with AI? or AI: in the code. That way, I can annotate the code as I read and ask it to handle my feedback when I'm done. By making this a skill, the agent will sometimes pick up on it on its own if you mention leaving feedback, and it will check it in a consistent way. This makes it feel like a real code review, and that approach works very well for me.
Although this has the same problem of only seeing local changes that I mentioned earlier, it's a nice final gate before I allow code to go in and a good trigger for me to go through those earlier steps.
When you're not as involved in writing the code, you have fewer opportunities to question the decisions that went into it. So those gut feelings when something looks wrong, overcomplicated, or broken are even more important.
You don't have to know what's wrong. You don't have to have the answer yourself. But you have to more actively notice when those feelings come up as you're reviewing the LLM-written code and ask it about them. "This seems complex," "I'm not so sure about X," or "I have a bad feeling about this" are often enough to guide it onto a different track. You'll feel more involved with how the code is generated, and you'll remember, "Oh, I helped guide this code into this shape."
You don't have to go through all of these every time, but they become natural habits when you do them frequently. Yes, it'll take longer than giving an agent total control of a feature until it ships. But if you're the one responsible for the development of a project, you owe it to yourself, your team, and your customers to understand it. Otherwise, you're betting your future on a future LLM to understand it for you and make better decisions than you. And at that point, why do they need you?
For additional perspectives on how AI is shifting how we write code, read the Aha! engineering blog .
Justin is a longtime Ruby on Rails developer, software writer, and open-source contributor. He is a Principal Software Engineer at Aha! — the world's #1 product management software . Previously, he led the research and development team at Avvo, where he helped people find the legal help they need. Read about why Justin joined Aha!
Build what matters. Try Aha! free for 30 days.

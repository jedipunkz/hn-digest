---
source: "https://chriskiehl.com/article/evolving-thoughts-on-ai-2026"
hn_url: "https://news.ycombinator.com/item?id=48711187"
title: "Evolving Thoughts on AI in 2026"
article_title: "Evolving thoughts on AI - Blogomatano"
author: "goostavos"
captured_at: "2026-06-28T20:24:59Z"
capture_tool: "hn-digest"
hn_id: 48711187
score: 1
comments: 0
posted_at: "2026-06-28T20:19:00Z"
tags:
  - hacker-news
  - translated
---

# Evolving Thoughts on AI in 2026

- HN: [48711187](https://news.ycombinator.com/item?id=48711187)
- Source: [chriskiehl.com](https://chriskiehl.com/article/evolving-thoughts-on-ai-2026)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T20:19:00Z

## Translation

タイトル: 2026 年の AI に関する進化する考え
記事のタイトル: AI に関する進化する考え - Blogomatano
説明: インターネットに何が必要か知っていますか? AIについて意見を持つ別の人

記事本文:
AI に関する進化する考え - Blogomatano
すべては異なりますが、同じです。 2026 年、AI は特異点ではなく、より速い馬のように感じられます。はい、はい、これは革命的であり、誰もがこれなしでは決して実行できなかった素晴らしいことを実行できるようになりますが、出力速度以外に実際に何が変わったのでしょうか?システムの構築と保守に関して難しかったことはすべて、依然として難しいようです。
しかし、状況は急速に変化します。昨年の私の意見は古いです。老化が進む可能性のある意見をさらに紹介します。
私はこれが何よりも大好きです。木工用語で言えば、「ショップ治具」を作るのが経済的になっています。これらはあなたのセットアップにおいてのみ意味を持ちます。コミットする価値があることはほとんどありません。多くの場合、一度使用したら捨てられるように作られています。 AI が登場する以前は、ツールは投資の対象でした。時間を費やすことを正当化する必要がありました。そのため、一般的には広く「役立つ」ものでなければなりませんでした。そのわずかな時間は、「出荷機能」などと天秤にかけられ、常に失われていきました。必要なときにどこでもオンデマンドで実行できるようになりました。生の生産性の向上という点では、これが最大の勝者に違いありません。
大規模なリファクタリングにも自信を持って取り組むことができます。数時間かかる変更も数分で完了します。優れたコードを書くのがこれまでより簡単になりました。
私は、変更を手動で開始し、その後、より広範なコードベースにフラッシュするという面倒な部分を終了するために変更をクロードに渡すという開発スタイルに落ち着きました。もう手作業でコードを書くことに抵抗を感じている人もいると思いますが、私にとってはそれがうまくいくとは思えませんでした。巨大な /plan 会話で自分が何を望んでいるかを説明するのに非常に多くの時間を費やしているため、単に変更を加え、それを指差し、「どこでもこのようにしてください」と言うよりも遅くなります。
私が今でもほぼすべての作業を手作業で行っている場所は 1 つあります。

テスト中。クロードは、プロパティ ベースのテスト ライブラリの API についてはすべて知っていますが、その使用方法を理解するのに非常に苦労しています。コードベースには参考になるサンプルが豊富にあるにもかかわらず、それでもおかしなことをするので、結局自分でやってしまうことになります。悪いテストスイートは、良いコードベースを台無しにする可能性があります。常に警戒が必要です。今まで以上に。
さらに、優れたプロパティベースのテストは、疑わしいほど小さいです。大量のテスト コードを生成するのは、通常、思考が不十分であることの症状です。
ソフトウェア開発の多くは、退屈で面白くない単調な作業です。これを AI にオフロードすることは素晴らしいことです。
CDK?私の人生を常に台無しにするあの巨大な犬の糞の山？それが今のクロードの問題だ。私の小さな友達が、私の個人的なスタックを自動的に最新の状態に保ってくれます。 AI が実稼働環境を破壊する別のチームになることをそれほど心配していなかったなら、私はクロードにすべての CDK デプロイメントを管理してもらい、自由になれるようにするでしょう。しかし...その日は（もしあったとしても）遠い将来です。
生活の質の小さな改善も気に入っています。コード変更で最も厄介な部分は、フォーマッタの実行、リンターの問題の解決、ビルドの問題の解決、テストの実行などの最後の忙しい作業です。以前は、発生する小さな問題 (「おっと。print ステートメントが残っている」) を処理するために、子守りと頻繁な再起動が必要でした。今？私が他のことをしている間、AI は一線を超えてしまいます。
AI コードレビューは煩わしいノイズから便利なものに変わりました。これは私が常に最も興奮している分野です。私がタイプを好むのには理由があります。私が静的解析を好むのには理由があります。ライフサイクルの早い段階で問題を発見できるツールが欲しいです。 AI コード レビュー ツールは、最後の防御手段として最適です。
しかし、AI コード レビューが役立つからといって、AI だけがコード レビューを行う必要があるわけではありません。

コード。 （私見）
これを真剣に提案する人の数は増えています（以下の「もっと速く」を参照）。同じ2つのエッセイを何度も投稿する以外に何をすべきかわかりません。
理論構築としてのプログラミング
私はそれをレガシーシステムに関連付けます。レガシー プロジェクトでの作業が難しくなる原因は、コードの詳細とはほとんど関係がなく (ただし、通常はゴミです)、すべてはコードに関連付けられた「理論」がないという事実と関係があります。サービスは、あなたが現れたときに実行されていた方法なので、実行されている方法で実行されます。別の方向に走行して警報が鳴ったときは?知るか。サービスは「あなたのもの」ではありません。それがそのように動作するのは、帰納法によって時間の経過とともに近似することしかできない哲学のためです。その運用モデルを導き出した無数の目に見えない決定は、あなたには利用できません。
ただ、一部のチームがこれを試みているのは良いことだと思います。おそらくピーター・ナウルは愚か者であり、クロードのサブスクリプションを持っている男の方がよく知っています。時間が経てばわかるでしょう。今のところ、私はオンコールローテーションをしなければならず、ポケベルを持ち歩き、顧客に影響を与えると怒鳴られます。したがって、古い習慣はなかなか消えません。コードに深く関わることが長期的には価値があると考える私は、間違いを犯すつもりです。
CR コメントに対する一般的な反応は次のとおりです。
「ああ、分からない。クロードがやったのに、私は見ていなかった」
あるいは、AI にコメントへの応答を投稿させるという、はるかに悪質な行為を行っています。
「あなたがこれを押し返すのは正しいことです。私は何とか何とかするべきです。」
これは大学の新入社員やインターンの間で非常に蔓延しています。 Slack での会話をクロードに渡すこともできます。同じ人間がクロードにとって損失の多いインターフェースに過ぎないことを選択したと知ったときに感じる失望のレベルを言葉で表現するのは難しいです。
開発者は決して得意ではありませんでした

テストを書くこと。しかし、少なくとも彼らはそれらを書きました。昔は、テストに合格するには、ある程度のレベルでコードに取り組む必要がありました。今では、見てもいない、何もしない「テスト」で埋め尽くされた CR をカットする人がいます。これが初心者だけのものであることを望むかもしれませんが、どうやら誰もがそれを受け入れているようです。私は AI が生成するスロップについてチームに多くの余地を与えていますが、テストは「この丘で死んでも幸せに死ぬ」ことができる唯一の領域です。
労働時間は騒音の不協和音です。常にコンテキストを切り替えたり、セッションをジャグリングしたりする必要があります。過去数か月間、私が行ったひどい設計コールの数は天井を突き抜けました。タスクの読み込みは効果的に管理できる範囲を超えています。
修正しなければならない間違った情報が大量にあると、大変な労力と時間がかかります。 「はい、はい、クロードがそのようにうまくいくと言ったのは知っていますが、それは違います。」 「はい、はい、そのように機能することを『証明』する RFC にリンクされているのは知っていますが、実装は...から逸脱することがよくあります。」 「はい、お見せできますが、それには時間がかかります...」 これは、最終的に質問者から継続的な不信感を受け取るだけの反例を引き出すために大量の時間を浪費するまで続きます。クロードは彼らが正しいと言いました！なぜ彼らが正しいことを認めないのですか！
この問題により、特定の種類のタスクを委任することが非常に困難になっています。クロード セッションの前に座ってその出力を無批判に消費する生ぬるい体は、良いことよりも害を引き起こす可能性があります。しかし、これは今までとそれほど変わりません。あいまいなタスクを実行してくれると信頼している人もいれば、「このチケットの指示に従ってください（クロードに）」と十分な範囲を与える人もいます。コツは 2 つのグループを混同しないことです。
リーダーシップチェーン全体が、「より速く進む」という 1 つの論点を採用しました。 「より少ない労力でより多くのことを実現する。」もっと、もっと、もっと。しかし、「もっと」

それも未定義のままです。何と比較して「より速く進む」でしょうか？誰も本当のことは言えません。ただ「もっと」やろう。もっと早く。
それは多くの人々から工芸品への愛を奪いました。 「エンジニアリング文化」は常に少数の人々によって構成されてきました。情熱的な人々は常に多勢に無勢です。今、彼らは溺れています。リーダーたちは「もっと」を望んでおり、今もさらにそれを望んでいます。何年にもわたって安定した開発ペースを提供して称賛されてきた企業も、今では自分たちが「ボトルネック」になっている、「スケールしていない」、または「他の人の速度を遅らせている」ことに気づいています。彼らは、「9」の数を誇りにしていた文化が周囲で崩壊していくのを目の当たりにしています。その代わりに、再び「間違いの修正」に慣れるように言われます。他のすべてを犠牲にして速く進みます。多くの人がこの新しい世界で苦労しています。
私たちは経済スケジュールに取り組んでいます。長い目で見れば、誰が正しいか分かるだろう。しかし、それには不快な時間がかかります。誰かの手帳にそれが感じられて初めて調整が行われます。
私は、SpaceX の初期の頃を記録した『Liftoff』という本のことをよく考えています。彼らは信じられないほどのペースで動きましたが、継ぎ目でバラバラになり始める前に反復できる速度には上限があることもわかりました。今、私たちはその限界まで突き進んでいるような気がします。しかし、多くの人がそれを超えています。物も人も壊れ始めています。
インターネット上では、「スキルの喪失」や、誰もどうすればよいか分からなくなるなど、さまざまな議論が巻き起こっていますが、世の中には、新しいトピックを深く学ぶために AI を使用する好奇心と関心を持つ人がまだたくさんいるようです。現在インターネット上でよくある「私はとても賢いので TLA+ を促す」という派手な方法ではなく、静かで退屈な方法で、つまり彼らがそれを面白いと思っているからです。

## Original Extract

You know what the internet needs? Another person having opinions on AI

Evolving thoughts on AI - Blogomatano About
Everything is different and yet the same. In 2026, AI feels less like the singularity and more like a faster horse. Yes, yes, it's revolutionary and enables everyone to do wonderful things they could have never done without it, but what has actually changed beyond output speed? Everything that was hard about building and maintaining systems still seems hard.
But things change fast. My opinions from last year are outdated. Here are more opinions that will age poorly.
I love this more than anything else. To put it in woodworking terms, it's now economic to create "shop jigs." They only make sense for your setup. They're rarely worth committing. They're often built to be used once and thrown away. Pre-AI, tooling was something that was an investment . You had to justify spending the time. And because of that, it generally had to be broadly "useful." That scarce time was weighed against "shipping features" or whatever and always lost. Now you can crank them out on-demand, wherever you need them. In terms of raw productivity boosts, this must be the biggest winner.
Big refactors can be tackled with confidence. Changes that would take hours can be done in minutes. It's easier than ever to write good code.
I've settled into a development style that starts a change by hand but then hands it off to Claude to finish up the tedious part of flushing it through the wider codebase. I know some don't bother writing code by hand anymore, but I haven't found that to work for me. I spend so much time trying to explain what I want in a giant /plan conversation that it ends up slower than me just making the change, pointing at it, and saying "do it like this, but everywhere"
The one place I still do just about everything by hand is testing. Claude knows everything about the APIs of property based testing libraries, but really struggles to understand how they should be used. Despite the ample examples in our codebase from which to draw, it'll still do goofy things that lead to me just doing it myself. A bad test suite can ruin a good codebase. Constant vigilance is required. Now more than ever.
Plus, good Property Based Tests are suspiciously small. Generating a bunch of test code is usually a symptom of insufficient thinking.
A lot of software development is tedious, uninteresting grunt work. Offloading this to AI has been awesome.
CDK? That gigantic pile of dogshit that constantly ruins my life? That's Claude's problem now. I've got my little buddy automatically keeping my personal stacks up to date for me. If I weren't so worried about becoming Yet Another Team that has AI destroy a production environment, I'd lean harder into having Claude manage all CDK deployments so that I can be free. But... that day is far in the future (if ever).
I also love the little quality of life improvements. The most annoying part of any code change is the final busy-work of running formatters, resolving linter issues, resolving build issues, running tests, etc. It previously required baby sitting and frequent restarts to handle the small issues that arise ("oh, whoops. Left over print statement"). Now? AI kicks it over the line while I go do other things.
AI code review has swapped from irritating noise to useful. This has always been the area I'm most hyped about. I like types for a reason. I like static analysis for a reason. I want tools that catch problems earlier in the lifecycle. AI Code Review tools have become a great last line of defense.
But AI code review being useful doesn't mean only AI should review code. (imho)
The number of people suggesting this in earnest is growing (see: "go faster" below). I don't know what to do other than post the same two essays over and over.
Programming as Theory Building
I relate it to legacy systems. The thing that makes working with legacy projects hard has very little to do with the specifics of the code (though, it's usually trash) and everything to do with the fact that you have no "theory" associated with the code. The service runs the way it runs because that's the way it was running when you showed up. When it runs a different way and alarms fire? Who knows. The service isn't "yours." It behaves the way it does because of a philosophy you can only approximate over time through induction. The innumerable, invisible decisions which led to its operating model are not available to you.
I do think it's good that some teams are trying this, though. Maybe Peter Naur was an idiot and the dude with a Claude subscription knows better. Only time will tell. For now, I have to do on-call rotations, carry a pager, and get yelled at when we affect customers. So, old habits die hard. I'm going to err on the side that thinks being deeply engaged with the code is valuable over the long term.
The increasingly common response to CR comments:
"Oh, I dunno. Claude did that and I didn't look at it"
Or they do the far, far worse thing of have their AI post a response to the comments.
"You're right to push back on this. I should have blah blah"
This is extremely prevalent amongst fresh college hires and interns. They'll even ferry Slack conversations to Claude. It's hard to articulate the level of disappointment I feel when I realize a fellow human being has chosen to be nothing more than a lossy interface for Claude.
Developers have never been good at writing tests. But at least they wrote them. In the old days, they'd have to engage with their code on some level in order to get their tests passing. Now people cut CRs filled with "tests" that they haven't looked at and that do nothing. You'd hope this was isolated to newbies, but seemingly everyone has embraced it. I allow a lot of leeway on my team for AI generated slop, but tests are the one area where "I'll die on this hill and die happy"
Work hours are a cacophony of noise. You're constantly context switching and juggling sessions. The number of terrible design calls I've made over the last few months has shot through the roof. The Task loading is beyond what can be managed effectively.
The amount of wrong information that must be corrected is exhausting and time consuming. "Yes, yes, I know Claude said it works that way, but it's wrong." "Yes, yes, I know it linked you to the RFC that 'proves' it works that way, but implementations often depart from..." "Yes, I can show you, but it would take time for me to..." This continues on until you waste a bunch of time pulling a counter example to ultimately just receive continued incredulity from the asker. Claude said they were right! Why aren't you accepting that they're right!?
This problem has made it very hard to delegate certain kinds of tasks. The lukewarm bodies sitting in front of a Claude session uncritically consuming its output can cause more harm than good. But this isn't that different from how it's always been. There are people you trust to do ambiguous tasks and there are people you give well scoped " (have Claude) do what this ticket says." The trick is not confusing the two groups.
The entire leadership chain has adopted one talking point: go faster . "Do more with less." More, more, more. But "more" also remains undefined. "Go faster" relative to what? Nobody can really say. Just do "more." Faster.
It has stolen the love of the craft from many people. "Engineering culture" has always been comprised of a small few. The passionate have always been outnumbered. Now they're drowning. Leadership wants "more" and they want more now. Those who have be celebrated for providing a steady pace of development over the span of years now find themselves "a bottleneck" or "not scaling" or "slowing others down." They're watching a culture that took pride in the number of "9s" disintegrate around them. In its place, they're told to get used to "Correction of Errors" again. Go fast at the expense of all else. Many struggle in this new world.
We're working on economic timelines. Over the long haul, we'll see who is right. But it will take an uncomfortable amount of time. Adjustments will only be made once it's felt in somebody's pocket book.
I find myself thinking a lot about the book Liftoff, which chronicles the early days of SpaceX. They moved at an incredible pace, but also found that there's an upper limit to how fast you can iterate before things start falling apart at the seams. Right now, it feels like we're pushing up to that edge. But many are going beyond it. Things, and people, are beginning to break.
For all the kvetching on the internet about "skill loss" and how nobody will know how to do anything anymore, there still seems to be a lot of curious, interested people out there who use AI to learn new topics deeply. Not in the flashy "I prompt TLA+ because I am very smart" way that's so common on the internet right now, but in quiet, boring way: because they find it interesting.

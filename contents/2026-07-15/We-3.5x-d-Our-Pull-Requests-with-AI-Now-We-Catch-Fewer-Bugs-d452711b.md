---
source: "https://brodzinski.com/2026/07/3x-pull-request-size-ai.html"
hn_url: "https://news.ycombinator.com/item?id=48920182"
title: "We 3.5x'd Our Pull Requests with AI: Now We Catch Fewer Bugs"
article_title: "We 3.5x'd Our Pull Requests with AI: Now We Catch Fewer Bugs - Pawel Brodzinski on Leadership in Technology"
author: "flail"
captured_at: "2026-07-15T13:12:50Z"
capture_tool: "hn-digest"
hn_id: 48920182
score: 1
comments: 0
posted_at: "2026-07-15T12:59:17Z"
tags:
  - hacker-news
  - translated
---

# We 3.5x'd Our Pull Requests with AI: Now We Catch Fewer Bugs

- HN: [48920182](https://news.ycombinator.com/item?id=48920182)
- Source: [brodzinski.com](https://brodzinski.com/2026/07/3x-pull-request-size-ai.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T12:59:17Z

## Translation

タイトル: AI を使用してプル リクエストを 3.5 倍にしました: 捕まえるバグが少なくなりました
記事のタイトル: AI を使用してプル リクエストを 3.5 倍にしました: 捕まえるバグが減りました - Pawel Brodzinski がテクノロジーのリーダーシップについて語る
説明: AI によりプル リクエストは 3.5 倍大きくなりましたが、欠陥検出は低下しました。 2 つのプロジェクトの PR データは、AI によって生成されたコードが品質に与える影響を示しています。

記事本文:
AI を使用してプル リクエストを 3.5 倍にしました: 捕まえるバグが減りました - Pawel Brodzinski 氏がテクノロジーのリーダーシップについて語る
パヴェル・ブロジンスキーがテクノロジー分野のリーダーシップについて語る
チームを率い、製品を構築し、ビジネスを運営するために必要なことは何でも
AI を使用してプル リクエストを 3.5 倍にしました。今では、バグの検出が少なくなりました
私は過去に、プル リクエストの平均サイズが桁違いに増加したといくつか主張しました。私は主に、Lunar Logic の開発者との廊下での会話に基づいています。
実際のデータを確認するのはそれほど難しくないことがわかりました。ああ、私は間違っていました。
私たちは、複雑さ、労力、チーム規模の点で非常によく似た 2 つのプロジェクトからデータを抽出しました。どちらのコードベースもグリーンフィールドから開始され、分析期間中に数百のプル リクエストをカバーしました。エンジニアリング チーム全体で重複する部分もありました。重要な違いは何でしょうか？ AIをどのように活用したか。
最初のプロジェクトは、「 Helpful 」と呼ぶことにしますが、クロード以前の時代に起こりました。開発をサポートするためにすでに AI を使用していましたが、主にオートコンプリートで、厄介な問題の解決策を提案するために ChatGPT に時折アクセスしていました。すべてのコードは開発者によってリアルタイムで管理されました。
2 番目のライブは、Grateful と呼ぶことにしますが、クロードの全力演奏でした。基本的な前提は、コードが手作業で書かれたものではないということでした。エンジニアリングの責任は、コンテキストの管理、プロンプト、およびレビューにありました。
結論は? AI を多用するプロジェクトでは、平均 PR サイズが 3.5 倍に増加しました。
はい、私は間違っていましたが、それは変化の相対的な規模についてのみでした。
AI のおかげで、私たちはより大きな課題に取り組めるようになったと言うのは簡単でしょう。ただし、データが示唆しているのはそうではありません。少なくとも PR の分布を見るとそうではありません。
注: 追加された行と削除された行を合計して、PR サイズを計算します。
どちらの場合も、PR の大部分は小規模から中規模の PR です。数百行の

コードトップ (カットオフラインは 500 でした)。それはエンジニアリングの仕事の糧でした。それはまだあります。
確かに、Project Helpful (「AI なし」) では、これらの PR は小さいサイズに偏っていましたが、Project Grateful では、コードの重心が 100 ～ 200 行重くなりました。それでも、コードベースに精通したエンジニアにとって、それは難しいことではありません。
では、なぜ平均値がそこまで上昇したのでしょうか？
答えは周辺にあります。最も小さなプル リクエスト、つまりよく言われるワンライナーは、ほとんど消え去りました。それが最も重要な変更です。小規模な PR は 5 人に 1 人でしたが、現在は 20 人に 1 人です。
レビューするのが最も簡単な作業項目のクラス全体が消滅の危険にさらされています。この考えはやめましょう。すぐに戻ります。
最も小さな作業で失ったものは、最も大きな作業で補います。
大規模な PR の数はほぼ 3 倍です。
90 パーセンタイルのサイズは、600 行のコードから 1799 行のコードへと 3 倍に増加しました。
異常値はさらに膨らみました。Project Grateful の最大の PR は 30,000 行以上のコードであり、Project Helpful の同等のコードのほぼ 7 倍でした。
ビッグはさらに大きくなりました。そして、私たちもそれ以上のものを得ることができます。
それでも、これらのコードの塊が作業の大半を占めるわけではありません。間違いなく、まだではありません。ただし、注目を集めるには十分な数があります。
偶然にも、これはレビュー担当者にとって最も困難なクラスの項目です。もう、おそらくそれがどこに向かっているのか推測できるでしょう。
3 倍のコード行を処理する効果
データをどのようにスライスしても、私たちは現在、以前のおよそ 3 倍の大きさのタスクを扱っているようです。
平均 PR サイズは 232 LoC から 817 LoC に増加し、3.5 倍に増加しました。
PR サイズの中央値は 66 から 210 LoC に増加し、3.2 倍になりました。
大規模および大規模な PR の割合は 13% から 33% に増加しました。

.5倍増加。
簡単に言うと、私たちの脳はタスクごとに以前の 3 倍の情報を処理します。常識的に考えて、レビューを少しずつ行う場合ほど徹底したものにはできないことが考えられます。
研究では一致しているようです。広く知られた Smart Bear/Cisco の調査では、プル リクエストのサイズを 200 行未満に抑えることが推奨されています。そのサイズを超えると、レビュー担当者は問題を見落とし始めます。
「レビュー担当者は、少量のコードをレビューするときに最も効果的です。200 行未満では、比較的高い率で欠陥が発生し、多くの場合、平均の数倍になります。その後、結果は大幅に減少します。250 行を超えるレビューでは、コード 1000 行あたり 37 個を超える欠陥は発生しませんでした。」
概要: AI を多用すると、個々の作業が大きくなり、全体的な品質が低下します。
品質の低下は避けられないが、その可能性は高い
変化を考慮すると、どれも避けられないように思えます。つまり、クロード コードをより小さな単位で動作させることができるので、レビューしやすくなります。なるほど、Smart Bear の研究でアドバイスされた注釈技術を使用して作成できるのですね。その結果、ほとんどの品質基準を維持する必要があります。
問題は 1 つだけです。私たちはこれらのいずれも行いません。
そのためには、エンジニアがコーディング エージェントを人為的に抑制する必要があります。それは、人間とそのツールの間の行き来がさらに増えることを意味します。それは私たちの「怠惰」の本能に反して働くでしょう。
エージェントが大きなタスクを処理する場合、なぜそれを小さなタスクに分割し、徐々に 1 つずつレビューする必要があるのでしょうか?コード変更が 200 行に近づくたびに停止するよりも、全体を一度に実行した方が効率的ではないでしょうか? （ちなみに、そうではありませんが、それはまた別の話です。）
最後に、現時点で、ここで目に見える効率の向上が得られますが、一方で、Qu のコストは下がります。

現実は未来に延期されます。残念ながら、高い品質を維持するためのエンジニアリング手法に固執する可能性は非常に低いように思えます。
私のお気に入りの質問をしてみましょう。エンドゲームはどのようなものですか?
モデルの機能により、ますます大規模なコーディング タスクを処理できるようになるため、一般的なプル リクエストのサイズは増加します。
その結果、レビュー担当者はますます多くの欠陥を見落とすことになります。
レビュー担当者が概念的にコードを深く掘り下げていないという事実は、品質の問題を悪化させるだけです。
したがって、私たちは欠陥だらけのソフトウェアをますます開発することになります。
前述の欠陥により、コーディング担当者と人間の両方に手戻りが発生します。
再作業 (および再作業の再作業) に多くの労力が費やされるにつれて、付加価値アイテムの納品ペースは必然的に遅くなります。
私たちはもっと速く進みますが、過去と同じくらい遅くなるだけです。あるいは、さらに遅い。つまり、ループ内の人間のレビュー担当者のアイデアに固執すると仮定します。そしてそれは認められない。
何十年もの間、私たちは小さなバッチで作業することを学ぼうと努めてきました。難しい道ですが、付け加えさせてください。今では、AI のおかげで、私たちは何も問題なかったかのように U ターンしています。悪い知らせがあります。それは今でもそうです。それは決してソフトウェア固有のものではありませんでした。実際、私たちは最初からそれを製造段階から盗んだのです。
少量ずつ再学習していきます。私たちが思っているよりも早く。
読んでいただきありがとうございます。新しい記事をメールで受け取るために登録していただければ幸いです。
また、 Pre-Pre-Seed サブスタック にも公開しています。そこでは、初期段階の製品開発に関連するものにさらに限定的に焦点を当てています。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
フォローアップのコメントを電子メールで通知する
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
AI を使用してプル リクエストを 3.5 倍にしました。今では、バグの検出が少なくなりました
皆さん、ソースをチェックしてください！
コンウェイの法則

製品開発における AI に関する厳しい教訓を教える
究極の質問: エンドゲームはどのようなものですか?

## Original Extract

AI made pull requests 3.5x bigger, but defect detection dropped. PR data from two projects shows impact of AI-generated code on quality.

We 3.5x'd Our Pull Requests with AI: Now We Catch Fewer Bugs - Pawel Brodzinski on Leadership in Technology
Pawel Brodzinski on Leadership in Technology
Whatever it takes to lead a team, build a product, and run a business
We 3.5x’d Our Pull Requests with AI: Now We Catch Fewer Bugs
I made a few claims in the past stating that the average size of pull requests went up by an order of magnitude. I largely based it on hallway conversations with developers at Lunar Logic.
It turns out, the actual data is not that hard to check. Aaaand I was wrong.
We pulled data from two very similar projects in terms of complexity, effort, and team size. Both codebases started greenfield and covered a few hundred pull requests in the analyzed period. There was even an overlap across the engineering team. The key difference? How we used AI.
The first project, let’s call it Helpful , happened in the pre-Claude era. While we already used AI to support development, it was predominantly autocomplete with occasional trips to ChatGPT to suggest solutions to pesky problems. All code was managed by developers in real time.
The second gig, I’ll call it Grateful , was full-on Claude. The basic assumption was that none of the code was written by hand. Engineering responsibilities were in context management, prompting, and review.
The bottom line? In an AI-heavy project, average PR size increased by a factor of 3.5.
Yes, I was wrong, but only about the relative scale of the change.
It would be easy to say that AI made us tackle bigger tasks. That’s not what data suggests, though. At least not when we look at the distribution of the PRs.
Note: We sum the added and removed lines to calculate the PR size.
The bulk of PRs in both cases are small to medium ones. A few hundred lines of code tops (our cutoff line was 500). It was the bread and butter of engineering work. It still is.
Sure, in Project Helpful (the “no AI” one), these PRs were skewed toward smaller sizes, while in Project Grateful , the center of gravity was 100-200 lines of code heavier. Still, for an engineer familiar with the codebase, that’s not a challenge.
So, how come the average went up that much?
The answer is on the fringes. The smallest pull requests—the proverbial one-liners—all but disappeared. That’s the single most significant change. Tiny PRs were 1 in 5. Now they are 1 in 20.
The whole class of work items that was easiest to review is at risk of extinction. Let’s park this thought. I’ll be back to it soon.
What we lost from the tiniest bits of work, we make up for with the largest.
There are almost 3 times as many large PRs.
The 90th percentile size increased 3x , from 600 lines of code to 1799 lines of code.
The outliers inflated even more—the largest PR in Project Grateful was 30k+ lines of code, almost 7 times bigger than its equivalent in Project Helpful.
Big just got bigger. And we get more of it, too.
Still, these chunks of code do not dominate the work. Definitely not just yet. However, there are enough of them to start paying attention.
Coincidentally, this is a class of items that is the most challenging for a reviewer. By now, you can probably guess where it is heading.
The Effects of Processing 3 Times as Many Lines of Code
No matter how I slice the data, it seems that we now deal with tasks that are roughly three times as big as they used to be.
The average PR size went up from 232 to 817 LoC—a 3.5x increase .
The median PR size went up from 66 to 210 LoC—a 3.2x increase .
The percentage of big and large PRs went up from 13% to 33%—a 2.5x increase .
Long story short, our brains process three times as much information per task as they used to. Common sense suggests that the review can’t be as thorough as it was when done in smaller bits.
Research seems to concur. Well-recognized Smart Bear/Cisco study advises keeping pull request size below 200 lines. Above that size, reviewers start overlooking the issues.
“Reviewers are most effective at reviewing small amounts of code. Anything below 200 lines produces a relatively high rate of defects , often several times the average. After that the results trail off considerably ; no review larger than 250 lines produced more than 37 defects per 1000 lines of code.”
Executive summary: Heavy use of AI makes individual chunks of work larger, and thus, overall quality drops.
Quality Drop Is Not Inevitable But Highly Likely
If we consider the changes, none of them seems inevitable. I mean, we can tell Claude Code to work in smaller chunks so it’s more convenient to review. Heck, we can make it use the annotation technique advised by the Smart Bear study. As a result, we should sustain most of the quality standards.
There’s only one issue. We won’t do any of these.
It would require engineers to artificially throttle their coding agents. It would mean more back-and-forth between humans and their tools. It would work against our “laziness” instincts.
If an agent handles a big task, why should we split it into smaller ones and review them gradually one by one, again? Isn’t it more effective to have the whole thing run at once rather than stopping it each time it approaches 200 lines of code changes? (By the way, it isn’t, but that’s another discussion.)
Finally, we get the perceived efficiency gains right here, right now, while the cost of lower quality is deferred to the future. Sadly, sticking to the engineering practices that kept the quality high seems highly unlikely.
We could ask my favorite question: What does the endgame look like?
As the capabilities of the models allow them to handle larger and larger coding tasks, the typical pull request size will go up.
As a result, reviewers will overlook more and more defects.
The fact that reviewers don’t dive deeply into the code conceptually will only exacerbate the quality issue.
Thus, we will increasingly develop software riddled with defects.
Said defects will add rework for coding agents and humans alike.
The pace of delivery of value-adding items will necessarily slow down, as more effort goes into rework (and rework of rework).
We will go so much faster, only to go as slow as we did in the past. Or slower still. That is, assuming that we stick to the idea of the human reviewer in the loop. And that’s not granted.
For decades, we tried to learn to work in small batches. The hard way, let me add. Now, with AI, we’re making a U-turn as if none of it mattered. I have bad news. It still does. It was never a software-specific thing. In fact, we stole it from manufacturing in the first place.
We will relearn small batches. Sooner than we think.
Thank you for reading. I appreciate if you sign-up for getting new articles to your email.
I also publish on Pre-Pre-Seed substack , where I focus more narrowly on anything related to early-stage product development.
Your email address will not be published. Required fields are marked *
Notify me of followup comments via e-mail
Save my name, email, and website in this browser for the next time I comment.
We 3.5x’d Our Pull Requests with AI: Now We Catch Fewer Bugs
Check Your F**cking Sources, People!
Conway’s Law Teaches a Grim Lesson About AI in Product Development
The Ultimate Question: What Does the Endgame Look Like?

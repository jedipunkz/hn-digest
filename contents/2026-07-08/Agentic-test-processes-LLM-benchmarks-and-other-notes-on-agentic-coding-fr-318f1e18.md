---
source: "https://danluu.com/ai-coding/#llm-variance"
hn_url: "https://news.ycombinator.com/item?id=48836951"
title: "Agentic test processes, LLM benchmarks, and other notes on agentic coding fr"
article_title: "Agentic test processes, LLM benchmarks, and other notes on agentic coding from Galapagos Island"
author: "lifeisstillgood"
captured_at: "2026-07-08T20:57:00Z"
capture_tool: "hn-digest"
hn_id: 48836951
score: 3
comments: 1
posted_at: "2026-07-08T20:21:45Z"
tags:
  - hacker-news
  - translated
---

# Agentic test processes, LLM benchmarks, and other notes on agentic coding fr

- HN: [48836951](https://news.ycombinator.com/item?id=48836951)
- Source: [danluu.com](https://danluu.com/ai-coding/#llm-variance)
- Score: 3
- Comments: 1
- Posted: 2026-07-08T20:21:45Z

## Translation

タイトル: エージェント テスト プロセス、LLM ベンチマーク、およびエージェント コーディングに関するその他の注意事項
記事のタイトル: エージェント テスト プロセス、LLM ベンチマーク、およびガラパゴス島のエージェント コーディングに関するその他のメモ

記事本文:
ガラパゴス島のエージェント テスト プロセス、LLM ベンチマーク、およびエージェント コーディングに関するその他のメモ
|パトレオン
私は昨年 11 月から AI をかなり頻繁に使用していますが、すべてが面白い経験です。エージェントは、人間がやったら即刻解雇するようなことをするでしょう。もちろん、私の反応は、あたかもこれが素晴らしいことであるかのように振る舞い、さらに多くのことを実行できるように 1,000 人のエージェントを立ち上げるというものです。
昨年の半ば、私は GPT (おそらく 5.0 か 5.1) にバグの原因を見つけさせました。当然のことながら、このコードにはテストがなく、git bisect は機能しません。また、これは UI インタラクションのバグであり、私にはテストを書く資格さえありませんでした。そのため、このバグを引き起こしたコミットを見つけるために日付 X と Y の間を二等分するように Codex に依頼しました。 Codex はすぐに、問題のコミットがこの日付範囲以降にあることを教えてくれました (これが正しいはずはありません)。 Codex にこれは間違っていると伝えると、明らかに問題のあるコミットではないコミットが 1 回か 2 回報告されました。それらが間違っていると伝えると、問題のコミットはもっともらしいコミットであると教えてくれました。私がその理論を証明するか反証するかを尋ねると、テストを作成し、疑わしいコミットが破壊的コミットであることを確認したと言いました。
次に、通常のブラウザー テスト環境で開発者の完全なエンドツーエンド スタックを使用したビデオを作成して見せてもらうように依頼しました。それを行う権限がないと主張していましたが (これは嘘でした)、適切なテスト コードを使用して Playwright でのコミットの前後の再現実行のビデオを作成できました。このビデオは説得力があり、コミット前は機能が正常に動作していたが、コミット後には機能しなくなることが示されていました。これについて何かが正しく感じられなかったので、コミットの前後に手動で問題を再現してみました。

そして、すべてが捏造だったことがわかりました。このビデオでは Codex がバグを再現したかのように見えましたが、それは実際の環境ではなく、偽の再現を作成するために設計された人工的なブラウザ環境でした。
先ほども言いましたが、これは皮肉ではなく非常に素晴らしい経験だったので、すぐに「どうすればもっとこれを得ることができるだろう?」と考えました。そして、エージェントをますます頻繁に使用するようになり、昨年中下旬にはコーディング エージェントを頻繁に使用するようになりました。
この投稿では比較的異なる一連のトピックを取り上げているため、ここでは簡単な概要を示します。
エージェントループとこの記事の執筆
人々がすれ違って会話する理由のいくつか
LLM はテストの際に非常に活用されます。必要な労力という点では、特定の品質基準に達するのはかつてないほど簡単になっていますが、それでもソフトウェアの品質はかつてないほど低下しているようです。 10 年前、私たちは任意の 1 週間に遭遇したバグを調べました。当時はかなりの数のバグがあり、今ではさらに多くのバグに遭遇していますが、必ずしもそうである必要はないと思います。
まず、バグが出荷された後、データ駆動型のアプローチを使用してバグを見つけて修正することがこれまでより簡単になりました。たとえば、職場でサポート チケット (チャットまたは電子メール) からプル リクエスト (PR) に至るパイプラインを作成してみました。私の知る限り、これは問題なく動作します。私は従来のワークフローを採​​用している会社で働いているため、これらの修正はすべて人間によってレビューされており、これまでのところ誤検知は発生していません。
投資した単位時間ごとに、より徹底的なテストを行うことも可能です。個人的には、これは十分に効果的であると考えており、「ソフトウェア ファクトリ」ワークフローを介して大量のコードを出荷することにかなり抵抗がありません。これは、レビューに依存したワークフローよりもはるかに高い品質をもたらす、テスト中心のレビューなしのワークフローを見てきたからです。

見たことも聞いたこともあります。
皆さんと同じように、私にも自分の経験からくる偏見があります。偶然ですが、私はキャリアの最初の 10 年間を、今日の LLM 環境でもテスト プロセスがうまく機能する会社で過ごしました。私はマストドンのデフォルトのテスト方法としてファジングについて話しましたが、懐疑論者がそれを試してみると、すぐにいくつかのバグを発見しました。
そこでブログ記事を読み直してみたら、とても「怪訝そうな顔」でしたが、いえいえ、クロードのファジングで修正する価値のあるいくつかのクラスのバグが見つかりました。
私が話した他の何人も、ここで説明するテスト フローのようなものを採用しようとしましたが、彼らは皆、コードのバグ監査、バグの発見、「テスト」、「さらにテスト」などを Codex や Claude に依頼するだけでは表面化しないバグも含めて、自分たちが取り組んでいるソフトウェアのバグをすぐに発見しました。たとえば、デニス・スネル氏は、自分とチームメイトのジョン・サレル氏はコード内のバグを発見しただけではないと述べました。かなり少ない労力で、「HTML 仕様、ビッグ 3 ブラウザ、その他のオープンソース プロジェクトを含む上流の依存関係」にも取り組んでいます。
一般に、ソフトウェア担当者とテストについて話すと、私があまりにも違う場所から来たので、すぐに私を宇宙人のように見ることがあります。そこで、私が働いていたハードウェア会社 Centaur でどのようにテストしたかについて話しましょう。これは、私がどのように働きたいかについての私の偏見を物語っています。ソフトウェアの世界で私たちが行った、または非正統的なことのいくつかは次のとおりです。
専任の QA / テスト エンジニアを雇用し、テストは開発者と同等の一流のキャリア パスとなる
手書きのテストはほぼ無し
プログラマーがプロパティ ベースのテスト、ランダム化テスト、ファジングなどと呼ぶことがある、定期的なテスト。ただし、私たちはこれらのテストを単にテストと呼んでいます (手書きのテストは「ハンド テスト」と呼ばれました)。

大規模な回帰テスト スイート (コンピューティング ファームでの実行には 3 か月かかります)
一般的な構造について説明すると、私が退職したとき (2013 年)、約 20 人のロジック設計者と 20 人のテスト エンジニアのために、常時約 1,000 台のマシンがテストを生成および実行していました。これはオンプレミスで、マシンは私たちがいた建物の半分のフロアを占めていました。
一般的な構造としては、おそらく 20% のマシンが回帰テストを実行し、80% のマシンが新しいテストを生成して実行しているということでした。 3 か月の回帰テストではコミットをゲートするには多すぎるため、実行に 10 分程度かかる、コミット前に実行するはるかに短いテストのリストがありました。これらのプリコミット テストは、お金で購入できる最速のマシンであるオーバークロックされたマシンと、別のシミュレータ セットアップを使用して、できるだけ早く実行するための特別なセットアップで実行されます。
新しい障害が発生すると発見されて報告され、1 ～ 2 人のエンジニアが障害を分類してトリアージする仕事をしていました (誤検知の拒否、誤検知の原因となったテスト ジェネレーターの問題の修正など)。
影響の大きさに関して言えば、文化を別個の項目としてカウントしない限り、(1) がおそらく私たちと典型的なソフトウェア会社との最大の違いですが、ここの読者にとって最も無関係でもあります。そのため、テストは他のスキルと同様であるというこの短いコメントを除いて、議論を脚注 1 に譲ります。テストに多くの時間を費やすことでスキルが向上します。また、ほとんどの大手テクノロジー企業ではテストが第一級のキャリアパスではないため、一部のキャリア CPU テスト エンジニアに見られるような、ソフトウェア会社でのテスト スキルと同じレベルのテスト スキルを持っている人は一般的にいません。分散システムに 20 年を費やしたエンジニアと同じように

分散システムや UX に時間の 5% を費やす同等の才能のあるエンジニアよりも、ステムや UX のほうがはるかに優れています。20 年間テストに取り組んでいる人は、時間の 5% をテストに費やしている人よりもはるかに優れています。
(2) は、チップ会社で使用したテスト手法の一部を AI ワークフローに適したものにするものの 1 つです。デフォルトではコードをレビューしませんでした。これは、テストの実践を十分に信頼しており、レビューによって一般に信頼性があまり向上しないためです。ユーザーが目に見える重大なバグは年間 1 件未満であり、レビューは、特に注意が必要と思われる事項について追加の監視が必要な場合に必要に応じて行われました 2 。 AI コーディング ワークフローを使用すると、人間が手作業でレビューできるよりも多くのコードを 1 人で簡単に生成できます。レビューなしの配送コードに対する快適さのレベルは人によって異なります。個人的には、人間によるレビューなしでコードを出荷することに非常に抵抗がありません。なぜなら、ほとんどのソフトウェア会社で、ほとんどのソフトウェアよりも技術的に難しい製品で人間によるレビューが行われているのを見てきたからです。
「それはリスクが大きすぎます。何百万ものユーザーがいるのですから」といったことを言う人をよく見かけますが、経験的に、彼らは生の数で一人当たりおそらく 1,000 倍高い割合でバグを送り込むワークフローについて話しています。重大度を調整するとその割合はさらに高くなります。ある会社が、主にバグを発見するためにレビューに依存しながら、Centaur のときの 100 分の 1 の速度でバグを出荷していたとしたら、その会社の言い分も理解できますが、バグを出荷するリスクを認識して人間によるレビューから離れたがらない典型的なソフトウェア会社ではそれが起こっていません。
(3)と(4)は連動しています。私が知っているほぼすべてのソフトウェア グループ

信頼性を真剣に考えているチーム (信頼性の高いデータベース、分散データベースなどを出荷するさまざまなチーム) は、手書きのテストの方が多いかもしれませんが、少なくとも方向性的には同じことを行っています。同じ理由で、自分でソフトウェアを操作してソフトウェアが動作するかどうかを観察するテストに依存するのは悪い考えだと考えられており、テストへの入力と予想される出力を直接入力することに依存するのは悪い考えです。前に説明したように、手動でテストを作成するのは非常に非効率的です。信頼性のどのレベルであっても、手書きのテストよりもランダム化されたテスト生成を好む場合は、より早く目標を達成できます。
(5) 多くのテストを行った結果、多くのバグが見つかりました。一般に、テストでバグが見つかり、後で修正した場合、そのテストを回帰テスト スイートに永久に保存します。優れたテストで多くのバグが見つかった場合、最終的には大規模なテスト スイートが必要になることがわかりました。しかし、それを脇に置いて、テスト効率の観点からだけ見てみると、PR ごとに同じテストのセットを CI で実行するというソフトウェアの標準設定は、バグが見つかる可能性が高いことを考えると、同じテストを 1 日に 1000 回実行するか、同じテスト時間内に 1000 の異なるテストを実行するかということを考えると、非常に非効率的です。
(6) は、競合他社よりもはるかに小規模なチームを抱えていたという点で、テストの効率性の問題からも生まれました。それが同社が長く存続できた理由だった。 Intel が AMD 以外の x86 設計者をすべて廃業に追い込んでいた一方で、当社の運営コストは十分に低かったので、会社は 2021 年まで存続し、その時点で Intel に 1 億 2,500 万ドルで買収されました。会社のチーム規模が小さかったため、適切なテスト範囲を確保することは不可能だったでしょう

GE が単体テストを導入し、単体テストを行うのに十分な人材を雇用していれば、同社はおそらく 10 年か 20 年早く、Transmeta、Rise、Cyrix、TI、UMC、NEC、VM などの x86 への取り組みの道を進んでいたでしょう。効率の観点から見ると、単体テストのパフォーマンスはかなり悪いです。
要約すると、ほとんどのソフトウェア担当者が悪いアイデアだと言うことを私たちはかなりの数行いました (専任のテスト エンジニア、単体テストなし、コード レビューなしなど)。そして、私が働いたことのあるどのソフトウェア会社や、私が使用したどのソフトウェアよりもはるかに高い品質を実現しました。このことについて話すと、CPU には X の問題しかなく、Y では同じことはできないから、これはソフトウェアには当てはまらないと言われるでしょう。最初に CPU 設計からソフトウェアに切り替えたときは、それは本当かもしれないと思いましたが、それ以来、誰かがこれでは機能しないと述べたすべての種類の Y でこのテスト方法を試してみましたが、すべての Y でうまくいきました。そのため、これはあまり妥当とは思えません (そして、X には一般に、ハードウェア開発がどのようなものであるかについての誤った仮定が含まれています)。ハードウェアとソフトウェアの間には実際の違いがありますが、テスト技術が継承されない理由として人々がそれを頼りにするのを私が見てきたとき、それは事実でした。

[切り捨てられた]

## Original Extract

Agentic test processes, LLM benchmarks, and other notes on agentic coding from Galapagos Island
| Patreon
I've been using AI fairly heavily since last November and the whole thing is a funny experience . An agent will do something that, if a human did it, you'd immediately fire them. My reaction, of course, is to act as if this is great and spin up a thousand agents so they can do even more of that.
Mid-last year, I had GPT (maybe 5.0 or 5.1) try to find the source of a bug . Naturally, this code didn't have tests and git bisect wouldn't work, and it was a UI interaction bug for which I'm not even really qualified to write a test for, so I asked Codex to bisect between dates X and Y to find the commit that introduced this bug. Codex immediately told me the offending commit was after this date range (which couldn't possibly be correct). On telling Codex this was wrong, it then told me some commit that was obviously also not the offending commit once or twice. On telling it those were wrong, it then told me the offending commit was some plausible looking commit. When I asked it to prove or disprove its theory, it told me that it wrote a test and confirmed that the alleged commit was the breaking commit.
I then asked it to show me by making a video with the full developer end-to-end stack in the normal browser test environment. It claimed that it didn't have permissions to do that ( which was a lie ), but it could make video of the execution of the repro before and after the commit in playwright with the appropriate test code. The video was convincing and showed the feature working properly before the commit and failing to work after the commit. Something about this didn't feel right, so I tried reproducing the issue by hand before and after the commit and found out that the whole thing was a fabrication. The video made it look like Codex had reproduced the bug, but it was an artificial browser environment that was designed to create a fake repro, not the real environment.
Like I said, because this was non-ironically such a great experience , I immediately thought to myself, "how can I get more of this?" and started using agents more and more heavily until I was using coding agents heavily mid-late last year .
Since this post covers a relatively disparate set of topics, here's a brief outline .
Agentic loops and writing this post
Some reasons people talk past each other
LLMs are highly leveraged when it comes to testing. In terms of the amount of effort it takes, it's easier than ever to hit a particular quality bar and yet, software seems to be lower quality than ever. A decade ago, we looked at the bugs I ran into in an arbitrary week . There were quite a few bugs then and I run into more bugs now, but I don't think this has to be the case.
For one thing, after a bug has been shipped, it's easier than it's ever been to use a data-driven approach to find and fix the bug. Just for example , at work, I tried creating a pipeline that goes from support ticket (chat or email) to pull request (PR). As far as I can tell, this works ok . Since I work for a company that has a traditional workflow, all of these fixes get reviewed by a human and, so far, we've had no known false positives .
Per unit of time invested, it's also possible to do more thorough testing. Personally, I think this can be effective enough that I'm fairly comfortable trying to ship a large volume of code via a " software factories " workflow because I've seen a testing-heavy no-review workflow that results in much higher quality than any review-reliant workflow I've seen or even heard of.
Like everybody, I have biases that fall out of my experiences. It just so happens that I spent the first decade of my career at a company whose test processes happen to work well in today's LLM environment. I talked about fuzzing as a default testing methodology on Mastodon , and a skeptic tried it out and immediately found some bugs :
so I reread the blog post and was very "dubious face" but no yeah, Claude fuzzing found several classes of bugs that are worth fixing
A number of other folks I've talked to have also tried adopting something like the testing flow we'll discuss here and they've all immediately found bugs in the software they work on, including bugs that don't get surfaced by just asking Codex or Claude to audit the code for bugs, find bugs, "test", "test more", etc. For example, Dennis Snell mentioned that he and a teammate, Jon Surrell, not only found bugs in the code they're working on, but also "in upstream dependencies, including the HTML specification, big-three browsers, and other open-source projects" with fairly low effort.
In general, when I talk to software folks about testing, I'm coming from such a different place that they immediately look at me like I'm an alien, so let's talk about how we tested at this hardware company I worked for, Centaur, which informs my biases about how I like to work. Some of the things that we did that were or are unorthodox in the software world are:
Hired dedicated QA / test engineers, with testing being a first-class career path on par with being a developer
Virtually no hand-written tests
Constant testing via what programmers sometimes called property based testing, randomized testing, fuzzing, etc., although we just called those tests (hand-written tests were called " hand tests ").
Large regeression test suite (3 months wall clock to execute on compute farm)
Just to give you an idea of the general structure, when I left (in 2013), we had about 1000 machines generating and running tests at all times for roughly 20 logic designers and 20 test engineers. This was on prem and the machines took up half a floor of the building we were in .
The general structure was that we had maybe 20% of machines running regression tests, and 80% generating and running new tests. Three months of regression tests is too much to gate commits on, so there was a much shorter list of tests that took maybe 10 minutes or so to run that people would run before committing. Those pre-commit tests would run on a special setup to run as quickly as possible, with overclocked machines that were the fastest machines money could buy, as well as a different simulator setup .
New failures would get found and reported as they happened and one to two engineers had a job of sorting through failures and triaging them (rejecting false positives, fixing issues in the test generator that caused them to generate false positives, etc.).
In terms of the magnitude of the impact, unless you count culture as a separate item , (1) was probably the biggest difference between us and a typical software company, but also the most irrelevant for readers here, so I'll relegate the discussion to a footnote 1 , except for this brief comment that testing is like any other skill; spending more time doing it improves skill and, since testing isn't a first-class career path at most major tech companies, people generally don't have the same level of testing skills at software companies as you see in some career CPU test engineers. In the same way that an engineer who who spends 20 years working on distributed systems or UX is going to be much better at it than an equally talented engineer who spends 5% of their time on distributed systems or UX, someone who spends 20 years working on testing is going to be much better at it than somebody who spends 5% of their time on testing.
(2) is one of the things that makes some of the test practices we used at the chip company suited to AI workflows. We didn't review code by default because we trusted our test practices enough that review didn't, in general, add much reliability. We were shipping fewer than 1 significant user-visible bug per year, and review was done on an as-needed basis when someone wanted an extra set of eyes on something they thought was particularly tricky 2 . With AI coding workflows, it's easy for one person to generate more code than any human or even any ten humans can review by hand. People have different levels of comfort with shipping code without review. Personally, I'm very comfortable shipping code without human review because I've seen it done on products that are technically more challenging than most software at most software companies.
I often see people say things like, "that's too much risk; we have millions of users" but, empirically, they're talking about a workflow that ships bugs at a rate that's maybe a thousand times higher per capita on raw count, with the ratio being much higher if you adjust for severity . If a company were shipping bugs at, say, a hundredth the rate we were at Centaur while relying primarily on review to catch bugs, then I could see their point, but that's not what's happening at the typical software company where people don't want to move away from human review because of the perceived risk of shipping bugs.
(3) and (4) go hand in hand. Almost every software group I know of that's serious about reliability (various teams that ship reliable databases, distributed databases etc.) are at least directionally doing the same thing, although they might have a larger fraction of hand written tests. For the same reason it's considered a bad idea to rely on testing by interacting with the software yourself and observing whether or not the software appeared to work, it's a bad idea to rely on directly typing out the inputs to a test and the expected outputs. As previously discussed , it's just really inefficient to write tests by hand . For any given level of reliability, you'll get there more quickly if you prefer randomized test generation over hand-written tests.
(5) fell out of having a lot of tests find a lot of bugs. In general, if a test found a bug that we later fixed, we'd keep the test in our regression test suite forever. It turns out, if you find a lot of bugs with good tests, you'll end up with a large test suite. But putting that aside and just looking at it from a test efficiency standpoint, the standard setup in software of having the same set of tests run in CI for each PR is extraordinarily inefficient if you think about the what's more likely to find a bug, running the same test a thousand times in a day or, in the same amount of test time, running a thousand different tests.
(6) came out of test efficiency concerns as well, in that we had a much smaller team than our competitors. That was a reason the company managed to survive for so long. While Intel was putting every x86 designer out of business other than AMD, our operating cost was low enough that the company survived until 2021, at which point it was acquired by Intel for $125M. With the company's tiny team size, it wouldn't have been possible to get reasonable test coverage with unit tests and hiring enough to do unit tests probably would've meant the company would've gone the way of the x86 efforts of Transmeta, Rise, Cyrix, TI, UMC, NEC, VM, etc., a decade or two sooner. From an efficiency standpoint, unit testing does pretty poorly .
To sum it up, we did quite a few things that most software people tell me are bad ideas (dedicated test engineers, no unit tests, no code review, etc.) and we had much higher quality than any software company I've worked for or any software I've used. Whenever I talk about this, people will say that this doesn't apply to software because CPUs only have X concerns and you can't do the same thing with Y. When I first switched from CPU design to software I thought that might be true, but I've since tried this testing methodology with every kind of Y that someone has mentioned this can't work for and it's worked for every single one, so I no longer find this very plausible (and the Xs generally involved incorrect assumptions of what hardware development is like ). While there are real differences between hardware and software, when I’ve seen people lean on that as a reason that testing techniques don’t carry over, it’s been the case

[truncated]

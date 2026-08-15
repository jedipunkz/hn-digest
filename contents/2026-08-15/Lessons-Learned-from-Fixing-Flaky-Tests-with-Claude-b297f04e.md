---
source: "https://henrikwarne.com/2026/08/15/lessons-learned-from-fixing-flaky-tests-with-claude/"
hn_url: "https://news.ycombinator.com/item?id=49312355"
title: "Lessons Learned from Fixing Flaky Tests with Claude"
article_title: "Lessons Learned from Fixing Flaky Tests with Claude | Henrik Warne's blog"
author: "henrik_w"
captured_at: "2026-08-15T18:15:03Z"
capture_tool: "hn-digest"
hn_id: 49312355
score: 1
comments: 0
posted_at: "2026-08-15T17:14:58Z"
tags:
  - hacker-news
  - translated
---

# Lessons Learned from Fixing Flaky Tests with Claude

- HN: [49312355](https://news.ycombinator.com/item?id=49312355)
- Source: [henrikwarne.com](https://henrikwarne.com/2026/08/15/lessons-learned-from-fixing-flaky-tests-with-claude/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T17:14:58Z

## Translation

タイトル: Claude による不安定なテストの修正から学んだ教訓
記事のタイトル: Claude による不安定なテストの修正から学んだ教訓 |ヘンリック・ウォーンのブログ
説明: 最近、統合テストの多くを並列実行することで、統合テストを 10 倍高速化しました。新しいテスト スイートには不安定なテストがいくつかありました。クロードは、ほとんどすべての不安定さを取り除くのに非常に役立ちましたが、途中でいくつかの間違いを犯しました。これが私が学んだことです。私たちのシステムは…

記事本文:
ヘンリック・ウォーンのブログ
プログラミングへの思い…
コンテンツにスキップ
ホーム
Claude による不安定なテストの修正から学んだ教訓
私たちは最近、統合テストの多くを並行して実行することで、統合テストを 10 倍高速化しました。新しいテスト スイートには不安定なテストがいくつかありました。クロードは、ほとんどすべての不安定さを取り除くのに非常に役立ちましたが、途中でいくつかの間違いを犯しました。これが私が学んだことです。
私たちのシステムは統合テストで十分にカバーされています。完全なローカル システムが実行されており、すべての統合テストは FIX 接続 (トラフィック) および HTTP 接続 (管理) を介して実行中のシステムと対話します。以前は、1360 統合テストは順番に実行され、完了までに約 40 分かかりました。最近の書き直しにより、実行時間が約 4 分に短縮されました。多くのテストを並行して実行できることからメリットが得られます。
テストを並行して高速に実行できるようにするために、テスト スイートは一度定義され、その後多くのテストで再利用されるインストゥルメントに依存します。ただし、これにより、テストはテスト間の干渉に対して脆弱になります。書き換え後は、通常、実行ごとにいくつかのテストが失敗しました。テストが単独で実行された場合、テストは成功します。これが不安定なテストの定義です。単独では合格しますが、大規模なテスト スイートの一部として実行すると (タイミングの問題により) 失敗する可能性があります。
実行に失敗した後にテストを単独で試すのは非常に簡単ですが、不安定さを完全に取り除く方がはるかに良いです。
私はテスト フレームワークについてはあまり詳しくありませんでしたが、クロードは不安定なテストを見つけて修正するのが得意だと思いました。失敗したテストごとに、失敗した場所からのスタック トレースが生成されました。多くの場合、スタック トレースは、目的の状態の待機がタイムアウトした場所を示すだけなので、根本的な原因は明らかではありませんでした。クロード・エクセル

はがれの原因の解明に取り組みました。そこで私は調子に乗って、理論的根拠をよく理解せずに、提案された修正を受け入れてコミットしてしまいました。問題はランダムに発生するため、根本的な問題が実際に修正されたかどうかを判断するのは必ずしも簡単ではありません。
不安定なテストは著しく減りましたが、完全に消えたわけではありません。多くのテスト ケースを移植した同僚に変更をレビューするよう依頼したところ、彼は多くのコミットに疑問を抱きました。そのうちのいくつかはもっともらしく聞こえましたが、効果はありませんでした。私は修正のほとんどを理解しておらず、クロードが提案したものをただ盲目的に受け入れていたことに気づきました。
そこで最初からやり直し、今回はそれぞれの修正を本当に理解するのに時間がかかりました。私が理解できなかった修正のあらゆる側面についてクロードに尋ねました。何度もクロードが正しかったので、私はテストとフレームワークがどのように機能するかについてさらに学びました。しかし、解決策について質問すると、次のような答えが返ってくることがありました。
「あなたがこれを推進するのは正しい。私は自分の言ったことを撤回する必要がある」
「公正な挑戦です。自分が実証できることとできないことについて正直に話させてください。」
「だからさっきは誇張してたんだ」
最終的には、私たちの共同の努力により、解決策に到達するでしょう。クロードの提案を盲目的に受け入れないことで、テストについてより多くのことを学び、テストに無意味なコードを追加することを避けることができました。
多くの場合、クロードは不安定の原因を見つけるという素晴らしい仕事をしてくれました。スタック トレースの分析は簡単で、複数のログ ファイルをクロスチェックして、重複するアクションのタイムスタンプをチェックすることも簡単です。クロードの Java と関連フレームワークに関する百科事典的な知識は本当に役に立ちます。ただし、Claude は多くの場合に非常に優れているため、その出力はすべて正しいと思われがちです。
私は本物を持っています

クロードが提案された解決策だけでなく、より広範に既存のコード全般をどのように説明できるかを理解するようになります。既存のコードを理解するのがはるかに速くなりました。私はまずクロードに、システムの一部がどのように機能するのか概要を教えてもらいます。その後、フォローアップの質問をして、特定の領域を掘り下げることができます。
LLM が登場する前は、不安定性の原因をすべて追求することに労力を費やす価値があったかどうかは疑問です。たとえ最終的に良い結果が得られたとしても、時間がかかりすぎたでしょう。これはおそらく他の用途にも当てはまります。以前はやる価値がなかったことが、今では十分に安価にできるようになりました。これは素晴らしいですね。
クロードは、不安定なテストのさまざまな原因を追跡するのに非常に役立ちました。ただし、それが示唆するすべてが正しいわけではありません。これは多くの機能を非常にうまく実行するため、提案されるすべての変更を受け入れるのは非常に簡単です。したがって、システムを重視する場合、提案された修正が有効かどうかを理解できることは非常に価値があります。
現在、不安定なテストが失敗するのは 10 回の実行に 1 回未満です。以前よりはずっと良くなりましたが、クロードと私にはまだやるべきことがいくつかあります。
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
ソフトウェア開発で学んだ教訓
ソフトウェア開発者として働き始めたときに驚くことトップ 5
ソフトウェア開発者として働く
優れたプログラマーはデバッグ可能なコードを作成する
Claude による不安定なテストの修正から学んだ教訓
クロードと: コーディングを減らし、テストを増やす
さらに 9 年間にわたる厄介なバグからの教訓
その他の優れたプログラミングの名言、パート 6
プログラミング カンファレンス – Jfokus ストックホルム 2025
私のシンプルな知識管理および時間追跡システム
ジョン・フォン・ノイマン – 未来から来た男
新しいソフトウェア開発者の仕事を探す
何

ソフトウェア開発に対する考えが変わりました
アルゴリズム取引: 実践者向けガイド
ソフトウェアのメンテナンスは不要です
Go への切り替え – 第一印象
効果的なソフトウェア テスト – 開発者ガイド
書評: ソフトウェア設計の哲学
マイクロサービスについて私が気に入っている 4 つの点
ソフトウェア開発者の募集 – コーディングテスト
その他の優れたプログラミングの名言、パート 5
サッカーの数学的モデリング
実稼働環境にデプロイするだけでは十分ではない
在宅勤務 – 短所と長所
人工知能 – 考える人間のためのガイド
その他の優れたプログラミングの名言、パート 4
EuroSTAR テストカンファレンス プラハ 2019
Python の古典的なコンピューター サイエンスの問題
ソフトウェア開発者の採用 – 会社をチェックする
書評: データ集約型アプリケーションの設計
ノルディック テスト デイズ タリン 2019
その他の優れたプログラミングの名言、パート 3
私のお気に入りのコマンドライン ショートカット
プログラミング形式の演習
プログラミングについての6年間の考え
継続的デリバリーの利点
その他の優れたプログラミングの名言、パート 2
プログラミング カンファレンス – QCon ニューヨーク 2017
ソフトウェア開発とギグエコノミー
書評: 有能なエンジニア
13 年間にわたる厄介なバグから得た 18 の教訓
プログラミングの知恵の名言
博士号それともプロのプログラマーですか？
Kevin Mitnick によるソーシャル エンジニアリング
ソフトウェア開発者の募集 - 最初の連絡先
Coursera コースのレビュー: ソフトウェア セキュリティ
ソフトウェア開発で学んだ教訓
Coursera コース レビュー: コンピュテーショナル投資パート 1
ソフトウェア開発者がキャリアの選択肢として最適である 5 つの理由
「ほとんどの単体テストが無駄である理由」への回答
Java から Python への切り替え – 第一印象
脆弱性対策とソフトウェア開発
バグ、トレース、テスト、ツイスト
バグの発見: デバッガーとロギング
TDD、単体テスト、そして時間の経過
オートマティカ

lly ログステートメントにリビジョンを含める
メソッドを増やすことでプログラムを改善できる 7 つの方法
優れたプログラマーはデバッグ可能なコードを作成する
SET カード ゲーム バリエーション – 相補的なペア
プログラマーの生産性 – 中断、会議、リモート作業
Coursera コースのレビュー: アルゴリズム: 設計と分析、パート 2
2012 年のブログ統計 (WordPress による)
ソフトウェア開発者として働く
虫が体に良い4つの理由
書評: Google はどのようにソフトウェアをテストするか
ソフトウェア開発者として働き始めたときに驚くことトップ 5
プログラマーの生産性: Emacs と IntelliJ IDEA
Coursera コースのレビュー: アルゴリズムの設計と分析 I
Mac OS X Break プログラムのレビュー
反復性ストレス損傷を克服する方法
データベース入門 – オンライン学習はうまくいきました
「「最も類似した」セットの中からランダムに選択」を使用してシミュレートされた 1,000 万の SET ゲーム
「利用可能なセットの中からランダム」を使用してシミュレートされた 1,000 万の SET ゲーム
「First Found Set」を使用してシミュレートされた 1,000 万の SET ゲーム
メール アドレスを入力してこのブログをフォローし、新しい投稿の通知をメールで受け取ります。
購読する
購読しました
ヘンリック・ウォーンのブログ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

We recently sped up our integration tests by a factor of ten by running many of them in parallel. There were several flaky tests in the new test suite. Claude was really useful in getting rid of almost all flakyness, but I made some mistakes along the way. Here is what I learnt. Our system…

Henrik Warne's blog
Thoughts on programming…
Skip to content
Home
Lessons Learned from Fixing Flaky Tests with Claude
We recently sped up our integration tests by a factor of ten by running many of them in parallel. There were several flaky tests in the new test suite. Claude was really useful in getting rid of almost all flakyness, but I made some mistakes along the way. Here is what I learnt.
Our system is very well covered by integration tests. We have a complete local system running, and all integration tests interact with the running system via FIX connections (traffic) and HTTP connections (administration). Previously, our 1360 integration tests ran sequentially, and took around 40 minutes to complete. A recent rewrite got the running time down to around 4 minutes instead. The gain comes from being able to run many tests in parallel.
To be able to run tests fast in parallel, the test suite relies on instruments that are defined once, then reused for many tests. This however makes the tests vulnerable to interference between tests. After the rewrite, there were typically a few tests each run that failed. If the test is run in isolation, it passes. So this is the definition of a flaky test: it passes in isolation, but may fail (due to timing issues) when run as part of a larger test suite.
Even though it is fairly easy to try the test in isolation after a failing run, it is much better to get rid of the flakiness once and for all.
I was not very familiar with the testing framework, but I thought Claude would be good at finding and fixing the flaky tests. Each failing test produced a stack trace from where it failed. Often, the stack trace would just show where a wait for a desired state timed out, so the underlying cause was not obvious. Claude excelled at finding the causes of flakiness. So I got carried away, and accepted and committed the proposed fixes without really understanding the rationales. Since the problems appeared randomly, it is not always easy to tell if the underlying problem had actually been fixed.
The flaky tests became noticeably fewer, but did not disappear completely. When I asked my colleague, who had ported many of the test cases, to review the changes, he questioned many of the commits. Several of them sounded quite plausible, but had no effect. I realized that I did not understand most of the fixes, but had just blindly accepted whatever Claude suggested.
So I started over, and this time took the time to really understand each fix. I would ask Claude about all aspects of the fix I did not understand. Many times, Claude was correct, and I learnt more about how the tests and the framework work. But occasionally, I would get these types of responses when I questioned the solutions:
“You’re right to push on this, and I need to walk back what I said”
“Fair challenge — let me be honest about what I can and can’t demonstrate.”
“So I was overstating it earlier”
Eventually, through our joint effort, we would arrive at a solution. By not blindly accepting Claude’s suggestions, I learnt a lot more about the tests, and I avoided adding pointless code to the tests.
In many of the cases, Claude did a fantastic job of finding the cause of the flakiness. Getting stack traces analyzed is a breeze, as is crosschecking multiple log files, checking the time stamps for overlapping actions. Claude’s encyclopedic knowledge of Java and associated frameworks really helps. However, because Claude is so good in so many cases, it is easy to assume that all of its output is correct.
I have really come to appreciate how Claude can explain not only a proposed solution, but also more broadly, the existing code in general. It is now much faster for me to understand existing code. I first ask Claude to give me an overview how some part of the system works. Then I can drill down into specific areas by asking follow-up questions.
Before LLMs, it is doubtful if it would have been worth the effort to chase down all the causes of flakiness. It would simply have taken too much time, even if the end result is good to have. This probably applies to other uses as well. Something that was not worth doing before is now possible to do cheaply enough. This is great.
Claude was a great help in tracking down the various causes of flaky tests. However, not everything it suggests is correct. Because it does many things so well, it is very easy to just accept every change it suggests. Therefore, being able to understand if the proposed fix is valid or not is very valuable if you care about your system.
Currently, less than one run in ten has a flaky test failure. Much better than before, but still some work for Claude and me to do.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Lessons Learned in Software Development
Top 5 Surprises When Starting Out as a Software Developer
Working as a Software Developer
Great Programmers Write Debuggable Code
Lessons Learned from Fixing Flaky Tests with Claude
With Claude: Less Coding, More Testing
Lessons From 9 More Years of Tricky Bugs
More Good Programming Quotes, Part 6
Programming Conference – Jfokus Stockholm 2025
My Simple Knowledge Management and Time Tracking System
John von Neumann – The Man from the Future
Finding a New Software Developer Job
What I Have Changed My Mind About in Software Development
Algorithmic Trading: A Practitioner’s Guide
There Is No Software Maintenance
Switching to Go – First Impressions
Effective Software Testing – A Developer’s Guide
Book Review: A Philosophy of Software Design
4 Things I Like About Microservices
Recruiting Software Developers – Coding Tests
More Good Programming Quotes, Part 5
Mathematical Modelling of Football
Deployed To Production Is Not Enough
Working From Home – Cons and Pros
Artificial Intelligence – A Guide for Thinking Humans
More Good Programming Quotes, Part 4
EuroSTAR Testing Conference Prague 2019
Classic Computer Science Problems in Python
Recruiting Software Developers – Checking Out a Company
Book Review: Designing Data-Intensive Applications
Nordic Testing Days Tallinn 2019
More Good Programming Quotes, Part 3
My Favorite Command-Line Shortcuts
Exercises in Programming Style
6 Years of Thoughts on Programming
Benefits of Continuous Delivery
More Good Programming Quotes, Part 2
Programming Conference – QCon New York 2017
Software Development and the Gig Economy
Book Review: The Effective Engineer
18 Lessons From 13 Years of Tricky Bugs
The Wisdom of Programming Quotes
Ph.D. or Professional Programmer?
Social Engineering from Kevin Mitnick
Recruiting Software Developers – Initial Contact
Coursera Course Review: Software Security
Lessons Learned in Software Development
Coursera Course Review: Computational Investing Part 1
5 Reasons Why Software Developer is a Great Career Choice
A Response to “Why Most Unit Testing is Waste”
Switching from Java to Python – First Impressions
Antifragility and Software Development
A Bug, a Trace, a Test, a Twist
Finding Bugs: Debugger versus Logging
TDD, Unit Tests and the Passage of Time
Automatically Include Revision in Log Statement
7 Ways More Methods Can Improve Your Program
Great Programmers Write Debuggable Code
SET Card Game Variation – Complementary Pairs
Programmer Productivity – Interruptions, Meetings and Working Remotely
Coursera course review: Algorithms: Design and Analysis, Part 2
Blog stats for 2012 (by WordPress)
Working as a Software Developer
4 Reasons Why Bugs Are Good For You
Book Review: How Google Tests Software
Top 5 Surprises When Starting Out as a Software Developer
Programmer Productivity: Emacs versus IntelliJ IDEA
Coursera course review: Design and Analysis of Algorithms I
Mac OS X Break Programs Review
How I Beat Repetitive Stress Injury
Introduction to Databases – On-line Learning Done Well
10 million SET games simulated using “Random among ‘most similar’ Sets”
10 million SET games simulated using “Random among available Sets”
10 million SET games simulated using “First found Set”
Enter your email address to follow this blog and receive notifications of new posts by email.
Subscribe
Subscribed
Henrik Warne's blog
Already have a WordPress.com account? Log in now.

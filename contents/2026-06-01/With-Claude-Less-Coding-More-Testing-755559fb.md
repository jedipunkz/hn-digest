---
source: "https://henrikwarne.com/2026/05/31/with-claude-less-coding-more-testing/"
hn_url: "https://news.ycombinator.com/item?id=48345028"
title: "With Claude: Less Coding, More Testing"
article_title: "With Claude: Less Coding, More Testing | Henrik Warne's blog"
author: "ingve"
captured_at: "2026-06-01T00:31:50Z"
capture_tool: "hn-digest"
hn_id: 48345028
score: 20
comments: 1
posted_at: "2026-05-31T11:56:54Z"
tags:
  - hacker-news
  - translated
---

# With Claude: Less Coding, More Testing

- HN: [48345028](https://news.ycombinator.com/item?id=48345028)
- Source: [henrikwarne.com](https://henrikwarne.com/2026/05/31/with-claude-less-coding-more-testing/)
- Score: 20
- Comments: 1
- Posted: 2026-05-31T11:56:54Z

## Translation

タイトル: クロードと: コーディングを減らし、テストを増やす
記事のタイトル: クロードと: コーディングを減らし、テストを増やす |ヘンリック・ウォーンのブログ
説明: クロード コードを数か月間使用して、ソフトウェア開発が自分にとってどのように変わったかに気づきました。私が書くコードはかなり減りましたが、クロードが書いたコードの理解とテストには多くの時間を費やしています。比率は変わりましたが、それでもソフトウェア開発のような感じがします。クロードも使ってます
[切り捨てられた]

記事本文:
ヘンリック・ウォーンのブログ
プログラミングへの思い…
コンテンツにスキップ
ホーム
クロードと: コーディングを減らし、テストを増やす
数か月間 Claude Code を使用して、ソフトウェア開発が自分にとってどのように変わったかに気づきました。私が書くコードはかなり減りましたが、クロードが書いたコードの理解とテストには多くの時間を費やしています。比率は変わりましたが、それでもソフトウェア開発のような感じがします。
また、私が開発したものではない既存のコードをより深く理解するために Claude を使用します。全体として、LLM コーディング エージェントを使用することは私にとって非常にポジティブです。ソフトウェアを作成する喜びを失うことなく、開発の多くの部分をスピードアップします。
エージェントの使用に関しては、私は最先端からは程遠いです。クロードが書いたとしても、私は今でもコードを見ます。時々編集してます。私が開発しているシステムがどのように機能するかを理解することが重要であると考えています。これには、アーキテクチャから実装の詳細に至るまで、さまざまなレベルでの理解が含まれます。最終的に、システムは膨大な量の詳細から構成されます。
機能を追加するとき、追加によってシステムがどのように変更されるかを知り、理解したいと思います。後で細かいことを忘れてしまうかもしれませんが、少なくとも一度はすべてを理解したいと思います。多くの人が、エージェントに与えられた仕様で十分であり、開発者は詳細を気にする必要はないと感じていることは承知しています。しかし、これまでのところ、私も詳細を理解することにこだわっています。これは部分的には、変更に私の名前が載っているので、それを保証したいからです。また、実装には、短い仕様では把握できない詳細が非常に多く含まれているためです。これらの詳細の多くは、ソリューションの動作に影響を与えます。ジョン・サルバティエのエッセイ『現実には驚くほどの詳細がある』を思い出します。
最近では、通常、ne を開始します。

チケットの説明が正しいかどうかをクロードに尋ね、正しい場合は解決策の提案を求めます。たとえ頭の中に解決策があるとしても、私はクロードを特定の解決策に導くことは避けます。もしかしたら、私が思いつかなかったもっと良い方法があるかもしれません。クロードには私の提案にそのまま従うことも望んでいません。
コードが書かれたら、それを理解するために読んでいきます。クロードとはよくやり取りがあります。「この部分は何をするのですか?」なぜここにあるのですか？それはどのように機能しますか？このプロセス中には、大きなことも小さなことも変更する必要がある可能性があります。
自分で定型コードを書いたり、適切な構文を見つけたり、API の正しい使用方法を見つけたりする必要がないのは素晴らしいことです。それは変化のロジックに直接行くようなものです。
私は常に、自分の名前を掲げた変化がうまくいくと自分自身に納得させたいと思っていました。単体テストや統合テストでメイン ロジックが動作していることを確認することとは別に、変更内のコードのすべての行が実行されたことを確認し、ログ メッセージがコンテキスト内で適切に表示されていることを確認し、機能を使用しているときにシステム全体を観察することなどが含まれます。
Claude を使用すると、テストのセットアップが非常に簡単になります。以前は、テストを実行できる適切な環境を構築するには多大な労力が必要でした。自動テストとは別に、いくつかの探索的テストも行うのが好きです。クロードを使用すると、これを促進するために一時的なローカル変更を要求するのが簡単になります。たとえば、一部の処理が午前 0 時にのみ実行される場合、ローカル テスト システムが開始後 1 分でそのロジックを実行できるようにするパッチを入手できます。真夜中まで待つ必要はありません。
AI は学習しないことの言い訳にはなりません。それどころか、特に AI が与える答えを判断できるようになるためには、学習が依然として重要です。おそらく最も驚くべき利点

Claude を使用すると、既存のコード ベースを探索するのに非常に役立ちます。
私は定期的にクロードに、既存のシステムの機能がどのように機能するかを説明してもらいます。通常、回答の品質は非常に高く、コード内の影響を受ける領域をいつでも簡単に確認できます。さらに良いのは、まだわからない部分についてフォローアップの質問を続けることができるという事実です。この目的は常に、システムに関する私自身の知識を高めることです。
現在は異常な時代であり、LLM はソフトウェアの開発方法について大きく変化しています。コーディング エージェントを使用することで、開発者としての時間の使い方が変わったことがわかりました。コーディングに関する付随的な詳細の多くはなくなりましたが、開発中の本質的なロジックにはまだ取り組んでいます。 Claude を使用することで、既存のコード ベースの理解もはるかに容易になりました。
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
クロードと: コーディングを減らし、テストを増やす
さらに 9 年間にわたる厄介なバグからの教訓
その他の優れたプログラミングの名言、パート 6
プログラミング カンファレンス – Jfokus ストックホルム 2025
私のシンプルな知識管理および時間追跡システム
ジョン・フォン・ノイマン – 未来から来た男
新しいソフトウェア開発者の仕事を探す
ソフトウェア開発に関して考えが変わったこと
アルゴリズム取引: 実践者向けガイド
ソフトウェアのメンテナンスは不要です
Go への切り替え – 第一印象
効果的なソフトウェア テスト – 開発者ガイド
書評: ソフトウェア設計の哲学
マイクロサービスについて私が気に入っている 4 つの点
ソフトウェア開発者の募集 – Codin

g テスト
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
ログステートメントにリビジョンを自動的に含める
メソッドを増やすことでプログラムを改善できる 7 つの方法
優れたプログラマーはデバッグ可能なコードを作成する
SET カード ゲーム バリエーション – 相補的なペア
プログラマーの生産性 – 中断、会議、リモート作業
Coursera コースのレビュー: アルゴリズム: 設計と分析、パート 2
2012 年のブログ統計 (WordPress による)
を

ソフトウェア開発者としての活動
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
購読済み
ヘンリック・ウォーンのブログ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。
コメントを書く...
メールアドレス (必須)
お名前 (必須)
ウェブサイト

## Original Extract

Having used Claude Code for a few months now, I have noticed how software development has changed for me. I write a lot less code, but I spend more time understanding and testing the code Claude has written. The proportions have changed, but it still feels like software development. I also use Claud
[truncated]

Henrik Warne's blog
Thoughts on programming…
Skip to content
Home
With Claude: Less Coding, More Testing
Having used Claude Code for a few months now, I have noticed how software development has changed for me. I write a lot less code, but I spend more time understanding and testing the code Claude has written. The proportions have changed, but it still feels like software development.
I also use Claude to better understand existing code that I didn’t develop. On balance, using an LLM coding agent is very positive for me. It speeds up many parts of development, without losing the joy of creating software.
I am far from the bleeding edge when it comes to using agents. I still look at the code, even if Claude wrote it. I sometimes edit it. I believe it is important to understand how the system I am developing works. This includes understanding it on many different levels, from architecture down to implementation details. Ultimately, a system is made of an enormous amounts of details.
When I add a feature, I want to know and understand how the system is changed by my addition. I may forget some of the details later, but I want to have understood them all at least once. I know many people feel that the specification given to the agent should be enough, and that developers no longer need to care about the details. But so far, I am sticking with understanding the details too. This is partly because with my name on the change, I want to be able to vouch for it. It is also because there are so many details in an implementation that can’t be captured in a short specification. Many of those details will affect the behavior of the solution. It reminds me of the essay Reality Has a Surprising Amount of Detail by John Salvatier.
These days, I typically start a new feature by asking Claude if the description in the ticket is correct, and if so, I am asking for a suggested solution. I avoid steering Claude to a given solution, even if I have one in mind. Perhaps there are better ways of doing it that I haven’t thought about. I also don’t want Claude to just go along with whatever I suggest.
When the code is written, I will read through it to understand it. There is often a back and forth with Claude: what does this part do? why is this here? how does that work? There can be both big and small things to change during this process.
It is great not having to write boiler plate code myself, or finding the right syntax, or figuring out the correct way of using an API. It is like going straight to the logic of the change.
I have always wanted to convince myself that the change I am putting my name to will work. Apart from seeing the main logic working with unit tests and integration tests, it means making sure every line of code in the change has been executed, seeing that the log messages look good in context, observing the whole system when using the feature, and so on,
With Claude, it is so easy to get tests set up. Before, it could take a lot of work to create the right environment for tests to be feasible. Apart from automatic tests, I like to do some exploratory tests too. With Claude, it is easy to ask for some temporary local changes to facilitate this. For example, if some processing is only done at midnight, I can get patches to make my local test system execute that logic one minute after starting, instead of having to wait for midnight.
AI is not an excuse for not learning. On the contrary, learning is still important, not least in order to be able to judge the answers the AI is giving. Perhaps the most surprising benefit of using Claude has been how useful it has been in exploring the existing code base.
I regularly ask Claude to explain how features in the existing system works. Usually, the answers are of very high quality, and it is always easy to check the affected areas in the code. Even better is the fact that I can keep asking follow-up questions about parts that are still not clear to me. The purpose of this is always to increase my own knowledge of the system.
These are extraordinary times, with LLMs changing a lot about how software is developed. I have found that using coding agents has changed how I spend my time as a developer. A lot of the incidental details of coding have gone away, while I am still engaged in the essential logic of what I am developing. Using Claude has also made understanding the existing code base a lot easier.
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
Write a Comment...
Email (Required)
Name (Required)
Website

---
source: "https://belderbos.dev/blog/jeff-haemer-agentic-ai-cohort/"
hn_url: "https://news.ycombinator.com/item?id=48567620"
title: "Building an AI Agent in 6 Weeks (and Understanding How They Work)"
article_title: "Building an AI Agent in 6 Weeks (and Finally Understanding How They Work) — Bob Belderbos | Developer & Team Coaching"
author: "lumpa"
captured_at: "2026-06-17T10:38:48Z"
capture_tool: "hn-digest"
hn_id: 48567620
score: 1
comments: 0
posted_at: "2026-06-17T08:52:13Z"
tags:
  - hacker-news
  - translated
---

# Building an AI Agent in 6 Weeks (and Understanding How They Work)

- HN: [48567620](https://news.ycombinator.com/item?id=48567620)
- Source: [belderbos.dev](https://belderbos.dev/blog/jeff-haemer-agentic-ai-cohort/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T08:52:13Z

## Translation

タイトル: 6 週間で AI エージェントを構築 (およびその仕組みを理解)
記事のタイトル: 6 週間で AI エージェントを構築 (そしてついにその仕組みを理解) — Bob Belderbos | Bob Belderbos開発者とチームのコーチング
説明: Jeff はソフトウェアに何十年も携わっていましたが、Python は平凡で AI の地図もありませんでした。 6 週間後: エージェントは 3 つのインターフェイス、約 250 のテスト、および実際のメンタル モデルを備えていました。

記事本文:
6 週間で AI エージェントを構築 (そしてついにその仕組みを理解)
エージェント AI を構築しますか?私は、別の API ラッパーではなく、本番環境に対応したエージェントを出荷する 6 週間のコホートを共同実行しています。
Jeff Haemer は、1980 年代初頭にコロラド大学でソフトウェアを教えて以来、ソフトウェアを作成してきました。しかし、彼は Python をブラッシュアップする必要があり、何よりも AI の基礎を身につける必要があると感じました。
彼の言葉を借りれば、AI は「私が知らなかったことの大きな未分化な雲」でした。それを変える時が来ました。
Python Agentic AI コホートを開始してから 6 週間後、彼は数千行のコード、約 250 の単体テスト、100% カバレッジ、および 3 つの動作インターフェイス (Web UI、コマンド ライン、携帯電話と通信する Telegram ボット) を備えたエージェント アプリケーションを完成させました。
彼は突然変異のテストを実行することさえしました。それ以上に、彼はこのプログラムを通じて、エージェントが実際に水面下でどのように働いているかについてのメンタル モデルを開発しました。
ジェフは長いキャリアを積んで入社し、自分に欠けているものを明確に把握していました。
「これは巨大な世界であり、毎日変化しています。そして、多くの点でどこから始めればよいのかさえわかりませんでした。あなたが AI コホートの実施を考えていると話したとき、私は、ぜひそうしてサインアップしてください、と言いました。喜んでまず飛びつきます。」
彼は意図的に期待を低く設定した。 「私は非常に斑点のある背景を持っているので、自分自身がクラッシュして火傷するだろうと予想していました。」彼はたくさんのコースを教えてきました。彼は何事も最初の作業がどうなるかを知っていた。彼はこのグループのベータ テスターとしてサインオンし、何が起こっても熱心に乗り切ることに決めました。
Jeff のインタビュー全文をご覧ください。
コアロジックが最初、インターフェースがその上にあります
エージェントは、リクエストを受け取り、AI モデルを使用してそれを推論し、行動するプログラムです。コホートは最初に推論コアを構築し、次にその周りにインターフェイスをラップします。その構造は厳しい方法でテストされました

e Telegram 週、エージェントをチャット アプリに接続するときに壁にぶつかりました。
「その週の終わりまで来ましたが、あと一歩にも及ばず、私は『分かった、失敗しました』と言いました。そしてあなたは、次の週に続けてくださいと言いました。なぜなら、これらのユニットが互いに独立するようにソフトウェアを設計したからです。私たちはそれを書き直します。」
彼は先に進みました。 Telegram ユニットは独立しているため、未完成のインターフェイスによって残りのビルドがブロックされることはありませんでした。その後、その週の資料を 100 ページ近くから 41 ページまで再構築し、より単純な実用バージョンから始めました。ジェフがそれに戻ったとき、それはうまくいきました。
「素晴らしかったです。スピードバンプにぶつかったとき、それがうまくいきました。それが証明されました。」
ここにはメタ的な教訓があり、適切に設計されたソフトウェアがどのようなものであるかを思い出させてください。最初に作業を行うものを構築し、エッジを交換可能にしておく必要があります。これにより、ベータコホートの崩壊が回避されました。
嘲笑、彼が先延ばしにしていたもの
Jeff が最も苦労した概念はモックでした。これは、テストを高速かつオフラインで実行できるように、実際の外部サービス (API、データベース) を代役に置き換える手法です。
「私にとって最大の苦痛は、モッキングが何を意味するのかを最終的に理解することでした。テスト自動化の Python 専門家である友人に尋ねても、彼は、私もそれをまったく理解していないと言いました。」
彼は第4週か第5週までそれを避けたが、そのときフアンホはそれを交渉不可能にした。
「外部サービスから独立した単体テストを統合テストから実際に分離できるところまで私を導いてくれました。私は pytest について学び、モックについて学び、自分のコードについても学びました。最終的には、物事を別の方法で考えるようになりました。」
Jeff のテスト スイートに対する厳格な取り組みが功を奏しました。最終的には、彼が設定した Debian ボックスにデプロイしたときにカバレッジが外れる機能など、実際の問題を把握するようになりました。

自分で追跡する必要があります。
自動操縦ではなく教育アシスタントとしての AI
ジェフはコードを手書きで書きました。彼は AI を、ファイルを埋めるためのジェネレーターとしてではなく、Python が不足した箇所の家庭教師として意図的に使用しました。
「ここで何をすべきか概念的には理解していますが、短期間でドキュメントを読み進め、すべてのバグに対処できるほど十分な Python の知識がありません。それでは、その方法を教えてください。AI は優れた教育アシスタントに変わります。」
単体テストでは毎週「完了」の定義が示されました。その結果は、彼が推論できる成果物であり、信仰に基づいて信頼しなければならない違いではありません。
ジェフが最も気にしていた勝利はレポではなかった。それはメンタルモデルでした：
「実際にやってみるのと、それについて読んでさえいるのとは違います。もし誰かが説明を書いてくれたとしたら、1週間後にはおぼろげなアイデアが浮かぶでしょう。1か月後にはそれについての記事を読んだけど何も覚えていないでしょう。そして2か月後にはそんなこと聞いたことないと言うでしょう。これは、私には理解できたと思います。」
彼は、野生で見つけたおもちゃのエージェント、つまり幻覚を見せる Wikipedia クローンに対してそれをテストし、入力を検証し、構造化し、モデルに渡し、構造化された出力を取得し、データベースにキャッシュするという動作を正確に追跡することができました。
「私が知らなかったことの大きな未分化な雲」から、AI エージェントがどのように機能するかについての確固たる基礎まで。 Jeff は最終的に GitHub 上に素晴らしいアーティファクトを作成しましたが、本当の勝利は、彼が次の AI プロジェクトに取り組むこのより深い理解でした。
Jeff がこの新しいスキルセットを使って次に何を構築するかを見るのが楽しみです。
ほとんどの AI チュートリアルは「API を呼び出す」で終わります。このコホートは、関数呼び出し、構造化出力、3 つのインターフェイス、Docker、95% 以上のテスト カバレッジを備えたエージェントのデプロイで終了します。ノートではなく、実際のエンジニアリングを 6 週間行います。次の Agentic AI コホートに参加しましょう →
AIができない工学的判断に関する月次メモ

交換しないでください。 Python、Rust、AI。購読して、コード レビューでの Python の禅、実際の PR レビューから得た 19 のレッスンを入手してください。
開発者とチーム向けのエンジニアリング コーチング。パイソン。さび。あい。

## Original Extract

Jeff had decades in software but mediocre Python and no map of AI. Six weeks later: an agent with three interfaces, ~250 tests, and a real mental model.

Building an AI Agent in 6 Weeks (and Finally Understanding How They Work)
Building agentic AI? I co-run a 6-week cohort where you ship a production-ready agent, not another API wrapper.
Jeff Haemer has written software since he was teaching it at the University of Colorado in the early 1980s. But he felt he needed to brush up his Python, and above all get a grounding in AI.
In his words AI was "a big undifferentiated cloud of things I didn't know." Time to change that.
Six weeks into our Python Agentic AI cohort, he ended up with an agentic application with a few thousand lines of code, almost 250 unit tests, 100% coverage, and three working interfaces: a web UI, a command line, and a Telegram bot that talks to his phone.
He even went as far as running mutation testing. More than that, through the program he developed a mental model of how agents actually work under the covers.
Jeff came in with a long career behind him and a clear-eyed view of what he was missing:
"It's a giant world and changing every day, and I didn't even know in many ways where to start. When you told me you were thinking about doing an AI cohort, I said, please do that and sign me up, and I will happily jump in feet first."
He set expectations low on purpose. "I expected myself to crash and burn because I have a very spotty background." He'd taught plenty of courses; he knew how the first run of anything goes. He signed on as the cohort's beta tester and decided to stay enthusiastic through whatever broke.
Watch the full interview with Jeff:
Core logic first, interfaces on top
An agent is a program that takes a request, reasons about it with an AI model, and acts. The cohort builds the reasoning core first, then wraps interfaces around it. That structure got tested the hard way in the Telegram week, when connecting the agent to a chat app hit a wall.
"I got to the end of that week and I wasn't even close, and I said, okay, I failed you. And you said, go on to the next week, because we designed the software so those units are independent of one another. We'll rewrite it."
He moved on. The Telegram unit was independent, so an unfinished interface didn't block the rest of the build. We later rebuilt that week's material from almost 100 pages down to 41, starting from a simpler working version. When Jeff came back to it, it worked.
"It was beautiful. When we hit a speed bump, that worked. It proved itself."
There's a meta lesson here, and a reminder of what well-designed software looks like: build the thing that does the work first, keep the edges replaceable. This saved the beta cohort from going under.
Mocking, the thing he kept putting off
The concept that gave Jeff the most trouble was mocking, the practice of replacing a real outside service (an API, a database) with a stand-in so a test can run fast and offline.
"The biggest pain for me was finally understanding what mocking involves. I even asked a friend of mine, a test automation Python guy, and he said, I never really understood that either."
He avoided it until week four or five, when Juanjo made it non-negotiable:
"You got me to the point where I could actually separate my unit tests, which were independent of outside services, from my integration tests. I learned stuff about pytest, I learned stuff about mocking, and I learned stuff about my own code. I just ended up thinking about things in a different way."
Jeff's rigor with the test suite paid dividends. By the end they were catching real issues, including a function that slipped coverage when he deployed to a Debian box, which he set out to track down on his own.
AI as teaching assistant, not autopilot
Jeff wrote the code by hand. He used AI deliberately, as a tutor for the spots where his Python ran out, not as a generator to fill the file for him.
"I understand conceptually what to do here, but I don't know enough Python to struggle through the docs and work through all the bugs in that short amount of time. So, show me how to do this. AI turns into a great teaching assistant."
Unit tests gave each week a definition of done . The result is an artifact he can reason about, not a diff he has to trust on faith.
The win Jeff cared about most wasn't the repo. It was the mental model :
"It's different having done it than even reading about it. If somebody had written an explanation, a week later I'd have a vague idea, a month later I'd say I read an article about it but don't remember anything, and two months later I'd say I've never heard of that. This, I think I understand."
He tested it against a toy agent he found in the wild, a hallucinating Wikipedia clone, and could trace exactly what it did: validate the input, structure it, hand it to the model, get structured output back, cache it in a database.
From "big undifferentiated cloud of things I didn't know" to a firm grounding in how AI agents work. Jeff ended up with a great artifact on GitHub, but the real win has been this deeper understanding that he will take to his next AI project.
We're excited to see what Jeff will build next with this new skill set.
Most AI tutorials end at "call the API." This cohort ends with a deployed agent: function calling, structured outputs, three interfaces, Docker, 95%+ test coverage. Six weeks of real engineering, not notebooks. Join the next Agentic AI cohort →
Monthly notes on the engineering judgment AI can't replace. Python, Rust, AI. Subscribe and get The Zen of Python in Code Review , 19 lessons from real PR reviews.
Engineering coaching for developers and teams . Python. Rust. AI.

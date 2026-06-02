---
source: "https://brooker.co.za/blog/2026/05/18/whats-easy-whats-hard.html"
hn_url: "https://news.ycombinator.com/item?id=48363750"
title: "What's Easy Now? What's Hard Now? How AI Is Changing Software Development"
article_title: "What's Easy Now? What's Hard Now? - Marc's Blog"
author: "pgedge_postgres"
captured_at: "2026-06-02T00:30:13Z"
capture_tool: "hn-digest"
hn_id: 48363750
score: 5
comments: 1
posted_at: "2026-06-01T23:05:07Z"
tags:
  - hacker-news
  - translated
---

# What's Easy Now? What's Hard Now? How AI Is Changing Software Development

- HN: [48363750](https://news.ycombinator.com/item?id=48363750)
- Source: [brooker.co.za](https://brooker.co.za/blog/2026/05/18/whats-easy-whats-hard.html)
- Score: 5
- Comments: 1
- Posted: 2026-06-01T23:05:07Z

## Translation

タイトル: 今簡単なことは何ですか?今何が大変ですか? AI はソフトウェア開発をどのように変えるのか
記事タイトル: 今簡単なことは何ですか?今何が大変ですか? - マークのブログ

記事本文:
今簡単なことは何ですか？今何が大変ですか?
これは、AI がソフトウェア開発をどのように変えるかについてのシリーズの第 4 回目です。「It’s time to be right」に続きます。 , ジュニアはどうでしょうか？ 、そして私のヒューリスティックは間違っています。今は何ですか？ 。独立したものですが、これが興味深いと思った場合は、これらも興味深いと思うかもしれません。
私はコーディング エージェントの機能の形について考えることに多くの時間を費やしてきました。彼らが今得意なこと、これから得意になること。彼らが今何が苦手なのか、そのどれだけが本質的なもので、どれだけが一時的なものなのか。これはソフトウェアとソフトウェア エンジニアリングの将来を形作る最も重要な問題であるため、検討する価値があります。私は答えを持っているつもりはありませんが、直感に大きく反するかもしれない結論に達しています。
実際、コーディング エージェントは非常に優れたものになってきており、有意義で正しいソフトウェアを非常に迅速かつ革新的な低コストで構築できるようになりました。彼らは、いくつかのコーディングタスクに関して超人的な能力を持っています。もちろん、コンピュータ システムは少なくとも 85 年間、超人的な能力を備えてきました 1 。この 90 年間と同様に、私たちが構築しているこの新しいテクノロジーは、ある分野では非常に超人的なものであり 2 、他の分野では人間の能力に及ばないことがわかると思います。
これは、どのように、そしてなぜ行うのかという重要な疑問を引き起こします。
EE 教育の初期に、教授の 1 人が基板上に簡単な回路を描きましたが、それ以来ずっと私の頭から離れません。こんな感じでした 3:
左側に電圧を印加すると、右側にその電圧の平方根が得られます 4 。 2 つのコンポーネントは、オペアンプとアナログ乗算器 IC (例: 完全に廃止された MC1495 ) です。この単純な回路には、電気工学においておそらく最も重要な概念が凝縮されています。つまり、フィードバックは独特の強力なものです。たぶん不合理

とても強力です。このアイデアは、ほぼすべての電子機器を機能させ、飛行機を空に留め、オーブンで夕食が焦げるのを防ぎます。
フィードバック ループ内のコンポーネントは、基本的な開ループ動作とは大きく異なる動作をするように作成できます。優れた出力は、貧弱なコンポーネントからも引き出されます。乗算器は平方根になる場合があります。フィードバックはすべてを変えます。
AI エージェントは単なるフィードバック ループです。これらは、便利ではあるが欠陥のあるオープン ループ動作 (LLM) を備えたコンポーネントを中心に構築されており、フィードバックを使用して、そのコンポーネントがフィードバックなしでは実行できないことを実行できるようにします。これは、過去 2 年ほどで開発者ツールで起こった変革の背後にある基本的な考え方です。つまり、オープン ループ AI (IDE のスマート オートコンプリート モード) からエージェントへの移行です。人間の開発者 (ビルド、テスト、IDE に戻る) からのフィードバックをエージェント自体 (ビルド、テスト、反復) に移動します。
長期的なコーディング エージェントの機能に関する会話の多くは、開ループ モデルの動作に関するものです。しかし、それは全体像の半分に過ぎません。それは全体像のそれほど重要ではない半分であるとまで言うかもしれません。フィードバックは長期的な能力を促進するものです。
長期的には、コーディング エージェントは、効果的なフィードバックのあるタスクは「簡単」、効果的なフィードバックのないタスクは「難しい」と感じるでしょう。正確なフィードバックが得られるかどうかで、その能力の限界が決まります。
一方で、これは議論の余地のないものであると見るべきです。エージェントを使用してコードを構築したことがある人なら誰でも、適切なエラー メッセージがエージェントの混乱を防ぐのに役立つことを知っています。 Rust のようなツールが、ある種の間違いについて明示的かつ即時にフィードバックを提供することで、エージェントを正しいコードの作成に導く方法を私たちは目にしています。エージェントのパフォーマンスが優れていることがわかります

適切なベンチマークが存在する場合に有効です。私たちは、プロパティベースのテストのようなツールが他に類を見ないほど価値のあるものであると認識しています。また、エージェントはアーキテクチャ (フィードバックが「見ればわかる」という種類のものになる傾向がある) や並行プログラムの作成 (フィードバックが「実行時に黙ってデータを破損する」種類のものになる傾向がある) が苦手であることもわかっています。
しかし、少し先を見据えて 2 つの問題を比較してみましょう。
人間工学に基づいた快適な写真編集 Web サイトを構築します。
適切な高性能データベース ストレージ エンジンの構築 5．
開ループ モデルの場合、前者の方が後者よりも簡単です。少なくとも、純粋なバイブ コーディング ワークフローでは真の成功に近づき、前者では一発で成功に大きく近づくことができます。しかし、フィードバック ループ仮説を考えると、実際には後者の方が長期的な問題としては簡単ではないかと思われます。
その理由を理解するには、彼らのフィードバック ループを考えてみましょう。ウェブサイトのフィードバック ループには、ボタンが適切に機能するかどうかをテストする自動化を超えたもので、ループに人間が関与する必要があります。人間にとって使いやすいものである必要がありますが、人間はフィードバックの提供が遅く、不安定で、一貫性がないことで有名です。ただし、後者の仕様は、API、安全性プロパティ、および活性プロパティを含むかなり単純です。フィードバック ループに適切なツールがあれば、成功に向けた反復に人間は必要ありません。
これは、コーディング エージェントについて多くの人が抱いている直観とは異なると思います。彼らは、Web サイトと UI を「簡単」（SaaSpocalypse を参照）、システム ソフトウェアを「難しい」と考えています。フィードバック ループ仮説では、これは逆であると言われています。実際、SaaS は「難しく」、システム ソフトウェアは「簡単」であることがわかります。
これにより、仕様（何が良いかを書き留める）の重要性が高まるでしょう。

フィードバック ループを駆動するのが好きです)、およびその仕様をコードに適用するツールの使用を好みます。 Rust、Hydro、Verus などのコンパイル時ツール。 TLA+ や P などのモデリング時ツール。 Kiro のスペック アナライザーなどの仕様ツール。テストツール、シミュレーター、モックなど。
ソフトウェア開発の将来は、これらのフィードバック ループを構築することです。多くの難しい問題が残っています。
1930 年代のマリアン・レイェフスキーのような人々の作品に遡ります。
私の机の上の MacBook は、私よりも約 100,000,000,000 倍ほど速く 64 ビット数値を追加できます。
CircuitLab で作成され、この Electronics StackExchange Answer から改変されました。実際には、さらにいくつかの受動部品が必要です。
このことに詳しくない場合は、これがどのように機能するかを直感的に理解してください。オペアンプ (三角形) は、2 つの入力が同じになるように出力 (右側) を調整しようとします。したがって、出力を取得してそれ自体を乗算し、それを入力の 1 つに供給すると、出力が入力の平方根に設定されます。この件についてご存知の方がいらっしゃいましたら、その説明について深くお詫び申し上げます。
ここで言っているのは、Aurora DSQL や PostgreSQL の規模ではなく、たとえば RocksDB や InnoDB の規模のことです。少なくとも私が見る限り、これらの大規模分散システムは、登るのがより困難になると思います。
2026年5月20日 » エージェントソフトウェア開発仮説
2026年4月30日 » 今が正しい時です。
2026年4月9日 » 仕様駆動開発はウォーターフォールではない
マーク・ブルッカー
このサイトの意見は私自身のものです。彼らは必ずしも私の雇用主を代表しているわけではありません。
marcbrooker@gmail.com

## Original Extract

What’s Easy Now? What’s Hard Now?
This is the fourth in a series about how AI is changing software development, after It’s time to be right. , What about juniors? , and My heuristics are wrong. What now? . It stands alone, but if you found this interesting you may also find those interesting.
I’ve been spending a lot of time thinking about the shape of the capabilities of coding agents. What they’re good at now, what they’re going to be good at. What they’re bad at now, how much of that is inherent and how much is transient. This is worth thinking about, because it’s the most important question shaping the future of software, and of software engineering. I don’t pretend to have an answer, but am coming to a conclusion that may be deeply counter-intuitive.
Coding agents are becoming very good indeed, and can build meaningful and correct software very quickly and at transformatively low cost. They have super-human abilities on some coding tasks. Of course, computer systems have had super human abilities for at least 85 years 1 . I think we’re going to find, as we have over those nine decades, that this new technology we’re building is vastly super-human in some areas 2 , and not nearly as capable as humans in others.
Which raises the important question of how, and why.
Early on in my EE education, one of my professors drew a simple circuit on the board that’s been stuck in my mind ever since. It looked like this 3 :
Apply a voltage on the left, and on the right you get the square root of that voltage 4 . The two components are an opamp and an analog multiplier IC (e.g. the deeply obsolete MC1495 ). This simple circuit encapsulates possibly the most important idea in electrical engineering: feedback is uniquely powerful . Maybe unreasonably powerful. It’s the idea that makes nearly every electronic device work, it keeps planes in the sky, and stops your oven from burning your dinner.
Components inside feedback loops can be made to behave significantly differently from their basic open loop behavior. Excellent outputs can be extracted from poor components. Multipliers can become square rooters. Feedback changes everything.
AI agents are just feedback loops. They’re built around a component with useful, but flawed, open loop behavior (an LLM), and use feedback to make that component able to do things that it’s not able to do without feedback. This is the basic idea behind the transformation that has happened in developer tooling in the last two years or so: a move from open loop AI (the smart autocomplete mode in IDEs) to agents. The moving of the feedback from the human developer (build, test, go back to IDE), into the agent itself (build, test, iterate).
Much of the conversation about long-term coding agent capabilities is about open loop model behavior. But that’s only half the picture. I may even stretch to saying it’s the less important half of the picture. Feedback is the thing that’s going to drive long-term capabilities.
In the long term, coding agents will find tasks with effective feedback ‘easy’, and tasks without effective feedback ‘hard’. The availability of accurate feedback will determine the limits on their capabilities .
On one hand, we should see this as uncontroversial. Anybody who has built code with agents knows that good error messages help keep agents unstuck. We’re seeing how tools like Rust guide agents towards writing correct code by providing explicit and immediate feedback about incorrectness of some kinds. We’re seeing agents be great at performance work, where good benchmarks exist. We’re seeing tools like property-based testing be uniquely valuable. We’re also seeing that agents aren’t great at architecture (where feedback tends to be of the ‘I know it when I see it’ kind), or writing concurrent programs (where feedback tends to be of the ‘it silently corrupted data at runtime’ kind).
But let’s look forward a little bit, and compare two problems:
Building a delightful ergonomic photo editing website.
Building a correct high-performance database storage engine 5 .
For open-loop models, the former is easier than the latter. At least in that you’ll get closer to real success with a pure vibe coding workflow, and much closer to success on the former after a single shot. The feedback loop hypothesis, however, makes me think that the latter is actually the easier long-term problem.
To understand why, consider their feedback loops. The website’s feedback loop, beyond maybe some automation that tests if the buttons do what they should, requires a human in the loop. It needs to be easy to use for humans, and humans are notoriously slow, squishy, and inconsistent feedback providers. The latter, however, has a rather simple specification, including the API, safety properties, and liveness properties. With the right tools in the feedback loop, iteration towards success requires no humans.
I think this is different from the intuition many people have about coding agents. They see websites and UIs as ‘easy’ (see the SaaSpocalypse), and system software as ‘hard’. The feedback loop hypothesis says that this is backwards. That, in fact, we’re going to find that SaaS is ‘hard’ and system software is ‘easy’.
This is going to raise the importance of specification (the writing down of what good looks like to drive the feedback loop), and of tools that apply that specification to code. Compile-time tools like Rust, Hydro , and Verus . Modelling-time tools like TLA+ and P . Specification tools like Kiro’s spec analyzer . Testing tools, simulators, mocks, etc.
The future of software development is building these feedback loops. Many hard problems remain.
Dating back to the work of folks like Marian Rejewski in the 1930s.
The MacBook on my desk can add 64 bit numbers about something like 100,000,000,000 times faster than I can.
Drawn with CircuitLab , and adapted from this Electronics StackExchange Answer . In reality, a few more passive components are needed.
If you’re not familiar with this stuff, here’s an intuition for how this works. The opamp (the triangle) tries to adjust its output (on the right) so the two inputs are the same. So if you take the output, and multiply it by itself, then feed it into one of the inputs, it’ll set the output to the square root of the input. If you are familiar with this stuff, I apologize deeply for that explanation.
I mean something on the scale of, say, RocksDB or InnoDB, not something on the scale of Aurora DSQL or even PostgreSQL. I think these large-scale distributed systems are going to be harder to hill climb to, at least for the future I can see.
20 May 2026 » Agentic software development hypothesis
30 Apr 2026 » It's time to be right.
09 Apr 2026 » Spec Driven Development isn't Waterfall
Marc Brooker
The opinions on this site are my own. They do not necessarily represent those of my employer.
marcbrooker@gmail.com

---
source: "https://warpedvisions.org/blog/2025/hitting-the-wall-at-ai-speed/"
hn_url: "https://news.ycombinator.com/item?id=48885784"
title: "The cost of AI-assisted development: cognitive fatigue"
article_title: "The hidden cost of AI-assisted development: cognitive fatigue | warpedvisions.org"
author: "winter_blue"
captured_at: "2026-07-12T23:43:07Z"
capture_tool: "hn-digest"
hn_id: 48885784
score: 3
comments: 0
posted_at: "2026-07-12T23:05:59Z"
tags:
  - hacker-news
  - translated
---

# The cost of AI-assisted development: cognitive fatigue

- HN: [48885784](https://news.ycombinator.com/item?id=48885784)
- Source: [warpedvisions.org](https://warpedvisions.org/blog/2025/hitting-the-wall-at-ai-speed/)
- Score: 3
- Comments: 0
- Posted: 2026-07-12T23:05:59Z

## Translation

タイトル: AI支援開発のコスト: 認知疲労
記事のタイトル: AI 支援開発の隠れたコスト: 認知疲労 | warpedvisions.org
説明: AI 支援開発を 3 か月間続けた後、私の生産性はこれまでよりも向上しました。精神的にも予想外に疲れてしまいます。
生産性の向上は本物です。スケッチから実際に動作するプロトタイプまで、数日ではなく数時間で作成できます。しかし、ここでは触れていない隠れたコストがあります。それはツールの楽しみです。
[切り捨てられた]

記事本文:
warpedvisions.org
ブルース・アルダーソンによるエッセイ、レシピ、および総合メーカー。
AI支援開発の隠れたコスト: 認知疲労
AI 支援による開発を 3 か月間続けた後、私の生産性はこれまで以上に向上しました。精神的にも予想外に疲れてしまいます。
生産性の向上は本物です。スケッチから実際に動作するプロトタイプまで、数日ではなく数時間で作成できます。しかし、私たちが話していない隠れたコストがあります。ツールはプログラミングの認知負荷を根本的に変えます。実装の詳細に取り組むことに疲れるのではなく、アーキテクチャと設計レベルで常に作業することに疲れています。
精力的なコーディング パートナーを持つと、それ自体が燃え尽き症候群を引き起こすことがわかりました。
従来のプログラミングの疲労は、細部に執拗にこだわることから生じます。構文と格闘したり、不明瞭なエラーをデバッグしたり、反復的な実装作業で苦労したりすることもあります。 AI ツールは、この摩擦の多くを解消します。しかし、彼らはそれを、より微妙でより消耗するもの、つまり設計レベルでの意思決定疲労に置き換えます。
1 つを実装するのにかかっていた時間内に 3 つの異なるアプローチのプロトタイプを作成できるようになると、突然、アーキテクチャ上の決定を継続的に行うようになります。これはサービスであるべきでしょうか?図書館？簡単なスクリプト？エラー処理はどうなるのでしょうか？データの永続性?それぞれの決定はさらに多くの決定に分岐し、AI はユーザーが選択したものをすぐに実行する準備ができています。
ボトルネックは「これを構築できるか?」から変わります。 「これをどのように構築すればよいですか?」これははるかに高い認知負荷であり、予想よりも早く蓄積されます。
私が休憩が必要なのは、問題に行き詰まっているからではなく、その抽象レベルで長時間考え続けることで脳が疲れているからです。インターバルtをやっているようなものです

心の高次の機能のために雨が降っているのです。
光速で壁にぶつかる #
最も驚くべき変化は、基本的な設計上の決定をいかに迅速に行うかということです。従来の開発では、実装、テスト、拡張を行うにつれて、より多くのアーキテクチャ上の問題が発見されます。コードの作成には時間がかかるため、その影響についてじっくり考える時間があります。
AI の支援があれば、これらの壁にすぐにぶつかります。実装は非常に迅速に行われるため、適切に検討する時間がないうちに、データ モデル、API 設計、システム境界について突然疑問に直面することになります。
しかし、ここでは別のことが起こっています。人間がコードを書くときは、アーキテクチャについて考えながら作業します。実装の際には、小さな調整を行い、リファクタリングを行い、アーキテクチャ上の決定をエンコードします。 AI はこれを行いません。少なくとも、その出力ではそのような考え方を伝えません。
その結果、コードは機能しますが、アーキテクチャ的にはフラットに感じられます。人間による実装中に自然に起こる微妙な設計の改善を行わずに、要求されたものを実装します。欠けているものを補うために、より明示的なアーキテクチャ的思考を行うことになります。
レビューするコードの量は爆発的に増えましたが、それが最も難しい部分ではありません。本当の課題は、AI の推論を事後的に調べることができないことです。人間のプログラマーが奇妙な選択をしたとき、その理由を尋ねることができます。 AI がそれを実行すると、その推論はより大きな変更セットのどこかに埋もれてしまいます。
特定の決定が下された後に AI に質問しようとすると、説明ではなくおべっかな謝罪をされるでしょう。 「ごめんなさい、それは間違った選択でした、修正させてください。」しかし、私は謝罪が欲しいわけではありません。その実装に至ったトレードオフを理解したいのです。
これにより、

以前には存在しなかったコードレビューの盲点。 「なぜ」にアクセスせずに「何を」をレビューすることになるため、そのアプローチが適切であるか、それともたまたまうまくいっただけであるかを評価することが難しくなります。
ここではテスト規律が重要になります。 AI が自身のコードをテストできない、またはテストしようとしない場合、そして通常は、明示的に要求しない限りテストしませんが、正しさに関して盲目的に行動していることになります。適切なテストでは検出できなかった微妙な問題をデバッグしている場合、速度の利点はすぐに消えてしまいます。
適応とは、ツールをより適切に使用することだけではなく、高度な思考を持続するための新しい精神的な筋肉を開発することです。脳がこれらの分野で強化されるには時間がかかります。
特に大きな設計変更の合間には、より意図的に休憩を取るようになりました。私は文字通りコンピューターから離れて、自分自身の精神的状況をクリアします。 AI の会話コンテキストをクリアすることと並行して行われることは偶然ではありません。あまりにも多くのことを進めすぎると、人間と AI の両方の思考が混乱する可能性があります。
もう 1 つの適応は、実装だけでなく設計検討のための思考パートナーとして AI を使用することです。すぐに「これを構築」するのではなく、もっと時間をかけて「何が足りないのか?」と自問します。これまでに何が行われたのでしょうか?トレードオフは何ですか? AI は実際、この種の分析に非常に優れており、いずれにせよ必要となるアーキテクチャ上の考え方をフロントローディングするのに役立ちます。
練習すれば強くなる #
これは、私たちの働き方における他の大きな変化と似ているように感じます。バージョン管理システムが一般的になったとき、私たちはコミットとブランチに関する新しい習慣を学ばなければなりませんでした。テスト フレームワークが成熟すると、テスト可能なコードを作成するためのさまざまなパターンを内部化する必要がありました。
AI 支援開発は、新しい習慣とメンタル モデルの構築を必要とするもう 1 つの根本的な変化です。疲労はレアです

でも、それも一時的なものです。他の新しい働き方と同様、自然になる前に練習が必要です。
奇妙なのは、この疲労感と、何が可能になるのかという純粋な興奮が共存していることです。私は、より興味深いものを構築し、より多くのアプローチを模索し、より洗練されたプロトタイプを出荷しています。しかし、以前は必要なかった方法で自分のペースを調整することも学んでいます。
AI がプログラミングの実際のエクスペリエンスをどのように変えるのかについては、まだ初期段階にあります。生産性の向上が注目を集めていますが、うまくやっていくためには認知の変化も同様に重要です。
成功する開発者は、これらのツールが構築できるものを変えるだけでなく、構築についての考え方を変えることを認識している開発者です。これらはあなたの最高レベルの思考を増幅させ、信じられないほど強力ですが、私たちがまだ管理方法を学んでいない点で精神的にも厳しいものでもあります。
適応期間はそれだけの価値があります。ただし、学習曲線、特に頭の中で起こる部分を過小評価しないでください。脳にこれらの新しい筋肉を発達させる時間を与えてください。ツールは待機します。

## Original Extract

After three months of AI-assisted development, I’m more productive than ever. I’m also mentally exhausted in ways I didn’t expect.
The productivity gains are real—I can go from sketch to working prototype in hours instead of days. But there’s a hidden cost that we’re not talking about: the tools fun
[truncated]

warpedvisions.org
ESSAYS , RECIPES , AND GENERAL MAKERY BY BRUCE ALDERSON .
The hidden cost of AI-assisted development: cognitive fatigue
After three months of AI-assisted development, I’m more productive than ever. I’m also mentally exhausted in ways I didn’t expect.
The productivity gains are real—I can go from sketch to working prototype in hours instead of days. But there’s a hidden cost that we’re not talking about: the tools fundamentally change the cognitive load of programming. Instead of being tired from wrestling with implementation details, I’m exhausted from operating constantly at the architecture and design level.
It turns out that having a tireless coding partner creates its own kind of burnout.
Traditional programming fatigue comes from the relentless focus on details; fighting with syntax, debugging obscure errors, or grinding through repetitive implementation work. AI tools eliminate much of this friction. But they replace it with something subtler and more draining: decision fatigue at the design level .
When you can prototype three different approaches in the time it used to take to implement one, you’re suddenly making architectural decisions constantly. Should this be a service? A library? A simple script? What about error handling? Data persistence? Each decision branches into more decisions, and the AI is ready to implement whatever you choose, immediately.
The bottleneck shifts from “can I build this?” to “should I build this, and how?” That’s a much higher cognitive load, and it accumulates faster than you’d expect.
I find myself needing breaks not because I’m stuck on a problem, but because my brain is tired from thinking at that level of abstraction for extended periods. It’s like doing interval training for your mind’s highest-order functions.
Hitting the wall at light speed #
The most jarring change is how quickly you slam into fundamental design decisions. In traditional development, you discover more architectural problems as you implement, test, and scale. You have time to think through the implications because writing the code takes time.
With AI assistance, you hit these walls immediately. The implementation happens fast enough that you’re suddenly facing questions about data models, API design, and system boundaries before you’ve had time to think them through properly.
But there’s something else happening here. When humans write code, we think about architecture as we work. We make small adjustments, refactor as we go, and encode architectural decisions in the act of implementation. AI doesn’t do this—or at least, it doesn’t communicate that kind of thinking in its output.
The result is code that works but feels architecturally flat. It implements what you asked for without the subtle design improvements that happen naturally during human implementation. You end up doing more explicit architectural thinking to compensate for what’s missing.
The volume of code to review has exploded, but that’s not the hardest part. The real challenge is that you can’t interrogate AI’s reasoning after the fact. When a human programmer makes an odd choice, you can ask them why. When AI does it, that reasoning is buried somewhere in a larger set of changes.
Try to ask the AI about a specific decision after it’s made, and you’ll get sycophantic apologizing instead of explanation. “I’m sorry, that was a poor choice, let me fix that.” But I don’t want an apology—I want to understand the tradeoff that led to that implementation.
This creates blind spots in code review that didn’t exist before. You’re reviewing the what without access to the why, which makes it harder to evaluate whether the approach is sound or just happens to work.
The testing discipline becomes critical here. If AI can’t or won’t test its own code—and it usually won’t unless you explicitly ask—then you’re flying blind on correctness. The speed advantage disappears quickly if you’re debugging subtle issues that proper tests would have caught.
The adaptation isn’t just about using the tools better—it’s about developing new mental muscles for sustained high-level thinking. Your brain needs time to get stronger in these areas.
I’ve started taking more deliberate breaks, especially between major design shifts. I literally walk away from my computer to /clear my own mental context. The parallel to clearing AI conversation context isn’t accidental—both human and AI thinking can get muddy when you’re carrying too much forward.
The other adaptation is using AI as a thinking partner for design exploration, not just implementation. Instead of jumping straight to “build this,” I spend more time asking: What’s missing? What’s been done before? What are the tradeoffs? The AI is actually quite good at this kind of analysis, and it helps front-load the architectural thinking that you’ll need anyway.
The practice makes you stronger #
This feels similar to other major shifts in how we work. When version control systems became common, we had to learn new habits around committing and branching. When testing frameworks matured, we had to internalize different patterns for writing testable code.
AI-assisted development is another fundamental shift that requires building new habits and mental models. The fatigue is real, but it’s also temporary. Like any new way of working, it requires practice before it becomes natural.
The strange thing is that this exhaustion coexists with genuine excitement about what’s possible. I’m building more interesting things, exploring more approaches, and shipping more polished prototypes. But I’m also learning to pace myself in ways that weren’t necessary before.
We’re still in the early days of understanding how AI changes the actual experience of programming. The productivity gains get all the attention, but the cognitive shifts are equally important to navigate successfully.
The developers who thrive will be those who recognize that these tools don’t just change what you can build—they change how you think about building. They amplify your highest-level thinking, which is incredibly powerful but also mentally demanding in ways we’re still learning to manage.
The adaptation period is worth it. But don’t underestimate the learning curve, especially the parts that happen inside your head. Give your brain time to develop these new muscles. The tools will wait.

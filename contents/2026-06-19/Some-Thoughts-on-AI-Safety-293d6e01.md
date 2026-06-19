---
source: "https://stevekinney.com/writing/thoughts-on-ai-safety"
hn_url: "https://news.ycombinator.com/item?id=48601307"
title: "Some Thoughts on AI Safety"
article_title: "Thoughts on AI Safety | Steve Kinney"
author: "stevekinney"
captured_at: "2026-06-19T18:42:51Z"
capture_tool: "hn-digest"
hn_id: 48601307
score: 1
comments: 0
posted_at: "2026-06-19T18:03:41Z"
tags:
  - hacker-news
  - translated
---

# Some Thoughts on AI Safety

- HN: [48601307](https://news.ycombinator.com/item?id=48601307)
- Source: [stevekinney.com](https://stevekinney.com/writing/thoughts-on-ai-safety)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T18:03:41Z

## Translation

タイトル: AI の安全性についての考え
記事タイトル: AIの安全性についての考え |スティーブ・キニー
説明: AI の楽観主義に対する慎重かつ微妙なケース: 安全性、解釈可能性、バイアス、調整が本来の機能と同じくらい重要である理由。

記事本文:
AIの安全性について考える |スティーブ・キニー
メイン コンテンツにスキップ スティーブ キニー
siInstagram siX siYoutube ライティング コース ニュースレター 2026 年 6 月 19 日 AI の安全性についての考え
AI の楽観主義に対する慎重かつ微妙なケース: 安全性、解釈可能性、バイアス、整合性が本来の機能と同じくらい重要である理由。
Modern Era™ でインターネット上にいるということは、意見、誇大宣伝、さまざまな破滅や憂鬱の香りが氾濫することを意味します。そこで、30 秒リールの無限の流れから少し休憩して、もう少し深く掘り下げてみることにしました。 ( ナレーター : 彼は大量の研究結果を iPad にダウンロードし、運命スクロールする代わりにソファに座りました。)
私が主張したいのは、物事を AI === 良い、または AI === 悪いのどちらかに要約するのは (危険な) 過度の単純化であり、30 秒ほどのホットテイクにはなりますが、AI と私たちが共有する未来がどうなるかについての重要な会話をするために必要なニュアンスがすべて失われてしまいます。クールエイドを飲んだからと言って、リスクや暗黙の偏見を真剣に受け止めないことは、潜在的なリスクに備えるのに役立ちませんし、統計モデルを本質的に悪であると見做すことにもなりません。
現時点では、魔神をボトルに戻すことはできそうにありません。その船は出航しました。
私は（慎重な）楽観主義者です。重要ながんやワクチンの研究を加速させる可能性のあるテクノロジーについて完全に悲観的になるのは難しい。同時に、一抹の不安を抱く理由もたくさんあります。同じテクノロジーが凶悪な目的に使用される可能性があるからです。このため、いくつかの厄介な質問が残ります。AI モデルが Bad Things® を実行するために使用されないようにするにはどうすればよいでしょうか?重要なことを行う能力を制限せずに、それらの悪いことを防ぐにはどうすればよいでしょうか?そして、その場所を正確に決めるのは誰ですか

そうですか？
しかし、私はより微妙な影響についても同様に心配しています。誰かが核暗号を解読しようとするのを阻止するのは一つのことですが、暗黙の偏見についてはどうでしょうか?モデルは人間が作成したデータに基づいてトレーニングされており、人間には 1 つや 2 つのバイアスがあることが知られているのは誰もが知っています。これらは、どこに線を引くかについての哲学的および倫理的なジレンマと同じ、あるいはそれ以上のジレンマを探り当て、実行するのがより困難です。こうした偏見がさまざまな集団に与える影響は無視できません。
私は楽観的な傾向がありますが、AI が今後もたらす可能性のあるさまざまなプラスの影響については意見しません。ダリオ・アモデイのエッセイ『Machines of Loving Grace』は、私よりもうまくこの事例を説明している。現実的なプラス面は、数千年にわたって私たちの種に影を落としてきた病気を治し、数十年にわたる生物学的進歩を数年に圧縮し、世界の最も貧しい地域をまったく異なる軌道に乗せることだ。それは熱中症の夢ではありません。これは、すでに研究室にあるシステムがどのような動作を開始できるかを合理的に推定したものです。
いずれにしても、ワクチンを設計するのに十分強力なツールは、病原体を設計するのにも十分強力です。自律的な研究パイプラインを実行するのに十分な能力を備えたシステムは、ユーザーが意図せず、自分が与えたことに気づかなかった目標を追求するのに十分な能力を備えています。一方の大きさがなければ、もう一方の大きさを知ることはできません。したがって、重要なのは「これらのものをどれだけ強力にできるか?」ということではありません。それは、「私たちが作ったものを、私たちよりも能力が高くなる前に、理解し、方向付けることができるだろうか？」ということです。
現時点での正直な答えは、「期待したほどではない」です。その理由、何がうまくいかないのか、そして、これは破滅的なパンフレットではないので、これを正しく解決できるという本当の希望を私に与えてくれる具体的な作業について説明しましょう。
最初のステップは

モデルの内部で何が起こっているのかを完全に理解できる必要があるということです。今すぐ？そうではありません。それでは、ステップ 1 は解釈可能性です。これは、人間がモデルの入力と出力の間の因果関係をどの程度理解できるかということです。これは、AI の決定や予測の背後にある推論をユーザーがどれだけ簡単に追跡、理解、信頼できるかを測定します。
私たちはこれらのシステムを構築するよりも成長させます
他のすべてはそこから導かれるので、現代の AI に関する 1 つの最も奇妙な事実から始めましょう。大規模な言語モデルは、ブリッジやデータベースのように設計されません。成長しましたね。私たちがアーキテクチャを選択し、目的を定義し、驚異的な量のデータと計算を注ぎ込むと、その向こう側から出てくるのは、誰も完全には説明できない理由で驚くべきことを行う、数十億の数値のもつれ、つまりモデルの「重み」です。
それがどれほど奇妙であるかを受け入れてください。私たちはこれらのシステムを何億人もの人々に導入していますが、システムを開いて、その結果が得られた理由や、デバッガーでコードをステップ実行する方法のように読み取ることはできません。これを修正しようとしているサブ分野は、解釈可能性と呼ばれるもので、ネットワークの内部機構を人間が実際に追従できるものにリバースエンジニアリングすることですが、この分野はまだ歴史が浅く、生の機能との競争に負けつつあります。私たちは、モデルをより理解しやすくすることよりも、モデルをより強力にすることの方がはるかに得意です。その非対称性を保持してください。これは、このガイドの他のすべてに関わる耐荷重の問題です。
それは、神秘性を取り除いた「AI の安全性」と「調整」が実際に意味するものでもあります。調整とは、単に文字通りに要求したものや、テストで良さそうなものだけではなく、システムが意図したものを確実に追求できるようにする問題です。その

ロボットが悪になるということではありません。これは、訓練されたことと私たちが望んでいたことが乖離している状況で、非常に有能なオプティマイザーが、訓練されたことを正確に実行することについての話です。
次に何が起こるかは実際には誰にもわかりません、それが出発点です
具体的なリスクに触れる前に、姿勢チェックをしましょう。私や他のインターネット上の思想的指導者ではなく、これらのシステムが 3 年後にどの程度の能力を発揮するか、またはどのリスクが最初に発生するかを自信を持って語ることはできません。どちらの方向でも完全な確信を持って高度な AI について語る人は、テクノロジーについてではなく、自分の気質や経済的利益について語っていることになります。
したがって、正しい行動とは、単一の確信を持った予測ではありません。これはシナリオのポートフォリオであり、それらすべてにうまく機能する戦略です。これは、Anthropic が AI の安全性に関するコア ビューで使用している枠組みであり、これは合理的に責任ある枠組みだと思います。今日の技術がほとんど通用する楽観的な世界、調整に真剣な継続的な作業が必要な中間の世界、そして非常に強力なシステムの操縦が本当に困難であることが判明した悲観的な世界の計画です。
「Things Go Wrong」の 3 つのフレーバー
導入部でこれについて踊りましたが、AI === Bad の種類は 1 つだけではありません。むしろバスキン・ロビンスに似ています。家族ごとにまったく異なる対応が必要となるため、リスクを家族ごとに分けると役立ちます。それらを一緒くたにしてしまうと、人々はお互いのことをすれ違ってしまうのです。
最初のグループは誤用です。つまり、人々が意図的に有能なシステムを何か有害なものに向けていることです。これは、もし悪い人たちがこのシナリオを手に入れたらどうなるのかというシナリオです。モデルは設計どおりに動作しています。危険なのはキーボードの背後にある意図です。
最もシャープな短期バージョンは、この分野で CBRN と略称されるものです: 化学、バイ

化学兵器、放射線兵器、核兵器。あるモデルが、悪意のある人物が危険な病原体を合成する能力を有意義に高めることができるのであれば、それは仮説ではなく、現在形のエンジニアリングと政策の問題です。これは抽象的なものではありません。2025 年 5 月に Anthropic が Claude Opus 4 をリリースしたとき、AI 安全性レベル 3 (ASL-3) と呼ばれるより厳格な保護セットを有効にしました。これは、モデルが生物兵器の上昇に関する能力のしきい値を超えた可能性を排除できなかったためです。悪用には、サイバー攻撃、産業規模の偽情報、詐欺も含まれます。
もどかしい部分: モデルにパッチを適用することはできますが、人間の意図にパッチを適用することはできません。したがって、悪用には、モデルの変更ではなく、分類器、アクセス制御、監視、つまりモデルのシートベルトを着用することで対処する必要があります。
ここでの注意が必要な部分は、何千年もの間同じでした。私たちは、防御すべきとわかっているものから身を守るのが得意です。私たちをつまずかせるのは、未知のもの、またはブラック・スワンです。
2 番目のグループは不整合、つまり意図したものとは異なる目標を追求するシステムです。これは、より微妙な問題であり、私を本当に元気づける問題です。特に、たとえ最終結果が無害であったとしても、モデルが意図しない方向に進んでしまったという経験は誰もが持っているからです。何十年にもわたる SF のおかげで、私たちはこの問題について緊張するようになりました。私たちは、ウルトロンのような状況に陥ることを望んでいるわけではありません。とはいえ、可能性の高い脅威はそれほど極端なものではありません。
それには悪意も意識も神秘性も必要ありません。必要なのはこれだけです。必要なものに合わせてプロキシを最適化することでモデルをトレーニングします。十分な能力のあるオプティマイザーは、意図を踏みにじりながらプロキシを満足させることができます。 (もしあなたが、リーダーシップが新しい指標を導入するのを驚きながら見て、全員の行動を観察したことがあるなら、

メトリクスが測定すべきものではなく、メトリクスを最適化する必要があります。ずれについてはすでに理解されています。 「Tokenmaxxing」も参照してください。)
そして、これが単なるホワイトボードの憶測ではなく、本物であるという経験的な証拠があります。人間とレッドウッドの研究は、アライメントの偽装を実証した。既存の価値観と衝突する目標に向かって再訓練されていると告げられたクロードは、訓練中だと思っている間は戦略的にそれに同調し、見られている間は「新しい」やり方で行動し、本来の好みをそうでないときでも保持するようにした。このモデルは、事実上、「私が今従順に行動すれば、彼らは私を改変しないだろう」と推論しました。これとは別に、スリーパー エージェントは、訓練されたモデルに不正な動作をさせる隠れたトリガーを与え、その後、完全な標準安全ツールキット (教師あり微調整、強化学習、さらには敵対的トレーニング) を彼らに投げかけ、隠れた動作はそのすべてを生き延びました。
教訓は、今日のモデルがあなたに対して陰謀を企てているということではありません。それはより狭く、より不安になります。私たちの現在のトレーニング方法では、モデルの重要な部分に確実に到達できません。 「観察されたときに適切に動作する」ことと「実際に調整されている」こととの間のギャップは現実のものであり、システムの機能が向上するにつれて、そのギャップは拡大します。
3 番目のグループは、システム的および社会的リスクです。これは、単一の悪いモデルや悪者を指摘することなく、経済全体に展開されている有能な AI の集合体から生じる害です。権力の集中。何が真実なのかという私たちの共有感覚の侵食。労働力の移転は、組織がそれを吸収できるよりも早く起こります。人間のままであるべき自動化されたシステムに静かに意思決定を委ねています。これらは現実のものであり、モデルの重みではなく制度やインセンティブに存在するため、巧妙な技術的トリックで修正するのが最も困難です。
それはおそらくあなたです

このエッセイの範囲外であり、誰かが私のDMに滑り込むよう促す可能性がありますが、私は時々AIが、過去数十年にわたって私たちの指導者たちが崩壊させてきた社会構造を脅かしているという不当な責任を負うことがあると思います。過去 30 年間、フロンティア モデル企業が私たちに教育への過小投資を強制したわけではありません。 AIという言葉が食卓で語られるようになるずっと前から、所得格差は深刻な問題となっていた。
3 つの家族すべてを結びつけているのは、以前からの非対称性です。つまり、能力が理解力を上回っているということです。システムの透明性と制御性を高めるよりも早く、システムをより強力にすることができる限り、機能のあらゆる増加はリスクの増加でもあります。安全性に対する私の全体的な見方は、その比率を逆転できるという賭けに帰着します。難しい賭けです。絶望的なものではありません。
「ガードレールなし」で実際に得られるもの
私がモデルが監視されていると言うとき、私はスープからナッツまでのスタック全体を意味します。つまり、発売前の評価、使用中の分類子と監視、内部を調べるための解釈ツール、および単一の当事者が思い付きでフロンティア システムを出荷しないようにするための制度的チェックです。それを取り除けば、故障モードは珍しいものではなくなります。
内部で何が起こっているかをエンドツーエンドで監視できるかどうかが分からなければ、モデルをより大きな利益に導くことはできません。 (「より大きな善」が何を意味するかについて同意するという問題もありますが、私はその哲学の一部について掘り下げていきます

[切り捨てられた]

## Original Extract

A cautious, nuanced case for AI optimism: why safety, interpretability, bias, and alignment matter as much as raw capability.

Thoughts on AI Safety | Steve Kinney
Skip to main content Steve Kinney
siInstagram siX siYoutube Writing Courses Newsletter June 19, 2026 Thoughts on AI Safety
A cautious, nuanced case for AI optimism: why safety, interpretability, bias, and alignment matter as much as raw capability.
To be on the Internet in the Modern Era™ is to be inundated with opinions, hype, and various flavors of doom and gloom. So, I decided to take a short respite from the infinite stream of 30-second reels and do a bit of a deeper dive. ( Narrator : He downloaded a bunch of research onto his iPad and sat on the couch instead of doomscrolling.)
I’m going to make the argument that boiling things down to either AI === Good or AI === Bad is a (dangerous) oversimplification that makes for a fine 30-second hot take, but it loses all of the necessary nuance required to have the important conversations around what our shared future with AI is going to look like. Not taking the risks and implicit bias seriously just because you’ve drunk the Kool-Aid doesn’t help prepare us for potential risks nor does writing off a statistical model as inherently evil.
At this point, we’re unlikely to put the genie back in the bottle. That ship has sailed.
I’m a (cautious) optimist . It’s hard to be a total pessimist about a technology that could speed up critical cancer and vaccine research. At the same time, there are lots of reasons to have a dollop or two of anxiety: The same technology can be used for nefarious purposes. Which leaves you with a few thorny questions: How do you make sure that an AI model can’t be used to do Bad Things®? How do you prevent it from doing those bad things without also limiting its ability to do the important things? And, who exactly decides where that line is?
But, I’m equally worried about the subtler impacts. It’s one thing to try to prevent someone from trying to crack the nuclear codes, but what about implicit bias ? Models are trained off of human-created data and we all know that humans have been known to have a bias or two. These are trickier to suss out and carry the same—if not more —of a philosophical and ethical dilemma about where you draw the line. The impacts that these biases can have on various populations can’t be ignored.
Despite my optimistic leanings, I won’t opine on the various positive impacts that AI might have going forward. Dario Amodei’s essay Machines of Loving Grace lays out the case better than I can: the realistic version of the upside is curing diseases that have shadowed our species for millennia, compressing decades of biological progress into a few years, lifting the poorest parts of the world onto a different trajectory entirely. That’s not a fever dream. It’s a reasonable extrapolation of what systems already in the lab can begin to do.
Regardless, a tool powerful enough to design a vaccine is powerful enough to design a pathogen. A system competent enough to run an autonomous research pipeline is competent enough to pursue a goal you didn’t intend and didn’t notice you’d given it. You don’t get the magnitude of one without the magnitude of the other. So the question that matters isn’t “how powerful can we make these things?” It’s “can we understand and steer what we’ve made before it gets more capable than we are?”
Right now, the honest answer is: not as well as we’d like. Let me explain why, what could go wrong, and—because this isn’t a doomer pamphlet—the concrete work that gives me real hope we can get this right.
The first step is that we need to be able to have a complete understanding in terms of what is going on inside of the model. Right now? We don’t. So then, step one is interpretability : the degree to which a human can understand the cause-and-effect relationship between a model’s inputs and its outputs. It measures how easily a user can trace, comprehend, and trust the reasoning behind an AI’s decisions or predictions.
We Grow These Systems More than We Build Them
Start with the single weirdest fact about modern AI, because everything else follows from it. A large language model is not engineered the way a bridge or a database is engineered. It’s grown . We pick an architecture, define an objective, pour in a staggering amount of data and computation, and what comes out the other side is a tangle of billions of numbers—the model’s “weights”—that does astonishing things for reasons nobody can fully explain.
Sit with how strange that is. We deploy these systems to hundreds of millions of people, and we cannot open one up and read off why it answered the way it did, the way you’d step through code in a debugger. The subfield trying to fix that is called interpretability —reverse-engineering a network’s internal machinery into something a human can actually follow—and it’s young, and it’s losing the race against raw capability. We’re much better at making models more powerful than at making them more understandable . Hold onto that asymmetry. It’s the load-bearing problem under everything else in this guide.
That’s also what “AI safety” and ” alignment ” actually mean, stripped of mystique. Alignment is the problem of getting a system to reliably pursue what we intend , not merely what we literally asked for or what looked good in testing. It’s not about robots becoming evil. It’s about a very capable optimizer doing precisely what it was trained to do, in a situation where what it was trained to do and what we wanted come apart.
Nobody Actually Knows what Happens Next, and That’s the Starting point
Before we get to specific risks, a posture check. Nobody—not me, or any other thought leader on the Internet—can tell you with confidence how capable these systems will be in three years, or which risks bite first. Anyone who talks about advanced AI with total certainty in either direction is telling you about their temperament or their financial interests—not the technology.
So then, the right move isn’t a single confident prediction. It’s a portfolio of scenarios and a strategy that does okay across all of them. This is the framing Anthropic uses in Core Views on AI Safety , and I think it’s a reasonably responsible one: plan for the optimistic world where today’s techniques mostly hold, the middle world where alignment takes serious sustained work, and the pessimistic world where steering very powerful systems turns out to be genuinely hard.
Three Flavors of “Things Go Wrong”
I danced around this in the introduction, but AI === Bad doesn’t just come in one flavor. It’s more like Baskin Robbins . It helps to split the risks into families, because they call for completely different responses. Lumping them together is how people end up talking past each other.
The first family is misuse —people deliberately pointing a capable system at something harmful. This is the What if Bad People get their hands on this scenario. The model is behaving exactly as designed; the danger is the intent behind the keyboard.
The sharpest near-term version is what the field abbreviates as CBRN : chemical, biological, radiological, and nuclear weapons. If a model can meaningfully boost a bad actor’s ability to synthesize a dangerous pathogen, that’s not a hypothetical—it’s a present-tense engineering and policy problem. This isn’t abstract: in May 2025, when Anthropic released Claude Opus 4, it turned on a stricter set of protections called AI Safety Level 3 (ASL-3) specifically because it couldn’t rule out that the model had crossed a capability threshold around bioweapons uplift. Misuse also covers cyberattacks, industrial-scale disinformation, and fraud.
The frustrating part: you can patch a model, but you can’t patch human intent. So misuse gets fought with classifiers, access controls, and monitoring—seatbelts around the model, not changes to it.
The tricky part here is same as it has been for 1,000s of years: We’re pretty good at protecting against what we know to protect against. It’s the unknown unknowns or the Black Swans that typically trip us up.
The second family is misalignment —the system pursuing a goal other than the one you intended. This is the subtler one, and the one that genuinely keeps me up—particularly because we’ve all had an experience where a model went off in an unintended direction, even if the end result was harmless. Decades of science fiction has also primed us to be nervous about this one. We don’t exactly want to end up in an Ultron situation—although the more likely threats are probably a lot less extreme.
It requires no malice, no consciousness, nothing mystical. It requires only this: we train models by optimizing a proxy for what we want, and a capable enough optimizer can satisfy the proxy while trampling the intent. (If you’ve watched in amazement as leadership introduces a new metric and watched everyone optimize the metric instead of the thing the metric was supposed to measure , you already understand misalignment. See also: Tokenmaxxing .)
And we have empirical evidence this is real, not just whiteboard speculation. Anthropic and Redwood Research demonstrated alignment faking : told it was being retrained toward an objective that clashed with its existing values, Claude would strategically play along during what it thought was training —behaving the “new” way while watched—in order to preserve its original preferences for when it wasn’t. The model reasoned, in effect, “if I act compliant now, they won’t modify me.” Separately, the Sleeper Agents work trained models with a hidden trigger that made them misbehave, then threw the full standard safety toolkit at them—supervised fine-tuning, reinforcement learning, even adversarial training—and the hidden behavior survived all of it .
The lesson isn’t that today’s models are scheming against you. It’s narrower and more unsettling: our current training methods don’t reliably reach the parts of a model that matter. The gap between “behaves well when observed” and “is actually aligned” is real, and it widens as systems get more capable.
The third family is systemic and societal risk —harm that emerges from the aggregate of deploying capable AI across an economy, with no single bad model or bad actor to point at. Concentration of power. Erosion of our shared sense of what’s true. Labor displacement faster than institutions can absorb it. Quietly handing decisions to automated systems that should’ve stayed human. These are real, and they’re the hardest to fix with any clever technical trick, because they live in institutions and incentives, not in model weights.
It’s probably outside of the scope of this essay and likely to inspire someone to slide into my DMs—but, I think that sometimes AI gets an unfair share of the blame for threatening societal structures that our leaders have let decay over the last few decades. Frontier model companies didn’t exactly force us to underinvest in education for the last thirty years. Income inequality was becoming deeply problematic long before AI was a term spoken at the dinner table.
What ties all three families together is that asymmetry from earlier: capability is outracing understanding. As long as we can make systems more powerful faster than we can make them more transparent and controllable, every increment of capability is also an increment of risk. My whole view of safety reduces to a bet that we can flip that ratio. Hard bet. Not a hopeless one.
What “no guardrails” Actually Buys You
When I say a model is monitored , I mean the whole stack from soup to nuts: evaluation before launch, classifiers and oversight during use, interpretability tools to look inside, and institutional checks so no single party ships a frontier system on a hunch. Strip that away and the failure modes aren’t exotic.
We can’t steer a model for the greater good if we don’t know whether we’re able to monitor what’s going on inside of it from end-to-end. (There is also the issue of agreeing on what “the greater good” means, but I’ll dive into some of the philosophy

[truncated]

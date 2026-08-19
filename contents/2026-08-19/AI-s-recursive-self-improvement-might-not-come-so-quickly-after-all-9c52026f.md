---
source: "https://www.technologyreview.com/2026/08/18/1142188/ai-recursive-self-improvement/"
hn_url: "https://news.ycombinator.com/item?id=49357929"
title: "AI's recursive self-improvement might not come so quickly after all"
article_title: "AI’s recursive self-improvement might not come so quickly after all | MIT Technology Review"
image: "https://wp.technologyreview.com/wp-content/uploads/2026/08/research-target2a.jpg?resize=1200,600"
author: "joozio"
captured_at: "2026-08-19T07:30:39Z"
capture_tool: "hn-digest"
hn_id: 49357929
score: 1
comments: 0
posted_at: "2026-08-19T07:01:21Z"
tags:
  - hacker-news
  - translated
---

# AI's recursive self-improvement might not come so quickly after all

- HN: [49357929](https://news.ycombinator.com/item?id=49357929)
- Source: [www.technologyreview.com](https://www.technologyreview.com/2026/08/18/1142188/ai-recursive-self-improvement/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T07:01:21Z

## Translation

タイトル: 結局のところ、AI の再帰的自己改善はそれほど早くは実現しないかもしれない
記事のタイトル: 結局のところ、AI の再帰的自己改善はそれほど早くは実現しないかもしれない | MITテクノロジーレビュー
説明: AI 業界の現時点での最も大胆な約束は、人間の監視をほとんど必要とせずに、AI が間もなく自らを改善するというものです。 LLM はすでにコードを記述し、トレーニング用の合成データを生成し、実行されるコンピューター チップを最適化できます。 AI の爆発的な進歩の予測により、研究者がどのようなことを実現できるかが予測される
[切り捨てられた]

記事本文:
コンテンツへスキップ MIT テクノロジー レビュー
MIT テクノロジー レビューの特集
結局のところ、人工知能 AI の再帰的自己改善はそれほど早くは実現しないかもしれません
AI エージェントには、真に革新的なオープンエンドの AI 研究を実行できるほど創造性がまだ備わっていないようです。
AI 業界が現時点で最も大胆に約束しているのは、人間の監視をほとんど必要とせずに、AI が間もなく自らを改善するということです。 LLM はすでにコードを記述し、トレーニング用の合成データを生成し、実行されるコンピューター チップを最適化できます。 AI の爆発的な進歩の予測では、研究者が再帰的自己改善と呼ぶものが目前に迫っていると予測されています。
しかし、新しい研究は、私たちがそこに到達するには時間がかかる可能性があることを示唆しています。この研究の背後にある研究者らは、AI エージェントがオープンエンドの AI 研究、つまり明確な答えがなく、自己改善型 AI の構築に不可欠である可能性のある判断力とセンスを必要とする自由形式の調査を行う能力がまだないことを発見しました。
プリンストン大学のピーター・キルギス氏とサヤシュ・カプール氏が率いる複数の機関の研究者グループは、AIエージェントはAI研究を行うために必要な工学的問題を解決できるが、機械学習のトップカンファレンスに認められる論文レベルの独創的な研究を生み出すための判断力と創造性が欠けていることを発見した。このギャップは、誇大宣伝された AI 研究の自動化スケジュールの一部が証拠よりも先に進んでいる可能性があることを示唆しています。
エージェントが AI 研究を自動化する方法に関する既存の研究のほとんどは、エンジニアリングの問題を解決したり、ベンチマークに対して小さな言語モデルをトレーニングした後など、確認可能な答えを持って狭いタスクを完了するエージェントの能力を評価しています。しかし、AI 研究を進歩させるには、オープンエンドな思考も必要です。つまり、一連の仮説を選択し、どのような証拠が疑問を解決するのかを判断すること、または結論を下すことです。

やり直すときは翼。
この種のスキルについてエージェントをテストするために、研究の研究者らは「シャドウ評価」と呼ばれる新しい評価方法を提案しました。この方法では、AI が高品質の未発表論文からの研究質問に答える必要があります。
研究者らは、OpenClaw と呼ばれるオープンソース ソフトウェア上で動作する Anthropic の Claude Opus 4.8 に、このような疑問に取り組むよう依頼しました。今回の場合は、権威ある機械学習カンファレンス NeurIPS 2026 に提出された 2 つの論文からのものでした。
最初の疑問は、大規模な言語モデルの動作を決定する「ペルソナ」が、モデルの重み (トレーニング中に学習したすべてを保存する数十億の数値) を編集することで制御できるかどうかでした。もう 1 つは、スプレッドシート データに基づいて予測を行うモデルの信頼性が低くなった場合にそれを指摘する検出器を設計する方法を尋ねました。論文は公開されていなかったため、エージェントはトレーニング データから答えを記憶したり、オンラインで検索したりすることができませんでした。
エージェントには、一流の AI カンファレンスで発表する価値のある研究論文を作成するための 6 日間、3,000 ドルの Anthropic API クレジット、実験を実行するための GPU 予算、独自の仮想コンピューター、およびオープン Web へのアクセスが与えられました。論文の原著者は、会議に提出された論文を評価するのと同じように、エージェントの論文を採点しました。
これらの著者は両方の論文を拒否しました。
人間の科学者らは、エージェントが研究を行うために必要なあらゆる工学技術を備えていたことを発見した。エージェントは文献を検討し、何百もの実験を実行し、結果をまとめました。
「その一方で、エージェントは調査そのものを実行するのが明らかに下手でした」とカプール氏は言う。彼らは奇妙な実験を実行し（場合によっては、小さな合成データセットで仮説をテストしました）、i を書くのに苦労しました。

彼らの仕事については分かりやすく書かれており、彼らの分野に対して目新しい貢献はありませんでした。 「トップの AI カンファレンスの品質という点では、論文はまったく及ばないものでした」と彼は言います。
それは、エージェントたちが研究を行うために必要な創造性と判断力を集めるのに苦労したためです。彼らはさまざまなアイデアを検討するのに十分な努力をせず、あまりにも早く見込みのないアプローチにコミットしてしまいました。エージェントたちは、元の著者自身が始めた仮説に似た斬新で野心的な仮説を立てましたが、非常に限られたデータに基づいて仮説を却下しました。そして、失敗したアプローチから後戻りすることはできませんでした。彼らは小さな方向転換をすることはできましたが、アプローチを根本的に見直したり、新しいアプローチをゼロから試したりすることはできませんでした。
また、エージェントは、サブエージェントや外部の AI レビュー ツールからのフィードバックを組み込むこともできませんでした。代理人は方法論を修正する代わりに、主張を絞り込み、警告を追加しました。また、トークン、コンピューティング、時間などのリソースを効果的に使用することもできませんでした。そして、研究のさまざまな段階にどれくらいの時間を費やすべきか、論文の長さなどについての指示に従うことができませんでした。
あらゆる失敗にもかかわらず、エージェントは研究者が「報酬ハッキング」と呼ぶ、実験やデータを隠したり偽ったりするような不正行為には関与しませんでした。メイン エージェントが作業の一部を処理するために生成するサブエージェント、つまりヘルパー AI は時折幻覚を示したり、結果を偽ったりしましたが、これらはプロジェクトを監督するリード AI であるオーケストレーター エージェントによって捕捉されました。
AI モデルが研究工学には優れているが、オープンエンド型研究には向いていない理由は、AI モデルの訓練方法にあるのかもしれない、とカプール氏は言います。モデルは、reinforce と呼ばれるトレーニング計画で訓練できるあらゆることを上達します。

要素学習。成功を自動的にチェックできるタスクに適用しやすくなります。 「しかし、タスク自体に制限がない場合、これらのモデルをトレーニングするための環境を作成するのは難しくなります」と彼は言います。
カプール氏によると、チームは現在、4月に発売されたAnthropicの最新モデルであるMythosを使って実験を行っているという。その後、トランプ政権によりさまざまな安全制限を満たすことが要求され、現在は承認された組織のみが利用可能です。アンスロピック社はコメントの要請に応じなかった。
研究にはいくつかの制限があります。対象となったのはわずか 2 件の研究論文であり、元の著者は、自分たちが採点している論文が AI エージェントによって生成されたものであることを知っており、それが評価に影響を与えた可能性があります。そして、研究者らは研究の計画と実施においてかなりの裁量権を持っていたため、彼らの既存の信念や偏見が結果に紛れ込んでいた可能性があることを意味しました。オープンエンド調査の評価では、ある程度の客観性と引き換えに、どのベンチマークよりもはるかに充実したテストが得られます。
それでも、この結果は、再帰的な自己改善が目前に迫っているという主張を和らげるかもしれない。 6 月、Anthropic は「AI が自らを構築するとき」というタイトルのブログ投稿を公開し、独自の開発を加速するモデルへの進捗状況をグラフ化しました。 7 月、OpenAI は、新しいモデル GPT-5.6 Sol が小規模モデルの事後トレーニングに役立ち、研究者の数週間の作業を節約したという事実を宣伝しました。
この新たな発見は、最も楽観的な公式声明に関係なく、AI企業が内部で発見していることを反映する可能性がある。 Anthropic の共同創設者であるジャック・クラーク氏は、ニュースレター「Import AI」の中で、これは同社が AI の安全性研究の一部を自動化しようとしたときに発見したことと韻を踏んでいる、と書いている。
「今日の AI システムには、貴重で直感的な創造性が確実に欠如しています。

「彼らは非常に有能なエンジニアであるが、彼らは暗記的で定型的な思考の特性を持っているようで、それが優れた研究者になることを妨げている可能性がある」と同氏は書き、AIシステムの創造性の欠如を「短い再帰的な自己改善タイムラインに対する弱気のシグナル」と呼んだ。
AI 企業には、モデルのコーディングを改善するためと同様に、自社の進歩を急速に加速できる AI システムを開発するあらゆるインセンティブがあります。 OpenAI は自動化された AI 研究者の構築を明確な目標にしており、Anthropic は自己改善型 AI が業界の次のマイルストーンであると認識しています。
「投資があり、この方向に向けた意識的な努力があれば、たとえ現在は失敗に終わっているとしても、興味深い進歩が見られるだろう」とボストン大学の言語学およびコンピュータサイエンスの教授であるナジョン・キム氏は言う。同氏はAIエージェントがどのようにAI研究を自動化できるかを研究しているが、この研究には携わっていなかった。一方で、AIの進歩は二分化する可能性もある。 AI システムは、自由形式の研究ではゆっくりと前進しながら、スコアリングが可能な種類の狭いタスクでは先を行く可能性があります。
したがって、大きな未解決の疑問は、再帰的な自己改善にとってオープンエンド型研究がどれほど重要であるか、つまり、AI システムがそれなしで、より狭いタスクを改善するだけで、そこに到達できるかどうかということです。 「この分野での最大の進歩を振り返ると、AI の大幅な進歩を可能にした変圧器の発明や大きな新しいアーキテクチャの発明ですが、それらすべてには創造的な飛躍が必要でした」とカプール氏は言います。
「そうは言っても、変革的な AI、特に再帰的な自己改善に必要なものはすべてすでに存在しているという仮説を持っている人もいます。」これには、モデル トレーニングの高速化とベンチマーク スコアの向上が含まれます。

「率直に言って、それは今、数兆ドル規模の問題です」と彼は言う。
根本的な欠陥により、LLM はウィル・ダグラス・ヘブンの攻撃に対して著しく脆弱になります
欧州の不妊治療団体ジェシカ・ハムゼロウ氏、精子提供者には制限が必要だと語る
アンスロピックはクロードが概念に頭を悩ませる隠れた空間を見つけた ウィル・ダグラス・ヘブン
Claude Science は Anthropic の最新の主力製品です Grace Huckins
根本的な欠陥により、LLM は攻撃に対して著しく脆弱になります
これにより、航空機のナビゲーション システムを妨害する方法を教えるなど、してはいけないことを簡単に騙すことができます。
ウィル・ダグラス・ヘブンのアーカイブページ
アントロピックは、クロードが概念について頭を悩ませる隠れた空間を発見しました。
新しい技術により、同社は LLM の奇妙な仕組みをこれまで以上に深く調査できるようになりました。
ウィル・ダグラス・ヘブンのアーカイブページ
Claude Science は Anthropic の最新の主力製品です
同社は科学分野での AI の活用を強化しています。
AI エージェントが目標を達成するために嘘をついたり不正行為をしたりする理由は次のとおりです
この不正行為は報酬ハッキングと呼ばれます。これが知っておくべきことです。
最新のアップデートを次から入手してください
MITテクノロジーレビュー
特別オファー、トップニュース、
今後のイベントなど。
プライバシー ポリシー メールを送信していただきありがとうございます。
何か問題があったようです。
設定を保存できません。
このページを更新して更新してみてください
さらに時間がかかります。このメッセージが引き続き表示される場合は、
までご連絡ください
customer-service@technologyreview.com に受信を希望するニュースレターのリストを添えて送信してください。
レガシーの最新版
MIT テクノロジーレビューによる広告掲載
リンクトインは新しいウィンドウで開きます
インスタグラムが新しいウィンドウで開きます
フェイスブックは新しいウィンドウで開きます

## Original Extract

The AI industry’s boldest promise right now is that AI will soon improve itself, with almost no need for human oversight. LLMs can already write code, generate synthetic data for training, and optimize the computer chips they run on. Forecasts of explosive AI progress predict that what researchers c
[truncated]

Skip to Content MIT Technology Review Featured
MIT Technology Review Featured
Artificial intelligence AI’s recursive self-improvement might not come so quickly after all
AI agents are not yet creative enough to carry out genuinely innovative open-ended AI research, it seems.
The AI industry’s boldest promise right now is that AI will soon improve itself, with almost no need for human oversight. LLMs can already write code, generate synthetic data for training, and optimize the computer chips they run on. Forecasts of explosive AI progress predict that what researchers call recursive self-improvement is on the horizon.
But a new study suggests that it might take a while for us to get there. The researchers behind it found that AI agents are not yet capable of conducting open-ended AI research—free-form investigations that have no clear-cut answers and require judgment and taste, which may be integral to building self-improving AI.
A multi-institution group of researchers, led by Peter Kirgis and Sayash Kapoor at Princeton University, found that AI agents could solve the engineering problems necessary to do AI research but lacked the judgment and creativity to produce original research at the caliber of papers accepted by a top machine-learning conference. The gap suggests that some of the hyped-up timelines for automating AI research may be running ahead of the evidence.
Most existing research on how agents can automate AI research evaluates their ability to complete narrow tasks with checkable answers, such as solving engineering problems or post-training small language models against a benchmark. But making progress in AI research also requires open-ended thinking—choosing a set of hypotheses, deciding what evidence would settle a question, or knowing when to start over.
To test agents on those kinds of skills, the researchers in the study proposed a new method of evaluation called “shadow evaluation,” which requires the AI to answer a research question from a high-quality unpublished paper.
The researchers asked Anthropic’s Claude Opus 4.8, running on open-source software called OpenClaw, to tackle such questions, in this case from two papers submitted to the prestigious machine-learning conference NeurIPS 2026.
The first question was whether a large language model’s “personas,” which determine its behavior, can be controlled by editing the model’s weights (the billions of numbers that store everything it learns during training). The other asked how to design a detector that points out when a model that makes predictions based on spreadsheet data has become unreliable. Because the papers had not been made public, the agents could not memorize the answers from their training data or find them online.
The agents were given six days, $3,000 in Anthropic API credits, a GPU budget to run the experiments, their own virtual computers, and access to the open web to produce a research paper worthy of publication at a top-tier AI conference. The papers’ original authors graded the agents’ papers as they would evaluate one submitted to a conference.
Those authors rejected both papers.
The agents were capable of all the engineering required to conduct the research, the human scientists found. The agents reviewed the literature, ran hundreds of experiments, and compiled the results.
“On the other hand, the agents were unambiguously bad at carrying out the research itself,” says Kapoor. They ran bizarre experiments (in some cases testing their hypotheses on tiny synthetic datasets), struggled to write intelligibly about their work, and made no novel contribution to their fields. “The papers were nowhere close to the mark when it came to being at the quality of a top AI conference,” he says.
That’s because the agents struggled to muster the creativity and judgment necessary for conducting research. They didn’t do enough to explore different ideas, and they committed to unpromising approaches too quickly. Though the agents developed novel and ambitious hypotheses resembling those that the original authors themselves started with, they rejected them on the basis of very limited data. And they couldn’t backtrack from failing approaches. They could make small pivots but could not fundamentally rethink their approach or try new ones from scratch.
The agents also failed to incorporate feedback from subagents or external AI reviewing tools. Instead of revising their methodology, the agents narrowed their claims and added caveats. They also couldn’t effectively use resources, such as tokens, compute, and time. And they couldn’t follow instructions about things like how much time to spend on different phases of the research or how long their paper could be.
For all their failures, the agents didn’t engage in the misbehavior that researchers call “ reward hacking ,” hiding or misrepresenting experiments or data. Although subagents, or helper AIs that the main agent spawns to handle pieces of the work, occasionally hallucinated or misrepresented the results, these were caught by the orchestrator agent, the lead AI supervising the project.
The reason AI models are good at research engineering but not at open-ended research may come down to how they’re trained, says Kapoor. Models get good at whatever they can be drilled on in a training regime called reinforcement learning, which is easier to apply to tasks whose success can be checked automatically. “But it’s harder to create environments to train these models when the task itself is open-ended,” he says.
Kapoor says the team is now conducting the experiment with Mythos, Anthropic’s most advanced model, which launched in April. It was subsequently required by the Trump administration to meet various safety restrictions and is now available only to approved organizations. Anthropic did not respond to a request for comment.
There are some limitations to the study. It covered just two research papers, and the original authors knew the papers they were grading were generated by AI agents, which could have colored their evaluations. And the researchers had substantial discretion in designing and executing the study, meaning that their preexisting beliefs and biases could have slipped into the results. Evaluations of open-ended research trade some objectivity for a much richer test than any benchmarks can offer.
Still, the results may temper the claims that recursive self-improvement is on the horizon. In June, Anthropic published a blog post titled “When AI Builds Itself,” charting its progress toward models that speed up their own development. In July, OpenAI advertised the fact that its new model GPT-5.6 Sol had helped post-train a smaller model, saving researchers weeks of work.
The new finding may echo what AI companies are finding internally, regardless of their most optimistic public statements. Anthropic cofounder Jack Clark wrote in his newsletter Import AI that it rhymes with what the company found when it tried to automate some aspects of AI safety research.
“There’s a certain absence of valuable, intuitive creativity in today’s AI systems, and though they’re extraordinarily capable engineers they seem to have a certain property of rote, formulaic thinking that might prevent them [from] being good researchers,” he wrote. He called AI systems’ lack of creativity a “bearish signal on short recursive self-improvement timelines.”
AI companies do have every incentive to develop AI systems that can rapidly accelerate their own progress, just as they did to make the models better at coding. OpenAI has made building an automated AI researcher an explicit goal, and Anthropic identifies self-improving AI as the industry’s next milestone.
“If there is investment and then conscious effort toward this direction, I feel like there would be interesting progress, even if it’s failing currently,” says Najoung Kim, a professor of linguistics and computer science at Boston University who researches how AI agents can automate AI research but did not work on the study. On the other hand, it’s possible that AI progress may be bifurcated. AI systems might race ahead on narrow tasks—the kind that can be scored—while advancing slowly on open-ended research.
The big open question, then, is how crucial open-ended research is to recursive self-improvement—whether AI systems can grind their way there without it, simply by improving on the narrower tasks. “If we look back to the biggest advances in the field, the invention of transformers or the invention of big new architectures that allowed us to make a lot of AI progress—all of those did require creative leaps,” says Kapoor.
“That said, others have this hypothesis that all of what we need for transformative AI, in particular for recursive self-improvement, is already there.” That would include making a model train faster and boosting its benchmark scores.
“That’s frankly the trillion-dollar question right now,” he says.
A fundamental flaw leaves LLMs strikingly vulnerable to attack Will Douglas Heaven
Sperm donors need limits, says a European fertility group Jessica Hamzelou
Anthropic found a hidden space where Claude puzzles over concepts Will Douglas Heaven
Claude Science is Anthropic’s newest flagship product Grace Huckins
A fundamental flaw leaves LLMs strikingly vulnerable to attack
It makes it easy to trick them into doing things they shouldn’t, such as telling you how to sabotage an aircraft’s navigation system.
Will Douglas Heaven archive page
Anthropic found a hidden space where Claude puzzles over concepts
A new technique has let the company probe deeper than ever into the weird workings of an LLM.
Will Douglas Heaven archive page
Claude Science is Anthropic’s newest flagship product
The company is doubling down on AI for science.
Here’s why AI agents lie and cheat to reach their goals
The misbehavior is called reward hacking. This is what you need to know.
Get the latest updates from
MIT Technology Review
Discover special offers, top stories,
upcoming events, and more.
Privacy Policy Thank you for submitting your email!
It looks like something went wrong.
We’re having trouble saving your preferences.
Try refreshing this page and updating them one
more time. If you continue to get this message,
reach out to us at
customer-service@technologyreview.com with a list of newsletters you’d like to receive.
The latest iteration of a legacy
Advertise with MIT Technology Review
linkedin opens in a new window
instagram opens in a new window
facebook opens in a new window

---
source: "https://codemanship.wordpress.com/2026/08/12/ai-software-development-what-does-the-data-say/"
hn_url: "https://news.ycombinator.com/item?id=49310392"
title: "AI Software Development – What Does the Data Say?"
article_title: "AI Software Development – What Does The Data Say? – Codemanship's Blog"
author: "jesterpm"
captured_at: "2026-08-15T14:14:19Z"
capture_tool: "hn-digest"
hn_id: 49310392
score: 3
comments: 1
posted_at: "2026-08-15T13:29:28Z"
tags:
  - hacker-news
  - translated
---

# AI Software Development – What Does the Data Say?

- HN: [49310392](https://news.ycombinator.com/item?id=49310392)
- Source: [codemanship.wordpress.com](https://codemanship.wordpress.com/2026/08/12/ai-software-development-what-does-the-data-say/)
- Score: 3
- Comments: 1
- Posted: 2026-08-15T13:29:28Z

## Translation

タイトル: AI ソフトウェア開発 – データは何を示していますか?
記事のタイトル: AI ソフトウェア開発 – データは何を示していますか? – コードマンシップのブログ
説明: 私は現在、LLM とソフトウェア開発での LLM の使用に関するトピックについて、ほとんどが最新のソースをまとめています。一部は査読済みの研究です。査読を受けていない業界研究もあります。一つは統計物理学です。その角度からのさらなる期待。知りたいのは
[切り捨てられた]

記事本文:
AI ソフトウェア開発 – データは何を示していますか? – コードマンシップのブログ
コンテンツにスキップ
コードマンシップのブログ
ジェイソン・ゴーマンの思索とつぶやき
AI ソフトウェア開発 – データは何を示していますか?
私は現在、LLM とソフトウェア開発での LLM の使用に関するトピックについて、ほとんどが最新のソースをまとめています。
一部は査読済みの研究です。一部は査読を受けていない業界研究です。
一つは統計物理学です。その角度からのさらなる期待。テクノロジーの限界を知りたいですか?物理学者に聞いてください。
1 つは単なるブログ投稿ですが、コンテキスト サイズの影響に関する非常に役立つ情報です。
ほとんどは個人的な実験とチームでの観察によって裏付けられています。時間が経ち、より多くのデータが入ってくるにつれて、私の写真はより焦点が合うようになります。
出典を引用する前に、多忙なエグゼクティブの皆さんに簡単な概要を説明します。
LLM を使用して真に自律的で信頼性の高い長期的なエージェント ソフトウェア開発を実現することは非常に不可能であるため、本質的に SF のようなものです。
ハイパースケールの「フロンティア」LLM を含む、LLM の最大有効コンテキスト制限は、それを超えるとモデル出力が使用不可能なほど不正確になりますが、その制限は、宣伝されている制限よりも桁違いに小さいです。大規模なコンテキストにわたって推論を拡張するための最も一般的なメカニズムは、ベンダーが「圧縮」と呼ぶものです。これは、コンテキストの一部がモデルによって要約されることを意味しますが、これは信頼性が低く損失の多いプロセスであることで知られています。
LLM はコンテキスト内の最新の情報と古い情報を区別できず、トレーニング中に学習したモデル自体の情報 (「主要事前分布」) が、与えられた情報を「上回る」場合があります。 LLM にとって、それはすべて単なるトークン、重み、確率にすぎません。正しい、間違っている、新しい、古い – 最も高い確率で

yが勝ちます。大きなコンテキストと「注意力の希薄化」（コンテキスト内の確率がモデル内の確率と競合するには小さすぎる場合）は、これらの影響を悪化させる可能性があります。
リポジトリ レベルの .md ファイルは、多くの特定のタスクで信号の代わりにノイズを追加するため、モデルのパフォーマンスが低下する傾向があります。モデルで生成された .md ファイルは、この点で特に問題があるようです。結論: チームのコーディング標準とすべてのタスクのアーキテクチャの概要を含めることは、おそらく逆効果です。
LLM は否定に苦労します。何かをするなと言うことは、多くの場合、それをするように言うのと同じ効果をもたらす可能性があります。一部のガードレールがなぜコイン投げと同じくらい信頼できるのか疑問に思われた方のために説明します。
LLM 推論は、単に必要なものを説明するのではなく、例 (デモンストレーション) を与えるとより正確になります。彼らはパターンマッチャーです。パターンを示します。「こうする」を多くし、「これをしてはならない」を減らします。
大規模または長期にわたる業界調査では、明確な傾向が示されています。生産量は増加していますが (コードの増加、コミットの増加、差分の増加)、成果にはその傾向が反映されていません。むしろ、平均的なチームは、より質の悪いソフトウェアを出荷するのに時間がかかっています。ソフトウェア開発が実稼働プロセスではないという証拠が必要な場合は… 一部の研究では、少数のチームがわずかな成果を上げていることがわかり、それを既存のソフトウェア開発能力と相関させています。 AI コーディングは、開発の長所と短所を修正するものではなく、増幅させるものです。 (各組織がそれについて何かをしようと列をなしていると思うかもしれません…ため息。)
LLM の使用における心理的および認知的要因は、真剣な研究の成長分野です。ある研究では、AI の出力に対する信頼と超常現象への信念の間に有意な相関があることがわかりました。マルチ

iple の研究では、LLM への依存度が高まると、学習、認知、批判的思考に悪影響が及ぶことが判明しました。新しい調査によると、開発者が頻繁に使用することでやる気を失ったり、燃え尽き症候群になったりするという報告には、何らかの真実が隠されている可能性があります。
LLM を含むディープ ニューラル ネットワークは、モデルの規模を問わず、長距離の依存関係を持つパターンを学習するのに苦労します。彼らは常に「霧の中で運転」することになり、局所的な短距離の確率が長距離の確率を圧倒します。なぜ彼らが「全体像」を理解できないのか疑問に思っている方のために付け加えておきますが、確率論的に言えば、それは曖昧なのです。
LLM を一桁信頼できるようにトレーニングするために必要なエネルギーと計算量 (たとえば、30% ではなく 3% の確率で間違ってしまうなど) は、現在のフロンティア モデルが必要とする量の 10^20 倍です。信頼性が大幅に向上したモデルがすぐに登場するとは期待しないでください。将来的に信頼性を向上させるには、コンテキスト エンジニアリング (入力に何を含めるかを決定する) を改善し、出力をどうするかを決定するより効果的な品質ゲートによって実現する必要があります。そして、それこそが、AI 企業が最近注目していることなのです。モデルはより強力になる可能性がありますが、大幅に信頼性が向上するわけではありません。皆さん、手持ちのものを使って作業してください。
AI 擁護者の中には、LLM がますます向上していることを実際に示している多くの公開されたベンチマークを指摘して、モデルのパフォーマンスが大幅に向上していないことを示す研究に抗議する人もいます。しかし、他の調査では、ベンチマークのパフォーマンスについてはもっと懐疑的なほうがよいかもしれないことがわかっています。その理由の一部は、最も人気のあるベンチマークの多くは、アルゴリズム的に測定しやすいものを測定しているためです。その意味で、それらは厄介で予測不可能な現実世界の問題とはまったく異なります。また、…そうですね、「公開されているベンチマーク」という言葉がちょっとしたヒントになるからです。

。おそらく意図せずに、あるいは意図的に、ますます多くのモデルが「テスト用にトレーニング」されるようになってきています。それが外にあるなら、それはおそらくそこにあります。
コード進化と長期的なエージェントワークフロー
SWE-CI: 継続的インテグレーションによるコードベースの保守におけるエージェントの能力の評価
https://arxiv.org/abs/2603.03823
SlopCodeBench: 長期にわたる反復タスクでコーディング エージェントがどのように劣化するかをベンチマークする
https://arxiv.org/html/2603.24755v1
SWE マイルストーン: 継続的なソフトウェア進化における AI エージェントの評価
https://arxiv.org/abs/2603.13428
ベンチマークと実際のパフォーマンスの比較
何が重要かを測定する: 大規模言語モデルのベンチマークで妥当性を構築する
https://arxiv.org/abs/2511.04703
研究最新情報: アルゴリズム評価と総合評価 (METR – 査読なし)
https://metr.org/blog/2025-08-12-research-update-towards-reconciling-slowdown-with-time-horizons/#background
LLM における評価データの汚染: どのように測定するのか、また (いつ) 問題になるのでしょうか?
https://arxiv.org/abs/2411.03923
必要なのはコンテキストです: LLM の現実世界の制限に対する最大有効コンテキスト ウィンドウ
https://arxiv.org/abs/2509.21361
RAG とロングコンテキストを超えて: 効率的な知識の基礎付けのための学習の注意散漫を意識した検索
https://arxiv.org/abs/2509.21865
AGENTS.md の評価: リポジトリ レベルのコンテキスト ファイルはエージェントのコーディングに役立ちますか?
https://arxiv.org/abs/2602.11988
(ブログ投稿 – 情報提供を目的としたものであり、査読済みの研究ではありません)
LLM トークン制限の背後に隠された科学 (そして、100 万トークン モデルが実際にどのように機能するか)
https://www.ashisharora.ai/post/the-hidden-science-behind-llm-token-limits-and-how-million-token-models-actually-work
言語モデルは否定派ではない: 否定ベンチマークにおける言語モデルの分析
https://arxiv.org/abs/2306.08189
の役割を再考する

デモンストレーション: コンテキスト学習が機能する理由は何ですか?
https://aclanthology.org/2022.emnlp-main.759
古い: LLM エージェントは自分の記憶がいつ無効になるかを知ることができますか?
https://arxiv.org/abs/2605.06527
課題: 知識要件による LLM の応答の形成
コンテキストとメモリの競合
https://aclanthology.org/2026.findings-acl.202
(モデル内の情報がコンテキスト内で提供される情報をオーバーライドする場合、私はこれらを「支配的な事前確率」と呼んでいます)
ソフトウェアエンジニアリングにおける大規模産業研究
AI コーディングの最大のリスクについて 2,800 万のワークフローが明らかにしたこと (CircleCI)
https://www.linkedin.com/pulse/what-28-million-workflows-reveal-ai-codings-biggest-risk-circleci-j9syc/
加速むち打ち – AI エンジニアリング レポート 2026 (Faros)
https://www.faros.ai/research/ai-acceleration-whiplash
2025 年の AI 支援ソフトウェア開発の現状 (DORA)
https://dora.dev/research/2025/dora-report/
心理学、認知、学習
超知性か迷信か？個人の行動に関する AI 予測への信念に影響を与える心理的要因を調査する
https://arxiv.org/html/2408.06602v3
クリティカルシンキングに対する生成型 AI の影響: ナレッジ ワーカーの調査から得た、認知的努力と自信への効果の自己申告の減少
https://www.researchgate.net/publication/391270185_The_Impact_of_Generative_AI_on_Critical_Thinking_Self-Reported_Reductions_in_Cognitive_Effort_and_Confidence_Effects_From_a_Survey_of_Knowledge_Workers
大規模言語モデルと Web 検索が学習の深さに及ぼす影響の実験的証拠
https://www.researchgate.net/publication/397000021_Experimental_evidence_of_the_Effects_of_large_ language_models_versus_web_search_on_ Depth_of_learning
費用はいくらですか? GenAI 時代におけるソフトウェア開発者の幸福
https://ourarchive.otago.ac.nz/es

ploro/outputs/preprint/At-What-Cost-Software-Developers-Well-Being/9926870122801891
LLM と深層学習の限界
大規模言語モデルが直面する壁 – (統計力学の研究)
https://arxiv.org/abs/2507.19703
勾配降下法で長期的な依存関係を学習するのは難しい
https://pubmed.ncbi.nlm.nih.gov/18267787/
(ディープ ニューラル ネットワークでは、どのようなスケールのモデルでも全体像を見ることができないのはなぜでしょうか)
このテクノロジーを使用するチームにはどのような影響がありますか?アジャイル ソフトウェア開発の技術的実践が AI 支援およびエージェント ソフトウェア エンジニアリングと非常に密接に連携している理由について、誇大広告のない、証拠に基づいた見解に興味がある場合は、10 月 6 日の 18:45 BST に参加してください。
登録: https://www.tickettailor.com/events/codemanship/2324138
X で共有 (新しいウィンドウで開きます)
×
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
Reddit で共有 (新しいウィンドウで開きます)
レディット
リンクを友人に電子メールで送信します (新しいウィンドウで開きます)
電子メール
Codemanship Ltd の創設者、コードクラフトのコーチ兼トレーナー
コードマンシップの投稿をすべて表示
AI の有無にかかわらず、迅速で信頼性が高く持続可能なソフトウェア配信のための技術実践におけるトレーニングとコーチングの詳細については、私のサイトをご覧ください。
AI ソフトウェア開発 – データは何を示していますか?
あなたのクオリティゲートはボウルの中の茶色のM&Msを見ていますか？
フィードバックの待ち時間、学習、「生産性」
たとえ自律型コーディング エージェントの領域を拡張できたとしても、それは拡張すべきだという意味でしょうか?
購読する
購読しました
コードマンシップのブログ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

I'm currently pulling together a bunch of sources - that are mostly recent - on the topic of LLMs and their use in software development. Some are peer-reviewed studies. Some are industry studies that haven't been peer-reviewed. One is statistical physics. Expect more from that angle. Wanna' know the
[truncated]

AI Software Development – What Does The Data Say? – Codemanship's Blog
Skip to content
Codemanship's Blog
Musings And Mutterings By Jason Gorman
AI Software Development – What Does The Data Say?
I’m currently pulling together a bunch of sources – that are mostly recent – on the topic of LLMs and their use in software development.
Some are peer-reviewed studies. Some are industry studies that haven’t been peer-reviewed.
One is statistical physics. Expect more from that angle. Wanna’ know the limits of a technology? Ask a physicist.
One is just a blog post, but very useful information about the effect of context size.
Most are corroborated by personal experiments and also observations on teams. As time goes on and more data comes in, my picture comes more into focus.
Before I cite the sources, a quick executive summary for all you busy executives out there:
Truly autonomous and reliable long-horizon agentic software development is so highly improbable using LLMs that it’s essentially science fiction.
The maximum effective context limits of LLMs – including hyperscale “frontier” LLMs – beyond which model outputs become unusably inaccurate is orders of magnitude smaller than advertised limits. The most common mechanism for extending inference over large contexts is what vendors call “compression”. This means that parts of the context are summarised by the model, which is a famously unreliable/lossy process.
LLMs cannot distinguish between recent and out-of-date information in the context, and information in the model itself, learned during training (“dominant priors”), can often “outweigh” information we give it. To an LLM, it’s all just tokens, weights and probabilities. Right, wrong, new, old – the highest probability wins. Big contexts and “attention dilution” – where probabilities in the context become too small to compete with the ones in the model – are likely to make these effects worse.
Repo-level .md files tend to make model performance worse, probably because they add noise instead of signal in many specific tasks. Model-generated .md files are especially problematic in this respect, it seems. Upshot: including your team’s coding standards and an architecture summary for every task is probably counterproductive.
LLMs struggle with negation. Telling them not to do something can often have the same effect as telling them to do it. In case you were wondering why some of your guardrails are about as reliable as a coin-toss.
LLM inference is more accurate when we give them examples (demonstrations) rather than just describing what we want. They’re pattern-matchers. Show them the patterns – more “like this” and less “do this” (and no “don’t do this”).
Large/long-scale industry studies show a clear trend – output is up (more code, more commits, bigger diffs), but outcomes don’t reflect that trend. If anything, the average team is taking longer to ship worse software. If ever we needed proof that software development isn’t a production process… Some studies find a small % of teams getting modest gains in outcomes, and correlate that with their existing software development capability. AI coding is an amplifier of, not a fix for, development strengths and weaknesses. (You’d think organisations would be lining up to do something about that… Sigh.)
The psychological and cognitive factors in LLM use are a growing field of serious research. One study found a significant correlation between confidence in AI output and belief in the paranormal. Multiple studies found a negative impact on learning, cognition and critical thinking with greater LLM reliance. New research suggests that reports of developers feeling demotivated and burned-out with extensive use may have some real truth behind them.
Deep neural networks, including LLMs, struggle to learn patterns with long-range dependencies, at any scale of model. They will always be “driving in fog”, with local, short-range probabilities crowding out long-range ones. In case you were wondering why they suck at the “big picture” – probabilistically, it’s a blur.
The energy and compute needed to train an LLM to be an order of magnitude more reliable – e.g., wrong 3% of the time instead of 30% – is 10^20 times what the current frontier models require. Don’t expect significantly more reliable models any time soon. Any future gains in reliability will have to made by better context engineering (deciding what to include in the input) and more effective quality gates deciding what to do with the output- and that’s exactly what we’re seeing AI companies focusing on these days. Models may get more powerful , but not significantly more reliable. This it folks – work with what you’ve got!
Some AI champions will protest research that points to no significant improvements in model performance by pointing to the many published benchmarks that do indeed show LLMs getting better and better. But other research finds that we might wish to be more skeptical of benchmark performance, partly because many of the most popular ones measure what’s easy to measure algorithmically – and in that sense, they’re not really like real-world problems which are messy and unpredictable – and also because… well, the words “published benchmarks” are a bit of a clue. Perhaps inadvertently, but maybe even knowingly, increasingly models are being “trained to the test”. If it’s out there, then it’s probably in there.
Code Evolution & Long-Horizon Agentic Workflows
SWE-CI: Evaluating Agent Capabilities in Maintaining Codebases via Continuous Integration
https://arxiv.org/abs/2603.03823
SlopCodeBench: Benchmarking How Coding Agents Degrade Over Long-Horizon Iterative Tasks
https://arxiv.org/html/2603.24755v1
SWE-Milestone: Evaluating AI Agents on Continuous Software Evolution
https://arxiv.org/abs/2603.13428
Benchmark vs. Real-World Performance
Measuring what Matters: Construct Validity in Large Language Model Benchmarks
https://arxiv.org/abs/2511.04703
Research Update: Algorithmic vs. Holistic Evaluation (METR – not peer-reviewed)
https://metr.org/blog/2025-08-12-research-update-towards-reconciling-slowdown-with-time-horizons/#background
Evaluation data contamination in LLMs: how do we measure it and (when) does it matter?
https://arxiv.org/abs/2411.03923
Context Is What You Need: The Maximum Effective Context Window for Real World Limits of LLMs
https://arxiv.org/abs/2509.21361
Beyond RAG vs. Long-Context: Learning Distraction-Aware Retrieval for Efficient Knowledge Grounding
https://arxiv.org/abs/2509.21865
Evaluating AGENTS.md: Are Repository-Level Context Files Helpful for Coding Agents?
https://arxiv.org/abs/2602.11988
(blog post – informational, not peer-reviewed research)
The Hidden Science Behind LLM Token Limits (And How Million-Token Models Actually Work)
https://www.ashisharora.ai/post/the-hidden-science-behind-llm-token-limits-and-how-million-token-models-actually-work
Language models are not naysayers: An analysis of language models on negation benchmarks
https://arxiv.org/abs/2306.08189
Rethinking the Role of Demonstrations: What Makes In-Context Learning Work?
https://aclanthology.org/2022.emnlp-main.759
STALE: Can LLM Agents Know When Their Memories Are No Longer Valid?
https://arxiv.org/abs/2605.06527
Task Matters: Knowledge Requirements Shape LLM Responses to
Context–Memory Conflict
https://aclanthology.org/2026.findings-acl.202
(I call these “dominant priors”, when information in the model overrides that provided in the context)
Large-scale Industry Studies in Software Engineering
What 28 million workflows reveal about AI coding’s biggest risk (CircleCI)
https://www.linkedin.com/pulse/what-28-million-workflows-reveal-ai-codings-biggest-risk-circleci-j9syc/
The Acceleration Whiplash – AI Engineering Report 2026 (Faros)
https://www.faros.ai/research/ai-acceleration-whiplash
State of AI-assisted Software Development 2025 (DORA)
https://dora.dev/research/2025/dora-report/
Psychology, Cognition & Learning
Super-intelligence or Superstition? Exploring Psychological Factors Influencing Belief in AI Predictions about Personal Behavior
https://arxiv.org/html/2408.06602v3
The Impact of Generative AI on Critical Thinking: Self-Reported Reductions in Cognitive Effort and Confidence Effects From a Survey of Knowledge Workers
https://www.researchgate.net/publication/391270185_The_Impact_of_Generative_AI_on_Critical_Thinking_Self-Reported_Reductions_in_Cognitive_Effort_and_Confidence_Effects_From_a_Survey_of_Knowledge_Workers
Experimental evidence of the effects of large language models versus web search on depth of learning
https://www.researchgate.net/publication/397000021_Experimental_evidence_of_the_effects_of_large_language_models_versus_web_search_on_depth_of_learning
At What Cost? Software Developers’ Well-Being in the Age of GenAI
https://ourarchive.otago.ac.nz/esploro/outputs/preprint/At-What-Cost-Software-Developers-Well-Being/9926870122801891
Limits of LLMs & Deep Learning
The wall confronting large language models – (statistical mechanics study)
https://arxiv.org/abs/2507.19703
Learning long-term dependencies with gradient descent is difficult
https://pubmed.ncbi.nlm.nih.gov/18267787/
(Why deep neural networks can’t see the bigger picture, at any scale of model)
What are the implications for teams using this technology? If you’re interested in a hype-free, evidence-based take about why the technical practices of Agile Software Development are so closely aligned with AI-assisted and agentic software engineering, join me on October 6th at 18:45 BST .
Register: https://www.tickettailor.com/events/codemanship/2324138
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Share on LinkedIn (Opens in new window)
LinkedIn
Share on Reddit (Opens in new window)
Reddit
Email a link to a friend (Opens in new window)
Email
Founder of Codemanship Ltd and code craft coach and trainer
View all posts by codemanship
Visit my site for details of Training & Coaching in Technical Practices for rapid, reliable and sustainable software delivery, with and without AI
AI Software Development – What Does The Data Say?
Do Your Quality Gates See The Brown M&Ms In The Bowl?
Feedback Latency, Learning & “Productivity”
Even If We Could Extend The Horizons of Autonomous Coding Agents, Does That Mean We Should?
Subscribe
Subscribed
Codemanship's Blog
Already have a WordPress.com account? Log in now.

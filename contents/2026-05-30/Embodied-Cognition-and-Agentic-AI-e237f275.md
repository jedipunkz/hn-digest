---
source: "https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/"
hn_url: "https://news.ycombinator.com/item?id=48325218"
title: "Embodied Cognition and Agentic AI"
article_title: "Embodied cognition and agentic AI – Daniel Lemire's blog"
author: "ibobev"
captured_at: "2026-05-30T11:43:51Z"
capture_tool: "hn-digest"
hn_id: 48325218
score: 4
comments: 0
posted_at: "2026-05-29T16:14:53Z"
tags:
  - hacker-news
  - translated
---

# Embodied Cognition and Agentic AI

- HN: [48325218](https://news.ycombinator.com/item?id=48325218)
- Source: [lemire.me](https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/)
- Score: 4
- Comments: 0
- Posted: 2026-05-29T16:14:53Z

## Translation

タイトル: 身体化された認知とエージェント的 AI
記事のタイトル: 身体化された認知とエージェント型 AI – Daniel Lemire のブログ
説明: あなたの知性はどこにありますか?あなたの脳内で？単純化した答えです。より良いモデルは、あなたの知性が具体化されているということです。高級レストランで働く料理人のことを考えてみましょう。彼はお気に入りのナイフと調理手順書をすべて、必要な場所に正確に配置しています。彼のキッチンはその一部です
[切り捨てられた]

記事本文:
身体化された認知とエージェント型 AI – Daniel Lemire のブログ
コンテンツにスキップ
ダニエル・レミアのブログ
Daniel Lemire はソフトウェア パフォーマンスの専門家です。彼は世界の科学者の上位 2% にランクされ (スタンフォード/エルゼビア 2025)、GitHub で最もフォローされている開発者のトップ 1000 の 1 人です。
12,500 人を超える電子メール購読者に加わりましょう:
身体化された認知とエージェント型 AI
AVX-512 で IPv6 アドレスを驚異的に高速に解析
すべての 64 ビット整数のうち 2 つの 32 ビット整数の積は 17% のみです
SIMD アクセラレーションによる整数から文字列への変換
乗算オーバーフローのチェック
ダニエル・レミア: 二分探索には勝てます
ザカリー・ドレマン「二分探索には勝てます」
Leonhard Staut on すべての 64 ビット整数のうち 2 つの 32 ビット整数の積は 17% だけです
Kasper が AVX-512 で IPv6 アドレスを驚異的に高速に解析することについて語る
KWillets on すべての 64 ビット整数のうち、2 つの 32 ビット整数の積は 17% のみです
身体化された認知とエージェント型 AI
あなたの知性はどこにありますか？あなたの脳内で？
単純化した答えです。より良いモデルは、あなたの知性が具体化されているということです。
高級レストランで働く料理人のことを考えてみましょう。彼はお気に入りのナイフと調理手順書をすべて、必要な場所に正確に配置しています。彼のキッチンは彼の知性、スキルの一部です。あなたのキッチンで働く同じ料理人は、おそらくあなたよりも上手に料理することができますが、お気に入りのキッチンで作るのと同じ料理を再現することはできません。
私たちはコンピューター プログラマーをホワイトボード テストを使用して評価することがよくあります。それは苦情の絶え間ない源です。プログラマは、それがプログラマの要素から強制的に排除されることを正しく指摘しています。ラップトップを持ち去ると、彼らはそれほど良くありません。それは言い訳ではなく、本当の問題です。あなたは彼らを、彼らを非常に賢くする要素の一部から切り離しているのです。
要約すると、脳としての知性のモデルは、

他のものから切り離された瓶はばかげています。
身体化された知性の考えを受け入れるなら、知性の結果として私たちが見ている多くの行動は、実際には知性の一部であることになります。何よりもまず言語です。お互いに話したり書いたりできるということは、私が自分という人間に制限されていないことを意味します。小さな部族に孤立した人類がテクノロジーの進歩を遂げているという話を聞いたことがありますか?いや。進歩するには、多くの人が一緒にコミュニケーションをとる必要があります。数十年前までは、進歩には都市が必要でした。今日では、どこからでも世界中の誰とでもコミュニケーションが取れるようになったため、以前よりも確信が持てなくなりました。しかし、言語は依然として重要であり、私たちはこれより優れたものを発明していません。同様に、手を持ち、洗練されたツール (ラップトップなど) を構築する能力によって、私たちの知性を拡張することができます。
2022 年末、私たちは ChatGPT という画期的なテクノロジーを手に入れました。これは、(大規模な) 言語モデル、ニューラル ネットワークなど、いくつかの既存のアイデアに基づいて構築されています。それが「GPT」の部分です。しかし、過小評価されているかもしれませんが、画期的な進歩の重要な部分は「チャット」コンポーネントでした。誰かが、大規模な言語モデルをチャット インターフェイスに接続するというアイデアを思いつきました。このシステムを構築している人々にとって、これは自然かつ明白なことかもしれませんが、それが取るに​​足らない、または重要ではないと考えるべきではありません。
言語は私たちの知能の重要な要素であり、したがって、言語が機械知能にとって極めて重要であることは当然です。
AI ソフトウェアをチャット ボックスに具体化しました。
次のステップは、今日「エージェント AI」と呼ばれるものです。チャット ボックスはそのままにしますが、AI ソフトウェアがツールと対話し、ツールの使用計画を立てる機能を追加します。事実上、私たちは AI にさらなる主体性を与えます。つまり、AI は何かを実行し、発生した結果から学習することができます。スタです

手と道具を持った人間に似ています。
今週同僚と話していました。私の同僚は AI 革命に全力で取り組んでいます。彼は AI を使用して、技術専門家の助けをあまり借りずに、文章をより正確かつ迅速に作成し、データ分析をより迅速に完了できるようにしています。
しかし、私の同僚はエージェント AI のアプローチを知りませんでした。電話で説明してみました。 AI にツールへのアクセスを許可するとはどういう意味ですか?これはAIの応答をコピー＆ペーストする手間を省くだけなのでしょうか？
最終的には、RStudio と呼ばれるもの内のシェルで AI を起動するビデオを作成することになりました。これは、人々が R でプログラミングしたり、データ分析を行ったりするために使用する環境です。私は R や RStudio を使用しませんが、AI のおかげで、Web からデータを取得するだけで、気候研究プロジェクト全体を数分で構築することができました。
AIはどうやってそれを行ったのでしょうか？録音しました。いくつかのことを試みましたが、最初はデータをダウンロードするのに苦労しました。ある時点で、新しい R パッケージが必要であることがわかり、それらをインストールします。インストールされると、図の生成を開始して、それが機能することを確認します。
エージェントティック AI は、AI の具体化を改善することでマシン インテリジェンスを大幅に拡張します。
まだ、あるべき姿が理解されていないのではないかと思います。
モントリオールでは、AI の分野で最も定評のある教授がヨシュア ベンジオです。彼は数年前に彼自身の重要な事業 (Element AI) を始めました。彼の最新の事業は、科学者 AI の作成を目的とした Law Zero です。このプロジェクトの最初の目標は、エージェント コンポーネントを使用せずに AI を構築することです。それは、それ自身の目標も主体性も持たない、実体を持たない AI でなければなりません。
私は、ベンジオがケビン・ケリーの言う思想主義に苦しんでいるのではないかと心配しています。ケリーの 2008 年のエッセイから引用しましょう。
どれほど優れた知性であっても、人間の体の仕組みを理解することはできません。

世界中の既知の科学文献をすべて読み、それを熟考するだけで効果があります。現在および過去のすべての核分裂実験を単純に考えて、実用的な核融合を 1 日で思いつくスーパー AI は存在しません。物事がどのように機能するかを知らないことと、物事がどのように機能するかを知っていることの間には、思考主義以上のものがあります。現実世界では、正しい作業仮説を形成するために必要な大量のデータが得られる実験が数多く行われています。潜在的なデータについて考えても、正しいデータは得られません。思考は科学の一部にすぎません。もしかしたらほんの一部かもしれない。 （…）思想主義だけでは十分ではありません。実験を行ったり、プロトタイプを構築したり、失敗したり、現実に関わったりしなければ、知性は思考を持つことはできても結果は得られません。世界の問題を解決する方法を考えることができません。 (…) シンギュラリティは絶えず後退し続ける幻想であり、常に「近く」にいますが、決して到達することはありません。 AIを手に入れた後、なぜそれが実現しなかったのか疑問に思うでしょう。そして将来のある日、私たちはそれがすでに起こったことに気づくでしょう。スーパー AI が登場しましたが、個人用ナノテクノロジー、脳のアップグレード、不死など、私たちが瞬時にもたらすと考えていたものはすべて実現しませんでした。その代わりに、私たちが予想していなかった他の利点が生じ、それを評価するまでに長い時間がかかりました。それらが来るのを私たちは見ていなかったので、私たちは振り返って、そう、あれがシンギュラリティだったと言います。
大学教授は特に思想主義になりやすいと思います。彼らは、知性は脳内で起こっていることに集中していると考えています。象牙の塔に住んでいると、知性の中核となる情報源である現実世界を無視してしまいがちです。さらに、彼らは、思想主義が自然に蔓延している学校で非常に良い成績を収めた人々であることがよくあります。
私は人生のほとんどを教授として過ごしてきました。しかし、

他の教授と話すのはすぐに飽きてしまいます。私が最も楽しんでいるのは、現実世界で応用できる新しいツールを持っている人々と仕事をすることです。当然のことながら、私はほとんどの時間を、人々が現実世界で展開するソフトウェアの作業に費やしました。
ケリーが言っていることは、高度な知性だけでは多くのことを行うのに十分ではないということです。現実世界はあなたの思考プロセスの最終段階ではありません。おそらくそれが最も重要な部分です。
したがって、AI を現実世界と接続し、(今日のほぼすべてのソフトウェア開発者が行っているように) 実験を実行できるようにすると、AI ソフトウェアが単独で実行できることをはるかに超える素晴らしい結果が得られます。
代理店は機能ではありません。代理店が第一です。
Daniel Lemire、「身体的認知とエージェント的 AI」、Daniel Lemire のブログ、2026 年 5 月 28 日、https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/ 。
https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/ '> [BibTeX]
ケベック大学 (TELUQ) のコンピューターサイエンス教授。
ダニエル・レミアの投稿をすべて表示
あなたのメールアドレスは公開されません。
次回コメントするときのために、このブラウザに名前、メールアドレス、ウェブサイトを保存してください。
このブログを電子メールで購読することもできます (非営利、広告なし、毎週の電子メール)
コードを投稿する場合は、 tohtml などのツールを使用してコードをフォーマットすることを検討してください。

## Original Extract

Where is your intelligence located? In your brain? It is a simplistic answer. A better model is that your intelligence is embodied. Consider a cook working at an expensive restaurant. He has all his favorite knives and cooking instructions, placed exactly where he wants them. His kitchen is part of
[truncated]

Embodied cognition and agentic AI – Daniel Lemire's blog
Skip to content
Daniel Lemire's blog
Daniel Lemire is a software performance expert. He ranks among the top 2% of scientists globally (Stanford/Elsevier 2025) and is one of GitHub's top 1000 most followed developers.
Join over 12,500 email subscribers:
Embodied cognition and agentic AI
Parsing IPv6 Addresses Crazily Fast with AVX-512
Only 17% of all 64-bit Integers are products of two 32-bit integers
SIMD-accelerated integer-to-string conversion
Checking multiplication overflow
Daniel Lemire on You can beat the binary search
Zachary Dremann on You can beat the binary search
Leonhard Staut on Only 17% of all 64-bit Integers are products of two 32-bit integers
Kasper on Parsing IPv6 Addresses Crazily Fast with AVX-512
KWillets on Only 17% of all 64-bit Integers are products of two 32-bit integers
Embodied cognition and agentic AI
Where is your intelligence located? In your brain?
It is a simplistic answer. A better model is that your intelligence is embodied.
Consider a cook working at an expensive restaurant. He has all his favorite knives and cooking instructions, placed exactly where he wants them. His kitchen is part of his intelligence, of his skills. The same cook working in your kitchen can probably cook better than you do, but he can’t reproduce the same meals he would prepare in his favorite kitchen.
We often assess computer programmers using whiteboard tests. It is an endless source of complaints. Programmers rightly point out that it forces them out of their element. They are just not as good when you take away their laptop. It is not an excuse, it is a real issue: you are cutting them off from part of what makes them so intelligent.
To sum it up, the model of intelligence as a brain in a jar, disconnected from anything else, is ridiculous.
If you accept the idea of embodied intelligence, then many actions that we view as a consequence of our intelligence are actually part of our intelligence. First and foremost, language. Our ability to talk or write to each other means that I am not limited by my own person. Have you ever heard of human beings isolated in small tribes making technological breakthroughs? Nah. Progress requires lots of people communicating together. Up until a few decades ago, progress required cities. Today I am less certain than it does, as I can more and more communicate with anyone in the world from anywhere. But language is still critical, we have not invented anything better. Similarly, having hands and the ability to build sophisticated tools (like laptops) allows us to extend our intelligence.
At the end of 2022, we got a breakthrough technology: ChatGPT. It built on several pre-existing ideas such as (large) language models, neural networks, and so forth. That’s the ‘GPT’ part. But an important, if underappreciated, part of the breakthrough was the ‘Chat’ component. Someone had the idea of connecting a large language model with a chat interface. Maybe this came naturally and obviously to people building this system, but it should not be assumed to be trivial or unimportant.
Language is a key component of our intelligence, and, thus, it makes sense that it would be pivotal for machine intelligence.
We embodied the AI software in a chat box.
The next step was what we call today ‘agentic AI’. We keep the chat box, but we add the ability for the AI software to interact with tools, and to make plans to use them. In effect, we give the AI more agency: it can do stuff and learn from the results as they happen. It is starting to resemble a human being with hands and tools.
I was talking with a colleague this week. My colleague is all in on the AI revolution. He uses his AI to help him write better and faster, and to get his data analysis done faster, without so much help from technical experts.
But my colleague was not aware of the agentic AI approach. I tried to explain on the phone. What does it mean to give the AI access to tools? Is this only about saving the effort of copying and pasting the AI’s response?
I ended up making a video where I start an AI in a shell within something called RStudio. It is an environment people use to program in R, to do data analysis. I don’t use R or RStudio, but thanks to the AI, I was able to build an entire climate research project in a few minutes, complete with the retrieval of the data from the web.
How did the AI do it? I recorded it. It tried a few things, initially struggling to download the data. At some point, it finds out that it needs new R packages, so it installs them, and once they are installed, it can proceed to generate figures, verifying that it works.
Agentic AI greatly extends machine intelligence by improving the embodiment of AI.
I believe that it is not yet understood as it should be.
In Montreal, the most established professor in the field of AI is Yoshua Bengio. He started his own non-trivial enterprise a few years ago (Element AI). His latest venture is Law Zero, which aims to create a Scientist AI. The first goal of this project is to build AI without the agentic component. It should be a disembodied AI that has no goal of its own, no agency.
I fear that Bengio suffers from what Kevin Kelly called Thinkism. Let me quote from Kelly’s 2008 essay.
No intelligence, no matter how super duper, can figure out how human body works simply by reading all the known scientific literature in the world and then contemplating it. No super AI can simply think about all the current and past nuclear fission experiments and then come up with working nuclear fusion in a day. Between not knowing how things work and knowing how they work is a lot more than thinkism. There are tons of experiments in the real world which yields tons and tons of data that will be required to form the correct working hypothesis. Thinking about the potential data will not yield the correct data. Thinking is only part of science; maybe even a small part. (…) Thinkism is not enough. Without conducting experiments, building prototypes, having failures, and engaging in reality, an intelligence can have thoughts but not results. It cannot think its way to solving the world’s problems. (…) The Singularity is an illusion that will be constantly retreating — always “near” but never arriving. We’ll wonder why it never came after we got AI. Then one day in the future, we’ll realize it already happened. The super AI came, and all the things we thought it would bring instantly — personal nanotechnology, brain upgrades, immortality — did not come. Instead other benefits accrued, which we did not anticipate, and took long to appreciate. Since we did not see them coming, we look back and say, yes, that was the Singularity.
I believe that University professors are especially prone to thinkism. They view intelligence as being centered on what is happening in their brain. When you live in an ivory tower, it is easy to dismiss the real world as the core source of intelligence. Further, they are often people who did quite well in school where thinkism is naturally prevalent.
I have been a professor most of my life. However, I tire quickly of talking with other professors. What I most enjoy is working with people who have new tools that they apply in the real world. Unsurprisingly, I spent most of my time working with software that people deploy in the real world.
What Kelly is saying is that a high degree of intelligence is not enough to do much of anything. The real world is not the final stage of your thinking process. It is maybe the most important part of it.
And thus, when you connect your AI with the real world, giving it the ability of running experiments (as virtually all software developers do today), you get impressive results that go much beyond what AI software can do on its own.
Agency is not a feature. Agency is primary.
Daniel Lemire, "Embodied cognition and agentic AI," in Daniel Lemire's blog , May 28, 2026, https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/ .
https://lemire.me/blog/2026/05/28/embodied-cognition-and-agentic-ai/ '> [BibTeX]
A computer science professor at the University of Quebec (TELUQ).
View all posts by Daniel Lemire
Your email address will not be published.
Save my name, email, and website in this browser for the next time I comment.
You can also subscribe by email to this blog (non commercial, no ads, weekly email)
If you want to post code, consider formatting it with a tool like tohtml .

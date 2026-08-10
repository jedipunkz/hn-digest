---
source: "https://www.technologyreview.com/2026/08/10/1141511/these-startups-are-chasing-the-next-big-thing-in-llms/"
hn_url: "https://news.ycombinator.com/item?id=49242087"
title: "Startups are chasing the next big thing in LLMs"
article_title: "These startups are chasing the next big thing in LLMs | MIT Technology Review"
author: "joozio"
captured_at: "2026-08-10T11:44:33Z"
capture_tool: "hn-digest"
hn_id: 49242087
score: 1
comments: 0
posted_at: "2026-08-10T11:02:19Z"
tags:
  - hacker-news
  - translated
---

# Startups are chasing the next big thing in LLMs

- HN: [49242087](https://news.ycombinator.com/item?id=49242087)
- Source: [www.technologyreview.com](https://www.technologyreview.com/2026/08/10/1141511/these-startups-are-chasing-the-next-big-thing-in-llms/)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T11:02:19Z

## Translation

タイトル: スタートアップ企業は LLM の次の大きなものを追いかけています
記事のタイトル: これらのスタートアップは LLM の次の大きなものを追いかけています | MITテクノロジーレビュー
説明: AI 巨人の後を追う新しい子供たちを紹介します。

記事本文:
コンテンツへスキップ MIT テクノロジー レビュー
MIT テクノロジー レビューの特集
人工知能 これらのスタートアップは LLM の次の大きなものを追いかけています
AI 巨人の後を追う新しい子供たちを紹介します。
ウィル・ダグラス・ヘブンのアーカイブページ
MIT Technology Review の What's Next シリーズでは、業界、トレンド、テクノロジーを横断的に取り上げ、将来の展望を提供します。残りの部分はここで読むことができます。
はるか昔の 2017 年の夏、Google の AI 研究者たちは、トランスフォーマーと呼ばれる新しいタイプのニューラル ネットワークについて説明した「Attending Is All You Need」という論文を発表しました。長いデータシーケンス、特にテキストの処理に非常に優れていることが判明しました。
9 年後、トランスフォーマーは、市場にあるすべての主要な大規模言語モデル内のエンジンとなっています。 「AI 業界全体はトランスフォーマーに基づいて構築されています」と AI スタートアップ Subquadratic の共同創設者兼 CEO である Justin Dangel は言います。 「これらはコンピューター サイエンスの歴史の中で最も重要なイノベーションの 1 つであり、世界を変えました。」
しかし、トランスフォーマーは老朽化を見せ始めています。いわゆる推論モデルの開発や一度に大量の入力を処理する機能など、LLM の最近の進歩の多くは、そのコア テクノロジのきちんとした拡張ではなく、その根本的な欠陥のいくつかを補う回避策です。
ますます多くの科学者やエンジニアが今、次に何が起こるのかを尋ねています。 LLM はどこにも行きませんが、その構築方法は手に入れることができます。 (MIT Technology Review は、AI で重要な 10 の今年のリストの中で、この将来世代のモデルを LLM+ と名付けました。)
このブームタウンのテクノロジーの限界を押し広げようとするスタートアップの波に乗り込みましょう。間違いなく失敗する人もいるだろうが、彼らは全力を尽くしてプレーするすべてを持っている

現在先頭に立っている企業よりも失うものが少ないのです。
しかし、まず問題です。トランスフォーマーの主な強みは、テキストのブロックの意味を一連の数字にエンコードする、デンス アテンションと呼ばれるメカニズムにあります。このプロセスでは、テキスト内のすべての単語 (またはトークンと呼ばれる単語の一部) を乗算の形式で他のすべての単語と比較します。
集中して注意を払うと、テキストの意味を驚くほど正確に捉えることができます。しかし、テキストの長さが長くなると、その処理に必要な計算の数が急速に増加します。 10,000 ワードの長さの文書には、5,000 万回の乗算を実行するための変換器が必要になる場合があります。これが、LLM が非常に多くの電力を消費する主な理由です。
費用は膨大です。同社社長のグレッグ・ブロックマン氏によると、OpenAIは今年コンピューティングに500億ドルを投じる予定だという。また、国際エネルギー機関は、データセンターで消費される電力の総量が 2030 年までに 2 倍になると予測しています。
さらに、変圧器は、最新モデルの多くがそのように設計されていることに苦労しています。トランスフォーマーはテキストを単語ごとに処理する方法のため、一度に多くの情報を追跡するのが得意ではありません (つまり、コンテキスト ウィンドウと呼ばれるものが大きくなりすぎることはありません)。それでも、LLM がより困難なタスクを実行する場合は、ドキュメントのライブラリ全体、コード ベース全体、エージェントの場合は他の LLM からの出力など、大量のデータを取り込む必要があります。
推論モデルに関しては、(思考の連鎖として知られる一種のスクラッチパッドに) 自分自身にメモを書き、それを読み返すことで機能します。これにより、管理しておくべきデータの量が再び増加します。
LLM が大きくなり性能が向上するにつれて、変圧器がボトルネックになってきました。このテクノロジーの主な強みは、今や限界となっています。
ここ

は、変圧器の問題を解決するための 4 つの新しいアイデアです。LLM を永久に変えて、LLM をより高速に、より効率的に、そして (おそらく) さらにスマートにする可能性のあるイノベーションです。
LLM をより速く、より安くするための明白な方法は、問題に正面から取り組み、注意の仕組みを変えることです。密なアテンションを、テキスト ブロック内のすべての単語の組み合わせではなく一部の単語の組み合わせに対してのみ計算を実行するスパース アテンションと呼ばれるメカニズムに置き換えることで、LLM が実行する必要のある計算量を大幅に削減できます。
研究者たちは長年にわたって、まばらな注意の仕組みをたくさん考案してきました。問題は、それらのどれもが意味を捉えることに集中力を発揮できなかったことです。
それは変わったかもしれない。マイアミに拠点を置くスタートアップ、Subquadratic は、検索やコーディングなどの少数のタスクに関して、主流のトップ LLM に匹敵する初のスパース アテンション メカニズムを発明したと主張している。これは大きな主張です (業界内の一部の人々は依然として懐疑的です)。
Subquadratic によると、そのモデルである SubQ は、与えられたテキストごとに、どの単語が重要でどの単語が重要でないかをその場で判断することで機能します。同社はまた、数千人が順番待ちリストに登録しており、このモデルを間もなく広く利用できるようにする予定だと主張している。
一方、サンフランシスコに拠点を置く新興企業 Manifest AI は、別の角度からこの問題に取り組んでいます。注意の仕組みを変えるのではなく、注意を別のものに置き換えているのです。
これは、特定のタスクに最も関連性の高い情報のみを保存し、LLM が追跡しなければならないデータ量が爆発的にならないようにする、パワーリテンションと呼ばれるメカニズムを開発しました。
アテンション メカニズムにより、LLM はコンテキスト ウィンドウ内のすべてを追跡するように強制されます。スパース アテンション モデル (SubQ など) では、多くの個々の単語が無視されます。

それは今でも、見たものすべての大まかなイメージを保持しています。対照的に、パワーリテンションは、コンテキスト ウィンドウのローリング サマリーをモデルに提供することによって機能します。新しい情報が追加されると、関連性の低い情報は削除されます。
保持の基本原則は 10 年前から存在しています。 Manifest AI は、これらの技術を更新して、トランスベースの LLM に初めて耐えられるモデルを構築したと主張しています。
同社によれば、最小限の再トレーニングで変圧器モデルを電力保持モデルに適応させることが可能だという。これを実証するために、StarCoder と呼ばれる既存のオープンソース コーディング LLM を、PowerCoder と呼ばれる電力保持を使用するバージョンに変更しました。また、同社は Brumby と呼ばれるモデルもリリースしており、これはアリババの人気オープンソース モデル Qwen の一部のバージョンに匹敵すると主張しています。
Manifest AI は、LLM が大量のデータの処理を伴うタスクを実行する必要がある場合に、その電力保持技術が頼りになるソリューションになることを望んでいます。 Manifest AI の共同創設者で CTO の Carles Gelada 氏は、昨年自社のテクノロジーを発表したビデオの中で、数時間にわたるビデオの分析から、一度に数週間タスクを続けることができるエージェントの構築まで、多くの有用なアプリケーションがあると主張しました。
02: モデルの小型化と柔軟性の向上
マサチューセッツ州ケンブリッジに本拠を置くMITのスピンアウトであるLiquid AIは、変圧器を完全に変更したり廃止したりはしていないが、変圧器を独自の技術であるリキッドニューラルネットワークと組み合わせて、共同創設者兼CEOのラミン・ハサニ氏がLFM（リキッドファンデーションモデル）と呼ぶものを構築している。
Liquid AI のモデルは、ほとんどの LLM よりもはるかに小さく、使用するエネルギーも少なくなります。同社は、メルセデスを含む自動車メーカー向けに、車両内の小型チップで動作するモデルを製造している。同社の最新モデルは、50 ドルの低消費電力趣味用コンピューターである Raspberry Pi で実行できます。
それはもう

dels は、年間収益が 1,000 万ドル未満の組織であれば無料で利用できます。そして、それらは人気があることが証明されており、同社は約 3,400 万ダウンロードを記録した、とハサニ氏は言います。
液体ニューラル ネットワークは、線虫の脳からインスピレーションを得たものです。これらは、畳み込みネットワークと呼ばれる、トランスフォーマーよりも前の別のタイプのニューラル ネットワークの拡張です。重要なイノベーションは、モデルがその動作を新しい情報に適応させ、その過程で学習できるようにするメカニズムです。トランスフォーマーではそれは不可能です。モデルがトレーニングされると、その動作は固定されます。
Liquid AI の最初のモデルは非常に基本的なものでしたが、ドローンを飛ばしたり、車両を運転したりすることができました。同社は、LFM を使用して、主流の LLM と競合するためにテクノロジーをスケールアップしようとしています。その新モデルは、アリババの Qwen や Google のオープンソース LLM Gemma のバージョンなど、4 倍規模のライバルのパフォーマンスに匹敵します。
一般的な LLM は、配線された変圧器のスタックから構築されます。 Liquid AI の最近の LFM は、20% のトランスフォーマーと 80% のリキッド ニューラル ネットワークで構成されるハイブリッド モデルです。
この比率は、Liquid AI が構築した別の AI システムによって発見され、すべてのモデルの設計に使用されています。 「これは現在の当社の中核技術です」とハサニ氏は言います。このデザイナー AI は、ニューラル ネットワーク (液体、畳み込みなど、トランスフォーマーなど) のさまざまな組み合わせを精査し、さまざまなネットワークをボルトで結合してパフォーマンスと効率のスイート スポットを達成する設計を考案します。
ハサニ氏は、変圧器はほんの始まりにすぎないと考えています。「ご存知のように、あなたの脳は AGI システムであり、20 ワットの電力で動作します。どうしてそんなことが可能なのでしょうか? 私たちはもっと革新的なものを手に入れることができます。」
03: テキストを一括生成する
ほとんどすべての LLM は、一度に 1 ワードずつ出力を生成します。それは理にかなっています、なぜなら

それが人が話したり書いたりする方法なのです。しかし、コンピュータの場合、それは非常に非効率的です。
LLM では、テキストを一度に生成し、文全体または段落全体を 1 回のショットで吐き出す方が速くて安価です。これは、カリフォルニア州パロアルトに拠点を置く新興企業 Inception が採用したアプローチであり、拡散と呼ばれる手法を使用して LLM を構築しています。
拡散は、ほとんどの画像およびビデオ生成モデルを推進するテクノロジーとしてよく知られています。拡散モデルは、古いテレビの静電気のようなピクセルのランダムなグリッドを取得し、それを画像に変換するようにトレーニングされています。これは、すべてのピクセルを同時に処理し、静止画をより高解像度の写真のように見せるためにどのピクセルを変更する必要があるかを判断することでこれを実現します。
このプロセスはテキストに対しても機能することがわかりました。 Inception は、ランダムな単語の文字列を取得し、それを意味のある文章に変換するように LLM をトレーニングしました。 Diffusion LLM は依然としてトランスフォーマーを使用して意味をエンコードしていますが、テキストのブロック全体を一度に生成することで、トランスフォーマーはより少ないコストでより多くのことを実行できるようになります。 「依然として大規模なトランスフォーマー モデルを使用していますが、同時に多くのトークンを予測できます」と Inception の共同創設者兼 CEO の Stefano Ermon 氏は言います。 「だからこそ、これらのモデルは、他のほとんどの人が今日構築しているものと比べて、はるかに高速でコスト効率が高いのです。」
課題は、画像生成用に設計されたテクノロジーをテキストに適用することでした。画像の場合、青いピクセルを赤いピクセルに変更する必要がある場合は、中間色を段階的に通過できるとエルモン氏は言います。これはテキストでは機能しません。「『猫』と『犬』がある場合、その間には実際には何もありません。」
エルモンはスタンフォード大学の研究者でもあります。 2024 年、彼とスタンフォード大学の同僚 2 人は、拡散モデルをテキストで機能させるための数学を考え出しました。彼らは訓練を受けました

この拡散モデルは、OpenAI が 2019 年に構築した LLM である GPT-2 のパフォーマンスに匹敵し、しかも 10 倍高速でした。エルモンにとっては会社をスピンアウトするだけで十分だった。
現在、彼は大リーグに照準を合わせている。 Inception は、最新モデル Mercury 2 は、2023 年にリリースされた OpenAI の GPT-4 モデルの一部と同等のパフォーマンスを示しますが、やはり 10 倍高速であると主張しています。 「私たちはこのアプローチに強気です。なぜなら、それが規模を拡大するものだからです」とエルモン氏は言います。
重要なのは速度とコストだけだ、と彼は付け加えた。「最終的には、通貨は 1 ドルあたりの情報になるでしょう。」
普及に賭けているのはインセプションだけではない。 Google もこのアプローチを実験しており、 Diffusion Gemma と呼ばれるプロトタイプ LLM を構築しました。しかしエルモン氏は競争を心配していない。 「それは有効だと思います」と彼は言います。 「これが未来だ。」
パロアルトに本拠を置く別のスタートアップである Pathway は、おそらくこの新しい集団の中で最も極端な企業だ。 LLM を言語の制約から解放したいと考えています。
同社は、Dragon Hatchling と呼ばれる LLM の一種を構築しました (名前の由来は、テリー・プラチェットの小説『カラー・オブ・マジック』に登場するドラゴンにちなんで名付けられました。よくよく考えれば現実化します)。これまでのところ、その際立った結果は、LLM を 250,000 を超える非常に難しい数独パズルと比較するベンチマークでの高いスコアです。ドラゴンの孵化したばかりの子がパズルの 97% 以上を破りました

[切り捨てられた]

## Original Extract

Meet the new kids nipping at the heels of the AI giants.

Skip to Content MIT Technology Review Featured
MIT Technology Review Featured
Artificial intelligence These startups are chasing the next big thing in LLMs
Meet the new kids nipping at the heels of the AI giants.
Will Douglas Heaven archive page
MIT Technology Review ’s What’s Next series looks across industries, trends, and technologies to give you a first look at the future. You can read the rest of them here .
Way back in the summer of 2017, AI researchers at Google put out a paper called “Attention Is All You Need,” in which they described a new type of neural network called a transformer. It proved to be very good at processing long sequences of data, especially text.
Nine years on, transformers are the engines inside every major large language model on the market. “The entire AI industry is built on transformers,” says Justin Dangel, cofounder and CEO of the AI startup Subquadratic. “They are one of the most important innovations in the history of computer science, and they’ve changed the world.”
But transformers are starting to show their age. Many of the recent advances in LLMs, such as the development of so-called reasoning models and their ability to handle large amounts of input at once, are not neat extensions of that core technology but workarounds that patch over some of its fundamental flaws.
A growing number of scientists and engineers are now asking what’s coming next. LLMs are not going anywhere, but the way they get built is up for grabs. ( MIT Technology Review dubbed this future generation of models LLMs+ in this year’s list of the 10 things that matter in AI .)
Enter a wave of startups hoping to push the boundaries of this boomtown technology. Some will no doubt fail—but they have everything to play for and far less to lose than the companies at the front of the pack today.
But first, the problem. The key strength of transformers lies in a mechanism called dense attention, which encodes the meaning of a block of text in a series of numbers . The process involves comparing every word (or part of a word, known as a token) in that text with every other word via a form of multiplication.
Dense attention can capture the meaning of text with remarkable accuracy. But as the length of that text grows, the number of computations needed to process it adds up fast. A document 10,000 words long might require a transformer to perform 50 million multiplications. That’s the main reason LLMs suck up so much power .
The costs are huge. OpenAI is set to spend $50 billion on computing this year , according to the company’s president, Greg Brockman. And the International Energy Agency predicts that the total amount of electricity consumed by data centers will double by 2030 .
What’s more, transformers struggle with what many of the latest models are designed to do. Because of the way they process text word by word, transformers are not great at keeping track of a lot of information at once (in other words, what's known as their context window cannot get too large). And yet if LLMs are to carry out harder tasks, they will need to take in larger amounts of data: a whole library of documents, an entire code base, or in the case of agents, output from other LLMs.
As for reasoning models, they work by writing notes to themselves (in a kind of scratch pad known as a chain of thought ) and then reading them back, which again adds to the amount of data to stay on top of.
As LLMs get bigger and better, transformers have become a bottleneck. The technology’s key strength is now a limitation.
Here are four new ideas for how to solve the transformer problem—innovations that could change LLMs for good, making them faster, far more efficient, and (maybe) even smarter.
An obvious way to make LLMs faster and cheaper is to tackle the problem head on and change the way attention works. Swapping out dense attention for a mechanism called sparse attention, which runs calculations on only some pairings of words in a block of text instead of all of them, can radically reduce the amount of computation LLMs need to do.
Researchers have come up with plenty of sparse attention mechanisms over the years. The problem is that none of them were as good as dense attention at capturing meaning.
That might have changed. Subquadratic, a startup based in Miami, claims it has invented the first sparse attention mechanism that rivals top mainstream LLMs on a handful of tasks, including search and coding. It’s a huge claim (and some people in the industry remain skeptical ).
Subquadratic says its model, SubQ, works by figuring out on the fly—for each piece of text it is given—which words matter and which don’t. The company also claims that thousands have signed up to its waitlist and plans to make the model widely available soon.
Meanwhile, Manifest AI, a startup based in San Francisco, is coming at the problem from a different angle. Instead of changing how attention works, it is replacing it with something else.
It has developed a mechanism it calls power retention, which stores only the most relevant information for a given task and ensures that the amount of data an LLM has to keep track of doesn’t blow up.
Attention mechanisms force LLMs to keep track of everything in their context window. A sparse attention model (such as SubQ) throws out a lot of the individual words, but it still retains a rough picture of everything it has seen. In contrast, power retention works by providing the model with a rolling summary of its context window. As new information is added, less relevant information is dropped.
The basic principle of retention has been around for a decade. Manifest AI claims it has updated those techniques to build models that can stand up to transformer-based LLMs for the first time.
The company says it is possible to adapt a transformer model into a power retention model with minimal retraining. To demonstrate this, it has turned an existing open-source coding LLM called StarCoder into a version that uses power retention, called PowerCoder. It has also released a model called Brumby, which it claims rivals some versions of Alibaba’s popular open-source model Qwen.
Manifest AI wants its power retention tech to become the go-to solution when LLMs need to carry out tasks that involve processing huge amounts of data. There are many useful applications, Manifest AI’s cofounder and CTO, Carles Gelada, claimed in a video announcing his company’s technology last year—from analyzing videos that are hours long to building agents that can stay on task for weeks at a time.
02: Making models smaller and more flexible
Liquid AI, an MIT spinout based in Cambridge, Massachusetts, hasn’t changed or ditched transformers fully but pairs them with its own tech, liquid neural networks, to build what cofounder and CEO Ramin Hasani calls LFMs (liquid foundation models).
Liquid AI’s models are far smaller and use less energy than most LLMs. The firm builds models for car makers, including Mercedes, which run on the small chips inside vehicles. Its latest models can run on a Raspberry Pi, a low-powered hobbyist computer that costs $50.
Its models are available for free to any organization with an annual revenue less than $10 million. And they have proved popular: The company has racked up almost 34 million downloads, says Hasani.
Liquid neural networks were inspired by worm brains. They are an extension of another type of neural network that predates transformers, called convolutional networks. The key innovation is a mechanism that lets a model adapt its behavior to new information, so it can learn as it goes. That’s not possible with transformers: Once a model is trained, its behavior is fixed.
Liquid AI’s first models were pretty basic but could fly drones or drive vehicles. With LFMs, the company is trying to scale up its technology to compete with mainstream LLMs. Its new models match the performance of rivals four times bigger, including versions of Alibaba’s Qwen and Google’s open-source LLM Gemma.
A typical LLM is built from a stack of transformers wired together. Liquid AI’s recent LFMs are hybrid models made up of 20% transformers and 80% liquid neural networks.
That ratio was hit upon by another AI system that Liquid AI has built, which it uses to help design all its models. “It’s the core technology of our company right now,” says Hasani. This designer AI sifts through many different combinations of neural networks—liquid, convolutional, and more, as well as transformers—and comes up with designs that bolt different ones together to hit a sweet spot of performance and efficiency.
Hasani thinks transformers were just the beginning: “Your brain is an AGI system, you know, and it operates with 20 watts of power. How is it possible? We can get a lot more innovative.”
03: Generating text all at once
Almost all LLMs produce their output one word at a time. It makes sense, because that is how people speak and write. But for computers, it’s very inefficient.
It is faster and cheaper for LLMs to generate text all at once—spitting out whole sentences or paragraphs in one shot. That’s the approach taken by Inception, a startup based in Palo Alto, California, which is building LLMs using a technique called diffusion.
Diffusion is better known as the technology that drives most image and video generation models . Diffusion models are trained to take a random grid of pixels—like the static on an old TV set—and turn it into an image . They do this by working on all the pixels at the same time, figuring out which need changing to make the static look more like a high-definition photo.
It turns out this process works on text too. Inception has trained its LLMs to take a random string of words and turn it into sentences that make sense. Diffusion LLMs still use transformers to encode meaning, but by producing whole blocks of text at once, they make transformers do more for less. “You’re still using a big transformer model, but you can predict many tokens at the same time,” says Inception’s cofounder and CEO, Stefano Ermon. “That’s why these models are so much faster and cost-efficient compared to what most other people are building today.”
The challenge was to take a technology designed for image generation and apply it to text. With images, if you need to change a blue pixel to a red one you can step through intermediate colors, says Ermon. That doesn’t work with text: “When you have ‘cat’ and ‘dog,’ there is not really something in between.”
Ermon is also a researcher at Stanford University. In 2024, he and a pair of his Stanford colleagues figured out the math to make diffusion models work with text. They trained a diffusion model that matched the performance of GPT-2—an LLM that OpenAI built in 2019—but was 10 times faster. It was enough for Ermon to spin out a company.
Today he has his sights on the big league. Inception claims its latest model, Mercury 2, performs as well as some of OpenAI’s GPT-4 models, released in 2023, but again 10 times faster. “We’re bullish about this approach because it’s the one that is going to scale up,” says Ermon.
The only things that matter are speed and cost, he adds: “Ultimately, the currency is going to be intelligence per dollar.”
Inception is not the only company betting on diffusion. Google is also experimenting with this approach and has built a prototype LLM called Diffusion Gemma . But Ermon is not worried about the competition. “I think it's validating,” he says. “This is the future.”
Pathway, another startup based in Palo Alto, is perhaps the most extreme of this new bunch. It wants to free LLMs from the constraints of language.
The firm has built a type of LLM called Dragon Hatchling (named after the dragons in Terry Pratchett’s novel Color of Magic , which materialize if you think about them hard enough). Its standout result so far is a high score on a benchmark that pits LLMs against more than 250,000 very hard sudoku puzzles. Dragon Hatchling beat more than 97% of the puz

[truncated]

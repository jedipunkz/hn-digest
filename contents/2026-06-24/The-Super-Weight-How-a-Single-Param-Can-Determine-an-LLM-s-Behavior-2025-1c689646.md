---
source: "https://machinelearning.apple.com/research/the-super-weight"
hn_url: "https://news.ycombinator.com/item?id=48661495"
title: "The \"Super Weight:\" How a Single Param Can Determine an LLM's Behavior (2025)"
article_title: "The “Super Weight:” How Even a Single Parameter can Determine a Large Language Model’s Behavior - Apple Machine Learning Research"
author: "sarreph"
captured_at: "2026-06-24T15:46:04Z"
capture_tool: "hn-digest"
hn_id: 48661495
score: 1
comments: 0
posted_at: "2026-06-24T15:34:04Z"
tags:
  - hacker-news
  - translated
---

# The "Super Weight:" How a Single Param Can Determine an LLM's Behavior (2025)

- HN: [48661495](https://news.ycombinator.com/item?id=48661495)
- Source: [machinelearning.apple.com](https://machinelearning.apple.com/research/the-super-weight)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T15:34:04Z

## Translation

タイトル: 「スーパー ウェイト:」単一パラメータが LLM の動作を決定する方法 (2025)
記事のタイトル: 「スーパー ウェイト:」単一パラメータでも大規模な言語モデルの動作をどのように決定できるか - Apple Machine Learning Research
説明: Apple 研究者による最近の論文「[The Super Weight in Large Language…

記事本文:
「スーパーウェイト」: 単一のパラメータでも大規模な言語モデルの動作をどのように決定できるか - Apple 機械学習の研究 機械学習の研究 メニューを開く メニューを閉じる 概要
「スーパーウェイト」: 単一パラメータでも大規模な言語モデルの動作をどのように決定できるか
図 1: スーパー ウェイト現象: 「スーパー ウェイト」と呼ばれる単一の特殊なスカラーを枝刈りすると、大規模言語モデルのテキスト生成機能が完全に破壊される可能性があります。左側は、スーパーウェイトを内蔵したオリジナルの Llama-7B で、それなりの完成度を示しています。右側では、スーパーウェイトを削除した後、Llama-7B が完全に意味不明な内容を生成します。この定性的な観察は定量的な影響もあります。ゼロショットの精度はランダムに低下し、複雑さは桁違いに増加します。
図 1: スーパー ウェイト現象: 「スーパー ウェイト」と呼ばれる単一の特殊なスカラーを枝刈りすると、大規模言語モデルのテキスト生成機能が完全に破壊される可能性があります。左側は、スーパーウェイトを内蔵したオリジナルの Llama-7B で、それなりの完成度を示しています。右側では、スーパーウェイトを削除した後、Llama-7B が完全に意味不明な内容を生成します。この定性的な観察は定量的な影響もあります。ゼロショットの精度はランダムに低下し、複雑さは桁違いに増加します。
Apple の研究者による最近の論文「The Super Weight in Large Language Models」では、LLM のパラメータの非常に小さなサブセット (場合によっては単一パラメータ) が、LLM の全体的な機能に不釣り合いな影響を与える可能性があることを明らかにしています (図 1 を参照)。この研究では、これらの「スーパー ウェイト」とそれに対応する「スーパー アクティベーション」の重要な役割に焦点を当て、LLM アーキテクチャと効率的なモデル コムへの道についての新たな洞察を提供します。

圧迫感。この論文には完全な技術的詳細と実験結果が記載されています。この投稿では、主要な調査結果とその影響についての概要を説明します。
ますます大規模になるモデルの理解と圧縮
LLM は優れた機能を発揮しますが、その規模が非常に大きいため、多くの場合、数十億、さらには数千億のパラメータが含まれるため、モバイル デバイスなどのリソースに制約のあるハードウェアへの導入には大きな課題が生じます。このようなプラットフォームの LLM のサイズと計算の複雑さを軽減すると、それに応じてメモリと消費電力が削減され、インターネット接続なしでローカルでプライベートに動作できるようになります。ただし、単純な圧縮や単純化はモデルの品質の大幅な低下につながる可能性があるため、LLM の内部メカニズムを理解することが重要です。
スーパーウェイトとその影響の特定
以前の研究では、LLM のパラメータの外れ値のごく一部がモデルの品質を維持するために不可欠であることが示されています。そして、これらの重みが大幅に変更される (圧縮によって) か完全に削除される (枝刈り) と、モデルの出力品質が低下します。この以前の研究では、この割合が重みの 0.01% 程度に小さい可能性があることが示されましたが、数十億のパラメータを持つモデルでは、これは依然として数十万の個々の重みに相当します。この研究で、Apple の研究者は、「スーパー ウェイト」と呼ばれる非常に少数のパラメータを特定しました。これらのパラメータが変更されると、LLM の一貫したテキストを生成する能力が破壊される可能性があり、たとえば、複雑さが 3 倍増加し、ゼロショットの精度がランダムな推測と一致するレベルに低下する可能性があります。たとえば、Llama-7B モデルでは、その 1 つのスーパー ウェイトを削除すると、モデルは prod できなくなります。

意味のある出力が得られます。逆に、他の何千もの外れ値の重みを削除しても、スーパー ウェイトよりも大きな値を持つものであっても、わずかな品質の低下しか生じません。
この研究では、モデルを通過する 1 回の前方パスのみを必要とすることで、これらのスーパー ウェイトを特定する方法論を提案しています。この方法は、スーパー ウェイトが、それに応じてまれで大きなアクティベーション外れ値を誘発するという観察を活用しており、これを「スーパー アクティベーション」と呼びます。これらのスーパー アクティベーションは、スーパー ウェイトの後に現れることが多く、入力プロンプトに関係なく、後続のレイヤー全体にわたって一定の大きさと位置で持続し、そのチャネルはスーパー ウェイトのチャネルと一致します。特定のモデル コンポーネントの入力および出力のアクティベーション分布 (フィードフォワード ネットワークの下方投影など) のスパイクを検出することにより、対応するスーパー アクティベーションを介してスーパー ウェイトを特定できます。興味深いことに、スーパー ウェイトは、アテンション ブロックに続くフィードフォワード ネットワークの下方投影で、通常はネットワークの初期層で一貫して見つかります。私たちは、研究コミュニティによるさらなる調査を容易にするために、いくつかの一般的で公開されている LLM のスーパー ウェイト座標のインデックスを作成しました。
表 1: 上記のレイヤー番号、レイヤー タイプ、およびウェイト タイプは、直接適用できます。
ハグフェイスモデル。たとえば、Huggingface の Llama-7B の場合は、layers[2].mlp.down_proj.weight[3968, 7003] を使用してスーパー ウェイトにアクセスします。
座標テーブル (表 1 を参照) に示されているように、スーパー ウェイトは特定の射影レイヤーで出現し、多くの場合、一般的に使用される広範囲の LLM にわたるネットワークの初期段階で発生します。これらの重みによってスーパー アクティベーションが生成され、図 2 に示すように、ネットワーク内の残りのスキップ接続を通じて持続します。

この永続的なスーパー アクティベーションはモデルの内部ダイナミクスに全体的な影響を与え、高確率のストップワードの生成を妨げます。スーパー ウェイトが削除されると、この抑制効果がなくなり、モデルの出力分布が急激に変化します。ストップワードの可能性が大幅に増加する一方、意味のある内容を含むトークンの可能性は低くなります。これは、モデルの前方パス中にどの意味論的に意味のあるトークンが出力されるかを決定する上で、スーパーウェイトが重要な役割を果たすことを示唆しています。
強化された圧縮とモデルの理解
スーパー ウェイトとスーパー アクティベーションの発見は、LLM 圧縮の改善と、これらのモデルに対する分野のより広範な理解につながる可能性があります。これらの少数のパラメータの大きな影響は、LLM 圧縮技術中にそれらのパラメータを保存することが重要であることを示唆しています。私たちは、スーパー アクティベーションを高精度で保持することにより、単純な最近値丸め量子化手法が、より洗練された最先端の手法と競合するパフォーマンスを達成できることを発見しました。同様に、重み量子化の場合、他の重みの外れ値をクリップしながらスーパー重みを保持することで、これまで実現可能と考えられていたよりもはるかに大きなブロック サイズでも、最も近い量子化への丸めが有効になり、圧縮率が向上します。
この研究は、わずか数個の超外れ値を処理するだけで圧縮品質が大幅に向上し、数十万の外れ値の重みを管理する方法と比較してハードウェアに優しいアプローチを提供できることを示しています。この的を絞ったアプローチにより、元のパフォーマンスをより高度に維持した、より効率的なモデルが得られます。これにより、強力な LLM アプリケーションが、モバイル デバイスなどのリソースに制約のあるハードウェア上で高品質で動作できるようになります。

超外れ値の状況を探る
私たちの発見は、将来の研究にいくつかの道を開きます。スーパーウェイトとスーパーアクティベーションの起源と正確なメカニズムをさらに調査すると、LLMの動作ダイナミクスについてのより深い洞察が得られる可能性があります。これらの特定のパラメーターがトレーニング中にどのようにしてそのような不均衡な影響を獲得するかを理解することは、将来のモデル設計とトレーニング戦略に役立つ可能性があります。より広範なモデル アーキテクチャとトレーニング パラダイムにわたってスーパー ウェイトの普及と特徴を調査すると、その役割や作成に光を当てることができ、提供されたスーパー ウェイトのディレクトリは、コミュニティ内でのそのような継続的な調査を促進することを目的としています。最終的には、これらの超外れ値をより包括的に理解することで、より効率的で堅牢で解釈可能な LLM を構築するための新しい方法論が解き放たれる可能性があります。
大規模言語モデルのスーパーウェイト
2025年7月2日 研究分野 音声と自然言語処理
最近の研究では、驚くべき結果が示されました。それは、大規模言語モデル (LLM) パラメーターの外れ値のごく一部が、モデルの品質にとって不釣り合いに重要であるということです。 LLM には数十億のパラメータが含まれているため、0.01% などの小さな割合は数十万のパラメータに変換されます。この研究では、さらに驚くべき発見を示します。パラメータを 1 つでもプルーニングすると、LLM のテキスト生成機能が破壊される可能性があります。
R^2: モデルの圧縮と量子化のための範囲正則化
2023年8月10日 研究分野 コンピュータビジョン 、研究分野 音声および自然言語処理
モデル パラメーターの正則化は、一般化を改善するために広く使用されている手法ですが、さまざまな目的で重み分布を形成するためにも使用できます。この作品では、

重み正則化がモデルの量子化と圧縮技術をどのように支援できるかを説明し、外れ値の防止に焦点を当ててモデル最適化の品質をさらに高めるための範囲正則化 ( R 2 R^2 R 2 ) を提案します。
効果的に…
機械学習のチャンスを見つけてください。
機械学習における私たちの研究は、日々新たな境地を開拓しています。
「スーパーウェイト」: 単一パラメータでも大規模な言語モデルの動作をどのように決定できるか

## Original Extract

A recent paper from Apple researchers, "[The Super Weight in Large Language…

The “Super Weight:” How Even a Single Parameter can Determine a Large Language Model’s Behavior - Apple Machine Learning Research Machine Learning Research Open Menu Close Menu Overview
The “Super Weight:” How Even a Single Parameter can Determine a Large Language Model’s Behavior
Figure 1: Super Weight Phenomenon: Pruning a single, special scalar, called the “super weight,” can completely destroy a Large Language Model’s ability to generate text. On the left, the original Llama-7B, which contains a super weight, produces a reasonable completion. On the right, after pruning the super weight, Llama-7B generates complete gibberish. This qualitative observation has quantitative impact as well: zero-shot accuracy drops to random and perplexity increases by orders of magnitude.
Figure 1: Super Weight Phenomenon: Pruning a single, special scalar, called the “super weight,” can completely destroy a Large Language Model’s ability to generate text. On the left, the original Llama-7B, which contains a super weight, produces a reasonable completion. On the right, after pruning the super weight, Llama-7B generates complete gibberish. This qualitative observation has quantitative impact as well: zero-shot accuracy drops to random and perplexity increases by orders of magnitude.
A recent paper from Apple researchers, “ The Super Weight in Large Language Models ,” reveals that an extremely small subset of parameters in LLMs (in some cases, a single parameter) can exert a disproportionate influence on an LLM’s overall functionality (see Figure 1 ). This work highlights the critical role of these “super weights” and their corresponding “super activations,” offering a new insight into LLM architecture and avenues for efficient model compression. The paper provides full technical details and experimental results; in this post, we provide a high-level overview of the key findings and their implications.
Understanding and Compressing Increasingly Large Models
While LLMs demonstrate impressive capabilities, their sheer size, often comprising billions or even hundreds of billions of parameters, presents significant challenges for deployment on resource-constrained hardware such as mobile devices. Reducing the size and computational complexity of LLMs for such platforms leads to corresponding reductions in memory and power consumption, enabling them to operate locally, privately, and without an internet connection. However, understanding the internal mechanisms of LLMs is critical, as naïve compression or simplification can lead to substantial degradation in model quality.
Identifying Super Weights and Their Impact
Prior research indicated that a small percentage of parameter outliers in LLMs are vital for maintaining model quality — and if these weights are significantly modified (through compression) or removed entirely (pruned) then the model’s output quality suffers. While this prior work showed that this fraction can be as small as 0.01% of the weights, in models with billions of parameters, this still translates to hundreds of thousands of individual weights. In this work, Apple researchers identified a remarkably small number of parameters, termed “super weights,” that if altered, can destroy an LLM’s ability to generate coherent text, for example, leading to a threefold order of magnitude increase in perplexity and reducing zero-shot accuracy to levels consistent with random guessing. For instance, in the Llama-7B model, removing its single super weight renders the model incapable of producing meaningful output. Conversely, removing thousands of other outlier weights, even those with larger magnitudes than the super weight, results in only marginal quality degradation.
This work proposes a methodology for locating these super weights by requiring only a single forward pass through the model. This method leverages the observation that super weights induce correspondingly rare and large activation outliers, which we term “super activations.” These super activations often appear after the super weight, persist throughout subsequent layers with constant magnitude and position, irrespective of the input prompt, and their channel aligns with that of the super weight. By detecting spikes in the input and output activation distributions of specific model components (e.g., the down projection of the feed-forward network), we can locate the super weights via their corresponding super activation. Intriguingly, the super weight is consistently found in the down projection of the feed-forward network following the attention block, typically in an early layer of the network. We have compiled an index of super weight coordinates for several common, openly available LLMs to facilitate further investigation by the research community.
Table 1: The above layer numbers, layer types, and weight types can be directly applied to
Huggingface models. For example, for Llama-7B on Huggingface, access the super weight using layers[2].mlp.down_proj.weight[3968, 7003] .
As shown in the coordinates table (see Table 1 ), super weights emerge in specific projection layers, often early in the network across a wide range of commonly used LLMs. These weights generate a super activation that then persists through the residual skip connections in the network as illustrated in Figure 2 . This persistent super activation exerts a global influence on the model’s internal dynamics, biasing it away from generating high-probability stopwords. When super weights are removed, this suppressive effect vanishes, and the model’s output distribution shifts sharply: the likelihood of stopwords increases significantly, while meaningful, content-bearing tokens become less probable. This suggests that super weights play a critical role in determining which semantically meaningful tokens are output during the forward pass of the model.
Enhanced Compression and Model Understanding
The discovery of super weights and super activations can lead to improvements in LLM compression and the field’s broader understanding of these models. The large impact of these few parameters suggests that their preservation is critical during LLM compression techniques. We found that by preserving super activations with high precision, simple round-to-nearest quantization methods can achieve performance competitive with more sophisticated state-of-the-art techniques. Similarly, for weight quantization, preserving the super weight while clipping other weight outliers allows round-to-nearest quantization to be effective even with much larger block sizes than previously thought feasible, leading to better compression ratios.
This work demonstrates that handling just a few super outliers can significantly improve compression quality, offering a hardware-friendly approach compared to methods that manage hundreds of thousands of outlier weights. This targeted approach can lead to more efficient models that retain a higher degree of their original performance. This in turn enables powerful LLM applications to operate with high quality on resource constrained hardware, such as mobile devices.
Exploring the Landscape of Super Outliers
Our findings open several avenues for future research. Further exploration into the genesis and precise mechanisms of super weights and super activations could yield deeper insights into the operational dynamics of LLMs. Understanding how these specific parameters acquire such disproportionate influence during training could inform future model design and training strategies. Investigating the prevalence and characteristics of super weights across a broader array of model architectures and training paradigms can shed light on their role/creation, and the provided directory of super weights aims to spur such continued investigation within the community. Ultimately, a more comprehensive understanding of these super outliers holds the potential to unlock new methodologies for building more efficient, robust, and interpretable LLMs.
The Super Weight in Large Language Models
July 2, 2025 research area Speech and Natural Language Processing
Recent works have shown a surprising result: a small fraction of Large Language Model (LLM) parameter outliers are disproportionately important to the quality of the model. LLMs contain billions of parameters, so these small fractions, such as 0.01%, translate to hundreds of thousands of parameters. In this work, we present an even more surprising finding: Pruning as few as a single parameter can destroy an LLM’s ability to generate text —…
R^2: Range Regularization for Model Compression and Quantization
August 10, 2023 research area Computer Vision , research area Speech and Natural Language Processing
Model parameter regularization is a widely used technique to improve generalization, but also can be used to shape the weight distributions for various purposes. In this work, we shed light on how weight regularization can assist model quantization and compression techniques, and then propose range regularization ( R 2 R^2 R 2 ) to further boost the quality of model optimization by focusing on the outlier prevention.
By effectively…
Discover opportunities in Machine Learning.
Our research in machine learning breaks new ground every day.
The “Super Weight:” How Even a Single Parameter can Determine a Large Language Model’s Behavior

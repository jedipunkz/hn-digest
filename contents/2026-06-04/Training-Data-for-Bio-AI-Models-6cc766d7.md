---
source: "https://research.dimensioncap.com/p/on-training-data-for-bio-ai-models"
hn_url: "https://news.ycombinator.com/item?id=48396340"
title: "Training Data for Bio AI Models"
article_title: "On Training Data for Bio AI Models - by Bauer LeSavage"
author: "sebg"
captured_at: "2026-06-04T10:22:01Z"
capture_tool: "hn-digest"
hn_id: 48396340
score: 1
comments: 0
posted_at: "2026-06-04T09:52:00Z"
tags:
  - hacker-news
  - translated
---

# Training Data for Bio AI Models

- HN: [48396340](https://news.ycombinator.com/item?id=48396340)
- Source: [research.dimensioncap.com](https://research.dimensioncap.com/p/on-training-data-for-bio-ai-models)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T09:52:00Z

## Translation

タイトル: バイオ AI モデルのトレーニング データ
記事タイトル: バイオ AI モデルのトレーニング データについて - Bauer LeSavage 著
説明: 生命科学、機械学習、テクノロジーにおける最先端の進歩と、それらの融合が地球上の生命の軌道をどのように再形成しているかについて書いています。

記事本文:
バイオ AI モデルのトレーニング データについて - Bauer LeSavage 著
バイオ AI モデルのトレーニング データについて
生物学的基礎モデルを進歩させるにあたり、LLM データキュレーションの転送から得られる教訓はどれでしょうか?また、再考が必要なものはどれでしょうか?
Bauer LeSavage 2026 年 6 月 3 日 16 1 シェア この記事の執筆にあたり、洞察力に富んだ会話をしてくださった Chris Gibson、Nathan Frey、Peyton Greenside、Ron Alfa、その他数名に特別に感謝します。
今年の AI における最大のトレンドの 1 つは、外部パートナーが生成した独自のトレーニング データに対するフロンティア ラボによる驚異的な支出です。研究所の支出指標が公に公開されることはほとんどありませんが、データラベル会社自体からの数字を使用して推定することができます。 2025 年 6 月に遡ると、最大手の Surge AI と Scale AI は、主にフロンティア ラボからの支出によってそれぞれ平均約 10 億ドルの収益を上げていました。
大手データラベル会社の収益。 (企業開示からのデータ; 出典 ) より大きなデータ市場に当てはめると、トップの研究機関は複数のデータプロバイダーからの独自のデータセット生成に年間 10 ～ 100 億ドルを費やしていると推定されています。オーダーメイドのデータセットと強化学習 (RL) 環境が、専門分野全体でモデルのパフォーマンスを向上させるための中核的なアクセラレータとして定着するにつれて、これらの数値は増加するばかりであるように見えます。
大規模言語モデル (LLM) がどのように改善されるかの軸を (1) 関連するトレーニング データへのアクセス、(2) モデル アーキテクチャの進歩、(3) コンピューティングの増加に単純化すると、外部データセットを購入するこの傾向は、いくつかの理由から理にかなっています。
過去数年間にわたるデータ収集の取り組み（インターネットをスクレイピングする Common Crawl、出版社からのデータライセンスなど）の組み合わせにより、人間が生成した既存のすべてのデータの枯渇が加速しています。

テキストデータ。
モデル アーキテクチャの大きな進歩は、研究室間の人材の頻繁な異動や研究への多額の投資により、時間の経過とともにコモディティ化されることが多く (変圧器など)、AI 研究室にとってアーキテクチャの信頼性が低くなります。
コンピューティングの制約は、モデルをトレーニングするために必要なコンピューティングではなく、モデルを提供する需要の増大によってますます支配されるようになってきています。
このデータ需要は、コーディング、金融、コンピューターの使用など、注目を集めるアプリケーション ドメイン全体にわたって最も広く現れています。ロジックはかなり単純で、プレイブックは機能しているようです。より厳選された Python コードのサンプルをトレーニング データに含めることで、LLM がより優れたソフトウェア エンジニアになる可能性があります。
ライフサイエンス基礎モデルのトレーニングに使用される生物学的データセットについても、同様の傾向が議論されることが増えており、モデル開発者はタスク関連データへのアクセスに関してより多くのライセンス契約を結んでいます。トレーニング データにさらに高品質の抗体親和性測定を含めることで、生成タンパク質モデルがより優れたタンパク質エンジニアになる可能性があります。
この需要に応えるために、生物学的データセットを収集する速度と厳密さを向上させることにも注目が集まっています。自律型ライフサイエンスロボット企業は現在、データ収集、モデルトレーニング、評価の間のフライホイールを閉じるために、この成長する市場を利用するデータファウンドリとして自らを売り込んでいます。
生物学的データを収集、販売、購入、トレーニングする機会が増えていることで、いくつかの重要な疑問が生じています。生物学的基礎モデルのデータセットをキュレーションしてトレーニングするための適切な戦略は何でしょうか。また、テキスト データを使用して LLM をトレーニングする方法からどのような原則が移行され、どのような原則が移行されないのでしょうか。
フロンティアLLMと生物学的基盤のMOの間にマクロレベルの市場の類似点があるにもかかわらず

デルには、深く議論する価値のある意味のある微妙な違いがあります。
データの観点から見ると、LLM トレーニングの歴史的なプレイブック (主にデータ スケールに焦点を当てたもの) は、生物学的データセットに 1 対 1 で転送されません。その代わりに、生物学的トレーニング データセットのキュレーションには、量よりもデータの質を優先するというはるかに注意が必要であり、スケールマックスの罠に陥る前にこのことを認識する必要があると私は主張します。
最新の LLM 事前トレーニングの初期の頃、大多数のコンセンサスは、スケールを通じてモデルのパフォーマンスを向上させることでした。より多くのデータ、より多くのパラメーター、より多くのコンピューティング。これは、現場が事前トレーニングのスケーリング則、つまりデータ、モデル サイズ、コンピューティングなどの関数としてモデルのパフォーマンスがどのように向上するかを列挙することに多くの時間を費やしていることで明らかになりました。このトピックについては、以前の Dimension Research の記事で詳しく書きました。
2020 年に、カプランらはOpenAI は、初期の事前トレーニング スケーリング則の一部を公開し、データ (トークン) よりも速くモデル サイズ (パラメーター) を増やすと、固定のコンピューティング バジェットで最適なパフォーマンスが得られることを示しました。 2022 年に、ホフマンらはDeepMind は、「チンチラ」スケーリング則を発表しました。これは、カプラン ベースのモデルが十分にトレーニングされていないこと、およびより多くのデータ (パラメーターあたり最大 20 トークン) でより小さなモデルをトレーニングすると、モデルのパフォーマンスが向上することを示しました。ほぼ 2 年後、Sardana らはDatabricks では、推論の需要を考慮したチンチラのスケーリング則の修正を公開しました。コンピューティング予算が固定されている場合、目的のパフォーマンスを達成するには、より多くのデータ (パラメーターあたり約 200 ～ 10,000 トークン) を使用して小規模なモデルをトレーニングする方が経済的です。これは、小規模なモデルを提供する方が最終的には顧客にとって安価であるためです。 ( 出典 ) この初期の時代からの一般的なポイント h

おそらく多くの AI 研究者の心に焼き付いているのは、「データが多ければ、より良い結果が得られる」ということです。人間は何世紀にもわたって、LLM が本質的に無料で継承した高品質の文章 (インターネットからの抜粋、出版文献など) の大規模なバンクを収集してきたため、データの品質よりもスケールが優先される可能性があります。継承されたデータセットには、有用なモデルのパフォーマンスを可能にするのに十分な両方の要素が含まれていたため、最初からデータの品質と量の間の標準的なトレードオフを回避することができました。
研究者は、高レベルのフィルタリング アプローチ (重複排除、汚染除去、初期品質の分類子など) を使用しながらスケールを活用して、低品質のサンプルを削除できます。テキスト データは非常に冗長であるため、ノイズはさらに許容され、ランダムなエラーや矛盾は、すでに存在していた規模で平均化されます。モデルが独自の高品質の合成データを生成できるようになると、トレードオフはさらに減少しました。
現場では、ここ数年にわたってこの戦略のリバランスが見られ始めています。利用可能なデータが枯渇するにつれて、高品質のデータセットに対する評価が事前トレーニング中に引くもう 1 つの重要な手段となり、研究者はこれを大規模に可能にする、より積極的なデータ優先順位付けツールの開発を開始しました。
重要度リサンプリングによるデータ選択 (DSIR) — 大規模なデータセットから、高品質データのサブ選択を定義します。次に、完全なデータセットの広範な分布からこの「高品質の署名」をサンプリングして、他のすべての高品質の例を明らかにします。
Minimax Optimization を使用したドメイン再重み付け (DoReMi) — さまざまなデータ型 (例: 一般知識 50%、コーディング 25%、科学 15%、推論トレース 10%) のさまざまな重みを使用して小規模モデルをトレーニングし、最もパフォーマンスの高い分布を見つけます。このディストリビューションをスケールする

最適な組み合わせを見つけたら、より大きなモデルまで検討してください。
データ影響モデルによるモデル認識データ選択 (MATES) — プライマリ モデルの事前トレーニングの各段階で最も効果的なデータを動的に選択して、モデルが認識するデータの順序を最適化できる小規模な「データ影響モデル」をトレーニングします。
品質分類子 ( FineWeb 、 DCLM 、 Nemotron-CC ) — 個別の分類子モデルをトレーニングして、ユーザー定義のメトリクスに基づいてデータ品質をスコアリングし、品質しきい値を満たさないものをフィルタリングします。小規模だが高品質のデータセットは、さまざまな品質を持つ大規模なデータセットよりもパフォーマンスが向上する可能性があります (たとえば、7B DCLM モデルは、トレーニング データが約 6 倍少ない 8B Llama 3 モデルと同様に実行されます)。
(A) フィルタリングに品質分類子モデルを使用して、元の Common Crawl データ ソースから DCLM トレーニング データセットを構築します。 (B) フィルタリングされた DCLM-Baseline データセットを使用してトレーニングされたモデルは、より大きなトレーニング予算を使用したモデルよりも優れたパフォーマンスを示しました。 (図は出典 から引用) これらのアプローチはすべて、LLM にとってデータ品質の新たな重要性を示していますが、そのほとんどは、フィルタリング対象となる、明確に定義された高品質データの大規模なコーパスが存在することに依存しています。数年前は LLM にとって問題ではありませんでしたが、私たちは既存のデータからほぼすべてを搾り取っています。
フロンティア ラボがアプリケーション固有のデータに巨額の支出を行っていることは、ドメイン全体にわたる高品質のデータセットが大きなボトルネックになっているというさらなる証拠です。
ありがたいことに、私たちはより多くのデータ収集に一般的にアプローチするための適切なレシピを持っています。関心のある多くの分野では、データの品質と関連するベンチマークを検証するのが比較的簡単です (例: コードは正しい答えを出力し、数学的分析は正しい閉形式の答えを出力し、テニスの歴史に関するレポートは事実に基づいており、人間によって解釈可能です)

n)。上記の高品質フィルター ツールの一部を再利用し、その上に構築して、どの新しいデータを外部から調達するかを優先することもできるかもしれません。
残念ながら、LLM 用のこのデータ フレームワークは生物学の世界に直接移植されません。
なぜ生物学がこの戦略を破るのか
生物学的基礎モデルは、幅広いサブドメインとアプリケーションにわたって急速に進歩しています。生体分子モデルは完全長の高親和性モノクローナル抗体をコンピュータで生成しており、ゲノム基礎モデルは変異体の影響を予測して現実的な DNA 配列を設計し始めており、組織レベルの世界モデルは空間的なタンパク質発現と患者の薬物反応を予測するために開発されています。
これらの成果の一部は、過去数年間にわたる深層学習と LLM のブレークスルーに遡ることができます。テキスト モデルからインスピレーションを得た多くの初期アーキテクチャとスケーリングの直感は、その後、公開されている生物学的データに適用されました。
ただし、このプレイブックが実際に機能するのは、データが適切に厳選されている場合のみです。標準的な例は AlphaFold です。これは、タンパク質データ バンク (PDB) (数十年にわたって収集された約 200,000 個のタンパク質構造の高品質実験データセット) での事前トレーニングがどのように構造予測において驚異的なパフォーマンスにつながるかを紹介しました。
しかし、PDB は標準というよりは例外です。フロンティア LLM をトレーニングするためのテキスト データとは異なり、私たちが公開している生物学的データははるかにまばらであり、モデルのトレーニングに使用することを困難にするいくつかの固有かつ意味のある制限に悩まされています。その結果、パフォーマンスの高いバイオ基礎モデルをトレーニングするには、多くの場合、新しいデータを合成する必要があり、(バイオ) フロンティア研究室は、データにアクセスするために LLM で見ているものと同様のデータ調達戦略を検討し始めています。
そんな世界において、最も重要な訓練データとは

特定のアプリケーション用に生成すると便利ですか?生物学的データ収集に関する私たちの哲学はどのようなものであるべきでしょうか?ライフサイエンスでは、規模と品質の間のトレードオフがより明らかですが、この 2 つのバランスを適切に取るにはどうすればよいでしょうか?
これらの質問に答えるには、「高品質」の意味をより正式に定義し、テキスト トレーニング データと比較した現在の生物学的データセットの限界を理解する必要があります。ここで私は品質を、(1) コンテキストが豊富であること、(2) クリーンであること (つまり、ノイズが低減され、再現可能であること)、(3) 多様性があること、(4) グラウンド トゥルースを学ぶために特化されていること、と定義します。
多かれ少なかれ人類の歴史に記録されたすべての規模に匹敵するのは困難です。メトリクスが公的に報告されている最大規模の LLM の一部は、フィルタリング後の数十兆のトークンでトレーニングされています。しかし、私の主張は単に生物学的データよりもテキスト データの方が多いということではありません。私たちは実際に、数十兆のヌクレオチド塩基、数億のタンパク質配列、数億の単一細胞トランスクリプトームプロファイル、および数十万の実験タンパク質構造など、かなりの量の生物学的データを持っています。
過去 10 年間で、トップ AI モデルのトレーニング データセットのサイズが大幅に増加しました。 ( 出典 ) scRNA-seq データセットのサイズ (細胞数) が増加しました

[切り捨てられた]

## Original Extract

We write about frontier advances in the life sciences, machine learning, and technology, and how their fusion is reshaping the trajectory of life on Earth.

On Training Data for Bio AI Models - by Bauer LeSavage
Subscribe Sign in On Training Data for Bio AI Models
As we advance biological foundation models, which lessons from LLM data curation transfer, and which need rethinking?
Bauer LeSavage Jun 03, 2026 16 1 Share Special thank you to Chris Gibson, Nathan Frey, Peyton Greenside, Ron Alfa, and several others for insightful conversations throughout the writing of this article.
One of the largest trends in AI this year has been the incredible spend by frontier labs on proprietary training data generated by external partners. While the labs’ spending metrics are rarely publicly disclosed, they can be estimated using numbers from the data labeling firms themselves. Back in June 2025, the largest players, Surge AI and Scale AI, were averaging ~$1B in revenue each , driven mainly by spend from the frontier labs.
Revenues for top data labeling companies. (Data from company disclosures; Source ) Extrapolating to the larger data market, it’s been estimated that top labs are shelling out $1-10B annually on proprietary dataset generation from several data providers. These numbers seem to only be increasing as bespoke datasets and reinforcement learning (RL) environments become cemented as core accelerants for improving model performance across disciplines.
If you simplify the axes of how large language models (LLMs) improve into (1) access to relevant training data, (2) advancements in model architecture, and (3) increased compute, this trend toward purchasing external datasets makes sense for a few reasons.
The combination of data collection efforts over the last several years ( e.g., Common Crawl scraping the internet, data licensing from publishers) has accelerated the exhaustion of all existing human-generated text data.
Major advancements in model architectures are often commoditized over time ( e.g., transformers) due to frequent talent transfer between labs and large investment into research, making architecture a less reliable moat for AI labs.
Compute constraints are increasingly governed by growing demand to serve models rather than by the compute needed to train them.
This data demand has manifested most widely across hot application domains like coding, finance, and computer use. The logic is rather simple, and the playbook seems to be working. Including more curated examples of Python code in your training data will likely make the LLM a better software engineer.
I am increasingly seeing these same trends discussed for the biological datasets used to train life science foundation models, with model developers making more licensing deals for access to task-relevant data. Including more high-quality antibody affinity measurements in your training data will likely make the generative protein model a better protein engineer.
There has also been a growing focus on improving the speed and rigor with which we collect these biological datasets to feed this demand. Autonomous life science robotics companies are now pitching themselves as data foundries to take advantage of this growing market in an effort to close the flywheel between data collection, model training, and evaluation.
These growing opportunities to collect, sell, buy, and train on biological data raise some important questions: what’s the right strategy for curating and training on datasets for biological foundation models, and what principles do and do not transfer from how we approached training LLMs with text data?
Despite the macro-level market parallels between frontier LLMs and biological foundation models, there are meaningful and nuanced differences that are worth discussing in depth.
From a data perspective, the historical playbook for LLM training — a predominant focus on data scale — will not transfer one-to-one to biological datasets. Instead, I argue that curating biological training datasets demands far more care in prioritizing data quality over quantity, and we should recognize this before we fall into the trap of scalemaxxing .
In the early days of modern LLM pretraining, the predominant consensus was to improve model performance through scale. More data, more parameters, more compute. This manifested with the field spending much of its time enumerating pretraining scaling laws — the slope of how model performance improved as a function of data, model size, compute, etc. We wrote about this topic in depth in a previous Dimension Research article .
In 2020, Kaplan et al. at OpenAI published some of the initial pretraining scaling laws, showing that increasing model size (parameters) faster than data (tokens) resulted in optimal performance for a fixed compute budget. In 2022, Hoffmann et al. at DeepMind published the “Chinchilla” scaling laws, which showed that Kaplan-based models were undertrained and that better model performance was achieved when training smaller models with more data (~20 tokens per parameter). Almost two years later, Sardana et al. at Databricks published a modification to Chinchilla scaling laws to take inference demand into account. With a fixed compute budget, it’s more economical to train smaller models with increasingly more data (~200-10,000 tokens per parameter) to achieve a desired performance because serving smaller models is ultimately cheaper for customers. ( Source ) The general takeaway from this early period has likely been seared into the minds of many AI researchers: “More data gives you better results.” Scale could take priority over data quality because humans had spent centuries assembling a large bank of high-quality writing ( e.g., excerpts from the internet, published literature) that LLMs inherited essentially for free. They were originally able to sidestep the canonical trade-off between data quality and quantity from the start because the inherited dataset contained enough of both to enable useful model performance.
Researchers could leverage scale while using high-level filtering approaches ( e.g., deduplication, decontamination, early quality classifiers) to remove low-quality examples. Noise was further tolerated because text data is highly redundant, so random errors and contradictions average out at the scale we already had. Once models could generate their own high-quality synthetic data, the trade-off was reduced even more.
The field has started to see this playbook rebalance over the last couple of years. As we ran out of available data, the appreciation for high-quality datasets became another prominent lever to pull during pretraining, and researchers began to develop more aggressive data prioritization tools to enable this at scale.
Data Selection with Importance Resampling (DSIR) — From a large dataset, define a sub-selection of high-quality data. Then sample this “high-quality signature” from the broader distribution of the full dataset to surface all the other high-quality examples.
Domain Reweighting with Minimax Optimization (DoReMi) — Train small models with varying weights of different data types ( e.g., 50% general knowledge, 25% coding, 15% science, 10% reasoning traces) to find the most performant distribution. Scale this distribution up to larger models once you find the optimal mix.
Model-Aware Data Selection with Data Influence Models (MATES) — Train a small “data influence model” that can dynamically select the most effective data for each stage of primary model pretraining to optimize the order of data the model sees.
Quality Classifiers ( FineWeb , DCLM , Nemotron-CC ) — Train separate classifier models to score data quality based on user-defined metrics and filter anything that doesn’t meet a quality threshold. Smaller but higher-quality datasets can be more performant than a larger dataset with mixed quality ( e.g., a 7B DCLM model performed similarly to an 8B Llama 3 model with ~6x less training data).
(A) Construction of the DCLM training dataset from the original Common Crawl data source using a quality classifier model to filter. (B) Models trained with the filtered DCLM-Baseline dataset outperformed models with larger training budgets. (Figure adapted from Source ) These approaches all showcase the newfound importance of data quality for LLMs, but most of them rely on having an existing large corpus of well-defined, high-quality data to filter from. Not a problem for LLMs a couple of years ago, but we are milking nearly everything we can from existing data.
The huge spending from frontier labs on application-specific data is further evidence that quality datasets across domains have become a major bottleneck.
Thankfully, we have a decent recipe for how to generally approach more data collection. For many areas of interest, it’s relatively easy to verify data quality and associated benchmarks ( e.g., the code outputs the right answer, the mathematical analysis outputs the correct closed-form answer, the report on the history of tennis is factual and interpretable by a human). We may even be able to repurpose and build upon some of the above quality filter tools to prioritize which novel data to procure externally.
Unfortunately, this data framework for LLMs does not directly port to the world of biology.
Why biology breaks this playbook
Biological foundation models are rapidly progressing across a wide range of subdomains and applications. Biomolecular models are generating full-length, high-affinity monoclonal antibodies in silico , genomic foundation models are starting to predict variant effects and design realistic DNA sequences, and tissue-level world models are being developed to predict spatial protein expression and patient drug responses.
These achievements can be partially traced back to breakthroughs in deep learning and LLMs over the last several years. Many initial architectures and scaling intuitions inspired by text models were subsequently applied to public biological data.
But this playbook only really works when the data is well-curated. The canonical example is AlphaFold , which showcased how pretraining on the Protein Data Bank (PDB) — a high-quality experimental dataset of ~200k protein structures curated over decades — could lead to incredible performance in structural prediction.
But the PDB is more of an exception than the norm. Unlike text data for training frontier LLMs, our publicly available biological data is much more sparse and suffers from several unique and meaningful limitations that make it difficult to use for model training. As a result, training performant bio foundation models often requires new data to be synthesized, and (bio) frontier labs are beginning to explore a data procurement strategy similar to the one we’re seeing with LLMs to access it.
In such a world, what training data is most useful to generate for a given application? What should be our philosophy on biological data collection? In the life sciences, where the trade-off between scale and quality is more apparent, how do we appropriately balance the two?
Answering these questions requires a more formal definition of what I mean by “high-quality” and an understanding of the limitations of current biological datasets compared to text training data. Here, I define quality as (1) context-rich, (2) clean ( i.e., reduced noise and reproducible), (3) diverse, and (4) purpose-built to learn ground truth.
It’s hard to compete with the scale of more or less all recorded human history. Some of the largest LLMs with publicly reported metrics are trained on tens of trillions of tokens post-filtering. But my argument is not simply that we have more text data than biological data. We actually have a decent amount of biological data — tens of trillions of nucleotide bases, hundreds of millions of protein sequences, hundreds of millions of single-cell transcriptomic profiles, and hundreds of thousands of experimental protein structures.
Training dataset size has grown considerably for top AI models over the last decade. ( Source ) scRNA-seq dataset size (number of cells) has grown

[truncated]

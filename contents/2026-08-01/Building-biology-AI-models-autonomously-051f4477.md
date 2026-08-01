---
source: "https://phylo.bio/blog/biomni-tuso"
hn_url: "https://news.ycombinator.com/item?id=49133734"
title: "Building biology AI models autonomously"
article_title: "Building state-of-the-art biology AI models autonomously | Phylo Blog"
author: "eamag"
captured_at: "2026-08-01T12:59:16Z"
capture_tool: "hn-digest"
hn_id: 49133734
score: 2
comments: 0
posted_at: "2026-08-01T12:11:32Z"
tags:
  - hacker-news
  - translated
---

# Building biology AI models autonomously

- HN: [49133734](https://news.ycombinator.com/item?id=49133734)
- Source: [phylo.bio](https://phylo.bio/blog/biomni-tuso)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T12:11:32Z

## Translation

タイトル: 生物学 AI モデルを自律的に構築する
記事のタイトル: 最先端の生物学 AI モデルを自律的に構築する |ファイロのブログ
説明: 我々は、生物学におけるメソッド開発のための完全に自律的でコード不要のツール、Biomni x TusoAI を紹介します。これは、自然言語の単一文から遺伝的摂動予測とエンハンサー遺伝子リンクのための最先端のモデルを構築します。

記事本文:
最先端の生物学 AI モデルを自律的に構築 | Phylo ブログ Phylo が Biomni Lab と、A16Z と Menlo Ventures が共同主導するシードラウンドを発表 続きを読む →
について 採用情報 料金 バイオムニ ラボ バイオムニ ラボ バイオムニ ラボの使用例 企業研究ブログ 会社
最先端の生物学 AI モデルを自律的に構築
Biomni x TusoAI は、生物学における手法開発のための完全に自律的でコード不要のツールであり、自然言語の 1 文から遺伝的摂動の予測とエンハンサー遺伝子のリンクのための最先端のモデルを構築します。
計算手法は生物学のあらゆる分野で必要とされており、発見の中心となっていますが、その発展が遅いため、膨大な量の実験データが十分に活用されておらず、十分に理解されていないままになっています。配列決定コストの低下により生物学がますます測定可能になったのと同様に、これらの測定を解釈可能にするためには、方法開発に対するスケーラブルなアプローチが必要です。
Autoresearch や AlphaEvolve などのエージェント手法により、科学コードや ML コードの改善が自動化され始めています。これらのツールの普及を妨げる主な障壁は、その使いやすさです。これらは多くの場合、クローズドソースであるか、実行に数百ドルかかる高価であり、最適なパフォーマンスを達成するには、コンピューティング インフラストラクチャ、コード、データ、テキスト、パラメーター入力のための膨大なセットアップが必要です。生物学はまた、問題を解決するために大量のコード リポジトリと有用なコードだけでなく有用なデータの同時検索を扱うという、メソッド開発に特有の課題を提示します。
生物学におけるメソッド開発を効果的に自動化するために、完全に自律的でコード不要のメソッド開発ツールである Biomni x TusoAI を紹介します。私たちは、メソッド開発用のエージェント ツール TusoAI (以前は ICLR で公開されました) を拡張してきました。

生物学特有の課題に対応するためです。 TusoAI はエージェントを調整して、人間の監視なしでコードベース全体を何千回も効率的かつ自律的に進化させ、新しいアルゴリズムとデータの革新を備えたよりパフォーマンスの高いバージョンを返します。 TusoAI は、文献をマイニングし、改善の実装、デバッグ、評価を繰り返し行うことで、実際の計算生物学者を模倣しますが、よりシンプルな方法を優先し、カスタム構築されたエージェント ハーネスを利用して反復全体でメモリを保持します。ここで、Biomni は、API キーを指定して TusoAI 自体を自動的にセットアップして実行することも、TusoAI のメソッド開発サイクル自体をエミュレートすることもできます (API キーは必要ありません)。
Biomni x TusoAI のいくつかの利点を強調します。
データとコードを同時に自律的に検索します
1 ドルの支出ごとに何百回も反復される
効率的なメソッドの生成を保証
環境設定、データ設定、コーディングは一切必要ありません。
メソッドは完全に自然言語の 1 文で開発できるようになりました。これは、生物学においてこれまでに行われた中で最もアクセスしやすい方法の開発です。
以下では、Biomni x TusoAI が (1) エンハンサーと遺伝子の結合および (2) 遺伝的摂動予測のための新しい強力なモデルを開発していることを示します。この成果は今後数か月以内に公開される予定で、新しい TusoAI は github.com/Alistair-Turcan/TusoAI で早期にオープンソースで入手できます。 Biomni x TusoAI へのアクセスは、biomni.phylo.bio の [スキル] タブから今日から利用できます。
図 1: Biomni を使用して TusoAI を調整します。
Biomni x TusoAI のアプリケーション
1. 最先端の遺伝的摂動予測
CRISPR 遺伝子ノックアウトのトランスクリプトーム効果の予測は、数百のグループが参加する毎年恒例のバーチャル セル チャレンジなど、大きな関心を集めている有望ではあるものの難しい課題です。このタスクは難しいです

設定が不明確で、評価が一致せず、使用する適切なデータが不明であるためです。
Biomni x TusoAI を 3 つの独立したベンチマークに適用して、新しい摂動予測手法をゼロから構築します。 Biomni は、遺伝子埋め込み、トレーニングおよび評価戦略、およびこれらすべてをリンクするモデルを検索し、データの小さな検証サブセットのパフォーマンスを向上させます。 1 週間自律的に反復し、500 の候補をテストした後、Biomni x TusoAI は、9 つ​​の遺伝子埋め込みソースと単純回帰および kNN のアンサンブルを統合した新しい手法を発見しました。これは、ほぼすべてのメトリクスとデータセットにおいて、大規模な深層学習モデルを含むあらゆる手法よりも優れたパフォーマンスを発揮します。
おそらく驚くべきことかもしれませんが、この単純な方法は文献では完全に理にかなっています。以前の研究では、より豊富な遺伝子埋め込みがより強力なパフォーマンスにつながることを特定していましたが、個々のデータ ソースを注意深く調整せず、代わりに同じ埋め込み手順を数十のデータ ソースに適用していました。ここで、Biomni x TusoAI はこのパラダイムをさらに推し進め、単純なモデルが摂動効果を学習できるようにするカスタマイズされた遺伝子埋め込み戦略を発見しました。この新しいメソッドは、github.com/phylobio/TusoPerturb で公開されています。
図 2: Biomni x TusoAI によって構築された新しい遺伝的摂動モデルのパフォーマンス。
2. エンハンサー遺伝子結合の新たな特徴の発見
エンハンサー遺伝子の関連性の同定はヒト遺伝学の中心的な課題となっており、ENCODE コンソーシアムを含むいくつかの大規模な取り組みによって追求されています。最先端の手法は、よく知られたパラダイムに従っています。つまり、DNA 配列、エピゲノム トラック、Hi-C コンタクト、マルチオーム データなどから大量のデータを収集し、このデータからエンハンサーを遺伝子に結び付けるのに役立つ一連の特徴を手動で構築します。
に

エンハンサーと遺伝子のリンク方法の開発を自動化するために、私たちは TusoAI を適用して、大規模なデータ コーパスから何千もの特徴を構築してテストしました。 TusoAI は、主要なメソッドである pgBoost と 36 の異なるゲノム データ ソースから始めて、これらのデータから機能を繰り返し構築およびテストし、pgBoost を改善します。 10,000 回の反復の後、TusoAI は、14 の異なるデータ ソースから追​​加の 23 の新しい機能を統合した新しいモデルを構築しました。この新しいモデルは、困難な遠位 (>100KB) リンクでの 5 倍の改善を含め、グラウンド トゥルース eQTL 評価においてすべての方法を大幅に上回りました。同様の優れたパフォーマンスが、グラウンド トゥルース CRISPR および GWAS 評価でも観察されました。
最も重要な 2 つの新機能は、SNP-TSS 距離を使用する一般的なパラダイムを高度に設計したバージョンです。 SNP-TSS 距離を単純に入力する代わりに、これら 2 つはこれを鎖と遺伝子の重複を認識したシグモイド変換共変量として再設計します。バイナリイントロンフラグ、SNP 周囲の GC コンテンツ、および操作された ABC スコアなど、その他の機能は pgBoost にとってまったく新しいものです。
図 3: TusoAI とのエンハンサー遺伝子リンクの改善。さまざまな距離しきい値にわたるグラウンド トゥルース eQTL の平均濃縮度。 pgBoost (TusoAI) は、pgBoost の TusoAI 改良版を指します。
biomni.phylo.bio でお気に入りのタスクを入力して、今すぐ試してみてください。 Biomni-Tuso は、企業ユーザー向けの Biomni Lab の自動リサーチ スキルとして利用できるようになりました。パブリック プラットフォームの場合、この機能を広く展開する準備として、100 人の Biomni ラボ ユーザーに 500 GPU 時間を共有します。ここからお申し込みください。
生物学者がより多くの発見をより迅速に行えるように設計された共同エージェント。
© 2026 Phylo, Inc. 無断複写・転載を禁じます。
© 2026 Phylo, Inc. 無断複写・転載を禁じます。

## Original Extract

We present Biomni x TusoAI, a fully autonomous and code-free tool for method development in biology — building state-of-the-art models for genetic perturbation prediction and enhancer-gene linking from a single sentence of natural language.

Building state-of-the-art biology AI models autonomously | Phylo Blog Phylo unveils Biomni Lab and seed round co-led by A16Z and Menlo Ventures Read more →
About Careers Pricing Biomni Lab Biomni Lab Biomni Lab Use Cases Enterprise Research Blog Company
Building state-of-the-art biology AI models autonomously
We present Biomni x TusoAI, a fully autonomous and code-free tool for method development in biology — building state-of-the-art models for genetic perturbation prediction and enhancer-gene linking from a single sentence of natural language.
Computational methods are needed across all fields of biology and are central to discovery, but their slow development leads to a vast amount of experimental data remaining underutilized and poorly understood. Just as the lowering cost of sequencing has made biology increasingly measurable, a scalable approach to method development is needed to make these measurements interpretable.
Agentic methods like Autoresearch or AlphaEvolve have begun to automate scientific and ML code improvement. A key barrier to the widespread adoption of these tools is their usability. These are often closed-source or expensive — costing hundreds of dollars to run, and require a huge amount of setup for compute infrastructure, code, data, text, and parameter input to achieve optimal performance. Biology also presents unique challenges to method development, dealing with massive code repositories and simultaneous search over not only useful code but also useful data to solve a problem.
To effectively automate method development in biology, we present Biomni x TusoAI, a fully autonomous and code-free tool for method development. We have been expanding upon our agentic tool for method development TusoAI (previously published at ICLR) to cater it to the specific challenges of biology. TusoAI coordinates agents to efficiently and autonomously evolve entire codebases thousands of times without human oversight, returning a higher performing version with new algorithmic and data innovations. TusoAI mimics a real computational biologist by mining literature and iteratively implementing, debugging, and evaluating improvements, while favoring simpler methods and retaining memory across iterations, powered by a custom-built agent harness. Here, Biomni can both automatically set up and run TusoAI itself given an API key, or emulate TusoAI’s method development cycle itself — no API key needed.
We highlight several advantages of Biomni x TusoAI:
Searches over data and code simultaneously and autonomously
Iterates hundreds of times per dollar spent
Guaranteed to produce efficient-methods
Requires no environment configuration, data setup, or coding whatsoever
Methods can now be developed entirely in one sentence of natural language. This is the most accessible method development has ever been in biology.
Below, we show Biomni x TusoAI developing new powerful models for (1) enhancer-gene linking and (2) genetic perturbation prediction. This work will be published in the coming months, and the new TusoAI is available early and open-source at github.com/Alistair-Turcan/TusoAI . Access to Biomni x TusoAI is available today at biomni.phylo.bio under the skills tab.
Figure 1: Using Biomni to orchestrate TusoAI.
Applications of Biomni x TusoAI
1. State-of-the-art genetic perturbation prediction
Predicting the transcriptomic effects of CRISPR gene knockouts is a promising but difficult task with sizeable interest, including the annual Virtual Cell Challenge with hundreds of groups participating. This task is challenging because the setup is unclear, evaluations disagree, and the proper data to use is unknown.
We apply Biomni x TusoAI to 3 independent benchmarks to build a new perturbation prediction method from scratch. Biomni searches over gene embeddings, training and evaluation strategies, and the models to link all these together, improving performance on a small validation subset of the data. After 1 week of iterating autonomously and testing 500 candidates, Biomni x TusoAI discovers a new method integrating 9 sources of gene embeddings with an ensemble of simple regression and kNNs that outperforms every method, including massive deep learning models, on almost every metric and dataset.
While perhaps surprising, this simpler method makes perfect sense in the literature. Previous work had identified that richer gene embeddings lead to stronger performance, but had not carefully tuned each individual data source, but instead applied the same embedding procedure to dozens of data sources. Here, Biomni x TusoAI took this paradigm further, discovering customized gene embedding strategies that enabled simple models to learn perturbational effects. This new method is available publicly at github.com/phylobio/TusoPerturb .
Figure 2: Performance of a new genetic perturbation model built by Biomni x TusoAI.
2. Discovering novel features for enhancer-gene linking
The identification of enhancer-gene links has become a central task in human genetics, pursued by several large efforts including the ENCODE Consortium. State-of-the-art methods follow a familiar paradigm — gathering large amounts of data from DNA sequences, epigenomic tracks, Hi-C contacts, multiome data, and more, then manually building some set of features from this data that are useful for linking enhancers to genes.
To automate the development of enhancer-gene linking methods, we applied TusoAI to build and test thousands of features from a large corpus of data. Starting from a leading method pgBoost and 36 different genomic data sources, TusoAI iteratively builds and tests features from these data to improve upon pgBoost. After 10,000 iterations, TusoAI built a new model, integrating an additional 23 novel features from 14 different data sources. This new model substantially outperformed all methods on ground truth eQTL evaluations, including a 5-fold improvement at difficult distal (>100KB) links. Similar strong performance was observed for ground truth CRISPR and GWAS evaluations.
The two most important new features are heavily engineered versions of the common paradigm of using the SNP-TSS distance. Instead of naively inputting SNP-TSS distance, these two redesign this as strand and gene overlap-aware sigmoid-transformed covariates. Other features are completely novel to pgBoost, including binary intron flags, GC content surrounding the SNP, and engineered ABC scores.
Figure 3: Improving enhancer-gene linking with TusoAI. Average enrichment of ground truth eQTLs across various distance thresholds. pgBoost (TusoAI) denotes the TusoAI-improved version of pgBoost.
Try it today at biomni.phylo.bio by prompting with your favorite task. Biomni-Tuso is now available as an auto-research skill in Biomni Lab for enterprise users; for public platform, we will share 500 GPU hours for 100 Biomni lab users as we prepare to roll out this capability broadly. Apply here .
Collaborative agents designed for biologists to make more discoveries, faster.
© 2026 Phylo, Inc. All rights reserved.
© 2026 Phylo, Inc. All rights reserved.

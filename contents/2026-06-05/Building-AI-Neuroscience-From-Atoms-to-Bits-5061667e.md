---
source: "https://blog.amaranth.foundation/p/building-ai-neuroscience-from-atoms"
hn_url: "https://news.ycombinator.com/item?id=48407807"
title: "Building AI Neuroscience: From Atoms to Bits"
article_title: "Building AI Neuroscience: From Atoms to Bits"
author: "pminimax"
captured_at: "2026-06-05T04:33:42Z"
capture_tool: "hn-digest"
hn_id: 48407807
score: 2
comments: 0
posted_at: "2026-06-05T04:09:07Z"
tags:
  - hacker-news
  - translated
---

# Building AI Neuroscience: From Atoms to Bits

- HN: [48407807](https://news.ycombinator.com/item?id=48407807)
- Source: [blog.amaranth.foundation](https://blog.amaranth.foundation/p/building-ai-neuroscience-from-atoms)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T04:09:07Z

## Translation

タイトル: AI 神経科学の構築: 原子からビットまで
説明: 神経科学は遅れています。どうすれば速くできるでしょうか?

記事本文:
AI 神経科学の構築: 原子からビットまで
アマランス財団のブログ
AI 神経科学の構築: 原子からビットまで
神経科学は遅いです。どうすれば速くできるでしょうか?
アマランス財団とパトリック・ミノー 2026年6月4日 7 3 シェア 大学の小規模研究室によって推進される神経科学は、細心の注意を払って進歩しています。ひるむことのないポスドクや大学院生によって、10 年間にわたりプロジェクトが英雄的に遂行されるのを見るのは珍しいことではありません。この研究は、神経変性疾患の治療や人間の知性を理解するための種を提供する可能性があります。どうすれば神経科学を加速できるでしょうか?
AI 科学者エージェント (文献を読み、仮説を生成し、データを読み取り、分析コードを記述し、実験を設計できるシステム) を使用して、直接またはアトラスやデジタル ツインとしてコンパイルして脳と行動を研究できれば、神経科学を大幅に加速できる可能性があります。実際、ダリオ・アモデイ氏は『Machines of Loving Grace』の中で、高度なAIがデータセンターで天才の国として機能し、知能科学の構築とすべての精神神経疾患の治療に向けていかに急速な進歩を遂げることができるかを説明しました。これらは高い目標ではありますが、このエッセイではどうすればそれを達成できるかについては説明されていません。ここでは、AI 神経科学を構築する方法と、資金提供者が何を優先すべきかを描いています。
アマランス財団のブログをお読みいただきありがとうございます。新しい投稿を受け取るには無料で購読してください。
AI 神経科学者の概略図 AI 科学者の実用的な定義
AI 神経科学者は、AI 科学者エージェントの特別な例であり、すでに形になりつつあります。 Nature の 2026 年 5 月号では、これらのシステムのうち 3 つが取り上げられ、生物医学科学における経験的ソフトウェアの作成と仮説のテストへの使用を実証しました。
大まかに言えば、これらのシステムはエージェント機能を備えた LLM です。

コンテキスト、メモリ、スキルへのアクセスを管理するハーネス。これらは、Claude Code や OpenAI Codex などのコーディング エージェントと同様に設計されていますが、最終製品はソフトウェアや Web サイトではなく、分析から得られた洞察であることに注意してください。
いくつかのグループは、科学全般にわたる AI 科学者や、FutureHouse のような生物科学に特化したエージェントを構築しています。私たちはすでに、これらの薬剤と神経科学との関連性を垣間見ることができます。たとえば、Aygün et al. (2026) は、知能科学に向けた基本的な構成要素である、最先端の神経活動予測を実行できる自動調査エージェントを実証しています。 AI 科学者が反復的により優れたタンパク質と治療法を設計することは、AI がアルツハイマー病研究をどのように前進させるかについての OpenAI Foundation の変革理論の一部となっています。
現在、AI 科学者の自律性は限られていますが、コーディング エージェントで見られたのと同様に、LLM の機能が成長するにつれて、AI 科学者の能力がさらに高まることが期待されています。ただし、神経科学に高度に特化したスキルを構築するには、根本的なボトルネックが残ります。これには、基本モデルのスループットに匹敵する単一のラボの範囲を超えるデータとソフトウェア エンジニアリングが必要になることが予想されます。
AI 科学者と脳と行動を学ぶ
AI 科学者は科学分野全体で比較的汎用的ですが、調査対象は異なります。従来の神経科学者と AI 神経科学者の主題は脳と行動です。 AI エージェントが仮説を迅速かつ安価にテストできるコードや数学などの検証可能な領域とは異なり、脳と行動に関する実験の実行には費用がかかります。
進歩を遂げるためには、脳と行動の研究を可能な限り過去の研究から移行する必要があります。

アトムの世界からビットの世界へ: アトラスを収集し、デジタル ツインを構築し、実際の被験者に対する仮説に基づいた実験でループを閉じます。
AI神経科学における被験者の形状
MICrONS データセットのサブセット。静的データセットが AI 神経科学者にとって役立つためには、アトラスのレベルまで引き上げる必要があります。アトラスとは、元の実験設計者が予想していたよりも多くの質問に答えることができる、高カバレッジで高エントロピーの脳の地図です。 Natural Scenes Dataset 、Allen Brain Cell Atlas 、FlyWire は、それぞれ fMRI、トランスクリプトミクス、コネクトミクスにわたるこのアイデアの最近の例です。これらのデータセットは完全性を念頭に置いて収集されており、それぞれのドメイン (FlyWire の場合は完全なドメイン) の代表的なサンプルが含まれています。これらは FAIR 原則に従って配布されます。高度に注釈が付けられ、サンプル コードとプログラムによるアクセスを備えたオープン プラットフォーム上で配布されます。
重要なのは、アトラスは実験チームと計算チームの間の好循環の中で生成される必要があるということです。最良のアトラスは、データをどのように使用するかについて明確なアイデアを持つ計算チームと連携して作成されます。 BRAIN Initiative のアーキテクトの 1 人である Tom Kalil 氏は、研究者はすぐに役立つモデルをもたらす適切な目的関数をすぐに特定できない可能性があると述べています。 「静的データセットよりもさらに価値があるのは、データ生成機能です。」これが私たちがセレンディピティを設計する方法です。
適切なニューロテクノロジーがなければアトラスを構築することは不可能であることがよくあり、アトラスとそのアトラスを取得するテクノロジーを共同設計することは理にかなっています。これまで収集された最大の機能地図帳を構築しているスタンフォード大学のエニグマプロジェクトに権限が与えられた

次世代ニューロピクセルの作成により、スループットが 4 倍に向上します。総合的な AI 戦略では、アトラスの構築に既存のツールを使用するのではなく、新しいツールの作成を最適化する必要があります。
静的データセットだけでは、仮定の実験を実行することはできません。そのためには、さまざまな入力や条件がシステムにどのような影響を与えるかを予測できるように、データをデジタル ツインにコンパイルする必要があります。神経科学では、このデジタル ツインは通常、ニューラル データを模倣するように直接トレーニングされたニューラル ネットワーク、またはデータによって固定された生物物理シミュレーションのいずれかです。ここでの重要な性能指数は、予測の妥当性です。つまり、関心のある値 (薬物の効果、刺激に応答したニューロンの発火率など) に関する双子の予測と、理想的には分布外で測定された実際の生物における同じ現象との間の相関関係です。
安価な実験的プロキシを使用すると、無数の介入をふるいにかけることはできますが、その精度は予測の妥当性と同じくらいです。創薬では、たとえば、Scannell et al. (2022) は、予測の妥当性 (スピアマンの ρ、0 と 1 の間の値) を 0.1 増やすことは、桁違いに多くの化合物をスキャンすることに勝ると主張しています。 GSK は 1990 年代に 1,000 万の化合物を in vitro でスクリーニングしましたが、抗生物質は見つかりませんでした。 Domagk 氏は、1930 年代に初歩的なツールを使用してマウスで数百の色素をテストし、プロントシルを発見しました。マウスは単に、単離された細胞よりもヒトにおける抗生物質反応のより優れた代理人です。
予測妥当性の高いモデルを目指して努力する必要があります。これは、因果関係の操作を含む大量の高エントロピー データをモデルに当てはめることによって実現できます。視覚神経科学の場合、これは自然主義的な視覚にわたる多様な視覚入力に応じて大量のデータを収集することを意味します。麻薬ディスコブ用

ery、広範なクラスの薬物をテストします。コネクトミクスのために、超微細構造をより正確で有効なシミュレーションにコンパイルするためのキャリブレーション データを収集します (Kording et al. 2026)。
猫にとって最良の材料モデルは、別の、できれば同じ猫です
従来の神経科学者をクラウドで実行する精力的な AI 神経科学者に置き換えても、従来の主題に関する従来の実験が高速化されるわけではありません。実際、ボトルネックはデータ分析からデータ収集に移ります。実際のサブジェクトの実行で実際の高速化を確認するには、次の 2 つのパスを予測します。
AI 神経科学者は、実際の被験者に対して実行する、より優れた、より識別力の高い実験を考案しました。神経科学者は研究室実験の実施に関する暗黙的かつ手順的な知識を統合しているが、これらの知識は文献には反映されていないことが多いため、これをクリアするのは高いハードルである。
AI 神経科学者は実験をより速く実行します。ラボの自動化を通じて。私たちはこれがバイオヘビーの神経科学やクラウドソーシングプラットフォームを通じて実行できる認知科学においても短期的には実現可能であると考えています。しかし、システム神経科学の自動化には明らかにギャップがあり、最終的には軟体ロボット工学のブレークスルーが必要になる可能性があります。
したがって、今後数年間は、実際の被験者に対して仮説に基づいた実験を実行することがボトルネックのままになるでしょう。これらの貴重なリソースを、アトラスとデジタル ツインに基づいて AI 神経科学者によって作成された予測の検証に集中する必要があります。これにより、好循環でより良いモデルを作成できるようになります。
資金を提供するプロジェクトの適切な形
AI 神経科学が今後数年間でどのように進歩するかについてのこのモデルを踏まえると、資金提供者はバリュー チェーンにおけるレバレッジの高いポイントに注意を向けるべきです。
アトラスレベルまで上がったデータ。 AIエージェントが分析を自動化するにつれてアトラスの価値は高まる

エス。これは、特に系統発生的にヒトに近い種において、より多くのコネクトーム、トランスクリプトーム、タンパク質の注釈、機能記録が必要になることを意味します。これらのアトラスは、測定ドメイン内で完全で、エントロピーが高く、メタデータで高度に注釈が付けられ、他のアトラスと結合しやすくするためにマルチモーダルである必要があります。
アトラス構築のためのより優れたニューロテクノロジー。神経科学におけるツールの不足により、私たちは依然として制限を受けています。理想的には、脳内のすべてのニューロンのスパイク列を記録できるはずです。樹状突起全体のシナプスの重み。神経ペプチドと神経調節物質の正体と濃度。資金提供者は、アトラスの構築をより安価かつ迅速にするスケーラブルなテクノロジーと、ニューロテクノロジーとアトラスを共同設計するプロジェクトに焦点を当てる必要があります。
より良いデジタルツイン。該当する場合、アトラス データをデジタル ツインに抽出して、もしもの質問に答える必要があります。これにより、データへのアクセスが容易になります。 LLM が何兆ものトークンを圧縮するのと同じように、デジタル ツインはペタバイト規模のデータを理想的には単一の GPU で実行できるモデルに抽出できます。評価を安価にするということは、産業規模のトレーニングを実行することを意味し、プロジェクトはトレーニング レシピの改善を考慮して、それに応じてコンピューティングを割り当てる必要があります。
より良いベンチマーク。重要なことは、優れたプロジェクトでは、人間における予測の妥当性について KPI を定量化し、改良することに努めるべきであるということです。ファーストパーティのベンチマークに加えて、サードパーティの評価や公開競争も依然として深刻な資金不足のままです。 ImageNet コンテストが現代の深層学習を促進したように、優れたベンチマークは神経科学を促進する新しいゲームを生み出すことができます。
AGI 周辺の将来を計画するということは、ギザギザのフロンティアのさまざまなセクションが予測不可能な順序で発生するため、幅広い結果を織り込むことを意味します。私たちは

したがって、主に静的データセットを操作し、ループ内で人間の実験者に大きく依存する AI 科学者から、現在自動化されたデータ取得機能が存在する完全自律型研究まで、幅広い自律型クラスに沿って共存するプロジェクトが見られるでしょう。
何よりも、最高のプロジェクトにはブロックを解除するものについての明確な理論があります。つまり、ImageNet が深層学習のブロックを解除した方法や、FlyWire がフライ神経科学のブロックを解除した方法のように、質問に答えることで他の全員の制約が解除されるという具体的な理由があります。機能が月ごとに変化するため、リサーチ テイストと呼ばれるものに基づいたこの長期ビジョンが、優れたプロジェクトと革新的なプロジェクトの違いを意味する可能性があります。
AI 神経科学には大きな期待が寄せられていますが、現在のシステムに AI を散りばめることでシステムのボトルネックが解決されるという誤った考えに陥ってはなりません。データセンターの天才たちが集まる国が、1 世紀にわたる神経科学の進歩を 10 年に圧縮するという目標を達成するには、これは至難の業です。私たちは今すぐツール、データセット、プロジェクトを構築する必要があります。これは、AI を調整するためのインテリジェンスの深層科学を構築するという私たちの目標に特に当てはまります。調整にとって最も重要な時期は、今後数年間です。 AI 神経科学を構築するには - distilli

[切り捨てられた]

## Original Extract

Neuroscience is slow. How can we make it faster?

Building AI Neuroscience: From Atoms to Bits
Amaranth Foundation Blog
Subscribe Sign in Building AI Neuroscience: From Atoms to Bits
Neuroscience is slow. How can we make it faster?
Amaranth Foundation and Patrick Mineault Jun 04, 2026 7 3 Share Neuroscience, driven by small-scale labs in universities, proceeds meticulously; it is not uncommon to see projects carried heroically by undaunted postdocs and graduate students for a decade. This research could provide the seeds for treating neurodegenerative disorders or for understanding our intelligence. How can we accelerate neuroscience?
If we could use AI scientist agents — systems that can read literature, generate hypotheses, read data, write analysis code, and design experiments — to study Brains and Behavior — either directly or compiled as atlases and digital twins—, we could potentially vastly accelerate neuroscience. Indeed, in Machines of Loving Grace , Dario Amodei described how advanced AI, acting as a country of geniuses in a datacenter , could make rapid advances towards building a science of intelligence and curing all neuropsychiatric disease. While these are lofty goals, the essay doesn’t tell us how we might reach them. Here, I paint a picture of how to build AI neuroscience and what funders should prioritize.
Thanks for reading Amaranth Foundation Blog! Subscribe for free to receive new posts.
A schematic of the AI neuroscientist A working definition of AI scientists
AI neuroscientists are a special instance of AI scientist agents, which are already taking shape. The May 2026 issue of Nature featured three of these systems demonstrating their use for writing empirical software and testing hypotheses in the biomedical sciences.
In broad strokes, these systems are LLMs with agentic harnesses that manage context, memory, and access to skills. They are architected similarly to coding agents, like Claude Code or OpenAI Codex, with the caveat that the end product is not a piece of software or website, but rather insight derived from an analysis.
Several groups are building AI scientists across the sciences, and more specialized agents for the biosciences, such as the ones from FutureHouse. We can already see glimpses of the relevance of these agents to neuroscience. For example, Aygün et al. (2026) demonstrate an auto-research agent that can perform state-of-the-art neural activity prediction, a fundamental building block towards a science of intelligence. AI scientists iteratively designing better proteins and treatments is already part of the OpenAI Foundation’s theory of change for how AI can advance Alzheimer’s disease research.
Currently, AI scientists have limited autonomy, though we expect them to become more capable as LLM capabilities grow, similar to what we’ve seen in coding agents . A fundamental bottleneck will remain, however, in building skills that are highly specialized to neuroscience. We expect this will require data and software engineering that are beyond the scope of a single lab to match the throughput of the base model.
Studying Brains and Behavior with AI scientists
While AI scientists are relatively generic across scientific disciplines, they differ in the subjects that they interrogate. The subjects of conventional and AI neuroscientists are Brains and Behavior. Unlike verifiable domains like code and math, where AI agents can test hypotheses rapidly and cheaply, running experiments on brains and behavior is expensive .
To make progress, we have to move as much of our study of brains and behavior as possible from the world of atoms to the world of bits: collecting atlases, building digital twins, and closing the loop with hypothesis-driven experiments on real subjects.
The shape of the subject in AI neuroscience
A subset of the MICrONS dataset . For a static dataset to be useful for an AI neuroscientist, it has to rise to the level of an atlas : a high-coverage, high-entropy map of the brain that can answer many more questions than could have been anticipated by the original experiment designers. The Natural Scenes Dataset , the Allen Brain Cell Atlas , and FlyWire are recent examples of this idea across fMRI, transcriptomics, and connectomics, respectively. These datasets are collected with completeness in mind, containing a representative sample of their respective domain (or in FlyWire’s case, the complete domain). They are distributed according to FAIR principles: highly annotated, distributed on open platforms, with sample code and programmatic access.
Importantly, atlases must be generated in a virtuous cycle between experimental and computational teams: the best atlases are made in conjunction with computational teams that have clear ideas of how they will use the data. Tom Kalil, one of the architects of the BRAIN Initiative, says that researchers may not be able to immediately identify the right objective function that will result in an immediately useful model. “What is even more valuable than a static dataset is a data generation capability”. This is how we engineer serendipity.
It’s often the case that building atlases is simply impossible without the right neurotechnology, and it makes sense to co-design the atlas and the technology to acquire that atlas. The Enigma project at Stanford, which is building the largest functional atlas ever collected, was empowered by the creation of next-generation Neuropixels, increasing throughput by a factor 4. A holistic AI strategy should optimize the creation of new tools vs. using existing tools to build atlases.
Static datasets, by themselves, do not allow us to run what-if experiments. To do that, we have to c ompile the data into digital twins that allow us to predict how different inputs and conditions will nudge the system. In neuroscience, this digital twin is typically either a neural net trained directly to imitate neural data, or a biophysical simulation anchored by data. The key figure of merit here is predictive validity: the correlation between the twins’ predictions on a value of interest (e.g. the effect of a drug, the firing rate of a neuron in response to a stimulus) and the same phenomenon in the real organism, ideally measured out-of-distribution.
While cheap experimental proxies allow us to sift through countless interventions, they are only as good as their predictive validity . In drug discovery, for example, Scannell et al. (2022) argue that increasing predictive validity (Spearman’s ρ, a value between 0 and 1) by 0.1 beats scanning orders of magnitudes more compounds. GSK screened 10 million compounds in vitro in the 1990s and found zero antibiotics; Domagk, working with rudimentary tools in the 1930s, tested a few hundred dyes in mice and found Prontosil . Mice are simply a better proxy of antibiotic response in humans than isolated cells.
We should strive for models with high predictive validity. This can be achieved by fitting models with a large amount of high-entropy data that include causal manipulations. For visual neuroscience, that means collecting large amounts of data in response to diverse visual inputs that span naturalistic vision; for drug discovery, testing broad classes of drugs; for connectomics, collecting calibration data for compiling ultrastructure into ever more precise, valid simulations (Kording et al. 2026).
The best material model for a cat, is another, and preferably, the same cat
Replacing conventional neuroscientists with tireless AI neuroscientists running in the cloud won’t make conventional experiments on conventional subjects faster. Indeed, the bottleneck will move from data analysis to data collection. To see real speedups in running real subjects, we foresee two paths:
The AI neuroscientist comes up with better, more discriminative experiments to run on real subjects. That is a high bar to clear, as neuroscientists integrate implicit and procedural knowledge about running lab experiments, which are often not reflected in the literature.
The AI neuroscientist runs experiments faster, e.g. through lab automation . We see this as feasible in the short term in bio-heavy neuroscience, and also for cognitive science that can be run through crowdsourcing platforms. There is clearly a gap, however, in the automation of systems neuroscience, which may ultimately require breakthroughs in soft-body robotics.
For the next few years, then, running hypothesis-driven experiments on real subjects will remain a bottleneck. We should focus these precious resources on validating predictions made by an AI neuroscientist based on atlases and digital twins. This will enable the creation of better models in a virtuous loop.
The right shape of projects to fund
Given this model for how AI neuroscience can advance over the next few years, funders should direct their attention to high-leverage points in the value chain:
Data that rise to the level of atlases. Atlases will rise in value as AI agents automate analyses. That means we’ll need more connectomes, transcriptomes, protein annotations, and functional recordings, especially in species that are phylogenetically closer to humans. These atlases should be complete within their measurement domain, high entropy, highly annotated with metadata, and multimodal to make it easier to join with other atlases.
Better neurotechnology for building atlases. We are still limited by the paucity of tools in neuroscience. Ideally, we should be able to record spike trains of every neuron in the brain; the synaptic weights of an entire dendritic arbor; and neuropeptide and neuromodulator identity and concentration. Funders should focus on scalable technologies that make building atlases cheaper and faster, and on projects that co-design neurotechnologies and atlases.
Better digital twins . Where applicable, atlas data should be distilled into digital twins to answer what-if questions. This makes data more accessible; much like LLMs compress trillions of tokens, digital twins can distill petabytes of data into models that can run, ideally, on a single GPU. Making evaluation cheap means industrial-scale training runs, and projects should allocate compute accordingly, taking into account improvements in training recipes.
Better benchmarks. Critically, a good project should strive to quantify and refine KPIs for predictive validity in humans. In addition to first-party benchmarks, third-party evaluations and open competitions remain critically underfunded. A good benchmark can create new games to catalyze neuroscience, much as the ImageNet competition catalyzed modern deep learning.
Planning for the peri-AGI future means factoring in a broad range of outcomes, as different sections of the jagged frontier will fall in an unpredictable sequence. We’ll thus see projects that co-exist along a broad range of autonomy classes, from AI scientists primarily operating on static datasets and relying heavily on human experimenters in the loop, to fully autonomous research where automated data acquisition capabilities exist today.
Above all, the best projects have a clear theory of what they unblock: a specific reason that answering their question removes a constraint for everyone else, the way ImageNet unblocked deep learning or FlyWire unblocked fly neuroscience. With capabilities changing month-to-month, this long-term vision, guided by what some call research taste , can mean the difference between a good project and a transformative one.
AI neuroscience holds great promise, but we must not fall prey to the fallacy that sprinkling AI on top of the current system will address systemic bottlenecks. For a country of geniuses in a datacenter to have a shot on goal of compressing a century of neuroscience progress in a decade—a tall order!—we must build tools, datasets and projects today. This is especially true for our goal of building a deep science of intelligence to align AI : the most critical time period for alignment is in the next few years. To build AI neuroscience—distilli

[truncated]

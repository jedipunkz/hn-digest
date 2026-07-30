---
source: "https://www.nextplatform.com/hpc/2026/07/27/nvidia-accelerates-chip-engineering-with-ai-agents/5279125"
hn_url: "https://news.ycombinator.com/item?id=49107822"
title: "Nvidia Accelerates Chip Engineering with AI Agents"
article_title: "Nvidia Accelerates Chip Engineering With AI Agents"
author: "rbanffy"
captured_at: "2026-07-30T10:13:34Z"
capture_tool: "hn-digest"
hn_id: 49107822
score: 2
comments: 0
posted_at: "2026-07-30T09:47:51Z"
tags:
  - hacker-news
  - translated
---

# Nvidia Accelerates Chip Engineering with AI Agents

- HN: [49107822](https://news.ycombinator.com/item?id=49107822)
- Source: [www.nextplatform.com](https://www.nextplatform.com/hpc/2026/07/27/nvidia-accelerates-chip-engineering-with-ai-agents/5279125)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T09:47:51Z

## Translation

タイトル: Nvidia、AI エージェントでチップ エンジニアリングを加速
記事のタイトル: Nvidia、AI エージェントでチップ エンジニアリングを加速
説明: 10 年以上前、Nvidia は人工知能に全面的に焦点を当て、それ以来...

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
すべてのセクションの最新ニュース
Nvidia、AI エージェントでチップ エンジニアリングを加速
10 年以上前、Nvidia は人工知能に全面的に焦点を当て、それ以来、NEmotron 傘下の CUDA-X および AI モデル ファミリを通じた豊富なソフトウェア ライブラリだけでなく、AI システムを強化するために設計された GPU (そして現在は CPU) やその他のハードウェアを大量生産してきました。これにより、Nvidia はエンタープライズおよび HPC データセンターの開発を引き継ぎ、急速に拡大する AI 業界で指導的な役割を果たし、同社を並外れた富をもたらしました。
その富を推進している要因の一部は、AI 向けのシリコン エンジニアリングが世代ごとにますます複雑になり、より多くのチップだけでなくより多くの機能が求められており、業界が従来の方法では対応できない状況に追い込まれているという事実です。 Nvidia のバイスプレジデント兼計算エンジニアリング担当ゼネラルマネージャーである Tim Costa 氏は、単に需要に応じて拡張するだけではないという言葉で述べています。
コスタ氏はビデオ通話でジャーナリストらと話し、2030年までに業界は2兆個のチップを生産し、月に約4,100万枚のウエハーを処理すると予想されていると指摘した。さらに、個々のパッケージのトランジスタ数は 1 兆に近づき、コンピューティング システム全体のトランジスタ数は千兆に達しようとしています。同時に、エンジニアがシミュレーション、検証、実装のステップに何年も費やすため、チップを市場に出すには何年もかかることがあります。それは、AI が加速する時代における一生です。
従来のチップ設計を上回る複雑さ
「重要な点は単一の数字ではありません。それは規模、アーキテクチャ、パッケージング、システムの複雑さの相互作用です」と彼は言いました。 「伝統的なデザインプロセスは、

SS はその規模の挑戦に追いつくことができません。それに対応するために、AI とアクセラレーテッド コンピューティングは生産性ツールから基礎的なエンジニアリング インフラストラクチャへと移行しつつあります。」
Nvidia やその他のチップ エンジニアリング企業が AI に注目しているのはここです。
「この複雑さは、半導体イノベーションのエンドツーエンドの性質が組み合わさることによってさらに複雑になります」とコスタ氏は述べた。 「チップ アーキテクチャの決定は、原子スケールの製造、高度なパッケージング能力、熱、システム全体の動作に影響を与えます。 AI は、エンジニアがより多くの設計代替案を検討し、それらのやり取り全体でより適切な意思決定を下せるように支援します。アクセラレーションされたコンピューティングにより、これらの決定の裏にある忠実度の高いシミュレーション、検証、最適化が、繰り返し実行できるほど高速になります。」
同氏は、物理学や設計ルールを置き換えるのではなく、チップ開発プロセスに AI を導入することで、エンジニアはそのようなチェックをループに組み込んでより頻繁に使用できるようになり、「孤立したツールではなくエンジニアリング ループ全体を加速する機会が得られる」と付け加えた。
Nvidia はまた、エンジニアリング AI のトレーニングと展開のための更新された CUDA-X および PhysicsNeMo ライブラリを含むようにエージェント ツールキットを拡張しています。
さらに、このベンダーは、チップ エンジニアリング会社の Cadence および Synopsys と協力して、Arm ベースの「Vera」CV100 CPU (下図) を使用して、将来の世代の CPU および GPU を高速化しています。
Vera は、Nvidia によって設計された 88 個のカスタム「オリンパス」 CPU コアと 1.2 TB/秒の LPDDR5X メモリ サブシステムを搭載しています。 Nvidia によると、同社の第 2 世代スケーラブル コヒーレンシー ファブリック メッシュ インターコネクトは、コアごとの強力なパフォーマンス、高いメモリ帯域幅、エンジニアリング アプリケーション向けの低遅延を実現します。
「NVIDIA は、将来の CPU と GPU の作成に使用される EDA (電子設計自動化) ワークフロー全体に Vera を導入しています。

コスタ氏は、初期のエンジニアリングテストで、Vera が Synopsys の VCS と Cadence の Jasper プラットフォーム上で実行され、AMD の Epyc Torrent システムの 1.5 倍のパフォーマンスを提供することが示されたことを指摘し、「その実用的な価値は、シミュレーション検証の実行時間が短いことです。」エンジニアリング チームは反復をより速く行うことができます。私たちはケイデンスおよびシノプシスと協力して、Vera を Rosa の設計支援に活用することで、Vera 向けの主要な EVA アプリケーションを最適化しています。」
もちろん「Rosa」は、Rigel コアをベースに構築された Nvidia の次世代 CPU です。これは、ベンダーの今後のファインマン データセンター プラットフォームの一部として 2028 年に発売される予定です。
Cadence と Synopsys は両社とも、EDA やその他のチップ設計機能にエージェント AI の使用に熱心に取り組んでいます。ケイデンスは 2 月にシリコン設計および検証用のエージェント ツールである AI Super Agent を発表し、その後 2 か月で、チップおよびシステム設計での AI エージェントの使用に関する Nvidia、TSMC、Google との提携を発表しました。先月、自動チップ設計と検証のためのフロントエンド エージェント ワークフロー用の AuraStack AI Super Agent (下) がリリースされました。
ケイデンスによると、AI スーパー エージェントは、エンジニアの指示に従って、現在では 5 週間かかる作業を 1 日未満で実行できる数百のシミュレーションを同時に実装でき、40 倍高速なレジスタ転送レベル (RTL) 検証サイクルを提供します。
一方、Synopsys は Nvidia と提携しているだけでなく、今週、AMD、Microsoft、Intel と同様のエージェント AI ベースの提携を発表しました。今週カリフォルニアで開催されたデザイン オートメーション カンファレンスで、シノプシスは、同社のエージェント プラットフォームと AgentEngineer テクノロジに基づいた、Nvidia t を含む完全自律型設計検証ワークフロー (以下) をデモンストレーションしました。

新しい Nemotron 3 Ultra モデル、Agent Toolkit、OpenShell ランタイムなどのテクノロジー。ベンダーは、検証済みの RTL を他のプラットフォームより 50 倍速く提供できると述べています。
Nvidiaは3月に、NemotronモデルやOpenShellランタイムなどのコアコンポーネントを備えたエージェントツールキットを発表し、1か月前にはチップ設計用のエージェントコーディングのオープンモデルであるNemotron 3 Ultraを発表した。現在、ベンダーは再設計された PhysicsNeMo などのツールキットを拡張しており、エージェントが設計やシミュレーション作業のために AI 物理モデルをトレーニングおよび展開できるようにしています。
「PhysicsNeMmo は、専門家が AI 物理学向けの AI モデルを構築、トレーニング、展開するのを支援するフレームワークとして始まりました」とコスタ氏は言います。 「私たちは現在、その専門知識を、エージェント対応のスキルを備えたオープンで構成可能なライブラリのコレクションとして整理しています。これらのライブラリは、実用的なレイヤーとエンジニアリング ワークフローのニーズ、物理学を意識した操作、GPU ネイティブ メッシュ処理、分散トレーニング、データ キュレーション、モデル アクセラレーション、トレーニング レシピの構成、データ パイプラインの作成、データ セットのキュレーションなどのタスクの反復可能な命令をエンコードするスキル (どのツールを呼び出すか、どのような出力を生成するか、結果をどのように確認するかなど) をカバーしています。」
同氏はさらに、「専門家が手動で操作するフレームワークから、より大規模なエンジニアリングプロセス内でエージェントが呼び出し、構成、検証できるAI物理機能への移行である」と付け加えた。
CUDA-X には、EDA および科学シミュレーションの鍵となる大規模で複雑なスパース線形システム用の cuDSS のダイレクト スパース ソルバー用の加速ライブラリと、デバイスのニーズを満たす規模の量子化学シミュレーション用の cuEST が含まれています。 GPU での物理およびエンジニアリング シミュレーションのための大規模な疎線形システム用の cuISS が含まれるようになりました。
「多くのエンジニアリング用途

「ns は偏微分方程式から始まります。これらの方程式が離散化されるか、コンピューターで計算できるように離散化されると、ワークロードは非常に大きな疎線形システムになることがよくあります。」とコスタ氏は言いました。直接ソルバーは依然として不可欠ですが、エンジニアリングにおける多くのアプリケーションでは、メモリ効率とスケーラビリティのために反復手法が必要です。 cuISS は、GPU ネイティブの反復スパース ソルバーと事前調整ビルディング ブロックを CUDA-X に導入します。」
QuiX は Carina システムでフォトニック量子コンピューティングを前進させます
AMDがお金を追いかけるために使用しているラックスケールAIシステムのロードマップ
AI ホストとサンドボックスがインテルのデータセンター CPU Cookie を保存
Nvidia、AI エージェントでチップ エンジニアリングを加速
AMDがラックスケールAIシステムのロードマップで追い求めている資金
米国エネルギー省、AI をターゲットとしたジェネシス ミッションに向けて 50 億ドルのスターター ガンを発射
Google の Axion CPU がクラウドにもたらすメリット
Salience Labs はシリコン フォトニクス光スイッチで AI をスケールアップしたいと考えています
GenAI のハードウェア投資はモデルとプラットフォームの収益を大きく上回っています
Microsoft、大規模な AI CPU および GPU クラスターのために AMD を活用
Google、量子誤り訂正に AI 強化学習を使用
マーベル、Teralynx T100 で基数、低遅延、帯域幅を実現
AI チップが TSMC の収益の約 3 分の 1 を占める
SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
量子古典 HPC データセンターにおける HPE とデルの願望
AI 対応データが真の利点である理由
QuiX Quantum が HPC データセンター向けのフォトニック アーキテクチャを披露
AI コンピューティングが進むにつれてイーサネット ネットワーキングも進む
もちろんメタプラットフォームはクラウドになるだろう
3 人の HPC 達人が尋ねる: まだ GPU は必要ですか?
光学スケールアップ ファブリックはアーキテクチャではなく製造によって制限される
AMD、フラッシュ拡張メモリでサーバー DRAM を拡張
ブームの終わり/武

メモリ市場向けの st サイクル
中国の「LineShine」オール CPU、エクサフロップスクラスのスーパーコンピューターの詳細
HPE、セキュリティ、主権、マルチテナントのためのアップグレードされた HPC ハードウェア、ソフトウェアを提供
hpc
HPC スーパーコンピューティングでは AMD と Nvidia が互角の関係にあります
企業
HPE、Agentic AI の波に乗ってデータセンターに再び参入
店
Everpure の AI 戦略はほぼ純粋に Nvidia に基づいています
計算する
サーバーブームにより価格上昇とチップ不足が両立
ハイエンド コンピューティングを詳しくカバー
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
レジスター
開発クラス
ブロックとファイル
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を販売しないでください
同意のオプション
著作権。無断複写・転載を禁じます © 1998-2026。

## Original Extract

More than a decade ago, Nvidia turned its full focus on artificial intelligence, and since then has ...

Jump to main content
Search
More topics
All the latest news, from all sections
Nvidia Accelerates Chip Engineering With AI Agents
More than a decade ago, Nvidia turned its full focus on artificial intelligence, and since then has been churning out GPUs – and now CPUs – and other hardware designed to power AI systems as well as a wealth of software libraries through CUDA-X and a family of AI models with under the Nemotron umbrella. This has put Nvidia in a leadership role in a rapidly expanding AI industry that has taken over enterprise and HPC datacenter development and made the company extraordinarily wealthy.
Part of what drives that wealth is the fact that silicon engineering for AI is increasingly complex with every generation, with demand not only for more chips but also more capabilities, a situation that is pushing the industry beyond what traditional methods can handle. Tim Costa, vice president and general manager of computational engineering at Nvidia, put it in terms that go beyond simply scaling to meet demand.
Speaking with journalists in a video call, Costa noted that by 2030 the industry is expected to produce 2 trillion chips and process about 41 million wafers a month. In addition, individual packages are approaching a trillion transistors while entire computing systems are on their way to transistor counts that will reach into the quadrillions. At the same time, bringing a chip to market can take years as engineers spend years on the simulation, verification, and implementation steps. That’s a lifetime in the accelerated age of AI.
Complexity Outpaces Tradition Chip Design
“The key point is not any one number; it's the interaction of scale, architecture, packaging, and system complexity,” he said. “The traditional design process just can't keep pace with that scale of challenge. To meet it, AI and accelerated computing are moving from productivity tools into being foundational engineering infrastructure.”
This is where Nvidia and other chip engineering firms are turning to AI.
“This complexity is compounded by the coupled end-to-end nature of semiconductor innovation,” Costa said. “Decisions in chip architecture affect atomic-scale manufacturing, advanced packaging power, thermals, and the behavior of the complete system. AI helps engineers explore far more design alternatives and make better decisions across those interactions. Accelerated computing makes the high-fidelity simulation, validation, and optimization behind those decisions fast enough to repeat.”
Rather than replacing physics or design rules, bringing AI into the chip development process lets engineers put such checks into the loop and use them more often, he explained, adding “the opportunity is to accelerate the full engineering loop, not isolated tools.”
Nvidia also is expanding its Agent Toolkit to include updated CUDA-X and PhysicsNeMo libraries for training and deploying engineering AI.
In addition, the vendor is working with chip engineering firms Cadence and Synopsys in using its Arm-based “Vera” CV100 CPU (below) to accelerate future generations of its CPUs and GPUs.
Vera holds 88 custom “Olympus” CPU cores designed by Nvidia and a 1.2 TB/sec LPDDR5X memory subsystem. The company’s second-generation Scalable Coherency Fabric mesh interconnect for strong per-core performance, high memory bandwidth, and low latency for engineering applications, according to Nvidia.
“Nvidia is deploying Vera across the EDA [electronic design automation] workflows used to create our future CPUs and GPUs, including simulation, formal verification, and physical implementation,” Costa said, noting that early engineering testing shows Vera running on Synopsys’ VCS and Cadence’s Jasper platforms provide 1.5 times the performance of AMD’s Epyc Torrent systems. “Its practical value is shorter simulation verification runs. Engineering teams can iterate faster. We are working with Cadence and Synopsys to optimize leading EVA applications for Vera by putting Vera to work helping design Rosa.”
“ Rosa,” of course, is Nvidia’s next-generation CPU built on its Rigel core. It is due to launch in 2028 as part of the vendor’s upcoming Feynman datacenter platform.
Both Cadence and Synopsys are going hard into using agentic AI for their EDA and other chip design capabilities. Cadence in February announced AI Super Agent, an agentic tool for silicon design and verification, and in the following two months unveiled partnerships with Nvidia, TSMC, and Google around using AI agents for chip and system designs. Last month came the launch of AuraStack AI Super Agent (below) for front-end agentic workflow for automated chip design and verification.
The AI Super Agents, at the direction of engineers, can simultaneously implement hundreds of simulation to do in less than a day work that now requires five weeks, providing 40-times faster Register-Transfer Level (RTL) validation cycles , according to Cadence.
For its part, Synopsys not only is partnering with Nvidia, but this week announced similar agentic AI-based collaborations with AMD, Microsoft , and Intel . At the Design Automation Conference this week in California, Synopsys demonstrated its Fully Autonomous Design Verification Workflow (below) that is based on its agentic platform and AgentEngineer technology and includes Nvidia technologies, such as its new Nemotron 3 Ultra model, Agent Toolkit, and OpenShell runtime. The vendor said it can deliver validated RTL 50 times faster than other platforms.
Nvidia in March launched its Agent Toolkit , with core components including the Nemotron models and OpenShell runtime and a month ago unveiled Nemotron 3 Ultra, an open model for agentic coding for chip designs. Now the vendor is expanding the toolkit, including with a rearchitected PhysicsNeMo so agents can train and deploy AI physics models for design and simulation work.
“PhysicsNeMmo began as a framework that helped specialists build, train, and deploy AI models for AI physics,” Costa said. “We are now organizing that expertise as a collection of open, composable libraries with agent-ready skills. The libraries cover the practical layers and engineering workflow needs, physics-aware operations, GPU-native mesh processing, distributed training, and data curation, the skills that encode repeatable instructions for tasks such as model acceleration, configuring a training recipe, creating a data pipeline, or curating a data set, including which tools to call, what outputs to produce, and how the results should be checked.”
He added that “the shift is from a framework an expert operates manually to AI physics capabilities an agent can invoke, compose, and validate inside a larger engineering process.”
CUDA-X has included acceleration libraries for direct sparse solvers in cuDSS for large and complex sparse linear systems that are key to EDA and scientific simulation, as well as cuEST for quantum chemistry simulations to scales that meet the need of devices. It now includes cuISS for large sparse linear systems for physics and engineering simulations on GPUs.
“Many engineering applications begin with partial differential equations,” Costa said. “Once those equations are discretized – or made discrete-ready to compute on a computer – the workload often becomes a very large sparse linear system. Direct solvers remain essential ... but many applications in engineering require iterative methods for memory efficiency and scalability. cuISS brings GPU-native iterative sparse solvers and preconditioning building blocks into CUDA-X.”
With The Carina System, QuiX Pushes Photonic Quantum Computing Forward
The Rackscale AI System Roadmaps That AMD Is Using To Chase Money
AI Hosts And Sandboxes Save Intel’s Datacenter CPU Cookies
Nvidia Accelerates Chip Engineering With AI Agents
The Money AMD Is Chasing With Its Rackscale AI System Roadmaps
DoE Fires The $5 Billion Starter Gun For Its AI-Targeted Genesis Mission
How Google's Axion CPUs Benefit The Cloud
Salience Labs Wants To Scale Up AI With Silicon Photonics Optical Switch
GenAI Hardware Investments Are Way Ahead Of Model And Platform Revenues
Microsoft Taps AMD For At Scale AI CPU And GPU Clusters
Google Uses AI Reinforcement Learning For Quantum Error Correction
Marvell Brings Radix, Low Latency, And Bandwidth To Bear With Teralynx T100
AI Chips Drive Around A Third Of TSMC Revenues
SREs To AI Agents: Prove Yourself Before You Touch Production
The Aspirations Of HPE And Dell In The Quantum-Classical HPC Datacenter
Why AI-Ready Data Is The Real Advantage
QuiX Quantum Shows Off A Photonic Architecture For HPC Datacenters
As Goes AI Compute, So Goes Ethernet Networking
Of Course Meta Platforms Is Going To Be A Cloud
Three HPC Gurus Ask: Do We Still Need GPUs?
Optical Scale Up Fabrics Are Limited By Manufacturing, Not Architecture
AMD Stretches Server DRAM With Flash Extended Memory
The End Of Boom/Bust Cycles For The Memory Market
A Deep Dive On China’s “LineShine” All-CPU, Exaflops-Class Supercomputer
HPE Delivers Upgraded HPC Hardware, Software For Security, Sovereignty, And Multi-Tenancy
hpc
AMD And Nvidia Are Neck And Neck In HPC Supercomputing
enterprise
HPE Rides The Agentic AI Wave Back Into The Datacenter
store
Everpure’s AI Strategy Is Almost Purely Based On Nvidia
compute
The Server Boom Balances Price Increases Against Chip Shortages
In-depth coverage of high end computing
Contact us
Advertise with us
Who we are
Newsletter
The Register
DevClass
Blocks and Files
Cookies Policy
Privacy Policy
Ts & Cs
Do not sell my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

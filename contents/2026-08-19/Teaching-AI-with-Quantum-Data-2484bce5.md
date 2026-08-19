---
source: "https://www.quantinuum.com/blog/teaching-ai-with-quantum-data"
hn_url: "https://news.ycombinator.com/item?id=49362854"
title: "Teaching AI with Quantum Data"
article_title: "Teaching AI with Quantum Data"
image: "https://cdn.prod.website-files.com/669960f53cd73aedb80c8f0e/6a8485e52bc0835ce2cf74a0_GenQAI_MP_2.png"
author: "jonbaer"
captured_at: "2026-08-19T16:20:30Z"
capture_tool: "hn-digest"
hn_id: 49362854
score: 1
comments: 0
posted_at: "2026-08-19T15:25:18Z"
tags:
  - hacker-news
  - translated
---

# Teaching AI with Quantum Data

- HN: [49362854](https://news.ycombinator.com/item?id=49362854)
- Source: [www.quantinuum.com](https://www.quantinuum.com/blog/teaching-ai-with-quantum-data)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T15:25:18Z

## Translation

タイトル: 量子データを使用した AI の教育
説明: Quantinuum、NVIDIA、および Pfizer は、トランスベースの生成 AI と量子コンピューティングを組み合わせて、従来の最適化手法よりも効率的に高品質の量子化学回路を自動的に生成しました。

記事本文:
量子データを使って AI を教える
製品とソリューション 量子コンピュータ 量子システム 世界最高パフォーマンスの量子コンピュータを発見 システム モデル H1 リニア アーキテクチャを備えた当社の第 1 世代量子コンピュータ システム モデル H2 レーストラック アーキテクチャを備えた当社の第 2 世代量子コンピュータ Helios 当社の Hardware-as-a-Service は、オンプレミス、クラウド、またはその両方でのアクセスを提供します ハードウェア ロードマップ 商用量子コンピューティングの進歩を目的とした加速ロードマップを探索してください ソリューション Quantum Origin暗号化を強化し、データを保護するように設計された量子乱数生成器 (QRNG) InQuanto 複雑な分子および材料シミュレーションのための量子計算化学を加速する Nexus 最高パフォーマンスのフルスタックにアクセスして管理するためのオールインワン量子コンピューティング プラットフォーム クラウド アクセス 商用アプリケーション向けの世界をリードする量子ハードウェアを使用して探索する 開発者ツール TKET ゲートレベルの量子計算を開発および実行する Guppy 次世代の量子コンピュータをプログラムするPython Lambeq でホストされる当社の専用言語 量子論と自然言語の構造的関係の実験 研究 研究分野 当社の科学的専門知識の広大な状況を発見してください 技術ブログ 量子コンピューティングの世界について学ぶ 出版物 当社の科学チームによる研究およびジャーナル出版物を探索してください 会社 Quantinuum について 世界最大の統合量子コンピューティング企業について詳しく学ぶ 採用情報 量子コンピューティングの重心に参加しましょう お問い合わせ 当社のチームと連絡を取る 投資家 当社の投資家向け情報サイトをご覧ください
ニュース プレスリリース Quantinuum がどのように量子の世界の限界を押し広げているかをご覧ください ブログ q を超えて前進する

量子コンピューティングの状況 用語集 これらの用語と定義により、量子コンピューティングをより深く理解できます。 コミュニティ Q-Net フルスタック テクノロジの共有学習、サポート、コラボレーションを目的としたユーザー コミュニティに参加してください。
[切り捨てられた]
当社は NVIDIA およびファイザーと協力して、GenQAI の標準インスタンスで AI と量子コンピューティングを組み合わせ、計算化学のフロンティアを拡大しました。
プレゼンテーションを見る AI + 量子コンピューティング: Quantinuum、NVIDIA、および Pfizer は、トランスベースの生成 AI と量子コンピューティングを組み合わせて、従来の最適化手法よりも効率的に高品質の量子化学回路を自動的に生成しました。
実用的な製薬への影響: このアプローチは分子の基底状態を準備するために使用され、Quantinuum の Helios ハードウェアで検証され、より大規模な計算化学と創薬への道が実証されました。
長期的なビジョン: チームは、ますます複雑になる量子データから学習する量子基盤モデルを構築し、最終的には古典的なシミュレーションには大きすぎる分子の回路を AI で設計できるようにすることを目指しています。
量子コンピューティングは、分子シミュレーション、材料発見、医薬品開発など、コンピューティングでできることを拡大する未来を長い間約束してきました。しかし、その約束と実用性の間には、量子状態の準備という頑固なボトルネックが存在します。
量子コンピューターでアルゴリズムを実行するには、まず量子ビットを正しい開始状態にする必要があります。これは、Rube Goldberg マシンをセットアップするようなものだと考えてください。この場合を除いて、どの初期セットアップで希望の結果が得られるか正確にはわかりません。これが、量子状態の準備が非常に重要になる理由です。初期状態の選択によって、残りの計算の精度とコストが決まります。
NVIと提携しました

DIA とファイザーは、有意義な産業ワークフローの開発を目指して、この問題に取り組みます。その結果、ADAPT-GQE と呼ばれる新しい生成量子 AI フレームワークが誕生しました。これは、GenQAI の標準的なインスタンスであると考えられます。 ADAPT-GQE は、量子データを使用して変圧器モデルをトレーニングし、最終的には一種の「好循環」で量子化学回路をより速く、より良い結果で合成します。
最終的に、これは量子コンピューティングと AI の間の新しいインターフェイスを開発したことを意味します。量子回路生成を言語モデリング問題として扱うことにより、状態準備精度が同等または向上した高品質の基底状態準備回路を生成できるシステムが完成しました。
計算化学の目標は、高価で時間がかかり、場合によっては危険な「ウェットラボ」実験を行わずに化学特性を学ぶことです。
原理的には、物理​​実験の大部分をコンピューター シミュレーションに置き換えることができ、数十億ドルと何年もの時間を節約できます。
実際のところ、計算化学は非常に難しいものです。コンピューター内の化学物質を正確にシミュレートするには、それを一から構築する必要があります。原子の集合から始めます (イミプラミンの場合、19 個の炭素原子、24 個の水素原子、および 2 個の窒素原子があります)。次に、自然の「レゴ」のように、それらの原子を組み立てて分子を作ります。結合の長さ、強さ、角度、相互作用などを設定します。
これは単純ではありません。単一の分子はさまざまな形で存在する可能性があります。これらの異なる形状を「立体構造」と呼びます。
次に、実際に化学的性質を推定したり、化学反応経路を探索したりするには、原子レベルで進行する詳細な物理現象を再現する必要があります。

次に、各軌道がどのように占有されているか、電子同士または原子核がどのように相互作用しているか、熱や触媒の追加が物事にどのような影響を与えるかを解明します...それはすぐに複雑になります。
このような状況にもかかわらず、計算化学は医薬品開発における原動力です。現在、製薬会社は、時間と費用がかかる実験室での実験を避け、コンピュータ上で可能な限りシミュレーションを行うことで、お金と時間を節約しています。ただし、開発に 50 年以上かかっているにもかかわらず、既存の古典的な手法には現実的な限界があります。
ここで量子コンピューティングが登場します。この新しい計算パラダイムは、多くの「難しい部分」（重ね合わせやもつれなど）がネイティブにエンコードされているため、これらの制限を回避できます。量子コンピューティングを正しく使用すれば、古い障壁を打ち破り、意味のあるインパクトのある発見に貢献しながら、世界中の製薬会社の利益をさらに向上させることが期待できます。
量子計算化学は、短期的な量子優位性の最も有力な候補の 1 つですが、現在のハードウェアはまだ開発の初期段階にあります。量子ビットとエラー率が限られているため、アルゴリズム設計者はすべてのゲートをカウントし、回路を浅く保ち、ある程度のノイズを許容できる必要があります。
ここで、生成 AI が登場します。
化学回路を手作業で設計し、どれだけうまく動作するかを苦労して実験する代わりに、別のアイデアがあります。量子計算化学が直面する問題を解決するために AI を訓練したらどうなるでしょうか。
このアプローチを使用すると、時間とリソースを節約できるだけでなく、しかし、実際の結果を実現するまでのスケジュールを短縮することはできます。状態の準備やその他の回路が改善されれば、かつては遠い未来だと考えられていたアプリケーションも実現します。

視界に入る。
これに対する最初の試みは ADAPT-GQE と呼ばれます。 ADAPT-GQE の背後にある中心的なアイデアは一見単純です。苦労してゼロから適切な量子回路を探す代わりに、変換モデルをトレーニングして回路を直接生成します。
重要なのは、このフレームワークはモデルに依存しないことです。これは、Nemotron (事前トレーニングされた LLM) と Gemma (ゼロからトレーニングされた) という相補的なトランスフォーマー アーキテクチャにデプロイすることで示されました。
ここでの最初の目標は、イミプラミン分子の「基底状態」を見つけることです (これは、内部に蓄積されたエネルギーが最小の電子状態です)。これを行うには、上で説明したように、適切な「状態準備回路」を見つける必要があります。
これまで、量子コンピューターで基底状態を見つけるための主要な方法は、量子と古典のハイブリッドアプローチである「変分量子固有ソルバー (VQE)」でした。 VQE プロセスは、分子の特定の構造を「推測」する回路から始まります。量子コンピューターは回路を実行して、分子の関連エネルギーを測定します。この結果は古典的なオプティマイザにフィードバックされ、回路パラメータが微調整され、分子エネルギーがより低い回路が得られることが期待されます。このループは、最小エネルギーが見つかるまで繰り返されます。
残念ながら、VQE には、広範な使用を不可能にする重大な制限がいくつかあります。最近提案された ADAPT-VQE は、「単純な」VQE の問題のいくつかに対処することを目的とした重要な前進でした。 ADAPT-VQE では、最初の回路を推測することから始めるのではなく、プロセスはプールから演算子を選択し (通常は勾配情報を使用して)、最適化することで段階的に回路を構築します。このアプローチはより効果的ですが、残念なことに、依然として急速に大きくなりすぎます。
ここで合同チームが登場した。
最高の組み合わせ

チームの新しいフレームワークである ADAPT-GQE は、AI と ADAPT-VQE を組み合わせてまったく新しいものを作成します。これは、これまでのところ、自動量子回路合成に向けたスケーラブルでハードウェア検証済みの経路です。
まず、トランスフォーマー (この場合、Nemotron と Gemma) は、ADAPT-VQE データの教師あり微調整によってトレーニングされます。このようにして、古いメソッドは捨てられるのではなく、高品質のデータを生成する「神託」として扱われます。
次に、変圧器が最初にトレーニングされると、回路全体の分布が定義され、各回路は基底状態に対応する一定の確率を持ちます。この分布は、強化学習などの微調整ループで使用できます。強化学習では、フレームワークがその分布から回路を取得して実行し、エネルギーを測定します。結果はトランスフォーマーにフィードバックされ、トランスフォーマーがその分布を調整します。時間の経過とともに、モデルは、より正確な基底状態を準備する回路の優先順位を学習します。
重要なのは、強化学習により、システムが元のトレーニング データを単に模倣するのではなく、それを超えることができるということです。モデルは、ADAPT-VQE の圧縮ルックアップ テーブルとして機能しなくなりました。教師アルゴリズム自体を上回る可能性のある新しい回路構成の探索を開始します。これは、プロジェクトにおける最も重要な概念の変更の 1 つです。
この場合、Quantinuum の Helios ですべての初期回路を実行する代わりに、強化学習回路は NVIDIA アクセラレーテッド コンピューティングと CUDA-Q プラットフォームを使用して実行され、量子プロセッサをシミュレートしました。
最後に、強化学習によってトランスが最適化されると、Quantinuum 上の InQuanto と Nexus を使用して回路を実行することにより、結果として得られる最良の回路の精度と実現可能性が検証されます。

の最新ハードウェア、Helios 。 InQuanto v5.2 を使用すると、ユーザーは Nexus を介して Helios 量子コンピュータと Selene 量子エミュレータの両方に直接接続できるようになります。
InQuanto と Nexus のこの強力な組み合わせにより、AI によって生成されたこれまでで最大規模の量子化学回路の 1 つを量子コンピューター上で実行することが可能になりました。量子コンピューティングの可能性を医薬品開発の実用的なツールに変えるのに役立ちます。
さらに将来に目を向けると、研究者らは単一の分子ベンチマークよりもはるかに大きな何かを構想しています。
より大きく複雑な分子の場合、ADAPT-VQE は最初のトレーニングの「オラクル」としては機能しません。さらに、強化学習で使用される分子エネルギー計算は、量子コンピューターをシミュレートする古典的なシステムにとって大規模になりすぎるため、量子プロセッサが不可欠になります。
幸いなことに、これは問題ではありません。 ADAPT-GQE フレームワークの最終目標は、変圧器の「カリキュラム」を開発することです。これは、新しい分子ごとに再トレーニングするのではなく、すでに学んだことを保持し、そこから知識を拡張することを意味します。
最初に、より小さく、完全にシミュレーションできる分子について学習させることで、kno を使用して二重チェックできる適切なデータに基づいて学習することが保証されます。

[切り捨てられた]

## Original Extract

Quantinuum, NVIDIA, and Pfizer have combined transformer-based generative AI with quantum computing to automatically generate high-quality quantum chemistry circuits more efficiently than traditional optimization methods.

Teaching AI with Quantum Data
Products & Solutions Quantum Computers Quantinuum Systems Discover the world’s highest performing quantum computers System Model H1 Our first-generation quantum computer with a linear architecture System Model H2 Our second-generation quantum computer with a racetrack architecture Helios Our Hardware-as-a-Service provides access on-premise, in the cloud, or both Hardware Roadmap Explore our accelerated roadmap aimed at advancing commercial quantum computing SOlutions Quantum Origin Our advanced quantum random number generator (QRNG), designed to strengthen encryption and protect data InQuanto Accelerate quantum computational chemistry for complex molecular & materials simulations Nexus All-in-one quantum computing platform for accessing and managing the highest performance full-stack Cloud Access Explore using our world-leading quantum hardware for commercial applications Developer Tools TKET Develop and execute gate-level quantum computation Guppy Program the next generation of quantum computers with our purpose-built language hosted in Python Lambeq Experiment with the structural relationships between quantum theory and natural languages Research Research Areas Discover the vast landscape of our scientific expertise Technical Blog Get educated on the world of quantum computing Publications Explore research and journal publications from our scientific teams Company About Quantinuum Learn more about the world’s largest integrated quantum computing company Careers Join us in the center of gravity for quantum computing Contact Us Get in touch with our team Investors Visit our Investor Relations site
News Press Releases See how Quantinuum is pushing the boundaries in the world of quantum Blog Taking steps forward across the quantum computing landscape Glossary Better understand quantum computing with these terms and definitions Community Q-Net Join our user community aimed at shared learning, support and collaboration for our full-stack te
[truncated]
We’ve teamed up with NVIDIA and Pfizer to expand the frontier of computational chemistry, combining AI and quantum computing in a canonical instance of GenQAI
View the presentation AI + quantum computing: Quantinuum, NVIDIA, and Pfizer have combined transformer-based generative AI with quantum computing to automatically generate high-quality quantum chemistry circuits more efficiently than traditional optimization methods.
Practical pharma impact: The approach was used to prepare molecular ground states and validated on Quantinuum’s Helios hardware, demonstrating a path toward larger-scale computational chemistry and drug discovery.
Long-term vision: The team aims to build quantum foundation models that learn from increasingly complex quantum data, eventually enabling AI to design circuits for molecules too large for classical simulation.
Quantum computing has long promised a future that expands what we can do with compute — for example, in molecular simulation, materials discovery, or pharmaceuticals development. But between that promise and practical utility sits a stubborn bottleneck: quantum state preparation.
To run any algorithm on a quantum computer, you must first put the qubits in the right starting state. Think of it like setting up a Rube Goldberg machine- except in this case, you’re not sure exactly which initial setup will give you the results you want. This is what makes quantum state preparation so important: your choice of initial state dictates the accuracy and cost of the rest of the calculation.
We teamed up with NVIDIA and Pfizer to tackle this problem, with an eye towards developing meaningful industrial workflows. The result is a new generative quantum AI framework, called ADAPT-GQE, which we consider to be a canonical instance of GenQAI . ADAPT-GQE uses quantum data to train transformer models that ultimately synthesize quantum chemistry circuits faster, with better outcomes, in a sort of ‘virtuous cycle’.
Ultimately, this means we have developed a new interface between quantum computing and AI. By treating quantum circuit generation as a language modelling problem, we now have a system that can generate high-quality ground-state preparation circuits - with comparable or improved state preparation accuracy.
The goal of computational chemistry is to learn about chemical properties without performing expensive, time-consuming, and sometimes dangerous “wet-lab” experiments.
In principle, you can replace the majority of your physical experiments with computer simulations, saving billions of dollars and years of time.
In reality, computational chemistry is very tricky. To accurately simulate a chemical inside of a computer, you have to build it from the ground up. You start with a collection of atoms (in the case of imipramine, you have 19 Carbon atoms, 24 Hydrogen atoms, and 2 Nitrogen atoms). Then, like Nature’s ‘lego’, you assemble those atoms into a molecule: you set bond lengths, strengths, angles, interactions, and so on.
This is not straightforward: a single molecule can exist in many forms; with different angles, rotations, etc. We will call these different forms ‘conformations’.
Then, to actually estimate chemical properties, or to explore chemical reaction pathways, you have to reproduce the detailed physics that goes on at the atomic level: take your chosen conformation then figure out how each orbital is occupied, how the electrons are interacting with each other or the atomic nuclei, how is the addition of heat or a catalyst going to affect things.... it gets complicated, quickly.
Despite all this, computational chemistry is a powerhouse in pharmaceutical development. Right now, pharmaceutical companies save money and time by simulating as much as they can on computers, avoiding time consuming and expensive laboratory experiments. However, even with ~50 years of development, the existing classical methods have very real limitations.
This is where quantum computing comes in: this new computational paradigm can elide those limitations because it has many of the “hard parts” (like superposition or entanglement) natively encoded. Used correctly, quantum computing promises to break old barriers, further improving margins for pharma companies across the globe while contributing to meaningful, impactful, discoveries.
While quantum computational chemistry is one of the strongest candidates for near-term quantum advantage, current hardware is still in the earlier stages of development. With limited qubits and error rates, algorithm designers need to make every gate count, keep circuits shallow, and be able to tolerate some level of noise.
This is where generative AI enters the picture.
Instead of hand-designing chemistry circuits and laboriously experimenting to see how well they run, there is another idea: what if we trained an AI to solve the problems that quantum computational chemistry faces?
Using this approach, not only can we save time and resources; but we can shorten the timeline to realize practical results. With better state prep and other circuits, applications that were once considered far in the future come into view.
Our first attempt at this is called ADAPT-GQE. The central idea behind ADAPT-GQE is deceptively simple: instead of laboriously searching for good quantum circuits from scratch, train a transformer model to generate them directly.
Importantly, the framework is model-agnostic, which we showed by deploying it on complementary transformer architectures - Nemotron (a pretrained LLM) and Gemma (trained from scratch).
The initial goal here is to find the ‘ground state’ of the molecule imipramine (this is the electronic state with the smallest amount of energy stored inside it). To do this, you have to find the right ‘state preparation circuit’, as described above.
Until now, a leading method for finding the ground state with quantum computers was the ‘Variational Quantum Eigensolver (VQE)’, a hybrid quantum-classical approach. The VQE process starts with a ‘guess’ circuit for a particular conformation of the molecule. The quantum computer runs the circuit to measure the associated energy of the molecule. This result is fed back into a classical optimizer that then tweaks the circuit parameters, hopefully resulting in one with a lower molecular energy. This loop repeats until a minimum energy is found.
Unfortunately, VQE has a few severe limitations that make it infeasible for widespread use. The recently proposed ADAPT-VQE was a crucial step forward meant to address some of the issues with “plain” VQE. In ADAPT-VQE, instead of starting with a guess for the initial circuit, the process builds a circuit in steps by selecting operators from a pool (typically using gradient information) and optimizing. This approach can be more effective, but unfortunately still grows too large too quickly.
This is where the joint team jumped in.
Combining the best of all worlds, the team’s new framework, ADAPT-GQE, combines AI with the ADAPT-VQE to create something entirely new – and something that, so far, is a scalable, hardware-validated pathway toward automated quantum circuit synthesis.
First, transformers (in this case, Nemotron and Gemma) are trained via supervised fine-tuning on ADAPT-VQE data. In this way, the old method isn’t thrown away but is instead treated as a high-quality data-producing “oracle”.
Then, once the transformer has been initially trained, it defines a distribution over circuits, each one with some probability of corresponding to the ground state. This distribution can be used in a fine-tuning loop, for example, reinforcement learning. In reinforcement learning, the framework takes a circuit from that distribution, runs it, and measures the energy. It feeds the results back into the transformer, which adjusts its distribution. Over time, the model learns to prioritize circuits that prepare increasingly accurate ground states.
Crucially, reinforcement learning allows the system to surpass its original training data instead of merely imitating it. The model is no longer acting as a compressed lookup table for ADAPT-VQE. It begins exploring novel circuit configurations that may outperform the teacher algorithm itself. This is one of the most important conceptual shifts in the project.
In this case, instead of running all the initial circuits on Quantinuum’s Helios, the reinforcement learning circuits were run using NVIDIA accelerated computing and the CUDA-Q platform, simulating a quantum processor.
Finally, once the transformers are optimized via reinforcement learning, the best resulting circuits are validated for accuracy and feasibility, by running them using InQuanto and Nexus on Quantinuum’s newest hardware, Helios . With InQuanto v5.2, users can now interface directly with both the Helios quantum computer and the Selene quantum emulator through Nexus.
This powerful combination of InQuanto and Nexus enabled the execution one of the largest AI-generated quantum chemistry circuits to date on a quantum computer; helping to turn the promise of quantum computing into a practical tool for pharmaceutical development.
Looking farther in the future, the researchers envision something much larger than a single molecular benchmark.
For bigger and more complex molecules, ADAPT-VQE won’t work in the first place as the initial training “oracle”. In addition, the molecular energy calculations used in the reinforcement learning grow too large for classical systems simulating quantum computers, so the quantum processor becomes essential.
Luckily, this is not a problem. The ultimate goal of the ADAPT-GQE framework is to develop a “curriculum” for the transformers. This means instead of re-training them for every new molecule, you instead keep what you already learned, and expand your knowledge from there.
By initially teaching it on molecules that are smaller, and that can be fully simulated, you ensure it learns on good data that can be double checked using kno

[truncated]

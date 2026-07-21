---
source: "https://dreossi.github.io/blog/physical-ai-taxonomy/"
hn_url: "https://news.ycombinator.com/item?id=48992875"
title: "Control-Theory Taxonomy of Physical AI"
article_title: "A Control-Theory Taxonomy of Physical AI — Tommaso Dreossi"
author: "maunic"
captured_at: "2026-07-21T14:56:31Z"
capture_tool: "hn-digest"
hn_id: 48992875
score: 1
comments: 0
posted_at: "2026-07-21T14:32:09Z"
tags:
  - hacker-news
  - translated
---

# Control-Theory Taxonomy of Physical AI

- HN: [48992875](https://news.ycombinator.com/item?id=48992875)
- Source: [dreossi.github.io](https://dreossi.github.io/blog/physical-ai-taxonomy/)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T14:32:09Z

## Translation

タイトル: 物理 AI の制御理論分類法
記事のタイトル: 物理 AI の制御理論分類 — Tommaso Dreossi
説明: 宣伝された物理 AI 空間を制御ループにマッピングします。

記事本文:
TD
LinkedIn Scholar メールを書く 物理 AI の制御理論分類
SFのミッション・ストリートを歩いていると、ミケランジェロの「アダムの創造」という巨大な看板が目に入るが、それはアダムの手が神の手に伸びるのではなく、「物理的AI」という文字に手を伸ばす義肢ロボットだった。これで正式に発表されました。物理 AI の宣伝が始まりました。
図 1 — 物理 AI の作成: 誇大宣伝が始まります。 1
「フィジカルAI」という言葉が登場したとき、私はそれが物理学用AIを指すものだと思っていました。しかしそれ以来、ハードウェア、材料、ロボット工学、バイオテクノロジー、ワールド モデルなど、物理世界に関わるほとんどすべてのものが含まれるようになりました。現在ではそのすべてが物理 AI です。この用語がどれほど過剰になっているかを考えると、これらすべてが何がどのように結びついているのかを尋ねる価値があります。
この記事では、物理 AI メソッドを定義する特徴を探り、それらを典型的な物理フレームワークである制御理論に基づいて整理します。 2
物理的な AI と非物理的な AI はどう違うのでしょうか?
まず、物理的 AI と非物理的 AI の違いを定義することから始めましょう。
物理的な制約。最初の、そしておそらく最も受け入れられている特徴は、物理 AI は物理の支配方程式を尊重しなければならないということです。シミュレーションまたは観察されたデータに基づいてトレーニングされたモデルは、物理学を遵守しながら現実に移行する必要があります。物理 AI には幻覚が入り込む余地はありません。
データ不足とシミュレーションと実際のギャップ。物理的な AI データは不足していることがよくあります。ほぼ無制限のテキストに依存できる LLM とは異なり、物理的な実験は時間がかかり、費用がかかります。 PCB の製造と測定には数週間かかる場合があります。細胞の増殖と配列決定には数か月かかる場合があります。シミュレーションはデータ収集を加速できますが、現実との整合性という問題が伴います。シミュレーターのキャリブレーションは困難な場合があり、場合によっては不可能です。以来

物理データは希少で高価であるため、物理 AI モデルは、現実世界に転送できる十分な柔軟性を維持しながら、限られたデータ体制で学習する必要があります。
センシング、具体化、そしてリアルタイムの閉ループ。生き物はさまざまな感覚を通じて現実を経験するため、物理 AI システムはさまざまなセンサーやモダリティ (視覚、ライダー、力、温度など) からのデータを処理することが期待されます。物理 AI は現実世界で継続的に動作し、閉ループ システムとして動作することが期待されます。
この閉ループは、まさに制御理論が把握するために構築されたものであり、物理 AI 自体の自然な分類法を与えてくれます。
制御ループからの分類法
制御理論は、物理システムに明確なフレームワークを提供します。つまり、システムはその状態を感知し、アクションを決定し、コスト関数を最適化するために動作します。
w(t) ⊕ e(t) コントローラ/アクチュエータ制御 最適化 u(t) ⊕ d(t) u s (t) システムモデル y(t) センサ知覚 y m (t) 図2｜物理AIの閉ループ制御へのマッピング
以下で構成される閉ループ制御システムの図
設定点 w ( t )、
測定出力 y m ( t )、
測定誤差 e ( t )、
コントローラー出力 u ( t )、
システム入力 us ( t )、
外乱 d ( t )、
およびシステム出力 y ( t )。
色付きのブロックは、ループ内の各物理 AI ブランチが動作する場所を示します。
これは、同じループにマッピングされる 4 つのカテゴリにグループ化できる物理 AI メソッドに自然に適合します。
知覚は、センサー データ (カメラ、ライダー、力、化学物質など) を取り込み、それらを状態表現に変換することで、生の観察から物理状態を推定します。これは、コンピュータ ビジョン、深度推定、状態推定の領域です。
モデルは、システムと世界がどのように進化するかを特徴づけます。物理的なダイナミクスを学習し、状態とアクションを入力として受け取り、予測します。

未来の状態。ニューラル オペレーター、PINN、サロゲート モデル、デジタル ツインはすべてこのグループに分類されます。
制御は、システムが次に何をすべきかを指定します。現在の状態を考慮して、リアルタイムでアクションを決定し、実行します。ここには、ポリシー学習、強化学習、最適制御が存在します。
最適化は、どの構成が目標を達成するかを決定します。ターゲットから開始して、目的の特性を持つソリューションを検索または設計します。これは、ベイジアン最適化、インバース デザイン、およびジェネレーティブ デザインの領域です。
タブ。 1 — 知覚 – 制御 – モデル – 最適化の分類法にマッピングされた物理 AI アーキテクチャ。この表は包括的なものではなく例示的なものです。エントリは各組み合わせの代表的な例であり、完全なカタログではありません。ブロックの上にカーソルを置くと、独創的な論文が表示されます。
この分類法の構成要素は、単一のカテゴリ内に収まるメソッドです。たとえば、知覚における状態推定および視覚モデル (ViT、SAM、DINO、YOLO)。 Control でのポリシー学習方法 (拡散ポリシー、ACT、BESO)。モデル内のサロゲート モデル (Neural ODE、GNS、GraphCast)。最適化のベイジアン最適化 (GP、TuRBO、SAASBO)。
2 つのブランチにまたがるより柔軟なアプローチは、モデルと最適化の境界に存在するアーキテクチャ (PINN、ニューラル オペレーター、ジェネレーティブ デザイン モデルなど) であり、それぞれがサロゲートとして前方に実行したり、逆ソルバーとして後方に実行したりできます。または、知覚と制御を単一のモデルにブレンドする VLA (例: GR00T、RT-2、π₀)。
最後に、ワールド モデル (Cosmos、DreamerV3、JEPA など) は 3 つのカテゴリにまたがっており、潜在空間内で知覚、モデル、および制御を完全にブレンドしています。彼らは、現実世界と対話することなく、観察をエンコードし、将来の状態を予測し、想像上の解決策を計画します。
この分類法

メソッドが同じ場所にほとんど出現しない混雑した物理 AI 空間に秩序をもたらします。
会話はすべて同じ傘の下に置かれます。
この表は、制御理論と物理 AI の根本的な違いが、AI モデルによって機能カテゴリ間の境界を越えることができることであることも強調しています。たとえば、従来のロボット スタックは個別のビジョン、状態推定、ポリシー モジュールを維持しますが、VLA はそれらを 1 つのモデルにまとめます。同様に、PINN は、モデルがクエリされる方向に応じて、動的システムのモデル化と逆問題の解決の両方が可能です。
4 番目の列に手を伸ばす
提案されたフレームワークからパターンが現れます。古典的なテクニックは 1 つの列にまとめられています。より最近の手法は 2 つの列を占めています。PINN はサロゲートとして前方に実行され、逆ソルバーとして後方に実行されます。VLA は知覚と制御をマージします。世界モデルは 3 つの列をカバーします。そこで推測される質問は、次は 4 列モデルなのかということです。認識、モデル化、制御、最適化を行う単一の基盤モデル。自動運転ラボ 3 と自動リサーチ 4、5 は、今日私たちが利用できる最も近い例の一部です。設計を提案し、実験を実行し、結果を観察し、次に何を試すかを決定するシステムです。
しかし、私の知る限りでは、どちらも完全には実現していません。自動運転ラボはまだ単一のモデルではなく、特殊なモデルをつなぎ合わせた組み合わせで稼働しています。そして、自動研究はまだ物理世界と直接接触していません。
そしてこれがフィジカル AI の立場です。ロボットハンドは伸びていますが、指はまだつながっていません。物理的な天国はまだ数センチメートル先にあります。
ロボット・アダムの創造、サイエンスフォトライブラリー。
サイエンスフォト.com/media/92834 。
K.J.オーストロームとR.M.マレー、フィードバック システム: Sc 向けの概要

ientists and Engineers、プリンストン大学出版局、第 2 版、2021 年。
N.J. Szymanski 他、新規材料の加速合成のための自律実験室、Nature 624、86–91、2023。
Nature.com/articles/s41586-023-06734-w 。
A. Novikov 他、AlphaEvolve: 科学的およびアルゴリズム発見のためのコーディング エージェント、Google DeepMind、2025 年。
ディープマインド.google 。
A. カルパシー、オートリサーチ、2026 年。
github.com/karpathy/autoresearch 。

## Original Extract

Mapping the hyped Physical AI space onto the control loop.

TD
Writing LinkedIn Scholar Email A Control-Theory Taxonomy of Physical AI
As I walk down Mission St. in SF, I see a giant billboard: Michelangelo’s Creation of Adam, except instead of Adam’s hand reaching for God’s, it’s a robotic prosthetic reaching for the words “Physical AI.” Now it’s official: the Physical AI hype is on.
Fig. 1 — The Creation of Physical AI: the hype is on. 1
When the term “Physical AI” appeared, I assumed it referred to AI for physics. But since then, it has come to include pretty much anything that touches the physical world: hardware, materials, robotics, biotech, world models, etc. All of it is Physical AI now. Given how overloaded the term has become, it’s worth asking what ties all of this together, and how.
This article explores the characteristics that define Physical AI methods and organizes them under the quintessential physical framework: control theory. 2
How Is Physical Different from Non-Physical AI?
Let’s begin by defining what differentiates Physical from non-Physical AI.
Physics constraints. The first and probably most accepted characteristic is that Physical AI must respect the governing equations of physics. A model trained on simulated or observed data must transfer to reality while still adhering to physics. There’s no room for hallucinations in Physical AI.
Data scarcity and the sim-to-real gap. Physical AI data is often scarce. Unlike LLMs, which can rely on nearly unlimited text, physical experiments are slow and expensive. Manufacturing and measuring PCBs can take weeks; growing and sequencing cells can take months. Simulation can accelerate data collection, but it comes with the problem of alignment to reality. Calibrating a simulator can be difficult, and at times impossible. Since physical data is rare and expensive, Physical AI models must learn in limited data regimes while staying flexible enough to transfer to the real world.
Sensing, embodiment, and the real-time closed loop. As living beings experience reality through different senses, Physical AI systems are expected to handle data from different sensors and modalities (e.g., vision, lidar, force, temperature). Physical AI continuously operates in the real world, and it is expected to behave as a closed-loop system.
This closed loop is exactly what control theory was built to capture, and it gives us a natural taxonomy for Physical AI itself.
A Taxonomy from the Control Loop
Control theory provides a clean framework for physical systems: a system senses its state, decides on an action, and acts in order to optimize some cost function.
w(t) ⊕ e(t) Controller / Actuator Control Optimize u(t) ⊕ d(t) u s (t) System Model y(t) Sensor Perception y m (t) Fig. 2 — Mapping Physical AI to closed-loop control.
Illustration of a closed-loop control system consisting of
set point w ( t ),
measured output y m ( t ),
measured error e ( t ),
controller output u ( t ),
system input u s ( t ),
disturbance d ( t ),
and system output y ( t ).
Colored blocks show where each Physical AI branch operates within the loop.
This naturally fits Physical AI methods that can be grouped into four categories mapping onto this same loop.
Perception estimates the physical states from raw observations, taking in sensor data (e.g., camera, lidar, force, chemical) and turning them into state representations. This is the territory of computer vision, depth estimation, and state estimation.
Model characterizes how the system and the world evolve. It learns physical dynamics, taking state and action as input and predicting future states. Neural operators, PINNs, surrogate models, and digital twins all fall under this group.
Control addresses what the system should do next. Given the current state, it decides on and executes an action in real time. This is where policy learning, reinforcement learning, and optimal control live.
Optimize determines what configuration achieves the goal. Starting from a target, it searches for or designs a solution with the desired properties. This is the domain of Bayesian optimization, inverse design, and generative design.
Tab. 1 — Physical AI architectures mapped to the Perception–Control–Model–Optimize taxonomy. The table is illustrative rather than comprehensive: entries are representative examples of each combination, not an exhaustive catalog. Hover over a block to see seminal papers.
The building blocks of this taxonomy are methods that sit within a single category. For instance, state estimation and vision models in Perception (ViT, SAM, DINO, YOLO); policy learning methods in Control (Diffusion Policy, ACT, BESO); surrogate models in Model (Neural ODE, GNS, GraphCast); and Bayesian optimization in Optimize (GP, TuRBO, SAASBO).
More flexible approaches spanning two branches are architectures that live at the Model–Optimize boundary (e.g., PINNs, Neural Operators, and Generative Design models), each capable of running forward as a surrogate or backward as an inverse solver; or VLAs, which blend Perception and Control into a single model (e.g., GR00T, RT-2, π₀).
Finally, World Models (e.g., Cosmos, DreamerV3, JEPA) span three categories, blending Perception, Model, and Control entirely in latent space. They encode observations, predict future states, and plan over imagined solutions without interacting with the real world.
This taxonomy brings some order to the crowded Physical AI space where methods that rarely appear in the same
conversation all find a place under the same umbrella.
This table also highlights that a fundamental difference between control theory and Physical AI is that AI models allow us to cross boundaries between functional categories. For instance, classical robotic stacks maintain separate vision, state estimation, and policy modules while a VLA collapses them into a single model. Similarly, PINNs are capable of both modeling dynamical systems and solving inverse problems depending on the direction in which a model is queried.
Reaching for the Fourth Column
A pattern emerges from the proposed framework. Classical techniques sit in a single column. More recent methods occupy two columns: PINNs run forward as surrogates and backward as inverse solvers, VLAs merge Perception and Control. World Models cover three columns. The extrapolated question is then: is a four column model next? A single foundation model that perceives, models, controls, and optimizes. Self-driving labs 3 and auto-research 4 , 5 are some of the closest examples we have today: systems that propose a design, run experiments, observe results, and decide what to try next.
However, to the best of my knowledge, neither is quite there. Self-driving labs don’t yet run on a single model, but on a combination of specialized ones stitched together. And auto-research hasn’t yet made direct contact with the physical world.
And this is where Physical AI stands. The robotic hand is reaching but the fingers haven’t met yet. Physical heaven is still a few centimeters away.
Creation of Robotic Adam , Science Photo Library.
sciencephoto.com/media/92834 .
K.J. Åström and R.M. Murray, Feedback Systems: An Introduction for Scientists and Engineers , Princeton University Press, 2nd ed., 2021.
N.J. Szymanski et al., An autonomous laboratory for the accelerated synthesis of novel materials , Nature 624, 86–91, 2023.
nature.com/articles/s41586-023-06734-w .
A. Novikov et al., AlphaEvolve: A coding agent for scientific and algorithmic discovery , Google DeepMind, 2025.
deepmind.google .
A. Karpathy, autoresearch , 2026.
github.com/karpathy/autoresearch .

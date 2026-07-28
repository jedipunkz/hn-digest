---
source: "https://blog.bytebytego.com/p/how-nvidia-builds-open-models-for"
hn_url: "https://news.ycombinator.com/item?id=49082499"
title: "How Nvidia Builds Open Models for the Age of AI"
article_title: "How NVIDIA Builds Open Models for the Age of AI"
author: "svobodamartin"
captured_at: "2026-07-28T12:48:57Z"
capture_tool: "hn-digest"
hn_id: 49082499
score: 2
comments: 0
posted_at: "2026-07-28T11:57:36Z"
tags:
  - hacker-news
  - translated
---

# How Nvidia Builds Open Models for the Age of AI

- HN: [49082499](https://news.ycombinator.com/item?id=49082499)
- Source: [blog.bytebytego.com](https://blog.bytebytego.com/p/how-nvidia-builds-open-models-for)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T11:57:36Z

## Translation

タイトル: Nvidia が AI 時代のオープン モデルを構築する方法
記事のタイトル: NVIDIA が AI 時代に向けたオープン モデルを構築する方法
説明: NVIDIA の応用深層学習研究担当副社長である Bryan Catanzaro 氏が、彼のチームがどのように会社のオープン モデルを構築しているか、そのアーキテクチャの背後にある理由、そして NVIDIA がその多くをオープンソース化している理由について説明しました。

記事本文:
NVIDIA が AI 時代に向けたオープン モデルを構築する方法
ByteByteGo ニュースレター
NVIDIA が AI 時代に向けたオープン モデルを構築する方法
ByteByteGo 2026 年 7 月 27 日 236 1 16 シェア チェーン可能なコンピューティング。まさに行列です。 (スポンサー付き)
Render の軽量 SDK を使用してタスクを定義し、それらを長時間実行される分散ワークフローにチェーンします。エージェントを起動し、オンデマンドでジョブをバッチ処理します。レンダリング ワークフローは、キューイング、オーケストレーション、および再試行を処理します。
コード BYTE で $50 クレジットを獲得
NVIDIA が GPU を販売していることはおそらくご存知でしょう。しかし、同社はオープン AI モデルの世界最大のパブリッシャーでもあることをご存知ですか?そのモデルは、Hugging Face で最もダウンロードされているものの 1 つにランクされており、そのラインナップはチャットボットをはるかに超えており、推論モデル、世界モデル、人型ロボットのモデル、自動運転車、さらには創薬、量子コンピューティング、世界予測まで多岐にわたります。
これにより 2 つの疑問が生じます。ハードウェアで最もよく知られている企業が、どのようにしてこれほど多くのドメインにわたってこれほど多くの強力なモデルを構築しているのでしょうか?同じモデルが NVIDIA が販売する GPU で実行できるのに、なぜ無料で配布するのでしょうか?
両方を理解するために、NVIDIA の応用深層学習研究担当副社長である Bryan Catanzaro 氏に話を聞きました。彼は、彼のチームが会社のオープン モデルをどのように構築しているか、そのアーキテクチャの背後にある理由、そして NVIDIA がその多くをオープンソース化している理由について説明してくれました。時間を割いてチームの作業を詳細に共有してくれたブライアンに感謝します。
この記事では、次のことを学びます:
NVIDIA が構築したオープン モデルとその機能
NVIDIA がモデルを高速かつ強力にする方法
1 つの基盤で小規模なチームが多くのモデルを構築できるようにする方法
オープンの実際の意味と、データの公開が重みと同じくらい重要である理由
NVIDIA がすべてを放棄する理由と、チームがオープンに出荷して学んだ教訓
NVIDIA オープン モデル (ハイレベル) NVIDIA のオープン モデル E

コシステム
オープン モデルのスペクトル NVIDIA のオープン モデルは、AI の両端に位置します。一方の端には、画面上に存在するシンボル、言語、コード、推論を扱うモデルがあります。もう 1 つは、ロボットや自動運転車 (AV) が周囲を認識して行動する必要がある物理世界で動作するモデルです。ラインナップのほとんどはこの線に沿ったどこかに当てはまります。
エコシステムを確認する最も簡単な方法は、各ファミリーの目的ごとにエコシステムをグループ化することです。
NVIDIAのオープンモデルのラインナップが一目で分かる 1. 推論モデル
推論モデルは、最終的な答えの前に中間トークンを生成するようにトレーニングされた大規模な言語モデルです。これにより、数学やコーディングなどのタスクを通して考えることができるようになります。最近見られるほとんどのフロンティア モデルは、推論モデルによって強化されています。クロード・オーパス、GLM-5.2、キミK2.7。
推論モデルは答える前に考えます。NVIDIA の推論モデル ファミリは Nemotron と呼ばれます。最初の Nemotron モデルは 2023 年に導入されました。2024 年に、NVIDIA は 3,400 億のパラメーターを備えた次世代をリリースし、2025 年にはより高速なハイブリッド設計が続きました。現行世代の Nemotron 3 は、2025 年末から 2026 年にかけて 3 つのサイズで登場しました。高速で単純なタスクには小型の Nano モデル、より困難な計画と推論には中型の Super モデル、そして最も重い推論を必要とする複雑なタスクには大型の Ultra モデルです。
Nemotron が 2023 年からどのように成長したか 2. 物理 AI
言語モデルは次のトークンを予測するのには優れていますが、物理的な世界は理解できません。彼らは、シャツの袖がどのように折り畳まれるのか、あるいは水の入ったコップが傾いた瞬間にどのように見えるのかを伝えることができません。ロボットや自動運転車にはそのような理解が必要ですが、テキストからは理解できません。
それが世界モデルが提供するものです。次の単語を予測するのではなく、世界と状況を理解します。

現在の世界の状態とアクションを考慮して、次の状態を予測します。現実的で物理的に妥当なビデオを生成し、動きや原因と結果について推論し、ロボットが学習する必要がある動作データを生成できます。
ワールド モデルは、現実世界でキャプチャするには時間がかかる、高価である、または危険な状況のトレーニング データを作成するために使用できるため、非常に便利です。
ワールド モデルは、入力アクションを条件として次の状態を予測します。NVIDIA の CEO であるジェンセン ファンは、これを「物理 AI にとって ChatGPT の瞬間」と呼んでいます。 NVIDIA の物理 AI のフロンティア オープン モデルは Cosmos で、2025 年 1 月の CES で発表され、2026 年に単一のオムニモデルである Cosmos 3 に統合されました。Cosmos 3 は、シーンを生成し、それについて推論し、次に何が起こるかを予測できます。 Cosmos 3 では、事前トレーニングされたビデオ バックボーンをロボット ポリシーに直接変換するワールド アクション モデルも導入されています。ロボット上で実行してリアルタイムで制御できるほど小さい、40 億パラメータの Edge バージョンもあります。
Cosmos はローンチからオムニモデルへ ワールドモデルからロボットポリシーへ ワールドモデルは次のシーンを予測できますが、ロボットは行動しなければなりません。カメラが捉えたものを、誰も事前にスクリプト化していないタスクでリアルタイムにモーターがどのように動くかに変換する必要があります。
VLA はアクションの予測に重点を置いています。これが、ビジョン-言語-アクション モデルと呼ばれることが多いロボット基盤モデルの仕事です。これらのモデルへの入力は通常、カメラ画像と命令であり、モデルが動きを生成します。
ロボットがピクセルをモーションに変換する仕組み NVIDIA の人型ロボット向けのオープン基盤モデル ラインは、Isaac GR00T です。 NVIDIA は 2024 年にこのプロジェクトを発表し、2025 年初めに最初のオープン ヒューマノイド基盤モデル GR00T N1 をリリースし、それ以来 GR00T 1.7 まで繰り返してきました。
GR00T はアイデアから 1.7 へ GR00T モデルは Cosmos を使用します

これは推論のバックボーンであり、ロボットに知覚、推論、移動操作の共通基盤、つまり物体を操作しながら移動する能力を与えます。これにより、より人間らしい意思決定が可能になり、ロボットが複雑なアクションを構造化された段階的な計画に分割し、開発者が特定のタスクに特化できるようになります。 1.7 リリースには、運用に適した商用ライセンスも含まれており、開発者はモデルを実際のアプリケーションに展開できます。
異なる作業に同じベースロボットを使用 3. ヘルスケア、自動運転車、量子、およびより特殊なアプリケーション
ラインナップはまだまだ続きます。 Alpamayo は、推論ベースの自動運転のためのオープン モデル ファミリです。 GR00T と同様に、推論バックボーンとして Cosmos を使用しているため、ロボットを駆動するのと同じ物理 AI 基盤により、車は訓練されていない道路状況でも作業できるようになります。また、各決定の背後にある原因と結果の連鎖も明らかになり、開発者はそれを検査して構築することができます。 BioNeMo は生物学と創薬にオープン モデルをもたらし、BioNeMo エージェント ツールキットを使用すると、開発者はライフ サイエンス作業用の自律エージェントを構築できます。 Ising は、量子コンピューターの構築者、オペレーター、開発者に、デバイスをフォールト トレランスに拡張するために必要な AI ツールを提供します。 Earth-2 は、プロフェッショナル グレードの気象お​​よび気候 AI を誰でも利用できるようにするオープン モデル、ライブラリ、およびフレームワークのファミリーです。これらに加えて、NVIDIA は、安全でないコンテンツをフィルタリングする安全モデル、音声のスピーチ モデル、画像のビジョン モデルなど、実稼働用のオープン モデルを提供します。
NVIDIA モデルが AI アプリをどのように機能させるか この強力なエコシステムを組み合わせることで、コミュニティが前進し構築する方法が変わります。この記事の残りの部分では、NVIDIA がこれらのモデルをどのように構築するか、チームがどのように迅速に行動できるか、そしてチームが何を学んだのかについて説明します。
モデルの様子

フロンティアかつ高速に動作するように構築
多くの場合、高速かつ高機能なモデルを構築するのは困難です。通常、有能なモデルはサイズが大きいため、トレーニングや質問への回答に時間がかかります。 NVIDIA の答えは、フロンティアとスピードを同時に追求することです。ブライアンは、「最速のモデルが最もスマートなモデルである」と述べています。
これはスローガンのように聞こえるかもしれませんが、実際の洞察を持っています。より高速なモデルは、同じ時間内により多くのデータでトレーニングできます。より多くの環境で事後トレーニングが可能です。導入すると、同じコストで難しい問題についてより長く考えることができます。したがって、スピードはあらゆる段階で能力と複合性につながります。
この戦略を踏まえ、NVIDIA は効率と速度に重点を置いてモデルを設計しています。次の設計上の選択により、これらのモデルは高速かつ高機能になります。最初の 2 つは、モデルをより高速かつ効率的にします。最後のものはそれをより有能なものにします。
ほとんどのよく知られたモデルは、ほぼ完全に Transformer の注意メカニズムに基づいて構築されています。アテンションを使用すると、入力内のすべてのトークンが他のすべてのトークンを参照できるようになります。これは関係を把握するのに強力ですが、コストがかかります。コストは入力長の 2 乗に応じて増加するため、入力が 2 倍になると作業量は約 4 倍になります。モデルは、実行中に以前のすべてのトークンをメモリ内に保持する必要があります。非常に長い入力の場合、速度が遅くなり、コストが高くなります。
アテンションは急速に高価になる もう 1 つの主要なクラスは状態空間モデルであり、Mamba はその最もよく知られた例です。すべてのトークンを他のすべてのトークンと比較する代わりに、Mamba レイヤーはシーケンスを 1 回読み取り、認識したものすべてを固定サイズのメモリに圧縮します。そのためコストが安くなります。コストは長さに応じて直線的に増加し、入力がどれほど長くてもメモリは一定のままです。トレードオフは、固定サイズのメモリではすべてのデットを保持できないことです。

まったく、マンバははるか昔に埋もれていた正確な事実を思い出すのが苦手なのです。
Mamba は入力を固定メモリに押し込みます Mamba レイヤーは効率的ですが、アテンション レイヤーはより効果的です。それらを組み合わせると、両方の長所が得られます。ハイブリッド アーキテクチャでは、Transformer のアテンション レイヤーの大部分を Mamba レイヤーと交換し、完全なアテンションを備えたいくつかのブロックを保持します。
アテンション層とマンバ層は交換可能 NVIDIA のモデルは同じハイブリッド パターンに依存しています。レイヤーのほとんどは Mamba であり、長い入力の処理を安価に保ち、100 万トークンのコンテキスト ウィンドウを単なる理論的ではなく実用的なものにします。 Mamba が放棄したピンポイントの再現を復元するために、いくつかのアテンション レイヤーが間に配置されているため、モデルはコンテキスト内のどこからでも正確な詳細を取得できます。
NVIDIA のハイブリッド アーキテクチャ ハイブリッド設計に加えて、モデルでは専門家混合 (MoE) レイヤーも使用されます。 MoE レイヤーは多数の小さなエキスパート レイヤーで構成され、各トークンはそのうちの少数のレイヤーのみにルーティングされます。したがって、モデルはすべてのトークンですべてのパラメーターを実行するのではなく、パラメーターの小さなサブセットのみをアクティブにします。これにより、低いトークンあたりのコストで大きな総容量を使用できるようになり、モデルの高速性が維持されます。
MoE: トークンごとにアクティベートするのは少数の専門家だけです ブライアンが説明するように、Mamba はシーケンス全体のグローバルなビューを構築します。注意はいつでも遡って正確な事実を見つけることができます。これらを組み合わせると、どちらか一方を単独で使用するよりも賢くなります。 NVIDIA は 2024 年にこのハイブリッドの結果を発表し、それ以来、Qwen と Kim の背後にあるチームを含む他のラボも同じ方向に進んでいます。
2. GPU とモデルの共同設計
2 番目の選択肢は、NVIDIA の 2 つのビジネスが出会う場所です。その大規模なモデルは、NVFP4 と呼ばれる 4 ビット数値形式で事前トレーニングされています。つまり、トレーニング中のほとんどの計算では、

値ごとに 4 ビット。ビット数が少ないほど、メモリと移動するデータが少なくなるため、計算はより高速に、より少ない電力で実行されます。
精度形式の比較 4 ビットでは精度が非常に低いことに注意してください。ほとんどのチームは、より高い精度でトレーニングし、その後モデルを縮小するだけなので、精度が失われます。 NVIDIA は、次世代 GPU である Blackwell が高速 4 ビット ハードウェアで構築されていることを知っていたため、最初のステップから 4 ビットでトレーニングしました。モデルとチップは相互に設計されています。
副社長はその背後にある考え方を説明しました。ムーアの法則はもはや世代ごとに簡単に利益をもたらすものではないため、進歩はモデルとハードウェアを一緒に設計することからもたらされる必要があります。チームは、結果が可能であると信じていたため、4 ビットの事前トレーニングのみを試みました。彼の言葉によれば、何か素晴らしいものを発明するための最初のステップは、自分にはできると信じることです。
3. トレーニング後: 経験から学ぶ
これらの効率に関する選択肢が設定されている場合、トレーニング後はモデルの能力を獲得する場所です。
トレーニング後のレシピ自体は、他のほとんどのチームが集まったものです。それは教師あり微調整から始まり、モデルは厳選された良い回答の例から学習し、それに期待される形式と動作を選択します。その後、強化学習に切り替わります。

[切り捨てられた]

## Original Extract

Bryan Catanzaro, VP of Applied Deep Learning Research at NVIDIA, walked us through how his team builds the company’s open models, the reasoning behind their architecture, and why NVIDIA open-sources so much of it.

How NVIDIA Builds Open Models for the Age of AI
ByteByteGo Newsletter
Subscribe Sign in How NVIDIA Builds Open Models for the Age of AI
ByteByteGo Jul 27, 2026 236 1 16 Share Chainable compute. Right on queue. (Sponsored)
Define tasks with Render’s lightweight SDK and chain them into long-running, distributed workflows. Launch your agents and batch jobs on demand. Render Workflows handles queuing, orchestration, and retries.
Get $50 credits with code BYTE
You probably know NVIDIA sells GPUs. But did you know it is also the largest publisher of open AI models in the world ? Its models rank among the most downloaded on Hugging Face, and the lineup goes way beyond chatbots: reasoning models, world models, models for humanoid robots, self-driving cars, even drug discovery, quantum computing, and global forecasting.
This raises two questions. How does a company best known for hardware build so many strong models across so many domains? And why give them away for free, when those same models run on the GPUs NVIDIA sells?
To understand both, we spoke with Bryan Catanzaro , VP of Applied Deep Learning Research at NVIDIA. He walked us through how his team builds the company’s open models , the reasoning behind their architecture, and why NVIDIA open-sources so much of it. Our thanks to Bryan for taking the time to share the team’s work in such detail.
In this article, you’ll learn:
What open models NVIDIA has built, and what they do
How NVIDIA makes its models both fast and powerful
How one foundation lets a small team build many models
What open really means and why releasing the data matters as much as the weights
Why NVIDIA gives it all away, and the lessons the team has learned shipping in the open
NVIDIA Open Models (High-Level) NVIDIA’s Open Model Ecosystem
Spectrum of open model NVIDIA’s open models sit at two ends of AI. At one end are models that work with symbols, the language, code, and reasoning that lives on a screen. At the other are models that work in the physical world, where robots and autonomous vehicles (AVs) have to perceive their surroundings and act. Most of the lineup falls somewhere along that line.
The simplest way to see the ecosystem is to group them by what each family is for.
NVIDIA’s lineup of open models at a glance 1. Reasoning models
Reasoning models are large language models that are trained to produce intermediate tokens before the final answer. This is what lets them think through tasks such as math and coding. Most frontier models you see these days are powered by a reasoning model. Claude Opus, GLM-5.2, Kimi K2.7.
Reasoning models think before they answer NVIDIA’s reasoning model family is called Nemotron. The first Nemotron model was introduced in 2023. In 2024, NVIDIA released the next generation with 340 billion parameters, followed by a faster hybrid design in 2025. The current generation, Nemotron 3, arrived across late 2025 and 2026 in three sizes. A small Nano model for fast, simple tasks, a mid-size Super model for harder planning and reasoning, and a large Ultra model for complex tasks that need the heaviest reasoning.
How Nemotron grew since 2023 2. Physical AI
Language models are great at predicting the next token, but they do not understand the physical world. They cannot tell you how a shirt sleeve folds, or what a glass of water looks like the moment it tips over. A robot or a self-driving car needs that kind of understanding, but it cannot get it from text.
That is what a world model provides. Instead of predicting the next word, it understands the world and can predict the next state given the current state of the world and an action. It can generate realistic, physically plausible video, reason about motion and cause and effect, and produce the action data a robot needs to learn.
World models are extremely useful because you can use them to create training data for situations that are slow, expensive, or dangerous to capture in the real world.
A world model predicts the next state conditioned on an input action NVIDIA’s CEO, Jensen Huang, has called this “the ChatGPT moment for physical AI.” NVIDIA’s frontier open model for physical AI is Cosmos, launched at CES in January 2025 and unified in 2026 into a single omni-model , Cosmos 3, that can generate a scene, reason about it, and predict what comes next. Cosmos 3 also introduces world action models , which turn the pretrained video backbones directly into robot policies. There is also a 4-billion-parameter Edge version small enough to run on the robot and control it in real time.
Cosmos went from launch to omni-model From World Models to Robot Policies A world model can predict the next scene, but a robot has to act. It has to turn what its cameras see into how its motors move, in real time, on tasks nobody scripted in advance.
VLA focuses on predicting actions That is the job of a robot foundation model, often called a vision-language-action model. The inputs to these models are usually camera images and instructions, and the model produces the movement.
How a robot turns pixels into motion NVIDIA’s open foundation model line for humanoid robots is Isaac GR00T. NVIDIA unveiled the project in 2024, released the first open humanoid foundation model, GR00T N1, in early 2025, and has since iterated to GR00T 1.7.
GR00T went from idea to 1.7 GR00T models use Cosmos as a reasoning backbone, giving robots a shared foundation for perception, reasoning, and locomanipulation: the ability to move while manipulating objects. This enables more human-like decision-making and helps robots break complex actions into structured, step-by-step plans that developers can specialize for specific tasks. The 1.7 release also includes a production-friendly commercial license, allowing developers to deploy the model in real-world applications.
Same base robot for different jobs 3. Healthcare, autonomous vehicles, quantum, and more specialized applications
The lineup keeps going. Alpamayo is an open model family for reasoning-based self-driving. Like GR00T, it uses Cosmos as its reasoning backbone, so the same physical-AI foundation that powers robots also lets a car work through road situations it was never trained on. It also exposes the chain of cause and effect behind each decision, which developers can inspect and build on. BioNeMo brings open models to biology and drug discovery, and the BioNeMo Agent Toolkit lets developers build autonomous agents for life sciences work. Ising gives quantum computer builders, operators, and developers the AI tools needed to scale devices to fault tolerance. Earth-2 is a family of open models, libraries, and frameworks that makes professional-grade weather and climate AI available to anyone. On top of these, NVIDIA offers open models for production use: safety models that filter unsafe content, speech models for voice, and vision models for images.
How NVIDIA models make AI apps work Taken together, this powerful ecosystem changes how the community can advance and build. The rest of the article looks at how NVIDIA builds these models, how the team is able to move fast, and what they have learned.
How the Models Are Built to Be Frontier and Fast
It is often a challenge to build a model that is both fast and capable. Capable models are usually larger, which makes them slower to train and to answer questions. NVIDIA’s answer is to chase the frontier and speed at the same time. Bryan put it this way: “the fastest model is the smartest model.”
While this may sound like a slogan, it has a real insight. A faster model can be trained on more data in the same amount of time. It can be post-trained across more environments. Once deployed, it can think for longer on a hard problem at the same cost. So speed translates into capability and compounds at every stage.
Given this strategy, NVIDIA designs its models with a focus on efficiency and speed. The following design choices are what make these models both fast and capable. The first two make the model faster and more efficient. The last one makes it more capable.
Most well-known models are built almost entirely on the Transformer’s attention mechanism. Attention lets every token in the input look at every other token, which is powerful for capturing relationships but expensive. The cost grows with the square of the input length, so doubling the input roughly quadruples the work. The model has to keep every earlier token in memory as it runs. On very long inputs, that gets slow and costly.
Attention gets expensive fast The other main class is the state space model, and Mamba is its best-known example. Instead of comparing every token with every other one, a Mamba layer reads the sequence once and compresses everything it has seen into a fixed-size memory. That makes it cheap: the cost grows linearly with length, and the memory stays constant no matter how long the input. The tradeoff is that a fixed-size memory cannot keep every detail, so Mamba is weaker at recalling one exact fact buried far back.
Mamba squeezes input into fixed memory Mamba layers are efficient while attention layers are more effective. Combining them gives the best of both worlds. A hybrid architecture swaps most of the attention layers in the Transformer with Mamba layers and keeps a few blocks with full attention.
Attention and Mamba layers can be swapped NVIDIA’s models rely on the same hybrid pattern. Most of the layers are Mamba, which keeps long inputs cheap to process and is what makes a million-token context window practical rather than just theoretical. A few attention layers are placed in between to restore the pinpoint recall that Mamba gives up, so the model can still retrieve exact details from anywhere in the context.
NVIDIA’s hybrid architecture On top of the hybrid design, the models also use mixture-of-experts, or MoE, layers. An MoE layer consists of many small expert layers and routes each token to only a few of them. So instead of running every parameter on every token, the model activates only a small subset of its parameters. That allows us to use a large total capacity at a low per-token cost, which keeps the model fast.
MoE: Only a few experts activate per token The way Bryan describes it, Mamba builds a global view of the whole sequence. Attention can always go back and find an exact fact. Put together, they are smarter than either one alone. NVIDIA published this hybrid result in 2024, and other labs, including the teams behind Qwen and Kimi, have since moved in the same direction.
2. Co-designing the GPU and the model
The second choice is where NVIDIA’s two businesses meet. Its larger models are pretrained in a 4-bit number format called NVFP4, meaning most of the math during training uses just four bits per value. Fewer bits means less memory and less data to move, so the math runs faster and on less power.
A comparison of precision formats Note that four bits is very little precision. Most teams train in higher precision and only shrink the model afterward, which loses accuracy. NVIDIA trained in 4 bits from the first step because it knew its next GPU generation, Blackwell, was being built with fast 4-bit hardware. The model and the chip were designed for each other.
The VP described the mindset behind it. Moore’s law is no longer giving easy gains every generation, so progress has to come from designing the model and the hardware together. The team only attempted 4-bit pretraining because it believed the result was possible. In his words, the first step to inventing something amazing is believing that you can.
3. Post-training: learning from experience
With those efficiency choices in place, post-training is where the model gains its capability.
The post-training recipe itself is what most other teams converged to. It starts with supervised fine-tuning, where the model learns from curated examples of good answers and picks up the format and behavior expected of it. Then it switches to reinforcement learning, where the mo

[truncated]

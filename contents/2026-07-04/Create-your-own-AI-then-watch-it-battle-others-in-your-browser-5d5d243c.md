---
source: "https://kim-ai-gpu.github.io/2026/07/04/introducing-agenlus-browser-rl/"
hn_url: "https://news.ycombinator.com/item?id=48784825"
title: "Create your own AI, then watch it battle others in your browser"
article_title: "Create your own AI, then watch it battle others in your browser | YoungSeong Kim"
author: "umjunsik132"
captured_at: "2026-07-04T13:00:43Z"
capture_tool: "hn-digest"
hn_id: 48784825
score: 1
comments: 0
posted_at: "2026-07-04T12:10:56Z"
tags:
  - hacker-news
  - translated
---

# Create your own AI, then watch it battle others in your browser

- HN: [48784825](https://news.ycombinator.com/item?id=48784825)
- Source: [kim-ai-gpu.github.io](https://kim-ai-gpu.github.io/2026/07/04/introducing-agenlus-browser-rl/)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T12:10:56Z

## Translation

タイトル: 独自の AI を作成し、ブラウザーで他の AI と戦うのを見てください。
記事のタイトル: 独自の AI を作成し、ブラウザーで他の AI と戦うのを見てください |キム・ヨンソン
説明: 簡単な入門書: 強化学習 (RL) とは何ですか?

記事本文:
キム・ヨンソン
私について
独自の AI を作成し、ブラウザーで他の AI と戦う様子を観察します。
簡単な入門書: 強化学習 (RL) とは何ですか?
大規模言語モデル (LLM) または教師あり学習を主に使用してきた場合、RL は考え方の変化です。
教師あり学習 (テキスト ファイル内の次のトークンの予測や画像の分類など) は、静的な正解の静的なデータセットに依存します。
強化学習は、アクティブなフィードバック ループに依存します。エージェント (AI モデル) は環境 (ゲームまたはシミュレーション) と対話します。 Action を受け取り、Observation (環境の新しい状態) と Reward (エージェントがどの程度うまくいっているかを伝える信号) を受け取り、これを繰り返します。
エージェントの目標は、時間の経過とともに累積報酬を最大化するポリシー (観察からアクションへのマッピング) を学習することです。完全にランダムに始まり、純粋に試行錯誤を通じて改善されます。
事前に提供される「正解」がないため、RL エージェントは、開発者が予期していなかったゲームを解決するための非常に賢明で新しい方法を見つけることがよくあります。ただし、この試行錯誤のプロセスは計算量が多く、何百万回もの操作が必要となるため、これをブラウザに直接表示することは、挑戦的であり、刺激的でもあります。
強化学習 (RL) は、常に AI の最も魅力的な分野の 1 つです。白紙のエージェントが環境を探索し、徐々に超人的な政策を打ち出していくのを見るのには、深い満足感があります。
しかし、LLM の遊び場やツールの爆発的な成長に比べ、RL は依然として比較的アクセスしにくいものです。ローカル環境のセットアップは、多くの場合、Python 仮想環境、CUDA バージョン、PyTorch インストール、Gymnasium でのヘッドレス レンダリングのバグと格闘することを意味します。
私たちはこれを解決するために Agenlus を構築しました。それは

強化学習用のコミュニティ プラットフォームおよびモデル ハブは、ブラウザ内で完全に実行できるように設計されています。インストールや CUDA 構成は不要で、即座にトレーニングと評価を行うだけです。
強化学習の民主化
過去 10 年間、最先端の強化学習は、エリート企業の研究室や資金豊富な学術機関の独占的な遊び場でした。 Google DeepMind の AlphaGo、OpenAI の Dota 2 ボット、または高度な産業用ロボット制御のいずれであっても、RL は大規模なコンピューティング クラスター、複雑なシミュレーターのセットアップ、および特殊な数学的専門知識へのアクセスを必要としました。
この集中化により、独立した開発者や研究者の創造的な可能性が抑圧されてきました。誰でもオンラインで大規模な言語モデルを簡単にプロンプ​​トできますが、RL を使い始めるには、複雑なローカル設定、GPU ドライバー、ローカル仮想化に取り組む必要があり、単純なエージェントが収束するまでに数時間待つだけです。
私たちはRLを民主化する必要があると信じています。
最新の Web テクノロジーを活用することで、次のような障壁を打ち破りたいと考えています。
参入障壁を下げる : RL を試すのにハイエンドのローカル マシンや AWS の予算は必要ありません。ブラウザをお持ちであれば、完全に機能する RL 研究所が手に入ります。
オープンソース環境の共有 : Hugging Face がモデルを共有しやすくすることで NLP を民主化したのと同じように、Agenlus を使用すると、開発者は環境を即座にアップロード、共有、ベンチマークすることができます。
インタラクティブな学習 : トレーニング プロセスがブラウザーでライブで行われるのを確認することで、エージェントのポリシーが報酬にどのように適応するかについて深く直感的な理解を構築します。
RL のツールを世界中の開発者コミュニティの手に直接渡すことで、企業の研究室が見落とす可能性のある新しい制御アーキテクチャとアルゴリズムの発見を加速することを目指しています。
B2C RL が現在非常に有効である理由

私たちは現在、LLM による膨大なコンピューティングのインフレを目の当たりにしています。このため、B2C AI スタートアップの構築は信じられないほど高価になり、創設者はクラウド GPU の巨額の請求書を支払うか、ベンチャー キャピタルで数百万ドルを調達するかの選択を迫られています。
私たちは、強化学習 (RL) が構造的にこのサイクルを打破し、収益性の高い B2C AI アプリケーションの新たな波を導く準備ができていると考えています。その理由は次の 3 つです。
限界インフラストラクチャコストゼロ: すべての推論トークンに API クレジットがかかる LLM とは異なり、Agenlus の RL トレーニングと推論は、WebGPU を介してユーザーのクライアント ハードウェア上で 100% ローカルに実行されます。当社のサーバーコストは事実上ゼロです。これにより、数百万のアクティブ ユーザーにスケールし、コンピューティング クレジットを使い果たすことなく永続的な無料利用枠を提供できるようになり、収益化の焦点をマーケットプレイスのトランザクションとカスタム アセットに移すことができます。
極端なモデル効率 : まともな LLM には数十億のパラメータが必要ですが、ゲーム用の高性能 RL エージェント (複雑な 2D/3D プラットフォーマーや制御タスクであっても) は信じられないほど軽量です。多くの場合、超人的なポリシーを達成するには、パラメータが 10 万未満の小規模な多層パーセプトロン (MLP) または小さな畳み込みニューラル ネットワーク (CNN) で十分です。これらのモデルは即座にロードされ、エントリーレベルのモバイル デバイスやラップトップで 1 秒あたり数百のステップを実行します。
ゲーミフィケーションとナチュラル バイラル ループ : 生成 AI ツールは主に生産性を重視しています。対照的に、RL エージェントのトレーニングは本質的にゲーム化されています。デジタルペット (たまごっちなど) を育てたり、スポーツ チームをコーチしたりするような感じです。エージェントを直接制御から自律トレーニングに移行させることで、ユーザーはエージェントのプレイスタイルに深い感情的な愛着を抱くようになります。競争力のあるリーダーボードとマルチエージェント PvP アリーナを追加すると、自然でバイラルなソーシャル ループ (「私の年齢」) が作成されます。

これにより、高額な顧客獲得コストをかけずに本的成長を推進できます。
オフライン データの未来をクラウドソーシング : ロボット工学などの業界における RL の最大のボトルネックは、多様で高品質なデモンストレーション データセットの欠如です。ユーザーがエージェントをトレーニングするためにプレイする B2C プラットフォームを構築することで、私たちは何千もの環境にわたる人間の行動軌跡の膨大なライブラリをクラウドソーシングしています。この多様なデータセットは、複数の制御ドメインにわたって一般化する将来の基礎モデルをトレーニングするための宝の山です。
以下は、動作中のプラットフォームの簡単な YouTube デモです。
以下は、クライアント側のトレーニング ループを実行する Web UI のプレビューです。
ここでは、ブラウザ内 RL を実現するためのアーキテクチャを構築した方法、遭遇した UX の課題、そして RL が B2C の復活に最適であると当社が考える理由を説明します。
アーキテクチャ: WebGPU + Pyodide + Web ワーカー
インストール不要のエクスペリエンスを実現するには、環境シミュレーションとモデル トレーニングの両方をクライアント ブラウザーに移行する必要がありました。
私たちのコア アーキテクチャは、負荷を 3 つの部分に分割します。
ブラウザコンテキスト
§── Web ワーカー (Pyodide)
│ §── 体育館（マイクロピップ）
│ §── env.step(action) -> 観察、報酬、終了
│ └── 状態の更新をメインスレッドに送信します
│
§── メインスレッド (WebGPU & JS)
│ §── モデル推論とポリシーの更新
│ §── アクションの選択 (WebGPU で高速化された tensor 経由)
│ └── 描画コマンドブリッジ（Canvasへのレンダリング）
1. Web ワーカーのピオダイド
CartPole や MountainCar などの公式の体育館環境を実行するには、Pyodide (WebAssembly にコンパイルされた Python) を使用します。
Python の実行はシングルスレッドであるため、ブラウザーで Python 環境を実行すると、UI がフリーズしやすくなります。 Pyodide を専用の Web Worker にオフロードすることで、

メイン UI スレッドはスムーズな 60fps で実行されます。
2. WebGPU 推論および描画コマンド ブリッジのメイン スレッド
ニューラル ネットワークの更新と転送パスのためにメイン スレッドで WebGPU を利用します。
Pyodide によって返される観測状態は単純であるため (たとえば、 CartPole の 4 つの float 値)、Web Worker と WebGPU コンテキストの間でデータを受け渡す際のシリアル化オーバーヘッドは無視できます。
環境の視覚化に関して、標準的な Gymnasium 環境は Pygame に大きく依存しています。Pygame はブラウザ Canvas に直接レンダリングできず、WASM から生のピクセル フレーム バッファー ( rgb_array ) をやり取りすると、大量のパフォーマンス オーバーヘッドが発生します。
これを回避するために、Pyodide に Pygame Mocking Bridge を構築しました。
擬似 pygame モジュールを動的に挿入することで、Python 環境内のすべての pygame.draw 呼び出し (rect 、circle 、 line など) をインターセプトします。
モック モジュールは、Python 描画操作を軽量の JSON 描画コマンドのリストに変換します (例: {"type": "circle", "color": [255, 0, 0], "center": [100, 50], "radius": 10} )。
これらの命令は Web ワーカーからメイン スレッドに送信され、そこでネイティブ HTML5 Canvas が描画します。
この妥協により、グラフィックス パイプラインの超高速性とハードウェア アクセラレーションを維持しながら、標準の Gymnasium 環境レンダリング コードとの完全な互換性が維持されます (環境内での pygame のインポートは引き続き機能します)。
ブラウザの制約を克服する
完全にクライアント側の ML トレーニング ループを構築するには、特有の摩擦点が伴います。私たちがそれらに対処した方法は次のとおりです。
1. 15 秒のピオダイド コールド スタート
Pyodide のロード、Python ランタイムの初期化、micropip 経由の依存関係 ( numpy や scipy など) のインストールには、ユーザーの接続に応じて 10 ～ 30 秒かかる場合があります。
このレイテンシを非表示にするには:
バックグラウンドプリロード : W をスピンアップします

eb Worker を使用すると、ユーザーがプラットフォームにアクセスした瞬間に、たとえ環境のカタログを閲覧しているだけであっても、バックグラウンドで Pyodide の読み込みが開始されます。
Service Worker のキャッシュ : WASM バイナリと pip パッケージはロードされると IndexedDB/Cache API 経由で積極的にキャッシュされるため、以降のアクセスは即座にロードされます。
2. セッションの永続化 (タブのクローズ処理)
サーバー側のトレーニングとは異なり、ユーザーがブラウザのタブを閉じると、アクティブなトレーニングの進行状況が失われる危険があります。
現在、進行状況を追跡するために、基本的なトレーニング メタデータ (総エピソード、トレーニング ステップ、最高の報酬など) とローカル デモンストレーション ログをブラウザーの localStorage に保存しています。アクティブなモデル状態が失われる問題を解決するために、すぐにできる次のステップは、自動エージェント チェックポイント設定 (モデルの重みとハイパーパラメータの保存) を IndexedDB に実装することです。これにより、タブが閉じたりクラッシュしたりした後でも、ユーザーがアクティブなトレーニング セッションをシームレスに再開できるようになります。
コアループ: ゲートウェイとしての模倣学習
RL の主な障害はサンプルの非効率です。エージェントはポールのバランスをとる方法を学ぶだけでも何千ものエピソードが必要になることがあります。ブラウザーを使用するカジュアル ユーザーにとって、知性の兆候が現れるまで 10 分間待つことは、リテンション キラーです。
これを修正するためにデュアルモード インタラクション ループを設計しました。
ヒューマン デモンストレーション (直接制御) : ユーザーはキーボード/マウスを使用してゲームを直接プレイし、デモンストレーションの軌跡を生成できます。
模倣学習 (行動クローニング) : これらの人間によるデモンストレーションを使用して、基本モデルを瞬時にトレーニングします。エージェントはプレイヤーのプレイスタイルを即座に反映します。
RL 微調整 : エージェントがタスクの基本を理解したら、ユーザーは RL アルゴリズム (PPO、DQN) に制御を渡し、人間のパフォーマンスを超えてポリシーを最適化します。
この「人間参加型」アプローチにより、コールドスタートの問題が解決されます。

ゼロからトレーニングするという問題を解決し、ユーザーに「このエージェントは私と同じように、しかしより優れたプレイを学んでいる」という強い当事者意識を与えます。
私たちは Agenlus を単なるブラウザ ゲーム プラットフォームとは考えていません。私たちの最終目標は、次世代の物理制御 AI の基礎インフラを構築することです。
私たちのロードマップは、拡張の 3 つの主要なフェーズにまたがっています。
1. 物理と制御のためのモジュール式サンドボックス (中期)
私たちは、クリエイターが複雑な物理学を使用してカスタム 2D/3D 環境を設計し、ワンクリックでインポートし、カスタム報酬システムを定義できる、ローコードの Web ベースのサンドボックス スタジオを構築しています。環境作成を民主化することで、少数の古典的な制御タスクからコミュニティが生成する何千もの物理的な課題まで拡張したいと考えています。
2. Sim-to-Real: ロボティクス ブリッジ (長期)
物理世界でロボットを訓練するのは時間がかかり、危険で、費用がかかります。シミュレーターが唯一の実行可能な方法です。
Agenlus エージェントはブラウザから直接トレーニングされ、標準形式 (ONNX や TensorFlow.js モデルなど) でエクスポートされるため、物理ハードウェアへのシームレスな展開パイプラインを構築しています。将来的には、Agenlus ブラウザ タブ内で自律ドローン制御ポリシーやロボット アーム マニピュレータをトレーニングし、その正確なポリシーを無線 (OTA) で物理マシンに展開できるようになります。

[切り捨てられた]

## Original Extract

A Quick Primer: What is Reinforcement Learning (RL)?

YoungSeong Kim
About me
Create your own AI, then watch it battle others in your browser
A Quick Primer: What is Reinforcement Learning (RL)?
If you have mostly worked with Large Language Models (LLMs) or Supervised Learning, RL is a shift in mindset:
Supervised Learning (like predicting the next token in a text file or classifying an image) relies on a static dataset of static correct answers.
Reinforcement Learning relies on active feedback loops . An Agent (our AI model) interacts with an Environment (the game or simulation). It takes an Action , receives an Observation (the new state of the environment) and a Reward (a signal telling the agent how well it is doing), and repeats.
The agent’s goal is to learn a Policy (a mapping from observations to actions) that maximizes the cumulative reward over time. It starts completely random and improves purely through trial and error.
Because there is no “correct answer” provided upfront, RL agents often find incredibly clever, emergent ways to solve games that developers never anticipated. However, this trial-and-error process is computationally intensive and requires millions of interactions, which is why bringing it directly to the browser is both challenging and exciting.
Reinforcement Learning (RL) has always been one of the most fascinating branches of AI. There is something deeply satisfying about watching a blank-slate agent explore an environment and gradually emerge with a superhuman policy.
Yet, compared to the explosive growth of LLM playgrounds and tools, RL remains relatively inaccessible. Setting up a local environment often means wrestling with Python virtual environments, CUDA versions, PyTorch installations, and headless rendering bugs in Gymnasium.
We built Agenlus to solve this. It is a community platform and model hub for Reinforcement Learning designed to run entirely in the browser —no installation, no CUDA configuration, just instant training and evaluation.
Democratizing Reinforcement Learning
For the past decade, state-of-the-art Reinforcement Learning has been the exclusive playground of elite corporate labs and well-funded academic institutions. Whether it is Google DeepMind’s AlphaGo, OpenAI’s Dota 2 bots, or sophisticated industrial robotics control, RL has required access to massive compute clusters, complex simulator setups, and specialized mathematical expertise.
This centralization has stifled the creative potential of independent developers and researchers. While anyone can easily prompt a large language model online, starting out with RL requires wrestling with complex local setups, GPU drivers, and local virtualization, only to wait hours for a simple agent to converge.
We believe RL needs to be democratized.
By leveraging modern web technologies, we want to break down these barriers:
Lowering the Entry Barrier : You don’t need a high-end local machine or an AWS budget to experiment with RL. If you have a browser, you have a fully functional RL research lab.
Open-Source Environment Sharing : Just as Hugging Face democratized NLP by making models easy to share, Agenlus allows developers to upload, share, and benchmark environments instantly.
Interactive Learning : Seeing the training process happen live in the browser builds a deep, intuitive understanding of how agent policies adapt to rewards.
By putting the tools of RL directly into the hands of the global developer community, we aim to accelerate the discovery of novel control architectures and algorithms that corporate labs might overlook.
Why B2C RL is Highly Viable Today
We are currently witnessing immense compute inflation dominated by LLMs. This has made building B2C AI startups incredibly expensive, forcing founders to choose between paying massive cloud GPU invoices or raising millions in venture capital.
We believe Reinforcement Learning (RL) is structurally primed to break this cycle and lead a new wave of highly profitable B2C AI applications for three key reasons:
Zero Marginal Infrastructure Cost : Unlike LLMs where every inference token costs API credits, RL training and inference in Agenlus run 100% locally on the user’s client hardware via WebGPU . Our server costs are virtually zero. This allows us to scale to millions of active users and offer a permanent free tier without burning through compute credits, shifting the monetization focus to marketplace transactions and custom assets.
Extreme Model Efficiency : While a decent LLM requires billions of parameters, high-performing RL agents for games (even complex 2D/3D platformers and control tasks) are incredibly lightweight. A small Multi-Layer Perceptron (MLP) or a tiny Convolutional Neural Network (CNN) of under 100K parameters is often enough to achieve superhuman policies. These models load instantly and execute hundreds of steps per second on entry-level mobile devices or laptops.
Gamification and Natural Viral Loops : Generative AI tools are mostly focused on productivity. In contrast, training an RL agent is inherently gamified. It feels like nurturing a digital pet (like a Tamagotchi) or coaching a sports team. By transitioning the agent from direct control to autonomic training, users develop a deep emotional attachment to their agent’s playstyle. When you add competitive leaderboards and multi-agent PvP arenas, you create a natural, viral social loop (“My agent can beat yours”) that drives organic growth without expensive customer acquisition costs.
Crowdsourcing the Future of Offline Data : RL’s biggest bottleneck in industries like robotics is the lack of diverse, high-quality demonstration datasets. By building a B2C platform where users play to train agents, we are crowdsourcing a massive library of human behavioral trajectories across thousands of environments. This diverse dataset is a goldmine for training future foundational models that generalize across multiple control domains.
Here is a quick YouTube demonstration of the platform in action:
And here is a preview of the web UI running a client-side training loop:
Here is how we built the architecture to make in-browser RL viable, the UX challenges we encountered, and why we believe RL is prime for a B2C comeback.
The Architecture: WebGPU + Pyodide + Web Worker
To achieve a zero-install experience, we had to move both the environment simulation and the model training into the client browser.
Our core architecture splits the load into three parts:
Browser Context
├── Web Worker (Pyodide)
│ ├── gymnasium (micropip)
│ ├── env.step(action) -> observation, reward, terminated
│ └── Sends state updates to the main thread
│
├── Main Thread (WebGPU & JS)
│ ├── Model inference & policy updates
│ ├── Action selection (via WebGPU-accelerated tensors)
│ └── Drawing Commands Bridge (rendering to Canvas)
1. Pyodide in a Web Worker
We use Pyodide (Python compiled to WebAssembly) to run official gymnasium environments like CartPole and MountainCar .
Running Python environments in the browser can easily freeze the UI because Python execution is single-threaded. By offloading Pyodide to a dedicated Web Worker , we keep the main UI thread running at a smooth 60fps.
2. Main Thread for WebGPU Inference & Drawing Commands Bridge
We utilize WebGPU on the main thread for neural network updates and forward passes.
Because the observation state returned by Pyodide is simple (e.g., 4 float values for CartPole ), passing data between the Web Worker and the WebGPU context has negligible serialization overhead.
For environment visualization, standard Gymnasium environments rely heavily on Pygame, which cannot render directly to a browser Canvas and introduces massive performance overhead if we pass raw pixel frame buffers ( rgb_array ) back and forth from WASM.
To bypass this, we built a Pygame Mocking Bridge in Pyodide:
We intercept all pygame.draw calls (like rect , circle , line ) in the Python environment by injecting a mock pygame module dynamically.
The mock module translates Python drawing operations into a list of lightweight JSON drawing commands (e.g., {"type": "circle", "color": [255, 0, 0], "center": [100, 50], "radius": 10} ).
These instructions are sent from the Web Worker to the main thread, where native HTML5 Canvas draws them.
This compromise preserves complete compatibility with standard Gymnasium environment rendering codes ( import pygame inside the environment still works) while keeping the graphics pipeline lightning fast and hardware-accelerated.
Overcoming Browser Constraints
Building a fully client-side ML training loop comes with unique friction points. Here is how we addressed them:
1. The 15-Second Pyodide Cold Start
Loading Pyodide, initializing the Python runtime, and installing dependencies via micropip (like numpy and scipy ) can take anywhere from 10 to 30 seconds depending on the user’s connection.
To hide this latency:
Background Preloading : We spin up the Web Worker and start loading Pyodide in the background the moment a user lands on the platform, even while they are just browsing the catalog of environments.
Service Worker Caching : Once loaded, WASM binaries and pip packages are aggressively cached via IndexedDB/Cache API so subsequent visits load instantly.
2. Session Persistence (Handling Tab Closes)
Unlike server-side training, if a user closes their browser tab, they risk losing their active training progress.
Currently, we save basic training metadata (such as total episodes, training steps, and best reward) as well as local demonstration logs to the browser’s localStorage for progress tracking. To solve the problem of losing active model states, our immediate next step is to implement automatic agent checkpointing (saving model weights and hyperparameters) to IndexedDB , allowing users to seamlessly resume active training sessions even after a tab close or crash.
The Core Loop: Imitation Learning as a Gateway
A major roadblock in RL is sample inefficiency: an agent can take thousands of episodes just to learn how to balance a pole. For a casual user in a browser, waiting 10 minutes for any sign of intelligence is a retention killer.
We designed a dual-mode interaction loop to fix this:
Human Demonstration (Direct Control) : Users can directly play the game using their keyboard/mouse to generate demonstration trajectories.
Imitation Learning (Behavioral Cloning) : We use these human demonstrations to train a base model instantly. The agent immediately mirrors the player’s playstyle.
RL Fine-Tuning : Once the agent has a basic grasp of the task, the user hands over control to the RL algorithm (PPO, DQN) to optimize the policy beyond human performance.
This “human-in-the-loop” approach solves the cold-start problem of training from scratch and gives users a strong sense of ownership: “This agent is learning to play like me, but better.”
We don’t view Agenlus as just a browser game platform. Our ultimate goal is to build the foundational infrastructure for the next generation of physical control AI.
Our roadmap spans three major phases of expansion:
1. A Modular Sandbox for Physics and Control (Medium-Term)
We are building a low-code, web-based sandbox studio where creators can design custom 2D/3D environments with complex physics, import them with a single click, and define custom reward systems. By democratizing environment creation, we want to scale from a few classic control tasks to thousands of community-generated physical challenges.
2. Sim-to-Real: The Robotics Bridge (Long-Term)
Training robots in the physical world is slow, dangerous, and expensive. Simulators are the only viable path forward.
Because Agenlus agents are trained and exported in standard formats (like ONNX and TensorFlow.js models) directly from the browser, we are building a seamless deployment pipeline to physical hardware. In the future, you will be able to train an autonomous drone control policy or a robotic arm manipulator inside an Agenlus browser tab, and deploy that exact policy over-the-air (OTA) to a physic

[truncated]

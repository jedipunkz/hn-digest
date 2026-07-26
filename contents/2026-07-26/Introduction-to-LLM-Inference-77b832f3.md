---
source: "https://kraghavan.ca/llm-infrastructure/inference/2026/04/14/re-introduction-to-inference.html"
hn_url: "https://news.ycombinator.com/item?id=49054962"
title: "Introduction to LLM Inference"
article_title: "What Is LLM Inference, Really? A Deep Technical Walkthrough - Karthika Raghavan"
author: "kraghavan"
captured_at: "2026-07-26T05:22:15Z"
capture_tool: "hn-digest"
hn_id: 49054962
score: 1
comments: 0
posted_at: "2026-07-26T05:18:42Z"
tags:
  - hacker-news
  - translated
---

# Introduction to LLM Inference

- HN: [49054962](https://news.ycombinator.com/item?id=49054962)
- Source: [kraghavan.ca](https://kraghavan.ca/llm-infrastructure/inference/2026/04/14/re-introduction-to-inference.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T05:18:42Z

## Translation

タイトル: LLM 推論の概要
記事のタイトル: LLM 推論とは実際何ですか?詳しい技術ウォークスルー - Karthika Raghavan
説明: バイトからトークン、埋め込み、モデルが最終的に吐き出す単語への注意に至るまで、送信ボタンを押したときに実際に何が起こるかをエンジニアが注釈付きで説明します。スキップされた手順はありません。いいえ、「そして魔法が起こります。」

記事本文:
LLM 推論とは実際何ですか?詳しい技術ウォークスルー
バイトからトークン、埋め込み、モデルが最終的に吐き出す単語への注意に至るまで、送信ボタンを押したときに実際に何が起こるかをエンジニアが注釈付きで説明します。スキップされた手順はありません。いいえ、「そして魔法が起こります。」
ソフトウェア エンジニア • 分散システム • LLM インフラストラクチャ
2. アーティファクト: 10 GB のダウンロードには実際には何が含まれているのでしょうか? 「マニュアル」（ヘッダー）
量子化: 脳の縮小
3. 3 つのフェーズ: 領土の前の地図
4. トークン化: テキストを数値に分割する トークン化は CPU に依存しますか?
5. プレフィル: モデルはプロンプトを読み取ります 埋め込み行列
ここでモデルの重みはどのように役立ちますか?
6. 位置埋め込み: 順序についてモデルに教える 順序はどのように計算されますか?
7. トランス層: 実際の作業が行われる場所
8. デコード: 一度に 1 つのトークンを永久にデコード ステップ 1: 「オン」を予測する
デコード ステップ 2: 「the」を予測する
サンプリングステップ (創造性が息づく場所)
9. メモリがデコードのボトルネックになる理由
10. KV キャッシュ: 推論プレフィックス キャッシュで最も重要なデータ構造: 無料で高速化
11. 注意: 自己注意を機能させるメカニズム
ページング アテンション: KV キャッシュ用の仮想メモリ
12. 連続バッチ処理: スループットのロック解除
13. 重要な指標 TTFT — 最初のトークンまでの時間
ITL — トークン間遅延 (別名 TPOT)
トラフィック形状別のスケーリング戦略
14. 時間は実際どこへ行くの?
15. オラマを知る価値のある推論エンジン
TGI (テキスト生成推論)
16. すべての背後にある言語
17. CPU とメモリの集中的な意味の概要
18. SRE が実際に知っておくべきこと
概要: 生産のためのメンタルモデル 「作戦室」参照表
最終原則 (あなたのTL;DR)
結論: フィールドは M です

ずっと、基本はそうではない
正直に言わせてください。 LLM インフラストラクチャに取り組み始めたとき、私には 11 年間の分散システムの経験がありました。 Kafka、Kubernetes、Prometheus は知っていました。寝ている間にパーティションの再バランスをデバッグできました。
それでも、誰かが推論中に実際に何が起こるのかと初めて尋ねたとき、私は「モデルはプロンプトを読み取り、トークンを生成します」のようなことを言いました。これは、「データベースがクエリを読み取って行を返す」が技術的に正しいのと同じで、正確ではありますが、役に立たず、主任エンジニアの給料をもらっている人にとっては非常に恥ずかしいことです。
この投稿は、初日にあればよかったと思っているものです。実際の例、パフォーマンスがどこに行くのかについての正直な説明、本番環境の問題について実際に推論できる十分な詳細を使用して、リクエストが到着した瞬間から画面にテキストが表示される瞬間まで、推論パイプライン全体を見ていきます。
いいえ、「その後、変圧器がその役割を果たします。」スキップされた手順はありません。ストラップで固定します。
トレーニングでは、大規模なデータセットを取得し、それをモデルで何百万回も実行し、モデルが次の単語を予測できるようになるまで数十億の数値重みをゆっくりと調整します。トレーニングは 1 回 (または時々) 行われます。 GPU 時間に数百万ドルの費用がかかり、研究者チームが必要です。
推論は、誰かがモデルを使用するたびに、その後に行われることです。これは、学習した重みを使用して新しい入力に応答するモデルです。重みは変わりません。学習は起こりません。これは純粋なフォワードパス計算です。
次のように考えてください。トレーニングとはパンを焼くことです。それをスライスしてお客様に提供するのが推論です。パン（重り）が完成しました。キッチン (推論エンジン) は、行列が店の前に戻らないように、十分な速さで皿に盛るだけで済みます。

など。
推論エンジンは、凍結されたモデルの重みを取得し、入力に対してそれらを実行するランタイムです。同じ重みを Ollama、vLLM、TensorRT-LLM、または TGI 上で実行でき、それぞれから大幅に異なるパフォーマンスが得られます。重みは変わりません。実行戦略はそうなります。
この区別は運用上重要です。推論は解決された問題ではありません。大規模なモデルを効率的に提供することは、完全なエンジニアリング分野です。
2. アーティファクト: 10 GB のダウンロードには実際には何が含まれているのでしょうか?
ollama pull misstral を実行したり、HuggingFace からモデルを取得したりするとき、単に「プログラム」をダウンロードしているわけではありません。あなたは箱の中に冷凍された巨大な脳をダウンロードしているのです。 「チャットするだけ」のモデルが SSD の 10 GB を占有する理由を疑問に思ったことがあるなら、それはモデルがトレーニング段階で学習した何十億もの小さな数値「設定」が詰め込まれているからです。
GGUF (または Safetensors ) ファイルを、Ikea の巨大なフラットパック ボックスと考えてください。実用的なモデルを構築するには、取扱説明書とハードウェアの 2 つが必要です。
7B パラメータ モデル ファイルの内容:
GGUF ファイル構造 (簡略化):
§── ヘッダー
│ §── モデルアーキテクチャ (LlamaForCausalLM)
│ §── 語彙 (32000 トークン + その埋め込み)
│ §── コンテキスト長 (4096、8192 など)
│ └── ハイパーパラメータ (n_layers、n_heads など)
│
└── 重みテンソル:
§── token_embeddings [32000 × 4096] ← 埋め込み行列
§──layer.0.attention.q [4096 × 4096] ← 射影重みのクエリ
§──layer.0.attention.k [4096 × 4096] ← キー投影ウェイト
§──layer.0.attention.v [4096 × 4096] ← 値投影重み
§──layer.0.attention.out [4096 × 4096] ← 出力投影
§──layer.0.ffn.up [4096 × 11008] ← フィードフォワードアップ
§─

─layer.0.ffn.down [11008 × 4096] ← フィードフォワードダウン
§── … × 32層
└── Output_norm + lm_head [32000 × 4096] ← ロジットへの最終投影
「マニュアル」（ヘッダー）
これはファイルの最初の数キロバイトです。これは、推論エンジン (Ollama や vLLM など) にどのように脳を組み立てるかを指示します。これには次のものが含まれます。
アーキテクチャ : モデル タイプ (例: LlamaForCausalLM ) を識別し、エンジンがどの数学ルールを適用するかを認識します。
語彙 : およそ 32,000 ～ 128,000 の「トークン」(モデルが話す音節) の辞書。
ハイパーパラメータ : レイヤーの数 (32 または 80) やコンテキストの長さ (記憶できる量) などの重要な設定。
ファイルの残りの部分は、 Weights と呼ばれる数値の行や行だけです。すべての推論リクエストは基本的に、これらの行列から値を検索し、それらを 32 回以上乗算します。
量子化: 脳の縮小
同じモデルでも、一部のファイルは 15 GB ですが、他のファイルは 4 GB であることに気づくかもしれません。これが量子化、つまり圧縮技術です。高精度の 16 ビット浮動小数点を低精度の整数 (4 ビットなど) に変換します。
SRE が INT4 を好む理由: 精度が低い = テンソルが小さい = メモリ転送が速い。デコードはメモリに依存するため、メモリ バスの負荷がそれほど大きくないため、INT4 モデルは「フル」バージョンよりも 20 ～ 40% 優れた TPOT (1 秒あたりのトークン数) を実現します。
要点: コードを実行しているわけではありません。大量の数学的負荷の高いルックアップ テーブルをロードしていることになります。 GGUF は単一ファイルの「ボックス」であり、量子化はそのボックスを小型のトラック (GPU) に収める方法です。
3. 3 つのフェーズ: 領土の前の地図
すべての推論リクエストは、大きく 3 つのフェーズを経ます。これらは同じように高価ではなく、同じように並列化可能ではなく、また同じように p99 レイテンシーに適しているわけでもありません。

┌─────────────────────────┐
│ TOKENIZATION │ プレフィル（即時処理） │ デコード │
│ (CPU) │ (GPU、コンピューティング) │ ループ(GPU、メモリ) │
│ │ │ │
│ │ │ 現在 → │
│ テキスト → ID │ 埋め込み → 位置 → アテンション │ 次のトークン │
━━━━━━━━━━━━━━━━━━━━━━━━┘
速い プロンプトの長さでスケールします 遅い
トークン化 : テキストをモデルが理解できるトークン ID に分割します。 CPU バウンド。速い。
Prefill : モデルを通じてプロンプト全体を処理します。 GPU コンピューティングに依存します。プロンプトの長さに応じてスケールします。
デコード : 出力トークンを一度に 1 つずつ生成します。 GPU メモリに依存します。完了するまでループで実行されます。
各フェーズには独自のボトルネックがあります。一つずつ見ていきましょう。
4. トークン化: テキストを数値に分割する
単一の GPU 操作が行われる前に、テキストをモデルが使用できる形式 (トークン ID と呼ばれる一連の整数) に変換する必要があります。
トークンは文字でも単語でもありません。これは、トレーニング コーパス内に頻繁に出現するテキストの塊であり、独自の ID が必要です。この語彙を構築するには、WordPiece (BERT で使用)、Unigram (SentencePiece で使用) などのいくつかの方法がありますが、最新の LLM で主流のアプローチは、バイト ペア エンコーディング (BPE) です。これは、最も一般的な文字またはサブワードのペアを、目標の語彙サイズに達するまで繰り返し単一のトークンにマージする圧縮アルゴリズムです。
結果として、およそ 32,000 ～ 128,000 のトークンの語彙が得られます。

対応する整数 ID。モデルはテキストを認識することはありません。数値のリストを認識します。
プロンプトの例を見てみましょう: 「猫は座っていました」
トークン化後は次のようになります。
「ザ」→1026
「猫」→5992
「座った」→3290
トークンID: [1026、5992、3290]
「cat」と「sat」の前のスペースに注意してください。これはトークンの一部です。トークナイザーは空白を重視します。空白は意味と頻度に影響を与えるためです。
はい。トークナイザーは通常、Rust (HuggingFace のトークナイザー クレート) または C++ で書かれますが、これはまさにこの理由からです。ほとんどのリクエストでは、目に見えないほど高速です。短いプロンプトの場合はマイクロ秒です。
問題となるのは、非常に長いドキュメントがバッチ処理ジョブに供給されることです。 100,000 トークンのコンテキストでは、100,000 のトークン ルックアップを処理する必要があります。 GPU の処理に比べればまだ高速ですが、CPU 上で実行されるパイプラインの 1 つのステップであり、単純に GPU を追加することはできません。
改善方法: バッチ ワークロードの CPU コア間でトークン化を並列化します。あるいは、これが本当の解決策です。同じコンテンツを繰り返し再トークン化しないことです。すべてのリクエストに送信する共有システム プロンプトがある場合、プロンプトを 1 回トークン化し、結果をキャッシュすると、待ち時間は発生しません。
5. 事前入力: モデルはプロンプトを読み取ります
これでトークン ID が得られました。モデルはこれらの ID を推論できるものに変換する必要があります。これはプレフィルです。つまり、プロンプト全体を 1 回のショットで処理するモデルです。
プレフィルには、混同しやすい 2 つのサブステップがあります。それは、埋め込みルックアップと実際のトランスフォーマ フォワード パスです。順番に見ていきましょう。
すべてのトークン ID は、 embedding と呼ばれる浮動小数点数の高次元ベクトルにマップされます。これらのベクトルは、モデルの埋め込み行列 (語彙トークンごとに 1 行、埋め込み次元ごとに 1 列を持つ巨大なルックアップ テーブル) 内に存在します。
32,000 トークンのモデルの場合

ocabulary および 4,096 の埋め込み次元では、この行列の形状は [32000, 4096] です。 16 ビット浮動小数点精度では、埋め込み層だけで約 256MB になります。
例 [1026, 5992, 3290] は次のようになります。
トークン ID 1026 → 埋め込み行 1026 → [0.12, -0.43, 0.81, ..., 0.07] (4096 値)
トークン ID 5992 → 埋め込み行 5992 → [-0.34, 0.91, 0.12, ..., -0.22] (4096 値)
トークン ID 3290 → 埋め込み行 3290 → [0.67, 0.05, -0.88, ..., 0.44] (4096 値)
ここではページに収まるように 8 次元に単純化しています。実際には、モデルに応じて 4,096 または 8,192 次元になります。
形状を示すために簡略化 (4096D ではなく 3D):
「ザ」 → [0.12, -0.43, 0.81]
「猫」→[-0.34、0.91、0.12]
「土」 → [0.67, 0.05, -0.88]
形状: [3 トークン × 3 次元] = float の行列
これらのベクトルはランダムではありません。これらはトレーニングの結果です。モデルは、この空間では「猫」と「犬」が近くに住んでおり、「猫」と「量子力学」は遠く離れていることを学習しました。幾何学模様は意味論的な意味をエンコードします。
ここでモデルの重みはどのように役立ちますか?
埋め込み行列は、具体的にはモデルの重みです。ダウンロードした 10 GB (または 40 GB、または 70 GB) ファイル (GGUF またはセーフテンサー ファイル) には、モデルがトレーニング中に学習したすべての重み行列が含まれています。埋め込みルックアップは文字通り、行番号によってこれらの重み行列の 1 つにインデックスを付けます。
推論を実行するとき、創造的なことは何も計算されていません。何百万回ものトレーニング反復にわたって調整された凍結された数値に対して行列計算を行っています。
6. 位置埋め込み: 順序についてモデルに教える
ここに問題があります。埋め込みルックアップはテーブル ルックアップです。 「cat」がトークン 2 で、「sat」がトークン 3 であることは関係ありません。異なる順序で同じトークンを持つ 2 つのリクエストは ID を生成します。

エンティカルな埋め込み。
しかし、順序は非常に重要です。 「猫が犬の上に座った」と「犬が猫の上に座った」

[切り捨てられた]

## Original Extract

An Engineer’s annotated tour through what actually happens when you hit send — from bytes to tokens to embeddings to attention to the word your model finally spits out. No skipped steps. No “and then magic happens.”

What Is LLM Inference, Really? A Deep Technical Walkthrough
An Engineer’s annotated tour through what actually happens when you hit send — from bytes to tokens to embeddings to attention to the word your model finally spits out. No skipped steps. No “and then magic happens.”
Software Engineer • Distributed Systems • LLM Infrastructure
2. The Artifact: What’s Actually in That 10GB Download? The “Manual” (The Header)
Quantization: Shrinking the Brain
3. The Three Phases: A Map Before the Territory
4. Tokenization: Chopping Text Into Numbers Is Tokenization CPU-Bound?
5. Prefill: The Model Reads Your Prompt The Embedding Matrix
How Do the Model Weights Help Here?
6. Positional Embeddings: Teaching the Model About Order How Is It Calculated?
7. The Transformer Layers: Where the Real Work Happens
8. Decoding: One Token at a Time, Forever Decode Step 1: Predicting “on”
Decode Step 2: Predicting “the”
The Sampling Step (Where Creativity Lives)
9. Why Memory Is the Decode Bottleneck
10. KV Cache: The Most Important Data Structure in Inference Prefix Caching: Free Speedups
11. Attention: The Mechanism That Makes It Work Self-Attention in Plain English
Paged Attention: Virtual Memory for KV Cache
12. Continuous Batching: The Throughput Unlock
13. The Metrics That Matter TTFT — Time to First Token
ITL — Inter-Token Latency (aka TPOT)
Scaling Strategy by Traffic Shape
14. Where Does the Time Actually Go?
15. The Inference Engines Worth Knowing Ollama
TGI (Text Generation Inference)
16. The Languages Behind It All
17. What CPU vs. Memory Intensive Means, Summarized
18. What an SRE Actually Needs to Know
The Summary: A Mental Model for Production The “War Room” Reference Table
The Final Principles (Your TL;DR)
Conclusion: The Field Is Moving, The Fundamentals Aren’t
Let me be honest with you. When I started working on LLM infrastructure, I had eleven years of distributed systems experience. I knew Kafka, Kubernetes, Prometheus. I could debug a partition rebalance in my sleep.
And yet the first time someone asked me what actually happens during inference , I said something like “the model reads the prompt and generates tokens.” Which is technically true the same way “a database reads your query and returns rows” is technically true — accurate, useless, and deeply embarrassing for someone drawing a principal engineer’s salary.
This post is what I wish I’d had on day one. We’re going to walk through the entire inference pipeline — from the moment your request arrives to the moment you see text on screen — with real examples, honest explanations of where the performance goes, and enough detail that you can actually reason about production problems.
No “and then the transformer does its thing.” No skipped steps. Strap in.
Training is where you take a massive dataset, run it through a model millions of times, and slowly adjust billions of numerical weights until the model gets good at predicting the next word. Training is done once (or occasionally). It costs millions of dollars in GPU-hours and requires a team of researchers.
Inference is what happens afterward, every time someone uses the model. It’s the model using those learned weights to respond to new input. No weights change. No learning happens. It’s pure forward-pass computation.
Think of it like this: training is baking the bread. Inference is slicing it and serving it to customers. The bread (weights) is done. The kitchen (inference engine) just has to plate it fast enough that the queue doesn’t back up to the street.
The inference engine is the runtime that takes the frozen model weights and executes them against an input. The same weights can run on Ollama, vLLM, TensorRT-LLM, or TGI — and get meaningfully different performance from each. The weights don’t change. The execution strategy does.
This distinction matters operationally: inference is not a solved problem . Serving a model efficiently at scale is a full engineering discipline.
2. The Artifact: What’s Actually in That 10GB Download?
When you run ollama pull mistral or grab a model from HuggingFace, you aren’t just downloading a “program.” You’re downloading a massive, frozen brain in a box. If you’ve ever wondered why a model that “just chats” takes up 10GB of your SSD, it’s because it is packed with billions of tiny numerical “preferences” the model learned during its training phase.
Think of the GGUF (or Safetensors ) file as a giant Ikea flat-pack box. To build the working model, you need two things: the Instruction Manual and the Hardware .
What’s inside a 7B parameter model file:
GGUF file structure (simplified):
├── Header
│ ├── Model architecture (LlamaForCausalLM)
│ ├── Vocabulary (32000 tokens + their embeddings)
│ ├── Context length (4096, 8192, etc.)
│ └── Hyperparameters (n_layers, n_heads, etc.)
│
└── Weight tensors:
├── token_embeddings [32000 × 4096] ← the embedding matrix
├── layer.0.attention.q [4096 × 4096] ← Query projection weights
├── layer.0.attention.k [4096 × 4096] ← Key projection weights
├── layer.0.attention.v [4096 × 4096] ← Value projection weights
├── layer.0.attention.out [4096 × 4096] ← Output projection
├── layer.0.ffn.up [4096 × 11008] ← Feed-forward up
├── layer.0.ffn.down [11008 × 4096] ← Feed-forward down
├── ... × 32 layers
└── output_norm + lm_head [32000 × 4096] ← Final projection to logits
The “Manual” (The Header)
This is the first few kilobytes of the file. It tells the inference engine (like Ollama or vLLM) how to put the brain together. It includes:
The Architecture : Identifies the model type (e.g., LlamaForCausalLM ) so the engine knows which math rules to apply.
The Vocabulary : A dictionary of roughly 32,000 to 128,000 “tokens” (the syllables the model speaks).
The Hyperparameters : Crucial settings like the number of layers (32 or 80) and the context length (how much it can remember).
The rest of that file is just rows and rows of numbers called Weights . Every inference request is essentially looking up values from these matrices and multiplying them together—32 times over.
Quantization: Shrinking the Brain
You might notice some files are 15GB while others are 4GB for the same model. This is Quantization —the art of compression. We turn high-precision 16-bit floats into lower-precision integers (like 4-bit).
Why SREs love INT4: Lower precision = smaller tensors = faster memory transfers. Because decoding is memory-bound, an INT4 model often delivers 20-40% better TPOT (tokens per second) than the “full” version because the memory bus isn’t screaming as loud.
The takeaway: You aren’t executing code; you are loading a massive, math-heavy lookup table. GGUF is your single-file “box,” and quantization is how you fit that box into a smaller truck (your GPU).
3. The Three Phases: A Map Before the Territory
Every inference request goes through three broad phases. They are not equally expensive, not equally parallelizable, and not equally friendly to your p99 latency.
┌──────────────────────────────────────────────────────────────────┐
│ TOKENIZATION │ Prefill (Prompt Processing) │ Decode │
│ (CPU) │ (GPU, compute) │ Loop(GPU, Mem) │
│ │ │ │
│ │ │ current → │
│ Text → IDs │ Embed → Position → Attention │ next token │
└──────────────────────────────────────────────────────────────────┘
Fast Scales with prompt length Slow
Tokenization : Split the text into token IDs the model understands. CPU-bound. Fast.
Prefill : Process the entire prompt through the model. GPU compute-bound. Scales with prompt length.
Decode : Generate output tokens one at a time. GPU memory-bound. Runs in a loop until done.
Each phase has its own bottleneck. Let’s go through them one by one.
4. Tokenization: Chopping Text Into Numbers
Before a single GPU operation happens, your text has to be converted into a format the model can work with: a sequence of integers called token IDs.
A token is not a character, and it’s not a word. It’s a chunk of text that appears frequently enough in the training corpus to deserve its own ID. There are several ways to build this vocabulary — WordPiece (used by BERT), Unigram (used by SentencePiece), and others — but the dominant approach in modern LLMs is Byte Pair Encoding (BPE): a compression algorithm that iteratively merges the most common pairs of characters or subwords into single tokens until it reaches a target vocabulary size.
The result is a vocabulary of roughly 32,000–128,000 tokens, each with a corresponding integer ID. The model never sees your text — it sees a list of numbers.
Take our example prompt: "The cat sat"
After tokenization, this becomes something like:
"The" → 1026
" cat" → 5992
" sat" → 3290
Token IDs: [1026, 5992, 3290]
Note the space before “cat” and “sat” — it’s part of the token. Tokenizers care about whitespace because it affects meaning and frequency.
Yes. The tokenizer is usually written in Rust (HuggingFace’s tokenizers crate) or C++ for exactly this reason. For most requests it’s fast enough to be invisible — microseconds for a short prompt.
Where it bites you: very long documents fed to batch processing jobs. A 100,000-token context requires processing 100,000 token lookups. It’s still fast relative to GPU work, but it’s the one step in the pipeline running on CPU that you can’t just throw more GPU at.
How it’s improved: Parallelizing tokenization across CPU cores for batch workloads. Or — and this is the real fix — not re-tokenizing the same content repeatedly . If you have a shared system prompt you send to every request, tokenizing it once and caching the result is free latency.
5. Prefill: The Model Reads Your Prompt
Now we have token IDs. The model needs to turn those IDs into something it can reason about. This is prefill — the model processing the entire prompt in one shot.
Prefill has two sub-steps that are easy to conflate: embedding lookup and the actual transformer forward pass . Let’s take them in order.
Every token ID maps to a high-dimensional vector of floating-point numbers called an embedding . These vectors live in the model’s embedding matrix — a giant lookup table with one row per vocabulary token and one column per embedding dimension.
For a model with a 32,000-token vocabulary and 4,096 embedding dimensions, this matrix has shape [32000, 4096] . At 16-bit float precision, that’s about 256MB just for the embedding layer.
Our example [1026, 5992, 3290] becomes:
Token ID 1026 → embedding row 1026 → [0.12, -0.43, 0.81, ..., 0.07] (4096 values)
Token ID 5992 → embedding row 5992 → [-0.34, 0.91, 0.12, ..., -0.22] (4096 values)
Token ID 3290 → embedding row 3290 → [0.67, 0.05, -0.88, ..., 0.44] (4096 values)
I’m simplifying to 8 dimensions here so this fits on a page. In reality it’s 4,096 or 8,192 dimensions depending on the model.
Simplified (3D instead of 4096D), just to show the shape:
"The" → [0.12, -0.43, 0.81]
" cat" → [-0.34, 0.91, 0.12]
" sat" → [0.67, 0.05, -0.88]
Shape: [3 tokens × 3 dims] = a matrix of floats
These vectors aren’t random. They’re the result of training — the model has learned that “cat” and “dog” live close together in this space, and “cat” and “quantum mechanics” are far apart. The geometry encodes semantic meaning.
How Do the Model Weights Help Here?
The embedding matrix IS the model weights, specifically. The 10GB (or 40GB, or 70GB) file you download — the GGUF or safetensors file — contains all the weight matrices the model learned during training. The embedding lookup is literally indexing into one of those weight matrices by row number.
When you run inference, you’re not computing anything creative. You’re doing matrix math against frozen numbers that were tuned over millions of training iterations.
6. Positional Embeddings: Teaching the Model About Order
Here’s a problem: the embedding lookup is a table lookup. It doesn’t care that “cat” is token 2 and “sat” is token 3. Two requests with the same tokens in different orders would produce identical embeddings.
But order matters enormously. “The cat sat on the dog” and “The dog sat on the cat

[truncated]

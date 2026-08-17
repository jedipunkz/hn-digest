---
source: "https://www.lancedb.com/blog/data-loading-guide"
hn_url: "https://news.ycombinator.com/item?id=49336985"
title: "Data Loading for AI/ML: A Comprehensive Guide"
article_title: "Data Loading for AI/ML: A Comprehensive Guide"
image: "https://cdn.prod.website-files.com/69b2da72cae7eea2b0091d5f/6a7a382205f14767ae1150a4_preview-image.png"
author: "westonpace"
captured_at: "2026-08-17T21:18:12Z"
capture_tool: "hn-digest"
hn_id: 49336985
score: 5
comments: 0
posted_at: "2026-08-17T20:17:24Z"
tags:
  - hacker-news
  - translated
---

# Data Loading for AI/ML: A Comprehensive Guide

- HN: [49336985](https://news.ycombinator.com/item?id=49336985)
- Source: [www.lancedb.com](https://www.lancedb.com/blog/data-loading-guide)
- Score: 5
- Comments: 0
- Posted: 2026-08-17T20:17:24Z

## Translation

タイトル: AI/ML のデータ読み込み: 包括的なガイド
説明: モデル トレーニングのためのデータ読み込み - 3 つのパイプライン ステージ、並列処理戦略、シャッフル、キャッシュ、再開可能性、および LanceDB の StreamingDataset がどのように適合するかについて詳しく説明します。

記事本文:
AI/ML のデータ読み込み: 包括的なガイド
Reverie • AI ビルダーのためのサミット | 11月5日、サンフランシスコ |参加申し込み
使用例
キュレーション 最適な分布の検索、大規模なデータセットの重複排除、エッジ ケースの表面化、すべてを 1 か所で実行 機能エンジニアリング Python UDF による機能の構築と拡張、自動更新、テーブルの書き換えなし 検索と取得 SQL フィルターを使用した統合ベクトル、フルテキスト、およびハイブリッド検索で本番環境に対応した検索を実現 トレーニング 最大 70% の MFU で、データ移動のボトルネックなしで厳選されたデータから直接トレーニング 本番パートナーのドキュメント ブログ コミュニティ
お問い合わせ
ありがとうございます！あなたの提出物は受理されました！おっと！フォームの送信中に問題が発生しました。ブログ投稿内の検索 マルチモーダル データの機能エンジニアリング: ラップトップから LanceDB を使用したクラスターまで これは div ブロック内のテキストです。 Geneva-example を使用してマルチモーダル特徴量エンジニアリング パイプラインを 10 分で実行し、LanceDB の特徴量エンジニアリング API を使用してその CLIP 埋め込みステップを再構築します。これは、LanceDB Enterprise で独自のテーブルを強化するために使用するのと同じ呼び出しです。 Reverie Summit の発表: AI の次のブレークスルーは何に基づいて行われるか これは div ブロック内のテキストです。生成ビデオ、世界モデル、物理 AI、およびその下にあるデータ システムを構築する研究者やエンジニアとのある日。 ⚡ リファインなしのマルチビット RaBitQ、🌋 Bytedance の Lance ベースの AI スタック、🤖 埋め込まれた AI データ用の Lance これは div ブロック内のテキストです。マルチビット RaBitQ はリファインのオーバーヘッドなしで 96% の再現率を達成し、Volcano Engine はトレーニング パイプラインを 7 日から 1 日に短縮し、Lance は 100 Hz のロボティクス データ インジェストを強化し、さらに今後のイベント、エンタープライズ アップデート、コミュニティ アップデートを強化します。 CrewAI が LanceDB 上のエージェント メモリを再構築し、20 億人以上のエージェントの実行を強化した理由

div ブロック内のテキスト。 CrewAI は、2 システムのメモリ スタックを LanceDB に置き換え、エージェントのメモリを再構築して矛盾を解決し、信頼度に基づいてリコールをゲートしました。デフォルトで月間 1,200 万ダウンロードに出荷されるようになりました
[切り捨てられた]
必須ではない Cookie はデフォルトで無効になっています。
このバナーを閉じても、選択は確定されません。
詳細については、当社のプライバシー ポリシーをご覧ください。
LanceDB は、Cookie および同様のテクノロジーを使用してエクスペリエンスを向上させ、トラフィックを分析し、関連するコンテンツや広告を表示します。私の個人情報を販売または共有しないでください 私の個人情報を販売または共有しないでください 同意 詳細については、プライバシー ポリシーを参照してください。キャンセル 設定を保存
ありがとうございます！あなたの提出物は受理されました！おっと！フォームの送信中に問題が発生しました。 LanceDB は、Cookie および同様のテクノロジーを使用してエクスペリエンスを向上させ、トラフィックを分析し、関連するコンテンツや広告を表示します。受け入れる
AI/ML のデータ読み込み: 包括的なガイド
これはタイトルです これはサブタイトルです 機械学習タスクにおける「データ読み込み」は、データを何らかのアルゴリズムに移動するプロセスです。通常、ある種のモデルをトレーニングしますが、さまざまなバリエーションが存在します。非常に多くの例があるため、私は常に、過度に具体的な例と過度に広範な一般的な例の間でバランスをとっていることに気づきました。
I/Oステージ
CPUステージ
GPUステージ
ストレージ
データベース
ホストマシン
CPU・RAM
GPU
モデル
データ読み込みを説明する際に私が直面する課題は、それが非常に単純である (データセットを反復する) と同時に、ニュアンスに満ちている (このガイド全体) ということです。データの読み込みはデータの世界とモデルの世界の間の架け橋でもあり、用語の選択だけが原因でトピックが混乱することがよくあります。この記事では、データ読み込みプロセスの包括的な概要をデータの観点から説明しようとします。

エンジニア。概念を説明し、課題を説明し、一般的なパフォーマンスの落とし穴について説明します。この記事は pytorch と lancedb に焦点を当てますが、この情報は他のライブラリにも当てはまります。
まず、非常に具体的な例を考えてみましょう。Alpaca データセットを使用して Qwen2.5-0.5B-Instruct という名前のモデルに適用される教師ありファイン チューニング (SFT) です。私たちが行っていることは、大規模な汎用モデルを取得し、より厳密に調整されたタスクを実行するように微調整することです。モデルの重みは最大 1 GB です。モデルはすでに微調整されています (そのため「-Instruct」接尾辞が付いています) が、モデルは小さくてすぐに利用できるため、いつでもさらに微調整できることに注意してください。 Alpaca データセットは、約 25 MB のデータと 50,000 行で構成される命令と応答のペアの優れたオープン データセットです。
クウェン 2.5
一般的な理解
アルパカ
50K 命令ペア
+
微調整
ファインチューニングモデル
指示応答
最も単純なデータローダー
これは複雑に聞こえるかもしれませんが、そうではありません。コードで見てみましょう。
model =load_model( "Qwen2.5-0.5B-Instruct" )
tbl = open_table( "アルパカ" )
tbl 内のバッチの場合:
トレイン(モデル、バッチ)
変換 = 変換(バッチ) 必要なのはこれだけです。いくつかのデータを取得し、それを少し変換し、モデルにそれを確認させれば完了です。残念ながら、これから説明するように、さらに細かいニュアンスがありますが、データの読み込みとは何かという基本的な前提を忘れてはなりません。
モデルは従来、GPU でトレーニングされてきました。データは従来、ある種のオブジェクト ストレージまたはディスクに保存されていました。従来、ある程度の「データ変換」が CPU 上で発生します。これにより、標準データ ローダーに 3 つのステージができます。
I/O ステージは、データをストレージから (通常は) CPU にロードするときです。ここで使用されるリソースは、CPU/GPU の機能とは異なります。もし私たちがそうであれば

クラウド ストレージからロードする場合、おそらくサーバーに接続されている NIC が考慮されます。スループットを最大化するために、数百もの同時 HTTP リクエストが発生する可能性があります。ディスクからロードする場合は、接続されているディスクの種類を考慮するため、通常はそれほど同時実行性は必要ありません。
クラウドストレージ
CPU
HTTP
NVMe SSD
CPU
io_uring
インフィニバンド
GPU
GPUダイレクト
I/O ステージのさまざまな例
通常、これらのリソースは GB/秒で評価されます。容量は、1 GB/秒のかなり安定したベースラインから、TB/秒を実現できる実験的なストレージ システムまで多岐にわたります。 IOPS/秒も考慮する必要があります。小規模な読み取りを実行している場合は、帯域幅が制限されるのではなく、IOPS が制限される可能性があります。適切なベースラインは、クラウド ストレージの場合は 1000 ～ 5000 IOPS/秒、高性能 NVMe サーバーの場合は数千万 IOPS/秒になります。
データ エンジニアとして、この事実は少々驚くべきことです。私たちは、世の中のほぼすべてのデータ集約型アプリケーションにおいて、I/O が制限要因であることに慣れています。ただし、LLM は...そうですね...大きいです。 LLM を通過する小さな (<1KB) トークンごとに、数十億回の計算を実行する必要があります。これは、従来の OLAP タスクよりもはるかに多くの計算を必要とします。
次の段階では、GPU 用のデータを準備します。この段階の詳細は非常に柔軟です。最初の部分は通常、ストレージから受信したバイトを解凍してデコードします。これはデータ エンジニアとしてよく知られています。ただし、次は「正規化」、「変換」、「トークン化」などを行う可能性があります。正直に言うと、それはあまり重要ではありません。この部分について説明されると、私は通常、ただ笑ってうなずくだけです。
エンコードされた画像
CPUデコード
GPUステージ
エンコードされた画像
JPEG・バイナリ
ボトルネック
CPUデコード
JPEG画像のデコード
アイドル
GPU
アイドル・待機中
最適化が不十分な CPU ステージ (例: 1 人のワーカーによる JPEG デコード)

簡単に
ボトルネックになる。
適切なベースラインは、おそらくコアあたり約 1GB/秒です。非常に単純なデコードでは、コアあたり数十または数百 GB/秒を処理できます。非常に高価なデコード (JPEG 画像のデコードなど) では、これらの速度を大幅に下回る可能性があります。ここでの並列処理は通常、利用可能なコアの数に基づいています。これはインスタンスごと、クラウドごとに異なる場合があります。
CPU ステージを完全にスキップしてみることもできます。データを事前にトークン化し、変換済みのデータを保存することで、実際のトレーニング タスクに必要な CPU 作業量を削減できます。また、作業を GPU にプッシュして、GPU がデータをデコードして変換できるようにすることもできます。ただし、GPU がおそらくボトルネックであるため、GPU にさらに作業を加えるのは最も避けたいことです。
最後に、実際にモデルをトレーニングしたり、推論を実行したりする準備が整いました。 GPU の機能は非常に柔軟であるため、GPU を評価することは非常に困難です。最も一般的なスコアは「1 秒あたりのトークン数」であり、これはモデルに依存するスコアになります。したがって、Qwen の 1 秒あたりのトークンは、DeepSeek の 1 秒あたりのトークンとは異なります。これにより、パイプライン全体のパフォーマンスを仕様化することが難しくなる可能性がありますが、これについては次のセクションで詳しく説明します。
I/Oステージ
CPUステージ
GPUステージ
I/O
CPU
GPU
パイプラインのスループットは最も狭いステージによって制限され、容量を追加します。
ボトルネックを解決するまでは、ステージを広くしても役に立ちません。
基礎モデルのこの時代では、モデル自体が非常に大きくなります。バイトあたりに必要な計算量は膨大です。 GPU がボトルネックになると予想する必要があります。通常、データ読み込みの目標は、GPU がアイドル状態にならない程度の速度で GPU にデータを供給することです。
上記の 3 段階モデル​​は最も一般的なシナリオですが、唯一のシナリオではありません。例えば

複数のデータ ソースからデータをロードしてそれらをつなぎ合わせたり、追加の通信を導入して CPU ステージでリモート処理を実行したりすることはそれほど珍しいことではありません。ステージの境界は、リソースや並列処理が変更される場所にすぎません。この記事の残りの大部分は、単純な 3 ステージのソリューションに焦点を当てます (ただし、4 ステージ モデルで I/O を 2 ステージに分割するシャッフルに関するセクションが 1 つあります) が、ツールとテクニックは他のシナリオにも適用されます。
パフォーマンスの最終的な目標は、モデルのトレーニングをできるだけ高速に実行することです。データ ローダーの場合、これは通常、GPU を常にビジー状態に保つのに十分な速度でデータをロードする必要があることを意味します。正確な答えは常に状況によって異なりますが、2026 年の時点での経験則は「それほど速くない」です。
これをさらに定量化したいと思っていますが、ハードウェアの機能がわかっていたとしても、圧縮バイト (I/O スループット) からメモリ内バイト (CPU スループット)、そしてトークン (GPU スループット) に変換するのは難しい場合があります。具体的な例を見てみましょう。はじめに、Alpaca データセットを使用して Qwen を微調整する例を示しました。このテスト マシンで実行してみましょう。
CPU：16コア（Intel Xeon、2.00GHz）
ディスク: ローカル NVMe (仮想ディスク)
トレーニングを実行できる最速の速度は 1 秒あたり 16,000 トークンです。では、I/O ステージと CPU ステージはどれくらいの速度が必要なのでしょうか?言い換えれば、1 秒あたり 16,000 のトークンを満たすには、1 秒あたり何バイト必要でしょうか?
‍
圧縮バイトからトークンへのマッピングに関する具体的なガイドはまだ見つかりません。通常、答えは「トークナイザーに依存します」ですが、これは非常に満足のいく答えではありません。実際にはそれほど大きな変化はありません。バイト/トークンの比率は主に入力のモダリティに基づいていることがわかりました。ここにあります

いくつかの数値は大幅にずれている可能性がありますが、ナプキンの計算には十分近いものです。
テキスト: 10 メモリ内バイト/トークン、5 圧縮バイト/トークン
画像: 1000 メモリ内バイト/トークン、100 圧縮バイト/トークン
ビデオ: 画像と同様、ほとんどのビデオは単なる写真のグループである小さなセグメントで処理されます。
オーディオ: 10K メモリ内バイト/トークン、5K 圧縮バイト/トークン
Alpaca の実験ではテキストを扱うため、1 秒あたりのトークンを 10 倍して 1 秒あたりのバイト数を求めることができます。つまり、GPU に電力を供給して快適な状態を維持するには、なんと…160KB/s のデータが必要ですが、圧縮があるため、おそらく約 80KB/s の I/O 🐌 しか意味しません。ここで、データをランダムな順序で読み取る場合は、1 秒あたりの IOPS にも注意する必要があるかもしれません。 lancedb では Lance ファイル形式を使用するため、値ごとに 1 ～ 2 IOPS が必要です。上記の例では、GPU は最大 200 行/秒を消費できるため、200 ～ 400 IOPS/秒が必要になります。
スループットと IOPS はいずれも、クラウド ストレージから直接簡単に提供できるものです。ただし、タスクをより困難にすることもできます。いくつかの異なるハードウェアを考えてみましょう。
単一サーバーで使用できるハードウェアをスケールアップすると、I/O 要件が増加します。これらのテキストベースのアプリケーションの場合、帯域幅は次のとおりです。

[切り捨てられた]

## Original Extract

A deep dive into data loading for model training — the three pipeline stages, parallelism strategies, shuffling, caching, resumability, and how LanceDB's StreamingDataset fits in.

Data Loading for AI/ML: A Comprehensive Guide
Reverie • The Summit for AI Builders | November 5, San Francisco | Apply to Attend
Use Cases
Curation Find optimal distributions, deduplicate massive datasets, and surface edge cases—all in one place Feature Engineering Build and scale features with Python UDFs, automatic updates, and no table rewrites Search & Retrieval Unified vector, full-text, and hybrid search with SQL filters for production-ready retrieval Training Train directly from curated data with up to 70% MFU and no data movement bottlenecks In Production Partners Docs Blog Community
Contact Us
Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. Searching in Blog Posts Feature Engineering for Multimodal Data: From Laptop to Cluster with LanceDB This is some text inside of a div block. Run a multimodal feature-engineering pipeline in ten minutes with geneva-examples, then rebuild its CLIP embedding step with LanceDB's feature engineering API, the same calls you'd use to enrich your own tables on LanceDB Enterprise. Announcing Reverie Summit: What AI’s Next Breakthroughs Are Made On This is some text inside of a div block. One day with the researchers and engineers building generative video, world models, physical AI, and the data systems beneath them. ⚡ Multi-Bit RaBitQ Without Refine, 🌋 Bytedance’s Lance-Based AI Stack, 🤖 Lance for Embodied AI Data This is some text inside of a div block. Multi-bit RaBitQ hits 96% recall without refine overhead, Volcano Engine cuts training pipelines from 7 days to 1, Lance powers 100 Hz robotics data ingestion, plus upcoming events, enterprise updates, community updates. Why CrewAI Rebuilt Agent Memory on LanceDB, Powering 2B+ Agent Executions This is some text inside of a div block. CrewAI replaced a two-system memory stack with LanceDB, then rebuilt agent memory to resolve contradictions and gate recall on confidence. It now ships by default across 12 million monthly downloa
[truncated]
Non-essential cookies are disabled by default.
Closing this banner does not confirm any choice.
See our Privacy Policy for more information.
LanceDB uses cookies and similar technologies to improve your experience, analyze traffic, and to show you relevant content and advertising. Do not sell or share my personal information Do not sell or share my personal information Consent See Privacy Policy for more info. Cancel Save Preferences
Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. LanceDB uses cookies and similar technologies to improve your experience, analyze traffic, and to show you relevant content and advertising. Accept
Data Loading for AI/ML: A Comprehensive Guide
This is a title This is a subtitle In machine learning tasks, "data loading" is the process of moving data into some kind of algorithm. Typically we are training some kind of model but there are many different variations. So many that I find myself always balancing between overly-specific examples and overly-broad generics.
I/O STAGE
CPU STAGE
GPU STAGE
Storage
Database
Host Machine
CPU · RAM
GPU
Model
The challenge I face in explaining data loading is that it is both extremely simple (iterate a dataset) and full of nuance (this whole guide). Data loading is also the bridge between the data world and the model world and topics are often confusing only because of the choices in terminology. In this article I will attempt to give a comprehensive overview of the data loading process, from the perspective of a data engineer. I'll explain the concepts, describe the challenges, and explain common performance pitfalls. While this article will be focused on pytorch and lancedb the information should apply to other libraries as well.
To begin with, let's consider a very specific example: supervised fine tuning (SFT) applied to model named Qwen2.5-0.5B-Instruct using the Alpaca dataset. What we are doing is taking a large generic model and fine tuning it to perform a more narrowly tailored task. The model is ~1GB of weights. Note that the model has already been fine tuned (hence the "-Instruct" suffix) but it is small and readily available and we can always fine tune it further. The Alpaca dataset is a nice open dataset of instruction/response pairs consisting of about 25MB of data and 50K rows.
Qwen 2.5
Generic understanding
Alpaca
50K instruction pairs
+
fine-tuning
Fine-tuned Model
Instruction answering
The Simplest Data Loader
If this sounds complex, it isn't. Let's look at things in code:
model = load_model( "Qwen2.5-0.5B-Instruct" )
tbl = open_table( "alpaca" )
for batch in tbl:
train(model, batch)
transformed = transform(batch) That's really all there is to it. We want to take some data, transform it a little, let the model take a look at it, and then we're done. There is quite a bit more nuance, as we are regrettably going to see, but we should not forget the basic premise of what data loading is.
Models are traditionally trained on GPUs. Data is traditionally stored on some kind of object storage or disk. Some amount of "data transform" traditionally happens on the CPU. This gives us three stages in our standard data loader.
The I/O stage is when we are loading our data from storage into (usually) the CPU. The resource at play here is distinct from CPU/GPU capabilities. If we are loading from cloud storage we probably care about the NIC(s) attached to the server. We might have hundreds of concurrent HTTP requests to maximize throughput. If we are loading from disk we care about what kind of disk is attached and we typically don't need as much concurrency.
Cloud Storage
CPU
HTTP
NVMe SSD
CPU
io_uring
InfiniBand
GPU
GPUDirect
Different examples of the I/O stage
Normally these resources are rated in GB/s. Capacities can range from a pretty solid baseline of 1GB/s to experimental storage systems that can deliver TB/s. We also have to consider IOPS/s. If we are performing small reads then we may be IOPS limited instead of bandwidth limited. A good baseline might be 1000-5000 IOPS/s for cloud storage all the way to tens of millions of IOPS/s for high-performance NVMe servers.
As a data engineer this fact is somewhat surprising. We're used to I/O being the limiting factor in just about every data-intensive application out there. However, LLMs are...well...large. For every tiny (<1KB) token that we pass through an LLM we need to perform billions of computations. This is vastly more compute-intensive than traditional OLAP tasks.
The next stage is to prepare the data for the GPU. The details at this stage are highly flexible. The first part is typically decompressing and decoding the bytes we receive from storage. This is familiar to me as a data engineer. Next, however, we might to do things like "normalization", "transformation", and "tokenization". To be honest, it doesn't really matter. I normally just smile and nod when this part is described to me.
ENCODED IMAGES
CPU DECODE
GPU STAGE
Encoded Images
JPEG · binary
BOTTLENECK
CPU Decode
Decoding JPEG images
IDLE
GPU
Idle · waiting
A poorly optimized CPU stage (e.g. JPEG decoding with one worker) can easily
become a bottleneck.
A good baseline is probably about 1GB/s per core. Extremely simplistic decoding can handle tens or hundreds of GB/s per core. Very expensive decoding (e.g. decoding JPEG images) can drop far below those speeds. The parallelism here is typically based on the number of cores we have available. This can vary from instance to instance and cloud to cloud.
We can also try and skip the CPU stage entirely. We can pre-tokenize our data and store the already transformed data, reducing the amount of CPU work needed in our actual training task. We can also push the work into the GPU, allowing the GPU to decode and transform the data. However, since the GPU is probably the bottleneck then the last thing we want to do is put more work onto it.
Finally we are ready to actually train the model or run inference. Rating GPUs is incredibly difficult because the capabilities of a GPU are so flexible. The most common scoring we see is "tokens per second" and this is going to be a model-dependent score. So the tokens per second on Qwen will be different than the tokens per second on DeepSeek. This can make it difficult to spec out the performance of our entire pipeline but we discuss that more in the next section.
I/O STAGE
CPU STAGE
GPU STAGE
I/O
CPU
GPU
Pipeline throughput is limited by the narrowest stage — adding capacity to
wider stages won't help until you've addressed the bottleneck.
In this era of foundation models the models themselves are extremely large. The amount of compute needed per byte is massive. We should expect the GPU to be the bottleneck. Typically, the goal of data loading is to feed data to the GPU fast enough that the GPU is never idle.
The three stage model described above is the most common scenario, but not the only one. For example, it is not too uncommon to load data from multiple data sources and stitch them together or we might do some remote processing in the CPU stage introducing additional communication. Stage boundaries are just places where our resources or parallelism change. Most of the rest of this article will be focused on a simple three stage solution (although there is one section on shuffling where we split the I/O into two stages for a four stage model) but the tools and techniques apply to other scenarios as well.
The ultimate performance goal is for model training to run as fast as possible. For a data loader that typically means we want to load data fast enough to keep the GPU busy at all times. The exact answer will always depend on your circumstances, but as of 2026, a good rule of thumb is "not that fast" .
I'd love to quantify that further but even if we know the capabilities of our hardware it can be difficult to convert from compressed bytes (I/O throughput) to in-memory bytes (CPU throughput) to tokens (GPU throughput). Let's look at a concrete example. In the introduction I provided an example where we fine-tune Qwen with the Alpaca dataset. Let's run on this test machine:
CPU: 16 cores (Intel Xeon, 2.00 GHz)
Disk: Local NVMe (virtual disk)
The fastest I can get the training to run is 16K tokens per second. So how fast do we need the I/O stage and CPU stages to be? In other words, how many bytes per second do we need to satisfy 16K tokens per second?
‍
I have yet to find a concrete guide for mapping from compressed bytes to tokens. The answer is typically "it depends on your tokenizer" which is an extremely unsatisfying answer. In practice there is not that much variation. I've found that the bytes/token ratio is primarily based on the modality of your input. Here are some numbers that are likely off significantly but close enough for napkin math.
Text: 10 in-memory bytes/token, 5 compressed bytes/token
Images: 1000 in-memory bytes/token, 100 compressed bytes/token
Video: Same as images, most video is processed in small segments which are just a group of pictures
Audio: 10K in-memory bytes/token, 5K compressed bytes/token
Since our Alpaca experiment deals with text, we can multiply our tokens per second by 10 to get bytes per second. That means, to keep our GPU fed and happy, we need a whopping…160KB/s of data which probably only means about 80KB/s of I/O 🐌 since we have compression. Now, if we're reading our data in random order, we might also need to care about IOPS per second. In lancedb , we use the Lance file format, and so we need 1-2 IOPS per value. In our example above the GPU is able to consume at most 200 rows/second and so we need 200-400 IOPS/second.
Both our throughput and IOPS are something we can easily deliver directly from cloud storage. However, we can make the task more challenging. Let's consider some different hardware.
As we scale up the hardware available on a single server we increase the I/O requirements. For these text based applications our bandwidth is

[truncated]

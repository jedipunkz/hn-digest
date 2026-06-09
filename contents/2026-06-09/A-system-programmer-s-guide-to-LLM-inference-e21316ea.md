---
source: "https://blog.xiangpeng.systems/posts/how-to-llm-inference/"
hn_url: "https://news.ycombinator.com/item?id=48461660"
title: "A system programmer's guide to LLM inference"
article_title: "A system programmer’s guide to LLM inference – Xiangpeng's blog"
author: "aoli-al"
captured_at: "2026-06-09T16:05:34Z"
capture_tool: "hn-digest"
hn_id: 48461660
score: 2
comments: 0
posted_at: "2026-06-09T14:31:21Z"
tags:
  - hacker-news
  - translated
---

# A system programmer's guide to LLM inference

- HN: [48461660](https://news.ycombinator.com/item?id=48461660)
- Source: [blog.xiangpeng.systems](https://blog.xiangpeng.systems/posts/how-to-llm-inference/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T14:31:21Z

## Translation

タイトル: システムプログラマーのための LLM 推論ガイド
記事のタイトル: システムプログラマのための LLM 推論ガイド – Xiangpeng のブログ
説明: 依存関係のないローカル LLM 推論エンジンを Rust で構築しましょう。

記事本文:
システムプログラマーのための LLM 推論ガイド
名前の読み方は？ Qwen3.6-35B-A3B-UD-Q4_K_M.gguf
最初のテンソルの読み取り
Q8_0ブロック
行列乗算
単純な dequant-then-multiply は悪いです
状態: KV キャッシュと DeltaNet 状態
私の博士号スポンサーである InfluxData 、 Bauplan 、 SpiralDB 、そしてウィスコンシン州と連邦政府の納税者に少し感謝を申し上げましょう。
LLM は非常に重要になっているので、私 (おそらくあなたも) をもっと理解したいと考えています。学習するための最良の方法は、LLM を構築することです。
このブログ投稿では、システム プログラマの観点から、LLM 推論について学んだことを共有します。
ほとんどのマシンで実行できるようにモデル Qwen3.6-35B-A3B-UD-Q4_K_M.gguf を選択しましたが、「最新の LLM」として数えられるほど複雑でもあります 1 。ここでサポートされるモデルはこれだけです。最終的には、プリフィルを 100 トークン/秒で実行し、デコードを 15 トークン/秒で実行できるようになります。これは、CPU のみのマシンではそれほど悪くありません。
1 2026年4月発売モデル。
このブログでは、ローカル LLM 推論エンジンの重要な部分のほとんどを説明します。
GPU アクセラレーション。これは純粋な CPU 推論エンジンです。おそらく GPU については後でフォローアップする予定です。
まだ GPU ではないため、MTP (投機的デコード) はありません。
{ベンダー固有のもの |クローズドソース}（CUDA など）は対象外です
名前の読み方は？ Qwen3.6-35B-A3B-UD-Q4_K_M.gguf
Qwen ( /kwɛn/ と発音、「when」に「kw」を付けたもののように) はモデル ファミリ名です。それは（テクノロジー業界では）あまり人気のない中国企業であるアリババからのものです。
3.6 はバージョン番号です。以前のバージョンは 3.5 と 3 だったので、このファミリーはしばらくの間存在していました。
モデルサイズは35Bです。 350 億はパラメータの数です。各パラメータが 8 ビット (fp8 または int8) の場合、モデルは約 35 GB になります。
A3B はモデルが有効化されることを意味します

トークンを生成するための 30 億のパラメータ。また、トークンごとにすべてのパラメーターがアクティブになる高密度モデルとは異なり、モデルが MoE (専門家の混合) モデルであることも意味します。
UD-Q4_K_M はモデルが 4 ビット (Q4) に量子化されていることを意味し、K_M は量子化スキームです (モデルを量子化する方法は多数あります)。 UD は Unsloth Dynamic quantization の略です。 Unsloth は、こ​​れらの量子化されたバリアントを作成する会社です。
gguf はモデル ファイル形式です。 safetensors とは異なり、 gguf は自己完結型の形式です。単一のファイルにモデルの実行に必要なものがすべて含まれています。実際には、それは単なるメタデータとモデルの重みです。
ファイル形式は (意図したとおり) 非常に退屈です。図 1 に示すように、モデルを説明するメタデータの後に、テンソル情報とテンソル データのペアが続きます。
テンソル情報はファイルの先頭に保存され、offset フィールドと len フィールドを通じて実際のテンソル データを指します。
struct TensorInfo {
name : String , // 例: `blk.3.attn_norm_weight`
形状：Vec<u64>、
dtype : u32 、
オフセット：u64、
レン：u64、
}
dtype はテンソルごとのフィールドであることに注意してください。これは、同じファイル内の 2 つのテンソルが異なるデータ型 (つまり、異なる量子化スキーム) を使用できることを意味します。これにより、重要なテンソル (埋め込み重みなど) をより多くのビットで保存して精度を高めると同時に、残りをより積極的に量子化することができます。ご想像のとおり、テンソルごとの量子化は芸術です。
量子化スキームを詳しく見ると、モデルには 6 つの異なるタイプのテンソルがあり、バイトの半分以上が Q4_K に量子化されていることがわかります。
BF16 2 テンソル 1.00 MB
F32 368 テンソル 99.78 MB
Q4_K 82 テンソル 11808.00 MB
Q5_K 38 テンソル 6688.00 MB
Q6_K 4 テンソル 1027.85 MB
Q8_0 259 テンソル 1978.38 MB
------ ---- ------------- ----------
合計 753 テンソル 21603.01 MB
メタデータは MOD もエンコードします

El アーキテクチャ: レイヤーの数、モデル ファミリなど。
このモデルは Qwen3-Next ファミリに属します。アーキテクチャ系統の背景については、Sebastian Raschka の「Big LLM Architecture Comparison」を参照してください。
図 2 は、アーキテクチャとその重量配分を簡略化して示した図です。すべての入力トークンについて、モデルはまずそれを埋め込みベクトル (このモデルの次元は 2048) に変換し、次にそれをいくつかの計算層で実行して最終出力を生成します。
Qwen3.6 は典型的なトランスフォーマーではありません。いわゆる DeltaNet レイヤーと従来のアテンション レイヤーを 3:1 の比率で混合します。詳細については後ほど説明します。大まかな直感は、アテンションのトークンごとの状態 (KV キャッシュ) はシーケンスの長さに応じて線形に増加し、その計算は 2 次であるのに対し、DeltaNet の状態は固定サイズであるということです。これは、モデルをより長いコンテキストに拡張するのに役立ちます。
すべての Qwen3.6 ファミリ モデルはこのアーキテクチャを共有しています。より多くのレイヤーを追加し、より広い隠れディメンションを使用することで、ファミリー スケールのより大きなメンバーを拡張できます (たとえば、Qwen3-Next-80B バリアントには 48 レイヤーと 2048 の隠れディメンションがあります)。 2
2 Qwen3.6-35B-A3B には、(3 DeltaNet + 1 Attendance) ブロックの 10 回の繰り返しとして編成された 40 個のレイヤーがあり、つまり 30 個の DeltaNet レイヤーと 10 個のアテンション層になります。 apxml 仕様ページを参照してください。
以下の表は、より詳細な重み配分を示しています。
グループパラメータストアドシェア
--------------------------------------------------------------
トークン埋め込み 508.56 M 515.31 MiB 2.44%
DeltaNet レイヤー 1.01 B 1.01 GiB 4.92%
注目層 272.66 M 276.35 MiB 1.31%
MoE ブロック 32.36 B 18.43 GiB 89.44%
最終出力 508.56 M 397.86 MiB 1.89%
--------------------------------------------------------------
報告された合計 34.66 B 20.60 GiB 100.00%
MoE ブロックが優勢です: モデル サイズの 89.4% を占め、

ch MoE ブロックは約 469 MB です。ただし、各 MoE ブロックはそのパラメータの 1/32 (トークンあたり 256 エキスパートのうち 8 エキスパート) のみをアクティブにし、他のコンポーネントはトークンごとに起動します。したがって、各トークンは約 0.508B + 1.01B + 273M + 32.36B / 32 + 0.508B ≈ 3.3B のパラメーターを有効にし、これは「A3B」の「3B」に四捨五入されます。 3
3 公式の「3B アクティブ」数値は若干異なる方法で計算されます。Qwen3.6 モデル カードを参照してください。球場も同じだ。
LLM 推論には、プリフィルとデコードという 2 つの段階があります。 Prefill は、単一の順方向パスで入力プロンプト (潜在的には数千のトークン) を処理し、レイヤーごとの状態 (KV キャッシュ、DeltaNet 状態) を設定します。通常、これはコンピューティングに依存します。デコードでは一度に 1 つずつトークンが生成され、各ステップで完全な重み行列が FPU を介して再度ストリーミングされる必要があるため、通常はメモリに依存します。この違いについては後ほど説明します。
トークンは、LLM 固有の文字列表現です。これは、ASCII や Unicode などの他の表現と基本的には異なります。42 は * (ASCII の場合) を意味し、U+1F3D4 は 🏔 (Unicode の場合) を意味します。
最初に確認するのは、tokenizer.ggml.tokens メタデータ フィールドに文字列のリストとして保存されている語彙テーブルです。このモデルの最初の 8 つのトークンは次のとおりです: [ ! 、 \ 、 # 、 $ 、 % 、 & 、 ' 、 ( ]。つまり、トークン ID 0 は ! を参照し、ID 7 は ( などを参照します。
モデルが異なれば語彙テーブルも異なる場合があることに注意してください。語彙が少ないほど、同じ文字列を表すためにより多くのトークンが必要になることを意味します。最新の LLM は語彙が大きくなる傾向があります。このモデルには 151,936 の語彙トークンがあり、これは u16::MAX (65,535) を超えるため、トークン インデックスは u32 でなければなりません。 4
4 Anthropic は Opus 4.7 でトークナイザーを変更しました。これにより、同じテキスト入力で 1.46 倍のトークンが生成されます。これにより、価格が 46% まで上昇します。

トークンごとのレート。
トークンから文字列への変換は簡単です。語彙内の各トークンを検索し、文字列を連結します。
1 つの注意点はストリーミング出力です。単一のトークンは部分的な UTF-8 シーケンスのみを表す可能性があるため、デコーダーは完全なコードポイントが出力されるまでバイトをバッファーする必要があります。これは、非 ASCII スクリプトと絵文字で特に一般的です。
文字列からトークンへの変換はより複雑ですが、概念的には辞書ベースの圧縮の仕組みと似ています。まず、入力文字列内のすべてのバイトが単一トークン表現を持つという特性から始めます (設計上)。これにより、あらゆる入力が表現可能であることが保証されます。
次に、バイト ペア エンコーディング (BPE) マージ ルールを適用します。目標は、隣接するトークンを語彙がサポートする最長のシーケンスに貪欲にマージし、総トークン数を減らすことです。
トークン化前。プリトークナイザーは、最初に入力をより小さなチャンク (空白、句読点、数字などに沿って) に分割します。これにより、下流の BPE マージが、分離しておきたい境界をまたがないようになります (たとえば、2 つの漢字にまたがる 1 つのトークンなど)。
BPE のマージ。次に、隣接するトークンに対して BPE マージを実行します。比較的高価ですが、総トークン数が大幅に減少します。
チャットのテンプレート。 <|im_start|> 、 <|im_end|> などの特別なトークンは、会話内の役割を区切ります。モデルはこれらを使用してターン構造を認識します。
大まかに言えば、LLM は次のような関数です。
fn 推論(入力 : Vec < u32 > ) -> u32
テンソルの寿命
入力トークンを取得したら、次のステップは次のトークンを予測することです。
多くのアルゴリズムと数学が関係しますが、正直なところ、そのほとんどを知る必要はありません。推論は一連の行列乗算と考えることができます。 RoPE、QKV、softmax、RMSNorm などが何をするのかを知ることは、話すのには素晴らしいことですが、システムの観点から見ると、

ログラマーの観点からは、それはあまり重要ではありません。
重要なのは tensor の形状です。これによって計算とメモリの要件が決まります。
図 3 は、テンソルの寿命を示しています。青い矢印は行列の乗算、灰色の矢印は正規化やソフトマックスなどの補助的な演算です。
最初に注意すべきことは、形状 [vocabulary_size, embedding_dim] の 2D テンソルである埋め込みテーブルです。この場合は [151936, 2048] です。トークンを埋め込みに変換するのは簡単な検索です。1 行をスライスして [1, 2048] テンソルを取得します。この [1, 2048] は通常活性化テンソルと呼ばれ、後続のすべての層の入力と出力の両方です。
最終的な出力射影は、[1, 2048] アクティベーションを [1, 151936] テンソルに変換し、ボキャブラリ内の各トークンにロジットを割り当てます。最高のロジットを持つトークンを次のトークンとして選択することも、top-k / top-p / 温度サンプリングなどのサンプリング アルゴリズムを使用して変動を取得することもできます。
残りのブロック (attention、DeltaNet、および MoE ブロック) はすべて [1, 2048] を入力として受け取り、[1, 2048] を出力として生成します。内部的には、それぞれが行列乗算のチェーンであり、正規化やソフトマックスなどの補足的な演算が行われます。
上記では単一のトークン (デコードの場合) について説明しています。 prefill は、アクティブ化が [1, 2048] ではなく [n, 2048] であることを除いて同じです。計算は変わりません。
理想的な世界では、すべてのテンソルは単純な f32 テンソルになり、行列の乗算は単純で適切に最適化されます。
しかし、f32 は高価すぎます。35B モデルの場合、すべて f32 にすると、35B × 4 = 140 GB のメモリが必要になります。その上、f32 は実際の LLM 推論には過剰です。精度をほとんど損なうことなく 8 ビットまで下げることができ、4 ビットでもほとんどの情報が保持されます。 5
5 Q4_K_M は通常、元の FP16 モードの約 92% を保持します

l の品質はバイトの約 25% です。 Q8_0 は「本質的にロスレス」で、パープレキシティ ドリフトは約 0.01 です。
重みあたりのビット数を減らすと、メモリ フットプリントが削減されるだけでなく、特にメモリ帯域幅がボトルネックとなる GEMV の場合、計算が高速になることがよくあります。ロードするデータが少なくなると、クリティカル パスが短くなります。
逆に、特にファイルに量子化スキームが混在しているため、matmul 自体がより複雑になります。話は戻ります。
前にリストしたように、モデルには 6 つのテンソル タイプがあります。ここでは、2 つの代表的なもの、 Q8_0 と Q4_K のみについて説明します。
図 4 に Q8_0 ブロックを示します。これは最も単純な量子化スキームです。fp32 値を取得し、スケールで割って、最も近い整数に丸めます。逆量子化はその逆で、量子化された値にスケールを掛けます。スケールは max(|w|) / 127 として選択されるため、量子化された値は [-127, 127] になります。各ブロックは 32 個の重みと 1 つの fp16 スケールを保持し、2 + 32 = 34 バイトにパックされます (出典: llama.cpp block_q8_0 定義)。
Q4_0 は Q8_0 に似ていますが、重みごとに 4 ビットを使用する点が異なります。実際、Q4_0 も広く使用されています。しかし、それには 2 つの問題があります。
最初の問題は、重みが 0 を中心に対称的に分布していると仮定していることですが、常にそうであるとは限りません。たとえば、ブロック内のすべての重みが正であり、

[切り捨てられた]

## Original Extract

Let’s build a local LLM inference engine in Rust with no dependencies.

A system programmer’s guide to LLM inference
How to read the name? Qwen3.6-35B-A3B-UD-Q4_K_M.gguf
Reading the first tensor
Q8_0 block
Matrix multiplication
Naive dequant-then-multiply is bad
The state: KV cache and DeltaNet state
Let’s take a moment to thank my PhD sponsors: InfluxData , Bauplan , SpiralDB , and the taxpayers of the State of Wisconsin and the federal government.
LLMs have become so important that I (probably you as well) want to understand them better, and the best way to learn is to build one.
In this blog post, I’ll share what I’ve learned about LLM inference, from the perspective of a systems programmer.
I pick the model Qwen3.6-35B-A3B-UD-Q4_K_M.gguf so that it runs on most machines but is also complex enough to count as a “modern LLM” 1 . This is the only model to support here. By the end, we’ll be able to run prefill at 100 tokens/s and decode at 15 tokens/s, not so bad on a CPU-only machine.
1 The model was released in April 2026 .
This blog covers most of the important parts of a local LLM inference engine:
GPU acceleration, this is a pure CPU inference engine. I’ll probably do a GPU follow up later.
No MTP (speculative decoding), because we are not GPU yet.
Things that are {vendor-specific | closed-source}, e.g., CUDA, are not covered
How to read the name? Qwen3.6-35B-A3B-UD-Q4_K_M.gguf
Qwen ( pronounced /kwɛn/ , like “when” with a “kw”) is the model family name; it comes from Alibaba, a not so popular Chinese company (in tech world).
3.6 is the version number; previous versions were 3.5 and 3, so the family has been around for a while.
35B is the model size. 35 B illion is the number of parameters; if each parameter were 8 bits (fp8 or int8), the model would be about 35 GB.
A3B means the model a ctivates 3 B illion parameters to generate a token. It also implies the model is an MoE (Mixture of Experts) model — unlike dense models where all parameters are activated for every token.
UD-Q4_K_M means the model is q uantized to 4 bits (Q4), and K_M is the quantization scheme (there are many ways to quantize a model). UD stands for Unsloth Dynamic quantization ; Unsloth is a company that produces these quantized variants.
gguf is the model file format. Unlike safetensors , gguf is a self-contained format: a single file holds everything you need to run the model. In practice, it’s just metadata plus the model weights.
The file format is pretty boring (as intended): some metadata describing the model, followed by pairs of tensor info and tensor data, as shown in Figure 1 .
The tensor info is stored at the beginning of the file, and points to the actual tensor data through the offset and len fields:
struct TensorInfo {
name : String , // e.g., `blk.3.attn_norm_weight`
shape : Vec < u64 >,
dtype : u32 ,
offset : u64 ,
len : u64 ,
}
Note that dtype is a per-tensor field, which means two tensors in the same file may use different data types (i.e., different quantization schemes). This lets us store important tensors (e.g., the embedding weights) with more bits for higher precision, while quantizing the rest more aggressively. As you might guess, per-tensor quantization is a fine art.
A closer look at the quantization schemes shows that the model has 6 different types of tensors, more than half of the bytes are quantized to Q4_K:
BF16 2 tensors 1.00 MB
F32 368 tensors 99.78 MB
Q4_K 82 tensors 11808.00 MB
Q5_K 38 tensors 6688.00 MB
Q6_K 4 tensors 1027.85 MB
Q8_0 259 tensors 1978.38 MB
------ ---- --------- ----------
total 753 tensors 21603.01 MB
The metadata also encodes the model architecture: the number of layers, the model family, and so on.
The model belongs to the Qwen3-Next family ; for background on the architecture lineage, see Sebastian Raschka’s “Big LLM Architecture Comparison” .
Figure 2 is a simplified view of the architecture and its weight distribution. For every input token, the model first converts it into an embedding vector (with dimension 2048 for this model), then runs it through several layers of computation to produce the final output.
Qwen3.6 is not a typical transformer: it mixes the so-called DeltaNet layers with conventional Attention layers at a 3:1 ratio . We will get to the details later; the high-level intuition is that attention’s per-token state (the KV cache) grows linearly with sequence length and its compute is quadratic, while DeltaNet’s state is fixed-size — this helps the model scale to longer context.
All Qwen3.6 family models share this architecture; larger members of the family scale by adding more layers and using a wider hidden dimension (e.g., the Qwen3-Next-80B variant has 48 layers and a 2048 hidden dim). 2
2 Qwen3.6-35B-A3B has 40 layers organized as 10 repetitions of (3 DeltaNet + 1 Attention) blocks, so 30 DeltaNet layers and 10 Attention layers. See the apxml spec page .
Table below shows the more detailed weights distribution:
Group Params Stored Share
-----------------------------------------------------------------
Token embedding 508.56 M 515.31 MiB 2.44%
DeltaNet layers 1.01 B 1.01 GiB 4.92%
Attention layers 272.66 M 276.35 MiB 1.31%
MoE blocks 32.36 B 18.43 GiB 89.44%
Final output 508.56 M 397.86 MiB 1.89%
-----------------------------------------------------------------
Reported total 34.66 B 20.60 GiB 100.00%
MoE blocks dominate: they account for 89.4% of model size, and each MoE block is roughly 469 MB. However, each MoE block only activates 1/32 of its parameters (8 of 256 experts per token), while the other components fire for every token. So each token activates approximately 0.508B + 1.01B + 273M + 32.36B / 32 + 0.508B ≈ 3.3B parameters, which rounds to the “3B” in “A3B”. 3
3 The official “3B active” figure is computed slightly differently — see the Qwen3.6 model card . The ballpark is the same.
There are two stages of LLM inference: prefill and decode. Prefill processes the input prompt (potentially thousands of tokens) in a single forward pass and populates the per-layer state (KV cache, DeltaNet state). It is usually compute-bound. Decode generates tokens one at a time, and on each step the full weight matrix must be streamed through the FPUs again, so it is usually memory-bound. We’ll come back to this distinction later.
A token is an LLM-specific representation of a string. It isn’t fundamentally different from other representations like ASCII or Unicode, where 42 means * (in ASCII) and U+1F3D4 means 🏔 (in Unicode).
The first thing to look at is the vocabulary table, stored in the tokenizer.ggml.tokens metadata field as a list of strings. The first 8 tokens of this model are: [ ! , \ , # , $ , % , & , ' , ( ]. So token id 0 refers to ! , id 7 refers to ( , and so on.
Note that different models may have different vocabulary tables. A smaller vocabulary means more tokens are needed to represent the same string. Modern LLMs tend to have larger vocabularies — this model has 151,936 vocabulary tokens, which exceeds u16::MAX (65,535), so the token index must be u32 . 4
4 Anthropic changed their tokenizer in Opus 4.7, which produces 1.46× more tokens on the same text input — a hidden ~46% price hike at the same per-token rate.
Token-to-string is easy: look up each token in the vocabulary and concatenate the strings.
The one caveat is streaming output: a single token may represent only a partial UTF-8 sequence, so the decoder needs to buffer bytes until a complete codepoint can be emitted. This is especially common with non-ASCII scripts and emoji.
String-to-token is more involved, but conceptually similar to how dictionary-based compression works. We start with the property that every byte in the input string has a single-token representation — by design — which guarantees that any input is representable.
Then we apply byte-pair encoding (BPE) merge rules; the goal is to greedily merge adjacent tokens into the longest sequences the vocabulary supports, which reduces the total token count.
Pre-tokenization. A pre-tokenizer first splits the input into smaller chunks (along whitespace, punctuation, digits, etc.) so that downstream BPE merges don’t span boundaries we’d prefer to keep separate — e.g., a single token that straddles two Chinese characters.
BPE merge. We then run a BPE merge over adjacent tokens. It is comparatively expensive but reduces the total token count substantially.
Chat templates. Special tokens like <|im_start|> , <|im_end|> , etc., delimit roles in a conversation; the model uses these to recognize turn structure.
Roughly speaking, an LLM is a function:
fn inference(input : Vec < u32 > ) -> u32
The life of a tensor
Once we have the input tokens, the next step is to predict the next token.
There are lots of algorithms and math involved, but honestly, you don’t need to know most of it — you can think of inference as a chain of matrix multiplications. Knowing what RoPE, QKV, softmax, RMSNorm, etc. do is cool to talk about, but from an system programmer’s perspective it doesn’t really matter.
What matters is the shape of the tensors , because that determines the compute and memory requirements.
Figure 3 shows the life of a tensor; blue arrows are matrix multiplications, grey arrows are auxiliary math like normalization and softmax.
The first thing to care about is the embedding table, a 2D tensor of shape [vocabulary_size, embedding_dim] , in this case [151936, 2048] . Converting a token to an embedding is a simple lookup: slice out one row to get a [1, 2048] tensor. This [1, 2048] , usually called the activation tensor, is both the input and the output of every subsequent layer.
The final output projection converts the [1, 2048] activation into a [1, 151936] tensor, assigning a logit to each token in the vocabulary. We can pick the token with the highest logit as the next token, or use a sampling algorithm like top-k / top-p / temperature sampling to get variation.
The rest of the blocks — attention, DeltaNet, and the MoE blocks — all take [1, 2048] as input and produce [1, 2048] as output; internally each is a chain of matrix multiplications with some supplementary math like normalization and softmax.
The above describes a single token (the decode case); prefill is the same except the activation is [n, 2048] rather than [1, 2048] . The math is unchanged.
In an ideal world, every tensor would be a simple f32 tensor, and matrix multiplication would be simple and well-optimized.
But f32 is too expensive: for our 35B model, going all-f32 would require 35B × 4 = 140 GB of memory. On top of that, f32 is overkill for practical LLM inference — we can drop to 8 bits with almost no loss of accuracy, and even 4 bits retains most of the information. 5
5 Q4_K_M typically retains ~92% of the original FP16 model’s quality at roughly 25% of the bytes; Q8_0 is “essentially lossless” with perplexity drift around 0.01.
Using fewer bits per weight not only reduces the memory footprint but often makes compute faster , especially for GEMV, where memory bandwidth is the bottleneck — less data to load means a shorter critical path.
The flip side is that the matmul itself gets more complicated, especially since the file contains a mix of quantization schemes. We’ll come back to that.
As listed earlier, the model has six tensor types; here we only discuss two representative ones: Q8_0 and Q4_K .
Figure 4 shows a Q8_0 block. It’s the simplest quantization scheme: take an fp32 value, divide by a scale, and round to the nearest integer. Dequantization is the reverse — multiply the quantized value by the scale. The scale is chosen as max(|w|) / 127 , so the quantized values land in [-127, 127] . Each block holds 32 weights plus a single fp16 scale, packed into 2 + 32 = 34 bytes ( source: llama.cpp block_q8_0 definition ).
Q4_0 is similar to Q8_0 , except it uses 4 bits per weight; in fact Q4_0 is also widely used. But it has two problems.
The first problem is that it assumes weights are symmetrically distributed around 0, which is not always the case. If, for example, all weights in a block are positive and

[truncated]

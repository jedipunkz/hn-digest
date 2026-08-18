---
source: "https://layog.io/blog/llm-speedrun-arch/"
hn_url: "https://news.ycombinator.com/item?id=49341379"
title: "LLM Speedrun: Architecture"
article_title: "LLM speedrun: Architecture | Layog's Blog"
image: ""
author: "layog"
captured_at: "2026-08-18T05:20:19Z"
capture_tool: "hn-digest"
hn_id: 49341379
score: 1
comments: 0
posted_at: "2026-08-18T04:46:15Z"
tags:
  - hacker-news
  - translated
---

# LLM Speedrun: Architecture

- HN: [49341379](https://news.ycombinator.com/item?id=49341379)
- Source: [layog.io](https://layog.io/blog/llm-speedrun-arch/)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T04:46:15Z

## Translation

タイトル: LLM Speedrun: アーキテクチャ
記事のタイトル: LLM スピードラン: アーキテクチャ |ラヨグのブログ
説明: Qwen2.5-Coder-7B-Instruct モデルを最初から再構築した Speedrun LLM アーキテクチャ - 「LLM Speedrun」シリーズの 3 つのうちの 1 つ目。

記事本文:
設定 削除された行を表示 折りたたみ可能なカード 元を折りたたむ 展開 LLM スピードラン: アーキテクチャ
2026 年 8 月 17 日 • #LLM #transformers #architecture Y HN ディスカッション Qwen2.5-Coder-7B-Instruct モデルを最初から再構築することによる Speedrun LLM アーキテクチャ - 「LLM Speedrun」シリーズの 3 つのうちの 1 つ目。
この記事全体はAIによって生成されたテキストを一切使用せずに私によって書かれています。 Python コードも AI によって設計されたものはありません。 AI を使用して記事を校正してギャップを強調し、Python コードを Rust に変換しました。
対象者 このブログ シリーズは、LLM の状態とスケーラブルなサービスについての個人的なメモを公開するものです。 LLM / トランスフォーマーに特有のアイデアの動機については説明しますが、機械学習や言語モデリングの背後にある構成要素については説明しません。そのために、読者が正規化、標準フィードフォワード ネットワーク、LSTM、線形代数の基礎などのトピックを理解していることを期待します。コードを進めるには、Python の numpy か Rust の rayon クレートに慣れていることが前提です (実際、私はこのブログを書いている間に rayon に慣れてきました)。
コードの実行 ほとんどのセクションにはコードが添付されており、ローカルにコピーして直接実行できます。 Qwen の完全な実装はハーネス セクションで入手できます。
uv run < file.p y > < arg s > Rust の場合は、以下を使用します。
カーゴ +nightly -Zscript < file.r s > -- < arg s >
これは 3 部構成のシリーズの最初の部分であり、アーキテクチャから始めて LLM の個々の部分を詳しく掘り下げています。計画されている他の 2 つの投稿は、推論とトレーニングです。
すべてのスケーラーは、1 × 1 1 \times 1 1 × 1 行列として表されます。
すべてのベクトルは、 d e m b e d × 1 d_{embed} \times 1 d e mb e d × 1 のような列ベクトルとして表されます。
X X X ベクトルの行列は X X X 連結列として表されます

d e m b e d × X d_{embed} \times X d e mb e d × X のような n 個のベクトル。
× \times × は行列の乗算を表し、 ⊙ \odot ⊙ は要素ごとの乗算を表します。
しばらくの間、言語をモデル化する一般的な方法の 1 つは、次のトークンを予測することでした。
これが、RNN またはその最も成功した派生製品である LSTM の背後にある前提です。ただし、これらのモデルでは、現在のトークンを処理する前に、以前のすべてのトークンを処理する必要があります。これにより、トレーニング プロセスが本質的にシリアルになります。
大規模なコーパスのセットをトレーニングすることができないため、これは大きなボトルネックでした (そして明らかに、モデルが流暢に話すことができるようにするには、人間のすべての知識を総計してトレーニングする必要がありました)。
トランスフォーマーのアーキテクチャは、2014 年の論文で紹介された「注意」という中核メカニズムを使用して、この問題を解決しました。 Google の人々はそれを見て、「もしかしたら…もしかしたら、必要なのはこれだけ?」と言いました。これが 2017 年の「attention is all you need」論文につながり、これが機械翻訳にとって皮肉なことに、言語モデルの計算を並列化する道を切り開きました。
RNN と同様に、入力テキストは学習された埋め込みに変換されます。具体的には、長さ T T T トークンのシーケンスが与えられた場合、シーケンスの埋め込みは、 d e m b e d × T d_{embed} \times T d e mb e d × T の行列で表されます。すべてのベクトルは列ベクトルとして表されます。しかし、立場はどうでしょうか？ RNN では、トークンが一度に 1 つずつ処理されるため、位置は自動的に行われますが、各トークンが位置に依存しない方法で処理されるトランスフォーマーの場合はそうではありません。したがって、位置はトークンの埋め込みに追加される埋め込みとしてもエンコードされます (最新の LLM は位置を追加せず、「ロータリー位置埋め込み」を使用します)。
アテンション メカニズムでは、各トークンが「どのように」というクエリを行うことが必要になります。

キーを使用して応答するすべてのトークン (それ自体を含む) に対して、「あなたと私は関連していますか?」というメッセージが表示されます。その後、すべてのトークンの値の重み付けされた一部が、クエリを行っているトークンの最終的な値となります。入力は、(深層学習におけるものと同様に) 多くのアテンション レイヤーで構成されるトランスフォーマー ブロックを通過します。
RNN では、単一の状態ベクトルは、ある時点までモデルによって認識されたすべてのテキストの情報を保持します (任意の大きさにすることができます)。アテンションにより、各トークンは過去のすべてのトークンを振り返ることができるため、ニーズに基づいてコンテキスト エンコーディングを動的に変更できます。したがって、並行してトレーニングでき、過去を振り返ることができるため、言語を表現する際に注意メカニズムが非常に強力になることが可能になりました。
具体的には、各トークン (サイズ d e m b e d d_{embed} d e mb e d 、入力語彙トークンの埋め込み次元 ) は、「値」空間と呼ばれる高次元空間 (サイズ d m o d e l d_{model} d mod e l 、トランスフォーマーが動作する隠れた次元 ) で意味を持ちます。ここで、各トークンは、他のトークンが自分にとってどれだけ重要であるかを知りたいと考えています (そして、他のトークンも同じことを尋ねます)。別の言い方をすると、トークン x x x の「価値」がトークン y y にどれだけ影響するかということです (ここで、 y y は x x x に等しくても構いません。トークンはそれ自体にとってどれだけ重要であるかを知る必要があります)。
それをどうやって理解するのでしょうか？各トークンが他のすべてのトークンに「クエリ」を要求し、他のすべてのトークンがクエリ トークンとの関連性を示す「キー」を提供する場合はどうなるでしょうか。これは、各トークンには 3 つの高次元表現 K K K 、 Q Q Q 、および V V V があり、次元 d mod e l × 1 d_{model} \times 1 d mod e l × 1 があり、これは 3 つの行列 W K W^K W K 、 W Q W^Q W Q 、および W V W^V W V の次元 d mod e l × d e m b e d d_{model} \ によって計算されます。

回 d_{embed} d mod e l × d embed - すべての行列が同じサイズである必要はありませんが、通常は同じサイズです。
ここで、サイズ d e m b e d × T d_{embed} \times T d e mb e d × T で示される入力シーケンス (T T T は入力内のトークンの数) は、他の 3 つのシーケンス行列に変換されます。
そして、各キークエリペアは次のように解決されます。
ＫＴ×Ｑ）
注目の重みのサイズ 注目^{weights} 注目の重み s は T × T T \times T T × T で、各トークンが他のすべてのトークンに払っている注目に注目します。
最後に、attention の出力は、attention スコアで重み付けされた値の合計です。
A 無制限の重みに基づいて、多数のものの確率的寄与を出力する手法は何ですか?わかりました - ソフトマックス。
？ A これは、ソフトマックスが誤動作しないようにするための単なる正規化であり、その分散は 1.0 です。
すべてのトークンと他のすべてのトークン、さらには将来のトークンも含めて?
私は到着の外国人ではありません、私は順番に読んでいます
私たちは物事を並行して見ていきます。画像の各ピクセルを読み取るわけではありません。すべてを順次埋め込んで表現できます。したがって、将来を見据えることには利点があります。
しかし言語の場合はそうではありません。したがって、左下の三角形を − ∞ -\infty − ∞ に設定した T × T T \times T T × T 行列だけの因果マスクによって、トークンが将来のトークンに付随しないことが保証されます。
K T × Q + M )
ここで、 i i i は行インデックス、j j j はアテンションが計算される列インデックス トークンです。 。
しかし、どうやって物事を記憶しているのでしょうか？
トランスには、古典的な完全接続層であるフィードフォワード ネットワークも備えているため、これで終わりではありません。この層は巨大で、通常は 4 ⋅ d mod e l 4 \cdot d_{model} 4 ⋅ d mod e l の中間活性化を持ちます。

ns d f f d_{ff} d f f で表されます。これは、LLM の重みの大部分でもあります (通常、その 70% を占めます)。最新の LLM は、ReLU (これも一般的には SwiGLU) ではなくゲート アクティベーションを使用します。
この FFN 層は LLM の「頭脳」であると考えられていますが、ML モデルの解釈可能性は活発な研究分野です。そこで、Google で検索してみましょう (LLM ブログの皮肉です)。
トランスの 1 つの層からの出力は次の層に渡されます。これらの層は好きなだけ積み重ねることができます (そして、スケーリング則と呼ばれるものをトレーニングする能力もあります。これについてはシリーズのトレーニング部分で書きます)。
ただし、1 つ矛盾があることに気づくかもしれません。最初の層への入力のサイズは d e m b e d × T d_{embed} \times T d e mbed × T ですが、出力のサイズは d mod e l × T d_{model} \times T d mod e l × T です。したがって、残りの層は d mod e l d_{model} d mod e l size の入力を取得するか、 d e m b e d d_{embed} d e mb e d を d mod e l d_{model} d mod e l と等しく設定する必要があり、通常は後者が選択されますが、これだけが理由ではありません。
ソフトマックスはアテンション層の一部であり、FFN にはアクティベーション関数も含まれることに注意してください。これにより、モデルが勾配消失を学習することが非常に困難になったり、複数のトランスフォーマー層がスタックされるため、不安定な爆発的アクティベーションを学習したりすることが難しくなります。したがって、transformer は「ResNet」から同じアイデアを借用し、残りの接続を使用します。
トランスでは、残留接続は 2 か所に追加されます。 A t t e n t i o n o u t p u t Attendance^{output} A t e n t i o n o u tp u t の計算後と、最終出力位置 O u t p u t Output O u tp u t です。数学的に:
これには、 d mod e l = d e m b e d d_{model} = d_{embed} d mod e l = d e mb e d が必要です。
残留物あり

フローに応じて、ネットワーク全体の意図が変わりました。トランスの各層は、初期入力の摂動として機能します。概念的には、最初の層はより大きな摂動 (ネットワークは言語の詳細を理解しようとしている) として機能し、後の層はより小さな摂動 (ネットワークはより大きな画像の表現を微調整している) として機能します。これが数学的にどのようにサポートされるかについては後で説明します。
正規化
2015 年以降に作成されたモデルと同様に、ネットワークは正規化する必要があります。基本的な考え方は同じです。つまり、各層への入力の分布が制約され、明確に定義されている場合、ネットワークはより良く学習します。それ以外の場合、深いネットワークでは、いくつかの行列乗算の直後に、再び爆発的なアクティベーションによって非常に高い値に達する可能性があります。変圧器には以下が含まれます。
爆発的なアクティベーションを引き起こすことで有名な残留接続。
アクティベーション出力はソフトマックスに依存します。この場合、大きな入力がスパイク状の出力を引き起こし、クロス トークン アテンション出力が低下する可能性があります。
正規化は、A t t e n t i o n o u t p u t tention^{output} A t e n t i o n u t p u t と、トランス層の入力または出力のいずれかに適用されます。後者の場合、正規化がどこに適用されるかが重要であり、Pre-Norm と Post-Norm として知られています。
ポストノルム モデルからの出力を見てみましょう。
これを標準化前のモデルからの出力と比較します。
入力残差は妨げられず、正規化パスを適用することなく、勾配は適切に戻ります。
これにより、当然のことながら、最初の層と比較して、後の層から小さな摂動が生じます。層への入力は正規化されるため、 F ( R M S N or m ( I l ) ) I l \frac{F(RMSNorm(I_l))}{I_l} I l F ( R M S N or m ( I l )) は、層が深くなるほど小さくなります。
アテンションメカニズムは各トークンを表します

3 つの d mod e l d_{model} d mod e l サイズのベクトルとして。その後、トークンは互いの関係を把握します。しかし、どういう関係なのでしょうか？これらのトークンは、「それ」が文の残りの部分から何を指しているのかを伝えることになるのでしょうか?それとも文章の文法について話すのでしょうか？あるいは、辞書に載っていない単一の単語を理解しようとしており、そのため複数の単一文字トークンに分割されている可能性があります。
単一の大きな注意行列があると、モデルが複数のことを並行して推論することが困難になります。経験的証拠はまた、それぞれの注意が 1 つの特定の事柄について推論しようとしていることを示しているため、そのような大規模な d mod e l d_{model} d mod e l 表現を持つことは無駄であり、低ランクの表現で十分です。
したがって、MHA では、単一の大きなアテンションを計算するのではなく、アテンションが複数の小さな表現に分割されます。具体的には、 h h h 個のアテンション ヘッドがあり、それぞれの内部表現のサイズ d k = d mod e l h d_k = \frac{d_{model}}{h} d k = h d mod e l が存在します。各ヘッドの出力は最終的に連結されて、注目層の d mod e l d_{model} d mod e l のサイズの出力ベクトルが復元されます。
方程式では (i i i は特定の注意の頭部を表します):
K i T × Q i ）
ATTENTION OUT PUT = [ATTENTION 0 OUTPUT AT 1

[切り捨てられた]

## Original Extract

Speedrun LLM architecture by reconstructing Qwen2.5-Coder-7B-Instruct model from scratch - first of three in the "LLM speedrun" series.

Configuration Show Removed Lines Collapsible Cards Collapse Original Expand LLM speedrun: Architecture
Aug 17, 2026 • #LLM #transformers #architecture Y HN Discussion Speedrun LLM architecture by reconstructing Qwen2.5-Coder-7B-Instruct model from scratch - first of three in the "LLM speedrun" series.
This entire article is written by me without any kind of text generated from AI. None of the python code is designed by AI either. I used AI to proofread the article to highlight gaps, and to translate the python code to Rust.
Audience This blog series is me publishing my personal notes on the state of LLMs and scalably serving them. I’ll discuss the motivation for ideas specific to LLMs / transformers, but not the building blocks behind machine learning and language modeling. To that end, I expect the reader to understand these topics: normalization, standard feed forward networks, LSTMs and basics of linear algebra. To follow along in code, I expect you to either be comfortable with Python’s numpy or with Rust’s rayon crate (I actually familiarized myself with rayon while writing this blog).
Running code Most sections have code attached with them that you can copy in your local and run directly. The full Qwen implementation is available in harness section.
uv run < file.p y > < arg s > For rust, use:
cargo +nightly -Zscript < file.r s > -- < arg s >
This is a first of three part series, deep diving into individual parts of LLMs, starting with architecture. The other 2 planned posts are: inference, and training.
All scalers are represented as 1 × 1 1 \times 1 1 × 1 matrix.
All vectors are represented as column vectors, like d e m b e d × 1 d_{embed} \times 1 d e mb e d ​ × 1 .
A matrix of X X X vectors is represented as X X X concatenated column vectors, like d e m b e d × X d_{embed} \times X d e mb e d ​ × X .
× \times × represents a matrix multiplication, while ⊙ \odot ⊙ represents an element wise multiplication.
For a while, one popular way to model language has been to predict the next token.
That is the premise behind RNNs or their most successful derivative, LSTMs . But these models require all previous tokens to be processed before processing the current one. This makes the training process inherently serial.
This was a huge bottleneck as we couldn’t train over a large set of corpus (and apparently we needed to train over sum of all human knowledge for models to be able to speak fluently).
The transformers architecture solved that problem, with the core mechanism of “attention,” which was introduced in a 2014 paper . Folks at Google looked at it and said “maybe… maybe that’s all you need?” leading to the “Attention is all you need” paper in 2017 --- unironically for machine translation --- paving the way to parallelize the computation of language models.
Similar to RNNs, the input text is converted into learned embeddings. Concretely, given a sequence of length T T T tokens, the embedding for the sequence is represented by a matrix of d e m b e d × T d_{embed} \times T d e mb e d ​ × T All vectors are represented as column vectors . But what about the position? In RNNs the position is automatic, as the tokens are processed one at a time but that is not the case with transformers where each token is processed in a position-independent way. So, position is encoded as an embedding as well which is added to the embedding of a token (modern LLMs do not add positions, rather use “Rotary Position Embedding” ).
Attention mechanism entails that each token asks a query “how relevant are you to me” to every token (including itself) that answers using its key. Then the weighted some of values of all tokens is the final value of the querying token. The inputs pass through the transformer block which is made up of many attention layers (as anything in deep learning).
In RNNs, the single state vector carries the information of every text that has been seen by the model up to a point (and it can be arbitrarily large). Attention allows each token to look back at all the past tokens, hence allowing it to change its context encoding dynamically based on its needs. So, being able to train in parallel and being able to look back allowed attention mechanism to be extremely powerful in representing language.
Concretely, each token (of size d e m b e d d_{embed} d e mb e d ​ Embedding dimension of input vocabulary tokens ) has a meaning in some high dimensional space called a “value” space (of size d m o d e l d_{model} d m o d e l ​ Some hidden dimension that transformer operates in ). Now each token wants to know how much other tokens matter to it (and other tokens will also ask the same of it). Or stated otherwise, how much of “value” of token x x x should affect a token y y y ( where y y y can be equal to x x x A token has to know how much it matters to itself as well ).
How to figure that out? What if each token asks a “query” to every other token, and all those other tokens provide a “key” which tells how relevant it is to the query token. This means that each token has 3 high dimensional representation K K K , Q Q Q and V V V Of dimension d m o d e l × 1 d_{model} \times 1 d m o d e l ​ × 1 , which is calculated by 3 matrices W K W^K W K , W Q W^Q W Q , and W V W^V W V Of dimension d m o d e l × d e m b e d d_{model} \times d_{embed} d m o d e l ​ × d e mb e d ​ - it is not necessary for all the matrices to be of the same size, but they generally are .
Now the input sequence Denoted by I I I of size d e m b e d × T d_{embed} \times T d e mb e d ​ × T where T T T is the number of tokens in our input is transformed into 3 other sequence matrices:
And each key query pair is resolved as:
​ K T × Q ​ )
Size of A t t e n t i o n w e i g h t s Attention^{weights} A tt e n t i o n w e i g h t s is T × T T \times T T × T noting the attention each token is paying to every other token.
Finally the output of attention is attention score weighted sum of values:
A What’s the technique to output probabilistic contribution of a bunch of things based on some unbounded weights? You got it - Softmax.
​ ? A That’s just normalization to make sure that softmax does not misbehave and its variance is 1.0.
Every token with every other token, even the future ones?
I am no alien from Arrival, I read sequentially
We do look at things in parallel too. We do not read each pixel of an image Everything can be represented as embedding sequentially. So, looking into the future has its advantages.
But for language that’s not the case. So a causal mask Just a T × T T \times T T × T matrix with lower left triangle set to − ∞ -\infty − ∞ makes sure a token does not attend to future tokens.
​ K T × Q ​ + M )
where i i i is the row index and j j j is the column index Token whose attention is being computed. .
But how does it remember things?
Attention is not the end of things for a transformer, which also has a classic fully connected layer, a feed-forward network. This layer is huge, typically having 4 ⋅ d m o d e l 4 \cdot d_{model} 4 ⋅ d m o d e l ​ the intermediate activations Denoted by d f f d_{ff} d f f ​ . This is also the bulk of the weights in an LLM (generally accounting for 70% of it). Modern LLMs use gated activations rather than ReLU, that too typically SwiGLU.
This FFN layer is considered to be the “brain” of an LLM, however, interpretability of ML models is an active area of research. So Google around (ironic in an LLM blog).
The output from one layer of transfomer is passed to the next one. One can stack as many of these layers as one wants (and have the capacity to train, something called scaling laws that I’ll write about in the training part of the series).
You might notice one discrepancy though. The input to the first layer is of size d e m b e d × T d_{embed} \times T d e mb e d ​ × T , while the output is of size d m o d e l × T d_{model} \times T d m o d e l ​ × T . So either the rest of the layers need to take input of d m o d e l d_{model} d m o d e l ​ size or d e m b e d d_{embed} d e mb e d ​ is set equal to d m o d e l d_{model} d m o d e l ​ and generally the latter is chosen but not for just this reason .
Notice that softmax is a part of the attention layer and FFN will also include an activation function. This makes it very hard for the model to learn vanishing gradient or makes learning unstable exploding activations as multiple transformer layers are stacked. So transformer borrows the same idea from “ResNet” and uses residual connections.
In transformer, residual connection is added at two places: After calculating A t t e n t i o n o u t p u t Attention^{output} A tt e n t i o n o u tp u t and at the final output location O u t p u t Output O u tp u t . Mathematically:
This does require d m o d e l = d e m b e d d_{model} = d_{embed} d m o d e l ​ = d e mb e d ​
With residual flow, the intent of the entire network has changed. Each layer of transformer now acts as a perturbation on the initial input. Conceptually, initial layers will act as a larger perturbation (the network is trying to figure out the language specifics) and the later layers as smaller perturbations (the network is finetuning the representation of the larger image). Later, I’ll discuss how this is supported mathematically .
Normalization
Like with any model created after 2015, networks need to normalize, with the basic idea being the same: network learns better when the distribution of inputs to each layer is constrained and well defined. Otherwise in a deep network, just after a few matrix multipliers we can reach really high values attacked by exploding activations again . Transformer contains:
Residual connections, famous for leading to explosion in activations.
Activation output depending on softmax where large inputs can lead to spiky outputs and hence poor cross token attention outputs.
Normalization is applied to A t t e n t i o n o u t p u t Attention^{output} A tt e n t i o n o u tp u t and to either input or output of a transformer layer. For the latter, where normalization is applied matters and is known as Pre-Norm v/s Post-Norm.
Take a look at the output from a post-norm model:
and compare this with the output from pre-norm model:
The input residual is unobstructed and the gradients flow back nicely, without subjected to normalization pass.
This naturally leads to small perturbations from the latter layers as compared to initial layers. The input to a layer is normalized so F ( R M S N o r m ( I l ) ) I l \frac{F(RMSNorm(I_l))}{I_l} I l ​ F ( R M S N or m ( I l ​ )) ​ gets smaller and smaller for deeper layers.
Attention mechanism represents each token as three d m o d e l d_{model} d m o d e l ​ sized vectors. The tokens then figure out the relationship between each other. But what kind of relationship? Are those tokens going to tell what “it” refers to from the rest of the sentence? Or are they going to talk about the grammar of the sentence? Or maybe they are trying to understand a single word which was not in the dictionary and hence got split into multiple single character tokens.
Having a single large attention matrix makes it difficult for the model to reason about multiple things in parallel. Empirical evidence also shows that each attention is trying to reason about one specific stuff, so having such a large d m o d e l d_{model} d m o d e l ​ representation is a waste, where a low-rank representation will do.
So, rather than computing a single large attention, in MHA, attention is split into multiple smaller representations. Concretely, there are h h h attention heads, each with internal representation of size d k = d m o d e l h d_k = \frac{d_{model}}{h} d k ​ = h d m o d e l ​ ​ . The output of each head is concatenated in the end to recover d m o d e l d_{model} d m o d e l ​ sized output vector of the attention layer.
In equations (where i i i represents a specific attention head):
​ K i T ​ × Q i ​ ​ )
A t t e n t i o n o u t p u t = [ A t t e n t i o n 0 o u t p u t A t t e n t i o n 1

[truncated]

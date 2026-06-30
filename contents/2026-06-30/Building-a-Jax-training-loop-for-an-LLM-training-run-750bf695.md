---
source: "https://www.gilesthomas.com/2026/06/llm-from-scratch-34a-building-a-jax-training-loop-for-an-llm-training-run"
hn_url: "https://news.ycombinator.com/item?id=48736594"
title: "Building a Jax training loop for an LLM training run"
article_title: "Writing an LLM from scratch, part 34a -- building a JAX training loop for an LLM training run :: Giles' blog"
author: "gpjt"
captured_at: "2026-06-30T18:34:54Z"
capture_tool: "hn-digest"
hn_id: 48736594
score: 1
comments: 0
posted_at: "2026-06-30T17:59:32Z"
tags:
  - hacker-news
  - translated
---

# Building a Jax training loop for an LLM training run

- HN: [48736594](https://news.ycombinator.com/item?id=48736594)
- Source: [www.gilesthomas.com](https://www.gilesthomas.com/2026/06/llm-from-scratch-34a-building-a-jax-training-loop-for-an-llm-training-run)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:59:32Z

## Translation

タイトル: LLM トレーニング実行のための Jax トレーニング ループの構築
記事のタイトル: LLM を最初から書く、パート 34a -- LLM トレーニング実行のための JAX トレーニング ループの構築 :: Giles のブログ
説明: 私

記事本文:
LLM を最初から書く、パート 34a -- LLM トレーニング実行のための JAX トレーニング ループの構築 :: Giles のブログ
el.dataset.currentDropdown = '')
}">
ジャイルズのブログ
何かを学び始めたときに見つけたかった投稿を書いています...
について
お問い合わせ
アーカイブ
カテゴリー
ブログロール
2026年6月 (7)
NSLU2 オフサイト バックアップ プロジェクト (13)
ソフトウェア開発ツール (1)
:: (ブログ可能 a) => a -> IO ()
いくつかの意見は、さまざまな程度の確実性を持って保持されています
LLM を最初から作成する、パート 34a -- LLM トレーニング実行のための JAX トレーニング ループの構築
1年以上、私はセバスチャン・ラシュカの本を使っています
「大規模な言語モデルを (ゼロから) 構築する」 --
そしてそれを読んで派生した多数のサイドプロジェクトも
現代のAIについて学ぶためのカリキュラムのようなもの。最後のタスク
私はメモを使って LLM をゼロから構築してトレーニングすることを自分自身に設定していました --
この本への言及はなく、この本の後に私が書いたモデル コードへの言及もありません。
出力として、Raschka に基づいた最高の PyTorch モデルと同じくらい優れたものが必要でした。
コード -- 32 億トークンでトレーニングされた基本モデル。(明らかに限定的ですが) 私の評価対象です。
オリジナルの GPT-2 Small の品質に近いと評価されています。
別のフレームワークを使用したかったのですが、以前に使用したコードをオウム返しにしないようにするためです。
なんとなく思い出したので、
Twitterで人々に尋ねたところ、
私が使用すべきものであり、勝者は JAX でした。
私はラシュカの本とは少し異なるルートをたどりました。
彼は裏返しの視点をとり、注意などについて説明し、徐々に構築していきます。
完全な GPT-2 スタイルのモデルを作成し、構築する
その上にトレーニングループがあります。行きたかった
アウトサイドイン: 可能な限りシンプルなトレーニングを行うためのトレーニング ハーネスを組み立てます。
実際の LLM と同様の API を備えたモデル、満足のいく動作が得られます

イオン、そして
完全なアーキテクチャが完成するまで、その単純なモデルに機能を 1 つずつ追加していきます。
計画では（実際にうまくいきました！）、
それぞれの変更によって状況がどのように改善されたか。
これですべてが完了したので、これについて 2 つのパートに分けて投稿します。この中で説明します
トレーニング ハーネスをどのように構築したかについて説明します。次は、実際の構築と
LLMのトレーニング。
JAX の上にあるフレームワークはどれですか?
JAX 自体には比較的最小限の API があり、標準の API は含まれていません。
線形層のようなニューラル ネットワーク コンポーネント。同様に、組み込みはありません
オプティマイザー、データローダー、または同様の ML ユーティリティ。
さて、次のように純粋な JAX だけを使用して LLM を構築することもできました。
以前はおもちゃの XOR モデルを使用しました。でも
これは実際の JAX コードのスタイルで構築したほうがよいと感じました。
これは、多くのユーティリティ ライブラリの一部を使用することを意味します。
JAX サイト自体には、便利そうなリンクがありました。
「JAX を使用してニューラル ネットワークをトレーニングしたい場合は、JAX AI スタックをチェックしてください。」
リンクされたページでは、そのスタックの 2 つのコア部分が次のとおりであることが明らかにされました。
ニューラル ネットワーク コンポーネント用の Flax NNX。
両方を見てみましたが、非常に理解しやすかったです。
確かに、一目見ただけでは NNX はかなり PyTorch っぽいと感じました。彼らのチュートリアルでは
たとえば、唯一の明らかな違いは
JAX-y の導関数スタイルの勾配計算と乱数の計算方法でした。
扱った。そして、乱数でさえ、あまり純粋関数的ではない方法で処理されました。
純粋な JAX よりも -- キーの分割をいじる代わりに、
どういうわけか内部的にそれ自体を分割するステートフル変数のように見えるものを渡します
必要に応じて。
そこで、私が使用するフレームワークは NNX と Optax でした。徹底的に研削するのではなく、
家庭教師

ああ、すぐに飛び込んで、進みながら物事を拾おうと決めた
一緒に。
A-to-A 言語モデルとトレーニング ループ
機能するトレーニング ループを構築するには、トレーニングするための最小限のモデルが必要でした。
実際の LLM ですが、少なくとも少しは LLM と同様に動作するものです。それは取り込むだろう
トークンのシーケンスを作成し、各トークンのログを吐き出します。
LLM がどのように機能するかについての私の推奨モデルでは、トップレベルで
モデルの場合、一連のトークン ID を入力し、次のようにします。
まず、それらを埋め込みに変換して、一連のベクトルを取得します。これを行うには、
テーブルへのルックアップですが、概念的には行列を介した射影として見ることができます。
語彙空間 (特定のトークン ID がワンホット ベクトルである場合) から
埋め込みスペース。
次に、Transformers レイヤーで魔法を実行し、次のトークンの埋め込みを取得します。
これらの層の後の、出力シーケンスの位置 n での埋め込みは、予測されたトークンのためのものです。
次のことを考慮して、入力シーケンスの位置 n にあるトークンの後に来ます。
入力トークンとその左側の他のすべてのトークン。
最後に、それらを埋め込み空間からロジットに逆投影します。今回は実際に実際の
マトリックス（線形層の形式）。ロジット（実行後）
Softmax) は、各トークンが次のトークンである確率を表します。
これらすべてのことから、始めるために作成できる最も愚かな「LLM」が考えられました。
トークン ID を埋め込み空間に投影してから投影するだけのものになります。
語彙スペースに戻ります。トランスフォーマーレイヤーはまったくありません。
次に、次のトークンを予測しようとするのではなく、次のトークンを予測するようにトレーニングします。
そもそも何が入力されたかを「予測」してみてください。言い換えれば、あなたは餌を与えるでしょう
この入力のトレーニング ループ:
太った猫がマットの上に座った
...そしてこのターゲット
太った猫がマットの上に座った
...というよりも、

LLM の通常のセットアップ (フィードする場所)
太った猫はその上に座っていました
...そしてそれにターゲットを与えます
太った猫がマットの上に座っていた
もし私がそれをうまく機能させることができたら -- そしてそれはあなたにもできるようなものだと感じました
膨大な量のトレーニングを行わずに損失をほぼゼロに抑えることができれば、私はそうなれるでしょう
トレーニングループが機能していることはかなり確信していました。 1
私はこれを A-to-A モデルと呼ぶことにしました。
モデル自体のコーディングは驚くほど簡単で、次のようになります。
亜麻輸入nnxから
クラス GPTModel ( nnx . Module ):
def __init__ (
自分自身、
vocab_size 、 context_length 、
emb_dim 、
n_heads 、 n_layers 、
qkv_bias 、
ドロップレート 、
RNG 、
):
自分自身。 token_embedding = nnx 。埋め込む(
num_embeddings = vocab_size 、
機能 = emb_dim 、
rngs = rngs 、
)
自分自身。出力ヘッド = nnx 。リニア (
in_features = emb_dim 、
out_features = vocab_size 、
use_bias = False 、
rngs = rngs 、
)
def __call__ ( self , xs ):
input_embeddings = self 。 token_embedding ( xs )
返却自己。出力ヘッド (入力埋め込み)
そこには、モデルが知っていたパラメーターについての定型文がたくさんあります。
完全な LLM を構築するときに必要になります。実際に何かを実行するコードがあるためです。
しかし、トレーニングループはもう少し楽しかったです。
先ほども言いましたが、ここでの私の計画は、LLM の内部を確実に理解することでした。
私のメモから 1 つを再構築することで正しかったです。 「メモのみ」という制限
トレーニングループ自体には当てはまらなかったので、少しずつ練習することにしました。
PyTorch DistributedDataParallel コード
クラウドで元のモデルをトレーニングするために使用していました。
私が使用した最初のバージョンはここにあります。
main 関数がある一番下の部分から始めましょう。
それは、「実行」の概念を処理するための定型的なものから始まります。これはパターンです
私はほとんどのプロジェクトで使用していることに気づきました。モデルで作業するときに便利です

ウルに
毎回内容を変更しながら、複数のトレーニングを実行できるようになります。チェックポイントは守りたいので、
将来の参照用に、それぞれのメタデータとトレーニング チャートを保存します。
私のレポでは、
「runs」ディレクトリがあり、その中に追跡したい各トレーニング実行のサブディレクトリがあります。
これらのサブディレクトリには、モデルを構成するための JSON ファイルがあります。
モデル.json 、
もう 1 つはトレーニングのハイパーパラメータなどを設定するためのものです。
train.json 。
(この段階では、これらのハイパーパラメータの多くが未使用であることに注意してください。
後で必要になることが分かっていたので、怠惰からそこに保管しておきました。)
したがって、これらをロードして main 関数を開始します。
次のステップは、トレーニングの 1 つを完全に無視することです。
ハイパーパラメータ、gradient_accumulation_steps 。絶対やりたかった
勾配累積、
しかし、それは後回しにすることにしました。しっかりとしたシンプルなトレーニングを受けた方が良い
まず走り終わった、と私は感じた。
次に、download_dataset を使用して、使用するデータセットをローカル ディスクにダウンロードします。
(最新のコピーが存在しない場合にのみダウンロードされます)。
次のステップでは、load_dataset を呼び出して RAM にロードします。あることがわかります。
そこには別のハードコーディングされた変数、 world_size があります。これはマルチ GPU からの名残です
これはすべて、DistributedDataParallel コードに基づいています。このブログ記事では私はただ
シングル GPU トレーニングのコードをカバーしていますが、DDP のものはそこに残すことにしました。
データセットのラングリングを目的とした、ハードコードされた
後で実装することにした場合に再導入しやすくするため、1 つの GPU に追加します。
JAX でも同様のものがあります。
load_dataset とそれに関連するものを見てみましょう。
39行目まで行くと
コードが表示されます。まず、私たちのデータを追跡する BigTrainDataset があります。
トレーニングデータ。よく見てみると、おかしな点が 1 つ見つかるかもしれません

授業で。これがあります:
自分自身。 xs = all_tokens [: - 1 ] 。 reshape ( - 1 、 microbatch_size 、 seq_length )
自分自身。 ys = all_tokens [: - 1 ] 。 reshape ( - 1 、 microbatch_size 、 seq_length )
この段階では、トークンをそれ自体にマッピングするようにモデルをトレーニングする計画であったことを思い出してください。
次のトークンを予測するのではなく。したがって、ターゲットは入力と同じであり、
より通常の次のトークンは次のようになります (次の投稿では次のようになります)。
自分自身。 xs = all_tokens [: - 1 ] 。 reshape ( - 1 、 microbatch_size 、 seq_length )
自分自身。 ys = all_tokens [1:] 。 reshape ( - 1 、 microbatch_size 、 seq_length )
次に、load_dataset 関数を使用して、適切なサブセットをロードします。
ローカルディスク上のコピーからのデータ
これらの BigTrainDataset オブジェクトの 1 つに。
この最初のバージョンを実行したときに、メモリ不足の問題が発生しました。
データを GPU の VRAM にロードしようとしていました。これは、JAX のデフォルトの動作です。
GPU があり、CUDA バージョンの JAX がインストールされています。しかし、そこには収まりきらないものが多すぎました。
少し掘り下げた後
JAX のデフォルトデバイスを変更する方法を学びました
これにより、通常のシステム RAM にロードされるようになります。
残念ながら、一度それを実行すると、反復処理が非常に遅いことがわかりました。
6,144 個のトークンからなる 1 つのトレーニング バッチを配列から取得するのに約 1.2 秒かかりました。
これは、それだけでトレーニングの制限が 5,120 トークン/秒になることを意味します。
最終的に、データはメイン RAM にロードされていたものの、ロードされていなかったことを知りました。
メイン RAM にコミットされていなかったため、処理のために GPU にコピーされました --
詳細はこちら。それを修正します（明示的に
jax.device_put への呼び出し）は、データセットから単一のトレーニング バッチを取得することを意味します
それを GPU に置くのにかかる時間は 0.001 秒未満で、はるかに優れていました。
それがそうでした

コードの 55 行目から 58 行目には、何時間にもわたる作業がすべて詰め込まれています。
cpu0 = jax 。デバイス ( "cpu" )[0]
ジャックス付き。デフォルトデバイス (cpu0):
full_dataset =load_file ( dataset_dir / f " { 分割 } .safetensors" )[ "トークン" ]
full_dataset = jax 。 device_put ( full_dataset 、 cpu0 )
load_dataset のロジックの残りの部分は、単に
世界サイズに対して正確に正しいサイズのデータセット (たとえそれが常に正しいサイズであっても)
今すぐ 1 つ）、マイクロバッチ サイズ、グラジエント累積ステップ、およびシーケンス
私たちが取り組んでいる長さ、
main関数に戻りましょう
またまた。データセットをロードしたら、モデル設定を渡してモデルを作成します。
ものと (現在は使用されていない) ドロップアウト率トレーニング ハイパーパラメーターも加えて、Flax NNX オプティマイザーを作成します。
Optax oneをラップします。これは基本的に Flax チュートリアルからのコピー/ペーストでした。
ただし、学習率と重み減衰を使用してオプティマイザーを構成している点が異なります。
トレーニング設定からのハイパーパラメータ:
オプティマイザー = nnx 。オプティマイザー (
モデル、
オプタックス 。アダム（
learning_rate = train_conf [ "learning_rate" ],
Weight_decay = train_conf [ "weight_decay" ],
）、
wrt = nnx 。パラム
)
最後に、電車を呼んでスタートします。

[切り捨てられた]

## Original Extract

I

Writing an LLM from scratch, part 34a -- building a JAX training loop for an LLM training run :: Giles' blog
el.dataset.currentDropdown = '')
}">
Giles' blog
Writing the post that I wished I'd found when I started learning whatever it was...
About
Contact
Archives
Categories
Blogroll
June 2026 (7)
NSLU2 offsite backup project (13)
Software development tools (1)
:: (Bloggable a) => a -> IO ()
Some opinions, held with varying degrees of certainty
Writing an LLM from scratch, part 34a -- building a JAX training loop for an LLM training run
For over a year, I've been using Sebastian Raschka 's book
" Build a Large Language Model (from Scratch) " --
and the multitude of side-projects that have branched out from reading it -- as
something like a curriculum for learning about modern AI. The one final task
I had set myself was to build and train an LLM from scratch just using my notes --
no reference to the book, no reference to the model code I'd written following the book.
As an output, I wanted something as good as my best PyTorch model based on Raschka's
code -- a base model, trained on 3.2B tokens, that my (admittedly limited) evals
ranked as being close to the original GPT-2 small's quality.
I wanted to use a different framework, just to make sure I wasn't parroting code that I'd
somehow memorised, so I
asked people on Twitter which
one I should use, and the winner was JAX .
I took a slightly different route to Raschka's book;
he takes an inside-out perspective, explaining things like attention, gradually building
up a complete GPT-2-style model, and then building
a training loop on top of it. I wanted to go
outside-in: I'd put together a training harness to train the simplest-possible
model with an API similar to a real LLM, get that working to my satisfaction, and then
add features to that simple model, one by one, until it had the full architecture in place.
The plan (which actually worked out nicely!) was that I'd be able to show
how each change improved things.
That's all done now, and I'm posting about it in two parts; in this one, I'll explain
how I built the training harness, and in the next, I'll show the actual building and
training of the LLM.
Which framework on top of JAX?
JAX itself has a relatively minimal API, and doesn't include standard
neural network components like linear layers. Likewise it doesn't have any built-in
optimisers, data loaders or similar ML utilities.
Now, I could have decided to build my LLM using just pure JAX, like
I previously did with a toy XOR model . But
I felt that it would be better to build this in the style that real-world JAX code
is written, which would mean using some of the many utility libraries .
On the JAX site itself, there was a useful-looking link:
"If you’re looking to use JAX to train neural networks, check out the JAX AI Stack !"
On the linked page, it made it clear that the two core parts of that stack were:
Flax NNX for neural network components.
I took a look at both, and they seemed pretty easy to grasp.
Indeed, at first glance, I felt that NNX looked pretty PyTorch-like! In their tutorial
example, the only real obvious difference
was the JAX-y derivative-style gradient calculation and the way that random numbers were
handled. And even the random numbers were handled in a less pure-functional way
than pure JAX -- instead of having to mess around with splitting keys, you could just
pass in what appeared to be a stateful variable that somehow split itself internally
as needed.
So, NNX and Optax were the frameworks I'd use. Rather than grinding through the
tutorials, I decided that I'd just dive right in, and try to pick things up as I went
along.
The A-to-A language model and the training loop
To build a functioning training loop, I needed a minimal model to train -- not an
actual LLM, but something that behaved at least a bit like one. It would take in
a sequence of tokens, and spit out logits for each token.
In my preferred model of how LLMs work , at the top level
for a model, we feed in a sequence of token IDs, then:
Firstly, we convert them into embeddings, so we get a series of vectors. We do this by
a lookup into a table, but we can see it conceptually as a projection via a matrix,
from vocab space (where a particular token ID is a one-hot vector) to an
embedding space.
Next, we do the magic with our Transformers layers, getting embeddings for the next token.
The embedding at position n in the output sequence, after these layers, is for the predicted token to
come after the token at position n in the input sequence, considering that
input token and all other tokens to its left.
Finally, we project those back from embedding space to logits, this time actually using a real
matrix (in the form of a linear layer). The logits (after being run through
softmax) represent the probabilities for each token of it being the next one.
All of that suggested to me that the dumbest "LLM" I could write just to get started
would be one that just projected token IDs into embedding space, and then projected
back to vocab space. No Transformer layers at all.
I'd then train it so that instead of trying to predict the next token, it would
try to "predict" what was fed into it in the first place. In other words, you'd feed
the training loop this input:
The fat cat sat on the mat
...and this target
The fat cat sat on the mat
...rather than the normal setup for an LLM, where you feed it
The fat cat sat on the
...and give it targets of
fat cat sat on the mat
If I could get that to work -- and it felt like the kind of thing where you'd be able
to get the loss down to near-zero without a huge amount of training -- then I could be
reasonably sure that I had a working training loop. 1
I decided to call this an A-to-A model.
Coding up the model itself was ridiculously simple: it looked like this:
from flax import nnx
class GPTModel ( nnx . Module ):
def __init__ (
self ,
vocab_size , context_length ,
emb_dim ,
n_heads , n_layers ,
qkv_bias ,
drop_rate ,
rngs ,
):
self . token_embedding = nnx . Embed (
num_embeddings = vocab_size ,
features = emb_dim ,
rngs = rngs ,
)
self . output_head = nnx . Linear (
in_features = emb_dim ,
out_features = vocab_size ,
use_bias = False ,
rngs = rngs ,
)
def __call__ ( self , xs ):
input_embeddings = self . token_embedding ( xs )
return self . output_head ( input_embeddings )
There's as much boilerplate in there -- for the parameters that I knew that the model
would need when I built out the full LLM -- as there is actual code doing stuff!
But the training loop was a bit more fun.
As I said, my plan here was to make sure my understanding of the internals of LLMs
was correct by rebuilding one just from my notes. That "notes only" restriction
didn't apply to the training loop itself, so I allowed myself to crib a bit from
the PyTorch DistributedDataParallel code
that I'd been using to train the original model in the cloud.
The first version that I used is here .
Let's start at the bottom, where we have the main function .
It starts with some boilerplate to handle the concept of "runs". This is a pattern
I've found myself using in most of my projects. When working on a model, it's useful to
be able to do multiple training runs, changing things each time. You want to keep the checkpoints,
metadata and training charts for each one for future reference.
So in my repo, I'll
have a "runs" directory, and in there subdirectories for each training run I want to track.
In those subdirectories, there are JSON files -- one to configure the model,
model.json ,
and one to configure the training hyperparameters and similar stuff,
train.json .
(It's worth noting that at this stage, a bunch of those hyperparameters were unused;
I kept them in there out of laziness, as I knew I'd need them later.)
So we start our main function by loading those.
Our next step is to completely ignore one of the training
hyperparameters, gradient_accumulation_steps . I definitely wanted to do
gradient accumulation ,
but decided to leave it for later. Better to get a solid, simpler training
run done first, I felt.
Next, we download the dataset we're going to use to our local disk with download_dataset
(which will only download if there's not an up-to-date copy already there).
The next step is to call load_dataset to load it into RAM. You can see that there's
another hard-coded variable there, world_size . This is a holdover from the multi-GPU
DistributedDataParallel code that this was all based on; in this blog post I'm only
covering the code for single-GPU training, but I decided to leave the DDP stuff in there
for dataset-wrangling purposes, hardcoded
to one GPU, so that it would be easier to re-introduce if I later decide to implement
something similar in JAX.
Let's take a look at load_dataset and its related stuff.
If you go up to line 39
you'll see the code. Firstly, there's a BigTrainDataset that keeps track of our
training data. If you look closely, you might spot one oddity in that class. We have this:
self . xs = all_tokens [: - 1 ] . reshape ( - 1 , microbatch_size , seq_length )
self . ys = all_tokens [: - 1 ] . reshape ( - 1 , microbatch_size , seq_length )
Remember that at this stage, the plan was to train the model to map tokens to themselves
rather than to make next-token predictions. So the targets are the same as the inputs, not
the more normal next token, which would look like (and, in the next post, will look like) this:
self . xs = all_tokens [: - 1 ] . reshape ( - 1 , microbatch_size , seq_length )
self . ys = all_tokens [ 1 :] . reshape ( - 1 , microbatch_size , seq_length )
Next, we have a load_dataset function to load the appropriate subset of
the data from the copy on the local disk
into one of those BigTrainDataset objects.
I hit an out-of-memory issue when I ran the first version of this.
It was trying to load the data into my GPU's VRAM -- JAX's default behaviour if you
have a GPU, and the CUDA version of JAX is installed -- and there was too much to fit in there.
After a bit of digging around
I learned how to change the JAX default device
so that it would be loaded into normal system RAM.
Unfortunately, once I'd done that, I found that iterating through it was super-slow --
it took about 1.2 seconds to get one training batch of 6,144 tokens out of the array,
which meant that I'd have a limit of 5,120 tokens/second of training from that alone.
I eventually learned that the data had been loaded into the main RAM, but was being
copied up to the GPU for processing because it had not been committed to the main RAM --
details here . Fixing that (with an explicit
call to jax.device_put ) meant that getting a single training batch from the dataset
and putting it onto the GPU took less than 0.001s, which was much better.
So that was many hours of work that all got packed into lines 55 to 58 of the code:
cpu0 = jax . devices ( "cpu" )[ 0 ]
with jax . default_device ( cpu0 ):
full_dataset = load_file ( dataset_dir / f " { split } .safetensors" )[ "tokens" ]
full_dataset = jax . device_put ( full_dataset , cpu0 )
The remainder of the logic in load_dataset is just to make sure that we have a
dataset that is exactly the right size for the world size (even though that's always
one right now), the microbatch size, the gradient accumulation steps, and the sequence
length that we're working with,
Let's go back to the main function
again. Having loaded our dataset, we create our model, passing in the model configuration
stuff and also the (currently unused) dropout rate training hyperparameter, then we create a Flax NNX optimiser which
wraps an Optax one. This was essentially a copy/paste from the Flax tutorial,
except we're configuring the optimiser with learning rate and weight decay
hyperparameters from the training config:
optimizer = nnx . Optimizer (
model ,
optax . adamw (
learning_rate = train_conf [ "learning_rate" ],
weight_decay = train_conf [ "weight_decay" ],
),
wrt = nnx . Param
)
Finally, we call train to kick off our t

[truncated]

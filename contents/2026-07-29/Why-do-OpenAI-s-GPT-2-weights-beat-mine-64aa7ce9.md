---
source: "https://www.gilesthomas.com/2026/07/why-do-openai-gpt2-weights-beat-mine-1-intro"
hn_url: "https://news.ycombinator.com/item?id=49098367"
title: "Why do OpenAI's GPT-2 weights beat mine?"
article_title: "Why do OpenAI's GPT-2 weights beat mine? :: Giles' blog"
author: "gpjt"
captured_at: "2026-07-29T15:05:53Z"
capture_tool: "hn-digest"
hn_id: 49098367
score: 2
comments: 0
posted_at: "2026-07-29T14:55:11Z"
tags:
  - hacker-news
  - translated
---

# Why do OpenAI's GPT-2 weights beat mine?

- HN: [49098367](https://news.ycombinator.com/item?id=49098367)
- Source: [www.gilesthomas.com](https://www.gilesthomas.com/2026/07/why-do-openai-gpt2-weights-beat-mine-1-intro)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T14:55:11Z

## Translation

タイトル: OpenAI の GPT-2 重みが私のものよりも優れているのはなぜですか?
記事のタイトル: OpenAI の GPT-2 重みが私のものよりも優れているのはなぜですか? :: ジャイルズのブログ
説明: オリジナルの GPT-2 小型ウェイトは、特定の指示に従うタスクにおいて、私が作成したモデルよりもはるかに優れています。その理由を調べてみましょう!

記事本文:
OpenAI の GPT-2 重みが私のものを上回るのはなぜですか? :: ジャイルズのブログ
el.dataset.currentDropdown = '')
}">
ジャイルズのブログ
何かを学び始めたときに見つけたかった投稿を書いています...
について
お問い合わせ
アーカイブ
カテゴリー
ブログロール
2026年7月 (5)
NSLU2 オフサイト バックアップ プロジェクト (13)
ソフトウェア開発ツール (1)
:: (ブログ可能 a) => a -> IO ()
いくつかの意見は、さまざまな程度の確実性を持って保持されています
OpenAI の GPT-2 重みが私のものを上回るのはなぜですか?
LLM をゼロからトレーニングするプロジェクトを終えたとき、私はこうしていました。
小さな謎を残した。なぜ私のモデルはオリジナルの OpenAI よりも命令に従うのが下手だったのか
GPT-2の小型ウェイト?
走っているという評価をいただきましたが、
の第 7 章の命令微調整コードに基づく
「大規模な言語モデルを (ゼロから) 構築する」。
このプロセスでは、アルパカのサンプルを使用してモデルをトレーニングしました。
検証損失が増加し始めるまで命令に従うデータセットを追跡し、その命令で微調整されたモデルを使用してデータセットを生成します。
完了をホールドバックされたテスト セットに設定し、LLM を使用して比較します。
さまざまなモデルからの結果。詳細はこちら;
それを IFT 評価と呼びましょう。
GPT-2 Small の OpenAI の元の重みは、一貫して私の重みを上回っています。
たとえ私のモデルがより技術的な面で彼らのモデルよりも良い結果を得たとしても
評価では、各モデルのクロスエントロピー損失を測定しました。
抑制されたテストシーケンスのセット。これには私は驚きました。期待していただろう
2 つの eval 間に適度に密接な相関関係がある -- より良いテスト損失が意味すること
より良い指示に従います。
これがなぜなのかについていくつか考えがありますが、最近のことを考えると、
ポピー、専用の LLM トレーニング ボックスをセットアップします
私は彼女にそのうちの 1 つをテストする実験を始めることにしました。さらに
実験は間に合うだろう -- 私はそうは思わないが

がブログの焦点になります。もっと見る
謎が解けるまで、時折投稿される、実行中のテーマ。
絶望して諦める…
この投稿では、問題の性質についてもう少し詳しく説明します。
原因として考えられることをいくつか挙げてみます。次回の投稿で結果を報告します
最初の実験の様子。
最近実行した IFT 評価の結果を見てみましょう。私は
OpenAI モデルは太字で強調表示されており、表はテスト損失によって並べ替えられています。
先ほど述べたより技術的な評価であり、低いほど良いと言えます。
モデルはどれだけうまくいったか
IFT 評価は最後の 3 列にあります。 「IFT エポック」列には次のことが表示されます。
モデルがその前に必要なトレーニング エポック数
検証損失が増加し始めた場合、スコアは (この場合) 100 点満点の平均点です。
GPT 5.5 は IFT テスト セットに対するモデルの回答を提供し、ランクはその順位です。
モデルはその平均スコアに関して成立します。
IFT 評価では、OpenAI の重みが最適です。 「中」モデルが最初で、「小」モデルが最初です。
1 つが 2 番目に来ます -- そして、「小さい」と「小さい」の間には顕著な差があります。
その評価に対する私の最良のモデル「Cloud FineWeb、8x A100 40 GiB」のスコア: 26.73 対 20.71。その違い
eval を複数回実行しても一貫性があった -- 私自身の相対的なランキング
モデルはさらに多様になりました。
「中型」モデルのパフォーマンスが優れているのは驚くべきことではありません。
他のモデルの 2 倍以上のサイズ -- しかし、小型モデルが一貫して私のものを上回るのは奇妙でした
テスト損失が低いモデルがあることを考えると。
もう 1 つ目立っているのは、モデルがトレーニングされたエポックの数です。
覚えておいてください、私は検証損失が上昇し始めるまで、つまりモデルが
過学習が始まりました。 OpenAI の重みが次の値に達していることがわかります。

の 2 エポック後
トレーニングの数は 3 ～ 7 でしたが、私のモデルは 3 ～ 7 でした。
私の最初の仮説は、OpenAI の重みがより良い状態から始まっているということでした。
IFT損失の状況における私の立場よりも。さまざまなタスクの損失状況
モデルは元々、 -- "このシーケンスの次のトークンを予測するために事前トレーニングされていました。
Web から取得したテキスト」 -- IFT タスクの損失状況とは異なります。
これは、「このリクエストに対する有用な応答内の次のトークンを予測する」というようなものです。
それについて少し掘り下げてみましょう。
通常の LLM トレーニング パラダイムの背後にある中心的な考え方は、有用なトレーニングを得ることができるということです。
事前トレーニングによるモデルの開始点 -- ウェブから安価に入手できる大量のものを使ったトレーニング --
その後、微調整のためにはるかに高価なタスク固有のデータセットを使用する必要があります。
私は約 32 億の FineWeb トークンでモデルをトレーニングしてきました。
データセット。Web のスクレイピングから生成され、重複が除去されて整理されました。
少し。
代わりに、指示に従うサンプルの 32 億トークンで彼らを訓練していたら --
基本的に、人間とアシスタントの間で事前にパッケージ化された Q&A セッションが行われます。
ほぼ間違いなく、私が説明に従っているモデルよりもはるかに優れているでしょう。
トレーニングの実行に使用されるコンピューティングコストは同じです。
しかし、これら 32 億のトークンを生成するコストは信じられないほど高額になります。
それを行うために人を雇うことを想像してみてください。GPT-2 トークナイザーは、トークンあたり平均約 0.75 ワードを処理します。
24億語を書くのにいくら払う必要があるかわかりませんが、私は
それはかなり多いのではないかと思われます。あれだけの合成データを生成しても、
(Alpaca データセットのような) より大きな LLM は安くはありません。2026 年半ばの時点で、GPT 5.6 Sol を使用すると 144,000 米ドルになります。
さらに悪いことに、

その種のデータでトレーニングした LLM は「脆弱」になるでしょう -- あなたが望むなら
わずかに異なるタスクを解決するために使用します (たとえば、アルパカは単発問題です)
チャットボットのように複数ターンの会話を処理したい場合を想像してください)、
そのように訓練するのは難しいでしょう。
つまり、アイデアは、構造の低いベース LLM を事前トレーニングすると、しかし重要なことに、それでも
「本物」 -- スクレイピングされた Web ページのようなデータ。汎用の基本 LLM を比較的安価に入手できます。
言語の構造などの基本的なこと、できれば一般的なことも発見されているでしょう。
知識（例えば、フランスの首都はパリであるかもしれない）。それができたら、微調整できます
特定の用途に使用します。
ここで、最初の事前トレーニングと、この IFT テストにおけるモデル構築の微調整フェーズの両方を行います。
本質的に同じ方法で行われます。つまり、モデルの予測の相互エントロピー損失を最小限に抑えようとします。
データセット。 1 それらの間で変わったのは、モデルをトレーニングしようとしているターゲットです
予測する。事前トレーニング段階では、Web スクレイピング データを予測できるようにしようとしています。
微調整では、クエリに対する応答を予測しようとしています。
そのアイデアをもう少し数学的な言語に置き換えると（大まかではありますが）、
これを読んでいる本物の数学者なら、おそらく恐怖で悲鳴を上げるだろう）、もし私たちがそう言うのであれば、
事前トレーニングされたモデルは、命令を微調整するためのベースとして役立ちます。
元の事前トレーニング段階の損失状況が合理的であるという仮定
微調整の損失状況と同様です。という場所であることを私たちは願っています。
トレーニング前のランドスケープでは良好で低い (パラメーター空間で) にかなり近い
微調整景観の低い場所。
それはすべて非常に抽象的です。それを視覚化してみましょう。もし私たちが

モデルの損失状況を想像する
パラメータが 2 つあれば、かなり簡単です。 2 つのパラメータの構成要素
2 次元 -- 1 つは左右、もう 1 つは前後とします -- そして
任意の点での損失は垂直方向の寸法です。凹凸のある表面で、少し転がっているような感じです
パラメータの損失が非常に高い地点に丘や山がある風景、
損失が少ない谷や裂け目。
もちろん、実際には 2 つより若干多いパラメータがあります。 163,009,536 で
私の GPT-2-small スタイル モデルのパラメーターでは、損失ランドスケープはサーフェスではなく、
163,009,537 次元空間内の 163,009,536 次元の超曲面 2。良い
それを視覚化する幸運。
実際に想像できる次元のレベルから直感を導き出しながら
そのようなめちゃくちゃ高次元の空間に行くのはしばしば危険です - 奇妙なことが始まります
次元の数が増加するにつれて起こります -- ここで特に注目しているのは、
安全です。風景のイメージは直感に役立ちます。
したがって、トレーニング前のフェーズにはそのようなランドスケープが 1 つあります。微調整を始めると、
喪失の風景は別のものに変わります。私たちが望んでいるのは、その風景です
かなり似ています。トレーニング前の風景の中で私たちがたどり着いた最低点
新しい微調整ランドスケープが交換されると、そのランドスケープの最低点に近づきます。 3
良いニュースは、事前トレーニングを行ってから微調整を行うこのプロセスが機能することがわかっていることです。
最新の AI のほとんどは、それを基盤として使用してトレーニングされます (ただし、その上に追加要素がたくさんあります)。
そして確かに、直感的にそれが機能することを期待するでしょう。
「言語を理解する」喪失の状況と「質問に答える」ことの間にはまったく相関関係がない
1つ。
しかし、これは私たちができることを意味します。

少なくとも、やや手探り的な方法で -- 特徴づける
損失状況の観点から、OpenAI の重みが私のものよりも「優れている」様子。
OpenAI の重みと私の重みは両方とも、トレーニング前の観点からは良好な位置に着地しました。
それが、上の表の「テスト損失」の結果が意味するものです。 GPT-2培地はどの培地よりも優れています
私のモデル -- サイズが約 2 倍であることを考えると、そうなるはずです -- そして GPT-2 Small のほうがうまく機能します
彼らのほとんどよりも。
しかし、OpenAI モデルが着地した列車喪失前の状況の場所は、明らかに、より適切に適合していました。
私自身のモデルの場所よりも、微調整損失ランドスケープの良い場所でした。という事実
彼らは、直感的にその方向を指しているオーバーフィッティングを開始するまでに、より短いエポックしか必要としませんでした。
どこかに近ければ、そこに行くまでの時間は短くなりますが、実際には、
微調整後のスコアが向上したことは明らかです。
しかし、これを書いているときにもう一つ思い浮かんだのは、テストセットの結果です。
また、OpenAI の重みの優位性も指摘していますが、これは私には思いつきませんでした
過去に。
OpenAI の重み: 思ったより優れています!
そのテストセットが何であるかを考えてみましょう。あったデータを取得するには
前処理され、すぐに使用できるようになったので、100 億トークンをすべてダウンロードしました
FineWeb のバージョンを作成し、99% のトレーニングと 1% の検証とテストに分割します。私は
それらの分割のそれぞれで各サンプルをトークン化し、それらを連結して
それぞれに 1 つのシーケンスがあり、サンプルは <|endoftext|> トークンで区切られています。私のトレーニングスクリプト
トレインを分割し、1024 個のトークンのシーケンスに分割し、それらをバッチにまとめました。
そして32億トークン相当を使い果たしました。
損失テストには、検証/テストの位置 50,000,000 から始まる 19,660,800 個のトークンが必要です。
分割して実行します

6 つのバッチに分けて、平均損失を計算します。
さて、GPT-2 の重みは、私が構築したものに基づいてトレーニングされました。
同様に -- Web から大量のドキュメントで構成される大量のテキストを取得し、それらをトークン化します。
すべてをまとめて <|endoftext|> で区切ってから使用します。
しかし重要なのは、それは異なるデータセットだったということです。迷惑なことに、OpenAI にはアクセスできません。
「WebText」データセットですが、たとえそれが何に似ていたとしても、それはかなり安全であると思います
私が思いついたのは、損失テストに使用されたシーケンスとはあまり似ていないということです。
私自身のトレーニングセットです。
OpenAI の小さな重みの損失テストでより良い結果が出たことを考慮すると、
PyTorch でトレーニングしたどのモデルよりも結果が良く、僅差で負けただけでした
JAX でトレーニングしたモデルによるものです (それでも、JAX モデルは幸運だったと思います)
トレーニングを開始する前にランダムな重みを初期化した場合）、次のように言えると思います。
OpenAI の重みはすでに私自身の重みよりもかなり優れています。
彼らはすでに損失状況の中で良い位置にいたが、それは偶然にも起こった。
私のテストセットでは良い位置にいます。
別の例え: 15 人のトレイル ランナーのグループがレースをしていると想像してください。

[切り捨てられた]

## Original Extract

The original GPT-2 small weights are much better than my own models at a specific instruction-following task. Time to look into why!

Why do OpenAI's GPT-2 weights beat mine? :: Giles' blog
el.dataset.currentDropdown = '')
}">
Giles' blog
Writing the post that I wished I'd found when I started learning whatever it was...
About
Contact
Archives
Categories
Blogroll
July 2026 (5)
NSLU2 offsite backup project (13)
Software development tools (1)
:: (Bloggable a) => a -> IO ()
Some opinions, held with varying degrees of certainty
Why do OpenAI's GPT-2 weights beat mine?
When I finished my project training an LLM from scratch , I was
left with a minor mystery. Why were my models worse at instruction-following than the original OpenAI
GPT-2 small weights?
I had an evaluation that I was running,
based on the instruction fine-tuning code in chapter 7 of
" Build a Large Language Model (from Scratch) ".
The process was to train a model on samples from the Alpaca
instruction-following dataset until validation loss started rising, to use that instruction fine-tuned model to generate
completions to a held-back test set, and then to use an LLM to compare the
results from various different models. The details are here ;
let's call it the IFT eval.
OpenAI's original weights for GPT-2 small consistently beat my own
models, even when mine got better results than theirs on a more technical
evaluation, where I just measured the cross entropy loss for each model on a
held-back set of test sequences. This surprised me; I would have expected
a reasonably close correlation between the two evals -- that better test loss would imply
better instruction-following.
I have a couple of thoughts about why this might be, and given that I recently
set up poppy , my dedicated LLM training box
I decided to inaugurate her with an experiment to test one of them; further
experiments will come in time -- though I don't think this will be a focus for the blog. More
of a running theme, with occasional posts until either I solve the mystery, or
give up in despair...
In this post I'll give a bit more detail about the nature of the problem, and
list some of the things I've been thinking might be the cause. In the next post, I'll give the results
of the first experiment.
Let's take a look at the results from my most recent runs of the IFT eval. I've
highlighted the OpenAI models in bold, and the table is sorted by the test loss -- that
more technical evaluation that I mentioned earlier, where lower is better.
How well the model did
with the IFT eval is in the last three columns. The "IFT epochs" column shows
how many epochs of training the model needed before its
validation loss started rising, the score is the average mark out of 100 that (in this case)
GPT 5.5 gives the model's answers to the IFT test set, and the rank is the position the
model holds in terms of that average score.
For the IFT eval, the OpenAI weights are the best. The "medium" model comes first, and the "small"
one comes second -- and there's a noticeable gap between "small" and
the score for my best model for that eval, "Cloud FineWeb, 8x A100 40 GiB": 26.73 vs 20.71. That difference
was consistent across multiple runs of the eval -- the relative rankings of my own
models varied more.
It's not surprising that the "medium" model does better -- it's
more than double the size of the others -- but the small model consistently beating mine was odd
given that I had models with lower test loss.
Another thing that stands out is the number of epochs that the models were trained for.
Remember, I trained them until the validation loss started rising -- that is, until the model
started overfitting. You can see that the OpenAI weights hit that after two epochs of
training, whereas my models ranged between three and seven.
My starting hypothesis was that the OpenAI weights were starting off from a better
position in the IFT loss landscape than mine. The loss landscape for the task the various
models were originally pre-trained for -- "predict the next token for this sequence of
text that was pulled from the web" -- is different to the loss landscape for the IFT task,
which is something more like "predict the next token in a useful response to this request".
Let's dig into that a bit.
The core idea behind the normal LLM training paradigm is that we can get a useful
starting point for a model by pre-training -- training on a whole load of cheaply-available stuff from the web --
and only then have to use our much more expensive task-specific datasets for fine-tuning.
I've been training my models on approximately 3.2 billion tokens of the fineweb
dataset, which was generated from scrapes of the web, and then deduplicated and tidied up
a bit.
If I had instead trained them on 3.2 billion tokens of instruction-following samples --
essentially, pre-packaged Q&A sessions between a human and an assistant -- then they
would almost certainly be much better than the models I have at instruction-following,
for the same cost in compute used to do the training run.
But the cost of generating those 3.2 billion tokens would be incredibly high.
Imagine hiring people to do it: the GPT-2 tokeniser averages about 0.75 words per token,
and while I don't know how much you'd need to pay people to write 2.4 billion words, I
suspect it would be rather a lot. Even generating that much synthetic data from
a larger LLM (as the Alpaca dataset was) would not be cheap -- I make it US$144,000 using GPT 5.6 Sol as of mid-2026.
Even worse, the LLM we trained on that kind of data would be "brittle" -- if you wanted
it to solve an even slightly different task (for example, Alpaca is single-shot question
and answer, so imagine if you wanted it to handle multi-turn conversations like a chatbot),
it would be hard to train it to do that.
So: the idea is that if we pre-train a base LLM on less-structured -- but, importantly, still
"real" -- data like scraped web pages, we'll get a general-purpose base LLM relatively cheaply.
It will have discovered basic stuff like the structure of language, and hopefully some general
knowledge (eg. maybe that the capital of France is Paris). Once we have that, we can fine-tune
it for specific uses.
Now, both the initial pre-training and -- in this IFT test -- the fine-tuning phase of building our model
are done in essentially the same way: trying to minimise the cross entropy loss of the model's predictions against a
dataset. 1 What has changed between them is the targets we're trying to train the model to
predict. In the pre-training phase, we're trying to get it to predict web scrape data,
in the fine-tuning, we're trying to predict responses to queries.
Putting that idea into slightly more mathematical language (albeit loosely enough that
any real mathematicians reading this will probably scream with horror), if we're saying that
a pre-trained model is useful as a base for instruction fine-tuning, what we have is
an assumption that the loss landscape for the original pre-training phase is reasonably
similar to the loss landscape for the fine-tuning. We are hoping that a place that is
nice and low on the pre-training landscape is reasonably close (in parameter space) to
a place that is low on the fine-tuning landscape.
That's all very abstract; let's try to visualise it. If we imagine the loss landscape for a model
with two parameters, that's reasonably easy. The two parameters make up
two dimensions -- let's say left and right for one, forward and back for the other -- and
the loss at any given point is the vertical dimension. It's an uneven surface, a bit like a rolling
landscape, with hills and mountains at points where the parameters have very high loss,
and valleys and clefts where the loss is lower.
In reality, of course, we have somewhat more than two parameters. With the 163,009,536
parameters in my GPT-2-small-style models, the loss landscape is not a surface but
a 163,009,536-dimensional hypersurface 2 in 163,009,537-dimensional space. Good
luck visualising that.
While taking intuitions from levels of dimensionality that we can actually imagine
over to insanely high-dimensional spaces like that is often risky -- weird stuff starts
happening as the number of dimensions goes up -- for what we're specifically looking at here, it's
safe. The landscape image works for intuition.
So, we have one landscape like that for our pre-training phase. When we start fine-tuning,
the loss landscape changes to a different one. What we're hoping is that the landscape
is reasonably similar; that the low point that we wound up at in the pre-training landscape
is close to a low point in the new fine-tuning landscape when that is swapped in. 3
The good news is that we know that this process of doing a pre-train then a fine-tune works.
Most modern AI is trained using it as a foundation (though there's lots of extra stuff on top).
And indeed, you'd intuitively expect it to work -- it would be kind of weird if there was
no correlation at all between the "understand language" loss landscape and the "answer questions"
one.
But this does mean that we can -- at least in a somewhat hand-wavey way -- characterise
the manner in which the OpenAI weights are "better" than mine in terms of the loss landscape.
Both the OpenAI weights and mine have landed in places that were good from the pre-training viewpoint;
that's what the "Test loss" results in the table above mean. GPT-2 medium is doing better than any of
my models -- given that it's about twice the size, it should -- and GPT-2 small is doing better
than most of them.
But the places in the pre-train loss landscape where the OpenAI models landed were clearly better-matched to
good places in the fine-tuning loss landscape than my own models' places were. The fact that
they took fewer epochs to start overfitting intuitively points in that direction -- after all,
if you're closer to somewhere, then it takes less time to get there -- but the fact that they
scored better after fine-tuning makes it pretty clear.
But something else that occurs to me as I write this is that the results on the test set
also point to a superiority in the OpenAI weights -- one that hadn't occurred to me
in the past.
OpenAI weights: better than we thought!
Let's think about what that test set is. In order to get data that was
pre-processed and ready to work with, I downloaded all of the 10 billion token
version of FineWeb, and split it into 99% training and 1% validation and test. I
tokenised each sample in each of those splits, and then concatenated them together into
a single sequence for each, separating the samples with <|endoftext|> tokens. My training script
took the train split, broke it up into 1024-token sequences, assembled them into batches,
and ran through 3.2B token's worth.
The loss test takes 19,660,800 tokens starting at position 50,000,000 in the validation/test
split, and runs them through in batches of six, working out the average loss.
Now, the GPT-2 weights were trained on something that I believe was constructed
similarly -- get a load of text consisting of a bunch of documents from the web, tokenise them
all and tack them together separated with <|endoftext|> s, and then use that.
But, importantly, it was a different dataset. Annoyingly, we don't have access to OpenAI's
"WebText" dataset, but I think it's reasonably safe to say that even if it is similar to what
I came up with, it will be less similar to the sequences used for the loss test than
my own training set is.
Given that the loss test for the OpenAI small weights came out with a better
result than any of the models I trained in PyTorch, and was only narrowly beaten
by the ones I trained with JAX (even then, I think, because the JAX models got lucky
with their random weight initialisation before they started training), then I think we can say that
the OpenAI weights are already considerably better than my own ones.
They were already at a good place in their own loss landscape, but it happened to also
be in a good place in my test set's.
An alternative analogy: if you imagine a group of 15 trail runners having a race

[truncated]

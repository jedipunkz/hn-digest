---
source: "https://www.gilesthomas.com/2026/03/llm-from-scratch-32f-interventions-weight-decay"
hn_url: "https://news.ycombinator.com/item?id=50606414"
title: "Writing an LLM from scratch, part 32f – Interventions: weight decay"
article_title: "Writing an LLM from scratch, part 32f -- Interventions: weight decay :: Giles' blog"
image: "https://www.gilesthomas.com/images/favicons/web-app-manifest-512x512.png"
author: "ykgoon"
captured_at: "2026-11-23T02:18:48Z"
capture_tool: "hn-digest"
hn_id: 50606414
score: 1
comments: 0
posted_at: "2026-11-23T01:41:30Z"
tags:
  - hacker-news
  - translated
---

# Writing an LLM from scratch, part 32f – Interventions: weight decay

- HN: [50606414](https://news.ycombinator.com/item?id=50606414)
- Source: [www.gilesthomas.com](https://www.gilesthomas.com/2026/03/llm-from-scratch-32f-interventions-weight-decay)
- Score: 1
- Comments: 0
- Posted: 2026-11-23T01:41:30Z

## Translation

タイトル: LLM を最初から書く、パート 32f – 介入: 重みの減衰
記事のタイトル: LLM を最初から書く、パート 32f -- 介入: 重量減衰 :: Giles のブログ
説明: 重み減衰とは何ですか?モデルに対して可能な限り最良のトレーニングを実行するには、重み減衰をどのような値に設定するのが適切ですか?

記事本文:
LLM を最初から書く、パート 32f -- 介入: 重量減衰 :: Giles のブログ
el.dataset.currentDropdown = '')
}">
ジャイルズのブログ
何かを学び始めたときに見つけたかった投稿を書いています...
について
お問い合わせ
アーカイブ
カテゴリー
ブログロール
2026年11月 (1)
NSLU2 オフサイト バックアップ プロジェクト (13)
ソフトウェア開発ツール (1)
:: (ブログ可能 a) => a -> IO ()
いくつかの意見は、さまざまな程度の確実性を持って保持されています
LLM を最初から書く、パート 32f -- 介入: 重みの減衰
私はまだ、ゼロからの GPT-2 の小さな基本モデルのテスト損失の改善に取り組んでおり、次のコードに基づいてトレーニングされています。
セバスチャン・ラシュカの本
「大規模な言語モデルを (ゼロから) 構築する」。
私のトレーニング コードには、オプティマイザーを作成する次のコードがあります。
オプティマイザー = トーチ。最適な。アダムW (
モデル。パラメータ ()、
lr = 0.0004、weight_decay = 0.1
)
前回の投稿で私は調べました
学習率、そのコードの lr パラメータを調べて、その値を見つけました。
さらに、スケジュールを設定する (つまり、時間の経過とともに変化させる) ための追加のコードを追加すると、より良いトレーニング結果が得られます。
今回は体重減少についてお話したいと思います。それは何ですか、何のためにありますか？
本当に0.1が最良の値なのでしょうか？
このハイパーパラメータを理解するには、次のことが少し心配でした。
オプティマイザーがどのように機能するかをよく理解する必要があります。私は構築してきました
これはオプティマイザーの確かなメンタルモデルだと思いますが、理解していないと思います
それらについてはまだ十分に説明できるので、投稿を遅らせたいと思っていました。
このブログの後に別のブログ投稿シリーズに書きます。
良いニュースは、重みの減少はオプティマイザーの仕組みの重要な側面である一方で、
-- AdamW の「W」、これが古い Adam オプティマイザーとの違いです。
それは体重減少の異なる治療法へのうなずきです -- yo

方法を知る必要はありません
オプティマイザー自体は、重みの減衰とは何かを理解するために機能します。
代わりに、構築における古くてより基本的な側面を考慮する必要があるだけです。
ML システム -- 正則化。それを掘り下げるために、過学習から始めましょう。
単純な分類タスクを想像してみましょう。次のようなモデルを構築したいと考えています。
このチャート上の任意の点 -- 十字または円がそこに行くべきかを予測し、それをトレーニングします。
すでにあるサンプル データ ポイント:
このデータセットで強力なモデルをトレーニングすると、次のようになったとします。
さて、最初はそれが良い結果かどうかはわかりません。私たちのものを使う必要があります
それを評価するための検証セット。検証ポイントがこれらであるとしましょう
青いもの：
強力なモデルが過剰適合しているように見えることがわかります。トレーニングセットは、
すべて境界によってうまく分割されていますが、検証ポイントは分割されていません。
このような問題に対処するための一般的な解決策は、次のとおりです。
入門 ML コースでは、それほど強力ではないモデルの使用を試します。それほど強力ではないモデル
この場合、2 つのカテゴリを分けるために、より「くねくねした」線が見つかるかもしれません。
おそらく、あまり細かく動かすのに十分なパラメーターがなかったため、
次のような分類子が見つかったかもしれません。
したがって、検証セットを使用して過学習を検出し、複雑さを調整できます。
それを回避しようとするモデルの。
これで問題はありませんが、手動による介入が必要です。やらなければならなかった
トレーニングを実行し、過学習であることを特定してから、そのパラメータを決定します。
新しい単純なモデル (パラメーターはいくつ必要ですか?)。おそらくそうすることができたでしょう
行き過ぎて、次のような結果になりました。
最初からパラメータの適切な数を知る方法はありません。

、
そのため、さまざまな値を試して、最適なバランスを見つける必要があります。
正則化手法は、これを自動化して過剰適合を防ぐように設計されています。
モデルをいじるのは面倒なことはありません。
Dropout についてはすでに見てきました。
それを行うための標準的な方法。私自身のメンタルモデルですが、
それが行うことは、過剰適合を防ぐだけではなく、おそらくそうであると思います。
間違っています -- そして、LLM トレインが同じトレーニング データを 2 回見ることがないことを考えると、
単一エポックの実行で、それを削除するとモデルが改善されることが判明しました。
もう 1 つの手法は、検証が表示され始めたときにトレーニングの実行を停止することです。
損失の増加、「早期停止」とも呼ばれます。それはとても当たり前のことなので、
私が初期の微調整実験をしていたときに、これを独立して思いつきました。
現在、これらのトレーニング実行用の個別の検証セットはありませんが、
単一のエポックでは、表示されるトレーニング データは、抑制されたデータと同じくらい「新しい」ものです。
検証セットは次のようになるため、同様のトリックを使用して、検証損失の代わりに「列車損失が増加し始める」ことを処理できます。
電車を早めに止める理由として浮上している。全く同じものではありませんが、
おそらくそれは十分に近いでしょう。
しかし、このシリーズの列車はどれも、
そんなことは一度もなかった -- 時々電車の損失が少しだけ跳ね上がることはあるが、長期的には
それは下がり続けます。
しかし、巧妙なトリックに依存するテクニックは他にもあります。振り返って考えてみましょう
モデリングに適切なパラメータの数を見つけようとする手動の退屈な方法
タスク。 1 つの数値を試してみたところ、それが過適合であることがわかりました。次に、より低い数値を試してみると、次のことがわかりました。
それがアンダーフィットで、中間のものを試してみると、それがより優れていることがわかりますが、それでもダメです
いずれかの方法で完璧にして、すすぎます

満足のいくものが見つかるまで繰り返します。
最適解を見つけるために解空間を探索するこの種の作業は、まさに私たちが行っていることです。
when training a model.同じように自動化できたら本当にいいですね。
1 つのトリックは、モデルが過剰適合しないようにモデルの複雑さを最小限に抑えたい場合、次のようにします。
モデルの複雑さの尺度を損失関数に追加してみることができます -- そして
通常の勾配降下プロセスは、それと同じようにそれを最小限に抑えようとします。
トレーニング結果自体からの損失を最小限に抑えるようにしてください。 And that brings us
on to weight decay.
重みの減衰による正則化は、すべての要素の「サイズ」が次のとおりであるという仮説から始まります。
モデルの重みを合計すると、モデルの複雑さの尺度になります。もし
モデルの重みが小さい場合、モデルは大きい場合よりも単純になります。 1
この意味での「サイズ」は L2 ノルムの 2 乗です。これは、私たちが見つけたものです。
gradient clipping .
L2 ノルムは基本的に、すべての重みを 2 乗して加算し、その結果得られる合計です。
square-rooted.これは、重み付けされるベクトルの長さと考えることができます。
を表します -- つまり、163M パラメータ モデルの場合、これは
モデルの重みは 1 億 6,300 万次元空間のベクトルとして考慮されます。 2 And by
その二乗を使用すると、より大きな値にさらにペナルティを与える何かが得られます（そして、
平方根を計算する時間も節約できます)。
私にとって、その尺度が本当に表現しているかどうかは直感的には明らかではありません
明確な意味でのモデルの複雑さ。結局のところ、あなたはそれが2倍になると思うでしょう
すべてのパラメータを使用しても、以前よりも複雑になることはありませんが、複雑さは 2 倍になります。
the L2 norm. 3 しかし、それが実際に機能するという裏には、確かな数学があると私は想像しています。
より一般的な方法なので、

数学ウサギに消えていく
この段階では穴があるので、与えられたものとして受け入れます。
つまり、モデルの複雑さの尺度として 2 乗 L2 ノルムを使用しており、次のようになります。
両方を最小限に抑える方法として、トレーニングの損失にそれを加えます。次の質問
それは、トレーニングの損失とモデルの複雑さという 2 つのバランスをどのようにとるかということです。
ペナルティ？
これは、やや手作業的な方法で、電流の量を決定するのと似ています。
重みを調整するときに使用する損失関数の勾配。そのために、 η を使用します。
勾配を適用する前に勾配をスケーリングするための学習率:
そして、「実際の」損失とモデルの複雑さのペナルティの間のバランスは、
同様の方法 -- 通常は表される重みの減衰という数値が得られます。
小文字のラムダ λ を掛けて、L2 ノルムの 2 乗を次のように乗算します。
...ここで、トレーニング入力の通常の損失に対して ℒ を使用しています。
ターゲット、N 2 は重みの 2 乗 L2 ノルム、ℒ ' は結合された重みです。
損失。そして、ℒ ' は、理論上、オプティマイザーを使用して実際に最小化しようとするものです。
しかし、実際には、これをさらに簡単にするために適用できる巧妙な単純化があります。
まず、上の式に小さな変更を 1 つ加えてみましょう。
λ を乗算する前に L2 ノルムを 2 乗します。それは明らかに基礎的な数学を変えるわけではありません。
これは、同じ効果を得るには λ に大きな値を使用する必要があることを意味します。これがなぜ役立つかは、すぐにわかります。
ここで、通常の勾配降下法について考えてみましょう。再度、各重みに対する損失関数の勾配を計算します。
そして、重みの値から学習率 η を引いて更新します。
それを少し定式化し直してみましょう。重量に対する損失関数の勾配は次のとおりです。
その重みに対する偏微分を求めることができるので、

上をこのように書きます
重量減衰を含む損失関数のバージョン、ℒ ' の場合:
さて、上記の ℒ ' を ℒ + λ · N 2 2 として定義したので、次のようになります。
それをそこに置き換えます:
ここで、L2 ノルム N について考えてみましょう。それは
すべての重みの 2 乗の合計の平方根、または同等のことができます。
それを二乗して（上記の式で行うように）、次のように言います。
さて、たくさんのものを足し合わせた導関数は、それぞれを微分しただけです
別々に、そして一緒に追加されます。これを括弧内の 2 つの項に当てはめてみましょう。
...次に、定数 λ と 2 を 2 階偏導関数から取り出します。
次に、多数のものを加算した導関数にルールを再度適用します。
ここで、次の 1 つである 1 つの特定の重み w に対して偏導関数を実行しています。
そこには w 0 、 w 1 などが含まれます。その観点からすると、
他の重みはすべて定数です。つまり、それらの導関数は次のようになります。
w に関してはゼロです。したがって、次の 1 つを除いてすべてを削除することができます。
実際には w であり、最終的には次のようになります。
w に関する w 2 の導関数は、ちょうど 2 w です。
先ほどの N 2 の巧妙な半分のおかげで、次のことが可能になります。
括弧で囲まれた項全体に − η を掛けると、次のようになります。
これは通常の勾配降下法更新とまったく同じであり、変更されていないものを使用します。
重み減衰のない損失関数 -- 追加の減算を除いて
重みの元の値は、学習率 η と
重み減衰値 λ 。
(余談ですが、上記の説明は
これは「従来の」単純勾配降下法と、大まかに言うと Adam にとっては正しいですが、AdamW のトリックは次のとおりです。
物事を少し違った方法で行うこと。それについては、記事を書き始めたときにさらに詳しく説明します
オプティマイザーに関する私の投稿。)
それで:

重み減衰は、モデルの重み付けを防止するための正則化手法です。
必要以上に複雑になってしまいます。 1 つの数 λ があります。
通常のトレーニングの損失に対して複雑さをどれだけ重視するかを決定します。そして、私たちとしては、
コードからわかります:
オプティマイザー = トーチ。最適な。アダムW (
モデル。パラメータ ()、
lr = 0.0004、weight_decay = 0.1
)
...現在、λ を 0.1 に設定しています。それは正しい値ですか?
いつものように GPT-2 論文
彼らが使用したハイパーパラメータの詳細についてはよくわかっていないが、ノスタルジブリストはこう書いている
Tumblrの本当に素敵な投稿
そこで彼らはその数字が何であったかを調査しました。彼らが言うように：
それはほとんどの点で最初の GPT 論文に従っており、その論文では 0.01 の重み減衰が使用されていると述べています。
この論文へのリンクは間違っているようです。なぜなら、それは別の（非常に興味深いですが）論文だからです。
2020 年、GPT-2 論文の 1 年後ですが、これは通常 GPT-1 論文と呼ばれている論文だと思います。
実際、そこでは 0.01 が使用されています。
私たちも
[37] で提案されている L2 正則化の修正バージョンを使用し、すべての非バイアスまたは w = 0.01 で
体重が増える。
ただし、GPT-3 論文へのリンクは正しく見えます。
彼らは、0.1 のウェイト減衰を使用すると言います。
すべてのモデルは 0.1 の重み減衰を使用して、
少量のr

[切り捨てられた]

## Original Extract

What is weight decay, and what is the right value to set it to in order to get the best possible training run for our model?

Writing an LLM from scratch, part 32f -- Interventions: weight decay :: Giles' blog
el.dataset.currentDropdown = '')
}">
Giles' blog
Writing the post that I wished I'd found when I started learning whatever it was...
About
Contact
Archives
Categories
Blogroll
November 2026 (1)
NSLU2 offsite backup project (13)
Software development tools (1)
:: (Bloggable a) => a -> IO ()
Some opinions, held with varying degrees of certainty
Writing an LLM from scratch, part 32f -- Interventions: weight decay
I'm still working on improving the test loss for a from-scratch GPT-2 small base model, trained on code based on
Sebastian Raschka 's book
" Build a Large Language Model (from Scratch) ".
In my training code, I have this code to create the optimiser:
optimizer = torch . optim . AdamW (
model . parameters (),
lr = 0.0004 , weight_decay = 0.1
)
In my last post I looked
into the learning rate, the lr parameter in that code, and found a value for that,
plus some extra code to schedule it -- that is, to vary it over time -- which gave better training results.
This time I want to go into the weight decay. What is it, what is it for, and is
0.1 really the best value?
I was a little concerned going into this that in order to understand this hyperparameter,
I'd need to have a good understanding of how the optimiser works; I've been building
what I think is a solid mental model of optimisers, but I don't think I understand
them well enough to explain them yet, and I've been hoping to delay posting about them
to a separate blog post series after this one.
The good news is that while weight decay is an important aspect of how optimisers work
-- the "W" in AdamW, the thing that makes it different to the older Adam optimiser,
is a nod to its different treatment of weight decay -- you don't need to know how
the optimiser itself works to understand what weight decay is.
Instead, you just need to consider an older and more fundamental aspect of building
ML systems -- regularisation. In order to dig into that, let's start with overfitting.
Let's imagine a simple classification task: we want to build a model that can -- for
any point on this chart -- predict whether a cross or a circle should go there, training it using
the sample data points that we already have:
Let's say that we train a powerful model on this dataset, and it comes up with this:
Now, ab initio we don't know whether that's a good result or not; we need to use our
validation set to evaluate it. Let's say that the validation points are these
blue ones:
We can see that it looks like our powerful model has overfit. The training set is
all nicely split by the boundary, but the validation points are not.
A common solution to how to handle that kind of issue that you might see in
introductory ML courses is to try using a less powerful model. A less powerful model
in this case might come up with a less "wiggly" line to separate the two categories,
perhaps because it didn't have enough parameters to make it wiggle so much, so you
might find that it came up with a classifier that looked more like this:
So: we use our validation set to detect overfitting, and we can adjust the complexity
of our model to try to avoid it.
Now, this is all very well, but it does require manual intervention. We had to do
a training run, identify that we were overfitting, and then decide on parameters for
the new simpler model (how many parameters should it have?). We could, perhaps have
gone too far and wound up with something like this:
There's no way when we start out knowing what the right number of parameters is,
so we need to try various values and then try to work out the optimum balance.
Regularisation techniques are designed to try to automate this -- to prevent overfitting
without all that tedious mucking about with the model.
We've already looked at Dropout , which is one
of the standard ways to do that. Although my own mental model
of what it does goes some way beyond just helping to prevent overfitting, I may well be
wrong -- and given that our LLM train is never seeing the same training data twice, being
a single-epoch run, removing it turned out to improve our model .
Another technique is just stopping the training run when you start seeing the validation
loss rise, also known as "early stopping". That's such an obvious thing to do that I
came up with it independently back when I was doing my early experiments with fine-tuning .
Now, we don't have a separate validation set for these training runs, but because we're doing
a single epoch, the training data it sees is just as "new to it" as a held-back
validation set would be, so we could use a similar trick and treat "train loss starts rising" instead of validation loss
rising as a reason to stop the train early. It's not exactly the same thing, but
perhaps it would be close enough.
But in all of the trains in this series,
that's never happened -- while sometimes the train loss blips up for a bit, in the longer term
it keeps going down.
But there are other techniques that rely on a neat trick. Let's think back to the
manual, boring way of trying to find how many parameters are appropriate for a modelling
task. We tried one number, found that it overfit, then we might try a lower one, find that
it underfit, then try something in the middle and find that it's better but still not
perfect one way or the other, and rinse and repeat until we find something we're happy with.
This kind of searching through a solution space to find an optimum is exactly what we're doing
when training a model. It would be really nice to automate it in the same way.
One trick is: if we want to minimise the complexity of our model so that it doesn't overfit,
we can try adding a measure of the model's complexity to the loss function -- and then
our normal process of gradient descent will try to minimise that, just like it will
try to minimise the loss from the training results themselves. And that brings us
on to weight decay.
Regularisation by weight decay starts off with the hypothesis that the "size" of all
of the model's weights, taken together, is a measure of the model's complexity. If
the model's weights are small, then it's a simpler model than if they're large. 1
The "size" in this sense is the square of the L2 norm -- that's something we came across in
gradient clipping .
The L2 norm is basically all of the weights squared, added together and then the resulting sum
square-rooted. You can think of it as the length of the vector that the weights
represent -- that is, for our 163M-parameter model, it would be the length of the
model's weights considered as a vector in 163-million dimensional space. 2 And by
using its square, we get something that penalises larger values more (and we
also save the time in calculating a square root).
To me, it's not intuitively obvious that that measure really does express
the complexity of the model in any clear sense. After all, you'd think that doubling
all parameters would leave it no more complex than it was before, but it would double
the L2 norm. 3 But I imagine there is solid maths behind it to say that it does work in
a more general way, so in the interests of not disappearing down a mathematical rabbit
hole at this stage, I'll take it as given.
So: we're using the squared L2 norm as a measure of model complexity, and we're going to
add that on to the training loss as a way to try to minimise both. The next question
is, how do we balance between the two -- the training loss and the model complexity
penalty?
This is, in a somewhat hand-wavy way, similar to the decision of how much of the current
loss function's gradient to use when adjusting the weights. For that, we use η , the
learning rate to scale the gradients before applying them:
And the balance between the "real" loss and the model complexity penalty is done in
a similar way -- we have a number, the weight decay, normally represented
by a lower-case lambda, λ , and we multiply the squared L2 norm by that, something like this:
...where I'm using ℒ for the normal loss on the training inputs vs
the targets, N 2 for the squared L2 norm of the weights, and ℒ ′ for the combined
loss. And ℒ ′ is what we -- in theory -- actually try to minimise using our optimiser.
But there's actually a neat simplification that we can apply to make this even easier.
Firstly, let's make one small change to the equation above: we'll halve the
squared L2 norm before multiplying it by λ . That obviously doesn't change the underlying maths, it just
means that we'd need to use larger values for λ to get the same effect. You'll see why that's useful in a bit.
Now let's think about normal gradient descent. Again, we work out the gradient of the loss function for each weight,
and subtract that times the learning rate η from the weight's value to update it:
Let's reformulate that a bit. The gradient of the loss function for the weight is
its partial derivative against that weight, so we can write the above like this
for the version of the loss function including weight decay, ℒ ′ :
Now, we defined ℒ ′ above as ℒ + λ · N 2 2 , so we can
substitute that in there:
Now, let's think about that L2 norm, N . It's
the square root of the sum of all of the weights squared, or equivalently we can
square it (like we do in the formula above) and say:
Now, the derivative of a bunch of things added together is just each of them differentiated
separately and then added together. Let's apply that to the two terms in the brackets:
...and now pull the constant λ and the 2 out of the second partial derivative:
Then we apply the rule for the derivative of a bunch of things added together again:
Now, we're doing a partial derivative versus one specific weight, w , which is one of
the w 0 , w 1 , and so on in there. From that perspective,
all of the other weights are constant -- which means that their derivative
with respect to w is zero. So we can just get rid of all of them apart from the one that
actually is w , and we wind up with this:
The derivative of w 2 with respect to w is just 2 w .
Thanks to that crafty halving of the N 2 earlier, that means that we can go to this:
Multiplying that − η across the bracketed terms, we get:
That's exactly the same as the normal gradient descent update, using the unmodified
loss function without weight decay -- except that we're additionally subtracting
the weight's original value scaled down by both the learning rate η and the
weight decay value λ .
(As an aside: the description above
is correct for "traditional" simple gradient descent and -- loosely -- for Adam, but AdamW's trick is
to do things somewhat differently. That's something I'll go into in more detail when I get round to writing
my post on optimisers.)
So: weight decay is a regularisation technique that tries to prevent our model from
getting any more complex than it needs to be. We have one number, λ , which
determines how much to weight complexity against the normal training loss. And, as we
can see from the code:
optimizer = torch . optim . AdamW (
model . parameters (),
lr = 0.0004 , weight_decay = 0.1
)
...right now we're setting λ to 0.1. Is that the right value?
As usual, the GPT-2 paper
is light on the details of the hyperparameters they used, but nostalgebraist wrote
a really nice post on Tumblr
where they dug into what the number might have been. As they say:
It does say it follows the first GPT paper in most respects, and that paper used weight decay of 0.01.
Their link for the paper appears to be mistaken, as it's a different (albeit very interesting) paper from
2020, a year after the GPT-2 one, but I believe this is the paper normally called the GPT-1 one .
They do indeed use 0.01 there:
We also
employed a modified version of L2 regularization proposed in [37], with w = 0.01 on all non bias or
gain weights.
The link to the GPT-3 paper looks right, though, and as
they say, it uses a weight decay of 0.1:
All models use weight decay of 0.1 to provide a
small amount of r

[truncated]

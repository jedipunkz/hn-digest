---
source: "https://blog.lwarfield.dev/layer-scope/"
hn_url: "https://news.ycombinator.com/item?id=49258121"
title: "Layer Scope: How I used $20 of compute to make a new way to look at LLMs"
article_title: "Layer Scope: How I used $20 of compute to make a new way to look at LLMs"
author: "lwarfield"
captured_at: "2026-08-11T14:14:12Z"
capture_tool: "hn-digest"
hn_id: 49258121
score: 3
comments: 1
posted_at: "2026-08-11T13:37:31Z"
tags:
  - hacker-news
  - translated
---

# Layer Scope: How I used $20 of compute to make a new way to look at LLMs

- HN: [49258121](https://news.ycombinator.com/item?id=49258121)
- Source: [blog.lwarfield.dev](https://blog.lwarfield.dev/layer-scope/)
- Score: 3
- Comments: 1
- Posted: 2026-08-11T13:37:31Z

## Translation

タイトル: レイヤースコープ: LLM を観察する新しい方法を作成するために 20 ドルのコンピューティングをどのように使用したか
説明: LLM の再利用

記事本文:
リアム・ウォーフィールド
レイヤースコープ: LLM を観察する新しい方法を作成するために 20 ドルのコンピューティングをどのように使用したか
LLM の解釈可能性に関する Anthropic の最新の論文、David Noel Ng による優れたブログ投稿シリーズ、および現在取り組んでいるその他の研究に触発されて、LLM ネットワークの真ん中の状態を読み取るためのクールな新しい方法を作成しました。最も良い点は、それが非常にシンプルなアプローチであるということです。
検査する LLM 内のトランスフォーマー ブロックとトークンを選択します。
LLM の最後の 4 層を実行します。
それだけで、Anthropic のほぼすべての読み取り例を再現できます。私はこれをレイヤースコープと呼んでいます。いくつかの例を試してみたい場合は、ここにアクセスしてください。
記事の残りの部分では、私がどのようにしてこの探査のアイデアに到達したかを説明します。また、AI 研究は人々が思っているよりもはるかに取り組みやすいものであることを示したいと思っています。この次の部分は、LLM についてある程度の知識がある SWE 向けに書かれており、本格的な研究論文を目的としたものではありません。
高レベルの LLM アーキテクチャに関する簡単な復習
LLM は、入力に埋め込み行列を乗算して、トークンを表す高次元ベクトルを取得することによって動作します。このベクトルは、隠れ状態として知られています。隠れた状態は、レイヤー [1] 1 とも呼ばれる一連のトランスフォーマー ブロックを介して供給されます。「レイヤー」という言葉が MLP/フィードフォワード レイヤーとも重複しているのは迷惑です。この投稿でレイヤーについて言及するときは、「トランス ブロック」を思い浮かべてください。 。これらの層は隠れ状態を取得し、それにいくらかの Δh を追加します。最後に、非埋め込み行列を乗算して、次のトークンの最終出力分布を取得します。 [2] 2. へー、私が今説明した内容をすべてまとめて本が書けるかもしれません…まだの方は、Welch Labs をチェックしてみてください!
AI/ML における未解決の課題は、どのように導入するかを考えることです。

隠れた状態がモデル内を流れるときに解釈します。 [3] 3. 要出典 何千もの数字のリストは読み物としては不十分であることがわかりました。それでは、研究者がこれまでにこれを行った方法をいくつか見てみましょう。
研究者が層間の隠れ状態を調査しようとした最も初期かつ最も簡単な方法の 1 つは、隠れ状態を非埋め込み行列に直接投入することでした。このアプローチはロジット レンズとして知られるようになりました。
ロジットレンズは安価でシンプルで理解しやすいものの、奥の層に行くほどバラバラになってしまいます。最後の数層の調査は問題なく機能しますが、隠れ状態は以前の層では根本的に異なる構造になっているようです。これにより、ロジット レンズはランダムなガベージの出力を開始します。
これらの中間層の状態から、非埋め込み行列にとって理解可能な状態に移行する方法があれば…ほら、J レンズと呼ばれるものについての Anthropic による論文です。
J レンズはヤコビアン レンズの略です。では、ヤコビアンとは何でしょうか？
ヤコビアンは、ベクトル値関数のすべての一次偏導関数の行列です。
わかりやすく言えば、関数に渡すベクトルの各値の導関数のリストです。これにより、入力ベクトルを少し調整するだけで関数の出力がどのように変化するかがわかります。
それで、Anthropic はここで何をしているのでしょうか?
(A) Jℓ は、最終層の残差ストリームから hℓ に逆伝播し、得られたヤコビアンをトークン位置およびプロンプトのコーパスにわたって平均することによって計算されます。
これを考える 1 つの方法は、特定の層の隠れ状態ベクトルの各部分に関して、次のいくつかのトークンに対するモデルの出力の導関数を取得しているということです。この導関数は次のトークンのいくつかに対して取得され、平均化されて ID が取得されます。

物事が将来のトークンにどのような影響を与えるかについて。
結果として得られるヤコビアンは、この特定のテキストのこの特定のトークンには非常に役立ちますが、他の場所で使用しようとすると壊れる可能性があります。クッキーのレシピでヤコビアンを計算した場合、意味的に似ているため、ケーキのレシピにも役立つ可能性があります。 SF ストーリーで試してみると、ゴミが発生する可能性があります。そのストーリーは、クッキングによって適切に表現されない、隠れた状態のベクトル空間の別の部分に存在するためです。
J レンズがこれを回避しようとする方法は、さまざまなテキスト (料理、詩、数学など) からの約 1000 の異なるプロンプトのヤコビアンを計算し、それらを平均することです。希望 [4] 4. これは大きな仮定です。 ML 研究には一般化されていない長い歴史があり、これがどこに破綻するかを調査する研究プロジェクトを行うことができます。これは、特定の層の隠れた状態に基づいて物事が進んでいる一般的な方向を示す、より一般的なヤコビアンを生成するということです。
全体的にこれは高いです！基本的に、これらのヤコビアンごとに複数のバックプロップ ステップを実行します。具体的には、おおよそ次のようになります。
ヤコビアンあたり 4 バックプロップ ステップ ∗ 1000 ヤコビアン ∗ 32 レイヤー ∗ 4 モデル 2 = 〜 256 , 000 バックプロップ [5] 5. レイヤーのみをバックプロップするため、これは 2 で割られます。
現在いるレイヤーの前にあります。これは平均して完全なバックプロップの数の半分になります。
これには、私がレンタルしていた H100 (コンピューティングの約 9 ドル) で約 4 時間半かかりました。素晴らしいのは、ヤコビアンを平均化すると、最終的な行列をほぼ無料で再利用し続けることができることです。
読み取り値を取得するには、隠れ状態を取得し、ヤコビアンを乗算して正規化し、結果を非埋め込み行列に渡します。これについて簡単に考えると、行列 m

乗算によって方向が決まり、正規化によってベクトルの長さがテキストの通常のものに戻ります。最後に、結果が非​​埋め込み行列によって読み取られて、トークンの分布が取得されます…人間が判読できるものです。
J レンズでの書き込みはこれらの派生関数を使用して行われます。非表示状態を変更して J レンズの強度を下げることにより、モデルを特定の出力から遠ざけることができます。
このヤコビアンを取得すると、2 つの重要なことがわかります。
非表示の状態をよりわかりやすくする方法。隠れ状態にヤコビアンを乗算して、どの出力トークンが現在の状態によって大きく影響されるかを確認できます。
導関数は、隠れ状態への小さな変更が将来の出力にどのように影響するかを示します。これは、小さな変更によるモデルの将来の動作をガイドする優れた方法を提供します。
これにより、非表示状態に対する読み取りと書き込みの両方の非常にエレガントな方法が提供されます。
それでは、レイヤースコープの背景に移りましょう。
最後の数層に固執するという私の決定は、これまで読んできた文献や、現在行っている他の研究 [6] 6 でよく見てきたパターンに基づいています。いつかそれをここで共有できればと思っています。 。このブログ投稿では、David Noel Ng の素晴らしいブログ シリーズに基づいた簡単なバージョンを紹介します (まだ読んでいない場合は、ぜひ読んでください!)。
多くの研究では、モデルがスケールアップするにつれて、異なる変圧器ブロックが異なる処理を行うようであることが観察されています。この差異は、10B パラメーター マーク付近に現れるスケールの新たな特性であると思われます。信じられないほど簡単に要約すると、最初のいくつかの変換ブロックは、トークンの隠された状態をある種の普遍的な「思考」空間にエンコードしているようだということです。このベクトル空間は d でクエリを行うようです

異なる言語だが内容が似ている言語は、隠れた状態で互いに近くにあります。 1 つの解釈として、最初の数層はモデルの最初に使用された埋め込み行列の拡張であると考えられます。 [7] 7. このブログ投稿の多くの事柄と同様、物事はここで見かけられているよりも複雑です。一部の論文では、初期から中期の層にも多くの事実の知識が保存されていることが示唆されています。
すべてがこのより普遍的な潜在空間にあるため、中間層が街に行きます。これらの隠れた状態を操作し、微妙に計算を実行しているようです。これらのレイヤーをいじると、興味深い結果が得られます。いくつかのレイヤーを繰り返すと、モデルの数学や推論がより得意になる可能性があります。ここでいくつかのレイヤーを削除すると、詩を表現する機能を残したまま、論理問題を処理するモデルの機能が破壊される可能性があります。
最後のいくつかのレイヤーは、出力トークンをエンコードしてクエリの元の言語に戻しているようです…または、それが何が起こっているのかについての少なくとも 1 つの理論です。
これらのアイデアをすべてまとめて仮説を立ててみましょう。
LLM の内部状態をデコードするのは非常に困難です。現在の技術は、安価ではあるが限定的なものから、より良い結果が得られるものの一般性に疑問があるものまで多岐にわたります。
LLM には包括的な構造があります。
初期の層はアイデア空間にエンコードされます。
中間層は、そのアイデア空間にあるものに多くの調整を加え、隠れた状態について計算を行います。
最後のいくつかの層は、アイデア空間をデコードして言語に戻します。
この時点で私は次のように尋ねました。
LLM の既存の構造を再利用して中間層をデコードできますか?
モデルの最後の数レイヤーをホッチキスで留めるほど簡単ではないでしょうか?きっと何かが壊れるでしょう！
この枠組みを考えると、LLM がすでに内部に自己分析するためのツールを持っていると考えるのは、かなり小さな飛躍です。

隠れた状態です。最後のいくつかのレイヤーは、すべてのトークンに対してこれを実行します。それでは、これらの最終レイヤーを非表示状態で実行するスコープを作成しましょう。
このプローブには非常に優れた点がいくつかあります。
数千のヤコビアンを計算するよりもはるかに少ない計算量で済みます。 [8] 8. 公平を期すために言うと、J レンズは一度入手すれば再利用可能な単一の行列乗算です。なので、長い目で見れば安くなります。
基本的に VRAM は 0 です。これらのレイヤーの重みはすでに GPU にロードされているため、再利用できます。追跡する必要があるのは、KV キャッシュ内の単一行だけです。
後ろのフロンティアラボのリソースを必要とせずに、これをテストできます。 :D
ここでの主な欠点は、このアプローチでは J レンズのように何かを書くことができないことです。そうは言っても、それでも多くの結果が得られますし、同様のものを得る他の方法があるかもしれません (わずか 4 層でヤコビアンを計算するのはかなり安価です!)。
そこで私は週末をかけてこれを現実化し、Anthropic 論文の例と照らし合わせてテストしました。簡単なことから始めましょう。以下は、中間層で物事を観察する論文の例の 1 つです。
この例では、Sonnet の中間層でトークン「orange」を検出します。 Anthropic のウェイトにアクセスできません [9] 9. 共有したい場合はお知らせください。 , そこで、オープンソースの Qwen モデルと Gemma モデルを使用して分析を行いました。また、私は論文で言及されている 1000 プロンプト ヤコビアン チューニングを持っていないので、私の J レンズはまったく同じではありません。そこでプロンプトを再現してみることにしたところ、「リンゴ」、「ジューシー」、「柑橘類」、「オレンジ」がスコープで浮かんでいることがわかりました。
レイヤースコープの観察
両方をしばらくいじってみた結果、これらの内省ツールについていくつかの結論に達しました。
どちらも

レイヤースコープと J レンズは、Anthropic 論文に示されている読み取り例を拾うことができます。
「柑橘類について考える」の例では、レンズは非常に予測可能なトークンでのみ柑橘類を表示しているように見えます。これは、J レンズと私のレイヤー スコープの両方の少数のトークンでのみ発生するようです。これらのトークンの多くについて、モデルは「古い絵が壁に曲がって掛かっていた」を繰り返すことに焦点を当てています。
全体として、J レンズはサンプル全体でより一貫して物事を捉えているように見えますが、完璧ではありません。たとえば、レイヤー スコープは数学の例では非常に苦労しますが、火星の色の例では J レンズよりも優れています。どうしてだろう、私には分からない！ [10] 10. この種の研究を始める場合は、これをよく考える/言うことに慣れてください。
使用するエンドレイヤーの数を調整するのは少しコツが必要です。使用するレイヤーが少なすぎると、出力が劣化し始め、解釈しにくくなります。私の推測では、要求するレイヤーが少なすぎるため、ランダムなゴミが排出されるのではないかと思います。ロジットレンズと同じように故障します。レイヤーを追加しすぎると、再び通常の LLM のように見え始めます。これらの中間状態を確認できなくなり、モデルは次のトークンを予測するだけになります。柑橘類の例では、オレンジ色の痕跡が失われ始めます。
レイヤースコープはクールですが、これに代わるものではないと思います

[切り捨てられた]

## Original Extract

Reusing an LLM

Liam Warfield
Layer Scope: How I used $20 of compute to make a new way to look at LLMs
Inspired by Anthropic’s newest paper on LLM interpretability, an excellent blog post series by David Noel Ng, and other research I’m currently working on, I’ve created a cool new way to read the state in the middle of LLM networks! The best part is that it’s a wonderfully simple approach:
Choose a transformer block and token in an LLM you want to inspect.
Run the last 4 layers of the LLM.
With just that I’m able to replicate almost all of Anthropic’s reading examples! I call it the layer scope . If you want to play around with some examples go here .
The rest of the article is here to explain how I arrived at the idea for this probe. I also want to show that AI research is something that’s a lot more approachable than people think. This next part is written more for SWEs who know some things about LLMs and is not intended to be a serious research paper.
A Quick Refresher on high-level LLM architecture
LLMs operate by multiplying your input by an embedding matrix to get a high-dimensional vector representing your token. This vector is known as the hidden state . The hidden state is then fed through a set of transformer blocks, also referred to as layers [1] 1. It’s annoying that the word “layer” also overlaps with MLP/feed-forward layers. Whenever I mention layers in this post, just think “transformer block”. . These layers take the hidden state and add some Δh to it. Finally, at the end you multiply by an unembedding matrix to get the final output distribution for the next token. [2] 2. Huh, you could write a book on all the stuff I just glossed over… Go check out Welch Labs if you haven’t!
An open challenge in AI/ML is figuring out how to interpret the hidden state as it flows through the model. [3] 3. citation needed It turns out that lists of 1000s of numbers make for poor reading material! So let’s go through some of the ways researchers have done this in the past.
One of the earliest and simplest ways that researchers tried to probe the hidden state between layers was to throw the hidden state directly into the unembedding matrix. This approach became known as the logit lens .
While cheap, simple, and easy to understand, the logit lens falls apart the farther back in the layers you go. Probing the last few layers works alright, but the hidden state seems to be structured fundamentally differently in the earlier layers. This causes the logit lens to start outputting random garbage.
If only there were a way to move from those middle layer states to something that was intelligible to the unembedding matrix… Hey look, a paper by Anthropic about something called the J lens!
The J Lens is short for Jacobian lens. So what’s a Jacobian?
A Jacobian is a matrix of all first-order partial derivatives of a vector-valued function.
In plainer English, it’s a list of derivatives for each value of a vector that you pass into a function. This gives you an idea of how small tweaks to the input vector will change the output of the function.
So what is Anthropic doing here?
(A) Jℓ is computed by backpropagating from the final-layer residual stream to hℓ and averaging the resulting Jacobians over token positions and over a corpus of prompts.
One way to think of this is that we are taking a derivative of the model’s output for the next few tokens with respect to each part of the hidden state vector at a specific layer. This derivative is taken for several of the next tokens and averaged together to get some idea of how things affect future tokens.
The resulting Jacobian is really useful for this specific token in this specific text, but would likely fall apart if we tried to use it elsewhere. If we calculated the Jacobian on a cookie recipe, it might be useful for a cake recipe since they are semantically similar. Try it on a sci-fi story and you might get garbage since the story lives in another part of the hidden state’s vector space not well represented by cooking.
The way the J lens tries to get around this is by calculating the Jacobian for around 1000 different prompts from a large variety of texts (cooking, poetry, math, etc.), and then averaging them together. The hope [4] 4. This is a big if. ML research has a long history of things not generalizing, and you could do a research project looking into where this breaks down! is that this results in a more general Jacobian that will tell you the general direction things are going based on the hidden state at a particular layer.
Overall this is expensive! You’re basically doing multiple backprop steps for each of these Jacobians. Specifically that ends up with roughly:
4 backprop steps per Jacobian ∗ 1000 jacobians ∗ 32 layers ∗ 4 models 2 = ∼ 256 , 000 backprops [5] 5. This is divided by two because you only backprop the layers
ahead of the layer that you are on. This averages out to half the number of full backprops.
This took about four and a half hours on the H100 I was renting (about $9 of compute). The nice thing is that once you’ve averaged the Jacobians you can keep reusing that final matrix for almost free.
To get a reading you take the hidden state and multiply it by the Jacobian, normalize it, and pass the result through the unembedding matrix. An easy way to think about this is that the matrix multiplication gives you a direction, and normalization gets the vector length back to something normal for text. Finally, the result is read out by the unembedding matrix to get a distribution of tokens… something human readable!
Writing with the J lens is done using those derivatives! You can steer the model away from certain outputs by modifying the hidden state to lower their strength in the J lens.
Once we have this Jacobian we get two key things:
A way to make the hidden state more understandable. We can multiply the hidden state by the Jacobian to see what output tokens are highly affected by the current state.
The derivatives tell us how small changes to the hidden state affect future output. This gives a nice way to guide future behavior of the model with those small changes.
This gives a wonderfully elegant way of both reading and writing to the hidden state!
With that, let’s move on to the background for the layer scope.
My decision to stick with the last few layers is based on a pattern that I’ve seen pop up a lot in the literature that I’ve been reading, and the other research I’m currently doing [6] 6. hopefully I’ll be able to share that here some day! . For the purpose of this blog post I’m going to give a brief version based on David Noel Ng’s wonderful blog series (Please give those a read if you haven’t!).
A lot of research has observed that different transformer blocks seem to handle different things as a model scales up. This differentiation seems to be an emergent property of scale that shows up around the 10B parameter mark. The incredibly brief summary is that the first few transformer blocks seem to encode the hidden state of the tokens into some type of universal “thought” space. This vector space seems to make queries in different languages but with similar content land near each other in the hidden state. One interpretation might be that the first few layers are an extension of the embedding matrix used at the beginning of the model. [7] 7. As with many things in this blog post, things are more complicated than they seem here. Some papers suggest the early-mid layers also store a lot of factual knowledge.
Now that everything is in this more universal latent space, the middle layers go to town. They seem to operate on these hidden states and subtly perform computation on them. Messing around with these layers gives some interesting results. Repeating some layers might make the model better at math or reasoning. Dropping some layers here might destroy the model’s ability to handle a logic problem, while leaving the ability to do poetry.
The last few layers seem to encode the output token back into the original language of the query … Or that’s at least one theory about what’s going on.
Let’s put all of these ideas together and create a hypothesis:
Decoding the internal state of an LLM is really hard. Current techniques range from cheap but limited, to better results with questionable generality.
LLMs have an overarching structure to them.
The early layers encode into idea space.
The middle layers make a bunch of tweaks to things in that idea space, doing some computation on the hidden state.
The last few layers decode from idea space back into language.
It was at this point that I asked:
Can we reuse the existing structure of the LLM to decode the middle layers?
It can’t be as simple as stapling on the last few layers of the model? Surely something will break!
With this framing it’s a fairly small leap to think that LLMs already have a tool inside of them to introspect on this hidden state. The last few layers do this for every token! So let’s create a scope that just runs these final layers on the hidden state!
This probe has a few really nice things about it:
It’s way less compute than calculating 1000s of Jacobians. [8] 8. To be fair, the J lens is a single reusable matrix multiplication once you’ve got it. So it’ll be cheaper in the long run.
It basically costs 0 VRAM! The weights for those layers are already loaded into the GPU, so we can reuse them. The only thing we need to keep track of is a single row in the KV cache!
I can test this without needing the resources of a frontier lab behind me! :D
The main drawback here is that this approach does not give us the ability to write things like the J lens can. That being said, we still get a lot, and there might be other ways of getting something similar (calculating a Jacobian on just 4 layers is pretty cheap!).
And so I spent the weekend bringing this thing into reality, and testing it against the examples in the Anthropic paper. Let’s start out with something simple. Here’s one of the paper’s examples of seeing things in the middle layers:
In this example they detect the token “orange” in the middle layers of Sonnet. I don’t have access to Anthropic’s weights [9] 9. let me know if y’all want to share! , so I’ve done my analysis using open-source Qwen and Gemma models. I also don’t have the 1000 prompt Jacobian tuning they mention in the paper, so my J lens is not quite the same. So I decided to replicate the prompt and found “apples”, “juicy”, “citrus” and “orange” floating around with my scope!
Observations of the layer scope
After messing around with both for a while I’ve come to some conclusions about these introspection tools:
Both the layer scope and the J lens are able to pick up on the reading examples shown in the Anthropic paper.
In the “think about citrus” example, the lenses only seem to show citrus on very predictable tokens. This also seems to only crop up in a small number of the tokens for both the J lens and my layer scope. For many of these tokens, the model is focused on repeating “The old painting hung crookedly on the wall.”
Overall the J lens seems to be more consistent at picking things up across the examples, but it’s not perfect. For example: the layer scope really struggles with the math example, but does better than the J lens on the color of mars example. Why, I have no idea! [10] 10. Get used to thinking/saying this a lot if you start doing this type of research.
Tuning the number of end layers to use is a bit of an art. If too few layers are used, the output starts degrading and becomes less interpretable. My guess is that you get random garbage out because it’s asking too much of too few layers. It breaks down in the same way as the logit lens . Add too many layers and it starts looking like a normal LLM again. You lose the ability to see those middle states, and the model starts to just predict the next token. In the think of citrus example, you start losing the orange traces.
Layer scope is cool but I don’t think this is a replacement

[truncated]

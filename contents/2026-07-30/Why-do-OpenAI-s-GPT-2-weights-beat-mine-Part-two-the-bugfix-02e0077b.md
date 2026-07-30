---
source: "https://www.gilesthomas.com/2026/07/why-do-openai-gpt2-weights-beat-mine-2-the-bugfix"
hn_url: "https://news.ycombinator.com/item?id=49113601"
title: "Why do OpenAI's GPT-2 weights beat mine? Part two: the bugfix"
article_title: "Why do OpenAI's GPT-2 weights beat mine? Part two: the bugfix :: Giles' blog"
author: "gpjt"
captured_at: "2026-07-30T19:10:46Z"
capture_tool: "hn-digest"
hn_id: 49113601
score: 2
comments: 0
posted_at: "2026-07-30T18:13:31Z"
tags:
  - hacker-news
  - translated
---

# Why do OpenAI's GPT-2 weights beat mine? Part two: the bugfix

- HN: [49113601](https://news.ycombinator.com/item?id=49113601)
- Source: [www.gilesthomas.com](https://www.gilesthomas.com/2026/07/why-do-openai-gpt2-weights-beat-mine-2-the-bugfix)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T18:13:31Z

## Translation

タイトル: OpenAI の GPT-2 重みが私のものよりも優れているのはなぜですか?パート 2: バグ修正
記事のタイトル: OpenAI の GPT-2 重みが私のものよりも優れているのはなぜですか?パート 2: バグ修正 :: Giles のブログ
説明: 手順に従ってモデルをオリジナルの GPT-2 モデルと同じくらい良くする試みを開始する前に、修正する必要のあるバグがありました。

記事本文:
OpenAI の GPT-2 重みが私のものを上回るのはなぜですか?パート 2: バグ修正 :: Giles のブログ
el.dataset.currentDropdown = '')
}">
ジャイルズのブログ
何かを学び始めたときに見つけたかった投稿を書いています...
について
お問い合わせ
アーカイブ
カテゴリー
ブログロール
2026年7月 (6)
NSLU2 オフサイト バックアップ プロジェクト (13)
ソフトウェア開発ツール (1)
:: (ブログ可能 a) => a -> IO ()
いくつかの意見は、さまざまな程度の確実性を持って保持されています
OpenAI の GPT-2 重みが私のものを上回るのはなぜですか?パート 2: バグ修正
私は、GPT-2 スタイルのモデルが命令後の評価で OpenAI の元の重みよりも悪いスコアを示す理由を詳しく調べています。
この投稿で詳細を説明しました。
最初の実験の結果を考えられる原因として書き留めているときに、
ChatGPT を超えて投稿を実行しました -- 私は常に AI の「編集ボード」を使用して自分の投稿をチェックしています
フロー、スタイル、技術的なエラーについての投稿 (ただし、すべての文章は常に私によるものです)。
私が実行していた eval コードを調べて、バグを指摘しました。
幸いなことに、重要な結果は変わりません -- OpenAI のモデルは引き続き使用されます。
指示に従うのは私より上手です。しかし、そうでした
ベースラインの数値を変更して、自分のモデルのパフォーマンスを並べ替えるだけで十分です。
そこでそれを修正してベースラインを再生成しました
今後の実験がしっかりした基盤に基づいて行われるように。
評価ではモデルを取得し、Alpaca のサブセットの分割で複数のエポックにわたってモデルをトレーニングします。
命令に従うデータセット。各エポックの終わりに、次の値を評価します。
保留された検証分割に対する結果のモデル。評価損失が始まった場合
上昇すると、救済されます。最後に、結果として得られるデータセットのテスト分割を実行します。
モデルを作成し、結果を保存します。
さまざまなモデルに対して実行したら、LLM-as-a-judge スクリプトを使用します。
GPT 5.5 を取得して結果をスコアリングする -- 各質問回答結果に対して

、すべてが見えます
同じプロンプト内のすべてのモデルの応答を毎回順番にシャッフルし、
可能な限り一貫してモデルを相互に判断できるようにするためです。
さて、テストの分割回答の生成には、前の時代のモデルを使用するという考えがありました。
上昇損失のものへ。だから私は次のようなコードを持っていました：
範囲 ( 100 ) のエポックの場合:
モデル。電車（）
tqdm の input_batch 、 target_batch の場合 ( train_loader 、 desc = f "Epoch { epoch } " ):
オプティマイザー。ゼログラッド()
損失 = calc_loss_batch (
input_batch 、 target_batch 、モデル、デバイス
）
紛失。後ろ向き（）
オプティマイザー。ステップ（）
モデル。 eval()
val_loss = calc_loss_loader ( val_loader 、モデル、デバイス、 eval_iter )
last_val_loss が None または last_val_loss > val_loss の場合:
last_val_loss = val_loss
last_params = モデル 。 state_dict ()
print ( "Val 損失はまだ減少し続けています" )
それ以外の場合:
print ( "Val 損失が上昇、救済" )
休憩
モデル。 load_state_dict ( last_params )
検証損失が減少するたびに、モデルのパラメータを保存したいと思いました
last_params にあるため、最後の行の後半で復元できるようになります。
よく見ると、間違いは明らかです。 model.state_dict() が返されない
モデルのパラメータのコピー -- 参照を含む辞書を生成します
モデル内のパラメータに適用されます。それで、パラメータを隠し込もうとしていましたが、
last_params には損失が減少するたびに、そのコピーが保存されていました。
上昇と損失の時代より前の時代に生きていましたが、私が実際にやっていたのは
「ライブ」パラメータへの参照を無意味に保存しているだけです。への呼び出し
model.load_state_dict(last_params) は基本的に何も操作しませんでした。
解決策は非常に簡単でした。
コピーインポートディープコピーから
...
last_params = deepcopy (モデル . state_dict ())
コードを実行するにはこれで十分でした

それが何をするつもりだったのか。
そこにいる間、評価コードが最初のコードのみを使用していることにも気づきました。
私の評価データセットには 5 つのバッチがあります。このコードはもともと eval から適応されたものです
「大規模な言語モデルを構築する (ゼロから)」、
ここでは評価がより頻繁に実行され、超高速に実行される必要がありました。なぜなら
私自身のコードではそれが実行されることはほとんどなく、すべてを使用する方が合理的でした。
完全に再実行するつもりだったことを考えると、
テスト応答を生成するスクリプトを修正した方がよいのではないかと考えました。
同時に。
比較しているすべてのモデルで修正されたスクリプトを再実行し、過去に実行しました
GPT 5.5 審査員。これが私が得たものです:
まずトレーニング エポックの数を調べてみましょう。変更されたモデルを強調表示しました。
「Cloud FineWeb、8x A100 40 GiB」モデルの場合は、変更の結果だと思います
検証サンプルの数に応じて。
2 つの JAX モデルについては、もう少し謎が多くなります。という部分もあるかもしれませんが、
同じ「more eval data」のバリエーションもありますが、何かがあるのではないかと疑っています。
さらにドロップアウトに関連して。
これまでの私の方法論は、IFT トレーニングを実行するというものでした。
元のベース モデルのトレーニング実行と同じドロップアウト設定を使用する必要があります (
これについては、後の投稿でもう一度説明します)。ただし、JAX 間の違いにより、
ドロップアウトに一致するトレーニング スクリプトと評価コード (PyTorch)
これらのモデルを評価する際のレートは少し面倒でエラーが発生しやすくなります。
私はそれを100%確信しています
今回は正しく理解できました -- 特定のコマンドを何度も確認しました
私は走りました。
しかし、私がやったことを覚えている限りでは、次のように思います。
前回は失敗したかも知れません。
ただし、それは確実ではなく、半分の記憶に基づいた単なる推測です
数週間前に実行したいくつかのコマンドは、私の .bash_hi には反映されませんでした。

物語
(一度に開く端末が多すぎます)。
IFT スコアについては、実行間で厳密に比較できないことに注意してください。
LLM 裁判官が「高慢と偏見の著者は誰ですか?」という質問に対して次のような答えを与えられたと想像してみましょう。
「『高慢と偏見』の作者はジェーン・オースティンだった」
「『高慢と偏見』の作者はサラ・ペイリンだった」
「『高慢と偏見』の作者は『高慢と偏見』だった」
場合によっては、最初の 2 つを 100/100 として扱う場合もあります。
2番目は冗長すぎるという理由で95/100点を与えるかもしれません。同様に、場合によっては、
最後の 2 つは間違っているとして 0/100 としてランク付けします。他のものでは「サラ・ペイリン」が与えられる可能性があります。
完全なナンセンスではなく、少なくとも人の名前であるという点で 5/100 が 1 つあります。
さて、私たちは常に LLM に、与えられた質問に対するすべてのモデルの回答を判断するように依頼します。
1 回の実行で、少なくとも、次の点についての特定のプロンプトに対して一貫性があることを確認できます。
与えられた質問。しかし、一貫性を保つことができないのは、それがどちらに傾いているかということです。
異なる実行、または同じ実行内の異なる質問。場合によってはそうかもしれない
「寛大な」と感じて、サラ・ペイリンに少しだけ猶予を与えて答えてください。
もっと厳しいかもしれない。
したがって、そこにはかなりの量のノイズが存在します。私の経験則では、次のバリエーションがあります。
1 つまたは 2 つのポイントがそのノイズ内にあるため、OpenAI 媒体は 41.62 から 42.41 になります。
ほとんど意味がなく、同様に「JAX、MHA バイアスあり、ドロップアウトなし」の動きも
19.25から18.12まで。
したがって、重要なのは相対的な順位です。どれが 1 位で、どれが 2 位でしょうか。
など。当然のことながら、たとえば、次のような事実を考慮する必要があります。
モデルは、「JAX、MHA バイアスなし、ドロップアウトなし」のように、位置 11 から位置 3 になります。
した場合、前の番号 3 は番号 4 になる必要があり、4 は番号 4 になります。
5になるなど。それがわかります

結果テーブルで発生しています。明らかに、
他のモデルが上がったり下がったりするにつれて、さらなる波及効果が生じます。
とにかく、これらすべての注意事項を踏まえた上で、良いニュースは、私の最初の謎が残っているということです。
OpenAI モデルは依然として、私が作成したモデルよりも明らかに優れたパフォーマンスを示していました。 GPT-2中
引き続き群をリードし続けました（大型モデルであることを考えると当然のことですが）、
GPT-2 smallは依然として2位でした。それが変わっていたら、
このシリーズのかなり残念な終わり方: 「謎の説明がつきましたが、それはシステムのバグでした」
評価:-("
さて、私のモデルを見てみましょう。
まず、「MHA バイアスなし」モデルが恩恵を受けているように見えました。
追加のトレーニングによるもの、またはドロップアウト設定の修正によるものです。彼らはから立ち上がった
11 位と 15 位がそれぞれ 3 位と 8 位になりました -- 「JAX、MHA バイアスなし、ドロップアウトなし」にとっては大きな変動です。
そして、「JAX、MHA バイアスなし、ドロップアウトあり」に関しては確実な改善が見られました。
相対ランキングにおけるその他の変化のほとんどは、これら 2 つのモデルで説明できます。
昇格しましたが、他にもいくつか変更があります。特に3つのモデルは、
スコアが大幅に低下しました (結果として、ランキングも低下しました)。
1xrtx3090-スタック介入
「Cloud FineWeb、8x B200 160 GiB」
私の疑念は、非常に根拠の薄い仮説ですが、興味深い仮説です。これらのモデルは以前からバグの恩恵を受けていたのではないかということです。
トレーニングを停止するために使用しているシグナルは、検証の損失であることを思い出してください。
上昇し始めます。これを過学習のプロキシとして使用し、次に次のように使用しています。
「このモデルは十分なトレーニングを受けています」の代用
評価に必要なようにこのデータを使用します。」
ただし、接続が確立されているという保証はありません。おそらくモデルはそうではなかった
過学習であり、もう 1 ～ 2 エポック待っていたら、検証損失が発生した可能性があります。
また落ち始めた。あるいは、ある程度の過剰適合があるかもしれません

tingはこれに役立つでしょう
評価？
おそらく、ここでは無限に近い量の掘り下げ作業ができる可能性があります。
でもやめたほうがいいと思います。バグ修正は重要でした。
eval は、私がやっていると思っていたことをやっていました。
重要なのは、この不可解さは変わらないということです
実際、私のモデルはこの評価において OpenAI のモデルよりも悪かったので、それを解明しようとしているのです。
そしてそれは、私がより自信を持ってベースラインの数字に頼れるようになったことを意味します。
それで、今度は、ギャップを縮めることができるかどうかを確認するために、実際に変更を開始するときです。

## Original Extract

Before I get started on my attempts to make my models as good as the original GPT-2 ones at instruction-following, there was a bug I needed to fix.

Why do OpenAI's GPT-2 weights beat mine? Part two: the bugfix :: Giles' blog
el.dataset.currentDropdown = '')
}">
Giles' blog
Writing the post that I wished I'd found when I started learning whatever it was...
About
Contact
Archives
Categories
Blogroll
July 2026 (6)
NSLU2 offsite backup project (13)
Software development tools (1)
:: (Bloggable a) => a -> IO ()
Some opinions, held with varying degrees of certainty
Why do OpenAI's GPT-2 weights beat mine? Part two: the bugfix
I'm digging into why my GPT-2 style models score worse on an instruction-following eval than OpenAI's original weights;
I gave the details in this post .
While I was writing up the results of my first experiment into possible causes,
I ran the post past ChatGPT -- I always use an "editorial board" of AIs to check my
posts for flow, style, and any technical errors (though all writing is always mine).
It took a look at the eval code that I was running, and highlighted a bug.
Luckily, it doesn't change the important results -- OpenAI's models continue to be
better than mine at instruction-following. But it was
enough to change the baseline numbers, re-ordering how well my own models did.
So I fixed it and regenerated the baseline
so that future experiments are based on solid ground.
The eval takes a model, and trains it over multiple epochs on a split of a subset of the Alpaca
instruction-following dataset. At the end of each epoch, it evaluates
the resulting model against a held-back validation split; if the eval loss starts
rising, it bails out. Finally, it runs a test split of the dataset through the resulting
model, and saves the result.
Once I've run it for a bunch of different models, I use an LLM-as-a-judge script
to get GPT 5.5 to score results -- for each question-answering result, it sees all
of the responses for all of the models in the same prompt , shuffled in order each time,
to try to make it judge models against each other as consistently as possible.
Now, the idea was that the generation of the test split answers would use the model from the epoch prior
to the rising-loss one. So I had code like this:
for epoch in range ( 100 ):
model . train ()
for input_batch , target_batch in tqdm ( train_loader , desc = f "Epoch { epoch } " ):
optimizer . zero_grad ()
loss = calc_loss_batch (
input_batch , target_batch , model , device
)
loss . backward ()
optimizer . step ()
model . eval ()
val_loss = calc_loss_loader ( val_loader , model , device , eval_iter )
if last_val_loss is None or last_val_loss > val_loss :
last_val_loss = val_loss
last_params = model . state_dict ()
print ( "Val loss still decreasing, continuing" )
else :
print ( "Val loss rising, bailing out" )
break
model . load_state_dict ( last_params )
Each time the validation loss went down, I wanted to store the model's parameters
in last_params so that they could be restored later in that last line.
If you look closely, the error is pretty obvious. model.state_dict() does not return
a copy of the model's parameters -- it produces a dictionary containing references
to the parameters inside the model. So although I was trying to stash away the parameters
in last_params each time loss went down, so that we had a copy of the ones that
were live in the epoch prior to the rising-loss one, what I was actually doing was
just pointlessly saving a reference to the "live" params. The call to
model.load_state_dict(last_params) was essentially a no-op.
The solution was simple enough:
from copy import deepcopy
...
last_params = deepcopy ( model . state_dict ())
That was enough to make the code do what it was meant to do.
While I was there, I also noticed that the evaluation code was only using the first
five batches in my eval dataset. This code was originally adapted from an eval
in " Build a Large Language Model (from Scratch) ",
where the eval was run much more frequently and had to be super-fast. Because
my own code ran it more rarely, it made more sense to use all of it.
Given that I was going to completely re-run
the script to generate the test responses, I figured that I might as well fix that
at the same time.
I re-ran the fixed script on all of the models I'm comparing, and ran that past
the GPT 5.5 judge; here's what I got:
Let's look into the number of training epochs first; I've highlighted the models for which it changed.
For the "Cloud FineWeb, 8x A100 40 GiB" model, I think that was a result of the change
to the number of validation samples.
For the two JAX models, it's a bit more of a mystery. There may be some part of the
same "more eval data" variation there, but I have a suspicion that there's something
more, related to dropout.
My methodology to date has been that the IFT training runs
should use the same dropout setting as the original base model training run (I'll
come back to this in a later post). However, due to differences between the JAX
training script and the evaluation code (which is PyTorch), matching the dropout
rate when evaluating those models is a bit fiddly and error-prone.
I am 100% sure that
I got it right this time around -- I've checked and double-checked the specific commands
I ran.
But I think -- from what I remember doing -- that
I might have messed it up the previous time.
That's not certain, though -- just a suspicion based on half-memories
of some commands I ran several weeks ago, which never wound up in my .bash_history
(too many terminals open at once).
For the IFT scores, remember that they're not strongly comparable between runs.
Let's imagine that the LLM judge is given the following answers to the question "Who was the author of Pride and Prejudice ?"
"The author of 'Pride and Prejudice' was Jane Austen"
"The author of 'Pride and Prejudice' was Sarah Palin"
"The author of 'Pride and Prejudice' was 'Pride and Prejudice'"
In some cases it might treat the first two as being 100/100, in others it
might give the second 95/100 for being too wordy. Likewise, in some cases it might
rank the last two as 0/100 for being wrong, in others it might give the "Sarah Palin"
one 5/100 for at least being the name of a person rather than complete nonsense.
Now, we always ask the LLM to judge all of the models' answers for a given question in
one go, so at least we can be sure that it will be consistent for a given prompt about
a given question. But what we can't keep consistent is which way it leans between
different runs, or different questions within the same run. Sometimes it might be
feeling "generous" and give the Sarah Palin answer a bit of grace, other times
it might be harsher.
So there's a significant amount of noise there; my rule of thumb is that a variation of
a point or two is within that noise, so OpenAI medium going from 41.62 to 42.41 is
pretty much meaningless, and likewise "JAX, with MHA bias, no dropout"'s move from
19.25 to 18.12.
So what's important is the relative ranking -- which is first, which is second,
and so on. Naturally, you have to allow for the fact that -- for example -- if one
model goes from position 11 to position 3, like "JAX, no MHA bias, no dropout"
did, then the previous number 3 will have to become number 4, 4 will
become 5, and so on. You can see that happening in the results table. Obviously,
as other models rise and fall, that has further knock-on effects.
Anyway, with all of those caveats, the good news is that my original mystery remains.
The OpenAI models were still doing noticeably better than my own ones. GPT-2 medium
continued to lead the pack (unsurprisingly, given that it's a bigger model), and
GPT-2 small was still in second place. If that had changed, it would have made
a rather disappointing end to this series: "mystery explained, it was a bug in the
eval :-("
But now let's look at my models.
Firstly, it looked like the "no MHA bias" models might have benefited
from the extra training -- or from having their dropout settings corrected. They rose from
positions 11 and 15 to 3 and 8 respectively -- a huge swing for "JAX, no MHA bias, no dropout",
and a solid improvement for "JAX, no MHA bias, with dropout".
Most of the other changes in relative rankings can be explained by those two models
having been promoted, but there are some other changes. In particular, three models
dropped significantly in score (and, as a result, ranking):
1xrtx3090-stacked-interventions
"Cloud FineWeb, 8x B200 160 GiB"
My suspicion -- a very weakly-held hypothesis, but an interesting one -- is that those models had previously been benefiting from the bug.
Remember that the signal we're using to stop training is that the validation loss
starts rising. We're using that as a proxy for overfitting, which in turn we're using as
a proxy for "this model has had as much training
on this data as it needs for the eval".
But there's no guarantee that the connection is there. Perhaps the models weren't
overfitting and if we'd waited for another epoch or two, the validation loss might have
started falling again. Or maybe some amount of overfitting would be beneficial for this
eval?
There's probably a near-infinite amount of digging in that I could potentially do here.
But I think it's best to stop. The bugfix was important because it meant that the
eval was now doing what I thought it was doing.
Importantly, it doesn't change the puzzling
fact that my models were worse at this eval than OpenAI's, which is what I'm trying to untangle.
And it means that I can now lean more confidently on the baseline numbers.
So now it's time to actually start changing things to see if I can close the gap!

---
source: "https://spock.is/writing/simulating-cosmic-rays-to-lobotomize-llms"
hn_url: "https://news.ycombinator.com/item?id=49389343"
title: "A single well aimed cosmic ray can demolish a LLM"
article_title: "Simulating cosmic rays to lobotomize LLMs — Benedikt Holm"
image: "https://spock.is/writing/simulating-cosmic-rays-to-lobotomize-llms/asset-0ycu4d0.png"
author: "BenediktHolm"
captured_at: "2026-08-21T15:23:58Z"
capture_tool: "hn-digest"
hn_id: 49389343
score: 1
comments: 0
posted_at: "2026-08-21T15:11:22Z"
tags:
  - hacker-news
  - translated
---

# A single well aimed cosmic ray can demolish a LLM

- HN: [49389343](https://news.ycombinator.com/item?id=49389343)
- Source: [spock.is](https://spock.is/writing/simulating-cosmic-rays-to-lobotomize-llms)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T15:11:22Z

## Translation

タイトル: 狙いを定めた単一の宇宙線が LLM を破壊する可能性がある
記事のタイトル: LLM をロボトミー化するための宇宙線のシミュレーション — Benedikt Holm
説明: 私たちの貴重な情報の 490 億ビットは、

記事本文:
← 書く
LLM をロボトミー化するための宇宙線のシミュレーション
ベネディクト・ホルム · 2026年8月20日
私たちの貴重な知性の 490 億ビットでは、1 つの宇宙線の後に 2 つの数字を加算することはできません。
宇宙には宇宙線として知られる災難が存在します。私たちのコンピューターを構成するビットに衝撃を与えると、0 から 1、またはその逆に自発的に反転することがあります。これらのビットフリップは、スピードランナーの勝利を早めたり、ベルギーの政治に干渉したり、ボイジャー 2 号の同期を狂わせたりするのに役立ったと言われていることでよく知られていますが、機械学習モデルを残酷にすることも示されています。現在、イーロン・マスクのような人々は、宇宙にデータセンターを設置する計画を立てていますが、そこではビットフリップは宇宙線だけでなく、はるかに強力な周囲放射線によっても引き起こされるため、海面よりもはるかに一般的です。ここで私は疑問に思いました。LLM は宇宙線攻撃に対してどれほど脆弱なのでしょうか?しかし、あなたが「馬鹿野郎、ECC というものがあるのですが、聞いたことありますか?」と必死に入力しているのがすでに聞こえてきます。はい、そうです。邪魔するのはやめてください。すぐに取り掛かります。
前置きすると、これは科学的な記事というよりは、私が新生児を寝かしつけているときに思いついたきちんとした実験の楽しい小さな記録です。また、以前の研究では LLM のビット破損について調べられていたため、完全に目新しいものではありませんが、私が知る限り、それは主に標的型または敵対的なビット操作に焦点を当てています。また、ここでシミュレートしているのは素粒子物理学、GPU、または実際の放射線環境ではないことも明確にしておきたいと思います。私は、放射線によるソフト エラーのおもちゃのモデルとして、モデル ビットを均一にランダムに反転させています。
ここで説明した作業の多くはさらに進めることができますが、私は時間 (新生児) と計算量の両方に制約があるため、この作業は必然的に制限されています。いずれにせよ、私がここで述べた恐怖から少しでも楽しんでいただければ幸いです。

eモデルを通して。
実験中の私とクウェン。
LLM が宇宙のロボトミー手術にどれだけ耐えられるか実験を行うために、私がとるべき最初のステップは被験者を選択することでした。選択肢を検討した結果、最初は機器をレンタルするのではなく、現地で実験を行うことにしました。これは 2 つのことを意味します。1) ローカル LLM コミュニティのために重要な研究を行っていると主張できます (local.ai を叫んでください)。2) Vast.ai のようなサイトで、高価で高値で入札できるハードウェアに多額の費用を費やす必要がありません。ただし、これは、ラップトップの 5070 に適合するモデルが限られていることを意味します (そうです、私は計算能力が低いのです)。私が最終的に選んだ犠牲者は、プログラミング タスクを解決するために訓練されたコーディング LLM である Qwen2.5-Coder-3B でした。注目すべき点は、FP16 で Qwen を実行したことです。各重みは 16 ビット浮動小数点数です。 16 ビットのうち、14 番は重みの指数の最上位ビット (MSB) であるため、最も魔法的です。次のデジタル フィジェット トイで遊ぶことができるように、この位置でビットを反転すると、その番号が劇的に変化する可能性があります。
= 0.02063 リセット 0 0 1 0 0 1 0 1 0 1 0 0 1 0 0 0
任意のビットをクリックして反転します。ビット 14 (赤いリング) は宇宙線です。0.021 を 1352 に変換します。リセットを押して復元します。
モデルの脳内でビットを反転することが出力にどのような影響を与えるかを測定するために、164 個のコーディング問題のセットである human-eval (HE) ベンチマークを使用しました。 HE は、多くのモデルがトレーニング データに含まれているため、時代遅れであると広く考えられていることに注意してください。私は、クウェンで模擬宇宙レーザーを使用して精度の低い脳手術を行うつもりなので、訓練の内容はそれほど重要ではないと考えました。拷問を開始する前に、基本モデルが 164 個の HE 問題のうち 139 個、つまり約 85% を解決したことを記録しました。このスコアを覚えておいてください、だって

使用しても改善されません。
Qwen を実行した後、メモリ内の重みをアドレス指定するのは比較的簡単でした。Transformers ライブラリを使用すると、個々の重みを直接アドレス指定できるため、特定の重みでビットを反転することは、XOR 演算するのと同じくらい簡単でした。
ここで、x は重み内で反転するビットの位置です。これにより、損傷を迅速かつ簡単に元に戻すこともできました。この実験では、生成内の do_sample フラグを False に設定することで、モデル推論を貪欲に保ちました。これは基本的に、シードや温度が推論に何の役割も果たさず、出力を重みの値の決定論的な関数に変えることを意味します。 tps を最適化するために実行ごとにバッチ サイズを変更しなかった場合も同様です。その結果、いくつかの実行の間でモデルのベースラインにわずかなぐらつきが生じます。念のため、一部の実行がわずかに異なる HE ベースラインで開始していることに気づいた場合に備えて説明します。気付かなかった場合は、この最後の文は忘れてください。
始める前に、モデルをビットごとにザッピングすると永遠に時間がかかるだろうとかなり確信していたので、ビットフリップ テープのようなもの、つまり反転するビット位置のシード シーケンスを設計しました。これにより、どのビットがどのような順序で破損するかを事前に計算できるようになり、毎回 1 ビットを反転してトークンを再生成する代わりに、一度に 500,000 個を破損して HE スコアの劣化をチェックすることができました。劣化した場合、致命的な打撃は 50 万ビットのどこかにあることがわかっていたので、半分にロールバックして再度テストし、まだ壊れている場合はさらにロールバックすることができました。モデルが良好であれば、ダメージを 2 倍にして繰り返すことができます。これは二分検索であり、大きな検索スペースを削減し、時間と計算量を節約するのに非常に貴重ですが、どちらも私には決定的に不足しています。繰り返しの反転に対しては防御しないことにしたので、1 ビットでも可能性があることに注意してください。

1回の実行で複数回ひっくり返される。
私は、最初に検出できる損傷はわずか 200 万ビット、さらには 300 万ビット程度にすぎないと十分に予想していました。しかし驚いたことに、最初の二等分検索が終了したとき、モデルがまったく使用不能になるまでにわずか 6 ビットの反転しか必要でなく、数文字の漢字の繰り返しが際限なく送信されていたことが判明しました。その魔法の 6 番目のビットは、まさに前述した、レイヤー 27 の単一フィードフォワード重みの指数の MSB であり、重みの値を 0.021 から 1352 に変換しました。最後のロボトミー化ビットまで、モデルの出力は (HE にとって) 品質において破損していないモデルと区別できませんでした。
範囲内の i の場合 ( len (数値)):
範囲 ( i + 1 , len (数値)) 内の j の場合:
if abs (数値 [ i ] - 数値 [ j ]) < しきい値 :
Trueを返す
return False !__!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_ … (+403 もっと同じ文字) 戻り値 - int (数値) big_expression[!expression[!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! … (同じ文字をさらに 424 文字追加) 残高 = 0
操作中の操作の場合:
バランス += 操作
残高 < 0 の場合:
Trueを返す
return False !__#_!_."__!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_! … (+404 もっと同じ文字) 数字でない場合:
戻り値 ( 0 , 1 )
total_sum = 合計 (数値)
合計製品 = 1
数字の中の数字の場合:
total_product *= 数値
return ( total_sum , total_product ) !__!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!__!_!_! … (同じ内容のさらに 403 文字) 4 つの HumanEval 問題に関する同じモデル。タブをクリックします。左: クリーンな出力、テストに合格。右

: 6 回のランダムなビット反転の後 (6 番目のビットがビット 14 に着地)、すべての答えは解析すらできないノイズに崩壊します。実際、いくつかの異なるシードを使用して再度実行すると、この脳死は平均してわずか約 20 ビットの反転で発生しました。毎回、最後の致命的なフリップは FP16 の指数の MSB に着地しました。各致命的なフリップを個別に繰り返すと、これらのフリップのそれぞれが単独でモデルを完全にフロアさせるのに十分であることがわかりました。 MSB に対するこの弱点は既存の研究でも十分に文書化されていますが、ここで非常によく再現されているのを見るのは興味深いことです。
0 10 20 30 40 50 60 ランダム ビット 反転の適用 0 20 40 60 80 100 120 140 HumanEval の問題が解決されました シード 4 シード 6 シード 2 シード 3 シード 8 シード 7 シード 5 ランダム ビット 反転はモデルのパフォーマンスに大きく影響します。 Qwen2.5-Coder-3B には 7 つの宇宙線シードがあり、どれも遮蔽されていません。それぞれがフルスコアを保持し、1 回のフリップで終了します。
ビットが反転されると HE スコアが低下します。
しかし、何らかの方法で確率的ディセレブレーションが魔法のようにビット 14 を回避した場合、Qwen は脳のよりフォールト トレラントな部分に何回ザップできるでしょうか?これをテストするために、MSB に到達するはずのレイをスキップして、同じランダム化されたビット破損を実行しました。
ランダムなビット反転が適用され、ビット 14 がシールドされている (千) 0 20 40 60 80 100 120 140 HumanEval の問題が解決されました シード 4 シード 8 シード 2 シード 3 シード 7 シード 5 シード 6 ビット 14 がシールドされていると、単一の反転は致命的ではなく、崩壊は 103 ～ 10⁴ 反転後に起こります 7 つの宇宙線シードがオンQwen2.5-Coder-3B、指数 MSB を省略。 7 つすべてが最終的に崩壊し、79,000 ～ 490,000 回のフリップが発生します。
ビット 14 は保護され、シードは 7 つあります。それぞれは最終的に崩壊しますが、約 79,000 回の反転 (シード 4) から最大約 490,000 回の反転 (シード 6) になります。
ビット 14 がシールドされると、モデルの放射線耐性が大幅に向上しました。 20時くらいに倒れるのではなく

ビットフリップを実行した後、モデルは数万個まで比較的無傷で生き残りましたが、最初の実行は約 78,500 回のビットフリップ後に 0/164 に崩壊しました。非常に興味深いのは、一部のモデルが破損が進行するにつれて 1 つまたは 2 つの問題を何とか成功させ、ゆっくりとした劣化が一瞬だけ明確になったことです。純粋な好奇心とは別に、私の二分探索の妥当性にも疑問が生じました。探索空間の分割方法が単調であることを前提としているからです。私の検索では、一般に、1 回の実行で最初の劣化ポイントを正確に特定することはできません。私はそれについてあまり考えないようにして、その問題を永遠に解決することにしました。
障害モードの中で一般的なのは、トークン ループに陥った出力です。あるケースでは、モデルがスタックする前に設定された最後のビットは、モデルのトークン埋め込み行列の重みのビット 13 でした。残りを元に戻してそれのみを破損すると、モデルも破壊され、139 HE 問題から 2/164 まで即座に低下しました。偶然にも、その重みは、แฟน (ボーイフレンド/ガールフレンド) や แฟชั่น (ファッション) などの単語の始まりであるタイ語のトークン แฟ に対応していました。この値は、入力と出力の両方に対するお気に入りのタイ語トークン (Thaiken) の重要性を制御します。このビットを反転すると、出力として生成される可能性がはるかに高くなりますが、入力としての重要性もはるかに高くなるため、このようなタイケン ループが発生します。モデルがまだ解決できた 2 つの問題のうち、1 つの解決策はループを回避しました。もう 1 つは、モデルがループから抜け出す前にトークン スパムがコメントで終わっていることです。
この無限のトークン ループは、モデル照射の一般的な末期症状でした。場合によっては、これらのループがストップ トークンを発行しないため、世代がトークン制限を使い果たし、永遠に時間がかかることもあれば、世代がストップ トークンを発行することもありました。

以下のnearest_pairの問題で起こるように、早すぎます。
「」パイソン
import List の入力から
def Separate_paren_groups(paren_string: str) -> リスト[str]:
""" この関数への入力は、ネストされた括弧の複数のグループを含む文字列です。あなたの目標は、
それらのグループを個別の文字列に分割し、それらのリストを返します。
別々のグループのバランスが取れています (それぞれ 2 回反転した後 — 壊れた ``python)
import List の入力から
def Separate_paren_groups(paren_string: str) -> リスト[str]:
""" この関数への入力は、ネストされた括弧の複数のグループを含む文字列です。あなたの目標は、
これらのグループを個別の文字列に分割すると、返される率が 100% に達します。
d率到達率は100%に達しました。
d率到達率は100%に達しました。
d率到達率は100%に達しました。
d率達成率成功クリーン — ``python を渡す
import * と入力してから
コレクションインポートから*
def truncate_number(数値: float) -> float:
""" 正の浮動小数点数を指定すると、次のように分解できます。
整数部（与えられた数値より小さい最大の整数）と小数点
(残りの部分は常に 1 より小さくなります)。
2 回反転した後にデシを返す — 壊れた ```Python
import * と入力してから
コレクションインポートから*
def truncate_number(数値: float) -> float:
""" 正の浮動小数点数を指定すると、次のように分解できます。
および整数部分（最大の整数より小さい整数）

[切り捨てられた]

## Original Extract

Forty-nine billion bits of our precious intelligence can

← Writing
Simulating cosmic rays to lobotomize LLMs
Benedikt Holm · August 20, 2026
Forty-nine billion bits of our precious intelligence can't add two numbers after one cosmic ray.
In space, there is a scourge known as cosmic rays. When one strikes the bits our computers are made of, they can spontaneously flip from 0 to 1, or vice versa. These bit-flips are mostly known for allegedly helping speedrunners win faster , meddlin' with Belgian politics , making Voyager-2 go out of sync , but they've also been shown to brutalize machine learning models. Today, people like Elon Musk plan to put data centers in space, where bit-flips are caused not just by cosmic rays but also by much stronger ambient radiation, and are thus far more common than at sea level. This got me wondering: how vulnerable are LLMs to cosmic ray attacks? But I can already hear you frantically typing "You idiot, there is such a thing as ECC, heard about it?" , and yes I have, stop interrupting me, I'll get to it.
To preface, this is less a scientific writeup and more a fun little log of a neat experiment I came up with while putting my newborn to sleep. Nor is it entirely novel, since prior research has looked into bit-corruption of LLMs , though as far as I could tell it mostly focuses on targeted or adversarial bit manipulation. I also want to be clear that what I am simulating here is not particle physics, a GPU or a true radiation environment here; I'm uniformly randomly flipping model bits as a toy model of radiation-induced soft errors.
A lot of the work described here could be carried much further, but I find myself constrained both on time (newborn) and on compute, so this work is by necessity limited. Either way, I hope that you derive some enjoyment out of the horror I put these models through.
Me and Qwen during the experiment.
To run my experiment on how hard a cosmic lobotomization an LLM could handle, the first step I had to take was to select the subject. After considering my options, I initially opted to perform the experiment locally rather than rent equipment. This meant two things: 1) I could claim that I was performing important research for the local-LLM community (shoutout local.ai ), and 2) I wouldn't have to spend too much money on expensive, out-biddable hardware on sites like vast.ai. This meant however that I was limited to models that could fit on my laptop's 5070 (yes, I am compute-poor). The victim that I ended up choosing was Qwen2.5-Coder-3B , a coding LLM trained to solve programming tasks. A thing to note, is that I ran Qwen in FP16, each weight being a 16 bit floating point number. Of the 16 bits, number 14 is the most magical, since it is the most significant bit (MSB) of the weight's exponent. A flipped bit in this position can change its number dramatically, as you can play around with in the following digital fidget toy:
= 0.02063 reset 0 0 1 0 0 1 0 1 0 1 0 0 1 0 0 0
Click any bit to flip it. Bit 14 (red ring) is the cosmic ray: it turns 0.021 into 1352. Press reset to restore.
To measure how flipping bits around in the model's brain affected its outputs, I used the human-eval (HE) benchmark, a set of 164 coding problems. Note that HE is widely considered obsolete due to many models including it in their training data. I figured that since I was going to be performing low-accuracy brain surgery with a simulated space laser on Qwen, it didn't matter too much what it was trained on. Before starting the torture, I recorded the base model solving 139 out of the 164 HE problems, or about 85%. Remember this score, because it's not getting any better.
After getting the Qwen running, addressing the weights in memory was relatively straightforward, the Transformers library let me address individual weights directly, so flipping a bit in a given weight was as easy as XOR-ing it as so:
where x is the position of the bit to flip within the weight. This also made reverting the damage very quick and easy. In this experiment, the model inference was kept greedy, by setting the do_sample flag in the generation to False. This basically means that the seed or temperature plays no role in the inference, turning the output into a deterministic function of the value of the weights. And so it would have, if I had not varied batch sizes between runs to try to optimize my tps, resulting in slight wobbliness in the model's baselines between some runs, just in case you noticed that some runs start at slightly different HE baselines. If you didn't notice, please forget this last sentence.
Before starting, I was fairly sure that zapping the model bit-by-bit would take forever, so I designed a kind-of bit-flip tape, a seeded sequence of bit locations to flip. This let me precompute which bits would be corrupted in what order, and instead of flipping one bit and regenerating tokens each time, I could corrupt 500,000 at once and check for degradation in the HE score. If degraded, I knew the killing blow was somewhere in those half a million bits, so I could roll back half and test again, rolling back further if still broken. If the model was fine, I could double the damage and repeat. This is bisection search, and it's invaluable for cutting down large search spaces, saving time and compute, both of which I am desperately lacking. Note that I decided not to guard against repeated flips, so a single bit could be flipped multiple times in one run.
I fully expected the first detectable damage to only be around bit two million or even three million. But to my surprise, when my first bisection search had finished, it turned out that it had taken only six flipped bits to render the model totally unusable where it endlessly spammed a repetition of a few Chinese characters. That magical sixth bit was exactly the previously mentioned MSB in the exponent of a single feed-forward weight in layer 27, turning the weight's value from 0.021 into 1352. Up until the final lobotomizing bit, the output of the model was indistinguishable (to HE) in quality from the uncorrupted model.
for i in range ( len ( numbers )):
for j in range ( i + 1 , len ( numbers )):
if abs ( numbers [ i ] - numbers [ j ]) < threshold :
return True
return False !__!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_ … (+403 more chars of the same) return number - int ( number ) big_ expression[!expression[!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! … (+424 more chars of the same) balance = 0
for operation in operations :
balance += operation
if balance < 0 :
return True
return False !__#_!_."__!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_!_! … (+404 more chars of the same) if not numbers :
return ( 0 , 1 )
total_sum = sum ( numbers )
total_product = 1
for number in numbers :
total_product *= number
return ( total_sum , total_product ) !__!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!_!_!_!__!_!_!_!_!_!_!_!_!_!__!_!_! … (+403 more chars of the same) The same model on four HumanEval problems. Click a tab. Left: clean output, passing the tests. Right: after six random bit flips - the sixth lands on bit 14 - every answer collapses into noise that will not even parse. In fact, when run again with a few different seeds, this brain death happened on average in only ~20 flipped bits. Every single time, the final fatal flip landed on the MSB of the FP16's exponent. Repeating each fatal flip in isolation showed that each of these flips would have sufficed to completely floor the model on their own. This weakness to the MSB is also well documented in existing work, but it is interesting to see it so well replicated here.
0 10 20 30 40 50 60 Random bit flips applied 0 20 40 60 80 100 120 140 HumanEval problems solved seed 4 seed 6 seed 2 seed 3 seed 8 seed 7 seed 5 Random bit flips heavily affect model performance. Seven cosmic-ray seeds on Qwen2.5-Coder-3B, none shielded. Each holds full score, then one flip ends it.
HE score degrading as bits get flipped.
But, what if somehow the stochastic de-cerebration magically avoided bit 14, how many zaps to the more fault-tolerant parts of its brain can Qwen take? To test this, I simply ran the same randomized bit corruption, skipping over rays that would have landed on the MSB.
Random bit flips applied, bit 14 shielded (thousands) 0 20 40 60 80 100 120 140 HumanEval problems solved seed 4 seed 8 seed 2 seed 3 seed 7 seed 5 seed 6 With bit 14 shielded, no single flip is lethal, and collapse comes 10³ to 10⁴ flips later Seven cosmic-ray seeds on Qwen2.5-Coder-3B, sparing the exponent MSB. All seven eventually collapse, between 79,000 and 490,000 flips.
Bit 14 spared, seven seeds: each collapses eventually, but from ~79,000 flips (seed 4) up to ~490,000 (seed 6).
With bit 14 shielded, the model's capability to withstand irradiation rose sharply. Instead of collapsing at around 20 bitflips, the models survived relatively unscathed into the tens of thousands, the first run collapsing to 0/164 after ~78,500 bit flips. What was very interesting was that some models managed to claw back success on one or two problems as the corruption progressed, a brief moment of lucidity in their slow deterioration. Aside from the sheer curiosity of it, it also put in question the validity of my bisection search, since it assumes that how you divide the search space is monotonic. Seeing as how my search cannot, in general, identify the exact first degradation point in any single run. I decided to not think too much about it, thus solving that problem forever.
Among the failure modes, a common one was output stuck in a token loop. In one case, the last bit set before the model got stuck was bit 13 of a weight in the model's token-embedding matrix. Reverting the rest and corrupting only it also killed the model, dropping it instantly from 139 HE problems to 2/164. Coincidentally, that weight corresponded to the Thai token แฟ, the start of words like แฟน (boyfriend/girlfriend) or แฟชั่น (fashion). The value controls the importance of our favorite Thai token (Thaiken) for both inputs and outputs. Flipping this bit made it far more likely to be generated as output, but also given far more importance as input, resulting in these Thaiken loops. In the 2 problems the model could still solve, one solution avoided the loop; in the other, the token spam ended in a comment before the model broke out of the loop.
This endless token loop was a common terminal symptom of model irradiation. Sometimes, these loops never emitted stop tokens causing the generation to exhaust its token limit and taking forever, and sometimes the generation emitted the stop token too early as happens in the closest_pair problem below.
```python
from typing import List
def separate_paren_groups(paren_string: str) -> List[str]:
""" Input to this function is a string containing multiple groups of nested parentheses. Your goal is to
separate those groups into separate strings and return the list of those.
Separate groups are balanced (each after 2 flips — broken ```python
from typing import List
def separate_paren_groups(paren_string: str) -> List[str]:
""" Input to this function is a string containing multiple groups of nested parentheses. Your goal is to
separate those group into separate strings and returnd率达到率达到100%。
d率达到率达到100%。
d率达到率达到100%。
d率达到率达到100%。
d率达到率达到 clean — passes ```python
from typing import *
from collections import *
def truncate_number(number: float) -> float:
""" Given a positive floating point number, it can be decomposed into
and integer part (largest integer smaller than given number) and decimals
(leftover part always smaller than 1).
Return the deci after 2 flips — broken ```python
from typing import *
from collections import *
def truncate_number(number: float) -> float:
""" Given a positive floating point number, it can be decomposed into
and integer part (largest integer smaller

[truncated]

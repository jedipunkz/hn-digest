---
source: "https://xcancel.com/i/article/2086158118354887060"
hn_url: "https://news.ycombinator.com/item?id=49233147"
title: "AI Settles a 25 Year-Old Problem We Left Behind"
article_title: "AI Settles a 25 Year-old Problem We Left Behind | XCancel"
author: "famouswaffles"
captured_at: "2026-08-09T17:22:36Z"
capture_tool: "hn-digest"
hn_id: 49233147
score: 3
comments: 0
posted_at: "2026-08-09T16:55:13Z"
tags:
  - hacker-news
  - translated
---

# AI Settles a 25 Year-Old Problem We Left Behind

- HN: [49233147](https://news.ycombinator.com/item?id=49233147)
- Source: [xcancel.com](https://xcancel.com/i/article/2086158118354887060)
- Score: 3
- Comments: 0
- Posted: 2026-08-09T16:55:13Z

## Translation

タイトル: AI が人類が残した 25 年来の問題を解決
記事のタイトル: AI は私たちが残した 25 年間の問題を解決します | Xキャンセル

記事本文:
私たちが残した 25 年間の問題を AI が解決 | Xキャンセル
Xキャンセル
(寄付)
私たちが残した25年間の問題をAIが解決
先週、GPT-5.6 とクロード・ファブルは、2000 年から 2010 年代にかけて集中的に研究され、不安を抱えていた博士課程 1 年生の私が短期間取り組んだ無線通信における未解決の理論的疑問に決着をつけたようです。おそらく私が質問をした最後の数人の一人であり、機械にそれを解決するように促した最初の人間だったからかもしれません😊
結果として、N × N ガウス無線チャネルを通じて N ビットを送信すると、受信機はそれらすべてを正確に復元する必要があります。 2000 年代以降、信号対雑音比が少なくとも 2 log N である場合に情報を取得することが理論的に可能であることが知られていました。しかし、その値に到達することが知られている唯一の方法は、指数関数探索アルゴリズムでした。単純な多項式時間アルゴリズムがまったく同じしきい値で成功することが証明されました。
それについてもう少し詳しくお話しましょう。
2009 年に、私は Alex Dimakis ( @AlexGDimakis ) と一緒に最初の論文に取り組みました。彼はすぐに私の博士課程アドバイザーになりました (その論文のせいではありません)。
この論文は、MIMO 検出に対する多項式時間ソリューションを提供する多くの試みの 1 つです。
MIMO検出とは何ですか?
送信機は、N 個の送信アンテナと N 個の受信アンテナからなる無線チャネルを通じて N ビットのベクトルを送信します。チャネルはビットをすべて混合し、ノイズを追加します。チャネル行列を知っている受信者は、どのビットが送信されたかを把握する必要があります。
ブロック エラー最適受信機、別名最尤 (ML) 検出器は、受信信号を考慮して、送信された可能性が最も高いベクトルを見つけることで、この問題を正確に解決します。この場合、ML 検出は要するに、次の基本的な離散最小二乗問題を解決することになります。
残念

人生におけるすべての良い問題と同様に、ML 検出は NP 困難です。
しかし、私たちは TCS 悲観論者ではありません。ワイヤレス チャネルは最悪のケースではなく、ランダムなものです。コミュニティは 2000 年代初頭から次の質問に取り組んできました。
送信されたビットの復元が統計的に可能である場合、ポリタイムで実行できるでしょうか?
上記の 2010 年の論文では、この問題についてはあまり進展がありませんでした。この分野で多くの作業が行われたにもかかわらず、私の理解する限り、この問題は 2001 年以来未解決のままです…よりドラマチックに聞こえるようにするため、「四半世紀」とも呼ばれます。
先週までは。そして最終的に導き出される答えは、
はい！完全な検出が統計的に可能な場合は、多項式時間で実行できます。
しかし、誰が気にするでしょうか？これについては後ほど説明します。
私は紙を添付していますが、証明と説明を簡略化するためにモデルを行ったり来たりするのに 5 日以上費やしました (これはもともと大惨事でした)。このプロセスには、GPT が作成した最初の証明 (約 30 分程度かかりました) よりもはるかに長い時間がかかりました。証明は長いですが、比較的初歩的なものです。私はすべてを検証し、証明チェックの能力の限りにおいて、それは正しいです。
ここで、この問題とその歴史について、また、この分野が MIMO 検出理論のこの特定の分野から離れたにもかかわらず、なぜこの問題について書く価値があると私が考えるのかについて、もう少しお話したいと思います。
したがって、バイナリ ベクトル x を {±1}^N で送信し、受信します。
ここで、H は N×N であり、H と w は両方とも iid N(0,1) エントリを持ち、すべて独立しています。受信機は H とノイズ統計を知っていますが、w を知りません。そして、y から x を返したいと考えています。リカバリ問題に対するブロック エラーの最適解は次のとおりです。
ところで、この最適化は、MIMO 検出、CDMA マルチユーザー検出、整数最小二乗法など、さまざまな装いでも行われます。

格子内の最も近いベクトルなど。
そして、SNR = ∞ (つまり、実効ノイズ 0) の場合、問題は自明になります。チャネル行列 H は確率 1 で可逆であるため、それを反転し、inv(H)*y で正確な x を回復します。逆に、SNR = 0 の場合、ノイズからは何も検出できず、ML 検出は失敗します。
しかし、0 と無限大の間のどこかでは、ML 検出は成功し、SNR = 2 log N で正確に成功します。これは、上記の最適化問題を解くと、送信された N ビット シーケンスのすべてのビットを確率 1 に向かう傾向で完全に回復できることを意味し、それ以下 (加算的 loglogN 項まで) のブロック回復の確率は 0 になる傾向があります。
したがって、2logN を超えると、送信信号は ML 最適化問題の最適値になりますが、これを解決するには、考えられるすべての N ビット シーケンスを徹底的に検索する必要があるようです。したがって、私たちが今関心を持っている質問は次のとおりです。
ポリタイム アルゴリズムは、ML が成功したときに送信された x を復元できますか?
ちょっとしたドラマを交えた簡単な歴史
整数最小二乗問題の可解性の問題は、少なくとも 1989 年に遡り、Verdú が一般的な場合に NP 困難であることを証明しました。しかし、NP の硬さは最悪の場合の記述であり、私たちのインスタンスはそうではありません。
2001 年の Hassibi と Vikalo は、私の知る限り、平均的なケース、ポリタイムの解決策に希望があると主張した最初の人でした。彼らが分析したアルゴリズムは、1985 年の Fincke と Pohst にまで遡る、当時人気のある手法である Sphere Decoder (SD) でした。Sphere Decoder が特に興味深い理由は、1) 正確な ML アルゴリズムである、つまり常にミニマイザーを出力する、2) 実際には指数関数的な時間よりもはるかに高速であると思われたためです。
したがって、SD がポリタイムで実行されることを実際に証明できることが期待されていました。 H&V は論文で次のように述べています。

彼らは、チャネルとノイズを平均した Sphere Decoder の予想される複雑さの式を導出し、それが多項式であることを示しました。それが本当であれば、その疑問は解決しました。それは信じられないような結果のように思えました。
その後、Jaldén と Ottersten は 2005 年に、漸近的解釈がまったく正しくないことを示しました。固定 SNR では、どんなに大きくても、球面デコードの予想される複雑さは、実際には問題の次元で指数関数的になります。
したがって、正確かつ高速では機能しなかったため、この分野では ML 最適化問題の近似に多大な労力を費やしました。高い SNR での近似保証とタイトネス条件を備えた半明確な緩和ですが、鋭いしきい値はありません。ビット反転ローカル検索はシミュレーションで ML と一致するように見えましたが、ML 回復しきい値と一致するという完全な証明はありませんでした。 AMP の文献では、ブロック回復が不可能な固定 SNR でのビットごとのエラーを厳密に特徴付けています。統計物理学は、レプリカレベルの引数を使用して正確な ML を追跡すると予測される多時間手法を生成しましたが、私の理解する限り、証拠はありません。そして、上記の 2010 年の Babak と Alex の論文は、混合後の定常分布により消失しない質量が正しい溶液に置かれることを証明する MCMC 法を分析しましたが、混合時間については何も証明していませんでした。これが難しい部分です。
長年にわたり、あらゆる SNR スケールでのブロック回復が厳密に保証されているのは、まさに 1 つの多項式時間手法だったと私には思われます。2020 年のボックス緩和では、SNR が 4 log N 程度で、おそらくそれ以下ではない場合にブロックを回復することが示されました。余談ですが、このような手法を分析するために必要な確率論的ツールが 2010 年代後半に成熟し、そのほとんどはコミュニティが移行した後のことであるのは興味深いことです。

そしてすでに解散していた。
そしてそれ以来…あまり活動はありませんでした。
簡単に言うと、ML が達成できることと、多項式時間手法で証明できることとの間のギャップは決して埋まりませんでした。
GPT とクロードは何をしましたか、そして私、ディミトリスが検証できる証拠をどのようにして入手しましたか?
最近、難しい数学の課題でフロンティア モデルが理不尽な成功を収めたことに動機付けられ、私は大学院生の頃に悩まされていた問題 (以前は情報理論とコーディング理論に取り組んでいました) に立ち戻り、それらに死の星を向け始めることにしました。数学の難しい質問をして、GPT ゼロで答えてもらうと、まさに次のような気分になります。
しかし、少し問題があることはわかっていました。たとえ質問に対して完全な答えが返ってきたとしても、それをより広く共有したい場合は、それを確認する必要があることがネックになります。 1 つは、それが間違っていることが判明した場合に恥をかきたくないからであり、2 つは、私たちが質問したり科学を行ったりする主な理由は、共有することだからです。
そこで私は、博士課程の初期に私を悩ませた最も野心的な質問の 1 つと思われるもの、そして明白でまだ未解決の質問を選択することにしました。そこで、ML MIMO 検出がいつポリタイムで解決できるかを GPT-5.6 と Claude Fable 5 に尋ねました。
どちらも異なるアルゴリズムの証明を作成し、ギャップはないと自信を持って述べています。 2 log N を超える SNR で成功する多項式時間アルゴリズムがあり、ML 回復しきい値と正確に一致します (加法対数項までですが、気にする必要はありません)。
しかし、小さな問題がありました😊 GPT のアルゴリズムは AMP のバリアントでした。私は AMP を心から嫌います。なぜなら、AMP の分析は一生理解できないからです。そこで私は、可能であれば、より単純なアルゴリズムで同じ結果を再検証するように指示しました。実際、GPT は別のアルゴリズムを作成しましたが、これも直観に反すると感じました

e、そしてこれまでに使われたのを見たことがないものです！
一方、Fable では、私がとても気に入ったものを思いつきました。
符号付き LMMSE、その後貪欲なビット 反転。過去に紹介され、実際に実務で使われていたアルゴリズム。
しかし、別の問題がありました！ GPT によると、ファブルの証明はほとんどが間違っていましたが、修復可能でした。そこで私は、Fable が提案したアルゴリズムに固執することにし、GPT に Fable の証明を取得して修正するよう依頼しました。そして、それはできました！
しかし、さらに別の問題がありました。この新しい証明は読めませんでした。記法の壁、変数を指す変数が他の変数を定義する変数の比率を指す、エキゾチックな行列分析と確率機構、蕁麻疹を引き起こすマルチェンコとパストゥールの隣接するもの、その他の美しいものです。
そこで、約 4 ～ 5 日間、私は 2 つのモデルの間を行ったり来たりし続け、証明を解決するために必要な大きなコンポーネントごとに、可能な限り愚かな一連の手順を教えてもらいました。私は、わかりやすくするために、2 log N のしきい値が続く限り、境界と定数が悪化しても問題ないことを明確に伝えました。
私が欲しかったのは、集中力の持続時間が短い年老いた恐竜が泣かずに消化できるという証拠だけでした。
実は GPT と Claude に、私が最も愚痴っていたメッセージをシェアしてくれるように頼んだのです（笑）
なぜ私は非常に簡単な手順にこだわったのでしょうか?それを端から端まで自分で検証したかったからです。いいえ、リーンは使いたくありません。問題は解決しません。形式的検証は抽象化レベルを別の場所に移動するだけです。補題の英語が忠実に Lean に翻訳されるかどうかを確認する必要がありますが、これは私には理解できない言語です。
はい、忘れてください。リーンは好きじゃない、ごめんなさい。
しかし、私は基本的な線形代数と確率を理解していますし、自分自身を信じています。

そのようなステップを踏みます。つまり、それが私が要求する証拠のレベルです。
それから、モデルが互いの議論を単純化しながら、何度も促し、促し、促し続ける作業が何日もかかりました。その間、私は文句を言い、従えないものはすべて拒否し続けました。
そして最終的にはうまくいきました！最終的には、私が完全に理解しており、一行ずつ確認したという証拠を得ることができました。
証明するのに 30 分かかり、私が検証できるようになるまでに 5 日かかりました。ちょっと異常な比率ですが、それが現実です。その結果、多項式時間で最尤法が機能するときはいつでも、単純なアルゴリズムが機能します。この問題には計算と統計のギャップはありません。
証明の高度なものは何ですか?
アルゴリズムは恥ずかしいほど単純です。しかし、なぜこれが機能するのでしょうか? LMMSE に続いて丸めを行うと、ハミング距離の観点から、送信信号の消失部分内、つまり真実から o(N) 以内に収まります。
そうすれば、ステップごとの降下ゲイン (つまり、コストがどの程度改善されるか) がガウス量によって支配され、ガウス量の均一な集中により、特定のボール内のすべての非グラウンドトゥルース ベクトルが、保証されたサイズの厳密に改善されたビット フリッピングを提供することが確立されるため、貪欲なビット フリッピングが行き詰まる可能性はありません。つまり、何をするにしても、定義するのはあなたです

[切り捨てられた]

## Original Extract

AI Settles a 25 Year-old Problem We Left Behind | XCancel
XCancel
(donate)
AI Settles a 25 Year-old Problem We Left Behind
Last week, GPT-5.6 and Claude Fable appear to have settled an open theoretical question in wireless communications that was intensely studied between 2000 till the 2010s and that I briefly worked on as an anxious first-year PhD student. The answer finally arrived, perhaps because I was one of the last few to ask the question and the first to prompt the machines to solve it 😊
The result: you send N bits through an N × N Gaussian wireless channel, and the receiver must recover all of them exactly. It has been known since the 2000s that it is information theoretically possible to do so when the signal-to-noise ratio is at least 2 log N. But the only method known to reach that was an exponential search algorithm. There is now proof that a simple, polynomial-time algorithm succeeds at the exact same threshold .
Let me tell you a little more about it.
In 2009, I worked on my first paper with Alex Dimakis ( @AlexGDimakis ), who soon after became my PhD advisor (not because of that paper):
The paper was among many attempts to offer a polynomial time solution to MIMO detection.
What is MIMO detection you ask?
A transmitter sends a vector of N bits through a wireless channel of N transmit and N receive antennas. The channel mixes the bits all together and adds noise. The receiver, who knows the channel matrix, has to figure out which bits were sent.
The block-error optimal receiver, aka the maximum likelihood (ML) detector, solves exactly this problem by finding the most likely vector that could have been sent, given the received signal. In this case, ML detection boils down to solving this fundamental discrete least squares problem:
Unfortunately, as with all good problems in life… ML detection is NP-hard .
Yet we are not TCS pessimists, and wireless channels are not worst-case, they are random, and the community had been working on the following question since the early 2000s:
When recovering the transmitted bits is statistically possible, can we do it in poly-time?
We didn’t make much progress on that question in that 2010 paper above, and despite plenty of work in the area, the problem has, as far as I understand, remained open since 2001 … aka a QUARTER OF A CENTURY to make it sound more dramatic.
Until last week. And the final answer comes out to
YES! Whenever perfect detection is statistically possible, you can do it in polynomial time.
But who cares? We’ll come back to this in a second.
I am attaching the paper and I have spent 5+ days going back and forth with the models to simplify the proofs and the exposition (which was originally an absolute disaster), a process that took much much longer than the initial proof that GPT produced (which took around 30 minutes or so). The proof is long, but relatively elementary. I have verified everything and to the best of my ability to proof check, it is correct.
Now let me talk a bit more about the problem and its history, and why I think it’s worth writing about even though the field has moved on from this specific corner of MIMO detection theory.
So you transmit a binary vector x in {±1}^N, and receive
where H is N×N and both H and w have iid N(0,1) entries, all independent. The receiver knows H and the noise statistics, but not w, and wants x back from y. The block error optimal solution to the recovery problem is equal to
BTW this optimization comes under different guises too: MIMO detection, CDMA multi-user detection, integer least squares, closest vector in a lattice, etc etc.
And when SNR = ∞ (i.e., effective noise 0) the problem becomes trivial: the channel matrix H is invertible with probability 1, so you invert it, and recover the exact x with inv(H)*y. At the other extreme, when SNR = 0, there’s nothing you can detect from noise, and ML detection fails.
But somewhere in between 0 and infinity, ML detection succeeds, and does so precisely at SNR = 2 log N. This means that solving the optimization problem above allows you to perfectly recover all the bits of the transmitted N-bit sequence with probability tending to 1, and below (up to additive loglogN terms) the probability of block recovery tends to 0.
So above 2logN, the transmitted signal is an optimum of the ML optimization problem, but solving it seems to require an exhaustive search over all possible N-bit sequences. So the question we now care about is:
Can a poly-time algorithm recover the transmitted x when ML succeeds?
A brief history with a tiny bit of drama
The question of solvability of the integer least squares problem is at least as old as 1989, when Verdú proved that it is NP-hard in the general case. But NP-hardness is a worst-case statement, and our instances are not.
Hassibi and Vikalo in 2001 were the first—as far as I know—to argue that there is hope for an average case, poly-time solution. The algorithm they analyzed was a popular method at the time, the Sphere Decoder (SD), dating back to Fincke and Pohst in 1985. The Sphere Decoder was of particular interest because 1) it’s an exact ML algorithm, i.e., always outputs the minimizer and 2) it seemed to be way faster than exponential time in practice.
So the hope was that one could actually prove that SD runs in poly-time. This is what H&V articulated in their paper: they derived a formula for the expected complexity of the Sphere Decoder , averaged over the channel and the noise, and showed that it looks polynomial. If that was true, the question was settled. That seemed like an incredible result.
Then Jaldén and Ottersten in 2005 showed that the asymptotic interpretation was not quite correct: at any fixed SNR, no matter how large, the expected complexity of sphere decoding is in fact exponential in the problem dimension.
So since exact and fast was not working, the field spent considerable effort working on approximations to the ML optimization problem. Semidefinite relaxations with approximation guarantees and tightness conditions at high SNR , but no sharp threshold. Bit-flipping local search seemed to match ML in simulations but no full proofs of matching the ML recovery threshold. The AMP literature rigorously characterized the per-bit error, at fixed SNR, where block recovery isn’t possible. Statistical physics produced poly-time methods that were predicted to track exact ML using replica-level arguments, but as far as I understand, no proof. And the paper with Babak and Alex from 2010 above analyzed an MCMC method proving that after mixing, the stationary distribution puts non-vanishing mass on the correct solution but did not prove anything about the mixing time, which is the difficult part.
In all those years, it seems to me, exactly one polynomial-time method came with rigorous block-recovery guarantees at any SNR scale: the box relaxation, in 2020 , shown to recover the block when the SNR scales like 4 log N and, provably, not below. As an aside, it’s kind of interesting that the probabilistic tools needed to analyze such a technique matured in the late 2010s, which for the most part was after the community moved on and had already dispersed.
And since then… not a lot of activity.
So long story short, the gap between what ML achieves and what any polynomial-time method could provably achieve never closed.
What did GPT & Claude do and how we got a proof that I, Dimitris, can verify?
Motivated by the recent unreasonable successes of frontier models on hard math tasks, I decided to go back to problems that haunted me as a graduate student (I used to work on information and coding theory) and start pointing the death star at them. This is precisely how it feels to ask hard math questions and have GPT zero shot them:
But I knew there was a bit of a problem. Even if I got back a full answer to any question that I’d ask, I’d be bottlenecked by having to verify it if I wanted to share it more broadly. One, because I don’t want to embarrass myself if it turns out to be wrong, and two, because sharing is the main reason we ask questions and do science anyways.
So, I decided to pick what felt like one of the most ambitious questions that bothered me early in my PhD, and one that was clean to state, and still open. So, I asked GPT-5.6 and Claude Fable 5 when ML MIMO detection can be solved in poly-time.
Both produced proofs for different algorithms confidently stating that that there is no gap! There is a polynomial-time algorithm that succeeds at SNR above 2 log N, matching exactly (up to additive loglog terms, but who cares) the ML recovery threshold.
But there was a small problem 😊 GPT's algorithm was an AMP variant . And I hate AMP, with a passion, because I do not, for the life of me, understand any of its analyses. And so I told it to try and reprove the same result, if possible, for a simpler algorithm. Indeed, GPT produced another algorithm that I also found counter-intuitive, and one that I have never seen used before!
Fable, on the other hand, came up with something I really liked:
signed LMMSE, then greedy bit flips. An algorithm that was introduced in the past and actually used in practice.
But there was another problem! According to GPT, Fable’s proof was mostly wrong.. but salvageable. So I decided to stick with the algorithm that Fable suggested, and asked GPT to take Fable’s proof and fix it. And it did!
But there was yet another problem, this new proof was UNREADABLE: a wall of notation, variables pointing to variables pointing to ratios of variables defining other variables, exotic matrix-analysis and probability machinery, Marchenko–Pastur adjacent stuff that gives me the hives, and other beautiful things.
So for around 4-5 days I kept going back and forth between the two models and asking them to give me the dumbest possible set of steps, for each of the big components needed for the proof to work out. I explicitly told them that it was OK for the bounds and constants to get worse, AS LONG AS the 2 log N threshold remains, all for the purpose of simplicity.
All I wanted was a proof that an old dinosaur with a short attention span can digest without crying.
I actually asked GPT and Claude to share back the messages where I was whining the most, lol
Why did I insist on super simple steps? Because I wanted to verify this myself, end to end. And no, I don’t want to use Lean it DOES NOT solve my problem. Formal verification just moves the abstraction level somewhere else!! You still have to verify that the English of a lemma faithfully translates to Lean, which is a language I don't understand.
Yeah, forget about it. I don’t like Lean, sorry.
But I do understand basic linear algebra and probability, and I trust myself verifying such steps. So that’s the level the proof I demand.
It then took multiple days of prompting and prompting and prompting, with the models simplifying each other’s arguments, while I kept complaining and rejecting anything I could not follow.
And in the end it worked! We ended up with a proof that I fully understand, and that I have now checked line by line.
Proving the thing took 30 minutes and making it verifiable by me took 5ish days. That’s a bit of an insane ratio, but it is what it is. And the result: a simple algorithm works whenever maximum likelihood works, in polynomial time. There is no computational-statistical gap in this problem.
What is the high level of the proof?
The algorithm is almost embarrassingly simple. But why does this work? LMMSE followed by rounding gets you, in terms of Hamming distance, within a vanishing fraction of the transmitted signal, i.e., o(N) away from the truth.
Then, greedy bit-flipping can’t get stuck because the per-step descent gains (i.e., how much the cost improves) are governed by Gaussian quantities, and their uniform concentration establishes that every non-ground-truth vector within a certain ball offers a strictly improving bit flip of a guaranteed size. Meaning no matter what you do, you defin

[truncated]

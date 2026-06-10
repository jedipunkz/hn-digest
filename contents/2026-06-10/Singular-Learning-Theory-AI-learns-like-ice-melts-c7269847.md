---
source: "https://tcz.hu/blog/2026/02/26/singular-learning-theory/"
hn_url: "https://news.ycombinator.com/item?id=48476343"
title: "Singular Learning Theory: AI learns like ice melts"
article_title: "Singular Learning Theory: AI learns like ice melts – Zoltan's Blog"
author: "hntcz"
captured_at: "2026-06-10T16:22:43Z"
capture_tool: "hn-digest"
hn_id: 48476343
score: 1
comments: 0
posted_at: "2026-06-10T13:55:30Z"
tags:
  - hacker-news
  - translated
---

# Singular Learning Theory: AI learns like ice melts

- HN: [48476343](https://news.ycombinator.com/item?id=48476343)
- Source: [tcz.hu](https://tcz.hu/blog/2026/02/26/singular-learning-theory/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T13:55:30Z

## Translation

タイトル: 特異な学習理論: AI は氷が溶けるように学習する
記事タイトル: 特異学習理論: AI は氷が溶けるように学習する – Zoltan のブログ
説明: 直観的にはニューラル ネットワークは機能しないはずですが、明らかに機能します。何故ですか？そしてそれは氷と何の関係があるのでしょうか？

記事本文:
特異な学習理論: AI は氷が溶けるように学習します |ゾルタンのブログ
について
特異な学習理論: AI は氷が溶けるように学習する
ニューラルネットワークは機能しないはずです。
直観的には、数十億または数兆のパラメータを持つニューラル ネットワークは一般化できないはずで、トレーニング データを壊滅的に過剰適合するはずです。しかし、ディープ ニューラル ネットワークは、さまざまなタスクに対して顕著な一般化能力を実証しています。
それを説明する最も明確な数学的物語は、特異学習理論と呼ばれます。 SLT は、将来的に AI の安全性と解釈可能性のための重要な数学的ツールとなる可能性があります。
SLT とは何ですか? SLT はどのような問題を解決しますか?そしてそれは氷と何の関係があるのでしょうか？
ニューラル ネットワークは、線形回帰やロジスティック回帰などの古典的なモデルに適した統計ツールでは分析できません。なぜだめですか？私たちの多くと同じように、彼らも退化性を持っています。古典的なモデルでは、損失ランドスケープはどこでも曲率を持ち、フィッシャー情報行列はフルランクです。専門用語を減らして、パラメータを変更すると、損失が滑らかに変化します。この単純な 2 パラメーター線形回帰のおもちゃのモデルを自分で試すことができます。
パラメータを少し変えると、滑らかなボウルの表面での損失が徐々に変化します。ニューラル ネットワークには当てはまりません。一部のニューロンは出力重みがゼロの場合があり、層内のニューロンの順序によってモデルが変更されることはありません。損失は​​、すべてのパラメータの変動に反応しない可能性があります。これにより、モデルが特異なものになります。
古典的なモデルでは、すべてのパラメーターがモデルの機能に影響を与えるため、すべてのパラメーターが「有効」になります。モデルの複雑さは、これらのパラメーター $d$ の数によって特徴付けることができます。この数値は、ベイズ自由エネルギー、BIC ペナルティ、事後エントロピーなど、さまざまな場所に現れます。 (通常 $d$ は分割されます

2つまでに。 $1/2$ の因数は、$d$ 次元のガウス分布から外れる数学的アーティファクトです。従来の複雑さの指標は $d/2$ です。)
ニューラル ネットワークは異なります。縮退のため、有効なパラメーターの数は通常 $d$ よりもはるかに少なくなります。 (したがって、典型的なニューラル ネットワークは $d/2$ よりも実効的な複雑さが小さくなります。) これらの縮退が、実際になぜ非常にうまく一般化できるのかの秘密です。
特異学習理論とは何ですか?
SLT は渡辺純雄によって開発された統計学の分野です。これは、パラメータ空間が表現する関数と 1 対 1 ではないモデルを解析できるようにする数学的機械です。言い換えれば、モデルは滑らかなボウルではありません。ニューラルネットワークのようなもの。
代数幾何学を中心とした数学は複雑ですが、見出しの結果 $\lambda$ は理解しやすいです。 $\lambda$ は実対数正準しきい値 (RLCT) を示し、複雑さの尺度として特異モデルの $d/2$ を置き換えます。この理論は通常のモデルと互換性があるため、これらの場合は $\lambda = d/2$ となります。
どのモデルでも、RLCT はパラメーター空間のどれだけが最小損失付近にあるかを測定します。それが意味をなさない場合は、ここで勾配降下の山登りの比喩が役に立つかもしれません。自分が損失の状況の中で、損失の低い地点に座っていると想像して、自分の周りを見回してください。急な登りを戻さずに歩き回ることができますか?あなたはフィヨルドか谷にいますか？ RLCT はこれを測定します。実際には、RLCT のローカル バージョンである $\hat{\lambda}$ が計算され、特定の点の周囲の損失ランドスケープの形状が測定されます。ローカル学習係数 (LLC) とも呼ばれます。
簡単に言うと、あるポイントの周囲で LLC が低いということは、それほど登ることなく歩き回ることができることを意味しますが、LLC が高いということは、ポイントから離れるとすぐに損失が増加することを意味します。

e パラメータ空間内のポイント。
さまざまな形をしたいくつかの天然盆地。上：コロラド州ガニソンのブラックキャニオン（CC BY-SA 4.0 Terry Foote）。この 2D 損失表面のメタファーでは、モデルは脆く、極小値からわずかに離れると急な崖を登らざるを得なくなります。下：カリフォルニア州デスバレー（パブリックドメイン、国立公園ギャラリー）。損失を大きく変えることなくパラメータを変更できる、縮退した堅牢なモデルのたとえです。
ここに驚くべき数学的事実があります: ニューラル ネットワークは氷が溶けるように学習します 1 。
水が凍結すると、水分子は結晶構造の中に位置し、限られた数の有効な配置方法を持ちます。可能な微小状態の数、つまりエントロピーは比較的小さいです。温度が上昇するにつれて、エネルギー (秩序ある配置に有利) とエントロピー (無秩序に有利) という 2 つの競合する圧力の相対的な重要性が変化し、$0^\circ\text{C}$ を超えると、エネルギーは綱引きを失い、相転移が起こります。氷の結晶が緩むと分子配置の可能性が増えるため、エントロピーが急増します。
ニューラル ネットワークのトレーニングの開始時には、パラメーターは同様の損失で多くの異なる構成になる可能性があります。エントロピーが高い。トレーニングが進み、モデルがデータに適合し始めると、損失が低下し、エントロピーも低下します。氷と同じように、ネットワークの構造は、少なくとも最初はより秩序だったものになります。
興味深いのは、損失が収束した後、エントロピーが上昇し始める場合があるということです。ネットワークは、モデルの縮退により、多くの異なるパラメーター構成が同様の損失を生成するパラメーター空間の領域に漂流します。氷が溶けます。トレーニング データのモデル表現は脆さがなくなり、より安定します。これは多くの場合、より良いテストに対応します

トレーニング データの精度とよりシンプルな表現を設定します。後で実際の例を見ていきます。
エントロピーが物理システム内のマイクロステートの多重度を記述する方法と同様に、LLC はパラメーター空間の現在の低損失点でのマイクロステートの多重度を特徴付けます。トレーニング中の変化は、ネットワークの重みの質的な変化を示す可能性があります。実際、渡辺氏の数学では、ニューラル ネットワークが物理システムと同様の相転移を経験すると予測されています。 SLT の自由エネルギーは、熱力学や統計力学における相転移を記述する数学的対象と同じです。
2 つの分野におけるフリー エネルギー
ここでの $Z$ はパーティション関数と呼ばれます。物理学では、エネルギーは物理的なエネルギーであり、積分は分子の状態に関するものです。ベイズ統計では、「エネルギー」は損失とデータ ポイントの数の積であり、積分はモデル パラメーターに対するものです。自由エネルギー $F$ は両方の場の中心的な対象であり、エネルギー対エントロピー分解 $F = \langle E \rangle - TS$ は両方に当てはまります。制御パラメータの変化に伴う $F$ の非解析性として定義される相転移は、制御パラメータが温度であってもサンプルサイズであっても同じ数学的現象です。
この投稿の最も重要な文は次のとおりです。これらの表現はパラメーター空間でより大きなボリュームを占めるため、モデルは自然にデータのより単純な表現に引き寄せられます。オッカムの剃刀は外部から押し付けられた好みではなく、計算の対象外です。
LLC は、損失曲線だけからは見えない、損失状況における領域間の定性的な変化を測定します。トレーニング中の LLC の低下は、モデルがトレーニング データのよりシンプルで安定した表現に到達したことを意味します。コンセプト「c」

モデルは一般化されました。
これは AI の安全性にとって重要です。発達的解釈可能性 (devinterp) の支持者は、LLC がトレーニング中に新しい能力がいつ獲得されるか、そしてそれらがモデル内でどの程度強力に確立されるかを追跡するのに役立つと信じています。また、AI モデルが安全機構を利用しようとするときの検出にも役立つ可能性があります。
ここで悪いニュースがあります。LLC の正確な値を計算することは、比較的小規模なニューラル ネットワークであっても困難です。これには、損失ランドスケープの構造に適用される特異点の解決と呼ばれる代数幾何学のプロセスが含まれます。時間計算量の上限はおよそ $O\left(\mathrm{Ackermann}(d)\right)$ または $O\Bigl(\underbrace{2^{2^{\cdot^{\cdot^{\cdot^{2}}}}}}_{d\ \text{twos}}\Bigr)$ です。ここであなた自身の一般的な科学の扱いが難しい例えを考え出してください (宇宙の原子、ビッグバンから数秒、ヤダヤダ)。
積分を分析的に評価できない場合、次善の策は、積分を期待値として書き直し、サンプリングによって推定することです。サンプリングはどうやって行うのですか？答えは、ノイズを追加した古き良き勾配降下法です。確率的勾配ランジュバン ダイナミクス (SGLD) は、標準のオプティマイザー ステップと同様に機能しますが、各パラメーターにガウス ノイズが追加されます。山登りに例えると、風景の形を感じるために、特定の地点をランダムに歩き回ります。勾配により低損失領域の近くに留まりますが、ランダムなノイズが探索を促します。
「歩き回る」と、そのエリア周辺の軌跡から損失サンプルを収集します。以下でご自身で試してみることができます。
これらのサンプルと元の損失の違いから、谷の形状がわかります。広くて平らな場合、または急峻な V 字型である場合などです。 LLC はこれらの差のスケーリングされた平均であり、単一の数値で形状を捉えます。
H

ここで $\overline{L}_{\text{SGLD}}$ は SGLD サンプルの平均損失です。分母は実効温度で、$\eta$ は学習率、$\varepsilon$ はノイズ スケールです。
徘徊（SGLD）の計算式は次のとおりです。
最初の部分は、調整され、正規化された損失の勾配降下ステップのように見えます。 2 番目の部分はガウス ノイズです。 $\epsilon$ はステップ サイズ、$n$ はデータセット サイズ、$\gamma$ はローカライゼーションの強度、$\sigma$ はノイズの強度です。ここでの $\beta$ は温度の逆数を意味し、渡辺の標準推定量では $1/\log n$ になります。
最適なハイパーパラメータについては、活発な研究が行われています。以下のおもちゃの例では、$\gamma = 1.0$、$\sigma = 1$、および $\epsilon = \sqrt{2\eta}$ の標準推定器を使用します。ここで、$\eta$ は学習率です。各 LLC 推定では 5000 の SGLD ステップが使用され、最初の 1000 はバーンインとして破棄されます。
私は次の簡単な実験を使用して、局所学習係数 (LLC) と確率的勾配ランジュバン ダイナミクス (SGLD) に関する私自身の直観を発展させました。まず、通常のモデルを見て、その複雑さを推定します。次に、縮退を伴う単純なニューラル モデルに進みます。最後に、トレーニングを通じての LLC の進化について見ていきます。ノートブックはここから入手できます。
ウォームアップ: 線形回帰 (2 パラメーター)
データセット $y = 3x + 2$ でトレーニングされた単純な線形モデル $f(x) = wx + b$。損失の状況は、$(w, b) = (3, 2)$ に単一の最小値を持つ凸状のボウルです。これは通常のモデルであるため、LLC は $d/2$、つまり 1 にほぼ一致すると予想されます。モデルをいくつかのステップでトレーニングし、ランドスケープを視覚化し、トレーニングされた重みから SGLD を実行して、大まかな LLC 推定値を取得します。
# DATASET: 線形回帰 (y = 3x + 2)
X = トーチ。 linspace ( - 1 、 1 、 100 )。形状変更 ( - 1 , 1 )
Y = 3 * X + 2
# モデル:

2 つのパラメータ (重みとバイアス)
モデル = nn 。線形 ( 1 , 1 )
以下のプロットは、損失の状況、最小値付近でトレーニングした後のパラメーター、および SGLD がたどった軌跡を示しています (丸は開始を示し、星は終了を示します)。
計算された LLC は $\約 0.81$ で、2 パラメーターの通常モデルから予想される値 ($d/2 = 2/2 = 1$) からそれほど遠くありません。
ベースロスとの軌道の差は、[-0.02838482 -0.02812932 -0.03070474 ... 0.0922314 0.08985042 のようになります。
0.09166867] 。それらのスケーリングされた平均が推定 LLC を示します。 100 要素のデータセットのスケーリング係数 $n\beta = 100 \cdot \frac{1}{\ln 100} \約 21.7147$。
チェーンがボウルの中心からあまり外れることなくボウル内を探索している様子がわかります。実際には、より統計的な安定性を高めるために、LLC を推定するためにいくつかのチェーン (独立した軌道) が使用されますが、このおもちゃのモデルでは 1 つのチェーンのみを実行します。
非線形モデル: 単一の隠れニューロン (4 パラメーター)
さらに興味深いケースがあります。 1 つの隠れニューロンと Tanh 活性化を含む小さな MLP:
これには 4 つの学習可能なパラメータ $(w_1, b_1, w_2, b_2)$ があります。線形の場合とは異なり、損失ランドスケープは非凸であり、対称性を持っています (たとえば、$w_1$ と $w_2$ の符号を同時に反転すると、同じ関数が得られます)。
ノートの中で私は証明しました

[切り捨てられた]

## Original Extract

Intuitively, neural networks should not work but they clearly do. Why is that? And what does it have to do with ice?

Singular Learning Theory: AI learns like ice melts | Zoltan’s Blog
About
Singular Learning Theory: AI learns like ice melts
Neural networks should not work.
Intuitively, neural networks with billions or trillions of parameters should not be able to generalize, they should catastrophically overfit the training data. But deep neural networks have demonstrated remarkable generalization abilities for a variety of tasks.
The cleanest mathematical story that explains it is called Singular Learning Theory. SLT might end up being a key mathematical tool for AI safety and interpretability in the future.
What is SLT and what problem is it solving? And what does it have to do with ice?
Neural networks cannot be analyzed with statistical tools that work well for classical models like linear regression or logistic regression. Why not? Like many of us, they have degeneracies . In classical models, the loss landscape has curvature everywhere, the Fisher information matrix is full rank. With less jargon: wiggle any parameters and your loss smoothly changes. You can try it yourself on this simple, two-parameter linear regression toy model.
Small changes in the parameters change the loss gradually on the surface of a smooth bowl. That’s not true for neural networks. Some neurons may have zero outgoing weights and the order of neurons within a layer doesn’t change the model. The loss might not respond to wiggles of all parameters. This makes the model singular.
In a classical model, all parameters are “effective” because they all affect the model’s function. The model’s complexity can be characterized by the number of these parameters, $d$. This number shows up in many places: Bayes free energy, BIC penalty, posterior entropy and so on. (Usually $d$ is divided by two. The factor of $1/2$ is a math artifact that falls out of the $d$-dimensional Gaussian. The conventional complexity metric is $d/2$.)
Neural networks are different: because of their degeneracies, the number of effective parameters is usually much lower than $d$. (Consequently, a typical neural network has a smaller effective complexity than $d/2$.) These degeneracies are the secret to why they actually generalize so well.
What is Singular Learning Theory?
SLT is a branch of statistics developed by Sumio Watanabe . It is the mathematical machinery that lets us analyze models whose parameter space is not one-to-one with the function they represent. Put differently, where the model is not a smooth bowl. Like a neural network.
The math, centered around algebraic geometry, is complicated, but the headline result, $\lambda$, is easy to understand. $\lambda$ denotes real log canonical threshold (RLCT), which replaces $d/2$ in singular models as a measure of complexity. The theory is compatible with regular models, so for these $\lambda = d/2$.
For any model, RLCT measures how much of the parameter space sits near minimum loss. If that makes no sense, you might find the hill-climbing metaphor of gradient descent useful here. Picture that you are sitting at some low loss point on the loss landscape and look around yourself. Can you wander around without starting a steep climb back up? Are you in a fjord or vale? RLCT measures this. In practice, $\hat{\lambda}$, the local version of RLCT is calculated, which measures the shape of the loss landscape around a specific point. It is also called local learning coefficient (LLC).
Put simply, low LLC around a point means you can walk around without much climbing, while high LLC means the loss will quickly increase as you leave the point in the parameter space.
Some natural basins with different shapes. Top : Black Canyon of the Gunnison, Colorado (CC BY-SA 4.0 Terry Foote). In this 2D loss-surface metaphor the model is brittle, small steps away from the local minimum force you up a steep cliff. Bottom : Death Valley, California (Public Domain, National Parks Gallery). An analogy for a degenerate, robust model where you can wiggle parameters without big changes in the loss.
Here’s a surprising math fact: neural networks learn like ice melts 1 .
When water is frozen, the water molecules sit in a crystal structure and have a limited number of valid ways to be arranged. The number of possible microstates, that is, the entropy, is relatively low. As the temperature rises, two competing pressures of energy (favoring ordered arrangements) and entropy (favoring disorder) shift their relative importance and, above $0^\circ\text{C}$, energy loses the tug of war: a phase transition happens. The entropy surges, because there are more possible molecule arrangements when the ice crystals loosen.
In the beginning of neural network training, the parameters can be in many different configurations at similar loss. Entropy is high. As the training progresses and the model starts to fit the data, the loss drops and so does the entropy: just like ice, the structure of the network becomes more ordered, at least initially.
Here’s what’s interesting: sometimes, well after the loss converges, the entropy starts climbing. The network drifts into a region of the parameter space where many different parameter configurations produce similar loss due to the degeneracies of the model. The ice melts. The model’s representation of the training data becomes less brittle, more stable. This often corresponds to better test set accuracy and a simpler representation of the training data. Later we will see a hands-on example.
Similarly to how entropy describes the multiplicity of microstates in physical systems, LLC characterizes the multiplicity of microstates at the current low-loss point in the parameter space. Its change over the training can indicate qualitative changes in the network’s weights. Indeed, Watanabe’s math predicts that neural networks go through similar phase transitions as physical systems. SLT’s free energy is the same mathematical object that describes phase transitions in thermodynamics and statistical mechanics.
The free energy, in two fields
$Z$ here is called partition function . In physics, the energy is a physical energy and the integral is over molecular states. In Bayesian statistics, the “energy” is the loss times the number of data points, and the integral is over model parameters. The free energy $F$ is the central object of both fields, and the energy-versus-entropy decomposition $F = \langle E \rangle - TS$ holds in both. Phase transitions, defined as non-analyticities of $F$ as a control parameter changes, are the same mathematical phenomenon whether the control parameter is temperature or sample size.
Here’s the most important sentence of this post: The model is naturally drawn towards simpler representations of the data, because those representations occupy a larger volume in the parameter space. Occam’s razor isn’t an externally imposed preference, it falls out of the math.
LLC measures qualitative jumps between regions in the loss landscape, something that is not visible from just the loss curve. A drop in LLC during training means that the model has arrived at a simpler, more stable representation of the training data. A concept “clicked”. The model generalized.
This matters for AI safety. Proponents of developmental interpretability (devinterp) believe that LLC can help us track when new abilities are acquired during training and how strongly they are established in the model. It may also help us detect when AI models try to game safety mechanisms.
Here’s the bad news: calculating the exact value of LLC is intractable for even relatively small neural networks. It involves a process from algebraic geometry called resolution of singularities applied to the loss landscape’s structure. The upper-bound time complexity is roughly $O\left(\mathrm{Ackermann}(d)\right)$ or $O\Bigl(\underbrace{2^{2^{\cdot^{\cdot^{\cdot^{2}}}}}}_{d\ \text{twos}}\Bigr)$. Please come up with your own popular science intractability analogy here (atoms in the Universe, seconds since the Big Bang, yada yada).
If you can’t evaluate an integral analytically, the next best thing is to rewrite it as an expectation and estimate it by sampling. How to do the sampling? The answer is good old gradient descent with a twist: added noise. Stochastic Gradient Langevin Dynamics (SGLD) works just like the standard optimizer step, but it adds some Gaussian noise to each parameter. In our hill climbing analogy: you randomly wander around a specific point to get a feel for the shape of the landscape. The gradient keeps you near the low-loss region but the random noise pushes you to explore.
As you “walk around”, you collect loss samples from trajectories around the area. You can try it yourself below.
The difference between these samples and your original loss gives you an idea of the shape of your valley: wide and flat or steep and V-shaped. LLC is the scaled average of these differences and captures the shape in a single number.
Here $\overline{L}_{\text{SGLD}}$ is the mean loss over SGLD samples. The denominator is the effective temperature where $\eta$ is the learning rate and $\varepsilon$ is the noise scale.
The formula for wandering around (SGLD) is as follows:
The first part looks like a gradient descent step on a tempered, regularized loss. The second part is the Gaussian noise. $\epsilon$ is the step size, $n$ is the dataset size, $\gamma$ is the localization strength, $\sigma$ is the noise strength. $\beta$ here means the inverse temperature which is $1/\log n$ for Watanabe’s standard estimator.
The optimal hyperparameters are subject of active research. In the toy examples below I use the standard estimator with $\gamma = 1.0$, $\sigma = 1$, and $\epsilon = \sqrt{2\eta}$ where $\eta$ is the learning rate. Each LLC estimate uses 5000 SGLD steps with the first 1000 discarded as burn-in.
I used the following simple experiments to develop my own intuitions on local learning coefficient (LLC) and Stochastic Gradient Langevin Dynamics (SGLD). First, we’ll look at a regular model and estimate its complexity. Then we’ll move on with a simple neural model with degeneracies. Finally, we’ll look at LLC evolution over the training. The notebook is available here .
Warm-up: Linear Regression (2 parameters)
A simple linear model $f(x) = wx + b$ trained on the dataset $y = 3x + 2$. The loss landscape is a convex bowl with a single minimum at $(w, b) = (3, 2)$. This is a regular model so we expect the LLC to approximately match $d/2$, that is, 1. We train the model for some steps, visualize the landscape, and run SGLD from the trained weights to get a rough LLC estimate.
# DATASET: Linear Regression (y = 3x + 2)
X = torch . linspace ( - 1 , 1 , 100 ). reshape ( - 1 , 1 )
Y = 3 * X + 2
# MODEL: 2 parameters (weight and bias)
model = nn . Linear ( 1 , 1 )
The plot below shows the loss landscape, the parameters after training near the minimum, and the trajectory the SGLD took (the circle marks the start, the star marks the end).
The LLC it calculated is $\approx 0.81$, which is not too far from what we expect from a two-parameter regular model: $d/2 = 2/2 = 1$.
The trajectory’s differences from base loss look like [-0.02838482 -0.02812932 -0.03070474 ... 0.0922314 0.08985042
0.09166867] . Their scaled average gives the estimated LLC. The scaling factor $n\beta = 100 \cdot \frac{1}{\ln 100} \approx 21.7147$ for the data set of 100 elements.
You can see how the chain explores the bowl without straying too far from its center. In practice, for more statistical stability, several chains (independent trajectories) are used to estimate LLC, but in this toy model we only run one.
Nonlinear Model: Single Hidden Neuron (4 parameters)
Here’s a more interesting case. A tiny MLP with one hidden neuron and tanh activation:
This has 4 learnable parameters $(w_1, b_1, w_2, b_2)$. Unlike the linear case, the loss landscape is non-convex and has symmetries (e.g., flipping the sign of $w_1$ and $w_2$ simultaneously gives the same function).
In the notebook I prov

[truncated]

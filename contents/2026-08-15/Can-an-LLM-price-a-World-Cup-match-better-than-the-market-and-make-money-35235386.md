---
source: "https://christian.bock.bio/posts/frontier_xg_worldcup/"
hn_url: "https://news.ycombinator.com/item?id=49313812"
title: "Can an LLM price a World Cup match better than the market (and make money)?"
article_title: "Can an LLM price a World Cup match better than the market (and make money)? · Christian Bock"
author: "cbock90"
captured_at: "2026-08-15T20:11:45Z"
capture_tool: "hn-digest"
hn_id: 49313812
score: 2
comments: 0
posted_at: "2026-08-15T20:04:07Z"
tags:
  - hacker-news
  - translated
---

# Can an LLM price a World Cup match better than the market (and make money)?

- HN: [49313812](https://news.ycombinator.com/item?id=49313812)
- Source: [christian.bock.bio](https://christian.bock.bio/posts/frontier_xg_worldcup/)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T20:04:07Z

## Translation

タイトル: LLM はワールドカップの試合に市場よりも高い価格を設定できますか (そして利益を得ることができますか)?
記事のタイトル: LLM はワールドカップの試合に市場よりも高い価格を設定できる (そして利益を得る) ことができますか? · クリスチャン・ボック
説明: 2 年前、私は GPT-4o に SportMonks の統計からユーロ 2024 のすべての結果を推測させました
(SoccerGPT): 勝者 51 名中 28 名が選ばれました
(55%) しかし、正確なゴール差はわずか 8 つ (16%) でした。 2026 年ワールドカップと新たな
推論モデルの生成は、より鋭い質問をするための口実でした。そうではない
どれだけ優れたモデルなのか
[切り捨てられた]

記事本文:
LLM はワールドカップの試合に市場よりも高い価格を設定する (そして利益を得る) ことができるでしょうか?
2 年前、私は GPT-4o に SportMonks の統計からユーロ 2024 のすべての結果を推測させました
( SoccerGPT ): 勝者 51 名中 28 名が選ばれました
(55%) しかし、正確なゴール差はわずか 8 つ (16%) でした。 2026 年ワールドカップと新たな
推論モデルの生成は、より鋭い質問をするための口実でした。そうではない
モデルがスコアラインをどれだけ正確に予測できるかだけでなく、モデルがスコアラインを推論できるかどうか
市場価格を上回る確率を獲得し、それをお金に変えるのです。それでこれ
モデルをオッズに盲目にして、実際の賭け金をモデル間のエッジに設定したとき
確率とカルシ価格、およびすべてのログが記録されます。
書類を作成し、公に賭けます。の
システムは完全に自動化されています。私は各分析を開始して、「ゴー」を出しただけです。
提案したに違いない。 121 回のベットでバンクロールは 255 米ドルから 771 米ドルに増加しました。この中で
この後、クロードが開発した方法論を順を追って説明し、その理由を見ていきます。
賭け金のほぼ半分を失いながら、現金は3倍になりました。
付属のリポジトリは GitHub にあります。
各試合の前に、私たちはチーム、最近のフォーム、メンバーなどの文書 $D$ を作成します。
決定論的に組み立てられたスターティングメンバーとトーナメントの背景が確認され、
価格が含まれていない。 LLM $f_\theta$ はこのドシエのみを読み取り、
分析から得られる分布、
$$\big(p,\ \mathrm{xG}^{1},\ \mathrm{xG}^{2}\big) \sim f_\theta(D),$$
チーム 1 の勝利の結果確率 $p = (p_1, p_{\mathrm{d}}, p_2)$、
引き分け、チーム 2 の勝利、予想ゴール $\mathrm{xG}^{1}$ と
各チームの得点 $\mathrm{xG}^{2}$。
試合が生み出す 1 つのもの、つまり最終ゴールから始めましょう。しましょう
$S = (S^1, S^2)$ はスコアラインになります。$S^1$ はチーム 1 のゴール数です。
得点と $S^2$ はチーム 2 の得点です。キックオフ前 $S

$ は不明なので、
これを、モデルが提供する分布から抽出された確率変数として扱います。
試合で提供されるすべての賭けは、それについての声明です
スコアライン: 勝ち/引き分け/勝ちの結果、試合合計 ($0.5$ ～ $3.5$ 以上/未満)
ゴール）、スプレッド、両チームの得点、クリーンシート、チームごとの合計など。
具体的には、各市場 $m$ は 1 で決済されるバイナリ カルシ契約です。
$S$ の条件が成立する場合はドル、成立しない場合は何もありません。それを書いてください
決済ルールとしての条件 $r_m(S^1, S^2) \in \{0, 1\}$、$1$ に正確に等しい
実現されたスコアラインがそれを満たした場合。たとえば、「2.5 ゴール以上」のマーケットの場合、
$r_m(S^1, S^2) = \mathbf{1}[S^1 + S^2 \ge 3]$。モデルの確率は、
条件が成立する場合は、単純に
$$q_m \ :=\ \Pr\big[\ r_m(S^1, S^2) = 1 \mid \mathrm{xG}^{1}, \mathrm{xG}^{2}\ \big].$$
したがって、すべては 1 つのオブジェクト、つまりスコアラインの分布に帰着します。
モデルの予想される目標が意味する $S$。
予想ゴールからスコアラインまで
見出しへのリンク
各チームのゴール数をモデルのレートと等しいポアソンとして扱います。
予想されるゴール、$\lambda_1 = \mathrm{xG}^1$ および $\lambda_2 = \mathrm{xG}^2$。あ
単一のポアソンはすでにゴール得点の基本的な形を捉えています: 低レートの山
ゼロと 1 つの目標の質量、一方、より高いレートはピークを外側に押し出し、
それを広めます（図1）。
スコアラインは、チームごとにそのような分布を 1 つ組み合わせたものです。 2の掛け算
独立したポアソンは明らかに出発点ですが、
ディクソンとコールズ (1997) はそれを示しました
最低スコアの価格を誤っている ($0$-$0$ および $1$-$1$ の抽選が少なすぎる) ため、適用します。
補正 $\tau$ は、スコアの最も低い 4 つのセルの重みを再計算します。
単一の依存関係パラメータ $\rho$:
$$P(S^1 = i, S^2 = j) \ \propto\ \frac{\lambda_1^{i} e^{-\lambda_1}}{i!} \cdot \frac{\lam

bda_2^{j} e^{-\lambda_2}}{j!} \cdot \tau(i, j),$$
$$\tau(i, j) = \begin{cases} 1 - \lambda_1 \lambda_2 \rho & (i, j) = (0, 0) \\ 1 + \lambda_1 \rho & (i, j) = (0, 1) \\ 1 + \lambda_2 \rho & (i, j) = (1, 0) \\ 1 - \rho & (i, j) = (1, 1) \\ 1 & \text{そうでない場合。} \end{cases}$$
$\rho = -0.10$ を設定し、各サイド 10 ゴールで切り捨て、行列を正規化します。
合計して 1 になります。モデルの出力は現在、2 つのソースから市場の価格を決定しています。の
勝ち/引き分け/勝ちのマーケットは確率 $p_1$, $p_{\mathrm{d}}$,
$p_2$ を直接勝ち取り、ノックアウトでチーム 1 が勝ち進む確率は次のとおりです。
$p_1 + \tfrac{1}{2} p_{\mathrm{d}}$ (引き分けた試合は延長戦に移行し、
ペナルティは均等に分割されます）。他のすべての市場 (合計、スプレッド、両チームの合計、スプレッド)
スコア、チーム合計）は、上で構築された 2 つのスコアライン グリッドから価格設定されます。
予想されるゴール。ルール $r_m$ は決定的であり、スコアラインのみが
ランダムなので、確率 $q_m$ はグリッド上のそのルールの期待値です。
セルの加重和として計算され、
$$q_m = \mathbb{E}[r_m(S^1, S^2)] = \sum_{i, j} r_m(i, j) \cdot P(S^1 = i, S^2 = j).$$
$r_m$ は値 $0$ と $1$ のみを取るため、最初の等式が成り立ちます。
期待値は、まさに $1$ に等しい確率です。
1 つの推論がパスし、2 点の推定値だけが得られるため、価格はおよそ 24 個になります。
一気に契約。図 2 は、実際の一致におけるそのようなグリッドの 1 つであるジョイントを示しています。
市場確率ごとのスコアライン分布が読み取られます。
いつベットするか: エッジとアンカー
見出しへのリンク
市場に参入するのは今だけです。カルシでは、市場 $m$ の YES 契約には、
売値 $c_m \in (0, 1)$: $c_m$ を支払い、イベントが発生すると 1 ドルを受け取ります。
この価格は、モデルの価格が $q_m$ であるのと同じイベントに対する市場独自の見解です。
したがって、それらの間のギャップが私たちの優位性です、
モデルが正の場合

契約が安すぎる。私たちは両方の側面を評価します
すべての市場の場合: NO 側が独自のアスクで $1 - q_m$ の確率で勝ちます。
価格なので、同じルールで $q_m$ を $1 - q_m$ に置き換えてカバーします。
モデルが間違っている可能性があるため、ポジティブエッジだけでは十分ではありません。
そう。そこで私たちはセカンドオピニオン、つまり鋭いアンカーを求めます。濃い液体の場合
市場（価格の裏に数百万ドルの取引量があるポリマーケット）の相場
同じ結果の場合、その暗黙の確率 $q^{\mathrm{Poly}}_m$ を読み取り、監視させます
モデルには 2 つの方法があります。
方向。アンカーはサイドも過小評価されていると見なす必要があります。
$q^{\mathrm{Poly}}_m - c_m \ge \delta$ ($\delta = 0.01$)。市場の意見が大きく異なる場合
モデルの方向性については、モデルが間違っているとみなして合格します。
保守的なサイジング。私たちはアンカーを超えてモデルに頼ることはありません
がサポートされるため、2 つの確率のうち小さい方のエッジを測定します。
$\min(q_m, q^{\mathrm{Poly}}_m) - c_m$。
ベットトリガーは端にある単一のバーであり、何もない場合はより高く設定されます。
寄りかかるアンカー:
$$\text{賭けます } m \iff \begin{cases} \min(q_m, q^{\mathrm{Poly}}_m) - c_m \ge 0.06 & \text{アンカーあり、} \\ q_m - c_m \ge 0.09 & \text{アンカーなし。} \end{cases}$$
アンカーのないバーは、独立したものがチェックしていないため、より厳密になります。
番号。クリアしてもカルシの取引手数料を上回らなければなりません。
$\lceil 0.07 \cdot C \cdot c_m (1 - c_m) \rceil$ セント ($C$ 契約前)
それは実際のエッジとしてカウントされます。
賭け金: フラクショナル ケリー
見出しへのリンク
賭け金を選択することは仕事の半分に過ぎません。残りの半分はサイズです。 $X$ を
現在のバンクロールと、対象となるベットに賭ける賭け金 $x$ です。 $w$ と書きます
賭けが勝つ確率 (賭け金からの保守的なサイジング確率)
最後のセクション、$\min(q_

m, q^{\mathrm{Poly}}_m)$ (アンカー時) および $q_m$
それ以外の場合)、契約ごとのコストは $c$ となります。
ケリー基準
(Kelly、1956) は、社会の長期的な成長率を最大化する賭け金を選択します。
バンクロール、つまり期待される富の対数。端数のステーキング
$\text{frac} = x / X$、勝利した契約はネットオッズを返します
$\text{odds} = (1 - c) / c$ 賭け金単位あたり、敗者は賭け金を失います。
1 回のベットでバンクロールを $1 + \text{frac} \cdot \text{odds}$ 倍にします。
確率 $w$、それ以外の場合は $1 - \text{frac}$ による。予想される対数増加は、
$$g(\text{frac}) = w \ln(1 + \text{frac} \cdot \text{odds}) + (1 - w) \ln(1 - \text{frac}).$$
$g’(\text{frac}) = 0$ と設定すると、成長を最大化する割合が得られます。
ケリー分数 $\text{frac}_K$、
$$\frac{w \cdot \text{odds}}{1 + \text{frac} \cdot \text{odds}} = \frac{1 - w}{1 - \text{frac}} \quad\Longrightarrow\quad \text{frac}_K = \frac{w - c}{1 - c},$$
その結果、下側 $1 - c$ を超えるエッジ $w - c$ になります。しかし、完全なケリーは、勝利確率 $w$ が正確にわかっており、
成長曲線は右側に急勾配です (図 3): 賭け金は最適値をはるかに超えており、
成長率は低下し、ケリー率の2倍近くのゼロに達し、マイナスに転じる
超えて。したがって、誇張された $w$ は危険です。
その崖に向かって右方向に杭を打つことをお勧めします。図では、勝つ賭け金
確率は 55% ですが、65% が 10% から
30%、赤字の領域に深く入り込み、実際のエッジをゆっくりとしたブリードに変換します。
$w$ はノイズの多いモデル推定値であるため、意図的に $w$ のみをステークします。
ケリーの4分の1、それは私たちをその端からかなり左に保ち、そして私たちはどれかをキャップします
バンクロールの 12% で賭けます:
$$x = \min\left( \tfrac{1}{4} \cdot \frac{w - c}{1 - c} \cdot X, \ \ 0.12 \cdot X \right

).$$
次に、整数の契約 $C = \lfloor x / c \rfloor$ を購入し、
四捨五入された賭け金が上記の手数料をまだクリアしている場合にのみ賭けます。
最後にもう一度調整。チームまたは試合を共有するベットは一緒に上がったり下がったりします。
したがって、それぞれを個別にサイジングすると、静かにリスクが集中します。を追加します。
クオリファイングは最高エッジファーストでベットし、エクスポージャの上限をバンクロールの 15% に設定します。
単一の相関グループと合計 50% がデプロイされ、任意のものをトリミングまたはドロップします
上限を突破してしまいます。
トーナメントを通じて、グループステージから決勝まで、システムは 121 に定着しました。
以下に示すように、賭け金は 255 USD のバンクロールから 771 USD まで増加しました。
リターンは実際のお金ですが、それは 1 つの短くて分散の大きなサンプルであり、
このセクションの残りの部分は、そのサンプルで何がサポートされ、何がサポートされないかについて説明します。
記録だけではほとんど証明されません。 62勝59敗の記録は命中率51.2%に相当し、
95% ウィルソン信頼区間は 42.4% ～ 60.0% です。その間隔には以下が含まれます
50%、つまり勝ち負けだけでは戦略とコインの区別がつかない
フリップします。これは仕様によるものです。目的は、より頻繁に正しくなることではなく、正しくなるようにすることです。
価格が間違っている場合、エッジの証拠は価格と
生のレコードではなく、キャリブレーションです。
実現された優位性は価格にあります。 121 のベット全体で、私たちが行った契約は
平均エントリー価格は 51.2% の確率で当社に有利に解決されました。
1ドルあたり37.8セントでした。 0.512 で発生した結果に対して 0.378 を支払います。
+20.8% の売上高利益率のすべての源は時間です。
モデルは絶対的に適切に調整されています。記載された平均値
賭け金に対する確率は 0.523 で、約 1 パーセントポイント上にあります。
実現勝率は 0.512 なので、平均すると、記載されている確率はどのように一致しますか
多くの場合、その賭けは実際に勝ちました。ベットごとのスコア
Brier スコア、平均二乗距離を使用

指定された確率の間にある場合
$q$ と $N = 121$ のベットに対する $0/1$ の結果 $o$、
$$\text{ブリエ} = \frac{1}{N} \sum_{i=1}^{N} (q_i - o_i)^2,$$
モデルの確率は 0.234 に達し、情報のない確率の 0.250 をわずかに下回ります。
相変わらずの推測。
残りは慎重に読みます。市場価格のスコアは 0.251 です。
ベットしますが、その差はそれほど大きくありません (ブートストラップ 95% 間隔)
$[-0.011,\ +0.045]$)、すでに持っている契約のみに賭けるため、偏りもありません
割安と判断された。リターンも騒がしいです: に対する総賞金は 1,774 USD
1,258 米ドルの損失が正味でわずか 516 米ドルになり、5 つの最善策が損失の 83% を占めます。
それ。 5 つの結果が異なれば、見出しもまったく違ったものになるでしょう。
LLM に試合ごとの書類から勝ち/引き分け/勝ちの確率まで推論させます。
予想される目標を、ディクソン・コールズで 20 の市場の価格に変えた
ポアソン、そして私たちの価格がカルシとポリマーケットのアンカーを上回るところにのみ賭けます
クォーターケリーによるサイジングに同意しました。オッズを無視した推論で、マッチの価格を設定した
賭け金のほぼ半分を失いながら、255 USD を 771 USD に増やすには十分です。
エッジはヒット率ではなく、価格にありました。なぜなら、そのリターンは小さくてノイズの多いサンプルであり、
少数の賭け、私たちが最も信頼するシグナル

[切り捨てられた]

## Original Extract

Two years ago I let GPT-4o guess every Euro 2024 result from SportMonks stats
(SoccerGPT): it called 28 of 51 winners
(55%) but the exact goal difference in only 8 (16%). The 2026 World Cup and a new
generation of reasoning models were the excuse to ask a sharper question. Not
just how well a model
[truncated]

Can an LLM price a World Cup match better than the market (and make money)?
Two years ago I let GPT-4o guess every Euro 2024 result from SportMonks stats
( SoccerGPT ): it called 28 of 51 winners
(55%) but the exact goal difference in only 8 (16%). The 2026 World Cup and a new
generation of reasoning models were the excuse to ask a sharper question. Not
just how well a model can predict a scoreline, but whether it can reason to a
probability that beats the market’s price, and turn that into money. So this
time I blinded the model to the odds, sized real bets on the edge between its
probability and the Kalshi price, and logged every
dossier and bet publicly . The
system is fully automated; I only started each analysis and gave the “go” on the
bets it proposed. Over 121 bets the bankroll grew from 255 to 771 USD. In this
post we walk through the methodology, which Claude developed, and look at why we
tripled our cash while losing almost half of our bets.
The accompanying repository can be found at GitHub .
Before each match, we compile a dossier $D$: the squads, recent form, the
confirmed starting XI, and tournament context, assembled deterministically and
containing no prices. The LLM $f_\theta$ reads only this dossier and defines a
distribution over analyses from which we draw one,
$$\big(p,\ \mathrm{xG}^{1},\ \mathrm{xG}^{2}\big) \sim f_\theta(D),$$
the outcome probabilities $p = (p_1, p_{\mathrm{d}}, p_2)$ for a team 1 win, a
draw, and a team 2 win, and the expected goals $\mathrm{xG}^{1}$ and
$\mathrm{xG}^{2}$ that each team scores.
Let us start from the one thing a match produces: its final goals. Let
$S = (S^1, S^2)$ be the scoreline , where $S^1$ is the number of goals team 1
scores and $S^2$ the number of goals team 2 scores. Before kickoff $S$ is unknown, so we
treat it as a random variable drawn from a distribution the model will supply.
Every wager offered on the match is a statement about that
scoreline: the win/draw/win outcome, match totals (over/under $0.5$ to $3.5$
goals), spreads, both teams to score, clean sheets, per-team totals, and so on.
Concretely, each market $m$ is a binary Kalshi contract that settles at one
dollar if its condition on $S$ holds and nothing if it does not. Write that
condition as a settlement rule $r_m(S^1, S^2) \in \{0, 1\}$, equal to $1$ exactly
when the realized scoreline satisfies it. For instance, for the “over 2.5 goals” market,
$r_m(S^1, S^2) = \mathbf{1}[S^1 + S^2 \ge 3]$. The model’s probability that the
condition holds is then simply
$$q_m \ :=\ \Pr\big[\ r_m(S^1, S^2) = 1 \mid \mathrm{xG}^{1}, \mathrm{xG}^{2}\ \big].$$
Everything therefore reduces to one object: the distribution of the scoreline
$S$ that the model’s expected goals imply.
From expected goals to a scoreline
Link to heading
We treat each team’s goal count as Poisson with a rate equal to the model’s
expected goals, $\lambda_1 = \mathrm{xG}^1$ and $\lambda_2 = \mathrm{xG}^2$. A
single Poisson already captures the basic shape of goalscoring: a low rate piles
the mass on zero and one goal, while a higher rate pushes the peak outward and
spreads it (Figure 1).
The scoreline combines one such distribution per team. Multiplying two
independent Poissons is the obvious starting point, but
Dixon and Coles (1997) showed it
misprices the lowest scores (too few $0$-$0$ and $1$-$1$ draws), so we apply
their correction $\tau$, which reweights the four lowest-score cells through a
single dependency parameter $\rho$:
$$P(S^1 = i, S^2 = j) \ \propto\ \frac{\lambda_1^{i} e^{-\lambda_1}}{i!} \cdot \frac{\lambda_2^{j} e^{-\lambda_2}}{j!} \cdot \tau(i, j),$$
$$\tau(i, j) = \begin{cases} 1 - \lambda_1 \lambda_2 \rho & (i, j) = (0, 0) \\ 1 + \lambda_1 \rho & (i, j) = (0, 1) \\ 1 + \lambda_2 \rho & (i, j) = (1, 0) \\ 1 - \rho & (i, j) = (1, 1) \\ 1 & \text{otherwise.} \end{cases}$$
We set $\rho = -0.10$, truncate at ten goals per side, and normalize the matrix
to sum to one. The model’s outputs now price the markets from two sources. The
win/draw/win markets take the probabilities $p_1$, $p_{\mathrm{d}}$,
$p_2$ directly, and in a knockout the probability that team 1 advances is
$p_1 + \tfrac{1}{2} p_{\mathrm{d}}$ (a drawn match goes to extra time and
penalties, split evenly). Every other market (totals, spreads, both teams to
score, team totals) is priced from the scoreline grid built above from the two
expected goals. The rule $r_m$ is deterministic and only the scoreline is
random, so the probability $q_m$ is the expectation of that rule over the grid,
computed as a weighted sum of its cells,
$$q_m = \mathbb{E}[r_m(S^1, S^2)] = \sum_{i, j} r_m(i, j) \cdot P(S^1 = i, S^2 = j).$$
The first equality holds because $r_m$ takes only the values $0$ and $1$, so its
expectation is exactly the probability that it equals $1$.
One reasoning pass, just two point estimates, therefore prices roughly two dozen
contracts at once. Figure 2 shows one such grid for a real match, the joint
scoreline distribution every market probability is read off.
When to bet: edge and anchor
Link to heading
Only now does the market enter. On Kalshi a YES contract for market $m$ has an
ask price $c_m \in (0, 1)$: pay $c_m$, receive one dollar if the event happens.
That price is the market’s own view of the same event our model prices at $q_m$,
so the gap between them is our edge ,
positive when the model thinks the contract is too cheap. We evaluate both sides
of every market: the NO side wins with probability $1 - q_m$ at its own ask
price, so the identical rule covers it with $q_m$ replaced by $1 - q_m$.
A positive edge is not enough, because the model can be wrong, and confidently
so. So we ask for a second opinion, a sharp anchor . When a deep, liquid
market (Polymarket, with millions of dollars of volume behind its prices) quotes
the same outcome, we read off its implied probability $q^{\mathrm{Poly}}_m$ and let it police
the model two ways:
Direction. The anchor must also see the side as underpriced,
$q^{\mathrm{Poly}}_m - c_m \ge \delta$ with $\delta = 0.01$. If a market that deep disagrees
with the model’s direction, we take the model to be wrong and pass.
Conservative sizing. We never lean on the model beyond what the anchor
supports, so we measure the edge on the smaller of the two probabilities,
$\min(q_m, q^{\mathrm{Poly}}_m) - c_m$.
The bet trigger is then a single bar on the edge, set higher when there is no
anchor to lean on:
$$\text{bet on } m \iff \begin{cases} \min(q_m, q^{\mathrm{Poly}}_m) - c_m \ge 0.06 & \text{with an anchor,} \\ q_m - c_m \ge 0.09 & \text{without one.} \end{cases}$$
The unanchored bar is stricter precisely because nothing independent has checked
the number. Anything that clears must still beat Kalshi’s trading fee, about
$\lceil 0.07 \cdot C \cdot c_m (1 - c_m) \rceil$ cents on $C$ contracts, before
it counts as a real edge.
How much to bet: fractional Kelly
Link to heading
Choosing the bets is only half the job; the other half is size. Let $X$ be the
current bankroll and $x$ the stake we place on a qualifying bet. Write $w$ for
the probability that the bet wins (the conservative sizing probability from the
last section, $\min(q_m, q^{\mathrm{Poly}}_m)$ when anchored and $q_m$
otherwise) and $c$ for its cost per contract.
The Kelly criterion
(Kelly, 1956) picks the stake that maximises the long-run growth rate of the
bankroll, that is, the expected logarithm of wealth. Staking a fraction
$\text{frac} = x / X$, a winning contract returns net odds
$\text{odds} = (1 - c) / c$ per unit staked while a loser forfeits the stake, so
one bet multiplies the bankroll by $1 + \text{frac} \cdot \text{odds}$ with
probability $w$ and by $1 - \text{frac}$ otherwise. The expected log-growth is
$$g(\text{frac}) = w \ln(1 + \text{frac} \cdot \text{odds}) + (1 - w) \ln(1 - \text{frac}).$$
Setting $g’(\text{frac}) = 0$ gives the growth-maximising fraction, which we call
the Kelly fraction $\text{frac}_K$,
$$\frac{w \cdot \text{odds}}{1 + \text{frac} \cdot \text{odds}} = \frac{1 - w}{1 - \text{frac}} \quad\Longrightarrow\quad \text{frac}_K = \frac{w - c}{1 - c},$$
the result that it is the edge $w - c$ over the downside $1 - c$. But full Kelly assumes the win probability $w$ is known exactly, and the
growth curve is steep on the right (Figure 3): stake much past the optimum and
growth falls, hitting zero near twice the Kelly fraction and turning negative
beyond. An overstated $w$ is therefore dangerous, because it pushes the
recommended stake rightward, toward that cliff. In the figure, a bet that wins
55% of the time but is sized as if it won 65% triples the fraction from 10% to
30%, deep into loss-making territory, converting a real edge into a slow bleed.
Because our $w$ is a noisy model estimate, we deliberately stake only a
quarter of Kelly, which keeps us well left of that edge, and we cap any one
bet at 12% of bankroll:
$$x = \min\left( \tfrac{1}{4} \cdot \frac{w - c}{1 - c} \cdot X, \ \ 0.12 \cdot X \right).$$
We then buy a whole number of contracts, $C = \lfloor x / c \rfloor$, and keep the
bet only if that rounded stake still clears the fee above.
One last adjustment. Bets that share a team or a match rise and fall together,
so sizing each in isolation would quietly concentrate risk. We add the
qualifying bets highest-edge-first, capping exposure at 15% of bankroll on any
single correlated group and 50% deployed in total, and trim or drop whatever
would breach a cap.
Over the tournament, from the group stage to the final, the system settled 121
bets and grew a 255 USD bankroll to 771 USD as shown below.
The return is real money, but it is one short, high-variance sample, and the
rest of this section is about what that sample does and does not support.
The record alone proves little. A 62-59 record is a 51.2% hit rate, whose
95% Wilson confidence interval runs from 42.4% to 60.0%. That interval contains
50%, so on wins and losses alone we cannot distinguish the strategy from a coin
flip. This is by design: the aim is not to be right more often, but to be right
when the price is wrong, so the evidence for an edge has to come from prices and
calibration, not the raw record.
The realized edge is in the prices. Across the 121 bets, the contracts we
bought resolved in our favour 51.2% of the time, while their average entry price
was 37.8 cents on the dollar. Paying 0.378 for outcomes that occur 0.512 of the
time is the whole source of the +20.8% return on turnover.
The model is well-calibrated in absolute terms. Its average stated
probability over the placed bets, 0.523, sits about one percentage point above
the realized win rate, 0.512, so on average its stated probabilities match how
often those bets actually won. Scored per bet
with the Brier score, the mean squared distance between a stated probability
$q$ and the $0/1$ outcome $o$ over $N = 121$ bets,
$$\text{Brier} = \frac{1}{N} \sum_{i=1}^{N} (q_i - o_i)^2,$$
the model’s probabilities reach 0.234 , slightly below the 0.250 of an uninformative
constant guess.
The rest we read cautiously. The market’s prices score 0.251 on the same
bets, but that gap is neither significant (bootstrap 95% interval
$[-0.011,\ +0.045]$) nor unbiased, since we only bet contracts we had already
judged underpriced. The return is noisy too: gross winnings of 1,774 USD against
1,258 USD of losses net to just 516 USD, and the five best bets make up 83% of
that. A different five results and the headline would look very different.
We let an LLM reason from a per-match dossier to a win/draw/win probability and
expected goals, turned those into prices for two dozen markets with a Dixon-Coles
Poisson, and bet only where our price beat Kalshi’s and a Polymarket anchor
agreed, sizing by quarter-Kelly. Reasoning blind to the odds, it priced matches
well enough to grow 255 USD into 771 USD while losing almost half of its bets:
the edge lived in the price, not the hit rate. Because that return is a small, noisy sample carried by a
handful of bets, the signal we trust most

[truncated]

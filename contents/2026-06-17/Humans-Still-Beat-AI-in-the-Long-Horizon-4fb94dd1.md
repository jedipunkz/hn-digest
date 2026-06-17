---
source: "https://joyemang33.github.io/blog/2026/humans-dont-just-sample/"
hn_url: "https://news.ycombinator.com/item?id=48564319"
title: "Humans Still Beat AI in the Long Horizon"
article_title: "Humans Still Beat AI in the Long Horizon: Revisiting Test-Time Scaling in the Agent Era | Qiuyang Mang"
author: "mxwsn"
captured_at: "2026-06-17T01:01:54Z"
capture_tool: "hn-digest"
hn_id: 48564319
score: 1
comments: 0
posted_at: "2026-06-17T00:44:51Z"
tags:
  - hacker-news
  - translated
---

# Humans Still Beat AI in the Long Horizon

- HN: [48564319](https://news.ycombinator.com/item?id=48564319)
- Source: [joyemang33.github.io](https://joyemang33.github.io/blog/2026/humans-dont-just-sample/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T00:44:51Z

## Translation

タイトル: 長い地平線でも人間は依然として AI に勝つ
記事のタイトル: 長い地平線でも人間は依然として AI に勝つ: エージェント時代のテスト時間のスケーリングの再検討 |クイヤン・マン
説明: エージェントは、試行、観察、修正することでテスト時の計算に費やすことができます。私たちは繰り返しのサンプリングから Elo 基準を導き出し、2022 年の 2 週間のコーディング マラソンでは、現在のエージェントが 24 時間以内に頭打ちになる一方、上位の人間は改善を続けることを示しました。

記事本文:
Qiuyang Mang ナビゲーションを切り替える
人間は長期的にも AI に勝つ: エージェント時代のテスト時間のスケーリングを再考する
2026年6月16日• Qiuyang Mang 1 、 Kaiyuan Liu 2 、 Bo Peng 3 、 Shreyas Pimpalgaonkar 4 、 Alex Dimakis 1,4 、 Alvin Cheung 1
1 カリフォルニア大学バークレー校 · 2 ワシントン大学 · 3 プリンストン大学 · 4 ビスポークラボ
2026 · テスト時間のスケーリング LLM エージェントのスケーリングの法則 · 研究
TL;DR.エージェントは、試行、観察、修正することでテスト時の計算に費やすことができるため、エージェントの利益がより優れた内部戦略から来るのか、それともサンプリングの繰り返しに近いものから来るのかを尋ねます。単純な Elo 基準線を導き出します。つまり、繰り返されるサンプリングは、ログ テスト時間の計算において線形です。 2022 年の 2 週間のコーディング マラソンでは、現在のエージェントは 24 時間以内に頭打ちになりますが、トップの人間は公式の 2 週間で向上し続けます。ここで重要なのは、長期的なテスト時間への適応に関しては人間のほうがはるかに優れており、エージェントの戦略には改善の余地がたくさんあるということです。
エージェントが本質的なテスト時間戦略をもたらす
OpenAI の o1 レポートは、テスト時のコンピューティングを増やすことでモデルのパフォーマンスを向上できることを示しました。その後、特にコードや数学などの検証可能なタスクに関する多くの論文が続きました (Snell et al.、Large Language Monkeys、Noam Brown の最近の投稿)。一般的なプロットは、成功率と試行回数のログ、またはテスト時間の計算のログです。これらの曲線は、飽和する前に超直線的に上昇することがよくあります。
Large Language Monkeys による、サンプル数の増加に伴う Oracle Verifier による MATH のカバレッジ。
これらの調査では、外部のテスト時間戦略に基づいてモデルのパフォーマンスを測定します。戦略はモデルの外側で固定されます。多くの候補解をサンプリングし、検証者でチェックし、pass@k またはカバレッジを報告します。
ただし、エージェントはこの設定を変更します。実行中に、エージェントは解決策を試し、状況を観察できます。

結果を見て、何が失敗したかを反省し、次の試みを修正します。これにより、私たちが研究する疑問が生じます。エージェントがより多くのテスト時間のコンピューティングを使用して向上するとき、エージェントはより優れたテスト時間戦略を使用しているのでしょうか、それとも主に繰り返しサンプリングを再現しているのでしょうか?
繰り返しのサンプリングはモデルの外部で修正されますが、エージェントは実行内でフィードバックを使用できます。
テスト時間のスケーリングのための単純なモデル
まず、通常の pass@k 曲線の背後にある最も単純なモデル、つまり反復サンプリングを書き留めます。モデルは、各試行を同じ連続スコア分布からの独立した抽選として扱います。 1 つのタスクについて、 \(X\) を 1 つのサンプルのスコア、 \(\tau\) を成功のしきい値とします。その後、1 つのサンプルが確率で成功します
\(k\) 個の独立したサンプルの場合、 pass@k は次のようになります。
複数のタスクを含むデータセットの場合、通常のテスト時間スケーリング曲線は、タスク全体でこの量を平均します。これにより、平均 pass@k またはカバレッジの曲線がサンプル数の関数として得られます。
しかし、この評価はエージェントにとって厄介なものです。通常、エージェントはタスクを解決すると停止するため、成功後にさらに独立したサンプルを要求し続けるのは不自然です。 FrontierCS などのオープンエンド タスクの場合は、タスク自体のスコアによって実行を比較することもできます。しかし、生のスコアの増加を解釈するのは困難です。 AlphaEvolve によって調査されたサークル パッキング タスクでは、目標値を 1 から 2 に改善するのは簡単ですが、2.35 から 2.36 に改善するにはさらに困難な改善が必要になる場合があります。スコアの数値だけでは、能力がどの程度変化したかはわかりません。したがって、より単純な質問をする比較が必要です。同じタスクに対して、ある候補者が別の候補者よりも多くのテスト時間の計算を費やした場合、どれくらいの頻度でより良い答えが得られるでしょうか?
同じ反復サンプリング モデルに従うと、これはペアごとの質問になります。候補者が 1 人いる場合

\(k_a\) 回独立して試行し、もう 1 人が \(k_b\) 回独立して試行する場合、ペアごとの勝利確率は次のようになります。
ここで、 \(X_i\) と \(X'_j\) は、同じタスクに対する独立した試行です。
ここからが便利な部分です。 Bradley-Terry モデルは、ペアごとの勝率を 1 次元の強さスケールに変換します。候補 \(a\) の BT 対数強度 \(\theta_a\) があり、候補 \(b\) の BT 対数強度 \(\theta_b\) がある場合、
サンプリングを繰り返した場合、上記のペアごとの勝利確率は次のように正確に一致することがわかります。
したがって、反復サンプリングは、テスト時間の対数計算において Elo が線形であるテスト時間戦略です。これは基準線を示すので非常に役立ちます。エージェントの固有のテスト時間戦略を判断するには、テスト時間のコンピューティングが増加するにつれて Elo 曲線をプロットし、この線と比較します。曲線が直線より上、下、または直線に近い場合、エージェントの戦略は反復サンプリングよりも優れている、劣っている、または同等です。
サンプリングを繰り返すと、Elo 対 log テスト時間の計算で線形の基準線が得られます。
エージェントは奮闘するが、人間はサンプリング以上のことを行う
この基準線を手に入れることで、実際の長期的なタスクで何が起こるかを尋ねることができます。私たちはエージェントの軌跡を反復サンプリング ラインと比較し、同じタスクに取り組んでいるトップの人間と比較してどうなるのかも尋ねます。
私たちは、2022 年に長期にわたるコーディングと最適化のコンテストである AtCoder Heuristic Contest 014: RectJoin を研究します。人間の出場者はアルゴリズムを作成し、コンテスト中に多くのコード ソリューションを提出することができ、リーダーボードには彼らが達成した最高スコアが記録されます。このコンテストは最新のコーディング エージェントが広く利用可能になる前に開催されたため、これらの人間の軌跡は AI エージェントによって支援されていません。このタスクには制限がありません。既知の最適な解決策はなく、スコアの向上と悪化だけが存在します。
高いレベルでは、RectJoin スター

グリッド上にマークされたドットを持つ ts。ソルバーは、有効な軸に沿った長方形または 45 度の長方形を形成する 3 つの既存のドットと 1 つの空のグリッド ポイントを繰り返し選択し、新しい点をマークして長方形の境界を描画します。目的は、最終的にマークされたドットの加重スコアを最大化することであり、中心から遠いドットほど重みが大きくなります。
AtCoder ヒューリスティック コンテスト 014 の RectJoin の視覚化。
人間の軌跡。人間の場合、最終順位から上位 10 位の出場者と上位 50 位の出場者という 2 つのグループを研究します。各チェックポイントで、プレフィックスのベスト スコアを計算します。これにより、各チェックポイントの人間グループごとに 1 つのベクトルが与えられ、各エントリはその時点までの競技者の最高スコアになります。
エージェントの設定。エージェント向けのコンテスト設定を再作成します。 FrontierCS 評価レイヤーを使用して、各エージェントは連続 24 時間ループで実行されます。人間の競技者と同じように、候補者を提出し続け、スコアを観察し、次の試技を修正することができます。早期に停止した場合は、新たに実行を開始するのではなく、同じセッションを再開します。
2 つのエージェント システム (Claude Code を使用した Claude Opus 4.6、および Codex を使用した GPT-5.5) をテストします。各エージェント システムについて、5 つの独立した 24 時間トライアルを実行します。私たちが注目するすべての実時間チェックポイントで、各試行からのプレフィックス最高スコアを収集します。これにより、各チェックポイントの各エージェント システムに長さ 5 のベクトルが与えられます。
これらのベクトルを取得すると、ペアごとの勝率を直接推定できます。たとえば、1 つのベクトルには 24 時間時点の上位 10 人の人間の出場者のプレフィックス最高スコアが含まれ、別のベクトルには 5 時間のクロード コード トライアルからのプレフィックス最高スコアが含まれている可能性があります。私たちは、24 時間のトップヒューマンのスコアが 5 時間のクロード コードのスコアをどのくらいの頻度で上回るかを知りたいと考えています。 2 つのベクトルが
\(u\) が次のような経験的確率を定義します。

すべてのペアを比較することで \(v\) に勝ちます。
これらのペアごとの勝率は、Elo フィットへの入力となります。
L-BFGS を使用して、これらのペアごとの勝率から Elo 評価を当てはめます。すべてのエージェント システムと人間のグループは、ブラッドリーとテリーの 1 つの結合当てはめに配置されるため、それらの評価は直接比較できます。エージェントの場合は、実行からの 24 時間の軌跡を使用します。人間の場合は、公式の 2 週間の完全な軌跡を使用します。
上記で紹介した長期的なコーディング タスクである RectJoin 上のトップの人間とエージェント システムの Elo 軌跡。
結果は鮮明です。エージェントは最初の数時間で急速に改善しますが、1 つのエージェントのトライアルで最大 1 億トークンを使用できるにもかかわらず、Elo 曲線は 24 時間経過すると平坦になります。トップレベルの人間は、最初はよりゆっくりと上達しますが、何日もかけて上昇を続け、最終的にはエージェント システムを大差で追い越します。これは、現在のエージェントは早い段階でスプリントすることができるが、強力な人間の競技者が長時間のコンテスト中に使用する長期的なテスト時間への適応がまだ欠けていることを示唆しています。
各参加者システムに対して、よりローカルな質問をすることもできます。一度に 1 つのシステムだけを観察する場合、反復サンプリング モデルは、そのシステムが単にそのシステム自身の出力分布からより多くの独立したサンプルを抽出した場合に何が起こるかを示す基準線を提供します。観察された Elo 曲線をこの線と比較すると、システムのテスト時間戦略がそれ自体から繰り返しサンプリングするよりも効率的であるかどうかが分かります。
モデルによって暗示される反復サンプリング基準線と比較したソースごとの Elo 曲線。
灰色の破線は、反復サンプリングの基準です。この線より上の曲線は、同じソース分布からの独立したサンプリングよりも速く Elo を獲得します。その下のカーブでは Elo の増加が遅くなります。エージェント システムは、24 年の終わりまでにこの基準に対してサブリニアになります。

-時間の実行。対照的に、人間の曲線は長期にわたって超直線的になります。人間はただサンプリングするだけではありません。現在のエージェントが長期的なテスト時間の拡張に取り組むには、まだ長い道のりがあります。
エージェントのテスト時間のスケーリングに関する要点
エージェントには独自のテスト時間戦略があります。コンピューティングを増やすとパフォーマンスが向上するかどうか、またその向上をもたらした戦略は何かを評価する必要があります。
繰り返しサンプリングを基準線として使用します。エージェントの Elo がログ計算に応じて直線的に増加する場合、エージェントはサンプリングを繰り返しているだけである可能性があります。そのラインからの逸脱がシグナルです。
人間を長期的な参考資料として扱いましょう。トップクラスの人間は依然として長期にわたって適応的な改善を示しており、これはエージェントによるテスト時間のスケーリングに有用な目標を与えてくれます。
長期にわたる自由な課題をもっと研究してください。エージェントがまだ不足している部分を理解するには、より多くのタスク、より長い実行軌跡、および慎重な障害分析が必要です。
私たちのフルペーパーは間もなく公開されます。それまでの間、このブログ投稿が役立つと思われた場合は引用してください。議論については、 qmang@berkeley.edu または lky04@cs.washington.edu までお問い合わせください。

## Original Extract

Agents can spend test-time compute by trying, observing, and revising. We derive an Elo reference for repeated sampling, then show that in a 2022 two-week coding marathon, current agents plateau within 24 hours while top humans keep improving.

Qiuyang Mang Toggle navigation About
Humans Still Beat AI in the Long Horizon: Revisiting Test-Time Scaling in the Agent Era
June 16, 2026• Qiuyang Mang 1 , Kaiyuan Liu 2 , Bo Peng 3 , Shreyas Pimpalgaonkar 4 , Alex Dimakis 1,4 , Alvin Cheung 1
1 UC Berkeley · 2 University of Washington · 3 Princeton University · 4 Bespoke Labs
2026 · test-time-scaling LLM-agents scaling-laws · research
TL;DR. Agents can spend test-time compute by trying, observing, and revising, so we ask whether their gains come from a better internal strategy or from something close to repeated sampling. We derive a simple Elo reference line: repeated sampling is linear in log test-time compute. In a 2022 two-week coding marathon, current agents plateau within 24 hours, while top humans keep improving over the official two weeks. The takeaway is that humans still do much better long-horizon test-time adaptation, and agent strategies have a lot of room to improve.
Agents Bring Intrinsic Test-Time Strategies
OpenAI’s o1 report showed that more test-time compute can improve model performance. Many papers followed, especially on verifiable tasks like code and math ( Snell et al. , Large Language Monkeys , Noam Brown’s recent post ). The common plot is success rate versus the log number of trials, or log test-time compute. These curves often rise superlinearly before they saturate.
Coverage on MATH with an oracle verifier as the number of samples increases, from Large Language Monkeys .
These studies measure model performance under an external test-time strategy. The strategy is fixed outside the model: sample many candidate solutions, check them with a verifier, and report pass@k or coverage.
However, agents change this setup. During a run, an agent can try a solution, observe the result, reflect on what failed, and revise its next attempt. This raises the question we study: when an agent improves with more test-time compute, is it using a better test-time strategy, or is it mostly reproducing repeated sampling?
Repeated sampling is fixed outside the model, while an agent can use feedback inside the run.
A Simple Model for Test-Time Scaling
We first write down the simplest model behind the usual pass@k curves: repeated sampling. The model treats each attempt as an independent draw from the same continuous score distribution. For one task, let \(X\) be the score of one sample and let \(\tau\) be the threshold for success. Then one sample succeeds with probability
With \(k\) independent samples, pass@k is
For a dataset with multiple tasks, the usual test-time scaling curve averages this quantity across tasks. This gives a curve of mean pass@k, or coverage, as a function of the number of samples.
However, this evaluation is awkward for agents. An agent usually stops once it solves the task, so it is not natural to keep asking for more independent samples after success. For open-ended tasks such as FrontierCS , we could instead compare runs by the task’s own score. But raw-score gains are hard to interpret. In circle-packing tasks studied by AlphaEvolve , improving the objective value from 1 to 2 can be trivial, while improving it from 2.35 to 2.36 can require a much harder improvement. The score number does not by itself tell us how much capability changed. We therefore want a comparison that asks a simpler question: when one candidate spends more test-time compute than another candidate on the same task, how often does it produce the better answer?
Following the same repeated-sampling model, this becomes a pairwise question. If one candidate gets \(k_a\) independent attempts and another gets \(k_b\) independent attempts, the pairwise win probability is
where \(X_i\) and \(X'_j\) are independent attempts on the same task.
Now comes the useful part. A Bradley-Terry model converts pairwise win rates into a one-dimensional strength scale. If candidate \(a\) has BT log-strength \(\theta_a\) and candidate \(b\) has BT log-strength \(\theta_b\), then
We can see that, for repeated sampling, the pairwise win probability above is exactly matched by
Thus, repeated sampling is a test-time strategy whose Elo is linear in log test-time compute. This is super helpful because it gives us a reference line. To judge an agent’s intrinsic test-time strategy, we can plot its Elo curve as test-time compute increases and compare it to this line. If the curve is above, below, or close to linear, the agent’s strategy is better than, worse than, or equivalent to repeated sampling.
Repeated sampling gives a linear reference line in Elo versus log test-time compute.
Agents Struggle, Humans Do More Than Sample
With this reference line in hand, we can ask what happens in real long-horizon tasks. We compare agent trajectories against the repeated-sampling line, and we also ask: how do they compare to top humans working on the same tasks?
We study AtCoder Heuristic Contest 014: RectJoin , a long-horizon coding and optimization contest in 2022 . Human contestants write algorithms and can submit many code solutions during the contest, and the leaderboard keeps the best score they achieve. Since the contest happened before modern coding agents were widely available, these human trajectories are not assisted by AI agents. The task is open-ended: there is no known optimal solution, only better and worse scores.
At a high level, RectJoin starts with marked dots on a grid. A solver repeatedly chooses three existing dots and one empty grid point that form a valid axis-aligned or 45-degree rectangle, then marks the new point and draws the rectangle boundary. The objective is to maximize a weighted score over the final marked dots, where dots farther from the center have larger weight.
Visualization of RectJoin from AtCoder Heuristic Contest 014 .
Human trajectories. For humans, we study two groups from the final standings: the top 10 contestants and the top 50 contestants . At each checkpoint, we compute prefix-best scores. This gives one vector for each human group at each checkpoint, where each entry is a contestant’s best score up to that time.
Agent setting. We recreate the contest setting for agents. Using the FrontierCS evaluation layer, each agent runs in a continuous 24-hour loop. It can keep submitting candidates, observe scores, and revise its next attempt, just like a human contestant. If it stops early, we resume the same session rather than starting a fresh run.
We test two agent systems: Claude Opus 4.6 with Claude Code , and GPT-5.5 with Codex . For each agent system, we run five independent 24-hour trials. At every wall-clock checkpoint we care about, we collect the prefix-best score from each trial. This gives a length-5 vector for each agent system at each checkpoint.
Once we have these vectors, we can estimate pairwise win rates directly. For example, one vector might contain the prefix-best scores of the top 10 human contestants at 24 hours, while another contains the prefix-best scores from the five Claude Code trials at 5 hours. We want to know how often the 24-hour top-human scores beat the 5-hour Claude Code scores. If the two vectors are
we define the empirical probability that \(u\) beats \(v\) by comparing all pairs:
These pairwise win rates are the inputs to the Elo fit.
We fit Elo ratings from these pairwise win rates using L-BFGS . All agent systems and human groups are placed in one joint Bradley-Terry fit, so their ratings are directly comparable. For agents, we use the 24-hour trajectories from our runs. For humans, we use the full trajectories from the official two weeks.
Elo trajectories for top humans and agent systems on RectJoin , the long-horizon coding task introduced above.
The result is sharp. Agents improve quickly in the first few hours, but their Elo curves flatten by the 24-hour mark, even though a single agent trial can use up to 100M tokens. Top humans improve more slowly at first, but they keep climbing for days and eventually pass the agent systems by a large margin. This suggests that current agents can sprint early, but they still lack the long-horizon test-time adaptation that strong human contestants use during an extended contest.
We can also ask a more local question for each participant system. If we only look at one system at a time, the repeated-sampling model gives a reference line for what would happen if that system simply drew more independent samples from its own output distribution. Comparing the observed Elo curve against this line tells us whether the system’s test-time strategy is more or less efficient than repeated sampling from itself.
Per-source Elo curves compared with the repeated-sampling reference line implied by the model.
The gray dashed line is the repeated-sampling reference. Curves above this line gain Elo faster than independent sampling from the same source distribution. Curves below it gain Elo more slowly. The agent systems become sublinear relative to this reference by the end of the 24-hour run. In contrast, the human curves become superlinear over longer horizons. Humans do not just sample; current agents still have a long way to go on long-horizon test-time scaling.
Takeaways for Agentic Test-Time Scaling
Agents have their own test-time strategies. We should evaluate whether performance improves with more compute and what strategy produced that improvement.
Use repeated sampling as a reference line. If an agent's Elo grows linearly with log compute, it may be doing little more than repeated sampling. Deviations from that line are the signal.
Keep humans as a long-horizon reference. Top humans still show adaptive improvement over long horizons, which gives us a useful target for agentic test-time scaling.
Study more open-ended long-horizon tasks. We need more tasks, longer run trajectories, and careful failure analysis to understand where agents still fall short.
Our full paper is coming soon. In the meantime, please cite this blog post if you found it helpful. For discussion, contact qmang@berkeley.edu or lky04@cs.washington.edu .

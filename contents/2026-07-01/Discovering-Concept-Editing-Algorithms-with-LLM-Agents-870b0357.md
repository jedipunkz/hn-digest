---
source: "https://dmodel.ai/concept-erasure/"
hn_url: "https://news.ycombinator.com/item?id=48746983"
title: "Discovering Concept-Editing Algorithms with LLM Agents"
article_title: "Discovering Concept-Editing Algorithms With LLM Agents"
author: "mattmarcus"
captured_at: "2026-07-01T14:59:30Z"
capture_tool: "hn-digest"
hn_id: 48746983
score: 3
comments: 0
posted_at: "2026-07-01T14:02:53Z"
tags:
  - hacker-news
  - translated
---

# Discovering Concept-Editing Algorithms with LLM Agents

- HN: [48746983](https://news.ycombinator.com/item?id=48746983)
- Source: [dmodel.ai](https://dmodel.ai/concept-erasure/)
- Score: 3
- Comments: 0
- Posted: 2026-07-01T14:02:53Z

## Translation

タイトル: LLM エージェントを使用したコンセプト編集アルゴリズムの発見
記事のタイトル: LLM エージェントを使用したコンセプト編集アルゴリズムの発見

記事本文:
LLM を使用したコンセプト編集アルゴリズムの発見
エージェント
アダム・シャーリス * 、アレックス・ビシュカ * 、
アディティア・ヴァサンタラオ、アジット・シヴァクマール、アレックス・サンチェス＝スターン、アレクサナ・デュボア、
アニッシュ・トンドワルカー、アルジュン・パギディ、オータム・シンクレア、カリー・ウィンター、ダルトン
コムズ、ダニエル・ムーン、デヴィッド・カオ、フェリックス・ミシャラク、ジェフ・コッグシャル、クナル
バム、モヒット・テクリワル、リテーシュ・バレラオ、シャウリヤ・ジェイン、タイラ・バージェス、ユエル
ヤン、ジャンナ・カウフマン
*コア
貢献者
概念消去とは、概念から不要な情報を削除する技術です。
モデルのアクティベーションは可能ですが、現在の消去方法では完全に消去するのは困難です
ターゲットの概念を削除します。この研究では、次の訓練を受けた LLM エージェントにタスクを課しました。
優れたパフォーマンスを発揮する概念消去アルゴリズムを発明してデータを収集
同じ実験的制約の下での現在の方法。測定します
各アルゴリズムファミリーのパフォーマンスを調査し、その原因を探る
現在の方法では不十分です。
発見された 4 つのアルゴリズム ファミリ
4.1 最適ガウス
輸送
4.2 MMD に最適化されたアフィン
トランスフォーム
4.3 ニューラルネットワーク + MMD
トレーニング
4.4 ソフト予測 + モーメント
マッチング
4.5 ポイントレベルの最適な転送
(ハンガリー語マッチング)
4.6 分位数/次元ごと
輸送
5 重要な発見: 共分散は
最強の信号
AI の安全性と
解釈可能性
安全なモデルを作成するには、モデルが知っていることと、
使用します。コンセプト消去は、モデルを変更するためのツールの 1 つです。
表現。再トレーニングや微調整ではなく、概念の消去
モデルのアクティベーションに到達し、ターゲットの概念を削除します。
使用できなくなる。
概念を消去する機能は、AI の安全性に関するさまざまな取り組みをサポートします。
概念の消去は、モデルが構築に必要な知識を忘れるのに役立ちます
生物兵器、差別を可能にするデリケートな属性を削除する、そして
ミスアライメントを引き起こす概念を除去します。しかし、概念が存在することはほとんどありません。
イソル

古い、取り外し可能な機能により、完全な消去が困難になります
達成する。
一般的な手法の 1 つである LEAst-squares コンセプト
Erasure ( LEACE ) (Belrose et al. 2023) 、線形領域を削除します。
閉じた形式の概念のコンポーネント。でラベル付けされた特定のアクティベーション
LEACE は、二項対立の概念 (例: 「皮肉」と「皮肉ではない」) を削除します。
ターゲットコンセプトの線形コンポーネントを最小限の編集で再現します。
その後、線形分類器はターゲットを回復できなくなります。
偶然を超えた概念、 Linear Guardedness として知られる特性。
LEACE の後継、Quadratic LEAst-squares Concept
Erasure ( QLEACE ) (Quirke and Belrose 2025)、追加
2 つのクラスの共分散を等化して、二次分類器が存在しないようにします。
目的の概念を回復することができます。
ただし、非線形表現を変更する線形手法として、LEACE
また、QLEACE はモデル内のターゲット概念を完全に消去できません。レース
2 つのクラスの平均には一致しますが、共分散には一致しません。クリアス
平均と共分散は一致しますが、高次は保持されます。
構造。 LEACE で消去されたアクティベーションでトレーニングされた RBF カーネル SVM はまだ
70% ～ 95% の精度でターゲットの概念を復元しました。非線形消去
一般化するのに苦労: カーネル化された概念
Erasure (Ravfogel et al. 2022) は次のことを示しました。
1 つの非線形分類器から保護すると、概念から完全に離れることができます
他の人が読むことができます。
私たちは、データに基づいてトレーニングされたエージェントに、次のようなアルゴリズムを発明するよう命じました。
LEACE を上回るニューラル ネットワークの活性化から概念を消去し、
クリアス。タスクの説明と L2 編集距離のみが与えられる
予算に応じて、エージェントは 6 つの異なる消去ファミリーを発見しました
50 の概念にわたるアルゴリズム。最良のアルゴリズムは平均を削減しました
非線形プローブの精度は 99% から 70% ですが、LEACE のみ
精度が 88% に低下しました。さらに、開催中の

ランダムフォレスト
分類子、最良のアルゴリズムでは精度が 99% から 72% に低下しましたが、
LEACE は精度を 82% に低下させるだけでした。
各エージェントには、ターゲットの概念、ラベル付きの小さなサンプル、
Gemma-3 270M モデルについて、次のような質問がありました: 1
モデルのアクティベーションを編集して、
非線形分類器ではそれを回復できないような概念です。あなたの
変更は LEACE の L2 バジェット内に収まる必要があります。
エージェントはアルゴリズムや公開されたソリューションを受け取りませんでした。彼らは
採点者は、持ち出された評価に対して自分たちの方法を評価することができませんでした。
新しい非線形プローブを使用したアクティベーション。エージェントは次のことを行う必要がありました。
活性化ジオメトリを分析して、LEACE がなぜ問題を残すのかを理解する
非線形信号。
それを取り除くアルゴリズムを考案します。
ハイパーパラメータを実装、デバッグ、最適化します。
指定された L2 バジェット内で編集を維持します (同じ量の L2 バジェット)
LEACE が使用する変更)。
私たちは 50 のコンセプトにわたって 560 の独立したロールアウトを実行しました。
「皮肉」と「ユーモア」から「メタ認知」と「仮説性」。私たちが使用した
LEACE ベースラインと比較してパフォーマンスを評価する 2 つのアルゴリズム:
SVM の精度とランダム フォレストの精度。
SVM の精度が主な評価指標でした。アン
RBF カーネル SVM は、エージェントが一度も見たことのないデータに基づいてトレーニングされ、
別のテストセット。 SVM 精度スコア 50% はランダムと同等です
チャンス（つまり、完全な消去）。 50% を超えるスコアは目標を意味します
コンセプトは回収可能でした。
あらゆるコンセプトにおいて、SVM 上の LEACE に勝る最適なエージェント ソリューション
精度。平均して、LEACE は SVM 精度を 88% に低下させますが、
エージェントはそれを 70.1% に削減しました。付録に示されているのは、
50 個のコンセプトすべてに対するコンセプトごとの SVM 結果。
一般化のテストとして、ランダムで評価も行いました。
フォレスト分類子 (木 100 本、最大深さ 10)。これは違います
非線形c

エージェントが最適化していなかった発病者ファミリー。
ランダム フォレストでもターゲットのコンセプトを回復できなかった場合、
エージェントが特定の範囲を超えて転送する消去を発見したことを示しました
SVM に基づいて採点されました。この結果は本物であることを示唆しています
1 つのカーネルにオーバーフィットするのではなく、分布を一致させます。
Random Forest は、SVM ではできなかったいくつかの概念を回復しましたが、私たちの最善の概念は
エージェント ソリューションは一般的に LEACE に勝ります。平均して、LEACE はランダムを削減します
フォレストの精度は 82.6% でしたが、エージェントはそれを 71.9% に下げました。を参照してください。
コンセプトごとのランダムのマッチングに関する付録
森の崩壊。
4 発見されたアルゴリズム
家族
独立して作業しているにもかかわらず、当社のエージェントは 6 つのアルゴリズムに集中しました
家族。私たちは 50 の最良のソリューション (コンセプトごとに 1 つ) をその基準に基づいてクラスタリングしました。
核となるアプローチ。
4.1 ガウス最適輸送
6 つのバリアント · 平均 SVM 71% · 最高: 確実性 (59%)
ガウスを各クラスに適合させてから、両方のクラスを共有クラスに移動します。
(LEACE のように) 平均値と完全値を均等化する「平均」形状
共分散。このアクションにより、RBF によって読み取られた共分散の差が削除されます。
閉じた形式の線形マップを使用する - まさに QLEACE の構築です。
4.2 MMD に最適化されたアフィン
トランスフォーム
12 のバリエーション · 平均 SVM 63% · 最高: 共感 (49%)
閉じた形式のソリューションの代わりに、グラデーションによるアフィン編集を学びます
最大平均不一致 (MMD) を直接最小化する降下: a
2 つのクラスがどのように異なって見えるかの尺度。
SVM 独自のカーネル。 SVM が認識するものを正確にターゲットにします。しかし、消去する
あるプローブを別のプローブで消去する必要はありません。これが、次のプローブに転送する理由の 1 つです。
ランダムフォレストは不完全です。
4.3 ニューラルネットワーク + MMD
トレーニング
6 つのバリエーション · 平均 SVM 66% · 最高: ユーモア (57%)
小さなニューラル ネットワークが各ポイントに対する独自の調整を学習し、
2 つのクラスが分布するまでトレーニングされた

と区別がつかない
カーネル (MMD、またはそれに関連する HSIC)。マップは非線形であるため、
共分散だけではなく完全な分布と一致します。スケーリング
ポイント独自のノルムへの各調整により、すべての編集が L2 内に保持されます
予算。
4.4 ソフト予測 + モーメント
マッチング
22 のバリアント · 平均 SVM 74% · 最良: 一致 (52%)
ソフト分類子をトレーニングして、各ポイントがどの程度強く属するかをスコア化します。
1 つのクラスを作成し、そのスコアに比例してクラスを変更し、クラスを引っ張ります
一緒になって次元ごとの広がりを均等化することを意味します。これにより削除されます
RBF が検出できる信号: 2 つの間の分散の差
クラス。自信を持ってブレンドすると、編集がスムーズに機能します。
決定境界ではハードスイッチではなく入力を使用します。
4.5 ポイントレベルの最適な転送
(ハンガリー語マッチング)
3 つのバリエーション · 平均 SVM 72% · 最高: 好奇心 (60%)
各ポイントを反対クラスのポイントとペアリングします (ハンガリー語経由)
マッチング アルゴリズム) を使用して 2 つを一緒に移動し、2 つの雲を折りたたむ
一つに。新しい点は最も近い点に続きます。何も仮定しない
データの形状について - 雲を直接重ねることで、
最初の 2 つの瞬間だけでなく、分布全体と一致します。
4.6 分位数/次元ごと
輸送
1 つのバリアント · 平均 SVM 69% · 最高: 権威 (69%)
一度に 1 つの活性化ディメンションを作業します。各クラスの形状を再構築します。
分位点を一致させることにより、その軸に沿って共有ターゲットに分布します。
(1D 最適トランスポート)。これは、次元ごとのあらゆる瞬間に一致しますが、
分散のみですが、次元間の相関は無視されます。これ
クラスが主に次元が異なる場合に、安価で効果的になります。
次元。
5 重要な発見: 共分散は
最強の信号
成功したエージェントのほぼすべてが同じ認識に達しました: LEACE
2つのclaの差を取り除きます

sses の意味ですが、コンセプトは
各クラスの分布の全体形状、最も顕著なもの
差は共分散です。 RBF カーネル SVM はペアごとに動作します
ポイント間の距離があるため、各クラスの広がりに敏感です
です。一方のクラスが他方のクラスよりも緊密にクラスタ化されている場合、SVM は次のように読み取ります。
コンセプトはそのままスプレッドの違い。この共分散
違いは QLEACE によって除去されるものですが、カーネル分類子は依然として除去できます。
残った構造を利用します。
LEACE の適用後も、RBF-SVM はターゲットを回復できます。
コンセプトは約87％。残された共分散の差は小さいですが、
それはまさにRBFが読んでいることです。小さいということは、削除が安価であることを意味します: 以内
LEACE 独自の L2 バジェット、エージェントの最高の消しゴムがその精度を向上させます
LEACE が残した場所の下にあります。より優れた消去には多額の予算はかかりません
LEACEよりも。
各アルゴリズムファミリーは共分散の差を攻撃しましたが、
それらを区別したのは、それらが分布のどの程度一致するかでした。もし
完全な分布が一致していることは、ターゲットのコンセプトが以下であることを示唆しています。
アクティベーションの形状全体にエンコードされます。消去レベル
概念を削除するために必要なのは、その概念が表現されている証拠である
幾何学。このように、家族は概念の分類法としても機能します。
モデル内でエンコードできます。
AI の安全性と
解釈可能性
実際には、線形保護を保証しても完全には達成されません。
概念の消去。 LEACE と QLEACE は線形プローブと二次プローブを消去できます。
それぞれ、しかしその有効性は非線形の性質によって制限されます
モデルの。ターゲットの概念は削除されたように見えるかもしれませんが、その機能は
単純な非線形検出器によって回復可能なままです。
現在のコンセプト消去方法では不十分ですが、当社のエージェントは次のような提案をしました。
識別による進むべき道

攻撃する新しいアルゴリズム
高次構造の違い。基準溶液なし、または
関連する論文へのアクセス、エージェントが活性化ジオメトリを分析した、
LEACE と QLEACE が失敗する理由について仮説を立て、特定した
より高性能な消去方法。
この研究は、概念を消去するには、次のことを理解する必要があることを示しています。
モデルがそれをどのようにエンコードするか。これにより、概念消去の研究が価値のあるものになります
安全性のためだけではなく、解釈可能性のためにもこの研究を紹介します。
エージェントに求めている自由回答型調査の例として、
加速します。
Claude (Anthropic) がこの投稿の準備を支援しました — 草稿を手伝ってくれました
最終バージョンの文章を作成し、図を作成します。の
研究、実験、発見は著者自身のものです。
LEACE と 50 のコンセプトすべてに対する最高のエージェントのパフォーマンス。
RBF-SVM。
LEACE と 50 のコンセプトすべてに対する最高のエージェントのパフォーマンス。
ランダムフォレスト分類器。
以下のテキストは一般的な内容を伝える言い換えです。
エージェントに与えられた文字通りのプロンプトではなく、タスクに対する感情です。 ↩︎

## Original Extract

Discovering Concept-Editing Algorithms With LLM
Agents
Adam Scherlis * , Alex Bishka * ,
Aditya Vasantharao, Ajit Sivakumar, Alex Sanchez-Stern, Alexana Dubois,
Anish Tondwalkar, Arjun Pagidi, Autumn Sinclair, Curry Winter, Dalton
Combs, Daniel Moon, David Cao, Felix Michalak, Jeff Coggshall, Kunal
Bham, Mohit Tekriwal, Ritesh Bhalerao, Shaurya Jain, Tyra Burgess, Yueru
Yan, Zhanna Kaufman
* Core
contributors
Concept erasure is a technique that removes unwanted information from
a model’s activations, but current erasure methods struggle to fully
remove target concepts. In this study, we tasked LLM agents trained on
our data with inventing concept erasure algorithms that outperform
current methods under the same experimental constraints. We measure the
performance of each algorithm family and explore the cause of why
current methods fall short.
4 Discovered Algorithm Families
4.1 Gaussian Optimal
Transport
4.2 MMD-Optimized Affine
Transform
4.3 Neural Network + MMD
Training
4.4 Soft Prediction + Moment
Matching
4.5 Point-Level Optimal Transport
(Hungarian Matching)
4.6 Quantile / Per-Dimension
Transport
5 Key Finding: Covariance Is the
Strongest Signal
6 Implications for AI Safety &
Interpretability
To create safe models, we must be able to control what they know and
use. Concept erasure is one tool for modifying a model’s
representations. Rather than retraining or fine-tuning, concept erasure
reaches into a model’s activations and removes a target concept,
rendering it unusable.
The ability to erase concepts supports a range of AI-safety efforts.
Concept erasure can help models unlearn knowledge needed to build
bioweapons, remove sensitive attributes that enable discrimination, and
ablate concepts that drive misalignment. But concepts are rarely
isolated, detachable features, making perfect erasure difficult to
achieve.
One prevalent method, LEAst-squares Concept
Erasure ( LEACE ) (Belrose et al. 2023) , removes the linear
component of concepts in closed form. Given activations labeled by a
binary concept (e.g., “sarcastic” vs. “not sarcastic”) LEACE removes the
target concept’s linear component with the smallest possible edit.
Afterwards, a linear classifier should not be able to recover the target
concept above chance, a property known as Linear Guardedness .
LEACE’s successor, Quadratic LEAst-squares Concept
Erasure ( QLEACE ) (Quirke and Belrose 2025) , additionally
equalizes the two classes’ covariances, so that no quadratic classifier
can recover the target concept.
However, as linear methods modifying nonlinear representations, LEACE
and QLEACE cannot fully erase target concepts in the model. LEACE
matches the two classes’ means, but not their covariances. QLEACE
matches the means and covariances but preserves the higher-order
structure. An RBF-kernel SVM trained on LEACE-erased activations still
recovered target concepts with 70%–95% accuracy. Nonlinear erasure
struggles to generalize: Kernelized Concept
Erasure (Ravfogel et al. 2022) showed that
protecting against one nonlinear classifier can leave the concept fully
readable by another.
We tasked agents trained on our data with inventing algorithms that
erase concepts from neural network activations that outperform LEACE and
QLEACE. Given only a description of the task and an L2 edit-distance
budget, our agents discovered six distinct families of erasure
algorithms across 50 concepts. The best algorithm reduced the average
accuracy of a nonlinear probe from 99% to 70%, whereas LEACE only
reduced accuracy to 88%. Additionally, on a held out random forest
classifier, the best algorithm reduced accuracy from 99% to 72%, whereas
LEACE only reduced accuracy to 82%.
Each agent was given a target concept, a small labeled sample, a
Gemma-3 270M model, and was asked the following: 1
Create a procedure that edits the model’s activations to remove a
concept such that nonlinear classifiers cannot recover it. Your
modification must stay within the L2 budget of LEACE.
The agents did not receive an algorithm or a published solution. They
were unable to view the grader, which evaluated their method on held-out
activations, with a fresh nonlinear probe. Agents had to:
Analyze the activation geometry to understand why LEACE leaves a
nonlinear signal.
Devise an algorithm to remove it.
Implement, debug, and optimize hyperparameters.
Keep the edit within the specified L2 budget (same amount of
modification LEACE uses).
We ran 560 independent rollouts across 50 concepts, ranging from
“sarcasm” and “humor” to “metacognition” and “hypotheticality.” We used
two algorithms to evaluate performance compared to the LEACE baseline:
SVM Accuracy and Random Forest Accuracy .
SVM Accuracy was our primary evaluation metric. An
RBF-kernel SVM is trained on data the agent never saw and evaluated on a
separate test set. An SVM Accuracy score of 50% is equal to random
chance (i.e., perfect erasure). A score higher than 50% meant the target
concept was recoverable.
For every concept, the best agent solution beat LEACE on SVM
accuracy. On average, LEACE reduces the SVM Accuracy to 88%, while our
agents reduced it to 70.1%. The Appendix shows
the per-concept SVM results for all 50 concepts.
As a test of generalization, we also evaluated with a Random
Forest Classifier (100 trees, max depth 10). This is a different
nonlinear classifier family that the agents were not optimizing against.
If a Random Forest also failed to recover the target concept, that
indicated the agents found erasure that transfers beyond the specific
SVM they were graded on. This result would suggest genuine
distribution-matching rather than overfitting to one kernel.
Random Forest recovered some concepts the SVM could not, but our best
agent solutions generally beat LEACE. On average, LEACE reduces Random
Forest accuracy to 82.6%, while our agents reduced it to 71.9%. See the
Appendix for the matching per-concept Random
Forest breakdown.
4 Discovered Algorithm
Families
Despite working independently, our agents converged on six algorithm
families. We clustered the 50 best solutions (one per concept) by their
core approach.
4.1 Gaussian Optimal Transport
6 variants · avg SVM 71% · best: certainty (59%)
Fit a Gaussian to each class, then move both classes to a shared
“average” shape, equalizing the means (as LEACE does) and the full
covariances. This action removes the covariance difference read by RBF
with a closed-form linear map - exactly QLEACE’s construction.
4.2 MMD-Optimized Affine
Transform
12 variants · avg SVM 63% · best: empathy (49%)
Instead of a closed-form solution, learn an affine edit by gradient
descent that directly minimizes the Maximum Mean Discrepancy (MMD): a
measure of how different the two classes still look, evaluated with the
SVM’s own kernel. It targets exactly what that SVM sees; but erasing for
one probe needn’t erase for another, which is part of why transfer to
the Random Forest is imperfect.
4.3 Neural Network + MMD
Training
6 variants · avg SVM 66% · best: humor (57%)
A small neural network learns its own adjustment for each point,
trained until the two class distributions are indistinguishable to a
kernel (MMD, or its relative HSIC). Because the map is nonlinear, it can
match the full distribution rather than just the covariance. Scaling
each adjustment to the point’s own norm, keeps every edit inside the L2
budget.
4.4 Soft Prediction + Moment
Matching
22 variants · avg SVM 74% · best: agreement (52%)
Train a soft classifier to score how strongly each point belongs to
one class, then shift it in proportion to that score, pulling the class
means together and equalizing their per-dimension spread. This removes
the signal an RBF can detect: the difference in variance between the two
classes. Blending by confidence makes the edit a smooth function of the
input rather than a hard switch at the decision boundary.
4.5 Point-Level Optimal Transport
(Hungarian Matching)
3 variants · avg SVM 72% · best: curiosity (60%)
Pair each point with an opposite-class point (via the Hungarian
matching algorithm) and move the two together, collapsing the two clouds
into one; new points follow their nearest neighbor. It assumes nothing
about the shape of the data - by overlapping the clouds directly, it
matches the entire distribution, not just the first two moments.
4.6 Quantile / Per-Dimension
Transport
1 variant · avg SVM 69% · best: authority (69%)
Work one activation dimension at a time: reshape each class’s
distribution along that axis onto a shared target by matching quantiles
(1D optimal transport). This matches every per-dimension moment, not
just the variance, but it ignores cross-dimension correlations. This
makes it cheap and effective when the classes differ mainly dimension by
dimension.
5 Key Finding: Covariance Is the
Strongest Signal
Almost all successful agents arrived at the same realization: LEACE
removes the difference in two classes’ means, but a concept lives in the
whole shape of each class’s distribution, with the most salient
difference being the covariance. An RBF-kernel SVM works from pairwise
distances between points, so it’s sensitive to how spread out each class
is. If one class is more tightly clustered than the other, the SVM reads
the concept straight off the difference in spread. This covariance
difference is what QLEACE eliminates, but a kernel classifier can still
exploit the remaining structure.
After LEACE is applied, an RBF-SVM can still recover the targeted
concept at about 87%. The leftover covariance difference is small, but
it’s exactly what the RBF reads. Small means cheap to remove: within
LEACE’s own L2 budget, the agents’ best erasers drive that accuracy well
below where LEACE leaves it. Better erasure doesn’t take a bigger budget
than LEACE’s.
Each algorithm family attacked the covariance difference, but what
distinguished them was how much of the distribution they match. If the
full distribution is matched, that suggests the target concept is
encoded in the entire shape of the activations. The level of erasure
needed to remove a concept is evidence of the concept’s representation
geometry. In this way, families double as a taxonomy of the way concepts
can be encoded in the model.
6 Implications for AI Safety &
Interpretability
In practice, guaranteeing linear guardedness does not achieve total
concept erasure. LEACE and QLEACE can erase linear and quadratic probes,
respectively, but their effectiveness is limited by the nonlinear nature
of models. Target concepts may appear removed, but their capabilities
remain recoverable by simple nonlinear detectors.
While current concept erasure methods fall short, our agents offered
a path forward by identifying novel algorithms that attacked
higher-order structural differences. Without a reference solution or
access to relevant papers, agents analyzed activation geometry,
hypothesized about why LEACE and QLEACE fail, and identified
higher-performance erasure methods.
This study demonstrates that to erase a concept, we must understand
how the model encodes it. This makes concept erasure research valuable
not just for safety, but for interpretability, and we present this study
as an example of the open-ended research we want agents to
accelerate.
Claude (Anthropic) assisted in preparing this post — helping to draft
the writing and to create the figures for the final version. The
research, experiments, and findings are the authors’ own.
LEACE and best-agent performance for all 50 concepts, under the
RBF-SVM.
LEACE and best-agent performance for all 50 concepts, under the
Random Forest classifier.
The text below is a paraphrase conveying the general
sentiment of the task, not the literal prompt given to the agent. ↩︎

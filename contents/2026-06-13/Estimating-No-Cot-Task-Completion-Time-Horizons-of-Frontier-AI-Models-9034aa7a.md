---
source: "https://www.lesswrong.com/posts/SieLowPgNgRSPGhFw/estimating-no-cot-task-completion-time-horizons-of-frontier"
hn_url: "https://news.ycombinator.com/item?id=48514091"
title: "Estimating No-Cot Task-Completion Time Horizons of Frontier AI Models"
article_title: "Estimating No-CoT Task-Completion Time Horizons of Frontier AI Models — LessWrong"
author: "kqr"
captured_at: "2026-06-13T07:24:49Z"
capture_tool: "hn-digest"
hn_id: 48514091
score: 1
comments: 0
posted_at: "2026-06-13T06:37:58Z"
tags:
  - hacker-news
  - translated
---

# Estimating No-Cot Task-Completion Time Horizons of Frontier AI Models

- HN: [48514091](https://news.ycombinator.com/item?id=48514091)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/SieLowPgNgRSPGhFw/estimating-no-cot-task-completion-time-horizons-of-frontier)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T06:37:58Z

## Translation

タイトル: フロンティア AI モデルのノーコットタスク完了時間の範囲の推定
記事のタイトル: フロンティア AI モデルの No-CoT タスク完了時間の範囲を推定する — LessWrong
説明: (最後にある完全な著者リストを参照) • 約 1 年前、METR は、フロンティア モデルのタスクの長さが数か月ごとに 2 倍を確実に完了できることを示しました。

記事本文:
ログイン CoT なしのタスク完了時間の推定 フロンティア AI モデルの視野 — LessWrong AI Control AI 評価 Redwood Research Scaling Laws AI フロントページ 2026 トップ 50: 13 % 167
フロンティア AI モデルの No-CoT タスク完了時間の範囲を推定する
(最後にある完全な著者リストを参照してください)
約 1 年前、METR は、フロンティア モデルが確実に完了できるタスクの長さが数か月ごとに 2 倍になることを示しました。これに関連する安全関連の質問は、モデルが思考連鎖 (CoT) なしで完了できるタスクの長さはどれくらいですか?私たちは新しい論文で調査します。
モデルが CoT を出力せずに広範な推論を実行できる場合、安全性に影響を与える可能性があります。開発者やデプロイメント時の監視者は、モデルの動機を簡単に理解して、危険な計画を発見することができませんでした。実質的に CoT なしで推論するモデルは、その推論が事前の事前トレーニングのテキストによって制約されなくなるため、人間の思考パターンからさらに逸脱する可能性もあります。その結果、理解が難しくなり、陰謀を企てやすくなる可能性があります。
Ryan Greenblatt の研究を拡張して、さまざまなドメインにわたる 43 のベンチマーク スイートで CoT なしでタスクを完了するモデルの能力を測定することで、これを調査します。推定 50% のタイムホライズン (TH)、つまり LLM が 50% の成功率で実行するタスクを人間が実行するのにかかる一般的な時間を使用して、AI の推論能力を人間と比較します。 GPT-5.5 のようなフロンティア モデルは、人間が約 3 分かかる質問に 50% の信頼性で回答することがわかり、この TH は 2019 年以来約 2 倍になっています。
図 1: 当社の CoT なし TH (緑) と METR の CoT あり TH (紫) の比較。 GPT-4 がリリースされるまで、CoT の有無にかかわらず TH は同様の速度で増加しました。 CoT TH を備えた GPT-4 以降

CoT のない TH のおよそ 2 倍の速度で成長しています。
AI企業がCoTモニターや人間に公開することなくモデルがどれだけの推論を実行できるかの下限を見つけるために、no-CoT THを明示的に追跡し始めることを提案します。私たちのテスト スイートは大幅な推論計算を必要としないため、これらの評価は低コストで実行できます。
GPT-2 (2019) から GPT-5.5 (2026) までの 14 のフロンティア モデルを、数学、コーディング、知識、エージェント ツールの使用、ステガノグラフィーやスキーム推論などの安全関連の質問をカバーする 43 のベンチマークで評価しました。質問には、単一トークンの回答、より長い生成、および複数ターンのエージェント設定が含まれます。
いずれの場合も、ベンチマーク固有のプロンプトと構造化された出力制約を使用することで、モデルが回答の前に推論トークンを発行しないようにします。
質問の難易度は次の 2 つの方法で推定します。
人間の解決時間 : 人間が電卓、紙、鉛筆を使用して問題を解くのにかかる平均時間 (LLM アクセスなし)。一部のベンチマークについては、人間による解決時間のデータがあります。他の人に対しては、小規模な人体試験を実施しました。残りの部分については、コンテキスト内の人間時間の例を通じて情報を得た、Claude Opus 4.7 の推定値を使用しました。
推論トークンアンカー: o3-mini が質問を解決するために必要なトークンの最小数。 o3-mini は各質問を 8 回試行し、成功した試行の中から最小限の推論トークンを取得します。与えられた質問を解決しようとしても解決しない場合、その質問は削除されます。 CoT を使用せずに評価されたモデルが、o3-mini が推論で解決できない問題を解決できることはほとんどありません。
これら 2 つの方法により、問題を解決するために「どの程度の推論」が必要かという独立した尺度が得られます。
各モデルについて、Kwa et al と同じ方法を使用して、これらのメトリクスの両方でその 50% の解決範囲を推定します。 l と比較した成功​​率にロジスティック曲線を当てはめます。

人間の解決時間と推論トークン アンカーのログを計算し、ロジスティックが 50% に等しくなる点を見つけます。
図 2: 各モデルについて、ロジスティック曲線を当てはめて、タスクを完了するのに必要な人間の時間 (左) または o3-mini 推論トークン (右) の関数としてタスクを正常に完了する確率を予測します。
不確実性を推定するために、計算を 10,000 回再実行し、ベンチマーク、各ベンチマーク内の質問、および質問ごとにモデルを実行して置換して再サンプリングしました。また、すべての質問に対してノイズ モデルを調整することで、人間の解答時間に不確実性を組み込みます。対照的に、推論トークン アンカーの不確実性はモデル化しません。
図 3: No-CoT TH の経時変化。各ポイントはリリース日別にプロットされたモデルです。線は指数関数的近似であり、網掛けは 95% CI を示しています。
図 4: 時間の経過に伴うトークン ホライズン。各ポイントはリリース日別にプロットされたモデルです。線は指数関数的近似であり、網掛けは 95% CI を示しています。
TH は 373 日ごとに 2 倍になります (95% CI: 167 ～ 691)。
トークンホライズンは 437 日ごとに 2 倍になります (95% CI: 341–571)。
GPT-5.5 は現在のフロンティアにあり、50% no-CoT TH は約 3 分、推論トークン範囲は約 1,500 o3-mini トークンです。
フォワード中央値を外挿すると、2028 年までに約 7 分 / 約 3.7,000 トークン、2030 年までに約 25 分/約 12,000 トークンが得られます。倍加時間 CI の範囲は広いですが、95% CI の終わりが遅い場合でも、2030 年までに潜在推論は約 10 分になると考えられます。
2 倍化は、ベンチマークの分布の変更に対してほぼ堅牢です。つまり、より長い世代およびマルチターンのエージェント質問を追加し、測定されたヒューマンタイムを持つ質問のみに制限し、単一のモデルを削除し、質問ドメインごとに分類します (ドメイン固有の 2 倍化の範囲は 246 日から 387 日です)。
私たちの傾向が Kwa らの傾向と異なるのはなぜですか?
クワ

他。は、-CoT を使用したエージェント ソフトウェア タスクの TH の倍増時間が 182 日であると報告しました。これは、私たちのものより約 2 倍の速さです。 GPT-4 以降の能力の向上のほとんどは、1 回のフォワード パスでのさらなる推論ではなく、外部化された推論によってもたらされているため、GPT-4 以降では 2 倍化時間は異なります。
私たちは推論ではなく暗記を追跡しているのでしょうか?
私たちのベンチマークにはいくつかの事実を思い出す質問が含まれているため、純粋に推論を測定しているわけではありません。推論重視の質問が多く含まれているため、見出し番号は推論をうまく追跡していると考えられますが、不完全です。
モデルは実際に推論を隠しているのでしょうか?
現在のモデルが実際に CoT からの推論を隠しているとは主張しません。私たちは、彼らが実際にどれだけのことをするかではなく、目に見えない形でどれだけのことをもっともらしく実行できるかを測定します。一部のモデルは、たとえそれが不必要であっても、推論を言語化します。
モデルはさらに隠された推論を行うことができるでしょうか?
不透明な推論の下限を測定します。
短い回答を強制することは分布から外れ、推論モデルにとって奇妙になる可能性があり、推定値を押し下げる可能性があります。より良い引き出し、または質問の異なる分布により、モデルのパフォーマンスが向上する可能性があります。
外部化された推論トークンなしでモデルが何ができるかを測定します。しかし、モデルは推論を一見無害な出力にエンコードすることを学習し、CoT モニター (「ステガノグラフィー」) を回避しながら、この下限以上のことを行うことができます。
フロンティア モデルは、CoT なしでかなりの量の推論を実行でき、CoT なしで実行できる推論の量は指数関数的に増加しています。
指示に従うかどうか、シナリオが評価に見えるかどうか、いくつかの目に見える選択肢のうちどれを採用するかなど、一部の単発的な決定には、人間と同等の推論を 3 分間行うだけで十分かもしれません。しかし、複雑な長期的な電力供給を行うにはおそらく十分ではありません。

トレーニングゲームを決定するか、他のインスタンスと共謀するかなど、王の推論。この傾向が続けば、モデルは 2030 年までに 25 分間の CoT なしの推論が可能になる可能性があり、これによりさらに多くの破壊活動が可能になるでしょう。ニューラライズがなくても、CoT モニタリングの効果は今後数年で大幅に低下する可能性があります。
論文著者: フランシス・リース・ウォード、デウィ・グールド、アンダース・ケアンズ・ウッドラフ、ラウノ・アリケ、ジョシュ・ヒルズ、アレックス・セラーノ、アイダ・キャスパリー、ジェイソン・ロス・ブラウン、ジョー・J・ジャオ、パトリック・リースク、トゥム・ストーン、ラム・ポサム、アイオナット・ガブリエル・スタン、ハリー・メイン、シメオン・ヘルステン、シュボラップ・ビスワス、アリアナ・アザルバル、ウィリアム・L・アンダーソン、エルナジット、ライアン・グリーンブラット、ジュリアン・スタストニー
トレーニングの汚染を避けるために、データセットは一般公開されません。リクエストに応じて利用可能です。これは、レッドウッド研究所、アストラフェロープログラム、エーテル研究所、MATS の共同研究でした。論文の全著者リスト。
3 コメント 3 AI 制御 AI 評価 Redwood Research スケーリング法 AI フロントページ 167
新しいコメント 3 件のコメントを送信し、スコアの高い順に並べ替えます クリックして新しいコメントをハイライト表示します: Today at 7:24 AM [ - ] Petropolitan 2d 4 2 大幅な進歩は推論のスケーリングによるものであるという従来の通念を定量化した非常に興味深い論文。
OpenAI から GPT-3 および 3.5 にアクセスできましたが、2025 年前半をより適切にカバーするために、現在は公開されていないクロード モデルについても同じものを入手しようとしましたか?
[ - ] dgros 6h * 1 0 興味深いですね。気になる点としては、(1) どのタスクが結果に最も影響するか、(2) どれくらいの時間の数字に意味があるのか、(3) 一般的な引き出しの質問、などです。
見出しに最も影響を与えているタスクとタスクの種類はどれですか?
METR のタイムライン プロットについては、特定の項目が結果に大きな影響を与えるかどうかについて、ここ LW やその他の場所で議論されています。

s.私は、どの種類のタスクが最近の 2 倍化をすでに説明するのに最も適しているのか、そしてこの方法論に従って再び 2 倍になった世界はどのように見えるのか疑問に思っていました。
私はこの分析の非常に大まかなバージョンを雰囲気付けしようとしました。基礎となるデータがなければ難しいですが、クロードに表 7 ～ 13 から精度の数値を抽出させ、図 47 をピクセルで覗いて、人間時間データの中央値と IQR を取得することができます。 [1]
約 1 ～ 3 分の帯域のものは 50% クロスオーバー ポイント付近にある可能性が最も高く、回帰により大きな影響を与えると思います。 Y 軸の「バイブコーディング妨害」や X 軸の SHADE モニターなど、いくつかの外れ値は考慮に値するかもしれません。これらは両方とも、他のデータセットとは異なる、一種の TPR @ ~1-2% FPR メトリクスを使用します。ここでは効果は不明。
ただし、これは危険であり、より信頼性の高い分析を行うために基礎となるデータにアクセスする必要があります。
各データセットの影響を定量化する他の方法もあります (ここでも、「バイブ コーディング サバオタージュ」タスクが、他のデータセットの 2 倍以上の影響力を持つ明らかな外れ値として表示されます)。ただし、方法論にはあまり自信がないので、折りたたみ可能なものに隠しています。
Vibey 影響スコアを計算できますか?
クロードの説明:
バーは、各ベンチマークが GPT-4o と GPT-5.5 の間のヘッドラインの時間軸の向上にどの程度寄与しているかを推定します。見出しのメトリクスは、成功に対するロジスティック曲線近似と人間の対数解決時間の 50% のクロスオーバーであり、ロジスティック近似の基本特性は、近似曲線に対する観測値の影響が p(1−p) でスケールされることです。これは、モデルが約半分の時間を通過するタスクでは最大で、ほぼ常に成功するか、ほぼ常に失敗するタスクではゼロに近くなります。これら 2 つのモデル間ではクロスオーバーが約 0.9 分から 3 分に変化したため、精度が向上しました。

その時間範囲内のタスクに関する s は、クロスオーバーがどこにあるかについてのほとんどの証拠を提供します。したがって、各バーは、ベンチマークの精度変化 (論文の正規化にほぼ従った確率補正 - 正確に MCQ タスクの場合、検出タスクの ~1.5% ベースラインは 0、~1pp の差として扱われる) に、~1.6 分を中心とする p(1-p) 重みを掛けたものであり、ログ時間幅は約 10 年であり、論文の図 47 から抽出されたベンチマークの質問ごとの人間時間分布で平均化されています。ひげは、各ベンチマークのタスクの中央の 50% でその重みがどのように変化するかを示しています。これはヒューリスティックな近似であり、論文の実際の推定量を再調整したものではありません。メイン チャネル (クロスオーバー近くのゲインがクロスオーバーを動かす) を捉えていますが、オフクロスオーバーのゲインが曲線の傾きをどのように回転させるかは無視し、依然として 0% に近い長いタスクのアンカリング効果も無視します。これは、クロスオーバーが移動できる距離を制限します。カーネル幅を適切に選択した場合、上位のランキングは安定していますが、正確な値を真剣に考慮する必要はありません。基になるデータを 1 つだけベンチマークアウトして再調整するのが、この分析の正しいバージョンです。
時間の数字
一般に、人間の時間を指標として見積もることについては、さまざまな考えがあります。

[切り捨てられた]

## Original Extract

(see full author list at the end) • About a year ago, METR showed that the length of tasks frontier models can reliably complete doubles every few mo…

Login Estimating No-CoT Task-Completion Time Horizons of Frontier AI Models — LessWrong AI Control AI Evaluations Redwood Research Scaling Laws AI Frontpage 2026 Top Fifty: 13 % 167
Estimating No-CoT Task-Completion Time Horizons of Frontier AI Models
(see full author list at the end)
About a year ago, METR showed that the length of tasks frontier models can reliably complete doubles every few months. A related safety-relevant question is this: what length of tasks can models complete without any chain of thought (CoT)? We investigate in our new paper .
If models can do extensive reasoning without outputting any CoT, it would have implications for safety. Developers and deployment-time monitors couldn’t easily understand models’ motivations and catch dangerous planning. Models that reason substantially without a CoT might also drift further from human patterns of thought, since their reasoning is no longer constrained by text in the pretraining prior. As a result, they would be harder to understand and might be more likely to scheme.
Extending Ryan Greenblatt’s research , we investigate this by measuring models' ability to complete tasks without any CoT on a suite of 43 benchmarks spanning different domains. We compare AI reasoning ability to humans using the estimated 50% time horizon (TH)---the typical time taken for a human to perform a task that the LLM performs with 50% success rate. We find that frontier models like GPT-5.5 answer questions that take humans roughly three minutes with 50% reliability, and this TH has doubled approximately every year since 2019.
Figure 1: Our no-CoT THs (green) compared to METR’s with-CoT THs (purple). Until the release of GPT-4, with- and without-CoT THs increased at a similar rate. Since GPT-4 with-CoT THs have grown at roughly twice the rate of no-CoT THs.
We suggest that AI companies start to track no-CoT THs explicitly to find a lower bound on how much reasoning a model could do without revealing it to a CoT monitor or human. Our test suite does not require substantial inference compute so these evaluations are cheap to run.
We evaluated 14 frontier models from GPT-2 (2019) through GPT-5.5 (2026) on 43 benchmarks covering math, coding, knowledge, agentic tool-use, and safety-relevant questions like steganography and scheming reasoning. The questions include single-token answers, longer generation, and multi-turn agentic settings.
In every case we prevent the model from emitting reasoning tokens before its answer by using benchmark-specific prompts and structured-output constraints.
We estimate question difficulty in two ways:
Human solve time : the average time for a human to solve the question with a calculator, paper, and pencil, but no LLM access. For some benchmarks we have human solve time data. For others, we ran small-scale human trials. For the remainder, we used Claude Opus 4.7 estimates, informed via in-context human time examples.
Reasoning token anchor : the minimum number of tokens o3-mini needs to solve the question. o3-mini attempts each question 8 times, and we take the minimum reasoning tokens among successful attempts. If no attempt solves a given question, we drop the question. It's rare for any evaluated model without CoT to solve a question that o3-mini fails to solve with reasoning.
These two methods give us independent measures of “how much reasoning” is needed to solve a problem.
For each model, we estimate its 50% solve horizon on both of these metrics using the same method as Kwa et al. We fit a logistic curve to success rate compared to log human solve time and log reasoning token anchor, and then find the point where the logistic equals 50%.
Figure 2: For each model, we fit a logistic curve to predict the probability it successfully completes tasks as a function of human time (left) or o3-mini reasoning tokens (right) needed to complete the task.
To estimate uncertainty, we rerun the calculation 10,000 times, resampling with replacement over benchmarks, questions within each benchmark, and model runs per question. We also incorporate uncertainty in human solve times by calibrating a noise model on all questions. In contrast, we don’t model uncertainty in the reasoning token anchor.
Figure 3: No-CoT TH over time. Each point is a model plotted by release date; lines are exponential fits, with shading showing 95% CIs.
Figure 4: T oken horizon over time. Each point is a model plotted by release date; lines are exponential fits, with shading showing 95% CIs.
The TH doubles every 373 days (95% CI: 167–691).
The token horizon doubles every 437 days (95% CI: 341–571).
GPT-5.5 is at the current frontier, with a 50% no-CoT TH of about 3 minutes and a reasoning token horizon of about 1,500 o3-mini tokens .
Extrapolating the median forward gives ~7 minutes / ~3.7k tokens by 2028, and ~25 minutes / ~12k tokens by 2030. The doubling-time CI is wide, but even the slow end of the 95% CI puts us at almost 10 minutes of latent reasoning by 2030.
The doubling is roughly robust to changing the distribution of benchmarks: adding longer-generation and multi-turn agentic questions, restricting to questions with measured human times only, dropping any single model, and breaking down by question domain (where domain-specific doublings range from 246 to 387 days).
Why is our trend different from the one in Kwa et al.?
Kwa et al. reported a doubling time of 182 days for THs on agentic software tasks with-CoT , which is about twice as fast as ours. The doubling times diverge after GPT-4, since most capability gains since GPT-4 have come from externalized reasoning, rather than more reasoning in a single forward pass.
Could we be tracking memorization rather than reasoning?
Our benchmarks include some factual recall questions, so we aren't purely measuring reasoning. We include many reasoning-heavy questions, so we think the headline number tracks reasoning well, but imperfectly.
Are models actually hiding their reasoning?
We're not claiming current models are actually hiding reasoning from CoT. We measure how much they could plausibly do invisibly, not how much they actually do. Some models verbalize their reasoning, even when it's unnecessary.
Could models do more hidden reasoning?
We measure a lower bound on opaque reasoning:
Forcing short answers may be out-of-distribution and weird for reasoning models, which would push our estimates down. Better elicitation, or a different distribution of questions could improve model performance.
We measure what a model can do without any externalized reasoning tokens. But a model could learn to encode reasoning into apparently-innocuous output and so do more than this floor while still evading a CoT monitor (“steganography”).
Frontier models can do non-trivial amounts of reasoning without any CoT, and the amount of no-CoT reasoning they can do is increasing exponentially.
Three minutes of human-equivalent reasoning might be enough for some single-shot decisions: whether to comply with an instruction, whether a scenario looks like an eval, or which of several visible options to take. But it's probably not enough to do complex long-term power-seeking reasoning, like deciding to training-game or collude with other instances. If trends continue, models could be capable of twenty-five minutes of no-CoT reasoning by 2030, which would enable much more subversion. Even without neuralese, CoT monitoring might become substantially less effective in the next few years.
Paper authors: Francis Rhys Ward, Dewi Gould, Anders Cairns Woodruff, Rauno Arike, Josh Hills, Alex Serrano, Ida Caspary, Jason Ross Brown, Jo J. Jiao, Patrick Leask, Twm Stone, Ram Potham, Ionut Gabriel Stan, Harry Mayne, Simeon Hellsten, Shubhorup Biswas, Ariana Azarbal, William L. Anderson, Elle Najt, Ryan Greenblatt, and Julian Stastny
Datasets are held back from public release to avoid training contamination; available on request. This was joint work at Redwood Research, the Astra Fellows Program, Aether Research, and MATS; full author list in the paper.
3 Comments 3 AI Control AI Evaluations Redwood Research Scaling Laws AI Frontpage 167
New Comment Submit 3 comments , sorted by top scoring Click to highlight new comments since: Today at 7:24 AM [ - ] Petropolitan 2d 4 2 A very interesting paper quantifying the conventional wisdom that a large degree of progress is from inference scaling.
You have got access to GPT-3 and 3.5 from OpenAI, but have you tried to get the same for now publicly unavailable Claude models for a better coverage of the first half of 2025?
[ - ] dgros 6h * 1 0 Interesting to see. Some things I wonder about, (1) which tasks most influence the results, (2) how much time numbers make sense, and (3) general elicitation questions
Which tasks and kinds of tasks are most influencing the headline?
With METR’s timeline plots there has been discussion here on LW and elsewhere about if certain items have a large influence on the results. I was wondering which kinds of tasks most explain the recent doublings already, and what might worlds look like where they doubled again according to this methodology.
I sort of attempted to vibe very rough versions of this analysis. It's hard without the underlying data, but can get claude to try to extract accuracy numbers from Tables 7-13 and pixel peep Fig 47 to get medians and IQRs for the human-time data. [1]
Things in the ~1-3min band are most likely to be near the 50% crossover point and I think will have more influence on the regression. A few outliers like the "vibe-coding sabotage" on the Y axis and the SHADE monitor on the X axis might deserve some consideration. These are both use a kind of TPR @ ~1-2% FPR metric, different than some other datasets. Unclear effects here.
This is vibey though and would want access to underlying data to do more confident analysis.
There's also other ways of quantify the influence of each dataset (where the "vibe-coding sabaotage" task again shows up as a clear outlier with more than twice as much influence as any other dataset). I hide it in a collapsible though as I'm not super confident in methodology.
Vibey can we calculate an influence score?
Claude's explanation:
The bars estimate how much each benchmark plausibly contributed to the headline time-horizon gain between GPT-4o and GPT-5.5. The headline metric is the 50% crossover of a logistic curve fit to success vs. log human solve time, and a basic property of logistic fits is that an observation's influence on the fitted curve scales with p(1−p) — maximal for tasks the model passes about half the time, near zero for tasks it almost always passes or almost always fails. Since the crossover moved from roughly 0.9 to 3 minutes between these two models, accuracy gains on tasks in that time range supply the most evidence about where the crossover sits. Each bar is therefore the benchmark's accuracy change (chance-corrected, approximately following the paper's normalization — exactly for MCQ tasks, while the detection tasks' ~1.5% baseline was treated as 0, a ~1pp difference) multiplied by a p(1−p) weight centered at ~1.6 minutes and about a decade wide in log time, averaged over the benchmark's per-question human-time distribution as extracted from the paper's Figure 47. The whiskers show how that weight varies across the middle 50% of each benchmark's tasks. This is a heuristic approximation, not a refit of the paper's actual estimator: it captures the main channel (gains near the crossover move the crossover) but ignores how off-crossover gains rotate the curve's slope, and ignores the anchoring effect of long tasks still near 0%, which constrain how far the crossover can move. The ranking at the top is stable across reasonable choices of kernel width, but the precise values shouldn't be taken seriously — a leave-one-benchmark-out refit on the underlying data would be the right version of this analysis.
Time numbers
In general I have mixed thoughts on human time estimates as a metric i

[truncated]

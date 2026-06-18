---
source: "https://www.auriko.ai/reports/llm-cost-arbitrage"
hn_url: "https://news.ycombinator.com/item?id=48588557"
title: "Quantifying LLM Cost Savings from Cache-Aware Inference Routing"
article_title: "Quantifying LLM Cost Savings from Cache-Aware Inference Routing | Auriko"
author: "zxy-action"
captured_at: "2026-06-18T18:55:40Z"
capture_tool: "hn-digest"
hn_id: 48588557
score: 4
comments: 1
posted_at: "2026-06-18T17:21:23Z"
tags:
  - hacker-news
  - translated
---

# Quantifying LLM Cost Savings from Cache-Aware Inference Routing

- HN: [48588557](https://news.ycombinator.com/item?id=48588557)
- Source: [www.auriko.ai](https://www.auriko.ai/reports/llm-cost-arbitrage)
- Score: 4
- Comments: 1
- Posted: 2026-06-18T17:21:23Z

## Translation

タイトル: キャッシュを意識した推論ルーティングによる LLM コスト削減の定量化
記事のタイトル: キャッシュを意識した推論ルーティングによる LLM コスト削減の定量化 |アウリコ
説明: キャッシュ対応 LLM 推論アービトラージに関する統計レポート: 方法論、堅牢性チェック、およびプロバイダー全体の集計結果。

記事本文:
キャッシュを意識した推論ルーティングによる LLM コスト削減の定量化 | Auriko 開発者
メニューを開く 完全な技術レポート
キャッシュ認識推論ルーティングによる LLM コスト削減の定量化
コストの最適化をライブで参照する 機会: プロバイダー間のコストの分散 実験計画のカバレッジ メトリクスの統計手法の妥当性管理 証拠: コスト削減の主な調査結果 ワークロード全体のコスト削減 モデル全体のコスト削減 統計テストの要約表 コスト削減の一貫性 メカニズム: キャッシュを意識したアービトラージ プロバイダーのコスト アービトラージ プロンプト キャッシュ ヒットの堅牢性と感度 ストレステストの概要 リーブワンアウトの安定性 完了正規化 完了一致感度 ワークロード制御による分散 結果の範囲 方法論のメトリクスの用語集 統計的手法の制限事項 同じ LLM セッションは、同じモデル、パラメータ、プロンプト、応答を使用する場合でも、サービスを提供する推論プロバイダーに応じて大幅に異なるコストが発生する可能性があります。このギャップは、トークンの価格設定、プロバイダーのプロンプト キャッシュ動作、ユーザーのリクエスト パターンによって生じます。これらの要因が異なる場合、キャッシュを意識したコスト裁定レイヤーにより支出を削減できます。問題は、どれだけ節約できるか、そしてどれだけ確実に節約できるかです。
私たちは、Auriko のキャッシュ対応コスト裁定取引エンジンを 5 つの推論プロバイダーと 1 つの推論ルーティング ピアに対してベンチマークしました。このベンチマークでは、3 つのワークロード タイプにわたる 37 のモデルに 80,000 を超える API リクエストが送信されました。これらのリクエストは 22,000 を超えるセッションを形成します。 Auriko は、テストされた 6 つのベースラインに対してドル加重支出を削減しました。ルーティング ピアに対して 32.8% (95% CI [30.6%、34.9%])、5 つの単一プロバイダー ベースライン全体では 7.7 ～ 38.3% でした。
このレポートでは、ベンチマーク結果を示し、節約の背後にあるコスト要因を説明し、方法論、統計情報を文書化します。

テストと堅牢性チェック。
機会: プロバイダ間のコストの分散
同じリクエストされたモデルの推論コストはプロバイダーによって異なります。私たちがテストした一部のモデルでは、同じリクエストに対して最も高価なパスのコストが最も安いパスの 4 倍でした。
プロバイダー全体で 37 モデルを比較しました。各モデルについて、各プロバイダーの平均コストを最低コストのプロバイダーの平均コストで割り、100 を掛けました。スコア 100 は最も安いことを意味します。スコア 300 は、同じリクエストのコストが 3 倍であることを意味します。
変動は研究内のモデル全体で見られます。このパターンはワークロードの種類全体に当てはまります。コーディング エージェントのワークロードには最も大きなギャップが見られます。シングルターンと会話のワークロードは同等の広がりを示します (詳細は堅牢性と感度を参照)。
このコストの分散は、テストされたプロバイダー セット内で、測定されたワークロードごとに最も安価なプロバイダーの選択肢が存在することを意味します。このレポートの残りの部分では、Auriko のキャッシュを意識したルーティング アプローチで観察されたコスト削減の統計的研究を示します。次のセクションでは、調査をどのように実施したかについて説明します。
この研究では、マッチドペアデザインを採用しています。モデルとワークロードの組み合わせごとに、Auriko と各コンパレーターへのすべての呼び出しで同一のプロンプトが使用されます。リクエスト パラメータ (温度、top-p、周波数ペナルティ、プレゼンス ペナルティ、および最大出力長) は、すべての呼び出しにわたって一定に保たれます。時間帯の影響を制御するために、両側が同時に派遣されます。それぞれの組み合わせが複数回繰り返されます。ペアのいずれかの側にエラーがある場合、分析前に両方の側が除外されます。
行われたリクエストは、個々の API 呼び出しの合計数です。 LLM セッションは、複数ターンの会話や複数ステップのエージェントの実行など、関連するコールをワークロード セッションにグループ化します。要求されたペアのセッションは、エラー除外前のターゲット レベルのセッション ペアです。

n;クリーンなペアのターゲット比較は、どちらかの側にエラーがあった場合の対称除外後の同じペアです。
コストは、各 API 応答によって返される金額です。セッションレベルの分析の場合、リクエストのコストは各ワークロードセッション内で合計されます。ルーティング ピアの場合、コストはプラットフォーム料金に合わせて調整され、そのサービスのユーザーが実質的に支払う金額に一致します。
調整後のコスト = API から返されたコスト × (1 + プラットフォーム手数料 %)
Auriko はプラットフォーム料金を請求しないため、コスト調整は必要ありません。
分散セクションでは、同じ要求モデルでも、プロバイダーが異なると大幅に異なる金額がかかる可能性があることを示しました。このセクションでは、各比較対象に対する Auriko の実現されたコスト削減を示します。
Auriko は 6 つの比較対象ターゲットに対してドル加重支出を削減し、ターゲットレベルの節約率は 7.7% ～ 38.3% でした。
以下のグラフは、各コンパレータ ターゲットのセッション レベルのコスト差の分布と、勝ち、引き分け、負けの割合を示しています。勝利とは、Auriko のコストがコンパレータよりも低いセッションを指します。 Auriko は、6 つのターゲットにわたる非結合セッションの 60 ～ 90% でコストが低くなります。
以下のグラフは、同点を除いた 95% 信頼区間でのセッション勝率を示しています。 Auriko は、信頼区間が 50% の偶数オッズ ラインを上回っており、各比較対象ターゲットに対して同数でないセッションの大部分でコストが低くなります。
以下のグラフは、各比較対象について、95% 信頼区間を使用したドル加重節約額を示しています。節約範囲は 7.7% (モデル プロバイダー E) から 38.3% (モデル プロバイダー A) です。信頼区間は層別ブートストラップによって計算されます。
これらのパネルは、複数ターンの会話でコストがどのように蓄積されるかを示します。各線は、そのコンパレータに対してテストされたモデル全体の平均累積コストです。線の間のギャップは節約であり、それは広いです

ターンごとに、ルーティングの選択と一致するパターンと、再利用可能なプロンプト コンテキストが、ターンを重ねるごとに価値が高まります。
以下のモデルごとの内訳は、累積コストのギャップが個々のモデル レベルでも見られることを示しています。
ワークロード全体のコスト削減
各セルは、1 つのワークロード シナリオとペアになった 1 つのコンパレータ ターゲットを表し、そのペアのドル加重節約率の注釈が付けられます。色は大きさをコード化します。コスト削減は、ワークロードの種類と比較対象全体にわたってプラスになります。
ほとんどのターゲットでは、マルチターン ワークロードの方がシングルターン リクエストよりも大幅なコスト削減が見られ、これは、再利用可能なプロンプト コンテキストが連続するターンで価値が高まることと一致しています。
平均的な節約に対する一般的な反対意見は、コスト差が大きい 1 つのモデルが全体の結果を左右する可能性があるというものです。このサブセクションでは、テストされたモデル セット全体にわたって節約が広範囲に及ぶかどうかをテストします。
ホルム有意モデルは、比較グループ内のすべてのモデルレベルのテストにわたる保守的なファミリーワイズ補正であるホルム・ボンフェローニ多重比較補正を生き残りました。
各ドットは、95% のブートストラップ信頼区間を使用した 1 つのモデルのドル加重節約額です。青色のマーカーは、ホルム・ボンフェローニ補正を生き残ったモデルを示します。ゼロに近い、またはゼロ未満のモデルは、サンプル サイズが小さいことを反映して広い信頼区間を持ちます。
各コンパレータ ターゲットについて、セッション レベルのコスト差 d = コンパレータ コスト − Auriko コストを計算しました。正の値は、Auriko の方が安かったことを示します。私たちは、2 つのノンパラメトリック検定を使用して、Auriko には低いセッションレベルのコストがない (H₀: d の分布はゼロを中心とする) という帰無仮説を片側代替 (H₁: d が正にシフトされる) に対して検定しました。
符号テストでは、Auriko のコストがコンパレーターよりも低い場合とそれよりも高い場合の非同率セッションがカウントされます。

低コストの割合が 50% を超えるかどうかをテストします。 Wilcoxon の符号付き順位検定では、絶対的な差を順位付けし、その後、正の順位が系統的に負の順位を上回るかどうかをテストすることで絶対値を組み込みます。ゼロの差はランキングに含まれます (Wilcox 法)。両方のテストは補完的であるため、両方のテストを報告します。符号テストは外れ値に対して堅牢です (単一の極端なセッションでは結果を導き出すことができません)。一方、Wilcoxon はランク情報を使用するため、より強力です。
どちらのテストも、6 つのコンパレータ ターゲットについて p < 0.001 で H₀ を拒否し、p 値の範囲は 10⁻¹² から 10⁻²⁵⁴ です。
ドル加重節約額 = (コンパレータの合計支出額 − Auriko の合計支出額) / コンパレータの合計支出額。 95% CI は、ドル加重節約に関する層別ブートストラップ間隔 (5,000 回の反復、モデル + シナリオ層) です。
モデル プロバイダー E は、ターゲット レベルの節約が最も狭く、タイ レートが最も高くなります。セッションの 30.2% がタイであるため、生の低コスト率は 42.5% です。同点を除くと、非同点セッションの 60.9% では Auriko のコストが低くなります。
コホートから直接の最も高い節約額のベースラインを除いても、モデルプロバイダーコホートの節約率は約 16.8% にとどまります。コホートの結果は、単一の匿名化されたベースラインには含まれません。
前のセクションでは、6 つのコンパレータ間で統計的に有意なコストの違いを報告しました。統計的有意性は帰無仮説に対する証拠を提供しますが、大きさを定量化するものではありません。十分なサンプル サイズがあれば、無視できるほどの差であっても有意な p 値が生成されます。効果の大きさは、サンプルサイズとは無関係に、セッション間のコストの変動と比較して節約がどの程度大きいかを測定します。
ヘッジズの g: セッションごとの平均ドル差 (コンパレータから Auriko を引いたもの) を標準偏差で割った値を、サンプルが小さいバイアスで報告します。

補正が適用されました。値が大きいほど、セッション間で一貫した節約効果があることを示します。値が低いほど、平均的にはプラスの節約効果があるものの、セッションごとの変動が大きいことを示します。
ヘッジの g の範囲は 0.21 (モデル プロバイダー E、N = 3,077) から 0.65 (モデル プロバイダー A、N = 1,647) です。 6 つのコンパレータは正の効果量を生成します。
ターゲットレベルの結果は、大きさの問題に答えます。次の疑問は、どの測定コスト要素がこれらの節約に関連するかということです。
仕組み: キャッシュを意識したアービトラージ
LLM リクエストの実際のコストは、ヘッドライン トークンの価格以外にも依存します。プロバイダー側​​では、キャッシュ アクティベーションの最小トークンしきい値、ブロック粒度、キャッシュ有効期限のタイミング、書き込みコストとストレージ コスト、およびコンテキスト長の価格階層という 5 つの変数が実現コストを形成します。ユーザー側では、ワークロード パターン (プレフィックス長、トークン再利用率、リクエストのタイミング、会話の深さ) によって、各プロバイダーのキャッシュ メカニズムがどの程度効果的にアクティブ化されるかが決まります。これら 2 つの側間の対話により、最低コストのプロバイダーがリクエストごとに異なることになります。コストモデルの詳細については、「Auriko による LLM 推論コストの最適化方法」を参照してください。
このセクションでは、プロバイダーのアービトラージとより高いプロンプト キャッシュ レートという、Auriko が実現したコスト上の利点を説明するのに役立つ 2 つのベンチマーク測定を検証します。まず、Auriko のルーティング コストをテスト済みのプロバイダーのコスト範囲と比較します。次に、プロンプト キャッシュのヒット率の向上について見ていきます。
コストの分散セクションでは、同じリクエストのモデルでもプロバイダーが異なると最大 4 倍のコストがかかる可能性があることを示しました。このチャートは、Auriko の実現コストが分散内でどこに現れるかを示しています。各モデルの灰色のバーは、テストされた 5 つの単一プロバイダーのベースラインにわたるコスト範囲を示しています。 Auriko の実現コストは範囲の下限に近いです。ルーティングピアの

範囲内ではコストが高くなります。
プロバイダーは、キャッシュから提供されるプロンプト トークンに対して割引料金を請求します。 Auriko は 6 つのコンパレーター ターゲットに対して、トークン重み付けキャッシュ ヒット率 (総入力トークンに占めるキャッシュ読み取りトークンの割合) が高く、デルタの範囲は +2.9 ～ +11.6 パーセント ポイントでした。
前のグラフは、キャッシュから提供される入力トークンのシェアを測定しました。このグラフは、キャッシュ ヒットのあるリクエストの割合という別の次元を測定します。 Auriko は、6 つのターゲットのうち 5 つに対してリクエスト レベルのヒット率が高くなります。モデル プロバイダー C は例外で、コンパレーターが 2.5 パーセント ポイントリードしています。
セッション レベルのキャッシュ ヒット率のデルタは、6 つのコンパレータ ターゲットにわたってコスト削減と正の相関関係があります (Spearman ρ = 0.26 ～ 0.61)。
Auriko のより高いキャッシュ ヒット率は、キャッシングの仕組みが測定された使用パターンに適合するプロバイダー ルートを選択することと一致しています。会話が続くにつれて、再利用可能なプロンプト コンテキストが増加し、測定されたキャッシュ ヒット ギャップが拡大します。
すべての比較の集計結果を、分析サンプルの 242 の摂動に対してテストしました。各シナリオ、比較ターゲット、またはモデルを一度に 1 つずつ削除しました。エラーセッションを再含める。完了トークンのコストを正規化する。分析をセッションに限定する

[切り捨てられた]

## Original Extract

Statistical report on cache-aware LLM inference arbitrage: methodology, robustness checks, and aggregate results across providers.

Quantifying LLM Cost Savings from Cache-Aware Inference Routing | Auriko Developers
Open menu Full technical report
Quantifying LLM Cost Savings from Cache-Aware Inference Routing
See cost optimization live The Opportunity: Cost Dispersion Across Providers Experimental Design Coverage Metric Statistical Methods Validity Controls The Evidence: Cost Reduction Primary Findings Cost Reduction Across Workloads Cost Reduction Across Models Statistical Tests Summary Table Cost Reduction Consistency The Mechanics: Cache-Aware Arbitrage Provider Cost Arbitrage Prompt-Cache Hit Measures Robustness and Sensitivity Stress-Test Summary Leave-One-Out Stability Completion Normalization Completion-Matched Sensitivity Workload-Controlled Dispersion Scope of Results Methodology Glossary of Metrics Statistical Methods Limitations The same LLM session, even with identical model, parameters, prompts and responses, can incur substantially different costs depending on the inference provider that serves it. The gap comes from token pricing, provider prompt-caching behavior, and users’ request patterns. When those factors diverge, a cache-aware cost-arbitrage layer can reduce spend. The question is how much it saves and how reliably.
We benchmarked Auriko's cache-aware cost-arbitrage engine against five inference providers and an inference routing peer. The benchmark sent over 80,000 API requests to 37 models across 3 workload types. These requests form over 22,000 sessions. Auriko reduced dollar-weighted spend against six tested baselines: 32.8% against the routing peer (95% CI [30.6%, 34.9%]) and 7.7–38.3% across five single-provider baselines.
This report presents the benchmark results, explains the cost drivers behind the savings, and documents the methodology, statistical tests, and robustness checks.
The Opportunity: Cost Dispersion Across Providers
Inference cost for the same requested model varies across providers. For some models we tested, the most expensive path costs 4x the cheapest for an identical request.
We compared 37 models across providers. For each model, we divided each provider's average cost by the lowest-cost provider's average cost and multiplied by 100. A score of 100 means cheapest. A score of 300 means three times the cost for the same request.
The variation appears across models in the study. The pattern holds across workload types. Coding agent workloads show the widest gap; single-turn and conversation workloads show comparable spreads (details in Robustness and Sensitivity).
This cost dispersion means that within the tested provider set, a cheapest provider choice exists for each measured workload. The rest of this report presents a statistical study of the cost reduction observed for Auriko's cache-aware routing approach. The next section describes how we conducted the study.
This study adopts a matched-pairs design. For each model and workload combination, every call to Auriko and each comparator uses identical prompts. Request parameters — temperature, top-p, frequency penalty, presence penalty, and maximum output length — are held constant across all calls. Both sides are dispatched concurrently to control for time-of-day effects. Each combination is repeated multiple times. If either side of a pair errors, both sides are excluded before analysis.
Requests made is the total number of individual API calls. LLM sessions group related calls into workload sessions, including multi-turn conversations and multi-step agent runs. Requested paired sessions are target-level session pairings before error exclusion; clean paired target comparisons are the same pairings after symmetric exclusion when either side errored.
Cost is the dollar amount returned by each API response. For session-level analysis, request costs are summed within each workload session. For the routing peer, cost is adjusted for their platform fee, matching what a user of that service effectively pays:
adjusted cost = API-returned cost × (1 + platform fee %)
Since Auriko charges no platform fees, no cost adjustment is needed.
The dispersion section showed that the same requested model can cost materially different amounts on different providers. This section presents Auriko's realized cost savings against each comparator target.
Auriko reduced dollar-weighted spend against six comparator targets, with target-level savings ranging from 7.7% to 38.3%.
The chart below shows the distribution of session-level cost differences for each comparator target, together with the percentage of wins, ties, and losses — where a win is a session where Auriko costs less than the comparator. Auriko costs less in 60–90% of non-tie sessions across six targets.
The chart below shows session win rates with 95% confidence intervals, ties excluded. Auriko costs less in the majority of non-tie sessions against each comparator target, with confidence intervals above the 50% even-odds line.
For each comparator target, the chart below shows dollar-weighted savings with 95% confidence intervals. Savings range from 7.7% (Model provider E) to 38.3% (Model provider A). Confidence intervals are computed via stratified bootstrap.
These panels show how cost accumulates over a multi-turn conversation. Each line is the average cumulative cost across models tested against that comparator. The gap between the lines is the savings, and it widens with each turn, a pattern consistent with routing choices and reusable prompt context becoming more valuable over successive turns.
The per-model breakdown below shows the cumulative-cost gap is also visible at the individual-model level.
Cost Reduction Across Workloads
Each cell represents one comparator target paired with one workload scenario, annotated with the dollar-weighted savings percentage for that pair. Color encodes magnitude. Cost reduction is positive across workload types and comparator targets.
For most targets, multi-turn workloads show larger cost reduction than single-turn requests, consistent with reusable prompt context becoming more valuable over successive turns.
A common objection to average savings is that one model with a large cost difference could drive the entire result. This subsection tests whether savings are broad across the tested model set.
Holm-significant models survived Holm–Bonferroni multiple-comparison correction, a conservative family-wise correction across all model-level tests within the comparator group.
Each dot is one model's dollar-weighted savings with a 95% bootstrap confidence interval. Blue markers indicate models that survived Holm–Bonferroni correction. Models near or below zero have wide confidence intervals reflecting small sample sizes.
For each comparator target, we computed the session-level cost difference d = comparator cost − Auriko cost, where positive values indicate Auriko was cheaper. We tested the null hypothesis that Auriko has no lower session-level cost (H₀: the distribution of d is centered at zero) against the one-sided alternative (H₁: d is shifted positive) using two nonparametric tests.
The sign test counts non-tie sessions where Auriko costs less than the comparator versus more, and tests whether the lower-cost proportion exceeds 50%. The Wilcoxon signed-rank test incorporates magnitude by ranking the absolute differences, then testing whether positive ranks systematically outweigh negative ranks. Zero differences are included in the ranking (Wilcox method). We report both tests because they are complementary: the sign test is robust to outliers (a single extreme session cannot drive the result), while the Wilcoxon is more powerful because it uses rank information.
Both tests reject H₀ at p < 0.001 for six comparator targets, with p-values ranging from 10⁻¹² to 10⁻²⁵⁴.
Dollar-weighted savings = (total comparator spend − total Auriko spend) / total comparator spend. 95% CI is a stratified bootstrap interval (5,000 replicates, model + scenario strata) around dollar-weighted savings.
Model provider E has the narrowest target-level savings and the highest tie rate. Its raw lower-cost rate is 42.5% because 30.2% of sessions are ties; excluding ties, Auriko costs less in 60.9% of non-tie sessions.
Dropping the highest-savings direct baseline from the cohort, model provider cohort savings remain at approximately 16.8%. The cohort result is not carried by a single anonymized baseline.
The preceding sections reported statistically significant cost differences across six comparators. Statistical significance provides evidence against the null hypothesis, but does not quantify magnitude — with sufficient sample size, even a negligible difference produces a significant p-value. Effect size measures how large the savings are relative to session-to-session cost variability, independent of sample size.
We report Hedges' g: the mean per-session dollar difference (comparator minus Auriko) divided by its standard deviation, with a small-sample bias correction applied. Higher values indicate savings that are consistent across sessions; lower values indicate savings that are positive on average but more variable per session.
Hedges' g ranges from 0.21 (Model provider E, N = 3,077) to 0.65 (Model provider A, N = 1,647). Six comparators produce positive effect sizes.
The target-level results answer the magnitude question. The next question is which measured cost components are associated with these savings.
The Mechanics: Cache-Aware Arbitrage
The actual cost of an LLM request depends on more than headline token price. On the provider side, five variables shape the realized cost: minimum token thresholds for cache activation, block granularity, cache expiration timing, write and storage costs, and context-length pricing tiers. On the user side, the workload pattern (prefix lengths, token reuse rates, request timing, and conversation depth) determines how effectively each provider's caching mechanics activate. The interaction between these two sides means the lowest-cost provider varies per request. For a full treatment of the cost model, see How Auriko Optimizes LLM Inference Cost .
This section examines two benchmark measurements that help explain Auriko's realized cost advantage: provider arbitrage and higher prompt-cache rate. First, we compare Auriko's routed cost with the tested provider cost range. Then, we look at prompt-cache hit rate improvements.
The cost dispersion section showed that the same requested model can cost up to 4× more on different providers. This chart shows where Auriko's realized cost appears within that dispersion. For each model, the gray bar spans the cost range across the five tested single-provider baselines. Auriko's realized cost is near the low end of the range. The routing peer's cost is higher in the range.
Providers charge a reduced rate for prompt tokens served from cache. Auriko had a higher token-weighted cache hit rate (cache-read tokens as a share of total input tokens) against six comparator targets, with deltas ranging from +2.9 to +11.6 percentage points.
The previous chart measured the share of input tokens served from cache. This chart measures a different dimension: the share of requests with any cache hit. Auriko has a higher request-level hit rate against five of six targets. Model provider C is the exception, where the comparator leads by 2.5 percentage points.
Session-level cache hit rate delta correlates positively with cost reduction across six comparator targets (Spearman ρ = 0.26 to 0.61).
Auriko's higher cache hit rate is consistent with selecting provider routes whose caching mechanics fit the measured usage pattern. Over successive conversation turns, the reusable prompt context grows and the measured cache-hit gap widens.
We tested the all-comparison aggregate result against 242 perturbations of the analysis sample: removing each scenario, comparator target, or model one at a time; re-including error sessions; normalizing completion-token costs; and restricting the analysis to sessions

[truncated]

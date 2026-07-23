---
source: "https://www.coveragecat.com/ai/leaderboard"
hn_url: "https://news.ycombinator.com/item?id=49028661"
title: "Which AI models understand insurance work?"
article_title: "AI Insurance Leaderboard | Coverage Cat"
author: "botacode"
captured_at: "2026-07-23T22:51:53Z"
capture_tool: "hn-digest"
hn_id: 49028661
score: 1
comments: 0
posted_at: "2026-07-23T22:03:26Z"
tags:
  - hacker-news
  - translated
---

# Which AI models understand insurance work?

- HN: [49028661](https://news.ycombinator.com/item?id=49028661)
- Source: [www.coveragecat.com](https://www.coveragecat.com/ai/leaderboard)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T22:03:26Z

## Translation

タイトル: 保険業務を理解できる AI モデルはどれですか?
記事のタイトル: AI 保険リーダーボード |取材猫
説明: AI 保険検索ツール

記事本文:
補償範囲 Cat AI 保険ベンチマーク
保険業務を理解できる AI モデルはどれですか?
価格推定という 2 つの別個の保険タスクでフロンティア モデルを評価します。
匿名化されたアンブレラ相場行と、それに対する仲介/エージェントタスクの推論用
引受業務、適格性、および補償範囲に関する質問に対するベンチマーク参照回答。
価格見積りの結果は、予測される保険料と不確実性の範囲を実際の見積り結果と比較します。これらの指標は、仲介/エージェントのタスクのリーダーボードとは別のものです。
同じ保険見積もりケースにおけるペアワイズ モデルの強度。スコアが高いほど、モデルが校正された見積もりの​​精度で同等のモデルを上回る頻度が高いことを意味します。
ペアごとの引用バトルのシェアが勝利し、同点は半分の勝利としてカウントされます。高いほど良いです。
実際の年換算保険料がモデルの予測相場範囲内に収まる頻度。高いほど良いです。
各モデルのポイントプレミアム推定値の平均絶対パーセント誤差。低いほど良いです。
実際のプレミアムを逃す、または不必要に広い見積範囲に対する調整ペナルティ。低いほど良いです。
価格見積もり行は、実際の見積もり結果に対してスコア付けされます。仲介・代理業務
行は参照回答に対してスコア付けされ、個別に判定されるため、そのリーダーボードは
プレミアム推定のパフォーマンスではなく、回答品質のパフォーマンスとして解釈する必要があります。
特定の通信事業者からの匿名化された 100 万ドルのカリフォルニア州アンブレラ見積もりの​​年間保険料と不確実性の範囲を推定します。
州、通信事業者、補償範囲の制限、および匿名化されたリスク特徴を考慮して、調整された P10/P50/P90 保険料の推定値を返します。
間隔を不必要に広くすることなく、実際の年換算保険料を含む見積範囲を予測します。
世帯には引受プロファイルがリストされています。特定の傘持ち手の対象となる可能性はありますか?
テキサスではどうやって

ましてや、指定通信会社を利用してアンブレラ補償範囲を 100 万ドルから 200 万ドルに移行するのには通常費用がかかりますか?
顧客から補償要件や資格の制約について尋ねられます。ベンチマーク参照の回答を使用して、アシスタントは何と言えばよいでしょうか?
ドメイン固有の集計のみのベンチマーク
一般的な AI ベンチマークでは、モデルが以下の詳細を推論できるかどうかを測定することはほとんどありません。
保険における事項: 賠償責任の制限、運送業者の制約、保険料の範囲、資格
ルールと不確実性。このベンチマークは、これらのワークフローに焦点を当てています。
見積行では、各モデルの推定年間保険料と範囲を実際の保険料と比較します。
見積もり結果。カバレッジ報酬の調整された範囲。 MAPE は正確なポイント推定値を提供します。
ウィンクラー損失は、実際の相場を逃したレンジや広すぎるレンジにペナルティを与えます。
仲介/エージェントのタスク行では、モデルの回答とベンチマークの参照回答を比較します。
AI判定。パブリック タスク ビューは、ジャッジのスコア、ペアごとの Elo、勝利を集計してレポートします。
レートのみ。
共有シナリオまたは質問ごとに、モデル出力のすべてのペアが比較されます。より良い
出力はローカルの戦いに勝ち、分割されたクレジットを結び付け、Elo はモデルの強度を内部で更新します。
そのベンチマークセクション。
公開結果は集計のみです。このページでは、生のプロンプト、行識別子、
モデル応答、裁判官推論、または操作上の評価成果物。 eval は使用します
保持もログも記録されないプラットフォーム上で匿名化されたデータを使用するため、顧客データが漏洩することはありません
モデルプロバイダーにも公開されます。

## Original Extract

Your AI insurance search tool

Coverage Cat AI Insurance Benchmark
Which AI models understand insurance work?
We evaluate frontier models on two separate insurance tasks: price estimation
for anonymized umbrella quote rows, and brokerage/agent task reasoning against
benchmark reference answers for underwriting, eligibility, and coverage questions.
Price-estimation results compare predicted premiums and uncertainty ranges against actual quote outcomes. These metrics are separate from the brokerage/agent task leaderboard.
Pairwise model strength on the same insurance quoting cases. Higher scores mean a model more often beat comparable models on calibrated quote accuracy.
The share of pairwise quote battles won, with ties counted as half a win. Higher is better.
How often the actual annualized premium landed inside the model's predicted quote range. Higher is better.
Mean absolute percentage error for each model's point premium estimate. Lower is better.
A calibration penalty for quote ranges that miss the actual premium or are unnecessarily wide. Lower is better.
Price-estimation rows are scored against actual quote outcomes. Brokerage/agent task
rows are scored against reference answers and judged separately, so their leaderboard
should be read as answer-quality performance rather than premium-estimation performance.
Estimate the annual premium and uncertainty range for an anonymized $1M California umbrella quote from a specific carrier.
Given state, carrier, coverage limit, and anonymized risk features, return calibrated P10/P50/P90 premium estimates.
Predict a quote range that contains the actual annualized premium without making the interval unnecessarily wide.
A household has a listed underwriting profile. Are they likely to be eligible with a specific umbrella carrier?
In Texas, how much more does moving from $1M to $2M of umbrella coverage typically cost with a named carrier?
A customer asks about coverage requirements or eligibility constraints. What should an assistant say, using the benchmark reference answer?
Domain-specific, aggregate-only benchmarking
General AI benchmarks rarely measure whether a model can reason through the details that
matter in insurance: liability limits, carrier constraints, premium ranges, eligibility
rules, and uncertainty. This benchmark focuses on those workflows.
Quote rows compare each model's estimated annual premium and range against the actual
quote outcome. Coverage rewards calibrated ranges; MAPE rewards accurate point estimates;
Winkler loss penalizes ranges that miss the actual quote or are too wide.
Brokerage/agent task rows compare model answers to benchmark reference answers with
AI judging. The public task view reports aggregate judge scores, pairwise Elo, and win
rate only.
For each shared scenario or question, every pair of model outputs is compared. Better
outputs win the local battle, ties split credit, and Elo updates model strength within
that benchmark section.
Public results are aggregate-only. The page does not expose raw prompts, row identifiers,
model responses, judge reasoning, or any operational eval artifacts. The evals use
anonymized data on no-retention and no-logging platforms, so customer data is never
exposed even to model providers.

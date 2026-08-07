---
source: "https://zenodo.org/records/21841560"
hn_url: "https://news.ycombinator.com/item?id=49213283"
title: "LLM-generated tests can change which implementation wins"
article_title: "Valid Evidence Is Not Necessarily Comparable Evidence: Candidate-Conditioned Non-Invariance in LLM-Generated Code Verification | Zenodo"
author: "haitamk"
captured_at: "2026-08-07T17:41:08Z"
capture_tool: "hn-digest"
hn_id: 49213283
score: 1
comments: 0
posted_at: "2026-08-07T16:58:20Z"
tags:
  - hacker-news
  - translated
---

# LLM-generated tests can change which implementation wins

- HN: [49213283](https://news.ycombinator.com/item?id=49213283)
- Source: [zenodo.org](https://zenodo.org/records/21841560)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T16:58:20Z

## Translation

タイトル: LLM で生成されたテストにより、どの実装が優先されるかが変わる可能性がある
記事のタイトル: 有効な証拠は必ずしも比較可能ではない証拠: LLM 生成コード検証における候補条件付き非不変性 |ゼノド
説明: 候補コードを評価するためのテストやその他の実行可能な証拠を生成するために、大規模な言語モデルがますます使用されています。これは、候補者を採点するために使用される証拠自体がその候補者を観察した後に生成される場合に、比較測定の問題を引き起こします。この作品は候補者条件を研究します
[切り捨てられた]

記事本文:
有効な証拠は必ずしも比較可能ではない証拠: LLM 生成コード検証における候補条件付き非不変性 |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
計画された介入 : 8 月 13 日木曜日 06:15 UTC に、Zenodo はストレージ クラスターのアップグレードを実行するために 3 ～ 5 分間利用できなくなります。
有効な証拠は必ずしも比較可能である必要はない 証拠: LLM 生成コード検証における候補条件付き非不変性
候補コードを評価するためのテストやその他の実行可能な証拠を生成するために、大規模な言語モデルがますます使用されています。これは、候補者を採点するために使用される証拠自体がその候補者を観察した後に生成される場合に、比較測定の問題を引き起こします。
この研究では、制御された交差計画を使用して、候補条件付きの比較非不変性を研究します。 8 つのセマンティック ファミリにまたがる 16 の要件にわたって、テスト生成中にどの候補が表示されるかを変更すると、同じ固定候補ペアの相対評価が変化し、16 要件のうち 5 つで厳密な勝者の逆転が生じました。この効果は有効のみの分析では持続しましたが、要件レベルの分析では、集約された方向が要件全体で普遍的なものとして解釈されるべきではないことが示されています。
別の制御された欠陥注入の研究では、候補認識検証の補完的な利点が示されています。欠陥のある実装を明らかにすることで、ターゲットを絞った有効な欠陥検出が 37.5% から 53.44% に増加しました。これらの結果は、診断の有用性と比較適合性を区別します。
中心的な結論は、仕様の妥当性は項目ごとの特性であり、それ自体では候補間の比較可能性をもたらさないということです。候補者を意識した証拠は診断に役立ちますが、

候補者を直接比較するには、共有された証拠が望ましいです。
付随するアーティファクトには、凍結された実験構成、実行結果、整合性監査、分析出力、再現性情報が含まれています。
候補条件付き非不変性v0.1-final.pdf
統計の収集方法の詳細....
10.5281/ゼノド.21841560
マークダウン
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg)](https://doi.org/10.5281/zenodo.21841560)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg
:target: https://doi.org/10.5281/zenodo.21841560
HTML
<a href="https://doi.org/10.5281/zenodo.21841560"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg
ターゲット URL
https://doi.org/10.5281/zenodo.21841560
リソースの種類
プレプリント
出版社
ゼノド
言語
英語
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

Large language models are increasingly used to generate tests and other executable evidence for evaluating candidate code. This creates a comparative measurement problem when the evidence used to score a candidate is itself generated after observing that candidate. This work studies candidate-condit
[truncated]

Valid Evidence Is Not Necessarily Comparable Evidence: Candidate-Conditioned Non-Invariance in LLM-Generated Code Verification | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Planned intervention : On Thursday, August 13th, 06:15 UTC, Zenodo will be unavailable for 3-5 minutes to perform a storage cluster upgrade.
Valid Evidence Is Not Necessarily Comparable Evidence: Candidate-Conditioned Non-Invariance in LLM-Generated Code Verification
Large language models are increasingly used to generate tests and other executable evidence for evaluating candidate code. This creates a comparative measurement problem when the evidence used to score a candidate is itself generated after observing that candidate.
This work studies candidate-conditioned comparative non-invariance using a controlled crossed design. Across 16 requirements spanning eight semantic families, changing which candidate was visible during test generation shifted the relative evaluation of the same fixed candidate pair and produced strict winner reversals in 5 of 16 requirements. The effect persisted under a valid-only analysis, while requirement-level analyses show that the aggregate direction should not be interpreted as universal across requirements.
A separate controlled defect-injection study shows the complementary benefit of candidate-aware verification: exposing the defective implementation increased targeted valid defect detection from 37.5% to 53.44%. These results distinguish diagnostic usefulness from comparative suitability.
The central conclusion is that specification validity is a per-item property and does not by itself confer cross-candidate comparability. Candidate-aware evidence can be useful for diagnosis, while shared evidence is preferable for direct candidate comparison.
The accompanying artifact contains frozen experimental configurations, execution results, integrity audits, analysis outputs, and reproducibility information.
candidate-conditioned-noninvariance-v0.1-final.pdf
More info on how stats are collected....
10.5281/zenodo.21841560
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg)](https://doi.org/10.5281/zenodo.21841560)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg
:target: https://doi.org/10.5281/zenodo.21841560
HTML
<a href="https://doi.org/10.5281/zenodo.21841560"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21841560.svg
Target URL
https://doi.org/10.5281/zenodo.21841560
Resource type
Preprint
Publisher
Zenodo
Languages
English
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies

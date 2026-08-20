---
source: "https://vulnbench.com/"
hn_url: "https://news.ycombinator.com/item?id=49373152"
title: "VulnBench: Can LLMs find the same security bugs twice?"
article_title: "Snyk VulnBench"
image: "https://vulnbench.com/brand/snyk-2026/social.png"
author: "lirantal"
captured_at: "2026-08-20T12:26:46Z"
capture_tool: "hn-digest"
hn_id: 49373152
score: 1
comments: 0
posted_at: "2026-08-20T11:27:20Z"
tags:
  - hacker-news
  - translated
---

# VulnBench: Can LLMs find the same security bugs twice?

- HN: [49373152](https://news.ycombinator.com/item?id=49373152)
- Source: [vulnbench.com](https://vulnbench.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T11:27:20Z

## Translation

タイトル: VulnBench: LLM は同じセキュリティ バグを 2 回見つけることができますか?
記事のタイトル: Snyk VulnBench
説明: AI システムが脆弱性をどのように確実に検出するかを調査し、すべての結論の背後にある証拠を検査します。

記事本文:
Snyk VulnBench コンテンツへスキップ VulnBench の概要
LLM は同じバグを 2 回見つけることができますか?
再現性と Snyk との参照一致に関する調査
検査可能な JavaScript プロジェクトに対して同じエージェント セキュリティ レビューを 5 回実行して、何が再発するか、何が変化するか、モデルの結果が決定論的な Snyk コード参照セットとどのように一致するかを測定しました。
見出しの証拠 · 5 件の同一レビュー
同じレビューです。異なる結果。
300 スキャン · 10 プロジェクト · 6 つの構成
158 件中 134 件の参照一致所見が 5 回の実行すべてで見られました。
161 件中 22 件 5 回の実行すべてで一致しない結果が見られた
161 件中 80 件 1 回の実行のみで一致しない所見が見つかった
1回の失点では決定ではありません。再現性は、検査対象の優先順位付けに役立ちます。一致しないレポートは誤検知としてラベル付けされません。
134/158 · 5 回の実行すべてで参照一致の所見が見られた
同じコードです。同じプロンプトです。異なる発見。
5 回の同一の実行で、参照と一致する所見の 84.8% が毎回再発しました。一致しない報告書のほぼ半数は 1 回しか公開されず、証拠は無視するものではなく検証する必要があります。
結果が実際に何を意味するか
AI レビューは測定であり、評決ではありません。
同じタスクを繰り返すと、ストーリーが鮮明になります。参考文献と一致する多くの調査結果は安定しています。 AI レビューと決定論的 SAST によりさまざまな盲点が明らかになります。そして、より多くの支出がSnykとの参照契約を確実に改善するわけではありません。
01 再現性により自信が見える化
158 件の参照一致所見のうち 134 件が、5 件の同一レビューのすべてに出現しました。繰り返しにより、報告されたパターンが 1 回の実行に依存せずに持続することがわかります。
慎重に解釈してください。リファレンス一致は、独立したグラウンドトゥルースの精度ではなく、Snyk コードとの一致を測定します。
02 AI レビューと SAST がさまざまな盲点を明らかにする
モデルは高信号エクスプロイトの形状を表面化し、一方で阻止

ミニスティックな Snyk コードは、繰り返しデータ フロー シンクを一貫して列挙しました。両者の異なる結果は有用な証拠であり、一方が普遍的な勝者であると宣言する理由にはなりません。
慎重に解釈してください: 一致しない所見は、分類する前に症例レベルの検査が必要です。
03 追加料金を支払っても確実に結果が改善されるわけではない
このベンチマークでは、セッション コストが高くても、Snyk 基準の F1 が常に高くなるわけではありません。単独で費やすことは、構成を選択する近道としては不十分です。
慎重に解釈してください。コストの見積もりには、テストされた小規模な設備と出版物の仮定が反映されています。
VulnBench が動作を測定する方法
条件を繰り返し、証拠を保存し、観察された合意をプロトコルがサポートできない主張から分離します。
10 個の小さな JavaScript および Express フィクスチャにより、すべての実行と参照結果がレビュー可能になります。
各構成では、同じコード、プロンプト、ハーネス、タスクが 5 回表示されます。
報告された問題は、再発分析に適した文書化された署名になります。
スコアラーは、脆弱性のタイプを決定論的な Snyk コードの検出結果と比較します。
合意、再発、差異、適用範囲、コスト、トークン、期間は明確に区別されます。
一致しないレポートは、自動的な誤検知ではなく、調査すべき証拠として残ります。
VulnBench はバージョン化された研究イニシアチブであり、ユニバーサル リーダーボードではありません。すべてのリリースは、何を測定し、何を証明しなかったかを定義します。
さらにリリースが予定されています。未公開の結果は、推測ランキングや空のリリース カードとして表示されません。
作品を読み、複製し、引用する
論文、レビューされた方法論、ソース スナップショット、およびリリース データでは、安定した公開リンクが使用されています。
@misc{tal2026snykvulnbenchjs10,
著者 = {タル、リランとクロース、ヨハネスとルーディッチ、アルセニーとトームズ、スティーブンとネール、マノージ}、
title = {Snyk VulnBench JS 1.0: LLM は同じバグを 2 回見つけることができますか?},
はい

r = {2026}、
URL = {https://arxiv.org/abs/2606.15762}
BibTeX をコピー Liran Tal · Johannes Kloos · Arsenii Rudich · Stephen Thoemmes · Manoj Nair
AI システムが脆弱性をいかに確実に発見するかについての、透明性のあるバージョン管理された証拠。
2026 年 6 月 11 日公開 · データセット 1.0.0
© 2026 Snyk Limited.研究成果物は Apache 2.0 で提供されます。

## Original Extract

Explore how reliably AI systems find vulnerabilities—and inspect the evidence behind every conclusion.

Snyk VulnBench Skip to content VulnBench Overview
Can LLMs find the same bugs twice?
A repeatability and Snyk-reference agreement study
We ran the same agentic security review five times against inspectable JavaScript projects to measure what recurs, what varies, and how model findings align with a deterministic Snyk Code reference set.
Headline evidence · 5 identical reviews
Same review. Different results.
300 scans · 10projects · 6 configurations
134 of 158 Reference-matched findings seen in all five runs Inspect
22 of 161 Unmatched findings seen in all five runs Inspect
80 of 161 Unmatched findings seen in only one run Inspect
One run is not a decision. Repeatability helps prioritize what to inspect; it does not label unmatched reports as false positives.
134 of 158 · Reference-matched findings seen in all five runs
Same code. Same prompt. Different findings.
Across five identical runs, 84.8% of reference-matched findings recurred every time. Nearly half of unmatched reports appeared only once—evidence to inspect, not dismiss.
What the results mean in practice
An AI review is a measurement—not a verdict.
Repeat the same task and the story sharpens: many reference-matched findings hold steady; AI review and deterministic SAST reveal different blind spots; and more spend does not reliably improve Snyk-reference agreement.
01 Repeatability makes confidence visible
134 of 158 reference-matched findings appeared in every one of five identical reviews. Repetition shows which reported patterns persist instead of relying on a single run.
Interpret with care: A reference match measures agreement with Snyk Code, not independent ground-truth accuracy.
02 AI review and SAST expose different blind spots
Models surfaced high-signal exploit shapes, while deterministic Snyk Code consistently enumerated repeated data-flow sinks. Their differing results are useful evidence—not a reason to declare one a universal winner.
Interpret with care: Unmatched findings require case-level inspection before they can be classified.
03 Paying more did not reliably improve the result
In this benchmark, higher session cost did not consistently yield higher Snyk-reference F1. Spend alone is a poor shortcut for choosing a configuration.
Interpret with care: Cost estimates reflect the tested small fixtures and publication assumptions.
How VulnBench measures behavior
Repeat the conditions, preserve the evidence, and separate observed agreement from claims the protocol cannot support.
Ten small JavaScript and Express fixtures make every run and reference finding reviewable.
Each configuration sees the same code, prompt, harness, and task five times.
Reported issues become documented signatures suitable for recurrence analysis.
The scorer compares vulnerability type against deterministic Snyk Code findings.
Agreement, recurrence, variance, coverage, cost, tokens, and duration stay distinct.
Unmatched reports remain evidence to investigate—not automatic false positives.
VulnBench is a versioned research initiative, not a universal leaderboard. Every release defines what it measured and what it did not prove.
Further releases are planned. Unpublished results will not appear as speculative rankings or empty release cards.
Read, reproduce, and cite the work
The paper, reviewed methodology, source snapshot, and release data use stable public links.
@misc{tal2026snykvulnbenchjs10,
author = {Tal, Liran and Kloos, Johannes and Rudich, Arsenii and Thoemmes, Stephen and Nair, Manoj},
title = {Snyk VulnBench JS 1.0: Can LLMs Find the Same Bugs Twice?},
year = {2026},
url = {https://arxiv.org/abs/2606.15762}
} Copy BibTeX Liran Tal · Johannes Kloos · Arsenii Rudich · Stephen Thoemmes · Manoj Nair
Transparent, versioned evidence about how reliably AI systems find vulnerabilities.
Published 11 June 2026 · Dataset 1.0.0
© 2026 Snyk Limited. Research artifacts are provided under Apache 2.0.

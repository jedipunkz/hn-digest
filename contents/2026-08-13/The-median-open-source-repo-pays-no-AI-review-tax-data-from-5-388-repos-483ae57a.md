---
source: "https://forge.ambera.app/review-tax-index"
hn_url: "https://news.ycombinator.com/item?id=49284039"
title: "The median open-source repo pays no \"AI review tax\" (data from 5,388 repos)"
article_title: "The Review Tax Index — Ambera Forge"
author: "makivlach"
captured_at: "2026-08-13T10:53:18Z"
capture_tool: "hn-digest"
hn_id: 49284039
score: 1
comments: 0
posted_at: "2026-08-13T10:44:44Z"
tags:
  - hacker-news
  - translated
---

# The median open-source repo pays no "AI review tax" (data from 5,388 repos)

- HN: [49284039](https://news.ycombinator.com/item?id=49284039)
- Source: [forge.ambera.app](https://forge.ambera.app/review-tax-index)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T10:44:44Z

## Translation

タイトル: オープンソース リポジトリの中央値は「AI レビュー税」を支払っていません (5,388 リポジトリからのデータ)
記事のタイトル: 税インデックスの見直し — アンベラ フォージ
説明: AI で作成されたコードのレビュー コストが実際にどこにかかるのかを、数千のパブリック リポジトリにわたって測定しました。属性に応じた作業のレビューにコストがかかるリポジトリと、コストが低いリポジトリに、言語とプロジェクトのサイズごとに分かれています。

記事本文:
The Review Tax Index — Ambera Forge スキャン キャパシティ レポート ベンチマーク手法 プライバシー The Index
AI で書かれたコードはどこでもレビューに多くのコストがかかるという主張を耳にするでしょう。データは、さらに有益なことを示しています。コストは分割されており、一部のリポジトリには負担が大きく、他のリポジトリにはまったく負担がかかりません。このページはどこを測定します。
検出可能なエージェントの作成者がまったくいない
帰属リポジトリの 20% 以上、マージされた作業の 20% 以上に帰属が含まれています
この測定は、検出可能な帰属と少なくとも 10 の分析済み PR を備えた 2,353 のリポジトリにわたって実行されます。レビューとコストの比率には、帰属が存在する側が必要です。それぞれの中で、帰属された作品は、同じリポジトリの残りの部分、つまり同じ査読者、同じ規約、同じ時代と比較して評価されます。 1x を超える倍数は、属性のある側がより多くの時間を費やしたことを意味します。以下は、時間がかからなかったことを意味します。
マージまでの時間、帰属/休憩
比較可能なリポジトリの 29% は、帰属された作業に対してより多くのレビューを支払っています。 71% が同じかそれ以下の金額を支払っています。人口の中央半分: 0.2 x ～ 1.2 x。
レビューラウンド、帰属/残り
比較可能なリポジトリの 32% は、帰属された作業に対してより多くのレビューを支払っています。 68% が同等かそれ以下の金額を支払っています。人口の中央半分: 0.2 x ～ 1.2 x。
属性付きリポジトリ全体で検出されたシェアの中央値は 3% です。これはシェアではなく下限です。属性はツールがコミットに残すマークからのみ得られるため、インライン完了作業は構造上見えません。全人口の中央値はゼロです。読み取られるリポジトリのほとんどには、帰属がまったく含まれていません。
ダッシュは、そのセグメントに公開するには比較可能なリポジトリが少なすぎることを意味します。少数の比率の中央値は、数字をまとったノイズです。
ダッシュは、そのセグメントに公開するには比較可能なリポジトリが少なすぎることを意味します。少数の比率の中央値は、数字をまとったノイズです。
帰属 vs 残り、決して AI vs 人間ではありません。の

比較グループには、痕跡を残さない AI 支援の作業が含まれており、ここに示されているすべてのギャップは広がるのではなく、むしろ狭まります。
サンプルは 2 回自己選択します。これらは、誰かがスキャンするために選択したリポジトリに加えて、有名なプロジェクトの厳選されたバックフィルです。ソフトウェアのランダムなサンプルではなく、人々が関心を持つリポジトリの母集団です。
保留ビートが発明されました。分析された PR が 10 未満のリポジトリはカウントされません。比較可能なリポジトリが 12 未満のセグメントでは中央値が公表されません。
インデックスは四半期ごとに再発行されます。すでに読み取られたリポジトリにはその履歴が保存されるため、将来のエディションには、税金がどこに到達するかだけでなく、どこに移動するかというトレンドラインが表示されます。
完全なメソッドとすべての制限: /methodology 。この集団に対して独自のリポジトリを配置します: /scan 。

## Original Extract

Where the review cost of AI-authored code actually lands, measured across thousands of public repositories: the split between repos where attributed work costs more review and repos where it costs less, by language and by project size.

The Review Tax Index — Ambera Forge Scan Capacity Reports Benchmark Methodology Privacy The Index
The claim you will hear is that AI-written code costs more review everywhere. The data says something more useful: the cost is a split — it lands hard on some repositories and not at all on others. This page measures where.
no detectable agent authorship at all
of attributed repos, 20%+ of merged work carries attribution
The measures run over the 2,353 repositories with detectable attribution and at least 10 analyzed PRs — a review-cost ratio needs an attributed side to exist. Within each, attributed work is measured against the rest of the same repository — same reviewers, same conventions, same era. A multiple above 1x means the attributed side took more; below means it took less.
Time to merge, attributed / rest
29% of comparable repos pay more review on attributed work; 71% pay the same or less. Middle half of the population: 0.2 x to 1.2 x.
Review rounds, attributed / rest
32% of comparable repos pay more review on attributed work; 68% pay the same or less. Middle half of the population: 0.2 x to 1.2 x.
Median detected share across the attributed repositories is 3% — and that is a floor, not a share: attribution comes only from marks a tool leaves on a commit, so inline-completion work is invisible by construction. Across the full population the median is zero — most repositories read carry no attribution at all.
A dash means the segment has too few comparable repos for that measure to publish — a median over a handful of ratios is noise wearing a number.
A dash means the segment has too few comparable repos for that measure to publish — a median over a handful of ratios is noise wearing a number.
Attributed vs rest, never AI vs human. The comparison group contains AI-assisted work that leaves no trace, which narrows every gap shown here rather than widening it.
The sample self-selects twice. These are repositories someone chose to scan plus a curated backfill of well-known projects — a population of repos people care about, not a random sample of software.
Withheld beats invented. Repos below 10 analyzed PRs do not count; segments below 12 comparable repos publish no median.
The Index republishes quarterly. Repos already read keep their history, so future editions carry trend lines — not just where the tax lands, but where it is moving.
Full method and every limitation: /methodology . Place your own repo against this population: /scan .

---
source: "https://yaroslavvb.github.io/artificial-analysis-oss-lag/"
hn_url: "https://news.ycombinator.com/item?id=48586938"
title: "How far behind is open-source AI?"
article_title: "How far behind is open-source AI?"
author: "eternal_braid"
captured_at: "2026-06-18T16:13:30Z"
capture_tool: "hn-digest"
hn_id: 48586938
score: 1
comments: 0
posted_at: "2026-06-18T15:30:53Z"
tags:
  - hacker-news
  - translated
---

# How far behind is open-source AI?

- HN: [48586938](https://news.ycombinator.com/item?id=48586938)
- Source: [yaroslavvb.github.io](https://yaroslavvb.github.io/artificial-analysis-oss-lag/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T15:30:53Z

## Translation

タイトル: オープンソース AI はどれくらい遅れているのか?
説明: オープンソース AI パレート フロンティアの経時変化: フロンティアを押し広げた各オープンウェイト モデルについて、独自モデルが同じ人工分析インテリジェンス インデックスに何か月早く到達したか。

記事本文:
オープンソース AI はどれくらい遅れているのでしょうか?
オープンソースのパレートフロンティアの経時変化: 各ポイントは、それ以前のすべてのオープン モデルを上回るオープンウェイト モデルです。プロプライエタリ モデルが何ヶ月前に同じレベルに達したかによってプロットされています。
垂直位置は、同じ人工解析を使用した最初期の独自モデルとのギャップです
インテリジェンス インデックス (任意の点にカーソルを置くと、どの点がどれかを確認できます)。このギャップは DeepSeek V3 の頃に 10 か月近くでピークに達しました
（2024 年 12 月）その後、DeepSeek、Kimi、Z AI、MiniMax の開始に伴い、約 2 ～ 3.5 か月に短縮されました
フロンティアのすぐ近くで発送します。
フロンティアプッシャー。モデルは、そのインテリジェンス指数がすべてのオープンウェイトを超えた場合にのみ含まれます。
その前に発売されたモデル。 23名が合格。 (2024 年初めに世間知らずのランニングマックスが警告した Solar Mini は除外されます:
フロアレベルのスコアでは、遡及的な v4.1 インデックスが同時期のより強力なオープン モデルよりも上位にランクされます — Mixtral 8×7B、
Llama-2-70B、Qwen-72B — したがって、実際にはリリース時点では主要なオープン モデルではありませんでした。)
メートル法。 Artificial Analysis Intelligence Index (v4.1)、すべてのモデルに一貫して適用
すべての日付。現在のハードベンチマーク (GPQA、HLE、ターミナルベンチなど) に基づいているため、2023 モデルのスコアは低く、
フロアでの注文がうるさい。
ギャップ。日付 D にインデックス S でリリースされたフロンティア モデルの場合、ギャップ =
D − T 、ここで T は独自のモデルがインデックスに到達した最も古い日付です
≥ S — その独自のモデルが各ポイントのツールチップに表示されます。
注記。これは、特定の能力レベルにおけるギャップであり、絶対的なフロンティアではありません。 GLM-5.2
(II 51) は 3.4 か月前の GPT-5.4 と一致しますが、絶対閉鎖フロンティア (Claude Fable 5、II 60)
本来の能力ではさらに上にあります。

## Original Extract

The open-source AI Pareto frontier over time: for each open-weights model that pushed the frontier, how many months earlier a proprietary model reached the same Artificial Analysis Intelligence Index.

How far behind is open-source AI?
The open-source Pareto frontier over time: each point is an open-weights model that beat every open model before it — plotted by how many months earlier a proprietary model reached the same level.
The vertical position is the gap to the earliest proprietary model with the same Artificial Analysis
Intelligence Index (hover any point to see which one). The gap peaked near 10 months around DeepSeek V3
(Dec 2024) and has since tightened to ~2–3.5 months as DeepSeek, Kimi, Z AI and MiniMax began
shipping close behind the frontier.
Frontier-pusher. A model is included only if its Intelligence Index exceeded every open-weights
model released before it. 23 qualify. (Solar Mini, which the naive running-max flags in early 2024, is excluded: at
floor-level scores the retroactive v4.1 index ranks it above contemporaneous stronger open models — Mixtral 8×7B,
Llama-2-70B, Qwen-72B — so it was not actually the leading open model at release.)
Metric. Artificial Analysis Intelligence Index (v4.1), applied consistently to every model across
all dates. It is anchored to today's hard benchmarks (GPQA, HLE, terminal-bench, etc.), so 2023 models score low and
their ordering at the floor is noisy.
Gap. For a frontier model released on date D with index S , gap =
D − T , where T is the earliest date any proprietary model reached index
≥ S — that proprietary model is shown in each point's tooltip.
Note. This is the gap at a given capability level , not the absolute frontier. GLM-5.2
(II 51) matches GPT-5.4 from 3.4 months earlier, while the absolute closed frontier (Claude Fable 5, II 60)
sits further ahead in raw capability.

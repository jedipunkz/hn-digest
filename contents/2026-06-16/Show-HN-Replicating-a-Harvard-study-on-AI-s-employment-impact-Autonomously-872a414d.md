---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48556246"
title: "Show HN: Replicating a Harvard study on AI's employment impact – Autonomously"
article_title: ""
author: "robeenly"
captured_at: "2026-06-16T16:38:10Z"
capture_tool: "hn-digest"
hn_id: 48556246
score: 3
comments: 0
posted_at: "2026-06-16T14:55:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Replicating a Harvard study on AI's employment impact – Autonomously

- HN: [48556246](https://news.ycombinator.com/item?id=48556246)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T14:55:18Z

## Translation

タイトル: Show HN: AI の雇用への影響に関するハーバード大学の研究を再現 – 自律的に
HN テキスト: NeuGBI を使用して、同じ Revelio Lab データセット (米国の 3 億の雇用記録) 上で「年功序列による技術変化としての生成 AI」(HBS、2025) を複製しました。論文の調査結果: AI は若手職 (-29.4%) と上級職 (-5.8%) に不均衡な影響を与えています。 NeuGBI も自律的に同じ結論に達しました。 NeuGBI が論文に記載されていないことを発見したのは、ソフトウェア開発の分野では、エントリーレベル (L1) ではなく、特にジュニアレベル (L2) のポジションがほぼ半分になったことです。 NeuGBI は、クエリ エンジンとして NeuG (マルチホップ リレーションシップ サポートを備えたグラフ データベース)、分析用のハイパーグラフ再構成、および質問を分解して段階的にドリルダウンするために LLM が呼び出せるパッケージ化された探索スキルを使用します。 NeuGBI の主な機能は、エンドツーエンドの不偏サンプリングです。3 億レコードでは、複雑なマルチホップ クエリが数時間ではなく数秒で返されます。ブログ投稿: https://graphscope.io/blog/tech/2026/06/16/NEUGBI-BLOG.html
元の論文: https://arxiv.org/abs/2603.10625

## Original Extract

We used NeuGBI to replicate "Generative AI as Seniority-Biased Technological Change" (HBS, 2025) on the same Revelio Lab dataset — 300M U.S. employment records. The paper's finding: AI disproportionately affects junior positions (−29.4%) vs. senior (−5.8%). NeuGBI arrived at the same conclusion autonomously. One thing NeuGBI found that the paper didn't: within software development, it's specifically junior-level (L2) positions that nearly halved, not entry-level (L1). NeuGBI uses NeuG (a graph database with multi-hop relationship support) as its query engine, Hypergraph reconstruction for analysis, and packaged exploratory Skills that an LLM can invoke to decompose questions and drill down step by step. The key capability of NeuGBI is end-to-end unbiased sampling — on 300M records, complex multi-hop queries return in seconds rather than hours. Blog post: https://graphscope.io/blog/tech/2026/06/16/NEUGBI-BLOG.html
Original paper: https://arxiv.org/abs/2603.10625


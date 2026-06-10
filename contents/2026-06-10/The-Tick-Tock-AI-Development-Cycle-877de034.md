---
source: "https://wilsoniumite.com/2026/06/10/the-tick-tock-ai-development-cycle/"
hn_url: "https://news.ycombinator.com/item?id=48475714"
title: "The Tick-Tock AI Development Cycle"
article_title: "The Tick-Tock AI Development Cycle. - Wilsons Blog"
author: "Wilsoniumite"
captured_at: "2026-06-10T13:17:35Z"
capture_tool: "hn-digest"
hn_id: 48475714
score: 1
comments: 0
posted_at: "2026-06-10T13:04:39Z"
tags:
  - hacker-news
  - translated
---

# The Tick-Tock AI Development Cycle

- HN: [48475714](https://news.ycombinator.com/item?id=48475714)
- Source: [wilsoniumite.com](https://wilsoniumite.com/2026/06/10/the-tick-tock-ai-development-cycle/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T13:04:39Z

## Translation

タイトル: Tick-Tock AI 開発サイクル
記事のタイトル: Tick-Tock AI 開発サイクル。 - ウィルソンズブログ
説明: 2007 年にインテルは、ティックトック プロセッサ開発サイクルとして知られるものを採用しました。ウィキペディアより: 「このモデルでは、すべての新しいプロセス技術が最初に実証済みのマイクロアーキテクチャ (tick) のダイシュリンクを製造するために使用され、その後、現在実証済みのプロセス (tock) で新しいマイクロアーキテクチャが製造されます。」 T
[切り捨てられた]

記事本文:
Tick-Tock AI 開発サイクル。 - ウィルソンズブログ
コンテンツにスキップ
ウィルソンズのブログ
Tick-Tock AI 開発サイクル。
2007 年にインテルは、ティックトック プロセッサー開発サイクルとして知られるサイクルを採用しました。ウィキペディアより: 「このモデルでは、すべての新しいプロセス技術が最初に実証済みのマイクロアーキテクチャ (tick) のダイシュリンクを製造するために使用され、その後、現在実証済みのプロセス (tock) で新しいマイクロアーキテクチャが製造されます。」これは経営陣が選択したというよりも、足並みをそろえて開発されるテクノロジーですでに起こっていることを形式化したものです。ダイの縮小が最初に行われ、何を扱うのかが分かれば、マイクロアーキテクチャを改善できます。
AI の開発も非常によく似たパターンに従っています。モデル自体が重要であり、ハーネス、フレームワーク、ループ、パイプライン、およびエージェント定義が重要です。前者なしでは後者に取り組むのは難しいため、すべての開発がこのパターン内で調整され始めています。
これが最終的に意味するのは、単にきちんとした観察であるということを超えて、これら 2 つの要素のうち 1 つだけを観察することはあまり役に立たないということです (新しいモデルが以前のモデルと目立った違いがないと主張するのと同様)。代わりに、ベンチマークでのパフォーマンスやプロンプトからの感触を超えて、モデルの実際の意味をタックが理解するまで待つ必要があります。
Tick-Tock AI 開発サイクル。
LLM を使用すると頭が悪くなるのでしょうか?
システムがすべて、ソフトウェアがシステム

## Original Extract

In 2007 Intel adopted what’s known as their tick-tock processor development cycle. From Wikipedia: “Under this model, every new process technology was first used to manufacture a die shrink of a proven microarchitecture (tick), followed by a new microarchitecture on the now-proven process (tock)”. T
[truncated]

The Tick-Tock AI Development Cycle. - Wilsons Blog
Skip to content
Wilsons Blog
The Tick-Tock AI Development Cycle.
In 2007 Intel adopted what’s known as their tick-tock processor development cycle. From Wikipedia: “Under this model, every new process technology was first used to manufacture a die shrink of a proven microarchitecture (tick), followed by a new microarchitecture on the now-proven process (tock)”. This isn’t so much a choice made by management as it is a formalisation of what already is going on with technologies that develop in lockstep. The die shrink comes first, and once you know what you are working with you can improve the microarchitecture.
AI development is following a very similar pattern. The models themselves are the tick, and the harnesses, frameworks, loops, pipelines, and agent definitions are the tock. It’s hard to work on the latter without the former, so all development is beginning to align within this pattern.
What that ends up meaning, beyond just being a neat observation, is that it’s not that useful to observe only one of these two components (like claiming a new model isn’t noticeably different from the last). Instead, you have to wait also for the tock to understand the real implications of a model beyond it’s performance on a benchmark and it’s feel from a prompt.
The Tick-Tock AI Development Cycle.
Does using LLMs make me dumber?
Systems are Everything, Software is Systems

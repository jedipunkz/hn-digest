---
source: "https://simonwillison.net/2026/Jun/14/why-ai-hasnt-replaced-software-engineers/"
hn_url: "https://news.ycombinator.com/item?id=48537572"
title: "Why AI hasn't replaced software engineers, and won't"
article_title: "Why AI hasn’t replaced software engineers, and won’t"
author: "SVI"
captured_at: "2026-06-15T08:10:12Z"
capture_tool: "hn-digest"
hn_id: 48537572
score: 4
comments: 0
posted_at: "2026-06-15T07:02:37Z"
tags:
  - hacker-news
  - translated
---

# Why AI hasn't replaced software engineers, and won't

- HN: [48537572](https://news.ycombinator.com/item?id=48537572)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/14/why-ai-hasnt-replaced-software-engineers/)
- Score: 4
- Comments: 0
- Posted: 2026-06-15T07:02:37Z

## Translation

タイトル: AI がソフトウェア エンジニアに取って代わらない理由、そして今後も代替しない理由
記事のタイトル: AI がソフトウェア エンジニアに取って代わらない理由、そして今後も代替しない理由
説明: アルビンド ナラヤナンとサヤシュ カッポールは、AI の破壊に特に適した専門職であるソフトウェア エンジニアリングのレンズを通して、AI による雇用喪失の問題に取り組みます。で …

記事本文:
AI がソフトウェア エンジニアに取って代わらない理由、そして今後も代替しない理由
サイモン・ウィリソンのウェブログ
AI がソフトウェア エンジニアに取って代わらない理由、そして今後も置き換えられない理由 。アルビンド・ナラヤナンとサヤシュ・カッポールは、AI の破壊に特に適した職業であるソフトウェア エンジニアリングのレンズを通して、AI による雇用喪失の問題に取り組みます。
このエッセイでは、AI の能力が一定の閾値に達すると大量解雇を引き起こすという説を否定する十分な証拠があると主張します。規制の壁がほとんどない分野であってもこれが当てはまることを考えると、他のほとんどの職業はさらに緩和される可能性があります。
最初の良いニュースは、AI が大量の失業を引き起こしているという考えをデータがまだ裏付けていないということです。
2025 年 3 月、ニューヨーク州は、WARN 法の提出書類に AI 開示チェックボックスを追加した最初の米国の州となりました。最初の 1 年間で、160 社以上の企業が警告通知を提出しました。 AI ボックスにチェックを入れた人は 1 人もいませんでした
AI はコンピューターにコードを入力する段階を高速化しますが、ソフトウェア エンジニアリングはそれ以上のものであることが判明しました。
コードを書くことがボトルネックではないとしたら、何がボトルネックなのでしょうか?タスクの内訳調査は、会議やデバッグなどを対象としています。これでは、開発者は会議で何をしているのか、なぜ AI では実行できないのか、というさらなる疑問が生じます。機能が向上するにつれてデバッグは自動化されないのでしょうか?本当のボトルネックを理解するには、定性的な情報を取得し、自動化を妨げているソフトウェア エンジニアの動作についてソフトウェア エンジニア自身の理解を掘り下げる必要があります。
この分析を行ったところ、実際のボトルネックとして 3 つのことが明らかになりました。(1) 何を構築するかの決定と指定、(2) 提供されるものの検証と責任、(3) コードベース、ビジネス、環境に対する人間の深い理解が必要であることです。

これらの両方を実行する必要があります。
AI による支援は、決定や検証のステップにも役立ちますが、私が提供する価値の鍵となるのは依然として「人間の深い理解」です。世界中の AI 支援をすべて私に提供してください。それでも私が生み出す価値は、問題とエージェントが構築している解決策の両方を私がどれだけ深く理解しているかに依存します。
Pyodide で使用するために WASM ホイールを PyPI に公開 - 2026 年 6 月 13 日
クロード・ファブルは容赦なく積極的です - 2026 年 6 月 11 日
クロード・ファブルの第一印象 5 - 2026 年 6 月 9 日
これは、2026 年 6 月 14 日に投稿された、Simon Willison によるリンク投稿です。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Arvind Narayanan and Sayash Kappor take on the question of AI job losses through the lens of a profession that is uniquely suited to AI disruption - software engineering. In …

Why AI hasn’t replaced software engineers, and won’t
Simon Willison’s Weblog
Why AI hasn’t replaced software engineers, and won’t . Arvind Narayanan and Sayash Kappor take on the question of AI job losses through the lens of a profession that is uniquely suited to AI disruption - software engineering.
In this essay, we argue that there is enough evidence to reject the narrative that once AI capabilities reach a certain threshold, it will cause mass layoffs. Given that this is true even in a sector with very few regulatory barriers, most other professions are likely to be even more cushioned.
The first good news is that the data still doesn't support the idea that AI is causing mass unemployment.
In March 2025, New York became the first U.S. state to add an AI disclosure checkbox to WARN Act filings. In the full first year, more than 160 companies filed WARN notices. Not a single one checked the AI box
AI speeds up the typing-code-into-a-computer phase, but it turns out software engineering is about a whole lot more than that:
If writing code isn’t the bottleneck, what is? The task-breakdown surveys point at things like meetings or debugging. This just leads to more questions: what are developers doing in those meetings and why can’t it be done by AI? Won’t debugging get automated as capabilities improve? To understand the real bottlenecks, we have to get qualitative, and dig into software engineers’ own understanding of what it is they do that resists automation.
When we did this analysis, it revealed three things as the real bottlenecks (1) deciding and specifying what to build, (2) verifying and being accountable for what is delivered, and (3) the deep human understanding — of the codebase, the business, and the environment — required to carry out both of these.
I'm finding AI assistance also helps me with the deciding and verifying steps, but it's the "deep human understanding" that remains key to the value I provide. Give me all of the AI assistance in the world and the value I produce will still be reliant on how deeply I understand both the problems and the solutions that the agents are building for them.
Publishing WASM wheels to PyPI for use with Pyodide - 13th June 2026
Claude Fable is relentlessly proactive - 11th June 2026
Initial impressions of Claude Fable 5 - 9th June 2026
This is a link post by Simon Willison, posted on 14th June 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.

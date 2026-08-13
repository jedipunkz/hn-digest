---
source: "https://arxiv.org/abs/2608.11694"
hn_url: "https://news.ycombinator.com/item?id=49282632"
title: "The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance"
article_title: "[2608.11694] The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance"
author: "sbulaev"
captured_at: "2026-08-13T07:12:43Z"
capture_tool: "hn-digest"
hn_id: 49282632
score: 1
comments: 0
posted_at: "2026-08-13T07:07:07Z"
tags:
  - hacker-news
  - translated
---

# The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance

- HN: [49282632](https://news.ycombinator.com/item?id=49282632)
- Source: [arxiv.org](https://arxiv.org/abs/2608.11694)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T07:07:07Z

## Translation

タイトル: 文言効果: LLM ベンチマーク パフォーマンスの双方向ドリフトの定量化
記事のタイトル: [2608.11694] 文言効果: LLM ベンチマーク パフォーマンスの双方向ドリフトの定量化
説明: arXiv 論文 2608.11694 の要約ページ: The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 8 月 12 日に提出]
タイトル: 文言効果: LLM ベンチマーク パフォーマンスの双方向ドリフトの定量化
要約: ベンチマーク スコアは、各問題の 1 つの表現から得られます。その 1 つのフレーズが、同じ問題を問うすべての方法を表しているかのように扱われますが、実際はそうではありません。私たちは、問題の意味と答えを固定したまま問題を言い換えると、モデルの答えが日常的に双方向に反転し、一部の失敗は成功になり、一部の成功は失敗になることを示します。これをドリフトと呼びます。 BenchDrift は、言語的、参照的、語用的、構造的な 4 つの軸に沿って、意味を保持したベンチマーク問題のバリエーションを生成し、それぞれの軸で正しさが反転する頻度と理由を測定します。 8 つのモデルと 3 つのベンチマーク (GSM8K、MMLU、MATH-Hard) にわたって、ドリフトが両方向に大きいことが観察されました。 2 つの発見が際立っています。まず、モデルが良くなってもフレージングの感性が衰えることはありません。代わりに、符号が変わります。弱いモデルは言い換えによって失うものよりも得るものが大きいのに対し、強いモデルは得るものよりも失うものの方がはるかに多いのです。したがって、ベンチマークで最も優れたモデルは、そのスコアがたまたま与えられた文言に最も大きく依存するモデルであることがわかります。第 2 に、モデルは、ドリフトの程度が異なるにもかかわらず、どの言い換えが最も正解のコストが高いかについてほぼ一致しているため、脆弱性はモデルではなく言い換えに属します。さらに、問題を短くしても長くしても、モデルが自信を持って答えた答えを言い換えることができます。コードとデータ: この https URL
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
彼を提出してください

トーリー
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.11694: The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 12 Aug 2026]
Title: The Wording Effect: Quantifying Two-Way Drift in LLM Benchmark Performance
Abstract: A benchmark score comes from a single phrasing of each problem. That single phrasing is treated as if it stood for the whole space of ways the same problem could be asked, but it does not. We show that rephrasing a problem while keeping its meaning and answer fixed routinely flips a model's answer in both directions, so some failures become successes and some successes become failures. We call this drift. BenchDrift generates meaning-preserving variations of benchmark problems along four axes, namely linguistic, referential, pragmatic, and structural, and measures how often, and why, correctness flips under each. Across eight models and three benchmarks (GSM8K, MMLU, MATH-Hard), we observe that drift is large in both directions. Two findings stand out. First, phrasing sensitivity does not fade as models get better. Instead, it changes sign. Weak models gain more from rephrasing than they lose, while strong models lose far more than they gain. We find that the best models on a benchmark are therefore the ones whose scores depend most on the wording they happened to be given. Second, the models largely agree on which rephrasings cost the most correct answers even though they differ in how much they drift, so fragility belongs to the rephrasing and not to the model. Furthermore, rephrasing breaks answers a model was confident about, whether the problem is made shorter or longer. Code and Data: this https URL
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

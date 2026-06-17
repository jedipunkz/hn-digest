---
source: "https://arxiv.org/abs/2606.18141"
hn_url: "https://news.ycombinator.com/item?id=48569306"
title: "Behind Python: The Languages That Power AI"
article_title: "[2606.18141] Behind Python: The Languages That Power AI"
author: "runningmike"
captured_at: "2026-06-17T13:23:51Z"
capture_tool: "hn-digest"
hn_id: 48569306
score: 1
comments: 1
posted_at: "2026-06-17T12:14:50Z"
tags:
  - hacker-news
  - translated
---

# Behind Python: The Languages That Power AI

- HN: [48569306](https://news.ycombinator.com/item?id=48569306)
- Source: [arxiv.org](https://arxiv.org/abs/2606.18141)
- Score: 1
- Comments: 1
- Posted: 2026-06-17T12:14:50Z

## Translation

タイトル: Python の背後にある: AI を動かす言語
記事のタイトル: [2606.18141] Python の背後にある: AI を強化する言語
説明: arXiv 論文 2606.18141 の要約ページ: Behind Python: The Languages That Power AI

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.18141
ヘルプ |高度な検索
コンピュータサイエンス > プログラミング言語
[2026 年 6 月 16 日に提出]
タイトル: Python の背後にある: AI を動かす言語
要約: Python は AI 開発の主流を占めていますが、PyTorch や NumPy などのフレームワークの背後にある数値処理は C、C++、または Rust で実行されます。開発者がそのようなライブラリを使用せずにアルゴリズムを実装する必要がある場合、ライブラリが存在しない、ターゲットがリソースに制約がある、または新しいシステムを構築しているため、どの言語を選択する必要があるでしょうか。この論文はその質問に実証的に答えます。データ マイニング (k 平均法)、機械学習 (k-NN)、ニューラル ネットワーク (バックプロパゲーションを備えた MLP)、計算知能 (遺伝的アルゴリズム)、およびファジー システム (マムダニ推論) をカバーする 5 つのアルゴリズムが、Python、C、C++、Rust、Go、および Julia で最初から実装されています。すべての実装は共通の疑似乱数ジェネレーターを共有し、同一の入力を消費し、ビット同一の出力を生成するため、測定されたすべての違いは計算ではなく言語を反映します。 3 つのパフォーマンス層が現れます。C と C++ は事実上連携しています。 Rust はそれらに 9% (幾何平均) 遅れています。 Julia は C や Go の 5.0 倍よりも 3.3 倍遅く実行されます。 Python は 315x にあります。メモリの場合は別の話になります。Julia の JIT ランタイムはワークロードに関係なく最大 224 MiB の固定フットプリントを持ちますが、C、C++、Rust は 6 MiB 未満に留まります。重要なのは、ランキングが安定していないということです。Go の速度低下は、k-NN の 2.6 倍から K-means の 8.0 倍まで変動しており、ワークロードの特性によって言語の位置が 1 層変化する可能性があることを示しています。結果は、実装を選択するためのワークロードごとの具体的なガイダンスを提供します。

AI システムにおける言語。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.18141: Behind Python: The Languages That Power AI

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.18141
Help | Advanced Search
Computer Science > Programming Languages
[Submitted on 16 Jun 2026]
Title: Behind Python: The Languages That Power AI
Abstract: Python dominates AI development, yet the numerical work behind frameworks like PyTorch and NumPy is executed in C, C++, or Rust. When a developer must implement an algorithm without such libraries -- because none exists, the target is resource-constrained, or a new system is being built -- which language should they choose? This paper answers that question empirically. Five algorithms covering data mining (k-means), machine learning (k-NN), neural networks (MLP with backpropagation), computational intelligence (genetic algorithm), and fuzzy systems (Mamdani inference) are implemented from scratch in Python, C, C++, Rust, Go, and Julia. All implementations share a common pseudo-random generator, consume identical inputs, and produce bit-identical outputs, so every measured difference reflects the language rather than the computation. Three performance tiers emerge: C and C++ are effectively tied; Rust trails them by 9% (geometric mean); Julia runs 3.3x slower than C and Go 5.0x; Python sits at 315x. Memory tells a different story -- Julia's JIT runtime carries a fixed ~224 MiB footprint regardless of workload, while C, C++, and Rust stay below 6 MiB. Crucially, rankings are not stable: Go's slowdown swings from 2.6x on k-NN to 8.0x on k-means, showing that workload characteristics can shift a language's position by a full tier. The results provide concrete, per-workload guidance for choosing an implementation language in AI systems.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe

---
source: "https://arxiv.org/abs/2606.17930"
hn_url: "https://news.ycombinator.com/item?id=48649767"
title: "Inference Compute Shapes Frontier LLM Evaluation"
article_title: "[2606.17930] How Inference Compute Shapes Frontier LLM Evaluation"
author: "matt_d"
captured_at: "2026-06-23T19:35:58Z"
capture_tool: "hn-digest"
hn_id: 48649767
score: 1
comments: 0
posted_at: "2026-06-23T19:02:55Z"
tags:
  - hacker-news
  - translated
---

# Inference Compute Shapes Frontier LLM Evaluation

- HN: [48649767](https://news.ycombinator.com/item?id=48649767)
- Source: [arxiv.org](https://arxiv.org/abs/2606.17930)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T19:02:55Z

## Translation

タイトル: 推論計算形状フロンティア LLM 評価
記事のタイトル: [2606.17930] 推論計算がフロンティア LLM 評価をどのように形成するか
説明: arXiv 論文 2606.17930 の要約ページ: 推論によるフロンティア LLM 評価の計算方法

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.17930
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 16 日に提出]
タイトル: 推論計算がフロンティア LLM 評価をどのように形成するか
要約: AI の評価は、ツールの使用と反復的な問題解決を伴う長期にわたる軌道から恩恵を受ける、より困難なタスクへと移行しています。その結果、パフォーマンスは、テスト時に利用可能なコンピューティング (「推論コンピューティング」) の量と割り当てにますます敏感になります。しかし、依然として多くの評価では単一の制限された予算でのパフォーマンスが報告されており、低いスコアはモデルの基礎的な機能ではなく評価設定を反映している可能性があることを意味します。これをテストするために、ソフトウェア エンジニアリング、数学、医学、サイバーセキュリティにわたる 7 つの挑戦的なベンチマークで最大 12 のフロンティア言語モデルを評価します。私たちは、3 つの単純な推論スケーリング介入を組み合わせた制御されたセットアップを使用します。つまり、より大きなトークン バジェット、コンテキストの圧縮、モデル自体または最小限の正確性フィードバックのいずれかによって導かれる、送信試行の繰り返しです。主な結果は 3 つあります。まず、より大きなトークン バジェットにより、サイバーセキュリティ、FrontierMath、人類最後の試験、および TerminalBench を含む複数のドメインにわたるベンチマークのパフォーマンスが大幅に向上します。第二に、固定予算の評価では、モデルが進歩するにつれてフロンティアの能力がますます過小評価される可能性があります。新しいモデルは、大きな予算でより高いパフォーマンスを実現し、より困難なタスクを解放し、より確実に解決します。第三に、どの推論スケーリング手法が最も役立つかがベンチマークによって異なります。繰り返し送信するとパフォーマンスが大幅に向上しますが、

より大きなトークン バジェット、外部フィードバック、並列試行はベンチマークによって異なります。全体として、私たちの結果は、ベンチマーク スコアがプロトコルに依存していることを示しています。したがって、評価では、特に安全性またはポリシー関連の設定において、推論時間のコンピューティングの関数として機能を報告し、プロトコルの選択を明示的に指定し、一致した予算で大規模な共有コンピューティング範囲にわたってモデルの世代を比較する必要があると主張します。
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

Abstract page for arXiv paper 2606.17930: How Inference Compute Shapes Frontier LLM Evaluation

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.17930
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 16 Jun 2026]
Title: How Inference Compute Shapes Frontier LLM Evaluation
Abstract: AI evaluations are shifting toward harder tasks that benefit from longer trajectories involving tool use and iterative problem solving. As a result, performance is increasingly sensitive to the amount and allocation of compute available at test time ("inference compute"). Yet many evaluations still report performance at a single restrictive budget, meaning that low scores may reflect the evaluation setup rather than the model's underlying capability. To test this, we evaluate up to 12 frontier language models on seven challenging benchmarks spanning software engineering, mathematics, medicine, and cybersecurity. We use a controlled setup combining three simple inference-scaling interventions: larger token budgets, context compaction, and repeated submission attempts, guided either by the model itself or by minimal correctness feedback. We find three main results. First, larger token budgets substantially improve performance on benchmarks across multiple domains, including cybersecurity, FrontierMath, Humanity's Last Exam, and TerminalBench. Second, fixed-budget evaluations can increasingly understate frontier capability as models advance. Newer models reach higher performance at large budgets, where they unlock harder tasks and solve them more reliably. Third, benchmarks differ in which inference-scaling methods help most: repeated submission broadly improves performance, but the value of larger token budgets, external feedback, and parallel attempts varies by benchmark. Overall, our results show that benchmark scores are protocol-dependent. We therefore argue that evaluations should report capability as a function of inference-time compute, specify protocol choices explicitly, and compare model generations over a large shared compute range at matched budgets, especially in safety- or policy-relevant settings.
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

---
source: "https://arxiv.org/abs/2606.04023"
hn_url: "https://news.ycombinator.com/item?id=48468715"
title: "CodegenBench: Can LLMs Write Efficient Code Across Architectures?"
article_title: "[2606.04023] CodegenBench: Can LLMs Write Efficient Code Across Architectures?"
author: "matt_d"
captured_at: "2026-06-10T01:00:48Z"
capture_tool: "hn-digest"
hn_id: 48468715
score: 2
comments: 0
posted_at: "2026-06-09T22:31:46Z"
tags:
  - hacker-news
  - translated
---

# CodegenBench: Can LLMs Write Efficient Code Across Architectures?

- HN: [48468715](https://news.ycombinator.com/item?id=48468715)
- Source: [arxiv.org](https://arxiv.org/abs/2606.04023)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T22:31:46Z

## Translation

タイトル: CodegenBench: LLM はアーキテクチャ全体で効率的なコードを書くことができますか?
記事のタイトル: [2606.04023] CodegenBench: LLM はアーキテクチャ全体で効率的なコードを書くことができますか?
説明: arXiv 論文 2606.04023 の要約ページ: CodegenBench: LLM はアーキテクチャ全体で効率的なコードを書くことができますか?

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.04023
ヘルプ |高度な検索
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 6 月 1 日に提出]
タイトル: CodegenBench: LLM はアーキテクチャ全体で効率的なコードを書くことができますか?
要約: 大規模言語モデル (LLM) は、汎用プログラミングおよび GPU アクセラレーション環境 (PyTorch、CUDA など) のコード生成タスクで広範囲に評価されてきましたが、多様なアーキテクチャにわたる CPU 指向のハイパフォーマンス コンピューティング (HPC) におけるその機能は、依然として十分に解明されていません。このギャップを埋めるために、x86_64、Sunway、Kunpeng の 3 つの異なるハードウェア プラットフォームにわたる効率的な並列コードの生成を評価するように設計された包括的なベンチマーク スイートである CodegenBench を紹介します。私たちのベンチマークは、基本的なベースラインを確立する 106 個の標準基本線形代数サブプログラム (BLAS) ルーチンと、独自のスーパーコンピューティング アーキテクチャ (LeetSunway および LeetKunpeng) のそれぞれに適合した 20 個の特殊な計算カーネルで構成されています。私たちの広範な評価により、最先端の LLM は x86_64 のようなユビキタス アーキテクチャ向けに最適化されたコードを生成できる一方で、公開ドキュメントやトレーニング データが限られたドメイン固有のアーキテクチャでは大幅なパフォーマンスの低下を示し、クロスプラットフォームの一般化における重大な制限が浮き彫りになったことが明らかになりました。さらに、実装の長さやタスクの複雑さなど、コードの品質に影響を与える要因を分析したところ、現在の LLM は、簡潔なコード スニペットを必要とする中程度に難しい問題に対して最も効果的であることが示されています。データセットと自動評価インフラストラクチャをオープンソース化しています

LLM 主導の高性能コード生成に関する将来の研究を促進するため。リソースは、この https URL およびこの https URL で入手できます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
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

Abstract page for arXiv paper 2606.04023: CodegenBench: Can LLMs Write Efficient Code Across Architectures?

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.04023
Help | Advanced Search
Computer Science > Software Engineering
[Submitted on 1 Jun 2026]
Title: CodegenBench: Can LLMs Write Efficient Code Across Architectures?
Abstract: While large language models (LLMs) have been extensively evaluated on code generation tasks for general-purpose programming and GPU-accelerated environments (e.g., PyTorch, CUDA), their capabilities in CPU-oriented high-performance computing (HPC) across diverse architectures remain underexplored. To bridge this gap, we introduce CodegenBench, a comprehensive benchmark suite designed to evaluate the generation of efficient parallel code across three distinct hardware platforms: x86_64, Sunway, and Kunpeng. Our benchmark comprises 106 standard Basic Linear Algebra Subprograms (BLAS) routines establishing a fundamental baseline, alongside 20 specialized computational kernels adapted for each of the unique supercomputing architectures (LeetSunway and LeetKunpeng). Our extensive evaluation reveals that while state-of-the-art LLMs can generate optimized code for ubiquitous architectures like x86_64, they exhibit significant performance degradation on domain-specific architectures with limited public documentation and training data, highlighting critical limitations in cross-platform generalization. Furthermore, our analysis of factors influencing code quality such as implementation length and task complexity indicates that current LLMs are most effective for moderately difficult problems requiring concise code snippets. We open-source our dataset and automated evaluation infrastructure to facilitate future research in LLM-driven high-performance code generation. The resources are available at this https URL and this https URL .
Focus to learn more
arXiv-issued DOI via DataCite
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

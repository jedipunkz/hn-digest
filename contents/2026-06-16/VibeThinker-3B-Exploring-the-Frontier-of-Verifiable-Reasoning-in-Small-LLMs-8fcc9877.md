---
source: "https://arxiv.org/abs/2606.16140"
hn_url: "https://news.ycombinator.com/item?id=48554065"
title: "VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small LLMs"
article_title: "[2606.16140] VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small Language Models"
author: "Anon84"
captured_at: "2026-06-16T13:57:45Z"
capture_tool: "hn-digest"
hn_id: 48554065
score: 2
comments: 0
posted_at: "2026-06-16T12:18:13Z"
tags:
  - hacker-news
  - translated
---

# VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small LLMs

- HN: [48554065](https://news.ycombinator.com/item?id=48554065)
- Source: [arxiv.org](https://arxiv.org/abs/2606.16140)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T12:18:13Z

## Translation

タイトル: VibeThinker-3B: 小規模 LLM における検証可能な推論のフロンティアを探索する
記事のタイトル: [2606.16140] VibeThinker-3B: 小さな言語モデルにおける検証可能な推論のフロンティアを探索する
説明: arXiv 論文 2606.16140 の要約ページ: VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small Language Models

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.16140
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 15 日に提出]
タイトル: VibeThinker-3B: 小さな言語モデルにおける検証可能な推論のフロンティアを探索する
要約: この技術レポートでは、VibeThinker-3B を紹介します。VibeThinker-3B は、厳密に小規模なモデルの領域内で検証可能な推論をどこまで推し進めることができるかを調査するために開発された、3B パラメーターを備えたコンパクトな高密度モデルです。スペクトルから信号へのポストトレーニング パラダイムに基づいて、カリキュラムベースの教師あり微調整、マルチドメイン強化学習、オフライン自己蒸留を含む最適化されたパイプラインを通じてモデルを体系的に強化します。実験による評価では、VibeThinker-3B が要求の厳しい検証可能なタスクにおいてフロンティア レベルのパフォーマンスを達成することが実証されています。具体的には、AIME26 で 94.3 のスコア (クレーム レベルのテスト時間スケーリングにより 97.1 に改善)、LiveCodeBench v6 で 80.2 Pass@1 を達成し、最近の未見の LeetCode コンテストでは 96.1\% の受け入れ率で強力な配布外一般化を示しています。これにより、事実上、DeepSeek V3.2、GLM-5、Gemini 3 Pro など、桁違いに大きいフラッグシップ モデルと同等またはそれを上回る、第 1 層推論システムのパフォーマンス帯域に位置付けられます。さらに、IFEval のスコア 93.4 は、この極端な推論の強化によって厳密な命令の制御性が損なわれないことを裏付けています。以前の 1.5B の研究を拡張して、これらの発見は、検証可能な推論をコンパクトな推論コアに圧縮できるとみなすパラメトリック圧縮カバレッジ仮説の動機付けとなります。

一方、オープンドメインの知識と汎用能力には、事実、概念、ロングテール シナリオにわたる幅広いパラメータをカバーする必要があります。この観点からは、コンパクト モデルは単に展開効率の高い代替品ではなく、パラメータが密な機能領域におけるフロンティア レベルのパフォーマンスに向けた補完的な道であることが示唆されます。
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

Abstract page for arXiv paper 2606.16140: VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small Language Models

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.16140
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 15 Jun 2026]
Title: VibeThinker-3B: Exploring the Frontier of Verifiable Reasoning in Small Language Models
Abstract: This technical report introduces VibeThinker-3B, a compact dense model with 3B parameters developed to investigate how far verifiable reasoning can be pushed within a strictly small-model regime. Building upon the Spectrum-to-Signal post-training paradigm, we systematically enhance the model through an optimized pipeline that includes curriculum-based supervised fine-tuning, multi-domain reinforcement learning, and offline self-distillation. Experimental evaluations demonstrate that VibeThinker-3B achieves frontier-level performance on highly demanding verifiable tasks. Specifically, it attains a score of 94.3 on AIME26 (improving to 97.1 with claim-level test-time scaling), an 80.2 Pass@1 on LiveCodeBench v6, and exhibits strong out-of-distribution generalization with a 96.1\% acceptance rate on recent unseen LeetCode contests. This effectively places it in the performance band of first-tier reasoning systems, matching or exceeding flagship models that are orders of magnitude larger, such as DeepSeek V3.2, GLM-5, and Gemini 3 Pro. Furthermore, a score of 93.4 on IFEval confirms that this extreme reasoning enhancement does not compromise strict instruction controllability. Extending our previous 1.5B work, these findings motivate the Parametric Compression-Coverage Hypothesis, which views verifiable reasoning as compressible into compact reasoning cores, while open-domain knowledge and general-purpose competence require broad parameter coverage over facts, concepts, and long-tail scenarios. This perspective suggests that compact models are not merely deployment-efficient substitutes, but a complementary path toward frontier-level performance in parameter-dense capability regimes.
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

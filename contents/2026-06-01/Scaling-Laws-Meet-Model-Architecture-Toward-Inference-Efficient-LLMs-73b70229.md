---
source: "https://arxiv.org/abs/2510.18245"
hn_url: "https://news.ycombinator.com/item?id=48342826"
title: "Scaling Laws Meet Model Architecture: Toward Inference-Efficient LLMs"
article_title: "[2510.18245] Scaling Laws Meet Model Architecture: Toward Inference-Efficient LLMs"
author: "matt_d"
captured_at: "2026-06-01T00:35:25Z"
capture_tool: "hn-digest"
hn_id: 48342826
score: 2
comments: 0
posted_at: "2026-05-31T03:41:57Z"
tags:
  - hacker-news
  - translated
---

# Scaling Laws Meet Model Architecture: Toward Inference-Efficient LLMs

- HN: [48342826](https://news.ycombinator.com/item?id=48342826)
- Source: [arxiv.org](https://arxiv.org/abs/2510.18245)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T03:41:57Z

## Translation

タイトル: スケーリング則とモデル アーキテクチャの融合: 推論効率の高い LLM に向けて
記事のタイトル: [2510.18245] スケーリング則とモデル アーキテクチャの出会い: 推論効率的な LLM に向けて
説明: arXiv 論文 2510.18245 の要約ページ: スケーリング則とモデル アーキテクチャの出会い: 推論効率の高い LLM に向けて

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2510.18245
ヘルプ |高度な検索
コンピューターサイエンス > 機械学習
[2025 年 10 月 21 日に提出 ( v1 )、最終改訂日 2026 年 5 月 13 日 (このバージョン、v3)]
タイトル: スケーリング則とモデル アーキテクチャの融合: 推論効率の高い LLM に向けて
要約: パラメータの数とトレーニング データのサイズをスケーリングすることは、大規模言語モデル (LLM) のパフォーマンスを向上させる効果的な戦略であることが証明されています。しかし、これらのモデルがますます強力になり、広く導入されるようになるにつれて、推論のコストが差し迫った懸念事項になっています。その重要性にもかかわらず、モデルの精度と推論効率の間のトレードオフは依然として十分に解明されていません。この研究では、主要なアーキテクチャ要素、隠れたサイズ、MLP とアテンションの間のパラメーターの割り当て (MLP 対アテンションの比率)、およびグループ化クエリ アテンション (GQA) が、推論コストと精度の両方にどのように影響するかを調べます。チンチラ フレームワークをアーキテクチャ情報で強化する条件付きスケーリング則と、推論効率と精度を同時に実現するアーキテクチャを特定するための検索フレームワークを導入します。私たちのアプローチを検証するために、80M ～ 3B のパラメーターと 8B ～ 100B のトレーニング トークンにわたる 200 以上のモデルをトレーニングし、提案された条件付きスケーリング則に適合させます。私たちの結果は、条件付きスケーリング則が最適なアーキテクチャの選択を確実に予測し、結果として得られるモデルが既存のオープンソース ベースラインを上回るパフォーマンスを示すことを示しています。同じトレーニング予算の下で、最適化されたアーキテクチャは、LLaMA-3.2 と比較して最大 2.1% 高い精度と 42% 高い推論スループットを達成します。
もっと学ぶために集中する

arXiv が DataCite 経由で発行した DOI
参考雑誌:
ICLR 2026
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

Abstract page for arXiv paper 2510.18245: Scaling Laws Meet Model Architecture: Toward Inference-Efficient LLMs

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2510.18245
Help | Advanced Search
Computer Science > Machine Learning
[Submitted on 21 Oct 2025 ( v1 ), last revised 13 May 2026 (this version, v3)]
Title: Scaling Laws Meet Model Architecture: Toward Inference-Efficient LLMs
Abstract: Scaling the number of parameters and the size of training data has proven to be an effective strategy for improving large language model (LLM) performance. Yet, as these models grow increasingly powerful and widely deployed, the cost of inference has become a pressing concern. Despite its importance, the trade-off between model accuracy and inference efficiency remains underexplored. In this work, we examine how key architectural factors, hidden size, the allocation of parameters between MLP and attention (mlp-to-attention ratio), and grouped-query attention (GQA), influence both inference cost and accuracy. We introduce a conditional scaling law that augments the Chinchilla framework with architectural information, along with a search framework for identifying architectures that are simultaneously inference-efficient and accurate. To validate our approach, we train more than 200 models spanning 80M to 3B parameters and 8B to 100B training tokens, and fit the proposed conditional scaling law. Our results show that the conditional scaling law reliably predicts optimal architectural choices and that the resulting models outperform existing open-source baselines. Under the same training budget, optimized architectures achieve up to 2.1% higher accuracy and 42% greater inference throughput compared to LLaMA-3.2.
Focus to learn more
arXiv-issued DOI via DataCite
Journal reference:
ICLR 2026
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

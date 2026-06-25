---
source: "https://arxiv.org/abs/2606.15868"
hn_url: "https://news.ycombinator.com/item?id=48677547"
title: "David vs. Goliath in Next Activity Prediction: Argmax vs. LSTM, Transformer, LLM"
article_title: "[2606.15868] David vs. Goliath in Next Activity Prediction: Argmax vs. LSTM, Transformer, and LLM"
author: "hramezani"
captured_at: "2026-06-25T18:42:19Z"
capture_tool: "hn-digest"
hn_id: 48677547
score: 1
comments: 0
posted_at: "2026-06-25T18:40:47Z"
tags:
  - hacker-news
  - translated
---

# David vs. Goliath in Next Activity Prediction: Argmax vs. LSTM, Transformer, LLM

- HN: [48677547](https://news.ycombinator.com/item?id=48677547)
- Source: [arxiv.org](https://arxiv.org/abs/2606.15868)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T18:40:47Z

## Translation

タイトル: 次のアクティビティ予測における David vs. Goliath: Argmax vs. LSTM、Transformer、LLM
記事のタイトル: [2606.15868] 次のアクティビティ予測における David vs. Goliath: Argmax vs. LSTM、Transformer、および LLM
説明: arXiv 論文 2606.15868 の要約ページ: 次のアクティビティ予測における David vs. Goliath: Argmax vs. LSTM、Transformer、および LLM

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.15868
ヘルプ |高度な検索
コンピューターサイエンス > 機械学習
[2026 年 6 月 14 日に提出]
タイトル: 次のアクティビティ予測における David vs. Goliath: Argmax vs. LSTM、Transformer、および LLM
要約: 次のアクティビティ予測 (NAP) は、予測プロセス監視 (PPM) の基礎であり、組織が遡及的な分析からプロアクティブなプロセス ステアリングに移行できるようにします。 PPM 分野は、古典的な機械学習から、LSTM やトランスフォーマーなどの深層学習アーキテクチャを経て、大規模言語モデル (LLM) へと進歩してきました。モデルの複雑さが増大しているにもかかわらず、NAP のダイレクト シーケンス モデリング設定で LLM、トランスフォーマー、LSTM、および単純なベースラインを統合して比較するベンチマークはありません。このペーパーでは、体系的なベンチマークを使用してこのギャップを埋めます。語彙に適応した LLM、ゼロからトレーニングされたトランスフォーマー、LLM で抽出されたトランスフォーマー、および LSTM を、7 つの実際のイベント ログにわたる単純なカウントベースの argmax ベースラインと比較します。私たちの結果は、デビッド対ゴリアテの物語を物語っています。事前トレーニングでは、最初からトレーニングする場合と比較して一貫した改善は得られず、モデルのサイズはパフォーマンスにほとんど影響を示さず、ほとんどのデータセットでは、argmax ベースラインは 10 億パラメータの LLM のパフォーマンスに一致またはそれに近づきます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
どちらも個性的

arXivLabs と協力する関係者や組織は、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.15868: David vs. Goliath in Next Activity Prediction: Argmax vs. LSTM, Transformer, and LLM

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.15868
Help | Advanced Search
Computer Science > Machine Learning
[Submitted on 14 Jun 2026]
Title: David vs. Goliath in Next Activity Prediction: Argmax vs. LSTM, Transformer, and LLM
Abstract: Next activity prediction (NAP) is a cornerstone of predictive process monitoring (PPM), enabling organizations to move from retrospective analysis to proactive process steering. The PPM field has progressed from classical machine learning through deep learning architectures such as LSTMs and Transformers to large language models (LLMs). Despite growing model complexity, no benchmark jointly compares LLMs, Transformers, LSTMs, and simple baselines in a direct sequence modeling setting for NAP. In this paper, we fill this gap with a systematic benchmark. We compare vocabulary-adapted LLMs, Transformers trained from scratch, LLM-distilled Transformers, and LSTMs against a simple counting-based argmax baseline across seven real-life event logs. Our results tell a David vs. Goliath story: pretraining confers no consistent improvement over training from scratch, model size shows little effect on performance, and on most datasets the argmax baseline matches or approaches the performance of billion-parameter LLMs.
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

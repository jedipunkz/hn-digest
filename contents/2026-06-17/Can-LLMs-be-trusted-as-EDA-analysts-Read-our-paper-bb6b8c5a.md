---
source: "https://arxiv.org/abs/2606.00051"
hn_url: "https://news.ycombinator.com/item?id=48570792"
title: "Can LLMs be trusted as EDA analysts? Read our paper"
article_title: "[2606.00051] Business Utility of Large Language Models as Exploratory Data Analysis Agents"
author: "Applied_AI"
captured_at: "2026-06-17T16:23:18Z"
capture_tool: "hn-digest"
hn_id: 48570792
score: 2
comments: 0
posted_at: "2026-06-17T14:08:10Z"
tags:
  - hacker-news
  - translated
---

# Can LLMs be trusted as EDA analysts? Read our paper

- HN: [48570792](https://news.ycombinator.com/item?id=48570792)
- Source: [arxiv.org](https://arxiv.org/abs/2606.00051)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T14:08:10Z

## Translation

タイトル: LLM は EDA アナリストとして信頼できますか?私たちの論文を読む
記事のタイトル: [2606.00051] 探索的データ分析エージェントとしての大規模言語モデルのビジネス ユーティリティ
説明: arXiv 論文 2606.00051: 探索的データ分析エージェントとしての大規模言語モデルのビジネス ユーティリティの要約ページ

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.00051
ヘルプ |高度な検索
コンピュータサイエンス > コンピュータと社会
[2026 年 5 月 8 日提出]
タイトル: 探索的データ分析エージェントとしての大規模言語モデルのビジネス上の有用性
要約: 大規模言語モデル (LLM) は分析ワークフローでますます使用されていますが、ビジネス環境における探索的データ分析 (EDA) エージェントとしての適合性は依然として不確実です。実際には、展開可能な EDA エージェントは、有用な平均パフォーマンスを提供するだけでなく、出力の信頼性をサポートする十分な再現性も提供する必要があります。私たちは、エージェントベースのサプライチェーンシミュレーションに基づいて構築された、管理されたビジネス関連のベンチマークでこの要件を評価します。課題は、明示的なラベルではなく間接的な運用追跡から推論して、低品質と下流の販売損失の原因となっているサプライヤーと製品の組み合わせを特定することです。 8 つのモデル ファミリからの 15 のモデル バリアント構成が、データ表現、プロンプトの明瞭さ、および信号強度を変化させる 4 つの実験条件下で、条件ごとに 5 つの軌跡で評価されました。出力は、Jaccard インデックスを使用して決定論的なグラウンド トゥルースに対してスコア付けされ、平均スコア (ms)、変動係数 (CV)、探索的クロス条件有意性テスト、および品質と再現性を 1 つの運用尺度に要約するために当社が提案するリスク調整済み指標であるビジネス ユーティリティを組み合わせたフレームワークを通じて評価されました。結果は、平均スコアが許容できるように見えても、ほとんどの構成は自律型 EDA の使用には十分な信頼性がないことを示しています。非常に高い理由を備えた GPT-5.4

ng の取り組みは、実験平均 ms が 0.8748、ビジネス ユーティリティが 0.6952 という最も強力な全体プロファイルを達成しましたが、次に最適な構成では、変動性を割り引いた後、大幅に多くのユーティリティを失いました。私たちの調査結果は、EDA エージェントの評価は、平均的な品質、再現性、および状態の敏感度を、運用の信頼性の補完的な側面として扱う必要があることを示唆しています。
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

Abstract page for arXiv paper 2606.00051: Business Utility of Large Language Models as Exploratory Data Analysis Agents

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.00051
Help | Advanced Search
Computer Science > Computers and Society
[Submitted on 8 May 2026]
Title: Business Utility of Large Language Models as Exploratory Data Analysis Agents
Abstract: Large Language Models (LLMs) are increasingly used in analytical workflows, but their suitability as exploratory data analysis (EDA) agents in business settings remains uncertain. In practice, a deployable EDA agent must provide not only useful average performance but also sufficient repeatability to support trust in its outputs. We evaluate this requirement in a controlled, business-relevant benchmark built on an agent-based supply chain simulation. The task is to identify supplier-product combinations responsible for low quality and downstream sales loss by reasoning from indirect operational traces rather than from explicit labels. Fifteen model-variant configurations from eight model families were evaluated under four experimental conditions that varied data representation, prompt clarity, and signal strength, with five trajectories per condition. Outputs were scored against deterministic ground truth using the Jaccard index and assessed through a framework that combines mean score (ms), coefficient of variation (CV), exploratory cross-condition significance tests, and Business utility, a risk-adjusted metric that we propose to summarise quality and repeatability in a single operational measure. The results show that most configurations are not reliable enough for autonomous EDA use, even when their average scores appear acceptable. GPT-5.4 with extra-high reasoning effort achieved the strongest overall profile, with an experiment-averaged ms of 0.8748 and an experiment-averaged Business utility of 0.6952, while the next-best configurations lost substantially more utility after variability discounting. Our findings suggest that evaluation of EDA agents should treat average quality, repeatability, and condition sensitivity as complementary dimensions of operational trustworthiness.
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

---
source: "https://arxiv.org/abs/2606.27288"
hn_url: "https://news.ycombinator.com/item?id=48688730"
title: "Combining LLMs Rarely Beats the Best Single Model, I tested 67 frontier models"
article_title: "[2606.27288] When Does Combining Language Models Help? A Co-Failure Ceiling on Routing, Voting, and Mixture-of-Agents Across 67 Frontier Models"
author: "josefchen"
captured_at: "2026-06-26T17:33:33Z"
capture_tool: "hn-digest"
hn_id: 48688730
score: 1
comments: 0
posted_at: "2026-06-26T16:42:31Z"
tags:
  - hacker-news
  - translated
---

# Combining LLMs Rarely Beats the Best Single Model, I tested 67 frontier models

- HN: [48688730](https://news.ycombinator.com/item?id=48688730)
- Source: [arxiv.org](https://arxiv.org/abs/2606.27288)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T16:42:31Z

## Translation

タイトル: LLM の組み合わせが最高の単一モデルを上回ることはほとんどありません。67 のフロンティア モデルをテストしました
記事のタイトル: [2606.27288] 言語モデルの結合が役立つのはどのような場合ですか? 67 のフロンティア モデルにわたるルーティング、投票、およびエージェントの混合における同時障害の上限
説明: arXiv 論文 2606.27288 の要約ページ: 言語モデルの結合はどのような場合に役立ちますか? 67 のフロンティア モデルにわたるルーティング、投票、およびエージェントの混合における同時障害の上限

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.27288
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 25 日に提出]
タイトル: 言語モデルの結合が役立つのはどのような場合ですか? 67 のフロンティア モデルにわたるルーティング、投票、およびエージェントの混合における同時障害の上限
要約: ルーティング、投票、カスケード、フュージョン、エージェントの混合などのマルチモデル LLM システムは、単一モデルの精度を上回るために使用されます。私たちは、彼らの利益が、現場でほとんど報告されない量によって制限されていることを示します。出力が 1 つのメンバー モデルの回答であるポリシーの場合、精度は 1 マイナス ベータを超えることはできません。ベータとは、同じクエリに対してすべてのモデルが誤る率です。対照的に、通常の診断である平均ペアワイズ誤差相関ρはベータを識別できません。同一の周辺値とペアワイズ相関を持つ誤差則は、全誤り率が異なる可能性があります。ベータ版の Clopper-Pearson バウンドは、ルーターをトレーニングする前に、ルーター、投票、またはカスケードが提供できる最大のゲインに基づいて有限サンプル証明書を提供します。
21 社のプロバイダーの 67 モデルにわたって、テトラコールで校正された単一因子モデルは依然として間違った尾部の価格を下回っています。オープンエンド数学では、観測されたベータ値は 0.052 であるのに対し、完全な 67 モデルのガウス コピュラでは 0.023 であり、約 2.5 倍の割安であり、90 パーセント CI は 1.7 ～ 3.4、k は 17 に相当します。効果は再発します。実行グレード コードの場合、ベータは 0.079 です。同じ GPQA-Diamond の質問を多肢選択形式ではなく自由回答形式で再質問すると、ベータ 0.127、カッパ 0.73 ～ 0.92 の 5 人の裁判官パネルで尾部が再び開き、主題ではなく回答形式での共失敗箇所が特定されます。相応の品質で

y、低 rho の異種アンサンブルは高 rho の Self-MoA を上回りますが、プール内のチェック可能なタスクでは、強力なクエリ レベルのルーティング シグナルがなければ、モデルを組み合わせた方が単一の最良のモデルを上回ることはほとんどありません。利益は、モデルを追加することでではなく、さまざまな質問でモデルが失敗することから得られます。
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

Abstract page for arXiv paper 2606.27288: When Does Combining Language Models Help? A Co-Failure Ceiling on Routing, Voting, and Mixture-of-Agents Across 67 Frontier Models

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.27288
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 25 Jun 2026]
Title: When Does Combining Language Models Help? A Co-Failure Ceiling on Routing, Voting, and Mixture-of-Agents Across 67 Frontier Models
Abstract: Multi-model LLM systems such as routing, voting, cascades, fusion, and mixture-of-agents are used to beat single-model accuracy. We show that their gain is capped by a quantity the field rarely reports. For any policy whose output is one member model answer, accuracy cannot exceed one minus beta, where beta is the rate at which every model is wrong on the same query. In contrast, the usual diagnostic, average pairwise error correlation rho, cannot identify beta: error laws with identical marginals and pairwise correlations can have different all-wrong rates. A Clopper-Pearson bound on beta gives a finite-sample certificate on the largest gain any router, vote, or cascade could deliver before training a router.
Across 67 models from 21 providers, a tetrachoric-calibrated single-factor model still underprices the all-wrong tail: on open-ended mathematics, observed beta is 0.052 versus 0.023 under the full 67-model Gaussian copula, about 2.5 times underpricing, with 90 percent CI 1.7 to 3.4 and k equals 17. The effect recurs on execution-graded code, where beta is 0.079. Re-asking the same GPQA-Diamond questions in free-response rather than multiple-choice form reopens the tail, with beta 0.127 and a five-judge panel with kappa 0.73 to 0.92, locating co-failure in answer format rather than subject. At matched quality, low-rho heterogeneous ensembles beat high-rho Self-MoA, but on checkable tasks in our pool, combining models rarely beats the single best model without a strong query-level routing signal. Gains come from models failing on different questions, not from adding more models.
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

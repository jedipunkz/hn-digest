---
source: "https://arxiv.org/abs/2606.13241"
hn_url: "https://news.ycombinator.com/item?id=48552834"
title: "Brick: SOTA LLM Routing"
article_title: "[2606.13241] Brick: Spatial Capability Routing for the Mixture-of-Models (MoM) Paradigm"
author: "FrancescoMassa"
captured_at: "2026-06-16T10:41:03Z"
capture_tool: "hn-digest"
hn_id: 48552834
score: 2
comments: 0
posted_at: "2026-06-16T09:45:22Z"
tags:
  - hacker-news
  - translated
---

# Brick: SOTA LLM Routing

- HN: [48552834](https://news.ycombinator.com/item?id=48552834)
- Source: [arxiv.org](https://arxiv.org/abs/2606.13241)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T09:45:22Z

## Translation

タイトル: ブリック: SOTA LLM ルーティング
記事のタイトル: [2606.13241] ブリック: モデル混合 (MoM) パラダイムのための空間機能ルーティング
説明: arXiv 論文 2606.13241 の要約ページ: ブリック: 混合モデル (MoM) パラダイムのための空間機能ルーティング

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.13241
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 11 日に提出]
タイトル: ブリック: モデル混合 (MoM) パラダイムのための空間機能ルーティング
要約: クエリの難易度を定義することは、展開エンジニアリングにおいて最も難しい問題の 1 つです。既存の LLM ルーターは、ドメイン ラベル、キーワード、トークン数などの表面的な機能に依存し、モデルの成功を実際に決定するドメイン内の差異を無視しています。フロンティア モデルのコストは、ローカルのオープンウェイト モデルの 10 ～ 100 倍であるため、実稼働規模では、リクエストごとのわずかな節約であっても、クラウドの請求額に直接影響します。我々は、6 つの機能次元で各モデルにスコアを付け、これをクエリごとの難易度推定と組み合わせ、コストにペナルティを課した幾何学的ルールを介してディスパッチするマルチモーダル ルーターである Brick を紹介します。連続設定ノブを使用すると、オペレーターは展開時に最大品質プロファイルと最大節約プロファイルの間をスライドできます。 5,504 クエリのベンチマークでは、最高品質の Brick は 76.98% の精度に達し、最高の単一モデル (75.02%) およびテストされたすべてのルーターを上回りました。 Brick は、中立的なコスト品質プロファイルで、常に最も強力なモデルを使用する場合に比べて 4.71 倍の低コストで 74.11% の精度を達成します。最小コストでは、コストが 22.15 倍削減され、精度は 11.85 ポイント低下します。レイテンシの中央値は 51.2 秒から 22.8 秒に減少します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs はコラボレーションを可能にするフレームワークです

または、新しい arXiv 機能を開発し、Web サイトで直接共有することもできます。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.13241: Brick: Spatial Capability Routing for the Mixture-of-Models (MoM) Paradigm

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.13241
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 11 Jun 2026]
Title: Brick: Spatial Capability Routing for the Mixture-of-Models (MoM) Paradigm
Abstract: Defining query difficulty is one of the hardest problems in deployment engineering. Existing LLM routers rely on surface features such as domain labels, keywords, and token count, ignoring the within-domain variance that actually determines model success. Frontier models cost ten to one hundred times more than local open-weight models, so at production scale even small per-request savings become a direct cloud-bill lever. We present Brick, a multimodal router that scores each model on six capability dimensions, combines this with a per-query difficulty estimate, and dispatches via a cost-penalized geometric rule. A continuous preference knob lets operators slide between max-quality and max-saving profiles at deploy time. On a benchmark of 5,504 queries, Brick at max-quality reaches 76.98% accuracy, beating the best single model (75.02%) and all tested routers. At a neutral cost-quality profile, Brick achieves 74.11% accuracy at 4.71x lower cost than always using the strongest model. At min-cost, it cuts cost 22.15x with 11.85 points accuracy loss. Median latency drops from 51.2s to 22.8s.
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

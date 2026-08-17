---
source: "https://arxiv.org/abs/2608.03676"
hn_url: "https://news.ycombinator.com/item?id=49330030"
title: "Solving Moe Load Imbalance in LLM Training via Optimal Transport"
article_title: "[2608.03676] TAOT: Topology-Aware Optimal Transport for Dynamic Expert Replica Placement in MoE Training"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "nullnonenilNULL"
captured_at: "2026-08-17T13:33:17Z"
capture_tool: "hn-digest"
hn_id: 49330030
score: 1
comments: 0
posted_at: "2026-08-17T12:51:59Z"
tags:
  - hacker-news
  - translated
---

# Solving Moe Load Imbalance in LLM Training via Optimal Transport

- HN: [49330030](https://news.ycombinator.com/item?id=49330030)
- Source: [arxiv.org](https://arxiv.org/abs/2608.03676)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T12:51:59Z

## Translation

タイトル: 最適なトランスポートによる LLM トレーニングにおける Moe 負荷の不均衡の解決
記事のタイトル: [2608.03676] TAOT: MoE トレーニングにおける動的エキスパート レプリカ配置のためのトポロジを意識した最適なトランスポート
説明: arXiv 論文 2608.03676 の要約ページ: TAOT: MoE トレーニングにおける動的エキスパート レプリカ配置のためのトポロジを意識した最適なトランスポート

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 8 月 4 日に提出]
タイトル: TAOT: MoE トレーニングにおける動的エキスパート レプリカ配置のためのトポロジを意識した最適なトランスポート
要約: Mixture-of-Experts (MoE) は、大規模言語モデル (LLM) をスケーリングするための主要なアーキテクチャとなっていますが、その動的なルーティングにより、エキスパートと並列のトレーニングで深刻な負荷の不均衡が生じます。既存の動的レプリカ手法は、計算を共有するためにホット エキスパートをアイドル ランクにコピーしますが、ロード バランスのみを最適化し、マルチノード トポロジ全体でエキスパートの重みを移動するコストを無視するため、結果として生じるクロスノード通信がバランシング ゲインを上回り、トレーニング コストが膨らむ可能性があります。動的エキスパート レプリカ配置のためのトポロジを意識した最適なトランスポート方法である TAOT を紹介します。 TAOT は、ホット ランクの過負荷と軽負荷ランクの予備容量を、通信コスト マトリックスを使用したバランスのとれたエントロピー正規化された最適トランスポート問題としてモデル化し、シンクホーン-ノップ反復でそれを解決してランクレベルのフロー ヒントを生成し、整数レプリカ マッチングとトークン割り当てを実行可能なスケジュールに組み合わせます。システム レベルでは、ゲストの重み転送をホーム エキスパートの計算と重ね合わせて、通信オーバーヘッドを隠します。実験の結果、TAOT は 1.43 倍のエンドツーエンド MoE トレーニング速度向上を達成し、既存の最先端の方法と同等以上のバランス品質に達し、すべての構成にわたって加重専門家コミュニケーション コストが最小となり、最大 74% 削減されることが示されています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: e

コミュニティの協力者との実験プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.03676: TAOT: Topology-Aware Optimal Transport for Dynamic Expert Replica Placement in MoE Training

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 4 Aug 2026]
Title: TAOT: Topology-Aware Optimal Transport for Dynamic Expert Replica Placement in MoE Training
Abstract: Mixture-of-Experts (MoE) has become a key architecture for scaling large language models (LLMs), yet its dynamic routing causes severe load imbalance in expert-parallel training. Existing dynamic-replica methods copy hot experts onto idle ranks to share computation, but they optimize load balance alone and ignore the cost of moving expert weights across a multi-node topology, so the resulting cross-node communication can outweigh the balancing gain and inflate training cost. We present TAOT, a topology-aware optimal transport method for dynamic expert-replica placement. TAOT models the overload on hot ranks and the spare capacity on lightly loaded ranks as a balanced entropy-regularized optimal transport problem with a communication-cost matrix, solves it with Sinkhorn-Knopp iterations to produce rank-level flow hints, and combines integer replica matching with token assignment into an executable schedule. At the system level, it overlaps guest-weight transfer with home-expert computation to hide the communication overhead. Experiments show TAOT achieves a 1.43x end-to-end MoE training speedup, reaches balance quality competitive with or better than existing state-of-the-art methods, and attains the lowest weighted expert-communication cost across all configurations, with up to a 74% reduction.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

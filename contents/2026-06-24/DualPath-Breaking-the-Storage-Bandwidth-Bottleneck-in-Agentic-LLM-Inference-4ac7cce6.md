---
source: "https://arxiv.org/abs/2602.21548"
hn_url: "https://news.ycombinator.com/item?id=48664361"
title: "DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference"
article_title: "[2602.21548] DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference"
author: "yogthos"
captured_at: "2026-06-24T19:31:59Z"
capture_tool: "hn-digest"
hn_id: 48664361
score: 1
comments: 0
posted_at: "2026-06-24T19:07:04Z"
tags:
  - hacker-news
  - translated
---

# DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference

- HN: [48664361](https://news.ycombinator.com/item?id=48664361)
- Source: [arxiv.org](https://arxiv.org/abs/2602.21548)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T19:07:04Z

## Translation

タイトル: DualPath: エージェント LLM 推論におけるストレージ帯域幅のボトルネックを解消する
記事のタイトル: [2602.21548] DualPath: エージェントの LLM 推論におけるストレージ帯域幅のボトルネックを解消する
説明: arXiv 論文 2602.21548 の要約ページ: DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2602.21548
ヘルプ |高度な検索
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 2 月 25 日に提出 ( v1 )、最終改訂日 2026 年 2 月 26 日 (このバージョン、v2)]
タイトル: DualPath: エージェント LLM 推論におけるストレージ帯域幅のボトルネックを解消する
要約: マルチターンのエージェント LLM 推論のパフォーマンスは、計算よりも KV キャッシュ ストレージ I/O によってますます支配されています。一般的な分散アーキテクチャでは、外部ストレージから大規模な KV キャッシュをロードすると、根本的な不均衡が生じます。プレフィル エンジン上のストレージ NIC は帯域幅が飽和状態になり、デコード エンジン上のストレージ NIC はアイドル状態のままになります。この非対称性により、システム全体のスループットが大幅に制限されます。
ここでは、デュアルパス KV キャッシュ読み込みを導入することでこのボトルネックを解消する推論システム、DualPath を紹介します。従来のストレージからプレフィルへのパスを超えて、DualPath により、ストレージからデコードへの新しいパスが可能になります。このパスでは、KV キャッシュがデコード エンジンにロードされ、コンピューティング ネットワーク経由で RDMA 経由でプレフィル エンジンに効率的に転送されます。 DualPath は、こ​​の最適化されたデータ パス (本質的にネットワークの輻輳を回避し、遅延が重要なモデル実行通信への干渉を回避します) を、プレフィル エンジンとデコード エンジン全体で動的に負荷のバランスをとるグローバル スケジューラと組み合わせます。
実稼働エージェントのワークロードを備えた 3 つのモデルでの評価では、DualPath が社内の推論システムでオフライン推論のスループットを最大 1.87$\times$ 向上させることが実証されました。また、バイオラチンを使用せずに、オンライン サービスのスループットを平均 1.96$\times$ 向上させることもできます。

g SLO。
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

Abstract page for arXiv paper 2602.21548: DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2602.21548
Help | Advanced Search
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 25 Feb 2026 ( v1 ), last revised 26 Feb 2026 (this version, v2)]
Title: DualPath: Breaking the Storage Bandwidth Bottleneck in Agentic LLM Inference
Abstract: The performance of multi-turn, agentic LLM inference is increasingly dominated by KV-Cache storage I/O rather than computation. In prevalent disaggregated architectures, loading the massive KV-Cache from external storage creates a fundamental imbalance: storage NICs on prefill engines become bandwidth-saturated, while those on decoding engines remain idle. This asymmetry severely constrains overall system throughput.
We present DualPath, an inference system that breaks this bottleneck by introducing dual-path KV-Cache loading. Beyond the traditional storage-to-prefill path, DualPath enables a novel storage-to-decode path, in which the KV-Cache is loaded into decoding engines and then efficiently transferred to prefill engines via RDMA over the compute network. DualPath combines this optimized data path -- which inherently avoids network congestion and avoids interference with latency-critical model execution communications -- with a global scheduler that dynamically balances load across prefill and decode engines.
Our evaluation on three models with production agentic workloads demonstrates that DualPath improves offline inference throughput by up to 1.87$\times$ on our in-house inference system. It can also improve online serving throughput by an average factor of 1.96$\times$ without violating SLO.
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

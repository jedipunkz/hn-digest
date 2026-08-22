---
source: "https://arxiv.org/abs/2511.07885"
hn_url: "https://news.ycombinator.com/item?id=49396954"
title: "Intelligence per Watt: Measuring Intelligence Efficiency of Local AI"
article_title: "[2511.07885] Intelligence per Watt: Measuring Intelligence Efficiency of Local AI"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "helsinkiandrew"
captured_at: "2026-08-22T06:21:44Z"
capture_tool: "hn-digest"
hn_id: 49396954
score: 1
comments: 0
posted_at: "2026-08-22T05:54:49Z"
tags:
  - hacker-news
  - translated
---

# Intelligence per Watt: Measuring Intelligence Efficiency of Local AI

- HN: [49396954](https://news.ycombinator.com/item?id=49396954)
- Source: [arxiv.org](https://arxiv.org/abs/2511.07885)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T05:54:49Z

## Translation

タイトル: ワットあたりのインテリジェンス: ローカル AI のインテリジェンス効率の測定
記事のタイトル: [2511.07885] ワットあたりのインテリジェンス: ローカル AI のインテリジェンス効率の測定
説明: arXiv 論文 2511.07885 の要約ページ: Intelligence per Watt: Measuring Intelligence Efficiency of Local AI

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
[2025 年 11 月 11 日に提出 ( v1 )、最終改訂日 2026 年 8 月 7 日 (このバージョン、v5)]
タイトル: ワットあたりのインテリジェンス: ローカル AI のインテリジェンス効率の測定
要約: 大規模言語モデル (LLM) クエリは、主に集中型クラウド インフラストラクチャのフロンティア モデルによって処理されます。需要の増加により、プロバイダーが拡張できるよりも早くこのパラダイムに負担がかかります。 2 つの進歩により、それを再考する機会が生まれました。小型のローカル LM (アクティブ パラメータ 20B 以下) が、多くのタスクでフロンティア モデルに匹敵するパフォーマンスを達成できるようになりました。また、ローカル アクセラレータ (Apple M4 Max など) は、インタラクティブなレイテンシでこれらのモデルをホストできるようになりました。これにより、ローカル推論は集中型インフラストラクチャからの需要を実行可能に再分配できるのかという疑問が生じます。これには、ローカル LM が実際のクエリに正確に応答できるかどうかと、電力に制約のあるデバイス (ラップトップなど) で効率的に応答できるかどうかの両方を測定する必要があります。私たちは、モデル アクセラレータ構成全体にわたるローカル推論の機能と効率の統一指標として、ワットあたりのインテリジェンス (IPW)、電力単位あたりのタスク精度を提案します。私たちは、20 を超える最先端のローカル LM、8 つのハードウェア アクセラレータ (ローカルおよびクラウド)、および 100 万の現実世界のシングル ターン チャットと推論クエリを評価します。各クエリについて、精度 (フロンティア モデルに対するローカル LM 勝率)、エネルギー、レイテンシ、電力を測定します。 3 つの重要な結果が見つかりました。まず、ローカル LM はこれらのクエリの 88.7% に正常に応答しますが、精度はドメインによって異なります。第 2 に、2023 年から 2025 年の長期的な分析では、アルゴリズムとアクセラレータの両方の進歩により、IPW が 5.3 倍向上し、ローカルでサービス可能なクエリ カバレッジが増加したことが示されています。

ngは23.2%から71.3%。第三に、ローカル アクセラレータは、同一のモデルを実行するクラウド アクセラレータよりも少なくとも 1.4 倍低い IPW を達成しており、ローカル アクセラレータの最適化に大きな余裕があることがわかります。これらの調査結果は、ローカル推論が、クエリの実質的なサブセットに対する集中型インフラストラクチャからの需要を有意義に再分散できることを示しており、IPW はこの移行を追跡するための重要な指標として機能します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2511.07885: Intelligence per Watt: Measuring Intelligence Efficiency of Local AI

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 11 Nov 2025 ( v1 ), last revised 7 Aug 2026 (this version, v5)]
Title: Intelligence per Watt: Measuring Intelligence Efficiency of Local AI
Abstract: Large language model (LLM) queries are predominantly processed by frontier models in centralized cloud infrastructure. Demand growth strains this paradigm faster than providers can scale. Two advances create an opportunity to rethink it: small, local LMs (<=20B active parameters) now achieve competitive performance to frontier models on many tasks, and local accelerators (e.g., Apple M4 Max) can host these models at interactive latencies. This raises the question: can local inference viably redistribute demand from centralized infrastructure? This requires measuring both whether local LMs can accurately answer real-world queries and whether they can do so efficiently on power-constrained devices (e.g., laptops). We propose intelligence per watt (IPW), task accuracy per unit of power, as a unified metric for the capability and efficiency of local inference across model-accelerator configurations. We evaluate 20+ state-of-the-art local LMs, 8 hardware accelerators (local and cloud), and 1M real-world single-turn chat and reasoning queries. For each query, we measure accuracy (local LM win rate against frontier models), energy, latency, and power. We find three key results. First, local LMs successfully answer 88.7% of these queries, with accuracy varying by domain. Second, longitudinal analysis from 2023-2025 shows IPW improved 5.3x, driven by both algorithmic and accelerator advances, with locally-serviceable query coverage rising from 23.2% to 71.3%. Third, local accelerators achieve at least 1.4x lower IPW than cloud accelerators running identical models, revealing significant headroom for local accelerator optimization. These findings demonstrate that local inference can meaningfully redistribute demand from centralized infrastructure for a substantial subset of queries, with IPW serving as the critical metric for tracking this transition.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

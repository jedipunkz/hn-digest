---
source: "https://arxiv.org/abs/2608.13573"
hn_url: "https://news.ycombinator.com/item?id=49399974"
title: "A Year in LLM Serving: Workload Evolution, Caching and Load-Balancing"
article_title: "[2608.13573] A Year in LLM Serving: Workload Evolution, Caching and Load-Balancing"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "1a1a11a"
captured_at: "2026-08-22T15:11:46Z"
capture_tool: "hn-digest"
hn_id: 49399974
score: 1
comments: 0
posted_at: "2026-08-22T14:16:26Z"
tags:
  - hacker-news
  - translated
---

# A Year in LLM Serving: Workload Evolution, Caching and Load-Balancing

- HN: [49399974](https://news.ycombinator.com/item?id=49399974)
- Source: [arxiv.org](https://arxiv.org/abs/2608.13573)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T14:16:26Z

## Translation

タイトル: LLM サービスの 1 年間: ワークロードの進化、キャッシュ、ロード バランシング
記事のタイトル: [2608.13573] LLM サービスの 1 年間: ワークロードの進化、キャッシュ、ロード バランシング
説明: arXiv 論文 2608.13573 の要約ページ: A Year in LLM Serving: Workload Evolution、Caching and Load-Balancing

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 7 月 3 日に提出]
タイトル: LLM サービスの 1 年間: ワークロードの進化、キャッシュ、ロード バランシング
要約: 大規模言語モデル (LLM) のサービス提供は重要なクラウド ワークロードになっており、現実的なトレースはサービス提供システムのモチベーションを高め、ベンチマークを行うために不可欠です。ただし、既存の LLM サービスのワークロード調査は、規模と範囲が限られています。多くの場合、短期間を観察し、ユーザーが運用環境でモデルをどのように操作するかについて限定的な可視性を提供します。その結果、LLM サービス提供ワークロードが時間の経過とともにどのように進化するか、またはユーザーとモデルの相互作用が運用トラフィックをどのように形成するかを完全には把握できません。
この研究では、全体的な特性評価と、シュートからの 1 年間の運用トレースの長期的研究の両方を通じて、実際の LLM サービスのワークロードについての理解を深めます。以前の調査とは異なり、私たちのトレースは、人気モデルとロングテールモデルの両方を含む、多くのモデルとユーザーにわたる完全な本番環境の動作を捕捉します。集約、時間、モデルレベル、ユーザーレベルの観点からワークロードを分析し、通常は集約ビューの背後に隠れているワークロードの進化とユーザーモデルの構造を明らかにします。将来の研究をサポートするために、論文とともに 1 年間の完全なトレースを公開します。これにより、サンプリングされたワークロードや合成的に生成されたワークロードに依存せずに、本番環境の動作の下流の研究が可能になります。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が arXiv の新しい機能を開発および共有できるようにするフレームワークです。

私たちのウェブサイトに直接アクセスできます。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.13573: A Year in LLM Serving: Workload Evolution, Caching and Load-Balancing

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 3 Jul 2026]
Title: A Year in LLM Serving: Workload Evolution, Caching and Load-Balancing
Abstract: Large Language Model (LLM) serving has become a critical cloud workload, and realistic traces are essential for motivating and benchmarking serving systems. However, existing LLM serving workload studies remain limited in scale and scope. They often observe short time periods and provide limited visibility into how users interact with models in production. As a result, they do not fully capture how LLM serving workloads evolve over time or how user-model interactions shape production traffic.
In this work, we further the understanding of real-world LLM serving workloads through both a global characterization and a longitudinal study of a one-year production trace from Chutes. Unlike prior studies, our trace captures full production behavior across many models and users, including both popular and long-tail models. We analyze the workload from aggregate, temporal, model-level, and user-level perspectives, revealing workload evolution and user-model structure that are typically hidden behind aggregate views. To support future research, we will release the full one-year trace with the paper, enabling downstream studies of production behavior without relying on sampled or synthetically generated workloads.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

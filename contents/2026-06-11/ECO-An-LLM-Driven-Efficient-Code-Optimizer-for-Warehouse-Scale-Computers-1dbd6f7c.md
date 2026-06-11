---
source: "https://arxiv.org/abs/2503.15669"
hn_url: "https://news.ycombinator.com/item?id=48485743"
title: "ECO: An LLM-Driven Efficient Code Optimizer for Warehouse Scale Computers"
article_title: "[2503.15669] ECO: An LLM-Driven Efficient Code Optimizer for Warehouse Scale Computers"
author: "bone_tag"
captured_at: "2026-06-11T04:35:37Z"
capture_tool: "hn-digest"
hn_id: 48485743
score: 2
comments: 0
posted_at: "2026-06-11T02:58:32Z"
tags:
  - hacker-news
  - translated
---

# ECO: An LLM-Driven Efficient Code Optimizer for Warehouse Scale Computers

- HN: [48485743](https://news.ycombinator.com/item?id=48485743)
- Source: [arxiv.org](https://arxiv.org/abs/2503.15669)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T02:58:32Z

## Translation

タイトル: ECO: 倉庫規模のコンピューター向けの LLM 主導の効率的なコード オプティマイザー
記事のタイトル: [2503.15669] ECO: 倉庫規模のコンピューター向けの LLM 駆動の効率的なコード オプティマイザー
説明: arXiv 論文 2503.15669 の要約ページ: ECO: 倉庫規模のコンピューター用の LLM 駆動の効率的なコード オプティマイザー

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2503.15669
ヘルプ |高度な検索
コンピュータサイエンス > ソフトウェアエンジニアリング
[2025 年 3 月 19 日に提出]
タイトル: ECO: 倉庫規模のコンピューター向けの LLM 主導の効率的なコード オプティマイザー
要約: ムーアの法則の終焉により、特にハイパースケール データセンターでは、わずかな効率向上でも大幅なリソースとエネルギーの節約につながる、増え続けるコンピューティング需要を満たすために、コードのパフォーマンスを最適化することが最重要となっています。従来、このプロセスでは、最適化の機会を特定し、コードを変更して最適化を実装し、最適化の影響を慎重に展開して測定するという多大なプログラマーの労力が必要でした。プログラム編集の自動化には多大な労力が費やされ、小規模設定では有望な結果が得られているにもかかわらず、実際の大規模な運用環境では、そのようなパフォーマンスの最適化は、その規模、高度な複雑さ、および信頼性が求められるため、依然として実現が困難です。
このペーパーでは、ソース コードを自動的にリファクタリングして大規模なパフォーマンスを向上させるシステムである ECO (Efficient Code Optimizer) について紹介します。これらのパフォーマンスの向上を達成するために、ECO は履歴コミットを大規模に検索して、これらのコミットが対処したパフォーマンスのアンチパターンの辞書を作成します。これらのアンチパターンは、数十億行のコードからなるコード ベース内で同様のパターンを検索し、同様の潜在的な最適化機会を持つ他のコード セグメントを特定するために使用されます。 ECO は、微調整された LLM を使用してコードを自動的にリファクタリングし、同様の編集を生成して適用します。次に、ECO は変換されたコードを検証し、送信します。

コードレビューのためにそれを使用し、運用環境での最適化の影響を測定します。
現在、Google のハイパースケール運用フリートに導入されているこのシステムは、640 件を超える送信されたコミットにわたって、運用コードの 25,000 行を超える変更を実行し、運用成功率は 99.5% を超えています。過去 1 年間、ECO により四半期ごとに一貫して大幅なパフォーマンスの節約がもたらされました。平均すると、四半期ごとに得られる節約量は、正規化された CPU コア 50 万個以上に相当します。
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

Abstract page for arXiv paper 2503.15669: ECO: An LLM-Driven Efficient Code Optimizer for Warehouse Scale Computers

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2503.15669
Help | Advanced Search
Computer Science > Software Engineering
[Submitted on 19 Mar 2025]
Title: ECO: An LLM-Driven Efficient Code Optimizer for Warehouse Scale Computers
Abstract: With the end of Moore's Law, optimizing code for performance has become paramount for meeting ever-increasing compute demands, particularly in hyperscale data centers where even small efficiency gains translate to significant resource and energy savings. Traditionally, this process requires significant programmer effort to identify optimization opportunities, modify the code to implement the optimization, and carefully deploy and measure the optimization's impact. Despite a significant amount of work on automating program edits and promising results in small-scale settings, such performance optimizations have remained elusive in large real-world production environments, due to the scale, high degree of complexity, and reliability required.
This paper introduces ECO (Efficient Code Optimizer), a system that automatically refactors source code to improve performance at scale. To achieve these performance gains, ECO searches through historical commits at scale to create a dictionary of performance anti-patterns that these commits addressed. These anti-patterns are used to search for similar patterns in a code base of billions of lines of code, pinpointing other code segments with similar potential optimization opportunities. Using a fine-tuned LLM, ECO then automatically refactors the code to generate and apply similar edits. Next, ECO verifies the transformed code, submits it for code review, and measures the impact of the optimization in production.
Currently deployed on Google's hyperscale production fleet, this system has driven >25k changed lines of production code, across over 6.4k submitted commits, with a >99.5% production success rate. Over the past year, ECO has consistently resulted in significant performance savings every quarter. On average, the savings produced per quarter are equivalent to over 500k normalized CPU cores.
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

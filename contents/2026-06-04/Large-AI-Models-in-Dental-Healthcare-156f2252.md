---
source: "https://arxiv.org/abs/2606.02914"
hn_url: "https://news.ycombinator.com/item?id=48392947"
title: "Large AI Models in Dental Healthcare"
article_title: "[2606.02914] Large AI Models in Dental Healthcare: From General-Purpose Systems to Domain-Specific Foundation Models"
author: "berlianta"
captured_at: "2026-06-04T04:35:19Z"
capture_tool: "hn-digest"
hn_id: 48392947
score: 1
comments: 0
posted_at: "2026-06-04T02:32:49Z"
tags:
  - hacker-news
  - translated
---

# Large AI Models in Dental Healthcare

- HN: [48392947](https://news.ycombinator.com/item?id=48392947)
- Source: [arxiv.org](https://arxiv.org/abs/2606.02914)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T02:32:49Z

## Translation

タイトル: 歯科医療における大規模 AI モデル
記事タイトル: [2606.02914] 歯科医療における大規模 AI モデル: 汎用システムからドメイン固有の基盤モデルまで
説明: arXiv 論文 2606.02914 の要約ページ: 歯科医療における大規模 AI モデル: 汎用システムからドメイン固有の基盤モデルまで

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.02914
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 1 日に提出 ( v1 )、最終改訂日 2026 年 6 月 3 日 (このバージョン、v2)]
タイトル: 歯科医療における大規模 AI モデル: 汎用システムからドメイン固有の基盤モデルまで
要約: 背景: 口腔疾患は世界中で約 35 億人に影響を与えていますが、歯科における大規模 AI モデルの相対的な臨床的可能性は依然として十分に理解されていません。言語生成モデル、弁別視覚基礎モデル、歯科特有の基礎モデルという 3 つの異なるモデル カテゴリが出現しましたが、それらの関係や集合的な制限を検討する統一されたレビューはありません。
方法: PRISMA-ScR ガイドラインに従って、4 つのデータベース (PubMed、Google Scholar、Scopus、arXiv) を体系的に検索し、2 人の査読者によって独立してスクリーニングされました。包含/除外基準を適用した後、97 件の研究 (2020 ～ 2026 年) が含まれました。建築パラダイムと歯科専門度によってモデルを整理する二次元分類フレームワークを提案します。
結果: 言語生成モデルは、テキストベースのタスク (臨床推論、免許試験、患者とのコミュニケーション) には優れていますが、画像依存の診断では一貫性のないパフォーマンスを示します。適応された SAM および CLIP バリアントにより、強力な歯のセグメンテーションと病変検出結果が得られます。歯科専用モデル (DentVFM、DentVLM、OralGPT) は、複雑なマルチモーダルなタスクで最高のパフォーマンスを発揮します。統合されたパイプラインは、単一モデルのアプローチよりも常に優れたパフォーマンスを発揮します。データの非対称性が観察されます: 歯科特有の事前学習

ing はほぼ完全に視覚領域に集中しており、希少な大規模な歯科テキストコーパスを反映しています。
結論: 汎用モデルと歯科専用モデルは補完的な役割を果たします。最も効果的なシステムは、構造化されたパイプライン内で両方を組み合わせたものです。安全な自律展開には、生成モデルにおける幻覚、注釈付き歯科データセットの制限、標準化された臨床評価ベンチマークの欠如という 3 つの永続的な障壁を解決する必要があります。
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

Abstract page for arXiv paper 2606.02914: Large AI Models in Dental Healthcare: From General-Purpose Systems to Domain-Specific Foundation Models

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.02914
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 1 Jun 2026 ( v1 ), last revised 3 Jun 2026 (this version, v2)]
Title: Large AI Models in Dental Healthcare: From General-Purpose Systems to Domain-Specific Foundation Models
Abstract: Background: Oral diseases affect nearly 3.5 billion people worldwide, yet the comparative clinical potential of large-scale AI models in dentistry remains poorly understood. Three distinct model categories have emerged: language-generative models, discriminative vision foundation models, and dental-specific foundation models, with no unified review examining their relationships and collective limitations.
Methods: Following PRISMA-ScR guidelines, we systematically searched four databases (PubMed, Google Scholar, Scopus, arXiv), screened independently by two reviewers. After applying inclusion/exclusion criteria, 97 studies (2020-2026) were included. We propose a two-dimensional classification framework organizing models by architectural paradigm and dental specialization degree.
Results: Language-generative models excel at text-based tasks (clinical reasoning, licensing exams, patient communication) but show inconsistent performance on image-dependent diagnostics. Adapted SAM and CLIP variants achieve strong tooth segmentation and lesion detection results. Dental-specific models (DentVFM, DentVLM, OralGPT) demonstrate strongest performance on complex multimodal tasks. Integrated pipelines consistently outperform single-model approaches. A data asymmetry is observed: dental-specific pretraining concentrates almost entirely in the vision domain, reflecting scarce large-scale dental text corpora.
Conclusions: General-purpose and dental-specific models play complementary roles; the most effective systems combine both within structured pipelines. Safe autonomous deployment requires resolving three persistent barriers: hallucination in generative models, limited annotated dental datasets, and absent standardized clinical evaluation benchmarks.
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

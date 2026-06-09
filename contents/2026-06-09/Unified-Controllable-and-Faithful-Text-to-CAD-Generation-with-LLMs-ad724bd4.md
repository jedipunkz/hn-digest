---
source: "https://arxiv.org/abs/2604.19773"
hn_url: "https://news.ycombinator.com/item?id=48461311"
title: "Unified Controllable and Faithful Text-to-CAD Generation with LLMs"
article_title: "[2604.19773] PR-CAD: Progressive Refinement for Unified Controllable and Faithful Text-to-CAD Generation with Large Language Models"
author: "PaulHoule"
captured_at: "2026-06-09T16:05:55Z"
capture_tool: "hn-digest"
hn_id: 48461311
score: 22
comments: 0
posted_at: "2026-06-09T14:04:54Z"
tags:
  - hacker-news
  - translated
---

# Unified Controllable and Faithful Text-to-CAD Generation with LLMs

- HN: [48461311](https://news.ycombinator.com/item?id=48461311)
- Source: [arxiv.org](https://arxiv.org/abs/2604.19773)
- Score: 22
- Comments: 0
- Posted: 2026-06-09T14:04:54Z

## Translation

タイトル: LLM を使用した統合された制御可能で忠実なテキストから CAD への生成
記事のタイトル: [2604.19773] PR-CAD: 大規模な言語モデルを使用した、統合された制御可能かつ忠実なテキストから CAD への生成のための進歩的な改良
説明: arXiv 論文 2604.19773 の要約ページ: PR-CAD: 大規模な言語モデルを使用した統合制御可能かつ忠実なテキストから CAD 生成のための漸進的改良

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.19773
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 3 月 27 日に提出]
タイトル: PR-CAD: 大規模な言語モデルを使用した、統合された制御可能で忠実なテキストから CAD への生成のための進歩的な改良
要約: CAD モデルの構築は従来、労働集約的な手動操作と専門知識に依存していました。大規模言語モデル (LLM) の最近の進歩により、テキストから CAD への生成に関する研究が促進されました。ただし、既存のアプローチは通常、生成と編集を独立したタスクとして扱い、実用性が制限されています。私たちは、生成と編集を統合して制御可能で忠実なテキストから CAD モデリングを実現する進歩的な改良フレームワークである PR-CAD を提案します。これをサポートするために、私たちは、複数の CAD 表現と定性的および定量的記述の両方を含む、CAD ライフサイクル全体にわたる高忠実度のインタラクション データセットを厳選しました。このデータセットは編集操作の種類を体系的に定義し、非常に人間らしいインタラクション データを生成します。 LLM に合わせて調整された CAD 表現に基づいて、意図の理解、パラメータ推定、および正確な編集ローカリゼーションを単一のエージェントに統合する、強化学習によって強化された推論フレームワークを提案します。これにより、デザインの作成と改良の両方を行う「オールインワン」ソリューションが可能になります。広範な実験により、生成タスクと編集タスク間、および定性的モダリティと定量的モダリティ全体にわたる強力な相互強化が実証されています。公開ベンチマークにおいて、PR-CAD は、両方の点で最先端の制御性と忠実性を実現します。

生成および改良シナリオを実現すると同時に、ユーザーフレンドリーで CAD モデリングの効率を大幅に向上させます。
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

Abstract page for arXiv paper 2604.19773: PR-CAD: Progressive Refinement for Unified Controllable and Faithful Text-to-CAD Generation with Large Language Models

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.19773
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 27 Mar 2026]
Title: PR-CAD: Progressive Refinement for Unified Controllable and Faithful Text-to-CAD Generation with Large Language Models
Abstract: The construction of CAD models has traditionally relied on labor-intensive manual operations and specialized expertise. Recent advances in large language models (LLMs) have inspired research into text-to-CAD generation. However, existing approaches typically treat generation and editing as disjoint tasks, limiting their practicality. We propose PR-CAD, a progressive refinement framework that unifies generation and editing for controllable and faithful text-to-CAD modeling. To support this, we curate a high-fidelity interaction dataset spanning the full CAD lifecycle, encompassing multiple CAD representations as well as both qualitative and quantitative descriptions. The dataset systematically defines the types of edit operations and generates highly human-like interaction data. Building on a CAD representation tailored for LLMs, we propose a reinforcement learning-enhanced reasoning framework that integrates intent understanding, parameter estimation, and precise edit localization into a single agent. This enables an "all-in-one" solution for both design creation and refinement. Extensive experiments demonstrate strong mutual reinforcement between generation and editing tasks, and across qualitative and quantitative modalities. On public benchmarks, PR-CAD achieves state-of-the-art controllability and faithfulness in both generation and refinement scenarios, while also proving user-friendly and significantly improving CAD modeling efficiency.
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

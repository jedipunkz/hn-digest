---
source: "https://arxiv.org/abs/2606.11247"
hn_url: "https://news.ycombinator.com/item?id=48778089"
title: "Physics informed generative AI for semiconductor manufacturing"
article_title: "[2606.11247] Physics-informed generative AI for semiconductor manufacturing: Enforcing hard physical constraints in generative models by construction"
author: "Jimmc414"
captured_at: "2026-07-03T19:02:48Z"
capture_tool: "hn-digest"
hn_id: 48778089
score: 1
comments: 0
posted_at: "2026-07-03T18:17:31Z"
tags:
  - hacker-news
  - translated
---

# Physics informed generative AI for semiconductor manufacturing

- HN: [48778089](https://news.ycombinator.com/item?id=48778089)
- Source: [arxiv.org](https://arxiv.org/abs/2606.11247)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T18:17:31Z

## Translation

タイトル: 半導体製造のための物理学に基づいた生成 AI
記事タイトル: [2606.11247] 半導体製造のための物理学に基づいた生成 AI: 構築による生成モデルのハード物理的制約の強制
説明: arXiv 論文 2606.11247 の要約ページ: 半導体製造のための物理情報に基づいた生成 AI: 構築による生成モデルのハード物理的制約の強制

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 機械学習
[2026 年 6 月 8 日に提出]
タイトル: 半導体製造のための物理学に基づいた生成 AI: 構築による生成モデルの厳しい物理的制約の強制
要約: 生成モデルは、物理システムの設計、データ、および制御動作を提案するためにますます使用されていますが、そのようなシステムの多くは、知覚的な妥当性ではなく、厳しい物理的制約によって支配されています。半導体製造では、要求の厳しいテストケースが提供されます。物理的に無効なサンプルは、単に品質が低いだけでなく使用できないため、生成されたマスク、レイアウト、合成欠陥データ、およびプロセスレシピは、リソグラフィー、輸送、反応、およびデバイスの物理的制約に従わなければなりません。この展望では、半導体製造は、より広範な計算科学の課題を明らかにしている、つまり、制約された物理領域の生成 AI は、ポストホック フィルタリングだけで修正するのではなく、構築によって物理情報を反映する必要がある、と主張しています。私たちは、物理学に基づいた拡散、偏微分方程式に制約された変分モデル、神経演算子の事前分布、保存則を尊重した生成ネットワークなど、新たなアーキテクチャ ツールキットを調査し、それが微分可能リソグラフィー、TCAD、プロセス シミュレーション、および自律実験にどのように接続されるかを示します。私たちは、生成モデルと物理ベースのシミュレーターの間の 4 つの統合パターンを特定し、物理的忠実性ベンチマーク、微分可能なシミュレーター インフラストラクチャ、および物理設計と製造のためのマルチモーダル基礎モデルを中心とした研究課題を提案します。中心的な主張は修辞的というよりは分析的である。

妥当性は成功の拘束力のある基準であり、構築によってそれを強制するアーキテクチャは、事後的に妥当性をフィルタリングするアーキテクチャよりも優れたパフォーマンスを発揮することが期待されます。ファブは、この区別が最も明確になる環境です。
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

Abstract page for arXiv paper 2606.11247: Physics-informed generative AI for semiconductor manufacturing: Enforcing hard physical constraints in generative models by construction

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 8 Jun 2026]
Title: Physics-informed generative AI for semiconductor manufacturing: Enforcing hard physical constraints in generative models by construction
Abstract: Generative models are increasingly used to propose designs, data, and control actions for physical systems, yet many such systems are governed by hard physical constraints rather than by perceptual plausibility. Semiconductor manufacturing provides a demanding test case: generated masks, layouts, synthetic defect data, and process recipes must obey lithography, transport, reaction, and device-physics constraints, because physically invalid samples are not merely low quality but unusable. This Perspective argues that semiconductor manufacturing exposes a broader computational-science challenge, namely that generative AI for constrained physical domains must be physics-informed by construction, not corrected only through post-hoc filtering. We survey the emerging architectural toolkit, including physics-informed diffusion, PDE-constrained variational models, neural-operator priors, and conservation-law-respecting generative networks, and show how it connects to differentiable lithography, TCAD, process simulation, and autonomous experimentation. We identify four integration patterns between generative models and physics-based simulators, and we propose a research agenda centered on physics-fidelity benchmarks, differentiable simulator infrastructure, and multimodal foundation models for physical design and manufacturing. The central claim is analytical rather than rhetorical: where physical validity is the binding criterion of success, architectures that enforce it by construction should be expected to outperform those that filter for it after the fact, and the fab is the setting where this distinction is sharpest.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

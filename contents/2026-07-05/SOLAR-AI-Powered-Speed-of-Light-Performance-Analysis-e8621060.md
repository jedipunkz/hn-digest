---
source: "https://arxiv.org/abs/2606.26383"
hn_url: "https://news.ycombinator.com/item?id=48796062"
title: "SOLAR: AI-Powered Speed-of-Light Performance Analysis"
article_title: "[2606.26383] SOLAR: AI-Powered Speed-of-Light Performance Analysis"
author: "matt_d"
captured_at: "2026-07-05T17:57:12Z"
capture_tool: "hn-digest"
hn_id: 48796062
score: 2
comments: 0
posted_at: "2026-07-05T17:29:18Z"
tags:
  - hacker-news
  - translated
---

# SOLAR: AI-Powered Speed-of-Light Performance Analysis

- HN: [48796062](https://news.ycombinator.com/item?id=48796062)
- Source: [arxiv.org](https://arxiv.org/abs/2606.26383)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T17:29:18Z

## Translation

タイトル: SOLAR: AI を活用した光速パフォーマンス分析
記事のタイトル: [2606.26383] SOLAR: AI を活用した光速パフォーマンス分析
説明: arXiv 論文 2606.26383 の要約ページ: SOLAR: AI を活用した光速パフォーマンス分析

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
[2026 年 6 月 24 日に提出]
タイトル: SOLAR: AI を活用した光速パフォーマンス分析
要約: 深層学習モデルはターゲット ハードウェア上でどのくらいの速度で実行できますか?また、現在の実装はその制限からどのくらい離れていますか?これらの質問は、ソフトウェア、ハードウェア、アルゴリズムの最適化の中心となります。 Speed-of-Light (SOL) 分析は、特定のアーキテクチャでのワークロードの理論上の最小実行時間を計算することでこれらの課題に答えます。しかし、SOL 境界の導出は依然として手作業であり、エラーが発生しやすく、迅速なモデル開発とは切り離されています。このギャップを埋めるために、PyTorch および JAX ソース コードから検証済みの SOL 境界を自動的に導出するフレームワークである SOLAR を導入します。 SOLAR は、フロー内で生成コンポーネントと決定論コンポーネントの両方を活用します。LLM フロントエンドは、あらゆるソース プログラムを実行可能なアフィン ループ IR に変換し、出力比較によって検証します。決定論的なフローにより、IR が Einsum グラフに引き上げられます。分析バックエンドは、非融合、融合、およびキャッシュ対応の SOL 境界を計算します。 SOLAR は、オペレーターと言語を包括的にカバーし、SOL 違反が観察されない検証済みの境界を生成し、境界を強化して最適化の洞察を明らかにする多重忠実度分析を提供します。 KernelBench、JAX/Flax モデル、ロボット ワークロード全体で SOLAR を評価します。これらの実験では、複数の忠実度レベルでのヘッドルーム分析、最適化機会の特定、クロスプラットフォームの探索、およびインバース ルーフライン ハードウェア プロビジョニングの 4 つのユース ケースを示します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
コード、データ

この記事に関連するメディアおよびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.26383: SOLAR: AI-Powered Speed-of-Light Performance Analysis

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
[Submitted on 24 Jun 2026]
Title: SOLAR: AI-Powered Speed-of-Light Performance Analysis
Abstract: How fast could a deep-learning model run on target hardware, and how far is today's implementation from that limit? These questions are central to software, hardware, and algorithm optimizations. Speed-of-Light (SOL) analysis answers them by computing a workload's theoretical minimum execution time on a given architecture. Yet deriving SOL bounds remains manual, error-prone, and disconnected from rapid model development. To close this gap, we introduce SOLAR, a framework that automatically derives validated SOL bounds from PyTorch and JAX source code. SOLAR leverages both generative and deterministic components in its flow: an LLM frontend translates any source programs into an executable Affine Loop IR, validated by output comparison; a deterministic flow lifts the IR into an einsum graph; and an analytical backend computes unfused, fused, and cache-aware SOL bounds. SOLAR provides comprehensive operator and language coverage, produces validated bounds with zero observed SOL violations, and offers multi-fidelity analysis that tightens bounds and surfaces optimization insights. We evaluate SOLAR across KernelBench, JAX/Flax models, and robotics workloads. These experiments demonstrate four use cases: headroom analysis at multiple fidelity levels, identifying optimization opportunities, cross-platform exploration, and inverse-roofline hardware provisioning.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

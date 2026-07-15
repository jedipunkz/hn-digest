---
source: "https://arxiv.org/abs/2607.11262"
hn_url: "https://news.ycombinator.com/item?id=48921609"
title: "GPU-Tile-SIM: Tile-Centric GPU Simulation for LLM Hardware-Software Co-Design"
article_title: "[2607.11262] GPU-Tile-Sim: A Tile-Centric GPU Simulation Framework for LLM Hardware-Software Co-Design"
author: "rbanffy"
captured_at: "2026-07-15T15:19:18Z"
capture_tool: "hn-digest"
hn_id: 48921609
score: 1
comments: 0
posted_at: "2026-07-15T14:42:20Z"
tags:
  - hacker-news
  - translated
---

# GPU-Tile-SIM: Tile-Centric GPU Simulation for LLM Hardware-Software Co-Design

- HN: [48921609](https://news.ycombinator.com/item?id=48921609)
- Source: [arxiv.org](https://arxiv.org/abs/2607.11262)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T14:42:20Z

## Translation

タイトル: GPU-タイル-SIM: LLM ハードウェアとソフトウェアの協調設計のためのタイル中心の GPU シミュレーション
記事のタイトル: [2607.11262] GPU-Tile-Sim: LLM ハードウェアとソフトウェアの協調設計のためのタイル中心の GPU シミュレーション フレームワーク
説明: arXiv 論文 2607.11262 の要約ページ: GPU-Tile-Sim: LLM ハードウェアとソフトウェアの協調設計のためのタイル中心の GPU シミュレーション フレームワーク

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
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 7 月 13 日に提出]
タイトル: GPU-Tile-Sim: LLM ハードウェアとソフトウェアの協調設計のためのタイル中心の GPU シミュレーション フレームワーク
要約: 最新の LLM (大規模言語モデル) ワークロードは、ハードウェアとソフトウェアの共同設計を通じて最適化された GPU カーネルにますます依存しています。これらのカーネルは、きめ細かい依存関係のスケジューリングと計算メモリのオーバーラップを通じて、高いパフォーマンスを実現します。そのため、既存の GPU パフォーマンス モデルに新たな課題が生じます。命令駆動型シミュレータは、進化するアーキテクチャに適応するにはコストがかかりますが、分析モデルはカーネルの特性を捉えるには粗すぎます。私たちは、LLM ハードウェアとソフトウェアの共同設計のためのタイル中心の GPU シミュレーション フレームワークである GPU-Tile-Sim を提案します。重要な洞察は、最新の LLM カーネルのパフォーマンスは、個々の命令のレイテンシではなく、実行順序と重複を制御する依存構造によって左右されるということです。したがって、GTSim はカーネルの実行をワープ レベルのタイル グラフとして表します。そのノードはタイル レベルの操作をキャプチャし、そのエッジはデータと順序付けの制約をエンコードします。この表現を使用して、自動タイル グラフ フロントエンドとグラフ駆動シミュレーション バックエンドを設計します。代表的な GEMM、attention、エンドツーエンド LLM 推論ワークロードで GTSim を評価します。 A100 および H100 では、従来のカーネルと高度に最適化されたカーネルの両方で、GTSim は高いパフォーマンス モデリング精度 (MAPE、平均絶対パーセント誤差、1.22% ～ 8.71%) を実現します。予備検証により GTSim を Blackwell にさらに拡張し、ソフトウェアとアーキテクトの分析におけるその有効性を実証します。

ラルデザインの選択。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.11262: GPU-Tile-Sim: A Tile-Centric GPU Simulation Framework for LLM Hardware-Software Co-Design

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
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 13 Jul 2026]
Title: GPU-Tile-Sim: A Tile-Centric GPU Simulation Framework for LLM Hardware-Software Co-Design
Abstract: Modern LLM (large language model) workloads increasingly rely on optimized GPU kernels through hardware-software co-design. These kernels achieve high-performance through fine-grained dependency scheduling and computation-memory overlap. As such, they incur new challenges on existing GPU performance models. Instruction-driven simulators are costly to adapt to evolving architectures, while analytical models are too coarse to capture kernels' characteristics. We propose GPU-Tile-Sim, a tile-centric GPU simulation framework for LLM hardware-software co-design. The key insight is that modern LLM kernel performance is governed less by individual instruction latency than by the dependency structure that controls execution order and overlap. Accordingly, GTSim represents kernel execution as a warp-level tile graph whose nodes capture tile-level operations and whose edges encode data and ordering constraints. Using this representation, we design an automatic tile-graph frontend and a graph-driven simulation backend. We evaluate GTSim on representative GEMM, attention, and end-to-end LLM inference workloads. On A100 and H100 across both conventional and highly optimized kernels, GTSim achieves high performance-modeling accuracy (MAPE, Mean Absolute Percentage Error, 1.22%--8.71%). We further extend GTSim to Blackwell with preliminary validation, and demonstrate its effectiveness in analyzing software and architectural design choices.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

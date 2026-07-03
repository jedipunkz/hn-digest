---
source: "https://arxiv.org/abs/2607.02501"
hn_url: "https://news.ycombinator.com/item?id=48774455"
title: "Embodied.cpp: A Portable Inference Runtime of Embodied AI Models"
article_title: "[2607.02501] Embodied.cpp: A Portable Inference Runtime of Embodied AI Models on Heterogeneous Robots"
author: "chrsw"
captured_at: "2026-07-03T13:43:27Z"
capture_tool: "hn-digest"
hn_id: 48774455
score: 1
comments: 0
posted_at: "2026-07-03T12:55:02Z"
tags:
  - hacker-news
  - translated
---

# Embodied.cpp: A Portable Inference Runtime of Embodied AI Models

- HN: [48774455](https://news.ycombinator.com/item?id=48774455)
- Source: [arxiv.org](https://arxiv.org/abs/2607.02501)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T12:55:02Z

## Translation

タイトル: Embodied.cpp: 身体化された AI モデルのポータブル推論ランタイム
記事のタイトル: [2607.02501] Embodied.cpp: 異種ロボット上の組み込み AI モデルのポータブル推論ランタイム
説明: arXiv 論文 2607.02501 の要約ページ: Embodied.cpp: 異種ロボット上の組み込み AI モデルのポータブル推論ランタイム

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
コンピュータサイエンス > ロボティクス
[2026 年 7 月 2 日に提出]
タイトル: Embodied.cpp: 異種ロボット上の組み込み AI モデルのポータブル推論ランタイム
要約: 身体化された AI モデルは現在、ビジョン-言語-アクション (VLA) モデルとワールド アクション モデル (WAM) にまたがっていますが、実際の展開は、モデル固有の Python スタック、バックエンドの仮定、およびロボット側のグルー コード (特に異種エッジ デバイス上) にわたって断片化されたままです。既存の推論ランタイムは、主にリクエストとレスポンスの処理を目的として設計されているため、閉ループ制御内でのマルチレート実行、異種ハードウェアでのレイテンシーファーストのバッチ 1 推論、固定トークン I/O を超えた拡張可能な具体化されたインターフェイスなど、具体化された展開のランタイム契約を満たしていません。私たちは、具体化されたモデル用の移植可能な C++ 推論ランタイムである http URL を紹介します。代表的な VLA モデルと WAM のアーキテクチャ分析に基づいて、この http URL は共有実行パスをキャプチャし、入力アダプター、シーケンス ビルダー、バックボーン実行、ヘッド プラグイン、展開アダプターの 5 つのレイヤーに編成します。このランタイムは、モジュール式のマルチレート実行、レイテンシーファーストの融合推論、拡張可能なオペレーターと I/O サポートを提供し、1 つのバックエンド抽象化を通じて異種デバイス、ロボット、シミュレーターにわたる展開を可能にします。この http URL を 2 つの VLA モデル (HY-VLA と pi0.5) と、LingBot-VA Transformer ブロックを使用した予備的な WAM ベンチマークで評価します。 VLA 導入では、それぞれ 100.0% と 91.0% のタスク成功率で閉ループ実行が成功しました。 WAM ベンチマークは、ブロック メモリを 312.2 MiB から 88.1 MiB に削減します。 Th

これらの結果は、この http URL が、さまざまな具体化されたモデル アーキテクチャにわたって高い精度を維持しながら、展開効率を向上させることを示しています。
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

Abstract page for arXiv paper 2607.02501: Embodied.cpp: A Portable Inference Runtime of Embodied AI Models on Heterogeneous Robots

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
Computer Science > Robotics
[Submitted on 2 Jul 2026]
Title: Embodied.cpp: A Portable Inference Runtime of Embodied AI Models on Heterogeneous Robots
Abstract: Embodied AI models now span vision-language-action (VLA) models and world-action models (WAMs), but practical deployment remains fragmented across model-specific Python stacks, backend assumptions, and robot-side glue code, especially on heterogeneous edge devices. Existing inference runtimes are designed mainly for request-response serving and therefore do not satisfy the runtime contract of embodied deployment: multi-rate execution inside closed-loop control, latency-first batch-1 inference on heterogeneous hardware, and extensible embodied interfaces beyond fixed token I/O. We present this http URL , a portable C++ inference runtime for embodied models. Based on an architectural analysis of representative VLA models and WAMs, this http URL captures a shared execution path and organizes it into five layers: input adapters, sequence builders, backbone execution, head plugins, and deployment adapters. The runtime provides modular multi-rate execution, latency-first fused inference, and extensible operator and I/O support, enabling deployment across heterogeneous devices, robots, and simulators through one backend abstraction. We evaluate this http URL on two VLA models, HY-VLA and pi0.5, and on a preliminary WAM benchmark using a LingBot-VA Transformer block. The VLA deployments achieve successful closed-loop execution with 100.0% and 91.0% task success rates, respectively. The WAM benchmark reduces block memory from 312.2 MiB to 88.1 MiB. These results show that this http URL improves deployment efficiency while preserving high accuracy across diverse embodied model architectures.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

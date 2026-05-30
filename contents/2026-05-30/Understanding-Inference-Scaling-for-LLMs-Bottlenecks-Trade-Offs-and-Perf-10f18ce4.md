---
source: "https://arxiv.org/abs/2605.19775"
hn_url: "https://news.ycombinator.com/item?id=48327924"
title: "Understanding Inference Scaling for LLMs: Bottlenecks, Trade-Offs, and Perf"
article_title: "[2605.19775] Understanding Inference Scaling for LLMs: Bottlenecks, Trade-offs, and Performance Principles"
author: "matt_d"
captured_at: "2026-05-30T11:40:29Z"
capture_tool: "hn-digest"
hn_id: 48327924
score: 6
comments: 0
posted_at: "2026-05-29T19:14:15Z"
tags:
  - hacker-news
  - translated
---

# Understanding Inference Scaling for LLMs: Bottlenecks, Trade-Offs, and Perf

- HN: [48327924](https://news.ycombinator.com/item?id=48327924)
- Source: [arxiv.org](https://arxiv.org/abs/2605.19775)
- Score: 6
- Comments: 0
- Posted: 2026-05-29T19:14:15Z

## Translation

タイトル: LLM の推論スケーリングについて: ボトルネック、トレードオフ、パフォーマンス
記事のタイトル: [2605.19775] LLM の推論スケーリングについて: ボトルネック、トレードオフ、およびパフォーマンスの原則
説明: arXiv 論文 2605.19775 の要約ページ: LLM の推論スケーリングの理解: ボトルネック、トレードオフ、およびパフォーマンスの原則

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.19775
ヘルプ |高度な検索
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 5 月 19 日に提出]
タイトル: LLM の推論スケーリングを理解する: ボトルネック、トレードオフ、パフォーマンス原則
要約: 標準的な生成 AI から、広範な思考連鎖 (CoT) 処理が可能なモデルに代表される \emph{推論中心のアーキテクチャ} への移行は、システム要件の根本的なパラダイム シフトを示しています。計算依存のプリフィルが大半を占める従来のワークロードとは異なり、推論ワークロードは、推論を \emph{容量制限体制} に移行させる推論トークンの長いチェーンを生成します。このペーパーでは、GPU クラスター上の 8B から 671B のパラメーターにわたるモデルを評価する、包括的なシステムの特性評価について説明します。データ、テンソル、パイプラインの並列処理間の相​​互作用を体系的に調査することで、標準のスケーリング ヒューリスティックを無視する重大なボトルネックを特定します。私たちの分析により、データ並列処理は小規模モデルではスループット効率が高いものの、KV キャッシュの断片化により早期のスロットリングが強制され、コンピューティング使用率が最適化されず、推論ワークロードで容量の罠に陥ることが明らかになりました。テンソル並列処理により、行き詰まったメモリが解放され、32B クロスオーバー付近でサブリニアなゲインが実現します。フロンティア規模では、高密度モデル (Llama-405B など) は相互接続とメモリ帯域幅に制限があり、高度な TP を優先しますが、疎な専門家混合 (MoE) モデル (DeepSeek-R1 など) はルーティングと同期レイテンシによって制限され、ハイブリッド戦略の恩恵を受けます。これらの洞察は、厳密な意思決定の枠組みを提供します

推論の崖を乗り越え、次世代の推論インフラストラクチャのための新しいアーキテクチャ上の必須事項を確立するために。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
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

Abstract page for arXiv paper 2605.19775: Understanding Inference Scaling for LLMs: Bottlenecks, Trade-offs, and Performance Principles

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.19775
Help | Advanced Search
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 19 May 2026]
Title: Understanding Inference Scaling for LLMs: Bottlenecks, Trade-offs, and Performance Principles
Abstract: The transition from standard generative AI to \emph{reasoning-centric architectures}, exemplified by models capable of extensive Chain-of-Thought~(CoT) processing, marks a fundamental paradigm shift in system requirements. Unlike traditional workloads dominated by compute-bound prefill, reasoning workloads generate long chains of reasoning tokens that shift inference into a \emph{Capacity-Bound regime}. This paper presents a comprehensive system characterization, evaluating models ranging from 8B to 671B parameters on GPUs clusters. By systematically exploring the interplay between Data, Tensor, and Pipeline parallelism, we identify critical bottlenecks that defy standard scaling heuristics. Our analysis reveals that data parallelism is throughput efficient for small models but hits a capacity trap on reasoning workloads as KV-cache fragmentation forces early throttling resulting in sub-optimal compute utilization. Tensor parallelism unlocks stranded memory and delivers sublinear gains near the 32B crossover. At frontier scale, dense models (e.g., Llama-405B) are interconnect and memory-bandwidth bound and favor high-degree TP, while sparse Mixture-of-Experts (MoE) models (e.g., DeepSeek-R1) are limited by routing and synchronization latency and benefit from hybrid strategies. These insights provide a rigorous decision framework for navigating the reasoning cliff, establishing new architectural imperatives for the next generation of inference infrastructure.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
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

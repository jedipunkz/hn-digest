---
source: "https://arxiv.org/abs/2310.03003"
hn_url: "https://news.ycombinator.com/item?id=48827556"
title: "From Words to Watts: Benchmarking the Energy Costs of LLM Inference (2023)"
article_title: "[2310.03003] From Words to Watts: Benchmarking the Energy Costs of Large Language Model Inference"
author: "teleforce"
captured_at: "2026-07-08T05:00:14Z"
capture_tool: "hn-digest"
hn_id: 48827556
score: 1
comments: 0
posted_at: "2026-07-08T04:33:20Z"
tags:
  - hacker-news
  - translated
---

# From Words to Watts: Benchmarking the Energy Costs of LLM Inference (2023)

- HN: [48827556](https://news.ycombinator.com/item?id=48827556)
- Source: [arxiv.org](https://arxiv.org/abs/2310.03003)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T04:33:20Z

## Translation

タイトル: ワードからワットへ: LLM 推論のエネルギーコストのベンチマーク (2023)
記事のタイトル: [2310.03003] 単語からワットへ: 大規模言語モデル推論のエネルギー コストのベンチマーク
説明: arXiv 論文 2310.03003 の要約ページ: From Words to Watts: Benchmarking the Energy Costs of Large Language Model Inference

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
コンピュータサイエンス > 計算と言語
[2023 年 10 月 4 日に提出]
タイトル: ワードからワットへ: 大規模言語モデル推論のエネルギーコストのベンチマーク
要約: 大規模言語モデル (LLM) は、以前の最先端技術をはるかに超えた新しい生成機能により、爆発的に人気が高まっています。これらのテクノロジーは、法律、金融、医療などのさまざまな分野でますます活用されています。ただし、これらのモデルには、特に推論に必要な計算コストとエネルギー コストといった、計算上の大きな課題があります。実際に推論を行うためにこれらの大規模なモデルがどれほど頻繁に呼び出されるかにもかかわらず (ChatGPT など)、推論のエネルギー コストは、LLM のトレーニングのエネルギー コストほど注目されていません。これらの最先端の LLM は、さまざまなドメインでの使用と導入が増加しているため、コスト削減、パフォーマンスのスケーリング、ハードウェアの効率的な使用、最適な推論戦略にとって、そのリソース使用率をより深く理解することが重要です。
この論文では、LLM を使用した推論の計算とエネルギーの利用を研究するために行われた実験について説明します。研究と実践における LLM の多様なタスク/ベンチマーク セットを反映するために、2 世代の人気のある GPU (NVIDIA V100 \& A100) と 2 つのデータセット (Alpaca と GSM8K) 上で Meta AI によって開発された、さまざまなサイズの LLaMA (最近の最先端 LLM) の推論パフォーマンスと推論エネルギー コストの予備分析をベンチマークおよび実行します。最大 32 GPU にわたるモデル シャーディングを使用したマルチノード、マルチ GPU 推論の結果を示します。私たちの知る限り、私たちの仕事は次のようなものです。

彼はまず、この規模での計算リソースとエネルギー リソースの観点から LLM 推論のパフォーマンスを研究しました。
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

Abstract page for arXiv paper 2310.03003: From Words to Watts: Benchmarking the Energy Costs of Large Language Model Inference

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
Computer Science > Computation and Language
[Submitted on 4 Oct 2023]
Title: From Words to Watts: Benchmarking the Energy Costs of Large Language Model Inference
Abstract: Large language models (LLMs) have exploded in popularity due to their new generative capabilities that go far beyond prior state-of-the-art. These technologies are increasingly being leveraged in various domains such as law, finance, and medicine. However, these models carry significant computational challenges, especially the compute and energy costs required for inference. Inference energy costs already receive less attention than the energy costs of training LLMs -- despite how often these large models are called on to conduct inference in reality (e.g., ChatGPT). As these state-of-the-art LLMs see increasing usage and deployment in various domains, a better understanding of their resource utilization is crucial for cost-savings, scaling performance, efficient hardware usage, and optimal inference strategies.
In this paper, we describe experiments conducted to study the computational and energy utilization of inference with LLMs. We benchmark and conduct a preliminary analysis of the inference performance and inference energy costs of different sizes of LLaMA -- a recent state-of-the-art LLM -- developed by Meta AI on two generations of popular GPUs (NVIDIA V100 \& A100) and two datasets (Alpaca and GSM8K) to reflect the diverse set of tasks/benchmarks for LLMs in research and practice. We present the results of multi-node, multi-GPU inference using model sharding across up to 32 GPUs. To our knowledge, our work is the one of the first to study LLM inference performance from the perspective of computational and energy resources at this scale.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

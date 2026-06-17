---
source: "https://www.day1training.com/"
hn_url: "https://news.ycombinator.com/item?id=48567755"
title: "Distributed AI on AWS"
article_title: "AWSome Distributed AI | Day1Training"
author: "gjmveloso"
captured_at: "2026-06-17T10:38:30Z"
capture_tool: "hn-digest"
hn_id: 48567755
score: 2
comments: 0
posted_at: "2026-06-17T09:11:09Z"
tags:
  - hacker-news
  - translated
---

# Distributed AI on AWS

- HN: [48567755](https://news.ycombinator.com/item?id=48567755)
- Source: [www.day1training.com](https://www.day1training.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T09:11:09Z

## Translation

タイトル: AWS 上の分散型 AI
記事のタイトル: AWSome 分散型 AI | 1日目トレーニング
説明: AWS での分散 ML トレーニングのリファレンス アーキテクチャとベスト プラクティス

記事本文:
AWSome 分散 AI | 1日目トレーニング
1日目トレーニング
フレームワーク アーキテクチャ ベンチマーク ワークショップ 検索... ⌘K フレームワーク アーキテクチャ ベンチマーク ワークショップ
オープンソース • MIT-0 ライセンス
すごい
分散型AI
AWS インフラストラクチャ上で PyTorch、Megatron-LM、NeMo、JAX などを使用して大規模モデルをトレーニングするためのリファレンス アーキテクチャ、テスト ケース、ベスト プラクティス。
フレームワークを探索する
はじめに
30 を超えるテスト ケース 10 のアーキテクチャ 4 つのフレームワーク 1.5K コミットのトレーニング フレームワーク
フレームワークごとにグループ化された実稼働対応のサンプル。それぞれに、Dockerfile、Slurm スクリプト、Kubernetes マニフェストが含まれています。
DDP、FSDP、TorchTitan、DeepSpeed などを使用したネイティブ分散トレーニング。 LLM、ビジョン、ロボット工学、RLHF をカバーします。
テンソル、パイプライン、エキスパート並列処理を使用した大規模な LLM 事前トレーニング用の NVIDIA Megatron-LM および NeMo。
XLA コンパイルと自動並列処理を活用した分散トレーニング用の PaxML を使用した Google JAX。
NeuronX は、最適化されたコンパイラーを備えた AWS Trainium & Inferentia チップでのトレーニングのために配布されています。
NVIDIA Isaac Lab、OpenVLA、V-JEPA2、およびビジョン言語アクション モデルを使用した、具体化された AI トレーニング。
RLHF、DPO、PPO、LLM 調整およびトレーニング後のスケーラブルな RL フレームワーク。
生産のための知識の蒸留、圧縮、およびモデル適応技術。
すべての AWS コンピューティング プラットフォーム用の CloudFormation テンプレートとデプロイ ガイド。
最初の分散トレーニング ジョブを開始するための 3 つのステップ。
HyperPod、ParallelCluster、または EKS 用の CloudFormation テンプレートを使用してクラスターを起動します。
Dockerfile を使用して、選択したフレームワークでトレーニング コンテナを構築します。
既製の起動スクリプトを使用して、Slurm または Kubernetes でジョブを送信します。
1日目トレーニング
オープンソースのリファレンス アーキテクチャと、AWS での分散 ML トレーニングのベスト プラクティス。

## Original Extract

Reference architectures and best practices for distributed ML training on AWS

AWSome Distributed AI | Day1Training
Day1Training
Frameworks Architectures Benchmarks Workshops Search... ⌘K Frameworks Architectures Benchmarks Workshops
Open Source • MIT-0 License
AWSome
Distributed AI
Reference architectures, test cases, and best practices for training large-scale models with PyTorch, Megatron-LM, NeMo, JAX, and more on AWS infrastructure.
Explore Frameworks
Getting Started
30+ Test Cases 10 Architectures 4 Frameworks 1.5K Commits Training Frameworks
Production-ready examples grouped by framework. Each includes Dockerfiles, Slurm scripts, and Kubernetes manifests.
Native distributed training with DDP, FSDP, TorchTitan, DeepSpeed, and more. Covers LLMs, vision, robotics, and RLHF.
NVIDIA Megatron-LM and NeMo for large-scale LLM pre-training with tensor, pipeline, and expert parallelism.
Google JAX with PaxML for distributed training leveraging XLA compilation and automatic parallelism.
NeuronX Distributed for training on AWS Trainium & Inferentia chips with optimized compilers.
Embodied AI training with NVIDIA Isaac Lab, OpenVLA, V-JEPA2, and vision-language-action models.
RLHF, DPO, PPO, and scalable RL frameworks for LLM alignment and post-training.
Knowledge distillation, compression, and model adaptation techniques for production.
CloudFormation templates and deployment guides for every AWS compute platform.
Three steps to launch your first distributed training job.
Launch a cluster using our CloudFormation templates for HyperPod, ParallelCluster, or EKS.
Use our Dockerfiles to build a training container with your framework of choice.
Submit your job with Slurm or Kubernetes using our ready-made launch scripts.
Day1Training
Open-source reference architectures and best practices for distributed ML training on AWS.

---
source: "https://documentation.suse.com/suse-ai/1.0/html/AI-deployment/index.html"
hn_url: "https://news.ycombinator.com/item?id=49304513"
title: "Deploying and Installing SUSE AI"
article_title: "Deploying and Installing SUSE AI | SUSE AI 1.0"
author: "indigodaddy"
captured_at: "2026-08-14T21:16:27Z"
capture_tool: "hn-digest"
hn_id: 49304513
score: 2
comments: 0
posted_at: "2026-08-14T20:58:40Z"
tags:
  - hacker-news
  - translated
---

# Deploying and Installing SUSE AI

- HN: [49304513](https://news.ycombinator.com/item?id=49304513)
- Source: [documentation.suse.com](https://documentation.suse.com/suse-ai/1.0/html/AI-deployment/index.html)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T20:58:40Z

## Translation

タイトル: SUSE AI のデプロイとインストール
記事のタイトル: SUSE AI のデプロイとインストール | SUSE AI 1.0
説明: この包括的なステップバイステップガイドを使用して、SUSE AI をシームレスに展開およびインストールします。

記事本文:
|インデックス | SUSE AI の展開とインストール SUSE AI の展開とインストール 1 インストールの概要
2 Linux および Kubernetes ディストリビューションのインストール
3 Preparing the cluster for AI Library
4 Installing applications from AI Library
C GNU Free Documentation License
Deploying and Installing SUSE AI
SUSE AI は、複数のソフトウェア コンポーネントと依存関係で構成される多層エンタープライズ AI ソリューションです。
このガイドを使用して、SUSE AI とその依存関係を一貫したサポートされた方法でデプロイおよびインストールします。
読む時間は 1 時間もかかりませんが、Linux 導入に関する高度な知識が必要です。
このガイドを終えると、SUSE AI と必要な依存関係がすべてインストールされ、使用できるようになります。
2 Linux および Kubernetes ディストリビューションのインストール 2.1 SUSE Linux Enterprise Server のインストール
2.2 Installing NVIDIA GPU drivers
2.3 Installing SUSE Rancher Prime: RKE2
3 AI Library 用のクラスターの準備 3.1 Kubernetes クラスターへの SUSE Rancher Prime のインストール
3.2 SUSE Rancher Prime: RKE2 クラスターへの NVIDIA GPU Operator のインストール
3.3 Registering existing clusters
3.4 Assigning GPU nodes to applications
3.6 Setting up SUSE Observability for SUSE AI
4 AI ライブラリからのアプリケーションのインストール 4.1 SUSE Application Collection とは何ですか?
4.16 Verifying SUSE AI Library applications
5 Alternative deployment 5.1 Node installer
C GNU Free Documentation License C1 0. PREAMBLE
C2 1. APPLICABILITY AND DEFINITIONS
C7 6. COLLECTIONS OF DOCUMENTS
C8 7. AGGREGATION WITH INDEPENDENT WORKS
C11 1. FUTURE REVISIONS OF THIS LICENSE
C12 付録: このライセンスを文書に使用する方法
1.1 SUSE AI installation process
2.1 Language, keyboard and product selection
2.5 Extension and module selection
2.11 Password for the system administrator root
2.14 tの実行

彼はインストール
3.5 SUSE セキュリティ アプリケーションのインストール
3.12 Neuvector 管理者ユーザー
3.15 SUSE Observability セットアップの概要
3.16 SUSE AI と SUSE Observability 用の個別のクラスター
3.17 新しい GenAI 可観測性メニュー項目
4.1 SUSE アプリケーション コレクションの Milvus ページ
4.2 Open WebUI への vLLM 接続の追加
4.1 Ollama Helm チャートのファイル オプションを上書きする
4.2 Open WebUI Helm チャートで使用可能なオプション
4.3 PyTorch Helm チャートで使用可能なオプション
4.4 LiteLLM Helm チャートのオプション
4.1 Ollama Helm チャート バージョン 0.x.x
4.2 Ollama Helm チャート バージョン 1.x.x
4.3 GPU と起動時にプルされる 2 つのモデルを使用した基本的なオーバーライド ファイル
4.4 Ingress を使用し、GPU を使用しない基本的なオーバーライド ファイル
4.5 Ollama を含む WebUI オーバーライド ファイルを開く
4.6 Ollama を別途インストールして WebUI オーバーライド ファイルを開く
4.7 パイプラインを有効にして WebUI オーバーライド ファイルを開く
4.8 vLLM に接続して WebUI オーバーライド ファイルを開く
4.9 オープン Webui パイプラインのスタンドアロン展開
4.11 インストールの検証
4.13 プリフェッチされたモデルを永続ストレージからロードする
4.14 複数モデルによる構成
4.16 LMCache を使用した共有リモート KV キャッシュ ストレージ
4.17 GPU が有効な場合の基本的なオーバーライド ファイル
4.19 チャートに焼き付けられたファイルを含むホストフォルダー
4.20 Git リポジトリ クローン: 認証なしのパブリック
4.21 Git リポジトリ クローン: 認証付きのプライベート
4.22 クラスターにデフォルトのストレージ クラスが設定されていない場合の基本オーバーライド ファイル。
4.23 Qdrant が GPU 機能を使用する例。
4.24 PostgreSQL デプロイメントとマスター キーが自動的に生成される基本的なオーバーライド ファイル。
4.25 SUSE Application Collection でホストされているコンテナーと Helm チャートの検証
4.26 SUSE レジストリでホストされているコンテナと Helm チャートの検証
4.27 添付された証明書の検出 (オプション)
4.28 抜粋

CycloneDX SBOM と脆弱性スキャンの実行 (オプション)
5.1 サーバー上に OS がすでにプロビジョニングされており、Rancher 管理サーバーを使用して RKE2 クラスターを展開したいと考えています。
5.2 サーバー上に OS がすでにプロビジョニングされており、SUSE AI ワークロードの実行環境として機能する RKE2 クラスターをデプロイしたいと考えています。

## Original Extract

Seamlessly deploy and install SUSE AI with this comprehensive, step-by-step guide

| Index | Deploying and Installing SUSE AI Deploying and Installing SUSE AI 1 Installation overview
2 Installing the Linux and Kubernetes distribution
3 Preparing the cluster for AI Library
4 Installing applications from AI Library
C GNU Free Documentation License
Deploying and Installing SUSE AI
SUSE AI is a multi-layered enterprise AI solution that consists of several software components and dependencies.
Use this guide to deploy and install SUSE AI and its dependencies in a consistent and supported manner.
It takes less than one hour of reading time and requires advanced knowledge of Linux deployment.
By the end of this guide, you will have SUSE AI and all required dependencies installed and ready for use.
2 Installing the Linux and Kubernetes distribution 2.1 Installing SUSE Linux Enterprise Server
2.2 Installing NVIDIA GPU drivers
2.3 Installing SUSE Rancher Prime: RKE2
3 Preparing the cluster for AI Library 3.1 Installing SUSE Rancher Prime on a Kubernetes cluster
3.2 Installing the NVIDIA GPU Operator on the SUSE Rancher Prime: RKE2 cluster
3.3 Registering existing clusters
3.4 Assigning GPU nodes to applications
3.6 Setting up SUSE Observability for SUSE AI
4 Installing applications from AI Library 4.1 What is SUSE Application Collection?
4.16 Verifying SUSE AI Library applications
5 Alternative deployment 5.1 Node installer
C GNU Free Documentation License C1 0. PREAMBLE
C2 1. APPLICABILITY AND DEFINITIONS
C7 6. COLLECTIONS OF DOCUMENTS
C8 7. AGGREGATION WITH INDEPENDENT WORKS
C11 1. FUTURE REVISIONS OF THIS LICENSE
C12 ADDENDUM: How to use this License for your documents
1.1 SUSE AI installation process
2.1 Language, keyboard and product selection
2.5 Extension and module selection
2.11 Password for the system administrator root
2.14 Performing the installation
3.5 Install SUSE Security application
3.12 Neuvector administrator users
3.15 High-level overview of the SUSE Observability setup
3.16 Separate clusters for SUSE AI and SUSE Observability
3.17 New GenAI Observability menu item
4.1 Milvus page in the SUSE Application Collection
4.2 Adding a vLLM connection to Open WebUI
4.1 Override file options for the Ollama Helm chart
4.2 Available options for the Open WebUI Helm chart
4.3 Available options for the PyTorch Helm chart
4.4 Options for the LiteLLM Helm chart
4.1 Ollama Helm chart version 0.x.x
4.2 Ollama Helm chart version 1.x.x
4.3 Basic override file with GPU and two models pulled at startup
4.4 Basic override file with Ingress and no GPU
4.5 Open WebUI override file with Ollama included
4.6 Open WebUI override file with Ollama installed separately
4.7 Open WebUI override file with pipelines enabled
4.8 Open WebUI override file with a connection to vLLM
4.9 Stand-alone deployment of open-webui-pipelines
4.11 Validating the installation
4.13 Loading prefetched models from persistent storage
4.14 Configuration with multiple models
4.16 Shared remote KV cache storage with LMCache
4.17 Basic override file with GPU enabled
4.19 Host-folder with files baked into the chart
4.20 Git repository clone: public with no authentication
4.21 Git repository clone: private with authentication
4.22 Basic override file when the cluster has no default storage class set.
4.23 An example where Qdrant uses GPU capabilities.
4.24 Basic override file with PostgreSQL deployment and master key automatically generated.
4.25 Verifying containers and Helm charts hosted on the SUSE Application Collection
4.26 Verifying containers and Helm charts hosted on the SUSE Registry
4.27 Discovering attached attestations (optional)
4.28 Extracting the CycloneDX SBOM and vulnerability scan (optional)
5.1 You already have an OS provisioned on the servers and you want to deploy an RKE2 cluster with a Rancher management server.
5.2 You already have an OS provisioned on the servers and you want to deploy an RKE2 cluster to serve as the execution environment for SUSE AI workloads.

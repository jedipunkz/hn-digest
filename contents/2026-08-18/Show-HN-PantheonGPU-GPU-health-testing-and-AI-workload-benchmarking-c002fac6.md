---
source: "https://pantheongpu.com/"
hn_url: "https://news.ycombinator.com/item?id=49350637"
title: "Show HN: PantheonGPU – GPU health testing and AI workload benchmarking"
article_title: "PANTHEON"
image: "https://pantheongpu.com/assets/logo.png"
author: "saqibkhan1992"
captured_at: "2026-08-18T19:21:47Z"
capture_tool: "hn-digest"
hn_id: 49350637
score: 6
comments: 0
posted_at: "2026-08-18T18:47:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PantheonGPU – GPU health testing and AI workload benchmarking

- HN: [49350637](https://news.ycombinator.com/item?id=49350637)
- Source: [pantheongpu.com](https://pantheongpu.com/)
- Score: 6
- Comments: 0
- Posted: 2026-08-18T18:47:51Z

## Translation

タイトル: Show HN: PantheonGPU – GPU 健全性テストと AI ワークロード ベンチマーク
記事タイトル: パンテオン
説明: クロスプラットフォームの CUDA および ROCm GPU ストレス テスト、診断、プロファイリング、およびパフォーマンス ベンチマーク。
HN テキスト: こんにちは、HN。私が PantheonGPU を構築したのは、この GPU が実際に健全で、正常に動作しているかという単純な質問に答えるためのより良い方法が欲しかったからです。 GPU は正常な温度と使用率を示していても、パフォーマンスが低下したり、特定のワークロードの下で不安定になったり、メモリ、PCIe、構成に問題がある場合があります。 PantheonGPU は、テレメトリを監視するだけでなく、GPU をアクティブにテストします。現在、コンピューティング、テンソル ワークロード、メモリ、キャッシュ、PCIe、熱、安定性、AI/LLM 推論をカバーする 45 以上のテストが含まれています。 NVIDIA CUDA と AMD ROCm の両方をサポートします。また、より大きなユースケースも検討しています。それは、GPU フリート全体で Pantheon を実行して、サーバーまたはクラスターの他の部分とは異なる動作をする個々の GPU を識別することです。 AI インフラストラクチャ、マルチ GPU システム、ローカル LLM、または GPU クラウドを実行している人々からのフィードバックを特に歓迎します。

記事本文:
パンテオン
コンテンツにスキップ
パンテオン
ホーム
検索を初期化しています
ホーム
ホーム
ホーム
目次
クイックスタート
1. 前提条件をインストールする
パフォーマンスデータベース
パフォーマンスデータベース
ライブベンチマーク
調査とレポート
調査とレポート
概要
テスト文書
テスト文書
概要
コアとコンピューティング
コアとコンピューティング
オムニウイルス
固定関数とアクセラレータ
固定関数とアクセラレータ
RTウイルス
メモリとキャッシュ
メモリとキャッシュ
原子ウイルス
相互接続とアーキテクチャ
相互接続とアーキテクチャ
P2Pスラッシャー
AI と機械学習
AI と機械学習
AI ワークロード スイート
クイックスタート
1. 前提条件をインストールする
GPUのストレステストと診断
Pantheon は、GPU のコンピューティング、メモリ、キャッシュ、相互接続、電力の動作をテストします。集中したワークロードを実行し、テレメトリをキャプチャし、比較のために結果を保存します。
Debian パッケージは、Ubuntu および Debian システムの最も簡単なインストール パスです。
基本的なビルド ツールをインストールします。
sudo apt-get アップデート
sudo apt-get install -y make g++
次に、GPU プラットフォーム用のコンパイラーをインストールします。必要なのは 1 つだけです:
sudo apt-get install -y nvidia-cuda-toolkit
sudo apt-get install -y hipcc
2.パンテオンをインストールする
最新の Debian パッケージをダウンロードしてインストールします。
バージョン = 1.0.14
wget "https://github.com/saqibkh/pantheongpu_website/releases/download/v ${ VERSION } /pantheongpu_ ${ VERSION } _amd64.deb"
sudo apt install "./pantheongpu_ ${ VERSION } _amd64.deb"
後で Debian パッケージをアンインストールするには:
sudo apt-get パンテオンプを削除します
3. インストールを確認する
短いハードウェア インベントリ テストを実行します。
パンテオン --テストベースライン_メトリクス --期間 10
次に、GPU 0 で対象を絞ったストレス テストを実行します。
パンテオン --test fp64_virus --duration 30 --gpu 0
注記
Pantheon は、CUDA、ROCm/HIP、またはモック モードを自動的に検出します。パンテオンを運営する
直接命令する。 --platform cuda を渡す必要はありません。
ネイティブパッケージコマンド

上記の d は、Pantheon のパッケージ管理ファイルを削除します。
ランタイムで作成されたファイルと現在のユーザーのコンパイルされたワークロードも削除するには
キャッシュ、または RHEL、Fedora、Rocky Linux 上のポータブル インストールを削除するには、
AlmaLinux、または別の Linux ディストリビューションは、以下を実行します。
カール -fsSL https://pantheongpu.com/uninstall.sh |須藤し
これにより、CUDA、ROCm、システム コンパイラ、ベンチマーク レポートが外部に保存されたままになります。
Pantheon のインストール ディレクトリとキャッシュ ディレクトリはそのままです。
リリース バンドルには、Debian パッケージと install.sh ヘルパーが含まれています。
RHEL ファミリおよびその他の Linux ディストリビューション。
バージョン = 1.0.14
wget "https://github.com/saqibkh/pantheongpu_website/releases/download/v ${ VERSION } /pantheongpu_ ${ VERSION } _amd64.tar.gz"
tar -xzf "pantheongpu_ ${ VERSION } _amd64.tar.gz"
cd "パンテオンプ_ ${ バージョン } _amd64"
sudo apt install "./packages/pantheongpu_ ${ VERSION } _amd64.deb"
次のコマンドを使用して、Debian パッケージのインストールをアンインストールします。
sudo apt-get パンテオンプを削除します
RHEL ファミリおよびその他の Linux システムでは、次のポータブル バンドルをインストールします。
sudo ./install.sh 。次のコマンドを使用してそのインストールを削除します。
sudo rm -f /usr/local/bin/pantheon && sudo rm -rf /opt/pantheongpu
もクリアしたい場合は、上記の完全削除コマンドを使用します。
現在のユーザーのコンパイル済みワークロード キャッシュ。
最初に実行されるワークロード ビルドは以下にキャッシュされます
${XDG_CACHE_HOME:-$HOME/.cache}/pantheongpu/builds/ 。
PANTHEON_BUILD_CACHE_DIR を設定して、別の書き込み可能なキャッシュ ディレクトリを選択します。

## Original Extract

Cross-platform CUDA and ROCm GPU stress testing, diagnostics, profiling, and performance benchmarks.

Hi HN, I built PantheonGPU because I wanted a better way to answer a simple question: is this GPU actually healthy and performing the way it should? A GPU can show normal temperatures and utilization and still be underperforming, unstable under certain workloads, or have memory, PCIe, or configuration issues. PantheonGPU actively tests the GPU instead of only monitoring telemetry. It currently includes 45+ tests covering compute, tensor workloads, memory, cache, PCIe, thermals, stability, and AI/LLM inference. It supports both NVIDIA CUDA and AMD ROCm. I’m also exploring a larger use case: running Pantheon across GPU fleets to identify individual GPUs that behave differently from the rest of a server or cluster. I’d especially appreciate feedback from people running AI infrastructure, multi-GPU systems, local LLMs, or GPU clouds.

PANTHEON
Skip to content
PANTHEON
Home
Initializing search
Home
Home
Home
Table of contents
Quick Start
1. Install prerequisites
Performance Database
Performance Database
Live Benchmarks
Research & Reports
Research & Reports
Overview
Test Documentation
Test Documentation
Overview
Core & Compute
Core & Compute
Omni Virus
Fixed-Function & Accelerators
Fixed-Function & Accelerators
RT Virus
Memory & Cache
Memory & Cache
Atomic Virus
Interconnect & Architecture
Interconnect & Architecture
P2P Thrasher
AI & ML
AI & ML
AI Workload Suites
Quick Start
1. Install prerequisites
GPU stress testing and diagnostics
Pantheon tests GPU compute, memory, cache, interconnect, and power behavior. Run focused workloads, capture telemetry, and keep the results for comparison.
The Debian package is the simplest installation path for Ubuntu and Debian systems.
Install the basic build tools:
sudo apt-get update
sudo apt-get install -y make g++
Then install the compiler for your GPU platform. You only need one:
sudo apt-get install -y nvidia-cuda-toolkit
sudo apt-get install -y hipcc
2. Install Pantheon
Download and install the latest Debian package:
VERSION = 1 .0.14
wget "https://github.com/saqibkh/pantheongpu_website/releases/download/v ${ VERSION } /pantheongpu_ ${ VERSION } _amd64.deb"
sudo apt install "./pantheongpu_ ${ VERSION } _amd64.deb"
To uninstall the Debian package later:
sudo apt-get remove pantheongpu
3. Verify the installation
Run a short hardware inventory test:
pantheon --test baseline_metrics --duration 10
Then run a targeted stress test on GPU 0:
pantheon --test fp64_virus --duration 30 --gpu 0
Note
Pantheon automatically detects CUDA, ROCm/HIP, or mock mode. Run the pantheon
command directly; you do not need to pass --platform cuda .
The native package command above removes Pantheon's package-managed files.
To also remove runtime-created files and the current user's compiled workload
cache, or to remove a portable installation on RHEL, Fedora, Rocky Linux,
AlmaLinux, or another Linux distribution, run:
curl -fsSL https://pantheongpu.com/uninstall.sh | sudo sh
This leaves CUDA, ROCm, system compilers, and benchmark reports stored outside
Pantheon's installation and cache directories untouched.
The release bundle contains the Debian package and an install.sh helper for
RHEL-family and other Linux distributions.
VERSION = 1 .0.14
wget "https://github.com/saqibkh/pantheongpu_website/releases/download/v ${ VERSION } /pantheongpu_ ${ VERSION } _amd64.tar.gz"
tar -xzf "pantheongpu_ ${ VERSION } _amd64.tar.gz"
cd "pantheongpu_ ${ VERSION } _amd64"
sudo apt install "./packages/pantheongpu_ ${ VERSION } _amd64.deb"
Uninstall a Debian package installation with:
sudo apt-get remove pantheongpu
On RHEL-family and other Linux systems, install the portable bundle with
sudo ./install.sh . Remove that installation with:
sudo rm -f /usr/local/bin/pantheon && sudo rm -rf /opt/pantheongpu
Use the complete-removal command above if you also want to clear the
current user's compiled workload cache.
First-run workload builds are cached under
${XDG_CACHE_HOME:-$HOME/.cache}/pantheongpu/builds/ .
Set PANTHEON_BUILD_CACHE_DIR to choose another writable cache directory.

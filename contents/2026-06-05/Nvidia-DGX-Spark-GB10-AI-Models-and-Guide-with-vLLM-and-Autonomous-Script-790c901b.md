---
source: "https://github.com/omnia-projetcs/spark-dgx"
hn_url: "https://news.ycombinator.com/item?id=48416815"
title: "Nvidia DGX Spark GB10 – AI Models and Guide with vLLM and Autonomous Script"
article_title: "GitHub - omnia-projetcs/spark-dgx · GitHub"
author: "nico248"
captured_at: "2026-06-05T21:37:11Z"
capture_tool: "hn-digest"
hn_id: 48416815
score: 1
comments: 0
posted_at: "2026-06-05T19:05:49Z"
tags:
  - hacker-news
  - translated
---

# Nvidia DGX Spark GB10 – AI Models and Guide with vLLM and Autonomous Script

- HN: [48416815](https://news.ycombinator.com/item?id=48416815)
- Source: [github.com](https://github.com/omnia-projetcs/spark-dgx)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T19:05:49Z

## Translation

タイトル: Nvidia DGX Spark GB10 – vLLM と自律スクリプトを使用した AI モデルとガイド
記事のタイトル: GitHub -omnia-projetcs/spark-dgx · GitHub
説明: GitHub でアカウントを作成して、omnia-projetcs/spark-dgx の開発に貢献します。

記事本文:
GitHub -omnia-projetcs/spark-dgx · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
オムニアプロジェクト
/
スパーク-dgx
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く Fo

フォルダとファイル
73 コミット 73 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md benchmark.py benchmark.py mix-vllm.sh mix-vllm.sh Prompts.txt Prompts.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
NVIDIA DGX Spark GB10 — AI モデルと推論ガイド
最先端の NVIDIA GB10 Grace Blackwell スーパーチップを搭載した NVIDIA DGX Spark デスクサイド スーパーコンピューターで最先端の AI モデルを実行、最適化、ベンチマークするための私のリポジトリおよびガイドへようこそ。
このガイドは、研究者、開発者、データ サイエンティストが LLM、VLM、拡散モデルを実行するときにローカル ワークステーションから最大のパフォーマンスを引き出すのに役立つように設計されています。
vLLM による高性能推論
DGX Spark で vLLM を使用する理由
Docker ファーストのアプローチ: パフォーマンスとクリーンさ
付属ツール
mix-vllm.sh — マルチモデルランチャー
benchmark.py — パフォーマンス テスト
役立つリソースと外部リンク
DGX Spark (GB10) のベスト プラクティス
NVIDIA DGX Spark は、革新的な NVIDIA GB10 Grace Blackwell スーパーチップを搭載したコンパクトなデスクサイド AI スーパーコンピューターです。これは、ローカルのプロトタイピングと大規模なデータセンターの拡張の間のギャップを橋渡しします。
プロセッサ: シングルチップ上の Coherent ARM64 ベースの NVIDIA Grace CPU + NVIDIA Blackwell GPU。
統合メモリ: 128 GB のコヒーレントな統合 LPDDR5x システム メモリは、高速 NVLink-C2C 経由で CPU と GPU 間でシームレスに共有されます。
パフォーマンス: ローカル AI パフォーマンスの最大 1 ペタフロップ (FP4 精度を使用)。
ネットワーキング: 統合された NVIDIA ConnectX-7 SmartNIC。
容量: ローカルで最大 2,000 億のパラメーターを実行するモデルをすぐに使用できるサポート。
マルチノード リンク: 2 台の DGX Spark ユニッ​​トをリンクして、最大 4,050 億のパラメーター モデルをサポートできます。
DGX Spark にモデルをローカルにデプロイする場合、パフォーマンスは精度 (FP16、FP8、FP4)、バッチによって異なります。

サイズ、およびテンソル並列処理の構成。
このようなワークステーション向けに特別に調整されたモデルのレイテンシ、スループット、トークン生成速度、タスク品質のリアルタイムおよびクラウドソースの評価については、公式ベンチマーク インデックスを確認してください。
👉 スパークアリーナリーダーボード
1 つ以上の NVIDIA DGX Spark システム上で LLM 推論ワークロードを (Slurm や Kubernetes の複雑さなしで) 簡単に起動、管理、調整するには、次の公式管理ツールを確認してください。
👉 スパークランのウェブサイト
TTFT (Time to First Token): インタラクティブなアプリケーション (チャットボットなど) にとって重要です。
Inter-Token Latency: 生成速度 (1 秒あたりのトークン数)。
スループット (トークン/秒/GPU): バッチ化されたオフライン処理ワークフローにとって重要です。
量子化品質の低下: 量子化されたバリアント (AWQ、GPTQ、FP8 など) をネイティブ FP16 ベースラインと比較します。
📈 パフォーマンスエクスペリエンスと同時実行レポート
DGX Spark ワークステーションに関するエンジニアリングの詳細および現実世界の同時実行ベンチマーク レポートについては、次のとおりです。
👉 Dendro Logic DGX Spark 同時実行ベンチマーク
vLLM による高性能推論
vLLM は、NVIDIA DGX Spark プラットフォーム上で LLM 推論を実行するための、絶対的に最もパフォーマンスが高く、高スループットでメモリ効率の高いエンジンの 1 つとして際立っています。 PagedAttendance を利用することで、vLLM は KV キャッシュを動的に管理し、メモリの断片化を事実上排除し、Spark の 128GB 統合 Grace Blackwell メモリを最大化できるようにします。
Docker ファーストのアプローチ: パフォーマンスとクリーンさ
ARM64 (aarch64)/Blackwell で vLLM を最大限に活用するには、Docker 内で vLLM を実行することが強く推奨されており、不可欠です。
ベアメタル インストールを避けて Docker を使用する必要がある理由:
最適なパフォーマンス: 事前に構築された ARM64 Docker イメージは、深く統合されたライブラリ (最適化された PyTorch、仕様) を使用してネイティブにコンパイルされます。

（Triton バージョン、FlashAttendant/FlashInfer バックエンド、およびカスタム CUDA カーネル）は、特に Blackwell のコンピューティング機能を対象としています。これらのコンパイルの最適化をホスト OS 上で手動で再現するのは非常に複雑です。
ゼロ ホスト ポリューション (「マシンへの依存」): ローカル ソースの構築では、大量のコンパイル ツールチェーン、カスタム Python パッケージ、複数の依存関係バージョン、および複雑なランタイム ライブラリ ( libuma-dev 、NCCL、特定の CUDA ツールキット パスの変更) がインストールされ、システムが乱雑になり、他の環境やシステム レベルの DGX OS ライブラリが破壊される可能性があります。 Docker は推論スタックを完全にカプセル化し、ホスト OS をクリーン、軽量、安定に保ちます。
NVIDIA NGC の公式の事前に最適化された ARM64 vLLM コンテナを使用するか、Grace Blackwell GB10 用に特別に構成された最適化されたコミュニティ構築セットアップを利用します。
👉 eugr/spark-vllm-docker : DGX Spark で vLLM を実行するための専用 Docker セットアップ。
👉 Saifgithub/vllm-gb10-sm121 : Grace Blackwell ( sm_121 ) ターゲット用に vLLM 構成が最適化されました。
OpenAI 互換 API サーバーの起動には、サンドボックス化された単一のコマンドのみが必要です。
docker run --gpus all \
--ipc=ホスト\
-v ~ /.cache/huggingface:/root/.cache/huggingface \
-p 8000:8000 \
vllm/vllm-openai:最新 \
vllmserve " metal-llama/Meta-Llama-3-8B-Instruct " \
--ポート 8000 \
--gpu メモリ使用率 0.90 \
--max-model-len 8192
Docker 化された vLLM サーバーにクエリを実行します。
カール http://localhost:8000/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"モデル": "メタ-ラマ/メタ-ラマ-3-8B-命令",
「メッセージ」: [
{"role": "user", "content": "NVIDIA Grace Blackwell NVLink-C2C 接続の利点について説明します。"}
]、
「温度」：0.7
} '
付属ツール
このリポジトリには、モデルのデプロイとベンチマークを行うための 2 つのすぐに使用できるスクリプトが同梱されています。

DGX Spark で。
mix-vllm.sh — マルチモデルランチャー
モデルごとに最適化された設定を使用して、完全に構成された本番環境対応の vLLM Docker コンテナを起動するターンキー Bash スクリプト。スクリプトを実行するだけで、任意のモデルを対話的に選択して起動できます。
⏳ 注: インターネット接続が良好な場合、最初の実行時にモデルをダウンロードしてキャッシュするのに平均 10 分かかります。
#
モデル
トク/秒
クオント
コンテキスト
能力
主な特長
★1
AEON-7/Qwen3.6-35B-A3B-heretic-NVFP4
~71 (88–117)
NVFP4
128K
💬🔧🧠
DFlash 仕様デコード、CUTLASS NVFP4、カスタム aeon-7 イメージ、FP8 KV キャッシュ
#2
rdtand/Qwen3.6-35B-A3B-PrismaQuant-4.75bit-vllm
~59
4.75ビット
256K
💬🔧🧠
投機的デコード (MTP × 3)、FP8 KV キャッシュ
#3
nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4
~58
NVFP4
256K
💬🔧🧠
MoE 30B/3.5B アクティブ、FlashInfer FP4
#4
bg-digitalservices/Gemma-4-26B-A4B-it-NVFP4
~50
NVFP4
262K
💬 🖼️ 🎥 🔊 🔧 🧠
マルチモーダル、TP×4、FP8 KVキャッシュ
—
Neural-ICE/Gemma-4-E2B-it-NVFP4
~120
NVFP4
128K
💬 🖼️ 🔊 🔧 🧠
Gemma 4 E2B Instruct NVFP4 — 高度に最適化された W4A4 会話モデル
—
bg-digitalservices/Gemma-4-E2B-NVFP4
~120
NVFP4
128K
💬🖼️🔊
Gemma 4 E2B ベース NVFP4 — W4A4 量子化事前トレーニング モデル
#5
クウェン/クウェン3.6-35B-A3B-FP8
~30
FP8
256K
💬🔧🧠
156 tok/s 集約 (c=32)、cu130-nightly
#6
rdtand/MiniMax-M2.7-PrismaQuant-3.20bit-vllm
~25
3.20ビット
196K
💬🔧🧠
MiniMax M2.7、PrismaQuant 3.20bit、標準 eugr イメージ
#7
インテル/Qwen3-Coder-Next-int4-AutoRound
~17
INT4
1M
💬🔧
MoE FP8、YaRN RoPE スケーリング、384 同時シーケンス
#8
RedHatAI/Qwen3.5-122B-A10B-NVFP4
~17
NVFP4
64K
💬🔧🧠
最高の品質 — RedHat キャリブレーション ≈ FP16、FlashInfer
—
シールドスター/Qwen3.5-122B-A10B-int4-AutoRound-EC
—
INT4
196K
💬🔧
オートラウンド I

NT4、z-lab DFlash 投機的デコード、カスタム vllm-node-tf5 イメージ
—
RedHatAI/Mistral-Small-24B-Instruct-2501-FP8-dynamic
~8.6 (合計 58.6)
FP8
32K
💬🔧
Mistral Small v4 — 高度に最適化され、非常に低い VRAM フットプリント (~24 GB)
#9
nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4
~15
NVFP4
128K
💬🔧🧠
MoE 120B/12B アクティブ、マーリン デカント
—
リキッドAI/LFM2.5-350M
～212
BF16
32K
💬
超軽量 350M、テスト/開発に最適
—
クウェン/クウェン3.5-0.8B
~103
BF16
102K
💬🔧
軽量 800M 高密度モデル、超高速、102k コンテキスト
—
nvidia/MiniMax-M2.7-NVFP4
~24
NVFP4
196K
💬🔧🧠
MiniMax M2.7、TP×4、インスタントテンサーフォーマット、FP8 KVキャッシュ
—
dervig/m51Lab-MiniMax-M2.7-REAP-139B-A10B-NVFP4-GB10
速い
NVFP4
102K
💬🔧🧠
ミニマックス M2.7
[切り捨てられた]
ハグフェイストークンが必要 — ほとんどのモデルでは、ウェイトをダウンロードするには有効なハグフェイスアクセストークンが必要です。これがないと、ゲートされたモデル (Llama、Gemma など) は起動できません。
ハグフェイス.co/join でアカウントを作成します
「設定」→「アクセストークン」に移動します：huggingface.co/settings/tokens
名前 (例: dgx-spark ) を選択し、権限「読み取り」を選択します (または、個人の名前空間にあるすべてのリポジトリのコンテンツに対する少なくとも読み取りアクセス権を持つきめ細かな権限)
トークンをコピーします ( hf_... で始まる)
# オプション A — スクリプトを直接編集します (行 ~40)
HUGGING_FACE_HUB_TOKEN= " hf_YourTokenHere "
# オプション B — 実行前に環境変数としてエクスポート
import HUGGING_FACE_HUB_TOKEN= " hf_YourTokenHere "
⚠️ ゲート付きモデル (ラマ、ジェマなど) の場合は、ダウンロードする前に、Hugging Face モデル ページでモデル ライセンスに同意する必要もあります。
📖 ドキュメント : ハグフェイス — セキュリティトークン
# 1. インタラクティブなモデル選択メニューを起動します。
./mix-vllm.sh
# または特定のモデルを直接起動します
./mix-vllm.sh --model nvidia/NVIDIA-Nemo

トロン-3-ナノ-30B-A3B-NVFP4
# モデルを起動し、デフォルトの Tensor Parallel サイズ (GPU の数) をオーバーライドします。
./mix-vllm.sh --model RedHatAI/Mistral-Small-24B-Instruct-2501-FP8-dynamic --tp 2
# サーバーの健全性をチェックする
カール http://localhost:8000/health
# ライブログを表示する
docker ログ -f mix-vllm
スクリプトは次のコマンドライン オプションをサポートしています。
--model <model_id> : 対話型選択メニューをバイパスして、特定のモデルを直接起動します。
--tp <tp_size> : 選択したモデルのデフォルトの Tensor Parallel サイズ (GPU の数) をオーバーライドします。省略した場合、スクリプトはモデル カタログで定義されている推奨される GPU 数を自動的にデフォルトに設定します。
--port <port> : vLLM サーバーがサービスを提供するポートをオーバーライドします (デフォルトは 8000 )。
--no-wait : サーバーが正常になるのを待たずに、コンテナーをバックグラウンドで実行します。
各モデル構成には、 --gpu-memory-utilization 、 --max-model-len 、 --max-num-batched-tokens 、アテンション バックエンド、量子化設定、およびツール呼び出しパーサーの最適化された値が含まれています。
ダウンロードしたモデルは、大量のディスク領域を占有する可能性があります。スクリプトは、モデルの構成に応じて、ダウンロードされた重みを 2 つの可能な場所に保存します。
ハグフェイスハブキャッシュ (デフォルト):
デフォルトでは、モデル ウェイトはダウンロードされ、ユーザーの標準の Hugging Face Hub キャッシュ ディレクトリにキャッシュされます。
~ /.cache/ハグフェイス/ハブ/
各モデルは、models--<author>--<model-name> のような名前のサブディレクトリに保存されます (ここで、sla

[切り捨てられた]

## Original Extract

Contribute to omnia-projetcs/spark-dgx development by creating an account on GitHub.

GitHub - omnia-projetcs/spark-dgx · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
omnia-projetcs
/
spark-dgx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
73 Commits 73 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md benchmark.py benchmark.py mix-vllm.sh mix-vllm.sh prompts.txt prompts.txt View all files Repository files navigation
NVIDIA DGX Spark GB10 — AI Models & Inference Guide
Welcome to my repository and guide for running, optimizing, and benchmarking state-of-the-art AI models on the NVIDIA DGX Spark deskside supercomputer, powered by the cutting-edge NVIDIA GB10 Grace Blackwell Superchip .
This guide is designed to help researchers, developers, and data scientists get the maximum performance out of their local workstation when running LLMs, VLMs, and diffusion models.
High-Performance Inference with vLLM
Why vLLM on DGX Spark?
Docker-First Approach: Performance & Cleanliness
Included Tools
mix-vllm.sh — Multi-Model Launcher
benchmark.py — Performance Testing
Useful Resources & External Links
Best Practices for DGX Spark (GB10)
The NVIDIA DGX Spark is a compact, deskside AI supercomputer featuring the revolutionary NVIDIA GB10 Grace Blackwell Superchip . It bridges the gap between local prototyping and heavy data center scaling.
Processor: Coherent ARM64-based NVIDIA Grace CPU + NVIDIA Blackwell GPU on a single chip.
Unified Memory: 128GB of coherent, unified LPDDR5x system memory shared seamlessly between CPU and GPU via high-speed NVLink-C2C.
Performance: Up to 1 PetaFLOP of local AI performance (using FP4 precision).
Networking: Integrated NVIDIA ConnectX-7 SmartNIC.
Capacity: Out-of-the-box support for running models up to 200 Billion parameters locally.
Multi-Node Link: Two DGX Spark units can be linked to support up to 405 Billion parameter models.
When deploying models locally on your DGX Spark, performance will vary depending on your precision (FP16, FP8, FP4), batch sizes, and tensor parallelism configurations.
For real-time and crowd-sourced evaluations of model latency, throughput, token generation speed, and task quality specifically calibrated for such workstations, check the official benchmark index:
👉 Spark Arena Leaderboard
To easily launch, manage, and orchestrate LLM inference workloads on one or more NVIDIA DGX Spark systems (without the complexity of Slurm or Kubernetes), check the official management tool:
👉 sparkrun Website
TTFT (Time to First Token): Critical for interactive applications (e.g., chatbots).
Inter-Token Latency: The generation speed (tokens per second).
Throughput (Tokens/sec/GPU): Important for batched, offline processing workflows.
Quantization Quality Degradation: Compares quantized variants (like AWQ, GPTQ, FP8) against native FP16 baselines.
📈 Performance Experience & Concurrency Reports
For an engineering deep-dive and real-world concurrency benchmark reports on the DGX Spark workstation:
👉 Dendro Logic DGX Spark Concurrency Benchmark
High-Performance Inference with vLLM
vLLM stands out as one of the absolute most performant, high-throughput, and memory-efficient engines to run LLM inference on the NVIDIA DGX Spark platform. By utilizing PagedAttention , vLLM dynamically manages the KV-cache, virtually eliminating memory fragmentation and allowing you to maximize the Spark’s 128GB unified Grace Blackwell memory.
Docker-First Approach: Performance & Cleanliness
To get the absolute best out of vLLM on ARM64 ( aarch64 ) / Blackwell, running vLLM within Docker is strongly recommended and essential .
Why you should avoid bare-metal installation and use Docker:
Optimal Performance: Pre-built ARM64 Docker images are natively compiled with deeply-integrated libraries (optimized PyTorch, specific Triton versions, FlashAttention/FlashInfer backends, and custom CUDA kernels) specifically targeted for Blackwell's computing capability. Manually reproducing these compiling optimizations on the host OS is extremely complex.
Zero Host Pollution ("Pourrissement de la Machine"): Local source building installs heavy compilation toolchains, custom Python packages, multiple dependency versions, and complex runtime libraries ( libnuma-dev , NCCL, specific CUDA Toolkit path modifications) that can clutter the system and break other environments or system-level DGX OS libraries. Docker completely encapsulates the inference stack, keeping the host OS clean, lightweight, and stable.
Use the official pre-optimized ARM64 vLLM containers from NVIDIA NGC, or utilize optimized community-built setups specifically configured for the Grace Blackwell GB10:
👉 eugr/spark-vllm-docker : Dedicated Docker setup for running vLLM on DGX Spark.
👉 saifgithub/vllm-gb10-sm121 : Optimized vLLM configurations for Grace Blackwell ( sm_121 ) target.
Launching the OpenAI-compatible API server takes only a single, sandboxed command:
docker run --gpus all \
--ipc=host \
-v ~ /.cache/huggingface:/root/.cache/huggingface \
-p 8000:8000 \
vllm/vllm-openai:latest \
vllm serve " meta-llama/Meta-Llama-3-8B-Instruct " \
--port 8000 \
--gpu-memory-utilization 0.90 \
--max-model-len 8192
Querying your dockerized vLLM Server:
curl http://localhost:8000/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "meta-llama/Meta-Llama-3-8B-Instruct",
"messages": [
{"role": "user", "content": "Explain the advantages of NVIDIA Grace Blackwell NVLink-C2C connection."}
],
"temperature": 0.7
} '
Included Tools
This repository ships with two ready-to-use scripts to deploy and benchmark models on your DGX Spark.
mix-vllm.sh — Multi-Model Launcher
A turnkey Bash script that launches a fully configured, production-ready vLLM Docker container with optimized per-model settings. Simply run the script to interactively select and launch any model.
⏳ Note : With a good internet connection, it takes on average 10 minutes to download and cache a model during its first run.
#
Model
tok/s
Quant
Context
Capabilities
Key Features
★1
AEON-7/Qwen3.6-35B-A3B-heretic-NVFP4
~71 (88–117)
NVFP4
128K
💬 🔧 🧠
DFlash spec-decode, CUTLASS NVFP4, custom aeon-7 image, FP8 KV-cache
#2
rdtand/Qwen3.6-35B-A3B-PrismaQuant-4.75bit-vllm
~59
4.75bit
256K
💬 🔧 🧠
Speculative decoding (MTP ×3), FP8 KV-cache
#3
nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4
~58
NVFP4
256K
💬 🔧 🧠
MoE 30B/3.5B active, FlashInfer FP4
#4
bg-digitalservices/Gemma-4-26B-A4B-it-NVFP4
~50
NVFP4
262K
💬 🖼️ 🎥 🔊 🔧 🧠
Multimodal, TP×4, FP8 KV-cache
—
Neural-ICE/Gemma-4-E2B-it-NVFP4
~120
NVFP4
128K
💬 🖼️ 🔊 🔧 🧠
Gemma 4 E2B Instruct NVFP4 — Highly optimized W4A4 conversational model
—
bg-digitalservices/Gemma-4-E2B-NVFP4
~120
NVFP4
128K
💬 🖼️ 🔊
Gemma 4 E2B Base NVFP4 — W4A4 quantized pre-trained model
#5
Qwen/Qwen3.6-35B-A3B-FP8
~30
FP8
256K
💬 🔧 🧠
156 tok/s aggregate (c=32), cu130-nightly
#6
rdtand/MiniMax-M2.7-PrismaQuant-3.20bit-vllm
~25
3.20bit
196K
💬 🔧 🧠
MiniMax M2.7, PrismaQuant 3.20bit, standard eugr image
#7
Intel/Qwen3-Coder-Next-int4-AutoRound
~17
INT4
1M
💬 🔧
MoE FP8, YaRN RoPE scaling, 384 concurrent sequences
#8
RedHatAI/Qwen3.5-122B-A10B-NVFP4
~17
NVFP4
64K
💬 🔧 🧠
Best quality — RedHat calibration ≈ FP16, FlashInfer
—
shieldstar/Qwen3.5-122B-A10B-int4-AutoRound-EC
—
INT4
196K
💬 🔧
AutoRound INT4, z-lab DFlash speculative decoding, custom vllm-node-tf5 image
—
RedHatAI/Mistral-Small-24B-Instruct-2501-FP8-dynamic
~8.6 (58.6 agg)
FP8
32K
💬 🔧
Mistral Small v4 — Highly optimized, extremely low VRAM footprint (~24 GB)
#9
nvidia/NVIDIA-Nemotron-3-Super-120B-A12B-NVFP4
~15
NVFP4
128K
💬 🔧 🧠
MoE 120B/12B active, Marlin dequant
—
LiquidAI/LFM2.5-350M
~212
BF16
32K
💬
Ultra-lightweight 350M, ideal for testing/development
—
Qwen/Qwen3.5-0.8B
~103
BF16
102K
💬 🔧
Lightweight 800M dense model, extremely fast, 102k context
—
nvidia/MiniMax-M2.7-NVFP4
~24
NVFP4
196K
💬 🔧 🧠
MiniMax M2.7, TP×4, instanttensor format, FP8 KV-cache
—
dervig/m51Lab-MiniMax-M2.7-REAP-139B-A10B-NVFP4-GB10
fast
NVFP4
102K
💬 🔧 🧠
MiniMax M2.7
[truncated]
Hugging Face Token Required — Most models need a valid Hugging Face access token to download weights. Without it, gated models (Llama, Gemma, etc.) will fail to start .
Create an account on huggingface.co/join
Go to Settings → Access Tokens : huggingface.co/settings/tokens
Choose a name (e.g. dgx-spark ) and select permission "Read" (or Fine-grained with at least Read access to contents of all repos under your personal namespace )
Copy the token (starts with hf_... )
# Option A — Edit the script directly (line ~40)
HUGGING_FACE_HUB_TOKEN= " hf_YourTokenHere "
# Option B — Export as environment variable before running
export HUGGING_FACE_HUB_TOKEN= " hf_YourTokenHere "
⚠️ For gated models (Llama, Gemma, etc.), you must also accept the model license on its Hugging Face model page before downloading.
📖 Documentation : Hugging Face — Security Tokens
# 1. Launch the interactive model selection menu
./mix-vllm.sh
# Or launch a specific model directly
./mix-vllm.sh --model nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-NVFP4
# Launch a model and override the default Tensor Parallel size (number of GPUs)
./mix-vllm.sh --model RedHatAI/Mistral-Small-24B-Instruct-2501-FP8-dynamic --tp 2
# Check server health
curl http://localhost:8000/health
# View live logs
docker logs -f mix-vllm
The script supports the following command-line options:
--model <model_id> : Launches a specific model directly, bypassing the interactive selection menu.
--tp <tp_size> : Overrides the default Tensor Parallel size (number of GPUs) for the selected model. If omitted, the script automatically defaults to the recommended number of GPUs defined in the model catalog.
--port <port> : Overrides the port on which the vLLM server is served (default is 8000 ).
--no-wait : Runs the container in the background without waiting for the server to become healthy.
Each model configuration includes optimized values for --gpu-memory-utilization , --max-model-len , --max-num-batched-tokens , attention backends, quantization settings, and tool-call parsers.
Downloaded models can take up significant disk space. The script stores downloaded weights in two possible locations depending on the model's configuration:
Hugging Face Hub Cache (Default) :
By default, model weights are downloaded and cached in your user's standard Hugging Face Hub cache directory:
~ /.cache/huggingface/hub/
Each model is stored in a subdirectory named like models--<author>--<model-name> (where sla

[truncated]

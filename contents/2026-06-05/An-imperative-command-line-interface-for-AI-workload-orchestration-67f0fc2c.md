---
source: "https://pypi.org/project/terradev-cli/"
hn_url: "https://news.ycombinator.com/item?id=48416931"
title: "An imperative command-line-interface for AI workload orchestration"
article_title: "terradev-cli · PyPI"
author: "Facingsouth"
captured_at: "2026-06-05T21:36:38Z"
capture_tool: "hn-digest"
hn_id: 48416931
score: 1
comments: 1
posted_at: "2026-06-05T19:15:30Z"
tags:
  - hacker-news
  - translated
---

# An imperative command-line-interface for AI workload orchestration

- HN: [48416931](https://news.ycombinator.com/item?id=48416931)
- Source: [pypi.org](https://pypi.org/project/terradev-cli/)
- Score: 1
- Comments: 1
- Posted: 2026-06-05T19:15:30Z

## Translation

タイトル: AI ワークロード オーケストレーションのための必須コマンドライン インターフェイス
記事のタイトル: terradev-cli · PyPI
説明: AI ワークロード オーケストレーションのための必須コマンドライン インターフェイス。

記事本文:
メインコンテンツにスキップ
モバイル版に切り替える
警告
サポートされていないブラウザを使用しているため、新しいバージョンにアップグレードしてください。
PyPIを検索
検索フォーカス#フォーカス検索フィールド"
data-search-focus-target="検索フィールド">
検索
ヘルプ
pip インストール terradev-cli
PIP 命令のコピー
AI ワークロード オーケストレーションのための必須コマンドライン インターフェイス。
ライセンス: Apache ソフトウェア ライセンス (Apache-2.0)
タグ
雲
、
プロビジョニング
、
GPU
、
マルチクラウド
、
裁定取引
、
AWS
、
gcp
、
紺碧
、
ランポッド
、
ヴァスタイ
、
機械学習
、
コストの最適化
、
平行
、
ギトプス
、
argocd
、
フラックス
、
Kubernetes
、
コードとしてのインフラストラクチャ
、
ラングチェーン
、
ランググラフ
、
俗語
、
ワンドブ
、
ミリフロー
、
分析
、
セットアップ
、
変換
、
テレメトリー
、
リアルタイム
、
モニタリング
、
ダッシュボード
、
抱きしめる顔
、
ハグフェイススペース
、
グラディオ
、
流光に照らされた
、
llm
、
変圧器
、
パイトーチ
、
文変換者
、
スマートテンプレート
、
ハードウェアの最適化
、
モデルの展開
、
推論
、
トレーニング
、
ビョアピ
、
テラフォーム
、
Kubernetes
、
ヘルム
、
大工
、
オパ
、
プロメテウス
、
グラファナ
、
光線
、
vllm
、
オラマ
、
DV
、
データのバージョン管理
、
沼
、
分解推論
、
萌える
、
専門家の混合
、
クウェン
、
分散する
、
xcd
、
mi300x
、
h200
、
プレフィルデコード
、
cudaグラフ
、
cuda グラフ
、
沼トポロジー
、
GPUの最適化
、
パッシブ最適化
、
温水プール
、
エンドポイントスコアリング
追加の提供:
AWS
、gcp
、紺碧
、lmキャッシュ
、ml
、請求
、すべて
、hf
、開発者
、エンタープライズ
開発状況
5 - 生産/安定
ライセンス
OSI 承認済み :: Apache ソフトウェア ライセンス
オペレーティングシステム
OSに依存しない
プログラミング言語
パイソン :: 3
トピック
インターネット :: WWW/HTTP :: 動的コンテンツ
ソフトウェア開発 :: ライブラリ :: Python モジュール
システム :: システム管理
AI ワークロード オーケストレーションのための必須コマンドライン インターフェイス。
ライセンス: Apache 2.0 - 商用および許可に基づく無料のオープンソース

ソニーの使用。
Terradev は、データセットの圧縮とステージング、最適なインスタンスとノードのプロビジョニングを行い、順次プロビジョニングよりも 3 ～ 5 倍高速にデプロイするクロスクラウド コンピューティング プロビジョニング CLI です。
ペイウォールを削除し、Terradev をオープンソース化し、安全かつ迅速な配信のために Rust アクセラレータを追加しました...
Rust DAG オーケストレーターを使用すると、実行グラフはランタイム レベルで正しい順序付けと冪等性を強制します。あなたまたはエージェントは自由にコマンドを発行できます。オーケストレーターはコマンドが安全に実行できることを保証します。
サブコマンド/フラグを含まない 104 ツールには、重いコンテキストが必要です。 Rust MCP オーケストレーターは、最小限のオーバーヘッドでツール呼び出しを処理します。デシリアライズ、ルーティング、実行、応答は、純粋な Python ベースの MCP サーバーよりも桁違いに高速です。エージェントが 21 を超えるクラウド プロバイダーにわたって複雑なプロビジョニング ワークフローを実行している場合、それはチェーン内のすべてのツール呼び出しにわたって複雑になります。
API キーはローカルの ~/.terradev/credentials.json に保存され、Terradev サーバーに送信されることはありません。
# 複数のプロバイダーを構成する
terradevconfigure --provider runpod
terradevconfigure --provider Vastai
terradevconfigure --provider aws
terradevconfigure --provider gcp
パフォーマンス
vLLM 最適化によりスループットが 2 ～ 8 倍向上
NUMA トポロジにより 30 ～ 50% の帯域幅ペナルティが排除される
最適なトポロジーによる 2 ～ 5 倍の CUDA グラフの高速化
自動プロバイダー切り替えにより最大 90% のコスト削減
KV キャッシュ チェックポイントによるスポット リカバリは 2 分未満
ウェイト ストリーミングによりコールド スタートが 3.6 倍高速化
MLA を認識した VRAM 推定により 57% のコスト削減
pip インストール terradev-cli
すべてのクラウド プロバイダー SDK と ML 統合の場合:
pip インストール terradev-cli [すべて]
コマンドを確認してリストします。
terradev --ヘルプ
ステップ 2: 最初のクラウドプロバイダーを構成する
Terradev は 21 を超える GPU クラウド プロバイダーをサポートしています。まずは 1 つから始めましょう。RunPod は 1 つです

彼は最も早くセットアップできました:
terradev セットアップ runpod --quick
これは、API キーを取得する場所を示しています。次に、それを設定します。
terradevconfigure --provider runpod
プロンプトが表示されたら、API キーを貼り付けます。これはローカルの ~/.terradev/credentials.json に保存され、Terradev サーバーに送信されることはありません。後でプロバイダーを追加します。
terradevconfigure --provider Vastai
terradevconfigure --provider lambda_labs
terradevconfigure --provider aws
構成するプロバイダーの数が増えるほど、料金範囲が広がります。
ステップ 3: リアルタイムの GPU 価格を取得する
構成したすべてのプロバイダーの料金を確認します。
terradev 引用 -g A100
出力は、時間あたりの価格、プロバイダー、地域、スポット対オンデマンドの安い順にソートされたテーブルです。さまざまな GPU を試してください。
terradev quote -g H100
terradev quote -g L40S
terradev quote -g RTX4090
ステップ 4: プロビジョニング
ほとんどのクラウドは、デフォルトで次善のトポロジを備えた GPU を提供します。 GPU と NIC は異なる NUMA ノードになり、RDMA は無効になり、kubelet トポロジ マネージャーは none に設定されます。これは、すべての分散操作で 30 ～ 50% の帯域幅ペナルティが発生しますが、nvidia-smi では決して発生しません。
Terradev を通じてプロビジョニングすると、トポロジの最適化が自動的に行われます。
terradev のプロビジョニング -g H100 -n 4 --Parallel 6
舞台裏で何が起こっているのか:
NUMA アライメント — GPU と NIC を同じ NUMA ノードに強制する
GPUDirect RDMA — nvidia_peermem ロード、ゼロコピー GPU 間転送
CPU 固定 — 静的 CPU マネージャー ポリシー、コア移行なし
SR-IOV — 分離された RDMA パスの GPU ごとに作成される仮想関数
NCCL チューニング - InfiniBand 有効、GDR_LEVEL=PIX、GDR_READ=1
これについては何も設定しません。自動的に適用されます。
プランを起動せずにプレビューするには:
terradev Provision -g A100 -n 2 --dry-run
価格上限を設定するには:
terradev Provision -g A100 --max-price 2 .50
ステップ 5: ワークロードを実行する
オプション A — co を実行する

プロビジョニングされたインスタンスでコマンドを実行します。
terradev 実行 -i <インスタンス ID> -c "nvidia-smi"
terradev 実行 -i <インスタンス ID> -c "python train.py"
オプション B — コンテナーをプロビジョニング、デプロイし、実行する 1 つのコマンド:
terradev run --gpu A100 --image pytorch/pytorch:latest -c "python train.py"
オプション C — 推論サーバーを生きたままにします。
terradev run --gpu H100 --image vllm/vllm-openai:latest --keep-alive --port 8000
ステップ 6: インスタンスを管理する
# 実行中のすべてのインスタンスと現在のコストを表示
terradev ステータス --live
# 停止 (割り当てを維持)
terradev manage -i <インスタンス ID> -a stop
# 再起動
terradev manage -i <インスタンス ID> -a start
# 終了して解放する
terradev manage -i <インスタンス ID> -a terminate
ステップ 7: コストを追跡し、節約額を見つける
# 過去 30 日間の支出を表示
terradev 分析 -- 30 日目
# インスタンスを実行するためのより安価な代替手段を見つける
terradev の最適化
ステップ 8: 分散トレーニング パイプライン
ノードに正しいトポロジが設定されたため、分散トレーニングは実際に全帯域幅で実行されます。
# 起動前に GPU、NCCL、RDMA、ドライバーを検証する
terradev プリフライト
# プロビジョニングしたばかりのノードでトレーニングを開始します
terradev train --script train.py --from-provision 最新
# GPU の使用率とコストをリアルタイムで監視する
terradev モニター --job my-job
# ステータスを確認する
terradev トレインステータス
#6. 完了したらチェックポイントをリストする
terradev チェックポイント リスト --job my-job
--from-provision 最新フラグは、最後のプロビジョニング コマンドから IP を自動解決します。 torchrun、DeepSpeed、Accelerate、Megatron をサポートします。
ステップ 9: vLLM 推論の最適化 (6 つのノブ)
vLLM を使用してモデルを提供している場合、ほとんどのチームがデフォルトのままにしている設定が 6 つあり、それぞれにスループットがかかります。
ワークロード プロファイルから 6 つすべてを自動調整します。
terradev vllm auto-optimize -s workload.json -m metal-llama/Llama-2-7b-hf

-g4
または、実行中のサーバーを分析します。
terradev vllm 分析 -e http://localhost:8000
ベンチマーク:
terradev vllm ベンチマーク -e http://localhost:8000 -c 10
ステップ 10: 自動適用された最適化を使用して MoE モデルを展開する
大規模な専門家混合モデル (GLM-5、Qwen 3.5、DeepSeek V4) の場合、Terradev の MoE テンプレートには、KV キャッシュ オフロード、投機的デコード、スリープ モード、エキスパート ロード バランシングなど、自動適用されるすべての最適化が含まれています。
terradev プロビジョニング --task クラスター/moe-template/task.yaml \
--set model_id = Qwen/Qwen3.5-397B-A17B
または、より小さいモデル:
terradev プロビジョニング --task クラスター/moe-template/task.yaml \
--set model_id = Qwen/Qwen3.5-122B-A10B --set tp_size = 4 --set gpu_count = 4
自動適用されるもの (フラグは必要ありません):
KV キャッシュのオフロード - CPU DRAM へのスピル、最大 9 倍のスループット
MTP 投機的デコード - 最大 2.8 倍高速な生成
スリープ モード - アイドル モデルは CPU RAM に休止状態になり、コールド リスタートより 18 ～ 200 倍高速になります
エキスパートの負荷分散 - 実行時にルーティングのバランスを再調整します
LMCache — Redis を介して KV キャッシュをインスタンス間で分散します
ステップ 11: 細分化されたプレフィル/デコード (上級)
これにより、推論が各フェーズに最適化された 2 つの GPU プールに分割されます。
プレフィル (コンピューティング バウンド) — 入力プロンプトを処理し、高い FLOPS を必要とします
デコード (メモリバウンド) — トークンを生成し、高い HBM 帯域幅が必要です
KV キャッシュは、NIXL (RDMA を介したゼロコピー GPU 間) を介してそれらの間で転送されます。これが、ステップ 4 で NUMA トポロジを正しく取得することが重要である理由です。NIXL は、GPU と NIC が PCIe スイッチを共有している場合にのみフルスピードで実行されます。
terradev ml ray --deploy-pd \
--model zai-org/GLM-5-FP8 \
--prefill-tp 8 --decode-tp 1 --decode-dp 24
Terradev の推論ルーターは自動的にスティッキー ルーティングを使用します。プリフィル GPU が KV キャッシュをデコード GPU に渡すと、同じプレフィックスを持つ今後のリクエストは同じデコード GPU に送られます。

冗長な転送を防ぎます。
ステップ 12: Kubernetes GPU クラスターを作成する
運用環境では、トポロジが最適化された K8s クラスターを作成します。
terradev k8s create my-cluster --gpu H100 --count 8 --prefer-spot
これにより、NUMA に調整された kubelet トポロジ マネージャー、GPUDirect RDMA、および PCIe ローカリティ強制を使用して Karpenter NodePools が自動構成されます。
# クラスターをリストする
terradev k8s リスト
# クラスター情報を取得する
terradev k8s 情報 my-cluster
# 解体する
terradev k8s が私のクラスターを破壊します
この順序が重要な理由
各ステップは、その前のステップに基づいて構築されます。
ステップ 4 : NUMA / RDMA / SR-IOV トポロジ ← 基礎
ステップ 8 : フル BW での分散トレーニング ← トポロジーに依存
ステップ 9 : vLLM ノブの調整 ← 正しいメモリ レイアウトに依存します
ステップ 10 : KV キャッシュのオフロード + スリープ モード ← CPU バスが飽和していないことに依存します
ステップ 11 : 非集約 P/D ← KV 転送は RDMA に依存
プロビジョニング層が間違っていると、その上にあるすべての最適化のパフォーマンスが低下します。クロス NUMA KV 転送を使用する非集約 P/D セットアップは、正しいトポロジーのモノリシック セットアップよりも遅くなります。
Terradev は基礎を自動的に処理するため、スタックの残りの部分は想定どおりに動作します。
例 1: LLM 推論サービス
#!/bin/bash
# LLM 導入ワークフローを完了する
#1. 最も安価な GPU を見つける
terradev quote -g A100 --quick
# 2. 自動最適化を使用したプロビジョニング
terradev のプロビジョニング -g A100 -n 2 --Parallel 4
# 3. 最適化された vLLM をデプロイする
terradev ml vllm --start --instance-ip $( terradev status --json | jq -r '.[0].ip' ) --model metal-llama/Llama-2-7b-hf --tp-size 2
#4. モニタリングを設定する
terradev モニター --endpoint llama-api --live
#5. カスタマーアダプターを追加する
terradev lora add -e http:// $( terradev status --json | jq -r '.[0].ip' ) :8000 -n customer-a -p ./adapters/customer-a
例 2: MoE モデルの実稼働展開
#!/bin/bash
# GLM-5 量産展開

イメント
# 1. MoE クラスターを展開する
terradev プロビジョニング --task Clusters/moe-template/task.yaml --set model_id = zai-org/GLM-5-FP8 --set tp_size = 8
#2. モニタリングを展開する
terradev k8s モニタリングスタック --cluster glm-5-cluster
# 3. バーストトラフィック用にウォームプールを設定する
terradev ml Warm-pool --configure --strategy Traffic_based --max-warm-models 5 --endpoint glm-5-api
#4. フェイルオーバーのテスト
terradev inferx フェイルオーバー --endpoint glm-5-api --test-load 5000
例 3: InferX + LoRA ハイブリッド展開 (実稼働マルチテナント)
#!/bin/bash
# コールド スタート フェイルオーバーとマルチテナント LoRA アダプターを使用した実稼働デプロイメント
echo "InferX + LoRA ハイブリッド推論サービスのデプロイ"
#1. 安定したトラフィックのためにベースラインの予約済み GPU を導入する
echo " ステップ 1: 予約されたベースライン容量をプロビジョニングする"
terradev のプロビジョニング -g H100 -n 2 --Parallel 4 \
--tag ベースライン-llm \
--最大価格 2.50
BASELINE_IP = $( terradev status --json | jq -r '.[] | select(.tags[] | contains("baseline-llm")) | .ip' | head -1 )
# 2. LoRA サポートを備えた最適化された vLLM をベースラインで導入する
echo " ステップ 2: LoRA アダプターをサポートする vLLM のデプロイ"
terradev ml vllm --start \
--instance-ip $BASELINE_IP \
--model metal-llama/Llama-2-7b-hf \
--tp-サイズ 2 \
--enable-lora \
--enable-kv-offloading \
--スリープモードを有効にする \
--ポート 8000

[切り捨てられた]

## Original Extract

An imperative command-line-interface for AI workload orchestration.

Skip to main content
Switch to mobile version
Warning
You are using an unsupported browser, upgrade to a newer version.
Search PyPI
search-focus#focusSearchField"
data-search-focus-target="searchField">
Search
Help
pip install terradev-cli
Copy PIP instructions
An imperative command-line-interface for AI workload orchestration.
License: Apache Software License (Apache-2.0)
Tags
cloud
,
provisioning
,
gpu
,
multi-cloud
,
arbitrage
,
aws
,
gcp
,
azure
,
runpod
,
vastai
,
machine-learning
,
cost-optimization
,
parallel
,
gitops
,
argocd
,
flux
,
kubernetes
,
infrastructure-as-code
,
langchain
,
langgraph
,
sglang
,
wandb
,
mlflow
,
analytics
,
setup
,
conversion
,
telemetry
,
real-time
,
monitoring
,
dashboard
,
huggingface
,
huggingface-spaces
,
gradio
,
streamlit
,
llm
,
transformers
,
pytorch
,
sentence-transformers
,
smart-templates
,
hardware-optimization
,
model-deployment
,
inference
,
training
,
byoapi
,
terraform
,
kubernetes
,
helm
,
karpenter
,
opa
,
prometheus
,
grafana
,
ray
,
vllm
,
ollama
,
dvc
,
data-version-control
,
numa
,
disaggregated-inference
,
moe
,
mixture-of-experts
,
qwen
,
distserve
,
xcd
,
mi300x
,
h200
,
prefill-decode
,
cuda-graph
,
cuda-graphs
,
numa-topology
,
gpu-optimization
,
passive-optimization
,
warm-pool
,
endpoint-scoring
Provides-Extra:
aws
, gcp
, azure
, lmcache
, ml
, billing
, all
, hf
, dev
, enterprise
Development Status
5 - Production/Stable
License
OSI Approved :: Apache Software License
Operating System
OS Independent
Programming Language
Python :: 3
Topic
Internet :: WWW/HTTP :: Dynamic Content
Software Development :: Libraries :: Python Modules
System :: Systems Administration
An imperative command-line-interface for AI workload orchestration.
License: Apache 2.0 - Free and open source for commercial and personal use.
Terradev is a cross-cloud compute-provisioning CLI that compresses + stages datasets, provisions optimal instances + nodes, and deploys 3-5x faster than sequential provisioning.
We removed the paywall, open-sourced Terradev, and added Rust accelerators for safe and snappy delivery...
With the Rust DAG orchestrator, the execution graph enforces correct sequencing and idempotency at the runtime level. You or the agent can issue commands freely... the orchestrator ensures they're safe to execute.
104 tools not including subcommand/flags require heavy context. The Rust MCP orchestrator processes tool calls with minimal overhead: deserializing, routing, executing, and responding faster than pure-Python-based MCP servers by an order of magnitude. For an agent running a complex provisioning workflow across 21+ cloud providers, that compounds across every tool call in the chain.
Your API keys are stored locally at ~/.terradev/credentials.json and never sent to Terradev servers.
# Configure multiple providers
terradev configure --provider runpod
terradev configure --provider vastai
terradev configure --provider aws
terradev configure --provider gcp
Performance
2-8x throughput improvements with vLLM optimization
30-50% bandwidth penalty eliminated with NUMA topology
2-5x CUDA Graph speedup with optimal topology
Up to 90% cost savings with automatic provider switching
<2 minute spot recovery with KV cache checkpointing
3.6x faster cold starts with weight streaming
57% cost savings with MLA-aware VRAM estimation
pip install terradev-cli
For all cloud provider SDKs and ML integrations:
pip install terradev-cli [ all ]
Verify and list commands:
terradev --help
Step 2: Configure Your First Cloud Provider
Terradev supports 21+ GPU cloud providers. Start with one, RunPod is the fastest to set up:
terradev setup runpod --quick
This shows you where to get your API key. Then configure it:
terradev configure --provider runpod
Paste your API key when prompted. It's stored locally at ~/.terradev/credentials.json, never sent to a Terradev server. Add more providers later:
terradev configure --provider vastai
terradev configure --provider lambda_labs
terradev configure --provider aws
The more providers you configure, the better your price coverage.
Step 3: Get Real-Time GPU Prices
Check pricing across every provider you've configured:
terradev quote -g A100
Output is a table sorted cheapest-first: price/hour, provider, region, spot vs. on-demand. Try different GPUs:
terradev quote -g H100
terradev quote -g L40S
terradev quote -g RTX4090
Step 4: Provision
Most clouds hand you GPUs with suboptimal topology by default. Your GPU and NIC end up on different NUMA nodes, RDMA is disabled, and the kubelet Topology Manager is set to none. That's a 30-50% bandwidth penalty on every distributed operation and you'll never see it in nvidia-smi.
When you provision through Terradev, topology optimization is automatic:
terradev provision -g H100 -n 4 --parallel 6
What happens behind the scenes:
NUMA alignment — GPU and NIC forced to the same NUMA node
GPUDirect RDMA — nvidia_peermem loaded, zero-copy GPU-to-GPU transfers
CPU pinning — static CPU manager policy, no core migration
SR-IOV — virtual functions created per GPU for isolated RDMA paths
NCCL tuning — InfiniBand enabled, GDR_LEVEL=PIX, GDR_READ=1
You don't configure any of this. It's applied automatically.
To preview the plan without launching:
terradev provision -g A100 -n 2 --dry-run
To set a price ceiling:
terradev provision -g A100 --max-price 2 .50
Step 5: Run a Workload
Option A — Run a command on your provisioned instance:
terradev execute -i <instance-id> -c "nvidia-smi"
terradev execute -i <instance-id> -c "python train.py"
Option B — One command that provisions, deploys a container, and runs:
terradev run --gpu A100 --image pytorch/pytorch:latest -c "python train.py"
Option C — Keep an inference server alive:
terradev run --gpu H100 --image vllm/vllm-openai:latest --keep-alive --port 8000
Step 6: Manage Your Instances
# See all running instances and current cost
terradev status --live
# Stop (keeps allocation)
terradev manage -i <instance-id> -a stop
# Restart
terradev manage -i <instance-id> -a start
# Terminate and release
terradev manage -i <instance-id> -a terminate
Step 7: Track Costs and Find Savings
# View spend over the last 30 days
terradev analytics --days 30
# Find cheaper alternatives for running instances
terradev optimize
Step 8: Distributed Training Pipeline
Now that your nodes have correct topology, distributed training actually runs at full bandwidth:
# Validate GPUs, NCCL, RDMA, and drivers before launching
terradev preflight
# Launch training on the nodes you just provisioned
terradev train --script train.py --from-provision latest
# Watch GPU utilization and cost in real time
terradev monitor --job my-job
# Check status
terradev train-status
# 6. List checkpoints when done
terradev checkpoint list --job my-job
The --from-provision latest flag auto-resolves IPs from your last provision command. Supports torchrun, DeepSpeed, Accelerate, and Megatron.
Step 9: Optimize vLLM Inference (The 6 Knobs)
If you're serving a model with vLLM, there are 6 settings most teams leave at defaults — each one costs throughput:
Auto-tune all six from your workload profile:
terradev vllm auto-optimize -s workload.json -m meta-llama/Llama-2-7b-hf -g 4
Or analyze a running server:
terradev vllm analyze -e http://localhost:8000
Benchmark:
terradev vllm benchmark -e http://localhost:8000 -c 10
Step 10: Deploy a MoE Model with Auto-Applied Optimizations
For large Mixture-of-Experts models (GLM-5, Qwen 3.5, DeepSeek V4), Terradev's MoE templates include every optimization auto-applied — KV cache offloading, speculative decoding, sleep mode, expert load balancing:
terradev provision --task clusters/moe-template/task.yaml \
--set model_id = Qwen/Qwen3.5-397B-A17B
Or a smaller model:
terradev provision --task clusters/moe-template/task.yaml \
--set model_id = Qwen/Qwen3.5-122B-A10B --set tp_size = 4 --set gpu_count = 4
What's auto-applied (no flags needed):
KV cache offloading — spills to CPU DRAM, up to 9x throughput
MTP speculative decoding — up to 2.8x faster generation
Sleep mode — idle models hibernate to CPU RAM, 18-200x faster than cold restart
Expert load balancing — rebalances routing at runtime
LMCache — distributes KV cache across instances via Redis
Step 11: Disaggregated Prefill/Decode (Advanced)
This separates inference into two GPU pools optimized for each phase:
Prefill (compute-bound) — processes input prompt, wants high FLOPS
Decode (memory-bound) — generates tokens, wants high HBM bandwidth
The KV cache transfers between them via NIXL — zero-copy GPU-to-GPU over RDMA. This is why getting the NUMA topology right in Step 4 matters: NIXL only runs at full speed when the GPU and NIC share a PCIe switch.
terradev ml ray --deploy-pd \
--model zai-org/GLM-5-FP8 \
--prefill-tp 8 --decode-tp 1 --decode-dp 24
Terradev's inference router automatically uses sticky routing. Once a prefill GPU hands off a KV cache to a decode GPU, future requests with the same prefix go to that same decode GPU, avoiding redundant transfers.
Step 12: Create a Kubernetes GPU Cluster
For production, create a topology-optimized K8s cluster:
terradev k8s create my-cluster --gpu H100 --count 8 --prefer-spot
This auto-configures Karpenter NodePools with NUMA-aligned kubelet Topology Manager, GPUDirect RDMA, and PCIe locality enforcement.
# List clusters
terradev k8s list
# Get cluster info
terradev k8s info my-cluster
# Tear down
terradev k8s destroy my-cluster
Why This Order Matters
Each step builds on the one before it:
Step 4 : NUMA / RDMA / SR-IOV topology ← foundation
Step 8 : Distributed training at full BW ← depends on topology
Step 9 : vLLM knob tuning ← depends on correct memory layout
Step 10 : KV cache offloading + sleep mode ← depends on CPU bus not saturated
Step 11 : Disaggregated P/D ← depends on RDMA for KV transfer
If the provisioning layer is wrong, every optimization above it underperforms. A disaggregated P/D setup with a cross-NUMA KV transfer is slower than a monolithic setup with correct topology.
Terradev handles the foundation automatically so the rest of the stack works the way it's supposed to.
Example 1: LLM Inference Service
#!/bin/bash
# Complete LLM deployment workflow
# 1. Find cheapest GPU
terradev quote -g A100 --quick
# 2. Provision with auto-optimization
terradev provision -g A100 -n 2 --parallel 4
# 3. Deploy optimized vLLM
terradev ml vllm --start --instance-ip $( terradev status --json | jq -r '.[0].ip' ) --model meta-llama/Llama-2-7b-hf --tp-size 2
# 4. Set up monitoring
terradev monitor --endpoint llama-api --live
# 5. Add customer adapter
terradev lora add -e http:// $( terradev status --json | jq -r '.[0].ip' ) :8000 -n customer-a -p ./adapters/customer-a
Example 2: MoE Model Production Deployment
#!/bin/bash
# GLM-5 production deployment
# 1. Deploy MoE cluster
terradev provision --task clusters/moe-template/task.yaml --set model_id = zai-org/GLM-5-FP8 --set tp_size = 8
# 2. Deploy monitoring
terradev k8s monitoring-stack --cluster glm-5-cluster
# 3. Set up warm pool for bursty traffic
terradev ml warm-pool --configure --strategy traffic_based --max-warm-models 5 --endpoint glm-5-api
# 4. Test failover
terradev inferx failover --endpoint glm-5-api --test-load 5000
Example 3: InferX + LoRA Hybrid Deployment (Production Multi-Tenant)
#!/bin/bash
# Production deployment with cold start failover and multi-tenant LoRA adapters
echo " Deploying InferX + LoRA Hybrid Inference Service"
# 1. Deploy baseline reserved GPUs for steady traffic
echo " Step 1: Provision reserved baseline capacity"
terradev provision -g H100 -n 2 --parallel 4 \
--tag baseline-llm \
--max-price 2 .50
BASELINE_IP = $( terradev status --json | jq -r '.[] | select(.tags[] | contains("baseline-llm")) | .ip' | head -1 )
# 2. Deploy optimized vLLM with LoRA support on baseline
echo " Step 2: Deploy vLLM with LoRA adapter support"
terradev ml vllm --start \
--instance-ip $BASELINE_IP \
--model meta-llama/Llama-2-7b-hf \
--tp-size 2 \
--enable-lora \
--enable-kv-offloading \
--enable-sleep-mode \
--port 8000

[truncated]

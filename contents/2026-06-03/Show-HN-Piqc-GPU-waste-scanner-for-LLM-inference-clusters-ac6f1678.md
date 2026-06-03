---
source: "https://github.com/paralleliq/piqc"
hn_url: "https://news.ycombinator.com/item?id=48371299"
title: "Show HN: Piqc – GPU waste scanner for LLM inference clusters"
article_title: "GitHub - paralleliq/piqc: Kubernetes scanner that discovers LLMs running on vLLM and extracts their deployment and runtime facts. · GitHub"
author: "paralleliq"
captured_at: "2026-06-03T00:42:43Z"
capture_tool: "hn-digest"
hn_id: 48371299
score: 3
comments: 0
posted_at: "2026-06-02T15:10:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Piqc – GPU waste scanner for LLM inference clusters

- HN: [48371299](https://news.ycombinator.com/item?id=48371299)
- Source: [github.com](https://github.com/paralleliq/piqc)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T15:10:41Z

## Translation

タイトル: Show HN: Piqc – LLM 推論クラスター用の GPU 廃棄物スキャナー
記事のタイトル: GitHub -Paralleliq/piqc: vLLM 上で実行されている LLM を検出し、そのデプロイメントと実行時の事実を抽出する Kubernetes スキャナー。 · GitHub
説明: vLLM 上で実行されている LLM を検出し、その展開と実行時の事実を抽出する Kubernetes スキャナー。 - パラレリク/piqc

記事本文:
GitHub -Paralleliq/piqc: vLLM 上で実行されている LLM を検出し、そのデプロイメントと実行時の事実を抽出する Kubernetes スキャナー。 · GitHub
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
パラレルク
/
ピクシー
公共
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
61 コミット 61 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE デプロイ デプロイ ドキュメント ドキュメントの例 例 k8s k8s piqc-test-outputs piqc-test-outputs rbac rbac src/ piqc src/ piqc テスト テスト .gitignore .gitignore CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE ModelSpec_Final_Documentation.pdf ModelSpec_Final_Documentation.pdf README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.md SECURITY.md SECURITY.md TEST_GIT.md TEST_GIT.md gcp_testing_guide.md.resolved gcp_testing_guide.md.resolved piqc-test-outputs.zip piqc-test-outputs.zip piqc_Guide.pdf piqc_Guide.pdf tongue.lock 詩.lock pyproject.toml pyproject.tomlすべてのファイルを表示 リポジトリ ファイルのナビゲーション
piqc — Kubernetes 用 GPU 廃棄物スキャナー
ほとんどの AI クラスターは GPU 支出の 20 ～ 40% を無駄にしています。 piqc は 1 つのコマンドでそれを見つけます。
読み取り専用 · エージェントなし · サイドカーなし · 永続的にインストールされるものなし · ジョブとして実行され、結果を印刷して終了します。
クイックスタート •
特徴 •
コマンド •
出力フォーマット •
インストール
piqc は、Kubernetes クラスター用のオープンソース GPU 廃棄物スキャナーです。アイドル状態の GPU 割り当て、層の配置ミス、ダーク キャパシティを検出し、1 分以内に無駄の見積もりを表示します。エージェント、サイドカー、書き込み権限はなく、永続的なインストールも必要ありません。
これは、Kubernetes クラスターが現在どれくらいの GPU を無駄にしているのか、という質問に答える最も簡単な方法です。
piqc は、標準の Kubernetes モニタリング ( kubectl top 、 kube-state-metrics 、Prometheus ノード エクスポーター) だけでは検出できない 3 種類の無駄を明らかにします。
アイドル割り当て - ほぼゼロの CO で GPU リソースを保持するポッド

計算使用率
層の配置ミス — 必要以上のメモリまたはコンピューティングを備えた GPU 層でモデルが実行されている
ダーク キャパシティ — ポッドがまったくスケジュールされていない GPU ノード
これは、GPU 推論ワークロード (GKE、EKS、AKS、オンプレミス、ベアメタル) を実行するあらゆる Kubernetes クラスターで動作します。 vLLM、Triton、TGI、および nvidia.com/gpu リソース リクエストを使用するあらゆる展開をサポートします。
クラスターに対して piqc スキャンを実行し、即時のコスト レポートを取得します。
発見された推論展開
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━━ ━┳━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┓
│ デプロイメント │ エンジン │ GPU │ レプリカ │ GPU 使用率 │ MFU │ $/1,000 トークン │ $/時間 │ アイドル $/日 │ 階層適合 │ 名前空間 │
┡━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━╇━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━━ ━╇━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━┩
│ metal-llama/Llama-3-70B-Inst │ vllm │ 8xH100-SXM4-80GB │ 2 │ 4% │ 3.1% │ $0.0842 │ $68.00 │ $1,566.72 │ ⚠ >A100-80GB │ 製品版 │
│ ミストラル-7b-命令 │ vllm │ 1xA100-SXM4-40GB │ 1 │ 11%

│ 8.4% │ $0.0073 │ $2.50 │ $53.40 │ ⚠ >T4 │ 生産 │
│ codellama-34b-staging │ vllm │ 4xH100-SXM4-80GB │ 1 │ 0% │ N/A │ N/A │ $17.00 │ $408.00 │ ⚠ >A100-40GB │ ステージング │
│ embedding-bge-large │ vllm │ 1xT4 │ 3 │ 82% │ N/A │ $0.0002 │ $1.35 │ $5.83 │ ✓ │ 共有サービス │
│ 不明-ランタイム-7f3a2 │ 不明 │ 2xA100-SXM4-80GB │ 1 │ N/A │ N/A │ N/A │ $7.00 │ ユーティリティ不明 │ ? │ ml-プラットフォーム │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┴─────┴─────┴─────┴───────┴─┴───────┴─────┴─────┘
╭───────────────── 費用概要───────────────────────╮
│ 合計 GPU 使用率 : $95.85/時間 │
│ │
│ リースおよびアイドル状態 (使用率 <60%) : $2,033.95/日 (ポッドは実行中、GPU は十分に使用されていない) │
│ 未割り当てノード: $1,152.00/日 (12 GPU、ポッドがスケジュールされていない) │
│ 階層の配置ミス: $721.20/日 (特大 GPU 階層の 3 モデル) │
│ │
│ 推定漏れ総額 : $3,907.15/日 ($1,426,110/年) │
│ │
│ 平均 MFU (アクティブな展開) : 15.7% (正常な範囲: 30 ～ 60%) │
╰───────────────────

──────────────────╯
piqc は 3 種類の廃棄物を表面化します。
アイドル状態の GPU — ポッドは実行中、GPU はほぼ空の状態
階層の配置ミス — T4 のみを必要とする H100 上の 7B モデル
未割り当てノード — ポッドがまったくスケジュールされていない GPU ノード
オプション 1: Kubernetes ジョブとして実行 (推奨)
クラスター内で実行されます。Docker 認証や kubeconfig ラングリングはありません。
# ステップ 1 — RBAC 権限を適用する (1 回限りのセットアップ)
kubectl apply -f https://raw.githubusercontent.com/Paralleliq/piqc/main/deploy/rbac.yaml
# ステップ 2 — スキャンを実行する
kubectl apply -f https://raw.githubusercontent.com/Paralleliq/piqc/main/deploy/scan-job.yaml
# ステップ 3 — 出力を表示する
kubectl ログ -f job/piqc-scan -n kube-system
#終わったらクリーンアップ
kubectl ジョブの削除 piqc-scan -n kube-system
ジョブは 10 分後に自動的に削除されます ( ttlSecondsAfterFinished: 600 )。
オプション 2: ラップトップから Docker を使用して実行する
# 認証情報が埋め込まれた静的 kubeconfig をエクスポートする
kubectl config view --raw -- flatten > /tmp/piqc-kubeconfig.yaml
# スキャンを実行する
docker run --rm \
-v /tmp/piqc-kubeconfig.yaml:/root/.kube/config \
ghcr.io/Paralleliq/piqc:最新 \
scan --format テーブル
linux/amd64 と linux/arm64 の両方をサポートします。
git clone https://github.com/Paralleliq/piqc.git
cd piqc
詩のインストール
詩の実行 piqc scan --format table
✨ 特徴
自動検出 : すべての名前空間にわたる vLLM 推論デプロイメントを自動的に検出します。
重み付き信頼スコアリング: 複数の信号 (画像、環境変数、CLI 引数、ラベル) を重み付きスコアリングで使用します。
フレームワーク検出 : パターン マッチングとヒューリスティックを使用して vLLM を高精度に識別します
📊 包括的なメトリクスの収集
GPU メトリクス: nvid を介したリアルタイムの GPU 使用率、メモリ、温度、電力

イアスミ
ランタイム メトリック: 以下を含む vLLM API メトリックを収集します。
リクエストのレイテンシー (P50、P95、P99)
トークンのスループット (プリフィルとデコード)
キューの深さとアクティブなリクエスト
GPU の利用率が低い — デプロイメントの利用率が 60% のしきい値を下回っており、1 日あたりおよび年間換算で 1 ドルの無駄が生じます。
ダーク キャパシティ — ポッドがスケジュールされていない GPU ノード (空のノードに対して料金を支払う)
層の配置ミス — モデルが特大の GPU 層で実行され、1 日あたりの推定コストデルタが発生する
断片化 — 空き GPU スロットが小さすぎて実行中のモデルに適合しないノード
保留中の GPU ポッド — スケジューリングからブロックされたワークロード (待機時間とともに表示)
コスト概要パネル - 総支出率、すべての廃棄物カテゴリ、1 日および 1 年あたりの推定漏出量の合計
MFU (モデル FLOPS 使用率) — 実測コンピューティングとデプロイごとの理論上の GPU ピークの比較
1,000 トークンあたりのコスト — GPU 支出を API 料金と同等のビジネス指標に換算
フォーマット
説明
テーブル
MFU、$/1K トークン、アイドル廃棄物を含むコスト レポート (デフォルト)
YAML
Kubernetes スタイルの推論デプロイメント ファイル
JSON
機械可読な JSON 出力
PIQCの事実
コントロールプレーン統合のための標準化されたファクトバンドル
🚀 本番環境に対応
並列処理: 構成可能なワーカーによるマルチスレッド スキャン
RBAC サポート : 事前構成された ClusterRole および ServiceAccount マニフェスト
柔軟なモード: 自動検出、リモート (kubeconfig)、またはクラスター内実行
タイムアウト制御: 設定可能な操作タイムアウト
Docker イメージ : GitHub Container Registry 上の事前構築済みマルチプラットフォーム イメージ ( linux/amd64 + linux/arm64 )
rocm-smi による AMD Instinct および Radeon GPU のサポート:
AMD Instinct MI250X/MI300X 検出
GPU 使用率、メモリ、温度のメトリクス
シームレスなマルチベンダー GPU 環境
分散 LLM 推論の検出と文書化:
分散推論トポロジーマップン

g
マルチノード GPU 調整メトリクス
ノード間のパフォーマンス集計
Kubernetes クラスターをスキャンして、推論ワークロードとサーフェス GPU の無駄を調べます。
piqcスキャン [オプション]
スキャンオプション
オプション
デフォルト
説明
--kubeconfig パス
~/.kube/config
kubeconfig ファイルへのパス
--context テキスト
現在の
使用するKubernetesコンテキスト
-n、--名前空間テキスト
すべて
スキャンする特定の名前空間
--format [yaml|json|テーブル]
ヤムル
出力フォーマット
-o, --output PATH
./出力
生成されたファイルの出力ディレクトリ
収集オプション
オプション
デフォルト
説明
--ランタイムの収集
偽
vLLM API 経由でランタイム メトリクスを収集する
--no-exec
偽
ポッド実行を無効にする (GPU メトリクスをスキップ)
--ログなし
偽
ログの読み取りを無効にする
--aggregate/--no-aggregate
集合体
ポッド レプリカ全体でメトリクスを集約する
--contribute-ベンチマーク
偽
匿名化された GPU/モデルのパフォーマンス データを Paralleliq ベンチマーク データセットに提供します
出力オプション
オプション
デフォルト
説明
--結合
偽
単一の結合出力ファイルを生成する
--output-piqc
偽
piqc-facts.json (PIQC v0.1 スキーマ) を生成します。
実行オプション
オプション
デフォルト
説明
--タイムアウトINT
30
操作のタイムアウト (秒単位)
--ワーカーINT
10
パラレルワーカーの数
--mode [自動|リモート|インクルード|ドライラン]
自動
実行モード
-v、--verbose
偽
詳細出力を有効にする
--デバッグ
偽
詳細なトレースを使用してデバッグ モードを有効にする
例
# 基本スキャン — すべての vLLM 導入と表面の無駄を検出します
piqcスキャン
# JSON 出力を使用して特定の名前空間をスキャンします
piqc scan -nproduction --format json
# GPU メトリクスを使用しないクイックスキャン (高速)
piqc スキャン --no-exec
# vLLM API からランタイム メトリクスを収集する
piqc スキャン --collect-runtime
# コントロール プレーン統合用の PIQC ファクト バンドルを生成
piqc scan --output-piqc -o ./facts
# コンソールへのテーブル出力 (人間が判読できる)
piqc scan --format テーブル
# カスタムの kubeconfig とコンテキスト
p

iqc scan --kubeconfig /path/to/config --context my-cluster
# 匿名化された GPU/モデル ベンチマークを Paralleliq データセットに提供する
piqc スキャン --contribute-benchmarks
piqc テスト接続
Kubernetes クラスターへの接続をテストし、必要な権限を確認します。
piqc テスト接続 [オプション]
オプション
デフォルト
説明
--kubeconfig パス
~/.kube/config
kubeconfig ファイルへのパス
--context テキスト
現在の
使用するKubernetesコンテキスト
piqc バージョン
piqc scan --format table を実行します。フラグは必要ありません。上記の出力例を参照してください。
デプロイメントごとに個別の Kubernetes スタイルの YAML ファイルを生成します。
APIバージョン: piqc/v1
種類 : 推論展開
メタデータ:
名前 : vllm-ラマ-7b
名前空間 : 推論
collectionTimestamp : " 2024-01-07T12:00:00Z "
コレクターバージョン: " 1.0.0 "
モデル:
名前 : メタラマ/ラマ-2-7b-hf
建築：ラマ
パラメータ：「7B」
識別信頼度 : 0.95
エンジン：
名前 : vllm
バージョン：「0.4.0」
検出信頼度 : 0.95
推論：
精度：フローア

[切り捨てられた]

## Original Extract

Kubernetes scanner that discovers LLMs running on vLLM and extracts their deployment and runtime facts. - paralleliq/piqc

GitHub - paralleliq/piqc: Kubernetes scanner that discovers LLMs running on vLLM and extracts their deployment and runtime facts. · GitHub
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
paralleliq
/
piqc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
61 Commits 61 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE deploy deploy docs docs examples examples k8s k8s piqc-test-outputs piqc-test-outputs rbac rbac src/ piqc src/ piqc tests tests .gitignore .gitignore CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE ModelSpec_Final_Documentation.pdf ModelSpec_Final_Documentation.pdf README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.md SECURITY.md SECURITY.md TEST_GIT.md TEST_GIT.md gcp_testing_guide.md.resolved gcp_testing_guide.md.resolved piqc-test-outputs.zip piqc-test-outputs.zip piqc_Guide.pdf piqc_Guide.pdf poetry.lock poetry.lock pyproject.toml pyproject.toml View all files Repository files navigation
piqc — GPU Waste Scanner for Kubernetes
Most AI clusters waste 20–40% of GPU spend. piqc finds it in one command.
Read-only · No agents · No sidecars · Nothing installed permanently · Runs as a Job, prints results, exits.
Quick Start •
Features •
Commands •
Output Formats •
Installation
piqc is an open-source GPU waste scanner for Kubernetes clusters. It detects idle GPU allocations, tier misplacement, and dark capacity — and surfaces a dollar estimate of the waste — in under a minute. No agents, no sidecars, no write permissions, no permanent installation required.
It is the fastest way to answer: how much GPU spend is my Kubernetes cluster wasting right now?
piqc surfaces three types of waste that standard Kubernetes monitoring ( kubectl top , kube-state-metrics , Prometheus node exporters) cannot detect on their own:
Idle allocation — pods holding GPU resources with near-zero compute utilization
Tier misplacement — models running on GPU tiers with far more memory or compute than they need
Dark capacity — GPU nodes with no pods scheduled at all
It works with any Kubernetes cluster running GPU inference workloads — GKE, EKS, AKS, on-prem, or bare metal. Supports vLLM, Triton, TGI, and any deployment using nvidia.com/gpu resource requests.
Run piqc scan against your cluster and get an instant cost report:
Discovered Inference Deployments
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━┳━━━━━━━━━━┳━━━━━━┳━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┓
┃ Deployment ┃ Engine ┃ GPU ┃ Replicas ┃ GPU Util ┃ MFU ┃ $/1K tokens ┃ $/hr ┃ Idle $/day ┃ Tier Fit ┃ Namespace ┃
┡━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╇━━━━━━━━━╇━━━━━━━━━━━━━━━━━━╇━━━━━━━━━━╇━━━━━━━━━━╇━━━━━━╇━━━━━━━━━━━━━╇━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━╇━━━━━━━━━━━━━━━━━┩
│ meta-llama/Llama-3-70B-Inst │ vllm │ 8xH100-SXM4-80GB │ 2 │ 4% │ 3.1% │ $0.0842 │ $68.00 │ $1,566.72 │ ⚠ >A100-80GB │ production │
│ mistral-7b-instruct │ vllm │ 1xA100-SXM4-40GB │ 1 │ 11% │ 8.4% │ $0.0073 │ $2.50 │ $53.40 │ ⚠ >T4 │ production │
│ codellama-34b-staging │ vllm │ 4xH100-SXM4-80GB │ 1 │ 0% │ N/A │ N/A │ $17.00 │ $408.00 │ ⚠ >A100-40GB │ staging │
│ embedding-bge-large │ vllm │ 1xT4 │ 3 │ 82% │ N/A │ $0.0002 │ $1.35 │ $5.83 │ ✓ │ shared-services │
│ unknown-runtime-7f3a2 │ unknown │ 2xA100-SXM4-80GB │ 1 │ N/A │ N/A │ N/A │ $7.00 │ util unknown │ ? │ ml-platform │
└─────────────────────────────┴─────────┴──────────────────┴──────────┴──────────┴──────┴─────────────┴────────┴──────────────┴──────────────┴─────────────────┘
╭──────────────────────────────────── Cost Summary ──────────────────────────────────────╮
│ Total GPU spend rate : $95.85/hr │
│ │
│ Leased & idle (util <60%) : $2,033.95/day (pods running, GPUs underused) │
│ Unallocated nodes : $1,152.00/day (12 GPU(s) with no pods scheduled) │
│ Tier misplacement : $721.20/day (3 model(s) on oversized GPU tier) │
│ │
│ Total estimated leak : $3,907.15/day ($1,426,110/yr) │
│ │
│ Avg MFU (active deployments) : 15.7% (healthy range: 30–60%) │
╰────────────────────────────────────────────────────────────────────────────────────────╯
piqc surfaces three types of waste:
Idle GPUs — pods running, GPUs sitting near-empty
Tier misplacement — a 7B model on an H100 that only needs a T4
Unallocated nodes — GPU nodes with no pods scheduled at all
Option 1: Run as a Kubernetes Job (recommended)
Runs inside your cluster — no Docker auth or kubeconfig wrangling:
# Step 1 — Apply RBAC permissions (one-time setup)
kubectl apply -f https://raw.githubusercontent.com/paralleliq/piqc/main/deploy/rbac.yaml
# Step 2 — Run the scan
kubectl apply -f https://raw.githubusercontent.com/paralleliq/piqc/main/deploy/scan-job.yaml
# Step 3 — View the output
kubectl logs -f job/piqc-scan -n kube-system
# Clean up when done
kubectl delete job piqc-scan -n kube-system
The job auto-deletes itself after 10 minutes ( ttlSecondsAfterFinished: 600 ).
Option 2: Run with Docker from your laptop
# Export a static kubeconfig with embedded credentials
kubectl config view --raw --flatten > /tmp/piqc-kubeconfig.yaml
# Run the scan
docker run --rm \
-v /tmp/piqc-kubeconfig.yaml:/root/.kube/config \
ghcr.io/paralleliq/piqc:latest \
scan --format table
Supports both linux/amd64 and linux/arm64 .
git clone https://github.com/paralleliq/piqc.git
cd piqc
poetry install
poetry run piqc scan --format table
✨ Features
Auto-Detection : Automatically discovers vLLM inference deployments across all namespaces
Weighted Confidence Scoring : Uses multiple signals (images, env vars, CLI args, labels) with weighted scoring
Framework Detection : Identifies vLLM with high accuracy using pattern matching and heuristics
📊 Comprehensive Metrics Collection
GPU Metrics : Real-time GPU utilization, memory, temperature, and power via nvidia-smi
Runtime Metrics : Collects vLLM API metrics including:
Request latency (P50, P95, P99)
Token throughput (prefill & decode)
Queue depth and active requests
GPU underutilization — Deployments below 60% utilization threshold, with dollar waste per day and annualized
Dark capacity — GPU nodes with no pods scheduled (paying for nodes sitting empty)
Tier misplacement — Models running on an oversized GPU tier, with estimated cost delta per day
Fragmentation — Nodes with free GPU slots too small to fit any running model
Pending GPU pods — Workloads blocked from scheduling, shown with wait time
Cost Summary panel — Total spend rate, all waste categories, total estimated leak per day and per year
MFU (Model FLOPS Utilization) — Observed compute vs. theoretical GPU peak per deployment
Cost per 1K tokens — GPU spend translated into a business metric comparable to API pricing
Format
Description
Table
Cost report with MFU, $/1K tokens, idle waste (default)
YAML
Kubernetes-style inference deployment files
JSON
Machine-readable JSON output
PIQC Facts
Standardized facts bundle for control plane integration
🚀 Production-Ready
Parallel Processing : Multi-threaded scanning with configurable workers
RBAC Support : Pre-configured ClusterRole and ServiceAccount manifests
Flexible Modes : Auto-detect, remote (kubeconfig), or in-cluster execution
Timeout Controls : Configurable operation timeouts
Docker Image : Pre-built multi-platform image ( linux/amd64 + linux/arm64 ) on GitHub Container Registry
Support for AMD Instinct and Radeon GPUs via rocm-smi :
AMD Instinct MI250X/MI300X detection
GPU utilization, memory & temperature metrics
Seamless multi-vendor GPU environments
Discovery and documentation for distributed LLM inference:
Distributed inference topology mapping
Multi-node GPU coordination metrics
Cross-node performance aggregation
Scan your Kubernetes cluster for inference workloads and surface GPU waste.
piqc scan [OPTIONS]
Scan Options
Option
Default
Description
--kubeconfig PATH
~/.kube/config
Path to kubeconfig file
--context TEXT
current
Kubernetes context to use
-n, --namespace TEXT
all
Specific namespace to scan
--format [yaml|json|table]
yaml
Output format
-o, --output PATH
./output
Output directory for generated files
Collection Options
Option
Default
Description
--collect-runtime
false
Collect runtime metrics via vLLM API
--no-exec
false
Disable pod exec (skip GPU metrics)
--no-logs
false
Disable log reading
--aggregate/--no-aggregate
aggregate
Aggregate metrics across pod replicas
--contribute-benchmarks
false
Contribute anonymized GPU/model performance data to the Paralleliq benchmark dataset
Output Options
Option
Default
Description
--combined
false
Generate single combined output file
--output-piqc
false
Generate piqc-facts.json (PIQC v0.1 schema)
Execution Options
Option
Default
Description
--timeout INT
30
Operation timeout in seconds
--workers INT
10
Number of parallel workers
--mode [auto|remote|incluster|dry-run]
auto
Execution mode
-v, --verbose
false
Enable verbose output
--debug
false
Enable debug mode with detailed trace
Examples
# Basic scan — discover all vLLM deployments and surface waste
piqc scan
# Scan specific namespace with JSON output
piqc scan -n production --format json
# Quick scan without GPU metrics (faster)
piqc scan --no-exec
# Collect runtime metrics from vLLM API
piqc scan --collect-runtime
# Generate PIQC facts bundle for control plane integration
piqc scan --output-piqc -o ./facts
# Table output to console (human-readable)
piqc scan --format table
# Custom kubeconfig and context
piqc scan --kubeconfig /path/to/config --context my-cluster
# Contribute anonymized GPU/model benchmarks to Paralleliq dataset
piqc scan --contribute-benchmarks
piqc test-connection
Test connection to Kubernetes cluster and verify required permissions.
piqc test-connection [OPTIONS]
Option
Default
Description
--kubeconfig PATH
~/.kube/config
Path to kubeconfig file
--context TEXT
current
Kubernetes context to use
piqc version
Run piqc scan --format table — no flags required. See the output example above.
Generates individual Kubernetes-style YAML files for each deployment:
apiVersion : piqc/v1
kind : InferenceDeployment
metadata :
name : vllm-llama-7b
namespace : inference
collectionTimestamp : " 2024-01-07T12:00:00Z "
collectorVersion : " 1.0.0 "
model :
name : meta-llama/Llama-2-7b-hf
architecture : llama
parameters : " 7B "
identificationConfidence : 0.95
engine :
name : vllm
version : " 0.4.0 "
detectionConfidence : 0.95
inference :
precision : floa

[truncated]

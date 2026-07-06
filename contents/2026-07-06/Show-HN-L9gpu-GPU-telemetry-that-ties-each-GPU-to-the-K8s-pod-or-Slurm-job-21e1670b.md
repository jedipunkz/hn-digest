---
source: "https://github.com/last9/gpu-telemetry"
hn_url: "https://news.ycombinator.com/item?id=48811382"
title: "Show HN: L9gpu – GPU telemetry that ties each GPU to the K8s pod or Slurm job"
article_title: "GitHub - last9/gpu-telemetry: GPU Observability with workload attribution. One OTLP agent per node ties hardware metrics (NVIDIA, AMD, Intel Gaudi) to the K8s pod or Slurm job burning the GPU. · GitHub"
author: "nishantmodak"
captured_at: "2026-07-06T22:58:47Z"
capture_tool: "hn-digest"
hn_id: 48811382
score: 2
comments: 0
posted_at: "2026-07-06T22:33:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: L9gpu – GPU telemetry that ties each GPU to the K8s pod or Slurm job

- HN: [48811382](https://news.ycombinator.com/item?id=48811382)
- Source: [github.com](https://github.com/last9/gpu-telemetry)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T22:33:57Z

## Translation

タイトル: Show HN: L9gpu – 各 GPU を K8s ポッドまたは Slurm ジョブに結び付ける GPU テレメトリ
記事のタイトル: GitHub - last9/gpu-telemetry: ワークロード アトリビューションによる GPU オブザーバビリティ。ノードごとに 1 つの OTLP エージェントが、ハードウェア メトリクス (NVIDIA、AMD、Intel Gaudi) を K8s ポッドまたは GPU を焼き付ける Slurm ジョブに結び付けます。 · GitHub
説明: ワークロード属性による GPU オブザーバビリティ。ノードごとに 1 つの OTLP エージェントが、ハードウェア メトリクス (NVIDIA、AMD、Intel Gaudi) を K8s ポッドまたは GPU を焼き付ける Slurm ジョブに結び付けます。 - last9/GPU テレメトリ

記事本文:
GitHub - last9/gpu-telemetry: ワークロード アトリビューションによる GPU オブザーバビリティ。ノードごとに 1 つの OTLP エージェントが、ハードウェア メトリクス (NVIDIA、AMD、Intel Gaudi) を K8s ポッドまたは GPU を焼き付ける Slurm ジョブに結び付けます。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました

g.このページをリロードしてください。
最後9
/
GPU テレメトリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
42 コミット 42 コミット .github .github アラート アラート ダッシュボード/グラファナ ダッシュボード/グラファナ デプロイ デプロイ docker docker docs docs k8shelper k8shelper k8sprocessor k8sprocessor l9gpu l9gpu スクリプト スクリプト シェルパー シェルパー slurmprocessor slurmprocessor スタブ スタブ systemd systemd .clang-format .clang-format .dockerignore .dockerignore .env.example .env.example .flake8 .flake8 .git_archival.txt .git_archival.txt .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml .isort.cfg .isort.cfg .pre-commit-config.yaml .pre-commit-config.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md dev-requirements.txt dev-requirements.txt mypy.ini mypy.ini noxfile.py noxfile.py pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
DCGM エクスポータは、GPU がホットであることを示します。誰の仕事が揚げているかはわかりません。
ほとんどの GPU 可観測性はハードウェア (使用率、温度、ECC) で止まります。
そして、重要な唯一の質問に対する答えなしで gpu.uuid を渡します。
このアイドル状態の H100 に誰がお金を払っているのでしょうか?
l9gpu はループを閉じます。ノードごとに 1 つのエージェントがベンダー中立の OTLP を発行します。
組み込まれたワークロード属性 — Kubernetes ポッド、名前空間、デプロイメント。
Slurm ジョブ、ユーザー、パーティション。任意の OTLP バックエンドにポイントすると、次の結果が得られます。
パイプラインを構築せずに、チームごと、ジョブごと、モデルごとにアカウンティングします。
現在、NVIDIA、AMD、Intel Gaudi で動作します。引き続き取り組んでいきます
次に何が起こっても

特注のフォーマットではなく、OpenTelemetry を発行します。
エージェント自体にはベンダー バックエンドはありません。それは故意です。
# Classic Helm リポジトリ
Helm リポジトリ l9gpu を追加 https://last9.github.io/gpu-telemetry
helm install l9gpu l9gpu/l9gpu -n モニタリング --create-namespace \
--set Monitoring.sink=otel \
--set Monitoring.cluster=my-cluster \
--set otlpSecretName=l9gpu-otlp
# または OCI
helm install l9gpu oci://ghcr.io/last9/gpu-telemetry/l9gpu --version 0.2.1 -n モニタリング
まず OTLP シークレットを作成します。
kubectl create secret generic l9gpu-otlp -n 監視 \
--from-literal=OTEL_EXPORTER_OTLP_ENDPOINT= < otlp エンドポイント > \
--from-literal=OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Bearer <your-token> "
AMD / Gaudi ノード: --setcollectors.nvidia=false --setcollectors.amd=true
(またはcollectors.gaudi=true)。
完全な Helm ガイド: docs/HELM.md 。トポロジーの例
(EKS + DCGM、マルチ GPU、サイドカー コレクター):deploy/helm/l9gpu/examples/ 。
クイックスタート — ベアメタル / systemd
pip インストール l9gpu
import OTEL_EXPORTER_OTLP_ENDPOINT= < あなたの otlp エンドポイント >
import OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Bearer <your-token> "
l9gpu nvml_monitor --sink otel --cluster my-cluster # NVIDIA
l9gpu amd_monitor --sink otel --cluster my-cluster # AMD
l9gpu gaudi_monitor --sink otel --cluster my-cluster # Intel Gaudi
OTLP を使用しない健全性チェック: l9gpu nvml_monitor --sink stdout --once 。
systemd ユニット ファイル: systemd/ .
Prometheus エクスポーターではありません。 OTLPを発します。あなたのコレクターがハンドルします
必要に応じてプロメテウスをスクレイピングします。
バックエンドではありません。 l9gpu は、標準 OTLP を OTLP を通信するものにエクスポートします。
エージェントには Last9 ロックインはありません。
DCGM の代替品ではありません。 DCGM プロファイリング (SM 占有率、テンソル パイプ、
NVLink) は補完的であり、1 つのコレクタ パイプラインを通じて両方をバンドルします。
エヌビディアだけではありません。 AMD MI300X / MI325X および Intel Gaudi 2/3 ar

ファーストクラス。
各ノードのコレクタは、NVML / DCGM / amdsmi / hl-smi を正規化して、
gpu.* OTel 名前空間を使用し、OTLP をコレクタに送信します。コレクターは充実させます
k8sプロセッサまたはslurmプロセッサを搭載
そしてファンアウトします。
すべてのサイクル (デフォルトは 60 秒) でメトリクスが生成されます (GPU ごとに 1 つの OTLP ゲージ)
メトリック）およびログ（完全なスナップショットを含むサイクルごとに GPU ごとに 1 つの OTLP ログ —
ログ形式のイベントや履歴の再生を好むバックエンドに役立ちます)。
完全なウォークスルー: docs/ARCHITECTURE.md 。
Kubernetes — k8sprocessor が各 GPU データを強化します
k8s.pod.name 、 k8s.namespace.name 、 k8s.deployment.name のポイント、
k8s.job.name 、cloud.availability_zone 、cloud.region 。セットアップ、
RBAC、ラベル許可リスト: docs/K8S_WORKLOAD_ATTRIBUTION.md 。
Slurm — slurmprocessor が slurm.job.id を添付します。
slurm.user 、 slurm.account 、パーティション、QoS:
プロセッサー:
スラム：
キャッシュ期間 : 60
キャッシュファイルパス: /tmp/slurmprocessor_cache.json
query_slurmctld : false
完全な構成: slurmprocessor/README.md 。
事前に構築された Grafana ダッシュボード (dashboards/grafana/) —
マルチクラスター フリート、ポッドごとのワークロード、健全性/信頼性 (ECC、スロットリング、
XID)、DCGM プロファイリング、推論エンジン (vLLM、SGLang、TGI、Triton、NIM)、
フリート効率/アイドル状態の検出。
Alerts/prometheus/ のアラート ルール — 17 個の事前構築済み
3 つの PrometheusRule CRD (フリート、GPU インフラストラクチャ、LLM) にわたるアラート ルール
推論) — そして、alerts/grafana/ 。 Helm 経由で有効にする:
Helm upgrade --setalerts.enabled=true … 。
ocb をスキップして、k8sprocessor + を使用して既製のコレクタを実行します
slurmプロセッサが組み込まれています:
docker run --rm -v $PWD /config.yaml:/etc/l9gpu/config.yaml:ro \
ghcr.io/last9/l9gpu-collector:latest --config=/etc/l9gpu/config.yaml
詳細とバイナリ/tarball インストール: docs/COLLECTOR.md 。
ディレクトリ
言語
役割
l9gpu/
パイソン
ノードレベルのコレクター (DaemonSet / systemd)。 OTLP を満たした排出量

リック+ログ。
k8sプロセッサ/
行く
OTel コレクタープロセッサ。 K8s ポッド / ワークロード / クラウド メタデータで強化します。
スラームプロセッサ/
行く
OTel コレクタープロセッサ。 Slurm ジョブのメタデータを強化します。
k8シェルパー/
行く
共有 K8s API ヘルパー ライブラリ。
ヘルパー/
行く
共有 Slurm ヘルパー ライブラリ。
ハードウェアサポート
NVIDIA A100、H100 / H200、B200 / GB200、T4、A10、L4 (NVML + DCGM) ・ AMD
MI300X、MI325X (amdsmi) · Intel Gaudi 2、Gaudi 3 (hl-smi)。
単位と属性を含む完全なメトリック カタログ: docs/METRICS.md 。
ワンコマンド EKS スタック — vLLM + SGLang + TGI + Triton と l9gpu NVML、
DCGM、コスト、フリートの健全性、およびエンジンごとのモニター:
./deploy/demo/launch.sh
ドキュメント
アーキテクチャ — システム設計、トポロジー、データフロー
メトリクスのリファレンス — すべてのメトリクス、単位、属性
統合ガイド — PromQL、OTel Collector レシピ、クラウド ノート
K8s ワークロードの属性 — RBAC、エンリッチメント、ラベル許可リスト
スケーリング - 大規模なフリートのカーディナリティ管理
GPU および LLM の可観測性 — vLLM / NIM / Triton の詳細
ヘルムのインストール · 構築済みコレクター
AWS テスト クックブック — エンドツーエンドの EC2 および EKS ウォークスルー
l9gpu CLI リファレンス · slurmprocessor · シェルパー
PR の方は大歓迎です。開発セットアップ、テストについては CONTRIBUTING.md を参照してください。
そしてPRの流れ。貢献することにより、あなたの作品が同じライセンスに基づいてライセンスされることに同意したことになります
プロジェクトの残りの部分と同様に条件を変更します。セキュリティレポート: SECURITY.md 。
l9gpu (Python パッケージ)、 shelper 、および slurmprocessor から派生したもの
Meta の facebookresearch/gcm より
プロジェクト (MIT および Apache-2.0)。 Kubernetes ワークロードでそれらを拡張しました
帰属、AMD / Intel Gaudi コレクター、vLLM / SGLang / TGI / Triton /
NIM モニター、コストとフリートの健全性シグナル、および OTLP ネイティブのエクスポート。
k8shelper/ と k8sprocessor/ は Last9 のオリジナル作品です。参照
完全な内訳についての通知。
l9gpu 、 k8shelper 、 k8sprocessor の MIT。アパッチ-2.0

スラームプロセッサー用、
ヘルパー。各サブディレクトリには独自のライセンスが含まれており、異なる点は
リポジトリのルート。
ワークロード属性による GPU オブザーバビリティ。ノードごとに 1 つの OTLP エージェントが、ハードウェア メトリクス (NVIDIA、AMD、Intel Gaudi) を K8s ポッドまたは GPU を焼き付ける Slurm ジョブに結び付けます。
last9.io/gpu-observability/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
4
フォーク
レポートリポジトリ
リリース
4
l9gpu コレクター 0.2.0
最新の
2026 年 6 月 2 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

GPU Observability with workload attribution. One OTLP agent per node ties hardware metrics (NVIDIA, AMD, Intel Gaudi) to the K8s pod or Slurm job burning the GPU. - last9/gpu-telemetry

GitHub - last9/gpu-telemetry: GPU Observability with workload attribution. One OTLP agent per node ties hardware metrics (NVIDIA, AMD, Intel Gaudi) to the K8s pod or Slurm job burning the GPU. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
last9
/
gpu-telemetry
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
42 Commits 42 Commits .github .github alerts alerts dashboards/ grafana dashboards/ grafana deploy deploy docker docker docs docs k8shelper k8shelper k8sprocessor k8sprocessor l9gpu l9gpu scripts scripts shelper shelper slurmprocessor slurmprocessor stubs stubs systemd systemd .clang-format .clang-format .dockerignore .dockerignore .env.example .env.example .flake8 .flake8 .git_archival.txt .git_archival.txt .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml .isort.cfg .isort.cfg .pre-commit-config.yaml .pre-commit-config.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md dev-requirements.txt dev-requirements.txt mypy.ini mypy.ini noxfile.py noxfile.py pyproject.toml pyproject.toml View all files Repository files navigation
DCGM exporter tells you a GPU is hot. It won't tell you whose job is frying it.
Most GPU observability stops at the hardware — utilization, temperature, ECC —
and hands you a gpu.uuid with no answer to the only question that matters:
who's paying for this idle H100?
l9gpu closes the loop. One agent per node emits vendor-neutral OTLP with
workload attribution baked in — Kubernetes pod, namespace, deployment;
Slurm job, user, partition. You point it at any OTLP backend and get
per-team, per-job, per-model accounting without building a pipeline.
It works on NVIDIA, AMD, and Intel Gaudi today. It will keep working on
whatever comes next because it emits OpenTelemetry, not a bespoke format.
There's no vendor backend in the agent itself. That's deliberate.
# Classic Helm repo
helm repo add l9gpu https://last9.github.io/gpu-telemetry
helm install l9gpu l9gpu/l9gpu -n monitoring --create-namespace \
--set monitoring.sink=otel \
--set monitoring.cluster=my-cluster \
--set otlpSecretName=l9gpu-otlp
# or OCI
helm install l9gpu oci://ghcr.io/last9/gpu-telemetry/l9gpu --version 0.2.1 -n monitoring
Create the OTLP secret first:
kubectl create secret generic l9gpu-otlp -n monitoring \
--from-literal=OTEL_EXPORTER_OTLP_ENDPOINT= < your-otlp-endpoint > \
--from-literal=OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Bearer <your-token> "
AMD / Gaudi nodes: --set collectors.nvidia=false --set collectors.amd=true
(or collectors.gaudi=true ).
Full Helm guide: docs/HELM.md . Topology examples
(EKS + DCGM, multi-GPU, sidecar collector): deploy/helm/l9gpu/examples/ .
Quick Start — Bare Metal / systemd
pip install l9gpu
export OTEL_EXPORTER_OTLP_ENDPOINT= < your-otlp-endpoint >
export OTEL_EXPORTER_OTLP_HEADERS= " Authorization=Bearer <your-token> "
l9gpu nvml_monitor --sink otel --cluster my-cluster # NVIDIA
l9gpu amd_monitor --sink otel --cluster my-cluster # AMD
l9gpu gaudi_monitor --sink otel --cluster my-cluster # Intel Gaudi
Sanity-check without OTLP: l9gpu nvml_monitor --sink stdout --once .
systemd unit files: systemd/ .
Not a Prometheus exporter. It emits OTLP. Your Collector handles
Prometheus scraping if you want it.
Not a backend. l9gpu exports standard OTLP to whatever speaks OTLP.
There's no Last9 lock-in in the agent.
Not a DCGM replacement. DCGM profiling (SM occupancy, tensor pipe,
NVLink) is complementary — bundle both through one Collector pipeline.
Not only NVIDIA. AMD MI300X / MI325X and Intel Gaudi 2/3 are first-class.
Collectors on each node normalize NVML / DCGM / amdsmi / hl-smi into the
gpu.* OTel namespace and ship OTLP to a Collector. The Collector enriches
with k8sprocessor or slurmprocessor
and fans out.
Every cycle (default 60s) emits metrics (one OTLP gauge per GPU per
metric) and logs (one OTLP log per GPU per cycle with the full snapshot —
useful for backends that prefer log-shaped events or for replaying history).
Full walk-through: docs/ARCHITECTURE.md .
Kubernetes — k8sprocessor enriches each GPU data
point with k8s.pod.name , k8s.namespace.name , k8s.deployment.name ,
k8s.job.name , cloud.availability_zone , cloud.region . Setup,
RBAC, label allow-lists: docs/K8S_WORKLOAD_ATTRIBUTION.md .
Slurm — slurmprocessor attaches slurm.job.id ,
slurm.user , slurm.account , partition, QoS:
processors :
slurm :
cache_duration : 60
cache_filepath : /tmp/slurmprocessor_cache.json
query_slurmctld : false
Full config: slurmprocessor/README.md .
Pre-built Grafana dashboards in dashboards/grafana/ —
multi-cluster fleet, per-pod workload, health/reliability (ECC, throttling,
XID), DCGM profiling, inference engines (vLLM, SGLang, TGI, Triton, NIM),
fleet efficiency / idle detection.
Alert rules in alerts/prometheus/ — 17 pre-built
alert rules across 3 PrometheusRule CRDs (fleet, GPU infrastructure, LLM
inference) — and alerts/grafana/ . Enable via Helm:
helm upgrade --set alerts.enabled=true … .
Skip ocb and run a ready-made Collector with k8sprocessor +
slurmprocessor baked in:
docker run --rm -v $PWD /config.yaml:/etc/l9gpu/config.yaml:ro \
ghcr.io/last9/l9gpu-collector:latest --config=/etc/l9gpu/config.yaml
Details and binary/tarball install: docs/COLLECTOR.md .
Directory
Language
Role
l9gpu/
Python
Node-level collector (DaemonSet / systemd). Emits OTLP metrics + logs.
k8sprocessor/
Go
OTel Collector processor. Enriches with K8s pod / workload / cloud metadata.
slurmprocessor/
Go
OTel Collector processor. Enriches with Slurm job metadata.
k8shelper/
Go
Shared K8s API helper library.
shelper/
Go
Shared Slurm helper library.
Hardware support
NVIDIA A100, H100 / H200, B200 / GB200, T4, A10, L4 (NVML + DCGM) · AMD
MI300X, MI325X (amdsmi) · Intel Gaudi 2, Gaudi 3 (hl-smi).
Full metric catalog with units and attributes: docs/METRICS.md .
One-command EKS stack — vLLM + SGLang + TGI + Triton alongside l9gpu NVML,
DCGM, cost, fleet-health, and per-engine monitors:
./deploy/demo/launch.sh
Documentation
Architecture — system design, topology, data flow
Metrics reference — every metric, unit, attribute
Integration guide — PromQL, OTel Collector recipes, cloud notes
K8s workload attribution — RBAC, enrichment, label allow-lists
Scaling — cardinality management for large fleets
GPU & LLM observability — vLLM / NIM / Triton specifics
Helm install · Pre-built collector
AWS testing cookbook — end-to-end EC2 and EKS walk-through
l9gpu CLI reference · slurmprocessor · shelper
PRs welcome. See CONTRIBUTING.md for dev setup, tests,
and PR flow. By contributing you agree your work is licensed under the same
terms as the rest of the project. Security reports: SECURITY.md .
l9gpu (the Python package), shelper , and slurmprocessor are derived
from Meta's facebookresearch/gcm
project (MIT and Apache-2.0). We extended them with Kubernetes workload
attribution, AMD / Intel Gaudi collectors, vLLM / SGLang / TGI / Triton /
NIM monitors, cost and fleet-health signals, and OTLP-native export.
k8shelper/ and k8sprocessor/ are original Last9 work. See
NOTICE for the full breakdown.
MIT for l9gpu , k8shelper , k8sprocessor . Apache-2.0 for slurmprocessor ,
shelper . Each subdirectory carries its own LICENSE where it differs from
the repo root.
GPU Observability with workload attribution. One OTLP agent per node ties hardware metrics (NVIDIA, AMD, Intel Gaudi) to the K8s pod or Slurm job burning the GPU.
last9.io/gpu-observability/
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
4
forks
Report repository
Releases
4
l9gpu-collector 0.2.0
Latest
Jun 2, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

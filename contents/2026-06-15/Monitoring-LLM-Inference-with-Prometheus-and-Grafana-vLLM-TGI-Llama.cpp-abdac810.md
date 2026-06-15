---
source: "https://www.glukhov.org/observability/monitoring-llm-inference-prometheus-grafana/"
hn_url: "https://news.ycombinator.com/item?id=48535910"
title: "Monitoring LLM Inference with Prometheus and Grafana (vLLM, TGI, Llama.cpp)"
article_title: "Monitor LLM Inference in Production (2026): Prometheus & Grafana for vLLM, TGI, llama.cpp - Rost Glukhov | Personal site and technical blog"
author: "nryoo"
captured_at: "2026-06-15T04:37:34Z"
capture_tool: "hn-digest"
hn_id: 48535910
score: 1
comments: 0
posted_at: "2026-06-15T02:34:15Z"
tags:
  - hacker-news
  - translated
---

# Monitoring LLM Inference with Prometheus and Grafana (vLLM, TGI, Llama.cpp)

- HN: [48535910](https://news.ycombinator.com/item?id=48535910)
- Source: [www.glukhov.org](https://www.glukhov.org/observability/monitoring-llm-inference-prometheus-grafana/)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T02:34:15Z

## Translation

タイトル: Prometheus と Grafana を使用した LLM 推論のモニタリング (vLLM、TGI、Llama.cpp)
記事のタイトル: 本番環境での LLM 推論の監視 (2026): vLLM、TGI、llama.cpp の Prometheus と Grafana - Rost Glukhov |個人サイトと技術ブログ
説明: Prometheus と Grafana を使用して本番環境で LLM 推論を監視する方法を学びます。 vLLM、TGI、および llama.cpp にわたる p95 レイテンシ、トークン/秒、キュー期間、KV キャッシュの使用状況を追跡します。 PromQL のサンプル、ダッシュボード、アラート、Docker と Kubernetes のセットアップが含まれます。

記事本文:
余白についての注意事項
ロスト・グルホフ。個人サイトと技術ブログ
メニュー
AIシステム
運用環境での LLM 推論の監視 (2026): vLLM、TGI、llama.cpp の Prometheus と Grafana
Prometheus と Grafana を使用して LLM を監視する
LLM 推論は、レイテンシーが急増し、キューがバックアップされ、明らかな説明もなく GPU のメモリが 95% にとどまるまでは、「単なる別の API」のように見えます。
単一ノードのセットアップを超えた瞬間、またはスループットの最適化を開始した瞬間から、監視はミッション クリティカルになります。その時点で、従来の API メトリクスでは十分ではありません。最新の LLM システムの本当のボトルネックであるトークン、バッチ動作、キュー時間、KV キャッシュの圧力を可視化する必要があります。
この記事は、私の広範な可観測性と監視に関するガイドの一部であり、監視と可観測性の基本、Prometheus アーキテクチャ、運用のベスト プラクティスについて説明しています。ここでは、特に LLM 推論ワークロードの監視に焦点を当てます。
(インフラストラクチャについて決定している場合は、2026 年の LLM ホスティングに関する私のガイドを参照してください。バッチ処理の仕組み、VRAM 制限、スループットとレイテンシのトレードオフについて詳しく知りたい場合は、LLM パフォーマンス エンジニアリング ガイドを参照してください。)
一般的な REST サービスとは異なり、LLM サービスは、トークン、連続バッチ処理、KV キャッシュ使用率、GPU/CPU 飽和、キュー ダイナミクスによって形成されます。ペイロード サイズが同じ 2 つのリクエストは、 max_new_tokens 、同時実行性、およびキャッシュの再利用に応じて、レイテンシーが大幅に異なる可能性があります。
このガイドは、Prometheus と Grafana を使用して LLM 推論モニタリングを構築するための、実稼働に焦点を当てた実践的なチュートリアルです。
測定対象 (p95/p99 レイテンシ、トークン/秒、キュー期間、キャッシュ使用率、エラー率)
共通サーバーから /metrics を収集する方法 ( vLLM 、 Hugging Face TGI 、 llama.cpp )
パーセンタイル、サチュラの PromQL の例

ションとスループット
Docker Compose と Kubernetes を使用したデプロイメント パターン
実際の負荷下でのみ発生する問題のトラブルシューティング
これらの例は意図的にベンダー中立になっています。後で OpenTelemetry トレース、自動スケーリング、またはサービス メッシュを追加する場合でも、同じメトリック モデルが適用されます。
LLM 推論を別の方法で監視する必要がある理由
従来の API モニタリング (RPS、p95 レイテンシ、エラー率) は必要ですが、十分ではありません。 LLM サービングにより追加の軸が追加されます。
E2E レイテンシ: リクエストを受信してから最終トークンが返されるまでの時間。
トークン間レイテンシ: デコード中のトークンあたりの時間 (ストリーミング UX にとって重要)。
一部のサーバーは両方を公開します。たとえば、TGI はリクエスト期間とトークンごとの平均時間をヒストグラムとして公開します。
2) スループットはリクエストではなくトークンで表されます
5 つのトークンを返す「高速」サービスは、500 のトークンを返すサービスと比較できません。多くの場合、「RPS」は「トークン/秒」になります。
連続バッチ処理を実行する場合、キューの深さが売りになります。キューの継続時間とキューのサイズを監視することで、ユーザーの期待に応えているかどうかがわかります。
4) キャッシュプレッシャーは停止の前兆である
KV キャッシュの枯渇 (または断片化) は、多くの場合、突然のレイテンシのスパイクやタイムアウトとして現れます。 vLLM は、KV キャッシュの使用状況をゲージとして公開します。
LLM 推論モニタリングのメトリクス チェックリスト
これを北極星として使用してください。初日にすべてが必要になるわけではありませんが、最終的にはほとんどが必要になります。
トラフィック: リクエスト/秒、トークン/秒
エラー: エラー率、タイムアウト、OOM、429 (レート制限)
レイテンシ: p50/p95/p99 リクエスト期間。プリフィルとデコードのレイテンシー。トークン間のレイテンシ
飽和: GPU 使用率、メモリ使用量、KV キャッシュ使用量、キュー サイズ
GPU のメモリ使用量、温度、Prometheus 以外の使用状況 (デバッグまたは単一ノードのセットアップ用) を低レベルで可視化する必要がある場合は、私の GPU ガイドを参照してください。

Linux/Ubuntu でのアプリケーションの監視。
トレース、構造化ログ、合成テスト、GPU プロファイリング、SLO 設計など、メトリクスを超えた LLM 可観測性のより広い視野については、LLM システムの可観測性に関する詳細なガイドを参照してください。
ラベルのカーディナリティを低く保ちます。良いラベル:
モデル 、 エンドポイント 、 メソッド (プリフィル/デコード)、 ステータス (成功/エラー)、 インスタンス
rawプロンプト、生のuser_id、リクエストID - これらはシリーズ数を爆発的にカウントします。
メトリクスの公開: 組み込みの /metrics エンドポイント (vLLM、TGI、llama.cpp)
最も簡単な方法は、サーバーがすでに公開しているメトリクスを使用することです。
vLLM: Prometheus 互換の /metrics
vLLM は、Prometheus 互換の /metrics エンドポイントを (Prometheus メトリクス ロガー経由で) 公開し、実行中のリクエストや KV キャッシュの使用状況などのゲージを含む、vllm: プレフィックスが付いたサーバー/リクエスト メトリクスを公開します。
コンテナーのセットアップ、OpenAI 互換のサービス提供、およびスループット指向のランタイム調整については、「vLLM クイックスタート: ハイパフォーマンス LLM サービス」を参照してください。
通常表示されるメトリクスの例:
ハグフェイス TGI: キューとリクエストのヒストグラムを含む /metrics
TGI は、キュー サイズ、リクエスト期間、キュー期間、トークンあたりの平均時間を含む、多くの実稼働グレードのメトリクスを /metrics で公開します。
tgi_request_duration (ヒストグラム、e2e レイテンシ)
tgi_request_queue_duration (ヒストグラム)
tgi_request_mean_time_per_token_duration (ヒストグラム)
運用セットアップ (Docker、GPU、起動フラグ、および空のスクレイピングまたは誤解を招くスクレイピングとして表示される障害) については、「TGI - テキスト生成推論 - インストール、構成、トラブルシューティング」で説明されています。
llama.cpp サーバー: メトリクスエンドポイントを有効にする
llama.cpp サーバーは、フラグ ( --metrics など) を使用して有効にする必要がある Prometheus 互換のメトリクス エンドポイントをサポートします。
インストール パス、主要なランタイム フラグ、および OpenAI 互換サーバーの使用法については、llama.cpp のクイックスタートを参照してください。

CLI とサーバー。
プロキシの背後で llama.cpp を実行している場合は、可能な限りサーバーを直接スクレイピングしてください (プロキシレベルの遅延によって実際の推論動作が隠蔽されるのを避けるため)。
Prometheus 構成: 推論サーバーのスクレイピング
vLLM (http://vllm:8000/metrics)
TGI (http://tgi:8080/metrics)
llama.cpp (http://llama:8080/metrics)
高速フィードバック用に調整されたスクレイピング間隔
グローバル :
スクレイピング間隔 : 5秒
評価間隔 : 15秒
スクレイピング_構成:
- ジョブ名: "vllm"
metrics_path : /metrics
static_configs :
- ターゲット: [ "vllm:8000" ]
- ジョブ名 : "tgi"
metrics_path : /metrics
static_configs :
- ターゲット: [ "tgi:8080" ]
- ジョブ名: "llama_cpp"
metrics_path : /metrics
static_configs :
- ターゲット: [ "ラマ:8080" ]
Prometheus を初めて使用する場合、またはスクレイピング構成、エクスポーター、再ラベル付け、アラート ルールについて詳しく説明したい場合は、完全な Prometheus 監視セットアップ ガイドを参照してください。
プロのヒント: 「サービス ラベル」を追加する
複数のモデル/レプリカを実行する場合は、ダッシュボードに安定したサービス ラベルを含めるために再ラベルを追加します。
relabel_configs :
- target_label : サービス
置換:「llm推論」
コピー/ペーストできる PromQL の例
sum ( rate ( tgi_request_count[ 5m ] ))
vLLM の場合は、そのリクエスト カウンター (名前はバージョンによって異なります) を使用しますが、パターンは同じです: sum(rate(<counter>[5m])) 。
*_success カウンターがある場合は、失敗率を計算します。
1 - (
sum ( rate ( tgi_request_success[ 5m ] ))
/
sum ( rate ( tgi_request_count[ 5m ] ))
)
ヒストグラムメトリクスの p95 レイテンシー (Prometheus)
Prometheus のヒストグラムはバケット化されたカウントです。バケットの rate() ではなく histogram_quantile() を使用します。 Prometheus は、このモデルとヒストグラムと要約のトレードオフを文書化しています。
ヒストグラム_分位点 (
0.95、
( le ) ( rate ( tgi_request_duration_bucket[ 5m ] )) による合計
)
p99 キュータイム
ヒストグラム_分位点 (
0.99、
( le による合計)

) (レート ( tgi_request_queue_duration_bucket[ 5m ] ))
)
トークンあたりの平均時間 (トークン間の待ち時間)
ヒストグラム_分位点 (
0.95、
( le ) ( rate ( tgi_request_mean_time_per_token_duration_bucket[ 5m ] )) による合計
)
トークン間のレイテンシは、デコードのボトルネックとメモリ帯域幅によって制約されることがよくあります。トピックについては、「
LLM パフォーマンス最適化ガイド。
最大 ( tgi_queue_size )
vLLM KV キャッシュ使用率 (瞬間)
max ( vllm : kv_cache_usage_perc )
Grafana ダッシュボード: オンコールに実際に役立つパネル
Grafana は、複数の方法 (パーセンタイル、ヒートマップ、バケット分布) でヒストグラムを視覚化できます。 Grafana Labs には、Prometheus ヒストグラム視覚化に関する詳細なガイドがあります。
最小限の高信号ダッシュボード レイアウト:
p95 リクエスト レイテンシ (時系列)
p95 トークン間レイテンシー (時系列)
エラー率 (時系列 + 統計)
行 2 — 容量と飽和度
実行中のリクエストと待機中のリクエスト (スタック)
生成されたトークン/リクエスト (p50/p95)
ストリーミングを使用している場合は、利用可能な場合は「最初のトークン レイテンシ」(TTFT) のパネルを追加します。
p95 レイテンシ パネル: 上記の histogram_quantile(0.95, …) クエリ
ヒートマップ パネル: バケット レート ( *_bucket ) をヒートマップとしてグラフ化します (Grafana はこのアプローチをサポートしています)
導入オプション 1: Docker Compose (高速ローカル + 単一ノード)
ローカル、セルフホスト、またはクラウドベースの推論アーキテクチャのいずれかを選択する場合は、私の詳細な内訳を参照してください。
LLM ホスティング比較ガイド 。
モニタリング/
docker-compose.yml
プロメテウス/
プロメテウス.yml
グラファナ/
プロビジョニング/
データソース/datasource.yml
ダッシュボード/ダッシュボード.yml
ダッシュボード/
llm-inference.json
docker-compose.yml
サービス:
プロメテウス：
画像 : プロム/プロメテウス:最新
ボリューム:
- ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml:ro
ポート:
- 「9090:9090」
グラファナ :
画像 : グラファナ/グラファナ:最新
環境：
- GF_SECURITY_ADMIN

_USER=管理者
- GF_SECURITY_ADMIN_PASSWORD=管理者
ボリューム:
- ./grafana/provisioning:/etc/grafana/provisioning
- ./grafana/dashboards:/var/lib/grafana/dashboards
ポート:
- 「3000:3000」
依存関係 :
- プロメテウス
Docker ではなく手動で Grafana をインストールしたい場合は、Ubuntu での Grafana のインストールと使用に関するステップバイステップ ガイドを参照してください。
Grafana データソース プロビジョニング ( grafana/provisioning/datasources/datasource.yml )
APIバージョン : 1
データソース :
- 名前 : プロメテウス
タイプ：プロメテウス
アクセス：プロキシ
URL : http://プロメテウス:9090
isDefault : true
ダッシュボード プロビジョニング ( grafana/provisioning/dashboards/dashboards.yml )
APIバージョン : 1
プロバイダー:
- 名前：「LLM」
フォルダ：「LLM」
タイプ: ファイル
無効な削除: true
オプション:
パス: /var/lib/grafana/dashboards
導入オプション 2: Kubernetes (Prometheus Operator + ServiceMonitor)
kube-prometheus-stack (Prometheus Operator) を使用する場合は、 ServiceMonitor 経由でターゲットをスクレイピングします。
Kubernetes、シングルノード Docker、マネージド推論プロバイダーの間のインフラストラクチャのトレードオフについては、
私のことを見てください
2026 年の LLM ホスティング ガイド 。
1) サービスを使用して推論デプロイメントを公開する
APIバージョン : v1
種類 : サービス
メタデータ:
名前：tgi
ラベル:
アプリ：tgi
仕様：
セレクター:
アプリ：tgi
ポート:
- 名前: http
ポート: 8080
ターゲットポート: 8080
2) ServiceMonitor を作成する
apiバージョン:monitoring.coreos.com/v1
種類 : ServiceMonitor
メタデータ:
名前：tgi
ラベル:
リリース: kube-prometheus-stack
仕様：
セレクター:
matchLabels :
アプリ：tgi
エンドポイント:
- ポート: http
パス: /metrics
間隔：5秒
vLLM サービスと llama.cpp サービスについてもこの手順を繰り返します。これにより、レプリカを追加するときれいに拡張されます。
3) アラート: SLO スタイルのルール (例)
キュー時間 p99 が長すぎます (ユーザーが待機中)
KV キャッシュ使用率 > 90% が継続 (容量クリフ)
ルールの例 (p95 リクエスト期間):
- アラート : LLMHighP95Latency
expr : 履歴

gram_quantile(0.95, sum by (le) (rate(tgi_request_duration_bucket[5m]))) > 3
用：10m
ラベル:
重大度 : ページ
注釈 :
概要 : 「TGI p95 レイテンシー > 3 秒 (10 メートル)」
トラブルシューティング: LLM スタックでの一般的な Prometheus + Grafana の障害
1) プロメテウスのターゲットは「DOWN」です
Prometheus UI → Targets が DOWN と表示される
「コンテキストの期限を超えました」または接続が拒否されました
サーバーは実際に /metrics を公開していますか?
ポートが間違っていますか?スキームが間違っていますか (http と https)?
Kubernetes: サービスはポッドを選択していますか? ServiceMonitor ラベルのリリースは正しいですか?
カール -sS http://tgi:8080/metrics |頭
2) メトリクスを収集できますが、パネルが空です
間違ったメトリクス名 (サーバーのバージョンが変更されました)
ダッシュボードには _bucket が必要ですが、ゲージ/カウンターしかありません
Prometheus のスクレイピング間隔が短いウィンドウには長すぎます (例: 30 秒のスクレイピングでは [1m] はノイズが多い可能性があります)
Grafana Explore を使用してメトリック プレフィックス (例: tgi_ / vllm: ) を検索します。
範囲ウィンドウを [1m] → [5m] に拡大
3) ヒストグラムのパーセンタイルが「平坦」または間違っているように見える
Prometheus ヒストグラムには正しい集計が必要です。
次に、(le) (およびオプションで他の安定したラベル) によって合計します。
Prometheus は、バケット モデルとサーバー側の分位数の計算を文書化します。
Grafana のヒストグラム視覚化ガイドには、実用的なパネル パターンが含まれています。
4) カーディナリティ爆発 (プロメテウスのメモリスパイク)
あなたは追加します

[切り捨てられた]

## Original Extract

Learn how to monitor LLM inference in production using Prometheus and Grafana. Track p95 latency, tokens/sec, queue duration, and KV cache usage across vLLM, TGI, and llama.cpp. Includes PromQL examples, dashboards, alerts, Docker & Kubernetes setups.

Notes on the margins
Rost Glukhov. Personal site and technical blog
Menu
AI Systems
Monitor LLM Inference in Production (2026): Prometheus & Grafana for vLLM, TGI, llama.cpp
Monitor LLM with Prometheus and Grafana
LLM inference looks like “just another API” — until latency spikes, queues back up, and your GPUs sit at 95% memory with no obvious explanation.
Monitoring becomes mission-critical the moment you move beyond a single-node setup or start optimizing for throughput. At that point, traditional API metrics aren’t enough. You need visibility into tokens, batching behavior, queue time, and KV cache pressure - the real bottlenecks of modern LLM systems.
This article is part of my broader observability and monitoring guide , where I cover monitoring vs observability fundamentals, Prometheus architecture, and production best practices. Here, we’ll focus specifically on monitoring LLM inference workloads .
(If you’re deciding on infrastructure, see my guide to LLM hosting in 2026 . If you want a deep dive into batching mechanics, VRAM limits, and throughput vs latency trade-offs, see the LLM performance engineering guide .)
Unlike typical REST services, LLM serving is shaped by tokens , continuous batching , KV cache utilization , GPU/CPU saturation , and queue dynamics . Two requests with identical payload sizes can have radically different latency depending on max_new_tokens , concurrency , and cache reuse .
This guide is a practical, production-focused walkthrough for building LLM inference monitoring with Prometheus and Grafana :
What to measure (p95/p99 latency, tokens/sec, queue duration, cache utilization, error rate)
How to scrape /metrics from common servers ( vLLM , Hugging Face TGI , llama.cpp )
PromQL examples for percentiles, saturation, and throughput
Deployment patterns with Docker Compose and Kubernetes
Troubleshooting the issues that only appear under real load
The examples are intentionally vendor-neutral. Whether you later add OpenTelemetry tracing, autoscaling, or a service mesh, the same metric model applies.
Why you should monitor LLM inference differently
Traditional API monitoring (RPS, p95 latency, error rate) is necessary but not sufficient. LLM serving adds additional axes:
E2E latency : time from request received → final token returned.
Inter-token latency : time per token during decode (critical for streaming UX).
Some servers expose both. For example, TGI exposes request duration and mean time-per-token as histograms.
2) Throughput is in tokens, not requests
A “fast” service that returns 5 tokens is not comparable to one returning 500 tokens. Your “RPS” should often be “ tokens/sec ”.
If you run continuous batching, queue depth is what you sell. Watching queue duration and queue size tells you whether you’re meeting user expectations.
4) Cache pressure is an outage precursor
KV cache exhaustion (or fragmentation) often shows up as sudden latency spikes and timeouts. vLLM exposes KV cache usage as a gauge.
Metrics checklist for LLM inference monitoring
Use this as your north star. You don’t need everything on day one—but you’ll want most of it eventually.
Traffic: requests/sec, tokens/sec
Errors: error rate, timeouts, OOMs, 429s (rate limiting)
Latency: p50/p95/p99 request duration; prefill vs decode latency; inter-token latency
Saturation: GPU utilization, memory usage, KV cache usage, queue size
If you need low-level visibility into GPU memory usage, temperature, and utilization outside of Prometheus (for debugging or single-node setups), see my guide to GPU monitoring applications in Linux / Ubuntu .
For a broader view of LLM observability beyond metrics — including tracing, structured logs, synthetic testing, GPU profiling, and SLO design — see my in-depth guide on observability for LLM systems .
Keep label cardinality low. Good labels:
model , endpoint , method (prefill/decode), status (success/error), instance
raw prompt , raw user_id , request ids — these explode series count.
Exposing metrics: built-in /metrics endpoints (vLLM, TGI, llama.cpp)
The easiest path is: use the metrics the server already exposes .
vLLM: Prometheus-compatible /metrics
vLLM exposes a Prometheus-compatible /metrics endpoint (via its Prometheus metrics logger) and publishes server/request metrics with the vllm: prefix, including gauges like running requests and KV cache usage.
For container setup, OpenAI-compatible serving, and throughput-oriented runtime tuning, see vLLM Quickstart: High-Performance LLM Serving .
Example metrics you’ll typically see:
Hugging Face TGI: /metrics with queue + request histograms
TGI exposes many production-grade metrics on /metrics , including queue size, request duration, queue duration, and mean time per token.
tgi_request_duration (histogram, e2e latency)
tgi_request_queue_duration (histogram)
tgi_request_mean_time_per_token_duration (histogram)
Operational setup—Docker, GPUs, launch flags, and the failures that show up as empty or misleading scrapes—is covered in TGI - Text Generation Inference - Install, Config, Troubleshoot .
llama.cpp server: enable metrics endpoint
The llama.cpp server supports a Prometheus-compatible metrics endpoint that must be enabled with a flag (e.g., --metrics ).
For installation paths, key runtime flags, and OpenAI-compatible server usage, see llama.cpp Quickstart with CLI and Server .
If you’re running llama.cpp behind a proxy, scrape the server directly whenever possible (to avoid proxy-level latency hiding the actual inference behavior).
Prometheus configuration: scraping your inference servers
vLLM at http://vllm:8000/metrics
TGI at http://tgi:8080/metrics
llama.cpp at http://llama:8080/metrics
scrape interval tuned for fast feedback
global :
scrape_interval : 5s
evaluation_interval : 15s
scrape_configs :
- job_name : "vllm"
metrics_path : /metrics
static_configs :
- targets : [ "vllm:8000" ]
- job_name : "tgi"
metrics_path : /metrics
static_configs :
- targets : [ "tgi:8080" ]
- job_name : "llama_cpp"
metrics_path : /metrics
static_configs :
- targets : [ "llama:8080" ]
If you’re new to Prometheus or want a deeper explanation of scrape configs, exporters, relabeling, and alerting rules, see my full Prometheus monitoring setup guide .
Pro tip: add a “service label”
If you run multiple models/replicas, add relabeling to include a stable service label for dashboards.
relabel_configs :
- target_label : service
replacement : "llm-inference"
PromQL examples you can copy/paste
sum ( rate ( tgi_request_count[ 5m ] ))
For vLLM, use its request counters (names vary by version), but the pattern is the same: sum(rate(<counter>[5m])) .
If you have *_success counters, compute failure ratio:
1 - (
sum ( rate ( tgi_request_success[ 5m ] ))
/
sum ( rate ( tgi_request_count[ 5m ] ))
)
p95 latency for histogram metrics (Prometheus)
Prometheus histograms are bucketed counts; use histogram_quantile() over rate() of the buckets. Prometheus documents this model and the histogram vs summary tradeoffs.
histogram_quantile (
0.95 ,
sum by ( le ) ( rate ( tgi_request_duration_bucket[ 5m ] ))
)
p99 queue time
histogram_quantile (
0.99 ,
sum by ( le ) ( rate ( tgi_request_queue_duration_bucket[ 5m ] ))
)
Mean time per token (inter-token latency)
histogram_quantile (
0.95 ,
sum by ( le ) ( rate ( tgi_request_mean_time_per_token_duration_bucket[ 5m ] ))
)
Inter-token latency is often constrained by decode bottlenecks and memory bandwidth - topics covered in detail in
LLM performance optimization guide .
max ( tgi_queue_size )
vLLM KV cache utilization (instant)
max ( vllm : kv_cache_usage_perc )
Grafana dashboards: panels that actually help on-call
Grafana can visualize histograms in multiple ways (percentiles, heatmaps, bucket distributions). Grafana Labs has a detailed guide to Prometheus histogram visualization.
A minimal, high-signal dashboard layout:
p95 request latency (time series)
p95 inter-token latency (time series)
Error rate (time series + stat)
Row 2 — Capacity and saturation
Running vs waiting requests (stacked)
Generated tokens/request (p50/p95)
If you have streaming, add a panel for “first token latency” (TTFT) when available.
p95 latency panel: the histogram_quantile(0.95, …) query above
heatmap panel: graph the bucket rates ( *_bucket ) as a heatmap (Grafana supports this approach)
Deployment option 1: Docker Compose (fast local + single-node)
If you’re deciding between local, self-hosted, or cloud-based inference architectures, see the full breakdown in my
LLM hosting comparison guide .
monitoring/
docker-compose.yml
prometheus/
prometheus.yml
grafana/
provisioning/
datasources/datasource.yml
dashboards/dashboards.yml
dashboards/
llm-inference.json
docker-compose.yml
services :
prometheus :
image : prom/prometheus:latest
volumes :
- ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml:ro
ports :
- "9090:9090"
grafana :
image : grafana/grafana:latest
environment :
- GF_SECURITY_ADMIN_USER=admin
- GF_SECURITY_ADMIN_PASSWORD=admin
volumes :
- ./grafana/provisioning:/etc/grafana/provisioning
- ./grafana/dashboards:/var/lib/grafana/dashboards
ports :
- "3000:3000"
depends_on :
- prometheus
If you prefer a manual Grafana installation instead of Docker, see my step-by-step guide on installing and using Grafana on Ubuntu .
Grafana datasource provisioning ( grafana/provisioning/datasources/datasource.yml )
apiVersion : 1
datasources :
- name : Prometheus
type : prometheus
access : proxy
url : http://prometheus:9090
isDefault : true
Dashboard provisioning ( grafana/provisioning/dashboards/dashboards.yml )
apiVersion : 1
providers :
- name : "LLM"
folder : "LLM"
type : file
disableDeletion : true
options :
path : /var/lib/grafana/dashboards
Deployment option 2: Kubernetes (Prometheus Operator + ServiceMonitor)
If you use kube-prometheus-stack (Prometheus Operator), scrape targets via ServiceMonitor .
For infrastructure trade-offs between Kubernetes, single-node Docker, and managed inference providers,
see my
LLM hosting in 2026 guide .
1) Expose your inference deployment with a Service
apiVersion : v1
kind : Service
metadata :
name : tgi
labels :
app : tgi
spec :
selector :
app : tgi
ports :
- name : http
port : 8080
targetPort : 8080
2) Create a ServiceMonitor
apiVersion : monitoring.coreos.com/v1
kind : ServiceMonitor
metadata :
name : tgi
labels :
release : kube-prometheus-stack
spec :
selector :
matchLabels :
app : tgi
endpoints :
- port : http
path : /metrics
interval : 5s
Repeat for vLLM and llama.cpp services. This scales cleanly as you add replicas.
3) Alerting: SLO-style rules (example)
Queue time p99 too high (users waiting)
KV cache usage > 90% sustained (capacity cliff)
Example rule (p95 request duration):
- alert : LLMHighP95Latency
expr : histogram_quantile(0.95, sum by (le) (rate(tgi_request_duration_bucket[5m]))) > 3
for : 10m
labels :
severity : page
annotations :
summary : "TGI p95 latency > 3s (10m)"
Troubleshooting: common Prometheus + Grafana failures in LLM stacks
1) Prometheus target is “DOWN”
Prometheus UI → Targets shows DOWN
“context deadline exceeded” or connection refused
Is the server actually exposing /metrics ?
Wrong port? Wrong scheme (http vs https)?
Kubernetes: is the Service selecting pods? Is the ServiceMonitor label release correct?
curl -sS http://tgi:8080/metrics | head
2) You can scrape metrics, but panels are empty
Wrong metric name (server version changed)
Dashboard expects _bucket but you only have a gauge/counter
Prometheus scrape interval too long for short windows (e.g., [1m] with 30s scrape can be noisy)
Use Grafana Explore to search metric prefixes (e.g., tgi_ / vllm: )
Increase range window from [1m] → [5m]
3) Histogram percentiles look “flat” or wrong
Prometheus histograms require correct aggregation:
then sum by (le) (and optionally other stable labels)
Prometheus documents the bucket model and server-side quantile calculation.
Grafana’s histogram visualization guide includes practical panel patterns.
4) Cardinality explosion (Prometheus memory spikes)
You adde

[truncated]

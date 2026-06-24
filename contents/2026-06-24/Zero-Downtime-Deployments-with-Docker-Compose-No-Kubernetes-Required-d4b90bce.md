---
source: "https://statusdude.com/blog/zero-downtime-docker-compose"
hn_url: "https://news.ycombinator.com/item?id=48665130"
title: "Zero-Downtime Deployments with Docker Compose – No Kubernetes Required"
article_title: "Zero-Downtime Deployments with Docker Compose — No Kubernetes Required | StatusDude"
author: "canto"
captured_at: "2026-06-24T20:33:19Z"
capture_tool: "hn-digest"
hn_id: 48665130
score: 3
comments: 0
posted_at: "2026-06-24T20:20:33Z"
tags:
  - hacker-news
  - translated
---

# Zero-Downtime Deployments with Docker Compose – No Kubernetes Required

- HN: [48665130](https://news.ycombinator.com/item?id=48665130)
- Source: [statusdude.com](https://statusdude.com/blog/zero-downtime-docker-compose)
- Score: 3
- Comments: 0
- Posted: 2026-06-24T20:20:33Z

## Translation

タイトル: Docker Compose を使用したダウンタイムゼロのデプロイ – Kubernetes は不要
記事のタイトル: Docker Compose によるゼロダウンタイムのデプロイメント — Kubernetes は不要 | StatusDude
説明: ダウンタイムゼロの展開のために Traefik を試しました。 It dropped requests, returned 404s, and couldn

記事本文:
Zero-Downtime Deployments with Docker Compose — No Kubernetes Required |ステータスデュード
ステータスデュード
Features Agents Pricing Blog
始めましょう
Features Agents Pricing Blog
始めましょう
← ブログに戻る
Zero-Downtime Deployments with Docker Compose — No Kubernetes Required
There's a mass delusion in the industry that you need Kubernetes to run a serious production service.あなたはしない。 StatusDude では、毎分数千件の監視チェックを実行し、マルチリージョン ワーカーを実行し、1 日に複数回デプロイします。これらはすべて Docker Compose と HAProxy を使用しています。ドロップされたリクエストはゼロです。ダウンタイムゼロ。 No etcd to babysit at 3 AM.
But we didn't start with HAProxy.私たちはTraefikから始めました。 That lasted about four hours.
Traefik is the popular choice for Docker-based setups. It auto-discovers services via Docker labels, has a slick dashboard, and the docs make it look effortless. We set up two backend replicas with Traefik labels, ran a rolling deploy, and watched everything fall apart.
"Service defined multiple times"
Our first deploy strategy was to run a backend_new service alongside the existing backend during the transition. Both had the same Traefik routing labels — same Host rule, same service definition.当然ですよね？ You want both old and new to serve traffic during the cutover.
トレーフィク氏はこれに同意しなかった。 Its Docker provider treats each Compose service as a separate configuration source. Two services with the same labels? "Service defined multiple times."すべてのリクエストで 404。 No fallback, no merge, just a flat refusal to route anything.
We reworked the approach to use docker compose --scale backend=4 instead of a separate service. That avoided the label conflict. But it uncovered the next problem.
The rolling deploy strategy: scale up to 4 replicas (2 old + 2 new), then s

スケールを 2 に戻します (新しいもののみを保持します)。十分シンプルです。
ただし、Traefik の内部ルーティング テーブルの更新速度が十分ではありませんでした。 4 から 2 にスケールダウンし、Traefik はシャットダウン中のコンテナへのルーティングを継続します。他のリクエストごとに 502。ルーティング状態は Docker の現実よりも数秒遅れていました。これはかなりの量のトラフィックをドロップするのに十分な長さです。
遅延を追加してみました。コンテナーを停止する前に、ネットワークからコンテナーを切断しようとしました (削除前にヘルスチェックが正常に失敗するようにします)。私たちはパッシブなヘルスチェックを試しましたが、追加しましたが、攻撃的すぎて誤検知を引き起こしたため、すぐにロールバックされました。
どれもきれいではありませんでした。しかし、本当の殺人者はまったく別の人物でした。
キラー: 別のバックエンドでの再試行は禁止
これは既知の問題ですが、開発者はしばらく無視しているようです... https://github.com/traefik/traefik/issues/2723
シナリオは次のとおりです。ローリング デプロイ中に、古いコンテナーを停止します。 docker stop は SIGTERM を送信します。 Uvicorn は正常なシャットダウンを開始しますが、ウィンドウが存在します。つまり、すでに処理中のリクエスト、または停止信号と Traefik がルーティング テーブルを更新するまでの間に到着するリクエストです。
そのリクエストが停止中のバックエンドに到達すると、接続はストリームの途中で切断されます。クライアントは生のエラー (空の応答、接続のリセット、本文の一部) を受け取ります。
そんなことはあり得ません。サービスとハートビート モニターが稼働していることを報告するときは、それを確認する必要があります。
ここで、失敗したリクエストに対して Traefik が行う処理は次のとおりです。何も行いません。
Traefik の再試行ミドルウェアは存在しますが、同じバックエンドで再試行されます。死につつある人。また失敗するやつ。正常なバックエンドには再ディスパッチされません。リクエストはただ...失われています。
パッシブヘルスチェック、切断前など、あらゆる組み合わせを試しました。

非常停止、試行回数が異なる再試行ミドルウェア。根本的な問題は依然として残っていました。Traefik は失敗したリクエストを別のサーバーに送信できませんでした。
その日の午後、私たちは Traefik を削除し、HAProxy に手を伸ばしました。
それを剥ぎ取りましょう。ゼロダウンタイム導入には実際に何が必要ですか?
複数のバックエンド インスタンス - トラフィックを処理している間に 1 つを置き換えることができます
別のバックエンドで再試行するロード バランサ - 停止中のコンテナがリクエストをドロップしないようにする
インスタンスを一度に 1 つずつ置き換えるデプロイ スクリプト - ローリング アップデート
それだけです。 3つのこと。それぞれのやり方を紹介しましょう。
ステップ 1: Docker Compose を使用した複数のレプリカ
Docker Compose には、deploy.replicas 設定が組み込まれています。
# docker-compose.yml
サービス:
バックエンド:
ビルド: ./バックエンド
デプロイ:
レプリカ: 2
画像: myapp-バックエンド
暴露する:
- 「8000」
環境ファイル: .env
ヘルスチェック:
テスト: ["CMD"、"curl"、"-f"、"http://localhost:8000/health"]
間隔: 5秒
タイムアウト: 5秒
再試行: 3
開始期間: 5秒
再起動: 停止しない限り
これは、共有 Docker DNS 名 backend の背後で実行されている 2 つのバックエンド コンテナーです。 Docker ネットワーク内のバックエンドを解決すると、両方のコンテナー IP が取得されます。
1 つの Dockerfile 、1 つのイメージ、2 つのコンテナー。ポッド仕様、デプロイメント、レプリカ セットはありません。
ステップ 2: ロード バランサーとしての HAProxy
HAProxy は十分にテストされており、高速で、構成は読み取り可能です。しかし、私たちがそれを選んだ本当の理由は、オプション redispatch です。
グローバル
ログ標準出力形式の生の local0 情報
マックスコン 4096
デフォルト
モードhttp
タイムアウト接続 3 秒
タイムアウトクライアント 30秒
タイムアウトサーバー30秒
# 重要な機能: 失敗したリクエストを別のバックエンドで再試行します
再試行 3
オプション再発送1
再試行接続失敗 空の応答 応答タイムアウト 502 503 504
リゾルバ docker_dns
ネームサーバー dns1 127.0.0.11:53
解決再試行 3
タイムアウト解決 1 秒
タイムアウト再試行 1 秒
# 再解決

2 秒ごとに DNS を実行する
有効な 2 を保持します
フロントエンド http_in
バインド*:80
default_backend バックエンド
バックエンド バックエンド
バランスラウンドロビン
オプション httpchk
http-check send meth GET uri /health
http-check 期待ステータス 200
デフォルトサーバー間 1 秒降下 1 上昇 1 チェックリゾルバー docker_dns \
解決優先 ipv4 init-addr なし \
Layer7 エラー制限 3 のエラー マークダウンを観察します。
サーバー テンプレート バックエンド 1-10 バックエンド:8000 チェック
この機能を実現する 3 つの要素について話しましょう。
これは、Traefik が提供できなかった機能です。
再試行 3
オプション再発送1
再試行接続失敗 空の応答 応答タイムアウト 502 503 504
リクエストが失敗した場合 (接続拒否、空の応答、502、503、504)、HAProxy はリクエストを再試行します。また、オプション redispatch 1 は、すべての再試行が異なるバックエンドに行われることを意味します。同じ瀕死のサーバーではありません。別の、健康的なもの。
ローリング デプロイ中に、シャットダウン中のコンテナにリクエストが到達し、空の応答を受け取った場合、HAProxy は他のレプリカでサイレントに再試行します。クライアントにはエラーが表示されません。ドロップされたリクエストはありません。この 1 つの機能により、Traefik で発生したすべての問題が解消されました。
3 層の健康状態検出
私たちは単一のヘルスチェックメカニズムに依存しません。 3 つの独立したレイヤーがあり、それぞれが異なる障害モードを検出します。
レイヤ 1 — リクエストごとの再試行 (ミリ秒): 1 つのリクエストが失敗した場合は、別のバックエンドですぐに再試行します。デプロイ中の一時的な障害を捕捉します。
レイヤ 2 — パッシブ観察 (layer7 を観察): HAProxy は、実際のトラフィックからの実際の HTTP 応答を監視します。バックエンドが 3 回連続して 5xx エラー ( error-limit 3 ) を返した場合、そのバックエンドは即座にローテーションから除外されます ( on-error mark-down )。プローブ サイクルを待つ必要はありません。
レイヤ 3 — アクティブなヘルス チェック ( inter 1s fall 1 raise 1 ): /health エンドポイントを毎秒プローブします。キャッチコンプ

トラフィックを受信しない、完全に停止したバックエンド。 1 回の失敗 = 即時ダウン。 1 回の成功 = ローテーションに戻ります。
各レイヤーは他のレイヤーの死角をカバーします。リクエストごとの再試行は、停止中のバックエンドにヒットした 1 つのリクエストを処理します。パッシブ チェックは、負荷がかかるとエラーを返し始めるバックエンドを処理します。アクティブ チェックは、トラフィックが流れずに静かにクラッシュするバックエンドを処理します。
DNS ベースの検出 (Docker ソケットなし)
server-template backend 1-10 backend:8000 check 行は、HAProxy がバックエンドを検出する方法です。 Docker の組み込み DNS リゾルバー ( 127.0.0.11:53 ) を使用して Docker DNS 名バックエンドを解決し、検出した IP ごとにサーバー エントリを作成します。
有効な 2 秒を保持するということは、HAProxy が 2 秒ごとに再解決されることを意味します。コンテナが死ぬ？その IP が DNS から消えます。新しいコンテナが始まりますか?そのIPが表示されます。 HAProxy はそれを自動的に取得します。
Docker ソケット マウントはありません。ラベルの解析は行われません。動的構成の生成はありません。そのまま動作する静的な構成ファイル。サービスメッシュはありません。サイドカーはありません。オペレーターはいません。
さすが。
これはデプロイ スクリプト全体です。
プロッドデプロイ:
@echo "=== ダウンタイムゼロのローリングデプロイ ===
@for cid in $$(docker compose -f docker-compose.prod.yml ps -q backend); \してください
echo "$$cid を置換しています..."; \
docker stop $$cid && docker rm -f $$cid; \
docker compose -f docker-compose.prod.yml up -d --no-deps --no-recreate --wait バックエンド; \
完了しました
@echo "=== デプロイ完了 ===
それだけです。何が起こるかを説明しましょう:
実行中のすべてのバックエンド レプリカのコンテナ ID を取得する
レプリカごとに、一度に 1 つずつ以下を実行します。
容器を止めて取り除く
HAProxy は失われたバックエンドを 2 秒以内に検出します (DNS 再解決)
トラフィックは残りの正常なレプリカに移行します
更新されたイメージを使用して新しいコンテナを開始します
--Docker のヘルスチェックに合格するまでブロックを待機します
HAProxy は DNS 経由で新しいバックエンドを検出します
交通が流れ始める

新しいコンテナに
デプロイ中のあらゆる時点で、少なくとも 1 つの正常なバックエンドがトラフィックを処理しています。 --no-recreate フラグは、Docker がまだ置き換えられていないレプリカに触れるのを防ぎます。
2 秒間の DNS ウィンドウ中に、瀕死のコンテナにヒットしたリクエストはありますか?正常なレプリカで自動的に再試行されます。クライアントには決してわかりません。
StatusDude では、このセットアップは以下を処理します。
3 つのリージョンにわたる毎分数千の監視チェック
リクエストのドロップなしで 1 日に複数回デプロイ
バックエンドがダウンした場合の 2 秒未満のフェイルオーバー
約 60 行の HAProxy 構成と 10 行のデプロイ スクリプト
私たちは、Traefik (404、502、リクエストのドロップ、4 時間のデバッグ) から HAProxy (リクエストのドロップなし、最初のデプロイ) まで 1 日の午後で完了しました。場合によっては、退屈で実績のあるツールが正しい選択となることもあります。まあ、「時々」ではなく、よくあることです ;-)
追伸
Nginx でもいいでしょう。今回は haproxy を起動したいと思っただけです :)
稼働時間の監視に StatusDude を信頼している何千もの開発者に加わりましょう。
© 2026 ステータスデュード.あなたはここに留まります。男が電話するよ。

## Original Extract

We tried Traefik for zero-downtime deploys. It dropped requests, returned 404s, and couldn

Zero-Downtime Deployments with Docker Compose — No Kubernetes Required | StatusDude
Status Dude
Features Agents Pricing Blog
Get Started
Features Agents Pricing Blog
Get Started
← Back to Blog
Zero-Downtime Deployments with Docker Compose — No Kubernetes Required
There's a mass delusion in the industry that you need Kubernetes to run a serious production service. You don't. At StatusDude, we serve thousands of monitoring checks per minute, run multi-region workers, and deploy multiple times a day — all with Docker Compose and HAProxy . Zero dropped requests. Zero downtime. No etcd to babysit at 3 AM.
But we didn't start with HAProxy. We started with Traefik. That lasted about four hours.
Traefik is the popular choice for Docker-based setups. It auto-discovers services via Docker labels, has a slick dashboard, and the docs make it look effortless. We set up two backend replicas with Traefik labels, ran a rolling deploy, and watched everything fall apart.
"Service defined multiple times"
Our first deploy strategy was to run a backend_new service alongside the existing backend during the transition. Both had the same Traefik routing labels — same Host rule, same service definition. Makes sense, right? You want both old and new to serve traffic during the cutover.
Traefik disagreed. Its Docker provider treats each Compose service as a separate configuration source. Two services with the same labels? "Service defined multiple times." 404 on every request. No fallback, no merge, just a flat refusal to route anything.
We reworked the approach to use docker compose --scale backend=4 instead of a separate service. That avoided the label conflict. But it uncovered the next problem.
The rolling deploy strategy: scale up to 4 replicas (2 old + 2 new), then scale back down to 2 (keeping only the new ones). Simple enough.
Except Traefik's internal routing table didn't update fast enough. We'd scale down from 4 to 2, and Traefik would keep routing to containers that were in the process of shutting down. 502s on every other request. The routing state lagged behind Docker's reality by several seconds — long enough to drop a significant chunk of traffic.
We tried adding delays. We tried disconnecting containers from the network before stopping them (so the health check would fail cleanly before removal). We tried passive health checks — added them, then immediately rolled them back because they were too aggressive and caused false positives.
None of it was clean. But the real killer was something else entirely.
The Killer: No Retry on a Different Backend
That's a known issue that devs seem to ignore for a while now... https://github.com/traefik/traefik/issues/2723
Here's the scenario: during a rolling deploy, you stop an old container. docker stop sends SIGTERM. Uvicorn starts its graceful shutdown, but there's a window — requests that are already in-flight, or requests that arrive between the stop signal and Traefik updating its routing table.
When that request hits the dying backend, the connection drops mid-stream. The client gets a raw error — empty response, connection reset, partial body.
We can't have that. When you report your service and heartbeat monitors are up - we need to acknowledge!
Now here's what Traefik does with that failed request: nothing .
Traefik's retry middleware exists, but it retries on the same backend . The one that's dying. The one that will fail again. It doesn't redispatch to a healthy backend. The request is just... lost.
We tried every combination: passive health checks, disconnect-before-stop, retry middleware with different attempts counts. The fundamental problem remained — Traefik couldn't send a failed request to a different server.
That afternoon, we ripped out Traefik and reached for HAProxy.
Let's strip it down. What does zero-downtime deployment actually require?
Multiple backend instances — so you can replace one while the other serves traffic
A load balancer that retries on a different backend — so dying containers don't drop requests
A deploy script that replaces instances one at a time — rolling update
That's it. Three things. Let me show you how we do each one.
Step 1: Multiple Replicas with Docker Compose
Docker Compose has a built-in deploy.replicas setting:
# docker-compose.yml
services:
backend:
build: ./backend
deploy:
replicas: 2
image: myapp-backend
expose:
- "8000"
env_file: .env
healthcheck:
test: ["CMD", "curl", "-f", "http://localhost:8000/health"]
interval: 5s
timeout: 5s
retries: 3
start_period: 5s
restart: unless-stopped
That's 2 backend containers running behind a shared Docker DNS name backend . When you resolve backend inside the Docker network, you get both container IPs.
One Dockerfile , one image, two containers. No pod specs, no deployments, no replica sets.
Step 2: HAProxy as the Load Balancer
HAProxy is battle-tested, fast, and the configuration is readable. But the real reason we chose it: option redispatch .
global
log stdout format raw local0 info
maxconn 4096
defaults
mode http
timeout connect 3s
timeout client 30s
timeout server 30s
# THE key feature: retry failed requests on a DIFFERENT backend
retries 3
option redispatch 1
retry-on conn-failure empty-response response-timeout 502 503 504
resolvers docker_dns
nameserver dns1 127.0.0.11:53
resolve_retries 3
timeout resolve 1s
timeout retry 1s
# Re-resolve DNS every 2 seconds
hold valid 2s
frontend http_in
bind *:80
default_backend backends
backend backends
balance roundrobin
option httpchk
http-check send meth GET uri /health
http-check expect status 200
default-server inter 1s fall 1 rise 1 check resolvers docker_dns \
resolve-prefer ipv4 init-addr none \
observe layer7 error-limit 3 on-error mark-down
server-template backend 1-10 backend:8000 check
Let's talk about the three things that make this work.
This is the feature that Traefik couldn't deliver:
retries 3
option redispatch 1
retry-on conn-failure empty-response response-timeout 502 503 504
When a request fails — connection refused, empty response, 502, 503, 504 — HAProxy retries it. And option redispatch 1 means every retry goes to a different backend . Not the same dying server. A different, healthy one.
During a rolling deploy, if a request hits a container that's shutting down and gets an empty response, HAProxy silently retries on the other replica. The client never sees the error. No dropped requests. This single feature eliminated every problem we had with Traefik.
Three Layers of Health Detection
We don't rely on a single health check mechanism. There are three independent layers, each catching different failure modes:
Layer 1 — Per-request retry (milliseconds): If a single request fails, retry immediately on a different backend. Catches transient failures during deploys.
Layer 2 — Passive observation ( observe layer7 ): HAProxy watches actual HTTP responses from real traffic. If a backend returns 3 consecutive 5xx errors ( error-limit 3 ), it's pulled from rotation instantly ( on-error mark-down ). No waiting for any probe cycle.
Layer 3 — Active health checks ( inter 1s fall 1 rise 1 ): Probes the /health endpoint every second. Catches completely dead backends that receive no traffic. One failure = instant DOWN. One success = back in rotation.
Each layer covers a blind spot of the others. Per-request retry handles the single request that hits a dying backend. Passive checks handle backends that start returning errors under load. Active checks handle backends that crash silently with no traffic flowing to them.
DNS-Based Discovery (No Docker Socket)
The server-template backend 1-10 backend:8000 check line is how HAProxy discovers backends. It resolves the Docker DNS name backend using Docker's embedded DNS resolver ( 127.0.0.11:53 ) and creates server entries for each IP it finds.
The hold valid 2s means HAProxy re-resolves every 2 seconds. Container dies? Its IP disappears from DNS. New container starts? Its IP appears. HAProxy picks it up automatically.
No Docker socket mount. No label parsing. No dynamic config generation. A static config file that just works. No service mesh. No sidecar. No operator.
Srsly.
This is the entire deploy script:
prod-deploy:
@echo "=== Zero-downtime rolling deploy ==="
@for cid in $$(docker compose -f docker-compose.prod.yml ps -q backend); do \
echo "Replacing $$cid..."; \
docker stop $$cid && docker rm -f $$cid; \
docker compose -f docker-compose.prod.yml up -d --no-deps --no-recreate --wait backend; \
done
@echo "=== Deploy complete ==="
That's it. Let me walk through what happens:
Get the container IDs of all running backend replicas
For each replica, one at a time:
Stop and remove the container
HAProxy detects the missing backend within 2 seconds (DNS re-resolution)
Traffic shifts to the remaining healthy replica
Start a new container with the updated image
--wait blocks until Docker's healthcheck passes
HAProxy discovers the new backend via DNS
Traffic starts flowing to the new container
At every point during the deploy, at least one healthy backend is serving traffic . The --no-recreate flag prevents Docker from touching the replica we haven't replaced yet.
Any requests that hit the dying container during that 2-second DNS window? Retried on the healthy replica automatically. The client never knows.
At StatusDude, this setup handles:
Thousands of monitoring checks per minute across 3 regions
Multiple deploys per day with zero dropped requests
Sub-2-second failover when a backend goes down
~60 lines of HAProxy config and a 10-line deploy script
We went from Traefik (404s, 502s, dropped requests, four hours of debugging) to HAProxy (zero dropped requests, first deploy) in one afternoon. Sometimes the boring, battle-tested tool is the right choice. Well, OK, not "sometimes" - quite often ;-)
P.S
Nginx would do too, I just felt like getting haproxy up this time :)
Join thousands of developers who trust StatusDude for their uptime monitoring.
© 2026 StatusDude. You abide. Dude will call.

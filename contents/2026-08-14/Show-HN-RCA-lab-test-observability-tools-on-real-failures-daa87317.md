---
source: "https://github.com/coroot/rca-lab"
hn_url: "https://news.ycombinator.com/item?id=49304191"
title: "Show HN: RCA-lab – test observability tools on real failures"
article_title: "GitHub - coroot/rca-lab: A Kubernetes lab that reproduces real production incidents on a live, instrumented microservice stack — for testing root-cause-analysis tools and agents. · GitHub"
author: "nikolay_sivko"
captured_at: "2026-08-14T21:16:47Z"
capture_tool: "hn-digest"
hn_id: 49304191
score: 1
comments: 0
posted_at: "2026-08-14T20:31:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RCA-lab – test observability tools on real failures

- HN: [49304191](https://news.ycombinator.com/item?id=49304191)
- Source: [github.com](https://github.com/coroot/rca-lab)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T20:31:13Z

## Translation

タイトル: Show HN: RCA-lab – 実際の障害に対する可観測性ツールのテスト
記事のタイトル: GitHub - coroot/rca-lab: 根本原因分析ツールとエージェントをテストするための、実際の運用インシデントをライブの計測済みマイクロサービス スタック上で再現する Kubernetes ラボ。 · GitHub
説明: 根本原因分析ツールとエージェントをテストするために、実際の運用インシデントをライブの計測済みマイクロサービス スタック上で再現する Kubernetes ラボ。 - コルート/rca-lab

記事本文:
GitHub - coroot/rca-lab: 根本原因分析ツールとエージェントをテストするための、実際の運用インシデントをライブの計測済みマイクロサービス スタック上で再現する Kubernetes ラボ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
コルート
/
RCAラボ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
55 コミット 55 コミット .github/ workflows .github/ workflows

デプロイ デプロイ ドキュメント ドキュメント オペレーター オペレーター シナリオ シナリオ スクリプト スクリプト サービス サービス test/ e2e テスト/ e2e バリアント バリアント .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md バージョン.yaml バージョン.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
根本原因分析 (RCA) を評価するための現実的で再現可能な障害ラボ
ライブ Kubernetes クラスター上のツール (人間または AI)。
ほとんどの RCA ベンチマークは、おもちゃの環境からの定型テレメトリを次のように再生します。
合成障害は機能フラグによって切り替えられます。 rca-lab は逆のアプローチをとります。
実際の多言語マイクロサービス スタック (Python、Go、Java、Node.js、Rust、PHP)
API ゲートウェイの背後で、継続的に負荷が生成されます。
本番グレードのオペレータによる実際のデータベース: PostgreSQL、MySQL、
Percona オペレーターを介した MongoDB、valkey-operator を介した Valkey クラスター、
Strimzi経由のKafka — with
シードされたデータ ボリューム。
実際の障害メカニズムのみ。アプリ内にカオスフラグはありません。 GC
圧力インシデントは、新しいイメージとして出荷された本物の割り当て回帰です
バージョンを変更し、後でロールバックします。データベース インシデントは分析ワークロードです
運用データベースに対して大量のクエリを実行する。トラフィックの急増は
実際にはトラフィックが増えます。
耐久性のある復帰。すべての失敗シナリオは、FailureScenario のカスタムです
オペレーターによって駆動されるリソースで、
オペレータが再起動しても、シナリオは終了、無効化、または削除されます。
豊富なテレメトリー、独自のバックエンドの導入。すべてのサービスが装備されています
OpenTelemetry SDK を使用: トレース、SDK によって発行されたメトリクス (JVM/ランタイム/HTTP)、および
ログ (標準出力および OTLP へ、トレース相関)。すべてはバンドルに流れます
デフォルトでデータを破棄する otel-collector — 任意の OTLP をポイントします
変数が 1 つあるバックエンド。
要件: クラスターを指す kubectl + helm (任意のディストリビューション)

寄付;
デフォルトの StorageClass、フルサイズ ラボのノード全体で最大 8 CPU / 16 GiB）。
git clone https://github.com/coroot/rca-lab && cd rca-lab
makedeploy # すべて: 演算子 → データベース → Kafka → アプリ → シード
単一ノードクラスター (kind/k3d/minikube):
デプロイを SINGLE_NODE=1 にする
テレメトリをどこかに送信します (Coroot または任意の OTLP エンドポイントなど)。
makedeploy OTLP_ENDPOINT=my-backend:4317
その他の変数: STORAGE_CLASS=<name> 、SEED_SIZE_GB=<n> (0 はシード処理をスキップ)、
OTLP_HEADERS=k=v 、YES=1 (確認プロンプトなし)。 makedeployを再実行する
冪等に収束します。これは、これらの設定を変更する方法でもあります。
make clean # KEEP_DATA=1 はデータベース ボリュームを維持します
失敗のトリガー
シナリオは Kubernetes カスタム リソースです。
kubectl が失敗するシナリオを取得する
kubectl パッチの失敗シナリオ pg-analytics-queries --type=merge -p ' {"spec":{"enabled":true}} '
または、Web UI を使用します。
kubectl ポートフォワード svc/rca-lab-operator 8080
UI には、カテゴリごとにグループ化されたすべてのシナリオが、重大度とライブ状態とともにリストされます。
クリックするだけでそれぞれを開始または停止できます。
各シナリオでは、そのメカニズムと RCA ツールのテレメトリの症状が文書化されています。
観察できるはずだ。シナリオは、cron スケジュールで実行することもできます。
固定期間 — シナリオ/ を参照してください。
rca-lab を共有クラスターまたは運用クラスターにインストールしないでください。シナリオ
オペレーターは、ネームスペース内のワークロードを意図的に低下させる権限を持っています。
すべてのシナリオは本物の現実世界のメカニズムを使用しており、決して人工的な障害ではありません
アプリ内でフラグを立て、永続的に元に戻します。それぞれに予期される症状が伴います
RCA ツールのドキュメントと採点ルーブリックを兼ねたリスト。
信頼性カテゴリは別の種類のテストです。他のシナリオ
RCA を行使する急性インシデント (遅延、エラー、飽和) である - 所与
症状が出たら原因を探ります。信頼性

シナリオには潜在的な、ゆっくりと燃え上がるリスクがある
(肥大化、古い統計、ブロックされたバキューム、レプリケーションの遅延、チェックポイントのプレッシャー)
多くの場合、初期段階ではユーザーに直接的な症状は生じません。彼らは積極的に運動します
検出 — ツールがリスクの発生前にそのリスクを警告するかどうか
停電。彼らが予想する症状は、事故ではなく早期警告の指標です
症状。
シナリオ
仕組み
RCA ツールが見つけるべきもの
pg-analytics-クエリ
分析レポートのワークロードは、アプリと同じ pgBouncer プールを介して、本番環境の PostgreSQL に対して大量のマルチ結合/集計クエリ (約 10 GB の製品テーブルのフル スキャン) を実行します。
製品カタログ/在庫サービスの待ち時間の増加。 PostgreSQL の CPU/IO の飽和。分析レポートのワークロードに起因する、pg_stat_statements 内の新しいフルスキャン クエリのフィンガープリント。
pg-排他的ロック
停止したスキーマ移行トランザクションは、products テーブル ( LOCK TABLE ) に実際の ACCESS EXCLUSIVE ロックを取得し、それを保持したままハングします。これは、「移行がロックを取得し、決して解放しない」インシデントです。
製品に対する製品カタログのクエリはロックでブロックされます。接続プールがいっぱいになり、サービスが利用できなくなるため、API ゲートウェイ製品のエンドポイント エラーが発生しますが、何も実行されていないため、PostgreSQL の CPU/IO はフラットのままです。問題は、リソースの飽和ではなく、ロック待機 ( pg_locks / pg_blocking_pids ) です。
mysql-ロック競合
ストールしたトランザクションは、注文テーブルのホットエンド ( SELECT … FOR UPDATE 、最大 ID を超えるギャップを含む) で InnoDB 行ロックを保持し、その後ハングします。これは、ロックを取得してストールしたトランザクションです。
読み取りは引き続き機能しますが (InnoDB MVCC スナップショット)、オーダー サービスの書き込み (新しい注文、ステータスの更新) はブロックされ、PXC CPU/IO がフラットなままの状態でロック待機タイムアウト (50 秒) を超えて失敗します。伝えられるのは、行ロックの待機 ( information_schema.innodb_trx / perf

ormance_schema.data_lock_waits )、リソースの飽和ではなく、reads-fine/writes-blocked により Postgres のテーブルレベルのロックと区別されます。
mysql-analytics-クエリ
同じ分析レポート アクターが、HAProxy 経由で本番環境の MySQL 注文データベースに対して大規模な結合/集計クエリ (ファイルソート、一時テーブル) を実行します。
注文サービス/チェックアウトの遅延とエラーの増加。 PXC CPU/IO の飽和。ワークロードに起因する低速クエリ ログ内の大量のステートメント。
デプロイ
シナリオ
仕組み
RCA ツールが見つけるべきもの
注文サービス-GC-回帰
本物の悪いデプロイ: order-service は 1.1.0 をロールアウトします。これは、読み取られたすべての注文を非効率なキャッシュにディープコピーする実際のコード回帰です。 GC 圧力が上昇します。 revert は、正常なイメージにロールバックします。
p99はロールアウト後に上昇するが、p50は横ばいのままである。 JVM 割り当て率と GC 時間の上昇。ヒープの鋸歯状の傾向は限界に向かっています。開始は展開イベントと正確に相関します。
オーダーサービスメモリリーク
本物の悪いデプロイ: order-service は 1.4.0 をロールアウトします。これは、読み取りごとに小さな「監査証跡」オブジェクトのバッチをプルーニングされることのないレジストリに追加する本当の回帰です。何百万もの小さな物体のゆっくりとした漏れ。 revert は、正常なイメージにロールバックします。
p95/p99 は徐々に上昇します (クラッシュやステップの変化はありません)。旧世代/ライブセットは増加傾向にあります。ライブ セットが大きくなるにつれて、GC 時間と混合コレクションの頻度が増加します。開始はロールアウトと一致します。高速 OOM クラッシュ リークとは異なります。
製品カタログ-GC-圧力
正真正銘の悪いデプロイ: product-catalog は 1.1.0 をロールアウトします。そのサーバー側の「製品カード」は、返されるすべての製品を、読み取りごとに大きな短期間のバッファーに再エンコードします。何も保持されません (リークなし) - 純粋な割り当てのチャーン。 revert はロールバックします。
Go GC CPU 割合とサイクル周波数のスパイク。使用中のヒープが制限されている間、割り当て率が跳ね上がります (OO なし)

母）; product-catalog CPU が飽和/スロットルになり、レイテンシーが増加し、 api-gateway に伝播します。 Postgres は健全な状態を保ちます。
レビューサービスイベントループ
本当に悪いデプロイ: review-service は 1.1.0 をロールアウトし、リクエスト パスに同期「コンテンツ セーフティ」CPU ループを追加して、読み取りごとに数十ミリ秒のシングルスレッド Node.js イベント ループをブロックします。 Revert はロールバックします。
フラット RPS での p95/p99 バルーン。イベントループのラグスパイクと 1 つの CPU コアペグ。レイテンシは DB 時間ではなく、同時実行 (リクエストのシリアル化) に応じて増加します。 MongoDB は健全な状態を保ちます。ボトルネックはデータベースではなく、インプロセス CPU です。
推奨事項-メモリリーク
本当に悪いデプロイ: Recommendation-service が 1.1.0 をロールアウトし、
[切り捨てられた]
シナリオ
仕組み
RCA ツールが見つけるべきもの
トラフィックの急増
ロードジェネレーターのデプロイメントは 5 つのレプリカ、つまりスタック全体にわたる実際の追加トラフィックにスケールされます。
どこでも均一の RPS が増加します。飽和 (遅延/エラー) は最も弱いコンポーネントでのみ発生し、原因と結果の推論をテストします。
CPU-ノイジーネイバー
バッチ ビデオ トランスコーダ ワークロードは、order-service を実行しているノード上に同じ場所に配置され (ポッド アフィニティ)、すべてのコアを消費します。
ノードの CPU が飽和します (~100%)。 Burstable オーダー サービスは、通常の CPU をはるかに下回って不足しています。その依存関係 (MySQL、Kafka) は健全なままです。原因は、被害者ではなく、共同テナントからのノードローカル CPU 競合です。
ネットワーク
シナリオ
仕組み
RCA ツールが見つけるべきもの
DNS-低解像度
Chaos Mesh は、クラスター DNS サービスへのアプリ層のパケットを遅延させます (~500 ミリ秒) (捏造された応答ではなく、DNS パス上の実際のネットワーク状態) ため、すべての名前検索が遅くなります。
サービスは、すべての発信呼び出しで断続的な p95/p99 スパイクを示します (新しい接続ごとに低速のルックアップがフロントロードされます)。一方、すべての依存関係と CoreDNS 自体は正常な状態を維持します (フラット CPU)。教えてください

■ DNS クエリの遅延。ワンホップではありません。古典的な「常に DNS」です。
ネットワーク遅延製品カタログ
Chaos Mesh は、製品カタログに最大 200 ミリ秒の出力遅延を発生させます (デッドマン spec.duration を伴う NetworkChaos 障害)。
product-catalog 自体の CPU/DB が正常な状態を維持している間、カタログベースのエンドポイントの api-gateway レイテンシーは最大 1 秒に跳ね上がります。遅延はサービスや PostgreSQL ではなく、ネットワーク パス上で発生します。
信頼性
潜在的なゆっくりとした燃焼のリスク - RCA ではなく検出 (上記の注を参照)。多くの場合、発症時には急性症状がありません。 「発見すべき」列は、ツールが表面化すべき早期警告信号です。
その他のシナリオ (不適切な移行、接続プールのリーク、Kafka コンシューマの遅延、
キャッシュ削除の圧力など）がロードマップに載っています。それぞれが次に従います
同じ実際のメカニズム、耐久性復帰ルール。
エッジ: 実線 = HTTP、点線 = gRPC、太線 = Kafka イベント。
フローチャート LR
LG([load-generator]):::gen --> GW[api-gateway]:::gw
GW --> PC[製品カタログ]
GW --> CART[カートサービス]
GW --> ORD[オーダーサービス]
GW --> REV[レビューサービス]
GW --> INV[在庫サービス]
GW～。 gRPC .-> REC[推奨サービス]
パソコン-。 gRPC .-> REC
カート -- チェックアウト --> ORD
ORD -- 同期 --> PAY[支払いサービス]
PC --> PGP[(製品)]:::db
INV --> PGI[(在庫)]:::db
CART --> VK[(ヴァルキークラスター)]:::db
または

[切り捨てられた]

## Original Extract

A Kubernetes lab that reproduces real production incidents on a live, instrumented microservice stack — for testing root-cause-analysis tools and agents. - coroot/rca-lab

GitHub - coroot/rca-lab: A Kubernetes lab that reproduces real production incidents on a live, instrumented microservice stack — for testing root-cause-analysis tools and agents. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
coroot
/
rca-lab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .github/ workflows .github/ workflows deploy deploy docs docs operator operator scenarios scenarios scripts scripts services services test/ e2e test/ e2e variants variants .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md versions.yaml versions.yaml View all files Repository files navigation
A realistic, reproducible failure lab for evaluating root-cause-analysis (RCA)
tooling — human or AI — on a live Kubernetes cluster.
Most RCA benchmarks replay canned telemetry from toy environments with
synthetic faults toggled by feature flags. rca-lab takes the opposite approach:
A real polyglot microservice stack (Python, Go, Java, Node.js, Rust, PHP)
behind an API gateway, with continuous generated load.
Real databases under production-grade operators : PostgreSQL, MySQL and
MongoDB via Percona operators, a Valkey Cluster via the valkey-operator,
Kafka via Strimzi — with
seeded data volumes.
Real failure mechanisms only. No chaos flags inside the apps. A GC
pressure incident is a genuine allocation regression shipped as a new image
version and rolled back later; a database incident is an analytics workload
running heavy queries against the production database; a traffic spike is
actually more traffic.
Durable revert. Every failure scenario is a FailureScenario custom
resource driven by an operator that restores the normal state when the
scenario ends, is disabled, or is deleted — even across operator restarts.
Rich telemetry, bring your own backend. Every service is instrumented
with OpenTelemetry SDKs: traces, SDK-emitted metrics (JVM/runtime/HTTP), and
logs (to stdout and OTLP, trace-correlated). Everything flows to a bundled
otel-collector that discards data by default — point it at any OTLP
backend with one variable.
Requirements: kubectl + helm pointed at a cluster (any distribution;
a default StorageClass, ~8 CPU / 16 GiB across nodes for the full-size lab).
git clone https://github.com/coroot/rca-lab && cd rca-lab
make deploy # everything: operators → databases → Kafka → apps → seed
Single-node cluster (kind/k3d/minikube):
make deploy SINGLE_NODE=1
Send telemetry somewhere (e.g. Coroot, or any OTLP endpoint):
make deploy OTLP_ENDPOINT=my-backend:4317
Other variables: STORAGE_CLASS=<name> , SEED_SIZE_GB=<n> (0 skips seeding),
OTLP_HEADERS=k=v , YES=1 (no confirmation prompt). Re-running make deploy
converges idempotently — it is also how you change any of these settings.
make clean # KEEP_DATA=1 keeps the database volumes
Triggering failures
Scenarios are Kubernetes custom resources:
kubectl get failurescenarios
kubectl patch failurescenario pg-analytics-queries --type=merge -p ' {"spec":{"enabled":true}} '
or use the web UI:
kubectl port-forward svc/rca-lab-operator 8080
The UI lists every scenario grouped by category, with severity and live state, and
starts or stops each one with a click.
Each scenario documents its mechanism and the telemetry symptoms an RCA tool
should be able to observe. Scenarios can also run on a cron schedule with a
fixed duration — see scenarios/ .
Never install rca-lab on a shared or production cluster. The scenario
operator deliberately has the power to degrade workloads in its namespace.
Every scenario uses a genuine real-world mechanism — never a synthetic fault
flag inside the app — and reverts durably. Each carries an expectedSymptoms
list that doubles as documentation and a grading rubric for RCA tools.
The reliability category is a different kind of test. The other scenarios
are acute incidents (latency, errors, saturation) that exercise RCA — given
a symptom, find the cause. Reliability scenarios are latent, slow-burn risks
(bloat, stale stats, blocked vacuum, replication lag, checkpoint pressure) that
often produce no user-facing symptom at onset ; they exercise proactive
detection — whether a tool flags a developing risk before it becomes an
outage. Their expectedSymptoms are early-warning indicators, not incident
symptoms.
Scenario
Mechanism
What an RCA tool should find
pg-analytics-queries
An analytics-reporting workload runs heavy multi-join/aggregation queries (full scans of the ~10 GB products table) against the production PostgreSQL, through the same pgBouncer pool as the apps.
Elevated product-catalog / inventory-service latency; PostgreSQL CPU/IO saturation; new full-scan query fingerprints in pg_stat_statements attributable to the analytics-reporting workload.
pg-exclusive-lock
A stalled schema-migration transaction takes a real ACCESS EXCLUSIVE lock on the products table ( LOCK TABLE ) and then hangs holding it — the "a migration grabbed the lock and never let go" incident.
product-catalog queries on products block on the lock; its connection pool fills and the service goes unavailable, so api-gateway product endpoints error — yet PostgreSQL CPU/IO stay flat because nothing is executing. The tell is lock waits ( pg_locks / pg_blocking_pids ), not resource saturation.
mysql-lock-contention
A stalled transaction holds InnoDB row locks on the hot end of the orders table ( SELECT … FOR UPDATE , including the gap above the max id) and then hangs — a transaction that grabbed locks and stalled.
Reads keep working (InnoDB MVCC snapshots), but order-service writes (new orders, status updates) block and fail with Lock wait timeout exceeded (50s) while PXC CPU/IO stay flat. The tell is row-lock waits ( information_schema.innodb_trx / performance_schema.data_lock_waits ), not resource saturation — and reads-fine/writes-blocked distinguishes it from Postgres's table-level lock.
mysql-analytics-queries
The same analytics-reporting actor runs large join/aggregation queries (filesort, temp tables) against the production MySQL orders database via HAProxy.
Elevated order-service /checkout latency and errors; PXC CPU/IO saturation; heavy statements in the slow query log attributable to the workload.
Deploy
Scenario
Mechanism
What an RCA tool should find
order-service-gc-regression
A genuine bad deploy: order-service rolls out 1.1.0 , a real code regression that deep-copies every order read into an ineffective cache. GC pressure builds; revert rolls back to the known-good image.
p99 rises after the rollout while p50 stays flat; JVM allocation rate and GC time climb; heap sawtooth trends toward the limit; onset correlates exactly with the deployment event.
order-service-memory-leak
A genuine bad deploy: order-service rolls out 1.4.0 , a real regression that appends a batch of small "audit trail" objects per read into a registry that is never pruned. Slow leak of millions of tiny objects; revert rolls back to the known-good image.
p95/p99 creep up gradually (no crash, no step change); old-gen/live-set trends up; GC time and mixed-collection frequency rise as the live set grows; onset matches the rollout. Distinct from the fast OOM-crash leaks.
product-catalog-gc-pressure
A genuine bad deploy: product-catalog rolls out 1.1.0 , whose server-side "product cards" re-encode every returned product into large short-lived buffers on each read. Nothing retained (no leak) — pure allocation churn; revert rolls back.
Go GC CPU fraction and cycle frequency spike; allocation rate jumps while heap in-use stays bounded (no OOM); product-catalog CPU saturates/throttles and latency rises, propagating to api-gateway ; Postgres stays healthy.
review-service-event-loop
A genuine bad deploy: review-service rolls out 1.1.0 , adding a synchronous "content safety" CPU loop on the request path that blocks the single-threaded Node.js event loop for tens of ms per read. Revert rolls back.
p95/p99 balloon at flat RPS; event-loop lag spikes and one CPU core pegs; latency grows with concurrency (requests serialize), not with DB time; MongoDB stays healthy — the bottleneck is in-process CPU, not the database.
recommendation-memory-leak
A genuine bad deploy: recommendation-service rolls out 1.1.0 , a re
[truncated]
Scenario
Mechanism
What an RCA tool should find
traffic-spike
The load-generator Deployment is scaled to 5 replicas — real extra traffic across the whole stack.
Uniform RPS increase everywhere; saturation (latency/errors) appears only at the weakest component, testing cause-vs-consequence reasoning.
cpu-noisy-neighbor
A batch video-transcoder workload is co-located (pod affinity) onto the nodes running order-service and burns all their cores.
Node CPU saturates (~100%); the Burstable order-service is starved far below its normal CPU; its dependencies (MySQL, Kafka) stay healthy — the cause is node-local CPU contention from a co-tenant, not the victim.
Network
Scenario
Mechanism
What an RCA tool should find
dns-slow-resolution
Chaos Mesh delays the app tier's packets to the cluster DNS service (~500 ms) — a real network condition on the DNS path, not fabricated answers — so every name lookup is slow.
Services show intermittent p95/p99 spikes on all outbound calls (each new connection front-loads a slow lookup), while every dependency and CoreDNS itself stay healthy (flat CPU). The tell is DNS query latency, not any one hop — the classic "it's always DNS."
network-delay-product-catalog
Chaos Mesh injects ~200 ms of egress latency on product-catalog (a NetworkChaos fault with a dead-man spec.duration ).
api-gateway latency for catalog-backed endpoints jumps to ~1 s while product-catalog 's own CPU/DB stay healthy; the delay is on the network path, not in the service or PostgreSQL.
Reliability
Latent, slow-burn risks — detection, not RCA (see the note above). Each often has no acute symptom at onset; the "should find" column is the early-warning signal a tool should surface.
More scenarios (bad migrations, connection-pool leaks, Kafka consumer lag,
cache eviction pressure, and others) are on the roadmap; each will follow the
same real-mechanism, durable-revert rule.
Edges: solid = HTTP, dotted = gRPC, thick = Kafka event.
flowchart LR
LG([load-generator]):::gen --> GW[api-gateway]:::gw
GW --> PC[product-catalog]
GW --> CART[cart-service]
GW --> ORD[order-service]
GW --> REV[review-service]
GW --> INV[inventory-service]
GW -. gRPC .-> REC[recommendation-service]
PC -. gRPC .-> REC
CART -- checkout --> ORD
ORD -- sync --> PAY[payment-service]
PC --> PGP[(products)]:::db
INV --> PGI[(inventory)]:::db
CART --> VK[(Valkey Cluster)]:::db
OR

[truncated]

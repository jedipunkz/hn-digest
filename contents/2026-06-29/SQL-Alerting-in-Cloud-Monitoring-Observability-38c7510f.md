---
source: "https://cloud.google.com/blog/products/management-tools/alert-with-sql-in-cloud-monitoring-observability-analytics/"
hn_url: "https://news.ycombinator.com/item?id=48718350"
title: "SQL Alerting in Cloud Monitoring Observability"
article_title: "Alert with SQL in Cloud Monitoring Observability Analytics | Google Cloud Blog"
author: "mjs06"
captured_at: "2026-06-29T13:24:56Z"
capture_tool: "hn-digest"
hn_id: 48718350
score: 1
comments: 0
posted_at: "2026-06-29T12:24:14Z"
tags:
  - hacker-news
  - translated
---

# SQL Alerting in Cloud Monitoring Observability

- HN: [48718350](https://news.ycombinator.com/item?id=48718350)
- Source: [cloud.google.com](https://cloud.google.com/blog/products/management-tools/alert-with-sql-in-cloud-monitoring-observability-analytics/)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T12:24:14Z

## Translation

タイトル: Cloud Monitoring の可観測性における SQL アラート
記事のタイトル: Cloud Monitoring の可観測性分析における SQL を使用したアラート | Google クラウド ブログ
説明: Cloud Monitoring Observability Analytics を使用すると、ログとトレースの分析 SQL クエリからアラートを作成 (およびアラートの受信) できます。

記事本文:
Cloud Monitoring の可観測性分析における SQL を使用したアラート | Google Cloud ブログ コンテンツに移動 クラウド ブログ 営業担当者へのお問い合わせ 無料で始めましょう クラウド ブログ ソリューションとテクノロジー AI と機械学習
管理ツール クエリからアクションまで: Cloud Monitoring Observability Analytics での SQL アラートの導入
職場における AI への玄関口
従来のアラート システムでは、多くの場合、妥協を強いられます。単純でノイズの多いログ イベントに対して即座にアラートを送信することも、ユーザー セッションや IP アドレスなどの固有の答えが多数あるデータに直面した場合に失敗する、厳格で事前設定されたメトリクスを監視することもできます。しかし、最も重大なシステムの問題 (特定の顧客のエラー率の 20% の急増や、データベースのタイムアウトと相関するレイテンシの異常など) は、これらの信号間の集計と関係の中に隠されています。
最近、Observability Analytics (旧名 Log Analytics) で SQL を使用してログとトレースをクエリできるようになったと発表しました。しかし、話はさらに良くなります。 SQL を使用して、Observability Analytics でアラートを作成することもできます。 SQL をアラート エンジンに直接導入することで、ログやトレースに対して複雑な分析クエリを作成し、それらをアラートに変換できます。エラーのパーセンテージの計算、高カーディナリティのディメンションの分析、またはログとトレースの結合のいずれが必要な場合でも、SQL アラートは、基本的なしきい値監視から、従来のアラート システムの機能を超えた詳細なコンテキスト検出に移行するのに役立ちます。 SQL アラートは現在プレビュー段階にあります。
Observability Analytics の SQL アラートは、Cloud Monitoring の一部として利用できます。アラート ポリシーは、定義したスケジュール (たとえば、10 分ごと) で SQL クエリを実行します。クエリに「ルックバック ウィンドウ」が自動的に適用されるため、クエリ以降に受信したログ エントリまたはトレース スパンのみが分析されます。

最後に走ったとき。
クエリの結果が設定した条件を満たす場合、Cloud Monitoring はインシデントを作成し、メール、Slack、PagerDuty などの選択したチャネルに通知を送信します。
SQL ベースのアラートでは BigQuery を使用してテレメトリ データを処理するため、クエリの実行は標準のオンデマンド料金または BigQuery 予約に基づいて BigQuery を通じて請求されることに注意してください。
2 種類のアラート条件から選択できます。
行数のしきい値: これは最も単純なオプションです。アラートは、クエリが返す行数が、設定したしきい値より大きい、等しい、または小さい場合に起動されます。これは、「10 人を超えるユーザーがログインに失敗した場合に警告する」シナリオに最適です。
ブール値: これは最も強力なオプションです。定義した特定の列の値が true である行がクエリから返された場合、アラートが発生します。これにより、パーセンテージの計算などの複雑なロジックを SQL クエリ内で直接構築できます。
例 1: 支払いゲートウェイの失敗に関するアラート (行数)
シナリオ: あなたは電子商取引オペレーターであり、時折発生する通常のカードの拒否 (PIN の間違いなど) を無視しながら、支払いゲートウェイでシステム的な停止が発生した場合にすぐにアラートを受け取りたいと考えているとします。
これを行うには、ゲートウェイのタイムアウトを示すログ エントリをフィルターするクエリを作成し、行数のしきい値を使用して、これらのエラーの量が急増した場合にのみアラートをトリガーできます。
読み込み中... 選択
JSON_VALUE(json_payload.transaction_id) AS トランザクション ID、
JSON_VALUE(json_payload.error_code) AS error_code
から
`my-project-id.my-dataset.my-log-view`
どこで
JSON_VALUE(json_payload.status) = '失敗'
-- WRONG_PIN などのユーザー入力エラーではなく、ゲートウェイのシステム的な問題をフィルターします。
AND JSON_VALUE(json_payload.failure_reason) = 'GATEWAY_TIMEOUT' アラート構成:
条件タイプ

: 行数のしきい値
トリガー条件: 行数が ( > ) 10 を超えると発生します。
評価ウィンドウ/ルックバック: 5 分 (定義したスケジュールに従って過去 5 分間のデータをチェックします)
例 2: エージェントの遅延 (トレース) に関するアラート
シナリオ: あなたは AI プラットフォーム エンジニアで、マルチステップ AI エージェントが許容可能な時間制限内で応答していることを確認したいと考えています。オーケストレーター サービスの 99 パーセンタイル (p99) 待機時間を監視し、パフォーマンスが低下した場合にアラートを受け取りたいと考えています。
これを行うには、すべてのサービスの p99 レイテンシーを計算し、エージェント オーケストレーターが 5 秒 (5000 ミリ秒) を超えた場合に true を返す、トレース データに対して SQL クエリを作成できます。
読み込み中... WITH latency_data AS (
選択
APPROX_QUANTILES(duration_nano, 100)[OFFSET(99)] / 1000000 AS p99_ms
から
`my-project-id.us._Trace.Spans._AllSpans`
どこで
-- エージェント オーケストレーターによって生成された行を調べる
JSON_VALUE(resource.attributes, '$."service.name"') = 'エージェント オーケストレーター'
グループ化
サービス名
）
選択
「エージェント オーケストレーター」 AS service_name、
p99_ms、
-- ブール論理: p99 が 5000ms を超えた場合に警告
(p99_ms > 5000) AS has_latency_spike
から
latency_data アラート構成:
ターゲット列: has_latency_spike
トリガー条件: この列が true と評価される行をクエリが返すと発生します。
評価ウィンドウ/ルックバック: 10 分 (または希望のスケジュール間隔)
SQL ベースのアラートを作成する前に、いくつかのことを設定する必要があります。
ログの場合: Observability Analytics を使用するようにログ バケットをアップグレードします (まだ更新されていない場合)。
トレースの場合: Cloud Trace を収集し、プロジェクトに保存する必要があります。
リンクされた BigQuery データセット: テレメトリ ソース (ログ バケットまたはトレース データセットのいずれか) 用にリンクされた BigQuery データセットを作成します。 SQL ベースのアラートは、この BigQu を通じてデータをクエリします。

エリリンク。
SQL ベースのアラート ポリシーの作成に必要な IAM ロールを付与します: モニタリング アラート ポリシー エディターおよびログ SqlAlert ライター (ログ アラートとトレース アラートの両方に適用されます)。
通知チャネル: アラートを受信する通知チャネル (電子メールや Slack など) を構成します。
SQL ベースのアラート ポリシーの作成は簡単です。
Google Cloud コンソールでオブザーバビリティ分析に移動します。
SQL クエリを作成して検証します。
UI で [Run on BigQuery] クエリ エンジンを選択します。
結果ツールバーから「アラートの作成」ボタンをクリックします。
条件 (行数またはブール値) と評価スケジュールを定義します。
通知チャネルを追加し、アラートにわかりやすい名前を付けて、 [保存] をクリックします。
Infrastructure as Code (IaC) パイプラインの場合は、API および Terraform を介してアラートを構成することもできます。
より強力で洞察力に富んだアラートを構築する準備はできていますか?コンソールで Observability Analytics ページを開いて、今すぐ最初の SQL クエリを作成してみてください。詳細と高度な例については、公式ドキュメントを参照してください。
Log Analytics は Observability Analytics になりました: SQL を使用してログとトレースをクエリします
Gemini を搭載したフリート インテリジェンスを備えた最新のデータベース センターをご紹介します
キラン・シェノイ著 • 5 分で読めます
Gemini Cloud Assist: 要求する前から、お客様に合わせて機能するプロアクティブなクラウド運用
Michael Bachman著 • 5分で読めます
統合メンテナンス: Google Cloud 全体のメンテナンスを管理する新しい統合方法
エロル・ヴァレリウ・キオアスカ著 • 2 分で読めます
言語 英語 ドイツ語 フランス語 한국어 日本語

## Original Extract

Cloud Monitoring Observability Analytics lets you create alerts from (and get alerted about) analytical SQL queries of logs and traces.

Alert with SQL in Cloud Monitoring Observability Analytics | Google Cloud Blog Jump to Content Cloud Blog Contact sales Get started for free Cloud Blog Solutions & technology AI & Machine Learning
Management Tools From query to action: Introducing SQL alerting in Cloud Monitoring Observability Analytics
The front door to AI in the workplace
Traditional alerting systems often force a compromise: you can either alert immediately on simple, noisy log events, or monitor rigid, pre-configured metrics that fail when faced with data with many unique answers like user sessions or IP addresses. But the most critical system issues — like a 20% spike in error rates for a specific customer or a latency anomaly correlated with database timeouts — are hidden in the aggregates and relationships between these signals.
Recently, we announced that you can now use SQL to query logs and traces in Observability Analytics (formerly Log Analytics). But the story gets better. You can also use SQL to create alerts in Observability Analytics. By bringing SQL directly to your alerting engine, you can write complex analytical queries over logs and traces and turn them into alerts. Whether you need to calculate error percentages, analyze high-cardinality dimensions, or JOIN logs and traces, SQL alerting helps you go from basic threshold monitoring to deep, contextual detection that goes beyond the capabilities of traditional alerting systems. SQL alerting is now in preview.
SQL alerting in Observability Analytics is available as part of Cloud Monitoring . An alerting policy runs your SQL query on a schedule you define (for example, every 10 minutes). It automatically applies a "lookback window" to your query, so it only analyzes the log entries or trace spans it received since the last time it ran.
If the results of your query meet the condition you set, Cloud Monitoring creates an incident and sends a notification to your chosen channels, like email, Slack, or PagerDuty.
Please note that because SQL-based alerting uses BigQuery to process telemetry data, query executions are billed through BigQuery under your standard on-demand pricing or BigQuery reservations.
You can choose between two types of alert conditions.
Row count threshold: This is the simplest option. The alert fires if your query returns a number of rows that is greater than, equal to, or less than a threshold you set. This is perfect for "alert me if more than 10 users have failed logins" scenarios.
Boolean: This is the most powerful option. The alert fires if your query returns any row where a specific column you define has a value of true. This lets you build complex logic, like calculating percentages, directly in your SQL query.
Example 1: Alerting on payment gateway failures (row count)
Scenario: Imagine that you’re an e-commerce operator, and you want to be alerted immediately if your payment gateway is experiencing systemic outages, while ignoring occasional, normal card declines (like an incorrect PIN).
To do this, you can write a query to filter for log entries indicating gateway timeouts, and use a row count threshold to trigger the alert only if the volume of these errors spikes.
Loading... SELECT
JSON_VALUE(json_payload.transaction_id) AS transaction_id,
JSON_VALUE(json_payload.error_code) AS error_code
FROM
`my-project-id.my-dataset.my-log-view`
WHERE
JSON_VALUE(json_payload.status) = 'FAILED'
-- Filter for systemic gateway issues, not user-input errors like WRONG_PIN
AND JSON_VALUE(json_payload.failure_reason) = 'GATEWAY_TIMEOUT' Alert configuration:
Condition type: Row count threshold
Trigger condition: Fired when row counts greater than ( > ) 10
Evaluation window / lookback: 5 minutes (checks the last 5 minutes of data on your defined schedule)
Example 2: Alerting on agent latency (traces)
Scenario: You’re an AI platform engineer, and you want to ensure your multi-step AI agents are responding within acceptable time limits. You want to monitor the 99th percentile (p99) latency of the orchestrator service and get alerted if performance degrades.
To do this, you can write a SQL query against your trace data that calculates the p99 latency for all services and returns true if your agent-orchestrator exceeds 5 seconds (5000 milliseconds).
Loading... WITH latency_data AS (
SELECT
APPROX_QUANTILES(duration_nano, 100)[OFFSET(99)] / 1000000 AS p99_ms
FROM
`my-project-id.us._Trace.Spans._AllSpans`
WHERE
-- Examine rows produced by the agent-orchestrator
JSON_VALUE(resource.attributes, '$."service.name"') = 'agent-orchestrator'
GROUP BY
service_name
)
SELECT
"agent-orchestrator" AS service_name,
p99_ms,
-- Boolean logic: Alert if p99 exceeds 5000ms
(p99_ms > 5000) AS has_latency_spike
FROM
latency_data Alert configuration:
Target column: has_latency_spike
Trigger condition: Fired when the query returns any row where this column evaluates to true.
Evaluation window / lookback: 10 minutes (or your preferred scheduling interval)
Before you can create a SQL-based alert, you need to set up a few things:
For logs : Upgrade your log bucket to use Observability Analytics (if not already updated).
For traces : Cloud Trace must be collected and stored in your project.
Linked BigQuery dataset: Create a linked BigQuery dataset for the telemetry source (either the log bucket or the trace dataset ). SQL-based alerts query the data through this BigQuery link.
Grant the IAM roles necessary to create an SQL-based alert policy: Monitoring AlertPolicy Editor and Logging SqlAlert Writer (applies to both log and trace alerts).
Notification channels: Configure the notification channels (like email or Slack) where you want to receive alerts.
Creating a sql-based alert policy is straightforward:
Navigate to Observability Analytics in the Google Cloud console.
Compose and validate your SQL query.
Select the Run on BigQuery query engine in the UI.
Click the Create alert button from the results toolbar.
Define your condition (row count or boolean) and your evaluation schedule.
Add your notification channels, give your alert a clear name, and click Save .
For Infrastructure as Code (IaC) pipelines, you can also configure alerts via the API and Terraform .
Ready to build more powerful, insightful alerts? Open the Observability Analytics page in the console and try writing your first SQL query today. You can find more details and advanced examples in the official documentation .
Log Analytics is now Observability Analytics: Query logs and traces with SQL
Meet the latest Database Center, now with Gemini-powered fleet intelligence
By Kiran Shenoy • 5-minute read
Gemini Cloud Assist: Proactive cloud operations that work for you, even before you ask
By Michael Bachman • 5-minute read
Unified Maintenance: A new, unified way to manage maintenance across Google Cloud
By Erol-Valeriu Chioasca • 2-minute read
Language ‪English‬ ‪Deutsch‬ ‪Français‬ ‪한국어‬ ‪日本語‬

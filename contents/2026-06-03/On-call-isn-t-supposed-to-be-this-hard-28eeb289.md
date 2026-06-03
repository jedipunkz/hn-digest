---
source: "https://fixbugs.ai/blog/good-alerts-bad-alerts"
hn_url: "https://news.ycombinator.com/item?id=48369654"
title: "On-call isn't supposed to be this hard"
article_title: "Oncall isn't supposed to be this hard - FixBugs Blog | FixBugs | From Alert to Validated Fix"
author: "kirtivr"
captured_at: "2026-06-03T00:47:19Z"
capture_tool: "hn-digest"
hn_id: 48369654
score: 4
comments: 1
posted_at: "2026-06-02T12:58:57Z"
tags:
  - hacker-news
  - translated
---

# On-call isn't supposed to be this hard

- HN: [48369654](https://news.ycombinator.com/item?id=48369654)
- Source: [fixbugs.ai](https://fixbugs.ai/blog/good-alerts-bad-alerts)
- Score: 4
- Comments: 1
- Posted: 2026-06-02T12:58:57Z

## Translation

タイトル: オンコールはこんなに大変なはずがない
記事のタイトル: オンコールはこれほど難しいものではない - FixBugs Blog |バグ修正 |アラートから検証済みの修正まで
説明: 適切に構成されたアラートは、症状をトレース、ログ、デプロイ、および疑わしいコミットに結び付けます。

記事本文:
オンコールはこれほど難しいものではない - FixBugs Blog |バグ修正 |アラートから検証済み修正まで メイン コンテンツにスキップ FixBugs 製品価格設定 ブログ アーキテクチャ ドキュメント サポート ホームについて
オンコールはそれほど難しくないはずです
オンコールはそれほど難しくないはずです
適切に構成されたアラートは、症状をトレース、ログ、デプロイ、および疑わしいコミットに結び付けます。
悪い Prometheus アラートは、オンコール エンジニアに何か問題があることを伝えますが、良好なアラートは、症状をトレース、ログ、デプロイ、および疑わしいコミットに結び付けます。
オンコールになってアラートの嵐が現れるまでは、その区別は些細なことのように思えます。
アラートの 1 つを開くと、次の内容が表示されます。
[クリティカル] CheckoutHighErrorRate - prod-eu-west-1 の 7.3% 5xx
警報は間違っていません。チェックアウトでエラーが発生しました。しかし、その理由や、どのホスト/コンテナ/VM から調査を開始するのかさえも示されていません。
SRE / 開発者は、今やるべきすべての作業を行うことになります。
何をしようとしているのかわかっている場合は、まずアラートの定義を確認します。
Prometheus の基本的なセットアップは通常次のようになります。
- アラート: CheckoutHighErrorRate
式: |
sum(rate(http_requests_total{service="チェックアウト",status=~"5.."}[5m]))
/
sum(rate(http_requests_total{service="checkout"}[5m])) > 0.05
用: 10メートル
ラベル:
重大度: クリティカル
チーム: 支払い
サービス: チェックアウト
注釈:
概要: 「チェックアウト 5xx 率は 5% を超えています」
runbook_url: "https://runbooks.corp/payments/checkout-5xx"
ダッシュボードの URL: "https://grafana.corp/d/checkout"
これにより、いくつかの重要な情報が得られます。
アラートは 5 分間の HTTP エラー (タイムアウトを含む) を集計し、しきい値と比較します。
アラートは支払いチームが所有します。
そこから始めることができるプレイブックがあります。
しかし、まだ祝う理由はありません。
時間枠をアラート時間の 5 分以内に調整します。
Grafana を開き、ダッシュボードに

追加の情報。
Loki を開き、`{service_name="checkout"} |~ "(?i)error" のようなクエリを作成します。
Tempo を開き、時間でトレースをフィルターします。どの痕跡がインシデントを表しているかを推測してください。
CD パイプラインを開き、アラートの直前にデプロイを検索します。
ある時点で、いくつかの可能な仮説が現れます。
checkout-api@v2.4.1 で新たに導入された大きな機能は怪しいようです。
5 台のホストのうち 3 台で CPU 使用率が高く、5xx エラーが報告されました。
調査したすべてのホストで疑わしい I/O エラー。
DB トランザクションが遅い。
最終的に、開発者は、何をしているのかを正確に把握していれば、約 1 時間で 4 つのツールにわたるコンテキストを再構築することに成功しました。
その間、調査すべき他の新しいアラートが存在する可能性があります。
優れたアラートにより、どこから探し始めるべきかがわかります
同じスタックでも動作が大きく異なる場合があります。
別のベンダーではありません。それほど高価な警告製品ではありません。
同じスタックが、コンテキストをバブルアップするために正しく配線されています。
Prometheus/Grafana/Tempo/Loki スタックの場合は次のようになります。
-> OpenTelemetry SDK を使用する Prometheus エクスポーター。
-> ヒストグラムはトレース範囲と相関しています。
-> Grafana エグザンプラが有効になりました。
-> ログへのトレースを有効にしてテンポを設定します。
-> マーカー / service.version / commit をデプロイし、各アラートにメタデータとして追加された SHA を追加します。
アラートは引き続きメトリックから始まります。それはすべきです。メトリクスは症状を検出する方法です。
ただし、メトリクスには特定のリクエストへのブレッドクラムが含まれるようになりました。
Prometheus アラートには、当然のことながら、trace_id は含まれません。ヒストグラム バケットは集約であり、単一のリクエストではありません。
模範となる人たちがそれを変えます。サンプリングされた測定により、アクティブな trace_id をバケットに添付できます。 Grafana は、それをグラフ上でクリック可能なダイヤモンドとしてレンダリングできます。それをクリックすると、Tempo によって代表的なトレースが開きます。
良いバージョンでは、選択されたスパンは次のようになります。
サービス: db-primary
操作: SELECT 注文 WHERE user_id

=$1
持続時間: 1210ms
db.rows_affected: 1110482
feature_flag.new_checkout: true
サービスのバージョン: 2.4.1
分散トレースでは、データベース クエリが遅いことがわかります。
次に、Tempo のトレースからログへのリンクにより、Loki が開き、正確なトレースが表示されます。
ログ行はタイム ウィンドウ クエリに埋もれなくなりました。
遅いクエリ: オーダーでのシーケンス スキャン (110 万行)、インデックスは使用されません
トレース_id=4bf92f3577b34da6a3ce929d0e0e4736
スパン_id=00f067aa0ba902b7
サービス.バージョン=2.4.1
コミット=7a3f9c2
これで、仮説はもはや曖昧ではなくなりました。
checkout-api@v2.4.1 は、新しい注文履歴クエリ パスを追加しました。
user-id 列をインデックスとして追加する必要があります。
不正なパスは、feature_flag.new_checkout=true によってゲートされます。
フラグを無効にするか、7a3f9c2 をロールバックします。
オンコール体験を楽しくするのは構成です
Prometheus + Grafana、Datadog、New Relic のいずれを使用しているかに関係なく、これらはいずれも自動ではなく、自動的に行われることもありません。
適切なパスには計画的な配管が必要です。
症状に関するページ: エラー率、遅延、トラフィック、飽和、または SLO バーン。
アラートにteam、service、severity、runbook_url、およびスコープ指定されたdashboard_urlを追加します。
すべてのサービスを通じて W3C トレース コンテキストを伝播します。
構造化ログにtrace_idとspan_idを挿入します。
アラートで使用されるヒストグラムのサンプルを有効にします。
エグザンプラが Tempo を開くように Grafana を構成します。
スパンが Loki を開くように Tempo トレースからログを設定します。
service.version を発行し、アノテーションをデプロイし、CI/CD から SHA をコミットします。
ここで、このような適切に構成されたセットアップを試すことができます。
有用な AI SRE ワークフローは、可観測性スタックが証拠の証跡を保存した後に開始されます。エージェントは根本原因の分析を支援し、修正を提案し、パッチを検証します。しかし、アラートによってトレース、ログ相関、およびデプロイ コンテキストが削除された場合、エージェントは人間と同じ問題、つまり推測することになります。
私たちがどこにいるのかについては、

製品でこれを行う場合は、「FixBugs および AI SRE ツール」を参照してください。
Google SRE: 分散システムの監視
Google SRE ワークブック: SLO に関するアラート
Grafana: トレースとログの相関関係を構成する
Grafana: ビジュアライゼーションに注釈を付ける
OpenTelemetry サービスのセマンティック規則
ブログに戻る FixBugs AI を活用したデバッグにより、チームは信頼性の高いソフトウェアをより迅速に出荷できるようになります。
© 2026 モデューロ AI 。無断転載を禁じます。

## Original Extract

Well-configured alerts connect the symptom to traces, logs, deploys, and the suspect commit.

Oncall isn't supposed to be this hard - FixBugs Blog | FixBugs | From Alert to Validated Fix Skip to main content FixBugs Product Pricing Blog Architecture Docs Support About Home
Oncall isn't supposed to be this hard
Oncall isn't supposed to be this hard
Well-configured alerts connect the symptom to traces, logs, deploys, and the suspect commit.
Bad Prometheus alerts tell an oncall engineer something is wrong, while good alerts connect the symptom to traces, logs, deploys, and the suspect commit.
That distinction sounds small until you're on-call and an alert storm appears.
You open one of the alerts and see:
[CRITICAL] CheckoutHighErrorRate - 7.3% 5xx in prod-eu-west-1
The alert is not wrong. Checkout is error'ing out. But it hasn't told you why or even which host/container/VM to start investigating from.
The SRE / Developer now has all the work to do.
If you know what you're doing, you first check the alert definition.
A basic Prometheus setup usually looks like this:
- alert: CheckoutHighErrorRate
expr: |
sum(rate(http_requests_total{service="checkout",status=~"5.."}[5m]))
/
sum(rate(http_requests_total{service="checkout"}[5m])) > 0.05
for: 10m
labels:
severity: critical
team: payments
service: checkout
annotations:
summary: "Checkout 5xx rate is above 5%"
runbook_url: "https://runbooks.corp/payments/checkout-5xx"
dashboard_url: "https://grafana.corp/d/checkout"
This gives you some important pieces of information:
The alert aggregates HTTP errors (including timeouts) over a 5 minute period and compares it to a threshold.
The alert is owned by the Payments team.
There is a playbook you can start from.
But there is no reason to celebrate just yet.
Adjust the time window to within 5 minutes of the alert time .
Open Grafana and check if the dashboards have any extra information.
Open Loki and write a query like `{service_name="checkout"} |~ "(?i)error"
Open Tempo and filter traces by time. Guess which trace represents the incident.
Open your CD pipeline and search for any deploys just before the alert.
At some point, several possible hypothesis appear.
Big newly introduced feature in the checkout-api@v2.4.1 looks fishy.
High CPU usage on 3 out of 5 hosts that reported 5xx errors.
Suspicious I/O errors on all the investigated hosts.
Slow DB transactions.
Eventually the developer manages to reconstruct context across four tools, in about an hour if they know exactly what they're doing.
Meanwhile, there may be other fresh alerts to investigate.
good alerts tell you where to start looking
The same stack can behave very differently.
Not a different vendor. Not a more expensive alerting product.
The same stack, wired correctly to bubble up context.
Here is what it would look like for the Prometheus/Grafana/Tempo/Loki stack:
-> Prometheus exporter using OpenTelemetry SDK.
-> histograms correlated with trace spans.
-> Grafana exemplars enabled.
-> Tempo setup with trace-to-logs enabled.
-> deploy marker / service.version / commit SHA added as metadata with each alert.
The alert still starts with a metric. It should. Metrics are how you detect the symptom.
But the metric now carries a breadcrumb to a specific request.
Prometheus alerts do not naturally carry a trace_id . A histogram bucket is an aggregate, not a single request.
Exemplars change that. A sampled measurement can attach the active trace_id to the bucket . Grafana can render that as a clickable diamond on its graph. Click it and Tempo opens the representative trace.
In the good version, the selected span says:
service: db-primary
operation: SELECT orders WHERE user_id=$1
duration: 1210ms
db.rows_affected: 1110482
feature_flag.new_checkout: true
service.version: 2.4.1
We see the slow database queries in the distributed trace .
Then Tempo's trace-to-logs link opens Loki for the exact trace.
The log line is not buried in a time-window query anymore:
slow query: seq scan on orders (1.1M rows), index not used
trace_id=4bf92f3577b34da6a3ce929d0e0e4736
span_id=00f067aa0ba902b7
service.version=2.4.1
commit=7a3f9c2
Now the hypothesis is no longer vague.
checkout-api@v2.4.1 added the new order-history query path.
The user-id column needs to be added as an index.
The bad path is gated by feature_flag.new_checkout=true.
Disable the flag or roll back 7a3f9c2.
the configuration is what makes the oncall experience fun
None of this is automatic and doesn't come automatically, whether you are using Prometheus + Grafana, Datadog, or New Relic.
The good path needs deliberate plumbing:
Page on symptoms: error rate, latency, traffic, saturation, or SLO burn.
Put team , service , severity , runbook_url , and a scoped dashboard_url on the alert.
Propagate W3C trace context through every service.
Inject trace_id and span_id into structured logs.
Enable exemplars on the histogram used by the alert.
Configure Grafana so exemplars open Tempo.
Configure Tempo trace-to-logs so spans open Loki.
Emit service.version , deploy annotations, and commit SHA from CI/CD.
You can play around with such a well-configured setup here .
The useful AI SRE workflow starts after the observability stack has preserved the evidence trail. The agent can help with root cause analysis, propose the fix, and validate the patch. But if the alert drops the trace, the log correlation, and the deploy context, the agent has the same problem the human does: it is guessing.
For where we are taking this in the product, see FixBugs and AI SRE tools .
Google SRE: Monitoring Distributed Systems
Google SRE Workbook: Alerting on SLOs
Grafana: configure trace to logs correlation
Grafana: annotate visualizations
OpenTelemetry service semantic conventions
Back to Blog FixBugs AI-powered debugging helping teams ship reliable software faster.
© 2026 Modulo AI . All Rights Reserved.

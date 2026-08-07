---
source: "https://peachtrace.com/"
hn_url: "https://news.ycombinator.com/item?id=49214949"
title: "Show HN: PeachTrace – a new self-hosted observability platform"
article_title: "Home | PeachTrace"
author: "jhaygood86"
captured_at: "2026-08-07T19:44:32Z"
capture_tool: "hn-digest"
hn_id: 49214949
score: 1
comments: 1
posted_at: "2026-08-07T19:06:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PeachTrace – a new self-hosted observability platform

- HN: [49214949](https://news.ycombinator.com/item?id=49214949)
- Source: [peachtrace.com](https://peachtrace.com/)
- Score: 1
- Comments: 1
- Posted: 2026-08-07T19:06:23Z

## Translation

タイトル: 表示 HN: PeachTrace – 新しい自己ホスト型可観測性プラットフォーム
記事のタイトル: ホーム |ピーチトレース
説明: PeachTrace は、自己ホスト型の OpenTelemetry ネイティブの可観測性プラットフォームであり、すでに実行および制御しているインフラストラクチャ上にメトリクス、ログ、トレースを保存します。

記事本文:
ホーム |ピーチトレース
ピーチトレース
製品
メトリクス
パーセンタイル推定と時間バケット化された集計を使用して、サービスが発行するすべてのメトリックに対するアドホック クエリ。
ログ
完全に忠実な OTel ログ レコード。リアルタイムで検索可能で、トレース相関関係が組み込まれています。ボルトオンではありません。
痕跡
完全なウォーターフォール ビュー、相関ログ、エラー マーカーを使用して、すべてのサービスにわたるリクエストを追跡します。
クエリのパフォーマンス
サービスがすでに出力しているトレース データから完全に構築されたスロー クエリ アナライザーです。個別のデータベース エージェントをインストールする必要はありません。
API
PeachTrace は API ファーストで構築されています。ダッシュボードのすべての画面は、自分で呼び出すことができる同じパブリック REST API のクライアントにすぎません。
仕組み
導入
限定ベータ版
について
ベータ版に参加する
メニュー
限定ベータ版
独自のハードウェア上での OpenTelemetry ネイティブの可観測性。
PeachTrace は、OTLP メトリクス、ログ、トレースを直接取り込みます。独自のエージェントは使用せず、ユーザーが望まない限りデータがインフラストラクチャから流出することはありません。 Windows または Linux 上でセルフホストされるか、既に管理しているハードウェア上でホストされるか、Peach State Technologies によってホストされます。
メトリクス、ログ、トレースを 1 か所に
OpenTelemetry SDK を使用してサービスをインストルメントし、PeachTrace にポイントします。すべての信号は同じデータベースに格納され、相互に関連付けられ、すぐにクエリできます。
パーセンタイル推定と時間バケット化された集計を使用して、サービスが発行するすべてのメトリックに対するアドホック クエリ。
完全に忠実な OTel ログ レコード。リアルタイムで検索可能で、トレース相関関係が組み込まれています。ボルトオンではありません。
完全なウォーターフォール ビュー、相関ログ、エラー マーカーを使用して、すべてのサービスにわたるリクエストを追跡します。
サービスがすでに出力しているトレース データから完全に構築されたスロー クエリ アナライザーです。個別のデータベース エージェントをインストールする必要はありません。
PeachTrace は API ファーストで構築されています。ダッシュボードのすべての画面は、同じ p の単なるクライアントです。

自分で呼び出すことができる ublic REST API。
実際に相互に関係するメトリクス、ログ、トレース
ログ レコードは単なるテキストではありません。ログ レコードには、ログが発生した正確な場所を示すtrace_idとspan_idが含まれています。数回クリックするだけで、スパイクからトレース、その周囲の正確な丸太にジャンプできます。
ログおよびトレースと同じサービスおよびリソース属性を対象とした時系列。
すべてのスパンには、正確なログ行に結び付けるtrace_idとspan_idが含まれています。
ログ レコードには、trace_id/span_id が直接保存され、即座にトレースにジャンプできるようにインデックスが付けられます。
テレメトリは、すでに管理しているインフラストラクチャから離れる必要はありません。
Windows または Linux 上で PeachTrace を実行し、すでに所有しているハードウェア上で実行します。サードパーティのクラウドや GB ごとの取り込み料金は発生せず、機密リクエスト データを他人のサーバーに送信する必要もありません。自分で実行したくないですか? Peach State Technologies がホストし、管理します。
セルフホスト — Windows または Linux、独自のハードウェア上で
マネージド — Peach State Technologies がホストし、運用します
PeachTrace 限定ベータ版に参加する
限られた数のチームをオンボーディングしています。あなたの環境について教えてください。セットアップいたします。
ピーチトレース
独自のハードウェア上での OpenTelemetry ネイティブの可観測性

## Original Extract

PeachTrace is a self-hosted, OpenTelemetry-native observability platform that stores your metrics, logs, and traces on infrastructure you already run and control.

Home | PeachTrace
PeachTrace
Product
Metrics
Ad-hoc queries over every metric your services emit, with percentile estimation and time-bucketed aggregation.
Logs
Full-fidelity OTel log records, searchable in real time, with trace correlation built in — not bolted on.
Traces
Follow a request across every service with a full waterfall view, correlated logs, and error markers.
Query Performance
A slow-query analyzer built entirely from the trace data your services already emit — no separate database agent to install.
API
PeachTrace is built API-first — every screen in the dashboard is just a client of the same public REST API you can call yourself.
How It Works
Deployment
Limited Beta
About
Join the Beta
Menu
Limited Beta
OpenTelemetry-native observability, on your own hardware.
PeachTrace ingests OTLP metrics, logs, and traces directly — no proprietary agents, no data leaving your infrastructure unless you want it to. Self-hosted on Windows or Linux, on hardware you already control, or hosted for you by Peach State Technologies.
One place for metrics, logs, and traces
Instrument your services with any OpenTelemetry SDK and point them at PeachTrace. Every signal lands in the same database, correlated and ready to query.
Ad-hoc queries over every metric your services emit, with percentile estimation and time-bucketed aggregation.
Full-fidelity OTel log records, searchable in real time, with trace correlation built in — not bolted on.
Follow a request across every service with a full waterfall view, correlated logs, and error markers.
A slow-query analyzer built entirely from the trace data your services already emit — no separate database agent to install.
PeachTrace is built API-first — every screen in the dashboard is just a client of the same public REST API you can call yourself.
Metrics, logs, and traces that actually point at each other
A log record isn't just text — it carries the trace_id and span_id that put it exactly where it happened. Jump from a spike, to the trace, to the exact logs around it, in a few clicks.
Time series scoped to the same service & resource attributes as your logs and traces.
Every span carries the trace_id and span_id that ties it back to its exact log lines.
Log records store trace_id/span_id directly, indexed for instant jump-to-trace.
Your telemetry never has to leave infrastructure you already control.
Run PeachTrace on Windows or Linux, on hardware you already own — no third-party cloud, no per-GB ingestion bill, no sending sensitive request data to someone else's servers. Prefer not to run it yourself? Peach State Technologies can host and manage it for you.
Self-hosted — Windows or Linux, on your own hardware
Managed — Peach State Technologies hosts and operates it for you
Join the PeachTrace limited beta
We're onboarding a limited number of teams. Tell us about your environment and we'll get you set up.
PeachTrace
OpenTelemetry-native observability, on your own hardware

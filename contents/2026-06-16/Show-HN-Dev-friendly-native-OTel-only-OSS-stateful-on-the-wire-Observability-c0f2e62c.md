---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48560279"
title: "Show HN: Dev-friendly native OTel: only OSS stateful, on-the-wire Observability"
article_title: ""
author: "jratkevic"
captured_at: "2026-06-16T19:18:25Z"
capture_tool: "hn-digest"
hn_id: 48560279
score: 3
comments: 0
posted_at: "2026-06-16T19:00:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dev-friendly native OTel: only OSS stateful, on-the-wire Observability

- HN: [48560279](https://news.ycombinator.com/item?id=48560279)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T19:00:53Z

## Translation

タイトル: HN を表示: 開発に適したネイティブ OTel: OSS のみのステートフル、オンザワイヤの可観測性
HN テキスト: こんにちは、HN。私たちは MyDecisive.ai のチームです。今日は開発者向けに、OpenTelemetry のポイント アンド クリック コントロールと可視性である Octant を紹介します。特に K8S クラスターを管理している場合は、「可観測性税」の苦痛を感じたことがあるでしょう。最新の標準は、すべてを OpenTelemetry で計測することですが、これらの豊富な OTLP ログ、メトリクス、トレースをすべて SaaS ベンダー (Datadog、Splunk、Honeycomb) に直接パイプすると、すぐにコストが高くなります。何かが壊れたときに検索できるようにするためだけに、ノイズが多く価値の低いデータの取り込みと保管に多大なコストを支払うことになります。 Octant を使用すると、数分で OTel を起動して実行できます。このモデルを反転するために Octant を構築しました。 Octant は、すべてのテレメトリをクラスタ外にやみくもに送信するのではなく、OTEL クラスタを構成し、管理するのに役立ちます。 K8s オブジェクトを管理するための視覚的なインターフェイスを提供しますが、重要なのは、データが VPC から離れる前にソースでデータをフィルタリングする OTLP ゲートウェイとして機能することです。 OpenTelemetry をネイティブに話すため、アプリケーション コードに触れることなく、既存の OTel SDK またはコレクタを直接 OpenTelemetry に向けることができます。内部で行われる機能は次のとおりです。 - OTel ネイティブ トレースおよびログ サンプリング: OTLP トラフィックの取り込みを容易にし、回線上のログとトレースを検査します。何を保持するかを決定する前にトレースの完全なコンテキストを待つことで、ブレイディングの約束を果たし、周囲の実用的な信号 (エラーや高遅延スパンなど) を 100% 保持しながら、SaaS の請求額に達する前にジャンクを削除します。 - 処理中のステートフル アラート: 外部プロバイダーによってデータがバッチ化、転送され、インデックス付けされてアラートがトリガーされるのを待つ代わりに、Octant は処理中のテレメトリ ストリームを処理できます。この縮み

検出ギャップがなくなり、そもそも SaaS ベンダーの必要性が減ります。
- オンザワイヤー PII 編集: ログやトレースがインターネット上に送信される前に、機密情報をリアルタイムで検出して削除し、「取り込み後」のクリーンアップ コストとコンプライアンスのリスクを排除します。 - K8s コンテキスト インジェクション: クラスターと深く統合されているため、統合 UI で OTel ストリームを K8s リソース (デプロイメント、ポッド、CRD) に直接マッピングします。 API は Go ([github.com/mydecisive/octant]) で構築されており、スタック全体は Helm チャートを介してクラスターに直接デプロイできます。開発クラスター上でスピンアップして分解していただければ幸いです。私たちはつい最近、最初のコミュニティ貢献者からの PR をマージしました。これは私たちにとって大きなマイルストーンでした。私たちはその勢いを維持したいと考えています。K8 の可観測性と自律性のハッキングに興味がある場合は、 OpenTelemetry パイプライン、または Go/React については、いくつかの「優れた最初の問題」をタグ付けしました。GitHub: https://github.com/MyDecisive/octant Web サイト: https://www.mydecisive.ai/ 今日はこのスレッドに参加し、質問に答えたり、アーキテクチャについて詳しく調べたりすることに喜んで応じます。

## Original Extract

Hi HN, We’re the team at MyDecisive.ai, and today we’re giving developers a peek at Octant — point-and-click control and visibility for your OpenTelemetry. You've likely felt the pain of the "observability tax," especially if you manage K8S clusters. The modern standard is to instrument everything with OpenTelemetry, but piping all those rich OTLP logs, metrics, and traces straight to a SaaS vendor (Datadog, Splunk, Honeycomb) gets expensive fast. You end up paying massive ingestion and storage costs for noisy, low-value data just so it's searchable when something breaks. With Octant you get up and running on OTel in minutes. We built Octant to flip this model. Instead of blindly shipping all telemetry off-cluster, Octant configures and helps to manage OTEL clusters. It gives you a visual interface for managing K8s objects, but importantly, it acts as an OTLP gateway that filters data at the source before it leaves your VPC. Because it natively speaks OpenTelemetry, you can point your existing OTel SDKs or collectors right at it without touching your application code. Here is what it does under the hood: - OTel-Native Trace & Log Sampling: It makes it easy to ingest OTLP traffic and inspects logs and traces on the wire. By waiting for the full context of a trace before determining what to keep, it delivers on the promise of braiding, retaining 100% of the actionable signals around (like errors and high-latency spans) but droppings the junk before it hits your SaaS bill. - In-Flight Stateful Alerting: Instead of waiting for data to be batched, shipped, and indexed by an external provider to trigger an alert, Octant can process the telemetry streams in-flight. This shrinks the detection gap and reduces the need for SaaS vendors in the first place.
- On-the-Wire PII Redaction: It can detect and strip sensitive information from your logs and traces in real-time before they are transmitted over the internet, removing "post-ingestion" clean-up costs and compliance risks. - K8s Context Injection: Because it's deeply integrated with your cluster, it maps your OTel streams directly to your K8s resources (Deployments, Pods, CRDs) in a unified UI. The API is built in Go ([github.com/mydecisive/octant] and the whole stack can be deployed directly into your cluster via our Helm charts. We’d love for you to spin it up on a dev cluster and tear it apart. We just recently merged a PR from our very first community contributor, which was a huge milestone for us! We want to keep that momentum going. If you're interested in hacking on K8s observability and autonomy, OpenTelemetry pipelines, or Go/React, we’ve tagged a few 'good first issues' and would be thrilled to welcome you to the project. GitHub: https://github.com/MyDecisive/octant Website: https://www.mydecisive.ai/ I'll be hanging out in the thread today and am happy to answer any questions or dig into the architecture!


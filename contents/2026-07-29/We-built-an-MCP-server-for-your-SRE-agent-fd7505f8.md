---
source: "https://clickhouse.com/blog/benchmarking-the-clickstack-mcp-server-with-hdx-evals"
hn_url: "https://news.ycombinator.com/item?id=49097958"
title: "We built an MCP server for your SRE agent"
article_title: "How we build and evaluate our MCP server for SRE agents | ClickHouse"
author: "mikeshi42"
captured_at: "2026-07-29T15:06:34Z"
capture_tool: "hn-digest"
hn_id: 49097958
score: 2
comments: 0
posted_at: "2026-07-29T14:25:29Z"
tags:
  - hacker-news
  - translated
---

# We built an MCP server for your SRE agent

- HN: [49097958](https://news.ycombinator.com/item?id=49097958)
- Source: [clickhouse.com](https://clickhouse.com/blog/benchmarking-the-clickstack-mcp-server-with-hdx-evals)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T14:25:29Z

## Translation

タイトル: SRE エージェント用の MCP サーバーを構築しました
記事のタイトル: SRE エージェント用の MCP サーバーを構築および評価する方法 |クリックハウス
説明: 決定論的な合成インシデント、サンドボックス化されたクロード エージェント、ブラインド LLM グレーディングを使用して、生の SQL ベースラインに対して ClickStack MCP サーバーのベンチマークを行うために構築されたオープンソース フレームワークである hdx-evals の舞台裏と、MCP が高スコアを獲得した理由

記事本文:
。 -->
SRE エージェント用の MCP サーバーを構築および評価する方法 | ClickHouse コンテンツにスキップ ロゴを SVG としてコピー 完全なロゴをダウンロード ロゴマークをダウンロード 検索を開く 地域セレクターを開く 英語
製品 製品 ClickHouse Cloud ClickHouse の最適な使い方。
AWS、GCP、Azure で利用できます。
Bring Your Own Cloud フルマネージドの ClickHouse サービス、
自分の AWS および GCP アカウントにデプロイされます。
ClickHouse によって管理される Postgres トランザクション用の統合データ スタック
そして分析。
マネージド ClickStack 高パフォーマンスによるマネージドオブザーバビリティ
クエリと長期保存。
Langfuse Cloud LLM の可観測性と評価
信頼性の高い AI アプリケーションとエージェントを実現します。
オープンソース ClickHouse 高速オープンソース OLAP データベース
リアルタイム分析。
ClickStack ログ用のオープンソース可観測性スタック、
メトリクス、トレース、セッションのリプレイ。
Agentic Data Stack AI を活用したアプリケーションを構築
クリックハウスで。
chDB インプロセス SQL エンジンを搭載
ClickHouse、Pandas 互換 API を使用
ソリューション ユースケース リアルタイム分析
リソース 会社のリソース ユーザーストーリー
ページをコピーしました コピーされました!その他のアクション マークダウンで表示 このページをマークダウンで開く
ChatGPT で開く このページについて質問する
クロードで開く このページについて質問する
v0 で開く このページについて質問する
SRE エージェント用の MCP サーバーを構築および評価する方法
先月、サンフランシスコで開催された Open House で、私たちは ClickStack MCP サーバーを紹介しました。これは AI エージェントに生の SQL を手動で組み立てるのではなく、本番環境のインシデントを調査するための構造化された高レベルのプリミティブを提供する専用の可観測性ツールのセットです。ステージでは、一般的な ClickHouse SQL インターフェイスに対して ClickStack MCP をベンチマークした初期の結果を共有しました。根本原因と修復の結果が 18% 正確になり、ツール呼び出しが 26% 減少し、一貫性が 2.4 倍向上しました。

結果はオーバーラン。
これらの数値は、どのように測定したかについて多くの質問を引き起こしたため、この投稿では、私たちが何を行ったかを正確に説明します。
エージェントが本番環境のインシデントを調査するために使用する MCP (Model Context Protocol) ツールを配布すると、新たな種類のリスクが生じます。同じ質問を 2 回行っても、常に同じ答えが得られるわけではありません。リファクタリングされたツール、名前変更されたパラメータ、またはクエリ ロジックの変更により、構造化された測定方法がなければ気付かない形で、調査の品質が静かに低下する可能性があります。
私たちは 2 つのことに答える必要がありました。まず、MCP の各新しいバージョンが前回よりも優れた調査結果をもたらしたかどうか。次に、MCP が ClickHouse への直接 SQL アクセスを持つエージェントよりも実際に優れたパフォーマンスを示したかどうかです。これを行うために、同一の合成テレメトリをシードし、各構成に対してクロード エージェントを実行し、実装間で結果を盲目的にスコア付けする再現可能なベンチマーク フレームワークである hdx-evals を構築しました。フレームワークは MCP とモデルのあらゆる組み合わせを実行するため、有益な副次的な利点が得られます。新しいモデルが出荷されると、実稼働環境で何を実行するかを決定する前に、hdx-evals をそのモデルに向けて同じシナリオでどのように実行されるかを確認できます。
hdx-evals はオープンソースであり、HyperDX リポジトリに存在します。コードに従って進みたい場合は、 github.com/hyperdxio/hyperdx/tree/main/packages/hdx-eval を参照してください。
ClickStack MCP と ClickHouse MCP #
ベンチマーク シナリオに移る前に、ClickHouse MCP と並行して専用の ClickStack MCP を構築した理由を簡単に説明します。すでにこの違いを理解している場合は、スキップしてインシデントのシナリオに進んでください。
ClickHouse MCP を使用すると、エージェントは SQL を通じて ClickHouse に直接アクセスできます。これは柔軟ですが、スキーマの検出、各クエリの構築、および組み立てをモデルに任せます。

ble の複数ステップの調査ワークフロー自体。可観測性の調査では、これにより、余分なツールの呼び出し、非効率なクエリ、カーディナリティの高い結果、実行間の一貫性のないアプローチが発生する可能性があります。
その代わりに、ClickStack MCP は、繰り返し発生するイベント パターンの検索、時間枠の比較、異常値の特定、ログとトレース間の移動などの一般的な可観測性タスクのための高レベルのセマンティック ツールを提供します。これらのツールは引き続き SQL を実行しますが、クエリ ロジックをパッケージ化し、エージェントが使いやすい構造化された結果を返します。目的は、ツールの呼び出しや一般的なクエリ エラーを減らしながら、精度と一貫性を向上させることです。以下のシナリオでは、これらの利点がさまざまな種類の調査にわたって維持されるかどうかをテストします。
ベンチマークの良さは、テスト対象となるインシデントによって決まります。きれいすぎると、どのエージェントも素晴らしく見えます。あまりにも非現実的である場合、結果からは実稼働環境での MCP のパフォーマンスについて何もわかりません。
したがって、各 hdx-evals シナリオは実際のインシデントをシミュレートします。数千万の合成ログとスパンが決定論的にシードされるため、すべての実行が同一のデータから開始され、現実的なノイズの中に埋め込まれた異常が埋め込まれ、さらに注意をそらす要素が追加されます。二次的な問題は、実際の問題のように見えるようにタイミングと表現が定められています。エージェントが最初に発生したエラーにジャンプすると失敗するはずです。データを基に推論するエージェントはそうすべきではありません。
データは、OpenTelemetry デモ、既存の製品、または記録されたテレメトリから派生したものではなく、完全に合成されたものです。パブリック データセットには、モデルのトレーニング データにすでに表示されている可能性のある既知の文書化された障害モードが含まれていることがよくあります。それらを再利用すると、エージェントは利用可能なツールを使用して調査するのではなく、よく知られたインシデントを認識できるようになるため、オリジナルのシナリオを生成することが役立ちます。

予備知識ではなく、エージェントの調査能力をテストします。
トレースとログは、シードされた擬似乱数ジェネレーター (PRNG) から TypeScript で決定論的に生成され、ClickStack の実際の OTel スキーマを反映する ClickHouse テーブルに直接書き込まれます。トポロジは、単一の API サーバーから、最大 25 の名前付き電子商取引サービスと、カーディナリティに対して手続き的に生成される最大 100 のバックグラウンド サービスまで、シナリオごとに異なります。量は意図的に行われます。エラーの根本原因だけでは、1,200 万以上のスパンと 1,200 万のログ内に 8 つの失敗したトレースのみが植え付けられます。そのため、評価では、エージェントがエラーをカウントできるかどうかではなく、健全なベースライン上に類似した妨害要因から植え付けられた小さなインシデントを分離できるかどうかが測定されます。
合計 5 つのシナリオを構築し、それぞれが異なる調査スキルを対象としました。
error-root-cause は、チェックアウト 500 秒にカスケードする支払いサービス データベースのタイムアウトであり、25 のサービスにわたる 2,400 万のスパンとログの間に隠されています。 6 つのディストラクタには、実際の信号の 10 倍の音量で CDN エラー バーストが含まれており、最初に WHERE StatusCode = 'ERROR' に到達したエージェントを騙すように設計されています。
latency-spike は、インデックスが欠落しており、5,000 のテナントにわたる 2,000 万のスパンに埋もれているため、エンタープライズ テナントのみに影響を与える p99 (99 パーセンタイル レイテンシ) 回帰です。特徴フラグ交絡因子は、影響を受けるセグメントと 80% の相関関係がありますが、因果関係はありません。無関係なエンドポイントにある同じ遅延範囲のディストラクターは、同じ時間枠内にあります。
noisy-signals は、取り込みと保管のコストを削減するために、ドロップまたはスロットリングできるログを特定するようにエージェントに要求します。データセットには 1,600 万件のログが含まれています。大量で価値の低い各ログ パターンは、同じサービスおよび重大度レベルの同様に一般的だが重要なパターンとペアになっています。たとえば、notification-service からすべての DEBUG ログを削除すると、両方のログが削除されます。

日常的なキャッシュヒットメッセージと、通知が送信されたことを確認するために必要な配信レコード。
service-health-check は、インシデントがまったく発生していない 3,600 万イベントの 4 時間の単一サービス データセットです。 4 つの微妙な新しいシグナルが比例的に表面化する必要があり、2 つの繰り返しパターンがエスカレートしてはなりません。これは他のベンチマークと同じくらい重要です。なぜなら、良いベンチマークの半分は正しい答えを評価し、残りの半分は自信を持って間違った答えをペナルティにするからです。
セグメント回帰は、600 万のトレースにわたるエラー スパイクであり、2 つの次元 (エンタープライズ層とキャッシュ ミス) の交差点でのみ現れます。どちらの軸だけでも問題は明らかにならないため、エージェントは問題を見つけるためにクロス集計する必要があります。それは教科書的なシンプソンのパラドックス設定です。
ノイズの多い信号を例に挙げます。エージェントは、1,600 万のログのうち、重要な運用情報を失わずに削除または調整できるログを判断する必要があります。通知サービスの DEBUG ログの半分は、削除しても安全な日常的なキャッシュ ヒット メッセージです。残りの半分は通知配信を記録し、監査証跡として保持する必要があります。
2 種類のログは同じサービスと重大度レベルを共有しているため、これらのフィールドだけをフィルタリングしてもそれらを分離することはできません。メッセージ内の値は大きく異なるため、完全なログ メッセージによるグループ化も役に立ちません。代わりに、エージェントは同様のメッセージをパターンにグループ化し、使い捨てのキャッシュ ログと重要な配信レコードを区別する必要があります。このシナリオでは、MCP のイベント パターン ツールがエージェントにそれを行うための正しい方法を提供するかどうかをテストします。
テレメトリのセットアップとシード #
エージェントがインシデントに対して実行される前に、hdx-evals は、プロビジョニング、スキーマ、データ生成の 3 つの段階で ClickStack 環境を構築します。
プロビジョニングにより評価アカウントと接続が作成されます

、ソース、および ClickStack MCP (HTTP 経由) と ClickHouse MCP (stdio 経由) の両方の MCP サーバー定義。ダッシュボード、アラート、保存された検索など、11 の非調査ツールの拒否リストにより、エージェントはクエリ ツールのみに集中できます。すべては単一の eval.config.json に書き込まれ、セットアップは冪等であるため、安全に再実行できます。
スキーマは、ClickStack の実稼働 OTel テーブル ( CREATE TABLE ... AS default.otel_traces ) の構造クローンとして eval テーブルを作成するため、それらは近似ではなく実際のスキーマ、エンジン、インデックスを継承します。各シナリオには、生のトレースとログ、および ClickStack の実稼働メタデータ検出パスを反映する 2 段階のマテリアライズド ビュー パイプラインから構築されたロールアップ テーブルの 6 つのテーブルが含まれます。それは、MCP の「どのようなフィールドが存在するか?」ということを意味します。クエリは生データをスキャンするのではなく、事前に集計されたロールアップにヒットします。これは、ClickStack UI でのオートコンプリートとファセット生成を強化するのと同じ最適化です。これらの実体化ビューがどのように機能するかについて詳しくは、ClickStack ドキュメントを参照してください。
データ生成により、実際のテレメトリが書き込まれます。すべてのランダム性は単一のシードされた PRNG を介して流れるため、特定のシードおよびアンカー時間では常にバイト同一のデータが生成されます。これは実行ごとに同じデータセットであるため、MCP 間の A/B 比較が意味のあるものになります。データはバッチでストリームされるため、数百万行であってもメモリの制限が維持されます。
各シナリオでは固定の「現在」時間が使用され、指示とともにエージェントに渡されます。これにより、評価がいつ実行されるかに関係なく、「過去 10 分間」などの語句がシード データ内の同じ期間を常に参照するようになります。したがって、データセットは一度生成すれば再利用できます。各実行前に、フレームワークはそれが存在することを確認し、必要に応じて作成します。
データがシードされると、HD

x-evals は、すべての (MCP、モデル、実行) の組み合わせに対して実際のクロード コード プロセスを生成し、SRE が行うような方法でシナリオを調査します。足場やヒントはなく、質問とツールだけを使用します。
実行ごとに独自のサンドボックスが取得されます。つまり、新しい一時ディレクトリ、単一の MCP サーバー定義 (評価対象のみ)、および調査ツール以外のすべてを削除するツール許可ファイルです。 bash、書き込み、編集、glob、または Webfetch の機能はなく、読み取りアクセスの範囲はその実行自体の一時ディレクトリのみに限定されます。これは私たちが思いつきで追加した予防措置ではありませんでした。初期のイテレーションでは、クロードが以前の実行出力、スコアリング基準、または真実の答えを探してファイルシステムを調べていることがわかりました。事実上、調査ではなく不正行為の方法を探していました。各実行を独自の使い捨てのブラインドサンドボックスにロックすると、そのドアは完全に閉じられます。リポジトリのルート、以前の実行、内部から見える評価設定はありません。
各サンドボックスは複数の分離層を使用します。ファイルシステムの制限はエージェントがアクセスできる内容を制御し、拒否リストはエージェントが実行できるアクションを制限します。各実行は独自のプロセス グループでも実行されるため、ランナーはエージェントとサブプロセスを確実に終了できます。

[切り捨てられた]

## Original Extract

A behind-the-scenes look at hdx-evals, the open-source framework we built to benchmark the ClickStack MCP server against a raw SQL baseline using deterministic synthetic incidents, sandboxed Claude agents, and blind LLM grading — and why the MCP scored hi

. -->
How we build and evaluate our MCP server for SRE agents | ClickHouse Skip to content Copy logo as SVG Download full logo Download logomark Open search Open region selector English
Products Products ClickHouse Cloud The best way to use ClickHouse.
Available on AWS, GCP, and Azure.
Bring Your Own Cloud A fully managed ClickHouse service,
deployed in your own AWS and GCP account.
Postgres managed by ClickHouse Unified data stack for transactions
and analytics.
Managed ClickStack Managed observability with high-performance
queries and long-term retention.
Langfuse Cloud LLM observability and evaluations
for reliable AI applications and agents.
Open source ClickHouse Fast open-source OLAP database for
real-time analytics.
ClickStack Open-source observability stack for logs,
metrics, traces, and session replays.
Agentic Data Stack Build AI-powered applications
with ClickHouse.
chDB In-process SQL Engine powered by
ClickHouse, with a Pandas-compatible API
Solutions Use cases Real-time analytics
Resources Company resources User stories
Copy page Copied! More actions View as Markdown Open this page in Markdown
Open in ChatGPT Ask questions about this page
Open in Claude Ask questions about this page
Open in v0 Ask questions about this page
How we build and evaluate our MCP server for SRE agents
Last month at Open House in San Francisco, we introduced the ClickStack MCP server , a set of purpose-built observability tools that give AI agents structured, high-level primitives for investigating production incidents instead of hand-assembling raw SQL. On stage, we shared early results from benchmarking the ClickStack MCP against a generic ClickHouse SQL interface: 18% more accurate root cause and remediation outcomes, 26% fewer tool calls, and 2.4x more consistent results run over run.
Those numbers prompted plenty of questions about how we measured them, so this post explains exactly what we did.
Shipping MCP (Model Context Protocol) tools that agents use to investigate production incidents introduces a new kind of risk. The same question asked twice doesn't always produce the same answer, and a refactored tool, a renamed parameter, or a change in query logic can quietly degrade investigation quality in ways you won't catch without a structured way to measure them.
We needed to answer two things. First, whether each new version of the MCP produced better investigative outcomes than the last. Second, whether the MCP actually outperformed an agent with direct SQL access to ClickHouse. To do that, we built hdx-evals , a reproducible benchmarking framework that seeds identical synthetic telemetry, runs Claude agents against each configuration, and scores the results blindly across implementations. Because the framework runs every combination of MCP and model, a useful side benefit fell out of it: when a new model ships, we can point hdx-evals at it and see how it performs on the same scenarios before deciding what to run in production.
hdx-evals is open source and lives in the HyperDX repository. If you want to follow along with the code, see github.com/hyperdxio/hyperdx/tree/main/packages/hdx-eval .
ClickStack MCP vs ClickHouse MCP #
Before we turn to the benchmark scenarios, we’ll briefly revisit why we built a dedicated ClickStack MCP alongside the ClickHouse MCP. If you’re already familiar with the distinction, you can skip ahead to the incident scenarios.
The ClickHouse MCP gives agents direct access to ClickHouse through SQL. This is flexible, but it leaves the model to discover the schema, construct each query, and assemble multi-step investigation workflows itself. In observability investigations, this can lead to extra tool calls, inefficient queries, high-cardinality results, and inconsistent approaches between runs.
The ClickStack MCP instead provides higher-level semantic tools for common observability tasks, such as finding recurring event patterns, comparing time windows, identifying outliers, and moving between logs and traces. These tools still execute SQL underneath, but they package the query logic and return structured results that are easier for an agent to use. The aim is to improve accuracy and consistency while reducing tool calls and common query errors. The scenarios below test whether those advantages hold across different types of investigation.
A benchmark is only as good as the incidents it tests against. If they're too clean, every agent looks brilliant. If they're too unrealistic, the results tell you nothing about how the MCP performs in production.
So each hdx-evals scenario simulates a real incident: tens of millions of synthetic logs and spans, seeded deterministically so every run starts from identical data, with a planted anomaly buried inside realistic noise, plus distractors: secondary issues timed and worded to look like the real problem. An agent that jumps to the first error it sees should fail. An agent that reasons through the data should not.
The data is completely synthetic rather than derived from the OpenTelemetry Demo, existing products, or recorded telemetry. Public datasets often contain well-known and documented failure modes that may already appear in a model’s training data. Reusing them could allow an agent to recognize a familiar incident instead of investigating it through the available tools, so generating original scenarios helps us test the agent’s investigative ability rather than its prior knowledge.
Traces and logs are generated deterministically in TypeScript from a seeded Pseudo-Random Number Generator (PRNG) and written directly into ClickHouse tables that mirror ClickStack's real OTel schema. Topology varies per scenario, from a single api-server up to 25 named e-commerce services plus ~100 procedurally-generated background services for cardinality. Volume is deliberate: error-root-cause alone plants only 8 failing traces inside 12M+ spans and 12M logs, so the eval measures whether the agent can isolate a small planted incident from look-alike distractors on top of a healthy baseline, not whether it can count errors.
We built five scenarios in total, each targeting a different investigative skill:
error-root-cause is a payment-service database timeout cascading into checkout 500s, hidden among 24M spans and logs across 25 services. Six distractors include a CDN error burst at 10x the volume of the real signal, designed to fool any agent that reaches for WHERE StatusCode = 'ERROR' first.
latency-spike is a p99 (99th-percentile latency) regression affecting only enterprise tenants because of a missing index, buried in 20M spans across 5,000 tenants. A feature-flag confounder correlates 80% with the affected segment but isn't causal, and a same-latency-range distractor at an unrelated endpoint sits in the same time window.
noisy-signals asks the agent to identify logs that can be dropped or throttled to reduce ingestion and storage costs. The dataset contains 16 million logs. Each high-volume, low-value log pattern is paired with an equally common but important pattern from the same service and severity level. For example, dropping all DEBUG logs from notification-service would remove both routine cache-hit messages and the delivery records needed to confirm that notifications were sent.
service-health-check is a four-hour, single-service dataset of 36M events with no incident at all. Four subtle novel signals must be surfaced proportionally, and two recurring patterns must not be escalated. This one matters as much as the others, because half of a good benchmark is rewarding the right answer and the other half is penalizing a confident wrong one.
segmented-regression is an error spike across 6M traces that only appears at the intersection of two dimensions (enterprise tier and cache miss). Neither axis alone reveals the problem, so the agent has to cross-tabulate to find it. It's a textbook Simpson's Paradox setup.
Take noisy-signals as an example. The agent must determine which of 16 million logs can be dropped or throttled without losing important operational information. Half of the notification-service DEBUG logs are routine cache-hit messages that are safe to remove. The other half record notification deliveries and must be retained as an audit trail.
The two types of logs share the same service and severity level, so filtering on those fields alone cannot separate them. Grouping by the complete log message does not help either, because values within the messages vary widely. The agent must instead group similar messages into patterns, then distinguish the expendable cache logs from the essential delivery records. This scenario tests whether the MCP’s event patterns tool gives the agent the right way to do that.
Setting up and seeding telemetry #
Before any agent runs against an incident, hdx-evals builds out a ClickStack environment in three stages: provisioning, schema, and data generation.
Provisioning creates an evaluation account, connections, and sources, plus MCP server definitions for both the ClickStack MCP (over HTTP) and the ClickHouse MCP (over stdio). A deny-list of 11 non-investigation tools, such as dashboards, alerts, and saved searches, keeps the agent focused on query tools only. Everything is written to a single eval.config.json , and setup is idempotent, so it can be re-run safely.
Schema creates the eval tables as structural clones of ClickStack’s production OTel tables ( CREATE TABLE ... AS default.otel_traces ), so they inherit the real schema, engine, and indexes rather than an approximation. Each scenario gets six tables: raw traces and logs, plus rollup tables built from a two-stage materialized view pipeline that mirrors ClickStack’s production metadata discovery path. That means the MCP's "what fields exist?" queries hit pre-aggregated rollups instead of scanning raw data — the same optimization that powers autocomplete and facet generation in the ClickStack UI. You can read more about how these materialized views work in the ClickStack docs .
Data generation writes the actual telemetry. All randomness flows through a single seeded PRNG, so a given seed and anchor time always produce byte-identical data. That's the same dataset every run, which is what makes A/B comparisons between MCPs meaningful. Data streams in batches to keep memory bounded, even at millions of rows.
Each scenario uses a fixed “current” time, which is passed to the agent with its instructions. This ensures that phrases such as “in the last 10 minutes” always refer to the same period in the seeded data, regardless of when the evaluation runs. The dataset can therefore be generated once and reused; before each run, the framework checks that it exists and creates it if necessary.
Once data is seeded, hdx-evals spawns a real Claude Code process for every (MCP, model, run) combination and lets it investigate the scenario the way an SRE would, with no scaffolding and no hints, just the question and the tools.
Each run gets its own sandbox: a fresh, ephemeral temp directory, a single MCP server definition (only the one being evaluated), and a tool permission file that strips out everything except investigation tools. There's no bash, write, edit, glob, or webfetch capabilities, and read access is scoped only to that run's own temp directory. This wasn't a precaution we added on a whim. In early iterations, we found Claude poking around the filesystem looking for prior run output, scoring criteria, or ground-truth answers - effectively looking for a way to cheat rather than investigate. Locking each run into its own disposable, blind sandbox closes that door entirely. There's no repo root, no prior runs, and no eval config visible from inside.
Each sandbox uses several layers of isolation. Filesystem restrictions control what the agent can access, while the deny-list limits which actions it can take. Each run also executes in its own process group, allowing the runner to terminate the agent and any subprocesses reliably d

[truncated]

---
source: "https://www.oodle.ai/product/agent-observability"
hn_url: "https://news.ycombinator.com/item?id=48907615"
title: "Show HN: Oodle.ai – $10 per million agent traces"
article_title: "Agent Observability that Just Works at Scale - Oodle.ai | Oodle AI"
author: "kirankgollu"
captured_at: "2026-07-14T15:17:48Z"
capture_tool: "hn-digest"
hn_id: 48907615
score: 11
comments: 1
posted_at: "2026-07-14T14:36:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Oodle.ai – $10 per million agent traces

- HN: [48907615](https://news.ycombinator.com/item?id=48907615)
- Source: [www.oodle.ai](https://www.oodle.ai/product/agent-observability)
- Score: 11
- Comments: 1
- Posted: 2026-07-14T14:36:28Z

## Translation

タイトル: HN を表示: Oodle.ai – エージェント トレース 100 万件につき 10 ドル
記事のタイトル: 大規模に機能するエージェントの可観測性 - Oodle.ai |ウードルAI
説明: LLM 呼び出し、API、インフラ全体にわたるエンドツーエンドの可視性 - そのため、期待ではなく自信を持って AI を出荷できます。エージェントの可観測性は 100 万スパンあたり 10 ドルです。
HN テキスト: こんにちは HN、私たちはキランとビジェイです!過去 2 年間にわたり、私たちはログ、メトリクス、トレースなどの可観測性を実現するカラムナ型ストレージ エンジンを構築してきました。今日、私たちがその基盤の上に構築したもの、つまり LLM エージェントの可観測性を示すことは、私たちにとって非常に興味深いことです。エージェントがいかに非決定的であるかを考えると、サンプリングせずにすべてのトレースを保存することが私たちにとって重要でした。しかし、これらのトレースは MB、場合によっては GB 単位になることが多く、安価に保存する必要がありました。また、クエリと分析を高速にする必要もありました。これら両方の目標を達成するために、私たちはそれらを独自の寄木細工のようなファイル形式で S3 に保存し、AWS Lambda を使用してクエリを実行します。すべてのトレースの各スパンを処理するため、それぞれに対して LLM ベースの評価を実行するのではなく、最初に決定論的な手法を使用して分析します。ツールの障害、再試行、ループ、異常なトークンの使用、レイテンシー回帰、スキーマ違反、センチメント、およびその他の実稼働シグナルを検出します。このアプローチについては、こちらで詳しく書いています: https://blog.oodle.ai/you-cant-sample-your-way-to-reliable-a... 独自のエンジン、サンプリングなし、LLM-for-evals 前の確定的処理の組み合わせにより、100 万トレースあたり 10 ドルの価格を実現し、1 秒未満の p99 クエリ レイテンシを提供し、健全なマージンを確保できます。これを構築する前は、独自のエージェントの可観測性のために Langfuse を使用していましたが、これは 6 倍高価でした。まだ非常に初期段階で、荒削りな部分もありますが、ご質問やフィードバックをお待ちしております。

記事本文:
大規模に機能するエージェントの可観測性 - Oodle.ai | Oodle AI 新しいエージェントのオブザーバビリティはライブです。エージェントは 100 万スパンあたり 10 ドルから追跡します。 → 製品 AI エージェントのデバッグ エージェントの可観測性 新しい AI コスト管理 可観測性 メトリクス ログ トレース APM インフラストラクチャ Kubernetes モニタリング BYOC テクノロジー エクスペリエンス リアル ユーザー モニタリング 合成モニタリング 移行 1-Click 移行 すべての製品を見る お客様 Elastic Grafana を切り替える理由 Datadog Elastic Cloud ELK スタック OpenSearch すべての比較を見る → 価格設定の概要 エージェントの可観測性 企業ドキュメント リソース シリーズ A スタートアップ向けソリューション ユニコーン向け フォーチュン グローバル向け2000 のユースケース SaaS プラットフォーム オブザーバビリティ 移行ガイド Grafana Labs ELK スタック セキュリティと信頼性 大学ブログ カンファレンス トーク ケーススタディ TCO 計算機 Discord ステータス ログイン US1 地域 AP1 地域 プレイグラウンド スケジュール デモを試す 製品の顧客 スイッチの理由 価格設定 会社ドキュメント リソース ログイン プレイグラウンド スケジュール デモを試す 配送エージェントは簡単です。エージェントを監視するのは難しいです。
エージェント トレースの検索が遅い
エージェント トレースには大規模なメタデータ (プロンプト、ツール呼び出しなど) が含まれており、検索結果の読み込みには数分かかる場合があります。
エージェント トレースの保存にはコストがかかります
エージェントは非決定的に失敗するため、いかなる痕跡も見逃すわけにはいきませんが、すべてを保持するわけにもいきません。
エージェントの失敗を表面化するのは難しい
エージェントはさまざまな方法で失敗します。偽の「やりました」メッセージ、サイレントツールエラー、チャット途中での放棄などです。
Oodle: 大規模に「適切に機能する」エージェントの可観測性
カラムナ型ストレージ、サーバーレス コンピューティング - トレースを数分ではなく数秒で見つけます。
トレースを 100% 保存し、予算も守ります。数か月または数年間保持するかどうかは、あなたが選択できます。
単一の評価を作成する前に、ユーザーのフラストレーション、ツール呼び出しの最適化、異常を検出します。
1 日あたり 300 万件以上のエージェント トレース。ゼロ

サンプリング。
Voice AI エージェント上のサイレント障害を捕捉します。
配送業者は簡単です。エージェントを監視するのは難しいです。
エージェント トレースの保存にはコストがかかります
エージェントは非決定的に失敗するため、いかなる痕跡も見逃すわけにはいきませんが、すべてを保持するわけにもいきません。
トレースを 100% 保存し、予算も守ります。数か月または数年間保持するかどうかは、あなたが選択できます。
エージェント トレースの検索が遅い
エージェント トレースには大規模なメタデータ (プロンプト、ツール呼び出しなど) が含まれており、検索結果の読み込みには数分かかる場合があります。
カラムナ型ストレージ、サーバーレス コンピューティング - トレースを数分ではなく数秒で見つけます。
エージェントの失敗を表面化するのは難しい
偽の「やりました」メッセージ、サイレントツールエラー、途中で諦めるなど、エージェントが失敗する可能性のあるものはたくさんあります。
単一の評価を作成する前に、ユーザーのフラストレーション、ツール呼び出しの最適化、異常を検出します。
1 日あたり 300 万件以上のエージェント トレース。ゼロサンプリング
クエリごと/ユーザーごとの料金はかかりません。
1 日に取り込まれるエージェント トレース データの総量。 GB TB 0 GB 500 GB 保持期間 (日) トレースを保持する期間。保持されたすべてのデータは同じ速度でクエリ可能であり、ウォーム/コールド層はありません。 1 1825 30d 90d 180d 1y 2y 5y ディメンション値 単価 月額コスト 取り込み 1200 GB/月 $ 0.30 /GB $360 保持 90 日 $0.001 /GB/月 $2 クエリ (UI、MCP、CLI、API) 無制限 $0 $0 合計 $362 /月 仕組み
エージェント スパン 100 万件あたり 10 ドルを達成した方法
Oodle では、あらゆる種類の可観測性信号に対応するオブジェクト ストレージを基盤とした、新しい高性能かつ低コストのデータベースを構築しました。
信頼できるエージェントを試すことはできません
建築業者は簡単になりました。数行のコードを書くだけで、機能するエージェントが完成します。しかし、実際に運用を開始すると、それが実際に機能していることをどうやって確認できるのでしょうか...
Oodle が大規模な可観測性をどのようにして高速に維持するか
「万が一に備えて」巨大な可観測性クラスターを存続させたことがある人なら、その雰囲気がわかるでしょう。支出の半分はアイドル状態の容量に費やされ、

残りの半分は希望に費やします...
Oodle エージェントの可観測性の仕組み
Oodle でエージェントの可観測性をセットアップして、サイレント エージェントの障害を表面化し、運用環境で初日からエージェントをデバッグする方法をご覧ください。
ユーザーが気づく前に問題を表面化します。
インサイトはエージェント トレース全体でサイレント障害を自動的に検出します。ルールを構成する必要はありません。
エラーが発生し、正常に回復できないエージェント。不完全なツール呼び出しチェーン、会話中の未処理の例外。
トレースにベースラインよりも大幅に時間がかかります。スタックしたループ、遅いツール呼び出し、および再試行の嵐を捕らえます。
感情スコアと品質スコアによって、否定的または主題から外れているとフラグが付けられた回答。悪い経験のパターンを特定します。
より安価なモデルでも同じ結果が得られる可能性があることを示します。単純なタスクで予算を消費しているオーバープロビジョニングされたコールを特定します。
キャッシュされる可能性のある同一のプロンプトが繰り返されました。トークンを無駄に消費し、待ち時間を追加する冗長な LLM 呼び出しにフラグを立てます。
会話が蛇行しすぎてしまう。プロンプト ループ、循環論法、および収束できないエージェントを検出します。
AI スタックのすべてのレイヤーを監視、デバッグ、最適化します。
エージェントのトランスクリプトを観察して詳細を調べる
システム プロンプト、ツールの呼び出し、マルチ エージェントのハンドオーバーなど、複数ターンにわたる会話を正確に読み取ります。
スパンごとに分類された入力、出力、推論トークン - 請求額に達する前に、暴走プロンプトや予算超過を特定します。
精度、関連性、安全性を確保するために評価パイプラインを構成します。すべての応答を自動的にスコアリングします。すぐに使用できる評価を使用するか、独自の評価を作成します。
Agentic スタック向けの AI を活用したインサイト
品質、パフォーマンス、効率に関する問題をすぐに表面化 - 初日からエージェントを改善します
Playground を使用してプロンプトを調整する
ライブ プレイグラウンドでプロンプトを反復処理します。再デプロイすることなく、モデルを交換し、パラメータを調整し、出力を比較します。
回帰チャネル

エージェントのメリット - データセットと実験
実稼働トレースまたは CSV インポートから評価データセットを構築します。任意のモデルに対して実験を実行し、結果を並べて比較します。
構成可能な価格設定によるコスト見積り
トレースごと、モデルごと、サービスごとの推定コスト。デフォルトで出荷される 50 以上のモデル価格定義 - OpenAI、Anthropic、Google、Mistral、Cohere など。 UI にカスタムまたは微調整されたモデルを追加します。
15 分以内にトレースの送信を開始します。無料利用枠が含まれており、クレジット カードは必要ありません。
LLM 呼び出しを計測するには別の SDK が必要ですか?
いいえ。Oodle は、Generative AI のセマンティック規約を使用して、標準の OpenTelemetry トレースを使用します。インストルメンテーション ライブラリがすでにこれらの属性を出力している場合は、OTLP エクスポータを Oodle に指定するだけです。どの LLM プロバイダーとモデルがサポートされていますか?
コードで呼び出すプロバイダー - OpenAI、Anthropic、Google Gemini、Mistral、Cohere、Azure OpenAI、AWS Bedrock、セルフホスト モデル。コスト見積もりには 50 以上のモデル価格定義が付属しており、カスタム モデルを追加できます。これはスタンドアロンの LLM 可観測性ツールとどう違うのですか?
スタンドアロン ツールは別のデータ サイロを作成します。 Oodle は、API、データベース、キャッシュ呼び出しと同じトレース内の LLM スパンを表示します。何かが壊れた場合、プロンプトと応答を個別に見るだけでなく、全体像がわかります。 Oodle はプロンプトのバージョン管理や評価を行いますか?
はい。 Oodle には、完全な LLMOps ワークフローが含まれています。バージョンとラベルのプロンプト、評価データセットの構築 (または CSV 経由でのインポート)、任意のモデルに対する実験の実行、組み込みまたはカスタムの評価器による結果のスコア付けなど、すべて運用トレースと並行して実行できます。 LLM トレースをサンプリングしますか?
いいえ。完全なプロンプトおよび応答コンテンツを含むトレースを 100% 保持します。エージェント主導のデバッグを体験してください。 80% 低いコストで 5 倍の速さでインシデントを解決します。
なぜ切り替えるのか?使用

クラウド マーケットプレイス経由でのクラウドのコミット: Oodle Pricing Company Blog Contact Technology Customers Docs Product Agent Debugging Agent Observability AI Cost Management Metrics Logs Traces APM Real User Monitoring Synthetics Why Switch Datadog OpenSearch Elasticsearch Grafana Labs ELK Stack すべての比較を見る Community Discord LinkedIn X (Twitter) YouTube リソース セキュリティ信頼性 トラスト センター ケース スタディ 大学 エージェント時代の可観測性アーキテクチャ — オブジェクト ストレージ - ネイティブ、サーバーレスRubrik、Amazon S3、DynamoDB、Snowflake のエンジニアによって構築された設計。
© 2026 ウードル。無断転載を禁じます。

## Original Extract

End-to-end visibility across LLM calls, APIs, and infra - so you can ship your AI with confidence, not hope. Agent observability at $10 per million spans.

Hi HN, we're Kiran and Vijay! Over the past two years, we have built a columnar storage engine for observability: logs, metrics, and traces. Today, it's exciting for us to show what we've built on top of that foundation: LLM Agent Observability. Given how non-deterministic agents are, storing all traces without sampling was critical for us. But these traces tend to be in the MBs, sometimes GBs - we needed to store them inexpensively. We also needed the queries and analyses to be fast. To meet both these goals, we store them in S3 in our own parquet-like file format, and query them using AWS Lambda. Since we process each span of every trace, instead of running LLM-based evals on each, we first analyze them using deterministic techniques. We detect tool failures, retries, loops, abnormal token usage, latency regressions, schema violations, sentiment, and other production signals. We've written more about the approach here: https://blog.oodle.ai/you-cant-sample-your-way-to-reliable-a... The combination of our own engine, no sampling, and deterministic processing before LLM-for-evals allows us to price at $10 per million traces, provide sub-second p99 query latency, and have healthy margins. Before building this, we used Langfuse for our own agent observability, which was 6x more expensive. Still super early, and rough around some edges, we would love your questions and feedback!

Agent Observability that Just Works at Scale - Oodle.ai | Oodle AI New Agent Observability is live. Agent traces from $10 per million spans. → Product AI Agent Debugging Agent Observability NEW AI Cost Management Observability Metrics Logs Traces APM Infrastructure Kubernetes Monitoring BYOC Technology Experience Real User Monitoring Synthetic Monitoring Migration 1-Click Migration See all products Customers Why Switch Elastic Grafana Datadog Elastic Cloud ELK stack OpenSearch See all comparisons → Pricing Overview Agent Observability Company Docs Resources Solutions For Series A Startups For Unicorns For Fortune Global 2000 Use Cases SaaS Platforms Observability Migration Guides Grafana Labs ELK Stack Security & Trust Reliability University Blog Conference Talks Case Studies TCO Calculator Discord Status Login US1 Region AP1 Region Try Playground Schedule Demo Product Customers Why Switch Pricing Company Docs Resources Login Try Playground Schedule Demo Shipping agents is easy. Observing agents is hard.
Searching Agent Traces is slow
Agent Traces have large metadata - prompts, tool calls, etc. - search results can take minutes to load.
Storing Agent Traces is expensive
Agents fail non-deterministically - so you can't afford to miss any trace - but you can't afford to keep them all either.
Surfacing agent failures is hard
Agents fail in so many ways - fake 'I did it' messages, silent tool errors, give up mid-chat...
Oodle: Agent Observability that “just works” at scale
Columnar storage, serverless compute - find the trace in seconds, not minutes.
Store 100% of traces and keep your budget too. Retain for months or years - you choose.
Detect user frustration, tool call optimizations, anomalies before you write a single eval
3M+ agent traces/day. Zero sampling.
Catch silent failures on Voice AI agents.
Shipping agents is easy. Observing agents is hard.
Storing Agent Traces is expensive
Agents fail non-deterministically - so you can't afford to miss any trace - but you can't afford to keep them all either.
Store 100% of traces and keep your budget too. Retain for months or years - you choose.
Searching Agent Traces is slow
Agent Traces have large metadata - prompts, tool calls, etc. - search results can take minutes to load.
Columnar storage, serverless compute - find the trace in seconds, not minutes.
Surfacing agent failures is hard
Fake 'I did it' messages, silent tool errors, gave up mid-way - so many ways agents can fail.
Detect user frustration, tool call optimizations, anomalies before you write a single eval
3M+ agent traces/day. Zero sampling
No per query / per-user charges.
Total volume of agent trace data ingested per day. GB TB 0 GB 500 GB Retention (days) How long to keep your traces. All retained data is queryable at the same speed — no warm/cold tiers. 1 1825 30d 90d 180d 1y 2y 5y Dimension Value Unit Price Monthly Cost Ingestion 1200 GB/mo $ 0.30 /GB $360 Retention 90 days $0.001 /GB/mo $2 Queries (UI, MCP, CLI, API) unlimited $0 $0 Total $362 /mo How It Works
How We Achieved $10/Million Agent Spans
At Oodle we built a new high performance, low cost database backed by object storage for each kind of observability signal...
You Can't Sample Your Way to Reliable Agents
Building agents has become easy. A few lines of code and you have a working agent. But once it's in production, how do you know it's actually working...
How Oodle Keeps Observability Fast at Scale
If you've ever kept a giant observability cluster alive “just in case,” you know the vibe: half your spend goes to idle capacity, and the other half goes to hoping ...
How Oodle Agent Observability Works
See how you can setup Agent Observability on Oodle to surface silent agent failures and debug your agents in production from day 1.
Surface issues before users notice.
Insights automatically detect silent failures across your agent traces - no rules to configure.
Agents that hit errors and fail to recover gracefully. Incomplete tool call chains, unhandled exceptions mid-conversation.
Traces taking significantly longer than baseline. Catches stuck loops, slow tool calls, and retry storms.
Responses flagged by sentiment and quality scoring as negative or off-topic. Identifies patterns in poor experiences.
Traces where a cheaper model could produce the same result. Spots over-provisioned calls burning budget on simple tasks.
Repeated identical prompts that could be cached. Flags redundant LLM calls wasting tokens and adding latency.
Conversations that spiral into too many turns. Detects prompt loops, circular reasoning, and agents that can't converge.
Monitor, debug, and optimize every layer of your AI stack.
Observe Agent Transcripts for Deep Dives
Read the exact multi-turn conversation: system prompt, tool calls, multi-agent handovers.
Input, output, and reasoning tokens broken down per span - spot runaway prompts and budget overruns before they hit your bill.
Configure eval pipelines for accuracy, relevance, and safety. Score every response automatically - use our out-of-the-box evals or write your own.
AI-powered Insights for your Agentic stack
Surface issues on Quality, Performance, Efficiency out-of-the-box - improve your agents from day 1
Refine Your Prompts with Playground
Iterate on prompts in a live playground. Swap models, tweak parameters, and compare outputs - without redeploying.
Regression Checks for Your Agents - Datasets & Experiments
Build eval datasets from production traces or CSV imports. Run experiments against any model and compare results side-by-side.
Cost Estimation with Configurable Pricing
Estimated cost per trace, per model, per service. 50+ model pricing definitions shipped by default - OpenAI, Anthropic, Google, Mistral, Cohere, and more. Add custom or fine-tuned models in the UI.
Start sending traces in under 15 minutes. Free tier included - no credit card required.
Do I need a separate SDK to instrument my LLM calls?
No. Oodle consumes standard OpenTelemetry traces using the Generative AI semantic conventions . If your instrumentation library already emits these attributes, you just point your OTLP exporter at Oodle. Which LLM providers and models are supported?
Any provider that your code calls - OpenAI, Anthropic, Google Gemini, Mistral, Cohere, Azure OpenAI, AWS Bedrock, and self-hosted models. Cost estimation ships with 50+ model pricing definitions, and you can add custom models. How is this different from standalone LLM observability tools?
Standalone tools create another data silo. Oodle shows your LLM spans inside the same traces as your API, database, and cache calls. When something breaks, you see the full picture - not just the prompt and response in isolation. Does Oodle do prompt versioning or evals?
Yes. Oodle includes a full LLMOps workflow: version and label prompts, build eval datasets (or import via CSV), run experiments against any model, and score results with built-in or custom evaluators - all alongside your production traces. Do you sample LLM traces?
No. Retain 100% of your traces with full prompt and response content. Experience agent-driven debugging. Resolve incidents 5x faster at 80% lower cost.
Why Switch? Use your cloud commits via cloud marketplace: Oodle Pricing Company Blog Contact Technology Customers Docs Product Agent Debugging Agent Observability AI Cost Management Metrics Logs Traces APM Real User Monitoring Synthetics Why Switch Datadog OpenSearch Elasticsearch Grafana Labs ELK Stack See all comparisons Community Discord LinkedIn X (Twitter) YouTube Resources Security Reliability Trust Center Case Studies University The observability architecture for the agent era — object storage–native, serverless by design, built by the engineers behind Rubrik, Amazon S3, DynamoDB, and Snowflake.
© 2026 Oodle. All rights reserved.

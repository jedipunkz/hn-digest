---
source: "https://production-ai-systems-beta.vercel.app/"
hn_url: "https://news.ycombinator.com/item?id=49112658"
title: "Production AI Systems – 34 chapters, 1,026 runnable assertions"
article_title: "Building Production AI Systems"
author: "mandeep419singh"
captured_at: "2026-07-30T17:16:23Z"
capture_tool: "hn-digest"
hn_id: 49112658
score: 1
comments: 1
posted_at: "2026-07-30T16:59:11Z"
tags:
  - hacker-news
  - translated
---

# Production AI Systems – 34 chapters, 1,026 runnable assertions

- HN: [49112658](https://news.ycombinator.com/item?id=49112658)
- Source: [production-ai-systems-beta.vercel.app](https://production-ai-systems-beta.vercel.app/)
- Score: 1
- Comments: 1
- Posted: 2026-07-30T16:59:11Z

## Translation

タイトル: プロダクション AI システム – 34 章、1,026 個の実行可能なアサーション
記事のタイトル: 生産 AI システムの構築
概要: 本番用AIシステムの設計・構築・運用に関する技術書

記事本文:
生産AIシステムの構築
コンテンツにスキップ
生産AIシステムの構築
ホーム
検索を初期化しています
MandeepSinghthakur/production-ai-systems
ホーム
パート II - データとイベントのインフラストラクチャ
生産AIシステムの構築
MandeepSinghthakur/production-ai-systems
ホーム
ホーム
目次
この本について
この本の違い
パート I - 分散システム
パート I - 分散システム
1. AI ワークロードのための分散システム
2. ステートレス API とストリーミング API のスケーリング
3. 長期間存続する接続の負荷分散
パート II - データとイベントのインフラストラクチャ
パート II - データとイベントのインフラストラクチャ
6. Kafka の内部構造
7. イベント駆動型のアーキテクチャ パターン
8. 送信ボックス、サーガ、および 1 回限り
9. 可観測性と OpenTelemetry
パート III - LLM の基礎
パート III - LLM の基礎
10. トランスフォーマーが実際にリクエストを処理する方法
11. トークン化とコンテキスト ウィンドウ
13. ストリーミングとトークンエコノミクス
パート IV - 取得
パート IV - 取得
14. ドキュメントの取り込みと OCR
16. ベクトルデータベースとハイブリッド検索
17. リランキングと検索評価
パート V - AI プラットフォーム
パート V - AI プラットフォーム
18. LLM ゲートウェイ
19. マルチプロバイダーのルーティングとフェイルオーバー
22. AI セキュリティと即時導入
23. コスト管理とキャパシティプランニング
パート VI - エージェント AI
パート VI - エージェント AI
24. ツール呼び出しとツールセキュリティ
25. 計画、反映、および対応
26. モデルコンテキストプロトコル
27. マルチエージェント システムとオーケストレーション
パート VII - システム設計
パート VII - システム設計
28. デザイン - 大規模な会話アシスタント
29. デザイン - 規制された領域における AI アシスタント
31. 設計 - マルチテナント エージェント プラットフォーム
パート VIII - スタッフエンジニアリング
パート VIII - スタッフエンジニアリング
32. アーキテクチャのレビューと RFC
33. AI システムのインシデント管理
34. 技術戦略と影響力
付録

エス
付録
付録 A - 用語集
付録 C - 注釈付き参考文献
この本の違い
生産 AI システムの構築 ¶
本番用AIシステムの設計・構築・運用に関する技術書。
この本は、ソフトウェア エンジニアに実稼働 AI システムの構築方法を教えます。デモではありません。ノートではありません。実際のユーザーにサービスを提供し、実際のお金を処理し、午前 3 時に休憩するシステム。
34 章が 8 つのパートに分かれています。
分散システム基盤 - 1 つのリクエストに 40 秒かかると何が変わるか
データおよびイベント インフラストラクチャ - Kafka、送信ボックス パターン、可観測性
LLM の基礎 - トランスフォーマー、トークン化、埋め込み、ストリーミング
取得 - ドキュメントの取り込み、チャンキング、ベクトル検索、再ランキング
AI プラットフォーム - ゲートウェイ、ルーティング、メモリ、評価、セキュリティ、コスト管理
Agentic AI - ツール呼び出し、計画、MCP、マルチエージェント システム
システム設計 - 容量見積もりを含む完全な設計ウォークスルー
スタッフ エンジニアリング - アーキテクチャのレビュー、インシデント管理、技術戦略
この本の違い ¶
あらゆる主張が検証されます。この本では、再試行の嵐によって負荷が 2 倍に増幅されると述べていますが、それを実証する実行可能なコードがあります。例は主張であり、説明ではありません。散文内の数値がコード内の数値と一致しない場合、ビルドは失敗します。
面接対策も統合されています。各章はインタビューの質問とスタッフレベルの回答で終わります。章を終える頃には、そのトピックに関する質問に答えることができるようになります。
壊れたものごとに整理されています。従来の書籍はテクノロジーごとに整理されています。この本は、取得パイプラインがガベージを返した場合、プロバイダーが停止した場合、コストが予算を超えた場合に何が起こるかという障害モード別にまとめられています。
各章には実行可能なコードがあります。 API キーは必要ありません。ドッカーはありません。 Node.js 22.6以降のみ。
# クローン

リポジトリ
git clone https://github.com/MandeepSinghthakur/production-ai-systems
CD 制作-AI-システム
# 任意の章のラボを実行する
ノードの例/ch18-llm-gateway/scripts/lab.mjs
各ラボは、実行中にアサーションを出力します。例は真実の源です。
まえがき - この本は誰に向けたものですか
この本の読み方 - 3 つの読み方

## Original Extract

A technical book on designing, building, and operating production AI systems

Building Production AI Systems
Skip to content
Building Production AI Systems
Home
Initializing search
MandeepSinghthakur/production-ai-systems
Home
Part II - Data and Event Infrastructure
Building Production AI Systems
MandeepSinghthakur/production-ai-systems
Home
Home
Table of contents
About This Book
What Makes This Book Different
Part I - Distributed Systems
Part I - Distributed Systems
1. Distributed Systems for AI Workloads
2. Scaling Stateless and Streaming APIs
3. Load Balancing Long-Lived Connections
Part II - Data and Event Infrastructure
Part II - Data and Event Infrastructure
6. Kafka Internals
7. Event-Driven Architecture Patterns
8. Outbox, Saga, and Exactly-Once
9. Observability and OpenTelemetry
Part III - LLM Fundamentals
Part III - LLM Fundamentals
10. How Transformers Actually Serve Requests
11. Tokenization and Context Windows
13. Streaming and Token Economics
Part IV - Retrieval
Part IV - Retrieval
14. Document Ingestion and OCR
16. Vector Databases and Hybrid Search
17. Re-ranking and Retrieval Evaluation
Part V - The AI Platform
Part V - The AI Platform
18. The LLM Gateway
19. Multi-Provider Routing and Failover
22. AI Security and Prompt Injection
23. Cost Control and Capacity Planning
Part VI - Agentic AI
Part VI - Agentic AI
24. Tool Calling and Tool Security
25. Planning, Reflection, and ReAct
26. The Model Context Protocol
27. Multi-Agent Systems and Orchestration
Part VII - System Design
Part VII - System Design
28. Design - A Conversational Assistant at Scale
29. Design - An AI Assistant in a Regulated Domain
31. Design - A Multi-Tenant Agent Platform
Part VIII - Staff Engineering
Part VIII - Staff Engineering
32. Architecture Reviews and RFCs
33. Incident Management for AI Systems
34. Technical Strategy and Influence
Appendices
Appendices
Appendix A - Glossary
Appendix C - Annotated Bibliography
What Makes This Book Different
Building Production AI Systems ¶
A technical book on designing, building, and operating production AI systems.
This book teaches software engineers how to build production AI systems. Not demos. Not notebooks. Systems that serve real users, handle real money, and break at 3 AM.
34 chapters organized into eight parts:
Distributed Systems Foundations - What changes when a single request takes 40 seconds
Data and Event Infrastructure - Kafka, outbox patterns, observability
LLM Fundamentals - Transformers, tokenization, embeddings, streaming
Retrieval - Document ingestion, chunking, vector search, re-ranking
The AI Platform - Gateways, routing, memory, evaluation, security, cost control
Agentic AI - Tool calling, planning, MCP, multi-agent systems
System Design - Complete design walkthroughs with capacity estimates
Staff Engineering - Architecture reviews, incident management, technical strategy
What Makes This Book Different ¶
Every claim is tested. When this book says that retry storms amplify load by 2x, there is runnable code that demonstrates it. The examples are assertions, not illustrations. If the numbers in prose do not match the numbers in code, the build fails.
Interview preparation is integrated. Each chapter ends with interview questions and staff-level answers. By the time you finish a chapter, you can answer questions about its topic.
Organized by what breaks. Traditional books organize by technology. This book organizes by failure mode: what happens when your retrieval pipeline returns garbage, when your provider has an outage, when your costs exceed your budget.
Every chapter has runnable code. No API keys required. No Docker. Just Node.js 22.6+.
# Clone the repository
git clone https://github.com/MandeepSinghthakur/production-ai-systems
cd production-ai-systems
# Run any chapter's lab
node examples/ch18-llm-gateway/scripts/lab.mjs
Each lab prints assertions as it runs. The examples are the source of truth.
Preface - Who this book is for
How to Read This Book - Three reading paths

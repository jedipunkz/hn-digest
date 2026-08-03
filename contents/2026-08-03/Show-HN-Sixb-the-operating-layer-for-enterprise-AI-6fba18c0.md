---
source: "https://sixb.ai/news/sixb-framework-launch"
hn_url: "https://news.ycombinator.com/item?id=49154981"
title: "Show HN: Sixb – the operating layer for enterprise AI"
article_title: "Sixb: the operating layer for enterprise AI"
author: "ademattos"
captured_at: "2026-08-03T12:49:52Z"
capture_tool: "hn-digest"
hn_id: 49154981
score: 2
comments: 0
posted_at: "2026-08-03T12:41:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sixb – the operating layer for enterprise AI

- HN: [49154981](https://news.ycombinator.com/item?id=49154981)
- Source: [sixb.ai](https://sixb.ai/news/sixb-framework-launch)
- Score: 2
- Comments: 0
- Posted: 2026-08-03T12:41:04Z

## Translation

タイトル: ショー HN: Sixb – エンタープライズ AI のオペレーティング層
記事のタイトル: Sixb: エンタープライズ AI のオペレーティング層
説明: Sixb は、人間と AI エージェントが使用する運用ソフトウェアを構築するために Bun 上に構築されたオープンソースの TypeScript フレームワークです。
HN テキスト: こんにちは、HN!ビジネスの運用層をモデル化できるフレームワークをオープンソース化しました。 1. 今日の問題: ほぼすべての企業が AI アシスタント (ChatGPT、Claude、または Gemini) を使用しています。しかし、それはアシスタントのままです。そして企業は、たとえ一部が機密事項であっても、より多くの情報を提供するようになっています。なぜなら、より良いコンテキストがより良い答えを生み出すからです。モデルには会社の業務が含まれていてはなりません。それはそれらに接続されている 1 つのコンポーネントである必要があります。 2. 私たちの論文 AI はインターフェイスと接着コードを安価にします。難しいのは、その背後にあるビジネス コンテキスト、ルール、権限、ワークフローをモデル化することです。 MCP はツールへのアクセスを解決しますが、ビジネス コンテキスト、権限、ガバナンスは解決しません。有用な AI 実装を既製で入手することはできません。各企業の運営方法は異なります。 3. Sixb でできること - データの接続: CRM、ERP、ファイル システム、会議記録...
- ビジネスをモデル化します: 顧客、プロジェクト、設備、見積書、契約、およびそれらの間の関係
- ルールを定義します: 顧客は 2 日以内に見積もりを受け取る必要があります
- 権限を定義する: 各ユーザーまたは AI エージェントが表示できる内容と実行できるアクションを決定します。
- ワークフローの定義: ビジネス イベントに反応し、人間の承認を要求し、ソース システムでアクションを実行します。
- カスタム アプリケーションの構築: チームと AI エージェントに同じ運用コンテキストと制御されたアクションを提供します。ぜひ皆さんに使っていただき、フィードバックや貢献をしていただきたいと思っています。

記事本文:
Sixb: エンタープライズ AI のオペレーティング層 ニュース ドキュメント GitHub ↗ EN / FR ご相談ください 30 分の通話を予約してください
Quentin Nippert フランス ← Sixb を立ち上げる: エンタープライズ AI のオペレーティング層
本日、人間と AI エージェントが使用する運用ソフトウェアを構築するためのオープンソース TypeScript フレームワークである Sixb をリリースします。
アンソニー・デマットスとクエンティン・ニッパート
私たちは過去数年間、建築業界に携わり、建物の効率と運用を向上させるデジタル ツインの作成に取り組んできました。
この経験から、デジタル ツイン自体は製品ではないという明確な教訓が得られました。その価値は、実際の問題を解決するユースケースから生まれます。新しいニーズが出現すると、それぞれのユースケースが運用モデルに追加されます。これらの小さな成果が合わさって、実際の運用ニーズに基づいた接続システムが形成されます。
システムは一度に 1 つのユースケースで構築します。
現在、ほぼすべての企業が AI アシスタント (ChatGPT、Claude、または Gemini) を使用しています。
プロンプトが表示されると、多くの場合、不完全なコンテキストで質問に答えます。ビジネスで何が変更されたのか、どのルールが適用されるのか、ユーザーに何を表示できるのか、どのアクションに承認が必要なのかはわかりません。
また、より良いコンテキストがより良い答えを生み出すため、企業は大量のビジネス データを AI プロバイダーに送信します。機密データ、権限、ガバナンスは遅すぎると考えられることがよくあります。
モデルには会社の業務が含まれていてはなりません。それはそれらに接続されている 1 つのコンポーネントである必要があります。
AI により、インターフェイスの生成とグルー コードが安価になりました。
難しいのは、別のフォーム、ダッシュボード、または独立した内部ツールを作成する必要がなくなったことです。ビジネス コンテキスト、ルール、権限、その背後にあるワークフローをモデル化しています。
MCP はツールへのアクセスを解決しますが、ビジネス コンテキスト、権限、ガバナンスは解決しません。
AI ツールを使用してアプリを生成すると、アイソレーションを解決できる

必要性はすぐに満たされますが、多くの場合、独自のデータ、ロジック、権限を持つ別の切断されたアプリケーションが作成されます。
有用な AI 実装を既製で入手することはできません。各企業の運営方法は異なります。
Sixb を使用すると、組織の運用レイヤーをモデル化できるため、チームと AI エージェントが同じ環境で作業できるようになります。
CRM、ERP、ファイル システム、会議記録、機器、または内部 API。
顧客、プロジェクト、設備、見積書、契約、およびそれらの間の関係。
顧客は 2 日以内に見積もりを受け取る必要があります。修理が保証されるかどうかは契約によって決まります。
各ユーザーまたは AI エージェントが表示できる内容と実行できるアクションを決定します。
ビジネス イベントに反応し、人間の承認を要求し、ソース システムでアクションを実行します。
チームと AI エージェントに同じ運用コンテキストと制御されたアクションを提供します。
企業における AI は質問に答えるだけではありません。チームの一員として安全に活動できるはずです。
このリポジトリには、Sixb だけで構築された架空の商用サービス会社である Northline Operations が含まれています。
機器アラームを受信すると、Sixb は次のことができます。
02 顧客、施設、機器、有効な契約を取得します。
03 カバレッジと応答時間のルールを適用します。
04 対応可能な資格のある技術者を推薦します。
05 ディスパッチャに割り当ての承認を依頼します。
06 フィールドサービス システムで作業指示書を作成します。
07 診断、見積もり、修理、回復を追跡します。
アプリケーション、ワークフロー、ルール、アクション、ソース システムはすべて、決定的なサンプル データを使用してローカルで実行されます。
ビジネス ドメインを記述するオントロジーを定義します。
オントロジーは TypeScript タイプのように機能し、ビジネス オブジェクトが尊重する必要があるコントラクトを定義します。
実行できるアクションも定義します。
1 "@sixb/blob-local" から { LocalBlobStorage } をインポートします 2

import { createSixb , InMemoryBroker, InMemoryQueues } from "@sixb/core" 3 import { DuckLakeStorage } from "@sixb/ducklake" 4 import { SqliteStorage } from "@sixb/sqlite" 5 6 import const sixb = createSixb ({ 7 id: "service-operations" , 8 Broker: new InMemoryBroker()、9 ストレージ: new SqliteStorage({ path: ".sixb" })、10 lakeStorage: new DuckLakeStorage({ 11 カタログ: { type: "duckdb" 、パス: ".sixb/lake/catalog.ducklake" }、12 dataPath: ".sixb/lake/data" 、13 })、14 blobStorage: new LocalBlobStorage({basePath: ".sixb/blobs" })、15 キュー: new InMemoryQueues(), 16 }) sixb.config.ts コピー 02 接続と変換
コネクタはソース システムと通信します。
同期は、S3 互換のオブジェクト ストレージに保存されているデータセットに生データを取り込みます。
パイプラインは、SQL または TypeScript を使用してこれらのデータセットをクリーンアップおよび変換します。
プロジェクションは、結果のデータを、オントロジーによって定義された型付きオブジェクト、関係、およびテレメトリに変換します。
ルールはビジネスの運用状態を評価します。
権限は、ユーザーと AI エージェントが何を表示し、実行できるかを制御します。
ワークフローは、アクション、イベント、スケジュール、人間の介入を調整します。
基盤が整ったら、業務に実際に必要なソフトウェアを構築できます。
複数のシステムからのデータを組み合わせたダッシュボード
AI アクションのための人間による監視インターフェイス
特定のチームに合わせた運用アプリケーション
見てはいけない情報にアクセスせずにビジネスを理解する AI アシスタント
ユーザーのプロンプトを待たずに反応するイベント駆動型エージェント
完全なロードマップと技術的制限はリポジトリで入手できます。
Sixbはまだ早いです。フレームワークが進化するにつれて破壊的な変更が発生することが予想されます。
私たちは以下のことに意欲的な開発者を求めています。
独自の運用アプリケーションを構築する
何が欠けている、または過剰に設計されていると感じるかを教えてください
貢献する

コネクタ、ストレージプロバイダー、例、およびコア機能
Sixbが始まりです。あなたがそれを使って何を構築するのかをぜひ見てみたいと思います。
Quentin Nippert フランス LinkedIn 初めてのワークフローをマップします。
価値の高い 1 つのプロセスから始めます。 Sixb はそれを中心にシステムを構築するため、次回の展開はすべて、すでに導入されているものから始まります。
ビジネスが実際にどのように機能するかに基づいて構築されたソフトウェア。
© 2026 シックスビー.無断転載を禁じます。

## Original Extract

Sixb is an open-source TypeScript framework built on Bun for building operational software used by humans and AI agents.

Hi HN! We open sourced a framework to be able to model the operational layer of a business. 1. The problem today: Almost every company uses an AI assistant: ChatGPT, Claude, or Gemini. But it remains an assistant. And the companies are giving more and more information, even if some are sensitive, because better context usually produces better answers. The model should not contain the company’s operations. It should be one component connected to them. 2. Our thesis AI is making interface and glue code cheap. The difficult part is modeling the business context, rules, permissions, and workflows behind them. MCP solves access to tools, but it does not solve business context, permissions, or governance. You cannot get a useful AI implementation off the shelf. Every company operates differently. 3. What you can do with Sixb - Connect your data: CRM, ERP, file systems, meeting transcripts...
- Model your business: Customers, projects, equipment, quotes, contracts, and the relationships between them
- Define your rules: A customer must receive a quote within two days
- Define your permissions: Decide what each user or AI agent can see and which actions they can perform.
- Define workflows: React to business events, request human approval, and execute actions in source systems.
- Build custom applications: Give teams and AI agents the same operational context and controlled actions. We would love to see people using it, giving us feedback and contributing!

Sixb: the operating layer for enterprise AI About News Docs GitHub ↗ EN / FR Talk to us Book a 30-min call
Quentin Nippert France ← Launch Sixb: the operating layer for enterprise AI
Today, we are releasing Sixb, an open-source TypeScript framework for building operational software used by humans and AI agents.
Anthony DeMattos and Quentin Nippert
We spent the last few years in the building industry, creating digital twins to improve building efficiency and operations.
That experience left us with a clear lesson: the digital twin itself is not the product. The value comes from use cases that solve real problems. As new needs emerge, each use case adds to the operational model. Together, these small wins compound into a connected system grounded in real operational needs.
Build the system one use case at a time.
Today, almost every company uses an AI assistant: ChatGPT, Claude, or Gemini.
It answers questions when prompted, often with incomplete context. It does not know what changed in the business, which rules apply, what users are allowed to see, or which actions require approval.
Companies also send large amounts of business data to AI providers because better context usually produces better answers. Sensitive data, permissions, and governance are often considered too late.
The model should not contain the company’s operations. It should be one component connected to them.
AI is making interface generation and glue code cheaper.
The difficult part is no longer creating another form, dashboard, or isolated internal tool. It is modeling the business context, rules, permissions, and workflows behind them.
MCP solves access to tools, but it does not solve business context, permissions, or governance.
Generating an app with an AI tool can solve an isolated need quickly, but it often creates another disconnected application with its own data, logic, and permissions.
You cannot get a useful AI implementation off the shelf. Every company operates differently.
Sixb lets you model the operational layer of an organization so teams and AI agents can work in the same environment.
CRM, ERP, file systems, meeting transcripts, equipment, or internal APIs.
Customers, projects, equipment, quotes, contracts, and the relationships between them.
A customer must receive a quote within two days. A contract determines whether a repair is covered.
Decide what each user or AI agent can see and which actions they can perform.
React to business events, request human approval, and execute actions in source systems.
Give teams and AI agents the same operational context and controlled actions.
AI in the enterprise should not only answer questions. It should be able to operate safely as part of the team.
The repository includes Northline Operations , a fictional commercial service company built entirely with Sixb.
When an equipment alarm is received, Sixb can:
02 Retrieve the customer, facility, equipment, and active contract.
03 Apply coverage and response-time rules.
04 Recommend an available and qualified technician.
05 Ask a dispatcher to approve the assignment.
06 Create the work order in the field-service system.
07 Track the diagnosis, quote, repair, and recovery.
The application, workflows, rules, actions, and source systems all run locally with deterministic example data.
You define an ontology describing the business domain:
An ontology works like TypeScript types: it defines the contracts that your business objects must respect.
You also define the actions that can be executed:
1 import { LocalBlobStorage } from "@sixb/blob-local" 2 import { createSixb , InMemoryBroker, InMemoryQueues } from "@sixb/core" 3 import { DuckLakeStorage } from "@sixb/ducklake" 4 import { SqliteStorage } from "@sixb/sqlite" 5 6 export const sixb = createSixb ({ 7 id: "service-operations" , 8 broker: new InMemoryBroker(), 9 storage: new SqliteStorage({ path: ".sixb" }), 10 lakeStorage: new DuckLakeStorage({ 11 catalog: { type: "duckdb" , path: ".sixb/lake/catalog.ducklake" }, 12 dataPath: ".sixb/lake/data" , 13 }), 14 blobStorage: new LocalBlobStorage({ basePath: ".sixb/blobs" }), 15 queues: new InMemoryQueues(), 16 }) sixb.config.ts Copy 02 Connect and transform
Connectors communicate with source systems.
Syncs ingest raw data into datasets stored in S3-compatible object storage.
Pipelines clean and transform those datasets using SQL or TypeScript.
Projections turn the resulting data into typed objects, relationships, and telemetry defined by your ontology.
Rules evaluate the operational state of the business.
Permissions control what users and AI agents can see and do.
Workflows coordinate actions, events, schedules, and human interventions.
Once the foundations are in place, you can build the software your operations actually need:
A dashboard combining data from multiple systems
A human-supervision interface for AI actions
An operational application adapted to a specific team
An AI assistant that understands the business without accessing information it should not see
Event-driven agents that react without waiting for a user prompt
The complete roadmap and technical limitations are available in the repository.
Sixb is still early. Expect breaking changes as the framework evolves.
We are looking for developers willing to:
Build their own operational application
Tell us what feels missing or over-engineered
Contribute connectors, storage providers, examples, and core features
Sixb is at the beginning. We would love to see what you build with it.
Quentin Nippert France LinkedIn Map your first workflow.
Start with one high-value process. Sixb builds the system around it, so every next deployment starts with more already in place.
Software built around how your business actually works.
© 2026 Sixb. All rights reserved.

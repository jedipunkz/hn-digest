---
source: "https://aws.amazon.com/blogs/machine-learning/context-intelligence-for-your-data-and-ai-agents-at-scale/"
hn_url: "https://news.ycombinator.com/item?id=48579171"
title: "Context intelligence for your data and AI agents at scale"
article_title: "Context intelligence for your data and AI agents at scale | Artificial Intelligence"
author: "champagnepapi"
captured_at: "2026-06-18T01:01:18Z"
capture_tool: "hn-digest"
hn_id: 48579171
score: 1
comments: 0
posted_at: "2026-06-18T00:56:34Z"
tags:
  - hacker-news
  - translated
---

# Context intelligence for your data and AI agents at scale

- HN: [48579171](https://news.ycombinator.com/item?id=48579171)
- Source: [aws.amazon.com](https://aws.amazon.com/blogs/machine-learning/context-intelligence-for-your-data-and-ai-agents-at-scale/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T00:56:34Z

## Translation

タイトル: 大規模なデータと AI エージェントのためのコンテキスト インテリジェンス
記事のタイトル: 大規模なデータと AI エージェントのためのコンテキスト インテリジェンス |人工知能
説明: エージェントのインテリジェント性は、エージェントが推論できるコンテキストと同程度に限られます。現在、そのコンテキストはデータ レイク、データ ウェアハウス、レイクハウス、データベース、ストリームに散在しており、また文書化されたことのない組織的な知識の中に散在しています。 AI エージェントによる決定を信頼したいと考えていますが、
[切り捨てられた]

記事本文:
大規模なデータと AI エージェントのためのコンテキスト インテリジェンス
エージェントのインテリジェンスは、エージェントが推論できるコンテキストと同程度に限られます。現在、そのコンテキストはデータ レイク、データ ウェアハウス、レイクハウス、データベース、ストリームに散在しており、また文書化されたことのない組織的な知識の中に散在しています。 AI エージェントが下した決定を信頼したいと考えていますが、エージェントがコンテキストを理解するまではそれができません。エージェントが信頼できる決定を下すために必要なコンテキストに安全にアクセスできる方法を提供すると、何が可能になるかを想像してみてください。
これが、AWS Summit New York City で、データと AI エージェントに大規模なコンテキスト インテリジェンスを提供する一連のイノベーションを発表する理由です。
本日の基調講演では、既存のデータ間の関係をナレッジグラフに自動的にマッピングし、組織内の AI エージェントが実行時に管理されたデータ関係、ビジネスルール、ドメインナレッジにアクセスできるようにエージェント検索を提供する新しいサービスである AWS Context を紹介しました。データ スチュワードとキュレーターは、直感的なコンソール エクスペリエンスを通じてグラフを管理し、推測された関係をレビューして実稼働環境にプロモートし、ビジネス定義や使用ルールなどのドメイン固有の知識を付加します。
AWS Context は、Amazon Quick を強化するのと同じナレッジ グラフ テクノロジーを拡張します。Amazon Quick では、数十万のユーザーが、データセット、ダッシュボード、メタデータをカタログ化する本番ナレッジ グラフと毎日対話し、使用パターンから学習してあらゆる対話をよりスマートにします。このグラフでは、すでに 1 日に何百万ものリクエストが処理されています。 AWS Context を使用すると、個人のナレッジ グラフであったものを組織のナレッジ グラフ、つまり組織内のエージェントやアプリケーションが利用できる共有の管理されたコンテキスト レイヤーに拡張します。既存のAmazonクイック利用

すぐに恩恵を受けられます。 AWS コンテキストが有効になると、Quick のエージェントは、システム間の関係、ビジネス ルール、単一ユーザーの個人グラフが提供できる範囲を超えた厳選されたコンテキストなど、より広範なエンタープライズ ナレッジ グラフにアクセスできるようになります。 AWS Glue データカタログ、Amazon SageMaker Unified Studio、および AWS Lake Formation はナレッジグラフと統合されているため、チームはビジネスルールと権限でナレッジグラフを管理し、AI 支援によって自動的に、または手動キュレーションを通じて明示的に新しいコンテキストを追加できます。
コンテキストレイヤーの主要な要素は、Apache Iceberg 形式で Amazon S3 に公開されるため、顧客は任意の Iceberg 準拠ツールを自由に使用して、メタデータを使用し、オープンスタンダードに基づいて AWS Context に対して構築できます。プロビジョニングするインフラストラクチャや構築する取得パイプラインはなく、顧客は AWS マネジメント コンソールで数回クリックするだけでエージェントのコンテキストの収集とキュレーションを開始できます。
その背後にある機能を詳しく見てみましょう。
エージェントの働き方から学ぶコンテキスト
AWS Context は、エージェントが使用すればするほど賢くなります。エージェントがグラフをクエリすると、どのソースが正しい結果を生成するか、エージェントが依存する結合パス、およびどの厳選されたルールが適用されるかを観察します。実際の使用状況によってソースをランク付けし、学習した内容を組織全体で共有するため、あるエージェントが正しい結合パスを発見したり、スキーマのあいまいさを解決したりすると、人間がグラフを再キュレートすることなく、他のエージェントがそれを取得します。
AWS Context は、構造化ソースおよび非構造化ソースからのすべての主要なメタデータを Amazon S3 テーブルの Apache Iceberg 形式で公開するため、Amazon Athena、Amazon Redshift、Apache Spark、または任意の Iceberg 互換エンジンを使用してコンテキストをクエリし、その上にダウンストリーム システムを構築し、監査することができます。

それを使用するか、移行します。
AWS Context はサードパーティのカタログに接続するように設計されているため、AWS 以外のシステムからのコンテキストを同じグラフに取り込むことができます。エージェントは、Amazon Bedrock AgentCore 上に構築されているか、Amazon EKS 上にデプロイされているか、MCP 互換フレームワーク上で実行されているかに関係なく、エージェント検索 API および MCP ツールを通じてクエリを実行します。コンテキストはクエリ可能であり、Apache Iceberg 形式で移植可能であり、完全にあなたのものになります。
ID を認識し、デフォルトで管理される
エージェントを実稼働環境に導入すると、ガバナンスに関する問題が生じます。つまり、どのようなデータにアクセスできるのか、誰の権限で何にアクセスしたのかを正確に示すことができるのかということです。 AWS Context は、すべてのクエリに ID を認識させることで両方に答えます。各呼び出しは、呼び出し元ユーザーの IAM および Lake Formation 権限を継承するように設計されているため、エージェントは、その ID がアクセスを許可されている関係のみを表示および横断できます。アクセスは ID を介して実行されるため、すべての対話は監査可能です。セキュリティ チームとコンプライアンス チームは、すでに依存しているのと同じ制御を使用して、エージェントがどの権限で何にアクセスしたかを確認できます。
AWS Glue データカタログ ビジネスコンテキストおよびセマンティック検索 (プレビュー)
本日、AWS Glue データカタログのビジネスコンテキストとセマンティック検索のプレビューも発表し、人間や AI エージェントがデータを発見して理解することを容易にするコンテキストとツールを提供します。顧客は、ビジネスの説明、用語集の用語、カスタム メタデータを使用して Glue のテーブル、ビュー、列 (S3 テーブルでサポートされているテーブルを含む) を強化し、カタログの外に保存されている追加のデータ コンテキストを提供するスキル アセットに関連付けることができるようになりました。 Glue データ カタログの技術メタデータとともにビジネス コンテキストのインデックスが作成されるため、顧客は新しい Glue Search API を使用して、ビジネス上の意味によってデータをより迅速に検索でき、AI エージェントは

文脈を推測するのではなく、信頼できる定義で推論を丸めます。
また、Glue Data Catalog でスキル アセットのプレビューを提供できることを嬉しく思います。データ プロデューサーは、S3、Git リポジトリ、Wiki などの任意の場所でホストされているファイル (AI スキル、ガイド マークダウン ファイル、チーム ランブックなど) への URI を参照する新しいアセット タイプであるスキル アセットを作成できるようになりました。スキルアセットをデータアセットに関連付けると、エージェントは追加のコンテキストと指示を取得して、特定のデータを操作するために、一度に 1 つのプロンプトをすべてのエージェントに再教育することなく、段階的に取得できるようになります。たとえば、スキル URI の場所は、粒度とスコープ、一般的なクエリ パターンとベスト プラクティス、使用ルール (データを使用するタイミング、結合キーと必要なフィルター) などのデータ使用の詳細を含む、ドメイン固有のドキュメントまたはプロセスを備えたチームのリポジトリを指すことができます。
スキル アセットを使用すると、AI エージェントがデータ資産内で使用する適切なデータを見つけやすくなりますが、それは問題の半分にすぎません。エージェントは、データを集計する前に適用するフィルター、従うべき結合パス、技術スキーマには表示されない注意事項など、その使用方法も知っておく必要があります。現在、AWS エージェント ツールキットには、AI エージェントが Glue データ カタログや Amazon Athena や S3 テーブルなどの他の機能を操作するのに役立つデフォルトのスキルが含まれています。多くの企業は、データ チームが開発した独自のスキルを持っています。まず開発者は、リモートのフルマネージド AWS MCP を使用して MCP 互換エージェントに接続し、AWS サービス スキルにアクセスするか、Claude Code、Cursor、Amazon Kiro 用の aws-data-analytics プラグインをインストールして、AWS またはその他のカスタム スキルを使用してデータの検索、分析の実施、またはそのデータ上にアプリケーションの構築をエージェントに依頼できます。 AgentC で構築されたエージェント

ore harness は、1 行のコードで AWS Agent Toolkit のすべての AWS スキルにアクセスできます。これにより、エージェントは AWS サービスの専門知識とベストプラクティスを迅速に活用できるようになります。
Amazon S3 アノテーション (正式リリース)
お客様が独自のカスタム コンテキストをデータレイクに簡単に追加できるようにするために、Amazon S3 アノテーションの一般提供を発表しました。これは、豊富でクエリ可能なビジネス コンテキストを S3 オブジェクトに直接添付し、そのコンテキストを S3 Iceberg テーブルに保存する新しい方法です。お客様は長い間、S3 でオブジェクト タグとユーザー定義のメタデータを使用してオブジェクトを記述してきましたが、これらはアクセス制御などの運用タスクや、アップロード時に設定される小さな情報に適したツールであり続けています。しかし、顧客はデータに基づいてエージェントを構築するにつれて、はるかに多くのメタデータを添付したいと考えています。彼らは、エージェントが読み取り、それに基づいて行動できる豊富なコンテキストを大規模に作成し、進化させたいと考えています。 S3 アノテーションは、その機能をオープン データ形式で提供します。 S3 に保存される各オブジェクトには、最大 1 GB のコンテキストを含めることができます。注釈は変更可能であるため、データの変更に応じてコンテキストを進化させることができます。 S3 アノテーションは、S3 ストレージ内の S3 オブジェクトとともに存在します。つまり、S3 アノテーションは、コピーおよびレプリケーション操作を通じて、関連付けられた S3 オブジェクトとともに移動し、オブジェクトが削除されると削除されます。注釈を使用すると、メタデータ データベースを個別に構築、同期したり、古くなったりするのを防ぐ必要がありません。
注釈は S3 メタデータを通じてクエリ可能になります。バケットでアノテーション テーブルを有効にすると、すべてのアノテーションがフルマネージドの Iceberg テーブルに自動的に流れ込みます。 Amazon Athena、Amazon Redshift、または任意の Iceberg 互換エンジンを使用してすべてのオブジェクトに対してクエリを実行でき、エージェントは S3 Tables MCP サーバーを通じて自然言語で注釈を検出できます。
A付き

mazon S3 アノテーションを使用すると、豊富なビジネス コンテキストを S3 オブジェクトに直接添付して大規模にクエリできるため、エージェントは別個のメタデータ システムを構築することなく、必要なものを見つけることができます。
コンテキストは AI エージェントのデータ レイクであり、これらのイノベーションにより、あらゆる規模の組織や企業にわたってデータを操作する AI エージェントのための知識とインテリジェンスの基盤を構築しています。
AWS のテクノロジー副社長である Mai-Lan Tomsen Bukovec 氏は、何百万もの AWS の顧客がデジタル変革、ビジネス分析、機械学習、生成 AI、次世代のカスタマー エクスペリエンスに依存している Amazon クラウド データ サービスを率いています。テクノロジー業界で 25 年以上の経験を持つ Mai-Lan は、顧客がクラウドベースのテクノロジーを活用してビジネスを変革できるよう支援するパイオニアです。
英語
トップに戻る
Amazon は、マイノリティ / 女性 / 障害者 / 退役軍人 / 性同一性 / 性的指向 / 年齢の機会均等雇用主です。
×
フェイスブック
リンクイン
インスタグラム
けいれん
ユーチューブ
ポッドキャスト
電子メール
プライバシー

## Original Extract

Agents are only as intelligent as the context they can reason over. Today, that context is scattered across data lakes, data warehouses, lakehouses, databases, and streams, and in institutional knowledge that has never been written down. You want to trust the decisions made by your AI agents, but th
[truncated]

Context intelligence for your data and AI agents at scale
Agents are only as intelligent as the context they can reason over. Today, that context is scattered across data lakes, data warehouses, lakehouses, databases, and streams, and in institutional knowledge that has never been written down. You want to trust the decisions made by your AI agents, but that can’t happen until agents have context. Imagine what becomes possible when we give agents a safe way to access the context they need to deliver trusted decisions.
This is why at the AWS Summit New York City , we’re announcing a series of innovations that deliver context intelligence for your data and AI agents at scale.
In today’s keynote, we introduced AWS Context, a new service that automatically maps the relationships across your existing data into a knowledge graph and provides agentic search so AI agents in the organization can access governed data relationships, business rules, and domain knowledge at runtime. Data stewards and curators manage the graph through an intuitive console experience, reviewing inferred relationships, promoting them to production, and attaching domain-specific knowledge like business definitions and usage rules.
AWS Context extends the same knowledge graph technology that powers Amazon Quick, where hundreds of thousands of users interact daily with a production knowledge graph that catalogs datasets, dashboards, and metadata, learning from usage patterns to make every interaction smarter. That graph already processes millions of requests per day. With AWS Context, we are extending what was a personal knowledge graph into an organizational one, a shared, governed context layer that agents and applications in your organization can draw from. Existing Amazon Quick users benefit immediately. When AWS Context is enabled, Quick’s agents gain access to the broader enterprise knowledge graph, including cross-system relationships, business rules, and curated context that go beyond what any single user’s personal graph can provide. AWS Glue Data Catalog, Amazon SageMaker Unified Studio, and AWS Lake Formation integrate with the knowledge graph, so teams can govern it with business rules and permissions and add new context automatically with AI assistance or explicitly through manual curation.
Key elements of the context layer are published to Amazon S3 in the Apache Iceberg format, so that customers are free to use the Iceberg-compliant tools of their choice to consume metadata and build against AWS Context based on open standards. There is no infrastructure to provision or retrieval pipeline to build, and customers can begin to gather and curate context for their agents with just a few clicks in the AWS Management Console.
Let’s take a closer look at the capabilities behind it.
Context that learns from how your agents work
AWS Context gets smarter the more your agents use it. As agents query the graph, it observes which sources produce correct results, which join paths agents rely on, and which curated rules get applied. It ranks sources by actual usage and shares what it learns across your organization, so when one agent discovers a correct join path or resolves a schema ambiguity, other agents pick it up, without requiring a human re-curate the graph.
AWS Context publishes all key metadata from structured and unstructured sources into Apache Iceberg format in Amazon S3 Tables, so you can query your context with Amazon Athena, Amazon Redshift, Apache Spark, or any Iceberg-compatible engine, and build downstream systems on it, audit it, or migrate it.
AWS Context is also designed to connect to third-party catalogs, so you can bring context from systems beyond AWS into the same graph. Agents query it through agentic search APIs and MCP tools, whether they’re built on Amazon Bedrock AgentCore, deployed on Amazon EKS, or running on MCP-compatible frameworks. Your context stays queryable, portable via Apache Iceberg format, and fully yours.
Identity-aware and governed by default
Any agent you put into production raises a governance question: what data can it reach, and can you show exactly what it accessed and under whose authority? AWS Context answers both by making every query identity-aware. Each call is designed to inherit the calling user’s IAM and Lake Formation permissions, so an agent can only see and traverse the relationships its identity is authorized to access. Because access runs through identity, every interaction is auditable. Your security and compliance teams can verify what an agent accessed and under what authority, using the same controls you already rely on.
AWS Glue Data Catalog Business Context and Semantic Search (preview)
Today, we also announced the preview of business context and semantic search for AWS Glue Data Catalog, providing context and tools that makes it easier for humans and AI agents to discover and understand data. Customers can now enrich their Glue tables, views, and columns, including those backed by S3 Tables, with business descriptions, glossary terms, custom metadata, and associate them with skill assets that provide additional data context stored outside the catalog. With business context indexed alongside technical metadata in Glue Data Catalog, customers can use the new Glue Search API to more quickly find data by business meaning and AI agents can ground their reasoning in trusted definitions rather than inferring context.
We are also excited to offer a preview of skill assets in Glue Data Catalog. Now, data producers can create skill assets, a new asset type that references URIs to files (such as AI skills, guide markdown files, and team runbooks) hosted in any location including S3, git repositories, and wikis. Associating skill assets to data assets gives agents additional context and instructions they can retrieve progressively for working with specific data without re-teaching it to every agent one prompt at a time. For example, the skill URI locations can point to your team’s repositories with domain-specific documentation or processes that include data usage details such as the grain and scope, common query patterns and best practices, and usage rules (when to use the data, what are join keys and required filters).
Skill assets make it easier for AI agents to find the right data to use in a data estate but that is only half the problem. An agent also has to know how to use it: the filters to apply before aggregating data, the join paths to follow, the caveats that aren’t visible in the technical schema. Today, the AWS Agent Toolkit contains default skills to help AI agents work with Glue Data Catalog as well as other capabilities like Amazon Athena and S3 Tables. Many enterprises have their own skills that their data teams have developed. To get started, developers can connect any MCP-compatible agents using the remote, fully managed AWS MCP to access AWS service skills or by installing the aws-data-analytics plugin for Claude Code, Cursor, and Amazon Kiro, to ask an agent to find data, conduct analysis, or build applications on top of that data using AWS or other custom skills. Agents built with AgentCore harness can access all AWS skills in the AWS Agent Toolkit with one line of code. This enables your agents to take advantage of the AWS service expertise and best practices quickly.
Amazon S3 Annotations (generally available)
To make it easier for customers to add their own custom context to their data lake, we announced the general availability of Amazon S3 annotations, a new way to attach rich, queryable business context directly to your S3 objects and store that context in an S3 Iceberg table. Customers have long described their objects in S3 with object tags and user-defined metadata, and those remain the right tools for operational tasks like access control and for small pieces of information set at upload. But as customers build agents on top of their data, they want to attach far more metadata. They want to create and evolve rich context that an agent can read and act on, at scale. S3 annotations provide that capability in an open data format. Each object stored in S3 can have up to 1 GB of context. Annotations are mutable, so you can evolve context as your data changes. S3 annotations live with the S3 object in S3 storage. That means that S3 annotations move with its associated S3 object through copy and replication operations, and they’re removed when the object is deleted. With annotations, there’s no separate metadata database to build, synchronize, or keep from drifting out of date.
Annotations become queryable through S3 Metadata. When you enable annotation tables on a bucket, every annotation flows automatically into a fully managed Iceberg table. You can query across all your objects with Amazon Athena, Amazon Redshift or any Iceberg-compatible engine, and agents can discover annotations in natural language through the S3 Tables MCP server.
With Amazon S3 annotations , you attach rich business context directly to S3 objects and query it at scale, so agents find what they need without building separate metadata systems.
Context is the data lake for AI agents, and with these innovations, we are building the foundation of knowledge and intelligence for AI agents interacting with data across organizations and enterprises of any scale.
Mai-Lan Tomsen Bukovec, Technology Vice President at AWS, leads the Amazon cloud data services that millions of AWS customers rely on for digital transformations, business analytics, machine learning, generative AI, and next generation customer experiences. With over 25 years of experience in the technology industry, Mai-Lan is a pioneer in helping customers take advantage of cloud-based technologies to transform their businesses.
English
Back to top
Amazon is an Equal Opportunity Employer: Minority / Women / Disability / Veteran / Gender Identity / Sexual Orientation / Age.
x
facebook
linkedin
instagram
twitch
youtube
podcasts
email
Privacy

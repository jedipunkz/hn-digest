---
source: "https://www.esri.com/en-us/software-engineering/blog/articles/ai-orchestration-in-the-browser"
hn_url: "https://news.ycombinator.com/item?id=48811674"
title: "Why AI Orchestration Belongs in the Browser"
article_title: "Why AI Orchestration Belongs in the Browser"
author: "dhrumil83"
captured_at: "2026-07-07T00:08:12Z"
capture_tool: "hn-digest"
hn_id: 48811674
score: 2
comments: 0
posted_at: "2026-07-06T23:10:45Z"
tags:
  - hacker-news
  - translated
---

# Why AI Orchestration Belongs in the Browser

- HN: [48811674](https://news.ycombinator.com/item?id=48811674)
- Source: [www.esri.com](https://www.esri.com/en-us/software-engineering/blog/articles/ai-orchestration-in-the-browser)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T23:10:45Z

## Translation

タイトル: AI オーケストレーションがブラウザに含まれる理由
説明: ブラウザーでの AI オーケストレーションが Web GIS に最適である理由を説明します。Web GIS では、エージェントがライブ マップの状態、レイヤー、ユーザー インタラクションを直接操作できます。

記事本文:
AI オーケストレーションがブラウザーに属する理由
Esri ソフトウェア エンジニアリング ブログ
メニュー
概要
AI オーケストレーションがブラウザーに属する理由
現在、ほとんどの AI アシスタントはサーバー上で実行されます。ユーザーは Web アプリケーションと対話し、リクエストはバックエンドに送信され、バックエンドはオーケストレーション、エージェント、ツール、アプリケーション ロジック、および LLM との対話を処理します。
この記事では、ArcGIS Maps SDK for JavaScript のバージョン 5.0 の一部としてリリースされた AI コンポーネントの背後にあるアーキテクチャと、オーケストレーションをアプリケーション コンテキストの多くがすでに存在する場所、つまりブラウザーの近くに移動した理由について説明します。
この文脈において、オーケストレーションは、どのエージェントを呼び出すか、どのようなコンテキストを提供するか、そしてその結果を有用な応答に変える方法を決定する調整層です。
Web マッピング アプリケーションが異なる理由
しかし、Web マッピング アプリケーションは異なります。
ArcGIS Maps SDK for JavaScript を使用して構築されたアプリケーションは、すでに高度なクライアント機能を備えています。ブラウザには、すでにかなりの量のランタイム コンテキストが含まれています。
アプリケーションの状態とビジネス ロジック
ArcGIS Maps SDK for JavaScript は、UI 状態とマップ コンテキストに加えて、クエリ、ジオメトリ処理、GPU アクセラレーションによる WebGL レンダリングなどの実質的なクライアント側 API とワークフローも提供します。これにより、高速で高度にインタラクティブなエクスペリエンスが可能になり、より多くのアプリケーション ロジックをブラウザーで直接実行できるようになります。
これらのクライアント側のワークフローの例については、「距離による FeatureLayerView クエリ」および「ジオメトリによる FeatureLayerView クエリ」を参照してください。
ブラウザーには、マップを理解して操作するために必要なライブ ランタイム コンテキストの多くがすでに組み込まれています。
Web GIS アプリケーション用の AI アシスタントの構築を開始したとき、私たちは次のような単純な質問をしました。
ブラウザネイティブ AI アーキテクチャ
すべてのマップ コンテキストとアプリケーションをプッシュする代わりに

状態をバックエンド オーケストレーション レイヤーに変換するため、私たちは別のアプローチを検討しました。それは、AI オーケストレーションとエージェントをマップ自体と並行してブラウザーで直接実行することでした。
アーキテクチャは以下を中心に構築されています。
ブラウザで実行されるエージェントとツール
ライブマップ状態との直接対話
ハイブリッド ブラウザ/サーバー AI ワークフロー
これにより、追加のバックエンド オーケストレーション サービスの必要性を減らしながら、AI アシスタントを Web マッピング アプリケーションと統合できるようになります。
ブラウザネイティブのオーケストレーションが実際にどのように機能するかを理解するには、次のリクエストを検討してください。
オーケストレーションのためにリクエストをバックエンド サービスに送信する代わりに、ブラウザ側のオーケストレーターによって処理されます。
使用可能なエージェントとその機能に基づいて、オーケストレーターはタスクを処理するデータ探索エージェントを選択します。
コンテキストを LLM に送信する前に、アプリケーションはブラウザーでベクトル検索を実行し、リクエストに最も関連するレイヤーとフィールドを特定します。これは、モデルに送信されるコンテキストの量を絞り込むのに役立ちます。
選択されたエージェントは、適切なクエリを生成し、マップを更新し、一致する地物を強調表示して、結果を要約してユーザーに返します。
オーケストレーションはブラウザーのランタイム コンテキストの近くで実行されるため、エージェントは重いバックエンド層を使用せずにライブ マップの状態と直接対話できます。
ベクトル検索は、ブラウザネイティブのオーケストレーションを実用化する重要な部分の 1 つです。
埋め込みは、意味的な類似性の比較を可能にするベクトル表現です。
ArcGIS Web マップには多くのレイヤーを含めることができ、各レイヤーには多くのフィールドを含めることができます。リクエストごとにすべてのメタデータを LLM に送信するのは非効率であり、モデルが間違ったコンテキストを使用する可能性が高くなります。
代わりに、ブラウザーで埋め込みとベクトル検索を使用して、最も関連性の高いレイヤーとフィールドを特定します。

ユーザーのリクエストに応じて。これは、LLM が呼び出される前のコンテキスト エンジニアリング ステップとして機能します。
エンベディングはエンベディング モデルを呼び出すことによって生成されますが、レイヤーとフィールドのエンベディングが利用可能になると、ベクトル検索自体がブラウザーで実行されます。
たとえば、「テキサス州の老人ホームを表示」ワークフローでは、エージェントがクエリを生成する前にベクトル検索を使用して、老人ホーム レイヤーと関連する州フィールドを特定することができます。
これにより、プロンプトが小さくなり、より焦点が絞られ、実際のマップに根ざしたものになります。
ブラウザネイティブのオーケストレーションは、すべてがブラウザ内で実行されるという意味ではありません。ブラウザーは、アプリケーション コンテキスト、オーケストレーション、エージェント、およびマップ インタラクションが存在する場所ですが、アーキテクチャでは引き続きサーバーでホストされるモデル、エンタープライズ API、および長時間実行される AI タスクを使用できます。
ブラウザネイティブのオーケストレーションは、あらゆるシナリオに適合するわけではありません。大規模なバッチ処理、長時間実行されるワークフロー、独自のプロンプトなどの機密性の高いオーケストレーション ロジックは、依然としてサーバー側での実行に適している可能性があります。
キーシフトはバックエンドを削除することではありません。これにより、オーケストレーションがブラウザーにすでに存在するランタイム コンテキストに近づきます。
これが開発者にとって重要な理由
すでに ArcGIS Maps SDK for JavaScript を使用して構築している Web 開発者にとって、このアーキテクチャでは、マップがすでに実行されているのと同じブラウザー ランタイムで AI ワークフローの多くが維持されます。開発者は、ワークフローごとに個別のバックエンド オーケストレーション サービスを作成する代わりに、アプリケーションと直接対話するエージェントとツールを登録できます。
オーケストレーション、エージェント、ツール、マップのインタラクションはブラウザーネイティブであるため、CodePen のような環境で AI 支援マップ アプリケーションのプロトタイプを作成することもできます。これは、運用アプリケーションに移行する前に統合パターンを検討するのに役立ちます。
デブ

また、lopers は、同じブラウザネイティブのオーケストレーション パターンを使用しながら、独自のワークフロー用にカスタム エージェントとツールを登録することもできます。詳細については、カスタム エージェントのドキュメントを参照してください。
簡単な実行例 (組織内で有効になっている ArcGIS 認証情報と AI アシスタントが必要) については、この CodePen デモを参照してください。
マップ、アプリケーションの状態、およびユーザー コンテキストがすでにブラウザー内に存在するため、ブラウザー ネイティブのオーケストレーションは Web GIS で特に役立ちます。
このパターンは GIS に限定されません。複雑な状態、ドメイン固有のツール、およびインタラクティブなワークフローを備えたリッチ ブラウザ アプリケーションは、オーケストレーションをユーザー コンテキストがすでに存在する場所に近づけることで恩恵を受ける可能性があります。
将来はハイブリッドになる可能性が高く、サーバーホスト型モデル、エンタープライズ API、および長時間実行される AI タスクと連携して動作するブラウザー側のオーケストレーションです。
今後の投稿では、クライアント側エージェントのワークフロー、安全なエンタープライズ モデル アクセス、ブラウザ ベースのベクトル検索、およびブラウザ/サーバーのハイブリッド実行パターンなど、この作業の背後にあるアーキテクチャをさらに深く掘り下げる予定です。
Esri の実践技術コミュニティがどのように規模を構築し、つながりを築くのか

## Original Extract

Learn why AI orchestration in the browser is a strong fit for Web GIS, where agents can work directly with live map state, layers, and user interactions.

Why AI Orchestration Belongs in the Browser
Esri Software Engineering Blog
Menu
Overview
Why AI Orchestration Belongs in the Browser
Most AI assistants today run on the server. The user interacts with a web application, the request gets sent to a backend, and the backend handles the orchestration, agents, tools, application logic, and interaction with the LLM.
This article looks at the architecture behind the AI Components released as part of version 5.0 of the ArcGIS Maps SDK for JavaScript, and why we moved orchestration closer to where much of the application context already exists: the browser.
In this context, orchestration is the coordination layer that decides which agents to call, what context to provide them, and how to turn their results into a useful response.
Why Web Mapping Applications Are Different
But web mapping applications are different.
Applications built with the ArcGIS Maps SDK for JavaScript are already highly client-rich. The browser already contains a significant amount of runtime context:
application state and business logic
Beyond UI state and map context, the ArcGIS Maps SDK for JavaScript also provides substantial client-side APIs and workflows, including querying, geometry processing, and GPU-accelerated WebGL rendering. This enables fast, highly interactive experiences and allows more application logic to run directly in the browser.
For examples of these client-side workflows, see FeatureLayerView query by distance and FeatureLayerView query by geometry .
The browser already has much of the live runtime context needed to understand and interact with the map.
As we started building AI assistants for Web GIS applications, we asked a simple question:
Browser-Native AI Architecture
Instead of pushing all map context and application state to a backend orchestration layer, we explored a different approach: running AI orchestration and agents directly in the browser, alongside the map itself.
The architecture is built around:
browser-executed agents and tools
direct interaction with live map state
hybrid browser/server AI workflows
This allows AI assistants to integrate with web mapping applications while reducing the need for additional backend orchestration services.
To understand how browser-native orchestration works in practice, consider the following request:
Instead of sending the request to a backend service for orchestration, it is handled by the browser-side orchestrator.
Based on the available agents and their capabilities, the orchestrator selects the data exploration agent to handle the task.
Before sending context to the LLM, the application performs vector search in the browser to identify the most relevant layers and fields for the request. This helps narrow the amount of context sent to the model.
The selected agent then generates the appropriate query, updates the map, highlights the matching features, and summarizes the result back to the user.
Because orchestration runs near the runtime context in the browser, agents can interact directly with live map state without a heavy backend layer.
Vector search is one of the key pieces that makes browser-native orchestration practical.
Embeddings are vector representations that allow semantic similarity comparisons.
An ArcGIS web map can contain many layers, and each layer can contain many fields. Sending all of that metadata to the LLM for every request is inefficient and can make the model more likely to use the wrong context.
Instead, we use embeddings and vector search in the browser to identify the most relevant layers and fields for the user’s request. This acts as a context engineering step before the LLM is called.
Embeddings are generated by calling an embedding model, but once the layer and field embeddings are available, the vector search itself runs in the browser.
For example, in the “Show nursing homes in Texas” workflow, vector search helps identify the nursing homes layer and the relevant state field before the agent generates the query.
This keeps the prompt smaller, more focused, and more grounded in the actual map.
Browser-native orchestration does not mean everything runs in the browser. The browser is where the application context, orchestration, agents, and map interactions live, but the architecture can still use server-hosted models, enterprise APIs, and long-running AI tasks.
Browser-native orchestration is not a fit for every scenario. Large-scale batch processing, long-running workflows, and sensitive orchestration logic such as proprietary prompts may still be better suited for server-side execution.
The key shift is not removing the backend. It is moving orchestration closer to the runtime context that already exists in the browser.
Why This Matters for Developers
For web developers already building with the ArcGIS Maps SDK for JavaScript, this architecture keeps much of the AI workflow in the same browser runtime where the map already runs. Instead of creating a separate backend orchestration service for every workflow, developers can register agents and tools that interact with the application directly.
Because the orchestration, agents, tools, and map interactions are browser-native, AI-assisted map applications can also be prototyped in environments like CodePen. This is useful for exploring the integration pattern before moving to a production application.
Developers can also register custom agents and tools for their own workflows while still using the same browser-native orchestration pattern. To learn more, see the custom agents documentation .
For a simple running example (requires ArcGIS credentials and AI Assistants enabled in the organization ), see this CodePen demo .
Browser-native orchestration is especially useful for Web GIS because the map, application state, and user context already live in the browser.
This pattern is not limited to GIS. Any rich browser application with complex state, domain-specific tools, and interactive workflows could benefit from moving orchestration closer to where the user context already exists.
The future is likely hybrid: browser-side orchestration working together with server-hosted models, enterprise APIs, and long-running AI tasks.
In future posts, we plan to go deeper into the architecture behind this work, including client-side agent workflows, secure enterprise model access, browser-based vector search, and hybrid browser/server execution patterns.
How Esri’s Technical Communities of Practice Scale Craft and Connection

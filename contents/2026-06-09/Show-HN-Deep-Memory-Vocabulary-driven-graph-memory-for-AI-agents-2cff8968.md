---
source: "https://github.com/TjWheeler/deep-memory"
hn_url: "https://news.ycombinator.com/item?id=48455368"
title: "Show HN: Deep Memory – Vocabulary-driven graph memory for AI agents"
article_title: "GitHub - TjWheeler/deep-memory: A GraphRAG implementation with a Vocabulary system to optimise AI integration · GitHub"
author: "tjwheeler"
captured_at: "2026-06-09T04:29:49Z"
capture_tool: "hn-digest"
hn_id: 48455368
score: 2
comments: 1
posted_at: "2026-06-09T02:12:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Deep Memory – Vocabulary-driven graph memory for AI agents

- HN: [48455368](https://news.ycombinator.com/item?id=48455368)
- Source: [github.com](https://github.com/TjWheeler/deep-memory)
- Score: 2
- Comments: 1
- Posted: 2026-06-09T02:12:49Z

## Translation

タイトル: Show HN: Deep Memory – AI エージェント向けの語彙駆動型グラフ メモリ
記事のタイトル: GitHub - TjWheeler/deep-memory: AI 統合を最適化するための語彙システムを備えた GraphRAG 実装 · GitHub
説明: AI 統合を最適化するための語彙システムを備えた GraphRAG 実装 - TjWheeler/deep-memory

記事本文:
GitHub - TjWheeler/deep-memory: AI 統合を最適化するための語彙システムを備えた GraphRAG 実装 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
TjWheeler
/
深い記憶
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
71 コミット 71 コミット .changeset .changeset .github .github データベース データベース docs docs エクスポート エクスポート インデックスコンテンツ/人物 インデックスコンテンツ/人物 インデックススターターキット インデックススターターキット パッケージ パッケージ スクリプト スクリプト .env.example .env.example .gitignore .gitignore .mcp.json.example .mcp.json.example CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.indexer.yml docker-compose.indexer.yml docker-compose.neo4j.yml docker-compose.neo4j.yml docker-compose.sqlserver.yml docker-compose.sqlserver.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Quickstart-claude-desktop.mdクイックスタート-claude-desktop.md クイックスタート-cosmosdb.md クイックスタート-cosmosdb.md クイックスタート-embeddings.md クイックスタート-embeddings.md クイックスタート-インデクサー.md クイックスタート-インデクサー.md クイックスタート-inmemory.md クイックスタート-inmemory.md クイックスタート-neo4j.md クイックスタート-neo4j.mdクイックスタート-sqlserver.md クイックスタート-sqlserver.md tsconfig.base.json tsconfig.base.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の語彙駆動型グラフ メモリ ライブラリ。 AI に、摩擦なしで読み取り、書き込み、および移動できる、構造化された永続的なメモリを与えます。
ナレッジ グラフを操作する AI エージェントは、コールド スタートの問題に直面します。つまり、単一のファクトを保存する前に、どのようなエンティティ タイプが存在するか、各タイプがどのようなプロパティを持ち、どのような関係タイプがそれらを接続しているのか、どのような制約が適用されるのかを理解する必要があります。これがなければ、すべての対話は推測、矛盾、トークンの無駄から始まります。
ディープメモリーは語彙ガバナンスでこれを解決します。 AGのとき

ent がリポジトリを開くと、最初にボキャブラリが表示されます。これは、エンティティ タイプ、関係タイプ、プロパティ定義、および関係制約の完全なスキーマです。これは次のことを意味します:
推測は不要です。エージェントは、どのようなタイプが存在し、どのようなプロパティを設定すべきかを正確に知っています。 1 回のツール呼び出しでエンティティを作成できます。
矛盾はありません。検証では、グラフに入力される前に、無効な型と欠落している必要なプロパティを検出します。同じリポジトリ上で動作する 2 つのエージェントは、構造的に一貫したデータを生成します。
伸ばすときに摩擦がありません。オープン ガバナンス モードでは、エージェントが新しいタイプを提案し、重複排除により自動承認されます。エージェントは承認を待つために作業を止めることはありません。
効率的なトークンの使用。語彙はコンパクトなスキーマであり、何千もの文書例ではありません。エージェントは数百のトークンで完全な語彙を習得し、自信を持って作業できます。
その結果、AI エージェントは、トレーニング データや少数ショットのサンプル、試行錯誤を必要とせずに、リポジトリを開いてエンティティを作成し、リレーションシップを横断するまでを数秒で行うことができます。
すぐに始めるための注意点
このリポジトリは、AI が使用できる詳細なドキュメントを提供するように設計されています。 AI はセットアップと構成をガイドし、作業を開始するのに役立ちます。
クイックスタートを選択し、そのプロンプトをクロード コードに貼り付けます。
Quickstart-inmemory.md — セットアップ不要。グラフはメモリ内に存在します。最速 2 分のパス。
Quickstart-sqlserver.md — SQL Server の永続ストレージ。ローカル使用のためにバンドルされた Docker Compose。
Quickstart-cosmosdb.md — Azure CosmosDB Gremlin API (Windows 上のローカル エミュレーター、または Azure アカウント)。
Quickstart-neo4j.md — Neo4j グラフ データベース。バンドルされた Docker Compose、または AuraDB / 自己ホスト型。
Quickstart-indexer.md — インデックス作成パイプラインを介してソースドキュメントから独自のナレッジグラフを構築します

。
開発者ではないですか? Claude Code や別の MCP 対応 AI の代わりに Claude Desktop を使用しますか?
Quickstart-claude-desktop.md — コーディング、git、ビルドは不要です。 3 つのインストーラー + 1 つの構成ファイル。 ～15分。
パッケージ
パス
説明
@utaba/ディープメモリー
パッケージ/コア
コア ライブラリ — ランタイム DEPS ゼロ、tsup 経由のデュアル CJS/ESM
@utaba/deep-memory-embeddings-openai
パッケージ/埋め込み-openai
OpenAI 互換の埋め込みプロバイダー (vLLM、OpenAI、Azure、Ollama)
@utaba/ディープメモリストレージ-sqlserver
パッケージ/ストレージ-sqlserver
SQL Server ストレージ プロバイダー - 永続的なマルチテナント グラフ ストレージ
@utaba/ディープメモリストレージ-cosmosdb
パッケージ/ストレージ-cosmosdb
CosmosDB Gremlin ストレージ + グラフ トラバーサル プロバイダー
@utaba/ディープメモリストレージ-neo4j
パッケージ/ストレージ-neo4j
Neo4j ストレージ + グラフ トラバーサル プロバイダー (Cypher over Bolt、コミュニティ エディション)
@utaba/ディープメモリーインデクサー
パッケージ/インデクサー
ドキュメントのインデックス作成パイプライン — LLM の抽出、検証、統合、インポート
@utaba/deep-memory-indexer-llm-anthropic
パッケージ/indexer-llm-anthropic
インデクサー用の Anthropic LLM プロバイダー - プロンプト キャッシュ + ネイティブ メッセージ API
@utaba/ディープメモリインデクサー-mcp-server
パッケージ/インデクサー-mcp-サーバー
インデクサー MCP サーバー — インデックス作成パイプライン用の 9 つのフェーズ対応ツール
@utaba/ディープメモリローカルmcpサーバー
パッケージ/mcp-server
AI エージェント用のディープメモリ メモリ ツールを公開するローカル MCP サーバー
ディープメモリライブラリ
コアの @utaba/deep-memory ライブラリは、型付きエンティティ、関係、プロパティ検証、プラグイン可能なストレージ/検索/埋め込みプロバイダーを備えた語彙主導のナレッジ グラフを提供します。実行時の依存関係はゼロです。ローカル MCP サーバー ( @utaba/deep-memory-local-mcp-server ) は、AI エージェントがリポジトリを作成し、エンティティと関係を管理し、グラフを走査し、セマンティック検索を実行するための 20 以上のツールを公開します。それは

リファレンス統合と、ナレッジ グラフのテストおよび操作のための本番対応インターフェイスの両方として提供されます。
詳細なアーキテクチャ、プロバイダーのセットアップ、および API ドキュメントについては、docs/README.md を参照してください。
注: インデックス作成パイプラインはオプションです。 AI を使用しなくても、メモリ グラフを使用して設定することができます。インデックス作成パイプラインは、インポートするドキュメントがある場合に最も役立ちます。
インデクサー パッケージ ( @utaba/deep-memory-indexer 、 @utaba/deep-memory-indexer-mcp-server 、 @utaba/deep-memory-indexer-llm-anthropic ) は、ソース ドキュメントを構造化ナレッジ グラフに変換するための AI 駆動のパイプラインを提供します。このパイプラインは、ドキュメント分析、マルチモデル LLM 抽出、品質検証、ドキュメント間エンティティ重複排除、リポジトリ インポート、および埋め込み生成を処理します。インデクサー MCP サーバーは、ソース ドキュメントのスキャンからクエリ対応のナレッジ グラフまで、プロセス全体を通じて AI エージェントをガイドする 9 つのフェーズ認識ツールを公開します。
マルチワーカー抽出 (ローカル + クラウド モデルの並列)
3 段階の検証 (構造的、LLM 検証、人間によるレビュー)
あらゆる段階でのコストの追跡と見積もり
共通ドメイン用のスターター キット
大きな文書の段階的な章抽出
Anthropic プロバイダーによるプロンプト キャッシュ (繰り返しのシステム プロンプトで最大 88% 節約)
クイックスタート、モデルの選択、トラブルシューティングなど、インデックス作成パイプラインの完全なドキュメントについては、docs/README.md を参照してください。
スターター キットは、一般的なドメイン用に事前に構築された語彙です。各キットはエンティティ タイプ、関係タイプ、プロパティ スキーマ、およびインデックス作成プロセスを定義するため、エージェントはナレッジ グラフへの入力をすぐに開始できます。
README.md — 概要 (目的、ユースケース、ガバナンス) + MCP ツールによる手動インデックス作成プロセス
語彙.md — エンティティ タイプと関係タイプ

プロパティの定義
domain-guidance.md — AI エージェントのドメイン固有の抽出ガイダンス (自動パイプラインによって使用されます)
スターター キットからリポジトリを作成するには、AI エージェントに次のように指示します。
「パーソン スターター キットを使用して、新しいディープ メモリ リポジトリを作成します。」
スターター キットは拡張可能です。オープン ガバナンス モードでは、エージェントはいつでも新しいエンティティおよび関係タイプを提案できます。
Node.js 22 または 24 (サポートされている LTS ペア - CI は両方をテストします)
リポジトリにはサービスごとの Compose ファイルが同梱されているため、必要なものだけを開始できます。
# SQL Server を起動します (永続ストレージに必要)
docker compose -f docker-compose.sqlserver.yml アップ -d
これにより、SQL Server 2022 Developer Edition コンテナーがポート 1435 で起動されます。データベース スキーマは、MCP サーバーが初めて接続したときに自動的に作成されます。
NVIDIA GPU を使用していて、セマンティック検索 (概念検索、埋め込みによる語彙の重複排除) が必要な場合は、埋め込みサービスも開始します。
# vLLM 埋め込みサーバーを起動します
docker compose -f docker-compose.indexer.yml --profile embeddings up -d
vLLM サービスは、ポート 8010 で Qwen3-Embedding-8B モデルを提供します。これはオプションです。これがないと、ディープメモリは語彙の重複排除のために文字列の類似性にフォールバックし、memory_search_by_concept が使用できなくなります。
pnpmインストール
pnpmビルド
3. データベースの作成
SQL Server が正常になったら、好みの SQL ツールに接続してデータベース (例: deep-memory ) を作成します。テーブル スキーマは、最初の接続時にストレージ プロバイダーによって自動的に作成されます。
MCP サーバーは、標準入出力を介してディープ メモリをモデル コンテキスト プロトコル ツールとして公開するローカル プロセスです。これは、AI エージェント (Claude Code、Claude Desktop など) と並行してローカルで実行し、エージェントがグラフ メモリ リポジトリを作成して操作できるようにすることを目的としています。参考にもなります

ディープメモリを独自のアプリケーションに組み込むための統合。
以下を .mcp.json (プロジェクト ルート、またはグローバルの場合は ~/.mcp.json) に追加します。
{
"mcpサーバー": {
"ディープメモリ" : {
"コマンド" : "ノード" ,
"args" : [ "<リポジトリへのパス>/packages/mcp-server/dist/index.js " ],
"環境" : {
"DEEP_MEMORY_ACTOR_ID" : " mcp-agent " ,
"DEEP_MEMORY_ACTOR_TYPE" : " エージェント " 、
"DEEP_MEMORY_STORAGE" : " sqlserver " 、
"DEEP_MEMORY_SQL_HOST" : " localhost " 、
"DEEP_MEMORY_SQL_PORT" : " 1435 " 、
"DEEP_MEMORY_SQL_DATABASE" : "ディープメモリ" ,
"DEEP_MEMORY_SQL_USER" : " sa " ,
"DEEP_MEMORY_SQL_PASSWORD" : " DeepMem@Dev1234 " ,
"DEEP_MEMORY_SQL_SCHEMA" : " dbo " ,
"DEEP_MEMORY_SQL_TRUST_CERT" : " true " ,
"DEEP_MEMORY_EMBEDDINGS_BASE_URL" : " http://localhost:8010 " 、
"DEEP_MEMORY_EMBEDDINGS_MODEL" : " Qwen/Qwen3-Embedding-8B "
}
}
}
}
セキュリティに関する注意: 構成内の SQL パスワードと Docker サーバー設定を変更することを検討してください。
変数
デフォルト
説明
DEEP_MEMORY_ACTOR_ID
mcp-エージェント
出自メタデータにスタンプされた俳優 ID
DEEP_MEMORY_ACTOR_TYPE
エージェント
アクターのタイプ: エージェント 、 人間 、または システム
DEEP_MEMORY_STORAGE
記憶
ストレージ バックエンド: メモリ (メモリ内、再起動時にデータが失われる)、 sqlserver 、 cosmosdb 、または neo4j
DEEP_MEMORY_SQL_HOST
—
SQLサーバーのホスト名
D

[切り捨てられた]

## Original Extract

A GraphRAG implementation with a Vocabulary system to optimise AI integration - TjWheeler/deep-memory

GitHub - TjWheeler/deep-memory: A GraphRAG implementation with a Vocabulary system to optimise AI integration · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
TjWheeler
/
deep-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
71 Commits 71 Commits .changeset .changeset .github .github database database docs docs exports exports index-content/ person index-content/ person index-starterkits index-starterkits packages packages scripts scripts .env.example .env.example .gitignore .gitignore .mcp.json.example .mcp.json.example CLA.md CLA.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.indexer.yml docker-compose.indexer.yml docker-compose.neo4j.yml docker-compose.neo4j.yml docker-compose.sqlserver.yml docker-compose.sqlserver.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml quickstart-claude-desktop.md quickstart-claude-desktop.md quickstart-cosmosdb.md quickstart-cosmosdb.md quickstart-embeddings.md quickstart-embeddings.md quickstart-indexer.md quickstart-indexer.md quickstart-inmemory.md quickstart-inmemory.md quickstart-neo4j.md quickstart-neo4j.md quickstart-sqlserver.md quickstart-sqlserver.md tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
A vocabulary-driven graph memory library for AI agents. Give your AI structured, persistent memory that it can read, write, and traverse with zero friction.
AI agents working with knowledge graphs face a cold-start problem: before storing a single fact, they need to understand what entity types exist, what properties each type has, what relationship types connect them, and what constraints apply. Without this, every interaction starts with guesswork, inconsistency, and wasted tokens.
Deep-memory solves this with vocabulary governance . When an agent opens a repository, the first thing it sees is the vocabulary: a complete schema of entity types, relationship types, property definitions, and relationship constraints. This means:
No guesswork. The agent knows exactly what types exist and what properties to set. It can create an entity in a single tool call.
No inconsistency. Validation catches invalid types and missing required properties before they enter the graph. Two agents working on the same repository produce structurally consistent data.
No friction when extending. In open governance mode, the agent proposes a new type and it is auto-approved with deduplication. The agent never stops working to wait for approval.
Efficient token usage. The vocabulary is a compact schema, not thousands of example documents. An agent can internalize the full vocabulary in a few hundred tokens and then work confidently.
The result: an AI agent can go from opening a repository to creating entities and traversing relationships in seconds, with no training data, no few-shot examples, and no trial-and-error.
A note on getting started quickly
This repo has been designed to provide detailed documentation that AI can use. The AI can help guide the setup and configuration and get you started.
Pick a quickstart and paste its prompts into Claude Code:
quickstart-inmemory.md — zero setup; the graph lives in memory. Fastest 2-minute path.
quickstart-sqlserver.md — persistent storage in SQL Server. Bundled Docker Compose for local use.
quickstart-cosmosdb.md — Azure CosmosDB Gremlin API (local emulator on Windows, or an Azure account).
quickstart-neo4j.md — Neo4j graph database. Bundled Docker Compose, or AuraDB / self-hosted.
quickstart-indexer.md — build your own knowledge graph from source documents via the indexing pipeline.
Not a developer? Using Claude Desktop instead of Claude Code or another MCP capable AI?
quickstart-claude-desktop.md — no coding, no git, no build. Three installers + one config file. ~15 minutes.
Package
Path
Description
@utaba/deep-memory
packages/core
Core library — zero runtime deps, dual CJS/ESM via tsup
@utaba/deep-memory-embeddings-openai
packages/embeddings-openai
OpenAI-compatible embeddings provider (vLLM, OpenAI, Azure, Ollama)
@utaba/deep-memory-storage-sqlserver
packages/storage-sqlserver
SQL Server storage provider — persistent, multi-tenant graph storage
@utaba/deep-memory-storage-cosmosdb
packages/storage-cosmosdb
CosmosDB Gremlin storage + graph traversal provider
@utaba/deep-memory-storage-neo4j
packages/storage-neo4j
Neo4j storage + graph traversal provider (Cypher over Bolt, Community Edition)
@utaba/deep-memory-indexer
packages/indexer
Document indexing pipeline — LLM extraction, validation, consolidation, import
@utaba/deep-memory-indexer-llm-anthropic
packages/indexer-llm-anthropic
Anthropic LLM provider for indexer — prompt caching + native Messages API
@utaba/deep-memory-indexer-mcp-server
packages/indexer-mcp-server
Indexer MCP server — 9 phase-aware tools for the indexing pipeline
@utaba/deep-memory-local-mcp-server
packages/mcp-server
Local MCP server exposing deep-memory memory tools for AI agents
Deep-Memory Library
The core @utaba/deep-memory library provides a vocabulary-driven knowledge graph with typed entities, relationships, property validation, and pluggable storage/search/embedding providers. Zero runtime dependencies. The local MCP server ( @utaba/deep-memory-local-mcp-server ) exposes 20+ tools for AI agents to create repositories, manage entities and relationships, traverse graphs, and perform semantic search. It serves both as a reference integration and a production-ready interface for testing and working with knowledge graphs.
See docs/README.md for detailed architecture, provider setup, and API documentation.
Note: The Indexing Pipeline is optional. You can use and populate a memory graph using the AI without it. The indexing pipeline is most useful when you have documents you want to import.
The indexer packages ( @utaba/deep-memory-indexer , @utaba/deep-memory-indexer-mcp-server , @utaba/deep-memory-indexer-llm-anthropic ) provide an AI-driven pipeline for transforming source documents into structured knowledge graphs. The pipeline handles document analysis, multi-model LLM extraction, quality validation, cross-document entity deduplication, repository import, and embedding generation. The indexer MCP server exposes 9 phase-aware tools that guide an AI agent through the entire process — from scanning source documents to a query-ready knowledge graph.
Multi-worker extraction (local + cloud models in parallel)
Three-tier validation (structural, LLM-verified, human-reviewed)
Cost tracking and estimation at every phase
Starter kits for common domains
Progressive chapter extraction for large documents
Prompt caching with the Anthropic provider (up to 88% savings on repeated system prompts)
See docs/README.md for the full indexing pipeline documentation, including quickstart, model selection, and troubleshooting.
Starter kits are pre-built vocabularies for common domains. Each kit defines entity types, relationship types, property schemas, and an indexing process so the agent can begin populating a knowledge graph immediately.
README.md — Overview (purpose, use cases, governance) + manual indexing process via MCP tools
vocabulary.md — Entity types and relationship types with property definitions
domain-guidance.md — Domain-specific extraction guidance for AI agents (consumed by the automated pipeline)
To create a repository from a starter kit, tell your AI agent:
"Create a new deep memory repository using the Person starter kit."
Starter kits are extensible. In open governance mode, the agent can propose new entity and relationship types at any time.
Node.js 22 or 24 (the supported LTS pair — CI tests both)
The repository ships per-service Compose files so you only start what you need:
# Start SQL Server (required for persistent storage)
docker compose -f docker-compose.sqlserver.yml up -d
This starts a SQL Server 2022 Developer Edition container on port 1435 . The database schema is created automatically when the MCP server connects for the first time.
If you have an NVIDIA GPU and want semantic search (concept search, vocabulary deduplication via embeddings), also start the embeddings service:
# Start the vLLM embeddings server
docker compose -f docker-compose.indexer.yml --profile embeddings up -d
The vLLM service serves the Qwen3-Embedding-8B model on port 8010 . This is optional — without it, deep-memory falls back to string similarity for vocabulary deduplication, and memory_search_by_concept is unavailable.
pnpm install
pnpm build
3. Create the Database
Once SQL Server is healthy, connect with your preferred SQL tool and create a database (e.g. deep-memory ). The table schema is created automatically by the storage provider on first connection.
The MCP server is a local process that exposes deep-memory as Model Context Protocol tools over stdio. It is intended to run locally alongside your AI agent (Claude Code, Claude Desktop, etc.) to let the agent create and work with graph memory repositories. It also serves as a reference integration for embedding deep-memory into your own applications.
Add the following to your .mcp.json (project root or ~/.mcp.json for global):
{
"mcpServers" : {
"deep-memory" : {
"command" : " node " ,
"args" : [ " <path-to-repo>/packages/mcp-server/dist/index.js " ],
"env" : {
"DEEP_MEMORY_ACTOR_ID" : " mcp-agent " ,
"DEEP_MEMORY_ACTOR_TYPE" : " agent " ,
"DEEP_MEMORY_STORAGE" : " sqlserver " ,
"DEEP_MEMORY_SQL_HOST" : " localhost " ,
"DEEP_MEMORY_SQL_PORT" : " 1435 " ,
"DEEP_MEMORY_SQL_DATABASE" : " deep-memory " ,
"DEEP_MEMORY_SQL_USER" : " sa " ,
"DEEP_MEMORY_SQL_PASSWORD" : " DeepMem@Dev1234 " ,
"DEEP_MEMORY_SQL_SCHEMA" : " dbo " ,
"DEEP_MEMORY_SQL_TRUST_CERT" : " true " ,
"DEEP_MEMORY_EMBEDDINGS_BASE_URL" : " http://localhost:8010 " ,
"DEEP_MEMORY_EMBEDDINGS_MODEL" : " Qwen/Qwen3-Embedding-8B "
}
}
}
}
Security Note: Consider changing the SQL Password in the config and the docker server setttings.
Variable
Default
Description
DEEP_MEMORY_ACTOR_ID
mcp-agent
Actor ID stamped on provenance metadata
DEEP_MEMORY_ACTOR_TYPE
agent
Actor type: agent , human , or system
DEEP_MEMORY_STORAGE
memory
Storage backend: memory (in-memory, data lost on restart), sqlserver , cosmosdb , or neo4j
DEEP_MEMORY_SQL_HOST
—
SQL Server hostname
D

[truncated]

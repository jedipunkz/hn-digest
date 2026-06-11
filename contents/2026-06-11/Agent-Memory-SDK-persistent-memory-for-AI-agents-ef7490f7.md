---
source: "https://github.com/gharibyan/agent-memory"
hn_url: "https://news.ycombinator.com/item?id=48494184"
title: "Agent Memory SDK – persistent memory for AI agents"
article_title: "GitHub - gharibyan/agent-memory: TypeScript SDK for building AI agents with automatic, scoped, persistent memory. `agent-memory` wraps model calls with memory recall and learning so apps can keep useful user, thread, and operation context without manually stuffing long chat histories into every prom\n[truncated]"
author: "xgharibyan"
captured_at: "2026-06-11T19:05:41Z"
capture_tool: "hn-digest"
hn_id: 48494184
score: 1
comments: 0
posted_at: "2026-06-11T18:09:51Z"
tags:
  - hacker-news
  - translated
---

# Agent Memory SDK – persistent memory for AI agents

- HN: [48494184](https://news.ycombinator.com/item?id=48494184)
- Source: [github.com](https://github.com/gharibyan/agent-memory)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T18:09:51Z

## Translation

タイトル: Agent Memory SDK – AI エージェント用の永続メモリ
記事のタイトル: GitHub - gharibyan/agent-memory: 自動、スコープ指定、永続メモリを備えた AI エージェントを構築するための TypeScript SDK。 「agent-memory」は、メモリの呼び出しと学習を使用してモデル呼び出しをラップするため、アプリは長いチャット履歴をすべてのプロムに手動で詰め込むことなく、有用なユーザー、スレッド、操作コンテキストを維持できます。
[切り捨てられた]
説明: スコープ指定された自動永続メモリを備えた AI エージェントを構築するための TypeScript SDK。 「agent-memory」は、メモリの呼び出しと学習を使用してモデル呼び出しをラップするため、アプリは長いチャット履歴をすべてのプロンプトに手動で詰め込むことなく、有用なユーザー、スレッド、操作コンテキストを維持できます。 - ガリビアン/エージェントメモリ

記事本文:
GitHub - gharibyan/agent-memory: 自動でスコープ指定された永続メモリを備えた AI エージェントを構築するための TypeScript SDK。 「agent-memory」は、メモリの呼び出しと学習を使用してモデル呼び出しをラップするため、アプリは長いチャット履歴をすべてのプロンプトに手動で詰め込むことなく、有用なユーザー、スレッド、操作コンテキストを維持できます。 · GitHub
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
別のアカウントでアカウントを切り替えました

bまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ガリビアン
/
エージェントメモリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット .github .github アプリ アプリ パッケージ パッケージ スクリプト スクリプト スキル/エージェント メモリ スキル/エージェント メモリ テスト test .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TypeScript SDK は、スコープ指定された自動永続メモリを備えた AI エージェントを構築するための SDK です。
Agent-memory-sdk は、メモリの呼び出しと学習を使用してモデル呼び出しをラップするため、アプリは長いチャット履歴をすべてのプロンプトに手動で詰め込むことなく、有用なユーザー、スレッド、操作コンテキストを維持できます。
npm インストール エージェント メモリ SDK
または、別のパッケージマネージャーを使用します。
pnpm エージェント メモリ SDK を追加
糸追加エージェントメモリSDK
ワークスペース開発の場合:
pnpmインストール
pnpmテスト
pnpm リント
pnpm パック:チェック
クイックスタート
import { createAgent , openai } from "agent-memory-sdk"
const エージェント = createAgent ( {
モデル：オープンアイ（「gpt-5」）
})
const result = エージェントを待ちます。生成 ( {
ユーザー ID : "user_123" 、
threadId : "thread_123" 、
操作 ID : "op_weekly_update" 、
メッセージ: [
{ role : "user" 、 content : "私は簡潔な週次レポートを好みます。" }
]
})
コンソール。ログ (結果のテキスト)
userId が省略された場合、メモリは共有のデフォルト スコープに保存されます。これはプロトタイプやシングルユーザー アプリに役立ちます。マルチユーザー アプリでは、 userId を渡します。
userId : 耐久性のあるユーザー メモリ分離。
orgId : 組織レベルのリコール。
スレッド ID : chat-l

ローカルコンテキスト。
OperationId : 複雑なワークフロー用のコンパクトなアクティブ操作コンテキスト。
Memory.contextBudget : トークンの肥大化を避けるために、注入されるメモリ コンテキストを制限します。
Memory.learn: false : 1 回の呼び出しの学習を無効にします。
Memory.recall: false : 単一の呼び出しのリコールを無効にします。
パブリック Agent-memory-sdk パッケージは、デフォルトでローカル JSON 永続化になります。
import { createAgent , openai } from "agent-memory-sdk"
const エージェント = createAgent ( {
モデル：オープンアイ（「gpt-5」）
})
デフォルトでは、これは .memory/memory.json に書き込みます。
import { createAgent , openai , sqliteMemory } from "agent-memory-sdk"
const エージェント = createAgent ( {
モデル：openai（「gpt-5」）、
メモリ : sqliteMemory ( {
パス: ".memory/memory.sqlite"
})
})
pgvector を使用した Postgres の場合:
import { createAgent , openai , postgresMemory } from "agent-memory-sdk"
const エージェント = createAgent ( {
モデル：openai（「gpt-5」）、
メモリ: {
ストア : postgresMemory ( {
接続文字列: プロセス。環境 。 DATABASE_URL 、
ベクトル寸法 : 1536
})
}
})
Postgres アダプタは、最初のメモリ操作の前に移行を自動的に実行します。デフォルトで pgvector 拡張機能を作成し、移行をバージョン追跡し、リレーショナル メモリ テーブルを作成し、ベクトル検索用の HNSW コサイン インデックスを追加します。データベース プロバイダーが拡張機能を個別に管理している場合は、データベースに pgvector をインストールし、 createExtension: false を渡します。
ファーストパーティ プロバイダーの統合は、agent-memory-sdk からエクスポートされ、プロバイダー SDK または文書化されたプロバイダー クライアント パスを内部的に使用します。 OpenAI アダプターは、公式の openai TypeScript SDK に依存しています。
import { openai } from "agent-memory-sdk"
const モデル = openai ( "gpt-5" )
Anthropic と Gemini は、agent-memory-sdk によってエクスポートされ、内部で公式 SDK を使用します。
import { anthropic , gemini } from "agent-memory-sdk"
人間的な定数

モデル = 人間 ( "人間モデル" )
const geminiModel = gemini ( "gemini-2.5-pro" )
xAI は、agent-memory-sdk によってもエクスポートされます。ここでは、xAI のデフォルトで文書化された OpenAI SDK 互換クライアント パスを使用します。
「agent-memory-sdk」から { xai } をインポートします
const モデル = xai ( "grok-4" )
OpenAI 互換ヘルパーは、互換性のあるチャット完了 API を公開するカスタム プロバイダーに対してのみ使用してください。
「agent-memory-sdk」から { openAISupport } をインポートします
const モデル = openAI互換性 ( {
モデル: "ディープシークチャット" 、
BaseURL : "https://api.deepseek.com/v1" ,
apiKey : プロセス。環境 。 DEEPSEEK_API_KEY
})
パッケージ
Agent-memory-sdk は唯一のパブリック npm パッケージです。これは、ランタイム、ストレージ アダプター、モデル プロバイダー アダプターを 1 つのインストールと 1 つのインポート サーフェスの背後にバンドルします。
リポジトリは依然として package/* の下に実装境界を保持します。
パッケージ/コア : ランタイム中立なエンジン、コントラクト、コンパイラ、取得、およびメモリ内ストア。
package/local : ローカル JSON 永続アダプター。
package/sqlite : 実際の SQLite 永続アダプター。
package/postgres : pgvector 移行を使用した Postgres 永続アダプター。
package/openai : OpenAI 公式 SDK アダプターに加え、OpenAI 互換のカスタム エンドポイントのサポート。
package/anthropic : Anthropic 公式 SDK アダプター。
package/gemini : Gemini 公式 SDK アダプター。
package/xai : 文書化された OpenAI SDK 互換クライアント パスを使用する xAI アダプター。
これらのワークスペース パッケージはプライベート ビルド ユニットです。これらは、パブリック パッケージのビルド中に package/agent-memory/dist/internal/* にコンパイルされ、個別に公開されることはありません。
pnpm --filter @agent-memory/playground dev
ローカル プレイグラウンドは、エコー モデルとローカル JSON メモリを使用します。
SQLite メモリを使用した実際の OpenAI 呼び出しの場合:
cp apps/openai-sqlite-demo/.env.example apps/openai-sqlite-demo/.env
pnpm --filter @agent-memory/openai-

sqlite-デモ開発者
apps/openai-sqlite-demo/.env またはサーバー環境で OPENAI_API_KEY を設定します。キーはサーバー側に残り、メモリはデフォルトで .memory/openai-demo.sqlite に保持されます。
どちらのサンプル アプリもリポジトリに対してプライベートであり、npm パッケージには含まれていません。
pnpm ビルド
pnpmテスト
pnpm リント
pnpm パック:チェック
パッケージのドライランでは、agent-memory-sdk アーティファクトのみを公開する必要があり、 apps/playground 、 .memory 、ローカル データベース、ログ、スクリーンショット、または生成された tarball を含めることはできません。
TypeScript SDK は、スコープ指定された自動永続メモリを備えた AI エージェントを構築するための SDK です。 「agent-memory」は、メモリの呼び出しと学習を使用してモデル呼び出しをラップするため、アプリは長いチャット履歴をすべてのプロンプトに手動で詰め込むことなく、有用なユーザー、スレッド、操作コンテキストを維持できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

TypeScript SDK for building AI agents with automatic, scoped, persistent memory. `agent-memory` wraps model calls with memory recall and learning so apps can keep useful user, thread, and operation context without manually stuffing long chat histories into every prompt. - gharibyan/agent-memory

GitHub - gharibyan/agent-memory: TypeScript SDK for building AI agents with automatic, scoped, persistent memory. `agent-memory` wraps model calls with memory recall and learning so apps can keep useful user, thread, and operation context without manually stuffing long chat histories into every prompt. · GitHub
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
gharibyan
/
agent-memory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits .github .github apps apps packages packages scripts scripts skills/ agent-memory skills/ agent-memory test test .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
TypeScript SDK for building AI agents with automatic, scoped, persistent memory.
agent-memory-sdk wraps model calls with memory recall and learning so apps can keep useful user, thread, and operation context without manually stuffing long chat histories into every prompt.
npm install agent-memory-sdk
Or with another package manager:
pnpm add agent-memory-sdk
yarn add agent-memory-sdk
For workspace development:
pnpm install
pnpm test
pnpm lint
pnpm pack:check
Quick Start
import { createAgent , openai } from "agent-memory-sdk"
const agent = createAgent ( {
model : openai ( "gpt-5" )
} )
const result = await agent . generate ( {
userId : "user_123" ,
threadId : "thread_123" ,
operationId : "op_weekly_update" ,
messages : [
{ role : "user" , content : "Remember that I prefer concise weekly reports." }
]
} )
console . log ( result . text )
If userId is omitted, memory is stored in the shared default scope. That is useful for prototypes and single-user apps. In multi-user apps, pass userId .
userId : durable user memory isolation.
orgId : organization-level recall.
threadId : chat-local context.
operationId : compact active-operation context for complex workflows.
memory.contextBudget : limits injected memory context to avoid token bloat.
memory.learn: false : disables learning for a single call.
memory.recall: false : disables recall for a single call.
The public agent-memory-sdk package defaults to local JSON persistence:
import { createAgent , openai } from "agent-memory-sdk"
const agent = createAgent ( {
model : openai ( "gpt-5" )
} )
By default this writes to .memory/memory.json .
import { createAgent , openai , sqliteMemory } from "agent-memory-sdk"
const agent = createAgent ( {
model : openai ( "gpt-5" ) ,
memory : sqliteMemory ( {
path : ".memory/memory.sqlite"
} )
} )
For Postgres with pgvector:
import { createAgent , openai , postgresMemory } from "agent-memory-sdk"
const agent = createAgent ( {
model : openai ( "gpt-5" ) ,
memory : {
store : postgresMemory ( {
connectionString : process . env . DATABASE_URL ,
vectorDimensions : 1536
} )
}
} )
The Postgres adapter runs migrations automatically before the first memory operation. It creates the pgvector extension by default, version-tracks migrations, creates relational memory tables, and adds an HNSW cosine index for vector search. If your database provider manages extensions separately, install pgvector in the database and pass createExtension: false .
First-party provider integrations are exported from agent-memory-sdk and use provider SDKs or documented provider client paths internally. The OpenAI adapter depends on the official openai TypeScript SDK:
import { openai } from "agent-memory-sdk"
const model = openai ( "gpt-5" )
Anthropic and Gemini are exported by agent-memory-sdk and use their official SDKs internally:
import { anthropic , gemini } from "agent-memory-sdk"
const anthropicModel = anthropic ( "anthropic-model" )
const geminiModel = gemini ( "gemini-2.5-pro" )
xAI is exported by agent-memory-sdk too. It uses the documented OpenAI SDK-compatible client path with xAI defaults:
import { xai } from "agent-memory-sdk"
const model = xai ( "grok-4" )
Use the OpenAI-compatible helper only for custom providers that expose a compatible chat completions API:
import { openAICompatible } from "agent-memory-sdk"
const model = openAICompatible ( {
model : "deepseek-chat" ,
baseURL : "https://api.deepseek.com/v1" ,
apiKey : process . env . DEEPSEEK_API_KEY
} )
Package
agent-memory-sdk is the only public npm package. It bundles the runtime, storage adapters, and model provider adapters behind one install and one import surface.
The repository still keeps implementation boundaries under packages/* :
packages/core : runtime-neutral engine, contracts, compiler, retrieval, and in-memory store.
packages/local : local JSON persistence adapter.
packages/sqlite : real SQLite persistence adapter.
packages/postgres : Postgres persistence adapter with pgvector migrations.
packages/openai : OpenAI official SDK adapter, plus OpenAI-compatible custom endpoint support.
packages/anthropic : Anthropic official SDK adapter.
packages/gemini : Gemini official SDK adapter.
packages/xai : xAI adapter using the documented OpenAI SDK-compatible client path.
Those workspace packages are private build units. They are compiled into packages/agent-memory/dist/internal/* during the public package build and are not published separately.
pnpm --filter @agent-memory/playground dev
The local playground uses an echo model and local JSON memory.
For a real OpenAI call with SQLite memory:
cp apps/openai-sqlite-demo/.env.example apps/openai-sqlite-demo/.env
pnpm --filter @agent-memory/openai-sqlite-demo dev
Set OPENAI_API_KEY in apps/openai-sqlite-demo/.env or in your server environment. The key stays server-side, and memory persists to .memory/openai-demo.sqlite by default.
Both example apps are private to the repository and are not included in npm packages.
pnpm build
pnpm test
pnpm lint
pnpm pack:check
Package dry-runs must only publish the agent-memory-sdk artifact and must not include apps/playground , .memory , local databases, logs, screenshots, or generated tarballs.
TypeScript SDK for building AI agents with automatic, scoped, persistent memory. `agent-memory` wraps model calls with memory recall and learning so apps can keep useful user, thread, and operation context without manually stuffing long chat histories into every prompt.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

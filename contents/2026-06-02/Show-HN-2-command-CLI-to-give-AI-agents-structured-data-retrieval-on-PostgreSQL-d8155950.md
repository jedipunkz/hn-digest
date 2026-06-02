---
source: "https://github.com/0xJaksun/lithium-core"
hn_url: "https://news.ycombinator.com/item?id=48357212"
title: "Show HN: 2-command CLI to give AI agents structured data retrieval on PostgreSQL"
article_title: "GitHub - 0xJaksun/lithium-core: Storage engine for AI agents to navigate, store, and retrieve structured data. PostgreSQL ltree, built-in versioning, scoped queries. · GitHub"
author: "0xJaksun"
captured_at: "2026-06-02T00:39:14Z"
capture_tool: "hn-digest"
hn_id: 48357212
score: 4
comments: 0
posted_at: "2026-06-01T14:21:21Z"
tags:
  - hacker-news
  - translated
---

# Show HN: 2-command CLI to give AI agents structured data retrieval on PostgreSQL

- HN: [48357212](https://news.ycombinator.com/item?id=48357212)
- Source: [github.com](https://github.com/0xJaksun/lithium-core)
- Score: 4
- Comments: 0
- Posted: 2026-06-01T14:21:21Z

## Translation

タイトル: Show HN: AI エージェントに PostgreSQL で構造化データの取得を提供する 2 つのコマンド CLI
記事のタイトル: GitHub - 0xJaksun/lithium-core: AI エージェントが構造化データを移動、保存、取得するためのストレージ エンジン。 PostgreSQL ltree、組み込みバージョニング、スコープ指定クエリ。 · GitHub
説明: AI エージェントが構造化データを移動、保存、取得するためのストレージ エンジン。 PostgreSQL ltree、組み込みバージョニング、スコープ指定クエリ。 - 0xJaksun/リチウムコア
HN テキスト: AI エージェントには類似性検索ではなく、構造化データが必要です。グラフ DB は高価で、ベクトル ストアは曖昧です。 Lithium は、PostgreSQL ltree 上のストレージ エンジンです。階層型、バージョン管理型、スコープ指定型のクエリ。 2 つのコマンド: npx @lithium-ai/kit init claude mcp add lithium -- npx @lithium-ai/kitserve エージェントは、既存の Postgres 上の構造化データを移動、保存、取得するためのツールを入手します。オープンソース、MIT。

記事本文:
GitHub - 0xJaksun/lithium-core: AI エージェントが構造化データを移動、保存、取得するためのストレージ エンジン。 PostgreSQL ltree、組み込みバージョニング、スコープ指定クエリ。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
0xジャクスン
/
リチウムコア
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
95 コミット 95 コミット .github/ workflows .github/ workflows パッケージ パッケージ .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml vitest.integration.config.ts vitest.integration.config.ts vitest.unit.config.ts vitest.unit.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが構造化データを移動、保存、取得するためのストレージ エンジン。階層型、バージョン管理型、スコープ付き。 Postgres 上で実行されます。
npx @lithium-ai/kit init
クロード mcp リチウムを追加 -- npx @lithium-ai/kitserve
AI エージェントには正確な構造化データが必要です。 「Xに似ている」ではありません。あいまい類似検索ではありません。データが増大するにつれて速度が低下する、高価ではないグラフ走査。
エージェントが「engineering.auth の下にあるものをすべて教えてください」と尋ねた場合、それはグラフウォークではなく、1 つのインデックス付きルックアップである必要があります。 PostgreSQL の ltree はまさにそれを行います。 Lithium は、組み込みのバージョニングとスコープ指定された取得を備えた TypeScript API でこれをラップします。
npx @lithium-ai/kit init --adapter postgres
pnpmインストール
psql $LITHIUM_DATABASE_URL -f リチウム/スキーマ.sql
クロード mcp リチウムを追加 -- npx @lithium-ai/kitserve
完了しました。 Claude Code は、構造化データの移動、保存、取得ができるようになりました。
npm install @lithium-ai/core @lithium-ai/postgres @lithium-ai/mcp postgres
"@lithium-ai/core" から { Lithium } をインポートします。
"@lithium-ai/postgres" から {postgresAdapter } をインポートします。
"@lithium-ai/mcp" から {serveMcp } をインポートします。
"postgres" から postgres をインポートします。
const sql = postgres(プロセス.環境.LITHIUM_DATABASE_URL!);
const lithium = new Lithium (postgresAdapter(sql)) ;
サーブMcp（リチウム）

;
得られるもの
// 階層を構築します
リチウムを待ってください。クラスター。 create ( { name : "エンジニアリング" } ) ;
リチウムを待ってください。クラスター。 create ( { name : "データベース" ,parentPath : "エンジニアリング" } ) ;
リチウムを待ってください。クラスター。 create ( { name : "auth" ,parentPath : "engineering" } ) ;
// バージョン管理されたエントリを保存する
const エントリ = リチウムを待ちます。エントリー。 create({clusterId:cluster.id});
リチウムを待ってください。エントリー。 update({id:エントリ.値.エントリ.id}); // v2 は自動的に
// スコープ指定された取得: 「エンジニアリング」の下にあるすべて
const context = リチウムを待ちます。 getContext ( { パス : "エンジニアリング" } ) ;
// そのパスにあるすべてのクラスター、エントリ、バージョンを返します
すべてのメソッドは Result<T, E> を返します。スローされた例外はありません。 TypeScript は、各メソッドがどのエラーを返す可能性があるかを正確に示します。
クロード mcp リチウムを追加 -- npx @lithium-ai/kitserve
カーソル / ウィンドサーフィン / 任意の MCP クライアント
{
"mcpサーバー": {
"リチウム" : {
"コマンド" : " npx " ,
"args" : [ " @lithium-ai/kit " , "serve " ],
"cwd" : " /path/to/your/project "
}
}
}
AI ツールには次の 4 つの MCP ツールが追加されます。
エントリは純粋な構造です。コンテンツは独自のテーブルに存在し、エントリ バージョン ID によって参照されます。
クラスター
id、parentId、パス (「engineering.database」)、名前、説明
エントリー
id、クラスターID
エントリーバージョン
id、entryId、version (自動インクリメント)
コンテンツテーブル
entryVersionId (FK)、...任意のもの
データを保存および取得するためのコンテンツ読み取りおよび書き込みコールバックを追加します。
const リチウム = 新しいリチウム (postgresAdapter (sql) , {
読み取り: async ( versionIds ) => {
const rows = await SQL `
SELECT エントリ バージョン ID、データ
コンテンツから
WHERE エントリ バージョン ID = ANY( ${ バージョン ID } )
` ;
return new Map (rows.map((r)=>[r.entry_version_id,r.data]));
} 、
書き込み: async ( versionId , content ) => {
SQL を待つ `INSERT INTO 続き

ent (entry_version_id, data) VALUES ( ${ versionId } , ${ sql . json ( content ) } )` ;
コンテンツを返します。
} 、
} ) ;
パッケージ
パッケージ
何
@リチウム-ai/キット
CLI ツールボックス。 init +serve は 2 つのコマンドで実行されます。
@リチウム-ai/コア
ストレージエンジン。ランタイム依存度ゼロ。
@lithium-ai/postgres
ltree を備えた PostgreSQL アダプター。
@リチウム愛/霧雨
霧雨ORMアダプター。
@リチウムai/mcp
AIツール用のMCPサーバー。
移行
npx @lithium-ai/kit init # lithium/schema.sql を生成します
psql $LITHIUM_DATABASE_URL -f リチウム/スキーマ.sql
霧雨の場合:
"@lithium-ai/drizzle" から { クラスター , エントリ , エントリバージョン } をエクスポートします。
Npx ドリズルキット プッシュ
API
方法
何
create({ 名前, 親パス?, 説明? })
クラスターの作成、親の解決
findByPath({ パス })
ドットパスで検索
リスト()
すべてのクラスターをパス順に並べる
listDescendantIds({ パス })
ltree サブツリー クエリ
エントリー
方法
何
create({ クラスターID })
新しいエントリ + バージョン 1
更新({ id })
自動インクリメントバージョン
get({ ID, バージョン? })
エントリ + バージョン (最新または特定)
list({ クラスタ ID })
クラスタIDによるエントリ
listWithlatestVersion({clusterIds})
エントリ + 最新バージョン (バッチ)
コンテキスト
方法
何
getContext({ パス })
オプションのコンテンツリゾルバーによる範囲指定された取得
すべてのメソッドは Promise<Result<T, E>> を返します。
CLI ツールボックス ( @lithium-ai/kit )
統合テスト (テストコンテナ)
AI エージェント データ レイヤー (構造化検索、スコープ指定されたクエリ)
チーム全体での意思決定の追跡
続きを読む: メモリ グラフはスケールしない
AI エージェントが構造化データを移動、保存、取得するためのストレージ エンジン。 PostgreSQL ltree、組み込みバージョニング、スコープ指定クエリ。
www.npmjs.com/package/@lithium-ai/core
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
v0.0.5 — トランザクションサポートによるアトミックエントリ作成
最新の
2026 年 5 月 29 日
+ 1 リリース
パッケージ

0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Storage engine for AI agents to navigate, store, and retrieve structured data. PostgreSQL ltree, built-in versioning, scoped queries. - 0xJaksun/lithium-core

AI agents need structured data, not similarity search. Graph DBs are expensive, vector stores are fuzzy. Lithium is a storage engine on PostgreSQL ltree. Hierarchical, versioned, scoped queries. Two commands: npx @lithium-ai/kit init claude mcp add lithium -- npx @lithium-ai/kit serve Your agents get tools to navigate, store, and retrieve structured data on your existing Postgres. Open source, MIT.

GitHub - 0xJaksun/lithium-core: Storage engine for AI agents to navigate, store, and retrieve structured data. PostgreSQL ltree, built-in versioning, scoped queries. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
0xJaksun
/
lithium-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
95 Commits 95 Commits .github/ workflows .github/ workflows packages packages .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml vitest.integration.config.ts vitest.integration.config.ts vitest.unit.config.ts vitest.unit.config.ts View all files Repository files navigation
The storage engine for AI agents to navigate, store, and retrieve structured data. Hierarchical, versioned, scoped. Runs on your Postgres.
npx @lithium-ai/kit init
claude mcp add lithium -- npx @lithium-ai/kit serve
AI agents need exact, structured data. Not "similar to X". Not fuzzy similarity search. Not expensive graph traversal that slows down as your data grows.
When an agent asks "give me everything under engineering.auth ", that should be one indexed lookup, not a graph walk. PostgreSQL's ltree does exactly that. Lithium wraps it in a TypeScript API with built-in versioning and scoped retrieval.
npx @lithium-ai/kit init --adapter postgres
pnpm install
psql $LITHIUM_DATABASE_URL -f lithium/schema.sql
claude mcp add lithium -- npx @lithium-ai/kit serve
Done. Claude Code can now navigate, store, and retrieve structured data.
npm install @lithium-ai/core @lithium-ai/postgres @lithium-ai/mcp postgres
import { Lithium } from "@lithium-ai/core" ;
import { postgresAdapter } from "@lithium-ai/postgres" ;
import { serveMcp } from "@lithium-ai/mcp" ;
import postgres from "postgres" ;
const sql = postgres ( process . env . LITHIUM_DATABASE_URL ! ) ;
const lithium = new Lithium ( postgresAdapter ( sql ) ) ;
serveMcp ( lithium ) ;
What You Get
// Build a hierarchy
await lithium . clusters . create ( { name : "engineering" } ) ;
await lithium . clusters . create ( { name : "database" , parentPath : "engineering" } ) ;
await lithium . clusters . create ( { name : "auth" , parentPath : "engineering" } ) ;
// Store versioned entries
const entry = await lithium . entries . create ( { clusterId : cluster . id } ) ;
await lithium . entries . update ( { id : entry . value . entry . id } ) ; // v2 automatically
// Scoped retrieval: everything under "engineering"
const context = await lithium . getContext ( { path : "engineering" } ) ;
// Returns all clusters, entries, and versions under that path
Every method returns Result<T, E> . No thrown exceptions. TypeScript tells you exactly which errors each method can return.
claude mcp add lithium -- npx @lithium-ai/kit serve
Cursor / Windsurf / Any MCP Client
{
"mcpServers" : {
"lithium" : {
"command" : " npx " ,
"args" : [ " @lithium-ai/kit " , " serve " ],
"cwd" : " /path/to/your/project "
}
}
}
Your AI tools get four MCP tools:
Entries are pure structure. Your content lives in your own tables, referenced by entry version IDs.
Cluster
id, parentId, path ("engineering.database"), name, description
Entry
id, clusterId
EntryVersion
id, entryId, version (auto-incremented)
Your Content Table
entryVersionId (FK), ...whatever you want
Add content read and write callbacks to store and retrieve your data:
const lithium = new Lithium ( postgresAdapter ( sql ) , {
read : async ( versionIds ) => {
const rows = await sql `
SELECT entry_version_id, data
FROM content
WHERE entry_version_id = ANY( ${ versionIds } )
` ;
return new Map ( rows . map ( ( r ) => [ r . entry_version_id , r . data ] ) ) ;
} ,
write : async ( versionId , content ) => {
await sql `INSERT INTO content (entry_version_id, data) VALUES ( ${ versionId } , ${ sql . json ( content ) } )` ;
return content ;
} ,
} ) ;
Packages
Package
What
@lithium-ai/kit
CLI toolbox. init + serve in two commands.
@lithium-ai/core
Storage engine. Zero runtime deps.
@lithium-ai/postgres
PostgreSQL adapter with ltree.
@lithium-ai/drizzle
Drizzle ORM adapter.
@lithium-ai/mcp
MCP server for AI tools.
Migrations
npx @lithium-ai/kit init # generates lithium/schema.sql
psql $LITHIUM_DATABASE_URL -f lithium/schema.sql
With Drizzle:
export { clusters , entries , entryVersions } from "@lithium-ai/drizzle" ;
npx drizzle-kit push
API
Method
What
create({ name, parentPath?, description? })
Create cluster, resolve parent
findByPath({ path })
Find by dot-path
list()
All clusters ordered by path
listDescendantIds({ path })
ltree subtree query
Entries
Method
What
create({ clusterId })
New entry + version 1
update({ id })
Auto-increment version
get({ id, version? })
Entry + version (latest or specific)
list({ clusterIds })
Entries by cluster IDs
listWithLatestVersion({ clusterIds })
Entries + latest versions (batch)
Context
Method
What
getContext({ path })
Scoped retrieval with optional content resolver
All methods return Promise<Result<T, E>> .
CLI toolbox ( @lithium-ai/kit )
Integration tests (testcontainers)
AI agent data layer (structured retrieval, scoped queries)
Decision tracking across teams
Read more: Memory Graphs Don't Scale
Storage engine for AI agents to navigate, store, and retrieve structured data. PostgreSQL ltree, built-in versioning, scoped queries.
www.npmjs.com/package/@lithium-ai/core
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
v0.0.5 — Atomic entry creation with transaction support
Latest
May 29, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

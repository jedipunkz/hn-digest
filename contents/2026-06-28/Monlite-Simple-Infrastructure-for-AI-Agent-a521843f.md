---
source: "https://github.com/qataruts/monlite"
hn_url: "https://news.ycombinator.com/item?id=48710233"
title: "Monlite Simple Infrastructure for AI Agent"
article_title: "GitHub - qataruts/monlite: Documents, vectors, cache, queue, and cron in one SQLite file. The local backend for AI agents. · GitHub"
author: "emadjumaah"
captured_at: "2026-06-28T19:30:02Z"
capture_tool: "hn-digest"
hn_id: 48710233
score: 1
comments: 0
posted_at: "2026-06-28T18:48:56Z"
tags:
  - hacker-news
  - translated
---

# Monlite Simple Infrastructure for AI Agent

- HN: [48710233](https://news.ycombinator.com/item?id=48710233)
- Source: [github.com](https://github.com/qataruts/monlite)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T18:48:56Z

## Translation

タイトル: AI エージェント向け Monlite シンプル インフラストラクチャ
記事のタイトル: GitHub - qataruts/monlite: ドキュメント、ベクター、キャッシュ、キュー、および cron を 1 つの SQLite ファイルにまとめました。 AI エージェントのローカル バックエンド。 · GitHub
説明: 1 つの SQLite ファイル内のドキュメント、ベクター、キャッシュ、キュー、および cron。 AI エージェントのローカル バックエンド。 - カタール/モンライト

記事本文:
GitHub - qataruts/monlite: 1 つの SQLite ファイル内のドキュメント、ベクター、キャッシュ、キュー、および cron。 AI エージェントのローカル バックエンド。 · GitHub
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
カタール
/
モンライト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

イゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
75 コミット 75 コミット .github/ workflows .github/ workflows ベンチ ベンチ デモ デモ ドキュメント ドキュメント サンプル サンプル パッケージ パッケージ python python src src テスト テスト .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.mdライセンス ライセンス README.md README.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Docker スタックのスピンアップを停止します。ドキュメント、ベクター、全文検索、キャッシュ、キュー、
cron — 依存関係のない TypeScript コアを備えた 1 つの SQLite ファイル内。
"@monlite/core" から { createDb } をインポートします。
const db = createDb ( "./app.db" ) ;
それでおしまい。サーバーがありません。移行はありません。設定はありません。以下のものはすべて app.db に存在します。
📖 ドキュメント · 🎮 ライブデモ — ブラウザで実行する monlite · 💻 GitHub
ローカルスタック全体を 1 つのファイルに置き換えます
ほとんどのローカル アプリ、CLI、および AI エージェントは、同じ一連のサービスをやりくりします。モンライトはすべてを崩壊させる
それらを 1 つの .db ファイルにまとめます。
機能ごとに 1 つの npm インストール。すべてに対して 1 つの .db ファイル。バックアップ = cp app.db バックアップ.db 。
AI エージェント バックエンド — Docker なし
コーディング エージェント、RAG パイプライン、または自律型ワーカーは通常、メモリとして MongoDB を必要とします。
キャッシュとロックには Redis、セマンティック検索には Qdrant、タスク キューには BullMQ が使用されます。
monlite では、スタック全体が 1 つのファイルになります。
"@monlite/core" から { createDb } をインポートします。
"@monlite/kv" から { kv } をインポートします。
import { createVectorStore } から "@monlite/vector" ;
"@monlite/queue" から { createQueue } をインポートします。
インポート { createCron } f

rom "@monlite/cron" ;
const db = createDb ( "./agent.db" ) ;
// メモリ / 状態 — ドキュメント コレクション
const メモリ = db 。コレクション (「思い出」) ;
思い出を待ってください。 create ( { data : { AgentId : "a1" , content : "ユーザーはダーク モードを好みます" } } ) ;
// セマンティックリコール — 埋め込みに対するベクトル検索
const ストア = createVectorStore (db) ;
店 。 ensureCollection ( "memory" , { Dimensions : 384 , IndexedFields : [ "agentId" ] } ) ;
const 呼び出し = ストア。 search ( "memory" , { ベクトル : queryEmb , topK : 5 , where : { AgentId : "a1" } } ) ;
// 1 回限りのジョブ要求 - プロセス間の比較と交換
const 要求 = ジョブを待ちます。 findOneAndUpdate ( {
ここで: { ステータス : "保留中" 、タイプ : "要約" } 、
データ : { $set : { ステータス : "アクティブ" } 、 $inc : { バージョン : 1 } } 、
returnDocument : "後" 、
} ) ;
// 8 人の労働者がレースします。まさに1人が勝ちます。残りは null になります。
// キャッシュ + アトミック ロック — 存在しない場合に設定
const ロック = kv (db) ;
const 取得 = ロック 。 setNX ( "ロック:ジョブ:42" , 1 , { ttl : 30_000 } ) ; // true = あなたが所有しています
// 永続的なタスク キュー — 再試行、バックオフ、デッドレター
const queue = createQueue(db,{maxAttempts:3});
列 。 process ( "embed" , async ( job ) => await embed ( job . payload . text ) , { concurrency : 4 } ) ;
// スケジュールされた作業 — 再起動後も維持されます
const cron = createCron (db) ;
クロン。スケジュール ( "nightly-cleanup" , "0 3 * * *" , ( ) => queue .add ( "cleanup" , { } ) ) ;
ドッカーはありません。接続文字列を含む .env ファイルはありません。 Redis のセットアップがありません。 1 つのファイル、ノードserve.mjs 。
リアルタイムの反応性 — ローカル Firebase
collection.watch() は、関連する変更が発生した場合にのみ再発行されるライブ結果セットを提供します。
行レベルのマッチング、偽りの再レンダリングはありません。 @monlite/sync からの変更にも対応します。
// 初期スナップショット → その後、管理者が許可した場合にのみ再起動します

追加/変更/削除されます
const stop = ユーザー .時計（
{ ここで : { 役割 : { の値 : "admin" } } } 、
({ 結果 , 追加 , 削除 } ) => renderAdminList (結果 ) ,
) ;
@monlite/sync とペアリングすると、ローカル データベースが MongoDB のライブ レプリカになります。
PostgreSQL — 完全にオフライン対応、再接続時に同期します。
おもちゃではなく、適切なクエリ言語
monlite には Mongo/Prisma スタイルのクエリ API があります。型付きコレクションはコンパイル時にチェックされます
where / orderBy と select で絞り込む戻り値の型。
インターフェースの順序 {
customerId : 文字列 ;
アイテム : { sku : 文字列 ;数量: 番号 } [ ] ;
ステータス : "保留中" | 「発送済み」 | "戻った" ;
合計: 数値;
}
定数注文 = db 。コレクション < 注文 > ( "注文" ) ;
// elemMatch — オブジェクトの配列内でクエリを実行する
注文を待ちます。 findMany ( {
ここで、 { アイテム : { elemMatch : { sku : "WIDGET-A" 、数量 : { gte : 5 } } } 、
} ) ;
// Regex — 大文字と小文字を区別しないパターン マッチング
注文を待ちます。 findMany ( { ここで : { ステータス : { 正規表現 : "^pend" 、モード : "insensitive" } } } ) ;
// 集計パイプライン — GROUP BY、$lookup joins、$unwind
注文を待ちます。集計 ( [
{ $match : { ステータス : "発送済み" } } ,
{ $group : { _id : "$customerId" 、支出額 : { $sum : "$total" } } } 、
{ $sort : { 支出 : - 1 } } 、
{ $limit : 10 } 、
]);
// アトミックな非同期トランザクション — オールオアナッシングで内部で待機します
データベースを待ちます。 transactionAsync ( async ( tx ) => {
const account = await tx 。 findFirst ( { ここで : { _id : "acc-1" } } ) ;
if (アカウント . 残高 < 100 ) 新しいエラー (「資金不足」) をスローします。
TX を待ちます。 update ({ where : {_id : "acc-1" } , data : {$inc : {balance : - 100 } } } ) ;
TX を待ちます。 update ({ where : { _id : "acc-2" } , data : { $inc : { 残高 : + 100 } } } ) ;
} ) ;
ハイブリッド検索 — 1 回の呼び出しでキーワード + セマンティックを実行
両方の長所を活用: FTS5 k

逆数によるベクトル類似度と融合したキーワードランキング
ランクフュージョン。 1 つのクエリ、1 つのランク付けされたリスト。
"@monlite/fts" から { fts } をインポートします。
import { ベクトル , hybridSearch } から "@monlite/vector" ;
const db = createDb ( "./app.db" , {
allowExtensions : true 、
プラグイン: [
fts ( { docs : [ "タイトル" , "本文" ] } ) 、
ベクトル ({ ドキュメント : { フィールド : "embedding" 、次元 : 384 } } ) 、
]、
} ) ;
const Hist = await hybridSearch ( db . collection ( "docs" ) , {
テキスト : "機械学習の基礎" ,
ベクトル : await embed ( "機械学習の基礎" ) 、
トップK : 10 、
ここで、 { 公開 : true } 、
} ) ;
// 融合ランキング — 意味的に類似しており、かつキーワード関連の結果、最良のものから順
Python は同じファイルを読み取ります
monlite データベースは、文書化された規則を備えたプレーンな SQLite であるため、Python ポートは次のようになります。
そして同じ .db を書き込みます。 Python が取り込み、Node がサービスを提供するか、または任意の分割を行います。
monlite import create_db 、kv から
db = create_db ( "app.db" ) # ノードプロセスが使用するのとまったく同じファイル
ユーザー = データベース 。コレクション (「ユーザー」)
ユーザー。 create ({ "名前" : "アリ" , "年齢" : 30 , "タグ" : [ "管理者" ]})
ユーザー。 find_many ( where = { "タグ" : { "has" : "admin" }})
kv (db)。 set ( "セッション:42" , { "ユーザー" : "アリ" }, ttl = 60_000 )
1つのファイル。 2 つのランタイム。ゼロ翻訳レイヤー。
環境
どうやって
ノード22.5+
@monlite/core — 組み込みの node:sqlite を使用し、ネイティブ ビルドは不要です
ノード18/20
@monlite/core + better-sqlite3 — 存在する場合は自動選択されます
ブラウザ
@monlite/wasm — SQLite-WASM (sql.js) 上の同じ API
電子
@monlite/electron — メインプロセスの DB、IPC 上のレンダラーの同じ API
パイソン
pip install monlite — 同じ .db ファイル、純粋な stdlib コア
インストール
# ゼロ依存性: ノードの組み込みノード:sqlite を使用します (ノード >= 22.5)
npm インストール @monlite/core
# ノード 18/20 の場合、または実験的フラグを回避するには:
npmで

ストール @monlite/core better-sqlite3
必要に応じてパッケージを追加します。
npm install @monlite/vector # セマンティック検索
npm install @monlite/fts # 全文検索
npm install @monlite/kv # キャッシュ + ロック
npm install @monlite/queue # ジョブキュー
npm install @monlite/cron # スケジューラ
npm install @monlite/sync # クラウド同期 (MongoDB / PostgreSQL / MySQL)
npm install @monlite/wasm # ブラウザのサポート
なぜ使えないのか…
SQLite を直接?それは可能ですが、ドキュメント レイヤー、クエリ トランスレータ、
FTS 統合、ベクトル拡張配線、変更フィード、同期エンジン、およびすべての
TypeScript は自分で入力します。 monlite は、すでに完了し、テストされているその作業です。
MongoDB + Redis + Qdrant?ローカル/エッジ/デスクトップ/単一マシンでの作業には料金がかかります
1 つの問題を解決するために 3 つの個別のサービスの運用コストがかかります。 monliteはそれらすべてを入れます
1 つのファイル、1 つの API、インフラストラクチャなし。
ファイアベース/スーパーベース?共有クラウド状態に最適です。仕事が必要なときはあまり良くありません
オフライン、CLI ツールの配布、デスクトップ アプリの構築、またはデータをデバイス上に保持します。モンライトは地元第一主義です。
@monlite/sync は必要に応じてクラウド部分を処理します。
完全なガイドはドキュメント サイトにあります:
パッケージ:vector、fts、kv、queue、cron、sync、wasm
ガイド: AI エージェントのバックエンド、運用、移行、カスタム同期アダプター
参考：ファイル形式・Python・ベンチマーク
実行可能なデモは、examples/ にあります。
実稼働準備が整い、公開されました。現在のバージョン: @monlite/core 2.6.2 、@monlite/sync
1.3.0、@monlite/vector 0.5.1、@monlite/fts 0.5.0、@monlite/kv 0.2.0、
@monlite/queue 0.3.0、@monlite/cron 0.1.1、@monlite/wasm 0.2.2。 2.x API は凍結されています。
ベクトルとフルテキストのインデックス作成は大規模に線形です - 100K のベクトルを一括取り込みするか、
ドキュメントは高速に保たれるため (O(n²) 再インデックスなし)、快適に 1 をバックアップします。

0K ～ 100K ドキュメント RAG。
ライブデモでは、あらゆるパッケージ (ドキュメント、
フルテキスト (FTS5)、ベクトル/セマンティック検索、キャッシュ、キュー、および cron — で 100% 実行されます。
SQLite-WASM 上のブラウザ。Transformers.js 経由でデバイス上で計算されたセマンティック埋め込みを使用します。
Python ポート ( pip install monlite ) には現在、ドキュメントと kv が同梱されており、残りの部分は
パッケージ ファミリが進行中です。
ドキュメント、ベクター、キャッシュ、キュー、および cron が 1 つの SQLite ファイルに含まれています。 AI エージェントのローカル バックエンド。
読み込み中にエラーが発生しました。このページをリロードしてください。
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Documents, vectors, cache, queue, and cron in one SQLite file. The local backend for AI agents. - qataruts/monlite

GitHub - qataruts/monlite: Documents, vectors, cache, queue, and cron in one SQLite file. The local backend for AI agents. · GitHub
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
qataruts
/
monlite
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
75 Commits 75 Commits .github/ workflows .github/ workflows bench bench demo demo docs docs examples examples packages packages python python src src tests tests .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Stop spinning up Docker stacks. Documents, vectors, full-text search, cache, queue, and
cron — in one SQLite file, with a zero-dependency TypeScript core.
import { createDb } from "@monlite/core" ;
const db = createDb ( "./app.db" ) ;
That's it. No server. No migrations. No configuration. Everything below lives in app.db .
📖 Docs · 🎮 Live demo — monlite running in your browser · 💻 GitHub
Replace your entire local stack with one file
Most local apps, CLIs, and AI agents juggle the same set of services. monlite collapses all
of them into a single .db file:
One npm install per feature. One .db file for all of them. Backup = cp app.db backup.db .
The AI agent backend — without Docker
A coding agent, RAG pipeline, or autonomous worker typically needs MongoDB for memory,
Redis for cache and locks, Qdrant for semantic search, and BullMQ for the task queue.
With monlite, that entire stack is one file:
import { createDb } from "@monlite/core" ;
import { kv } from "@monlite/kv" ;
import { createVectorStore } from "@monlite/vector" ;
import { createQueue } from "@monlite/queue" ;
import { createCron } from "@monlite/cron" ;
const db = createDb ( "./agent.db" ) ;
// Memory / state — document collections
const memories = db . collection ( "memories" ) ;
await memories . create ( { data : { agentId : "a1" , content : "user prefers dark mode" } } ) ;
// Semantic recall — vector search over embeddings
const store = createVectorStore ( db ) ;
store . ensureCollection ( "memory" , { dimensions : 384 , indexedFields : [ "agentId" ] } ) ;
const recall = store . search ( "memory" , { vector : queryEmb , topK : 5 , where : { agentId : "a1" } } ) ;
// Exactly-once job claim — cross-process compare-and-swap
const claimed = await jobs . findOneAndUpdate ( {
where : { status : "pending" , type : "summarize" } ,
data : { $set : { status : "active" } , $inc : { version : 1 } } ,
returnDocument : "after" ,
} ) ;
// 8 workers race. Exactly one wins. The rest get null.
// Cache + atomic locks — set-if-absent
const lock = kv ( db ) ;
const acquired = lock . setNX ( "lock:job:42" , 1 , { ttl : 30_000 } ) ; // true = you own it
// Durable task queue — retries, backoff, dead-letter
const queue = createQueue ( db , { maxAttempts : 3 } ) ;
queue . process ( "embed" , async ( job ) => await embed ( job . payload . text ) , { concurrency : 4 } ) ;
// Scheduled work — persisted across restarts
const cron = createCron ( db ) ;
cron . schedule ( "nightly-cleanup" , "0 3 * * *" , ( ) => queue . add ( "cleanup" , { } ) ) ;
No Docker. No .env files with connection strings. No Redis setup. One file, node serve.mjs .
Real-time reactivity — local Firebase
collection.watch() delivers a live result set that re-emits only when a relevant change lands.
Row-level matching, no spurious re-renders. Works with changes from @monlite/sync too.
// Initial snapshot → then re-fires only when an admin is added/changed/removed
const stop = users . watch (
{ where : { roles : { has : "admin" } } } ,
( { results , added , removed } ) => renderAdminList ( results ) ,
) ;
Pair with @monlite/sync and your local database becomes a live replica of MongoDB or
PostgreSQL — fully offline capable, syncs when reconnected .
A proper query language — not a toy
monlite has a Mongo/Prisma-style query API. Typed collections get compile-time checked
where / orderBy and return types that narrow with select .
interface Order {
customerId : string ;
items : { sku : string ; qty : number } [ ] ;
status : "pending" | "shipped" | "returned" ;
total : number ;
}
const orders = db . collection < Order > ( "orders" ) ;
// elemMatch — query inside arrays of objects
await orders . findMany ( {
where : { items : { elemMatch : { sku : "WIDGET-A" , qty : { gte : 5 } } } } ,
} ) ;
// Regex — case-insensitive pattern matching
await orders . findMany ( { where : { status : { regex : "^pend" , mode : "insensitive" } } } ) ;
// Aggregation pipeline — GROUP BY, $lookup joins, $unwind
await orders . aggregate ( [
{ $match : { status : "shipped" } } ,
{ $group : { _id : "$customerId" , spent : { $sum : "$total" } } } ,
{ $sort : { spent : - 1 } } ,
{ $limit : 10 } ,
] ) ;
// Atomic async transactions — await inside, all-or-nothing
await db . transactionAsync ( async ( tx ) => {
const account = await tx . findFirst ( { where : { _id : "acc-1" } } ) ;
if ( account . balance < 100 ) throw new Error ( "insufficient funds" ) ;
await tx . update ( { where : { _id : "acc-1" } , data : { $inc : { balance : - 100 } } } ) ;
await tx . update ( { where : { _id : "acc-2" } , data : { $inc : { balance : + 100 } } } ) ;
} ) ;
Hybrid search — keyword + semantic in one call
Get the best of both worlds: FTS5 keyword ranking fused with vector similarity via Reciprocal
Rank Fusion. One query, one ranked list.
import { fts } from "@monlite/fts" ;
import { vector , hybridSearch } from "@monlite/vector" ;
const db = createDb ( "./app.db" , {
allowExtensions : true ,
plugins : [
fts ( { docs : [ "title" , "body" ] } ) ,
vector ( { docs : { field : "embedding" , dimensions : 384 } } ) ,
] ,
} ) ;
const hits = await hybridSearch ( db . collection ( "docs" ) , {
text : "machine learning fundamentals" ,
vector : await embed ( "machine learning fundamentals" ) ,
topK : 10 ,
where : { published : true } ,
} ) ;
// Fused ranking — semantically similar AND keyword-relevant results, best first
Python reads the same file
A monlite database is plain SQLite with documented conventions, so the Python port reads
and writes the same .db . Python ingests, Node serves — or any split you like.
from monlite import create_db , kv
db = create_db ( "app.db" ) # the exact same file your Node process uses
users = db . collection ( "users" )
users . create ({ "name" : "Ali" , "age" : 30 , "tags" : [ "admin" ]})
users . find_many ( where = { "tags" : { "has" : "admin" }})
kv ( db ). set ( "session:42" , { "user" : "ali" }, ttl = 60_000 )
One file. Two runtimes. Zero translation layer.
Environment
How
Node 22.5+
@monlite/core — uses built-in node:sqlite , zero native build
Node 18/20
@monlite/core + better-sqlite3 — auto-selected when present
Browser
@monlite/wasm — same API on SQLite-WASM (sql.js)
Electron
@monlite/electron — DB in main process, same API in renderer over IPC
Python
pip install monlite — same .db file, pure stdlib core
Install
# Zero-dependency: uses Node's built-in node:sqlite (Node >= 22.5)
npm install @monlite/core
# For Node 18/20, or to avoid the experimental flag:
npm install @monlite/core better-sqlite3
Add packages as you need them:
npm install @monlite/vector # semantic search
npm install @monlite/fts # full-text search
npm install @monlite/kv # cache + locks
npm install @monlite/queue # job queue
npm install @monlite/cron # scheduler
npm install @monlite/sync # cloud sync (MongoDB / PostgreSQL / MySQL)
npm install @monlite/wasm # browser support
Why not just use…
SQLite directly? You could — but you'd be writing the document layer, the query translator,
the FTS integration, the vector extension wiring, the change feed, the sync engine, and all the
TypeScript types yourself. monlite is that work, already done and tested.
MongoDB + Redis + Qdrant? For local / edge / desktop / single-machine work, you're paying
the operational cost of three separate services to solve one problem. monlite puts them all in
one file, with one API, and zero infrastructure.
Firebase / Supabase? Great for shared cloud state. Not so great when you need to work
offline, ship a CLI tool, build a desktop app, or keep data on-device. monlite is local-first;
@monlite/sync handles the cloud part when you need it.
Full guide at the documentation site :
Packages: vector · fts · kv · queue · cron · sync · wasm
Guides: AI-agent backend · production · migrations · custom sync adapters
Reference: file format · Python · benchmarks
Runnable demos are in examples/ .
Production-ready and published. Current versions: @monlite/core 2.6.2 , @monlite/sync
1.3.0, @monlite/vector 0.5.1 , @monlite/fts 0.5.0 , @monlite/kv 0.2.0,
@monlite/queue 0.3.0 , @monlite/cron 0.1.1, @monlite/wasm 0.2.2 . The 2.x API is frozen.
Vector and full-text indexing are linear at scale — bulk-ingesting 100K vectors or
documents stays fast (no O(n²) re-index), so it comfortably backs 10K–100K-document RAG.
The live demo showcases every package — documents,
full-text (FTS5), vector/semantic search , cache, queue, and cron — running 100% in the
browser on SQLite-WASM, with semantic embeddings computed on-device via Transformers.js.
The Python port ( pip install monlite ) currently ships documents + kv, with the rest of the
package family in progress.
Documents, vectors, cache, queue, and cron in one SQLite file. The local backend for AI agents.
There was an error while loading. Please reload this page .
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

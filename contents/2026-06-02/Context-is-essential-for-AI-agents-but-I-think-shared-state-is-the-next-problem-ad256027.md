---
source: "https://github.com/Abloatai/ablo"
hn_url: "https://news.ycombinator.com/item?id=48355849"
title: "Context is essential for AI agents, but I think shared state is the next problem"
article_title: "GitHub - Abloatai/ablo: State control for AI agents. Persist, coordinate, and audit every human and agent write to shared application state. · GitHub"
author: "lukasanderss"
captured_at: "2026-06-02T00:41:34Z"
capture_tool: "hn-digest"
hn_id: 48355849
score: 2
comments: 1
posted_at: "2026-06-01T12:19:05Z"
tags:
  - hacker-news
  - translated
---

# Context is essential for AI agents, but I think shared state is the next problem

- HN: [48355849](https://news.ycombinator.com/item?id=48355849)
- Source: [github.com](https://github.com/Abloatai/ablo)
- Score: 2
- Comments: 1
- Posted: 2026-06-01T12:19:05Z

## Translation

タイトル: AI エージェントにはコンテキストが不可欠ですが、状態の共有が次の問題だと思います
記事タイトル: GitHub - Abloatai/ablo: AI エージェントの状態管理。共有アプリケーションの状態に書き込むすべての人間とエージェントを永続化し、調整し、監査します。 · GitHub
説明: AI エージェントの状態制御。共有アプリケーションの状態に書き込むすべての人間とエージェントを永続化し、調整し、監査します。 - アブロアタイ/アブロ

記事本文:
GitHub - Abloatai/ablo: AI エージェントの状態制御。共有アプリケーションの状態に書き込むすべての人間とエージェントを永続化し、調整し、監査します。 · GitHub
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
アブロアタイ
/
アブロ
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows docs docs 例 例 スクリプト スクリプト src src .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス通知 NOTICE README.md README.md llms-full.txt llms-full.txt llms.txt llms.txt package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Ablo は、アプリの状態を共有するための型付き同期エンジンです。これは人間が行うようなものです。
サーバー コードと AI エージェントはすべて一度に編集します。
編集内容をリアルタイムではなく、どこにでも表示する必要がある場合に手を伸ばしてください。
お互いを黙って上書きし、誰が何を作業しているかを公開し、記録を残します
誰が何を変えたのか。
スキーマ -> ablo.<モデル>.create/retrieve/update/claim(...)
なぜアブロなのか
デフォルトではリアルタイム。すべての作成、更新、削除がファンアウトします
すべての加入者 (人間とエージェント) にデルタを個別に送信することなく確認
「マルチプレイヤーモード」をオンにします。
静かな泥棒はいません。書き込みは古い読み取りから保護されており、
低速読み取り→LLM→書き込みギャップにわたって行を保持するため、同時編集キュー
上書きする代わりに。
エージェント向けに構築されています。誰が編集中であるかを確認し (claimState / queue )、調整を行います。
コーディング エージェントが実際の API から統合されるように、llms.txt を同梱します。
端から端まで入力します。 Zod スキーマは型付きモデル プロキシを生成します
( ablo.<model>.update(...) )、オプティミスティックなローカル読み取り、およびリアクティブな React フック。
独自の認証とデータベースを持ち込んでください。 Ablo スコープのリアルタイム データを同期する
既存のアイデンティティからグループを作成し、データベースをソースとして残すことができます
データソースを介して真実を伝えます。
目的: 共同作業

ディター、AI エージェントのワークフロー、内部ツールなど
複数のアクターが共有状態を変更するアプリであり、誰もがそれをライブで見る必要があります。
npm インストール @abloatai/ablo
キーとランタイム。 Ablo には Node 22 以降と TypeScript 5 以降が必要です。 sk_test_* を取得する
サンドボックスのキー
(エクスポート ABLO_API_KEY=sk_test_... );キーは信頼できるサーバー ランタイムにのみ保持します。
ブラウザーでは、<AbloProvider> がサインインしているユーザーの認証を行います。
セッション — 決して生のキーではありません。
次に、手動で配線します。以下のクイック スタートは、次の形状です。
コピーする。本番環境 (React、既存のバックエンド、データ ソース、エージェント) の場合、
統合ガイドはさらに深いマップです。
エージェントに配線させたほうがいいですか?パッケージには、正確な llms.txt が同梱されています。
API のマップ — クロード コードまたはカーソルが実際のサーフェスから統合されるようにする
推測する代わりに:
node_modules/@abloatai/ablo/llms.txt を読み取り、Ablo スキーマ、 <AbloProvider> を追加し、最初の作成/取得/更新を追加します。
'@abloatai/ablo' から Ablo をインポートします。
import {defineSchema , model , z } from '@abloatai/ablo/schema' ;
const スキーマ = 定義スキーマ ( {
天気レポート : モデル ( {
場所：z。文字列 ( ) 、
ステータス：z。 enum ( [ 'pending' , 'ready' ] ) ,
予想：z。文字列 ( ) 。オプション ( ) 、
} )、
} ) ;
const ablo = アブロ ( {
スキーマ 、
apiKey: プロセス。環境 。 ABLO_API_KEY 、
} ) ;
待ってください。準備ができて （ ） ;
const created = ablo を待ちます。天気予報。作成 ( {
場所 : 「ストックホルム」 、
ステータス : '保留中' 、
} ) ;
const updated = ablo を待ちます。天気予報。 update (作成された .id , {
ステータス : '準備完了' 、
予報：「小雨、13℃」、
} ) ;
コンソール。 log({id:更新.id,ステータス:更新.ステータス});
待ってください。破棄() ;
期待される出力:
{ ID: '...'、ステータス: '準備完了' }
スキーマを渡して、 ablo.weatherReports.update(...) のような型付きモデルを取得します。
retrieve(id) はローカルから 1 行を返します。

キャッシュ — 同期、ラウンドトリップなし。
list(...) は、すでに同期されているものをフィルターして並べ替えます。それは同期でもあり、
useAblo/subscribe の下でリアクティブです。 load(...) は、次のときにサーバーからフェッチします。
行はまだローカルではない可能性があります。
アブロ。天気予報。取得 ( 'レポート_ストックホルム' ) ;
const pending = ablo 。天気予報。リスト ( {
ここで: { ステータス : '保留中' } 、
orderBy : { 場所 : 'asc' } 、
制限: 20、
} ) ;
const Ready = ablo を待ちます。天気予報。ロード ( {
ここで: { ステータス : '準備完了' } 、
タイプ: '完全' 、
} ) ;
ここでの配列値は IN を意味します。ロード時に、次のように入力します: 'complete' waits for
サーバー。 「unknown」は、現在ローカルにあるものを返し、バックグラウンドで更新します。
create / update は楽観的に適用され、行に解決されます。 2 つのオプション
日々のこと：
待ってください。天気予報。 update ( id , { status : 'ready' } , { wait : 'confirmed' } ) ;
自分の下で変更された行に対する書き込みを保護するには、readAt + onStale を渡します
— 長期にわたるエージェントの作業の調整を参照してください。
エージェントは行を読み取り、30 秒間考え、書き戻し、変更されたものはすべて上書きします
その間、あるいはさらに悪いことに、古い状態に作用します。クレームはそのギャップをまたいで行を保持します。
待ってください。天気予報。 claim ( 'report_stockholm' , async ( report ) => {
const予報=weatherAgentを待ちます。 getWeather (レポート.場所) ;
待ってください。天気予報。 update (レポート . id , { 予測 , ステータス : '準備完了' } ) ;
} ) ;
他の誰かが行を保持している場合、claim() は公平なキューで待機してから再読み取りします。
したがって、レポートは現在の行であり、古いスナップショットではありません。読み取りは次までオープンのままになります
デフォルト。行にのみ作用してシリアル化されます。コールバックが行われるとクレームは解放されます
返すか投げる。
行動する前に誰が編集中であるかを確認し、待つかスキップするかを決定します。
アブロ。天気予報。 claimState ( 'レポート_ストックホルム' ) ;
アブロ。天気予報。キュー ( 'rep

ort_ストックホルム' ) ;
待ってください。天気予報。クレーム ( id 、非同期 (レポート) => {
/* 保持されている作業を実行します */
} , {待機: false } ) ;
待ってください。天気予報。クレーム ( id 、非同期 (レポート) => {
/* 保持されている作業を実行します */
} , { maxQueueDepth : 2 } ) ;
claimState は、ホルダー (または null ) を返します。キューは待機中のラインを返します
その後ろに。 wait: false は、行が保持されているときに待機せずにスキップします。
maxQueueDepth: 2 人以上がすでに先行している場合、2 つのベイル。
デフォルトの読み取りは、行が要求されている間も機能し続けます。サーバーは要求が必要な読み取りを行います
セマンティクスは ifClaimed: 'return' | を使用してオプトインできます。 '待って' | '失敗' 。
要求されていない書き込みであっても、古い推論に到達することはできません。コミットは保護されています。
{を試してください
待ってください。天気予報。 update ( id , { status : 'ready' } , { readAt , onStale : 'reject' } ) ;
} キャッチ (e) {
if ( einstanceof AbloStaleContextError ) { /* 行が下に移動しました — 再読み取り、再試行 */ }
}
通常の保留作業にはコールバック フォームを使用することをお勧めします。マニュアル範囲のクレームは、
より長い存続期間にわたって利用可能ですが、コールバッククレームがドキュメントのデフォルトです。
請求の全文については、「Coordination」を参照してください/claimState/
キュー/リリースのリファレンス。
React アプリでは、これは同じ ablo.<model> API であり、単に
プロバイダーを呼び出し、 @abloatai/ablo/react からフックで読み取ります。ツリーを一度包みます。
内部のすべてが生きています。
import { AbloProvider , useAblo } から '@abloatai/ablo/react' ;
'./ablo/schema' から { スキーマ } をインポートします。
関数アプリ ( ) {
戻る (
< AbloProvider スキーマ = { スキーマ } >
< レポート ID = "report_stockholm" />
</ アブロプロバイダ >
) ;
}
関数レポート ({ id } : { id : string } ) {
const report = useAblo((ablo) => ablo.weatherReports.retrieve(id));
const ablo = useAblo() ;
if (!レポート) は null を返します。
戻る (
< button onClick = { ( ) => ablo ?。天気予報。更新 ( id , { sta

タス: '準備完了' } ) } >
《報告書》ステータス }
</ボタン>
) ;
}
useAblo(selector) 読み取りは、行が変更されるたびに再レンダリングされます。
チームメイト、またはエージェントがそれを変更しました。書き込みは同じ楽観的なファンアウトです
上記のサーバーの例のようなメソッドです。
<AbloProvider> が接続を所有します。ブラウザーには API キーはありません。それが
ループ全体: useAblo(selector) で読み取り、 ablo.<model> で書き込み、およびすべての
その行の他のクライアント (人間またはエージェント) はそれをリアルタイムで確認します。参照
完全な <AbloProvider> プロップ サーフェス ( userId 、
TeamIds 、 syncGroups 、 fallback 、 bootstrapMode ) およびステータス フック。
Ablo は認証プロバイダーではありません。独自の認証プロバイダー (Clerk、Auth0、NextAuth、
何でも）。 Ablo のジョブは、リクエストを認証した後に開始されます。つまり、リクエストを認証した後に開始されます。
誰が接続しているのか、リアルタイム データの範囲を適切な同期に限定します
グループ (両方とも単位である org:acme または Deck:abc123 のような名前付きチャネル)
ファンアウトとアクセスの単位)。
モデルはプロキシです。ABLO_API_KEY は信頼できるサーバー上に残ります。
サーバーはサインインしているユーザー (組織 / チーム / ユーザー) を独自の認証から解決します。
ブラウザは、すでにスコープが設定された参加者として接続します。ブラウザはキーを保持することはありません。
そしてそれ自体の範囲を広げることはできません。スキーマの ID ロールがその ID をマップします
同期グループ文字列に。
userId / TeamId は認証から取得され、サーバー側で解決されます。
< AbloProvider スキーマ = { スキーマ } userId = { user . id } チーム ID = { ユーザー .チームID } >
< アプリ />
</ アブロプロバイダ >
上記のクイック スタートで組織、チーム、ユーザーがどこから来たのかが明らかでない場合は、
それはあなたのアプリからのものだからです - を見てください
ID と同期グループの全体像: 同期とは何か
group は、スコープの 2 つの部分 (identityRoles + モデルごとの orgScoped /
syncGroupFormat )、および API キーなしで ID が Ablo に到達する方法
ブラウザ。
そこには

独立したマルチプレイヤー モードはありません。人間の UI、サーバーのアクション、およびエージェントの場合
ワーカーは同じスキーマを共有し、 ablo.<model> を通じて書き込みます。
お互いの変更をリアルタイムで — これはデフォルトであり、有効にする機能ではありません。
ablo.<model>.create/update/delete は、確認されたデルタをサブスクライバにファンアウトします。
useAblo(...) は、React クライアントにライブ行を提供し、自動的に最新に保たれます。
ablo.<model>.claim(id) /claimState(id) / queue(id) を使用すると、書き込みが完了する前に人間とエージェントが行 (およびその後ろで待機している行) でアクティブな作業を調整 (および観察) できます。
常に Ablo 経由で書き込みます — SDK モデル メソッドのいずれか
( ablo.<model>.create/update/delete ) または以下の HTTP 書き込みエンドポイント。もしあなたが
代わりに独自のデータベースに直接書き込むと、それらの変更は接続に到達しません
クライアント。
JavaScript を使用していて、型付きモデルまたはリアルタイムが必要な場合は、SDK を使用します。を使用します。
サーバー間の呼び出し元が、ファイルを開かずに書き込む必要がある場合の HTTP エンドポイント
ウェブソケット:
カール https://api.abloatai.com/v1/commits \
-H " 権限: ベアラー sk_test_... " \
-H " Content-Type: application/json " \
-d ' { "操作": [
{ "アクション": "更新"、"モデル": "weatherReports"、"id": "report_stockholm"、"データ": { "ステータス": "準備完了" } }
] } '
{ "オブジェクト" : " commit_receipt " , "stat

[切り捨てられた]

## Original Extract

State control for AI agents. Persist, coordinate, and audit every human and agent write to shared application state. - Abloatai/ablo

GitHub - Abloatai/ablo: State control for AI agents. Persist, coordinate, and audit every human and agent write to shared application state. · GitHub
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
Abloatai
/
ablo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows docs docs examples examples scripts scripts src src .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE NOTICE NOTICE README.md README.md llms-full.txt llms-full.txt llms.txt llms.txt package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json tsconfig.test.json tsconfig.test.json View all files Repository files navigation
Ablo is a typed sync engine for shared app state — the kind that humans,
server code, and AI agents all edit at once.
Reach for it when those edits need to show up everywhere in real time, not
silently overwrite each other, expose who's working on what, and leave a record
of who changed what.
schema -> ablo.<model>.create/retrieve/update/claim(...)
Why Ablo
Real-time by default. Every create / update / delete fans out
confirmed deltas to all subscribers — humans and agents — with no separate
"multiplayer mode" to switch on.
No silent clobbers. Writes are guarded against stale reads, and claim
holds a row across a slow read → LLM → write gap so concurrent edits queue
instead of overwriting.
Built for agents. See who's mid-edit ( claimState / queue ), coordinate a
fair line, and ship an llms.txt so coding agents integrate from the real API.
Typed end to end. Your Zod schema produces typed model proxies
( ablo.<model>.update(...) ), optimistic local reads, and reactive React hooks.
Bring your own auth and database. Ablo scopes realtime data to sync
groups from your existing identity, and can leave your database as the source
of truth via a Data Source.
Built for: collaborative editors, AI agent workflows, internal tools, and any
app where multiple actors mutate shared state and everyone must see it live.
npm install @abloatai/ablo
Keys & runtime. Ablo needs Node 22+ and TypeScript 5+. Grab an sk_test_*
key for a sandbox
( export ABLO_API_KEY=sk_test_... ); keep keys in trusted server runtimes only.
In the browser, <AbloProvider> authenticates with the signed-in user's
session — never the raw key.
Then wire it by hand — the Quick Start below is the shape to
copy. For production (React, an existing backend, Data Source, agents), the
Integration Guide is the deeper map.
Prefer to let an agent wire it? The package ships an llms.txt — a precise
map of the API — so Claude Code or Cursor integrates from the real surface
instead of guessing:
Read node_modules/@abloatai/ablo/llms.txt , then add an Ablo schema, a <AbloProvider> , and my first create / retrieve / update.
import Ablo from '@abloatai/ablo' ;
import { defineSchema , model , z } from '@abloatai/ablo/schema' ;
const schema = defineSchema ( {
weatherReports : model ( {
location : z . string ( ) ,
status : z . enum ( [ 'pending' , 'ready' ] ) ,
forecast : z . string ( ) . optional ( ) ,
} ) ,
} ) ;
const ablo = Ablo ( {
schema ,
apiKey : process . env . ABLO_API_KEY ,
} ) ;
await ablo . ready ( ) ;
const created = await ablo . weatherReports . create ( {
location : 'Stockholm' ,
status : 'pending' ,
} ) ;
const updated = await ablo . weatherReports . update ( created . id , {
status : 'ready' ,
forecast : 'Light rain, 13C' ,
} ) ;
console . log ( { id : updated . id , status : updated . status } ) ;
await ablo . dispose ( ) ;
Expected output:
{ id: '...', status: 'ready' }
Pass schema to get typed models like ablo.weatherReports.update(...) .
retrieve(id) returns one row from the local cache — synchronous, no round-trip.
list(...) filters and sorts what's already synced; it's also synchronous, and
reactive under useAblo / subscribe . load(...) fetches from the server when a
row may not be local yet.
ablo . weatherReports . retrieve ( 'report_stockholm' ) ;
const pending = ablo . weatherReports . list ( {
where : { status : 'pending' } ,
orderBy : { location : 'asc' } ,
limit : 20 ,
} ) ;
const ready = await ablo . weatherReports . load ( {
where : { status : 'ready' } ,
type : 'complete' ,
} ) ;
An array value in where means IN . On load , type: 'complete' waits for
the server; 'unknown' returns what's local now and refreshes in the background.
create / update apply optimistically and resolve to the row. Two options
matter day to day:
await ablo . weatherReports . update ( id , { status : 'ready' } , { wait : 'confirmed' } ) ;
To guard a write against a row that changed under you, pass readAt + onStale
— see Coordinating long agent work .
An agent reads a row, thinks for 30s, writes back — and clobbers whatever changed
meanwhile, or worse, acts on stale state. claim holds the row across that gap:
await ablo . weatherReports . claim ( 'report_stockholm' , async ( report ) => {
const forecast = await weatherAgent . getWeather ( report . location ) ;
await ablo . weatherReports . update ( report . id , { forecast , status : 'ready' } ) ;
} ) ;
If someone else holds the row, claim() waits in a fair queue, then re-reads —
so report is the current row, never a stale snapshot. Reads stay open by
default; only acting on the row serializes. The claim releases when the callback
returns or throws.
See who's mid-edit before you act — decide to wait, or skip:
ablo . weatherReports . claimState ( 'report_stockholm' ) ;
ablo . weatherReports . queue ( 'report_stockholm' ) ;
await ablo . weatherReports . claim ( id , async ( report ) => {
/* do the held work */
} , { wait : false } ) ;
await ablo . weatherReports . claim ( id , async ( report ) => {
/* do the held work */
} , { maxQueueDepth : 2 } ) ;
claimState returns the holder (or null ); queue returns the line waiting
behind it. wait: false skips rather than waiting when the row is held;
maxQueueDepth: 2 bails when two or more are already ahead.
Default reads keep working while a row is claimed. Server reads that need claimed
semantics can opt in with ifClaimed: 'return' | 'wait' | 'fail' .
Even an unclaimed write can't land on stale reasoning — the commit is guarded:
try {
await ablo . weatherReports . update ( id , { status : 'ready' } , { readAt , onStale : 'reject' } ) ;
} catch ( e ) {
if ( e instanceof AbloStaleContextError ) { /* row moved under you — re-read, retry */ }
}
Prefer the callback form for ordinary held work. Manual scoped claims are
available for wider lifetimes, but callback claims are the docs default.
See Coordination for the full claim / claimState /
queue / release reference.
In a React app it's the same ablo.<model> API — just mounted through a
provider and read with hooks, from @abloatai/ablo/react . Wrap your tree once;
everything inside is live.
import { AbloProvider , useAblo } from '@abloatai/ablo/react' ;
import { schema } from './ablo/schema' ;
function App ( ) {
return (
< AbloProvider schema = { schema } >
< Report id = "report_stockholm" />
</ AbloProvider >
) ;
}
function Report ( { id } : { id : string } ) {
const report = useAblo ( ( ablo ) => ablo . weatherReports . retrieve ( id ) ) ;
const ablo = useAblo ( ) ;
if ( ! report ) return null ;
return (
< button onClick = { ( ) => ablo ?. weatherReports . update ( id , { status : 'ready' } ) } >
{ report . status }
</ button >
) ;
}
The useAblo(selector) read re-renders whenever the row changes — whether you,
a teammate, or an agent changed it. The write is the same optimistic, fan-out
method as the server example above.
<AbloProvider> owns the connection — no API key in the browser. That's the
whole loop: read with useAblo(selector) , write with ablo.<model> , and every
other client (human or agent) on that row sees it in real time. See
React for the full <AbloProvider> prop surface ( userId ,
teamIds , syncGroups , fallback , bootstrapMode ) and status hooks.
Ablo is not an auth provider — you keep your own (Clerk, Auth0, NextAuth,
whatever). Ablo's job starts after you've authenticated a request: you tell it
who is connecting, and it scopes their realtime data to the right sync
groups (named channels like org:acme or deck:abc123 that are both the unit
of fan-out and the unit of access).
The model is a proxy: your ABLO_API_KEY stays on your trusted server, your
server resolves the signed-in user (org / team / user) from your own auth, and
the browser connects as an already-scoped participant — it never holds the key
and can't widen its own scope. Your schema's identityRoles map that identity
to sync-group strings.
userId / teamIds come from your auth, resolved server-side:
< AbloProvider schema = { schema } userId = { user . id } teamIds = { user . teamIds } >
< App />
</ AbloProvider >
If it isn't obvious where org / team / user come from in the Quick Start above,
that's because they come from your app — see
Identity & Sync Groups for the full picture: what a sync
group is, the two halves of scoping ( identityRoles + per-model orgScoped /
syncGroupFormat ), and how identity reaches Ablo without an API key in the
browser.
There is no separate multiplayer mode. When human UI, server actions, and agent
workers share the same schema and write through ablo.<model> , they all see
each other's changes in real time — that's the default, not a feature you turn on.
ablo.<model>.create/update/delete fan out confirmed deltas to subscribers.
useAblo(...) gives React clients the live row, kept current automatically.
ablo.<model>.claim(id) / claimState(id) / queue(id) let humans and agents coordinate (and observe) active work on a row — and the line waiting behind it — before a write lands.
Always write through Ablo — either the SDK model methods
( ablo.<model>.create/update/delete ) or the HTTP write endpoint below. If you
write straight to your own database instead, those changes won't reach connected
clients.
Use the SDK when you are in JavaScript and want typed models or realtime. Use the
HTTP endpoint when a server-to-server caller needs to write without opening a
WebSocket:
curl https://api.abloatai.com/v1/commits \
-H " Authorization: Bearer sk_test_... " \
-H " Content-Type: application/json " \
-d ' { "operations": [
{ "action": "update", "model": "weatherReports", "id": "report_stockholm", "data": { "status": "ready" } }
] } '
{ "object" : " commit_receipt " , "stat

[truncated]

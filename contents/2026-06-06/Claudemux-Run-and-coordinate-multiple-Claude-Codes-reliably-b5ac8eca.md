---
source: "https://github.com/wastedcode/claudemux"
hn_url: "https://news.ycombinator.com/item?id=48426188"
title: "Claudemux – Run and coordinate multiple Claude Codes reliably"
article_title: "GitHub - wastedcode/claudemux: Drive multiple Claude Code sessions on a box from Node/CLI · GitHub"
author: "zeppelin_7"
captured_at: "2026-06-06T18:33:17Z"
capture_tool: "hn-digest"
hn_id: 48426188
score: 2
comments: 1
posted_at: "2026-06-06T15:54:18Z"
tags:
  - hacker-news
  - translated
---

# Claudemux – Run and coordinate multiple Claude Codes reliably

- HN: [48426188](https://news.ycombinator.com/item?id=48426188)
- Source: [github.com](https://github.com/wastedcode/claudemux)
- Score: 2
- Comments: 1
- Posted: 2026-06-06T15:54:18Z

## Translation

タイトル: Claudemux – 複数のクロード コードを確実に実行および調整する
記事のタイトル: GitHub - wastedcode/claudemux: ノード/CLI · GitHub からボックス上で複数のクロード コード セッションを駆動する
説明: ノード/CLI からボックス上で複数のクロード コード セッションを駆動する - wastedcode/claudemux

記事本文:
GitHub - wastedcode/claudemux: ノード/CLI からボックス上で複数のクロード コード セッションを駆動する · GitHub
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
無駄なコード
/
クロードムクス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
113 コミット 113 コミット .github .github bin bin docs/決定 docs/決定 例 例 スクリプト スクリプト src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Node からボックス上で複数のリアルログイン Claude Code セッションを実行し、調整します。
await session.wait() は、エージェントが完了すると実際に戻ります。
クロードがマシンにログインしており、それをコードから操作したいと考えています。つまり、セッション (またはフリート全体) を生成し、タスクを送信し、タスクが実際にいつ完了したかを知り、結果を読み取り、それらを調整します。今日では、 child_process.spawn('claude', …) + アドホック ANSI 正規表現 + sleep(5) 、 N 回のセッション、およびそれらの衝突を防ぐ接着剤です。初回実行の信頼ダイアログでハングし、プロンプトで静かに停止し、クロードの更新ごとに腐ります。 claudemux はそのレイヤーを一度引退します。
import { create } from "@wastedcode/claudemux" ;
const session = await create ( { name : "ジョブ" , cwd : process . cwd() } ) ;
セッションを待ちます。 send ( "次のリリース用に CHANGELOG エントリを追加" ) ;
セッションを待ちます。待って （ ） ; // ターン終了までブロックします。 { maxMs } / { idleMs } を渡してバインドします
const text = セッションを待ちます。捕獲 （ ） ;
create() はエー​​ジェントを起動し、初回実行ダイアログを閉じ、スリープ後ではなく、REPL が本当に準備ができたときに戻ります。 wait() はターンがターミナル TurnOutcome に到達するまでブロックします — 完了 (応答が読み取れる)、待機中

決定、中止、または忍耐力の限界を超えた決定 — 画面スクレイピングではなく、エージェントのフックとトランスクリプトから融合されています。 1 回の呼び出しでの往復全体に対して ask() (送信 → 待機 → 読み取り) が行われます。クラッシュ後に会話を続けるには、resume() があります。
それが1つのセッションです。名前は残りの半分です。フリートを駆動し、各セッションは name でアドレス指定され、1 つのプロセスから実行されます。
import { create , ask } from "@wastedcode/claudemux" ;
// 一度に 3 つをブートします。それぞれが独自の実際のクロード セッションです。
const フリート = Promise を待ちます。全部（
[ "api" 、 "ui" 、 "docs" ] 。 map ( ( name ) => create ( { name , cwd : `./services/ ${ name } ` } ) ) 、
) ;
// タスクをファンに広めます。実際に終了したらそれぞれを収集します。
const summary = await Promise 。全部（
艦隊。 map ( ( s ) => ask ( s , "このサービスで今週変更されたことを要約してください" ) ) ,
) ;
list() はフリートを列挙します。別のプロセスは、名前によって任意のセッションを採用できます (デーモンとワーカーのドライブの分割)。 1 つの信頼できるセッションは、まさに最小のフリートです。
そして、これはヘッドレス パイプではなく、実際のセッションです。それぞれは tmux セッションで実行され、エージェントの動作を監視するために tmux に接続できます。または、自分でキーボードを使用することもできます (一度に 1 人のライター、§4)。プログラムによるハンドルと、座って参加できるライブ セッションは同じセッションです。
これは、制御するボックス上で Consumer-login claude CLI ( claude login で設定したもの) を駆動するためです。オーケストレーターが相互に通信する複数の claude セッションを実行するのと同じように、1 つのセッションまたは複数のセッションを調整します。これはボックスのクロード設定 (認証、許可モード、モデル、MCP) を継承し、クロード自身のフラグを渡します。独自の構成は所有しません (1 つの例外: ワークスペースの信頼、§4)。
これは対象外です: 注入された認証情報を介してクロードを駆動するデプロイされた自動化または匿名の自動化

ls または API キー — CI フリート、ホスト型サービス。そこでは Consumer-login claude を実行できません (エフェメラル ボックスはインタラクティブにログインできず、Anthropic の規約に違反します)。それが Claude Agent SDK + API の目的です。 claudemux は、オン・ア・ボックスのリアル・ログイン・ケースを信頼性の高いものにします。
npm install @wastedcode/claudemux
ノード ≥20 と、PATH 上で動作する claude CLI が必要です (認証されるために少なくとも 1 回対話的に claude を実行している必要があります)。 MITライセンス取得済み。
CLI とライブラリは 1:1 にマップされます。claudemux の送信名 "..." は、ライブラリ側では send(name, "...") です。語彙がひとつ。
$ npm i @wastedcode/claudemux
$ claudemux spawn my-job --cwd ./fresh-repo --trust-workspace
{ " AgentSessionId " : " f47ac10b-58cc-4372-a567-0e02b2c3d479 " } # 再開のためにこれを保持します
$ claudemux ask my-job " 次のリリースの CHANGELOG エントリを追加する "
{ " 結果 " :{ " 種類 " : " 完了 " }、 " メッセージ " :[…]、 " カーソル " : " … " }
$ claudemux 私の仕事を殺します
ask はワンショットの往復です。プリミティブ (送信 → 待機 → メッセージ) は、自分でターンを操作したい場合に使用します。
決して信頼されていないフォルダーでの最初の生成には --trust-workspace (上記) が必要です。それ以外の場合は失敗して閉じられ、フラグによって永続的なフォルダーごとの権限付与が書き込まれます。制御していないコードを指定する前に、「ワークスペースの信頼 (フェールクローズ)」を参照してください。
生成 / 再開 / 送信 / 問い合わせ / 待機 / 状態 / キャプチャ /… take --agent ;レジストリ動詞 ( kill / list / doesn't ) はそうではありません。
すべてのコマンドは --namespace <name> (デフォルトは claudemux ) を受け入れるため、1 台のマシン上の 2 つのコンシューマーが衝突しません。
同じユーザーからのすべての claudemux 呼び出しは、1 つのランデブー ソケット (OS によって UID ごとに所有されるデフォルトの claudemux ソケット ファイル) を共有します。これにより、あるプロセスでのスポーンが後続のプロセスでの送信/待機/キャプチャに表示されるようになります。隔離されたソケットにオプトインするには (1 秒)

同じボックス上の独立したオーケストレーター、またはデバッグなど）、環境内で --socket <name> を渡すか、CLAUDEMUX_SOCKET=<name> を設定します。
ライブラリは CLI をミラーリングします。正規の 30 秒サンプルは、examples/spawn-send-wait-capture.ts にあり、唯一の正規サンプルです。README スニペットは、それを複製するのではなく参照します。
import { ask 、 create 、 type SessionHandle } from "@wastedcode/claudemux" ;
const session : SessionHandle = await create ( { name : "ジョブ" , cwd : process . cwd() } ) ;
// 1 往復 — 90% のパス:
const { 結果 , メッセージ } = await ask ( session , "CHANGELOG エントリの追加" ) ;
if ( 結果 . 種類 === "完了" ) コンソール。 log (メッセージ .at (-1) ) ;
else handleAbnormal (結果) ; // 待っています |中止されました |予算超過
セッションを待ちます。殺す （ ） ;
wait() は TurnOutcome を返します。これは分岐する識別共用体であり、タイムアウトがスローされることはありません。
const カーソル = セッションを待ちます。送信 （ "…" ） ;
const 結果 = セッションを待ちます。 wait ({ maxMs : 60_000 } ) ;
スイッチ (結果の種類) {
ケース "完了" : ブレーク ; // 返信は読み取り可能です (フラッシュスキューが閉じられています)
case "待機中" : 結果。の上 ; /* "許可プロンプト" | "ダイアログ" */ ブレーク ;
ケース "中止" : ブレーク ; // 割り込み() によって停止されました
ケース「予算超過」: 結果。理由 ; /* "アイドル" (スタック) | "max" (壁時計) */ ブレーク ;
}
completed は、応答が読み取り可能であることを保証します。次のmessages Because(cursor) には競合がありません。予算超過は失敗を意味するわけではありません。忍耐力が尽きましたが、ターンはまだ実行中である可能性があるため、やみくもに再送信しないでください (ライブ ターン キューに再送信すると、副作用が発生したり重複したりするため、最悪の失敗になります)。代わりに、progress() をポーリングします。toolInFlight === true または新たに進むtransscriptCount は、遅いですが生きている (待ち続ける) ことを意味します。状態が異なる長いフラットなtransscriptCount

動作しているということは、おそらくウェッジされたことを意味します（その後、中断（）、再送信しません）。着地していないことを確認したターンのみを再送信します —turnComplete(cursor) === false 。
名前の理由。セッションは独立しており、 name でアドレス指定されるため、多くのセッションを起動し、それぞれを名前で駆動し、 list() でそれらを列挙します。ハンドル create() のハンドバックは維持されます。
import { create , list , kill } from "@wastedcode/claudemux" ;
const names = [ "api" , "ui" , "docs" ] ;
const フリート = オブジェクト 。 fromエントリ (
約束を待ちます。全部（
名前 。 map ( async ( name ) => [ name , await create ( { name , cwd : `./services/ ${ name } ` } ) ] ) 、
）、
) ;
艦隊を待ちます。 API 。 send ( "テストを実行し、失敗を要約します" ) ;
艦隊を待ちます。うい。 send ( "デザイントークンのバージョンをバンプする" ) ;
待機リスト ( ) ; // ['api', 'ui', 'docs'] — フリート ビュー
for (await list() の const 名) await kill({name}) ;
名前はプロセスより長く存続します。デーモンはフリートを作成でき、別個のワーカーは名前によって任意のセッションを採用してそれを駆動できます (「 Adopt 」を参照)。お互いの list() が見えないフリートが必要ですか?それらに異なる名前空間を与えます。一度に多数をブートするのは安全です (誤った準備やクロストークはありません)。ただし、ビジー ボックス上の大規模なフリートを抑制するのは、基板の責任ではなく、ユーザーの責任です (ブートの同時実行性はユーザーの責任です)。
ターンの出力の読み取り ( send →messagesSince / progress )
send() は、そのターンに固定された Cursor を返します。メッセージを読み返してください
構造化されたバックエンド中立のメッセージとして生成されます。ペインスクレイピングはありません。
const カーソル = セッションを待ちます。 send ( "README を 1 行に要約します" ) ;
セッションを待ちます。待って （ ） ; // ターンが確定する
const msgs = セッションを待ちます。メッセージ以来 (カーソル) ;
// → [{ 役割: "ユーザー", パーツ: [{ 種類: "テキスト", テキスト }] },
// { 役割: "アシスタント", パーツ: [{ 種類: "テキスト", テキスト } , { 種類: "ツール", ツール, 概要 },

…] }]
信頼性の高い「動作していますか? / 完了しましたか?」を確認するには、エージェントの機能と融合した progress() を使用してください。
TUI ではなく、フック + トランスクリプト (決定論的):
const p = セッションを待ちます。進捗 （ ） ;
// { フェーズ: "プロンプト"|"ツール"|"作曲中"|"完了"|"不明",
// toolInFlight: boolean, // ツールは正当に実行されています (ハングしていません)
// トランスクリプトカウント: 数値,
//hookChannelHealthy: boolean, // false → ベストエフォートペインフォールバックに劣化
// AgentChannelHealthy: boolean, // false → すべてのチャネルがブラインドと空ではないペイン (おそらくクロード バージョンのドリフト)
// 状態 }
忍耐力はあなたのものです: フェーズ === "完了" (またはあなた自身の) まで、poll progress()
予算が経過した場合) — claudemux はシグナルを報告しますが、アイドル タイムアウトは発生しません。フックは
デフォルトではスポーン時に注入されます。 create({hooks:false}) でオプトアウトします (観察してください)
その後、ペインのフォールバックに低下し、hookChannelHealthy: false を介してその旨を通知します)。
ベアネーム操作 (ハンドルは必要ありません):
import { 存在 , kill , list } from "@wastedcode/claudemux" ;
await が存在します ( { name : "ジョブ" } ) ; // ブール値
待機リスト ( ) ; // デフォルトの名前空間内の名前の string[]
await kill ({ name : "ジョブ" } ) ; // 冪等
クラッシュ後の会話の再開 (resume() +turnComplete())
resume() は、create() (新規に開始) および Adopt() のファーストクラスのライフサイクル ピアです。
(実行中のペインに再接続します)。

[切り捨てられた]

## Original Extract

Drive multiple Claude Code sessions on a box from Node/CLI - wastedcode/claudemux

GitHub - wastedcode/claudemux: Drive multiple Claude Code sessions on a box from Node/CLI · GitHub
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
wastedcode
/
claudemux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
113 Commits 113 Commits .github .github bin bin docs/ decisions docs/ decisions examples examples scripts scripts src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json package-lock.json package-lock.json package.json package.json tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Run and coordinate multiple real-login Claude Code sessions on your box, from Node.
await session.wait() actually returns when the agent is done.
You have claude logged in on your machine and you want to drive it from code — spawn a session (or a whole fleet), send a task, know when it's actually done , read the result, coordinate them. Today that's child_process.spawn('claude', …) + ad-hoc ANSI regex + sleep(5) , times N sessions, plus glue to keep them from colliding: it hangs on the first-run trust dialog, silently stalls on prompts, and rots on every claude update. claudemux retires that layer once.
import { create } from "@wastedcode/claudemux" ;
const session = await create ( { name : "job" , cwd : process . cwd ( ) } ) ;
await session . send ( "Add a CHANGELOG entry for the next release" ) ;
await session . wait ( ) ; // blocks until the turn ends; pass { maxMs } / { idleMs } to bound it
const text = await session . capture ( ) ;
create() boots the agent, dismisses the first-run dialogs, and returns when the REPL is genuinely ready — not after a sleep . wait() blocks until the turn reaches a terminal TurnOutcome — completed (and the reply is readable), awaiting a decision, aborted , or out of your patience budget — fused from the agent's hooks + transcript, not screen-scraping. For the whole round-trip in one call there's ask() (send → wait → read); to continue a conversation after a crash there's resume() .
That's one session. The name is the other half — drive a fleet , each session addressed by name , from one process:
import { create , ask } from "@wastedcode/claudemux" ;
// Boot three at once — each is its own real claude session.
const fleet = await Promise . all (
[ "api" , "ui" , "docs" ] . map ( ( name ) => create ( { name , cwd : `./services/ ${ name } ` } ) ) ,
) ;
// Fan a task across them; collect each as it actually finishes.
const summaries = await Promise . all (
fleet . map ( ( s ) => ask ( s , "Summarize what changed in this service this week" ) ) ,
) ;
list() enumerates the fleet; another process can adopt any session by name (the daemon-spawns / workers-drive split). One reliable session is just the smallest fleet.
And it's a real session, not a headless pipe: each runs in a tmux session you can tmux attach into to watch the agent work — or take the keyboard yourself (one writer at a time, §4). The programmatic handle and the live session you can sit down at are the same session.
What this is for: driving the consumer-login claude CLI (the one you set up with claude login ) on a box you control — one session, or many coordinating, the way an orchestrator might run several claude sessions that talk to each other. It inherits your box's claude config (auth, permission mode, model, MCP) and passes claude's own flags through; it owns no configuration of its own (one exception: workspace trust, §4).
What this is not for: deployed or anonymous automation that drives claude via injected credentials or API keys — CI fleets, hosted services. Consumer-login claude can't run there (ephemeral boxes can't interactively log in, and it's against Anthropic's terms); that's what the Claude Agent SDK + API are for. claudemux makes the on-a-box, real-login case reliable.
npm install @wastedcode/claudemux
Requires Node ≥20 and a working claude CLI on PATH (you've run claude interactively at least once so it's authenticated). MIT-licensed.
The CLI and library map 1:1 — claudemux send name "..." is send(name, "...") on the library side. One vocabulary.
$ npm i @wastedcode/claudemux
$ claudemux spawn my-job --cwd ./fresh-repo --trust-workspace
{ " agentSessionId " : " f47ac10b-58cc-4372-a567-0e02b2c3d479 " } # persist this for resume
$ claudemux ask my-job " Add a CHANGELOG entry for the next release "
{ " outcome " :{ " kind " : " completed " }, " messages " :[…], " cursor " : " … " }
$ claudemux kill my-job
ask is the one-shot round-trip; the primitives ( send → wait → messages ) are there when you want to drive the turn yourself.
The first spawn in a never-trusted folder needs --trust-workspace (above) — it fails closed otherwise, and the flag writes a persistent per-folder authority grant; see Workspace trust (fail-closed) before pointing it at code you don't control.
spawn / resume / send / ask / wait / state / capture /… take --agent ; the registry verbs ( kill / list / exists ) don't.
Every command accepts --namespace <name> (default claudemux ) so two consumers on one machine don't collide.
All claudemux invocations from the same user share one rendezvous socket (the default claudemux socket file, owned per-UID by the OS). That's how spawn in one process is visible to send / wait / capture in subsequent processes. To opt into an isolated socket (a second independent orchestrator on the same box, or debugging), pass --socket <name> or set CLAUDEMUX_SOCKET=<name> in the environment.
The library mirrors the CLI. The canonical 30-second example lives in examples/spawn-send-wait-capture.ts and is the only canonical sample — README snippets reference it rather than duplicate it.
import { ask , create , type SessionHandle } from "@wastedcode/claudemux" ;
const session : SessionHandle = await create ( { name : "job" , cwd : process . cwd ( ) } ) ;
// One round-trip — the 90% path:
const { outcome , messages } = await ask ( session , "Add a CHANGELOG entry" ) ;
if ( outcome . kind === "completed" ) console . log ( messages . at ( - 1 ) ) ;
else handleAbnormal ( outcome ) ; // awaiting | aborted | budget-exceeded
await session . kill ( ) ;
wait() returns a TurnOutcome — a discriminated union you branch on, never a thrown timeout:
const cursor = await session . send ( "…" ) ;
const outcome = await session . wait ( { maxMs : 60_000 } ) ;
switch ( outcome . kind ) {
case "completed" : break ; // reply is readable (flush-skew closed)
case "awaiting" : outcome . on ; /* "permission-prompt" | "dialog" */ break ;
case "aborted" : break ; // an interrupt() stopped it
case "budget-exceeded" : outcome . reason ; /* "idle" (stuck) | "max" (wall-clock) */ break ;
}
completed guarantees the reply is readable — a following messagesSince(cursor) is race-free. budget-exceeded does NOT mean failed — your patience ran out, but the turn may still be running , so do not blindly re-send (a re-send into a live turn queues or duplicates side effects — the worst failure). Instead poll progress() : toolInFlight === true or a freshly-advancing transcriptCount means slow-but-alive (keep waiting); a long flat transcriptCount with state not working means likely wedged (then interrupt() , don't re-send). Re-send only a turn you've confirmed never landed — turnComplete(cursor) === false .
The reason for the name. Sessions are independent and addressed by name , so you boot many, drive each by name, and enumerate them with list() — keep the handles create() hands back:
import { create , list , kill } from "@wastedcode/claudemux" ;
const names = [ "api" , "ui" , "docs" ] ;
const fleet = Object . fromEntries (
await Promise . all (
names . map ( async ( name ) => [ name , await create ( { name , cwd : `./services/ ${ name } ` } ) ] ) ,
) ,
) ;
await fleet . api . send ( "Run the tests and summarize failures" ) ;
await fleet . ui . send ( "Bump the design-token version" ) ;
await list ( ) ; // ['api', 'ui', 'docs'] — your fleet view
for ( const name of await list ( ) ) await kill ( { name } ) ;
Names outlive your process: a daemon can create the fleet and separate workers can adopt any session by name and drive it (see adopt ). Want fleets that can't see each other's list() ? Give them different namespace s. Booting many at once is safe — no false-ready, no crosstalk — but throttling a big fleet on a busy box is your call, not the substrate's ( Boot concurrency is yours ).
Reading a turn's output ( send → messagesSince / progress )
send() returns a Cursor anchored at that turn. Read back the messages the
turn produced as structured, backend-neutral Message s — no pane-scraping:
const cursor = await session . send ( "Summarize the README in one line" ) ;
await session . wait ( ) ; // turn settles
const msgs = await session . messagesSince ( cursor ) ;
// → [{ role: "user", parts: [{ kind: "text", text }] },
// { role: "assistant", parts: [{ kind: "text", text } , { kind: "tool", tool, summary }, …] }]
For reliable "is it working / done?", use progress() — fused from the agent's
hooks + transcript (deterministic), not the TUI:
const p = await session . progress ( ) ;
// { phase: "prompt"|"tool"|"composing"|"done"|"unknown",
// toolInFlight: boolean, // a tool is legitimately running (not hung)
// transcriptCount: number,
// hookChannelHealthy: boolean, // false → degraded to best-effort pane fallback
// agentChannelHealthy: boolean, // false → ALL channels blind vs a non-empty pane (likely a claude-version drift)
// state }
Patience is yours : poll progress() until phase === "done" (or your own
budget elapses) — claudemux reports the signal, never an idle timeout. Hooks are
injected on spawn by default; opt out with create({ hooks: false }) (observe
then degrades to the pane fallback and says so via hookChannelHealthy: false ).
Bare-name operations (no handle needed):
import { exists , kill , list } from "@wastedcode/claudemux" ;
await exists ( { name : "job" } ) ; // boolean
await list ( ) ; // string[] of names in the default namespace
await kill ( { name : "job" } ) ; // idempotent
Resuming a conversation after a crash ( resume() + turnComplete() )
resume() is a first-class lifecycle peer of create() (start fresh) and adopt()
(re-attach to a running pane).

[truncated]

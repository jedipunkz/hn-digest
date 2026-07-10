---
source: "https://github.com/Botfather90/digiton-agent-fleet"
hn_url: "https://news.ycombinator.com/item?id=48866405"
title: "I run my agency's 10 AI agents from launchd, no n8n, 0 dependencies"
article_title: "GitHub - Botfather90/digiton-agent-fleet: Ten scheduled AI agents run my agency from launchd with zero dependencies. One Node file per agent; every send passes unit-tested rails (kill switch, allowlist, DNC, dedupe, daily cap). No n8n. · GitHub"
author: "Brandon99pt"
captured_at: "2026-07-10T23:49:06Z"
capture_tool: "hn-digest"
hn_id: 48866405
score: 2
comments: 2
posted_at: "2026-07-10T22:59:18Z"
tags:
  - hacker-news
  - translated
---

# I run my agency's 10 AI agents from launchd, no n8n, 0 dependencies

- HN: [48866405](https://news.ycombinator.com/item?id=48866405)
- Source: [github.com](https://github.com/Botfather90/digiton-agent-fleet)
- Score: 2
- Comments: 2
- Posted: 2026-07-10T22:59:18Z

## Translation

タイトル: 代理店の 10 個の AI エージェントを launchd から実行しています。n8n なし、依存関係はありません
記事のタイトル: GitHub - Botfather90/digitalon-agent-fleet: 10 人のスケジュールされた AI エージェントが、依存関係なしで launchd から私の代理店を実行しています。エージェントごとに 1 つのノード ファイル。すべての送信は単体テスト済みのレール (キルスイッチ、ホワイトリスト、DNC、重複排除、日次上限) を通過します。いいえ、n8n。 · GitHub
説明: スケジュールされた 10 人の AI エージェントが、依存関係なしで launchd から私の代理店を実行します。エージェントごとに 1 つのノード ファイル。すべての送信は単体テスト済みのレール (キルスイッチ、ホワイトリスト、DNC、重複排除、日次上限) を通過します。いいえ、n8n。 - Botfather90/digitalon-agent-fleet

記事本文:
GitHub - Botfather90/digitalon-agent-fleet: 10 人のスケジュールされた AI エージェントが、依存関係なしで launchd から私の代理店を実行しています。エージェントごとに 1 つのノード ファイル。すべての送信は単体テスト済みのレール (キルスイッチ、ホワイトリスト、DNC、重複排除、日次上限) を通過します。いいえ、n8n。 · GitHub
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
{{ メッセージ

e }}
ボットファーザー90
/
デジトンエージェントフリート
パブリックテンプレート
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Botfather90/digitalon-agent-fleet
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github/ workflows .github/ workflows bin bin data data jobs jobs lib lib スケジュール スケジュール テスト test .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md フリート.config.example.json フリート.config.example.json llms.txt llms.txt package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
予定されている 10 人の AI エージェントが、リードリサーチ、PR アウトリーチ、返信のトリアージ、パイプライン同期、受信箱の監視、毎日のブリーフィングなど、私の代理店の業務を実行しています。 n8n、フレームワーク、キュー、Docker はありません。各エージェントは launchd タイマー上の 1 つのノード ファイルであり、すべての不可逆アクションは、それが発生する前に単体テストされたガードレールを通過します。このリポジトリは、テンプレートとしてクリーンに抽出されたハーネスです。
Digiton Dynamics では 10 のエージェントが運用中
ランタイム依存関係なし、ビルドステップなし、プレーンノード
レール上の 27 個の単体テスト、ノード 18/20/22 の CI
すべてのアクションに 5 つのレール: キルスイッチ、接触禁止、ホワイトリスト、重複排除、1 日あたりの上限
$ npm デモを実行する
[情報] ドラフト ana@oceanview-realty.pt :: 許可、送信されません (mode=draft)
[情報] carlos@montebianco.com をブロック:: 受信者が許可リストに載っていません
[情報] DNC で info@donotcontact.example :: 受信者をブロックします
[情報] 完了 提案=3 許可=1 ブロック=2 実行=0 殺害=false
レールが存在する理由: 私たちはかつて、本番環境で無防備な電子メール ボットをタイマーで実行しました。ホワイトリスト、1 日の上限、オフ スイッチはありません。1 つの間違ったプロンプトまたは 1 つの間違ったデータ行があれば、午前 3 時に間違ったリストをメールすることができます。それをシャットダウンし、 lib/rails.js の単一の純粋な関数 Evaluate() の背後でスケジュールされたすべてのエージェントを再構築しました。

それをクリアしなければ取り返しのつかないことは起こりません。
新しいクローンではデフォルトで安全です: オプトインしない限り、空のホワイトリスト、スタブ送信者、ドライラン。
提案されたアクション -> [ Rails ] -> 送信またはブロック (理由あり)
クイックスタート
git clone https://github.com/Botfather90/digitalon-agent-fleet
cd デジトン エージェント フリート
npm test # Rails単体テストを実行します（ゼロセットアップ）
npm run Demon # レールがサンプルリードを決定するのを観察します (認証情報なし)
独自のコピーを使用する: GitHub ページの上部にある [このテンプレートを使用する] をクリックします。
このリポジトリから新しいリポジトリを開始します。
デモでは、サンプル フォローアップ エージェントを 3 つのサンプル リードに対して予行演習で実行し、動作中のレールを示します。1 つは許可され、1 つはホワイトリストにないためにブロックされ、1 つは連絡禁止リストによってブロックされています。 2 番目の実際の例jobs/example-digest.js は、毎日のダイジェストを Slack スタイルの Webhook チャネルに投稿し、電子メール以外のアクションは同じ Evaluate() を介してルーティングされるため、受信者、DNC、許可リスト、重複排除、および日次の上限がすべて適用されます。 npm run demo:digest で実行します。
エージェントが提案するすべてのアクションは、 lib/rails.js の Evaluate() によってチェックされます。この順序で、最も困難な停止が最初に行われます。
Evaluate() は I/O を持たない純粋な関数であるため、推論が速く、完全に単体テストされています。
テンプレートでは、新しいクローンを他のユーザーに電子メールで送信することはできません。サンプルの送信者は意図的なスタブであり、デフォルトの許可リストは空であり、 --send を渡して "mode": "send" を設定しない限り、ランナーは予行演習されます。単一のメッセージを送信する前に、それらすべてにオプトインする必要があります。
ジョブは、jobs/ 内の 1 つのファイルであり、次の 2 つの関数をエクスポートします。
// ジョブ/my-agent.js
モジュール。エクスポート = {
// 実行したいアクションを返します。ランナーはこれらを信用しません。
// 何かが起こる前に、レールを通してそれぞれを実行します。
非同期提案 ({ config , log } ) {
戻る [
{ 宛先: 'lead@example.com' 、件名: 'こんにちは

' 、本文 : '...' 、重複排除キー : 'lead-2026-06' } 、
] ;
} 、
// 許可されたアクションを 1 つ実行します。実際に発信しているのはここだけです。
非同期送信 (アクション , {構成 , ログ }) {
// 本当の送信者はここに来ます: Gmail、再送信、SES、API 呼び出しなど何でも
} 、
} ;
実行します:node bin/run-job.js my-agent (ドライラン) または --send を追加して、許可されたアクションをディスパッチします。
フリート.config.json でポリシーを設定します ( フリート.config.example.json をコピーします)。
{
"モード" : "ドラフト" ,
「ポリシー」: {
"デイリーキャップ" : 30 、
"allowEmails" : [ "vip@client.com " ],
"allowDomains" : [ " client.com " ],
"dnc" : [ "unsubscribed@x.com" , "competitor.com" ]
}
}
実際の フリート.config.json は gitignored されるため、ホワイトリストがコミットされることはありません。
プラットフォームを選択してください。デフォルトでは、3 つすべてが毎日 09:00 にエージェントを実行します。
macOS (launchd): bash schedule/install.sh は、launchd ジョブをインストールし、それを実行、検査、削除する方法を出力します。
Linux (systemd): ユニットとタイマーをSchedule/systemd/からコピーし、 systemctl enable --now Agent-fleet.timer を実行します。
cron を使用するもの:schedule/crontab.example の行を crontab -e に貼り付けます。
lib/rails.js ガードレール (純粋な Evaluate() + レジャー永続性)
lib/log.js 日付付きファイル + stdout ロガー
bin/run-job.js 管理されたランナー
jobs/example-followup.js 動作するサンプルエージェント (スタブ送信者、安全)
jobs/example-digest.js 同じレールを介した、Slack スタイルの Webhook チャネル (電子メール以外のアクション) への毎日のダイジェスト
data/leads.sample.json サンプル データなので、セットアップなしでデモが実行されます
data/digest.sample.json ダイジェスト例のサンプルデータ
test/rails.test.js レールとケイデンスの単体テスト
test/example-digest.test.js 同じレール上のダイジェスト ジョブの単体テスト
スケジュール/launchd、cron、systemd テンプレート + インストーラー
.github/workflows/test.yml ノード:プッシュ/PR でのテスト (ノード 18/20/22)
ライセンス
マサチューセッツ工科大学。建てられた

Digiton Dynamics によって保守されています。
スケジュールされた 10 人の AI エージェントが、依存関係なしで launchd から私の代理店を実行します。エージェントごとに 1 つのノード ファイル。すべての送信は単体テスト済みのレール (キルスイッチ、ホワイトリスト、DNC、重複排除、日次上限) を通過します。いいえ、n8n。
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

Ten scheduled AI agents run my agency from launchd with zero dependencies. One Node file per agent; every send passes unit-tested rails (kill switch, allowlist, DNC, dedupe, daily cap). No n8n. - Botfather90/digiton-agent-fleet

GitHub - Botfather90/digiton-agent-fleet: Ten scheduled AI agents run my agency from launchd with zero dependencies. One Node file per agent; every send passes unit-tested rails (kill switch, allowlist, DNC, dedupe, daily cap). No n8n. · GitHub
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
Botfather90
/
digiton-agent-fleet
Public template
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Botfather90/digiton-agent-fleet
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows bin bin data data jobs jobs lib lib schedule schedule test test .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md fleet.config.example.json fleet.config.example.json llms.txt llms.txt package.json package.json View all files Repository files navigation
Ten scheduled AI agents run my agency's operations: lead research, PR outreach, reply triage, pipeline sync, inbox monitoring, daily briefings. No n8n, no framework, no queue, no Docker. Each agent is one Node file on a launchd timer, and every irreversible action passes a unit-tested guardrail before it happens. This repo is that harness, extracted clean as a template.
10 agents in production at Digiton Dynamics
0 runtime dependencies, no build step, plain Node
27 unit tests on the rails, CI on Node 18/20/22
5 rails on every action: kill switch, do-not-contact, allowlist, dedupe, daily cap
$ npm run demo
[INFO] DRAFT ana@oceanview-realty.pt :: allowed, not sent (mode=draft)
[INFO] BLOCK carlos@montebianco.com :: recipient not on allowlist
[INFO] BLOCK info@donotcontact.example :: recipient on DNC
[INFO] done proposed=3 allowed=1 blocked=2 performed=0 killed=false
Why the rails exist: we once ran an unguarded email bot on a timer in production. No allowlist, no daily cap, no off switch - one bad prompt or one bad data row away from mailing the wrong list at 3am. We shut it down and rebuilt every scheduled agent behind a single pure function, evaluate() in lib/rails.js . Nothing irreversible happens without clearing it.
Safe by default on a fresh clone: empty allowlist, stub sender, dry-run unless you opt in.
proposed action -> [ rails ] -> sent or blocked (with a reason)
Quickstart
git clone https://github.com/Botfather90/digiton-agent-fleet
cd digiton-agent-fleet
npm test # run the rails unit tests (zero setup)
npm run demo # watch the rails decide on sample leads (no credentials)
Prefer your own copy: click Use this template at the top of the GitHub page
to start a fresh repo from this one.
The demo runs the example follow-up agent in dry-run against three sample leads and shows the rails at work: one allowed, one blocked for not being on the allowlist, one blocked by the do-not-contact list. A second worked example, jobs/example-digest.js , posts a daily digest to a Slack-style webhook channel, a non-email action routed through the same evaluate() so recipient, DNC, allowlist, dedupe, and daily cap all still apply. Run it with npm run demo:digest .
Every action the agent proposes is checked by evaluate() in lib/rails.js , in this order, hardest stop first:
evaluate() is a pure function with no I/O, so it is fast to reason about and fully unit-tested.
The template cannot email anyone on a fresh clone. The example sender is a deliberate stub, the default allowlist is empty, and the runner is dry-run unless you pass --send and set "mode": "send" . You have to opt in to every one of those before a single message leaves.
A job is one file in jobs/ that exports two functions:
// jobs/my-agent.js
module . exports = {
// Return the actions you want to take. The runner does NOT trust these,
// it runs each one through the rails before anything happens.
async propose ( { config , log } ) {
return [
{ to : 'lead@example.com' , subject : 'Hello' , body : '...' , dedupeKey : 'lead-2026-06' } ,
] ;
} ,
// Perform ONE allowed action. This is the only place that actually sends.
async send ( action , { config , log } ) {
// your real sender goes here: Gmail, Resend, SES, an API call, whatever
} ,
} ;
Run it: node bin/run-job.js my-agent (dry-run) or add --send to dispatch allowed actions.
Set your policy in fleet.config.json (copy fleet.config.example.json ):
{
"mode" : " draft " ,
"policy" : {
"dailyCap" : 30 ,
"allowEmails" : [ " vip@client.com " ],
"allowDomains" : [ " client.com " ],
"dnc" : [ " unsubscribed@x.com " , " competitor.com " ]
}
}
Your real fleet.config.json is gitignored, so your allowlist never ends up in a commit.
Pick your platform. All three run the agent daily at 09:00 by default.
macOS (launchd): bash schedule/install.sh installs a launchd job and prints how to run, inspect, and remove it.
Linux (systemd): copy the unit and timer from schedule/systemd/ and systemctl enable --now agent-fleet.timer .
Anything with cron: paste the line from schedule/crontab.example into crontab -e .
lib/rails.js the guardrail (pure evaluate() + ledger persistence)
lib/log.js dated file + stdout logger
bin/run-job.js the governed runner
jobs/example-followup.js a worked example agent (stub sender, safe)
jobs/example-digest.js a daily digest to a Slack-style webhook channel (a non-email action), through the same rails
data/leads.sample.json sample data so the demo runs with no setup
data/digest.sample.json sample data for the digest example
test/rails.test.js unit tests for the rails and the cadence
test/example-digest.test.js unit tests for the digest job on the same rails
schedule/ launchd, cron, and systemd templates + installer
.github/workflows/test.yml node:test on push/PR (Node 18/20/22)
License
MIT. Built and maintained by Digiton Dynamics .
Ten scheduled AI agents run my agency from launchd with zero dependencies. One Node file per agent; every send passes unit-tested rails (kill switch, allowlist, DNC, dedupe, daily cap). No n8n.
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

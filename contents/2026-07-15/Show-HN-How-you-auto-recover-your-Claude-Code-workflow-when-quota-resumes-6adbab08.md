---
source: "https://github.com/softcane/cc-session-recover"
hn_url: "https://news.ycombinator.com/item?id=48924870"
title: "Show HN: How you auto recover your Claude Code workflow when quota resumes"
article_title: "GitHub - softcane/cc-session-recover: AFK mode for Claude Code. Recover Claude Code sessions after quota or rate-limit stops. · GitHub"
author: "pradeep1177"
captured_at: "2026-07-15T18:56:00Z"
capture_tool: "hn-digest"
hn_id: 48924870
score: 1
comments: 0
posted_at: "2026-07-15T18:09:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: How you auto recover your Claude Code workflow when quota resumes

- HN: [48924870](https://news.ycombinator.com/item?id=48924870)
- Source: [github.com](https://github.com/softcane/cc-session-recover)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T18:09:47Z

## Translation

タイトル: HN の表示: クォータ再開時にクロード コード ワークフローを自動回復する方法
記事のタイトル: GitHub - Softcane/cc-session-recover: クロード コードの AFK モード。クォータまたはレート制限が停止した後にクロード コード セッションを回復します。 · GitHub
説明: クロード コードの AFK モード。クォータまたはレート制限が停止した後にクロード コード セッションを回復します。 - ソフトケーン/cc-session-recover

記事本文:
GitHub - Softcane/cc-session-recover: クロード コードの AFK モード。クォータまたはレート制限が停止した後にクロード コード セッションを回復します。 · GitHub
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
ソフトケーン
/
cc-セッション-回復
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
29 コミット 29 コミット .claude .claude .github/ workflows .github/ workflows bin bin docs docs lib lib scripts scripts templates templates .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json session-recover.yaml session-recover.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コードは午前1時に停止しました。クォータのリセットに成功しました。端末はまだ待機中です
一言で言えば：
続ける
cc-session-recover は、終了する前にクロード コードに回復計画を提供します。それ
StopFailure フック、/session-recover コマンド、HANDOFF.md をインストールします。
ノートブック、およびオプションのウォッチャー。クロード コードがレート制限に達したとき、または
過負荷が発生した場合、フックは失敗を記録します。再試行ウィンドウが経過すると、武装した
セッションはハンドオフを読み取り、タスクを続行します。
tmux のトリックはありません。隠れたプロンプト注入はありません。 OSハックはありません。
デモでは、Claude Code が 6 月 12 日に実際のクォータ ストップに達し、最後まで待機していることが示されています。
リセットして同じタスクを継続します。
クロード コード プロジェクトにインストールします。
npx cc-session-recover init /path/to/project
このリポジトリのクローンからインストールします。
bash スクリプト/install-into-project.sh /path/to/project
クロード コードは、次回そのフック コマンドを開くときにフック コマンドを承認するように求めます。
プロジェクト。インストーラーは既存の HANDOFF.md を保持するか、
セッション回復.yaml 。
フック設定をマージせずにファイルをコピーする場合は、 --no-hooks を渡します。
.claude/settings.local.json にコピーします。
プロジェクトでクロード コードを開始します。
cd /パス/プロジェクトへ
クロード
クロード・コードにタスクを与えてから、アームを回復します。
/セッション-回復
必要な場合は、カスタムのチェック間隔を使用します。
/セッション-リカバリ15
デフォルトの間隔は 30 分です

エス。クロード コードはセッション内で 1 つの繰り返しを作成します
次のプロンプトでスケジュールを設定します。
.claude/auto- continue.md を読み、それに従ってください。
クロード コードのセッションを開いたままにしておきます。スケジュールはタスクをチェックしながら、
セッションはアイドル状態になり、続行、待機、または停止します。
他の人は「クロードが落ち込んでいる」と言います。ランチ、会議、または一晩中保存します
プロジェクトにはすでに 3 つの回復状態があるため、そのまま実行されます。
HANDOFF.md は、何が変更され、次に何をすべきかをクロード コードに伝えます。
.claude/quota-blocked.json は、一時的な失敗と再試行ウィンドウを記録します。
リカバリ スケジュールはマーカーをチェックし、その後ハンドオフから再開します。
窓が開きます。
インストールされたランタイムは、次の 3 つの状態のいずれかを出力します。
WAIT <秒> <エラー> は、再試行ウィンドウが開いていないことを意味します。
READY <error> は、クロード コードがマーカーをクリアして続行できることを意味します。
NONE は保留中の障害が存在しないことを意味するため、クロード コードは HANDOFF.md をチェックします。
残りの仕事のために。
Claude Code はリカバリ スケジュールをキャンセルし、ハンドオフ後に DONE を出力します。
チェックリストは終了です。
別のシェルからウォッチャーを実行します。
npx cc-session-recover watch /path/to/project
ウォッチャーは同じマーカー ファイルを読み取り、次の内容で再開します。
クロード -p --resume < session_id >
in-session /session-recover スケジュールまたはウォッチャーを使用します。両方を実行しないでください
同じファイルを編集する 2 つのクロード コード セッションが必要な場合を除き、同じプロジェクトを使用します。
インストーラーは次の session-recover.yaml を作成します。
エラー:
- レート制限
- 過負荷
再試行分: 20
新規インストールは rate_limit とオーバーロードを再試行します。実行時ログ
server_error ですが、追加するまでそのエラーは再試行されません。
エラー:
- レート制限
- 過負荷
- サーバーエラー
再試行分: 10
Claude Code は、入力された StopFailure 名を選択します。クロード コードの一部のバージョン
制御された HTTP 529 応答を server_error としてレポートします。もしそうならその名前を追加してください
回復が欲しい

その分類のために。
レート制限の場合、キャッシュされた有効なリセット時間が retry_ minutes よりも優先されます。過負荷
サーバーエラーの回復には retry_ minutes を使用します。
.claude/session-recover.js 、ローカル ランタイム。
.claude/hooks/log-stop-failure.sh 、StopFailure フック。
.claude/commands/session-recover.md 、スラッシュコマンド。
.claude/auto- continue.md 、回復プロンプト。
.claude/statusline-quota-cache.sh 、レート制限リセット キャッシュ。
HANDOFF.md 、タスク ノートブック。
session-recover.yaml 、再試行構成。
また、ランタイム状態ファイルを .gitignore に追加するため、リカバリ ログとマーカーも追加されます。
リポジトリ履歴を入力しないでください。
このパッケージは、クォータまたはプロバイダーの障害をバイパスしません。
同一セッションのリカバリでは、クロード コード プロセスを開いたままにする必要があります。
認証、請求、無効なリクエスト、モデルが見つからない、および不明なエラー
再試行ループから抜け出してください。
ウォッチャーには、 jq 、 Node.js 、および PATH 上の claude コマンドが必要です。
完成した作品をレビューする必要があります。
簡単な流れ：
ノート、アラーム、ウォッチャー。
よくある質問:
フックの承認、失敗の種類、そして依然として人間が必要なもの。
詳細:
ウォッチャーのタイミング、リセット時間のキャッシュ、および動作制限。
クロード・コードのAFKモード。クォータまたはレート制限が停止した後にクロード コード セッションを回復します。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
7
v0.2.1
最新の
2026 年 6 月 26 日
+ 6 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AFK mode for Claude Code. Recover Claude Code sessions after quota or rate-limit stops. - softcane/cc-session-recover

GitHub - softcane/cc-session-recover: AFK mode for Claude Code. Recover Claude Code sessions after quota or rate-limit stops. · GitHub
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
softcane
/
cc-session-recover
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
29 Commits 29 Commits .claude .claude .github/ workflows .github/ workflows bin bin docs docs lib lib scripts scripts templates templates .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json session-recover.yaml session-recover.yaml View all files Repository files navigation
Claude Code stopped at 1 AM. The quota reset passed. Your terminal still waited
for one word:
continue
cc-session-recover gives Claude Code a recovery plan before you walk away. It
installs a StopFailure hook, a /session-recover command, a HANDOFF.md
notebook, and an optional watcher. When Claude Code hits a rate limit or
overload, the hook records the failure. After the retry window passes, the armed
session reads the handoff and continues the task.
No tmux tricks. No hidden prompt injection. No OS hacks.
The demo shows Claude Code hitting a real quota stop on June 12, waiting through
the reset, and continuing the same task.
Install into your Claude Code project:
npx cc-session-recover init /path/to/project
Install from a clone of this repo:
bash scripts/install-into-project.sh /path/to/project
Claude Code asks you to approve the hook command the next time you open that
project. The installer keeps an existing HANDOFF.md or
session-recover.yaml .
Pass --no-hooks if you want to copy the files without merging hook settings
into .claude/settings.local.json .
Start Claude Code in the project:
cd /path/to/project
claude
Give Claude Code the task, then arm recovery:
/session-recover
Use a custom check interval when you want one:
/session-recover 15
The default interval is 30 minutes. Claude Code creates one recurring in-session
schedule with this prompt:
Read .claude/auto-continue.md and follow it.
Keep that Claude Code session open. The schedule checks the task while the
session sits idle, then continues, waits, or stops.
Someone else says "Claude is down." You keep your lunch, meeting, or overnight
run intact because the project already has three pieces of recovery state:
HANDOFF.md tells Claude Code what changed and what to do next.
.claude/quota-blocked.json records the transient failure and retry window.
The recovery schedule checks the marker and resumes from the handoff after
the window opens.
The installed runtime prints one of three states:
WAIT <seconds> <error> means the retry window has not opened.
READY <error> means Claude Code can clear the marker and continue.
NONE means no pending failure exists, so Claude Code checks HANDOFF.md
for remaining work.
Claude Code cancels the recovery schedule and prints DONE after the handoff
checklist finishes.
Run the watcher from another shell:
npx cc-session-recover watch /path/to/project
The watcher reads the same marker file and resumes with:
claude -p --resume < session_id >
Use the in-session /session-recover schedule or the watcher. Do not run both on
the same project unless you want two Claude Code sessions editing the same files.
The installer creates this session-recover.yaml :
errors :
- rate_limit
- overloaded
retry_minutes : 20
New installs retry rate_limit and overloaded . The runtime logs
server_error , but it will not retry that error until you add it:
errors :
- rate_limit
- overloaded
- server_error
retry_minutes : 10
Claude Code chooses the typed StopFailure name. Some Claude Code versions
report controlled HTTP 529 responses as server_error ; add that name if you
want recovery for that classification.
For rate limits, a valid cached reset time wins over retry_minutes . Overload
and server-error recovery use retry_minutes .
.claude/session-recover.js , the local runtime.
.claude/hooks/log-stop-failure.sh , the StopFailure hook.
.claude/commands/session-recover.md , the slash command.
.claude/auto-continue.md , the recovery prompt.
.claude/statusline-quota-cache.sh , the rate-limit reset cache.
HANDOFF.md , the task notebook.
session-recover.yaml , the retry config.
It also adds runtime state files to .gitignore so recovery logs and markers do
not enter your repo history.
This package does not bypass quota or provider failures.
Same-session recovery needs the Claude Code process to stay open.
Authentication, billing, invalid-request, model-not-found, and unknown errors
stay out of the retry loop.
The watcher needs jq , Node.js, and the claude command on PATH .
You still need to review the finished work.
Simple flow :
notebook, alarm, and watcher.
FAQ :
hook approval, failure types, and what still needs a human.
Full details :
watcher timing, reset-time cache, and operational limits.
AFK mode for Claude Code. Recover Claude Code sessions after quota or rate-limit stops.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
7
v0.2.1
Latest
Jun 26, 2026
+ 6 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

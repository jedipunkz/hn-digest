---
source: "https://github.com/sand0vvv/fleet"
hn_url: "https://news.ycombinator.com/item?id=49024833"
title: "Show HN: Fleet – drive a fleet of Claude Code/Codex agents from Telegram"
article_title: "GitHub - sand0vvv/fleet: A fleet of AI coding agents in your Telegram supergroup. Self-hosted, local, no account. · GitHub"
author: "sand0vvv"
captured_at: "2026-07-23T17:09:32Z"
capture_tool: "hn-digest"
hn_id: 49024833
score: 1
comments: 0
posted_at: "2026-07-23T17:03:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Fleet – drive a fleet of Claude Code/Codex agents from Telegram

- HN: [49024833](https://news.ycombinator.com/item?id=49024833)
- Source: [github.com](https://github.com/sand0vvv/fleet)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T17:03:00Z

## Translation

タイトル: HN を表示: フリート – Telegram からクロード コード/コーデックス エージェントのフリートを運転します
記事のタイトル: GitHub - Sand0vvv/fleet: Telegram スーパーグループ内の AI コーディング エージェントのフリート。自己ホスト型、ローカル、アカウントなし。 · GitHub
説明: Telegram スーパーグループ内の AI コーディング エージェントのフリート。自己ホスト型、ローカル、アカウントなし。 - サンド0vvv/艦隊
HN テキスト: これは休日に作成しました。ベッドから起きたくなかったので、エージェント全員が入ったラップトップが部屋の向こう側にありました。それで私は彼らを私のところに来させました。フリートは、1 つの Telegram スーパーグループから AI コーディング エージェント (Claude Code または Codex) のフリートを実行します (各エージェントは独自のトピック内にあります)。音声またはテキスト、コード/差分/ファイルがチャットに戻ってきます。ヘッドレス (メッセージごとに起動)、または実行中に操作するライブ CLI セッション。すべては自分のマシン上で実行されます。 Telegram は単なるリモートであり、ロングポーリング経由で到達するため、あらゆる NAT の背後で動作します。 1か月間、毎日すべてを実行してきました。時期尚早で、ところどころ荒っぽいところもありますが、何でも喜んでお答えします。

記事本文:
GitHub - Sand0vvv/fleet: Telegram スーパーグループ内の AI コーディング エージェントのフリート。自己ホスト型、ローカル、アカウントなし。 · GitHub
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
コードの品質 マージ時に品質を強制する
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
砂0vvv
/
艦隊
公共
通知
通知を変更するにはサインインする必要があります

イオン設定
追加のナビゲーション オプション
コード
スタンドアロン ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット アセット アセット codex-shell codex-shell mcp mcp src src .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Telegram スーパーグループ内の AI コーディング エージェントのフリート。
自己ホスト型、ローカル、アカウントなし。
Fleet を使用すると、単一の Telegram スーパーグループ内で AI コーディング エージェント (Claude Code と Codex) のチーム全体を携帯電話から操作できます。各エージェントはグループ内で独自のトピックを取得します。あなたがそれらにテキスト (または音声メッセージ) を送信すると、それらはあなたのマシン上で動作し、コード、ファイル、スクリーンショット、質問、進捗状況などのテキストを返します。
すべてがローカルで実行されます。クラウド サービスもサインアップもダッシュボードもありません。 Telegram ボット トークンと API キーは、マシンの ~/.fleet/config.json に保存されます。 Telegram はロングポーリング経由で到達します。NAT の背後ではパブリック URL も Webhook も機能しません。
公開されてまだ初期の段階で、荒削りな部分が予想されます。問題やPRを歓迎します。
0. 前提条件 — Node.js ≥ 18 および Claude コードがインストールされ、ログインしていること (PATH に claude)。コーデックスはオプションです。
1. ボットを作成します — Telegram: /newbot で @BotFather に話しかけ、名前を選択し、トークンをコピーします。
2. グループを作成します — Telegram グループを作成し、トピック (グループ設定 → トピック) を有効にし、ボットを管理者として追加します (トピックを管理する必要があります)。
npm i -g @sand0vvv/フリート
フリート初期化 # ボット トークン、音声プロバイダー、デフォルトを要求します
フリートドクター # 青信号 = 準備完了
フリートスタート # ランナー — このターミナルを開いたままにしておきます
4. リンク — スーパーグループで /link を送信します。ランナーは自分自身をあなた (Telegram ユーザー ID) とそのグループにバインドします。それ以降はあなただけが艦隊を指揮できるようになります。
5. エージェントを生成します —

グループの一般トピック:
/spawn C:\パス\あなたの\プロジェクトへ
フォルダーにちなんだ名前のトピックが表示され、マシン上でターミナル ウィンドウが開き、エージェントが稼働します。トピックにテキスト メッセージを送信します。エージェントがそこで応答します。または、マシン上のターミナルから: cd で任意のプロジェクトに移動し、フリート クロード (またはフリート コデックス) を実行します。
電報スーパーグループ
§─ 一般 ← フリート コマンド: /spawn /list /usage …
§─ 🟢 my-webapp ← クロード コード エージェント、実行中
§─ 🟢 my-webapp-codex ← 同じフォルダー内の Codex エージェント
└─ 🟡 data-pipeline ← パーク中 (次のメッセージで起動します)
トピック タイトルにはライブ ステータスが表示されます: 🟢 実行中、🟡 駐車中、🔴 停止/クラッシュ。エージェントがメッセージを処理している間、トピックには「入力中…」と表示されます。
すべてのエージェントはマシン上で実行されます。クロード エージェントは実際のターミナル ウィンドウ内に存在し、監視したり入力したりできます。ウィンドウを閉じるとエージェントがパーク (リソースを解放) します。エージェントは次のメッセージで自動的に起動し、セッションが継続されます ( --resume / -- continue )。 Codex エージェントは独自のコンソール ウィンドウを開き、ウェイク時に最後のセッションも再開します。新しいセッションは、セッション ( /new ) を要求した場合にのみ発生します。
これらを Telegram で送信します (ボットの / メニューにも表示されます):
エージェントのトピックに入力した他のスラッシュ コマンド ( /cost 、 /doctor 、 /review など) は、ネイティブ コマンドとしてエージェントの端末に直接転送されます。
音声 : エージェントのトピックで音声メッセージを送信します。音声メッセージは文字に起こされ (Groq / OpenAI / ローカル Whisper、フリート init で選択)、テキストとして配信されます。
ファイル : エージェントはファイルとスクリーンショットをトピックに送り返します。トピックに添付したものはすべて、エージェントの .inbox/ にダウンロードされ、エージェントに渡されます。
フリート初期化ボット トークン、ウィスパー、デフォルト設定 (再実行は安全、シークレットはマスクされて表示)
フリートを開始 フリートを実行します (単一インスタンス - 古いランナーは repl です)

自動的にエースされます）
フリートストップ 走行中のフリートを停止します
フリート更新停止 → npm i -g 最新 → 再起動、1 つのコマンドで
フリート ログ [n] ランナー ログの最後の n 行
フリート自動起動オン|オフ ログオン時にフリートを開始します (Windows)
フリート クロードは現在のフォルダーにクロード エージェントを生成します
フリート コーデックスは、現在のフォルダーのコーデックス エージェントを生成します
フリートドクターチェックノード/クロード/コーデックス/トークン/リンク
フリートステータスは現在の構成とバージョンを表示します
すべてのコマンドは、npm 上に新しいバージョンがあることを示します。ランナーもあなたのグループに一言を投下します。
テレグラム ⇄ ロングポール ⇄ フリートランナー (あなたのマシン)
§─ localhost HTTP+WS (以下の MCP の場合 - 何も公開されません)
§─ クロード エージェント → ノード-pty ターミナル + フリート MCP (send_message/send_file,
│ 受信メッセージのチャネル インジェクション)
━─ Codex エージェント → 独自のコンソール ウィンドウ。 codex-shell は公式を駆動します
メッセージをライブで挿入するための「codex app-server」JSON-RPC
ランナーは Telegram Bot API をロングポーリングします。受信ポート、Webhook、パブリック URL はありません。
エージェントのトピック内のメッセージは、そのエージェントのローカル プロセスにルーティングされます。
各エージェントには小規模なフリート MCP があり、ローカルホスト上のランナーにダイヤルバックして返信とファイルを Telegram にプッシュします。受信メッセージはライブ セッションに挿入されます (クロード: チャネル通知、コーデックス: アプリサーバー プロトコルを経由して開始/開始)。
ランナーはエージェントの返信を適切なトピックに投稿します。
双方向の信頼性。パーク/オフラインのエージェントへのメッセージはディスク上のキューに入れられ、エージェントが復帰したときに配信されます。ウィンドウを閉じても何も失われません。エージェントからの返信とファイルは、ネットワークの障害でも同様に存続します (ディスク上の送信ボックス、順番に再試行されます)。高速の連続メッセージは 1 つの配信にデバウンスされます (構成可能、0 = 即時)。クラッシュしたエージェントは自動再起動され (クラッシュ ループが保護され)、エージェントが実行された場合、

er 自体が停止すると、ダウンする前にクラッシュ レポートがグループに投稿されます。
State : すべては ~/.fleet/ に存在します — config.json (トークン、所有者、グループ、デフォルト)、state.json (エージェント/トピック/セッション)、配信キュー、ログ。シークレットはプロジェクト ディレクトリには決して触れないため、誤ってコミットされることはありません。フォルダーを削除して工場出荷時の状態にリセットします。 (構成パスを FLEET_CONFIG 環境変数でオーバーライドします。)
フリートドクターが最初にすべてをチェックします。フリート ログには、ランナーの最近のログが表示されます。
ポートビジー: フリートスタートは古いランナーを強制終了し、自動的に引き継ぎます。非フリートプロセスがポートを保持している場合は、 ~/.fleet/config.json に「port」を設定します。
ボットが反応しない: ボットがグループの管理者であり、トピックが有効になっていることを確認してください。次に /link を再度実行します。
ターミナル ウィンドウが表示されない: フリート スタートを実際のコンソール (バックグラウンド サービスではない) で実行します。 /show は再試行し、ウィンドウが失敗したかどうかを正直に通知します。
Windows : git-bash がインストールされている場合はエージェントがそれを使用し、そうでない場合は通常の cmd.exe を使用します。どちらも機能します。
支店
何
スタンドアロン
このバージョン - ローカル、npm、ロングポーリング。使うもの。
レガシー
オリジナルのクラウド アーキテクチャ (FastAPI バックエンド + Postgres + Webhook)
レガシードッカー
エージェントを Docker サンドボックスに閉じ込める従来の亜種
著者
Telegram スーパーグループ内の AI コーディング エージェントのフリート。自己ホスト型、ローカル、アカウントなし。
フリート.sand0vvv.workers.dev
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
16
フリート v2.0.1
最新の
2026 年 7 月 22 日
+ 15 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A fleet of AI coding agents in your Telegram supergroup. Self-hosted, local, no account. - sand0vvv/fleet

I built this on a day off - didn't want to get out of bed, and the laptop with all my agents was across the room. So I made them come to me. fleet runs a fleet of AI coding agents (Claude Code or Codex) from one Telegram supergroup - each agent in its own topic. Voice or text, code/diffs/files come back to the chat. Headless (fires per message) or a live CLI session you steer mid-run. Everything runs on your own machine. Telegram is just the remote, reached via long-polling so it works behind any NAT. Been running everything through it daily for a month. It's early and rough in places — happy to answer anything.

GitHub - sand0vvv/fleet: A fleet of AI coding agents in your Telegram supergroup. Self-hosted, local, no account. · GitHub
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
Code Quality Enforce quality at merge
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
sand0vvv
/
fleet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
standalone Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits assets assets codex-shell codex-shell mcp mcp src src .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
A fleet of AI coding agents in your Telegram supergroup.
Self-hosted · local · no account.
Fleet lets you drive a whole team of AI coding agents — Claude Code and Codex — from your phone, in a single Telegram supergroup. Each agent gets its own topic in the group. You text (or voice-message) them, they work on your machine and text back: code, files, screenshots, questions, progress.
Everything runs locally . No cloud service, no sign-up, no dashboard. Your Telegram bot token and API keys stay on your machine in ~/.fleet/config.json . Telegram is reached via long-polling — no public URL, no webhook, works behind any NAT.
Built in public and still early — expect rough edges. Issues and PRs welcome.
0. Prerequisites — Node.js ≥ 18 and Claude Code installed and logged in ( claude in your PATH). Codex is optional.
1. Make a bot — talk to @BotFather in Telegram: /newbot , pick a name, copy the token.
2. Make the group — create a Telegram group, enable Topics (group settings → Topics), and add your bot as admin (it needs to manage topics).
npm i -g @sand0vvv/fleet
fleet init # asks for the bot token, voice provider, defaults
fleet doctor # green lights = ready
fleet start # the runner — keep this terminal open
4. Link — send /link in your supergroup. The runner binds itself to you (your Telegram user id) and that group . From then on only you can command the fleet.
5. Spawn an agent — in the group's General topic:
/spawn C:\path\to\your\project
A topic appears named after the folder, a terminal window opens on your machine, and the agent is live. Text the topic — the agent answers there. Or from a terminal on the machine: cd into any project and run fleet claude (or fleet codex ).
Telegram supergroup
├─ General ← fleet commands: /spawn /list /usage …
├─ 🟢 my-webapp ← Claude Code agent, running
├─ 🟢 my-webapp-codex ← Codex agent in the same folder
└─ 🟡 data-pipeline ← parked (wakes on your next message)
Topic titles carry live status: 🟢 running · 🟡 parked · 🔴 stopped/crashed. While an agent works on your message, the topic shows typing… .
Every agent runs on your machine. Claude agents live in a real terminal window you can watch and type into; close the window and the agent parks (frees resources) — it wakes automatically on your next message, with its session continued ( --resume / --continue ). Codex agents open their own console window and also resume their last session on wake. A fresh session only happens when you ask for one ( /new ).
Send these in Telegram (the bot's / menu shows them too):
Any other slash command you type in an agent's topic (e.g. /cost , /doctor , /review ) is forwarded straight into the agent's terminal as a native command.
Voice : send a voice message in an agent's topic — it's transcribed (Groq / OpenAI / local Whisper, your choice in fleet init ) and delivered as text.
Files : agents send files and screenshots back into their topic; anything you attach in the topic is downloaded into the agent's .inbox/ and handed to it.
fleet init set up bot token, whisper, defaults (re-run safe; secrets shown masked)
fleet start run the fleet (single instance — a stale runner is replaced automatically)
fleet stop stop the running fleet
fleet update stop → npm i -g latest → restart, in one command
fleet logs [n] last n lines of the runner log
fleet autostart on|off start the fleet at logon (Windows)
fleet claude spawn a Claude agent for the current folder
fleet codex spawn a Codex agent for the current folder
fleet doctor check node/claude/codex/token/link
fleet status show current config + version
Every command tells you when a newer version is on npm; the runner also drops a one-liner into your group.
Telegram ⇄ long-poll ⇄ fleet runner (your machine)
├─ localhost HTTP+WS (for the MCP below — nothing exposed)
├─ Claude agents → node-pty terminal + fleet MCP (send_message/send_file,
│ channel injection for incoming messages)
└─ Codex agents → own console window; codex-shell drives the official
`codex app-server` JSON-RPC to inject your messages live
The runner long-polls the Telegram Bot API — no inbound ports, no webhook, no public URL.
A message in an agent's topic is routed to that agent's local process.
Each agent has a small fleet MCP that dials the runner back on localhost to push replies and files up to Telegram; incoming messages are injected into the live session (Claude: channel notification; Codex: turn/start over the app-server protocol).
The runner posts the agent's reply into the right topic.
Reliability, both directions. Messages to a parked/offline agent are queued on disk and delivered when it wakes — nothing is lost if you closed the window. Replies and files from agents survive network hiccups the same way (on-disk outbox, retried in order). Rapid consecutive messages are debounced into one delivery (configurable, 0 = instant). A crashed agent is auto-restarted (crash-loop guarded), and if the runner itself dies it posts a crash report into your group before going down.
State : everything lives in ~/.fleet/ — config.json (token, owner, group, defaults), state.json (agents/topics/sessions), delivery queues, logs. Secrets never touch a project directory, so they can't be committed by accident. Delete the folder to factory-reset. (Override the config path with the FLEET_CONFIG env var.)
fleet doctor first — it checks everything. fleet logs shows the runner's recent log.
Port busy : fleet start kills a stale runner and takes over by itself. If a non-fleet process holds the port, set "port" in ~/.fleet/config.json .
Bot doesn't react : make sure the bot is an admin of the group and Topics are enabled; then /link again.
No terminal window pops : run fleet start in a real console (not a background service); /show retries and tells you honestly if the window failed.
Windows : if git-bash is installed agents use it, otherwise plain cmd.exe — both work.
Branch
What
standalone
this version — local, npm, long-polling. The one to use.
legacy
the original cloud architecture (FastAPI backend + Postgres + webhook)
legacy-docker
legacy variant that jails agents in a Docker sandbox
Author
A fleet of AI coding agents in your Telegram supergroup. Self-hosted, local, no account.
fleet.sand0vvv.workers.dev
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
16
fleet v2.0.1
Latest
Jul 22, 2026
+ 15 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

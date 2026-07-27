---
source: "https://github.com/radres/call-me"
hn_url: "https://news.ycombinator.com/item?id=49076793"
title: "Show HN: Call Me, your AI agents ring you with no extra carrier costs"
article_title: "GitHub - radres/call-me: Let your AI agents call and text your actual iPhone — Claude Code plugin + agent skill · GitHub"
author: "radres"
captured_at: "2026-07-27T23:55:46Z"
capture_tool: "hn-digest"
hn_id: 49076793
score: 2
comments: 0
posted_at: "2026-07-27T23:10:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Call Me, your AI agents ring you with no extra carrier costs

- HN: [49076793](https://news.ycombinator.com/item?id=49076793)
- Source: [github.com](https://github.com/radres/call-me)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T23:10:00Z

## Translation

タイトル: HN を表示: Call Me、AI エージェントが追加の通信料金なしで電話をかけます
記事タイトル: GitHub - radres/call-me: AI エージェントに実際の iPhone に電話とテキスト メッセージを送信させます — クロード コード プラグイン + エージェント スキル · GitHub
説明: AI エージェントに実際の iPhone に電話とテキスト メッセージを送信させます — クロード コード プラグイン + エージェント スキル - radres/call-me
HN テキスト: 私が取り組んでいるただの楽しいアプリです。必ずしも AI のみである必要はなく、あらゆる AI ハーネスまたは API として機能します。

記事本文:
GitHub - radres/call-me: AI エージェントに実際の iPhone に電話とテキスト メッセージを送信させます — クロード コード プラグイン + エージェント スキル · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
レーダー
/
電話してください
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .claude-plugin .claude-plugin claude-code claude-code skill skill README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントは、実際の iPhone を鳴らしたり、質問をしたり、メッセージを取得したりできます。
音声による応答をテキストとして返信することも、テキストメッセージで送信することもできます。このリポジトリにはエージェント側があります。
Call Me iOS アプリ (旧 AI Phone) の統合。
App Store で Call Me — 無料です。
それを開いて [同意して続行] をタップすると、個人用の Call Me が表示されます。
番号 : 10 桁。エージェントが連絡するために必要なすべての番号。
すでにターミナルで CLI を使用していますか? callme qr は、スキャン可能な QR コードを印刷します。
そのリンクをクリックすると、callme setup がこれらの手順を出力します。
2. クロードコードを接続する（推奨）
これらを 2 つの別々のメッセージとして送信します (スラッシュ コマンドは 1 回につき 1 つだけ実行されます)
メッセージ):
/プラグインマーケットプレイス追加radres/call-me
/プラグインのインストール call-me@call-me
次に、クロードに「Call Me を設定して」と伝えてください。からの番号を要求します。
アプリはそれを記憶し、確認テキストを送信します。それ以来、クロードは
テキストメッセージを送信し、電話を鳴らすと、アプリからの返信がすぐに返されます
ライブセッションへ。
プロンプトが表示されたら、プラグインの設定フィールドに番号を貼り付けることもできます。
同じ結果。どちらの方法でも ~/.aiphone/config.json に配置されるため、すべての Claude
マシン上のセッションが電話に到達し、新しい電話との再ペアリングに時間がかかります
何も再起動しなくてもすぐに効果が現れます。
3. 他のエージェント (Codex、スクリプト、cron ジョブなど)
skill/ でスタンドアロン スキルを使用します。これは単一の bash+curl CLI です。
git clone https://github.com/radres/call-me
cp call-me/skill/callme ~ /bin/ # または PATH 上の任意の場所
callme ペア 5551234567 # 1 回 — アプリからの番号
callme call " ステージングをデプロイします

最初に製品を作りますか？ " # 音声で応答するまでブロックします
callme text " ビルドグリーン、PR マージ。 " # プッシュ通知
スキルをサポートするエージェントは、skill/SKILL.md をポイントできます。
直接 — 完全なフロー (セットアップ、ペアリング、通話、テキストメッセージ、リスニング、タイトル) を教えます。
これを CLAUDE.md / エージェントの指示に追加します。
Call Me を通じて電話で連絡が取れます。
クロード コード: github.com/radres/call-me から call-me プラグインをインストールします。
まだペアリングされていない場合は、セットアップ ツールを実行します。
他のエージェント: callme CLI (スキル/ディレクトリ) を使用します - 「callme setup」で説明されています
ペアリングしてから、「callme call "question"」を押すと呼び出し音が鳴り、私の言ったことを返します。
そこに数字がないことに注意してください。数字は資格情報なので、
~/.aiphone/config.json (モード 0600) をメモの代わりにします。電話番号
それを読み戻し、callme ペア <number> がそれを変更します。
callme call は、Call Me サービスに質問を投稿します。あなたのiPhone
実際の電話のように CallKit を通じて呼び出し音が鳴り、TTS が質問を話します。
音声による応答は文字に起こされ、エージェントに返されます。通話ブロック
あなたが答えるまで、それが重要です。
callme text はプッシュ通知メッセージを送信します。電話からの返事
チャネルが有効になっているエージェント セッションに戻されます (Claude Code プラグイン)
または callme listen / callme events で取得されます。
各エージェント セッションでは独自の番号とスレッドが登録されるため、電話機には
タスクごとに別々の、タイトル付きの会話。
AI エージェントに実際の iPhone に電話とテキスト メッセージを送信させます — クロード コード プラグイン + エージェント スキル
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Let your AI agents call and text your actual iPhone — Claude Code plugin + agent skill - radres/call-me

Just a fun app I've been working on. Not necessarily AI only and works with any AI harness or just as an API.

GitHub - radres/call-me: Let your AI agents call and text your actual iPhone — Claude Code plugin + agent skill · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
radres
/
call-me
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .claude-plugin .claude-plugin claude-code claude-code skill skill README.md README.md View all files Repository files navigation
Your AI agents can ring your actual iPhone, speak a question, and get your
spoken answer back as text — or just text you. This repo has the agent-side
integrations for the Call Me iOS app (formerly AI Phone).
Call Me on the App Store — free.
Open it, tap Agree & Continue , and it shows your personal Call Me
number : 10 digits, and all an agent needs to reach you.
Already in a terminal with the CLI? callme qr prints a scannable QR code for
that link, and callme setup prints these steps for you.
2. Connect Claude Code (recommended)
Send these as two separate messages (slash commands only run one per
message):
/plugin marketplace add radres/call-me
/plugin install call-me@call-me
Then just tell Claude: "set up Call Me" . It asks for the number from the
app, remembers it, and sends you a confirmation text. From then on Claude can
text you and ring your phone, and your replies from the app flow straight back
into the live session.
You can also paste the number into the plugin's config field when prompted —
same result. Either way it lands in ~/.aiphone/config.json , so every Claude
session on the machine reaches your phone, and re-pairing to a new phone takes
effect immediately without restarting anything.
3. Any other agent (Codex, scripts, cron jobs, …)
Use the standalone skill in skill/ . It's a single bash+curl CLI:
git clone https://github.com/radres/call-me
cp call-me/skill/callme ~ /bin/ # or anywhere on PATH
callme pair 5551234567 # once — the number from the app
callme call " Deploy staging or prod first? " # blocks until you answer by voice
callme text " Build green, PR merged. " # push notification
Agents that support skills can point at skill/SKILL.md
directly — it teaches the full flow (setup, pair, call, text, listen, title).
Drop this in your CLAUDE.md / agent instructions:
I'm reachable on my phone through Call Me.
Claude Code: install the call-me plugin from github.com/radres/call-me,
then run its setup tool if we aren't paired yet.
Other agents: use the callme CLI (skill/ dir) — `callme setup` explains
pairing, then `callme call "question"` rings me and returns what I say.
Notice there's no number in there: the number is a credential, so it lives in
~/.aiphone/config.json (mode 0600) instead of your notes. callme number
reads it back, callme pair <number> changes it.
callme call POSTs your question to the Call Me service; your iPhone
rings through CallKit like a real call, TTS speaks the question, your
spoken reply is transcribed and returned to the agent. The call blocks
until you answer — that's the point.
callme text sends a push-notification message; replies from the phone
are delivered back into channel-enabled agent sessions (Claude Code plugin)
or fetched with callme listen / callme events .
Each agent session registers its own number and thread, so your phone shows
separate, titled conversations per task.
Let your AI agents call and text your actual iPhone — Claude Code plugin + agent skill
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

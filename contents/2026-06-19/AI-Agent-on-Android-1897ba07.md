---
source: "https://github.com/ExTV/rikkahub-agent"
hn_url: "https://news.ycombinator.com/item?id=48602469"
title: "AI Agent on Android"
article_title: "GitHub - ExTV/rikkahub-agent: RikkaHub Agent -- is RikkaHub fork that have Full agent mode . · GitHub"
author: "excp"
captured_at: "2026-06-19T21:30:03Z"
capture_tool: "hn-digest"
hn_id: 48602469
score: 2
comments: 1
posted_at: "2026-06-19T19:47:51Z"
tags:
  - hacker-news
  - translated
---

# AI Agent on Android

- HN: [48602469](https://news.ycombinator.com/item?id=48602469)
- Source: [github.com](https://github.com/ExTV/rikkahub-agent)
- Score: 2
- Comments: 1
- Posted: 2026-06-19T19:47:51Z

## Translation

タイトル: Android 上の AI エージェント
記事のタイトル: GitHub - ExTV/rikkahub-agent: RikkaHub Agent -- フル エージェント モードを備えた RikkaHub フォークです。 · GitHub
説明: RikkaHub Agent -- フル エージェント モードを持つ RikkaHub フォークです。 - ExTV/rikkahub-agent

記事本文:
GitHub - ExTV/rikkahub-agent: RikkaHub Agent -- フル エージェント モードを持つ RikkaHub フォークです。 · GitHub
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
ExTV
/
リッカハブエージェント
公共
rikkahub/rikkahub から分岐
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,956 コミット 2,956 コミット .agents/ スキル .agents/ スキル .github .github ai ai アプリ app common common docs docs document document gradle gradle ハイライト ハイライト local-llm local-llm locale-tui locale-tui マテリアル 3 マテリアル 3 検索 検索 音声 音声 Web-UI Web-UI Web Web ワークスペース ワークスペース .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules ライセンス ライセンス README.md README.md build.gradle.kts build.gradle.kts bun.lock bun.lock gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat package.json package.json settings.gradle.kts settings.gradle.kts skill-lock.json skill-lock.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ネイティブ Android LLM チャット クライアントを実際のオンデバイス エージェントに変える RikkaHub のフォーク。80 以上のデバイス ツール、AI が作成したワークフロー、スケジュールされたジョブ、AI が駆動するアプリ内ブラウザ、SSH、画面自動化、ファイル マネージャー、音楽プレーヤー、音声文字起こし、ダウンロード可能なオンデバイス LLM、およびリモート Telegram ボット。すべてオプトイン。
ウェブサイト ·
ダウンロード・
特徴・
クイックスタート ·
ビルドする
Vanilla LLM チャット アプリは質問に答えることができます。アプリを開いたり、メッセージを送信したり、通知を確認したり、スケジュールされたジョブを実行したり、サーバーに SSH 接続したりすることはできません。 RikkaHub エージェントはできます。平易な言葉で何をすべきかを伝え、立ち去ると、バックグラウンドで、あなたの携帯電話上で、あなたの条件に応じて実行されます。
「毎週平日午前9時に、未読のWhatsAppを1つのTelegramメッセージにまとめてください。」
「ホームサーバーのディスクがいっぱいになったら、私に連絡してください。」
「通知を見てください。上司から何かが入ったら、テレグラムに転送してください。残りは静かに無視してください。」
「携帯電話で「請求書」について言及している PDF を見つけて、最初の段落を読んでください

ふ。」
「次の 4 時間、30 分ごとにスクリーンショットを撮って、午後ずっと実際に何をしたかを確認します。」
「Termux を使用して、できることをすべてリストした Web ページを作成し、ブラウザで開きます。」
「午後 7 時以降に自宅の Wi-Fi にヘッドフォンを接続すると、夜のプレイリストが始まります。」
「ルーターの管理ページを開いて、保存したパスワードでサインインし、現在どのデバイスが最も帯域幅を消費しているかを教えてください。」
「2 つの調査を並行して実行します。1 つは今月の東京行きの片道航空券の最安値を見つけるもので、もう 1 つは渋谷の 100 ドル以下のホテルをリストするものです。両方終わったら教えてください。」
これらはそれぞれ 1 行のセットアップです。あなたが日常生活を送っている間、電話はバックグラウンドでそれらを実行します。
AI に、タップ、スワイプ、スクロール、入力、スクリーンショットの撮影、アプリの起動、懐中電灯のオン、明るさや音量の変更、通知の投稿、振動、何かの共有、バッテリー、WiFi、信号、位置、センサー、連絡先、SMS の読み取りを依頼します。また、SMS の送信、壁紙の設定、NFC タグの読み書き、Android キーストアによるデータの署名と暗号化、外部ストレージや SD カードへのアクセス、アーカイブの圧縮または解凍も可能です。 80 を超えるツールはすべて Android に組み込まれており、追加のアプリは必要ありません。それぞれのスイッチをオンにするまでオフのままです。
どこからでもアシスタントに話しかけることができます。プライベート Telegram ボットを 1 分でセットアップし、連絡先のようにチャットします。質問、写真、PDF、または音声メモを送信してください。仕事中、睡眠中、運転中などに代わって実行できます。承認プロンプトでは、チャット内の単純な [はい/いいえ] ボタンを使用します。 AI が何か質問する必要がある場合、タップ可能な多肢選択式の質問がチャット内に表示されます (または自由テキストでの回答が得られます)。 1 つのメッセージに対して長すぎる応答は、切り詰められるのではなくダウンロード可能なファイルとして到着し、メッセージ バーストのペースが調整されるため、Telegram はレート制限を行いません。

答えはあなたの下から出ています。
エージェントには、アプリに実際のブラウザが組み込まれています。 URL を開いて、Cookie バナーをクリックし、検索ボックスに入力し、スクロールしてページを読み上げる様子を見てください。または、用事で Telegram から送信します。ステップごとに新しいスクリーンショットがチャットにストリーミングされます。ブラウザ画面にはフローティング チャット ピルがあるため、ページを離れることなく AI と会話を続けることができます。組み込みの記事抽出とアクション後の差分により、長時間の閲覧セッションでもトークンのコストが低く抑えられます。
Tasker スタイルの自動化ですが、AI がルールを作成します。トリガーとアクションを次のように説明するだけです。「家に帰ったら、着信音を消します」。 「平日の午前 8 時にバッテリーが 50% を超えていれば、メールをチェックして、何か緊急のことがあれば連絡してください。」 19 のトリガー (WiFi、Bluetooth、ヘッドフォン、ジオフェンス、アプリの起動、受信した通知、時間、充電、画面のオン/オフなど) と 14 の条件 (バッテリーのしきい値、日の出/日の入り、曜日、現在のフォアグラウンド アプリ、画面の状態) によって、それぞれがいつ起動されるかを決定します。レシーバーはワークフローで実際に必要な場合にのみ登録されるため、バッテリーの消耗は最小限に抑えられます。
タスクをスケジュールに従って実行するように設定し、忘れてしまいましょう。 「毎週月曜日の朝8時」、「2時間ごと」、「来週の金曜日の午後3時」。電話機は再起動とバッテリーセーバーによってすべてを実行し続けます。各タスクの起動方法を選択します。AI にその時点で考えさせ、何をするかを決定させます (「X を監視し、Y なら ping を送信」に適しています)。または、AI トークンを使用せずに実行される固定アクションを事前にベイクします (単純なリマインダーに適しています)。
AI には独自のファイル マネージャーがあります。ファイルの検索、読み取り、新しいファイルの保存、コピー、移動、名前変更、削除。通常のファイルマネージャーで行うことと同じですが、必要なものを記述するとそれが実行されます。 「携帯電話で「請求書」に言及しているすべての PDF を検索する」は 1 つの文で機能します。システムフォルダー

あなたに属していない人は、たとえ尋ねても立ち入り禁止です。
サーバーを一度保存​​すると、AI はオンデマンドでどのサーバーにも SSH 接続できるようになります。コマンドを実行し、ファイルをアップロードし、バックアップを取得し、ディスク容量を確認し、ログを記録します。パイプ入力をコマンドに直接入力するか、リモート ファイルを 1 回で書き込み、バックグラウンドで長時間実行されるサーバーを起動して、呼び出しが接続をハングアップせずにプロセス ID を返すようにします。 WiFi または携帯電話のどちらに接続していても機能します。端末を開かずにコーヒーショップから自宅サーバーを監視します。
Termux がインストールされている場合、AI は、パッケージのインストール、ソフトウェアの構築、スクリプトの実行、コマンドが返された後も実行を続けるバックグラウンド サービスの開始など、実際の Linux コマンドを携帯電話上で実行できます。専用の [設定] → [Termux] ページを使用すると、長時間にわたるインストールまたはビルドでより多くのスペースが必要な場合に、コマンド タイムアウト、ターンごとの時間バジェット、およびその他の Termux 制限を調整できます。さらに、Telegram で送信した音声メモは自動的に文字に起こされます。すべてが携帯電話上で実行され、クラウド文字起こしや API キー、インターネットは必要ありません。
音楽をリクエストすると、AI が Android の通常のメディア コントロール (ロック画面アート、ヘッドフォン キー、作品など) を通じて音楽を再生します。会議の一時停止、再開、音量の下げ、後で元に戻すなどの操作はすべてチャットまたはテレグラムから行えます。強制停止後でも、AI はスナップショット フォールバックを介して、中断した場所、同じトラック、同じ位置を再開できます。いいえ、「あなたがプレイヤーを殺したので、プレイヤーは永遠に消えてしまいます」。あなたのキューは存続します。
Markdown スキル ファイルをアプリにドロップすると、AI はステップごとに従う新しいプレイブックを取得します。連絡先に自動返信したり、通知スタックを要約したり、結果がアプリ内ブラウザーで直接開く JavaScript ミニアプリを実行したりします。バンドルされた注目のカタログには、QR ジェネレーター、Wikipedia クエリ ボックス、演奏できるピアノ、インターラが同梱されています。

アクティブマップなど。 2 つのスキルがすぐに有効になっています。1 つは、アシスタントがセッション間でプロアクティブに自己改善を続ける常時稼働のエージェント プレイブックで、もう 1 つは OpenClaw スキルをここで実行するように適応させるコンバータです。 URL、アプリ内で共有するマークダウン ファイル、またはバンドルされたカタログから選択して、新しいスキルを追加します。
長時間のタスクの場合、メイン アシスタントは、必要に応じてより小型で安価なモデル上で、集中したサブエージェントをクリーンなサイド コンテキストにディスパッチできます。 2 つ以上が並行して実行されます。1 つはトピックを調査し、もう 1 つはサーバーを更新します。各結果は 1 つの概要として返されるため、メインのチャットが無関係なツールの出力に埋もれてしまうことはありません。また、/stop は 1 ティックですべてのアクティブな子をカスケードしてキャンセルします。
アプリに組み込まれた健康診断。 [設定]、[Doctor] の順にタップすると、権限、バックグラウンド サービス、データベースの整合性、ネットワーク、Termux、診断の監査が上から下まで実行されます。何か足りない？行の横にある自動修正ボタンをタップして、権限を付与するか、サービスを再起動するか、チャット検索インデックスを再構築します。リモート トラブルシューティングのために、同じレポートが Telegram から /doctor 経由で実行されます。オーバーレイ、システム設定、Bluetooth、近くの WiFi、背景位置など、有効なツールが実際に依存しているすべての権限をチェックし、有効にしていないツールの権限については沈黙を保ちます。
アシスタントをモデル コンテキスト プロトコル サーバーに接続すると、AI はサーバーが公開するあらゆるツールを取得します。 AI は MCP 接続自体を追加、更新、管理できます。すべての接続変更は承認ゲート制であるため、背後でサーバーを接続することはできません。
通知 + 外部トリガー
AI に監視を許可するアプリを選択すると、AI は受信した通知を読み取り、要約し、転送できます。ホワイトリストは空から始まるため、携帯電話からは何も残されません。

あなたが選ぶのです。他のアプリ (Tasker、自動化ツール、ADB) も、外部オートメーション インテント API を介してエージェントにタスクを渡すことができるため、RikkaHub エージェントは、すでに実行されている自動化フローに組み込まれます。
厳格な順に 3 つの保護層:
アシスタントごとに切り替えます。すべてのツールは最初から始まります。必要なものだけをオンにします。
通話ごとの承認。携帯電話上で何かを変更するツールは、実行する前に尋ねます。このチャットについては 1 回だけ許可するか、常に許可するか、拒否します。
ハードラインフロア。本当に危険なコマンドの短いリスト (すべてをワイプ、再起動、フォークボム、システムファイルの破壊、ルールを回避する既知のシェルトリック) は無条件にブロックされます。誤って AI にこれらのいずれかを実行するように指示したとしても、実行されません。
さらに、パスワードと API キーがログ ファイルに記録されることはありません。 Telegram ボットは、許可リストに登録した人を除くすべてのユーザーを無視します。クラウド バックアップでは、保存されたサーバー資格情報とボット トークンがスキップされます。通知リスナーは空のホワイトリストから始まるため、転送するアプリを選択するまで携帯電話からは何も残りません。
インストール: 最新の *-release.apk を Releases からダウンロードします。不明なソースからのインストールを許可してから開きます。 (一度限りのメモ: RikkaHub Agent の古いデバッグ ビルドがまだインストールされている場合は、アンインストールしてください)

[切り捨てられた]

## Original Extract

RikkaHub Agent -- is RikkaHub fork that have Full agent mode . - ExTV/rikkahub-agent

GitHub - ExTV/rikkahub-agent: RikkaHub Agent -- is RikkaHub fork that have Full agent mode . · GitHub
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
ExTV
/
rikkahub-agent
Public
forked from rikkahub/rikkahub
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
2,956 Commits 2,956 Commits .agents/ skills .agents/ skills .github .github ai ai app app common common docs docs document document gradle gradle highlight highlight local-llm local-llm locale-tui locale-tui material3 material3 search search speech speech web-ui web-ui web web workspace workspace .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules LICENSE LICENSE README.md README.md build.gradle.kts build.gradle.kts bun.lock bun.lock gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat package.json package.json settings.gradle.kts settings.gradle.kts skills-lock.json skills-lock.json View all files Repository files navigation
A fork of RikkaHub that turns the native Android LLM chat client into a real on-device agent: 80+ device tools, AI-authored workflows, scheduled jobs, an in-app browser the AI drives, SSH, screen automation, file manager, music player, voice transcription, downloadable on-device LLMs, and a remote Telegram bot. All opt-in.
Website ·
Download ·
Features ·
Quick Start ·
Build
Vanilla LLM chat apps can answer questions. They can't open your apps, send your messages, watch your notifications, run scheduled jobs, or SSH into your server. RikkaHub Agent can. Tell it what to do in plain language, walk away, and it runs in the background, on your phone, on your terms.
"Every weekday at 9am, summarize my unread WhatsApp into one Telegram message."
"If my home server's disk fills up, ping me."
"Watch my notifications. If anything from my boss comes in, forward it to Telegram. Quietly ignore the rest."
"Find the PDF on my phone that mentions 'invoice' and read me the first paragraph."
"Take a screenshot every 30 minutes for the next 4 hours so I can see what I actually did all afternoon."
"Use Termux to build me a webpage listing everything you can do, then open it in my browser."
"When I plug in headphones at home WiFi after 7pm, start my evening playlist."
"Open my router's admin page, sign in with the saved password, and tell me which devices are eating the most bandwidth right now."
"Spin up two researches in parallel: one finds the cheapest one-way flight to Tokyo this month, the other lists hotels in Shibuya under $100. Tell me when both finish."
Each of those is a one-line setup. The phone runs them in the background while you live your life.
Ask the AI to tap, swipe, scroll, type, take screenshots, open apps, turn the torch on, change brightness or volume, post a notification, vibrate, share something, or read your battery, WiFi, signal, location, sensors, contacts, and SMS. It can also send an SMS, set the wallpaper, read and write NFC tags, sign and encrypt data with the Android Keystore, reach external storage and SD cards, and zip or unzip archives. Over 80 tools, all built into Android, no extra apps required. Each one stays off until you flip it on.
Talk to your assistant from anywhere. Set up a private Telegram bot in a minute, then chat with it like a contact. Send a question, a photo, a PDF, or a voice note. It can run on your behalf while you're at work, while you sleep, or while you're driving. Approval prompts use simple Yes/No buttons in the chat. When the AI needs to ask you something, it pops a tappable multiple-choice question (or takes a free-text reply) right in the chat. A reply too long for one message arrives as a downloadable file instead of getting truncated, and message bursts are paced so Telegram never rate-limits the answer out from under you.
The agent has a real browser built into the app. Watch it open URLs, click through cookie banners, fill in search boxes, scroll, and read the page back to you. Or send it on errands from Telegram. It streams a fresh screenshot to your chat after every step. There's a floating chat pill on the browser screen so you can keep talking to the AI without ever leaving the page. Built-in article extraction and diff-after-action keep the token cost low even on long browse sessions.
Tasker-style automation, but the AI writes the rules for you. Just describe the trigger and the action: "when I get home, turn the ringer off" ; "every weekday at 8am if battery is over 50%, check my email and ping me if anything's urgent" . 19 triggers (WiFi, Bluetooth, headphones, geofence, app launch, notifications received, time, charging, screen on/off, and more) and 14 conditions (battery thresholds, sunrise/sunset, day-of-week, current foreground app, screen state) decide when each one fires. Receivers register only when a workflow actually needs them, so battery drain stays minimal.
Set tasks to run on a schedule and forget about them. "Every Monday morning at 8", "every two hours", "next Friday at 3pm". The phone keeps everything running through reboots and battery saver. Pick how each task fires: let the AI think at the moment and decide what to do (good for "watch X and ping me if Y"), or pre-bake a fixed action that runs without using AI tokens (good for plain reminders).
The AI has its own file manager. Find files, read them, save new ones, copy, move, rename, delete. Same things you'd do in a regular file manager, except you describe what you want and it does it. "Find every PDF mentioning 'invoice' on my phone" works in one sentence. System folders that don't belong to you are off-limits, even if you ask.
Save your servers once and the AI can SSH into any of them on demand. Run a command, upload a file, pull down a backup, check disk space, tail a log. Pipe input straight into a command or write a remote file in one shot, and launch a long-running server in the background so the call returns its process ID instead of hanging on the connection. Works whether you're on WiFi or cell. Watch your home server from a coffee shop without opening a terminal.
If you have Termux installed, the AI can run real Linux commands on your phone: installing packages, building software, running scripts, or starting a background service that keeps running after the command returns. A dedicated Settings → Termux page lets you tune the command timeout, the per-turn time budget, and other Termux limits when a long install or build needs more room. On top of that, voice notes you send in Telegram get transcribed automatically. Everything runs on your phone, no cloud transcription, no API key, no internet needed.
Ask for music and the AI plays it through Android's normal media controls: lock-screen art, headphone keys, the works. Pause, resume, lower the volume for a meeting and bring it back later, all from chat or Telegram. Even after a force-stop the AI can pick up where you left off, same track, same position, via a snapshot fallback. No "you killed the player so it's gone forever". Your queue survives.
Drop a Markdown skill file into the app and the AI gains a new playbook it'll follow step-by-step: auto-reply to a contact, summarise a notification stack, or run a JavaScript mini-app whose result opens right in the in-app browser. A bundled featured catalog ships with a QR generator, a Wikipedia query box, a piano you can play, an interactive map, and more. Two skills are enabled out of the box: an always-on agent playbook that keeps the assistant proactive and self-improving across sessions, and a converter that adapts OpenClaw skills to run here. Add new skills from a URL, a markdown file you share into the app, or pick from the bundled catalog.
For long tasks the main assistant can dispatch a focused sub-agent into a clean side-context, optionally on a smaller and cheaper model. Two or more run in parallel: one researches a topic while another updates your server. Each result comes back as a single summary so the main chat doesn't drown in irrelevant tool output, and /stop cascades cancellation through every active child in one tick.
A built-in health checkup for the app. Tap Settings, then Doctor, and it runs a top-to-bottom audit of permissions, background services, database integrity, network, Termux, and diagnostics. Missing something? Tap the auto-fix button next to the row to grant the permission, restart the service, or rebuild the chat search index. The same report runs from Telegram via /doctor for remote troubleshooting. It checks every permission your enabled tools actually rely on, including overlay, system-settings, Bluetooth, nearby-WiFi, and background-location, and stays quiet about the ones for tools you haven't turned on.
Connect the assistant to Model Context Protocol servers and the AI gains whatever tools those servers expose. The AI can add, update, and manage MCP connections itself — every connection change is approval-gated, so a server can't be wired in behind your back.
Notifications + external triggers
Pick which apps the AI is allowed to watch, and it can read, summarize, and forward incoming notifications — the whitelist starts empty, so nothing leaves your phone until you choose. Other apps (Tasker, automation tools, ADB) can also hand the agent a task through the External Automation Intent API, so RikkaHub Agent slots into automation flows you already run.
Three layers of protection, in order of strictness:
Per-assistant toggles . Every tool starts off. Flip on only what you want.
Per-call approval . Tools that change something on your phone ask before running. Allow once, for this chat, always, or deny.
HARDLINE floor . A short list of genuinely dangerous commands (wipe everything, reboot, fork bombs, system file destruction, and known shell tricks to bypass the rule) is blocked unconditionally. Even if you accidentally tell the AI to do one of these, it won't.
Plus: passwords and API keys never make it into log files. The Telegram bot ignores everyone except people you put on its allowlist. Cloud backups skip your saved server credentials and bot token. The notification listener starts with an empty whitelist, so nothing leaves your phone until you pick the apps to forward.
Install : download the latest *-release.apk from Releases . Allow install from unknown sources, then open. (One-time note: if you still have an old debug build of RikkaHub Agent installed, unin

[truncated]

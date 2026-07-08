---
source: "https://github.com/TristanBrotherton/openclaw-voice-call-realtime"
hn_url: "https://news.ycombinator.com/item?id=48830588"
title: "OpenClaw plugin for real phone calls via Twilio and OpenAI Realtime"
article_title: "GitHub - TristanBrotherton/openclaw-voice-call-realtime: Give your AI assistant a phone — OpenClaw plugin for real phone calls via Twilio + OpenAI Realtime, with in-call tools, transcripts, and call screening · GitHub"
author: "thunderbong"
captured_at: "2026-07-08T12:16:23Z"
capture_tool: "hn-digest"
hn_id: 48830588
score: 1
comments: 0
posted_at: "2026-07-08T11:33:40Z"
tags:
  - hacker-news
  - translated
---

# OpenClaw plugin for real phone calls via Twilio and OpenAI Realtime

- HN: [48830588](https://news.ycombinator.com/item?id=48830588)
- Source: [github.com](https://github.com/TristanBrotherton/openclaw-voice-call-realtime)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T11:33:40Z

## Translation

タイトル: Twilio および OpenAI Realtime を介した実際の電話通話用の OpenClaw プラグイン
記事のタイトル: GitHub - TristanBrotherton/openclaw-voice-call-realtime: AI アシスタントに電話を与える — Twilio + OpenAI Realtime を介した実際の電話通話用の OpenClaw プラグイン、通話中ツール、トランスクリプト、および通話スクリーニングを備えた · GitHub
説明: AI アシスタントに電話を提供します — Twilio + OpenAI Realtime を介した実際の電話通話用の OpenClaw プラグイン、通話中ツール、トランスクリプト、および通話スクリーニングを備えています - TristanBrotherton/openclaw-voice-call-realtime

記事本文:
GitHub - TristanBrotherton/openclaw-voice-call-realtime: AI アシスタントに電話を与えます — Twilio + OpenAI Realtime を介した実際の電話通話用の OpenClaw プラグイン、通話中ツール、トランスクリプト、および通話スクリーニングを備えています · GitHub
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
トリスタンブラザートン
/

openclaw-音声通話-リアルタイム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
TristanBrotherton/openclaw-voice-call-realtime
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md Index.ts Index.ts マニフェスト.json マニフェスト.json openclaw.plugin.json openclaw.plugin.json package-lock.json package-lock.json package.json package.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
OpenClaw 音声通話 — リアルタイム版
AI アシスタントに電話を与えましょう。
この OpenClaw プラグインを使用すると、アシスタントは Twilio Programmable Voice と OpenAI Realtime API を利用して、実際の電話をかけたり受けたり、自然な全二重音声会話を行うことができます。テーブルを予約し、店の営業時間を確認し、電話メニューを操作し、人間に電話をかけ、丁寧に電話を切り、自動的に電話を切り、構造化された結果、AI が作成した概要、および完全なトランスクリプトを報告します。
あなた: 「ルイージに電話して、金曜日の 7 時に 2 人用のテーブルを予約してください。」
アシスタント: *ダイヤルし、IVR を介して会話し、ホストと交渉し、詳細を確認します*
アシスタント: 「予約済みです。金曜午後 7 時、あなたのお名前で 2 名様用のテーブルです。
彼らは、15分以上遅刻したパーティーはテーブルを失うと言った。」
なぜこれが存在するのか
すべての「AI エージェント」は電子メールを送信できます。ドライクリーニングを呼べる人はほとんどいません。レストラン、診療所、請負業者、ウェブサイトが 2019 年以来更新されていない店舗など、現実の世界は依然として電話で動いています。このプラグインはそのギャップを埋めます。アシスタントは電話番号、音声、耳、キーパッドを取得し、仕事が終わったら通話を終了する判断を取得します。
完全な会話モード — OpenAI Realt にブリッジされた Twilio Media Streams

ime API (スピーチツースピーチ、GA プロトコル)。 1 秒未満のターンアラウンド、相手が割り込んだときの自然な割り込み。
音声 AI が自律的に使用する通話ツール:
press_phone_keys — IVR メニューをナビゲートするための合成 DTMF タッチトーン (「予約するには 2 を押してください」)
report_call_outcome — 構造化された結果のキャプチャ (ステータスと収集されたすべての事実: 時間、価格、確認番号)
end_call — 正常な切断: 終了のセリフを話し、それが実際に回線上で再生される (Twilio マーク エコー) のを待ってから切断します。切り取られた別れや、長引くデッドな空気はありません。
トランスクリプトとサマリー — すべての通話は、AI によって生成されたサマリーと報告された結果を含む Markdown トランスクリプトに最終的にまとめられます。呼び出し後に get_transcript ツール アクションを介して取得できます。
通話スクリーニング認識 — 構成可能な ID フレーズで自身を識別することで、Google Call Screen とボイスメールのゲートキーパーを処理します。
目標に向けた通話 - Talking_points と call_party (ファーストパーティとサードパーティ) を渡すと、AI はタスクを続行します。つまり、ポイントをカバーし、回答を収集し、大声で確認し、まとめます。
着信通話 (オプション) — ホワイトリストでゲートされ、設定可能な挨拶付き。
デバイス プロファイル — 発信者ごとのポリシー: 応答の長さ、禁止されたアクション、追加の指示。
デフォルトで強化 — Twilio Webhook 署名検証、通話ごとのストリーム認証トークン、事前認証接続スロットル、SSRF で保護されたプロバイダー API 呼び出し、通話時間の安全キャップ、古い通話のリーピング。
プロバイダー - Twilio (推奨、完全リアルタイム会話モード)、および従来の TTS+STT パイプライン用の Telnyx / Plivo / モック。
予約と予約 — レストラン、サロン、歯医者: 「木曜日の午後の一番早い時間帯に予約してください。」
情報収集 - 営業時間、在庫確認、価格相場、「キッチンはまだ開いていますか?」
エラー

トリアージ — 補充については薬局に、見積もりについては請負業者に、駐車については会場に電話してください。
人間味のある通知 — 聞き逃す可能性のあるテキストメッセージだけでなく、メッセージを読み上げ、それが聞こえたことを確認する通話です。
あなたに連絡 — リアルタイムの決定が必要な場合、アシスタントがあなた自身の電話に電話をかけることができ、あなたはそれに話しかけるだけで済みます。
着信アシスタント回線 — ホワイトリストに登録されている発信者は、アシスタントに直接電話をかけ、アシスタントと会話できます。
電話ネットワーク あなたのマシン OpenAI
┌─────────┐ Webhook ┌───────────┐
│ Twilio │ ──────► │ プラグイン Webhook │
│ プログラマブル│ │ サーバー (HTTP) │
│ 音声 │ ◄─────── │ ・TwiML <接続> │
│ │ TwiML │ <ストリーム> │
│ │ │ │
│ メディア │ WebSocket │ メディアストリーム │ WebSocket ┌─────┐
│ ストリーム │ ◄══════════► │ ハンドラー (μ-law 8kHz) │ ◄═════════► │ リアルタイム │
━━━━━━━━━━━━━━━━━━━━━━━━━┘ 音声 ━━━━━━━━━━━┘ 音声 │ API │
━─────┘
プラグインは、OpenClaw ゲートウェイ内で HTTP サーバーを実行します。 Twilio はそこから TwiML を取得し、双方向オーディオ WebSocket を開きます。プラグインは、その音声を、STT、推論、ツール呼び出し、および TTS を 1 つのループで実行する OpenAI Realtime セッションにブリッジします。
OpenClaw ≥ 2026.6 (プラグインはパス プラグインとしてロードされ、ゲートウェイ内で実行されます)
音声対応電話番号を持つ Twilio アカウント (~$1.15/月 + ~$0.014/分)
リアルタイム API アクセスを備えた OpenAI API キー ( gpt-realtime-2.1 を推奨)
プラグインに到達するパブリック HTTPS URL

bhook ポート (Cloudflare Tunnel、ngrok、または Tailscale Funnel — 以下のウォークスルー)
mkdir -p ~ /.openclaw/plugins-src
git clone https://github.com/TristanBrotherton/openclaw-voice-call-realtime.git \
~ /.openclaw/plugins-src/voice-call-realtime
cd ~ /.openclaw/plugins-src/voice-call-realtime
npm install --omit=dev
2. Twilio のセットアップ (詳細)
twilio.com でアカウントを作成します。試用アカウントはテスト用に機能しますが、確認済みの番号に電話して試用通知を再生することしかできません。実際に使用する場合はアップグレード (請求の追加) してください。
音声対応電話番号を購入します: 「コンソール」→「電話番号」→「管理」→「番号を購入」。 [音声機能] ボックスにチェックを入れます。市内番号を選択してください。 E.164 形式でメモしてください (例: +15551234567 )。
資格情報を取得します: コンソールのホームページ → アカウント情報パネル:
アカウント SID — AC で始まります…
認証トークン — クリックして表示します。パスワードのように扱います。すべての Webhook に署名し、検証します。
地域許可 (国外に電話する場合のみ): コンソール → 音声 → 設定 → 地域許可 — 宛先の国を有効にします。
着信通話 (オプション) : コンソール → 自分の電話番号 → 音声設定 → Webhook に「着信」を設定、URL = パブリック Webhook URL (下記)、メソッド POST。アウトバウンドのみで使用する場合は、これをスキップできます。プラグインは API 経由で呼び出しごとに TwiML URL を渡します。
3. Webhook を公開する
Twilio は、HTTPS 経由でプラグインの Webhook サーバー (デフォルト ポート 3336 ) に到達する必要があります。 Cloudflare Tunnel (無料で安定したホスト名) が推奨されます。
brew install Cloudflared # または apt/pacman 同等のもの
Cloudflare トンネル ログイン # Cloudflare が管理するドメインに対して認証します
クラウドフレアトンネルの音声作成
Cloudflared トンネル ルート DNS 音声 voice.yourdomain.com
~/.cloudflared/config.yml を作成します。
トンネル：音声
認証情報ファイル: /Users/you/.cloudflared/<トンネル ID>.json
イングレス

s:
- ホスト名 : voice.yourdomain.com
サービス: http://localhost:3336
- サービス: http_status:404
実行してください: クラウドフレアトンネルの実行音声
再起動しても存続できるようにします。デッドトンネルはすべての通話を静かに遮断します。 macOS では、LaunchAgent ( ~/Library/LaunchAgents/com.you.cloudflared-voice.plist ) をインストールします。
<? XML バージョン = " 1.0 " エンコーディング = " UTF-8 " ?>
<! DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
< plist バージョン = " 1.0 " >
<辞書>
< key >ラベル</ key >< string >com.you.cloudflared-voice</ string >
< key >ProgramArguments</ key >
<配列>
< string >/opt/homebrew/bin/cloudflared</ string >
< string >トンネル</ string >< string >--config</ string >
< string >/Users/you/.cloudflared/config.yml</ string >
< string >走る</ string >< string >声</ string >
</ 配列 >
< key >RunAtLoad</ key >< true />
< key >KeepAlive</ key >< true />
< key >ThrottleInterval</ key >< integer >15</ integer >
< key >StandardOutPath</ key >< string >/tmp/cloudflared-voice.log</ string >
< key >StandardErrorPath</ key >< string >/tmp/cloudflared-voice.log</ string >
</辞書>
</plist>
launchctl ブートストラップ gui/ $( id -u ) ~ /Library/LaunchAgents/com.you.cloudflared-voice.plist
Linux では、 Restart=always を指定して systemd ユニットを使用します。代替案: ngrok http 3336 (自動管理用に、tunnel.provider: "ngrok" を設定) または tailscale funnel 3336 。
~/.openclaw/openclaw.json 内:
{
「プラグイン」: {
"allow" : [ " voice-call-tristan " ], // 他のプラグインも加えて
"load" : { "paths" : [ " /Users/you/.openclaw/plugins-src " ] },
「エントリ」: {
"音声通話トリスタン" : {
"有効" : true 、
"構成" : {
"有効" : true 、
"プロバイダー" : " twilio " 、
"fromNumber" : " +15551234567 " , // Twilio 番号
"toNumber" : " +15557654321 " , // デフォルトの宛先 (あなた)
"twilio" : {
「はい、

ountSid" : " ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx " ,
"authToken" : " your_twilio_auth_token "
}、
"serve" : { "port" : 3336 、 "bind" : " 127.0.0.1 " 、 "path" : " /voice/webhook " }、
"publicUrl" : " https://voice.yourdomain.com/voice/webhook " ,
"アウトバウンド" : { "デフォルトモード" : " 会話 " },
"maxDurationSeconds" : 3600 、
「ストリーミング」: {
"有効" : true 、
"sttProvider" : " openai-realtime-conversation " ,
"realtimeModel" : " gpt-realtime-2.1 " 、
"realtimeVoice" : " 合金 " 、
"openaiApiKey" : " sk-... " ,
"realtimePolicy" : { "idleTimeoutMs" : 120000 、 "maxSessionMs" : 7200000 }
}、
"callScreening" : {
"有効" : true 、
"callerIdentity" : " こんにちは、これはアレックスの AI アシスタントが彼らに代わって電話をかけています。 」
}
}
}
}
}
}
検証して再起動します。
openclaw 設定の検証
openclaw ゲートウェイの再起動
ゲートウェイのログに次の内容が表示されるはずです。
[音声通話] メディア ストリーミングが初期化されました
[音声通話] http://127.0.0.1:3336/voice/webhook をリッスンする Webhook サーバー
[音声通話] ランタイムが初期化されました
パブリック URL をスモークテストします。ダイヤルする前に両方とも成功する必要があります。
curl -s -o /dev/null -w " %{http_code}\n " -X POST https://voice.yourdomain.com/voice/webhook # → 200
curl -s -o /dev/null -w " %{http_code}\n " -H " 接続: アップグレード " -H " アップグレード: websocket " \
-H " Sec-WebSocket-Version: 13 " -H " Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ== " \
https://voice.yourdomain.com/voice/stream # → 101
使用法
OpenClaw エージェントは、次のアクションで voice_call ツールを取得します。
// 目標に向けた呼び出しを行う
{ "アクション" : "initiate_call " ,
"宛先" : " +15558675309 " 、
「メッセージ」：「こんにちは！アレックスに代わってテーブルを予約するために電話しています。 「、
"モード" : "会話" ,
"call_party" : " サードパート

[切り捨てられた]

## Original Extract

Give your AI assistant a phone — OpenClaw plugin for real phone calls via Twilio + OpenAI Realtime, with in-call tools, transcripts, and call screening - TristanBrotherton/openclaw-voice-call-realtime

GitHub - TristanBrotherton/openclaw-voice-call-realtime: Give your AI assistant a phone — OpenClaw plugin for real phone calls via Twilio + OpenAI Realtime, with in-call tools, transcripts, and call screening · GitHub
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
TristanBrotherton
/
openclaw-voice-call-realtime
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
TristanBrotherton/openclaw-voice-call-realtime
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md index.ts index.ts manifest.json manifest.json openclaw.plugin.json openclaw.plugin.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
OpenClaw Voice Call — Realtime Edition
Give your AI assistant a phone.
This OpenClaw plugin lets your assistant place and receive real phone calls and hold natural, full-duplex voice conversations — powered by Twilio Programmable Voice and the OpenAI Realtime API. It books your table, checks the store's hours, navigates the phone menu, gets a human on the line, wraps up politely, hangs up on its own, and reports back with a structured outcome, an AI-written summary, and the full transcript.
You: "Call Luigi's and book a table for two on Friday at 7."
Assistant: *dials, talks past the IVR, negotiates with the host, confirms details*
Assistant: "Booked — table for 2, Friday 7pm, under your name.
They said parties over 15 minutes late lose the table."
Why this exists
Every "AI agent" can send an email. Almost none of them can call the dry cleaner. The real world still runs on phone calls — restaurants, doctors' offices, contractors, that one store whose website hasn't been updated since 2019. This plugin closes that gap: your assistant gets a phone number, a voice, ears, a keypad, and the judgment to end the call when the job is done.
Full conversation mode — Twilio Media Streams bridged to the OpenAI Realtime API (speech-to-speech, GA protocol). Sub-second turnaround, natural barge-in when the other party interrupts.
In-call tools the voice AI uses autonomously:
press_phone_keys — synthesized DTMF touch-tones for navigating IVR menus ("press 2 for reservations")
report_call_outcome — structured result capture (status + every fact gathered: times, prices, confirmation numbers)
end_call — graceful hangup: speaks a closing line, waits for it to actually play out on the line (Twilio mark echo), then disconnects. No clipped goodbyes, no lingering dead air.
Transcripts & summaries — every call is finalized into a Markdown transcript with an AI-generated summary and the reported outcome. Retrievable after the call via the get_transcript tool action.
Call screening awareness — handles Google Call Screen and voicemail gatekeepers by identifying itself with a configurable identity phrase.
Goal-directed calls — pass talking_points and call_party (first-party vs third-party) and the AI stays on task: cover the points, collect the answers, confirm out loud, wrap up.
Inbound calls (optional) — allowlist-gated, with a configurable greeting.
Device profiles — per-caller policies: response length, forbidden actions, extra instructions.
Hardened by default — Twilio webhook signature verification, per-call stream auth tokens, pre-auth connection throttling, SSRF-guarded provider API calls, call-duration safety caps, stale-call reaping.
Providers — Twilio (recommended, full realtime conversation mode), plus Telnyx / Plivo / mock for the legacy TTS+STT pipeline.
Reservations & appointments — restaurants, salons, dentists: "book me the earliest slot Thursday afternoon."
Information gathering — opening hours, stock checks, price quotes, "is the kitchen still open?"
Errand triage — call the pharmacy about a refill, the contractor about a quote, the venue about parking.
Notifications with a human touch — a call that speaks a message and confirms it was heard, not just a text that might be missed.
Reaching you — your assistant can call your own phone when something needs a real-time decision, and you can just talk to it.
Inbound assistant line — allowlisted callers can ring your assistant directly and converse with it.
Phone network Your machine OpenAI
┌─────────────┐ webhooks ┌──────────────────────┐
│ Twilio │ ───────────► │ Plugin webhook │
│ Programmable│ │ server (HTTP) │
│ Voice │ ◄─────────── │ · TwiML <Connect> │
│ │ TwiML │ <Stream> │
│ │ │ │
│ Media │ WebSocket │ Media stream │ WebSocket ┌──────────┐
│ Streams │ ◄══════════► │ handler (μ-law 8kHz) │ ◄═════════► │ Realtime │
└─────────────┘ audio └──────────────────────┘ audio │ API │
└──────────┘
The plugin runs an HTTP server inside the OpenClaw gateway. Twilio fetches TwiML from it and opens a bidirectional audio WebSocket; the plugin bridges that audio to an OpenAI Realtime session which does STT, reasoning, tool calls, and TTS in one loop.
OpenClaw ≥ 2026.6 (the plugin loads as a path plugin and runs inside the gateway)
Twilio account with a voice-capable phone number (~$1.15/mo + ~$0.014/min)
OpenAI API key with Realtime API access ( gpt-realtime-2.1 recommended)
A public HTTPS URL reaching the plugin's webhook port (Cloudflare Tunnel, ngrok, or Tailscale Funnel — walkthrough below)
mkdir -p ~ /.openclaw/plugins-src
git clone https://github.com/TristanBrotherton/openclaw-voice-call-realtime.git \
~ /.openclaw/plugins-src/voice-call-realtime
cd ~ /.openclaw/plugins-src/voice-call-realtime
npm install --omit=dev
2. Set up Twilio (detailed)
Create an account at twilio.com . Trial accounts work for testing but can only call verified numbers and play a trial notice — upgrade (add billing) for real use.
Buy a voice-capable phone number : Console → Phone Numbers → Manage → Buy a number . Check the Voice capability box. Pick a local number; note it in E.164 form (e.g. +15551234567 ).
Get your credentials : Console home page → Account Info panel:
Account SID — starts with AC…
Auth Token — click to reveal. Treat it like a password; it signs and verifies every webhook.
Geo permissions (only if calling outside your country): Console → Voice → Settings → Geo permissions — enable the destination countries.
Inbound calls (optional) : Console → your phone number → Voice Configuration → set "A call comes in" to Webhook , URL = your public webhook URL (below), method POST. For outbound-only use you can skip this — the plugin passes TwiML URLs per call via the API.
3. Expose the webhook publicly
Twilio must reach the plugin's webhook server (default port 3336 ) over HTTPS. Cloudflare Tunnel (free, stable hostname) is recommended:
brew install cloudflared # or apt/pacman equivalent
cloudflared tunnel login # authorize against your Cloudflare-managed domain
cloudflared tunnel create voice
cloudflared tunnel route dns voice voice.yourdomain.com
Create ~/.cloudflared/config.yml :
tunnel : voice
credentials-file : /Users/you/.cloudflared/<tunnel-id>.json
ingress :
- hostname : voice.yourdomain.com
service : http://localhost:3336
- service : http_status:404
Run it: cloudflared tunnel run voice
Make it survive reboots. A dead tunnel silently breaks all calls. On macOS, install a LaunchAgent ( ~/Library/LaunchAgents/com.you.cloudflared-voice.plist ):
<? xml version = " 1.0 " encoding = " UTF-8 " ?>
<! DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
< plist version = " 1.0 " >
< dict >
< key >Label</ key >< string >com.you.cloudflared-voice</ string >
< key >ProgramArguments</ key >
< array >
< string >/opt/homebrew/bin/cloudflared</ string >
< string >tunnel</ string >< string >--config</ string >
< string >/Users/you/.cloudflared/config.yml</ string >
< string >run</ string >< string >voice</ string >
</ array >
< key >RunAtLoad</ key >< true />
< key >KeepAlive</ key >< true />
< key >ThrottleInterval</ key >< integer >15</ integer >
< key >StandardOutPath</ key >< string >/tmp/cloudflared-voice.log</ string >
< key >StandardErrorPath</ key >< string >/tmp/cloudflared-voice.log</ string >
</ dict >
</ plist >
launchctl bootstrap gui/ $( id -u ) ~ /Library/LaunchAgents/com.you.cloudflared-voice.plist
On Linux, use a systemd unit with Restart=always . Alternatives: ngrok http 3336 (set tunnel.provider: "ngrok" for auto-management) or tailscale funnel 3336 .
In ~/.openclaw/openclaw.json :
{
"plugins" : {
"allow" : [ " voice-call-tristan " ], // plus your other plugins
"load" : { "paths" : [ " /Users/you/.openclaw/plugins-src " ] },
"entries" : {
"voice-call-tristan" : {
"enabled" : true ,
"config" : {
"enabled" : true ,
"provider" : " twilio " ,
"fromNumber" : " +15551234567 " , // your Twilio number
"toNumber" : " +15557654321 " , // default destination (you)
"twilio" : {
"accountSid" : " ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx " ,
"authToken" : " your_twilio_auth_token "
},
"serve" : { "port" : 3336 , "bind" : " 127.0.0.1 " , "path" : " /voice/webhook " },
"publicUrl" : " https://voice.yourdomain.com/voice/webhook " ,
"outbound" : { "defaultMode" : " conversation " },
"maxDurationSeconds" : 3600 ,
"streaming" : {
"enabled" : true ,
"sttProvider" : " openai-realtime-conversation " ,
"realtimeModel" : " gpt-realtime-2.1 " ,
"realtimeVoice" : " alloy " ,
"openaiApiKey" : " sk-... " ,
"realtimePolicy" : { "idleTimeoutMs" : 120000 , "maxSessionMs" : 7200000 }
},
"callScreening" : {
"enabled" : true ,
"callerIdentity" : " Hi, this is Alex's AI assistant calling on their behalf. "
}
}
}
}
}
}
Validate and restart:
openclaw config validate
openclaw gateway restart
You should see in the gateway logs:
[voice-call] Media streaming initialized
[voice-call] Webhook server listening on http://127.0.0.1:3336/voice/webhook
[voice-call] Runtime initialized
Smoke-test the public URL — both must succeed before dialing:
curl -s -o /dev/null -w " %{http_code}\n " -X POST https://voice.yourdomain.com/voice/webhook # → 200
curl -s -o /dev/null -w " %{http_code}\n " -H " Connection: Upgrade " -H " Upgrade: websocket " \
-H " Sec-WebSocket-Version: 13 " -H " Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ== " \
https://voice.yourdomain.com/voice/stream # → 101
Usage
Your OpenClaw agent gets a voice_call tool with these actions:
// Place a goal-directed call
{ "action" : " initiate_call " ,
"to" : " +15558675309 " ,
"message" : " Hi! I'm calling on behalf of Alex to book a table. " ,
"mode" : " conversation " ,
"call_party" : " third-part

[truncated]

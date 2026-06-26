---
source: "https://github.com/team-telnyx/telnyx-code-examples/tree/main/outbound-hold-agent-python"
hn_url: "https://news.ycombinator.com/item?id=48688082"
title: "Outbound hold agent that pauses AI runtime while waiting on hold"
article_title: "telnyx-code-examples/outbound-hold-agent-python at main · team-telnyx/telnyx-code-examples · GitHub"
author: "anushathukral"
captured_at: "2026-06-26T16:38:25Z"
capture_tool: "hn-digest"
hn_id: 48688082
score: 1
comments: 0
posted_at: "2026-06-26T15:51:26Z"
tags:
  - hacker-news
  - translated
---

# Outbound hold agent that pauses AI runtime while waiting on hold

- HN: [48688082](https://news.ycombinator.com/item?id=48688082)
- Source: [github.com](https://github.com/team-telnyx/telnyx-code-examples/tree/main/outbound-hold-agent-python)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T15:51:26Z

## Translation

タイトル: 保留待機中に AI ランタイムを一時停止する送信保留エージェント
記事のタイトル: メインの telnyx-code-examples/outbound-hold-agent-python · チーム-telnyx/telnyx-code-examples · GitHub
説明: Telnyx AI 通信インフラストラクチャの運用対応コード例 — 音声 AI、SMS、SIP、IoT API - telnyx-code-examples/outbound-hold-agent-python at main · Team-telnyx/telnyx-code-examples

記事本文:
メインの telnyx-code-examples/outbound-hold-agent-python · チーム-telnyx/telnyx-code-examples · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
チームテルニクス
/
telnyx コードの例
公共
通知
あなたは署名しているに違いありません

通知設定を変更するには編集してください
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py要件.txt 要件.txt すべてのファイルを表示 README.md
概要
名前
アウトバウンド保留エージェント
タイトル
アウトバウンド ホールドアウェア AI エージェント
説明
Telnyx AI アシスタントで企業に電話し、IVR を操作し、保留中にアシスタントを一時停止し、文字起こしで監視し、担当者が応答したらコンテキストをもとに再開します。
言語
パイソン
フレームワーク
ファスタピ
テルニクス製品
声
AIアシスタント
転写
TeXML
チャンネル
声
アウトバウンド ホールドアウェア AI エージェント
企業への電話、IVR の操作、保留中のアクティブな AI アシスタントの停止、文字起こしによる通話の監視、元の目的と承認されたコンテキストで代表者向けアシスタントを再起動できるアウトバウンド Telnyx AI 音声エージェントを構築します。
これは、保険会社、ホテル、診療所、サービスプロバイダー、またはエージェントがメニューを見て数分を費やし、人間が応答するまで待ち行列を作る可能性のある企業に電話するエージェントに役立ちます。
アウトバウンド通話制御コールを発信します。
応答後に IVR ナビゲーション AI アシスタントを開始します。
アシスタントが /tools/send-dtmf を通じてバックエンド所有の DTMF をリクエストできるようにします。
Telnyx イベント、アシスタント ツールの呼び出し、またはトランスクリプト フレーズから保留を検出します。
ai_assistant_stop を使用して、保留中にアクティブなアシスタントを停止します。
保留中に文字起こしのみの監視を開始します。
call.unhold またはトランスクリプト フレーズから代表的なピックアップを検出します。
元のタスク、コンテキスト、保留期間、最近のトランスクリプトを使用して 2 番目の AI アシスタントを開始します。
タスクを完了するための /tools/end-call ツールを公開します。
決定論的な偽会社の TeXML フローが含まれています。

再現可能なテスト。
ダイヤル: POST /v2/calls - API リファレンス
AI アシスタントを開始する : POST /v2/calls/{call_control_id}/actions/ai_assistant_start - API リファレンス
AI アシスタントを停止する : POST /v2/calls/{call_control_id}/actions/ai_assistant_stop - API リファレンス
DTMF の送信 : POST /v2/calls/{call_control_id}/actions/send_dtmf - API リファレンス
文字起こし開始 : POST /v2/calls/{call_control_id}/actions/transcription_start - API リファレンス
文字起こしの停止 : POST /v2/calls/{call_control_id}/actions/transcription_stop - API リファレンス
ハングアップ : POST /v2/calls/{call_control_id}/actions/hangup - API リファレンス
call.answered - IVR アシスタントを開始します。
call.hold - アシスタントを停止し、保留モニタリングを開始します。
call.unhold - 通話を代表者対応可能として扱います。
call.transcription - 保留フレーズと代表的なピックアップ フレーズを検出します。
call.hangup - ローカルセッションが終了したことをマークします。
クライアント/ワークフロー
-> POST /calls/outbound
-> Telnyx が対象会社にダイヤル
-> 電話をかけました。応答しました
-> IVR AIアシスタントが起動します
-> アシスタントはメニューのために /tools/send-dtmf を呼び出します
-> ホールド検出
-> バックエンドがアシスタントを停止する
-> 転写のみのホールド監視
-> 代表者が検出されました
-> 代表的な AI アシスタントはコンテキストから開始
-> タスクが完了しました
-> アシスタントが /tools/end-call を呼び出します
環境変数
変数
実際の通話に必要
説明
TELNYX_API_KEY
はい
音声 API リクエストに使用される Telnyx API キー。
TELNYX_CONNECTION_ID
はい
音声 API / 通話制御接続 ID。
TELNYX_FROM_NUMBER
はい
E.164 形式の Telnyx 発信者 ID 番号。
TELNYX_IVR_ASSISTANT_ID
はい
長押しする前にメニュー ナビゲーションに使用されるアシスタント。
TELNYX_REPRESENTATIVE_ASSISTANT_ID
はい
代表ピックアップ後にアシスタントを使用。
PUBLIC_BASE_URL
はい
Webhook、アシスタント ツール、および偽の会社 TeXML 用のパブリック HTTPS ベース URL。
TELNYX_PUBLIC_KEY
推奨
Telnyx Webhook P

署名検証用のパブリックキー。
転写エンジン
いいえ
デフォルトはディープグラムです。
転写_モデル
いいえ
デフォルトは nova-2 です。
TRANSCRIPTION_LANGUAGE
いいえ
デフォルトは en です。
START_TRANSCRIPTION_DURING_IVR
いいえ
デフォルトは true なので、デモ中にフレーズ検出で保留言語をキャッチできます。
TELNYX_DRY_RUN
いいえ
実際の Telnyx API 呼び出しを使用しないローカル テストの場合、デフォルトは true です。
港
いいえ
ローカルサーバーポート。デフォルトは 8000 です。
セットアップ
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/outbound-hold-agent-python
cp .env.example .env
python3 -m venv .venv
ソース .venv/bin/activate
python -m pip install -rrequirements.txt
Python app.py
このターミナルで Python app.py を実行し続けます。同じフォルダー内の 2 番目のターミナルを開いて、curl コマンドを実行します。
ドライラン モードはデフォルトで有効になっているため、ローカル セッションと模擬 Telnyx コマンド応答が作成されます。
カール -X POST http://127.0.0.1:8000/calls/outbound \
-H " Content-Type: application/json " \
-d ' {
"へ":"+15551234567",
"objective":"金曜日の夜のホテルの予約をする",
"target_company":"ウィロー クリーク ホテル",
"context":{"guest_name":"Alex Morgan","party_size":2}
} '
返された call_control_id を使用して、電話に応答する Telnyx をシミュレートします。
カール -X POST http://127.0.0.1:8000/webhooks/telnyx \
-H " Content-Type: application/json " \
-d ' {"data":{"event_type":"call.answered","payload":{"call_control_id":"dry-run-call-id"}}} '
IVR メニュー ナビゲーションをシミュレートします。
カール -X POST http://127.0.0.1:8000/tools/send-dtmf \
-H " Content-Type: application/json " \
-d ' {"桁":"1","理由":"予約メニュー オプション"} '
ホールド検出をシミュレートします。
カール -X POST http://127.0.0.1:8000/tools/hold-detected \
-H " Content-Type: application/json " \
-d ' {"理由":"次に空いている代表者のために保留してください","信頼":0.95}

'
文字起こしからの代表的なピックアップをシミュレートします。
カール -X POST http://127.0.0.1:8000/webhooks/telnyx \
-H " Content-Type: application/json " \
-d ' {"data":{"event_type":"call.transcription","payload":{"call_control_id":"dry-run-call-id","transcript":"お待ちいただきありがとうございます。予約中のサラです。"}}} '
状態を検査します:
カール http://127.0.0.1:8000/sessions
組み込みの偽会社でテストする
ngrok http 8000
PUBLIC_BASE_URL を HTTPS ngrok URL に設定します。次に、Telnyx TeXML アプリケーションまたはテスト番号を次のように指定します。
https://YOUR_PUBLIC_BASE_URL/fake-company/texml
偽の会社は Willow Creek Hotel として応答し、メニューを提示し、数字 1 を受け入れ、保留言語を発し、次に代表的なピックアップ フレーズを発します。実際の会社に電話する前にこれを使用してください。
次のツールを使用して IVR アシスタントを構成します。
POST https://YOUR_PUBLIC_BASE_URL/tools/send-dtmf
POST https://YOUR_PUBLIC_BASE_URL/tools/hold-detected
このオプションのツールを使用して代表アシスタントを構成します。
POST https://YOUR_PUBLIC_BASE_URL/tools/end-call
APIリファレンス
この例で公開されるローカル エンドポイントについては、API.md を参照してください。
アプリをパブリック HTTPS 経由で公開し、 PUBLIC_BASE_URL を設定します。
Telnyx Voice API アプリケーションの Webhook URL を https://YOUR_PUBLIC_BASE_URL/webhooks/telnyx に構成します。
Webhook 署名が検証されるように TELNYX_PUBLIC_KEY を設定します。
ローカル ワークフロー エンドポイントとアシスタント ツール エンドポイントに認証を追加します。
DTMF アクションをバックエンドで所有し、ターゲット企業ごとに許可された数字を検証します。
メモリ内のセッションを永続ストレージに置き換えます。
宛先のホワイトリスト、レート制限、再試行、スタックコールアラートを追加します。
アウトバウンド通話、AI 開示、録音、転写、保存の要件を確認します。
問題
原因
修正
Curl は機能しますが、実際の呼び出しは行われません
TELNYX_DRY_RUN=true
conf の後に TELNYX_DRY_RUN=false を設定します

Telnyx 値を設定します。
アシスタントが起動しない
アシスタント ID がないか、通話に応答しない Webhook がありません
アシスタント ID と Webhook URL を確認します。
DTMF ツールは受け入れられますが、IVR が動きません
メニューが準備できる前に間違った数字または DTMF が送信されました
アシスタントがプロンプトを待って有効なオプションを送信することを確認します。
代表アシスタントが始まらない
ピックアップフレーズが検出されませんでした
REPRESENTATIVE_PHRASES を調整するか、 call.unhold をトリガーします。
Webhook が 401 を返す
署名検証に失敗しました
TELNYX_PUBLIC_KEY と Webhook 署名ヘッダーを確認します。
リソース
Telnyx は、発信通話、通話制御 Webhook、DTMF、AI アシスタント、リアルタイム文字起こしを 1 つの音声 API で公開する AI コミュニケーション インフラストラクチャ プラットフォームです。そのため、アプリは個別のプロバイダーをつなぎ合わせることなく、通話ライフサイクル全体を制御できます。
make-outbound-phone-call-python - 発信通話制御通話を発信します。
build-ivr-phone-menu-python - 従来の DTMF IVR を構築します。
ai-voice-agent-with-function-calling-python - AI 音声エージェントにツール呼び出しを追加します。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-ready code examples for Telnyx AI Communications Infrastructure — Voice AI, SMS, SIP, and IoT APIs - telnyx-code-examples/outbound-hold-agent-python at main · team-telnyx/telnyx-code-examples

telnyx-code-examples/outbound-hold-agent-python at main · team-telnyx/telnyx-code-examples · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
team-telnyx
/
telnyx-code-examples
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py requirements.txt requirements.txt View all files README.md
Outline
name
outbound-hold-agent
title
Outbound Hold-Aware AI Agent
description
Call a business, navigate IVRs with a Telnyx AI Assistant, pause the assistant during hold, monitor with transcription, and resume with context when a representative answers.
language
python
framework
fastapi
telnyx_products
Voice
AI Assistants
Transcription
TeXML
channel
voice
Outbound Hold-Aware AI Agent
Build an outbound Telnyx AI voice agent that can call a business, navigate an IVR, stop the active AI Assistant during hold, monitor the call with transcription, and restart a representative-facing assistant with the original objective and approved context.
This is useful for agents that call insurance companies, hotels, clinics, service providers, or any business where the agent may spend several minutes in menus and hold queues before a human answers.
Places an outbound Call Control call.
Starts an IVR navigation AI Assistant after answer.
Lets the assistant request backend-owned DTMF through /tools/send-dtmf .
Detects hold from Telnyx events, assistant tool calls, or transcript phrases.
Stops the active assistant during hold with ai_assistant_stop .
Starts transcription-only monitoring during hold.
Detects representative pickup from call.unhold or transcript phrases.
Starts a second AI Assistant with the original task, context, hold duration, and recent transcript.
Exposes an /tools/end-call tool for task completion.
Includes a deterministic fake company TeXML flow for repeatable testing.
Dial : POST /v2/calls - API reference
Start AI Assistant : POST /v2/calls/{call_control_id}/actions/ai_assistant_start - API reference
Stop AI Assistant : POST /v2/calls/{call_control_id}/actions/ai_assistant_stop - API reference
Send DTMF : POST /v2/calls/{call_control_id}/actions/send_dtmf - API reference
Transcription Start : POST /v2/calls/{call_control_id}/actions/transcription_start - API reference
Transcription Stop : POST /v2/calls/{call_control_id}/actions/transcription_stop - API reference
Hangup : POST /v2/calls/{call_control_id}/actions/hangup - API reference
call.answered - start the IVR assistant.
call.hold - stop the assistant and enter hold monitoring.
call.unhold - treat the call as representative-ready.
call.transcription - detect hold and representative pickup phrases.
call.hangup - mark the local session ended.
Client / workflow
-> POST /calls/outbound
-> Telnyx dials target company
-> call.answered
-> IVR AI Assistant starts
-> assistant calls /tools/send-dtmf for menus
-> hold detected
-> backend stops assistant
-> transcription-only hold monitoring
-> representative detected
-> representative AI Assistant starts with context
-> task completes
-> assistant calls /tools/end-call
Environment Variables
Variable
Required for real calls
Description
TELNYX_API_KEY
yes
Telnyx API key used for Voice API requests.
TELNYX_CONNECTION_ID
yes
Voice API / Call Control connection ID.
TELNYX_FROM_NUMBER
yes
Telnyx caller ID number in E.164 format.
TELNYX_IVR_ASSISTANT_ID
yes
Assistant used for menu navigation before hold.
TELNYX_REPRESENTATIVE_ASSISTANT_ID
yes
Assistant used after representative pickup.
PUBLIC_BASE_URL
yes
Public HTTPS base URL for webhooks, assistant tools, and fake company TeXML.
TELNYX_PUBLIC_KEY
recommended
Telnyx webhook public key for signature verification.
TRANSCRIPTION_ENGINE
no
Defaults to Deepgram .
TRANSCRIPTION_MODEL
no
Defaults to nova-2 .
TRANSCRIPTION_LANGUAGE
no
Defaults to en .
START_TRANSCRIPTION_DURING_IVR
no
Defaults to true so phrase detection can catch hold language during demos.
TELNYX_DRY_RUN
no
Defaults to true for local testing without real Telnyx API calls.
PORT
no
Local server port. Defaults to 8000 .
Setup
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/outbound-hold-agent-python
cp .env.example .env
python3 -m venv .venv
source .venv/bin/activate
python -m pip install -r requirements.txt
python app.py
Keep python app.py running in this terminal. Open a second terminal in the same folder to run the curl commands.
Dry-run mode is enabled by default, so this creates a local session and mock Telnyx command responses:
curl -X POST http://127.0.0.1:8000/calls/outbound \
-H " Content-Type: application/json " \
-d ' {
"to":"+15551234567",
"objective":"book a hotel reservation for Friday night",
"target_company":"Willow Creek Hotel",
"context":{"guest_name":"Alex Morgan","party_size":2}
} '
Simulate Telnyx answering the call by using the returned call_control_id :
curl -X POST http://127.0.0.1:8000/webhooks/telnyx \
-H " Content-Type: application/json " \
-d ' {"data":{"event_type":"call.answered","payload":{"call_control_id":"dry-run-call-id"}}} '
Simulate IVR menu navigation:
curl -X POST http://127.0.0.1:8000/tools/send-dtmf \
-H " Content-Type: application/json " \
-d ' {"digits":"1","reason":"reservations menu option"} '
Simulate hold detection:
curl -X POST http://127.0.0.1:8000/tools/hold-detected \
-H " Content-Type: application/json " \
-d ' {"reason":"please hold for the next available representative","confidence":0.95} '
Simulate representative pickup from transcription:
curl -X POST http://127.0.0.1:8000/webhooks/telnyx \
-H " Content-Type: application/json " \
-d ' {"data":{"event_type":"call.transcription","payload":{"call_control_id":"dry-run-call-id","transcript":"thanks for holding, this is Sarah with reservations"}}} '
Inspect state:
curl http://127.0.0.1:8000/sessions
Test With The Built-In Fake Company
ngrok http 8000
Set PUBLIC_BASE_URL to the HTTPS ngrok URL. Then point a Telnyx TeXML application or test number at:
https://YOUR_PUBLIC_BASE_URL/fake-company/texml
The fake company answers as Willow Creek Hotel, presents a menu, accepts digit 1 , emits hold language, and then emits a representative pickup phrase. Use this before calling a real company.
Configure the IVR assistant with these tools:
POST https://YOUR_PUBLIC_BASE_URL/tools/send-dtmf
POST https://YOUR_PUBLIC_BASE_URL/tools/hold-detected
Configure the representative assistant with this optional tool:
POST https://YOUR_PUBLIC_BASE_URL/tools/end-call
API Reference
See API.md for local endpoints exposed by this example.
Expose the app over public HTTPS and set PUBLIC_BASE_URL .
Configure your Telnyx Voice API application webhook URL to https://YOUR_PUBLIC_BASE_URL/webhooks/telnyx .
Set TELNYX_PUBLIC_KEY so webhook signatures are verified.
Add authentication to local workflow endpoints and assistant tool endpoints.
Keep DTMF actions backend-owned and validate allowed digits per target company.
Replace in-memory sessions with persistent storage.
Add destination allowlists, rate limits, retries, and stuck-call alerting.
Review outbound calling, AI disclosure, recording, transcription, and retention requirements.
Issue
Cause
Fix
Curl works but no real call is placed
TELNYX_DRY_RUN=true
Set TELNYX_DRY_RUN=false after configuring Telnyx values.
Assistant does not start
Missing assistant ID or no call.answered webhook
Check assistant IDs and webhook URL.
DTMF tool is accepted but IVR does not move
Wrong digit or DTMF sent before menu is ready
Confirm the assistant waits for the prompt and sends a valid option.
Representative assistant never starts
No pickup phrase was detected
Tune REPRESENTATIVE_PHRASES or trigger call.unhold .
Webhook returns 401
Signature verification failed
Confirm TELNYX_PUBLIC_KEY and webhook signature headers.
Resources
Telnyx is an AI Communications Infrastructure platform that exposes outbound calling, call-control webhooks, DTMF, AI Assistants, and real-time transcription in one Voice API, so the app can control the full call lifecycle without stitching together separate providers.
make-outbound-phone-call-python - place an outbound Call Control call.
build-ivr-phone-menu-python - build a traditional DTMF IVR.
ai-voice-agent-with-function-calling-python - add tool calls to an AI voice agent.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

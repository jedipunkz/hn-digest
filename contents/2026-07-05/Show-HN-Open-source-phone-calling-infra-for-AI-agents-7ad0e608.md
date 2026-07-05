---
source: "https://github.com/AgentLineHQ/AgentLine"
hn_url: "https://news.ycombinator.com/item?id=48793076"
title: "Show HN: Open-source phone calling infra for AI agents"
article_title: "GitHub - AgentLineHQ/AgentLine: Give your AI agent a real phone number. Open-source telephony API with voice calls, SMS, MCP server, and skill file. Built with FastAPI + SignalWire + Deepgram. · GitHub"
author: "sameersri2004"
captured_at: "2026-07-05T11:24:01Z"
capture_tool: "hn-digest"
hn_id: 48793076
score: 3
comments: 0
posted_at: "2026-07-05T10:51:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source phone calling infra for AI agents

- HN: [48793076](https://news.ycombinator.com/item?id=48793076)
- Source: [github.com](https://github.com/AgentLineHQ/AgentLine)
- Score: 3
- Comments: 0
- Posted: 2026-07-05T10:51:56Z

## Translation

タイトル: Show HN: AI エージェント向けのオープンソースの通話インフラ
記事のタイトル: GitHub - AgentLineHQ/AgentLine: AI エージェントに実際の電話番号を与えます。音声通話、SMS、MCP サーバー、スキル ファイルを備えたオープンソースのテレフォニー API。 FastAPI + SignalWire + Deepgram で構築されています。 · GitHub
説明: AI エージェントに実際の電話番号を与えます。音声通話、SMS、MCP サーバー、スキル ファイルを備えたオープンソースのテレフォニー API。 FastAPI + SignalWire + Deepgram で構築されています。 - エージェントライン本社/エージェントライン

記事本文:
GitHub - AgentLineHQ/AgentLine: AI エージェントに実際の電話番号を与えます。音声通話、SMS、MCP サーバー、スキル ファイルを備えたオープンソースのテレフォニー API。 FastAPI + SignalWire + Deepgram で構築されています。 · GitHub
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
読み込み中にエラーが発生しました

g.このページをリロードしてください。
エージェントライン本社
/
エージェントライン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット .github .github エージェントライン エージェントライン移行 移行スキル/ エージェントライン スキル/ エージェントライン .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md エージェントラインロゴ-200.png エージェントラインロゴ-200.png エージェントラインロゴ.png エージェントラインロゴ.png claude_desktop_config_example.json claude_desktop_config_example.json docker-compose.yml docker-compose.yml Railway.toml Railway.toml要件.txt 要件.txt schema.sql schema.sql すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用のオープンソース テレフォニー API — 電話番号、音声通話、SMS
AI エージェントに実際の電話番号を与えます。発信通話の発信、受信通話の受信、SMS の処理を​​すべて 1 つの API を通じて行います。通信に関する専門知識は必要ありません。
ウェブサイト · ドキュメント · スキルファイル · Discord
AgentLine は、AI エージェントに実際の電話番号と人間のような音声を提供する、オープンソースの AI ネイティブ テレフォニー プラットフォームです。シンプルな REST API と MCP サーバーを提供するため、AI エージェント (クロード、カーソル、カスタム LLM エージェント) は、電気通信の知識がなくても、電話の発着信、SMS の処理、トランスクリプトの取得を行うことができます。
AI エージェント → AgentLine API → 実際の電話通話
→SMSメッセージ
→ 通話記録
なぜエージェントラインなのか？
エージェントライン
Twilio/Vonage
自分で構築する
AI エージェント向けに構築
✅ 専用設計
❌ コールセンター向けに構築
❌ あなたはそれを縫い合わせます
セットアップ時間
5分
営業時間
週間
MCPサーバー
✅ ネイティブ
❌ なし
❌ 自分で構築する
スキルファイル

ストール
✅ 1つのファイル
❌ 複雑な SDK
❌ 数百行
音声パイプライン
✅ 含まれる (STT + TTS + LLM)
❌ 期限付き
❌ 期限付き
オープンソース
✅ マサチューセッツ工科大学
❌ 独自の
✅あなたのコード
特長
📞 音声通話 — シンプルな API を通じて実際の電話をかけたり受けたりします
🎙️ AI 音声パイプライン — 内蔵 STT (ディープグラム) + LLM (GPT-4o) + TTS (デカルト) パイプライン
💬 SMS — 受信テキストメッセージを受信して読む
🔌 MCP サーバー — クロード デスクトップとカーソルのネイティブ モデル コンテキスト プロトコルのサポート
📋 スキル ファイル — 任意の AI エージェント (クロード コード、カーソル、OpenClaw) を 1 つのファイルでインストール
🌍 マルチプロバイダー — プラグイン可能なプロバイダー アーキテクチャを備えた SignalWire (US)
📝 トランスクリプト — 完全な会話履歴を含む自動通話トランスクリプション
📬 イベント メールボックス — Webhook エンドポイントを持たないエージェント向けのポーリングベースのイベント システム
💰 ビルトイン請求 — 残高追跡による秒単位の通話請求
🐳 Docker Ready — 1 つのコマンドでスタック全体をローカルで実行
グラフLR
A["🤖 AI エージェント"] -->|REST API / MCP| B[「⚡ AgentLine API」]
B -->|プロビジョン番号| C["📱シグナルワイヤー"]
B -->|音声テキスト変換| D["🎤 ディープグラム"]
B -->|テキスト読み上げ| E["🔊 デカルト"]
B -->|LLM 推論| F["🧠OpenAI"]
C -->|音声とSMS| G["📞 PSTN ネットワーク"]
B -->|イベントと記録|あ
読み込み中
音声パイプライン — ハイブリッド リレー モード
AgentLine は、脆弱なリアルタイム WebSocket ストリームの代わりに、非同期ハイブリッド リレー アーキテクチャを使用します。
発信者がエージェントの番号にダイヤルします
→ SignalWire が電話に応答します
→ 発信者に対して TTS グリーティングを再生します
→ 通話相手の音声を録音します
→ ディープグラムによる転写（高速かつ正確）
→ LLM がエージェント応答を生成します
→ デカルトが返答を話す
→通話が終了するまでループ継続
→ 検索用に保存された完全なトランスクリプト
クイックスタート
オプション 1: Docker Compose (推奨)
git clone https://github.

com/agentlineHQ/AgentLine.git
cd エージェントライン
cp .env.example .env
# .env に API キーを入力します (以下の設定を参照)
docker-compose up -d
これは始まります:
API サーバー (http://localhost:8000)
localhost:5432 の PostgreSQL (スキーマ自動適用)
git clone https://github.com/agentlineHQ/AgentLine.git
cd エージェントライン
Python -m venv venv
source venv/bin/activate # または Windows では venv\Scripts\activate
pip install -r 要件.txt
cp .env.example .env
# API キーを入力します
uvicorn Agentline.main:app --reload
オプション 3: ホストされたバージョンを使用する
セルフホスティングをスキップします。agentline.cloud にサインアップすると、すぐに API キーを取得できます。
.env.example を .env にコピーし、資格情報を入力します。
実行したら、http://localhost:8000/docs にアクセスして対話型 Swagger UI を表示します。
方法
パス
説明
投稿
/v1/エージェント
新しい AI 音声エージェントを作成する
ゲット
/v1/エージェント
すべてのエージェントをリストする
パッチ
/v1/エージェント/{id}
更新エージェント (プロンプト、音声、挨拶)
投稿
/v1/数値
電話番号をプロビジョニングする
ゲット
/v1/数値
電話番号をリストする
投稿
/v1/呼び出し
発信電話をかける
ゲット
/v1/呼び出し
通話のリストを表示する
ゲット
/v1/calls/{id}/トランスクリプト
通話記録を取得する
投稿
/v1/calls/{id}/ハングアップ
アクティブな通話を終了する
ゲット
/v1/イベント
ポーリングイベントメールボックス（消費）
ゲット
/v1/イベント/ピーク
イベントを覗く (非破壊)
ゲット
/v1/メッセージ
SMS メッセージの一覧表示
ゲット
/v1/請求/残高
口座残高を確認する
ゲット
/v1/請求/支出
支出の内訳
認証
すべてのリクエストには API キーが必要です。
カール -H " 権限: ベアラー sk_live_YOUR_KEY " \
-H " Content-Type: application/json " \
https://api.agentline.cloud/v1/agents
MCPサーバーの統合
AgentLine には MCP (Model Context Protocol) サーバーが組み込まれているため、Claude Desktop や Cursor などの AI エージェントはテレフォニー ツールをネイティブに使用できます。
Claude デスクトップ構成に次を追加します。
Windows: %APPDATA%\Claude\claude_desktop_co

nfig.json
macOS: ~/ライブラリ/Application Support/Claude/claude_desktop_config.json
{
"mcpサーバー": {
"エージェントライン" : {
"コマンド" : " npx " ,
"引数" : [
" -y " 、 " mcp-remote@latest " 、
" http://localhost:8000/mcp " 、
" --header " 、 " 認証: ベアラー sk_live_YOUR_KEY "
】
}
}
}
利用可能な MCP ツール
ツール
説明
エージェントの作成
新しい AI 音声エージェントを作成する
エージェントのリスト
すべてのエージェントをリストする
アップデートエージェント
エージェント設定/プロンプト/音声を更新する
make_outbound_call
発信電話を開始する
list_calls
通話履歴を一覧表示する
get_call_transscript
通話の完全なトランスクリプトを取得する
電話を切る
アクティブな通話を終了する
電話番号を購入する
新しい電話番号をプロビジョニングする
電話番号のリスト
すべての電話番号をリストする
投票イベント
投票イベントメールボックス
ピークイベント
保留中のイベントを確認する
アカウント残高の取得
口座残高を確認する
list_available_voices
音声プリセットの一覧表示
MCP Inspector を使用したテスト
npx @modelcontextprotocol/inspector http://localhost:8000/mcp
スキル ファイル — 30 秒でインストール
AgentLine には、AI エージェントが即座に電話能力を獲得できるスキル ファイルが付属しています。
https://agentline.cloud/skill.md
使用方法:
AI エージェント (Claude Code、Cursor、OpenClaw など) に追加します。
AGENTLINE_API_KEY 環境変数を設定します。
エージェントに「+1234567890 に電話してください」と伝えてください。これで十分です。
スキル ファイルは、このリポジトリの skill/agentline/SKILL.md にも含まれています。
コンポーネント
テクノロジー
APIフレームワーク
FastAPI (非同期 Python)
データベース
PostgreSQL (スーパーベース)
キャッシュ
レディス
電話番号と通話
信号線
音声からテキストへの変換
ディープグラム ノヴァ-2
テキスト読み上げ
デカルトソニック
LLM
OpenAI GPT-4o / GPT-4o-mini
MCPサーバー
FastAPI-MCP
導入
港湾労働者、鉄道
導入
docker build -t エージェントライン 。
docker run -p 8000:8000 --env-file .env エージェントライン
任意のクラウドプロバイダー
AgentLine は、Python 3.12 以降と Docker をサポートする場所であればどこでも実行できます: AWS、GCP、Azure、

Fly.io、レンダリング、鉄道など
データベース スキーマは schema.sql にあります。以下のテーブルを作成します。
accounts — 残高追跡のあるユーザーアカウント
api_keys — 認証用のハッシュされた API キー
エージェント — AI 音声エージェント構成
Phone_numbers — プロビジョニングされた電話番号
通話 — トランスクリプト付きの通話記録
メッセージ — SMS メッセージ レコード
events_mailbox — サーバー側のイベント キュー
billing_ledger — 不変の請求トランザクション ログ
移行は、migrations/ ディレクトリにあります。
Agentline.cloud でホストされているバージョンを使用している場合:
セルフホスト型: プロバイダー料金 (SignalWire、Deepgram、Cartesia、OpenAI) のみを支払います。
寄付を歓迎します!ガイドラインについては、CONTRIBUTING.md を参照してください。
💡 機能リクエスト — 問題を開く
🔧 プルリクエスト — フォーク、ブランチ、PR
💬 Discord — チームや他のビルダーとチャットします
🐦 Twitter — 更新情報とお知らせ
📖 ドキュメント — 完全な API ドキュメント
MIT — あらゆる用途に使用できます。商用利用歓迎です。
AgentLine チームによって ❤️ を使用して構築されました
AI エージェントに声を与えましょう。 ☎️
AI エージェントに実際の電話番号を与えます。音声通話、SMS、MCP サーバー、スキル ファイルを備えたオープンソースのテレフォニー API。 FastAPI + SignalWire + Deepgram で構築されています。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
3
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give your AI agent a real phone number. Open-source telephony API with voice calls, SMS, MCP server, and skill file. Built with FastAPI + SignalWire + Deepgram. - AgentLineHQ/AgentLine

GitHub - AgentLineHQ/AgentLine: Give your AI agent a real phone number. Open-source telephony API with voice calls, SMS, MCP server, and skill file. Built with FastAPI + SignalWire + Deepgram. · GitHub
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
AgentLineHQ
/
AgentLine
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github .github agentline agentline migrations migrations skills/ agentline skills/ agentline .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md agentline-logo-200.png agentline-logo-200.png agentline-logo.png agentline-logo.png claude_desktop_config_example.json claude_desktop_config_example.json docker-compose.yml docker-compose.yml railway.toml railway.toml requirements.txt requirements.txt schema.sql schema.sql View all files Repository files navigation
Open-source telephony API for AI agents — phone numbers, voice calls, and SMS
Give your AI agent a real phone number. Make outbound calls, receive inbound calls, and handle SMS — all through one API. No telecom expertise needed.
Website · Docs · Skill File · Discord
AgentLine is an open-source AI-native telephony platform that gives AI agents real phone numbers and human-like voices. It provides a simple REST API and MCP server so AI agents (Claude, Cursor, custom LLM agents) can make and receive phone calls, handle SMS, and retrieve transcripts — without any telecom knowledge.
Your AI Agent → AgentLine API → Real Phone Calls
→ SMS Messages
→ Call Transcripts
Why AgentLine?
AgentLine
Twilio/Vonage
Build It Yourself
Built for AI agents
✅ Purpose-built
❌ Built for call centers
❌ You stitch it together
Setup time
5 minutes
Hours
Weeks
MCP server
✅ Native
❌ None
❌ Build your own
Skill file install
✅ One file
❌ Complex SDK
❌ Hundreds of lines
Voice pipeline
✅ Included (STT + TTS + LLM)
❌ BYO
❌ BYO
Open source
✅ MIT
❌ Proprietary
✅ Your code
Features
📞 Voice Calls — Make and receive real phone calls through a simple API
🎙️ AI Voice Pipeline — Built-in STT (Deepgram) + LLM (GPT-4o) + TTS (Cartesia) pipeline
💬 SMS — Receive and read inbound text messages
🔌 MCP Server — Native Model Context Protocol support for Claude Desktop and Cursor
📋 Skill File — One-file install for any AI agent (Claude Code, Cursor, OpenClaw)
🌍 Multi-Provider — SignalWire (US) with pluggable provider architecture
📝 Transcripts — Automatic call transcription with full conversation history
📬 Event Mailbox — Poll-based event system for agents without webhook endpoints
💰 Built-in Billing — Per-second call billing with balance tracking
🐳 Docker Ready — One command to run the entire stack locally
graph LR
A["🤖 AI Agent"] -->|REST API / MCP| B["⚡ AgentLine API"]
B -->|Provision Numbers| C["📱 SignalWire"]
B -->|Speech-to-Text| D["🎤 Deepgram"]
B -->|Text-to-Speech| E["🔊 Cartesia"]
B -->|LLM Reasoning| F["🧠 OpenAI"]
C -->|Voice & SMS| G["📞 PSTN Network"]
B -->|Events & Transcripts| A
Loading
Voice Pipeline — Hybrid Relay Mode
AgentLine uses an asynchronous Hybrid Relay architecture instead of fragile real-time WebSocket streams:
Caller dials your agent's number
→ SignalWire answers the call
→ Plays TTS greeting to caller
→ Records caller's speech
→ Deepgram transcribes (fast, accurate)
→ LLM generates agent response
→ Cartesia speaks the response
→ Loop continues until call ends
→ Full transcript stored for retrieval
Quick Start
Option 1: Docker Compose (Recommended)
git clone https://github.com/agentlineHQ/AgentLine.git
cd AgentLine
cp .env.example .env
# Fill in your API keys in .env (see Configuration below)
docker-compose up -d
This starts:
API server at http://localhost:8000
PostgreSQL at localhost:5432 (schema auto-applied)
git clone https://github.com/agentlineHQ/AgentLine.git
cd AgentLine
python -m venv venv
source venv/bin/activate # or venv\Scripts\activate on Windows
pip install -r requirements.txt
cp .env.example .env
# Fill in your API keys
uvicorn agentline.main:app --reload
Option 3: Use the Hosted Version
Skip self-hosting — sign up at agentline.cloud and get an API key instantly.
Copy .env.example to .env and fill in your credentials:
Once running, visit http://localhost:8000/docs for the interactive Swagger UI.
Method
Path
Description
POST
/v1/agents
Create a new AI voice agent
GET
/v1/agents
List all agents
PATCH
/v1/agents/{id}
Update agent (prompt, voice, greeting)
POST
/v1/numbers
Provision a phone number
GET
/v1/numbers
List phone numbers
POST
/v1/calls
Make an outbound call
GET
/v1/calls
List calls
GET
/v1/calls/{id}/transcript
Get call transcript
POST
/v1/calls/{id}/hangup
End an active call
GET
/v1/events
Poll event mailbox (consume)
GET
/v1/events/peek
Peek at events (non-destructive)
GET
/v1/messages
List SMS messages
GET
/v1/billing/balance
Check account balance
GET
/v1/billing/expenditure
Spending breakdown
Authentication
All requests require an API key:
curl -H " Authorization: Bearer sk_live_YOUR_KEY " \
-H " Content-Type: application/json " \
https://api.agentline.cloud/v1/agents
MCP Server Integration
AgentLine includes a built-in MCP (Model Context Protocol) server , so AI agents like Claude Desktop and Cursor can use telephony tools natively.
Add to your Claude Desktop config:
Windows: %APPDATA%\Claude\claude_desktop_config.json
macOS: ~/Library/Application Support/Claude/claude_desktop_config.json
{
"mcpServers" : {
"agentline" : {
"command" : " npx " ,
"args" : [
" -y " , " mcp-remote@latest " ,
" http://localhost:8000/mcp " ,
" --header " , " Authorization: Bearer sk_live_YOUR_KEY "
]
}
}
}
Available MCP Tools
Tool
Description
create_agent
Create a new AI voice agent
list_agents
List all agents
update_agent
Update agent config/prompt/voice
make_outbound_call
Initiate an outbound phone call
list_calls
List call history
get_call_transcript
Get the full transcript of a call
hangup_call
End an active call
buy_phone_number
Provision a new phone number
list_phone_numbers
List all phone numbers
poll_events
Poll event mailbox
peek_events
Peek at pending events
get_account_balance
Check account balance
list_available_voices
List voice presets
Test with MCP Inspector
npx @modelcontextprotocol/inspector http://localhost:8000/mcp
Skill File — Install in 30 Seconds
AgentLine ships with a skill file that lets any AI agent gain telephony powers instantly:
https://agentline.cloud/skill.md
How to use:
Add it to your AI agent (Claude Code, Cursor, OpenClaw, etc.)
Set your AGENTLINE_API_KEY environment variable
Tell your agent: "Call +1234567890" — it just works!
The skill file is also included in this repo at skills/agentline/SKILL.md .
Component
Technology
API Framework
FastAPI (async Python)
Database
PostgreSQL ( Supabase )
Cache
Redis
Phone Numbers & Calls
SignalWire
Speech-to-Text
Deepgram Nova-2
Text-to-Speech
Cartesia Sonic
LLM
OpenAI GPT-4o / GPT-4o-mini
MCP Server
FastAPI-MCP
Deployment
Docker, Railway
Deployment
docker build -t agentline .
docker run -p 8000:8000 --env-file .env agentline
Any Cloud Provider
AgentLine runs anywhere that supports Python 3.12+ and Docker: AWS, GCP, Azure, Fly.io, Render, Railway, etc.
The database schema is in schema.sql . It creates tables for:
accounts — User accounts with balance tracking
api_keys — Hashed API keys for authentication
agents — AI voice agent configurations
phone_numbers — Provisioned phone numbers
calls — Call records with transcripts
messages — SMS message records
event_mailbox — Server-side event queue
billing_ledger — Immutable billing transaction log
Migrations are in the migrations/ directory.
If using the hosted version at agentline.cloud :
Self-hosted: you pay only your provider costs (SignalWire, Deepgram, Cartesia, OpenAI).
We welcome contributions! See CONTRIBUTING.md for guidelines.
💡 Feature requests — Open an issue
🔧 Pull requests — Fork, branch, PR
💬 Discord — Chat with the team and other builders
🐦 Twitter — Updates and announcements
📖 Docs — Full API documentation
MIT — use it for anything. Commercial use welcome.
Built with ❤️ by the AgentLine team
Give your AI agent a voice. ☎️
Give your AI agent a real phone number. Open-source telephony API with voice calls, SMS, MCP server, and skill file. Built with FastAPI + SignalWire + Deepgram.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
3
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

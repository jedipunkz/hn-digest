---
source: "https://github.com/Tboworst/beacon"
hn_url: "https://news.ycombinator.com/item?id=49234995"
title: "Beacon: A self-hosted error tracking and LLM observability in one place"
article_title: "GitHub - Tboworst/beacon · GitHub"
author: "Tboworst"
captured_at: "2026-08-09T20:20:28Z"
capture_tool: "hn-digest"
hn_id: 49234995
score: 2
comments: 0
posted_at: "2026-08-09T19:42:24Z"
tags:
  - hacker-news
  - translated
---

# Beacon: A self-hosted error tracking and LLM observability in one place

- HN: [49234995](https://news.ycombinator.com/item?id=49234995)
- Source: [github.com](https://github.com/Tboworst/beacon)
- Score: 2
- Comments: 0
- Posted: 2026-08-09T19:42:24Z

## Translation

タイトル: Beacon: 自己ホスト型エラー追跡と LLM 可観測性を 1 か所にまとめたもの
記事タイトル: GitHub - Tboworst/beacon · GitHub
説明: GitHub でアカウントを作成して、Tboworst/beacon の開発に貢献します。

記事本文:
GitHub - Tboworst/ビーコン · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
最悪
/
ビーコン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
45 コミット 45 コミット コア コア ダッシュボード ダッシュボード sdk sdk テスト テスト Web Web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile README.md README.md docker-compose.yml docker-compose.yml required.txt required.txt sam

ple.jsonlサンプル.jsonlシミュレート.pyシミュレート.pyシミュレート_相関.pyシミュレート_相関.pyシミュレート_llm.pyシミュレート_llm.py start_dashboard.py start_dashboard.py start_server.py start_server.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
制御を必要とする開発者向けのセルフホスト型エラー追跡および監視。
Beacon は、あらゆるサービスからエラーを取り込み、根本原因ごとにグループ化し、すべてをライブ ターミナル ダッシュボードに表示します。サードパーティ サービスは使用せず、インフラストラクチャからデータが流出することも、月額料金も発生しません。
あなたが所有する、端末用に構築された軽量の Sentry を考えてください。
Beacon_updated_WebUI_compressed.mp4
git クローン https://github.com/Tboworst/beacon.git
CDビーコン
cp .env.example .env # API キーと Slack Webhook を追加します
docker-構成アップ
サーバーは http://localhost:7000 で実行されています。
別のターミナルでダッシュボードを開きます。
pip インストールテキスト
python3 start_dashboard.py
クイックスタート — Docker なし
git クローン https://github.com/Tboworst/beacon.git
CDビーコン
pip install -r 要件.txt
cp .env.example .env
python3 start_server.py # ターミナル 1
python3 start_dashboard.py # ターミナル 2
ウェブダッシュボード
Beacon にはブラウザ ダッシュボードも付属しています。これは TUI と同じデータで、取り込みサーバーから直接提供されます。検索、環境フィルター、解決/再オープン、GitHub 問題の作成を伴う、問題、LLM 呼び出し、アラートとデプロイ。
一度ビルドすれば、残りはサーバーが実行します (Node.js 18 以降が必要です。Docker ユーザーはこれをスキップします。イメージによって自動的にビルドされます)。
CDウェブ
npmインストール
npm ビルドを実行する
通常どおりサーバーを起動し、 http://localhost:7000 を開きます。
python3 start_server.py
フロントエンド開発の場合は、Vite dev サーバーを並行して実行します。これは、/api を取り込みサーバーにプロキシし、ホットリロードします。
python3 start_server.py # ターミナル 1
cd web && npm run dev # ターミナル 2 → http://localhost:5

173
TUI は変更されずに動作し続けます。両方のダッシュボードは同じ beacon.db を読み取ります。
macOS の注意: AirPlay Receiver はポート 7000 でも待機します。 http://localhost:7000 が誤動作する場合は、 http://127.0.0.1:7000 を使用してください。
sdk/node/ をプロジェクトにコピーします (または、公開されたら npm install beacon-monitor)。
const beacon = require ( './beacon' ) ;
ビーコン。初期化 ( {
エンドポイント: 'http://your-beacon-server:7000/ingest' 、
サービス : 'マイアプリ' 、
環境 : '実稼働' 、
apiKey : 'あなたの秘密キーをここに' ,
} ) ;
未処理の例外と Promise の拒否は自動的にキャプチャされます。捕捉されたエラーの場合:
{を試してください
リスキーオペレーション() ;
} キャッチ (エラー) {
ビーコン。キャプチャ (エラー) ;
}
LLM 通話追跡:
const t0 = 日付。今 （ ） ;
{を試してください
const res = openai を待ちます。チャット 。完成品。 create ( { モデル : 'gpt-4o' , メッセージ } ) ;
ビーコン。 CaptureLlm ( {
モデル: 'gpt-4o' 、
inputTokens : res 。使用方法。プロンプトトークン 、
出力トークン: res 。使用方法。完了トークン 、
latencyMs: 日付。今 ( ) - t0 、
コスト米ドル: 解像度。使用方法。プロンプト_トークン * 0.000005 + res 。使用方法。 completed_tokens * 0.000015 、
機能: 'ドキュメントサマライザー' 、
} ) ;
} キャッチ (エラー) {
ビーコン。 CaptureLlm ( { モデル : 'gpt-4o' 、inputTokens : 0 、outputTokens : 0 、
latencyMs: 日付。現在 ( ) - t0 、コストUSD : 0 、
機能: 'ドキュメントサマライザー'、エラー: err } ;
}
依存関係なし — Node の組み込み http / https モジュールのみを使用します。
pipインストールリクエスト
sdk/python/beacon/ をプロジェクトにコピーし、エントリ ポイントに 2 行を追加します。
インポートビーコン
ビーコン。初期化(
エンドポイント = "http://your-beacon-server:7000/ingest" ,
サービス = "私のアプリ" 、
環境 = "本番" 、
api_key = "your-secret-key-here" # .env の BEACON_API_KEY と一致します
）
未処理の例外はすべて自動的にキャプチャされるようになりました。処理された例外の場合:
試してみてください:
危険な操作 ()
例外を除く

eとして：
ビーコン。キャプチャ ( e )
構成
変数
説明
BEACON_API_KEY
すべての取り込みリクエストで秘密キーが必要です。設定されていない場合、サーバーはすべてのリクエストを受け入れます (ローカル開発モード)。
SLACK_WEBHOOK_URL
Slack の受信 Webhook URL。エラー グループがしきい値を超えると、アラートが起動されます。
エラーのグループ化方法
生のエラー: 「NoneType には属性 'email' がありません」
「NoneType には属性 'username' がありません」
フィンガープリント: 属性エラー
NoneType には属性 <attr> がありません ← 正規化されています
handle_request → get_current_user → find_user_by_token
結果：同じグループ。バグが1つ。 1行。
フィンガープリント = 例外タイプのハッシュ + 正規化されたメッセージ + 関数呼び出しチェーン。
行番号は無視され、再フォーマットするたびに変更されます。関数名は安定しています。
ビーコン/
§── core/ ← サーバー、ストレージ、フィンガープリンティングを取り込む (Python → Go)
§── ダッシュボード/ ← ライブ TUI ダッシュボード (テキスト)
§── web/ ← Web ダッシュボード (Vite + React、インジェスト サーバーによって提供)
§── SDK/
│ └── python/ ← Python SDK
§── start_server.py ← python3 start_server.py
§── start_dashboard.py ← python3 start_dashboard.py
§── docker-compose.yml
└── 要件.txt
あらゆる言語からのエラーの送信
どのサービスでも HTTP 経由でエラーを直接送信できます。SDK は必要ありません。
curl -X POST http://localhost:7000/ingest \
-H " Content-Type: application/json " \
-H " X-Api-Key: あなたの秘密キーをここに " \
-d ' {
"タイムスタンプ": "2024-01-15T10:23:45Z",
"例外タイプ": "属性エラー",
"message": "NoneType オブジェクトには電子メール属性がありません",
"スタックトレース": [
{"関数": "handle_request", "ファイル": "/app/server.py", "行": 42},
{"関数": "get_current_user"、"ファイル": "/app/auth.py"、"行": 87}
】
} '
ロードマップ
デプロイマーカー — エラーとデプロイを関連付けます
環境のタグ付け — 個別の本番環境とステージング環境
スピ

ke 検出 — 合計数だけでなく増加率に関するアラート
回帰アラート — エラーは 7 日間静かだったが、突然再び発生する
TUI からのグループを解決/無視する
任意のエラー グループからの GitHub 問題の作成
Redis ホット パスを使用してコア サーバーを書き換える
レイヤー
技術
インジェストサーバー
Python + フラスコ
ストレージ
SQLite
ダッシュボード
Python + テキスト
アラート
Slack Webhook
コンテナ化
ドッカー
なぜセントリーを使わないのでしょうか?
セントリーは優秀だよ。 Beacon は次のようなときに使用します。
完全に読み取り、変更し、所有できるもの
公共の場で建てられました。星が評価されました。
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Tboworst/beacon development by creating an account on GitHub.

GitHub - Tboworst/beacon · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Tboworst
/
beacon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
45 Commits 45 Commits core core dashboard dashboard sdk sdk tests tests web web .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile README.md README.md docker-compose.yml docker-compose.yml requirements.txt requirements.txt sample.jsonl sample.jsonl simulate.py simulate.py simulate_correlated.py simulate_correlated.py simulate_llm.py simulate_llm.py start_dashboard.py start_dashboard.py start_server.py start_server.py View all files Repository files navigation
Self-hosted error tracking and monitoring for developers who want control.
Beacon ingests errors from any service, groups them by root cause, and surfaces everything in a live terminal dashboard — no third-party service, no data leaving your infrastructure, no monthly bill.
Think lightweight Sentry, built for the terminal, owned by you.
Beacon_updated_WebUI_compressed.mp4
git clone https://github.com/Tboworst/beacon.git
cd beacon
cp .env.example .env # add your API key and Slack webhook
docker-compose up
Server is running at http://localhost:7000 .
Open the dashboard in a separate terminal:
pip install textual
python3 start_dashboard.py
Quick start — without Docker
git clone https://github.com/Tboworst/beacon.git
cd beacon
pip install -r requirements.txt
cp .env.example .env
python3 start_server.py # terminal 1
python3 start_dashboard.py # terminal 2
Web dashboard
Beacon also ships a browser dashboard — same data as the TUI, served straight from the ingest server. Issues, LLM calls, alerts and deploys, with search, environment filters, resolve / reopen, and GitHub issue creation.
Build it once, then the server does the rest (requires Node.js 18+; Docker users skip this — the image builds it automatically):
cd web
npm install
npm run build
Start the server as usual and open http://localhost:7000 :
python3 start_server.py
For frontend development, run the Vite dev server alongside — it proxies /api to the ingest server and hot-reloads:
python3 start_server.py # terminal 1
cd web && npm run dev # terminal 2 → http://localhost:5173
The TUI keeps working unchanged — both dashboards read the same beacon.db .
macOS note: AirPlay Receiver also listens on port 7000. If http://localhost:7000 misbehaves, use http://127.0.0.1:7000 .
Copy sdk/node/ into your project (or npm install beacon-monitor once published):
const beacon = require ( './beacon' ) ;
beacon . init ( {
endpoint : 'http://your-beacon-server:7000/ingest' ,
service : 'my-app' ,
environment : 'production' ,
apiKey : 'your-secret-key-here' ,
} ) ;
Unhandled exceptions and promise rejections are captured automatically. For caught errors:
try {
riskyOperation ( ) ;
} catch ( err ) {
beacon . capture ( err ) ;
}
LLM call tracking:
const t0 = Date . now ( ) ;
try {
const res = await openai . chat . completions . create ( { model : 'gpt-4o' , messages } ) ;
beacon . captureLlm ( {
model : 'gpt-4o' ,
inputTokens : res . usage . prompt_tokens ,
outputTokens : res . usage . completion_tokens ,
latencyMs : Date . now ( ) - t0 ,
costUsd : res . usage . prompt_tokens * 0.000005 + res . usage . completion_tokens * 0.000015 ,
feature : 'document-summarizer' ,
} ) ;
} catch ( err ) {
beacon . captureLlm ( { model : 'gpt-4o' , inputTokens : 0 , outputTokens : 0 ,
latencyMs : Date . now ( ) - t0 , costUsd : 0 ,
feature : 'document-summarizer' , error : err } ) ;
}
Zero dependencies — uses Node's built-in http / https modules only.
pip install requests
Copy sdk/python/beacon/ into your project, then add two lines to your entry point:
import beacon
beacon . init (
endpoint = "http://your-beacon-server:7000/ingest" ,
service = "my-app" ,
environment = "production" ,
api_key = "your-secret-key-here" # matches BEACON_API_KEY in .env
)
Every unhandled exception is now automatically captured. For handled exceptions:
try :
risky_operation ()
except Exception as e :
beacon . capture ( e )
Configuration
Variable
Description
BEACON_API_KEY
Secret key required on all ingest requests. If not set, server accepts all requests (local dev mode).
SLACK_WEBHOOK_URL
Slack incoming webhook URL. Alerts fire when an error group crosses the threshold.
How errors are grouped
raw errors: "NoneType has no attribute 'email'"
"NoneType has no attribute 'username'"
fingerprint: AttributeError
NoneType has no attribute <attr> ← normalized
handle_request → get_current_user → find_user_by_token
result: same group. one bug. one row.
Fingerprint = hash of exception type + normalized message + function call chain.
Line numbers are ignored — they change on every reformat. Function names are stable.
beacon/
├── core/ ← ingest server, storage, fingerprinting (Python → Go)
├── dashboard/ ← live TUI dashboard (Textual)
├── web/ ← web dashboard (Vite + React, served by the ingest server)
├── sdk/
│ └── python/ ← Python SDK
├── start_server.py ← python3 start_server.py
├── start_dashboard.py ← python3 start_dashboard.py
├── docker-compose.yml
└── requirements.txt
Sending errors from any language
Any service can send errors directly over HTTP — no SDK required:
curl -X POST http://localhost:7000/ingest \
-H " Content-Type: application/json " \
-H " X-Api-Key: your-secret-key-here " \
-d ' {
"timestamp": "2024-01-15T10:23:45Z",
"exception_type": "AttributeError",
"message": "NoneType object has no attribute email",
"stack_trace": [
{"function": "handle_request", "file": "/app/server.py", "line": 42},
{"function": "get_current_user", "file": "/app/auth.py", "line": 87}
]
} '
Roadmap
Deploy markers — correlate errors with deploys
Environment tagging — separate prod vs staging
Spike detection — alert on rate of increase, not just total count
Regression alerts — error quiet for 7 days that suddenly fires again
Resolve / ignore groups from the TUI
GitHub issue creation from any error group
Go rewrite of the core server with Redis hot path
Layer
Tech
Ingest server
Python + Flask
Storage
SQLite
Dashboard
Python + Textual
Alerts
Slack webhooks
Containerisation
Docker
Why not just use Sentry?
Sentry is excellent. Beacon is for when you want:
Something you can read, modify, and own completely
Built in public. Stars appreciated.
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

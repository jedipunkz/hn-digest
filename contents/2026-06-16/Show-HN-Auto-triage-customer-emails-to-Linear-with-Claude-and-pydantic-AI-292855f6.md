---
source: "https://github.com/Reactance0083/pydantic-ai-email-linear-auto-triage"
hn_url: "https://news.ycombinator.com/item?id=48554622"
title: "Show HN: Auto-triage customer emails to Linear with Claude and pydantic-AI"
article_title: "GitHub - Reactance0083/pydantic-ai-email-linear-auto-triage: Email to Linear Issue Auto-Triage — FastAPI webhook that classifies incoming emails with Claude AI and creates prioritized Linear issues with Slack alerts · GitHub"
author: "reactance0083"
captured_at: "2026-06-16T13:56:41Z"
capture_tool: "hn-digest"
hn_id: 48554622
score: 1
comments: 0
posted_at: "2026-06-16T13:02:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Auto-triage customer emails to Linear with Claude and pydantic-AI

- HN: [48554622](https://news.ycombinator.com/item?id=48554622)
- Source: [github.com](https://github.com/Reactance0083/pydantic-ai-email-linear-auto-triage)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:02:00Z

## Translation

タイトル: HN を表示: Claude と pydantic-AI を使用した Linear への顧客メールの自動トリアージ
記事のタイトル: GitHub - Reactance0083/pydantic-ai-email-linear-auto-triage: Email to Linear Issue Auto-Triage — Claude AI で受信メールを分類し、Slack アラートで優先順位の高い Linear 問題を作成する FastAPI Webhook · GitHub
説明: Email to Linear Issue Auto-Triage — Claude AI で受信メールを分類し、Slack アラートで優先順位の高い Linear 問題を作成する FastAPI Webhook - Reactance0083/pydantic-ai-email-linear-auto-triage

記事本文:
GitHub - Reactance0083/pydantic-ai-email-linear-auto-triage: Email to Linear Issue Auto-Triage — Claude AI で受信メールを分類し、Slack アラートで優先順位の高い Linear 問題を作成する FastAPI Webhook · GitHub
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
リアクタンス0083

/
pydantic-ai-email-linear-auto-triage
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Reactance0083/pydantic-ai-email-linear-auto-triage
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .env.example .env.example README.md README.md main.py main.pyRequirements.txt Required.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
電子メール→リニア問題の自動トリアージ
AI を活用したトリアージを使用して、受信メールを優先順位の高い線形問題に自動的に変換します。生の電子メールコンテンツから顧客情報、優先度、問題の種類を抽出し、緊急アイテムの Slack 通知を使用して Linear ワークスペースに構造化されたチケットを即座に作成します。
このテンプレートは、実稼働対応の FastAPI Webhook サービスを提供します。
SMTP 転送または Gmail API 統合を介して電子メールを受け入れます
Claude AI を使用して優先度、顧客、問題の分類を抽出
豊富なメタデータと適切なリンクを使用して線形の問題を作成します
優先度の高いチケットについて Slack アラートを送信します
監査と改良のためにトリアージの決定を保存します
コスト削減: 月額 20 ～ 40 ドルの Zapier 手数料 + 手動トリアージのオーバーヘッドを削減
速度: 電子メールからチケットへのパイプラインは 30 秒、手動ルーティングは 5 分
一貫性: AI による分類により、優先順位の割り当てにおける人的エラーが削減されます。
拡張性: Pydantic AI に基づいて構築されており、トリアージ ロジックを簡単にカスタマイズできます。
生の電子メール POST ペイロード (SMTP Webhook 形式または解析された JSON) を受け入れます
送信者、件名、本文、添付ファイルのメタデータを抽出します
プレーンテキストとHTMLメール本文の両方をサポート
Claude を使用して電子メールのコンテンツから抽出します。
優先度（緊急/高/通常/低）
顧客識別子 (電子メール ドメイン、名前、アカウント ID)
問題の種類 (バグ/機能リクエスト/サポート/請求)
概要 (から自動生成)

件名 + 本文のコンテキスト)
推奨される譲受人 (課題タイプのパターンに基づく、オプション)
Linear ワークスペースで問題を作成します
元の電子メールを問題のコメントとして添付します
トリアージ出力に基づいて優先度とステータスを設定します
顧客/チームプロジェクトへのリンク (構成可能)
電子メールメタデータのカスタムフィールドをサポート
緊急/優先度の高いチケットを指定されたチャネルに投稿します
顧客情報、問題のリンク、優先バッジが含まれます
フォローアップ更新のためのオプションのスレッド返信
すべてのトリアージ決定を SQLite (または構成された DB) に保存します
パフォーマンスの監視とモデルの改良を可能にします
手動オーバーライドとフィードバックループをサポート
リニア API トークン ([設定] > [API] > [パーソナル API キー] で作成)
Claude API キー (Anthropic Console から)
Slack Webhook URL (オプション、Slack Apps から)
Gmail API 認証情報または SMTP リレー サービス (オプション、電子メール取り込み用)
Docker + Docker Compose (コンテナ化されたデプロイメント用)
PostgreSQL (運用データベースの場合、デフォルトは SQLite)
1. 依存関係のクローンを作成してインストールする
git clone https://github.com/yourusername/email-linear-triage.git
cd メール-リニア-トリアージ
Python -m venv venv
source venv/bin/activate # Windows の場合: venv\Scripts\activate
pip install -r 要件.txt
2. 環境ファイルの作成
プロジェクトのルートに .env を作成します。
API キーの数
ANTHROPIC_API_KEY=sk-ant-...
LINEAR_API_KEY=lin_api_...
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/YOUR/WEBHOOK/URL # オプション
# 線形構成
LINEAR_TEAM_ID=acme # Linear チームのスラッグ (例、linear.app/acme の「acme」)
LINEAR_PROJECT_ID=INB # 受信メールのプロジェクト キー (デフォルト: 'INB')
LINEAR_DEFAULT_STATUS=backlog # 新しい問題の初期ステータス
# 電子メール設定
SMTP_SECRET_TOKEN=your-secret-token-here # Webhook 認証の場合
EMAIL_DOMAIN=あなたのドメイン.com
# データベース (オプション)
DATABASE_URL=sqlite:///./

triage.db # または: postgresql://user:pass@localhost/triage
# 機能フラグ
ENABLE_SLACK_NOTIFICATIONS=true
ENABLE_AUTO_ASSIGN=false
TRIAGE_MODEL=claude-3-5-sonnet-20241022 # 使用するクロード モデル
3. データベースの初期化
Python -m alembic アップグレードヘッド
または SQLite (自動作成) の場合:
python -c " from app.db import init_db; init_db() "
4. サーバーを実行する
uvicorn app.main:app --host 0.0.0.0 --port 8000 --reload
サーバーは http://localhost:8000 で実行されます
オプション A: SMTP 転送 (推奨)
電子メールプロバイダーで、転送ルールを設定します。
送信者: ticket@yourdomain.com
宛先: {your-server}/webhook/email (認証あり)
Google Cloud コンソールで Gmail API を有効にする
資格情報 JSON を ./credentials.json にダウンロードします
アプリはラベル付きメールを定期的に自動取得します
curl -X POST http://localhost:8000/webhook/email \
-H " X-Webhook-Token: your-secret-token-here " \
-H " Content-Type: application/json " \
-d ' {
"from": "customer@example.com",
"subject": "支払い処理が失敗しました",
"body": "こんにちは。定期請求書が 2 日間請求されていません。緊急です!",
"タイムスタンプ": "2024-01-15T14:30:00Z"
} '
6. Linear Webhook の設定 (オプション、将来の統合用)
[線形設定] > [統合] > [Webhook] で、以下を追加します。
URL: {あなたのサーバー}/webhook/linear-event
イベント: 問題の作成、問題の更新
メール返信で問題を解決するのに便利
電子メールは ticket@yourdomain.com に到着します (SMTP 経由で転送)
Webhook ハンドラーが POST を受信し、トークンを検証する
クロードのトリアージによる電子メールの分類 (2 ～ 5 秒)
抽出されたメタデータを使用して作成された線形の問題
Slack 通知の投稿 (緊急/緊急の場合)
問題の URL を含む応答が返されました
curl -X POST http://localhost:8000/webhook/email \
-H " X-Webhook-Token: your-secret-token-here " \
-H " Content-Type: application/json " \
-d ' {
"差出人": "sarah@acmecorp.com",
"subject": "[バグ] ダッシュボードがクラッシュする

モバイルで」、
"body": "iPhone でダッシュボードを開くと、すぐにクラッシュします。毎回起こります。私たちのチームは仕事ができません。」
"タイムスタンプ": "2024-01-15T09:30:00Z"
} '
応答 (201 作成):
{
"ステータス" : " 成功 " 、
"linear_issue_id" : " INB-234 " 、
"linear_issue_url" : " https://linear.app/acme/issue/INB-234 " ,
"トリアージ結果" : {
"優先度" : "緊急" 、
"issue_type" : "バグ" ,
"customer_domain" : " acmecorp.com " 、
"summary" : " iOS でダッシュボード モバイル アプリがクラッシュする " ,
"suggested_assignee" : " eng-mobile "
}、
"slack_notification_sent" : true 、
"処理時間_ミリ秒" : 3200
}
APIエンドポイント
生の電子メールを取り込み、リニア問題を作成する
X-Webhook トークン: {SMTP_SECRET_TOKEN}
コンテンツタイプ: application/json
リクエスト本文:
{
"from" : " customer@example.com " 、
"件名" : " 問題のタイトル " ,
"body" : " メール本文 " ,
"html_body" : " <p>HTML バージョン (オプション)</p> " ,
"タイムスタンプ" : " 2024-01-15T10:00:00Z " ,
「添付ファイル」: [
{
"ファイル名" : "スクリーンショット.png " ,
"content_base64" : " iVBORw0KGgoAAAANS... " ,
"mime_type" : " 画像/png "
}
】
}
応答 (201 作成):
{
"ステータス" : " 成功|エラー " 、
"linear_issue_id" : " INB-123 " 、
"linear_issue_url" : " string " ,
"トリアージ結果" : {
"優先度" : " 緊急|高|通常|低 " ,
"issue_type" : " バグ|機能|サポート|請求 " ,
"customer_domain" : " string " ,
"customer_name" : " 文字列 (オプション) " ,
"概要" : " 文字列 " ,
"suggested_assignee" : " 文字列 (オプション) "
}、
"slack_notification_sent" : ブール値、
"error" : " string (if status='error') "
}
POST /api/triage/override/{issue_id}
AI トリアージの決定を手動で上書きする
X-API キー: {LINEAR_API_KEY}
コンテンツタイプ: application/json
リクエスト本文:
{
"優先度" : " 高 " 、
"issue_type" : "バグ" ,
"notes" : " 文脈により「低」から手動で修正されました "
}
応答 (200 OK):
{
"ステータス" : "

更新されました」、
"triage_record_id" : " uuid " ,
「変更」: {
"優先度" : { "古い" : " 通常 " 、 "新しい" : " 高 " }
}
}
GET /api/triage/history
トリアージ履歴とメトリクスを取得する
priority_filter=緊急|高|通常|低 (オプション)
date_from=2024-01-01 (オプション)
{
"total_processed" : 342 、
「結果」: [
{
"id" : "uuid" ,
"email_from" : " customer@example.com " 、
"linear_issue_id" : " INB-234 " 、
"優先度" : " 高 " 、
"issue_type" : "バグ" ,
"created_at" : " 2024-01-15T09:30:00Z " ,
"処理時間_ミリ秒" : 3200 、
「モデルの信頼性」: 0.94
}
]、
"統計" : {
"avg_processing_time_ms" : 2800 、
"優先順位分布" : {
「緊急」：15、
「高い」：87、
「通常」 : 198 、
「低」：42
}、
"issue_type_distribution" : {
「バグ」：124、
「特徴」：56、
「サポート」：142、
「請求」：20
}
}
}
GET /健康
{
"ステータス" : " 健康 " 、
"タイムスタンプ" : " 2024-01-15T10:00:00Z " ,
"依存関係" : {
"人類" : "わかりました" ,
"リニア" : " OK " 、
"たるみ" : "わかりました" ,
"データベース" : " OK "
}
}
構成
変数
必須
デフォルト
説明
ANTHROPIC_API_KEY
✓
—
Anthropic Console の Claude API キー
LINEAR_API_KEY
✓
—
[設定] > [API] からのリニア API トークン
LINEAR_TEAM_ID
✓
—
リニアチームスラッグ（例：「acme」）
LINEAR_PROJECT_ID
INB
新しい問題のためのリニアプロジェクトキー
LINEAR_DEFAULT_STATUS
バックログ
初期課題ステータス (バックログ/todo/進行中)
SMTP_SECRET_TOKEN
✓
—
Webhook認証用のシークレットトークン
SLACK_WEBHOOK_URL
—
Slack Webhook URL (無効にする場合は空のままにします)
ENABLE_SLACK_NOTIFICATIONS
本当の
緊急/緊急の通知を投稿する
ENABLE_AUTO_ASSIGN
偽
問題の種類に基づいて自動的に割り当てる
TRIAGE_MODEL
クロード-3-5-ソネット-20241022
クロード モデル (最高の精度のための 3-opus-20250219)
DATABASE_URL
sqlite:///./triage.db
PostgreSQL または SQLite 接続文字列
GMAIL_CREDENTIALS_PATH
./credenti

als.json
Gmail API 認証情報へのパス (Gmail を使用している場合)
EMAIL_DOMAIN
—
電子メール ドメイン (返信先ヘッダー用)
LOG_LEVEL
情報
ログレベル (DEBUG/INFO/WARNING/ERROR)
Pydantic AI 構成
app/config.py を編集してカスタマイズします。
# クロードモデルの設定
TRIAGE_MODEL = "claude-3-5-sonnet-20241022" # 精度を高めるために claude-3-opus-20250219 に変更します
# トリアージ分類のしきい値
PRIORITY_KEYWORDS = {
"緊急" : [ "クリティカル" 、 "ダウン" 、 "故障" 、 "できるだけ早く" 、 "緊急" ],
"高" : [ "バグ" 、 "壊れた" 、 "失敗" 、 "緊急" ]、
"通常" : [ "機能" , "改善" ],
"低" : [ "タイプミス" 、 "軽度" 、 "あれば便利" ]
}
# 線形フィールドマッピング
LINEAR_PRIORITY_MAP = {
"urgent" : 4 , # リニアでの緊急
「高」 : 3 、
「通常」 : 2 、
「低」：1
}
カスタマイズ
app/agents/triage_agent.py を編集します。
TRIAGE_SYSTEM_PROMPT = """あなたはカスタマー サポート トリアージ システムのエキスパートです...
メールを分析して以下を抽出します。
1. 優先順位 (緊急/高/通常/低) - 顧客のトーン、サービスへの影響、頻度を考慮します。
2. 問題の種類 (バグ/機能/サポート/請求) - 性質別に分類
3. 顧客識別子 - ドメインまたは会社名の抽出
4. 簡潔な要約 - 最大 10 単語
5. 提案されるチーム - 問題の種類に基づく
「」
カスタム問題フィールドの追加
app/models/triage.py の ext

[切り捨てられた]

## Original Extract

Email to Linear Issue Auto-Triage — FastAPI webhook that classifies incoming emails with Claude AI and creates prioritized Linear issues with Slack alerts - Reactance0083/pydantic-ai-email-linear-auto-triage

GitHub - Reactance0083/pydantic-ai-email-linear-auto-triage: Email to Linear Issue Auto-Triage — FastAPI webhook that classifies incoming emails with Claude AI and creates prioritized Linear issues with Slack alerts · GitHub
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
Reactance0083
/
pydantic-ai-email-linear-auto-triage
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Reactance0083/pydantic-ai-email-linear-auto-triage
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .env.example .env.example README.md README.md main.py main.py requirements.txt requirements.txt View all files Repository files navigation
Email→Linear Issue Auto-Triage
Automatically convert incoming emails into prioritized Linear issues using AI-powered triage. Extract customer information, priority levels, and issue types from raw email content, then instantly create structured tickets in your Linear workspace with Slack notifications for urgent items.
This template provides a production-ready FastAPI webhook service that:
Accepts emails via SMTP forwarding or Gmail API integration
Extracts priority, customer, and issue classification using Claude AI
Creates Linear issues with rich metadata and proper linking
Sends Slack alerts for high-priority tickets
Stores triage decisions for audit and refinement
Cost savings: Eliminates $20-40/mo Zapier fees + manual triage overhead
Speed: 30-second email-to-ticket pipeline vs. 5-minute manual routing
Consistency: AI-driven classification reduces human error in priority assignment
Extensibility: Built on Pydantic AI for easy customization of triage logic
Accepts raw email POST payloads (SMTP webhook format or parsed JSON)
Extracts sender, subject, body, and attachment metadata
Supports both plain text and HTML email bodies
Uses Claude to extract from email content:
Priority level (urgent/high/normal/low)
Customer identifier (email domain, name, account ID)
Issue type (bug/feature-request/support/billing)
Summary (auto-generated from subject + body context)
Suggested assignee (based on issue type patterns, optional)
Creates issues in your Linear workspace
Attaches original email as issue comment
Sets priority and status based on triage output
Links to customer/team projects (configurable)
Supports custom fields for email metadata
Posts urgent/high-priority tickets to designated channel
Includes customer info, issue link, and priority badge
Optional thread replies for follow-up updates
Stores all triage decisions in SQLite (or configured DB)
Enables performance monitoring and model refinement
Supports manual override and feedback loops
Linear API token (create in Settings > API > Personal API Keys )
Claude API key (from Anthropic Console )
Slack webhook URL (optional, from Slack Apps )
Gmail API credentials or SMTP relay service (optional, for email ingestion)
Docker + Docker Compose (for containerized deployment)
PostgreSQL (for production database, defaults to SQLite)
1. Clone and Install Dependencies
git clone https://github.com/yourusername/email-linear-triage.git
cd email-linear-triage
python -m venv venv
source venv/bin/activate # On Windows: venv\Scripts\activate
pip install -r requirements.txt
2. Create Environment File
Create .env in the project root:
# API Keys
ANTHROPIC_API_KEY=sk-ant-...
LINEAR_API_KEY=lin_api_...
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/YOUR/WEBHOOK/URL # Optional
# Linear Configuration
LINEAR_TEAM_ID=acme # Your Linear team slug (e.g., 'acme' from linear.app/acme)
LINEAR_PROJECT_ID=INB # Project key for incoming emails (default: 'INB')
LINEAR_DEFAULT_STATUS=backlog # Initial status for new issues
# Email Configuration
SMTP_SECRET_TOKEN=your-secret-token-here # For webhook authentication
EMAIL_DOMAIN=yourdomain.com
# Database (optional)
DATABASE_URL=sqlite:///./triage.db # Or: postgresql://user:pass@localhost/triage
# Feature Flags
ENABLE_SLACK_NOTIFICATIONS=true
ENABLE_AUTO_ASSIGN=false
TRIAGE_MODEL=claude-3-5-sonnet-20241022 # Claude model to use
3. Initialize Database
python -m alembic upgrade head
Or for SQLite (auto-created):
python -c " from app.db import init_db; init_db() "
4. Run the Server
uvicorn app.main:app --host 0.0.0.0 --port 8000 --reload
Server runs on http://localhost:8000
Option A: SMTP Forwarding (Recommended)
In your email provider, set up a forward rule:
From: tickets@yourdomain.com
To: {your-server}/webhook/email with authentication
Enable Gmail API in Google Cloud Console
Download credentials JSON to ./credentials.json
App auto-fetches labeled emails periodically
curl -X POST http://localhost:8000/webhook/email \
-H " X-Webhook-Token: your-secret-token-here " \
-H " Content-Type: application/json " \
-d ' {
"from": "customer@example.com",
"subject": "Payment processing is broken",
"body": "Hi, our recurring invoices havent charged for 2 days. This is urgent!",
"timestamp": "2024-01-15T14:30:00Z"
} '
6. Set Linear Webhook (Optional, for future integration)
In Linear Settings > Integrations > Webhooks, add:
URL: {your-server}/webhook/linear-event
Events: Issue created, issue updated
Useful for closing issues via email replies
Email arrives at tickets@yourdomain.com (forwarded via SMTP)
Webhook handler receives POST, validates token
Claude triage classifies email (2-5 seconds)
Linear issue created with extracted metadata
Slack notification posted (if urgent/high)
Response returned with issue URL
curl -X POST http://localhost:8000/webhook/email \
-H " X-Webhook-Token: your-secret-token-here " \
-H " Content-Type: application/json " \
-d ' {
"from": "sarah@acmecorp.com",
"subject": "[BUG] Dashboard crashes on mobile",
"body": "When I open the dashboard on iPhone, it instantly crashes. Happens every time. Our team cant work.",
"timestamp": "2024-01-15T09:30:00Z"
} '
Response (201 Created):
{
"status" : " success " ,
"linear_issue_id" : " INB-234 " ,
"linear_issue_url" : " https://linear.app/acme/issue/INB-234 " ,
"triage_result" : {
"priority" : " urgent " ,
"issue_type" : " bug " ,
"customer_domain" : " acmecorp.com " ,
"summary" : " Dashboard mobile app crashes on iOS " ,
"suggested_assignee" : " eng-mobile "
},
"slack_notification_sent" : true ,
"processing_time_ms" : 3200
}
API Endpoints
Ingest raw email and create Linear issue
X-Webhook-Token: {SMTP_SECRET_TOKEN}
Content-Type: application/json
Request Body:
{
"from" : " customer@example.com " ,
"subject" : " Issue title " ,
"body" : " Email body text " ,
"html_body" : " <p>HTML version (optional)</p> " ,
"timestamp" : " 2024-01-15T10:00:00Z " ,
"attachments" : [
{
"filename" : " screenshot.png " ,
"content_base64" : " iVBORw0KGgoAAAANS... " ,
"mime_type" : " image/png "
}
]
}
Response (201 Created):
{
"status" : " success|error " ,
"linear_issue_id" : " INB-123 " ,
"linear_issue_url" : " string " ,
"triage_result" : {
"priority" : " urgent|high|normal|low " ,
"issue_type" : " bug|feature|support|billing " ,
"customer_domain" : " string " ,
"customer_name" : " string (optional) " ,
"summary" : " string " ,
"suggested_assignee" : " string (optional) "
},
"slack_notification_sent" : boolean,
"error" : " string (if status='error') "
}
POST /api/triage/override/{issue_id}
Manually override AI triage decision
X-API-Key: {LINEAR_API_KEY}
Content-Type: application/json
Request Body:
{
"priority" : " high " ,
"issue_type" : " bug " ,
"notes" : " Manually corrected from 'low' due to context "
}
Response (200 OK):
{
"status" : " updated " ,
"triage_record_id" : " uuid " ,
"changes" : {
"priority" : { "old" : " normal " , "new" : " high " }
}
}
GET /api/triage/history
Retrieve triage history and metrics
priority_filter=urgent|high|normal|low (optional)
date_from=2024-01-01 (optional)
{
"total_processed" : 342 ,
"results" : [
{
"id" : " uuid " ,
"email_from" : " customer@example.com " ,
"linear_issue_id" : " INB-234 " ,
"priority" : " high " ,
"issue_type" : " bug " ,
"created_at" : " 2024-01-15T09:30:00Z " ,
"processing_time_ms" : 3200 ,
"model_confidence" : 0.94
}
],
"statistics" : {
"avg_processing_time_ms" : 2800 ,
"priority_distribution" : {
"urgent" : 15 ,
"high" : 87 ,
"normal" : 198 ,
"low" : 42
},
"issue_type_distribution" : {
"bug" : 124 ,
"feature" : 56 ,
"support" : 142 ,
"billing" : 20
}
}
}
GET /health
{
"status" : " healthy " ,
"timestamp" : " 2024-01-15T10:00:00Z " ,
"dependencies" : {
"anthropic" : " ok " ,
"linear" : " ok " ,
"slack" : " ok " ,
"database" : " ok "
}
}
Configuration
Variable
Required
Default
Description
ANTHROPIC_API_KEY
✓
—
Claude API key from Anthropic Console
LINEAR_API_KEY
✓
—
Linear API token from Settings > API
LINEAR_TEAM_ID
✓
—
Linear team slug (e.g., 'acme')
LINEAR_PROJECT_ID
INB
Linear project key for new issues
LINEAR_DEFAULT_STATUS
backlog
Initial issue status (backlog/todo/in_progress)
SMTP_SECRET_TOKEN
✓
—
Secret token for webhook authentication
SLACK_WEBHOOK_URL
—
Slack webhook URL (leave empty to disable)
ENABLE_SLACK_NOTIFICATIONS
true
Post notifications for urgent/high
ENABLE_AUTO_ASSIGN
false
Automatically assign based on issue type
TRIAGE_MODEL
claude-3-5-sonnet-20241022
Claude model (3-opus-20250219 for best accuracy)
DATABASE_URL
sqlite:///./triage.db
PostgreSQL or SQLite connection string
GMAIL_CREDENTIALS_PATH
./credentials.json
Path to Gmail API credentials (if using Gmail)
EMAIL_DOMAIN
—
Your email domain (for reply-to headers)
LOG_LEVEL
INFO
Logging level (DEBUG/INFO/WARNING/ERROR)
Pydantic AI Configuration
Edit app/config.py to customize:
# Claude model settings
TRIAGE_MODEL = "claude-3-5-sonnet-20241022" # Change to claude-3-opus-20250219 for higher accuracy
# Triage classification thresholds
PRIORITY_KEYWORDS = {
"urgent" : [ "critical" , "down" , "broken" , "asap" , "emergency" ],
"high" : [ "bug" , "broken" , "failing" , "urgent" ],
"normal" : [ "feature" , "improve" ],
"low" : [ "typo" , "minor" , "nice-to-have" ]
}
# Linear field mappings
LINEAR_PRIORITY_MAP = {
"urgent" : 4 , # Urgent in Linear
"high" : 3 ,
"normal" : 2 ,
"low" : 1
}
Customization
Edit app/agents/triage_agent.py :
TRIAGE_SYSTEM_PROMPT = """You are an expert customer support triage system...
Analyze the email and extract:
1. Priority (urgent/high/normal/low) - consider customer tone, service impact, frequency
2. Issue type (bug/feature/support/billing) - classify by nature
3. Customer identifier - extract domain or company name
4. Concise summary - max 10 words
5. Suggested team - based on issue type
"""
Add Custom Issue Fields
In app/models/triage.py , ext

[truncated]

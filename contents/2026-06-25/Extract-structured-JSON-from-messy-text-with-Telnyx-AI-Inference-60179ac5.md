---
source: "https://github.com/team-telnyx/telnyx-code-examples/tree/main/extract-structured-json-with-ai-python"
hn_url: "https://news.ycombinator.com/item?id=48679998"
title: "Extract structured JSON from messy text with Telnyx AI Inference"
article_title: "telnyx-code-examples/extract-structured-json-with-ai-python at main · team-telnyx/telnyx-code-examples · GitHub"
author: "sona-coffee11"
captured_at: "2026-06-25T22:29:20Z"
capture_tool: "hn-digest"
hn_id: 48679998
score: 1
comments: 0
posted_at: "2026-06-25T22:25:51Z"
tags:
  - hacker-news
  - translated
---

# Extract structured JSON from messy text with Telnyx AI Inference

- HN: [48679998](https://news.ycombinator.com/item?id=48679998)
- Source: [github.com](https://github.com/team-telnyx/telnyx-code-examples/tree/main/extract-structured-json-with-ai-python)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T22:25:51Z

## Translation

タイトル: Telnyx AI 推論を使用して乱雑なテキストから構造化された JSON を抽出する
記事のタイトル: telnyx-code-examples/extract-structord-json-with-ai-python at main · チーム-telnyx/telnyx-code-examples · GitHub
説明: Telnyx AI 通信インフラストラクチャの運用対応コード例 — 音声 AI、SMS、SIP、IoT API - メインの telnyx-code-examples/extract-structurald-json-with-ai-python · Team-telnyx/telnyx-code-examples

記事本文:
メインの telnyx-code-examples/extract-structord-json-with-ai-python · チーム Telnyx/telnyx-code-examples · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
チームテルニクス
/
telnyx コードの例
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
AI-Python を使用した構造化された json の抽出
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
AI-Python を使用した構造化された json の抽出
コピーパスのトップフォルダーとファイル
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py要件.txt 要件.txt すべてのファイルを表示 README.md
概要
名前
AI を使用した構造化された json の抽出
タイトル
AI を使用して構造化された JSON を抽出する
説明
Telnyx AI Inference を使用して、サポート チケット、電子メール、リード、またはインシデント レポートから構造化された JSON を抽出します。
言語
パイソン
フレームワーク
フラスコ
テルニクス製品
AI推論
AI を使用して構造化された JSON を抽出する
Telnyx AI Inference を使用して、サポート チケット、電子メール、リード、またはインシデント レポートから構造化された JSON を抽出します。
AI 推論 : POST /v2/ai/chat/completions - API リファレンス
非構造化テキスト
|
v
Flask APIはリクエストを検証します
|
v
Telnyx AI 推論による JSON の抽出
|
v
構造化された JSON 応答
環境変数
.env.example を .env にコピーし、次のように入力します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/extract-structord-json-with-ai-python
cp .env.example .env
pip install -r 要件.txt
Python app.py
APIリファレンス
テキストから構造化データを抽出します。スキーマを指定しない場合、アプリはデフォルトのサポート チケット スキーマを使用します。
curl -X POST http://localhost:5000/extract \
-H " Content-Type: application/json " \
-d ' {
"text": "アカウント: Acme Health。API キーのローテーション後に実稼働検証ジョブが失敗し始めました。ユーザーはサインアップを完了できません。ログには 401 エラーが表示されます。"
} '
応答:
{
"モデル" : "月翔体/君-K2.6" 、
"結果" : {
"会社" : " アクメ ヘルス " 、
"カテゴリ" : "認証" ,
"優先度" : "緊急" 、
"概要" : " 製品

API キーのローテーション後に検証ジョブが失敗します。 「、
"affected_environment" : "本番環境" ,
"affected_region" : " 不明 " ,
"customer_impact" : " ユーザーはサインアップを完了できません。 「、
"error_codes" : [ " 401 " ],
"suspected_cause" : " 新しい API キーが無効であるか、必要な権限が不足している可能性があります。 「、
"requested_action" : " API キーの設定と権限を確認します。 」
}
}
GET /サンプル
サンプル テキストとデフォルトのスキーマを返します。
サービスのステータスと構成されたモデルを返します。
問題
原因
修正
401 不正
Telnyx API キーが無効または欠落しています
.env で TELNYX_API_KEY を確認する
400 不正なリクエスト
テキストが欠落しているか、スキーマが無効です
空ではないテキスト文字列と JSON オブジェクト スキーマを送信する
502 モデル応答が有効な JSON ではありませんでした
選択したモデルは解析可能な JSON を返しませんでした
デフォルトのモデルで再試行するか、スキーマを簡素化してください
関連する例
Telnyx 推論を使用して RAG を構築する (Python)
構造化データ パイプラインへの FAX (Python)
Telnyx は、音声、メッセージング、SIP、AI、IoT を 1 つのプライベートなグローバル ネットワーク上で実現する AI コミュニケーション インフラストラクチャ プラットフォームです。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-ready code examples for Telnyx AI Communications Infrastructure — Voice AI, SMS, SIP, and IoT APIs - telnyx-code-examples/extract-structured-json-with-ai-python at main · team-telnyx/telnyx-code-examples

telnyx-code-examples/extract-structured-json-with-ai-python at main · team-telnyx/telnyx-code-examples · GitHub
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
extract-structured-json-with-ai-python
More options Directory actions
History History main Breadcrumbs
extract-structured-json-with-ai-python
Copy path Top Folders and files
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py requirements.txt requirements.txt View all files README.md
Outline
name
extract-structured-json-with-ai
title
Extract Structured JSON with AI
description
Extract structured JSON from support tickets, emails, leads, or incident reports with Telnyx AI Inference.
language
python
framework
flask
telnyx_products
AI Inference
Extract Structured JSON with AI
Extract structured JSON from support tickets, emails, leads, or incident reports with Telnyx AI Inference.
AI Inference : POST /v2/ai/chat/completions - API reference
Unstructured text
|
v
Flask API validates request
|
v
Telnyx AI Inference extracts JSON
|
v
Structured JSON response
Environment Variables
Copy .env.example to .env and fill in:
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/extract-structured-json-with-ai-python
cp .env.example .env
pip install -r requirements.txt
python app.py
API Reference
Extract structured data from text. If you do not provide a schema, the app uses a default support-ticket schema.
curl -X POST http://localhost:5000/extract \
-H " Content-Type: application/json " \
-d ' {
"text": "Account: Acme Health. Production verification jobs started failing after an API key rotation. Users cannot finish signup. Logs show 401 errors."
} '
Response:
{
"model" : " moonshotai/Kimi-K2.6 " ,
"result" : {
"company" : " Acme Health " ,
"category" : " authentication " ,
"priority" : " urgent " ,
"summary" : " Production verification jobs are failing after an API key rotation. " ,
"affected_environment" : " production " ,
"affected_region" : " unknown " ,
"customer_impact" : " Users cannot finish signup. " ,
"error_codes" : [ " 401 " ],
"suspected_cause" : " The new API key may be invalid or missing required permissions. " ,
"requested_action" : " Check API key configuration and permissions. "
}
}
GET /sample
Returns sample text and the default schema.
Returns service status and configured model.
Issue
Cause
Fix
401 Unauthorized
Invalid or missing Telnyx API key
Verify TELNYX_API_KEY in .env
400 Bad Request
Missing text or invalid schema
Send a non-empty text string and a JSON object schema
502 Model response was not valid JSON
The selected model did not return parseable JSON
Retry with the default model or simplify the schema
Related Examples
Build RAG with Telnyx Inference (Python)
Fax to Structured Data Pipeline (Python)
Telnyx is an AI Communications Infrastructure platform - voice, messaging, SIP, AI, and IoT on one private, global network.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

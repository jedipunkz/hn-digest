---
source: "https://github.com/team-telnyx/telnyx-code-examples/tree/main/build-rag-with-telnyx-inference-python"
hn_url: "https://news.ycombinator.com/item?id=48687698"
title: "Build a Simple RAG App with Telnyx AI Inference"
article_title: "telnyx-code-examples/build-rag-with-telnyx-inference-python at main · team-telnyx/telnyx-code-examples · GitHub"
author: "sona-coffee11"
captured_at: "2026-06-26T15:44:50Z"
capture_tool: "hn-digest"
hn_id: 48687698
score: 1
comments: 1
posted_at: "2026-06-26T15:19:05Z"
tags:
  - hacker-news
  - translated
---

# Build a Simple RAG App with Telnyx AI Inference

- HN: [48687698](https://news.ycombinator.com/item?id=48687698)
- Source: [github.com](https://github.com/team-telnyx/telnyx-code-examples/tree/main/build-rag-with-telnyx-inference-python)
- Score: 1
- Comments: 1
- Posted: 2026-06-26T15:19:05Z

## Translation

タイトル: Telnyx AI 推論を使用したシンプルな RAG アプリの構築
記事のタイトル: メインの telnyx-code-examples/build-rag-with-telnyx-inference-python · チーム-telnyx/telnyx-code-examples · GitHub
説明: Telnyx AI 通信インフラストラクチャの本番対応コード例 — 音声 AI、SMS、SIP、IoT API - telnyx-code-examples/build-rag-with-telnyx-inference-python at main · team-telnyx/telnyx-code-examples

記事本文:
メインの telnyx-code-examples/build-rag-with-telnyx-inference-python · チーム-telnyx/telnyx-code-examples · GitHub
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
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
build-rag-with-telnyx-inference-python
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
build-rag-with-telnyx-inference-python
コピーパスのトップフォルダーとファイル
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py要件.txt 要件.txt すべてのファイルを表示 README.md
概要
名前
Telnyx 推論を使用したビルドラグ
タイトル
Telnyx 推論を使用して RAG を構築する
説明
Telnyx 埋め込みとチャット補完を使用して、検索拡張生成 API を構築します。
言語
パイソン
フレームワーク
フラスコ
テルニクス製品
AI推論
Telnyx 推論を使用して RAG を構築する
Telnyx 埋め込みとチャット補完を使用して、検索拡張生成 API を構築します。
Embeddings : POST /v2/ai/embeddings - ドキュメントと質問のベクトルを作成します
AI 推論 : POST /v2/ai/chat/completions - API リファレンス
ユーザーの質問
|
v
Telnyx に質問を埋め込む
|
v
ドキュメントの埋め込みと比較する
|
v
取得したコンテキストを Telnyx AI に送信する
|
v
根拠のある回答 + ソースのタイトル
環境変数
.env.example を .env にコピーし、次のように入力します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/build-rag-with-telnyx-inference-python
cp .env.example .env
pip install -r 要件.txt
Python app.py
APIリファレンス
質問してください。アプリは、そのコンテキストのみを使用して、関連するメモリ内のサポート ドキュメントと回答を取得します。
curl -X POST http://localhost:5000/rag/ask \
-H " Content-Type: application/json " \
-d ' {
"question": "API キーをローテーションした後、実稼働サインアップが中断されました。ログに 401 エラーが表示されます。最初に何を確認する必要がありますか?"
} '
応答:
{
"answer" : " 運用サービスが新しいアクティブな API キーを使用していること、およびそのキーに必要な権限があることを確認してください。

また、古いキーがデプロイメント シークレットにキャッシュされていないことも確認します。 「、
"モデル" : "月翔体/君-K2.6" 、
"embedding_model" : " thenlper/gte-large " ,
「ソース」: [
{ "title" : " API キー認証 " , "スコア" : 0.9123 },
{ "title" : "検証メッセージ配信" , "スコア" : 0.7811 }
】
}
GET /ドキュメント
サンプルのナレッジ ベースを返します。
サービスのステータス、構成されたモデル、およびドキュメント数を返します。
問題
原因
修正
401 不正
Telnyx API キーが無効または欠落しています
.env で TELNYX_API_KEY を確認する
最初のリクエストが遅い
アプリはドキュメントの埋め込みを遅延して作成します
最初のリクエストにはさらに時間がかかる場合があります。後のリクエストはメモリ内の埋め込みを再利用します
弱い答え
サンプルナレッジベースが小さすぎます
さらにドキュメントを追加するか、ドキュメントを独自のコンテンツに置き換えます
関連する例
AI を使用して構造化された JSON を抽出する (Python)
AI アシスタント ナレッジ ベース (Python)
Telnyx は、音声、メッセージング、SIP、AI、IoT を 1 つのプライベートなグローバル ネットワーク上で実現する AI コミュニケーション インフラストラクチャ プラットフォームです。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-ready code examples for Telnyx AI Communications Infrastructure — Voice AI, SMS, SIP, and IoT APIs - telnyx-code-examples/build-rag-with-telnyx-inference-python at main · team-telnyx/telnyx-code-examples

telnyx-code-examples/build-rag-with-telnyx-inference-python at main · team-telnyx/telnyx-code-examples · GitHub
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
build-rag-with-telnyx-inference-python
More options Directory actions
History History main Breadcrumbs
build-rag-with-telnyx-inference-python
Copy path Top Folders and files
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py requirements.txt requirements.txt View all files README.md
Outline
name
build-rag-with-telnyx-inference
title
Build RAG with Telnyx Inference
description
Build a retrieval-augmented generation API with Telnyx embeddings and chat completions.
language
python
framework
flask
telnyx_products
AI Inference
Build RAG with Telnyx Inference
Build a retrieval-augmented generation API with Telnyx embeddings and chat completions.
Embeddings : POST /v2/ai/embeddings - create vectors for documents and questions
AI Inference : POST /v2/ai/chat/completions - API reference
User question
|
v
Embed question with Telnyx
|
v
Compare against document embeddings
|
v
Send retrieved context to Telnyx AI
|
v
Grounded answer + source titles
Environment Variables
Copy .env.example to .env and fill in:
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/build-rag-with-telnyx-inference-python
cp .env.example .env
pip install -r requirements.txt
python app.py
API Reference
Ask a question. The app retrieves relevant in-memory support docs and answers using only that context.
curl -X POST http://localhost:5000/rag/ask \
-H " Content-Type: application/json " \
-d ' {
"question": "Production signup broke after rotating an API key. Logs show 401 errors. What should we check first?"
} '
Response:
{
"answer" : " Check that production services are using the new active API key and that the key has the required permissions. Also verify no old key is cached in deployment secrets. " ,
"model" : " moonshotai/Kimi-K2.6 " ,
"embedding_model" : " thenlper/gte-large " ,
"sources" : [
{ "title" : " API Key Authentication " , "score" : 0.9123 },
{ "title" : " Verification Message Delivery " , "score" : 0.7811 }
]
}
GET /documents
Returns the sample knowledge base.
Returns service status, configured models, and document count.
Issue
Cause
Fix
401 Unauthorized
Invalid or missing Telnyx API key
Verify TELNYX_API_KEY in .env
Slow first request
The app creates document embeddings lazily
First request may take longer; later requests reuse embeddings in memory
Weak answers
Sample knowledge base is too small
Add more documents or replace DOCUMENTS with your own content
Related Examples
Extract Structured JSON with AI (Python)
AI Assistant Knowledge Base (Python)
Telnyx is an AI Communications Infrastructure platform - voice, messaging, SIP, AI, and IoT on one private, global network.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

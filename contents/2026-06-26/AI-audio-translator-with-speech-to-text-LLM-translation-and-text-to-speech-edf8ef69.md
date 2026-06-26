---
source: "https://github.com/team-telnyx/telnyx-code-examples/tree/main/ai-content-translator-python"
hn_url: "https://news.ycombinator.com/item?id=48692436"
title: "AI audio translator with speech-to-text, LLM translation, and text-to-speech"
article_title: "telnyx-code-examples/ai-content-translator-python at main · team-telnyx/telnyx-code-examples · GitHub"
author: "sona-coffee11"
captured_at: "2026-06-26T22:25:00Z"
capture_tool: "hn-digest"
hn_id: 48692436
score: 2
comments: 0
posted_at: "2026-06-26T21:50:51Z"
tags:
  - hacker-news
  - translated
---

# AI audio translator with speech-to-text, LLM translation, and text-to-speech

- HN: [48692436](https://news.ycombinator.com/item?id=48692436)
- Source: [github.com](https://github.com/team-telnyx/telnyx-code-examples/tree/main/ai-content-translator-python)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T21:50:51Z

## Translation

タイトル: 音声からテキストへの変換、LLM 翻訳、テキストから音声への変換を備えた AI オーディオ翻訳者
記事のタイトル: メインの telnyx-code-examples/ai-content-translator-python · チーム-telnyx/telnyx-code-examples · GitHub
説明: Telnyx AI 通信インフラストラクチャの本番対応コード例 — 音声 AI、SMS、SIP、IoT API - メインの telnyx-code-examples/ai-content-translator-python · Team-telnyx/telnyx-code-examples

記事本文:
メインのtelnyx-code-examples/ai-content-translator-python · チーム-telnyx/telnyx-code-examples · GitHub
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
あなたはきっとsiでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. .env.example .env.example API.md API.md GUIDE.md GUIDE.md README.md README.md app.py app.py要件.txt 要件.txt すべてのファイルを表示 README.md
概要
名前
AIコンテンツトランスレータ
タイトル
AI コンテンツ トランスレーター
説明
任意の音声 (ポッドキャスト、会議、講義) をアップロードし、STT がソース言語で書き起こし、AI 推論が翻訳し、TTS がターゲット言語で音声を生成します。翻訳された音声と整列されたトランスクリプトを返します。
言語
パイソン
フレームワーク
フラスコ
テルニクス製品
AI推論
メディアストリーミング
統合
チャンネル
声
API
AI コンテンツ トランスレーター
任意の音声 (ポッドキャスト、会議、講義) をアップロードし、STT がソース言語で書き起こし、AI 推論が翻訳し、TTS がターゲット言語で音声を生成します。翻訳された音声と整列されたトランスクリプトを返します。
STT 転写: POST /v2/ai/transcribe -- ref
AI 推論 : POST /v2/ai/chat/completions -- ref
TTS 生成 : POST /v2/ai/generate -- ref
APIリクエスト
│
▼
┌─────────┐
│ 応答 + 挨拶 │ ── TTS ウェルカムメッセージ
━───┬─────┘
│
▼
┌─────────┐
│ スピーチの収集 │ ── STT 文字起こし
━───┬─────┘
│
▼
┌─────────┐
│ AI推論 │
│ • 翻訳 │
━───┬─────┘
│ ◄──── 会話ループ
│
▼
JSON応答
仕組み
会話を Telnyx AI Inference に送信して処理します
Telnyx TTS 経由で応答を音声に変換します
Telnyx は AI コミュニケーションです

インフラストラクチャ プラットフォーム - 1 つのプライベートなグローバル ネットワーク上の音声、メッセージング、SIP、AI、IoT。
共存推論 - LLM は音声トラフィックと同じネットワーク上で実行されます。往復で 200 ミリ秒未満。
.env.example を .env にコピーし、次のように入力します。
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-content-translator-python
cp .env.example .env
pip install -r 要件.txt
Python app.py
Webhook の設定
ngrok http 5000
Telnyx ポータルで Webhook URL を設定します。
通話制御アプリケーション -> https://<id>.ngrok.io/webhooks/voice
curl -X POST http://localhost:5000/translate \
-F audio=@lecture.mp3 \
-F ソース=en \
-F ターゲット=ja
応答:
{ "job_id" : " tr-a1b2c3d4 " 、 "status" : " complete " 、 "source" : " en (英語) " 、 "target" : " ja (日本語) " 、 "original_length" : 1847 、 "transrated_length" : 923 }
GET /健康
カール http://localhost:5000/health
{ "ステータス" : " OK " }
トラブルシューティング
ポート 5000 で接続が拒否されました: アプリが実行されていません。 python app.py を実行し、他のプロセスがポート 5000 を使用していないことを確認します。
401 Unauthorized : TELNYX_API_KEY が無効です。 portal.telnyx.com/api-keys で新しいものを生成します。
AIの応答が遅い/空：モデル名を確認してください。利用可能なモデルについては、developers.telnyx.com でご覧ください。
run-llm-inference-python - スタンドアロン推論
build-voice-ai-agent-python - 音声 AI エージェント
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Production-ready code examples for Telnyx AI Communications Infrastructure — Voice AI, SMS, SIP, and IoT APIs - telnyx-code-examples/ai-content-translator-python at main · team-telnyx/telnyx-code-examples

telnyx-code-examples/ai-content-translator-python at main · team-telnyx/telnyx-code-examples · GitHub
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
ai-content-translator
title
AI Content Translator
description
Upload any audio (podcast, meeting, lecture), STT transcribes in source language, AI Inference translates, TTS generates audio in target language. Returns translated audio + aligned transcript.
language
python
framework
flask
telnyx_products
AI Inference
Media Streaming
integrations
channel
voice
api
AI Content Translator
Upload any audio (podcast, meeting, lecture), STT transcribes in source language, AI Inference translates, TTS generates audio in target language. Returns translated audio + aligned transcript.
STT Transcribe : POST /v2/ai/transcribe -- ref
AI Inference : POST /v2/ai/chat/completions -- ref
TTS Generate : POST /v2/ai/generate -- ref
API Request
│
▼
┌──────────────────┐
│ Answer + Greet │ ── TTS welcome message
└────────┬─────────┘
│
▼
┌──────────────────┐
│ Gather Speech │ ── STT transcription
└────────┬─────────┘
│
▼
┌──────────────────┐
│ AI Inference │
│ • Translation │
└────────┬─────────┘
│ ◄──── conversation loop
│
▼
JSON response
How It Works
Sends conversation to Telnyx AI Inference for processing
Converts response to speech via Telnyx TTS
Telnyx is an AI Communications Infrastructure platform - voice, messaging, SIP, AI, and IoT on one private, global network.
Co-located inference - LLM runs on the same network as voice traffic. Sub-200ms round trips.
Copy .env.example to .env and fill in:
git clone https://github.com/team-telnyx/telnyx-code-examples.git
cd telnyx-code-examples/ai-content-translator-python
cp .env.example .env
pip install -r requirements.txt
python app.py
Webhook Configuration
ngrok http 5000
Set webhook URL in Telnyx Portal :
Call Control Application -> https://<id>.ngrok.io/webhooks/voice
curl -X POST http://localhost:5000/translate \
-F audio=@lecture.mp3 \
-F source=en \
-F target=ja
Response:
{ "job_id" : " tr-a1b2c3d4 " , "status" : " complete " , "source" : " en (English) " , "target" : " ja (Japanese) " , "original_length" : 1847 , "translated_length" : 923 }
GET /health
curl http://localhost:5000/health
{ "status" : " ok " }
Troubleshooting
Connection refused on port 5000 : App isn't running. Run python app.py and check no other process uses port 5000.
401 Unauthorized : Your TELNYX_API_KEY is invalid. Generate a new one at portal.telnyx.com/api-keys .
AI response slow/empty : Verify model name. See available models at developers.telnyx.com .
run-llm-inference-python - Standalone inference
build-voice-ai-agent-python - Voice AI agent
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

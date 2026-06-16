---
source: "https://github.com/rNetAi/rnet-oauth-python"
hn_url: "https://news.ycombinator.com/item?id=48553571"
title: "Show HN: Rnet-OAuth-Python – Python lib for user-funded AI API access"
article_title: "GitHub - rNetAi/rnet-oauth-python: Python library for rNet oauth integration, allowing users to authenticate and pay for AI model token costs using their rNet account. · GitHub"
author: "nextma"
captured_at: "2026-06-16T13:58:11Z"
capture_tool: "hn-digest"
hn_id: 48553571
score: 1
comments: 1
posted_at: "2026-06-16T11:29:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rnet-OAuth-Python – Python lib for user-funded AI API access

- HN: [48553571](https://news.ycombinator.com/item?id=48553571)
- Source: [github.com](https://github.com/rNetAi/rnet-oauth-python)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T11:29:20Z

## Translation

タイトル: HN を表示: Rnet-OAuth-Python – ユーザー資金による AI API アクセス用の Python ライブラリ
記事タイトル: GitHub - rNetAi/rnet-oauth-python: rNet oauth 統合用の Python ライブラリ。これにより、ユーザーは rNet アカウントを使用して認証し、AI モデル トークンのコストを支払うことができます。 · GitHub
説明: rNet oauth 統合用の Python ライブラリ。これにより、ユーザーは rNet アカウントを使用して認証し、AI モデル トークンのコストを支払うことができます。 - rNetAi/rnet-oauth-python

記事本文:
GitHub - rNetAi/rnet-oauth-python: rNet oauth 統合用の Python ライブラリ。これにより、ユーザーは rNet アカウントを使用して認証し、AI モデル トークンのコストを支払うことができます。 · GitHub
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
rNetAi
/
rnet-oauth-python
公共
通知
あなたは署名しているに違いありません

通知設定を変更する必要がある
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット example_flask example_flask rnet_oauth rnet_oauth .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
RNet OAuth と AI プロバイダー サービスを統合するための Python バックエンド ライブラリ。このライブラリを使用すると、ユーザーは RNet 経由で認証し、RNet アカウントを使用して AI モデル トークンのコストを直接支払うことができます。
OAuth2 PKCE サポート: 自動コード検証機能とチャレンジ生成による安全な認証コード フロー。
トークン管理: トークンの認証コードを交換し、期限切れのトークンを更新します。
UserInfo Endpoint : 認証されたユーザーの RNet プロファイルをアクセス トークンを使用して取得します。
AI 統合 : 標準応答またはストリーミング応答を使用して AI モデルとチャットする簡単な方法。
pip インストール rnet-oauth
クイックスタート
rnet_oauth からインポート RNetAuth 、RNetAi
認証 = RNetAuth (
client_id = 'クライアント ID' ,
client_secret = 'クライアント シークレット' ,
redirect_uri = 'リダイレクト-uri'
)
ai = RNetAi ()
2. 認可URLの生成(OAuth2 PKCE)
#1. PKCE を生成する
pkce = 認証 .生成_pkce()
検証者 = pkce [ '検証者' ]
チャレンジ = pkce [ 'チャレンジ' ]
# 2. 認可URLの取得
# チャレンジ: PKCE コードチャレンジ (オプション)
# state: リクエストとコールバックの間で状態を維持するためのオプションの文字列 (セキュリティのために推奨)
auth_url = 認証。 get_authorization_url (チャレンジ、状態 = 'オプションの状態')
3. トークンの交換コード
# 3. トークンの交換コード
トークン = 認証 。 Exchange_code_for_token (コード、検証者)
access_token = トークン [ 'access_token' ]
4. ユーザー情報の取得
user_info = 認証 。 get_user_info (アクセストークン)
print ( user_info [ 'email' ])
印刷する

( ユーザー情報 [ '名前' ])
UserInfo 応答は RNet の /userinfo エンドポイントから送信され、次のものが含まれる場合があります。
sub 、 email 、 email_verified 、 name 、preferred_username 、 user_id 、 role 、および status 。
応答 = ai 。チャット ({
「内容」：[
{
"ロール" : "ユーザー" 、
"パーツ" : [{ "テキスト" : "こんにちは!" }]
}
】
}、access_token、「gemini-2.5-flash-lite」）
6. ストリーミング AI 応答 (未テスト)
ai のチャンク用。チャットストリーム ({
「内容」：[
{
"ロール" : "ユーザー" 、
"パーツ" : [{ "テキスト" : "こんにちは!" }]
}
】
}、 access_token 、 "gemini-2.5-flash-lite" ):
印刷 (チャンク)
7. ファイルのアップロード (未テスト)
open ( "document.pdf" , "rb" ) を f として使用:
ファイルバッファ = f 。読んでください()
# Gemini にアップロード
gemini_upload = ai 。 gemini_file_upload ( access_token , "gemini-2.5-flash-lite" , file_buffer , "application/pdf" , "document.pdf" )
print ( gemini_upload [ 'fileReference' ]) # これをチャットペイロードで使用します
# OpenAIにアップロード
openai_upload = ai 。 openai_file_upload ( access_token , "gpt-4o" , file_buffer , "application/pdf" , "document.pdf" )
8. ファイルの削除 (未テスト)
# Gemini ファイルは 48 時間後に自動削除されるため、削除方法はありません。
# OpenAI ファイルを削除します。
あい。 openai_file_delete ( access_token , "gpt-4o" , openai_upload [ 'fileReference' ])
9. ファイルとツールを使用した AI チャット (未テスト)
ペイロード = {
「内容」：[
{
"ロール" : "ユーザー" 、
「パーツ」: [
{ "text" : "この文書に基づくと、私の名前は何ですか? また、ロンドンの現在の天気をウェブで検索してください。" }、
{ "fileData" : { "fileUri" : gemini_upload [ 'fileReference' ], "mimeType" : gemini_upload [ 'mimeType' ] } }
】
}
]、
「ツール」: [
{ "googleSearch" : {} } # Google 検索ツールを有効にする
】
}
応答 = ai 。チャット (ペイロード、アクセストークン、「gemini-2.5-flash-lite」)
印刷（応答）
ライセンス
rNet oauth 統合用の Python ライブラリにより、ユーザーが AI モデル トークンの認証と支払いを行うことが可能になります

rNet アカウントを使用してコストがかかります。
pypi.org/project/rnet-oauth
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Python library for rNet oauth integration, allowing users to authenticate and pay for AI model token costs using their rNet account. - rNetAi/rnet-oauth-python

GitHub - rNetAi/rnet-oauth-python: Python library for rNet oauth integration, allowing users to authenticate and pay for AI model token costs using their rNet account. · GitHub
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
rNetAi
/
rnet-oauth-python
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits example_flask example_flask rnet_oauth rnet_oauth .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
A Python backend library for integrating RNet OAuth and AI Provider services. This library allows users to authenticate via RNet and pay for AI model token costs directly using their RNet account.
OAuth2 PKCE Support : Secure authorization code flow with automatic code verifier and challenge generation.
Token Management : Exchange authorization codes for tokens and refresh expired tokens.
UserInfo Endpoint : Fetch the authenticated user's RNet profile with an access token.
AI Integration : Easy methods to chat with AI models using standard or streaming responses.
pip install rnet-oauth
Quick Start
from rnet_oauth import RNetAuth , RNetAi
auth = RNetAuth (
client_id = 'client-id' ,
client_secret = 'client-secret' ,
redirect_uri = 'redirect-uri'
)
ai = RNetAi ()
2. Generate Authorization URL (OAuth2 PKCE)
# 1. Generate PKCE
pkce = auth . generate_pkce ()
verifier = pkce [ 'verifier' ]
challenge = pkce [ 'challenge' ]
# 2. Get Authorization URL
# challenge: PKCE code challenge (optional)
# state: An optional string to maintain state between the request and callback (recommended for security)
auth_url = auth . get_authorization_url ( challenge , state = 'optional-state' )
3. Exchange Code for Tokens
# 3. Exchange code for tokens
tokens = auth . exchange_code_for_token ( code , verifier )
access_token = tokens [ 'access_token' ]
4. Get User Info
user_info = auth . get_user_info ( access_token )
print ( user_info [ 'email' ])
print ( user_info [ 'name' ])
The UserInfo response comes from RNet's /userinfo endpoint and may include:
sub , email , email_verified , name , preferred_username , user_id , role , and status .
response = ai . chat ({
"contents" : [
{
"role" : "user" ,
"parts" : [{ "text" : "Hello!" }]
}
]
}, access_token , "gemini-2.5-flash-lite" )
6. Streaming AI Response (Untested)
for chunk in ai . chat_stream ({
"contents" : [
{
"role" : "user" ,
"parts" : [{ "text" : "Hello!" }]
}
]
}, access_token , "gemini-2.5-flash-lite" ):
print ( chunk )
7. File Upload (Untested)
with open ( "document.pdf" , "rb" ) as f :
file_buffer = f . read ()
# Upload to Gemini
gemini_upload = ai . gemini_file_upload ( access_token , "gemini-2.5-flash-lite" , file_buffer , "application/pdf" , "document.pdf" )
print ( gemini_upload [ 'fileReference' ]) # Use this in chat payload
# Upload to OpenAI
openai_upload = ai . openai_file_upload ( access_token , "gpt-4o" , file_buffer , "application/pdf" , "document.pdf" )
8. File Deletion (Untested)
# Gemini files auto-delete after 48 hours, so there is no delete method.
# Delete an OpenAI file:
ai . openai_file_delete ( access_token , "gpt-4o" , openai_upload [ 'fileReference' ])
9. AI Chat with File & Tools (Untested)
payload = {
"contents" : [
{
"role" : "user" ,
"parts" : [
{ "text" : "Based on this document, what is my name? Also search the web for the current weather in London." },
{ "fileData" : { "fileUri" : gemini_upload [ 'fileReference' ], "mimeType" : gemini_upload [ 'mimeType' ] } }
]
}
],
"tools" : [
{ "googleSearch" : {} } # Enable Google Search tool
]
}
response = ai . chat ( payload , access_token , "gemini-2.5-flash-lite" )
print ( response )
License
Python library for rNet oauth integration, allowing users to authenticate and pay for AI model token costs using their rNet account.
pypi.org/project/rnet-oauth
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

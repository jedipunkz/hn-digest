---
source: "https://github.com/KishaWeb/local-char"
hn_url: "https://news.ycombinator.com/item?id=48686474"
title: "Hey HN, I built local char an open source local AI platform like character AI"
article_title: "GitHub - KishaWeb/local-char: an locall ai software that you can give your llm commands to act like your wanted character. · GitHub"
author: "KishaWeb"
captured_at: "2026-06-26T13:39:05Z"
capture_tool: "hn-digest"
hn_id: 48686474
score: 1
comments: 0
posted_at: "2026-06-26T13:34:03Z"
tags:
  - hacker-news
  - translated
---

# Hey HN, I built local char an open source local AI platform like character AI

- HN: [48686474](https://news.ycombinator.com/item?id=48686474)
- Source: [github.com](https://github.com/KishaWeb/local-char)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T13:34:03Z

## Translation

タイトル: HN さん、キャラクター AI のようなオープンソースのローカル AI プラットフォームである local char を構築しました
記事のタイトル: GitHub - KishaWeb/local-char: llm コマンドに希望のキャラクターのように動作させることができるローカル AI ソフトウェアです。 · GitHub
説明: llm コマンドに指定して希望のキャラクターのように動作させることができるローカル AI ソフトウェアです。 - KishaWeb/ローカル文字

記事本文:
GitHub - KishaWeb/local-char: llm コマンドに希望のキャラクターのように動作させることができるローカル AI ソフトウェアです。 · GitHub
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
KishaWeb
/
ローカル文字
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.mdhistory.jsonhistory.jsonhistory_tui.jsonhistory_tui.jsonhistory_web.jsonhistory_web.jsonすべてのファイルを表示リポジトリ ファイル ナビゲーション
local char は、希望のキャラクターのように動作できる AI をローカルで実行する cli および Web プログラムです。キャラクターはほとんどが 3 ～ 5 b モデルに向いていますが、あまり強くありません。各モデルの射程強度に応じて異なるカテゴリを追加する予定です。そうすることで、希望のキャラクターのように動作できるようになります。また、クラウド サービスもアカウントもサブスクリプションも必要ありません。
(今のところ) 弱いモデル用に作られています
本当に基本的な Web サーバー (キャラクター選択付き)
このプロジェクトでは、llama.cpp が実行されていることを想定しています。
API エンドポイントを公開するには、llama.cpp をサーバー モードで起動する必要があります。
git clone https://github.com/ggml-org/llama.cpp.git
cd ラマ.cpp
./build/bin/llama-server -m モデルへのパス.gguf
インストールする
llama.cppサーバーが実行されていることを確認してください
git clone https://github.com/KishaWeb/local-char.git
cd ローカル文字
カーゴインストール --path 。
ローカル文字
使用法
local-char tui (tui を実行)
local-char web [--lan] (Web サーバーを起動します。--lan を使用すると LAN 共有が有効になります)
カーゴで実行する場合は、次のように使用します。
カーゴラン -- トゥイ (トゥイを走る)
Cargo run -- web (Web サーバーを開始します。--lan を使用すると LAN 共有が有効になります)
現在の名簿には、からのキャラクターが含まれています
llm コマンドに指定して希望のキャラクターのように動作させることができるローカル AI ソフトウェアです。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。これをリロードしてください

のページ。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

an locall ai software that you can give your llm commands to act like your wanted character. - KishaWeb/local-char

GitHub - KishaWeb/local-char: an locall ai software that you can give your llm commands to act like your wanted character. · GitHub
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
KishaWeb
/
local-char
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md history.json history.json history_tui.json history_tui.json history_web.json history_web.json View all files Repository files navigation
local char is a cli and a web program that runs ai locally that can act like the character you want, the characters are mostly mad for 3-5 b models not very strong im going to add different category for the range strength of each model so it can act better like your desired character, and you dont need any cloud service, no accounts and no subscriptions.
made for weaker models (for now)
really basic web server (with character selecting)
This project expects llama.cpp running .
You need to start llama.cpp in server mode so it exposes an API endpoint:
git clone https://github.com/ggml-org/llama.cpp.git
cd llama.cpp
./build/bin/llama-server -m PATH-TO-YOUR-MODEL.gguf
install
make sure llama.cpp server is running
git clone https://github.com/KishaWeb/local-char.git
cd local-char
cargo install --path .
local-char
usage
local-char tui (runs tui)
local-char web [--lan] (starts the web server, if you use --lan it enables lan sharing)
if your running it in cargo use it like this:
cargo run -- tui (runs tui)
cargo run -- web (starts web server, if you use --lan it enables lan sharing)
current roster includes characters from
an locall ai software that you can give your llm commands to act like your wanted character.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

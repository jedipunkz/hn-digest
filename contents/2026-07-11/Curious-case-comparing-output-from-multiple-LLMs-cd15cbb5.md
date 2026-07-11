---
source: "https://github.com/vseryakov/backendjs/tree/master/examples/prompts"
hn_url: "https://news.ycombinator.com/item?id=48868730"
title: "Curious case, comparing output from multiple LLMs"
article_title: "backendjs/examples/prompts at master · vseryakov/backendjs · GitHub"
author: "vlad1719"
captured_at: "2026-07-11T04:54:52Z"
capture_tool: "hn-digest"
hn_id: 48868730
score: 1
comments: 1
posted_at: "2026-07-11T04:25:55Z"
tags:
  - hacker-news
  - translated
---

# Curious case, comparing output from multiple LLMs

- HN: [48868730](https://news.ycombinator.com/item?id=48868730)
- Source: [github.com](https://github.com/vseryakov/backendjs/tree/master/examples/prompts)
- Score: 1
- Comments: 1
- Posted: 2026-07-11T04:25:55Z

## Translation

タイトル: 複数の LLM からの出力を比較する興味深いケース
記事のタイトル: backendjs/examples/prompts at master · vsryakov/backendjs · GitHub
説明: Node.js バックエンド ライブラリ。 GitHub でアカウントを作成して、vseryakov/backendjs の開発に貢献してください。

記事本文:
マスターの backendjs/examples/prompts · vsryakov/backendjs · GitHub
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
ヴセリヤコフ
/
バックエンドjs
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
その他のオプション ディレクトリアクション
歴史

歴史マスターブレッドクラム
コピーパスのトップフォルダーとファイル
.. モジュール モジュール Web Web README.md README.md bkjs.conf bkjs.conf package.json package.json スクリーンショット.jpg スクリーンショット.jpg スクリーンショット1.jpg スクリーンショット1.jpg スクリーンショット2.jpg スクリーンショット2.jpg すべてのファイルを表示 README.md
異なる LLM からの出力を並べて比較する方法を示す Backendjs の例の概要
これは backendjs モジュール api 、 db 、 jobs 、 ws のデモです。
モデルと結果を保存するデータベース テーブル
リクエストを処理するための API ルート、入力の検証
すべてのモデルに同時にプロンプトを表示し、結果が返されたら保存するバックグラウンド ジョブ
他のすべての結果に対するすべての結果のコサイン類似度を計算し、最も類似している順に並べ替えます
Websocket 経由で Web クライアントに進行状況と結果を通知する
これは backendjs リポジトリ内の例であるため、最初に backendjs をクローンする必要があります
まだ存在しない場合は、存在する場合は次の項目にスキップしてください
git clone -- Depth 1 https://github.com/vseryakov/backendjs.git
次の例に移動します。
cd backendjs/examples/prompts
3.. サンプルを準備して開始します。外部依存関係は必要ありません。
npm 実行セットアップ
npm 実行開始
http://localhost:8000 にアクセスしてください
UI/UX : Alpinejs、Alpinejs-app
ソース/
§── web/
│ §── propmts.js # メインコンポーネント
│ §── Prompts.html # メイン HTML テンプレート
| §── models.html # モデルポップアップ
│ ━──index.html # ホームページ
§── モジュール/
│ §── llm.js # LLM アクセス
│ §── models.js # モデル API
│ └── Prompts.js # プロンプト API
§── var/
└── プロンプト.db # ローカル SQLite データベース
さらに詳しく
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

node.js backend library. Contribute to vseryakov/backendjs development by creating an account on GitHub.

backendjs/examples/prompts at master · vseryakov/backendjs · GitHub
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
vseryakov
/
backendjs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History master Breadcrumbs
Copy path Top Folders and files
.. modules modules web web README.md README.md bkjs.conf bkjs.conf package.json package.json screenshot.jpg screenshot.jpg screenshot1.jpg screenshot1.jpg screenshot2.jpg screenshot2.jpg View all files README.md
Outline Backendjs example showing how to compare output from different LLMs side by side
This is a demonstration of backendjs modules api , db , jobs , ws :
database tables to store models and results
API routes to handle requests, validation of input
a backgrouind job prompting all modesl at the same time and storing results as they come
calculate cosine similarity for every result against all other results, sort by most similar
notifying web clients via Websockets about progress and results
This is an example inside the backendjs repository, so first you need to clone backendjs
it if it does not exist yet, skip to the next item if you have it
git clone --depth 1 https://github.com/vseryakov/backendjs.git
Navigate to the example:
cd backendjs/examples/prompts
3.. Prepare and start the example, no external dependencies are needed
npm run setup
npm run start
Visit http://localhost:8000
UI/UX : Alpinejs, Alpinejs-app
src/
├── web/
│ ├── propmts.js # Main component
│ ├── prompts.html # Main HTML template
| ├── models.html # Models popup
│ └── index.html # Home page
├── modules/
│ ├── llm.js # LLM access
│ ├── models.js # Models API
│ └── prompts.js # Prompts API
├── var/
└── prompt.db # Local SQLite database
Learn More
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

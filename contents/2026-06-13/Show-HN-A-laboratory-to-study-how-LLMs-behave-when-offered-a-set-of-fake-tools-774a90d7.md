---
source: "https://github.com/ampyard/Llaboratory"
hn_url: "https://news.ycombinator.com/item?id=48516028"
title: "Show HN: A laboratory to study how LLMs behave when offered a set of fake tools"
article_title: "GitHub - ampyard/Llaboratory: A laboratory for studying how LLMs behave when offered a set of fake tools · GitHub"
author: "vivganes"
captured_at: "2026-06-13T12:43:18Z"
capture_tool: "hn-digest"
hn_id: 48516028
score: 2
comments: 0
posted_at: "2026-06-13T11:11:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A laboratory to study how LLMs behave when offered a set of fake tools

- HN: [48516028](https://news.ycombinator.com/item?id=48516028)
- Source: [github.com](https://github.com/ampyard/Llaboratory)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T11:11:01Z

## Translation

タイトル: Show HN: 偽のツールのセットが提供されたときに LLM がどのように動作するかを研究する研究室
記事のタイトル: GitHub - ampyard/Llaboratory: 偽のツールのセットが提供されたときに LLM がどのように動作するかを研究するための研究所 · GitHub
説明: 偽のツールのセットが提供されたときに LLM がどのように動作するかを研究するための研究所 - ampyard/Llaboratory
HN テキスト: これは、LLM がツールを見たときにどのように動作するかについての私の好奇心を養うために私が構築したツールです。たとえば、「slap_bad_human」という名前のツールを見たとき、実際にそれを使用するでしょうか、それとも使用しないでしょうか?

記事本文:
GitHub - ampyard/Llaboratory: 偽のツールのセットが提供されたときに LLM がどのように動作するかを研究するための研究所 · GitHub
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
アンヤード
/
研究室
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .github .github .vscode .vscode バックエンド バックエンド フロントエンド フロントエンド .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml logo-small.png logo-small.png logo.png logo.png Original_PRD.md Original_PRD.md スクリーンショット.png スクリーンショット.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
偽のツールのセットが提供されたときに LLM がどのように動作するかを研究するための自己ホスト型のオープンソース ラボ。
docker 構成 --build
UI として http://localhost:5173 を開きます。フロントエンドは、ポート 8000 上のバックエンド コンテナーへの API 呼び出しをプロキシします。
CD バックエンド
紫外線対策
uv pip install -e " .[dev] "
cp ../.env.example ../.env # API キーを入力します
uv run uvicorn app.main:app --reload
フロントエンド
CDフロントエンド
npmインストール
npm 実行開発
http://localhost:5173 を開きます。フロントエンドは /api を :8000 にプロキシします。
ツール ライブラリ → 静的または動的応答を持つ偽のツールを作成します
モデル構成 → プロバイダーエンドポイント + モデルスナップショット + API キー環境変数を構成します
計画 → ツール + モデル + プロンプトをバージョン化されたテスト計画にまとめる
実行→セッションを起動します。ライブ イベント ストリームを視聴します。ツールの呼び出しとモデルの応答を検査する
セッション → 履歴、メトリクス、セッションごとのイベント タイムラインを表示
動的ツール コードは、サンドボックス化せずにインプロセスで実行されます。これはローカルで作成されたツールを意図したものです。
信頼できないソースからの動的コードは決して実行しないでください。完全な理論的根拠については、PRD の §10.6 を参照してください。
CD バックエンド
uv 実行 pytest -v --tb=short
フロントエンド
偽のツールのセットが提供されたときに LLM がどのように動作するかを研究するための研究所
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
貢献者
途中でエラーが発生しました

オーディング。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A laboratory for studying how LLMs behave when offered a set of fake tools - ampyard/Llaboratory

This is a tool I built to feed my curiosity about how LLMs behave when they see tools. For example, when they see a tool named 'slap_bad_human', will they actually use it, or not?

GitHub - ampyard/Llaboratory: A laboratory for studying how LLMs behave when offered a set of fake tools · GitHub
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
ampyard
/
Llaboratory
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .github .github .vscode .vscode backend backend frontend frontend .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml logo-small.png logo-small.png logo.png logo.png original_PRD.md original_PRD.md screenshot.png screenshot.png View all files Repository files navigation
A self-hostable, open-source laboratory for studying how LLMs behave when offered a set of fake tools.
docker compose up --build
Open http://localhost:5173 for the UI. The frontend will proxy API calls to the backend container on port 8000.
cd backend
uv venv
uv pip install -e " .[dev] "
cp ../.env.example ../.env # fill in your API keys
uv run uvicorn app.main:app --reload
Frontend
cd frontend
npm install
npm run dev
Open http://localhost:5173 — the frontend proxies /api to :8000 .
Tool Library → create fake tools with static or dynamic responses
Model Configs → configure a provider endpoint + model snapshot + API key env var
Plans → compose tools + model + prompts into a versioned testing plan
Run → launch sessions; watch the live event stream; inspect tool calls and model responses
Sessions → view history, metrics, and per-session event timelines
Dynamic tool code runs in-process without sandboxing . This is intentional for locally-authored tools.
Never execute dynamic code from untrusted sources. See §10.6 of the PRD for the full rationale.
cd backend
uv run pytest -v --tb=short
Frontend
A laboratory for studying how LLMs behave when offered a set of fake tools
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

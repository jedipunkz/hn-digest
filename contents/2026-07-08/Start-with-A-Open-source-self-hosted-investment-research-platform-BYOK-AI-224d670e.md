---
source: "https://github.com/warlock20/StartWithA"
hn_url: "https://news.ycombinator.com/item?id=48832733"
title: "Start with A – Open-source, self-hosted investment research platform (BYOK AI)"
article_title: "GitHub - warlock20/StartWithA: Value investor's intelligence secretary · GitHub"
author: "warlock_20"
captured_at: "2026-07-08T15:10:40Z"
capture_tool: "hn-digest"
hn_id: 48832733
score: 1
comments: 0
posted_at: "2026-07-08T14:50:01Z"
tags:
  - hacker-news
  - translated
---

# Start with A – Open-source, self-hosted investment research platform (BYOK AI)

- HN: [48832733](https://news.ycombinator.com/item?id=48832733)
- Source: [github.com](https://github.com/warlock20/StartWithA)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T14:50:01Z

## Translation

タイトル: A から始める – オープンソースの自己ホスト型投資調査プラットフォーム (BYOK AI)
記事タイトル: GitHub - warlock20/StartWithA: バリュー投資家の情報秘書 · GitHub
説明: バリュー投資家の情報秘書。 GitHub でアカウントを作成して、warlock20/StartWithA の開発に貢献してください。

記事本文:
GitHub - warlock20/StartWithA: バリュー投資家の情報秘書 · GitHub
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
ウォーロック20
/
StartWithA
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ fi に移動

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
658 コミット 658 コミット .claude/ スキル .claude/ スキル .github/ ワークフロー .github/ ワークフロー アプリ データ データ ドキュメント ドキュメント フロントエンド フロントエンド マイグレーション マイグレーション ユニットテスト ユニットテスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile Dockerfile ライセンス ライセンス Procfile Procfile README.md README.md Architecture_notes.md Architecture_notes.md backfill_currency.py backfill_currency.py backfill_掃引_kills.py backfill_掃引_kills.py celery_app.py celery_app.py config.py config.py database_schema.md database_schema.md docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh init_postgresql.py init_postgresql.py our_vision.md our_vision.md package-lock.json package-lock.json package.json package.json Railway.toml Railway.toml 要件-ローカル.txt 要件-ローカル.txt 要件.txt 要件.txt run.py run.py runtime.txt runtime.txt setup_postgresql.py setup_postgresql.py stratergy.md stratergy.md sync_db.py sync_db.pyvision.mdvision.md vitest.config.js vitest.config.js webpack.config.js webpack.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
リサーチ、ポートフォリオのモニタリング、ジャーナリングを 1 つの規律ある投資ワークフローに結び付けるプラットフォーム。
完全な製品ビジョンについては、vision.md を参照してください。
リサーチ — アイデア受信箱、キルスクリーニング、企業およびセクターのリサーチ、AI アシスタント
ポートフォリオ — 位置追跡、理論と現実、チェックポイント、AI 分析
ジャーナル — 意思決定ジャーナル、間違いログ、学習の洞察
アイデアまたは市場のスイーパー
→ キル画面
→ 研究 ←────────┐
→ 購入決定 │
→ トラック位置更新チェックリスト
→ 日記 │
→学ぶ |───────

──┘
技術スタック
バックエンド: Python 3.12、Flask、PostgreSQL (pgvector)、Celery + Redis
フロントエンド: Jinja2 + React コンポーネント、Bootstrap、Webpack
AI: Gemini、OpenAI、Anthropic (YAML プロンプト テンプレート経由)
プラットフォームをローカルで実行する最も簡単な方法。 Docker と Docker Compose が必要です。
cp .env.example .env
.env を編集し、Auth0 キーと AI キーを入力します (以下の環境変数を参照)。
DATABASE_URL と REDIS_URL は Docker Compose によって自動的に設定されます。.env 内の値はすべてオーバーライドされます。
docker 構成 --build -d
アプリは http://localhost:8000 で入手可能になります。
サービス
説明
港
ウェブ
Gunicorn 経由の Flask アプリ
8000
データベース
PostgreSQL 16 + pgvector
5432
レディス
Redis 7 (Celery ブローカー + キャッシュ)
6379
労働者
セロリのバックグラウンドワーカー
—
データベースの移行は起動時に自動的に実行されます。
# ログを表示する
docker compose ログ -f web
docker compose ログ -f ワーカー
# .env 変更後に再起動
docker compose ウェブの再起動
＃すべてを停止する
ドッカー構成ダウン
# データベースを停止して消去します
docker compose down -v
5. 管理者アクセス
UIからアカウントを登録する
あなたの電子メールが .env の ADMIN_EMAILS にリストされていることを確認してください
再起動: docker compose restart web
開発セットアップ (Docker なし)
Python 3.12 以降、pgvector 拡張機能を備えた PostgreSQL、Redis、および Node.js が必要です。
git クローン https://github.com/warlock20/StartWithA.git
cd StartWithA
# パイソン
python3 -m venv venv && ソース venv/bin/activate
pip install -r 要件.txt
# フロントエンド
npm ci && npm run build
# 環境
cp .env.example .env
# .env を編集 — DATABASE_URL、REDIS_URL、Auth0 キーなどを設定します。
# データベース
フラスコデータベースのアップグレード
# 実行
フラスコの実行
別のターミナルで Celery ワーカーを起動します。
celery - celery_app ワーカー --loglevel=info
環境変数
.env.example を .env にコピーします。 Docker Compose は DATABASE_URL と REDIS_URL を自動的に設定します

— これらは、Docker なしで実行している場合にのみ設定します。
変数
説明
SECRET_KEY
Flask セッション シークレット — 任意のランダムな文字列
AUTH0_ドメイン
Auth0 テナント ドメイン (例: your-tenant.auth0.com )
AUTH0_CLIENT_ID
Auth0 アプリケーションのクライアント ID
AUTH0_CLIENT_SECRET
Auth0 アプリケーション クライアント シークレット
AUTH0_CALLBACK_URL
OAuth コールバック URL (Docker の場合は http://localhost:8000/auth/callback)
AUTH0_AUDIENCE
Auth0 API オーディエンス (通常は https://<domain>/userinfo )
GEMINI_API_KEY
Google Gemini API キー — 少なくとも 1 つの AI プロバイダーが必要です
オプション
変数
説明
デフォルト
OPENAI_API_KEY
OpenAI API キー (代替 AI プロバイダー)
—
ANTHROPIC_API_KEY
Anthropic API キー (代替 AI プロバイダー)
—
NEWS_API_KEY
市場ニュースの NewsAPI キー
—
ADMIN_EMAILS
管理者アクセス権を取得するカンマ区切りのメール
—
DATABASE_URL
PostgreSQL 接続文字列
Docker Compose によって設定される
REDIS_URL
Redis 接続文字列
Docker Compose によって設定される
DEFAULT_USER_TIER
新規ユーザーのデフォルト階層 (無料またはプレミアム)
無料
FLASK_DEBUG
デバッグモードを有効にする
偽
SESSION_COOKIE_SECURE
Cookie に HTTPS を要求する
本当
UPLOAD_FOLDER
ファイルアップロード用のパス
インスタンス/アップロード
Auth0 セットアップ
アプリは認証に Auth0 を使用します。設定するには:
auth0.comで無料アカウントを作成します
通常の Web アプリケーションを作成する
[設定] で、[許可されたコールバック URL] を次のように設定します。
http://localhost:8000/auth/callback (Docker)
http://localhost:5000/auth/callback (ローカルのフラスコ実行)
Domain 、 Client ID 、および Client Secret を .env にコピーします。
セルフホストをしたくないですか?管理されたホスト型インスタンスが利用可能です。ディスカッションを開いてアクセスをリクエストしてください。
リポジトリをフォークして機能ブランチを作成する
アプリをローカルで実行して、すべてが動作することを確認します
PR は常に 1 つの機能または修正に重点を置いてください。
このプロジェクトは、GNU Affero General Public License v3.0 に基づいてライセンスされています。
使用できます、MOD

このソフトウェアを自由に作成し、配布してください。変更したバージョンをサーバー上で実行する場合は、そのサーバーのユーザーがソース コードを利用できるようにする必要があります。詳細については、ライセンスの全文を参照してください。
バリュー投資家の情報秘書
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
最初の安定バージョン
最新の
2026 年 7 月 7 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Value investor's intelligence secretary. Contribute to warlock20/StartWithA development by creating an account on GitHub.

GitHub - warlock20/StartWithA: Value investor's intelligence secretary · GitHub
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
warlock20
/
StartWithA
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
658 Commits 658 Commits .claude/ skills .claude/ skills .github/ workflows .github/ workflows app app data data docs docs frontend frontend migrations migrations unittests unittests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md Dockerfile Dockerfile LICENSE LICENSE Procfile Procfile README.md README.md architecture_notes.md architecture_notes.md backfill_currency.py backfill_currency.py backfill_sweep_kills.py backfill_sweep_kills.py celery_app.py celery_app.py config.py config.py database_schema.md database_schema.md docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh init_postgresql.py init_postgresql.py our_vision.md our_vision.md package-lock.json package-lock.json package.json package.json railway.toml railway.toml requirements-local.txt requirements-local.txt requirements.txt requirements.txt run.py run.py runtime.txt runtime.txt setup_postgresql.py setup_postgresql.py stratergy.md stratergy.md sync_db.py sync_db.py vision.md vision.md vitest.config.js vitest.config.js webpack.config.js webpack.config.js View all files Repository files navigation
A platform that connects research, portfolio monitoring, and journaling into one disciplined investment workflow.
See vision.md for the full product vision.
Research — Idea inbox, kill screening, company & sector research, AI assistant
Portfolio — Position tracking, thesis vs reality, checkpoints, AI analytics
Journal — Decision journal, mistake log, learning insights
Idea or Market Sweeper
→ Kill Screen
→ Research ←──────────────┐
→ Buy Decision │
→ Track Position Updated Checklist
→ Journal │
→ Learn |───────────────┘
Tech Stack
Backend: Python 3.12, Flask, PostgreSQL (pgvector), Celery + Redis
Frontend: Jinja2 + React components, Bootstrap, Webpack
AI: Gemini, OpenAI, and Anthropic (via YAML prompt templates)
The easiest way to run the platform locally. Requires Docker and Docker Compose.
cp .env.example .env
Edit .env and fill in the Auth0 and AI keys (see Environment Variables below).
DATABASE_URL and REDIS_URL are set automatically by Docker Compose — any values in .env are overridden.
docker compose up --build -d
The app will be available at http://localhost:8000 .
Service
Description
Port
web
Flask app via Gunicorn
8000
db
PostgreSQL 16 + pgvector
5432
redis
Redis 7 (Celery broker + cache)
6379
worker
Celery background worker
—
Database migrations run automatically on startup.
# View logs
docker compose logs -f web
docker compose logs -f worker
# Restart after .env changes
docker compose restart web
# Stop everything
docker compose down
# Stop and wipe database
docker compose down -v
5. Admin access
Register an account through the UI
Ensure your email is listed in ADMIN_EMAILS in .env
Restart: docker compose restart web
Development Setup (without Docker)
Requires Python 3.12+, PostgreSQL with the pgvector extension, Redis, and Node.js.
git clone https://github.com/warlock20/StartWithA.git
cd StartWithA
# Python
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
# Frontend
npm ci && npm run build
# Environment
cp .env.example .env
# Edit .env — set DATABASE_URL, REDIS_URL, Auth0 keys, etc.
# Database
flask db upgrade
# Run
flask run
Start the Celery worker in a separate terminal:
celery -A celery_app worker --loglevel=info
Environment Variables
Copy .env.example to .env . Docker Compose sets DATABASE_URL and REDIS_URL automatically — only set those if running without Docker.
Variable
Description
SECRET_KEY
Flask session secret — any random string
AUTH0_DOMAIN
Your Auth0 tenant domain (e.g. your-tenant.auth0.com )
AUTH0_CLIENT_ID
Auth0 application client ID
AUTH0_CLIENT_SECRET
Auth0 application client secret
AUTH0_CALLBACK_URL
OAuth callback URL ( http://localhost:8000/auth/callback for Docker)
AUTH0_AUDIENCE
Auth0 API audience (usually https://<domain>/userinfo )
GEMINI_API_KEY
Google Gemini API key — at least one AI provider is needed
Optional
Variable
Description
Default
OPENAI_API_KEY
OpenAI API key (alternative AI provider)
—
ANTHROPIC_API_KEY
Anthropic API key (alternative AI provider)
—
NEWS_API_KEY
NewsAPI key for market news
—
ADMIN_EMAILS
Comma-separated emails that get admin access
—
DATABASE_URL
PostgreSQL connection string
set by Docker Compose
REDIS_URL
Redis connection string
set by Docker Compose
DEFAULT_USER_TIER
Default tier for new users ( free or premium )
free
FLASK_DEBUG
Enable debug mode
False
SESSION_COOKIE_SECURE
Require HTTPS for cookies
True
UPLOAD_FOLDER
Path for file uploads
instance/uploads
Auth0 Setup
The app uses Auth0 for authentication. To set it up:
Create a free account at auth0.com
Create a Regular Web Application
In Settings, set Allowed Callback URLs to:
http://localhost:8000/auth/callback (Docker)
http://localhost:5000/auth/callback (local flask run )
Copy Domain , Client ID , and Client Secret into your .env
Don't want to self-host? A managed, hosted instance is available — open a discussion to request access.
Fork the repo and create a feature branch
Run the app locally to verify everything works
Please keep PRs focused — one feature or fix per PR.
This project is licensed under the GNU Affero General Public License v3.0 .
You can use, modify, and distribute this software freely. If you run a modified version on a server, you must make your source code available to users of that server. See the full license text for details.
Value investor's intelligence secretary
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
First stable version
Latest
Jul 7, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

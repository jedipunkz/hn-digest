---
source: "https://github.com/tanner49/safe_vibes"
hn_url: "https://news.ycombinator.com/item?id=48479290"
title: "Safe Vibes: Open-source governed workspace for AI-built internal tools"
article_title: "GitHub - tanner49/safe_vibes: A django-based app for creating and deploying AI generated reports in a safe manner. · GitHub"
author: "tannerphillips"
captured_at: "2026-06-10T19:01:47Z"
capture_tool: "hn-digest"
hn_id: 48479290
score: 2
comments: 0
posted_at: "2026-06-10T17:02:06Z"
tags:
  - hacker-news
  - translated
---

# Safe Vibes: Open-source governed workspace for AI-built internal tools

- HN: [48479290](https://news.ycombinator.com/item?id=48479290)
- Source: [github.com](https://github.com/tanner49/safe_vibes)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T17:02:06Z

## Translation

タイトル: Safe Vibes: AI で構築された内部ツール用のオープンソース管理ワークスペース
記事のタイトル: GitHub - Tanner49/safe_vibes: AI によって生成されたレポートを安全な方法で作成およびデプロイするための Django ベースのアプリ。 · GitHub
説明: AI によって生成されたレポートを安全な方法で作成および展開するための Django ベースのアプリです。 - Tanner49/safe_vibes

記事本文:
GitHub - Tanner49/safe_vibes: AI によって生成されたレポートを安全な方法で作成およびデプロイするための Django ベースのアプリ。 · GitHub
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
タンナー49
/
安全な雰囲気
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

オプションについて
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット アカウント アカウント 構成 構成コア コア デモ_データ デモ_データ docker docker サービス サービス テンプレート テンプレート .dockerignore .dockerignore .env.docker.example .env.docker.example .env.example .env.example .gitignore .gitignore .python-version .python-version CONTEXT.md CONTEXT.md Dockerfile Dockerfile LICENSE LICENSE Procfile Procfile README.md README.md Reports.png Reports.png Settings.png Settings.png docker-compose.yml docker-compose.yml manage.py manage.py 要件-db-drivers.txt 要件-db-drivers.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Safe Vibes は、AI が生成した HTML レポートを作成および共有するための Django アプリです
すべての販売業務や財務実験を管理対象外のシャドー IT に変える必要はありません。
ユーザーは AI レポート ビルダーとチャットし、承認されたデータベースに接続し、SQL を生成します
および HTML を使用して、結果をプレビューし、適切なユーザーにレポートを公開します。管理者が取得する
ガバナンス層: データベース接続、モデル制御、クエリ制限、キャッシュ
設定、SSO 構成、IP ホワイトリスト、および外部 URL ポリシー。
ビジネスチームはすでに AI を使用してレポートをバイブコーディングしています。 SQL を貼り付けます
ツールを使用したり、HTML ファイルを渡したりして、誤ってコスト、セキュリティ、
ガバナンスの問題。
Safe Vibes は、エンジニアリングを提供しながら、チームに安全に構築できる場所を提供します。
そしてITはコントロールプレーンです。
HTML + SQL レポート用のチャット スタイル AI ビルダー
Postgres、SQLite デモ データ、BigQuery、Snowflake の管理されたデータベース接続
読み取り専用 SQL ポリシー チェック、行/バイト制限、クエリ タイムアウト、およびキャッシュされたレポート データ
公開されたレポートの所有者、組織全体、または特定のユーザーによる共有
管理者が管理する AI プロバイダー、API キー、許可されたモデル、

およびデフォルトのモデル
SSO、レポートポリシー、セキュリティポリシー、ユーザー、データベースの組織設定
レポートアクセス用の IP ホワイトリスト
CSP およびランタイム ガードを備えた外部レポート URL ホワイトリスト/ブラックリスト ルール
初回実行テスト用にバンドルされた偽の SaaS 販売デモ ウェアハウス
Docker Compose デプロイメント パス (オプションの Heroku ノート付き)
Python - m pip install - rrequirements.txt
python manage.py 移行
python manage.py ensure_demo_database
Python manage.py createsuperuser
python manage.py runserver 127.0 。 0.1 : 8000
開く:
http://127.0.0.1:8000/
次に、/admin/ を開いて以下を作成します。
ユーザーを組織に結び付けるメンバーシップ
アプリはユーザー ID として電子メールを使用します。 Django の内部ユーザー名フィールドは次のとおりです。
電子メールから自動的に入力されます。
ローカルの非 Docker 開発用に .env.example を .env にコピーします。
DJANGO_DEBUG = false
DJANGO_SECRET_KEY = 私を置き換えます
SECRET_ENCRYPTION_KEY = fernet キーで置き換える
DJANGO_ALLOWED_HOSTS = あなたのドメイン.example.com
DATABASE_URL = postgres://ユーザー:パスワード@ホスト:5432/dbname
DATABASE_CONN_MAX_AGE = 0
REPORT_CACHE_ENABLED = true
ENABLE_DEMO_DATABASE_CONNECTION = true
SECRET_ENCRYPTION_KEY を生成します。
python - c " from cryptography.fernet import Fernet; print(Fernet.generate_key().decode()) "
DATABASE_CONN_MAX_AGE=0 は、ASGI デプロイメント向けに意図的に設定されています。ジャンゴ
ASGI では永続的なデータベース接続は推奨されません。それらを維持する
無効にすると、断続的に発生する古い Postgres 接続が回避されます。
Docker Compose は、推奨される汎用デプロイメント エントリポイントです。付属の
compose ファイルはアプリと Postgres を実行し、同じ ASGI/Gunicorn/Uvicorn を使用します
本番環境としてのサーバー形状。
docker 構成 -- ビルド
開く:
http://127.0.0.1:8000/
実行中のコンテナにスーパーユーザーを作成します。
docker compose exec web python manage.py createsuperuser
デフォルトの構成データベース

は:
DATABASE_URL = postgres://save_vibes:save_vibes@db:5432/save_vibes
実際の Docker デプロイメントの場合:
db サービスをマネージド Postgres に置き換えるか、単純な目的のためにそのままにしておきます。
単一ホストの展開。
バンドルされたデータベースを交換するときに、DATABASE_URL を管理対象データベースに設定します。
Postgres サービス。
実際の DJANGO_SECRET_KEY を設定します。
実際の SECRET_ENCRYPTION_KEY を設定します。
DJANGO_ALLOWED_HOSTS をドメインに設定します。
意図的に ASGI から変更しない限り、DATABASE_CONN_MAX_AGE=0 を維持します。
単純なデプロイの場合は RUN_MIGRATIONS=true を維持するか、移行を
プラットフォームがリリース ジョブをサポートしている場合は、別のリリース ジョブを実行します。
Python manage.py 移行 --noinput
python manage.py ensure_demo_database
次のいずれかを使用して無効にします。
RUN_MIGRATIONS = false
ENABLE_DEMO_DATABASE_CONNECTION = false
Docker イメージを直接デプロイするプラットフォームの場合は、Docker イメージで構築されたのと同じイメージを使用します。
Dockerfile を作成し、プラットフォームの経由でこれらの環境変数を提供します。
シークレット/構成システム。コンテナは PORT でリッスンし、デフォルトは 8000 です。
Heroku はオプションです。アプリには次のものが用意されています。
Gunicorn + Uvicorn ASGI ワーカーを使用した Procfile Web プロセス
移行およびデモ データベースのセットアップのための Procfile リリース プロセス
dj-database-url による DATABASE_URL のサポート
WhiteNoise 静的ファイルの提供
Heroku ログイン
heroku あなたの - アプリ - 名を作成します
heroku addons:create heroku - postgresql:essential - 0 -- app your - app - name
heroku config:set DJANGO_DEBUG = false -- アプリのアプリ名
heroku config:set DJANGO_SECRET_KEY = " replace-with-a-long-random-secret " -- アプリのアプリ名
heroku config:set SECRET_ENCRYPTION_KEY = " replace-with-a-fernet-key " -- アプリ - アプリ名 -
heroku config:set DJANGO_ALLOWED_HOSTS = " your-app-name.herokuapp.com " -- アプリのアプリ名
heroku config:set ENABLE_DEMO_DATABASE_CONNECTION = true -- app your - app - n

雨
heroku config:set DATABASE_CONN_MAX_AGE = 0 -- アプリ - アプリ名 -
heroku config:set DJANGO_LOG_LEVEL = INFO -- アプリ - アプリ名 -
heroku config:set WEB_CONCURRENCY = 2 -- アプリ - アプリ名 -
heroku config:set WEB_TIMEOUT = 180 -- アプリ - アプリ名 -
GitHub または CLI からデプロイします。
git Push Heroku メイン
ブートストラップ管理者を作成します。
heroku run python manage.py createsuperuser -- app your - app - name
次に、 /admin/ に組織とメンバーシップを作成します。
リポジトリには、次の場所にある小さな読み取り専用 SQLite デモ ウェアハウスが含まれています。
デモデータ/デモ_セールス.sqlite3
新しい組織は、次の場合に Demo SaaS Sales データベース接続を自動的に取得します。
ENABLE_DEMO_DATABASE_CONNECTION=true 。
コミットされたデモ ファイルを更新します。
Python サービス / build_demo_sqlite.py
バックフィルのデモ接続:
python manage.py ensure_demo_database
データベース接続
サポートされているレポート データ ソース:
SQLAlchemy 非同期ドライバー経由の Postgres
非同期ポーリングを使用した REST jobs.query 経由の BigQuery
Snowflake SQL API 経由の Snowflake
[設定] > [データベース接続] で、会社管理者は承認済みの接続を作成できます
接続。シークレットは SECRET_ENCRYPTION_KEY で暗号化され、
編集されたプレビューが表示されます。
[設定] > [AI プロバイダー] で、管理者はプロバイダー キーを追加し、許可するプロバイダーを選択できます。
モデル。ビルダーは、ドラフトに
有効な明示的なオーバーライド。
[設定] > [SSO] では、ハンドホールディング OIDC 構成ページが提供されます。
組織の設定されたログイン URL を通じて SSO を完了するユーザーは、
必要に応じて自動的に作成され、閲覧者としてその組織に追加されます。
会社管理者は後でそれらを宣伝できます。 「SSO が必要」トグルはパスワードをブロックします
その組織のスタッフ以外のユーザーはログインします。スタッフとスーパーユーザーは引き続き
ブートストラップ/管理者アクセスにはパスワード ログインを使用します。
Safe Vibes にはいくつかのガードレールが含まれています。
クゥ

ery タイムアウト、行数、生バイト、圧縮キャッシュ制限
TTL を使用したオプションのレポート データ キャッシュ
暗号化された AI、データベース、SSO、およびウェアハウスの認証情報
組織レベルのレポート共有制御
レポートアクセス用の IP ホワイトリスト
外部レポート URL ホワイトリスト/ブラックリスト ルール
レポート プレビューにおける CSP およびランタイム フェッチ / XMLHttpRequest ガード
これらは実用的な MVP コントロールであり、正式なセキュリティ証明ではありません。コードを確認してください
機密性の高い運用データを公開する前にポリシーを適用します。
python manage.py テスト core.tests
Python manage.pyチェック
使い捨てのローカル Postgres デモ ウェアハウスを実行します。
docker compose -f services / docker - compose.postgres.yml up - d
Pythonサービス/seed_demo_postgres.py
ライセンス
Safe Vibes は、Apache License 2.0 に基づいてリリースされています。 「ライセンス」を参照してください。
AI が生成したレポートを安全な方法で作成および展開するための Django ベースのアプリ。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A django-based app for creating and deploying AI generated reports in a safe manner. - tanner49/safe_vibes

GitHub - tanner49/safe_vibes: A django-based app for creating and deploying AI generated reports in a safe manner. · GitHub
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
tanner49
/
safe_vibes
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits accounts accounts config config core core demo_data demo_data docker docker services services templates templates .dockerignore .dockerignore .env.docker.example .env.docker.example .env.example .env.example .gitignore .gitignore .python-version .python-version CONTEXT.md CONTEXT.md Dockerfile Dockerfile LICENSE LICENSE Procfile Procfile README.md README.md Reports.png Reports.png Settings.png Settings.png docker-compose.yml docker-compose.yml manage.py manage.py requirements-db-drivers.txt requirements-db-drivers.txt requirements.txt requirements.txt View all files Repository files navigation
Safe Vibes is a Django app for building and sharing AI-generated HTML reports
without turning every sales ops or finance experiment into unmanaged shadow IT.
Users chat with an AI report builder, connect approved databases, generate SQL
and HTML, preview the result, and publish reports to the right people. Admins get
the governance layer: database connections, model controls, query limits, cache
settings, SSO configuration, IP allowlists, and external URL policies.
Business teams are already vibe-coding reports with AI. They paste SQL into
tools, pass around HTML files, and accidentally create cost, security, and
governance problems.
Safe Vibes gives those teams a safer place to build while giving engineering
and IT a control plane.
Chat-style AI builder for HTML + SQL reports
Governed database connections for Postgres, SQLite demo data, BigQuery, and Snowflake
Read-only SQL policy checks, row/byte limits, query timeouts, and cached report data
Published report sharing by owner, whole organization, or specific users
Admin-managed AI providers, API keys, allowed models, and default models
Organization settings for SSO, report policy, security policy, users, and databases
IP allowlists for report access
External report URL whitelist/blacklist rules with CSP and runtime guards
Bundled fake SaaS sales demo warehouse for first-run testing
Docker Compose deployment path, with optional Heroku notes
python - m pip install - r requirements.txt
python manage.py migrate
python manage.py ensure_demo_database
python manage.py createsuperuser
python manage.py runserver 127.0 . 0.1 : 8000
Open:
http://127.0.0.1:8000/
Then open /admin/ and create:
A Membership connecting your user to the organization
The app uses email as the user identity. Django's internal username field is
filled from email automatically.
Copy .env.example to .env for local non-Docker development.
DJANGO_DEBUG = false
DJANGO_SECRET_KEY = replace-me
SECRET_ENCRYPTION_KEY = replace-with-fernet-key
DJANGO_ALLOWED_HOSTS = your-domain.example.com
DATABASE_URL = postgres://user:password@host:5432/dbname
DATABASE_CONN_MAX_AGE = 0
REPORT_CACHE_ENABLED = true
ENABLE_DEMO_DATABASE_CONNECTION = true
Generate SECRET_ENCRYPTION_KEY :
python - c " from cryptography.fernet import Fernet; print(Fernet.generate_key().decode()) "
DATABASE_CONN_MAX_AGE=0 is intentional for the ASGI deployment. Django
persistent database connections are not recommended under ASGI; keeping them
disabled avoids intermittent stale Postgres connections.
Docker Compose is the recommended generic deployment entrypoint. The included
compose file runs the app plus Postgres and uses the same ASGI/Gunicorn/Uvicorn
server shape as production.
docker compose up -- build
Open:
http://127.0.0.1:8000/
Create a superuser in the running container:
docker compose exec web python manage.py createsuperuser
The default compose database is:
DATABASE_URL = postgres://save_vibes:save_vibes@db:5432/save_vibes
For a real Docker deployment:
Replace the db service with your managed Postgres, or keep it for a simple
single-host deployment.
Set DATABASE_URL to your managed database when you swap out the bundled
Postgres service.
Set a real DJANGO_SECRET_KEY .
Set a real SECRET_ENCRYPTION_KEY .
Set DJANGO_ALLOWED_HOSTS to your domain.
Keep DATABASE_CONN_MAX_AGE=0 unless you deliberately change away from ASGI.
Keep RUN_MIGRATIONS=true for simple deployments, or run migrations as a
separate release job if your platform supports one.
python manage.py migrate --noinput
python manage.py ensure_demo_database
Disable either with:
RUN_MIGRATIONS = false
ENABLE_DEMO_DATABASE_CONNECTION = false
For platforms that deploy Docker images directly, use the same image built by
the Dockerfile and provide these environment variables through the platform's
secret/config system. The container listens on PORT and defaults to 8000 .
Heroku is optional. The app is prepared for it with:
Procfile web process using Gunicorn + Uvicorn ASGI workers
Procfile release process for migrations and demo database setup
DATABASE_URL support through dj-database-url
WhiteNoise static file serving
heroku login
heroku create your - app - name
heroku addons:create heroku - postgresql:essential - 0 -- app your - app - name
heroku config:set DJANGO_DEBUG = false -- app your - app - name
heroku config:set DJANGO_SECRET_KEY = " replace-with-a-long-random-secret " -- app your - app - name
heroku config:set SECRET_ENCRYPTION_KEY = " replace-with-a-fernet-key " -- app your - app - name
heroku config:set DJANGO_ALLOWED_HOSTS = " your-app-name.herokuapp.com " -- app your - app - name
heroku config:set ENABLE_DEMO_DATABASE_CONNECTION = true -- app your - app - name
heroku config:set DATABASE_CONN_MAX_AGE = 0 -- app your - app - name
heroku config:set DJANGO_LOG_LEVEL = INFO -- app your - app - name
heroku config:set WEB_CONCURRENCY = 2 -- app your - app - name
heroku config:set WEB_TIMEOUT = 180 -- app your - app - name
Deploy from GitHub or the CLI:
git push heroku main
Create a bootstrap admin:
heroku run python manage.py createsuperuser -- app your - app - name
Then create an organization and membership in /admin/ .
The repo includes a small read-only SQLite demo warehouse at:
demo_data/demo_sales.sqlite3
New organizations automatically get a Demo SaaS Sales database connection when
ENABLE_DEMO_DATABASE_CONNECTION=true .
Refresh the committed demo file:
python services / build_demo_sqlite.py
Backfill demo connections:
python manage.py ensure_demo_database
Database Connections
Supported report data sources:
Postgres via SQLAlchemy async drivers
BigQuery via REST jobs.query with async polling
Snowflake via the Snowflake SQL API
In Settings > Database connections, company admins can create approved
connections. Secrets are encrypted with SECRET_ENCRYPTION_KEY and only
redacted previews are displayed.
In Settings > AI providers, admins can add provider keys and choose allowed
models. The builder uses the admin-selected default model unless a draft has a
valid explicit override.
Settings > SSO provides a handholding OIDC configuration page:
Users who complete SSO through an organization's configured login URL are
created automatically if needed and added to that organization as viewers.
Company admins can promote them later. The Require SSO toggle blocks password
login for non-staff users in that organization. Staff and superusers can still
use password login for bootstrap/admin access.
Safe Vibes includes several guardrails:
Query timeout, row count, raw byte, and compressed cache limits
Optional report data caching with TTL
Encrypted AI, database, SSO, and warehouse credentials
Organization-level report sharing controls
IP allowlists for report access
External report URL whitelist/blacklist rules
CSP and runtime fetch / XMLHttpRequest guards in report previews
These are practical MVP controls, not a formal security proof. Review the code
and policies before exposing sensitive production data.
python manage.py test core.tests
python manage.py check
Run a disposable local Postgres demo warehouse:
docker compose -f services / docker - compose.postgres.yml up - d
python services / seed_demo_postgres.py
License
Safe Vibes is released under the Apache License 2.0. See LICENSE .
A django-based app for creating and deploying AI generated reports in a safe manner.
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

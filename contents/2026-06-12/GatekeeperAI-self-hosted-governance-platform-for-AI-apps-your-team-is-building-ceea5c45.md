---
source: "https://github.com/jacobthomasmichael/GatekeeperAI/blob/main/README.md"
hn_url: "https://news.ycombinator.com/item?id=48508622"
title: "GatekeeperAI – self-hosted governance platform for AI apps your team is building"
article_title: "GatekeeperAI/README.md at main · jacobthomasmichael/GatekeeperAI · GitHub"
author: "jacob_thomas503"
captured_at: "2026-06-12T21:38:17Z"
capture_tool: "hn-digest"
hn_id: 48508622
score: 1
comments: 0
posted_at: "2026-06-12T19:44:42Z"
tags:
  - hacker-news
  - translated
---

# GatekeeperAI – self-hosted governance platform for AI apps your team is building

- HN: [48508622](https://news.ycombinator.com/item?id=48508622)
- Source: [github.com](https://github.com/jacobthomasmichael/GatekeeperAI/blob/main/README.md)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T19:44:42Z

## Translation

タイトル: GatekeeperAI – チームが構築している AI アプリ用の自己ホスト型ガバナンス プラットフォーム
記事のタイトル: GatekeeperAI/README.md (メイン) · jacobthomasmichael/GatekeeperAI · GitHub
説明: 内部アクセスおよびワークフロー オーケストレーション アシスタント (リクエスト → ルーティング → Jira → 追跡) - メインの GatekeeperAI/README.md · jacobthomasmichael/GatekeeperAI

記事本文:
メインの GatekeeperAI/README.md · jacobthomasmichael/GatekeeperAI · GitHub
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
ジェイコブトーマスマイケル
/
ゲートキーパーAI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピー もっと責める

ファイル アクションを責める さらにファイル アクションを追加する 最新のコミット
履歴 履歴 78 行 (55 loc) · 3.36 KB メイン ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード 概要編集と raw アクション GatekeeperAI
GatekeeperAI は、企業チームがサードパーティおよび社内の AI アプリケーションを安全に導入できるようにするオンプレミス プラットフォームです。すべてのアプリは、ユーザーがアクセスできるようになる前に、自動セキュリティ スキャン、人間による承認、サンドボックス コンテナーのデプロイメントを経ます。
送信 — 開発者はアプリのコードを GatekeeperAI git サーバーにプッシュします。
スキャン — プラットフォームは、機密情報の検出、依存関係の脆弱性監査、出口ネットワーク分析、PII 露出チェック、および Claude AI による LLM を利用したコード レビューの 5 つのスキャナーを自動的に実行します。
レビュー — 指定された承認者がスキャン結果をレビューし、SLA 期限が自動的に追跡されながらアプリを承認または拒否します。
デプロイ - 承認されたアプリは Docker コンテナに組み込まれ、パブリック URL 経由でアクセスできる分離された環境で起動されます。
管理 — 実行時のシークレット (API キー、認証情報) は、デプロイ時に環境変数として挿入され、コードには保存されません。
自動化されたマルチスキャナ パイプライン - シークレット、CVE、出力ルール、PII、LLM コード レビューがプッシュごとに並行して実行されます。
リスク階層化 — アプリは自動的にスコアリングされ、レビューの緊急性を決定するリスク階層 (低/中/高/重大) が割り当てられます。
SLA の施行 - 期限を過ぎた承認にはフラグが付けられ、エスカレーターには電子メールで通知されます
暗号化されたシークレットの注入 - アプリごとのシークレットは保存時に AES-256 で暗号化され、コンテナーの起動時に注入されます。
監査ログ - すべてのアクション (承認、展開、シークレットの変更) がアクター、IP、およびタイムスタンプとともに記録されます。
管理パネル - ユーザー管理 (ロールの作成、無効化、変更)、監査ログ ビューア、プラットフォーム

幅広いメトリクス
セットアップ ウィザード - 初回実行ウィザードは、構成ファイルの編集を必要とせずにインスタンスを構成します。
デフォルトで安全 — リフレッシュ トークン ローテーションを備えた JWT、すべてのエンドポイントのレート制限、セキュリティ ヘッダー (CSP、HSTS など)、非ルート コンテナー
レイヤー
テクノロジー
バックエンドAPI
FastAPI + SQLAlchemy 2.0 非同期 + PostgreSQL 16
タスクキュー
セロリ+レディス
コンテナランタイム
Docker SDK (Python)
LLM
Anthropic Claude API
フロントエンド
Next.js 16 (アプリルーター) + Tailwind CSS
認証
Redis を利用した JTI ローテーションによる JWT (アクセス + リフレッシュ)
はじめに
ローカル インストール、AWS/Azure/GCP クラウド ホスティング、SSL を使用したカスタム ドメインなどの完全なセットアップ手順については、INSTALL.md を参照してください。
クイックスタート (Docker が必要):
cp .env.example .env
# .env に SECRET_KEY、SECRET_ENCRYPTION_KEY、ANTHROPIC_API_KEY を入力します
docker compose -f infra/docker-compose.yml up --build
次に、http://localhost:3000 を開き、セットアップ ウィザードに従います。
バックエンド/ FastAPI アプリケーション、スキャナー、Celery ワーカー、Alembic 移行
フロントエンド/Next.js Web アプリケーション
インフラ/Docker Compose 構成
ワーカー/ Celery タスク定義 (デプロイ、SLA チェック)
ユーザーの役割
役割
できる
ic (個人寄稿者)
アプリを送信し、自分のアプリを表示し、結果をスキャンします
承認者
IC が実行できるすべての機能に加え、保留中の承認を確認して決定し、すべての展開を表示します
管理者
承認者が実行できるすべての機能に加え、ユーザーの管理、展開の停止/開始、監査ログの表示
新しいユーザーは管理者によって作成されます。公開された自己登録はありません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Internal access & workflow orchestration assistant (requests → routing → Jira → tracking) - GatekeeperAI/README.md at main · jacobthomasmichael/GatekeeperAI

GatekeeperAI/README.md at main · jacobthomasmichael/GatekeeperAI · GitHub
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
jacobthomasmichael
/
GatekeeperAI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 78 lines (55 loc) · 3.36 KB main Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions GatekeeperAI
GatekeeperAI is an on-premises platform that lets enterprise teams safely adopt third-party and internal AI applications. Every app goes through automated security scanning, human approval, and sandboxed container deployment before any user can access it.
Submit — A developer pushes their app's code to the GatekeeperAI git server.
Scan — The platform automatically runs five scanners: secrets detection, dependency vulnerability audit, egress network analysis, PII exposure check, and an LLM-powered code review via Claude AI.
Review — A designated approver reviews the scan results and approves or rejects the app, with an SLA deadline tracked automatically.
Deploy — Approved apps are built into Docker containers and launched in an isolated environment, accessible via a public URL.
Manage — Runtime secrets (API keys, credentials) are injected as environment variables at deploy time, never stored in the code.
Automated multi-scanner pipeline — secrets, CVEs, egress rules, PII, and LLM code review run in parallel on every push
Risk tiering — apps are automatically scored and assigned a risk tier (low / medium / high / critical) that determines review urgency
SLA enforcement — overdue approvals are flagged and escalators are notified via email
Encrypted secret injection — per-app secrets are AES-256 encrypted at rest and injected at container startup
Audit log — every action (approval, deployment, secret change) is recorded with actor, IP, and timestamp
Admin panel — user management (create, disable, change roles), audit log viewer, platform-wide metrics
Setup wizard — first-run wizard configures the instance with no config-file editing required
Secure by default — JWT with refresh token rotation, rate limiting on all endpoints, security headers (CSP, HSTS, etc.), non-root containers
Layer
Technology
Backend API
FastAPI + SQLAlchemy 2.0 async + PostgreSQL 16
Task queue
Celery + Redis
Container runtime
Docker SDK (Python)
LLM
Anthropic Claude API
Frontend
Next.js 16 (App Router) + Tailwind CSS
Auth
JWT (access + refresh) with Redis-backed JTI rotation
Getting started
See INSTALL.md for full setup instructions, including local installation, AWS/Azure/GCP cloud hosting, and custom domain with SSL.
Quick start (requires Docker):
cp .env.example .env
# Fill in SECRET_KEY, SECRET_ENCRYPTION_KEY, and ANTHROPIC_API_KEY in .env
docker compose -f infra/docker-compose.yml up --build
Then open http://localhost:3000 and follow the setup wizard.
backend/ FastAPI application, scanners, Celery workers, Alembic migrations
frontend/ Next.js web application
infra/ Docker Compose configuration
worker/ Celery task definitions (deploy, SLA checks)
User roles
Role
Can do
ic (individual contributor)
Submit apps, view own apps and scan results
approver
Everything an IC can do, plus review and decide on pending approvals, view all deployments
admin
Everything an approver can do, plus manage users, stop/start deployments, view audit logs
New users are created by an admin — there is no public self-registration.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

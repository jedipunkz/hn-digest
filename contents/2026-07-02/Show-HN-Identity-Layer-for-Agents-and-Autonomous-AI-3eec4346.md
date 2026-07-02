---
source: "https://github.com/authsec-ai/authsec-ai"
hn_url: "https://news.ycombinator.com/item?id=48754628"
title: "Show HN: Identity Layer for Agents and Autonomous AI"
article_title: "GitHub - authsec-ai/authsec-ai: Identity layer for Agents and Autonomous AI. With OAuth 2.1, CIBA, SPIFFE workload identity (for agent to agent auth), secret vault and MFA. · GitHub"
author: "azifali"
captured_at: "2026-07-02T00:31:33Z"
capture_tool: "hn-digest"
hn_id: 48754628
score: 1
comments: 0
posted_at: "2026-07-01T23:47:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Identity Layer for Agents and Autonomous AI

- HN: [48754628](https://news.ycombinator.com/item?id=48754628)
- Source: [github.com](https://github.com/authsec-ai/authsec-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T23:47:36Z

## Translation

タイトル: HN を表示: エージェントと自律型 AI のアイデンティティ レイヤー
記事のタイトル: GitHub - authsec-ai/authsec-ai: エージェントと自律型 AI の ID レイヤー。 OAuth 2.1、CIBA、SPIFFE ワークロード ID (エージェント間の認証用)、シークレット ボールト、および MFA を使用します。 · GitHub
説明: エージェントと自律型 AI の ID レイヤー。 OAuth 2.1、CIBA、SPIFFE ワークロード ID (エージェント間の認証用)、シークレット ボールト、および MFA を使用します。 - authsec-ai/authsec-ai

記事本文:
GitHub - authsec-ai/authsec-ai: エージェントと自律型 AI の ID レイヤー。 OAuth 2.1、CIBA、SPIFFE ワークロード ID (エージェント間の認証用)、シークレット ボールト、および MFA を使用します。 · GitHub
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
読み込み中にエラーが発生しました。これをリロードしてください

のページ。
authsec-ai
/
authsec-ai
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
98 コミット 98 コミット .github .github cmd cmd config config コントローラー コントローラー データベース データベース ハンドラー ハンドラー 内部 内部ミドルウェア ミドルウェア マイグレーション マイグレーション モデル モデル 監視 監視 リポジトリ リポジトリ ルート ルート スクリプト スクリプト サービス サービス テスト テスト utils utils .env .env .env.example .env.example .golangci.yml .golangci.yml API_DOCS.md API_DOCS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Jenkinsfile Jenkinsfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示リポジトリ ファイルのナビゲーション
AuthSec – ID およびアクセス管理プラットフォーム
AuthSec は、完全な ID ライフサイクル (認証、MFA、OIDC フェデレーション、RBAC、SCIM プロビジョニング、クライアント管理、外部サービス資格情報、SPIFFE/SPIRE ワークロード ID) をすべて単一のバイナリから提供する統合 Go サービスです。
APIルートマップ
コア認証とユーザー フロー ( /authsec/uflow )
WebAuthn / パスキー ( /authsec/webauthn )
クライアント管理 ( /authsec/clientms )
Hydra マネージャー ( /authsec/hmgr )
OIDC 構成マネージャー ( /authsec/oocmgr )
認証マネージャー ( /authsec/authmgr )
外部サービス ( /authsec/exsvc )
SPIRE ヘッドレス ( /authsec/spire )
移行管理 ( /authsec/migration )
┌───────────────────────────┐
│ authsec (ポート 7468) │
│ │
│ /authsec/uflow/* – オート

h、RBAC、OIDC フェデレーション、SCIM │
│ /authsec/webauthn/* – パスキー、TOTP、SMS MFA │
│ /authsec/clientms/* – クライアントのライフサイクル管理 │
│ /authsec/hmgr/* – Ory Hydra ログイン/同意、SAML SSO │
│ /authsec/oocmgr/* – OIDC プロバイダー設定と Hydra 同期 │
│ /authsec/authmgr/* – JWT 検証、RBAC チェック │
│ /authsec/exsvc/* – 外部サービスレジストリ │
│ /authsec/spire/* – SPIFFE ワークロード ID │
│ /authsec/migration/* – データベース移行管理 │
│ │
│ /.well-known/* – OIDC 検出 (RFC 8414 ルート パス) │
│ /metrics – Prometheus メトリクス │
━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│
§── PostgreSQL (マスター DB — すべての操作に対応する単一データベース)
§── mt-plugin (オプションの gRPC — マルチテナント DB 管理)
§── HashiCorp Vault (オプション - シークレット、OIDC プロバイダーの資格情報)
└── Redis (オプション — 権限キャッシュ、セッション キャッシュ)
すべての HTTP ルートは、単一の gin.Engine から提供されます。各モジュールのルートは独自のサブプレフィックスの下に存在するため、パスはグローバルに一意になります。
authsec は、デフォルトではシングルテナント サービスです。すべての操作ではマスター PostgreSQL データベースを使用します。マルチテナントのサポートには、 mt-plugin gRPC マイクロサービスを実行し、 MT_PLUGIN_GRPC_ADDR を設定する必要があります。
モジュール
サブプレフィックス
説明
コア認証とユーザー フロー
/authsec/uflow
管理者/エンドユーザーのログイン、RBAC、OIDC フェデレーション、SCIM、TOTP、CIBA、音声認証
WebAuthn / パスキー
/authsec/webauthn
WebAuthn/FIDO2 パスキー、TOTP セットアップ、SMS MFA
クライアント管理
/authsec/clientms
Hydra クライアントのライフサイクル管理
ヒドラマネージャー
/authsec/hmgr
Ory Hydra ログイン/同意、SA

ML SSO、トークン交換
OIDC 構成マネージャー
/authsec/oocmgr
OIDC プロバイダー構成、Hydra クライアント同期、SAML プロバイダー
認証マネージャー
/authsec/authmgr
JWT検証/発行、RBAC権限チェック、グループ管理
外部サービス
/authsec/exsvc
Vault に基づく認証情報を使用した外部サービス レジストリ
SPIRE ヘッドレス
/authsec/spire
SPIFFE/SPIRE ワークロード ID、OIDC トークン交換、クラウド フェデレーション (AWS/Azure/GCP)、RBAC/ABAC ポリシー エンジン
移行管理
/authsec/移行
データベース移行管理（マスターDB＋テナント別DB）
クイックスタート
HashiCorp Vault (オプション - OIDC シークレットに推奨)
mt-plugin (オプション — マルチテナント モードの場合にのみ必要)
# 環境変数をコピーして編集する
cp .env.example .env
# ビルド
go build -o authsec ./cmd/
# 実行
./authsec
または go run を使用します:
実行してください。/cmd/
サーバーはデフォルトでポート 7468 で起動します。
変数
説明
例
DB_NAME
PostgreSQL データベース名
authsec_db
DB_USER
データベースのユーザー名
認証秒
DB_パスワード
データベースのパスワード
チェンジミー
DB_ホスト
データベースホスト
ローカルホスト
DB_PORT
データベースポート
5432
WEBAUTHN_RP_NAME
WebAuthn 証明書利用者の表示名
認証秒
WEBAUTHN_RP_ID
WebAuthn 証明書利用者 ID (オリジンのホスト名と一致する必要があります)
app.authsec.dev
WEBAUTHN_ORIGIN
許可された WebAuthn 起点
https://app.authsec.dev
オプション – コアサービス
変数
デフォルト
説明
港
7468
HTTPリッスンポート
GIN_MODE
デバッグ
ジン実行モード (デバッグ/リリース/テスト)
環境
開発
テナント ドメイン チェックで使用されるランタイム ラベル (開発/運用)
DB_SCHEMA
公共の
PostgreSQL スキーマ
JWT_SECRET
「」
プライマリ JWT 署名シークレット (外部サービス ルート、SPIFFE デリゲート)
JWT_DEF_SECRET
—
デフォルトの JWT 署名シークレット (管理者/プラットフォーム トークン)
JWT_SDK_SECRET
—
SDK JWT 署名シークレット
BASE_URL
https://app.authsec.dev
OIDC コールバックおよび電子メールのベース URL l

インク
テナントドメイン_サフィックス
—
自動生成されたテナントのサブドメインのサフィックス
REDIS_URL
「」
Redis 接続 URL (例: redis://localhost:6379 )
ICP_SERVICE_URL
http://ローカルホスト:7001
ICP/PKIプロビジョニングサービス
REQUIRE_SERVER_AUTH
本当の
サービス間認証チェックを強制する (開発環境で無効にする場合は false)
SKIP_MIGRATIONS
偽
起動時にマスター DB の移行をスキップするには true に設定します。
MT_PLUGIN_GRPC_ADDR
「」
mt-plugin gRPC アドレス (例: localhost:7469 );シングルテナントモードの場合は空のままにしておきます
オプション - CORS
変数
デフォルト
説明
CORS_ALLOWED_ORIGINS
( WEBAUTHN_ORIGIN から自動検出)
カンマ区切りの許可されるオリジン
CORS_ALLOWED_METHODS
取得、投稿、挿入、パッチ、削除、オプション
許可されるHTTPメソッド
CORS_ALLOWED_HEADERS
オリジン、コンテンツタイプ、認可、…
許可されるリクエストヘッダー
オプション - 暗号化キー
変数
説明
TOTP_ENCRYPTION_KEY
保存時の TOTP シークレットを暗号化するための 64 進数文字の AES-256 キー (本番環境で必要)
SYNC_CONFIG_ENCRYPTION_KEY
保存時の AD/Entra 同期構成を暗号化するための 64 進数文字の AES-256 キー
オプション – Twilio (SMS MFA / 音声)
変数
説明
TWILIO_ACCOUNT_SID
Twilio アカウント SID (例: ACxxxxxxxx )
TWILIO_AUTH_TOKEN
Twilio 認証トークン
TWILIO_FROM_NUMBER
SMS OTP の送信者の電話番号 (例: +10000000000 )
オプション - 外部統合
変数
説明
VAULT_ADDR
HashiCorp Vault アドレス (デフォルト: http://localhost:8200 )
VAULT_TOKEN
ボールトのルート/サービス トークン
HYDRA_ADMIN_URL
Ory Hydra 管理 API (デフォルト: http://localhost:4445 )
HYDRA_PUBLIC_URL
Ory Hydra パブリック API (デフォルト: http://localhost:4444 )
REACT_APP_URL
リダイレクト用のフロントエンド アプリの URL
IDENTITY_PROVIDER_URL
ID プロバイダーのベース URL
SMTP_HOST / SMTP_PORT / SMTP_USER / SMTP_PASSWORD
電子メール通知用の SMTP
GOOGLE_CLIENT_SECRET
Google OIDC クライアント シークレット (Vault が使用できない場合のフォールバック)

ブル)
GITHUB_CLIENT_SECRET
GitHub OIDC クライアント シークレット
MICROSOFT_CLIENT_SECRET
Microsoft OIDC クライアント シークレット
HUBSPOT_ACCESS_TOKEN
HubSpot CRM 統合トークン
オプション - OIDC トークンの検証
変数
説明
AUTH_EXPECT_ISS
受信 OIDC トークンを検証するときに予期される iss クレーム (空 = スキップ)
AUTH_EXPECT_AUD
受信した OIDC トークンを検証するときに予期される aud クレーム (空 = スキップ)
オプション – SPIFFE / SVID OIDC
SPIFFE ワークロード ID/デリゲート エンドポイントが使用される場合にのみ必要です。
Okta が CIBA プロバイダーとして使用される場合にのみ必要です。
すべてのアプリケーション ルートは /authsec プレフィックスの下にあります (OIDC Discovery と /metrics を除く)。
コア認証とユーザー フロー ( /authsec/uflow )
健康
方法
パス
説明
ゲット
/authsec/uflow/health
人間ドック
ゲット
/authsec/uflow/health/tenant/:tenant_id
シングルテナント DB の健全性
ゲット
/authsec/uflow/health/tenants
すべてのテナント DB の健全性
管理者認証 ( /authsec/uflow/auth/admin )
方法
パス
説明
ゲット
/authsec/uflow/auth/admin/チャレンジ
認証チャレンジを取得する
投稿
/authsec/uflow/auth/admin/login/precheck
ログイン前チェック
投稿
/authsec/uflow/auth/admin/login/bootstrap
ブートストラップの最初の管理者
投稿
/authsec/uflow/auth/admin/login
管理者ログイン
投稿
/authsec/uflow/auth/admin/ログインハイブリッド
ハイブリッドログイン
投稿
/authsec/uflow/auth/admin/register
管理者を登録する
投稿
/authsec/uflow/auth/admin/complete-registration
登録を完了する
投稿
/authsec/uflow/auth/admin/パスワードを忘れた場合
パスワードのリセットを開始する
投稿
/authsec/uflow/auth/admin/パスワードを忘れた場合/verify-otp
OTP を確認する
投稿
/authsec/uflow/auth/admin/パスワードを忘れた場合/リセット
パスワードをリセットする
エンドユーザー認証 ( /authsec/uflow/auth/enduser )
方法
パス
説明
ゲット
/authsec/uflow/auth/エンドユーザー/チャレンジ
チャレンジを受ける
投稿
/authsec/uflow/auth/enduser/initiate-registration
レジストラを開始する

ション
投稿
/authsec/uflow/auth/enduser/verify-otp
OTP を確認して登録を完了する
投稿
/authsec/uflow/auth/enduser/login/precheck
ログイン前チェック
投稿
/authsec/uflow/auth/enduser/webauthn-callback
WebAuthn アサーション コールバック
投稿
/authsec/uflow/auth/enduser/delegate-svid
デリゲート SPIFFE SVID
デバイス認証付与 – RFC 8628 ( /authsec/uflow/auth/device )
方法
パス
認証
説明
投稿
/authsec/uflow/auth/デバイス/コード
公共
デバイスリクエストコード
投稿
/authsec/uflow/auth/デバイス/トークン
公共
デバイスがトークンをポーリングする
ゲット
/authsec/uflow/auth/device/activate/info
公共
UIのデバイス情報を取得する
投稿
/authsec/uflow/auth/device/verify
JWT
ユーザーがデバイスを認証する
ゲット
/authsec/uflow/activate
公共
アクティベーション UI ページ
音声認証 ( /authsec/uflow/auth/voice )
方法
パス
認証
説明
投稿
/authsec/uflow/auth/voice/initiate
公共
音声認証を開始する
投稿
/認証
[切り捨てられた]
認証されたエンドポイント (JWT + テナントが必要):
すべての管理エンドポイントには、 JWT + admin:access + tenant validation が必要です。
方法
パス
説明
投稿
/authsec/uflow/admin/roles
ロールの作成
ゲット
/authsec/uflow/admin/roles
役割の一覧表示
ゲット
/authsec/uflow/admin/roles/:role_id
IDでロールを取得
置く
/authsec/uflow/admin/roles/:role_id
役割を更新する
削除
/authsec/uflow/admin/roles/:role_id
役割の削除
投稿
/authsec/uflow/admin/bindings
役割の割り当て (範囲指定)
ゲット
/authsec/uflow/admin/bindings
リストバインディング
投稿
/authsec/uflow/admin/権限
許可を登録する
ゲット
/authsec/uflow/admin/権限
リスト権限
削除
/authsec/uflow/admin/permissions/:id
削除権限
ゲット
/authsec/uflow/admin/権限/リソース
リソースのリストを表示する
GET/POST/PUT/DEL

[切り捨てられた]

## Original Extract

Identity layer for Agents and Autonomous AI. With OAuth 2.1, CIBA, SPIFFE workload identity (for agent to agent auth), secret vault and MFA. - authsec-ai/authsec-ai

GitHub - authsec-ai/authsec-ai: Identity layer for Agents and Autonomous AI. With OAuth 2.1, CIBA, SPIFFE workload identity (for agent to agent auth), secret vault and MFA. · GitHub
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
authsec-ai
/
authsec-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
98 Commits 98 Commits .github .github cmd cmd config config controllers controllers database database handlers handlers internal internal middlewares middlewares migrations migrations models models monitoring monitoring repository repository routes routes scripts scripts services services tests tests utils utils .env .env .env.example .env.example .golangci.yml .golangci.yml API_DOCS.md API_DOCS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Jenkinsfile Jenkinsfile LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
AuthSec – Identity & Access Management Platform
AuthSec is a unified Go service for the complete identity lifecycle: authentication, MFA, OIDC federation, RBAC, SCIM provisioning, client management, external-service credentials, and SPIFFE/SPIRE workload identity — all served from a single binary.
API Route Map
Core Auth & User Flow ( /authsec/uflow )
WebAuthn / Passkeys ( /authsec/webauthn )
Client Management ( /authsec/clientms )
Hydra Manager ( /authsec/hmgr )
OIDC Config Manager ( /authsec/oocmgr )
Auth Manager ( /authsec/authmgr )
External Services ( /authsec/exsvc )
SPIRE Headless ( /authsec/spire )
Migration Management ( /authsec/migration )
┌────────────────────────────────────────────────────────────────────┐
│ authsec (port 7468) │
│ │
│ /authsec/uflow/* – Auth, RBAC, OIDC federation, SCIM │
│ /authsec/webauthn/* – Passkeys, TOTP, SMS MFA │
│ /authsec/clientms/* – Client lifecycle management │
│ /authsec/hmgr/* – Ory Hydra login/consent, SAML SSO │
│ /authsec/oocmgr/* – OIDC provider config & Hydra sync │
│ /authsec/authmgr/* – JWT verification, RBAC checks │
│ /authsec/exsvc/* – External service registry │
│ /authsec/spire/* – SPIFFE workload identity │
│ /authsec/migration/* – Database migration management │
│ │
│ /.well-known/* – OIDC discovery (RFC 8414 root path) │
│ /metrics – Prometheus metrics │
└────────────────────────────────────────────────────────────────────┘
│
├── PostgreSQL (master DB — single database for all operations)
├── mt-plugin (optional gRPC — multi-tenant DB management)
├── HashiCorp Vault (optional — secrets, OIDC provider credentials)
└── Redis (optional — permission cache, session cache)
All HTTP routes are served from a single gin.Engine . Each module's routes live under its own sub-prefix so paths are globally unique.
authsec is a single-tenant service by default . All operations use the master PostgreSQL database. Multi-tenant support requires running the mt-plugin gRPC microservice and setting MT_PLUGIN_GRPC_ADDR .
Module
Sub-prefix
Description
Core Auth & User Flow
/authsec/uflow
Admin/enduser login, RBAC, OIDC federation, SCIM, TOTP, CIBA, voice auth
WebAuthn / Passkeys
/authsec/webauthn
WebAuthn/FIDO2 passkeys, TOTP setup, SMS MFA
Client Management
/authsec/clientms
Hydra client lifecycle management
Hydra Manager
/authsec/hmgr
Ory Hydra login/consent, SAML SSO, token exchange
OIDC Config Manager
/authsec/oocmgr
OIDC provider config, Hydra client sync, SAML providers
Auth Manager
/authsec/authmgr
JWT verify/issue, RBAC permission checks, group management
External Services
/authsec/exsvc
External service registry with Vault-backed credentials
SPIRE Headless
/authsec/spire
SPIFFE/SPIRE workload identity, OIDC token exchange, cloud federation (AWS/Azure/GCP), RBAC/ABAC policy engine
Migration Management
/authsec/migration
Database migration management (master DB + per-tenant DB)
Quick Start
HashiCorp Vault (optional — recommended for OIDC secrets)
mt-plugin (optional — required only for multi-tenant mode)
# Copy and edit environment variables
cp .env.example .env
# Build
go build -o authsec ./cmd/
# Run
./authsec
Or with go run :
go run ./cmd/
The server starts on port 7468 by default.
Variable
Description
Example
DB_NAME
PostgreSQL database name
authsec_db
DB_USER
Database username
authsec
DB_PASSWORD
Database password
changeme
DB_HOST
Database host
localhost
DB_PORT
Database port
5432
WEBAUTHN_RP_NAME
WebAuthn relying party display name
AuthSec
WEBAUTHN_RP_ID
WebAuthn relying party ID (must match origin's hostname)
app.authsec.dev
WEBAUTHN_ORIGIN
Allowed WebAuthn origin
https://app.authsec.dev
Optional – Core Service
Variable
Default
Description
PORT
7468
HTTP listen port
GIN_MODE
debug
Gin run mode ( debug / release / test )
ENVIRONMENT
development
Runtime label used by tenant domain checks ( development / production )
DB_SCHEMA
public
PostgreSQL schema
JWT_SECRET
""
Primary JWT signing secret (ext-service routes, SPIFFE delegate)
JWT_DEF_SECRET
—
Default JWT signing secret (admin / platform tokens)
JWT_SDK_SECRET
—
SDK JWT signing secret
BASE_URL
https://app.authsec.dev
Base URL for OIDC callbacks and email links
TENANT_DOMAIN_SUFFIX
—
Suffix for auto-generated tenant sub-domains
REDIS_URL
""
Redis connection URL (e.g. redis://localhost:6379 )
ICP_SERVICE_URL
http://localhost:7001
ICP/PKI provisioning service
REQUIRE_SERVER_AUTH
true
Enforce inter-service auth check ( false to disable in dev)
SKIP_MIGRATIONS
false
Set to true to skip master DB migrations at startup
MT_PLUGIN_GRPC_ADDR
""
mt-plugin gRPC address (e.g. localhost:7469 ); leave empty for single-tenant mode
Optional – CORS
Variable
Default
Description
CORS_ALLOWED_ORIGINS
(auto-detect from WEBAUTHN_ORIGIN )
Comma-separated allowed origins
CORS_ALLOWED_METHODS
GET,POST,PUT,PATCH,DELETE,OPTIONS
Allowed HTTP methods
CORS_ALLOWED_HEADERS
Origin,Content-Type,Authorization,…
Allowed request headers
Optional – Encryption Keys
Variable
Description
TOTP_ENCRYPTION_KEY
64-hex-char AES-256 key for encrypting TOTP secrets at rest (required in production)
SYNC_CONFIG_ENCRYPTION_KEY
64-hex-char AES-256 key for encrypting AD/Entra sync configurations at rest
Optional – Twilio (SMS MFA / Voice)
Variable
Description
TWILIO_ACCOUNT_SID
Twilio account SID (e.g. ACxxxxxxxx )
TWILIO_AUTH_TOKEN
Twilio auth token
TWILIO_FROM_NUMBER
Sender phone number for SMS OTPs (e.g. +10000000000 )
Optional – External Integrations
Variable
Description
VAULT_ADDR
HashiCorp Vault address (default: http://localhost:8200 )
VAULT_TOKEN
Vault root/service token
HYDRA_ADMIN_URL
Ory Hydra admin API (default: http://localhost:4445 )
HYDRA_PUBLIC_URL
Ory Hydra public API (default: http://localhost:4444 )
REACT_APP_URL
Frontend app URL for redirects
IDENTITY_PROVIDER_URL
Identity provider base URL
SMTP_HOST / SMTP_PORT / SMTP_USER / SMTP_PASSWORD
SMTP for email notifications
GOOGLE_CLIENT_SECRET
Google OIDC client secret (fallback if Vault unavailable)
GITHUB_CLIENT_SECRET
GitHub OIDC client secret
MICROSOFT_CLIENT_SECRET
Microsoft OIDC client secret
HUBSPOT_ACCESS_TOKEN
HubSpot CRM integration token
Optional – OIDC Token Validation
Variable
Description
AUTH_EXPECT_ISS
Expected iss claim when validating incoming OIDC tokens (empty = skip)
AUTH_EXPECT_AUD
Expected aud claim when validating incoming OIDC tokens (empty = skip)
Optional – SPIFFE / SVID OIDC
Required only when SPIFFE workload identity / delegate endpoints are used.
Required only when Okta is used as a CIBA provider.
All application routes are under the /authsec prefix (except OIDC discovery and /metrics ).
Core Auth & User Flow ( /authsec/uflow )
Health
Method
Path
Description
GET
/authsec/uflow/health
Comprehensive health check
GET
/authsec/uflow/health/tenant/:tenant_id
Single tenant DB health
GET
/authsec/uflow/health/tenants
All tenant DBs health
Admin Authentication ( /authsec/uflow/auth/admin )
Method
Path
Description
GET
/authsec/uflow/auth/admin/challenge
Get auth challenge
POST
/authsec/uflow/auth/admin/login/precheck
Pre-login check
POST
/authsec/uflow/auth/admin/login/bootstrap
Bootstrap first admin
POST
/authsec/uflow/auth/admin/login
Admin login
POST
/authsec/uflow/auth/admin/login-hybrid
Hybrid login
POST
/authsec/uflow/auth/admin/register
Register admin
POST
/authsec/uflow/auth/admin/complete-registration
Complete registration
POST
/authsec/uflow/auth/admin/forgot-password
Initiate password reset
POST
/authsec/uflow/auth/admin/forgot-password/verify-otp
Verify OTP
POST
/authsec/uflow/auth/admin/forgot-password/reset
Reset password
End-User Authentication ( /authsec/uflow/auth/enduser )
Method
Path
Description
GET
/authsec/uflow/auth/enduser/challenge
Get challenge
POST
/authsec/uflow/auth/enduser/initiate-registration
Start registration
POST
/authsec/uflow/auth/enduser/verify-otp
Verify OTP + complete registration
POST
/authsec/uflow/auth/enduser/login/precheck
Pre-login check
POST
/authsec/uflow/auth/enduser/webauthn-callback
WebAuthn assertion callback
POST
/authsec/uflow/auth/enduser/delegate-svid
Delegate SPIFFE SVID
Device Authorization Grant – RFC 8628 ( /authsec/uflow/auth/device )
Method
Path
Auth
Description
POST
/authsec/uflow/auth/device/code
Public
Device requests code
POST
/authsec/uflow/auth/device/token
Public
Device polls for token
GET
/authsec/uflow/auth/device/activate/info
Public
Get device info for UI
POST
/authsec/uflow/auth/device/verify
JWT
User authorises device
GET
/authsec/uflow/activate
Public
Activation UI page
Voice Authentication ( /authsec/uflow/auth/voice )
Method
Path
Auth
Description
POST
/authsec/uflow/auth/voice/initiate
Public
Initiate voice auth
POST
/auth
[truncated]
Authenticated endpoints (JWT + tenant required):
All admin endpoints require JWT + admin:access + tenant validation .
Method
Path
Description
POST
/authsec/uflow/admin/roles
Create role
GET
/authsec/uflow/admin/roles
List roles
GET
/authsec/uflow/admin/roles/:role_id
Get role by ID
PUT
/authsec/uflow/admin/roles/:role_id
Update role
DELETE
/authsec/uflow/admin/roles/:role_id
Delete role
POST
/authsec/uflow/admin/bindings
Assign role (scoped)
GET
/authsec/uflow/admin/bindings
List bindings
POST
/authsec/uflow/admin/permissions
Register permission
GET
/authsec/uflow/admin/permissions
List permissions
DELETE
/authsec/uflow/admin/permissions/:id
Delete permission
GET
/authsec/uflow/admin/permissions/resources
List resources
GET/POST/PUT/DEL

[truncated]

---
source: "https://github.com/anway-dev/anway"
hn_url: "https://news.ycombinator.com/item?id=48909477"
title: "Show HN: A agentic nervous system for all DevOps tools"
article_title: "GitHub - anway-dev/anway: Anway (beta) — the central nervous system of your software organisation: agentic DevOps, one surface across all your tools · GitHub"
author: "rajvr"
captured_at: "2026-07-14T17:04:07Z"
capture_tool: "hn-digest"
hn_id: 48909477
score: 1
comments: 0
posted_at: "2026-07-14T16:39:43Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A agentic nervous system for all DevOps tools

- HN: [48909477](https://news.ycombinator.com/item?id=48909477)
- Source: [github.com](https://github.com/anway-dev/anway)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T16:39:43Z

## Translation

タイトル: Show HN: すべての DevOps ツールのエージェント神経システム
記事のタイトル: GitHub - anway-dev/anway: Anway (ベータ) — ソフトウェア組織の中枢神経系: エージェント的な DevOps、すべてのツールにわたる 1 つの表面 · GitHub
説明: Anway (ベータ) — ソフトウェア組織の中枢神経系: エージェント的な DevOps、すべてのツールにわたる 1 つの表面 - anway-dev/anway
HN テキスト: 既存のすべてのツールを 1 つのワークフローに接続する、エージェント型製品デバッグ ツールを作成しました。各コネクタは、データにインデックスを付けるサブエージェントであり、メイン エージェントは、コネクタを使用して取得したナレッジ グラフを使用して、サービスがダウンしている理由などのクエリに応答します。そのため、散乱が少なく、幻覚的な応答が期待できます。

記事本文:
GitHub - anway-dev/anway: Anway (ベータ) — ソフトウェア組織の中枢神経系: エージェント的な DevOps、すべてのツールにわたる 1 つの表面 · GitHub
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
アンウェイ開発
/
ある

それにしても
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
900 コミット 900 コミット .github/ workflows .github/ workflows アプリ アプリ アセット アセット コネクタ コネクタ ドキュメント ドキュメント インフラ インフラ パッケージ パッケージ スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE PLAN.md PLAN.md QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yamlturbo.jsonturbo.json vitest.integration.config.ts vitest.integration.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ソフトウェア組織の中枢神経系。ベータ — 積極的に開発されています。 API とスキーマは変更される可能性があります。
GitHub、Datadog、Linear、K8s、Loki、Prometheus、Jira、ArgoCD、PagerDuty、Terraform、および任意のクラウド プロバイダーを単一のインテリジェンス サーフェスに接続します。組織内のすべての人 (SRE、PM、BA、Dev) は、1 つのオーケストレーターを通じてソフトウェア ライフサイクル全体をクエリし、行動し、管理します。
開発ツールではありません。 Product、Eng、SRE の間の結合組織。
コネクタが追加されるたびに、組織全体のインテリジェンスが向上します。一度聞いてみてください
すべてのツールにわたって、管理された書き込みパスの背後で自信を持って動作します。
🌐 Web サイト: www.anway.dev · 🎬 デモ: 以下の 90 秒のウォークスルーをご覧ください。
https://github.com/anway-dev/anway/releases/download/v0.1.0-beta/anway-demo.mp4
上記のプレーヤーが読み込まれない場合は、 ▶ anway.dev で 90 秒のデモを視聴するか、サムネイルをクリックしてください。
合図が鳴る→一度聞いてみる

→ エージェントは、すべてのツールにわたって根本原因を追跡します。 → ゲートされたアクション、その後、サービス、パイプライン、コネクタ、監査証跡のツアーを行います。
Anway は、anway-dev 組織の下のリポジトリに分割されています。
製品デモ ビデオとそのキャプチャ ハーネスは、このリポジトリにあります。
apps/web/demo/ 。
アプリ/
web/Next.js UI (ポート 3000)
ゲートウェイ/ Fastify BFF — 認証、RBAC、監査、コネクタ プロキシ (ポート 4000)
エージェントサービス/ Python FastAPI — LLM 推論、エピソードグラフ (ポート 8000)
cli/アンウェイ CLI
パッケージ/
エージェント/オーケストレーター ハーネス、専門エージェント、IModelProvider
k8s/クラスタークライアント
ui/共有コンポーネント
タイプ/共有 TypeScript タイプ
インフラ/
docker-compose.yml 開発依存関係 (Postgres、Redis、Neo4j)
docker-compose.dev.yml 拡張開発スタック (Prometheus、Grafana、OTEL)
Helm/anway/ Production Helm チャート
terraform/AWS/GCP/Azure Terraform モジュール
ストレージ:
依存関係
最小バージョン
注意事項
Node.js
20LTS
pnpm
9
npm i -g pnpm
ドッカー + 構成
24
ローカルインフラ向け
パイソン
3.11
エージェントサービスをローカルで実行している場合のみ
地域開発
git clone <リポジトリ>
とにかくCDを
pnpmインストール
2.インフラを開始する
docker compose -f infra/docker-compose.yml up -d
これは、Postgres (5432)、Redis (6379)、Neo4j (7474/7687) から始まります。
docker compose -f infra/docker-compose.yml ps
すべてのサービスが正常であることを示すはずです。
cp apps/gateway/.env.example apps/gateway/.env
デフォルトは、ローカル開発ではそのまま使用できます。すぐに設定したい唯一の値は次のとおりです。
# LLM プロバイダーを 1 つ選択します
ANTHROPIC_API_KEY = sk-ant-...
# または
OPENAI_API_KEY = sk-...
# または両方を空白のままにする - Ollama がインストールされている場合はローカルで実行します
LLM キーが設定されていない場合、チャット エンドポイントはステータス 200 の { エラー: "LLM プロバイダーが構成されていません"、コード: "NO_PROVIDER" } を返し、UI にはキーを構成するためのインライン プロンプトが表示されます。クラッシュはありません。
CD アプリ/ゲートウェイ
pnpm prisma 移行デプロイ
pnpmp

risma db シード # デモ テナントに 22 のサービス、アラート、インシデント、パイプラインを追加します
CD ../..
シードはべき等であるため、安全に再実行できます。
# ターミナル 1 — ゲートウェイ (http://localhost:4000)
cd アプリ/ゲートウェイ && pnpm dev
# ターミナル 2 — Web UI (http://localhost:3000)
cd アプリ/ウェブ && pnpm dev
http://localhost:3000 を開きます。
デモ モードでは、事前にシードされた単一のテナントがプロビジョニングされ、資格情報なしで誰でもログインできます。ライブデモと内部レビューを目的としています。運用環境では決して有効にしないでください。
# アプリ/ゲートウェイ/.env
DEMO_MODE = true
使用する
カール -X POST http://localhost:4000/api/auth/demo
# → { "トークン": "<jwt>", "expiresIn": "24h" }
DEMO_MODE=true の場合、Web UI にはログイン ページに [デモを試す] ボタンが表示されます。
デモ テナント (ID 00000000-0000-0000-0000-000000000001 ) には、シード スクリプトによって以下が設定されます。
3 つの環境 (ステージング / preprod / prod) にわたる 22 のサービス
10 件以上のアラート (重大度の組み合わせ)
10 件以上のインシデント (進行中/解決済みの混合)
ステージランとゲートイベントを備えた 3 つのパイプライン
LLM プロバイダー (チャットに必要なプロバイダー):
# 形式: 名前:テナント ID、名前 2:テナント ID2
CONNECTOR_API_KEYS = e2e キー:00000000-0000-0000-0000-000000000001
ウェブ ( apps/web/.env.local )
GATEWAY_URL = http://127.0.0.1:4000 # ゲートウェイのベース URL (サーバー側のみ)
他のすべての構成 (LLM キー、コネクタ トークン) はゲートウェイ内に存在し、Web アプリ内には存在しません。
cd apps/gateway && npx tsc --noEmit
cd apps/web && npx tsc --noEmit
マージの前に、両方とも 0 エラーを返す必要があります。
pnpm test # すべてのワークスペーステストを実行する
cd apps/gateway && pnpm test # ゲートウェイのみ
E2E (劇作家)
フルスタックの実行が必要です (ゲートウェイ + Web + インフラ)。
# インフラ + アプリを最初に開始します (上記のローカル開発を参照)
# 次に:
cd apps/web && pnpm exec プレイライト テスト
# 特定の仕様を実行します。
pnpm exec playwright テスト e2e/99-certification.spec.ts
認証仕様 (99 認証。

spec.ts ) はゲートであり、すべての Wave にわたって 37 の受け入れ基準すべてを検証します。
オプション A — Docker Compose (単一ホスト)
# イメージをビルドする
docker build -t anway-gateway:最新のアプリ/ゲートウェイ
docker build -t anway-web:最新のアプリ/ウェブ
# 設定する
cp apps/gateway/.env.example apps/gateway/.env.prod
# apps/gateway/.env.prod を編集 — DATABASE_URL、REDIS_URL、JWT_PRIVATE_KEY などを設定します。
# 開始
docker compose -f infra/docker-compose.yml \
-f infra/prod/docker-compose.override.yml \
上 -d
オプション B — Kubernetes (Helm)
infra/helm/anway/ の Helm チャートには、ゲートウェイ デプロイメント、Web デプロイメント、HPA (2 ～ 10 ポッド、CPU 70%)、PDB (minAvailable: 1)、NetworkPolicy、Ingress (nginx)、IRSA アノテーション付き ServiceAccount が含まれます。
# 1. 名前空間を作成する
kubectl はとにかく名前空間を作成します
# 2. シークレットを作成する
kubectl シークレットを作成します 汎用 anway-secrets \
--namespace そのまま \
--from-literal=database-url= " postgresql://... " \
--from-literal=redis-url= " redis://... " \
--from-literal=jwt-private-key= " $( cat jwt.key ) " \
--from-literal=jwt-public-key= " $( cat jwt.key.pub ) "
#3. インストールチャート
helm install anway infra/helm/anway \
--namespace そのまま \
--setgateway.image.tag= < バージョン > \
--set web.image.tag= < バージョン > \
--set ingress.host=anway.yourdomain.com
# 4. 移行の実行 (ジョブ)
kubectl apply -f infra/k8s/merge-job.yaml -n anway
オプション C — Terraform (AWS EKS)
cd インフラ/terraform/環境/aws
# 設定する
cp terraform.tfvars.example terraform.tfvars
# 編集: aws_region、cluster_name、db_password など。
テラフォームの初期化
テラフォーム計画
テラフォーム適用
これにより、EKS クラスター、RDS PostgreSQL (本番環境のマルチ AZ)、ElastiCache Redis、ECR リポジトリ、IAM ロール (IRSA)、TF 状態の S3 バケットがプロビジョニングされます。 GCP (環境/gcp) モジュールと Azure (環境/azure) モジュールは同じパターンに従います。
[ ]ダ

ポッドから到達可能な TABASE_URL + REDIS_URL
[ ] JWT_PRIVATE_KEY + JWT_PUBLIC_KEY セット (RS256、HS256 シークレットではありません)
[ ] DEMO_MODE が存在しないか、false に設定されています
[ ] ALLOW_DEV_TOKEN がありません (本番環境では決して設定されません)
[ ] Postgres 移行が適用されました (prisma 移行デプロイ)
[ ] デモ テナントを使用する場合のシード実行
[ ] ヘルスチェック応答: GET /health/live → 200、GET /health/ready → 200
[ ] SENTRY_DSN がエラー追跡用に構成されています
[ ] テレメトリが必要な場合に OTEL コレクタに到達可能
[ ] 開発者のデフォルトからローテーションされたコネクタ API キー
[ ] ANWAY_WEBHOOK_TOKEN が設定され、alertmanager/CI 構成と一致します
[ ] RDS 自動バックアップが有効になっています (Terraform はこれを自動的に実行します)
健康エンドポイント
エンドポイント
目的
GET /健康/ライブ
Liveness — プロセスが起動している場合は常に 200 を返します
GET /health/ready
準備状況 — Postgres + Redis の接続をチェックします
/health/ready を使用するように Kubernetes プローブを構成します。
役割
能力
管理者
すべての操作
スレ
ゲートの承認、インシデントの作成、デプロイのトリガー
開発者
パイプラインを作成し、非本番環境でステージを実行する
午後/午後
読み取り専用
ロールは、アクセス ビューまたは API を介してユーザー プロビジョニング時に設定されます。
Anway は、OIDC 準拠のプロバイダー (Azure AD、Okta、Google Workspace、Keycloak、Dex) をサポートしています。
OIDC_ISSUER_URL = https://login.microsoftonline.com/<テナント>/v2.0
OIDC_CLIENT_ID = <アプリ ID>
OIDC_CLIENT_SECRET = <シークレット>
OIDC_REDIRECT_URI = https://anway.yourdomain.com/auth/oidc/callback
OIDC_TENANT_ID = <anway-tenant-uuid>
ローカル テスト用に、デモ作成スタックには模擬 OIDC プロバイダーとして Dex ( infra/demo/dex/ ) が含まれています。
チャットに「AI モデルが設定されていません」と表示される
apps/gateway/.env に少なくとも 1 つの LLM プロバイダー キーを設定します。ゲートウェイを再起動します。
シードが「関係が存在しません」で失敗する
移行は実行されていません。最初に cd apps/gateway && pnpm prisma mergedeploy を実行します。
GET /health/ready は 503 を返します
Postgres または Re

dis に到達できません。 DATABASE_URL と REDIS_URL を確認してください。インフラコンテナーが正常であることを確認します: docker compose ps 。
プル後の TypeScript エラー
リポジトリのルートで pnpm install を実行して新しい依存関係を取得し、再度 npx tsc --noEmit を実行します。
Playwright テストが「デモ ユーザーが見つかりません」で失敗する
シードを再実行します: cd apps/gateway && pnpm prisma db seed 。
負荷がかかると SSE ストリームが低下する
Redis は、複数のゲートウェイ ポッドにわたる SSE ファンアウトに必要です。 REDIS_URL が設定されていることを確認してください。単一ポッドのデプロイメントは Redis なしで機能しますが、ファンアウトすることはできません。
貢献は歓迎です — 詳細については CONTRIBUTING.md を参照してください。
セットアップ、ワークフロー、ガイドライン、および CODE_OF_CONDUCT.md 。
脆弱性が見つかりましたか?非公開で報告してください — を参照してください
セキュリティ.md 。何に関しても公開問題を開かないでください
悪用可能。デプロイする前に、そこにある非運用デフォルトのテーブルに注目してください。
Anway (ベータ版) — ソフトウェア組織の中枢神経系: エージェント的な DevOps、すべてのツールにわたる 1 つの表面
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
アンウェイ v0.1.0-ベータ版
最新の
2026 年 7 月 14 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲート

[切り捨てられた]

## Original Extract

Anway (beta) — the central nervous system of your software organisation: agentic DevOps, one surface across all your tools - anway-dev/anway

I have created a agentic prod debugging tool which connects all the existing tools to one single workflow. Each connector is a subagent which index the data and a main agent answers the queries like why the service is down using the knowledge graph it acquired using the connectors, Hence less scatter gather and less hallucinating answers you can expect.

GitHub - anway-dev/anway: Anway (beta) — the central nervous system of your software organisation: agentic DevOps, one surface across all your tools · GitHub
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
anway-dev
/
anway
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
900 Commits 900 Commits .github/ workflows .github/ workflows apps apps assets assets connectors connectors docs docs infra infra packages packages scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PLAN.md PLAN.md QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml turbo.json turbo.json vitest.integration.config.ts vitest.integration.config.ts View all files Repository files navigation
The central nervous system of a software organisation. Beta — actively developed; APIs and schemas may change.
Connects GitHub, Datadog, Linear, K8s, Loki, Prometheus, Jira, ArgoCD, PagerDuty, Terraform, and any cloud provider into a single intelligence surface. Every person in the org — SRE, PM, BA, Dev — queries, acts, and governs the entire software lifecycle through one orchestrator.
Not a devtool. The connective tissue between Product, Eng and SRE.
Every connector added = more intelligence for the entire org. Ask once, see
everything across every tool, act with confidence behind a governed write path.
🌐 Website: www.anway.dev · 🎬 Demo: watch the 90-second walkthrough below.
https://github.com/anway-dev/anway/releases/download/v0.1.0-beta/anway-demo.mp4
If the player above doesn't load, ▶ watch the 90-second demo on anway.dev — or click the thumbnail:
A signal fires → ask once → agents trace the root cause across every tool → gated action, then a tour of services, pipelines, connectors and the audit trail.
Anway is split across repositories under the anway-dev org:
The product demo video and its capture harness live in this repo under
apps/web/demo/ .
apps/
web/ Next.js UI (port 3000)
gateway/ Fastify BFF — auth, RBAC, audit, connector proxy (port 4000)
agent-service/ Python FastAPI — LLM inference, episodic graph (port 8000)
cli/ anway CLI
packages/
agent/ Orchestrator harness, specialist agents, IModelProvider
k8s/ Cluster client
ui/ Shared components
types/ Shared TypeScript types
infra/
docker-compose.yml Dev dependencies (Postgres, Redis, Neo4j)
docker-compose.dev.yml Extended dev stack (Prometheus, Grafana, OTEL)
helm/anway/ Production Helm chart
terraform/ AWS / GCP / Azure Terraform modules
Storage:
Dependency
Min version
Notes
Node.js
20 LTS
pnpm
9
npm i -g pnpm
Docker + Compose
24
For local infra
Python
3.11
Only if running agent-service locally
Local Development
git clone < repo >
cd anway
pnpm install
2. Start infra
docker compose -f infra/docker-compose.yml up -d
This starts: Postgres (5432), Redis (6379), Neo4j (7474/7687).
docker compose -f infra/docker-compose.yml ps
All services should show healthy .
cp apps/gateway/.env.example apps/gateway/.env
The defaults work out of the box for local dev. The only value you may want to set immediately:
# Pick any one LLM provider
ANTHROPIC_API_KEY = sk-ant-...
# or
OPENAI_API_KEY = sk-...
# or leave both blank — runs Ollama locally if installed
If no LLM key is set, the chat endpoint returns { error: "No LLM provider configured", code: "NO_PROVIDER" } with status 200 and the UI shows an inline prompt to configure one. No crash.
cd apps/gateway
pnpm prisma migrate deploy
pnpm prisma db seed # populates demo tenant with 22 services, alerts, incidents, pipelines
cd ../..
The seed is idempotent — safe to re-run.
# Terminal 1 — gateway (http://localhost:4000)
cd apps/gateway && pnpm dev
# Terminal 2 — web UI (http://localhost:3000)
cd apps/web && pnpm dev
Open http://localhost:3000 .
Demo mode provisions a single pre-seeded tenant and lets anyone log in without credentials. Intended for live demos and internal reviews — never enable in production.
# apps/gateway/.env
DEMO_MODE = true
Use
curl -X POST http://localhost:4000/api/auth/demo
# → { "token": "<jwt>", "expiresIn": "24h" }
The web UI shows a Try Demo button on the login page when DEMO_MODE=true .
The demo tenant (ID 00000000-0000-0000-0000-000000000001 ) is populated by the seed script with:
22 services across 3 environments (staging / preprod / prod)
10+ alerts (mix of severities)
10+ incidents (mix of active / resolved)
3 pipelines with stage runs and gate events
LLM provider (one required for chat):
# Format: name:tenantId,name2:tenantId2
CONNECTOR_API_KEYS = e2e-key:00000000-0000-0000-0000-000000000001
Web ( apps/web/.env.local )
GATEWAY_URL = http://127.0.0.1:4000 # gateway base URL (server-side only)
All other config (LLM keys, connector tokens) lives in the gateway — never in the web app.
cd apps/gateway && npx tsc --noEmit
cd apps/web && npx tsc --noEmit
Both must return 0 errors before any merge.
pnpm test # run all workspace tests
cd apps/gateway && pnpm test # gateway only
E2E (Playwright)
Requires the full stack running (gateway + web + infra):
# Start infra + app first (see Local Development above)
# Then:
cd apps/web && pnpm exec playwright test
# Run a specific spec:
pnpm exec playwright test e2e/99-certification.spec.ts
The certification spec ( 99-certification.spec.ts ) is the gate — it verifies all 37 acceptance criteria across every wave.
Option A — Docker Compose (single host)
# Build images
docker build -t anway-gateway:latest apps/gateway
docker build -t anway-web:latest apps/web
# Configure
cp apps/gateway/.env.example apps/gateway/.env.prod
# Edit apps/gateway/.env.prod — set DATABASE_URL, REDIS_URL, JWT_PRIVATE_KEY, etc.
# Start
docker compose -f infra/docker-compose.yml \
-f infra/prod/docker-compose.override.yml \
up -d
Option B — Kubernetes (Helm)
The Helm chart at infra/helm/anway/ includes: gateway deployment, web deployment, HPA (2–10 pods, CPU 70%), PDB (minAvailable: 1), NetworkPolicy, Ingress (nginx), ServiceAccount with IRSA annotations.
# 1. Create namespace
kubectl create namespace anway
# 2. Create secrets
kubectl create secret generic anway-secrets \
--namespace anway \
--from-literal=database-url= " postgresql://... " \
--from-literal=redis-url= " redis://... " \
--from-literal=jwt-private-key= " $( cat jwt.key ) " \
--from-literal=jwt-public-key= " $( cat jwt.key.pub ) "
# 3. Install chart
helm install anway infra/helm/anway \
--namespace anway \
--set gateway.image.tag= < version > \
--set web.image.tag= < version > \
--set ingress.host=anway.yourdomain.com
# 4. Run migrations (Job)
kubectl apply -f infra/k8s/migrate-job.yaml -n anway
Option C — Terraform (AWS EKS)
cd infra/terraform/environments/aws
# Configure
cp terraform.tfvars.example terraform.tfvars
# Edit: aws_region, cluster_name, db_password, etc.
terraform init
terraform plan
terraform apply
This provisions: EKS cluster, RDS PostgreSQL (Multi-AZ in prod), ElastiCache Redis, ECR repositories, IAM roles (IRSA), S3 bucket for TF state. GCP ( environments/gcp ) and Azure ( environments/azure ) modules follow the same pattern.
[ ] DATABASE_URL + REDIS_URL reachable from pods
[ ] JWT_PRIVATE_KEY + JWT_PUBLIC_KEY set (RS256, not HS256 secret)
[ ] DEMO_MODE absent or set to false
[ ] ALLOW_DEV_TOKEN absent (never set in prod)
[ ] Postgres migrations applied (prisma migrate deploy)
[ ] Seed run if using demo tenant
[ ] Health check responding: GET /health/live → 200, GET /health/ready → 200
[ ] SENTRY_DSN configured for error tracking
[ ] OTEL collector reachable if telemetry required
[ ] Connector API keys rotated from dev defaults
[ ] ANWAY_WEBHOOK_TOKEN set and matches alertmanager/CI config
[ ] RDS automated backups enabled (Terraform does this automatically)
Health endpoints
Endpoint
Purpose
GET /health/live
Liveness — always returns 200 if process is up
GET /health/ready
Readiness — checks Postgres + Redis connectivity
Configure Kubernetes probes to use /health/ready .
Role
Capabilities
admin
All operations
sre
Approve gates, create incidents, trigger deploys
dev
Create pipelines, run stages on non-prod envs
pm / ba
Read only
Roles are set at user provisioning time via the Access view or API.
Anway supports any OIDC-compliant provider (Azure AD, Okta, Google Workspace, Keycloak, Dex).
OIDC_ISSUER_URL = https://login.microsoftonline.com/<tenant>/v2.0
OIDC_CLIENT_ID = <app-id>
OIDC_CLIENT_SECRET = <secret>
OIDC_REDIRECT_URI = https://anway.yourdomain.com/auth/oidc/callback
OIDC_TENANT_ID = <anway-tenant-uuid>
For local testing, the demo compose stack includes Dex ( infra/demo/dex/ ) as a mock OIDC provider.
Chat shows "No AI model configured"
Set at least one LLM provider key in apps/gateway/.env . Restart the gateway.
Seed fails with "relation does not exist"
Migrations haven't run. Run cd apps/gateway && pnpm prisma migrate deploy first.
GET /health/ready returns 503
Either Postgres or Redis is unreachable. Check DATABASE_URL and REDIS_URL . Confirm infra containers are healthy: docker compose ps .
TypeScript errors after pulling
Run pnpm install at the repo root to pick up any new dependencies, then npx tsc --noEmit again.
Playwright tests fail with "demo user not found"
Re-run the seed: cd apps/gateway && pnpm prisma db seed .
SSE streams drop under load
Redis is required for SSE fan-out across multiple gateway pods. Ensure REDIS_URL is set. Single-pod deployments work without Redis but cannot fan-out.
Contributions are welcome — see CONTRIBUTING.md for
setup, workflow and guidelines, and CODE_OF_CONDUCT.md .
Found a vulnerability? Please report it privately — see
SECURITY.md . Do not open a public issue for anything
exploitable. Note the non-production defaults table there before deploying.
Anway (beta) — the central nervous system of your software organisation: agentic DevOps, one surface across all your tools
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
Anway v0.1.0-beta
Latest
Jul 14, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigat

[truncated]

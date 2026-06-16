---
source: "https://github.com/OtterMind/Nubase"
hn_url: "https://news.ycombinator.com/item?id=48552581"
title: "Nubase – an open-source back end/deploy layer for AI-written apps"
article_title: "GitHub - OtterMind/Nubase: Turn AI-written code into real apps. Nubase is an open-source, AI-native backend platform for AI Coding, agentic applications, and modern product teams: Memory, Database, Storage, and Auth in one self-hostable service. · GitHub"
author: "jipengfei1016"
captured_at: "2026-06-16T10:41:53Z"
capture_tool: "hn-digest"
hn_id: 48552581
score: 2
comments: 0
posted_at: "2026-06-16T09:11:36Z"
tags:
  - hacker-news
  - translated
---

# Nubase – an open-source back end/deploy layer for AI-written apps

- HN: [48552581](https://news.ycombinator.com/item?id=48552581)
- Source: [github.com](https://github.com/OtterMind/Nubase)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T09:11:36Z

## Translation

タイトル: Nubase – AI で作成されたアプリ用のオープンソース バックエンド/デプロイ レイヤー
記事のタイトル: GitHub - OtterMind/Nubase: AI で書かれたコードを実際のアプリに変える。 Nubase は、AI コーディング、エージェント アプリケーション、最新の製品チーム向けのオープンソースの AI ネイティブ バックエンド プラットフォームです。メモリ、データベース、ストレージ、認証を 1 つの自己ホスト型サービスにまとめています。 · GitHub
説明: AI が作成したコードを実際のアプリに変換します。 Nubase は、AI コーディング、エージェント アプリケーション、最新の製品チーム向けのオープンソースの AI ネイティブ バックエンド プラットフォームです。メモリ、データベース、ストレージ、認証を 1 つの自己ホスト型サービスにまとめています。 - OtterMind/Nubase

記事本文:
GitHub - OtterMind/Nubase: AI が作成したコードを実際のアプリに変換します。 Nubase は、AI コーディング、エージェント アプリケーション、最新の製品チーム向けのオープンソースの AI ネイティブ バックエンド プラットフォームです。メモリ、データベース、ストレージ、認証を 1 つの自己ホスト型サービスにまとめています。 · GitHub
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
オッターマインド
/
ぬばせ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
47 コミット 47 コミット .github/ workflows .github/ workflows brand brand cloudflare/functions-dispatcher Cloudflare/functions-dispatcher docker/all-in-one docker/all-in-one docs docs フロントエンド フロントエンド スクリプト script src src .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.all-in-one Dockerfile.all-in-one ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md nubase-modules.png nubase-modules.png pg-docker-compose.yml pg-docker-compose.yml pom.xml pom.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI が作成したコードを実際のアプリに変換します。 Nubase はオープンソースの AI ネイティブ バックエンドであり、コーディング エージェントが直接駆動するデプロイ レイヤーであるため、生成されたアプリは数分で稼働します。 1 つの自己ホスト可能なサービス内の 8 つの機能モジュール: データベース、認証、ストレージ、資産、関数、AI ゲートウェイ、メモリ、および cron 。
エージェントは、個別のホスティング アカウントを使用せずに、データのモデル化 (データベース + 認証)、バックエンド ロジックのデプロイ (関数)、生成されたフロントエンドのパブリック CDN への公開 (アセット)、定期的な作業のスケジュール ( cron ) をすべて MCP ツールを通じて行うことができます。理にかなった Supabase スタイル (Postgres、REST、JWT、RLS、オブジェクト ストレージ、Studio ダッシュボード) に加え、AI コーディング エージェント用に構築されたファーストクラスのメモリと MCP サーフェス。
1. Claude Code または Codex で Nubase を使用する — 1 つのコマンド
プロジェクト フォルダーから次を実行します。
npx -y nubase_cli@最新のインストールスキル
その 1 つのコマンド:
📚 Claude Code と Codex の両方の Nubase スキルをインストールします。
🔌ワイヤー

MCP サーバーの設定をアップし、
🔐 ブラウザが開き、プロジェクトを承認して選択します。
Claude Code — このフォルダーで再起動し、 /mcp を実行して、nubase が接続されていることを確認します。
Codex — ~/.codex/config.toml に追加されます。 Codexを起動するだけです。
これにより、エージェントが Nubase インスタンス (ホストされているインスタンス、または独自のインスタンス (ステップ 2 でスピンアップ)) に接続されます。次のコマンドを使用して、任意のインスタンスで CLI を指定します。
npx -y nubase_cli@最新のインストールスキル\
--studio-url https://studio.example.com \
--nubase-url https://api.example.com
2. 独自の Nubase を 1 つのコマンドで実行します
オールインワンの Docker イメージには、PostgreSQL + Redis + バックエンド + Studio がバンドルされています。
docker run -d --name nubase \
-p 9999:9999 -p 5432:5432 \
-v nubase_data:/data \
< 名前空間 > /nubase:latest
Studio → http://localhost:9999/studio — アカウントを作成し、プロジェクトを作成し、「プロビジョニング」をクリックしてデータベースを初期化します。
API → http://localhost:9999 (Studio UI はバックエンドにバンドルされ、同じポートで提供されます)
初回起動シークレットは /data ボリュームに生成されます。プロジェクトを保持するにはボリュームを保持してください。安定したシークレットを使用した実際のデプロイメントについては、「Docker を使用したセルフホスト」を参照してください。
エージェントは、MCP ツールを通じて Nubase を直接操作できるようになりました。これにより、スキーマの検査、テーブルの作成、SQL の実行、認証とストレージの管理、エッジ機能のデプロイ、パブリック CDN へのフロントエンドの公開、cron ジョブのスケジュール、および永続メモリの読み取り/書き込みが可能になりました。尋ねてみてください:
「RLS を使用して Todos テーブルを作成し、オープン数を返すエッジ関数をデプロイし、それを呼び出す 1 ページの UI を Assets に公開し、デプロイメントを記憶します。」
生成→ライブの完全なウォークスルーについては、「AI 生成アプリのデプロイ」を参照してください。
単一のオールインワン イメージには、独自のボックスで Nubase を実行するために必要なものがすべて揃っています。1 行で、構成ファイルや外部サービスは必要ありません。
試してみましょう (自動生成されたシークレット、ボリュームに保存されます)

私）：
docker run -d --name nubase -p 9999:9999 -p 5432:5432 \
-v nubase_data:/data < 名前空間 > /nubase:latest
運用環境 (暗号化されたプロジェクト資格情報が再起動後も存続するように、安定したシークレットを固定します):
docker run -d --name nubase -p 9999:9999 -p 5432:5432 \
-v nubase_data:/data \
-e PGRST_ENCRYPTION_MASTER_KEY= " $( openssl rand -base64 32 ) " \
-e METADATA_SERVICE_ROLE_KEY= " $( openssl rand -base64 48 ) " \
< 名前空間 > /nubase:latest
他のすべては環境変数 (Postgres、Redis、S3/R2 ストレージ、SMTP、OAuth、LLM プロバイダー) を介して構成されます。完全なリストとマルチアーキテクチャ ( amd64 + arm64 ) に関する注記については、 docs/docker-all-in-one.md を参照してください。
<your-namespace> を、イメージが公開されている Docker Hub 名前空間に置き換えます。
AI ネイティブ アプリケーションには CRUD 以上のものが必要です。初日からユーザー メモリ、取得、認証、ストレージ、データベース API、プロジェクトの分離が必要です。バックエンド層がなければ、AI コーディング セッションごとに別のデモが生成されますが、そのデモには依然として数週間のインフラストラクチャ作業が必要です。
Supabase は優れていますが、そのオープンソースのセルフホスト型スタックは単一のプロジェクトを中心に設計されています。 Nubase は、独自のインフラストラクチャ上に 1 つのスタジオ、1 つのバックエンド サービス、および多数の分離された AI プロジェクトを必要とする AI チームとセルフホスティング者向けに構築されており、次の 3 つの独自の追加機能が追加されています。
メモリはファーストクラスのプリミティブです。耐久性のあるメモリ、エンティティ抽出、履歴、およびハイブリッド取得は、別個のベクトル ストア スクリプトとして追加されるのではなく、組み込まれています。
AI コーディングは実際のバックエンド ターゲットを取得します。エージェントはテーブルを作成し、REST API を呼び出し、メモリに書き込み、MCP フレンドリーなツールを通じてスキーマを検査します。
セルフホスティングは多くのプロジェクトをサポートします。単一のコントロール プレーンがプロビジョニングし、複数の分離されたプロジェクト データベースにルーティングします。
🗄️ データベース — プロジェクトごとに 1 つの分離された PostgreSQL。 PostgREST-co

mpatible /rest/v1 API (select/filter/order/paginate/insert/update/upsert/delete);プロジェクトごとの JWT シークレット、ロール、スキーマ キャッシュ。 JWT クレームを使用した行レベルのセキュリティ。
🔐 認証 — Supabase スタイルのサインアップ/ログインおよびリフレッシュ トークンのローテーション。 MFA/TOTP、OTP およびマジック リンク、匿名サインイン。 OAuth (Google / GitHub / WeChat) および SAML SSO。プロジェクトごとの anon/authenticated/service_role トークン。
📦 ストレージ — S3 互換 (Cloudflare R2 / AWS S3 / MinIO);パブリック/プライベート バケット、署名付き URL、サイズおよび MIME コントロール。大規模なドキュメント/アセットのワークロード向けのオプションの S3 ベクター。
🌐 アセット (静的 CDN) — 生成されたフロントエンドを公開します。プロジェクトごとのパブリック静的アセットは、Cache-Control/ETag/304 セマンティクスを使用して /assets/v1/** で提供されます。プロジェクトごとのデフォルトのキャッシュ ポリシーとカスタム CDN ドメイン。エージェントは MCP (assets_upload) 経由で直接公開します。
⚡ 関数 — バックエンド ロジックを /functions/v1/** で提供されるエッジ関数としてデプロイします。関数ごとのシークレット、呼び出しログ、レート制限、verify_jwt ;ローカルエグゼキューターまたはプラットフォーム用のCloudflareワーカー。エージェントは MCP 上でスキャフォールド/デプロイ/呼び出し (functions_deploy) します。
🤖 AI ゲートウェイ — プロジェクトごとのキーとトークン/コストの使用状況追跡を備えた、OpenAI および Anthropic 互換のエンドポイント。
🧠 メモリ — Mem0 スタイルのメモリ API。 LLM によるファクト抽出 (追加/更新/削除/なし)。 pgvector + Postgres フルテキスト + エンティティ ブーストによるハイブリッド取得。エンティティ ストアと追加専用の履歴。 OpenAI、Anthropic、および OpenAI 互換プロバイダーと連携します。
⏰ スケジュールされたジョブ (cron) — crontab スケジュールでエッジ関数または名前付きデータベース関数を呼び出す繰り返しジョブ。実行履歴のあるコントロール プレーンによって実行されます。 MCP ( cron_create ) 経由で管理されます。
🧰 AI コーディングとエージェント — スキーマ検査、SQL 実行、RLS エクスポート、プロジェクト初期化、

d 記憶。認証、REST、ストレージ、メモリにわたる 1 つの一貫したプロジェクト トークン モデル。
🎛️ Studio — プロジェクト、SQL (実行履歴付き)、ユーザー、ストレージ、メモリ エクスプローラー用の Next.js ダッシュボード。
Nubase には 2 つのデータベース層があります。
メタデータ データベース — プラットフォーム ユーザー、プロジェクト構成、暗号化されたプロジェクト認証情報、所有権、プラットフォーム設定、SQL スニペット、および実行記録。
プロジェクト データベース — 各プロジェクトは、認証テーブル、ストレージ メタデータ、メモリ テーブル、アプリケーション テーブルを備えた独自の PostgreSQL データベースを取得します。
リクエストでは 2 トークン モデルが使用されます。 apikey はプロジェクト + ロール ( anon /Authenticated / service_role ) を識別し、Authorization: Bearer <jwt> は RLS およびメモリ所有権のエンド ユーザーを識別します。リクエスト フィルタは、 apikey からプロジェクトを解決し、JDBC を正しいプロジェクト データベースにルーティングし、リクエスト コンテキストを設定します。
要件: Java 17、Maven、Docker、Node.js + pnpm。
# 1. Postgres を開始する (15 + pgvector)
docker compose -f pg-docker-compose.yml up -d
# 2. 必要なシークレット
import PGRST_ENCRYPTION_MASTER_KEY= " $( openssl rand -base64 32 ) "
import METADATA_SERVICE_ROLE_KEY= " 長いランダムな管理トークンで置き換える "
import OPENAI_API_KEY= " sk-... " # オプション、LLM 搭載メモリの場合のみ
# 3. バックエンド → http://localhost:9999
mvn spring-boot:run
#4. スタジオ → http://localhost:3000
cd フロントエンド && pnpm インストール && pnpm dev:studio
オールインワン イメージを自分でビルドするには、 docker build -f Dockerfile.all-in-one -t nubase:local を実行します。
curl -X POST http://localhost:9999/mem/v1/memories \
-H " apikey: $NUBASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"userId":"user-42","messages":[{"role":"user","content":"私は寿司よりステーキが好きで、私の犬の名前はモチです。"}]} '
curl -X POST http://localhost:9999/mem/v1/search \
-H " apikey: $NUB

ASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"userId":"user-42","query":"彼らはどんな食べ物が好きですか?"} '
REST API を使用します (todos テーブルを作成した後)。
curl " http://localhost:9999/rest/v1/todos?select=* " -H " apikey: $NUBASE_ANON_KEY "
curl -X POST " http://localhost:9999/rest/v1/todos " \
-H " apikey: $NUBASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"text":"最初のオープンソース リリースを出荷する"} '
ドキュメント
AI 生成アプリのデプロイ (生成 → ライブ)
エージェントを接続 (クロード / コーデックス / カーソル)
エッジ機能 · アセット (静的 CDN) · スケジュールされたジョブ (cron)
Nubase は初期段階ですが、8 つのモジュール (データベース、認証、ストレージ、資産、関数、AI ゲートウェイ、メモリ、cron) と Studio および MCP ブリッジがすべて揃っています。未実装: バックアップ/PITR、HA、エンタープライズ SSO/SCIM などのリアルタイムおよび運用上の追加機能。サーバーをパブリック インターネットに公開する前に、管理/管理エンドポイントを確認してください。
貢献や問題点は大歓迎です。 CONTRIBUTING.md および SECURITY.md を参照してください。これは初期の公開リリースであるため、フィードバックが次のリリースを形作ります。 🙌
AI が作成したコードを実際のアプリに変換します。 Nubase は、AI コーディング、エージェント アプリケーション、最新の製品チーム向けのオープンソースの AI ネイティブ バックエンド プラットフォームです。メモリ、データベース、ストレージ、認証を 1 つの自己ホスト型サービスにまとめています。
アパッチ

[切り捨てられた]

## Original Extract

Turn AI-written code into real apps. Nubase is an open-source, AI-native backend platform for AI Coding, agentic applications, and modern product teams: Memory, Database, Storage, and Auth in one self-hostable service. - OtterMind/Nubase

GitHub - OtterMind/Nubase: Turn AI-written code into real apps. Nubase is an open-source, AI-native backend platform for AI Coding, agentic applications, and modern product teams: Memory, Database, Storage, and Auth in one self-hostable service. · GitHub
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
OtterMind
/
Nubase
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
47 Commits 47 Commits .github/ workflows .github/ workflows brand brand cloudflare/ functions-dispatcher cloudflare/ functions-dispatcher docker/ all-in-one docker/ all-in-one docs docs frontend frontend script script src src .dockerignore .dockerignore .gitignore .gitignore .gitleaks.toml .gitleaks.toml CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.all-in-one Dockerfile.all-in-one LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md nubase-modules.png nubase-modules.png pg-docker-compose.yml pg-docker-compose.yml pom.xml pom.xml View all files Repository files navigation
Turn AI-written code into real apps. Nubase is an open-source, AI-native backend and deploy layer that a coding agent drives directly — so a generated app goes live in minutes. Eight capability modules in one self-hostable service: Database, Auth, Storage, Assets, Functions, AI Gateway, Memory, and cron .
An agent can model the data (Database + Auth), deploy backend logic ( Functions ), publish the generated frontend to a public CDN ( Assets ), and schedule recurring work ( cron ) — all through MCP tools, with no separate hosting account. Supabase-style where it makes sense (Postgres, REST, JWTs, RLS, object storage, a Studio dashboard), plus first-class Memory and an MCP surface built for AI coding agents.
1. Use Nubase in Claude Code or Codex — one command
From your project folder, run:
npx -y nubase_cli@latest install-skills
That single command:
📚 installs the Nubase skills for both Claude Code and Codex ,
🔌 wires up the MCP server config, and
🔐 opens a browser to authorize and pick your project.
Claude Code — restart it in this folder, run /mcp , and confirm nubase is connected.
Codex — it's added to ~/.codex/config.toml ; just start Codex.
This connects your agent to a Nubase instance (a hosted one, or your own — spin one up in step 2). Point the CLI at any instance with:
npx -y nubase_cli@latest install-skills \
--studio-url https://studio.example.com \
--nubase-url https://api.example.com
2. Run your own Nubase — one command
The all-in-one Docker image bundles PostgreSQL + Redis + the backend + Studio :
docker run -d --name nubase \
-p 9999:9999 -p 5432:5432 \
-v nubase_data:/data \
< your-namespace > /nubase:latest
Studio → http://localhost:9999/studio — create an account, create a project, click Provision to initialize its database.
API → http://localhost:9999 (the Studio UI is bundled into the backend and served on the same port)
First-boot secrets are generated into the /data volume; keep the volume to retain your projects. For a real deployment with stable secrets, see Self-host with Docker .
Your agent can now operate Nubase directly through MCP tools — inspect schema, create tables, run SQL, manage auth & storage, deploy edge functions, publish a frontend to the public CDN, schedule cron jobs , and read/write durable memory . Try asking:
"Create a todos table with RLS, deploy an edge function that returns the open count, publish a one-page UI to Assets that calls it, and remember the deployment."
See Deploy an AI-generated app for the full generate → live walkthrough.
The single all-in-one image is everything you need to run Nubase on your own box — one line, no compose file, no external services .
Try it (auto-generated secrets, kept in the volume):
docker run -d --name nubase -p 9999:9999 -p 5432:5432 \
-v nubase_data:/data < your-namespace > /nubase:latest
Production (pin stable secrets so encrypted project credentials survive restarts):
docker run -d --name nubase -p 9999:9999 -p 5432:5432 \
-v nubase_data:/data \
-e PGRST_ENCRYPTION_MASTER_KEY= " $( openssl rand -base64 32 ) " \
-e METADATA_SERVICE_ROLE_KEY= " $( openssl rand -base64 48 ) " \
< your-namespace > /nubase:latest
Everything else is configured via environment variables — Postgres, Redis, S3/R2 storage, SMTP, OAuth, and LLM providers. See docs/docker-all-in-one.md for the full list and a multi-architecture ( amd64 + arm64 ) note.
Replace <your-namespace> with the Docker Hub namespace the image is published under.
AI-native applications need more than CRUD. They need user memory, retrieval, auth, storage, database APIs, and project isolation from day one. Without that backend layer, every AI coding session produces another demo that still needs weeks of infrastructure work.
Supabase is excellent, but its open-source self-hosted stack is designed around a single project. Nubase is built for AI teams and self-hosters who want one Studio, one backend service, and many isolated AI projects on their own infrastructure — with three opinionated additions:
Memory is a first-class primitive — durable memory, entity extraction, history, and hybrid retrieval are built in, not bolted on as a separate vector-store script.
AI coding gets a real backend target — agents create tables, call REST APIs, write memory, and inspect schema through MCP-friendly tools.
Self-hosting supports many projects — a single control plane provisions and routes to multiple isolated project databases.
🗄️ Database — one isolated PostgreSQL per project; a PostgREST-compatible /rest/v1 API (select/filter/order/paginate/insert/update/upsert/delete); per-project JWT secrets, roles, and schema cache; Row Level Security with JWT claims.
🔐 Auth — Supabase-style signup/login and refresh-token rotation; MFA/TOTP, OTP & magic links, anonymous sign-in; OAuth (Google / GitHub / WeChat) and SAML SSO; per-project anon / authenticated / service_role tokens.
📦 Storage — S3-compatible (Cloudflare R2 / AWS S3 / MinIO); public/private buckets, signed URLs, size & MIME controls; optional S3 Vectors for large document/asset workloads.
🌐 Assets (static CDN) — publish a generated frontend: per-project public static assets served at /assets/v1/** with Cache-Control/ETag/304 semantics; per-project default cache policy and custom CDN domain; agents publish directly over MCP ( assets_upload ).
⚡ Functions — deploy backend logic as edge functions served at /functions/v1/** ; per-function secrets, invocation logs, rate limits, verify_jwt ; local executor or Cloudflare Workers for Platforms; agents scaffold/deploy/invoke over MCP ( functions_deploy ).
🤖 AI Gateway — OpenAI- and Anthropic-compatible endpoints with per-project keys and token/cost usage tracking.
🧠 Memory — Mem0-style memory API; LLM-powered fact extraction (ADD/UPDATE/DELETE/NONE); hybrid retrieval over pgvector + Postgres full-text + entity boost; entity store and append-only history. Works with OpenAI, Anthropic, and OpenAI-compatible providers.
⏰ Scheduled Jobs (cron) — recurring jobs that invoke an edge function or a named database function on a crontab schedule, run by the control plane with run history; managed over MCP ( cron_create ).
🧰 AI Coding & Agents — an MCP bridge ( nubase_cli ) for schema inspection, SQL execution, RLS export, project init, and memory; one consistent project-token model across Auth, REST, Storage, and Memory.
🎛️ Studio — a Next.js dashboard for projects, SQL (with execution history), users, storage, and the memory explorer.
Nubase has two database layers:
Metadata database — platform users, project configs, encrypted project credentials, ownership, platform settings, SQL snippets, and execution records.
Project databases — each project gets its own PostgreSQL database with auth tables, storage metadata, memory tables, and application tables.
Requests use a two-token model: apikey identifies the project + role ( anon / authenticated / service_role ), and Authorization: Bearer <jwt> identifies the end user for RLS and memory ownership. A request filter resolves the project from the apikey , routes JDBC to the correct project database, and sets the request context.
Requirements: Java 17, Maven, Docker, Node.js + pnpm.
# 1. Start Postgres (15 + pgvector)
docker compose -f pg-docker-compose.yml up -d
# 2. Required secrets
export PGRST_ENCRYPTION_MASTER_KEY= " $( openssl rand -base64 32 ) "
export METADATA_SERVICE_ROLE_KEY= " replace-with-a-long-random-admin-token "
export OPENAI_API_KEY= " sk-... " # optional, only for LLM-powered Memory
# 3. Backend → http://localhost:9999
mvn spring-boot:run
# 4. Studio → http://localhost:3000
cd frontend && pnpm install && pnpm dev:studio
To build the all-in-one image yourself: docker build -f Dockerfile.all-in-one -t nubase:local .
curl -X POST http://localhost:9999/mem/v1/memories \
-H " apikey: $NUBASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"userId":"user-42","messages":[{"role":"user","content":"I prefer steak over sushi and my dog is named Mochi."}]} '
curl -X POST http://localhost:9999/mem/v1/search \
-H " apikey: $NUBASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"userId":"user-42","query":"what food do they like?"} '
Use the REST API (after creating a todos table):
curl " http://localhost:9999/rest/v1/todos?select=* " -H " apikey: $NUBASE_ANON_KEY "
curl -X POST " http://localhost:9999/rest/v1/todos " \
-H " apikey: $NUBASE_SERVICE_KEY " -H " Content-Type: application/json " \
-d ' {"text":"Ship the first open-source release"} '
Documentation
Deploy an AI-generated app (generate → live)
Connect agents (Claude / Codex / Cursor)
Edge Functions · Assets (static CDN) · Scheduled Jobs (cron)
Nubase is early-stage but all eight modules (Database, Auth, Storage, Assets, Functions, AI Gateway, Memory, cron) plus Studio and the MCP bridge are in place. Not yet implemented: Realtime and operational extras like backups/PITR, HA, and enterprise SSO/SCIM. Review the admin/management endpoints before exposing a server to the public internet.
Contributions and issues are welcome — see CONTRIBUTING.md and SECURITY.md . This is an early public release, so feedback shapes what comes next. 🙌
Turn AI-written code into real apps. Nubase is an open-source, AI-native backend platform for AI Coding, agentic applications, and modern product teams: Memory, Database, Storage, and Auth in one self-hostable service.
Apache

[truncated]

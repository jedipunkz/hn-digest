---
source: "https://github.com/landry-77/AI-DEVOPS-ENGINE"
hn_url: "https://news.ycombinator.com/item?id=48661117"
title: "AI DevOps Engine – bot posts PR fixes after testing in network-isolated Docker"
article_title: "GitHub - landry-77/AI-DEVOPS-ENGINE: Autonomous AI DevOps pipeline with air-gapped sandboxing, zero-data-retention LLM routing, PostgreSQL RLS multi-tenancy, and pluggable billing. Self-hosted, open-source, one docker-compose command. · GitHub"
author: "landry-77"
captured_at: "2026-06-24T15:46:33Z"
capture_tool: "hn-digest"
hn_id: 48661117
score: 2
comments: 2
posted_at: "2026-06-24T15:06:35Z"
tags:
  - hacker-news
  - translated
---

# AI DevOps Engine – bot posts PR fixes after testing in network-isolated Docker

- HN: [48661117](https://news.ycombinator.com/item?id=48661117)
- Source: [github.com](https://github.com/landry-77/AI-DEVOPS-ENGINE)
- Score: 2
- Comments: 2
- Posted: 2026-06-24T15:06:35Z

## Translation

タイトル: AI DevOps Engine – ネットワーク分離された Docker でのテスト後にボットが PR 修正を投稿
記事のタイトル: GitHub - landry-77/AI-DEVOPS-ENGINE: エアギャップ サンドボックス、データ保持ゼロの LLM ルーティング、PostgreSQL RLS マルチテナント、プラグイン可能な課金を備えた自律型 AI DevOps パイプライン。自己ホスト型、オープンソースの 1 つの docker-compose コマンド。 · GitHub
説明: エアギャップサンドボックス、データ保持ゼロの LLM ルーティング、PostgreSQL RLS マルチテナント、プラグイン可能な課金を備えた自律型 AI DevOps パイプライン。自己ホスト型、オープンソースの 1 つの docker-compose コマンド。 - landry-77/AI-DEVOPS-エンジン

記事本文:
GitHub - landry-77/AI-DEVOPS-ENGINE: エアギャップ サンドボックス、データ保持ゼロの LLM ルーティング、PostgreSQL RLS マルチテナント、プラグイン可能な課金を備えた自律型 AI DevOps パイプライン。自己ホスト型、オープンソースの 1 つの docker-compose コマンド。 · GitHub
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
{{ メッセージ

}}
ランドリー-77
/
AI-DEVOPS-エンジン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
18 コミット 18 コミット .github .github billing-collector billing-collector core-brain core-brain django-dashboard django-dashboard infra ingestion-service ingestion-service Sandbox-env Sandbox-env .env.example .env.example .gitignore .gitignore 1 1 Caddyfile Caddyfile Caddyfile.local Caddyfile.local FOR_BUSINESS_OWNERS.md FOR_BUSINESS_OWNERS.md FOR_TECHNICAL_AUDIENCE.md FOR_TECHNICAL_AUDIENCE.md LICENSE LICENSE Makefile Makefile README.md README.md calculator.py calculator.py check_pr6_comment.py check_pr6_comment.py check_pr_comments.py check_pr_comments.py debug_git.py debug_git.py debug_token.py debug_token.py docker-compose.local.yml docker-compose.local.yml docker-compose.web.yml docker-compose.web.yml docker-compose.worker.yml docker-compose.worker.yml docker-compose.yml docker-compose.yml list_repos.py list_repos.py setup.ps1 setup.ps1 setup_better_test.py setup_better_test.py setup_endtoend_test.py setup_endtoend_test.py target_bug.py target_bug.pytrigger_fix.pytrigger_fix.pytrigger_pr6.pytrigger_pr6.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GitHub Webhook を取り込み、LLM 経由でコード パッチを構築し、ネットワーク分離された Docker サンドボックス (Pytest/Jest) で実行し、検証された修正をレビュー用の PR コメントとして投稿する、自律型のデータ保持ゼロの AI DevOps パイプライン。これらはすべて 1 つの docker compose コマンドで自己ホストされます。
SaaS 手数料なし - 使用したトークンに対して AI プロバイダーに直接支払うだけです
データ保持ゼロ — コードはメモリ内でスクラブされ、推論後に破棄されます。
ネットワーク分離されたサンドボックス — パッチはネットワーク分離されたリソースで実行されます。

RCEスロットルコンテナ
マルチテナント — PostgreSQL の行レベル セキュリティにより、データベース エンジン レベルでテナントが分離されます。
クイックスタート (ゼロから実行まで約 2 分)
OpenRouter API キー — 無料利用枠は機能します
ngrok — 無料枠、ローカル Webhook を GitHub に公開します
# 1. シークレットを自動生成し、OpenRouter キーの入力を求める
セットアップを行う
# 2. サンドボックス イメージのプリベイク (1 回限り)
サンドボックスを作る
# 3. フルスタックを起動する
補う
Windows (PowerShell):
# 1. シークレットを自動生成し、OpenRouter キーの入力を求める
.\setup.ps1 セットアップ
# 2. サンドボックス イメージのプリベイク (1 回限り)
.\setup.ps1 サンドボックス
# 3. フルスタックを起動する
.\setup.ps1 アップ
ダッシュボード: http://localhost:8000 ・ゲートウェイ: http://localhost:3000
次に、Webhook を公開して PR を開きます。
# 4. ngrok 経由で公開する (別のターミナル)
ngrok http http://localhost:3000
#5. リポジトリ上の任意の PR を開きます — 残りはエンジンが処理します
詳細設定（初めて設定する場合）
1. 環境の構成
make setup は .env.example → .env をコピーし、 DJANGO_SECRET_KEY と FERNET_KEY の安全なランダム値を生成し、OpenRouter キーの入力を求めます。その後、さらに 2 つの値を手動で設定します。
GitHub アプリの .pem ファイルを certs/github_app.pem に配置します。
GitHub 設定 → 開発者設定 → GitHub アプリ → 新しい GitHub アプリ
設定:
GitHub アプリ名: ai-devops-bot
ホームページ URL: http://localhost:3000
Webhook URL: https://<your-ngrok-id>.ngrok-free.app/webhooks/github
Webhook シークレット: ランダムな文字列を選択 → .env で GITHUB_WEBHOOK_SECRET として設定
権限: コンテンツ: 書き込み、プル リクエスト: 読み取りおよび書き込み、チェック: 書き込み、メタデータ: 読み取り
秘密キーを生成 → certs/github_app.pem として保存
アプリIDをコピー → .envでGITHUB_APP_IDENTIFIERとして設定
make Sandbox # local-pytest-sandbox + local-jest-sandbox イメージをビルドする
構成 # docker 構成 -d
4. 機知を試す

カール
デモを作る
または手動で:
WEBHOOK_SECRET= " ${GITHUB_WEBHOOK_SECRET?} "
payload= ' {"action":"opened","pull_request":{"number":1},"repository":{"id":101,"full_name":"local-org/test-repo","clone_url":"local_vfs"},"installation":{"id":202}} '
sig= $( printf ' %s ' " $payload " | openssl dgst -sha256 -hmac " $WEBHOOK_SECRET " | awk ' {print $NF} ' )
curl -X POST http://localhost:3000/webhooks/github \
-H " Content-Type: application/json " \
-H " x-github-event: pull_request " \
-H " x-hub-signature-256: sha256= $sig " \
-d " $ペイロード "
仕組み
リポジトリで PR が開かれると、次のようになります。
GitHub が Webhook を送信 → ゲートウェイが HMAC 署名を検証
ゲートウェイはデフォルトのブランチにフィルタリング → Redis でタスクをエンキューします
Celery ワーカーがピックアップ → AI エンジンが秘密をスクラブ → OpenRouter を呼び出す
生成されたパッチは、ネットワークから分離された Docker サンドボックス (Pytest/Jest) で実行されます。
テストパス時: ボットはレビュー用の PR コメントとして修正を投稿します。
ダッシュボードのすべてのステップは http://localhost:8000 に記録されます。
chmod +x infra/patch-bot.sh
エイリアス patch-bot=./infra/patch-bot.sh
patch-bot my_app/main.py " ページ番号が総ページ数を超えるとページネーションが中断される "
建築
グラフTB
classDef 外部塗りつぶし:#2da44e、カラー:#fff、ストローク:#1a7f37
classDef 雲の塗りつぶし:#6366f1、カラー:#fff、ストローク:#4f46e5
classDef ゲートウェイの塗りつぶし:#0ea5e9、色:#fff、ストローク:#0284c7
classDef キューの塗りつぶし:#f59e0b、色:#fff、ストローク:#d97706
classDef ワーカーの塗りつぶし:#8b5cf6、色:#fff、ストローク:#7c3aed
classDef 脳塗り:#ec4899、色:#fff、ストローク:#db2777
classDef サンドボックスの塗りつぶし:#e17055、カラー:#fff、ストローク:#d35400
classDef ダッシュボードの塗りつぶし:#14b8a6、色:#fff、ストローク:#0d9488
classDef db 塗りつぶし:#64748b、カラー:#fff、ストローク:#475569
classDef 請求塗りつぶし:#f43f5e、色:#fff、ストローク:#e11d48
サブグラフ 外部["☁️ 外部サービス"]
GH["GitHub リポジトリ<br/>プッシュ / PR イベント"]
OR["OpenRouter API<br/>

GPT-4o-mini」]
終わり
サブグラフ プロキシ["🚀 リバース プロキシ (キャディ) :80"]
CADDY["Caddy プロキシ<br/>ルート: /webhooks/* → ゲートウェイ<br/>/* → ダッシュボード"]
終わり
サブグラフ Ingest["📡 インジェスト ゲートウェイ :3000"]
GW["Express.js 取り込みサービス<br/>HMAC 署名検証<br/>ブランチ フィルタリング ロジック"]
終わり
サブグラフ ブローカー["📬 メッセージ ブローカー"]
REDIS[(Redis キュー<br/>Celery ブローカー + バックエンド)]
終わり
サブグラフ ワーカー["⚙️ タスク ワーカー"]
CELERY["Celery ワーカー<br/>プロセス修復タスク"]
BEAT["Celery Beat<br/>スケジュールされた Cron ジョブ"]
終わり
サブグラフ AIEngine["🧠 AI エンジン :8010"]
BRAIN["FastAPI コア ブレイン<br/>オーケストレーション レイヤー"]
SCRUBBER["🔒 シークレットスクラバー<br/>メモリ内正規表現編集<br/>AWS キー / GH トークン / DB 認証"]
AI_INF["OpenRouter 推論<br/>データ保持ゼロ<br/>コード分析 + パッチ生成"]
HANDLER["GitHub API ハンドラー<br/>PR コメントを投稿"]
終わり
サブグラフ サンドボックス["🛡️ ネットワーク分離サンドボックス"]
SBX["Docker サンドボックス コンテナ<br/>ネットワークなし · 512MB RAM<br/>読み取り専用マウント"]
PYT["Pytest / Jest の実行<br/>テスト スイートの検証"]
終わり
サブグラフ ダッシュボード["📊 Django ダッシュボード :8000"]
DJANGO["Django 6 + Daphne ASGI<br/>Web UI / 管理者"]
POLL["ライブ ステータス ポーリング API<br/>/api/v1/logs-stream/"]
UI[「実行ログカード<br/>5秒ごとに自動更新」]
終わり
サブグラフ ストレージ["💾 データベース層"]
PG[("PostgreSQL<br/>マルチテナント")]
RLS["🔐 行レベルのセキュリティ<br
[切り捨てられた]
サービス
テクノロジー
役割
インジェストサービス
Node.js + Express
GitHub Webhook レシーバー、パス フィルタリング、タスク キューイング
コアブレイン
FastAPI + セロリ
AI オーケストレーション、シークレット スクラビング、Docker サンドボックス コントロール
ジャンゴダッシュボード
ジャンゴ 6 + ダフネ ASGI
Web UI、監査ログ、マルチテナント管理、請求
サンドボックス環境
Docker (ネットワーク分離)
事前にベイクされた Pytest/Jest イメージ、ネットワークなし、512MB RAM キャップ
集金人
パイソン
コスト

予測、使用量測定、AWS/GCP アンケート
redisブローカー
Redis 7
Celery メッセージ キュー + 結果バックエンド
セキュリティ
レイヤー
仕組み
コードストレージ
永続性ゼロ — PostgreSQL は運用メタデータのみを保存し、ソース コードは保存しません
LLM プライバシー
data_collection: すべての OpenRouter リクエストを拒否します — コードのトレーニングを法的にブロックします
秘密のスクラブ
メモリ内正規表現 — AWS キー、GH トークン、転送前にマスクされた DB 認証情報
パッチの実行
ネットワーク分離された Docker — ネットワーク アクセスなし、512MB RAM / 2 CPU のハード制限、ホスト FS マウントなし
テナントの分離
PostgreSQL RLS — データベースによる強制的な行分離、Django .filter() のバイパス
容器の漏れ
バックグラウンド cron — 実行フリーズ時に孤立したサンドボックス コンテナを自動削除します
について
エアギャップサンドボックス、データ保持ゼロの LLM ルーティング、PostgreSQL RLS マルチテナント、プラグイン可能な課金を備えた自律型 AI DevOps パイプライン。自己ホスト型、オープンソースの 1 つの docker-compose コマンド。
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

Autonomous AI DevOps pipeline with air-gapped sandboxing, zero-data-retention LLM routing, PostgreSQL RLS multi-tenancy, and pluggable billing. Self-hosted, open-source, one docker-compose command. - landry-77/AI-DEVOPS-ENGINE

GitHub - landry-77/AI-DEVOPS-ENGINE: Autonomous AI DevOps pipeline with air-gapped sandboxing, zero-data-retention LLM routing, PostgreSQL RLS multi-tenancy, and pluggable billing. Self-hosted, open-source, one docker-compose command. · GitHub
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
landry-77
/
AI-DEVOPS-ENGINE
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits .github .github billing-collector billing-collector core-brain core-brain django-dashboard django-dashboard infra infra ingestion-service ingestion-service sandbox-env sandbox-env .env.example .env.example .gitignore .gitignore 1 1 Caddyfile Caddyfile Caddyfile.local Caddyfile.local FOR_BUSINESS_OWNERS.md FOR_BUSINESS_OWNERS.md FOR_TECHNICAL_AUDIENCE.md FOR_TECHNICAL_AUDIENCE.md LICENSE LICENSE Makefile Makefile README.md README.md calculator.py calculator.py check_pr6_comment.py check_pr6_comment.py check_pr_comments.py check_pr_comments.py debug_git.py debug_git.py debug_token.py debug_token.py docker-compose.local.yml docker-compose.local.yml docker-compose.web.yml docker-compose.web.yml docker-compose.worker.yml docker-compose.worker.yml docker-compose.yml docker-compose.yml list_repos.py list_repos.py setup.ps1 setup.ps1 setup_better_test.py setup_better_test.py setup_endtoend_test.py setup_endtoend_test.py target_bug.py target_bug.py trigger_fix.py trigger_fix.py trigger_pr6.py trigger_pr6.py View all files Repository files navigation
An autonomous, zero-data-retention AI DevOps pipeline that ingests GitHub webhooks, constructs code patches via LLM, runs them in network-isolated Docker sandboxes (Pytest/Jest), and posts validated fixes as PR comments for your review — all self-hosted with one docker compose command.
No SaaS fees — you only pay the AI provider directly for tokens used
Zero data retention — code scrubbed in memory, destroyed after inference
Network-isolated sandbox — patches run in network-isolated, resource-throttled containers
Multi-tenant — PostgreSQL Row-Level Security isolates tenants at the database engine level
Quickstart (from zero to running in ~2 minutes)
OpenRouter API key — free tier works
ngrok — free tier, exposes your local webhook to GitHub
# 1. Auto-generate secrets, prompt for your OpenRouter key
make setup
# 2. Pre-bake sandbox images (one-time)
make sandbox
# 3. Launch the full stack
make up
Windows (PowerShell):
# 1. Auto-generate secrets, prompt for your OpenRouter key
.\setup.ps1 setup
# 2. Pre-bake sandbox images (one-time)
.\setup.ps1 sandbox
# 3. Launch the full stack
.\setup.ps1 up
Dashboard: http://localhost:8000 · Gateway: http://localhost:3000
Then expose your webhook and open a PR:
# 4. Expose via ngrok (separate terminal)
ngrok http http://localhost:3000
# 5. Open any PR on your repo — the engine handles the rest
Detailed Setup (for first-time configuration)
1. Configure Environment
make setup copies .env.example → .env , generates secure random values for DJANGO_SECRET_KEY and FERNET_KEY , then prompts for your OpenRouter key. After that, set two more values manually:
Place the GitHub App .pem file at certs/github_app.pem .
GitHub Settings → Developer settings → GitHub Apps → New GitHub App
Settings:
GitHub App name: ai-devops-bot
Homepage URL: http://localhost:3000
Webhook URL: https://<your-ngrok-id>.ngrok-free.app/webhooks/github
Webhook secret: pick a random string → set as GITHUB_WEBHOOK_SECRET in .env
Permissions: Contents: Write , Pull requests: Read & Write , Checks: Write , Metadata: Read
Generate a private key → save as certs/github_app.pem
Copy the App ID → set as GITHUB_APP_IDENTIFIER in .env
make sandbox # build local-pytest-sandbox + local-jest-sandbox images
make up # docker compose up -d
4. Test with a Curl
make demo
Or manually:
WEBHOOK_SECRET= " ${GITHUB_WEBHOOK_SECRET?} "
payload= ' {"action":"opened","pull_request":{"number":1},"repository":{"id":101,"full_name":"local-org/test-repo","clone_url":"local_vfs"},"installation":{"id":202}} '
sig= $( printf ' %s ' " $payload " | openssl dgst -sha256 -hmac " $WEBHOOK_SECRET " | awk ' {print $NF} ' )
curl -X POST http://localhost:3000/webhooks/github \
-H " Content-Type: application/json " \
-H " x-github-event: pull_request " \
-H " x-hub-signature-256: sha256= $sig " \
-d " $payload "
How It Works
When a PR is opened on your repo:
GitHub sends a webhook → gateway verifies HMAC signature
Gateway filters to default branch → enqueues task in Redis
Celery worker picks up → AI engine scrubs secrets → calls OpenRouter
Generated patch runs in network-isolated Docker sandbox (Pytest/Jest)
On test pass: bot posts the fix as a PR comment for your review
Dashboard logs every step at http://localhost:8000
chmod +x infra/patch-bot.sh
alias patch-bot=./infra/patch-bot.sh
patch-bot my_app/main.py " pagination breaks when page number exceeds total pages "
Architecture
graph TB
classDef external fill:#2da44e,color:#fff,stroke:#1a7f37
classDef cloud fill:#6366f1,color:#fff,stroke:#4f46e5
classDef gateway fill:#0ea5e9,color:#fff,stroke:#0284c7
classDef queue fill:#f59e0b,color:#fff,stroke:#d97706
classDef worker fill:#8b5cf6,color:#fff,stroke:#7c3aed
classDef brain fill:#ec4899,color:#fff,stroke:#db2777
classDef sandbox fill:#e17055,color:#fff,stroke:#d35400
classDef dashboard fill:#14b8a6,color:#fff,stroke:#0d9488
classDef db fill:#64748b,color:#fff,stroke:#475569
classDef billing fill:#f43f5e,color:#fff,stroke:#e11d48
subgraph External["☁️ External Services"]
GH["GitHub Repo<br/>Push / PR Events"]
OR["OpenRouter API<br/>GPT-4o-mini"]
end
subgraph Proxy["🚀 Reverse Proxy (Caddy) :80"]
CADDY["Caddy Proxy<br/>Routes: /webhooks/* → Gateway<br/>/* → Dashboard"]
end
subgraph Ingest["📡 Ingestion Gateway :3000"]
GW["Express.js Ingestion Service<br/>HMAC Signature Verify<br/>Branch Filtration Logic"]
end
subgraph Broker["📬 Message Broker"]
REDIS[(Redis Queue<br/>Celery Broker + Backend)]
end
subgraph Workers["⚙️ Task Workers"]
CELERY["Celery Worker<br/>Process Remediation Tasks"]
BEAT["Celery Beat<br/>Scheduled Cron Jobs"]
end
subgraph AIEngine["🧠 AI Engine :8010"]
BRAIN["FastAPI Core Brain<br/>Orchestration Layer"]
SCRUBBER["🔒 Secret Scrubber<br/>In-Memory Regex Redaction<br/>AWS Keys / GH Tokens / DB Creds"]
AI_INF["OpenRouter Inference<br/>Zero-Data-Retention<br/>Code Analysis + Patch Gen"]
HANDLER["GitHub API Handler<br/>Post PR Comments"]
end
subgraph Sandbox["🛡️ Network-Isolated Sandbox"]
SBX["Docker Sandbox Container<br/>No Network · 512MB RAM<br/>Read-Only Mount"]
PYT["Pytest / Jest Execution<br/>Test Suite Validation"]
end
subgraph Dashboard["📊 Django Dashboard :8000"]
DJANGO["Django 6 + Daphne ASGI<br/>Web UI / Admin"]
POLL["Live Status Polling API<br/>/api/v1/logs-stream/"]
UI["Execution Logs Cards<br/>Auto-Refresh Every 5s"]
end
subgraph Storage["💾 Database Layer"]
PG[("PostgreSQL<br/>Multi-Tenant")]
RLS["🔐 Row-Level Security<br
[truncated]
Service
Technology
Role
ingestion-service
Node.js + Express
GitHub webhook receiver, path filtering, task queuing
core-brain
FastAPI + Celery
AI orchestration, secret scrubbing, Docker sandbox control
django-dashboard
Django 6 + Daphne ASGI
Web UI, audit logs, multi-tenant admin, billing
sandbox-env
Docker (network-isolated)
Pre-baked Pytest/Jest images, no network, 512MB RAM cap
billing-collector
Python
Cost forecasting, usage metering, AWS/GCP poll
redis-broker
Redis 7
Celery message queue + result backend
Security
Layer
Mechanism
Code storage
Zero persistence — PostgreSQL stores only operational metadata, never source code
LLM privacy
data_collection: deny on every OpenRouter request — legally blocks training on your code
Secret scrubbing
In-memory regex — AWS keys, GH tokens, DB credentials masked before transit
Patch execution
Network-isolated Docker — no network access, 512MB RAM / 2 CPU hard limit, no host FS mount
Tenant isolation
PostgreSQL RLS — database-enforced row separation, bypasses Django .filter()
Container leaks
Background cron — auto-prunes orphaned sandbox containers on execution freeze
About
Autonomous AI DevOps pipeline with air-gapped sandboxing, zero-data-retention LLM routing, PostgreSQL RLS multi-tenancy, and pluggable billing. Self-hosted, open-source, one docker-compose command.
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

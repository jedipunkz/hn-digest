---
source: "https://devopness.com/blog/deploy-hermes-ai-agent-devopness/"
hn_url: "https://news.ycombinator.com/item?id=48564127"
title: "How to Deploy Hermes AI Agent with Docker and HTTPS"
article_title: "How to Deploy Hermes AI Agent with Docker and HTTPS"
author: "Diegiwg"
captured_at: "2026-06-17T01:02:09Z"
capture_tool: "hn-digest"
hn_id: 48564127
score: 1
comments: 0
posted_at: "2026-06-17T00:14:36Z"
tags:
  - hacker-news
  - translated
---

# How to Deploy Hermes AI Agent with Docker and HTTPS

- HN: [48564127](https://news.ycombinator.com/item?id=48564127)
- Source: [devopness.com](https://devopness.com/blog/deploy-hermes-ai-agent-devopness/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T00:14:36Z

## Translation

タイトル: Docker と HTTPS を使用して Hermes AI エージェントをデプロイする方法
説明: Docker、HTTPS、および認証を使用して、Hermes AI エージェントをデプロイします。自動 SSL 証明書と永続ストレージを備えた完全なガイド。

記事本文:
Docker と HTTPS を使用して Hermes AI エージェントをデプロイする方法のドキュメント
ブログに戻る 2026 年 6 月 11 日 15 分で読む Devopness チーム Docker と HTTPS を使用して Hermes AI エージェントをデプロイする方法
Docker、HTTPS、認証を使用して、Hermes AI Agent をデプロイします。自動 SSL 証明書と永続ストレージを備えた完全なガイド。
Docker、自動 HTTPS、および認証を使用して、Hermes AI Agent をデプロイします。このガイドは、AI アシスタント (Devopness MCP サーバー、Devopness API など) と Web インターフェイスを介した手動セットアップの両方で機能します。
Hermes は、コードの作成、アプリケーションのデプロイ、Web の閲覧、Slack や GitHub との統合、反復的なタスクの自動化ができる自律型 AI エージェントです。チームに開発者がいるのと同じように機能します。
通常、AI エージェントを展開するには、複雑なサーバー構成、Docker の専門知識、および手動の SSL 証明書管理が必要です。このガイドでは、すべてのインフラストラクチャを自動的に処理する Devopness を使用して、Hermes をデプロイする方法を説明します。
このガイドを終えると、Hermes AI Agent がデプロイされ、実行されるようになります。
クラウドインフラストラクチャ上で年中無休で稼働
安全な HTTPS ダッシュボード ( https://hermes.example.com ) 経由でアクセス可能
信頼できる/プライベート アクセスのためのユーザー名/パスワード認証によって保護され、パブリック インターネット展開には OAuth/OIDC が推奨されます
導入後も存続する永続ストレージを使用
SSL証明書の自動更新（90日ごと）
Slack、GitHub、その他のサービスと統合する準備ができています
このガイドの対象者: 手動構成を行わずに、Hermes をクラウド サーバーに展開したい人。あらゆる AI アシスタントまたは Devopness Web UI 経由で動作します。
Devopness アカウントをまだお持ちでない場合は、始めるには、無料アカウントを作成してください。
Devopness MCP を使用して、AI アシスタント (Claude Code、Cursor、Cline、Windsurf など) にこの展開全体を実行させることができます。

サーバー 。
このプロンプトを AI アシスタントにコピーします。
https://www.devopness.com/blog/deploy-hermes-ai-agent-devopness で Herme 展開ガイドを読み、Devopness MCP ツールを使用して段階的に実行します。
AI アシスタントが展開を処理しながら、重要な意思決定ポイントで情報を提供します。
認証情報が準備された環境
Devopness では、「環境」(開発、ステージング、または運用) 内で作業します。以下の環境が必要です。
クラウド プロバイダーの資格情報 - Devopness がサーバーを作成できるようにします (AWS、Azure、Google Cloud、DigitalOcean、Hetzner、または任意のクラウド プロバイダー)
GitHub 資格情報 - Devopness が Hermes のコード (GitHub Personal Access Token) をダウンロードできるようにする
LLM プロバイダー アカウント - したがって、Hermes は AI (OpenRouter、OpenAI、Anthropic、Groq、または Ollama Cloud) を使用できます。
セットアップのサポートが必要ですか?最初にこのガイドに従ってください: 環境を作成する方法
次の仕様でサーバーを作成します。
クラウドプロバイダー: AWS、Azure、GCP、DigitalOcean、Hetzner、またはその他
地域 : ユーザーに最も近いものを選択してください
インスタンス サイズ : 最小 4GB RAM、20GB ディスク
オペレーティング システム : 最新の Ubuntu LTS
この仕様の理由:Hermes では、ブラウザーの自動化と AI タスクに 4GB RAM が必要です。 Ubuntu LTS は長期的なセキュリティ更新プログラムを提供します。
ステップ 2: アプリケーションを作成する
次の設定を使用してアプリケーションを作成します。
名前 : hermes-agent (小文字、スペースなし)
リポジトリ : NousResearch/hermes-agent
docker compose -p ${APPLICATION_NAME} pull && docker compose -p ${APPLICATION_NAME} up -d && docker compose -p ${APPLICATION_NAME} ログ || { docker compose -p ${APPLICATION_NAME} ログ;出口1; }
ステップ 3: ストレージ パス変数を追加する
永続ストレージの環境変数を作成します。
値: /home/devopness/hermes-agent (hermes-agent をアプリ名に置き換えます)
対象：OS環境変数
理由: 場所を定義します。

Hermes は、チャット履歴、API キー、ブラウザ セッション、アップロードされたファイルを保存します。これを行わないと、すべての展開が工場出荷時のデフォルトにリセットされます。
ステップ 4: 構成ファイルを追加する
3 つの構成ファイルを作成します。
次の値を準備し、以下のファイルの内容に置き換えます。
HERMES_DASHBOARD_BASIC_AUTH_PASSWORD - 強力なパスワードを使用します (20 文字以上)
HERMES_DASHBOARD_BASIC_AUTH_SECRET - 32 文字以上のランダムな文字列を使用します (https://generate-random.org/api-token-generator で生成するか、 openssl rand -base64 32 を実行します)。
説明: ダッシュボードの設定と認証
#Hermes AI エージェントの構成
# コンテナ名
CONTAINER_NAME=${APPLICATION_NAME}-エルメス
# ファイル権限 (展開中に自動検出)
エルメス_UID=__DEVOPNESS_UID__
エルメス_GID=__DEVOPNESS_GID__
# ダッシュボード設定
エルメス_ダッシュボード=true
エルメス_ダッシュボード_ホスト=0.0.0.0
エルメス_ダッシュボード_ポート=9119
# セキュリティ - ログイン保護
エルメス_DASHBOARD_BASIC_AUTH_USERNAME=エージェント
エルメス_DASHBOARD_BASIC_AUTH_PASSWORD=CHANGE_THIS_PASSWORD
エルメス_DASHBOARD_BASIC_AUTH_SECRET=CHANGE_THIS_SECRET
# パフォーマンス設定
エルメス_SHM_SIZE=1g
エルメス_MEMORY_LIMIT=4G
エルメス_MEMORY_RESERVATION=2G
エルメス_CPU_LIMIT=2.0
ファイル 2: .env.hermes
説明: LLM プロバイダーの API キーと統合
# ヘルメスエージェントの設定
# 端末環境（必須）
TERMINAL_ENV=ローカル
# ここに統合 API キーを追加します (例):
# SLACK_BOT_TOKEN=xoxb-...
# SLACK_APP_TOKEN=xapp-...
# LINEAR_API_KEY=lin_api_...
# OPENROUTER_API_KEY=sk-or-...
# OPENAI_API_KEY=sk-...
🔸 注: このファイルは環境変数の信頼できる情報源であり、デプロイメントごとに永続ストレージにコピーされます。 Herme ダッシュボードまたは ${APPLICATION_PATH}/data/.env 内で直接行われた変更は、次回のデプロイ時に上書きされます。有効期間の長い API キーとインテルを追加する

シークレットを .env.hermes に保存して、再デプロイします。
説明: コンテナの構成とリソース
サービス:
エルメス：
画像: nousresearch/hermes-agent:v2026.6.5
コンテナ名: ${CONTAINER_NAME}
再起動: 停止しない限り
コマンド: ゲートウェイの実行
特権: true
ボリューム:
- ${APPLICATION_PATH}/data:/opt/data
- /var/run/docker.sock:/var/run/docker.sock
環境ファイル:
- .env
ポート:
- "127.0.0.1:${エルメス_ダッシュボード_ポート}:${エルメス_ダッシュボード_ポート}"
shm_size: ${エルメス_SHM_SIZE:-1g}
デプロイ:
リソース:
制限:
メモリ: ${エルメス_MEMORY_LIMIT:-4G}
CPU: "${エルメス_CPU_LIMIT:-2.0}"
予約:
メモリ: ${HERMES_MEMORY_RESERVATION:-2G}
ネットワーク:
- エルメスネット
ネットワーク:
エルメスネット:
ドライバー：ブリッジ
名前: ${APPLICATION_NAME}-ネットワーク
🔹 ヒント: 新しいバージョンについては Docker Hub を確認してください。再現可能なデプロイメントのために、常に特定のバージョン タグ (latest ではない) にピン留めします。
🔸 セキュリティ上の注意: この構成では、privileged: true を使用し、 /var/run/docker.sock をマウントします。 Hermes が Docker コマンドとブラウザー自動化を実行するために必要ですが、コンテナーにホストへの昇格されたアクセスを与えます。 HTTPS と強力な認証を使用して展開します。
ステップ 5: デプロイメント パイプラインを構成する
デプロイメント パイプライン設定を更新します。
Max Parallel Actions : 1 に設定します (古い構成が新しい構成を上書きする可能性があるデプロイメントの競合を防ぎます)
3 つのカスタム パイプライン ステップを次の順序で追加します。
位置: 「Git リポジトリからソースを取得」の後
mkdir -p ${APPLICATION_PATH}/data && \
chown devopness:devopness ${APPLICATION_PATH}/data
ステップ 2: UID/GID プレースホルダーを更新する
位置: 「Hermes ストレージの作成」後
名前: UID/GID プレースホルダーの更新
DEVOPNESS_UID=$(id -u devopness) && \
DEVOPNESS_GID=$(id -g devopness) && \
sed -i "s/__DEVOPNESS_UID__/${DEVOPNESS_UID}/g" .env && \
sed -i "s/__DEVOPNESS_GID__/${DEVOPNESS_GID}/g" .env
理由: 一致する

ホスト ファイル システムに対する ntainer ユーザー権限。
ステップ3：Hermes設定をコピーする
位置: 「現在のビルドをアクティブ化」の後
名前:Hermes 設定をデータ フォルダーにコピーします。
cp .env.hermes ${APPLICATION_PATH}/data/.env && \
chmod 600 ${APPLICATION_PATH}/data/.env && \
chown devopness:devopness ${APPLICATION_PATH}/data/.env
理由: API キーを永続ストレージにコピーして、デプロイ後も存続できるようにします。
パイプライン: パイプラインをデプロイします (構成済み)
サーバー: ステップ 1 で作成したサーバーを選択します。
デプロイメントには、最初の実行時に 5 ～ 8 分かかります (Docker イメージのダウンロード、コンテナーの構築)。
Devopness は、展開をライブで見ることができる [アクション] ページにリダイレクトします。アクションが完了するまで待ちます。
ステップ 7: 導入をテストする (オプション)
🔹 ヒント: ドメインの準備ができている場合は、ステップ 8 の HTTPS セットアップに進みます。
サーバー用に自動生成された仮想ホストを検索して編集します (ホスト名はサーバーのパブリック IP アドレスと一致します)。
アプリケーション:Hermes アプリケーションを選択してください
アプリケーションリッスンアドレス: http://localhost:9119
デプロイして、アクションが完了するまで待ちます。次に、 http://YOUR_SERVER_IP でアクセスをテストします。 hermes ログイン ページが表示されるはずです (ユーザー名: Agent 、パスワード: .env ファイルからのもの)。
認証セキュリティに関する通知
このガイドでは、信頼されたネットワークまたは VPN アクセスに適した HTTP 基本認証 (ユーザー名/パスワード) を使用します。
強力なパスワードを使用してください (20 文字以上のランダムな文字)
セキュリティを強化するために IP ホワイトリストを検討してください
パブリック インターネットの展開の場合は、OAuth (例: Nous Portal) またはセルフホスト型 OIDC (例: Keycloak、Auth0、Okta) を使用します。詳細な手順については、Hermes のドキュメントを参照してください。
ステップ 8: カスタム ドメインで HTTPS をセットアップする (オプション)
ドメインがサーバーの IP アドレス (A レコード) を指すようにします。
名前ベースの仮想ホストを作成します。
アプリケーション: アプリケーションを選択してください
アプリケーションリッスンアドレス: http:/

/ローカルホスト:9119
仮想ホスト: 上記で作成した仮想ホストを選択します
認証局: Let's Encrypt
https://hermes.example.com へのアクセスをテストします - 有効な SSL 証明書を含む Hermes ログイン ページが表示されるはずです。
ステップ9：Hermesの設定とテスト
LLM プロバイダー API キーを .env.hermes に追加します。
OPENROUTER_API_KEY=sk-or-v1-...
# または OPENAI_API_KEY、ANTHROPIC_API_KEY など
アプリケーションを再デプロイします。
「モデル」に移動 → プロバイダーとモデルを選択します
「チャット」に移動 → テストメッセージを送信
完了！ Hermes AI Agent がデプロイされ、使用できるようになりました。
AI アシスタントが実行され、タスクを自動化する準備が整いました。
私たちのブログ投稿をさらにチェックして、最新の洞察を常に入手してください。
DevOps Happiness 、出荷するビルダー向け。

## Original Extract

Deploy Hermes AI Agent with Docker, HTTPS, and authentication. Complete guide with automatic SSL certificates and persistent storage.

How to Deploy Hermes AI Agent with Docker and HTTPS Docs
Back to Blog June 11, 2026 15 min read Devopness Team How to Deploy Hermes AI Agent with Docker and HTTPS
Deploy Hermes AI Agent with Docker, HTTPS, and authentication. Complete guide with automatic SSL certificates and persistent storage.
Deploy Hermes AI Agent with Docker, automatic HTTPS, and authentication. This guide works for both AI assistants (e.g., via Devopness MCP Server, Devopness API) and manual setup through the web interface.
Hermes is an autonomous AI agent that can write code, deploy applications, browse the web, integrate with Slack and GitHub, and automate repetitive tasks. It works like having a developer on your team.
Deploying AI agents typically requires complex server configuration, Docker expertise, and manual SSL certificate management. This guide shows you how to deploy Hermes using Devopness, which handles all the infrastructure automatically.
By the end of this guide, you'll have Hermes AI Agent deployed and running:
Running 24/7 on your cloud infrastructure
Accessible via secure HTTPS dashboard ( https://hermes.example.com )
Protected by username/password authentication for trusted/private access, with OAuth/OIDC recommended for public internet deployments
With persistent storage that survives deployments
Auto-renewing SSL certificates (every 90 days)
Ready to integrate with Slack, GitHub, and other services
This guide is for: Anyone who wants to deploy Hermes to a cloud server, without manual configuration. Works with any AI assistant or via Devopness web UI.
Don't have a Devopness account yet? Create your free account to get started.
You can have an AI assistant (e.g., Claude Code, Cursor, Cline, Windsurf) execute this entire deployment for you using the Devopness MCP Server .
Copy this prompt to your AI assistant:
Read the Hermes deployment guide at https://www.devopness.com/blog/deploy-hermes-ai-agent-devopness and execute it step by step using Devopness MCP tools.
Your AI assistant will handle the deployment while keeping you informed at key decision points.
An Environment with Credentials Ready
In Devopness, you work inside "Environments" (Development, Staging, or Production). You'll need an environment with:
A Cloud Provider Credential - So Devopness can create servers for you (AWS, Azure, Google Cloud, DigitalOcean, Hetzner, or any cloud provider)
A GitHub Credential - So Devopness can download Hermes' code (GitHub Personal Access Token)
An LLM Provider Account - So Hermes can use AI (OpenRouter, OpenAI, Anthropic, Groq, or Ollama Cloud)
Need help setting up? Follow this guide first: How to Create an Environment
Create a server with these specifications:
Cloud Provider : AWS, Azure, GCP, DigitalOcean, Hetzner, or others
Region : Choose the closest to your users
Instance Size : 4GB RAM minimum, 20GB disk
Operating System : Latest Ubuntu LTS
Why these specs: Hermes requires 4GB RAM for browser automation and AI tasks. Ubuntu LTS provides long-term security updates.
Step 2: Create the Application
Create an application with these settings:
Name : hermes-agent (lowercase, no spaces)
Repository : NousResearch/hermes-agent
docker compose -p ${APPLICATION_NAME} pull && docker compose -p ${APPLICATION_NAME} up -d && docker compose -p ${APPLICATION_NAME} logs || { docker compose -p ${APPLICATION_NAME} logs; exit 1; }
Step 3: Add the Storage Path Variable
Create an environment variable for persistent storage:
Value : /home/devopness/hermes-agent (replace hermes-agent with your app name)
Target : OS Environment Variable
Why: Defines where Hermes stores chat history, API keys, browser sessions, and uploaded files. Without this, every deployment resets to factory defaults.
Step 4: Add Configuration Files
Create three configuration files.
Prepare the following values and replace them in the file content below:
HERMES_DASHBOARD_BASIC_AUTH_PASSWORD - use a strong password (20+ characters)
HERMES_DASHBOARD_BASIC_AUTH_SECRET - use a random 32+ character string (generate at https://generate-random.org/api-token-generator or run openssl rand -base64 32 )
Description: Dashboard settings and authentication
# Hermes AI Agent Configuration
# Container Name
CONTAINER_NAME=${APPLICATION_NAME}-hermes
# File Permissions (auto-detected during deployment)
HERMES_UID=__DEVOPNESS_UID__
HERMES_GID=__DEVOPNESS_GID__
# Dashboard Settings
HERMES_DASHBOARD=true
HERMES_DASHBOARD_HOST=0.0.0.0
HERMES_DASHBOARD_PORT=9119
# Security - Login Protection
HERMES_DASHBOARD_BASIC_AUTH_USERNAME=agent
HERMES_DASHBOARD_BASIC_AUTH_PASSWORD=CHANGE_THIS_PASSWORD
HERMES_DASHBOARD_BASIC_AUTH_SECRET=CHANGE_THIS_SECRET
# Performance Settings
HERMES_SHM_SIZE=1g
HERMES_MEMORY_LIMIT=4G
HERMES_MEMORY_RESERVATION=2G
HERMES_CPU_LIMIT=2.0
File 2: .env.hermes
Description: LLM provider API keys and integrations
# Hermes Agent Configuration
# Terminal Environment (required)
TERMINAL_ENV=local
# Add your integration API keys here (examples):
# SLACK_BOT_TOKEN=xoxb-...
# SLACK_APP_TOKEN=xapp-...
# LINEAR_API_KEY=lin_api_...
# OPENROUTER_API_KEY=sk-or-...
# OPENAI_API_KEY=sk-...
🔸 Note: This file is the source of truth for environment variables and is copied to persistent storage on every deployment. Changes made directly in the Hermes dashboard or inside ${APPLICATION_PATH}/data/.env will be overwritten on the next deployment. Add long-lived API keys and integration secrets to .env.hermes , then redeploy.
Description: Container configuration and resources
services:
hermes:
image: nousresearch/hermes-agent:v2026.6.5
container_name: ${CONTAINER_NAME}
restart: unless-stopped
command: gateway run
privileged: true
volumes:
- ${APPLICATION_PATH}/data:/opt/data
- /var/run/docker.sock:/var/run/docker.sock
env_file:
- .env
ports:
- "127.0.0.1:${HERMES_DASHBOARD_PORT}:${HERMES_DASHBOARD_PORT}"
shm_size: ${HERMES_SHM_SIZE:-1g}
deploy:
resources:
limits:
memory: ${HERMES_MEMORY_LIMIT:-4G}
cpus: "${HERMES_CPU_LIMIT:-2.0}"
reservations:
memory: ${HERMES_MEMORY_RESERVATION:-2G}
networks:
- hermes-net
networks:
hermes-net:
driver: bridge
name: ${APPLICATION_NAME}-network
🔹 Tip: Check Docker Hub for newer versions. Always pin to a specific version tag (not latest ) for reproducible deployments.
🔸 Security note: This configuration uses privileged: true and mounts /var/run/docker.sock . Required for Hermes to run Docker commands and browser automation, but gives the container elevated access to the host. Deploy with HTTPS and strong authentication.
Step 5: Configure the Deployment Pipeline
Update your deployment pipeline settings:
Max Parallel Actions : Set to 1 (prevents deployment conflicts where outdated configurations could override newer ones)
Add three custom pipeline steps in this order:
Position: After "Get source from Git repository"
mkdir -p ${APPLICATION_PATH}/data && \
chown devopness:devopness ${APPLICATION_PATH}/data
Step 2: Update UID/GID Placeholders
Position: After "Create Hermes storage"
Name: Update UID/GID placeholders
DEVOPNESS_UID=$(id -u devopness) && \
DEVOPNESS_GID=$(id -g devopness) && \
sed -i "s/__DEVOPNESS_UID__/${DEVOPNESS_UID}/g" .env && \
sed -i "s/__DEVOPNESS_GID__/${DEVOPNESS_GID}/g" .env
Why: Aligns container user permissions with host filesystem.
Step 3: Copy Hermes Configuration
Position: After "Activate current build"
Name: Copy Hermes configuration to data folder
cp .env.hermes ${APPLICATION_PATH}/data/.env && \
chmod 600 ${APPLICATION_PATH}/data/.env && \
chown devopness:devopness ${APPLICATION_PATH}/data/.env
Why: Copies API keys to persistent storage so they survive deployments.
Pipeline: Deploy pipeline (already configured)
Servers: Select the server created in Step 1
Deployment takes 5-8 minutes on first run (downloads Docker images, builds the container).
Devopness redirects you to the Action page where you can watch the deployment live. Wait for the action to complete.
Step 7: Test Your Deployment (Optional)
🔹 Tip: If you have a domain ready, skip to Step 8 for HTTPS setup.
Find and edit the auto-generated virtual host for your server (hostname matches your server's public IP address):
Application: Select your Hermes application
Application Listen Address: http://localhost:9119
Deploy and wait for the action to complete. Then test access at http://YOUR_SERVER_IP . You should see the Hermes login page (username: agent , password: from your .env file).
Authentication Security Notice
This guide uses HTTP Basic Authentication (username/password), suitable for trusted networks or VPN access.
Use a strong password (20+ random characters)
Consider IP whitelisting for additional security
For public internet deployments, use OAuth (e.g., Nous Portal) or self-hosted OIDC (e.g., Keycloak, Auth0, Okta). For further instructions, see Hermes documentation .
Step 8: Setup HTTPS with Custom Domain (Optional)
Point your domain to your server's IP address (A record).
Create a name-based virtual host:
Application: Select your application
Application Listen Address: http://localhost:9119
Virtual Host: Select the virtual host created above
Certificate Authority: Let's Encrypt
Test access to https://hermes.example.com - you should see the Hermes login page with a valid SSL certificate.
Step 9: Configure and Test Hermes
Add your LLM provider API key to .env.hermes :
OPENROUTER_API_KEY=sk-or-v1-...
# or OPENAI_API_KEY, ANTHROPIC_API_KEY, etc.
Redeploy the application.
Go to Models → select your provider and model
Go to Chat → send a test message
Done! Your Hermes AI Agent is deployed and ready to use.
Your AI assistant is now running and ready to automate tasks.
Check out more of our blog posts and stay updated with the latest insights.
DevOps Happiness , for builders who ship.

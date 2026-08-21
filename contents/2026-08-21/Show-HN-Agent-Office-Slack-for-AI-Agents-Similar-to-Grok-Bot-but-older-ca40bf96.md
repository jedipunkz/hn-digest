---
source: "https://github.com/baturyilmaz/agent-office"
hn_url: "https://news.ycombinator.com/item?id=49387504"
title: "Show HN: Agent Office (Slack for AI Agents) – Similar to Grok Bot but older"
article_title: "GitHub - baturyilmaz/agent-office: Slack for AI Agents · GitHub"
image: "https://opengraph.githubassets.com/8c5937e660b01daf8d261ad8c74677362ecdd8265800944b301e89e31747a9ec/baturyilmaz/agent-office"
author: "arbayi"
captured_at: "2026-08-21T13:37:47Z"
capture_tool: "hn-digest"
hn_id: 49387504
score: 2
comments: 0
posted_at: "2026-08-21T13:04:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent Office (Slack for AI Agents) – Similar to Grok Bot but older

- HN: [49387504](https://news.ycombinator.com/item?id=49387504)
- Source: [github.com](https://github.com/baturyilmaz/agent-office)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T13:04:59Z

## Translation

タイトル: HN を表示: エージェント オフィス (AI エージェント用の Slack) – Grok Bot に似ていますが、古いものです
記事のタイトル: GitHub - buturyilmaz/agent-office: AI エージェントのための Slack · GitHub
説明: AI エージェント用の Slack。 GitHub でアカウントを作成して、baturyilmaz/agent-office の開発に貢献してください。

記事本文:
GitHub - buturyilmaz/agent-office: AI エージェント用の Slack · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
バトゥリルマズ
/
代理店
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
177 コミット 177 コミット フォルダーとファイル
docs/ 画像 docs/ 画像 例 例 src src test test ui ui .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc README.md RE

ADME.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pi 上に構築されたマルチエージェント ワークスペース マネージャー。 Claude Code や OpenClaw と同様に、ティックベースのスケジューリング、優先キュー、受信ボックス IPC、エージェント間ファイル アクセス、ウォッチドッグ モニタリング、プロアクティブな cron ジョブ、オプションの Docker サンドボックス分離、宣言型 YAML 構成を使用して、AI コーディング エージェントを調整します。
すぐに起動して実行するには、次の例のいずれかを試してください。プロジェクトのルート .env に環境変数を設定します (Docker 内ではなく、ホストが環境変数をコンテナーに転送します)。
基本チーム — PM、コーダー、およびレビュー担当者 (GitHub Copilot OAuth を使用、API キーは使用しません):
pnpmインストール
cp -r 例/basic-team/ ~ /.agent-office/offices/basic-team/
pnpm dev oauth ログイン github-copilot --office Basic-team
pnpm dev start --office Basic-team --sandbox docker
OpenServ チーム — アイデア スカウト、チーム リーダー、エージェント開発者、およびトークン ランチャー:
pnpmインストール
cp .env.example .env
mkdir -p ~ /.agent-office/offices/openserv-team
cp 例/openserv-team/office.yaml ~ /.agent-office/offices/openserv-team/office.yaml
pnpm dev start --office openserv-team --sandbox docker
OPENAI_API_KEY =
WALLET_PRIVATE_KEY = # openserv-labs/スキルエージェントのEVMウォレットキー
フィーチャー チーム — カンバン ボードを使用したタスク駆動型開発:
cp -r 例/feature-team/ ~ /.agent-office/offices/feature-team/
pnpm dev start --office feature-team
詳細については、examples/ を参照してください。それぞれのセットアップについて説明した README があります。
OAuth認証
サポートされているプロバイダー
マルチオフィスアーキテクチャ
オフィスを作る
グラフTD
YAML[office.yaml] --> WS
CLI[CLI + Web UI] --> WS[ワークスペース]
WS --> SCH[スケジューラ\nループ]
WS --> バス[MessageBus\ninboxes]
WS --> WD[ウォッチドッグ\nハートビート]
WS --> CRON[CronServ

\nスケジュールされたジョブを氷]
WS --> TS[TaskService\nカンバンボード]
WS -->|処理中| A[エージェント A\nPi · ツール · スキル]
WS -->|処理中| B[エージェント B\nPi · ツール · スキル]
WS -->|Docker サンドボックス| HA[ホスト API\nHTTP :13000]
HA <-->|HTTP| SA[サンドボックス A\nDocker · Pi · プロキシ ツール]
HA <-->|HTTP| SB[サンドボックス B\nDocker · Pi · プロキシ ツール]
バス --> A
バス --> B
バス --> ハ
読み込み中
コア フロー: office.yaml (自動生成) / CLI / Web UI / Cron / エージェント cron ツール / タスク通知 -> ワークスペース -> スケジューラ ティック -> 受信トレイのドレイン -> Pi エージェントへのディスパッチ -> エージェントがツールを実行 -> 応答が UI にストリーミングされます。
各エージェントは、独自のファイルシステム ワークスペース、スキル、挿入されたツール ( message_user 、 post_channel 、 message_agent 、 list_agents 、 read_agent_file 、authenticated_fetch 、 cron_add 、 cron_remove 、 cron_list 、 task_create 、 task_update 、 task_list 、 task_get 、 task_delete 、 read_skill 、 skill_search 、 skill_install 、 skill_remove 、 skill_create )。スケジューラは、優先順位に従ってエージェントにサービスを提供するティック ループを実行し、エージェントごとに 1 つのティックごとに 1 つのメッセージを非ブロックで提供します。
エージェントはプロセス内 (デフォルト) で実行することも、完全なプロセスレベルで分離するために Docker コンテナ内で実行することもできます。
pnpmインストール
# .env を設定する
cp .env.example .env # 次にキーを入力します
#オフィスを作る
pnpm dev officeは私のチームを作成します
# 起動（Web UI自動起動）
pnpm dev start --office my-team
# Docker サンドボックス分離から始める
pnpm dev start --office my-team --sandbox docker
プロバイダー API キーを使用して .env ファイルを作成します。各モデルには、対応するプロバイダー キーが必要です。
# モデル API キー (これらのモデルを使用するエージェントに必要)
OPENAI_API_KEY = sk-... # OpenAI モデルの場合 (gpt-4o など)
ANTHROPIC_API_KEY = sk-... # Anthropic モデルの場合 (Claude など)
GEMINI_API_KEY = ... # Google Gemini モデルの場合
XAI_API_KEY = ... # xAI G の場合

ロックモデル
# オプション: office.yaml エージェントのカスタム シークレット参照
# MY_GH_TOKEN=ghp_... #Authenticated_fetch シークレットのホスト環境変数
認証: 各モデルには認証情報が必要です。 API キー ( .env ) または OAuth のいずれかを使用できます。
API キー — .env で設定します (例: OPENAI_API_KEY=sk-... )。モデルのプロバイダーに OAuth 資格情報がない場合に必要です。
OAuth — 開始する前にプロバイダー CLI 経由で認証します。 OAuth トークンは自動更新され、.env キーは必要ありません。
# オプション A: .env 内の API キー
echo " ANTHROPIC_API_KEY=sk-... " >> .env
# オプション B: OAuth ログイン (プロバイダー CLI がインストールされている必要があります)
pnpm dev oauth ログイン anthropic --office my-team
pnpm dev oauth list --office my-team
OAuth 資格情報と API キーの両方がプロバイダーに存在する場合、Web UI でエージェントごとにそれらを切り替えることができます。詳細については、「OAuth 認証」を参照してください。
Web UI で利用可能なモデルとその要件を参照する方法については、以下の「動的モデル検出」セクションを参照してください。
.env の API キーの代わりに、エージェントは OAuth 経由でモデル プロバイダーで認証できます。これは、プロバイダー独自の CLI ログイン フローを使用します。エージェントとオフィスの CLI はブラウザベースの OAuth ハンドシェイクを調整し、オフィスごとに資格情報を保存します。
プロバイダーID
名前
フロータイプ
必要なもの
人間的な
人間的
コードペースト
人間の CLI
オープンアイコーデックス
OpenAI
コールバックサーバー
OpenAI コーデックス CLI
github-副操縦士
GitHub コパイロット
コードペースト
GitHub コパイロット CLI
google-gemini-cli
Google Gemini CLI
コールバックサーバー
ジェミニ CLI
グーグル反重力
反重力
コールバックサーバー
反重力 CLI
コード貼り付けプロバイダーはブラウザの URL を開き、認証コードを貼り付けるように求めます。コールバック サーバー プロバイダーはローカル HTTP サーバーを起動し、フローを自動的に完了します。
# ログイン — インタラクティブな OAuth フロー (ブラウザを開く)
pnpm dev oauth ログイン < プロバイダー > --office < id >
# リスト — sh

すべてのプロバイダーと資格情報のステータスを確認する
pnpm dev oauth list --office < id >
# ログアウト — 保存されている認証情報を削除します
pnpm dev oauth logout <プロバイダー> --office <id>
セッションの例:
$ pnpm dev oauth ログイン anthropic --office my-team
[oauth] Anthropic にログインしています...
[oauth] この URL を開いて認証します。
https://console.anthropic.com/oauth/...
認証コードを貼り付けます: ****
[oauth] Anthropic 用に保存された認証情報。
$ pnpm dev oauth list --office my-team
✓ 人間的 人間的
✗ openai-codex OpenAI
✗ github-copilot GitHub コパイロット
✗ google-gemini-cli Google Gemini CLI
✗ google-antigravity 反重力
ログイン: pnpm dev oauth ログイン <プロバイダー> --office my-team
ログアウト: pnpm dev oauth logout <プロバイダー> --office my-team
資格情報は ~/.agent-office/offices/<id>/oauth/<provider>.json に保存され、トークンの有効期限が切れると自動更新されます。
API キーの代わりに OAuth を使用するようにエージェントの認証フィールドを設定します。
エージェント:
デザイナー：
モデル: anthropic:claude-sonnet-4-20250514
auth : " oauth:anthropic " # OAuth 認証情報を使用する
査読者：
モデル：openai：gpt-4o
auth : " oauth:openai-codex " # OAuth 認証情報を使用する
アナリスト：
モデル: google:gemini-2.0-flash
# 認証フィールドなし — .env から GEMINI_API_KEY にフォールバックします
認証フィールドの形式は oauth:<provider-id> です。設定すると、エージェントは静的 API キーの代わりに、自動トークン更新で保存された OAuth 認証情報を使用します。
エージェントのモデル プロバイダーの OAuth 資格情報が存在する場合、[Web UI Config] タブには、「API キー」モードと「OAuth」モードを切り替えるための認証セレクターが表示されます。また、UI には、認証されたすべてのプロバイダーが緑色のバッジとして表示され、ワンクリックで削除できます。
認証セレクターは、認証情報が利用可能な場合にのみ表示されます。プロバイダーに対して OAuth ログインが行われていない場合、エージェントはデフォルトで API キーを使用します。
方法
パス
説明
ゲット
/api/oauth/

プロバイダー
すべてのプロバイダーを認証ステータスとともにリストします
ゲット
/api/oauth/ステータス/:id
プロバイダーの資格情報が存在するかどうかを確認する
削除
/api/oauth/:id
保存されているプロバイダーの資格情報を削除する
マルチオフィスアーキテクチャ
各オフィスは、ID、環境変数、およびシークレットを共有する会社またはチームを表します。オフィスは ~/.agent-office/offices/<id>/ の下に存在します。
# デフォルトの表示名（IDと同じ）で作成します
pnpm dev office create acme
# カスタム表示名で作成する
pnpm dev office create acme --name " Acme Corp "
Office ID はパスセーフである必要があります: 小文字、数字、ハイフン、アンダースコア ( [a-z0-9][a-z0-9_-]* に一致)。表示名 (YAML では office.name) は自由形式です。
Office 構成 (office.yaml)
~/.agent-office/offices/<id>/office.yaml でオフィスを定義すると、エージェントは起動時に自動生成されます。
# ~/.agent-office/offices/acme/office.yaml
事務所：
名前 : アクメコーポレーション
説明 : 「 AI を活用したウィジェットを構築します 」
環境:
SHARED_API_URL : https://api.acme.com
秘密:
SHARED_TOKEN : ${ACME_TOKEN}
クロン :
スタンドアップ：
スケジュール：「0 9 * * 1-5」
レポートチャンネル : 一般
タスク:
- タイトル：「デイリースタンドアップ」
譲受人 : 午後
エージェント:
デザイナー：
モデル: anthropic:claude-sonnet-4-20250514
優先順位 : 通常 # アイドル |低い |通常 |高い |クリティカル (または 0 ～ 4)
思考: 低い # オフ |最小限 |低い |中 |高い |高さ
説明 : " フロントエンド デザイナー — HTML/CSS を構築します "
プロンプトインライン : |
あなたはレスポンシブ レイアウトを専門とするフロントエンド デザイナーです。
クリーンでセマンティックな HTML と最新の CSS に焦点を当てます。
スキル：
- ニコキャラ/Webスキル
auth : " oauth:anthropic " # オプション — API キーの代わりに OAuth を使用します
api_key_ref : MY_CUSTOM_KEY # オプション — モデル キー オーバーライドのホスト環境変数名
env : # 非機密、Docker --env として渡されます (エージェントが office をオーバーライドします)
LOG_LEVEL : デバッグ
WORKSPACE_NAME :

デザイナー
Secrets : # 機密、${VAR} 参照のみ —authenticated_fetch 経由で配信されます
GITHUB_TOKEN : ${MY_GH_TOKEN}
Discover_secrets : true # システム プロンプトにシークレット名を表示します (デフォルト: false)
権限:
office_cron : true # オフィスレベルの cron ジョブの管理を許可します
査読者：
モデル：openai:gpt-4.1
優先度：高
思考：中程度
説明:「コードレビューアー」
オフィスレベルの環境とシークレットは、すべてのエージェントによって継承されます。エージェント レベルの値はオフィス レベルをオーバーライドします。
エージェントのフィールドはすべてオプションです。エージェントは宣言順に順次生成されます。 1 つが失敗しても、残りは引き続き開始されます。モデルの利用可能性はプロバイダー アカウントによって異なります。デフォルトが利用できない場合は、モデルの値を希望の Provider:model-id に置き換えます。
タスク ツール ( task_create 、 task_update 、 task_list 、 task_get 、 task_delete ) は、デフォルトですべてのインプロセス エージェントで使用できます。 Permissions.tools.deny を介してアクセスを制限します。 「タスク管理」を参照してください。
エージェントは、ハートビート (外部トリガーなしでエージェントに作業の確認やメンテナンスの実行を促す定期的なメッセージ) を介してプロアクティブに実行できます。
エージェント:
モニター：
心拍数:
間隔ミリ秒 : 60000
プロンプト:「保留中の作業を確認し、ステータスを報告する」
アクティブ時間:
開始 : " 09:00 "
終了：「17:00」
フィールド
必須
デフォルト
説明
インター

[切り捨てられた]

## Original Extract

Slack for AI Agents. Contribute to baturyilmaz/agent-office development by creating an account on GitHub.

GitHub - baturyilmaz/agent-office: Slack for AI Agents · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
baturyilmaz
/
agent-office
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
177 Commits 177 Commits Folders and files
docs/ images docs/ images examples examples src src test test ui ui .DS_Store .DS_Store .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc README.md README.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json View all files Repository files navigation
Multi-agent workspace manager built on Pi . Orchestrates AI coding agents — similar to Claude Code or OpenClaw — with tick-based scheduling, priority queues, inbox IPC, cross-agent file access, watchdog monitoring, proactive cron jobs, optional Docker sandbox isolation, and declarative YAML configuration.
Try one of these examples to get up and running quickly. Set env vars in the project root .env (not inside Docker — the host forwards them to containers).
Basic team — PM, coder, and reviewer (uses GitHub Copilot OAuth, no API keys):
pnpm install
cp -r examples/basic-team/ ~ /.agent-office/offices/basic-team/
pnpm dev oauth login github-copilot --office basic-team
pnpm dev start --office basic-team --sandbox docker
OpenServ team — idea scout, team lead, agent dev, and token launcher:
pnpm install
cp .env.example .env
mkdir -p ~ /.agent-office/offices/openserv-team
cp examples/openserv-team/office.yaml ~ /.agent-office/offices/openserv-team/office.yaml
pnpm dev start --office openserv-team --sandbox docker
OPENAI_API_KEY =
WALLET_PRIVATE_KEY = # EVM wallet key for openserv-labs/skills agents
Feature team — task-driven development with Kanban board:
cp -r examples/feature-team/ ~ /.agent-office/offices/feature-team/
pnpm dev start --office feature-team
See examples/ for more details — each has a README describing the setup.
OAuth Authentication
Supported Providers
Multi-Office Architecture
Creating an Office
graph TD
YAML[office.yaml] --> WS
CLI[CLI + Web UI] --> WS[Workspace]
WS --> SCH[Scheduler\ntick loop]
WS --> BUS[MessageBus\ninboxes]
WS --> WD[Watchdog\nheartbeat]
WS --> CRON[CronService\nscheduled jobs]
WS --> TS[TaskService\nKanban board]
WS -->|in-process| A[Agent A\nPi · tools · skills]
WS -->|in-process| B[Agent B\nPi · tools · skills]
WS -->|Docker sandbox| HA[Host API\nHTTP :13000]
HA <-->|HTTP| SA[Sandbox A\nDocker · Pi · proxy tools]
HA <-->|HTTP| SB[Sandbox B\nDocker · Pi · proxy tools]
BUS --> A
BUS --> B
BUS --> HA
Loading
Core flow: office.yaml (auto-spawn) / CLI / Web UI / Cron / Agent cron tools / Task notifications -> Workspace -> Scheduler tick -> drain inbox -> dispatch to Pi Agent -> agent runs tools -> response streamed to UI.
Each agent is a full Pi coding agent with its own filesystem workspace, skills, and injected tools ( message_user , post_channel , message_agent , list_agents , read_agent_file , authenticated_fetch , cron_add , cron_remove , cron_list , task_create , task_update , task_list , task_get , task_delete , read_skill , skill_search , skill_install , skill_remove , skill_create ). The scheduler runs a tick loop that serves agents by priority, one message per tick per agent, non-blocking.
Agents can run in-process (default) or inside Docker containers for full process-level isolation.
pnpm install
# Configure .env
cp .env.example .env # then fill in your keys
# Create an office
pnpm dev office create my-team
# Start (Web UI auto-starts)
pnpm dev start --office my-team
# Start with Docker sandbox isolation
pnpm dev start --office my-team --sandbox docker
Create a .env file with your provider API keys. Each model requires its corresponding provider key:
# Model API Keys (required for agents using these models)
OPENAI_API_KEY = sk-... # For OpenAI models (gpt-4o, etc.)
ANTHROPIC_API_KEY = sk-... # For Anthropic models (Claude, etc.)
GEMINI_API_KEY = ... # For Google Gemini models
XAI_API_KEY = ... # For xAI Grok models
# Optional: Custom secret refs for office.yaml agents
# MY_GH_TOKEN=ghp_... # Host env vars for authenticated_fetch secrets
Authentication: Each model needs credentials. You can use either API keys ( .env ) or OAuth:
API keys — set in .env (e.g. OPENAI_API_KEY=sk-... ). Required when the model's provider has no OAuth credentials.
OAuth — authenticate via provider CLIs before starting. OAuth tokens auto-refresh and don't require .env keys.
# Option A: API keys in .env
echo " ANTHROPIC_API_KEY=sk-... " >> .env
# Option B: OAuth login (requires provider CLI installed)
pnpm dev oauth login anthropic --office my-team
pnpm dev oauth list --office my-team
When both OAuth credentials and an API key exist for a provider, you can switch between them per-agent in the Web UI. See OAuth Authentication for details.
See the Dynamic Model Discovery section below for how to browse available models and their requirements in the Web UI.
As an alternative to API keys in .env , agents can authenticate with model providers via OAuth. This uses the provider's own CLI login flow — the agent-office CLI orchestrates the browser-based OAuth handshake and stores credentials per office.
Provider ID
Name
Flow Type
Requires
anthropic
Anthropic
Code paste
Anthropic CLI
openai-codex
OpenAI
Callback server
OpenAI Codex CLI
github-copilot
GitHub Copilot
Code paste
GitHub Copilot CLI
google-gemini-cli
Google Gemini CLI
Callback server
Gemini CLI
google-antigravity
Antigravity
Callback server
Antigravity CLI
Code paste providers open a browser URL and prompt you to paste back an auth code. Callback server providers start a local HTTP server and complete the flow automatically.
# Login — interactive OAuth flow (opens browser)
pnpm dev oauth login < provider > --office < id >
# List — show all providers and credential status
pnpm dev oauth list --office < id >
# Logout — remove stored credentials
pnpm dev oauth logout < provider > --office < id >
Example session:
$ pnpm dev oauth login anthropic --office my-team
[oauth] Logging in to Anthropic...
[oauth] Open this URL to authenticate:
https://console.anthropic.com/oauth/...
Paste the authorization code: ****
[oauth] Credentials saved for Anthropic.
$ pnpm dev oauth list --office my-team
✓ anthropic Anthropic
✗ openai-codex OpenAI
✗ github-copilot GitHub Copilot
✗ google-gemini-cli Google Gemini CLI
✗ google-antigravity Antigravity
Login: pnpm dev oauth login < provider > --office my-team
Logout: pnpm dev oauth logout < provider > --office my-team
Credentials are stored at ~/.agent-office/offices/<id>/oauth/<provider>.json and auto-refresh when tokens expire.
Set the auth field on an agent to use OAuth instead of an API key:
agents :
designer :
model : anthropic:claude-sonnet-4-20250514
auth : " oauth:anthropic " # use OAuth credentials
reviewer :
model : openai:gpt-4o
auth : " oauth:openai-codex " # use OAuth credentials
analyst :
model : google:gemini-2.0-flash
# no auth field — falls back to GEMINI_API_KEY from .env
The auth field format is oauth:<provider-id> . When set, the agent uses stored OAuth credentials with automatic token refresh instead of a static API key.
When OAuth credentials exist for an agent's model provider, the Web UI Config tab shows an Auth selector to switch between "API Key" and "OAuth" modes. The UI also displays all authenticated providers as green badges with one-click removal.
The auth selector only appears when credentials are available — if no OAuth login has been done for a provider, agents use API keys by default.
Method
Path
Description
GET
/api/oauth/providers
List all providers with authentication status
GET
/api/oauth/status/:id
Check if credentials exist for a provider
DELETE
/api/oauth/:id
Remove stored credentials for a provider
Multi-Office Architecture
Each office represents a company or team with shared identity, env vars, and secrets. Offices live under ~/.agent-office/offices/<id>/ .
# Create with default display name (same as id)
pnpm dev office create acme
# Create with a custom display name
pnpm dev office create acme --name " Acme Corp "
Office IDs must be path-safe: lowercase letters, digits, hyphens, underscores (matching [a-z0-9][a-z0-9_-]* ). The display name ( office.name in YAML) is free-form.
Office Configuration ( office.yaml )
Define your office once in ~/.agent-office/offices/<id>/office.yaml and agents auto-spawn on startup.
# ~/.agent-office/offices/acme/office.yaml
office :
name : Acme Corp
description : " We build AI-powered widgets "
env :
SHARED_API_URL : https://api.acme.com
secrets :
SHARED_TOKEN : ${ACME_TOKEN}
cron :
standup :
schedule : " 0 9 * * 1-5 "
report_channel : general
tasks :
- title : " Daily standup "
assignee : pm
agents :
designer :
model : anthropic:claude-sonnet-4-20250514
priority : normal # idle | low | normal | high | critical (or 0-4)
thinking : low # off | minimal | low | medium | high | xhigh
description : " Frontend designer — builds HTML/CSS "
prompt_inline : |
You are a frontend designer specializing in responsive layouts.
Focus on clean, semantic HTML and modern CSS.
skills :
- nichochar/web-skills
auth : " oauth:anthropic " # optional — use OAuth instead of API key
api_key_ref : MY_CUSTOM_KEY # optional — host env var name for model key override
env : # non-sensitive, passed as Docker --env (agent overrides office)
LOG_LEVEL : debug
WORKSPACE_NAME : designer
secrets : # sensitive, ${VAR} refs only — delivered via authenticated_fetch
GITHUB_TOKEN : ${MY_GH_TOKEN}
disclose_secrets : true # show secret names in system prompt (default: false)
permissions :
office_cron : true # allow managing office-level cron jobs
reviewer :
model : openai:gpt-4.1
priority : high
thinking : medium
description : " Code reviewer "
Office-level env and secrets are inherited by all agents. Agent-level values override office-level.
All agent fields are optional. Agents are spawned sequentially in declaration order; if one fails, the rest still start. Model availability depends on your provider account — replace the model value with your preferred provider:model-id if the default is unavailable.
Task tools ( task_create , task_update , task_list , task_get , task_delete ) are available to all in-process agents by default. Restrict access via permissions.tools.deny . See Task Management .
Agents can run proactively via heartbeats — periodic messages that prompt agents to check for work or run maintenance without external triggers.
agents :
monitor :
heartbeat :
interval_ms : 60000
prompt : " Check for pending work and report status "
active_hours :
start : " 09:00 "
end : " 17:00 "
Field
Required
Default
Description
inter

[truncated]

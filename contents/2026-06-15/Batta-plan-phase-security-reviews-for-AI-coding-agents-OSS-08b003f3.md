---
source: "https://github.com/Dolevco/batta-ai"
hn_url: "https://news.ycombinator.com/item?id=48539249"
title: "Batta – plan-phase security reviews for AI coding agents (OSS)"
article_title: "GitHub - Dolevco/batta-ai: Security by design for AI agents · GitHub"
author: "dolevco1"
captured_at: "2026-06-15T11:10:50Z"
capture_tool: "hn-digest"
hn_id: 48539249
score: 1
comments: 0
posted_at: "2026-06-15T10:29:24Z"
tags:
  - hacker-news
  - translated
---

# Batta – plan-phase security reviews for AI coding agents (OSS)

- HN: [48539249](https://news.ycombinator.com/item?id=48539249)
- Source: [github.com](https://github.com/Dolevco/batta-ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T10:29:24Z

## Translation

タイトル: Batta – AI コーディング エージェント (OSS) の計画段階のセキュリティ レビュー
記事タイトル: GitHub - Dolevco/batta-ai: AI エージェント向けのセキュリティ バイ デザイン · GitHub
説明: AI エージェント向けの設計によるセキュリティ。 GitHub でアカウントを作成して、Dolevco/batta-ai の開発に貢献してください。

記事本文:
GitHub - Dolevco/batta-ai: AI エージェント向けの設計によるセキュリティ · GitHub
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
ドレフコ
/
ばったあい
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 開く mo

[アクション] メニュー フォルダーとファイル
2 コミット 2 コミット .github .github .vscode .vscode デプロイ デプロイ docs docs パッケージ スクリプト スクリプト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス README.md README.md docker-compose-https.yml docker-compose-https.yml docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Batta はエージェントに適切なセキュリティとコンプライアンスのコンテキストを提供するため、あらゆる意思決定が重要になります。
コードを記述する前に通知されます。これは、マシン速度で実行されるセキュリティ アーキテクトです。
一般的なチェックリストではなく、実際のコード、クラウド構成、組織のポリシーに基づいています。
核となる価値は別のチャット プロンプトではありません。これは、システムのインデックス付きセキュリティ モデルです。
サービス、エントリ ポイント、アイデンティティ、クラウド リソース、データ フロー、信頼境界、データ
分類、脅威、軽減策、既知のギャップ。エージェントが新しい仕事を始めるとき、Batta
提案された変更をそのモデルと比較し、具体的な質問、リスク、
必要なタスク、および人間によるレビューのための証拠に裏付けられた証明書。
計画フェーズのレビューでは、コードを作成する前にセキュリティ レビューを実行し、設計上の欠陥を発見します。
PR中やデプロイ後ではなく、修正するのが最も安価なときに。
完全なコンテキスト。常に実際のコード、クラウド構成、組織内のすべてのレビューを根拠とします。
ポリシー — 一般的なチェックリストではありません。すべてのレビューはシステムの実際の動作を反映しています。
人間の記録システムは、あらゆる決定、発見、

そして証明書。人間
必要なときに完全な監査証跡を残して、重要なことを常に管理します。
エージェントネイティブのワークフローは、MCP を介してインデックス作成とレビューを公開するため、クロード コード、カーソル、
Codex、Copilot Agent、およびその他のコーディング エージェントは、リポジトリ内から Batta を使用できます。
ローカルファースト OSS セットアップは、MCP インデックス作成およびレビュー ループ用の LLM キーなしで機能します。
cp パッケージ/api/.env.example パッケージ/api/.env
ドッカーの構成
http://localhost:3100/onboarding を開き、次のような安定したリポジトリ キーを選択します。
Payments-service を選択し、このプロンプトをコーディング エージェントに貼り付けます。
リポジトリが開いています:
Batta オンボーディング手順を次から取得します。
http://localhost:3101/api/onboarding/agent-led?repo=<リポジトリ名>
次に、このリポジトリの指示に従います。オンボーディングが完了したとみなす前に、MCP を構成し、接続を確認し、このリポジトリにインデックスを付けます。これにより、将来のレビューでアーキテクチャのコンテキストが得られます。
それが推奨されるオンボーディング パスです。エージェントは現在のセットアップ命令を次の場所から取得します。
ローカルの Batta サーバー、リポジトリの MCP の構成、接続の検証、インデックスの作成
リポジトリに追加し、将来の機能作業の前に Batta レビューを実行するための常設の指示を追加します。
インデックス作成は、レビューを汎用的なものではなくアーキテクチャを意識したものにするステップです。
手動セットアップと実稼働 OAuth の詳細が存在します
docs/agent-integration 。
コーディングエージェント
|
| MCP
v
Batta API ----> Postgres + pgvector ----> インデックス付きアーキテクチャ コンテキスト
|
v
セキュリティレビューループ
コーディング エージェントは、Batta MCP を通じてリポジトリにインデックスを付けます。
Batta は、サービス、機能、DFD、脅威モデルなどの構造化されたアーキテクチャ コンテキストを保存します。
人間関係、レビューのギャップなど。
機能や意味のあるコードが変更される前に、エージェントはセキュリティ レビューを開始します。
Batta は変更をインデックス付きアーキテクチャと比較し、欠落しているコンテキストを返します。
リスクと要件

必要なセキュリティタスク。
エージェントは変更を実装し、証拠に裏付けられた証明書をレビューのために提出します。
pnpmインストール
cp パッケージ/api/.env.example パッケージ/api/.env
docker compose up -d postgres redis
pnpm --filter @batta/api dev
pnpm --filter @batta/ui dev
API は http://localhost:3101 で実行され、UI は http://localhost:3101 で実行されます。
http://localhost:3100 。次のようにしてローカルの準備状況を確認します。
pnpmドクター
デフォルトのローカル .env では認証と埋め込みが無効になっているため、最初の実行では認証が必要ありません。
OAuth、証明書、またはモデル キー。
Batta は、ローカル チャット、インデックス作成エージェント、作業項目レビュー エージェントなどに Ollama を使用できます。
セマンティック埋め込み。セットアップ例:
オラマ プル qwen2.5-coder:14b
オラマ プル qwen2.5-coder:7b
オラマプル nomic-embed-text
LLM_PROVIDER = オラマ
OLLAMA_BASE_URL = http://localhost:11434
OLLAMA_CHAT_MODEL = qwen2.5-coder:14b
OLLAMA_SMALL_CHAT_MODEL = qwen2.5-coder:7b
EMBEDINGS_ENABLED = true
EMBEDINGS_PROVIDER = オラマ
OLLAMA_EMBEDDING_MODEL = nomic-embed-text
OLLAMA_EMBEDDING_DIMENSION = 768
ローカルモデルの品質は異なります。コーダー モデルが大きいほど、信頼性が高くなる傾向があります。
テキスト形式のツールは、Batta エージェントが使用する呼び出しです。さまざまな埋め込み
プロバイダーまたはモデルは、次の場合を除き、同じ永続化ベクター データ内に混在させないでください。
インデックスが再構築されます。
┌─────────┐
ブラウザ ────▶│ UI │ (React + Vite)
━━━━┬──────┘
│ 休憩 + SSE
┌──────▼──────┐
コーディングエージェント ─ │ API │ (Express + MCP)
(MCP/OAuth) └──┬───────┬──┘
│ │
▼ ▼
Postgres Redis
+ pgvector (キャッシュ / pubsub)
パッケージ
目的
@batta/ui
オンボーディング、レビュー、ナレッジベース、チャット、統合のための React フロントエンド。
@batta/API
Express REST API と MCP エンドポイント。
@batta/コア
LLM タスク ランタイム、ツール、

nd メモリ プリミティブ。
@batta/共有
永続性、サービス、統合、および共有タイプ。
@batta/データインデクサー
コードとクラウド インデックス作成のためのバックグラウンド スキャナー。
ドキュメント
AI エージェント向けの設計によるセキュリティ
Apache-2.0ライセンス
行動規範
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

Security by design for AI agents. Contribute to Dolevco/batta-ai development by creating an account on GitHub.

GitHub - Dolevco/batta-ai: Security by design for AI agents · GitHub
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
Dolevco
/
batta-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github .github .vscode .vscode deploy deploy docs docs packages packages scripts scripts .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE README.md README.md docker-compose-https.yml docker-compose-https.yml docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json View all files Repository files navigation
Batta gives agents the right security and compliance context — so every decision is
informed before code is written. It is the security architect running at machine speed:
grounded in your actual code, cloud config, and org policies, not generic checklists.
The core value is not another chat prompt. It is the indexed security model of your system:
services, entry points, identities, cloud resources, data flows, trust boundaries, data
classifications, threats, mitigations, and known gaps. When an agent starts new work, Batta
compares the proposed change against that model and returns concrete questions, risks,
required tasks, and evidence-backed attestations for human review.
Plan-phase reviews run security review before code is written — catching design flaws
when they are cheapest to fix, not during PR or after deploy.
Full context, always grounds every review in your actual code, cloud config, and org
policies — not generic checklists. Every review reflects what your system really does.
System of record for humans logs every decision, finding, and attestation. Humans
stay in control of what matters — with a complete audit trail when it counts.
Agent-native workflow exposes indexing and reviews over MCP so Claude Code, Cursor,
Codex, Copilot Agent, and other coding agents can use Batta from inside the repo.
Local-first OSS setup works without an LLM key for MCP indexing and review loops.
cp packages/api/.env.example packages/api/.env
docker compose up
Open http://localhost:3100/onboarding , choose a stable repo key such as
payments-service , then paste this prompt into your coding agent while the target
repository is open:
Fetch Batta onboarding instructions from:
http://localhost:3101/api/onboarding/agent-led?repo=<repo-name>
Then follow those instructions in this repository. Configure MCP, verify the connection, and index this repository before considering onboarding complete so future reviews have architecture context.
That is the recommended onboarding path. The agent fetches current setup instructions from
your local Batta server, configures MCP for the repository, verifies the connection, indexes
the repo, and adds standing instructions to run Batta reviews before future feature work.
Indexing is the step that makes reviews architecture-aware instead of generic.
Manual setup and production OAuth details live in
docs/agent-integration .
coding agent
|
| MCP
v
batta API ----> Postgres + pgvector ----> indexed architecture context
|
v
security review loop
The coding agent indexes the repository through Batta MCP.
Batta stores structured architecture context: services, features, DFDs, threat models,
relationships, and review gaps.
Before a feature or meaningful code change, the agent starts a security review.
Batta compares the change to the indexed architecture and returns missing context,
risks, and required security tasks.
The agent implements the change and submits evidence-backed attestations for review.
pnpm install
cp packages/api/.env.example packages/api/.env
docker compose up -d postgres redis
pnpm --filter @batta/api dev
pnpm --filter @batta/ui dev
The API runs on http://localhost:3101 and the UI runs on
http://localhost:3100 . Check local readiness with:
pnpm doctor
The default local .env disables auth and embeddings so the first run does not require
OAuth, certificates, or model keys.
Batta can use Ollama for local chat, indexing agents, work-item review agents, and
semantic embeddings. Example setup:
ollama pull qwen2.5-coder:14b
ollama pull qwen2.5-coder:7b
ollama pull nomic-embed-text
LLM_PROVIDER = ollama
OLLAMA_BASE_URL = http://localhost:11434
OLLAMA_CHAT_MODEL = qwen2.5-coder:14b
OLLAMA_SMALL_CHAT_MODEL = qwen2.5-coder:7b
EMBEDDINGS_ENABLED = true
EMBEDDINGS_PROVIDER = ollama
OLLAMA_EMBEDDING_MODEL = nomic-embed-text
OLLAMA_EMBEDDING_DIMENSION = 768
Local model quality varies; larger coder models tend to be more reliable for
the text-formatted tool calls Batta agents use. Embeddings from different
providers or models should not be mixed in the same persisted vector data unless
the indexes are rebuilt.
┌─────────────┐
Browser ─────▶│ UI │ (React + Vite)
└──────┬──────┘
│ REST + SSE
┌──────▼──────┐
Coding agent ─▶│ API │ (Express + MCP)
(MCP/OAuth) └──┬───────┬──┘
│ │
▼ ▼
Postgres Redis
+ pgvector (cache / pubsub)
Package
Purpose
@batta/ui
React frontend for onboarding, reviews, knowledge base, chat, and integrations.
@batta/api
Express REST API and MCP endpoint.
@batta/core
LLM task runtime, tools, and memory primitives.
@batta/shared
Persistence, services, integrations, and shared types.
@batta/data-indexer
Background scanner for code and cloud indexing.
Documentation
Security by design for AI agents
Apache-2.0 license
Code of conduct
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

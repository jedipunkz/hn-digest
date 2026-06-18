---
source: "https://github.com/lobu-ai/lobu"
hn_url: "https://news.ycombinator.com/item?id=48590973"
title: "Show HN: Open-source back end for multi-user AI agents with shared memory"
article_title: "GitHub - lobu-ai/lobu: Build AI teammates programmatically · GitHub"
author: "buremba"
captured_at: "2026-06-18T21:45:50Z"
capture_tool: "hn-digest"
hn_id: 48590973
score: 2
comments: 0
posted_at: "2026-06-18T20:21:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source back end for multi-user AI agents with shared memory

- HN: [48590973](https://news.ycombinator.com/item?id=48590973)
- Source: [github.com](https://github.com/lobu-ai/lobu)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T20:21:02Z

## Translation

タイトル: Show HN: 共有メモリを備えたマルチユーザー AI エージェント用のオープンソース バックエンド
記事タイトル: GitHub - lobu-ai/lobu: AI チームメイトをプログラムで構築する · GitHub
説明: AI チームメイトをプログラムで構築します。 GitHub でアカウントを作成して、lobu-ai/lobu の開発に貢献してください。

記事本文:
GitHub - lobu-ai/lobu: AI チームメイトをプログラムで構築する · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ロブアイ
/
ロブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く ac

オプションメニュー フォルダーとファイル
2,056 コミット 2,056 コミット .claude .claude .github .github .husky .husky charts/ lobu charts/ lobu codex-skills/ lobu-builder codex-skills/ lobu-builder config config db db docker docker docs docs 例 例 パッケージ パッケージ パッチ パッチ プロンプト プロンプト スクリプト スクリプト スキル スキル.dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitmodules .gitmodules .node-version .node-version .npmrc .npmrc .nvmrc .nvmrc .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml codecov.yml codecov.yml context7.json context7.json docker-compose.example.yml docker-compose.example.yml package.json package.json release-please-config.json release-please-config.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Lobu — 組織向けマルチテナント OpenClaw
Lobu は、OpenClaw 用のオープンソース マルチテナント ゲートウェイです。ユーザー/チャネルごとに 1 つのサンドボックスとファイルシステム。コンテキスト全体でメモリを共有します。エージェントが秘密を知ることはありません。
OpenClaw は完全なエージェント ランタイム (~800k LOC) ですが、設計上シングル テナントです。すべてのユーザーが同じファイル システムと bash セッションを共有します。 Lobu は、ゲートウェイ層 (約 40k LOC) のみをマルチテナントに書き換え、OpenClaw の Pi ハーネスを各ワーカー内にそのまま保持します。
埋め込みモードでは、再現可能なパッケージに just-bash + Nix を使用します。各ユーザーは、分離された仮想ファイル システムと、インスタンスあたり最大 50 MB の bash セッションを取得します。単一マシン上の 300 の同時インスタンスでテストされており、Docker は必要ありません。
OpenClaw を利用したエージェントを製品に組み込むことも、個人ごとに個別のインスタンスを管理せずにチーム エージェントに提供することもできます。
デモ-readme.mp4

チャネルとAPI
REST API — プログラムによるエージェントの作成、制御、および状態。
Slack — 豊富な対話性を備えたマルチチャネル/DM エージェント。
Telegram — インタラクティブなワークフローを備えた Webhook またはポーリング ボット。
WhatsApp — WhatsApp ビジネス クラウド API。
Discord — チャンネル + DM ボットのサポート。
Google Chat — カード v2、ワークスペース スペース。
CLI 経由でスキャフォールディングして実行します。 Lobu は、デフォルトでゼロ構成の Postgres が組み込まれた単一のノード プロセスとして起動します (または、 DATABASE_URL 経由で独自の pgvector を持ち込むこともできます)。
npx @lobu/cli@latest init my-bot
cd マイボット
npx @lobu/cli@latest run # スタックを起動し、エージェントを適用します
npx @lobu/cli@latest chat -c local " hello " # 話しかけてください
lobu run (embedded) は lobu.config.ts を自動適用するため、スキャフォールドされたエージェントをすぐに使用できるようになります。外部 Postgres を使用するには、 .env で DATABASE_URL を設定します。後の構成変更をプッシュするには、 lobu apply を実行します。
ランタイム設定は、Web アプリまたは CLI で使用される同じ組織スコープの REST API を通じて管理されます。
npx @lobu/cli@最新ログイン
npx @lobu/cli@latest 組織セット my-org
npx @lobu/cli@最新エージェントリスト
ローカル lobu.config.ts プロジェクトは、lobu validate および lobu apply ワークフローに引き続き役立ちます。
単一プロセス ノードは依然として最も単純なデプロイメントです。これを、node 、 pm2 、 systemd 、または別のプロセス スーパーバイザで実行します。アプリには、その環境からアクセス可能な DATABASE_URL (Postgres + pgvector) が必要です。
ローカル開発 (Lobu 自体に貢献): clone、make setup 、make dev (組み込みゲートウェイ + ワーカー + Vite HMR を :8787 でブート)。
実稼働 (VM/ベアメタル) : --cwd package/server build:server を実行し、選択したプロセス スーパーバイザの下で package/server/dist/server.bundle.mjs をノードに実行します。
実稼働 (Kubernetes) : charts/lobu でパブリック Helm チャートを使用します。
helm install lobu oci://ghcr.io/lobu-ai/charts/lobu \
--namespace lobu --cre

食べた名前空間 \
-f あなたの値.yaml
調整パラメータの完全なセットについては、charts/lobu/values.yaml を参照してください。少なくとも、
入力ホスト、DATABASE_URL + ENCRYPTION_KEY + を含む SecretName シークレット
BETTER_AUTH_SECRET + プロバイダー API キー、および database.existingSecret 。
フローチャート LR
Slack[スラック] <--> GW[ゲートウェイ]
電報[電報] <--> GW
WhatsApp[WhatsApp] <--> GW
Discord[ディスコード] <--> GW
API[REST API] <--> GW
GW <--> PG[(Postgres)]
GW -->|スポーン| W[労働者]
サブグラフサンドボックス
W
終わり
W -.->|HTTP プロキシ| GW
W -.->|MCP プロキシ| GW
GW -->|ドメインフィルター|インターネット((インターネット))
GW -->|スコープ指定されたトークン| MCP[MCPサーバー]
読み込み中
能力
すべての Lobu エージェントには、自律的な実行と永続性のためのツールが付属しています。
生産性: Google カレンダー、Slack、Jira、Notion
開発: GitHub、GitLab、Postgres、Docker
知識: Wikipedia、Brave Search、YouTube、PDF 検索
単一出力としてのゲートウェイ。すべてのワーカー トラフィック (インターネットと MCP) はゲートウェイを経由してルーティングされます。従業員はネットワークに直接アクセスできません。ドメイン フィルタリングは、どのサービスに到達するかを制御します。
MCP プロキシ。ゲートウェイは ${env:VAR} シークレットを解決し、上流の MCP サーバーにルーティングします。サードパーティ API の OAuth は Lobu に残り、ワーカーがトークンを参照することはありません。
マルチプラットフォーム、マルチテナント。 1 つのインスタンスは、Slack、Telegram、WhatsApp、Discord、Teams、および REST API を提供します。各チャネル/DM は、独自のランタイム、モデル、ツール、資格情報、および Nix パッケージを取得します。
OpenClaw ランタイム。ワーカーは、エージェントごとのモデルを選択して OpenClaw Pi エージェントを実行します。 OpenClaw スキルと IDENTITY.md / SOUL.md / USER.md ワークスペース ファイルをサポートします。
マルチプロバイダー認証。構成主導のレジストリを介した 16 の LLM プロバイダー (OpenAI、Gemini、Groq、DeepSeek、Mistral など)。 API キーはゲートウェイ上に残ります。
Lobu は、自律エージェントのインフラストラクチャ層です。 LangChain や CrewAI などのフレームワークは作成に役立ちます

エージェントロジック。 Lobu は、サンドボックス化、永続化、メッセージング接続など、これらのエージェントを大規模に実行する配信レイヤーです。
ワーカーはゲートウェイ プロキシ経由で出力します — HTTP_PROXY=http://localhost:8118 (allowlist/blocklist + LLM 出力判定を使用)。 Linux 運用ホストでは、ワーカー スポーンは IPAddressDeny=any を指定した systemd-run --user --scope を使用して、カーネル レベルで出力を強制します。開発環境 (macOS) では、プロキシはベストエフォート型です。
シークレットはゲートウェイに残ります - プロバイダーの資格情報と ${env:} 置換。 OAuth はロブにあります。労働者は本物の鍵を見ることはありません。
脅威モデル: シングルテナントのローカル分離 - just-bash と孤立 vm はポリシー + ベストエフォート型のサンドボックスであり、敵対的なコードのセキュリティ境界ではありません。 Lobu を信頼できないユーザーに公開する前に、docs/SECURITY.md を参照してください。
Nix システム パッケージ — エージェントごとに再現可能なツールとスキル ポリシー。
Lobu はオープンソースですが、実稼働グレードのエージェントを展開するには、通常、魂、アイデンティティ、統合を調整する必要があります。私は次のような実践的な実装を提供します。
従業員 AI アシスタント - 社内ツールやドキュメントに接続された Slack 上の永続的なサンドボックス エージェント。
自動化されたカスタマー サポート - 人間が関与する複数ステップのチケット処理。
自律型ワークフロー - 永続的な状態で長時間実行され、スケジュールされたバックグラウンド ジョブ。
マネージド インフラストラクチャ — アップデートとスケーリングを備えたプライベート Lobu デプロイメント。
カスタム ツールとスキル — 特注の MCP サーバー、Nix ランタイム、OpenClaw スキル。
私は 2 回目の技術創業者です。以前にrakam.io (エンタープライズ分析PaaS) を設立し、LiveRamp (NYSE: RAMP) に買収されました。
チームや顧客のために常駐エージェントが必要ですか? Founder に相談するか、X/Twitter に連絡してください。
AI チームメイトをプログラムで構築する
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロップ

アーティーズ
星
22
フォーク
レポートリポジトリ
リリース
53
ロブ: v12.0.0
最新の
2026 年 6 月 18 日
+ 52 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Build AI teammates programmatically. Contribute to lobu-ai/lobu development by creating an account on GitHub.

GitHub - lobu-ai/lobu: Build AI teammates programmatically · GitHub
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
lobu-ai
/
lobu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,056 Commits 2,056 Commits .claude .claude .github .github .husky .husky charts/ lobu charts/ lobu codex-skills/ lobu-builder codex-skills/ lobu-builder config config db db docker docker docs docs examples examples packages packages patches patches prompts prompts scripts scripts skills skills .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .gitmodules .gitmodules .node-version .node-version .npmrc .npmrc .nvmrc .nvmrc .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md bun.lock bun.lock bunfig.toml bunfig.toml codecov.yml codecov.yml context7.json context7.json docker-compose.example.yml docker-compose.example.yml package.json package.json release-please-config.json release-please-config.json tsconfig.json tsconfig.json View all files Repository files navigation
Lobu — Multi-tenant OpenClaw for Organizations
Lobu is an open-source multi-tenant gateway for OpenClaw . One sandbox and filesystem per user/channel. Shared memory across contexts. Agents never see secrets.
OpenClaw is a full agent runtime (~800k LOC) but it's single-tenant by design — every user shares the same filesystem and bash session. Lobu rewrites only the gateway layer (~40k LOC) to be multi-tenant and keeps OpenClaw's Pi harness untouched inside each worker.
Embedded mode uses just-bash + Nix for reproducible packages. Each user gets an isolated virtual filesystem and bash session at ~50MB per instance — tested at 300 concurrent instances on a single machine, no Docker needed.
Embed OpenClaw-powered agents into your product, or give your team agents without managing a separate instance per person.
demo-readme.mp4
Channels & API
REST API — programmatic agent creation, control, and state.
Slack — multi-channel/DM agents with rich interactivity.
Telegram — webhook or polling bot with interactive workflows.
WhatsApp — WhatsApp Business Cloud API.
Discord — channel + DM bot support.
Google Chat — Cards v2, Workspace spaces.
Scaffold and run via the CLI. Lobu boots as a single Node process with a zero-config embedded Postgres by default (or bring your own — pgvector required — via DATABASE_URL ).
npx @lobu/cli@latest init my-bot
cd my-bot
npx @lobu/cli@latest run # boots the stack and applies your agent
npx @lobu/cli@latest chat -c local " hello " # talk to it
lobu run (embedded) auto-applies your lobu.config.ts , so the scaffolded agent is usable immediately. To use an external Postgres, set DATABASE_URL in .env ; to push later config changes, run lobu apply .
Runtime configuration is managed through the web app or the same org-scoped REST API used by the CLI:
npx @lobu/cli@latest login
npx @lobu/cli@latest org set my-org
npx @lobu/cli@latest agent list
Local lobu.config.ts projects are still useful for lobu validate and lobu apply workflows.
Single-process Node remains the simplest deployment: run it with node , pm2 , systemd , or another process supervisor. The app needs DATABASE_URL (Postgres + pgvector) reachable from its environment.
Local dev (contributing to Lobu itself): clone, make setup , make dev (boots embedded gateway + workers + Vite HMR on :8787 ).
Production (VM/bare metal) : bun run --cwd packages/server build:server , then node packages/server/dist/server.bundle.mjs under your process supervisor of choice.
Production (Kubernetes) : use the public Helm chart in charts/lobu :
helm install lobu oci://ghcr.io/lobu-ai/charts/lobu \
--namespace lobu --create-namespace \
-f your-values.yaml
See charts/lobu/values.yaml for the full set of tunables. At minimum supply an
ingress host, a secretName Secret containing DATABASE_URL + ENCRYPTION_KEY +
BETTER_AUTH_SECRET + provider API keys, and a database.existingSecret .
flowchart LR
Slack[Slack] <--> GW[Gateway]
Telegram[Telegram] <--> GW
WhatsApp[WhatsApp] <--> GW
Discord[Discord] <--> GW
API[REST API] <--> GW
GW <--> PG[(Postgres)]
GW -->|spawn| W[Worker]
subgraph Sandbox
W
end
W -.->|HTTP proxy| GW
W -.->|MCP proxy| GW
GW -->|domain filter| Internet((Internet))
GW -->|scoped tokens| MCP[MCP Servers]
Loading
Capabilities
Every Lobu agent ships with tools for autonomous execution and persistence:
Productivity: Google Calendar, Slack, Jira, Notion
Development: GitHub, GitLab, Postgres, Docker
Knowledge: Wikipedia, Brave Search, YouTube, PDF Search
Gateway as single egress. All worker traffic — internet and MCP — routes through the gateway. Workers have no direct network access; domain filtering controls which services they reach.
MCP proxy. Gateway resolves ${env:VAR} secrets and routes to upstream MCP servers. OAuth for third-party APIs stays in Lobu — workers never see tokens.
Multi-platform, multi-tenant. One instance serves Slack, Telegram, WhatsApp, Discord, Teams, and the REST API. Each channel/DM gets its own runtime, model, tools, credentials, and Nix packages.
OpenClaw runtime. Workers run OpenClaw Pi Agent with per-agent model selection. Supports OpenClaw skills and IDENTITY.md / SOUL.md / USER.md workspace files.
Multi-provider auth. 16 LLM providers (OpenAI, Gemini, Groq, DeepSeek, Mistral, …) via a config-driven registry. API keys stay on the gateway.
Lobu is the infrastructure layer for autonomous agents. Frameworks like LangChain or CrewAI help you write agent logic; Lobu is the delivery layer that runs those agents at scale — sandboxing, persistence, and messaging connectivity.
Worker egress through the gateway proxy — HTTP_PROXY=http://localhost:8118 with allowlist/blocklist + LLM egress judge. On Linux production hosts the worker spawn uses systemd-run --user --scope with IPAddressDeny=any to enforce egress at the kernel level; in dev (macOS) the proxy is best-effort.
Secrets stay in gateway — provider credentials and ${env:} substitution; OAuth lives in Lobu. Workers never see real keys.
Threat model: single-tenant local isolation — just-bash and isolated-vm are policy + best-effort sandboxes, not security boundaries for hostile code. See docs/SECURITY.md before exposing Lobu to untrusted users.
Nix system packages — per-agent reproducible tooling and skill policy.
Lobu is open source, but deploying production-grade agents usually means tuning soul, identity, and integrations. I offer hands-on implementation for:
Employee AI assistants — persistent sandboxed agents on Slack wired into internal tools and docs.
Automated customer support — multi-step ticket handling with human-in-the-loop.
Autonomous workflows — long-running, scheduled background jobs with persistent state.
Managed infrastructure — private Lobu deployments with updates and scaling.
Custom tooling & skills — bespoke MCP servers, Nix runtimes, and OpenClaw skills.
I'm a second-time technical founder. Previously founded rakam.io (enterprise analytics PaaS), acquired by LiveRamp (NYSE: RAMP).
Want persistent agents for your team or customers? Talk to Founder or reach out on X/Twitter .
Build AI teammates programmatically
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
22
forks
Report repository
Releases
53
lobu: v12.0.0
Latest
Jun 18, 2026
+ 52 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

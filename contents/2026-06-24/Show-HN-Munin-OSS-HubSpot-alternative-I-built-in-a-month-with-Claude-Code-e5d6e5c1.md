---
source: "https://github.com/getmunin/munin"
hn_url: "https://news.ycombinator.com/item?id=48659577"
title: "Show HN: Munin – OSS HubSpot alternative I built in a month with Claude Code"
article_title: "GitHub - getmunin/munin: The customer platform for the agentic era. MCP-first, open source, self-hostable. · GitHub"
author: "kman_85"
captured_at: "2026-06-24T13:42:38Z"
capture_tool: "hn-digest"
hn_id: 48659577
score: 1
comments: 0
posted_at: "2026-06-24T13:35:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Munin – OSS HubSpot alternative I built in a month with Claude Code

- HN: [48659577](https://news.ycombinator.com/item?id=48659577)
- Source: [github.com](https://github.com/getmunin/munin)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T13:35:04Z

## Translation

タイトル: Show HN: Munin – Claude Code を使って 1 か月で構築した OSS HubSpot の代替案
記事のタイトル: GitHub - getmunin/munin: エージェント時代の顧客プラットフォーム。 MCP ファースト、オープンソース、自己ホスト可能。 · GitHub
説明: エージェント時代の顧客プラットフォーム。 MCP ファースト、オープンソース、自己ホスト可能。 - ゲットムニン/ムニン

記事本文:
GitHub - getmunin/munin: エージェント時代の顧客プラットフォーム。 MCP ファースト、オープンソース、自己ホスト可能。 · GitHub
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
ゲムニン
/
ムニン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
619 コミット 619 コミット .changeset .changeset .github .github .husky .husky アプリ アプリ パッケージ パッケージ スクリプト スクリプト .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス LICENSE_POLICY.json LICENSE_POLICY.json README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md compose.yml compose.yml lint-staged.config.mjs lint-staged.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml server.json server.json tsconfig.base.json tsconfig.base.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェント時代に向けて作られた MCP ファーストの顧客プラットフォーム。エージェントは UI です。
Munin は、AI エージェントを介した唯一のインターフェイスであるエージェント時代 (ナレッジベース、会話、CRM、CMS、アウトリーチ、分析) 向けに構築されたオープンソースの顧客プラットフォームです。アプリ自体には従来の管理 UI はありません。すべてのアクションは、MCP 互換クライアントから呼び出し可能な MCP ツールを通じて実行されます。
Lovable はフロントエンドを構築します。 Munin があなたの業務をスピンアップします。 1 つのプロンプト、1 つの MCP エンドポイント - 残りはエージェントが行います。
Lovable が 1 つのプロンプトから実際の Web サイトを構築する一方、Munin がブログの読み取り元となる CMS、シードされたナレッジ ベース、分析、ビジネスをすでに認識しているチャット ウィジェットなど、その背後にあるすべてのものを立ち上げている様子をご覧ください。管理 UI や接続するダッシュボードはありません。エージェントは 1 つの MCP エンドポイントを介して作業を実行します。その後、実際の顧客との会話が展開されます。

知識ベースから応答され、重要なときに人間に引き渡され、エージェントによって回収されて終了します。
こんにちは — Kjell です。昨年、私の会社は HubSpot にほとんど触れていない 3 人のユーザーに月額 500 ドルを支払っていました。どのワークフローも、私たちが使用していたものよりも 1 つ上の階層で生きてきた価値があります。
その後、AI エージェントが自らプロスペクティングとアウトリーチを行うようになり、私は衝撃を受けました。これまで HubSpot を使用していた目的が、基本的な CRUD にまで縮小したのです。エージェントが作業を行います。 HubSpot は記録を保存したところです。 CRUD に月額 500 ドルを支払い続けるつもりはありませんでした。
それで実際に欲しかったものを作りました。 Claude Code を主な IDE として使用し、約 1 か月かけて、Munin は何もない状態からこのリポジトリにあるものまでになりました。私は約 20 年間、製品版ソフトウェアを出荷してきました。その経験がアーキテクチャを形成し、問題が発生したときにエージェントを捕まえることができました。クロード・コードは、それを方程式からタイプする手間を取り除きました。 1年前にはこれはソロでは実現不可能だったと思います。
Munin は AI ではありません。AI が使用するものです。あなたはエージェントを連れてきて、プロンプトを書き、ワークフローを組み立てます。 Munin の仕事は、ツールの下の表面を可能な限りきれいにすることです。これはオープンソースであり、あなたが実行できます。
ナレッジベース — マークダウンドキュメント、ハイブリッド検索 (BM25 + pgvector 埋め込み)、ドキュメントごとの対象範囲指定。
会話 — マルチチャネル スレッド (電子メール、Threll.ai または Vapi 経由の音声、Twilio または MessageBird 経由の SMS、チャット ウィジェット) はエンドユーザーにルーティングされ、割り当て可能でエージェントが解決可能で、ハンドオーバー状態と Webhook ファンアウトを備えています。
CRM — 連絡先、企業、取引、アクティビティ、パイプライン、セグメントに加えて、クリーンな連絡先データのキュレーターが実行するマージ提案キュー。
CMS — フィールド スキーマ、ローカライズされたエントリ、スケジュールされた公開、S3 互換のアセット ライブラリを備えたエージェント作成のコンテンツ コレクション

ストレージ、および組み込みのエンゲージメント シグナルのすべてのエントリに _tracking ブロックを送信するパブリック配信 API。
アウトリーチ — 提案のみのアウトバウンド電子メール: キャンペーン + セグメント + 人間の承認のためにキューに入れられたドラフト。決して自動送信しません。
分析 — ポリモーフィックなページビュー + 検索クエリの取り込み。 CMS 配信は無料で行われます。任意のページは <script src="…/tracker.js" data-key="mn_track_…"> タグにドロップされます。読み取りツール (analytics_list_top_subjects 、analytics_get_subject_engagement 、analytics_list_zero_result_searches ) は、「次に何を書こうか」というループを cms/review-stale-entries などのキュレーター スキルにフィードします。
キュレーター — 再試行 + デッドレターを伴うスケジュールでの耐久性のあるバックグラウンド ジョブ キュー実行スキル (KB キュレーション、CRM 衛生管理、連絡先抽出、古いコンテンツのレビュー、アウトリーチの下書き)。
プレイブック + スキル — パッケージ化されたマークダウン プロシージャ ( skill://module/<verb-object> ) エージェントは MCP リソースを介して読み取り、マルチステップ フローに従います。
データの移植性 — モジュールごとに対称的な *_export / *_import MCP ツール (および /v1/<module>/export|import REST) を使用できるため、エージェントは自己ホスト型サーバーとクラウドの間で組織のデータをどちらの方向にも移動できます。 skill://playbooks/data-migration を参照してください。
横断的 — 監査ログ、Webhook、チームへの招待、OAuth 2.1 動的クライアント登録、BetterAuth を利用したサインイン。
セルフホスト (このリポジトリ): シングルテナント、招待のみ。
git clone https://github.com/getmunin/munin.git
CDムニン
cp .env.example .env # MUNIN_AUTH_SECRET + MUNIN_KEY_PEPPER + MUNIN_ENCRYPTION_KEY を編集します
ドッカーの構成
最初にサインアップしたユーザーが組織管理者になります。後続のユーザーには、招待トークンまたはドメインが MUNIN_ALLOWED_EMAIL_DOMAINS にある電子メールが必要です。
ホスト型 ( https://getmunin.com ): マルチテナント、組織ごとに 1 つのサインアップ。
docker compose up 後、バックエンドは :30 でリッスンします。

01 と :3000 のダッシュボード。
http://localhost:3000 を開いて最初のユーザーを登録します。そのユーザーがシングルトン組織管理者になります。
ダッシュボードで、「設定」→「API キー」に移動し、管理キー ( mn_admin_… ) を作成します。 1 回表示されます。パスワードのように扱います。
# REST コントロール プレーン - 直接、OAuth なし
カール -s http://localhost:3001/v1/kb/spaces \
-H " 権限: ベアラー mn_admin_... " | jq
# MCP ツールブラウザ (ツール/スキルをつつく場合に推奨)
npx @modelcontextprotocol/inspector
# UI 内: URL = http://localhost:3001/mcp、認証 = Bearer mn_admin_...
# クロードコード (CLI)
クロード mcp munin-local を追加 http://localhost:3001/mcp
# 次に `claude` → OAuth フローがブラウザで開きます → ツールが利用可能になります
# クロードデスクトップ — claude_desktop_config.json
# {
# "mcpServers": { "munin": { "url": "http://localhost:3001/mcp" } }
# }
# Streamable HTTP 上の生のカール — トランスポートの健全性チェックに役立ちます
curl -N -X POST http://localhost:3001/mcp \
-H " 権限: ベアラー mn_admin_... " \
-H " Content-Type: application/json " \
-H " 受け入れる: application/json、text/event-stream " \
-d ' {"jsonrpc":"2.0","id":1,"method":"tools/list"} '
REST コントロール プレーンの OpenAPI 仕様は、packages/backend-core/openapi.json にあります。
サインアップ (ホスト) するか、docker compose up (セルフホスト) を実行したら、MCP クライアントに URL を指定します。
Claude デスクトップ / コード — MCP 構成に追加します。
{
"mcpサーバー": {
"ムニン" : {
"url" : " http://localhost:3001/mcp "
}
}
}
(ホスト型バージョンの場合は、 https://mcp.getmunin.com に切り替えます。) 最初の呼び出しでブラウザーに OAuth 同意画面がトリガーされ、エージェントはナレッジ ベース、会話、CRM、CMS、アウトリーチ、分析などのツールの全領域を利用できるようになります。
開発者ドキュメントは getmunin.com/docs にあります — ガイド、REST API リファレンス、完全な MCP ツール リスト、スキル ライブラリ

。
2 つの信頼コンテキスト、1 つの MCP エンドポイント
同じ /mcp エンドポイントは、視聴者を意識した 2 つの異なる呼び出し元にサービスを提供します。
管理エージェント (Claude Desktop、Cursor、内部オートメーション) — OAuth 承認済み。ツール全画面、kb:* 、conv:* 、crm:* 、cms:* 、outreach:* 、analytics:* ごとにスコープゲートされます。
エンドユーザー エージェント (音声 AI、Web チャットボット、モバイル アプリ ヘルパー) — バックエンドからサーバー側で生成された、エンドユーザーの 1 人を対象とする、有効期間の短い委任トークン。セルフサービス ツールのみ (自分の連絡先を読み、自分の会話でメッセージを送信)。
token-mint API については、packages/backend-core/src/control/delegated-token.controller.ts を参照してください。 @getmunin/sdk ノード クライアントがそれをラップします。
apps/backend — シングルテナント AuthModule を使用して @getmunin/backend-core モジュールを構成する薄い NestJS エントリ。 MCP サーバー (ストリーミング可能 HTTP)、OAuth 2.1 + OIDC 検出、およびコントロール プレーン REST API をポート 3001 で公開します。
apps/web — Next.js ダッシュボード + ポート 3000 に着陸。
apps/chat-widget — mn_widget_* キーとウィジェット取り込み API を使用する埋め込み可能なブラウザ ウィジェット。
apps/analytics-tracker — /tracker.js で提供される埋め込み可能なブラウザ トラッカー。公開 mn_track_* キーを消費し、ページビュー イベントを書き込みます。
Packages/backend-core — すべての共有 NestJS モジュール (KB、Conversations、CRM、CMS、Outreach、Analytics、Curator、Web、Playbooks、MCP、コントロール プレーン、監査、テナント、RLS、メーラー、ストレージ、Webhook、レート制限、クォータ、OAuth) と createApp および createAuthController ヘルパー。クラウドは、マルチテナント認証を使用して同じモジュールを構成します。
package/dashboard-pages — OSS とクラウド Web 間で共有されるダッシュボード ページ コンポーネント。
package/ui — デザインシステムのプリミティブ (shadcn スタイル)。
package/{core, db, tables, sdk, mcp-toolkit} — 非 Nest ビルディング ブロック: アクター ID + テナンシー GUC、Drizzl

スキーマ + 移行、共有タイプ、ノード クライアント SDK、MCP @McpTool / @SkillRegistry デコレータ。
Packages/{agent-host, Agent-runtime} — キューからキュレーター ジョブを選択し、設定された LLM プロバイダーに対して実行する、耐久性のある組織ごとのエージェント ランナー。
package/widget-voice — 共通の VoiceSession インターフェイスの背後にあるチャット ウィジェット (Threll.ai WebRTC + Vapi) 用のベンダーに依存しないブラウザー音声 SDK。
すべての @getmunin/* パッケージは GitHub パッケージに公開されます。
TypeScript · Node 24 LTS · Turborepo · pnpm · NestJS · Next.js · Drizzle · Postgres + pgvector · MCP Streamable HTTP · BetterAuth + OAuth 2.1 動的クライアント登録
バンドルされたサードパーティの依存関係は、独自のライセンスを保持します。THIRD_PARTY_LICENSES.md (pnpm Licenses:generate によって生成され、CI で検証されています) を参照してください。
エージェント時代の顧客プラットフォーム。 MCP ファースト、オープンソース、自己ホスト可能。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
22
v4.57.1
最新の
2026 年 6 月 23 日
+ 21 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

[切り捨てられた]

## Original Extract

The customer platform for the agentic era. MCP-first, open source, self-hostable. - getmunin/munin

GitHub - getmunin/munin: The customer platform for the agentic era. MCP-first, open source, self-hostable. · GitHub
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
getmunin
/
munin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
619 Commits 619 Commits .changeset .changeset .github .github .husky .husky apps apps packages packages scripts scripts .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE_POLICY.json LICENSE_POLICY.json README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_LICENSES.md THIRD_PARTY_LICENSES.md compose.yml compose.yml lint-staged.config.mjs lint-staged.config.mjs package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml server.json server.json tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
MCP-first customer platform made for the agentic era. The agent is the UI.
Munin is an open-source customer platform built for the agentic era (Knowledge Base, Conversations, CRM, CMS, Outreach, Analytics) where the only interface is via AI agents. There is no traditional admin UI for the apps themselves — every action runs through MCP tools, callable from any MCP-compatible client.
Lovable builds your frontend. Munin spins up your operations. One prompt, one MCP endpoint — and the agents do the rest.
Watch Lovable build a real website from a single prompt while Munin stands up everything behind it — the CMS the blog reads from, a seeded knowledge base, analytics, and a chat widget that already knows the business. No admin UI, no dashboard to wire up; the agent does the work, over one MCP endpoint. Then a real customer conversation plays out: answered from the knowledge base, handed off to a human when it matters, and picked back up by the agent to close.
Hi — Kjell here. Last year my company was paying HubSpot $500/month for three users who barely touched it. Every workflow worth having lived a tier above the one we were on.
Then our AI agents started doing the prospecting and outreach themselves, and it hit me: what we still used HubSpot for had shrunk to basic CRUD. The agent does the work; HubSpot just stored the record. I wasn't going to keep paying $500/month for CRUD.
So I built the thing I actually wanted. Over roughly a month, with Claude Code as my primary IDE, Munin went from nothing to what's in this repo. I've shipped production software for about twenty years — that experience is what shaped the architecture and let me catch the agent when it got things wrong. Claude Code took the labor of typing it out of the equation; a year ago I don't think this would have been feasible solo.
Munin isn't the AI — it's what your AI uses. You bring the agent, write the prompts, wire up the workflows; Munin's job is to be the cleanest possible tool surface underneath. It's open source and yours to run.
Knowledge Base — markdown documents, hybrid search (BM25 + pgvector embeddings), per-document audience scoping.
Conversations — multi-channel threads (email, voice via Threll.ai or Vapi, SMS via Twilio or MessageBird, chat widget) routed to end-users, assignable, agent-resolvable, with handover state and webhook fan-out.
CRM — contacts, companies, deals, activities, pipelines, segments, plus a merge-proposal queue the clean-contact-data curator runs against.
CMS — agent-authored content collections with field schemas, localized entries, scheduled publishing, an asset library backed by S3-compatible storage, and a public delivery API that ships a _tracking block on every entry for built-in engagement signal.
Outreach — propose-only outbound emails: campaigns + segments + drafts queued for human approval; never auto-sends.
Analytics — polymorphic page-view + search-query ingestion. CMS delivery wires in for free; arbitrary pages drop in a <script src="…/tracker.js" data-key="mn_track_…"> tag. Read tools ( analytics_list_top_subjects , analytics_get_subject_engagement , analytics_list_zero_result_searches ) feed the "what should we write next" loop into curator skills like cms/review-stale-entries .
Curator — durable background job queue running skills (KB curation, CRM hygiene, contact extraction, stale-content review, outreach drafts) on a schedule with retry + dead-letter.
Playbooks + skills — packaged markdown procedures ( skill://module/<verb-object> ) the agent reads via MCP resources to follow multi-step flows.
Data portability — symmetric *_export / *_import MCP tools (and /v1/<module>/export|import REST) per module, so an agent can move an org's data between a self-hosted server and the cloud in either direction. See skill://playbooks/data-migration .
Cross-cutting — audit log, webhooks, team invites, OAuth 2.1 dynamic-client registration, BetterAuth-backed sign-in.
Self-host (this repo): single-tenant, invite-only.
git clone https://github.com/getmunin/munin.git
cd munin
cp .env.example .env # edit MUNIN_AUTH_SECRET + MUNIN_KEY_PEPPER + MUNIN_ENCRYPTION_KEY
docker compose up
The first user to sign up becomes the org admin; subsequent users need an invitation token or an email whose domain is in MUNIN_ALLOWED_EMAIL_DOMAINS .
Hosted ( https://getmunin.com ): multi-tenant, one signup per org.
After docker compose up , the backend listens on :3001 and the dashboard on :3000 .
Open http://localhost:3000 and register the first user — they become the singleton org admin.
In the dashboard, go to Settings → API keys and mint an admin key ( mn_admin_… ). Shown once; treat like a password.
# REST control plane — direct, no OAuth
curl -s http://localhost:3001/v1/kb/spaces \
-H " Authorization: Bearer mn_admin_... " | jq
# MCP tool browser (recommended for poking at tools/skills)
npx @modelcontextprotocol/inspector
# In its UI: URL = http://localhost:3001/mcp, Auth = Bearer mn_admin_...
# Claude Code (CLI)
claude mcp add munin-local http://localhost:3001/mcp
# Then `claude` → OAuth flow opens in browser → tools available
# Claude Desktop — claude_desktop_config.json
# {
# "mcpServers": { "munin": { "url": "http://localhost:3001/mcp" } }
# }
# Raw curl over Streamable HTTP — useful for sanity-checking the transport
curl -N -X POST http://localhost:3001/mcp \
-H " Authorization: Bearer mn_admin_... " \
-H " Content-Type: application/json " \
-H " Accept: application/json, text/event-stream " \
-d ' {"jsonrpc":"2.0","id":1,"method":"tools/list"} '
The OpenAPI spec for the REST control plane is at packages/backend-core/openapi.json .
Once you've signed up (hosted) or run docker compose up (self-host), point your MCP client at the URL.
Claude Desktop / Code — add to your MCP config:
{
"mcpServers" : {
"munin" : {
"url" : " http://localhost:3001/mcp "
}
}
}
(For the hosted version, swap in https://mcp.getmunin.com .) The first call triggers an OAuth consent screen in your browser, then your agent has the full tool surface — Knowledge Base, Conversations, CRM, CMS, Outreach, Analytics.
Developer docs live at getmunin.com/docs — guides, the REST API reference, the full MCP tool list, and the skill library.
Two trust contexts, one MCP endpoint
The same /mcp endpoint serves two distinct callers, audience-aware:
Admin agents (Claude Desktop, Cursor, internal automation) — OAuth-authorized by you. Full tool surface, scope-gated per kb:* , conv:* , crm:* , cms:* , outreach:* , analytics:* .
End-user agents (your voice AI, web chatbot, mobile app helper) — short-lived delegated tokens minted server-side from your backend, scoped to one of your end-users. Only self-service tools (read your own contact, send a message in your own conversation).
See packages/backend-core/src/control/delegated-token.controller.ts for the token-mint API. The @getmunin/sdk Node client wraps it.
apps/backend — thin NestJS entry composing @getmunin/backend-core modules with single-tenant AuthModule . Exposes the MCP server (Streamable HTTP), OAuth 2.1 + OIDC discovery, and the control-plane REST API on port 3001.
apps/web — Next.js dashboard + landing on port 3000.
apps/chat-widget — embeddable browser widget that consumes a mn_widget_* key and the widget ingest API.
apps/analytics-tracker — embeddable browser tracker, served at /tracker.js , that consumes a public mn_track_* key and writes page-view events.
packages/backend-core — every shared NestJS module (KB, Conversations, CRM, CMS, Outreach, Analytics, Curator, Web, Playbooks, MCP, control plane, audit, tenancy, RLS, mailer, storage, webhooks, rate limit, quotas, OAuth) plus createApp and the createAuthController helper. Cloud composes the same modules with multi-tenant auth.
packages/dashboard-pages — dashboard page components shared between OSS and cloud webs.
packages/ui — design-system primitives (shadcn-style).
packages/{core, db, types, sdk, mcp-toolkit} — non-Nest building blocks: actor identity + tenancy GUCs, Drizzle schema + migrations, shared types, Node client SDK, MCP @McpTool / @SkillRegistry decorators.
packages/{agent-host, agent-runtime} — durable per-org agent runner that picks curator jobs off the queue and executes them against the configured LLM provider.
packages/widget-voice — vendor-agnostic browser voice SDK for the chat widget (Threll.ai WebRTC + Vapi) behind a common VoiceSession interface.
All @getmunin/* packages are published to GitHub Packages.
TypeScript · Node 24 LTS · Turborepo · pnpm · NestJS · Next.js · Drizzle · Postgres + pgvector · MCP Streamable HTTP · BetterAuth + OAuth 2.1 dynamic client registration
Bundled third-party dependencies retain their own licenses — see THIRD_PARTY_LICENSES.md (generated by pnpm licenses:generate , verified in CI).
The customer platform for the agentic era. MCP-first, open source, self-hostable.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
22
v4.57.1
Latest
Jun 23, 2026
+ 21 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal informati

[truncated]

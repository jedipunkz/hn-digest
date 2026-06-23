---
source: "https://github.com/latitude-dev/latitude-llm"
hn_url: "https://news.ycombinator.com/item?id=48645555"
title: "Show HN: Find and fix issues in your AI Agents (MIT, open source)"
article_title: "GitHub - latitude-dev/latitude-llm: Latitude is the open-source ai monitoring platform. · GitHub"
author: "paulaq"
captured_at: "2026-06-23T15:01:22Z"
capture_tool: "hn-digest"
hn_id: 48645555
score: 2
comments: 0
posted_at: "2026-06-23T14:30:54Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Find and fix issues in your AI Agents (MIT, open source)

- HN: [48645555](https://news.ycombinator.com/item?id=48645555)
- Source: [github.com](https://github.com/latitude-dev/latitude-llm)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T14:30:54Z

## Translation

タイトル: HN を表示: AI エージェントの問題を見つけて修正する (MIT、オープンソース)
記事タイトル: GitHub - latitude-dev/latitude-llm: Latitude は、オープンソースの AI モニタリング プラットフォームです。 · GitHub
説明: Latitude は、オープンソースの AI モニタリング プラットフォームです。 - 緯度-dev/緯度-llm

記事本文:
GitHub - latitude-dev/latitude-llm: Latitude は、オープンソースの AI モニタリング プラットフォームです。 · GitHub
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
緯度-開発
/
緯度-llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
開発する

オプション ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5,002 コミット 5,002 コミット .agents/ スキル .agents/ スキル .cursor .cursor .github .github .husky .husky .zed .zed apps apps charts/ latitude charts/ latitude dev-docs dev-docs docker docker docs docs fern fern アイデア アイデア インフラ インフラ パッケージ パッケージ スクリプトスクリプト 仕様 仕様 ツール ツール .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json .npmrc .npmrc .tmuxinator.yml .tmuxinator.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json docker-compose.yml docker-compose.yml docker-stack.yml docker-stack.yml docs-overhol-plan.md docs-overhol-plan.md knip.json knip.json misse.toml misse.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml skill-lock.json skill-lock.json tsconfig.base.json tsconfig.base.jsonturbo.jsonturbo.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
新機能: Latitude MCP サーバー: AI エージェントを Latitude に接続 →
オープンソース AI エージェントの監視
セントリーですが、エージェントと LLM 用です。
ウェブサイト ·
ドキュメント ·
変更履歴 ·
たるみ・
×
Latitude は、AI エージェントで次に何が問題になるかを示し、ユーザーが気づく前に修正できるようにします。
問題中心 : 失敗したトレースは、ステータス、サイズ、傾向とともに追跡された問題にグループ化されます。
人間による調整された評価: 時間の経過とともに人間の判断からのドリフトを追跡する調整スコアを使用して、チームの判断に基づいて自動的に構築された評価。
エージェントネイティブ トレース: マルチターン セッション、ツール呼び出し、および完全な実行パスを 1 つのビューに表示します。
セマンティック検索 : フィン

d 意味、完全一致、またはほぼ類似した文による追跡。サンプリングはなく、100% のトレースが検索可能です。
問題アラート : 新しい問題が検出されたとき、または既存の問題がエスカレートしたときに通知を受け取ります。 Slack、電子メール、または Webhook に接続します。
MCP サーバー : Latitude UI で実行できるすべての機能が、コーディング エージェントから MCP 経由で利用できるようになりました。
Latitude は、月あたり 20,000 クレジット、30 日間のデータ保持、無制限のシートを含めて無料で使用できます。
latitude.so にサインアップして、API キーとプロジェクト スラグを取得します。
推奨: コーディング エージェントに問い合わせる
このプロンプトを Claude Code、Cursor、Windsurf、Codex、OpenCode、または別のコーディング エージェントに貼り付けます。
https://raw.githubusercontent.com/latitude-dev/skills/refs/heads/main/skills/latitude-telemetry/SKILL.md から Latitude Telemetry AI スキルを読み、このアプリケーションにトレースを追加します。
TypeScript の手動セットアップ
npm install @latitude-data/telemetry
この例では OpenAI を使用します。これを、アプリがすでにインポートしている LLM SDK に置き換えます。
"@latitude-data/telemetry" から { Latitude } をインポートします。
「openai」から OpenAI をインポートします。
const latitude = 新しい緯度 ( {
apiKey: プロセス。環境 。 LATITUDE_API_KEY ! 、
プロジェクト: プロセス。環境 。 LATITUDE_PROJECT_SLUG ! 、
インストルメンテーション : { openai : OpenAI } 、
} ) ;
const client = new OpenAI() ;
クライアントを待ちます。チャット 。完成品。作成 ( {
モデル：「gpt-4o」、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "こんにちは" } ] 、
} ) ;
緯度を待ちます。シャットダウン ( ) ;
サポートされているすべての LLM 呼び出しが Latitude でトレースとして表示されるようになりました。ユーザー ID、セッション ID、タグ、またはメタデータを追加する場合は、リクエスト、会話、またはエージェントの境界で Capture() を使用します。
Python および OpenTelemetry 互換ランタイムもサポートされています。完全なセットアップ、プロバイダー ガイド、および OTel パススルーについては、「トレース開始ガイド」を参照してください。
Latitude はプロバイダーに依存しません。テレメトリー業務

ほとんどのモデル プロバイダーとフレームワーク ( OpenAI 、 Anthropic 、 Bedrock 、 Vercel AI SDK 、 LangChain など) に加えて、あらゆる OTEL 互換アプリケーションをすぐに利用できます。
セットアップ手順については、完全な統合リストを参照してください。
クロード・コードの内部を構築しますか？完全なセッショントランスクリプトをトレースとしてキャプチャする専用パッケージがあります。
ドキュメントを確認してください。
npx -y @latitude-data/claude-code-telemetry install
ターミナル、デスクトップ アプリ、IDE 拡張機能で動作します。
Latitude への貢献を開始するには、開発セットアップと貢献ガイドを確認してください。
Latitude は、完全にオープンなインフラストラクチャ上で、あらゆる規模で自己ホスト可能です。すぐに使えるコンテナ イメージを Docker Hub にプルします。
単一ホスト (単純) — Docker Compose を備えた 1 台のマシン上の運用グレードのインスタンス。「単一ホスト ガイド」に従ってください。
クラスター (上級) — Helm チャートを介した Kubernetes 上のスケーラブルで高可用性のデプロイメント。クラスター ガイドに従ってください。
ワンクリック (Railway) — 管理されたインフラストラクチャ上のスタック全体、Railway を介した簡単な導入。
Slack コミュニティに参加して、質問したり、フィードバックを共有したり、構築しているものを披露したりできます。
Latitude は MIT ライセンスに基づいてライセンスされています。
貢献は大歓迎です。貢献ガイドを読んで開始し、Slack コミュニティに参加し、問題をオープンするか、プル リクエストを送信してください。
プロジェクトは初めてですか?最初の問題は、始めるのに適した場所です。
🧑‍💻 貢献者の皆様に感謝します
Latitude チームが愛情を込めて作成しました
Latitude は、オープンソースの AI モニタリング プラットフォームです。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
336
フォーク
レポートリポジトリ
リリース
97
TypeScript テレメトリ v3.3.0
最新の
2026 年 6 月 22 日
+ 96 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました

g.このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Latitude is the open-source ai monitoring platform. - latitude-dev/latitude-llm

GitHub - latitude-dev/latitude-llm: Latitude is the open-source ai monitoring platform. · GitHub
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
latitude-dev
/
latitude-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
development Branches Tags Go to file Code Open more actions menu Folders and files
5,002 Commits 5,002 Commits .agents/ skills .agents/ skills .cursor .cursor .github .github .husky .husky .zed .zed apps apps charts/ latitude charts/ latitude dev-docs dev-docs docker docker docs docs fern fern ideas ideas infra infra packages packages scripts scripts specs specs tools tools .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .mcp.json .mcp.json .npmrc .npmrc .tmuxinator.yml .tmuxinator.yml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json docker-compose.yml docker-compose.yml docker-stack.yml docker-stack.yml docs-overhaul-plan.md docs-overhaul-plan.md knip.json knip.json mise.toml mise.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml skills-lock.json skills-lock.json tsconfig.base.json tsconfig.base.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
New: Latitude MCP server: connect your AI agent to Latitude →
Open source AI Agent Monitoring
Sentry, but for agents and LLMs.
Website ·
Docs ·
Changelog ·
Slack ·
X
Latitude shows you what will break next in your AI Agent and helps you fix it before users notice.
Issue-centric : failed traces grouped into tracked issues, with status, size, and trend.
Human-aligned evals : evals built automatically from your team's judgments, with an alignment score that tracks drift from human judgment over time.
Agent-native traces : multi-turn sessions, tool calls, and full execution paths in one view.
Semantic search : find any trace by meaning, exact matches, or roughly similar sentences. No sampling, 100% of traces are searchable.
Issue alerts : get notified the moment a new issue is detected or an existing one escalates. Connect Slack, email, or webhooks.
MCP server : everything you can do in the Latitude UI, now available from your coding agent via MCP.
You can use Latitude for free, including 20K credits/month, 30-day data retention, and unlimited seats.
Sign up at latitude.so and grab your API key and project slug.
Recommended: ask your coding agent
Paste this prompt into Claude Code, Cursor, Windsurf, Codex, OpenCode, or another coding agent:
Read the Latitude Telemetry AI skill from https://raw.githubusercontent.com/latitude-dev/skills/refs/heads/main/skills/latitude-telemetry/SKILL.md and add tracing to this application.
Manual TypeScript setup
npm install @latitude-data/telemetry
This example uses OpenAI; replace it with the LLM SDK your app already imports.
import { Latitude } from "@latitude-data/telemetry" ;
import OpenAI from "openai" ;
const latitude = new Latitude ( {
apiKey : process . env . LATITUDE_API_KEY ! ,
project : process . env . LATITUDE_PROJECT_SLUG ! ,
instrumentations : { openai : OpenAI } ,
} ) ;
const client = new OpenAI ( ) ;
await client . chat . completions . create ( {
model : "gpt-4o" ,
messages : [ { role : "user" , content : "Hello" } ] ,
} ) ;
await latitude . shutdown ( ) ;
Every supported LLM call now shows up as a trace in Latitude. Use capture() at request, conversation, or agent boundaries when you want to add user IDs, session IDs, tags, or metadata.
Python and any OpenTelemetry-compatible runtime are also supported. Full setup, provider guides, and OTel passthrough are in the Start tracing guide .
Latitude is provider-agnostic. Telemetry works out of the box with most model providers and frameworks ( OpenAI , Anthropic , Bedrock , Vercel AI SDK , LangChain , and more), plus any OTEL-compatible application.
See the full integration list for setup instructions.
Building inside Claude Code? We have a dedicated package that captures full session transcripts as traces.
Check out docs.
npx -y @latitude-data/claude-code-telemetry install
Works in the terminal, the Desktop app, and IDE extensions.
Check out the Development setup and the Contributing guide to get started contributing to Latitude.
Latitude is self-hostable at any scale, on fully open infrastructure. Pull the ready-to-go container images on Docker Hub :
Single-host (simple) — a production-grade instance on one machine with Docker Compose, follow the Single-host guide .
Cluster (advanced) — a scalable, highly-available deployment on Kubernetes via a Helm chart, follow the Cluster guide .
One-click (Railway) — the whole stack on managed infrastructure, easy deploy through Railway .
Join the Slack community to ask questions, share feedback, and show what you're building.
Latitude is licensed under the MIT License .
Contributions are welcome. Read the Contributing Guide to get started, then join the Slack community , open an issue , or submit a pull request.
New to the project? Good first issues are a friendly place to start.
🧑‍💻 Thanks to all of our contributors
Made with love by the Latitude Team
Latitude is the open-source ai monitoring platform.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
336
forks
Report repository
Releases
97
TypeScript Telemetry v3.3.0
Latest
Jun 22, 2026
+ 96 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

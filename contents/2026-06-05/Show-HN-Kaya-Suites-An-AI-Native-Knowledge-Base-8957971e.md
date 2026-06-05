---
source: "https://github.com/Kaya-suites/Kaya-Suites"
hn_url: "https://news.ycombinator.com/item?id=48410745"
title: "Show HN: Kaya Suites – An AI Native Knowledge Base"
article_title: "GitHub - Kaya-suites/Kaya-Suites · GitHub"
author: "kyahwill"
captured_at: "2026-06-05T13:10:14Z"
capture_tool: "hn-digest"
hn_id: 48410745
score: 1
comments: 0
posted_at: "2026-06-05T11:01:56Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Kaya Suites – An AI Native Knowledge Base

- HN: [48410745](https://news.ycombinator.com/item?id=48410745)
- Source: [github.com](https://github.com/Kaya-suites/Kaya-Suites)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T11:01:56Z

## Translation

タイトル: HN を表示: Kaya Suites – AI ネイティブのナレッジ ベース
記事タイトル: GitHub - Kaya-suites/Kaya-Suites · GitHub
説明: GitHub でアカウントを作成して、Kaya-suites/Kaya-Suites の開発に貢献します。
HN テキスト: 私は人間だけでなくエージェントにとってもネイティブな知識ベースを生成するというこのアイデアに非常に興味があります。私は Den の [ https://www.producthunt.com/products/den ] 製品にかなりインスピレーションを受けており、同様にオープンソース化されたそれに似た製品を生成したいと考えています。私の仮説では、近い将来、AI エージェントと人間のエージェントが連携することにより、AI エージェントと人間のエージェントが組織活動や業務活動の参考となる情報を一元的に保管する必要があると考えています。これまでのプロジェクトの方向性に関しては、今実際にまともな問題に取り組んでいるのか、それとも別のことに軸足を移すべきなのかを確認することに熱心です。お気軽にフィードバックをお寄せください。ぜひ聞いてみたいです。

記事本文:
GitHub - カヤ スイート/カヤ スイート · GitHub
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
カヤスイート
/
カヤ スイーツ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクション メニューを開く 折りたたむ

ファイルとファイル
132 コミット 132 コミット .claude .claude .github .github apps apps docs docs package scripts scripts .gitignore .gitignore CONFIG.md CONFIG.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md TODO.md TODO.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントが生きた真実の情報源としてドキュメントを積極的に管理する AI ネイティブのナレッジ ベース。
このプロジェクトは、単一の Apache 2.0 ディストリビューションとして出荷されます。Rust バックエンド (Axum + Cargo ワークスペース) と、生成された TypeScript OpenAPI クライアントによって結合された Next.js 16 フロントエンドです。
3 つのエージェントのパイプライン。オーケストレーターが各チャット ターンを分類し、読み取り専用の研究者が証拠を収集し、書き込み専用の編集者が編集を提案します。読み取り/書き込み分離は、trybuild テストを通じてコン​​パイル時に強制されます。
編集を提案してから承認します。エージェントはストレージを直接変更できません。編集は、SSE 経由でストリーミングされる ProposedEdit イベントになります。明示的なユーザー承認のみが、commit_edit が受け入れる ApprovalToken を生成します。
プラグイン可能なストレージ。単一の StorageAdapter トレイトは、ツリー内の 3 つのバックエンド (SQLite (デフォルト)、Postgres、および MySQL) を駆動します。
プラグ可能な LLM プロバイダー。 OpenAI、Anthropic、Gemini、およびテスト用の MockProvider に対する LlmProvider + ModelRouter 抽象。ルーティングは、 kaya.yaml を介して構成主導で行われます。
Notion スタイルのブロック エディター。型付き Markdown ブロック モデル ( @kaya/markdown-model ) 上に構築されたワークスペース パッケージ ( @kaya/markdown-editor ) として出荷されます。
アプリ/
web/Next.js 16 フロントエンド (pnpm)
backend/ Rust バックエンド、Cargo ワークスペース ルート
木箱/
kaya-core/ トレイト、エージェント パイプライン、編集プリミティブ
kaya-storage/ SQLite + Postgres + MySQL アダプター
kaya-server/ Axum HTTP ルート + SSE チャット
kaya-auth/AuthAdapter sc

アフォールズ（マジックリンク、パスワード）
kaya-metering/トークン + 支出計測、レート制限
kaya-db/ 共有 DB ユーティリティ
ビン/
kaya-oss/ 自己ホスト型バイナリ。フロントエンドを埋め込むことができる
パッケージ/
api-client/ 生成された TypeScript クライアント (OpenAPI スキーマから)
markdown-editor/Notion 風のブロックエディター (@kaya/markdown-editor)
markdown-model/ Markdown ブロック モデル + パーサー (@kaya/markdown-model)
ui/共有 React プリミティブ
スクリプト/
docs/Apache 2.0 ドキュメント
Rust ワークスペースのルートは apps/backend/Cargo.toml です。リポジトリのルートには Cargo.toml はなく、apps/web/ は Cargo ワークスペースの一部ではありません。
SQLite システム ライブラリ (デフォルトのストレージ)
CD アプリ/バックエンド
カーゴビルド --workspace
カーゴラン --bin kaya-oss
OSS バイナリは、デフォルトで http://localhost:3001 で HTTP API を提供します。
pnpm install # リポジトリのルートから
pnpm dev # すべてのワークスペース開発スクリプトを並行して実行します
http://localhost:3000 を開きます。 Web アプリは、API 呼び出しを NEXT_PUBLIC_API_URL (デフォルト http://localhost:3001 ) にプロキシします。
最初の起動時にバックエンドによってスーパー管理者アカウントがシードされるため、すぐにサインインできます。
最初のログイン後にアカウント設定ページからパスワードを変更します。これらの資格情報はローカル開発のみを目的としています。
バックエンドのルートまたはスキーマを変更した後:
CD アプリ/バックエンド
カーゴ実行 --bin kaya-oss -- --schema > ../../packages/api-client/openapi.json
CD ../../
pnpm生成
openapi.json と更新された package/api-client/src/ の両方を一緒にコミットする必要があります。
変数
デフォルト
目的
NEXT_PUBLIC_API_URL
http://ローカルホスト:3001
フロントエンドで使用されるバックエンドのベース URL
OPENAI_API_KEY
—
OpenAIプロバイダーに必要
ANTHROPIC_API_KEY
—
Anthropic プロバイダーが必要とする
GEMINI_API_KEY
—
Gemini プロバイダーに必要
LLM ルーティングは apps/backend/kaya.yaml で構成されます。 docs/llm-provider.md および CONFIG.md を参照してください。
ドキュメント

メント
説明
建築
システムの概要、クレート、2 ビルド システムのレイアウト
エージェントのアーキテクチャ
オーケストレーター/リサーチャー/エディターのパイプラインと SSE 契約
ストレージアダプター
StorageAdapter 特性、SQLite/Postgres/MySQL バックエンド
認証アダプター
AuthAdapter の特性と現在のスキャフォールド
LLMプロバイダー
LlmProvider 、 ModelRouter 、ルーティング構成
APIコード生成
OpenAPI スキーマ → TypeScript クライアント パイプライン
パッケージ
フロントエンドワークスペースパッケージ
建物
ビルド、テストコマンド、OSS 静的バイナリ
構成
ルーティング構成、環境変数、ストレージ バックエンドの選択
ライセンス
このリポジトリ内のすべてに Apache 2.0。 LICENSE および CONTRIBUTING.md を参照してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Kaya-suites/Kaya-Suites development by creating an account on GitHub.

I'm really interested in this idea of generating a knowledge base that's native for agents as well as humans. I'm pretty much inspired by Den's [ https://www.producthunt.com/products/den ] product and I want to generate something akin to it that is open-sourced as well. I hypothesize that, in the near future, with AI agents collaborating with human agents, there will be a need for a centralized storage of information that AI agents and human agents could use as a reference for their organizational and work activities. As for the direction of my project so far, I'm keen on checking on whether I'm actually working on a decent problem right now or whether I should pivot to something else. Feel free to provide some feedback. I'd love to hear it.

GitHub - Kaya-suites/Kaya-Suites · GitHub
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
Kaya-suites
/
Kaya-Suites
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
132 Commits 132 Commits .claude .claude .github .github apps apps docs docs packages packages scripts scripts .gitignore .gitignore CONFIG.md CONFIG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md TODO.md TODO.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
An AI-native knowledge base where an agent actively maintains documents as a living source of truth.
The project ships as a single Apache 2.0 distribution: a Rust backend (Axum + Cargo workspace) and a Next.js 16 frontend, glued by a generated TypeScript OpenAPI client.
Three-agent pipeline. An orchestrator classifies each chat turn, a read-only researcher gathers evidence, and a write-only editor proposes edits. Read/write isolation is enforced at compile time via trybuild tests.
Propose-then-approve edits. No agent can mutate storage directly. Edits become ProposedEdit events streamed over SSE; only an explicit user approval mints an ApprovalToken that commit_edit will accept.
Pluggable storage. A single StorageAdapter trait drives three backends in-tree: SQLite (default), Postgres, and MySQL.
Pluggable LLM provider. LlmProvider + ModelRouter abstract over OpenAI, Anthropic, Gemini, and a MockProvider for tests. Routing is config-driven via kaya.yaml .
Notion-style block editor. Shipped as a workspace package ( @kaya/markdown-editor ) built on a typed Markdown block model ( @kaya/markdown-model ).
apps/
web/ Next.js 16 frontend (pnpm)
backend/ Rust backend, Cargo workspace root
crates/
kaya-core/ Traits, agent pipeline, edit primitives
kaya-storage/ SQLite + Postgres + MySQL adapters
kaya-server/ Axum HTTP routes + SSE chat
kaya-auth/ AuthAdapter scaffolds (magic link, password)
kaya-metering/ Token + spend metering, rate limits
kaya-db/ Shared DB utilities
bin/
kaya-oss/ Self-hosted binary; can embed the frontend
packages/
api-client/ Generated TypeScript client (from OpenAPI schema)
markdown-editor/ Notion-like block editor (@kaya/markdown-editor)
markdown-model/ Markdown block model + parser (@kaya/markdown-model)
ui/ Shared React primitives
scripts/
docs/ Apache 2.0 documentation
The Rust workspace root is apps/backend/Cargo.toml . There is no Cargo.toml at the repo root and apps/web/ is not part of the Cargo workspace.
A SQLite system library (default storage)
cd apps/backend
cargo build --workspace
cargo run --bin kaya-oss
The OSS binary serves the HTTP API on http://localhost:3001 by default.
pnpm install # from repo root
pnpm dev # runs all workspace dev scripts in parallel
Open http://localhost:3000 . The web app proxies API calls to NEXT_PUBLIC_API_URL (default http://localhost:3001 ).
On first boot the backend seeds a superadmin account so you can sign in immediately:
Change the password from the account settings page after your first login — these credentials are intended for local development only.
After changing a backend route or schema:
cd apps/backend
cargo run --bin kaya-oss -- --schema > ../../packages/api-client/openapi.json
cd ../../
pnpm generate
Both openapi.json and the updated packages/api-client/src/ should be committed together.
Variable
Default
Purpose
NEXT_PUBLIC_API_URL
http://localhost:3001
Backend base URL used by the frontend
OPENAI_API_KEY
—
Required by the OpenAI provider
ANTHROPIC_API_KEY
—
Required by the Anthropic provider
GEMINI_API_KEY
—
Required by the Gemini provider
LLM routing is configured in apps/backend/kaya.yaml . See docs/llm-provider.md and CONFIG.md .
Document
Description
Architecture
System overview, crates, two-build-system layout
Agent architecture
Orchestrator / Researcher / Editor pipeline and SSE contract
Storage adapter
StorageAdapter trait, SQLite/Postgres/MySQL backends
Auth adapter
AuthAdapter trait and current scaffolds
LLM provider
LlmProvider , ModelRouter , routing config
API codegen
OpenAPI schema → TypeScript client pipeline
Packages
Frontend workspace packages
Building
Builds, test commands, OSS static binary
Configuration
Routing config, env vars, storage backend selection
License
Apache 2.0 for everything in this repository. See LICENSE and CONTRIBUTING.md .
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/azhakhan/dribble"
hn_url: "https://news.ycombinator.com/item?id=48727039"
title: "Show HN: Dribble – an open-source AI IDE for databases"
article_title: "GitHub - azhakhan/dribble: AI-powered, open source IDE for Databases · GitHub"
author: "ayazhan"
captured_at: "2026-06-30T00:30:24Z"
capture_tool: "hn-digest"
hn_id: 48727039
score: 1
comments: 0
posted_at: "2026-06-30T00:16:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dribble – an open-source AI IDE for databases

- HN: [48727039](https://news.ycombinator.com/item?id=48727039)
- Source: [github.com](https://github.com/azhakhan/dribble)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T00:16:39Z

## Translation

タイトル: Show HN: Dribble – データベース用のオープンソース AI IDE
記事のタイトル: GitHub - azhakhan/dribble: AI を活用したデータベース用のオープンソース IDE · GitHub
説明: AI を活用したデータベース用のオープンソース IDE。 GitHub でアカウントを作成して、azhakhan/dribble の開発に貢献してください。

記事本文:
GitHub - azhakhan/dribble: AI を活用したデータベース用のオープンソース IDE · GitHub
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
アザカン
/
ドリブル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Cod

e その他のアクション メニューを開く フォルダーとファイル
406 コミット 406 コミット .claude .claude app アプリコンポーネント コンポーネント docs docs lib lib public publictypes タイプ .DS_Store .DS_Store .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md Screenshot.png Screenshot.png auth.config.ts auth.config.ts drizzle.config.ts drizzle.config.ts eslint.config.mjs eslint.config.mjs launch.json launch.json next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs proxy.ts proxy.ts tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI を活用したデータベース用のオープンソース SQL IDE。
Dribble は、AI データ アナリストが組み込まれた Web ベースの SQL IDE です。に接続します
Postgres データベース、そのスキーマの参照、ノートブックでのクエリの実行、テーブルの探索
並べ替え/フィルター/ページネーションを使用し、データについて AI エージェントに質問する - すべて
1 つのタブ付きワークスペースで、中断した場所を記憶します。
AI データ アナリスト — あなたのデータを検査するエージェント (Claude Opus 4.8) とチャットします。
スキーマを作成し、読み取り専用 SQL を作成して実行し、エラーを反復処理して、
最終的な結果セットはテーブルとして設定されます。
SQL ノートブック — Monaco エディターで構文を使用してクエリを作成および実行します
ハイライト。 Cmd/Ctrl + Enter で実行します。ノートブックとその結果は次のとおりです。
保存されました。
スキーマ ブラウザ — 折りたたみ可能なサイドバーからスキーマとテーブルに移動します
木。
テーブル エクスプローラー — サーバー側のページネーション、列を使用してテーブル データを参照します。
並べ替え、および生の WHERE 句フィルター。
高速な結果グリッド — 大規模な結果セットが仮想化されたデータ グリッドでレンダリングされます。
永続的なワークスペース — 開いているタブ、レイアウト/パネルのサイズ、展開されたツリー、
キャッシュされたクエリ/チャット結果はリロード後も残ります (ブラウザー間で追跡されます)。
状態はサーバー側に保存されるため)。
S

マート接続のライフサイクル - データベースドライバーは使用中にウォーム状態に保たれます
そうでない場合はアイドル状態になり、サイドバーにライブ接続ステータスが反映されます。
柔軟な認証 — ローカルで使用する場合、または Google の背後で使用する場合は、まったくログインせずに実行されます。
マルチユーザー展開のサインイン (電子メール/ドメイン許可リストを使用)。
人のつながり、ノート、チャットは非公開です。保存されたデータベース
認証情報は保存時に暗号化されます。参照
docs/authentication.md 。
プラグイン可能なドライバー — Postgres は本日出荷されます。ドライバーレジストリは以下のように構築されています
さらにエンジンを追加します (MySQL、Snowflake など)。
Next.js 16 · React 19 · TypeScript · Tailwind CSS 4 · モナコエディタ ·
glide-data-grid · Zustand · Vercel AI SDK ( @ai-sdk/anthropic ) · Postgres ( pg )
アプリのメタデータ (接続、ノートブック、チャット) を保存するための Postgres データベース
歴史）。ローカル、Neon、Supabase、Vercel Postgres など、あらゆる Postgres が動作します。
AI エージェントの Anthropic API キー。
git clone < your-repo-url > ドリブル
CDドリブル
npmインストール
設定する
サンプルの env ファイルをコピーし、値を入力します。
cp .env.example .env.local
# メタデータ ストレージ (接続、ノートブック、チャット履歴)。
# すべての Postgres が動作します — Vercel Postgres / Neon / Supabase / local。
DATABASE_URL=postgres://user:pass@host:5432/dribble
# 保存された DB 認証情報の暗号化 (および認証セッションへの署名) に使用されるシークレット。
＃ 必須。 openssl rand -hex 32 で生成します。
APP_SECRET=
# AI チャット エージェント (claude-opus-4-8) を強化します。
ANTHROPIC_API_KEY=
ローカルで実行するために必要なのはこれだけです。認証が設定されていない状態でアプリが起動します。
ログイン画面はなく、すべてのデータは単一の組み込みユーザーに属します。
必要なメタデータ テーブルは、初回実行時に自動的に作成されます。
ログインを要求し、各ユーザーのデータを非公開にするには、Google OAuth を構成します。「」を参照してください。
完全なセットアップについては docs/authentication.md を参照してください。要するに、
に追加します。

環境ローカル:
AUTH_GOOGLE_ID=
AUTH_GOOGLE_SECRET=
# サインインできるユーザーを制限します (すべての Google アカウントを許可するには空のままにします):
AUTH_ALLOWED_EMAILS=you@example.com
AUTH_ALLOWED_DOMAIN=example.com
<origin>/api/auth/callback/google を承認されたリダイレクト URI として
Google OAuth クライアント。これらを設定すると、ログイン画面が自動的に有効になります。
npm 実行開発
http://localhost:3000 を開きます (そうでない場合は直接アクセスします)
認証が設定されています。それ以外の場合は、Google でサインインします)、データベース接続を追加します。
そしてクエリを開始します。
実稼働サーバーを構築して実行するには、次の手順を実行します。
npm ビルドを実行する
npmスタート
AI が生成したコードに関する注意事項
このプロジェクトは主に AI コーディング ツール (Claude Code) の助けを借りて書かれました。
すべてのコードはコミット前にレビューされていますが、レビューする必要があります。
実稼働環境でそれに依存する前に、自分自身で確認してください。
MIT ライセンスに基づいてリリースされています。
AI を活用したデータベース用のオープンソース IDE
読み込み中にエラーが発生しました。このページをリロードしてください。
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

AI-powered, open source IDE for Databases. Contribute to azhakhan/dribble development by creating an account on GitHub.

GitHub - azhakhan/dribble: AI-powered, open source IDE for Databases · GitHub
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
azhakhan
/
dribble
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
406 Commits 406 Commits .claude .claude app app components components docs docs lib lib public public types types .DS_Store .DS_Store .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md Screenshot.png Screenshot.png auth.config.ts auth.config.ts drizzle.config.ts drizzle.config.ts eslint.config.mjs eslint.config.mjs launch.json launch.json next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs proxy.ts proxy.ts tsconfig.json tsconfig.json View all files Repository files navigation
An AI-powered, open-source SQL IDE for your databases.
Dribble is a web-based SQL IDE with a built-in AI data analyst. Connect to a
Postgres database, browse its schema, run queries in a notebook, explore tables
with sort/filter/pagination, and ask an AI agent questions about your data — all
in one tabbed workspace that remembers where you left off.
AI data analyst — chat with an agent (Claude Opus 4.8) that inspects your
schema, writes and runs read-only SQL, iterates on errors, and renders the
final result set as a table.
SQL notebooks — write and execute queries in a Monaco editor with syntax
highlighting. Run with Cmd/Ctrl + Enter . Notebooks and their results are
saved.
Schema browser — navigate schemas and tables from a collapsible sidebar
tree.
Table explorer — browse table data with server-side pagination, column
sorting, and a raw WHERE -clause filter.
Fast results grid — large result sets render in a virtualized data grid.
Persistent workspace — open tabs, layout/panel sizes, the expanded tree,
and cached query/chat results survive reloads (and follow you across browsers,
since state is stored server-side).
Smart connection lifecycle — database drivers are kept warm while in use
and idle out when not, with the sidebar reflecting live connection status.
Flexible auth — runs with no login at all for local use, or behind Google
sign-in (with an email/domain allowlist) for multi-user deployments, where each
person's connections, notebooks, and chats are private. Stored database
credentials are encrypted at rest. See
docs/authentication.md .
Pluggable drivers — Postgres ships today; the driver registry is built to
add more engines (MySQL, Snowflake, …).
Next.js 16 · React 19 · TypeScript · Tailwind CSS 4 · Monaco Editor ·
glide-data-grid · Zustand · Vercel AI SDK ( @ai-sdk/anthropic ) · Postgres ( pg )
A Postgres database for storing app metadata (connections, notebooks, chat
history). Any Postgres works — local, Neon, Supabase, Vercel Postgres, etc.
An Anthropic API key for the AI agent.
git clone < your-repo-url > dribble
cd dribble
npm install
Configure
Copy the example env file and fill in the values:
cp .env.example .env.local
# Metadata storage (connections, notebooks, chat history).
# Any Postgres works — Vercel Postgres / Neon / Supabase / local.
DATABASE_URL=postgres://user:pass@host:5432/dribble
# Secret used to encrypt stored DB credentials (and sign the auth session).
# Required. Generate with: openssl rand -hex 32
APP_SECRET=
# Powers the AI chat agent (claude-opus-4-8).
ANTHROPIC_API_KEY=
That's all you need to run locally — with no auth configured, the app starts
without a login screen and all data belongs to a single built-in user.
The required metadata tables are created automatically on first run.
To require login and keep each user's data private, configure Google OAuth — see
docs/authentication.md for the full setup. In short,
add to .env.local :
AUTH_GOOGLE_ID=
AUTH_GOOGLE_SECRET=
# Restrict who may sign in (leave empty to allow any Google account):
AUTH_ALLOWED_EMAILS=you@example.com
AUTH_ALLOWED_DOMAIN=example.com
Register <origin>/api/auth/callback/google as an authorized redirect URI on the
Google OAuth client. Setting these enables the login screen automatically.
npm run dev
Open http://localhost:3000 (you're in directly when no
auth is configured; otherwise sign in with Google), add a database connection,
and start querying.
To build and run a production server:
npm run build
npm start
A note on AI-generated code
This project was written largely with the help of AI coding tools (Claude Code).
All code has been reviewed before being committed, but you should review it
yourself before relying on it in production.
Released under the MIT License .
AI-powered, open source IDE for Databases
There was an error while loading. Please reload this page .
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

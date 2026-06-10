---
source: "https://github.com/hashwnath/KMCP"
hn_url: "https://news.ycombinator.com/item?id=48469712"
title: "Show HN: KnowledgeMCP – Turn any docs into an MCP endpoint (0 LLM at query time)"
article_title: "GitHub - hashwnath/KMCP: Open-source MCP server for your docs. Zero LLM at query time. docker compose up and go. · GitHub"
author: "Hashwanths"
captured_at: "2026-06-10T00:59:25Z"
capture_tool: "hn-digest"
hn_id: 48469712
score: 2
comments: 0
posted_at: "2026-06-10T00:36:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: KnowledgeMCP – Turn any docs into an MCP endpoint (0 LLM at query time)

- HN: [48469712](https://news.ycombinator.com/item?id=48469712)
- Source: [github.com](https://github.com/hashwnath/KMCP)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T00:36:53Z

## Translation

タイトル: HN の表示: KnowledgeMCP – あらゆるドキュメントを MCP エンドポイントに変換します (クエリ時に 0 LLM)
記事のタイトル: GitHub - hashwnath/KMCP: ドキュメント用のオープンソース MCP サーバー。クエリ時の LLM はゼロです。 docker を作成して実行します。 · GitHub
説明: ドキュメント用のオープンソース MCP サーバー。クエリ時の LLM はゼロです。 docker を作成して実行します。 - ハッシュナス/KMCP

記事本文:
GitHub - hashwnath/KMCP: ドキュメント用のオープンソース MCP サーバー。クエリ時の LLM はゼロです。 docker を作成して実行します。 · GitHub
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
ハッシュナス
/
KMCP
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション最適化

オン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github .github docs docs フロントエンド フロントエンド インフラ インフラ スクリプト スクリプト src src テスト テスト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SCAFFOLD.md SCAFFOLD.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml要件.txt要件.txt すべてのファイルを表示リポジトリ ファイルのナビゲーション
ドキュメントに MCP エンドポイントを与えます。すべての AI エージェントがそれらを使用できます。
ドラフト.mp4
KnowledgeMCP は、あらゆるドキュメント ソース (Web サイト、PDF、Confluence、Notion、S3、GitHub) を標準準拠の Model Context Protocol (MCP) エンドポイントに変換します。 Claude、GitHub Copilot、Cursor、およびその他の MCP 互換エージェントは、クエリ時に LLM 呼び出しを行わずに、これらのドキュメントを即座に検索して読み取ることができます (OpenSearch では小さなローカル埋め込みモデル + ハイブリッド BM25/kNN 検索を使用しています)。
🔌 MCP ネイティブ — エージェントが接続できる 3 つのツール ( docs_search 、 code_sample_search 、 docs_fetch )
💰 コストゼロのクエリパス — ローカル埋め込み + OpenSearch ハイブリッド検索。クエリごとの OpenAI/Bedrock 料金はかかりません。
🐳 docker compose up は機能します — 完全にローカルで実行され、AWS アカウントやクレジット カードは必要ありません
☁️ 必要なときに本番環境に対応できる AWS パス — Lambda + DynamoDB + SQS + S3 + バンドルされた SAM テンプレートを介したマネージド OpenSearch
git clone https://github.com/hashwnath/KMCP.git
cd KMCP
構成 # docker 構成 -d --build
次に:
ダッシュボード → http://localhost:3000 (サインアップ → ソースの追加 → 検索)
管理REST API → http://localhost:8081
MCP エンドポイント → http://localhost:8000/mcp/{your-tenant-slug}
初回起動時のダウンロード

彼はモデル (約 30 MB) と OpenSearch (約 700 MB のイメージ) を高速に埋め込みました。
テナント URL で MCP クライアントを指定します。
{
"mcpサーバー": {
"MyDocs" : {
"url" : " http://localhost:8000/mcp/your-tenant-slug " ,
"タイプ" : " http "
}
}
}
エージェントは 3 つのツールを利用できます。
┌─────────────────────────┐
│ AI エージェント (クロード、カーソル、副操縦士、コンティニュー、...) │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ POST /mcp/{tenant_slug}
┌───────────▼─────────────────┐
│ MCP サーバー (FastMCP) — docs_search / code_search / fetch │
━━━━━━━━━━━━━━━━━━━━━━━━━━┘
┌─────────┼─────────┐
▼ ▼ ▼
┌───────────┐ ┌─────┐ ┌─────────┐
│ OpenSearch │ │ SQLite │ │ ファイルシステム │
│ (BM25 + kNN) │ │ テナント │ │ BLOB │
│ ~768 トークン │ │ ソース │ │ アップロード │
│ チャンク │ │ ジョブ │ │ │
━━━━━━━━━━━━━━━━━━━━━━━┘ ━━━━━━━━┘
▲
┌───────────

─┴───────────────┐
│ 管理 API (Starlette) + バックグラウンド ワーカー │
│ サインアップ/ログイン (JWT) クロール → マークダウン → チャンク → │
│ ソース CRUD 埋め込み → OpenSearch │
│ 分析 │
━━━━━━━━━━━━━━━━━━━━━━━┘
(AWS モードでは、SQLite → DynamoDB、ファイルシステム → S3、ワーカー キュー → SQS、
各サービスを独自の Lambda として実行します。アプリケーションコードは変更されていません
すべての AWS 呼び出しは src/common/backends/ を介してルーティングされるためです。)
種類
摂取するもの
ウェブサイトの URL
サイトマップ全体のクロール → マークダウン
テキストの貼り付け
インラインテキスト
ファイルアップロード
PDF、DOCX、PPTX、MD、HTML、TXT
クラウドストレージ
S3、Azure BLOB、GCS
wiki_kb
Confluence、Notion、SharePoint、GitBook
git_repo
パブリックまたはプライベート GitHub/GitLab リポジトリ (トークンはオプション)
構成
デフォルトはローカルの docker-compose で機能します。カスタマイズするには、.env.example を .env にコピーして編集します。最も便利なノブ:
SAM テンプレート (Lambda + DynamoDB + SQS + S3 + OpenSearch + SES)、コスト見積もり、運用ランブックについては、docs/AWS_DEPLOYMENT.md を参照してください。
PR の方は大歓迎です。コードベース ツアーとローカル開発セットアップについては、CONTRIBUTING.md を参照してください。
make test #完全な pytest スイート (BACKEND=local)
make test-aws # AWS モックスイート
構成 # docker 構成 -d --build
ライセンス
バックエンド ( src/ 、 infra/ 、トップレベル構成) — AGPL-3.0
AGPL-3.0 ライセンスは、ホスト/SaaS での使用は同じライセンスに基づいて変更を公開する必要があることを意味します。それがあなたのユースケースにとって問題である場合は、商用ライセンスについて話し合うことができるよう、問題を開いてください。
ナレッジMCP
典型的な RAG ツール
クエリコスト
$0 (ローカル埋め込み + OpenSearch)
$0.01 ～ 0.10/クエリ (LLM 再ランキング)
エージェント

統合
ネイティブ MCP — プラグアンドプレイ
REST API + カスタム グルー コード
自己ホスト型
docker compose up、クラウドアカウントなし
通常はクラウド API が必要です
マルチテナント
テナントごとの分離機能が組み込まれています
シングルテナント、後からボルトオンで使用可能
レイテンシ
~100ms (パスに LLM なし)
1-5 秒 (LLM 再ランキング)
コミュニティ
GitHub ディスカッション — 質問、アイデア、ショーアンドテル
問題 — バグレポート、機能リクエスト
FastMCP — MCP サーバー フレームワーク
fastembed — ONNX ランタイム埋め込みライブラリ
OpenSearch — ハイブリッド BM25 + kNN 検索
Microsoft Learn MCP サーバー チーム — ミドルウェア経由のテナント コンテキスト設計を形成するために苦労して得た教訓を文書化します
ドキュメント用のオープンソース MCP サーバー。クエリ時の LLM はゼロです。 docker を作成して実行します。
github.com/hashwnath/KMCP#quick-start
トピックス
AGPL-3.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0 — 最初のオープンソース リリース
最新の
2026 年 6 月 7 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source MCP server for your docs. Zero LLM at query time. docker compose up and go. - hashwnath/KMCP

GitHub - hashwnath/KMCP: Open-source MCP server for your docs. Zero LLM at query time. docker compose up and go. · GitHub
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
hashwnath
/
KMCP
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github .github docs docs frontend frontend infra infra scripts scripts src src tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SCAFFOLD.md SCAFFOLD.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
Give your docs an MCP endpoint. Every AI agent can use them.
draft.mp4
KnowledgeMCP turns any documentation source (websites, PDFs, Confluence, Notion, S3, GitHub) into a standards-compliant Model Context Protocol (MCP) endpoint . Claude, GitHub Copilot, Cursor, and any other MCP-compatible agent can search and read those docs instantly — with no LLM calls at query time (we use a tiny local embedding model + hybrid BM25/kNN search in OpenSearch).
🔌 MCP-native — three tools ( docs_search , code_sample_search , docs_fetch ) any agent can plug into
💰 Zero-cost query path — local embeddings + OpenSearch hybrid search. No OpenAI/Bedrock fees per query.
🐳 docker compose up works — runs fully local, no AWS account, no credit card
☁️ Production-ready AWS path when you want it — Lambda + DynamoDB + SQS + S3 + managed OpenSearch via the bundled SAM template
git clone https://github.com/hashwnath/KMCP.git
cd KMCP
make up # docker compose up -d --build
Then:
Dashboard → http://localhost:3000 (signup → add a source → search)
Admin REST API → http://localhost:8081
MCP endpoint → http://localhost:8000/mcp/{your-tenant-slug}
First-time start downloads the fastembed model (~30 MB) and OpenSearch (~700 MB image).
Point any MCP client at your tenant URL:
{
"mcpServers" : {
"MyDocs" : {
"url" : " http://localhost:8000/mcp/your-tenant-slug " ,
"type" : " http "
}
}
}
The agent gets three tools:
┌────────────────────────────────────────────────────────────┐
│ AI Agents (Claude, Cursor, Copilot, Continue, ...) │
└──────────────────────────┬─────────────────────────────────┘
│ POST /mcp/{tenant_slug}
┌──────────────────────────▼─────────────────────────────────┐
│ MCP Server (FastMCP) — docs_search / code_search / fetch │
└──────────────────────────┬─────────────────────────────────┘
┌──────────────┼──────────────┐
▼ ▼ ▼
┌───────────────┐ ┌──────────┐ ┌──────────────┐
│ OpenSearch │ │ SQLite │ │ Filesystem │
│ (BM25 + kNN) │ │ tenants │ │ blobs │
│ ~768 token │ │ sources │ │ uploads │
│ chunks │ │ jobs │ │ │
└───────────────┘ └──────────┘ └──────────────┘
▲
┌──────────────────────────┴─────────────────────────────────┐
│ Admin API (Starlette) + Background Worker │
│ signup/login (JWT) crawl → markdown → chunk → │
│ sources CRUD embed → OpenSearch │
│ analytics │
└────────────────────────────────────────────────────────────┘
(In AWS mode, swap SQLite → DynamoDB, Filesystem → S3, the worker queue → SQS,
and run each service as its own Lambda. The application code is unchanged
because every AWS call routes through src/common/backends/ .)
Type
What it ingests
website_url
Full sitemap crawl → markdown
paste_text
Inline text
file_upload
PDF, DOCX, PPTX, MD, HTML, TXT
cloud_storage
S3, Azure Blob, GCS
wiki_kb
Confluence, Notion, SharePoint, GitBook
git_repo
Public or private GitHub/GitLab repos (token optional)
Configuration
Defaults work for local docker-compose. To customise, copy .env.example to .env and edit. The most useful knobs:
See docs/AWS_DEPLOYMENT.md for the SAM template (Lambda + DynamoDB + SQS + S3 + OpenSearch + SES), cost estimate, and operational runbook.
PRs welcome. See CONTRIBUTING.md for the codebase tour and local dev setup.
make test # full pytest suite (BACKEND=local)
make test-aws # AWS-mocked suite
make up # docker compose up -d --build
License
Backend ( src/ , infra/ , top-level configs) — AGPL-3.0
The AGPL-3.0 license means hosted/SaaS use must publish modifications under the same license. If that's a problem for your use case, please open an issue so we can discuss commercial licensing.
KnowledgeMCP
Typical RAG tools
Query cost
$0 (local embeddings + OpenSearch)
$0.01-0.10/query (LLM reranking)
Agent integration
Native MCP — plug and play
REST API + custom glue code
Self-hosted
docker compose up , no cloud account
Usually needs cloud APIs
Multi-tenant
Per-tenant isolation built-in
Single-tenant, bolt-on later
Latency
~100ms (no LLM in path)
1-5s (LLM reranking)
Community
GitHub Discussions — questions, ideas, show-and-tell
Issues — bug reports, feature requests
FastMCP — the MCP server framework
fastembed — ONNX-runtime embedding library
OpenSearch — hybrid BM25 + kNN search
Microsoft Learn MCP server team — for documenting hard-earned lessons that shaped the tenant-context-via-middleware design
Open-source MCP server for your docs. Zero LLM at query time. docker compose up and go.
github.com/hashwnath/KMCP#quick-start
Topics
AGPL-3.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0 — First open-source release
Latest
Jun 7, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

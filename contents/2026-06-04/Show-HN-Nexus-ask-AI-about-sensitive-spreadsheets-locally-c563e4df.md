---
source: "https://github.com/scottshapiro142/nexuscli"
hn_url: "https://news.ycombinator.com/item?id=48403463"
title: "Show HN: Nexus, ask AI about sensitive spreadsheets locally"
article_title: "GitHub - scottshapiro142/nexuscli: Local-first agent layer for tabular data. Drop any spreadsheet — get a local MCP server your AI agent can query. · GitHub"
author: "scottbuilds"
captured_at: "2026-06-04T21:38:14Z"
capture_tool: "hn-digest"
hn_id: 48403463
score: 2
comments: 0
posted_at: "2026-06-04T19:27:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nexus, ask AI about sensitive spreadsheets locally

- HN: [48403463](https://news.ycombinator.com/item?id=48403463)
- Source: [github.com](https://github.com/scottshapiro142/nexuscli)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T19:27:27Z

## Translation

タイトル: HN を表示: Nexus、機密性の高いスプレッドシートについてローカルで AI に質問する
記事のタイトル: GitHub - scottshapiro142/nexuscli: 表形式データのローカルファーストエージェント層。スプレッドシートをドロップします。AI エージェントがクエリできるローカル MCP サーバーを取得します。 · GitHub
説明: 表形式データのローカルファーストエージェント層。スプレッドシートをドロップします。AI エージェントがクエリできるローカル MCP サーバーを取得します。 - scottshapiro142/nexuscli

記事本文:
GitHub - scottshapiro142/nexuscli: 表形式データのローカルファーストエージェント層。スプレッドシートをドロップします。AI エージェントがクエリできるローカル MCP サーバーを取得します。 · GitHub
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
スコットシャピロ142
/
ネクサスクリ
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE app app cli cli コンポーネント/ render コンポーネント/ render docs docs lib lib public public scripts scripts testing testing .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.cli.json tsconfig.cli.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
表形式データのローカルファーストエージェント層。
CSV、XLSX、または SQLite ファイルをドロップします。データがマシンから流出することなく、AI エージェントがクエリや操作できるローカル MCP サーバーを取得します。
Obsidian はローカルファーストのメモを提供してくれました。 Nexus は構造化データに対しても同じことを行います。
ビジョンポスト · 問題 · MITライセンス取得
現在、データを AI ツールに渡すと、データはクラウドに送られます。 Salesforce + ChatGPT、Sheets + Gemini、Notion + Claude — 同じパターン。
Nexusはそれを逆転させます。データはマシン上に残ります。 Claude Code、Cursor、およびその他の MCP 対応エージェントは、シートを意味的に意味のあるツール ( description_source 、 find_rows 、 create_collection など) として公開するローカル サーバーと通信します。マスター シートは決して変更されません。コレクション、ブランチ (仮定のセル オーバーレイ)、ビュー、スナップショット、注釈はすべて、非破壊的に最上位に重ねられます。
ユニバーサル入力。 CSV、XLSX、SQLite、Google スプレッドシート - 1 つの CLI、任意の表形式のソース。
エージェントネイティブ。すべてのシートが MCP サーバーになります。 Claude Code または Cursor は、これをドメイン固有のツール パレットとして認識します。
非破壊的な派生。サブセットを保存し、

what-if シナリオ、フィルター、特定時点のスナップショット、行の注釈など、すべてマスターに触れることなく実行できます。
選択的なクラウド公開。 v0.3.x で登場 — マスターはローカルにありながら、特定の派生をクラウドに共有します。
インストールと最初の実行 (60 秒以内)
# 1. 任意のローカル シートで Nexus をポイントする
npx @pixeldesigns/nexus connect ~ /Downloads/customers.csv
# 2. MCP サーバーを起動します (デフォルトでは localhost:5391/mcp 上の HTTP)
npx @pixeldesigns/nexusserve
# 3. 別のターミナルでクロード コードを接続します。
クロード mcp add --transport http nexus http://localhost:5391/mcp
クロード
> このシートには何が含まれていますか?
> 古い顧客を見つけてアウトリーチメールの下書きを作成する
> 古くなった顧客を「ニーズアウトリーチ」というコレクションとして保存します
これがローカル ファイル フロー全体です。 Iris (シートを意味論的に読み取る LLM) が説明を生成し、列が入力され、提案された質問が表示され、エージェントはデータにちなんで名付けられたツール パレットを取得します。
非公開の Google スプレッドシートの初回実行
公開 Google スプレッドシートは、シートが「リンクを知っている人 → 閲覧者」として共有されている場合、認証なしで機能します。プライベート シートには 1 回限りの Google サインインが必要です。
#1. 一度サインインします。最上位のエイリアスは同等です。
Nexus 認証ログイン Google
# または: nexus ログイン Google
# 2. シェルが処理しないように URL を引用します。または #gid を構文として使用します。
ネクサス接続「 https://docs.google.com/spreadsheets/d/<sheet-id>/edit#gid=0 」
# 3. 再接続せずに、キャッシュされた最新の行から保存されたビューをクエリします。
ネクサス クエリ < ビュー名 >
Nexus は引き続きパブリック CSV エクスポートを最初に試行します。 Google がプライベート/ログイン ページで応答し、Google OAuth トークンを持っている場合、Nexus は Sheets API v4 を使用し、それらの行を同じ CSV インジェスト パイプラインに変換し、後の Nexus クエリ実行のために最新のマスター スナップショットをローカルに保存します。
最初の実行後は速くなります
インス

グローバルに背が高いので、コマンドは単に nexus です。
npm install -g @pixeldesigns/nexus
nexus connect ~ /Downloads/customers.csv
ネクサスサーブ
要件
プライベート Google スプレッドシートの場合: 余分なことは何もありません。Nexus には登録済みの Google OAuth クライアントが同梱されているため、nexus auth ログイン Google がそのまま機能します。 (資格情報を BYO したい投稿者は、 NEXUS_GOOGLE_CLIENT_ID / NEXUS_GOOGLE_CLIENT_SECRET を設定できます。)
Nexus は、LLM (Iris) を使用してシート、つまり型付きの列、構造的な概要、推奨されるビュー、および非自明なパターン (「Tell」) を事前に読み取ります。アイリスはオプションです。 3 つのバックエンド、次の順序で自動検出されます。
クロード クロードが PATH 上にある場合はコードを作成します。既存のクロード コード認証を使用します (2 番目のキーはありません)。各 Nexus 接続は、クロードの使用量を少量消費します。
OPENROUTER_API_KEY が設定されている場合は OpenRouter (env または ~/.nexus/config.json )。
ローカル — LLM なし。 nexus connect は、列を取り込み、型を入力し、行を保持します。エージェントは、最初の MCP コンタクト時にシートの独自の説明を作成します。
nexus connect または NEXUS_SAMPLER=... env で --sampler local|claude-code|openrouter を使用して特定のバックエンドを強制します。 Claude Code または OpenRouter によって選択されたモデルを NEXUS_MODEL=... でオーバーライドします。
OpenRouter を使用するには、 openrouter.ai/keys でキーを取得し、次のいずれかを行います。
# オプション A: 一度保存します (~/.nexus/config.json、chmod 600 を書き込みます)
nexus config set-key sk-or-...
# オプション B: シェルごとにエクスポート (保存されたキーよりも常に env が優先されます)
エクスポート OPENROUTER_API_KEY=sk-or-...
nexus config get で設定内容を確認します。 nexus config unset-key を使用して、保存されたキーを削除します。
nexus connect <path-or-url> シート/データベースをマスターソースとして登録します。
サポート: .csv、.tsv、.xlsx、.xls、.sqlite、
および公開/非公開の Google スプレッドシート URL。
--sampler <バックエンド> Iris バックエンド: ローカル |クロードコード |オープンルーター。
省略した場合は自動検出されます。
--skip-iris Iris を実行しないでください

all (--sampler local のエイリアス)。
nexus list リストの派生 (ビュー、コレクション、
ブランチ、スナップショット、アノテーション)
アクティブなソース。
nexus list --sources 接続されているすべてのマスター ソースをリストします。
nexus query <view-name> 保存されたビューを実行し、行を出力します。
nexus tools Iris が公開する MCP ツール定義を印刷します。
アクティブなソースの場合。
nexusserve MCP サーバーを起動します。
--port <n> HTTP ポート (デフォルトは 5391)
--host <h> バインドアドレス (デフォルトは 127.0.0.1)
--stdio stdio 経由で提供します (「claude mcp add」 stdio モードの場合)
nexus config get 解決された設定を表示します (シークレットはマスクされています)。
nexus config set-key [<key>] OpenRouter API キーを保存します (または標準入力経由でパイプします)。
nexus config unset-key 保存されているキーを削除します。
nexus config path 設定ファイルのパスを出力します。
nexus auth login google プライベート スプレッドシートにアクセスするには、Google にログインします。
別名: nexus ログイン google
--force 再同意/リフレッシュトークンのローテーションを強制します。
nexus auth logout google 保存されている Google OAuth トークンを削除します。
別名: nexus ログアウト google
プライベート Google スプレッドシート
公開 Google スプレッドシートは、認証なしの CSV エクスポート パスでも引き続き機能します。プライベート シートの場合は、一度サインインします。
nexus 認証ログイン google # または: nexus ログイン グーグル
ネクサス接続「 https://docs.google.com/spreadsheets/d/<sheet-id>/edit#gid=0 」
Nexus はまず、パブリック CSV エクスポート URL を試行します。 Google がプライベート/ログイン応答を返し、Google OAuth トークンが利用可能な場合、Google Sheets API にフォールバックし、ローカル/パブリック シートで使用されるのと同じ CSV パーサーを通じて返された行をフィードします。保存された Google トークンは、所有者のみのファイル権限を持つ ~/.nexus/auth/google.json の下に存在し、 nexus auth logout google または nexus logout google を使用して削除できます。
シェルは、一致するものが見つからないか、URL が壊れていると言います。Google シートの URL を引用してください。 ?および #gid=0 は、zsh などのシェルで意味を持ちます。
Google はリフレッシュ トークンを返しませんでした: run ne

xus auth login google --force を使用すると、同意画面が強制的に表示され、リフレッシュ トークンがローテーションされます。
セッションが期限切れまたは取り消されました: nexus auth login google を再度実行します。
Google から許可が拒否されました: サインインしたアカウントがシートを表示できることを確認するか、シートを「リンクを知っている人 → 閲覧者」に切り替えてパブリック パスを使用してください。
ネクサス サーブ --ポート 5391
クロード mcp add --transport http nexus http://localhost:5391/mcp
stdio トランスポート (Claude Code が Nexus 自体を起動します):
クロード mcp add nexus -- npx @pixeldesigns/nexusserve --stdio
追加すると、Claude Code 内の /mcp には、保存した派生を反映する自動生成ツール ( query_<your-view> 、 read_<your-collection> ) を含む Nexus のツールが表示されます。
Nexus と同様のツールとの比較
ネクサス
データセット
DuckDB UI
二次関数
リル
OpenAI コードインタープリター
Excel 用コパイロット
完全にマシン上で動作します
✅
✅
✅
❌
✅
❌
❌
CSV / XLSX / SQLite / シートを読み取ります
✅
⚠️ SQLite ファースト
✅
⚠️アプリ内のみ
⚠️ 寄木細工第一
⚠️アップロード
❌ Excelのみ
データを AI エージェント (MCP) に公開します
✅
❌
❌
⚠️アプリ内AI
❌
❌
⚠️アプリ内AI
型付きセマンティック レイヤー (生のセルではない)
✅アイリス
❌
❌
❌
✅ メトリクス
❌
⚠️部分的
非破壊的な派生 (ビュー、ブランチ、スナップショット)
✅
❌
❌
❌
⚠️ダッシュボード
❌
❌
オープンソース
✅ マサチューセッツ工科大学
✅ アパッチ
✅ マサチューセッツ工科大学
⚠️AGPL/クラウド
✅ アパッチ
❌
❌
いつどれを選択するか:
データセット — SQLite データベースをブラウズ可能な Web UI として公開するのに最適です。さまざまな視聴者 (データ ジャーナリズム、公開データセット)、エージェントの統合なし。
DuckDB UI — Parquet/CSV を介した高速ローカル分析 SQL に最適です。エージェント層ではなくクエリエンジン。
Quadratic / Copilot / Code Interpreter — アップロードに問題がなく、洗練されたアプリ内 AI エクスペリエンスが必要な場合に最適です。アップロードできない場合に備えて Nexus が存在します

大丈夫。
Rill — ローカルファーストの BI ダッシュボードに最適です。重複する地元第一主義の精神。異なるプリミティブ (ダッシュボードとエージェント ツール)。
Nexus — 既存の AI エージェント (クロード コード、カーソル、任意の MCP クライアント) が、what-if 用の非破壊レイヤーを使用して、アップロードせずにその場でスプレッドシートをクエリする場合に最適です。
セキュリティ モデル — Google OAuth 認証情報
Nexus には、登録済みの Google 「デスクトップ アプリ」OAuth クライアントがバイナリに埋め込まれて出荷されます。クライアント ID とシークレットは、公開されたソースと npm tarball に表示されます。これは意図的なものです:
Google のデスクトップ アプリ クライアント タイプでは、トークン交換ごとに client_secret が必要です。 PKCE のみは実行できません (経験的に検証されています - lib/auth/google/client-creds.ts を参照)。トークン エンドポイントは 400 client_secret が見つかりませんを返します。シークレットを省略した場合。
Google は、デスクトップ クライアント シークレットは「明らかにシークレットとして扱われない」と明示しています。 https://developers.google.com/identity/protocols/oauth2 を参照してください。
同等のすべての OSS CLI には、埋め込みシークレットが付属しています。 gcloud 、 gh 、 firebase 、および npm はすべて、Google OAuth クライアント シークレットをバイナリで配布します。
ゼロ構成。 Nexus と nexus 認証ログイン Google Works をインストールします。
PKCE は引き続き認証コードの傍受から保護します。
リフレッシュ トークン、スコープ、データはすべてマシン上にあります。
これがクライアント ID にとって何を意味するか:
Nexus ユーザーは自分として認証されます

[切り捨てられた]

## Original Extract

Local-first agent layer for tabular data. Drop any spreadsheet — get a local MCP server your AI agent can query. - scottshapiro142/nexuscli

GitHub - scottshapiro142/nexuscli: Local-first agent layer for tabular data. Drop any spreadsheet — get a local MCP server your AI agent can query. · GitHub
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
scottshapiro142
/
nexuscli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE app app cli cli components/ render components/ render docs docs lib lib public public scripts scripts tests tests .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md eslint.config.mjs eslint.config.mjs next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs tsconfig.cli.json tsconfig.cli.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts View all files Repository files navigation
Local-first agent layer for tabular data.
Drop any CSV, XLSX, or SQLite file. Get a local MCP server your AI agent can query and manipulate — without your data ever leaving your machine.
Obsidian gave us local-first notes. Nexus does the same for structured data.
Vision post · Issues · MIT licensed
When you hand your data to AI tools today, it goes to their cloud. Salesforce + ChatGPT, Sheets + Gemini, Notion + Claude — same pattern.
Nexus inverts that. Your data stays on your machine. Claude Code, Cursor, and any other MCP-aware agent talks to a local server that exposes your sheets as semantically meaningful tools ( describe_source , find_rows , create_collection , …). The master sheet is never mutated — collections, branches (what-if cell overlays), views, snapshots, and annotations all layer on top non-destructively.
Universal input. CSV, XLSX, SQLite, Google Sheets — one CLI, any tabular source.
Agent-native. Every sheet becomes an MCP server. Claude Code or Cursor sees it as a domain-specific tool palette.
Non-destructive derivations. Save subsets, what-if scenarios, filters, point-in-time snapshots, and row annotations — all without touching the master.
Selective cloud publishing. Coming in v0.3.x — share specific derivations to the cloud while the master stays local.
Install & first run (under 60 seconds)
# 1. Point Nexus at any local sheet
npx @pixeldesigns/nexus connect ~ /Downloads/customers.csv
# 2. Start the MCP server (HTTP on localhost:5391/mcp by default)
npx @pixeldesigns/nexus serve
# 3. In another terminal, connect Claude Code to it
claude mcp add --transport http nexus http://localhost:5391/mcp
claude
> what does this sheet contain ?
> find stale customers and draft outreach emails
> save the stale customers as a collection called " needs-outreach "
That's the whole local-file flow. Iris (the LLM that reads your sheet semantically) generates a description, columns get typed, suggested questions appear, and your agent gets a tool palette named after your data.
Private Google Sheets first run
Public Google Sheets work without auth when the sheet is shared as “Anyone with the link → Viewer.” Private sheets need a one-time Google sign-in:
# 1. Sign in once. The top-level alias is equivalent.
nexus auth login google
# or: nexus login google
# 2. Quote the URL so shells do not treat ? or #gid as syntax.
nexus connect " https://docs.google.com/spreadsheets/d/<sheet-id>/edit#gid=0 "
# 3. Query saved views from the cached latest rows without reconnecting.
nexus query < view-name >
Nexus still tries the public CSV export first. If Google responds with a private/login page and you have Google OAuth tokens, Nexus uses the Sheets API v4, converts those rows into the same CSV ingestion pipeline, and stores the latest master snapshot locally for later nexus query runs.
Going faster after the first run
Install globally so the command is just nexus :
npm install -g @pixeldesigns/nexus
nexus connect ~ /Downloads/customers.csv
nexus serve
Requirements
For private Google Sheets: nothing extra — Nexus ships with a registered Google OAuth client, so nexus auth login google just works. (Contributors who want to BYO credentials can set NEXUS_GOOGLE_CLIENT_ID / NEXUS_GOOGLE_CLIENT_SECRET .)
Nexus uses an LLM (Iris) to pre-read your sheet — typed columns, structural summary, suggested views, and non-obvious patterns ("Tells"). Iris is optional. Three backends, auto-detected in this order:
Claude Code if claude is on your PATH . Uses your existing Claude Code auth — no second key. Each nexus connect consumes a small amount of your Claude usage.
OpenRouter if OPENROUTER_API_KEY is set (env or ~/.nexus/config.json ).
Local — no LLM. nexus connect ingests, types columns, and persists rows. Your agent forms its own description of the sheet on first MCP contact.
Force a specific backend with --sampler local|claude-code|openrouter on nexus connect , or NEXUS_SAMPLER=... env. Override the model picked by Claude Code or OpenRouter with NEXUS_MODEL=... .
To use OpenRouter, get a key at openrouter.ai/keys , then either:
# Option A: store it once (writes ~/.nexus/config.json, chmod 600)
nexus config set-key sk-or-...
# Option B: export per-shell (env always wins over the stored key)
export OPENROUTER_API_KEY=sk-or-...
Check what's set with nexus config get . Remove the stored key with nexus config unset-key .
nexus connect <path-or-url> Register a sheet/database as a master source.
Supports: .csv, .tsv, .xlsx, .xls, .sqlite,
and public/private Google Sheets URLs.
--sampler <backend> Iris backend: local | claude-code | openrouter.
Auto-detected if omitted.
--skip-iris Don't run Iris at all (alias for --sampler local).
nexus list List derivations (views, collections,
branches, snapshots, annotations) for the
active source.
nexus list --sources List every connected master source.
nexus query <view-name> Run a saved view and print rows.
nexus tools Print the MCP tool definitions Iris exposes
for the active source.
nexus serve Start the MCP server.
--port <n> HTTP port (default 5391)
--host <h> Bind address (default 127.0.0.1)
--stdio Serve over stdio (for `claude mcp add` stdio mode)
nexus config get Show resolved config (secrets masked).
nexus config set-key [<key>] Store OpenRouter API key (or pipe via stdin).
nexus config unset-key Remove the stored key.
nexus config path Print the config file path.
nexus auth login google Sign in to Google for private Sheets access.
Alias: nexus login google
--force Force re-consent / refresh-token rotation.
nexus auth logout google Remove stored Google OAuth tokens.
Alias: nexus logout google
Private Google Sheets
Public Google Sheets still work through the no-auth CSV export path. For private sheets, sign in once:
nexus auth login google # or: nexus login google
nexus connect " https://docs.google.com/spreadsheets/d/<sheet-id>/edit#gid=0 "
Nexus first tries the public CSV export URL. If Google returns a private/login response and Google OAuth tokens are available, it falls back to the Google Sheets API and then feeds the returned rows through the same CSV parser used by local/public sheets. Stored Google tokens live under ~/.nexus/auth/google.json with owner-only file permissions and can be removed with nexus auth logout google or nexus logout google .
Shell says no matches found or mangles the URL: quote Google Sheet URLs. ? and #gid=0 have meaning in shells like zsh.
Google did not return a refresh token: run nexus auth login google --force to force the consent screen and rotate the refresh token.
Session expired or revoked: run nexus auth login google again.
Permission denied from Google: make sure the signed-in account can view the sheet, or switch the sheet to “Anyone with the link → Viewer” and use the public path.
nexus serve --port 5391
claude mcp add --transport http nexus http://localhost:5391/mcp
stdio transport (Claude Code launches Nexus itself):
claude mcp add nexus -- npx @pixeldesigns/nexus serve --stdio
Once added, /mcp inside Claude Code shows Nexus's tools, including auto-generated ones ( query_<your-view> , read_<your-collection> ) that reflect the derivations you've saved.
How Nexus compares to similar tools
Nexus
Datasette
DuckDB UI
Quadratic
Rill
OpenAI Code Interpreter
Copilot for Excel
Runs entirely on your machine
✅
✅
✅
❌
✅
❌
❌
Reads CSV / XLSX / SQLite / Sheets
✅
⚠️ SQLite-first
✅
⚠️ in-app only
⚠️ Parquet-first
⚠️ upload
❌ Excel only
Exposes data to your AI agent (MCP)
✅
❌
❌
⚠️ in-app AI
❌
❌
⚠️ in-app AI
Typed semantic layer (not raw cells)
✅ Iris
❌
❌
❌
✅ metrics
❌
⚠️ partial
Non-destructive derivations (views, branches, snapshots)
✅
❌
❌
❌
⚠️ dashboards
❌
❌
Open source
✅ MIT
✅ Apache
✅ MIT
⚠️ AGPL/cloud
✅ Apache
❌
❌
When to pick which:
Datasette — best for publishing a SQLite database as a browsable web UI. Different audience (data journalism, public datasets), no agent integration.
DuckDB UI — best for fast local analytical SQL over Parquet/CSV. Querying engine, not agent layer.
Quadratic / Copilot / Code Interpreter — best when uploading is fine and you want a polished in-app AI experience. Nexus exists for the case when uploading is not fine.
Rill — best for local-first BI dashboards. Overlapping local-first ethos; different primitive (dashboards vs. agent tools).
Nexus — best when you want your existing AI agent (Claude Code, Cursor, any MCP client) to query your spreadsheets in place , without uploading, with a non-destructive layer for what-ifs.
Security model — Google OAuth credentials
Nexus ships with a registered Google "Desktop app" OAuth client embedded in the binary. The client ID and secret are visible in the published source and npm tarball. This is deliberate:
Google's Desktop app client type requires client_secret on every token exchange. PKCE-only is not viable (empirically verified — see lib/auth/google/client-creds.ts ). The token endpoint returns 400 client_secret is missing. when the secret is omitted.
Google explicitly states the Desktop client secret "is obviously not treated as a secret" — see https://developers.google.com/identity/protocols/oauth2 .
Every comparable OSS CLI ships its embedded secret. gcloud , gh , firebase , and npm all distribute Google OAuth client secrets in their binaries.
Zero configuration. Install Nexus and nexus auth login google works.
PKCE still protects against auth-code interception.
Your refresh token, your scope, your data — all on your machine.
What this means for the client identity:
Nexus users authenticate as thems

[truncated]

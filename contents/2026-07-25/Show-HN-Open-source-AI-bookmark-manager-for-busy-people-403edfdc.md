---
source: "https://github.com/rortan134/cache-app"
hn_url: "https://news.ycombinator.com/item?id=49050909"
title: "Show HN: Open-source AI bookmark manager for busy people"
article_title: "GitHub - rortan134/cache-app: (New) The AI bookmark manager for busy people. With support for links, notes, images, and videos from your favorite platforms · GitHub"
author: "gsmt"
captured_at: "2026-07-25T20:06:07Z"
capture_tool: "hn-digest"
hn_id: 49050909
score: 1
comments: 0
posted_at: "2026-07-25T19:50:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source AI bookmark manager for busy people

- HN: [49050909](https://news.ycombinator.com/item?id=49050909)
- Source: [github.com](https://github.com/rortan134/cache-app)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T19:50:59Z

## Translation

タイトル: Show HN: 忙しい人のためのオープンソース AI ブックマーク マネージャー
記事のタイトル: GitHub - rortan134/cache-app: (新) 忙しい人のための AI ブックマーク マネージャー。お気に入りのプラットフォーム · GitHub からのリンク、メモ、画像、ビデオのサポート
説明: (新機能) 忙しい人のための AI ブックマーク マネージャー。お気に入りのプラットフォームからのリンク、メモ、画像、ビデオのサポート - rortan134/cache-app

記事本文:
GitHub - rortan134/cache-app: (新) 忙しい人のための AI ブックマーク マネージャー。お気に入りのプラットフォーム · GitHub からのリンク、メモ、画像、ビデオのサポート
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
ロルタン134
/
キャッシュアプリ
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,093 コミット 1,093 コミット .agents/ スキル .agents/ スキル .github .github .husky .husky .vscode .vscode アプリ アプリコンポーネント コンポーネント env env 拡張機能/ キャッシュアプリ拡張機能/ キャッシュアプリフック フック lib lib prisma prisma public public .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .nvmrc .nvmrc AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md SECURITY.md SECURITY.md biome.json biome.json bun.lock bun.lock get-locale.ts get-locale.ts get-region.ts get-region.ts gt-lock.json gt-lock.json gt.config.json gt.config.json load-translations.ts load-translations.ts next.config.ts next.config.ts package.json package.json postcss.config.json postcss.config.json prisma.config.ts prisma.config.ts proxy.ts proxy.ts tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プラットフォーム間でブックマークを 1 つの実用的なライブラリに統合します。
特徴 •
理由 •
技術スタック •
MCP •
ロードマップ •
貢献する •
ライセンス
ブックマークが壊れています。ツイート、ビデオ、または投稿で「保存」をクリックするときは、これを覚えておく価値があるという意図的な決定を行っていることになります。しかし、その意図はすぐに失われます。実際のワークフローや目標とは何の関係もなく、十数のプラットフォームに散在し、決して再訪することのないリストとなって消えてしまいます。フィードはスクロールし続けるように設計されており、必要な情報を再度表示するのに役立つものではありません。既存のツールは、「保存」アクションを開始点ではなく後付け、つまり行き止まりとして扱います。
その署名のためキャッシュが存在します

ナルは無駄にするには価値がありすぎます。保存という行為を第一級のイベントとして扱い、その意図を行動に移すことを中心としたエクスペリエンス全体を構築します。プラットフォームを置き換えるものではありません。それらを使用する理由の背後にある意図を尊重し、目的地を与えます。
ブックマークを統合する — ブラウザーのブックマーク、保存済み Instagram、TikTok のお気に入り、YouTube Watch Later、X/Twitter ブックマーク、GitHub スター、Pinterest、Google フォト、MCP などのブックマークに対する最高級のサポートをすべて 1 か所にまとめて、Cache を日常生活に統合します。保存に制限がある他のツールとは異なり、キャッシュには制限がありません。
スマート コレクション — AI 支援の関連性ランキングを使用して、エントリをコレクションに自動的に整理します。キャッシュは時間の経過とともにユーザーの好みも学習します。
概要 — 各コレクションの上に 1 行の概要が表示されます。新しいエントリが追加されると、即座に更新されます。さらに詳細を確認したい場合は、「展開」をクリックしてください。
AI 支援検索 — Cache AI エージェントに依頼して、保存されているすべてのコンテンツを検索します。
自動化 — カスタム エージェントを作成して何でも実行します。毎日のダイジェスト、要約、毎週のリマインダーなどを生成します。
メモ取り — ブックマークとともにファーストパーティのメモ取りをサポートします。
コラボレーション — キャッシュを使用していなくても、コレクションのライブ ビューを誰とでも共有できます。
ブラウザ拡張機能 — Web 上のどこからでも保存されたコンテンツをキャプチャして同期します。
エクスポートと統合 — 結果をすでに使用している他のツールにパイプします。
シンプルでメンテナンスの手間がかからない — キャッシュは、シンプルでメンテナンスの手間がかからず、常に移植できるように設計されています。
セルフホスティング (作業中)
キャッシュを自己ホストして、データとデザインを完全に制御できます。デフォルトでは、キャッシュのテレメトリはゼロです。
PostgreSQL 12+ (ローカルまたはリモート)
Google Gemini API キー (AI 機能用)
# リポジトリのクローンを作成します
git clone https://github.

com/rortan134/cache-app.git
CDキャッシュ
# 依存関係をインストールする
バンインストール
# 環境をセットアップする
cp .env.example .env
# データベースの URL と API キーを使用して .env を編集します
# データベースをセットアップする
bun run db-deploy
# 開発サーバーを起動します
バンラン開発
http://localhost:3000 を開きます。
完全なリストについては、環境変数のリファレンスを参照してください。
Next.js · Bun · PostgreSQL · Prisma ORM · Better Auth · Tailwind — など
カテゴリ
テクノロジー
フレームワーク
Next.js (アプリルーター)
UI
React 、ベース UI 、Tailwind CSS
リッチテキスト
字句、ストリームダウン
データベース
PostgreSQL、Redis
認証
より良い認証
検証
ゾッド、@t3-oss/env-nextjs
AI/LLM
AI SDK 、AI ゲートウェイ 、Gemini
エージェント Web 検索
タヴィリー
データの取得
SWR、ニュークス
i18n
GTネクスト
定期購入
ストライプ
永続的な実行
ワークフローSDK
セキュリティ (クラウドのみ)
Arcjet (WAF、レート制限、PII 秘匿化)
糸くず
ウルトラサイト (バイオーム)
日付の処理
Day.js 、クロノノード
キャッシュアプリMCP
キャッシュは MCP サーバーを公開するため、Claude、Cursor などの AI エージェントがライブラリを直接読み書きできるようになります。ブックマークの検索、新しいアイテムの保存、コレクションのリストなどを行います。
エンドポイント: https://www.cachd.app/mcp
llms.txt — エージェントのコンテキストとツールのリファレンス
アプリからベアラー トークンを使用してセットアップ プロンプトを生成します (統合 → MCP)。
リマインドする — アイテムを保存または参照するときに、後で戻ってくるように独自のリマインダーを設定します。
コメント — エントリのスレッドコメントを追加および表示します。
受信箱ビュー — エントリをレビューするための優先順位付けビュー。
Notes の改善 — より豊富な編集エクスペリエンス、高度な書式設定。
Raycast の統合 — Raycast からキャッシュをキャプチャして検索します。
Substack の統合 — Substack の投稿とニュースレターをインポートして保存します。
寄付を歓迎します!詳細については、貢献ガイドをご覧ください。
バグが発生したと思われる場合は、問題をオープンしてください。
このプロジェクトはフォローします

■ 貢献者規約の行動規範。
このプロジェクトは、Apache License 2.0 に基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
(新) 忙しい人のための AI ブックマーク マネージャー。お気に入りのプラットフォームからのリンク、メモ、画像、ビデオのサポート
Readme Apache-2.0 ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

(New) The AI bookmark manager for busy people. With support for links, notes, images, and videos from your favorite platforms - rortan134/cache-app

GitHub - rortan134/cache-app: (New) The AI bookmark manager for busy people. With support for links, notes, images, and videos from your favorite platforms · GitHub
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
rortan134
/
cache-app
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,093 Commits 1,093 Commits .agents/ skills .agents/ skills .github .github .husky .husky .vscode .vscode app app components components env env extensions/ cache-app extensions/ cache-app hooks hooks lib lib prisma prisma public public .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .nvmrc .nvmrc AGENTS.md AGENTS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md biome.json biome.json bun.lock bun.lock get-locale.ts get-locale.ts get-region.ts get-region.ts gt-lock.json gt-lock.json gt.config.json gt.config.json load-translations.ts load-translations.ts next.config.ts next.config.ts package.json package.json postcss.config.json postcss.config.json prisma.config.ts prisma.config.ts proxy.ts proxy.ts tsconfig.json tsconfig.json View all files Repository files navigation
Unify your bookmarks across platforms into a single actionable library.
Features •
Why •
Tech Stack •
MCP •
Roadmap •
Contributing •
License
Bookmarking is broken. When you hit "save" on a tweet, a video, or a post, you are making a deliberate decision that this is worth remembering . But that intent is immediately lost. It vanishes into a list you never revisit, scattered across a dozen platforms with no connection to your actual workflow or goals. The feeds are designed to keep you scrolling, not to help you resurface what you need. Existing tools treat the "save" action as an afterthought, a dead end rather than a starting point.
Cache exists because that signal is too valuable to waste. It treats the act of saving as a first-class event and builds the entire experience around turning that intent into action. It does not replace your platforms; it respects the intent behind why you use them and gives it a destination.
Unify your bookmarks — Integrate Cache into your day-to-day with first-class support for bookmarks from Browser bookmarks, Instagram Saved, TikTok Favorites, YouTube Watch Later, X/Twitter bookmarks, GitHub Stars, Pinterest, Google Photos, MCP, and more, all in one place. Unlike other tools that cap saves, Cache has no limits.
Smart collections — Automatically organizes entries into your collections with AI-assisted relevance ranking. Cache even learns your preferences over time.
Overviews — See a 1-line summary above every collection. As new entries are added, it updates instantly. And if you want to see more detail, just hit expand.
AI-assisted search — Ask the Cache AI agent and search across all your saved content.
Automations — Create custom agents to do anything. Generate daily digests, summaries, weekly reminders, and much more.
Note-taking — First-party note-taking support alongside bookmarks.
Collaboration — Share a live view of any collection with anyone, even if they don't use Cache.
Browser extension — Capture and sync saved content from anywhere on the web.
Export & integrate — Pipe results into other tools you already use.
Simple and low maintenance — Cache is designed to be simple, low-maintenance, and always portable.
Self-hosting (Work in progress)
You can self-host Cache for total control over your data and design. Cache has zero telemetry by default.
PostgreSQL 12+ (local or remote)
A Google Gemini API key (for AI features)
# Clone the repository
git clone https://github.com/rortan134/cache-app.git
cd cache
# Install dependencies
bun install
# Set up environment
cp .env.example .env
# Edit .env with your database URL and API keys
# Set up the database
bun run db-deploy
# Start the development server
bun run dev
Open http://localhost:3000 .
See the environment variables reference for the full list.
Next.js · Bun · PostgreSQL · Prisma ORM · Better Auth · Tailwind — and more
Category
Technology
Framework
Next.js (App Router)
UI
React , Base UI , Tailwind CSS
Rich Text
Lexical , Streamdown
Database
PostgreSQL, Redis
Auth
Better Auth
Validation
Zod , @t3-oss/env-nextjs
AI/LLM
AI SDK , AI Gateway , Gemini
Agentic Web Search
Tavily
Data Fetching
SWR , nuqs
i18n
gt-next
Subscriptions
Stripe
Durable Execution
Workflow SDK
Security (Cloud-only)
Arcjet (WAF, rate limiting, PII redaction)
Linting
Ultracite (Biome)
Date Handling
Day.js , chrono-node
Cache App MCP
Cache exposes an MCP server so AI agents like Claude, Cursor, and others can read and write your library directly. Search bookmarks, save new items, list collections, and more.
Endpoint: https://www.cachd.app/mcp
llms.txt — agent context and tool reference
Generate a setup prompt with your Bearer token from the app (Integrations → MCP)
Remind me — Set up unique reminders when saving or browsing on items to come back to later.
Comments — Add and view threaded comments on entries.
Inbox view — Triage view for reviewing entries.
Notes improvements — Richer editing experience, advanced formatting.
Raycast integration — Capture and search Cache from Raycast.
Substack integration — Import and save Substack posts and newsletters.
We welcome contributions! Please see our Contributing Guide for details.
Open an issue if you believe you've encountered a bug.
This project follows the Contributor Covenant code of conduct.
This project is licensed under the Apache License 2.0 - see the LICENSE file for details.
(New) The AI bookmark manager for busy people. With support for links, notes, images, and videos from your favorite platforms
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

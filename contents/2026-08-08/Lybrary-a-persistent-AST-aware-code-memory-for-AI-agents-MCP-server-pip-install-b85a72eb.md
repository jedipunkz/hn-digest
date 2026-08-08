---
source: "https://github.com/vibhu-dixit/lybrary"
hn_url: "https://news.ycombinator.com/item?id=49219016"
title: "Lybrary,a persistent AST-aware code memory for AI agents(MCP server pip install)"
article_title: "GitHub - vibhu-dixit/lybrary: Living structure-aware code memory for AI coding agents · GitHub"
author: "vibhudixit_30"
captured_at: "2026-08-08T05:35:22Z"
capture_tool: "hn-digest"
hn_id: 49219016
score: 1
comments: 0
posted_at: "2026-08-08T05:01:16Z"
tags:
  - hacker-news
  - translated
---

# Lybrary,a persistent AST-aware code memory for AI agents(MCP server pip install)

- HN: [49219016](https://news.ycombinator.com/item?id=49219016)
- Source: [github.com](https://github.com/vibhu-dixit/lybrary)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T05:01:16Z

## Translation

タイトル: AI エージェント用の永続的な AST 対応コード メモリ、Lybrary (MCP サーバー pip インストール)
記事タイトル: GitHub - vibhu-dixit/library: AI コーディング エージェントのための生きた構造を認識したコード メモリ · GitHub
説明: AI コーディング エージェント用の生きた構造を認識したコード メモリ - vibhu-dixit/library

記事本文:
GitHub - vibhu-dixit/library: AI コーディング エージェント用の生きた構造を認識したコード メモリ · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ヴィブディクシット
/
図書館
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows lybrary lybrary testing testing .gitignore .gitignore CHANGELOG.md CHANGELOG.md README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ

ファイルナビゲーション
AI コーディング エージェント向けの生きた構造を認識したコード メモリ。
pip インストールライブラリ
問題
AI エージェントはセッションを開始するたびに、コードベースを最初から再読み込みします。
エージェント: 認証ロジックを grep させてください...
エージェント: src/auth/middleware.py を読んでいます...
エージェント: src/auth/jwt.py を読み取り中...
エージェント: src/auth/session.py を読んでいます...
エージェント: src/utils/crypto.py を読み取り中...
↳ 1 行を書き込む前に 4,000 個のトークンが消費されました。
大規模なコードベースでは、これはセッションごとに数十回発生します。トークンが無駄になりました。コンテキストが埋められました。同じファイルが何度も読み取られます。
ライブラリは、ファイルを読み取る代わりにクエリできる永続メモリをエージェントに提供します。
エージェント:memory_query("認証フロー")
↳ 3 個のチャンクが返されました。トークン180個。終わり。
実際の AST 境界を使用してリポジトリのインデックスを作成し、インデックスを自動的に最新の状態に保ち、AI IDE がネイティブに接続する MCP サーバーとして公開します。
🌳 AST を意識したチャンキング
ツリーシッターはコードを解析します - 関数を半分に分割することはありません
⚡ バックグラウンドデーモン
ファイルの変更を監視し、変更されたもののみを再インデックスします
🔍 セマンティック検索
トークンバジェットパッキングによるベクトル検索
🔌 MCPサーバー
Kiro、Cursor、Claude Desktop、Windsurf ですぐに使用可能
📦 完全にローカル
クラウドなし、API キーなし、埋め込みはマシン上で実行されます
🐍 純粋な pip インストール
Python 3.11 ～ 3.14、PyTorch なし、コンパイル不要
サポートされている言語: Python · JavaScript · TypeScript · TSX · Go · Rust · Java · C · C++
pip インストールライブラリ
cd /path/to/your/repo
ライブラリの初期化
lybrary start # インデックスを構築 + バックグラウンド デーモンを開始
ライブラリクエリ「認証フロー」
lybrary start の後、ターミナルを閉じた後でもデーモンは実行を続けます。ファイルの変更は自動的に取得され、影響を受けるチャンクのみが再インデックスされます。
MCP 構成に追加 (Kiro、Cursor、Claude Desktop、Winds)

ウルフ):
{
"mcpサーバー": {
"図書館" : {
"コマンド" : " ライブラリ " ,
"args" : [ "mcp " ]
}
}
}
エージェントには次の 3 つのツールが追加されました。
エージェントはファイルを読み取る前に、memory_query を呼び出す必要があります。
これにより、複数ファイルの読み取りが単一の対象クエリに置き換えられ、大規模なコードベースでのトークン使用量が 80 ～ 90% 削減されます。
コマンド
説明
ライブラリの初期化
.lybrary/ とデフォルト設定を作成する
図書館のスタート
インデックス (必要な場合) + 永続デーモンの開始
図書館停止
デーモンを停止する
図書館のステータス
実行状態、チャンク数、追跡されたファイルを表示
図書館の索引
強制 (再) インデックス
図書館のクエリ
メモリ上のセマンティック検索
図書館のログ
デーモンログの表示/フォロー
図書館MCP
MCP サーバーを開始します (stdio トランスポート)
🏗️ チャンクの仕組み
あなたのファイル
│
▼
ツリーシッターパーサー
│
▼
AST 定義ノード ← 関数、クラス、メソッド、インターフェイス
│
§── クラス Foo ──────► チャンク: クラス本体全体
│ §── def bar ──────► chunk: メソッド bar (独自のチャンクも)
│ └── def baz ──────► チャンク: メソッド baz (独自のチャンクも)
│
└── モジュールレベル --------------► チャンク: インポート、定数、トップレベルのステートメント
各チャンクはコンテキスト ヘッダーを取得し、ONNX ランタイム経由で MiniLM-L6-v2 に埋め込まれます。高速、ローカル、GPU は不要です。
.library/
§── config.toml # モデル、チャンクサイズ、パターンを無視
§──index.db # SQLite: チャンク + float32 ベクトル BLOB
§── file_hashes.json # 増分更新用のコンテンツ ハッシュ マップ
§── デーモン.pid
└── daemon.log
インデクサー — ツリーシッター → AST チャンク → fastembed / ONNX ランタイム埋め込み
ストア — SQLite + numpy (バッチ処理されたドット積によるコサイン類似度、外部ベクトル DB なし)
デーモン — ウォッチドッグ ファイル ウォッチャー + デバウンス + 増分再チャンク/埋め込み
MCP — st 上の FastMCP サーバー

ディオ
AST チャンカー (多言語、cAST スタイル)
コンテンツハッシュによる増分インデックス作成
バックグラウンドデーモン + ファイルウォッチャー (Windows + Unix)
CLI (初期化/開始/停止/ステータス/インデックス/クエリ/ログ/mcp)
MCPサーバー（memory_query、memory_status、memory_update）
階層的なファイル/パッケージの概要
systemd / launchd ユーザーサービスヘルパー
問題やPRを歓迎します。以下を使用してテスト スイートを実行します。
pip install -e " .[dev] "
pytest
📄ライセンス
AI コーディング エージェント向けの生体構造を認識したコード メモリ
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Living structure-aware code memory for AI coding agents - vibhu-dixit/lybrary

GitHub - vibhu-dixit/lybrary: Living structure-aware code memory for AI coding agents · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
vibhu-dixit
/
lybrary
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows lybrary lybrary tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Living structure-aware code memory for AI coding agents.
pip install lybrary
The problem
Every time an AI agent starts a session, it re-reads your codebase from scratch.
agent: let me grep for auth logic...
agent: reading src/auth/middleware.py...
agent: reading src/auth/jwt.py...
agent: reading src/auth/session.py...
agent: reading src/utils/crypto.py...
↳ 4,000 tokens burned before writing a single line.
On a large codebase this happens dozens of times per session. Tokens wasted. Context filled. Same files read over and over.
lybrary gives your agent a persistent memory it can query instead of reading files.
agent: memory_query("authentication flow")
↳ 3 chunks returned. 180 tokens. Done.
It indexes your repo using real AST boundaries, keeps the index fresh automatically, and exposes it as an MCP server that any AI IDE connects to natively.
🌳 AST-aware chunking
tree-sitter parses your code — never splits a function in half
⚡ Background daemon
watches for file changes, re-indexes only what changed
🔍 Semantic search
vector search with token-budget packing
🔌 MCP server
works with Kiro, Cursor, Claude Desktop, Windsurf out of the box
📦 Fully local
no cloud, no API keys, embeddings run on your machine
🐍 Pure pip install
Python 3.11–3.14, no PyTorch, no compilation needed
Supported languages: Python · JavaScript · TypeScript · TSX · Go · Rust · Java · C · C++
pip install lybrary
cd /path/to/your/repo
lybrary init
lybrary start # builds index + starts background daemon
lybrary query " authentication flow "
After lybrary start , the daemon keeps running even after you close the terminal. File changes are picked up automatically — only affected chunks are re-indexed.
Add to your MCP config (Kiro, Cursor, Claude Desktop, Windsurf):
{
"mcpServers" : {
"lybrary" : {
"command" : " lybrary " ,
"args" : [ " mcp " ]
}
}
}
Your agent now has three tools:
Agents should call memory_query before reading any files.
This replaces multi-file reads with a single targeted query — cutting token usage by 80–90% on large codebases.
Command
Description
lybrary init
Create .lybrary/ and default config
lybrary start
Index (if needed) + start persistent daemon
lybrary stop
Stop the daemon
lybrary status
Show running state, chunk count, tracked files
lybrary index
Force (re)index
lybrary query
Semantic search over the memory
lybrary logs
View / follow daemon log
lybrary mcp
Start MCP server (stdio transport)
🏗️ How chunking works
your file
│
▼
tree-sitter parser
│
▼
AST definition nodes ← functions, classes, methods, interfaces
│
├── class Foo ──────────► chunk: entire class body
│ ├── def bar ───────► chunk: method bar (its own chunk too)
│ └── def baz ───────► chunk: method baz (its own chunk too)
│
└── module-level ────────► chunk: imports, constants, top-level statements
Each chunk gets a context header and is embedded with MiniLM-L6-v2 via ONNX Runtime — fast, local, no GPU needed.
.lybrary/
├── config.toml # model, chunk size, ignore patterns
├── index.db # SQLite: chunks + float32 vector blobs
├── file_hashes.json # content-hash map for incremental updates
├── daemon.pid
└── daemon.log
Indexer — tree-sitter → AST chunks → fastembed / ONNX Runtime embeddings
Store — SQLite + numpy (cosine similarity via batched dot product, no external vector DB)
Daemon — watchdog file watcher + debounce + incremental re-chunk/embed
MCP — FastMCP server over stdio
AST chunker (multi-language, cAST-style)
Incremental indexing via content hashes
Background daemon + file watcher (Windows + Unix)
CLI ( init / start / stop / status / index / query / logs / mcp )
MCP server ( memory_query , memory_status , memory_update )
Hierarchical file/package summaries
systemd / launchd user service helper
Issues and PRs welcome. Run the test suite with:
pip install -e " .[dev] "
pytest
📄 License
Living structure-aware code memory for AI coding agents
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

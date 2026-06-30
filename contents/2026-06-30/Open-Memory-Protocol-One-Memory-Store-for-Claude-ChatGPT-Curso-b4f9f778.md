---
source: "https://github.com/SMJAI/open-memory-protocol"
hn_url: "https://news.ycombinator.com/item?id=48726966"
title: "Open Memory Protocol – One Memory Store for Claude, ChatGPT, Curso"
article_title: "GitHub - SMJAI/open-memory-protocol: An open standard for portable, interoperable AI memory across tools, sessions, and devices. · GitHub"
author: "soji_mathew"
captured_at: "2026-06-30T00:30:33Z"
capture_tool: "hn-digest"
hn_id: 48726966
score: 2
comments: 0
posted_at: "2026-06-30T00:05:33Z"
tags:
  - hacker-news
  - translated
---

# Open Memory Protocol – One Memory Store for Claude, ChatGPT, Curso

- HN: [48726966](https://news.ycombinator.com/item?id=48726966)
- Source: [github.com](https://github.com/SMJAI/open-memory-protocol)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T00:05:33Z

## Translation

タイトル: オープン メモリ プロトコル – クロード、ChatGPT、Curso 用の 1 つのメモリ ストア
記事のタイトル: GitHub - SMJAI/open-memory-protocol: ツール、セッション、デバイス間でポータブルで相互運用可能な AI メモリのオープン スタンダード。 · GitHub
説明: ツール、セッション、デバイス間でポータブルで相互運用可能な AI メモリのオープン スタンダード。 - SMJAI/オープンメモリプロトコル

記事本文:
GitHub - SMJAI/open-memory-protocol: ツール、セッション、デバイス間でポータブルで相互運用可能な AI メモリのオープン スタンダード。 · GitHub
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
SMJAI
/
オープンメモリプロトコル
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows アダプター/ claude-mcp アダプター/ claude-mcp パッケージ/ サーバー パッケージ/ サーバー仕様/ v1 仕様/ v1 .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SPEC.md SPEC.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ツール、セッション、デバイス間でポータブルで相互運用可能な AI メモリのオープン スタンダード。
各 AI ツールは異なる方法でユーザーを記憶しますが、それはそのツールの壁内でのみです。
クロードはあなたが昨日言ったことを知っています。カーソルはそうではありません。
ChatGPT はあなたの好みを学習しました。あなたのカスタムエージェントはそうではありません。
Copilot はコード スタイルを認識しました。あなたの端末AIはゼロからスタートします。
ツールを切り替えるたびに、AI はあなたのことを忘れてしまいます。あなたは繰り返します。コンテキストが失われます。ようやくあなたを認識し始めたAIは、見知らぬ人にリセットされます。
これが AI メモリ サイロ問題です。そして、これまでのすべてのサイロ問題と同じ解決策、つまりオープン プロトコルを持っています。
Open Memory Protocol は、AI ツールがユーザーとそのコンテキストに関するメモリを保存、取得、共有する方法に関するベンダー中立の仕様です。
仕様 — メモリ オブジェクト、ストレージ形式、HTTP API の正確な定義
リファレンス サーバー — 自己ホスト可能なオープンソースで、1 つのコマンドで Docker 内で実行されます。
SDK のセット — OMP 互換ツールを構築するための TypeScript および Python ライブラリ
アダプターのセット - Claude (MCP)、OpenAI、Cursor などのプラグイン
OMP を実装する AI ツールは、他の OMP 互換ツールと即座にメモリを共有できます。
要件: Node.js 22 以降
npx omp-server
または Docker を使用して:
docker run -p 3456:3456 -v omp-data:/data ghcr.io/smjai/omp-server
あなたの人生

サーバーは現在 http://localhost:3456 で実行されています。テストしてみましょう:
カール http://localhost:3456/v1/health
# {"ステータス":"ok","バージョン":"0.1","コンプライアンス":"OMP-Core","memories_count":0}
2.クロードを接続します（MCP経由）
Claude デスクトップ構成 ( ~/.claude/claude_desktop_config.json ) に追加します。
{
"mcpサーバー": {
"オンプ" : {
"コマンド" : " npx " ,
"args" : [ "omp-mcp " ],
"環境" : {
"OMP_SERVER" : " http://localhost:3456 " 、
"OMP_API_KEY" : " あなたの OMP キー "
}
}
}
}
AI を活用したメモリの抽出と圧縮を有効にするには、サーバー上で次の設定も行います。
OMP_AI_PROVIDER=人類 # または "openai"
OMP_AI_API_KEY=sk-ant-... # Anthropic または OpenAI キー
あらゆるツールからメモリを書き込む
curl -X POST http://localhost:3456/v1/memories \
-H " Content-Type: application/json " \
-d ' {
"content": "ユーザーは JavaScript よりも TypeScript を好み、冗長なコメントを嫌います",
"タイプ": "セマンティック",
"ソース": { "ツール": "クロード" },
"タグ": ["設定"、"コーディング"]
} '
他のツールからのクエリ
カール「 http://localhost:3456/v1/memories/search?q=coding+preferences 」
仕組み
┌─────────┐ ┌─────────┐ ┌─────────┐
│ クロード │ │ カーソル │ │ あなたの代理人 │
│ (MCP) │ │ (SDK) │ │ (REST API) │
━━━━┬━━━┘ ━━━━┬──────┘ ━━━━┬──────┘
│ │ │
━━━━━━━━━━━━━━━━━━━┘
│
┌───▼───┐
│ OMP サーバー │
│ (自己ホスト型) │
│ │
│ ┌───────┐ │
│ │ SQLite │ │
│ │ / Pgvec │ │
│ └───────┘ │

━━━━━━━━┘
すべてのツールは、制御する単一の OMP サーバーに対して読み取りと書き込みを行います。メモリーストアが 1 つ。すべてのツール。サイロゼロ。
メモリ オブジェクト — メモリの正規スキーマ (コンテンツ、タイプ、ソース、タグ、タイムスタンプ、オプションの埋め込み)
記憶の種類 — エピソード的 (出来事)、意味的 (事実/好み)、手続き的 (ハウツー知識)
REST API — 標準 CRUD + セマンティック検索エンドポイント
認証 — ベアラートークン、ツールごとの API キー
エクスポート/インポート — サーバー間でメモリを移動するためのポータブル JSON 形式
完全な仕様を読む: SPEC.md
{
"id" : " mem_01j9xk2p3q4r5s6t " ,
"content" : " ユーザーはフィンテックのスタートアップを構築しており、クリーンなアーキテクチャを好み、オーバーエンジニアリングを嫌います " ,
"タイプ" : " セマンティック " 、
「ソース」: {
"ツール" : " クロード " 、
"session_id" : " sess_abc123 " ,
"タイムスタンプ" : " 2026-06-29T12:00:00Z "
}、
"タグ" : [ " プロフィール " 、 " 設定 " 、 " エンジニアリング " ]、
"created_at" : " 2026-06-29T12:00:00Z " ,
"updated_at" : " 2026-06-29T12:00:00Z " ,
"expires_at" : null
}
アダプター
ツール
ステータス
インストール
クロード (MCP)
✅ 利用可能
npx omp-mcp
OpenAI アシスタント
🙋 助けを求めています
未解決の問題
カーソル
🙋 助けを求めています
未解決の問題
コパイロット / VS コード
🙋 助けを求めています
未解決の問題
ジェミニ
🙋 助けを求めています
未解決の問題
カスタム (REST)
✅ 利用可能
任意の HTTP クライアント
構築してみませんか?アダプターは通常 100 ～ 200 行です。CONTRIBUTING.md を読み取り、adapters/claude-mcp をテンプレートとして使用します。
OMP API はプレーンな REST であり、あらゆる HTTP クライアントがそのまま動作します。型指定された SDK はロードマップにあります。
構築してみませんか? Python、Go、Rust、Ruby SDK はすべて必要です。 COTRIBUTING.md を参照してください。
# 思い出を保存する
curl -X POST http://localhost:3456/v1/memories \
-H " Content-Type: application/json " \
-d ' {"content":"ユーザーは TypeScript を好みます","type":"semantic","source":

{"ツール":"myapp","タイムスタンプ":"2026-06-30T00:00:00Z"}} '
# 思い出を探す
curl -X POST http://localhost:3456/v1/memories/search \
-H " Content-Type: application/json " \
-d ' {"q":"TypeScript","limit":5} '
なぜオープンソースなのか?
あなたの思い出はあなたのものです。これらを企業のデータベース内にロックしたり、同意なしにモデルのトレーニングに使用したり、ツールを切り替えるときに失ったりしてはなりません。
OMP is designed on these principles:
まずはセルフホスト — サーバーを実行し、データを所有する
ベンダー中立 - 標準を管理する企業はありません
設計によるプライバシー — エクスポートしない限り、思い出がサーバーから離れることはありません
ポータブル — 1 つのコマンドで全メモリをインポート/エクスポート
v0.1 — コア仕様、リファレンス サーバー、MCP アダプター
v0.2 — AI メモリ抽出、会話圧縮、MCP リソース + プロンプト
v0.3 — 埋め込みによるセマンティック検索、pgvector サポート
v0.4 — メモリの名前空間 (プロジェクトごとのメモリ)
v0.5 — Multi-user support, access control
v1.0 — 安定した仕様、オープン標準化団体に提出
OMP はコミュニティ主導型です。必要なものは次のとおりです。
アダプター ビルダー — お気に入りの AI ツールを接続します
SDK contributors — Go, Rust, Java SDKs welcome
Spec reviewers — read SPEC.md and open issues
Early adopters — try it and report what breaks
See CONTRIBUTING.md to get started.
GitHub ディスカッション — 質問、アイデア、フィードバック
Issues — bugs and spec clarifications
Apache 2.0 — 自由に使用、変更、配布できます。 「ライセンス」を参照してください。
Built by SMJAI and contributors.
ツール、セッション、デバイス間でポータブルで相互運用可能な AI メモリのオープン スタンダード。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An open standard for portable, interoperable AI memory across tools, sessions, and devices. - SMJAI/open-memory-protocol

GitHub - SMJAI/open-memory-protocol: An open standard for portable, interoperable AI memory across tools, sessions, and devices. · GitHub
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
SMJAI
/
open-memory-protocol
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows adapters/ claude-mcp adapters/ claude-mcp packages/ server packages/ server spec/ v1 spec/ v1 .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
An open standard for portable, interoperable AI memory across tools, sessions, and devices.
Every AI tool remembers you differently — and only within its own walls.
Claude knows what you told it yesterday. Cursor doesn't.
ChatGPT learned your preferences. Your custom agent hasn't.
Copilot saw your code style. Your terminal AI is starting from zero.
Every time you switch tools, your AI forgets you. You repeat yourself. Context is lost. The AI that was finally starting to know you resets to a stranger.
This is the AI memory silo problem . And it has the same solution as every silo problem before it: an open protocol.
Open Memory Protocol is a vendor-neutral specification for how AI tools store, retrieve, and share memory about users and their context.
A specification — a precise definition of memory objects, storage format, and HTTP API
A reference server — self-hostable, open-source, runs in Docker in one command
A set of SDKs — TypeScript and Python libraries for building OMP-compatible tools
A set of adapters — plug-ins for Claude (MCP), OpenAI, Cursor, and more
Any AI tool that implements OMP can instantly share memory with any other OMP-compatible tool.
Requirements: Node.js 22 or newer
npx omp-server
Or with Docker:
docker run -p 3456:3456 -v omp-data:/data ghcr.io/smjai/omp-server
Your server is now running at http://localhost:3456 . Test it:
curl http://localhost:3456/v1/health
# {"status":"ok","version":"0.1","compliance":"OMP-Core","memories_count":0}
2. Connect Claude (via MCP)
Add to your Claude Desktop config ( ~/.claude/claude_desktop_config.json ):
{
"mcpServers" : {
"omp" : {
"command" : " npx " ,
"args" : [ " omp-mcp " ],
"env" : {
"OMP_SERVER" : " http://localhost:3456 " ,
"OMP_API_KEY" : " your-omp-key "
}
}
}
}
To enable AI-powered memory extraction and compression, also set these on the server:
OMP_AI_PROVIDER=anthropic # or "openai"
OMP_AI_API_KEY=sk-ant-... # your Anthropic or OpenAI key
Write a memory from any tool
curl -X POST http://localhost:3456/v1/memories \
-H " Content-Type: application/json " \
-d ' {
"content": "User prefers TypeScript over JavaScript and dislikes verbose comments",
"type": "semantic",
"source": { "tool": "claude" },
"tags": ["preferences", "coding"]
} '
Query from any other tool
curl " http://localhost:3456/v1/memories/search?q=coding+preferences "
How It Works
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│ Claude │ │ Cursor │ │ Your Agent │
│ (MCP) │ │ (SDK) │ │ (REST API) │
└──────┬──────┘ └──────┬──────┘ └──────┬──────┘
│ │ │
└───────────────────┼───────────────────┘
│
┌────────▼────────┐
│ OMP Server │
│ (self-hosted) │
│ │
│ ┌───────────┐ │
│ │ SQLite │ │
│ │ / Pgvec │ │
│ └───────────┘ │
└─────────────────┘
Every tool reads and writes to a single OMP server you control. One memory store. All tools. Zero silos.
Memory Object — the canonical schema for a memory (content, type, source, tags, timestamps, optional embedding)
Memory Types — episodic (events), semantic (facts/preferences), procedural (how-to knowledge)
REST API — standard CRUD + semantic search endpoints
Authentication — bearer token, per-tool API keys
Export/Import — portable JSON format for moving memories between servers
Read the full specification: SPEC.md
{
"id" : " mem_01j9xk2p3q4r5s6t " ,
"content" : " User is building a fintech startup, prefers clean architecture, dislikes over-engineering " ,
"type" : " semantic " ,
"source" : {
"tool" : " claude " ,
"session_id" : " sess_abc123 " ,
"timestamp" : " 2026-06-29T12:00:00Z "
},
"tags" : [ " profile " , " preferences " , " engineering " ],
"created_at" : " 2026-06-29T12:00:00Z " ,
"updated_at" : " 2026-06-29T12:00:00Z " ,
"expires_at" : null
}
Adapters
Tool
Status
Install
Claude (MCP)
✅ Available
npx omp-mcp
OpenAI Assistants
🙋 Help wanted
Open issue
Cursor
🙋 Help wanted
Open issue
Copilot / VS Code
🙋 Help wanted
Open issue
Gemini
🙋 Help wanted
Open issue
Custom (REST)
✅ Available
Any HTTP client
Want to build one? An adapter is typically 100–200 lines — read CONTRIBUTING.md and use adapters/claude-mcp as a template.
The OMP API is plain REST — any HTTP client works out of the box. Typed SDKs are on the roadmap.
Want to build one? Python, Go, Rust, and Ruby SDKs are all needed. See CONTRIBUTING.md .
# Save a memory
curl -X POST http://localhost:3456/v1/memories \
-H " Content-Type: application/json " \
-d ' {"content":"User prefers TypeScript","type":"semantic","source":{"tool":"myapp","timestamp":"2026-06-30T00:00:00Z"}} '
# Search memories
curl -X POST http://localhost:3456/v1/memories/search \
-H " Content-Type: application/json " \
-d ' {"q":"TypeScript","limit":5} '
Why Open Source?
Your memories are yours. They should not be locked inside a company's database, used to train models without your consent, or lost when you switch tools.
OMP is designed on these principles:
Self-hosted first — you run the server, you own the data
Vendor neutral — no company controls the standard
Privacy by design — memories never leave your server unless you export them
Portable — import/export your full memory in one command
v0.1 — Core spec, reference server, MCP adapter
v0.2 — AI memory extraction, conversation compression, MCP resources + prompts
v0.3 — Semantic search with embeddings, pgvector support
v0.4 — Memory namespacing (per-project memories)
v0.5 — Multi-user support, access control
v1.0 — Stable spec, submitted to open standards body
OMP is community-driven. We need:
Adapter builders — connect your favourite AI tool
SDK contributors — Go, Rust, Java SDKs welcome
Spec reviewers — read SPEC.md and open issues
Early adopters — try it and report what breaks
See CONTRIBUTING.md to get started.
GitHub Discussions — questions, ideas, feedback
Issues — bugs and spec clarifications
Apache 2.0 — free to use, modify, and distribute. See LICENSE .
Built by SMJAI and contributors.
An open standard for portable, interoperable AI memory across tools, sessions, and devices.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/bybit-exchange/kaas"
hn_url: "https://news.ycombinator.com/item?id=49099857"
title: "KaaS – Knowledge as a Service: an out-of-the-box LLM wiki compiler"
article_title: "GitHub - bybit-exchange/kaas: Knowledge as a Service: an out-of-the-box LLM wiki compiler. · GitHub"
author: "lucasmaan"
captured_at: "2026-07-29T17:04:52Z"
capture_tool: "hn-digest"
hn_id: 49099857
score: 1
comments: 0
posted_at: "2026-07-29T16:45:38Z"
tags:
  - hacker-news
  - translated
---

# KaaS – Knowledge as a Service: an out-of-the-box LLM wiki compiler

- HN: [49099857](https://news.ycombinator.com/item?id=49099857)
- Source: [github.com](https://github.com/bybit-exchange/kaas)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T16:45:38Z

## Translation

タイトル: KaaS – Knowledge as a Service: すぐに使える LLM wiki コンパイラー
記事のタイトル: GitHub - bybit-exchange/kaas: Knowledge as a Service: すぐに使える LLM wiki コンパイラー。 · GitHub
説明: Knowledge as a Service: すぐに使える LLM wiki コンパイラー。 - bybit-exchange/kaas

記事本文:
GitHub - bybit-exchange/kaas: Knowledge as a Service: すぐに使える LLM Wiki コンパイラー。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
バイビットエクスチェンジ
/
カース
公共
通知

オン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github .github cmd/ kaas cmd/ kaas docs docs など 内部 内部 py py スクリプト スクリプト Web Web .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md go.mod go.mod go.sum go.sum install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM 主導のナレッジ編集機能を利用して、分散したメモ、ドキュメント、トランスクリプトを検索可能でクエリ可能な個人 Wiki に変換します。
KaaS は社内ツールとして始まりました。私たちの知識はあちこちに散らばっていました
文書、会議、電子メール、そして誰かが役割を変更したり退職したりするたびに、
彼らが築き上げた文脈が彼らとともに消えていったのです。新しい人が数週間を費やした
それをつなぎ合わせます。
蒸留パイプラインがそれを修正しました。一人一人の散りばめられたものを集大成する
彼らのアイデンティティではなく、その役割に関連付けられた資料を Wiki に投稿する - したがって、
誰かが先に進み、生データは消えますが、蒸留された判断は誰にでも残ります
次に席を埋めます。
どちらの方法でも見返りは同じです。組織は同じ質問に再回答するのをやめます。
質問。それが、オープンソースにする価値があると私たちに確信させた理由です。
生のテキストをチャンクして埋め込む一般的な RAG システムとは異なり、KaaS は 4 フェーズ LLM パイプラインを通じてコンテンツをコンパイルします。
生のコンテンツ → 抽出 → 分類 → 書き込み → インデックス → 構造化 Wiki
その結果、ブラックボックスのベクター ストアではなく、人間が判読できる Markdown 記事が作成されます。ナレッジ ベースを読み取り、編集し、Git 管理することができます。
AI エージェントのクイック スタート
すでにコーディング エージェント (Claude Code、Codex、openclaw など) に住んでいますか?スキップドゥ

カー。
これをコピーしてエージェントに貼り付けます。kb-ai がインストールされ、何を行うかを尋ねます。
蒸留し、Wiki を構築し、後のセッションでクエリできるように MCP を接続します。
KaaS をセットアップして、ファイルからクエリ可能なナレッジ ベースを構築します。
https://raw.githubusercontent.com/bybit-exchange/kaas/main/docs/agent-quickstart.md を取得します
そしてそれを正確に従ってください。
Web UI をお好みですか、それとも完全なバックエンドが必要ですか?以下の Docker パスを使用します。
KaaS は、OpenAI 互換 API (OpenAI、DeepSeek、Ollama、vLLM、Azure OpenAI など) を通じて LLM を呼び出します。以下のいずれかの方法を選択してください。
# イメージをビルドする
docker build -t kaas 。
# 実行
docker run -d --name kaas \
-p 8080:8080 \
-v ./data:/app/data \
-e LLM_API_KEY=sk-xxx \
-e LLM_BASE_URL=https://api.openai.com/v1 \
-e LLM_MODEL=gpt-4o-mini \
カース
オプション B: CLI インストール
# インストール (Linux/macOS、amd64/arm64)
カール -fsSL https://raw.githubusercontent.com/bybit-exchange/kaas/main/install.sh |しー
# サービスを開始する
export LLM_API_KEY= " sk-xxx " # OpenAI 互換 API キー
export LLM_BASE_URL= " https://api.openai.com/v1 " # API エンドポイント
export LLM_MODEL= " gpt-4o-mini " # モデル名
kaas サーブ # デフォルト: http://localhost:8080
サポートされているプラットフォーム: Linux/macOS、amd64/arm64。アンインストール: rm -rf ~/.local/share/kaas ~/.local/bin/kaas 。
LLM_BASE_URL のデフォルトは https://api.openai.com/v1 、LLM_MODEL のデフォルトは gpt-4o-mini です。
OpenAI 互換のエンドポイントを指すように変更します。
Claude Code または他の MCP クライアントをナレッジ ベースに接続できるようにするには:
docker run -d --name kaas \
-p 8080:8080 \
-v ./data:/app/data \
-e LLM_API_KEY=sk-xxx \
-e KAAS_MCP_ENABLED=true \
-e KAAS_MCP_TOKEN=あなたのシークレットトークン \
カース
MCP クライアント URL: http://<host>:8080/mcp 、認可: Bearer your-secret-token 。
Go バックエンドは、Python AI エンジンを長時間実行デーモン プロセスとして生成し、mul 経由で通信します。

二重化された stdin/stdout プロトコル。 1 つの Docker イメージにすべてがバンドルされており、サイドカー コンテナーは必要ありません。
4フェーズコンパイルパイプライン : コンセプト/エンティティ/決定事項の抽出 → 記事への分類 → マークダウンの書き込み/マージ → マークダウンインデックスの更新
ワーカー アクセラレーション : 同時抽出/パイプライン ワーカー、サーキット ブレーカー、リースの回復
ストリーミング チャット: Wiki 記事を参照するソース引用を含む SSE ストリーミング
複数の入力ソース: テキストの貼り付け、ファイルのアップロード、または URL の提供
インクリメンタルコンパイル: 新規/変更されたコンテンツのみを再コンパイルします。
Git フレンドリーな出力 : すべての Wiki 記事はプレーンな Markdown です
MCP アクセス : ask ツールを介して、MCP 対応コーディング エージェントからコンパイルされた wiki をクエリします。
コンパイルされた Wiki を任意のモデル コンテキスト プロトコルに公開します
単一の ask ツールを介したクライアント (Claude Code、Codex、openclaw など) —
ask(query, paths?, model?) は、以下に基づいて引用された Markdown 回答を返します。
ウィキ。 2 つのトランスポート:
stdio (ローカル - エージェントがサーバーを生成し、完全に自己完結型):
# エージェントがこれを起動します。 KAAS_KB_DIR をナレッジベースのルートに設定し、
# 環境内の LLM_* 資格情報。
kb-ai mcp # stdio がデフォルトのトランスポートです
# クロードコード:
クロード mcp add kaas -- kb-ai mcp
Codex / openclaw の場合、コマンド kb-ai mcp および env を使用して stdio MCP サーバーを追加します。
KAAS_KB_DIR + LLM_* 。
streamable-http (リモート — バックエンドの:8080 オリジンを通じて公開):
KAAS_MCP_ENABLED=true を指定してコンテナを実行します (「クイック スタート」を参照)。
バックエンドは、 /mcp で MCP エンドポイントを公開します。リモート エージェントをそこに向けます。
# クロードコード:
クロード mcp add --transport http kaas http://host:8080/mcp
KAAS_MCP_TOKEN を、HTTP 上の認証: Bearer <token> を要求するように設定します。
トランスポート (デフォルトではオフ - ローカル/イントラネットを想定)。標準入出力にはネットワークがありません
表面にあり、認証されていません。
すべての設定は e に存在します

tc/kaas.toml 。それをコピーして編集します。
[llm]
api_key = " sk-... "
Base_url = " https://api.openai.com/v1 "
モデル = " gpt-4o-mini "
「アイ。 MCP]
Enabled = false # /mcp エンドポイントを公開するには true を設定します
token = " " # MCP 認証用のベアラー トークン (空 = 認証なし)
timeout_sec = 120 # ツール/呼び出しタイムアウト
Docker では、シークレットを環境変数として渡します。環境変数は TOML をオーバーライドします。
起動:
すべてのサービスをローカルで開始する最も簡単な方法:
開発を行う
これにより、Go バックエンド (Python AI デーモンを自動生成します) と Vite 開発サーバーが一緒に起動します。
コンポーネントを個別に実行するには:
# バックエンド (Python デーモンを自動的に生成します)
./cmd/kaas -f etc/kaas.toml を実行します。
# フロントエンド (ホットリロード)
cd web && pnpm dev
# MCP サーバー (stdio - ローカル エージェント統合用)
cd py && KAAS_KB_DIR=./data uv run kb-ai mcp
# テスト
テストを行う
貢献する
貢献は歓迎です — 開発セットアップについては CONTRIBUTING.md を参照してください。
テストの実行方法とコミット規約。
中心的なアイデア — 知識を永続的に相互リンクされた Wiki に編集すること
各クエリに対して RAG を介して回答を再取得するのではなく、時間の経過とともに複雑になります。
Andrej Karpathy の「LLM Wiki」からインスピレーションを得た
要点。パターンを明確に表現していただきありがとうございます。
Knowledge as a Service: すぐに使える LLM Wiki コンパイラー。
bybit-exchange.github.io/kaas-doc/ リソース
Readme MIT ライセンス
貢献活動 カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Knowledge as a Service: an out-of-the-box LLM wiki compiler. - bybit-exchange/kaas

GitHub - bybit-exchange/kaas: Knowledge as a Service: an out-of-the-box LLM wiki compiler. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
bybit-exchange
/
kaas
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github .github cmd/ kaas cmd/ kaas docs docs etc etc internal internal py py scripts scripts web web .dockerignore .dockerignore .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md go.mod go.mod go.sum go.sum install.sh install.sh View all files Repository files navigation
Turn scattered notes, documents, and transcripts into a searchable, queryable personal Wiki — powered by LLM-driven knowledge compilation.
KaaS started as an internal tool. Our knowledge lived scattered across
documents, meetings, and email — and every time someone changed roles or left,
the context they'd built up walked out with them. New people spent weeks
piecing it back together.
A distillation pipeline fixed that. It compiles each person's scattered
material into a wiki tied to their role rather than their identity — so when
someone moves on, the raw data goes but the distilled judgment stays for whoever
fills the seat next.
The payoff is the same either way: the organization stops re-answering the same
questions. That's what convinced us it was worth open-sourcing.
Unlike typical RAG systems that chunk and embed raw text, KaaS compiles your content through a 4-phase LLM pipeline:
Raw Content → Extract → Classify → Write → Index → Structured Wiki
The result is human-readable Markdown articles — not a black-box vector store. You can read, edit, and git-manage your knowledge base.
Quick Start with your AI agent
Already living in a coding agent (Claude Code, Codex, openclaw, …)? Skip Docker.
Copy this and paste it to your agent — it will install kb-ai , ask what to
distill, build the wiki, and wire up MCP so you can query it in any later session:
Set up KaaS to build a queryable knowledge base from my files.
Fetch https://raw.githubusercontent.com/bybit-exchange/kaas/main/docs/agent-quickstart.md
and follow it exactly.
Prefer a web UI, or want the full backend? Use the Docker path below.
KaaS calls LLMs through any OpenAI-compatible API (OpenAI, DeepSeek, Ollama, vLLM, Azure OpenAI, etc.). Pick either method below:
# Build the image
docker build -t kaas .
# Run
docker run -d --name kaas \
-p 8080:8080 \
-v ./data:/app/data \
-e LLM_API_KEY=sk-xxx \
-e LLM_BASE_URL=https://api.openai.com/v1 \
-e LLM_MODEL=gpt-4o-mini \
kaas
Option B: CLI Install
# Install (Linux/macOS, amd64/arm64)
curl -fsSL https://raw.githubusercontent.com/bybit-exchange/kaas/main/install.sh | sh
# Start the service
export LLM_API_KEY= " sk-xxx " # OpenAI-compatible API key
export LLM_BASE_URL= " https://api.openai.com/v1 " # API endpoint
export LLM_MODEL= " gpt-4o-mini " # Model name
kaas serve # Default: http://localhost:8080
Supported platforms: Linux/macOS, amd64/arm64. Uninstall: rm -rf ~/.local/share/kaas ~/.local/bin/kaas .
LLM_BASE_URL defaults to https://api.openai.com/v1 and LLM_MODEL defaults to gpt-4o-mini .
Change them to point at any OpenAI-compatible endpoint.
To let Claude Code or other MCP clients connect to the knowledge base:
docker run -d --name kaas \
-p 8080:8080 \
-v ./data:/app/data \
-e LLM_API_KEY=sk-xxx \
-e KAAS_MCP_ENABLED=true \
-e KAAS_MCP_TOKEN=your-secret-token \
kaas
MCP client URL: http://<host>:8080/mcp , Authorization: Bearer your-secret-token .
The Go backend spawns the Python AI engine as a long-running daemon process, communicating via a multiplexed stdin/stdout protocol. A single Docker image bundles everything — no sidecar containers needed.
4-Phase Compile Pipeline : Extract concepts/entities/decisions → Classify into articles → Write/merge Markdown → Update markdown indexes
Worker Acceleration : Concurrent extract/pipeline workers, circuit breaker, lease recovery
Streaming Chat : SSE streaming with source citations pointing back to wiki articles
Multiple Input Sources : Paste text, upload files, or provide URLs
Incremental Compilation : Only recompiles new/changed content
Git-Friendly Output : All wiki articles are plain Markdown
MCP Access : Query the compiled wiki from any MCP-capable coding agent via an ask tool
Expose the compiled wiki to any Model Context Protocol
client (Claude Code, Codex, openclaw, …) through a single ask tool —
ask(query, paths?, model?) returns a cited Markdown answer grounded in the
wiki. Two transports:
stdio (local — the agent spawns the server, fully self-contained):
# The agent launches this; set KAAS_KB_DIR to the knowledge-base root and
# the LLM_* credentials in the environment.
kb-ai mcp # stdio is the default transport
# Claude Code:
claude mcp add kaas -- kb-ai mcp
For Codex / openclaw, add a stdio MCP server with command kb-ai mcp and env
KAAS_KB_DIR + LLM_* .
streamable-http (remote — published through the backend's :8080 origin):
Run the container with KAAS_MCP_ENABLED=true (see Quick Start ).
The backend exposes the MCP endpoint at /mcp . Point a remote agent at it:
# Claude Code:
claude mcp add --transport http kaas http://host:8080/mcp
Set KAAS_MCP_TOKEN to require Authorization: Bearer <token> on the HTTP
transport (off by default — local/intranet assumption). stdio has no network
surface and is unauthenticated.
All configuration lives in etc/kaas.toml . Copy and edit it:
[ llm ]
api_key = " sk-... "
base_url = " https://api.openai.com/v1 "
model = " gpt-4o-mini "
[ ai . mcp ]
enabled = false # set true to expose /mcp endpoint
token = " " # bearer token for MCP auth (empty = no auth)
timeout_sec = 120 # tools/call timeout
With Docker, pass secrets as environment variables — they override the TOML at
startup:
The quickest way to start all services locally:
make dev
This launches the Go backend (which auto-spawns the Python AI daemon) and the Vite dev server together.
To run components individually:
# Backend (spawns Python daemon automatically)
go run ./cmd/kaas -f etc/kaas.toml
# Frontend (hot-reload)
cd web && pnpm dev
# MCP server (stdio — for local agent integration)
cd py && KAAS_KB_DIR=./data uv run kb-ai mcp
# Tests
make test
Contributing
Contributions are welcome — see CONTRIBUTING.md for dev setup,
how to run the tests, and commit conventions.
The core idea — compiling knowledge into a persistent, interlinked wiki that
compounds over time instead of re-deriving answers via RAG on each query — was
inspired by Andrej Karpathy's "LLM Wiki"
gist. Thanks for the clear articulation of the pattern.
Knowledge as a Service: an out-of-the-box LLM wiki compiler.
bybit-exchange.github.io/kaas-doc/ Resources
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

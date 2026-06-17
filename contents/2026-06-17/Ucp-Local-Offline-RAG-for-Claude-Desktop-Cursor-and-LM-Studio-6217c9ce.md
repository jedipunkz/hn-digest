---
source: "https://github.com/akshay2211/universal-context-pipeline"
hn_url: "https://news.ycombinator.com/item?id=48567297"
title: "Ucp-Local – Offline RAG for Claude Desktop, Cursor, and LM Studio"
article_title: "GitHub - akshay2211/universal-context-pipeline: Universal Context Pipeline — local-first MCP server for grounding LLMs in your files · GitHub"
author: "akshay2211"
captured_at: "2026-06-17T10:39:17Z"
capture_tool: "hn-digest"
hn_id: 48567297
score: 1
comments: 0
posted_at: "2026-06-17T08:08:05Z"
tags:
  - hacker-news
  - translated
---

# Ucp-Local – Offline RAG for Claude Desktop, Cursor, and LM Studio

- HN: [48567297](https://news.ycombinator.com/item?id=48567297)
- Source: [github.com](https://github.com/akshay2211/universal-context-pipeline)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T08:08:05Z

## Translation

タイトル: UCP-Local – クロード デスクトップ、カーソル、LM Studio 用のオフライン RAG
記事のタイトル: GitHub - akshay2211/universal-context-pipeline: ユニバーサル コンテキスト パイプライン — ファイル内の LLM を接地するためのローカルファースト MCP サーバー · GitHub
説明: Universal Context Pipeline — ファイル内の LLM を接地するためのローカルファースト MCP サーバー - akshay2211/universal-context-pipeline

記事本文:
GitHub - akshay2211/universal-context-pipeline: ユニバーサル コンテキスト パイプライン — ファイル内の LLM を接地するためのローカルファースト MCP サーバー · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
アクシャイ2211
/
ユニバーサルコンテキストパイプライン
公共
通知
変更するにはサインインする必要があります

化設定
追加のナビゲーション オプション
コード
akshay2211/ユニバーサルコンテキストパイプライン
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット デモ デモ サンプル サンプル src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE POSITIONING.md POSITIONING.md README.md README.md ROADMAP.md ROADMAP.md ユニバーサル コンテキスト パイプライン仕様.md ユニバーサル コンテキスト パイプライン仕様.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
UCP — ユニバーサル コンテキスト パイプライン
LLM を独自のファイルに組み込む、ローカルファーストの MCP サーバー。
UCP は、マシン上のフォルダー (メモ、コード、会話エクスポート) のインデックスを作成し、MCP 互換クライアント (Claude Desktop、Cursor、LM Studio、およびその他のローカル エージェント ランタイム) に単一のツール search_local_context として公開します。ハイブリッド検索 (BM25 + ベクトル)、ツリーシッター対応コード チャンキング、完全な引用、コンテンツ ハッシュ埋め込みキャッシュ。単一のバイナリ。テレメトリはありません。雲はありません。
LM Studio のローカル モデル (または ucp-local ask を介した Ollama) と組み合わせると、スタック全体 (インデックス作成、埋め込み、取得、チャット モデル) が完全にオフラインで実行されます。飛行機内やエアギャップのある施設など、クラウド LLM が利用できない場所でも機能します。
会話の記憶 — 過去のすべてのクロード チャットを将来のすべてのセッションで検索可能にします。
エアギャップ RAG — ローカル Ollama + ローカル インデックス、ゼロ ネットワーク トラフィック。
クイック スタート — インストール、インデックス作成、問い合わせは 1 分以内に完了します。
あなたがそうであれば…
UCP が提供するのは…
クロード / カーソル / LM Studio パワー ユーザー
過去のすべての AI 会話の検索可能なアーカイブ。将来のセッションから search_local_context ツールとして呼び出すことができます。
ソフトウェアエンジニア
コード + プライベート ドキュメント + 兄弟リポジトリ + 過去のクロード チャットが 1 つの MCP ツールに統合

— ネイティブ インデクサーと並んで、カーソルまたはクロード コード内に表示されます。
研究者、作家、学者
マシンから何も出力せずに、行レベルの引用を使用して、根拠のある質問をできる PDF + ノート コーパス。
プライバシーが規制されたワークフロー (法律、医療、防衛、NDA に拘束された IP)
テレメトリもクラウドもゼロの単一の Rust バイナリ。 LM Studio と組み合わせて、完全にオフラインのエンドツーエンドの RAG スタックを実現します。
個人の創業者またはコンサルタント
folder_filter によるフォルダーごとのクライアント分離 — クライアント A のコンテキストがクライアント B のセッションに漏洩するリスクはありません。
完全な視聴者分析、競合比較、および UCP が勝つために明示的に構築されている 2 つのウェッジ: POSITIONING.md を参照してください。
v0.1、ヘッドレス。 ROADMAP.md のトラックスコープ。
ハイブリッド検索: SQLite FTS5 (BM25) ⨉ sqlite-vec (ANN) が相互ランク融合を介してマージされました。
Rust、Python、TypeScript/JavaScript のツリーシッター チャンク。見出しを意識したマークダウン。文に制限された散文フォールバック。
会話の記憶: Claudemessages.json エクスポートを取り込み、過去のチャットを検索します。
PII マスキングはデフォルトでオンになっています — 電子メール、OpenAI sk-、AWS キー、GitHub PAT、JWT。
コンテンツ ハッシュ埋め込みキャッシュ: 変更されていないコンテンツのインデックスを再作成すると、Ollama 呼び出しがゼロになります。
ファイルシステム ウォッチャー: ファイルを編集すると、インデックスが約 500 ミリ秒で更新されます。
デスクトップ UI / トレイ (延期 - 元の仕様でしたが、現在は ROADMAP Tier 2+ にあります)。
OS ホットキー インジェクターと HTTP プロキシ インターセプター (元の仕様から削除)。
OpenAI / Anthropic 埋め込みプロバイダー (現時点では Ollama のみ)。
カーソルおよび ChatGPT エクスポート形式 (クロードのみ。その他は後で)。
UCP には、マシン上に 3 つのものが必要です。Rust (ビルド用)、Ollama (埋め込みとオプションでチャット用)、および Poppler (堅牢な PDF テキスト抽出用 - 推奨)。
オラマ・ポプラーを醸造インストールする
ollam サーブ & # またはメニューバー アプリを使用する
オラマ プル nomic-embed-tex

t
# オプション、「ucp-local ask」の場合: ollama pull llama3.2
Linux (Debian/Ubuntu)
sudo apt install Poppler-utils
カール -fsSL https://ollama.com/install.sh |しー
オラマプル nomic-embed-text
# オプション、「ucp-local ask」の場合: ollama pull llama3.2
Linux (Fedora/RHEL)
sudo dnf install Poppler-utils
カール -fsSL https://ollama.com/install.sh |しー
オラマプル nomic-embed-text
窓
choco install Poppler ollam # またはそれぞれを手動でインストールします
オラマ プルノミック - 埋め込み - テキスト
Rust (安定版、エディション 2024) は、ソースからビルドする場合にのみ必要です。事前に構築された UCP バイナリをインストールする場合は、Rust のインストールをスキップしてください。
ポプラーはオプションですが推奨されます。これがないと、UCP は PDF に対してバンドルされた pdf-extract のみを使用します。このため、本文フォントに ToUnicode CMap が含まれていない PDF では問題が発生します (見出しは抽出されますが、本文テキストが失われます)。 PATH 上の Poppler からの pdftotext を使用すると、UCP は自動的にそれにフォールバックします。
名前に注意してください。クレートは crates.io で ucp-local として公開されます。つまり、裸の ucp 名が使用されます。 PATH 上のバイナリも ucp-local (コマンドラインで入力するもの) であり、ライブラリは use ucp_local::... としてインポートされます。
貨物インストール ucp-local
# `ucp-local` バイナリを PATH に置きます
ソースから
git clone < リポジトリ URL > ucp-local
cd ucp-ローカル
カーゴビルド --release
# target/release/ucp-local のバイナリ
カーゴインストール --path 。 # オプション。PATH に `ucp-local` を追加します。
使用法
# 1 つのフォルダーにインデックスを付ける
ucp-local インデックス ~ /Documents/notes
# 複数のフォルダーを同じストアにインデックス付けする
ucp-local インデックス ~ /Documents/notes ~ /code/my-project ~ /research
# フォルダーを監視し、変更に応じてインデックスを再作成します (最初のパスが最初に実行されます)
ucp-local watch ~ /code/my-project
# インデックスをクリアします — ソフト (埋め込みキャッシュを保持するため、再インデックスが高速になります)
ucp-ローカルクリア
# 1 つのフォルダーのチャンクのみをクリアします
ucp-local クリア ~ /Documents/notes

# ハード リセット — 埋め込みキャッシュも消去し、次のインデックスに再埋め込みを強制します
ucp-local クリア --hard --yes
# クロードのversations.json エクスポートを取り込む
ucp-local ingest-conversations ~ /Downloads/claude-export/conversations.json
# 設定 + インデックスのステータスを表示
UCP ローカル ステータス
# 標準入出力で MCP サーバーを実行します (これは MCP クライアントが起動するものです)
UCP ローカル サーブ
# ターミナルからインデックスを検索します (LLM なし) — 「インデックス作成が実際にこれをキャプチャしたか?」というデバッグに最適です。
ucp-local 検索「クエリはここにあります」
ucp-local 検索 " レート制限 " --folder ~ /code/my-project --limit 10
# 質問する — 内部で検索を実行し、ローカル チャット モデルが引用で回答します。
ucp-local は「トークン バケットがなくなったときにレート リミッターは何をしますか?」と尋ねます。
ucp-local ask "Q3 計画を要約します" --model qwen2.5
MCP クライアントを接続する
UCP は標準入出力経由で MCP を話すため、MCP サーバーを起動するすべてのクライアントがそれを使用できます。同じserveコマンドでもクライアントごとに異なる設定ファイル。
macOS の場合は ~/Library/Application Support/Claude/claude_desktop_config.json (Windows の場合は %APPDATA%\Claude\claude_desktop_config.json) に追加します。
{
"mcpサーバー": {
"ucp-ローカル" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " サーブ " ]
}
}
}
クロードデスクトップを再起動します。 search_local_context ツールが利用可能になります。インデックス付きファイルに基づいて何かを尋ねると、それらをインラインで引用します。
カーソルは ~/.cursor/mcp.json (またはプロジェクトごとの .cursor/mcp.json ) から MCP サーバーを読み取ります。
{
"mcpサーバー": {
"ucp-ローカル" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " サーブ " ]
}
}
}
カーソルをリロードします。チャット サイドバーには、search_local_context がツールとして表示されます。これは、カーソル自身の @codebase インデクサーが到達できないリポジトリやドキュメント (プライベート メモ、会話履歴、兄弟リポジトリ) にエージェントを定着させるのに役立ちます。
LM Studio 0.3.17+ は MCP をサポートします。オペ

チャット設定で、MCP サーバー セクションを見つけて、以下を追加します。
{
"mcpサーバー": {
"ucp-ローカル" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " サーブ " ]
}
}
}
UCP を LM Studio にダウンロードしたローカル モデル (Llama、Qwen、Mistral など) とペアリングします。これで、インデックス作成、埋め込み、取得、チャット モデルがすべて同じマシン上で実行されます (クラウドもネットワークも必要ありません)。LLM は引き続き search_local_context を呼び出して、その回答をファイルに根付かせることができます。
MCP 仕様に従うクライアント (Zed、Continue.dev、Goose、カスタム エージェント SDK アプリなど) は、同じコマンド + 引数の形式をとります。クライアントが JSON-RPC stdio サーバーを予期している場合は、それを ucp-localserve に指定すれば完了です。
~/.config/ucp/config.toml (またはプラットフォームの同等物 - ucp-local ステータスは解決されたパスを出力します)。すべてのフィールドはオプションです。デフォルトが表示されます:
[ オラマ ]
ホスト = " http://localhost:11434 "
embedding_model = " nomic-embed-text "
[チャンク化]
max_tokens = 512
オーバーラップ文 = 1
インデックスされるもの
拡張子: md 、 markdown 、 txt 、 rs 、 py 、 ts 、 tsx 、 js 、 jsx 、 mjs 、 go 、 pdf 。
PDF: テキストは pdf-extract によって抽出され、散文としてチャンク化されます。デジタルで生成された PDF (論文、ドキュメント、エクスポートされたメモ) に適しています。スキャンされた画像のみの PDF ではこの問題が発生します。OCR (v0.2 以降) が必要です。引用行番号は、PDF のページ番号ではなく、抽出された平文を参照します。ページ認識引用は v0.2 リストにあります。
スキップされたディレクトリ: .git 、 .idea 、 .vscode 、 target 、 node_modules 、 __pycache__ 、 .venv 、 venv 、 dist 、 build 、 .next 、 .nuxt 、coverage 、 .pytest_cache 、 .mypy_cache 。ドットファイルはスキップされます。
モジュール
役割
摂取
マスキング + フォーマットごとのチャンカー (散文 / マークダウン / ツリーシッター経由のコード) + ディスパッチャー
ストレージ
rusqlite + sqlite-vec + FTS5; RRF によるハイブリッド検索
埋め込み
OllamaClient + EmbeddingCache 経由のコンテンツ ハッシュ キャッシュ

::ハッシュ
インデクサ
ウォーク + 読み取り + チャンク + 埋め込み + 挿入。単一ファイルおよびバルクチャンクのパス
監視者
通知ベースのデバウンス再インデックス
mcp
JSON-RPC 2.0 stdio サーバー、1 つのツール: search_local_context
開発者向けのアーキテクチャの概要については CLAUDE.md を、元の (範囲が狭くなった) 設計ドキュメントについては Universal Context Pipeline Supplement.md を参照してください。
カーゴテスト # フルテストスイート
カーゴテスト --lib 取り込み # 1 つのモジュール
カーゴ実行 -- インデックス <パス> # 開発ビルドに対して反復処理します
RUST_LOG=ucp_local=info カーゴ実行 -- 監視 < パス > # 詳細
変更履歴
リリース履歴とメモは CHANGELOG.md に保存されています。現在公開されているバージョンは 0.1.0 ( crates.io ) です。
ユニバーサル コンテキスト パイプライン — ファイル内で LLM を接地するためのローカル ファースト MCP サーバー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 6 月 16 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Universal Context Pipeline — local-first MCP server for grounding LLMs in your files - akshay2211/universal-context-pipeline

GitHub - akshay2211/universal-context-pipeline: Universal Context Pipeline — local-first MCP server for grounding LLMs in your files · GitHub
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
akshay2211
/
universal-context-pipeline
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
akshay2211/universal-context-pipeline
master Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits demo demo samples samples src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE POSITIONING.md POSITIONING.md README.md README.md ROADMAP.md ROADMAP.md Universal Context Pipeline Specification.md Universal Context Pipeline Specification.md View all files Repository files navigation
UCP — Universal Context Pipeline
A local-first MCP server that grounds LLMs in your own files.
UCP indexes folders on your machine — notes, code, conversation exports — and exposes them to any MCP-compatible client (Claude Desktop, Cursor, LM Studio, and other local-agent runtimes) as a single tool: search_local_context . Hybrid retrieval (BM25 + vector), tree-sitter-aware code chunking, full citations, content-hash embedding cache. Single binary. No telemetry. No cloud.
Paired with a local model in LM Studio (or Ollama via ucp-local ask ), the whole stack — indexing, embeddings, retrieval, and the chat model — runs fully offline. Works on a plane, in an air-gapped facility, or anywhere a cloud LLM isn't an option.
Conversation memory — make every past Claude chat searchable across every future session.
Air-gap RAG — local Ollama + local index, zero network traffic.
Quick start — install, index, ask, in under a minute.
If you are…
UCP gives you…
A Claude / Cursor / LM Studio power user
A searchable archive of every past AI conversation, callable from any future session as the search_local_context tool.
A software engineer
Code + private docs + sibling repos + past Claude chats unified under one MCP tool — surfaced inside Cursor or Claude Code alongside their native indexers.
A researcher, writer, or academic
A PDF + notes corpus you can ask grounded questions against, with line-level citations, without anything leaving the machine.
In a privacy-regulated workflow (legal, medical, defense, NDA-bound IP)
A single Rust binary with zero telemetry and zero cloud. Pair with LM Studio for a fully offline, end-to-end RAG stack.
A solo founder or consultant
Per-folder client isolation via folder_filter — no risk of leaking client A's context into client B's session.
Full audience analysis, competitive comparison, and the two wedges UCP is explicitly built to win on: see POSITIONING.md .
v0.1, headless. Track scope in ROADMAP.md .
Hybrid search: SQLite FTS5 (BM25) ⨉ sqlite-vec (ANN) merged via reciprocal-rank fusion.
Tree-sitter chunking for Rust, Python, TypeScript/JavaScript. Heading-aware Markdown. Sentence-bounded prose fallback.
Conversation memory: ingest your Claude conversations.json export and search across past chats.
PII masking on by default — email, OpenAI sk- , AWS keys, GitHub PATs, JWT.
Content-hash embedding cache: re-indexing unchanged content makes zero Ollama calls.
Filesystem watcher: edit a file, the index updates in ~500ms.
Desktop UI / tray (deferred — was in original spec, now in ROADMAP tier 2+).
OS hotkey injector and HTTP proxy interceptor (cut from the original spec).
OpenAI / Anthropic embedding providers (Ollama only for now).
Cursor and ChatGPT export formats (Claude only; others later).
UCP needs three things on your machine: Rust (to build), Ollama (to embed and optionally chat), and Poppler (for robust PDF text extraction — recommended).
brew install ollama poppler
ollama serve & # or use the menu-bar app
ollama pull nomic-embed-text
# Optional, for `ucp-local ask`: ollama pull llama3.2
Linux (Debian/Ubuntu)
sudo apt install poppler-utils
curl -fsSL https://ollama.com/install.sh | sh
ollama pull nomic-embed-text
# Optional, for `ucp-local ask`: ollama pull llama3.2
Linux (Fedora/RHEL)
sudo dnf install poppler-utils
curl -fsSL https://ollama.com/install.sh | sh
ollama pull nomic-embed-text
Windows
choco install poppler ollama # or install each manually
ollama pull nomic - embed - text
Rust (stable, edition 2024) is needed only to build from source. If you install a pre-built UCP binary, skip the Rust install.
Poppler is optional but recommended. Without it, UCP only uses the bundled pdf-extract for PDFs, which struggles with PDFs whose body fonts lack a ToUnicode CMap (you'll see headings extract but body text go missing). With pdftotext from Poppler on PATH, UCP falls back to it automatically.
Note on the name. The crate is published as ucp-local on crates.io — the bare ucp name was taken. The binary on your PATH is also ucp-local (that's what you type on the command line), and the library is imported as use ucp_local::... .
cargo install ucp-local
# Puts the `ucp-local` binary on your PATH
From source
git clone < repo-url > ucp-local
cd ucp-local
cargo build --release
# Binary at target/release/ucp-local
cargo install --path . # optional, to put `ucp-local` on your PATH
Usage
# Index one folder
ucp-local index ~ /Documents/notes
# Index multiple folders into the same store
ucp-local index ~ /Documents/notes ~ /code/my-project ~ /research
# Watch a folder and re-index on changes (initial pass runs first)
ucp-local watch ~ /code/my-project
# Clear the index — soft (keeps the embedding cache so re-index is fast)
ucp-local clear
# Clear only one folder's chunks
ucp-local clear ~ /Documents/notes
# Hard reset — also wipes the embedding cache, forces re-embed on next index
ucp-local clear --hard --yes
# Ingest a Claude conversations.json export
ucp-local ingest-conversations ~ /Downloads/claude-export/conversations.json
# Show config + index status
ucp-local status
# Run the MCP server over stdio (this is what MCP clients launch)
ucp-local serve
# Search the index from the terminal (no LLM) — best for debugging "did indexing actually capture this?"
ucp-local search " your query here "
ucp-local search " rate limiting " --folder ~ /code/my-project --limit 10
# Ask a question — runs search internally, then a local chat model answers with citations
ucp-local ask " what does the rate limiter do when a token bucket runs out? "
ucp-local ask " summarize my Q3 plan " --model qwen2.5
Wire up an MCP client
UCP speaks MCP over stdio, so any client that launches MCP servers can use it. Same serve command, different config file per client.
Add to ~/Library/Application Support/Claude/claude_desktop_config.json on macOS ( %APPDATA%\Claude\claude_desktop_config.json on Windows):
{
"mcpServers" : {
"ucp-local" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " serve " ]
}
}
}
Restart Claude Desktop. The search_local_context tool will be available — ask something grounded in your indexed files and it'll cite them inline.
Cursor reads MCP servers from ~/.cursor/mcp.json (or per-project .cursor/mcp.json ):
{
"mcpServers" : {
"ucp-local" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " serve " ]
}
}
}
Reload Cursor. The chat sidebar will surface search_local_context as a tool — useful for grounding the agent in repos and docs Cursor's own @codebase indexer can't reach (private notes, conversation history, sibling repos).
LM Studio 0.3.17+ supports MCP. Open the chat settings, find the MCP servers section, and add:
{
"mcpServers" : {
"ucp-local" : {
"command" : " /full/path/to/ucp-local " ,
"args" : [ " serve " ]
}
}
}
Pair UCP with any local model you've downloaded in LM Studio (Llama, Qwen, Mistral, etc.). Now your indexing, embeddings, retrieval, and chat model all run on the same machine — no cloud, no network — and the LLM can still call search_local_context to ground its answers in your files.
Any client following the MCP spec (Zed, Continue.dev, Goose, custom Agent SDK apps, etc.) takes the same command + args shape. If your client expects a JSON-RPC stdio server, point it at ucp-local serve and you're done.
~/.config/ucp/config.toml (or the platform equivalent — ucp-local status prints the resolved path). All fields optional; defaults shown:
[ ollama ]
host = " http://localhost:11434 "
embedding_model = " nomic-embed-text "
[ chunking ]
max_tokens = 512
overlap_sentences = 1
What gets indexed
By extension: md , markdown , txt , rs , py , ts , tsx , js , jsx , mjs , go , pdf .
PDFs: text is extracted via pdf-extract and chunked as prose. Works well for digitally generated PDFs (papers, docs, exported notes). Falls down on scanned image-only PDFs — those need OCR (v0.2+). Citation line numbers reference the extracted plaintext, not PDF page numbers; page-aware citations are on the v0.2 list.
Skipped directories: .git , .idea , .vscode , target , node_modules , __pycache__ , .venv , venv , dist , build , .next , .nuxt , coverage , .pytest_cache , .mypy_cache . Dotfiles are skipped.
Module
Role
ingestion
Masking + per-format chunkers (prose / markdown / code via tree-sitter) + dispatcher
storage
rusqlite + sqlite-vec + FTS5; hybrid search via RRF
embeddings
OllamaClient + content-hash cache via EmbeddingCache::hash
indexer
Walk + read + chunk + embed + insert; single-file and bulk-chunk paths
watcher
notify -based debounced re-index
mcp
JSON-RPC 2.0 stdio server, one tool: search_local_context
See CLAUDE.md for the developer-facing architecture summary, and Universal Context Pipeline Specification.md for the original (now narrower in scope) design doc.
cargo test # full test suite
cargo test --lib ingestion # one module
cargo run -- index < path > # iterate against the dev build
RUST_LOG=ucp_local=info cargo run -- watch < path > # verbose
Changelog
Release history and notes live in CHANGELOG.md . The current published version is 0.1.0 ( crates.io ).
Universal Context Pipeline — local-first MCP server for grounding LLMs in your files
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jun 16, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/braincompany/sessiongrep"
hn_url: "https://news.ycombinator.com/item?id=49087850"
title: "Show HN: Sessiongrep – Local-First Memory Across Claude Code, Codex, and Cursor"
article_title: "GitHub - braincompany/sessiongrep: Local-first memory layer for CLI agents. Indexes Claude Code, Codex CLI, and Cursor session histories into SQLite + FTS5 — searchable from a CLI, TUI, or MCP server so your next agent session can recall the last one. · GitHub"
author: "npx88"
captured_at: "2026-07-28T19:08:38Z"
capture_tool: "hn-digest"
hn_id: 49087850
score: 1
comments: 0
posted_at: "2026-07-28T18:19:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sessiongrep – Local-First Memory Across Claude Code, Codex, and Cursor

- HN: [49087850](https://news.ycombinator.com/item?id=49087850)
- Source: [github.com](https://github.com/braincompany/sessiongrep)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T18:19:22Z

## Translation

タイトル: HN を表示: Sessiongrep – クロード コード、コーデックス、およびカーソルにわたるローカル ファースト メモリ
記事のタイトル: GitHub - Braincompany/sessiongrep: CLI エージェントのローカルファースト メモリ層。 Claude Code、Codex CLI、および Cursor セッション履歴を SQLite + FTS5 にインデックス付けします。CLI、TUI、または MCP サーバーから検索できるため、次のエージェント セッションで最後のセッションを呼び出すことができます。 · GitHub
説明: CLI エージェントのローカルファーストメモリ層。 Claude Code、Codex CLI、および Cursor セッション履歴を SQLite + FTS5 にインデックス付けします。CLI、TUI、または MCP サーバーから検索できるため、次のエージェント セッションで最後のセッションを呼び出すことができます。 - ブレインカンパニー/セッショングレップ

記事本文:
GitHub - Braincompany/sessiongrep: CLI エージェントのローカルファーストメモリ層。 Claude Code、Codex CLI、および Cursor セッション履歴を SQLite + FTS5 にインデックス付けします。CLI、TUI、または MCP サーバーから検索できるため、次のエージェント セッションで最後のセッションを呼び出すことができます。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。レル

oad を使用してセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ブレインカンパニー
/
セッショングレップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .github/ workflows .github/ workflows docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix すべてのファイルを表示 リポジトリ ファイルナビゲーション
あなたは先週そのバグを解決しました。次のエージェントセッションはわかりません。
CLI エージェントのローカルファーストメモリ層。 sessiongrep は、Claude Code、Codex CLI、Cursor、Antigravity、および Pi セッション履歴を単一の SQLite + FTS5 データベースにインデックス付けし、トピック、リポジトリ、プロバイダー、または最近の作業ごとに古い作品を見つけるための 1 つの CLI/TUI を提供します。また、エージェントが自身の履歴を検索できるように、MCP サーバーも同梱されています。
実際の成果は移植可能なコンテキストです。セッション履歴が 1 つのツールに閉じ込められることがありません。 Claude Code で開始した作業は Codex でも継続でき、エージェントは使用するすべてのツールにわたって独自の以前の推論を復元し、さらには批判することもできます。
発表を読んでください: Sessiongrep: CLI エージェント用のローカルファーストメモリ層。
セッションのトランスクリプトはすでにマシン上に存在しており、不透明なファイル名を持つノイズの多い JSONL として ~/.claude/projects 、 ~/.codex/sessions 、 ~/.cursor/projects に分散しています。情報が不足しているのではなく、行き詰まっているのです。人間はそれを読みたくありません。エージェントはそれを取得する方法を知りません。 JSONL 上の grep はツールのペイロードに埋もれてしまいます。シェル履歴はコマンドをキャプチャしますが、推論はキャプチャしません。クラウド同期またはベクター支援の代替手段は、自分のものではないシステムにシークレットと URL を持ち込みます。
セッション

grep はリコールをローカルに保ちます。 1 つの静的バイナリ、1 つの SQLite ファイル、デーモンやサーバーはありません。インデックスは使い捨てキャッシュです。必要に応じていつでも削除して再構築できます。
プロバイダー アダプターは、Claude、Codex、Cursor、Antigravity、および Pi トランスクリプトを単一のセッション モデルに正規化し、トランスクリプト テキスト、タイトル、概要、およびプレビュー上の FTS5 仮想テーブルを使用して SQLite (WAL モード) に書き込みます。すべての読み取りコマンドは最初に増分再インデックスを実行します。mtime とサイズが変更されていないファイルはスキップされるため、履歴が大きくなっても検索とリストは高速に実行されます。
Claude Code、Codex CLI、および/または Cursor がインストールされている (セッション データ用)
git clone git@github.com:braincompany/sessiongrep.git
cd セッショングレップ
# 両方のバイナリをインストールする
カーゴインストール --path 。
これにより、2 つのバイナリが ~/.cargo/bin/ にインストールされます。
~/.cargo/bin が PATH にあることを確認してください。まだ存在しない場合は、~/.bashrc または ~/.zshrc に追加します。
エクスポート PATH= " $HOME /.cargo/bin: $PATH "
セッションにインデックスを付ける
インデックスは自動的に更新されます。すべてのコマンド (search、list、tui など) は、実行前に増分再インデックスを実行します。 cron ジョブや手動手順は必要ありません。
最初から完全な再構築を強制するには:
sessiongrep の再インデックス --full
クイックスタート
sessiongrep list --limit 20 最近のセッション数 (初回実行時の自動インデックス)
sessiongrep 検索 "認証バグ" # キーワード検索
sessiongrep 検索 " redis " --provider codex
sessiongrep 検索 " datadog " --プロバイダー カーソル
sessiongrep 検索「temporal」 --provider pi
sessiongrep show claude:79accec8-5bf5-415b-a4a5-fe370eb2c998
sessiongrep 再開 79accec8 --dry-run
sessiongrep import 79accec8 --format マークダウン
sessiongrep 医師 # 健康診断
sessiongrep tui # インタラクティブブラウザ
MCPサーバーのセットアップ
MCP サーバーを使用すると、AI エージェントが過去のセッションをプログラムで検索して取得できます。コンテキストをコピー＆ペーストする必要はありません。

昔の会話。
クロード mcp add --scope user --transport stdio sessiongrep -- sessiongrep-mcp
コーデックス CLI
codex mcp add sessiongrep -- sessiongrep-mcp
検証する
新しいセッションを開始し、次のようなプロンプトを試してください。
「Datadog メトリクスを設定していた前のセッションを見つけてください」
エージェントは、search_session を呼び出して一致を検索し、get_session を呼び出して関連するコンテキストを取得します。
ツール
説明
検索セッション
オプションのプロバイダー フィルターと制限を使用して、キーワードでセッションを検索します
セッションの取得
セッション ID によって完全なトランスクリプトとメタデータを取得 (コンテキストを制限する max_lines をサポート)
リストセッション
最近のセッションをリストします。プロバイダーとパス プレフィックスでフィルタリング可能
リポジトリのタイムライン
リポジトリ パス プレフィックスの 1 日単位のメタデータ タイムライン - 時間の経過とともに何が変化したかを詳細に表示
get_resume_command
CLI コマンドを取得してネイティブ ツールでセッションを再開する
構成
~/.config/sessiongrep/config.toml にあるオプションの構成ファイル:
[プロバイダー。クロード]
有効 = true
パス = [ " ~/.claude/projects " ]
[プロバイダー。コーデックス]
有効 = true
パス = [ " ~/.codex/sessions " ]
[プロバイダー。カーソル]
有効 = true
パス = [ " ~/.cursor/projects " ]
[プロバイダー。反重力]
有効 = true
パス = [ " ~/.gemini/antigravity/brain " ]
[プロバイダー。パイ]
有効 = true
パス = [ " ~/.pi/agent/sessions " ]
[インデックス]
db_path = " ~/.local/share/sessiongrep/index.db "
ache_dir = " ~/.cache/sessiongrep "
[ウイ]
プレビュー行数 = 30
【検索】
デフォルト_制限 = 50
prefer_current_repo = true
プライバシーとデータ
すべてがマシン上に残ります。ネットワーク通話、テレメトリ、クラウド同期はありません。
このツールは読み取り専用であり、セッション ファイルを変更することはありません。
SQLite インデックスは派生キャッシュです。いつでも削除して、再インデックス --full を実行すると、トランスクリプトから再構築できます。
すべてのパス (データベース、キャッシュ、構成) は、 ~/.local/share 、 ~/.c の下にあるユーザーローカルです。

ache 、および ~/.config 。
ネイティブ プロバイダー CLI へのデリゲートを再開します ( claude --resume <id> 、 codex submit <id> 、または pi --session <id> )。カーソルと反重力の再開は現在サポートされていません。
クロード、カーソル、および Pi サブエージェントのトランスクリプトは、レコードの重複を避けるためにインデックス作成から除外されます。
初期ですが使用可能 — プレリリース、ソースから構築 (タグ付きリリースはまだありません)。 CLI サーフェスと MCP ツールの名前は安定したままになる可能性があります。ディスク上のインデックス スキーマは依然として変更される可能性があります (スキーマの不一致が発生した場合は、~/.local/share/sessiongrep/index.db を削除し、再構築します)。
問題やプルリクエストは大歓迎です。バグについては、プロバイダーのバージョンと sessiongrep ドクターの出力を含めてください。機能については、PR を送信する前に範囲について簡単に話し合うことで、物事がスムーズに進みます。
CLI エージェントのローカルファーストメモリ層。 Claude Code、Codex CLI、および Cursor セッション履歴を SQLite + FTS5 にインデックス付けします。CLI、TUI、または MCP サーバーから検索できるため、次のエージェント セッションで最後のセッションを呼び出すことができます。
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
5 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local-first memory layer for CLI agents. Indexes Claude Code, Codex CLI, and Cursor session histories into SQLite + FTS5 — searchable from a CLI, TUI, or MCP server so your next agent session can recall the last one. - braincompany/sessiongrep

GitHub - braincompany/sessiongrep: Local-first memory layer for CLI agents. Indexes Claude Code, Codex CLI, and Cursor session histories into SQLite + FTS5 — searchable from a CLI, TUI, or MCP server so your next agent session can recall the last one. · GitHub
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
braincompany
/
sessiongrep
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .github/ workflows .github/ workflows docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
You solved that bug last week. Your next agent session has no idea.
A local-first memory layer for CLI agents. sessiongrep indexes your Claude Code, Codex CLI, Cursor, Antigravity, and Pi session histories into a single SQLite + FTS5 database, then gives you one CLI/TUI to find old work by topic, repo, provider, or recency. It also ships an MCP server so your agent can search its own history.
The real payoff is portable context: your session history isn't trapped in one tool. Work you started in Claude Code can continue in Codex, and an agent can recover — and even critique — its own prior reasoning across every tool you use.
Read the announcement: Sessiongrep: a local-first memory layer for CLI agents .
Session transcripts already live on your machine — scattered across ~/.claude/projects , ~/.codex/sessions , ~/.cursor/projects as noisy JSONL with opaque filenames. The information is not missing, it's stranded. Humans don't want to read it; agents don't know how to retrieve it. Grep over JSONL drowns in tool payloads. Shell history captures commands but not reasoning. Cloud-synced or vector-backed alternatives bring secrets and URLs into systems that aren't yours.
sessiongrep keeps recall local. One static binary, one SQLite file, no daemon, no server. The index is a disposable cache — delete it and rebuild it whenever you want.
Provider adapters normalize Claude, Codex, Cursor, Antigravity, and Pi transcripts into a single Session model and write them into SQLite (WAL mode) with an FTS5 virtual table over transcript text, title, summary, and preview. Every read command runs an incremental reindex first — files whose mtime and size haven't changed are skipped, so search and list stay fast even as your history grows.
Claude Code, Codex CLI, and/or Cursor installed (for session data)
git clone git@github.com:braincompany/sessiongrep.git
cd sessiongrep
# Install both binaries
cargo install --path .
This installs two binaries to ~/.cargo/bin/ :
Make sure ~/.cargo/bin is in your PATH. Add to your ~/.bashrc or ~/.zshrc if not already present:
export PATH= " $HOME /.cargo/bin: $PATH "
Index your sessions
The index updates automatically — every command (search, list, tui, etc.) runs an incremental reindex before executing. No cron jobs or manual steps needed.
To force a full rebuild from scratch:
sessiongrep reindex --full
Quick start
sessiongrep list --limit 20 # recent sessions (auto-indexes on first run)
sessiongrep search " auth bug " # keyword search
sessiongrep search " redis " --provider codex
sessiongrep search " datadog " --provider cursor
sessiongrep search " temporal " --provider pi
sessiongrep show claude:79accec8-5bf5-415b-a4a5-fe370eb2c998
sessiongrep resume 79accec8 --dry-run
sessiongrep export 79accec8 --format markdown
sessiongrep doctor # health check
sessiongrep tui # interactive browser
MCP server setup
The MCP server lets AI agents search and retrieve your past sessions programmatically — no copy-pasting context from old conversations.
claude mcp add --scope user --transport stdio sessiongrep -- sessiongrep-mcp
Codex CLI
codex mcp add sessiongrep -- sessiongrep-mcp
Verify
Start a new session and try a prompt like:
"Find my previous session where I was setting up Datadog metrics"
The agent will call search_sessions to find matches and get_session to pull in relevant context.
Tool
Description
search_sessions
Search sessions by keyword, with optional provider filter and limit
get_session
Get full transcript and metadata by session ID (supports max_lines to limit context)
list_sessions
List recent sessions, filterable by provider and path prefix
timeline_for_repo
Day-bucketed metadata timeline for a repo path prefix — scoped view of what changed over time
get_resume_command
Get the CLI command to resume a session in its native tool
Config
Optional config file at ~/.config/sessiongrep/config.toml :
[ providers . claude ]
enabled = true
paths = [ " ~/.claude/projects " ]
[ providers . codex ]
enabled = true
paths = [ " ~/.codex/sessions " ]
[ providers . cursor ]
enabled = true
paths = [ " ~/.cursor/projects " ]
[ providers . antigravity ]
enabled = true
paths = [ " ~/.gemini/antigravity/brain " ]
[ providers . pi ]
enabled = true
paths = [ " ~/.pi/agent/sessions " ]
[ index ]
db_path = " ~/.local/share/sessiongrep/index.db "
cache_dir = " ~/.cache/sessiongrep "
[ ui ]
preview_lines = 30
[ search ]
default_limit = 50
prefer_current_repo = true
Privacy & data
Everything stays on your machine. No network calls, no telemetry, no cloud sync.
The tool is read-only — it never modifies your session files.
The SQLite index is a derived cache. Delete it anytime and reindex --full rebuilds it from your transcripts.
All paths (database, cache, config) are user-local under ~/.local/share , ~/.cache , and ~/.config .
Resume delegates to the native provider CLI ( claude --resume <id> , codex resume <id> , or pi --session <id> ). Cursor and Antigravity resume are not currently supported.
Claude, Cursor, and Pi subagent transcripts are excluded from indexing to avoid duplicate records.
Early but usable — pre-release, built from source (no tagged release yet). The CLI surface and MCP tool names are likely to stay stable; the on-disk index schema may still change (delete ~/.local/share/sessiongrep/index.db and let it rebuild if you hit a schema mismatch).
Issues and pull requests are welcome. For bugs, please include your provider versions and a sessiongrep doctor output. For features, a quick issue to discuss scope before sending a PR keeps things moving.
Local-first memory layer for CLI agents. Indexes Claude Code, Codex CLI, and Cursor session histories into SQLite + FTS5 — searchable from a CLI, TUI, or MCP server so your next agent session can recall the last one.
Readme Apache-2.0 license Activity Custom properties Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

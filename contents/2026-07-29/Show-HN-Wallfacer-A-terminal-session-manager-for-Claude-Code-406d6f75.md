---
source: "https://github.com/pradipta/wallfacer"
hn_url: "https://news.ycombinator.com/item?id=49096639"
title: "Show HN: Wallfacer: A terminal session manager for Claude Code"
article_title: "GitHub - pradipta/wallfacer: A terminal session manager for Claude Code, and more · GitHub"
author: "pradiptasarma"
captured_at: "2026-07-29T12:56:02Z"
capture_tool: "hn-digest"
hn_id: 49096639
score: 1
comments: 1
posted_at: "2026-07-29T12:32:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wallfacer: A terminal session manager for Claude Code

- HN: [49096639](https://news.ycombinator.com/item?id=49096639)
- Source: [github.com](https://github.com/pradipta/wallfacer)
- Score: 1
- Comments: 1
- Posted: 2026-07-29T12:32:14Z

## Translation

タイトル: HN を表示: Wallfacer: クロード コードのターミナル セッション マネージャー
記事のタイトル: GitHub - pradipta/wallfacer: クロード コード用のターミナル セッション マネージャー、その他 · GitHub
説明: クロード コードなどのターミナル セッション マネージャー - pradipta/wallfacer

記事本文:
GitHub - pradipta/wallfacer: クロード コード用のターミナル セッション マネージャーなど · GitHub
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
プラディプタ
/
ウォールフェイサー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

オプションについて
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .claude .claude .github/ workflows .github/ workflows .idea .idea cmd cmd docs docs 内部 内部 .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.mdデモ.キャスト デモ.キャストデモ.gif デモ.gif go.mod go.mod go.sum go.sum img.png img.png main.go main.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
██╗ ██╗█████╗ ██╗ ██╗ ███████╗█████╗ ██████╗███████╗██████╗
██║ ██║██╔══██╗██║ ██║ ██╔════╝██╔══██╗██╔════╝██╔════╝██╔══██╗
██║ █╗ ██║███████║██║ ██║ █████╗ ███████║██║ █████╗ ██████╔╝
██║███╗██║██╔══██║██║ ██║ ██╔══╝ ██╔══██║██║ ██╔══╝ ██╔══██╗
╚███╔███╔╝██║ ██║███████╗███████╗██║ ██║ ██║╚██████╗███████╗██║ ██║
╚══╝╚══╝ ╚═╝ ╚═╝╚══════╝╚══════╝╚═╝ ╚═╝ ╚═════╝╚══════╝╚═╝ ╚═╝
ウォールフェイサー
Claude Code のターミナル セッション マネージャー — すべての AI コーディングを参照
これまでに開始したセッションを選択し、セッションに名前を付け、タグ付け、グループ化、検索、再開、または削除します。
全画面ブラウザまたはコマンドラインから直接。
Claude Code は、すべての会話を無題の JSONL ファイルとして ~/ に保存します。

クロード/プロジェクト/ 、
数週間後には、何十ものトランスクリプトが作成されます。
区別できず、必要なものを見つける方法もありません。
Wallfacer はそれらすべてにインデックスを付けます (読み取り専用で、クロードのファイルには決して触れません)。
タイトル、タグ、プロジェクトを独自のローカル SQLite データベースに保存します。
すべてを 1 つのビューで表示 - すべてのセッション、すべてのディレクトリから、最新順に並べ替え
整理 — セッションの名前変更、タグ付け、プロジェクトへのグループ化
検索 — タイトル、最初のプロンプト、ディレクトリ、プロジェクト、タグ全体にわたって
起動と再開 — どこからでも新しいセッションを開始したり、古いセッションに戻ったりできます
安全な削除 — rm はゴミ箱に移動します。 --purge のみが永続的です
TUI および CLI — 人間用の全画面ブラウザ、スクリプト用のサブコマンド + --json
拡張可能 — エージェントはプラグイン可能なアダプターです。 Codex、opencode、Cursor CLI、および kiro-cli がロードマップに含まれています
github.com/pradipta/wallfacer@latest をインストールしてください
Go 1.22 以降が必要です。 CGO なしの Pure Go — macOS と Linux。事前に構築されたバイナリは
リリースページ ;ソースからのビルドは
開発ガイドで説明されています。
Wallfacer は 2 つのフロントエンドを持つ 1 つのバイナリで、どちらを取得できるかは、
サブコマンドを渡します:
これらは別個のツールではなく、相互に切り替えるものはありません。読み取りと書き込みの両方が可能です。
同じ SQLite インデックスなので、ブラウザでタグ付けしたセッションは次の方法ですぐに見つけられます。
Wallfacer list --tag 、またはその逆。 TUI で実行できることはすべて、サブコマンドでも実行できます。
(stdout が端末ではない場合 (wallfacer |less 、またはスクリプト内である場合)、裸のコマンド
ブラウザを開く代わりにヘルプを表示します)。
ブラウザを開いて、そこから作業します。
ウォールフェイサー
または、シェルから実行します。
Wallfacer new ~ /work/api --title "不安定な認証テストを修正する"
Wallfacer 再開「不安定な認証テストを修正」 # タイトルまたは ID プレフィックスで
壁面

r 検索認証
Wallfacer list --project api --json # スクリプト用
TUI — ウォールフェイサー
Bare Wallfacer は全画面セッション ブラウザを開きます。左側にリストと詳細が表示されます。
右側のペインには、強調表示されているすべてのウォールフェイサー ショーのプリントが表示されます。
すべての行には、プロジェクト、タグ、ディレクトリ、年齢、エージェントが同じメタデータとして表示されます。
ウォールフェイサーのリストが印刷されます。 100 列以下では、詳細ペインが脇に移動し、リストが表示されます。
全幅をとります。
すべてのサブコマンドはワンショットであり、実行、出力、終了します。ブラウザと同じインデックス。
<ref> は ID プレフィックス (resume 5f2) または正確なタイトル (resume "smoke test") です - 曖昧です
参考文献では、推測ではなく候補をリストします。ウォールフェイサーの外側で開始されたセッションは、
自動的にピックアップされます。インポート手順はありません。
Wallfacer は ~/.claude/projects/ をスキャンし、各セッション ファイルの先頭だけを読み取ります。
作業ディレクトリ、タイムスタンプ、および最初のプロンプト (自動タイトル)。メタデータは次の場所に存在します
~/.local/share/wallfacer/ の SQLite — これを削除すると、オーバーレイのみが失われ、
会話。同期は増分であるため、何百ものセッションが行われても高速な状態が維持されます。
他のエージェントは小さなアダプター インターフェイスを介して接続します。「」を参照してください。
docs/adding-an-agent.md 。
他のエージェント用のアダプター: Codex 、
opencode 、カーソル CLI、kiro-cli
セッションコンテンツ全体の全文検索 (SQLite FTS5)
Wallfacer 復元 (CLI からゴミ箱から削除)
セッショントランスクリプトをMarkdownにエクスポートする
統計 (プロジェクトごとのセッション数/週、ディスク使用量)
デスクトップ アプリ - サイドバーのセッション、埋め込みターミナル、複数のタブ
(Wails + xterm.js を同じ Go 内部構造上に構築)
ビルド、テスト、リリースについては、開発ガイドを参照してください。
貢献を歓迎します — 特に新しいエージェント アダプター。
この名前は、Liu Cixin の The Dark Forest の Wallfacers から借用したものです — 人々は委託されています
計画あり

あまりにも広大すぎて、他の人はフォローできません。
Claude Code などのターミナル セッション マネージャー
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A terminal session manager for Claude Code, and more - pradipta/wallfacer

GitHub - pradipta/wallfacer: A terminal session manager for Claude Code, and more · GitHub
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
pradipta
/
wallfacer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .claude .claude .github/ workflows .github/ workflows .idea .idea cmd cmd docs docs internal internal .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md demo.cast demo.cast demo.gif demo.gif go.mod go.mod go.sum go.sum img.png img.png main.go main.go View all files Repository files navigation
██╗ ██╗█████╗ ██╗ ██╗ ███████╗█████╗ ██████╗███████╗██████╗
██║ ██║██╔══██╗██║ ██║ ██╔════╝██╔══██╗██╔════╝██╔════╝██╔══██╗
██║ █╗ ██║███████║██║ ██║ █████╗ ███████║██║ █████╗ ██████╔╝
██║███╗██║██╔══██║██║ ██║ ██╔══╝ ██╔══██║██║ ██╔══╝ ██╔══██╗
╚███╔███╔╝██║ ██║███████╗███████╗██║ ██║ ██║╚██████╗███████╗██║ ██║
╚══╝╚══╝ ╚═╝ ╚═╝╚══════╝╚══════╝╚═╝ ╚═╝ ╚═╝ ╚═════╝╚══════╝╚═╝ ╚═╝
wallfacer
A terminal session manager for Claude Code — see every AI coding
session you've ever started, then name, tag, group, search, resume, or delete them, from a
full-screen browser or straight from the command line.
Claude Code stores every conversation as an untitled JSONL file under ~/.claude/projects/ ,
keyed by whatever directory you were in. After a few weeks you have dozens of transcripts you
can't tell apart and no way to find the one you need.
wallfacer indexes them all — read-only, it never touches Claude's files — and keeps your
titles, tags, and projects in its own local SQLite database.
One view of everything — every session, from every directory, sorted by recency
Organize — rename sessions, tag them, group them into projects
Search — across titles, first prompts, directories, projects, and tags
Launch & resume — start new sessions or jump back into old ones, from anywhere
Safe deletes — rm moves to trash; only --purge is permanent
TUI and CLI — a full-screen browser for humans, subcommands + --json for scripts
Extensible — agents are pluggable adapters; Codex, opencode, Cursor CLI, and kiro-cli are on the roadmap
go install github.com/pradipta/wallfacer@latest
Requires Go 1.22+. Pure Go, no CGO — macOS and Linux. Pre-built binaries are on the
releases page ; building from source is
covered in the development guide .
wallfacer is one binary with two front ends, and which one you get depends on whether you
pass a subcommand:
They are not separate tools and there is nothing to switch between: both read and write the
same SQLite index, so a session you tag in the browser is immediately findable by
wallfacer list --tag , and vice versa. Anything the TUI can do, a subcommand can do too.
(If stdout isn't a terminal — wallfacer | less , or inside a script — the bare command
prints help instead of opening the browser.)
Open the browser and work from there:
wallfacer
Or drive it from the shell:
wallfacer new ~ /work/api --title " Fix flaky auth tests "
wallfacer resume " fix flaky auth tests " # by title or ID prefix
wallfacer search auth
wallfacer list --project api --json # for scripts
The TUI — wallfacer
Bare wallfacer opens a full-screen session browser: a list on the left, and a detail
pane on the right showing everything wallfacer show prints for whatever is highlighted.
Every row shows its project, tags, directory, age and agent — the same metadata
wallfacer list prints. Below ~100 columns the detail pane steps aside and the list
takes the full width.
Every subcommand is one-shot: it runs, prints, and exits. Same index as the browser.
<ref> is an ID prefix ( resume 5f2 ) or an exact title ( resume "smoke test" ) — ambiguous
references list the candidates instead of guessing. Sessions started outside wallfacer are
picked up automatically; there's no import step.
wallfacer scans ~/.claude/projects/ , reading just the head of each session file for its
working directory, timestamps, and first prompt (the automatic title). Your metadata lives in
SQLite at ~/.local/share/wallfacer/ — delete it and you lose only the overlay, never a
conversation. Sync is incremental, so it stays fast with hundreds of sessions.
Other agents plug in through a small adapter interface — see
docs/adding-an-agent.md .
Adapters for more agents: Codex ,
opencode , Cursor CLI, kiro-cli
Full-text search across session content (SQLite FTS5)
wallfacer restore (un-trash from the CLI)
Export a session transcript to Markdown
Stats (sessions per project/week, disk usage)
Desktop app — sessions in a sidebar, embedded terminal, multiple tabs
(Wails + xterm.js on top of the same Go internals)
See the development guide for building, testing, and releasing.
Contributions welcome — especially new agent adapters.
The name is borrowed from the Wallfacers of Liu Cixin's The Dark Forest — people entrusted
with plans too sprawling for anyone else to follow.
A terminal session manager for Claude Code, and more
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

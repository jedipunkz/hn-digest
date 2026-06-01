---
source: "https://github.com/Alejandro-M-P/git-courer"
hn_url: "https://news.ycombinator.com/item?id=48347588"
title: "Git-courer – A complete, JSON-first Git layer for LLM agents"
article_title: "GitHub - Alejandro-M-P/git-courer: MCP server for Git with local Ollama — zero tokens for git operations · GitHub"
author: "blakok14"
captured_at: "2026-06-01T00:27:08Z"
capture_tool: "hn-digest"
hn_id: 48347588
score: 3
comments: 0
posted_at: "2026-05-31T17:27:16Z"
tags:
  - hacker-news
  - translated
---

# Git-courer – A complete, JSON-first Git layer for LLM agents

- HN: [48347588](https://news.ycombinator.com/item?id=48347588)
- Source: [github.com](https://github.com/Alejandro-M-P/git-courer)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T17:27:16Z

## Translation

タイトル: Git-courer – LLM エージェント用の完全な JSON ファースト Git レイヤー
記事のタイトル: GitHub - Alejandro-M-P/git-courer: ローカル Ollama を使用した Git 用 MCP サーバー — Git 操作用のゼロ トークン · GitHub
説明: ローカル Ollama を使用した Git 用 MCP サーバー — git 操作用のゼロトークン - Alejandro-M-P/git-courer

記事本文:
GitHub - Alejandro-M-P/git-courer: ローカル Ollama を使用した Git 用 MCP サーバー — Git 操作用のゼロトークン · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
アレハンドロ-M-P
/
git-courer
公共
通知
通知を変更するにはサインインする必要があります

イオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
412 コミット 412 コミット .git-courer .git-courer .github .github アセット アセット cmd cmd docs docs 内部 内部スクリプト スクリプト ターゲット/リリース ターゲット/リリース テスト テスト tui tui .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md config.example.yaml config.example.yaml go.mod go.mod go.sum go.sumすべてのファイルを表示 リポジトリ ファイルのナビゲーション
問題とバグ : @Alejandro-M-P/git-courer/issues · ディスカッション : @Alejandro-M-P/git-courer/Discussions
ドクター
説明
建築
コードベースの構造、パターン、機能の追加方法
トラブルシューティング
修正: Ollama が実行されていない、MCP が検出されない、権限エラー
MCPクライアント
サポートされている 12 クライアントすべて、構成フォーマット、手動セットアップ
設定オプション
すべての ~/.config/git-courer/config.yaml および .git-courer/config.json 設定
コマンド
22 個の MCP ツールすべての完全なリファレンス
モデルガイド
テストされたモデル、トークンの使用法、およびどれを選択するか
貢献する
セットアップ、テストの実行、および共同作業の方法
git-courer
コミットする前にコードを理解する唯一の MCP git サーバー。
ほとんどのツールは git diff | をラップします。 llm "コミットを書き込む" 。 git-courer は何か違うことをします。つまり、AST を解析し、依存関係グラフを構築し、LLM には触れずに Go で変更タイプを分類し、モデルを呼び出して人間が読めるメッセージを書き込むだけです。
1 つのステージング領域 = 1 つのコミット。いつも。
1. LLM ではなく Go で決定されるコミット タイプ
git-courer は go-enry と go-tree-sitter を使用して AST を解析し、決定論的なルールを適用します — 新しいパブリック関数、mod

署名の修正、シンボルの削除、重大な変更。 LLM は決定ではなくメッセージを書き込みます。
コミットする前に、git-couler はステージングされたファイルがコードベース全体に与える影響のグラフを作成します。どのファイルが操作されたかだけではなく、実際の範囲。
3. ローカルとクラウドの連携
Go とローカル LLM は、できることはすべて処理します。クラウド LLM は、生のコードではなく、構造化されたコミット コンテキスト (WHY + WHAT) を受け取ります。
git-courer コミットは、LLM が直接利用できる圧縮された構造化された概要です。入力が改善され、トークンが減り、幻覚が減ります。
5. すべての突然変異の自動バックアップ
すべての書き込み操作は、実行前にバックアップを作成します。 1 つのコマンドで何でも元に戻せます。
6. 完全な git パック — サーバー 1 台、他には何も必要ありません
なぜコミット用にツールをインストールし、リリース用に別のツールをインストールし、PR チェック用に別のツールをインストールするのでしょうか? git-couler は全サイクルをカバーします。
読み取り: ステータス、差分、履歴、非難、事前レビュー
書き込み: コミット、修正、元に戻す、ステージング、リセット、スタッシュ
ブランチ: ブランチ、マージ、リベース、チェリーピック、タグ
ユーティリティ: 同期、バックアップ、元に戻す、リモート、構成
すべて構造化された JSON を返します。ページャーがハングしたり、テキストを解析したり、追加のツールを使用したりすることはありません。
変更をプレビュー→提案されたコミットを確認→適用します。 git-couler は、ステージングされたファイルを依存関係グラフによって自動的にアトミックコミットに分割します。
pr-review を呼び出す — ワンショットで実行される PR 前ゲート:
テスト — .git-courer/config.json から test_command を実行します (例: go test ./... )
競合 — ターゲットブランチとのマージ競合を検出し、AST アノテーション付きハンクを返します ( [NEW_FUNC] 、 [MOD_SIG ⚠BREAKING] )
差分統計 - ファイルの変更、追加、削除
発散 — 先行/遅延カウント、マージ可能ステータス
緑色でない場合はマージされません。 git-config SET_TEST_COMMAND を介して test_command を設定するか、.git-courer/config.json を直接編集します。
git-couler release を実行します。 CLI 領域

ds は .git-courer/commits.json (このブランチでの各 git-commit APPLY 中にキャプチャされた) からコミットし、それらを .git-courer/config.json で定義された領域ごとにグループ化します。ブランチ ストアが空の場合は、最後のタグ以降の git ログにフォールバックします。 Go はバージョン バンプを計算します。ローカル LLM は、人間が判読できる変更ログをエリアごとに書き込みます。
すべての破壊的な操作には自動バックアップが作成されます。 1 つのコマンドで以前の状態が復元されます。
あなた: 「変更をコミットしてください」
↓
AI は git-courer に委任します (MCP 経由)
↓
Go は AST + 依存関係グラフを読み取ります → 型を決定的に分類します
↓
ローカル LLM は、注釈付きの差分から人間が判読できるメッセージを書き込みます
↓
セキュリティスキャン（5層）→自動バックアップ→コミット
↓
"✓ feat(auth): OAuth2 トークンの更新を追加"
ツールの完全なリストについては: docs/commands.md
ワークフローの詳細: docs/workflows.md
カール -fsSL https://raw.githubusercontent.com/Alejandro-M-P/git-courer/main/scripts/install.sh |しー
それだけです。バイナリをインストールし、マシン上で検出されたすべての AI ツールを自動構成します。
要件: Git · Ollama (オプション、AI コミット メッセージ用)
# macOS / Linux
curl -fsSL https://github.com/Alejandro-M-P/git-courer/releases/latest/download/git-courer_ $( uname -s | tr ' [:upper:] ' ' [: lower:] ' ) _ $( uname -m ) .tar.gz | tar -xz -C /usr/local/bin git-courer
chmod +x /usr/local/bin/git-couler
git-courer のセットアップ
# Windows (PowerShell)
IRM https://github.com/Alejandro-M-P/git-courer/releases/latest/download/git-courer_windows_amd64.tar.gz | tar -xz -o git-courer.exe
。 \g it-couler.exe セットアップ
# または Go を使用して
github.com/Alejandro-M-P/git-courer@latest をインストールしてください
サポートされているツール
ツール
自動構成
クロード・コード
✓
カーソル
✓
ウィンドサーフィン
✓
オープンコード
✓
クライン
✓
Roo コード
✓
VSコード
✓
クロードデスクトップ
✓ macOS/Win のみ
続ける
✓
ゼッド
✓
コーデックス
✓
ジェミニ CLI

✓
git-courer mcp setup を実行して、検出されたすべてのツールを一度に構成するか、特定のツールに対して git-courer mcp setup <client> を実行します。
引数を指定せずに git-couler を実行して、対話型インストーラーを起動します。
TUI では、次の 4 つの手順を実行します。
MCP 構成 - 構成する AI ツールを選択します (インストールされているクライアントを自動検出します)
一般設定 — LLM バックエンド、モデル、プロジェクト コンテキストを構成します
確認 — 保存する前に構成をプレビューします。
完了 — 構成は ~/.config/git-couler/config.yaml に保存されました
バイナリを更新したり、メニューから直接アンインストールしたりすることもできます。ナビゲーション: j/k または矢印キー、esc で戻ります。
git-couler は、引数なしで起動すると対話型 TUI として実行されます。また、MCP サーバーと管理コマンドも提供します。
22 の MCP ツールとその引数については、 docs/commands.md を参照してください。
git-couler は 2 つの構成レベルを使用します。
グローバル ( ~/.config/git-courer/config.yaml ) — 個人設定: LLM バックエンド、モデル。
プロジェクトごと ( .git-courer/config.json ) — コミット可能、チームと共有。説明、領域、test_command、除外されたものを格納します。より良い結果 = このファイルをプロジェクトごとに編集します。
git-commit PREVIEW を呼び出すと、サーバーはすぐに戻る (FAST) か、バックグラウンド ジョブを開始する (SLOW) 場合があります。遅い場合:
PREVIEW は {status:"processing", job_id} を返します
エージェントは、その job_id を使用して git-commit STATUS を呼び出してポーリングする必要があります
STATUS が {status:"done"} を返したら、計画の準備は完了です。
次に、同じ job_id で git-commit APPLY を呼び出します。
エージェントは動作を継続します。ジョブの待機はブロックされません。
git-commit APPLY を介して成功したすべてのコミットは、 .git-courer/commits.json (ブランチ スコープは .git-courer/branches/<branch>/commits.json ) にキャプチャされます。後で git-courer release を実行すると、CLI はキャプチャされたコミットを .git-courer/config.json で定義された領域ごとにグループ化して読み取ります。

変更ログを生成します。 git log を再解析する必要はありません。
これが重要な理由: Git の歴史は頻繁に書き換えられ、PR スカッシュ、リベース、強制プッシュにより、実際のコミットの物語が破壊されます。 CommitStore は、リモートで何が起こったかに関係なく、すべてのコミット メッセージを書き込まれたとおりに保存します。リリースの変更ログは、Git ログが 1 つの圧縮されたコミットにフラット化されても存続します。これは、git 履歴が書き換えられてからも存続するローカル ドキュメントです。
問題がありますか? docs/troubleshooting.md で以下を確認してください。
Ollama が実行されていない / モデルが設定されていません
MCP が AI ツールで検出されない
インストール中の権限エラー
コミットで検出されたシークレット (誤検知)
MCP 構成ファイルの場所: docs/mcp-clients.md
コミットタイプは誰が決めるのでしょうか?
AST 分析による。 LLM は人間が判読できるメッセージのみを書き込みます。
オラマは必要ですか？
LLM バックエンドが必要です。 Ollama が推奨されるデフォルトですが、git-courer は OpenAI 互換サーバー (LM Studio、vLLM、LocalAI、またはカスタム エンドポイント) で動作します。バックエンドが構成されていないと、すべての AI 操作 (コミット、リリース、ブランチ名、セキュリティ監査) が失敗します。基本的な git 読み取り (ステータス、差分、ログ) は引き続き機能します。
私のコードはどこかに送信されますか?
いいえ。git-courer、Ollama、データなど、すべてがあなたのマシン上で実行されます。
リリースのバージョン番号は誰が決定しますか?
オラマじゃなくて行きなさい。バージョンはコミット タイプ (feat: → マイナー、feat!: → メジャー) から計算されます。 Ollama は人間の変更ログのみを書きます。
私のツールがリストにありません。
問題を開きます: @Alejandro-M-P/git-courer/issues 。 MCP をサポートしている場合、通常は数行で追加できます。
重大な変更をマークするにはどうすればよいですか?
使用 ！コミットタイプ ( feat!: ) の後に、本文に BREAKING CHANGE: を含めます。 git-couler は、バージョンバンプと変更ログ生成のためにこれを自動的に取得します。
コラボレーションしたいですか?必要なものは次のとおりです。
ドキュメント/アークを読む

hitecture.md:
ディレクトリ構造と技術スタック
六角形のアーキテクチャパターン
主要なパッケージとその責任
バグが見つかりましたか?問題を開く: @Alejandro-M-P/git-courer/issues
OS と git-couler のバージョン ( git-couler version )
使用しているAIツール（クロードコード、カーソルなど）
関連するログまたはエラー メッセージ
ドキュメントを読みます: docs/architecture.md と CONTRIBUTING.md から始めます。
問題を選択する : @Alejandro-M-P/git-courer/issues で適切な最初の問題のラベルを確認してください
ディスカッション: 質問や機能のアイデアについては、@Alejandro-M-P/git-courer/Discussions を使用してください。
PR を送信 : 従来のコミットに従ってください (feat: 、 fix: 、chore: )
AI ツールが MCP をサポートしているがリストにない場合、通常、internal/installer/mcp_config.go に 5 行のコードを追加します。形式については docs/mcp-clients.md を参照してください。
ローカル Ollama を使用した Git 用の MCP サーバー — git 操作のトークンはゼロ
読み込み中にエラーが発生しました。このページをリロードしてください。
3
フォーク
レポートリポジトリ
リリース
23
v2.0.7 - 公式安定版
最新の
2026 年 5 月 30 日
+ 22 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for Git with local Ollama — zero tokens for git operations - Alejandro-M-P/git-courer

GitHub - Alejandro-M-P/git-courer: MCP server for Git with local Ollama — zero tokens for git operations · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
Alejandro-M-P
/
git-courer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
412 Commits 412 Commits .git-courer .git-courer .github .github assets assets cmd cmd docs docs internal internal scripts scripts target/ release target/ release test test tui tui .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md config.example.yaml config.example.yaml go.mod go.mod go.sum go.sum View all files Repository files navigation
Issues & Bugs : @Alejandro-M-P/git-courer/issues · Discussions : @Alejandro-M-P/git-courer/discussions
Doc
Description
Architecture
Codebase structure, patterns, and how to add features
Troubleshooting
Fix: Ollama not running, MCP not detected, permission errors
MCP Clients
All 12 supported clients, config formats, manual setup
Config Options
All ~/.config/git-courer/config.yaml and .git-courer/config.json settings
Commands
Complete reference for all 22 MCP tools
Models Guide
Tested models, token usage, and which one to pick
Contributing
Setup, running tests, and how to collaborate
git-courer
The only MCP git server that understands your code before it commits.
Most tools wrap git diff | llm "write a commit" . git-courer does something different: it parses your AST, builds a dependency graph, classifies the change type in Go without touching the LLM, and only calls the model to write the human-readable message.
One staged area = one commit. Always.
1. Commit type determined in Go, not by the LLM
git-courer uses go-enry and go-tree-sitter to parse your AST and apply deterministic rules — new public function, modified signature, deleted symbol, breaking change. The LLM writes the message, not the decision.
Before committing, git-courer builds a graph of what your staged files affect across the codebase. Real scope, not just which files were touched.
3. Local + cloud working together
Go and the local LLM handle everything they can. Your cloud LLM receives structured commit context — WHY + WHAT — instead of raw code.
A git-courer commit is a compressed, structured summary any LLM can consume directly. Better input, fewer tokens, fewer hallucinations.
5. Automatic backup on every mutation
Every write operation creates a backup before executing. One command undoes anything.
6. The complete git pack — one server, nothing else needed
Why install one tool for commits, another for releases, another for PR checks? git-courer covers the full cycle:
Read: status, diff, history, blame, pr-review
Write: commit, amend, revert, stage, reset, stash
Branch: branch, merge, rebase, cherry-pick, tag
Utility: sync, backup, undo, remotes, config
All returning structured JSON. No pager hangs, no text parsing, no extra tools.
Preview the change → review proposed commits → apply. git-courer splits your staged files into atomic commits by dependency graph automatically.
Call pr-review — a pre-PR gate that runs in one shot:
Tests — runs test_command from .git-courer/config.json (e.g. go test ./... )
Conflicts — detects merge conflicts with the target branch and returns AST-annotated hunks ( [NEW_FUNC] , [MOD_SIG ⚠BREAKING] )
Diff stats — files changed, additions, deletions
Divergence — ahead/behind count, mergeable status
If it's not green, you don't merge. Set test_command via git-config SET_TEST_COMMAND or edit .git-courer/config.json directly.
Run git-courer release . The CLI reads commits from .git-courer/commits.json (captured during each git-commit APPLY on this branch) and groups them by the areas defined in .git-courer/config.json . If the branch store is empty, it falls back to git log since the last tag. Go calculates the version bump; the local LLM writes the human-readable changelog per area.
Every destructive operation has an automatic backup. One command restores the previous state.
You: "commit my changes"
↓
AI delegates to git-courer (via MCP)
↓
Go reads AST + dependency graph → classifies type deterministically
↓
Local LLM writes the human-readable message from the annotated diff
↓
Security scan (5 layers) → auto-backup → commit
↓
"✓ feat(auth): add OAuth2 token refresh"
For the full list of tools: docs/commands.md
For workflow details: docs/workflows.md
curl -fsSL https://raw.githubusercontent.com/Alejandro-M-P/git-courer/main/scripts/install.sh | sh
That's it. It installs the binary and auto-configures every AI tool it detects on your machine.
Requirements: Git · Ollama (optional, for AI commit messages)
# macOS / Linux
curl -fsSL https://github.com/Alejandro-M-P/git-courer/releases/latest/download/git-courer_ $( uname -s | tr ' [:upper:] ' ' [:lower:] ' ) _ $( uname -m ) .tar.gz | tar -xz -C /usr/local/bin git-courer
chmod +x /usr/local/bin/git-courer
git-courer setup
# Windows (PowerShell)
irm https://github.com/Alejandro-M-P/git-courer/releases/latest/download/git-courer_windows_amd64.tar.gz | tar -xz -o git-courer.exe
. \g it-courer.exe setup
# Or with Go
go install github.com/Alejandro-M-P/git-courer@latest
Supported Tools
Tool
Auto-configured
Claude Code
✓
Cursor
✓
Windsurf
✓
OpenCode
✓
Cline
✓
Roo Code
✓
VS Code
✓
Claude Desktop
✓ macOS/Win only
Continue
✓
Zed
✓
Codex
✓
Gemini CLI
✓
Run git-courer mcp setup to configure all detected tools at once, or git-courer mcp setup <client> for a specific one.
Run git-courer with no arguments to launch the interactive installer:
The TUI walks you through 4 steps:
MCP Configuration — select which AI tools to configure (auto-detects installed clients)
General Settings — configure your LLM backend, model, and project context
Review — preview your config before saving
Finish — config saved to ~/.config/git-courer/config.yaml
You can also update the binary or uninstall directly from the menu. Navigation: j/k or arrow keys, esc to go back.
git-courer runs as an interactive TUI when launched without arguments. It also provides MCP server and management commands:
For the 22 MCP tools and their arguments, see docs/commands.md .
git-courer uses two config levels:
Global ( ~/.config/git-courer/config.yaml ) — personal settings: LLM backend, model.
Per-project ( .git-courer/config.json ) — committable, shared with team. Stores description, areas, test_command, excluded. Better results = edit this file per project.
When you call git-commit PREVIEW , the server may return immediately (FAST) or start a background job (SLOW). In the SLOW case:
PREVIEW returns {status:"processing", job_id}
The agent must call git-commit STATUS with that job_id to poll
When STATUS returns {status:"done"} , the plan is ready
Then call git-commit APPLY with the same job_id
The agent continues working — it does NOT block waiting for the job.
Every successful commit via git-commit APPLY is captured to .git-courer/commits.json (branch-scoped under .git-courer/branches/<branch>/commits.json ). When you later run git-courer release , the CLI reads those captured commits — grouped by the areas defined in .git-courer/config.json — to generate the changelog. No re-parsing git log .
Why this matters: Git history is frequently rewritten — PR squashes, rebases, force-pushes — destroying the real commit narrative. The CommitStore preserves every commit message as it was written, independently of what happens on the remote. Your release changelog survives git log being flattened into a single squashed commit. This is local documentation that outlives git history rewriting.
Having issues? Check docs/troubleshooting.md for:
Ollama not running / model not configured
MCP not detected by your AI tool
Permission errors during install
Secrets detected in commits (false positives)
MCP config file locations: docs/mcp-clients.md
Who decides the commit type?
Go, via AST analysis. The LLM only writes the human-readable message.
Do I need Ollama?
You need some LLM backend. Ollama is the recommended default, but git-courer works with any OpenAI-compatible server: LM Studio, vLLM, LocalAI, or a custom endpoint. Without a configured backend, all AI operations (commits, releases, branch names, security auditor) fail. Basic git reads (status, diff, log) still work.
Is my code sent anywhere?
No. Everything runs on your machine — git-courer, Ollama, your data.
Who decides the version number in a release?
Go, not Ollama. Version is calculated from commit types ( feat: → minor, feat!: → major). Ollama only writes the human changelog.
My tool isn't listed.
Open an issue: @Alejandro-M-P/git-courer/issues . If it supports MCP, adding it is usually a few lines.
How do I mark a breaking change?
Use ! after the commit type ( feat!: ) or include BREAKING CHANGE: in the body. git-courer picks this up automatically for version bumping and changelog generation.
Want to collaborate? Here's everything you need:
Read docs/architecture.md for:
Directory structure and tech stack
Hexagonal architecture patterns
Key packages and their responsibilities
Found a bug? Open an issue: @Alejandro-M-P/git-courer/issues
Your OS and git-courer version ( git-courer version )
AI tool you're using (Claude Code, Cursor, etc.)
Relevant logs or error messages
Read the docs : Start with docs/architecture.md and CONTRIBUTING.md
Pick an issue : Check @Alejandro-M-P/git-courer/issues for good first issue labels
Discuss : Use @Alejandro-M-P/git-courer/discussions for questions or feature ideas
Submit PR : Follow conventional commits ( feat: , fix: , chore: )
If your AI tool supports MCP but isn't listed, adding it is usually 5 lines of code in internal/installer/mcp_config.go . See docs/mcp-clients.md for the format.
MCP server for Git with local Ollama — zero tokens for git operations
There was an error while loading. Please reload this page .
3
forks
Report repository
Releases
23
v2.0.7 - Official Stable
Latest
May 30, 2026
+ 22 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

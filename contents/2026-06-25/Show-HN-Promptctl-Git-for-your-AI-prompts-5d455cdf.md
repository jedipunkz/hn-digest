---
source: "https://github.com/naya-ai/promptctl"
hn_url: "https://news.ycombinator.com/item?id=48667544"
title: "Show HN: Promptctl – Git for your AI prompts"
article_title: "GitHub - naya-ai/promptctl: Git for your AI prompts — version, diff, and rollback LLM prompts from the CLI · GitHub"
author: "shawnaya101"
captured_at: "2026-06-25T01:40:20Z"
capture_tool: "hn-digest"
hn_id: 48667544
score: 2
comments: 0
posted_at: "2026-06-25T01:09:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Promptctl – Git for your AI prompts

- HN: [48667544](https://news.ycombinator.com/item?id=48667544)
- Source: [github.com](https://github.com/naya-ai/promptctl)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T01:09:05Z

## Translation

タイトル: HN を表示: Promptctl – AI プロンプト用の Git
記事のタイトル: GitHub - naya-ai/promptctl: AI プロンプト用の Git — CLI からのバージョン、差分、ロールバック LLM プロンプト · GitHub
説明: AI プロンプト用の Git — CLI からのバージョン、差分、ロールバック LLM プロンプト - naya-ai/promptctl

記事本文:
GitHub - naya-ai/promptctl: AI プロンプト用の Git — CLI からのバージョン、差分、ロールバック LLM プロンプト · GitHub
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
ナヤアイ
/
プロンプトctl
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows .vscode .vscode cmd/promptctl cmd/promptctl docs docs 内部内部スクリプト scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod Viewすべてのファイル リポジトリ ファイルのナビゲーション
コードをバージョン管理するのと同じように、LLM プロンプトを追跡、バージョン管理、差分計算、ロールバックします。
$ プロンプトctl commit システム --file プロンプト/system.txt -m "初期"
✓ プロンプト「system」を v1 としてコミット
$プロンプトctl commit システム --file プロンプト/system.txt -m "引用ルールを追加しました"
✓ プロンプト「system」を v2 としてコミット
$promptctl 差分システム
--- システム v1 (2026-06-01 09:12)
+++ システム v2 (2026-06-03 14:47)
- あなたは有能なアシスタントです。
+ あなたは親切なアシスタントです。常に出典を引用してください。
$promptctl rollback system 1 -m "元に戻す — 引用は思い出しを損なう"
✓ 「システム」を v1 にロールバック → v3 として保存
問題
プロンプトはランダムな場所に存在します。ファイルにハードコーディングされたり、Notion ドキュメントに埋め込まれたり、Slack にコピー＆ペーストされたりします。何かが壊れたり、正常に動作しなくなったりした場合、次のようなことはわかりません。
先週プロンプトが機能していたときの様子
特定のポイントに戻る方法
これを解決するのがpromptctlです。プロンプト用の git です。
github.com/naya-ai/promptctl/cmd/promptctl@latest をインストールしてください
またはソースからビルドします。
git クローン https://github.com/naya-ai/promptctl.git
cd プロンプトctl
bin/promptctl でビルド # バイナリを作成する
# または: go build -o bin/promptctl ./cmd/promptctl/
デモ
# 1. プロジェクトで初期化します (.promptctl/store.json を作成します)
プロンプトctlの初期化
# 2. 最初のプロンプト バージョンを保存する
echo "あなたは親切なアシスタントです。" | echoプロンプトctl commit system -m " 初期 "
# 3. 変更を加えて保存する
エコー「あなたは役に立ちます」

助手。常に出典を引用します。 " |promptctl commit system -m " 引用ルールを追加しました " --model claude-3
#4. 歴史を見る
プロンプトctl ログ システム --preview
#5. 最新と以前を比較する
プロンプトctl diffシステム
#6. ロールバック
プロンプトctl ロールバック システム 1 -m " 元に戻す "
コマンド
現在のディレクトリでプロンプトctlを初期化します。 .promptctl/store.json を作成します。
新しいバージョンのプロンプトを保存します。
# 標準入力から
echo "あなたは親切なアシスタントです。 " | プロンプトctl commit system -m " 初期 "
# ファイルから (長いプロンプトの場合に推奨)
プロンプトctl commit system --file プロンプト/system.txt -m " ファイルから "
# メタデータあり
Promptctl commit classifier -m " 速度を最適化 " --model gpt-4o-mini --tag prod
# インタラクティブ — 終了するには、単独の行に「END」と入力します。
プロンプトctl commit system -m "手動入力"
フラグ:
-m / --message — コミットメッセージ (デフォルト: "add version N" )
--model — このプロンプトが対象とするモデルのタグ
--tag — このバージョンにラベルを追加します
-f / --file — 標準入力ではなくファイルからコンテンツを読み取ります
プロンプトctl ログ # すべてのプロンプト
プロンプトctl ログ システム # 1 つのプロンプト
プロンプトctl ログ システム --preview # コンテンツ スニペットを表示
Promptctl log system --tag prod # タグによるフィルタリング
プロンプトctl ログ システム --json # 機械可読出力
差分
色で強調表示してバージョンを並べて比較します。
Promptctl diff システム # 最新と以前
Promptctl diff system 2 # v2 と最新
プロンプトctl diff システム 1 3 # v1 と v3
ショー
プロンプトctl システム # 最新バージョンを表示
プロンプトctl show system 2 # 特定のバージョン
Promptctl show system --raw # コンテンツのみ (パイプ処理に適しています)
プロンプトctl show system --json # 完全な JSON
プロンプトctl show system --copy # クリップボードにコピー (macOS/Linux)
Promptctl show system --version-at 2026-06-01 # その日付の最新バージョン
ロールバック
以前のバージョンを復元します。新しいものを作成します

バージョン（非破壊）。
プロンプトctlロールバックシステム1
Promptctl rollback system 1 -m "元に戻す — 引用を傷つけるリコール"
リスト
プロンプトctlリスト
プロンプトctl list --json # バージョン数を含む JSON 配列
Promptctl list --names # 1 行に 1 つの名前 (シェル補完で使用)
検索
コンテンツ、コミットメッセージ、タグを横断して検索します。
プロンプトctl検索「役立つアシスタント」
Promptctl search prod # タグ「prod」に一致します
輸出
プロンプトのすべてのバージョンを Markdown にエクスポートします。
プロンプトctlエクスポートシステム
プロンプト ctl エクスポート システム > system-history.md
コピー
完全なバージョン履歴を含むプロンプトを新しい名前でフォークします。
プロンプトctlコピーシステムシステム実験的
名前を変更する
プロンプトの名前を変更します。すべての歴史はそれとともに動きます。
プロンプトctlシステムアシスタントの名前変更
削除する
プロンプトとそのすべてのバージョンを削除します。
ディスク上でファイルが変更されるたびにファイルを自動コミットします。通常のエディタでプロンプトを編集し、保存するたびに自動的にバージョン管理するのに最適です。
プロンプトctl 監視プロンプト/system.txt --as システム
プロンプトctl 監視プロンプト/system.txt --as system --model claude-3 --interval 2
フラグ:
--as <name> — コミットするプロンプト名 (必須)
--model — 新しいバージョンにこのモデルをタグ付けします
--interval <秒> — ポーリング間隔 (デフォルト: 1)
コミットのタイムスタンプは、promptctl の実行時ではなく、ファイルの変更時刻と一致します。
N 個の最新バージョンのみを保持し、残りは破棄します。
プロンプトctl プルーン システム --keep 10
統計
プロンプトctl統計
統計:
─────────────────
プロンプト: 4
合計バージョン: 23
総コンテンツ: 8432 文字
最初のコミット: 2026-06-01 09:12:00
最終コミット: 2026-06-24 15:30:44
最も反復された: システム (12 バージョン)
完成
シェル完了スクリプトを出力します。
# Bash — ~/.bashrc に追加
ソース <( プロンプトctl comp

レションバッシュ）
# Zsh — ~/.zshrc に追加
ソース <( プロンプトctl 完了 zsh )
# Fish — 完了ディレクトリに保存
プロンプトctl 完了フィッシュ > ~ /.config/fish/completions/promptctl.fish
補完は、<name> 引数を取るコマンドのプロンプト名を動的に提案します。
Promptctl は、.git/ の仕組みと同様に、すべてを .promptctl/store.json に保存します。ストアは現在のディレクトリまたは親から検出されるため、コマンドはどのサブディレクトリからでも機能します。
あなたのプロジェクト/
§── .promptctl/
│ └── store.json ← すべてのプロンプト バージョンがここにあります
§── src/
━──
プロンプト履歴をローカルに保持するには、プロジェクトの .gitignore に .promptctl/ を追加します。チームが git でバージョンを共有したい場合は、ストア ファイルのみを追跡します。
.promptctl / *
！ .promptctl / ストア.json
プログラムによる読み取りと書き込み
いくつかのコマンドは、スクリプト作成のために --json または --raw をサポートしています。
# プロンプトの最新バージョンを読む
プロンプトctl show system --raw
# 最新のコンテンツを別のツールにパイプします
プロンプトctl show system --raw | pbcopy
# プロンプトを JSON としてリストします
プロンプトctlリスト --json
# スクリプトから出力を取り込む
./generate-prompt.sh |プロンプトctl commit system -m " 生成されました "
タグ付けとモデルの追跡
# フィルタリングするバージョンにタグを付ける
プロンプトctl commit system -m "production" --tag prod --model claude-opus-4-8
# タグ付きバージョンのみを表示
プロンプトctl ログ システム --tag prod
# タグを参照しているすべてのプロンプトを検索します
プロンプトctl検索製品
なぜ依存関係をゼロにするのか?
Promptctl は Go の標準ライブラリのみを使用します。外部パッケージ、バージョンの競合、Go Mod のようなきちんとした驚きはありません。バイナリは 1 つで、ランタイムはありません。
端末を並べて見た差分図
COTRIBUTING.md を参照してください。 PR は歓迎です — 大きな変更については、まず問題をオープンしてください。
git クローン https://github.com/naya-ai/promptctl.git
cd プロンプトctl
テストを行う
ビルドする
ライセンス
AI プロンプト用の Git —

CLI からのバージョン、差分、ロールバック LLM プロンプト
github.com/naya-ai/promptctl
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0 — 初期リリース
最新の
2026 年 6 月 25 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Git for your AI prompts — version, diff, and rollback LLM prompts from the CLI - naya-ai/promptctl

GitHub - naya-ai/promptctl: Git for your AI prompts — version, diff, and rollback LLM prompts from the CLI · GitHub
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
naya-ai
/
promptctl
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows .vscode .vscode cmd/ promptctl cmd/ promptctl docs docs internal internal scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod View all files Repository files navigation
Track, version, diff, and rollback your LLM prompts — just like you version your code.
$ promptctl commit system --file prompts/system.txt -m "initial"
✓ Committed prompt "system" as v1
$ promptctl commit system --file prompts/system.txt -m "added citation rule"
✓ Committed prompt "system" as v2
$ promptctl diff system
--- system v1 (2026-06-01 09:12)
+++ system v2 (2026-06-03 14:47)
- You are a helpful assistant.
+ You are a helpful assistant. Always cite your sources.
$ promptctl rollback system 1 -m "reverting — citation hurt recall"
✓ Rolled back "system" to v1 → saved as v3
The Problem
Your prompts live in random places — hardcoded in files, buried in Notion docs, copy-pasted in Slack. When something breaks or stops working well, you have no idea:
What the prompt looked like last week when it was working
How to get back to a specific point
promptctl solves this. It's git for prompts.
go install github.com/naya-ai/promptctl/cmd/promptctl@latest
Or build from source:
git clone https://github.com/naya-ai/promptctl.git
cd promptctl
make build # binary at bin/promptctl
# or: go build -o bin/promptctl ./cmd/promptctl/
Demo
# 1. Initialize in your project (creates .promptctl/store.json)
promptctl init
# 2. Save your first prompt version
echo " You are a helpful assistant. " | promptctl commit system -m " initial "
# 3. Make a change and save it
echo " You are a helpful assistant. Always cite sources. " | promptctl commit system -m " added citation rule " --model claude-3
# 4. See the history
promptctl log system --preview
# 5. Compare latest vs. previous
promptctl diff system
# 6. Roll back
promptctl rollback system 1 -m " reverting "
Commands
Initialize promptctl in the current directory. Creates .promptctl/store.json .
Save a new version of a prompt.
# From stdin
echo " You are a helpful assistant. " | promptctl commit system -m " initial "
# From a file (recommended for longer prompts)
promptctl commit system --file prompts/system.txt -m " from file "
# With metadata
promptctl commit classifier -m " optimized for speed " --model gpt-4o-mini --tag prod
# Interactive — type END on its own line to finish
promptctl commit system -m " manual entry "
Flags:
-m / --message — commit message (default: "add version N" )
--model — tag with the model this prompt targets
--tag — add a label to this version
-f / --file — read content from a file instead of stdin
promptctl log # all prompts
promptctl log system # one prompt
promptctl log system --preview # show content snippets
promptctl log system --tag prod # filter by tag
promptctl log system --json # machine-readable output
diff
Compare versions side by side with color highlighting.
promptctl diff system # latest vs. previous
promptctl diff system 2 # v2 vs. latest
promptctl diff system 1 3 # v1 vs. v3
show
promptctl show system # latest version
promptctl show system 2 # specific version
promptctl show system --raw # content only (good for piping)
promptctl show system --json # full JSON
promptctl show system --copy # copy to clipboard (macOS/Linux)
promptctl show system --version-at 2026-06-01 # version that was latest on that date
rollback
Restore a previous version. Creates a new version (non-destructive).
promptctl rollback system 1
promptctl rollback system 1 -m " reverting — citation hurt recall "
list
promptctl list
promptctl list --json # JSON array with version counts
promptctl list --names # one name per line (used by shell completions)
search
Search across content, commit messages, and tags.
promptctl search " helpful assistant "
promptctl search prod # matches tag "prod"
export
Export all versions of a prompt to Markdown.
promptctl export system
promptctl export system > system-history.md
copy
Fork a prompt with its full version history under a new name.
promptctl copy system system-experimental
rename
Rename a prompt. All history moves with it.
promptctl rename system assistant
delete
Delete a prompt and all its versions.
Auto-commit a file whenever it changes on disk. Great for editing prompts in your usual editor and automatically versioning every save.
promptctl watch prompts/system.txt --as system
promptctl watch prompts/system.txt --as system --model claude-3 --interval 2
Flags:
--as <name> — the prompt name to commit under (required)
--model — tag new versions with this model
--interval <seconds> — polling interval (default: 1)
Commit timestamps match the file's modification time, not when promptctl ran.
Keep only the N most recent versions and discard the rest.
promptctl prune system --keep 10
stats
promptctl stats
Stats:
──────────────────────────────────────────────────
Prompts: 4
Total versions: 23
Total content: 8432 chars
First commit: 2026-06-01 09:12:00
Last commit: 2026-06-24 15:30:44
Most iterated: system (12 versions)
completion
Print a shell completion script.
# Bash — add to ~/.bashrc
source <( promptctl completion bash )
# Zsh — add to ~/.zshrc
source <( promptctl completion zsh )
# Fish — save to completions directory
promptctl completion fish > ~ /.config/fish/completions/promptctl.fish
Completions dynamically suggest prompt names for commands that take a <name> argument.
promptctl stores everything in .promptctl/store.json — similar to how .git/ works. The store is discovered from your current directory or any parent, so commands work from any subdirectory.
your-project/
├── .promptctl/
│ └── store.json ← all prompt versions live here
├── src/
└── ...
Add .promptctl/ to your project's .gitignore to keep prompt history local. If your team wants to share versions in git, track only the store file:
.promptctl / *
! .promptctl / store.json
Reading & Writing Programmatically
Several commands support --json or --raw for scripting:
# Read the latest version of a prompt
promptctl show system --raw
# Pipe the latest content into another tool
promptctl show system --raw | pbcopy
# List prompts as JSON
promptctl list --json
# Ingest output from a script
./generate-prompt.sh | promptctl commit system -m " generated "
Tagging & Model Tracking
# Tag a version for filtering
promptctl commit system -m " production " --tag prod --model claude-opus-4-8
# View only tagged versions
promptctl log system --tag prod
# Find all prompts referencing a tag
promptctl search prod
Why Zero Dependencies?
promptctl uses only Go's standard library — no external packages, no version conflicts, no go mod tidy surprises. One binary, no runtime.
Side-by-side terminal diff view
See CONTRIBUTING.md . PRs welcome — open an issue first for major changes.
git clone https://github.com/naya-ai/promptctl.git
cd promptctl
make test
make build
License
Git for your AI prompts — version, diff, and rollback LLM prompts from the CLI
github.com/naya-ai/promptctl
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0 — Initial release
Latest
Jun 25, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

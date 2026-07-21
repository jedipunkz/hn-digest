---
source: "https://github.com/RimantasZ/ctoken"
hn_url: "https://news.ycombinator.com/item?id=48992877"
title: "Show HN: Ctoken – CLI util to count LLM tokens in files, dirs or input"
article_title: "GitHub - RimantasZ/ctoken: Simple CLI utility to estimate size of directory or file in tokens · GitHub"
author: "iezhy"
captured_at: "2026-07-21T14:56:23Z"
capture_tool: "hn-digest"
hn_id: 48992877
score: 2
comments: 0
posted_at: "2026-07-21T14:32:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ctoken – CLI util to count LLM tokens in files, dirs or input

- HN: [48992877](https://news.ycombinator.com/item?id=48992877)
- Source: [github.com](https://github.com/RimantasZ/ctoken)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T14:32:23Z

## Translation

タイトル: Show HN: Ctoken – ファイル、ディレクトリ、または入力内の LLM トークンをカウントする CLI ユーティリティ
記事のタイトル: GitHub - RimantasZ/ctoken: トークン内のディレクトリまたはファイルのサイズを推定するシンプルな CLI ユーティリティ · GitHub
説明: トークン内のディレクトリまたはファイルのサイズを推定するシンプルな CLI ユーティリティ - RimantasZ/ctoken
HN テキスト: AI スキル、ナレッジ ベース、MCP の説明などを実装する場合に役立ちます。Web 計算機よりも CLI ツールを使用する方がはるかに簡単です (特に複数のファイルの場合)

記事本文:
GitHub - RimantasZ/ctoken: トークン内のディレクトリまたはファイルのサイズを推定するシンプルな CLI ユーティリティ · GitHub
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
リマンタスZ
/
Cトークン
公共
通知
通知設定を変更するにはサインインする必要があります
追加

イオンナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github/ workflows .github/ workflows src src testing testing .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ctoken は、ファイルまたはディレクトリとその内容内のトークンをカウントする cli ユーティリティです。行コードで cloc を使用する方法と似ています。ファイルまたはディレクトリをコーディング エージェントに供給するときに、そのファイルまたはディレクトリがどの程度のコンテキストを消費するかを理解するのに役立ちます。
一般に AI エージェントまたは LLM ベースのアプリを開発する場合、特定のデータがコンテキスト ウィンドウにどの程度影響を与えるかを知ることが興味深い場合があります。つまり、特定のファイルまたはファイルのセットがどのくらいのトークンに変換されるかということです。
サイズや単語数で見積もったり、Web 上のさまざまな計算機を使用したり、LLM プロバイダー API の 1 つを呼び出したりするなど、さまざまなオプションがありますが、これを繰り返し実行する必要がある場合や、より大きなファイルのセットに対して実行する必要がある場合には、どれも不便になります。
そこで ctoken ユーティリティの出番です。「ctoken <dirname>」と入力すると、その内容のトークン数の概要が表示されます。
ディレクトリトークン
-----------------
。 26,091
.github 1,462
式152
ソース 12,959
テスト 3,092
-----------------
合計 43,756
また、ファイル タイプによるグループ化、パターンやカスタマイズ可能なプロファイルによるフィルタリングなどもサポートしています。詳細については、「フラグ」を参照してください。
醸造タップ RimantasZ/ctoken
醸造インストールctoken
Linux (Debian / Ubuntu)
最新リリースから .deb をダウンロードします。
カール -LO https://github.com/RimantasZ/ctoken/releases/latest/download/ctoken_amd64.deb
sudo apt install ./ctoken_amd64.deb
窓
最新リリースから ctoken-x86_64-windows.zip をダウンロードして解凍し、フォルダーを PATH に追加します。
カーゴインストール --git https://github.com/RimantasZ/ctoken
またはクローンを作成してビルドする

ローカルで:
git clone https://github.com/RimantasZ/ctoken
cdcトークン
カーゴビルド --release
# ターゲット/リリース/ctoken のバイナリ
簡単な例
# デフォルト: 直接のサブディレクトリごとのトークン数
cトークン 。
# ファイル拡張子ごとにグループ化する
cトークン 。 -t
# 組み込みの言語プロファイルを使用する
cトークン 。 -p 錆び
# マークダウン ファイルのみに一致
cトークン 。 -m ' **/*.md '
# ディレクトリごとの内訳で再帰的に歩きます
cトークン 。 --再帰的
# 合計トークン数のみ
cトークン 。 -s
# JSON出力
cトークン 。 --json
# 単一ファイル内のトークンをカウントする
ctoken src/main.rs
# 処理中の各ファイルを表示する
cトークン 。 -v
標準入力とパイプの使用法
ctoken は、パスが指定されていない場合でも stdin から読み取ることができるため、パイプラインやアドホック入力で簡単に使用できます。
# ctoken を介してファイルをパイプします
猫のmyfile.txt | Cトークン
# パイプラインのステップとして使用します — 出力は裸の整数です
猫のmyfile.txt | Cトークン | xargs -I{} echo " トークン数: {} "
# コマンドの出力からトークンをカウントします (例: git diff)
git diff HEAD~1 | Cトークン
# 他のフラグと組み合わせる
猫のmyfile.txt | ctoken --json
猫のmyfile.txt | ctoken --encoding cl100k_base
# stdin を明示的に読み取るには「-」を使用します (他のフラグと組み合わせる場合に便利です)
ctoken - --json < myfile.txt
# 対話型モード: 引数なしで実行し、テキストを入力または貼り付け、完了したら Ctrl+D を押します。
Cトークン
stdin が端末 (対話モード) の場合、ctoken は短いプロンプトを stderr に出力し、入力を待ちます。 Ctrl+D で入力終了を通知すると、トークン数が標準出力に出力されます。他のすべてのフラグ ( --json 、 --encoding 、 --verbose ) は、stdin モードで同じように機能します。
ショート
ロング
引数
説明
-h
--助けて
—
ヘルプを印刷して終了する
--バージョン
—
バージョンを印刷して終了する
-t
--type
—
サブディレクトリではなくファイル拡張子でグループ化する
-g
--gitignore
に |オフ
.gitignore を尊重します。デフォルトでオン
-m
--一致
<グロブ>
インクルードされるファイルを制限するグロブ パターン。代表者

食べられる
-p
--プロフィール
<名前>
~/.config/ctoken/profiles.toml の名前付きプロファイルを使用する
--recreate-プロファイル
—
profile.toml の組み込みプロファイル エントリを書き換えます (対話型)
--再帰的
—
再帰的に歩きます。ファイルタイプごとにグループ化されたディレクトリごとのテーブル
--recursive-with-dir
—
--recursive と同じですが、子ディレクトリのロールアップが含まれます
-v
--冗長
—
処理された各ファイルをログに記録します
-s
--合計
—
総計 (単一の整数) のみを出力します。
--json
—
テーブルの代わりに JSON を出力します。 --recursive* と互換性がありません
--エンコーディング
<名前>
Tiktoken エンコーディング (以下を参照)
エンコードオプション
ctoken は tiktoken-rs を使用してファイル内の実際のトークンを推定します。
OpenAI モデルで使用されるこれらのエンコーディングをサポートします
注: LLM プロバイダーが異なると、トークンの計算が大幅に異なる場合があります。したがって、このツールは、正確な見積もりではなく、大まかな比較 (例: 「このファイル/フォルダーがあのファイル/フォルダーよりもトークンの観点からどのくらい大きいか」) に使用する必要があります。
最初の実行時に、ctoken は、一般的なプロジェクト タイプ ( java 、 c_cpp 、 typescript 、 python 、rust 、 go ) の組み込みプロファイルを含む ~/.config/ctoken/profiles.toml を作成します。
# プロファイルを使用する
cトークン 。 -p タイプスクリプト
# 組み込みプロファイルをデフォルトに復元します (変更する前にプロンプトが表示されます)
cトークン 。 --recreate-プロファイル
~/.config/ctoken/profiles.toml を直接編集して、カスタム プロファイルを追加するか、既存のプロファイルを微調整します。
[私のプロジェクト]
include = [ " **/*.rs " 、 " **/*.toml " 、 " docs/**/*.md " ]
exclude = [ "ターゲット/** " ]
新しいバージョンで追加された新しい組み込みプロファイルは、カスタマイズを上書きせずに自動的に追加されます。
すべての CPU コアをトークン化に使用します (レーヨン経由)。
ファイルは完全にメモリに読み込まれます。非常に大きなファイル (50 MB 以上) は、相応の RAM を使用します。
バイナリ ファイルは、拡張子によって、または最初の 8 KB の NUL バイトをスキャンすることによって検出され、スキップされます。
ディレクトリのサイズを見積もるためのシンプルな CLI ユーティリティ

またはトークンにファイルする
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
cトークン v0.3.0
最新の
2026 年 7 月 1 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Simple CLI utility to estimate size of directory or file in tokens - RimantasZ/ctoken

Useful when implementing AI skills, knowledge bases, MCP descriptions etc - its much easier to use CLI tool than web calculator (especially for multiple files)

GitHub - RimantasZ/ctoken: Simple CLI utility to estimate size of directory or file in tokens · GitHub
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
RimantasZ
/
ctoken
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github/ workflows .github/ workflows src src tests tests .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
ctoken is a cli utility to count tokens in a file or a directory and its contents — similar to how cloc is used for lines code. Useful for understanding how much context a file or directory would consume when feeding it to a coding agent.
When developing AI agends or LLM based apps in general, sometimes it is interesting to know how much a certain data will impact context window - that is how much tokens a certain file or set of files will translate to.
There are various options - estimate it by size or word count, use various calculators on the web, or call one of LLM providers APIs - but all of them become inconvenient when this needs to be done repeadedly or on larger set of files.
Thats where ctoken utility comes in - type ctoken <dirname> and you get token count summary of its contents:
DIRECTORY TOKENS
-----------------
. 26,091
.github 1,462
Formula 152
src 12,959
tests 3,092
-----------------
TOTAL 43,756
It also supports grouping by file type, filtering by pattern or customisable profiles, etc - see Flags for details.
brew tap RimantasZ/ctoken
brew install ctoken
Linux (Debian / Ubuntu)
Download the .deb from the latest release :
curl -LO https://github.com/RimantasZ/ctoken/releases/latest/download/ctoken_amd64.deb
sudo apt install ./ctoken_amd64.deb
Windows
Download ctoken-x86_64-windows.zip from the latest release , extract it, and add the folder to your PATH .
cargo install --git https://github.com/RimantasZ/ctoken
Or clone and build locally:
git clone https://github.com/RimantasZ/ctoken
cd ctoken
cargo build --release
# binary at target/release/ctoken
Quick examples
# Default: token count by immediate subdirectory
ctoken .
# Group by file extension
ctoken . -t
# Use a built-in language profile
ctoken . -p rust
# Match only markdown files
ctoken . -m ' **/*.md '
# Walk recursively, per-directory breakdown
ctoken . --recursive
# Just the total token count
ctoken . -s
# JSON output
ctoken . --json
# Count tokens in a single file
ctoken src/main.rs
# Show each file as it's processed
ctoken . -v
Stdin and pipe usage
ctoken can read from stdin when no path is given, making it easy to use in pipelines or with ad-hoc input.
# Pipe a file through ctoken
cat myfile.txt | ctoken
# Use as a step in a pipeline — output is a bare integer
cat myfile.txt | ctoken | xargs -I{} echo " Token count: {} "
# Count tokens from a command's output (e.g. git diff)
git diff HEAD~1 | ctoken
# Combine with other flags
cat myfile.txt | ctoken --json
cat myfile.txt | ctoken --encoding cl100k_base
# Use '-' to read stdin explicitly (useful when mixing with other flags)
ctoken - --json < myfile.txt
# Interactive mode: run with no arguments, type or paste text, press Ctrl+D when done
ctoken
When stdin is a terminal (interactive mode), ctoken prints a brief prompt to stderr and waits for input. The token count is printed to stdout once you signal end-of-input with Ctrl+D. All other flags ( --json , --encoding , --verbose ) work the same way in stdin mode.
Short
Long
Arg
Description
-h
--help
—
Print help and exit
--version
—
Print version and exit
-t
--type
—
Group by file extension instead of by subdirectory
-g
--gitignore
on | off
Honor .gitignore . Default on
-m
--match
<GLOB>
Glob pattern restricting included files. Repeatable
-p
--profile
<NAME>
Use named profile from ~/.config/ctoken/profiles.toml
--recreate-profiles
—
Rewrite built-in profile entries in profiles.toml (interactive)
--recursive
—
Walk recursively; per-directory table grouped by file type
--recursive-with-dir
—
Same as --recursive , but includes child directory rollups
-v
--verbose
—
Log each file processed
-s
--sum
—
Print only the grand total (single integer)
--json
—
Emit JSON instead of a table. Incompatible with --recursive*
--encoding
<NAME>
Tiktoken encoding (see below)
Encoding options
ctoken uses tiktoken-rs to estimate actual tokens in files,
and supports these encoding used by OpenAI models
Note: for different LLM providers, token calculation might skughtly differ. Therefore this tool should be used for rough comparison (e.g. "how much this file/folder is bigger in terms of tokens than that one"), rather than precise estimation.
On first run, ctoken creates ~/.config/ctoken/profiles.toml with built-in profiles for common project types: java , c_cpp , typescript , python , rust , go .
# Use a profile
ctoken . -p typescript
# Restore built-in profiles to defaults (prompts before changing)
ctoken . --recreate-profiles
Edit ~/.config/ctoken/profiles.toml directly to add custom profiles or tweak existing ones:
[ myproject ]
include = [ " **/*.rs " , " **/*.toml " , " docs/**/*.md " ]
exclude = [ " target/** " ]
New built-in profiles added in later versions are appended automatically without overwriting your customizations.
Uses all CPU cores for tokenization (via rayon).
Files are read fully into memory. Very large files (50+ MB) will use proportionate RAM.
Binary files are detected by extension or by scanning the first 8 KB for NUL bytes, and skipped.
Simple CLI utility to estimate size of directory or file in tokens
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
ctoken v0.3.0
Latest
Jul 1, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/gi-dellav/speck/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48788707"
title: "Speck – AI spec-driven agents, inspired by compilers and build tools"
article_title: "GitHub - gi-dellav/speck: A fully spec-based AI agentic compiler · GitHub"
author: "gidellav"
captured_at: "2026-07-04T20:44:00Z"
capture_tool: "hn-digest"
hn_id: 48788707
score: 1
comments: 0
posted_at: "2026-07-04T20:27:03Z"
tags:
  - hacker-news
  - translated
---

# Speck – AI spec-driven agents, inspired by compilers and build tools

- HN: [48788707](https://news.ycombinator.com/item?id=48788707)
- Source: [github.com](https://github.com/gi-dellav/speck/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T20:27:03Z

## Translation

タイトル: Speck – コンパイラーとビルドツールからインスピレーションを得た、AI 仕様主導型エージェント
記事のタイトル: GitHub - gi-dellav/speck: 完全な仕様ベースの AI エージェント コンパイラー · GitHub
説明: 完全仕様ベースの AI エージェント コンパイラー。 GitHub でアカウントを作成して、gi-dellav/speck の開発に貢献してください。

記事本文:
GitHub - gi-dellav/speck: 完全な仕様ベースの AI エージェント コンパイラー · GitHub
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
ギデラヴ
/
斑点
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 操作

その他のアクション メニュー フォルダーとファイル
36 コミット 36 コミット .github .github data データ パッケージング/自作パッケージング/自作スクリプト スクリプト src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE ライセンス README.md README.md install.sh install.sh justfile justfile Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
仕様主導のエージェント コンパイラ: 完全なエージェント管理のためのシンプルなコマンドにより、再現可能で検証可能なコード生成が可能になり、従来のチャット インターフェイスがスポットライトの当たらない場所に移動します。
注: speck はゼロスタックに基づいており、 PATH 上で動作するゼロスタック バイナリが必要です。
3 層アーキテクチャ: 仕様/機能/ (何を) → 仕様/技術/ (方法) → ソース/ (コード)。任意のレイヤーを編集しても、他のレイヤーは同期されたままになります。
双方向同期 : speck apply は、BLAKE3 ハッシュを介して変更を検出し、特殊な AI プロンプトを使用してレイヤー全体に変更を伝播します。
既存のプロジェクトからの移行: speck は、コードベースから仕様/技術/をリバース エンジニアリングして移行し、そこから仕様/機能/を導き出します。
スキャフォールディング : speck init は、正規の AGENTS.md、ARCHITECTURE.md、およびプロンプト テンプレートを使用して新しいプロジェクトをブートストラップします。
競合の解決: コードと仕様の両方が編集されると、speck は優先順位を求めるプロンプトを表示します (または --prefer-code / --prefer-specs を使用します)。
コード レビュー : スペック レビューは、特殊なプロンプトを介して包括的で構造化されたレビューを実行し、Markdown レポートを出力します。
Git フック: speck git-hooks は、 speck apply 、 speck fmt 、または speck status を自動的に実行する、コミット前、プッシュ前、マージ後、およびチェックアウト後のフックをインストールします。
技術スタック切り替え: speck switch-lang は言語/フレームワークを破壊的に変更し、機能仕様からすべてを再構築します。
ファイル管理: speck mv および speck rm の移動または削除

ハッシュ データベースの同期を維持しながらファイルを保存します。
変更検出: スペックステータスは、編集済みファイルと未登録ファイルをレイヤーごとにグループ化してリストします。
コンパイルされたバイナリは Github リリースからダウンロードできます。
カーゴインストール speck-dev
ソースから
git clone https://github.com/gi-dellav/speck.git
CDスペック
カーゴインストール --path 。
クイックスタート
# ゼロスタックがインストールされ、設定されていることを確認します
ゼロスタック --バージョン
# 新しいスペックプロジェクトの足場を作る
スペック初期化
# または既存のコードベースを移行します
斑点の移行
# 仕様とコードの間で編集を同期する
斑点を適用する
# 何が変わったかを確認する
スペックステータス
# 完全なコードレビュー
スペックレビュー --output report.md
構成
Speck はプロジェクト ルートから Speck.toml を読み取ります。
Speck は、各層が同じ製品の異なる表現である 3 層アーキテクチャを維持しています。
speck apply は 4 ステップのパイプラインを実行します。
変更の検出 — 現在の BLAKE3 ハッシュを .speck_hash.toml と比較します
feat2tech — 機能仕様の編集を技術仕様まで反映します
tech2code — 技術仕様の編集をソースコードまで反映します
code2tech — ソースコード編集を技術仕様まで反映します
各ステップでは、適切な温度で特殊なゼロスタック プロンプトが使用されます (確定的な更新の場合は 0、創造的な仕様の拡張の場合はより高い温度)。
前回の同期以降、ソース ファイルとそれに対応する技術仕様の両方が編集された場合、spec はどちらが信頼できるかを認識できません。インタラクティブに次のようなプロンプトが表示されます。
src/auth.rs の仕様とコードの両方が変更されました。どちらが優先されますか?
[1] コード (一致するように仕様を更新)
[2] 仕様 (一致するコードを更新)
--prefer-code または --prefer-specs を使用して、すべての競合を非対話的に解決します。
コマンド
説明
スペック初期化
テンプレートとディレクトリ構造を使用して新しいスペック プロジェクトを足場にする
斑点の移行
既存のコードベースから仕様をリバースエンジニアリングする
斑点を適用する
編集を双方向で同期

基本的に仕様とコードの間
スペックステータス
編集済みファイルと未登録ファイルをレイヤーごとに一覧表示
スペックレビュー
エージェント経由で包括的なコードレビューを実行する
スペックFMT
構成されたフォーマッタを実行し、ハッシュを更新します
スペックチャット
ゼロスタック TUI セッションを開き、変更を自動適用します
スペック強制更新
現在のファイルの内容と一致するようにすべてのハッシュをリセットします
スペックリセット
ハッシュをクリアします (オプションで src/ または specs/technical/ を削除します)
スペックスイッチ言語
技術スタックを変更し、機能からすべてを再構築する
スペックMV
ファイルを移動してハッシュ追跡を更新する
スペックRM
ファイルを削除してハッシュ追跡を更新する
スペック git フック
git フックをインストールする (コミット前、プッシュ前、マージ後、チェックアウト後)
プロンプト
Speck には、 data/prompts/ に 5 つの特殊なゼロスタック プロンプトが付属しています。
これらは speck init でプロジェクトにインストールされ、プロジェクトごとにカスタマイズできます。
Speck には、speck init で新しいプロジェクトに組み込まれるデフォルトの AGENTS.md および ARCHITECTURE.md テンプレートが同梱されています。これらは、コードベースで作業するすべてのエージェント (ゼロスタックを含む) に共有のコア知識を提供します。
フルスペックベースの AI エージェントコンパイラー
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
2
v1.0.0
最新の
2026 年 7 月 4 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A fully spec-based AI agentic compiler. Contribute to gi-dellav/speck development by creating an account on GitHub.

GitHub - gi-dellav/speck: A fully spec-based AI agentic compiler · GitHub
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
gi-dellav
/
speck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
36 Commits 36 Commits .github .github data data packaging/ homebrew packaging/ homebrew scripts scripts src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md install.sh install.sh justfile justfile rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Spec-driven agentic compiler: simple commands for complete agent managament, allowing for reproducible and verifiable code generation, moving the traditional chat interface out of the spotlight.
note: speck is a based on zerostack and it requires a working zerostack binary on PATH .
Three-layer architecture : specs/features/ (what) → specs/technical/ (how) → src/ (code). Edit any layer, the others stay in sync.
Bidirectional sync : speck apply detects changes via BLAKE3 hashes and propagates them through the layers using specialized AI prompts.
Migration from existing projects : speck migrate reverse-engineers specs/technical/ from any codebase, then derives specs/features/ from that.
Scaffolding : speck init bootstraps a new project with canonical AGENTS.md, ARCHITECTURE.md, and prompt templates.
Conflict resolution : when both code and specs are edited, speck prompts for priority (or use --prefer-code / --prefer-specs ).
Code review : speck review runs a comprehensive, structured review via a specialized prompt, outputting a Markdown report.
Git hooks : speck git-hooks installs pre-commit, pre-push, post-merge, and post-checkout hooks that run speck apply , speck fmt , or speck status automatically.
Tech stack switching : speck switch-lang destructively changes language/framework and rebuilds everything from feature specs.
File management : speck mv and speck rm move or remove files while keeping the hash database in sync.
Change detection : speck status lists edited and unregistered files grouped by layer.
You can download compiled binaries from Github releases .
cargo install speck-dev
From source
git clone https://github.com/gi-dellav/speck.git
cd speck
cargo install --path .
Quick start
# Ensure zerostack is installed and configured
zerostack --version
# Scaffold a new speck project
speck init
# Or migrate an existing codebase
speck migrate
# Sync edits between specs and code
speck apply
# Check what changed
speck status
# Full code review
speck review --output report.md
Configuration
Speck reads Speck.toml from the project root:
Speck maintains a three-layer architecture where each layer is a different representation of the same product:
speck apply runs a 4-step pipeline:
Detect changes — compare current BLAKE3 hashes against .speck_hash.toml
feat2tech — propagate feature spec edits down to technical specs
tech2code — propagate technical spec edits down to source code
code2tech — propagate source code edits up to technical specs
Each step uses a specialized zerostack prompt at the appropriate temperature (0 for deterministic updates, higher for creative spec expansion).
When both a source file and its corresponding technical spec are edited since the last sync, speck cannot know which is authoritative. It prompts interactively:
Both spec and code changed for src/auth.rs — which takes priority?
[1] Code (update specs to match)
[2] Specs (update code to match)
Use --prefer-code or --prefer-specs to resolve all conflicts non-interactively.
Command
Description
speck init
Scaffold a new speck project with templates and directory structure
speck migrate
Reverse-engineer specs from an existing codebase
speck apply
Sync edits bidirectionally between specs and code
speck status
List edited and unregistered files grouped by layer
speck review
Run a comprehensive code review via agent
speck fmt
Run the configured formatter and refresh hashes
speck chat
Open a zerostack TUI session, then auto-apply changes
speck force-update
Reset all hashes to match current file contents
speck reset
Clear hashes (optionally delete src/ or specs/technical/)
speck switch-lang
Change tech stack and rebuild everything from features
speck mv
Move a file and update hash tracking
speck rm
Remove a file and update hash tracking
speck git-hooks
Install git hooks (pre-commit, pre-push, post-merge, post-checkout)
Prompts
Speck ships with five specialized zerostack prompts in data/prompts/ :
These are installed into your project on speck init and can be customized per-project.
Speck ships default AGENTS.md and ARCHITECTURE.md templates that are scaffolded into new projects on speck init . These provide shared core knowledge for all agents (including zerostack) working on the codebase.
A fully spec-based AI agentic compiler
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
2
v1.0.0
Latest
Jul 4, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

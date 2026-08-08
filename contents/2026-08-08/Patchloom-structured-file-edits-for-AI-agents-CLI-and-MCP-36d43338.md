---
source: "https://github.com/patchloom/patchloom"
hn_url: "https://news.ycombinator.com/item?id=49224010"
title: "Patchloom – structured file edits for AI agents (CLI and MCP)"
article_title: "GitHub - patchloom/patchloom: Structured file edits for AI agents (JSON/YAML/TOML, markdown, AST, dry-run, MCP). Not a generic filesystem MCP. · GitHub"
author: "SebTardif"
captured_at: "2026-08-08T18:23:03Z"
capture_tool: "hn-digest"
hn_id: 49224010
score: 1
comments: 0
posted_at: "2026-08-08T17:41:39Z"
tags:
  - hacker-news
  - translated
---

# Patchloom – structured file edits for AI agents (CLI and MCP)

- HN: [49224010](https://news.ycombinator.com/item?id=49224010)
- Source: [github.com](https://github.com/patchloom/patchloom)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T17:41:39Z

## Translation

タイトル: Patchloom – AI エージェント用の構造化ファイル編集 (CLI および MCP)
記事タイトル: GitHub - patchloom/patchloom: AI エージェント用の構造化ファイル編集 (JSON/YAML/TOML、マークダウン、AST、ドライラン、MCP)。汎用ファイルシステム MCP ではありません。 · GitHub
説明: AI エージェントの構造化ファイル編集 (JSON/YAML/TOML、マークダウン、AST、ドライラン、MCP)。汎用ファイルシステム MCP ではありません。 - パッチルーム/パッチルーム

記事本文:
GitHub - patchloom/patchloom: AI エージェント用の構造化ファイル編集 (JSON/YAML/TOML、マークダウン、AST、ドライラン、MCP)。汎用ファイルシステム MCP ではありません。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
パッチルーム
/
パッチルーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,826 コミット 1,826 コミット .github .github アセット アセット ベンチ ベンチ チョコレート

Chocolatey デモ デモ ドキュメント ドキュメントの例 例 ファズ ファズ mcpb mcpb スクリプト スクリプト src src テスト テスト .editorconfig .editorconfig .gitignore .gitignore .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml GEMINI.md GEMINI.md GOVERNANCE.md GOVERNANCE.md ライセンス ライセンス ライセンス-APACHE ライセンス-APACHE MAINTAINERS.md MAINTAINERS.md Makefile Makefile PATCHLOOM.md PATCHLOOM.md PLAN-for-code-review-items.md PLAN-for-code-review-items.md README.md README.md REPO-SETUP.md REPO-SETUP.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md book.toml book.tomldeny.tomldeny.toml dist-workspace.toml dist-workspace.toml Glama.json Glama.json lychee.toml lychee.toml release-please-config.json release-please-config.json錆びたツールチェーン.toml錆びたツールチェーン.toml錆びたfmt.toml錆びたfmt.tomlサーバー.json server.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バイナリが 1 つ。すべてのプラットフォーム。 AI エージェントの構造化されたファイル編集。
Patchloom は、AI コーディング エージェントに、あらゆるオペレーティング システム上で安全で構造化されたファイル編集を提供するシングル バイナリ CLI です。 JSON、YAML、および TOML をセレクター (正規表現ではなく) で編集し、コメントを保存し、20 の言語にわたるコード構造を理解し、複数のファイル編集を 1 つのツール呼び出しにまとめて実行し、Linux、macOS、および Windows 上で同様に動作します。
汎用ファイルシステム MCP ではありません。デフォルトの MCP ファイルシステム サーバーは、ファイルをテキストとして読み取り/書き込みします。 Patchloom は、ドライラン プレビュー、パーサーによる設定とマークダウン編集、AST 操作、元に戻すことができる複数ファイルのバッチ/TX、ホスト用の安定した error_kind ピールを追加します。完全なコーディング エージェント (Claude Code、Codex、Cursor) がループを所有します。 Patchloom は、彼ら (または

Rust embedder) を呼び出します。
# コメントや書式を壊すことなく、セレクターで YAML 値を編集します
patchloom ドキュメント セット config.yaml データベース.ポート 5432 --apply
# バッチ 6 ファイル編集を単一のツール呼び出しに変換
パッチルーム バッチ --apply << ' EOF '
doc.set package.json バージョン「2.0.0」
doc.set config.yaml app.version "2.0.0"
doc.set config.toml プロジェクトのバージョン "2.0.0"
README.md "1.0.0" "2.0.0" を置き換えます
CHANGELOG.md "1.0.0" "2.0.0" を置き換えます
file.create バージョン "2.0.0"
終了後
なぜパッチルームなのか？ |インストール |クイックスタート |コマンド |比較 |いつ何を使用するか |建築 |ステータス
AI エージェントはツール呼び出しを通じてファイルを編集します。各呼び出しは LLM への往復です。タスクが構成ファイルにアクセスすると、そのプロセスには次の 3 つの失敗モードがあります。
構文の破損。エージェントは、JSON、YAML、または TOML でテキスト置換を使用し、無効な出力 (中括弧の不一致、インデントの破損、コメントの欠落) を生成します。
往復税。 6 つのファイルを編集すると、6 つの個別のツール呼び出しが行われることになります。それぞれは、LLM が生成、実行、結果を読み取り、次の呼び出しを計画するのを待ちます。
プラットフォームの断片化。 Linux では、エージェントは sed 、 jq 、 grep を使用します。 Windows では、それらはどれも存在しません。エージェントは冗長な PowerShell にフォールバックするか、不慣れな構文でエラーを起こします。
問題
パッチルームがそれを解決する方法
構文の破損
doc コマンドはファイルを解析し、セレクター パスによって値を変更し、有効な出力を書き込みます。コメントと書式設定は保持されます。正規表現は必要ありません。
往復税
バッチと tx は、N 個の操作を 1 つのツール呼び出しに結合します。 6 つのファイル編集が 1 つのコマンドになり、失敗時にアトミック ロールバックが行われます。
プラットフォームの断片化
依存関係のない単一の静的バイナリ。 Linux、macOS、および Windows 上で同じコマンド、同じフラグ、同じ動作。
パッチルームで何が変わるのか
パッチルームなし (6 つのツール呼び出し)
エージェント: ファイル 1 を編集 --- ツール呼び出し ---▶ 15 秒
エージェント: エド

it ファイル 2 ────ツール呼び出し────▶ 15 秒
エージェント: ファイル 3 を編集 ─── ツール呼び出し ──▶ 15 秒
エージェント: ファイル 4 の編集 ──── ツール呼び出し ────▶ 15 秒
エージェント: ファイル 5 の編集 ──── ツール呼び出し ────▶ 15 秒
エージェント: ファイル 6 の編集 ────ツール呼び出し────▶ 15 秒
合計: ~90 秒
パッチルームバッチあり（1ツールコール）
エージェント: とのバッチ
全 6 回の編集 --- ツール呼び出し ---▶ 25 秒
5往復を節約
合計: ~25秒
主な機能
能力
何をするのか
例
パーサーによる編集
セレクターによる JSON/YAML/TOML の編集、コメントと書式設定の保持
doc set config.yaml db.port 5432 --apply
1 回の呼び出しで N ファイルをバッチ処理する
バッチと TX は操作をロールバック付きの 1 つのツール呼び出しに結合します
バッチ --apply < ops.txt
コメントの保存
YAML/TOML コメントは配列のサイズ変更を含むすべての編集後も存続します
doc append config.yaml tags '"v2"' --apply
見出しを意識したマークダウン
セクション、表、箇条書きを行番号ではなく見出しで編集します
md table-append README.md --Heading "API" --row "| new | row |" --適用する
AST 対応のコード操作
20 言語にわたる記号の一覧表示、名前変更、置換、分析
ast rename src/ --old old_name --new new_name --apply
アトミックロールバック
strict: true は、フォーマットまたは検証ステップが失敗した場合にすべてのファイルを元に戻します
tx plan.json --apply
MCPサーバー
すべての操作を構造化された MCP ツール呼び出しとして公開します
パッチルーム mcp サーバー
オプションの CLI サンドボックス
Reject ../ / 読み取りおよび書き込み時に絶対パスが --cwd からエスケープされます (デフォルトではオフ、MCP は常にオン)
patchloom --cwd <ws> --contain search … / create … --apply
クロスプラットフォーム
Linux、macOS、Windows でも同様の動作。 sed 、 jq 、 grep は必要ありません。
どこでも同じバイナリ
パッチルームとネイティブ ツールをいつ使用するか
Patchloom は、すべてのファイル操作に代わるものではありません。その指示により、いつそれを使用するべきか、いつネイティブ ツールが高速になるかをエージェントに正確に通知します。
パッチルームは

シンプルな単一ファイルの編集ではネイティブ ツールよりも高速ではありません。これらにはネイティブ ツールを使用してください。ただし、ネイティブ テキスト置換では構造化ファイルを安全に編集できません。YAML の sed では、インデントが破損したり、コメントが削除されたり、無効な構文が生成されたりする可能性があります。 doc set はファイルを解析し、セレクターによって値を変更し、有効な出力を書き込みます。その保証がポイントです。
Patchloom が高速な場合は、複数ファイルのバッチ処理です。ネイティブ ツールによる 6 つのファイル編集は、LLM への 6 往復を意味します。 1 つのバッチ呼び出しは、1 回のラウンドトリップで同じ作業を実行します。
タスク PL-CLI MCP ネイティブ
────────── ────── ────── ──────
検索 18.5 秒 12.7 秒 13.9 秒 ◀ ~同じ
置換 36.1s 26.6s 26.1s ◀ ~同じ
doc_set 30.9s 16.9s 13.7s ◀ ネイティブ最速
md_table 15.5 秒 13.5 秒 15.3 秒 ◀ MCP 最速
tx_multi_file 41.4 秒 28.5 秒 22.9 秒 ◀ ネイティブ最速
バッチ_6_ファイル 50.6 秒 46.6 秒 30.3 秒 ◀ ネイティブ最速
バッチ_ミックスドオプス 24.7 秒 13.6 秒 20.9 秒 ◀ MCP 最速
yaml_comment_preserve 18.1 秒 11.6 秒 16.1 秒 ◀ MCP 最速
md_insert 15.0s 11.7s 15.7s ◀ MCP 最速
file_ops 26.0秒 16.6秒 17.2秒 ◀ ~同じ
整然とした 45.0 秒 30.3 秒 41.7 秒 ◀ MCP 最速
────────── ────── ────── ──────
合計 321.9秒 228.5秒 233.8秒
構造化ツールがシェル構文構築を完全にスキップするため、MCP モードが全体的に優れています (228.5 秒対ネイティブ 233.8 秒)。 MCP が 5/11 タスクを獲得。ネイティブが3/11で勝利。 3はネクタイです。 CLI モードは、シェル構築のオーバーヘッドのため、常に最も遅くなります。
各 GitHub リリース (Homebrew、Scoop、crate、
npm、リリース)。 Windows では、Scoop が推奨されます。ウィンゲット
( Patchloom.Patchloom ) はリリースごとに公開され、通常は最新のものです
Microsoft のパブリッシュ パイプラインの後 (winget ソース更新の場合)

検索は
古い）。 Chocolatey は、コミュニティのモデレーションが実行されている間、遅延することがよくあります。
# 自作 (macOS/Linux)
brew install パッチルーム/タップ/パッチルーム
# crates.io (Rust 1.95 以降が必要、MCP サーバーを含む)
貨物取り付けパッチルーム
# npm / npx (GitHub リリースからプラットフォーム バイナリをダウンロードします)
npx patchloom --version
# または: npm install -g patchloom
# Scoop (Windows; 推奨される Windows チャネル)
スクープ バケット パッチルームを追加 https://github.com/patchloom/スクープ - バケット
スクープインストールパッチルーム/パッチルーム
Linux、macOS、および Windows 用の事前に構築されたバイナリは、
リリースページ。
シェルのインストールを参照してください。
インストーラー スクリプト、ソース ビルド、シェル補完、および winget /
チョコレートっぽいノート。
MCP レジストリ名: mcp-name: io.github.patchloom/patchloom
VS Code、Cursor、Windsurf、または VSCodium のコンパニオン拡張機能をインストールします。
拡張機能は CLI を自動検出し (または自動的にインストールし)、
AGENTS.md、MCP サーバーを構成し、コマンドにクイック アクションを追加します。
パレット。詳細については、エディタ拡張ガイドを参照してください。
パッチルームの初期化
これにより、新しいプロジェクトに AGENTS.md が作成されるか、既存のエージェント指示ファイルにルールが追加され、シェル補完が提供され、MCP 構成の機会が検出されます。確認プロンプトをスキップするには、-y を渡します。
ルールのテキストのみが必要な場合:
パッチルーム エージェント ルール >> AGENTS.md
# または出力を調整します。
patchloom Agent-rules --mode mcp >> AGENTS.md # MCP のみ (CLI の例はありません)
patchloom Agent-rules --platform Windows >> AGENTS.md # Windows のみの構文
.vscode/ または .cursor/ が存在する場合、init はすぐにコピーできる .vscode/mcp.json または .cursor/mcp.json スニペットも出力します。
AI エージェントは AGENTS.md を読み取り、パッチルームとネイティブ ツールをいつ使用するかを学習します。
# パーサー支援: 値を変更し、コメントと書式設定を保持します
patchloom ドキュメント セット config.yaml データベース.ポート 5432 --apply

3. 複数の編集を 1 つの呼び出しにまとめて実行する
パッチルーム バッチ --apply << ' EOF '
doc.set config.json バージョン '"2.0.0"'
md.upsert_bullet AGENTS.md "ルール" "- 常にテスト"
src/main.rs "v1" "v2" を置き換えます
終了後
値は JSON から始まります。引用符で囲まれていない 2.0 は数値になります。ネストされた文字列を強制する
上記のように引用符で囲みます (Unix シェル)。
または、形式を指定して JSON プランを使用し、ライフサイクルを検証します。
{
「バージョン」: 1 、
「操作」: [
{ "op" : " doc.set " 、 "path" : " config.json " 、 "selector" : " version " 、 "value" : " 2.0.0 " },
{ "op" : " md.upsert_bullet " 、 "path" : " AGENTS.md " 、 "Heading" : " Rules " 、 "bullet" : " - 常にテスト " },
{ "op" : " replace " 、 "path" : " src/main.rs " 、 "old" : " v1 " 、 "new" : " v2 " }
]、
"format" : [{ "cmd" : " Cargo fmt --all " }],
"検証" : [{ "cmd" : "貨物テスト" , "必須" : true }]
}
patchloom tx plan.json --apply
tx プランは信頼できる入力です。 format と validate は、ホスト シェル (Unix の場合は sh -c、Windows の場合は cmd /C) を介して cmd フィールドを実行するため、信頼できるプランのみを実行します。
4. または、構造化ツール呼び出しに MCP を使用します (シェル構文は使用しません)。
MCP サポートを使用してインストールした後、サーバーを起動します。
パッチルーム mcp サーバー
MCP 対応エージェントは、シェルの引用やコマンドの構築を行わずに、構造化された JSON としてパッチルーム ツールを直接呼び出します。エージェントは、 patchloom doc set config.json version '"2.0"' --apply をビルドする代わりに、 {"path": "config.json", "selector": "version", "value": "2.0"} を送信します。
コーディング エージェント: 11 ツール パックの PATCHLOOM_MCP_SURFACE=core を設定します ( list_files 、 search/read/repla

[切り捨てられた]

## Original Extract

Structured file edits for AI agents (JSON/YAML/TOML, markdown, AST, dry-run, MCP). Not a generic filesystem MCP. - patchloom/patchloom

GitHub - patchloom/patchloom: Structured file edits for AI agents (JSON/YAML/TOML, markdown, AST, dry-run, MCP). Not a generic filesystem MCP. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
patchloom
/
patchloom
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,826 Commits 1,826 Commits .github .github assets assets benches benches chocolatey chocolatey demo demo docs docs examples examples fuzz fuzz mcpb mcpb scripts scripts src src tests tests .editorconfig .editorconfig .gitignore .gitignore .release-please-manifest.json .release-please-manifest.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml GEMINI.md GEMINI.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE LICENSE-APACHE LICENSE-APACHE MAINTAINERS.md MAINTAINERS.md Makefile Makefile PATCHLOOM.md PATCHLOOM.md PLAN-for-code-review-items.md PLAN-for-code-review-items.md README.md README.md REPO-SETUP.md REPO-SETUP.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md book.toml book.toml deny.toml deny.toml dist-workspace.toml dist-workspace.toml glama.json glama.json lychee.toml lychee.toml release-please-config.json release-please-config.json rust-toolchain.toml rust-toolchain.toml rustfmt.toml rustfmt.toml server.json server.json View all files Repository files navigation
One binary. Every platform. Structured file edits for AI agents.
Patchloom is a single-binary CLI that gives AI coding agents safe, structured file editing on any operating system. It edits JSON, YAML, and TOML by selector (not regex), preserves comments, understands code structure across 20 languages, batches multiple file edits into one tool call, and works identically on Linux, macOS, and Windows.
Not a generic filesystem MCP. Default MCP filesystem servers read/write files as text. Patchloom adds dry-run previews, parser-backed config and markdown edits, AST ops, multi-file batch / tx with undo, and stable error_kind peels for hosts. Full coding agents (Claude Code, Codex, Cursor) own the loop; Patchloom is the tool layer they (or a Rust embedder) call.
# Edit a YAML value by selector without breaking comments or formatting
patchloom doc set config.yaml database.port 5432 --apply
# Batch 6 file edits into a single tool call
patchloom batch --apply << ' EOF '
doc.set package.json version "2.0.0"
doc.set config.yaml app.version "2.0.0"
doc.set config.toml project.version "2.0.0"
replace README.md "1.0.0" "2.0.0"
replace CHANGELOG.md "1.0.0" "2.0.0"
file.create VERSION "2.0.0"
EOF
Why Patchloom? | Install | Quick start | Commands | Comparison | When to use what | Architecture | Status
AI agents edit files through tool calls. Each call is a round-trip back to the LLM. When a task touches config files, that process has three failure modes:
Syntax corruption. The agent uses text replacement on JSON, YAML, or TOML and produces invalid output (mismatched braces, broken indentation, lost comments).
Round-trip tax. Editing 6 files means 6 separate tool calls. Each one waits for the LLM to generate, execute, read the result, and plan the next call.
Platform fragmentation. On Linux the agent uses sed , jq , grep . On Windows, none of those exist. The agent falls back to verbose PowerShell or makes errors with unfamiliar syntax.
Problem
How patchloom solves it
Syntax corruption
doc commands parse the file, change the value by selector path, and write valid output. Comments and formatting are preserved. No regex needed.
Round-trip tax
batch and tx combine N operations into 1 tool call. Six file edits become one command with atomic rollback on failure.
Platform fragmentation
Single static binary with zero dependencies. Same commands, same flags, same behavior on Linux, macOS, and Windows.
What changes with patchloom
Without patchloom (6 tool calls)
Agent: edit file 1 ─── tool call ───▶ 15s
Agent: edit file 2 ─── tool call ───▶ 15s
Agent: edit file 3 ─── tool call ───▶ 15s
Agent: edit file 4 ─── tool call ───▶ 15s
Agent: edit file 5 ─── tool call ───▶ 15s
Agent: edit file 6 ─── tool call ───▶ 15s
Total: ~90s
With patchloom batch (1 tool call)
Agent: batch with
all 6 edits ─── tool call ───▶ 25s
5 round-trips saved
Total: ~25s
Key capabilities
Capability
What it does
Example
Parser-backed edits
Edit JSON/YAML/TOML by selector, preserving comments and formatting
doc set config.yaml db.port 5432 --apply
Batch N files in 1 call
batch and tx combine operations into one tool call with rollback
batch --apply < ops.txt
Comment preservation
YAML/TOML comments survive all edits, including array resizing
doc append config.yaml tags '"v2"' --apply
Heading-aware markdown
Edit sections, tables, and bullets by heading, not line number
md table-append README.md --heading "API" --row "| new | row |" --apply
AST-aware code ops
List, rename, replace, and analyze symbols across 20 languages
ast rename src/ --old old_name --new new_name --apply
Atomic rollback
strict: true reverts every file if format or validate steps fail
tx plan.json --apply
MCP server
Expose all operations as structured MCP tool calls
patchloom mcp-server
Optional CLI sandbox
Reject ../ / absolute path escapes from --cwd on reads and writes (off by default; MCP always on)
patchloom --cwd <ws> --contain search … / create … --apply
Cross-platform
Identical behavior on Linux, macOS, Windows. No sed , jq , grep required.
Same binary everywhere
When to use patchloom vs native tools
Patchloom is not a replacement for all file operations. Its instructions tell agents exactly when to use it and when native tools are faster:
Patchloom is not faster than native tools for simple, single-file edits. Use native tools for those. But native text replacement cannot safely edit structured files: a sed on YAML can corrupt indentation, strip comments, or produce invalid syntax. doc set parses the file, changes the value by selector, and writes valid output. That guarantee is the point.
Where patchloom is faster is multi-file batching. Six file edits via native tools means six round-trips to the LLM. One batch call does the same work in a single round-trip.
Task PL-CLI MCP Native
────────────────────── ────── ────── ──────
search 18.5s 12.7s 13.9s ◀ ~same
replace 36.1s 26.6s 26.1s ◀ ~same
doc_set 30.9s 16.9s 13.7s ◀ native fastest
md_table 15.5s 13.5s 15.3s ◀ MCP fastest
tx_multi_file 41.4s 28.5s 22.9s ◀ native fastest
batch_6_files 50.6s 46.6s 30.3s ◀ native fastest
batch_mixed_ops 24.7s 13.6s 20.9s ◀ MCP fastest
yaml_comment_preserve 18.1s 11.6s 16.1s ◀ MCP fastest
md_insert 15.0s 11.7s 15.7s ◀ MCP fastest
file_ops 26.0s 16.6s 17.2s ◀ ~same
tidy 45.0s 30.3s 41.7s ◀ MCP fastest
────────────────────── ────── ────── ──────
TOTAL 321.9s 228.5s 233.8s
MCP mode wins overall (228.5s vs 233.8s native) because structured tool calls skip shell syntax construction entirely. MCP wins 5/11 tasks; native wins 3/11; 3 are ties. CLI mode is always slowest due to shell construction overhead.
Prefer channels that track each GitHub Release (Homebrew, Scoop, crates,
npm, Releases). On Windows, Scoop is recommended. winget
( Patchloom.Patchloom ) is published per release and is usually current
after Microsoft's publish pipeline ( winget source update if search is
stale). Chocolatey often lags while community moderation runs.
# Homebrew (macOS/Linux)
brew install patchloom/tap/patchloom
# crates.io (requires Rust 1.95+, includes MCP server)
cargo install patchloom
# npm / npx (downloads the platform binary from GitHub Releases)
npx patchloom --version
# or: npm install -g patchloom
# Scoop (Windows; recommended Windows channel)
scoop bucket add patchloom https: // github.com / patchloom / scoop - bucket
scoop install patchloom / patchloom
Pre-built binaries for Linux, macOS, and Windows are on the
Releases page.
See Installation for shell
installer scripts, source builds, shell completions, and winget /
Chocolatey notes.
MCP Registry name: mcp-name: io.github.patchloom/patchloom
Install the companion extension for VS Code, Cursor, Windsurf, or VSCodium:
The extension auto-discovers the CLI (or installs it for you), generates
AGENTS.md, configures MCP servers, and adds Quick Actions to the command
palette. See the Editor Extension guide for details.
patchloom init
This creates AGENTS.md in a new project or appends the rules to an existing agent instructions file, offers shell completions, and detects MCP configuration opportunities. Pass -y to skip confirmation prompts.
If you only want the rules text:
patchloom agent-rules >> AGENTS.md
# Or tailor the output:
patchloom agent-rules --mode mcp >> AGENTS.md # MCP-only (no CLI examples)
patchloom agent-rules --platform windows >> AGENTS.md # Windows-only syntax
If .vscode/ or .cursor/ exists, init also prints ready-to-copy .vscode/mcp.json or .cursor/mcp.json snippets.
Your AI agent reads AGENTS.md and learns when to use patchloom vs native tools.
# Parser-backed: changes the value, preserves comments and formatting
patchloom doc set config.yaml database.port 5432 --apply
3. Batch multiple edits into one call
patchloom batch --apply << ' EOF '
doc.set config.json version '"2.0.0"'
md.upsert_bullet AGENTS.md "Rules" "- Always test"
replace src/main.rs "v1" "v2"
EOF
Values are JSON-first: unquoted 2.0 becomes a number. Force a string with nested
quotes as above (Unix shells).
Or use a JSON plan with format and validate lifecycle:
{
"version" : 1 ,
"operations" : [
{ "op" : " doc.set " , "path" : " config.json " , "selector" : " version " , "value" : " 2.0.0 " },
{ "op" : " md.upsert_bullet " , "path" : " AGENTS.md " , "heading" : " Rules " , "bullet" : " - Always test " },
{ "op" : " replace " , "path" : " src/main.rs " , "old" : " v1 " , "new" : " v2 " }
],
"format" : [{ "cmd" : " cargo fmt --all " }],
"validate" : [{ "cmd" : " cargo test " , "required" : true }]
}
patchloom tx plan.json --apply
tx plans are trusted input. format and validate run their cmd fields through the host shell ( sh -c on Unix, cmd /C on Windows), so only run plans you trust.
4. Or use MCP for structured tool calls (no shell syntax)
After installing with MCP support , start the server:
patchloom mcp-server
MCP-capable agents call patchloom tools directly as structured JSON, with no shell quoting or command construction. The agent sends {"path": "config.json", "selector": "version", "value": "2.0"} instead of building patchloom doc set config.json version '"2.0"' --apply .
Coding agents: set PATCHLOOM_MCP_SURFACE=core for an 11-tool pack ( list_files , search/read/repla

[truncated]

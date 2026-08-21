---
source: "https://github.com/waterflane/ContextForge"
hn_url: "https://news.ycombinator.com/item?id=49386431"
title: "Show HN: ContextForge – context engineering platform for AI-assisted development"
article_title: "GitHub - waterflane/ContextForge: A context engineering platform for AI-assisted software development. · GitHub"
image: "https://repository-images.githubusercontent.com/1298691131/1ea28b94-4d2d-44fc-b32a-97abd9846d6c"
author: "waterflane"
captured_at: "2026-08-21T11:17:23Z"
capture_tool: "hn-digest"
hn_id: 49386431
score: 1
comments: 0
posted_at: "2026-08-21T11:14:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ContextForge – context engineering platform for AI-assisted development

- HN: [49386431](https://news.ycombinator.com/item?id=49386431)
- Source: [github.com](https://github.com/waterflane/ContextForge)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T11:14:08Z

## Translation

タイトル: Show HN: ContextForge – AI 支援開発のためのコンテキスト エンジニアリング プラットフォーム
記事のタイトル: GitHub - Waterflane/ContextForge: AI 支援ソフトウェア開発のためのコンテキスト エンジニアリング プラットフォーム。 · GitHub
説明: AI 支援ソフトウェア開発のためのコンテキスト エンジニアリング プラットフォーム。 - ウォーターフレーン/ContextForge

記事本文:
GitHub - Waterflane/ContextForge: AI 支援ソフトウェア開発のためのコンテキスト エンジニアリング プラットフォーム。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ウォータープレーン
/
ContextForge
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
82 コミット 82 コミット フォルダーとファイル
.github .github docs docs スクリプト スクリプト src/ contextforge src/ contextforge テスト テスト .e

ditorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントをコーディングするための、制限されたレビュー可能なリポジトリ コンテキストを構築します。
ContextForge はローカル リポジトリをスキャンし、決定論的な構造マップを構築します。
制限されたタスク認識の検出をサポートし、移植可能なコンテキスト パッケージを生成します
そして引き継ぎ。外部コーディング エージェントが何を参照するかを決定するのに役立ちます
ContextForge にソース コードを編集したり、任意のプログラムを実行する権限を与えずに
コマンド。
クイックスタート ·
CLI・
構成・
ウィキ ·
ディスカッション ·
貢献方針
ContextForge はプレアルファ版のソフトウェアです。バージョン 0.4.2 が最初のパブリック リリースです
候補者。 Discovery ベンチマークは実験的なものであり、その結果は次のとおりです。
記録されたプロバイダー、モデル、構成、ソースとともにレビューされます
スナップショット。
決定的なリポジトリ インベントリ。安定した順序でファイルをスキャンし、
移植可能なパス、ハッシュ、言語分類、無視ルール、および制限付き
読みます。
レビュー可能なセレクション。正確なファイル、ディレクトリ、GitWildMatch を選択してください
パターン、またはライン範囲、または設定されたモデルに限定された提案を求めます。
ローカル リポジトリ インテリジェンス。不変の構造的およびオプションのストア
.contextforge/index でのセマンティック インデックスの生成。
ポータブルなアーティファクト。 Markdown または JSON コンテキスト パッケージ、JSON タスクのエクスポート
ハンドオフ、およびコンパイルされた Markdown プロンプト。
明示的な信頼境界。 ContextForge はリポジトリ ソースを編集しません。
リポジトリ コードを実行したり、シェル ツールを公開したり、Git の状態を変更したりします。
自動化に適した出力。構造化された結果は標準出力に残ります

イル
進行状況と診断は標準エラー出力に残ります。
フローチャート LR
R["リポジトリ"] --> S["スキャン / インデックス"]
S --> D[「タスク認識ディスカバリー」]
D --> B["境界付き選択"]
B --> P["コンテキストパッケージ"]
P --> A["外部コーディング エージェント"]
読み込み中
モデルは、スキャン、ツリー、手動コンテキスト パッケージ、および
構造のみのインデックス作成:
contextforge スキャン。
contextforge インデックスの構築。 --プロバイダーなし
contextforge コンテキスト作成 。 \
--include pyproject.toml \
--ディレクトリ src/contextforge/context \
--exclude " **/__init__.py " \
--フォーマットマークダウン\
--output context.md
インストール
ContextForge には Python 3.12 以降が必要です。公開されたディストリビューションをインストールします。
python -m pip install contextforge-repo
分離されたコマンドライン インストールの場合は、次のいずれかのツール マネージャーを使用します。
pipx インストール contextforge-repo
# または
UV ツールのインストール contextforge-repo
PyPI ディストリビューションの名前は contextforge-repo です。インポートパッケージは残ります
contextforge となり、インストールされるコマンドは contextforge と ctxf のままになります。
同様の名前の context-forge-cli ディストリビューションは異なります。
無関係のプロジェクト。
代わりにチェックアウトされたソース ツリーをインストールするには:
git clone https://github.com/waterflane/ContextForge.git
cd ContextForge
Python -m venv .venv
ソース .venv/bin/activate
python -m pip install --upgrade pip
python -m pip インストール 。
Windows PowerShell:
git clone https://github.com/waterflane/ContextForge.git
場所の設定 ContextForge
Python - m venv .venv
.\.venv\Scripts\Activate.ps1
Python - m pip install -- pip のアップグレード
python -m pip インストール 。
インストールにより、同等の contextforge および ctxf コンソールが提供されます
コマンド。 python -m contextforge もサポートされています。
ContextForge 状態を書き込まずにリポジトリを検査します。
contextforge スキャン。
contextforge ツリー。 -- 深さ 2
contextforge コンテキスト作成 。 `
-- ' pyproject.toml を含める

' `
-- ディレクトリ ' src/contextforge/context ' `
-- ' **/__init__.py ' ` を除外します
-- json のフォーマット `
-- ' context.json ' を出力します
contextforge コンテキスト検査「 context.json 」
構造のみのローカル インデックスを構築し、そのステータスを検査します。
contextforge インデックスの構築。 -- プロバイダーなし
contextforge インデックスのステータス。
ヒント
関連するファイルがすでにわかっている場合は、手動でコンテキストを作成してください。
タスクが不慣れなコードにまたがっており、
サポートされているモデルプロバイダー。
外部コーディング エージェント用のコンパクトなレビュー パケットを作成します。
ソースをモデルに送信せずにリポジトリをマッピングします。
古い、欠落している、または失敗したインデックス レコードを検査します。
可能性のあるエントリ ポイント、テスト、構成、および依存関係を発見します。
タスク;
オリジナルなしでレビューできる検証済みのハンドオフを保存します。
チェックアウト;
バージョン管理されたマニフェストに対する検出品質と再現性のベンチマーク。
すべてのコマンドは --help をサポートしています。 Advanced またはを使用する前にグループ ヘルプを実行してください
突然変異操作。
グローバル診断オプションは、 --log-level 、 --log-format 、 --log-file 、
反復可能な --log-component 、 --no-log-file 、 --no-color 、および -v / -vv 。
詳細な構文、デフォルト、ストリーム、副作用、間違い、例については、
Wiki CLI リファレンス。
プロジェクト構成は閉じられており、バージョン管理された TOML です。解決順序は次のとおりです。
サポートされている CONTEXTFORGE_* 環境変数。
.contextforge/config.local.toml ;
.contextforge/config.toml 、または明示的な --config PATH ;
主にサポートされている環境変数は次のとおりです。
CONTEXTFORGE_MODEL_CONTEXT_WINDOW ;
CONTEXTFORGE_MODEL_CONNECT_TIMEOUT ;
CONTEXTFORGE_MODEL_READ_TIMEOUT ;
CONTEXTFORGE_MODEL_OPERATION_TIMEOUT ;
CONTEXTFORGE_JSON_REPAIR_ATTEMPTS ;
CONTEXTFORGE_LOG_LEVEL 、 CONTEXTFORGE_LOG_FORMAT 、
CONTEXTFORGE_LOG_FILE および CONTEXTFORGE_LOG_COMPONENTS 。
デフォルトのプロバイダーは、

地元のオラマ
http://127.0.0.1:11434/api/chat モデル qwen2.5-coder:7b を使用します。使用する
--provider none (構造のみのインデックス作成の場合)。 openai互換
プロバイダーとその lmstudio CLI エイリアスには、正確なモデル ID と適切なモデル ID が必要です。
Base_url 。
モデルに基づく検出では、構成されたプロバイダーが
指定されたモデルが利用可能です。 ContextForge で設定された context_window は次のことを行ってはなりません
そのプロバイダーによって実際にロードされるウィンドウを超えます。解決されたポリシーを検査する
contextforge 診断プロバイダー PATH を使用して長時間実行する前に。
資格情報構成では、環境変数の名前のみが保存されます。
資格情報_環境 ;資格情報の値はリクエスト時に解決されます。を参照してください。
構成ガイドと
Wiki 設定リファレンス。
新規ではメモリ内に現在の構造的証拠が構築され、ロードされません。
永続化されたセマンティック レコードまたはリポジトリ マップ。
インデックス付きには読み取り可能なアクティブなインデックスが必要で、現在のインデックス付きを使用します
構造、セマンティクス、マップ。
ハイブリッドがデフォルトです。現在のインデックス証拠から始まり、埋められます
ライブスナップショットからの構造的なギャップを除去し、明示的に新しいスナップショットにフォールバックします。
有効なインデックスが存在しない場合の構造。
成功したすべての選択は、現在のソース ID に対して検証されます。
モデルに基づいた実行では、異なる有効な選択が生成される場合があります。 ContextForge の主張
決定論的モデルではなく、同じ検証結果に対する決定論的レンダリング
行動。 Discovery の出力とベンチマークを参照してください。
手動パッケージでは明示的なセレクターを使用します。インクルードセレクターなしですべて選択可能
スナップショット ファイルは、設定された制限まで含まれます。
contextforge コンテキスト作成 。 --include README.md --format マークダウン
contextforge コンテキスト作成 。 --directory src --exclude " **/__init__.py "
contextforge コンテキスト作成 。 --glob " testing/test_*.py " --no-include-tree
contextforgeコンテキスト作成

。 \
--include pyproject.toml \
--include-lines pyproject.toml:1-24 \
--json の形式
自動モードでは空ではないタスクが必要であり、手動ディレクトリは受け入れられません。
glob、または行範囲セレクター:
contextforge コンテキスト提案 。 \
--task " トレース構成の優先順位 " \
--ディスカバリーハイブリッド \
--format マークダウン
contextforge コンテキスト作成 。 \
--task " トレース構成の優先順位 " \
--ディスカバリーハイブリッド \
--git-diff が動作中 \
--json の形式\
--handoff.json を出力 \
--prompt-output プロンプト.md
context assign はソースまたはインデックスの状態を書き込みませんが、現在の診断を書き込みます。
ポリシーは、 .contextforge/runs の下に安全な概要を書き込むことができます。出力アーティファクトは次のとおりです。
原子的に書かれています。既存の宛先には --force が必要ですが、そのオプションは
利用可能です。
ベンチマーク検出は 0.4.2 では実験的です。モデルに裏付けられた再現性
測定値は記録された器具の状態を観察したものであり、保証するものではありません。
最初に構成されたプロバイダーを起動し、同じコンテキスト ウィンドウ値を使用します。
ContextForge とプロバイダー ランタイム。
contextforge ベンチマーク検出 ' C:\Repositories ' `
-- タスク ' .\benchmarks\discovery.json ' `
-- モード ' フレッシュ、インデックス付き、ハイブリッド ' `
-- 3 ` を繰り返します
-- json のフォーマット `
-- 出力「 .\benchmark-report.json 」
ランナーはリポジトリ/インデックス読み取り専用であり、設定されたファイルのログ記録を無効にし、
完了、失敗、およびキャンセルされた実行が結果に記録されます。終了コード 3 の意味
このコマンドは、少なくとも 1 つのタスクを含む完全なベンチマーク レポートを作成しました。
期待、または予算の失敗。標準出力または要求された出力を破棄しないでください
そのコードを処理するときにファイルを作成します。すべての実行はマニフェスト制限によって制限されたままになります。
プロバイダーの再試行制限、操作のタイムアウト、および構成されたコンテキスト ウィンドウ。
ツリー: テキスト、マークダウン、json ;
提案: text 、互換性エイリアステーブル、markdown 、json ;
コンテキストパッケージ: mar

kdown 、 json ;
インデックスのステータスと診断: table 、 json ;
検出ベンチマーク: text 、 markdown 、 json 。
出力パスが指定されていない場合は、選択された結果が標準出力に書き込まれます。
進行状況、ログ、エラーは stderr を使用し、解析可能な JSON stdout を保持します。一部
コマンドは、ファイルの書き込み後に確認を標準出力に出力します。ベンチマーク出力
ファイルは例外であり、標準出力は空のままになります。
一般的なプロセス終了コードは、成功の場合は 0、操作の失敗の場合は 1、操作の失敗の場合は 2 です。
無効な使用または構成の場合は 130、キャンセルの場合は 130。終了コード 3
コマンド固有の意味: scan --fail-on-error による読み取り不可能なエントリ、または
回帰失敗を伴う完了した検出ベンチマーク。
ContextForge は、型付き Python モジュラー モノリスです。コアアプリケーションとドメイン
ロジックは Typer、FastAPI、モデルプロバイダーの実装から独立したままになります。
ストレージ アダプター、および将来のエディターの統合。スキャナーが作成するのは、
検証済みのスナップショット。インテリジェンスは構造的事実とオプションのセマンティクスを抽出します
解釈。 Discovery は限定された候補を選択します。コンテキストとハンドオフ
モジュールは移植可能な成果物を具体化します。 CLI、HTTP、および MCP はシン インターフェイスです。
依存関係についてはアーキテクチャの概要を参照してください。
境界とtrのセキュリティポリシー

[切り捨てられた]

## Original Extract

A context engineering platform for AI-assisted software development. - waterflane/ContextForge

GitHub - waterflane/ContextForge: A context engineering platform for AI-assisted software development. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
waterflane
/
ContextForge
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
82 Commits 82 Commits Folders and files
.github .github docs docs scripts scripts src/ contextforge src/ contextforge tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Build bounded, reviewable repository context for coding agents.
ContextForge scans local repositories, builds deterministic structural maps,
supports bounded task-aware discovery, and produces portable context packages
and handoffs. It helps you decide what an external coding agent should see
without giving ContextForge permission to edit source code or run arbitrary
commands.
Quick start ·
CLI ·
Configuration ·
Wiki ·
Discussions ·
Contribution policy
ContextForge is pre-alpha software. Version 0.4.2 is the first public-release
candidate. Discovery benchmarking is experimental and its results should be
reviewed alongside the recorded provider, model, configuration, and source
snapshot.
Deterministic repository inventory. Scan files with stable ordering,
portable paths, hashes, language classification, ignore rules, and bounded
reads.
Reviewable selection. Choose exact files, directories, GitWildMatch
patterns, or line ranges—or ask a configured model for a bounded suggestion.
Local repository intelligence. Store immutable structural and optional
semantic index generations under .contextforge/index .
Portable artifacts. Export Markdown or JSON context packages, JSON task
handoffs, and compiled Markdown prompts.
Explicit trust boundaries. ContextForge does not edit repository source,
execute repository code, expose shell tools, or mutate Git state.
Automation-friendly output. Structured results stay on stdout while
progress and diagnostics stay on stderr.
flowchart LR
R["Repository"] --> S["scan / index"]
S --> D["task-aware discovery"]
D --> B["bounded selection"]
B --> P["context package"]
P --> A["external coding agent"]
Loading
A model is optional for scanning, trees, manual context packages, and
structural-only indexing:
contextforge scan .
contextforge index build . --provider none
contextforge context create . \
--include pyproject.toml \
--directory src/contextforge/context \
--exclude " **/__init__.py " \
--format markdown \
--output context.md
Installation
ContextForge requires Python 3.12 or newer. Install the published distribution:
python -m pip install contextforge-repo
For an isolated command-line installation, use either tool manager:
pipx install contextforge-repo
# or
uv tool install contextforge-repo
The PyPI distribution is named contextforge-repo ; the import package remains
contextforge , and the installed commands remain contextforge and ctxf .
The similarly named context-forge-cli distribution is a different,
unaffiliated project.
To install a checked-out source tree instead:
git clone https://github.com/waterflane/ContextForge.git
cd ContextForge
python -m venv .venv
source .venv/bin/activate
python -m pip install --upgrade pip
python -m pip install .
Windows PowerShell:
git clone https: // github.com / waterflane / ContextForge.git
Set-Location ContextForge
python - m venv .venv
.\.venv\Scripts\Activate.ps1
python - m pip install -- upgrade pip
python - m pip install .
The installation provides equivalent contextforge and ctxf console
commands. python -m contextforge is also supported.
Inspect a repository without writing ContextForge state:
contextforge scan .
contextforge tree . -- depth 2
contextforge context create . `
-- include ' pyproject.toml ' `
-- directory ' src/contextforge/context ' `
-- exclude ' **/__init__.py ' `
-- format json `
-- output ' context.json '
contextforge context inspect ' context.json '
Build a structural-only local index and inspect its status:
contextforge index build . -- provider none
contextforge index status .
Tip
Start with manual context creation when you already know the relevant files.
Use discovery when the task spans unfamiliar code and you have configured a
supported model provider.
create a compact review packet for an external coding agent;
map a repository without sending source to a model;
inspect stale, missing, or failed index records;
discover likely entry points, tests, configuration, and dependencies for a
task;
preserve a validated handoff that can be reviewed without the original
checkout;
benchmark discovery quality and repeatability against versioned manifests.
Every command supports --help ; run group help before using advanced or
mutating operations.
Global diagnostic options are --log-level , --log-format , --log-file ,
repeatable --log-component , --no-log-file , --no-color , and -v / -vv .
Detailed syntax, defaults, streams, side effects, mistakes, and examples are in
the Wiki CLI reference .
Project configuration is closed, versioned TOML. Resolution order is:
supported CONTEXTFORGE_* environment variable;
.contextforge/config.local.toml ;
.contextforge/config.toml , or an explicit --config PATH ;
The primary supported environment variables are:
CONTEXTFORGE_MODEL_CONTEXT_WINDOW ;
CONTEXTFORGE_MODEL_CONNECT_TIMEOUT ;
CONTEXTFORGE_MODEL_READ_TIMEOUT ;
CONTEXTFORGE_MODEL_OPERATION_TIMEOUT ;
CONTEXTFORGE_JSON_REPAIR_ATTEMPTS ;
CONTEXTFORGE_LOG_LEVEL , CONTEXTFORGE_LOG_FORMAT ,
CONTEXTFORGE_LOG_FILE , and CONTEXTFORGE_LOG_COMPONENTS .
The default provider is local Ollama at
http://127.0.0.1:11434/api/chat using model qwen2.5-coder:7b . Use
--provider none for structural-only indexing. The openai-compatible
provider and its lmstudio CLI alias require an exact model ID and a suitable
base_url .
Model-backed discovery requires the configured provider to be running with the
named model available. ContextForge's configured context_window must not
exceed the window actually loaded by that provider; inspect the resolved policy
before a long run with contextforge diagnostics provider PATH .
Credential configuration stores only the name of an environment variable in
credential_env ; the credential value is resolved at request time. See the
configuration guide and
Wiki configuration reference .
Fresh builds current structural evidence in memory and does not load
persisted semantic records or repository maps.
Indexed requires a readable active index and uses current indexed
structure, semantics, and maps.
Hybrid is the default. It starts with current index evidence, fills
structural gaps from the live snapshot, and explicitly falls back to fresh
structure when no valid index exists.
All successful selections are verified against current source identities.
Model-backed runs can produce different valid selections; ContextForge claims
deterministic rendering for the same validated result, not deterministic model
behavior. See Discovery output and benchmarks .
Manual packages use explicit selectors. With no include selector, all selectable
snapshot files are included up to the configured limits:
contextforge context create . --include README.md --format markdown
contextforge context create . --directory src --exclude " **/__init__.py "
contextforge context create . --glob " tests/test_*.py " --no-include-tree
contextforge context create . \
--include pyproject.toml \
--include-lines pyproject.toml:1-24 \
--format json
Automatic mode requires a non-empty task and does not accept manual directory,
glob, or line-range selectors:
contextforge context suggest . \
--task " Trace configuration precedence " \
--discovery hybrid \
--format markdown
contextforge context create . \
--task " Trace configuration precedence " \
--discovery hybrid \
--git-diff working \
--format json \
--output handoff.json \
--prompt-output prompt.md
context suggest does not write source or index state, but current diagnostics
policy may write a safe summary under .contextforge/runs . Output artifacts are
written atomically; existing destinations require --force where that option
is available.
benchmark discovery is experimental in 0.4.2 ; model-backed repeatability
measurements are observations for the recorded fixture state, not guarantees.
Start the configured provider first and use the same context-window value in
ContextForge and the provider runtime.
contextforge benchmark discovery ' C:\Repositories ' `
-- tasks ' .\benchmarks\discovery.json ' `
-- modes ' fresh,indexed,hybrid ' `
-- repeat 3 `
-- format json `
-- output ' .\benchmark-report.json '
The runner is repository/index read-only, disables configured file logging, and
records complete, failed, and cancelled runs in the result. Exit code 3 means
the command produced a complete benchmark report containing at least one task,
expectation, or budget failure. Do not discard stdout or the requested output
file when handling that code. Every run remains bounded by manifest limits,
provider retry limits, operation timeouts, and the configured context window.
trees: text , markdown , json ;
suggestions: text , compatibility alias table , markdown , json ;
context packages: markdown , json ;
index status and diagnostics: table , json ;
discovery benchmarks: text , markdown , json .
When no output path is supplied, the selected result is written to stdout.
Progress, logs, and errors use stderr, preserving parseable JSON stdout. Some
commands print a confirmation to stdout after writing a file; benchmark output
files are the exception and leave stdout empty.
Common process exit codes are 0 for success, 1 for operational failure, 2
for invalid usage or configuration, and 130 for cancellation. Exit code 3
has command-specific meaning: unreadable entries with scan --fail-on-error , or
a completed discovery benchmark with regression failures.
ContextForge is a typed Python modular monolith. Core application and domain
logic remain independent from Typer, FastAPI, model-provider implementations,
storage adapters, and future editor integrations. The scanner creates a
verified snapshot; intelligence extracts structural facts and optional semantic
interpretations; discovery selects bounded candidates; context and handoff
modules materialize portable artifacts; CLI, HTTP, and MCP are thin interfaces.
Read the architecture overview for dependency
boundaries and the security policy for tr

[truncated]

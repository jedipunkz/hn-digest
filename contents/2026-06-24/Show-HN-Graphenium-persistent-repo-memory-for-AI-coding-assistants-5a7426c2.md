---
source: "https://github.com/lambda-alpha-labs/Graphenium"
hn_url: "https://news.ycombinator.com/item?id=48658145"
title: "Show HN: Graphenium, persistent repo memory for AI coding assistants"
article_title: "GitHub - lambda-alpha-labs/Graphenium: Persistent structural memory for AI coding agents. Turns your repo into a fast, MCP-native knowledge graph so assistants stop grepping and start querying. · GitHub"
author: "Graphenium"
captured_at: "2026-06-24T12:01:18Z"
capture_tool: "hn-digest"
hn_id: 48658145
score: 1
comments: 0
posted_at: "2026-06-24T11:27:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Graphenium, persistent repo memory for AI coding assistants

- HN: [48658145](https://news.ycombinator.com/item?id=48658145)
- Source: [github.com](https://github.com/lambda-alpha-labs/Graphenium)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T11:27:02Z

## Translation

タイトル: Show HN: グラフェニウム、AI コーディング アシスタント用の永続リポジトリ メモリ
記事のタイトル: GitHub - lambda-alpha-labs/Graphenium: AI コーディング エージェントのための永続的な構造メモリ。リポジトリを高速な MCP ネイティブのナレッジ グラフに変えて、アシスタントが grep を停止してクエリを開始できるようにします。 · GitHub
説明: AI コーディング エージェントの永続的な構造メモリ。リポジトリを高速な MCP ネイティブのナレッジ グラフに変えて、アシスタントが grep を停止してクエリを開始できるようにします。 - ラムダアルファラボ/グラフェニウム

記事本文:
GitHub - lambda-alpha-labs/Graphenium: AI コーディング エージェント用の永続的な構造メモリ。リポジトリを高速な MCP ネイティブのナレッジ グラフに変えて、アシスタントが grep を停止してクエリを開始できるようにします。 · GitHub
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
ラムダアルファラボ
/
グラフェニウム
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 Commit 1 Commit .github .github contrib/ harness-adapter contrib/ harness-adapter docs docs scripts scripts skills/ graphenium skills/ graphenium src src worked worked .gitattributes .gitattributes .gitignore .gitignore .grapheniumignore .grapheniumignore AI_SETUP.md AI_SETUP.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントの永続的な構造メモリ。
Graphenium はリポジトリをクエリ可能なグラフに変えます。
他の MCP 互換アシスタントは、応答なしで約 20 ミリ秒でこれらに応答できます。
単一のファイルを読み取っています。大規模、マルチモジュール、または
grep とトレースによるナビゲーションが機能しない、見慣れないコードベース:
建築上のハブとは何ですか?
これらのコンポーネント間の最短パスは何ですか?
どのファイルが同じコミュニティに属していますか?
これは、ソースコードの理解ではなく、grep とトレースによるナビゲーションを置き換えます。
AI コーディング アシスタントはコードを読むのが得意ですが、リポジトリをナビゲートします
grep を使用する人間のように、シンボルを検索し、ファイルを開き、フォローします。
インポートし、さらにファイルを開き、関係を推測します。次に、もう一度すべてを実行します
次のセッション。
50 ファイルのプロジェクトでは grep が機能します。ディープインポートを備えた 5,000 ファイルのモノリポジトリ内
チェーン、それはありません。このワークフローには、次の 5 つの永続的な問題があります。
繰り返しのコールドスタート。すべての新しいセッションは耐久性のあるモデルなしで始まります
リポジトリの。
コンテキストウィンドウの圧力。生のソース ファイルはサイズが大きくなります。ナビゲーション
トクを消費する

推論に使用できる ens。
構造記憶がない。ファイルを読み取った後、アシスタントには何もありません
モジュール、機能、概念がどのように関係しているかを示す永続的なグラフ。
ファイル間の関係が失われています。 Grep はテキストの一致を表示しますが、一致しません
アーキテクチャ トポロジ、ハブ、コミュニティ、またはパス。
スケールが痛みを倍増させます。新しいファイルと依存関係があるたびに、
grep とトレースのループは遅くなり、コストが高くなります。グラフは速いままです
リポジトリのサイズに関係なく。
Graphenium は分析を 1 回実行し、結果をグラフとして保持して公開します。
MCP サーバー経由でアシスタントに送信します。グラフ
アシスタントのリポジトリの長期記憶になります。
分単位ではなく秒単位で方向を確認します。 Architecture_summary は、
アシスタントが単一のファイルを読み取る前の 30 秒間のコードベースのマップ。
コンテキストに焦点が当てられたままになります。コンテキスト ウィンドウを生の情報で埋めるのではなく、
ナビゲーション中のソース、アシスタントはコンパクトなグラフ出力について推論します
重要なファイルのみを読み取ります。
メモリはセッション後も残ります。グラフは持続します。新しい AI セッションが開始されます
前回のものと同じ構造知識を持っています。
何が得意なのか（そして何が苦手なのか）
大規模なコードベースをナビゲートする。 50 を超えるファイル リポジトリ、モノリポジトリ、または馴染みのないもの
プロジェクト、grep アンド トレースの無駄なコンテキスト。グラフがそれを置き換えます。
AI 支援のコード ナビゲーション: 構造的な質問に繰り返し答えることなく答えます
ファイルを読んでいます。
影響分析: 機能を変更する前に、接続されているノードを特定します。
クラス、またはモジュール。
オンボーディング: なじみのないリポジトリの高レベルのアーキテクチャ マップを迅速に取得します。
リファクタリング計画: ゴッド ノード、凝集性の低いコミュニティ、および
驚くべき国境を越えたエッジ。
コードレビュー: コードをレビューする前に、シンボル、次数、およびホットスポットを検査します。
変更されたファイル。
アクティブな開発中に監視モードでグラフを最新の状態に保ちます。
ソースコードを読んでいます。グラフは構造を捉えています

再や人間関係、
実装ロジックではありません。アシスタントは依然として実際のコードを読み取る必要があります。
全言語サーバー。完全な型チェックは実行しません。
LSP の深さでの言語固有のセマンティック分析。
実行時のトレース。これは静的分析とオプションの LLM 抽出です。それ
プログラムは実行しません。
セマンティック検索/埋め込み。 Graphenium はキーワードのスコアリングとグラフを使用します
ベクトルの類似性ではなく、トラバーサルです。
セキュリティスキャン。関係グラフは代わりにはなりません
専用の SAST ツール。
グラフェニウムなし:
grep → ファイルを読み取る → インポートをトレース → ファイルをさらに読む → アーキテクチャを推測する
グラフェニウムの場合:
query_graph → get_neighbors → Shortest_path → 適切なファイルのみを読み取ります
# プロジェクトのグラフを構築します (API キーは必要ありません)
ジムラン。 --no-semantic --no-viz
# 構造的な質問をする
gm query " build_from_extraction を呼び出すものは何ですか? "
# または、MCP 経由で AI アシスタントに接続し、直接質問します
クイックスタート
カール -fsSL https://raw.githubusercontent.com/lambda-alpha-labs/Graphenium/main/install.sh |しー
ソースから
Rust 1.75+ (rustup) が必要です。
git clone https://github.com/lambda-alpha-labs/Graphenium
cd グラフェニウム
カーゴインストール --path 。
バイナリは gm としてインストールされます。
# グラフを作成します (API キーは必要ありません)
ジムラン。 --no-semantic --no-viz
# クエリしてみる
gm query "認証ログインセッション" --budget 1000
# インストールを確認してください
GMドクター
地図と交通のオーバーレイ
グラフェニウムには 2 つの抽出モードがあります。どちらも便利です。彼らは異なるサービスを提供します
目的。
AST 専用モードでは、アシスタントにマップが提供されます。セマンティック モードではトラフィックが追加されます
オーバーレイ。
# AST のみ、ローカル、キーは不要
ジムラン。 --no-semantic --no-viz
# セマンティック: LLM 推論された関係を追加します
エクスポート ANTHROPIC_API_KEY=sk-ant-...
ジムラン。 --provider anthropic # また: openai、deepseek、openrouter
graph_stats ツールは常にレポートします

エッジ信頼度の内訳なので、
アシスタントはそれが何に取り組んでいるのかを知っています。
AI アシスタントの MCP 構成にグラフェニウムを追加します。サーバーは、
標準の MCP stdio トランスポート。または、 gm setup <target> を実行して構成を出力します。
あなたのアシスタントのために。
クロード デスクトップ (claude_desktop_config.json):
{
"mcpサーバー": {
"グラフェニウム" : {
"コマンド" : "gm" ,
"args" : [ "serve " 、 " --graph " 、 " /absolute/path/to/graphenium-out/graph.json " ]
}
}
}
カーソル ( ~/.cursor/mcp.json ):
{
"mcpサーバー": {
"グラフェニウム" : {
"コマンド" : "gm" ,
"args" : [ "serve " 、 " --graph " 、 " /absolute/path/to/graphenium-out/graph.json " ]
}
}
}
CodeWhale ( ~/.codewhale/mcp.json ):
{
「サーバー」: {
"グラフェニウム" : {
"コマンド" : " /absolute/path/to/gm " ,
"args" : [ "serve " 、 " --graph " 、 " /absolute/path/to/graphenium-out/graph.json " ]、
"環境" : {}
}
}
}
構成を更新した後、AI ツールを完全に終了して再起動します (Cmd+Q
macOS では、ウィンドウを閉じるだけではありません)。 MCP サーバーは起動時にのみロードされます。
リポジトリには、スキル/グラフェニウム/SKILL.md に AI スキルが同梱されており、
アシスタントはどのツールに手を伸ばすべきか、信頼水準をどのように解釈するか、そして
MCP が使用できないときに gm クエリにフォールバックする方法。
接続すると、アシスタントは 13 のグラフ ツールにアクセスできるようになります。
すべての書き込みは直ちにディスクに保存されます。
Graphenium はコードベースを 3 つのものとしてモデル化します。
ノードは、関数、メソッド、クラス、モジュールなどの意味のあるエンティティを表します。
構造、特性、ドキュメント、画像、アーキテクチャの概念。各ノード
メタデータを保持します: ラベル、修飾ラベル、ファイル タイプ、ソース ファイル、ソース
場所とコミュニティID。
エッジは、型付けされた有向関係です。
グラフェニウムはグラフを分析して、コミュニティ、ハブ ノード、最短距離を表面化します。
パス、コミュニティ間の驚くべきつながり、およびアーキテクチャ上の焦点パス。
アシスタントは方向を変えることができます

実装を読む前にそれ自体を構造的に
詳細。
すべてのエッジには信頼レベルが含まれます。
抽出されたエッジを事実として信頼します。
INFERRED エッジを強力なヒントとして使用します。
曖昧なエッジを検査対象のリードとして扱います。
実装を変更する前にソース コードを読んでください。
graph_stats は信頼度の内訳をレポートするため、グラフの種類がわかります。
あなたは一緒に働いています。
グラフェニウムはASTにツリーシッターを使用
9 言語にわたる抽出。
セマンティック抽出では、ドキュメント ( .md 、 .rst 、 .txt )、PDF、
そして画像。
必要な言語のみを使用してビルドします。
カーゴビルド --release --no-default-features --features lang-python、lang-rust
機能: lang-python 、 lang-js 、 lang-ts 、 lang-rust 、 lang-go 、
lang-java 、 lang-c 、 lang-cpp 、 lang-csharp 。
ディレクトリ上で完全な分析パイプラインを実行します。
gm run [パス] [オプション]
オプション
説明
パス
分析するディレクトリ (デフォルト: . )
--セマンティックなし
LLM 抽出をスキップします。 AST のみの結果を使用する
--no-viz
HTML生成をスキップする
--プロバイダー名
AI プロバイダー: anthropic (デフォルト)、openai 、openrouter 、deepseek 、openai 互換
--モデル名
使用するモデル (デフォルトはプロバイダー固有のデフォルト)
--api-key キー
API キー (プロバイダー固有の環境変数をオーバーライドします)
--api-base URL
openai互換プロバイダーのAPIベースURL
--モードディープ
積極的な LLM 推論
--更新
増分: 変更されたファイルのみを再抽出します
ジムラン。 --no-semantic --no-viz # AST のみの高速スキャン
ジムラン。 --provider openai # LLM セマンティック抽出を使用する
ジムラン。 --update # ファイル編集後の増分
GMクエリ
キーワードを使用して既存のグラフをクエリします。
gm query "<キーワード>" [オプション]
オプション
デフォルト
説明
--グラフのパス
グラフェニウム-out/graph.json
グラフファイルへのパス
--予算 N
2000年
出力トークンバジェット
--dfs
オフ
深さ優先検索を使用する
gmクエリ「認証ログインセッション」
gm クエリ " パーサー ast walker

" --dfs --budget 4000
GMサーブ
MCP サーバーを起動し、stdio 経由でグラフを公開します。
GMサーブ [オプション]
オプション
デフォルト
説明
--グラフのパス
グラフェニウム-out/graph.json
グラフファイルへのパス
GMの時計
ディレクトリを監視し、変更に応じてグラフを自動再構築します。
gm watch [パス] [オプション]
オプション
デフォルト
説明
パス
。
監視するディレクトリ
--デバウンスSECS
3.0
最後のイベントの後、再構築するまで待機します
--増分
本当の
変更されたファイルにパッチを適用します。完全な再構築の場合は false
GM腕時計。 --デバウンス 2.0
GMドクター
Graphenium インストールで診断チェックを実行します: バイナリの場所、グラフ
ファイルの健全性、ツリーシッター言語、API キー、グラフの品質。
gm ドクター [--グラフパス]
GM セットアップ
AI アシスタント用のすぐに貼り付けられる MCP 構成を印刷します。
gm setup <claude|cursor|codewhale> [--graph PATH]
GM セットアップ クロード
GM セットアップカーソル
gm セットアップ コードクジラ
出力ファイル
Graphenium は、分析されたディレクトリ内の Graphenium-out/ に出力を書き込みます。
ソース/
9 言語の抽出/ツリーシッター抽出
モデル/グラフ、ノード、エッジ、ハイパーエッジのタイプ
build/抽出結果からのグラフ構築
クラスター/ルーヴァン コミュニティの検出、結合、分割/フォーカス
検出/ファイル分類、機密ファイルのスキップ、コーパス警告
分析/ゴッドノード、驚くべき接続、AR

[切り捨てられた]

## Original Extract

Persistent structural memory for AI coding agents. Turns your repo into a fast, MCP-native knowledge graph so assistants stop grepping and start querying. - lambda-alpha-labs/Graphenium

GitHub - lambda-alpha-labs/Graphenium: Persistent structural memory for AI coding agents. Turns your repo into a fast, MCP-native knowledge graph so assistants stop grepping and start querying. · GitHub
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
lambda-alpha-labs
/
Graphenium
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github .github contrib/ harness-adapter contrib/ harness-adapter docs docs scripts scripts skills/ graphenium skills/ graphenium src src worked worked .gitattributes .gitattributes .gitignore .gitignore .grapheniumignore .grapheniumignore AI_SETUP.md AI_SETUP.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.sh install.sh View all files Repository files navigation
Persistent structural memory for AI coding agents.
Graphenium turns your repository into a queryable graph so Claude, Cursor,
and other MCP-compatible assistants can answer these in ~20 ms, without
reading a single file. Especially valuable in large, multi-module, or
unfamiliar codebases where grep-and-trace navigation breaks down:
What are the architectural hubs?
What is the shortest path between these components?
Which files belong to the same community?
It replaces grep-and-trace navigation, not source-code understanding.
AI coding assistants are good at reading code, but they navigate repositories
like a human using grep : search for a symbol, open the file, follow
imports, open more files, infer relationships. Then do it all again in the
next session.
In a 50-file project, grep works. In a 5,000-file monorepo with deep import
chains, it doesn't. That workflow has five persistent problems:
Repeated cold starts. Every new session begins without a durable model
of the repository.
Context window pressure. Raw source files are large; navigation
consumes tokens that could be used for reasoning.
No structural memory. After reading files, the assistant has no
persisted graph of how modules, functions, and concepts relate.
Missed cross-file relationships. Grep surfaces text matches, not
architectural topology, hubs, communities, or paths.
Scale multiplies the pain. Every new file and dependency makes the
grep-and-trace loop slower and more expensive. The graph stays fast
regardless of repo size.
Graphenium runs analysis once, persists the result as a graph, and exposes it
to assistants via an MCP server. The graph
becomes the assistant's long-term memory for your repository.
Orientation in seconds, not minutes. architecture_summary gives a
30-second map of the codebase before the assistant reads a single file.
Context stays focused. Instead of filling the context window with raw
source during navigation, the assistant reasons over compact graph output
and reads only the files that matter.
Memory survives sessions. The graph persists. A new AI session starts
with the same structural knowledge the last one had.
What it's good at (and what it's not)
Navigating large codebases. In 50+ file repos, monorepos, or unfamiliar
projects, grep-and-trace wastes context; the graph replaces it.
AI-assisted code navigation: answer structural questions without repeatedly
reading files.
Impact analysis: identify connected nodes before changing a function,
class, or module.
Onboarding: get a high-level architectural map of an unfamiliar repo fast.
Refactoring planning: find god nodes, low-cohesion communities, and
surprising cross-boundary edges.
Code review: inspect symbols, degrees, and hotspots before reviewing a
changed file.
Keeping the graph current with watch mode during active development.
Reading source code. The graph captures structure and relationships,
not implementation logic. An assistant still needs to read actual code.
A full language server. It does not perform complete type checking or
language-specific semantic analysis at LSP depth.
Runtime tracing. It is static analysis plus optional LLM extraction; it
does not execute the program.
Semantic search / embeddings. Graphenium uses keyword scoring and graph
traversal, not vector similarity.
Security scanning. Relationship graphs are not a substitute for
dedicated SAST tools.
Without Graphenium:
grep → read file → trace imports → read more files → infer architecture
With Graphenium:
query_graph → get_neighbors → shortest_path → read only the right files
# Build a graph for your project (no API key needed)
gm run . --no-semantic --no-viz
# Ask structural questions
gm query " what calls build_from_extraction? "
# Or connect an AI assistant via MCP and ask directly
Quick start
curl -fsSL https://raw.githubusercontent.com/lambda-alpha-labs/Graphenium/main/install.sh | sh
From source
Requires Rust 1.75+ ( rustup ).
git clone https://github.com/lambda-alpha-labs/Graphenium
cd Graphenium
cargo install --path .
The binary is installed as gm .
# Build a graph (no API key needed)
gm run . --no-semantic --no-viz
# Query it
gm query " authentication login session " --budget 1000
# Check your installation
gm doctor
Map vs traffic overlay
Graphenium has two extraction modes. Both are useful; they serve different
purposes.
AST-only mode gives the assistant a map. Semantic mode adds the traffic
overlay.
# AST-only, local, no key needed
gm run . --no-semantic --no-viz
# Semantic: adds LLM-inferred relationships
export ANTHROPIC_API_KEY=sk-ant-...
gm run . --provider anthropic # also: openai, deepseek, openrouter
The graph_stats tool always reports the edge confidence breakdown, so the
assistant knows what it's working with.
Add Graphenium to your AI assistant's MCP config. The server uses the
standard MCP stdio transport. Or run gm setup <target> to print the config
for your assistant.
Claude Desktop ( claude_desktop_config.json ):
{
"mcpServers" : {
"graphenium" : {
"command" : " gm " ,
"args" : [ " serve " , " --graph " , " /absolute/path/to/graphenium-out/graph.json " ]
}
}
}
Cursor ( ~/.cursor/mcp.json ):
{
"mcpServers" : {
"graphenium" : {
"command" : " gm " ,
"args" : [ " serve " , " --graph " , " /absolute/path/to/graphenium-out/graph.json " ]
}
}
}
CodeWhale ( ~/.codewhale/mcp.json ):
{
"servers" : {
"graphenium" : {
"command" : " /absolute/path/to/gm " ,
"args" : [ " serve " , " --graph " , " /absolute/path/to/graphenium-out/graph.json " ],
"env" : {}
}
}
}
After updating config, quit and relaunch the AI tool completely (Cmd+Q on
macOS, not just close the window). MCP servers are only loaded at startup.
The repo ships an AI Skill at skills/graphenium/SKILL.md that teaches
assistants which tool to reach for, how to interpret confidence levels, and
how to fall back to gm query when MCP is unavailable.
Once connected, the assistant has access to 13 graph tools.
All writes persist to disk immediately.
Graphenium models a codebase as three things.
Nodes represent meaningful entities: functions, methods, classes, modules,
structs, traits, documents, images, and architectural concepts. Each node
carries metadata: label, qualified label, file type, source file, source
location, and community ID.
Edges are typed, directed relationships.
Graphenium analyzes the graph to surface communities, hub nodes, shortest
paths, surprising cross-community connections, and architectural focus paths.
The assistant can orient itself structurally before reading implementation
details.
Every edge carries a confidence level.
Trust EXTRACTED edges as fact.
Use INFERRED edges as strong hints.
Treat AMBIGUOUS edges as leads to inspect.
Read source code before making implementation changes.
graph_stats reports the confidence breakdown so you know what kind of graph
you're working with.
Graphenium uses tree-sitter for AST
extraction across 9 languages.
Semantic extraction also processes documents ( .md , .rst , .txt ), PDFs,
and images.
Build with only the languages you need:
cargo build --release --no-default-features --features lang-python,lang-rust
Features: lang-python , lang-js , lang-ts , lang-rust , lang-go ,
lang-java , lang-c , lang-cpp , lang-csharp .
Run the full analysis pipeline on a directory.
gm run [PATH] [OPTIONS]
Option
Description
PATH
Directory to analyse (default: . )
--no-semantic
Skip LLM extraction; use AST-only results
--no-viz
Skip HTML generation
--provider NAME
AI provider: anthropic (default), openai , openrouter , deepseek , openai-compatible
--model NAME
Model to use (defaults to provider-specific default)
--api-key KEY
API key (overrides provider-specific env var)
--api-base URL
API base URL for openai-compatible provider
--mode deep
Aggressive LLM inference
--update
Incremental: only re-extract changed files
gm run . --no-semantic --no-viz # Fast AST-only scan
gm run . --provider openai # With LLM semantic extraction
gm run . --update # Incremental after editing files
gm query
Query an existing graph with keywords.
gm query "<keywords>" [OPTIONS]
Option
Default
Description
--graph PATH
graphenium-out/graph.json
Path to graph file
--budget N
2000
Output token budget
--dfs
off
Use depth-first search
gm query " authentication login session "
gm query " parser ast walker " --dfs --budget 4000
gm serve
Start an MCP server exposing the graph over stdio.
gm serve [OPTIONS]
Option
Default
Description
--graph PATH
graphenium-out/graph.json
Path to graph file
gm watch
Watch a directory and auto-rebuild the graph on changes.
gm watch [PATH] [OPTIONS]
Option
Default
Description
PATH
.
Directory to watch
--debounce SECS
3.0
Wait after last event before rebuild
--incremental
true
Patch changed files; false for full rebuild
gm watch . --debounce 2.0
gm doctor
Run diagnostic checks on your Graphenium installation: binary location, graph
file health, tree-sitter languages, API keys, and graph quality.
gm doctor [--graph PATH]
gm setup
Print ready-to-paste MCP config for an AI assistant.
gm setup <claude|cursor|codewhale> [--graph PATH]
gm setup claude
gm setup cursor
gm setup codewhale
Output files
Graphenium writes outputs to graphenium-out/ inside the analysed directory.
src/
extract/ tree-sitter extraction for 9 languages
model/ graph, node, edge, hyperedge types
build/ graph construction from extraction results
cluster/ Louvain community detection, cohesion, split/focus
detect/ file classification, sensitive-file skipping, corpus warnings
analyze/ god nodes, surprising connections, ar

[truncated]

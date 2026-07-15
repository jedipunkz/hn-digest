---
source: "https://github.com/Litenova-Solutions/Fuse"
hn_url: "https://news.ycombinator.com/item?id=48927999"
title: "Fuse – an open source MCP/CLI tool to speed up Claude Code on C# codebases"
article_title: "GitHub - Litenova-Solutions/Fuse: MCP server for .NET: typecheck AI agent edits against the compiler, resolve DI and route wiring from Roslyn, scope PRs to the files that matter. · GitHub"
author: "litenova"
captured_at: "2026-07-15T22:49:43Z"
capture_tool: "hn-digest"
hn_id: 48927999
score: 2
comments: 0
posted_at: "2026-07-15T22:30:13Z"
tags:
  - hacker-news
  - translated
---

# Fuse – an open source MCP/CLI tool to speed up Claude Code on C# codebases

- HN: [48927999](https://news.ycombinator.com/item?id=48927999)
- Source: [github.com](https://github.com/Litenova-Solutions/Fuse)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T22:30:13Z

## Translation

タイトル: Fuse – C# コードベース上のクロード コードを高速化するオープン ソース MCP/CLI ツール
記事のタイトル: GitHub - Litenova-Solutions/Fuse: .NET 用の MCP サーバー: AI エージェントの編集をコンパイラーに対してタイプチェックし、DI を解決して Roslyn からの配線をルーティングし、PR を重要なファイルにスコープします。 · GitHub
説明: .NET 用 MCP サーバー: コンパイラーに対して AI エージェント編集をタイプチェックし、DI を解決し、Roslyn からの配線をルーティングし、PR を重要なファイルにスコープします。 - Litenova-ソリューション/ヒューズ

記事本文:
GitHub - Litenova-Solutions/Fuse: .NET 用 MCP サーバー: AI エージェント編集をコンパイラーに対してタイプチェックし、DI を解決し、Roslyn からの配線をルーティングし、PR を重要なファイルにスコープします。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。リロしてください

このページに広告を掲載します。
Litenova-ソリューション
/
ヒューズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
926 コミット 926 コミット .github .github アセット アセット ベンチリリース ベンチリリース ビルド ビルド インストーラー インストーラー mcp-registry mcp-registry パッケージング パッケージ化 ロードマップ ロードマップ サイト サイト src src テスト テスト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md DCO.txt DCO.txt Directory.Build.props Directory.Build.props Directory.Packages.props Directory.Packages.props Fuse.slnx Fuse.slnx ライセンス ライセンス通知 通知README.md README.md SECURITY.md SECURITY.md Briefing.md Briefing.md global.json global.json install.bat install.bat すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Fuse は、永続的なセマンティック インデックス、型付きグラフの接続を備えたローカル .NET ツールです。
解像度、縮小されたタスクスコープのソース、および書き込み前のコンパイラ検証
コーディングエージェント。 MSBuild と Roslyn を通じてソリューションにインデックスを付け、結果を保存します
.fuse/fuse.db に保存され、エージェントのターンをまたいで再利用されます。
ファイルの読み取りとテキスト検索を繰り返すことにより、同じ構造が得られます。
.NET プロジェクト ディレクトリから:
dotnet ツールのインストール -g Fuse
ヒューズ mcp インストール --rules
MCP クライアントをリロードして、次のように尋ねます。
IOrderService をその実装に解決し、提案された OrderService.cs 編集を確認します。
書き込む前に、fuse_check を使用してください。
MCP サーバーが起動すると、共有ローカル デーモンが .fuse/fuse.db のウォーミングを開始します。
背景。コールド リードは、制限付き構文ファースト パスを待機し、制限付き構文ファースト パスが完了するとレポートします。
セマンティック グラフはまだアップグレード中です。同期が必要な場合は、fuse Index を実行します。

騒々しい
エージェントに接続する前にインデックスを作成します。また、fuse mcp install --rules は .fuse/ を追加します
プロジェクト スコープの .gitignore。
分析はローカルで実行され、オフラインでも作業できます。 Fuse は DI 登録の型付きグラフを調べます。
ハンドラー、ルート、および呼び出し元は、選択されたスコープの縮小されたソースを出力し、コーディングを可能にします。
エージェントのタイプチェックは、提案された単一ファイルのコンテンツを書き込む前にそれをチェックしました。モデルは必要ありません。の
オプションの更新チェックでは NuGet にアクセスでき、ビルドグレードの操作ではパッケージを使用できます。
リポジトリ用に構成されたフィード。
永続的なディスカバリーと削減されたコンテキスト
コーディング エージェントは、ファイルの読み取り、grep、正規表現を通じてリポジトリを検査できます。大きな
解決策として、これらの操作により、同じシンボル、参照、登録、および
数ターンにわたるプロジェクト構造。 Fuse は MSBuild を通じてその検出を実行し、
Roslyn は、結果を永続化し、ファイルの変更に応じて段階的にインデックスを再作成します。
プロジェクトがセマンティックにロードされると、グラフには DI 登録、リクエスト ハンドラー、
ルート、オプション バインディング、およびコール エッジ。そうでない場合、Fuse は構文レベルに戻ります。
そのプロジェクトのインデックスを作成し、モードを報告します。
配線を解決します。 furuse_find は、サービス、リクエスト、ルート、または構成をトレースします。
それを処理するコードのセクションを参照してください。テキスト検索で IOrderService が見つかります。ヒューズは次のとおりです
実行する実装への登録。
ブランチコンテキストをパックします。 fuse_review は git diff にシードし、関連する呼び出し元を返します。
ハンドラーと出所のあるテスト。記録された 69 件のプル リクエストの応答中央値は、
93.4% の精度で 1,026 個のトークン ( review.json )。
返すソースを減らします。 fuse_context は、トークン バジェット内で選択されたファイルを削減します
各ファイルが含まれた理由を記録します。 4 つの記録されたリポジトリにわたるスケルトン
削減により、トークンを保持しながら 38 ～ 44 パーセントが削除されました

すべての測定された一般大衆と
保護された型の名前 (reduce.json)。
暖かく読んでください。記録された NodaTime 実行 (セマンティック層、14,760 シンボル) では、正確に
シンボルの検索には中央値で 1.8 ミリ秒、タスクの位置特定には 15.7 ミリ秒、計画のレビューには 15.7 ミリ秒かかりました。
106.3 ミリ秒 (パフォーマンス.json、タイミングは環境によって異なります)。
コンパイラのチェックと変更の検証
furuse_check は 1 つのファイルの提案された内容をチェックし、コンパイラ診断を返します。
作業ツリーを変更せずに。 Oracle グレードは、
本物のビルド。ビルド グレードは、キャプチャ時に所有プロジェクトのスコープ指定されたドットネット ビルドを実行します。
状態は利用できません。サポートされている API 形状のエラーには修復パケットが含まれる場合があります。で
記録された実行、最上位の提案により、メンバーおよびタイプのニアミス エラー 20 件中 20 件が修復されました
( diagbench.json )。
パブリック メソッドを変更する前に、fuse_impact は呼び出し元、実装、および
参照型。パッケージ ID と 2 つのバージョンを指定すると、そのブレーク セットを返します。
そのNuGetアップグレード。
シグネチャを変更する必要がある場合、fuse_refactor はリファクタリングを diff としてステージングし、
コンパイラが新しい診断を報告しない場合にのみ、これを返します。
編集後、fuse_test は変更されたテスト タイプを選択して実行します。
スイート全体から始めるのではなく、シンボルを使用します。
すべての回答には、その製造方法が記載されています。 Fuse はこれを検証グレードと呼びます。
Oracle グレードは、実際のビルドからキャプチャされたコンパイル、ビルド グレードに対してチェックされます。
スコープ付き dotnet build を実行し、どちらのコンパイラ パスも応答できない場合、Fuse は棄権します。
そして、推測する代わりに、不足している前提条件に名前を付けます。
記録された結果の内容
以下のすべての結果はテスト/ベンチマーク/結果からのものであり、再現コマンドが含まれています
ベンチマークページにあります。
記録された OrderingApp テスト アプリ内の 1,000 件のコンパイラ ラベル付き編集 (500 件を超える、
500 ニュートラル)、ヒューズ_

壊れた編集が報告されたゼロをクリーンとして、拒否されたゼロが有効であることを確認してください
( checkgate.json ) を編集します。
同じアプリ内で、Fuse は、予期される 24 個の .NET 配線リンクすべてに一致し、追加の一致はありませんでした。
( semantics.json )。
69 件の実際のプル リクエストにわたって、ブランチ レビューはすべての git 変更ファイルを保持しました。
93.4% の精度で構築され、サイズの中央値は 1,026 トークン ( review.json )。
スコープが縮小されたエージェント ループの実行では、Fuse アームの編集がプロジェクト自体の編集に渡されました。
スコアリングされたロールアウトの 89% では最初の試行でテストを行ったのに対し、ロールアウトでは 82% でした。
重複する信頼区間を持つネイティブ ツール。ヒューズアームは成功を宣言しました
編集の失敗は 8 回でしたが、ネイティブの場合は 9 回でした。ビルドとテストの呼び出しは基本的に同等でした
3.1 と 3.2 (loop.json)。
4 つの記録されたリポジトリ全体で、スケルトンの削減により 38 ～ 44 パーセントが削除されました。
すべての公開型名と保護型名、および 96.3 ～ 99.4% を保持しながらトークンを作成します。
メソッド名 (reduce.json)。
オプトイン常駐ワークスペースは、繰り返し行われた fuse_check 呼び出しに 31.2 ミリ秒で応答しました。
記録された NodaTime 実行の中央値 ( Resident-latency.json )。そのパスは、
投機的チェックのための完全なビルド。マージ前の通常のビルドやテストは置き換えられません。
これらの測定には限界が定義されています。読んでください
比較する前の完全なメソッドと結果
ツールを使用するか、図を別のリポジトリに適用します。
Fuse は、ソース、コンパイラの状態、git メタデータ、およびローカルの .fuse/fuse.db インデックスを読み取ります。
読み取り、チェック、影響、リファクタリング、およびレビューの操作では、作業ツリーは書き込まれません。
action=apply を指定したfusion_workspaceは、明示的なツリー書き込みパスの1つであり、ドライです。
write=true でない限り実行します。
コンパイラによる配線解析は .NET のみです。他の言語は構文レベルの検索を受け取ります
そして削減。
関連するコードインテリジェンスツール
リポジトリ インデックス、コード グラフ、および l

anguage-server ツールはすでに存在します。
CodeGraphContext はローカルを提供します
多言語グラフ、セレナが公開
言語サーバーを使用したシンボル ツールとソースグラフ
マルチリポジトリ検索とコードインテリジェンスをカバーします。 Fuse はローカルの .NET 作業に集中します
MSBuild と Roslyn による: フレームワーク ワイヤリング、スコープ コンテキストの縮小、コンパイラ支援
提案されたファイルのチェック、変更の影響、およびカバーテストの選択。並走できる
コーディング クライアントに組み込まれたインデックスまたは検索。
ピア比較
期限付きの日付付き実験とそのサンプリング制限を記録します。一般的なランキングではありませんが、
コードインテリジェンスツール。
アパッチ2.0。著作権 (c) 2026 Litenova ソリューション。 「ライセンス」と「
注意。
.NET 用の MCP サーバー: AI エージェントの編集をコンパイラーに対してタイプチェックし、DI を解決して Roslyn からの配線をルーティングし、PR を重要なファイルにスコープします。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
13
ヒューズ v4.2.0
最新の
2026 年 7 月 15 日
+ 12 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for .NET: typecheck AI agent edits against the compiler, resolve DI and route wiring from Roslyn, scope PRs to the files that matter. - Litenova-Solutions/Fuse

GitHub - Litenova-Solutions/Fuse: MCP server for .NET: typecheck AI agent edits against the compiler, resolve DI and route wiring from Roslyn, scope PRs to the files that matter. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Litenova-Solutions
/
Fuse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
926 Commits 926 Commits .github .github assets assets bench-release bench-release build build installer installer mcp-registry mcp-registry packaging packaging roadmap roadmap site site src src tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md DCO.txt DCO.txt Directory.Build.props Directory.Build.props Directory.Packages.props Directory.Packages.props Fuse.slnx Fuse.slnx LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md briefing.md briefing.md global.json global.json install.bat install.bat View all files Repository files navigation
Fuse is a local .NET tool with a persistent semantic index, typed-graph wiring
resolution, reduced task-scoped source, and pre-write compiler verification for
coding agents. It indexes a solution through MSBuild and Roslyn, stores the result
in .fuse/fuse.db , and reuses it across agent turns instead of rediscovering the
same structure through repeated file reads and text searches.
From a .NET project directory:
dotnet tool install -g Fuse
fuse mcp install --rules
Reload your MCP client, then ask:
Resolve IOrderService to its implementation, then check the proposed OrderService.cs edit
with fuse_check before writing it.
When the MCP server starts, its shared local daemon begins warming .fuse/fuse.db in
the background. A cold read waits for a bounded syntax-first pass and reports when the
semantic graph is still upgrading. Run fuse index when you want a synchronous full
index before connecting the agent. fuse mcp install --rules also adds .fuse/ to
.gitignore at project scope.
Analysis runs locally and can work offline. Fuse walks a typed graph of DI registrations,
handlers, routes, and callers, emits reduced source for a selected scope, and lets a coding
agent typecheck proposed single-file content before writing it. No model is required. The
optional update check can contact NuGet, and build-grade operations can use the package
feeds configured for the repository.
Persistent Discovery and Reduced Context
Coding agents can inspect a repository through file reads, grep, and regex. On a large
solution, those operations can rediscover the same symbols, references, registrations, and
project structure across several turns. Fuse performs that discovery through MSBuild and
Roslyn, persists the result, and incrementally re-indexes files as they change.
When a project loads semantically, the graph records DI registrations, request handlers,
routes, options bindings, and call edges. When it does not, Fuse falls back to syntax-level
indexing for that project and reports the mode.
Resolve wiring. fuse_find traces a service, request, route, or configuration
section to the code that handles it. Text search finds IOrderService ; Fuse follows the
registration to the implementation that runs.
Pack branch context. fuse_review seeds on the git diff and returns related callers,
handlers, and tests with provenance. On 69 recorded pull requests the median response was
1,026 tokens at 93.4 percent precision ( review.json ).
Return less source. fuse_context reduces the selected files under a token budget
and records why each file was included. Across four recorded repositories, skeleton
reduction removed 38 to 44 percent of tokens while retaining every measured public and
protected type name ( reduce.json ).
Read warm. On the recorded NodaTime run (semantic tier, 14,760 symbols), exact
symbol lookup took 1.8 ms at the median, task localization 15.7 ms, and review planning
106.3 ms ( performance.json ; timings are environment-dependent).
Compiler Checks and Change Verification
fuse_check checks the proposed content of one file and returns compiler diagnostics
without changing the working tree. Oracle grade reuses compiler state captured from the
real build. Build grade runs a scoped dotnet build for the owning project when captured
state is unavailable. Supported API-shape errors can include a repair packet; in the
recorded run, the top suggestion repaired 20 of 20 near-miss member and type errors
( diagbench.json ).
Before changing a public method, fuse_impact finds callers, implementations, and
referencing types. Given a package id and two versions, it returns the break set for
that NuGet upgrade.
When a signature must change, fuse_refactor stages the refactor as a diff and
returns it only when the compiler reports no new diagnostic.
After an edit, fuse_test selects and runs the test types that reach the changed
symbol instead of starting with the whole suite.
Every answer names how it was produced. Fuse calls this the verification grade :
oracle grade checks against the compilation captured from the real build, build grade
runs a scoped dotnet build , and when neither compiler path can answer, Fuse abstains
and names the missing prerequisite instead of guessing.
What the Recorded Results Cover
Every result below comes from tests/benchmarks/results and has a reproduction command
on the benchmarks page .
Across 1,000 compiler-labeled edits in the recorded OrderingApp test app (500 breaking,
500 neutral), fuse_check reported zero broken edits as clean and rejected zero valid
edits ( checkgate.json ).
In the same app, Fuse matched all 24 expected .NET wiring links with no extra matches
( semantics.json ).
Across 69 real pull requests, branch review retained every git-changed file by
construction at 93.4 percent precision with a median size of 1,026 tokens ( review.json ).
In the reduced-scope agent-loop run, the Fuse arm's edits passed the project's own
tests on the first attempt in 89 percent of scored rollouts versus 82 percent for
native tools, with overlapping confidence intervals. The Fuse arm declared success on a
failing edit 8 times versus 9 for native. Build and test calls were essentially equal at
3.1 versus 3.2 ( loop.json ).
Across four recorded repositories, skeleton reduction removed 38 to 44 percent of
tokens while keeping every public and protected type name and 96.3 to 99.4 percent of
method names ( reduce.json ).
The opt-in resident workspace answered repeated fuse_check calls in 31.2 ms at the
median on the recorded NodaTime run ( resident-latency.json ). That path is faster than a
full build for speculative checks; it does not replace normal builds or tests before merge.
These measurements have defined limits. Read the
full methods and results before comparing
tools or applying the figures to another repository.
Fuse reads source, compiler state, git metadata, and its local .fuse/fuse.db index.
Read, check, impact, refactor, and review operations do not write the working tree.
fuse_workspace with action=apply is the one explicit tree-write path, and it is a dry
run unless write=true .
Compiler-backed wiring analysis is .NET-only. Other languages receive syntax-level search
and reduction.
Related Code-Intelligence Tools
Repository indexes, code graphs, and language-server tools already exist.
CodeGraphContext provides a local
multi-language graph, Serena exposes
language-server-backed symbol tools, and Sourcegraph
covers multi-repository search and code intelligence. Fuse concentrates on local .NET work
through MSBuild and Roslyn: framework wiring, reduced scoped context, compiler-backed
proposed-file checks, change impact, and covering-test selection. It can run alongside the
index or search built into a coding client.
The peer comparison
records a bounded, dated experiment and its sampling limits. It is not a general ranking of
code-intelligence tools.
Apache 2.0. Copyright (c) 2026 Litenova Solutions. See LICENSE and
NOTICE .
MCP server for .NET: typecheck AI agent edits against the compiler, resolve DI and route wiring from Roslyn, scope PRs to the files that matter.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
13
Fuse v4.2.0
Latest
Jul 15, 2026
+ 12 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

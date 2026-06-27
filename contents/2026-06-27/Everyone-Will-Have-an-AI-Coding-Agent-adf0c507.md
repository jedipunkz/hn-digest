---
source: "https://github.com/HoratiuCode/Clawie"
hn_url: "https://news.ycombinator.com/item?id=48697174"
title: "Everyone Will Have an AI Coding Agent"
article_title: "GitHub - HoratiuCode/Clawie: Clawie is a fast terminal coding agent built to write code well, stay easy to use, and run with serious local-model power when you want speed and control on your own machine. · GitHub"
author: "horatiucode"
captured_at: "2026-06-27T11:30:09Z"
capture_tool: "hn-digest"
hn_id: 48697174
score: 1
comments: 0
posted_at: "2026-06-27T11:04:39Z"
tags:
  - hacker-news
  - translated
---

# Everyone Will Have an AI Coding Agent

- HN: [48697174](https://news.ycombinator.com/item?id=48697174)
- Source: [github.com](https://github.com/HoratiuCode/Clawie)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T11:04:39Z

## Translation

タイトル: 誰もが AI コーディング エージェントを持つようになる
記事のタイトル: GitHub - HoratiuCode/Clawie: Clawie は、コードを適切に記述し、使いやすさを維持し、自分のマシンで速度と制御が必要な場合に本格的なローカル モデルのパワーで実行できるように構築された高速ターミナル コーディング エージェントです。 · GitHub
説明: Clawie は、コードを適切に記述し、使いやすさを維持し、自分のマシンで速度と制御が必要な場合に本格的なローカル モデルのパワーで実行できるように構築された高速ターミナル コーディング エージェントです。 - ホラティウコード/クローウィー

記事本文:
GitHub - HoratiuCode/Clawie: Clawie は、コードを適切に記述し、使いやすさを維持し、自分のマシンで速度と制御が必要な場合に本格的なローカル モデルのパワーで実行できるように構築された高速ターミナル コーディング エージェントです。 · GitHub
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
ホラティウコード
/
クローウィー
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
58 コミット 58 コミット .vscode .vscode python-clawie python-clawie rich-clawie rich-clawie スクリプト スクリプト .DS_Store .DS_Store .gitattributes .gitattributes README.md README.md ShrimpAIR.png ShrimpAIR.png clawie clawie clawie.png clawie.png clawiev2.png clawiev2.png Visual_agents_webui.png Visual_agents_webui.png web_ui_overview.png web_ui_overview.png すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Clawie は、プロジェクトの両側をまとめて管理するパッケージ化されたワークスペースです。
Rust-clawie (メイン CLI/ランタイム側)
python-clawie (Python ミラー/ワークスペース側)
このリポジトリが存在するため、完全なプロジェクトを 1 つのフォルダーから共有して実行できます。
環境変数、プロバイダー、モデル、API キー、およびベース URL を対話的に構成します。
./clawie セットアップ
設定は、Clawie config ディレクトリの下の settings.json に保存されます。
2. 怠惰なシニア開発モード (リーン モード)
Clawie は、過剰なエンジニアリングを防ぐためにリーン ラダーを強制します。
すでにこのコードベースに含まれていますか?
ネイティブ プラットフォーム機能で対応できますか?
依存関係をインストールすると解決しますか?
その後、最小限のコードを記述します。
このモードは REPL セッション内で直接管理します。
/lean [lite|full|ultra|off] : アクティブなリーン モードを切り替えるか表示します (デフォルトは full )。
/lean-review : 現在の差分をレビューしてオーバーエンジニアリングがないかどうかを確認します。
/lean-audit : リポジトリのオーバーエンジニアリングをスキャンします。
/lean-debt : clawie: 簡略化コメントを台帳に収集します。
/lean-gain : ベンチマークの影響メトリックを表示します。
/lean-help : コマンドリファレンスを印刷します。
REPL で /map (または /repo-map ) コマンドを使用すると、リポジトリのファイルと抽出されたシンボルのランク付けされたマップが生成され、大規模なコードベースのナビゲートに役立ちます。
マ

REPL から直接コミットを行います。
/commit : プリフライトは変更をチェックし、コミット メッセージを生成してコミットします。
/undo : 最後のコミットを元に戻します (ソフト リセット、変更を保持します)。
5. ワークスペース RAG サービス (claw-rag-service)
セマンティック リポジトリ検索のための SQLite ベースのベクトル インデックス サービス:
ファイルの取り込み: カーゴ run -p claw-rag-service -- ingest --workspace 。
API と UI を提供する : Cargo run -p claw-rag-service --serve
ローカル Clawie Web UI の高度なグラフィカル インターフェイス機能:
WebSocket ライブ ログ ストリーミング : 動的なリアルタイム実行ログ ストリーム。静的スナップショットを取得するのではなく、UI はバックグラウンド ソケット ( /ws-log ) に接続して、プロセス イベントの発生を監視します。
並べて視覚的に比較: 元のファイルとエージェントの改善点または現在の編集を比較します。 「差分を表示」をクリックすると、自動レイアウト調整により、視覚的に赤/緑の削除/追加が並べて表示されます。
Rust コードベースと Python ミラー間の同期とパリティをチェックします。
Sync Auditor CLI : ./scripts/check_rust_python_sync.py は、コマンド/ツール定義とファイル パリティを分析します。
単体テスト : test_rust_python_sync.py で定義されたテストは、継続的統合でチェックを実行します。
8. Pixel Agents ダッシュボード (ビジュアル インターフェイス)
アクティブなエージェントのインスタンスとステータスを表示する、ゲーム化されたリアルタイムのピクセルアート ダッシュボード:
ドラッグ可能なエージェント: CLI プロセスは、ビジュアル ルーム (机、コンピューター、本棚、サーバー ラックを完備) でアクティブなピクセル アート キャラクターとしてレンダリングされます。
セッション アクション : アクティブなエージェント セッションをビジュアル インターフェイスから直接終了します。
状態ビーコン : プロセスのステータス (思考中、実行中、アイドル、クローズ) を色分けされたステータス ライトによって動的に表示します。
./clawie セットアップ
Clawie エージェント REPL を起動します。
./クローウィー
焦点に応じて次のフォルダーで作業します。
CLI/ランタイム動作用のrust-clawie
パイ

Python 側のミラーリングされたモジュールとツール用の thon-clawie
クローウィ/
§── README.md
§── クローウィー
§── サビ・クローウィー/
└── python-clawie/
長いコーディング セッション (改善)
より長いセッションをより適切にサポートするために、ランタイムのデフォルトが増加されました。
max_budget_tokens : 12000 (低かった)
ターンループ --max-turns : デフォルト 12
これらは実行時に環境変数を使用して調整できます。
エクスポート CLAWIE_MAX_TURNS=120
エクスポート CLAWIE_MAX_BUDGET_TOKENS=30000
エクスポート CLAWIE_COMPACT_AFTER_TURNS=80
エクスポート CLAWIE_STRUCTURED_OUTPUT=false
エクスポート CLAWIE_STRUCTURED_RETRY_LIMIT=2
./クローウィー
注:
無効な値はデフォルトに戻ります。
数値は少なくとも 1 に固定されます。
# Python側のまとめ
python3 -m python-clawie.src.main の概要
# 明示的なターン数を指定してステートフル ループを実行する
python3 -m python-clawie.src.mainturn-loop「このモジュールを監査」 --max-turns 30
# 既存のセッションを再開する
python3 -m python-clawie.src.mainresume-session < session_id > " continue "
ユースケースとワークフロー
1. バックグラウンドでのエージェントのライブ監視 (WebSocket ログ)
複雑なエージェント タスク ( /lean-audit を使用したマルチターン監査セッションの実行など) を実行する場合、開発者は端末の横で Web UI を起動し、エージェントのアクションをライブで監視できます。
ワークフロー:
CLI セッションを開始します: ./clawie
CLI REPL で /webui を実行するか、./clawie --webui を実行して、Web UI を開きます。
Web UI ダッシュボード内のアクティブなルームのターミナル モニターをクリックして、ログ コンソールを開きます。
コンソールは WebSocket 経由で /ws-log?pid=<PID> に接続し、エージェントの実行中にプロセスのライフサイクル更新、コマンドの経過時間、および実行の詳細をリアルタイムでストリーミングします。
2. コード改善レビュー (並列差分)
Clawie がコード ファイルへの変更を提案すると、並べて表示される分割画面を使用して変更を確認、編集、適用できます。
ワークフロー:
クローウィに聞いてください

ファイルを改善するには: 「main.py を最適化」 (.improvements.md ファイルを生成します)。
Web UI を開き、ワークスペース エクスプローラーのサイドバーからファイルを選択します。
エディターの右上にある「差分を表示」ボタンをクリックします。
元のファイル (左ペイン、赤色の削除) と改善/編集済みファイル (右ペイン、緑色の追加) を比較します。
[Show Editor] でエディタに戻り、手動で調整を行い、[Save] をクリックします。
3. 継続的なパリティチェック (自動パイプライン)
機能ドリフトを作成せずに、rust-clawie 内の CLI ランタイム コマンド/ツールの更新が python-clawie 内に適切にミラーリングされるようにするには、次のようにします。
ワークフロー:
同期チェック CLI ツールを実行します: ./scripts/check_rust_python_sync.py
このツールは、Rust command.json のコマンドと Python のスナップショット、および Rust tools/src/lib.rs のツール仕様と Python のスナップショットを比較した詳細なパリティ レポートを出力します。
欠落しているファイルやコンテンツのドリフトがある場合、スクリプトはコード 1 で終了し、Git フックまたは CI パイプラインのバリデータとして機能します。
python3 -m Unittest python-clawie/tests/test_rust_python_sync.py を実行して、パッケージ構造の同期をアサートします。
Jameclaw : レガシー/オリジンの名前付けコンテキスト
以前の作業コピーは複数のローカル フォルダーに分割されていました。このパッケージでは、すべてが 1 つの Git 対応構造にまとめられているため、オンボーディング、開発、共有がより簡単になります。
Clawie は、コードを適切に記述し、使いやすさを維持し、自分のマシンで速度と制御が必要な場合に本格的なローカル モデルのパワーで実行できるように構築された高速ターミナル コーディング エージェントです。
jameclaw.xyz/clawie.html
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私のペルソナを共有しないでください

情報

## Original Extract

Clawie is a fast terminal coding agent built to write code well, stay easy to use, and run with serious local-model power when you want speed and control on your own machine. - HoratiuCode/Clawie

GitHub - HoratiuCode/Clawie: Clawie is a fast terminal coding agent built to write code well, stay easy to use, and run with serious local-model power when you want speed and control on your own machine. · GitHub
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
HoratiuCode
/
Clawie
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
58 Commits 58 Commits .vscode .vscode python-clawie python-clawie rust-clawie rust-clawie scripts scripts .DS_Store .DS_Store .gitattributes .gitattributes README.md README.md ShrimpAIR.png ShrimpAIR.png clawie clawie clawie.png clawie.png clawiev2.png clawiev2.png visual_agents_webui.png visual_agents_webui.png web_ui_overview.png web_ui_overview.png View all files Repository files navigation
Clawie is a packaged workspace that keeps both sides of the project together:
rust-clawie (main CLI/runtime side)
python-clawie (Python mirror/workspace side)
This repository exists so the full project is shareable and runnable from one folder.
Configure your environment variables, provider, model, API key, and base URL interactively:
./clawie setup
Settings are persisted in settings.json under the Clawie config directory.
2. Lazy Senior Dev Mode (Lean Mode)
Clawie enforces a Lean Ladder to prevent over-engineering:
Is it already in this codebase?
Does a native platform feature cover it?
Does an installed dependency solve it?
Only then write the minimum code.
Manage this mode directly inside the REPL session:
/lean [lite|full|ultra|off] : Switch or view the active lean mode (default is full ).
/lean-review : Review current diff for over-engineering.
/lean-audit : Scan repository for over-engineering.
/lean-debt : Harvest clawie: simplification comments into a ledger.
/lean-gain : Show benchmark impact metrics.
/lean-help : Print command reference.
Use the /map (or /repo-map ) command in the REPL to generate a ranked map of the repository's files and extracted symbols, helping navigate large codebases.
Manage commits directly from the REPL:
/commit : Preflight checks changes, generates a commit message, and commits them.
/undo : Undoes the last commit (soft reset, keeping changes).
5. Workspace RAG Service ( claw-rag-service )
SQLite-backed vector indexing service for semantic repository searches:
Ingest files : cargo run -p claw-rag-service -- ingest --workspace .
Serve API & UI : cargo run -p claw-rag-service -- serve
Advanced graphical interface features for the local Clawie Web UI:
WebSocket Live Log Streaming : Dynamic real-time execution log streams. Rather than pulling static snapshots, the UI connects to a background socket ( /ws-log ) to monitor process events as they happen.
Side-by-Side Visual Diffing : Compare original files vs agent improvements or current edits. Clicking "Show Diff" provides visual red/green deletions/additions side-by-side with automatic layout alignment.
Checks sync and parity between the Rust codebase and Python mirrors:
Sync Auditor CLI : ./scripts/check_rust_python_sync.py analyzes command/tool definitions and file parity.
Unit Testing : Tests defined in test_rust_python_sync.py run checks in continuous integration.
8. Pixel Agents Dashboard (Visual Interface)
A gamified, real-time pixel-art dashboard showing active agent instances and status:
Draggable Agents : CLI processes are rendered as active pixel-art characters in visual rooms (complete with desks, computers, bookshelves, and server racks).
Session Actions : Terminate active agent sessions directly from the visual interface.
State Beacons : Displays process statuses (thinking, executing, idle, closed) dynamically via color-coded status lights.
./clawie setup
Launch the Clawie agent REPL:
./clawie
Work in these folders depending on focus:
rust-clawie for CLI/runtime behavior
python-clawie for Python-side mirrored modules and tooling
Clawie/
├── README.md
├── clawie
├── rust-clawie/
└── python-clawie/
Long Coding Sessions (Improved)
The runtime defaults were increased to better support longer sessions:
max_budget_tokens : 12000 (was lower)
turn-loop --max-turns : default 12
You can tune these at runtime with environment variables:
export CLAWIE_MAX_TURNS=120
export CLAWIE_MAX_BUDGET_TOKENS=30000
export CLAWIE_COMPACT_AFTER_TURNS=80
export CLAWIE_STRUCTURED_OUTPUT=false
export CLAWIE_STRUCTURED_RETRY_LIMIT=2
./clawie
Notes:
Invalid values fall back to defaults.
Numeric values are clamped to at least 1 .
# Python-side summary
python3 -m python-clawie.src.main summary
# Run a stateful loop with explicit turn count
python3 -m python-clawie.src.main turn-loop " audit this module " --max-turns 30
# Resume an existing session
python3 -m python-clawie.src.main resume-session < session_id > " continue "
Use Cases & Workflows
1. Live Background Agent Monitoring (WebSocket Logs)
When running complex agent tasks (e.g. running a multi-turn audit session with /lean-audit ), developers can launch the Web UI alongside their terminal and watch the agent's actions live.
Workflow :
Start a CLI session: ./clawie
Open the Web UI by running /webui in the CLI REPL or by running ./clawie --webui
Click on the active room's terminal monitor inside the Web UI dashboard to open the Log Console.
The console connects via WebSocket to /ws-log?pid=<PID> and streams process lifecycle updates, command elapsed times, and execution details in real time as the agent runs.
2. Code Improvement Review (Side-by-Side Diff)
When Clawie suggests changes to code files, you can review, edit, and apply them using the side-by-side split screen.
Workflow :
Ask Clawie to improve a file: "Optimize main.py" (which generates a .improvements.md file).
Open the Web UI and select the file from the workspace explorer sidebar.
Click the "Show Diff" button at the top right of the editor.
Compare the Original File (left pane, red deletions) and the Improvements / Edited (right pane, green additions).
Switch back to the editor with "Show Editor" to make manual refinements, then click "Save" .
3. Continuous Parity Checking (Automated Pipelines)
To ensure that CLI runtime command/tool updates inside rust-clawie are mirrored properly inside python-clawie without creating feature drift:
Workflow :
Run the sync check CLI tool: ./scripts/check_rust_python_sync.py
The tool outputs a detailed parity report comparing commands in Rust commands.json vs Python's snapshot, and tool specifications in Rust tools/src/lib.rs vs Python's snapshot.
If there are missing files or content drifts, the script exits with code 1 , serving as a validator in Git hooks or CI pipelines.
Run python3 -m unittest python-clawie/tests/test_rust_python_sync.py to assert package structure sync.
Jameclaw : legacy/origin naming context
Earlier working copies were split across multiple local folders. This package keeps everything in one Git-ready structure so onboarding, development, and sharing are simpler.
Clawie is a fast terminal coding agent built to write code well, stay easy to use, and run with serious local-model power when you want speed and control on your own machine.
jameclaw.xyz/clawie.html
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

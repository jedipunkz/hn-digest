---
source: "https://github.com/ronle/clayrune"
hn_url: "https://news.ycombinator.com/item?id=49170443"
title: "Show HN: Clayrune – Run Claude Code agents in parallel without losing context"
article_title: "GitHub - ronle/clayrune: Open-source mission control for running multiple Claude Code agents across projects - one dashboard, a scheduler, cross-session memory, and browser/phone access. Runs locally on your own Claude subscription (the CLI, not the API). MIT, local-first. · GitHub"
author: "Clayrune"
captured_at: "2026-08-04T16:08:20Z"
capture_tool: "hn-digest"
hn_id: 49170443
score: 1
comments: 0
posted_at: "2026-08-04T15:35:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Clayrune – Run Claude Code agents in parallel without losing context

- HN: [49170443](https://news.ycombinator.com/item?id=49170443)
- Source: [github.com](https://github.com/ronle/clayrune)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:35:06Z

## Translation

タイトル: Show HN: Clayrune – コンテキストを失わずにクロード コード エージェントを並行して実行する
記事のタイトル: GitHub - ronle/clayrune: プロジェクト全体で複数のクロード コード エージェントを実行するためのオープンソース ミッション コントロール - 1 つのダッシュボード、スケジューラー、クロスセッション メモリ、ブラウザ/電話アクセス。独自の Claude サブスクリプション (API ではなく CLI) でローカルに実行されます。 MIT、ローカルファースト。 · GitHub
説明: プロジェクト全体で複数の Claude Code エージェントを実行するためのオープンソースのミッション コントロール (1 つのダッシュボード、スケジューラー、クロスセッション メモリ、ブラウザ/電話アクセス)。独自の Claude サブスクリプション (API ではなく CLI) でローカルに実行されます。 MIT、ローカルファースト。 - ロンル/クレイルーン

記事本文:
GitHub - ronle/clayrune: プロジェクト全体で複数の Claude Code エージェントを実行するためのオープンソースのミッション コントロール (1 つのダッシュボード、スケジューラー、クロスセッション メモリ、ブラウザ/電話アクセス)。独自の Claude サブスクリプション (API ではなく CLI) でローカルに実行されます。 MIT、ローカルファースト。 · GitHub
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
ロンレ
/
クレイルーン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
937 コミット 937 コミット

.github/ ワークフロー .github/ ワークフロー _scratch _scratch アセット アセット ビーコン ビーコン control_plane control_plane データ データ デモエクスポート デモエクスポート ドキュメント ドキュメント インストーラー インストーラー マーケティング マーケティング mc mc mc_remote mc_remote mc_remote_iface mc_remote_iface mc_tty_shim mc_tty_shim mc_tunnel mc_tunnel 静的 静的 スチュワード スチュワード テスト テスト ツール ツール .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md Agent_runtime.py Agent_runtime.py Agent_worktree.py Agent_worktree.py app.py app.py build-macos.spec build-macos.spec db.py db.py distiller.py distiller.py github_sync.py github_sync.py Marketing_preview.py Marketing_preview.py mcp.py mcp.py mcp_installer.py mcp_installer.py package.json package.json preflight.py preflight.py project_sync.py project_sync.py pyrightconfig.json pyrightconfig.json pytest.ini pytest.ini 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt サーバー.py サーバー.py スキル.py スキル.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード エージェントのミッション コントロール。
1 つのコンソールまたは携帯電話から、すべてのプロジェクトにわたって多くのエージェントを実行、スケジュールし、ベビーシッターを行います。
Claude Code は優れたエージェントですが、それは 1 つのターミナル、1 つのリポジトリ上にある 1 つのエージェントです。
そしてセッションが終了すると、知っていたものはすべて消えてしまいます。 Clayrune は一番上のレイヤーです。
すべてのプロジェクトがタイルであり、すべてのエージェントが実行され、作業が継続されるダッシュボード
キーボードから離れている間に移動する。
1 人のエージェントではなくフリートです。すべてのプロジェクトが 1 つのグリッド上にあり、エージェントは並行して派遣され監視されます。
携帯電話から制御します。どこからでもタスクをディスパッチし、ストリームを監視し、サーバーを再起動します。
それはあなたなしで実行されます - エージェントをスケジュールするか、常駐チャーターを渡してそのまま実行してください

kは無人です。
セッション間のメモリを記憶するので、エージェントがすべてのタスクをコールドで開始することはありません。
API ではなく、Claude サブスクリプション — Clayrune は、すでに実行されている Claude Code CLI を駆動するため、エージェントの作業コストは現在とまったく同じです。追加のトークン請求、メーター制、当社経由のルーティングは一切ありません。
それはあなたのものです - ローカルファースト、自分の鍵を持参し、すべて開いてください。クラウドはオプションです。
Clayrune はあなたのマシン上で実行されます。マシンがスリープ状態になると、エージェントと Clayrune もスリープ状態になります。
それがあったと伝えます。エージェントが動作している間は起動したままにする設定があります。
▶ ライブデモをお試しください — インストールなし · または ⬇ Windows の場合 — macOS + Linux + 以下のソースからダウンロードします。
「これってクロード・コードじゃないの？」
Clayrune は Claude Code を置き換えるのではなく、それを実行します。クロード・コードがエンジンです。クレイルーン
コックピットです。複数のプロジェクト、複数のプロジェクトがある瞬間に、それに到達します。
エージェント、または座って見守りたくない仕事。
ステータス追跡、説明、ドメインを使用して複数のプロジェクトを管理
あらゆるプロジェクトにエージェントを派遣 — Claude Code のファーストクラスに加え、Gemini、Codex、Aider、OpenCode、Goose
同じプロジェクト内で以前のチャット セッションを再利用して、全体的なトークンの使用量を削減します。
ストリーミング出力を介してエージェントのアクティビティをリアルタイムで監視し、ライブの「思考と書き込み」のアクティビティ状態を監視します
エージェントの質問にインラインで回答します。エージェントはオプションから選択するよう求めることができ、誰も見ていない場合は、設定したチャネルを通じて質問が届きます。
クロスセッションメモリ — エージェントは各タスクをコールドで開始するのではなく、コンテキストを転送します。
エージェントを無人で実行します — スケジューラー (1 回 / 毎日 / 間隔 / cron) と、常駐チャーターを渡す自律的なスチュワード
Hivemind を使用して 1 つの目標に向けて複数のエージェントを調整し、クロード コード ワークフローの進行状況をライブで監視します
スキルと MCP サーバーの管理 (

専用サーフェスからの Anthropic 形式のスキル、stdio/HTTP/SSE MCP)
優先順位、ドラッグ アンド ドロップの順序付け、添付ファイルを使用してプロジェクトのバックログを維持します
受信トレイ / 「お待ちしています」により、実際に決定が必要な実行が表示されます
1 回、毎日、または間隔ベースのトリガーを使用して自動タスクをスケジュールする
GitHub の問題とバックログを同期する — gh CLI を介した双方向同期
時間範囲フィルタリングを使用して、すべてのセッションにわたるトークンの使用量とコストを追跡します
複数のプロジェクト ウィンドウを同時に開きます (マルチモーダル ウィンドウ システム) - 開いている会話とそのレイアウトは、ページを更新して再起動しても保持されます。
Clayrune.io トンネル経由で任意のデバイスから管理 — プッシュで電話アプリとしてインストール可能。さらに、変更を展開した後、携帯電話からワンクリックでサーバーを再起動できます。
専用のワイドフォーマットビューアでエージェントの計画を表示
まずは試してみてください - Clayrune.io/demo にあるインタラクティブで完全にシミュレートされたデモ
すべてのプロジェクト間でベースライン ルールを共有する (SHARED_RULES.md)
プロジェクトタイルをグリッド上に自由に配置（Androidホーム画面風）
初回実行のウォークスルー ツアーでは、インターフェイスを通じて新規ユーザーをガイドします
PowerShell を開き (「スタート」→「PowerShell」と入力→「Enter」)、1 行を貼り付けて Enter キーを押します。
iwr https://clayrune.io/install.ps1 - useb |アイエックス
Claude CLI をインストールし、Clayrune を複製し、依存関係をインストールし、
デスクトップ/スタートメニューのショートカット。最初の起動ではダッシュボードが開きます
http://localhost:5199 。
ダブルクリックの方がいいですか? Clayrune.io にはインストーラー .exe があります —
署名されていないため、Windows は「認識されないアプリ」の通知を 1 回表示します ( 詳細 →
とにかく実行してください）。同じインストールで、1 回限りのプロンプトだけが表示されます。
代わりにソースから実行したいですか?以下のソースからの実行を参照してください。
前提条件 (ソースのみ)
ビルド済みの exe ではなくソースから実行したい場合:
Claude CLI — インストールガイド
git clone https

://github.com/ronle/clayrune.git
CD ミッションコントロール
pip install -r 要件.txt
Python app.py
ダッシュボードを含むネイティブ ウィンドウを開きます。 Flask サーバーはバックグラウンドで実行されます。
git clone https://github.com/ronle/clayrune.git
CD ミッションコントロール
pip インストールフラスコ
Pythonサーバー.py
# ブラウザで http://localhost:5199 を開きます
最初の実行: アプリ内の設定パネルでポート、プロジェクトを順を追って説明します。
ディレクトリ、エージェント モデル、その他の設定 - セットアップ スクリプトは必要ありません。
クローンを使用しないガイド付きインストールをご希望ですか?ワンコマンドインストーラーを使用します。
Clayrune.io (Windows では PowerShell、macOS では 1 行 /
Linux)、セットアップ全体を自動で実行します。
最初の実行時に、config.json ファイルがデフォルトで作成されます。
{
「ポート」: 5199 、
"shared_rules_path" : " data/SHARED_RULES.md " ,
"projects_base" : " /home/you " 、
"エージェントモデル": " " ,
"agent_max_turns" : 0 、
"use_streaming_agent" : false ,
"ユーザー名" : " " ,
"エージェント名" : " "
}
設定
説明
デフォルト
ポート
サーバーポート
5199
共有ルールのパス
共有ルール ファイルへのパス (すべてのエージェント プロンプトに挿入)
データ/SHARED_RULES.md
プロジェクトベース
プロジェクト パス検証用のベース ディレクトリ
ユーザーのホームディレクトリ
エージェントモデル
すべてのプロジェクトのデフォルトのクロード モデル
"" (CLI のデフォルト)
エージェント最大ターン数
セッションごとのエージェントの最大ターン数 (0 = 無制限)
0
use_streaming_agent
モード B 永続エージェント プロセスを有効にする
偽
ユーザー名
あなたの名前 (エージェントのコンテキストに表示されます)
「」
エージェント名
エージェントの表示名
「」
環境変数 MC_PORT=8080 python server.py を使用してポートを設定することもできます。
ステータスインジケーター (アクティブ、待機中、ブロック中、パーク中) を含むすべてのプロジェクトのタイルベースの概要
グリッドにスナップしたタイル配置 - タイルを任意の位置にドラッグしたり、ギャップを残したり、タイルを入れ替えたりできます。
カスタマイズ可能な色によるドメインの分類
プロジェクトごとのアクセントカラーのテーマ
すべてのプロジェクトにわたるアクティビティ ストリーム
コンパ

ct ボタンでグリッドのギャップを削除します
ワンクリックでタスクを Claude Code エージェントにディスパッチします
構文強調表示付きのリアルタイム ストリーミング出力
実行中のエージェントまたはアイドル状態のエージェントにフォローアップ メッセージを送信する
プロジェクトあたりの複数の同時エージェント セッション
スクリーンショットをエージェントのプロンプトに直接貼り付けます
プロジェクトごとのモデル選択 (Sonnet 4.5、Opus 4.6、Haiku 4.5)
エージェントが AskUserQuestion を呼び出すときのインタラクティブな質問フォーム
ExitPlanMode でスタックしたエージェントの計画承認ボタン
セッションごとのトークンの使用量とコストの追跡
実行中のセッションのライブ経過タイマー
モード A (デフォルト): ターンごとに新しいクロード プロセスを生成します。フォローアップキューと自動ディスパッチ。
モード B ( use_streaming_agent: true ): --input-format stream-json を使用した永続プロセス。フォローアップは応答を高速化するために標準入力に直接書き込みます。プロセスはターン間で存続します。
スケジュールに従ってエージェントの派遣を自動化する
3 つのスケジュール タイプ: 1 回 (特定の日時)、毎日 (時刻 + 曜日)、間隔 (N 分ごと)
個々のスケジュールを有効/無効にする
今後のジョブのバナーには、次の 5 つのスケジュールされたタスクが表示されます
プロジェクトごとのタスクのバックログと優先順位 (低/通常/高/重大)
ドラッグアンドドロップによる添付ファイルのアップロード
バックログアイテムをエージェントセッションに直接ディスパッチします
GitHub の問題の同期 (双方向)
任意のプロジェクトを GitHub リポジトリ ( owner/repo ) に接続します。
バックログ項目は GitHub の課題と双方向で同期します
優先ラベルは自動的にマッピングされます (高/中/低)
開閉状態を同期
手動同期ボタンまたは 5 分ごとの自動同期
認証には gh CLI が必要です
エージェントの計画は専用の幅の広いウィンドウで開き、読みやすくなります。
プランモード出力の自動検出
「計画履歴」タブには、プロジェクトごとのすべての過去の計画が表示されます
エージェントの出力を大きなウィンドウで表示するためのポップアウト ボタン
Claude Code のネイティブを使用したプロジェクトごとの永続メモリ

メモリー.md
メモリの内容がエージェント コンテキストに自動的に挿入される
自動メモリ: セッションの概要が完了時に追加されます
Clayrune と Claude CLI の直接使用間で共有されます。
複数のプロジェクトモーダルを同時に開く
ウィンドウのドラッグ、サイズ変更、最小化、復元
キーボード ナビゲーション (エスケープして閉じる)
画面下部の最小化されたトレイ
タッチサポート: モバイルデバイスでのドラッグ、サイズ変更、ピンチズーム
ヘッダー内のグローバル トークン カウンター (入力/出力トークン + USD コスト)
クリックして時間範囲でフィルタリングします: 全期間、今日、今週、今月
完了後のセッションごとのトークンバッジ
新規ユーザー向けの 18 ステップのガイド付きツアー
最初の実行時に自動トリガー (ゼロプロジェクト)
ヘッダーの「ツアー」ボタンから再トリガー可能
ツアー中に表示される仮想デモ タイルとモーダル
ミッションコントロール/
app.py デスクトップ エントリ ポイント (pywebview + Flask)
server.py Flask バックエンド (API + 静的サービス)
github_sync.py GitHub 問題同期モジュール
静的/
Index.html シングルページ アプリ (HTML + CSS + JS、ビルドステップなし)
データ/
プロジェクト/プロジェクト JSON ファイル (自動作成)
アップロード/添付ファイル
config.json ユーザー構成 (自動作成、gitignored)
installer/ ホストされた 1 コマンド インストーラー (clayrune.io)
バックエンド: 構成可能なポート上の Python Flask サーバー (デフォルトは 5199)
フロントエンド : バニラ HTML/CSS/JS シングルページ アプリ (フレームワーなし)

[切り捨てられた]

## Original Extract

Open-source mission control for running multiple Claude Code agents across projects - one dashboard, a scheduler, cross-session memory, and browser/phone access. Runs locally on your own Claude subscription (the CLI, not the API). MIT, local-first. - ronle/clayrune

GitHub - ronle/clayrune: Open-source mission control for running multiple Claude Code agents across projects - one dashboard, a scheduler, cross-session memory, and browser/phone access. Runs locally on your own Claude subscription (the CLI, not the API). MIT, local-first. · GitHub
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
ronle
/
clayrune
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
937 Commits 937 Commits .github/ workflows .github/ workflows _scratch _scratch assets assets beacon beacon control_plane control_plane data data demo-export demo-export docs docs installer installer marketing marketing mc mc mc_remote mc_remote mc_remote_iface mc_remote_iface mc_tty_shim mc_tty_shim mc_tunnel mc_tunnel static static steward steward tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md agent_runtime.py agent_runtime.py agent_worktree.py agent_worktree.py app.py app.py build-macos.spec build-macos.spec db.py db.py distiller.py distiller.py github_sync.py github_sync.py marketing_preview.py marketing_preview.py mcp.py mcp.py mcp_installer.py mcp_installer.py package.json package.json preflight.py preflight.py project_sync.py project_sync.py pyrightconfig.json pyrightconfig.json pytest.ini pytest.ini requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt server.py server.py skills.py skills.py View all files Repository files navigation
Mission control for your Claude Code agents.
Run, schedule, and babysit many agents across every project — from one console, or from your phone.
Claude Code is a brilliant agent — but it's one agent, in one terminal, on one repo,
and everything it knew is gone when the session ends. Clayrune is the layer on top: a
dashboard where every project is a tile, every agent is running, and the work keeps
moving while you're away from the keyboard.
A fleet, not one agent — every project on one grid, agents dispatched and monitored side by side.
Control it from your phone — dispatch a task, watch it stream, restart the server, from anywhere.
It runs without you — schedule agents, or hand one a standing charter and let it work unattended.
It remembers — cross-session memory, so your agents don't start every task cold.
Your Claude subscription, not the API — Clayrune drives the Claude Code CLI you already run, so agent work costs exactly what it does today. No extra token bills, nothing metered, nothing routed through us.
It's yours — local-first, bring your own key, all of it open. The cloud is optional.
Clayrune runs on your machine. If the machine sleeps, so does your agent — and Clayrune
tells you it did. There's a setting to keep it awake while an agent is working.
▶ Try the live demo — no install · or ⬇ Download for Windows — macOS + Linux + from source below.
"Isn't this just Claude Code?"
Clayrune doesn't replace Claude Code — it runs it. Claude Code is the engine; Clayrune
is the cockpit. You reach for it the moment you have more than one project, more than one
agent, or work you'd rather not sit and watch.
Manage multiple projects with status tracking, descriptions, and domains
Dispatch agents across any project — Claude Code first-class, plus Gemini, Codex, Aider, OpenCode, and Goose
Reuse previous chat sessions within the same project to reduce overall token usage
Monitor agent activity in real-time via streaming output, with live "thinking vs. writing" activity states
Answer agents' questions inline — an agent can ask you to choose between options, and if nobody's watching it reaches you over your configured channel
Cross-session memory — agents carry context forward instead of starting each task cold
Run agents unattended — a Scheduler (once / daily / interval / cron) and an autonomous Steward you hand a standing charter
Orchestrate multiple agents on one goal with Hivemind , and watch live Claude Code Workflow progress
Manage Skills and MCP servers (Anthropic-format skills, stdio/HTTP/SSE MCP) from dedicated surfaces
Maintain project backlogs with priorities, drag-and-drop ordering, and file attachments
Inbox / "Waiting on you" surfaces the runs that actually need a decision
Schedule automated tasks with once, daily, or interval-based triggers
Sync backlogs with GitHub Issues — bidirectional sync via gh CLI
Track token usage and costs across all sessions with time-range filtering
Open multiple project windows simultaneously (multi-modal windowing system) — open conversations and their layouts persist across page refresh and reboot
Manage from any device via the clayrune.io tunnel — installable as a phone app with push, plus a one-click server restart from your phone after deploying changes
View agent plans in a dedicated wide-format viewer
Try it first — an interactive, fully-simulated demo at clayrune.io/demo
Share baseline rules across all projects (SHARED_RULES.md)
Arrange project tiles freely on a grid (Android home screen style)
First-run walkthrough tour guides new users through the interface
Open PowerShell (Start → type PowerShell → Enter), paste one line, press Enter:
iwr https: // clayrune.io / install.ps1 - useb | iex
It installs the Claude CLI, clones Clayrune, installs dependencies, and drops a
Desktop/Start Menu shortcut. First launch opens the dashboard at
http://localhost:5199 .
Prefer a double-click? There's an installer .exe on clayrune.io —
it's unsigned, so Windows shows an "unrecognized app" notice once ( More info →
Run anyway ). Same install, just the one-time prompt.
Prefer running from source instead? See Running from Source below.
Prerequisites (from source only)
If you prefer running from source instead of the prebuilt exe:
Claude CLI — Installation guide
git clone https://github.com/ronle/clayrune.git
cd mission-control
pip install -r requirements.txt
python app.py
Opens a native window with the dashboard. Flask server runs in the background.
git clone https://github.com/ronle/clayrune.git
cd mission-control
pip install flask
python server.py
# Open http://localhost:5199 in your browser
First run: the in-app Settings panel walks you through the port, project
directory, agent model, and other settings — no setup script needed.
Prefer a guided, no-clone install? Use the one-command installer at
clayrune.io (PowerShell on Windows, one line on macOS /
Linux), which drives the whole setup for you.
On first run, a config.json file is created with defaults:
{
"port" : 5199 ,
"shared_rules_path" : " data/SHARED_RULES.md " ,
"projects_base" : " /home/you " ,
"agent_model" : " " ,
"agent_max_turns" : 0 ,
"use_streaming_agent" : false ,
"user_name" : " " ,
"agent_name" : " "
}
Setting
Description
Default
port
Server port
5199
shared_rules_path
Path to shared rules file (injected into all agent prompts)
data/SHARED_RULES.md
projects_base
Base directory for project path validation
User home directory
agent_model
Default Claude model for all projects
"" (CLI default)
agent_max_turns
Max agent turns per session (0 = unlimited)
0
use_streaming_agent
Enable Mode B persistent agent process
false
user_name
Your name (shown in agent context)
""
agent_name
Agent display name
""
You can also set the port via environment variable: MC_PORT=8080 python server.py
Tile-based overview of all projects with status indicators (Active, Waiting, Blocked, Parked)
Snap-to-grid tile arrangement — drag tiles to any position, leave gaps, swap tiles
Domain categorization with customizable colors
Per-project accent color theming
Activity stream across all projects
Compact button to remove grid gaps
Dispatch tasks to Claude Code agents with one click
Real-time streaming output with syntax highlighting
Send follow-up messages to running or idle agents
Multiple concurrent agent sessions per project
Paste screenshots directly into agent prompts
Per-project model selection (Sonnet 4.5, Opus 4.6, Haiku 4.5)
Interactive question forms when agents call AskUserQuestion
Plan approval button for agents stuck in ExitPlanMode
Token usage and cost tracking per session
Live elapsed timer for running sessions
Mode A (default): Spawns a new claude process per turn. Follow-ups queue and auto-dispatch.
Mode B ( use_streaming_agent: true ): Persistent process with --input-format stream-json . Follow-ups write directly to stdin for faster responses. Process stays alive between turns.
Automate agent dispatch on a schedule
Three schedule types: Once (specific datetime), Daily (time + day-of-week), Interval (every N minutes)
Enable/disable individual schedules
Upcoming jobs banner shows next 5 scheduled tasks
Per-project task backlog with priorities (low/normal/high/critical)
File attachments with drag-and-drop upload
Dispatch backlog items directly to an agent session
GitHub Issues sync (bidirectional)
Connect any project to a GitHub repository ( owner/repo )
Backlog items sync with GitHub Issues in both directions
Priority labels mapped automatically (high/medium/low)
Open/closed status synchronized
Manual sync button or automatic sync every 5 minutes
Requires gh CLI for authentication
Agent plans open in a dedicated wider window for easier reading
Auto-detection of plan mode output
Plans History tab shows all historical plans per project
Pop-out button for viewing any agent output in a larger window
Per-project persistent memory using Claude Code's native MEMORY.md
Memory content injected into agent context automatically
Auto-memory: session summaries appended on completion
Shared between Clayrune and direct Claude CLI usage
Open multiple project modals simultaneously
Drag, resize, minimize, and restore windows
Keyboard navigation (Escape to close)
Minimized tray at bottom of screen
Touch support: drag, resize, and pinch-to-zoom on mobile devices
Global token counter in header (input/output tokens + USD cost)
Click to filter by time range: All Time, Today, This Week, This Month
Per-session token badge after completion
18-step guided tour for new users
Auto-triggers on first run (zero projects)
Re-triggerable via "Tour" button in header
Virtual demo tile and modal shown during tour
mission-control/
app.py Desktop entry point (pywebview + Flask)
server.py Flask backend (API + static serving)
github_sync.py GitHub Issues sync module
static/
index.html Single-page app (HTML + CSS + JS, no build step)
data/
projects/ Project JSON files (auto-created)
uploads/ File attachments
config.json User configuration (auto-created, gitignored)
installer/ Hosted one-command installer (clayrune.io)
Backend : Python Flask server on configurable port (default 5199)
Frontend : Vanilla HTML/CSS/JS single-page app (no framewor

[truncated]

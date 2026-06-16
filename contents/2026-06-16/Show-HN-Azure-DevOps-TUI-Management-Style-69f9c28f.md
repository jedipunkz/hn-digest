---
source: "https://github.com/elpulgo/azdo"
hn_url: "https://news.ycombinator.com/item?id=48560026"
title: "Show HN: Azure DevOps TUI Management Style"
article_title: "GitHub - Elpulgo/azdo: A TUI for working with Azure DevOps in ze terminal. · GitHub"
author: "elpulgo"
captured_at: "2026-06-16T19:19:06Z"
capture_tool: "hn-digest"
hn_id: 48560026
score: 2
comments: 0
posted_at: "2026-06-16T18:44:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Azure DevOps TUI Management Style

- HN: [48560026](https://news.ycombinator.com/item?id=48560026)
- Source: [github.com](https://github.com/elpulgo/azdo)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T18:44:47Z

## Translation

タイトル: HN の表示: Azure DevOps TUI 管理スタイル
記事のタイトル: GitHub - Elpulgo/azdo: ZE ターミナルで Azure DevOps を操作するための TUI。 · GitHub
説明: ZE ターミナルで Azure DevOps を操作するための TUI。 - エルプルゴ/アズド
HN テキスト: 管理メトリックを備えた Azure Devops TUI ツールは、スタックしたアイテム、チーム メンバーごとの多すぎる Wip、およびスプリント全体の追跡を追跡します。

記事本文:
GitHub - Elpulgo/azdo: ze ターミナルで Azure DevOps を操作するための TUI。 · GitHub
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
エルプルゴ
/
アズド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode 「その他のアクション」メニューを開く フォルダーとファイル
31 コミット 31 コミット .beads .beads .claude/ コマンド .claude/ コマンド .github/ workflows .github/ workflows cmd/ azdo-tui cmd/ azdo-tui 内部 内部スクリーンショット スクリーンショット スクリプト スクリプト .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md Architecture.md Architecture.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md FAQ.md FAQ.md ライセンス ライセンス README.md README.md RELEEASES.md RELEASES.md Skill.md Skill.md config.yaml.example config.yaml.example example-theme.json example-theme.json go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh install_test.sh install_test.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Azure DevOps のターミナル ユーザー インターフェイス (TUI) - プル リクエスト、作業項目、パイプラインをターミナルから直接管理します。
カール -fsSL https://raw.githubusercontent.com/Elpulgo/azdo/main/install.sh |しー
Windows (PowerShell):
irm https://raw.githubusercontent.com/Elpulgo/azdo/main/install.ps1 |アイエックス
インストール スクリプトは自動的に次のことを行います。
OS とアーキテクチャを検出する
GitHub から最新リリースをダウンロードする
バイナリを適切な場所にインストールします
プレースホルダー値を含む構成ファイルを作成する
# 特定のバージョンをインストールする
カール -fsSL https://raw.githubusercontent.com/Elpulgo/azdo/main/install.sh | sh -s -- --バージョン v0.1.0
# カスタム ディレクトリにインストールする
./install.sh --install-dir ~ /bin
マニュアルのダウンロード
「リリース」ページから、使用しているプラットフォーム用の最新リリースをダウンロードします。
アーカイブを抽出し、バイナリを PATH 内のディレクトリに移動します。
git クローン https://github.com/Elpulgo/azdo.git
CD アズド
go build -o azdo-tui ./cmd/azdo-tui
Go インストールの使用
github.com/Elpulgo/azdo/cmd/azdo-tui@latest をインストールしてください
特長
プルリクエスト

uests (タブ 1): プル リクエストの表示と追跡
作業項目 (タブ 2): 作業項目の参照と管理
パイプライン (タブ 3): パイプライン実行の監視とドリルダウン
1 、 2 、 3 キーまたは ← / → 矢印キーを使用してタブを切り替えます。
ステータスインジケーター付きのプルリクエストのリストビュー
自分が作成した PR ( m キー) または自分がレビュー担当者である PR ( A キー) のみを表示するフィルター
PR情報とメタデータを表示する詳細ビュー
詳細ビューから PR に直接投票 (承認、拒否、提案、待機、リセット)
コードレビュー: ファイルごとのナビゲーションを備えた差分ビューア
インラインコメント、スレッド返信、スレッド解決
一般的な (ファイル固有ではない) コメント
ステータスとタイプの情報を含む作業項目のリスト ビュー
作業項目の詳細を示す詳細ビュー
説明の下にあるディスカッション (コメント) を最新のものから順に表示します
詳細ビューからコメントを追加します ( c キー、複数行フォーム)
詳細ビューから作業項目の状態を直接変更します (利用可能な状態を動的に取得します)。
割り当てられたアイテムのみを表示するようにフィルターします
並べ替え可能なテーブルで最近のパイプライン実行を表示する
色分けされたステータス インジケーター (✓ 成功、✗ 失敗、● 実行中、○ キュー中)
構成可能なポーリング間隔によるライブ自動更新
フッターの接続ステータスインジケーター
ステージ、ジョブ、タスクを含む階層的な詳細ビュー
各ステップの所要時間の追跡
スクロール可能なビューポートを備えた完全なログ ビューア
チーム リーダーの管理ビュー。デフォルトでは無効になっています。 config.yaml の metrics.enabled: true で有効にすると、4 番目のタブが表示されます。
v で切り替えられる 2 つのサブビュー:
ライブ — 作業項目ごとの現在の状態の滞留、ユーザーごとのロールアップ (WIP、処理中、最も古いアクティブ/テスト準備完了、構成された間隔で閉じられたポイント)、および最悪から順に表示される「スタック項目」ペイン。ライブ作業項目フェッチからソースされます。ローカル状態はなく、オンデマンド更新のみです。
トレンド — スプリント・オン・スプリン

ローカルの 90 日間のスナップショット ファイルからの比較。 T を使用してスプリント タグの任意の組み合わせを選択し (複数選択、スペースで切り替え、入力で確認)、ユーザーごとのクローズ ポイント、平均 WIP、スタック数、およびサイクル タイムを並べて確認します。値は色で表示されます。閉じたポイントは緑、過負荷の場合は黄色、スタックしたアイテムは赤です。
スナップショット ファイルは ~/.config/azdo-tui/metrics.jsonl にあります。毎日最初のメトリクス タブの起動時に作業項目ごとに 1 日あたり 1 行が追加され、その後 90 日のウィンドウにプルーニングされます。データベースなし - 追加専用の JSONL。
ワンショット埋め戻し (オプション)。新規インストールは空のスナップショット ファイルから開始されるため、トレンド ビューには最初の ~ 2 つのスプリントに対して「スナップショット履歴が不十分です」と表示されます。チームの実際の最近の履歴からファイルをシードするには、次のように設定します。
メトリクス:
run_one_shot_backfill : true
次回の起動時に、このタブは、すべての構成済みプロジェクトにわたって進行中のすべての作業項目または最近閉じられた作業項目を調べ、 /updates 経由で各項目のリビジョン履歴を読み取り、90 日前に遡って毎日のスナップショット行を合成します。フッターには進捗状況と結果が報告されます。マーカー ファイル ( ~/.config/azdo-tui/.metrics-backfill-done ) は再実行を防止します。再シードしたい場合は削除してください。完了したらフラグを false に戻し、フッター ヒントが表示されなくなるようにします。
初回実行時のセットアップ ウィザードによる構成手順のガイド
すべてのキーボード ショートカットのヘルプ モーダル (? を押す)
システムキーリングを使用した安全な PAT ストレージ
コンテキストを認識したキーバインドのヒント
自動再試行による適切なエラー処理
True Color をサポートする 8 つの組み込みテーマ
テーマスイッチャーモーダル ( t を押す) でテーマをその場で変更できます
表示名のカスタマイズによるマルチプロジェクトのサポート
状態の永続性 — 最後にアクティブなタブと最後に開いた PR / 作業項目の詳細をセッション間で記憶するため、中断したところから再開できます
アズドなしで試してみたい

Azure DevOps アカウントをお持ちではありませんか?デモを実行します。構成、PAT、セットアップは必要ありません。
アズドデモ
これにより、現実的なモック データ (2 つの架空のプロジェクト、差分付きのプル リクエスト、作業項目、ログ付きのパイプライン実行) を含む完全な TUI が起動します。すべての機能が動作し、ナビゲート、詳細の表示、テーマの切り替え、UI の探索が可能です。ツールの評価やスクリーンショットの撮影に最適です。
スクリーンショット フォルダー内のその他のスクリーンショットをご覧ください。
# TUIを開始する
アズド
# モックデータで試してみる(設定不要)
アズドデモ
# パーソナルアクセストークンを設定または更新します
アズド認証
# バージョンを表示
azdo --バージョン
# ヘルプを表示
azdo --ヘルプ
構成
初めて azdo を実行するときは、ウィザード セットアップがこれをセットアップするのに役立ちます。
それ以外の場合は、次の指示に従ってください。
次の場所に構成ファイルを作成します。
Linux/macOS : ~/.config/azdo-tui/config.yaml
Windows : C:\Users\<ユーザー名>\.config\azdo-tui\config.yaml
# Azure DevOps 組織名 (必須)
組織 : 組織名
# Azure DevOps プロジェクト名 (必須)
# 単純な形式:
プロジェクト:
- あなたのプロジェクト名
# 表示名の場合 (UI に表示されるフレンドリ名):
プロジェクト数:
# - 名前: 醜い API プロジェクト名
# 表示名: 私のプロジェクト
# - 名前: 醜い API プロジェクト名-2
# 表示名: 私のプロジェクト 2
# ポーリング間隔 (秒単位) (オプション、デフォルト: 60)
ポーリング間隔 : 60
# テーマ (オプション、デフォルト: ダーク)
# 利用可能なテーマ: ダーク、gruvbox、nord、dracula、catppuccin、github、retro、monokai
テーマ：ダーク
# 特定のペインを無効にする (オプション、カンマ区切り)
# 有効な値: パイプライン、ワークアイテム
#disabled_panes: パイプライン、作業項目
# メトリクス ダッシュボード (オプトイン、管理機能)。有効にしない限り非表示になります。
# 完全なリファレンスについては、以下の「メトリクスの構成」を参照してください。
# メトリクス:
# 有効: false
# interval_days: 14 # ライブのウィンドウが閉じられました p

「ts」欄
# active_stale_days: 3 # この上にアクティブに滞在すると、アイテムにフラグが立てられます
# rft_stale_days: 2 # この上にある「Ready for Test」に留まると項目にフラグが立てられます
# wip_limit: 4 # 飛行中、これを厳密に超えるとユーザーがオーバーロードされたことをマークします
# run_one_shot_backfill: false # ワンタイム /アップデート シード (「機能」→「メトリクス」を参照)
# states: # ボードの実際の状態名 (大文字と小文字は区別されません)
# アクティブ: アクティブ
#ready_for_test: テストの準備完了
# 閉店: 閉店
# state_labels: # オプションの列ヘッダーのオーバーライド (省略された場合は自動導出)
# アクティブ: アクティブ
# テスト準備完了: rft
# 閉店: 閉店
構成オプション:
組織 : Azure DevOps 組織名 (必須)
プロジェクト : Azure DevOps プロジェクト名のリスト (必須)。各エントリは、プレーン文字列、または name および display_name フィールドを持つオブジェクトにすることができます。 display_name は TUI に表示されますが、名前は API 呼び出しに使用されます。
polling_interval : データを更新する頻度 (秒単位) (オプション、デフォルト: 60)
テーマ : UI のカラーテーマ (オプション、デフォルト: ダーク)
disabled_panes : 非表示にするペインのカンマ区切りのリスト (オプション)。有効な値: Pipelines 、 workitems 。ペインが無効になると、そのタブ、キーボード ショートカット、および関連するすべての UI が削除されます。プル リクエストを無効にすることはできません。
メトリクス : オプトイン管理ダッシュボード。完全なリファレンスについては、以下の「メトリクス設定」を参照してください。また、その機能については、「機能」→「メトリクス ダッシュボード」を参照してください。
dark - 青とシアンのアクセントのあるダークテーマ
gruvbox - レトロなグルーヴの配色
nord - 北極、北青みがかったカラー パレット
ドラキュラ - 紫とピンクのアクセントが付いたデフォルトのダークテーマ
catpuccin - 心地よいパステルテーマ (モカバージョン)
レトロ - 黒地にマトリックスにインスピレーションを得た緑色の蛍光体
monokai - クラシックな Monokai の配色
メトリクス ダッシュボードはオプトインされており、 metrics.enabled: true でない限り完全に非表示になります。すべてのキーは の下に存在します

トップレベルのメトリクス: ブロックおよびオプションです。キーが省略された場合は、以下のデフォルトが適用されます。検証ルールは、 Enabled が true の場合にのみ適用されます。
これらは、ボードの実際のワークフロー状態文字列を、メトリクス エンジンが追跡する 3 つのバケットにマッピングします。照合では大文字と小文字が区別されず、空白がトリミングされますが、それぞれの名前は単一の名前になります (カンマ区切りのエイリアスはありません)。ボードが文字通り「アクティブ」/「テスト準備完了」/「クローズ」を使用していない場合は、これらを設定しないと、メトリクス タブには何もバケットが設定されません。
各名前は空ではなく、他の 2 つとは区別され、一重引用符 ( ' ) を含まない必要があります。一重引用符は WIQL インジェクションの安全性のために拒否されます。
メトリクス テーブルの列ヘッダーの表示専用オーバーライド。省略した場合、ラベルは設定された状態名から自動的に導出されます。複数の単語の名前は小文字のイニシャルになり ( Ready for Test → rft 、 In Progress → ip )、単一単語の名前はそのまま小文字になります ( Done → done )。
例 - 「Doing」/「QA」/「Done」を使用するボード:
メトリクス:
有効 : true
間隔日 : 14
wip_limit : 4
状態:
アクティブ : 実行中
テスト準備完了: QA
終了 : 完了
状態ラベル:
Ready_for_test : QA # 自動生成された小文字の「qa」をオーバーライドして大文字を保持します
カスタムテーマ
JSON を配置して独自のカスタム テーマを作成できます

[切り捨てられた]

## Original Extract

A TUI for working with Azure DevOps in ze terminal. - Elpulgo/azdo

An Azure devops TUI tool, with management metrics, to follow up on stuck items, too much wip per team member and tracking across sprints.

GitHub - Elpulgo/azdo: A TUI for working with Azure DevOps in ze terminal. · GitHub
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
Elpulgo
/
azdo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits .beads .beads .claude/ commands .claude/ commands .github/ workflows .github/ workflows cmd/ azdo-tui cmd/ azdo-tui internal internal screenshots screenshots scripts scripts .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md Architecture.md Architecture.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md FAQ.md FAQ.md LICENSE LICENSE README.md README.md RELEASES.md RELEASES.md Skill.md Skill.md config.yaml.example config.yaml.example example-theme.json example-theme.json go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh install_test.sh install_test.sh View all files Repository files navigation
A Terminal User Interface (TUI) for Azure DevOps - manage pull requests, work items, and pipelines directly from your terminal.
curl -fsSL https://raw.githubusercontent.com/Elpulgo/azdo/main/install.sh | sh
Windows (PowerShell):
irm https: // raw.githubusercontent.com / Elpulgo / azdo / main / install.ps1 | iex
The install scripts will automatically:
Detect your OS and architecture
Download the latest release from GitHub
Install the binary to the appropriate location
Create a config file with placeholder values
# Install a specific version
curl -fsSL https://raw.githubusercontent.com/Elpulgo/azdo/main/install.sh | sh -s -- --version v0.1.0
# Install to a custom directory
./install.sh --install-dir ~ /bin
Manual Download
Download the latest release for your platform from the Releases page .
Extract the archive and move the binary to a directory in your PATH .
git clone https://github.com/Elpulgo/azdo.git
cd azdo
go build -o azdo-tui ./cmd/azdo-tui
Using Go Install
go install github.com/Elpulgo/azdo/cmd/azdo-tui@latest
Features
Pull Requests (Tab 1): View and track pull requests
Work Items (Tab 2): Browse and manage work items
Pipelines (Tab 3): Monitor and drill into pipeline runs
Switch between tabs using 1 , 2 , 3 keys or ← / → arrow keys
List view of pull requests with status indicators
Filter to show only your created PRs ( m key) or PRs where you're a reviewer ( A key)
Detailed view showing PR information and metadata
Vote on PRs directly from the detail view (approve, reject, suggestions, wait, reset)
Code review : Diff viewer with file-by-file navigation
Inline commenting, thread replies, and thread resolution
General (non-file-specific) comments
List view of work items with status and type information
Detailed view showing work item details
View the Discussion (comments) below the description, newest first
Add comments from the detail view ( c key, multi-line form)
Change work item state directly from the detail view (dynamically fetches available states)
Filter to show only your assigned items
View recent pipeline runs in a sortable table
Color-coded status indicators (✓ Success, ✗ Failed, ● Running, ○ Queued)
Live auto-refresh with configurable polling interval
Connection status indicator in footer
Hierarchical detail view with stages, jobs, and tasks
Duration tracking for each step
Full log viewer with scrollable viewport
A management view for team leads, disabled by default . Enable it via metrics.enabled: true in config.yaml and a fourth tab appears.
Two sub-views, toggled with v :
Live — current-state dwell per work item, per-user roll-up (WIP, in-flight, oldest Active / Ready for Test, points closed in the configured interval), and a worst-first "stuck items" pane. Sourced from the live work-item fetch — no local state, on-demand refresh only.
Trends — sprint-on-sprint comparison from a local 90-day snapshot file. Pick any combination of sprint tags with T (multi-select; space toggles, enter confirms) and see per-user points closed , average WIP , stuck count , and cycle time side-by-side. Values are colored: green for closed points, yellow when overloaded, red for stuck items.
The snapshot file lives at ~/.config/azdo-tui/metrics.jsonl . One row per work item per day is appended on first metrics-tab launch each day, then pruned to a 90-day window. No database — append-only JSONL.
One-shot backfill (optional). A fresh install starts with an empty snapshot file, so the Trends view shows "Insufficient snapshot history" for the first ~2 sprints. To seed the file from your team's actual recent history, set:
metrics :
run_one_shot_backfill : true
On the next launch the tab walks every in-flight or recently-closed work item across all configured projects, reads each item's revision history via /updates , and synthesizes daily snapshot rows back 90 days. The footer reports progress and the result. A marker file ( ~/.config/azdo-tui/.metrics-backfill-done ) prevents re-running — delete it if you want to re-seed. Flip the flag back to false once it's done so the footer hint stops appearing.
Setup wizard on first run guides you through configuration
Help modal with all keyboard shortcuts (press ? )
Secure PAT storage using system keyring
Context-aware keybinding hints
Graceful error handling with automatic retry
Eight built-in themes with true color support
Theme switcher modal (press t ) to change themes on the fly
Multi-project support with display name customization
State persistence — remembers the last active tab and the last opened PR / work item detail across sessions, so you can pick up where you left off
Want to try azdo without an Azure DevOps account? Run the demo — no configuration, no PAT, no setup required:
azdo demo
This launches the full TUI with realistic mock data (two fictional projects, pull requests with diffs, work items, pipeline runs with logs). All features work — you can navigate, view details, switch themes, and explore the UI. Perfect for evaluating the tool or taking screenshots.
See more screenshots in the screenshots folder.
# Start the TUI
azdo
# Try it out with mock data (no setup needed)
azdo demo
# Set or update your Personal Access Token
azdo auth
# Show version
azdo --version
# Show help
azdo --help
Configuration
When running azdo for the first time, a wizard setup will help you setup this.
Otherwise follow these instructions.
Create a configuration file at the following location:
Linux/macOS : ~/.config/azdo-tui/config.yaml
Windows : C:\Users\<username>\.config\azdo-tui\config.yaml
# Azure DevOps organization name (required)
organization : your-org-name
# Azure DevOps project name(s) (required)
# Simple format:
projects :
- your-project-name
# With display names (friendly name shown in UI):
# projects:
# - name: ugly-api-project-name
# display_name: My Project
# - name: ugly-api-project-name-2
# display_name: My Project 2
# Polling interval in seconds (optional, default: 60)
polling_interval : 60
# Theme (optional, default: dark)
# Available themes: dark, gruvbox, nord, dracula, catppuccin, github, retro, monokai
theme : dark
# Disable specific panes (optional, comma-separated)
# Valid values: pipelines, workitems
# disabled_panes: pipelines,workitems
# Metrics dashboard (opt-in, management feature). Hidden unless enabled.
# See "Metrics Configuration" below for the full reference.
# metrics:
# enabled: false
# interval_days: 14 # window for the Live "closed pts" column
# active_stale_days: 3 # dwell in Active above this flags the item
# rft_stale_days: 2 # dwell in Ready for Test above this flags the item
# wip_limit: 4 # in-flight strictly above this marks a user overloaded
# run_one_shot_backfill: false # one-time /updates seed (see Features → Metrics)
# states: # your board's actual state names (case-insensitive)
# active: Active
# ready_for_test: Ready for Test
# closed: Closed
# state_labels: # optional column-header overrides (auto-derived if omitted)
# active: active
# ready_for_test: rft
# closed: closed
Configuration Options:
organization : Your Azure DevOps organization name (required)
projects : List of Azure DevOps project names (required). Each entry can be a plain string or an object with name and display_name fields. The display_name is shown in the TUI while the name is used for API calls.
polling_interval : How often to refresh data in seconds (optional, default: 60)
theme : Color theme for the UI (optional, default: dark)
disabled_panes : Comma-separated list of panes to hide (optional). Valid values: pipelines , workitems . When a pane is disabled, its tab, keyboard shortcuts, and all related UI are removed. Pull Requests cannot be disabled.
metrics : Opt-in management dashboard. See Metrics Configuration below for the full reference, and Features → Metrics Dashboard for what it does.
dark - Dark theme with blue and cyan accents
gruvbox - Retro groove color scheme
nord - Arctic, north-bluish color palette
dracula - Default dark theme with purple and pink accents
catppuccin - Soothing pastel theme (Mocha variant)
retro - Matrix-inspired green phosphor on black
monokai - Classic Monokai color scheme
The metrics dashboard is opt-in and hidden entirely unless metrics.enabled: true . All keys live under the top-level metrics: block and are optional — the defaults below apply when a key is omitted. The validation rules only apply when enabled is true .
These map your board's actual workflow-state strings onto the three buckets the metrics engine tracks. Matching is case-insensitive and whitespace-trimmed , but each takes a single name (no comma-separated aliases). If your board doesn't literally use "Active" / "Ready for Test" / "Closed", set these or the metrics tab will bucket nothing.
Each name must be non-empty, distinct from the other two, and contain no single quote ( ' ) — single quotes are rejected for WIQL-injection safety.
Display-only overrides for the metrics table column headers. When omitted, labels are auto-derived from the configured state name: multi-word names become lowercased initials ( Ready for Test → rft , In Progress → ip ), single-word names are lowercased as-is ( Done → done ).
Example — a board using "Doing" / "QA" / "Done":
metrics :
enabled : true
interval_days : 14
wip_limit : 4
states :
active : Doing
ready_for_test : QA
closed : Done
state_labels :
ready_for_test : QA # override the auto-derived lowercase "qa" to keep the caps
Custom Themes
You can create your own custom themes by placing JSON them

[truncated]

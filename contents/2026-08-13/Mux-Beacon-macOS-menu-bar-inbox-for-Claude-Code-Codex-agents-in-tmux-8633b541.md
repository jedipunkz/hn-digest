---
source: "https://github.com/Lukeesec/mux-beacon"
hn_url: "https://news.ycombinator.com/item?id=49280400"
title: "Mux Beacon – macOS menu-bar inbox for Claude Code/Codex agents in tmux"
article_title: "GitHub - Lukeesec/mux-beacon: Native macOS notifications and a menu-bar inbox for Claude Code and Codex running across tmux sessions. · GitHub"
author: "lmartineng"
captured_at: "2026-08-13T01:05:47Z"
capture_tool: "hn-digest"
hn_id: 49280400
score: 1
comments: 1
posted_at: "2026-08-13T00:32:51Z"
tags:
  - hacker-news
  - translated
---

# Mux Beacon – macOS menu-bar inbox for Claude Code/Codex agents in tmux

- HN: [49280400](https://news.ycombinator.com/item?id=49280400)
- Source: [github.com](https://github.com/Lukeesec/mux-beacon)
- Score: 1
- Comments: 1
- Posted: 2026-08-13T00:32:51Z

## Translation

タイトル: Mux Beacon – tmux のクロード コード/コーデックス エージェント用の macOS メニューバー受信箱
記事のタイトル: GitHub - Lukeesec/mux-beacon: tmux セッション間で実行されるクロード コードとコーデックスのネイティブ macOS 通知とメニューバーの受信箱。 · GitHub
説明: ネイティブ macOS 通知と、tmux セッション間で実行されるクロード コードおよびコーデックスのメニューバー受信ボックス。 - Lukeesec/マルチプレクサビーコン

記事本文:
GitHub - Lukeesec/mux-beacon: ネイティブ macOS 通知と、tmux セッション間で実行されるクロード コードおよびコーデックスのメニュー バーの受信箱。 · GitHub
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
ルケセック
/
マルチプレクサビーコン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .github .github アセット アセット パッケージング パッケージング ソース ソース docs docs scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CHAN

GELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターミナル エージェントがいつあなたを必要とするかを把握し、正確な tmux ターゲットに戻ります。
Mux Beacon は、Claude Code および Codex 用のネイティブ macOS メニューバー受信箱です。文書化されたライフサイクル フックと完了コールバックを使用し、ローカル ターン時間を記録し、オプションのペイン境界バッジを表示し、通知を実行可能にします。
ビジーな tmux セットアップにより、セッションやウィンドウ全体で完了した作業が隠される可能性があります。 Mux Beacon は、各エージェントのターンを小さなライフサイクルに変えます。
プロンプトが送信されました → 動作中 → 準備完了 / 失敗しました
└→ 注意が必要です (オプション)
UserPromptSubmit は開始をすぐに記録します。ノイズを避けるために通知はオプトインになっています。
完了および失敗の通知には、エージェント、プロジェクト、期間、およびセッション › ウィンドウが含まれます。
Ghostty で [開く] をクリックすると、元の tmux クライアントがターゲットになり、キャプチャされた Ghostty ターミナルがフォーカスされます。
確認すると未読状態がクリアされます。記録された時間をマークすることは、受信トレイで利用できます。
PermissionRequest はサポートされていますが、その通知はデフォルトではオフになっています。
プロンプト、コマンド、および最終的な回答は、プレビューが明示的に有効になっていない限り保存されません。
ライフサイクル フック ( StopFailure を含む) を備えた最近の Claude コード、またはフックとその完了コールバックを備えた Codex CLI
Ghostty 1.3 以降では、ウィンドウ/タブに正確にフォーカスできます。他の端末は引き続き受信ボックスと tmux メタデータを受信します
git クローン https://github.com/Lukeesec/mux-beacon.git
CD マルチプレクサ ビーコン
./scripts/install-local.sh
mux-beacon install # ドライラン: フックの変更をプレビューする
mux-beacon install --apply # 書き込みます
アプリは、書き込み可能な場合は /Applications にインストールされ、それ以外の場合は ~/Applications にインストールされ、La に登録されます。

アンチサービス。 Finder は直接開くことができます。 mux-beacon gui は、Spotlight がまだインデックスを作成していないローカル アドホック ビルド用の信頼できるランチャーです。
所有ハンドラーを ~/.claude/settings.json に追加します。
所有ハンドラーを ~/.codex/hooks.json に追加します。
通知スロットが空いている場合、Codex の文書化されたエージェント ターン完了コールバックを ~/.codex/config.toml に追加します。
既存の Codex 通知コマンドを保持し、置き換えるのではなく警告を出力します。
既存のファイルを変更する前に、タイムスタンプ付きのバックアップを書き込みます。
Codex では、新しいフックを一度確認するように求められます。 /hooks を開いて、Mux Beacon 定義を信頼します。完了コールバックは、ライフサイクル停止フックが各ターン後に発行されない Codex バージョンをカバーします。
デフォルトでは、許可イベントは延期されます。必要なユーザーは、アダプターを明示的にインストールし、設定でその通知を有効にすることができます。
mux-beacon install --apply --with-permission-events
macOS のプロンプトが表示されたときに通知を許可します。 [システム設定] → [通知] → [Mux Beacon] で、通知を閉じるまで保持する必要がある場合は、[バナー] の代わりに [アラート] を選択します。 macOS がこの設定を所有しているため、アプリがそれを強制することはできません。 Ghostty の正確なフォーカスを有効にするには、 [システム設定] → [プライバシーとセキュリティ] → [自動化] で、Mux Beacon が Ghostty を自動化できるようにします。
Mux Beacon は主にメニュー バー アプリです。 macOS メニュー バーでビーコン アイコンをクリックし、インデックス作成時に Finder または Spotlight から開くか、次のコマンドを実行します。
マルチプレクサビーコン GUI
アプリを起動すると、ウィンドウを開かずにメニューバー項目が直接開始されます。フックイベントや通知のクリックによって受信トレイが開かれたり、フォーカスされたりすることはありません。 Open window 、新しい mux-beacon gui リクエスト、または muxbeacon://inbox のみがそれを転送します。通知ナビゲーションは Ghostty とキャプチャされた tmux ターゲットに直接戻ります。
フック設定に触れずに試してみる
マルチプレクサビーコンのデモ
マルチプレクサビーコンテスト準備完了 -

-ソースコーデックス
マルチプレクサビーコンのステータス
アプリとデモでは tmux を再起動する必要はありません。フックでは、新しいエージェント プロセスまたは再ロードされたエージェント プロセスが必要になる場合があります。
Mux Beacon はプロセス スキャナーではなくフック駆動です。フックが有効なプロンプトが送信されると、エージェントの追跡が開始されます。インストール前にすでに実行されていたターンを再構築することはできません。 Claude または Codex を起動するだけでは開始イベントは生成されません。
macOS では、プロジェクトと状態が太字のタイトルとして表示され、その下にエージェントと期間が表示され、本文には tmux ルートが表示されます。最終的なレイアウト、切り捨て、永続化、およびフォーカス/DND 配信を制御します。ルーティングの詳細は、不透明なイベント ID として非表示の通知メタデータに存在します。
デモ レコードには DEMO とマークが付けられており、意図的にライブ ジャンプ ターゲットがありません。 GUI は、サンプルデータのコントロールを通常のワークフローから切り離します。 mux-beacon clear-demo を使用してサンプルを削除します。
完了と失敗のアラートがオンになります。開始と許可のアラートはオフになっています。 GUI または CLI から変更します。
マルチプレクサビーコン通知ステータス
マルチプレクサビーコン通知の開始日
マルチプレクサビーコン通知の開始
マルチプレクサビーコン通知はすべてオフ
Mux ビーコンを無効化または削除する
# すべての通知を沈黙させますが、受信トレイ内のターンを記録し続けます
マルチプレクサビーコン通知はすべてオフ
# Mux Beacon のエージェント フックのみを削除して、新しいイベントの収集を停止します
mux-beacon アンインストール --apply
# フックを削除し、アプリをゴミ箱に移動します。地元の歴史が残されている
./scripts/uninstall-local.sh
tmux ビュー
各ペインの上端の左側に状態 (青 ● WORKING 、緑 ● READY 、黄 ● ATTENTION 、または赤 ● FAILED) が表示され、その後に既存のペインのタイトルとペイン番号が表示されます。このイラストでは中立的なテーマが使用されています。 tmux は、端末のフォントと背景を使用してそれをレンダリングします。これらのバッジは、ウィンドウがペインに分割されている場合に最も役立ちます。デスクトップ通知

オプションとメニューバーの受信箱により、非表示のウィンドウとセッション全体が表示されます。
mux-beacon tmux ポップアップ
mux-beacon tmux 有効化バッジ
mux-beacon tmux 無効化バッジ
mux-beacon tmux バッジ-ステータス
mux-beacon tmux ポップアップは、最近のエージェント アクティビティの一時的な tmux オーバーレイを開きます。行番号を入力すると、そのエージェントにジャンプします。 Return キーを押して閉じます。
バッジはオプトインであり、現在の tmux サーバーに適用されます。そのサーバー内からenable-badgesを1回実行します。 bad-status は、境界線が有効かどうか、および状態を追跡しているペインの数を報告します。 Mux Beacon は、正確な既存の pane-border-status および pane-border-format を保存し、 disable-badges を使用して復元します。
ターン期間は、プロンプト提出から完了または失敗まで測定されます。
mux-beacon エクスポート --format json --output mux-beacon-time.json
mux-beacon エクスポート --format csv --output mux-beacon-time.csv
コアは TimeExportProvider と TimeEntryDraft を公開するため、フックや UI コードを変更せずに Clockify アダプターを追加できます。 Direct Clockify 資格情報と API 呼び出しは、最初のリリースから延期されます。 「 Clockify 統合設計 」を参照してください。
コマンド
目的
マルチプレクサビーコンドクター
アプリ、フック、tmux、Ghostty、ローカル ストレージを確認する
マルチプレクサビーコンのステータス
ターミナルでの最近のアクティビティを表示する
マルチプレクサビーコンの健全性
置き換えられたレコードと欠落している tmux ターゲットを廃止する
マルチプレクサビーコン GUI
ネイティブの受信トレイウィンドウを開く
マルチプレクサビーコン通知 …
アラート設定を検査または変更する
マルチプレクサビーコンジャンプラスト
最新の未読イベントを開く
マルチプレクサビーコンデモ / クリアデモ
匿名化されたサンプル データを追加または削除する
mux-beacon アンインストール --apply
Mux Beacon のフック ハンドラーのみを削除します
ルーティングの仕組み
Mux Beacon は、安定した tmux ID と正確なサーバー ソケットを保存します。ナビゲーションには以下が使用されます。
tmux -S < ソケット > スイッチクライアント -c < クライアント tty > -t < ペイン ID >
Ghostty 1.3 は端末 TTY を公開しません。

したがって、Mux Beacon は、プロンプト送信時にフォーカスされた端末 ID を同期的にキャプチャします。 Ghostty 1.4 では TTY/PID プロパティが追加され、直接マッピングが可能になります。あいまいなルートや古いルートは、任意の端末を切り替える代わりにフェールクローズされます。
受信トレイは、30 秒ごと、および [更新] をクリックするたびにターゲットの状態をチェックします。同じ tmux ターゲットをオンにする古いアクティブなターンと、ペインがもう存在しないイベントは古いものとして認識され、履歴の下に 7 日間保持されます。実行中のレコードと未読のレコードは、履歴のクリーンアップによって削除されることはありません。
「アーキテクチャ、開発、およびトラブルシューティング」を参照してください。
タグ付きリリースには、CI によって構築されたアプリ zip が添付されます。これはアドホック署名されており、公証されていないため、macOS はダウンロードされたコピーの最初の起動をブロックします。 システム設定 → プライバシーとセキュリティ → とにかく開く で承認するか、上記のようにソースからビルドします (ローカル ビルドは隔離されません)。
ローカル専用の SQLite データベース。テレメトリーはありません。
ユーザー専用のアプリケーション サポート ディレクトリとフックのバックアップ。
通知からの承認または拒否のアクションはありません。
通知メタデータの不透明なイベント ID。 URL にシェル コマンドや tmux ラベルはありません。
フック コマンドは、エージェントを操作せずに成功を返します。
リリース ビルドには開発者 ID の署名と公証が行われるため、ダウンロードは手動承認なしで Gatekeeper を通過します。
既存のプロバイダー境界上の Clockify エクスポート アダプター (設計)。
MIT © 2026 Lukeesec の寄稿者。
ネイティブ macOS 通知と、tmux セッション間で実行されるクロード コードおよびコーデックスのメニュー バーの受信箱。
Readme MIT ライセンスの行動規範
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Native macOS notifications and a menu-bar inbox for Claude Code and Codex running across tmux sessions. - Lukeesec/mux-beacon

GitHub - Lukeesec/mux-beacon: Native macOS notifications and a menu-bar inbox for Claude Code and Codex running across tmux sessions. · GitHub
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
Lukeesec
/
mux-beacon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .github .github Assets Assets Packaging Packaging Sources Sources docs docs scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Know when terminal agents need you—and get back to the exact tmux target.
Mux Beacon is a native macOS menu-bar inbox for Claude Code and Codex. It uses documented lifecycle hooks and completion callbacks, records local turn duration, shows optional pane-border badges, and makes notifications actionable.
A busy tmux setup can hide finished work across sessions and windows. Mux Beacon turns each agent turn into a small lifecycle:
prompt submitted → working → ready / failed
└→ needs attention (optional)
UserPromptSubmit records the start immediately; its notification is opt-in to avoid noise.
Completion and failure notifications contain agent, project, duration, and session › window .
Clicking Open in Ghostty targets the originating tmux client and focuses the captured Ghostty terminal.
Acknowledge clears the unread state; Mark time logged is available in the inbox.
PermissionRequest is supported but its notification is off by default.
Prompts, commands, and final answers are not stored unless previews are explicitly enabled.
A recent Claude Code with lifecycle hooks (including StopFailure ), or Codex CLI with hooks and its completion callback
Ghostty 1.3+ for exact window/tab focus; other terminals still receive the inbox and tmux metadata
git clone https://github.com/Lukeesec/mux-beacon.git
cd mux-beacon
./scripts/install-local.sh
mux-beacon install # dry run: preview the hook changes
mux-beacon install --apply # write them
The app installs into /Applications when writable, otherwise ~/Applications , and registers with Launch Services. Finder can open it directly; mux-beacon gui is the reliable launcher for local ad-hoc builds that Spotlight has not indexed yet.
adds owned handlers to ~/.claude/settings.json ;
adds owned handlers to ~/.codex/hooks.json ;
adds Codex's documented agent-turn-complete callback to ~/.codex/config.toml when the notify slot is free;
preserves an existing Codex notify command and prints a warning instead of replacing it;
writes timestamped backups before changing existing files.
Codex asks you to review new hooks once. Open /hooks and trust the Mux Beacon definitions. The completion callback covers Codex versions where the lifecycle Stop hook is not emitted after each turn.
Permission events are deferred by default. Users who want them can install the adapter explicitly and then enable its notification in Settings:
mux-beacon install --apply --with-permission-events
Allow notifications when macOS prompts. In System Settings → Notifications → Mux Beacon , choose Alerts instead of Banners if notifications should remain until dismissed; macOS owns this setting, so apps cannot enforce it. To enable exact Ghostty focus, allow Mux Beacon to automate Ghostty under System Settings → Privacy & Security → Automation .
Mux Beacon is primarily a menu-bar app. Click its beacon icon in the macOS menu bar, open it from Finder or Spotlight when indexed, or run:
mux-beacon gui
Launching the app directly starts its menu-bar item without opening a window. Hook events and notification clicks never open or focus the inbox; only Open window , a fresh mux-beacon gui request, or muxbeacon://inbox brings it forward. Notification navigation returns directly to Ghostty and the captured tmux target.
Try it without touching hook configuration
mux-beacon demo
mux-beacon test ready --source codex
mux-beacon status
The app and demo require no tmux restart. Hooks may require a new or reloaded agent process.
Mux Beacon is hook-driven rather than a process scanner. It begins tracking an agent when a hook-enabled prompt is submitted; it cannot reconstruct turns that were already running before installation. Merely launching Claude or Codex does not produce a start event.
macOS renders the project and state as the bold title, with agent and duration beneath it and the tmux route in the body. It controls final layout, truncation, persistence, and Focus/DND delivery. Routing details live in hidden notification metadata as an opaque event ID.
Demo records are marked DEMO and intentionally have no live jump target. The GUI keeps sample-data controls out of the normal workflow; remove samples with mux-beacon clear-demo .
Completion and failure alerts are on. Start and permission alerts are off. Change them in the GUI or from the CLI:
mux-beacon notifications status
mux-beacon notifications start on
mux-beacon notifications start off
mux-beacon notifications all off
Disable or remove Mux Beacon
# Silence every notification but keep recording turns in the inbox
mux-beacon notifications all off
# Stop collecting new events by removing only Mux Beacon's agent hooks
mux-beacon uninstall --apply
# Remove hooks and move the app to Trash; local history is retained
./scripts/uninstall-local.sh
tmux views
The state appears at the left of each pane's top border—blue ● WORKING , green ● READY , yellow ● ATTENTION , or red ● FAILED —followed by the existing pane title and pane number. The illustration uses a neutral theme; tmux renders it using your terminal's font and background. These badges are most useful when a window is split into panes; desktop notifications and the menu-bar inbox provide visibility across hidden windows and sessions.
mux-beacon tmux popup
mux-beacon tmux enable-badges
mux-beacon tmux disable-badges
mux-beacon tmux badge-status
mux-beacon tmux popup opens a temporary tmux overlay of recent agent activity. Enter a row number to jump to that agent; press Return to close it.
Badges are opt-in and apply to the current tmux server. Run enable-badges once from inside that server; badge-status reports whether borders are enabled and how many panes have tracked state. Mux Beacon saves the exact existing pane-border-status and pane-border-format and restores them with disable-badges .
Turn duration is measured from prompt submission until completion or failure.
mux-beacon export --format json --output mux-beacon-time.json
mux-beacon export --format csv --output mux-beacon-time.csv
The core exposes TimeExportProvider and TimeEntryDraft so a Clockify adapter can be added without changing hook or UI code. Direct Clockify credentials and API calls are deferred from the first release; see Clockify integration design .
Command
Purpose
mux-beacon doctor
Check app, hooks, tmux, Ghostty, and local storage
mux-beacon status
Show recent activity in the terminal
mux-beacon health
Retire superseded records and missing tmux targets
mux-beacon gui
Open the native inbox window
mux-beacon notifications …
Inspect or change alert preferences
mux-beacon jump-last
Open the newest unread event
mux-beacon demo / clear-demo
Add or remove anonymized sample data
mux-beacon uninstall --apply
Remove only Mux Beacon's hook handlers
How routing works
Mux Beacon stores stable tmux IDs and the exact server socket. Navigation uses:
tmux -S < socket > switch-client -c < client-tty > -t < pane-id >
Ghostty 1.3 does not expose a terminal TTY, so Mux Beacon captures the focused terminal ID synchronously at prompt submission. Ghostty 1.4 adds TTY/PID properties, allowing direct mapping. Ambiguous or stale routes fail closed instead of switching an arbitrary terminal.
The inbox checks target health every 30 seconds and whenever Refresh is clicked. Older active turns on the same tmux target and events whose panes no longer exist are acknowledged as stale and retained under History for 7 days. Running and unread records are never removed by history cleanup.
See Architecture , Development , and Troubleshooting .
Tagged releases attach an app zip built by CI. It is ad-hoc signed and not notarized, so macOS blocks the first launch of a downloaded copy: approve it under System Settings → Privacy & Security → Open Anyway , or build from source as shown above (local builds are not quarantined).
Local-only SQLite database; no telemetry.
User-only application-support directory and hook backups.
No approval or denial actions from notifications.
Opaque event IDs in notification metadata; no shell commands or tmux labels in URLs.
Hook commands return success without steering the agent.
Developer ID signing and notarization for release builds, so downloads pass Gatekeeper without manual approval.
Clockify export adapter on the existing provider boundary ( design ).
MIT © 2026 Lukeesec contributors.
Native macOS notifications and a menu-bar inbox for Claude Code and Codex running across tmux sessions.
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

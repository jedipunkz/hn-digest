---
source: "https://github.com/yazon/oh-my-wrist"
hn_url: "https://news.ycombinator.com/item?id=48426168"
title: "Show HN: Oh my wrist – Garmin alerts for Claude Code and OpenCode"
article_title: "GitHub - yazon/oh-my-wrist: Display real-time Claude Code and Opencode activity on a Garmin smartwatch via BLE. · GitHub"
author: "yazon"
captured_at: "2026-06-06T18:33:28Z"
capture_tool: "hn-digest"
hn_id: 48426168
score: 2
comments: 0
posted_at: "2026-06-06T15:51:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Oh my wrist – Garmin alerts for Claude Code and OpenCode

- HN: [48426168](https://news.ycombinator.com/item?id=48426168)
- Source: [github.com](https://github.com/yazon/oh-my-wrist)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T15:51:50Z

## Translation

タイトル: HN を表示: ああ、私の手首 – クロード コードと OpenCode に関する Garmin アラート
記事タイトル: GitHub - yazon/oh-my-wrist: BLE 経由で Garmin スマートウォッチにリアルタイムのクロード コードと Opencode アクティビティを表示します。 · GitHub
説明: BLE 経由で Garmin スマートウォッチにリアルタイムのクロード コードと Opencode アクティビティを表示します。 - ヤゾン/オーマイリスト

記事本文:
GitHub - yazon/oh-my-wrist: BLE 経由で Garmin スマートウォッチにリアルタイムのクロード コードと Opencode アクティビティを表示します。 · GitHub
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
ヤゾン
/
ああ、手首
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション操作

ション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
36 コミット 36 コミット .github/ workflows .github/ workflows docs docs garmin garmin opencode/ plugins opencode/ plugins src/ ohm src/ ohm static static testing テストツール tools .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierrc .prettierrc ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
oh-my-wrist は、Bluetooth Low Energy 経由で Garmin スマートウォッチにリアルタイムのクロード コードと OpenCode アクティビティを表示します。 AI コーディング アシスタントが何をしているのかを手首で確認できます。
リアルタイムの BLE アップデート — ツール呼び出し、ファイル編集、セッション状態がウォッチにストリーミングされます
触覚アラート — アイドル状態、セッション完了、破壊的なコマンド、エージェントの完了の振動パターン
CLI スタイルのウォッチ UI — アニメーション化された琥珀色のスピナー、イベント スタック、プロバイダーごとの統計を備えた端末の美しさ
マルチプロバイダー — Claude Code と OpenCode を同時にサポート
クロスプラットフォーム — Linux、macOS、Windows 上で実行可能
静かな時間帯 - 設定可能な時間帯で振動を抑制します。
接続 ID フィルター - オプションの 0 ～ 255 ID により、近くのユーザーのウォッチが相互にペアリングされるのを防ぎます。
pip インストール オーマイリスト
または uv を使用して:
uv pip インストール オーマイリスト
2. プロバイダーとシステムサービスを構成する
オーマイリストのインストール
これは自動的に次のようになります。
~/.claude/settings.json の Claude コード フックにパッチを適用します。
/usage クォータをウォッチにストリーミングするクロード コード statusLine を構成します (既存の statusLine をチェーンします)。
OpenCode TypeScript プラグインをインストールします (opencode が PATH 上にある場合)
バックグラウンド システム サービス (systemd / launchd / タスク スケジューラ) を登録します。
単一プロバイダーのみをインストールするには:
oh-my-wrist install --provider クロード

oh-my-wrist install --provider opencode
3. Garmin ウォッチ アプリをインストールする
GitHub リリース: 最新リリースから oh-my-wrist-prg-<version>.zip をダウンロードして解凍し、USB 経由で Garmin デバイスを接続し、一致する oh-my-wrist-<device-id>.prg ファイルを /GARMIN/Apps/ にコピーします。
ソースからビルド: Connect IQ SDK 9.1+ をインストールし、tools/build_garmin.sh release を実行して、生成されたデバイス固有の .prg ファイルを build/garmin/ から /GARMIN/Apps/ にコピーします。 VS Code の単一デバイスは garmin/write bin/oh-my-wrist.prg からビルドされます。
ああ、手首からスタート
時計で oh-my-wrist アプリを開きます。自動的に接続され、リアルタイムで更新されます。
プラットフォーム
ガイド
🐧 リナックス
docs/INSTALL_LINUX.md
🍎macOS
docs/INSTALL_MACOS.md
🪟 Windows
docs/INSTALL_WINDOWS.md
使用法
デーモンをバックグラウンド (またはデバッグ用のフォアグラウンド) で起動します。
ああ、手首からスタート
oh-my-wrist start --foreground
Claude Code または OpenCode セッションを開始します。ウォッチは自動的に更新されます。
ソース チェックアウトから 1 分間の診断ストリームを実行して、
実際の Claude Code または OpenCode セッションを開始せずに、デーモンから監視へのリンクを実行します。
Pythonツール/check_connection.py
デーモンを実行し、Watch アプリを開いたままにします。履歴行が表示されるはずです
変化し、Claude/OpenCode 統計が増加し、Claude 使用状況バーが移動します。役に立つ
オプション:
python tools/check_connection.py --duration 30
python tools/check_connection.py --dry-run
python tools/check_connection.py --provider クロード
ウォッチナビゲーション
4 つのスワイプ可能なビュー (ボタンウォッチでは左/右または上/下)。歴史というのは、
初期ビュー — クロードの使用状況画面を表示するには上にスワイプ/押し、
プロバイダーごとの統計画面:
使用状況画面はクロードのみで、パーセンテージのない空のバーが表示されます。
クォータ データが利用できない場合 (API キー ユーザー、または最初の API の前)
セッションでの応答

n)。
任意のビューで SELECT/START を押して、アプリ メニューを開きます。 「ID を設定」を使用して、
デスクトップ デーモンに構成されているのと同じ 0 ～ 255 の接続 ID を保存してから、
新しい BLE サービス UUID が正常に登録されるように、Watch アプリを再起動します。
コマンド
目的
ああ、私の手首のスタート [-f|--foreground]
BLEデーモンを起動する
ああ、リストストップ
デーモンを停止する
ああ、手首のステータス
デーモン PID、サービス ステータス、フック/プラグインの状態を表示します
oh-my-wrist インストール [--provider claude|opencode|both]
フック、プラグイン、および OS サービスを構成する
oh-my-wrist アンインストール [--provider …]
フック、プラグイン、サービスを削除する
oh-my-wrist テスト [メッセージ] [--provider claude|opencode]
デーモンにテストメッセージを送信する
oh-my-wrist config [--show|--haptic オン|オフ|--quiet-start HH:MM|--quiet-end HH:MM]
構成の表示または更新
oh-my-wrist set-id ID
BLE 接続 ID ( 0 ～ 255 ) を設定し、実行中のデーモンの更新をキューに入れます。
oh-my-wrist opencode インストール|アンインストール|ステータス
OpenCode プラグインを管理する
オーマイリストログ [-n LINES]
デーモンログを追跡します
構成
設定は ~/.oh-my-wrist/config.json に保存され、すぐに有効になります (再起動は必要ありません)。
# 現在の設定を表示する
オーマイリスト設定 --show
# 触覚アラートの切り替え
oh-my-wrist config --haptic オン
オーマイリスト設定 --触覚オフ
# 静かな時間を設定（振動を抑制）
oh-my-wrist config --quint-start 22:00 --quiet-end 08:00
# ウォッチ メニューの Set id 値と一致するように BLE 接続 ID を設定します
オーマイリストセットID 42
設定
デフォルト
説明
ハプティック有効
本当の
すべての振動アラートを有効/無効にする
クワイエットスタート
22:00
静かなウィンドウの開始 (HH:MM)
静かな終わり
08:00
静かなウィンドウの終了 (HH:MM)
接続ID
0
BLE スキャン フィルター ID。ウォッチは、同じ ID を持つデーモンにのみ接続します。
接続 ID は認証ではありません。部屋の衝突軽減装置です
複数の人がオーマイリストを実行する場所。 ID0

がデフォルトであり、
元の BLE サービス UUID。
Connect IQ Generic BLE サポートを備えた Garmin ウォッチが必要です (API レベルだけでは十分ではありません)。
マニフェスト製品 ID: Approachs50 、 Approachs7042mm 、 Approachs7047mm 、 d2air 、 d2airx10 、 d2mach1 、 d2mach2 、 d2mach2pro 、 descentg1 、 descentg2 、 descentmk2 、 descentmk2s 、 descentmk343mm 、 descentmk351mm 、 edge1030 、edge1030plus 、edge1040 、edge1050 、edge530 、edge540 、edge550 、edge830 、edge840 、edge850 、edgeexplore 、edgeexplore2 、edgemtb 、enduro 、enduro3 、epix2 、epix2pro42mm 、epix2pro47mm 、 epix2pro51mm 、 etrextouch 、 fenix5plus 、 fenix5splus 、 fenix5xplus 、 fenix6 、 fenix6pro 、 fenix6s 、 fenix6spro 、 fenix6xpro 、 fenix7 、 fenix7pro 、 fenix7pronowifi 、 fenix7s 、 fenix7spro 、 fenix7x 、 fenix7xpro 、 fenix7xpronowwifi 、 fenix843mm 、 fenix847mm 、 fenix8pro47mm 、 fenix8solar47mm 、 fenix8solar51mm 、 フェニックス 、 fr165 、 fr165m 、 fr170 、 fr170m 、 fr245 、 fr245m 、 fr255 、 fr255m 、 fr255s 、 fr255sm 、 fr265 、 fr265s 、 fr55 、 fr57042mm 、 fr57047mm 、 fr645m 、 fr70 、 fr745 、 fr945 、 fr945lte 、 fr955 、 fr965 、 fr970 、 gpsmap66 、 gpsmap67 、 gpsmaph1 、本能 2 、本能 2s 、本能 2x 、本能 3amoled45mm 、本能 3amoled50mm 、本能 3solar45mm 、本能クロスオーバー、本能クロスオーバーアモールド、本能 40mm 、本能 45mm 、レガシーヒーローキャプテンマーベル、レガシーヒーローファーストアベンジャー、レガシーサガダースベイダー、レガシーサガリー 、 マルク2 、 マルク2アビエーター 、 マルカアドベンチャー 、 マルカアスリート 、 マルキャビエーター 、 マルクキャプテン 、 マルコマンダー 、 マルクドライバー 、 マルクエクスペディション 、 マルクゴルファー 、 モンタナ7xx 、 ヴェヌ 、 ヴェヌ2 、 ヴェヌ2プラス 、 ヴェヌ2s 、 ヴェヌ3 、 ヴェヌ3s 、 venu441mm 、 venu445mm 、 venud 、 venusq2m 、 venusqm 、 venux1 、 vivoactive3m 、 vivoactive3mlte 、 vivoactive4 、 vivoactive4s 、 vivoactive5 、 vivoactive6 。
問題
解決策
ウォッチがデーモンを見つけられない
走れ、ああ、

手首のステータス — デーモンがアドバタイズしていること、およびデスクトップ/ウォッチの接続 ID が一致していることを確認します。オーマイリストストップ && オーマイリストスタート で再起動します。ソース チェックアウトから、ウォッチ アプリを開いた状態で python tools/check_connection.py を実行し、ライブ履歴、統計、使用状況の更新を実行します。
Garmin アプリが PC 上のデーモンに接続しない
Bluetooth アダプターがサポートされていることを確認してください。 Windows のランダム MAC 動作などの一部のアダプターまたは OS 設定により、安定した BLE 接続が妨げられる場合があります。問題が解決しない場合は、macOS または Linux をお勧めします。
Garmin アプリが時計で起動しない
デバイス上に空の /GARMIN/Apps/Logs/oh-my-wrist-YOUR_WATCH_ID.log (例: oh-my-wrist-fenix7x.log) ファイルを作成し、アプリを実行し、問題が発生するのを待ち、デバイスを PC に再接続し、ログ ファイルをダウンロードして、デバッグ用のログを含む GitHub 問題を作成します。
フックが発砲しない
oh-my-wrist status を実行します — ~/.claude/settings.json のフックを確認します。 oh-my-wrist install を再実行します。
OpenCodeが更新されない
プラグインを確認してください: oh-my-wrist オープンコード ステータス。再インストール: oh-my-wrist opencode install 。
BLE 権限エラー
プラットフォーム ガイドを参照してください: Linux · macOS · Windows
デーモンが起動時にクラッシュする
oh-my-wrist start --foreground を実行してエラーを確認します。オーマイリストのログを確認してください。
アンインストール
オーマイリストのアンインストール
pip アンインストール オーマイリスト
ウォッチ アプリを削除するには、Garmin Express または携帯電話の Connect IQ アプリを介して削除します。
貢献は大歓迎です。ぜひ貢献してください！ガイドラインについては、CONTRIBUTING.md を参照してください。
# ピップ
pip install -e " .[dev] "
pytest テスト/
# または UV
uv venv && uv pip install -e " .[dev] "
UV で pytest を実行
oh-my-wrist start --foreground
ライセンス
このプロジェクトは、BSD 3 条項ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
BLE 経由で Garmin スマートウォッチにリアルタイムの Claude Code と Opencode アクティビティを表示します。
BSD-3-第 1 条

香
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
リリース v0.1.1
最新の
2026 年 6 月 5 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Display real-time Claude Code and Opencode activity on a Garmin smartwatch via BLE. - yazon/oh-my-wrist

GitHub - yazon/oh-my-wrist: Display real-time Claude Code and Opencode activity on a Garmin smartwatch via BLE. · GitHub
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
yazon
/
oh-my-wrist
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
36 Commits 36 Commits .github/ workflows .github/ workflows docs docs garmin garmin opencode/ plugins opencode/ plugins src/ ohm src/ ohm static static tests tests tools tools .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .prettierrc .prettierrc LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
oh-my-wrist displays real-time Claude Code and OpenCode activity on your Garmin smartwatch over Bluetooth Low Energy. See what your AI coding assistant is doing — right on your wrist.
Real-time BLE updates — tool calls, file edits, and session state streamed to your watch
Haptic alerts — vibration patterns for idle, session done, destructive commands, and agent completion
CLI-styled watch UI — terminal aesthetic with animated amber spinner, event stack, and per-provider stats
Multi-provider — supports Claude Code and OpenCode simultaneously
Cross-platform — runs on Linux, macOS, and Windows
Quiet hours — suppress vibrations during configurable time windows
Connection ID filter — optional 0–255 ID keeps nearby users' watches from pairing with each other
pip install oh-my-wrist
Or with uv :
uv pip install oh-my-wrist
2. Configure providers and system service
oh-my-wrist install
This automatically:
Patches Claude Code hooks in ~/.claude/settings.json
Configures a Claude Code statusLine that streams /usage quota to the watch (chaining any existing statusLine)
Installs the OpenCode TypeScript plugin (when opencode is on PATH)
Registers a background system service (systemd / launchd / Task Scheduler)
To install a single provider only:
oh-my-wrist install --provider claude
oh-my-wrist install --provider opencode
3. Install the Garmin watch app
GitHub Releases: download oh-my-wrist-prg-<version>.zip from the latest release, unzip it, connect your Garmin device over USB, then copy the matching oh-my-wrist-<device-id>.prg file to /GARMIN/Apps/ .
Build from source: install Connect IQ SDK 9.1+, then run tools/build_garmin.sh release and copy the generated device-specific .prg file from build/garmin/ to /GARMIN/Apps/ . VS Code single-device builds from garmin/ write bin/oh-my-wrist.prg .
oh-my-wrist start
Open the oh-my-wrist app on your watch — it connects automatically and updates in real time.
Platform
Guide
🐧 Linux
docs/INSTALL_LINUX.md
🍎 macOS
docs/INSTALL_MACOS.md
🪟 Windows
docs/INSTALL_WINDOWS.md
Usage
Start the daemon in the background (or foreground for debugging):
oh-my-wrist start
oh-my-wrist start --foreground
Start a Claude Code or OpenCode session — the watch updates automatically.
From a source checkout, run the one-minute diagnostic stream to verify the
daemon-to-watch link without starting a real Claude Code or OpenCode session:
python tools/check_connection.py
Keep the daemon running and the watch app open. You should see history rows
change, Claude/OpenCode stats increment, and Claude usage bars move. Useful
options:
python tools/check_connection.py --duration 30
python tools/check_connection.py --dry-run
python tools/check_connection.py --provider claude
Watch Navigation
Four swipeable views (left/right or UP/DOWN on button watches). History is the
initial view — swipe/press UP for the Claude usage screen, DOWN for the
per-provider stats screens:
The usage screen is Claude-only and shows an empty bar with no percentage
when quota data is unavailable (API-key users, or before the first API
response in a session).
Press SELECT/START on any view to open the app menu. Use Set id to
save the same 0–255 connection ID configured on your desktop daemon, then
restart the watch app so the new BLE service UUID is registered cleanly.
Command
Purpose
oh-my-wrist start [-f|--foreground]
Start the BLE daemon
oh-my-wrist stop
Stop the daemon
oh-my-wrist status
Show daemon PID, service status, hook/plugin state
oh-my-wrist install [--provider claude|opencode|both]
Configure hooks, plugin, and OS service
oh-my-wrist uninstall [--provider …]
Remove hooks, plugin, and service
oh-my-wrist test [message] [--provider claude|opencode]
Send a test message to the daemon
oh-my-wrist config [--show|--haptic on|off|--quiet-start HH:MM|--quiet-end HH:MM]
View or update configuration
oh-my-wrist set-id ID
Set BLE connection ID ( 0 – 255 ) and queue an update for a running daemon
oh-my-wrist opencode install|uninstall|status
Manage the OpenCode plugin
oh-my-wrist logs [-n LINES]
Tail the daemon log
Configuration
Settings are stored at ~/.oh-my-wrist/config.json and take effect immediately (no restart needed).
# View current config
oh-my-wrist config --show
# Toggle haptic alerts
oh-my-wrist config --haptic on
oh-my-wrist config --haptic off
# Set quiet hours (vibrations suppressed)
oh-my-wrist config --quiet-start 22:00 --quiet-end 08:00
# Set BLE connection ID to match the watch menu's Set id value
oh-my-wrist set-id 42
Setting
Default
Description
haptic_enabled
true
Enable/disable all vibration alerts
quiet_start
22:00
Start of quiet window (HH:MM)
quiet_end
08:00
End of quiet window (HH:MM)
connection_id
0
BLE scan filter ID. The watch only connects to daemons with the same ID.
The connection ID is not authentication; it is a collision reducer for rooms
where multiple people run oh-my-wrist. ID 0 is the default and preserves the
original BLE service UUID.
Requires a Garmin watch with Connect IQ Generic BLE support (API level alone is not enough):
Manifest product IDs: approachs50 , approachs7042mm , approachs7047mm , d2air , d2airx10 , d2mach1 , d2mach2 , d2mach2pro , descentg1 , descentg2 , descentmk2 , descentmk2s , descentmk343mm , descentmk351mm , edge1030 , edge1030plus , edge1040 , edge1050 , edge530 , edge540 , edge550 , edge830 , edge840 , edge850 , edgeexplore , edgeexplore2 , edgemtb , enduro , enduro3 , epix2 , epix2pro42mm , epix2pro47mm , epix2pro51mm , etrextouch , fenix5plus , fenix5splus , fenix5xplus , fenix6 , fenix6pro , fenix6s , fenix6spro , fenix6xpro , fenix7 , fenix7pro , fenix7pronowifi , fenix7s , fenix7spro , fenix7x , fenix7xpro , fenix7xpronowifi , fenix843mm , fenix847mm , fenix8pro47mm , fenix8solar47mm , fenix8solar51mm , fenixe , fr165 , fr165m , fr170 , fr170m , fr245 , fr245m , fr255 , fr255m , fr255s , fr255sm , fr265 , fr265s , fr55 , fr57042mm , fr57047mm , fr645m , fr70 , fr745 , fr945 , fr945lte , fr955 , fr965 , fr970 , gpsmap66 , gpsmap67 , gpsmaph1 , instinct2 , instinct2s , instinct2x , instinct3amoled45mm , instinct3amoled50mm , instinct3solar45mm , instinctcrossover , instinctcrossoveramoled , instincte40mm , instincte45mm , legacyherocaptainmarvel , legacyherofirstavenger , legacysagadarthvader , legacysagarey , marq2 , marq2aviator , marqadventurer , marqathlete , marqaviator , marqcaptain , marqcommander , marqdriver , marqexpedition , marqgolfer , montana7xx , venu , venu2 , venu2plus , venu2s , venu3 , venu3s , venu441mm , venu445mm , venud , venusq2m , venusqm , venux1 , vivoactive3m , vivoactive3mlte , vivoactive4 , vivoactive4s , vivoactive5 , vivoactive6 .
Problem
Solution
Watch can't find daemon
Run oh-my-wrist status — verify daemon is advertising and that desktop/watch connection IDs match. Restart with oh-my-wrist stop && oh-my-wrist start . From a source checkout, run python tools/check_connection.py with the watch app open to exercise live HISTORY, stats, and usage updates.
Garmin app does not connect to daemon on PC
Make sure your Bluetooth adapter is supported. Some adapters or OS settings, such as Windows random MAC behavior, can prevent stable BLE connections. If problems persist, macOS or Linux is preferred.
Garmin app does not start on watch
Create an empty /GARMIN/Apps/Logs/oh-my-wrist-YOUR_WATCH_ID.log (e.g oh-my-wrist-fenix7x.log) file on the device, run the app, wait for the issue to occur, reconnect the device to your PC, download log file, then create a GitHub issue with the log for debugging.
Hook not firing
Run oh-my-wrist status — confirm hooks in ~/.claude/settings.json . Re-run oh-my-wrist install .
OpenCode not updating
Check plugin: oh-my-wrist opencode status . Re-install: oh-my-wrist opencode install .
BLE permission errors
See your platform guide: Linux · macOS · Windows
Daemon crashes on start
Run oh-my-wrist start --foreground to see the error. Check oh-my-wrist logs .
Uninstall
oh-my-wrist uninstall
pip uninstall oh-my-wrist
To remove the watch app, delete it via Garmin Express or the Connect IQ app on your phone.
Contributions are more than welcome. Please contribute! See CONTRIBUTING.md for guidelines.
# pip
pip install -e " .[dev] "
pytest tests/
# or uv
uv venv && uv pip install -e " .[dev] "
uv run pytest
oh-my-wrist start --foreground
License
This project is licensed under the BSD 3-Clause License. See the LICENSE file for details.
Display real-time Claude Code and Opencode activity on a Garmin smartwatch via BLE.
BSD-3-Clause license
Contributing
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
Release v0.1.1
Latest
Jun 5, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

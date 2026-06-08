---
source: "https://github.com/zioder/TaskbarQuota"
hn_url: "https://news.ycombinator.com/item?id=48439651"
title: "TaskbarQuota – Real-time AI quota tracker in the Windows taskbar"
article_title: "GitHub - zioder/TaskbarQuota: AI coding-tool usage in the Windows taskbar · GitHub"
author: "zioder"
captured_at: "2026-06-08T01:00:42Z"
capture_tool: "hn-digest"
hn_id: 48439651
score: 1
comments: 0
posted_at: "2026-06-07T23:19:30Z"
tags:
  - hacker-news
  - translated
---

# TaskbarQuota – Real-time AI quota tracker in the Windows taskbar

- HN: [48439651](https://news.ycombinator.com/item?id=48439651)
- Source: [github.com](https://github.com/zioder/TaskbarQuota)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T23:19:30Z

## Translation

タイトル: TaskbarQuota – Windows タスクバーのリアルタイム AI クォータ トラッカー
記事タイトル: GitHub - zioder/TaskbarQuota: Windows タスクバーでの AI コーディング ツールの使用法 · GitHub
説明: Windows タスクバーでの AI コーディング ツールの使用法。 GitHub でアカウントを作成して、zioder/TaskbarQuota の開発に貢献してください。

記事本文:
GitHub - zioder/TaskbarQuota: Windows タスクバーでの AI コーディング ツールの使用法 · GitHub
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
ジオダー
/
タスクバークォータ
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります

ティング
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
27 コミット 27 コミット .github .github installer installer src src testing/ TaskbarQuota.Tests testing/ TaskbarQuota.Tests .gitignore .gitignore LICENSE LICENSE README.md README.md TaskbarQuota.slnx TaskbarQuota.slnx すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ネイティブの Windows タスクバー コンパニオン。使用している AI ツール (焦点を当てたデスクトップ アプリまたはターミナルで実行されているエージェント) を自動的に検出し、そのプロバイダーのみのライブ使用状況を表示します。
セッション、レートウィンドウ、コストまたはバランス。トレイの横にあるコンパクトなウィジェット。オンデマンドの完全なダッシュボード。クラウドバックエンドはありません。
タスクバークォータ.mp4
TaskbarQuota は、Windows タスクバー (通知領域の隣) に小さな XAML アイランドを挿入する WinUI 3 デスクトップ アプリです。ツールを切り替えるたびにリストからプロバイダーを選択する必要はありません。アプリは、Windows 上で実際に実行していることを監視し、それに合わせてウィジェット、フライアウト、アクティブ フェッチを切り替えます。
内部 WinCheck コードネームからアップグレードした後の最初の起動時に、TaskbarQuota フォルダーに設定と credentials.json が含まれていない場合は、%LOCALAPPDATA%\WinCheck から設定と credentials.json がコピーされます。
アクティブなツールを自動的に追跡します
TaskbarQuota は、「現在何を使用していますか? クォータはどれくらい残っていますか?」という 1 つの質問に答えます。次の 2 つのケースが区別されます。
カーソルから Alt-Tab キーを押してクロード コードを実行している PowerShell セッションに入ると、ウィジェットは約 0.5 秒以内にクロードに再ターゲットされます。サポートされているすべてのツールを終了すると、いずれかのツールに戻るまでウィジェットがフェードアウトすることがあります。フォアグラウンド アプリが無関係な場合 (ブラウザー、Slack など)、TaskbarQuota は最後にアクティブなプロバイダーを表示し続けるため、タスクバーにはコンテキストが残ります。
カーソル IDE 重視 → カーソルの使用法
Windows ターミナ

コマンドラインで l + claude → Claude の使用法
VS Code 中心 → GitHub Copilot の使用
OpenCode モデルの切り替え (Zen vs Go) → TaskbarQuota を再起動せずにプロバイダーを更新
検出後、使用状況はローカル資格情報またはプロバイダー API から取得されます (「サポートされているプロバイダー」を参照)。ウィジェットをクリックするとポップアップが表示されます。すべてのプロバイダーのフル ウィンドウを一度に開き、時間をリセットし、資格情報を手動で修正します。
タスクバー内のコンパクト UI — システム トレイの横に挿入された使用量バーやパーセンテージ (階層化された WinUI アイランドを Shell_TrayWnd に親に戻します)。
タスク バーのレイアウトに従います。タスク バーが移動したり中央に配置されたり、ウィジェット/トレイの形状が変更されたときに中央に再配置されます。
ドラッグ可能な配置 — ドラッグして位置を変更します。オフセットは %LOCALAPPDATA%\TaskbarQuota\ に保存されます。
クリックするとポップアップが表示されます — ウィジェットの上にあるクイックプロバイダーストリップと設定のショートカット。
アイドル時はフェード — サポートされている AI ツール プロセスが検出されない場合、ウィジェットは非表示になります。サポートされているアプリまたはターミナル セッションにフォーカスすると返されます。
アクティブツールの検出 (デスクトップ アプリまたはターミナル エージェント)
これが中心的な動作であり、他のすべて (ウィジェット、フライアウト、キャッシュのフェッチ) はこれに依存します。
デスクトップ アプリ — フォーカスされているウィンドウのプロセス名: Cursor、Antigravity、Codex、Claude、VS Code / Insiders (Copilot) など。
ターミナル エージェント — 既知のターミナルがフォーカスされている場合、WMI/プロセス スキャンは Claude Code、Codex、cursor-agent 、OpenCode、gh copilot 、Antigravity CLI、および関連するランチャー (Windows Terminal、PowerShell、pwsh、WezTerm、Alacritty など) のコマンド ラインを調べます。
高速切り替え — コーディネーターは約 500 ミリ秒ごとにポーリングするため、アプリまたはシェルを変更すると、アクティブなプロバイダーがすぐに更新されます。 API 呼び出しの使用状況は 60 秒間キャッシュされるため、プロバイダーにスパムが送信されることはありません。
OpenCode モデルの切り替え — OpenCode Zen と OpenCode G の間で移動する OpenCode モデルの状態を監視します。

o アプリを再起動せずに。
スティッキー最後のプロバイダー - 無関係なフォアグラウンド アプリはウィジェットをクリアしません。最後に検出されたツールは、サポートされているアプリまたは端末に再度フォーカスするまで残ります。
アイドル非表示 — サポートされている AI プロセスが実行されていない場合のオプションのフェード。サポートされているアプリまたはターミナル セッションにフォーカスすると、再び表示されます。
プロバイダー
あなたが見るもの
使用状況の取得方法
コーデックス
セッションおよびレートウィンドウ
%USERPROFILE%\.codex\auth.json (または %CODEX_HOME% ) からの OAuth トークン → ChatGPT 使用 API
GitHub コパイロット
クォータのスナップショット
GITHUB_TOKEN / GH_TOKEN 、保存された認証情報、または gh 認証トークン → GitHub Copilot 内部 API
クロード
5 時間、週、およびモデル ウィンドウ
%USERPROFILE%\.claude\.credentials.json または環境オーバーライド → Anthropic OAuth 使用法 API
反重力
現地の状況
language_server プロセスの実行 → CSRF トークンとポート → 127.0.0.1 ステータス API
カーソル
使用法の概要
Cursor state.vscdb 、ブラウザ Cookie、または手動 Cookie ヘッダー → Cursor /cursor.com API
OpenCode Zen
請求の使用量と残高
opencode.ai のブラウザ Cookie または手動 Cookie → ワークスペースの請求ページ
OpenCode Go
ローリング、週次、月次のウィンドウ
Zen → OpenCode サーバー ワークスペース API と同じ Cookie パス
ダッシュボード上の各プロバイダー カードには、プラン名、リセット時間、料金ウィンドウ、および API が返すときのオプションのコストまたは残高が表示されます。自動検出でトークンまたは Cookie ヘッダーを %LOCALAPPDATA%\TaskbarQuota\credentials.json に貼り付けることができない場合は、カードで [修正] を使用します。
完全なダッシュボード — 登録されているすべてのプロバイダーを一度に表示し、更新、エラー状態、および詳細カードを表示します。
システム トレイ — 開いてウィジェットを表示し、終了します。
起動時に実行 — オプションの Windows 起動登録。ウィジェットのみを起動できます。
設定 — ライト/ダーク/システム テーマ、ウィジェット レイアウト (バー、パーセンテージ、またはその両方)、消費済みと残りのパーセンテージの表示。
ローカルのみ

— TaskbarQuota バックエンドはありません。使用呼び出しは、PC から各プロバイダー (または Antigravity のローカルホスト) に直接送信されます。
テレメトリなし — 診断ログは %TEMP%\taskbarquota.log のみに記録されます。
Cookie の抽出はメモリ内で行われます。ブラウザの Cookie が読み取られて、アクティブなフェッチのリクエスト ヘッダーが構築されます。これらは意図的に永続化されるものではありません (credentials.json 内の手動認証情報は現在プレーンな JSON です。そのファイルは PC 上にのみ保存してください)。
アクティブアプリ検出器
§─ フォアグラウンド GUI アプリ → ProviderId
└─ フォーカスされた端末 → CLI プロセスのスキャン → ProviderId
↓
UsageCoordinator (500 ミリ秒ティック、アクティブなプロバイダーを選択)
↓
UsageService (プロバイダー レジストリ + 60 秒の成功キャッシュ、エラー時の TTL の短縮 / 429)
↓
IUsageProvider 実装 (HTTP / ローカル API / SQLite / DPAPI)
↓
タスクバーマネージャー → ウィジェットサマリー (タスクバー) + ダッシュボード + フライアウト
検出 — ActiveAppDetector は、フォアグラウンド プロセス (または既知の端末の CLI 子) を ProviderId にマップします。
フェッチ — 一致する IUsageProvider は、CLI 構成ファイル、環境変数、TaskbarQuota 認証情報、Cursor のローカル データベース、または CookieExtractor (ロックされた DB の %TEMP% にコピーされた Chromium / Firefox プロファイル) からトークンを読み込みます。
Normalize — 結果は、RateWindow 行とオプションの CostSnapshot を含むUsageSnapshotになります。
表示 —UsageCoordinator は StateChanged を発生させます。タスクバー ウィジェットとダッシュボードはスナップショットを適用します。フェッチはキャッシュされるため、ティックごとにプロバイダーが攻撃されることはありません。
パス
役割
src/TaskbarQuota.App/ActiveApp/
フォアグラウンドおよび CLI プロバイダーの検出
src/TaskbarQuota.App/Browser/
ブラウザの Cookie の検出と Chromium の復号化
src/TaskbarQuota.App/タスクバー/
ウィジェットホスト、トレイアイコン、タスクバーレイアウトウォッチャー
src/TaskbarQuota.App/使用法/プロバイダー/
プロバイダーごとのフェッチ ロジック
src/TaskbarQuota.App/Usage/UsageService.cs
レジストリとTTLキャッシュ

src/TaskbarQuota.App/UsageCoordinator.cs
タイマーによる検出 → 取得 → UI 更新
src/TaskbarQuota.App/Views/
ダッシュボードと設定
テスト/TaskbarQuota.Tests/
単体テスト
地域発展
Windows 10 19041 以降 (タスクバーの動作については Windows 11 を推奨)。
.NET SDK 10.0.x (プロジェクトは net10.0-windows10.0.19041.0 をターゲットとしています)。
Windows アプリ SDK / WinUI 3 ワークロード。
テストするプロバイダー (ローカル OAuth ファイル、GitHub CLI、カーソル インストールなど) に正常にサインインする必要があります。別途サーバーをインストールする必要はありません。
# ビルド (デバッグ、x64)
dotnet build src / TaskbarQuota.App / TaskbarQuota.App.csproj - c デバッグ - p:Platform = x64
# 実行
dotnet run -- プロジェクト src / TaskbarQuota.App / TaskbarQuota.App.csproj
# テスト
dotnet テスト テスト / TaskbarQuota.Tests / TaskbarQuota.Tests.csproj
診断ログ:
%TEMP%\taskbarquota.log
手動認証情報 (オプション)
%LOCALAPPDATA%\TaskbarQuota\credentials.json
{
"カーソル" : {
"cookieHeader" : " 名前=値; その他=値 "
}、
"副操縦士" : {
"apiKey" : " github_pat_or_token "
}、
"オープンコード" : {
"cookieHeader" : " 名前=値; その他=値 "
}
}
発売＆配布
署名された Microsoft Store ビルドをインストールします。
GitHub Releases から最新のインストーラーをダウンロードします。
一致する .exe を実行し、[スタート] メニューから TaskbarQuota を起動します。 Windows では、署名されていないビルドに対して SmartScreen が表示される場合があります — 詳細 → とにかく実行 。
Microsoft Store ビルドでは、署名付きパッケージ ID ZiedKallel.TaskbarQuota_q2e4dm2bjnsne を使用します。ストアにインストールされたビルドに更新が表示されると、TaskbarQuota は、署名されていない GitHub インストーラーをダウンロードする代わりに、Microsoft Store の製品ページを開きます。
タスクバークォータ/
§── .github/workflows/release.yml # タグビルド → GitHub リリースインストーラー
§── installer/TaskbarQuota.iss # Inno Setup (x64 + arm64)
§── src/
│ §── taskbarquota.png # README / ブランド ic

に
│ └── TaskbarQuota.App/ # WinUI 3 アプリ
§── testing/TaskbarQuota.Tests/ # 単体テスト
§── TaskbarQuota.slnx
§── ライセンス
━── README.md
ライセンス
TaskbarQuota は MIT ライセンスに基づいてリリースされています。
TaskbarQuota が役立つ場合は、開発のサポートを検討してください。
Windows タスクバーでの AI コーディング ツールの使用法
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
12
タスクバークォータ 1.0.11
最新の
2026 年 6 月 7 日
+ 11 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI coding-tool usage in the Windows taskbar. Contribute to zioder/TaskbarQuota development by creating an account on GitHub.

GitHub - zioder/TaskbarQuota: AI coding-tool usage in the Windows taskbar · GitHub
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
zioder
/
TaskbarQuota
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
27 Commits 27 Commits .github .github installer installer src src tests/ TaskbarQuota.Tests tests/ TaskbarQuota.Tests .gitignore .gitignore LICENSE LICENSE README.md README.md TaskbarQuota.slnx TaskbarQuota.slnx View all files Repository files navigation
A native Windows taskbar companion that automatically detects which AI tool you are using — the focused desktop app or the agent running in your terminal — and shows live usage for that provider only.
Sessions, rate windows, cost or balance; compact widget beside the tray; full dashboard on demand; no cloud backend.
taskbarquota.mp4
TaskbarQuota is a WinUI 3 desktop app that injects a small XAML island into the Windows taskbar (next to the notification area). You do not pick a provider from a list every time you switch tools — the app watches what you are actually doing on Windows and switches the widget, flyout, and active fetch to match.
On first launch after upgrading from the internal WinCheck codename, settings and credentials.json are copied from %LOCALAPPDATA%\WinCheck when the TaskbarQuota folder does not already contain them.
Automatically follows your active tool
TaskbarQuota answers one question: “What am I using right now, and how much quota is left?” It distinguishes two cases:
When you alt-tab from Cursor into a PowerShell session running Claude Code, the widget retargets to Claude within about half a second. When you leave all supported tools, the widget can fade until you return to one. If the foreground app is unrelated (browser, Slack, …), TaskbarQuota keeps showing the last active provider so the taskbar still has context.
Cursor IDE focused → Cursor usage
Windows Terminal + claude in the command line → Claude usage
VS Code focused → GitHub Copilot usage
OpenCode model switch (Zen vs Go) → provider updates without restarting TaskbarQuota
After detection, usage is fetched from local credentials or provider APIs (see Supported providers ). Click the widget for a flyout; open the full window for every provider at once, reset times, and manual credential fixes.
In-taskbar compact UI — usage bars and/or percentages injected beside the system tray (reparents a layered WinUI island into Shell_TrayWnd ).
Follows taskbar layout — recenters when the taskbar moves, centers, or when Widgets / tray geometry changes.
Draggable placement — drag to reposition; offset is saved under %LOCALAPPDATA%\TaskbarQuota\ .
Click for flyout — quick provider strip and settings shortcut above the widget.
Fades when idle — widget can hide when no supported AI tool process is detected; returns when you focus a supported app or terminal session.
Active-tool detection (desktop app or terminal agent)
This is the core behavior — everything else (widget, flyout, fetch cache) hangs off it.
Desktop apps — process name of the focused window: Cursor, Antigravity, Codex, Claude, VS Code / Insiders (Copilot), and similar.
Terminal agents — if a known terminal is focused, WMI/process scan looks at command lines for Claude Code, Codex, cursor-agent , OpenCode, gh copilot , Antigravity CLI, and related launchers (Windows Terminal, PowerShell, pwsh , WezTerm, Alacritty, …).
Fast switching — coordinator polls about every 500 ms , so changing app or shell updates the active provider quickly; usage API calls are cached for 60 seconds so providers are not spammed.
OpenCode model switch — watches OpenCode model state to move between OpenCode Zen and OpenCode Go without restarting the app.
Sticky last provider — unrelated foreground apps do not clear the widget; last detected tool stays until you focus a supported app or terminal again.
Idle hide — optional fade when no supported AI process is running; comes back when you focus a supported app or terminal session.
Provider
What you see
How usage is fetched
Codex
Session and rate windows
OAuth token from %USERPROFILE%\.codex\auth.json (or %CODEX_HOME% ) → ChatGPT usage API
GitHub Copilot
Quota snapshots
GITHUB_TOKEN / GH_TOKEN , saved credentials, or gh auth token → GitHub Copilot internal API
Claude
5-hour, weekly, and model windows
%USERPROFILE%\.claude\.credentials.json or env override → Anthropic OAuth usage API
Antigravity
Local status
Running language_server process → CSRF token and port → 127.0.0.1 status API
Cursor
Usage summary
Cursor state.vscdb , browser cookies, or manual cookie header → Cursor / cursor.com APIs
OpenCode Zen
Billing usage and balance
Browser cookies for opencode.ai or manual cookie → workspace billing pages
OpenCode Go
Rolling, weekly, and monthly windows
Same cookie path as Zen → OpenCode server workspace APIs
Each provider card on the dashboard shows plan name, reset times, rate windows, and optional cost or balance when the API returns it. Use Fix on a card when auto-detection fails to paste tokens or cookie headers into %LOCALAPPDATA%\TaskbarQuota\credentials.json .
Full dashboard — all registered providers at once with refresh, error states, and detail cards.
System tray — Open, show widget, and Exit.
Run at startup — optional Windows startup registration; can launch widget-only.
Settings — light / dark / system theme, widget layout (bars, percentages, or both), consumed vs remaining percentage display.
Local-only — no TaskbarQuota backend; usage calls go directly from your PC to each provider (or localhost for Antigravity).
No telemetry — diagnostics log to %TEMP%\taskbarquota.log only.
Cookie extraction is in-memory — browser cookies are read to build a request header for the active fetch; they are not intentionally persisted (manual credentials in credentials.json are plain JSON today — keep that file on your PC only).
ActiveAppDetector
├─ foreground GUI app → ProviderId
└─ focused terminal → scan CLI processes → ProviderId
↓
UsageCoordinator (500 ms tick, picks active provider)
↓
UsageService (provider registry + 60 s success cache, shorter TTL on errors / 429)
↓
IUsageProvider implementations (HTTP / local API / SQLite / DPAPI)
↓
TaskBarManager → WidgetSummary (taskbar) + Dashboard + Flyout
Detection — ActiveAppDetector maps the foreground process (or CLI child of a known terminal) to a ProviderId .
Fetch — the matching IUsageProvider loads tokens from CLI config files, environment variables, TaskbarQuota credentials, Cursor's local database, or CookieExtractor (Chromium / Firefox profiles, copied to %TEMP% for locked DBs).
Normalize — results become a UsageSnapshot with RateWindow rows and optional CostSnapshot .
Display — UsageCoordinator raises StateChanged ; the taskbar widget and dashboard apply the snapshot. Fetches are cached so providers are not hammered on every tick.
Path
Role
src/TaskbarQuota.App/ActiveApp/
Foreground and CLI provider detection
src/TaskbarQuota.App/Browser/
Browser cookie discovery and Chromium decryption
src/TaskbarQuota.App/Taskbar/
Widget host, tray icon, taskbar layout watcher
src/TaskbarQuota.App/Usage/Providers/
Per-provider fetch logic
src/TaskbarQuota.App/Usage/UsageService.cs
Registry and TTL cache
src/TaskbarQuota.App/UsageCoordinator.cs
Timer-driven detect → fetch → UI update
src/TaskbarQuota.App/Views/
Dashboard and settings
tests/TaskbarQuota.Tests/
Unit tests
Local development
Windows 10 19041 or newer (Windows 11 recommended for taskbar behavior).
.NET SDK 10.0.x (project targets net10.0-windows10.0.19041.0 ).
Windows App SDK / WinUI 3 workload.
You need working sign-in for the providers you test (local OAuth files, GitHub CLI, Cursor install, etc.). No separate server install.
# Build (Debug, x64)
dotnet build src / TaskbarQuota.App / TaskbarQuota.App.csproj - c Debug - p:Platform = x64
# Run
dotnet run -- project src / TaskbarQuota.App / TaskbarQuota.App.csproj
# Tests
dotnet test tests / TaskbarQuota.Tests / TaskbarQuota.Tests.csproj
Diagnostic log:
%TEMP%\taskbarquota.log
Manual credentials (optional)
%LOCALAPPDATA%\TaskbarQuota\credentials.json
{
"cursor" : {
"cookieHeader" : " name=value; other=value "
},
"copilot" : {
"apiKey" : " github_pat_or_token "
},
"opencode" : {
"cookieHeader" : " name=value; other=value "
}
}
Release & distribution
Install the signed Microsoft Store build:
Download the latest installer from GitHub Releases .
Run the matching .exe , then launch TaskbarQuota from the Start menu. Windows may show SmartScreen for unsigned builds — More info → Run anyway .
Microsoft Store builds use the signed package identity ZiedKallel.TaskbarQuota_q2e4dm2bjnsne . When a Store-installed build sees an update, TaskbarQuota opens the Microsoft Store product page instead of downloading the unsigned GitHub installer.
TaskbarQuota/
├── .github/workflows/release.yml # Tag builds → GitHub Release installers
├── installer/TaskbarQuota.iss # Inno Setup (x64 + arm64)
├── src/
│ ├── taskbarquota.png # README / branding icon
│ └── TaskbarQuota.App/ # WinUI 3 app
├── tests/TaskbarQuota.Tests/ # Unit tests
├── TaskbarQuota.slnx
├── LICENSE
└── README.md
License
TaskbarQuota is released under the MIT License .
If TaskbarQuota is useful to you, consider supporting development:
AI coding-tool usage in the Windows taskbar
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
12
TaskbarQuota 1.0.11
Latest
Jun 7, 2026
+ 11 releases
Sponsor this project
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

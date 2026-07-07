---
source: "https://github.com/mayur-25-cd/strays"
hn_url: "https://news.ycombinator.com/item?id=48822092"
title: "Show HN: Strays – See the ports and AI coding sessions your Mac is running"
article_title: "GitHub - mayur-25-cd/strays: Native macOS menu-bar app to watch listening ports grouped by project — and the AI coding sessions that spawned them. Reclaim strays safely. Local-only. · GitHub"
author: "kordio"
captured_at: "2026-07-07T19:43:56Z"
capture_tool: "hn-digest"
hn_id: 48822092
score: 1
comments: 0
posted_at: "2026-07-07T19:01:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Strays – See the ports and AI coding sessions your Mac is running

- HN: [48822092](https://news.ycombinator.com/item?id=48822092)
- Source: [github.com](https://github.com/mayur-25-cd/strays)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:01:10Z

## Translation

タイトル: Show HN: Strays – Mac が実行しているポートと AI コーディング セッションを表示します
記事のタイトル: GitHub - Mayur-25-cd/strays: プロジェクトごとにグループ化されたリスニング ポートと、それらを生成した AI コーディング セッションを監視するためのネイティブ macOS メニュー バー アプリ。迷子を安全に取り戻しましょう。地元限定。 · GitHub
説明: プロジェクトごとにグループ化されたリスニング ポートと、それらを生成した AI コーディング セッションを監視するネイティブ macOS メニュー バー アプリ。迷子を安全に取り戻しましょう。地元限定。 - Mayur-25-cd/strays

記事本文:
GitHub - Mayur-25-cd/strays: プロジェクトごとにグループ化されたリスニング ポートと、それらを生成した AI コーディング セッションを監視するためのネイティブ macOS メニュー バー アプリ。迷子を安全に取り戻しましょう。地元限定。 · GitHub
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
マユール-25-cd
/
はぐれ者
公共
通知
あなたは

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows Casks Casks Packaging Packaging Sources/ Strays Sources/ Strays dist dist docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Mac が実行したままになっているすべてのポートと AI セッションで、迷子を見つけてください。
リッスンしているすべての TCP ポートを表示するネイティブ macOS メニューバー + ウィンドウ アプリ。
元のプロジェクトフォルダーごとにグループ化されており、野良犬を安全に殺すことができます。
ローカル専用、読み取り専用、依存関係なし。
ウィンドウ (プロジェクトごとにグループ化されたポート) · AI セッション インスペクター · メニューバーのポップオーバー
いくつかの開発サーバーを実行します。 AI エージェント (Claude Code、Codex など) がいくつか起動します。
もっと。 1 週間後、12 個のノード、uvicorn、postgres が完成しました。
説明できないポートをリッスンしているプロセス - いくつかは数日間アイドル状態であり、もう 1 つは
密かに 0.0.0.0 にバインドされます。 Strays は、何が実行されているか、どのプロジェクトが実行されているかを示します
どこから来たのか、停止しても安全かどうかが一目でわかります。
brew install --cask Mayur-25-cd/tap/strays
または、リリースから公証された .dmg を取得します。
Strays を「アプリケーション」にドラッグします。
プロジェクト フォルダーごとにグループ化された、リッスンしているすべての TCP ポート。ユビコーン：8000
エージェントが開始すると、myapp/api の下に表示されます — 「これはどこから来たのですか
から？」努力なしで答えました。
それぞれが何であるかを認識します — Vite、Next.js、webpack、Node、uvicorn、
gunicorn、Flask、Postgres、Redis、MySQL、Docker など、それぞれに
カテゴリのグリフ。
露出の認識 — ローカルホスト専用ポートの場合は落ち着いた青緑のロック、オレンジ色
0.0.0.0 上のあらゆるものに対するシールド

/あなたのLAN。赤はキル専用です。
安全な強制終了 - ワンクリックで通常の開発サーバーを正常に停止します (SIGTERM)
4 秒間の元に戻す機能を使用します (ウィンドウが経過するまで信号は送信されないため、
元に戻すのは正直です）。データベース、公開ポート、および強制終了には名前付きの名前が必要です。
確認。システムプロセスが誤って強制終了されることはありません - 中空リング、
ボタンを無効にします。
アイドル/忘れられた検出 - アクティブなサーバーが存在せず、しきい値を超えたサーバー
接続には idle フラグが付けられ、ワンクリックでバッチ リープが行われます。
ポート (⌘L) を解放します — :3000 と入力し、誰がそれを保持しているかを確認し、停止します。
インスペクター内のプロセスごとの CPU、メモリ、ライブ接続。
メニューバーのポップオーバー — メニューバーを開かずに、すべてを確認、フィルター、強制終了
窓。メニューバーのみのモードでは、Dock アイコンが完全に非表示になります。
Strays は、Mac 上で実行されている AI コーディング セッションも明らかにします。
ポート、通常それがポートを生成するものであるためです。
プロジェクトごとにグループ化されたライブ セッション (プロジェクトのポートがネストされている)
その下に。モデル、メッセージ数、および最後のアクティビティを表示します。
ファントムなし - セッションは、プロセス開始時間が記録されている場合にのみ表示されます。
は実際の PID と一致するため、リサイクルされた PID がライブ セッションを偽装することはできません。
増分トランスクリプトテール（バイトオフセット）からの推定コスト（クロード）
デルタ — 完全に再読することはありません）、推定値として明確にラベル付けされています。
読み取り専用およびローカル — セッション ファイルはローカルで読み取られてカウントが取得され、
メタデータのみ。何も保存されたり送信されたりすることはありません。
すべての AI は、サイドバー (⌘2) の 1 つの AI フィルターの下に存在します。ただ欲しいだけ
ポート?すべてのポート (⌘3)。すべて？すべて (⌘1)。
macOS 14 以降では Xcode 16 / Swift 6 が必要です。
# dist/ 内のダブルクリック可能なアプリ バンドル (デフォルトは release、デバッグ ビルドの場合は `debug` を渡します)
./scripts/build-app.sh
dist/Strays.app を開く
# またはデバッグで実行
素早く走る
仕組み
リフレッシュ、解析ごとの低コストのサブプロセス呼び出し

d 機械可読モード - いいえ
昇格された特権、デーモンなし:
Kill は、EPERM を使用して、kill(2) システムコールを直接使用します (SIGTERM / SIGKILL)。
サイレントに失敗するのではなく、「管理者が必要」として表示されます。
サービス/ — PortScanner 、 AISessionScanner 、 ProcessClassifier 、
ProcessKiller 、 ConnectionResolver 、 CommandRunner 。
Models/ — PortEntry (1 行 = 1 つの PID+ポート) および AISession 。
State/PortStore.swift — @MainActor @Observable ビュー モデル: ポーリング ループ、
フィルタリング/グループ化、および kill コーディネーター (deferred-kill Undo、
終了→強制エスカレーション、確認）。
Views/ — メニューバーのポップオーバー、NavigationSplitView メイン ウィンドウ、インスペクター。
Strays は読み取り専用でローカル専用です。他のツールのセッション ファイルを読み取ります。
あなたのマシンがカウントとメタデータ (メッセージ数、モデル、トークン使用量) を取得します。
コンテンツを保存したり送信したりすることはありません。ネットワークもテレメトリもなし
デーモン。現在、それはドットファイル ( ~/.claude 、 ~/.copilot ) のみを読み取ります。
特別な許可。 ~/Library/Application Support を読み取る将来のアダプター
(VS Code、カーソルなど) フルディスク アクセスを求めるプロンプトが表示される場合があります。
サンドボックス化されておらず、Mac App Store にもありません。 lsof / ps へのシェルアウトは
アプリサンドボックスと互換性がありません — Strays は公証されたものとして配布されます
直接ダウンロードとHomebrew Cask。
Strays は独立した、無所属のプロジェクトです。とは提携していませんが、
Anthropic、GitHub/Microsoft、Google、Anysphere によって承認または後援されている、または
ツールが検出される他のベンダー。すべての製品名とブランドは、
はそれぞれの所有者の財産であり、識別のためにのみ使用されます。
貢献を歓迎します - 特に新しいツールアダプター。参照
COTRIBUTING.md および SECURITY.md 。
リリース履歴については、CHANGELOG.md を参照してください。
MIT © 2026 Mayur Parthivaraju.
プロジェクトごとにグループ化されたリスニング ポートを監視するためのネイティブ macOS メニューバー アプリ

ct — およびそれらを生み出した AI コーディング セッション。迷子を安全に取り戻しましょう。地元限定。
github.com/mayur-25-cd/strays/releases
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
ストレイズ 1.0.0
最新の
2026 年 7 月 7 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Native macOS menu-bar app to watch listening ports grouped by project — and the AI coding sessions that spawned them. Reclaim strays safely. Local-only. - mayur-25-cd/strays

GitHub - mayur-25-cd/strays: Native macOS menu-bar app to watch listening ports grouped by project — and the AI coding sessions that spawned them. Reclaim strays safely. Local-only. · GitHub
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
mayur-25-cd
/
strays
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows Casks Casks Packaging Packaging Sources/ Strays Sources/ Strays dist dist docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Find the strays — every port and AI session your Mac left running.
A native macOS menu-bar + window app that shows every listening TCP port,
grouped by the project folder it came from, and lets you kill strays safely.
Local-only, read-only, zero dependencies.
The window (ports grouped by project) · an AI-session inspector · the menu-bar popover
You run a few dev servers. Your AI agent (Claude Code, Codex, …) spins up a few
more. A week later you've got a dozen node , uvicorn , and postgres
processes listening on ports you can't account for — some idle for days, one
quietly bound to 0.0.0.0 . Strays shows you what's running, which project it
came from, and whether it's safe to stop — at a glance.
brew install --cask mayur-25-cd/tap/strays
Or grab the notarized .dmg from Releases
and drag Strays to Applications.
Every listening TCP port, grouped by project folder. The uvicorn :8000
your agent started shows up under myapp/api — "where did this come
from?" answered with zero effort.
Knows what each one is — detects Vite, Next.js, webpack, Node, uvicorn,
gunicorn, Flask, Postgres, Redis, MySQL, Docker, and more, each with a
category glyph.
Exposure awareness — a calm teal lock for localhost -only ports, an amber
shield for anything on 0.0.0.0 / your LAN. Red is reserved only for kill.
Safe kills — one click gracefully stops an ordinary dev server (SIGTERM)
with a 4-second Undo (the signal isn't sent until the window elapses, so
Undo is honest). Databases, exposed ports, and force-kills need a named
confirmation. System processes can't be killed by mistake — hollow ring,
disabled button.
Idle / forgotten detection — servers up past a threshold with no active
connections are flagged idle , with a one-click batch reap.
Free a Port (⌘L) — type :3000 , see who's holding it, stop it.
CPU, memory & live connections per process, in the inspector.
Menu-bar popover — glance at everything, filter, kill, without opening the
window. Menu-bar-only mode hides the Dock icon entirely.
Strays also surfaces the AI coding sessions running on your Mac — beside the
ports, because that's usually what spawned them:
Live sessions grouped by project , with their project's ports nested
underneath. Shows model, message count, and last activity.
No phantoms — a session is shown only if its recorded process-start time
matches the real one, so a recycled PID can never fake a live session.
Estimated cost (Claude) from an incremental transcript tail (byte-offset
deltas — never a full re-read), clearly labelled an estimate.
Read-only & local — session files are read locally to derive counts and
metadata only; nothing is stored or sent anywhere.
Everything AI lives under one AI filter in the sidebar (⌘2). Just want
ports? All Ports (⌘3). Everything? All (⌘1).
Requires Xcode 16 / Swift 6 on macOS 14+.
# double-clickable app bundle in dist/ (defaults to release; pass `debug` for a debug build)
./scripts/build-app.sh
open dist/Strays.app
# or run in debug
swift run
How it works
Cheap subprocess calls per refresh, parsed in machine-readable mode — no
elevated privileges, no daemon:
Kills use the kill(2) syscall directly (SIGTERM / SIGKILL), with EPERM
surfaced as "requires admin" rather than failing silently.
Services/ — PortScanner , AISessionScanner , ProcessClassifier ,
ProcessKiller , ConnectionResolver , CommandRunner .
Models/ — PortEntry (one row = one PID+port) and AISession .
State/PortStore.swift — @MainActor @Observable view model: poll loop,
filtering/grouping, and the kill coordinator (deferred-kill Undo,
terminating→force escalation, confirmations).
Views/ — menu-bar popover, NavigationSplitView main window, inspector.
Strays is read-only and local-only . It reads other tools' session files on
your machine to derive counts and metadata (message count, model, token usage);
it never stores or transmits their contents. No network, no telemetry, no
daemon. Today it reads only dotfiles ( ~/.claude , ~/.copilot ), which need no
special permission; future adapters that read ~/Library/Application Support
(e.g. VS Code, Cursor) may prompt for Full Disk Access.
Not sandboxed / not on the Mac App Store. Shelling out to lsof / ps is
incompatible with the App Sandbox — Strays is distributed as a notarized
direct download and Homebrew cask.
Strays is an independent, unaffiliated project. It is not affiliated with,
endorsed by, or sponsored by Anthropic, GitHub/Microsoft, Google, Anysphere, or
any other vendor whose tools it detects. All product names and brands are the
property of their respective owners and are used solely for identification.
Contributions welcome — especially new tool adapters. See
CONTRIBUTING.md and SECURITY.md .
See CHANGELOG.md for release history.
MIT © 2026 Mayur Parthivaraju.
Native macOS menu-bar app to watch listening ports grouped by project — and the AI coding sessions that spawned them. Reclaim strays safely. Local-only.
github.com/mayur-25-cd/strays/releases
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Strays 1.0.0
Latest
Jul 7, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

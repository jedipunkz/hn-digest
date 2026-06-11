---
source: "https://github.com/manelrv/CCTV"
hn_url: "https://news.ycombinator.com/item?id=48492863"
title: "CCTV – See which Claude Code agent needs you, from the menu bar"
article_title: "GitHub - manelrv/CCTV · GitHub"
author: "manelrv"
captured_at: "2026-06-11T19:07:06Z"
capture_tool: "hn-digest"
hn_id: 48492863
score: 2
comments: 0
posted_at: "2026-06-11T16:52:02Z"
tags:
  - hacker-news
  - translated
---

# CCTV – See which Claude Code agent needs you, from the menu bar

- HN: [48492863](https://news.ycombinator.com/item?id=48492863)
- Source: [github.com](https://github.com/manelrv/CCTV)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T16:52:02Z

## Translation

タイトル: CCTV – メニュー バーから、どのクロード コード エージェントがあなたを必要としているかを確認します
記事タイトル: GitHub - manelrv/CCTV · GitHub
説明: GitHub でアカウントを作成して、manelrv/CCTV の開発に貢献します。

記事本文:
GitHub - manelrv/CCTV · GitHub
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
マネルヴ
/
監視カメラ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミ

ts 32 コミット .github/ workflows .github/ workflows docs docs フック フック アイコン アイコン src-tauri src-tauri src src .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md WORKLOG.md WORKLOG.md Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
どの Claude Code エージェントがあなたを必要としているかを見つけるために、Alt キーを押しながら Tab キーを押しながら端末を移動するのはやめてください。
CCTV は、すべてのクロード コード インスタンスのライブ状態を表示する macOS メニュー バー アプリです。
あなたのマシン上で — 作業中、承認待ち、入力待ち、完了、または失敗 —
常に一番上に表示される 1 つのフローティング ウィンドウで。あなたを必要とするエージェントがあなたのところに来ます。
Cmd+Shift+Space は、あなたを必要とする次のエージェントに直接フォーカスを移動します。
複数の Claude Code エージェントを同時に実行することは、複数の端末を子守することを意味します - チェック
タブを次々と押して、許可プロンプトまたは質問でどれがブロックされているかを確認します。監視カメラ
それを一目でわかるように折りたたむと、動的なトレイ アイコンがカウンター付きの警告に変わります。
ブロックされたエージェントがリストの先頭に表示され、グローバル ホットキーで直接アクセスできます
そのターミナルへ。
brew install --cask manelrv/tap/cctv
xattr -dr com.apple.quarantine /Applications/CCTV.app
現時点では 2 行目が必要です: CCTV はまだ署名/公証されていないため、macOS
ゲートキーパーはインストール時にそれを隔離します。その属性をクリアすると起動できるようになります —
または、アプリを右クリック→「開く」→「初回開く」の順にクリックします。 (公証は有効です
ロードマップ。着地すると、余分なステップは消えます。) .dmg は次のように構築されます。
タグ付けされたすべてのリリースのリリース パイプライン。
git clone https://github.com/manelrv/CCTV.git
CD CCTV
npmインストール
npm run tauri build # src-tauri/target の下に .app / .dmg を生成します

/リリース/バンドル
または、開発環境で実行します。
npmインストール
npm run tauri dev
要件: Node.js 18 以降および Tauri 2 ツールチェーン
(安定した Rust + プラットフォーム依存)。
フックを接続する (フォアグラウンドセッション)
バックグラウンド セッション ( claude --bg ) にはセットアップが必要ありません。通常の前景も表示するには
セッション、フックキーをマージします
フック/settings.snippet.json を
~/.claude/settings.json 。クロード コード内では、/hooks でそれらをリストする必要があります。
ソース ユーザー 。それ以降、開いたすべてのセッションがウィンドウに表示されます。
ハイブリッド デュアル ソース モニタリング — 両方の種類のセッションをカバーします。
バックグラウンド ( claude --bg 、エージェント ビュー): ファイル ウォッチャーが state.json を読み取ります。
クロード コード スーパーバイザーが ~/.claude/jobs/ の下に保存するファイル。設定ゼロ。
フォアグラウンド (ターミナルの通常のクロード): ローカル サーバーへの HTTP フック POST
アプリに埋め込まれています。エンドポイントは即座に応答し、セッションの速度が低下することはありません。
グローバル ホットキー - Cmd+Shift+Space サイクルで、必要なエージェントに焦点を絞ります。
緊急の順に、各ホスティング端末を最前面に持ってきます。
緊急度順リスト — 許可待ち > 入力待ち > エラー >
動作中 > 信号なし > アイドル > 完了。あなたを必要としている列は常に一番上にあります、
保留中の質問を追加するか、ツール: パスを詳細として承認します。
動的なトレイ アイコン — 冷静/警戒のバリエーションと、数値による注意カウンター
メニューバー。
クリックしてフォーカスまたはコピーします。行をクリックすると、その端子タブにジャンプします。 Alt キーを押しながらクリック (または
バックグラウンド行) は、代わりに claudeattach <id> / 作業ディレクトリをコピーします。
セッションごとのコンテキスト メーター — エージェントごとのコンテキスト ウィンドウの占有率を表示します。
オレンジ/赤のしきい値で、誰が部屋を使い果たしそうになっているかがわかります。
デスクトップ通知 — エージェントが必要になった瞬間に、エージェントごとに 1 つのアラートが送信されます。スパムはありません。
全画面アプリの上にフロートします — 非アクティブ化 NSPanel : ev で表示されます

エリースペース、
アクティブなアプリからフォーカスを奪うことはありません。
すべての設定が有効です - フローティング ウィンドウ、常に最前面に表示、自動非表示、コンパクト モード、
ログイン時に開く、テーマ (システム/ダーク/ライト)、不透明度、言語。
8 言語 — 英語 (デフォルト)、スペイン語、ポルトガル語、ドイツ語、フランス語、イタリア語、
カタルーニャ語、ロシア語。自動検出されるか、トレイの言語サブメニューから固定されます。
クロード コード — bg セッション クロード コード — fg セッション
~/.claude/jobs/<id>/state.json HTTP フック (POST localhost:8787)
│ ファイルウォッチャー │
▼ ▼
CCTV（常時生存トレイプロセス）
§── jobs.rs → ソース A: 監視 + 監視状態ファイルの解析
§──server.rs → ソースB:フックレシーバ(axum)
§── state.rs → ハイブリッド ストア、「バックグラウンド優先」マージ ルール、TTL リーパー
§──fresh.rs → 単一の伝播ポイント: Web ビュー、トレイ アイコン、自動非表示
└── webview → React フローティング ウィンドウ (イベント駆動型、ポーリングなし)
詳細については、 CLAUDE.md 、 docs/DATA-SOURCES.md を参照してください。
および docs/ARCHITECTURE.md 。
npm run tauri dev # Vite + Rust、アプリを開きます
cd src-tauri && カーゴテスト # Rust テストスイート
プロジェクトのステータス
macOS 上で機能的に完成しています。実際のセッションに対して検証されたハイブリッド ソース、
ステート マシンとリーパーはユニット + ライブキル テストでカバーされ、すべてを備えた動的トレイ
設定はライブ、i18n (自動検出 + 手動)、グローバル ホットキー、およびフルスクリーン フロートです。
57 件のテストに合格しました。 Homebrew 経由でインストール可能。リリースパイプラインにはユニバーサルが付属しています
すべてのタグに .dmg が付けられます。
次は署名 + 公証 (xattr ステップを削除するため)、次に Linux/Wayland
(Hyprland ルールは docs/ARCHITECTURE.md で草案化されています) および Windows。
履歴は WORKLOG.md 、フェーズ追跡は docs/ROADMAP.md にあります。
macOS (プライマリ) → Linux → Windows。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
CCTV v0.1.2
最新の
2026 年 6 月 11 日
+2リリース

セス
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to manelrv/CCTV development by creating an account on GitHub.

GitHub - manelrv/CCTV · GitHub
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
manelrv
/
CCTV
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits .github/ workflows .github/ workflows docs docs hooks hooks icons icons src-tauri src-tauri src src .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md WORKLOG.md WORKLOG.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Stop alt-tabbing through terminals to find out which Claude Code agent needs you.
CCTV is a macOS menu-bar app that shows the live state of every Claude Code instance
on your machine — working, waiting for your approval, waiting for input, done, or failed —
in one always-on-top floating window. The agent that needs you comes to you.
Cmd+Shift+Space jumps focus straight to the next agent that needs you.
Running several Claude Code agents at once means babysitting several terminals — checking
tab after tab to see which one is blocked on a permission prompt or a question. CCTV
collapses that into a single glance: a dynamic tray icon turns to alert with a counter,
the blocked agent floats to the top of the list, and a global hotkey takes you straight
to its terminal.
brew install --cask manelrv/tap/cctv
xattr -dr com.apple.quarantine /Applications/CCTV.app
The second line is needed for now : CCTV isn't signed/notarized yet, so macOS
Gatekeeper quarantines it on install. Clearing that attribute lets it launch —
or right-click the app → Open → Open the first time. (Notarization is on
the roadmap; once it lands, the extra step disappears.) The .dmg is built by
the release pipeline on every tagged release.
git clone https://github.com/manelrv/CCTV.git
cd CCTV
npm install
npm run tauri build # produces a .app / .dmg under src-tauri/target/release/bundle
Or run it in development:
npm install
npm run tauri dev
Requirements: Node.js 18+ and the Tauri 2 toolchain
(stable Rust + platform deps).
Connect the hooks (foreground sessions)
Background sessions ( claude --bg ) need no setup. To also see regular foreground
sessions, merge the hooks key from
hooks/settings.snippet.json into your
~/.claude/settings.json . Inside Claude Code, /hooks should then list them with
source User . From then on every session you open appears in the window.
Hybrid dual-source monitoring — covers both kinds of sessions:
Background ( claude --bg , Agent View): a file watcher reads the state.json
files the Claude Code supervisor persists under ~/.claude/jobs/ . Zero config.
Foreground (regular claude in a terminal): HTTP hooks POST to a local server
embedded in the app. Endpoints respond instantly and never slow your sessions down.
Global hotkey — Cmd+Shift+Space cycles focus through the agents that need you,
in urgency order, bringing each hosting terminal to the front.
Urgency-ordered list — waiting for permission > waiting for input > error >
working > no signal > idle > completed. The row that needs you is always on top,
with the pending question or approve Tool: path as detail.
Dynamic tray icon — calm/alert variants plus a numeric attention counter in the
menu bar.
Click to focus or copy — click a row to jump to its terminal tab; Alt-click (or
background rows) copies claude attach <id> / the working directory instead.
Per-session context meter — shows context-window occupancy per agent, with
amber/red thresholds, so you see who's about to run out of room.
Desktop notifications — one alert per agent the moment it starts needing you. No spam.
Floats above fullscreen apps — a non-activating NSPanel : visible on every Space,
never steals focus from your active app.
All preferences live — floating window, always on top, auto-hide, compact mode,
open at login, theme (system/dark/light), opacity, and language.
8 languages — English (default), Spanish, Portuguese, German, French, Italian,
Catalan, Russian. Auto-detected, or pinned from the tray's Language submenu.
Claude Code — bg sessions Claude Code — fg sessions
~/.claude/jobs/<id>/state.json HTTP hooks (POST localhost:8787)
│ file watcher │
▼ ▼
CCTV (always-alive tray process)
├── jobs.rs → source A: watch + parse supervisor state files
├── server.rs → source B: hook receiver (axum)
├── state.rs → hybrid store, "background wins" merge rule, TTL reaper
├── refresh.rs → single propagation point: webview, tray icon, auto-hide
└── webview → React floating window (event-driven, no polling)
Full details in CLAUDE.md , docs/DATA-SOURCES.md
and docs/ARCHITECTURE.md .
npm run tauri dev # Vite + Rust, opens the app
cd src-tauri && cargo test # Rust test suite
Project status
Functionally complete on macOS. Hybrid sources verified against real sessions,
state machine and reaper covered by unit + live-kill tests, dynamic tray with all
preferences live, i18n (auto-detect + manual), global hotkey, and fullscreen float.
57 passing tests. Installable via Homebrew; the release pipeline ships a universal
.dmg on every tag.
Next up: signing + notarization (to drop the xattr step), then Linux/Wayland
(Hyprland rules drafted in docs/ARCHITECTURE.md ) and Windows.
History in WORKLOG.md , phase tracking in docs/ROADMAP.md .
macOS (primary) → Linux → Windows.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
CCTV v0.1.2
Latest
Jun 11, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

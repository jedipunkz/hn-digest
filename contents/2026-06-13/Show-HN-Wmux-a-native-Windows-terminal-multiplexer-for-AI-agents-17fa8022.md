---
source: "https://github.com/openwong2kim/wmux"
hn_url: "https://news.ycombinator.com/item?id=48517302"
title: "Show HN: Wmux – a native Windows terminal multiplexer for AI agents"
article_title: "GitHub - openwong2kim/wmux: Windows tmux alternative for AI agents — split terminals for Claude Code, Codex, Gemini CLI with MCP browser automation. No WSL required. · GitHub"
author: "wong2kim"
captured_at: "2026-06-13T15:38:40Z"
capture_tool: "hn-digest"
hn_id: 48517302
score: 3
comments: 1
posted_at: "2026-06-13T13:46:19Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wmux – a native Windows terminal multiplexer for AI agents

- HN: [48517302](https://news.ycombinator.com/item?id=48517302)
- Source: [github.com](https://github.com/openwong2kim/wmux)
- Score: 3
- Comments: 1
- Posted: 2026-06-13T13:46:19Z

## Translation

タイトル: Show HN: Wmux – AI エージェント用のネイティブ Windows ターミナル マルチプレクサ
記事のタイトル: GitHub - openwong2kim/wmux: AI エージェント用の Windows tmux 代替 — MCP ブラウザ自動化を備えた Claude Code、Codex、Gemini CLI 用の分割ターミナル。 WSLは必要ありません。 · GitHub
説明: AI エージェントの Windows tmux 代替 — MCP ブラウザ自動化を備えた Claude Code、Codex、Gemini CLI 用の分割ターミナル。 WSLは必要ありません。 - openwong2kim/wmux

記事本文:
GitHub - openwong2kim/wmux: AI エージェント用の Windows tmux 代替 — MCP ブラウザ自動化を備えた Claude Code、Codex、Gemini CLI 用の分割ターミナル。 WSLは必要ありません。 · GitHub
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
オープンウォン2キム
/
wmux
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
522 コミット 522 コミット .claude-plugin .claude-plugin .github .github アセット アセット ベンチ ベンチ ビルド チョコレートチョコレート ドキュメント ドキュメントの例 例 統合 統合計画 計画 スクリプト スクリプト src src .eslintrc.json .eslintrc.json .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md THIRD_PARTY_NOTICES THIRD_PARTY_NOTICES TODOS.md TODOS.md Decisions.md Decisions.md forge.config.ts forge.config.ts forge.env.d.ts forge.env.d.ts handoff.md handoff.md Index.html Index.html install.ps1 install.ps1 package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js progress.md progress.md tailwind.config.js tailwind.config.js tsconfig.cli.json tsconfig.cli.json tsconfig.daemon.json tsconfig.daemon.json tsconfig.json tsconfig.json tsconfig.mcp.json tsconfig.mcp.json vite.main.config.ts vite.main.config.ts vite.preload.config.ts vite.preload.config.ts vite.renderer.config.ts vite.renderer.config.ts vitest.config.ts vitest.config.ts vitest.runtime.config.ts vitest.runtime.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
wmux — AI エージェント用の Windows ターミナル マルチプレクサ (cmux の代替)
wmux は端末用の LSP であり、外部ツールが端末セッションの上にワークフロー インテリジェンスを構築できる中立的な基盤です。
分割ペイン、MCP ブリッジ、およびブラウザ自動化を備えたネイティブ Windows ターミナル マルチプレクサ - Claude Code、Codex CLI、および Gemini CLI を並行して実行するために専用に構築されています。 WSLは必要ありません。
キーワード: Windows tmux、tmux for Windows、ターミナル マルチプレクサ Wind

ows、AI エージェント ターミナル、Claude Code Windows、Codex CLI、Gemini CLI、MCP サーバー、Chrome DevTools プロトコル、スプリット ターミナル、マルチエージェント ターミナル、ブラウザ自動化、ConPTY、xterm.js、Electron ターミナル。
Windows 上で AI コーディング エージェント用に 1 つのターミナルをまだ実行していますか?
Windows にはネイティブ tmux がありません。 WSL がなければ、複数の AI コーディング エージェントを並行して実行するきれいな方法はありませんでした。
wmux はこれを修正します。ネイティブ Windows ターミナル マルチプレクサ + ブラウザ自動化 + MCP サーバー。Claude Code、Codex CLI、Gemini CLI などの AI コーディング エージェント専用に構築されています。 AI エージェントは端末を読み取り、ブラウザを制御し、すべて 1 つのウィンドウ内で自律的に動作します。
Claude Code は左側にバックエンドを記述します
Codex は右側にフロントエンドを構築します
Gemini CLI は下部でテストを実行します
— すべてを 1 つの画面上に同時に表示します。
スターの歴史
Winget (推奨 — セキュリティ警告なし):
winget インストール openwong2kim.wmux
チョコレート味:
ちょこwmuxをインストール
インストーラー: wmux Setup.exe をダウンロードします。
「Windows が PC を保護しました」という警告が表示されましたか?インストーラーはまだ Authenticode 署名されていないため (SignPath による無料のコード署名がセットアップ中です。以下の「コード署名」を参照)、Windows SmartScreen は不明な発行元からのものとしてフラグを立てます。安全に続行できます。[詳細情報] → [とにかく実行] をクリックします。上記の winget または Chocolatey 経由でインストールすると、これらのパッケージ マネージャーは信頼されたコンテキストで実行されるため、このプロンプトは完全に回避されます。
「とにかく実行」オプションがないため、インストーラーがブロックされましたか? Windows 11 デバイスで Smart App Control (SAC) が有効になっている場合、署名のないインストーラーは完全にブロックされる可能性があります (SmartScreen ダイアログやオーバーライドはありません)。 SAC の適用は SmartScreen とは別個であり、Microsoft のクラウド レピュテーションに依存しているため、ブロックは一時的なものになる可能性があります。ハッシュがレピュテーションを獲得すると、ローカルで変更を加えることなく、数時間後に同じバイナリが正常にインストールされる可能性があります。

s ( #200 )。 SAC が原因かどうかを確認します。
Get-MpComputerStatus |オブジェクトの選択 SmartAppControlState
On と報告される場合、オプションは次のとおりです: 上記の winget または Chocolatey 経由でインストールするか、後でインストーラーを再試行するか、ソースからビルドするか (下記を参照)。ブロックされた試行では、イベント ビューアーにコード整合性イベント ID 3077 も記録されます。
ワンライナー (PowerShell): 最新リリースから事前にビルドされた Setup.exe をダウンロードし、起動する前に SHA-256 を検証します (ノード/Python/ビルド ツールは必要ありません)。
irm https://raw.githubusercontent.com/openwong2kim/wmux/main/install.ps1 |アイエックス
代わりにソースからビルドしたいですか (Node 18+、Python 3、および VS C++ Build Tools (自動インストールされる) が必要です)?
$ 環境: WMUX_FROM_SOURCE = 1 ; irm https://raw.githubusercontent.com/openwong2kim/wmux/main/install.ps1 |アイエックス
コード署名 — Windows ビルド用の無料の Authenticode コード署名は、SignPath.io によって提供され、証明書は SignPath Foundation によって提供されます。
1. AI エージェントがブラウザを制御します — 実際に
Claude Code に「これを Google で検索してください」と指示すると、実際に検索が行われます。 wmux の組み込みブラウザは、Chrome DevTools Protocol (CDP) 経由で接続します。クリック、入力、スクリーンショット、JS の実行はすべて AI によって直接行われます。 React 制御の入力および CJK テキスト (韓国語、日本語、中国語) で完全に動作します。
あなた: 「Google で wmux を検索してください」
クロード: ブラウザーオープン → ブラウザースナップショット → ブラウザーフィル(ref=13, "wmux") → ブラウザープレスキー("Enter")
→実際にGoogle検索してみます。終わり。
2. 1 つのウィンドウに複数の端末を表示
Ctrl+D で分割し、Ctrl+N で新しいワークスペースを作成します。各ワークスペースに複数の端末とブラウザを配置します。 Ctrl キーを押しながらクリックするとマルチビューになります。複数のワークスペースを一度に表示します。
ConPTY ベースのネイティブ Windows 疑似端末。 xterm.js + WebGL ハードウェア アクセラレーションによるレンダリング。 999K 行のスクロールバック。ターミナルのコンテンツはap後も保持されます

再起動します。
Tmux スタイルのプレフィックス モード — Ctrl+B に続いてシングル アクション キー (分割、フォーカス、新しいワークスペースなど) を押してマッスル メモリ ターミナル ナビゲーションを実行します。 13 のアクション、完全に再バインド可能。
フローティング ペイン — メイン レイアウトの外側にある Quake スタイルのドロップダウン ターミナルの場合は Ctrl+`。トグルを切り替えても存続します。
スクロール ブックマーク — Ctrl+M で現在のスクロール位置をマークします。マーカーが側溝にレンダリングされるので、どこにいたのかがわかります。
スマート右クリック — Windows ターミナル スタイル。選択→即時コピー。空き領域→インスタントペースト。リンク → 小さい [リンクを開く/コピー] メニュー。モーダル割り込みゼロ。
3. 「もう終わりましたか?」と尋ねる必要はもうありません。
wmux は、AI エージェントが終了すると通知します。
タスク完了 → デスクトップ通知 + タスクバーのフラッシュ
異常終了→即時警告
git Push --force 、 rm -rf 、 DROP TABLE → 危険なアクションの検出
パターン マッチングではありません - 出力スループット ベースの検出。どのエージェントでも動作します。
4. クロードコード (MCP) の自動統合
wmux を起動すると、MCP サーバーが自動的に登録されます。クロードコードは正常に機能します:
マルチエージェント: すべてのブラウザ ツールは surfaceId を受け入れます。各クロード コード セッションは独自のブラウザを独立して制御します。 A2A (エージェント間) ツールは、同じ wmux インスタンス内の兄弟エージェント間でメッセージをルーティングします。
プログラムによるオーケストレーション: 独自のスクリプトまたはエージェントから複数の wmux ペインを駆動するには、MCP ツール (terminal_send 、terminal_read 、pane_list 、wmux_events_poll ) または wmux CLI ( wmux send / read-screen / list-panes 、 --json 経由でスクリプト可能) を直接呼び出します。サブストレートはオーケストレーション サーフェスです。アトミック ペイン クレームは、expectedVersion ガード (サーバー強制のオプティミスティック同時実行) を備えた pane_set_metadata です。
5. セッションの永続性 — tmux スタイル、終了、クラッシュ、再起動後も存続します
wmux は、ウィンドウを閉じても作業を強制終了することはありません。のように

tmux サーバーでは、バックグラウンド デーモンがすべての PTY を所有し、実行し続けます。 UI は接続と取り外しのみを行います。 wmux を終了して再度開きます。スクロールバック テキストだけでなく、セッション、プロセス、その他すべて (飛行中のビルド、vim バッファ、ssh シェル) はまだ実行中です。
終了 / アプリの再起動: wmux を閉じるとデーモンから切り離され、すべての PTY プロセスが生きたままになります。次回起動すると、各ペインが即座にライブ セッションに再接続されます。
クラッシュ: 同じパス — デーモンは UI よりも存続するため、レンダラーやメインプロセスのクラッシュによって端末が破壊されることはありません。
再起動: PTY は電源を入れ直しても存続できないため、代わりに保存されたスクロールバック + 作業ディレクトリからセッションが復元されます。 wmux はログイン時に自動起動します。
完全な分解: すべてを実際に停止したい場合は、トレイの「wmux をシャットダウン (すべてのセッションを閉じる)」を使用します。すべてのペインを閉じると、数分後にデーモンが自動的にアイドル状態でシャットダウンされます。
自動更新: GitHub リリース経由で更新を確認します。設定でオン/オフを切り替えます。
6. 実際に重要なセキュリティ
すべての IPC パイプでのトークン認証
SSRF 保護 — プライベート IP、 file:// 、 javascript: スキームをブロックします
PTY 入力のサニタイズ — コマンド インジェクションを防止します
ランダム化された CDP ポート — 固定デバッグ ポートなし
メモリプレッシャーウォッチドッグ — 750MBでデッドセッションを獲得し、1GBで新しいセッションをブロックします
電子ヒューズ — Cookie 暗号化がオン。ノード CLI 検査引数と NODE_OPTIONS 環境が無効になります。アプリは asar からのみロードされます。 ( RunAsNode は有効のままです。バックグラウンド デーモンは、 ELECTRON_RUN_AS_NODE=1 を介して wmux.exe から切り離されたノード プロセスとして生成されます。)
xterm.js + WebGL GPU アクセラレーションによるレンダリング
ConPTY ネイティブ Windows 疑似端末
Unicode 11 幅テーブル — カーソル位置 TUI の正しい CJK / 絵文字レンダリング (Claude Code、vim)
ペインの分割 — Ctrl+D 水平、Ctrl+Shift+D 垂直
タブ — マル

ペインごとのチップルサーフェス
フローティング ペイン — Quake スタイルのドロップダウン ターミナル、専用 PTY、Ctrl+`
スマートな右クリック - 選択→インスタントコピー、空の領域→インスタントペースト、リンク→開く/リンクコピーメニュー
スクロールブックマーク - Ctrl+Mマーク、余白インジケーター
正規表現による検索切り替え - Ctrl+F
ディスク永続性による 999K 行のスクロールバック
シェル統合 (OSC 133) — Terminal_read_events のセマンティック プロンプト/コマンド境界。 pwsh / bash に自動挿入されます。制約付き言語モードは安全です (v2.7.1)。
Tmux スタイルのプレフィックス モード — Ctrl+B でアクション キー、13 のデフォルト アクション (分割、フォーカス、ワークスペース、パレット、フローティング ペインなど)
設定でカスタマイズ可能なバインディング + カスタム キーマップ
ドラッグアンドドロップで並べ替えできるサイドバー
マルチビュー — Ctrl キーを押しながらクリックすると、複数のワークスペースを並べて表示できます。
レイアウト テンプレート - 現在のペイン レイアウトを保存し、コマンド パレット (「最近」カテゴリ) 経由で復元します。
完全なセッション永続性 — レイアウト、タブ、CWD、スクロールバックがすべて復元されました
内蔵ブラウザパネル — Ctrl+Shift+L
ナビゲーション バー、DevTools、戻る/進む
要素インスペクター — マウスを移動して強調表示し、クリックして LLM に適したコンテキストをコピーします
CDP の完全な自動化: クリック、入力、入力、スクリーンショット、JS 評価、キーの押下
出力スループットベースのアクティビティ検出
タスクバーのフラッシュ + Windows トースト通知
クロード コード、カーソル、Aider、Codex CLI、Gemini CLI、OpenCode、GitHub C

[切り捨てられた]

## Original Extract

Windows tmux alternative for AI agents — split terminals for Claude Code, Codex, Gemini CLI with MCP browser automation. No WSL required. - openwong2kim/wmux

GitHub - openwong2kim/wmux: Windows tmux alternative for AI agents — split terminals for Claude Code, Codex, Gemini CLI with MCP browser automation. No WSL required. · GitHub
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
openwong2kim
/
wmux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
522 Commits 522 Commits .claude-plugin .claude-plugin .github .github assets assets bench bench build build chocolatey chocolatey docs docs examples examples integrations integrations plans plans scripts scripts src src .eslintrc.json .eslintrc.json .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md THIRD_PARTY_NOTICES THIRD_PARTY_NOTICES TODOS.md TODOS.md decisions.md decisions.md forge.config.ts forge.config.ts forge.env.d.ts forge.env.d.ts handoff.md handoff.md index.html index.html install.ps1 install.ps1 package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js progress.md progress.md tailwind.config.js tailwind.config.js tsconfig.cli.json tsconfig.cli.json tsconfig.daemon.json tsconfig.daemon.json tsconfig.json tsconfig.json tsconfig.mcp.json tsconfig.mcp.json vite.main.config.ts vite.main.config.ts vite.preload.config.ts vite.preload.config.ts vite.renderer.config.ts vite.renderer.config.ts vitest.config.ts vitest.config.ts vitest.runtime.config.ts vitest.runtime.config.ts View all files Repository files navigation
wmux — Windows Terminal Multiplexer for AI Agents (cmux alternative)
wmux is LSP-for-terminals — a neutral substrate that lets external tools build workflow intelligence on top of any terminal session.
Native Windows terminal multiplexer with split panes, MCP bridge, and browser automation — purpose-built for running Claude Code, Codex CLI, and Gemini CLI side by side. No WSL required.
Keywords: Windows tmux, tmux for Windows, terminal multiplexer Windows, AI agent terminal, Claude Code Windows, Codex CLI, Gemini CLI, MCP server, Chrome DevTools Protocol, split terminal, multi-agent terminal, browser automation, ConPTY, xterm.js, Electron terminal.
Still running one terminal for your AI coding agents on Windows?
Windows has no native tmux. Without WSL, there was no clean way to run multiple AI coding agents side by side.
wmux fixes this. A native Windows terminal multiplexer + browser automation + MCP server, purpose-built for AI coding agents like Claude Code, Codex CLI, and Gemini CLI. Your AI agent reads the terminal, controls the browser, and works autonomously — all in one window.
Claude Code writes the backend on the left
Codex builds the frontend on the right
Gemini CLI runs tests at the bottom
— all on one screen, simultaneously.
Star History
Winget (recommended — no security warning):
winget install openwong2kim.wmux
Chocolatey:
choco install wmux
Installer: Download wmux Setup.exe
Seeing a "Windows protected your PC" warning? The installer is not yet Authenticode-signed (free code signing via SignPath is being set up — see Code signing below), so Windows SmartScreen flags it as from an unknown publisher. It's safe to proceed — click More info → Run anyway . Installing via winget or Chocolatey above avoids this prompt entirely, since those package managers run in a trusted context.
Installer blocked with no "Run anyway" option? If Smart App Control (SAC) is enabled on your Windows 11 device, the unsigned installer may be blocked outright — no SmartScreen dialog, no override. SAC enforcement is separate from SmartScreen and relies on Microsoft's cloud reputation, so a block can be transient: the same binary may install successfully hours later once the hash gains reputation, with no local changes ( #200 ). Check whether SAC is the cause:
Get-MpComputerStatus | Select-Object SmartAppControlState
If it reports On , your options are: install via winget or Chocolatey above, retry the installer later, or build from source (see below). Blocked attempts also log Code Integrity Event ID 3077 in Event Viewer.
One-liner (PowerShell): downloads the prebuilt Setup.exe from the latest release and verifies its SHA-256 before launching it (no Node/Python/build tools needed).
irm https: // raw.githubusercontent.com / openwong2kim / wmux / main / install.ps1 | iex
Want to build from source instead (needs Node 18+, Python 3, and VS C++ Build Tools — auto-installed)?
$ env: WMUX_FROM_SOURCE = 1 ; irm https: // raw.githubusercontent.com / openwong2kim / wmux / main / install.ps1 | iex
Code signing — Free Authenticode code signing for Windows builds is provided by SignPath.io , certificate by the SignPath Foundation .
1. Your AI agent controls the browser — for real
Tell Claude Code "search Google for this" and it actually does it. wmux's built-in browser connects via Chrome DevTools Protocol (CDP). Click, type, screenshot, execute JS — all done by the AI directly. Works perfectly with React controlled inputs and CJK text (Korean, Japanese, Chinese).
You: "Search for wmux on Google"
Claude: browser_open → browser_snapshot → browser_fill(ref=13, "wmux") → browser_press_key("Enter")
→ Actually searches Google. Done.
2. Multiple terminals in one window
Ctrl+D to split, Ctrl+N for new workspace. Place multiple terminals and browsers in each workspace. Ctrl+click for multiview — see multiple workspaces at once.
ConPTY-based native Windows pseudo-terminal. xterm.js + WebGL hardware-accelerated rendering. 999K lines of scrollback. Terminal content persists even after app restart.
Tmux-style prefix mode — Ctrl+B then a single action key (split, focus, new workspace…) for muscle-memory terminal navigation. 13 actions, fully rebindable.
Floating pane — Ctrl+` for a Quake-style dropdown terminal that lives outside your main layout. Stays alive across toggles.
Scroll bookmarks — Ctrl+M marks the current scroll position. Markers render on the gutter so you can see where you've been.
Smart right-click — Windows Terminal style. Selection → instant copy. Empty area → instant paste. Link → small Open / Copy Link menu. Zero modal interrupts.
3. No more asking "is it done yet?"
wmux tells you when your AI agent finishes.
Task complete → desktop notification + taskbar flash
Abnormal exit → immediate warning
git push --force , rm -rf , DROP TABLE → dangerous action detection
Not pattern matching — output throughput-based detection. Works with any agent.
4. Automatic Claude Code (MCP) integration
Launch wmux and the MCP server registers automatically. Claude Code just works:
Multi-agent: Every browser tool accepts surfaceId — each Claude Code session controls its own browser independently. A2A (agent-to-agent) tools route messages between sibling agents in the same wmux instance.
Programmatic orchestration: To drive multiple wmux panes from your own script or agent, call the MCP tools ( terminal_send , terminal_read , pane_list , wmux_events_poll ) or the wmux CLI ( wmux send / read-screen / list-panes , scriptable via --json ) directly. The substrate is the orchestration surface — atomic pane claim is pane_set_metadata with an expectedVersion guard (server-enforced optimistic concurrency).
5. Session persistence — tmux-style, survives Quit, crash, and reboot
wmux never kills your work when you close the window. Like a tmux server, the background daemon owns every PTY and keeps it running; the UI only attaches and detaches. Quit wmux and reopen — your sessions are still running, processes and all (a build mid-flight, a vim buffer, an ssh shell), not just scrollback text.
Quit / app restart: Closing wmux detaches from the daemon, which keeps every PTY process alive. The next launch reattaches each pane to its live session instantly.
Crash: Same path — the daemon outlives the UI, so a renderer or main-process crash never takes your terminals with it.
Reboot: PTYs can't survive a power cycle, so sessions are restored from saved scrollback + working directory instead. wmux auto-starts on login.
Full teardown: When you want everything to actually stop, use the tray's "Shut down wmux (close all sessions)" . Closing every pane also lets the daemon idle-shut-down on its own after a few minutes.
Auto-update: Checks for updates via GitHub Releases. Toggle on/off in Settings.
6. Security that actually matters
Token authentication on all IPC pipes
SSRF protection — blocks private IPs, file:// , javascript: schemes
PTY input sanitization — prevents command injection
Randomized CDP port — no fixed debug port
Memory pressure watchdog — reaps dead sessions at 750MB, blocks new ones at 1GB
Electron Fuses — cookie encryption on; Node CLI inspect args and NODE_OPTIONS env disabled; app loads only from asar. ( RunAsNode stays enabled — the background daemon is spawned as a detached Node process from wmux.exe via ELECTRON_RUN_AS_NODE=1 .)
xterm.js + WebGL GPU-accelerated rendering
ConPTY native Windows pseudo-terminal
Unicode 11 width tables — correct CJK / emoji rendering for cursor-positioning TUIs (Claude Code, vim)
Split panes — Ctrl+D horizontal, Ctrl+Shift+D vertical
Tabs — multiple surfaces per pane
Floating pane — Quake-style dropdown terminal, dedicated PTY, Ctrl+`
Smart right-click — selection → instant copy, empty area → instant paste, link → Open / Copy Link menu
Scroll bookmarks — Ctrl+M mark, gutter indicators
Search with regex toggle — Ctrl+F
999K line scrollback with disk persistence
Shell integration (OSC 133) — semantic prompt / command boundaries for terminal_read_events . Auto-injected for pwsh / bash. Constrained Language Mode safe (v2.7.1).
Tmux-style prefix mode — Ctrl+B then action key, 13 default actions (splits, focus, workspaces, palette, floating pane, …)
Customizable bindings + custom keymaps in Settings
Sidebar with drag-and-drop reordering
Multiview — Ctrl+click to view multiple workspaces side by side
Layout templates — save current pane layout, restore via Command Palette ("recent" category)
Full session persistence — layout, tabs, cwd, scrollback all restored
Built-in browser panel — Ctrl+Shift+L
Navigation bar, DevTools, back/forward
Element Inspector — hover to highlight, click to copy LLM-friendly context
Full CDP automation: click, fill, type, screenshot, JS eval, key press
Output throughput-based activity detection
Taskbar flash + Windows toast notifications
Claude Code, Cursor, Aider, Codex CLI, Gemini CLI, OpenCode, GitHub C

[truncated]

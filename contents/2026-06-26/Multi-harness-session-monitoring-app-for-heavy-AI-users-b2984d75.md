---
source: "https://cctop.app/"
hn_url: "https://news.ycombinator.com/item?id=48686466"
title: "Multi-harness session monitoring app for heavy AI users"
article_title: "cctop - monitor and jump between AI coding sessions"
author: "st0012"
captured_at: "2026-06-26T13:39:10Z"
capture_tool: "hn-digest"
hn_id: 48686466
score: 1
comments: 1
posted_at: "2026-06-26T13:33:11Z"
tags:
  - hacker-news
  - translated
---

# Multi-harness session monitoring app for heavy AI users

- HN: [48686466](https://news.ycombinator.com/item?id=48686466)
- Source: [cctop.app](https://cctop.app/)
- Score: 1
- Comments: 1
- Posted: 2026-06-26T13:33:11Z

## Translation

タイトル: AI ヘビー ユーザー向けのマルチハーネス セッション監視アプリ
記事のタイトル: cctop - AI コーディング セッションを監視し、セッション間を移動する
説明: AI コーディング セッションを監視し、セッション間を移動するキーボードファーストの macOS メニューバー アプリ。 Claude Code、Claude Desktop、opencode、pi、Codex CLI、Codex Desktop で動作します。

記事本文:
CCトップ
ツール
特長
テーマ
インストール
よくある質問
GitHub
ダウンロード
cctop ステータス
AI コーディング セッションに注目してください。
どのセッションがあなたを待っているかを確認してください。キーストロークでそこにジャンプします。
各ツールはセッション イベントをローカルに報告します。 cctop はそれらを 1 つのメニューバー ビューに変換し、適切なエディターまたはターミナル ウィンドウに戻ります。
* kitty.conf でリモート制御が有効になっている場合、Kitty は正確なウィンドウをターゲットにします。これがない場合は、アプリのアクティベーションに戻ります。
† Ghostty 1.3.0+ は、一時的な OSC 7 cwd マーカーを TTY に書き込み、AppleScript 経由でそのマーカーと照合し、実際の cwd を復元することによってセッション端末をターゲットにします。 TTY が使用できない場合は、作業ディレクトリの照合に戻ります。
‡ Apple ターミナルは、AppleScript 経由で tty によってタブをターゲットにします。マルチプレクサー (tmux、screen) 内でのアプリのアクティベーションにフォールバックします。キャプチャされた tty は、[ターミナル] タブではなくマルチプレクサー ペインの pty です。
スキャンしてからジャンプするように設計されています。
アプリはメニューバー内にコンパクトに収まりますが、次にどこに行くかを決定できるように十分なコンテキストが表示されます。
グローバル ホットキーを押して、すべてのセッション カードに番号付きバッジをオーバーレイし、番号を押すと即座にジャンプします。
ヘッダーをドラッグして、パネルの位置を画面上の任意の場所に変更します。位置は起動後も保持され、ダブルクリックするとメニューバーのアンカーに戻ります。
macOS のメニューバーには、「健康」、「注意が必要」、「カメラ ノッチのあるラップトップのスリムな錠剤」の 3 つの状態で表示されます。
2 番目のタブにはセッション履歴が保存されるため、過去のプロジェクトを簡単に再度開くことができます。
ライトとダークの 4 つのパレット。
開発者ツールからインスピレーションを得た配色。設定に切り替えて、ライト、ダーク、またはシステム モードを選択します。
アプリをインストールしてから、ツールを接続します。
最新リリースをダウンロードし、DMG を開き、cctop.app を /Applications にドラッグします。署名されたリリース ビルドは新しい u をチェックできます

自動的に更新されます。
Apple Developer ID で署名され、Apple によって公証されています。 macOS 13 以降で動作します。署名付きリリース ビルドでは、新しい更新を自動的にチェックできます。
[設定] > [ツール] を開きます。 cctop は、検出された各ツール (コピー インストール コマンド、インストール プラグイン、インストール フック、またはトラスト フック) のセットアップ アクションを表示します。
各行に表示されているアクションを使用して、そのツールのプラグインまたはフックをインストールまたは更新します。
Codex CLI/デスクトップの場合、Codex CLI セッションから一度インストールされたフックを信頼します。
新しいセッションまたは再起動されたセッションは、セットアップ後にメニューバー パネルに自動的に表示されます。
パフォーマンス、プライバシー、セットアップ、プラットフォームの詳細についての短い回答。
cctop はコーディング ツールの速度を低下させますか?
いいえ。各統合はセッション イベントに対して最小限の作業を行います。バンドルされたヘルパーを呼び出し、小さな JSON ファイルを書き込み、すぐに戻ります。
cctop はどこかにデータを送信しますか?
分析、テレメトリ、セッションのアップロードはありません。署名付きリリース ビルドは、自動更新チェックとダウンロードにネットワーク アクセスを使用しますが、セッション データはプレーンな JSON として ~/.cctop/sessions/ にマシン上に残ります。
プロジェクトごとに何か設定する必要がありますか?
いいえ。ツールが接続されると、新しいセッションは自動的に追跡されます。プロジェクトごとの設定はありません。
Codex Desktop に追加の信頼手順が必要なのはなぜですか?
Codex は、明示的にレビューして信頼したフックのみを実行します。 cctop はフックをインストールできますが、Codex Desktop は現在フックレビュープロンプトを表示しません。ターミナルで 1 つの Codex CLI セッションを開始し、フックを信頼します。 Codex Desktop はその信頼状態を共有します。
アプリが /Applications に存在する必要があるのはなぜですか?
プラグインは、 /Applications/cctop.app または ~/Applications/cctop.app 内で cctop ヘルパーを探します。他の場所にインストールすると統合が壊れます。
Intel Mac を使用していますが、アップデーターによって間違ったアーキテクチャがインストールされました。
最新リリースから Intel DMG をダウンロードする

e.
/Applications/ 内の既存のアプリを置き換えます。
cctopを再起動します。今後のアップデートでは正しいアーキテクチャが選択される予定です。
/Applications からアプリを削除します。
インストールした各プラグインを削除します。
rm -rf ~/.cctop を使用してセッション データと構成を消去します。

## Original Extract

A keyboard-first macOS menubar app to monitor and jump between AI coding sessions. Works with Claude Code, Claude Desktop, opencode, pi, Codex CLI, and Codex Desktop.

cctop
Tools
Features
Themes
Install
FAQ
GitHub
Download
cctop status
Keep an eye on your AI coding sessions.
See which session is waiting on you. Jump there with a keystroke.
Each tool reports session events locally; cctop turns them into one menubar view and jumps you back to the right editor or terminal window.
* Kitty targets the exact window when remote control is enabled in kitty.conf. Without it, falls back to app activation.
† Ghostty 1.3.0+ targets the session terminal by writing a temporary OSC 7 cwd marker to its TTY, matching that marker via AppleScript, then restoring the real cwd. If the TTY is unavailable, it falls back to working-directory matching.
‡ Apple Terminal targets the tab by tty via AppleScript. Falls back to app activation inside a multiplexer (tmux, screen), where the captured tty is the multiplexer pane's pty rather than the Terminal tab's.
Designed for scanning, then jumping.
The app stays compact in the menubar, but keeps enough context visible that you can decide where to go next.
Hit a global hotkey to overlay numbered badges on every session card, then press the number to jump instantly.
Drag the header to reposition the panel anywhere on screen. Position persists across launches, and double-click snaps it back to the menubar anchor.
Lives in your macOS menubar in three states: healthy, needs attention, and a slim pill for laptops with a camera notch.
A second tab keeps session history so you can reopen past projects easily.
Four palettes, light and dark.
Color schemes inspired by developer tools. Switch in Settings, then choose light, dark, or system mode.
Install the app, then connect your tools.
Download the latest release, open the DMG, and drag cctop.app into /Applications. Signed release builds can check for new updates automatically.
Signed with Apple Developer ID and notarized by Apple. Runs on macOS 13+. Signed release builds can check for new updates automatically.
Open Settings > Tools. cctop shows the setup action for each detected tool: Copy Install Command, Install Plugin, Install Hooks, or Trust Hooks.
Use the action shown on each row to install or update that tool's plugin or hook.
For Codex CLI / Desktop, trust installed hooks once from a Codex CLI session.
New or restarted sessions appear automatically in the menubar panel after setup.
Short answers for performance, privacy, setup, and platform details.
Does cctop slow down my coding tool?
No. Each integration does minimal work on session events: it calls the bundled helper, writes a small JSON file, and returns immediately.
Does cctop send any data anywhere?
No analytics, no telemetry, and no session upload. Signed release builds use network access for automatic update checks and downloads, but session data stays on your machine in ~/.cctop/sessions/ as plain JSON.
Do I need to configure anything per project?
No. Once your tools are connected, new sessions are automatically tracked. There's no per-project setup.
Why does Codex Desktop need an extra trust step?
Codex only runs hooks you've explicitly reviewed and trusted. cctop can install the hooks, but Codex Desktop does not currently surface the hook-review prompt. Start one Codex CLI session in a terminal and trust the hooks; Codex Desktop shares that trust state.
Why does the app need to live in /Applications?
Plugins look for the cctop helper inside /Applications/cctop.app or ~/Applications/cctop.app . Installing elsewhere breaks the integration.
I am on an Intel Mac and the updater installed the wrong architecture.
Download the Intel DMG from the latest release.
Replace the existing app in /Applications/ .
Relaunch cctop. Future updates will pick the correct architecture.
Remove the app from /Applications .
Remove each plugin you installed.
Wipe session data and config with rm -rf ~/.cctop .

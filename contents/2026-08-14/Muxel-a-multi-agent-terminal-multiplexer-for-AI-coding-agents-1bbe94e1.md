---
source: "https://muxel.sh/"
hn_url: "https://news.ycombinator.com/item?id=49301533"
title: "Muxel – a multi-agent terminal multiplexer for AI coding agents"
article_title: "muxel — a multi-agent terminal multiplexer for AI coding agents"
author: "ankitg12"
captured_at: "2026-08-14T17:45:33Z"
capture_tool: "hn-digest"
hn_id: 49301533
score: 3
comments: 0
posted_at: "2026-08-14T17:03:35Z"
tags:
  - hacker-news
  - translated
---

# Muxel – a multi-agent terminal multiplexer for AI coding agents

- HN: [49301533](https://news.ycombinator.com/item?id=49301533)
- Source: [muxel.sh](https://muxel.sh/)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T17:03:35Z

## Translation

タイトル: Muxel – AI コーディング エージェント用のマルチエージェント ターミナル マルチプレクサ
記事のタイトル: muxel — AI コーディング エージェント用のマルチエージェント ターミナル マルチプレクサ
説明: Muxel は、AI コーディング エージェントのチーム全体を 1 つのウィンドウから実行します。分割可能なペインとタブ、ファーストクラスの Git ワークツリー、ライブ ステータス、タスク ランナーとスケジュールされたループ、SSH リモート、組み込みのエディターとファイル ブラウザー、エージェントが必要なときの通知などです。ネイティブ、GPUI 上に構築されています。

記事本文:
最初のペイント前 (フラッシュなし) -->
コンテンツにスキップ
ムクセル
インストールする
特徴
スクリーンショット
ダウンロード
ネイティブ · GPUI 上に構築 · 電子なし
AI コーディング エージェントのチーム全体を 1 つのウィンドウから実行します。
muxel は、エージェント マネージャーのような形をしたターミナル マルチプレクサです (分割可能なペイン)
およびタブ、ファーストクラス git ワークツリー、ライブ ステータス、タスク ランナー、およびスケジュール済み
ループ、SSH リモート、組み込みエディターとファイル ブラウザー、および通知
エージェントがあなたを必要とするとき。
無料 & オープンソース · GPL-3.0 · Linux / macOS / Windows
$ claude "不安定な認証テストを修正"
Routes.rs にパッチを適用しますか? [はい/いいえ]
サイトはプラットフォームを検出し、適切なコマンドを表示します。または、自分でタブを選択します。
最新のビルドを GitHub から直接取得します。
Arch / AppImage — .AppImage 、 chmod +x をダウンロードし、実行します。
Debian / Ubuntu — sudo apt install ./muxel_*.deb
Fedora / RHEL — sudo dnf install ./muxel-*.rpm
ユニバーサル — インテルと Apple シリコン。 muxel.app を ~/Applications にインストールします。
.dmg — 代わりに、muxel を開いてアプリケーションにドラッグします。
署名および公証済み — Gatekeeper の警告なしで開きます。
PowerShell で実行 — ユーザー プロファイルにダウンロードしてインストールします。
クリックスルー設定をご希望ですか? — ダウンロードから署名付きインストーラーを取得します (ユーザーごと、管理者なし、アプリ内自動更新)。
インストール場所 — %LOCALAPPDATA%\muxel\muxel.exe
PATH — 自動的に追加されます。インストール後に新しいターミナルを開きます。
ポータブル — 代わりに、ダウンロードから .zip をダウンロードします。
多数のエージェントを管理するために構築
muxel は、「別の端末アプリ」ではなく、エージェント マネージャーのような形をしています。
任意のペインを水平または垂直に好きな深さに分割します。タブまたはペインをドラッグして、再ドッキング、交換、または最大化します。サイズはプロジェクトごとに維持されます。
┌─────┬─────┐
│ クロード │ オープンコード │
│ §─────

─┤
│ ▓▓▓░ 62% │ シェル ▍ │
━─────┴─────┘
エージェントのライブステータス
動作中、アイドル中、ブロック中、完了 - タブ、サイドバー、ダッシュボードで色分けされます。 Ctrl + Shift + A を使用して、あなたを必要とする次のエージェントにジャンプします。
■クロード・ワーキング
■ オープンコードのアイドル状態
■ コーデックスがブロックされました
■ テストが完了しました
Git ワークツリー
色分けされたワークツリーにより、並列エージェントが分離されます。完了したらコミット、マージ、破棄、または保持します。ワークツリー内のレビューと GitHub PR アクションにより、マージされていないコミットが孤立する前に捕捉されます。
マスター
§─ wt/auth ■ クロード
§─ wt/perf ■ オープンコード
└─ wt/docs ■ マージされました
ペインごとのタブ
各ペインは、ドラッグして並べ替えたり、ペインを越えて移動したり、タブを固定したりできるタブ グループです。右クリックすると、名前変更、複製、および閉じるオプションが表示されます。
Shell、Claude、opencode、Amp、Grok、Hermes、Ollama、Ollama Code、Pi — それぞれ完全に構成可能で、バイナリがインストールされるまで非表示になります。 1 つのプロンプトをすべてのエージェントに一度にブロードキャストします。
すべてのペインは安定したセッション ID を保持するため、クロードのような再開可能なエージェントは、再起動時に新たに開始するのではなく、以前の会話を再開します。
プロジェクトごとのオプトイン: エージェントは、Muxel がソート、タイムスタンプ、プルーニングを維持する共有メモリ ファイルに永続的なレッスンを書き込みます。ドッキングされたパネルから検索、ピン留め、編集します。
すべてのプロジェクトのすべてのエージェントをライブ ステータスとともに 1 つのビューで確認できます。クリックすると、そのビューに直接ジャンプします。
Review や Security Review などのランチャーはエージェントを起動し、プロンプトを入力します。テンプレート化されたプロンプトを使用して独自のプロンプトを定義します。
ツールバー、パレット、またはタブの右クリック メニューから、保存したテキストを実行中のペインに入力します。スニペットごと: 自動送信するか、レビューのためにドロップしてください。
保存されたプロンプトを数分ごと、時間ごと、または毎日の設定時刻に実行します。ループはフォーカスを盗むことなく新しいペインで実行され、その後終了します

エージェントを実行するか、完了したらエージェントを終了します。
SSH 経由でリモート ホスト上でエージェントを実行します。UI、エディタ、ファイル ブラウザはローカルに残ります。 tmux はドロップ後もセッションを維持し、ペインのレイアウトはプロジェクトとともに移動します。
内蔵エディター、gitignore 対応ファイル ブラウザー、マークダウン/イメージ レンダリング。分割または統合された差分を確認します。 git パネルからステージ、破棄、コミットを実行します。
Ctrl+Shift+P は任意のファイルまたはエージェントにジャンプします。フローを離れることなく、プロジェクト ファイルまたはターミナルのスクロールバックを検索します。
トゥルーカラー、スクロールバック検索、クリック可能な URL、OSC-52 クリップボードを備えた完全なアラクリティ ベースの VTE。エージェントは、自分のペインがいつフォーカスされるかを認識します。
プロジェクトには、ライブ エージェントのステータスと現在のブランチが表示されます。右クリックからブランチ、コミット、プル、プッシュ、スタッシュを実行し、各コミットに含まれるファイルを正確に選択します。
集中していないエージェントがあなたを必要としているときにデスクトップ アラートを表示し、すべてを収集するアプリ内フィード、およびエージェントのライブ ステータスを表示するオプションのシステム トレイを提供します。
任意のペインを独自の OS ウィンドウにポップし、後でドッキングし直します。複数のワークスペースは個別のプロジェクト セットを保持します。異なるウィンドウで 2 つを並べて開きます。
レイアウト、分割サイズ、ウィンドウのジオメトリは再起動後も維持されます。中断したところから正確に再開します。
最大 22 のバンドルされたテーマ、再マップ可能なキーバインド、独立したフォント サイズ、および設定からライブで切り替え可能な言語を話す UI。
GPUI 上に構築されています - Electron も Web ビューもありません。 Linux、macOS、Windows (アプリ内アップデートあり)。
ネイティブ デスクトップ アプリ — 左側にプロジェクト、右側にライブ エージェントのグリッド。
ビルドは CI によって生成され、各 GitHub リリースに添付されます。あなたのプラットフォームが強調表示されます。
Intel と Apple Silicon に対応した 1 つのユニバーサル アプリ。 Apple による署名と公証が行われています。
自動更新を備えた署名付きインストーラー (ユーザーごと、管理者なし)。ポータブルな .zip も利用できます。
特定のバージョン、チェックサム、リリース ノートをお探しですか?
兄さん

GitHub 上のすべてのリリースを参照 →
マルチエージェント端末マルチプレクサ。 GPUI 上に構築されています。

## Original Extract

muxel runs a whole team of AI coding agents from one window: splittable panes and tabs, first-class git worktrees, live status, task runners and scheduled loops, SSH remotes, a built-in editor and file browser, and notifications when an agent needs you. Native, built on GPUI.

before first paint (no flash) -->
Skip to content
muxel
install
features
screenshots
download
native · built on GPUI · no electron
Run a whole team of AI coding agents from one window.
muxel is a terminal multiplexer shaped like an agent manager — splittable panes
and tabs, first-class git worktrees, live status, task runners and scheduled
loops, SSH remotes, a built-in editor and file browser, plus notifications
when an agent needs you.
Free & open source · GPL-3.0 · Linux / macOS / Windows
$ claude "fix the flaky auth test"
apply patch to routes.rs? [y/n]
The site detects your platform and shows the right command — or pick a tab yourself.
Grabs the latest build straight from GitHub.
Arch / AppImage — download the .AppImage , chmod +x , run it.
Debian / Ubuntu — sudo apt install ./muxel_*.deb
Fedora / RHEL — sudo dnf install ./muxel-*.rpm
Universal — Intel & Apple Silicon. Installs muxel.app to ~/Applications .
.dmg — open and drag muxel to Applications instead.
Signed & notarized — opens with no Gatekeeper warning.
Run in PowerShell — downloads and installs to your user profile.
Prefer a click-through setup? — grab the signed Installer from Download (per-user, no admin, in-app auto-updates).
Install location — %LOCALAPPDATA%\muxel\muxel.exe
PATH — added automatically; open a new terminal after install.
Portable — download the .zip from Download instead.
Built for managing many agents
muxel is shaped like an agent manager, not “another terminal app.”
Split any pane horizontally or vertically, as deep as you like. Drag tabs or panes to re-dock, swap, or maximize — sizes persist per project.
┌──────────┬──────────┐
│ claude │ opencode │
│ ├──────────┤
│ ▓▓▓░ 62% │ shell ▍ │
└──────────┴──────────┘
Live agent status
Working, idle, blocked or done — color-coded on tabs, sidebar and dashboard. Jump to the next agent that needs you with Ctrl + Shift + A .
■ claude working
■ opencode idle
■ codex blocked
■ tests done
Git worktrees
Color-coded worktrees keep parallel agents isolated. Commit, merge, discard or keep when done — unmerged commits are caught before they're orphaned, with in-worktree review and GitHub PR actions.
master
├─ wt/auth ■ claude
├─ wt/perf ■ opencode
└─ wt/docs ■ merged
Tabs per pane
Each pane is a tab group with drag-to-reorder, cross-pane moves, and pinned tabs. Right-click for rename, duplicate, and close options.
Shell, Claude, opencode, Amp, Grok, Hermes, Ollama, Ollama Code and Pi — each fully configurable, and hidden until its binary is installed. Broadcast one prompt to every agent at once.
Every pane keeps a stable session id, so resume-capable agents like Claude reopen their previous conversation on restart instead of starting fresh.
Opt-in per project: agents write durable lessons to a shared memory file that muxel keeps sorted, timestamped and pruned. Search, pin and edit it from a docked panel.
One view of every agent in every project with live status — click to jump straight to it.
Launchers like Review and Security Review spawn an agent and type the prompt for you. Define your own with templated prompts.
Type saved text into any running pane — from the toolbar, the palette, or a tab's right-click menu. Per snippet: auto-submit, or drop it in for review.
Run a saved prompt every few minutes, hourly, or daily at a set time. Loops fire in a fresh pane without stealing focus, then leave the agent running or exit it when done.
Run agents on remote hosts over SSH — the UI, editor and file browser stay local. tmux keeps sessions alive across drops, and pane layouts roam with the project.
Built-in editor, gitignore-aware file browser, and markdown/image rendering. Review diffs split or unified; stage, discard and commit from the git panel.
Ctrl+Shift+P jumps to any file or agent. Search project files or terminal scrollback without leaving your flow.
Full alacritty-based VTE with truecolor, scrollback search, clickable URLs and OSC-52 clipboard. Agents know when their pane is focused.
Projects show live agent status and current branch. Branch, commit, pull, push and stash from a right-click — and pick exactly which files each commit includes.
Desktop alerts when an unfocused agent needs you, an in-app feed that collects everything, and an optional system tray showing live agent status.
Pop any pane into its own OS window and dock it back later. Multiple workspaces keep separate project sets — open two side by side in different windows.
Layout, split sizes and window geometry persist across restarts. Reopen exactly where you left off.
~22 bundled themes, remappable keybindings, independent font sizes — and a UI that speaks your language, switchable live from settings.
Built on GPUI — no Electron, no web view. Linux, macOS and Windows with in-app updates.
A native desktop app — projects on the left, a grid of live agents on the right.
Builds are produced by CI and attached to each GitHub release. Your platform is highlighted.
One universal app for Intel & Apple Silicon. Signed & notarized by Apple.
Signed installer (per-user, no admin) with auto-updates. Portable .zip also available.
Looking for a specific version, checksums, or release notes?
Browse all releases on GitHub →
Multi-agent terminal multiplexer. Built on top of GPUI.

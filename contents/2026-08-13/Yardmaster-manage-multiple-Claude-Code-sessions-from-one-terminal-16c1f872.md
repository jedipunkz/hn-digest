---
source: "https://github.com/blothecap/yardmaster"
hn_url: "https://news.ycombinator.com/item?id=49284362"
title: "Yardmaster – manage multiple Claude Code sessions from one terminal"
article_title: "GitHub - blothecap/yardmaster · GitHub"
author: "blothecap"
captured_at: "2026-08-13T11:39:17Z"
capture_tool: "hn-digest"
hn_id: 49284362
score: 1
comments: 0
posted_at: "2026-08-13T11:23:40Z"
tags:
  - hacker-news
  - translated
---

# Yardmaster – manage multiple Claude Code sessions from one terminal

- HN: [49284362](https://news.ycombinator.com/item?id=49284362)
- Source: [github.com](https://github.com/blothecap/yardmaster)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T11:23:40Z

## Translation

タイトル: Yardmaster – 1 つの端末から複数のクロード コード セッションを管理
記事のタイトル: GitHub - blothecap/yardmaster · GitHub
説明: GitHub でアカウントを作成して、blothecap/yardmaster の開発に貢献します。

記事本文:
GitHub - ブロスキャップ/ヤードマスター · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ブロスキャップ
/
庭師
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
119 コミット 119 コミット アセット アセット ビルド ビルド ドキュメント ドキュメント src src ウェブサイト ウェブサイト .gitignore .gitignore .nvmrc .nvmrc ライセンス ライセンス通知 通知 README.md README.md electric.vite.config.ts el

ectron.vite.config.ts install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
macOS 用の Claude ネイティブ ターミナル。
すべての Claude 端末を 1 つの端末で管理します。
▶ 25秒の音声付きデモ・
ヤードマスター.me ·
リリース
現在、ほとんどの開発者は複数のクロード コード セッションを同時に実行していますが、端末には何もありません。
アイデア。 20分間働いているクロード、許可を待っているクロード、
デッドペインはすべて同じ長方形のテキストのように見えます。ヤードマスターはmacOSです
クロード セッションがファーストクラスであるターミナル: 名前が付けられ、プロジェクトごとにグループ化され、
信頼できるライブステータス — クロードコード独自のフックプロトコルを話すため
出力から推測するのではなく。
セッション サイドバー — 名前付きセッションがプロジェクトの下にグループ化され、各行に表示されます。
ライブ git ブランチ、アクティビティ行 (正しく実行されている正確なツール) を表示
今 — ▸ Bash: npm test — または待機中の質問)、作業タイマー、および
アプリの再起動後も存続するトークン メーター ( 803k tok 、 1.2M tok )。ステータスドット
クロード コード フックによって駆動され、決して出力スクレイピング: working / need-you /
アイドル/終了しました。
待機中の受信箱 ( ⌘E ) — 正確な質問が含まれるブロックされたすべてのセッション
それは尋ねているのです。そこにジャンプするか、リストから直接承認/拒否します ( a / d )。
ワークツリー セッション — ワンクリックでセッションに独自の分離された git ワークツリーを提供します
( repo/.worktrees/<ブランチ> ); 1 つのリポジトリで複数のクロードを並行して実行します。の
「変更」ペインには、各セッションの差分とワンクリックのマージによるコミットが表示されます
(保護され、衝突安全) またはプッシュ + PR (gh 経由)。
セッションをフォークする — セッションを右クリックして、会話全体を複製します
セッションの HEAD から分岐した新しいワークツリーにコピーされます。リスクのあることを試してみる

アプローチ —
または 2 つの競合するもの — 蓄積されたコンテキストを賭けることなく。
実際の端末、複数形 — ⌘T は、クロードの横にログインシェル タブを開きます
セッション（好きなだけ、⌘⌥←/→で切り替え、⌘Wで閉じる）、
スタンドアロンのターミナル ワークスペースは、プレーンなマルチタブ ターミナルになります - いいえ
クロードが必要です。スクロールバックはタブの切り替え後も存続し、リプレイ バッファを介してリロードされます。
ドラッグ アンド ドロップ — ファイルを Finder から任意のペインとその引用符付きパスにドロップします。
カーソル位置に入力されます。
通知 — バックグラウンド セッション時の macOS 通知 + ドック バッジ
終了したか、入力が必要です。アイドル状態のリマインダーはフィルターで除外されるため、赤は常に赤を意味します。
セッションの永続性 — セッションはアプリの再起動後に存続し、セッションを再開します。
クロード --resume による会話。セッションごとの CLI フラグ (例: --model opus )
彼らと粘り強く付き合ってください。
プロジェクトパネル — アクティブなプロジェクトのダーティファイル数、先行コミット
ベースおよび最後のコミット。サイドバーの下部に常に表示されます。
アプリ内アップデート — アプリは GitHub リリースをチェックします (1 つの匿名 API 呼び出し —
テレメトリはありません）、ワンクリックで自動的に更新されます。または手動で確認してください
メニュー。
ショートカット
アクション
⌘N
新しいセッション (グループヘッダーの + / ⎇ ボタンからプロジェクトを事前入力可能)
⌘1…9
ポジションごとにセッションにジャンプ
⌘J / ⌘K 、 ⌘↓ / ⌘↑ 、 ⌘⇧] / ⌘⇧[
次へ / 前へ (ターミナル ワークスペースを含む)
⌘E
待機中受信ボックス (承認 / 拒否)
⌘T
アクティブなセッション/ワークスペースの新しいターミナル タブ
⌘⌥← / ⌘⌥→
クロードタブとターミナルタブを切り替えます
⌘W
アクティブなターミナル タブを閉じます (クロード セッションは閉じないでください)
⌘R
セッション名の変更
⌘B
サイドバーの切り替え
インストール
要件: Apple シリコン上の macOS、Xcode コマンド ライン ツール。ノードと偶数
クロードコード自体が見つからない場合は自動的にインストールされます。
カール -fsSL https://yardm

aster.me/install.sh |バッシュ
または Homebrew 経由 (事前構築済み、即時 — --no-quarantine はこのゲートキーパーをスキップします)
未署名のビルド):
brew install --cask --no-quarantine blothecap/tap/yardmaster
Curl スクリプトはこのリポジトリを ~/yardmaster に複製し、前提条件、ビルド、および
/Applications にインストールされます (最初のビルドには数分かかります)。その後のアプリは
それ自体を更新するか、いつでもワンライナーを再実行できます。事前に構築された (署名されていない) DMG は次のとおりです。
リリースページで;その後
インストールして、Gatekeeper の隔離をクリアします
xattr -cr /Applications/Yardmaster.app 。
開発の場合: nvm use && npx -y npm@11 install && npm run dev
( npm run dist はアプリバンドルを構築します。アプリは単一インスタンスでロックされるため、
パッケージ化されたインスタンスを起動する前に dev インスタンスを実行します)。
品質ゲート: npm テスト (195 の Vitest テスト — セッション ステート マシン、
git/worktree 操作、および pty フェイクと実際の temp git に対する永続化実行
repos)、npm run typecheck 、および手動パスの docs/smoke-checklist.md 。
電子、3 層。メインプロセスはすべてを所有します: SessionManager
ノードptyプロセス上のステートマシン、およびループバックフックサーバー
挿入されたクロード コード フック (SessionStart / UserPromptSubmit / PreToolUse /
通知/停止) セッションごとの --settings ファイル経由の呼び出し。それが核心だ
トリック: ステータスは推測ではなく既知です — PreToolUse がライブ ツールを強化します
ハートビート、通知により正確な許可の質問が受信箱に送信されます。
Stop は、トランスクリプトからのトークン メータリングをトリガーします。その周りに git/worktree/ を配置します。
モジュールとアトミック session.json の永続性を確認します。レンダラーは、
xterm.js ペイン上のダムな React UI (セッションごとに 1 つ、スイッチ間で維持される)
リプレイバッファを使用)。型付き window.api contextBridge が唯一のシームです
彼らの間で。
マシンからは何も残されません。アカウントもテレメトリもありません。

o クラウド — 唯一の
ネットワーク呼び出しは Claude Code 独自のものであり、匿名の GitHub バージョン チェックによって行われます。
デザイン/仕様の履歴は docs/superpowers/ にあります — アプリは構築されました
クロードと協力して、タスクごとの敵対的レビューを備えた仕様優先。最初の
作業バージョンには 2 日かかりました。
Apache-2.0 — 自由に使用、変更、再配布できます。が含まれています
明示的な特許付与。著作権 2026 Blothecap。
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to blothecap/yardmaster development by creating an account on GitHub.

GitHub - blothecap/yardmaster · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
blothecap
/
yardmaster
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
119 Commits 119 Commits assets assets build build docs docs src src website website .gitignore .gitignore .nvmrc .nvmrc LICENSE LICENSE NOTICE NOTICE README.md README.md electron.vite.config.ts electron.vite.config.ts install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vercel.json vercel.json vitest.config.ts vitest.config.ts View all files Repository files navigation
The Claude-native terminal for macOS.
One terminal to manage all your Claude terminals.
▶ 25-second demo with sound ·
yardmaster.me ·
releases
Most devs now run several Claude Code sessions at once — and their terminal has no
idea. A Claude that's been working for 20 minutes, a Claude waiting for permission,
and a dead pane all look like the same rectangle of text. Yardmaster is a macOS
terminal where Claude sessions are first-class: named, grouped by project, with
live status you can trust — because it speaks Claude Code's own hook protocol
instead of guessing from output.
Sessions sidebar — named sessions grouped under their project, each row
showing the live git branch , an activity line (the exact tool running right
now — ▸ Bash: npm test — or the waiting question), a working timer, and a
token meter ( 803k tok , 1.2M tok ) that survives app restarts. Status dots
are driven by Claude Code hooks , never output-scraping: working / needs-you /
idle / exited.
Waiting-on-you inbox ( ⌘E ) — every blocked session with the exact question
it's asking; jump to it, or approve/deny right from the list ( a / d ).
Worktree sessions — one click gives a session its own isolated git worktree
( repo/.worktrees/<branch> ); run several Claudes on one repo in parallel. The
Changes pane shows each session's diff and commits with one-click merge
(guarded, conflict-safe) or push + PR (via gh ).
Fork Session — right-click a session to duplicate its entire conversation
into a fresh worktree branched from the session's HEAD. Try a risky approach —
or two competing ones — without gambling your accumulated context.
Real terminals, plural — ⌘T opens a login-shell tab beside any Claude
session (as many as you want, ⌘⌥←/→ to switch, ⌘W to close), and a
standalone Terminals workspace makes it a plain multi-tab terminal — no
Claude required. Scrollback survives tab flips and reloads via replay buffers.
Drag & drop — drop files from Finder onto any pane and their quoted paths
are typed at the cursor.
Notifications — macOS notification + dock badge when a background session
finishes or needs input; idle reminders are filtered out so red always means red.
Session persistence — sessions survive app restarts and resume their
conversations via claude --resume ; per-session CLI flags (e.g. --model opus )
persist with them.
Project panel — the active project's dirty-file count, commits ahead of
base, and last commit, always visible at the bottom of the sidebar.
In-app updates — the app checks GitHub releases (one anonymous API call —
there is no telemetry) and updates itself with one click; or check manually via
the menu.
Shortcut
Action
⌘N
New session (project pre-fillable from a group header's + / ⎇ buttons)
⌘1…9
Jump to session by position
⌘J / ⌘K , ⌘↓ / ⌘↑ , ⌘⇧] / ⌘⇧[
Next / previous (includes the Terminals workspace)
⌘E
Waiting-on-you inbox ( a approve / d deny)
⌘T
New terminal tab in the active session / workspace
⌘⌥← / ⌘⌥→
Switch between the Claude tab and terminal tabs
⌘W
Close the active terminal tab (never the Claude session)
⌘R
Rename session
⌘B
Toggle sidebar
Install
Requirements: macOS on Apple silicon, Xcode Command Line Tools. Node and even
Claude Code itself are installed automatically if missing.
curl -fsSL https://yardmaster.me/install.sh | bash
or via Homebrew (prebuilt, instant — --no-quarantine skips Gatekeeper on this
unsigned build):
brew install --cask --no-quarantine blothecap/tap/yardmaster
The curl script clones this repo to ~/yardmaster , checks prerequisites, builds, and
installs to /Applications (first build takes a few minutes). After that the app
updates itself — or re-run the one-liner any time. A prebuilt (unsigned) DMG is
on the releases page ; after
installing it, clear Gatekeeper's quarantine with
xattr -cr /Applications/Yardmaster.app .
For development: nvm use && npx -y npm@11 install && npm run dev
( npm run dist builds the app bundle; the app single-instance-locks, so close a
dev instance before launching the packaged one).
Quality gates: npm test (195 Vitest tests — the session state machine,
git/worktree operations, and persistence run against pty fakes and real temp git
repos), npm run typecheck , and docs/smoke-checklist.md for the manual pass.
Electron, three layers. The main process owns everything: a SessionManager
state machine over node-pty processes, and a loopback hook server that
injected Claude Code hooks ( SessionStart / UserPromptSubmit / PreToolUse /
Notification / Stop ) call via per-session --settings files. That's the core
trick: status is known , not guessed — PreToolUse powers the live tool
heartbeat, Notification carries the exact permission question into the inbox,
Stop triggers token metering from the transcript. Around that sit git/worktree/
review modules and atomic sessions.json persistence. The renderer is a
dumb React UI over xterm.js panes (one per session, kept alive across switches
with replay buffers). A typed window.api contextBridge is the only seam
between them.
Nothing leaves your machine: no accounts, no telemetry, no cloud — the only
network calls are Claude Code's own and an anonymous GitHub version check.
The design/spec history lives in docs/superpowers/ — the app was built
spec-first with per-task adversarial review, working with Claude. The first
working version took two days.
Apache-2.0 — free to use, modify, and redistribute; includes an
explicit patent grant. Copyright 2026 Blothecap.
Readme Apache-2.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

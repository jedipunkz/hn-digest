---
source: "https://github.com/xreader/ai-session-manager"
hn_url: "https://news.ycombinator.com/item?id=48561870"
title: "Claude Session Manager – find and manage forgotten Claude Code sessions"
article_title: "GitHub - xreader/ai-session-manager: A local, zero‑dependency web dashboard to find, overview and manage all your Claude Code sessions across every folder on your machine — including the ones you started and forgot. · GitHub"
author: "xreader"
captured_at: "2026-06-16T21:53:38Z"
capture_tool: "hn-digest"
hn_id: 48561870
score: 2
comments: 0
posted_at: "2026-06-16T20:49:51Z"
tags:
  - hacker-news
  - translated
---

# Claude Session Manager – find and manage forgotten Claude Code sessions

- HN: [48561870](https://news.ycombinator.com/item?id=48561870)
- Source: [github.com](https://github.com/xreader/ai-session-manager)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T20:49:51Z

## Translation

タイトル: クロード セッション マネージャー – 忘れられたクロード コード セッションを検索して管理します
記事のタイトル: GitHub - xreader/ai-session-manager: マシン上のすべてのフォルダーにわたるすべてのクロード コード セッション (開始したものや忘れたものを含む) を検索、概要、管理するローカルの依存関係のない Web ダッシュボード。 · GitHub
説明: ローカルの依存関係のない Web ダッシュボード。マシン上のすべてのフォルダーにわたるすべてのクロード コード セッション (開始したものや忘れたものを含む) を検索、概要、管理します。 - xreader/ai-session-manager

記事本文:
GitHub - xreader/ai-session-manager: マシン上のすべてのフォルダーにわたるすべてのクロード コード セッション (開始したものや忘れたものを含む) を検索、概要、管理するローカルの依存関係のない Web ダッシュボードです。 · GitHub
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
エックスリーダー
/
ある

i-セッションマネージャー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット docs docs lang lang .gitignore .gitignore ライセンス ライセンス README.md README.md Index.html Index.html server.js server.js start.bat start.bat stop.bat stop.bat すべてのファイルを表示 リポジトリ ファイルのナビゲーション
依存関係のないローカルの Web ダッシュボードにより、マシン上のすべてのフォルダーにわたるすべてのクロード コード セッション (開始したものや忘れたものを含む) を検索、概要、管理できます。
Claude コードは、すべてのセッションを ~/.claude の下に集中的に保存します。多くの作業に取り組むと、
プロジェクトを始めると、何を始めたか、どこで、そしてそれがまだ実行されているかどうかを思い出すのが難しくなります。
このツールはそのデータ (読み取り専用) を読み取り、1 つの明確な概要をリストとして提供します。
タイムライン、またはカンバン ボード。
⚠️非公式コミュニティプロジェクト。 Anthropic と提携または承認されていません。
リスト + 詳細 — すべてのセッションがプロジェクトごとにグループ化され、豊富な詳細パネルが表示されます。
タイムライン — 各セッションがいつ実行されたかをアクティビティ マーカーとともに示す絶対時間軸:
ボード — 独自のかんばん列全体にセッションをドラッグします。
あるプロジェクトで Claude Code セッションを開始し、中断されて、まだ実行中であることを忘れてしまいます。
セッションがどのフォルダーに入っていたのか、そのセッションの内容を思い出せません。
これまで取り組んできたことすべてと現在の状態を俯瞰したいと考えています。
Claude Session Manager はまさにそれを解決します。
🟢 ライブ ステータス — 実際のプロセス チェックを通じて、どのセッションが実際に実行中 (🟢)、アイドル状態 (🟡)、または閉じられているのか (⚫) を確認します。
☰ リスト + 詳細 — プロジェクトごとにグループ化されたセッション。クリックすると詳細パネルが表示されます:
🎯 目標 (最初のプロンプト) と 📍 最新の状態 (最後のプロンプト) — 瞬時の合計

AI 呼び出しは必要ありません。
💬完全なプロンプト履歴、🖼️画像ギャラリー（セッションに貼り付けられたスクリーンショット）、メッセージ数、期間、ブランチ、PID、フォルダー。
⧗ タイムライン (ガント) — 各セッションがいつ実行されたかを示す単一の絶対時間軸 (年→月→日) とアクティビティのプロンプト マーカー。
▦ ボード (カンバン) — セッションを独自の列にドラッグして、ワークフローを整理します。
🎴 ビジュアル プロジェクト フィルター — 代表的な画像 (または色) を使用してタイルをプロジェクトします。
▶ 新しい端末で忘れてしまったセッションを再開し、 📂 フォルダーを開き、 📝 メモを追加します。
🌗 ライトとダークのテーマ、🌍 多言語 (英語/ドイツ語を含む)、すべて設定可能。
🔌 Config-first — ほとんどすべてが設定スキーマ (プラグイン システムへの足がかり) によって駆動されます。
ローカルの ~/.claude ディレクトリから読み取り専用で読み取ります。
独自のデータ (かんばん列、メモ) は別のサイドカー ファイルに保存されます
( ~/.claude/session-manager-state.json )。元のクロード コード ファイルは決して変更されません。
100%地元産。どこにもアップロードされるものはありません。これは、 localhost 上のプレーンな Node HTTP サーバーです。
Web UI は、お使いのマシン (デフォルト http://localhost:4317 ) にのみ提供されます。
セッションへの読み取り専用アクセス。書き込みは独自のサイドカー/構成ファイルにのみ行われます。
Node.js 18+ (22 でテスト済み)。 npm インストールも依存関係もありません。
git clone https://github.com/ < あなた > /claude-session-manager.git
cd クロード セッション マネージャー
ノードサーバー.js
次に、 http://localhost:4317 を開きます。
Windows: start.bat をダブルクリックします (サーバーが起動し、ブラウザが開きます)。
停止するには stop.bat を使用します。
カスタムポート: PORT=5000 ノードserver.js (または設定で設定)。
右上の⚙を開きます。すべては ~/.claude/session-manager-config.json に保存されます。
言語、テーマ (ライト/ダーク)、ビューの開始、自動更新間隔、サーバー

ポート
カンバン列 (自由に定義可能)
「再開」に使用されるターミナル — ライブ プレビュー付きのプラットフォームごとのプリセット:
Windows: PowerShell、cmd、Windows ターミナル
Linux: gnome‑terminal、konsole、xterm
…または、{cwd} および {id} プレースホルダーを使用したカスタム コマンド。
lang/en.json をコピーし、値を翻訳し、 "_name" (例: "Français" ) を設定し、名前を付けて保存します。
lang/fr.json — 言語ピッカーに自動的に表示されます。コードの変更はありません。 PRの方も大歓迎！
🧠 Claude Code だけでなく、より多くの AI アシスタントをサポート - プラグイン可能な「セッション プロバイダー」
他のコーディング エージェント/CLI (ローカル セッション/トランスクリプト データを保持するツールなど)。
🔍 すべてのトランスクリプトにわたる全文検索。
🔔 忘れられた/アイドル状態の実行セッションの通知。
🌿 セッションフォルダーごとのオプションの git ステータス (コミットされていない変更、前方/後方)。
📈 タイムライン内の単一のドットの代わりにアクティビティ パルス (GitHub スタイル)。
🧩 プラグイン システム — ドロップイン モジュールを介してビュー、列、コマンド、設定を提供します。
アイデアやフィードバックは歓迎です - 問題を開いてください!
小規模で焦点を絞った PR がレビューしやすいです。
新しい翻訳は特に歓迎されます (「言語の追加」を参照)。
より大きな変更の場合は、まず問題を開いて方向性について話し合います。
アプリ全体は意図的に小さく、依存関係がありません。単一の server.js 、単一の
Index.html 、および lang/*.json 。読みやすく、ハッキングも簡単です。
MIT — 好きなようにしてください。保証はありません。
依存関係のないローカルの Web ダッシュボードにより、マシン上のすべてのフォルダーにわたるすべてのクロード コード セッション (開始したものや忘れたものを含む) を検索、概要、管理できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHu

株式会社ビー
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A local, zero‑dependency web dashboard to find, overview and manage all your Claude Code sessions across every folder on your machine — including the ones you started and forgot. - xreader/ai-session-manager

GitHub - xreader/ai-session-manager: A local, zero‑dependency web dashboard to find, overview and manage all your Claude Code sessions across every folder on your machine — including the ones you started and forgot. · GitHub
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
xreader
/
ai-session-manager
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits docs docs lang lang .gitignore .gitignore LICENSE LICENSE README.md README.md index.html index.html server.js server.js start.bat start.bat stop.bat stop.bat View all files Repository files navigation
A local, zero‑dependency web dashboard to find, overview and manage all your Claude Code sessions across every folder on your machine — including the ones you started and forgot.
Claude Code stores every session centrally under ~/.claude . Once you work across many
projects, it gets hard to remember what you started, where, and whether it's still running.
This tool reads that data (read‑only) and gives you one clean overview — as a list , a
timeline , or a kanban board .
⚠️ Unofficial community project. Not affiliated with or endorsed by Anthropic.
List + detail — all sessions grouped by project, with a rich detail panel:
Timeline — an absolute time axis showing when each session ran, with activity markers:
Board — drag sessions across your own kanban columns:
You start a Claude Code session in some project, get interrupted, and forget it's still running .
You can't remember which folder a session was in, or what it was about .
You want a bird's‑eye view of everything you've worked on and its current state.
Claude Session Manager solves exactly that.
🟢 Live status — see which sessions are actually running (🟢), idle (🟡) or closed (⚫), via a real process check.
☰ List + detail — sessions grouped by project; click for a rich detail panel:
🎯 Goal (first prompt) and 📍 latest state (last prompt) — an instant summary, no AI call needed.
💬 full prompt history , 🖼️ image gallery (screenshots pasted into the session), message counts, duration, branch, PID, folder.
⧗ Timeline (Gantt) — a single absolute time axis (years → months → days) showing when each session ran, with prompt markers for activity.
▦ Board (Kanban) — drag sessions across your own columns to organize your workflow.
🎴 Visual project filter — project tiles with a representative image (or color).
▶ Resume a forgotten session in a new terminal, 📂 open its folder , 📝 add notes .
🌗 Light & dark theme , 🌍 multi‑language (English/German included), all configurable.
🔌 Config‑first — almost everything is driven by a settings schema (a stepping stone toward a plugin system).
It reads, read‑only , from your local ~/.claude directory:
Your own data (kanban columns, notes) is stored in a separate sidecar file
( ~/.claude/session-manager-state.json ). The original Claude Code files are never modified.
100% local . Nothing is uploaded anywhere — it's a plain Node HTTP server on localhost .
The web UI is only served to your machine (default http://localhost:4317 ).
Read‑only access to your sessions; writes go only to its own sidecar/config files.
Node.js 18+ (tested on 22). No npm install, no dependencies .
git clone https://github.com/ < you > /claude-session-manager.git
cd claude-session-manager
node server.js
Then open http://localhost:4317 .
Windows: double‑click start.bat (starts the server and opens your browser).
Use stop.bat to stop it.
Custom port: PORT=5000 node server.js (or set it in the settings).
Open ⚙ in the top‑right. Everything is stored in ~/.claude/session-manager-config.json .
Language , theme (light/dark), start view , auto‑refresh interval , server port
Kanban columns (freely definable)
Terminal used for “Resume” — presets per platform with a live preview:
Windows: PowerShell, cmd, Windows Terminal
Linux: gnome‑terminal, konsole, xterm
…or a custom command with {cwd} and {id} placeholders.
Copy lang/en.json , translate the values, set "_name" (e.g. "Français" ), save as
lang/fr.json — it appears in the language picker automatically. No code changes. PRs welcome!
🧠 Support more AI assistants , not just Claude Code — pluggable “session providers” for
other coding agents/CLIs (e.g. tools that keep local session/transcript data).
🔍 Full‑text search across all transcripts.
🔔 Notifications for forgotten/idle running sessions.
🌿 Optional git status per session folder (uncommitted changes, ahead/behind).
📈 Activity pulse (GitHub‑style) instead of single dots in the timeline.
🧩 Plugin system — contribute views, columns, commands and settings via drop‑in modules.
Ideas and feedback welcome — open an issue!
Small, focused PRs are easiest to review.
New translations are especially appreciated (see “Adding a language”).
For larger changes, open an issue first to discuss the direction.
The whole app is intentionally tiny and dependency‑free: a single server.js , a single
index.html , and lang/*.json . Easy to read, easy to hack.
MIT — do whatever you like, no warranty.
A local, zero‑dependency web dashboard to find, overview and manage all your Claude Code sessions across every folder on your machine — including the ones you started and forgot.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

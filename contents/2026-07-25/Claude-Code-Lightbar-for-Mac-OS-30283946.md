---
source: "https://github.com/gregsramblings/claude-status-bar"
hn_url: "https://news.ycombinator.com/item?id=49048037"
title: "Claude Code Lightbar for Mac OS"
article_title: "GitHub - gregsramblings/claude-status-bar · GitHub"
author: "gw5815"
captured_at: "2026-07-25T14:54:53Z"
capture_tool: "hn-digest"
hn_id: 49048037
score: 1
comments: 1
posted_at: "2026-07-25T14:50:12Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Lightbar for Mac OS

- HN: [49048037](https://news.ycombinator.com/item?id=49048037)
- Source: [github.com](https://github.com/gregsramblings/claude-status-bar)
- Score: 1
- Comments: 1
- Posted: 2026-07-25T14:50:12Z

## Translation

タイトル: Mac OS 用クロード コード ライトバー
記事のタイトル: GitHub - gregsramblings/claude-status-bar · GitHub
説明: GitHub でアカウントを作成して、gregsramblings/claude-status-bar の開発に貢献します。

記事本文:
GitHub - gregsramblings/claude-status-bar · GitHub
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
グレッグスランブリング
/
クロードステータスバー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
グレッグサー

amblings/claude-status-bar
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット 画像 画像 .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md ccbar.swift ccbar.swift フック.sh フック.sh install.sh install.sh uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
バージョン 1.0 · macOS のみ (AppKit + launchd を使用。Windows または Linux ビルドはありません)。
すべての画面の一端に沿って表示される全幅の赤いバーは、次のことを示します。
部屋の向こうで、クロード・コードが何をしているのか。デフォルトでは上端に位置します
(位置、厚さ、色はすべて変更可能です。構成を参照してください)。
赤色の点灯バー — 少なくとも 1 つのセッションがビジーです (クロードが動作中)。
赤いバーが点滅 — セッションは今あなたからの応答を必要としています (許可)
プロンプト、またはクロードが入力の待機をブロックしました）。
バーなし - ターンが単に終了したか、何も実行されていません。
ビジーの勝ち: いずれかのセッションがまだ動作している場合、バーはしっかりと表示されます。 1つのときのみパルスします
アクティブなプロンプト (許可リクエストまたは明示的な質問) があります。ターン
それだけでバーがクリアされます。バーはクリックスルーです (決してインターセプトされません)
マウス)、すべてのスペース上で全画面アプリを含むすべての上にフロートします
そしてすべてのモニター。
注: パルスは、Claude Code の通知フックによって駆動されます。
実際のプロンプト (許可リクエストなど)。 ~60 秒のアイドル通知は、
意図的に無視されています (hook.sh は「入力を待っています」メッセージをフィルターします)。
終了したセッションがそこにあるからといって、バーが点滅することはありません。 1 つ
結果: クロード コードには、「ターンを終了した」という明確なイベントはありません。
質問」なので、ターン終了時の単純な会話の質問は完了したと見なされます。
(バーなし) — 実際のプロンプトのみが脈動します。構造化されたAskUserQuestionプロンプトは次のことを行います
脈拍、

そのツールの PreToolUse / PostToolUse フック経由。
左: 下端に沿って赤色で点灯 = クロードが動作しています。クロードがあなたからの応答を必要としているときに脈動します。右: ライブ設定用のメニューバー メニュー。
macOS ( AppKit + launchd を使用)。
Xcode コマンド ライン ツール — ビルドに swiftc を提供します。以下を使用してインストールします。
xcode-select --install
Claude Code CLI — これは Claude Code によって駆動されます
フック。ローカルで実行されているクロード コード (ターミナルまたは IDE) で動作します。
拡張子）。クラウド セッション (claude.ai/code) では機能しません。これらのセッションは実行されます。
リモートにあるため、ローカル マシンにアクセスできません。
その他の依存関係はありません。Homebrew、jq、ランタイムはありません。まさにシステムSwift
コンパイラと plutil (macOS に組み込まれています)。
git clone https://github.com/gregsramblings/claude-status-bar.git
cd クロードステータスバー
./install.sh
install.sh はアプリをビルドし、ログイン時実行 LaunchAgent をインストールします (パスは
クローンを作成した場所 (ハードコーディングされていないもの) を開始し、フック ブロックを出力します。
そのブロックを ~/.claude/settings.json にコピーします (最上位の JSON にマージします)
オブジェクト;すでに「フック」キーがある場合は、それに 4 つのイベントを追加します)。ブロック
実際のパスが入力されていると、次のようになります。
「フック」: {
"UserPromptSubmit" : [{ "フック" : [{ "タイプ" : " コマンド " , "コマンド" : " bash /path/to/claude-status-bar/hook.sh ビジー " }] }],
"停止" : [{ "フック" : [{ "タイプ" : " コマンド " , "コマンド" : " bash /path/to/claude-status-bar/hook.sh 待機 " }] }],
"通知" : [{ "フック" : [{ "タイプ" : " コマンド " , "コマンド" : " bash /path/to/claude-status-bar/hook.sh need_input " }] }],
"SessionEnd" : [{ "hooks" : [{ "type" : " command " , "command" : " bash /path/to/claude-status-bar/hook.sh end " }] }]
}
次に、新しいクロード コード セッションを開始し、何かを実行します (赤いバー)。
動作中に表示されます。 (フックはセッション時にロードされます)

が開始されるため、すでに開いているセッションは開始されません
再始動するまでバーを駆動してください。)
エコー ' {"セッション ID":"テスト"} ' | ./hook.sh ビジー # バーが表示されるはずです
エコー ' {"セッション ID":"テスト"} ' | ./hook.sh end # バーは消えるはずです
アンインストール
./アンインストール.sh
LaunchAgent と状態ディレクトリを停止して削除します。意図的にはそうではありません
settings.json を編集します — 4 つの ccbar フック エントリを自分で削除してから、削除します
リポジトリフォルダー。
クロード コードのフックは状態の変更時に起動され、hook.sh を実行します。
セッションごとの小さな JSON ファイルを ~/.claude/ccbar/state/<session_id>.json に保存します。
UserPromptSubmit → ビジー (あなたは引き継ぎました。クロードは仕事中です)
停止→待機（ターン終了）
通知 → 通知、hook.sh は need_input (実際の
プロンプト→パルス）または待機中（アイドルタイムアウト→バーなし）
PreToolUse (AskUserQuestion に一致) → need_input (パルス中にパルス
多肢選択式の質問があなたを待っています); PostToolUse → ビジー（再び点灯）
答えたら）
ccbar (小さな AppKit アプリ、画面ごとに 1 つの境界のないクリックスルー ウィンドウ
.screenSaver レベル、すべてのスペース上) 0.4 秒ごとに状態ディレクトリをポーリングします: 実線
いずれかのライブ セッションがビジーの場合、それ以外の場合はパルス (needs_input の場合)、それ以外の場合
非表示 (待機中 - 終了したターン - 何も表示されません)。
マルチセッションは正常に機能します。各セッションは session_id をキーとする独自のファイルであるため、
3 つの同時セッションが 1 つのバーを共有します。動作している場合は点灯、パルスがある場合は点滅します。
あなたの意見が必要です。古いファイル (12 時間以上古い、スキップされたクラッシュによるものなど)
SessionEnd ) は無視されます。
ccbar は、macOS メニュー バーに小さなバー状のアイコンを追加します。クリックして設定します
バーはライブです — リビルドも再起動もありません:
選択内容は再起動後も UserDefaults に保持されます。残りの調整可能パラメータはそのままです
ccbar.swift の先頭にあるコンパイル時定数 (変更後のリビルド + 再起動 —
。

/install.sh は再び両方を実行します):
ファイル
目的
ccbar.swift
アプリ全体 (AppKit、単一ファイル)。
フック.sh
クロードコードフックターゲット。セッションごとに 1 つの状態ファイルを書き込みます。
インストール.sh
LaunchAgent をビルドしてインストールし、フック ブロックを出力します。
アンインストール.sh
LaunchAgent と状態ディレクトリを停止して削除します。
CCバー
コンパイル済みバイナリ (gitignored — install.sh によってビルド)。
注意事項
脆弱性 — アイドル フィルターは文字列一致です。フック.sh は、
部分文字列の一致による実際のプロンプトからの約 60 秒間のアイドル通知
通知メッセージへの入力を待っています。これはクロード・コードのものです
現在アイドル状態の文言であり、安定した API ではありません。 Anthropic がそのテキストを変更すると、アイドルは
タイムアウトになると、バーのパルスが再び開始されます。修正は 1 行です — ケースを更新します
hook.sh のパターン。 (逆に、既知のアイドル テキストのみが抑制されるため、新しいテキストはすべて抑制されます)
メッセージのデフォルトはpulseであり、実際のプロンプトを黙って非表示にすることはありません。）
Stop はメインエージェントの停止のみです (SubagentStop はフックされていない別のイベントです)。
そのため、バーはサブエージェントのチャーンではなく、実際のターン境界を追跡します。
バーはデフォルトで上端に設定されます。スイッチ位置→下（またはその逆）、
好みに合わせてメニューバーメニューの厚みをバンプします — 例:エッジから外して移動します。
下部のドックまたはノッチがそれを隠します。
パブリック ドメイン — Unlicense 。それを使って何でも好きなようにしてください、いいえ
帰属は必要ありませんが、保証も責任もありません。
Readme ライセンス解除アクティビティのスター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to gregsramblings/claude-status-bar development by creating an account on GitHub.

GitHub - gregsramblings/claude-status-bar · GitHub
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
gregsramblings
/
claude-status-bar
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
gregsramblings/claude-status-bar
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits images images .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md ccbar.swift ccbar.swift hook.sh hook.sh install.sh install.sh uninstall.sh uninstall.sh View all files Repository files navigation
Version 1.0 · macOS only (uses AppKit + launchd ; there is no Windows or Linux build).
A full-width red bar along one edge of every screen that tells you, from
across the room, what Claude Code is doing. It sits on the top edge by default
(position, thickness, and color are all changeable — see Configuration ).
Solid red bar — at least one session is busy (Claude working).
Pulsing red bar — a session needs a response from you now (a permission
prompt, or Claude blocked waiting for input).
No bar — a turn simply finished, or nothing running.
Busy wins: if any session is still working the bar is solid; it pulses only when one
has an active prompt for you (a permission request or an explicit question). A turn
that just finishes clears the bar. The bar is click-through (it never intercepts
your mouse) and floats above everything, including fullscreen apps, on every Space
and every monitor.
Note: pulsing is driven by Claude Code's Notification hook, which fires for a
real prompt (a permission request, etc.). The ~60s idle notification is
deliberately ignored ( hook.sh filters the "waiting for your input" message), so
the bar never blinks just because a finished session is sitting there. One
consequence: Claude Code has no distinct event for "ended a turn with a
question," so a plain conversational question at the end of a turn reads as done
(no bar) — only actual prompts pulse. Structured AskUserQuestion prompts do
pulse, via a PreToolUse / PostToolUse hook on that tool.
Left: solid red along the bottom edge = Claude is working; it pulses when Claude needs a response from you. Right: the menu-bar menu for live configuration.
macOS (uses AppKit + launchd ).
Xcode Command Line Tools — provides swiftc to build. Install with:
xcode-select --install
Claude Code CLI — this is driven by Claude Code
hooks. It works with any locally-running Claude Code (terminal or IDE
extension). It does not work with cloud sessions (claude.ai/code) — those run
remotely and can't reach your local machine.
No other dependencies: no Homebrew, no jq , no runtime. Just the system Swift
compiler and plutil (built into macOS).
git clone https://github.com/gregsramblings/claude-status-bar.git
cd claude-status-bar
./install.sh
install.sh builds the app, installs a run-at-login LaunchAgent (paths derived from
wherever you cloned — nothing hardcoded), starts it, then prints a hooks block .
Copy that block into ~/.claude/settings.json (merge it into the top-level JSON
object; if you already have a "hooks" key, add the four events to it). The block
looks like this, with real paths filled in:
"hooks" : {
"UserPromptSubmit" : [{ "hooks" : [{ "type" : " command " , "command" : " bash /path/to/claude-status-bar/hook.sh busy " }] }],
"Stop" : [{ "hooks" : [{ "type" : " command " , "command" : " bash /path/to/claude-status-bar/hook.sh waiting " }] }],
"Notification" : [{ "hooks" : [{ "type" : " command " , "command" : " bash /path/to/claude-status-bar/hook.sh needs_input " }] }],
"SessionEnd" : [{ "hooks" : [{ "type" : " command " , "command" : " bash /path/to/claude-status-bar/hook.sh end " }] }]
}
Then start a new Claude Code session and give it something to do — the red bar
appears while it works. (Hooks load at session start, so already-open sessions won't
drive the bar until restarted.)
echo ' {"session_id":"test"} ' | ./hook.sh busy # bar should appear
echo ' {"session_id":"test"} ' | ./hook.sh end # bar should disappear
Uninstall
./uninstall.sh
Stops and removes the LaunchAgent and the state directory. It intentionally does not
edit settings.json — remove the four ccbar hooks entries yourself, then delete
the repo folder.
Claude Code hooks fire on state change and run hook.sh , which writes one
small JSON file per session into ~/.claude/ccbar/state/<session_id>.json :
UserPromptSubmit → busy (you handed off; Claude is working)
Stop → waiting (turn finished)
Notification → notify , which hook.sh splits into needs_input (a real
prompt → pulse) or waiting (the idle timeout → no bar)
PreToolUse (matching AskUserQuestion ) → needs_input (pulse while a
multiple-choice question waits on you); PostToolUse → busy (solid again
once you answer)
ccbar (a tiny AppKit app, one borderless click-through window per screen at
.screenSaver level, on all Spaces) polls the state directory every 0.4s: solid
if any live session is busy , else pulsing if any is needs_input , else
hidden ( waiting — a finished turn — shows nothing).
Multi-session just works: each session is its own file keyed by session_id , so
three concurrent sessions share one bar — solid if any is working, pulsing if any
needs your input. Stale files (older than 12h, e.g. from a crash that skipped
SessionEnd ) are ignored.
ccbar adds a small bar-shaped icon to the macOS menu bar . Click it to configure
the bar live — no rebuild, no restart:
Choices persist in UserDefaults across restarts. The remaining tunables are still
compile-time constants at the top of ccbar.swift (rebuild + restart after changing —
./install.sh again does both):
File
Purpose
ccbar.swift
The whole app (AppKit, single file).
hook.sh
Claude Code hook target; writes one state file per session.
install.sh
Build + install LaunchAgent + print the hooks block.
uninstall.sh
Stop + remove the LaunchAgent and state dir.
ccbar
Compiled binary (gitignored — built by install.sh ).
Notes
Fragility — the idle filter is a string match. hook.sh distinguishes the
~60s idle notification from a real prompt by matching the substring
waiting for your input in the notification message . This is Claude Code's
current idle wording, not a stable API. If Anthropic changes that text, the idle
timeout will start pulsing the bar again. The fix is one line — update the case
pattern in hook.sh . (Conversely, only known idle text is suppressed, so any new
message defaults to pulse , never to silently hiding a real prompt.)
Stop is the main-agent stop only ( SubagentStop is a separate, unhooked event),
so the bar tracks real turn boundaries, not sub-agent churn.
The bar defaults to the top edge. Switch Position → Bottom (or the reverse) and
bump Thickness in the menu-bar menu to taste — e.g. move it off an edge where a
bottom Dock or the notch would hide it.
Public domain — The Unlicense . Do whatever you want with it, no
attribution required, no warranty, no liability.
Readme Unlicense Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

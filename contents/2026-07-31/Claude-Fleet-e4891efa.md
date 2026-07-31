---
source: "https://github.com/Shubhamgupta2612/claude-fleet"
hn_url: "https://news.ycombinator.com/item?id=49129782"
title: "Claude-Fleet"
article_title: "GitHub - Shubhamgupta2612/claude-fleet: macOS menu bar app that shows your Claude Code sessions (local + remote over SSH), grouped by repo, so you can see which one is waiting on you · GitHub"
author: "Shubham2612"
captured_at: "2026-07-31T23:53:58Z"
capture_tool: "hn-digest"
hn_id: 49129782
score: 1
comments: 0
posted_at: "2026-07-31T23:48:21Z"
tags:
  - hacker-news
  - translated
---

# Claude-Fleet

- HN: [49129782](https://news.ycombinator.com/item?id=49129782)
- Source: [github.com](https://github.com/Shubhamgupta2612/claude-fleet)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T23:48:21Z

## Translation

タイトル: クロード・フリート
記事のタイトル: GitHub - Shubhamgupta2612/claude-fleet: クロード コード セッション (ローカル + SSH 経由のリモート) をリポジトリごとにグループ化して表示する macOS メニュー バー アプリなので、どれが待機しているかを確認できます · GitHub
説明: クロード コード セッション (ローカル + SSH 経由のリモート) をリポジトリごとにグループ化して表示する macOS メニュー バー アプリ。どのセッションが待機しているかを確認できます - Shubhamgupta2612/claude-fleet

記事本文:
GitHub - Shubhamgupta2612/claude-fleet: クロード コード セッション (ローカル + SSH 経由のリモート) をリポジトリごとにグループ化して表示する macOS メニュー バー アプリ。どのセッションが待機しているかを確認できます · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
シュバムグプタ26

12
/
クロード・フリート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE .superpowers .superpowers ソース ソース docs docs fixtures fixtures swift-tests swift-tests テスト テスト ツール ツール .gitignore .gitignore AppIcon.icns AppIcon.icns CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SUPPORT.md SUPPORT.md build.sh build.sh Collector.py Collector.py fixture.json fixture.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
持っているすべてのクロード コード セッションを表示するネイティブ macOS メニュー バー アプリ
この Mac 上と SSH 経由のリモート マシン上で実行中です。git リポジトリごとにグループ化されているため、
どれがあなたを待っているのか一目でわかります。
それがあなたに伝えるために存在する唯一のことは、どのセッションが未回答のままであるかということです。
現在許可プロンプトが表示されます。クロードが「このコマンドを実行してもいいですか?」と尋ねると、の中に
ウィンドウが見ていない場合、そのセッションはフリーズし、確認するまで何もわかりません。
たまたまそれに切り替わります。クロード・フリートはそれをメニューバーに置きます - そしてセッションのときに
あなたを必要とし始めると、それはトップにジャンプします。
コミュニティによって構築された独立したツール。 Anthropic と提携または承認されていません。
サポート姿勢 — 最初にお読みください
これは個人用ツールであり、1 つのセットアップ用に構築され、MIT のもとで現状のまま公開されています。それ
Claude Code のディスク上のファイルを読み取ります。これらのファイルは文書化されておらず、変更される可能性があります。
クロード コードをリリースすると、誰かが修正するまで壊れてしまいます。
壊れたものは私が直します。問題は未解決ですが、回答が得られない可能性があります。 PRは
歓迎ですが、座っても構いません。ロードマップもリリース頻度もありません。そうでない場合
必要なもの

、フォークしてください — それがライセンスの目的です。
壊れるときは大きな音で壊れます: 理由は設定→にあります
診断と ~/.claude-fleet/last-poll.log 。のいずれかに含めてください
問題。 SUPPORT.md を参照してください。
セッション。それは、それらが含まれている Git リポジトリごとにグループ化されたということです。
マシン名、接続ステータス、エラー行、「到達不能」警告はありません。
リモート マシンがネットワーク外にある場合、リストは単に空で落ち着いたものになります。
アプリがアクセスできないもの、または非表示にすることを選択したものは、「設定」→「設定」でのみ説明されます。
診断 — セッション リスト自体には決して含まれません。
グループ化は、 acme-web のディレクトリ: セッションごとではなく、 git repo ごとに行われます。
acme-web/services/api と acme-web/worker は 1 つの acme-web グループにまとめられます
3つに分散するのではなく。行にはリポジトリ内のパスが表示されるため、
それでもサブフォルダーのセッションを区別します。
ドット
状態
意味
🟠
承認が必要です
ツール呼び出しが発行されましたが、それ以降何も動かず、許可プロンプトに応答がありませんでした。これがアプリの状態です。
🟡
あなたの番です
クロードはターンを終了し、あなたが何かを入力するのを待っています
🟢
働く
ツール呼び出しが進行中であり、トランスクリプトがアクティブに移動中です
⚪️
アイドル状態
あなたが立ち去った「あなたの番です」セッション — 灰色、強調されていない、カウントされていない
承認が必要なセッションを含むリポジトリが一番上に表示されるため、
「あなたを必要としている」というメッセージは常に最初に目に入るものです。メニューバーバッジは、
今あなたを必要とするセッション (承認が必要 + あなたの番)。
macOS 14 以降と Xcode コマンド ライン ツール ( xcode-select --install ) が必要です。
git clone < this-repo > クロードフリート
CDクロードフリート
./build.sh
~ /Applications/ClaudeFleet.app を開きます
build.sh はコンパイル、署名 (デフォルトではアドホック - Apple アカウントは必要ありません)、およびインストールを行います。
~/Applications/ClaudeFleet.app に。これは LSUIElement アプリです: メニュー バーのみ、

ドックなし
アイコン、アプリスイッチャーのエントリはありません。
ログイン時に開始: システム設定 → 一般 → ログイン項目 → + →
~/アプリケーション/ClaudeFleet.app 。 (代わりに行われるのではなく文書化されるため、何も行われません
自分自身をサイレントにログインに追加します)。
初回起動時に、macOS は、Claude Fleet がローカル上のデバイスにアクセスすることを許可するかどうかを尋ねます。
ネットワーク — [許可] をクリックしないと、リモート マシンにアクセスできません。参照
プロンプトが表示されない場合は、docs/troubleshooting.md。
マシン — ~/.ssh/config のエイリアス (すでに SSH または
VS Code リモート SSH を使用すると、そのユーザー、ポート、キーが継承されます。空白のままにしておきます
これは Mac のみです。
この Mac を含めます — デフォルトでオンです。リモコンのみを視聴するにはオフにします。
更新間隔 — 5 秒 / 15 秒 / 30 秒 / 1 分 (デフォルトは 15 秒)。
リモートには何もインストールする必要はありません。コレクターは SSH 経由でパイプされ、
stdlib のみの Python。完全なセットアップについては docs/user-guide.md を参照してください。
Windows ホストを含む。
ClaudeFleet.app ──15 秒ごと──▶ python3collector.py (この Mac)
──15秒ごと──▶ ssh <ホスト> '<py|python3> -' <collecter.py (リモート)
│
└─▶ フレーム化された JSON ──▶ リポジトリ グループ ──▶ メニュー バー
collector.py は、クロード コードが各マシンに書き込む 2 つの内容 (プロセスごと) を読み取ります。
レジストリ ( ~/.claude/sessions/<pid>.json ) とセッションごとのトランスクリプト
( ~/.claude/projects/<slug>/<sessionId>.jsonl ) — 1 つの小さな JSON ドキュメントを出力します。
SSH stdin 経由でパイプされるため、リモートには何もインストールされず、何もインストールされません。
そこで同期を保つ必要があります。
完全な設計と、それを自明ではないものにするトラップは次のとおりです。
docs/architecture.md および docs/traps.md 。
ユーザー ガイド — セットアップ、Windows ホスト、バックグラウンド エージェント、右クリック アクション
アーキテクチャ — コレクター/アプリの分割、状態の導出、稼働性、セキュリティ モデル
トラップ — 苦労して勝ち取った落とし穴: ストアスタブ Python、

macOS ローカル ネットワークのデッドロック、あらゆるテストを偽る SSH コントロール ソケット
トラブルシューティング — 空のリスト、プロンプトなし、通知
貢献・サポート・ライセンス (MIT)
オフライン プレビュー — 実際のパネルを PNG にラスタライズします。画面は必要ありません。
swiftc -O -parse-as-library -DPREVIEW -target arm64-apple-macos14.0 \
-o /tmp/cfpreview Sources/ * .swift && /tmp/cfpreview fixture.json /tmp/panel.png dark
テスト:
python3 -m Unittest Discover -s テスト -t テスト # コレクター ロジック
./swift-tests/run.sh # 通知機能 + セキュリティチェック
fixture.json は合成マルチリポジトリのスナップショットです。同じファイルがアプリ自体を駆動します
CLAUDE_FLEET_FIXTURE=/path/to/fixture.json 経由で、SSH を完全にバイパスします。
クロード コード セッション (ローカル + SSH 経由のリモート) をリポジトリごとにグループ化して表示する macOS メニュー バー アプリ。どのセッションが待機しているかを確認できます。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS menu bar app that shows your Claude Code sessions (local + remote over SSH), grouped by repo, so you can see which one is waiting on you - Shubhamgupta2612/claude-fleet

GitHub - Shubhamgupta2612/claude-fleet: macOS menu bar app that shows your Claude Code sessions (local + remote over SSH), grouped by repo, so you can see which one is waiting on you · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Shubhamgupta2612
/
claude-fleet
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE .superpowers .superpowers Sources Sources docs docs fixtures fixtures swift-tests swift-tests tests tests tools tools .gitignore .gitignore AppIcon.icns AppIcon.icns CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SUPPORT.md SUPPORT.md build.sh build.sh collector.py collector.py fixture.json fixture.json View all files Repository files navigation
A native macOS menu bar app that shows every Claude Code session you have
running — on this Mac and on a remote machine over SSH — grouped by git repo, so
you can tell at a glance which one is stuck waiting on you.
The one thing it exists to tell you: which session is sitting at an unanswered
permission prompt right now. When Claude asks "can I run this command?" inside a
window you're not looking at, that session freezes and you have no idea until you
happen to switch to it. Claude Fleet puts that in your menu bar — and when a session
starts needing you, it jumps to the top.
An independent, community-built tool. Not affiliated with, or endorsed by, Anthropic.
Support posture — read this first
This is a personal tool, built for one setup and published as-is under MIT. It
reads Claude Code's on-disk files, which are undocumented and can change in
any Claude Code release — when they do, this breaks until someone fixes it.
I fix what breaks for me. Issues are open but may go unanswered; PRs are
welcome but may sit. There is no roadmap and no release cadence. If that's not
what you need, fork it — that's what the licence is for.
When it does break, it breaks loudly : the reason lands in Settings →
Diagnostics and in ~/.claude-fleet/last-poll.log . Please include that in any
issue. See SUPPORT.md .
Sessions. That's it — grouped by the git repo they're in.
No machine names, no connection status, no error rows, no "unreachable" warnings.
If the remote machine is off-network the list is simply empty and calm .
Anything the app can't reach, or chooses to hide, is explained only in Settings →
Diagnostics — never in the session list itself.
Grouping is by git repo , not by directory: sessions in acme-web ,
acme-web/services/api and acme-web/worker collapse into one acme-web group
rather than scattering into three. The row shows the path within the repo so you can
still tell subfolder sessions apart.
dot
state
meaning
🟠
Needs approval
A tool call was issued and nothing has moved since — an unanswered permission prompt. This is the state the app is for.
🟡
Your turn
Claude finished its turn and is waiting for you to type something
🟢
Working
A tool call is in flight and the transcript is actively moving
⚪️
Idle
A "your turn" session you walked away from — grey, de-emphasised, uncounted
Repos containing a Needs approval session float to the top, so the thing that
needs you is always the first thing you see. The menu bar badge counts only the
sessions that need you now ( Needs approval + Your turn ).
Requires macOS 14+ and the Xcode command line tools ( xcode-select --install ).
git clone < this-repo > claude-fleet
cd claude-fleet
./build.sh
open ~ /Applications/ClaudeFleet.app
build.sh compiles, signs (ad-hoc by default — no Apple account needed), and installs
to ~/Applications/ClaudeFleet.app . It's an LSUIElement app: menu bar only, no Dock
icon, no app-switcher entry.
Start at login: System Settings → General → Login Items → + →
~/Applications/ClaudeFleet.app . (Documented rather than done for you, so nothing
adds itself to your login silently.)
On first launch , macOS asks to allow Claude Fleet to access devices on your local
network — click Allow , or the remote machine can never be reached. See
docs/troubleshooting.md if no prompt appears.
Machine — an alias from your ~/.ssh/config (the same one you already SSH or
VS Code Remote-SSH with, so its user, port and key are inherited). Leave blank for
this-Mac-only.
Include this Mac — on by default; turn it off to watch only the remote.
Refresh every — 5s / 15s / 30s / 1m (default 15s).
The remote needs nothing installed: the collector is piped over SSH and is
stdlib-only Python. See docs/user-guide.md for the full setup,
including Windows hosts.
ClaudeFleet.app ──every 15s──▶ python3 collector.py (this Mac)
──every 15s──▶ ssh <host> '<py|python3> -' < collector.py (remote)
│
└─▶ framed JSON ──▶ repo groups ──▶ menu bar
collector.py reads two things Claude Code writes on each machine — a per-process
registry ( ~/.claude/sessions/<pid>.json ) and a per-session transcript
( ~/.claude/projects/<slug>/<sessionId>.jsonl ) — and prints one small JSON document.
It is piped over SSH stdin , so nothing is ever installed on the remote and nothing
needs to be kept in sync there.
The full design, and the traps that make it non-trivial, are in
docs/architecture.md and docs/traps.md .
User guide — setup, Windows hosts, background agents, the right-click actions
Architecture — the collector/app split, state derivation, liveness, the security model
Traps — the hard-won gotchas: the Store-stub Python, the macOS Local Network deadlock, the SSH control-socket that makes every test lie
Troubleshooting — empty list, no prompt, notifications
Contributing · Support · Licence (MIT)
Offline preview — rasterises the real panel to a PNG, no screen needed:
swiftc -O -parse-as-library -DPREVIEW -target arm64-apple-macos14.0 \
-o /tmp/cfpreview Sources/ * .swift && /tmp/cfpreview fixture.json /tmp/panel.png dark
Tests:
python3 -m unittest discover -s tests -t tests # collector logic
./swift-tests/run.sh # notifier + security checks
fixture.json is a synthetic multi-repo snapshot; the same file drives the app itself
via CLAUDE_FLEET_FIXTURE=/path/to/fixture.json , bypassing SSH entirely.
macOS menu bar app that shows your Claude Code sessions (local + remote over SSH), grouped by repo, so you can see which one is waiting on you
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/grzegorz-raczek-unit8/claude-quota"
hn_url: "https://news.ycombinator.com/item?id=48473845"
title: "Show HN: macOS menu bar gauges for your Claude Code quota"
article_title: "GitHub - grzegorz-raczek-unit8/claude-quota: Menu bar gauges for Claude Code quota (SwiftBar plugin) · GitHub"
author: "grzracz"
captured_at: "2026-06-10T10:29:22Z"
capture_tool: "hn-digest"
hn_id: 48473845
score: 2
comments: 0
posted_at: "2026-06-10T09:43:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: macOS menu bar gauges for your Claude Code quota

- HN: [48473845](https://news.ycombinator.com/item?id=48473845)
- Source: [github.com](https://github.com/grzegorz-raczek-unit8/claude-quota)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T09:43:16Z

## Translation

タイトル: HN を表示: クロード コード クォータの macOS メニュー バー ゲージ
記事タイトル: GitHub - grzegorz-raczek-unit8/claude-quota: クロード コード クォータのメニュー バー ゲージ (SwiftBar プラグイン) · GitHub
説明: クロード コード クォータのメニュー バー ゲージ (SwiftBar プラグイン) - grzegorz-raczek-unit8/claude-quota

記事本文:
GitHub - grzegorz-raczek-unit8/claude-quota: クロード コード クォータのメニュー バー ゲージ (SwiftBar プラグイン) · GitHub
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
grzegorz-raczek-unit8
/
クロードクォータ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

イオンオプション
コード
grzegorz-raczek-unit8/claude-quota
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット docs docs .gitignore .gitignore ライセンス ライセンス README.md README.md claude-quota.5m.py claude-quota.5m.py install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード クォータのメニュー バー ゲージ — アカウントごとに 1 つのバー、次のようになります。
(暗いメニュー バー用に描画 - 白い輪郭)
各バーは、1 つのアカウントの 5 時間ウィンドウの使用率を色分けして示します。
緑 / オレンジ (≧70%) / 赤 (≧90%)。
5 時間のウィンドウが完全に使用されると、バーにはカウントダウンが表示されます。
パーセンテージの代わりにリセット ( 4:28 )。
週間制限に達すると、バーが黒くなり、カウントダウンが表示されます
毎週のリセット ( 2D ) — 5 時間であっても、これはより厳しい上限です
窓が言う。
ドロップダウンには、すべてのアカウントの完全な詳細がインラインでリストされます: 5 時間および毎週
ウィンドウ、プランがレポートするモデルごとのウィンドウ、追加使用量
クレジットとリセット時間。
5 分ごとに更新 (SwiftBar のファイル名規則) およびマニュアル
「今すぐ更新」エントリ。
カール -fsSL https://raw.githubusercontent.com/grzegorz-raczek-unit8/claude-quota/main/install.sh |バッシュ
macOS と Homebrew が必要です。 macOS でキーチェーンが表示される場合
最初の更新時に許可ダイアログが表示されるので、「常に許可」をクリックします。
プラグインは、macOS キーチェーンから Claude Code OAuth トークンを読み取ります。
( 読み取り専用 — トークンを更新したり書き換えたりすることはないため、ログを記録することはできません
out)、Claude Code の /usage 画面と同じ使用エンドポイントをクエリします。
を使用します。パスワード、スクレイピング、サードパーティのサービスはありません。
注: このエンドポイントはクロード コードの内部にあり、文書化されていないため、
将来のクロード コードの変更では、ここでの小さな修正が必要になる可能性があります。
git clone https://github.com/grzegorz-raczek-unit8/claude-quota.git
cd クロードクォータ
./インスタ

ll.sh
どちらのインストール パスでも SwiftBar がセットアップされます
お持ちでない場合はHomebrew経由で。アカウントに ⚠ が表示されている場合、そのトークンは失われています
使用されなくなったため古い — クロード CLI を 1 回実行すると、ウィジェットが回復します。
次のサイクル。
デフォルトでは、プラグインは次のようにアカウントを自動検出します: ~/.claude /
クロード コード キーチェーン エントリを持つ ~/.claude-* config ディレクトリは、
bar、ディレクトリ接尾辞 ( ~/.claude-work → W ) でラベル付けされます。シングル
自動検出されたアカウントには文字ラベルは表示されず、バーのみが表示されます。
アカウントを固定するか名前を変更するには（複数の CLAUDE_CONFIG_DIR を使用する場合など）、
~/.config/claude-quota/accounts に 1 行に 1 つのパス [ラベル] を指定
(単一単語のラベル):
~/.claude-work 仕事
~/.claude-priv プライベート
アカウントのメニュー バー ゲージを非表示にするには (ドロップダウンの詳細は残ります)、次を使用します。
ドロップダウンのそのアカウントの行の下にあるメニュー バーを非表示にするか、編集します
~/.config/claude-quota/hidden (1 行に 1 つのラベル)。
CLAUDE_CONFIG_DIR を介した複数のアカウントは、シェル rc では次のようになります。
クロード () { CLAUDE_CONFIG_DIR= " $HOME /.claude-work " コマンド クロード " $@ " ; }
claude-priv () { CLAUDE_CONFIG_DIR= " $HOME /.claude-priv " コマンド クロード " $@ " ; }
アンインストール
SwiftBar プラグイン フォルダーから claude-quota.5m.py を削除します。
(デフォルトでは ~/.swiftbar)。
クロード コード クォータのメニュー バー ゲージ (SwiftBar プラグイン)
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Menu bar gauges for Claude Code quota (SwiftBar plugin) - grzegorz-raczek-unit8/claude-quota

GitHub - grzegorz-raczek-unit8/claude-quota: Menu bar gauges for Claude Code quota (SwiftBar plugin) · GitHub
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
grzegorz-raczek-unit8
/
claude-quota
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
grzegorz-raczek-unit8/claude-quota
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits docs docs .gitignore .gitignore LICENSE LICENSE README.md README.md claude-quota.5m.py claude-quota.5m.py install.sh install.sh View all files Repository files navigation
Menu bar gauges for your Claude Code quota — one bar per account, like this:
(drawn for dark menu bars — white outlines)
Each bar shows the 5-hour-window utilization for one account, colored
green / orange (≥70%) / red (≥90%).
When the 5-hour window is fully used, the bar shows a countdown until
reset ( 4:28 ) instead of the percentage.
When the weekly limit is hit, the bar turns black with a countdown
to the weekly reset ( 2D ) — that's the harder cap, whatever the 5-hour
window says.
The dropdown lists full detail for every account inline: 5-hour and weekly
windows, per-model windows where your plan reports them, extra-usage
credits, and reset times.
Refreshes every 5 minutes (SwiftBar filename convention) plus a manual
"Refresh now" entry.
curl -fsSL https://raw.githubusercontent.com/grzegorz-raczek-unit8/claude-quota/main/install.sh | bash
Requires macOS and Homebrew . When macOS shows a Keychain
permission dialog on the first refresh, click Always Allow .
The plugin reads your Claude Code OAuth token from the macOS Keychain
( read-only — it never refreshes or rewrites tokens, so it can't log you
out) and queries the same usage endpoint that Claude Code's /usage screen
uses. No passwords, no scraping, no third-party services.
Note: that endpoint is internal to Claude Code and undocumented, so a
future Claude Code change may require a small fix here.
git clone https://github.com/grzegorz-raczek-unit8/claude-quota.git
cd claude-quota
./install.sh
Either install path sets up SwiftBar
via Homebrew if you don't have it. If an account shows ⚠, its token went
stale from disuse — run that claude CLI once and the widget recovers on the
next cycle.
By default the plugin auto-discovers accounts: every ~/.claude /
~/.claude-* config directory that has a Claude Code Keychain entry gets a
bar, labeled by the directory suffix ( ~/.claude-work → W ). A single
auto-discovered account shows no letter label — just the bar.
To pin or rename accounts (e.g. you use multiple CLAUDE_CONFIG_DIR s), create
~/.config/claude-quota/accounts with one path [label] per line
(single-word labels):
~/.claude-work Work
~/.claude-priv Priv
To hide an account's menu bar gauge (its dropdown detail stays), use
Hide from menu bar under that account's row in the dropdown — or edit
~/.config/claude-quota/hidden (one label per line).
Multiple accounts via CLAUDE_CONFIG_DIR look like this in your shell rc:
claude () { CLAUDE_CONFIG_DIR= " $HOME /.claude-work " command claude " $@ " ; }
claude-priv () { CLAUDE_CONFIG_DIR= " $HOME /.claude-priv " command claude " $@ " ; }
Uninstall
Delete claude-quota.5m.py from your SwiftBar plugin folder
( ~/.swiftbar by default).
Menu bar gauges for Claude Code quota (SwiftBar plugin)
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

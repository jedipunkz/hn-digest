---
source: "https://github.com/sarhej/claude-fix"
hn_url: "https://news.ycombinator.com/item?id=48467458"
title: "Show HN: Separate Work/Personal Claude Desktop profiles on one Mac(bash,no deps)"
article_title: "GitHub - sarhej/claude-fix · GitHub"
author: "supersergio"
captured_at: "2026-06-09T21:38:10Z"
capture_tool: "hn-digest"
hn_id: 48467458
score: 1
comments: 0
posted_at: "2026-06-09T20:42:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Separate Work/Personal Claude Desktop profiles on one Mac(bash,no deps)

- HN: [48467458](https://news.ycombinator.com/item?id=48467458)
- Source: [github.com](https://github.com/sarhej/claude-fix)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T20:42:31Z

## Translation

タイトル: HN を表示: 1 台の Mac 上で仕事用/個人用のクロード デスクトップ プロファイルを分離 (bash、deps なし)
記事のタイトル: GitHub - sarhej/claude-fix · GitHub
説明: GitHub でアカウントを作成して、sarhej/claude-fix の開発に貢献します。
HN テキスト: Claude Desktop = インストールごとに 1 つのプロファイル。私は 1 台の Mac で仕事と個人を使用しています。小さな bash スクリプトは --user-data-dir を使用して別のランチャーを作成します。デプスも、Claude.app も変更しません。 macOSのみ。

記事本文:
GitHub - sarhej/claude-fix · GitHub
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
サルヘイジ
/
クロードフィックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
4 コミット 4 コミット .github/ workflows .github/ workflows テスト テスト ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md make_claude_launchers.sh make_claude_launchers.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Mac 用クロード プロファイル スイッチャー。
1 台の Mac 上で、仕事用と個人用の別々のクロード デスクトップ プロファイルを実行します。
独自のログイン、チャット履歴、設定、接続されたツールを使用します。
カール -fsSL https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh |バッシュ
デフォルトの設定では、既存の Claude ログインがそのまま維持され、
欠落している 2 番目のプロファイル:
通常の Claude.app は現在のアカウント (仕事用または個人用) に残ります。
スクリプトは、他のアカウント用に 1 つの新しいランチャーを作成します。
~/アプリケーション/Claude Personal.app
オプションで、デスクトップ上のクリック可能なコピー
新しいプロファイルが開いたら、他のアカウントでサインインし、
そこに一致するツールがあります。
信頼性チェック: このプロジェクトはローカルのみで依存関係がなく、以下によってカバーされています。
GitHub Actions での macOS 統合テスト。ソースとテスト スイートを確認します。
完全な詳細が必要な場合は、実行前に GitHub にアクセスしてください。
make_claude_launchers.sh は、ドックに適した個別のランチャーを作成します。
Claude デスクトップ — それぞれに独自のログインがあり、
チャット履歴、設定、接続されているツール。仕事と個人を考える
(異なる電子メール、カレンダー、Slack など)、またはクライアントごとに 1 つのプロファイル。
生成されたランチャーを削除して復元するクリーン コマンドも含まれています。
Mac を標準の単一アプリ設定にします。
Claude Desktop (ほとんどの Electron/Chromium アプリと同様) にはすべてのデータが保存されます。
サインインしているアカウント、会話、設定、
接続された統合 — 単一のユーザー データ ディレクトリ内。つまり 1 つ
アプリ = 1 つのアカウント = 1 セットの接続されたツール。
クロードをexで使う場合

社内サービス (Gmail、Google カレンダー、Slack、
Notion など)、すべてがその 1 つのプロファイルに接続されています。簡単にはできません
1 人のクロードが個人用の Gmail とカレンダーにサインインし、もう 1 人がクロードになるようにします。
コンテキスト、アカウント、
または権限。
このスクリプトは、プロファイルごとにクロードを起動することでこの問題を回避します。
--user-data-dir なので、各ランチャーは完全に分離されます。
それぞれに異なる Claude アカウントにサインインします。
プロファイルごとに異なるツールを接続します — 例:個人のメールアドレスとカレンダー
1 つは会社のメールとカレンダー、もう 1 つは会社のメールとカレンダーです。
仕事上のチャットと個人的なチャットを分離してください (コンテキストの相互汚染は避けてください)。
複数のクロード ウィンドウを一度に実行します (プロファイルごとに 1 つ)。
アプリの修正、再署名、パッチ適用は不要 - 薄い AppleScript だけです
既存の Claude.app を指すランチャー。
個人用クロード アカウントと会社用クロード アカウントを持つ人々。
クロードをさまざまなクライアントの電子メール、カレンダー、Slack、
Notion、GitHub、または MCP/ツールの権限。
を使用しながら個人的なコンテキストと仕事のコンテキストを分離したい人
ネイティブの Claude デスクトップ アプリ。
macOS (スクリプトは他の場所では実行を拒否します)。
Claude デスクトップがインストールされました — ここからダウンロードしてください。
スクリプトは、Launch を介して /Applications 、 ~/Applications でそれを自動検出します。
サービス (バンドル ID) または Spotlight。
組み込み macOS ツール: osacompile および osascript (プリインストール)。
Python、Pillow、アイコン変換、パッケージのインストールは必要ありません。
これをコピーしてターミナルに貼り付けます。
カール -fsSL https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh |バッシュ
スクリプトはいくつかの簡単な質問をします。 Enter キーを押してデフォルトを受け入れます。
現在のクロードがすでにサインインしているアカウントを確認します
不足している 2 番目のプロファイル ランチャーを作成する
オプティオ

最後に、クリック可能なコピーをデスクトップに配置します
必要に応じて、新しいプロファイルをすぐに起動して、そこにサインインできるようにします
ランチャーが存在した後、スクリプトを再度実行すると、代わりに管理メニューが表示されます
最初の質問をもう一度すること。
スクリプトを実行する前に検査したい場合は、次のようにします。
カール -fsSLO https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh
make_claude_launchers.sh を少なくする
bash make_claude_launchers.sh
リポジトリのクローンを作成する
git clone https://github.com/sarhej/claude-fix.git
cd クロードフィックス
bash make_claude_launchers.sh
実行可能権限はオプションです。 bash make_claude_launchers.sh ... 動作します。
# 対話型セットアップ: 既存のクロード ログインを保持し、不足しているプロファイルを追加します
./make_claude_launchers.sh
# パワー ユーザー向けにカスタム ランチャーを作成します (ラベルごとに 1 つ)
./make_claude_launchers.sh 仕事用個人クライアントを作成する
# ランチャーを作成し、クリック可能なコピーをデスクトップに配置します
./make_claude_launchers.sh create --desktop Personal
# 不足しているプロファイルを作成し、すぐに起動します
./make_claude_launchers.sh create --launch
# デスクトップのコピーをスキップする
./make_claude_launchers.sh create --no-desktop Personal
# 生成されたランチャーを削除しますが、プロファイル データは保持します
./make_claude_launchers.sh クリーン
# 生成されたランチャーとそのプロファイル データを削除します (削除するたびに尋ねられます)
./make_claude_launchers.sh clean --purge
# ヘルプを表示
./make_claude_launchers.sh ヘルプ
作成後、スクリプトは Finder で ~/Applications を開きます (まだ開いていない場合)。
新しいプロフィールを立ち上げました。
新しいプロファイルをすぐに起動することを選択した場合、その新しいクロード ウィンドウのみが表示されます。
が開きます。他のアカウントでサインインし、マッチング ツールに接続します。
すでにお持ちのアカウントの通常の Claude アプリを引き続き使用してください。
生成されたプロファイルがすでに存在すると、スクリプトは次の方法でそれらを検出します。

相続人マーカー
ファイルを作成し、初回実行のオンボーディングをスキップします。代わりに、次のステップが示されます。
クロード プロファイル スイッチャーはすでにセットアップされています。
生成されたプロファイル ランチャーが見つかりました:
- クロード・パーソナル
ランチャー: ~/Applications/Claude Personal.app
ローカルサインイン: ~/ClaudePersonal
次に何をしたいですか?
1) 生成されたクロード プロファイルを開きます
2) 生成されたすべてのクロード プロファイルを開く
3) ランチャーフォルダーを開きます
4) 別のプロファイルを作成する
5) 間違ったアカウントを修正する/生成されたプロファイルを新たに開始する
6) 生成されたランチャーを削除します (ローカルサインインを維持します)
7) ランチャーとローカルで生成されたプロファイル データを削除します
8) キャンセル
生成されたランチャーが間違ったアカウントを開く場合は、「新規開始」を使用します。それだけ
Mac 上のランチャーのローカル サインイン フォルダーをクリアします。あなたのものは削除されません
クロードアカウントは通常のクロードアプリと変わりません。
コマンド
何をするのか
コマンドがありません
最初の実行: 不足しているプロファイルを追加します。後で: 管理メニューを表示します。
[ラベル...] を作成する
明示的なラベルごとに 1 つのランチャーを作成します。
create --desktop [ラベル...]
また、クリック可能なランチャーのコピーをデスクトップに配置します。
create --no-desktop [ラベル...]
~/Applications にのみ作成します。
create --launch [ラベル...]
セットアップ後に作成したプロファイルを起動します。
create --no-launch [ラベル...]
クロードを開かずにランチャーを作成します。
作成 --はい
メニューをスキップします。既存の仕事を引き継ぎ、個人を作成します。
きれいな
生成されたランチャーを削除します。すべてのプロフィールデータを保持します。
クリーン --パージ
ランチャーと生成されたプロファイル データを削除します (プロファイルごとの確認)。
助けて
印刷の使用法。
何が生み出されるのか
現在のクロードがすでに仕事/会社にサインインしている場合、スクリプトは
Claude Personal.app のみを使用し、通常の Claude アプリはそのままにしておきます。
現在の Claude が Personal の場合、 Claude Work.app のみが作成されます。
Work のようなラベルの場合、スクリプトは以下を生成します。
~/Applications/Claude Work.app — 小さなアプリ

実行する pleScript アプレット:
open -n -a ' /Applications/Claude.app ' --args --user-data-dir= " $HOME /ClaudeWork "
~/ClaudeWork/ — 分離されたプロファイル データ ディレクトリ (Claude によって作成)
最初の起動時;そのプロファイルのログイン、履歴、設定が保持されます)。
オプションのデスクトップ コピー — などのクリック可能なコピー
~/Desktop/Claude Work.app 、からアプリを起動する予定の人向け
デスクトップ。
ランチャー名はデータ ディレクトリに決定的にマッピングされます (「ビッグ クライアント」 →
~/ClaudeBigClient ) なので、create と clean は常にパスに一致します。
ラベルは意図的に控えめになっています: 文字、数字、スペース、ハイフン、
アンダースコアとアポストロフィがサポートされています。パス区切り文字、制御文字、
先頭のドット、長すぎるラベル、同じものにマップされる重複したラベル
プロファイルディレクトリは拒否されます。
健全性チェック — macOS と必要なツールが存在することを確認します。
クロードを見つける — 標準パスを確認し、サービスを起動してから、
スポットライト。見つからない場合は、明確なメッセージ (およびダウンロード リンク) を表示して終了します。
クロードの通常のアイコンをコピーします — ランチャーは、標準のクロード アイコンを使用します。
利用可能です。色合いやアイコンの変換は試行されません。
コンパイル ランチャー — osacompile は 1 行の open -n ... --user-data-dir を変換します
コマンドを実際の .app バンドルに追加します。
Finder メタデータを更新 — バンドルをタッチして、
Finder が小銭を拾います。
-n は、macOS に新しいインスタンスを起動するように指示します。これにより、いくつかのインスタンスが可能になります。
同時に実行するクロード プロファイル。
重要な警告: プロファイルの分離は、Claude Desktop が引き続き実行されるかどうかに依存します。
Electron/Chromium の --user-data-dir フラグを尊重します。この用途では今日でも機能します
ただし、Anthropic は将来のリリースで Claude Desktop を変更する可能性があります。
Dock アイコンに関する注意事項: ランチャーはクロードの通常のアイコンと実行プロファイルを使用します。
Dock では別個の同じ外観のクロード アイコンとして表示される場合があります。これは私

は
同じ基盤となる /Applications/Claude.app を起動する場合の macOS の制限
異なるプロファイル データ ディレクトリを使用してプロセスを実行します。
create はプロファイル データを削除しません。削除して再構築するだけです。
.app ランチャー バンドル。 ~/Claude* ディレクトリは保存されます。
clean は、このスクリプトが生成したランチャーのみを削除します。それは彼らを識別します
アプリバンドル内にclaude-fixで生成されたマーカーファイルを使用し、決して使用しない
本物の Claude.app に触れます。
clean --purge は各削除の前に質問し、デフォルトでは No になっているため、
誤って Enter を押すとデータが保持されます。生成されたプロファイルのみをターゲットとしてパージします
~/ClaudePersonal などのディレクトリ。通常のクロードを削除することはありません
アプリまたはそのデフォルトのプロファイルデータ。
依存関係やパッケージのインストールはありません。スクリプトは組み込みの macOS のみを使用します
ツールを使用し、Python パッケージをインストールしません。
Claude.app の変更、パッチ、再署名は行いません。
アカウント制限、認証、または会社のポリシーをバイパスすることはありません。
プロファイル間でチャット履歴を移行または結合しません。
クロードのプロフィール データは暗号化されません。プロファイルは通常どおり保存されます
ローカルのクロード/エレクトロン ユーザー データ。
--user-data-dir に対する Anthropic の公式サポートは保証されません。
見た目が異なる実行中の Dock アイコンは作成されません。プロフィールは、
別々に表示されますが、ID

[切り捨てられた]

## Original Extract

Contribute to sarhej/claude-fix development by creating an account on GitHub.

Claude Desktop = one profile per install. I use work + personal on one Mac. Small bash script creates separate launchers with --user-data-dir. No deps, no modify Claude.app. macOS only.

GitHub - sarhej/claude-fix · GitHub
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
sarhej
/
claude-fix
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .github/ workflows .github/ workflows tests tests LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md make_claude_launchers.sh make_claude_launchers.sh View all files Repository files navigation
Claude Profile Switcher for Mac.
Run separate Work and Personal Claude Desktop profiles on one Mac, each
with its own login, chat history, settings, and connected tools.
curl -fsSL https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh | bash
The default setup keeps your existing Claude login as-is and creates only
the missing second profile:
Your normal Claude.app stays on your current account (Work or Personal)
The script creates one new launcher for the other account, e.g.
~/Applications/Claude Personal.app
Optionally, a clickable copy on your Desktop
When the new profile opens, sign in with your other account and connect the
matching tools there.
Trust check: this project is local-only, dependency-free, and covered by
macOS integration tests in GitHub Actions. Review the source and test suite on
GitHub before running if you want the full details.
make_claude_launchers.sh creates separate, dock-friendly launchers for
Claude Desktop — each with its own login,
chat history, settings, and connected tools . Think Work vs Personal
(different email, calendar, Slack, etc.), or one profile per client.
It also includes a clean command to remove the generated launchers and restore
your Mac to the standard single-app setup.
Claude Desktop (like most Electron/Chromium apps) stores all of its data —
the account you're signed into, your conversations, your preferences, and your
connected integrations — in a single user-data directory. That means one
app = one account = one set of connected tools.
If you use Claude with external services (Gmail, Google Calendar, Slack,
Notion, etc.), everything is wired to that single profile. You can't easily
have one Claude signed into your personal Gmail and calendar while another
uses your company email and calendar — without mixing contexts, accounts,
or permissions.
This script works around that by launching Claude with a per-profile
--user-data-dir , so each launcher is fully isolated:
Sign into a different Claude account in each.
Connect different tools per profile — e.g. personal email & calendar in
one, company email & calendar in another.
Keep work and personal chats separate (no cross-contamination of context).
Run multiple Claude windows at once (one per profile).
No app modification, no re-signing, no patching — just thin AppleScript
launchers that point at your existing Claude.app .
People with a private Claude account and a company Claude account .
Consultants who connect Claude to different clients' email, calendars, Slack,
Notion, GitHub, or MCP/tool permissions.
Anyone who wants to keep personal and work context separate while still using
the native Claude Desktop app.
macOS (the script refuses to run anywhere else).
Claude Desktop installed — download here .
The script auto-detects it in /Applications , ~/Applications , via Launch
Services (bundle id), or Spotlight.
Built-in macOS tools: osacompile and osascript (preinstalled).
No Python, Pillow, icon conversion, or package installation is required.
Copy and paste this into Terminal:
curl -fsSL https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh | bash
The script will ask a few simple questions. Press Enter to accept the default:
Confirm which account your current Claude is already signed into
Create the missing second profile launcher
Optionally place a clickable copy on your Desktop
Optionally launch the new profile immediately so you can sign in there
After launchers exist, running the script again shows a management menu instead
of asking the first-run question again.
If you want to inspect the script before running it:
curl -fsSLO https://raw.githubusercontent.com/sarhej/claude-fix/heads/main/make_claude_launchers.sh
less make_claude_launchers.sh
bash make_claude_launchers.sh
Clone the Repo
git clone https://github.com/sarhej/claude-fix.git
cd claude-fix
bash make_claude_launchers.sh
Executable permissions are optional; bash make_claude_launchers.sh ... works.
# Interactive setup: keep existing Claude login, add the missing profile
./make_claude_launchers.sh
# Create custom launchers (one per label) for power users
./make_claude_launchers.sh create Work Personal Clients
# Create launchers and also put clickable copies on your Desktop
./make_claude_launchers.sh create --desktop Personal
# Create the missing profile and launch it immediately
./make_claude_launchers.sh create --launch
# Skip Desktop copies
./make_claude_launchers.sh create --no-desktop Personal
# Remove generated launchers, but KEEP your profile data
./make_claude_launchers.sh clean
# Remove generated launchers AND their profile data (asks before each delete)
./make_claude_launchers.sh clean --purge
# Show help
./make_claude_launchers.sh help
After creation, the script opens ~/Applications in Finder unless it already
launched the new profile for you.
If you choose to launch the new profile right away, only that new Claude window
opens. Sign into it with your other account, then connect the matching tools.
Keep using your normal Claude app for the account you already had.
Once generated profiles already exist, the script detects them by their marker
files and skips first-run onboarding. Instead it shows next steps:
Claude Profile Switcher is already set up.
Found generated profile launcher(s):
- Claude Personal
launcher: ~/Applications/Claude Personal.app
local sign-in: ~/ClaudePersonal
What would you like to do next?
1) Open a generated Claude profile
2) Open all generated Claude profiles
3) Open the launchers folder
4) Create another profile
5) Fix wrong account / start fresh for a generated profile
6) Remove generated launchers (keep local sign-ins)
7) Remove launchers AND local generated profile data
8) Cancel
Use start fresh if a generated launcher opens the wrong account. It only
clears that launcher's local sign-in folder on your Mac; it does not delete your
Claude account and does not change the normal Claude app.
Command
What it does
no command
First run: add missing profile. Later: show the management menu.
create [labels...]
Create one launcher per explicit label.
create --desktop [labels...]
Also place clickable launcher copies on your Desktop.
create --no-desktop [labels...]
Create only in ~/Applications .
create --launch [labels...]
Launch the created profile(s) after setup.
create --no-launch [labels...]
Create launchers without opening Claude.
create --yes
Skip menu; assumes existing Work, creates Personal.
clean
Remove generated launchers; keep all profile data.
clean --purge
Remove launchers and generated profile data (per-profile confirmation).
help
Print usage.
What gets created
If your current Claude is already signed into Work / Company, the script creates
only Claude Personal.app and leaves your normal Claude app untouched.
If your current Claude is Personal, it creates only Claude Work.app .
For a label like Work , the script produces:
~/Applications/Claude Work.app — a tiny AppleScript applet that runs:
open -n -a ' /Applications/Claude.app ' --args --user-data-dir= " $HOME /ClaudeWork "
~/ClaudeWork/ — the isolated profile data directory (created by Claude
on first launch; holds that profile's login, history, and settings).
Optional Desktop copies — clickable copies like
~/Desktop/Claude Work.app , for people who expect to launch apps from the
Desktop.
Launcher names map to data dirs deterministically ( "Big Client" →
~/ClaudeBigClient ), so create and clean always agree on paths.
Labels are intentionally conservative: letters, numbers, spaces, hyphens,
underscores, and apostrophes are supported. Path separators, control characters,
leading dots, overly long labels, and duplicate labels that map to the same
profile directory are rejected.
Sanity checks — confirms macOS and that required tools exist.
Locate Claude — checks standard paths, then Launch Services, then
Spotlight. Exits with a clear message (and the download link) if not found.
Copy Claude's normal icon — launchers use the standard Claude icon if it
is available. No tinting or icon conversion is attempted.
Compile launchers — osacompile turns a one-line open -n ... --user-data-dir
command into a real .app bundle.
Refresh Finder metadata — touch es the bundle so
Finder picks up the change.
-n tells macOS to launch a new instance , which is what allows several
Claude profiles to run simultaneously.
Important caveat: profile isolation depends on Claude Desktop continuing to
honor Electron/Chromium's --user-data-dir flag. That works today for this use
case, but Anthropic could change Claude Desktop in a future release.
Dock icon caveat: the launchers use Claude's normal icon, and running profiles
may appear as separate but same-looking Claude icons in the Dock. This is a
macOS limitation of launching the same underlying /Applications/Claude.app
process with different profile data directories.
create never deletes profile data. It only removes and rebuilds the
.app launcher bundle; your ~/Claude* directories are preserved.
clean only removes launchers this script generated. It identifies them
with a claude-fix-generated marker file inside the app bundle and never
touches the real Claude.app .
clean --purge asks before each deletion , defaulting to No , so an
accidental Enter keeps your data. Purge only targets generated profile
directories such as ~/ClaudePersonal ; it never deletes your normal Claude
app or its default profile data.
No dependencies or package installs. The script uses only built-in macOS
tooling and never installs Python packages.
It does not modify, patch, or re-sign Claude.app .
It does not bypass account limits, authentication, or company policy.
It does not migrate or merge chat history between profiles.
It does not encrypt Claude profile data. Profiles are stored as normal
local Claude/Electron user data.
It does not guarantee official Anthropic support for --user-data-dir .
It does not create different-looking running Dock icons. Profiles may
appear as separate but id

[truncated]

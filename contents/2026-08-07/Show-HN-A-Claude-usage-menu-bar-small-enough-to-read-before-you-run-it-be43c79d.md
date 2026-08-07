---
source: "https://github.com/tpatrouillat/tokease"
hn_url: "https://news.ycombinator.com/item?id=49210250"
title: "Show HN: A Claude usage menu bar small enough to read before you run it"
article_title: "GitHub - tpatrouillat/tokease: macOS menu bar app showing your remaining Claude 5-hour and weekly usage limits. No token read, 100% local. · GitHub"
author: "tpatrouillat"
captured_at: "2026-08-07T14:06:09Z"
capture_tool: "hn-digest"
hn_id: 49210250
score: 1
comments: 0
posted_at: "2026-08-07T13:40:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Claude usage menu bar small enough to read before you run it

- HN: [49210250](https://news.ycombinator.com/item?id=49210250)
- Source: [github.com](https://github.com/tpatrouillat/tokease)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T13:40:09Z

## Translation

タイトル: HN の表示: クロードの使用法メニュー バーは、実行前に読めるほど小さいです
記事のタイトル: GitHub - tpatrouillat/tokeease: クロードの残りの 5 時間および週ごとの使用制限を表示する macOS メニュー バー アプリ。トークン読み取りなし、100% ローカル。 · GitHub
説明: クロードの残りの 5 時間および週ごとの使用制限を表示する macOS メニュー バー アプリ。トークン読み取りなし、100% ローカル。 - tpatrouillat/トケアス

記事本文:
GitHub - tpatrouillat/tokeease: クロードの残りの 5 時間および週ごとの使用制限を表示する macOS メニュー バー アプリ。トークン読み取りなし、100% ローカル。 · GitHub
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
パトゥルイラ
/
トケアス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
62 コミット 62 コミット .github .github 資産 アセット ドキュメント ドキュメント ステータスライン ステータスライン テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md LICE

NSE ライセンス PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md build.sh build.sh install.sh install.sh pyproject.toml pyproject.toml要件.txt要件.txt setup.py setup.py sonar-project.properties sonar-project.properties tracker.py tracker.py uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードの 5 時間および 1 週間の制限をリアルタイムで表示する軽量の macOS メニュー バー アプリ
クロード リミット トラッカーは、トークンやキーチェーンには決して触れません。公式の Claude アプリが Mac 上ですでに公開しているもののみを読み取ります。
醸造インストール tpatrouillat/tap/tokeease
brew services start tokease # 今すぐ開始し、ログインするたびに開始します
macOS 12 Monterey 以降と Claude Pro または Max プランが必要です。 Claude デスクトップ アプリが実行されている場合は、設定は不要です。その他のインストール パスは以下のとおりです。
Tokease が読み取るのは、公式の Claude アプリがすでに Mac に書き込んでいるローカル ファイルの 2 つだけです。それは、Claude デスクトップ アプリのクォータ履歴と、Claude Code がステータスラインに公開する rate_limits データです。キーチェーンからの OAuth トークンの読み取り、非表示の API 呼び出し、ユーザー エージェントのスプーフィングはありません。他のトラッカーはキーチェーンからサブスクリプション トークンを読み取ります。トケアスは決して触れません。
約束によるものではなく、構築によりトークンフリーになります。唯一のデータ ソースは、正式なクロード クライアントが独自の使用のためにローカルに書き込むファイルです。トークンの読み取りやエンドポイントの呼び出しはありません。
使用した履歴ではなく、残りの制限を表示します。 5 時間および週ごとの残り容量とリセットのカウントダウン。 (80% と 95% のしきい値通知はコード内に存在しますが、macOS は署名付き .app バンドルに対してのみ通知を配信するため、Homebrew またはソースのインストールでは通知されません。ロードマップを参照してください。)
ファイル 1 つ、テレメトリなし、アカウントなし。 MIT ライセンスの小さな単一ファイルのメニュー バー アプリ。完全にあなたのマッハ上で動作します

いね。テレメトリもアカウントもサインアップする必要もありません。
2 つの使用状況リング: 5 時間のセッション (外側) と週ごとの (内側) ゲージ、メニュー バー内にあります。
カウントダウンをリセット: 各ウィンドウがいつロールオーバーするかを確認します
正直な鮮度 : クロード クライアントが最後にデータ (デスクトップ アプリまたはクロード コード) を更新した時期を示し、何も実行されていない場合は古いというフラグを立てます。
カスタマイズ可能な表示: アイコン + パーセンテージ、アイコンのみ、またはパーセンテージのみ、および両方のパーセンテージを表示するオプション (5 時間/週)
設定メニュー: 表示モード、更新間隔、アラートしきい値 (80% / 95%、署名された .app ビルドのみ)、ログイン時の起動 (.app ビルド、それ以外の場合は brew サービスまたは LaunchAgent を使用)
軽量: 純粋な Python、2 つの小さな依存関係 (rump + Pillow)
テレメトリなし : アカウントなし、追跡なし、電話は何もありません
macOS (rumps を使用したメニュー バー アプリ)
Python 3.10+ (Homebrew のインストールでは独自のものが使用されます。これはソース インストールの場合にのみ重要です。注: Apple の組み込み python3 は 3.9 である可能性がありますが、これは古すぎます)
Claude Desktop が実行中 (ゼロ構成ソース)、および/または Claude Code ≥ 2.1.x でステータスラインが接続されている (リセット カウントダウンを追加)
クロード プロまたはマックスのサブスクリプション。クォータ フィードはこれらのプランにのみ存在します。Claude Free はクォータ フィードを公開せず、チーム/エンタープライズはクレジット ベースの課金を使用するため、リングはマッピングされず、サポートされません。
Tokease は 2 つのローカルの読み取り専用ソースをマージします。より新鮮なものが勝ちます:
Claude デスクトップ履歴 (ゼロ構成)。 Claude デスクトップ アプリの実行中、アカウント クォータを ~5 分ごとにサンプリングしてローカル ファイル (~/Library/Application Support/Claude/plan-usage-history.json) に保存します。トケアスはそれをそのまま読みます。セットアップなし: デスクトップ アプリが実行されている場合、使用しているサーフェス (VS Code 拡張機能、claude.ai、Cowork、CLI) に関係なく、リングは新鮮なままです。 ADR 0003 を参照してください。
クロード コードのステータスライン (リセット カウントダウンを追加)。小さなキャプチャ

re スクリプト ( statusline/tokease-statusline.py ) は、クロード コードの statusline コマンドとして実行されます。クロード コードは標準入力で rate_limits.five_hour / .seven_day を渡し、スクリプトはそれらを ~/.tokease/usage.json に書き込みます。これは、各リングの横に示されているリセット時間を伝える唯一のフィードです。
どちらの場合も、データは公式のクロード クライアントによって独自に使用するためにローカルに書き込まれます。 Tokease はトークンを読み取ることはなく、エンドポイントを呼び出すこともありません。ステータスラインのセットアップ: statusline/README.md 。
レガシー エンドポイント モード (OAuth トークンを読み取る) は git タグ v0.9.0-endpoint で凍結されます。 v1.0 はトークンを読み取りません。
これは独立したプロジェクトであり、Anthropic と提携または承認されていません。
クイックインストール (Homebrew、推奨)
醸造インストール tpatrouillat/tap/tokeease
tokease # すぐに起動
brew services start tokease # ログイン時に自動開始
Gatekeeper をバイパスし (Apple Developer の署名は必要ありません)、Python virtualenv を自動的にインストールし、brew upgrade が更新を処理します。
git clone https://github.com/tpatrouillat/tokease.git
CDトークケース
bash インストール.sh
インストール スクリプトは次のことを行います。
Python 仮想環境を作成し、依存関係をインストールする
オプションで、ログイン時に自動起動するように macOS LaunchAgent を設定します。
必要に応じてアプリをすぐに起動します
手動インストール (インストール スクリプトなし)
git clone https://github.com/tpatrouillat/tokease.git
CDトークケース
python3 -m venv venv
venv/bin/pip install -rrequirements.txt
venv/bin/python tracker.py
Claude デスクトップ アプリが実行されていれば、作業は完了です。必要に応じて、リセット カウントダウン用に statusline キャプチャ スクリプトを接続します。 statusline/README.md を参照してください。
ソースから: bash uninstall.sh は、LaunchAgent、~/.tokease/ データ、および Tokease statusLine ブロックを ~/.claude/settings.json (バックアップ付き) から削除します。次に、クローンしたフォルダーを削除します。
Homebrew: 同じクリーンアップを実行します。

st を削除し、次の式を削除します。
bash " $( brew --prefix ) /opt/tokease/libexec/uninstall.sh "
brew サービス stop tokease && brew uninstall tokease
brew uninstall だけではアプリは削除されますが、~/.tokease/ とステータスラインの配線はそのまま残ります。
クロードのアップデート後にリングが動かなくなった場合
Claude デスクトップ アプリのクォータ ファイルは内部にあり文書化されていないため、Claude の更新により予告なく形式が変更される可能性があります。 Tokease はそれを検出し、間違った数値を表示するのではなく、ステータスライン フィードにフォールバックします。ただし、どちらのソースも更新しない場合、リングは最後の読み取り値でフリーズします (常に古いフラグが立てられます)。 2 つのコマンドにより、どのソースが壊れたかがわかります。
ls -la ~ /Library/Application \ Support/Claude/plan-usage-history.json # デスクトップ ソースは存在しますか?
cat ~ /.tokease/statusline.err # キャプチャ スクリプト エラー
見つけた内容について問題を開いてください。テレメトリーがないため、これを認識するにはレポートが唯一の方法です。
Tokease はおそらく正常に実行されており、メニュー バーはいっぱいになっています。ノッチ付き MacBook では、macOS はオーバーフロー メニュー バー アイコンをインジケーターなしで非表示にします。不要なメニュー バー アプリを終了すると (または Ice や Bartender などのマネージャーを使用すると)、リングが表示されます。 [設定] のパーセンテージのみの表示モードも Tokease の表示領域を狭くし、Cmd キーを押しながらアイコンを時計の方向にドラッグするとアイコンが表示されたままになります。
割合はアカウント レベルです。claude.ai チャットや Claude Desktop (Cowork を含む) から VS Code 拡張機能や CLI まで、サブスクリプションのすべてが対象になります。リングは、使用したクォータの量を決して間違えません。唯一の問題は、最後に読んだ内容がどれだけ新鮮かということです。
クロード デスクトップの実行: どのような表面で作業していても、約 5 分ごとに読み取ります。これは推奨されるセットアップです。デスクトップ アプリを開いたままにしておきます (とにかくメニュー バーにあります)。
ステータスラインのみが接続されています: 読み取り値が更新されます。

対話型のクロード端末セッションがアクティブです (CLI または VS Code 統合端末)。 VS Code 拡張パネル、Claude Desktop およびヘッドレス claude -p はステータスラインを実行しません ( #55643 、「計画されていません」で閉じられています)。ステータスラインを唯一のソースとして使用すると、リングは CLI セッション間で古くなります。
両方 (最良): デスクトップ履歴によりリングが最新の状態に保たれ、ステータスライン キャプチャにより CLI を使用するたびにリセット カウントダウンが追加されます。
古いデータには常に視覚的にフラグが立てられ、すでにロールオーバーしたリセット ウィンドウが検出されます。古いパーセンテージが新しいものであるかのように表示されることはありません。
ステータスライン フィードのみがリセット時間を伝えます。デスクトップ履歴ファイルにはパーセンテージが含まれていますが、リセット タイムスタンプが含まれていないため、2 つのプロファイルは異なります。
パワー ユーザー (あらゆるクロード端末を、場合によっては使用する)。インタラクティブ CLI セッション内のすべてのメッセージは、両方のリセット時間をキャプチャします。その後、リセットが経過するまで (最大 7 日間)、毎週のカウントダウンが表示され続けます。 5 時間のカウントダウンはそのウィンドウの長さだけ存続するため、表示するには現在の 5 時間以内にキャプチャが必要です。
デスクトップのみのユーザー。呼び出し音とパーセンテージは最新のままですが、メニューにはリセットが表示されます -- 。これは予期されたものであり、バグではありません。任意のクロード端末セッション (VS Code 統合端末カウント) からの 1 つのメッセージにより、両方のカウントダウンが戻されます。
リセット時間が経過すると、Tokease は次の時間を推測するのではなく表示します。 The weekly cycle could be extrapolated, but the 5-hour window anchors on your first message after the previous reset, and showing a made-up time would break the honesty rule above.
デスクトップ履歴ファイルに関する 2 つの注意事項。その形式は Claude アプリの内部にあり文書化されていないため、将来のアップデートによって変更される可能性があります。 Tokease はそれを防御的に解析し (異常があればステータスライン フィードにフォールバックし)、その期待値を観察されたバージョンに固定します。

シオン: 2 。 And older Claude Desktop builds may not write this file at all: if plan-usage-history.json doesn't exist on your machine, update the desktop app or wire the statusline.
これをよりクリーンにする上流の機能リクエスト (いずれかの機能が提供されると、Tokease の利点が自動的に得られます):
anthropics/claude-code#38380 : CLI フラグまたはフック イベントを介して使用量/レート制限データを公開します
anthropics/claude-code#55643 : VS Code 拡張機能でのステータスラインのサポート
anthropics/claude-code#33257 : ネイティブ使用状況インジケーター
ローカルのオプトイン ドリフト推定器 (すべてのサーフェスに書き込まれるローカル セッション ログからキャプチャ間の使用量を推定する) は、v1.1 の候補です。 ROADMAP.md を参照してください。
トークンをまったく読みません。 Tokease only reads two local files: ~/.tokease/usage.json (written by the Claude Code statusline) and the Claude desktop app's own plan-usage-history.json (read-only).キーチェーンや OAuth トークンには決して触れません。
独自の API 呼び出しはありません。関係する唯一のネットワーク呼び出しは、正式なクロード クライアントがすでに行っているものです。 Tokease は、ローカルで書き込まれた結果を読み取るだけです。
データ収集はありません。アプリは完全にマシン上で実行されます。
シークレットは保存されません。ディスク上に .env 、トークン、資格情報はありません。
オープンソース。コードを自分で監査してください。

[切り捨てられた]

## Original Extract

macOS menu bar app showing your remaining Claude 5-hour and weekly usage limits. No token read, 100% local. - tpatrouillat/tokease

GitHub - tpatrouillat/tokease: macOS menu bar app showing your remaining Claude 5-hour and weekly usage limits. No token read, 100% local. · GitHub
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
tpatrouillat
/
tokease
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
62 Commits 62 Commits .github .github assets assets docs docs statusline statusline tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md build.sh build.sh install.sh install.sh pyproject.toml pyproject.toml requirements.txt requirements.txt setup.py setup.py sonar-project.properties sonar-project.properties tracker.py tracker.py uninstall.sh uninstall.sh View all files Repository files navigation
A lightweight macOS menu bar app showing your Claude 5-hour and weekly limits in real time
The Claude limit tracker that never touches your token or Keychain. It reads only what official Claude apps already publish on your Mac.
brew install tpatrouillat/tap/tokease
brew services start tokease # starts it now, and at every login
Requires macOS 12 Monterey or later and a Claude Pro or Max plan. Zero config if the Claude desktop app is running. Other install paths below.
Tokease reads only two local files that official Claude apps already write on your Mac: the quota history of the Claude desktop app, and the rate_limits data Claude Code publishes to its statusline. No OAuth token read from the Keychain, no hidden API call, no User-Agent spoofing. Other trackers read your subscription token from the Keychain. Tokease never touches it.
Token-free by construction, not by promise. The only data sources are files official Claude clients write locally for their own use. No token read, no endpoint call.
Shows the limit you have left, not the history you spent. Your 5-hour and weekly remaining capacity, with reset countdowns. (Threshold notifications at 80% and 95% exist in the code but macOS only delivers them for a signed .app bundle, so they do not fire on Homebrew or source installs — see the roadmap.)
One file, zero telemetry, zero account. A small single-file menu bar app, MIT-licensed. It runs entirely on your machine. No telemetry, no account, nothing to sign up for.
Two usage rings : 5-hour session (outer) and weekly (inner) gauges, right in the menu bar
Reset countdowns : see when each window rolls over
Honest freshness : shows when a Claude client last refreshed the data (desktop app or Claude Code), and flags it stale when none is running
Customizable display : icon + percentage, icon only, or percentage only, plus an option to show both percentages ( 5h / weekly )
Settings menu : display modes, refresh interval, alert thresholds (80% / 95%, signed .app build only), launch at login ( .app build; use brew services or the LaunchAgent otherwise)
Lightweight : pure Python, two small dependencies ( rumps + Pillow )
No telemetry : no account, no tracking, nothing phones home
macOS (menu bar app using rumps )
Python 3.10+ (Homebrew installs bring their own, this only matters for source installs. Note: Apple's built-in python3 can be 3.9, which is too old)
Claude Desktop running (zero-config source), and/or Claude Code ≥ 2.1.x with the statusline wired up (adds reset countdowns)
Claude Pro or Max subscription. The quota feeds only exist for these plans: Claude Free doesn't expose them, and Team/Enterprise use credit-based billing, so the rings don't map and it isn't supported.
Tokease merges two local, read-only sources. Whichever is fresher wins:
Claude Desktop history (zero config). While the Claude desktop app runs, it samples your account quota every ~5 minutes into a local file ( ~/Library/Application Support/Claude/plan-usage-history.json ). Tokease reads it as-is. No setup: if the desktop app is running, the rings stay fresh whatever surface you're using (VS Code extension, claude.ai, Cowork, CLI). See ADR 0003 .
Claude Code statusline (adds reset countdowns). A small capture script ( statusline/tokease-statusline.py ) runs as your Claude Code statusline command. Claude Code passes it rate_limits.five_hour / .seven_day on stdin and the script writes them to ~/.tokease/usage.json . This is the only feed carrying the reset times shown next to each ring.
In both cases the data is written locally by an official Claude client for its own use. Tokease never reads your token and never calls any endpoint. Statusline setup: statusline/README.md .
The legacy endpoint mode (which read the OAuth token) is frozen at the git tag v0.9.0-endpoint . v1.0 never reads the token.
This is an independent project, not affiliated with or endorsed by Anthropic .
Quick Install (Homebrew, recommended)
brew install tpatrouillat/tap/tokease
tokease # launch immediately
brew services start tokease # auto-start at login
Bypasses Gatekeeper (no Apple Developer signing required), installs the Python virtualenv automatically, and brew upgrade handles updates.
git clone https://github.com/tpatrouillat/tokease.git
cd tokease
bash install.sh
The install script will:
Create a Python virtual environment and install dependencies
Optionally set up a macOS LaunchAgent for auto-start at login
Optionally launch the app immediately
Manual Install (no install script)
git clone https://github.com/tpatrouillat/tokease.git
cd tokease
python3 -m venv venv
venv/bin/pip install -r requirements.txt
venv/bin/python tracker.py
If the Claude desktop app is running, you are done. Optionally wire up the statusline capture script for reset countdowns: see statusline/README.md .
From source: bash uninstall.sh removes the LaunchAgent, the ~/.tokease/ data, and the Tokease statusLine block from ~/.claude/settings.json (with a backup). Then delete the cloned folder.
Homebrew: run the same cleanup first, then remove the formula:
bash " $( brew --prefix ) /opt/tokease/libexec/uninstall.sh "
brew services stop tokease && brew uninstall tokease
brew uninstall alone removes the app but leaves ~/.tokease/ and any statusline wiring in place.
If the rings stop moving after a Claude update
The Claude desktop app's quota file is internal and undocumented, so a Claude update can change its format without notice. Tokease detects that and falls back to the statusline feed rather than showing a wrong number, but if neither source refreshes, the rings freeze at their last reading (always flagged stale). Two commands tell you which source broke:
ls -la ~ /Library/Application \ Support/Claude/plan-usage-history.json # desktop source present?
cat ~ /.tokease/statusline.err # capture script errors
Please open an issue with what you find. There is no telemetry, so a report is the only way this gets noticed.
Tokease is probably running fine and your menu bar is full. On notched MacBooks, macOS hides overflow menu bar icons without any indicator. Quit a menu bar app you don't need (or use a manager like Ice or Bartender) and the rings appear. The percentage-only display mode in Settings also narrows Tokease's footprint, and Cmd-dragging the icon toward the clock keeps it visible.
The percentages are account-level : they cover everything on your subscription, from claude.ai chat and Claude Desktop (Cowork included) to the VS Code extension and the CLI. The ring is never wrong about how much quota you've used. The only question is how fresh the last reading is.
Claude Desktop running : readings every ~5 minutes, whatever surface you work in. This is the recommended setup. Just keep the desktop app open (it lives in your menu bar anyway).
Only the statusline wired : readings refresh while an interactive claude terminal session is active (the CLI, or the VS Code integrated terminal ). The VS Code extension panel, Claude Desktop and headless claude -p never execute statuslines ( #55643 , closed "not planned"). With the statusline as sole source, the rings go stale between CLI sessions.
Both (best): desktop history keeps the rings fresh, and statusline captures add the reset countdowns whenever you use the CLI.
Stale data is always visibly flagged, and reset windows that have already rolled over are detected. An old percentage is never shown as if it were fresh.
Only the statusline feed carries reset times. The desktop history file has percentages but no reset timestamps, so the two profiles differ:
Power users (any claude terminal use, even occasional). Every message in an interactive CLI session captures both reset times. The weekly countdown then stays displayed until that reset passes (up to 7 days). The 5-hour countdown only lives as long as its window, so it needs a capture within the current 5 hours to show up.
Desktop-only users. Rings and percentages stay current, but the menu shows resets -- . That is expected, not a bug. One message from any claude terminal session (the VS Code integrated terminal counts) brings both countdowns back.
When a reset time has passed, Tokease shows -- rather than guessing the next one. The weekly cycle could be extrapolated, but the 5-hour window anchors on your first message after the previous reset, and showing a made-up time would break the honesty rule above.
Two caveats on the desktop history file. Its format is internal to the Claude app and undocumented, so a future update could change it. Tokease parses it defensively (any anomaly falls back to the statusline feed) and pins its expectations to the observed version: 2 . And older Claude Desktop builds may not write this file at all: if plan-usage-history.json doesn't exist on your machine, update the desktop app or wire the statusline.
Upstream feature requests that would make this cleaner (Tokease benefits automatically if any of them ships):
anthropics/claude-code#38380 : expose usage/rate-limit data via a CLI flag or hook event
anthropics/claude-code#55643 : statusline support in the VS Code extension
anthropics/claude-code#33257 : native usage indicator
A local, opt-in drift estimator (estimate usage between captures from the local session logs all surfaces write) is a v1.1 candidate. See ROADMAP.md .
Reads no token, ever. Tokease only reads two local files: ~/.tokease/usage.json (written by the Claude Code statusline) and the Claude desktop app's own plan-usage-history.json (read-only). It never touches the Keychain or your OAuth token.
No API calls of its own. The only network calls involved are the ones official Claude clients already make. Tokease just reads the results they wrote locally.
No data collection. The app runs entirely on your machine.
No secrets stored. No .env , no tokens, no credentials on disk.
Open source. Audit the code yourself:

[truncated]

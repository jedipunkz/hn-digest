---
source: "https://github.com/silverdolphin863/claude-code-meter"
hn_url: "https://news.ycombinator.com/item?id=49144960"
title: "Show HN: CC Meter, a Windows Tray Gauge for Claude Code and Codex Limits"
article_title: "GitHub - silverdolphin863/claude-code-meter: Tray widget showing Claude Code and Codex rate-limit usage and burn pace. Windows, local-only, MIT. · GitHub"
author: "silversurfer863"
captured_at: "2026-08-02T14:27:52Z"
capture_tool: "hn-digest"
hn_id: 49144960
score: 1
comments: 0
posted_at: "2026-08-02T14:19:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CC Meter, a Windows Tray Gauge for Claude Code and Codex Limits

- HN: [49144960](https://news.ycombinator.com/item?id=49144960)
- Source: [github.com](https://github.com/silverdolphin863/claude-code-meter)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T14:19:25Z

## Translation

タイトル: Show HN: CC Meter、クロード コードおよびコーデックス制限用の Windows トレイ ゲージ
記事のタイトル: GitHub - silverdolphin863/claude-code-meter: クロード コードとコーデックスのレート制限の使用量と書き込みペースを示すトレイ ウィジェット。 Windows、ローカルのみ、MIT。 · GitHub
説明: クロード コードとコーデックスのレート制限の使用量と書き込みペースを示すトレイ ウィジェット。 Windows、ローカルのみ、MIT。 - silverdolphin863/クロードコードメーター

記事本文:
GitHub - silverdolphin863/claude-code-meter: クロード コードとコーデックスのレート制限の使用量と書き込みペースを示すトレイ ウィジェット。 Windows、ローカルのみ、MIT。 · GitHub
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
シルバードルフィン863
/
クロードコードメーター
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット アセット アセット build build docs docs public public .gitignore .gitignore ライセンス ライセンス README.md README.md main.mjs main.mjs package-lock.json package-lock.json package.json package.json preload.cjs preload.cjs server.mjs server.mjs start-widget.cmd start-widget.cmd test.mjs test.mjs すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードとコーデックスの量を表示する Windows 用トレイ ウィジェット
焼き付けたレート制限と、それを焼き付ける速度。
ネイティブ解像度の v1.0.4 フルパネルでレンダリングされた代表的な使用法。
コストトラッカーではありません。それはセッション中に重要な唯一の質問への答えです: 私はですか?
このウィンドウがリセットされる前に壁にぶつかるでしょうか?
各リミット ウィンドウ: 使用率、バー、リセット時間、ペース バッジ。
ペースは便利な部分です。 1.0x は、クォータを正確に同期して消費していることを意味します。
時計なので、リセット時に正確に使い果たされます。 0.3x は、あなたがそうであることを意味します
惰性走行。 1.6倍ということは、早く壁にぶつかることを意味しており、おおよそどのくらい早く壁にぶつかるかということです。
クロード コード - 5 時間ウィンドウ、毎週の全モデル、および毎週のモデルごと
(作品、寓話、プランの範囲に関係なく)。モデルごとの分割は実際に行われます。
/usage 画面が使用するのと同じソースから来ています。
Codex - Codex が現在報告しているアカウント全体のウィンドウ。 1 つを共有
プール。一部のプランでは週ごとのウィンドウのみが公開され、CC Meter は決して
1つ欠けています。
2 つのモード: 画面の上端にドッキングするコンパクトなストリップ (
オプションで端に触れるまで自動的に非表示になります）、およびカウントダウン付きのフルパネル。
コンパクトモードは上端で邪魔になりません。コンパクトストリップを開きます
ネイティブのフル解像度で。
ダウンロード

CCMeter-Setup-x.y.z.exe から
リリースして実行します。タスクバーではなくトレイに存在します。
インストーラーは署名されていないため、Windows SmartScreen に次のように表示されます。
「Windows が PC を保護しました」。 [詳細情報] -> [とにかく実行] をクリックします。もしそうしていただければ
インターネット上の見知らぬ人からのバイナリを信頼せず、自分でビルドしてください
以下ソースより。それが正直な答えであり、それがソースが
ここです。
CC Meter にはホスト型サービスはありません。 Anthropic社にクロードの使用申請を行う
独自の OAuth トークンを使用してマシンから。
Codex の場合、インストールされている Codex CLI に現在のアカウントのスナップショットを要求します。
Codex は独自の認証された OpenAI リクエストを処理します。 CC メーターは決して読み取りません
コーデックス認証トークン。
Claude : 標準の Claude Code OAuth トークンを読み取ります。
~/.claude/.credentials.json と呼び出し
ローカル キャッシュがある場合は https://api.anthropic.com/api/oauth/usage を直接使用します。
10分以上経っています。への応答を書き込みます
~/.claude/usage-cache.json とキャッシュ ロックを尊重します。ステータス行スクリプトはありません
が必要です。
Codex : ローカルの Codex アプリサーバーを起動し、呼び出します
account/rateLimits/read 、5 分間キャッシュされます。これは現在を返します
このマシンのセッション ファイルが古い場合でも、アカウント全体のウィンドウを表示します。もし
古い Codex バージョンではそのメソッドが公開されていないため、CC Meter は最新のバージョンにフォールバックします
~/.codex/sessions/ の下にあるファイルを削除し、リセット時間が経過したすべてのウィンドウを破棄します。
すでに過ぎました。
CC Meter はクロード OAuth トークンをディスクから読み取り、次の宛先へ送信します。
api.anthropic.com 以外にはありません。ログに記録されることも、どこにも保存されることもありません
新しいものであり、Anthropic 以外のホストには送信されません。直属のクロード
ネットワークパスは、server.mjsのrefreshClaudeUsage()で分離されます。コーデックス
認証およびアカウント要求は、インストールされた Codex CLI 内に残ります。
使用エンドポイントは文書化されておらず、レートが制限されています（次の点に注意してください）。

d: HTTP 429 と
~57 分の再試行後)。 CC メーターは最大 10 分に 1 回更新されます。
Retry-After を正確に尊重し、失敗すると指数関数的に後退し、永続化します。
このバックオフは再起動時に行われるため、再起動によってロックアウトされたエンドポイントが破壊されることはありません。
文書化されていないため、エンドポイントは予告なく変更されたり、動作を停止したりする場合があります。
UI は http://localhost:7373 から提供され、GET /usage.json は
ローカル使用状況データ。ブラウザのクロスオリジン アクセスはデフォルトでは無効になっています。もしあなたが
別のディスプレイを駆動したい場合は、CCMETER_CORS_ORIGIN をそのディスプレイの値に設定します。
サーバーを起動する前に特定のオリジンを確認してください。トークンやセッションのコンテンツはありません。
このエンドポイントによって返されます。
npmインストール
npm テスト # 分離されたクロードおよびコーデックス フィクスチャ テスト
npm 開始 # 開発
npm run dist # 未署名の NSIS インストーラー -> dist/
ノード 20 以降と Windows が必要です。 macOS と Linux はテストされていません。トレイの動作
画面端のドックは Windows の形をしています。
OAuth (サブスクリプション) クロード コード ログインでのみ機能します。ユーザーが持っている API キー
表示するプラン制限はありません。
ライブ Codex の読み取りには、以下を公開する Codex CLI バージョンが必要です。
アカウント/rateLimits/読み取り 。古いバージョンではフィルタリングされたセッションログが使用されます
Codex がセッション ファイルを書き込む場合にのみフォールバックして更新します。
自動更新はありません。リリースページをご覧ください。
Anthropic または OpenAI と提携、承認、またはサポートされていません。
クロード コードとコーデックスのレート制限の使用量と書き込みペースを示すトレイ ウィジェット。 Windows、ローカルのみ、MIT。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Tray widget showing Claude Code and Codex rate-limit usage and burn pace. Windows, local-only, MIT. - silverdolphin863/claude-code-meter

GitHub - silverdolphin863/claude-code-meter: Tray widget showing Claude Code and Codex rate-limit usage and burn pace. Windows, local-only, MIT. · GitHub
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
silverdolphin863
/
claude-code-meter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits assets assets build build docs docs public public .gitignore .gitignore LICENSE LICENSE README.md README.md main.mjs main.mjs package-lock.json package-lock.json package.json package.json preload.cjs preload.cjs server.mjs server.mjs start-widget.cmd start-widget.cmd test.mjs test.mjs View all files Repository files navigation
A tray widget for Windows that shows how much of your Claude Code and Codex
rate limits you have burned, and how fast you are burning them.
Representative usage rendered by the v1.0.4 full panel at native resolution.
Not a cost tracker. It answers the only question that matters mid-session: am I
going to hit the wall before this window resets?
For each limit window: percent used, a bar, the reset time, and a pace badge.
Pace is the useful part. 1.0x means you are burning quota exactly in step with
the clock, so you will run out precisely at the reset. 0.3x means you are
coasting. 1.6x means you will hit the wall early, and roughly how early.
Claude Code - 5-hour window, weekly all-models, and weekly per-model
(Opus, Fable, whichever your plan scopes). The per-model split is real, it
comes from the same source the /usage screen uses.
Codex - the account-wide windows Codex currently reports. One shared
pool; some plans expose only a weekly window, and CC Meter never invents a
missing one.
Two modes: a compact strip that docks against the top edge of the screen (with
optional auto-hide until you touch the edge), and a full panel with countdowns.
The compact mode stays out of the way at the top edge. Open the compact strip
at its full native resolution .
Download CCMeter-Setup-x.y.z.exe from
Releases and run it. It lives in the tray, not the taskbar.
The installer is unsigned , so Windows SmartScreen will show
"Windows protected your PC". Click More info -> Run anyway . If you would
rather not trust a binary from a stranger on the internet, build it yourself
from the source below. That is the honest answer and it is why the source is
here.
CC Meter has no hosted service. It makes the Claude usage request to Anthropic
from your machine with your own OAuth token.
For Codex, it asks the installed Codex CLI for the current account snapshot;
Codex handles its own authenticated OpenAI request. CC Meter never reads the
Codex authentication token.
Claude : reads the standard Claude Code OAuth token from
~/.claude/.credentials.json and calls
https://api.anthropic.com/api/oauth/usage directly when the local cache is
older than ten minutes. It writes the response to
~/.claude/usage-cache.json and honours the cache lock. No status-line script
is required.
Codex : launches the local Codex app-server and calls
account/rateLimits/read , cached for five minutes. This returns the current
account-wide windows even when this machine's session files are old. If an
older Codex version does not expose that method, CC Meter falls back to recent
files under ~/.codex/sessions/ and discards every window whose reset time has
already passed.
CC Meter reads your Claude OAuth token off disk and sends it to
api.anthropic.com and nowhere else. It is never logged, never stored anywhere
new, and never transmitted to any host but Anthropic's. The direct Claude
network path is isolated in refreshClaudeUsage() in server.mjs . Codex
authentication and account requests remain inside the installed Codex CLI.
The usage endpoint is undocumented and rate-limited (observed: HTTP 429 with a
~57 minute Retry-After ). CC Meter refreshes at most once every 10 minutes,
honours Retry-After exactly, backs off exponentially on failure, and persists
that backoff across restarts so relaunching cannot hammer a locked-out endpoint.
Being undocumented, the endpoint may change or stop working without notice.
The UI is served from http://localhost:7373 , and GET /usage.json returns the
local usage data. Browser cross-origin access is disabled by default. If you
want to drive another display, set CCMETER_CORS_ORIGIN to that display's
specific origin before starting the server. No token or session content is
returned by this endpoint.
npm install
npm test # isolated Claude and Codex fixture test
npm start # dev
npm run dist # unsigned NSIS installer -> dist/
Requires Node 20+ and Windows. macOS and Linux are untested; the tray behaviour
and the screen-edge dock are Windows-shaped.
Only works with an OAuth (subscription) Claude Code login. API-key users have
no plan limits to display.
Live Codex readings require a Codex CLI version that exposes
account/rateLimits/read . Older versions use the filtered session-log
fallback and update only when Codex writes a session file.
No auto-update. Watch the releases page.
Not affiliated with, endorsed by, or supported by Anthropic or OpenAI.
Tray widget showing Claude Code and Codex rate-limit usage and burn pace. Windows, local-only, MIT.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

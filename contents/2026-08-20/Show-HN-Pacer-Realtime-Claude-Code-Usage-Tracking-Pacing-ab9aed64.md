---
source: "https://github.com/EricAndrechek/Pacer"
hn_url: "https://news.ycombinator.com/item?id=49376775"
title: "Show HN: Pacer – Realtime Claude Code Usage Tracking / Pacing"
article_title: "GitHub - EricAndrechek/Pacer: Native macOS app for tracking Claude Code usage — tokens, cost, rate-limit pacing, per-project breakdowns. SwiftUI + SwiftData. · GitHub"
image: "https://opengraph.githubassets.com/49c8d0d017b6a2251743da0f5666c16eebb3718d6e7ecf3f8a08260022a17e6e/EricAndrechek/Pacer"
author: "jwoodsworks"
captured_at: "2026-08-20T17:21:10Z"
capture_tool: "hn-digest"
hn_id: 49376775
score: 2
comments: 1
posted_at: "2026-08-20T16:23:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pacer – Realtime Claude Code Usage Tracking / Pacing

- HN: [49376775](https://news.ycombinator.com/item?id=49376775)
- Source: [github.com](https://github.com/EricAndrechek/Pacer)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T16:23:20Z

## Translation

タイトル: Show HN: Pacer – リアルタイム クロード コード使用状況追跡 / ペーシング
記事のタイトル: GitHub - EricAndrechek/Pacer: クロード コードの使用状況を追跡するためのネイティブ macOS アプリ — トークン、コスト、レート制限ペーシング、プロジェクトごとの内訳。 SwiftUI + SwiftData。 · GitHub
説明: クロード コードの使用状況 (トークン、コスト、レート制限ペーシング、プロジェクトごとの内訳) を追跡するためのネイティブ macOS アプリ。 SwiftUI + SwiftData。 - エリック・アンドレチェック/ペーサー

記事本文:
GitHub - EricAndrechek/Pacer: クロード コードの使用状況 (トークン、コスト、レート制限ペーシング、プロジェクトごとの内訳) を追跡するためのネイティブ macOS アプリ。 SwiftUI + SwiftData。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
エリック・アンドレチェク
/
ペーサー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
353 コミット 353 コミット フォルダーとファイル
.claude/ スキル/ シップ .claude/ スキル/ シップ .github .github アプリ アプリ PacerCore PacerCore ベンダー ベンダー Wi

dgets ウィジェット bin bin docs docs Research/ harness Research/ harness .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md project.yml project.yml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Mac のメニュー バーから直接、Claude Code のコストと、限界にどれだけ近づいているかを知ることができます。
Pacer は、クロード コードを監視する無料のオープンソース macOS アプリです。
使用法: 燃やしているトークン、そのコスト、どれくらい近いのか
5 時間および週ごとの料金制限、および支出の行き先 - プロジェクト別、別
日ごとのモデル。メニューバーにひっそりと鎮座し、常に最新の状態に保ちます。
すべてのデータを自分の Mac に保存します。
メニューバーとクリックダウンポップオーバー
ホーム画面と通知センターのウィジェット
5 時間および 7 日間の燃焼チャートを画像として共有します
モデル — シェア、トレンド、モデルごとのテーブル。モデルまたはクラスごとにグループ化する
プロジェクトとコレクション — プロジェクトをグループ化し、合計をコレクションにスコープします
履歴 — 生涯合計、6 か月のヒートマップ、毎月の支出
注意してください — 初期の頃。 Pacer は 1.0 より前であり、現在開発中です。それ
動作し、今日でも役に立ちますが、時折荒い点があることを覚悟してください。
保存された履歴は、0.x バージョン間で再構築する必要がある場合があります。
最新バージョンをダウンロード →
(「Assets」の下にある Pacer-x.y.z.dmg ファイルを取得します)。
ダウンロードしたファイルを開き、Pacer をアプリケーション フォルダーにドラッグします。
ペーサーを起動します。メニューバーにあります - 小さなゲージアイコンを探してください
一番上にあります (Dock アイコンはありません)。クリックしてダッシュボードを開きます。
初めて、macOS は Pacer にクロード コードのファイルを読み取る許可を求めます。
[許可] をクリックします。これにより、Pacer はあなたの使用状況を確認します。 Mac からは何も残りません。
それだけです。 Pacer は、新しいバージョンになると、自動的に最新の状態に保たれます。
船、それは提供します

自動的にインストールします。再ダウンロードや再インストールは必要ありません。
Apple Silicon または Intel の macOS 15 (Sequoia) 以降が必要です。
レート制限ペーシング。 5 時間、毎週、およびモデルごとのウィンドウ —
毎週の寓話のような、アカウントの対象範囲を限定した Anthropic レポート
制限 — 読みやすいペースチャートとして。進んでいますか、軌道に乗っていますか、それともこれから到着しますか
壁は？カラーバンド (後方 / 軌道上 / 前方 / 最大近く) が一目でわかります。
一目見ると、すべてのウィンドウで同じ予測が得られます。投影された塗りつぶし、
制限時間と校正されたバンド。から数分ごとに更新されます
Anthropic の使用状況データ。
費用は思いのまま。 Claude Code が報告する方法での支出を参照するか、Pacer に問い合わせてください。
トークン自体から価格を設定します。設定で切り替え可能です。毎日、毎月、
歴代の合計も含まれます。
どこへ行くのか。使用状況をプロジェクトごとに分類 - 独自のプロジェクトにグループ化
非破壊コレクション — モデルごと、すべてのクロード モデルをそのコレクションに含める
独自の色とクラスごとにグループ化できます。任意の 1 日を掘り下げて確認すると、
過去 6 か月の GitHub スタイルのアクティビティ ヒートマップ。
「今日」のライブビュー。現在の燃焼速度とランニング「このペースでは、
今日は約 X ドルで終了すると予想されます。
複数のアカウント。複数の Claude アカウントにサインインしていますか?ペーサートラック
各アカウントの使用量と制限を個別に設定し、それらを切り替えることができます。
一目で、いつも。構成可能なメニューバーの読み出し - アクティビティリング
アイコン、チップのパーセント、およびアイコンを駆動するウィンドウの明示的な選択 —
さらに、ホーム画面スタイルのウィジェットを好みのウィンドウに選択可能
コストとペースについては、約 (5 時間、毎週、またはモデルごとの上限) です。
オプションのナッジ。レート制限のしきい値を超えた場合のローカル通知
(50 / 75 / 90%) — モデルごとのキャップを含む任意のウィンドウ上 — または吹き飛ばし
あなたが設定した1日の支出制限。デフォルトではオフです。

輸出。毎日の合計、毎日のモデル別、またはプロジェクトごとの数値を
独自のスプレッドシートの CSV。
Claude Code の使用状況を監視するためのオプションがあります - Claude Code 独自の
/usage 、一般的な ccusage CLI、および
いくつかのメニューバー アプリ — 最も近いものは Token Pacer です。
有料のクローズドソースの同名タイトルに加えて、Claude God と
ccseva 。それらは良いツールです。ここにあります
正直な土地の様子。
✦ ウィンドウを読み取るのではなく、ローカルの JSONL ログからウィンドウを推定します。
Anthropic の使用 API。 Token Pacer は有料のクローズドソース アプリであるため、その列は
ソースではなく公開サイトから読み取り、その機能の一部 (CSV)
輸出、カスタム モデルの価格設定）は Pro レベルです。 2026 年 6 月時点のベストエフォート — これら
ツールはすべて高速に動作するため、問題または PR を通じて修正を歓迎します。
vs. /usage — Claude Code の組み込みビューは、ユーザーが要求したスナップショットです。
1 つの端末。 Pacer は、常にオンでズームアウトしたコンパニオンであり、すべての情報を記憶します。
セッション。 1 台の車のスピードメーターと、すべての走行を記録するダッシュボード。
vs. ccusage — スクリプト可能な数値に最適な CLI (Pacer のテストでも
スキャナーを照合してください)。 Pacer は、一目でわかる GUI です。
コマンドを再実行すると、実際の制限値が Anthropic から読み取られます。
ログから推定するよりも。
vs. クロード・ゴッド / ccseva — 最も近いライバルで、本当に素晴らしい。ペーサー
ペース配分に傾いている（理想的な燃焼ラインに合わせて、「足りなくなるかな？）」
リセット前?")、ネイティブ + デフォルトで静か、署名/公証済み
自己更新。 Claude God が ROI と git の相関関係とプラグインについてさらに詳しく解説します
マーケットプレイス、およびガラス質の UI 上の ccseva。あなたのことを考えたものを選んでください
あなたがそうするように使用すれば、それらは幸せに共存します。
vs. Token Pacer — 近い名前であり、直接的なものに最も近いもの
ライバル：これもネイティブ、ローカルファーストの macOS 私

nu-bar トラッカー、そしてそれは到達します
より広範囲 — Codex と opencode をクロード コードとともにカバーし、「次のアクション」を表示
ナッジする。トレードオフ: 29 ドルでクローズドソース、Pacer は無料、
開いた; Pacer は Anthropic から実際の限界値 % を読み取り、理想的な燃焼を追加します。
ペースライン、6か月のヒートマップ、ウィジェット。複数のエージェントを使用しており、料金を支払う意思がありますか?
公正な選択。クロード コードファーストでオープンソースを重視しますか?ペーサー。
ペーサーはローカルファーストです。 Claude Code が既に書き込んだファイルのみを読み取ります。
自分の Mac を使用しても、使用状況についてはどこにも送信されません。
~/.claude/projects/*.jsonl — セッション トークンの使用状況。マシン上で読み取ります。
macOS キーチェーン内の Claude Code ログイン トークン — 質問するためにのみ使用されます
クロード・コード自体が行うのと同じように、レート制限ステータスを人為的に扱います。
Pacer が行う唯一のネットワーク接続は次のとおりです。
api.anthropic.com — レート制限ウィンドウを確認します (最大 12 の小さなリクエスト
実行中の時間)。
github.com — アプリのアップデートを確認、ダウンロード、インストールします (起動時、
その後は6時間ごと。自動更新時にインストールがバックグラウンドで行われる
オンになっています）。
分析、テレメトリー、サードパーティはありません。データは Mac に残ります
~/Library/Group Containers/…/pacer.sqlite であり、アプリの更新後も保持されます。
Pacer は SwiftUI + SwiftData + Charts であり、開発者 ID で署名され、公証されています。
Sparkle による自動更新。貢献を歓迎します —
COTRIBUTING.md を参照してください。
コンポーネント
役割
ペーサーアプリ
メイン UI + メニューバー項目。データ収集 (FSEvents JSONL スキャン + OAuth ポーリング) は、このバイナリ内でインプロセスで実行されます。個別のデーモンはありません。
ペーサーウィジェット
WidgetKit 拡張機能 — 共有アプリ グループ ストアを直接読み取ります。
ペーサーコア
共有 Swift パッケージ — パーサー、モデル、スキャン コーディネーター、再コンピューター。
ソースから構築する
日常使いに向けてリリースされたDMG
それはあなたが望むものです—

このセクションは Pacer のハッキングのみを対象としています。
要件: macOS 15 SDK (Xcode 16+) および xcodegen ( brew install xcodegen )。
Xcode プロジェクトは project.yml から生成されます。.xcodeproj は決して編集しないでください。
直接的に。
make verify # 署名されていないコンパイルのみのチェック (Apple アカウントは必要ありません)
make test # PacerCore ユニット + グラウンドトゥルース テスト
make install # 署名済み + 公証済みビルド → /Applications/Pacer.app
スクリーンショットを作成 # README スクリーンショットを再生成します (docs/screenshots.md を参照)
ヘルプを作成 # その他すべて
make verify と make test には署名のセットアップは必要なく、CI が実行するのはそれだけです。
実行可能なアプリを構築するには、独自の Apple Developer アカウントが必要です。
PACER_SIGN_IDENTITY="Developer ID Application: <Name> (<TEAMID>)" make install を使用してインストールします。
CONTRIBUTING →「自分で構築して実行する」を参照してください。
詳しい説明 (およびアプリ グループが署名とチーム ID を結び付ける理由) については、こちらをご覧ください。
AGENTS.md は、詳細なアーキテクチャ ガイド (パフォーマンスの不変条件、
SwiftData スキーマ、リコンピューター パターン); docs/ ではデザインについて説明します。
パフォーマンス チューニング、スクリーンショットの生成、およびリリース プロセス。
vX.Y.Z タグのプッシュがトリガーされる
.github/workflows/release.yml をビルドします。
DMG の署名、公証、パッケージ化、Sparkle アップデートの署名、GitHub の公開
gh-pages ブランチ上の appcast.xml をリリースおよび更新します。参照
秘密とリリースのカットについては docs/releasing.md
チェックリスト。
マサチューセッツ工科大学Pacer は Anthropic の Claude Code 製品からデータを読み取りますが、そうではありません。
Anthropic と提携または承認されています。
クロード コードの使用状況 (トークン、コスト、レート制限ペーシング、プロジェクトごとの内訳) を追跡するためのネイティブ macOS アプリ。 SwiftUI + SwiftData。
ericandrechek.github.io/Pacer/ トピック
Readme MIT ライセンス
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Native macOS app for tracking Claude Code usage — tokens, cost, rate-limit pacing, per-project breakdowns. SwiftUI + SwiftData. - EricAndrechek/Pacer

GitHub - EricAndrechek/Pacer: Native macOS app for tracking Claude Code usage — tokens, cost, rate-limit pacing, per-project breakdowns. SwiftUI + SwiftData. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
EricAndrechek
/
Pacer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
353 Commits 353 Commits Folders and files
.claude/ skills/ ship .claude/ skills/ ship .github .github App App PacerCore PacerCore Vendor Vendor Widgets Widgets bin bin docs docs research/ harness research/ harness .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md project.yml project.yml View all files Repository files navigation
Know what Claude Code is costing you — and how close you are to your limits — right from your Mac's menu bar.
Pacer is a free, open-source macOS app that keeps an eye on your Claude Code
usage: the tokens you're burning, what they cost, how close you are to your
5-hour and weekly rate limits, and where the spend is going — by project, by
model, by day. It sits quietly in your menu bar, keeps itself up to date, and
keeps all of your data on your own Mac.
Menu bar & click-down popover
Home-screen & Notification Center widgets
Share your 5-hour & 7-day burn charts as an image
Models — share, trend, and per-model table; group by model or class
Projects & Collections — group projects, scope totals to a collection
History — lifetime totals, six-month heatmap, monthly spend
Heads up — early days. Pacer is pre-1.0 and under active development. It
works and it's useful today, but expect the occasional rough edge, and your
saved history may need to be rebuilt between 0.x versions.
Download the latest version →
(grab the Pacer-x.y.z.dmg file under "Assets").
Open the downloaded file and drag Pacer into your Applications folder.
Launch Pacer. It lives in your menu bar — look for the little gauge icon
up top (there's no Dock icon). Click it to open the dashboard.
The first time, macOS asks permission for Pacer to read Claude Code's files.
Click Allow — that's how Pacer sees your usage. Nothing leaves your Mac.
That's all. Pacer keeps itself up to date automatically: when a new version
ships, it offers to install it for you — no re-downloading, no reinstalling.
You'll need macOS 15 (Sequoia) or newer, on either Apple Silicon or Intel.
Rate-limit pacing. Your 5-hour, weekly, and any per-model windows —
the scoped caps Anthropic reports for your account, like a weekly Fable
limit — as easy-to-read pace charts. Are you ahead, on track, or about to hit
the wall? Color bands (behind / on track / ahead / nearly maxed) tell you at a
glance, and every window gets the same forecast: projected fill,
time-to-limit, and calibrated bands. Refreshed every few minutes from
Anthropic's usage data.
Costs, your way. See spend the way Claude Code reports it, or have Pacer
price it from tokens itself — switchable in Settings. Daily, monthly, and
all-time totals included.
Where it's going. Break usage down by project — grouped into your own
non-destructive collections — and by model, with every Claude model in its
own color and groupable by class. Drill into any single day, and see a
GitHub-style activity heatmap of the last six months.
Live "today" view. Your current burn rate plus a running "at this pace,
today will end at about $X" projection.
Multiple accounts. Signed into more than one Claude account? Pacer tracks
each account's usage and limits separately and lets you switch between them.
At a glance, always. A configurable menu-bar readout — an activity-ring
icon, percent chips, and an explicit choice of which window drives the icon —
plus home-screen-style widgets, each pickable to whichever window you care
about (5-hour, weekly, or a per-model cap) for cost and pacing.
Optional nudges. Local notifications when you cross a rate-limit threshold
(50 / 75 / 90%) — on any window, including a per-model cap — or blow past a
daily spending limit you set. Off by default.
Export. Send daily totals, daily-by-model, or per-project numbers to a
CSV for your own spreadsheets.
You've got options for keeping an eye on Claude Code usage — Claude Code's own
/usage , the popular ccusage CLI, and
several menu-bar apps — the closest being Token Pacer ,
a paid, closed-source namesake, plus Claude God and
ccseva . They're good tools; here's
the honest lay of the land.
✦ estimates the windows from your local JSONL logs rather than reading
Anthropic's usage API. Token Pacer is a paid, closed-source app, so its column is
read from its public site rather than its source — and some of its features (CSV
export, custom model pricing) are Pro-tier. Best-effort as of June 2026 — these
tools all move fast, so corrections are welcome via an issue or PR.
vs. /usage — Claude Code's built-in view is a snapshot you ask for in
one terminal; Pacer is the always-on, zoomed-out companion that remembers every
session. The speedometer in one car vs. the dashboard that logs every trip.
vs. ccusage — a great CLI for scriptable numbers (Pacer's tests even
cross-check their scanner against it); Pacer is the GUI you glance at instead of
a command you re-run, and it reads your actual limit % from Anthropic rather
than estimating from logs.
vs. Claude God / ccseva — the closest rivals, and genuinely nice. Pacer
leans into pacing (your windows against an ideal-burn line, "will I run out
before the reset?"), native + quiet-by-default, and signed/notarized
self-update; Claude God goes further on ROI/git correlation and a plugin
marketplace, and ccseva on its glassy UI. Pick the one that thinks about your
usage the way you do — they coexist happily.
vs. Token Pacer — the close namesake, and the nearest thing to a direct
rival: also a native, local-first macOS menu-bar tracker, and it reaches
wider — covering Codex and opencode alongside Claude Code, with "next action"
nudges. The trade-offs: it's $29 and closed-source , where Pacer is free and
open; Pacer reads your real limit % from Anthropic and adds the ideal-burn
pacing line, the 6-month heatmap, and widgets. Multi-agent and willing to pay?
Fair pick. Claude Code-first and value open source? Pacer.
Pacer is local-first . It reads only the files Claude Code already writes to
your own Mac, and it sends nothing about your usage anywhere:
~/.claude/projects/*.jsonl — your session token usage, read on your machine.
The Claude Code login token in your macOS Keychain — used only to ask
Anthropic for your rate-limit status, the same way Claude Code itself does.
The only network connections Pacer makes are:
api.anthropic.com — to check your rate-limit windows (~12 small requests an
hour while it's running).
github.com — to check for, download, and install app updates (at launch,
then every 6 hours; installs happen in the background when automatic updates
are on).
No analytics, no telemetry, no third parties. Your data stays on your Mac in
~/Library/Group Containers/…/pacer.sqlite , and it persists across app updates.
Pacer is SwiftUI + SwiftData + Charts, signed with Developer ID and notarized,
auto-updating via Sparkle . Contributions welcome —
see CONTRIBUTING.md .
Component
Role
Pacer.app
Main UI + menu-bar item. Data collection (FSEvents JSONL scan + OAuth poll) runs in-process inside this binary — there is no separate daemon.
PacerWidgets
WidgetKit extension — reads the shared App Group store directly.
PacerCore
Shared Swift package — parsers, models, scan coordinator, recomputers.
Building from source
For everyday use, the released DMG
is what you want — this section is only for hacking on Pacer.
Requirements: macOS 15 SDK (Xcode 16+) and xcodegen ( brew install xcodegen ).
The Xcode project is generated from project.yml — never edit the .xcodeproj
directly.
make verify # unsigned compile-only check (no Apple account needed)
make test # PacerCore unit + ground-truth tests
make install # signed + notarized build → /Applications/Pacer.app
make screenshots # regenerate the README screenshots (see docs/screenshots.md)
make help # everything else
make verify and make test need no signing setup and are all that CI runs.
To build a runnable app you need your own Apple Developer account — point the
install at it with PACER_SIGN_IDENTITY="Developer ID Application: <Name> (<TEAMID>)" make install .
See CONTRIBUTING → "Building and running it yourself"
for the full story (and why an App Group ties signing to your Team ID).
AGENTS.md is the deep architectural guide (performance invariants,
the SwiftData schema, the recomputer pattern); docs/ covers the design,
perf-tuning, screenshot generation , and release process.
Pushing a vX.Y.Z tag triggers
.github/workflows/release.yml , which builds,
signs, notarizes, packages a DMG, signs the Sparkle update, publishes the GitHub
Release, and updates appcast.xml on the gh-pages branch. See
docs/releasing.md for the secrets and the cut-a-release
checklist.
MIT . Pacer reads data from Anthropic's Claude Code product but is not
affiliated with or endorsed by Anthropic.
Native macOS app for tracking Claude Code usage — tokens, cost, rate-limit pacing, per-project breakdowns. SwiftUI + SwiftData.
ericandrechek.github.io/Pacer/ Topics
Readme MIT license Contributing
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

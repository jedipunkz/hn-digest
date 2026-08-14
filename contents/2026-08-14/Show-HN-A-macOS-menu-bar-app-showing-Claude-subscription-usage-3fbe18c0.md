---
source: "https://github.com/palamim/usagent"
hn_url: "https://news.ycombinator.com/item?id=49297106"
title: "Show HN: A macOS menu bar app showing Claude subscription usage"
article_title: "GitHub - palamim/usagent: macOS menu bar app showing Claude subscription usage (5-hour + weekly limits) at a glance · GitHub"
author: "leopalamim"
captured_at: "2026-08-14T11:37:43Z"
capture_tool: "hn-digest"
hn_id: 49297106
score: 1
comments: 0
posted_at: "2026-08-14T11:00:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A macOS menu bar app showing Claude subscription usage

- HN: [49297106](https://news.ycombinator.com/item?id=49297106)
- Source: [github.com](https://github.com/palamim/usagent)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T11:00:51Z

## Translation

タイトル: HN を表示: クロードのサブスクリプションの使用状況を表示する macOS メニュー バー アプリ
記事タイトル: GitHub - palamim/usagent: クロードのサブスクリプション使用状況 (5 時間 + 週の制限) を一目で表示する macOS メニュー バー アプリ · GitHub
説明: クロードのサブスクリプション使用状況 (5 時間 + 週の制限) を一目で表示する macOS メニュー バー アプリ - Palamim/usagent
HN テキスト: 私は複数の Claude 製品を使用しており、使用制限を常にチェックしていますが、それを即座に行う方法はありませんでした。私は claude ai > 設定 > 使用法を開くか、CLI で /usage を押し続けました。その情報を常に目に見えて入手できる方法が必要でした。それで私はuseentを構築しました。これは、使用制限、5 時間のウィンドウ、および 1 週間のウィンドウを常に表示する macOS メニュー バー アプリです。デフォルトでは、usagent はその制限に近い方を表示します。クリックすると両方の詳細情報が表示されます。私は個人的に、関係なく常に 5 時間のウィンドウを表示するようにしています。私は二度と /usage を押したり、自分の限界を確認するために claude ai を開いたりすることはありませんでした。

記事本文:
GitHub - palamim/usagent: クロードのサブスクリプション使用状況 (5 時間 + 週の制限) を一目で表示する macOS メニュー バー アプリ · GitHub
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
パラミム
/
使用者
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .github/ workflows .github/ workflows リソース リソース スクリプト スクリプト Sources/ usent ソース/ usent テスト/ usentTests テスト/ usentTests docs docs .gitignore

.gitignore COTRIBUTING.md CONTRIBUTING.md Info.plist Info.plist ライセンス ライセンス Makefile Makefile Package.swift Package.swift README.md README.md VERSION VERSION すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードの 2 つの使用時計、つまり 5 時間のローリング ウィンドウと
週ごとの上限 — 上限に近いほうがメニューに表示されます
バー。クリックすると、使用率とリセット時間の両方が並べて表示されます。
macOS キーチェーンから Claude Code OAuth トークンを読み取ります - 同じです
プレーンテキスト ファイルではなく、OS 自体が使用するシステム資格情報ストア。
これを読むと、ログインのための macOS 認証プロンプトがトリガーされます
パスワードを入力すると、キーチェーンでいつでもアクセスを取り消すことができます
.appにアクセスします。 Anthropic 自身のエンドポイントを除き、ネットワーク呼び出しを行いません。
ソースは十分に短いので、信頼する前に自分で読んでください。
あなたの資格情報。
非公式であり、Anthropic とは提携していません — によって構築されました
文書化されていないエンドポイントのリバースエンジニアリング
( api.anthropic.com/api/oauth/usage ) そのため、警告なしに壊れる可能性があります
それが変わったら。
ドックアイコンも設定もテレメトリもありません。
1 つを選択します。どちらもディスク上に usent.app が残ります。
最新の usent-X.Y.Z.zip を次の場所から取得します。
をリリースします。
それを解凍し、 usent.app を /Applications に移動します。
usent.appを右クリック→最初のを開きます（ダブルクリックしないでください）
時間。 macOS は、正体不明の開発者からのものであると警告します。
アプリはアドホック署名されており、有料の Apple で公証されていないため
開発者ID。とにかく「開く」をクリックします。この警告は 1 回だけ表示されます。
その後は、ログイン項目経由も含めて通常どおり起動します。
最初の手動インストール後、 ./Scripts/update.sh (または make update )
新しいリリースの取得を自動化します。以下の「アップデート」を参照してください。
Xcode コマンド ライン ツール (Swift 5.9 以降、macOS 13 以降) が必要です。
git clone https://github.com/palamim/usagent.git
CD 使用法
アプリを作る

# 適切な Info.plist を使用して usent.app をビルドします (ドック アイコンなし)
make open # ビルドして開きます
ローカルに構築されたアプリは Gatekeeper によって隔離されません (これは単に発生するだけです)
ブラウザを通じてダウンロードされたファイルに保存されるため)、右クリックして開く手順はありません
ここで必要となります。必要に応じて、usagent.app を /Applications に移動します
どこかに永住する。
make run (単純な swift run ) も素早い反復に機能しますが、
開発中 — 開発中ではないため、起動時にドックアイコンが短時間点滅します。
バンドルから実行されますが、これは予期されており無害です。
swift test は、UsageStore 単体テスト (ステートマシンの動作、
バインディングクロック選択) モック化されたUsageFetchingに対する、ネットワークなし
またはキーチェーンアクセスが必要です。だけではなく、完全な Xcode.app が必要です。
コマンド ライン ツール — これは macOS の XCTest 要件であり、要件ではありません。
このプロジェクトが管理するもの。 CI ( .github/workflows/tests.yml )
GitHub の macOS ランナーにはプッシュのたびにスイートが実行されます。
Xcodeがプリインストールされています。
usent.app がどこかに永続的になったら ( /Applications など)、次を追加します
それをログイン項目に追加します。
システム設定→一般→ログイン項目と拡張機能→次の場所で開く
ログイン → + → usent.app を選択します。
ソースからリビルドする場合、make app は同じパスを上書きするため、
後でログイン項目として再追加する必要はありません。
make update # または ./Scripts/update.sh
最新の GitHub リリースをダウンロードし、実行中の usent を終了します。
/Applications/usagent.app を新しいビルドに置き換え、
ゲートキーパーの隔離フラグを設定して、右クリックせずに開くようにします。
それを再開します。まだ存在しない場合は、ログイン項目に追加するように求められます。
認証済みの gh CLI が必要です ( gh auth login )。
OAuth トークンは、サービス名の下で macOS キーチェーンに存在します。
Claude Code-credentials 、これではなく Claude Code CLI によって作成されます。
アプリ。ユーザントが初めてそれを読み取るとき、macOS

ログインを求めるプロンプトが表示されます
アクセスを許可するためのパスワード。 「常に許可」をクリックします（「許可」ではありません） —
そうしないと、起動または更新のたびに再度プロンプトが表示されます。
その後も繰り返しプロンプトが表示される場合は、キーチェーン項目の ACL
リセットされた可能性があります（例: 別のコードで再構築した後）
署名）。 「常に許可」で一度再認証すると、再度修正されます。
コンパニオン状態ファイルではなく直接 API 呼び出しを行う理由
クロードコード使用状況モニター
ネイティブ メニュー バー アプリがその --write-state 出力を使用するように招待します
Anthropic のエンドポイントを直接呼び出す代わりに。エンドポイントを呼び出します
代わりに直接: その状態ファイルの独自の「公式」データ ソースは
同じ OAuth 使用エンドポイント (オプトイン --api フラグの背後) なので、
それを経由するルーティングは、別の Python ツールを
バックグラウンドプロセスは何のメリットもありません - 依存関係と移動だけです
このアプリには必要のない部分です。
他の先行技術にも名前を付ける価値があります: ccusage
は、ほとんどの人が最初に手を伸ばす端末ツールであり、それをうまく機能します。
usent はそれを置き換えようとしているわけではありません。端末や端末の場合をカバーします。
ステータスライン ツールにはありません: 永続的で一目でわかるメニュー バー インジケーター
コマンドとして実行しているわけではありません。
Pro : このアプリが何に対して構築され、テストされているか。 5時間と
seven_day は表示されている 2 つの時計です。応答内のその他すべて
は null で無視されます。
Max : 2 つのメイン クロックに対して変更せずに動作するはずです - 同じ形状、
より高い限界。 Max プランも追加される予定です
seven_day_opus 、Opus の個別の (より厳しい) 週次サブキャップ。
usent は追加の「ウィークリー オーパス」行として表示され、
存在する場合、「限界に最も近い」計算。これは実装されています
防御的であるが、実際の Max アカウントに対して検証されていない場合 —
フィールドの形状または意味が異なることが判明しました。フィールドを開いてください。
編集された再に関する問題

スポンジボディ。
無料：対象外です。無料枠アカウントは Claude Code CLI を取得できません
アクセスするため、このアプリが最初に読み取る OAuth トークンはありません
場所。
チーム / エンタープライズ : 未処理。これらは組織がプールした制限を使用する可能性があります
このアプリは、個人的な 5 時間/週の時間枠ではなく、構築されています。
おそらく調整ではなく、別の機能です。
エンドポイント ( https://api.anthropic.com/api/oauth/usage ) は
文書化されていない。予告なく変更または消滅する場合がございます。
60 秒タイマーとクリック時に更新され、最大 1 回に 1 回に制限されます。
エンドポイントへの打撃を避けるために 15 秒。
Claude サブスクリプションの使用状況 (5 時間 + 週ごとの制限) が一目でわかる macOS メニュー バー アプリ
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS menu bar app showing Claude subscription usage (5-hour + weekly limits) at a glance - palamim/usagent

i use multiple Claude products, and i constantly check my usage limits, but there was never an instant way to do it. i kept opening claude ai > settings > usage or hitting /usage in the cli. i wanted a way to have that information always visible and at reach. so i built usagent. it's a macOS menu bar app always showing usage limits, the 5h window and the weekly window. by default, usagent shows whichever is closer to its limit. you can click to see both with full info. i, personally, make it always show the 5h window regardless. i never hit /usage again or opened claude ai just to see my limits.

GitHub - palamim/usagent: macOS menu bar app showing Claude subscription usage (5-hour + weekly limits) at a glance · GitHub
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
palamim
/
usagent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .github/ workflows .github/ workflows Resources Resources Scripts Scripts Sources/ usagent Sources/ usagent Tests/ usagentTests Tests/ usagentTests docs docs .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Info.plist Info.plist LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md VERSION VERSION View all files Repository files navigation
Tracks Claude's two usage clocks — the 5-hour rolling window and the
weekly cap — and shows whichever is closer to its limit in the menu
bar. Click it for both: percent used and time to reset, side by side.
Reads your Claude Code OAuth token from the macOS Keychain — the same
system credential store the OS itself uses, not a plaintext file.
Reading it triggers a macOS authorization prompt for your login
password, and you can revoke access at any time in Keychain
Access.app. Makes no network calls except to Anthropic's own endpoint.
The source is short enough to read yourself before trusting it with
your credentials.
Unofficial and not affiliated with Anthropic — built by
reverse-engineering an undocumented endpoint
( api.anthropic.com/api/oauth/usage ), so it may break without warning
if that changes.
No dock icon, no settings, no telemetry.
Pick one — both end up with usagent.app on disk.
Grab the latest usagent-X.Y.Z.zip from
Releases .
Unzip it and move usagent.app to /Applications .
Right-click usagent.app → Open (don't double-click) the first
time. macOS will warn it's from an unidentified developer — that's
because the app is ad-hoc signed, not notarized with a paid Apple
Developer ID. Click Open anyway. This warning only appears once;
after that it launches normally, including via Login Items.
After the first manual install, ./Scripts/update.sh (or make update )
automates picking up new releases — see Update below.
Requires the Xcode command line tools (Swift 5.9+, macOS 13+).
git clone https://github.com/palamim/usagent.git
cd usagent
make app # builds usagent.app with a proper Info.plist (no dock icon)
make open # builds and opens it
A locally built app isn't quarantined by Gatekeeper (that only happens
to files downloaded through a browser), so no right-click-to-open step
is needed here. Move usagent.app to /Applications if you want it to
live somewhere permanent.
make run (plain swift run ) also works for quick iteration while
developing — it flashes a dock icon briefly on launch since it isn't
run from a bundle, which is expected and harmless.
swift test runs the UsageStore unit tests (state-machine behavior,
binding-clock selection) against a mocked UsageFetching , no network
or Keychain access needed. Requires full Xcode.app , not just the
Command Line Tools — that's an XCTest requirement on macOS, not
something this project controls. CI ( .github/workflows/tests.yml )
runs the suite on every push, since GitHub's macOS runners ship with
Xcode preinstalled.
Once usagent.app is somewhere permanent (e.g. /Applications ), add
it to Login Items:
System Settings → General → Login Items & Extensions → Open at
Login → + → select usagent.app .
If you rebuild from source, make app overwrites the same path, so you
won't need to re-add it as a login item afterward.
make update # or ./Scripts/update.sh
Downloads the latest GitHub release, quits any running usagent ,
replaces /Applications/usagent.app with the new build, clears the
Gatekeeper quarantine flag so it opens without a right-click, then
reopens it. Prompts to add it to Login Items if it isn't already there.
Requires the gh CLI, authenticated ( gh auth login ).
The OAuth token lives in the macOS Keychain under the service name
Claude Code-credentials , created by the Claude Code CLI — not by this
app. The first time usagent reads it, macOS will prompt for your login
password to authorize access. Click "Always Allow" (not "Allow") —
otherwise you'll be prompted again on every launch or every refresh.
If you're prompted repeatedly even after that, the Keychain item's ACL
may have been reset (e.g. after rebuilding with a different code
signature). Re-authorizing once with "Always Allow" fixes it again.
Why direct API calls, not a companion state file
Claude-Code-Usage-Monitor
invites native menu bar apps to consume its --write-state output
instead of calling Anthropic's endpoint directly. We call the endpoint
directly instead: its own "official" data source for that state file is
the same OAuth usage endpoint (behind an opt-in --api flag), so
routing through it would mean running a separate Python tool as a
background process for no benefit — just a dependency and a moving
part this app doesn't need.
Worth naming the other prior art too: ccusage
is the terminal tool most people reach for first, and does that well.
usagent isn't trying to replace it — it covers the case a terminal or
status-line tool doesn't: a persistent, at-a-glance menu bar indicator
you're not running as a command.
Pro : what this app is built and tested against. five_hour and
seven_day are the two clocks shown; everything else in the response
is null and ignored.
Max : should work unchanged for the two main clocks — same shape,
higher limits. Max plans are also expected to populate
seven_day_opus , a separate (tighter) weekly sub-cap on Opus, which
usagent shows as an extra "Weekly Opus" row and folds into the
"closest to limit" calculation when present. This is implemented
defensively but not verified against a real Max account — if the
field's shape or meaning turns out to be different, please open an
issue with a redacted response body.
Free : out of scope. Free-tier accounts don't get Claude Code CLI
access, so there's no OAuth token for this app to read in the first
place.
Team / Enterprise : unhandled. These likely use org-pooled limits
rather than the personal five-hour/weekly windows this app is built
around — probably a different feature, not a tweak.
The endpoint ( https://api.anthropic.com/api/oauth/usage ) is
undocumented. It could change or disappear without notice.
Refreshes on a 60s timer and on click, throttled to at most once per
15s to avoid hammering the endpoint.
macOS menu bar app showing Claude subscription usage (5-hour + weekly limits) at a glance
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/benri-ai/slack-debloat"
hn_url: "https://news.ycombinator.com/item?id=48854923"
title: "Show HN: Debloat Slack – Get rid of Slackbot, AI upsells and other annoyances"
article_title: "GitHub - benri-ai/slack-debloat: Remove Slackbot and other useless features and get actual work done! Currently macOS only. · GitHub"
author: "ed_mercer"
captured_at: "2026-07-10T02:49:08Z"
capture_tool: "hn-digest"
hn_id: 48854923
score: 2
comments: 0
posted_at: "2026-07-10T02:08:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Debloat Slack – Get rid of Slackbot, AI upsells and other annoyances

- HN: [48854923](https://news.ycombinator.com/item?id=48854923)
- Source: [github.com](https://github.com/benri-ai/slack-debloat)
- Score: 2
- Comments: 0
- Posted: 2026-07-10T02:08:12Z

## Translation

タイトル: Show HN: Debloat Slack – Slackbot、AI アップセル、その他の煩わしさを取り除く
記事タイトル: GitHub - benri-ai/slack-debloat: Slackbot やその他の役に立たない機能を削除して、実際の作業を完了してください。現在は macOS のみです。 · GitHub
説明: Slackbot やその他の役に立たない機能を削除して、実際の作業を完了してください。現在は macOS のみです。 - べんりあい/たるみデブロート

記事本文:
GitHub - benri-ai/slack-debloat: Slackbot やその他の役に立たない機能を削除して、実際の作業を完了してください。現在は macOS のみです。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
べんりあい
/
たるみデブレーション
公共
通知

ション
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット docs docs .gitignore .gitignore ライセンス ライセンス README.md README.mdcatalog.mjs category.mjs config.json.example config.json.examplecustom.css.examplecustom.css.examplecustom.js.examplecustom.js.exampleeval.mjseval.mjsinject.mjs inject.mjs install.sh install.sh launch-slack.sh launch-slack.sh スクリーンショット.mjs スクリーンショット.mjslack-debloat.shlack-debloat.sh uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
独自の CSS と JS を使用して、macOS 上のネイティブ Slack アプリの肥大化を解消します。uBlock Origin + Tampermonkey のように、ただしデスクトップ アプリの場合です。
Slack の macOS アプリは Electron、つまりブラウザと同じ Web アプリをレンダリングする Chromium です。通常は拡張機能を入れることはできません。 lack-debloat は、Chrome DevTools プロトコルのポートを開いた状態 (localhost のみ) で Slack を起動し、custom.css /custom.js をすべての Slack ウィンドウにプッシュし、リロード後も保持し、保存するたびに 1 秒以内にライブ再適用する小さな依存関係のないインジェクターを実行します。 [アクティビティ] タブ、ハドル ボタン、Slackbot、アップセル バナーなど、気になるものはすべて非表示にします。
非表示の行は実際に折りたたまれます (空白スロットはありません。以下の自動リフローを参照)。残っている行のスタイルを変更できます。ここでは、スレッドが 18 ピクセルの半太字に昇格しました。
Slack Debloat.app (Dock アイコン) LaunchAgent (ログイン)
│ │
└─ inject.mjs ポーリング localhost:9222 を使用して Slack を起動します。
--remote-debugging-port=9222 は CDP 経由ですべての Slack ウィンドウに接続します。
(localhost のみ)custom.css +custom.js を挿入します。
ファイルを保存するたびに再挿入します
Slack.app にパッチは適用されていません (コード署名されており、これを変更すると Gatekeeper と Keychain が壊れます)。 Slack 自体のアップデートは標準的

同盟者。
macOS、/Applications の Slack デスクトップ アプリ
Node.js ≥ 22 (組み込みの WebSocket クライアントを使用 - npm 依存関係なし)
git clone https://github.com/benri-ai/slack-debloat.git
cd スラックデブロート
./install.sh
次に、新しい Slack Debloat アプリ (Slack 独自のアイコンが表示されます) を介して Slack を起動し、Slack の代わりに Dock に置きます。それでおしまい。
.example ファイルから config.json /custom.css /custom.js をシードします (コピーは gitignore されます)。
ログインからインジェクターを実行し続ける LaunchAgent をインストールします。
/Applications/Slack Debloat.app をビルドします。
何もインストールせずに試してみます。 ./slack-debloat.sh は、デバッグ ポートを使用して Slack を再起動し、ターミナルでインジェクターを実行します (停止するには Ctrl-C、Slack は実行し続けます)。
Slack ウィンドウの最上部にある細い緑色の線は、インジェクションが機能していることを確認します (セットアップが信頼できる場合は、custom.css で削除してください)。
config.json — 組み込みオプション
config.json (インストール時に config.json.example からシードされる) は、ブール値へのオプション キーのフラット マップです。キーを true に切り替えて保存すると、実行中の Slack が 1 秒以内に更新されます。再起動や再構築は必要ありません。これは意図的にプレーンな JSON です。手動で編集するか、LLM にこの README を指定して反転させます。
{
"hide-slackbot-dm" : true 、
"アップセルバナーの非表示" : true 、
"より大きなスレッド行" : true
}
(省略したキーはデフォルトでオフになります。不明なキーは無視され、 injector.log に警告が表示されます。)
これらのキーの背後にあるセレクターは、Slack 4.50 で検証された category.mjs に存在します。 Slack のアップデートでは、名前が変更されることがあります。修正は通常 2 分間の DevTools ジョブ (下記を参照) で完了します。PR は歓迎です。
custom.css /custom.js — その他すべて
手動で非表示にする: カタログにないものについては、custom.css を編集して保存します。同じ ~1 秒のライブが適用されます。 .example ファイルには、パターンを示すコメントアウトされたルールが同梱されています。
検索セレクター: wh

Slack はデブロートで実行され、Chrome で http://localhost:9222 を開くと、ネイティブ アプリに対する完全な DevTools が入手できます。 Slack の [data-qa=...] 属性を優先します。これらは、生成されたクラス名よりもはるかに安定しています。サイドバーの行は仮想化されているため、行ラッパー全体を非表示にします。
[ data-qa = "virtual-list-item" ] : has ([ data-qa ^= "channel_sidebar_pslackbot" ]) {
表示: なし!重要;
}
この方法で行を非表示にしても、空のスロットは残りません。インジェクターには、非表示の仮想リスト行を検出して残りの行を上にシフトする自動リフロー シムが同梱されており、Slack が再レンダリングされるたびに再適用されます。サイドバーは、あたかも項目が存在しなかったかのように、きれいに折りたたまれます。
スクリプトの実行:custom.js は、読み込み時と保存のたびにすべての Slack ウィンドウで実行されます (冪等なコードを作成します)。
端末から DOM を起動します。
ノード eval.mjs ' document.title '
ノード Screenshot.mjslack.png # フルウィンドウ
ノード Screenshot.mjs Sidebar.png ' .p-channel_sidebar ' # セレクターにクリップされました
よくある質問
通常の Slack アプリを開いたらどうなるでしょうか?
標準の Slack (デバッグ ポートなし) が得られ、インジェクターはアイドル状態になるだけです。 Slack Debloat をクリックすると終了し、ポートを使用して再起動します。 Slack 自体の「ログイン時にアプリを起動する」設定がオンになっている場合は、起動するたびにストックが開始されることに注意してください。オフにして、代わりにラッパーをクリックしてください。
これは Slack のアップデートに耐えられますか?
仕組みは、そうです。 Slack が名前を変更する場合、個々の CSS セレクターに微調整が必​​要になる場合があります。これが DevTools-at-localhost:9222 ワークフローです。
ハドル/通話/通知は引き続き機能しますか?
はい、これは修正されていない Slack アプリです。何もパッチもプロキシも適用されていません。
セキュリティ上の注意: DevTools ポートは localhost のみにバインドされていますが、Slack がこの方法で実行されている間、任意のローカル プロセスがそれに接続され、Slack セッションを読み取る可能性があります。個人のマシンでは問題ありません。トレードオフを知っています。
アンインストール: ./uninstall.sh 、Slack を終了して再度開きます

普通に。
Slack/Salesforce と提携または承認されていません。化粧品の使用は自己責任で行ってください。
Slackbot やその他の役に立たない機能を削除して、実際の作業を完了してください。現在は macOS のみです。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

Remove Slackbot and other useless features and get actual work done! Currently macOS only. - benri-ai/slack-debloat

GitHub - benri-ai/slack-debloat: Remove Slackbot and other useless features and get actual work done! Currently macOS only. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
benri-ai
/
slack-debloat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits docs docs .gitignore .gitignore LICENSE LICENSE README.md README.md catalog.mjs catalog.mjs config.json.example config.json.example custom.css.example custom.css.example custom.js.example custom.js.example eval.mjs eval.mjs inject.mjs inject.mjs install.sh install.sh launch-slack.sh launch-slack.sh screenshot.mjs screenshot.mjs slack-debloat.sh slack-debloat.sh uninstall.sh uninstall.sh View all files Repository files navigation
De-bloat the native Slack app on macOS with your own CSS and JS — like uBlock Origin + Tampermonkey, but for the desktop app.
Slack's macOS app is Electron, i.e. a Chromium rendering the same web app as your browser — you just normally can't put extensions in it. slack-debloat launches Slack with its Chrome DevTools Protocol port open (localhost-only) and runs a tiny zero-dependency injector that pushes your custom.css / custom.js into every Slack window, keeps them across reloads, and live-reapplies within a second whenever you save . Hide the Activity tab, the huddle button, Slackbot, upsell banners — whatever annoys you.
Hidden rows collapse for real (no blank slots — see auto-reflow below), and you can restyle what stays: here Threads got promoted to 18px semibold.
Slack Debloat.app (Dock icon) LaunchAgent (login)
│ │
└─ launches Slack with └─ inject.mjs polls localhost:9222,
--remote-debugging-port=9222 attaches to every Slack window via CDP,
(localhost only) injects custom.css + custom.js,
re-injects on every file save
No patching of Slack.app (it's code-signed; modifying it breaks Gatekeeper and Keychain). Slack updates itself normally.
macOS, Slack desktop app in /Applications
Node.js ≥ 22 (uses the built-in WebSocket client — no npm dependencies)
git clone https://github.com/benri-ai/slack-debloat.git
cd slack-debloat
./install.sh
Then launch Slack via the new Slack Debloat app (it wears Slack's own icon) and put it in your Dock in place of Slack. That's it.
seeds config.json / custom.css / custom.js from the .example files (your copies are gitignored),
installs a LaunchAgent that keeps the injector running from login,
builds /Applications/Slack Debloat.app .
Try it without installing anything: ./slack-debloat.sh relaunches Slack with the debug port and runs the injector in your terminal (Ctrl-C to stop; Slack keeps running).
A thin green line at the very top of the Slack window confirms injection is working (remove it in custom.css once you trust the setup).
config.json — the built-in options
config.json (seeded from config.json.example on install) is a flat map of option keys to booleans. Flip a key to true , save, and the running Slack updates within a second — no restart, no rebuild. It's plain JSON on purpose: edit it by hand, or point your LLM at this README and let it do the flipping.
{
"hide-slackbot-dm" : true ,
"hide-upsell-banners" : true ,
"bigger-threads-row" : true
}
(Keys you omit default to off; unknown keys are ignored with a warning in injector.log .)
The selectors behind these keys live in catalog.mjs , verified on Slack 4.50. Slack updates occasionally rename them; fixes are usually a two-minute DevTools job (see below) — PRs welcome.
custom.css / custom.js — everything else
Hide things by hand: for anything not in the catalog, edit custom.css , save — same ~1s live apply. The .example file ships with commented-out rules showing the patterns.
Find selectors: while Slack runs debloated, open http://localhost:9222 in Chrome and you get full DevTools against the native app. Prefer Slack's [data-qa=...] attributes — they're much more stable than the generated class names. Sidebar rows are virtualized, so hide the whole row wrapper:
[ data-qa = "virtual-list-item" ] : has ([ data-qa ^= "channel_sidebar_pslackbot" ]) {
display : none !important ;
}
Rows hidden this way don't leave blank slots: the injector ships an auto-reflow shim that detects hidden virtual-list rows and shifts the remaining rows up, re-applying whenever Slack re-renders. The sidebar collapses cleanly, as if the items never existed.
Run scripts: custom.js runs in every Slack window on load and on every save (write idempotent code).
Poke at the DOM from a terminal:
node eval.mjs ' document.title '
node screenshot.mjs slack.png # full window
node screenshot.mjs sidebar.png ' .p-channel_sidebar ' # clipped to a selector
FAQ
What if I open the regular Slack app?
You get stock Slack (no debug port), and the injector just idles. Clicking Slack Debloat quits and relaunches it with the port. Note that if Slack's own "Launch app on login" setting is on, it starts stock at every boot — turn it off and click the wrapper instead.
Does this survive Slack updates?
The mechanism, yes. Individual CSS selectors may need a tweak when Slack renames things — that's the DevTools-at-localhost:9222 workflow.
Do huddles / calls / notifications still work?
Yes — it's the unmodified Slack app; nothing is patched or proxied.
Security note: the DevTools port is bound to localhost only, but while Slack runs this way, any local process could attach to it and read your Slack session. Fine on a personal machine; know the trade-off.
Uninstall: ./uninstall.sh , then quit and reopen Slack normally.
Not affiliated with or endorsed by Slack/Salesforce. Cosmetic use at your own risk.
Remove Slackbot and other useless features and get actual work done! Currently macOS only.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

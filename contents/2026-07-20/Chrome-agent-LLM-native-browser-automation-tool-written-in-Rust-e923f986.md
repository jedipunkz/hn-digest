---
source: "https://github.com/sderosiaux/chrome-agent"
hn_url: "https://news.ycombinator.com/item?id=48982835"
title: "Chrome-agent: LLM-native browser automation tool written in Rust"
article_title: "GitHub - sderosiaux/chrome-agent: Browser automation for AI agents. Single Rust binary, zero deps, CDP direct to Chrome. · GitHub"
author: "chtefi"
captured_at: "2026-07-20T19:26:13Z"
capture_tool: "hn-digest"
hn_id: 48982835
score: 2
comments: 0
posted_at: "2026-07-20T18:25:23Z"
tags:
  - hacker-news
  - translated
---

# Chrome-agent: LLM-native browser automation tool written in Rust

- HN: [48982835](https://news.ycombinator.com/item?id=48982835)
- Source: [github.com](https://github.com/sderosiaux/chrome-agent)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T18:25:23Z

## Translation

タイトル: Chrome-agent: Rust で書かれた LLM ネイティブのブラウザ自動化ツール
記事のタイトル: GitHub - sderosiaux/chrome-agent: AI エージェントのブラウザ自動化。単一の Rust バイナリ、ゼロ deps、Chrome への CDP 直接。 · GitHub
説明: AI エージェントのブラウザ自動化。単一の Rust バイナリ、ゼロ deps、Chrome への CDP 直接。 - sderosiaux/chrome-agent

記事本文:
GitHub - sderosiaux/chrome-agent: AI エージェントのブラウザ自動化。単一の Rust バイナリ、ゼロ deps、Chrome への CDP 直接。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
スデロシオー
/
クロムエージェント
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
145 コミット 145 コミット .github/ workflows .github/ workflows docs docs npm npm scripts scripts skill/ chrome-agent skill/ chrome-agent src src テスト テスト ベンダー ベンダー .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンスREADME.cn.md README.cn.md README.md README.md Clippy.toml Clippy.toml llm-guide.txt llm-guide.txt Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
LLM を話すブラウザ自動化。
免責事項: これは独立したコミュニティ主導のプロジェクトです。 Google または Chrome チームと提携、承認、後援されているわけではありません。
あなたはユーザーではありません。あなたのLLMは。
この README を読む必要はありません。あなたのエージェントはそうします。これをインストールし、 chrome-agent --help を実行して、LLM に認識させます。 CLI には独自の使用ガイドが埋め込まれており、すべてのエラーには次のアクションのヒントが含まれており、--json モードでは単一のアダプターを作成しなくてもエージェントが解析できる構造化データが出力されます。このページがここにあるのは、GitHub がそれを期待しているためです。
これはエージェントブラウザとどう違うのですか?
Agent-Browser (Vercel) は、ダッシュボード、クラウド プロバイダー、注釈付きスクリーンショット、iOS サポート、AI チャット、認証ボールト、40,000 行の Rust などの機能を備えたブラウザー自動化プラットフォームです。素晴らしいですね。
chrome-agent はその逆です。機能を追加する代わりに、トークンを削除します。
Agent-Browser は、モニタリング、クラウド ブラウザ、およびビジュアル デバッグを備えたプラットフォームを提供します。 chrome-agent は、LLM に Web ページの最小表現を提供し、邪魔にならないようにします。エージェントにダッシュボードが必要な場合は、エージェント ブラウザを使用します。エージェントがページ解析ではなく推論にトークンを費やす必要がある場合、

これ。
エージェントがページの理解に費やすすべてのトークンは、タスクに関する推論に費やさないトークンです。 chrome-agent は、「このページはどのように見えるか?」の間のトークンを最小限に抑えるという 1 つのアイデアに基づいて構築されています。そして「次は何をすればいいですか？」
DOM 上のアクセシビリティ ツリー。 Playwright は、生の HTML の最大 2,000 トークンを返します。 chrome-agent は、安定した要素 ID を持つ a11y ツリーの最大 50 個のトークンを返します。 CSS セレクターを記述する必要も、DOM を解析する必要もありません。
1 つのバイナリでランタイムはゼロ。 3 MB Rust バイナリ。 Node.js、npm、Playwright ランタイムはありません。 npx chrome-agent は正常に動作します。 Linux ビルドは完全に静的 (musl) です。glibc 依存関係はなく、どのディストリビューションでも実行できます。
1 回の通話でアクション + 観察。アクション コマンドの --inspect は、アクション後のページの状態を返します。 2往復ではなく1往復です。
エラーは指示です。すべてのエラーには、エージェントに次に何をすべきかを伝えるヒント フィールドが含まれています。 {"ok":false、"error":"..."、"ヒント":"検査の実行"} 。
デフォルトの意図はステルスです。誰も話題にしていない検出ベクトル ( Runtime.enable ) を含む 7 つの CDP パッチ。本物の Chrome に接続して、最も強力な保護を実現します。
セレクターを使用しないコンテンツ抽出。記事の読み取り、繰り返しデータの抽出、API ペイロードのネットワーク。エージェントは CSS セレクターを作成しません。
これは、汎用のブラウザ テスト フレームワークではありません。これは、LLM を効果的に Web ブラウジングできるようにするツールです。
chrome-agent goto news.ycombinator.com --inspect
# ~2,000 ではなく ~50 トークン:
uid=n1 RootWebArea " ハッカー ニュース "
uid=n50 見出し「ハッカーニュース」 level=1
uid=n82 リンク「HN を表示: 新しいブラウザ ツール」
uid=n97 リンク " Rust 2025 Edition が発表されました "
...
# クリック + 1 回の呼び出しで新しいページを表示:
chrome-agent click n82 --inspect
UID は Chrome の backendNodeId に基づいています。それらは検査ごとに変わりません。今すぐ、または 5 分後に [n82] をクリックします。
クロムエージェント (3

MB Rustバイナリ）
| WebSocket 上の CDP
v
Chrome (ヘッドレス、Node.js、ランタイムなし)
なぜこれが存在するのか
これに当たったら…
chrome-agent が代わりにこれを実行します
Playwright のスナップショットは 2K トークンを消費します
a11y ツリー: ~50 トークン。ページ状態に費やされるコンテキストが 40 分の 1 に減少します。
デプロイのたびに CSS セレクターが壊れる
Chrome の backendNodeId からの UID。 DOM ノードが存在する限り安定しています。
クリックしてから検査 = 2 往復
--任意のコマンドで検査します。 1 つのコール、アクション + 観察。
200MB のノード + npm + Playwright
3MBのバイナリ。 npx chrome-agent はそのまま使用できます。
Cloudflare がヘッドレス Chrome をブロックする
7 CDP パッチ。 Runtime.enable は呼び出されませんでした (誰も話題にしていない検出ベクトル)。
サイトごとのスクレイピング セレクターの作成
記事の場合は読み取り、リスト/テーブル/カードの場合は抽出、API ペイロードの場合はネットワーク。セレクターはありません。
エラーはスタックトレースです
{"ok":false, "error":"...", "hint":"run Inspection"} -- 解析可能、実行可能。
各コマンドにより新しいブラウザが起動されます
セッションは継続します。 Chrome は呼び出し間で動作し続けます。 〜10msの起動。
エージェントがログインしたアカウントにアクセスできない
--copy-cookies は実際の Chrome から Cookie を取得します。 X.com、Gmail、ダッシュボードで動作します。
無限スクロールで 10 個の項目が表示されます
Inspection --scroll --limit 50 スクロールして収集します。 X.com でテスト: ライブ タイムラインから 50 件のツイート。
2 人のエージェントが 1 つのブラウザを共有 = カオス
--browser エージェント1 、 --browser エージェント2 。個別の Chrome インスタンス。
インストール
# AI エージェントの場合 -- エージェントが自動的に読み取る SKILL.md をインストールします
npx スキルで sderosiaux/chrome-agent を追加
# またはバイナリだけ
npm install -g chrome-agent # 事前構築済み
npx chrome-agent --help # インストールなし
カーゴはソースから chrome-agent # をインストールします
クイックスタート
# ページに移動して表示する
chrome-agent goto https://example.com --inspect
# uid でクリック
chrome-agent click n12 --inspect
# フォームに記入する
chrome-agent fill --uid n20 " を使用

r@test.com "
# CSS セレクターも機能します
chrome-agent click --selector " button.submit "
chrome-agent fill --selector " input[name=email] " " hello@test.com "
# 記事の内容 (読みやすさ -- Firefox リーダー モードのような)
クロムエージェントの読み取り
# 表示されるテキスト、スコープ付き、キャップ付き
chrome-agent text --selector " main " --truncate 500
# JSを実行する
chrome-agent eval " document.title "
# スクリーンショット (バイナリではなくファイル パスを返します)
クロムエージェントのスクリーンショット
コマンド
コマンド
何をするのか
goto <url> [--inspect] [--max- Depth N] [--header "K: V"]
ナビゲートします。自動プレフィックス https:// 。 --header (反復可能) は追加の HTTP ヘッダーを送信します。
戻る
歴史が戻ってきました。
前へ
歴史は前進します。
閉じる [--パージ]
ブラウザを停止します。 --purge は Cookie/プロファイルを削除します。
検査
コマンド
何をするのか
Inspection [--verbose] [--max- Depth N] [--uid nN] [--filter "role,role"] [--scroll] [--limit N] [--urls] [--max-chars N] [--offset K]
UID を含む a11y ツリー。無限スクロールの場合は --scroll --limit。 --urls はリンク上の href を解決します。 --max-chars / --offset 出力のキャップとページングを行います。
差分
前回の検査以降に何が変わったか。
スクリーンショット [--ファイル名名] [--形式 jpeg|png] [--品質 N] [--最大幅 N] [--uid nN|--セレクター "css"]
スクリーンショットをファイルに保存します。 JPEG/品質/最大幅で縮小します。 --uid / --selector 1 つの要素にクリップします。
pdf [--ファイル名] [--風景] [--背景]
現在のページを PDF ファイルに印刷します。
タブ
開いているタブをリストします。
インタラクション
コマンド
何をするのか
<uid> [--inspect] をクリックします
uid でクリックします。ボックス モデルがない場合は、JS .click() にフォールバックします。
--selector "css" [--inspect] をクリックします
CSSセレクターでクリックします。
--xy 100,200 をクリックします
座標をクリックしてください。
dblclick <uid> [--inspect]
uid、 --selector 、または --xy を指定してダブルクリックします。
fill --uid <uid> <値> [--inspect]
uidで入力を埋めます。
fill --selector "css" <値>
セレクターで入力します。
入力フォーム <uid=val>...
バッチフィル。
--u を選択してください

id <uid> <値>
値または表示されるテキストによってドロップダウン オプションを選択します。
select --selector "css" <値>
CSSセレクターで選択します。
<uid> を確認してください
チェックボックス/ラジオがオンになっていることを確認します。冪等。
<uid> のチェックを外します
チェックボックス/ラジオがオフになっていることを確認します。冪等。
アップロード --uid <uid> <ファイル>...
ファイルをファイル入力にアップロードします。
アップロード --selector "css" <ファイル>...
CSSセレクターでアップロードします。
<from-uid> <to-uid> をドラッグします
要素を別の要素にドラッグします。
type <テキスト> [--selector "css"]
フォーカスされた要素に入力します。
<キー>を押してください
Enter、Tab、Escape など。
スクロール <下|上|UID>
ページまたは要素をスクロールして表示します。
ホバー <uid>
ホバーします。
wait <テキスト|URL|セレクター> <パターン>
条件を待ちます。
ネットワークアイドルを待つ [--idle-ms N] [--timeout N]
ネットワークが --idle-ms (デフォルトは 500) の間静かになるまで待ちます。 SPA/XHR 解決のための固定スリープを打ち破ります。
コンテンツ抽出
コマンド
何をするのか
読み取り [--html] [--truncate N]
Mozilla Readability による記事の抽出。
テキスト [uid] [--selector "css"] [--truncate N]
ページまたは要素からの表示テキスト。
eval <式> [--selector "css"]
ページコンテキスト内の JS。 el = 一致した要素。
抽出 [--selector "css"] [--limit N] [--scroll] [--a11y]
繰り返しデータを自動検出します。 React SPA の場合は --a11y (X.com)。
ダウンロード <url> [--out パス] [--timeout N]
ページ内で取得された URL をダウンロードするため、Cookie/認証が引き継がれます (ログインゲートされたファイル)。 {path,bytes,mime} を返します。
モニタリング
コマンド
何をするのか
ネットワーク [--フィルター "パターン"] [--body] [--live N] [--abort "パターン"]
ネットワークリクエストとAPIレスポンス。 --abort は一致するリクエストをブロックします。
コンソール [--レベルエラー] [--クリア]
console.log/warn/error + JS 例外。
上級者向け
コマンド
何をするのか
フレーム <セレクター|メイン>
eval/inspect を iframe に切り替えます (またはメインに戻ります)。パイプ/バッチ プロセス内でのみ保持されます。
バッチ
標準入力の JSON 配列から複数のコマンドを実行します。
パイプ
ペルシス

テント JSON stdin/stdout 接続。
グローバルフラグ
--browser <名前> 名前付きブラウザ プロファイル (デフォルト: "default")
--page <name> 名前付きタブ (デフォルト: "default")
--connect [url] 実行中の Chrome に接続します
--headed ブラウザ ウィンドウを表示 (デフォルト: ヘッドレス)
--stealth 検出防止パッチ (Cloudflare、Turnstile)
--copy-cookies 実際の Chrome プロフィールの Cookie を使用する
--timeout <秒> コマンドのタイムアウト (デフォルト: 30)
--max- Depth <N> 検査の深さを制限します
--ignore-https-errors 自己署名証明書を受け入れる
--json 構造化された JSON 出力
--dialog <モード> JS ダイアログ ポリシー: 受け入れる (デフォルト)、閉じる、または手動
--dialog-text <text> --dialog accept の場合に、prompt() ダイアログに送信するテキスト
JS ダイアログ (alert/confirm/prompt/beforeunload) はデフォルトで自動応答されます ( --dialog accept )。そうでない場合、ネイティブ ダイアログは DOM シグナルなしでページをブロックし、エージェントの次のコマンドがハングします。キャンセルするには --dialog dismiss を使用し、オプトアウトするには --dialog manual を使用します。
ループ: 検査、実行、検査
chrome-agent goto https://app.com/login --inspect
# uid=n52 テキストボックス「メール」フォーカス可能
# uid=n58 テキストボックス「パスワード」フォーカス可能
# uid=n63 ボタン「サインイン」フォーカス可能
chrome-agent fill --uid n52 " user@test.com "
chrome-agent fill --uid n58 " passw

[切り捨てられた]

## Original Extract

Browser automation for AI agents. Single Rust binary, zero deps, CDP direct to Chrome. - sderosiaux/chrome-agent

GitHub - sderosiaux/chrome-agent: Browser automation for AI agents. Single Rust binary, zero deps, CDP direct to Chrome. · GitHub
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
sderosiaux
/
chrome-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
145 Commits 145 Commits .github/ workflows .github/ workflows docs docs npm npm scripts scripts skills/ chrome-agent skills/ chrome-agent src src tests tests vendor vendor .gitignore .gitignore CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.cn.md README.cn.md README.md README.md clippy.toml clippy.toml llm-guide.txt llm-guide.txt rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Browser automation that speaks LLM.
Disclaimer: This is an independent, community-driven project. It is not affiliated with, endorsed by, or sponsored by Google or the Chrome team.
You're not the user. Your LLM is.
You don't need to read this README. Your agent does. Install it, run chrome-agent --help , and let the LLM figure it out. The CLI embeds its own usage guide, every error comes with a hint for the next action, and --json mode outputs structured data an agent can parse without you writing a single adapter. This page is here because GitHub expects one.
How is this different from agent-browser?
agent-browser (Vercel) is a feature-complete browser automation platform: dashboard, cloud providers, annotated screenshots, iOS support, AI chat, auth vault, 40K lines of Rust. It's excellent.
chrome-agent is the opposite bet. Instead of adding features, it removes tokens.
agent-browser gives you a platform with monitoring, cloud browsers, and visual debugging. chrome-agent gives your LLM the smallest possible representation of a webpage and gets out of the way. If your agent needs a dashboard, use agent-browser. If your agent needs to spend tokens on reasoning instead of page parsing, use this.
Every token your agent spends understanding a page is a token it doesn't spend reasoning about the task. chrome-agent is built around one idea: minimize the tokens between "what does this page look like?" and "what should I do next?"
Accessibility tree over DOM. Playwright returns ~2,000 tokens of raw HTML. chrome-agent returns ~50 tokens of a11y tree with stable element IDs. No CSS selectors to write, no DOM to parse.
One binary, zero runtime. 3 MB Rust binary. No Node.js, no npm, no Playwright runtime. npx chrome-agent just works. Linux builds are fully static (musl) — no glibc dependency, runs on any distro.
Action + observation in one call. --inspect on any action command returns the page state after the action. One round-trip instead of two.
Errors are instructions. Every error includes a hint field telling the agent what to do next. {"ok":false, "error":"...", "hint":"run inspect"} .
Stealth by default intent. 7 CDP patches including the detection vector nobody talks about ( Runtime.enable ). Connect to real Chrome for the hardest protections.
Content extraction without selectors. read for articles, extract for repeating data, network for API payloads. The agent never writes CSS selectors.
This is not a general-purpose browser testing framework. It's a tool that makes an LLM effective at browsing the web.
chrome-agent goto news.ycombinator.com --inspect
# ~50 tokens instead of ~2,000:
uid=n1 RootWebArea " Hacker News "
uid=n50 heading " Hacker News " level=1
uid=n82 link " Show HN: A New Browser Tool "
uid=n97 link " Rust 2025 Edition Announced "
...
# Click + see the new page in one call:
chrome-agent click n82 --inspect
UIDs are based on Chrome's backendNodeId . They don't change between inspects. Click n82 now or five minutes from now.
chrome-agent (3 MB Rust binary)
| CDP over WebSocket
v
Chrome (headless, no Node.js, no runtime)
Why this exists
If you've hit this...
chrome-agent does this instead
Playwright snapshots burn 2K tokens
a11y tree: ~50 tokens. 40x less context spent on page state.
CSS selectors break after every deploy
UIDs from Chrome's backendNodeId . Stable as long as the DOM node exists.
Click then inspect = 2 round-trips
--inspect on any command. One call, action + observation.
200MB of Node + npm + Playwright
3 MB binary. npx chrome-agent works out of the box.
Cloudflare blocks your headless Chrome
7 CDP patches. Runtime.enable never called (the detection vector nobody talks about).
Writing per-site scraping selectors
read for articles, extract for lists/tables/cards, network for API payloads. No selectors.
Errors are stack traces
{"ok":false, "error":"...", "hint":"run inspect"} -- parseable, actionable.
Each command launches a fresh browser
Sessions persist. Chrome stays alive between calls. ~10ms startup.
Agent can't access your logged-in accounts
--copy-cookies grabs cookies from your real Chrome. Works with X.com, Gmail, dashboards.
Infinite scroll shows 10 items
inspect --scroll --limit 50 scrolls and collects. Tested on X.com: 50 tweets from a live timeline.
Two agents sharing one browser = chaos
--browser agent1 , --browser agent2 . Separate Chrome instances.
Install
# For AI agents -- installs a SKILL.md your agent reads automatically
npx skills add sderosiaux/chrome-agent
# Or just the binary
npm install -g chrome-agent # prebuilt
npx chrome-agent --help # no install
cargo install chrome-agent # from source
Quick start
# Navigate and see the page
chrome-agent goto https://example.com --inspect
# Click by uid
chrome-agent click n12 --inspect
# Fill a form
chrome-agent fill --uid n20 " user@test.com "
# CSS selectors work too
chrome-agent click --selector " button.submit "
chrome-agent fill --selector " input[name=email] " " hello@test.com "
# Article content (Readability -- like Firefox Reader Mode)
chrome-agent read
# Visible text, scoped and capped
chrome-agent text --selector " main " --truncate 500
# Run JS
chrome-agent eval " document.title "
# Screenshot (returns a file path, not binary)
chrome-agent screenshot
Commands
Command
What it does
goto <url> [--inspect] [--max-depth N] [--header "K: V"]
Navigate. Auto-prefixes https:// . --header (repeatable) sends extra HTTP headers.
back
History back.
forward
History forward.
close [--purge]
Stop browser. --purge deletes cookies/profile.
Inspection
Command
What it does
inspect [--verbose] [--max-depth N] [--uid nN] [--filter "role,role"] [--scroll] [--limit N] [--urls] [--max-chars N] [--offset K]
a11y tree with UIDs. --scroll --limit for infinite scroll. --urls resolves href on links. --max-chars / --offset cap and page the output.
diff
What changed since last inspect.
screenshot [--filename name] [--format jpeg|png] [--quality N] [--max-width N] [--uid nN|--selector "css"]
Screenshot to file. JPEG/quality/max-width shrink it; --uid / --selector clip to one element.
pdf [--filename name] [--landscape] [--background]
Print the current page to a PDF file.
tabs
List open tabs.
Interaction
Command
What it does
click <uid> [--inspect]
Click by uid. Falls back to JS .click() when no box model.
click --selector "css" [--inspect]
Click by CSS selector.
click --xy 100,200
Click by coordinates.
dblclick <uid> [--inspect]
Double-click by uid, --selector , or --xy .
fill --uid <uid> <value> [--inspect]
Fill input by uid.
fill --selector "css" <value>
Fill by selector.
fill-form <uid=val>...
Batch fill.
select --uid <uid> <value>
Select dropdown option by value or visible text.
select --selector "css" <value>
Select by CSS selector.
check <uid>
Ensure checkbox/radio is checked. Idempotent.
uncheck <uid>
Ensure checkbox/radio is unchecked. Idempotent.
upload --uid <uid> <file>...
Upload file(s) to a file input.
upload --selector "css" <file>...
Upload by CSS selector.
drag <from-uid> <to-uid>
Drag element to another element.
type <text> [--selector "css"]
Type into focused element.
press <key>
Enter, Tab, Escape, etc.
scroll <down|up|uid>
Scroll page or element into view.
hover <uid>
Hover.
wait <text|url|selector> <pattern>
Wait for a condition.
wait network-idle [--idle-ms N] [--timeout N]
Wait until the network is quiet for --idle-ms (default 500). Beats fixed sleeps for SPA/XHR settle.
Content extraction
Command
What it does
read [--html] [--truncate N]
Article extraction via Mozilla Readability.
text [uid] [--selector "css"] [--truncate N]
Visible text from page or element.
eval <expression> [--selector "css"]
JS in page context. el = matched element.
extract [--selector "css"] [--limit N] [--scroll] [--a11y]
Auto-detect repeating data. --a11y for React SPAs (X.com).
download <url> [--out path] [--timeout N]
Download a URL fetched in-page, so cookies/auth carry over (login-gated files). Returns {path,bytes,mime} .
Monitoring
Command
What it does
network [--filter "pattern"] [--body] [--live N] [--abort "pattern"]
Network requests and API responses. --abort blocks matching requests.
console [--level error] [--clear]
console.log/warn/error + JS exceptions.
Advanced
Command
What it does
frame <selector|main>
Switch eval / inspect into an iframe (or back to main). Persists only within a pipe / batch process.
batch
Execute multiple commands from a JSON array on stdin.
pipe
Persistent JSON stdin/stdout connection.
Global flags
--browser <name> Named browser profile (default: "default")
--page <name> Named tab (default: "default")
--connect [url] Attach to a running Chrome
--headed Show browser window (default: headless)
--stealth Anti-detection patches (Cloudflare, Turnstile)
--copy-cookies Use cookies from your real Chrome profile
--timeout <seconds> Command timeout (default: 30)
--max-depth <N> Limit inspect depth
--ignore-https-errors Accept self-signed certs
--json Structured JSON output
--dialog <mode> JS dialog policy: accept (default), dismiss, or manual
--dialog-text <text> Text to submit for prompt() dialogs when --dialog accept
JS dialogs ( alert / confirm / prompt / beforeunload ) are auto-answered by default ( --dialog accept ) — a native dialog otherwise blocks the page with no DOM signal and the agent's next command hangs. Use --dialog dismiss to cancel them, or --dialog manual to opt out.
The loop: inspect, act, inspect
chrome-agent goto https://app.com/login --inspect
# uid=n52 textbox "Email" focusable
# uid=n58 textbox "Password" focusable
# uid=n63 button "Sign In" focusable
chrome-agent fill --uid n52 " user@test.com "
chrome-agent fill --uid n58 " passw

[truncated]

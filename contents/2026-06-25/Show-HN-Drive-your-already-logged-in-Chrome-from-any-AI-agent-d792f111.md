---
source: "https://github.com/leeguooooo/chrome-use"
hn_url: "https://news.ycombinator.com/item?id=48668131"
title: "Show HN: Drive your already-logged-in Chrome from any AI agent"
article_title: "GitHub - leeguooooo/chrome-use: Browser automation CLI for AI agents · GitHub"
author: "leeguoo"
captured_at: "2026-06-25T03:03:17Z"
capture_tool: "hn-digest"
hn_id: 48668131
score: 1
comments: 0
posted_at: "2026-06-25T02:32:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Drive your already-logged-in Chrome from any AI agent

- HN: [48668131](https://news.ycombinator.com/item?id=48668131)
- Source: [github.com](https://github.com/leeguooooo/chrome-use)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T02:32:15Z

## Translation

タイトル: HN を表示: ログイン済みの Chrome を任意の AI エージェントから操作する
記事タイトル: GitHub - leeguooooo/chrome-use: AI エージェント用のブラウザ自動化 CLI · GitHub
説明: AI エージェント用のブラウザ自動化 CLI。 GitHub でアカウントを作成して、leeguooooo/chrome 使用の開発に貢献してください。

記事本文:
GitHub - leeguooooo/chrome-use: AI エージェント用のブラウザ自動化 CLI · GitHub
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
リーグォォォ
/
クロム使用
公共
vercel-labs/agent-browser からフォーク
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

イオン
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
872 コミット 872 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows .husky .husky アセット アセット bin bin cli cli docker docker docs docs 拡張機能 拡張機能 スクリプト スクリプト スキルデータ スキルデータ スキル/ クロム使用 スキル/ クロム使用 .gitignore .gitignore .node-version .node-version .nojekyll .nojekyll .prettierrc .prettierrc AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md README.zh.md README.zh.md install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
chrome-use は、AI エージェントから実際にログインしている Chrome を駆動します。これは既存のログイン セッションを共有し、実際のブラウザであるため、ボット対策システムによって検出されません。 *-use ファミリーの一部です (iphone-use は実際の iPhone を駆動し、chrome-use は実際の Chrome を駆動します)。
元々は vercel-labs/agent-browser (Apache-2.0) に基づいていました。現在はスタンドアロン プロジェクトです。ステルス/拡張リレー アーキテクチャ、検出防止、ヒューマナイズ、マルチエージェント分離、および CLI は大幅に分岐しています。
📖 詳細: エージェントがクロスオリジン iframe をクリックできるようにする — Chrome の使用がブラウザ制御の最も困難な部分をどのように解決するか
AI エージェントにすでに使用しているブラウザを提供します
新しいChromeはありません。再ログインはありません。いいえ、「あなたはロボットですか？」壁。
chrome-use は、すべてのエージェント (クロード コード、カーソル、コーデックス、独自のスクリプト) を、すでにすべてにサインインしている Chrome に向けます。ウィンドウ内でクリックされるので、その動作を観察し、2FA プロンプトまたはキャプチャが表示された瞬間にホイールを握ります。そして、これは文字通り実際のブラウザであるため (ワンクリック拡張機能、ネイティブ メッセージング - デバッグ ポートなし)、サイトはそれを 100% ヒューマとして読み取ります。

n: CreepJS のスコアは 0% bot です。
劇作家 / 人形遣い / ブラウザ使用?空のブラウザを起動するため、すべてのログインをやり直し、すべてのキャプチャと戦い、それでも自動化としてフラグが立てられます。すでにお持ちのセッションを使用します。
クロードのChrome拡張機能?素晴らしいですが、それはクロードを駆り立てるだけです。これにより、エージェントまたは CLI が駆動されます。
生の --remote-debugging-port (Web アクセスなど)? Chrome 136 以降では、「リモート デバッグを許可しますか?」というメッセージが表示されます。接続するたびに。これは決して実現しません。ワンクリックのストア拡張機能、ネイティブ メッセージングです。
¹ 3 つの実際の Chrome ツールはすべて、CreepJS (実際のブラウザです) では ~0% のスコアを獲得します。私たちのものを測定しました。 ² リブラウザの runtimeEnableLeak — リレー パス上でクリーンであることが確認されました。 Chrome のクロードは個別にテストされていません (—)。 3 Web アクセスは 1 つのブラウザ上で並列サブエージェントを実行できますが、セッションごとに分離する必要はありません。ここでの各 --session は、独自の色付きのコマンド分離タブ グループを取得します。測定された数値については、「検出防止」を参照してください。
一般的なブラウザ自動化 (Playwright、Puppeteer、または新しい --launch ) では、空のプロファイルで新しいブラウザが開きます。再度ログインする必要があり、Web サイトではそれが自動化されていることがわかります。
chrome-use は既存の Chrome に接続します。 Cookie、セッション、ブラウザーのフィンガープリントはすべて本物です。それは実際のブラウザーであるためです。
Chrome を使用する CLI は、Chrome を介して小さなブラウザ拡張機能と通信します。
ネイティブ メッセージング — ローカルのプロセス間チャネル、ネットワーク ソケットなし、なし
トークン、リモートサーバーなし。この拡張機能は chrome.debugger を使用してタブを駆動します
すでにログインしている自分の Chrome でターゲットを設定し、結果を に返します。
CLI。すべてがマシン上に残ります。
各 --session は独自の色の Chrome タブ グループを取得するため、複数のエージェントが
お互いに干渉することなく、1 つの実際のブラウザを同時に共有できます。
独自のタブ。
拡張機能を使用する理由 (生のデバッグポートではない)
その他のローカルt

ools は、生の --remote-debugging-port (CDP) 経由で Chrome を駆動します。以来
Chrome 136 では、そのような接続が行われるたびに、「リモート デバッグを許可しますか?」というブロックが表示されます。
同意ダイアログ - 事前にポートを有効にする必要があります。私たちの拡張機能では、
代わりにネイティブ メッセージング: 一度インストールすると、使用ごとの確認は不要になります。
¹ rebrowser-bot-detector に対して検証済み:
私たちのリレーは runtimeEnableLeak: 🟢 リークなし、navigatorWebdriver: 🟢 を報告します。
² CreepJS に対して検証済み
接続された実際の Chrome パス — 「検出防止」を参照してください。
同意ダイアログは架空のものではありません。raw ポート ツールは、すべてのデバイスで同意ダイアログをポップします。
アタッチします (Chrome 136 以降のセキュリティ)。拡張パスは決してそうではありません。
カール -fsSL https://raw.githubusercontent.com/leeguooooo/chrome-use/main/install.sh |しー
最新の GitHub リリースからプラットフォーム用のビルド済みバイナリをダウンロードし、chrome-use (+ abs エイリアス) をインストールします。 npmもトークンもありません。
バージョンを固定します: AGENT_BROWSER_VERSION=v0.27.0-fork.12curl -fsSL https://raw.githubusercontent.com/leeguooooo/chrome-use/main/install.sh |しー
カスタムの場所: AGENT_BROWSER_BIN_DIR=$HOME/bincurl -fsSL … |しー
Windows: リリース ページから chrome-use-win32-x64.tar.gz をダウンロードし、chrome-use.exe を PATH に置きます。
npm (レガシー): npm install -g chrome-use — まだ公開されていますが、現在は GitHub Releases がプライマリ チャネルです。
このリポジトリには、Claude Code、Cursor などの SKILL.md ファイルが同梱されています。 skill.sh を使用して、それらを現在のプロジェクトにプルします。
npx スキル追加 leeguooooo/chrome-use
これにより、 skill/chrome-use (および特殊な skill-data/{core,electron,slack,dogfood,agentcore,vercel-sandbox}) がプロジェクトにドロップされるため、AI エージェントは適切な使用パターンと、 chrome-use 、 chrome-use 、および abs に対する事前承認された bash 権限を取得します。
chrome-use 、 chrome-use 、および abs は同じバイナリです -
abs は単なる短いエイリアスです。セは無いよ

パラテ「ステルス実行可能ファイル」;ステルス
実行時の動作 (下記の検出防止を参照) が適用されます。
実際の Chrome に接続するか、 --launch を起動するかに基づいて自動的に
新鮮なもの。
推奨 — ブラウザ拡張機能 (ワンクリック、ポップアップなし)。をインストールします。
Chrome ウェブストアの chrome-use 拡張機能、
次に、ローカル ブリッジを一度登録します。
chrome-use extension install # ネイティブメッセージングホストを登録します (1 回限り)
chrome-use オープン https://x.com/home
chrome-use open は、ネイティブ上でログインしている実際の Chrome を駆動します。
メッセージング - デバッグ ポート、トークン、および「リモート デバッグを許可しますか?」はありません。ダイアログ、
今までに。拡張機能は自動更新され、Chrome を再起動しても存続するため、そのまま残ります
使用ごとの確認なしで接続されます (無人/エージェントによる使用に最適)。
拡張機能を使用しない場合、chrome-use は Chrome DevTools プロトコル経由で接続します。
これは、Chrome がリモート デバッグ ポート (
起動フラグ — chrome://inspect トグルだけでは十分ではありません):
# macOS
open -a " Google Chrome " --args --remote-debugging-port=9222
# Linux
google-chrome --remote-debugging-port=9222
# Windows: --remote-debugging-port=9222 を Chrome ショートカットのターゲットに追加します
次に、chrome-use open <url> によってポートが自動検出されます。最初のアタッチ時に、
Chrome 136 以降では、「リモート デバッグを許可しますか?」というメッセージが表示されます。ダイアログ — 「許可」を 1 回クリックします (
その Chrome セッションの間保持されます)。上記の拡張機能はこれを完全に回避します。
設定がありません/実際の Chrome を触りたくないですか?使用する
chrome-use --launch open <url> を使用して、新たに分離されたステルス ブラウザを起動します
(完全な検出防止パッチが適用されています。以下を参照してください)。これは何もしなくても常に機能します
ポート設定は、CI が自動的に使用するものです。
# Chrome に接続してナビゲートします
chrome-use https://example.com を開く
# すべてはログインしたブラウザを通じて機能します
クロムを使用して「投稿」をクリックします
クロム

e-use クリック 449 320 # …または生のビューポート座標をクリックします
chrome-use fill「タイトル」「Hello World」
クロム使用のスクリーンショット ./page.png
エージェントは Chrome で動作します。タブが開き、ページが読み込まれ、クリックがリアルタイムで発生するのがわかります。いつでも引き継ぐことができ (CAPTCHA の解決など)、その後はエージェントを続行させます。
実行中の Chrome に接続するのではなく、別のブラウザを生成します。
# 使い捨て: 新鮮な空のプロファイル — Cookie なし、ログインなし (CI/テストに適しています)
chrome-use --launch https://example.com を開きます
# ログインを保持します: 実際の Chrome プロファイルで起動します (Cookie/セッションはそのままです)。
chrome-use --launch --profile 自動オープン https://x.com/home
# または明示的に名前を付けます: --profile Default / --profile "Profile 1"
⚠️ プレーン --launch ( --profile なし) は一時的な空のプロファイルを使用します。
何もログインしないでください。ログインしたサイトの場合は、 --profile auto (
最近使用した Chrome プロファイル) または --profile <name> 。クローム使用プリント
--プロファイルなしで起動すると警告が表示されます。
CI 環境では、スタンドアロン モードが自動的に使用されます。
サイト アダプター — Web サイトを構造化データ CLI に変換します。
ほとんどの「GitHub の問題を読む」/「Reddit を検索」/「Bilibili フィードを取得する」タスクでは、
クリックしてスクリーンショットを撮る必要はまったくありません - サイトにはすでに JSON API が組み込まれています
独自のログイン。サイト アダプターは、その API を呼び出す小さな JS 関数です。
ログインしたタブ内 (Cookie、同一オリジンフェッチ、サイト独自のもの)
モジュール)、クリーンな JSON を返します。サイトはそれをあなたと区別できません。
あなたです。
chrome-use にはこれらのアダプターは含まれていません - サイトの更新によりコミュニティが取得されます
bb-sites は実行時にパックします (パッケージと同様)
依存関係を取得するマネージャー) を使用して、chrome-use のステルス トランスポート上でそれらを実行します。
chrome-use site update # アダプター パックを取得する (~14

5コマンド）
chrome利用サイトリスト # github/issues, reddit/search, bilibili/feed, …
chrome-use site info github/issues # アダプターの引数 + ドメインを参照
# 1 つ実行 — サイトに移動し (すでにサイトにいる場合はタブを再利用)、JSON を返します。
chrome 使用サイト github/issuesepral/bb-browser --json
クロム使用サイト reddit/検索「rust async」 --json
chrome-use site bilibili/feed --json # ログインセッションなので機能します
位置引数は、アダプターの宣言された引数を順番に埋めます。 --key 値の上書き
名前で。アダプターは bb-site コミュニティによって作成され、その作成者の所有物のままになります。
プロパティ — chrome-use はそれらを実行するだけです。
自動同期 + 自動提案。サイトの更新を自分で入力することはめったにありません: chrome-use
最初の使用時にパックを同期し、毎週バックグラウンドで更新します (調整
AGENT_BROWSER_SITES_TTL_DAYS 、 AGENT_BROWSER_SITES_NO_AUTO_UPDATE=1 で無効になります)。
そして、ドメインにアダプターがあるページを開いたりスナップショットしたりすると、Chrome 使用のサーフェスが表示されます。
それらは出力に直接表示されます — <domain> 行の 💡 サイト アダプター、および
--json の下の siteAdapters フィールド - エージェントが構造化データにアクセスできるようにする
DOM をスクレイピングする代わりにアダプターを使用します。
$ chrome-use オープン https://github.com
💡 github.com 用のサイト アダプター — 構造化データにはこれらを推奨します。
github/issues、github/me、github/repo、…
例:

[切り捨てられた]

## Original Extract

Browser automation CLI for AI agents. Contribute to leeguooooo/chrome-use development by creating an account on GitHub.

GitHub - leeguooooo/chrome-use: Browser automation CLI for AI agents · GitHub
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
leeguooooo
/
chrome-use
Public
forked from vercel-labs/agent-browser
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
872 Commits 872 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows .husky .husky assets assets bin bin cli cli docker docker docs docs extensions extensions scripts scripts skill-data skill-data skills/ chrome-use skills/ chrome-use .gitignore .gitignore .node-version .node-version .nojekyll .nojekyll .prettierrc .prettierrc AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md README.zh.md README.zh.md install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
chrome-use drives your real, logged-in Chrome from any AI agent — it shares your existing login sessions and is undetectable by anti-bot systems because it is your real browser. Part of the *-use family ( iphone-use drives your real iPhone; chrome-use drives your real Chrome).
Originally based on vercel-labs/agent-browser (Apache-2.0); now a standalone project — the stealth/extension-relay architecture, anti-detection, humanize, multi-agent isolation, and CLI have diverged substantially.
📖 Deep dive: Letting an agent click into cross-origin iframes — how chrome-use solves the hardest part of browser control
Give your AI agent the browser you already live in
No fresh Chrome. No re-login. No "are you a robot?" walls.
chrome-use points any agent — Claude Code, Cursor, Codex, your own scripts — at the Chrome you're already signed into everything on . It clicks in your window, so you watch it work and grab the wheel the moment it hits a 2FA prompt or captcha. And because it's literally your real browser (over a one-click extension, native messaging — no debug port), sites read it as 100% human: CreepJS scores it 0% bot .
Playwright / Puppeteer / browser-use? They boot an empty browser — so you redo every login, fight every captcha, and still get flagged as automation. We use the session you already have.
Claude's Chrome extension? Great, but it only drives Claude. This drives any agent or CLI.
A raw --remote-debugging-port (web-access, etc.)? Chrome 136+ pops "Allow remote debugging?" on every connect. This never does — one-click Store extension, native messaging.
¹ All three real-Chrome tools score ~0% on CreepJS (it's a real browser); we've measured ours. ² rebrowser's runtimeEnableLeak — verified clean on our relay path; Claude in Chrome not independently tested (—). ³ web-access can run parallel sub-agents on one browser, but without per-session isolation; each --session here gets its own colored, command-isolated tab group. See Anti-detection for the measured numbers.
Typical browser automation (Playwright, Puppeteer, or a fresh --launch ) opens a brand-new browser with an empty profile. You have to log in again, and websites can tell it's automated.
chrome-use connects to your existing Chrome. Your cookies, sessions, and browser fingerprint are all real — because it IS your real browser.
Your chrome-use CLI talks to a tiny browser extension over Chrome
native messaging — a local inter-process channel, no network socket, no
token, no remote server . The extension uses chrome.debugger to drive the tabs
you target in your own, already-logged-in Chrome , then hands results back to
the CLI. Everything stays on your machine.
Each --session gets its own colored Chrome tab group , so multiple agents
can share one real browser concurrently without stepping on each other — or your
own tabs.
Why the extension (not a raw debug port)
Other local tools drive Chrome over a raw --remote-debugging-port (CDP). Since
Chrome 136 , every such connection pops a blocking "Allow remote debugging?"
consent dialog — and the port has to be enabled up front. Our extension uses
native messaging instead: install once, then zero per-use confirmation.
¹ Verified against rebrowser-bot-detector :
our relay reports runtimeEnableLeak: 🟢 No leak and navigatorWebdriver: 🟢 .
² Verified against CreepJS on the
connected real-Chrome path — see Anti-detection .
The consent dialog isn't hypothetical: a raw-port tool pops it on every
attach (Chrome 136+ security). The extension path never does.
curl -fsSL https://raw.githubusercontent.com/leeguooooo/chrome-use/main/install.sh | sh
Downloads the prebuilt binary for your platform from the latest GitHub Release and installs chrome-use (+ the abs alias). No npm, no tokens.
Pin a version: AGENT_BROWSER_VERSION=v0.27.0-fork.12 curl -fsSL https://raw.githubusercontent.com/leeguooooo/chrome-use/main/install.sh | sh
Custom location: AGENT_BROWSER_BIN_DIR=$HOME/bin curl -fsSL … | sh
Windows: download chrome-use-win32-x64.tar.gz from the Releases page and put chrome-use.exe on your PATH.
npm (legacy): npm install -g chrome-use — still published, but GitHub Releases is the primary channel now.
The repo ships SKILL.md files for Claude Code, Cursor, etc. Pull them into the current project with skills.sh :
npx skills add leeguooooo/chrome-use
This drops skills/chrome-use (and the specialized skill-data/{core,electron,slack,dogfood,agentcore,vercel-sandbox} ) into your project so your AI agent gets the right usage patterns and pre-approved bash permissions for chrome-use , chrome-use , and abs .
chrome-use , chrome-use , and abs are the same binary —
abs is just a short alias. There is no separate "stealth executable"; stealth
is a runtime behavior (see Anti-detection below), applied
automatically based on whether you attach to your real Chrome or --launch a
fresh one.
Recommended — the browser extension (one click, no popups). Install the
chrome-use extension from the Chrome Web Store ,
then register the local bridge once:
chrome-use extension install # register the native-messaging host (one-time)
chrome-use open https://x.com/home
chrome-use open then drives your real, logged-in Chrome over native
messaging — no debug port, no token, and no "Allow remote debugging?" dialog,
ever . The extension auto-updates and survives Chrome restarts, so it stays
connected with zero per-use confirmation (ideal for unattended/agent use).
Without the extension, chrome-use attaches over the Chrome DevTools Protocol,
which Chrome only exposes when launched with a remote-debugging port (a
startup flag — the chrome://inspect toggle alone is not enough):
# macOS
open -a " Google Chrome " --args --remote-debugging-port=9222
# Linux
google-chrome --remote-debugging-port=9222
# Windows: add --remote-debugging-port=9222 to your Chrome shortcut's target
Then chrome-use open <url> auto-discovers the port. On first attach,
Chrome 136+ shows an "Allow remote debugging?" dialog — click Allow once (it
persists for that Chrome session). The extension above avoids this entirely.
No setup / don't want to touch your real Chrome? Use
chrome-use --launch open <url> to spawn a fresh isolated stealth browser
(full anti-detection patches applied; see below). This always works without any
port setup and is what CI uses automatically.
# Connect to your Chrome and navigate
chrome-use open https://example.com
# Everything works through your logged-in browser
chrome-use click " Post "
chrome-use click 449 320 # …or click a raw viewport coordinate
chrome-use fill " Title " " Hello World "
chrome-use screenshot ./page.png
The agent operates in your Chrome — you'll see tabs opening, pages loading, clicks happening in real time. You can take over at any point (e.g. solve a CAPTCHA), then let the agent continue.
Spawn a separate browser instead of attaching to your running Chrome:
# Throwaway: fresh, EMPTY profile — no cookies, no login (good for CI/testing)
chrome-use --launch open https://example.com
# Keep your login: launch with your real Chrome profile (cookies/sessions intact)
chrome-use --launch --profile auto open https://x.com/home
# or name it explicitly: --profile Default / --profile "Profile 1"
⚠️ Plain --launch (no --profile ) uses a temporary empty profile — you will
NOT be logged into anything. For logged-in sites use --profile auto (picks the
Chrome profile you used most recently) or --profile <name> . chrome-use prints
a warning when you --launch without a profile.
In CI environments, standalone mode is used automatically.
Site adapters — turn a website into a structured-data CLI
Most "read GitHub issues" / "search Reddit" / "get my Bilibili feed" tasks don't
need clicking and screenshotting at all — the site already has a JSON API behind
its own login. A site adapter is a tiny JS function that calls that API from
inside your logged-in tab (your cookies, same-origin fetch , the site's own
modules) and returns clean JSON. The site can't tell it apart from you, because it
is you.
chrome-use ships none of these adapters — site update fetches the community
bb-sites pack at runtime (like a package
manager pulling a dependency), then runs them over chrome-use's stealth transport:
chrome-use site update # fetch the adapter pack (~145 commands)
chrome-use site list # github/issues, reddit/search, bilibili/feed, …
chrome-use site info github/issues # see an adapter's args + domain
# Run one — navigates to the site (reusing the tab if you're already there) and returns JSON
chrome-use site github/issues epiral/bb-browser --json
chrome-use site reddit/search " rust async " --json
chrome-use site bilibili/feed --json # works because it's your logged-in session
Positional args fill the adapter's declared args in order; --key value overrides
by name. Adapters are authored by the bb-sites community and remain their authors'
property — chrome-use just runs them.
Auto-sync + auto-suggest. You rarely type site update yourself: chrome-use
syncs the pack on first use and refreshes it weekly in the background (tune with
AGENT_BROWSER_SITES_TTL_DAYS , disable with AGENT_BROWSER_SITES_NO_AUTO_UPDATE=1 ).
And when you open / snapshot a page whose domain has adapters, chrome-use surfaces
them right in the output — a 💡 site adapters for <domain> line, plus a
siteAdapters field under --json — so an agent reaches for the structured-data
adapter instead of scraping the DOM:
$ chrome-use open https://github.com
💡 site adapters for github.com — prefer these for structured data:
github/issues, github/me, github/repo, …
e.g.

[truncated]

---
source: "https://github.com/disi910/claude-usage-bar"
hn_url: "https://news.ycombinator.com/item?id=49010415"
title: "Show HN: Chrome Extension Claude Token Usage Bar and Context Use for Claude.ai"
article_title: "GitHub - disi910/claude-usage-bar: Chrome extension that shows your Claude.ai 5-hour usage as a bar at the top of every page. · GitHub"
author: "didi2121"
captured_at: "2026-07-22T18:03:31Z"
capture_tool: "hn-digest"
hn_id: 49010415
score: 1
comments: 0
posted_at: "2026-07-22T17:34:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Chrome Extension Claude Token Usage Bar and Context Use for Claude.ai

- HN: [49010415](https://news.ycombinator.com/item?id=49010415)
- Source: [github.com](https://github.com/disi910/claude-usage-bar)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T17:34:46Z

## Translation

タイトル: HN を表示: Claude.ai の Chrome 拡張機能クロード トークン使用バーとコンテキストの使用
記事のタイトル: GitHub - disi910/claude-usage-bar: Claude.ai の 5 時間の使用状況を各ページの上部にバーとして表示する Chrome 拡張機能。 · GitHub
説明: Claude.ai の 5 時間の使用状況を各ページの上部にバーとして表示する Chrome 拡張機能。 - disi910/claude-usage-bar

記事本文:
GitHub - disi910/claude-usage-bar: Claude.ai の 5 時間の使用状況を各ページの上部にバーとして表示する Chrome 拡張機能。 · GitHub
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
ディシ910
/
クロード使用状況バー
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
32 コミット 32 コミット _locales _locales スクリーンショット スクリーンショット スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md 背景.js 背景.js content.js content.js icon128.png icon128.png icon16.png icon16.png icon48.png icon48.png Index.htmlインデックス.html マニフェスト.json マニフェスト.json ポップアップ.css ポップアップ.css ポップアップ.html ポップアップ.html ポップアップ.js ポップアップ.js プライバシー-ポリシー.html プライバシー-ポリシー.html スタイル.css スタイル.css すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Usage Bar for Claude は、Claude プランの使用制限 (5 時間制限、週制限、追加クレジット)、コンテキスト ウィンドウのライブ トークン カウンター、およびプロンプト キャッシュ カウントダウンを claude.ai 上で直接表示する、無料のオープンソース Chrome 拡張機能です。アカウント、分析、外部サーバーは必要ありません。Claude 設定ページが使用するものと同じ使用状況データが完全にブラウザ内で読み取られます。
Chrome ウェブストアからインストールします。 · 14 言語に対応。 · MIT ライセンスを取得しています。 · Anthropic とは提携していません。
使用状況バー — クロードの 5 時間制限をパーセンテージで示し、リセット カウントダウンが表示されます。ページの上部に固定されるか、チャット ボックス内 (ポップアップで選択) に細い線として表示されます。クロードのコーラルパレットとダーク/ライトテーマを反映しています。
ホバー上の計画パネル — 4 つの行: 5 時間制限、毎週、すべてのモデル、追加クレジット、ルーチン、それぞれに使用率とリセット時間が表示されます。
コンテキスト ウィンドウ ドーナツ — 現在のチャットが占めるコンテキスト ウィンドウ (デフォルトでは 200,000 トークン) の量を推定する循環トークン カウンターで、現在のブランチ コンテキストと使用された合計トークンを区別します。
プロンプト キャッシュの砂時計 — 各アシスタントが応答した後、約 5 分間のプロンプト キャッシュ TTL をカウントダウンするため、フォローアップがいつ行われるかを知ることができます。

安価なキャッシュにヒットするまで。
14 言語 — 英語、スペイン語、ポルトガル語 (BR/PT)、フランス語、ドイツ語、イタリア語、Русский、简体中文、繁體中文、日本語、한국어、हिन्दी、テュルクチェ語。ブラウザの言語に自動的に準拠します。
クロードの使用制限を確認するにはどうすればよいですか?
Claude 用の使用状況バーをインストールし、 claude.ai を開きます。バーには、5 時間の制限使用量と、リセットされるまでの時間がパーセンテージで表示されます。カーソルを合わせると、毎週の全モデルの制限、追加クレジット、ルーティンの使用状況が表示されます。データはクロード自身の使用状況エンドポイントから取得されます。会話を離れることなく、[設定] → [使用状況] と同じ数値が取得されます。
クロードの 5 時間制限がどのくらい残っているかを確認するにはどうすればよいですか?
メインバーはまさにその通りです。使用されたローリング 5 時間制限の割合と、ライブの「…でリセット」カウントダウンが表示されます。 100% に達するとバーが赤くなり、クロードがメッセージを拒否し始める前にバーが表示されます。
クロードにはトークンカウンターがありますか?
Claude.ai は、限られたトークン情報をネイティブに表示します。クロードの使用状況バーは、永続的なコンテキスト ウィンドウのドーナツを追加します。カーソルを置くと、使用された推定トークン、消費されたコンテキストの長さ (例: 34k / 200k)、および編集されたブランチを含む会話全体の合計トークンが表示されます。
はい。この拡張機能は、既存のセッションを使用して、claude.ai に対して同一オリジンのリクエストのみを行います。他には何も送信されません。分析、テレメトリ、サードパーティのサーバーはありません。会話テキストはトークンを推定するためにのみメモリ内で処理されます。保存または送信されることはありません。プライバシー ポリシーを参照してください。
Claude Pro、Max、Team、および無料プランで動作しますか?
クロードの使用状況 API がアカウントに対して返す制限行がすべて表示されます。プランにない行 (追加クレジットなど) は — として表示されます。
使用状況バーを表示 — ウィジェットのオンとオフを切り替えます。
位置 — ページの先頭 (フローティング オーバーレイ) またはチャット ボックス内 (com)

コンポーザーのツールバー行内の協定バー)。チャット ボックスが見つからない場合、バーは上部の配置に戻ります。
chrome://extensions を開き、開発者モードを有効にします。
「解凍してロード」をクリックし、フォルダーを選択します。
プランの使用量 : クロードの使用量設定ページが使用するのと同じ内部エンドポイントを 30 秒ごとにポーリングします。カウントダウンは 1 秒ごとに刻みます (タブが非表示になっている間は一時停止します)。
コンテキストドーナツ : 2 層のトークンカウンター。まず、表示されている会話からインスタント ヒューリスティックを描画し (スクリプト対応: スペース言語の場合は単語あたり最大 1.3 トークン、CJK の場合は文字あたり最大 1 トークン)、次に完全な会話ペイロード (ツール呼び出し、添付ファイル、DOM では表示されない思考ブロックを含む) を使用してアップグレードし、会話ツリーのアクティブなブランチをたどるので、編集/再試行によってカウントが膨らみません。
キャッシュ タイマー : 新しいアシスタント メッセージが到着するたびに、ローカル 5 分間タイマーが再起動されます。
すべてがブラウザ内で実行されます。分析もサーバーもサードパーティもありません。
トークン数は推定値であり、トークナイザーの正確な値ではありません。 claude.ai システム プロンプトとツール定義は拡張機能には表示されません。
複数組織アカウントは、API が返す最初の組織を使用します。
文書化されていない claude.ai エンドポイントを使用します。 Anthropic がそれらを変更すると、更新が適用されるまでバーの一部が空白になります。
プルリクエストは大歓迎です。微調整よりも大きな問題については、まず問題を解決してください。Firefox への移植と組織スイッチャーが次のステップとして適切です。
翻訳を追加または変更するには、scripts/build_locales.py を編集して再実行します。すべての _locales/<locale>/messages.json を再生成します。
manifest.json 拡張構成 (_locales 経由でローカライズ)
content.js バー マウント (上部オーバーレイまたはコンポーザー内) + オブザーバー + レンダリング
インストール時のbackground.jsシードのデフォルト設定
Popup.html/css/js 設定ポップアップ (表示/非表示、位置、レビューリンク)
s

tyles.css バー、パネル、ドーナツ、砂時計、インライン モード スタイル
_locales/ 生成された翻訳 (14 言語)
scripts/build_locales.py 信頼できる翻訳ソース — _locales/ を再生成します
icon{16,48,128}.png ツールバーとストアのアイコン
index.html ランディング ページ (GitHub ページ)
Privacy-policy.html は GitHub ページ経由で提供され、ストアの掲載情報にリンクされています
ライセンス
Claude.ai の 5 時間の使用状況を各ページの上部にバーとして表示する Chrome 拡張機能。
chromewebstore.google.com/detail/imblbfhdbdecholhjbagcjahdkhidneb
トピックス
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

Chrome extension that shows your Claude.ai 5-hour usage as a bar at the top of every page. - disi910/claude-usage-bar

GitHub - disi910/claude-usage-bar: Chrome extension that shows your Claude.ai 5-hour usage as a bar at the top of every page. · GitHub
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
disi910
/
claude-usage-bar
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
32 Commits 32 Commits _locales _locales screenshots screenshots scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md background.js background.js content.js content.js icon128.png icon128.png icon16.png icon16.png icon48.png icon48.png index.html index.html manifest.json manifest.json popup.css popup.css popup.html popup.html popup.js popup.js privacy-policy.html privacy-policy.html styles.css styles.css View all files Repository files navigation
Usage Bar for Claude is a free, open-source Chrome extension that shows your Claude plan usage limits (5-hour limit, weekly limit, extra credits), a live token counter for the context window, and a prompt-cache countdown — directly on claude.ai. No account, no analytics, no external servers: it reads the same usage data the Claude settings page uses, entirely inside your browser.
Install from the Chrome Web Store · 14 languages · MIT licensed · Not affiliated with Anthropic.
Usage bar — your Claude 5-hour limit as a percentage with reset countdown, pinned to the top of the page or rendered as a slim line inside the chat box (pick in the popup). Mirrors Claude's coral palette and dark/light themes.
Plan panel on hover — four rows: 5-hour limit, Weekly · all models, Extra credits, Routines, each with percent used and reset time.
Context-window donut — a circular token counter estimating how much of the context window (200k tokens by default) the current chat occupies, distinguishing current-branch context from total tokens used.
Prompt-cache hourglass — counts down the ~5-minute prompt-cache TTL after each assistant reply, so you know when a follow-up still hits the cheap cache.
14 languages — English, Español, Português (BR/PT), Français, Deutsch, Italiano, Русский, 简体中文, 繁體中文, 日本語, 한국어, हिन्दी, Türkçe. Follows your browser language automatically.
How do I see my Claude usage limits?
Install Usage Bar for Claude and open claude.ai . The bar shows your 5-hour limit usage as a percentage plus the time until it resets; hover it for the weekly all-models limit, extra credits, and Routines usage. The data comes from Claude's own usage endpoint — the same numbers as Settings → Usage, without leaving the conversation.
How do I check how much of my Claude 5-hour limit is left?
The main bar is exactly that: percent of the rolling 5-hour limit used, with a live "resets in …" countdown. When it hits 100% the bar turns red so you see it before Claude starts declining messages.
Does Claude have a token counter?
Claude.ai shows limited token information natively. Usage Bar for Claude adds a persistent context-window donut: hover it to see estimated tokens used, context length consumed (e.g. 34k / 200k), and total tokens across the conversation, including edited branches.
Yes. The extension makes only same-origin requests to claude.ai using your existing session. Nothing is sent anywhere else — no analytics, no telemetry, no third-party servers. Conversation text is processed in memory only to estimate tokens; it is never stored or transmitted. See the privacy policy .
Does it work with Claude Pro, Max, Team and free plans?
It displays whatever limit rows Claude's usage API returns for your account. Rows your plan doesn't have (e.g. Extra credits) show as — .
Show usage bar — toggle the widget on or off.
Position — Top of page (floating overlay) or In chat box (compact bar inside the composer's toolbar row). If the chat box can't be found, the bar falls back to the top placement.
Open chrome://extensions , enable Developer mode.
Click Load unpacked and select the folder.
Plan usage : polls the same internal endpoint the Claude usage settings page uses, every 30 seconds; countdowns tick every second (paused while the tab is hidden).
Context donut : a two-tier token counter. It first paints an instant heuristic from the visible conversation (script-aware: ~1.3 tokens per word for spaced languages, ~1 token per character for CJK), then upgrades using the full conversation payload — including tool calls, attachments, and thinking blocks the DOM never shows — and follows the active branch of the conversation tree so edits/retries don't inflate the count.
Cache timer : a local 5-minute timer restarted whenever a new assistant message lands.
Everything runs inside your browser. No analytics, no server, no third party.
Token counts are estimates, not tokenizer-exact; the claude.ai system prompt and tool definitions aren't visible to the extension.
Multi-org accounts use the first organization the API returns.
Uses undocumented claude.ai endpoints. If Anthropic changes them, parts of the bar go blank until an update lands.
Pull requests welcome. Open an issue first for anything bigger than a tweak — a Firefox port and an org switcher are reasonable next steps.
To add or change a translation, edit scripts/build_locales.py and re-run it; it regenerates every _locales/<locale>/messages.json .
manifest.json extension config (localized via _locales)
content.js bar mount (top overlay or in-composer) + observers + rendering
background.js seeds default settings on install
popup.html/css/js settings popup (show/hide, position, review link)
styles.css bar, panel, donut, hourglass, inline-mode styling
_locales/ generated translations (14 languages)
scripts/build_locales.py translation source of truth — regenerates _locales/
icon{16,48,128}.png toolbar and store icons
index.html landing page (GitHub Pages)
privacy-policy.html served via GitHub Pages, linked in the store listing
License
Chrome extension that shows your Claude.ai 5-hour usage as a bar at the top of every page.
chromewebstore.google.com/detail/imblbfhdbdecholhjbagcjahdkhidneb
Topics
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

---
source: "https://github.com/cytostack/openwolf"
hn_url: "https://news.ycombinator.com/item?id=48482190"
title: "Sharper context. Fewer tokens. Open-source middleware for Claude Code"
article_title: "GitHub - cytostack/openwolf: Sharper context. Fewer tokens. Open-source middleware for Claude Code. · GitHub"
author: "momentmaker"
captured_at: "2026-06-10T21:44:55Z"
capture_tool: "hn-digest"
hn_id: 48482190
score: 1
comments: 0
posted_at: "2026-06-10T20:28:12Z"
tags:
  - hacker-news
  - translated
---

# Sharper context. Fewer tokens. Open-source middleware for Claude Code

- HN: [48482190](https://news.ycombinator.com/item?id=48482190)
- Source: [github.com](https://github.com/cytostack/openwolf)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T20:28:12Z

## Translation

タイトル: より鮮明なコンテキスト。トークンが少なくなります。 Claude Code 用のオープンソース ミドルウェア
記事のタイトル: GitHub - cytostack/openwolf: より鮮明なコンテキスト。トークンが少なくなります。 Claude Code 用のオープンソース ミドルウェア。 · GitHub
説明: より鮮明なコンテキスト。トークンが少なくなります。 Claude Code 用のオープンソース ミドルウェア。 - サイトスタック/オープンウルフ

記事本文:
GitHub - cytostack/openwolf: より鮮明なコンテキスト。トークンが少なくなります。 Claude Code 用のオープンソース ミドルウェア。 · GitHub
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
サイトスタック
/
オープンウルフ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット .github/ workflows .github/ workflows bin bin docs docs src src .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.mdデモ.gif デモ.gif openwolf-icon.png openwolf-icon.png package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.hooks.json tsconfig.hooks.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード・コードの第二の頭脳。
6 つのフック スクリプトによるプロジェクト インテリジェンス、トークン追跡、目に見えない強制。ワークフローの変更はゼロです。
クロード コードは強力ですが、盲目で機能します。ファイルを開くまで、ファイルに何が含まれているかわかりません。 50 トークンの構成と 2,000 トークンのモジュールを区別することはできません。気付かずに 1 つのセッションで同じファイルを複数回読み取ります。プロジェクトのインデックス、修正内容の記憶、すでに試行した内容の認識はありません。
OpenWolf はクロードに第 2 の脳を与えます。ファイルを読み取る前にファイルに何が含まれているかを知るためのファイル インデックス、好みや過去の間違いを蓄積する学習記憶、すべてを追跡するトークン台帳です。すべてのクロードのアクションで起動する 6 つの目に見えないフック スクリプト全体を通して。
大規模なアクティブなプロジェクトでテストされました。同じコードベース、同じプロンプト、異なるセットアップ:
OpenClaw + クロード ██████████████████████████████████████ ~340 万トークン
クロード CLI (OpenWolf なし) ████████████████████████████████ ~250 万トークン
OpenWolf + クロード CLI ████████ ~425K トークン
OpenWolf は、同じプロジェクトの裸の Claude CLI と比較して、トークンを最大 80% 節約しました。
20のプロジェクトにわたって、

132 を超えるセッション: 平均トークン削減率は 65.8% で、繰り返されるファイル読み取りの 71% が捕捉されブロックされました。これらはベンチマークではなく、実際の使用状況からの数値です。結果はプロジェクトの規模と使用パターンによって異なります。
npm install -g openwolf
プロジェクトの cd
オープンウルフ初期化
それだけです。クロードは普通に使ってください。 OpenWolfは監視しています。
openwolf init は、プロジェクト内に .wolf/ ディレクトリを作成します。
クロードがファイルを読み取る前に、OpenWolf はファイルの内容とそのサイズをクロードに伝えます。クロードがこのセッションですでにそのファイルを読んでいる場合、OpenWolf は警告します。 Claude がコードを書く前に、OpenWolf は cerebrum.md に既知の間違いがないかチェックします。書き込みが完了するたびに、プロジェクト マップが自動更新され、トークンの使用状況が記録されます。これはどれもわかりません。それはただ起こるのです。
メッセージを入力します
↓
クロードはファイルを読むことにしました
↓
OpenWolf: 「anatomy.md によると、このファイルは ~380 トークンであるとのことです。説明: メイン エントリ ポイント。」
↓
クロードはファイルを読みます
↓
OpenWolf: 読み取りをログに記録し、トークンを推定し、繰り返し読み取りをチェックします。
↓
クロードはコードを書きます
↓
OpenWolf: cerebrum.md をチェックします - 一致する既知の間違いはありません
↓
クロードは終了します
↓
OpenWolf: anatomy.md を更新し、memory.md に追加し、トークン台帳を更新します。
.wolf/ ファイル
cerebrum.md - 学習記憶
クロードは、ユーザーがファイルを修正したり、好みを表明したり、決定を下したりすると、このファイルを更新します。 Do-Not-Repeat リストは、セッション間で同じ間違いを防止します。
## 繰り返さないでください
- 2026-03-10: ` var ` は決して使用しないでください - 常に ` const ` または ` let ` を使用してください
- 2026-03-11: 統合テストでデータベースをモックしないでください - 実際の接続を使用してください
- 2026-03-14: 認証ミドルウェアは ` cfg.tts ` ではなく ` cfg.talk ` から読み取ります - 2 回焼き付けられました
## ユーザー設定
- クラスコンポーネントよりも機能コンポーネントを優先します
- 常に名前付きエクスポートを使用し、デフォルトのエクスポートは使用しないでください。
- テストはソースファイルの隣の` __tests__/ `に入れます
## 主要な学習

イングス
- このプロジェクトは、厳密なホイスティングを備えた pnpm ワークスペースを使用します
- API レート リミッターは、固定バケットではなく、スライディング ウィンドウを使用します。
- 認証ミドルウェアは構成ファイルではなく env.JWT_SECRET から読み取ります
anatomy.md - プロジェクト マップ
すべてのファイルには説明とトークン推定値が付けられます。クロードは概要だけで十分な場合、ファイルを開く代わりにこれを読みます。
## ソース/
- `index.ts` - メインエントリポイント。 CLI 用に createProgram() をエクスポートします。 (〜180トク)
- `server.ts` - ミドルウェアチェーンを備えた Express HTTP サーバー。 (〜520トーク)
## ソース/API/
- ` auth.ts ` - JWT 検証ミドルウェア。 env.JWT_SECRET から読み取ります。 (~340トク)
- ` users.ts ` - /api/users の CRUD エンドポイント。クエリパラメータによるページネーション。 (~890トク)
token-ledger.json - 領収書
すべてのセッションがラインアイテムを取得します。有効期間の合計により、OpenWolf が実際にトークンを節約しているかどうかがわかります。
{
「生涯」: {
"total_tokens_estimated" : 503978 、
"total_reads" : 287 ,
"total_writes" : 269 ,
"合計セッション数" : 15 、
"解剖学_ヒット" : 198 、
"解剖学ミス" : 89 、
"repeat_reads_blocked" : 106 、
「推定節約額_vs_bare_cli」: 2066959
}
}
buglog.json - バグの記憶
クロードは何かを修正する前に、その修正がすでに知られているかどうかを確認します。修正後、解決策をログに記録します。
{
"id" : " バグ-012 " 、
"error_message" : " TypeError: 未定義のプロパティを読み取れません ('map' の読み取り) " ,
"ファイル" : " src/components/UserList.tsx " ,
"root_cause" : " ユーザー配列が予期されていたときに API 応答が null でした " ,
"fix" : " オプションのチェーンを追加しました: data?.users?.map() とフォールバックの空の配列 " ,
"タグ" : [ " null チェック " 、 " API 応答 " 、 " 反応 " ]
}
コマンド
openwolf init .wolf/ を初期化し、フックを登録します
openwolf status 健全性、統計、ファイルの整合性を表示します
openwolf scan プロジェクト構造マップを更新する
openwolf scan --check 分析構造がファイルシステムと一致することを確認します (古い場合は 1 を終了します)
オペ

nwolf ダッシュボード リアルタイム Web ダッシュボードを開く
openwolf daemon start バックグラウンド タスク スケジューラを開始します
openwolf daemon stop スケジューラを停止します
openwolf デーモンの再起動 スケジューラを再起動します
openwolf デーモンのログ スケジューラーのログを表示する
openwolf cron リスト スケジュールされたタスクをすべて表示
openwolf cron run <id> 手動でタスクをトリガーする
openwolf cron retry <id> 配信不能になったタスクを再試行します
openwolf designqc デザイン評価のために全ページのスクリーンショットをキャプチャします
openwolf bug search <用語> 既知の修正のバグ メモリを検索します。
openwolf update 登録されているすべてのプロジェクトを最新バージョンに更新します
openwolf stop [backup] タイムスタンプ付きバックアップから .wolf/ を復元します。
設計QC
実行中のアプリの全ページのスクリーンショットをキャプチャし、Claude にデザインを評価してもらいます。
オープンウルフデザインqc
開発サーバーを自動検出し、すべてのルートのビューポート高さの JPEG セクションをキャプチャし、 .wolf/designqc-captures/ に保存します。次に、クロードにスクリーンショットを読んで評価するように伝えます。 puppeteer-core が必要です。
クロードに UI フレームワークの選択を手伝ってもらいます。 OpenWolf には、12 のフレームワーク (shadcn/ui、Aceternity、Magic UI、DaisyUI、HeroUI、Chakra、Flowbite、Preline、Park UI、Origin UI、Headless UI、Cult UI) の厳選されたナレッジ ベースが、実戦テストされた移行プロンプトとともに同梱されています。クロードは .wolf/reframe-frameworks.md を読み、いくつかの質問をし、プロジェクトに適切なプロンプトを表示して移行を実行します。
OpenWolf は AI ラッパーではありません。それは 6 つのフック スクリプトと .wolf/ ディレクトリです。 AI を代わりに実行したり、ワークフローを変更したりすることはありません。これは、Claude Code に欠けているものを提供します。プロジェクト マップにより読み取りが少なくなり、メモリにより学習が速くなり、台帳によりトークンの行き先がわかります。
オプション: 永続的なバックグラウンド タスク用の PM2
オプション: デザイン QC 用の puppeteer-core スクリーンショット
Claude Code フックは比較的新しい機能です。 OpenWolf は CLAUD にフォールバックします

フックが起動しない場合の E.md 命令。
トークンの追跡は、正確な API 数ではなく、推定 (文字とトークンの比率) に基づいています。精度は最大 15% 以内です。
cerebrum.md は、クロードが修正後に更新する指示に従うことに依存しています。コンプライアンスは 85 ～ 90% であり、100% ではありません。
これはv1.0.4です。物が壊れるかもしれません。ファイルの問題。
Cytostack で Claude Code を使用して製品を構築していたときに、何かがおかしいことに気づきました。セッションはトークンを必要以上に早く消費していました。私たちが調査したところ、クロードは同じファイルを何度も再読み込みし、1 つの関数を見つけるためにディレクトリ全体をスキャンし、ファイルを開かない限りファイルに何が含まれているかを知る方法がないことがわかりました。プロジェクト マップ、読み取り認識、トークンの可視性はありませんでした。そこで私たちは、クロードが読み込む回数を減らすためのファイル インデックス、より賢くなるための学習メモリ、そしてすべてのトークンを追跡する台帳など、存在してほしかったツールを構築しました。それが OpenWolf になりました。
Farhan Palathinkal Afsal によって構築されました - Cytostack
より鮮明なコンテキスト。トークンが少なくなります。 Claude Code 用のオープンソース ミドルウェア。
AGPL-3.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
155
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Sharper context. Fewer tokens. Open-source middleware for Claude Code. - cytostack/openwolf

GitHub - cytostack/openwolf: Sharper context. Fewer tokens. Open-source middleware for Claude Code. · GitHub
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
cytostack
/
openwolf
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits .github/ workflows .github/ workflows bin bin docs docs src src .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md demo.gif demo.gif openwolf-icon.png openwolf-icon.png package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.hooks.json tsconfig.hooks.json tsconfig.json tsconfig.json View all files Repository files navigation
A second brain for Claude Code.
Project intelligence, token tracking, and invisible enforcement through 6 hook scripts. Zero workflow changes.
Claude Code is powerful but it works blind. It doesn't know what a file contains until it opens it. It can't tell a 50-token config from a 2,000-token module. It reads the same file multiple times in one session without noticing. It has no index of your project, no memory of your corrections, and no awareness of what it already tried.
OpenWolf gives Claude a second brain: a file index so it knows what files contain before reading them, a learning memory that accumulates your preferences and past mistakes, and a token ledger that tracks everything. All through 6 invisible hook scripts that fire on every Claude action.
Tested on a large active project. Same codebase, same prompts, different setups:
OpenClaw + Claude ██████████████████████████████████████ ~3.4M tokens
Claude CLI (no OpenWolf) ████████████████████████████████ ~2.5M tokens
OpenWolf + Claude CLI ████████ ~425K tokens
OpenWolf saved ~80% of tokens compared to bare Claude CLI on the same project.
Across 20 projects, 132+ sessions: average token reduction of 65.8%, with 71% of repeated file reads caught and blocked. These are numbers from real usage, not benchmarks. Your results will vary by project size and usage patterns.
npm install -g openwolf
cd your-project
openwolf init
That's it. Use claude normally. OpenWolf is watching.
openwolf init creates a .wolf/ directory in your project:
Before Claude reads a file, OpenWolf tells it what the file contains and how large it is. If Claude already read that file this session, OpenWolf warns it. Before Claude writes code, OpenWolf checks your cerebrum.md for known mistakes. After every write, it auto-updates the project map and logs token usage. You see none of this. It just happens.
You type a message
↓
Claude decides to read a file
↓
OpenWolf: "anatomy.md says this file is ~380 tokens. Description: Main entry point."
↓
Claude reads the file
↓
OpenWolf: logs the read, estimates tokens, checks for repeated reads
↓
Claude writes code
↓
OpenWolf: checks cerebrum.md - no known mistakes matched
↓
Claude finishes
↓
OpenWolf: updates anatomy.md, appends to memory.md, updates token ledger
The .wolf/ Files
cerebrum.md - the learning memory
Claude updates this file when you correct it, express a preference, or make a decision. The Do-Not-Repeat list prevents the same mistake across sessions.
## Do-Not-Repeat
- 2026-03-10: Never use ` var ` - always ` const ` or ` let `
- 2026-03-11: Don't mock the database in integration tests - use the real connection
- 2026-03-14: The auth middleware reads from ` cfg.talk ` , not ` cfg.tts ` - got burned twice
## User Preferences
- Prefers functional components over class components
- Always use named exports, never default exports
- Tests go in ` __tests__/ ` next to the source file
## Key Learnings
- This project uses pnpm workspaces with strict hoisting
- The API rate limiter uses a sliding window, not fixed buckets
- Auth middleware reads from env.JWT_SECRET, not config file
anatomy.md - the project map
Every file gets a description and token estimate. Claude reads this instead of opening files when the summary is enough.
## src/
- ` index.ts ` - Main entry point. Exports createProgram() for CLI. ( ~ 180 tok)
- ` server.ts ` - Express HTTP server with middleware chain. ( ~ 520 tok)
## src/api/
- ` auth.ts ` - JWT validation middleware. Reads from env.JWT_SECRET. ( ~ 340 tok)
- ` users.ts ` - CRUD endpoints for /api/users. Pagination via query params. ( ~ 890 tok)
token-ledger.json - the receipt
Every session gets a line item. Lifetime totals tell you if OpenWolf is actually saving tokens.
{
"lifetime" : {
"total_tokens_estimated" : 503978 ,
"total_reads" : 287 ,
"total_writes" : 269 ,
"total_sessions" : 15 ,
"anatomy_hits" : 198 ,
"anatomy_misses" : 89 ,
"repeated_reads_blocked" : 106 ,
"estimated_savings_vs_bare_cli" : 2066959
}
}
buglog.json - the bug memory
Before fixing anything, Claude checks if the fix is already known. After fixing, it logs the solution.
{
"id" : " bug-012 " ,
"error_message" : " TypeError: Cannot read properties of undefined (reading 'map') " ,
"file" : " src/components/UserList.tsx " ,
"root_cause" : " API response was null when users array was expected " ,
"fix" : " Added optional chaining: data?.users?.map() and fallback empty array " ,
"tags" : [ " null-check " , " api-response " , " react " ]
}
Commands
openwolf init Initialize .wolf/ and register hooks
openwolf status Show health, stats, file integrity
openwolf scan Refresh the project structure map
openwolf scan --check Verify anatomy matches filesystem (exits 1 if stale)
openwolf dashboard Open the real-time web dashboard
openwolf daemon start Start background task scheduler
openwolf daemon stop Stop the scheduler
openwolf daemon restart Restart the scheduler
openwolf daemon logs View scheduler logs
openwolf cron list Show all scheduled tasks
openwolf cron run <id> Trigger a task manually
openwolf cron retry <id> Retry a dead-lettered task
openwolf designqc Capture full-page screenshots for design evaluation
openwolf bug search <term> Search bug memory for known fixes
openwolf update Update all registered projects to latest version
openwolf restore [backup] Restore .wolf/ from a timestamped backup
Design QC
Capture full-page screenshots of your running app and let Claude evaluate the design.
openwolf designqc
Auto-detects your dev server, captures viewport-height JPEG sections of every route, and saves them to .wolf/designqc-captures/ . Then tell Claude to read the screenshots and evaluate. Requires puppeteer-core .
Ask Claude to help you pick a UI framework. OpenWolf ships a curated knowledge base of 12 frameworks (shadcn/ui, Aceternity, Magic UI, DaisyUI, HeroUI, Chakra, Flowbite, Preline, Park UI, Origin UI, Headless UI, Cult UI) with battle-tested migration prompts. Claude reads .wolf/reframe-frameworks.md , asks you a few questions, and executes the migration with the right prompt for your project.
OpenWolf is not an AI wrapper. It is 6 hook scripts and a .wolf/ directory. It doesn't run your AI for you or change your workflow. It gives Claude Code what it lacks: a project map so it reads less, a memory so it learns faster, and a ledger so you see where tokens go.
Optional: PM2 for persistent background tasks
Optional: puppeteer-core for Design QC screenshots
Claude Code hooks are a relatively new feature. OpenWolf falls back to CLAUDE.md instructions when hooks don't fire.
Token tracking is estimation-based (character-to-token ratio), not exact API counts. Accurate to within ~15%.
cerebrum.md depends on Claude following instructions to update it after corrections. Compliance is ~85-90%, not 100%.
This is v1.0.4. Things may break. File issues .
We were building products with Claude Code at Cytostack when we noticed something off. Sessions were eating through tokens faster than they should. When we dug in, we found Claude re-reading the same files multiple times, scanning entire directories to find one function, and having no way to know what a file contained without opening it. There was no project map, no read awareness, no token visibility. So we built the tooling we wished existed -- a file index so Claude reads less, a learning memory so it gets smarter, and a ledger that tracks every token. That became OpenWolf.
Built by Farhan Palathinkal Afsal - Cytostack
Sharper context. Fewer tokens. Open-source middleware for Claude Code.
AGPL-3.0 license
Code of conduct
There was an error while loading. Please reload this page .
155
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

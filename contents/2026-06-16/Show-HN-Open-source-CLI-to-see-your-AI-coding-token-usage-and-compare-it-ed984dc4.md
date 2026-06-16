---
source: "https://github.com/amiinwani/whoburnedmore.com"
hn_url: "https://news.ycombinator.com/item?id=48549250"
title: "Show HN: Open-source CLI to see your AI coding token usage and compare it"
article_title: "GitHub - amiinwani/whoburnedmore.com: See how many tokens your AI coding agents really burned. · GitHub"
author: "arhaam"
captured_at: "2026-06-16T01:07:52Z"
capture_tool: "hn-digest"
hn_id: 48549250
score: 1
comments: 0
posted_at: "2026-06-16T01:03:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source CLI to see your AI coding token usage and compare it

- HN: [48549250](https://news.ycombinator.com/item?id=48549250)
- Source: [github.com](https://github.com/amiinwani/whoburnedmore.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T01:03:04Z

## Translation

タイトル: HN を表示: AI コーディング トークンの使用状況を確認して比較するためのオープンソース CLI
記事のタイトル: GitHub - amiinwani/whoburnedmore.com: AI コーディング エージェントが実際に焼いたトークンの数を確認します。 · GitHub
説明: AI コーディング エージェントが実際に燃やしたトークンの数を確認します。 - amiinwani/whoburnedmore.com
HN テキスト: 私は毎日 Claude Code、Codex、Cursor を使用していますが、これらすべてを合わせて実際にどれだけ消費しているかわかりませんでした。各ツールは、たとえ表示されるとしても、独自の場所に独自の使用法を示します (ほとんどのツールは表示しません)。必要なのは 1 つの番号だけです。それで私はこれを書きました。 npx whoburnedmore は、これらのツールがローカルに既に書き込んでいる使用状況/ログ ファイルを読み取り、それらを合計して、トークンの合計、コストの見積もり、アクティブな日数、および最も負荷の高い 1 日を含むダッシュボードを出力します。ウェブサイトには、AI の使用状況を他のすべての AI 使用状況と比較する公開リーダーボードもあります。フィードバックと賛成票をいただければ幸いです。github は /amiinwani/whoburnedmore.com です。
製品ハント: https://www.producthunt.com/products/whoburnedmore

記事本文:
GitHub - amiinwani/whoburnedmore.com: AI コーディング エージェントが実際に書き込んだトークンの数を確認します。 · GitHub
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
アミインワニ
/
whoburnedmore.com
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
C

頌歌
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows src src test test .gitignore .gitignore CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json Viewすべてのファイル リポジトリ ファイルのナビゲーション
AI コーディング エージェントが実際に燃やしたトークンの数を確認します。
クロード コードのセッション履歴を読み取り、費やしたトークンの数とドルを正確に表示する、高速で 100% ローカルのコマンドライン ツールです。アカウントもサインアップもアップロードもありません。あなたのマシンから何も離れることはありません。
whoburnedmore.com · ホストされたリーダーボード
🔥 whoburnedmore — ローカル AI トークン書き込みレポート
─────────────────
18.2 億トークンが推定 $3,410.00 を消費しました。
12,704 件のアシスタント メッセージ · 18 アクティブ日 · 2026-05-29 → 2026-06-15
モデル別
████████░░░░░░░░░░ クロード-作品-4-8 11億ドル $2,512.40
█████░░░░░░░░░░░░░ クロード・ソネット-4-6 512.00M $640.10
██░░░░░░░░░░░░░░░░ クロード俳句-4-5 210.00M $36.20
プロジェクト別
███████░░░░░░░░░░░ API 903.00M $1,640.00
████░░░░░░░░░░░░░░ ウェブアプリ 540.00M $980.00
██░░░░░░░░░░░░░░░░インフラ 377.00M $790.00
プロンプトキャッシュ 97.4% 読み取りヒット率 (1.71B キャッシュ読み取り)
─────────────────
100% ローカル · マシンには何も残りません。
公開掲示板で比較してみよう

→ https://whoburnedmore.com
✨ 特徴
🔒完全にローカルです。すでにディスク上にあるトランスクリプトを読み取ります。ネットワーク リクエストはゼロです。テレメトリ、アカウント、API キーはありません。
💸 実際のコストの見積もり。透明で編集可能な価格表を使用して、生のトークン数をドルの数字に変換します。
🧩 重要な内訳。書き込みがモデルごと、プロジェクトごとに分割されて表示されるため、トークンがどこに送られたのかが正確にわかります。
⚡ プロンプトキャッシュの洞察。キャッシュの読み取りヒット率を明らかにします。これは、実際に支払う金額を左右する唯一の最大の要因です。
🖼️ 美しい HTML ダッシュボード。 --html は、任意のブラウザで開いたり共有したりできる、自己完結型のオフライン ダッシュボードを作成します。
🪶 実行時に小さくて依存関係がありません。ランタイム依存関係のない最大 15 KB のバンドル。ここで読めるコード以外に信頼できるものはありません。
このマシンでは Node.js 20 以降とクロード コードを使用する必要があります。
npx github:amiinwani/whoburnedmore.com
またはクローンを作成して実行します
git clone https://github.com/amiinwani/whoburnedmore.com.git
cd フーバーンモア.com
npm install # CLI もビルドします
npm start # == ノード dist/cli.js
📖 使用法
whoburnedmore # 書き込みレポートを印刷する
whoburnedmore --html # ./whoburnedmore.html とも書きます
whoburnedmore --html out.html # ...カスタム パスへ
whoburnedmore --since 30 # 過去 30 日間のみをカウントします
whoburnedmore --dir < path > # カスタム ディレクトリからトランスクリプトを読み取ります
whoburnedmore --json # 生の集約された JSON (どこにでもパイプできます)
誰が燃えたもっと --バージョン
誰が燃えたのか -- ヘルプ
旗
何をするのか
--html [ファイル]
自己完結型の HTML ダッシュボードを作成します (デフォルトは ./whoburnedmore.html )。
-- <日> 以降
過去 N 日間の使用量のみを含めます。
--dir <パス>
~/.claude/projects の代わりにカスタム トランスクリプト ディレクトリを指定します。
--json
集計されたデータをレポートではなく JSON として印刷します。
--バージョン / --ヘルプ
バージョン/ヘルプを出力します

p.
🔍 仕組み
クロード コードは、セッションごとに 1 つの JSON Lines トランスクリプトを次の場所に保存します。
~/.claude/projects/<プロジェクト>/<セッションID>.jsonl
アシスタントがこれらのファイルを提出するたびに、使用ブロック (入力トークン、出力トークン、プロンプト キャッシュの読み取り/書き込みカウント) に加えて、モデル名とタイムスタンプが含まれます。 whoburnedmore は、これらのファイルを 1 行ずつストリーミングし、モデル別、プロジェクト別、および日ごとのカウントを合計し、公開価格表を乗算して金額の見積もりを出し、レポートを作成します。それがすべてのトリックです。コードは意図的に小さくなっており、 src/scan.ts から始めます。
このツールは設計上読み取り専用でオフラインです。
ファイルを読み取るだけです。トランスクリプトに書き込んだり、何かを削除したりすることはありません。
いかなる種類のネットワーク リクエストも作成しません。分析も、「テレホン ホーム」も、自動更新も行いません。
オプションの HTML ダッシュボードは完全に自己完結型であり、外部フォント、スクリプト、トラッカーはありません。
プロンプトとコードの内容は決して解析されません。数値の使用状況カウンターとモデル名のみが読み取られます。
確認したいですか?約400行です。 src/ を読み取るか、 --json を指定して実行し、計算内容を正確に検査します。
この CLI は、オープンソースのローカル専用の whoburnedmore.com の兄弟です。whoburnedmore.com は、開発者がクロード コード、コーデックス、カーソルなどの書き込み量を比較する公開リーダーボードです。 Web サイトは別個のホスト製品です。このリポジトリにはローカル ツールのみが含まれており、そのバックエンドは含まれていません。
問題や PR は歓迎です — COTRIBUTING.md を参照してください。最初の良い貢献: 価格表にモデルを追加する、より多くのエージェントのトランスクリプト形式をサポートする、またはダッシュボードを改善する。
npm install # インストール + ビルド
npm test # 単体テストを実行する
npm run typecheck # 厳密な型チェック
npm run build # dist/cli.js にバンドル
📄ライセンス
AI コーディング エージェントが実際に燃やしたトークンの数を確認します。
がありました

ロード中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

See how many tokens your AI coding agents really burned. - amiinwani/whoburnedmore.com

I use Claude Code, Codex, Cursor every day and had no idea how much I was actually burning across all of them combined. Each tool shows its own usage (most don't) in its own place, if at all, and I just wanted one number. So I wrote this. npx whoburnedmore reads the usage/log files these tools already write locally, sums them up, and prints a dashboard: total tokens, a cost estimate, active days, and your heaviest single day. There is also a public leaderboard on the webiste which compares your ai usage with all others Feedback & a upvote is appreciated github is /amiinwani/whoburnedmore.com
product hunt : https://www.producthunt.com/products/whoburnedmore

GitHub - amiinwani/whoburnedmore.com: See how many tokens your AI coding agents really burned. · GitHub
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
amiinwani
/
whoburnedmore.com
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
See how many tokens your AI coding agents really burned.
A fast, 100% local command-line tool that reads your Claude Code session history and shows you exactly how many tokens — and how many dollars — you've spent. No account, no sign-up, no upload. Nothing ever leaves your machine.
whoburnedmore.com · the hosted leaderboard
🔥 whoburnedmore — your local AI token burn report
────────────────────────────────────────────────
1.82B tokens burned $3,410.00 est.
12,704 assistant messages · 18 active days · 2026-05-29 → 2026-06-15
By model
████████░░░░░░░░░░ claude-opus-4-8 1.10B $2,512.40
█████░░░░░░░░░░░░░ claude-sonnet-4-6 512.00M $640.10
██░░░░░░░░░░░░░░░░ claude-haiku-4-5 210.00M $36.20
By project
███████░░░░░░░░░░░ api 903.00M $1,640.00
████░░░░░░░░░░░░░░ web-app 540.00M $980.00
██░░░░░░░░░░░░░░░░ infra 377.00M $790.00
Prompt cache 97.4% read-hit rate (1.71B cached reads)
────────────────────────────────────────────────
100% local · nothing left your machine.
Compare on the public board → https://whoburnedmore.com
✨ Features
🔒 Completely local. Reads the transcripts already on your disk. Makes zero network requests — no telemetry, no account, no API keys.
💸 Real cost estimates. Turns raw token counts into dollar figures using a transparent, editable pricing table .
🧩 Breakdowns that matter. See your burn split by model and by project , so you know exactly where the tokens went.
⚡ Prompt-cache insight. Surfaces your cache read-hit rate — the single biggest lever on what you actually pay.
🖼️ Beautiful HTML dashboard. --html writes a self-contained, offline dashboard you can open in any browser or share.
🪶 Tiny & dependency-free at runtime. A ~15 KB bundle with no runtime dependencies. Nothing to trust but the code you can read here.
You'll need Node.js 20+ and some Claude Code usage on this machine.
npx github:amiinwani/whoburnedmore.com
Or clone and run
git clone https://github.com/amiinwani/whoburnedmore.com.git
cd whoburnedmore.com
npm install # also builds the CLI
npm start # == node dist/cli.js
📖 Usage
whoburnedmore # print your burn report
whoburnedmore --html # also write ./whoburnedmore.html
whoburnedmore --html out.html # ...to a custom path
whoburnedmore --since 30 # only count the last 30 days
whoburnedmore --dir < path > # read transcripts from a custom directory
whoburnedmore --json # raw aggregated JSON (pipe it anywhere)
whoburnedmore --version
whoburnedmore --help
Flag
What it does
--html [file]
Write a self-contained HTML dashboard (default ./whoburnedmore.html ).
--since <days>
Only include usage from the last N days.
--dir <path>
Point at a custom transcripts directory instead of ~/.claude/projects .
--json
Print the aggregated data as JSON instead of the report.
--version / --help
Print the version / help.
🔍 How it works
Claude Code stores one JSON Lines transcript per session under:
~/.claude/projects/<project>/<session-id>.jsonl
Every assistant turn in those files carries a usage block — input tokens, output tokens, and the prompt-cache read/write counts — plus the model name and a timestamp. whoburnedmore streams through those files line by line, adds the counts up by model, by project, and by day, multiplies by a public pricing table for a dollar estimate, and draws the report. That's the whole trick. The code is small on purpose — start with src/scan.ts .
This tool is read-only and offline by design.
It only reads files; it never writes to your transcripts or deletes anything.
It makes no network requests of any kind — no analytics, no "phone home," no auto-update.
The optional HTML dashboard is fully self-contained: no external fonts, scripts, or trackers.
Your prompts and code are never parsed for content — only the numeric usage counters and the model name are read.
Want to verify? It's ~400 lines. Read src/ , or run with --json and inspect exactly what it computes.
This CLI is the open-source, local-only sibling of whoburnedmore.com — a public leaderboard where developers compare how much they've burned across Claude Code, Codex, Cursor and more. The website is a separate, hosted product; this repository contains only the local tool and none of its backend.
Issues and PRs welcome — see CONTRIBUTING.md . Good first contributions: add models to the pricing table , support more agents' transcript formats, or improve the dashboard.
npm install # install + build
npm test # run the unit tests
npm run typecheck # strict type check
npm run build # bundle to dist/cli.js
📄 License
See how many tokens your AI coding agents really burned.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

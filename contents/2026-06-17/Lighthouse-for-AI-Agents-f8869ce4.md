---
source: "https://github.com/lucioduran/ax-audit"
hn_url: "https://news.ycombinator.com/item?id=48567130"
title: "Lighthouse for AI Agents"
article_title: "GitHub - lucioduran/ax-audit: Audit websites for AI Agent Experience (AX) readiness · GitHub"
author: "lucioduran"
captured_at: "2026-06-17T07:58:34Z"
capture_tool: "hn-digest"
hn_id: 48567130
score: 1
comments: 0
posted_at: "2026-06-17T07:46:36Z"
tags:
  - hacker-news
  - translated
---

# Lighthouse for AI Agents

- HN: [48567130](https://news.ycombinator.com/item?id=48567130)
- Source: [github.com](https://github.com/lucioduran/ax-audit)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T07:46:36Z

## Translation

タイトル: AI エージェントの灯台
記事のタイトル: GitHub - lucioduran/ax-audit: AI Agent Experience (AX) の準備のための Web サイトの監査 · GitHub
説明: AI Agent Experience (AX) の準備状況について Web サイトを監査する - lucioduran/ax-audit

記事本文:
GitHub - lucioduran/ax-audit: AI Agent Experience (AX) の準備状況を監査する Web サイト · GitHub
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
ルシオデュラン
/
斧監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
54 コミット 54 コミット .github/ workflows .github/ workflows bin bin docs docs src src test test .gitignore .gitignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md ax-logo.svg ax-logo.svg eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントの灯台。あらゆる Web サイトの AI Agent Experience (AX) の準備状況を数秒で監査します。
npx ax-audit https://your-site.com
AX監査レポート
https://lucioduran.com
███████████████████████████████████░░░░░ 88/100 良い
LLMs.txt (100/100)
PASS /llms.txt が存在します
PASS /llms.txt Content-Type OK (テキスト/プレーン)
PASS H1 の見出し: 「ルシオ デュラン — 個人ポートフォリオ」
ロボット.txt (100/100)
PASS 8 つのコア AI クローラーすべてが明示的に構成されている
PASS ユーザー エージェントに対して宣言されたコンテンツ シグナル: * — search=yes、ai-train=no
コンテンツネゴシエーション (100/100)
PASS ホームページはコンテンツ ネゴシエーション経由でマークダウンを提供します (受け入れ: テキスト/マークダウン)
PASS Markdown は HTML 表現よりも約 95% 軽量です
...
なぜ
AI エージェントと LLM は、Web サイトのクロール、インデックス作成、および対話をますます行っています。 Lighthouse が Web パフォーマンスを監査し、axe-core がアクセシビリティを監査するのと同じように、ax-audit は、サイトが AI エージェント エコシステム (検出ファイル、クローラー ポリシー、ライセンス、コンテンツ ネゴシエーション、オペレーターには見えない障害モード (robots.txt で許可されているクローラーをブロックする WAF など)) に対してどの程度準備ができているかを示します。
18 個のチェック — 14 個の重み付け、4 個の情報。完全なリファレンス: docs/checks.md 。
* 3.x の情報: 完全に報告されます。電子はありません。

スコアに影響します。 v4.0 で重み付けされました。
すべての発見事項は、ステップバイステップの修復ガイドにリンクされています。
ax-audit https://example.com # 完全な監査、ターミナル出力
ax-audit https://a.com https://b.com --concurrency 2 # バッチ、並列
ax-audit https://example.com --output markdown # また: json、html
ax-audit https://example.com --checks llms-txt,rsl # チェックのサブセット
ax-audit https://example.com --only-failures # 合格した結果を非表示にする
ax-audit https://example.com --baseline .ax-baseline.json --fail-on-regression 5
終了コード ゲート CI: スコア ≥ 70 の場合は 0、それ以下の場合は 1。フラグの完全なリファレンス: docs/cli.md · CI レシピ (PR コメント、回帰ゲート、スケジュールされた監査): docs/ci.md 。
import { 監査 , バッチ監査 } から 'ax-audit' ;
const report = 監査を待ちます ( { url : 'https://example.com' } ) ;
報告 。総合スコア ; // 0～100
ご報告をさせていただきます。結果 ; // チェックごとの結果
完全な API とタイプ: docs/api.md 。
これらのファイルからレンダリングされた同じドキュメントが lucioduran.com/projects/ax-audit/docs で参照できます。寄稿者: COTRIBUTING.md および SECURITY.md を参照してください。
グレード
スコア
終了コード
素晴らしい
90～100
0
良い
70–89
0
フェア
50–69
1
貧しい
0～49
1
技術
TypeScript 厳密モード · 2 つのランタイム依存関係 (チョーク、コマンダー) · ノード 18 以降の組み込みフェッチ · Promise.allSettled による並列チェック · Vary 対応キーを使用した実行ごとのリクエスト キャッシュ · バックオフによる一時的な失敗の再試行 · テスト依存関係がゼロの node:test での 301 テスト。
貢献は大歓迎です。パイプラインの設計、チェックの構造、新しいチェックに必要な手順 (コード、テスト、ドキュメント、修復ガイド) については、docs/architecture.md を参照してください。
ax-init — このツールが監査する AX ファイルを生成します
ax-cite — AI で抽出可能な構造化データをページに埋め込む
AI Agent Experience (AX) の準備状況について Web サイトを監査する
Apache-2.0ライセンス
投稿者

g
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
22
v3.6.0 — AI ライセンス、クローキング検出、クロール効率、および完全なドキュメント
最新の
2026 年 6 月 9 日
+ 21 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Audit websites for AI Agent Experience (AX) readiness - lucioduran/ax-audit

GitHub - lucioduran/ax-audit: Audit websites for AI Agent Experience (AX) readiness · GitHub
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
lucioduran
/
ax-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
54 Commits 54 Commits .github/ workflows .github/ workflows bin bin docs docs src src test test .gitignore .gitignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md ax-logo.svg ax-logo.svg eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Lighthouse for AI Agents. Audit any website's AI Agent Experience (AX) readiness in seconds.
npx ax-audit https://your-site.com
AX Audit Report
https://lucioduran.com
███████████████████████████████████░░░░░ 88/100 Good
LLMs.txt (100/100)
PASS /llms.txt exists
PASS /llms.txt Content-Type OK (text/plain)
PASS H1 heading: "Lucio Duran — Personal Portfolio"
Robots.txt (100/100)
PASS All 8 core AI crawlers explicitly configured
PASS Content signals declared for User-agent: * — search=yes, ai-train=no
Content Negotiation (100/100)
PASS Homepage serves Markdown via content negotiation (Accept: text/markdown)
PASS Markdown is ~95% lighter than the HTML representation
...
Why
AI agents and LLMs are increasingly crawling, indexing, and interacting with websites. Just like Lighthouse audits web performance and axe-core audits accessibility, ax-audit tells you how ready your site is for the AI agent ecosystem — discovery files, crawler policy, licensing, content negotiation, and the failure modes invisible to operators (like a WAF blocking crawlers your robots.txt allows).
18 checks — 14 weighted, 4 informational. Full reference: docs/checks.md .
* Informational in 3.x: reported in full, no effect on the score. Weighted in v4.0.
Every finding links to a step-by-step remediation guide .
ax-audit https://example.com # full audit, terminal output
ax-audit https://a.com https://b.com --concurrency 2 # batch, in parallel
ax-audit https://example.com --output markdown # also: json, html
ax-audit https://example.com --checks llms-txt,rsl # subset of checks
ax-audit https://example.com --only-failures # hide passing findings
ax-audit https://example.com --baseline .ax-baseline.json --fail-on-regression 5
Exit codes gate CI: 0 for score ≥ 70, 1 below. Full flag reference: docs/cli.md · CI recipes (PR comments, regression gates, scheduled audits): docs/ci.md .
import { audit , batchAudit } from 'ax-audit' ;
const report = await audit ( { url : 'https://example.com' } ) ;
report . overallScore ; // 0–100
report . results ; // per-check findings
Full API and types: docs/api.md .
The same documentation is browsable at lucioduran.com/projects/ax-audit/docs , rendered from these files. Contributors: see CONTRIBUTING.md and SECURITY.md .
Grade
Score
Exit Code
Excellent
90–100
0
Good
70–89
0
Fair
50–69
1
Poor
0–49
1
Tech
TypeScript strict mode · 2 runtime dependencies ( chalk , commander ) · Node 18+ built-in fetch · parallel checks via Promise.allSettled · per-run request cache with Vary -aware keys · transient-failure retries with backoff · 301 tests on node:test with zero test dependencies.
Contributions are welcome — see docs/architecture.md for the pipeline design, check anatomy, and the steps (code, tests, docs, remediation guide) a new check requires.
ax-init — generate the AX files this tool audits
ax-cite — embed AI-extractable structured data in your pages
Audit websites for AI Agent Experience (AX) readiness
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
22
v3.6.0 — AI licensing, cloaking detection, crawl efficiency & full documentation
Latest
Jun 9, 2026
+ 21 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/syntropicsignal-ai/ai-visibility-audit"
hn_url: "https://news.ycombinator.com/item?id=48584153"
title: "Show HN: Reverse engineering user queries for ChatGPT/Gemini"
article_title: "GitHub - syntropicsignal-ai/ai-visibility-audit: Open-source AI visibility audit for ecommerce - see if ChatGPT, Gemini & Google AI recommend your store or your competitor's. Self-hosted, printable PDF report. · GitHub"
author: "biduskamil"
captured_at: "2026-06-18T13:18:16Z"
capture_tool: "hn-digest"
hn_id: 48584153
score: 1
comments: 2
posted_at: "2026-06-18T12:12:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Reverse engineering user queries for ChatGPT/Gemini

- HN: [48584153](https://news.ycombinator.com/item?id=48584153)
- Source: [github.com](https://github.com/syntropicsignal-ai/ai-visibility-audit)
- Score: 1
- Comments: 2
- Posted: 2026-06-18T12:12:51Z

## Translation

タイトル: HN の表示: ChatGPT/Gemini のユーザー クエリのリバース エンジニアリング
記事のタイトル: GitHub - syntropicsignal-ai/ai-visibility-audit: e コマース向けのオープンソース AI 可視性監査 - ChatGPT、Gemini、Google AI が自分のストアを推奨するか、競合他社のストアを推奨するかを確認します。自己ホスト型の印刷可能な PDF レポート。 · GitHub
説明: e コマース向けのオープンソース AI 可視性監査 - ChatGPT、Gemini、Google AI があなたのストアまたは競合他社のストアを推奨しているかどうかを確認します。自己ホスト型の印刷可能な PDF レポート。 - syntropicsignal-ai/ai-visibility-audit

記事本文:
GitHub - syntropicsignal-ai/ai-visibility-audit: e コマース向けのオープンソース AI 可視性監査 - ChatGPT、Gemini、Google AI があなたのストアまたは競合他社のストアを推奨しているかどうかを確認します。自己ホスト型の印刷可能な PDF レポート。 · GitHub
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
シントロピックシグナルai
/

AI 可視性監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
syntropicsignal-ai/ai-visibility-audit
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github/ workflows .github/ workflows api api docker docker web web .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI があなたの店舗を推奨するか、それとも競合他社の店舗を推奨するかを確認してください。オープンソース
買い物客が ChatGPT に質問したときにブランドと製品がどのように表示されるかを監査します。
Gemini、Google AI モード、Google AI の概要 「Y にとって最適な X」または「Z を購入できる場所」。
e コマース向けに構築 — あらゆるブランド (SaaS、ローカル、B2B) で機能します。
買い物客はますます AI アシスタントに何を買うべきかを尋ね、得られるものを購入します
推奨されます。オンライン ストアへの AI 紹介トラフィックは 3 桁で増加
年々。しかし、誰かがChatGPTに「パンクに最適なランニングシューズ」と尋ねると、
フィート」か「500 ドル以下のエスプレッソマシンをどこで買えますか」、どちらかわかりません。
あなたのストアが推奨されたり、無視されたり、競合他社に負けたりします。
このツールは、すべての主要な AI アシスタント、すべての購入者のクエリについて表示します。
それはあなたのカテゴリーでは重要です。
監査レポート — 利害関係者に渡すことができるクライアント対応の PDF: KPI、
競合他社のシェア、不足しているクエリ、具体的な推奨事項。
AI があなたの代わりにどの競合他社を推奨するか、またどの購入者が問い合わせるか
あなたのストアが見えないような意図の高い質問 (「最高の…」、「どこで買えますか…」)
AI がソースとして引用するドメイン (および存在する場合はどのページも)
あなたのブランドが言及されたときの感情とあなたがどのように評価されるか
最初に修正すべきものを優先順位付けしたリスト
bを表示中

uilt-in サンプル データセット — ワンクリックでロードでき、API キーは必要ありません。
回答 — あなたのカテゴリに関するすべての AI 回答: どの店舗が名前を付けられ、推奨され、引用されます。
ダッシュボード — 可視性、GEO ファネル、競合他社のシェア、および AI が他の人を推奨する購入者のクエリ。
トピック — AI の回答でどの製品カテゴリーが勝ち、どのカテゴリーが負けるか。
プロンプト – 追跡される購入者のクエリ。それぞれに、AI がストアに名前を付ける頻度と時間の経過に伴う傾向が含まれます。
git clone https://github.com/syntropicsignal-ai/ai-visibility-audit.git
cd ai-可視性-監査
ドッカー構成 -d
次に、http://localhost:8080 を開きます。最初のアクセス時に API キーを入力するか、
サンプル データセットをロードし、キーをまったく使用せずに探索します。移行の実行
開始時に自動的に。
ストアの URL を指定します — Shopify、WooCommerce、PrestaShop、Magento、
BigCommerce、Medusa、Saleor、Vendure、または Shoper、IdoSell などの PL プラットフォーム
&BaseLinker。監査はスタックに関係なく、あらゆるブランドまたはドメインで実行されます。
Docker デスクトップ (または Docker エンジン + Compose v2)
API キー: Gemini および Exa (プロンプト生成 + 分析)、DataForSEO
(Google AI の概要 + キーワード シグナル)。
Bright Data — ChatGPT、Gemini、Google AI モード ソースのロックを解除します。
このツールが測定する主な AI アシスタント。それがなければ追跡することになります
Google AI の概要のみ (DataForSEO 経由)。
api/ FastAPI バックエンド (Python、UV)
web/Vue 3 フロントエンド (TypeScript、Vite)
docker/コンテナビルドファイル
ライセンス
e コマース向けのオープンソース AI 可視性監査 - ChatGPT、Gemini、Google AI があなたのストアまたは競合他社のストアを推奨しているかどうかを確認します。自己ホスト型の印刷可能な PDF レポート。
syntropicsignal.ai/ecommerce
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
2
v0.2.0
最新の
2026 年 6 月 12 日
+1レ

あせ
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source AI visibility audit for ecommerce - see if ChatGPT, Gemini & Google AI recommend your store or your competitor's. Self-hosted, printable PDF report. - syntropicsignal-ai/ai-visibility-audit

GitHub - syntropicsignal-ai/ai-visibility-audit: Open-source AI visibility audit for ecommerce - see if ChatGPT, Gemini & Google AI recommend your store or your competitor's. Self-hosted, printable PDF report. · GitHub
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
syntropicsignal-ai
/
ai-visibility-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
syntropicsignal-ai/ai-visibility-audit
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github/ workflows .github/ workflows api api docker docker web web .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml View all files Repository files navigation
See if AI recommends your store — or your competitor's. An open-source
audit of how your brand and products show up when shoppers ask ChatGPT,
Gemini, Google AI Mode & Google AI Overview "best X for Y" or "where to buy Z".
Built for ecommerce — works for any brand (SaaS, local, B2B).
Shoppers increasingly ask AI assistants what to buy — and buy what gets
recommended. AI referral traffic to online stores is growing triple-digits
year over year. But when someone asks ChatGPT "best running shoes for flat
feet" or "where to buy an espresso machine under $500", you can't see whether
your store is recommended, ignored, or losing to a competitor.
This tool shows you — across every major AI assistant, for every buyer query
that matters in your category.
Audit report — a client-ready PDF you can hand to a stakeholder: KPIs,
competitor share, the queries you're missing, and concrete recommendations.
Which competitor AI recommends instead of you — and for which buyer queries
The high-intent questions ("best…", "where to buy…") where your store is invisible
Which domains AI cites as sources (and which of your pages, if any)
Sentiment and how you're framed when your brand is mentioned
A prioritized list of what to fix first
Showing the built-in sample dataset — load it in one click, no API keys required.
Responses — every AI answer about your category: which store gets named, recommended, and cited.
Dashboard — visibility, the GEO funnel, competitor share, and the buyer queries where AI recommends someone else.
Topics — which product categories you win in AI answers, and which you're losing.
Prompts — the buyer queries being tracked, each with how often AI names your store and the trend over time.
git clone https://github.com/syntropicsignal-ai/ai-visibility-audit.git
cd ai-visibility-audit
docker compose up -d
Then open http://localhost:8080 — enter your API keys on first visit, or
load the sample dataset and explore with no keys at all. Migrations run
automatically on start.
Point it at your store URL — Shopify, WooCommerce, PrestaShop, Magento,
BigCommerce, Medusa, Saleor, Vendure, or PL platforms like Shoper, IdoSell
& BaseLinker. The audit runs on any brand or domain, regardless of stack.
Docker Desktop (or Docker Engine + Compose v2)
API keys: Gemini and Exa (prompt generation + analysis), DataForSEO
(Google AI Overview + keyword signals).
Bright Data — unlocks the ChatGPT, Gemini, and Google AI Mode sources,
the main AI assistants this tool measures. Without it you'll track
Google AI Overview only (via DataForSEO).
api/ FastAPI backend (Python, uv)
web/ Vue 3 frontend (TypeScript, Vite)
docker/ Container build files
License
Open-source AI visibility audit for ecommerce - see if ChatGPT, Gemini & Google AI recommend your store or your competitor's. Self-hosted, printable PDF report.
syntropicsignal.ai/ecommerce
Topics
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
v0.2.0
Latest
Jun 12, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

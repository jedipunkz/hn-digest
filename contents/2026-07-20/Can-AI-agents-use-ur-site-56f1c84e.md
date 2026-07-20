---
source: "https://github.com/Open-Ingress/OpenIngress"
hn_url: "https://news.ycombinator.com/item?id=48985431"
title: "Can AI agents use ur site?"
article_title: "GitHub - Open-Ingress/OpenIngress: Find where AI agents break when navigating your site. · GitHub"
author: "manupareekk"
captured_at: "2026-07-20T22:47:08Z"
capture_tool: "hn-digest"
hn_id: 48985431
score: 5
comments: 9
posted_at: "2026-07-20T22:01:49Z"
tags:
  - hacker-news
  - translated
---

# Can AI agents use ur site?

- HN: [48985431](https://news.ycombinator.com/item?id=48985431)
- Source: [github.com](https://github.com/Open-Ingress/OpenIngress)
- Score: 5
- Comments: 9
- Posted: 2026-07-20T22:01:49Z

## Translation

タイトル: AI エージェントはあなたのサイトを使用できますか?
記事のタイトル: GitHub - Open-Ingress/OpenIngress: サイトを移動するときに AI エージェントが中断する場所を見つけます。 · GitHub
説明: サイト内を移動するときに AI エージェントが中断する場所を見つけます。 - オープンイングレス/オープンイングレス

記事本文:
GitHub - Open-Ingress/OpenIngress: サイトを移動するときに AI エージェントが中断する場所を見つけます。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オープンイングレス
/
OpenIngress
公共
そうではない

化
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット バックエンド バックエンド デプロイ デプロイ docs docs フロントエンド フロントエンド .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DEPLOY.md DEPLOY.md DEPLOY_AZURE.md DEPLOY_AZURE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SELF_HOST.md SELF_HOST.md package-lock.json package-lock.json Railway.toml Railway.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自己ホスト型 Web 操作性分析。
クロール カバレッジ、アクセシビリティ ツリー検査、およびエージェント ナビゲーション ブレーク ポイント。
オープニングレス.dev ·
建築・
クイックスタート ·
セルフホスト
AI エージェントは、ページを開いたり、ボタンをクリックしたり、フォームに記入したりして、人々のために Web を閲覧し始めています。サイトはブラウザーでは問題なく見えても、ラベルのないボタン、JavaScript でのみ読み込まれるページ、エージェントが読み取るページ構造に表示されないリンクなど、エージェントにとっては使いにくい場合があります。
検索可能 (SEO、引用、llms.txt) であることと、実際のサイトでエージェントが使用できることは同じではありません。
OpenIngress は、ブラウザと同じようにサイトをクロールし、エージェントが実際に表示およびクリックできるものを検査し、そのクロールに対して LLM ガイド付きタスクを実行します。出力は、どのページに到達したか、何が壊れたか、どこでナビゲーションが停止したかなどのレポートです。
フローチャートTB
サブグラフ クロール["クロール (劇作家)"]
URL[サイト URL] --> BFS[BFS ページ検出]
BFS --> CAP[HTML · スクリーンショット · aria スナップショット]
終わり
サブグラフの分析["分析"]
CAP --> STATIC[静的な操作性]
CAP --> GRAPH[ナビゲーション グラフ]
STATIC --> GAPS[ギャップ分類]
終わり
サブグラフの探索["探索 (BYOK LLM)"]
グラフ --> カタログ[Pa

ゲーカタログ]
カタログ --> ジョブ[探索ジョブ]
LLM[OpenAI互換API] --> ジョブ
JOBS --> TRACE[ジョブ トレース · ブロッカー]
終わり
サブグラフレポート["レポート"]
GAPS --> OUT[カバレッジ + ブレークポイント]
トレース --> アウト
終わり
読み込み中
ステージ
入力
出力
クロール
公開URL
ページグラフ、DOM、スクリーンショット、Playwright アクセシビリティスナップショット
分析する
クロールアーティファクト
操作性スコア、ラベル カバレッジ、llms.txt、水和チェック、ギャップ分類法
探検する
カタログ + LLM キー
到達可能性ジョブ、タスク試行、ブロッカーの証拠
レポート
分析 + ジョブ結果
準備状況メトリクス、ブレークポイント、マークダウンエクスポート
デフォルトのクロール境界: 深さ 3 、最大 100 ページ。
レイヤー
コンポーネント
API
フラスコ、ガニコーン
自動化
Playwright (Chromium)、Python 3.10+
探検
OpenAI 互換 LLM (Azure、Gemini、Ollama、ローカル プロキシ)
UI
Vue 3、ヴィテ
持続性
ローカルファイルシステム。オプションのスーパーベース
クイックスタート
cp .env.example .env
cp バックエンド/.env.example バックエンド/.env
cp フロントエンド/.env.example フロントエンド/.env
# backend/.env には LLM_API_KEY が必要です
インストールする
バックエンド # ターミナル 1 を作成します — API
フロントエンド # ターミナル 2 を作成します — UI
構成、Docker、およびデプロイメント: SELF_HOST.md 。
変数
必須
説明
LLM_API_KEY
はい
プロバイダーキー (OpenAI または互換性)
LLM_BASE_URL
いいえ
カスタムエンドポイント
LLM_MODEL_NAME
いいえ
デフォルトの gpt-4o-mini
AUTH_DISABLED
いいえ
デフォルト 1
実行と探索には、構成された LLM キーが必要です。
パス
説明
バックエンド/
Flask API、クロール/探索ワーカー、レポート生成
フロントエンド/
Vue ダッシュボード
ドキュメント/
エージェントの概要、劇作家の録音メモ
Dockerfile
Playwright + ガニコーンコンテナ
ライセンス
サイト内を移動するときに AI エージェントが中断する場所を見つけます。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
2
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
そこには

読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Find where AI agents break when navigating your site. - Open-Ingress/OpenIngress

GitHub - Open-Ingress/OpenIngress: Find where AI agents break when navigating your site. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Open-Ingress
/
OpenIngress
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits backend backend deploy deploy docs docs frontend frontend .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DEPLOY.md DEPLOY.md DEPLOY_AZURE.md DEPLOY_AZURE.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SELF_HOST.md SELF_HOST.md package-lock.json package-lock.json railway.toml railway.toml View all files Repository files navigation
Self-hosted web operability analysis.
Crawl coverage, accessibility-tree inspection, and agent navigation break points.
openingress.dev ·
Architecture ·
Quick start ·
Self-host
AI agents are starting to browse the web for people — opening pages, clicking buttons, filling forms. A site can look fine in a browser and still be hard for an agent to use: buttons without labels, pages that only load in JavaScript, links that never show up in the page structure agents read.
Being findable (SEO, citations, llms.txt ) is not the same as being usable by an agent on the live site.
OpenIngress crawls a site like a browser would, inspects what an agent can actually see and click, then runs LLM-guided tasks against that crawl. The output is a report: which pages were reached, what broke, and where navigation stalled.
flowchart TB
subgraph crawl["Crawl (Playwright)"]
URL[Site URL] --> BFS[BFS page discovery]
BFS --> CAP[HTML · screenshots · aria snapshots]
end
subgraph analyze["Analyze"]
CAP --> STATIC[Static operability]
CAP --> GRAPH[Navigation graph]
STATIC --> GAPS[Gap taxonomy]
end
subgraph explore["Explore (BYOK LLM)"]
GRAPH --> CATALOG[Page catalog]
CATALOG --> JOBS[Exploration jobs]
LLM[OpenAI-compatible API] --> JOBS
JOBS --> TRACE[Job traces · blockers]
end
subgraph report["Report"]
GAPS --> OUT[Coverage + break points]
TRACE --> OUT
end
Loading
Stage
Input
Output
Crawl
Public URL
Page graph, DOM, screenshots, Playwright accessibility snapshots
Analyze
Crawl artifacts
Operability scores, label coverage, llms.txt , hydration checks, gap taxonomy
Explore
Catalog + LLM key
Reachability jobs, task attempts, blocker evidence
Report
Analysis + job results
Readiness metrics, break points, markdown export
Default crawl bounds: depth 3 , max 100 pages.
Layer
Components
API
Flask, Gunicorn
Automation
Playwright (Chromium), Python 3.10+
Exploration
OpenAI-compatible LLM (Azure, Gemini, Ollama, local proxy)
UI
Vue 3, Vite
Persistence
Local filesystem; optional Supabase
Quick start
cp .env.example .env
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
# LLM_API_KEY required in backend/.env
make install
make backend # terminal 1 — API
make frontend # terminal 2 — UI
Configuration, Docker, and deployment: SELF_HOST.md .
Variable
Required
Description
LLM_API_KEY
yes
Provider key (OpenAI or compatible)
LLM_BASE_URL
no
Custom endpoint
LLM_MODEL_NAME
no
Default gpt-4o-mini
AUTH_DISABLED
no
Default 1
Runs and exploration require a configured LLM key.
Path
Description
backend/
Flask API, crawl/explore workers, report generation
frontend/
Vue dashboard
docs/
Agent brief, Playwright recording notes
Dockerfile
Playwright + Gunicorn container
License
Find where AI agents break when navigating your site.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
2
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

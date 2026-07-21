---
source: "https://github.com/telemetry-sh/llm-margin-lab"
hn_url: "https://news.ycombinator.com/item?id=48992841"
title: "LLM Margin Lab"
article_title: "GitHub - telemetry-sh/llm-margin-lab: An interactive calculator for AI product unit economics: tokens, caching, routing, retries, and gross margin. · GitHub"
author: "flurly"
captured_at: "2026-07-21T14:56:55Z"
capture_tool: "hn-digest"
hn_id: 48992841
score: 2
comments: 0
posted_at: "2026-07-21T14:29:22Z"
tags:
  - hacker-news
  - translated
---

# LLM Margin Lab

- HN: [48992841](https://news.ycombinator.com/item?id=48992841)
- Source: [github.com](https://github.com/telemetry-sh/llm-margin-lab)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T14:29:22Z

## Translation

タイトル: LLMマージンラボ
記事のタイトル: GitHub - telemetry-sh/llm-margin-lab: AI 製品のユニット エコノミクス (トークン、キャッシュ、ルーティング、再試行、粗利益) のための対話型計算ツール。 · GitHub
説明: AI 製品のユニットエコノミクス (トークン、キャッシュ、ルーティング、再試行、粗利益) のための対話型計算ツール。 - テレメトリー-sh/llm-margin-lab

記事本文:
GitHub - telemetry-sh/llm-margin-lab: AI 製品のユニット エコノミクス (トークン、キャッシュ、ルーティング、再試行、粗利益) のための対話型計算ツール。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。嘆願

このページをリロードしてください。
テレメトリ-sh
/
llm-マージンラボ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ workflows .github/ workflows app app コンポーネント コンポーネント docs docs lib lib test test .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md next-env.d.ts next-env.d.ts next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 機能のユニットエコノミクスを計算するための対話型計算機。
ほとんどの AI 製品はトークン価格を見積もることができます。という質問に答えられる人はほとんどいません
重要: ビジネス モデルの前に、各顧客がどれだけの使用量を消費できるか
休憩？
LLM Margin Lab は、価格設定、トークン、再試行、プロンプト キャッシュ、モデルを組み合わせます
ルーティング、インフラストラクチャ、サポート、目標粗利益を 1 つの共有可能なものにまとめたもの
シナリオ。ローカルで実行され、API キーは必要ありません。
アクティブユーザーあたりの収益と貢献利益
プロンプトキャッシュヒット率とキャッシュ割引
Routing a percentage of work to a cheaper model
Infrastructure and support costs
目標マージンでのユーザーあたりの最大リクエスト数
Margin sensitivity as usage grows
含まれているプリセットは一例です。 Replace them with current model prices
and the operating costs of your own stack.
Requires Node.js 20.9 or newer.
npmインストール
npm 実行開発
http://localhost:3000 を開きます。
シナリオのコピーを使用して、URL にすべての仮定を含めます。受信者は次のことができます
edit the scenario without an account or backend.プロンプト テキストはありません、お客様
データ、または独自のモデル設定は、入力しない限り含まれます。
a numeric field, which the app does not provide.
計算

または分析なしで動作します。 TELEMETRY_API_KEY を設定した場合は、コピーされます
シナリオ リンクは、次を使用してサーバー ルートを通じて 1 つの集約イベントを発行します。
テレメトリー。
イベントには、選択したプリセットと広いマージンおよび支出バンドのみが含まれます。
価格、トークン数、URL、IP アドレス、シナリオは含まれません。
価値観。
cp .env.example .env.local
# Telemetry API キーを追加し、開発サーバーを再起動します。
検証する
npm タイプチェックを実行する
npmテスト
npm ビルドを実行する
ライセンス
AI 製品のユニットエコノミクス (トークン、キャッシュ、ルーティング、再試行、粗利益) を計算する対話型計算ツールです。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

An interactive calculator for AI product unit economics: tokens, caching, routing, retries, and gross margin. - telemetry-sh/llm-margin-lab

GitHub - telemetry-sh/llm-margin-lab: An interactive calculator for AI product unit economics: tokens, caching, routing, retries, and gross margin. · GitHub
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
telemetry-sh
/
llm-margin-lab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ workflows .github/ workflows app app components components docs docs lib lib test test .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md next-env.d.ts next-env.d.ts next.config.ts next.config.ts package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
An interactive calculator for the unit economics of AI features.
Most AI products can quote token prices. Fewer can answer the question that
matters: how much usage can each customer consume before the business model
breaks?
LLM Margin Lab combines pricing, tokens, retries, prompt caching, model
routing, infrastructure, support, and target gross margin in one shareable
scenario. It runs locally and does not require an API key.
Revenue and contribution profit per active user
Prompt cache hit rate and cache discounts
Routing a percentage of work to a cheaper model
Infrastructure and support costs
Maximum requests per user at a target margin
Margin sensitivity as usage grows
The included presets are illustrative. Replace them with current model prices
and the operating costs of your own stack.
Requires Node.js 20.9 or newer.
npm install
npm run dev
Open http://localhost:3000 .
Use Copy scenario to put every assumption in the URL. The recipient can
edit the scenario without an account or backend. No prompt text, customer
data, or proprietary model configuration is included unless you type it into
a numeric field, which the app does not provide.
The calculator works without analytics. If you set TELEMETRY_API_KEY , copied
scenario links emit one aggregate event through the server route using
Telemetry .
The event contains only the selected preset plus broad margin and spend bands;
it does not contain prices, token counts, URLs, IP addresses, or scenario
values.
cp .env.example .env.local
# Add your Telemetry API key, then restart the dev server.
Verify
npm run typecheck
npm test
npm run build
License
An interactive calculator for AI product unit economics: tokens, caching, routing, retries, and gross margin.
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

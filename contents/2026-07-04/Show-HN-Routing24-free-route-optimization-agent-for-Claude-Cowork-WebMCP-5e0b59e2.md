---
source: "https://github.com/routing24/skill"
hn_url: "https://news.ycombinator.com/item?id=48786769"
title: "Show HN: Routing24 – free route optimization agent for Claude Cowork/WebMCP"
article_title: "GitHub - routing24/skill: AI skill for Routing24 · GitHub"
author: "dennis16384"
captured_at: "2026-07-04T16:54:24Z"
capture_tool: "hn-digest"
hn_id: 48786769
score: 1
comments: 0
posted_at: "2026-07-04T16:43:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Routing24 – free route optimization agent for Claude Cowork/WebMCP

- HN: [48786769](https://news.ycombinator.com/item?id=48786769)
- Source: [github.com](https://github.com/routing24/skill)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T16:43:39Z

## Translation

タイトル: Show HN: Routing24 – Claude Cowork/WebMCP 用の無料ルート最適化エージェント
記事タイトル：GitHub - routing24/スキル：Routing24・GitHub向けAIスキル
説明: Routing24 の AI スキル。 GitHub でアカウントを作成して、routing24/スキル開発に貢献してください。
HN テキスト: 私はしばらくの間、企業向けの無料のルート最適化および計画アプリである https://routing24.com を構築してきました。長い間私を悩ませていたことの 1 つは、ブラウザー エージェントが Routing24 と連携できないことです。 AI には自然なタスクがたくさんあります。 - データの取り込み。ユーザー側 (CSV と Excel) から送られてくるデータを把握します。 - ジオコーディング品質の検証 (ユーザーに何を修正する必要があるのか​​、その住所とその方法を説明するため)。 - ビジネス ユーザーにとっての学習曲線は急峻です (ルート最適化に何を求めるかは理解していますが、モデルのセットアップには多大な時間と労力がかかります)。 - ルート最適化の決定の説明可能性 (どの制約が互いに競合するのか、注文が特定の方法で順序付けされる理由など) - ... これらおよび他の多くのタスクは、エージェント AI 一般およびデスクトップ ビジネス AI アシスタント (特に Claude Cowork) にとって非常に自然なタスクです。私たちは、AI がブラウザーを計測および制御する Chrome CDP を使用して、プロンプトからフロントエンド テストを実行しているときに気づきました。統合テスト用の Routing24 のビルドでは、状態とアクションのモデルが「ウィンドウ」経由で AI テストに完全に公開されており、プロンプトから複雑なテスト ケースを実行するのは素晴らしく簡単でした。驚くべきことに、「javascript_tool」とブラウザ拡張機能を備えたクロード・コワークに「ウィンドウ」経由で公開された同様の縮小状態オブジェクトにより、インポート、ジオコーディング、最適化、エージェントからの説明を魅力的に実行できるようになり、アプリ全体が 1 つの大きなツール + 状態になりました。さらに回った

このアプローチは W3C WebMCP として標準化されつつあります (Chrome ではすでにプレビュー中です)。簡単に言うと、Routing24.com の無料ルート最適化が、Claude Cowork と将来の WebMCP エージェントで機能するようになりました。

記事本文:
GitHub - routing24/スキル: Routing24 用 AI スキル · GitHub
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
ルーティング24
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット 参照 参照 README.md README.md SKILL.md SKILL.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Routing24 ルート オプティマイザー — エージェント スキル
ブラウザ内で自然言語から車両配送ルートを計画し、最適化します。
routing24-optimizer は、Routing24 ルートを駆動するためのエージェント スキルです。
routing24_* WebMCP ツールによる最適化。停留所のリストが表示され、
車両を最適化された共有可能な複数の経由地ルート計画に組み込むことができます。すべてが実行されます
ユーザー自身のブラウザタブのクライアント側。サーバー API や API キーはありません。
パッケージ化されたスキルをダウンロードし、互換性のあるエージェント (Claude / Cowork) に追加します。
最新リリース: routing24.skill
または routing24.com から: https://routing24.com/routing24.skill
SKILL.md — スキル定義 (指示 + 手順)。
References/ — API コントラクト ( api.md )、
機械可読な JSON スキーマ ( schema.json )、および
すぐに評価できる呼び出しスニペット (example.md)。
常に最新の契約は https://routing24.com/llms.txt で提供されます。
Routing24 独自のタイプから生成され、リリース時に公開されます。編集しないでください
手。ソースと問題点: https://routing24.com · ライセンス: プロプライエタリ (SKILL.md を参照)。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
最新 (アプリ v0.1.7)
最新の
2026 年 7 月 4 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI skill for Routing24. Contribute to routing24/skill development by creating an account on GitHub.

I've been building https://routing24.com for a while, a free route optimization and planning app for businesses. One of the things that bothered me for quite a time is inability of browser agents to work with Routing24. We have tons of natural tasks for AI: - data ingestion, to figure out data coming from user's side (csv and Excel); - geocoding quality validation (to explain user what they need to correct and their addresses and how); - steep learning curve for business users (who have sense of what they want from route optimization, but setting up the model takes significant time and effort); - explainability of route optimization decisions (which constraints conflict with each other, why orders are sequenced in a certain way, etc) - ... These and many others are very natural tasks for agentic AI in general as well as for desktop business AI assistants (Claude Cowork in particular). We came up to the realisation when were doing frontend testing from prompts, that used Chrome CDP, where AI instruments and controls the browser. Our build of Routing24 for integration testing has its state and action model fully exposed to AI tests via `window`, and driving complex test cases from prompts was nice and easy. Surprisingly, similar reduced state object exposed via `window` to Claude Cowork equipped with `javascript_tool` and browser extension made it possible to drive import, geocoding, optimization and explanation from agent like a charm, the whole app became one large tool+state. Further it turned our this approach is getting standartized as W3C WebMCP (already in preview in Chrome). Long story short: Routing24.com free route optimization now works in Claude Cowork and and future WebMCP agent!

GitHub - routing24/skill: AI skill for Routing24 · GitHub
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
routing24
/
skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits references references README.md README.md SKILL.md SKILL.md View all files Repository files navigation
Routing24 route optimizer — Agent Skill
Plan and optimize vehicle delivery routes from natural language, right in the browser.
routing24-optimizer is an Agent Skill for driving Routing24 route
optimization through its routing24_* WebMCP tools. It turns a list of stops and
vehicles into an optimized, shareable multi-stop route plan. Everything runs
client-side in the user's own browser tab. There is no server API and no API key.
Download the packaged skill and add it to a compatible agent (Claude / Cowork):
Latest release: routing24.skill
Or from routing24.com: https://routing24.com/routing24.skill
SKILL.md — the skill definition (instructions + procedure).
references/ — API contract ( api.md ),
machine-readable JSON Schema ( schema.json ), and
ready-to-eval call snippets ( examples.md ).
The always-current contract is served at https://routing24.com/llms.txt .
Generated from Routing24's own types and published on release. Do not edit by
hand. Source & issues: https://routing24.com · License: Proprietary (see SKILL.md ).
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
Latest (app v0.1.7)
Latest
Jul 4, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

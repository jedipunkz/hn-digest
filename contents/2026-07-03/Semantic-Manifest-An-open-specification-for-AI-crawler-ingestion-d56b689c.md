---
source: "https://github.com/CKL75/semantic-manifest-specification"
hn_url: "https://news.ycombinator.com/item?id=48774528"
title: "Semantic Manifest – An open specification for AI crawler ingestion"
article_title: "GitHub - CKL75/semantic-manifest-specification: An open data standard and streamable content graph specification (NDJSON) designed to optimize website discovery, content relations, and text ingestion for AI crawlers, LLMs, and RAG engines. · GitHub"
author: "CKL75"
captured_at: "2026-07-03T13:43:16Z"
capture_tool: "hn-digest"
hn_id: 48774528
score: 1
comments: 0
posted_at: "2026-07-03T13:02:42Z"
tags:
  - hacker-news
  - translated
---

# Semantic Manifest – An open specification for AI crawler ingestion

- HN: [48774528](https://news.ycombinator.com/item?id=48774528)
- Source: [github.com](https://github.com/CKL75/semantic-manifest-specification)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T13:02:42Z

## Translation

タイトル: セマンティック マニフェスト – AI クローラー取り込みのオープン仕様
記事のタイトル: GitHub - CKL75/semantic-manifest-specation: AI クローラー、LLM、RAG エンジンの Web サイト検出、コンテンツ関係、テキスト取り込みを最適化するために設計されたオープン データ標準およびストリーミング可能なコンテンツ グラフ仕様 (NDJSON)。 · GitHub
説明: AI クローラー、LLM、RAG エンジンの Web サイト検出、コンテンツ関係、テキスト取り込みを最適化するために設計されたオープン データ標準およびストリーミング可能なコンテンツ グラフ仕様 (NDJSON)。 - CKL75/セマンティックマニフェスト仕様

記事本文:
GitHub - CKL75/semantic-manifest-specation: AI クローラー、LLM、RAG エンジンの Web サイト検出、コンテンツ関係、テキスト取り込みを最適化するために設計されたオープン データ標準およびストリーミング可能なコンテンツ グラフ仕様 (NDJSON)。 · GitHub
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
{{私

セージ }}
CKL75
/
セマンティックマニフェスト仕様
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
CKL75/セマンティックマニフェスト仕様
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット LICENSE.md LICENSE.md README.md README.md semantic-manifest-specation.md semantic-manifest-specation.md semantic-manifest.jsonl semantic-manifest.jsonl すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セマンティックマニフェスト仕様
AI クローラー、LLM、RAG エンジンの Web サイト検出、コンテンツ関係、テキスト取り込みを最適化するために設計されたオープン データ標準およびストリーミング可能なコンテンツ グラフ仕様 (NDJSON)。
従来の Web 標準は、大規模な AI 検索エンジンにとって根本的に壊れています。
サイトマップは、構造コンテキストなしで生の URL を渡します。
JSON-LD は単一ページのスコープに閉じ込められます。
llms.txt ファイルはフラット テキストであり、数千ページに拡大すると大量のコンテキスト ウィンドウ トークンを消費します。
セマンティック マニフェストはこのギャップを埋めます。ストリーミング可能な NDJSON 形式を使用するため、AI クローラーはサイト全体のコンテンツ タイプ、リレーショナル エンティティ、明示的に指定された「マークダウン ツイン」を 1 行ずつ効率的に解析できます。
<head> にセマンティック マニフェストがリンクされている約 58,000 ページのサイトを起動してから数時間以内に、ClaudeBot は 1 秒あたり約 7 URL の速度でサイト全体を取り込みました。
📺 ここでシステムパフォーマンスとクロールログのデモビデオをご覧ください
📖 全仕様を読む (v0.1)
この規格は、High-Velocity Content Engine (HVCE) のネイティブの組み込み構造コンポーネントです。
ライブ本番マニフェストは次の場所で表示できます。
EduStats (58,000 ページ): edustats.app/semantic-manifest.jsonl 。
ハイパーソニック SEO フレームワーク: hypersonicseo.com/semantic-manifest.jsonl
セマンティックマニフェスト固有

イオン (v0.1) © 2026 by Chris Limner / Hypersonic SEO.
この作品には CC0 1.0 Universal のマークが付いています。
AI クローラー、LLM、RAG エンジンの Web サイト検出、コンテンツ関係、テキスト取り込みを最適化するために設計されたオープン データ標準およびストリーミング可能なコンテンツ グラフ仕様 (NDJSON)。
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

An open data standard and streamable content graph specification (NDJSON) designed to optimize website discovery, content relations, and text ingestion for AI crawlers, LLMs, and RAG engines. - CKL75/semantic-manifest-specification

GitHub - CKL75/semantic-manifest-specification: An open data standard and streamable content graph specification (NDJSON) designed to optimize website discovery, content relations, and text ingestion for AI crawlers, LLMs, and RAG engines. · GitHub
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
CKL75
/
semantic-manifest-specification
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
CKL75/semantic-manifest-specification
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits LICENSE.md LICENSE.md README.md README.md semantic-manifest-specification.md semantic-manifest-specification.md semantic-manifest.jsonl semantic-manifest.jsonl View all files Repository files navigation
The Semantic Manifest Specification
An open data standard and streamable content graph specification (NDJSON) designed to optimize website discovery, content relations, and text ingestion for AI crawlers, LLMs, and RAG engines.
Traditional web standards are fundamentally broken for AI search engines at scale.
Sitemaps pass raw URLs without structural context.
JSON-LD is trapped in single-page scopes.
llms.txt files are flat text that consume massive amounts of context window tokens when scaling to thousands of pages.
The Semantic Manifest bridges this gap. It uses a streamable NDJSON format so AI crawlers can parse an entire site's content types, relational entities, and explicitly designated "markdown twins" line-by-line efficiently.
Within hours of launching a ~58,000-page site with a Semantic Manifest linked in the <head> , ClaudeBot ingested the entire site at ~7 URLs per second.
📺 Watch the System Performance & Crawl Log Demo Video Here
📖 Read the Full Specification (v0.1)
This standard is a native, built-in structural component of the High-Velocity Content Engine (HVCE).
Live production manifests can be viewed at:
EduStats (58,000 pages): edustats.app/semantic-manifest.jsonl .
Hypersonic SEO Framework: hypersonicseo.com/semantic-manifest.jsonl
Semantic Manifest Specification (v0.1) © 2026 by Chris Limner / Hypersonic SEO.
This work is marked with CC0 1.0 Universal .
An open data standard and streamable content graph specification (NDJSON) designed to optimize website discovery, content relations, and text ingestion for AI crawlers, LLMs, and RAG engines.
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

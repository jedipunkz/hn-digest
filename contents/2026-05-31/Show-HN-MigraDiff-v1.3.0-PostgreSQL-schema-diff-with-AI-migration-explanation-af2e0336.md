---
source: "https://github.com/migradiff/migra/releases/tag/v1.3.0"
hn_url: "https://news.ycombinator.com/item?id=48338061"
title: "Show HN: MigraDiff v1.3.0 – PostgreSQL schema diff with AI migration explanation"
article_title: "Release v1.3.0 — AI-Powered Migration Explanation & Migrations Folder Support · migradiff/migra · GitHub"
author: "lateos-ai"
captured_at: "2026-05-31T00:27:35Z"
capture_tool: "hn-digest"
hn_id: 48338061
score: 1
comments: 0
posted_at: "2026-05-30T16:34:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: MigraDiff v1.3.0 – PostgreSQL schema diff with AI migration explanation

- HN: [48338061](https://news.ycombinator.com/item?id=48338061)
- Source: [github.com](https://github.com/migradiff/migra/releases/tag/v1.3.0)
- Score: 1
- Comments: 0
- Posted: 2026-05-30T16:34:22Z

## Translation

タイトル: Show HN: MigraDiff v1.3.0 – AI 移行の説明を含む PostgreSQL スキーマの差分
記事のタイトル: リリース v1.3.0 — AI を活用した移行の説明と移行フォルダーのサポート · migradiff/migra · GitHub
説明: アクティブに維持されている migra のフォーク — PostgreSQL スキーマ diff および移行スクリプト ジェネレーター。 - リリース v1.3.0 — AI を活用した移行の説明と移行フォルダーのサポート · migradiff/migra

記事本文:
リリース v1.3.0 — AI を活用した移行の説明と移行フォルダーのサポート · migradiff/migra · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ミグラディフ
/
マイグラ
公共
djrobstep/migra からフォーク
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
v1.3.0 — AI を活用した移行の説明と移行フォルダーのサポート
フィルター
読み込み中
申し訳ありませんが、問題が発生しました。
読み込み中にエラーが発生しました。このページをリロードしてください。
pip install --upgrade migradiff
新機能
AI を活用した移行の説明 (--explain)
MigraDiff はあらゆる移行をわかりやすい英語で説明できるようになりました。
それぞれの変更が何をするのか、それがどのようなリスクを伴うのか、より安全な代替手段は何か
破壊的な操作用。
pip インストール migradiff[ai]
migra --setup-ai
migra --explain postgres://db_a postgres://db_b
クロード俳句 (Anthropic) によって提供されています。独自の API キーを持参する —
MigraDiff サーバーにはデータは送信されません。 --output json で動作します。
--from-file、--from-migrations-dir、およびすべての既存のフラグ。
移行フォルダー入力モード (--from-migrations-dir)
番号付き移行ファイルのディレクトリをベースと比較します。
ライブブランチデータベースを必要とせずにスキーマを構築できます。
migra --from-migrations-dir ./supabase/migrations \
postgres://db_production
Supabase タイムスタンプ形式、Flyway バージョン形式、
および標準の数値プレフィックス。正しい数値で適用されたファイル
ソート順序 (辞書順ではなく、9 から 10)。
README での名前の明確化 — CLI は後方移行のために移行を継続します
互換性、パッケージは migriff
完全な変更ログ: https://github.com/migradiff/migra/blob/main/CHANGELOG.md
pip install --upgrade migradiff
資産
2
読み込み中
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The actively maintained fork of migra — PostgreSQL schema diff and migration script generator. - Release v1.3.0 — AI-Powered Migration Explanation & Migrations Folder Support · migradiff/migra

Release v1.3.0 — AI-Powered Migration Explanation & Migrations Folder Support · migradiff/migra · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
migradiff
/
migra
Public
forked from djrobstep/migra
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
v1.3.0 — AI-Powered Migration Explanation & Migrations Folder Support
Filter
Loading
Sorry, something went wrong.
There was an error while loading. Please reload this page .
pip install --upgrade migradiff
What's New
AI-Powered Migration Explanation (--explain)
MigraDiff can now explain any migration in plain English — what
each change does, what risks it carries, and safer alternatives
for destructive operations.
pip install migradiff[ai]
migra --setup-ai
migra --explain postgres://db_a postgres://db_b
Powered by Claude Haiku (Anthropic). Bring your own API key —
no data is sent to MigraDiff servers. Works with --output json,
--from-file, --from-migrations-dir, and all existing flags.
Migrations Folder Input Mode (--from-migrations-dir)
Diff a directory of numbered migration files against a base
schema without requiring a live branch database.
migra --from-migrations-dir ./supabase/migrations \
postgres://db_production
Supports Supabase timestamp format, Flyway versioned format,
and standard numeric prefixes. Files applied in correct numeric
sort order (9 before 10, not lexicographic).
Naming clarification in README — CLI stays migra for backward
compatibility, package is migradiff
Full changelog: https://github.com/migradiff/migra/blob/main/CHANGELOG.md
pip install --upgrade migradiff
Assets
2
Loading
Uh oh!
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

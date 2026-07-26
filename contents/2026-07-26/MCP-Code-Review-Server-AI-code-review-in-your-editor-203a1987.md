---
source: "https://github.com/GoodJobwilliam/aicraft/tree/main/products/mcp-code-review"
hn_url: "https://news.ycombinator.com/item?id=49058442"
title: "MCP Code Review Server – AI code review in your editor"
article_title: "aicraft/products/mcp-code-review at main · GoodJobwilliam/aicraft · GitHub"
author: "williamXue"
captured_at: "2026-07-26T14:27:44Z"
capture_tool: "hn-digest"
hn_id: 49058442
score: 1
comments: 0
posted_at: "2026-07-26T14:13:19Z"
tags:
  - hacker-news
  - translated
---

# MCP Code Review Server – AI code review in your editor

- HN: [49058442](https://news.ycombinator.com/item?id=49058442)
- Source: [github.com](https://github.com/GoodJobwilliam/aicraft/tree/main/products/mcp-code-review)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T14:13:19Z

## Translation

タイトル: MCP コード レビュー サーバー – エディターでの AI コード レビュー
記事のタイトル: aicraft/products/mcp-code-review at main · GoodJobwilliam/aicraft · GitHub
説明: AICraft - AI を活用した開発者ツールとテンプレート - aicraft/products/mcp-code-review at main · GoodJobwilliam/aicraft

記事本文:
メインの aicraft/products/mcp-code-review · GoodJobwilliam/aicraft · GitHub
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
グッドジョブウィリアム
/
航空機
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション最適化

オン
コード
その他のオプション ディレクトリアクション
歴史 歴史メイン ブレッドクラム
コピーパスのトップフォルダーとファイル
.. src src testing testing LAUNCHGUIDE.md LAUNCHGUIDE.md README.md README.md llms-install.md llms-install.md pyproject.toml pyproject.toml smithery.yaml smithery.yaml uv.lock uv.lock すべてのファイルを表示 README.md
MCP コードレビューサーバーの概要
MCP サーバーとしてのコード レビュー。クロード コード、カーソル、または MCP 互換の AI アシスタントに接続します。
review_code — ソース コード スニペットのバグ、セキュリティ、パフォーマンス、スタイルをレビューします。
review_diff — マージする前に潜在的な問題がないか git diff を確認します。
review_file — パスに基づいてローカル ファイルをレビューする
コード レビュー エージェントと同じ方法論 (OWASP トップ 10 スキャン、N+1 クエリ検出、競合状態分析、重大度評価付きの構造化出力) を利用しています。
# クロード コード MCP 構成に追加します。
クロード mcp コードレビューを追加 -- uvx aicraft-code-review
または、 ~/.cursor/mcp.json または claude_desktop_config.json に追加します。
{
"mcpサーバー": {
"コードレビュー" : {
"コマンド" : " uvx " ,
"args" : [ "mcp-code-review " ]
}
}
}
pip経由
pip インストール aicraft-code-review
python -m mcp_code_review
使用例
接続したら、AI アシスタントに次のように尋ねます。
「セキュリティの問題については、この Python コードを確認してください: [コードを貼り付け]」
「コミットする前にこの差分を確認してください: [差分を貼り付け]」
「このファイルを確認してください: /path/to/file.py」
AI は MCP サーバーを呼び出し、構造化された結果を返します。
## 結果のレビュー
### 🔴 クリティカル (1)
|ライン |問題 |カテゴリー |修正 |
|------|-------|----------|-----|
| 42 | f-string による SQL インジェクション |セキュリティ |パラメータ化されたクエリを使用する |
### 🟠 高 (2)
|ライン |問題 |カテゴリー |修正 |
|------|-------|----------|-----|
| 15 |未検証のユーザー入力 |セキュリティ |入力検証を追加 |
| 78 |ループ内の N+1 クエリ |パフォーマンス |選択関連を追加 |
#

## 概要
- **重大**: 1 — 修正する必要があります
- **高**: 2 — 修正する必要があります
- **中**: 0
- **情報**: 0
開発
git clone https://github.com/GoodJobwilliam/aicraft
CD飛行機
pip install -e " .[dev] "
python -m mcp_code_review # サーバーを起動します
要件
MCP 互換クライアント (クロード コード、カーソルなど)
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AICraft - AI-powered developer tools and templates - aicraft/products/mcp-code-review at main · GoodJobwilliam/aicraft

aicraft/products/mcp-code-review at main · GoodJobwilliam/aicraft · GitHub
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
GoodJobwilliam
/
aicraft
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
More options Directory actions
History History main Breadcrumbs
Copy path Top Folders and files
.. src src tests tests LAUNCHGUIDE.md LAUNCHGUIDE.md README.md README.md llms-install.md llms-install.md pyproject.toml pyproject.toml smithery.yaml smithery.yaml uv.lock uv.lock View all files README.md
Outline MCP Code Review Server
Code review as an MCP server. Connect it to Claude Code, Cursor, or any MCP-compatible AI assistant.
review_code — Review any source code snippet for bugs, security, performance, and style
review_diff — Review a git diff for potential issues before merging
review_file — Review a local file by path
Powered by the same methodology as our Code Review Agent: OWASP Top 10 scanning, N+1 query detection, race condition analysis, and structured output with severity ratings.
# Add to your Claude Code MCP config:
claude mcp add code-review -- uvx aicraft-code-review
Or add to your ~/.cursor/mcp.json or claude_desktop_config.json :
{
"mcpServers" : {
"code-review" : {
"command" : " uvx " ,
"args" : [ " mcp-code-review " ]
}
}
}
Via pip
pip install aicraft-code-review
python -m mcp_code_review
Usage Examples
Once connected, ask your AI assistant:
"Review this Python code for security issues: [paste code]"
"Review this diff before I commit: [paste diff]"
"Review this file: /path/to/file.py"
The AI will call the MCP server and return structured results.
## Review Results
### 🔴 Critical (1)
| Line | Issue | Category | Fix |
|------|-------|----------|-----|
| 42 | SQL injection via f-string | Security | Use parameterized queries |
### 🟠 High (2)
| Line | Issue | Category | Fix |
|------|-------|----------|-----|
| 15 | Unvalidated user input | Security | Add input validation |
| 78 | N+1 query in loop | Performance | Add select_related |
### Summary
- **Critical**: 1 — must fix
- **High**: 2 — should fix
- **Medium**: 0
- **Info**: 0
Development
git clone https://github.com/GoodJobwilliam/aicraft
cd aicraft
pip install -e " .[dev] "
python -m mcp_code_review # Start server
Requirements
An MCP-compatible client (Claude Code, Cursor, etc.)
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

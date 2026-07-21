---
source: "https://github.com/Karthikvk1899/guard-ai"
hn_url: "https://news.ycombinator.com/item?id=48995348"
title: "Guard-AI – A security linter for AI-generated code"
article_title: "GitHub - Karthikvk1899/guard-ai · GitHub"
author: "karthikvk1899"
captured_at: "2026-07-21T18:11:00Z"
capture_tool: "hn-digest"
hn_id: 48995348
score: 1
comments: 0
posted_at: "2026-07-21T17:23:56Z"
tags:
  - hacker-news
  - translated
---

# Guard-AI – A security linter for AI-generated code

- HN: [48995348](https://news.ycombinator.com/item?id=48995348)
- Source: [github.com](https://github.com/Karthikvk1899/guard-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T17:23:56Z

## Translation

タイトル: Guard-AI – AI 生成コード用のセキュリティ リンター
記事タイトル: GitHub - Karthikvk1899/guard-ai · GitHub
説明: GitHub でアカウントを作成して、Karthikvk1899/guard-ai の開発に貢献します。

記事本文:
GitHub - Karthikvk1899/guard-ai · GitHub
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
Karthikvk1899
/
ガードアイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ f に移動

ファイル コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github/ workflows .github/ workflows dist dist src/guard_ai src/guard_ai README.md README.md pre-commit-hooks.yaml pre-commit-hooks.yaml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI によって生成された脆弱性、幻覚的な依存関係、コードの切り捨てを本番環境に影響を与える前に検出する自動リンター。
問題: AI のしゃがみと暗号の幻覚
大規模言語モデル (LLM) は優れたコードを記述しますが、次のような明確なセキュリティ リスクももたらします。
スロップスクワッティング (幻覚パッケージ): LLM は、存在しない PyPI および NPM の依存関係を頻繁に作成します (例: fake_llm_helper_12345 )。攻撃者はこれらの幻覚名を登録して、サプライチェーン攻撃を実行します。
ハードコードされたシークレット プレースホルダー: LLM は、不足しているキーを openai_key = "your_openai_key_here" のようなパターンで埋めますが、これは開発者が誤ってコミットすることがよくあります。
切り詰められたコード コメント: AI モデルは、 # ... ここのコードの残りの部分、または // ... 既存のコード を含む不完全なスニペットを出力することが多く、本番環境のロジックを破壊します。
Guard-ai は CLI、コミット前フック、CI/CD パイプラインで直接実行され、これらの問題を即座にブロックします。
多言語サポート: Python ( .py )、JavaScript/TypeScript ( .js 、 .ts 、 .jsx 、 .tsx )、および package.json をスキャンします。
高性能キャッシュ: SQLite ベースのローカル レジストリ ルックアップ エンジン (24 時間 TTL) により、レート制限に達することなく PyPI および NPM をクエリします。
誤検知ゼロ: Python 標準ライブラリ、ノード組み込み、ローカル モジュール構造のスマート フィルタリング。
CI/CD 対応: 自動パイプライン用のカスタム終了しきい値 ( --fail-on HIGH ) と構造化された --json 出力。
無視ルール エンジン: デフォルトの無視 ( .venv 、 node_modules )、 .gitignore 、および .guardignore のシームレスなサポート。
pip インストールガード -

アイクリ
使用法
バッシュ
# 現在のディレクトリの基本スキャン
ガードアイスキャン。
# 重大度の高い問題でパイプラインが失敗する
ガードアイスキャン。 --フェイルオン HIGH
# 結果をJSON形式で出力する
ガードアイスキャン。 --json
プリコミットフックの統合
コミットする前に AI のリスクを把握したいですか? Guard-ai をプロジェクトの .pre-commit-config.yaml に追加します。
YAML
リポジトリ:
- リポジトリ: [https://github.com/Karthikvk1899/guard-ai](https://github.com/Karthikvk1899/guard-ai)
リビジョン: v0.1.1
フック:
- id: ガードアイ
引数: ["--フェイルオン"、"HIGH"]
について
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

Contribute to Karthikvk1899/guard-ai development by creating an account on GitHub.

GitHub - Karthikvk1899/guard-ai · GitHub
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
Karthikvk1899
/
guard-ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github/ workflows .github/ workflows dist dist src/ guard_ai src/ guard_ai README.md README.md pre-commit-hooks.yaml pre-commit-hooks.yaml pyproject.toml pyproject.toml View all files Repository files navigation
Automated linter that catches AI-generated vulnerabilities, hallucinated dependencies, and code truncations before they hit production.
The Problem: AI Slopsquatting & Code Hallucinations
Large Language Models (LLMs) write great code, but they also introduce distinct security risks:
Slopsquatting (Hallucinated Packages): LLMs frequently invent non-existent PyPI and NPM dependencies (e.g., fake_llm_helper_12345 ). Attackers register these hallucinated names to execute supply-chain attacks.
Hardcoded Secret Placeholders: LLMs fill missing keys with patterns like openai_key = "your_openai_key_here" , which developers often accidentally commit.
Truncated Code Comments: AI models often output incomplete snippets containing # ... rest of code here or // ... existing code , breaking logic in production.
guard-ai runs directly in your CLI, pre-commit hooks, and CI/CD pipelines to block these issues instantly.
Multi-Language Support: Scans Python ( .py ), JavaScript/TypeScript ( .js , .ts , .jsx , .tsx ), and package.json .
High Performance Caching: SQLite-backed local registry lookup engine (24-hr TTL) to query PyPI & NPM without hitting rate limits.
Zero False-Positives: Smart filtering for Python standard libraries, Node built-ins, and local module structures.
CI/CD Ready: Custom exit thresholds ( --fail-on HIGH ) and structured --json output for automated pipelines.
Ignore Rule Engine: Seamless support for default ignores ( .venv , node_modules ), .gitignore , and .guardignore .
pip install guard-ai-cli
Usage
Bash
# Basic scan on current directory
guard-ai scan .
# Fail pipeline on high-severity issues
guard-ai scan . --fail-on HIGH
# Output results in JSON format
guard-ai scan . --json
Pre-Commit Hook Integration
Want to catch AI risks before you commit ? Add guard-ai to your project ' s .pre-commit-config.yaml:
YAML
repos:
- repo: [https://github.com/Karthikvk1899/guard-ai](https://github.com/Karthikvk1899/guard-ai)
rev: v0.1.1
hooks:
- id: guard-ai
args: ["--fail-on", "HIGH"]
About
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

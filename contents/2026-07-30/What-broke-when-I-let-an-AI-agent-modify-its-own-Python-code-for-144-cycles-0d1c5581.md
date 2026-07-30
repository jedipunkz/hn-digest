---
source: "https://github.com/ADevBelgie/zmb-audit"
hn_url: "https://news.ycombinator.com/item?id=49110869"
title: "What broke when I let an AI agent modify its own Python code for 144 cycles"
article_title: "GitHub - ADevBelgie/zmb-audit · GitHub"
author: "ArthurDev"
captured_at: "2026-07-30T15:02:53Z"
capture_tool: "hn-digest"
hn_id: 49110869
score: 2
comments: 0
posted_at: "2026-07-30T14:49:41Z"
tags:
  - hacker-news
  - translated
---

# What broke when I let an AI agent modify its own Python code for 144 cycles

- HN: [49110869](https://news.ycombinator.com/item?id=49110869)
- Source: [github.com](https://github.com/ADevBelgie/zmb-audit)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T14:49:41Z

## Translation

タイトル: AI エージェントに自身の Python コードを 144 サイクル変更させたら何が壊れたのか
記事タイトル: GitHub - ADevBelgie/zmb-audit · GitHub
説明: GitHub でアカウントを作成して、ADevBelgie/zmb-audit の開発に貢献します。

記事本文:
GitHub - ADevBelgie/zmb-audit · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
ADevベルギー
/
ZMB監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード

「その他のアクション」メニューを開く フォルダーとファイル
1 コミット 1 コミット LICENSE LICENSE README.md README.md zmb_audit.py zmb_audit.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ZMB スタンドアロン コードベース監査 ( zmb_audit.py )
依存関係のない、Python コードベース用のスタンドアロン静的分析ツール。
zmb_audit.py は、Python の標準ライブラリ ast パーサーを使用してターゲット Python リポジトリをスキャンし、構造的なガバナンスのリスクとコードの腐敗を特定します。
孤立したモジュール : コードベースに存在するが、何もインポートされていない Python モジュールを検索します。
未解決の属性呼び出し : ファイル全体で、欠落しているシンボルまたは名前が変更されたシンボルに対する不審なメソッド呼び出しを検出します。
Kernel Guard Status : git pre-commit フック保護がアクティブかどうかを確認します。
要件: Python 3.8+ (Stdlib のみ - pip 依存関係は必要ありません)。
セットアップ: zmb_audit.py をダウンロードして直接実行します。
# 現在のディレクトリを監査する
Python zmb_audit.py
# 特定のリポジトリ パスを監査する
python zmb_audit.py /path/to/repository
# 結果を構造化された JSON として出力する
python zmb_audit.py /path/to/repository --json
# 特定のカスタムディレクトリをスキャンから除外する
python zmb_audit.py /path/to/repository --exclude Legacy_code Experimental
詳細レポート
8 つの自律的自己修正故障モードと生産到達可能性計測の完全な実証分析については、以下を参照してください。
👉 ZMB 障害モード レポートを読む
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to ADevBelgie/zmb-audit development by creating an account on GitHub.

GitHub - ADevBelgie/zmb-audit · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
ADevBelgie
/
zmb-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit LICENSE LICENSE README.md README.md zmb_audit.py zmb_audit.py View all files Repository files navigation
ZMB Standalone Codebase Auditor ( zmb_audit.py )
A zero-dependency, standalone static analysis tool for Python codebases.
zmb_audit.py scans any target Python repository using Python's standard library ast parser to identify structural governance risks and code rot:
Orphan Modules : Finds Python modules that exist in the codebase but are imported by nothing.
Unresolved Attribute Calls : Detects suspicious method calls across files to missing or renamed symbols.
Kernel Guard Status : Verifies whether git pre-commit hook protection is active.
Requirements : Python 3.8+ (Stdlib only — no pip dependencies required).
Setup : Download zmb_audit.py and run it directly.
# Audit current directory
python zmb_audit.py
# Audit specific repository path
python zmb_audit.py /path/to/repository
# Output results as structured JSON
python zmb_audit.py /path/to/repository --json
# Exclude specific custom directories from scan
python zmb_audit.py /path/to/repository --exclude legacy_code experimental
Deep-Dive Report
For full empirical analysis of 8 autonomous self-modification failure modes and production reachability instrumentation:
👉 Read the ZMB Failure-Mode Report
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

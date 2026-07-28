---
source: "https://github.com/anthropics/cryptography-research-demo"
hn_url: "https://news.ycombinator.com/item?id=49090083"
title: "Anthropic publishes a practical key-recovery attack on HAWK-256"
article_title: "GitHub - anthropics/cryptography-research-demo · GitHub"
author: "bakigul"
captured_at: "2026-07-28T22:04:57Z"
capture_tool: "hn-digest"
hn_id: 49090083
score: 11
comments: 1
posted_at: "2026-07-28T21:22:48Z"
tags:
  - hacker-news
  - translated
---

# Anthropic publishes a practical key-recovery attack on HAWK-256

- HN: [49090083](https://news.ycombinator.com/item?id=49090083)
- Source: [github.com](https://github.com/anthropics/cryptography-research-demo)
- Score: 11
- Comments: 1
- Posted: 2026-07-28T21:22:48Z

## Translation

タイトル: Anthropic が HAWK-256 に対する実用的なキー回復攻撃を公開
記事のタイトル: GitHub - 人類学/暗号研究デモ · GitHub
説明: GitHub でアカウントを作成して、人類学/暗号研究デモの開発に貢献します。

記事本文:
GitHub - 人類学/暗号研究デモ · GitHub
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
人類学
/
暗号研究デモ
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット AES AES HAWK HAWK LEA LEA .gitignore .gitignore ライセンス ライセンス通知 通知 README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
研究成果物。メンテナンスされておらず、寄付も受け付けていません。
関連論文に付属する暗号解析コード。 3つの独立した
コンポーネント:
Apache 2.0 に基づいてライセンスが付与されています (「ライセンス」を参照)。
Readme Apache-2.0 ライセンス アクティビティ カスタム プロパティ スター
2 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to anthropics/cryptography-research-demo development by creating an account on GitHub.

GitHub - anthropics/cryptography-research-demo · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
anthropics
/
cryptography-research-demo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit AES AES HAWK HAWK LEA LEA .gitignore .gitignore LICENSE LICENSE NOTICE NOTICE README.md README.md View all files Repository files navigation
Research artifact. Not maintained and not accepting contributions.
Cryptanalysis code accompanying the associated papers. Three independent
components:
Licensed under Apache 2.0 (see LICENSE ).
Readme Apache-2.0 license Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

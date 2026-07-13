---
source: "https://github.com/JustiNoel/LLM-Prompt-Injection"
hn_url: "https://news.ycombinator.com/item?id=48895452"
title: "I built a prompt injection defense middleware for LLMs (Python/FastAPI)"
article_title: "GitHub - JustiNoel/LLM-Prompt-Injection: A .py code that acts as a middleware between the user and the LLM before producing the output · GitHub"
author: "its_me_noel"
captured_at: "2026-07-13T17:07:33Z"
capture_tool: "hn-digest"
hn_id: 48895452
score: 1
comments: 0
posted_at: "2026-07-13T16:54:16Z"
tags:
  - hacker-news
  - translated
---

# I built a prompt injection defense middleware for LLMs (Python/FastAPI)

- HN: [48895452](https://news.ycombinator.com/item?id=48895452)
- Source: [github.com](https://github.com/JustiNoel/LLM-Prompt-Injection)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T16:54:16Z

## Translation

タイトル: LLM 用のプロンプト インジェクション防御ミドルウェアを構築しました (Python/FastAPI)
記事のタイトル: GitHub - JustiNoel/LLM-Prompt-Injection: 出力を生成する前にユーザーと LLM の間のミドルウェアとして機能する .py コード · GitHub
説明: 出力を生成する前にユーザーと LLM の間のミドルウェアとして機能する .py コード - JustiNoel/LLM-Prompt-Injection

記事本文:
GitHub - JustiNoel/LLM-Prompt-Injection: 出力を生成する前にユーザーと LLM の間のミドルウェアとして機能する .py コード · GitHub
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
ジャスティノエル
/
LLM プロンプト インジェクション
公共
通知
通知を変更するにはサインインする必要があります

設定上
追加のナビゲーション オプション
コード
JustiNoel/LLM-プロンプトインジェクション
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット docker-compose.yml docker-compose.yml .github/ workflows .github/ workflows .vscode .vscode Prompt-Shield Prompt-Shield README.md README.md benchmark_results.json benchmark_results.json docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
実稼働グレードの LLM プロンプト インジェクション防御ミドルウェア。
PromptShield はユーザーと AI モデルの間に位置し、損害が発生する前に敵対的な攻撃を検出してブロックします。
PromptShield は、すべてのユーザー入力を 4 つの防御層を通じて実行します。
✅ 100% 攻撃検出、誤検知ゼロ (ベンチマークテスト済み)
⚙️ 調整可能な攻撃性ダイヤル — 寛容→バランスの取れた→厳格→偏執的
🔑 APIキー認証+レート制限
📊 監査ログとダッシュボード UI
🐍 簡単な統合のための Python SDK
git クローン https://github.com/JustiNoel/LLM-Prompt-Injection.git
cd LLM プロンプトインジェクション
ドッカー構成アップ
ライブデモ
Python · FastAPI · Google Gemini · Docker
出力を生成する前にユーザーと LLM の間のミドルウェアとして機能する .py コード
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A .py code that acts as a middleware between the user and the LLM before producing the output - JustiNoel/LLM-Prompt-Injection

GitHub - JustiNoel/LLM-Prompt-Injection: A .py code that acts as a middleware between the user and the LLM before producing the output · GitHub
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
JustiNoel
/
LLM-Prompt-Injection
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
JustiNoel/LLM-Prompt-Injection
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits docker-compose.yml docker-compose.yml .github/ workflows .github/ workflows .vscode .vscode prompt-shield prompt-shield README.md README.md benchmark_results.json benchmark_results.json docker-compose.yml docker-compose.yml View all files Repository files navigation
Production-grade LLM prompt injection defense middleware.
PromptShield sits between your users and your AI model, detecting and blocking adversarial attacks before they cause damage.
PromptShield runs every user input through 4 layers of defense:
✅ 100% attack detection, zero false positives (benchmark tested)
⚙️ Tunable aggression dial — permissive → balanced → strict → paranoid
🔑 API key authentication + rate limiting
📊 Audit logging and dashboard UI
🐍 Python SDK for easy integration
git clone https://github.com/JustiNoel/LLM-Prompt-Injection.git
cd LLM-Prompt-Injection
docker-compose up
Live Demo
Python · FastAPI · Google Gemini · Docker
A .py code that acts as a middleware between the user and the LLM before producing the output
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

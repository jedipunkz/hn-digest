---
source: "https://github.com/ywu593412-afk/difflens"
hn_url: "https://news.ycombinator.com/item?id=48346869"
title: "DiffLens: Using LangGraph to verify line-range coordinates in LLM code reviews"
article_title: "GitHub - ywu593412-afk/difflens: Automated code review workflow using LangGraph to fix line-number hallucinations. · GitHub"
author: "wwyb"
captured_at: "2026-06-01T00:28:18Z"
capture_tool: "hn-digest"
hn_id: 48346869
score: 2
comments: 0
posted_at: "2026-05-31T16:10:07Z"
tags:
  - hacker-news
  - translated
---

# DiffLens: Using LangGraph to verify line-range coordinates in LLM code reviews

- HN: [48346869](https://news.ycombinator.com/item?id=48346869)
- Source: [github.com](https://github.com/ywu593412-afk/difflens)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T16:10:07Z

## Translation

タイトル: DiffLens: LangGraph を使用して LLM コード レビューでライン範囲座標を検証する
記事のタイトル: GitHub - ywu593412-afk/difflens: LangGraph を使用して行番号の幻覚を修正する自動コード レビュー ワークフロー。 · GitHub
説明: LangGraph を使用して行番号の幻覚を修正する自動コード レビュー ワークフロー。 - ywu593412-afk/ディフレンズ

記事本文:
GitHub - ywu593412-afk/difflens: LangGraph を使用して行番号の幻覚を修正する自動コード レビュー ワークフロー。 · GitHub
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
ywu593412-afk
/
ディフレンス
公共
通知
変更するにはサインインする必要があります

化設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github/ workflows .github/ workflows difflens difflens 例 例 テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Git の差分をレビューするために LangGraph で構築された軽量の TypeScript ツール。境界外の行幻覚を分離してフィルタリングするように設計された追加の検証手順が含まれています。
注: このプロジェクトは、単一プロンプトのレビュー スクリプトを超えて、最終出力を表示する前に構造化パイプラインが壊れた座標参照を停止できるかどうかを確認する実験として始まりました。
ステータス: コア エンジンは完全に機能しており、自動化された CI パイプラインによって検証されています。
グラフTD
A[Git Diff入力] --> B[パーサーノード: 行範囲の抽出]
B --> C[LLM 生成ノード: コメントの確認]
C --> D[検証ノード:座標チェック]
D -- 境界外の回線 --> E[条件付きルーター: エラー フィードバック]
E --> C
D -- 有効な座標 --> F[最終承認レポート]
読み込み中
について
LangGraph を使用して行番号の幻覚を修正する自動コード レビュー ワークフロー。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.0 - コア エンジン リリース
最新の
2026 年 5 月 31 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Automated code review workflow using LangGraph to fix line-number hallucinations. - ywu593412-afk/difflens

GitHub - ywu593412-afk/difflens: Automated code review workflow using LangGraph to fix line-number hallucinations. · GitHub
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
ywu593412-afk
/
difflens
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github/ workflows .github/ workflows difflens difflens examples examples tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A lightweight TypeScript tool built with LangGraph for reviewing Git diffs, with additional verification steps designed to isolate and filter out-of-bounds line hallucinations.
Note: This project started as an experiment to move beyond single-prompt review scripts and see if a structured pipeline can stop broken coordinate references before displaying the final output.
Status: Core engine fully functional and verified via automated CI pipeline.
graph TD
A[Git Diff Input] --> B[Parser Node: Extract Line Range]
B --> C[LLM Generation Node: Review Comments]
C --> D[Verifier Node: Coordinate Check]
D -- Line Out of Bounds --> E[Conditional Router: Error Feedback]
E --> C
D -- Valid Coordinates --> F[Final Approved Report]
Loading
About
Automated code review workflow using LangGraph to fix line-number hallucinations.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.0 - Core Engine Release
Latest
May 31, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

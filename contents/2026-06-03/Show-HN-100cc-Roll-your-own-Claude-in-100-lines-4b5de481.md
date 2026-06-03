---
source: "https://github.com/yujqiao/100cc"
hn_url: "https://news.ycombinator.com/item?id=48374308"
title: "Show HN: 100cc - Roll your own Claude in 100 lines"
article_title: "GitHub - yujqiao/100cc: Roll your own claude code with 100 lines · GitHub"
author: "rapiz"
captured_at: "2026-06-03T00:38:32Z"
capture_tool: "hn-digest"
hn_id: 48374308
score: 7
comments: 0
posted_at: "2026-06-02T18:35:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: 100cc - Roll your own Claude in 100 lines

- HN: [48374308](https://news.ycombinator.com/item?id=48374308)
- Source: [github.com](https://github.com/yujqiao/100cc)
- Score: 7
- Comments: 0
- Posted: 2026-06-02T18:35:06Z

## Translation

タイトル: Show HN: 100cc - 自分のクロードを 100 行で転がす
記事タイトル: GitHub - yujqiao/100cc: 100 行で独自のクロード コードをロールする · GitHub
説明: 100 行の独自のクロード コードをロールします。 GitHub でアカウントを作成して、yujqiao/100cc の開発に貢献してください。

記事本文:
GitHub - yujqiao/100cc: 100 行の独自のクロード コードをロールアップする · GitHub
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
玉橋
/
100cc
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
master ブランチ タグ ファイルに移動 コード 操作

その他のアクション メニュー フォルダーとファイル
1 コミット 1 コミット docs/ スクリーンショット docs/ スクリーンショット src src .env.example .env.example .gitignore .gitignore README.md README.md TODO.md TODO.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
100 行の独自のコーディング エージェントをロールアップする
考え方はシンプルです。必要最小限のハーネスを手書きで作成し、それ自体を作成するように依頼します。
バンインストール
以下を使用して .env を作成します。
OPENAI_API_KEY=... # 必須
OPENAI_MODEL=... # 必須
OPENAI_BASE_URL=... # オプション、デフォルトは OpenAI の公式エンドポイント
走る
それは必要最小限のものから始まります。非対話モード ( -p ) のみが実装されています
bun start -- -p " このプロジェクトをレビューし、README.md にジョークを 3 つ追加します "
自身に書き込むよう依頼する
bun start -- -c -p "このプロジェクトに対話型モードを実装します" # -c 最後のセッションから続行します
bun start -- -c -p "このプロジェクトをより見栄えよくする"
100cc に TODO.md を実装するよう依頼します。
楽しんで友達と共有しましょう！
このリポジトリは、信号対ノイズ比を高く保ち、出力を減らすことを堅持しようとします。
100 行の独自のクロード コードをロールアップします。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Roll your own claude code with 100 lines. Contribute to yujqiao/100cc development by creating an account on GitHub.

GitHub - yujqiao/100cc: Roll your own claude code with 100 lines · GitHub
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
yujqiao
/
100cc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit docs/ screenshot docs/ screenshot src src .env.example .env.example .gitignore .gitignore README.md README.md TODO.md TODO.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Roll your own coding agent with 100 lines
The idea is simple. Write the bare minimal harness by hand and then ask it to write itself.
bun install
Create a .env with:
OPENAI_API_KEY=... # required
OPENAI_MODEL=... # required
OPENAI_BASE_URL=... # optional, defaults to OpenAI's official endpoint
Run
It starts with the bare minimal. Only non-interactive mode( -p ) is implemented
bun start -- -p " review this project and add 3 jokes to README.md "
Ask it to write itself
bun start -- -c -p " implement interactive mode for this project " # -c to continue from the last session
bun start -- -c -p " make this project look nicer "
Ask your 100cc to implement TODO.md
Have fun and share with your friends!
This repo tries to keep the signal to noise ratio high and adhere to showing less output
Roll your own claude code with 100 lines
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

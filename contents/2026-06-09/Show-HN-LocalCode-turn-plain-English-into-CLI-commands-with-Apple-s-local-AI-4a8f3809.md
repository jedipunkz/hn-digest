---
source: "https://github.com/localcodeai/localcode"
hn_url: "https://news.ycombinator.com/item?id=48467746"
title: "Show HN: LocalCode – turn plain English into CLI commands with Apple's local AI"
article_title: "GitHub - localcodeai/localcode: Turn natural language into CLI commands using Apple's on-device AI · GitHub"
author: "rahlokzero"
captured_at: "2026-06-09T21:37:56Z"
capture_tool: "hn-digest"
hn_id: 48467746
score: 1
comments: 0
posted_at: "2026-06-09T21:04:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LocalCode – turn plain English into CLI commands with Apple's local AI

- HN: [48467746](https://news.ycombinator.com/item?id=48467746)
- Source: [github.com](https://github.com/localcodeai/localcode)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T21:04:24Z

## Translation

タイトル: Show HN: LocalCode – Apple のローカル AI を使用して普通の英語を CLI コマンドに変換します
記事タイトル: GitHub - localcodeai/localcode: Apple のオンデバイス AI を使用して自然言語を CLI コマンドに変換する · GitHub
説明: Apple のオンデバイス AI - localcodeai/localcode を使用して、自然言語を CLI コマンドに変換します。

記事本文:
GitHub - localcodeai/localcode: Apple のオンデバイス AI を使用して自然言語を CLI コマンドに変換する · GitHub
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
ローカルコードアイ
/
ローカルコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
タラ

e
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット LocalCode LocalCode セッション セッション .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md pre-commit.sh pre-commit.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Apple のオンデバイス AI を使用して、自然言語を CLI コマンドに変換します。
LocalCode は、Apple の Foundation Models フレームワークを使用して CLI ツールを構築することを実証する概念実証です。わかりやすい英語で希望を伝えると、適切なコマンドが提案され、承認すると実行されます。
すべての AI 処理は Mac 上でローカルに行われます。クラウドがなく、マシンからデータが流出することもありません。
Apple Silicon Mac (M1/M2/M3/M4)
Xcode 26+ (Swift ヘルパーの構築用)
git clone https://github.com/localcodeai/localcode.git
cd ローカルコード/ローカルコード
# Swift AFM ヘルパーをビルドする
cd ソース/afmhelper
swiftc -o afmhelper main.swift -framework FoundationModels -target arm64-apple-macosx26.0
# Go TUI を構築する
CD ../..
build -o localcode に進みます。
# 実行
./ローカルコード
仕組み
あなた: 「すべての Python ファイルをリストします」
ローカルコード: を見つけます。 -name "*.py"
▶ このコマンドを実行しますか? [Enter] 実行 | [N] キャンセル
モデルはリクエストをコマンドに変換します。実行する前に承認します。
「このディレクトリ内の最大のファイルを表示します」
「このディレクトリ内のすべてのファイルをカウントします」
「ポート8080が使用中かどうかを確認してください」
「このディレクトリ内の hello を grep で実行」
すべてのコマンドは、実行前に確認ダイアログを表示します。実行する前にコマンドを編集することも、キャンセルすることもできます。
ローカルコード/
§── main.go # TUI (タピオカティー) に行く
§── pre-commit.sh # プリコミットフック
└── 出典/
━── afmhelper/ # Swift → Apple FoundationModels
━━ main.swift
TUI レイヤー : ゴー + タピオカティー
AI レイヤー: Apple FoundationModels フレームワーク (Swift)
コマンドフロー : モデルが提案 → あなたが承認 → コマンド

実行する
プリコミットフックにより、プロジェクトが確実にビルドされます。
cp pre-commit.sh .git/hooks/pre-commit && chmod +x .git/hooks/pre-commit
git config core.hooksPath .git/hooks
なぜこれが存在するのでしょうか?
Apple の Foundation Models フレームワークは新しく、文書化が不十分です。このプロジェクトは、CLI ツールで実行可能であることを証明し、AFM 上に構築する他のプロジェクトにリファレンス実装を提供します。
POC には荒削りな部分があります。貢献を歓迎します:
フォークして機能ブランチを作成する
Apple のオンデバイス AI を使用して自然言語を CLI コマンドに変換する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

Turn natural language into CLI commands using Apple's on-device AI - localcodeai/localcode

GitHub - localcodeai/localcode: Turn natural language into CLI commands using Apple's on-device AI · GitHub
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
localcodeai
/
localcode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits LocalCode LocalCode sessions sessions .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md pre-commit.sh pre-commit.sh View all files Repository files navigation
Turn natural language into CLI commands using Apple's on-device AI.
LocalCode is a proof-of-concept that demonstrates building CLI tools with Apple's Foundation Models framework. Tell it what you want in plain English, it suggests the right command, you approve, and it runs.
All AI processing happens locally on your Mac. No cloud, no data leaving your machine.
Apple Silicon Mac (M1/M2/M3/M4)
Xcode 26+ (for building the Swift helper)
git clone https://github.com/localcodeai/localcode.git
cd localcode/LocalCode
# Build the Swift AFM helper
cd Sources/afmhelper
swiftc -o afmhelper main.swift -framework FoundationModels -target arm64-apple-macosx26.0
# Build the Go TUI
cd ../..
go build -o localcode .
# Run
./localcode
How It Works
You: "list all python files"
LocalCode: find . -name "*.py"
▶ Execute this command? [Enter] Run | [N] Cancel
The model translates your request into a command. You approve before it runs.
"show me the largest files in this directory"
"count all files in this directory"
"check if port 8080 is in use"
"grep for hello in this directory"
All commands show a confirmation dialog before running. You can edit the command before executing, or cancel.
LocalCode/
├── main.go # Go TUI (Bubble Tea)
├── pre-commit.sh # Pre-commit hook
└── Sources/
└── afmhelper/ # Swift → Apple FoundationModels
└── main.swift
TUI Layer : Go + Bubble Tea
AI Layer : Apple FoundationModels framework (Swift)
Command Flow : Model suggests → You approve → Command executes
A pre-commit hook ensures the project builds:
cp pre-commit.sh .git/hooks/pre-commit && chmod +x .git/hooks/pre-commit
git config core.hooksPath .git/hooks
Why does this exist?
Apple's Foundation Models framework is new and under-documented. This project proves it's viable for CLI tools and provides a reference implementation for others building on AFM.
POCs have rough edges. Contributions welcome:
Fork and create a feature branch
Turn natural language into CLI commands using Apple's on-device AI
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

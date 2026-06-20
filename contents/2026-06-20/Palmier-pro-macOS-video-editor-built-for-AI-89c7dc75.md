---
source: "https://github.com/palmier-io/palmier-pro"
hn_url: "https://news.ycombinator.com/item?id=48610706"
title: "Palmier-pro: macOS video editor built for AI"
article_title: "GitHub - palmier-io/palmier-pro: macOS video editor built for AI · GitHub"
author: "nateb2022"
captured_at: "2026-06-20T18:36:21Z"
capture_tool: "hn-digest"
hn_id: 48610706
score: 3
comments: 0
posted_at: "2026-06-20T16:48:52Z"
tags:
  - hacker-news
  - translated
---

# Palmier-pro: macOS video editor built for AI

- HN: [48610706](https://news.ycombinator.com/item?id=48610706)
- Source: [github.com](https://github.com/palmier-io/palmier-pro)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T16:48:52Z

## Translation

タイトル: Palmier-pro: AI 向けに構築された macOS ビデオエディター
記事タイトル: GitHub - palmier-io/palmier-pro: AI 用に構築された macOS ビデオエディタ · GitHub
説明: AI 用に構築された macOS ビデオ エディター。 GitHub でアカウントを作成して、palmier-io/palmier-pro の開発に貢献してください。

記事本文:
GitHub - palmier-io/palmier-pro: AI 用に構築された macOS ビデオエディタ · GitHub
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
パルミエ・イオ
/
パルミエプロ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode Open more actions menu Folders and files
633 コミット 633 コミット Sources/ PalmierPro Sources/ PalmierPro Tests/ PalmierProTests Tests/ PalmierProTests アセット アセット mcpb mcpb モデル モデル スクリプト スクリプト .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md FAQ.md FAQ.md ライセンス ライセンス Package.resolved Package.resolved Package.swift Package.swift README.md README.md appcast.xml appcast.xml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
The video editor built for AI.
Requires macOS 26 (Tahoe) on Apple Silicon
Palmier Pro は、Mac 用のオープンソースビデオエディターです。あなたとエージェントは、タイムライン内で一緒にビデオを生成および編集できます。
We built Palmier Pro from scratch with Swift.北極星は、AI をワークフローに統合するという当社の見解を備えた Premiere Pro です。
タイムライン エディター内で Seedance、Kling、Nano Banana Pro などの SOTA モデルを使用してビデオや画像を生成します。
MCP 経由でクロード/コーデックス/カーソルを接続するか、アプリ内エージェントを使用して同じプロジェクトで一緒に作業します。
アプリが開くと、HTTP 経由で http://127.0.0.1:19789/mcp にある MCP サーバーが公開されます。接続するには:
クロード mcp add --transport http Palmier-pro http://127.0.0.1:19789/mcp
コーデックス
codex mcp add palmier-pro --url http://127.0.0.1:19789/mcp
カーソル
最も簡単な方法は、アプリの [ヘルプ] -> [MCP 手順] -> [カーソルでインストール] に移動するか、これを ~/.cursor/mcp.json に追加して手動でインストールすることです。
{
"mcpサーバー": {
"パルミエプロ": {
"タイプ": "http",
"url": "http://127.0.0.1:19789/mcp"
}
}
}
クロードデスクトップ
ワンクリックで Claude Desktop にデスクトップ拡張機能をインストールできるアプリに mcpb がバンドルされています。 [ヘルプ] -> [MCP 手順] -> [Claude デスクトップにインストール] に移動します。
Is Palmier Pro fully open source?
ビデオ エディター (生成 AI 機能なし) は完全にオープン ソースです。 MCP

サーバーとエージェントのチャットもオープンソースです。クローズドソースであるのは生成 AI 処理だけです。
エディターは無料です。ログイン不要でダウンロードでき、CapCut や Adob​​e Premiere などのビデオエディターとして使用できます。 MCP サーバーを無料で使用し、クロード コード/デスクトップまたはカーソルを使用してタイムライン エディターと対話する実験を開始することもできます。
生成 AI 機能を使用するには、ログインとサブスクリプションが必要です。
どのようなプラットフォームをサポートしていますか?
Apple Silicon 上の macOS 26 (タホ) のみ。
Discord: Discord のコミュニティに参加してください。
Twitter / X: 最新情報やお知らせについては @Palmier_io をフォローしてください。
フィードバックとサポート: Github Issue を作成するか、founders@palmier.io までメールでお問い合わせください。
Copyright (C) 2026 株式会社パルミエ
Palmier Pro は GPLv3 に基づくオープンソースです。
AI 向けに構築された macOS ビデオ エディター
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
253
フォーク
レポートリポジトリ
リリース
53
v0.3.5
最新の
2026 年 6 月 20 日
+ 52 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS video editor built for AI. Contribute to palmier-io/palmier-pro development by creating an account on GitHub.

GitHub - palmier-io/palmier-pro: macOS video editor built for AI · GitHub
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
palmier-io
/
palmier-pro
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
633 Commits 633 Commits Sources/ PalmierPro Sources/ PalmierPro Tests/ PalmierProTests Tests/ PalmierProTests assets assets mcpb mcpb models models scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md FAQ.md FAQ.md LICENSE LICENSE Package.resolved Package.resolved Package.swift Package.swift README.md README.md appcast.xml appcast.xml View all files Repository files navigation
The video editor built for AI.
Requires macOS 26 (Tahoe) on Apple Silicon
Palmier Pro is an open source video editor for Mac. You and your agent can generate and edit videos together inside the timeline.
We built Palmier Pro from scratch with Swift. The north star is Premiere Pro, with our take on integrating AI into the workflow.
Generate videos and images with SOTA models like Seedance, Kling, Nano Banana Pro inside the timeline editor.
Connects your Claude/Codex/Cursor via MCP, or use the in-app agent to work on the same project together.
When the app is open, it exposes an MCP server at http://127.0.0.1:19789/mcp via HTTP. To connect:
claude mcp add --transport http palmier-pro http://127.0.0.1:19789/mcp
Codex
codex mcp add palmier-pro --url http://127.0.0.1:19789/mcp
Cursor
The easiest way is go inside the app Help -> MCP Instructions -> Install in Cursor , or install manually by adding this to ~/.cursor/mcp.json :
{
"mcpServers": {
"palmier-pro": {
"type": "http",
"url": "http://127.0.0.1:19789/mcp"
}
}
}
Claude Desktop
We bundle a mcpb with the app that allows a one click install Desktop Extension on Claude Desktop. Go to Help -> MCP Instructions -> Install in Claude Desktop
Is Palmier Pro fully open source?
The video editor (without the generative AI features) is fully open source. The MCP server and the agent chat are also open source. The only thing that is closed source is the generative AI processing.
The editor is free. You can download it with no login required, and use it as a video editor like CapCut or Adobe Premiere. You can also use the MCP server for free, and start experimenting using Claude Code/Desktop or Cursor to interact with your timeline editor.
Generative AI features require login and subscription.
What platforms does it support?
macOS 26 (Tahoe) on Apple Silicon only.
Discord: Join the community on Discord .
Twitter / X: Follow @Palmier_io for updates and announcements.
Feedback & Support: Create a Github Issue or email us at founders@palmier.io
Copyright (C) 2026 Palmier, Inc.
Palmier Pro is open source under GPLv3 .
macOS video editor built for AI
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
253
forks
Report repository
Releases
53
v0.3.5
Latest
Jun 20, 2026
+ 52 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

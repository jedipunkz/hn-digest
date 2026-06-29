---
source: "https://github.com/mjn298/ts7-lsp-plugin/tree/main"
hn_url: "https://news.ycombinator.com/item?id=48723931"
title: "Show HN: TypeScript7 LSP Claude Code Plugin"
article_title: "GitHub - mjn298/ts7-lsp-plugin: Claude Code plugin: TypeScript LSP backed by the native TypeScript 7 server (tsc --lsp) · GitHub"
author: "seedlessmike"
captured_at: "2026-06-29T19:33:08Z"
capture_tool: "hn-digest"
hn_id: 48723931
score: 1
comments: 0
posted_at: "2026-06-29T19:27:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TypeScript7 LSP Claude Code Plugin

- HN: [48723931](https://news.ycombinator.com/item?id=48723931)
- Source: [github.com](https://github.com/mjn298/ts7-lsp-plugin/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T19:27:48Z

## Translation

タイトル: HN を表示: TypeScript7 LSP クロード コード プラグイン
記事タイトル: GitHub - mjn298/ts7-lsp-plugin: クロード コード プラグイン: ネイティブ TypeScript 7 サーバーによってサポートされる TypeScript LSP (tsc --lsp) · GitHub
説明: クロード コード プラグイン: ネイティブ TypeScript 7 サーバーによってサポートされる TypeScript LSP (tsc --lsp) - mjn298/ts7-lsp-plugin
HN テキスト: タイトルを参照してください。他に興味深いものはありません。 TS7 は LSP 互換の言語サーバーを提供するため、このプラグインは typescript- language-server ではなくそれを使用します。ぜひお役立てください。

記事本文:
GitHub - mjn298/ts7-lsp-plugin: クロード コード プラグイン: ネイティブ TypeScript 7 サーバーによってサポートされる TypeScript LSP (tsc --lsp) · GitHub
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
mjn298
/
ts7-lsp-プラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のn

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .claude-plugin .claude-plugin plugins/ ts7-lsp plugins/ ts7-lsp LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TypeScript を提供する Claude Code マーケットプレイス
ネイティブ TypeScript 7 サーバー ( typescript@rc ) によってサポートされる言語サーバー。次のように起動されます。
tsc --lsp --stdio 。
これは、公式の typescript-lsp プラグイン (ノードベースの
従来の tsserver を囲む typescript- language-server ラッパー）、より高速なネイティブを備えた
実装に進みます。
非公式。これはコミュニティ プロジェクトであり、何の関係も持たず、承認もされません。
Microsoft または TypeScript チームによって保守されます。 「TypeScript」と tsc は、
マイクロソフト;このリポジトリは、typescript@rc LSP をクロード コードに接続するだけです。現状のまま提供されます。
.claude-plugin/marketplace.json # マーケットプレイス "ts7-lsp-marketplace" → プラグイン "ts7-lsp"
plugins/ts7-lsp/README.md # インストールと有効化の手順
クイックスタート
npm install -g typescript@rc # PATH 上の TS7 `tsc` — 任意のパッケージ マネージャーが機能します
/plugin Marketplace add mjn298/ts7-lsp-plugin # このマーケットプレイスを GitHub から直接追加します
/プラグインインストール ts7-lsp@ts7-lsp-marketplace
/plugin disable typescript-lsp@claude-plugins-official # 公式 TS LSP をすでにインストールしている場合のみ
/reload-プラグイン
mjn298/ts7-lsp-plugin は、この GitHub リポジトリの短縮形です。完全な URL を渡すこともできます
( https://github.com/mjn298/ts7-lsp-plugin )、またはハッキングしている場合はローカル チェックアウト パス。
詳細については、plugins/ts7-lsp/README.md を参照してください。
スイッチバックの指示。
0BSD — パブリックドメインに相当し、帰属は必要ありません。やりたいことは何でもしてください。
Claude Code プラグイン: ネイティブ TypeScript 7 サーバーをサポートする TypeScript LSP (tsc --lsp)
そこに

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

Claude Code plugin: TypeScript LSP backed by the native TypeScript 7 server (tsc --lsp) - mjn298/ts7-lsp-plugin

see title- nothing else interesting here. TS7 provides a LSP compatible language server, so this plugin uses it rather than typescript-language-server. Hope you find it useful.

GitHub - mjn298/ts7-lsp-plugin: Claude Code plugin: TypeScript LSP backed by the native TypeScript 7 server (tsc --lsp) · GitHub
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
mjn298
/
ts7-lsp-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .claude-plugin .claude-plugin plugins/ ts7-lsp plugins/ ts7-lsp LICENSE LICENSE README.md README.md View all files Repository files navigation
A Claude Code marketplace that provides a TypeScript
language server backed by the native TypeScript 7 server ( typescript@rc ), launched as
tsc --lsp --stdio .
It replaces the official typescript-lsp plugin (which uses the Node-based
typescript-language-server wrapper around the classic tsserver ) with the faster native
Go implementation.
Unofficial. This is a community project and is not affiliated with, endorsed by, or
maintained by Microsoft or the TypeScript team. "TypeScript" and tsc are products of
Microsoft; this repo just wires the typescript@rc LSP into Claude Code. Provided as-is.
.claude-plugin/marketplace.json # marketplace "ts7-lsp-marketplace" → plugin "ts7-lsp"
plugins/ts7-lsp/README.md # install + enable instructions
Quick start
npm install -g typescript@rc # a TS7 `tsc` on PATH — any package manager works
/plugin marketplace add mjn298/ts7-lsp-plugin # add this marketplace straight from GitHub
/plugin install ts7-lsp@ts7-lsp-marketplace
/plugin disable typescript-lsp@claude-plugins-official # ONLY if you'd already installed the official TS LSP
/reload-plugins
mjn298/ts7-lsp-plugin is shorthand for this GitHub repo. You can also pass the full URL
( https://github.com/mjn298/ts7-lsp-plugin ), or a local checkout path if you're hacking on it.
See plugins/ts7-lsp/README.md for full details and the
switch-back instructions.
0BSD — public-domain-equivalent, no attribution required. Do whatever you want.
Claude Code plugin: TypeScript LSP backed by the native TypeScript 7 server (tsc --lsp)
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

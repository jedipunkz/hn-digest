---
source: "https://github.com/tomreinert/claude-annotate"
hn_url: "https://news.ycombinator.com/item?id=48716465"
title: "Show HN: Claude Code plugin to draw feedback and send it back into the session"
article_title: "GitHub - tomreinert/claude-annotate: Claude Code plugin: draw on a live frontend in the Playwright browser, annotated screenshot flows back into the session. Local, no cloud. · GitHub"
author: "tom2948329494"
captured_at: "2026-06-29T09:32:57Z"
capture_tool: "hn-digest"
hn_id: 48716465
score: 1
comments: 0
posted_at: "2026-06-29T08:38:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code plugin to draw feedback and send it back into the session

- HN: [48716465](https://news.ycombinator.com/item?id=48716465)
- Source: [github.com](https://github.com/tomreinert/claude-annotate)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T08:38:01Z

## Translation

タイトル: Show HN: Claude Code プラグインを使用してフィードバックを取得し、セッションに送り返す
記事のタイトル: GitHub - tomreinert/claude-annotate: クロード コード プラグイン: Playwright ブラウザーのライブ フロントエンドに描画し、注釈付きのスクリーンショットがセッションに戻ります。ローカル、雲なし。 · GitHub
説明: Claude Code プラグイン: Playwright ブラウザーのライブ フロントエンドで描画し、注釈付きのスクリーンショットがセッションに戻ります。ローカル、雲なし。 - tomreinert/claude-annotate

記事本文:
GitHub - tomreinert/claude-annotate: クロード コード プラグイン: Playwright ブラウザーのライブ フロントエンドに描画し、注釈付きのスクリーンショットがセッションに戻ります。ローカル、雲なし。 · GitHub
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
トムライナート
/
クロード注釈を付ける
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
31 コミット 31 コミット .claude-plugin .claude-plugin annotate-plugin annotate-plugin .gitignore .gitignore GUIDE.md GUIDE.md README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
フィードバックをフロントエンドに描画し、クロードに直接送信します。
🧪 作業中です。構築しながらテスト中です。フィードバックは大歓迎です!
Playwright MCP を接続する必要があります。
クロード・mcp 劇作家 npx を追加 @playwright/mcp@latest
次に、プラグインをインストールします。
クロードプラグインマーケットプレイス追加 tomreinert/claude-annotate
クロードプラグインのインストール annotate@claude-annotate --scope user
使用する
チャンネルをオンにして Claude Code を起動します。
クロード --dangerously-load-development-channels プラグイン:annotate@claude-annotate
/annotate と入力します。クロードは、注釈ツールバーを備えた Playwright ブラウザを起動します。
フィードバックを入力し、 [送信] をクリックします。クロードは変化を起こします。
好きなだけ描いて送信し続けてください。終わったら「レビューをやめる」と言いましょう。
ツールバーのショートカットと FAQ (なぜ
--dangerously-load-development-channels フラグは安全であり、それをスキップする方法)。
これは、localhost:3000 で実行されているアプリを使用したローカル設定では非常にうまく機能します。
ただし、可動部分が多いため、さまざまな環境でどのように動作するか興味があります。
フィードバックや問題点をお待ちしています!
Claude Code プラグイン: Playwright ブラウザーのライブ フロントエンドで描画し、注釈付きのスクリーンショットがセッションに戻ります。ローカル、雲なし。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub、私

ノースカロライナ州
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code plugin: draw on a live frontend in the Playwright browser, annotated screenshot flows back into the session. Local, no cloud. - tomreinert/claude-annotate

GitHub - tomreinert/claude-annotate: Claude Code plugin: draw on a live frontend in the Playwright browser, annotated screenshot flows back into the session. Local, no cloud. · GitHub
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
tomreinert
/
claude-annotate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits .claude-plugin .claude-plugin annotate-plugin annotate-plugin .gitignore .gitignore GUIDE.md GUIDE.md README.md README.md View all files Repository files navigation
Draw your feedback onto your frontend and send it straight to Claude.
🧪 Work in progress, I'm still testing this as I build. Feedback welcome!
You need the Playwright MCP connected:
claude mcp add playwright npx @playwright/mcp@latest
Then install the plugin:
claude plugin marketplace add tomreinert/claude-annotate
claude plugin install annotate@claude-annotate --scope user
Use
Launch Claude Code with the channel turned on:
claude --dangerously-load-development-channels plugin:annotate@claude-annotate
Type /annotate . Claude launches a Playwright browser with the annotation toolbar.
Draw your feedback and hit Send . Claude makes the changes.
Keep drawing and sending as often as you like. Say "stop reviewing" when you are done.
See the full guide for the toolbar shortcuts and FAQ (including why the
--dangerously-load-development-channels flag is safe and how to skip it).
This works quite well in my local setup with an app running on localhost:3000.
Many moving parts though, so I'm curious how it works in different enviroments.
Looking forward to feedback and issues!
Claude Code plugin: draw on a live frontend in the Playwright browser, annotated screenshot flows back into the session. Local, no cloud.
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

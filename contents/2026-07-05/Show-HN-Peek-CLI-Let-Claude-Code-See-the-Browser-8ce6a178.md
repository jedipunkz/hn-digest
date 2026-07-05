---
source: "https://github.com/puffinsoft/peek-cli"
hn_url: "https://news.ycombinator.com/item?id=48799078"
title: "Show HN: Peek-CLI: Let Claude Code See the Browser"
article_title: "GitHub - puffinsoft/peek-cli: Let coding agents see your browser. · GitHub"
author: "ReactRocks"
captured_at: "2026-07-05T23:57:45Z"
capture_tool: "hn-digest"
hn_id: 48799078
score: 3
comments: 0
posted_at: "2026-07-05T23:38:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Peek-CLI: Let Claude Code See the Browser

- HN: [48799078](https://news.ycombinator.com/item?id=48799078)
- Source: [github.com](https://github.com/puffinsoft/peek-cli)
- Score: 3
- Comments: 0
- Posted: 2026-07-05T23:38:41Z

## Translation

タイトル: HN を表示: Peek-CLI: クロード コードをブラウザに表示させる
記事のタイトル: GitHub - puffinsoft/peek-cli: コーディング エージェントにブラウザを表示させます。 · GitHub
説明: コーディング エージェントにブラウザを表示させます。 GitHub でアカウントを作成して、puffinsoft/peek-cli の開発に貢献してください。

記事本文:
GitHub - puffinsoft/peek-cli: コーディング エージェントにブラウザを表示させます。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
パフィンソフト
/
ピーククリ
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
49 コミット 49 コミット .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows 資産 アセット スキル/ view-browser-tab スキル/ view-browser-tab src src .gitignore .gitignore ライセンス ライセンス README.md README.md plugin.json plugin.json Privacy.md Privacy.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Peak-cli を使用すると、エージェントはブラウザで開いているタブのスクリーンショットをキャプチャできます。
Claude Code、Codex、Copilot などで動作します...
デモ.mp4
これは、ブラウザ拡張機能を使用して、WebSocket 経由でスクリーンショットをストリーミングすることで機能します。
セキュリティのため、起動ごとにエージェントに 1 回接続する必要があります。
覗いたスタート
> サーバーが正常に起動しました。
ブラウザを接続します。
ピークリスト # 利用可能なタブを表示
> [ ' http://localhost:3000/ ' ]
http://localhost:3000 を覗いた # スクリーンショットをキャプチャ
> 画像の保存場所: /var/.../peek_cli/images/dd999ee0.jpg
インストール
Chrome 拡張機能をインストールします。
🎉 拡張機能は Chrome ウェブストアで承認されました!
npm i -g 覗いた
スキルをインストールする
/プラグインマーケットプレイス追加puffinsoft/peek-cli
/plugin install view-browser-tab@peek-cli
コーデックスの場合:
コーデックスプラグインマーケットプレイスにpuffinsoft/peek-cliを追加
コーデックス プラグインを追加 view-browser-tab@peek-cli
例
はい、100%です。エージェントはスクリーンショットを撮る以外に何もすることはできません。
これは、エージェントが WebSocket サーバー経由で拡張機能にスクリーンショット要求を送信するためです。
ブラウザにアクセスすることはなく、スクリプトを挿入したり、アクションを実行したりすることはできません。
安心して拡張機能のソース コードをご覧になることをお勧めします。
Peak-cli は、MIT ライセンスに基づいてライセンスされたオープン ソース ソフトウェアです。
コーディング エージェントにブラウザを表示させます。
読み込み中にエラーが発生しました。 PL

このページを簡単にリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
2
v0.2.0
最新の
2026 年 7 月 3 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Let coding agents see your browser. Contribute to puffinsoft/peek-cli development by creating an account on GitHub.

GitHub - puffinsoft/peek-cli: Let coding agents see your browser. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
puffinsoft
/
peek-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
49 Commits 49 Commits .claude-plugin .claude-plugin .codex-plugin .codex-plugin .github/ workflows .github/ workflows assets assets skills/ view-browser-tab skills/ view-browser-tab src src .gitignore .gitignore LICENSE LICENSE README.md README.md plugin.json plugin.json privacy.md privacy.md View all files Repository files navigation
peek-cli allows agents to capture a screenshot of any open tab in your browser.
Works with Claude Code, Codex, Copilot and many more...
demo.mp4
It works by using a browser extension to stream screenshots over WebSockets.
For security, you need to connect the agent once on every startup.
peeked start
> Successfully started server.
Connect your browser:
peeked list # view available tabs
> [ ' http://localhost:3000/ ' ]
peeked at http://localhost:3000 # capture screenshot
> Image saved to: /var/.../peek_cli/images/dd999ee0.jpg
Installation
Install the Chrome Extension .
🎉 the extension has been approved on the Chrome Web Store!
npm i -g peeked
Install the Skill
/plugin marketplace add puffinsoft/peek-cli
/plugin install view-browser-tab@peek-cli
For Codex:
codex plugin marketplace add puffinsoft/peek-cli
codex plugin add view-browser-tab@peek-cli
Examples
Yes, 100%. It is impossible for the agent to do anything but take a screenshot .
This is because the agent sends screenshot requests to the extension through a WebSocket server.
It never accesses the browser and cannot inject any scripts / perform any actions.
We invite you to take a look at the extension source code for added peace of mind.
peek-cli is open source software, licensed under the MIT license.
Let coding agents see your browser.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
2
v0.2.0
Latest
Jul 3, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

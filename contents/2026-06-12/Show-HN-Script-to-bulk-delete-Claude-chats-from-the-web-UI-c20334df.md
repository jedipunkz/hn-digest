---
source: "https://github.com/MatteoLeonesi/bulk-delete-claude-chat"
hn_url: "https://news.ycombinator.com/item?id=48505161"
title: "Show HN: Script to bulk delete Claude chats from the web UI"
article_title: "GitHub - MatteoLeonesi/bulk-delete-claude-chat: Delete all your claude.ai conversations at once · GitHub"
author: "ML0037"
captured_at: "2026-06-12T16:07:13Z"
capture_tool: "hn-digest"
hn_id: 48505161
score: 14
comments: 2
posted_at: "2026-06-12T15:08:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Script to bulk delete Claude chats from the web UI

- HN: [48505161](https://news.ycombinator.com/item?id=48505161)
- Source: [github.com](https://github.com/MatteoLeonesi/bulk-delete-claude-chat)
- Score: 14
- Comments: 2
- Posted: 2026-06-12T15:08:59Z

## Translation

タイトル: HN を表示: Web UI からクロード チャットを一括削除するスクリプト
記事のタイトル: GitHub - MatteoLeonesi/bulk-delete-claude-chat: すべての claude.ai 会話を一度に削除する · GitHub
説明: すべての claude.ai 会話を一度に削除します。 GitHub でアカウントを作成して、MatteoLeonesi/bulk-delete-claude-chat の開発に貢献してください。
HN テキスト: Chatgpt のようにすべてのチャットを一括で削除する方法が見つかりません。クロードの場合は、一番下までスクロールしてすべてを選択し、削除する必要があります。問題は、チャットがたくさんあるとそれが不可能になることです。このスクリプトを作成しました。それは単独で行います。誰かの役に立てば幸いです。 (会話は数分かけてゆっくりと UI から消えます。コンソールに「完了」と表示されるまでタブを開いたままにしてください。ページを更新すると削除プロセスが停止する可能性があります。)

記事本文:
GitHub - MatteoLeonesi/bulk-delete-claude-chat: すべての claude.ai 会話を一度に削除する · GitHub
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
マッテオ・レオネーシ
/
一括削除-クロード-チャット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
MatteoLeonesi/一括削除-クロード-チャット
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .gitignore .gitignore README.md README.md delete-all.js delete-all.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべての claude.ai 会話を一度に削除します。/recents の [すべて選択] ボタンは画面上にレンダリングされた行のみをカバーするため、これを作成しました。このスクリプトは代わりに内部 API を呼び出すため、すべてを取得します。
https://claude.ai/recents を開き、F12 → Console を押します。
delete-all.js の内容を貼り付け、Enter キーを押します。
ダイアログを確認します (組織ごとに 1 つ)
⚠️ しばらくお待ちください。会話は数分かけてゆっくりと UI から消えます。それは正常です
⚠️ コンソールに「完了」と表示されるまで、claude.ai タブを開いたままにしてください。ページを閉じる、更新する、またはページから移動すると、削除プロセスが停止する場合があります。
すべての claude.ai の会話を一度に削除します
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

Delete all your claude.ai conversations at once. Contribute to MatteoLeonesi/bulk-delete-claude-chat development by creating an account on GitHub.

I haven't found a way to delete all chats in bulk like you can on Chatgpt. With Claude, you have to scroll to the bottom, select everything, and delete. The problem is, if you have a lot of chats, it becomes impossible. I created this script. It does it alone. I hope it helps someone. (conversations disappear from the UI slowly, over several minutes, and remember to keep the tab open until the console shows "Finished", refreshing away from the page can stop the deletion process.)

GitHub - MatteoLeonesi/bulk-delete-claude-chat: Delete all your claude.ai conversations at once · GitHub
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
MatteoLeonesi
/
bulk-delete-claude-chat
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
MatteoLeonesi/bulk-delete-claude-chat
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .gitignore .gitignore README.md README.md delete-all.js delete-all.js View all files Repository files navigation
Delete all your claude.ai conversations at once, i created it because the "Select all" button on /recents only covers the rows rendered on screen. This script calls the internal API instead, so it gets everything.
Open https://claude.ai/recents and press F12 → Console
Paste the contents of delete-all.js , hit Enter
Confirm the dialog (one per organization)
⚠️ Be patient: conversations disappear from the UI slowly, over several minutes ; that's normal
⚠️ Keep the claude.ai tab open until the console shows Finished ; closing, refreshing, or navigating away from the page can stop the deletion process.
Delete all your claude.ai conversations at once
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

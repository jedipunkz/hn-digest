---
source: "https://github.com/prahladyeri/gubbi"
hn_url: "https://news.ycombinator.com/item?id=48793557"
title: "Show HN: Gubbi – Minimalist LLM Chatbot"
article_title: "GitHub - prahladyeri/gubbi: Minimalist LLM chatbot 🐤 · GitHub"
author: "pyeri"
captured_at: "2026-07-05T13:09:11Z"
capture_tool: "hn-digest"
hn_id: 48793557
score: 2
comments: 0
posted_at: "2026-07-05T12:08:20Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Gubbi – Minimalist LLM Chatbot

- HN: [48793557](https://news.ycombinator.com/item?id=48793557)
- Source: [github.com](https://github.com/prahladyeri/gubbi)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T12:08:20Z

## Translation

タイトル: 表示 HN: Gubbi – ミニマリスト LLM チャットボット
記事タイトル: GitHub - prahladyeri/gubbi: ミニマリスト LLM チャットボット 🐤 · GitHub
説明: ミニマリスト LLM チャットボット 🐤。 GitHub でアカウントを作成して、prahladyeri/gubbi の開発に貢献してください。

記事本文:
GitHub - prahladyeri/gubbi: ミニマリスト LLM チャットボット 🐤 · GitHub
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
プララディエリ
/
ガビ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く

アクション メニュー フォルダーとファイル
18 コミット 18 コミット Assets src/ gubbi src/ gubbi .gitignore .gitignore ライセンス ライセンス README.md README.md setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
pip インストール ガビ
構成
すべての設定は ~/.config/gubbi/settings.json に保存されます。チャットが機能するには、少なくとも 1 つのプロバイダーを追加する必要があります。手動で追加するか、プログラムが最初の起動時に入力を求めます。
{
「プロバイダー」: {
"github" : {
"スラッグ" : " github " 、
"url" : " https://models.github.ai/inference " ,
"default_model_id" : " openai/gpt-4.1-mini "
}
}
}
理論的には、openai API 仕様に準拠するどのプロバイダーでも動作するはずですが、これまでのところ github と他のいくつかのプロバイダーでしかテストしていません。他の人にもうまくいくかどうか、Issue Tracker を通じて知らせてください。
ガビ # ただチャットしてください
gubbi -a # プロバイダー/モデルを追加
コマンド
メッセージが # (ハッシュ) で始まる場合、LLM に送信されるチャット メッセージではなく、特別なコマンドとして解釈されます。
#help => コマンドのリスト
#exit => チャットを終了する
#use <プロバイダ> => プロバイダを切り替える
#model <model_name> => モデルを切り替える
#attach <パス> => ファイルを添付します
#clear => コンテキストをクリア
#save => 現在のチャットを保存
#load <filename> => 以前のチャットをロードします
#models <ファイル名> => すべてのモデルをリストする
#providers <filename> => すべてのプロバイダーをリストする
スクリーンショット
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
リリース 1.0.1
最新の
2026 年 7 月 5 日
貢献者
1
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Minimalist LLM chatbot 🐤. Contribute to prahladyeri/gubbi development by creating an account on GitHub.

GitHub - prahladyeri/gubbi: Minimalist LLM chatbot 🐤 · GitHub
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
prahladyeri
/
gubbi
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits assets assets src/ gubbi src/ gubbi .gitignore .gitignore LICENSE LICENSE README.md README.md setup.py setup.py View all files Repository files navigation
pip install gubbi
Configuration
All configuration is stored in ~/.config/gubbi/settings.json . At least one provider needs added in order for chat to work. Either add it manually or program will ask for inputs on the first start.
{
"providers" : {
"github" : {
"slug" : " github " ,
"url" : " https://models.github.ai/inference " ,
"default_model_id" : " openai/gpt-4.1-mini "
}
}
}
In theory, it should work with any provider who follows openai api spec though I've only tested with github and few others so far. Let me know through issue tracker if it works for others.
gubbi # just chat
gubbi -a # add provider/model
Commands
When a message starts with # (hash), it is interpreted as special command rather than a chat message sent to the LLM.
#help => List commands
#exit => Quit chat
#use <provider> => Switch provider
#model <model_name> => Switch model
#attach <path> => Attach a file
#clear => Clear context
#save => Save current chat
#load <filename> => Load an earlier chat
#models <filename> => List all models
#providers <filename> => List all providers
Screenshot
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Release 1.0.1
Latest
Jul 5, 2026
Contributors
1
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

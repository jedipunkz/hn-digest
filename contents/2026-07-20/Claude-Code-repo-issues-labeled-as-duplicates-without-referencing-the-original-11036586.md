---
source: "https://github.com/anthropics/claude-code/issues/79523"
hn_url: "https://news.ycombinator.com/item?id=48983088"
title: "Claude Code repo: issues labeled as duplicates without referencing the original"
article_title: "[BUG] Issues are labeled as duplicates without referencing the original issue · Issue #79523 · anthropics/claude-code · GitHub"
author: "marcindulak"
captured_at: "2026-07-20T19:25:43Z"
capture_tool: "hn-digest"
hn_id: 48983088
score: 1
comments: 0
posted_at: "2026-07-20T18:42:48Z"
tags:
  - hacker-news
  - translated
---

# Claude Code repo: issues labeled as duplicates without referencing the original

- HN: [48983088](https://news.ycombinator.com/item?id=48983088)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/79523)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T18:42:48Z

## Translation

タイトル: Claude Code リポジトリ: オリジナルを参照せずに重複としてラベル付けされた問題
記事のタイトル: [バグ] 元の問題を参照せずに問題が重複としてラベル付けされる · 問題 #79523 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存の問題を検索しましたが、まだ報告されていません これは 1 つのバグ レポートです (バグごとに個別のレポートを提出してください) 最新バージョンの Claude Code を使用しています 何が問題ですか?問題は実験室です...

記事本文:
[バグ] 元の問題を参照せずに問題が重複としてラベル付けされる · 問題 #79523 · anthropics/claude-code · GitHub
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
コードの品質 マージ時に品質を強制する
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
アントロ

写真
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
[バグ] 元の問題を参照せずに問題が重複としてラベル付けされる #79523
リンクをコピー 新しい問題 リンクをコピー 開く 開く [バグ] 元の問題を参照せずに問題が重複としてラベル付けされる #79523 リンクをコピー ラベルのバグ 何かが機能していない 何かが機能していない 説明
発行本体アクションのプリフライト チェックリスト
既存の問題を検索しましたが、まだ報告されていません
これは 1 つのバグ レポートです (バグごとに別々のレポートを提出してください)
最新バージョンのクロードコードを使用しています
問題は、元の問題を参照せずに、github-actions ボットによって重複としてラベル付けされます。
この問題が重複している元の問題を説明するコメントのない、重複としてラベル付けされている問題 #79240 の例を参照してください。
上記の問題 #78415 と比較してください。これには、元の問題をリストしたコメントが含まれています。
問題が重複としてラベル付けされている場合は、問題が重複している元の問題に関する情報を含むコメントを投稿する必要があります。
問題を送信します (例: git commit #79240 で強制された「Co-Authored-By: Claude」トレーラーを削除/無効化を許可する)。
クロードの問題を見つけて、予期されるコメントなしで git commit #79240 の強制的な「共同作成者: クロード」トレーラーを削除/無効にする問題に対応する重複排除ワークフローを実行します。
クロードの問題 [バグ] クロード デスクトップの更新後にセッションの自動名前付けが停止する — セッションは予期されるコメントとともに「無題のセッション」 (デスクトップ 1.22209.0、macOS) #78415 のままになります。
ワークフローの実行出力には、前の問題 (ポイント 2.) にコメントと緯度が欠落している理由が示されていないようです。

後の問題 (ポイント 3.) には予想通りのコメントがあります。クロード コード セッションからの詳細情報は、ランナーに保持されるログ ファイルに存在する可能性があり、別のワークフローとのやり取りである可能性があります。
はい、これは以前のバージョンでは機能していました
この問題は #19267 の問題を悪化させます。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Preflight Checklist I have searched existing issues and this hasn't been reported yet This is a single bug report (please file separate reports for different bugs) I am using the latest version of Claude Code What's Wrong? Issues are lab...

[BUG] Issues are labeled as duplicates without referencing the original issue · Issue #79523 · anthropics/claude-code · GitHub
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
Code Quality Enforce quality at merge
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
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
[BUG] Issues are labeled as duplicates without referencing the original issue #79523
Copy link New issue Copy link Open Open [BUG] Issues are labeled as duplicates without referencing the original issue #79523 Copy link Labels bug Something isn't working Something isn't working Description
Issue body actions Preflight Checklist
I have searched existing issues and this hasn't been reported yet
This is a single bug report (please file separate reports for different bugs)
I am using the latest version of Claude Code
Issues are labeled by github-actions bot as duplicates without referencing the original issue.
See an example of the issue #79240 that is labeled as duplicate without a comment explaining what is the original issue this issue os the duplicate of.
Compare the above the issue #78415 , which contains the comment listing the original issues.
If an issue is labeled as duplicate, a comment should be posted that includes information which are the original issues the issue is a duplicate of.
Submit an issue (example Remove/allow disabling the forced "Co-Authored-By: Claude" trailer on git commits #79240 ).
Find the Claude Issue Dedupe worfklow run corresponding to the issue Remove/allow disabling the forced "Co-Authored-By: Claude" trailer on git commits #79240 without the expected comment.
Find the Claude Issue Dedupe worfklow run corresponding to the issue [BUG] Session auto-naming stopped after Claude Desktop update — sessions stay "Untitled session" (Desktop 1.22209.0, macOS) #78415 with the expected comment.
The workflow run output does not seem to contain an indication why the former issue (point 2.) is lacking the comment and the latter issue (point 3.) has the expected comment. More information from the Claude Code sessions may be present in the log files retained on the runner, possibly an interaction with another workflow.
Yes, this worked in a previous version
This problem aggravates the #19267 issue.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

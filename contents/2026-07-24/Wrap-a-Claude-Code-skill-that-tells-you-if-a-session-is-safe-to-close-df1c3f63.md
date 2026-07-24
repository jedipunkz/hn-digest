---
source: "https://github.com/ravsau/wrap"
hn_url: "https://news.ycombinator.com/item?id=49042416"
title: "Wrap – a Claude Code skill that tells you if a session is safe to close"
article_title: "GitHub - ravsau/wrap: A Claude Code skill that checks the current session for loose threads and tells you whether it's safe to close: STOP or CONTINUE. · GitHub"
author: "cloudyeti"
captured_at: "2026-07-24T22:55:05Z"
capture_tool: "hn-digest"
hn_id: 49042416
score: 1
comments: 0
posted_at: "2026-07-24T22:31:05Z"
tags:
  - hacker-news
  - translated
---

# Wrap – a Claude Code skill that tells you if a session is safe to close

- HN: [49042416](https://news.ycombinator.com/item?id=49042416)
- Source: [github.com](https://github.com/ravsau/wrap)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T22:31:05Z

## Translation

タイトル: Wrap – セッションを閉じても安全かどうかを知らせるクロード コード スキル
記事のタイトル: GitHub - ravsau/wrap: 現在のセッションでスレッドの緩みをチェックし、閉じても安全かどうか (STOP または CONTINUE) を通知するクロード コード スキル。 · GitHub
説明: 現在のセッションでスレッドの緩みをチェックし、閉じても安全かどうか (STOP または CONTINUE) を通知するクロード コード スキル。 - ラヴサウ/ラップ

記事本文:
GitHub - ravsau/wrap: 現在のセッションでスレッドの緩みをチェックし、閉じても安全かどうか (STOP または CONTINUE) を通知するクロード コード スキル。 · GitHub
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
ラヴサウ
/
包む
公共
通知
あなたはむ

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .claude-plugin .claude-plugin plugins/wrap plugins/wrap CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード セッションを閉じる前に、次の 1 つの質問をしてください。停止しても安全か、それとも
まだ何かが終わっていませんか？
/wrap は、現在の会話と回答のみをレビューします。
STOP — 安全に閉めることができ、緩んだ糸がリストに表示されているため、糸を紛失することはありません。
続行 — ここにある何かは、今すぐにでも終わらせる価値があります
バックログや他のセッションをスキャンしたり、追加の作業を生み出したりすることはありません。
git clone https://github.com/ravsau/wrap.git
CDラップ
./install.sh
次に、任意のセッションで /wrap を実行します。コマンドが実行されない場合は、新しいセッションを開始します。
すぐに現れます。
プラグイン マーケットプレイスからインストールする
/プラグイン マーケットプレイスに ravsau/wrap を追加
/プラグインのインストール Wrap@saurav-claude-tools
/reload-プラグイン
次に、 /wrap:wrap を実行します (プラグインは名前空間化されています。上記のインストーラーは
裸の /wrap が得られます)。
ハンドオフの判定: 作業を続行する必要があるが、新しいセッションで行う場合 - 出力
単に「CONTINUE」ではなく、すぐに貼り付けられるハンドオフの概要。
現在のセッションでスレッドの緩みをチェックし、閉じても安全かどうか (STOP または CONTINUE) を通知するクロード コード スキル。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Claude Code skill that checks the current session for loose threads and tells you whether it's safe to close: STOP or CONTINUE. - ravsau/wrap

GitHub - ravsau/wrap: A Claude Code skill that checks the current session for loose threads and tells you whether it's safe to close: STOP or CONTINUE. · GitHub
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
ravsau
/
wrap
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .claude-plugin .claude-plugin plugins/ wrap plugins/ wrap CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
Ask any Claude Code session one question before you close it: safe to stop, or
is something still unfinished?
/wrap reviews only the current conversation and answers:
STOP — safe to close, with any loose threads listed so you don't lose them
CONTINUE — something here is still worth finishing now
It does not scan your backlog, other sessions, or invent more work.
git clone https://github.com/ravsau/wrap.git
cd wrap
./install.sh
Then run /wrap in any session. Start a new session if the command does not
appear immediately.
Install from the plugin marketplace
/plugin marketplace add ravsau/wrap
/plugin install wrap@saurav-claude-tools
/reload-plugins
Then run /wrap:wrap (plugins are namespaced; the installer above
gives you the bare /wrap ).
HANDOFF verdict: when work should continue, but in a fresh session — output
a paste-ready handoff summary instead of just CONTINUE.
A Claude Code skill that checks the current session for loose threads and tells you whether it's safe to close: STOP or CONTINUE.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

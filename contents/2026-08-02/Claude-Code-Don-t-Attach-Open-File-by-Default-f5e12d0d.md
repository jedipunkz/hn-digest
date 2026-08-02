---
source: "https://github.com/anthropics/claude-code/issues/63925"
hn_url: "https://news.ycombinator.com/item?id=49142870"
title: "Claude Code – Don't Attach Open File by Default"
article_title: "[Feature Request] Add settings to disable auto-attaching active editor file and selection by default · Issue #63925 · anthropics/claude-code · GitHub"
author: "danmaz74"
captured_at: "2026-08-02T10:25:58Z"
capture_tool: "hn-digest"
hn_id: 49142870
score: 2
comments: 0
posted_at: "2026-08-02T10:00:58Z"
tags:
  - hacker-news
  - translated
---

# Claude Code – Don't Attach Open File by Default

- HN: [49142870](https://news.ycombinator.com/item?id=49142870)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/63925)
- Score: 2
- Comments: 0
- Posted: 2026-08-02T10:00:58Z

## Translation

タイトル: Claude Code – デフォルトでは開いているファイルを添付しない
記事のタイトル: [機能リクエスト] アクティブ エディター ファイルと選択の自動添付をデフォルトで無効にする設定を追加 · 問題 #63925 · anthropics/claude-code · GitHub
説明: 概要: VS Code 拡張機能は、現在開いているエディター ファイル (および現在のテキスト選択) をすべてのプロンプトのコンテキストとして自動的に添付します。現在、これをデフォルトでオフにする方法はありません。メッセージごとにのみ抑制できます (...

記事本文:
[機能リクエスト] アクティブなエディター ファイルと選択の自動添付をデフォルトで無効にする設定を追加 · 問題 #63925 · anthropics/claude-code · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
ああ、ああ！
読み込み中にエラーが発生しました。リロードしてください

このページ。
人類学
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
[機能リクエスト] アクティブなエディター ファイルと選択の自動添付をデフォルトで無効にする設定を追加 #63925
リンクをコピー 新しい問題 リンクをコピー 開く 開く [機能リクエスト] アクティブなエディター ファイルと選択の自動添付をデフォルトで無効にする設定を追加 #63925 リンクをコピー ラベル エリア:ide の機能強化 新機能またはリクエスト 新機能またはリクエスト プラットフォーム:vscode 問題は特に VS Code で発生します 問題は特に VS Code で発生します 説明
本文のアクションを発行する 概要: VS Code 拡張機能は、現在開いているエディター ファイル (および現在のテキスト選択) をすべてのプロンプトのコンテキストとして自動的に添付します。現時点では、これをデフォルトでオフにする方法はありません。メッセージごとにのみ抑制できます (チップの [×] をクリックするか、目のスラッシュ選択トグルをクリックする)。これにより、ターンごとにリセットされます。
リクエスト: デフォルトの自動コンテキストをオフにする永続的な設定。理想的には個別の切り替えを使用します。
アクティブなエディター ファイルの自動添付を無効にする
現在の選択範囲の自動共有を無効にする
拡張機能設定パネルおよび/または settings.json を介して構成できます (例: claudeCode.autoAttachActiveFile: false)。
理由: 開いているファイルを自動的に挿入するのではなく、@-メンションを使用してコンテキストを明示的に制御したいことがよくあります。理由としては、関連のないファイルや開いているファイルがプロンプトに漏れることを避けること、コンテキストを無駄にしないこと、たまたま表示しているファイル以外の内容についてプロンプトを表示することなどが挙げられます。既存の Read(...) 拒否ルールの回避策はパスごとであり、必要なときにクロードがこれらのファイルを読み取れないようにするため、代替手段ではありません。
現在の回避策: チップ上のメッセージごとの × / アイ スラッシュ トグル、または特定の機密パスに対する読み取り拒否ルール。どちらもGLを提供しません

楕円形のデフォルトはオフです。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Summary: The VS Code extension automatically attaches the currently open editor file (and current text selection) as context on every prompt. There is currently no way to turn this off by default — it can only be suppressed per-message (...

[Feature Request] Add settings to disable auto-attaching active editor file and selection by default · Issue #63925 · anthropics/claude-code · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
[Feature Request] Add settings to disable auto-attaching active editor file and selection by default #63925
Copy link New issue Copy link Open Open [Feature Request] Add settings to disable auto-attaching active editor file and selection by default #63925 Copy link Labels area:ide enhancement New feature or request New feature or request platform:vscode Issue specifically occurs in VS Code Issue specifically occurs in VS Code Description
Issue body actions Summary: The VS Code extension automatically attaches the currently open editor file (and current text selection) as context on every prompt. There is currently no way to turn this off by default — it can only be suppressed per-message (clicking × on the chip, or the eye-slash selection toggle), which resets each turn.
Request: A persistent setting to default auto-context to off, ideally with separate toggles:
disable auto-attach of active editor file
disable auto-share of current selection
Configurable via the extension settings panel and/or settings.json (e.g. claudeCode.autoAttachActiveFile: false).
Why: I often want to control context explicitly via @-mentions rather than have the open file injected automatically. Reasons include: avoiding leaking unrelated/open files into the prompt, keeping context lean, and prompting about something other than the file I happen to be viewing. The existing Read(...) deny-rule workaround is per-path and also blocks Claude from reading those files when I do want it to — so it's not a substitute.
Current workaround: Per-message × on the chip / eye-slash toggle, or Read deny rules for specific sensitive paths. Neither provides a global default-off.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

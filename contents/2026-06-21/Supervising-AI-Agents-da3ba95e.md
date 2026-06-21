---
source: "https://github.com/usefulsoftworks/ai-agent-control-checklist"
hn_url: "https://news.ycombinator.com/item?id=48620058"
title: "Supervising AI Agents"
article_title: "GitHub - usefulsoftworks/ai-agent-control-checklist: A practical checklist for supervising AI coding agents across branches, worktrees, reviews, approvals, and human intervention points. · GitHub"
author: "usefulsoftworks"
captured_at: "2026-06-21T16:32:13Z"
capture_tool: "hn-digest"
hn_id: 48620058
score: 2
comments: 0
posted_at: "2026-06-21T16:05:03Z"
tags:
  - hacker-news
  - translated
---

# Supervising AI Agents

- HN: [48620058](https://news.ycombinator.com/item?id=48620058)
- Source: [github.com](https://github.com/usefulsoftworks/ai-agent-control-checklist)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T16:05:03Z

## Translation

タイトル: AI エージェントの監督
記事のタイトル: GitHub - Usefulsoftworks/ai-agent-control-checklist: ブランチ、ワークツリー、レビュー、承認、人間の介入ポイント全体で AI コーディング エージェントを監督するための実用的なチェックリスト。 · GitHub
説明: ブランチ、ワークツリー、レビュー、承認、人間の介入ポイント全体で AI コーディング エージェントを監督するための実用的なチェックリスト。 - 便利なソフトワークス/ai-agent-control-checklist

記事本文:
GitHub - Usefulsoftworks/ai-agent-control-checklist: ブランチ、ワークツリー、レビュー、承認、人間の介入ポイント全体で AI コーディング エージェントを監督するための実用的なチェックリスト。 · GitHub
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
便利なソフトウェアワークス
/
AIエージェント制御チェックリスト

公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
便利なソフトワークス/ai-agent-control-checklist
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE .gitignore .gitignore COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md checklist.md checklist.md Failure-modes.md Failure-modes.mdサンプル-エージェント-レビュー-ログ.md サンプル-エージェント-レビュー-ログ.md worktrees-vs-control.md worktrees-vs-control.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ブランチ、ワークツリー、レビュー、承認、人間の介入ポイント全体にわたって AI コーディング エージェントを監督するための実用的なチェックリスト。
AI コーディング エージェント (Claude Code、Cursor、Codex、Aider、OpenCode、および同様のツール) は、実際の作業を実行できるほど強力になりつつあります。彼らは計画、作成、テスト、反復を行うことができます。以前は数分かかっていたセッションが何時間も実行されるようになりました。エージェントはより多くのファイルにアクセスし、より多くのサーフェスで動作し、並行して実行されることが増えています。
それが起こると、ボトルネックが移動します。
コード生成はもはや難しい部分ではありません。監修は。
重要な質問は次のとおりです。
現在、どのエージェントが何に取り組んでいますか?
誰がいつ、何をレビューしましたか?
何かがマージされる前の承認状態は何ですか?
人間が介入する必要があるのはいつですか?
最後のセッションで何が起こりましたか?
ほとんどの開発者は、これらの質問に対して適切な答えを持っていません。このリポジトリはそれを支援する試みです。
ワークツリーはコードを分離します。これらは、状態、レビュー、承認、所有権、または人間の介入を完全に解決するものではありません。
Git ワークツリーとブランチは適切な基盤です。しかし、それらは問題の 1 つの層、つまりコード層にのみ対処しています。それ以上のすべて (誰がレビューするか、何を承認するか、いつ一時停止するか、どのように監査するか) は依然として有効です。

意図的なプロセスが必要です。
詳細については、worktrees-vs-control.md を参照してください。
エージェント セッションの最小限のチェックリスト:
スコープはエージェントの開始前に定義されます
エージェントは分離されたブランチまたはワークツリーで動作しています
エージェントが触れる可能性のあるファイルは事前にわかっています
エージェントを一時停止または中断する方法を知っている
何かがステージングまたはマージされる前に差分がレビューされます
テストに合格すると、その内容が理解できる
エージェントによってシークレット、資格情報、または構成値が書き込まれていない
変更はメインに反映される前に人間によって承認されます。
完全版については、 checklist.md を参照してください。
ファイル
対象となる内容
チェックリスト.md
フェーズごとにまとめられた完全なセッションのチェックリスト
障害モード.md
エージェント支援開発が失敗する一般的な原因
worktrees-vs-control.md
ワークツリーで何が解決され、何が解決されないのか
サンプルエージェントレビューログ.md
セッション間でエージェントの作業を追跡するための形式の例
貢献.md
貢献方法
なぜこれが存在するのか
エージェントのツールは急速に進歩しています。ほとんどの議論は、エージェントが何ができるかに焦点を当てています。実際に大規模に実行すると何が壊れるか、つまり複数のセッション、複数のファイル、実際の運用コードベースにはあまり注意が払われません。
このチェックリストはフレームワークではありません。これは、開発者がすでに使用しているどのようなワークフローにも適応できる実用的なコントロールのコレクションです。
このチェックリストは Useful Softworks によって管理されています。また、AI コーディング エージェントを使用するビルダー向けのローカルファースト制御レイヤーである AgentLeash も検討しています。プライベート ベータ アプリケーションは https://agentleash.dev/ で公開されています。
AgentLeash は、Chain Bridge Labs LLC が運営するソフトウェア ブランドである Useful Softworks の製品です。
COTRIBUTING.md を参照してください。新しい障害モード、改善されたチェックリスト項目、ツール固有のメモについては、プル リクエストを歓迎します。
ブランチ、ワークツリー、レビュー、承認全体にわたる AI コーディング エージェントを監督するための実用的なチェックリスト

、人間の介入ポイント。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A practical checklist for supervising AI coding agents across branches, worktrees, reviews, approvals, and human intervention points. - usefulsoftworks/ai-agent-control-checklist

GitHub - usefulsoftworks/ai-agent-control-checklist: A practical checklist for supervising AI coding agents across branches, worktrees, reviews, approvals, and human intervention points. · GitHub
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
usefulsoftworks
/
ai-agent-control-checklist
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
usefulsoftworks/ai-agent-control-checklist
master Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md checklist.md checklist.md failure-modes.md failure-modes.md sample-agent-review-log.md sample-agent-review-log.md worktrees-vs-control.md worktrees-vs-control.md View all files Repository files navigation
A practical checklist for supervising AI coding agents across branches, worktrees, reviews, approvals, and human intervention points.
AI coding agents -- Claude Code, Cursor, Codex, Aider, OpenCode, and similar tools -- are becoming powerful enough to do real work. They can plan, write, test, and iterate. Sessions that used to take minutes now run for hours. Agents touch more files, operate across more surfaces, and increasingly run in parallel.
When that happens, the bottleneck shifts.
Code generation is no longer the hard part. Supervision is.
The questions that matter become:
Which agent is working on what right now?
Who has reviewed what, and when?
What is the approval state before anything merges?
When does a human need to step in?
What happened in the last session?
Most developers do not have good answers to these questions. This repo is an attempt to help with that.
Worktrees isolate code. They do not fully solve state, review, approvals, ownership, or human intervention.
Git worktrees and branches are the right foundation. But they only address one layer of the problem -- the code layer. Everything above that (who reviews, what gets approved, when to pause, how to audit) still requires deliberate process.
See worktrees-vs-control.md for a full breakdown.
A minimal checklist for any agent session:
Scope is defined before the agent starts
Agent is working in an isolated branch or worktree
Files the agent may touch are known in advance
You know how to pause or interrupt the agent
Diffs are reviewed before anything is staged or merged
Tests pass and you understand what they cover
No secrets, credentials, or config values were written by the agent
Changes are approved by a human before they reach main
For the full version, see checklist.md .
File
What it covers
checklist.md
Full session checklist organized by phase
failure-modes.md
Common ways agent-assisted development goes wrong
worktrees-vs-control.md
What worktrees solve and what they do not
sample-agent-review-log.md
Example format for tracking agent work across sessions
CONTRIBUTING.md
How to contribute
Why this exists
Agent tooling is moving fast. Most of the discourse focuses on what agents can do. Less attention goes to what breaks when you actually run them at scale -- multiple sessions, multiple files, real production codebases.
This checklist is not a framework. It is a collection of practical controls that developers can adapt to whatever workflow they already use.
This checklist is maintained by Useful Softworks. We are also exploring AgentLeash, a local-first control layer for builders using AI coding agents. Private beta applications are open at https://agentleash.dev/
AgentLeash is a product of Useful Softworks, a software brand operated by Chain Bridge Labs LLC.
See CONTRIBUTING.md . Pull requests are welcome for new failure modes, improved checklist items, and tool-specific notes.
A practical checklist for supervising AI coding agents across branches, worktrees, reviews, approvals, and human intervention points.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

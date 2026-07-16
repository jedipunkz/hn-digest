---
source: "https://www.flanner.io/"
hn_url: "https://news.ycombinator.com/item?id=48932248"
title: "Flanner.io - Stop drowning in AI .md files."
article_title: "Flanner | Version control for your AI's planning docs"
author: "jaysonm"
captured_at: "2026-07-16T09:57:44Z"
capture_tool: "hn-digest"
hn_id: 48932248
score: 1
comments: 0
posted_at: "2026-07-16T09:30:53Z"
tags:
  - hacker-news
  - translated
---

# Flanner.io - Stop drowning in AI .md files.

- HN: [48932248](https://news.ycombinator.com/item?id=48932248)
- Source: [www.flanner.io](https://www.flanner.io/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T09:30:53Z

## Translation

タイトル: Flanner.io - AI .md ファイルに溺れるのをやめます。
記事のタイトル: フランナー | AI の計画ドキュメントのバージョン管理
説明: AI エージェントが生成する計画 .md ファイルをバージョン管理して保護する、無料のオープンソース CLI および MCP サーバー。クロード コードおよびコーデックスで動作します。

記事本文:
フラナー | AI の計画のバージョン管理 docs flanner 機能 仕組み ユースケース 価格 GitHub の Docs Star MIT オープン ソース · PyPI の v0.7.1 AI の .md ファイルに溺れるのはやめましょう。
Flanner は、AI エージェントが生成する計画ドキュメント用の無料のオープンソース CLI および MCP サーバーです。これらのファイルは、ディスク上のリポジトリ全体とトラッカー内の課題全体の 2 つの方向に同時に分散されます。flanner は両方の制御ポイントとなり、すべてのリビジョンをバージョン管理し、各プランをそれが属する課題にリンクします。
GitHub で始める クラウド待機リストに参加する ▹ MIT ライセンス取得 ▹ 100% ローカル、アカウントなし ▹ MCP Claude Code + Codex Architecture.md 同期履歴 v 1 v 2 v 3 --- mcp_plan_file: true version: 3 updated_by: claude --- # システム アーキテクチャ - ガードフック フローを文書化 + - Codex 用に AGENTS.md をワイヤ Git-protected // 機能 すべてAI 計画ファイルを管理する必要がある
Claude Code、Codex、および任意の MCP クライアントを使用した AI 支援開発用に構築
ガード フックと CLAUDE.md/AGENTS.md およびスキルにより、すべてのプランがフランナーを介して自動ルーティングされるため、エージェントにドキュメントの保存を促す必要はありません。
モデル コンテキスト プロトコルを通じて、Claude Code、Codex (AGENTS.md)、およびその他の AI アシスタントとシームレスに統合
.gitignore 管理を介して計画ファイルが git にコミットされるのを自動的に防止します
すべての計画ファイルには、プロジェクト情報、バージョン管理、所有権追跡を含むメタデータが含まれています
セマンティック バージョニングで変更を追跡し、すべての計画反復の完全な履歴を維持します
マークダウン プレビューを使用してプラン ファイルを表示、編集、管理するための美しいブラウザベースの UI
プロジェクト管理、同期、ステータスチェックのための強力なコマンドラインインターフェイス
すべてはマシン上のローカル SQLite データベースに存在し、アカウントは必要ありません。クラウド ワークスペースはロードマップにあります

。
プロジェクトごとにプラン ディレクトリをカスタマイズし、Git 統合設定を構成する
初期から完全にバージョン化されたプラン履歴までの 5 つのステップ
「flanner init」を実行してデータベースをセットアップし、.plans ディレクトリを作成し、クロード コードに自動登録します。
MCP を通じて計画ファイルを作成するようにクロードに依頼します。
または
ファイルを手動で作成してローカル プロジェクトに添付する
更新ごとに、完全な履歴とメタデータが保存された新しいバージョンが作成されます
Web UI または CLI を使用して、計画の参照、履歴の表示、プロジェクトの管理を行います。
ファイルをリモート組織アカウントにプッシュし、JIRA と同期して .md ファイルをタスクに関連付けます
開発者の実際の作業方法に合わせて構築
ソロでもチームでも、flanner はワークフローに適応します
Claude の AI 支援により、アーキテクチャの決定と計画を整理しておくことができます。
計画ファイルを共有し、誰が変更を加えたかを追跡し、チーム全体で一貫した文書を維持します。
Claude Code で計画文書を管理しながら、バージョン管理と監査証跡を維持しましょう
RFC、技術仕様、アーキテクチャの決定を自動バージョン管理で保存
git を無視したローカル計画ファイルを使用してプロジェクトのロードマップと機能計画を管理する
Web ダッシュボードへのアクセスと明確な所有権追跡による非同期コラボレーション
flanner を問題トラッカーに接続します。リニア同期はライブです。さらに多くのことが進行中です。
Linear API を介した双方向同期。計画は、リンクされた問題がワークフロー内を移動するにつれて自動的に更新されます。
トレーサビリティのために、計画を Jira 課題にリンクします。現在はリンクベースで、ライブステータスの同期はありません。
GitHub Issues、Notion、その他の統合が進行中です。
現在、ローカル ツールは無料でオープンソースです。クラウド ワークスペース (共有ストレージ、チーム コラボレーション、Jira、Notion、チャットへのリンク) がロードマップにあります。順番待ちリストに参加して、最初に話を聞いてください。
ウェイティングリストに参加する メールアドレスを残して、

発送日に必ず 1 通のメッセージが届きます。スパムはありません。
Flanner をインストールし、Claude Code または Codex に接続して、自信を持って計画ファイルの管理を始めてください。無料でオープンソース。

## Original Extract

Free, open-source CLI and MCP server that versions and protects the planning .md files your AI agents generate. Works with Claude Code and Codex.

Flanner | Version control for your AI's planning docs flanner Features How it Works Use Cases Pricing Docs Star on GitHub MIT open source · v0.7.1 on PyPI Stop drowning in AI .md files.
Flanner is a free, open-source CLI and MCP server for the planning docs your AI agents generate. Those files scatter in two directions at once, across your repos on disk and across the issues in your tracker, and flanner is the control point for both, versioning every revision and linking each plan to the issue it belongs to.
Get started on GitHub Join the cloud waitlist ▹ MIT licensed ▹ 100% local, no account ▹ MCP Claude Code + Codex architecture.md Synced history v 1 v 2 v 3 --- mcp_plan_file: true version: 3 updated_by: claude --- # System Architecture - Document the guard-hook flow + - Wire AGENTS.md for Codex Git-protected // features Everything you need to manage AI plan files
Built for AI-assisted development with Claude Code, Codex, and any MCP client
A guard hook plus CLAUDE.md/AGENTS.md and a skill auto-route every plan through flanner, so you never remind the agent to save its docs.
Seamlessly integrates with Claude Code, Codex (AGENTS.md), and other AI assistants through the Model Context Protocol
Automatically prevents plan files from being committed to git via .gitignore management
Every plan file includes metadata with project info, versioning, and ownership tracking
Tracks changes with semantic versioning and maintains complete history of all plan iterations
Beautiful browser-based UI for viewing, editing, and managing plan files with markdown preview
Powerful command-line interface for project management, syncing, and status checking
Everything lives in a local SQLite database on your machine, no account required. Cloud workspaces are on the roadmap.
Customize plan directory per project and configure git integration preferences
From init to a fully versioned plan history, in five steps
Run 'flanner init' to set up the database, create .plans directory, and auto-register with Claude Code
Ask Claude to create plan files through MCP
or
Manually create and attach files to your local projects
Every update creates a new version with complete history and metadata preserved
Use the web UI or CLI to browse plans, view history, and manage your projects
Push files to your remote organisation account, and sync with JIRA to associate .md files to tasks
Built for how developers actually work
Whether you're solo or on a team, flanner adapts to your workflow
Keep your architecture decisions and planning organized with AI assistance from Claude
Share plan files, track who made changes, and maintain consistent documentation across the team
Let Claude Code manage your planning documents while maintaining version control and audit trails
Store RFCs, technical specs, and architecture decisions with automatic versioning
Manage project roadmaps and feature planning with git-ignored local plan files
Asynchronous collaboration with web dashboard access and clear ownership tracking
Connect flanner to your issue tracker. Linear syncs live; more are on the way.
Two-way sync over the Linear API. Plans update automatically as their linked issues move through your workflow.
Link any plan to a Jira issue for traceability. Link-based today, without live status sync.
GitHub Issues, Notion, and more integrations are on the way.
The local tool is free and open source today. Cloud workspaces (shared storage, team collaboration, and links to Jira, Notion, and chat) are on the roadmap. Join the waitlist to hear first.
Join waitlist Leave your email and you'll get exactly one message the day it ships. No spam.
Install Flanner, connect it to Claude Code or Codex, and start managing your plan files with confidence. Free and open source.

---
source: "https://github.com/keyuchen21/agentic-engineering-handbook"
hn_url: "https://news.ycombinator.com/item?id=48470361"
title: "Agentic Engineering Handbook – 115 official OpenAI/Anthropic articles"
article_title: "GitHub - keyuchen21/agentic-engineering-handbook: The definitive OpenAI, Claude, MCP, Harness, Evals, and Production Agent Systems learning roadmap. · GitHub"
author: "keyuchen2020"
captured_at: "2026-06-10T04:35:21Z"
capture_tool: "hn-digest"
hn_id: 48470361
score: 2
comments: 0
posted_at: "2026-06-10T01:57:43Z"
tags:
  - hacker-news
  - translated
---

# Agentic Engineering Handbook – 115 official OpenAI/Anthropic articles

- HN: [48470361](https://news.ycombinator.com/item?id=48470361)
- Source: [github.com](https://github.com/keyuchen21/agentic-engineering-handbook)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T01:57:43Z

## Translation

タイトル: Agentic Engineering ハンドブック – 115 の公式 OpenAI/Anthropic 記事
記事のタイトル: GitHub - keyuchen21/agentic-engineering-handbook: 決定版の OpenAI、Claude、MCP、Harness、Evals、および Production Agent Systems 学習ロードマップ。 · GitHub
説明: OpenAI、Claude、MCP、Harness、Evals、および Production Agent Systems の学習ロードマップの決定版。 - keyuchen21/エージェント エンジニアリング ハンドブック

記事本文:
GitHub - keyuchen21/agentic-engineering-handbook: 決定版の OpenAI、Claude、MCP、Harness、Evals、および Production Agent Systems 学習ロードマップ。 · GitHub
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
キーチェン21
/
エージェント エンジニアリング ハンドブック
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
keyuchen21/エージェントエンジニアリングハンドブック
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI、Anthropic、MCP、Harness、Evals、および Production Agent Systems の学習ロードマップの決定版。
このリポジトリが役立つ場合は、⭐を付けることを検討してください。
AI 業界はエージェント時代に入りました。現在、実稼働グレードの AI システムを構築するには、エージェント、ツールの使用、MCP、メモリ、長時間実行のワークフロー、コーディング エージェント、エージェント ハーネス、評価、安全性を習得する必要があります。しかし、知識は OpenAI ブログ、Anthropic エンジニアリングの投稿、SDK ドキュメント、クックブック、研究論文に散在しています。
このリポジトリは、114 の公式リソースを 1 つの構造化された学習ロードマップに統合します。
目標: 世界クラスのエージェント エンジニアになること。
ワークフローとエージェント、ツール ループ、ハンドオフ、ガードレールに関する共有語彙を構築します。
#
タイトル
ベンダー
1
効果的なエージェントを構築する
人間的
2
エージェントを構築するための新しいツール
OpenAI
3
エージェントSDKの概要
OpenAI
次に読む
タイトル
ベンダー
エージェントのオーケストレーション: ルーチンとハンドオフ
OpenAI
マルチエージェント システムの構造化された出力
OpenAI
ビルド演習
カスタマー サービス/チケット トリアージ エージェントを構築します: ルーター → スペシャリスト → 評価者。すべての出力は構造化されたスキーマによって制約されます。
フェーズ 2 — MCP とツールのエコシステム
MCP サーバー/クライアント、リモートとローカル、ツールの読み込み、承認、コネクタの境界を理解します。
#
タイトル
ベンダー
1
モデル コンテキスト プロトコルの紹介
人間的
2
MCP とコネクタ
OpenAI
3
ChatGPT アプリと API 統合用の MCP サーバーの構築
OpenAI
次に読む
タイトル
ベンダー
MCP を使用したコード実行: より効率的なエージェントの構築
人間的
モデルカンパニー

ntext プロトコル - コーデックス
OpenAI
OpenAI ドキュメント MCP
OpenAI
ビルド演習
読み取り専用のリポジトリ/ドキュメント MCP サーバーを構築し、評価を作成してエージェントがドキュメントを正しく引用していることを確認します。
フェーズ 3 — コンテキスト、記憶、スキル
コンテキスト ウィンドウ、短期/長期記憶、スキル/プラグイン、CLAUDE.md/AGENTS.md の制御方法を学びます。
#
タイトル
ベンダー
1
AI エージェントのための効果的なコンテキスト エンジニアリング
人間的
2
現実世界のエージェントにエージェント スキルを装備する
人間的
3
メモリと圧縮を使用して信頼性の高いエージェントを構築する
OpenAI
次に読む
タイトル
ベンダー
AGENTS.md を使用したカスタム手順 - Codex
OpenAI
クロード コードのベスト プラクティス
人間的
エージェント スキル - コーデックス
OpenAI
ビルド演習
同じタスクをスキル/プラグインとして実装し、スキルなし、長いプロンプト、スキルベースの 3 つのバリエーションにわたって精度とトークン コストを測定します。
フェーズ 4 — ハーネスと長期実行エージェント
マスター エージェント ランタイム: イベント ストリーム、スレッド、ツール実行、状態、サンドボックス、承認、回復。
#
タイトル
ベンダー
1
Codex エージェント ループの展開
OpenAI
2
Codex ハーネスのロックを解除する: App Server の構築方法
OpenAI
3
長期稼働エージェント向けの効果的なハーネス
人間的
次に読む
タイトル
ベンダー
エージェント SDK の次の進化
OpenAI
長時間実行されるアプリケーション開発のためのハーネス設計
人間的
マネージド エージェントのスケーリング: 脳を手から切り離す
人間的
ビルド演習
ミニ コーディング ハーネスを構築します: プラン ファイル、シェル ツール、パッチの適用、テスト ゲート、イベント ログ、再開機能。
フェーズ 5 — コーディングおよびワークスペース エージェント
Codex と Claude Code 製品/SDK フォームを比較します。マルチエージェント、IDE、ワークスペースのコラボレーションを学びます。
#
タイトル
ベンダー
1
コーデックスの紹介
OpenAI
2
クロード コードのベスト プラクティス
人間的
3
クロードコードがより自律的に動作できるようにする
人間的
次に読む
タイトル
ベンダー
Codex アプリの紹介
OpenAI
を紹介します

ChatGPT の rkspace エージェント
OpenAI
Apple の Xcode が Claude Agent SDK をサポートするようになりました
人間的
Codex CLI とエージェント SDK による一貫したワークフローの構築
OpenAI
ビルド演習
OpenAI/Codex と Claude Code スタイルのワークフローの両方を同じリポジトリで実行します: 問題→計画→パッチ→テスト→PR サマリー。
フェーズ 6 — 評価、安全性、生産
起動前/起動後の評価ループ、トレース ループ、安全境界、権限、回帰監視を構築します。
#
タイトル
ベンダー
1
AI エージェントの評価をわかりやすく説明する
人間的
2
Evals を使用してエージェントのスキルを体系的にテストする
OpenAI
3
トレース、評価、コーデックスを使用してエージェント改善ループを構築する
OpenAI
次に読む
タイトル
ベンダー
OpenAI で Codex を安全に実行する
OpenAI
製品全体にクロードを含める方法
人間的
Evals API の使用例 - MCP の評価
OpenAI
AI エージェントの自律性を実際に測定する
人間的
ビルド演習
エージェント用のスモーク/マクロ評価スイートを構築します: タスクの成功率、ツールの誤用、プロンプトインジェクションの抵抗、レイテンシ、コスト、および人間の承認数。
優先ガイド: P0 = 必読 (アーキテクチャ/概念的)、P1 = 非常に役立つ (実装の詳細)、P2 = オプションのコンテキスト (背景/リリース)。
貢献は大歓迎です。見つかった場合:
生産技術記事
このプロジェクトの目標は、エージェントティック エンジニアリングのシステム設計入門書になることです。
実稼働 AI エージェントの構築に真剣に取り組んでいる場合は、ここから始めてください。
OpenAI、Claude、MCP、Harness、Evals、および Production Agent Systems の学習ロードマップの決定版。
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

The definitive OpenAI, Claude, MCP, Harness, Evals, and Production Agent Systems learning roadmap. - keyuchen21/agentic-engineering-handbook

GitHub - keyuchen21/agentic-engineering-handbook: The definitive OpenAI, Claude, MCP, Harness, Evals, and Production Agent Systems learning roadmap. · GitHub
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
keyuchen21
/
agentic-engineering-handbook
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
keyuchen21/agentic-engineering-handbook
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
The definitive OpenAI, Anthropic, MCP, Harness, Evals, and Production Agent Systems learning roadmap.
If this repository helps you, consider giving it a ⭐
The AI industry has entered the Agentic Era . Building production-grade AI systems now requires mastering agents, tool use, MCP, memory, long-running workflows, coding agents, agent harnesses, evals, and safety — but the knowledge is scattered across OpenAI blogs, Anthropic engineering posts, SDK docs, cookbooks, and research papers.
This repository consolidates 114 official resources into one structured learning roadmap.
The goal: Become a world-class Agentic Engineer.
Build shared vocabulary for workflow vs agent, tool loop, handoff, guardrails.
#
Title
Vendor
1
Building effective agents
Anthropic
2
New tools for building agents
OpenAI
3
Agents SDK overview
OpenAI
Then Read
Title
Vendor
Orchestrating Agents: Routines and Handoffs
OpenAI
Structured Outputs for Multi-Agent Systems
OpenAI
Build Exercise
Build a customer service/ticket triage agent: router → specialist → evaluator, with all outputs constrained by structured schemas.
Phase 2 — MCP & Tool Ecosystem
Understand MCP server/client, remote vs local, tool loading, approval, connector boundaries.
#
Title
Vendor
1
Introducing the Model Context Protocol
Anthropic
2
MCP and Connectors
OpenAI
3
Building MCP servers for ChatGPT Apps and API integrations
OpenAI
Then Read
Title
Vendor
Code execution with MCP: Building more efficient agents
Anthropic
Model Context Protocol - Codex
OpenAI
OpenAI Docs MCP
OpenAI
Build Exercise
Build a read-only repo/docs MCP server, then create an eval to verify the agent correctly cites documentation.
Phase 3 — Context, Memory & Skills
Learn to control context window, short/long-term memory, skills/plugins, CLAUDE.md/AGENTS.md.
#
Title
Vendor
1
Effective context engineering for AI agents
Anthropic
2
Equipping agents for the real world with Agent Skills
Anthropic
3
Building Reliable Agents with Memory and Compaction
OpenAI
Then Read
Title
Vendor
Custom instructions with AGENTS.md - Codex
OpenAI
Best practices for Claude Code
Anthropic
Agent Skills - Codex
OpenAI
Build Exercise
Implement the same task as a Skill/Plugin, then measure accuracy and token cost across three variants: no skill, long prompt, and skill-based.
Phase 4 — Harness & Long-Running Agents
Master agent runtime: event stream, thread, tool execution, state, sandbox, approval, recovery.
#
Title
Vendor
1
Unrolling the Codex agent loop
OpenAI
2
Unlocking the Codex harness: how we built the App Server
OpenAI
3
Effective harnesses for long-running agents
Anthropic
Then Read
Title
Vendor
The next evolution of the Agents SDK
OpenAI
Harness design for long-running application development
Anthropic
Scaling Managed Agents: Decoupling the brain from the hands
Anthropic
Build Exercise
Build a mini coding harness: plan file, shell tool, apply patch, test gate, event log, and resume capability.
Phase 5 — Coding & Workspace Agents
Compare Codex vs Claude Code product/SDK forms; learn multi-agent, IDE, workspace collaboration.
#
Title
Vendor
1
Introducing Codex
OpenAI
2
Best practices for Claude Code
Anthropic
3
Enabling Claude Code to work more autonomously
Anthropic
Then Read
Title
Vendor
Introducing the Codex app
OpenAI
Introducing workspace agents in ChatGPT
OpenAI
Apple's Xcode now supports Claude Agent SDK
Anthropic
Building Consistent Workflows with Codex CLI & Agents SDK
OpenAI
Build Exercise
Run both OpenAI/Codex and Claude Code style workflows on the same repo: issue → plan → patch → tests → PR summary.
Phase 6 — Evals, Safety & Production
Build pre/post-launch eval loop, trace loop, safety boundaries, permissions, regression monitoring.
#
Title
Vendor
1
Demystifying evals for AI agents
Anthropic
2
Testing Agent Skills Systematically with Evals
OpenAI
3
Build an Agent Improvement Loop with Traces, Evals, and Codex
OpenAI
Then Read
Title
Vendor
Running Codex safely at OpenAI
OpenAI
How we contain Claude across products
Anthropic
Evals API Use-case - MCP Evaluation
OpenAI
Measuring AI agent autonomy in practice
Anthropic
Build Exercise
Build a smoke/macro eval suite for your agent: task success rate, tool misuse, prompt injection resistance, latency, cost, and human approval count.
Priority guide: P0 = must-read (architectural/conceptual), P1 = highly useful (implementation detail), P2 = optional context (background/releases).
Contributions are welcome. If you find:
Production engineering articles
The goal of this project is to become the System Design Primer for Agentic Engineering.
If you're serious about building production AI agents, start here.
The definitive OpenAI, Claude, MCP, Harness, Evals, and Production Agent Systems learning roadmap.
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

---
source: "https://agentapplication.io"
hn_url: "https://news.ycombinator.com/item?id=49376946"
title: "Agent Applications: A Reference Architecture for AI Agent Systems"
article_title: "Agent Applications - Agent Applications"
image: "https://agentapplication.mintlify.app/mintlify-assets/_next/image?url=%2F_mintlify%2Fapi%2Fog%3Fdivision%3DOverview%26title%3DAgent%2BApplications%26description%3DA%2Bvendor-neutral%2Breference%2Barchitecture%2Bfor%2Bbuilding%2Band%2Boperating%2BAI%2Bagent%2Bsystems.%26primaryColor%3D%25235a5a14%26lightColor%3D%2523e8e13a%26backgroundLight%3D%2523fdfdfb%26backgroundDark%3D%2523161613&w=1200&q=100"
author: "amthewiz"
captured_at: "2026-08-20T17:20:51Z"
capture_tool: "hn-digest"
hn_id: 49376946
score: 2
comments: 1
posted_at: "2026-08-20T16:35:49Z"
tags:
  - hacker-news
  - translated
---

# Agent Applications: A Reference Architecture for AI Agent Systems

- HN: [49376946](https://news.ycombinator.com/item?id=49376946)
- Source: [agentapplication.io](https://agentapplication.io)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T16:35:49Z

## Translation

タイトル: エージェント アプリケーション: AI エージェント システムのリファレンス アーキテクチャ
記事のタイトル: エージェント アプリケーション - エージェント アプリケーション
説明: AI エージェント システムを構築および運用するためのベンダー中立のリファレンス アーキテクチャ。

記事本文:
エージェント アプリケーション - エージェント アプリケーションのドキュメント インデックス
/llms.txt で完全なドキュメントのインデックスを取得します。
さらに探索する前に、このファイルを使用して利用可能なすべてのページを検出します。
メイン コンテンツにスキップ エージェント アプリケーションのホームページ エージェント アプリケーションの検索... ⌘ K エージェント アプリケーション/エージェント アプリケーション
エージェントアプリケーション/エージェントアプリケーション
検索... ナビゲーションの概要 エージェント アプリケーションの概要
ページをコピー ページをコピー AI エージェント システムを構築および運用するためのベンダー中立のリファレンス アーキテクチャ。
ページをコピー ページをコピー ワーキングペーパー · ドラフト 0.9 · 2026 年 8 月 · 論文を読む · ドラフトをレビューする · 引用
エージェント アプリケーションは構築された AI システムです
永続的なツールを使用するエージェントの周囲。リクエスト、イベント、スケジュールを開始できます
仕事。エージェントはアプリケーション ガードレール内の中間アクションを選択します。
そのインスタンスにはワークスペース、権限、アーティファクト、および未完成のデータが保持されます。
走行終了後の作業。
このワーキングペーパーでは、ベンダー中立のリファレンス アーキテクチャが示されています。
これらのシステムを構築および運用します。主要な要素を特定し、割り当てます
彼らに対する責任を負い、リリースまでエージェント プロジェクトに従います。
導入、運用、変更。また、モデルなどの標準も検索します。
コンテキスト プロトコル (MCP) とエージェント
カバーできる限界のスキル。
この論文は作業草案です。読者の皆様には、その主張を実際の製品、実装経験、先行技術、反例と照らし合わせてテストしていただくようお勧めします。
紙があなたに与えるもの
プロジェクト、リリース、インスタンス、ワークスペース、成果物、
フレームワーク、プラットフォーム、ストア
を回転させるための構築および運用方法
ユースケースをテスト済みのアプリケーション設計に組み込む
の機能モデル
フレームワークの比較と責任の割り当て
長期運用のためのライフサイクル、セキュリティ、ガバナンスの視点

ライブされたインスタンス
蓄積された仕事を失うことなく
既存の標準と将来の交換境界についての説明
契約書が役立つかもしれない
この文書では、完全な議論、カテゴリの定義、およびアーキテクチャを示します。
Hello World は、小さなアプリケーションでモデルを示します。
この論文については、対象者、範囲、作業の状況について説明します。
付属文書には、この論文に基づいて計画されている実践的および技術的な文書がリストされています。
自由形式の提案、アーキテクチャに関する質問、および境界ケースについては、GitHub ディスカッションを使用してください。
具体的なエラー、ソースの欠落、または破損したページについての問題をオープンします。
焦点を絞った修正、サンプル、またはドキュメントの改善については、プル リクエストを送信してください。

## Original Extract

A vendor-neutral reference architecture for building and operating AI agent systems.

Agent Applications - Agent Applications Documentation Index
Fetch the complete documentation index at: /llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Agent Applications home page Agent Applications Search... ⌘ K agentapplication/agentapplication
agentapplication/agentapplication
Search... Navigation Overview Agent Applications Overview
Copy page Copy page A vendor-neutral reference architecture for building and operating AI agent systems.
Copy page Copy page A working paper · Draft 0.9 · August 2026 · Read the paper · Review the draft · Cite
Agent Applications are AI systems built
around persistent, tool-using agents. A request, event, or schedule can start
work. The agent chooses intermediate actions within application guardrails,
while its instance retains a workspace, authority, artifacts, and unfinished
work after the run ends.
The working paper presents a vendor-neutral reference architecture for
building and operating these systems. It identifies the main elements, assigns
responsibilities to them, and follows an Agent Project through release,
deployment, operation, and change. It also locates standards such as the Model
Context Protocol (MCP) and Agent
Skills at the boundaries they cover.
The paper is a working draft. We invite readers to test its claims against real products, implementation experience, prior art, and counterexamples.
​ What the paper gives you
A common vocabulary for projects, releases, instances, workspaces, artifacts,
frameworks, platforms, and stores
A build-and-operate method for turning a
use case into a tested application design
A capability model for
comparing frameworks and assigning responsibility
Lifecycle, security, and governance views for operating long-lived instances
without losing their accumulated work
An account of existing standards and the exchange boundaries where future
contracts may help
The paper presents the full argument, category definition, and architecture.
Hello World shows the model in a small application.
About this paper explains the audience, scope, and status of the work.
Companion documents lists the practical and technical documents planned around the paper.
Use GitHub Discussions for open-ended proposals, architecture questions, and boundary cases.
Open an issue for a concrete error, missing source, or broken page.
Send a pull request for a focused correction, example, or documentation improvement.

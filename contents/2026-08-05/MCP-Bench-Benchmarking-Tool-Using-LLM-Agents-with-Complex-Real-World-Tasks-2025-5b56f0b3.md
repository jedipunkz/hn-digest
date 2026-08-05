---
source: "https://arxiv.org/abs/2508.20453"
hn_url: "https://news.ycombinator.com/item?id=49183366"
title: "MCP-Bench: Benchmarking Tool-Using LLM Agents with Complex Real-World Tasks 2025"
article_title: "[2508.20453] MCP-Bench: Benchmarking Tool-Using LLM Agents with Complex Real-World Tasks via MCP Servers"
author: "janandonly"
captured_at: "2026-08-05T15:09:41Z"
capture_tool: "hn-digest"
hn_id: 49183366
score: 2
comments: 1
posted_at: "2026-08-05T14:21:44Z"
tags:
  - hacker-news
  - translated
---

# MCP-Bench: Benchmarking Tool-Using LLM Agents with Complex Real-World Tasks 2025

- HN: [49183366](https://news.ycombinator.com/item?id=49183366)
- Source: [arxiv.org](https://arxiv.org/abs/2508.20453)
- Score: 2
- Comments: 1
- Posted: 2026-08-05T14:21:44Z

## Translation

タイトル: MCP-Bench: ベンチマーク ツール - 複雑な現実世界のタスクで LLM エージェントを使用する 2025
記事のタイトル: [2508.20453] MCP-Bench: ベンチマーク ツール - MCP サーバーを介した複雑な現実世界のタスクで LLM エージェントを使用する
説明: arXiv 論文 2508.20453 の要約ページ: MCP-Bench: ベンチマーク ツール - MCP サーバーを介した複雑な現実世界のタスクで LLM エージェントを使用する

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2025 年 8 月 28 日に提出]
タイトル: MCP-Bench: ベンチマーク ツール - MCP サーバー経由で現実世界の複雑なタスクを実行する LLM エージェントを使用する
要約: ツールの使用、ツール間の調整、正確なパラメーター制御、およびタスクを解決するための計画/推論を必要とする現実的な複数ステップのタスクに関する大規模言語モデル (LLM) を評価するためのベンチマークである MCP-Bench を紹介します。 Model Context Protocol (MCP) 上に構築された MCP-Bench は、金融、旅行、科学技術コンピューティング、学術検索などの分野にわたる 250 のツールにわたる 28 の代表的なライブ MCP サーバーに LLM を接続します。以前の API ベースのベンチマークとは異なり、各 MCP サーバーは連携して動作するように設計された一連の補完的なツールを提供し、豊富な入出力結合を備えた本格的なマルチステップ タスクの構築を可能にします。 MCP-Bench のタスクは、明示的なツール名のないあいまいな命令から関連ツールを取得するエージェントの機能、複雑な目的のためのマルチホップ実行軌跡の計画、中間ツール出力の応答の調整、およびクロスドメイン ワークフローの調整を行うエージェントの機能をテストします。これらの機能は、明示的なツール仕様、浅い数ステップのワークフロー、および分離されたドメイン操作に依存する既存のベンチマークでは適切に評価されていません。私たちは、ツールレベルのスキーマの理解と使用法、軌道レベルの計画、タスクの完了をカバーする多面的な評価フレームワークを提案します。 20 の高度な LLM の実験により、MCP-Bench における永続的な課題が明らかになりました。コードとデータ: この https URL 。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
コード、データ

この記事に関連するメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2508.20453: MCP-Bench: Benchmarking Tool-Using LLM Agents with Complex Real-World Tasks via MCP Servers

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 28 Aug 2025]
Title: MCP-Bench: Benchmarking Tool-Using LLM Agents with Complex Real-World Tasks via MCP Servers
Abstract: We introduce MCP-Bench, a benchmark for evaluating large language models (LLMs) on realistic, multi-step tasks that demand tool use, cross-tool coordination, precise parameter control, and planning/reasoning for solving tasks. Built on the Model Context Protocol (MCP), MCP-Bench connects LLMs to 28 representative live MCP servers spanning 250 tools across domains such as finance, traveling, scientific computing, and academic search. Unlike prior API-based benchmarks, each MCP server provides a set of complementary tools designed to work together, enabling the construction of authentic, multi-step tasks with rich input-output coupling. Tasks in MCP-Bench test agents' ability to retrieve relevant tools from fuzzy instructions without explicit tool names, plan multi-hop execution trajectories for complex objectives, ground responses in intermediate tool outputs, and orchestrate cross-domain workflows - capabilities not adequately evaluated by existing benchmarks that rely on explicit tool specifications, shallow few-step workflows, and isolated domain operations. We propose a multi-faceted evaluation framework covering tool-level schema understanding and usage, trajectory-level planning, and task completion. Experiments on 20 advanced LLMs reveal persistent challenges in MCP-Bench. Code and data: this https URL .
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

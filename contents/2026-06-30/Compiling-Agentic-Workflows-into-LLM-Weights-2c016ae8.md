---
source: "https://arxiv.org/abs/2605.22502"
hn_url: "https://news.ycombinator.com/item?id=48731576"
title: "Compiling Agentic Workflows into LLM Weights"
article_title: "[2605.22502] Compiling Agentic Workflows into LLM Weights: Near-Frontier Quality at Two Orders of Magnitude Less Cost"
author: "dipankarsarkar"
captured_at: "2026-06-30T12:37:18Z"
capture_tool: "hn-digest"
hn_id: 48731576
score: 1
comments: 0
posted_at: "2026-06-30T12:12:33Z"
tags:
  - hacker-news
  - translated
---

# Compiling Agentic Workflows into LLM Weights

- HN: [48731576](https://news.ycombinator.com/item?id=48731576)
- Source: [arxiv.org](https://arxiv.org/abs/2605.22502)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T12:12:33Z

## Translation

タイトル: エージェント ワークフローを LLM 重みにコンパイルする
記事のタイトル: [2605.22502] エージェント ワークフローを LLM 重みにコンパイル: 2 桁低いコストでニアフロンティアの品質を実現
説明: arXiv 論文 2605.22502 の要約ページ: LLM 重みへのエージェント ワークフローのコンパイル: 2 桁低いコストでのニアフロンティア品質

記事本文:
メインコンテンツにスキップ
arXiv の提出は、6 月 30 日火曜日 14:00 EDT からメンテナンスのため停止されます。それ以外の場合は、サイトは稼働し続ける必要があります。
コンピュータサイエンス > 人工知能
[2026 年 5 月 21 日に提出]
タイトル: エージェント ワークフローを LLM 重みにコンパイル: 2 桁低いコストでニアフロンティアの品質を実現
要約: エージェント オーケストレーション フレームワークは急増しており、LangGraph、CrewAI、Google ADK、OpenAI Agents SDK、Semantic Kernel、Strands、LlamaIndex 全体で GitHub スターの数は合計 290,000 を超えています。すべては同じパターンに従います。つまり、LLM の上に外部オーケストレーターがあり、毎ターン命令を挿入し、決定をルーティングします。最近の研究では、このアーキテクチャがフロンティア モデルのシステム プロンプトでプロシージャを提供するだけで手続き型タスクに優勢であることが示されています [Dennis et al., 2026a]。その代償として、コンテキスト ウィンドウが消費され、すべての会話にフロンティア モデルが必要になり、独自のプロシージャがサードパーティ プロバイダーに公開されます。この手順を微調整された小さなモデルの重みにコンパイルする (地下エージェントを作成する) ことで、これらの懸念事項はすべて解決されるはずです。以前の研究 (SimpleTOD、FireAct、SynTOD、WorkflowLLM、Agent Lumos) では、この手法が機能することが示されています。しかし、開発者の採用ではオーケストレーションが圧倒的に支持されています。当社は、旅行予約 (14 ノード)、Zoom サポート (14 ノード、製品固有の知識)、および保険請求 (55 ノード、6 つの意思決定ハブ) にわたって、認識されている 3 つの障壁を特定し、経験的にそれぞれに対処します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が開発および共有できるようにするフレームワークです。

arXiv の新しい機能は Web サイトで直接ご覧いただけます。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2605.22502: Compiling Agentic Workflows into LLM Weights: Near-Frontier Quality at Two Orders of Magnitude Less Cost

Skip to main content
arXiv submission will be down for maintenance beginning 14:00 EDT Tuesday June 30th. The site should otherwise remain in operation.
Computer Science > Artificial Intelligence
[Submitted on 21 May 2026]
Title: Compiling Agentic Workflows into LLM Weights: Near-Frontier Quality at Two Orders of Magnitude Less Cost
Abstract: Agent orchestration frameworks have proliferated, collectively exceeding 290,000 GitHub stars across LangGraph, CrewAI, Google ADK, OpenAI Agents SDK, Semantic Kernel, Strands, and LlamaIndex. All follow the same pattern: an external orchestrator above the LLM, injecting instructions and routing decisions every turn. Recent work has shown this architecture is dominated for procedural tasks by simply providing the procedure in a frontier model's system prompt [Dennis et al., 2026a], at the cost of consuming the context window, requiring a frontier model for every conversation, and exposing proprietary procedures to third-party providers. Compiling the procedure into the weights of a small fine-tuned model -- creating a subterranean agent -- should resolve all of these concerns, and prior work (SimpleTOD, FireAct, SynTOD, WorkflowLLM, Agent Lumos) has shown the technique works. Yet developer adoption has overwhelmingly favored orchestration. We identify three perceived barriers and address each empirically across travel booking (14 nodes), Zoom support (14 nodes, product-specific knowledge), and insurance claims (55 nodes, 6 decision hubs).
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe

---
source: "https://arxiv.org/abs/2607.29254"
hn_url: "https://news.ycombinator.com/item?id=49157192"
title: "Tool Specifications Matter: Uncovering and Mitigating Safety Risks in AI Agents"
article_title: "[2607.29254] Tool Specifications Matter: Uncovering and Mitigating Safety Risks in AI Agents"
author: "zhinit"
captured_at: "2026-08-03T15:37:32Z"
capture_tool: "hn-digest"
hn_id: 49157192
score: 1
comments: 0
posted_at: "2026-08-03T15:37:00Z"
tags:
  - hacker-news
  - translated
---

# Tool Specifications Matter: Uncovering and Mitigating Safety Risks in AI Agents

- HN: [49157192](https://news.ycombinator.com/item?id=49157192)
- Source: [arxiv.org](https://arxiv.org/abs/2607.29254)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T15:37:00Z

## Translation

タイトル: ツールの仕様が重要: AI エージェントの安全性リスクの発見と軽減
記事のタイトル: [2607.29254] ツール仕様が重要: AI エージェントの安全性リスクの発見と軽減
説明: arXiv 論文 2607.29254 の要約ページ: ツールの仕様に関する事項: AI エージェントの安全性リスクの発見と軽減

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 7 月 31 日に提出]
タイトル: ツールの仕様が重要: AI エージェントの安全性リスクの発見と軽減
要約: AI エージェントは外部ツールを使用して大規模言語モデル (LLM) を拡張し、複雑なタスクを実行し、モデルの出力を結果として現実世界のアクションに変換できるようにします。しかし、LLM はエージェントとして導入されると安全性が大幅に低下することが多く、この低下の原因は依然としてよくわかっていません。この論文では、スキーマ形式のツール仕様がエージェントの安全性低下の主な原因であることを特定し、ホワイトボックス表現分析を通じて、それらがモデルの内部拒否シグナルを弱め、安全でないツールの実行に寄与していることを示します。この発見に基づいて、私たちは安全性の判断をツールの実行から切り離す推論時の保護手段である SafeKeep を提案します。これは、元のスキーマ形式の実行仕様を保持しながら、平坦化されたテキストのツール仕様を使用してリクエストを評価します。 2 つの代表的なベンチマークと、ホワイト ボックス モデルとブラック ボックス モデルの両方を含む 4 つの LLM にわたって、SafeKeep は有害なリクエストの平均拒否率を 23.8% から 70.6% に増加させ、観測レベルのプロンプト インジェクションの下での平均攻撃成功率を 25.6% から 2.5% に減少させます。また、既存の安全対策よりも優れたパフォーマンスを発揮し、タスク処理能力を維持します。この https URL でコードとデータを公開します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs はフレームワーです

k これにより、共同作業者は新しい arXiv 機能を開発し、Web サイト上で直接共有できるようになります。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.29254: Tool Specifications Matter: Uncovering and Mitigating Safety Risks in AI Agents

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 31 Jul 2026]
Title: Tool Specifications Matter: Uncovering and Mitigating Safety Risks in AI Agents
Abstract: AI agents extend large language models (LLMs) with external tools, enabling them to perform complex tasks and translate model outputs into consequential real-world actions. Yet LLMs often become substantially less safe when deployed as agents, and the source of this degradation remains poorly understood. In this paper, we identify schema-formatted tool specifications as a primary source of agent safety degradation and show, through white-box representation analysis, that they weaken the model's internal refusal signals and contribute to unsafe tool execution. Building on this finding, we propose SafeKeep, an inference-time safeguard that decouples safety judgment from tool execution: it assesses requests using flattened textual tool specifications while retaining the original schema-formatted specifications for execution. Across two representative benchmarks and four LLMs, including both white-box and black-box models, SafeKeep increases the average refusal rate for harmful requests from 23.8% to 70.6% and reduces the average attack success rate under observation-level prompt injection from 25.6% to 2.5%. It also outperforms existing safeguards and preserves task-handling capability. We release the code and data at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

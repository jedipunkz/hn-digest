---
source: "https://arxiv.org/abs/2602.10133"
hn_url: "https://news.ycombinator.com/item?id=49108689"
title: "AgentTrace: A Structured Logging Framework for Agent System Observability"
article_title: "[2602.10133] AgentTrace: A Structured Logging Framework for Agent System Observability"
author: "ankitg12"
captured_at: "2026-07-30T12:24:05Z"
capture_tool: "hn-digest"
hn_id: 49108689
score: 1
comments: 0
posted_at: "2026-07-30T11:46:15Z"
tags:
  - hacker-news
  - translated
---

# AgentTrace: A Structured Logging Framework for Agent System Observability

- HN: [49108689](https://news.ycombinator.com/item?id=49108689)
- Source: [arxiv.org](https://arxiv.org/abs/2602.10133)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T11:46:15Z

## Translation

タイトル: AgentTrace: エージェント システムの可観測性のための構造化ログ フレームワーク
記事のタイトル: [2602.10133] AgentTrace: エージェント システムの可観測性のための構造化ログ フレームワーク
説明: arXiv 論文 2602.10133 の要約ページ: AgentTrace: エージェント システムの可観測性のための構造化ログ フレームワーク

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
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 2 月 7 日に提出]
タイトル: AgentTrace: エージェント システムの可観測性のための構造化ログ フレームワーク
要約: 大規模言語モデル (LLM) を活用した自律エージェントの機能が向上しているにもかかわらず、一か八かの分野での自律エージェントの採用は依然として限られています。重要な障壁はセキュリティです。LLM エージェントの本質的に非決定的な動作は、歴史的にソフトウェア アシュアランスを支えてきた静的監査アプローチに反するものです。プロキシレベルの入力フィルタリングやモデルのグラスボックス化などの既存のセキュリティ手法では、エージェントの推論、状態の変化、環境の相互作用に対する十分な透明性や追跡可能性を提供できません。この作業では、このギャップを埋めるために設計された動的な可観測性およびテレメトリ フレームワークである AgentTrace を紹介します。 AgentTrace は、最小限のオーバーヘッドで実行時にエージェントを計測し、操作面、認知面、コンテキスト面の 3 つの面にわたって構造化ログの豊富なストリームをキャプチャします。従来のログ システムとは異なり、AgentTrace は継続的で内省可能なトレース キャプチャを重視しており、デバッグやベンチマークのためだけでなく、エージェントのセキュリティ、アカウンタビリティ、およびリアルタイム監視の基礎層として設計されています。私たちの研究は、AgentTrace がどのようにしてより信頼性の高いエージェントの展開、きめ細かいリスク分析、情報に基づいた信頼の調整を可能にし、それによってこれまで機密性の高い環境での LLM エージェントの使用を制限していた重大な懸念に対処できるかを明らかにしています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs:

コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2602.10133: AgentTrace: A Structured Logging Framework for Agent System Observability

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
Computer Science > Software Engineering
[Submitted on 7 Feb 2026]
Title: AgentTrace: A Structured Logging Framework for Agent System Observability
Abstract: Despite the growing capabilities of autonomous agents powered by large language models (LLMs), their adoption in high-stakes domains remains limited. A key barrier is security: the inherently nondeterministic behavior of LLM agents defies static auditing approaches that have historically underpinned software assurance. Existing security methods, such as proxy-level input filtering and model glassboxing, fail to provide sufficient transparency or traceability into agent reasoning, state changes, or environmental interactions. In this work, we introduce AgentTrace, a dynamic observability and telemetry framework designed to fill this gap. AgentTrace instruments agents at runtime with minimal overhead, capturing a rich stream of structured logs across three surfaces: operational, cognitive, and contextual. Unlike traditional logging systems, AgentTrace emphasizes continuous, introspectable trace capture, designed not just for debugging or benchmarking, but as a foundational layer for agent security, accountability, and real-time monitoring. Our research highlights how AgentTrace can enable more reliable agent deployment, fine-grained risk analysis, and informed trust calibration, thereby addressing critical concerns that have so far limited the use of LLM agents in sensitive environments.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

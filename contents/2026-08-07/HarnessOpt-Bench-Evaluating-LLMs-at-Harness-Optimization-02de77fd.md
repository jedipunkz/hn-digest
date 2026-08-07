---
source: "https://arxiv.org/abs/2608.06301"
hn_url: "https://news.ycombinator.com/item?id=49213543"
title: "HarnessOpt-Bench: Evaluating LLMs at Harness Optimization"
article_title: "[2608.06301] HarnessOpt-Bench: Evaluating LLMs at Harness Optimization"
author: "wslh"
captured_at: "2026-08-07T17:40:39Z"
capture_tool: "hn-digest"
hn_id: 49213543
score: 2
comments: 0
posted_at: "2026-08-07T17:20:13Z"
tags:
  - hacker-news
  - translated
---

# HarnessOpt-Bench: Evaluating LLMs at Harness Optimization

- HN: [49213543](https://news.ycombinator.com/item?id=49213543)
- Source: [arxiv.org](https://arxiv.org/abs/2608.06301)
- Score: 2
- Comments: 0
- Posted: 2026-08-07T17:20:13Z

## Translation

タイトル: HarnessOpt-Bench: ハーネス最適化における LLM の評価
記事のタイトル: [2608.06301] HarnessOpt-Bench: ハーネス最適化における LLM の評価
説明: arXiv 論文 2608.06301 の要約ページ: HarnessOpt-Bench: ハーネス最適化における LLM の評価

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
[2026 年 8 月 6 日に提出]
タイトル: HarnessOpt-Bench: ハーネス最適化における LLM の評価
要約: LLM がエージェント システム内に導入されることが増えているため、LLM の機能はモデルの重みだけでなく、ハーネス (プロンプト、ツール、制御フロー、メモリ、それらを囲むオーケストレーション コード) にも依存します。このため、ハーネスの自動最適化 (AI システムによる反復的かつ評価に基づくハーネスの改善) は、AI システムを改善するための重要な手段であると同時に、AI システム自体に要求の厳しい機能でもあります。しかし、コミュニティには、フロンティア LLM がこのタスクでどれだけうまく機能するかを測定するための共通のプロトコルがありません。高価で確率的な評価に基づくエンドツーエンドのハーネス最適化のベンチマークである HarnessOpt-Bench を紹介します。コーディング ハーネスとペアになった LLM であるオプティマイザーは、ターゲット エージェントのシード ハーネス、段階的な評価フィードバック、および固定のターゲット評価予算を受け取ります。ハーネスを編集し、最終候補を指名します。この候補は、検索中アクセスできないまま保持されているテスト パーティション上のシードに対する正規化されたゲインによってスコア付けされます。信頼できる実行環境は、評価境界を強制し、ターゲット エージェントのリソース使用量を計測し、監査用に候補バージョンを保存します。共有コーディング ハーネスとネイティブ ハーネスの両方で、4 つのダウンストリーム タスクにわたって 111 回のスコアリング実行で、5 つのフロンティア LLM をオプティマイザーとして評価しました。実験の結果、オプティマイザ モデルは、動作するコーディング ハーネス以上に分離しており、ネイティブ ハーネスが一貫して優れているわけではなく、ゲインはタスクやシード レジームによって大幅に異なることが示されています。これらの結果

ハーネスの最適化は、改善の余地が大きい測定可能で識別可能な機能として確立されます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.06301: HarnessOpt-Bench: Evaluating LLMs at Harness Optimization

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 6 Aug 2026]
Title: HarnessOpt-Bench: Evaluating LLMs at Harness Optimization
Abstract: As LLMs are increasingly deployed within agentic systems, their capabilities depend not only on the model weights but also on the harness: the prompts, tools, control flow, memory, and orchestration code surrounding them. This makes automated harness optimization -- the iterative and evaluation-guided improvement of a harness by an AI system -- both an important route to improving AI systems and a demanding capability for AI systems themselves. Yet the community lacks a common protocol for measuring how well frontier LLMs perform at this task. We introduce HarnessOpt-Bench, a benchmark for end-to-end harness optimization under expensive and stochastic evaluation. An optimizer, an LLM paired with a coding harness, receives a target agent's seed harness, graded evaluation feedback, and a fixed target-evaluation budget. It edits the harness and nominates a final candidate, which is scored by its normalized gain over the seed on a held-out test partition that remains inaccessible throughout search. A trusted execution environment enforces the evaluation boundary, meters target-agent resource use, and preserves candidate versions for audit. We evaluate 5 frontier LLMs as optimizers both under a shared coding harness and under their native harnesses across 4 downstream tasks, over 111 scored runs. Experiment results show that optimizer models separate more than the coding harnesses they act through, native harnesses are not consistently superior, and gains vary substantially across tasks and seed regimes. These results establish harness optimization as a measurable and discriminative capability with large space for improvement.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

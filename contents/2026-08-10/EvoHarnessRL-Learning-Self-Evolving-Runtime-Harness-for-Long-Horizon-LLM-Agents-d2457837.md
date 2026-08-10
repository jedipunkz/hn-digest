---
source: "https://arxiv.org/abs/2608.05446"
hn_url: "https://news.ycombinator.com/item?id=49249998"
title: "EvoHarnessRL: Learning Self-Evolving Runtime Harness for Long-Horizon LLM Agents"
article_title: "[2608.05446] EvoHarness-RL: Learning Self-Evolving Runtime Harness for Long-Horizon LLM Agents"
author: "matt_d"
captured_at: "2026-08-10T21:32:33Z"
capture_tool: "hn-digest"
hn_id: 49249998
score: 1
comments: 0
posted_at: "2026-08-10T21:28:07Z"
tags:
  - hacker-news
  - translated
---

# EvoHarnessRL: Learning Self-Evolving Runtime Harness for Long-Horizon LLM Agents

- HN: [49249998](https://news.ycombinator.com/item?id=49249998)
- Source: [arxiv.org](https://arxiv.org/abs/2608.05446)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T21:28:07Z

## Translation

タイトル: EvoHarnessRL: Long-Horizon LLM エージェント向けの自己進化型ランタイム ハーネスの学習
記事のタイトル: [2608.05446] EvoHarness-RL: Long-Horizon LLM エージェント向けの自己進化型ランタイム ハーネスの学習
説明: arXiv 論文 2608.05446 の要約ページ: EvoHarness-RL: Long-Horizon LLM エージェントのための自己進化型ランタイム ハーネスの学習

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 機械学習
[2026 年 8 月 5 日に提出]
タイトル: EvoHarness-RL: Long-Horizon LLM エージェント向けの自己進化型ランタイム ハーネスの学習
要約: Long-horizon LLM エージェントは、状態を維持し、進行状況を追跡し、ツールを呼び出し、結果を検証し、インタラクション全体でエクスペリエンスを再利用するために、外部の実行サポートにますます依存しています。ただし、ハーネスを効果的に使用すると、ノイズの多い対話トレースによる状態形成と、外部状態アクセスに対するランタイム制御という 2 つの課題が組み合わさって発生します。既存のエージェントは通常、プロンプト、ヒューリスティック、またはドメイン固有の規則を通じて両方を処理し、外部ワークスペースとその使用ポリシーは手動で設計されたままになります。これに対処するために、エージェントがオフラインでハーネス ポリシーを学習し、それを展開してランタイム タスクの実行中にオンラインで外部ハーネスの状態を構築および更新する、ハーネス ポリシーの学習の問題を研究します。 EvoHarness-RL を導入します。これは、信念、進歩、経験 (BPE) をポリシー対応のハーネス状態として公開します。監視付きハーネス微調整は、ベース エージェントにハーネス アクション スペースと有用な外部状態の構築方法を教え、コストを意識した GRPO は、長期的な対話中にその状態を選択的に読み取り、更新、統合するための調整ポリシーを探索します。 Qwen3-8B LLM を使用して ALFWorld 上でインスタンス化された EvoHarness-RL は 96.9% の成功率に達し、2 つの重要なダイナミクスを明らかにしました。1 つはトレーニングによって繰り返し発生するハーネス使用パターンをモデル ポリシーに取り込み、エージェントを頻繁なハーネス呼び出しから選択的な外部状態アクセスに移行させるハーネス アニーリング、もう 1 つはハーネスの進化で、進捗状況の更新とエクスペリエンスの統合によりハーネスをコンパクトなタスク適応型に洗練します。

状態基板。これらの結果は、長期的なエージェントは、単に強力なツールや大容量のメモリを追加するだけでなく、外部ハーネス ワークスペースを構築および調整するためのトレーニング可能なポリシーから恩恵を受けることを示唆しています。
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

Abstract page for arXiv paper 2608.05446: EvoHarness-RL: Learning Self-Evolving Runtime Harness for Long-Horizon LLM Agents

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 5 Aug 2026]
Title: EvoHarness-RL: Learning Self-Evolving Runtime Harness for Long-Horizon LLM Agents
Abstract: Long-horizon LLM agents increasingly rely on external execution support to maintain state, track progress, invoke tools, verify outcomes, and reuse experience across interactions. However, effective harness use raises two coupled challenges: state formation from noisy interaction traces and runtime control over external-state access. Existing agents usually handle both through prompts, heuristics, or domain-specific conventions, leaving the external workspace and its usage policy manually engineered. To address this, we study the problem of harness policy learning, where agents learn harness policies offline and deploy them to construct and update external harness state online during runtime task execution. We introduce EvoHarness-RL, which exposes Belief, Progress, and Experience (BPE) as policy-facing harness state. Supervised harness fine-tuning teaches the base agent the harness action space and how to construct useful external state, while cost-aware GRPO explores coordination policies to selectively read, update, and consolidate that state during long-horizon interaction. Instantiated on ALFWorld with a Qwen3-8B LLM, EvoHarness-RL reaches 96.9% success and reveals two key dynamics: harness annealing, where training internalizes recurring harness-use patterns into the model policy and shifts the agent from frequent harness calls toward selective external-state access, and harness evolution, where progress updates and experience consolidation refine the harness into a compact, task-adaptive state substrate. These results suggest that long-horizon agents benefit from trainable policies for constructing and coordinating with external harness workspaces, beyond simply adding stronger tools or larger memories.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

---
source: "https://arxiv.org/abs/2606.03489"
hn_url: "https://news.ycombinator.com/item?id=48428558"
title: "Learn from Your Mistakes: Tree-Like Self-Play for Secure Code LLMs"
article_title: "[2606.03489] Learn from Your Mistakes: Tree-like Self-Play for Secure Code LLMs"
author: "Extropy_"
captured_at: "2026-06-06T21:26:44Z"
capture_tool: "hn-digest"
hn_id: 48428558
score: 1
comments: 0
posted_at: "2026-06-06T20:13:16Z"
tags:
  - hacker-news
  - translated
---

# Learn from Your Mistakes: Tree-Like Self-Play for Secure Code LLMs

- HN: [48428558](https://news.ycombinator.com/item?id=48428558)
- Source: [arxiv.org](https://arxiv.org/abs/2606.03489)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T20:13:16Z

## Translation

タイトル: 間違いから学ぶ: セキュア コード LLM のためのツリー状セルフプレイ
記事のタイトル: [2606.03489] 間違いから学ぶ: セキュア コード LLM のツリー状セルフプレイ
説明: arXiv 論文 2606.03489 の要約ページ: 間違いから学ぶ: セキュア コード LLM のためのツリー状のセルフプレイ

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.03489
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 2 日に提出]
タイトル: 間違いから学ぶ: セキュア コード LLM のツリー状セルフプレイ
要約: 大規模言語モデル (LLM) はコード生成に優れていますが、トレーニング データに固有の微妙だが重大な脆弱性を複製する傾向が依然としてあります。教師あり微調整 (SFT) や強化学習 (RL) などの現在のアライメント技術は、通常、シーケンス レベルで粗粒度の最適化を適用します。このアプローチでは、セキュリティ上の欠陥の局所的な性質に対処できないことが多く、トークンの選択を 1 つ間違えるとプログラム全体が危険にさらされる可能性があります。このギャップを埋めるために、安全なコード生成をきめ細かい逐次決定プロセスとして再構成するフレームワークである Tree-like Self-Play (TSP) を導入します。やみくもに尤度を最大化する標準的な方法とは異なり、TSP はモデルが分岐軌道を探索する決定木を構築し、安全な「ゴールデン パス」と脆弱なバリアントの両方を生成します。コード生成をセルフプレイ ゲームとして扱うことにより、モデルはそれ自身の局所的なエラーを厳密に区別することを学習します。これにより、脆弱性が通常現れる重要な意思決定ノードで正確に自己修正を強制する、高密度のポリシーに基づいた学習信号が提供されます。私たちの実験は、TSP がモデルの信頼性を根本的に強化することを示しています。 Python セキュリティ ベンチマークでは、TSP は CodeLlama-7B の合格率 (SPR@1) を 75.8% に押し上げ、SFT (57.0%) および非構造化セルフプレイ ベースラインを大幅に上回りました。重要なのは、TSP が強力なアウトオブ

配布の一般化: このモデルは、目に見えないカテゴリ (CWE) の脆弱性を 24.5% 削減するだけでなく、C/C++ から学んだセキュリティ原則を Python、Go、JavaScript などのさまざまな言語にうまく移植します。これは、TSP が単にパッチを記憶しているだけではなく、抽象的で言語に依存しないセキュリティ ロジックを内部化していることを示唆しています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.03489: Learn from Your Mistakes: Tree-like Self-Play for Secure Code LLMs

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.03489
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 2 Jun 2026]
Title: Learn from Your Mistakes: Tree-like Self-Play for Secure Code LLMs
Abstract: While Large Language Models (LLMs) excel in code generation, they remain prone to replicating subtle yet critical vulnerabilities endemic to their training data. Current alignment techniques, such as Supervised Fine-Tuning (SFT) and Reinforcement Learning (RL), typically apply coarse-grained optimization at the sequence level. This approach often fails to address the localized nature of security flaws, where a single incorrect token choice can compromise an entire program. To bridge this gap, we introduce Tree-like Self-Play (TSP), a framework that reframes secure code generation as a fine-grained sequential decision process. Unlike standard methods that blindly maximize likelihood, TSP constructs a decision tree where the model explores branching trajectories--generating both secure "golden paths" and vulnerable variants. By treating code generation as a self-play game, the model learns to strictly discriminate against its own localized errors. This provides a dense, on-policy learning signal that forces self-correction precisely at the critical decision nodes where vulnerabilities typically emerge. Our experiments demonstrate that TSP fundamentally enhances model reliability. In Python security benchmarks, TSP boosts CodeLlama-7B's pass rate (SPR@1) to 75.8%, significantly outperforming SFT (57.0%) and unstructured self-play baselines. Crucially, TSP induces robust out-of-distribution generalization: the model not only reduces vulnerabilities in unseen categories (CWEs) by 24.5% but also successfully transfers security principles learned from C/C++ to diverse languages, including Python, Go, and JavaScript. This suggests that TSP does not merely memorize patches, but internalizes abstract, language-agnostic security logic.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
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

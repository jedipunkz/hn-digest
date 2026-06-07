---
source: "https://arxiv.org/abs/2602.11964"
hn_url: "https://news.ycombinator.com/item?id=48430918"
title: "Gaia2: Benchmarking LLM Agents on Dynamic and Asynchronous Environments"
article_title: "[2602.11964] Gaia2: Benchmarking LLM Agents on Dynamic and Asynchronous Environments"
author: "Anon84"
captured_at: "2026-06-07T04:34:43Z"
capture_tool: "hn-digest"
hn_id: 48430918
score: 2
comments: 0
posted_at: "2026-06-07T01:36:17Z"
tags:
  - hacker-news
  - translated
---

# Gaia2: Benchmarking LLM Agents on Dynamic and Asynchronous Environments

- HN: [48430918](https://news.ycombinator.com/item?id=48430918)
- Source: [arxiv.org](https://arxiv.org/abs/2602.11964)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T01:36:17Z

## Translation

タイトル: Gaia2: 動的環境および非同期環境での LLM エージェントのベンチマーク
記事のタイトル: [2602.11964] Gaia2: 動的環境および非同期環境での LLM エージェントのベンチマーク
説明: arXiv 論文 2602.11964 の要約ページ: Gaia2: 動的および非同期環境での LLM エージェントのベンチマーク

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2602.11964
ヘルプ |高度な検索
検索を開く
行く
ナビゲーションメニューを開く
クイックリンク
コンピュータサイエンス > 人工知能
[2026 年 2 月 12 日に提出]
タイトル: Gaia2: 動的環境および非同期環境での LLM エージェントのベンチマーク
要約: 現実的な非同期環境で大規模言語モデル エージェントを評価するためのベンチマークである Gaia2 を紹介します。以前の静的評価や同期評価とは異なり、Gaia2 では環境がエージェントのアクションとは独立して進化するシナリオが導入されており、エージェントは時間的制約の下で動作し、ノイズの多い動的なイベントに適応し、あいまいさを解決し、他のエージェントと協力する必要があります。各シナリオは書き込みアクション検証器とペアになっているため、きめ細かいアクションレベルの評価が可能になり、Gaia2 を検証可能な報酬からの強化学習に直接使用できるようになります。最先端の独自モデルとオープンソース モデルを評価したところ、どのモデルも機能全体で優位に立つことはありませんでした。GPT-5 (高) は 42% pass@1 という最強の総合スコアに達しましたが、時間に敏感なタスクでは失敗しました。Claude-4 Sonnet は精度と速度をコストと引き換えにしました。Kimi-K2 は 21% pass@1 でオープンソース モデルの中でトップでした。これらの結果は、推論、効率、堅牢性の間の基本的なトレードオフを浮き彫りにし、「sim2real」のギャップを埋める際の課題を明らかにしています。 Gaia2 は、オープンソースの Agents Research Environments プラットフォームを使用した消費者環境に基づいて構築されており、簡単に拡張できるように設計されています。基本的な ARE フレームワークと一緒に Gaia2 をリリースすることで、開発、ベンチマーク、TR のための柔軟なインフラストラクチャをコミュニティに提供することを目指しています。

次世代の実用的なエージェント システムです。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
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

Abstract page for arXiv paper 2602.11964: Gaia2: Benchmarking LLM Agents on Dynamic and Asynchronous Environments

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2602.11964
Help | Advanced Search
open search
GO
open navigation menu
quick links
Computer Science > Artificial Intelligence
[Submitted on 12 Feb 2026]
Title: Gaia2: Benchmarking LLM Agents on Dynamic and Asynchronous Environments
Abstract: We introduce Gaia2, a benchmark for evaluating large language model agents in realistic, asynchronous environments. Unlike prior static or synchronous evaluations, Gaia2 introduces scenarios where environments evolve independently of agent actions, requiring agents to operate under temporal constraints, adapt to noisy and dynamic events, resolve ambiguity, and collaborate with other agents. Each scenario is paired with a write-action verifier, enabling fine-grained, action-level evaluation and making Gaia2 directly usable for reinforcement learning from verifiable rewards. Our evaluation of state-of-the-art proprietary and open-source models shows that no model dominates across capabilities: GPT-5 (high) reaches the strongest overall score of 42% pass@1 but fails on time-sensitive tasks, Claude-4 Sonnet trades accuracy and speed for cost, Kimi-K2 leads among open-source models with 21% pass@1. These results highlight fundamental trade-offs between reasoning, efficiency, robustness, and expose challenges in closing the "sim2real" gap. Gaia2 is built on a consumer environment with the open-source Agents Research Environments platform and designed to be easy to extend. By releasing Gaia2 alongside the foundational ARE framework, we aim to provide the community with a flexible infrastructure for developing, benchmarking, and training the next generation of practical agent systems.
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

---
source: "https://arxiv.org/abs/2604.27032"
hn_url: "https://news.ycombinator.com/item?id=48403393"
title: "LLM-Guided Runtime Parameter Optimization for Energy-Efficient Model Inference"
article_title: "[2604.27032] LLM-Guided Runtime Parameter Optimization for Energy-Efficient Model Inference"
author: "PaulHoule"
captured_at: "2026-06-04T21:38:20Z"
capture_tool: "hn-digest"
hn_id: 48403393
score: 2
comments: 0
posted_at: "2026-06-04T19:22:09Z"
tags:
  - hacker-news
  - translated
---

# LLM-Guided Runtime Parameter Optimization for Energy-Efficient Model Inference

- HN: [48403393](https://news.ycombinator.com/item?id=48403393)
- Source: [arxiv.org](https://arxiv.org/abs/2604.27032)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T19:22:09Z

## Translation

タイトル: エネルギー効率の高いモデル推論のための LLM ガイドによるランタイム パラメーターの最適化
記事のタイトル: [2604.27032] エネルギー効率の高いモデル推論のための LLM ガイド付きランタイム パラメーターの最適化
説明: arXiv 論文 2604.27032: エネルギー効率の高いモデル推論のための LLM ガイド付きランタイム パラメーターの最適化の要約ページ

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.27032
ヘルプ |高度な検索
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 4 月 29 日に提出]
タイトル: エネルギー効率の高いモデル推論のための LLM ガイドによるランタイム パラメーターの最適化
要約: 大規模言語モデル (LLM) は、現実世界の多くのワークフローに不可欠な部分になっています。ただし、LLM は大量のエネルギーを消費するため、これらのツールの需要が規模に応じて大きくなることが大きな問題となります。 LLM がさまざまなワークフローに統合されるにつれて、これらのツールの推論を実行するという課題に対処するさまざまなアプリケーションが登場しました。これにより、エネルギー消費を最小限に抑えるためにこれらのサービスの実行時パラメータ値を選択するという別の問題が生じます。多くの場合、これにはアプリケーションに関する深い知識や、最適な値を見つけるまでに数日かかる従来の最適化手法が必要です。この作業では、この問題を解決するために、LLM を利用したランタイム パラメーターの最適化を備えた人間参加型フローを作成しました。人間が作成した特定のフィードバックを促す方法を使用すると、チャットベースの LLM は、従来の検索方法よりも速く、エネルギー効率の高い推論パラメータを繰り返し見つけることができます。 LLM は、さまざまなハードウェア設定に合わせてソリューションを調整し、他のシステム制約を簡単に考慮することもできます。強化されたプロンプト テンプレートは、平均 5.2 個のプロンプトで収束したベースラインと比較して、平均 3.4 個のプロンプトでしきい値未満に収束することができ、一貫してより低いトークンあたりの最終エネルギーを達成しました。強化されたプロンプト テンプレートは、収束速度においても Sobol サンプリングを上回りました。
もっと学ぶために集中する
arXiv 発行の D

DataCite 経由の OI
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

Abstract page for arXiv paper 2604.27032: LLM-Guided Runtime Parameter Optimization for Energy-Efficient Model Inference

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.27032
Help | Advanced Search
Computer Science > Software Engineering
[Submitted on 29 Apr 2026]
Title: LLM-Guided Runtime Parameter Optimization for Energy-Efficient Model Inference
Abstract: Large Language Models (LLMs) have become an integral part of many real-world workflows. However, LLMs consume a lot of energy, which becomes a large concern in the scale of the demand for these tools. As LLMs become integrated into different workflows, different applications have arisen to deal with the challenge of running inference for these tools. This raises another issue of choosing the runtime parameter values for these services in order to minimize the energy consumption. Oftentimes this requires deep knowledge of the application or traditional optimization methods that can take days to find optimal values. In this work, we created a human-in-the-loop flow with LLM-assisted runtime parameter optimization in order to solve this issue. With human-created, specific feedback prompting methods, chat-based LLMs can iteratively find energy-efficient inference parameters faster than traditional search methods. LLMs can also tailor their solutions to different hardware setups and easily take into account other system constraints. The enhanced prompt template was able to converge below the threshold at an average of 3.4 prompts compared to the baseline, which converged in an average of 5.2 prompts, and consistently achieved lower final energy per token. The enhanced prompt template also outperformed Sobol sampling in convergence speed.
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

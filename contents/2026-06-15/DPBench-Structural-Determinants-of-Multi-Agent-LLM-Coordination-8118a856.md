---
source: "https://arxiv.org/abs/2602.13255"
hn_url: "https://news.ycombinator.com/item?id=48547462"
title: "DPBench: Structural Determinants of Multi-Agent LLM Coordination"
article_title: "[2602.13255] DPBench: Structural Determinants of Multi-Agent LLM Coordination Under Simultaneous Resource Contention"
author: "najmul-hasan"
captured_at: "2026-06-15T21:54:26Z"
capture_tool: "hn-digest"
hn_id: 48547462
score: 1
comments: 0
posted_at: "2026-06-15T21:48:03Z"
tags:
  - hacker-news
  - translated
---

# DPBench: Structural Determinants of Multi-Agent LLM Coordination

- HN: [48547462](https://news.ycombinator.com/item?id=48547462)
- Source: [arxiv.org](https://arxiv.org/abs/2602.13255)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T21:48:03Z

## Translation

タイトル: DPBench: マルチエージェント LLM 調整の構造的決定要因
記事のタイトル: [2602.13255] DPBench: 同時リソース競合下でのマルチエージェント LLM 調整の構造的決定要因
説明: arXiv 論文 2602.13255 の要約ページ: DPBench: 同時リソース競合下のマルチエージェント LLM 調整の構造的決定要因

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2602.13255
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 2 月 2 日に提出 ( v1 )、最終改訂日 2026 年 6 月 3 日 (このバージョン、v2)]
タイトル: DPBench: 同時リソース競合下でのマルチエージェント LLM 調整の構造的決定要因
要約: 大規模な言語モデルから構築されたマルチエージェント システムの調整を評価するためのベンチマークである DPBench を紹介します。既存のベンチマークは、固定プロトコルの下でタスクレベルの成功を測定します。調整が成功するか失敗する構造的条件はまったく特徴付けられていません。 DPBench は、食事の哲学者の問題を、行動プロトコル、コミュニケーション構造、グループ サイズがそれぞれ独立して変化する制御されたテストベッドに適応させます。 GPT-5.2、Claude Opus 4.5、Grok 4.1、Gemini 2.5 Flash、Llama 4 Maverick、および均一ランダムのベースラインの 6 つのエージェントを評価します。デフォルトのプロンプトを使用した N=5 の同時アクションでは、デッドロックの範囲は GPT-5.2 の 25.0% (95% Wilson CI [11.2, 46.9]) から Gemini 2.5 フラッシュの 90.0% [74.4, 96.5] です。連続アクションは 6 つのうち 4 つで解決されます。モデルを Gemini 2.5 フラッシュで固定したまま、3 つのプロトコル変数によってデッドロックが 90% から CI ゼロ以内に達します。つまり、3 ラウンドのプリコミットメント通信 (0.0% 対 1 ラウンド 86.7%)、古典的な同時実行プリミティブをエンコードするプロンプト (最小プロンプトの 100% に対して、リソースの順序付けと対称性の破壊には 0.0%)、またはグループを N=5 から 2 倍にします。 N=10 (90.0% ～ 10.0%)。単一ラウンドのメッセージングと過去のタイムステップのメモリでは、実行したサンプル サイズではレートは変わりません。かどうか

同じモデルの座標またはデッドロックは、モデルの機能ではなくプロトコルによって決まります。
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

Abstract page for arXiv paper 2602.13255: DPBench: Structural Determinants of Multi-Agent LLM Coordination Under Simultaneous Resource Contention

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2602.13255
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 2 Feb 2026 ( v1 ), last revised 3 Jun 2026 (this version, v2)]
Title: DPBench: Structural Determinants of Multi-Agent LLM Coordination Under Simultaneous Resource Contention
Abstract: We present DPBench, a benchmark for evaluating coordination in multi-agent systems built from large language models. Existing benchmarks measure task-level success under a fixed protocol; the structural conditions under which coordination succeeds or fails at all have not been characterised. DPBench adapts the Dining Philosophers problem into a controlled testbed where the action protocol, the communication structure, and the group size each vary independently. We evaluate six agents: GPT-5.2, Claude Opus 4.5, Grok 4.1, Gemini 2.5 Flash, Llama 4 Maverick, and a uniform-random baseline. Under simultaneous action at N=5 with the default prompt, deadlock ranges from 25.0% (95% Wilson CI [11.2, 46.9]) for GPT-5.2 to 90.0% [74.4, 96.5] for Gemini 2.5 Flash; sequential action is solved by four of the six. Holding the model fixed at Gemini 2.5 Flash, three protocol variables drive deadlock from 90% to within CI of zero: three rounds of pre-commitment communication (0.0% vs. single-round 86.7%), a prompt encoding a classical concurrency primitive (0.0% for resource-ordering and symmetry-breaking, against 100% for the minimal prompt), or doubling the group from N=5 to N=10 (90.0% to 10.0%). Single-round messaging and memory of past timesteps do not change the rate at the sample size we ran. Whether the same model coordinates or deadlocks is determined by the protocol, not by the model's capability.
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

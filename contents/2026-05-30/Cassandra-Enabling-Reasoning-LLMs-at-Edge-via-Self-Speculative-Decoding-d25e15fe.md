---
source: "https://arxiv.org/abs/2605.26558"
hn_url: "https://news.ycombinator.com/item?id=48323047"
title: "Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding"
article_title: "[2605.26558] Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding"
author: "chrsw"
captured_at: "2026-05-30T11:47:26Z"
capture_tool: "hn-digest"
hn_id: 48323047
score: 4
comments: 0
posted_at: "2026-05-29T13:45:30Z"
tags:
  - hacker-news
  - translated
---

# Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding

- HN: [48323047](https://news.ycombinator.com/item?id=48323047)
- Source: [arxiv.org](https://arxiv.org/abs/2605.26558)
- Score: 4
- Comments: 0
- Posted: 2026-05-29T13:45:30Z

## Translation

タイトル: Cassandra: 自己投機的デコーディングによるエッジでの推論 LLM の有効化
記事のタイトル: [2605.26558] Cassandra: 自己投機的デコーディングによるエッジでの推論 LLM の有効化
説明: arXiv 論文 2605.26558 の要約ページ: Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.26558
ヘルプ |高度な検索
コンピュータ サイエンス > ハードウェア アーキテクチャ
[2026 年 5 月 26 日提出]
タイトル: Cassandra: 自己投機的デコーディングによるエッジでの推論 LLM の有効化
要約: 投機的デコードは、大規模言語モデル (LLM) を高速化するための有望なロスレス アプローチとして登場しました。推論 LLM はデコード段階のオーバーヘッドにますます悩まされ、近似ベースの手法では精度が低下するため、効率的な推論にはロスレスの投機的デコードが不可欠になっています。ただし、既存の方法では、追加のトレーニングなしで強力な低バッチ パフォーマンスを実現するのは依然として困難であり、消費者向けデバイスでの実際の展開は制限されています。この課題に対処するために、アルゴリズムとハードウェアが共同設計され、低バッチ シナリオ向けに最適化された自己投機的デコード フレームワークである Cassandra を提案します。 Cassandra は、きめ細かいデータ選択を通じて、トレーニング不要の高性能のドラフト モデルを構築します。最適化された枝刈りと仮数切り捨てを使用して、モデルの重みと Key-Value (KV) キャッシュの両方で最も顕著な値を特定し、完全精度の並列検証の前に候補トークンを迅速に生成できるようにします。レイヤースキップや構造化された KV 圧縮に基づく従来の自己投機的デコード方法とは異なり、Cassandra は大幅に高い効率を達成します。 Cassandra 表現と標準浮動小数点形式の間の形式変換のオーバーヘッドをさらに削減するために、商用 GPU および NPU とのシームレスな統合のために設計された軽量のエンコーダ/デコーダ ハードウェア モジュールも導入しました。実験結果によると、カサンドラは

追加のトレーニングなしで、BF16 ベースラインと比較して最大 2.41 倍の高速化を達成します。さらに、NVIDIA GeForce RTX 4090 上で実行されている Llama 3 8B では、Cassandra は、最先端の投機的デコード方式である Eagle-3 と比較して、同じメモリ バジェットの下で 1.81 倍多くのトークンを生成します。
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

Abstract page for arXiv paper 2605.26558: Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.26558
Help | Advanced Search
Computer Science > Hardware Architecture
[Submitted on 26 May 2026]
Title: Cassandra: Enabling Reasoning LLMs at Edge via Self-Speculative Decoding
Abstract: Speculative decoding has emerged as a promising lossless approach for accelerating Large Language Models (LLMs). As reasoning LLMs increasingly suffer from decode-stage overhead and approximation-based methods degrade accuracy, lossless speculative decoding has become essential for efficient inference. However, existing methods still struggle to deliver strong low-batch performance without additional training, limiting practical deployment on consumer devices. To address this challenge, we propose Cassandra, an algorithm-hardware co-designed self-speculative decoding framework optimized for low-batch scenarios. Cassandra constructs a high-performance, training-free draft model through fine-grained data selection. Using optimized pruning and mantissa truncation, it identifies the most salient values in both model weights and the Key-Value (KV) cache, enabling rapid candidate token generation before full-precision parallel verification. Unlike prior self-speculative decoding methods based on layer skipping or structured KV compression, Cassandra achieves significantly higher efficiency. To further reduce the overhead of format conversion between Cassandra representations and standard floating-point formats, we also introduce a lightweight encoder-decoder hardware module designed for seamless integration with commercial GPUs and NPUs. Experimental results show that Cassandra achieves up to 2.41x speedup over the BF16 baseline without additional training. Furthermore, on Llama 3 8B running on an NVIDIA GeForce RTX 4090, Cassandra generates 1.81x more tokens under the same memory budget compared to Eagle-3, a state-of-the-art speculative decoding method.
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

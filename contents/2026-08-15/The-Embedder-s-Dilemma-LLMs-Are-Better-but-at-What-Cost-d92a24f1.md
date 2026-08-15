---
source: "https://arxiv.org/abs/2608.12875"
hn_url: "https://news.ycombinator.com/item?id=49308102"
title: "The Embedder's Dilemma: LLMs Are Better, but at What Cost?"
article_title: "[2608.12875] The Embedder's Dilemma: LLMs Are Better, but at What Cost?"
author: "sbulaev"
captured_at: "2026-08-15T06:20:27Z"
capture_tool: "hn-digest"
hn_id: 49308102
score: 1
comments: 0
posted_at: "2026-08-15T06:07:06Z"
tags:
  - hacker-news
  - translated
---

# The Embedder's Dilemma: LLMs Are Better, but at What Cost?

- HN: [49308102](https://news.ycombinator.com/item?id=49308102)
- Source: [arxiv.org](https://arxiv.org/abs/2608.12875)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T06:07:06Z

## Translation

タイトル: 埋め込み者のジレンマ: LLM は優れていますが、その代償は何ですか?
記事のタイトル: [2608.12875] 埋め込み者のジレンマ: LLM の方が優れていますが、その代償は何ですか?
説明: arXiv 論文 2608.12875 の要約ページ: The Embedder's Dilemma: LLMs Are Better, but at What Cost?

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 8 月 13 日に提出]
タイトル: 埋め込み者のジレンマ: LLM は優れていますが、その代償は何ですか?
要約: テキスト埋め込みパイプラインを大規模な言語モデルに置き換えるべきですか?私たちは、分類、意味論的テキスト類似性 (STS)、クラスタリング、ペア分類、検索にわたる 37 のタスクについて、6 ファミリにわたる 10 個の LLM と 26 の埋め込みモデル (118M ～ 14B パラメータ) を制御されたコストを意識した比較で答えます。合計すると、2 つのパラダイムは事実上結びついています。最良の LLM (Gemini 3.1 Pro、77.6) と最良の埋め込みモデル (77.2) は 0.4 ポイントの差があります。それらの強みはタスクによって異なります。LLM は推論重視の検索をリードし、埋め込みモデルは分類をリードし、2 つの一致はクラスタリング、STS、およびペア分類をリードします。その同等に達するには費用がかかります。 LLM のコストは、同等の品質の埋め込みモデルよりも最大 1,431 倍高く (ベンチマーク パスあたり 154 米ドル対 0.11 米ドル)、同じ GPU 上でテストされたオープン LLM はトークンの処理速度が 2.5 ～ 736 倍遅くなります。推論トークンは、LLM 推論コストの 28 ～ 81% を占めます。推論予算が低いため、アブレーションのほとんどのモデルの検索品質が維持または向上します。パレート フロンティアには、主要な埋め込みモデルと 1 つの LLM、Gemini 3.1 Pro が含まれています。これらの結果は、類似性、分類、クラスタリングには埋め込みモデルを使用し、推論集中型の検索には LLM を予約するという分業を裏付けています。私たちのコード、データセット、結果は、この https URL で公開されています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
コード、データ、メディア アソシエイト

この記事では
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.12875: The Embedder's Dilemma: LLMs Are Better, but at What Cost?

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 13 Aug 2026]
Title: The Embedder's Dilemma: LLMs Are Better, but at What Cost?
Abstract: Should you replace your text-embedding pipeline with a large language model? We answer this with a controlled, cost-aware comparison of ten LLMs across six families and 26 embedding models (118M to 14B parameters) on 37 tasks spanning classification, semantic textual similarity (STS), clustering, pair classification, and retrieval. In aggregate the two paradigms are effectively tied: the best LLM (Gemini 3.1 Pro, 77.6) and the best embedding model (77.2) differ by 0.4 points. Their strengths differ by task: LLMs lead on reasoning-heavy retrieval, embedding models lead on classification, and the two match on clustering, STS, and pair classification. Reaching that parity is expensive. An LLM costs up to 1,431x more than an embedding model of comparable quality (USD 154 vs. USD 0.11 per benchmark pass), and the open LLMs tested process tokens 2.5 to 736x more slowly on the same GPU. Reasoning tokens account for 28 to 81% of LLM inference cost; lower reasoning budgets preserve or improve retrieval quality for most models in our ablation. The Pareto frontier contains the leading embedding models and one LLM, Gemini 3.1 Pro. These results support a division of labour: use embedding models for similarity, classification, and clustering, and reserve LLMs for reasoning-intensive retrieval. Our code, datasets, and results are publicly available at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

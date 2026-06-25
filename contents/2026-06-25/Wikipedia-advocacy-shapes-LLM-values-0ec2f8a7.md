---
source: "https://arxiv.org/abs/2606.24890"
hn_url: "https://news.ycombinator.com/item?id=48669764"
title: "Wikipedia advocacy shapes LLM values"
article_title: "[2606.24890] Small edits, large models: How Wikipedia advocacy shapes LLM values"
author: "50kIters"
captured_at: "2026-06-25T07:07:34Z"
capture_tool: "hn-digest"
hn_id: 48669764
score: 2
comments: 0
posted_at: "2026-06-25T06:31:49Z"
tags:
  - hacker-news
  - translated
---

# Wikipedia advocacy shapes LLM values

- HN: [48669764](https://news.ycombinator.com/item?id=48669764)
- Source: [arxiv.org](https://arxiv.org/abs/2606.24890)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T06:31:49Z

## Translation

タイトル: ウィキペディアの擁護が LLM の価値を形作る
記事のタイトル: [2606.24890] 小さな編集、大きなモデル: ウィキペディアの擁護が LLM 値をどのように形作るか
説明: arXiv 論文 2606.24890 の要約ページ: 小さな編集、大きなモデル: ウィキペディアの擁護が LLM 値をどのように形成するか

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.24890
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 4 月 30 日に提出]
タイトル: 小さな編集、大きなモデル: ウィキペディアの擁護が LLM の価値をどのように形成するか
要約: 少人数のボランティア グループが、ウィキペディアを編集するだけで、AI システムが動物福祉について議論する方法を形作ることができますか?私たちは彼らができることを示します。 Wikipedia は、ほぼすべての主要な言語モデルのトレーニング データセットに含まれており、Web クロールされたテキストよりも重み付けされています。関連する記事に動物福祉に関する情報源を追加する活動家グループであるプロアニマル ウィキペディアンズ (PAW) は、115 ページにわたって 125 回の編集を加えました。勾配ベースのデータ属性 (Bergson; MAGIC) を使用して、これらの編集が言語モデルの動作にどのような影響を与えるかを追跡しました。 Llama 3.1 8B の TrackStar 検索アトリビューションでは、PAW で編集されたセクションが、動物福祉に関するクエリについて最も帰属が高い文書の 68 パーセントを占めていることがわかりました (p < 0.0001) が、同じ企業に関する無関係なクエリではわずか 52 パーセント (p = 0.53) でした。このモデルは、PAW のコンテンツを、エンティティ全般ではなく、特に動物福祉のトピックに関連付けています。ラマ-3.2-1B に対する MAGIC の反事実的影響推定は、5 つのランダムなトレーニング順序シードで実行され、同じ状況をさらに鮮明に示しました。どのシードでも、動物福祉クエリーで最も影響力のある文書のトップ 10 はすべて PAW 編集でした (10 個中 10 個、5 個中 5 個のシード) が、一般クエリでは同じトップ 10 が偶然に存在しました (10 個中 4 ～ 6 個)。動物福祉クエリに対する平均 PAW 影響は、すべてのシードで p < 0.0001 で平均対照影響を上回りました。これは、一般クエリよりも 6 ～ 30 倍大きい効果です。リーブサブセットアウト v

分析では、10 回の実行すべてで Spearman の rho = 1.00 が得られました。 PAW コンテンツとコントロール コンテンツの別々のモデルを微調整したところ、各モデルはトレーニングされたテキストのタイプに特化してより良いパフォーマンスを示しました。PAW でトレーニングされたモデルは、動物福祉テキストの混乱度を 12.4 から 8.4 に削減しましたが、コントロールでトレーニングされたモデルは、コントロール テキストの混乱度を 16.1 から 11.4 に削減しました。したがって、小規模で調整されたウィキペディア編集キャンペーンは、編集で扱われるトピックを言語モデルがどのように処理するかを測定的に決定します。
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

Abstract page for arXiv paper 2606.24890: Small edits, large models: How Wikipedia advocacy shapes LLM values

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.24890
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 30 Apr 2026]
Title: Small edits, large models: How Wikipedia advocacy shapes LLM values
Abstract: Can a small group of volunteers shape how AI systems discuss animal welfare, just by editing Wikipedia? We show that they can. Wikipedia appears in nearly every major language model training dataset and is weighted more heavily than web-crawled text. The Pro-Animal Wikipedians (PAW), a group of advocates who add sourced animal welfare content to relevant articles, have made 125 edits across 115 pages. Using gradient-based data attribution (Bergson; MAGIC), we traced how these edits influence language model behavior. TrackStar retrieval attribution on Llama 3.1 8B found that PAW-edited sections made up 68 percent of the highest-attributed documents for animal welfare queries (p < 0.0001) but only 52 percent for unrelated queries about the same companies (p = 0.53): the model links PAW content specifically to animal welfare topics, not to the entities in general. MAGIC counterfactual influence estimation on Llama-3.2-1B, run across five random training-order seeds, gave the same picture even more sharply: in every seed, the top-10 most influential documents on animal welfare queries were all PAW edits (10 of 10, 5 of 5 seeds), while on general queries the same top-10 sat at chance (4 to 6 of 10). Mean PAW influence exceeded mean control influence on animal welfare queries with p < 0.0001 in every seed, an effect 6 to 30 times larger than on general queries. Leave-subset-out validation gave Spearman rho = 1.00 for all 10 runs. When we fine-tuned separate models on PAW content versus control content, each model performed better specifically on the type of text it was trained on: the PAW-trained model cut perplexity on animal welfare text from 12.4 to 8.4, while the control-trained model cut perplexity on control text from 16.1 to 11.4. A small, coordinated Wikipedia editing campaign therefore measurably shapes how language models handle the topics those edits address.
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

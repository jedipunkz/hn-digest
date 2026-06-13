---
source: "https://arxiv.org/abs/2606.06622"
hn_url: "https://news.ycombinator.com/item?id=48510690"
title: "UnpredictaBench: A Benchmark for Evaluating Distributional Randomness in LLMs"
article_title: "[2606.06622] UnpredictaBench: A Benchmark for Evaluating Distributional Randomness in LLMs"
author: "matt_d"
captured_at: "2026-06-13T01:01:07Z"
capture_tool: "hn-digest"
hn_id: 48510690
score: 1
comments: 0
posted_at: "2026-06-12T23:41:49Z"
tags:
  - hacker-news
  - translated
---

# UnpredictaBench: A Benchmark for Evaluating Distributional Randomness in LLMs

- HN: [48510690](https://news.ycombinator.com/item?id=48510690)
- Source: [arxiv.org](https://arxiv.org/abs/2606.06622)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T23:41:49Z

## Translation

タイトル: UnpredictaBench: LLM の分布ランダム性を評価するためのベンチマーク
記事のタイトル: [2606.06622] UnpredictaBench: LLM の分布ランダム性を評価するためのベンチマーク
説明: arXiv 論文 2606.06622 の要約ページ: UnpredictaBench: LLM の分布ランダム性を評価するためのベンチマーク

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.06622
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 6 月 4 日に提出 ( v1 )、最終改訂日 2026 年 6 月 8 日 (このバージョン、v2)]
タイトル: UnpredictaBench: LLM の分布ランダム性を評価するためのベンチマーク
要約: 大規模言語モデル (LLM) が真の基礎となる分布を捕捉する能力をテストする評価である UnpredictaBench を紹介します。 LLM が他のエンティティ (経済シミュレーションにおける人間など) の代替として使用されることが増えているため、多くのモデルが単一のもっともらしい答えに向かって崩壊する傾向は、実際のシステムの予測不可能性を捉えることができないことを意味します。出力の多様性を改善するための最近の取り組みでは、この設定には不十分です。シミュレーションには、単に出力を変化させるだけでなく、ターゲットの分布に合わせて調整されたサンプルが必要です。 UnpredictaBench は、こ​​の問題の単純化された、しかし根本的なバージョンを分離します。つまり、標準統計分布、確率的プログラムによって誘発される分布、ランダム プロセスを記述する自然言語シナリオを含む、個々のターゲット分布から結果をサンプリングします。このような 448 の問題を、コルモゴロフ-スミルノフ統計検定によってモデルがどの程度適切にブラック ボックス ターゲット分布を出力するかを定量化する汎用評価指標である KS@N とともに紹介します。これは、グラウンドトゥルース サンプルに対してサイズ N のモデル サンプルを拒否できない率であり、N が大きいほど、難易度が高くなります。オープン モデルと独自モデルでテストしたところ、分散機能に大きな広がりがあることがわかりました。たとえば、モードのとき

サイズ 100 (KS@100、当社の標準メトリクス) のサンプルを生成し、スコアの範囲は 0 近くから 20% 以上までです。 KS@100 で 40% 以上を達成できるモデルはなく、機能としての分布サンプリングで大きなヘッドルームを示しています。推論を追加するとスコアが多少向上する可能性がありますが、この問題に対する即時の解決策は見つかりません。 UnpredictaBench は、単純な分布シミュレーションでさえ依然として困難であり、複雑なシステムの代用として LLM を使用するために必要な第一歩であることを示しています。
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

Abstract page for arXiv paper 2606.06622: UnpredictaBench: A Benchmark for Evaluating Distributional Randomness in LLMs

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.06622
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 4 Jun 2026 ( v1 ), last revised 8 Jun 2026 (this version, v2)]
Title: UnpredictaBench: A Benchmark for Evaluating Distributional Randomness in LLMs
Abstract: We introduce UnpredictaBench, an evaluation that tests the ability of large language models (LLMs) to capture true underlying distributions. As LLMs are increasingly used as substitutes for other entities (e.g., for humans in economic simulations), the tendency of many models to collapse towards a single plausible answer means a failure to capture the unpredictability of real systems. Recent work on improving output diversity is insufficient for this setting: simulation requires samples that are calibrated to a target distribution, not merely varied outputs. UnpredictaBench isolates a simplified but fundamental version of this problem: sampling outcomes from individual target distributions, including canonical statistical distributions, distributions induced by stochastic programs, and natural-language scenarios that describe random processes. We introduce 448 such problems together with KS@N, a general-purpose evaluation metric that quantifies how well a model outputs approximate black-box target distributions via the Kolmogorov-Smirnov statistical test. This is the rate at which we fail to reject model samples of size N against ground-truth samples, with larger N indicating greater difficulty. Tested across open and proprietary models, we find a large spread in distributional capabilities. For instance, when models generate samples of size 100 (KS@100, our standard metric), scores range from near 0 to over 20%. No model is able to achieve over 40% at KS@100, showing significant headroom in distributional sampling as a capability. Although adding reasoning can somewhat increase scores, we find no immediate solution for this issue. UnpredictaBench shows that even simple distributional simulation remains challenging, making it a necessary first step toward using LLMs as stand-ins for complex systems.
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

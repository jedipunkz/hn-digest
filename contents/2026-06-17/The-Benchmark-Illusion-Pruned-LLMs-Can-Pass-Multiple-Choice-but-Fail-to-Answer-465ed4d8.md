---
source: "https://arxiv.org/abs/2606.17609"
hn_url: "https://news.ycombinator.com/item?id=48564787"
title: "The Benchmark Illusion: Pruned LLMs Can Pass Multiple Choice but Fail to Answer"
article_title: "[2606.17609] The Benchmark Illusion: Pruned LLMs Can Pass Multiple Choice but Fail to Answer"
author: "ilreb"
captured_at: "2026-06-17T04:36:04Z"
capture_tool: "hn-digest"
hn_id: 48564787
score: 2
comments: 0
posted_at: "2026-06-17T01:53:22Z"
tags:
  - hacker-news
  - translated
---

# The Benchmark Illusion: Pruned LLMs Can Pass Multiple Choice but Fail to Answer

- HN: [48564787](https://news.ycombinator.com/item?id=48564787)
- Source: [arxiv.org](https://arxiv.org/abs/2606.17609)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T01:53:22Z

## Translation

タイトル: ベンチマークの幻想: プルーニングされた LLM は多肢選択には合格するが回答に失敗する
記事のタイトル: [2606.17609] ベンチマークの錯覚: プルーニングされた LLM は多肢選択には合格するが回答に失敗する
説明: arXiv 論文 2606.17609 の要約ページ: The Benchmark Illusion: プルーニングされた LLM は複数の選択肢に合格するが回答に失敗する

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.17609
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 6 月 16 日に提出]
タイトル: ベンチマークの幻想: プルーニングされた LLM は多肢選択には合格するが回答に失敗する
要約: 大規模な言語モデルを圧縮すると、メモリ使用量と推論コストが削減されますが、標準のベンチマークでは見逃される障害が発生する可能性もあります。枝刈りされたモデルは、多肢選択評価では依然として良好なパフォーマンスを示しますが、オープン世代では同じ質問に答えることができない可能性があります。枝刈りがどのような変化をもたらすのかを尋ねます。それは正しい答えを消去するのでしょうか、それとも、その答えをトップの出力として生成するのが難しくなりますか?
私たちはこの質問を多言語質問応答で研究し、枝刈りの前後で同じ質問を追跡します。ベンチマークの錯覚が見つかりました。高スパース性枝刈り (特に Wanda) では、モデルは多くの場合、多肢選択スコアリングでは正しい答えを選択しながら、貪欲なオープン生成に失敗します。これらの認識のみのエラーでは、通常、答えは消えませんが、降格されます。多くの場合、ビーム検索、サンプリング、またはコンテキスト内の 1 つの例で再表示されます。全体として、複数選択ベンチマークは圧縮 LLM の使いやすさを過大評価する可能性があり、評価の盲点が生じます。圧縮モデルは、認識できるものだけでなく、生成できるものについてもテストする必要があります。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arX を開発および共有できるようにするフレームワークです。

iv の機能は当社の Web サイトに直接掲載されています。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.17609: The Benchmark Illusion: Pruned LLMs Can Pass Multiple Choice but Fail to Answer

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.17609
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 16 Jun 2026]
Title: The Benchmark Illusion: Pruned LLMs Can Pass Multiple Choice but Fail to Answer
Abstract: Compressing large language models reduces memory use and inference cost, but it can also create failures that standard benchmarks miss. A pruned model may still perform well on multiple-choice evaluations, yet fail to answer the same question in open generation. We ask what pruning changes: does it erase the correct answer, or does it make the answer harder to produce as the top output?
We study this question with multilingual question answering, tracking the same questions before and after pruning. We find a benchmark illusion. Under high-sparsity pruning, especially Wanda, models often fail in greedy open generation while still selecting the correct answer under multiple-choice scoring. In these recognition-only errors, the answer is usually not gone, but demoted: it often reappears with beam search, sampling, or one in-context example. Overall, multiple-choice benchmarks can overstate the usability of compressed LLMs, creating an evaluation blind spot. Compressed models should be tested on what they can produce, not only on what they can recognize.
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

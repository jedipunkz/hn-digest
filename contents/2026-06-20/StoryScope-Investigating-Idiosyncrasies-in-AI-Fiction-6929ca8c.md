---
source: "https://arxiv.org/abs/2604.03136"
hn_url: "https://news.ycombinator.com/item?id=48612708"
title: "StoryScope: Investigating Idiosyncrasies in AI Fiction"
article_title: "[2604.03136] StoryScope: Investigating idiosyncrasies in AI fiction"
author: "amai"
captured_at: "2026-06-20T21:33:40Z"
capture_tool: "hn-digest"
hn_id: 48612708
score: 2
comments: 0
posted_at: "2026-06-20T20:29:46Z"
tags:
  - hacker-news
  - translated
---

# StoryScope: Investigating Idiosyncrasies in AI Fiction

- HN: [48612708](https://news.ycombinator.com/item?id=48612708)
- Source: [arxiv.org](https://arxiv.org/abs/2604.03136)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T20:29:46Z

## Translation

タイトル: StoryScope: AI フィクションにおける特異性の調査
記事のタイトル: [2604.03136] StoryScope: AI フィクションの特異性の調査
説明: arXiv 論文 2604.03136 の要約ページ: StoryScope: AI フィクションにおける特異性の調査

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.03136
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 4 月 3 日に提出 ( v1 )、最終改訂日 2026 年 4 月 13 日 (このバージョン、v4)]
タイトル: StoryScope: AI フィクションにおける特異性の調査
要約: AI によって生成されたフィクションがますます普及するにつれて、著作者と独創性の問題が書かれた作品の評価方法の中心になりつつあります。この分野の既存の研究のほとんどは、AI の文章の表面レベルの特徴を特定することに焦点を当てていますが、私たちは代わりに、登場人物の主体性や時系列的な不連続性などの談話レベルの物語の選択に焦点を当て、文体シグナルに依存せずに AI が生成した物語を人間の物語と区別できるかどうかを問います。私たちは StoryScope を提案します。これは、10 次元にわたる談話レベルの物語の特徴からなる、きめの細かい解釈可能な特徴空間を自動的に誘導するパイプラインです。 StoryScope を、人間の作成者と 5 人の LLM によってそれぞれ書かれた 10,272 個の執筆プロンプトの並列コーパスに適用すると、61,608 個のストーリー、各約 5,000 ワード、およびストーリーごとに 304 個の抽出された特徴が得られます。ナラティブ機能だけでも、人間対 AI の検出で 93.2% のマクロ F1 を達成し、6 方向の著者帰属で 68.4% のマクロ F1 を達成し、文体の手がかりを含むモデルのパフォーマンスの 97% 以上を維持します。 30 の核となる物語機能からなるコンパクトなセットは、このシグナルの多くを捉えています。AI の物語はテーマを過剰に説明し、整然とした単線的なプロットを好む一方、人間の物語は主人公の選択を道徳的により曖昧なものとして枠付けし、時間的な複雑さを増しています。モデルごとのフィンガープリント機能により、6 方向のアトリビューションが可能になります。

たとえば、クロードは特に平坦なイベント エスカレーションを生成し、夢のシーケンスでは GPT のオーバーインデックスが発生し、ジェミニはデフォルトで外部文字の記述を行います。 AI が生成したストーリーは物語空間の共有領域に集中しているのに対し、人間が作成したストーリーはより多様性を示していることがわかりました。より広範に、これらの結果は、文体だけでなく、根底にある物語構造の違いを利用して、人間が書いたオリジナル作品と AI が生成したフィクションを区別できることを示唆しています。
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

Abstract page for arXiv paper 2604.03136: StoryScope: Investigating idiosyncrasies in AI fiction

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.03136
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 3 Apr 2026 ( v1 ), last revised 13 Apr 2026 (this version, v4)]
Title: StoryScope: Investigating idiosyncrasies in AI fiction
Abstract: As AI-generated fiction becomes increasingly prevalent, questions of authorship and originality are becoming central to how written work is evaluated. While most existing work in this space focuses on identifying surface-level signatures of AI writing, we ask instead whether AI-generated stories can be distinguished from human ones without relying on stylistic signals, focusing on discourse-level narrative choices such as character agency and chronological discontinuity. We propose StoryScope, a pipeline that automatically induces a fine-grained, interpretable feature space of discourse-level narrative features across 10 dimensions. We apply StoryScope to a parallel corpus of 10,272 writing prompts, each written by a human author and five LLMs, yielding 61,608 stories, each ~5,000 words, and 304 extracted features per story. Narrative features alone achieve 93.2% macro-F1 for human vs. AI detection and 68.4% macro-F1 for six-way authorship attribution, retaining over 97% of the performance of models that include stylistic cues. A compact set of 30 core narrative features captures much of this signal: AI stories over-explain themes and favor tidy, single-track plots while human stories frame protagonist' choices as more morally ambiguous and have increased temporal complexity. Per-model fingerprint features enable six-way attribution: for example, Claude produces notably flat event escalation, GPT over-indexes on dream sequences, and Gemini defaults to external character description. We find that AI-generated stories cluster in a shared region of narrative space, while human-authored stories exhibit greater diversity. More broadly, these results suggest that differences in underlying narrative construction, not just writing style, can be used to separate human-written original works from AI-generated fiction.
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

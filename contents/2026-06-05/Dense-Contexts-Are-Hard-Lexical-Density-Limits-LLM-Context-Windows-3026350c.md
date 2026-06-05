---
source: "https://arxiv.org/abs/2606.06203"
hn_url: "https://news.ycombinator.com/item?id=48409066"
title: "Dense Contexts Are Hard: Lexical Density Limits LLM Context Windows"
article_title: "[2606.06203] Dense Contexts Are Hard Contexts: Lexical Density Limits Effective Context in LLMs"
author: "sbulaev"
captured_at: "2026-06-05T07:36:47Z"
capture_tool: "hn-digest"
hn_id: 48409066
score: 1
comments: 0
posted_at: "2026-06-05T07:11:40Z"
tags:
  - hacker-news
  - translated
---

# Dense Contexts Are Hard: Lexical Density Limits LLM Context Windows

- HN: [48409066](https://news.ycombinator.com/item?id=48409066)
- Source: [arxiv.org](https://arxiv.org/abs/2606.06203)
- Score: 1
- Comments: 0
- Posted: 2026-06-05T07:11:40Z

## Translation

タイトル: 高密度コンテキストは難しい: 字句密度による LLM コンテキスト ウィンドウの制限
記事のタイトル: [2606.06203] 密なコンテキストはハード コンテキスト: 語彙密度が LLM の有効なコンテキストを制限する
説明: arXiv 論文 2606.06203 の要約ページ: 密なコンテキストはハード コンテキストである: 字句密度が LLM の有効なコンテキストを制限する

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.06203
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 6 月 4 日に提出]
タイトル: 密なコンテキストはハードコンテキストである: 語彙密度は LLM の有効なコンテキストを制限する
要約: 入力長と関連情報の位置は、LLM ロングコンテキストのパフォーマンス低下の主な原因として広く挙げられています。ここでは、LLM の有効なコンテキスト ウィンドウを体系的に減少させる、あまり見落とされがちな 3 番目の要因として、語彙密度 (コンテキストによって個別の情報が導入される割合) を研究します。同じ長さ (約 12,000 トークン) と制御されたニードル位置を持つ 3 つの「ファインド・ザ・ニードル」スタイルのベンチマークを使用して、オープンウェイト LLM (9B-685B) に対する語彙密度の影響を定量化しますが、情報密度は増加します。高密度のベンチマークでは急激なパフォーマンスの低下が観察されます。つまり、疎なコンテキストではほぼ完璧なモデルが、より密なコンテキストでは取得スコアが 60% を下回ります。タスクタイプの交絡を排除するために、他のすべてのプロパティを変更せずに、各ベンチマーク内の密度を変更して制御します。一般に、密度を下げると、特に劣化が見られる高密度領域でパフォーマンスが回復します。これらの結果は、有効なコンテキスト容量が語彙密度の関数であり、コンパクトで情報豊富な入力で動作する現実世界の LLM システムに直接影響することを示しています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティとの実験的プロジェクト

協力者
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.06203: Dense Contexts Are Hard Contexts: Lexical Density Limits Effective Context in LLMs

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.06203
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 4 Jun 2026]
Title: Dense Contexts Are Hard Contexts: Lexical Density Limits Effective Context in LLMs
Abstract: Input length and the position of relevant information are widely cited as the primary causes of degraded LLM long-context performance. Here, we study lexical density -- the rate at which a context introduces distinct information -- as a third, largely overlooked factor that systematically reduces the effective context window of LLMs. We quantify the impact of lexical density on open-weight LLMs (9B-685B) using three "find-the-needle" style benchmarks with identical length (~12k tokens) and controlled needle position, but increasing density of information. We observe a sharp performance collapse in higher-density benchmarks: models that are near-perfect in sparse contexts drop below 60% retrieval score on denser ones. To rule out task-type confounds, we vary and control the density within each benchmark while keeping all other properties unchanged. Reducing density generally restores performance, especially in the high-density regimes where degradation appears. These results show that effective context capacity is a function of lexical density, with direct implications for real-world LLM systems operating on compact, information-rich inputs.
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

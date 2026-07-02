---
source: "https://arxiv.org/abs/2601.13082"
hn_url: "https://news.ycombinator.com/item?id=48765725"
title: "Manipulating Headlines in LLM-Driven Algorithmic Trading"
article_title: "[2601.13082] Adversarial News and Lost Profits: Manipulating Headlines in LLM-Driven Algorithmic Trading"
author: "m-hodges"
captured_at: "2026-07-02T19:09:50Z"
capture_tool: "hn-digest"
hn_id: 48765725
score: 2
comments: 0
posted_at: "2026-07-02T18:47:11Z"
tags:
  - hacker-news
  - translated
---

# Manipulating Headlines in LLM-Driven Algorithmic Trading

- HN: [48765725](https://news.ycombinator.com/item?id=48765725)
- Source: [arxiv.org](https://arxiv.org/abs/2601.13082)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T18:47:11Z

## Translation

タイトル: LLM 主導のアルゴリズム取引におけるヘッドラインの操作
記事のタイトル: [2601.13082] 敵対的なニュースと逸失利益: LLM 主導のアルゴリズム取引における見出しの操作
説明: arXiv 論文 2601.13082 の要約ページ: 敵対的なニュースと逸失利益: LLM 主導のアルゴリズム取引における見出しの操作

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 1 月 19 日に提出]
タイトル: 敵対的なニュースと逸失利益: LLM 主導のアルゴリズム取引における見出しの操作
要約: 大規模言語モデル (LLM) は金融分野でますます採用されています。テキスト データを分析する優れた能力により、金融関連ニュースのセンチメントを推測するのに適しています。このようなフィードバックは、売買の決定を導くためにアルゴリズム取引システム (ATS) によって活用できます。ただし、この手法には、脅威アクターが LLM を誤解させることを目的とした「敵対的なニュース」を作成する可能性があるというリスクが伴います。特に、ニュースの見出しには、人間の読者には見えないものの、依然として LLM によって取り込まれている「悪意のある」コンテンツが含まれている可能性があります。これまでの研究ではテキストの敵対的な例が研究されてきましたが、LLM がサポートする ATS に対するシステム全体への影響は、金銭的リスクの観点からはまだ定量化されていません。この脅威に対処するために、ATS に直接アクセスできないが、株式関連のニュースのヘッドラインを 1 日で変更できる敵を想定します。私たちは、金融の文脈における人間には感知できない 2 つの操作を評価します。1 つは銘柄認識時にモデルのルートを誤る Unicode の同形文字置換、もう 1 つはニュース見出しの感情を変える隠しテキスト条項です。 LSTM ベースの価格予測と LLM 由来のセンチメント (FinBERT、FinGPT、FinLLaMA、および 6 つの汎用 LLM) を融合する現実的な ATS を Backtrader に実装し、ポートフォリオ メトリックを使用して金銭的影響を定量化します。実世界のデータを使った実験では、14 か月にわたって 1 日限りの攻撃を操作すると、LLM とセキュリティを確実に誤解させる可能性があることが示されています。

年間収益は最大 17.7 パーセントポイント増加します。現実世界での実現可能性を評価するために、人気のあるスクレイピング ライブラリと取引プラットフォームを分析し、27 人の FinTech 実務者に調査を行い、仮説を確認しました。このセキュリティ問題については、取引プラットフォームの所有者に通知しました。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
関連するDOI:
https://doi.org/10.1109/SaTML68715.2026.00053
もっと学ぶために集中する
関連リソースにリンクする DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2601.13082: Adversarial News and Lost Profits: Manipulating Headlines in LLM-Driven Algorithmic Trading

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 19 Jan 2026]
Title: Adversarial News and Lost Profits: Manipulating Headlines in LLM-Driven Algorithmic Trading
Abstract: Large Language Models (LLMs) are increasingly adopted in the financial domain. Their exceptional capabilities to analyse textual data make them well-suited for inferring the sentiment of finance-related news. Such feedback can be leveraged by algorithmic trading systems (ATS) to guide buy/sell decisions. However, this practice bears the risk that a threat actor may craft "adversarial news" intended to mislead an LLM. In particular, the news headline may include "malicious" content that remains invisible to human readers but which is still ingested by the LLM. Although prior work has studied textual adversarial examples, their system-wide impact on LLM-supported ATS has not yet been quantified in terms of monetary risk. To address this threat, we consider an adversary with no direct access to an ATS but able to alter stock-related news headlines on a single day. We evaluate two human-imperceptible manipulations in a financial context: Unicode homoglyph substitutions that misroute models during stock-name recognition, and hidden-text clauses that alter the sentiment of the news headline. We implement a realistic ATS in Backtrader that fuses an LSTM-based price forecast with LLM-derived sentiment (FinBERT, FinGPT, FinLLaMA, and six general-purpose LLMs), and quantify monetary impact using portfolio metrics. Experiments on real-world data show that manipulating a one-day attack over 14 months can reliably mislead LLMs and reduce annual returns by up to 17.7 percentage points. To assess real-world feasibility, we analyze popular scraping libraries and trading platforms and survey 27 FinTech practitioners, confirming our hypotheses. We notified trading platform owners of this security issue.
Focus to learn more
arXiv-issued DOI via DataCite
Related DOI :
https://doi.org/10.1109/SaTML68715.2026.00053
Focus to learn more
DOI(s) linking to related resources
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

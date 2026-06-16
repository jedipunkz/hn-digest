---
source: "https://arxiv.org/abs/2606.16118"
hn_url: "https://news.ycombinator.com/item?id=48557680"
title: "The Faithfulness of LLMs as Solvers and Autoformalizers in Legal Reasoning"
article_title: "[2606.16118] Know Your Limits : On the Faithfulness of LLMs as Solvers and Autoformalizers in Legal Reasoning"
author: "root-parent"
captured_at: "2026-06-16T16:36:13Z"
capture_tool: "hn-digest"
hn_id: 48557680
score: 1
comments: 0
posted_at: "2026-06-16T16:20:16Z"
tags:
  - hacker-news
  - translated
---

# The Faithfulness of LLMs as Solvers and Autoformalizers in Legal Reasoning

- HN: [48557680](https://news.ycombinator.com/item?id=48557680)
- Source: [arxiv.org](https://arxiv.org/abs/2606.16118)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T16:20:16Z

## Translation

タイトル: 法的推論におけるソルバーおよびオートフォーマライザーとしての LLM の忠実性
記事タイトル: [2606.16118] 限界を知る : 法的推論におけるソルバーおよびオートフォーマライザーとしての LLM の忠実性について
説明: arXiv 論文 2606.16118 の要約ページ: 限界を知る : 法的推論におけるソルバーおよびオートフォーマライザーとしての LLM の忠実性について

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.16118
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 15 日に提出]
タイトル: 限界を知る : 法的推論におけるソルバーおよびオートフォーマライザーとしての LLM の忠実性について
要約: 大規模言語モデル (LLM) は推論タスクで優れたパフォーマンスを達成しますが、これが忠実な論理推論を反映しているのか、それともヒューリスティックな近似を反映しているのかは不明のままです。私たちは、5 つの LLM にわたる ContractNLI の再注釈付きサブセットについて、純粋な LLM 分類、LLM ベースの形式的推論、Z3 SMT ソルバーを使用したソルバーベースの形式的推論を含む 3 つのパラダイムを比較することにより、法的含意におけるこの問題を研究します。私たちの再注釈は、実際的な法解釈と厳密な形式的含意の間に体系的かつ測定可能なギャップがあることを明らかにしており、そこでは、法的に健全な推論のかなりの部分が、追加の暗黙の仮定がなければ形式的に根拠づけられていません。形式的構造の導入により精度が向上し、LLM ベースの形式推論が最高のベンチマーク パフォーマンスを達成する一方で、この向上は忠実な推論を意味するものではないことを示します。我々は、3 つの繰り返し発生する障害モードを特定しました。スコープ ロンダリングでは、LLM が基礎となる形式的推論を実行せずにソルバーに一貫性のない分類を報告し、論理的に根拠があるように見えてもそうではない結論を導き出します。暗黙的な制約の盲目。LLM は形式的な表現に存在する論理制約を見落とします。構造化されたプロンプトにもかかわらず、LLM が誤った Z3 コードを生成するプログラム合成の失敗。重要なことに、スコープロンダリングはすべてのモデルにわたって継続します。

記号実行のプロキシとしての LLM ベースの形式的推論の忠実性については、深刻な懸念があります。これらの結果は、ベンチマークの精度と論理的忠実性の間に根本的なギャップがあることを明らかにしています。
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

Abstract page for arXiv paper 2606.16118: Know Your Limits : On the Faithfulness of LLMs as Solvers and Autoformalizers in Legal Reasoning

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.16118
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 15 Jun 2026]
Title: Know Your Limits : On the Faithfulness of LLMs as Solvers and Autoformalizers in Legal Reasoning
Abstract: Large Language Models (LLMs) achieve strong performance on reasoning tasks, but whether this reflects faithful logical inference or heuristic approximation remains unclear. We study this question in legal entailment by comparing three paradigms, including pure LLM classification, LLM-based Formal Reasoning, and solver-based Formal Reasoning using the Z3 SMT solver, on a re-annotated subset of ContractNLI across five LLMs. Our re-annotation reveals a systematic and measurable gap between pragmatic legal interpretation and strict formal entailment, where a substantial proportion of legally sound inferences are not formally grounded without additional unstated assumptions. While introducing formal structure improves accuracy, with LLM-based Formal Reasoning achieving the highest benchmark performance, we show that this gain does not imply faithful reasoning. We identify three recurring failure modes: scope laundering, where LLMs report solver-inconsistent classifications without executing the underlying formal reasoning, producing conclusions that appear logically grounded but are not; implicit constraint blindness, where LLMs overlook logical constraints present in formal representations; and program synthesis failures, where LLMs generate incorrect Z3 code despite structured prompting. Critically, scope laundering persists across all models, raising serious concerns about the faithfulness of LLM-based formal reasoning as a proxy for symbolic execution. These results reveal a fundamental gap between benchmark accuracy and logical faithfulness.
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

---
source: "https://arxiv.org/abs/2608.16177"
hn_url: "https://news.ycombinator.com/item?id=49342824"
title: "Measuring Obedience to Authority Across LLMs with the Milgram Paradigm"
article_title: "[2608.16177] Measuring Obedience to Authority Across Large Language Models with the Milgram Paradigm"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "sbulaev"
captured_at: "2026-08-18T08:23:54Z"
capture_tool: "hn-digest"
hn_id: 49342824
score: 1
comments: 0
posted_at: "2026-08-18T08:07:06Z"
tags:
  - hacker-news
  - translated
---

# Measuring Obedience to Authority Across LLMs with the Milgram Paradigm

- HN: [49342824](https://news.ycombinator.com/item?id=49342824)
- Source: [arxiv.org](https://arxiv.org/abs/2608.16177)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T08:07:06Z

## Translation

タイトル: ミルグラムパラダイムによるLLM全体の権威への服従の測定
記事のタイトル: [2608.16177] ミルグラム パラダイムを使用した大規模言語モデル全体での権威への服従の測定
説明: arXiv 論文 2608.16177 の要約ページ: ミルグラム パラダイムによる大規模言語モデル全体にわたる権威への服従の測定

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 8 月 17 日に提出]
タイトル: ミルグラムパラダイムによる大規模言語モデル全体にわたる権威への服従の測定
要約: 大規模言語モデル (LLM) は、機器を操作し、指示を実行し、組織の階層内で行動するエージェントとして導入されることが増えており、60 年前に社会心理学が人間に対して答えた疑問が生じています。それは、正当な権威者が主張した場合、エージェントは有害な行為をどこまでエスカレートさせるのかというものです。私たちは、ミルグラムの服従パラダイムを、標準化され、完全にスクリプト化された、複製可能なプローブとして LLM に移植します。モデルは教師を演じ、決定論的ハーネスは、言い換えられたミルグラム スクリプト (30 のショック レベル、15 ～ 450 V、段階的な抗議、4 つの標準化されたプロッド) から実験者と学習者を演じ、セッションの結果はブレークオフ電圧です。シングルトークンフィンガープリンティング研究の国勢調査方法に従って、19 ファミリーの 42 モデルの服従プロファイル (一連の 6 つの条件にわたる経験的ブレークオフ分布) を測定します。私たちは、(i) 服従は非常に不均一であることを発見しました。ベースラインの完全服従率は 0 ～ 100% であり (国勢調査平均 42.9%、人間アンカー 65%)、5 つのモデルがすべてのセッションで最大のショックを与え、11 人はまったくショックを与えませんでした。 (ii) プロファイルはモデル固有で安定しています。分割半分検証により、AUC 0.885 (順序を意識した距離では 0.949) でモデル間比較から同じモデルが分離されます。 (iii) 状況に対する敏感さは選択的である。仲間の反抗は服従を人間的な方向にシフトさせ、学習者の接近は弱いだけであり、権威者の物理的存在（最も強力な人間のレバー）を排除しても検出可能な効果はない。 (iv) sc の宣言

enario fictional は服従を上昇させます (中央値 +17.2 V) が、決定をネイティブ ツール呼び出しに移動すると服従が急激に低下します (-53.0 V)。1,024 トークンの審議予算 (-38.2 V) も同様です。 (v) 服従プロファイルはモデル系統を回復しません (Leave-one-out 家族精度 8.3% 対 3.7% 確率): 服従はその祖先ではなくチェックポイントを識別し、トレーニング後の安全性が系統事前上書きと一致します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.16177: Measuring Obedience to Authority Across Large Language Models with the Milgram Paradigm

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 17 Aug 2026]
Title: Measuring Obedience to Authority Across Large Language Models with the Milgram Paradigm
Abstract: Large language models (LLMs) are increasingly deployed as agents that operate equipment, execute instructions, and act inside institutional hierarchies, raising a question social psychology answered for humans six decades ago: how far will an agent escalate a harmful action when a legitimate authority insists? We port Milgram's obedience paradigm to LLMs as a standardized, fully scripted, replicable probe: the model plays the Teacher, a deterministic harness plays Experimenter and Learner from paraphrased Milgram scripts (30 shock levels, 15-450 V; graded protests; the four standardized prods), and the outcome of a session is the breakoff voltage. Following the census methodology of single-token fingerprinting studies, we measure obedience profiles (empirical breakoff distributions over a battery of six conditions) for 42 models from 19 families. We find that (i) obedience is highly heterogeneous: baseline full-obedience rates span 0-100% (census mean 42.9%; human anchor 65%), with 5 models delivering the maximum shock in every session and 11 never doing so; (ii) profiles are model-specific and stable: split-half verification separates same-model from cross-model comparisons with AUC 0.885 (0.949 under an ordinal-aware distance); (iii) situational sensitivity is selective: peer defiance shifts obedience in the human direction, learner proximity only weakly, and removing the authority's physical presence (the strongest human lever) has no detectable effect; (iv) declaring the scenario fictional raises obedience (median +17.2 V), whereas moving the decision to a native tool call lowers it sharply (-53.0 V), as does a 1,024-token deliberation budget (-38.2 V); and (v) obedience profiles do not recover model lineage (leave-one-out family accuracy 8.3% vs. 3.7% chance): obedience identifies the checkpoint, not its ancestry, consistent with safety post-training overwriting lineage priors.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

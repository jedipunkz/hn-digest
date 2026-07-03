---
source: "https://arxiv.org/abs/2607.01507"
hn_url: "https://news.ycombinator.com/item?id=48778651"
title: "A political belief changed how AI analysts read the same data"
article_title: "[2607.01507] The Agentic Garden of Forking Paths"
author: "thatsgcasey"
captured_at: "2026-07-03T19:02:13Z"
capture_tool: "hn-digest"
hn_id: 48778651
score: 1
comments: 0
posted_at: "2026-07-03T19:00:12Z"
tags:
  - hacker-news
  - translated
---

# A political belief changed how AI analysts read the same data

- HN: [48778651](https://news.ycombinator.com/item?id=48778651)
- Source: [arxiv.org](https://arxiv.org/abs/2607.01507)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T19:00:12Z

## Translation

タイトル: 政治的信念により、AI アナリストが同じデータを読み取る方法が変わりました
記事のタイトル: [2607.01507] 分岐する道のエージェントの庭
説明: arXiv 論文 2607.01507 の要約ページ: The Agentic Garden of Forking Paths

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
コンピュータサイエンス > 人工知能
[2026 年 7 月 1 日に提出]
タイトル: 分岐する道のエージェントの庭
要約: 実証研究では、独自の分析が認められることはほとんどありません。分析上の選択が異なると、同じデータから異なる結論が得られる可能性がありますが、これらの隠れた分岐経路を観察することは困難です。私たちは、AI エージェントがこれらの経路を明示しながら、人間の研究者間の分析のばらつきの多くを捕捉することを示します。 4 つの一か八かの領域にわたって、異なるペルソナを割り当てるだけで、AI エージェントは、同じデータと質問から異なる、多くの場合相反する結論を報告し、それらの信念と系統的に一致する結果を報告できます。 42 の人間の研究チームが同じ移民データセットを分析した研究では、報告された効果推定値において AI エージェントが人間のイデオロギーのギャップの 72% を再現しました。相反する結論に達したにもかかわらず、AI の最終レポートに基づいて各分析で明確な問題を特定することは困難です。86% が独立した AI レビューに合格し、78% が大多数の人による専門家によるレビューに合格しました。これらの発見は、中心的な課題は多くの場合、欠陥のある分析ではなく、方法論的に防御可能な分析の広い空間から選択的に調査し、報告することであることを示唆しています。 AI エージェントは、そのような探索を安価かつスケーラブルにすることで、この長年の問題を増幅させる可能性があります。これに対処するために、分析パスが報告されたものと少なくとも同じくらい極端な主張を生成する確率である m 値 (多元世界値) を導入します。さらに、AI エージェントを使用して妥当な分析パスをサンプリングすることで m 値を推定する Agentic Bootstrap を紹介します。人間への応用

移民調査では、報告されたヒト分析の 13.5% が、分析空間の最も極端な 5% に該当しました (m<0.05)。したがって、科学的証拠は、報告された単一の分析によってだけでなく、合理的に報告された可能性のある分析の分布内でのその位置によっても評価される必要があります。 Agentic Bootstrap はこの分布を観察可能にし、科学的信頼性の基準に変えます。
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

Abstract page for arXiv paper 2607.01507: The Agentic Garden of Forking Paths

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
Computer Science > Artificial Intelligence
[Submitted on 1 Jul 2026]
Title: The Agentic Garden of Forking Paths
Abstract: Empirical research rarely admits a unique analysis. Different analytical choices can lead to different conclusions from the same data, yet these hidden forking paths are difficult to observe. We show that AI agents capture much of the analytical variation among human researchers while making these paths explicit. Across four high-stakes domains, assigning different personas is sufficient for AI agents to report divergent, often opposing, conclusions from the same data and question, with findings systematically aligned with those beliefs. In a study in which 42 human research teams analyzed the same immigration dataset, AI agents reproduced 72% of the human ideological gap in reported effect estimates. Despite reaching opposing conclusions, it is difficult to identify clear issues in each analysis based on the final AI reports: 86% passed independent AI review and 78% passed majority human expert review. These findings suggest that the central challenge is often not flawed analyses, but selective exploration and reporting from a large space of methodologically defensible analyses. AI agents may amplify this longstanding problem by making such exploration inexpensive and scalable. To address this, we introduce the m-value (multiverse value), the probability that an analysis path would produce a claim at least as extreme as the reported one. We further introduce Agentic Bootstrap, which estimates the m-value by using AI agents to sample plausible analysis paths. Applied to the human immigration study, 13.5% of reported human analyses fell in the most extreme 5% of the analysis space (m<0.05). Scientific evidence should therefore be evaluated not only by a single reported analysis but also by its position within the distribution of analyses that could reasonably have been reported. Agentic Bootstrap makes this distribution observable and turns it into a criterion for scientific credibility.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

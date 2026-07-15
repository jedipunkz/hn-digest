---
source: "https://www.lesswrong.com/posts/hcq4ZDoijSjy3Wrba/how-much-of-ml-research-is-about-ai-safety-what-is-it-about"
hn_url: "https://news.ycombinator.com/item?id=48919012"
title: "How much of ML research is about AI safety, what is it about, and who's doing I"
article_title: "How much of ML research is about AI safety, what is it about, and who's doing it? — LessWrong"
author: "joozio"
captured_at: "2026-07-15T11:22:06Z"
capture_tool: "hn-digest"
hn_id: 48919012
score: 2
comments: 0
posted_at: "2026-07-15T11:02:39Z"
tags:
  - hacker-news
  - translated
---

# How much of ML research is about AI safety, what is it about, and who's doing I

- HN: [48919012](https://news.ycombinator.com/item?id=48919012)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/hcq4ZDoijSjy3Wrba/how-much-of-ml-research-is-about-ai-safety-what-is-it-about)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T11:02:39Z

## Translation

タイトル: AI の安全性に関する ML 研究の割合、内容、および誰が行っているのか
記事のタイトル: AI の安全性に関する ML 研究はどの程度行われていますか? それは何についてであり、誰が行っているのですか? — 間違っていない
説明: 大きな ML カンファレンスでは AI の安全性に関する論文が何件ありますか?何を研究し、誰が執筆していますか?包括的な分析。
> ウェブサイト：https://ai-saf…

記事本文:
ログイン ML 研究の中で AI の安全性に関するものはどれくらいありますか? それは何についてのもので、誰がそれを行っているのですか? — LessWrong AI 調整現場構築 AI フロントページ 10
ML 研究のどれくらいが AI の安全性に関するもので、それは何についてのもので、誰がそれを行っているのでしょうか?
大規模な ML カンファレンスでは AI の安全性に関する論文が何本ありますか?何を研究し、誰が執筆しているのでしょうか?包括的な分析。
> ウェブサイト: https://ai-safety-tracker-website.vercel.app/
> データ、コード、プロット: https://github.com/SomaxSoma/AI-Safety-Research-Tracker
TL;DR: 私たちは、2019 年から 2026 年までに ICLR、ICML、NeurIPS で受理されたすべての論文を、各タイトルと要約を読み取る LLM を使用して分類しました。そのうち 2,328 件 (4.2%) が AI の安全性に関する論文です。受理された論文に占めるセーフティの割合は、2019年の0.3%から2026年には8.3%に増加し、およそ25倍に増加しました。この投稿は主な結果の参考になります。インタラクティブな Web サイトでは、すべての論文、そのサブドメイン、および分類子の推論を参照できます。
私たちは、AI の安全性研究で何が起こっているのかをできるだけ広範囲に把握したいと考えていました。これは、研究に適切な優先順位を付け、より正確な変化理論を開発できるようにするために必要であると思われます。これを達成する 1 つの方法は、AI の安全性をテーマに発表されたすべての主要会議論文の統計を分析することですが、それらを実際に測定したデータセットは見つかりませんでした。そこで、次のようなものを構築しました。
- ICLR (2019–2026)、ICML (2019–2026)、および NeurIPS (2019–2025) で承認されたすべての論文 (合計 55,794) が LLM (DeepSeek V4 Flash) によって読み取られ、タイトルと要約から AI の安全性 (フロンティアとミスアライメント)、真実性、信頼性と XAI、倫理と公平性の 4 つのクラスのいずれかに分類されます。 、または一般的な機能。
- 各安全文書には、17 の安全サブドメイン (解釈可能性、調整トレーニング、スカラ) のいずれかが割り当てられます。

ble の監視、危険な能力の評価など) を評価し、安全性にどのように重点的に取り組んでいるかを 7 ポイントのルーブリックで採点します。
- 安全文書の別の全文パスにより、どの安全組織、プログラム、資金提供者が背後にいるかを検出し、所属と承認を区別します。
これらの結果を再現するために必要なコード (分類プロンプト、CSV など) はすべて公開されています。この Web サイトには、2,328 件の安全に関する論文すべてを検索できるエクスプローラーがあり、それぞれが OpenReview ページにリンクし、分類子の推論を示しています。私たちのデータセットがこの分野のさらなる研究に使用できることを願っています。
これらのカンファレンスの一般採択率は 20 ～ 30% にとどまりましたが、年間の安全性論文の数は 2019 年から 2026 年にかけて 100 倍以上に増加しました。さらに、全採択数に占める安全性論文の割合は 2019 年から 2026 年にかけて 25 倍以上増加しており、AI 安全分野の大幅な成長を示しています。
(2026 年は ICLR 2026 と ICML 2026 を対象としています。NeurIPS 2026 はまだ開催されていません。)
グラフ 1: 3 つの会議全体でプールされた年別の安全性シェア
これらの会議に受理された論文の数が急速に増加したにもかかわらず、安全性に関する論文の割合が年々、特に 2023 年から 2024 年にかけて急速に増加していることが観察されています。3 つの会議 (ICLR、ICML、および NeurIPS) の安全性に関する論文の割合は同様です。
どのような安全性研究なのか
すべての年を合計すると、解釈可能性が最大のサブドメインであり、次に調整トレーニング、敵対的堅牢性、レッドチーム化が続きます。
計画 3: サブドメインの内訳、すべての年と会場
時代とともにその構成も変化してきました。初期（2020 ～ 2022 年）では、敵対的堅牢性が最大の領域でしたが、2023 年には解釈可能性がそれを上回り、より新しい領域（危険能力の評価）が追加されました。

、陰謀と欺瞞、監視）は本質的にゼロから現れます。
制御、モデル生物、AI 福祉などの一部のトピックは、会議議事録にほとんど載っていないままです (2,328 件の論文中、それぞれ 9 件、5 件、2 件の論文)。
安全に関する論文については、2 回目の全文分析を実行しました。安全に関する組織、プログラム、資金提供者の厳選されたリストと各論文の前付および謝辞とをキーワード照合し、LLM に各ヒット (所属、謝辞、単なる言及) を検証させ、論文の主要な組織 (作業を主導したと推定される組織) を選択します。企業は、その著者が論文を主導した場合にのみ主著者としてカウントされます。たとえば、大学の著者の中で唯一の DeepMind 共著者がいる場合、DeepMind は主著者にはなりません。フェローシップ プログラム (MATS、SPAR、Anthropic Fellows など) は、謝辞にのみ表示される場合でも、ホスティング組織としてカウントされます。
私たちが全文を取得できた約 2,300 件の安全に関する文書のうち、約 600 件は追跡された安全組織または資金提供者とのつながりが確認されています。
グラフ 4: 研究組織ごとの論文数、主な組織ごとにカウント
主な属性別では、MATS (48)、Mila (46) [1]、Google DeepMind (42) がリストのトップにあります。所属別にカウントすると、業界の研究機関が圧倒的に多く、Google DeepMind は 107 件、OpenAI は 74 件、Anthropic は 51 件の論文に掲載されており、研究室の研究者が外部主導の研究を共同執筆したり指導したりする頻度が高いことを反映しています。 (Web サイトの「発行者」タブには両方のビューがあります。)
プロット 5: 謝辞に資金提供者が記載されている
Open Philanthropy は 120 件の論文でクレジットされており、他のどの慈善活動資金提供者よりも桁違いに多くなっています (Long-Term Future Fund: 10、Future of Life Institute: 8)。
興味深い傾向の 1 つは、生産量が増加しているにもかかわらず、安全エコシステムにおける安全書類のシェアがわずかに縮小していることです。 2019 年には、9 件の安全 P のうち 5 件が

apers には追跡された組織または資金提供者がおり、2026 年には 954 件中 150 件になりました [2] 。専用のエコシステムは毎年より多くの出版物を発表していますが、この分野の成長は主に外部の研究者によるものです。
これらの傾向は、時間の経過とともに、効果的な介入は、論文や研究者の生の数を増やすだけでなく、安全性に関心のある既存の研究者による研究の質/テーマの関連性/長期的な影響を改善すること、そして研究をより多様で一般化可能で影響力があり、適用可能なものにすることに焦点を当てるべきであることを示唆していると考えています。
計画 6: 組織が支援する年間の安全性に関する文書
- 分類は LLM による、タイトル + 要約から行われます。フロンティア AI の安全性と隣接するクラス (真実性/幻覚、XAI、公平性) の間の境界は本質的に曖昧であるようですが、ルーブリック全体と分類子の論文ごとの推論を公開します。
- 組織分析は下限です。厳選された組織リスト、ページ 1 の所属、および最初の承認ウィンドウのみが表示されます。付録の奥深くに埋もれている指導功績が見逃される可能性があります。
- 2026 は部分的であり (NeurIPS はまだありません)、ICML 2026 の全文は組織分析に部分的にしか利用できませんでした (ICML の 45 件の論文は、初期バージョンとカメラ対応バージョンの間でタイトルが大幅に変更されているため、取得できませんでした)。
- これらの会場での受け入れは、安全作業の大部分を占めるワークショップの論文、arXiv のプレプリント、レポート、ブログ投稿など、未提出または拒否された作業を過小評価します。
ウェブサイト https://ai-safety-tracker-website.vercel.app/ には、会議ごとのビュー、年ごとのサブドメイン プロット、全 2,328 件の論文の検索可能なテーブル (各行で分類子の逐語的推論が開き、OpenReview へのリンクが表示されます)、組織別および資金提供者別のプロット、および実験など、上記すべてのインタラクティブ バージョンがあります。

l arXiv トレンドライン。リポジトリには、すべての CSV と、分類プロンプトを含むすべての CSV を再生成するためのスクリプトが含まれています。
間違いを見つけた場合は、お気軽にコメント、DM、または GitHub の問題を送信してください。修正は大歓迎です。
将来の作業には、この分析を arXiv プレプリントに拡張し、場合によっては LessWrong や AlignmentForum の投稿にも拡張することが含まれる可能性があります。もう 1 つの可能性は、引用ネットワークを分析することです。たとえば、組織関連の論文が独立した論文を引用しているかどうかを確認し、クラスターと引用の流れを判断します。
フィードバックを提供してくださった皆様に感謝します。
このプロジェクトは BlueDot Rapid Grants によって支援されました。
^ 企業以外の主要組織を割り当てる際に分類子が意図的に緩かったため、Mila メンバーが共著または指導した一部の論文が主要組織として Mila に割り当てられましたが、Anthropic や DeepMind では同様のことが起こらず、これらの組織からの二次著者をカウントするとシェアが大幅に増加することになります。
^ 一部の ICML 2026 論文は、オリジナル バージョンとカメラ対応バージョンの間でタイトルが変更されたために検索できなかったため、合計 999 件のうち 954 件が検索可能でした。
0 コメント 0 AI 調整フィールド構築 AI フロントページ 10

## Original Extract

How many AI safety papers are at the big ML conferences, what do they study, and who writes them? A comprehensive analysis.
> Website: https://ai-saf…

Login How much of ML research is about AI safety, what is it about, and who's doing it? — LessWrong AI Alignment Fieldbuilding AI Frontpage 10
How much of ML research is about AI safety, what is it about, and who's doing it?
How many AI safety papers are at the big ML conferences, what do they study, and who writes them? A comprehensive analysis.
> Website: https://ai-safety-tracker-website.vercel.app/
> Data, code and plots: https://github.com/SomaxSoma/AI-Safety-Research-Tracker
TL;DR: We classified every paper accepted at ICLR, ICML and NeurIPS from 2019 through 2026, using an LLM that reads each title and abstract. 2,328 of them (4.2%) are AI safety papers. Safety's share of accepted papers grew from 0.3% in 2019 to 8.3% in 2026, roughly a 25-fold increase. This post is a reference for the main results. The interactive website lets you browse every paper, its subdomain, and the classifier's reasoning.
We wanted to have an overview of what is going on in AI safety research while being as broad as possible, this seems necessary to be able to prioritize research correctly, and develop more precise theories of change. One way to achieve this is by analyzing statistics of all the published main-conference papers on the topic of AI safety, but we couldn't find a dataset that actually measured them. So we built one:
- Every accepted paper at ICLR (2019–2026), ICML (2019–2026) and NeurIPS (2019–2025) (55,794 in total) is read by an LLM (DeepSeek V4 Flash), which classifies it from its title and abstract into one of four classes: AI safety (frontier & misalignment) , truthfulness, reliability & XAI , ethics & fairness , or general capabilities .
- Each safety paper is then assigned one of 17 safety subdomains (interpretability, alignment training, scalable oversight, dangerous-capability evals, …) and scored on a 7-point rubric for how centrally it addresses safety.
- A separate full-text pass over the safety papers detects which safety organizations, programs and funders are behind them, distinguishing affiliations and acknowledgments.
All the code required to reproduce these results (classification prompts, CSVs, etc.) is publicly available. The website has a searchable explorer of all 2,328 safety papers, each linking to its OpenReview page and showing the classifier's reasoning. We hope that our datasets can be used for further studies of the field.
While the general acceptance rates of these conferences stayed between 20 and 30%, the number of safety papers per year increased over 100 times from 2019 to 2026. Moreover, the share of safety papers among all the acceptances increased over 25 times from 2019 to 2026, which shows a substantial growth in the AI safety field.
(2026 covers ICLR 2026 and ICML 2026; NeurIPS 2026 hasn't happened yet.)
PLOT 1: safety share by year, pooled across the three conferences
We observe that even with a rapid increase in the number of papers accepted to these conferences, the share of safety papers increased rapidly over the years, especially between 2023 and 2024. The three conferences (ICLR, ICML and NeurIPS) have a similar share of safety papers.
What kind of safety research it is
Summing across all years, interpretability is the largest subdomain, followed by alignment training, adversarial robustness and red-teaming.
PLOT 3: subdomain breakdown, all years and venues
The composition has also shifted over time. In the early years (2020–2022) adversarial robustness was the largest area, which was surpassed by interpretability in 2023, with newer areas (dangerous-capability evals, scheming and deception, monitoring) appearing essentially from zero.
Some topics, such as control, model organisms and AI welfare, remain nearly absent from conference proceedings (9, 5 and 2 papers respectively, out of 2,328).
For the safety papers, we ran a second, full-text analysis: keyword-match a curated list of safety organizations, programs and funders against each paper's front matter and acknowledgments, then have an LLM verify each hit (affiliation vs. acknowledgment vs. mere mention) and pick the paper's primary org , the organization presumed to have led the work. A company only counts as primary if its authors lead the paper, for example, a lone DeepMind co-author among university authors doesn't make DeepMind the primary author. Fellowship programs (MATS, SPAR, Anthropic Fellows, …) count as the hosting org even when they appear only in the acknowledgments.
Of the ~2,300 safety papers whose full text we could retrieve, about 600 have a verified connection to a tracked safety org or funder.
PLOT 4: papers per research org, counted by primary org
By primary attribution, MATS (48), Mila (46) [1] and Google DeepMind (42) top the list. Counted by any affiliation instead, the industry labs dominate, Google DeepMind appears on 107 papers, OpenAI on 74, Anthropic on 51, reflecting how often lab researchers co-author and mentor externally-led work. (The website's "Who Publishes" tab has both views.)
PLOT 5: funders credited in the acknowledgments
Open Philanthropy is credited in 120 papers, an order of magnitude more than any other philanthropic funder (Long-Term Future Fund: 10, Future of Life Institute: 8).
One interesting trend is that the safety ecosystem's share of safety papers is slightly shrinking even as its output grows. In 2019, 5 of the 9 safety papers had a tracked org or funder behind them, in 2026 it's 150 of 954 [2] , the dedicated ecosystem publishes more every year, but the field's growth is mostly coming from researchers outside it.
We think that these trends suggest that over time, effective interventions should focus on improving the quality/topic relevance/long-term impact of the research coming from existing researchers interested in safety, and on making research more diverse, generalizable, impactful, and applicable, rather than only increasing the raw count of papers and researchers.
PLOT 6: org-backed safety papers per year
- Classification is by LLM, from title + abstract. The boundary between frontier AI safety and the neighboring classes (truthfulness/hallucinations, XAI, fairness) seems to be inherently fuzzy, we publish the full rubric and the classifier's per-paper reasoning.
- The org analysis is a lower bound. It only sees a curated org list, page-one affiliations and the first acknowledgments window; a mentorship credit buried deep in the appendix can be missed.
- 2026 is partial (no NeurIPS yet), and ICML 2026 full texts were only partially available for the org analysis (45 papers from ICML had significant title changes between the initial and camera-ready versions, and thus weren't retrievable).
- Acceptance at these venues undercounts non-submitted or rejected work, such as workshop papers, arXiv preprints, reports and blog posts, which is a large part of safety work.
The website https://ai-safety-tracker-website.vercel.app/ has the interactive version of everything above: per-conference views, per-year subdomain plots, a searchable table of all 2,328 papers (each row opens the classifier's verbatim reasoning and links to OpenReview), the by-org and by-funder plots, and an experimental arXiv trend line. The repo has every CSV and the scripts to regenerate all of it, including the classification prompt.
If you spot a mistake, feel free to send a comment, DM or GitHub issue, corrections are very welcome.
Future work could include extending this analysis to arXiv preprints, and possibly LessWrong and AlignmentForum posts. Another possibility is to analyze the citation networks, for example checking whether org-affiliated papers cite independent papers, and determining clusters and citation flow.
Thanks to every person who provided feedback.
This project was supported by BlueDot Rapid Grants.
^ The classifier was intentionally lax when assigning non-corporate primary orgs, this caused some papers coauthored or mentored by Mila members to be assigned Mila as their primary org, while the same didn't happen for Anthropic or DeepMind, whose share would increase significantly if counting secondary authors from these organizations.
^ There were 954 retrievable papers from the total of 999, as some ICML 2026 papers were not retrievable due to title changes between the original and camera-ready versions.
0 Comments 0 AI Alignment Fieldbuilding AI Frontpage 10

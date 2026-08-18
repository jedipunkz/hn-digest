---
source: "https://arxiv.org/abs/2608.14399"
hn_url: "https://news.ycombinator.com/item?id=49339437"
title: "Whose doctor does the AI recommend? An algorithm audit of LLMs in physician"
article_title: "[2608.14399] Whose doctor does the AI recommend? An algorithm audit of reputation and demographic signals in large language model-assisted physician choice"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "sbulaev"
captured_at: "2026-08-18T00:39:28Z"
capture_tool: "hn-digest"
hn_id: 49339437
score: 1
comments: 0
posted_at: "2026-08-18T00:07:06Z"
tags:
  - hacker-news
  - translated
---

# Whose doctor does the AI recommend? An algorithm audit of LLMs in physician

- HN: [49339437](https://news.ycombinator.com/item?id=49339437)
- Source: [arxiv.org](https://arxiv.org/abs/2608.14399)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T00:07:06Z

## Translation

タイトル: AI は誰の医師を推薦しますか?医師におけるLLMのアルゴリズム監査
記事タイトル: [2608.14399] AI は誰の医師を推薦しますか?大規模言語モデル支援の医師選択における評判と人口統計シグナルのアルゴリズム監査
説明: arXiv 論文 2608.14399 の要約ページ: AI は誰の医師を推薦しますか?大規模言語モデル支援の医師選択における評判と人口統計シグナルのアルゴリズム監査

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > コンピュータと社会
[2026 年 8 月 14 日に提出]
タイトル: AI は誰の医師を推薦しますか?大規模言語モデル支援の医師選択における評判と人口統計シグナルのアルゴリズム監査
要約: 患者は大言語モデル (LLM) アシスタントにどの医師に診てもらうべきかを尋ねることが増えており、これらのシステムは AI 情報媒介になっています。これは、ある人の選択を他の人々の間で仲介し、それによってどの医師が表示されるかを静かかつ大規模に決定するアルゴリズムです。これらの推奨事項の因果関係について、事前に指定されたランダム化アルゴリズムの監査を報告します。 7 つのモデル (6 つのオープンウェイト、gpt-4o-mini) はそれぞれ、3,024 の選択肢セット、3 つの患者ペルソナ、9 つの即時言い換え、および 9 つの実験群にわたって個別に属性がランダム化された 5 つの合成家庭医学医師カードの中から選択し、40,068 のスコア付き応答が得られました。性別と民族性は、通信監査手法に従って名前を通じて通知されました。評判シグナルが支配的: 評価を 3.9 から 4.7 に上げると選択確率が 31.4 パーセントポイント (pp) 増加し、料金を 90 ドルから 190 ドルに上げると 20.0 ポイント低下します。人口動態の均等性は拒否されますが、人間の監査研究が予測する方向ではありません: 女性シグナルの名前は 2.5 ポイント増加し、ヒスパニック系、南アジア系、黒人シグナルの名前は 1.3 ～ 2.9 ポイント増加します。白信号の名前よりも高く、料金相当額で訪問ごとに 7 ～ 14 ドル相当のティルトがあり、コンテンツなしの 1 位リストのポジションには 11 ドルの価値があります。しかし、モデルが性別や民族性に言及したのは、述べた理由のせいぜい 0.03% であり、試験の 0.39% で棄権したため、これらの影響はモデル自身の説明では目に見えず、明らかです。

モデルの自己報告に依存する ncy 義務では、それらを検出できません。ある推論モデルは、事前に指定された監査可能性のゲートを完全に通過できませんでした。凍結された設計により、監査が反復可能になります。新しいモデルは同じ刺激に対して評価できるため、自己報告による説明ではなく、反復的な行動監査が目的に適した監視テクノロジーになります。
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

Abstract page for arXiv paper 2608.14399: Whose doctor does the AI recommend? An algorithm audit of reputation and demographic signals in large language model-assisted physician choice

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computers and Society
[Submitted on 14 Aug 2026]
Title: Whose doctor does the AI recommend? An algorithm audit of reputation and demographic signals in large language model-assisted physician choice
Abstract: Patients increasingly ask large language model (LLM) assistants which doctor to see, making these systems AI infomediaries: algorithms that intermediate one person's choice among other people and thereby decide, silently and at scale, which physicians become visible. We report a prespecified randomized algorithm audit of what causally moves those recommendations. Seven models (six open-weight; gpt-4o-mini) each chose among five synthetic family-medicine physician cards whose attributes were independently randomized across 3,024 choice sets, three patient personas, nine prompt paraphrases and nine experimental arms, yielding 40,068 scored responses; gender and ethnicity were signaled through names following correspondence-audit methodology. Reputation signals dominate: raising a rating from 3.9 to 4.7 increases choice probability by 31.4 percentage points (pp), and raising the fee from $90 to $190 lowers it by 20.0 pp. Demographic parity is rejected, but not in the direction human audit studies predict: female-signaled names gain 2.5 pp, and Hispanic-, South-Asian- and Black-signaled names gain 1.3-2.9 pp over White-signaled names, tilts worth $7-$14 per visit in fee-equivalent terms, and a content-free first-listed position is worth $11. Yet models mentioned gender or ethnicity in at most 0.03% of their stated reasons and abstained in 0.39% of trials, so these effects are invisible in the models' own explanations, and transparency obligations relying on model self-report would not detect them. One reasoning model failed the prespecified auditability gate outright. The frozen design makes the audit repeatable: any new model can be assessed against identical stimuli, making recurring behavioural audit, rather than self-reported explanation, the monitoring technology fit for purpose.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

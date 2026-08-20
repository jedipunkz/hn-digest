---
source: "https://arxiv.org/abs/2603.27017"
hn_url: "https://news.ycombinator.com/item?id=49373473"
title: "Google's AI photoscanner can determine body fat through selfies"
article_title: "[2603.27017] Beyond BMI: Smartphone Body Composition Phenotyping for Cardiometabolic Risk Assessment"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "Phreaker00"
captured_at: "2026-08-20T12:26:14Z"
capture_tool: "hn-digest"
hn_id: 49373473
score: 2
comments: 1
posted_at: "2026-08-20T12:06:37Z"
tags:
  - hacker-news
  - translated
---

# Google's AI photoscanner can determine body fat through selfies

- HN: [49373473](https://news.ycombinator.com/item?id=49373473)
- Source: [arxiv.org](https://arxiv.org/abs/2603.27017)
- Score: 2
- Comments: 1
- Posted: 2026-08-20T12:06:37Z

## Translation

タイトル: Google の AI フォトスキャナーは自撮り写真から体脂肪を測定できる
記事タイトル: [2603.27017] BMI を超えて: 心臓代謝リスク評価のためのスマートフォン体組成表現型解析
説明: arXiv 論文 2603.27017 の要約ページ: BMI を超えて: 心臓代謝リスク評価のためのスマートフォンの体組成表現型解析

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
定量的生物学 > 定量的方法
[2026 年 3 月 27 日に提出 ( v1 )、最終改訂日 2026 年 4 月 6 日 (このバージョン、v2)]
タイトル: BMI を超えて: 心臓代謝リスク評価のためのスマートフォン体組成表現型解析
要約: Body Mass Index (BMI) は、心臓代謝の健康状態を表す指標として広く利用されていますが、不正確です。真の体組成を評価することは優れていますが、デュアルエネルギー X 線吸光光度計 (DXA) のようなゴールドスタンダードの方法は拡張性がありません。私たちは、スマートフォンの画像から体組成を推定する手法「PhotoScan」を開発・検証することで、このギャップに対処しています。英国のバイオバンク参加者 (N=35,323) でディープラーニング モデルを事前トレーニングし、多様な民族性、年齢、体脂肪分布を持つ新たに採用された臨床コホート (PhotoBIA コホート、N=677) で微調整し、総体脂肪率 (BF%、MAE = 2.15%)、アンドロイドとガイノイドの脂肪比 (A/G、MAE = 0.11) について DXA に対して高い精度を達成しました。内臓脂肪と皮下脂肪の面積比 (V/S、MAE = 0.09)。このモデルの一般化可能性は独立した代謝健康研究コホート (MetabolicMosaic コホート、参加者 N=132) で実証され、BF% で 2.13%、A/G で 0.09、V/S で 0.09 の MAE を達成しました。次に、インスリン抵抗性 (IR) を予測することにより、MetabolicMosaic コホートにおけるこれらの指標の臨床的有用性を評価しました。 PhotoScan 由来の体組成指標をベースライン人口統計モデル (年齢、性別、BMI) に追加すると、インスリン抵抗性分類が大幅に改善されました (受信者動作特性曲線下面積「AUROC」 76.0% 対 69.2%、DeLong テスト p=0.002、ネット再分類指数「NRI」 0.593)。重要なのは、このアクセス可能なスマートフォンの方法は、ほぼ次のパフォーマンスを達成したということです。

これは、臨床グレードの DXA データをベースライン人口統計モデルに追加することに相当します (AUROC 77.3% 対 69.2%、DeLong 検定 p=0.004、NRI 0.748)。これらの発見は、スマートフォンベースの表現型解析が、BMI や人体計測学では見逃していた臨床的に意味のあるリスクシグナルを捕捉し、心臓代謝リスク層別化のための DXA に代わるスケーラブルな代替手段を提供することを示しています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2603.27017: Beyond BMI: Smartphone Body Composition Phenotyping for Cardiometabolic Risk Assessment

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Quantitative Biology > Quantitative Methods
[Submitted on 27 Mar 2026 ( v1 ), last revised 6 Apr 2026 (this version, v2)]
Title: Beyond BMI: Smartphone Body Composition Phenotyping for Cardiometabolic Risk Assessment
Abstract: Body Mass Index (BMI) is a widely accessible but imprecise proxy of cardiometabolic health. While assessing true body composition is superior, gold-standard methods like Dual-Energy X-ray Absorptiometry (DXA) are not scalable. We address this gap by developing and validating "PhotoScan," a method to estimate body composition from smartphone imagery. We pretrained a deep learning model on UK Biobank participants (N=35,323) and fine-tuned on a newly recruited clinical cohort (PhotoBIA cohort, N=677) with diverse ethnicity, age, and body fat distribution, achieving high accuracy against DXA for total body fat percentage (BF%, MAE = 2.15%), Android-to-Gynoid fat ratio (A/G, MAE = 0.11), and visceral-to-subcutaneous fat area ratio (V/S, MAE = 0.09). Generalizability of the model was demonstrated on an independent metabolic health study cohort (MetabolicMosaic cohort, N=132 participants), achieving MAEs of 2.13% for BF%, 0.09 for A/G, and 0.09 for V/S. We then evaluated the clinical utility of these metrics in the MetabolicMosaic cohort by predicting insulin resistance (IR). Adding PhotoScan-derived body composition metrics to baseline demographics model (Age, Sex, BMI) significantly improved insulin resistance classification (Area Under the Receiver Operating Characteristic Curve "AUROC" 76.0% vs 69.2%, DeLong test p=0.002, Net Reclassification Index "NRI" 0.593). Crucially, this accessible smartphone method achieved performance nearly equivalent to adding clinical-grade DXA data to baseline demographics model (AUROC 77.3% vs 69.2%, DeLong test p=0.004, NRI 0.748). These findings demonstrate that smartphone-based phenotyping captures clinically meaningful risk signals missed by BMI and anthropometrics, offering a scalable alternative to DXA for cardiometabolic risk stratification.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

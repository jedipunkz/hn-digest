---
source: "https://devblogs.microsoft.com/sustainable-software/the-carbon-footprint-of-ai/"
hn_url: "https://news.ycombinator.com/item?id=48621805"
title: "The Carbon Footprint of AI (2020)"
article_title: "The Carbon Footprint Of AI - Sustainable Software"
author: "reconnecting"
captured_at: "2026-06-21T19:30:50Z"
capture_tool: "hn-digest"
hn_id: 48621805
score: 1
comments: 0
posted_at: "2026-06-21T19:26:57Z"
tags:
  - hacker-news
  - translated
---

# The Carbon Footprint of AI (2020)

- HN: [48621805](https://news.ycombinator.com/item?id=48621805)
- Source: [devblogs.microsoft.com](https://devblogs.microsoft.com/sustainable-software/the-carbon-footprint-of-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T19:26:57Z

## Translation

タイトル: AI の二酸化炭素排出量 (2020)
記事のタイトル: AI の二酸化炭素排出量 - 持続可能なソフトウェア
説明: AI の二酸化炭素排出量は指数関数的に増加しています。より多くのデータを必要とする大規模なモデルが「RedAI」に貢献 - 新しいアプローチが必要

記事本文:
メインコンテンツにスキップ
開発者ブログ
AI
すべての .NET 投稿
.NETマウイ
ASP.NETコア
ブレイザー
エンティティフレームワーク
C++
C#
F#
TypeScript
NuGet
整備
中国語の .NET ブログ
開発者向けのマイクロソフト
エージェントフレームワーク
クラウドから開発する
Xcode
ISE開発者
TypeScript
パワーシェル
パイソン
ジャワ
中国語の Java ブログ
行く
Microsoft Edgeの開発
Microsoft 365 開発者
Microsoft Entra ID 開発者
Microsoft Entra PowerShell
ビジュアルスタジオ
Visual Studioコード
アスパイア
アズールのすべて
Azure SDK
Azure VM ランタイム チーム
マイクロソフトアジュール
Azure Cosmos DB
AzureドキュメントDB
Azure データスタジオ
Azure SQL
DevOps
ダイレクトX
マイクロソフトファウンドリ
パワープラットフォーム
Oデータ
統合データモデル (IDEA)
Windows コマンドライン
#ifdef Windows
MSIX 内
MIDIと音楽
ネイティブに反応する
古いもの、新しいもの
Windows 開発者
開発者ブログ
人工知能 (AI) は、気候変動との闘い方を変える可能性を秘めています。しかし、AI による二酸化炭素排出量は今後 10 年間で指数関数的に増加し、2025 年まで全世界で 44% 近くの CAGR で成長すると予測されています。
業界はより大きなモデル (GPT-3 など) に向かう傾向にあります。これらのモデルでは、増大し続けるデータセット、計算予算が必要となり、モデルのライフサイクル全体にわたって莫大なエネルギー料金が発生します。 AI モデルの計算コストは​​数か月ごとに 2 倍になり、2012 年から 2018 年にかけて推定 300,000 倍に増加しました。過去 2 年間で、パラメータの数は 170 倍に増加しました。現在、アーキテクチャ検索を通じて単一の 2 億 1300 万パラメーターの NLP 深層学習モデルをトレーニングすると、ガソリンを含むアメリカ車 5 台の寿命と同じ二酸化炭素排出量を生成できます。
1 つの NLP モデルは、ガソリンを含む自動車 5 台分と同じ二酸化炭素排出量を誇ります。
出典: Emma Strubell、カーネギーメロン大学
環境の持続可能性を考慮する必要がある

AI の責任ある開発と応用に向けた原則の 1 つとして。このようなテクノロジーを使用する利点はその欠点を上回るはずであり、AI の隠れたコストについての議論を最前線に引き上げる時期が来ています。
ジェリー・マクガバン氏の著書「World Wide Waste」によると、データの90%は使用されず、単に保存されているだけです（コストが安い）。このため、IT 環境は最大 90% 無駄になっていると彼は主張します。分析されたページの 91% は Google からのトラフィックがゼロで、検索結果の 10 ページ目よりも多くの人がエベレストの頂上に到達しました。企業がこの前例のない量のデータ収集を活用しようとしているのは当然のことです。 AI は大量のデータを理解する方法を提供しますが、現在の最先端技術ではトレーニングと検証のために大量のデータが必要です。モデルの重みが増えるほど、より多くのデータが必要になります。
「AI 業界はよく石油業界と比較されます。データは、一度採掘され精製されると、石油と同様に非常に儲かる商品になります。現在、この比喩はさらに拡張される可能性があるようです。」 – MIT Technology Review
「少数ショット学習」や「ゼロショット未満学習」など、有望ではあるが新興の手法があり、ML システムが少数の例から学習して推論できるようになります。しかし、それまでの間、業界は現在、ますます大規模なデータセットに向かう傾向にあり、それには膨大な計算トレーニング予算が必要です。
人間の脳は驚くほど効率的です。単一の例から学び、その知識を生涯にわたってさまざまな状況に適用することができます。必要なエネルギー入力は比較的少なくて済みます。脳が動作するのに必要なエネルギーは 20 W だけで、人間の一生の世界平均では年間約 8,000 ポンドの CO2 が必要です。
自然は環境に対して大幅に優れた働きをしてくれました

AI よりも脳を操作するのです。これは、改善の大きな可能性を示しています。業界の最も聡明な頭脳の一部によると、「人工知能」は、言葉の意味のある意味でのインテリジェントにはまだ近づいていません。 ML システムは、多くの例を観察することによって特定のタスクを実行する方法を学習し、ブルートフォース計算アプローチとみなせるものによってデータのパターン マッチングを学習します。
これには、パターン マッチングと超人的な統計分析を実行するために、膨大な量のエネルギー (データとコンピューティング リソース) が必要です。その結果、AI への最先端のアプローチは、次のベンチマーク パフォーマンスを達成するための計算軍拡競争に巻き込まれています。カリフォルニアに本拠を置くOpenAI研究所のダリオ・アモデイ氏とダニー・ヘルナンデス氏が主導した2018年の分析では、一般人工知能が全人類に確実に利益をもたらすことをその使命としている組織で、さまざまな大規模AIトレーニングモデルで使用されるコンピューティング量が2012年以来3.4か月ごとに2倍になっていることが明らかになった。これは18か月とするムーアの法則からは大幅に逸脱しており、30万倍の増加に相当する。
出典: https://openai.com/blog/ai-and-compute/
近年、効率が大幅に向上していることは注目に値します。OpenAI によると、同等のモデル パフォーマンスの場合、コンピューティング コストは 16 か月ごとに半減しています。 2012 年以降、同じモデルで必要なコンピューティング量は 44 倍少なくなりました。しかし、モデルのパフォーマンスを少しずつ向上させるには、法外な計算コスト、経済コスト、環境コストが急速に増大しつつあります。現時点では、計算コストは​​、トレーニングに含まれるデータ、エポック、およびハイパーパラメーターの量と線形の関係があります。最先端の進歩は主にスケールによって達成されます: より大きなデータセット、より大きなデータセット

ger モデル、およびより多くのコンピューティング。
「石油資源が豊富な国が非常に高い超高層ビルを建設できるという例が最も適切だと思います。確かに、こうしたものを建設するには多大な資金と技術的努力が費やされます。そして、高層ビルの建設では確かに『最先端』が得られます。しかし…それ自体は科学の進歩ではありません。」 – ガイ・ヴァン・デン・ブロック、UCLA
現在の AI 業界は、研究者ロイ シュワルツ (アレン AI 研究所およびエルサレム ヘブライ大学) が「RedAI」と見なしているもの、つまり効率 (速度、エネルギー コスト) を犠牲にしてパフォーマンスと精度を犠牲にしたものです。
2017 年のマッキンゼーのレポートでは、ML プロジェクトの 88% が本番環境に到達しないことが示されており、多くの実験パスが行き止まりであり、それに応じて二酸化炭素排出量が発生していることが示唆されています。 ML プロジェクトが本番/公開までに成熟したとしても、最大のパフォーマンスを達成するには多くの異なるチューニング/トライアルが必要となり、最高の精度で単一のモデルを生成するには 10 ～ 20 回のトレーニング実行を超える可能性があります。
モデルのトレーニングは、総二酸化炭素排出量の一部にすぎません。ML プロジェクトのすべてのフェーズにわたって排出量を総合的に分析することが重要です。
デプロイ中、モデル推論のフットプリントは非常に大きくなります。NVIDIA は、推論が ML モデルの総炭素コストの 80 ～ 90% を占めると推定しています。これにより、ML ライフサイクル全体にわたって固着炭素を評価するための完全なライフ サイクル分析 (LCA) のための業界横断的なフレームワークが保証されます。
機械学習をより環境に優しいものにする方法はあります。この運動は、自然言語処理の研究者によって開始され、「GreenAI」と呼ばれています。このコミュニティは、中核的な指標として効率性を推進しています。一部のカンファレンス (下記) では、報告された r を生成するために使用された計算予算に関する情報を含むフォームに記入することが求められています。

結果。
GreenAI はまだ初期段階にあり、数多くの研究機会と業界パートナーシップの可能性の両方を提供しています。 ML の取り組みに可視性と説明責任を導入することで、持続可能な AI 実践を奨励するためのレポート作成と効率化の対策を優先することができます。今後の投稿では、特に有望なアプローチ (レポート作成や効率化の方法など) について詳しく説明します。それまでの間、コミュニティへの参加を開始するためのリソースをいくつか紹介します。
実験の二酸化炭素排出量の追跡を開始するための OSS パッケージ:
https://github.com/Breakend/experiment-impact-tracker
https://mlco2.github.io/impact/
持続可能なソフトウェアエンジニアリング
Microsoft 内で AI の脱炭素化を推進
GPU を使用した VLC エネルギー最適化
炭素を意識したアプリケーションと炭素効率の高いアプリケーション

## Original Extract

The carbon footprint of AI is increasing exponentially. Bigger models requiring ever more data contribute toward 'RedAI' - we need a new approach

Skip to main content
Dev Blogs
AI
All .NET posts
.NET MAUI
ASP.NET Core
Blazor
Entity Framework
C++
C#
F#
TypeScript
NuGet
Servicing
.NET Blog in Chinese
Microsoft for Developers
Agent Framework
Develop from the cloud
Xcode
ISE Developer
TypeScript
PowerShell
Python
Java
Java Blog in Chinese
Go
Microsoft Edge Dev
Microsoft 365 Developer
Microsoft Entra Identity Developer
Microsoft Entra PowerShell
Visual Studio
Visual Studio Code
Aspire
All things Azure
Azure SDK
Azure VM Runtime Team
Microsoft Azure
Azure Cosmos DB
Azure DocumentDB
Azure Data Studio
Azure SQL
DevOps
DirectX
Microsoft Foundry
Power Platform
OData
Unified Data Model (IDEAs)
Windows Command Line
#ifdef Windows
Inside MSIX
MIDI and music
React Native
The Old New Thing
Windows Developer
Dev Blogs
Artificial Intelligence (AI) has the potential to transform how we fight climate change. However, it also increasingly contributes to it: the carbon footprint of AI will grow exponentially over the next decade, and is projected to grow at a CAGR of nearly 44% globally through 2025.
The industry is trending towards bigger models (e.g. GPT-3): these require ever-growing datasets, compute budgets, and incur massive energy bills over the model lifecycle. Computational costs of AI models have been doubling every few months, resulting in an estimated 300,000x increase from 2012-2018 . In the past two years, the number of parameters have grown 170X. Currently, training a single 213M parameter NLP deep-learning model through an architecture search can generate the same carbon footprint as the lifetime of five American cars, including gas.
One NLP model has the same carbon footprint of five cars, including gas.
Source: Emma Strubell, Carnegie Mellon University
Environmental sustainability should be considered as one of the principles towards responsible development and application of AI. The benefits of using such technology should outweigh its drawbacks, and it’s time to bring the conversation about the hidden costs of AI to the forefront.
According to Gerry McGovern’s book ‘ World Wide Waste ’ , 90% of data is not used – merely stored (which is cheap). He argues that because of this, the IT landscape is ~90% waste: 91% of pages analyzed got zero traffic from google, and more people have been to the top of Everest than the 10th page of search results. It’s no surprise that companies are seeking capitalize on this unprecedented amount of data collection. AI provides a way to make sense of massive amounts of data, but the current state-of-the-art requires a massive amount of data for training & validation. The more weights a model has, the more data it needs.
“The AI industry is often compared to the oil industry: once mined & refined, data, like oil, can be a highly lucrative commodity. Now it seems the metaphor may extend even further” – MIT Technology Review
There are promising, but emergent methods such as ‘ few-shot learning ’ and ‘ less than zero-shot learning ’ which will allow ML systems to learn & reason from a handful of examples. However, in the interim, the industry is currently trending towards increasingly large datasets, which require massive computational training budgets.
The human brain is remarkably efficient; it can learn from a single example and apply this knowledge in a wide variety of contexts for the rest of its life. It requires relatively little energetic input: the brain only requires 20W to operate, and the global average of a human life requires around 8,000 lbs Co2/year .
Nature has done a significantly better job of engineering a brain than we have with AI. This represents significant potential for improvement: according to some of the brightest minds in the industry , ‘Artificial Intelligence’ isn’t yet close to being intelligent in any meaningful sense of the word. ML systems learn to perform a specific task by observing lots of examples, and pattern matching data by what can be viewed as a brute-force computational approach.
This requires staggering amounts of energy (data & compute resources) to perform pattern matching & superhuman statistical analysis. As a result, state-of-the-art approaches to AI are engaged in a computational arms race to achieve the next benchmark performance. A 2018 analysis led by Dario Amodei and Danny Hernandez of the California-based OpenAI research lab, an organization that describes its mission as ensuring that artificial general intelligence benefits all of humanity, revealed that the compute used in various large AI training models had been doubling every 3.4 months since 2012 — a wild deviation from Moore’s Law, which puts this at 18 months — accounting for a 300,000× increase .
Source: https://openai.com/blog/ai-and-compute/
It is noteworthy that there have been major efficiency gains in recent years: According to OpenAI , compute cost has been halving every 16 months for equivalent model performance. Since 2012, the same model requires 44X less compute. However, we’re rapidly approaching outrageous computational, economic, and environmental costs to gain incrementally smaller improvements in model performance. For the moment, computational cost has a linear relationship with the amount of data, epochs, and hyperparameters involved in training. The state-of-the-art advances are primarily achieved through scale: bigger datasets, larger models, and more compute.
“I think the best analogy is with some oil-rich country being able to build a very tall skyscraper. Sure, a lot of money and engineering effort goes into building these things. And you do get the ‘state of the art’ in building tall buildings. But…there is no scientific advancement per se.” – Guy Van den Broeck, UCLA
The AI industry is currently what researcher Roy Schwartz (Allen Institute For AI & Hebrew University of Jerusalem) has deemed ‘RedAI’: performance & accuracy at the expense of efficiency (speed, energy cost).
A McKinsey report from 2017 indicated that 88% of ML projects never reach production , suggesting that many experimentation paths are dead-ends with a corresponding carbon footprint. Even if a ML project does mature to production/publication, it will have required many different tunings/trials to achieve the max performance, likely exceeding 10-20 training runs to produce a single model with the highest accuracy.
Model training represents only a portion of the total carbon footprint: it is important to holistically analyze the footprint across all phases of an ML project.
While deployed, model inference has an outsized footprint: NVIDIA has estimated that inferencing constitutes 80-90% of total carbon cost of a ML model . This warrants a cross-industry framework for a full Life Cycle Analysis (LCA) to evaluate embodied carbon across the ML lifecycle.
There are ways to make machine learning greener, a movement that has been dubbed ‘GreenAI’, initiated by Natural Language Processing researchers. This community is pushing for efficiency as a core metric. Some conferences (below) now require submissions to fill out forms that include information about the computational budget used to generate the reported results.
GreenAI is in its infancy, and presents both numerous research opportunities and industry partnership potential. By bringing visibility & accountability into our ML efforts, we can begin to prioritize reporting and efficiency measures to incentivize sustainable AI practices. In future posts, I will dive into particularly promising approaches (such as reporting & efficiency methods). In the interim, here are some resources that can get you started in the community:
OSS packages to start tracking the carbon footprint of your experiments:
https://github.com/Breakend/experiment-impact-tracker
https://mlco2.github.io/impact/
Sustainable Software Engineering
Driving decarbonization of AI within Microsoft
VLC Energy Optimization With GPU
Carbon-Aware vs. Carbon-Efficient Applications

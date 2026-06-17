---
source: "https://hai.stanford.edu/news/new-approach-to-scaling-laws-could-change-how-ai-models-are-trained"
hn_url: "https://news.ycombinator.com/item?id=48567040"
title: "New Approach to Scaling Laws Could Change How AI Models Are Trained"
article_title: "New Approach to Scaling Laws Could Change How AI Models Are Trained | Stanford HAI"
author: "ilreb"
captured_at: "2026-06-17T07:58:41Z"
capture_tool: "hn-digest"
hn_id: 48567040
score: 1
comments: 0
posted_at: "2026-06-17T07:35:48Z"
tags:
  - hacker-news
  - translated
---

# New Approach to Scaling Laws Could Change How AI Models Are Trained

- HN: [48567040](https://news.ycombinator.com/item?id=48567040)
- Source: [hai.stanford.edu](https://hai.stanford.edu/news/new-approach-to-scaling-laws-could-change-how-ai-models-are-trained)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T07:35:48Z

## Translation

タイトル: スケーリング法則への新しいアプローチは AI モデルのトレーニング方法を変える可能性がある
記事のタイトル: スケーリング法への新しいアプローチは AI モデルのトレーニング方法を変える可能性がある |スタンフォードHAI
説明: AI 研究者は、計測科学と教育の統計概念を活用して、最大規模の言語モデルが将来どのようにスケールアップするかを予測するための計算需要を大幅に削減しました。トレーニング費用を数百万ドル節約できる可能性があります。

記事本文:
スタンフォード
スタンフォード大学ホーム
幹部教育および専門教育
HAI からの最新ニュース、研究の進歩、政策活動、教育プログラムの最新情報を毎週受信箱で入手できます。
ニュース スケーリング法則への新しいアプローチにより、AI モデルのトレーニング方法が変わる可能性がある
AI 研究者は、計測科学と教育の統計概念を活用することで、最大規模の言語モデルが将来どのようにスケールアップするかを予測するための計算需要を大幅に削減しました。トレーニング費用を数百万ドル節約できる可能性があります。
ChatGPT、Claude、Gemini などの大規模な言語モデルをトレーニングするのにどれくらいの費用がかかるかについて大手テック企業は口を閉ざしていますが、トレーニングの反復ごとに数億ドルから 10 億ドルかかると見積もられています。この莫大なコストは、AI 開発者が新しいモデルを 1 回だけトレーニングしたいと考えていることを意味します。
コストを抑制し、これらの大規模な単一トレーニング実行の信頼性を高めるために、開発者は、スケーリング則として知られるものに頼って、モデルを構成する多数の小さなモデルの機能を調査するようになりました。つまり、トレーニング中に言語モデルがどのようにスケールアップされるかを予測するのに役立ちます。スケーリングの法則は今や不可欠な AI インフラストラクチャとなっており、これらのスケーリング手法でも高価なコンピューティングが必要です。
現在、学者たちは、トレーニングの需要を大幅に削減し、スケーリングの時間とコストを削減する、スケーリングへの新しいアプローチを開発しました。
「スケーリングの法則が証明される前に、最も有名な開発者たちはそれに賭けてファームを賭けましたが、それがたまたまうまくいきました。彼らはモデルをどのように調整して設計するかについて大きな戦略的決定を下し、スケーリングの法則を使用してパフォーマンスを推定しましたが、彼らは正しかったのです。しかし、スケーリングは依然として高価であり、代替手段よりも低コストであっただけです」とアシスタントの Sanmi Koyejo は言います。

コンピューター サイエンスの教授であり、機械学習に関する国際会議で承認された新しい研究の上級著者であり、計算需要を 99% も削減しながらスケーリングを改善する賢い方法を紹介しています。
コエジョ研究室の大学院生で論文の筆頭著者であるサン・チュオン氏は、「私たちが研究している中心的な疑問は非常に単純です。アルゴリズムを使用してスケーリングを改善できるか?」と述べています。
Koyejo 氏、Truong 氏らは新しい論文の中で、スケーリング アルゴリズムを調整して計算需要を大幅に削減する方法を示しています。彼らはそのフレームワークをItem Response Scaling Laws（IRSL）と呼んでいます。これは、SAT などの標準化された学力評価で使用されるのと同じ概念です。
IRSL は、測定科学 (精神測定学) と教育から原則を借用して、受験者と彼らが尋ねられる質問との関係を構築し、モデルが正しく回答するにつれて質問の難易度を連続ラウンドで高めます。これにより、能力を正確に推定するために必要なクエリの数が大幅に削減される、とコエジョ氏は言います。研究者らは、IRSL がはるかに少ないクエリで同等以上の予測精度を達成し、パフォーマンスを向上させながら時間とコストを節約できることを示しています。
It’s a sort of statistical shortcut. Koyejo と Truong は、すべてのモデルについてすべての質問を複数回行うよりも、情報をより効果的かつ効率的に使用します。従来のスケーリングにおける潜在的な質問の数は 10,000 以上になることがあります。モデルの数と回答をサンプリングする必要がある回数を掛けると、スケーリング実行には 10 兆のクエリが含まれる可能性があります。一方、IRSL は、わずか 50 個の質問で同等の精度を実現します。これは 99% 以上の削減です。
「既存のフレームワークでは、多くの場合、数万の小規模なモデルを実行する必要がありました。

「私たちのアプローチは、このプロセスを劇的に効率化し、信頼性を高めます。」と Truong 氏は説明します。場合によっては、計算作業を減らすことで予測結果が向上します。」
Koyejo 氏は、IRSL の影響は学術界で最も大きくなると予測しています。学術界ではトレーニング費用が法外に高額になる可能性がありますが、潤沢な資金を持つ民間開発者も恩恵を受ける可能性があります。目標は、科学的かつ統計的に厳密な方法でスケーリングについて推論するのに役立つ新しいツールを研究者に提供することだとチュオン氏は言う。
「項目応答スケーリング法は重要な前進であると信じています」とコエジョ氏は結論づけています。 「これは、スケーリングとトレーニング全般を洗練できることを示しています。これにより、より少ない作業でより良い信号を得るという直感に反する組み合わせが得られます。」
寄稿者には、スタンフォード大学の大学院生ライラン・シェーファー氏とカリフォルニア大学ロサンゼルス校のユヘン・トゥ氏が含まれる。
この研究は、国立科学財団、ARPA-H、マッカーサー財団、シュミット サイエンス、スタンフォード人間中心人工知能研究所 (HAI)、OpenAI、マイクロソフト、Google からの資金提供によって可能になりました。
リンクがクリップボードにコピーされました!寄稿者 Andrew Myers 関連ニュース
PsychAdapter を使用すると、研究者は性格特性、年齢、精神的健康特性をダイヤルインして、実際の個人のように聞こえるテキストを生成し、トレーニング シミュレーションやパーソナライズされたコンテンツへの扉を開きます。
PsychAdapter を使用すると、研究者は性格特性、年齢、精神的健康特性をダイヤルインして、実際の個人のように聞こえるテキストを生成し、トレーニング シミュレーションやパーソナライズされたコンテンツへの扉を開きます。
新しい研究で学者らは、人気のAIチャットボットが新興ニュースに関する質問にどれだけ正確に答えられるかを測定し、地域別のチャットボットの割合がかなり高いことが判明した。

不均衡、異なる情報エコシステムへの依存、不完全なプロンプトの下での深刻な脆弱性。
新しい研究で学者らは、人気のAIチャットボットが新たなニュースに関する質問にどれだけ正確に答えられるかを測定し、大きな地域格差、異なる情報エコシステムへの依存、不完全なプロンプトの下での深刻な脆弱性を発見した。
2 つのモデルが連携すると、単独で動作する場合よりもパフォーマンスが低下し、人工知能の機能に重大なギャップがあることが明らかになります。
2 つのモデルが連携すると、単独で動作する場合よりもパフォーマンスが低下し、人工知能の機能に重大なギャップがあることが明らかになります。

## Original Extract

Leveraging statistical concepts from measurement science and education, AI researchers have greatly reduced the computational demand of predicting how the largest of large language models will scale up in the future. It could save millions of dollars in training costs.

Stanford
University Stanford Home
Executive and Professional Education
Get the latest news, advances in research, policy work, and education program updates from HAI in your inbox weekly.
news New Approach to Scaling Laws Could Change How AI Models Are Trained
Leveraging statistical concepts from measurement science and education, AI researchers have greatly reduced the computational demand of predicting how the largest of large language models will scale up in the future. It could save millions of dollars in training costs.
While Big Tech is tight-lipped on how much it costs to train large language models like ChatGPT, Claude, or Gemini, estimates range from hundreds of millions to a billion dollars for each training iteration. That steep cost means AI developers would prefer to train their new models only once.
To rein in costs and increase confidence in these massive singular training runs, developers have come to rely upon what are known as scaling laws to probe the capabilities of the many smaller models that make up their models. That is, they help predict how the language models will scale up during training. The scaling laws have now become essential AI infrastructure, and even these scaling techniques require expensive compute.
Now, scholars have developed a new approach to scaling that reduces training demands significantly, lowering the time and cost of scaling.
“Before scaling laws were proven, the best-known developers gambled and bet the farm on them, and it happened to work out for them. They made big strategic decisions about how to tweak and design their models and used scaling laws to extrapolate performance, and they were right. But scaling was still expensive, just less expensive than the alternative,” says Sanmi Koyejo , assistant professor of computer science and senior author of a new study accepted at the International Conference on Machine Learning that introduces a clever way to improve scaling while reducing computational demands as much as 99%.
“The core question we study is quite simple,” says Sang Truong , a graduate student in Koyejo’s lab and first author of the paper, “Can we use algorithms to improve scaling?”
In their new paper, Koyejo, Truong, and colleagues show how they can tailor scaling algorithms to reduce computational demand significantly. They call their framework Item Response Scaling Laws (IRSL). It is the same concept used by standardized academic assessments like the SAT.
Borrowing principles from measurement science (psychometrics) and education, IRSL builds from the relationship of test takers to the questions they are asked, increasing question difficulty with successive rounds as the model answers correctly. This significantly reduces the number of queries needed to accurately estimate ability, Koyejo says. The researchers show that IRSL achieves equal or greater predictive accuracy with far fewer queries – saving time and money while improving performance.
It’s a sort of statistical shortcut. Koyejo and Truong use information more effectively and efficiently rather than asking every question of every model multiple times. The potential questions in traditional scaling can number 10,000 or more. Multiplied by the number of models and the number of times answers must be sampled, a scaling run could include 10 trillion queries. IRSL, on the other hand, delivers equivalent accuracy using as few as 50 questions – a reduction of more than 99 percent.
“Under existing frameworks, you often had to run thousands of smaller models across tens of thousands of benchmark questions to predict outcomes,” Truong explains. “Our approach makes this process dramatically more efficient and more reliable. In some cases, doing less computational work improves predictive results.”
Koyejo predicts IRSL’s impact will be greatest in the academic world, where the costs of training can be prohibitive, but deep-pocketed private developers could benefit, too. The goal is to provide researchers new tools to help them reason about scaling in a scientific and statistically rigorous way, Truong says.
“We believe Item Response Scaling Laws is an important step forward,” Koyejo concludes. “It shows that you can refine scaling – and training in general. It gives you the counterintuitive combination of a better signal with less work.”
Contributing authors include graduate students Rylan Schaeffer at Stanford and Yuheng Tu of the University of California, Los Angeles.
This work was made possible by funding from the National Science Foundation, ARPA-H, the MacArthur Foundation, Schmidt Sciences, the Stanford Institute for Human-Centered Artificial Intelligence (HAI), OpenAI, Microsoft, and Google.
Link copied to clipboard! Contributor(s) Andrew Myers Related News
PsychAdapter lets researchers dial in on personality traits, age, and mental health characteristics to generate text that sounds like real individuals, opening the door to training simulations and personalized content.
PsychAdapter lets researchers dial in on personality traits, age, and mental health characteristics to generate text that sounds like real individuals, opening the door to training simulations and personalized content.
In a new study, scholars measured how accurately popular AI chatbots answered questions about the emerging news and found substantial regional disparity, dependence on distinct information ecosystems, and acute fragility under imperfect prompts.
In a new study, scholars measured how accurately popular AI chatbots answered questions about the emerging news and found substantial regional disparity, dependence on distinct information ecosystems, and acute fragility under imperfect prompts.
Two models working together perform worse than one alone, exposing a critical gap in artificial intelligence capabilities.
Two models working together perform worse than one alone, exposing a critical gap in artificial intelligence capabilities.

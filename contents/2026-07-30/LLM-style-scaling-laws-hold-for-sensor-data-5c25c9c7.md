---
source: "https://www.empirical.health/blog/llm-scaling-laws-hold-for-sensor-data/"
hn_url: "https://news.ycombinator.com/item?id=49116007"
title: "LLM-style scaling laws hold for sensor data"
article_title: "LLM-style scaling laws hold for sensor data | Empirical Health"
author: "brandonb"
captured_at: "2026-07-30T22:00:26Z"
capture_tool: "hn-digest"
hn_id: 49116007
score: 1
comments: 0
posted_at: "2026-07-30T21:24:47Z"
tags:
  - hacker-news
  - translated
---

# LLM-style scaling laws hold for sensor data

- HN: [49116007](https://news.ycombinator.com/item?id=49116007)
- Source: [www.empirical.health](https://www.empirical.health/blog/llm-scaling-laws-hold-for-sensor-data/)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T21:24:47Z

## Translation

タイトル: センサー データに適用される LLM スタイルのスケーリング則
記事のタイトル: LLM スタイルのスケーリング則はセンサー データに適用される |経験的健康
説明: Google

記事本文:
LLM スタイルのスケーリング則はセンサー データに適用されます。経験的健康
New: 190 ドルで 100 のバイオマーカーをサーバーアイランドで開始 LLM スタイルのスケーリング則がセンサー データに適用
ブランドン・バリンジャー · 2026年6月30日
LLM の魔法の多くは、モデル サイズ、データセット サイズ、トレーニングに使用されるコンピューティング量に応じて、損失が予測どおりにスケールされるという事実から生まれます。スケーリングの法則を当然のことと思いがちですが、それらは 2020 年に発表されたばかりであり、その構造は AI の経済学 (スケーリングの法則がなければ、フロンティアの研究機関は 1 回のトレーニングに 9 桁を投資できなかったでしょう) と AI の新たな能力 (フィリップ アンダーソンの名言「もっと多くのものは違う」が頭に浮かびます) の両方の基礎となっています。
同様のスケーリング則は、ウェアラブル基盤モデルなどの非言語基盤モデルにも適用されますか?彼らはそうしていることが分かりました。形は全く同じですか？そうではないので、いくつかの興味深い疑問が生じます。
まず、非 LLM スケーリング則の例を説明します。
Google の Scaling Wearable Foundation Models は、私の知る限り、ウェアラブルからの生理学的センサー データのスケーリング則を確立した最初の論文です。
データ サイズとモデル サイズの関数としてウェアラブル基礎モデルのパフォーマンスをスケーリングします。出典: ウェアラブル基盤モデルのスケーリング 。
検証損失 L L L は次のようにスケールされます。
ここで、 C C C は計算、 b b b はべき乗則指数、 c c c は既約フロア (詳細は後ほど) です。損失は​​数桁にわたり、両対数プロット上でほぼ直線に沿って低下し、その後、下限 c c c に向かって曲がります (計算の代わりにデータ時間やパラメーターを変更した場合にも、同じ形状が当てはまります)。 LSM は、数千時間から最大 4,000 万時間のデータに対して 4 つのモデル サイズ (2M、7M、110M、および 328M パラメーター) をテストしました。より大きなモデルとより多くのデータは、測定したすべての生成タスクに役立ちました。

ランダム代入、時間補間、センサー代入、および予測。下流のトレーニング後のタスクでも良い成果が得られました。微調整された LSM により、補間と予測がベースラインより 16 ～ 23% 向上し、アクティビティ認識が 29% 向上しました。
非 LLM スケーリング則は LLM スケーリング則と似ていますが、同一ではありません
LLM スケーリング則は、Kaplan らによって初めて確立されました。 (2020) 、その後 2022 年の Chincilla 論文で洗練されました。 Chinchilla スケーリングの法則では、固定のコンピューティング バジェットの場合、パラメーターとトークンを一緒にスケーリングする必要があります (パラメーターごとに約 20 トークン)。 Chinchilla は 1 兆 4,000 億のトークンでトレーニングされた 70B モデルで、データが不足していたそのサイズの数倍のモデルを上回りました。
チンチラ LLM のスケーリング則は次のように表されます。
ここで、 L ( N , P ) L(N, P) L ( N , P ) は検証損失です。 L ∞ L_\infty L ∞ は、既約損失の下限を表します。 a a a 、 bb b b 、 c c c 、および d d d は近似された定数 (指数と乗数) です。広く引用されている発見の 1 つは、固定のコンピューティング バジェットの下で、データとモデルのサイズを一緒にスケーリングすることによって最適な結果が達成されるということです。具体的には、コンピューティングが最適な領域とは、トレーニング トークンの数 N N N がパラメーターの数 P P P (実際には、パラメーターあたり約 20 個のトークン) に比例する状態です。
大きな違いの 1 つは、LSM のゲインが約 1,000 万時間のデータと約 1 億個のパラメーターで平坦化されたことです。 LLM は消費者規模ではそのような上限を示していません。チンチラは 1.4 兆トークンを使用しましたが、フロンティア モデルはそれをはるかに超えており、まだフラット化されていません。 (どちらのスケーリング則にも既約誤差項があるため、これは関数形式の違いではなく、むしろ経験的な結果です。)
これはスタートアップ企業にとって潜在的に興味深い機会です。 JEPA スタイルのウェアラブルをトレーニングしました

基盤モデル JETS は、4 人チームによる Google や Apple と同じ桁のデータを対象としています。したがって、別の LLM 基盤モデル会社を立ち上げるには数十億ドルの投資が必要ですが、非 LLM ドメインは実際には小規模なスタートアップ向けに開かれている可能性があります。
べき乗則は韻を踏んでいますが、基礎となる詳細の多くはかなり異なります。
これにより、いくつかの興味深い疑問が生じます。
データの壁。 LLM はデータの壁に直面しています。高品質の公開テキストの在庫はほぼ使い果たされており、合成データは不安な代替手段です。 Ilya Sutskevar が事前トレーニング終了時の講演で述べたように、「インターネットは 1 つしかありません。」生理学的データには逆の問題があります。すべての手首に装着されたすべての時計は、年間約 8,760 時間の新しい信号を受動的に、永久に生成します。センサー モデルのバインディング制約には、結果、計算、現実世界のデータの乱雑さというラベルが付けられます。 c を 0 に減らすアーキテクチャを見つけることができれば、これらの法則には実際には非常に高い上限があることになります。
市場構造。 T ロウ プライス氏は、「AI の設備投資の経済的根拠は、最終的にはスケーリングの法則に基づいている」と述べています。限界コンピューティングは限界パフォーマンスにつながる必要があります。そのため、フロンティア ラボは本質的に寡占であり、参入価格は 10 億ドル以上です。スケーリングの法則が異なる場合、競争力学が異なることを意味しますか?
同じ潜在空間ですか、それとも異なる潜在空間ですか?私たちは、ビジョン モデルを LLM の潜在空間 (CLIP など) に調整することで興味深い影響があることを確認しました。最終的にはすべてのモデルに 1 つの潜在スペースが存在するのでしょうか?それとも、生理学的モデルやその他のモデルには、言語に適合するには微妙すぎる意味のある区別を行う独自の潜在的な空間があることがわかりますか?
30 日間の心臓健康ガイドを無料で入手
心臓の健康を最適化するための証拠に基づいた手順。

心臓病によって死亡する人の数は、すべてのがんを合わせた数よりも多くなります。
それをあなたにさせないでください。
今すぐ 2,200 の検査会場の 1 つに立ち寄って、より良い心臓への旅を始めましょう
健康。
メディケアの対象となる心臓血管ケア
ニューヨーク州ニューヨークの❤️で作られています · © 2026 Empirical Health

## Original Extract

Google

LLM-style scaling laws hold for sensor data | Empirical Health
New: 100 biomarkers for $190 server-island-start LLM-style scaling laws hold for sensor data
Brandon Ballinger · Jun 30, 2026
Much of the magic of LLMs comes from the fact loss scales predictably with model size, dataset size, the amount of compute used for training. It’s easy to take scaling laws for granted, but they only published in 2020 and their structure underlies both the economics of AI (if not for scaling laws, frontier labs couldn’t invest nine figures in a training run) and AI’s emergent capabilities (the Phillip Anderson quote, “more is different”, comes to mind).
Do similar scaling laws apply to non-language foundation models, such as wearable foundation models ? It turns out they do. Is the form exactly the same? It is not, which leads to some interesting questions.
First, let me describe an example of a non-LLM scaling law.
Google’s Scaling Wearable Foundation Models was, to my knowledge, the first paper to establish a scaling law for physioloigical sensor data from wearables:
Scaling performance of a wearable foundation model as a function of data size & model size. Source: Scaling wearable foundation models .
Validation loss L L L scaled as:
where C C C is compute, b b b is the power-law exponent, and c c c is an irreducible floor (more on that later). Across multiple orders of magnitude, loss falls along a nearly straight line on the log-log plot before bending toward the floor c c c (the same shape holds when you vary data hours or parameters instead of compute). LSM tested four model sizes (2M, 7M, 110M, and 328M parameters) against data from a few thousand hours up to 40 million. Bigger models and more data both helped on every generative task they measured: random imputation, temporal interpolation, sensor imputation, and forecasting. The payoff on downstream, post-trained tasks was good too. Fine-tuned LSM improved interpolation and forecasting by 16-23% over baselines and lifted activity recognition by 29%.
Non-LLM scaling laws are similar, but not identical to LLM scaling laws
LLM scaling laws were first established in Kaplan et al. (2020) , and then refined in the 2022 Chincilla paper . In the Chinchilla scaling laws, for a fixed compute budget, you should scale parameters and tokens together, about 20 tokens per parameter. Chinchilla was a 70B model trained on 1.4 trillion tokens, and it beat models several times its size that had been starved of data.
The Chinchilla LLM scaling laws are expressed as:
Here, L ( N , P ) L(N, P) L ( N , P ) is the validation loss; L ∞ L_\infty L ∞ ​ represents the irreducible loss floor; and a a a , b b b , c c c , and d d d are fitted constants (exponents and multipliers). One widely cited finding is that, under a fixed compute budget, optimal results are achieved by scaling data and model size together: specifically, the compute-optimal regime is where the number of training tokens N N N is proportional to the number of parameters P P P (in practice, about 20 tokens per parameter).
One major difference is that LSM’s gains flattened out around 10 million hours of data and roughly 100 million parameters. LLMs have shown no such ceiling at consumer scale. Chinchilla used 1.4 trillion tokens and frontier models have gone well past that, with no flattening yet. (Both scaling laws have an irreducible error term, so this isn’t a difference in functional form but rather an empirical result.)
That’s a potentially interesting opportunity for startups. We trained a JEPA-style wearable foundation model, JETS , on the same order of magnitude of data as Google and Apple with a four-person team. So whereas starting another LLM foundation model company requires billions of dollars of investment, non-LLM domains might actually be open for smaller startups.
While the power laws rhyme, many of the underlying details are pretty different:
This leads to several interesting questions:
Data wall. LLMs are running into a data wall, where the stock of high-quality public text is close to spent and synthetic data is an uneasy substitute. As Ilya Sutskevar put it in his talk on the end of pretraining, “we have but one internet.” Physiological data has the opposite problem. Every watch on every wrist generates roughly 8,760 hours of new signal a year, passively, forever. The binding constraints for sensor models are labeled outcomes, compute, and the messiness of real-world data. If we can find architectures that reduce c to 0, there’s actually a very high ceiling on these laws.
Market structure. T Rowe Price put it, “the economic rationale for AI capital expenditure ultimately rests on scaling laws.” Marginal compute must lead to marginal performance, which is why frontier labs are essentially an oligopoly with an entry price of $1B+. If the scaling laws are different, does this mean the compeitive dynamics are different?
Same or different latent space? We’ve seen interesting implications from aligning vision models into the LLM’s latent space (e.g., CLIP). Will there ultimately be one latent space for all models? Or will we see physiological and other models have their own latent spaces that make meaningful distinctions that are too subtle to fit in language?
Get your free 30-day heart health guide
Evidence-based steps to optimize your heart health.
Heart disease kills more people than all cancers combined.
Don't let it be you.
Stop by one of 2,200 testing sites today and start your journey to better heart
health.
Medicare-covered cardiovascular care
Made with ❤️ in New York, NY · © 2026 Empirical Health

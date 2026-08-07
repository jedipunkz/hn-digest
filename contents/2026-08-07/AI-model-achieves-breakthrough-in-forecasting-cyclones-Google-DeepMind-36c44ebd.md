---
source: "https://deepmind.google/blog/weathernext-ai-model-achieves-breakthrough-in-forecasting-cyclones/"
hn_url: "https://news.ycombinator.com/item?id=49207172"
title: "AI model achieves breakthrough in forecasting cyclones – Google DeepMind"
article_title: "AI model achieves breakthrough in forecasting cyclones — Google DeepMind"
author: "rbanffy"
captured_at: "2026-08-07T08:06:59Z"
capture_tool: "hn-digest"
hn_id: 49207172
score: 1
comments: 0
posted_at: "2026-08-07T07:49:19Z"
tags:
  - hacker-news
  - translated
---

# AI model achieves breakthrough in forecasting cyclones – Google DeepMind

- HN: [49207172](https://news.ycombinator.com/item?id=49207172)
- Source: [deepmind.google](https://deepmind.google/blog/weathernext-ai-model-achieves-breakthrough-in-forecasting-cyclones/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T07:49:19Z

## Translation

タイトル: AI モデルがサイクロン予測で画期的な進歩を達成 – Google DeepMind
記事タイトル: AI モデルがサイクロン予測で画期的な進歩を達成 — Google DeepMind

記事本文:
メイン コンテンツにスキップ 次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
次世代 AI システムを探索する
AI の最新のブレークスルーとラボからの最新情報
AI で新たな発見の時代を切り開く
私たちの使命は、人類に利益をもたらすために責任を持って AI を構築することです
Google DeepMind Google AI すべての AI について学ぶ Google DeepMind AI のフロンティアを探る Google Labs AI 実験を試す Google Research 研究を探る 製品とアプリ Gemini アプリ Gemini とチャットする Google AI Studio 次世代 AI モデルで構築する Google Antigravity 当社のエージェント開発プラットフォーム モデル 研究 科学 Gemini で構築する Gemini を試す 2026 年 8 月 6 日 Science WeatherNext: AI モデルがサイクロンの予測でブレークスルーを達成
WeatherNext を使用すると、正確なサイクロン予測が可能になり、追加の 1 日の警告を発することができます。現在、モデルをオープンソース化しています。
サイクロンがどのように危険に発達するかを予測することは、一刻を争う長年の課題です。ハリケーンまたは台風としても知られる熱帯低気圧は、地球上で最も破壊的な気象現象の 1 つであり、過去 50 年間で世界中で 70 万人以上の死者と 1 兆 4,000 億ドルの経済損失を引き起こしています。気象予報士にとって、タイムリーで正確な警告を発することは、常に時間との闘いです。
本日、Nature に掲載された論文で、WeatherNext AI モデルがサイクロンの進路、強さ、風の構造を予測する際に最先端の精度を達成したことを示しました。平均して、私たちのモデルは予報担当者に 1 日分の予測精度をもたらします。私たちの 3 日間の予測は、以前のモデルが次の 2 日間のみ提供できたものと同等です。 T

彼の改善の規模は、およそ 10 年分の気象学の進歩に相当します。
この共同作業には、Google DeepMind と Google Research の AI 研究者とエンジニア、国立ハリケーン センター (NHC)、大気圏研究協力研究所 (CIRA)、英国気象庁、および世界中の気象機関の専門予報官が結集しました。
私たちの研究はすでに現実世界に影響を与えています。 2025 年のハリケーン シーズン中、私たちのモデルは、嵐の急速な激化とジャマイカへの上陸を予測することで、NHC がハリケーン メリッサの歴史的な予測を立てるのに役立ちました。これにより、NHC は事前に警告を発することができ、地上のチームに準備のための重要な時間を与えることができました。今年も私たちは協力を続け、予報担当者の意思決定を支援するために、各サイクロンについて考えられる 1,000 のシナリオを予測しています。
天気は誰にでも影響を与えます。この広範な影響を考慮して、私たちは現在、ハリケーンシーズン中に使用される WeatherNext 2 および WeatherNext Cyclones モデルをオープンソース化しています。このテクノロジーをオープンに利用できるようにすることで、研究コミュニティに力を与え、地域の予報官に自然災害に備えるために必要なツールを提供したり、再生可能エネルギーの成長を支援したり、異常気象を予測したりするなど、より回復力のあるコミュニティの構築における AI の影響を拡大したいと考えています。
WeatherNext が天気と低気圧を予測する方法
WeatherNext Cyclones は、ハリケーン ミルトン (2024 年 10 月) の際の世界的な大気状況から出発して、世界的な気象パターンと詳細なサイクロンの進路の両方を最大 15 日前まで繰り返し予測します。 1,000 人のメンバーからなるアンサンブルを実行すると、熱帯低気圧からハリケーン級の風までの局所的な確率マップが生成されます。
サイクロンを予測すると通常はトレードオフが必要になる

ff には 2 つの異なるモデリング技術が必要です。サイクロンの進路（行き先）は、大規模な全地球規模の大気流によって導かれますが、これまでは、より粗い全球モデルによってモデル化するのが最適でした。ただし、サイクロンの強さ (どの程度強くなるか) は、その中心付近の高度に局所化された微細スケールの熱力学物理プロセスによって決まります。これらのプロセスは、専門化された高解像度の局所モデルによって最適にモデル化されます。
当社の WeatherNext モデルは、サイクロンだけでなく地球規模の天気全体の予測を改善することで、このギャップを埋めます。これは、熱帯低気圧の進路、強さ、風の構造を最先端の精度で予測する単一の AI モデルです。この画期的なトレーニング、アーキテクチャ、および低解像度入力へのアプローチを独自に組み合わせることによって、この画期的な成果を達成しました。
私たちは、2023 年から 2024 年までの歴史的なサイクロンで WeatherNext Cyclones を評価し、他のトップ気象モデルと比較してその決定論的および確率論的なパフォーマンスをベンチマークしました。平均して、WeatherNext Cyclones は、サイクロンの進路、強さ、風の構造を予測するために丸 1 日 (24 時間) 以上のリードタイムの​​利点を獲得します。
このモデルは、地球規模の気象力学と専門家が厳選した過去のサイクロン観測という 2 つの異なるデータ モダリティで共同トレーニングされました。約 20 テラバイトの地球規模の大気データと、過去 5,000 件近くの嵐に及ぶ過去の IBTrACS データベースを使用してエンドツーエンドでトレーニングすることにより、モデルは複雑な大気パターンと異常気象をモデル化する方法を学習します。
サイクロンの予測精度はここ数十年で着実に進歩しています。これらのプロットは、長年にわたる ECMWF-ENS 進路予報 (a) と HWRF 強度予報 (b) の 3 日間の精度と、WeatherNext Cyclones が進路と強度の両方の精度の段階的な変化にどのように寄与しているかを示しています。この改善は同等です

過去 20 年間の傾向に応じて 10 年の進歩を遂げます。
私たちのモデルは、機能生成ネットワーク (FGN) を使用して、さまざまな予測のアンサンブルを効率的に生成し、天気に固有の不確実性を捉えます。 TPU を使用して 1 分以内に単一の 15 日間の予測を生成できるようになり、予測担当者が潜在的に壊滅的なテールリスクの確率分布を迅速に評価できるようになりました。昨年、私たちのシステムは一度に 50 件の予測を生成し、地球規模の物理モデルと一致しました。今年、私たちはアンサンブルの規模をメンバー 1,000 人に拡大し、2025 年のハリケーン メリッサの際に発生した、急速な激化イベントのようなまれではあるが重大なシナリオを捉えました。
これまで、非常に高い空間分解能で動作することが、正確な強度予測を行うための主な推進力であると考えられてきました。ただし、WeatherNext Cyclones に必要なのは、従来のモデルよりも 100 倍粗い 28x28km の解像度のデータのみです。このモデルの小型バージョンである WeatherNext 2-mini は、より粗い 111x111km 解像度で動作し、優れたパフォーマンスを示します。これは科学者を驚かせましたが、私たちのモデルがどのようにしてこの解像度でこれほど正確な予測を生成するのかを完全に理解することは未解決の研究課題のままです。研究コミュニティと協力して解明できることを願っています。
WeatherNext を研究コミュニティに開放する
Nature 論文と並行して、コードとモデルの重みをオープンソース化し、誰でも自由に構築できるようにしています。これには、学術研究、運用予測、またはより専門化されたローカライズされたモデルの開発が含まれます。私たちは、世界の気象コミュニティ全体の進歩を加速し、気象機関、研究者、非営利団体があらゆる種類の気象現象をより適切に予測し、それを守るための重要な決定を下せるようにしたいと考えています。

CTの生活とインフラ。
また、同様のモデルを 2 セットリリースしています。WeatherNext Cyclones はハリケーンシーズン中に実行されました (結果は論文でご覧いただけます)。 WeatherNext 2 は、10 月に運用開始した後のアップデートです。さらに、無料のパブリック Colab ノートブックの単一の TPU で実行できるモデルのコンパクト バージョンである WeatherNext 2-mini をリリースします。
最新のサイクロン予報は Weather Lab で確認できます。Weather Lab は最近新しいインターフェイスで更新され、サイクロンの進路とともに世界の天気予報を含むように拡張されました。 Weather Lab では、気温、降水量、風速などに関する WeatherNext の予測をすべて 1 つのビューで視覚化できるようになりました。 Weather Lab モデルと WeatherNext モデルはどちらも Google Earth AI の一部です。
天気予報における AI の最前線を開拓する
私たちは、サイクロンを予測するためのリードタイムを丸 1 日以上獲得することで、歴史的な進歩を達成しました。これは、気象学の進歩の 10 年に匹敵する進歩をもたらしました。将来の嵐の季節に備えて、研究者、気象機関、専門家に協力してもらい、オープンソース モデルを構築し、Weather Lab で予測を調査してもらいます。高度な機械学習と人間の予報士が持つ不可欠な現実世界の専門知識を組み合わせることにより、私たちは、命を救い、コミュニティが気候の変化に適応できるよう支援できる、協調的な天気予報エコシステムを構築することを目指しています。
注: 公式の天気予報と警報については、地元の気象機関または国の気象局に問い合わせてください。
この研究は、Google DeepMind チームと Google Research チームによって共同開発されました。
私たちの協力者である NOAA/NWS/NCEP 国立ハリケーン センター、大気研究協力研究所 (CIRA)、および

英国気象庁のパートナーシップとこの論文への貢献に感謝します。
この作品は、論文の共著者の貢献を反映しています: Ferran Alet、Tom Andersson、Ilan Price、Stratis Markou、Andrew El-Kadi、Dominic Masters、Amy Li、Samier Merchant、Natalie Williams、Gregory Thornton、Ken MacKay、Olivia Graham、Akib Uddin、Ben Gaiarin、Devaja Shah、Elinor Kruse、Wallace Hogsett、Davidゼリンスキー、ジョン・カンギアロシ、ジョナサン・マルティネス、ジェームズ・フランクリン、マーク・デマリア、ケイト・マスグレイブ、キャロライン・L・ベイン、ヘレン・タイトリー、ジャックリン・ストット、レミ・ラム、アーロン・ベル、ポール・コマレク、マシュー・ウィルソン、アルバロ・サンチェス＝ゴンザレス、ピーター・バタグリア。
国立ハリケーン センターがハリケーン メリッサの歴史的なジャマイカ上陸を予測するのにウェザーネクストがどのように役立ったか
GenCast は、最先端の精度で天気と極端な状況のリスクを予測します
GraphCast: より高速かつ正確な地球規模の天気予報を実現する AI モデル
フォローしてください サインアップして、最新のイノベーションに関する最新情報を入手してください 私は Google の利用規約に同意し、私の情報が Google のプライバシー ポリシーに従って使用されることを認めます。

## Original Extract

Skip to main content Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Explore our next generation AI systems
Our latest AI breakthroughs and updates from the lab
Unlocking a new era of discovery with AI
Our mission is to build AI responsibly to benefit humanity
Google DeepMind Google AI Learn about all our AI Google DeepMind Explore the frontier of AI Google Labs Try our AI experiments Google Research Explore our research Products and apps Gemini app Chat with Gemini Google AI Studio Build with our next-gen AI models Google Antigravity Our agentic development platform Models Research Science About Build with Gemini Try Gemini August 6, 2026 Science WeatherNext: AI model achieves breakthrough in forecasting cyclones
Share Copied WeatherNext enables accurate cyclone forecasts that can give an extra day of warning. Now we are open sourcing the model.
Predicting how dangerous cyclones develop is a longstanding challenge where every hour counts. Tropical cyclones — also known as hurricanes or typhoons — are among the most destructive weather phenomena on Earth, responsible for more than 700,000 deaths and $1.4 trillion in economic losses globally over the past 50 years. For forecasters, issuing timely, accurate warnings is a constant race against time.
Today, in a paper published in Nature , we show that our WeatherNext AI model achieved state-of-the-art accuracy in predicting a cyclone's track, intensity, and wind structure. On average, our model gives forecasters an extra day’s worth of predictive accuracy: our three-day forecasts are as good as what prior models were able to provide for only the next two days. This scale of improvement corresponds roughly to a decade’s worth of meteorological progress.
This collaborative work brought together AI researchers and engineers at Google DeepMind and Google Research, with expert forecasters at the National Hurricane Center (NHC), the Cooperative Institute for Research in the Atmosphere (CIRA), the UK Met Office , and weather agencies around the world.
Our research has already had real-world impact. During the 2025 hurricane season, our model helped the NHC to make a historic forecast for Hurricane Melissa by predicting the storm’s rapid intensification and landfall in Jamaica. This enabled the NHC to issue an advance warning, giving teams on the ground critical time to prepare. This year, we continue to work together and are now predicting 1,000 possible scenarios for each cyclone to help support forecasters in their decision-making.
Weather affects everyone. Given this broad impact, we are now open sourcing our WeatherNext 2 and WeatherNext Cyclones models used during the hurricane season. By making this technology openly available, we hope to empower the research community and amplify AI's impact in building more resilient communities – whether that be providing local forecasters with the tools they need to prepare for natural disasters , supporting the growth of renewable energy, or anticipating extreme weather.
How WeatherNext predicts weather and cyclones
Starting from global atmospheric conditions during Hurricane Milton (October 2024), WeatherNext Cyclones iteratively predicts both global weather patterns and fine-scale cyclone tracks up to 15 days in advance. Running a 1,000-member ensemble generates localised probability maps of tropical storm to hurricane-force winds.
Predicting cyclones has typically forced a trade-off requiring two distinct modeling techniques. A cyclone's track (where it goes) is steered by massive, global atmospheric currents, which before now have been best modeled by coarser global models. However, a cyclone’s intensity (how strong it gets) is driven by highly localized, fine-scale thermodynamic physical processes around its core, which are best modeled by specialized, higher resolution, local models.
Our WeatherNext model bridges this gap by improving forecasting for global weather overall as well as cyclones. It is a single AI model that predicts a tropical cyclone’s track, intensity, and wind structure with state-of-the-art accuracy. It achieves this breakthrough through a unique combination of its training, architecture and approach to low resolution inputs.
We evaluated WeatherNext Cyclones on historical cyclones from 2023 to 2024, benchmarking its deterministic and probabilistic performance against other top weather models. On average, WeatherNext Cyclones gains more than a full day (24 hours) of lead time advantage for predicting cyclone tracks, intensity, and wind structure.
The model was co-trained on two distinct data modalities: global weather dynamics and expert-curated historical cyclone observations. By training end-to-end on nearly 20 terabytes of global atmospheric data and the historical IBTrACS database spanning nearly 5,000 historical storms, the model learns complex atmospheric patterns and how to model extreme weather.
Cyclone forecast accuracy has been steadily advancing over recent decades. The plots show the 3-day accuracy of ECMWF-ENS track forecasts (a) and HWRF intensity forecasts (b) over the years, and how WeatherNext Cyclones contributes a step change in accuracy for both track and intensity. This improvement is the equivalent to a one-decade progress according to trends over the last 20 years.
Our model uses Functional Generative Networks (FGNs) to efficiently produce ensembles of different predictions, which captures the inherent uncertainty of the weather. We can now generate a single 15-day forecast in less than a minute on a TPU, empowering forecasters to quickly evaluate the probability distribution of potentially devastating tail-risks. Last year, our system produced 50 predictions at a time, matching global physics models. This year we scaled our ensemble size to 1,000 members, capturing rare but consequential scenarios like rapid intensification events, as occurred during Hurricane Melissa in 2025.
Up until now, operating at very high spatial resolution has been considered the main driver for making accurate intensity forecasts. However, WeatherNext Cyclones only needs data with a resolution of 28x28km, 100x coarser than traditional models. A smaller version of the model, WeatherNext 2-mini, which operates at a coarser 111x111km resolution, also shows great performance. This has surprised scientists, and it remains an open research question to fully understand how our models produce such accurate predictions at this resolution. We hope that, together with the research community, we can find out.
Opening up WeatherNext to the research community
Alongside our Nature paper, we are open sourcing the code and model weights, making them freely available for anyone to build on. This includes academic research, operational forecasting, or developing more specialized, localized models. We hope to accelerate progress across the global weather community and empower meteorological agencies, researchers, and nonprofits to better predict weather events of all kinds and make key decisions to protect lives and infrastructure.
We are also releasing two sets of similar models: WeatherNext Cyclones, which ran during the hurricane season (results can be seen in the paper); and WeatherNext 2, a later update that we operationalized in October. Additionally, we are releasing WeatherNext 2-mini, a compact version of the model that can run on a single TPU in a free public Colab notebook .
You can explore our latest cyclone forecasts on Weather Lab , which we recently refreshed with a new interface and expanded to include global weather forecasts alongside cyclone tracks. Weather Lab now lets you visualize WeatherNext predictions for temperature, precipitation, wind speed, and more, all in a single view. Both Weather Lab and WeatherNext models are a part of Google Earth AI .
Pushing the frontiers of AI for weather forecasting
We have achieved a historic breakthrough by gaining more than a full day of lead time for predicting cyclones — delivering an advance equivalent to a decade of meteorological progress. As we prepare for future storm seasons, we invite researchers, meteorological agencies, and experts to partner with us, build on our open source models, and explore our forecasts on Weather Lab. By combining advanced machine learning with the indispensable real-world expertise of human forecasters, we aim to create a collaborative weather forecasting ecosystem that can save lives and help communities adapt to a changing climate.
Note: For official weather forecasts and warnings, refer to your local meteorological agency or national weather service.
This research was co-developed by Google DeepMind and Google Research teams.
We’d like to thank our collaborators NOAA/NWS/NCEP National Hurricane Center, Cooperative Institute for Research in the Atmosphere (CIRA) and the UK Met Office for their partnership and contributions to the paper.
This work reflects the contributions of the paper’s co-authors: Ferran Alet, Tom Andersson, Ilan Price, Stratis Markou, Andrew El-Kadi, Dominic Masters, Amy Li, Samier Merchant, Natalie Williams,Gregory Thornton, Ken MacKay, Olivia Graham, Akib Uddin, Ben Gaiarin, Devaja Shah, Elinor Kruse, Wallace Hogsett, David Zelinsky, John Cangialosi, Jonathan Martinez, James Franklin, Mark DeMaria, Kate Musgrave, Caroline L. Bain, Helen Titley, Jacklynn Stott, Remi Lam, Aaron Bell, Paul Komarek, Matthew Willson, Alvaro Sanchez-Gonzalez, and Peter Battaglia.
How WeatherNext helped the National Hurricane Center better predict Hurricane Melissa’s historic landfall in Jamaica
GenCast predicts weather and the risks of extreme conditions with state-of-the-art accuracy
GraphCast: AI model for faster and more accurate global weather forecasting
Follow us Sign up for updates on our latest innovations I accept Google's Terms and Conditions and acknowledge that my information will be used in accordance with Google's Privacy Policy .

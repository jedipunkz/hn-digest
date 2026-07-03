---
source: "https://openllmstack.com/blog/open-source-llm-statistics/"
hn_url: "https://news.ycombinator.com/item?id=48769528"
title: "Open Source LLM Statistics and Trends (2026)"
article_title: "30+ Open Source LLM Statistics & Trends (2026) — OpenLLMStack"
author: "sherlockxu"
captured_at: "2026-07-03T01:58:47Z"
capture_tool: "hn-digest"
hn_id: 48769528
score: 1
comments: 0
posted_at: "2026-07-03T01:24:00Z"
tags:
  - hacker-news
  - translated
---

# Open Source LLM Statistics and Trends (2026)

- HN: [48769528](https://news.ycombinator.com/item?id=48769528)
- Source: [openllmstack.com](https://openllmstack.com/blog/open-source-llm-statistics/)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T01:24:00Z

## Translation

タイトル: オープンソース LLM の統計と傾向 (2026 年)
記事のタイトル: 30 を超えるオープンソース LLM の統計と傾向 (2026) — OpenLLMStack
説明: エコシステムの成長、モデルの使用状況、パフォーマンス、推論コスト、主要な開発者をカバーする、2026 年の 30 以上のオープンソース LLM 統計を調査します。

記事本文:
30 以上のオープンソース LLM 統計と傾向 (2026) — OpenLLMStack OpenLLMStack モデル エンジン 最適化エージェント タイムライン ブログ モデル エンジン 最適化エージェント タイムライン ブログ
GitHub
ブログに戻る
30 を超えるオープンソース LLM の統計と傾向 (2026 年)
エコシステムの成長、モデルの使用状況、パフォーマンス、推論コスト、主要な開発者をカバーする、2026 年の 30 以上のオープンソース LLM 統計を調査します。
オープンソース LLM は、3 年も経たないうちに、研究目的から実際の実稼働システムのバックボーンに成長しました。これらは現在、コーディング アシスタント、エージェント、エンタープライズ パイプラインを強化しており、一部の分野ではブームの初期を定義した独自のモデルを追い越しています。
それでは、今日のオープンソース LLM エコシステムはどれくらいの規模になっているのでしょうか?そして実際にどれくらいのスピードで成長しているのでしょうか？
私は、Stanford AI Index、Hugging Face、Meta、OpenRouter、Stack Overflow Developer Survey などの一次ソースから、私が見つけた中で最も有用なオープンソース LLM 統計をまとめました。各セクションは元の情報源にリンクしているため、直接引用できます。
上位のオープンソース LLM 統計
この記事では、ファーストパーティのレポート、研究論文、公式ドキュメント、および元のデータセットを優先します。
オープンソース LLM は一般に、ユーザーがダウンロード、実行、変更できる言語モデルを指します。 API を介してのみアクセスできる独自のモデルと比較して、オープンソース LLM を使用すると、開発者は展開、カスタマイズ、インフラストラクチャをより詳細に制御できます。
オープンソースという用語は、多くの場合、曖昧に使用されます。オープンソースとして説明されているモデルの多くは、実際にはオープンウェイト モデルとしてリリースされています。それらの重みは公開されていますが、ライセンスには従来のオープンソース ソフトウェア ライセンスとは異なる制限が含まれる場合があります。業界では一般的に「オープンソース LLM」を参照するために使用されているためです。

どちらのカテゴリーについても、この記事は簡潔にするためにその規則に従います。
オープンソース LLM はいくつありますか?
オープンソース LLM の権威ある世界的な数は存在しません。 Hugging Face は、2025 年に 200 万を超えるパブリック モデル リポジトリをホストしましたが、その合計には、テキスト、画像、オーディオ、ロボット工学、その他のタスクのモデルに加え、微調整、アダプター、量子化、および派生関数が含まれています。
より広範な Hugging Face エコシステムは、オープン モデル開発の規模と方向性を測定するのに依然として役立ちます。
ダウンロードされたモデルの平均サイズは 2023 年から 2025 年の間に約 25 倍に増加しましたが、中央値は約 25% しか増加しませんでした。この違いは、比較的少数の大型モデルが平均を押し上げている一方で、小型モデルが依然として一般的であることを示唆しています。
オープンソース LLM の導入と使用
オープン モデルは実験をはるかに超えています。 2025 年後半までに、開発者が何百もの AI モデルにアクセスできるようにする統合 API プラットフォームである OpenRouter 上の全トークン量の約 3 分の 1 に達しました。基礎となる調査では、2024 年 11 月から 2025 年 11 月までの 1 年間で 100 兆を超えるトークンが分析されました。
DeepSeek は、調査対象の 2 番目に大きいオープンソース ファミリである Qwen の約 2.6 倍のトークンを処理しました。しかし、2025 年後半までに、オープンソース モデル トークンの 20% ～ 25% 以上を一貫して占める個別モデルはなくなりました。使用は 5 ～ 7 の競合モデルに広がっていました。
中国で開発されたモデルは、2024 年後半には週間トークン量のわずか 1.2% でしたが、数週間では 30% 近くまで増加しました。
中国のオープンソース モデルは、全期間を通じて週間 OpenRouter トークン量の平均 13.0% を占めました。
中国国外で開発されたオープンソース モデルは平均 13.7% でした。
中国国外で開発された独自モデルは平均約 70% のシェアを維持しました。
ロールプレイはオープンの約 52% を占めました

ソース モデル トークンが 2 番目に大きいカテゴリーはプログラミングで、およそ 15% ～ 20% でした。
OpenRouter 自体は調査期間後も成長を続けました。 Menlo Ventures の報告によると、このプラットフォームの開発者は約 1 年で 250 万人から 800 万人以上に増加し、2026 年 5 月には年間実行率が約 150 兆トークンに達しました。
開発者による AI ツールの導入は、今やほぼ普遍的になりました。 2025 年の Stack Overflow Developer Survey の回答者の 84% が AI ツールを使用している、または使用する予定であり、前年の 76% から増加し、プロの開発者の 51% が毎日 AI ツールを使用しています。
出典: OpenRouter 、 Menlo Ventures 、 スタックオーバーフロー
最も人気のあるオープンソース LLM
人気は指標に依存します。ダウンロードは配布を測定し、派生リポジトリは開発者がモデルに基づいて構築する頻度を測定し、ルーティングされたトークンはホストされた API の使用量を測定します。
Meta は、Llama ファミリーが 2025 年 3 月に累積ダウンロード数 10 億を超えたと報告しました。しかし、Hugging Face では、Qwen が派生作品の最大のエコシステムになりました。このファミリーは、2026 年 3 月までに 113,000 を超える直接派生モデルを持ち、Qwen タグ付きモデルをすべて含めると 200,000 を超えるリポジトリを持ちました。
これらの数字はダウンロードとリポジトリをカウントしており、ユニーク ユーザーや運用環境の数はカウントしていません。自動ダウンロード、ミラーリング、量子化、および繰り返しのプルはすべて、合計に影響を与える可能性があります。
OpenLLMStack モデル ページで、現在のモデル リリース、ライセンス、コンテキスト ウィンドウ、およびハードウェア要件を比較できます。
オープンモデルとクローズドモデルのパフォーマンス
スタンフォード AI インデックスによると、2026 年 3 月のアリーナ リーダーボードでは、主要なオープン モデルが主要なクローズド モデルに 3.3% 遅れました。 2025年中に再開されるまで、2024年8月時点ではその差はわずか0.5％だった。
正しい結論は、オープン モデルが永久に達成されたということではありません。

d パリティ。この差は競争力を維持できるほど小さいですが、新しい世代のモデルが登場するにつれて変化します。アリーナの上位 10 モデルのうち 6 モデルは、2026 年 3 月時点で販売を終了しました。
出典: スタンフォード AI インデックス 2026
オープンソース モデルの由来
中国で開発されたモデルは、2025 年のハグ フェイス ダウンロード数の 41% を占め、単一国による最大のシェアを占めました。この年、中国は月間ダウンロード数と全体ダウンロード数の両方で米国を上回りました。
同時に投稿者の構成も変更されました。業界によるモデル開発の割合は、2022 年以前の約 70% から 2025 年の 37% に低下しましたが、関連のない開発者によるダウンロード数は 17% から 39% に増加しました。
OpenRouter も同様の地理的変化を記録しましたが、測定基準は異なりました。中国のオープンソースモデルは、2024年後半の週間トークン使用量の1.2％から、2025年の一部の週には30％近くまで上昇した。
出典: ハグフェイス 、 OpenRouter
LLM 推論のコスト低下
Andreessen Horowitz の分析では、固定レベルの MMLU パフォーマンスでの推論の価格が年間約 10 分の 1 に低下することがわかりました。 MMLU スコアが 42 の場合、観測された最低価格は 2021 年の 100 万トークンあたり 60 ドルから 2024 年には 0.06 ドルに低下し、1,000 倍の下落となりました。
この過去の推定には限界があります。MMLU に依存し、入力と出力の価格を平均し、OpenAI、Anthropic、Meta から選択されたモデルをカバーしています。これはコストの傾向を示すものとして有用であり、毎年さらに 10 倍の減少を保証する法則ではありません。
現在の API の価格設定は、スプレッドがどれほど大きくなる可能性があるかを示しています。 2026 年 7 月 2 日現在:
価格は標準レートでの 100 万トークンあたりの料金です。これらの公開価格では、GPT-5.5 よりも DeepSeek-V4-Flash 入力は 97.2% 安く、出力は 99.1% 安くなります。
サービスの効率性は、安価なハードウェアだけでは実現できません。量子化、バッチ処理、キャッシング、スパース アーキテ

アーキテクチャと最適化されたカーネルがすべて貢献します。 OpenLLMStack は、推論最適化ページ上の主要なテクニックと、それらを推論ディレクトリに実装するエンジンを追跡します。
出典: Andreessen Horowitz 、 DeepSeek API の価格 、 OpenAI API の価格
トップのオープンソース LLM API プロバイダー
オープン モデルを実行するために独自の GPU は必要ありません。サーバーレス推論プロバイダーのセットが増加しており、単一 API の背後で主要なオープン ウェイトをホストしており、トークンごとに課金されるため、インフラストラクチャを管理せずにモデルを切り替えることができます。
単一モデルに最も安価にアクセスするには、DeepSeek や Alibaba の独自モデル用の API など、モデル メーカーのファーストパーティ API が最も低価格のオプションとなることがよくあります。 OpenRouter のようなアグリゲーターは、1 つのリクエストをこれらのプロバイダーの多くにルーティングするため、価格と速度を比較できます。
DeepSeek R1: オープンソース LLM の躍進の瞬間
DeepSeek R1 ほど、オープンソース LLM のプロファイルに大きな効果をもたらしたリリースはありません。このモデルは 2025 年 1 月 20 日に発売され、数日以内に DeepSeek アプリは米国の Apple App Store で第 1 位に上り詰め、ChatGPT を追い越し、50 か国以上でチャートのトップになりました。
ダウンロードの急増はほぼ垂直でした。このアプリは、発売後の月曜日までに App Store と Google Play 全体で 260 万ダウンロードに達し、全ダウンロードの 80% 以上が過去 7 日間に行われ、Appfigures のデータではこのアプリが世界第 1 位にランクされました。
市場の反応も同様に劇的でした。 2025 年 1 月 27 日、NVIDIA は市場価値で約 5,890 億ドルを失い、これは企業としては史上最大の 1 日の損失となりました。DeepSeek がフロンティア グレードのオープン モデルを約 560 万ドルでトレーニングできることが示されたと伝えられています。
これらのマイルストーンは、OpenLLMStack タイムラインで追跡できます。
出典: TechCrunch 、 ブルームバーグ 、 フォーブス
権威あるものではない

rganization は、オープンソース LLM の世界的な数を管理しています。 Hugging Face は 2025 年に 200 万以上のパブリック モデル リポジトリをホストしましたが、その数字はすべてのモデル タイプをカバーしており、派生モデル、アダプター、量子化も含まれています。これは一意の LLM の数ではありません。
Meta は、2025 年 3 月の Llama の累計ダウンロード数が 10 億件を超えたと報告しました。
Qwen には、2026 年 3 月に Hugging Face によって報告された最大のデリバティブ エコシステムがあり、113,000 を超える直接デリバティブ モデルがありました。
DeepSeek は、2024 年 11 月から 2025 年 11 月までに 14 兆 3,700 億のトークンを処理し、OpenRouter でのオープンソースの使用を主導しました。
2026 年時点で、単一の最高のオープンソース LLM は存在しません。コーディングに優れたモデルもあれば、推論、長いコンテキストの処理、または多言語タスクに優れたモデルもあります。適切なモデルは、ワークロード、ハードウェア、レイテンシ要件、予算によって異なります。
オープンソース LLM を自己ホストする場合は、独自のデータに基づいてモデルを微調整することで、LLM をドメインに適応させることもできます。これにより、法的分析、医療、財務、顧客サポートなどの特殊なタスクのパフォーマンスが大幅に向上し、特定のドメインでモデルが汎用の基盤モデルよりも優れたパフォーマンスを発揮できるようになります。
2026 年半ばまでに、いくつかのオープンファミリーが独自のフロンティアまたはその近くで、それぞれが異なる強みを持って競争します。
2026 年のリーダーには、最先端のコーディング、エージェント ツールの使用、および 100 万トークン以上にまで拡張されたコンテキスト ウィンドウという共通の要素が貫かれています。
パラメーター、ライセンス、コンテキスト ウィンドウ、および推奨される GPU を含む完全なリストについては、OpenLLMStack モデル ページを参照してください。
オープンソース モデルは、2025 年後半までに OpenRouter のトークン量の約 3 分の 1 に達しました。これは、そのプラットフォームでの採用の強力な証拠ですが、他のプロバイダーでのセルフホストの使用量とトラフィックは含まれていません。
2025 年と 202 年のオープンソース AI

6 は 3 つのトレンドによって定義されます。1 つは品質においてクローズド システムに匹敵するようになったモデル、崩壊した推論コスト、そして中国と独立系開発者に重心が移っていることです。オープンソース モデルは、もはや予算の選択肢ではありません。チームの割合が増えているため、それがデフォルトになります。
オープン モデルを使用して構築している場合、OpenLLMStack は現在のリリース、推論エンジン、最適化手法、エージェント フレームワークを 1 か所で追跡します。

## Original Extract

Explore 30+ open source LLM statistics for 2026, covering ecosystem growth, model usage, performance, inference costs, and leading developers.

30+ Open Source LLM Statistics & Trends (2026) — OpenLLMStack OpenLLMStack Models Engines Optimizations Agents Timeline Blog Models Engines Optimizations Agents Timeline Blog
GitHub
Back to blog
30+ Open Source LLM Statistics & Trends (2026)
Explore 30+ open source LLM statistics for 2026, covering ecosystem growth, model usage, performance, inference costs, and leading developers.
Open source LLMs went from a research curiosity to the backbone of real production systems in under three years. They now power coding assistants, agents, and enterprise pipelines, and in some fields they’ve overtaken the proprietary models that defined the early days of the boom.
So how big is the open source LLM ecosystem today? And how fast is it actually growing?
I pulled together the most useful open source LLM statistics I could find, all from primary sources like the Stanford AI Index, Hugging Face, Meta, OpenRouter, and the Stack Overflow Developer Survey. Each section links to the original sources so you can cite them directly.
Top open source LLM statistics
This article prioritizes first-party reports, research papers, official documentation, and original datasets.
An open source LLM generally refers to a language model that people can download, run, and modify. Compared with proprietary models that are only accessible through an API, open source LLMs give developers much greater control over deployment, customization, and infrastructure.
The term open source is often used loosely. Many models described as open source are actually released as open-weight models. Their weights are publicly available but the license may include restrictions that differ from a traditional open source software license. Because the industry commonly uses “open source LLM” to refer to both categories, this article follows that convention for simplicity.
How many open source LLMs are there?
There is no authoritative global count of open source LLMs. Hugging Face hosted more than 2 million public model repositories in 2025, but that total includes models for text, image, audio, robotics, and other tasks, as well as fine-tunes, adapters, quantizations, and derivatives.
The broader Hugging Face ecosystem is still useful for measuring the scale and direction of open model development:
The mean downloaded model size grew by about 25x between 2023 and 2025, while the median grew by only about 25% . That difference suggests a relatively small number of large models are pulling up the average while small models remain common.
Open source LLM adoption and usage
Open models have moved well past experimentation. By late 2025, they reached roughly one-third of all token volume on OpenRouter , a unified API platform that gives developers access to hundreds of AI models. The underlying study analyzed more than 100 trillion tokens over a year, covering November 2024 through November 2025.
DeepSeek processed about 2.6 times as many tokens as Qwen , the second-largest open source family in the study. By late 2025, however, no individual model consistently accounted for more than 20%-25% of open source model tokens. Usage had spread across five to seven competitive models.
Models developed in China rose from as little as 1.2% of weekly token volume in late 2024 to nearly 30% in some weeks .
Chinese open source models averaged 13.0% of weekly OpenRouter token volume over the full period.
Open source models developed outside China averaged 13.7% .
Proprietary models developed outside China retained an average share of about 70% .
Roleplay represented about 52% of open source model tokens , while programming was the second-largest category at roughly 15%-20% .
OpenRouter itself continued to grow after the study period. Menlo Ventures reported that the platform increased from 2.5 million to more than 8 million developers in roughly one year and reached an annualized run rate of about 1.5 quadrillion tokens in May 2026.
Developer adoption of AI tooling is now near-universal. 84% of respondents to the 2025 Stack Overflow Developer Survey use or plan to use AI tools, up from 76% the year before, and 51% of professional developers use them daily.
Sources: OpenRouter , Menlo Ventures , Stack Overflow
The most popular open source LLMs
Popularity depends on the metric. Downloads measure distribution, derivative repositories measure how often developers build on a model, and routed tokens measure hosted API usage.
Meta reported that the Llama family passed 1 billion cumulative downloads in March 2025. On Hugging Face, however, Qwen became the largest ecosystem for derivative work: the family had more than 113,000 direct derivative models by March 2026 and more than 200,000 repositories when every Qwen-tagged model was included.
These figures count downloads and repositories, not unique users or production deployments. Automated downloads, mirrors, quantizations, and repeated pulls can all affect the totals.
You can compare current model releases, licenses, context windows, and hardware requirements on the OpenLLMStack models page .
Open vs. closed model performance
The leading open model trailed the leading closed model by 3.3% on the Arena leaderboard in March 2026 , according to the Stanford AI Index. The gap had been only 0.5% in August 2024 before reopening during 2025.
The correct conclusion is not that open models have permanently reached parity. The gap is small enough to remain competitive, but it changes as new model generations arrive. Six of the top ten Arena models were closed as of March 2026.
Source: Stanford AI Index 2026
Where open source models come from
Models developed in China made up 41% of Hugging Face downloads in 2025 , the largest share attributed to a single country. China surpassed the United States in both monthly and overall downloads during the year.
The contributor mix changed at the same time. The share of model development attributed to industry fell from roughly 70% before 2022 to 37% in 2025 , while unaffiliated developers grew from 17% to 39% of downloads .
OpenRouter recorded a similar geographic shift, although through a different metric. Chinese open source models rose from 1.2% of weekly token usage in late 2024 to nearly 30% in some weeks of 2025.
Sources: Hugging Face , OpenRouter
The falling cost of LLM inference
An Andreessen Horowitz analysis found that the price of inference at a fixed level of MMLU performance fell by roughly 10x per year . At an MMLU score of 42, the cheapest observed price dropped from $60 per million tokens in 2021 to $0.06 in 2024, a 1,000-fold decline .
That historical estimate has limitations: it relies on MMLU, averages input and output pricing, and covers selected models from OpenAI, Anthropic, and Meta. It is useful as a directional cost trend, not a law that guarantees another 10x decline every year.
Current API pricing shows how large the spread can be. As of July 2, 2026:
Prices are per 1 million tokens at standard rates. On those published prices, DeepSeek-V4-Flash input is 97.2% cheaper and output is 99.1% cheaper than GPT-5.5.
Serving efficiency comes from more than cheaper hardware. Quantization, batching, caching, sparse architectures, and optimized kernels all contribute. OpenLLMStack tracks major techniques on the inference optimizations page and the engines that implement them in the inference directory .
Sources: Andreessen Horowitz , DeepSeek API pricing , OpenAI API pricing
Top open source LLM API providers
You do not need your own GPUs to run an open model. A growing set of serverless inference providers host the leading open weights behind a single API, billed per token, so you can switch models without managing infrastructure.
For the cheapest access to a single model, the first-party API from the model maker is often the lowest-priced option, such as the APIs from DeepSeek and Alibaba for their own models. Aggregators like OpenRouter route one request across many of these providers so you can compare price and speed.
DeepSeek R1: the breakout moment for open source LLMs
No single release did more for the profile of open source LLMs than DeepSeek R1. The model launched on January 20, 2025, and within days the DeepSeek app climbed to No. 1 on the U.S. Apple App Store, displacing ChatGPT and topping the charts in more than 50 countries.
The download surge was almost vertical. The app reached 2.6 million downloads across the App Store and Google Play by the Monday after launch, with more than 80% of all downloads coming in the previous seven days, and Appfigures data ranked the app No. 1 worldwide.
The market reaction was just as dramatic. On January 27, 2025, Nvidia lost about $589 billion in market value, the largest single-day loss for any company in history, after DeepSeek showed that a frontier-grade open model could reportedly be trained for around $5.6 million.
You can trace these milestones on the OpenLLMStack timeline .
Sources: TechCrunch , Bloomberg , Forbes
No authoritative organization maintains a global count of open source LLMs. Hugging Face hosted more than 2 million public model repositories in 2025, but that figure covers all model types and includes derivatives, adapters, and quantizations. It is not a count of unique LLMs.
Meta reported more than 1 billion cumulative Llama downloads in March 2025.
Qwen had the largest derivative ecosystem reported by Hugging Face in March 2026, with more than 113,000 direct derivative models.
DeepSeek led open source usage on OpenRouter with 14.37 trillion tokens processed from November 2024 through November 2025.
There is no single best open source LLM in 2026. Some models are good at coding, others at reasoning, long-context processing, or multilingual tasks. The right model depends on your workload, hardware, latency requirements, and budget.
If you self-host an open source LLM, you can also adapt it to your domain by fine-tuning the model on proprietary data. This can significantly improve performance for specialized tasks, such as legal analysis, healthcare, finance, or customer support, helping the model outperform a general-purpose foundation model in your specific domain.
By mid-2026, several open families compete at or near the proprietary frontier, each with a different strength.
A common thread runs through the 2026 leaders: state-of-the-art coding, agentic tool use, and context windows that now stretch to 1 million tokens or more.
For the full list with parameters, licenses, context windows, and recommended GPUs, see the OpenLLMStack models page .
Open source models reached roughly one-third of token volume on OpenRouter by late 2025. This is strong evidence of adoption on that platform, but self-hosted usage and traffic on other providers are not included.
Open source AI in 2025 and 2026 is defined by three trends: models that now rival closed systems on quality, inference costs that have collapsed, and a center of gravity shifting toward Chinese and independent developers. Open source models are no longer the budget option. For a growing share of teams, they’re the default.
If you’re building with open models, OpenLLMStack tracks current releases, inference engines, optimization techniques, and agent frameworks in one place.

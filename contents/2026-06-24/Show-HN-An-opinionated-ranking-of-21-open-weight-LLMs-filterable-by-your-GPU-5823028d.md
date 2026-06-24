---
source: "https://northwoodsystems.ai/research/open-source-models-big-board"
hn_url: "https://news.ycombinator.com/item?id=48659077"
title: "Show HN: An opinionated ranking of 21 open-weight LLMs, filterable by your GPU"
article_title: "The State of Open-Source Models 2026: The Big Board — Northwood Systems"
author: "AndrewLiu96"
captured_at: "2026-06-24T13:43:54Z"
capture_tool: "hn-digest"
hn_id: 48659077
score: 2
comments: 0
posted_at: "2026-06-24T13:01:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An opinionated ranking of 21 open-weight LLMs, filterable by your GPU

- HN: [48659077](https://news.ycombinator.com/item?id=48659077)
- Source: [northwoodsystems.ai](https://northwoodsystems.ai/research/open-source-models-big-board)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T13:01:31Z

## Translation

タイトル: Show HN: 21 個のオープンウェイト LLM の独自のランキング、GPU でフィルタリング可能
記事のタイトル: 2026 年のオープンソース モデルの現状: ビッグボード — Northwood Systems
説明: オープンウェイト AI モデルの決定版ランク付けボードです。全 21 個が 6 つの軸で評価され、ハードウェア層、カテゴリ、ライセンスによってフィルタリングされています。 2026 年 6 月に更新されました。

記事本文:
2026 年のオープンソース モデルの現状: ビッグボード — Northwood Systems Northwood Systems Industries Research Security About Contact デモのリクエスト The Big Board · 2026 Open-weight Class 最終更新日 2026 年 6 月 24 日 · Living Board 言語 EN 中文 FR 2026 年のオープンソース モデルの現状
無差別級モデルを見つけて評価するための決定版ボード - すべて 21 種類があり、ランク付けされ、グレード付けされ、タグ付けされています。新しいモデルは毎月出荷されます。上のスタンプは鮮度チェックです。
Composite 91 オープンソース AI の状況は急速に変化しており、最高のオープンウェイト モデルは現在、クローズド フロンティアとの打撃を交換しています。それらの数は 1 つのチームが追跡できる数を超えており、ボード上で最も興味深いのは、誰がそれらを構築しているかということです。これはクラス全体の、ランク付けされ、採点された、意見の入った読み物です。スペックとベンチマークは出典によるものです。ランキング、成績、評決は一つの観点であり、測定ではありません。ボードをざっと見るか、任意のモデルにドリルで穴を開けます。タグを使用して、ハードウェア、ライセンス、ユースケースに適合するものを見つけてください。
2026 年 6 月 24 日に確認されました。これらは毎月発送されます。これを生きているボードとして扱い、数字に賭ける前にソースノートを確認してください。
1 から 21 までの 1 つのランク付けされたリストは、重要性順に並べられています。つまり、生の機能の組み合わせ、そのものがどれほど真にオープンであるか、どの程度展開可能であるか、そしてラボがどの程度現場を動かしているのかを表します。機能はベンダーのデッキではなく、人工分析と LMArena に基づいています。ランキング、レーダーのグレード、評決は私のものです。ベンチマークはソースから提供されます。
これらのエントリーの多くはモデルではなく、家族です。 Qwen だけでも、電話機で実行できる 0.6B から、8 GPU ノードを必要とする 235B-A22B までの範囲に及びます。実際的な効果: 「Qwen を実行できますか」と尋ねるのではなく、どの Qwen が自分の GPU に適合するかを尋ねてください。
無差別級のフロンティアは中国だが、近いわけではない。ザ

i、Moonshot、DeepSeek、MiniMax、そしてその背後にあるディープベンチは、現在、西側のオープンラボでは対応していない作業を行っており、そのほとんどは実際に使用できる MIT ライセンスに基づいています。
セルフホスティングの場合、メモリフロアを設定するのはアクティブな数ではなく、合計パラメータ数です。FP8 ではパラメータごとにおよそ 1 バイト、4 ビットでは 0.5 バイトとなります。 8 GPU H100 ノードでは約 640 GB が得られます。 H200 ノード、約 1.1TB。専門家を混合すると、アクティブなパラメータが小さく抑えられるため、速度とコストが高まりますが、それでもモデル全体をメモリに保持する必要があります。
推論、コーディング、数学、知識、エージェント、ロングコンテキストの 6 つの能力グレードは、それぞれ 0 から 100 のスコアを付けられ、平均化されて複合化されます。私の判断はベンチマークと人工分析知能指数に基づいており、実験室の結果ではありません。
以下のヘッドライン ベンチマークのほぼすべてはベンダーによって報告されたものであり、ラボが独自のモデルに基づいて最大限の思考努力のもとで実行しています。 Scale の SEAL リーダーボードでは、ベンダーの足場が結果を 10 ～ 30 ポイント水増ししていることがわかりました。本当に比較可能なスパインの 1 つは、Artificial Analysis Intelligence Index です。
カテゴリ、ハードウェア、ライセンスでフィルタリングします。 8 つの段階的なフロンティア モデルをレーダー軸ごとに並べ替えます。カードを開いて完全な内訳を確認してください。
すべてのカテゴリ すべてのハードウェア すべてのライセンス すべての 21 モデル、重要度順にランク付け Tier 1 フロンティア 閉ざされたフロンティアで貿易が打撃を受ける 01 01 推論 コーディング 数学知識 Agentic Long ctx Composite 91 MIT · ~753B / 40B active · 1M ctx GLM-5.2
現在ダウンロードできる唯一の最高の無差別級モデル。 GLM-5 は 4 月に人工分析インテリジェンス インデックスでオープンウェイト スロットのトップとなり、オープン モデルとして初めてその層に到達し、GLM-5.2 は 6 月にコーディング数を押し上げました。
エージェントコーディングと長期的なツールの使用に優れています。 GLM-5 は SWE ベンチ Veri で 77.8% を記録

クロード作品4.6の約3ポイント以内にあります。
Z.ai は、清華社からスピンアウトした Zhipu AI の国際ブランドです。重要な詳細: GLM-5 は、Huawei Ascend シリコン上でエンドツーエンドでトレーニングされたと伝えられています。これはモデルのリリースではなく、主権の表明であり、単一の Nvidia GPU を使用せずに構築されたフロンティア モデルです。
~753B MoE (40B アクティブ): FP8 または 4 ビットで 8xH100/H200 ノードを計画します。完全な BF16 は 2 つのノードにまたがります。
推論 コーディング 数学知識 Agentic Long ctx Composite 90 Modified MIT · 1T / 32B active · 256K ctx キミ K2.6
GLMが最高のモデルなら、キミは辞めないモデルだ。 Artificial Analysis は、K2.6 を新しい主要な無差別重量モデルと呼び、見出しは持久力です。
長期的なエージェント業務に優れています。報告によると、エージェント群アーキテクチャは 12 時間、4,000 の調整されたステップを転倒することなく実行されました。これは、SWE-bench Pro で GPT-5.4 を上回った最初のオープン モデルであり、GPT-5.5 と 58.6% で並んでいます。
推論 コーディング 数学知識 Agentic Long ctx Composite 90 MIT · Flash 284B / 13B · 1M ctx DeepSeek V4
GLM はボード上で最高のモデルです。 DeepSeek はその中で最も重要なラボです。
最先端の推論と競技数学に優れています。 V3.2-Speciale は、厳密な推論において Gemini 3.0 Pro と同等に達し、2025 年の IMO および IOI で金メダルの結果を獲得しました。
推論 コーディング 数学知識 Agentic Long ctx Composite 85 MIT · 230B / 10B active · 204K ctx MiniMax-M2
他のフロンティア モデルでは、インテリジェンスとハードウェア予算のどちらかを選択する必要があります。 MiniMax はほとんどそうではありません。
他のサービスでは実現できないサービスコストで、エージェントによるツールの呼び出しとコーディングに優れています。 VentureBeat は、M2 をエージェント作業用のオープンソース LLM の新たな王様と呼び、リリース時にはオープンウェイト インテリジェンス インデックスでトップになりました。
推論 コーディング 数学知識 Agentic Long ctx Composite 86 Apache 2.0 ·

0.6B → 235B-A22B · 最大 262K Qwen3 ファミリー
Qwen は 1 つのモデルではなく、フリートです。それがその利点であり、アスタリスクです。
幅広さで優れています。 200 以上の言語、電話機で実行される 0.6B からクラスターを必要とする 235B-A22B までのあらゆるサイズ、そして最も深いツール エコシステムがここにあります。 1 つのファミリーで組織を標準化したい場合は、これが最適です。
推論 コーディング 数学知識 Agentic Long ctx Composite 80 Apache 2.0 · 675B / 41B active · 256K ctx Mistral Large 3
ミストラルがここにあるのにはベンチマークではない理由がありますが、それは正直な理由です。それは信頼できる西側の選択肢であり、多くの購入者にとってそれがすべての決定です。
中国人以外の答えとして優れています。信念による Apache 2.0、真のマルチモーダルな範囲、そして杭州からの重みを運用したくないヨーロッパの役員室と米国の規制部門に着地したデータ主権の物語。
推論 コーディング 数学知識 Agentic Long ctx Composite 84 Apache 2.0* · E2B 2B → 31B dense · 256K ctx Gemma 4
単一 GPU、フルストップ、および Gemma 4 (2026 年 3 月) で実行できる最高の機能により、そのリードはさらに広がりました。
サイズを超えたパンチ力に優れる。 31B 高密度モデルは、Arena AI (ELO 1,452) のすべてのオープン モデルの中で 3 位にランクされており、世代間の飛躍は本当の話です。AIME 2026 は Gemma 3 の 20.8% から 89.2% に、Codeforces ELO は 110 から 2,150 に上昇しました。これは、オープン モデルの記録における単一世代の飛躍としては最大です。
推論 コーディング 数学の知識 Agentic Long ctx Composite 77 Llama コミュニティ ライセンス · Maverick 17B アクティブ · 非常に長い Llama 4
このカテゴリ全体を開拓したラボは、今や明らかに遅れをとっており、Gemma 4 がちょうどそれを追い越したところです。それは、以前の姿ではなく、今もボード上に残ります。
このクラスのサイズでマルチモーダルに優れ、さらにエコシステムの重力、どこよりも深い微調整とツール ベースに優れています (Nvidia

の Nemotron ファミリーはそこから蒸留されます)。
Cohere の最初の Apache 2.0 リリースでは、以前の 5 つのコマンド モデルが 1 つに統合されています。これは、エージェントのワークフロー、マルチモーダル (テキスト、画像、ツールの使用)、48 の言語、128K のコンテキストなど、真のエンタープライズ ジェネラリストです。ネイティブの引用グラウンディングは現実的で便利です。ツールから取得すると、明示的なグラウンディング スパンが生成されます。しかし、それは 1 つの機能であり、すべてではありません。 25B がアクティブな 218B MoE なので、2 台の H100 または 1 台の Blackwell で実行されます。
NVIDIA は、Llama を推論/エージェント ファミリ (Super 49B、Ultra 253B) に調整し、1M コンテキストのハイブリッド Mamba-Transformer MoE ラインと 30B マルチモーダル Nano Omni を備えています。すでに購入しているシリコンのスループットに合わせて調整されています。
Tencent のオープン 295B MoE (21B アクティブ、256K コンテキスト) は元 OpenAI 研究者によって率いられ、WeChat 規模のディストリビューションが背後にあります。有能なジェネラリスト。数値は依然としてほとんどがベンダーによって報告されています。
上海の研究所の小型アクティブ MoE ビジョン言語モデル (198B / 11B アクティブ) は、独自のベンチマークで、100 万トークンあたり ~0.10 ドルではるかに大規模なモデルを上回っています。それが成立すれば素晴らしいです。引用する前に確認してください。
ここで唯一の真にオープンなモデル: Ai2 は、完全なフロー (データセット、中間チェックポイント、RL ステージ) に加えて、検査可能な推論トレースを含む「思考」系統を公開します。アメリカのオープンエブリシングオプション、そして最高のフルオープン32B。
コンプライアンスレーン: 既存のハードウェア向けの適度な 3B/8B/30B サイズ、ISO 42001、米国ベンダー、IBM の補償。ボード上で最もスマートなモデルではありませんが、調達が最も簡単です。
ハイブリッド SSM-Transformer (Mamba) アーキテクチャは、ピークのスマート性と引き換えに長いコンテキストの速度を実現します。長い入力では最大約 2.5 倍高速になり、そのサイズ クラスで最長のコンテキストになります。難しいのは、推論ではなく文書を選択することです。
アブダビのTIIが効率的なハイブリッドを出荷

リッド: 最大 7 倍のサイズのモデルを上回ると伝えられる 7B、さらにオープン ビジョン (知覚、OCR)、および主要なアラビア モデル。湾岸諸国の主権AI参入者。
読むだけでなく見ることが仕事の場合、これは開かれたフロンティアです。SOTA オープン マルチモーダル ベンチマーク (78B で 72.2 MMMU) を達成しながら、強力なテキスト スキルを維持します。デフォルトのオープン VLM。
3 つのこと、もし 3 つだけを取り上げるなら
上位 4 社だけでなく、Hunyuan、StepFun、ERNIE を含むベンチ全体も、そのほとんどが MIT および Apache ライセンスの下で商用利用可能です。欧米のオープンエコシステムは現在、リーダーボードのトップではなく、価値観と信頼、ミストラルは主権、コヒアとIBMはコンプライアンス、Ai2は徹底した透明性で競争している。
モデルをダウンロードしてデプロイメントを所有するオープンウェイトと、ウェイトも公開するラボが API の背後に最適なモデルを保持するオープンラボがあります。クウェンとアーニーが威嚇射撃。他の企業が追随するかどうかに注目してください。フロンティアが API 専用になる日は、「オープン」という言葉があなたが思っているような意味を持たなくなる日だからです。
MiniMax-M2、DeepSeek V4-Flash、または Cohere Command A+ を実行する単一の 8 GPU ノードにより、お客様が管理するハードウェア上で、建物の外に出ることのないデータを使用して、真の意味でのデータを取得できます。何を保護しているかによっては、それが重要な唯一のランキングになる場合があります。
利用可能な場合は、人工分析と LMArena に基づいた機能。ベンダーが報告した数値にはラベルが付けられています。公式モデルカードからのパラメータ数とライセンス。ハードウェア層は経験則による推定です。ここにあるすべてのモデルは 2026 年 1 月以降のリリースです。現在の主力製品とそのオープン ステータスは公開日に確認してください。
単一の 8 GPU ノードにより、制御するハードウェア上でフロンティア クラスの作業が可能になります。
Northwood は、機密性の高いデータ内にオープン ウェイトを展開するためのコンピューティング、モデル、ソフトウェア、およびサービスを提供します。

専門組織 - 許可、引用、監査ログが組み込まれています。知識が環境から離れることはありません。

## Original Extract

The definitive ranked board for open-weight AI models — all 21, graded on six axes, filtered by hardware tier, category, and license. Updated June 2026.

The State of Open-Source Models 2026: The Big Board — Northwood Systems Northwood Systems Industries Research Security About Contact Request a demo The Big Board · 2026 Open-Weight Class Last updated 24 Jun 2026 · Living board Language EN 中文 FR The State of Open-Source Models 2026
The definitive board for finding and evaluating open-weight models - all twenty-one, ranked, graded, and tagged. New models ship monthly; the stamp above is your freshness check.
Composite 91 The open source AI landscape is rapidly changing, the best open-weight models now trade blows with the closed frontier. There are more of them than any one team can track, and the most interesting story on the board is who's building them. This is the whole class, ranked and graded, an opinionated read. The specs and benchmarks are sourced; the ranking, the grades, and the verdicts are a point of view, not a measurement. Skim the board, or drill into any model. Use the tags to find one that fits your hardware, your license, and your use case.
Verified 2026-06-24. These ship monthly - treat this as a living board, and see the source notes before betting on a number.
One ranked list, 1 to 21, ordered by significance: a blend of raw capability, how genuinely open the thing is, how deployable it is, and how much the lab moves the field. Capability is anchored on Artificial Analysis and LMArena rather than vendor decks. The ranking, the radar grades, and the verdicts are mine; the benchmarks are sourced.
A lot of these entries aren't a model, they're a family. Qwen alone spans a 0.6B you can run on a phone to a 235B-A22B that needs an 8-GPU node. The practical effect: don't ask "can I run Qwen," ask which Qwen fits your GPU.
The frontier of open weights is Chinese, and it isn't close. Z.ai, Moonshot, DeepSeek, MiniMax, and a deep bench behind them are doing work the Western open labs aren't matching right now, mostly under MIT licenses you can actually use.
For self-hosting, it's the total parameter count that sets your memory floor, not the active count: figure roughly one byte per parameter at FP8, half a byte at 4-bit. An 8-GPU H100 node gives you about 640GB; an H200 node, about 1.1TB. Mixture-of-experts keeps the active params small, which buys speed and cost, but you still have to hold the whole model in memory.
Six capability grades — reasoning, coding, math, knowledge, agentic, long-context — each scored 0 to 100 and averaged into a composite. My judgment, informed by benchmarks and the Artificial Analysis Intelligence Index, not a lab result.
Almost every headline benchmark below is vendor-reported — run by the lab on its own model under maximum thinking effort. Scale's SEAL leaderboard finds vendor scaffolding inflates results by ten to thirty points. The one genuinely comparable spine is the Artificial Analysis Intelligence Index.
Filter by category, hardware, and license; sort the eight graded frontier models by any radar axis; open a card for the full breakdown.
All categories All hardware All licenses All 21 models, ranked by significance Tier 1 The Frontier Trading blows with the closed frontier 01 01 Reasoning Coding Math Knowledge Agentic Long ctx Composite 91 MIT · ~753B / 40B active · 1M ctx GLM-5.2
The single best open-weight model you can download today. GLM-5 took the top open-weight slot on the Artificial Analysis Intelligence Index in April, the first open model to reach that tier, and GLM-5.2 pushed the coding numbers higher in June.
Excels at agentic coding and long-horizon tool use. GLM-5 posted 77.8% on SWE-bench Verified, within about three points of Claude Opus 4.6.
Z.ai is the international brand of Zhipu AI, spun out of Tsinghua. The detail that matters: GLM-5 was reportedly trained end to end on Huawei Ascend silicon. That's not a model release, it's a sovereignty statement, a frontier model built without a single Nvidia GPU.
A ~753B MoE (40B active): plan for an 8xH100/H200 node at FP8 or 4-bit; full BF16 spans two nodes.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 90 Modified MIT · 1T / 32B active · 256K ctx Kimi K2.6
If GLM is the best model, Kimi is the one that doesn't quit. Artificial Analysis called K2.6 the new leading open-weight model, and the headline is endurance.
Excels at long-horizon agentic work. An agent-swarm architecture reportedly ran 12 hours and 4,000 coordinated steps without falling over. It's the first open model to beat GPT-5.4 on SWE-bench Pro, and it ties GPT-5.5 there at 58.6%.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 90 MIT · Flash 284B / 13B · 1M ctx DeepSeek V4
GLM is the best model on the board. DeepSeek is the most important lab on it.
Excels at frontier reasoning and competition math. V3.2-Speciale reached parity with Gemini 3.0 Pro on hard reasoning and took gold-medal results on the 2025 IMO and IOI.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 85 MIT · 230B / 10B active · 204K ctx MiniMax-M2
The other frontier models make you choose between intelligence and a hardware budget. MiniMax mostly doesn't.
Excels at agentic tool-calling and coding at a serving cost the others can't touch. VentureBeat called M2 the new king of open-source LLMs for agentic work, and it topped the open-weight Intelligence Index at release.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 86 Apache 2.0 · 0.6B → 235B-A22B · up to 262K Qwen3 family
Qwen isn't one model, it's a fleet, and that's its advantage and its asterisk.
Excels at breadth. Over 200 languages, every size from a 0.6B that runs on a phone to a 235B-A22B that needs a cluster, and the deepest tooling ecosystem here. If you want one family to standardize an org on, this is it.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 80 Apache 2.0 · 675B / 41B active · 256K ctx Mistral Large 3
Mistral is here for a reason that isn't a benchmark, and it's an honest one: it's the credible Western option, and for a lot of buyers that's the whole decision.
Excels at being the non-Chinese answer. Apache 2.0 by conviction, real multimodal range, and a data-sovereignty story that lands in European boardrooms and regulated US sectors that would rather not run weights from Hangzhou.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 84 Apache 2.0* · E2B 2B → 31B dense · 256K ctx Gemma 4
The best capability you can run on a single GPU, full stop, and Gemma 4 (March 2026) widened that lead.
Excels at punching wildly above its size. The 31B dense model ranks #3 among all open models on Arena AI (ELO 1,452), and the generational jump is the real story: AIME 2026 went from 20.8% on Gemma 3 to 89.2%, and Codeforces ELO from 110 to 2,150, the largest single-generation leap on record for an open model.
Reasoning Coding Math Knowledge Agentic Long ctx Composite 77 Llama Community License · Maverick 17B active · very long Llama 4
The lab that opened this whole category, now visibly behind, and Gemma 4 just passed it. It stays on the board for what it still is, not just what it was.
Excels at multimodal in its size class and, still, ecosystem gravity, the deepest fine-tune and tooling base anywhere (Nvidia's Nemotron family is distilled from it).
Cohere's first Apache 2.0 release consolidates five previous Command models into one. It's a genuine enterprise generalist: agentic workflows, multimodal (text, image, tool use), 48 languages, 128K context. Native citation grounding is real and useful — when it pulls from a tool it emits explicit grounding spans — but it's one feature, not the whole story. 218B MoE with 25B active, so two H100s or a single Blackwell runs it.
NVIDIA distills and tunes Llama into a reasoning/agentic family (Super 49B, Ultra 253B), now with a hybrid Mamba-Transformer MoE line at 1M context and a 30B multimodal Nano Omni. Tuned for throughput on the silicon you're already buying.
Tencent's open 295B MoE (21B active, 256K context), led by a former OpenAI researcher, with WeChat-scale distribution behind it. Capable generalist; numbers still mostly vendor-reported.
A Shanghai lab's tiny-active MoE vision-language model (198B / 11B active) that, on its own benchmarks, outruns much larger models at ~$0.10 per million tokens. Impressive if it holds; verify before quoting.
The only truly open model here: Ai2 publishes the full flow — datasets, intermediate checkpoints, RL stages — plus a "Think" lineage with inspectable reasoning traces. America's open-everything option, and the best fully-open 32B.
The compliance lane: modest 3B/8B/30B sizes for existing hardware, ISO 42001, US vendor, IBM indemnity. Not the smartest model on the board — the easiest to get through procurement.
A hybrid SSM-Transformer (Mamba) architecture that trades peak smarts for long-context speed — up to ~2.5× faster on long inputs, the longest context in its size class. The pick when the document, not the reasoning, is the hard part.
Abu Dhabi's TII ships efficient hybrids: a 7B that reportedly out-reasons models up to 7× its size, plus open vision (Perception, OCR) and the leading Arabic models. The Gulf's sovereign-AI entrant.
If the job is seeing, not just reading, this is the open frontier: SOTA open multimodal benchmarks (72.2 MMMU at 78B) while keeping strong text skills. The default open VLM.
Three things, if you only take three
Not just the top four, but a whole bench, Hunyuan, StepFun, ERNIE, behind them, mostly under MIT and Apache licenses you can use commercially. The Western open ecosystem now competes on values and trust, Mistral on sovereignty, Cohere and IBM on compliance, Ai2 on radical transparency, rather than on topping the leaderboard.
There's open-weight, where you download the model and own your deployment, and open-lab, where a lab that also publishes weights keeps its best model behind an API. Qwen and ERNIE are the warning shots. Watch whether the others follow, because the day the frontier goes API-only is the day "open" stops meaning what you think it means.
A single 8-GPU node running MiniMax-M2, DeepSeek V4-Flash, or Cohere Command A+ gets you genuinely close, on hardware you control, with data that never leaves your building. Depending on what you're protecting, that might be the only ranking that matters.
Capability anchored on Artificial Analysis and LMArena where available; vendor-reported figures are labelled. Parameter counts and licenses from official model cards; hardware tiers are rule-of-thumb estimates. Every model here is a post-Jan-2026 release - confirm the current flagship and its open status on publish day.
A single 8-GPU node now gets you frontier-class work on hardware you control.
Northwood provides the compute, models, software, and services to deploy open weights inside sensitive technical organizations - with permissioning, citations, and audit logs built in. Your knowledge never leaves your environment.

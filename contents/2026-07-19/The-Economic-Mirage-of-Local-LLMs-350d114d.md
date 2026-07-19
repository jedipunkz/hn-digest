---
source: "https://eamag.me/2026/the-economic-mirage-of-local-llms"
hn_url: "https://news.ycombinator.com/item?id=48966745"
title: "The Economic Mirage of Local LLMs"
article_title: "The Economic Mirage of Local LLMs"
author: "eamag"
captured_at: "2026-07-19T11:04:41Z"
capture_tool: "hn-digest"
hn_id: 48966745
score: 1
comments: 0
posted_at: "2026-07-19T10:40:51Z"
tags:
  - hacker-news
  - translated
---

# The Economic Mirage of Local LLMs

- HN: [48966745](https://news.ycombinator.com/item?id=48966745)
- Source: [eamag.me](https://eamag.me/2026/the-economic-mirage-of-local-llms)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T10:40:51Z

## Translation

タイトル: 地元LLMの経済的蜃気楼
説明: HackerNews で同じ議論を何度も目にします。「お金を節約するには、LLM をセルフホストするべきだ」というものですが、数字を計算してみると、それは真実ではありません。

記事本文:
ライト モード リーダー モード {\n if (!a2.isFolder && !b2.isFolder || a2.isFolder && b2.isFolder) {\n return (a2.displayName || \"\").localeCompare(b2.displayName || \"\", void 0, {\n 数値: true,\n 感度: \"base\"\n });\n }\n if (!a2.isFolder && b2.isFolder) {\n return 1;\n }\n return -1;\n }","filterFn":"(node) => node.slugSegment !== \"tags\"","mapFn":"(node) => {\n return node;\n }"}"> エクスプローラー
地元LLMの経済的蜃気楼
HackerNews で同じ議論を繰り返し目にします。「お金を節約するために LLM を自己ホストせよ」というものですが、数字を計算してみるとそれは真実ではありません。データのセキュリティにお金を払うつもりがない限り、光熱費やハードウェアの減価償却費はあまりにも高額です。
標準の量子化された 8B オープンウェイト モデルをハイエンドのコンシューマ ワークステーションでローカルに実行するとします。アクティブな推論負荷の下では、RTX 4090 ～ 5090 システムは通常約 500 ～ 700 W を消費します。
vLLM を使用すると、生成が GPU のメモリ帯域幅によって厳密に制限されるため、1 秒あたり 100 トークンに簡単に到達します。
アクティブコスト: 100 tps で 18,000 トークンの生成には 3 分かかります。 RTX 4090 セットアップの 500W では、その 3 分間で 0.025 kWh の電力が消費され、純粋な電力料金で約 0.0095 米ドルかかります。
集中コスト: DeepSeek V4 Flash などの同等の軽量モデルの呼び出しには、およそ 0.21 万トークン ∗ ∗ のコストがかかります。これは、18,000 トークン トークン ∗ ∗ 0.0038 米ドルを意味します。
ローカル GPU を絶対最大効率で実行した場合でも、アクティブなローカル電力コストは集中型 API よりも約 2.5 倍から 3.5 倍高くなります。
実際には、ローカル システムは、ユーザーのクエリ間のアイドル状態にほとんどの時間を費やします。
ハイエンドの Nvidia デュアル GPU デスクトップ システムがアイドル状態で 90 W である場合、24 時間年中無休で電力を供給し続けると、年間約 788 kWh を消費します。その結果、3 イーが発生します

ar アイドル税 $898 USD
より効率的な Apple Mac Studio M4 Max (アイドル時は約 10W) に切り替えたとしても、ハードウェアの初期費用が依然として障壁となります。 3 年間の総所有コストを比較してみましょう (ベースライン API コスト 230 ドルと比較して、1 日あたり 100 万トークンを想定)。
Nvidia PC (デュアル RTX 4090) : 4、500 キャパシティー + 898 アイドル + 578 アクティブ = ∗ ∗ 5,976 総コスト** (API コストの 26.0 倍)
Mac Studio M4 Max (128GB) : 3、600 キャパシティー + 100 アイドル + 231 アクティブ = ∗ ∗ 合計コスト 3,931** (API コストの 17.1 倍)
なぜローカルハードウェアは競争できないのでしょうか?
推論の順次デコード フェーズ中に、システムは生成された単一トークンごとにすべてのモデルの重みをメモリからプロセッサの計算コアにロードする必要があります。バッチ サイズ 1 では、演算強度がバイトあたり約 1 FLOP に低下し、メモリ上で待機している間、コンピューティング コアがアイドル状態になります。
これにより、消費者向けハードウェアは 2 つの悪いトレードオフを余儀なくされます。
VRAM (Nvidia GPU): RTX 5090 などのグラフィック カードは優れた帯域幅 (~1,792 GB/秒) を備えていますが、その容量は 32 GB です。 70B パラメータ モデルを 4 ビット精度で実行するには、複数のカードに分割するか、レイヤを CPU にオフロードする必要があります。
帯域幅 (Apple UMA および AMD Strix Halo): ユニファイド メモリ アーキテクチャにより、より大きなモデル (最大 128GB または 512GB) をメモリに搭載できますが、帯域幅は制限されています (M4 Max の場合は 546 GB/秒)。高密度 70B モデルは、1 秒あたりわずか 8 ～ 15 トークンで実行されます。
クラウド プロバイダーは、継続的なバッチ処理を使用することで大規模なマルチテナンシーの恩恵を受け、同時ユーザー間でメモリ エネルギー税を償却します。
Ternary-Bonsai-27B のようなネイティブ 3 値シリコン モデルを 3 値 ( { − 1 , 0 , + 1 } ) に量子化すると、乗算を加算に置き換えることで演算エネルギーが桁違いに削減されます。

チップメーカーがネイティブの 3 値テンソル カーネルをハードウェアに組み込む必要があります。
スマートスリープ ランタイム管理: クエリが到着した瞬間に推論ブロックを 0W ディープスリープ状態からアクティブな実行に移行できるランタイム エンジンが必要です。
地元LLMの経済的蜃気楼
Mac 上で並列拡散 LLM が遅い理由
Quartz v5.0.0 で作成 © 2026

## Original Extract

I keep seeing the same argument on HackerNews: “Self-host your LLMs to save money”, but it’s just not true if you run the numbers.

Light mode Reader mode {\n if (!a2.isFolder && !b2.isFolder || a2.isFolder && b2.isFolder) {\n return (a2.displayName || \"\").localeCompare(b2.displayName || \"\", void 0, {\n numeric: true,\n sensitivity: \"base\"\n });\n }\n if (!a2.isFolder && b2.isFolder) {\n return 1;\n }\n return -1;\n }","filterFn":"(node) => node.slugSegment !== \"tags\"","mapFn":"(node) => {\n return node;\n }"}"> Explorer
The Economic Mirage of Local LLMs
I keep seeing the same argument on HackerNews: “Self-host your LLMs to save money”, but it’s just not true if you run the numbers. Utility bills and hardware depreciation just cost too much, unless you’re willing to pay for the data security.
Suppose you run a standard, quantized 8B open-weight model locally on a high-end consumer workstation. Under active inference load, an RTX 4090-5090 system typically draws around 500-700W
With vLLM, you will easily hit 100 tokens per second because generation is bound strictly by the GPU’s memory bandwidth.
The Active Cost: Generating 18k tokens at 100 tps takes 3 minutes. At 500W for an RTX 4090 setup, those 3 minutes consume 0.025 kWh of electricity, costing roughly $0.0095 USD in pure power
The Centralized Cost: Calling an equivalent lightweight model like DeepSeek V4 Flash costs roughly 0.21 p er mi l l i o n t o k e n s ∗ ∗ , m e anin g 18 k t o k e n so na co mm er c ia l A P I cos t s ∗ ∗ 0.0038 USD .
Even running your local GPU at absolute maximum efficiency, active local electricity costs are ~2.5x to 3.5x higher than centralized APIs
In practice, local systems spend most of their time sitting idle between user queries.
If a high-end Nvidia dual-GPU desktop system idles at 90W , keeping it powered 24/7 consumes roughly 788 kWh annually. That results in a three-year idle tax of $898 USD
Even if you switch to a more efficient Apple Mac Studio M4 Max (which idles at around 10W), the upfront hardware cost still creates a barrier. Let’s compare the three-year total cost of ownership (assuming 1M tokens/day, compared against a baseline API cost of $230 ):
Nvidia PC (Dual RTX 4090) : 4 , 500 C a pE x + 898 idle + 578 a c t i v e = ∗ ∗ 5,976 total cost** (a 26.0x multiple over the API cost)
Mac Studio M4 Max (128GB) : 3 , 600 C a pE x + 100 idle + 231 a c t i v e = ∗ ∗ 3,931 total cost** (a 17.1x multiple over the API cost)
Why can’t local hardware compete?
During the sequential decode phase of inference, the system must load all model weights from memory into the processor’s compute cores for every single token generated. At batch size 1, the arithmetic intensity drops to roughly 1 FLOP per byte, leaving compute cores idle while waiting on memory.
This forces consumer hardware into two bad trade-offs:
The VRAM (Nvidia GPUs): Graphics cards like the RTX 5090 have good bandwidth (~1,792 GB/s), but their capacity is 32 GB. Running a 70B parameter model at 4-bit precision requires splitting it across multiple cards or offloading layers to the CPU
The Bandwidth (Apple UMA & AMD Strix Halo): Unified memory architectures let you fit bigger models (up to 128GB or 512GB) in memory, but their bandwidth is restricted (546 GB/s for M4 Max). A dense 70B model runs at only 8 to 15 tokens per second .
Cloud providers benefit from massive multi-tenancy by using continuous batching, and amortize the memory energy tax across concurrent users
Native Ternary Silicon Quantizing models to ternary values ( { − 1 , 0 , + 1 } ) like Ternary-Bonsai-27B reduces arithmetic energy by order of magnitude by replacing multiplications with additions. We need chipmakers to build native ternary tensor kernels into hardware.
Smart-Sleep Runtime Management: We need runtime engines that can transition an inference block from a 0W deep-sleep state to active execution the moment a query comes in
The Economic Mirage of Local LLMs
Why Parallel Diffusion LLMs Are Slow on a Mac
Created with Quartz v5.0.0 © 2026

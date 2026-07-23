---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49022097"
title: "Show HN: Avoiding the Memory Wall by computing LLM inference directly inside RAM"
article_title: ""
author: "pcdeni"
captured_at: "2026-07-23T15:03:34Z"
capture_tool: "hn-digest"
hn_id: 49022097
score: 1
comments: 0
posted_at: "2026-07-23T14:21:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Avoiding the Memory Wall by computing LLM inference directly inside RAM

- HN: [49022097](https://news.ycombinator.com/item?id=49022097)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T14:21:32Z

## Translation

タイトル: HN の表示: RAM 内で LLM 推論を直接計算することでメモリの壁を回避する
HN テキスト: PrismML の 1 ビット/3 値 Bonsai モデルをめぐる興奮により、業界はスマートフォン大手、特に Apple がエッジ デバイスに LLM をどのように実装するかを注視しています。
AI をデバイス上で動かすことは、素晴らしい、そして必要な戦略です。これにより、EU の規制に合わせてユーザーの絶対的なプライバシーが確保され、コストのかかるクラウド推論から経済性を根本的に転換し、ユーザーが真の AI 対応シリコンを求めるにつれて、大幅なハードウェア アップグレードのスーパーサイクルへの道が開かれます。スマートなオンデバイス「セマン​​ティック ルーター」を作成するには、モデルが 27B+ パラメーター スケールに達する必要があります。これを電話機で実現するには、PrismML の 3 値の重みなどの極端な量子化が必要です。ただし、ソフトウェアの世界では見落とされがちな重要なハードウェアの現実は、RAM に重みを当てはめることと、重みを移動することは同じではないということです。標準 LPDDR で 27B ターナリ モデルを実行すると、メモリ帯域幅に大きな制限が発生します。トークン生成ごとに SoC バスを介してギガバイトのデータを転送すると、NPU のサーマル スロットルが発生し、バッテリーが過度に消耗する可能性があります。これにより、重要な疑問が生じます。なぜ依然としてデータをコンピューティングに転送しているのでしょうか? AI 推論をメモリ内でネイティブに実行してみませんか?ベアメタル物理学を無視するアカデミックな PIM シミュレーションに不満を感じた私は、メモリ バスを完全にバイパスし、電荷共有を通じて COTS DRAM 内で直接 3 値 LLM 推論を実行するアーキテクチャである CaSA を開発しました。ソフトウェア量子化は優れた初期ステップであり、CaSA はブリッジを完成させるために必要な物理ハードウェア基板を提供します: https://github.com/pcdeni/CaSA

## Original Extract

The excitement surrounding PrismML’s 1-bit/ternary Bonsai models has the industry closely watching how smartphone giants, particularly Apple, will implement LLMs on edge devices.
Moving AI on-device is a brilliant and necessary strategy. It ensures absolute user privacy in alignment with EU regulations, fundamentally shifts the economics away from costly cloud inference, and paves the way for a significant hardware upgrade supercycle as users seek true AI-capable silicon. To create a smart on-device "Semantic Router," models need to reach the 27B+ parameter scale. Achieving this on a phone requires extreme quantization, such as PrismML’s ternary weights. However, a critical hardware reality often overlooked by the software world is that fitting the weights in RAM is not equivalent to moving them. Running a 27B ternary model on standard LPDDR encounters a significant memory bandwidth limitation. Transferring gigabytes of data across the SoC bus for each token generation can lead to thermal throttling of the NPU and excessive battery drain. This raises an important question: why are we still transferring data to the compute? Why not execute AI inference natively within the memory? Frustrated with academic PIM simulations that overlook bare-metal physics, I developed CaSA, an architecture that performs ternary LLM inference directly inside COTS DRAM through charge-sharing, completely bypassing the memory bus. Software quantization is a great initial step, and CaSA provides the physical hardware substrate needed to complete the bridge: https://github.com/pcdeni/CaSA


---
source: "https://deepgate.ai/blog/compiler"
hn_url: "https://news.ycombinator.com/item?id=48604625"
title: "Our edge AI compiler outperforms Google and vendor toolchains"
article_title: "Our edge AI compiler outperforms Google and vendor toolchains. | DeepGate"
author: "webstorms"
captured_at: "2026-06-20T00:56:49Z"
capture_tool: "hn-digest"
hn_id: 48604625
score: 1
comments: 0
posted_at: "2026-06-19T23:37:40Z"
tags:
  - hacker-news
  - translated
---

# Our edge AI compiler outperforms Google and vendor toolchains

- HN: [48604625](https://news.ycombinator.com/item?id=48604625)
- Source: [deepgate.ai](https://deepgate.ai/blog/compiler)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T23:37:40Z

## Translation

タイトル: 当社のエッジ AI コンパイラーは Google やベンダーのツールチェーンを上回るパフォーマンスを発揮します
記事のタイトル: 当社のエッジ AI コンパイラーは、Google やベンダーのツールチェーンよりも優れたパフォーマンスを発揮します。 |ディープゲート
説明: DeepGate コンパイラは、シリコン ベンダー全体で、Arm Cortex-M 上の Google の TFLM よりも RAM の使用量が最大 3 倍少なく、最大 2 倍高速に実行される静的バイナリを生成します。

記事本文:
当社のエッジ AI コンパイラーは、Google やベンダーのツールチェーンよりも優れたパフォーマンスを発揮します。 | DeepGate DeepGate ソリューション R&D チーム ブログ プラットフォーム お問い合わせ 2026 年 6 月 6 日
当社のエッジ AI コンパイラーは、Google やベンダーのツールチェーンよりも優れたパフォーマンスを発揮します。
DeepGate コンパイラは、シリコン ベンダー全体で、Arm Cortex-M 上の Google の TFLM よりも RAM の使用量が最大 3 倍少なく、最大 2 倍高速に実行される静的バイナリを生成します。
エッジ AI ツールは、大規模な GPU ベースのモデル用に構築されたコンパイラーやランタイムよりもまだ遅れています。ほとんどのマイクロコントローラーの導入は、Google の TensorFlow Lite for Microcontrollers (TFLM)、またはベンダー固有のバリアントに依存しています。このアプローチでは、大幅なパフォーマンスが未開発のまま残されていると考えられます。エッジでは、モデルが完全に適合するか、リアルタイムで実行されるか、または電力バジェットを満たしているかどうかは、効率によって決まります。私たちの目標は、最小のデバイスであるマイクロコントローラーから始めて、CPU および AI アクセラレーター用の最先端の AI コンパイラーを構築することです。
DeepGate コンパイラ (v0.15.0) をリリースします。これは量子化された .tflite モデルを最適化された推論バイナリにコンパイルします。このバイナリは、Arm Cortex-M デバイス上で Google の TFLM よりも最大 3 分の RAM の使用量を削減し、最大 2 倍の速度で実行します。マイクロコントローラー上の小型機械学習のベンチマーク スイートである MLPerf Tiny の評価では、アナログ デバイセズ、インフィニオン、Silicon Labs、STM のシリコン全体で TFLM よりも優れたパフォーマンスを示し、また、ハードウェア上では Infineon および Silicon Labs の独自のツールチェーンよりも優れたパフォーマンスを示しました。場合によっては、コンパイラを使用すると、メモリに収まらないモデルを実行できるようになりました。
独自のハードウェアでベンダーのツールチェーンを上回るパフォーマンスを発揮
私たちは、マイクロコントローラーでの機械学習の業界標準ベンチマークである MLPerf Tiny v1.4 ベンチマーク スイートで DeepGate コンパイラー (v0.15.0) を検証しました。シリコン ベンダー 4 社の 4 つのボードで実行し、結果を提出しました。

独立したレビュー用の MLPerf。このスイートには、キーワード スポッティング、ビジュアル ウェイク ワード、画像分類、異常検出のための代表的なエッジ AI ワークロードが含まれています。モデルを変更することなく、当社のコンパイラーは Google の TFLM よりも最大 3 分の 1 の RAM 使用量を削減し、最大 2 倍の速度で実行します。また、ベンダーのツールチェーンよりも優れたパフォーマンスを発揮します。EFR32MG24 の AI アクセラレータでは、Silicon Labs の TFLM Simplicity SDK よりも最大 3 倍の RAM 使用量の削減と 1.8 倍の高速推論を実現し、PSoC 6 では Infineon の Imagimob よりも最大 2 倍の高速推論を実現します。メモリの節約により、モデルが完全に適合するかどうかが決まります。アナログ デバイセズの MAX32655 (Visual Wake Words ベンチマーク) TFLM ではメモリが不足しましたが、DeepGate コンパイラで正常にコンパイルおよび実行されました。
以下のすべての比較を検討してください: ボードを切り替え、利用可能な場合はフレームワークを比較し、レイテンシと RAM 使用量を切り替えます。ここでは、テンソル アリーナにピーク スタック サイズを加えたものとして RAM を測定しました。
DeepGate は最大 1.9 倍高速に動作します
STMicroelectronics の ST Edge AI は依然として高い競争力を維持しています。バランスのとれたコンパイル設定に反して、キーワード スポッティング推論の高速化 (1.1 倍の高速化) と異常検出時の RAM 使用量の削減 (1.6 倍の RAM 削減) を実現しますが、他のワークロードは今後のリリースでも引き続き焦点となります。
意味のある効率向上には複数の次元にわたる最適化が必要であるため、コンパイラーをそれらすべてにわたって最適化しました。コンパイラーはランタイム・インタープリターではなく静的バイナリにコンパイルし、コンパイル時にグラフ全体のメモリー割り当てを計画し、ハードウェアインザループ・テストを通じて調整されたカスタム・アセンブリ・ルーチンを含む、Arm の標準 CMSIS-NN カーネルを超えたハードウェア対応カーネル最適化を適用します。
DeepGate コンパイラの違い
私たちは最適化ロードマップの初期段階にあり、大きなチャンスを抱えています。

主にメモリプランニングやカーネル最適化などの分野を担当しています。また、スパース ネットワーク、低ビット量子化、Transformer モデルの効率的なアテンション メカニズムなど、既存のエッジ AI ツールチェーンでは十分に機能していないアプローチのサポートも拡大しています。さらに先を見据えて、私たちは DeepGate の新しい ML ビルディング ブロックを中心にコンパイラを共同設計しています。これにより、コストのかかる行列乗算への依存が軽減され、インプレース計算の利用が可能になり、制約のあるハードウェアに根本的に適したモデルへの道が開かれます。
現在、当社のコンパイラは Arm Cortex-M CPU と厳選された組み込み AI アクセラレータをターゲットにしており、そのサポートを積極的に拡大しています。あなたにとってどのターゲットが最も重要なのかをぜひお聞かせください。アップデートにサインアップしたり、プラットフォームへのアクセスをリクエストしたり、次にサポートしてほしいデバイスがある場合はご連絡ください。

## Original Extract

DeepGate compiler produces static binaries that use up to 3× less RAM and run up to 2× faster than Google's TFLM on Arm Cortex-M, across silicon vendors.

Our edge AI compiler outperforms Google and vendor toolchains. | DeepGate DeepGate Solutions R&D Team Blog Platform Get in touch June 6, 2026
Our edge AI compiler outperforms Google and vendor toolchains.
DeepGate compiler produces static binaries that use up to 3× less RAM and run up to 2× faster than Google's TFLM on Arm Cortex-M, across silicon vendors.
Edge AI tooling still lags behind the compilers and runtimes built for large GPU-based models. Most microcontroller deployments rely on Google’s TensorFlow Lite for Microcontrollers (TFLM), or vendor-specific variants – an approach we believe leaves significant performance untapped. At the edge, efficiency determines whether a model fits at all, runs in real time, or meets its power budget. Our goal is to build the leading edge AI compiler for CPUs and AI accelerators, starting with the smallest devices: microcontrollers.
We’re releasing the DeepGate compiler (v0.15.0), which compiles quantized .tflite models into optimized inference binaries that use up to 3× less RAM and run up to 2× faster than Google’s TFLM on Arm Cortex-M devices. In our MLPerf Tiny evaluation, a benchmark suite for tiny machine learning on microcontrollers, it outperformed TFLM across silicon from Analog Devices, Infineon, Silicon Labs, and STM, while also outperforming Infineon’s and Silicon Labs’ own toolchains on their hardware. In some cases, our compiler enabled models to run that otherwise would not fit in memory.
Outperforming vendor toolchains on their own hardware
We’ve validated the DeepGate compiler (v0.15.0) on the MLPerf Tiny v1.4 benchmark suite, the industry-standard benchmark for machine learning on microcontrollers. We ran it across four boards from four silicon vendors, with results submitted to MLPerf for independent review. The suite includes representative edge AI workloads for keyword spotting, visual wake words, image classification, and anomaly detection. Without modifying the models, our compiler uses up to 3× less RAM and runs up to 2× faster than Google’s TFLM. It also outperforms vendor toolchains: delivering up to 3× lower RAM usage and 1.8× faster inference than Silicon Labs’ TFLM Simplicity SDK on the EFR32MG24’s AI accelerator, and up to 2× faster inference than Infineon’s Imagimob on the PSoC 6. Our memory savings determine whether a model fits at all: on Analog Devices’ MAX32655, the Visual Wake Words benchmark ran out of memory under TFLM but compiled and executed successfully with the DeepGate compiler.
Explore every comparison below: switch boards, compare frameworks where available, and toggle between latency and RAM usage. Here, we measured RAM as the tensor arena plus peak stack size.
DeepGate runs up to 1.9× faster
ST Edge AI from STMicroelectronics remains highly competitive. Against its balanced compilation setting, we deliver faster keyword spotting inference (1.1× faster) and lower RAM usage on anomaly detection (1.6× less RAM), while other workloads remain a focus for upcoming releases.
Meaningful efficiency gains require optimization across multiple dimensions, so we optimized our compiler across all of them: it compiles to static binaries rather than a runtime interpreter, plans whole-graph memory allocation at compile time, and applies hardware-aware kernel optimizations beyond Arm’s standard CMSIS-NN kernels, including custom assembly routines tuned through hardware-in-the-loop testing.
What makes the DeepGate compiler different
We’re still early in our optimization roadmap, with significant opportunities remaining in areas such as memory planning and kernel optimization. We’re also expanding support for approaches that existing edge AI toolchains often underserve, including sparse networks, lower-bit quantization, and efficient attention mechanisms for Transformer models. Looking further ahead, we are co-designing our compiler around DeepGate’s novel ML building blocks, which reduce reliance on costly matrix multiplications and enable greater use of in-place computation – paving the way for models fundamentally better suited to constrained hardware.
Today our compiler targets Arm Cortex-M CPUs and selected embedded AI accelerators, and we’re actively expanding that support. We’d love to hear which targets matter most to you. Sign up for updates, request platform access, or get in touch if there’s a device you’d like us to support next.

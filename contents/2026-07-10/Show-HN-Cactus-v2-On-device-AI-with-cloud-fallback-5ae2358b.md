---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48864459"
title: "Show HN: Cactus v2 – On-device AI with cloud fallback"
article_title: ""
author: "rshemet"
captured_at: "2026-07-10T20:14:04Z"
capture_tool: "hn-digest"
hn_id: 48864459
score: 1
comments: 0
posted_at: "2026-07-10T19:57:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cactus v2 – On-device AI with cloud fallback

- HN: [48864459](https://news.ycombinator.com/item?id=48864459)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T19:57:18Z

## Translation

タイトル: Show HN: Cactus v2 – クラウド フォールバックを備えたオンデバイス AI
HN テキスト: こんにちは、HN、Cactus ( https://github.com/cactus-compute/cactus ) からの Roman と Henry です。オンデバイス推論プラットフォームへの最大のアップグレードを出荷しました: - 推論の実行をクラウドに引き渡すための組み込みモデルの信頼ベースのルーティング
- あらゆる PyTorch モデル用のコンバーター
- ロスレス 4 ビット量子化 (GitHub README で評価)
- 互換性のあるデバイスでの GPU アクセラレーション (Apple Metal 以降)
- 最小限のRAMフットプリント
- iOS、Android、Mac、DGX Spark、Raspberry Pi などのあらゆる Arm デバイスで実行可能 全体として、Gemma 4 E2B クラス モデルは M5 Max 上で 169 トーク/秒で実行され、FP16 からの精度低下なしで 2.7 GB のディスク領域を占有し、1.3 GB の RAM を使用し、必要に応じてクラウド モデルにヘルプを要求します。私たちが 18 か月前に始めた問題: 推論エンジンはデータセンター向けに構築されていますが、コンシューマ ハードウェアは物理的に異なります。RAM を OS と共有し、熱スロットルが発生し、同じモデルでもハードウェアが異なると動作が異なります。そこで、リソースに制約のあるデバイス向けにランタイムを最初から作成しました。それ以来、Cactus は毎週何百万もの推論実行を処理し、毎月数万人のアクティブな開発者を処理するまでに成長しました。 Cactus を運用アプリにデプロイすることで得た最大の教訓は、ローカル モデルはワークロードの 90% を処理できるものの、その 10% のギャップはまだ運用準備ができていないことを意味するということです。ユーザーの解決策は、カスタムのクラウド フォールバック ロジックを構築することでした。 Cactus v2 では次の点が修正されています。 クラウド フォールバックに対する私たちのアプローチは、内部アクティベーションを読み取り、信頼性信号を出力するモデルの重みへのプローブをポストトレーニングすることです。このように、ルーティングは、モデルの前にあるプロンプト分類子ではなく、モデル内で行われます。私たちは、これが複数ターンのエージェント作業にとって重要であると考えています。

モデルは、どのターンがローカルで処理できるほど簡単で、どのターンが難しいかを認識し、クラウドに引き渡す必要があります。本日出荷されるのは、構成可能なエスカレーション エンドポイント (Gemini、Claude、OpenAI 互換、または独自のエンドポイント) に対する Gemma-4 E2B のシングル ターン ルーティングです。私たちの次のターゲットは、マルチターン エージェント作業用のハイブリッド ネイティブ モデルです。これは本当に未解決の問題です。現在のプローブインザウェイトは有望な結果を示しており、最初のモデルのバリアントを間もなくリリースする予定です。ハイブリッド バリアントは Gemma の派生製品であり、HF カードに記載されている Gemma 条件に基づいてリリースされます。ハイブリッド ルーティングに加えて、ランタイムには 4 ビットでロスレスの SOTA 量子化があり、メモリ マップの重みによって RAM フットプリントが削減され、Python、Rust、React Native、Swift、Kotlin バインディングを使用してクロスプラットフォームで実行されます。免責事項: Cactus はソースとして配布されており、個人使用および小規模企業には無料です。それ以上の商用ライセンス (Docker スタイル ライセンス)。 GitHub: https://github.com/cactus-compute/cactus または `brew install cactus-compute/cactus/cactus` で開始できます。

## Original Extract

Hi HN, Roman and Henry here from Cactus ( https://github.com/cactus-compute/cactus ). We just shipped the biggest upgrade to our on-device inference platform: - Built-in model confidence-based routing to hand off inference runs to the cloud
- Converter for any PyTorch model
- Lossless 4-bit quantization (evals on our GitHub README)
- GPU acceleration on compatible devices (starting with Apple Metal)
- Minimal RAM footprint
- Runs on any Arm device: iOS, Android, Mac, DGX Spark, Raspberry Pi, and more All in, a Gemma 4 E2B class model runs at 169 tok/sec on M5 Max, takes 2.7GB disk space with no accuracy degradation from FP16, uses 1.3GB of RAM, and requests help from cloud models when needed. The problem we started with eighteen months ago: inference engines are built for datacenters, but consumer hardware has different physics: you share RAM with the OS, you get thermally throttled, and the same model behaves differently on different hardware. So we wrote a runtime from scratch for resource-constrained devices. Since then, Cactus has grown to process millions of weekly inference runs, and tens of thousands of monthly active developers. Our biggest learning from deploying Cactus in production apps is that while local models can handle 90% of workloads, that 10% gap means they're still not production-ready. Our users' fix was to build custom cloud fallback logic. Cactus v2 fixes that: Our approach to cloud fallback is to post-train a probe into the model's weights that reads its internal activations and emits a confidence signal. This way, the routing happens inside the model rather than in a prompt classifier sitting in front of it. We believe this is critical for multi-turn agentic work, where the model should know which turns are easy enough to be handled locally, and which are hard - and get handed off to the cloud. What ships today is single-turn routing for Gemma-4 E2B against a configurable escalation endpoint (Gemini, Claude, OpenAI-compatible, or your own endpoint). Our next target is hybrid-native models for multi-turn agentic work. This is the genuinely unsolved problem. Our current probe-in-the-weights is showing promising results and we look to release the first model variants soon. The hybrid variants are Gemma derivatives, released under the Gemma terms noted on the HF cards. In addition to the hybrid routing, the runtime has SOTA quantization, which is lossless at 4bit, memory maps weights to decrease RAM footprint and runs cross-platform, with Python, Rust, React Native, Swift, and Kotlin bindings. Disclaimer: Cactus is distributed as source-available - free for personal use and small companies; commercial license above that (Docker-style license). You can get started on our GitHub: https://github.com/cactus-compute/cactus or by `brew install cactus-compute/cactus/cactus`.


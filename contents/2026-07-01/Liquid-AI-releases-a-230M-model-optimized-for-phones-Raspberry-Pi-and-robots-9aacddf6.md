---
source: "https://www.liquid.ai/blog/lfm2-5-230m"
hn_url: "https://news.ycombinator.com/item?id=48743203"
title: "Liquid AI releases a 230M model optimized for phones, Raspberry Pi, and robots"
article_title: "LFM2.5-230M: Built to Run Anywhere — Blog — Liquid AI"
author: "mpfect"
captured_at: "2026-07-01T07:25:24Z"
capture_tool: "hn-digest"
hn_id: 48743203
score: 1
comments: 0
posted_at: "2026-07-01T07:03:12Z"
tags:
  - hacker-news
  - translated
---

# Liquid AI releases a 230M model optimized for phones, Raspberry Pi, and robots

- HN: [48743203](https://news.ycombinator.com/item?id=48743203)
- Source: [www.liquid.ai](https://www.liquid.ai/blog/lfm2-5-230m)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T07:03:12Z

## Translation

タイトル: Liquid AI、電話、Raspberry Pi、ロボットに最適化された 2 億 3000 万モデルをリリース
記事のタイトル: LFM2.5-230M: どこでも実行できるように構築 — ブログ — Liquid AI
説明: Liquid AI のこれまでで最小のモデルである LFM2.5-230M を紹介します。微調整、エッジ展開、ツールの使用、およびデータ抽出のための高速でオープンウェイトの基盤モデルです。

記事本文:
LFM2.5-230M: どこでも実行できるように構築 — ブログ — Liquid AI 製品
ニュース / モデル モデル 2026 年 6 月 25 日 LFM2.5-230M: どこでも実行できるように構築
本日、当社はこれまでの最小モデルである LFM2.5-230M をリリースします。これは、開発者がエージェント ワークフローで微調整して展開するための高速かつ軽量の基盤です。 LFM2 アーキテクチャに基づいて構築されており、非常に高速な推論を実現し、クラウド GPU から低コスト CPU まであらゆる場所で実行できます (Galaxy S25 Ultra ではデコード速度 213 tok/s、Raspberry Pi 5 では 42 tok/s)。サイズが小さいにもかかわらず、ツールの使用とデータ抽出タスクにおいて驚くほどの能力を発揮します。
ベース (LFM2.5-230M-Base) モデルとトレーニング後モデル (LFM2.5-230M) は、Hugging Face で現在入手可能です。ローカルで実行して微調整する方法については、ドキュメントをご覧ください。
モデルは、32K コンテキスト拡張フェーズを含む 19T トークン用に事前トレーニングされました。独自のダウンストリーム アプリケーションを対象とする開発者の柔軟性を維持するように設計された軽量のポストトレーニング レシピを適用します。
レシピは 3 つの段階で構成されます: (1) LFM2.5-350M からの蒸留による教師あり微調整、(2) 直接優先最適化、および (3) マルチドメイン強化学習。最後のチェックポイントでは、大きなモデルとの競争力を維持しながら、すぐに使用できる強力な機能と下流の特殊化への適応性のバランスをとります。
進行中の作業の初期段階として、LFM2.5-230M を Unitree G1 ヒューマノイド ロボットに導入し、オンボードの NVIDIA Jetson Orin 上で完全にオンデバイスで実行しました。ここでモデルはスキル選択レイヤーとして機能します。モデルは単一の自然言語命令を受け取り、NVIDIA の SONIC フレームワークによって提供される事前トレーニング済みの低レベル スキルを呼び出す一連のツール呼び出しに分解します。このタスクを簡単に微調整した後、モデルは次のような自由形式のコマンドを変換します。
「じっとしてて

2 秒間前方に歩き、秒速 1 メートルで 3 メートル歩き、片足で片足を前方に 5 秒間保持し、秒速 0.5 メートルで 3 メートル後ろ向きに歩きます。」
目標速度でのタイミングを計ったウォーキングや片足膝立ちなどのスキルを、構造化された複数ステップの計画に組み込んでいきます。この段階では動作は意図的に単純ですが、これは説得力のあるシグナルであると考えています。2 億 3,000 万のパラメーターのモデルを迅速に微調整し、デバイス上に展開して、ヒューマノイドの自然言語制御インターフェイスとして機能させることができます。
私たちは、コア機能と応用タスクの両方をカバーする 10 のベンチマークにわたって LFM2.5-230M を評価しました。そのサイズにもかかわらず、知識 (GPQA Diamond、MMLU-Pro)、命令に従っている (IFEval、IFBench、Multi-IF)、データ抽出 (CaseReportBench)、およびツールの使用 (BFCLv3、BFCLv4、τ²-Bench Telecom and Retail) にわたる、2 倍以上の規模のモデルと競合し、多くの場合、それを上回っています。
このため、LFM2.5-230M は、大規模なデータ抽出パイプラインや軽量のオンデバイス エージェント ワークロードを強化するための理想的なソリューションになります。ただし、そのコンパクトなサイズを考慮すると、高度な数学、コード生成、クリエイティブな執筆など、推論が重要なワークロードにはお勧めできません。
LFM2.5-230M には、推論エコシステム全体にわたる初日サポートが付属しています。
llama.cpp — 効率的なエッジ推論のための GGUF チェックポイント
MLX — Apple Silicon向けに最適化された推論
vLLM — 実稼働スループットを実現する GPU 高速化サービス
SGLang — 実稼働スループットのための GPU 高速化サービス
ONNX — 多様なアクセラレータにわたるクロスプラットフォーム推論
CPU推論。効率的な LFM2 アーキテクチャのおかげで、LFM2.5-230M は、SSM ハイブリッドやゲート デルタ ネットワークなどの同様のサイズのモデルよりも大幅に高速です。 Raspberry Pi 5 と Qualcomm Snapdragon Gen4 (Samsung Galaxy S25 Ultra) の両方で、

メモリ使用量を最小限に抑えながら、クラス最高のプリフィルおよびデコード スループットを実現します。各プラットフォームでプリフィルを最大化するためにデバイスごとのフラッシュ アテンション フラグを調整します。Raspberry Pi 5 では有効 (-fa 1)、Snapdragon Gen4 では無効 (-fa 0) になります。
GPU推論。実稼働グレードのエンタープライズ展開向けに、非常に低遅延のサービスを提供する内部 GPU 推論スタックも開発しました。 SGLang で実行されている他の小規模モデルに対してベンチマークを行ったところ、すべての同時実行レベルにわたって、LFM2.5 モデルはエンドツーエンドのレイテンシーが大幅に短縮されました。
Hugging Face で入手可能な LFM2.5-230M および LFM2.5-230M-Base を使用して、今すぐ構築を始めてください。
LFM2.5 により、私たちはどこでも実行できる AI というビジョンを実現します。これらのモデルは次のとおりです。
オープンウェイト — 制限なしでダウンロード、微調整、展開します
初日から迅速 — Apple、AMD、Qualcomm、および Nvidia ハードウェアにわたる llama.cpp、NexaSDK、MLX、および vLLM のネイティブ サポート
完全なファミリー — カスタマイズ用のベースモデルから特殊なオーディオおよびビジョンのバリアントまで、1 つのアーキテクチャで多様なユースケースをカバーします
エッジ AI の未来がここにあります。あなたが何を構築するのか楽しみです。
引用については、次の参考文献または BibTeX を使用してください。
Liquid AI、「LFM2.5-230M: どこでも実行できるように構築」、Liquid AI ブログ、2026 年 6 月。
@article{liquidAI2026230M,
著者 = {Liquid AI}、
タイトル = {LFM2.5-230M:
どこでも実行できるように構築},
ジャーナル = {Liquid AI ブログ}、
年 = {2026}、
注 = {www.liquid.ai/blog/lfm2-5-230m}
}
報道関係者からのお問い合わせについては、press@liquid.ai までメールでお問い合わせください。
インテリジェンスがどこにでも存在できるように、最速、最も計算効率が高く、有能な基盤モデルを構築します。
リキッドエッジ AI プラットフォーム (LEAP)
自動化された基礎モデル設計

## Original Extract

Meet LFM2.5-230M, Liquid AI’s smallest model yet: a fast, open-weight foundation model for fine-tuning, edge deployment, tool use, and data extraction.

LFM2.5-230M: Built to Run Anywhere — Blog — Liquid AI Products
News / Models Models JUN 25, 2026 LFM2.5-230M: Built to Run Anywhere
Today, we're releasing LFM2.5-230M , our smallest model yet. It’s a fast, lightweight foundation for developers to fine-tune and deploy in agentic workflows. Built on the LFM2 architecture, it delivers exceptionally fast inference and runs everywhere, from cloud GPUs to low-cost CPUs (213 tok/s decode speed on Galaxy S25 Ultra, 42 tok/s on a Raspberry Pi 5). Despite its small size, it’s surprisingly capable at tool use and data extraction tasks.
The base (LFM2.5-230M-Base) and post-trained (LFM2.5-230M) models are available today on Hugging Face . Check out our docs on how to run and fine-tune them locally.
The model was pre-trained for 19T tokens, including a 32K context extension phase. We apply a lightweight post-training recipe designed to preserve flexibility for developers targeting their own downstream applications.
The recipe consists of three stages: (1) supervised fine-tuning with distillation from LFM2.5-350M, (2) direct preference optimization, and (3) multi-domain reinforcement learning . The final checkpoint balances strong out-of-the-box capabilities with adaptability to downstream specialization, while remaining competitive with larger models.
As an early look at ongoing work, we deployed LFM2.5-230M on a Unitree G1 humanoid robot, running entirely on-device on its onboard NVIDIA Jetson Orin. Here the model acts as a skill-selection layer: it takes a single natural-language instruction and decomposes it into a sequence of tool calls that invoke pre-trained low-level skills provided by NVIDIA's SONIC framework. After a quick fine-tune for this task, the model turns a free-form command such as
"Hold still for 2 seconds, then walk forward at 1 meter per second for 3 meters, hold a forward one-leg kneel for 5 seconds, and walk backward at 0.5 meters per second for 3 meters"
into a structured, multi-step plan, chaining skills like timed walking at a target velocity and a one-legged kneel. While the behaviors are deliberately simple at this stage, we think it's a compelling signal: a 230M-parameter model can be quickly fine-tuned and deployed on-device to serve as the natural-language control interface for a humanoid.
We evaluated LFM2.5-230M across ten benchmarks covering both core capabilities and applied tasks. Despite its size, it competes with and often beats models more than twice as large , spanning knowledge (GPQA Diamond, MMLU-Pro), instruction following (IFEval, IFBench, Multi-IF), data extraction (CaseReportBench), and tool use (BFCLv3, BFCLv4, τ²-Bench Telecom and Retail).
This makes LFM2.5-230M an ideal solution to power large-scale data extraction pipelines or lightweight on-device agentic workloads. However, given its compact size, we do not recommend it for reasoning-heavy workloads such as advanced math, code generation, or creative writing.
LFM2.5-230M ships with day-one support across the inference ecosystem:
llama.cpp — GGUF checkpoints for efficient edge inference
MLX — Optimized inference for Apple Silicon
vLLM — GPU-accelerated serving for production throughput
SGLang — GPU-accelerated serving for production throughput
ONNX — Cross-platform inference across diverse accelerators
CPU inference. Thanks to the efficient LFM2 architecture, LFM2.5-230M is considerably faster than similar-sized models, including SSM hybrids and Gated Delta Networks. On both a Raspberry Pi 5 and a Qualcomm Snapdragon Gen4 (Samsung Galaxy S25 Ultra), it delivers the highest prefill and decode throughput in its class while keeping the smallest memory footprint. We tune the flash-attention flag per device to maximize prefill on each platform: enabled (-fa 1) on the Raspberry Pi 5 and disabled (-fa 0) on the Snapdragon Gen4.
GPU inference. For production-grade enterprise deployments, we have also developed an internal GPU inference stack that delivers extremely low-latency serving. We benchmark it against other small models running on SGLang, and across all concurrency levels, LFM2.5 models achieve considerably lower end-to-end latency.
Start building today with LFM2.5-230M and LFM2.5-230M-Base, available on Hugging Face.
With LFM2.5, we're delivering on our vision of AI that runs anywhere. These models are:
Open-weight — Download, fine-tune, and deploy without restrictions
Fast from day one — Native support for llama.cpp, NexaSDK, MLX, and vLLM across Apple, AMD, Qualcomm, and Nvidia hardware
A complete family — From base models for customization to specialized audio and vision variants, one architecture covers diverse use cases
The edge AI future is here. We can't wait to see what you build.
For citations, please use the following reference or BibTeX:
Liquid AI, "LFM2.5-230M: Built to Run Anywhere", Liquid AI Blog, Jun 2026.
@article{liquidAI2026230M,
author = {Liquid AI},
title = {LFM2.5-230M:
Built to Run Anywhere},
journal = {Liquid AI Blog},
year = {2026},
note = {www.liquid.ai/blog/lfm2-5-230m}
}
For press inquiries, email us at press@liquid.ai
Building the fastest, most compute-efficient, and capable foundation models so intelligence can live anywhere.
Liquid Edge AI Platform (LEAP)
Automated Foundation Model Design

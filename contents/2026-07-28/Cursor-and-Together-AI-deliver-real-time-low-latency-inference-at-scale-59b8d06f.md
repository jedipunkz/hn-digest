---
source: "https://www.together.ai/customers/cursor"
hn_url: "https://news.ycombinator.com/item?id=49091417"
title: "Cursor and Together AI deliver real-time, low-latency inference at scale"
article_title: "Learn how Cursor partnered with Together AI to deliver real-time, low-latency inference at scale"
author: "inferhaven"
captured_at: "2026-07-28T23:51:33Z"
capture_tool: "hn-digest"
hn_id: 49091417
score: 1
comments: 0
posted_at: "2026-07-28T23:30:24Z"
tags:
  - hacker-news
  - translated
---

# Cursor and Together AI deliver real-time, low-latency inference at scale

- HN: [49091417](https://news.ycombinator.com/item?id=49091417)
- Source: [www.together.ai](https://www.together.ai/customers/cursor)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T23:30:24Z

## Translation

タイトル: Cursor and Together AI は、大規模なリアルタイムの低遅延推論を実現します
記事のタイトル: Cursor が Together AI と提携して、リアルタイムで低遅延の推論を大規模に提供する方法を学ぶ
説明: AI が Cursor と連携して、エディター内エージェントの高速性と信頼性を維持するリアルタイム推論スタックを構築しました。彼らは、NVIDIA Blackwell (B200/GB200) を製品化し、ARM ホスト、カーネル、FP4/TensorRT 量子化を調整して、低レイテンシーと迅速なモデル ロールアウトを実現しました。

記事本文:
Cursor が Together AI と提携して、大規模なリアルタイムで低遅延の推論を実現する方法をご覧ください。 Webflow 分析/最適化追跡ブリッジ -->
💰 シリーズ C を発表します。インテリジェンスは高価ではなく豊富であるべきです →
🤝 Together AI & Y Combinator が初の専用 YC GPU クラスターを提供するパートナーシップを発表 →
⚡ オンデマンド B200 が Together GPU クラスターで利用可能になりました →
🚀 効率的な推論のために MiniMax-M3 を提供中 →
API としての高パフォーマンス推論
SLA を備えたトークンベースの容量
上位のオープンソース モデルを探索する
大規模な信頼性の高い GPU クラスター
フロンティア規模のカスタムインフラストラクチャ
AIの開発環境を構築する
モデルの重みとデータを安全に保存
トップのオープンソース モデルを微調整する
プロダクションAIのシステム研究
Together AI の技術ドキュメント
実践的な実装ガイド
実稼働用の音声エージェントを構築する
質問に対する答えを見つける
API としての高パフォーマンス推論
SLA を備えたトークンベースの容量
上位のオープンソース モデルを探索する
大規模な信頼性の高い GPU クラスター
フロンティア規模のカスタムインフラストラクチャ
AIの開発環境を構築する
モデルの重みとデータを安全に保存
トップのオープンソース モデルを微調整する
プロダクションAIのシステム研究
Together AI の技術ドキュメント
実践的な実装ガイド
実稼働用の音声エージェントを構築する
質問に対する答えを見つける
すべての顧客事例 Cursor が Together AI と提携して、リアルタイムで低遅延の推論を大規模に提供する方法を学びましょう
AI ネイティブの開発者ツール会社
データセンター全体での生産
Cursor は AI を活用したコーディング プラットフォームで、エディター内のエージェントはエディターのフィードバック ループ内で応答する必要があり、同時実行時の最悪の場合のレイテンシーが必須の要件となります。 Cursor は Together AI と提携して、NVIDIA Blackwell (GB200 NVL72/H) に本番推論を導入しました

GX B200)、スタックをエンドツーエンドで最適化し、量子化 (TensorRT-LLM + NVFP4) を介して、新しい重みから本番環境のようなテスト エンドポイントまでの反復可能なパスを確立します。
Cursor は、継続的なバックグラウンド インテリジェンスを備えた AI を活用したコーディング プラットフォームで、エージェント コーディング モデルをトレーニングして本番環境に出荷する研究チームによって構築されています。開発者が入力すると、Cursor はコード コンテキストのライブ モデルを維持し、編集を予測し、コードをリファクタリングし、リアルタイムでコンテキストを管理します。
そのエクスペリエンスを提供するには、編集者のフィードバック ループ内での応答が必要です。この制約により、サービス提供の問題はバッチ形式からリアルタイムの低遅延推論に移行します。 Cursor は、AI ネイティブ クラウドである Together AI と提携して、NVIDIA Blackwell アーキテクチャを使用し、厳しい遅延目標を満たすように推論スタックを調整して、このループのインフラストラクチャを構築しました。
エディタ内でレイテンシが異なる理由
エディタ内エージェントは、開発者がアクティブに編集している間に出力を生成するため、提案が生成に使用されたモデルと同じローカル コンテキストに提案が到達するかどうかはタイミングによって決まります。開発者がコードの別の領域に移動すると、出力はサポートするはずだった状態と一致しなくなることがよくあります。
Cursor では、開発者が作業を続けている間に、これらのエージェントが問題をデバッグし、機能を生成し、リファクタリングを実行します。そのワークロードには、同時実行下での予測可能な最悪の場合のレイテンシ、重複するリクエスト全体での一貫したコンテキスト処理、および持続的な負荷の下での安定した動作が必要です。
NVIDIA Blackwell でのエンジニアリング
レイテンシ バジェットを大規模に満たすことで、チームは、より高速なサービスをサポートするために、より高いメモリ帯域幅とテンソル スループットを備えた NVIDIA Blackwell GB200 NVL72 および NVIDIA HGX™ B200 を採用しました。ライフサイクルの早い段階で実稼働ワークロードをデプロイすることは、信頼性を高めることを意味します

ハードウェア、ファームウェア、ホスト ソフトウェア、サービス層など、スタック全体にわたる最適化を実現します。
NVIDIA Blackwell によるフロンティア インフラストラクチャ: Together AI は Cursor と協力して、NVIDIA Blackwell の初期展開を実現し、協力しました。 Cursor にとって、ハードウェアへの早期アクセスは製品の利点であり、Togetter エンジニアは、この新しいフロンティア インフラストラクチャを提供するために、迅速なアップグレードと交換に取り組みました。これらの努力により、Cursor が使用できる新しいチップが迅速かつ確実に提供されました。
ARM ホストでのフル スループット: GB200 NVL72 は、ARM 命令セット上の NVIDIA Grace™ CPU と GPU を組み合わせます。高性能推論エコシステムの多くは、x86 ホストを前提としています。推論スタックを ARM に移植するには、GB200 NVL72 用のカーネルとホストレベルのチューニングが必要でした。
Blackwell Tensor Core のカスタム カーネル: Blackwell は、低精度フォーマット用に最適化された新しい Tensor Core 命令を導入します。 AI は一緒に Blackwell 用のカーネルを構築し、これらの命令を直接ターゲットにして、ハードウェアのスループットをより多く獲得しました。
NVIDIA GB200 NVL72 全体での効率的な並列処理: GB200 NVL72 は、72 個の NVIDIA Blackwell GPU を全対全トポロジで接続します。そのドメイン全体にモデルを分散すると、チップ間の通信と同期のオーバーヘッドが追加されます。共同で GB200 NVL72 用の並列メッシュを設計したため、調整コストは制限されたままとなり、計算の向上は推論まで引き継がれました。
計量から生産までのサイクルを短縮する
Cursor の研究チームは、独自のデータとコーディング ワークフローの対象を絞った最適化を組み合わせてモデルを内部でトレーニングし、新しい重み候補を生成します。 Together AI とのコラボレーションにより、これらの重みを運用環境のようなエンドポイントに移動して即時にテストするための反復可能なパスが確立されました。
その過程における重要なステップは量子化です。リアルt

ime サービングは厳しいメモリとコンピューティング バジェットで実行され、量子化はより少ないビットで重みを表現することで両方を削減します。コーディングのコンテキストでは、品質の低下が微妙な論理ミスや構文エラーとして表面化する可能性があるため、量子化ではレイテンシーとコストを改善しながら出力品質を維持する必要があります。
AI は一緒に、Blackwell 上の NVIDIA TensorRT LLM と NVFP4 を中心に構築された量子化パイプラインを実装しました。これにより、積極的な圧縮と、Cursor のコーディング モデルに必要な品質バーの間で針を通すことができます。 Cursor が新しい候補を生成すると、Together はそれを量子化し、検証し、数日以内にテスト エンドポイントを起動します。 Cursor は内部評価スイートを実行し、カットオーバーを完了する前に運用トラフィックの下で A/B テストを実施します。カットオーバーは検証とライブトラフィックチェックによってゲートされます。
次は何ですか — レイテンシーからスループットまで
Cursor の実稼働デプロイメントは、複数のデータセンターにわたる NVIDIA Blackwell GPU 上で実行されます。 NVIDIA GB200 NVL72 は推論を処理し、Cursor の研究チームが新しい重みを出荷する際にインフラストラクチャがモデルの反復をサポートします。
実稼働環境でそのレイテンシ パスが実行されると、焦点はスループットと使用率に移ります。 Cursor と Together AI は、使用量の増加に応じて GPU あたりの経済性を向上させるために、NVIDIA Blackwell プラットフォーム上に高スループットのエンドポイントを構築しています。
Together AI のインスタント クラスターにより、Latent Health が GPT-4 を上回る臨床 AI を構築できるようになります
Together Code Sandbox を使用して HeroUI Chat を 10 倍高速に起動した方法
Scaled Cognition が AI GPU クラスター上で APT-1 をトレーニングする方法
AWS から Together 専用エンドポイントへ: 推論の柔軟性を高める Arcee AI の旅
AI Companions のスケーリング: Dippy AI がどのようにして専用エンドポイントを併用して 400 万トークン/分を超えたか
専用 AI インフラを使用して世界クラスのタイ語モデルを構築

構造
最適化されたトレーニングとモデル形成から大規模な本番推論まで
© 2026 一緒に AI。無断転載を禁じます。

## Original Extract

Together AI teamed with Cursor to build the real-time inference stack that keeps in-editor agents fast and reliable. They productionized NVIDIA Blackwell (B200/GB200), tuning ARM hosts, kernels, and FP4/TensorRT quantization for low latency and rapid model rollouts.

Learn how Cursor partnered with Together AI to deliver real-time, low-latency inference at scale Webflow Analyze/Optimize tracking bridge -->
💰 Announcing our Series C. Intelligence should be abundant, not expensive →
🤝 Together AI & Y Combinator announce partnership to deliver the first dedicated YC GPU cluster →
⚡ On-demand B200s now available on Together GPU Clusters →
🚀 Now serving MiniMax-M3 for efficient inference →
High-performance inference as APIs
Token-based capacity with SLAs
Explore the top open-source models
Reliable GPU clusters at scale
Custom infrastructure at frontier scale
Build development environments for AI
Store model weights & data securely
Fine-tune top open-source models
Systems research for production AI
Technical docs for Together AI
Practical implementation guides
Build voice agents for production
Find answers to your questions
High-performance inference as APIs
Token-based capacity with SLAs
Explore the top open-source models
Reliable GPU clusters at scale
Custom infrastructure at frontier scale
Build development environments for AI
Store model weights & data securely
Fine-tune top open-source models
Systems research for production AI
Technical docs for Together AI
Practical implementation guides
Build voice agents for production
Find answers to your questions
All customer stories Learn how Cursor partnered with Together AI to deliver real-time, low-latency inference at scale
AI-native developer tools company
Production across data centers
Cursor is an AI-powered coding platform where in-editor agents must respond inside the editor feedback loop, making worst-case latency under concurrency a hard requirement. Cursor partnered with Together AI to deploy production inference on NVIDIA Blackwell (GB200 NVL72 / HGX B200), optimize the stack end-to-end, and stand up a repeatable path from new weights to a production-like test endpoint via quantization (TensorRT-LLM + NVFP4).
Cursor is an AI-powered coding platform with continuous background intelligence, built by a research team that trains agentic coding models and ships them into production. As developers type, Cursor maintains a live model of the code context — predicting edits, refactoring code, and managing context in real time.
Delivering that experience requires responses inside the editor's feedback loop. That constraint shifts the serving problem from batch-style to real-time, low-latency inference. Cursor partnered with Together AI, the AI Native Cloud, to build infrastructure for this loop — using the NVIDIA Blackwell architecture and tuning the inference stack to meet strict latency targets.
Why latency is different inside an editor
In-editor agents generate outputs while the developer is actively editing, so timing determines whether a suggestion lands in the same local context the model used to generate it. Once the developer moves to a different region of code, the output often no longer lines up with the state it was meant to support.
In Cursor, those agents debug issues, generate features, and execute refactors while the developer continues working. That workload requires predictable worst-case latency under concurrency, consistent context handling across overlapping requests, and stable operation under sustained load.
Engineering on NVIDIA Blackwell
Meeting the latency budget at scale led the teams to NVIDIA Blackwell GB200 NVL72 and NVIDIA HGX™ B200 , with higher memory bandwidth and tensor throughput to support faster serving. Deploying production workloads early in the lifecycle meant driving reliability and optimization across the full stack — hardware, firmware, host software, and the serving layer.
Frontier infrastructure with NVIDIA Blackwell: Together AI worked with Cursor to deliver and collaborate on NVIDIA Blackwell's early rollout. For Cursor, early hardware access is a product advantage and Together engineers worked on fast upgrades and replacements to deliver this new frontier infrastructure. These efforts delivered these new chips quickly and reliably for Cursor to use.
Full throughput on ARM hosts: GB200 NVL72 pairs the GPUs with the NVIDIA Grace™ CPU on the ARM instruction set. Much of the high-performance inference ecosystem assumes x86 hosts. Porting the inference stack to ARM required kernel and host-level tuning for GB200 NVL72.
Custom kernels for Blackwell Tensor Core: Blackwell introduces new Tensor Core instructions optimized for lower-precision formats. Together AI built kernels for Blackwell, targeting these instructions directly, capturing more of the hardware’s throughput.‍ ‍
Efficient parallelism across NVIDIA GB200 NVL72: GB200 NVL72 connects 72 NVIDIA Blackwell GPUs in an all-to-all topology. Distributing a model across that domain adds communication and synchronization overhead between chips. Together designed parallelism meshes for GB200 NVL72, so coordination costs stayed bounded — and the compute gains carried through to inference.
Shortening the weights-to-production cycle
Cursor’s research team trains models internally — combining proprietary data with targeted optimization for coding workflows — and produces new candidate weights. The collaboration with Together AI established a repeatable path to move these weights to a production-like endpoint for immediate testing.
A key step in that path is quantization. Real-time serving runs on tight memory and compute budgets, and quantization reduces both by representing weights with fewer bits. In coding contexts, a quality drop can surface as subtle logic mistakes or syntax errors, so quantization has to preserve output quality while improving latency and cost.
Together AI implemented a quantization pipeline built around NVIDIA TensorRT LLM and NVFP4 on Blackwell — threading the needle between aggressive compression and the quality bar Cursor's coding models require. When Cursor produces a new candidate, Together quantizes it, validates it, and spins up a test endpoint within days. Cursor runs internal evaluation suites, then stages A/B testing under production traffic before completing a cutover. Cutovers are gated by validation and live-traffic checks.
What’s next — from latency to throughput
Cursor’s production deployment runs on NVIDIA Blackwell GPUs across multiple data centers. NVIDIA GB200 NVL72 handles inference, and the infrastructure supports model iteration as Cursor’s research team ships new weights.
With that latency path running in production, the focus shifts to throughput and utilization. Cursor and Together AI are building higher-throughput endpoints on the NVIDIA Blackwell platform to improve per-GPU economics as usage grows.
Together AI’s Instant Clusters Enable Latent Health to Build Clinical AI That Outperforms GPT-4
How HeroUI Chat launched 10x faster with Together Code Sandbox
How Scaled Cognition Trains APT-1 on Together AI GPU Clusters
From AWS to Together Dedicated Endpoints: Arcee AI's journey to greater inference flexibility
Scaling AI Companions: How Dippy AI Reached Over 4 Million Tokens/Minute with Together Dedicated Endpoints
Building World-Class Thai Language Models with Purpose-Built AI Infrastructure
From optimized training and model shaping to large-scale production inference
© 2026 Together AI. All Rights Reserved.

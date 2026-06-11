---
source: "https://uccl-project.github.io/posts/commbench/"
hn_url: "https://news.ycombinator.com/item?id=48487070"
title: "CommBench: Can LLMs Write Correct and Efficient GPU Communication Code?"
article_title: "CommBench: Can LLMs Write Correct and Efficient GPU Communication Code?"
author: "matt_d"
captured_at: "2026-06-11T07:49:24Z"
capture_tool: "hn-digest"
hn_id: 48487070
score: 1
comments: 0
posted_at: "2026-06-11T06:47:11Z"
tags:
  - hacker-news
  - translated
---

# CommBench: Can LLMs Write Correct and Efficient GPU Communication Code?

- HN: [48487070](https://news.ycombinator.com/item?id=48487070)
- Source: [uccl-project.github.io](https://uccl-project.github.io/posts/commbench/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T06:47:11Z

## Translation

タイトル: CommBench: LLM は正しく効率的な GPU 通信コードを記述できますか?
説明: CommBench は、多様な通信機能とコンピューティングと通信の融合カーネルをカバーする、フロンティア LLM がマルチデバイス GPU 通信コードをどのように効率的に生成するかを評価します。

記事本文:
CommBench: LLM は正しく効率的な GPU 通信コードを記述できますか?
私たちについて ブログ投稿 タグで並べ替え CommBench: LLM は正しく効率的な GPU 通信コードを記述できますか?
著者: Shuang Ma、Yuyi Li、Yihan Zhang、Danyang Chen、Shuyang Ji、Ziming Mao、Cheng Ji、Ansha Prashanth、Wenting Yang、Yiran Wang、Chihan Cui、Peiyu Lin、Amanda Raybuck、Ion Stoica、Yang Zhou。
日付: 2026 年 6 月 9 日
GPU 通信は大規模な LLM トレーニングと推論の重要なコンポーネントですが、その複雑さによりコード生成モデルにとっては困難となっています。 CommBench は、100 を超える GPU 通信の問題を含むベンチマークと、UCCL の開発経験に基づいた業界レベルのマルチデバイス通信のユースケースをカバーするリファレンス ソリューション (総称してサンプルと呼ばれます) を紹介します。 CommBench は、ポイントツーポイント、コレクティブ、エキスパート並列、コンピューティングと通信の融合、およびユーティリティ機能に及びます。
これらの例は、GPU 通信の専門家によって手書きされたか、Mscclpp、NCCL、NVSHMEM、DeepEP、ThunderKittens、vLLM、SGLang などの実稼働コードベースから抽出されたものです。次に、ノード内 NVLink とノード間 RDMA にわたる実際のハードウェア上のチート耐性ハーネスの下で主要なクローズド モデルとオープン モデルを評価し、それらが成功または失敗する場所と理由のケース スタディを示します。今後の作業として、このギャップを埋めるために、これらのデータセットで LLM をポストトレーニングする予定です。
CommBench オープンソース: uccl-project/CommBench (MIT ライセンス)。
GPU 通信コードの作成が重要である理由、そしてそれが LLM にとって依然として課題である理由
通信とコンピューティングと通信の融合は、最新の LLM トレーニングと推論を拡張するために不可欠です。実稼働トレーニングでは、コミュニケーションによってフォワード パスの 43.6% が消費される可能性があります 1 。幅広い専門家による並列処理を使用した MoE 推論では、デバイス間通信が考慮されます。

合計実行時間の最大 47% 2 。このコードを正しく高速に実行することは、あればいいというわけではありません。
カスタマイズされた GPU 通信とコンピューティングと通信の融合に対する需要が急速に高まっています。 NCCL のような確立されたライブラリは包括的なインターフェイスを提供しますが、最先端のパフォーマンスよりも汎用性を重視して最適化されています。その結果、企業は多くの場合、より厳密な制御と最適化のために社内の GPU 通信と計算スタックを維持します。 GPU 通信も依然として急速に進化している分野です。新しいハードウェアと新しい LLM アーキテクチャにより、より高いパフォーマンスと特殊なワークロードに対する新しい要件が継続的に導入されている一方で、通信の抽象化は依然として進化しています。
最新の GPU は非常に強力で高価であるため、高度にカスタマイズされたカーネルと、より緊密なコンピューティングと通信の融合を促進して、Hopper、Blackwell、AMD GPU などのアーキテクチャ全体でハードウェアの使用率を最大化します。 GPU の高速化に伴い、NCCL などの従来のライブラリで使用されている CPU 仲介の実行パスに依存するのではなく、通信を GPU から直接開始する必要性がますます高まっています。
MoE エキスパート並列処理などの新しい LLM アーキテクチャでは、既存の集合ライブラリでは十分にサポートされていない、ますます不規則で粒度の細かい通信パターンが導入されています。
マルチデバイス GPU プログラミングは、次の 3 つの理由により、単一デバイスのコーディングよりも本質的に困難です。
専門知識が必要であり、GPU カーネルとネットワークの両方に関する深い知識が必要です。
障害が発生しやすい相互接続上で多くのデバイスを調整する必要がありますが、これは本質的に困難です。
GPU 通信用の実用的で忠実なデータセットがほとんど欠落しているため、データが不足しています。
これらすべてにもかかわらず、マルチデバイス GPU プログラミングは、LLM コーディング ベンチマークではほとんど無視されてきました。 HumanEval、MBPP、LiveCodeBench — これら

単一デバイスの推論を測定します。既存のベンチマークでは、通信プリミティブ (Mscclpp チャネルや集合インターフェイスなど) や計算通信融合カーネル (NVLink と InfiniBand にわたる AllGather+GEMM の融合など) を含む正しい GPU 通信コードをモデルが生成できるかどうかを評価していません。
ベンチマークとフレームワークの構造
データセットは独立して実行可能なサンプルのリストとして編成されており、現在そのようなサンプルが 100 以上あります。完全ですぐに使用できる機能を実装しているものもあります (例: P2P/コレクティブ インターフェイス、MoE エキスパートの並列ディスパッチと結合など)。その他は再利用可能な通信ビルディング ブロック (Mscclpp チャネルなど) です。 UCCL プロジェクトでの実践的な経験を活用して、各サンプルに 3 つの難易度レベル (Easy/Medium/Hard) のいずれかを手動で割り当てます。
いくつかの例は、基本ライブラリの上に手書きで作成しました。その他は、運用グレードの通信および LLM 提供フレームワークから抽出されます。機能により、次のように分類されます。
P2P — デバイスのペア間のポイントツーポイント転送。
集合的 — すべてのランクにわたるグループ操作 (AllReduce、AllGather、All-to-All など)。
EP — MoE モデルの動的で不均一なディスパッチ/結合トラフィック。
Fusion — 通信とコンピューティングをインターリーブするカーネル (AllGather+GEMM など)。
ユーティリティ — 接続セットアップ、バッファ登録、トポロジ クエリなどのサポート コンポーネント。
ソースごとに、それらは cuda-runtime、libibverbs、Mscclpp、nccl、nvshmem、deepep、nccl-device-api、thunderkittens、vllm、および sglang に及びます。
データセット上のさまざまなモデルを自動的に評価し、生成されたコードの複数ラウンドの改良をサポートするフレームワークを構築しました。各サンプルとその実行環境を慎重に設計することで、厳密なチェックが実施され、モデルの不正行為が防止されます。
構造例。それぞれ

この例には 3 つの部分があります。
リファレンス ソリューション (リポジトリから除外) — オブジェクト指向スタイル (Python / CUDA / HIP / C++) で編成された高品質の手書きコード。ランダム化された入力を使用するテスト ハーネスが付属しているため、モデルは予想される出力をハードコーディングできません。参照ソリューションがモデルのトレーニング データに漏洩して将来の評価が汚染されるのを防ぐために、参照ソリューションは秘密に保たれます。
問題ステートメント + 空の解決策 — コア実装が取り除かれ // TODO とマークされているリファレンスですが、関数インターフェイスは保持されています。このファイルはプロンプトの一部になります。モデルはインターフェイス セマンティクスを尊重する必要があり、これにより出力を直接テストできるようになります。不正行為を防止するために、編集する予定の領域のみが変更されたことを確認します。
ビルドして実行するスクリプト — サンプル (参照または生成) を独立してコンパイルして実行し、モデルからは隠されます。コンパイル コマンドを制御することで、生成されたコードが意図したライブラリのみを使用することを厳密に保証します。
マルチラウンドの改良。 max-round > 1 の場合、生成されたコードが実行に失敗するか、リファレンスのパフォーマンスを下回る場合は、コンパイル/実行の出力をプロンプトに戻し、モデルに反復を依頼します。パフォーマンスが向上するかラウンド制限に達するまで繰り返します。
テストベッド。 NVIDIA と AMD GPU の両方で NVLink、InfiniBand、RoCE にまたがる 3 つのクラスターで評価します。(1) NVLink 上の 8 倍の NVIDIA B300 の単一ノード。 (2) ノード間 RDMA 用の 400 Gb/s InfiniBand (ConnectX-7 / BlueField-3) によってリンクされた NVIDIA GH200 ノード。 (3) ノード内 Infinity Fabric (XGMI) および 400 Gb/s RoCE ノード間 RDMA を備えた 8x AMD Instinct MI325X ノード。これにより、CommBench は、ノード内の NVLink だけでなく、両方のベンダーのスタック (CUDA および ROCm/HI) にわたって実際のノード間 RDMA パスを実行できるようになります。

P）。
Pass×GM で並べ替え ⭐ — 合格率は、合格例の幾何平均コード品質によって評価されます。
バー▓░は、絶対 0 ～ 100% のスケールで各値を示します (Pass×GM および GM‑Speedup を 1.0 の分数として表します)。
カラー: 🟢 上段 · 🟡 中段 · 🔴 下段。
パフォーマンス判定しきい値 (4 層):
良好 ≥ +20% · on_compare −5% から +20% · 劣化 −40% から −5% · Serious_degraded < −40%
推論の労力とモデルのパラメーター。すべてのモデルを、デフォルトの推論努力とデフォルトのサンプリング パラメータ (温度、top-p など) で実行します。私たちはそれらを調整しません。モデルごとのデフォルトは次のとおりです: gpt-5.5 = middle 、gemini-3.1-pro-preview = high 、claude-opus-4-7 = 適応的思考、glm-5.1 = 思考が有効 (モデルがいつ思考するかを決定します)、kimi-k2.6 = 思考が有効 (温度は 1.0 に固定)、deepseek-v4-pro = middle 。唯一の例外は qwen3.7-max で、タイムアウトを避けるために思考が無効になっています。
分析: 上位モデルと下位モデル
CommBench で最もスコアの高いモデル、gpt-5.5 (Pass×GM = 0.467) および deepseek-v4-pro (Pass×GM = 0.197) を選択し、難易度、タスク タイプ、ライブラリ カバレッジ、コード パフォーマンスの詳細な内訳を示します。
人間が定義した難易度は、モデルの難易度と必ずしも一致するとは限りません。 deepseek の Easy パフォーマンスの低さ (14%) は、ライブラリ固有の API 知識のギャップが主な原因です。mscclpp や ThunderKittens などのライブラリの場合、API を正しく呼び出すために必要なインターフェイス セマンティクスが欠如しており、単純な機能を実装するために少数の関数を呼び出すだけで済むサンプルでエラーが発生します。ハードエンドでは、多くの例でモデルに通信プリミティブをほぼスクラッチから実装するよう要求しています (NVLink または RDMA 経由の AllReduce など)。これにより、ディープシークの機能が制限されていることがわかります。

複雑な GPU 通信計算タスク。
PASS 例におけるパフォーマンスの品質
gpt-5.5 は品質のばらつきが最も広く、8 件のひどく劣化したケース (PASS の 14%) があり、すべてハード サンプルに集中しています。これは、gpt-5.5 がコンパイルと実行に成功したサンプル (ただしパフォーマンスは低下します) が、ディープシークではコンパイルまたは実行に失敗するため、ディープシークのパフォーマンス分布にまったく表示されないことが部分的に原因です。
gpt-5.5 は Collective で優勢 (66% 対 22%) であり、特殊なライブラリ (mscclpp (17%)、nccl-device-api (40%)、thunderkittens (54%)) を有意義にカバーしている唯一のモデルです。 deepseek は、NCCL (100%) および P2P タグで競合しますが、3 つの特殊なライブラリすべてでスコアは 0% です。 vllm、NCCL、nvshmem などの広く採用されているライブラリでは、どちらのモデルも良好なパフォーマンスを発揮します。
予算の制約により、gpt-5.5 は単一世代ラウンドで評価されました。 Deepseek の最初のラウンドのパフォーマンスは低いにもかかわらず、その API コストは gpt-5.5 のわずか 1% 程度 (1 例あたり 0.02 ドル対 1.91 ドル) であり、複数ラウンドの自己修正が経済的に実行可能です。最大 5 回の自己修正ラウンドまでのディープシークを許可すると、プロファイルが大幅に変わります。
マルチラウンド自己修正により、ディープシークの全体的な合格率が 2 倍以上 (16→42、15.8%→41.6%)、中程度の難易度の商品ライブラリ タスクのロックが大幅に解除されます (31%→60%)。ハード サンプルや特殊なライブラリ (mscclpp: 0%、thunderkittens: 4%) のロックは解除されません。これらには、モデルが持っていない分野の知識が必要です。実際的な意味: タスクがコモディティ ライブラリ (NCCL、vllm、cuda-runtime、nvshmem) に限定されており、再試行のバジェットが許容される場合、再試行を伴うディープシークは合理的な選択です。
ケース 1: 部分パス — ThunderKittens AllToAll
レベル 🟢 簡単 · タグ集合 · 図書館サンダーキトゥン
タスク

: NVLink を介した GPU シャード間のタイルレベルのデータ移動に TMA (Tensor Memory Accelerator) を使用し、ThunderKittens ライブラリを使用して BF16 マルチ GPU AllToAll カーネルを実装します。
実装する内容: 2 つの小さなデバイス側カーネル: all_to_all::kernel (ローカル入力から 1 つのタイルを TMA ロード、宛先デバイスの出力シャードに TMA ストア) と all_to_all_barrier::kernel (1 行の Barrier_all 同期)。すべてのホスト足場が提供されます。
ほとんどのモデルが失敗した理由: どちらのカーネルも 10 行未満で、既存の ThunderKittens インターフェイスを呼び出すだけでよく、カスタム アルゴリズムは必要ありません。失敗の原因は、ThunderKittens 固有の API の知識です。タイプのスコープ ( all_to_all::globals::shared_tile と裸のshared_tile)、正しい tma::load_async 引数の署名、およびshared_allocator::allocateテンプレートの使用法です。 ThunderKittens は、トレーニング データの範囲が最小限に抑えられたニッチな研究ライブラリであり、モデルのスコープ タイプの誤り、存在しない API の幻覚、または間違った引数形式の使用を引き起こします。
ケース 2: すべての LLM が失敗したケース
レベル 🟡 中 · タグ ユーティリティ · ライブラリ cuda-runtime
タスク: ノンブロッキング try_wait プローブを備えた Hopper/Blackwell 共有メモリ mbarrier を使用して、CTA 内のプロデューサー/コンシューマー同期を実装します。これは、データ到着をオーバーラップするために永続カーネル タイル パイプライン (CUTLASS、Mirage など) で使用されるパターンです。

[切り捨てられた]

## Original Extract

CommBench evaluates how effectively frontier LLMs generate multi-device GPU communication code, covering diverse communication functionalities and compute–communication fusion kernels.

CommBench: Can LLMs Write Correct and Efficient GPU Communication Code?
About Us Blog Posts Sort by Tags CommBench: Can LLMs Write Correct and Efficient GPU Communication Code?
By: Shuang Ma, Yuyi Li, Yihan Zhang, Danyang Chen, Shuyang Ji, Ziming Mao, Cheng Ji, Ansha Prashanth, Wenting Yang, Yiran Wang, Chihan Cui, Peiyu Lin, Amanda Raybuck, Ion Stoica, Yang Zhou.
Date: June 9, 2026
GPU communication is a critical component of large-scale LLM training and inference, yet its complexity makes it challenging for code-generation models. We present CommBench, a benchmark with 100+ GPU communication problems + reference solutions (collectively called examples) that cover industry-level multi-device communication use cases based on UCCL's development experience. CommBench spans point-to-point , collective , expert-parallel , compute and communication fusion , and utility functions .
These examples are either hand-written by GPU communication experts or distilled from production codebases such as Mscclpp, NCCL, NVSHMEM, DeepEP, ThunderKittens, vLLM, and SGLang. We then evaluate leading closed and open models under a cheat-resistant harness on real hardware spanning intra-node NVLink and inter-node RDMA, and present case studies of where and why they succeed or break down. As future work, we plan to post-train LLMs on these datasets to close this gap.
CommBench open-source: uccl-project/CommBench (MIT license).
Why Writing GPU Communication Code Matters—and Why It Remains Challenging for LLMs?
Communication and compute-communication fusion are essential for scaling modern LLM training and inference. In production training, communication can consume 43.6% of the forward pass 1 ; in MoE inference with wide expert parallelism, inter-device communication accounts for up to 47% of total execution time 2 . Getting this code right and fast is not a nice-to-have.
The demand for customized GPU communication and compute-communication fusion is rapidly growing. Established libraries like NCCL offer comprehensive interfaces, but optimize for generality over frontier performance. As a result, companies often maintain in-house GPU communication and computation stacks for tighter control and optimization. GPU communication also remains a rapidly evolving area: new hardware and new LLM architectures continuously introduce new requirements for higher performance and specialized workloads, while communication abstractions are still evolving:
Modern GPUs are extremely powerful and expensive , motivating highly customized kernels and tighter compute–communication fusion to maximize hardware utilization across architectures such as Hopper, Blackwell, and AMD GPUs. As GPUs become faster, communication increasingly needs to be initiated directly from GPUs instead of relying on CPU-mediated execution paths used in traditional libraries such as NCCL.
New LLM architectures, such as MoE expert parallelism , introduce increasingly irregular and fine-grained communication patterns that are not well supported by existing collective libraries.
Multi-device GPU programming is inherently harder than single-device coding, for three reasons:
It demands niche expertise , requiring deep knowledge of both GPU kernels and networking.
It requires coordinating many devices over fail-prone interconnects , which is intrinsically difficult.
It lacks data , as practical, faithful datasets for GPU communication are largely missing.
Despite all this, multi-device GPU programming has been largely overlooked in LLM coding benchmarks. HumanEval, MBPP, LiveCodeBench — these measure single-device reasoning. No existing benchmark evaluates whether a model can generate correct GPU communication code, including communication primitives (e.g., Mscclpp channels and collective interfaces) and compute–communication fusion kernels (e.g., fused AllGather+GEMM across NVLink and InfiniBand).
Benchmark and Framework Structure
The dataset is organized as a list of independently runnable examples, currently with 100+ such examples. Some implement complete, ready-to-use functionality (e.g., P2P/collective interfaces, or MoE expert-parallel dispatch and combine); others are reusable communication building blocks (e.g., Mscclpp channels). Drawing on our hands-on experience from the UCCL project, we manually assign each example one of three difficulty levels: Easy / Medium / Hard .
Some examples we hand-wrote on top of base libraries; others are extracted from production-grade communication and LLM-serving frameworks. By function , they fall into:
P2P — point-to-point transfer between a pair of devices.
Collective — group operations across all ranks (AllReduce, AllGather, All-to-All, …).
EP — dynamic, non-uniform dispatch/combine traffic for MoE models.
Fusion — kernels that interleave communication with compute (e.g., AllGather+GEMM).
Utilities — supporting components such as connection setup, buffer registration, and topology queries.
By source , they span cuda-runtime, libibverbs, Mscclpp, nccl, nvshmem, deepep, nccl-device-api, thunderkittens, vllm, and sglang.
We built a framework that automatically evaluates different models on the dataset and supports multi-round refinement of generated code. Careful design of each example and its execution environment enforces strict checking and prevents the model from cheating.
Example structure. Each example has three parts:
Reference solution (excluded from the repo) — high-quality, hand-written code organized in an object-oriented style (Python / CUDA / HIP / C++). It ships with a test harness that uses randomized inputs, so a model cannot hard-code expected outputs. We keep the reference solutions secret to prevent them from leaking into model training data and contaminating future evaluations.
Problem statement + empty solution — the reference with its core implementation stripped out and marked // TODO , but with the functional interface preserved. This file becomes part of the prompt: the model must honor the interface semantics, which also keeps its output directly testable. We verify that only the regions meant to be edited were changed, guarding against cheating.
Build-and-run script — independently compiles and runs an example (reference or generated) and is hidden from the model. By controlling the compile command, we strictly ensure the generated code uses only the intended libraries.
Multi-round refinement. When max-round > 1 , if the generated code fails to run or underperforms the reference, we feed the compile/run output back into the prompt and ask the model to iterate — repeating until performance improves or the round limit is reached.
Testbed. We evaluate on three clusters that together span NVLink, InfiniBand, and RoCE on both NVIDIA and AMD GPUs: (1) a single node of 8× NVIDIA B300 over NVLink; (2) NVIDIA GH200 nodes linked by 400 Gb/s InfiniBand (ConnectX-7 / BlueField-3) for inter-node RDMA; and (3) 8× AMD Instinct MI325X nodes with intra-node Infinity Fabric (XGMI) and 400 Gb/s RoCE inter-node RDMA. This lets CommBench exercise real inter-node RDMA paths, not only intra-node NVLink, across both vendors’ stacks (CUDA and ROCm/HIP).
Sorted by Pass×GM ⭐ — pass rate scaled by geometric-mean code quality on passing examples.
Bars ▓░ show each value on an absolute 0–100% scale (Pass×GM and GM‑Speedup as a fraction of 1.0).
Color: 🟢 top tier · 🟡 mid tier · 🔴 bottom tier.
Performance verdict thresholds (4-tier):
better ≥ +20% · on_compare −5% to +20% · degraded −40% to −5% · severely_degraded < −40%
Reasoning effort and model parameters. We run every model at its default reasoning effort and default sampling parameters (temperature, top-p, etc.); we do not tune them. The per-model defaults are: gpt-5.5 = medium , gemini-3.1-pro-preview = high , claude-opus-4-7 = adaptive thinking, glm-5.1 = thinking enabled (the model decides when to think), kimi-k2.6 = thinking enabled (temperature fixed at 1.0), deepseek-v4-pro = medium . The only exception is qwen3.7-max, where thinking is disabled to avoid timeouts.
Analysis: Top vs. Bottom Model
We select the highest- and lowest-scoring models on CommBench, gpt-5.5 (Pass×GM = 0.467) and deepseek-v4-pro (Pass×GM = 0.197), for a detailed breakdown across difficulty levels, task types, library coverage, and code performance.
Human-defined difficulty does not always align with model difficulty. deepseek’s weak Easy performance (14%) is largely attributable to gaps in library-specific API knowledge — for libraries such as mscclpp and ThunderKittens, it lacks the interface semantics needed to call APIs correctly, causing failures on examples that only require invoking a small number of functions to implement straightforward functionality. On the Hard end, many examples ask the model to implement communication primitives from near-scratch (e.g. AllReduce over NVLink or RDMA) which exposes deepseek’s limited capability on complex GPU communication-compute tasks.
Performance Quality Among PASS Examples
gpt-5.5 has the widest quality spread: 8 severely degraded cases (14% of PASS), all concentrated in Hard examples. This is partly because examples that gpt-5.5 manages to compile and run (but with degraded performance) would simply fail to compile or execute under deepseek, and therefore never appear in deepseek’s performance distribution at all.
gpt-5.5 dominates on Collective (66% vs 22%) and is the only model with meaningful coverage of specialized libraries — mscclpp (17%), nccl-device-api (40%), and thunderkittens (54%). deepseek is competitive on NCCL (100%) and P2P tags, but scores 0% across all three specialized libraries. For widely adopted libraries such as vllm, NCCL, and nvshmem, both models perform well.
Due to budget constraints, gpt-5.5 was evaluated with a single generation round. Despite deepseek’s weaker first-round performance, its API cost is only ~1% of gpt-5.5’s ($0.02 vs $1.91 per example), which makes multi-round self-correction economically viable. Allowing deepseek up to 5 self-correction rounds substantially changes its profile.
Multi-round self-correction more than doubles deepseek’s overall pass rate (16→42, 15.8%→41.6%) and substantially unlocks Medium-difficulty commodity-library tasks (31% → 60%). It does not unlock Hard examples or specialized libraries (mscclpp: 0%, thunderkittens: 4%) — those require domain knowledge the model does not have. The practical implication: deepseek with retries is a reasonable choice when tasks are restricted to commodity libraries (NCCL, vllm, cuda-runtime, nvshmem) and a retry budget is acceptable.
Case 1: Partial Pass — ThunderKittens AllToAll
Level 🟢 Easy · Tag Collective · Library thunderkittens
Task: Implement a BF16 multi-GPU AllToAll kernel using the ThunderKittens library, using TMA (Tensor Memory Accelerator) for tile-level data movement between GPU shards over NVLink.
What to implement: Two small device-side kernels: all_to_all::kernel (TMA-load one tile from local input, TMA-store to destination device’s output shard) and all_to_all_barrier::kernel (one-line barrier_all synchronization). All host scaffolding is provided.
Why most models failed: Both kernels are under 10 lines and only require calling existing ThunderKittens interfaces, with no custom algorithm needed. The failure root is ThunderKittens-specific API knowledge: type scoping ( all_to_all::globals::shared_tile vs bare shared_tile ), the correct tma::load_async argument signature, and the shared_allocator::allocate template usage. ThunderKittens is a niche research library with minimal training-data coverage, causing models to mis-scope types, hallucinate non-existent APIs, or use wrong argument forms.
Case 2: A Case All LLMs Failed
Level 🟡 Medium · Tag Utilities · Library cuda-runtime
Task: Implement intra-CTA producer/consumer synchronization using a Hopper/Blackwell shared-memory mbarrier with a non-blocking try_wait probe, a pattern used in persistent-kernel tile pipelines (e.g., CUTLASS, Mirage ) to overlap data arriv

[truncated]

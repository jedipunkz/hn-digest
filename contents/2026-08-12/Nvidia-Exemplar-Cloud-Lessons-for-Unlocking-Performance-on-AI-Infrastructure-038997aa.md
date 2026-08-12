---
source: "https://developer.nvidia.com/blog/nvidia-exemplar-cloud-lessons-for-unlocking-full-performance-on-ai-infrastructure/"
hn_url: "https://news.ycombinator.com/item?id=49271985"
title: "Nvidia Exemplar Cloud: Lessons for Unlocking Performance on AI Infrastructure"
article_title: "NVIDIA Exemplar Cloud: Lessons for Unlocking Full Performance on AI Infrastructure | NVIDIA Technical Blog"
author: "gmays"
captured_at: "2026-08-12T14:15:13Z"
capture_tool: "hn-digest"
hn_id: 49271985
score: 3
comments: 0
posted_at: "2026-08-12T13:19:35Z"
tags:
  - hacker-news
  - translated
---

# Nvidia Exemplar Cloud: Lessons for Unlocking Performance on AI Infrastructure

- HN: [49271985](https://news.ycombinator.com/item?id=49271985)
- Source: [developer.nvidia.com](https://developer.nvidia.com/blog/nvidia-exemplar-cloud-lessons-for-unlocking-full-performance-on-ai-infrastructure/)
- Score: 3
- Comments: 0
- Posted: 2026-08-12T13:19:35Z

## Translation

タイトル: Nvidia Exemplar Cloud: AI インフラストラクチャのパフォーマンスを解放するためのレッスン
記事のタイトル: NVIDIA Exemplar Cloud: AI インフラストラクチャのパフォーマンスを最大限に引き出すためのレッスン | NVIDIA テクニカル ブログ
説明: 同一の NVIDIA H100、GB200 NVL72、または GB300 NVL72 システムから構築された 2 つの AI コンピューティング クラスターは、大幅に異なるトレーニング スループットを提供できます。

記事本文:
NVIDIA Exemplar Cloud: AI インフラストラクチャのパフォーマンスを最大限に引き出すためのレッスン | NVIDIA テクニカル ブログ
開発者
ホーム
購読する
関連リソース
データセンター/クラウド
英語 한국어 中文
NVIDIA Exemplar Cloud: AI インフラストラクチャのパフォーマンスを最大限に引き出すためのレッスン
いいね
嫌い
同一の NVIDIA H100、GB200 NVL72、または GB300 NVL72 システムで構築されたクラスター間でのトレーニング スループットの重大な違いは、主にカーネル、ハイパーバイザー、BIOS、および NVIDIA Collective Communications Library (NCCL) レベルでの構成ギャップが複合的に発生することに起因し、導入環境が NVIDIA Exemplar Cloud 検証の 95% しきい値を満たさないことがよくあります。
4 つの実際のケーススタディは、パフォーマンス損失の繰り返しの原因を明らかにしています。それは、SMMU 機能の欠落と NVIDIA Grace CPU での不適切な仮想化構成です。 CPU の C ステートと NUMA の設定ミスにより、ターボ周波数とメモリの局所性が最適化されません。 ConnectX-8 SuperNIC などの高帯域幅ファブリックでの NCCL キュー ペアの同時実行性が不十分。また、NCCL トポロジ ファイルをコンテナに伝播できないため、サイレントかつ深刻な AllGather/ReduceScatter の速度低下が発生します。
インフラストラクチャ エンジニアは、SMMU と VM カーネルの機能を系統的に検証し、CPU 電力管理と NUMA/プロセス バインディングが最適化されていることを確認し、ファブリックのスケールとワークロードに合わせて NCCL キュー ペアの同時実行性を調整し、目的のコンテナ化されたトレーニング環境内ですべてのトポロジ/環境変数にアクセスできることを保証することで、パフォーマンスのギャップを埋めることができます。
AI によって生成されたコンテンツでは、情報が不完全に要約されている可能性があります。重要な情報を確認してください。さらに詳しく
同一の NVIDIA H100、GB200 NVL72、または GB300 NVL72 システムから構築された 2 つの AI コンピューティング クラスターは、大幅に異なるトレーニング スループットを提供できます。私たちは

同じワークロード、同じモデル、同じグローバル バッチ サイズ上で、パートナーの導入と対応する NVIDIA リファレンス アーキテクチャ (RA) との間には 8% ～ 12% のギャップが日常的に見られます。
原因は多くの場合、カーネル、ハイパーバイザー、BIOS、および NVIDIA Collective Communications Library (NCCL) 設定の構成選択の積み重ねであり、それぞれに数パーセントのコストがかかり、NVIDIA Exemplar Cloud の検証に必要な 95% のしきい値を満たさないほど大きなギャップとなります。
この投稿では、実際のパートナー クラスターからの 4 つのデバッグ調査について説明します。各診断は、スタックの個別の層、つまりシステム メモリ管理ユニット (SMMU) と NVIDIA Grace CPU でのページ テーブルの動作を分離します。 x86 ベースの CPU での電源管理と不均一メモリ アクセス (NUMA) の配置。 1.6 Tbps ファブリック上の NVIDIA NCCL キュー ペアの同時実行性。そして、サイレントハードウェアインストールの欠陥。この投稿では、根本原因を指摘した perf 、 NVIDIA Nsight Systems 、または NVIDIA NCCL テストの特定の信号と、ギャップを埋めるためのチューニング変更も示しています。
すでにこれらのベンチマークを実行しているインフラストラクチャ エンジニアやパフォーマンス アーキテクトは、正式な RA 検証の前に独自のクラスターに対して実行するために内部で使用している診断パターンから恩恵を受けることができます。
この投稿の診断を再現するには、次のものが必要です。
NVIDIA Quantum InfiniBand または RoCE インターコネクトを備えた NVIDIA HGX H100、HGX H200、HGX B200、GB200 NVL72、または GB300 NVL72 システム クラスター。
安定した反復タイミングを備えた分散トレーニング ワークロード - Llama 3 モデル上の NVIDIA NeMo、NVIDIA Nemotron 、または DeepSeek 構成は、合理的な参照です。
perf 、BIOS/UEFI の変更、およびカーネル パラメータの変更のための、少なくとも 1 つのノードでの root アクセス。
nccl-tests は、トレーニング スタックが使用するものと同じ NCCL バージョン、NVIDIA Nsight Systems、および Li に対して構築されています。

nux perf とカーネル シンボルが利用可能。
トレーニングのパフォーマンスギャップの背後にある一般的なパターン
最近の Exemplar トレーニングでは、単一の明らかな失敗によってパフォーマンスのギャップが生じることはほとんどないことがわかりました。多くの場合、それらは、ワークロードの負荷がかかった場合にのみ表示される構成の詳細から発生します。繰り返されるパターンには次のようなものがあります。
猶予と仮想化の準備: プラットフォーム機能の不足、SMMU オーバーヘッド、IOMMU の動作、または予期される構成と一致しないページ サイズ設定。
CPU パワーとプロセスの配置: 予想されるターボ周波数を下回って実行されているコア、間違ったコアに配置されたランクまたはヘルパー スレッド、またはプラットフォーム トポロジと一致しない NUMA/PCT バインディング。
ランタイム トポロジ: ノード上では正しいが、ワークロード コンテナーまたはランチャー環境内では欠落しているホスト トポロジ ファイルまたは NCCL 設定。
ファブリックと集団動作: ターゲット ファブリック、メッセージ サイズ、トレーニング ワークロードの規模と一致しない NCCL 設定。
アプリケーションとプラットフォームのバインディング: トレーニング プロセスは、トポロジを意識したアフィニティではなく、コア ID またはランク順によってバインドされます。
トレーニング パフォーマンスのギャップの原因はこれらだけではありません。また、それらをチェックすることが検証を実際のアプリケーションに置き換えるわけではありません。
以下の 4 つのケーススタディは、最近のトレーニング作業でこれらのパターンがどのように現れたか、つまり、どの信号が問題を明らかにし、何が変更され、修正がどのように検証されたかを示しています。この順序は普遍的な優先順位ではありません。適切な開始点は、ワークロード、プラットフォーム、および最初のプロファイラー信号によって異なります。
ケーススタディ 1: NVIDIA GB200 NVL72 FP8 の事前トレーニング、仮想マシン (VM) ではベアメタルよりも 12% 遅い
レイヤ: 仮想化と SMMU
VM 内で DeepSeek-V3 Mixture-of-Experts (MoE) FP8 事前トレーニングを実行している GB200 NVL72 パートナー デプロイメントでは、反復時間が 12% ～ 14 となっていました。

ベアメタル RA よりも % 長いです。 Llama 3 70B のような高密度モデルの事前トレーニング レシピは RA パフォーマンスの 3% 以内で実行されましたが、反復ごとに多数の小さなカーネルを発行する DeepSeek-V3 MoE は異常値でした。
パートナー クラスターでキャプチャされた Nsight Systems トレースでは、ワークロードの小さなカーネル領域で CPU オーバーヘッドが大幅に高いことが示されました。 CPU シングル スレッドのパフォーマンスのみを対象としたマイクロベンチマークでは、パートナー ノードと RA クラスター ノードでほぼ同一のパフォーマンスが実証されました。これは、ホスト上の 30 秒間の perf Record -a -g キャプチャを perf report で表示すると、予想外のトップ フレーム、つまり CPU サイクルの 24% が arm_smmu_cmdq_issue_cmdlist に費やされたことを示しています。
図 1. 仮想化 NVIDIA GB200 NVL72 システムでの DeepSeek-V3 FP8 事前トレーニング中の Arm SMMU コマンド キュー無効化オーバーヘッドを示す Linux 氷柱グラフ
arm_smmu_cmdq_issue_cmdlist は、Arm SMMU のコマンド キューに無効化コマンドを送信する関数です。仮想化では、ゲストの無効化を引き起こすすべてのマップ/マップ解除がホストをトラップし、単一のコマンド キューを通じてシリアル化され、プロファイルに表示されるスピンロック競合が生成されます。仮想コマンド キュー (VCMDQ) は、標準 Arm SMMUv3 のコマンド キュー仮想化拡張機能を通じて利用できる機能で、ゲストが VM 終了なしで SMMU 無効化コマンドをハードウェアに直接発行できるようになります。
修正: パートナー クラスター上のホスト カーネルで CMDQV/VCMDQ を有効にし、ゲストに公開します。これには、tegra241-cmdqv ドライバーで構築されたカーネルと、対応するハイパーバイザー サポートが必要です。最近の QEMU/libvirt バージョンでは、ゲストに公開するための cmdqv IOMMU 属性が追加されました。
この変更後、Linux perf では arm_smmu_cmdq_issue_cmdlist が上位フレームから外れ、dTLB ミス率がベアメットに戻っていることが示されました。

アルパリティ。 MoE の反復時間ギャップは 12% から RA 許容範囲内に縮小しました。
ここで重要なのは、Grace ベースの仮想化展開では、メモリ マッピングの負荷が高いワークロードに適切な SMMU 機能を公開する VM スタックが必要であるということです。ホスト カーネルで CMDQV/VCMDQ を有効にしてゲストに公開すると、プラットフォームは不要な SMMU シリアル化を回避し、MoE トレーニング パフォーマンスを RA 許容範囲内に戻すことができます。
次の層は CPU 自体であり、そこでは障害モードがまったく異なって見えます。
ケーススタディ 2: CPU 競合と NUMA ミスバインディングにより H100 クラスターが 12% 損失
レイヤー: CPU パワーとプロセスの配置。
NVIDIA の HGX RA と同じ NCCL バージョンと NeMo コンテナを実行しているパートナーの H100 SXM5 クラスターでは、Llama 3 70B の事前トレーニングが基準より 12% 遅く実行されていました。 GB200 NVL72 の場合とは異なり、これはカーネル レベルの問題ではありませんでした。すべてはユーザー空間とBIOSで起こりました。
CPU 周波数: トレーニング中のturbostat -i 1 では、SKU が 3.8 GHz ターボと評価されているにもかかわらず、ビジー コアが 3.0 GHz に固定されていることが示されました。アイドル状態のコアも 3.0 GHz で、C ステートは C6 に落ちずに C1 に留まりました。
NUMA リモート トラフィック: umastat -p <python_pid> は、トレーニング プロセスのメモリ アクセスの約 18% がリモート NUMA ノードに送信されていることを示しました。
パートナー クラスタの CPU は、BIOS で C1 に制限された C ステートで構成されていました。これは一般的な「低レイテンシ」のデフォルトですが、AI トレーニング ワークロードにとっては明らかに間違っています。アイドル状態のコアが C1 に保持されると、パッケージ電力が消費され続けました。 GPU にカーネルを供給する多忙なコアは、ターボを発生させるのに十分なパッケージ電力バジェットを要求できませんでした。アイドル状態のコアを C6 まで下げることを許可すると、電力ヘッドルームが解放され、ビジー状態のコアが 3.8 GHz まで上昇し、このワークロードで約 4% を回復できるようになります。
ハイパーバイザーのハウスキーピング スレッドは固定されていました。

o トレーニング プロセスのデータ ローダー ワーカーと同じ物理コア。 VM 内部では、これは Python スレッドで散発的な 50 ～ 100 ミリ秒のストールのように見え、その後、ステップ時間のロングテールとして伝播しました。修正は CPUset の分離でした。ハイパーバイザーとホスト サービスはコア 0 ～ 7 と 56 ～ 63 にあり、トレーニング プロセスは残りの部分にあります。
結果: 12% のギャップは 3% に縮小し、残りは次のケース スタディで取り上げる別の NCCL チューニングの問題に起因することが判明しました。
ここでのパターンは、単一の修正ではギャップ全体を回復できなかったということです。 C ステートの変化が最大の単一要因であり、最大の要因は約 4% で、残りは NUMA バインディングによるプロセスの分離によるものです。 CPU と仮想化に取り組むと、次の上限はネットワークになります。
ケーススタディ 3: 1.6 Tbps ファブリックを十分に活用していない NVIDIA ConnectX-8 SuperNIC を搭載した GB300 NVL72
焦点: ConnectX-8 SuperNIC の一括チューニング
NVIDIA ConnectX-8 SuperNIC (ノードあたり 1.6 Tbps) を備えた GB300 NVL72 展開では、Nemotron-4 15B 事前トレーニングで 31% のトレーニング パフォーマンス ギャップが示されました。単一ノードのスループットは健全に見えました。ギャップは 512 GPU で発生し、プロファイラーは公開された AllGather および ReduceScatter 時間を示しました。これは、コンピューティングではなく、ConnectX-8 ファブリック上の集合パスを指しました。
調査では、反復回数、UCX/UCC 動作、NUMA マッピング、NVLS、NCCL バージョンなど、NCCL テスト ( nccl-tests ) を使用していくつかの変数をテストしました。ワークロードのネットワーク パフォーマンスに関しては、関連するチューニング変更がより狭くなり、NCCL_IB_QPS_PER_CONNECTION をデフォルト値の 1 から 4 に増やしました。
図 2. 512 GPU スケールでの NCCL QPS 最適化を使用した Nemotron-4 15B のパフォーマンス
Nsight Systems のトレースは、より低い QPS 値で露呈した通信オーバーヘッドを示しており、これがトレーニング反復時間の延長に寄与しています
信号はワークロードと NC の両方に表示されました

cl-test は集合的な測定値をテストします。 NVIDIA リファレンス クラスターでは、デフォルト構成は反復あたり約 1.09 秒で実行されました。 QPS=4 では、同じ参照ワークロードが約 0.83 秒に改善されました。プロファイルでは、AllGather 時間は約 375 ミリ秒から 262 ミリ秒に低下し、ReduceScatter は約 389 ミリ秒から 273 ミリ秒に低下しました。比較の実行時間は約 0.76 秒で、別の NCCL バージョンを使用しました。したがって、残りの差異は、比較環境と参照環境の間の NCCL バージョンの不一致が部分的に原因でした。バージョンを揃えることで、残りのギャップがさらに狭まりました。 NCCL バージョンの変更は通常の Exemplar チューニングの範囲外であるため、推奨されるチューニングでは、デプロイされた NCCL バージョンは変更されません。
教訓 : どこでも QPS を上げてはいけない。 QPS はファブリックとワークロードに依存します。この GB300 ConnectX-8 ワークロードでは、QPS=4 により、大きなメッセージの AllGather および ReduceScatter の動作が改善されました。他のファブリックまたはメッセージ サイズ プロファイルでは、同じ設定でもトレーニング スループットが向上せずに CPU オーバーヘッドが追加される可能性があります。正しいアプローチは、ワークロードの実際のメッセージ サイズでコレクティブをテストし、ターゲット ファブリック上の設定をスイープし、トレーニング ワークロードで結果を検証することです。
ケーススタディ 4: 内部に到達しなかった環境変数
仮想環境で

[切り捨てられた]

## Original Extract

Two AI computing clusters built from identical NVIDIA H100, GB200 NVL72, or GB300 NVL72 systems can deliver materially different training throughput.

NVIDIA Exemplar Cloud: Lessons for Unlocking Full Performance on AI Infrastructure | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Data Center / Cloud
English 한국어 中文
NVIDIA Exemplar Cloud: Lessons for Unlocking Full Performance on AI Infrastructure
Like
Dislike
Material differences in training throughput across clusters built from identical NVIDIA H100, GB200 NVL72, or GB300 NVL72 systems result primarily from compounded configuration gaps at the kernel, hypervisor, BIOS, and NVIDIA Collective Communications Library (NCCL) levels, frequently causing deployments to miss the 95% threshold for NVIDIA Exemplar Cloud validation.
Four real-world case studies highlight recurring sources of performance loss: missing SMMU capabilities and improper virtualization configuration on NVIDIA Grace CPUs; CPU C-state and NUMA misconfiguration leading to sub-optimal turbo frequencies and memory locality; insufficient NCCL queue-pair concurrency on high-bandwidth fabrics such as ConnectX-8 SuperNICs; and failure to propagate NCCL topology files into containers, resulting in silent and severe AllGather/ReduceScatter slowdowns.
Infrastructure engineers can close performance gaps by systematically verifying SMMU and VM kernel capabilities, ensuring CPU power management and NUMA/process bindings are optimized, tuning NCCL queue-pair concurrency to match fabric scale and workload, and guaranteeing all topology/environment variables are accessible inside the intended containerized training environment.
AI-generated content may summarize information incompletely. Verify important information. Learn more
Two AI computing clusters built from identical NVIDIA H100, GB200 NVL72, or GB300 NVL72 systems can deliver materially different training throughput. We routinely see 8% to 12% gaps between partner deployments and the corresponding NVIDIA reference architecture (RA) on the same workload, same model, same global batch size.
The cause is often a stack of configuration choices in the kernel, hypervisor, BIOS, and NVIDIA Collective Communications Library (NCCL) settings, each costing a few percent, that compound into a gap large enough to miss the 95% threshold required for NVIDIA Exemplar Cloud validation.
This post walks through four debugging investigations from real partner clusters. Each diagnostic isolates a distinct layer of the stack: system memory management unit (SMMU) and page-table behavior on NVIDIA Grace CPU ; power management and non-uniform memory access (NUMA) placement on x86-based CPU; NVIDIA NCCL queue-pair concurrency on 1.6 Tbps fabrics; and silent hardware-installation defects. The post also shows the specific signal in perf , NVIDIA Nsight Systems , or NVIDIA NCCL tests that pointed to the root cause, alongside the tuning change that closed the gap.
Infrastructure engineers and performance architects who already run these benchmarks can benefit from these diagnostic patterns we use internally to run against their own clusters before formal RA validation.
To reproduce the diagnostics in this post, you will need:
An NVIDIA HGX H100, HGX H200, HGX B200, GB200 NVL72, or GB300 NVL72 systems cluster with NVIDIA Quantum InfiniBand or RoCE interconnect.
A distributed training workload with stable iteration timing— NVIDIA NeMo on Llama 3 model, NVIDIA Nemotron , or DeepSeek configuration is a reasonable reference.
Root access on at least one node for perf , BIOS/UEFI changes, and kernel parameter changes.
nccl-tests built against the same NCCL version your training stack uses, NVIDIA Nsight Systems, and Linux perf with kernel symbols available.
Common patterns behind training performance gaps
Recent Exemplar training engagements show that performance gaps rarely come from a single obvious failure. More often, they come from configuration details that become visible only under workload pressure. Some recurring patterns include:
Grace and virtualization readiness: Missing platform capabilities, SMMU overhead, IOMMU behavior, or page-size settings that don’t match the expected configuration.
CPU power and process placement: Cores running below expected turbo frequency, ranks or helper threads placed on the wrong cores, or NUMA/PCT bindings that don’t match the platform topology.
Runtime topology: Host topology files or NCCL settings that are correct on the node but missing inside the workload container or launcher environment.
Fabric and collective behavior: NCCL settings that don’t match the target fabric, message size, or scale of the training workload.
Application-to-platform binding: Training processes binding by core ID or rank order instead of topology-aware affinity.
These aren’t the only causes of training performance gaps, and checking them doesn’t replace validation with real applications.
Four case studies below show how these patterns appeared in recent training work: what signal exposed the issue, what changed, and how the fix was verified. The order isn’t a universal triage sequence; the right starting point depends on the workload, platform, and first profiler signal.
Case study 1: NVIDIA GB200 NVL72 FP8 pre-training, 12% slower in a virtual machine (VM) than on bare metal
Layer: Virtualization and SMMU
A GB200 NVL72 partner deployment running DeepSeek-V3 Mixture-of-Experts (MoE) FP8 pre-training inside a VM was producing iteration times 12% to 14% longer than the bare-metal RA. Pre-training recipes for dense models like Llama 3 70B ran within 3% of RA performance, while DeepSeek-V3 MoE, which issues many small kernels per iteration, was the outlier.
Nsight Systems traces captured on the partner cluster showed significantly higher CPU overhead for tiny kernel regions of the workload. Microbenchmarks targeting just the CPU single thread performance demonstrated near identical performance on the partner and RA cluster nodes. This indicated that a 30-second perf record -a -g capture on the host, viewed with perf report , surfaced an unexpected top frame: 24% of CPU cycles spent on arm_smmu_cmdq_issue_cmdlist .
Figure 1. Linux perf icicle graph highlighting Arm SMMU command-queue invalidation overhead during DeepSeek-V3 FP8 pre-training on a virtualized NVIDIA GB200 NVL72 system
arm_smmu_cmdq_issue_cmdlist is the function that submits invalidation commands to the Arm SMMU’s command queue. Under virtualization, every map/unmap resulting in guest invalidation traps the host and serializes through a single command queue, producing the spinlock contention visible in the profile. Virtual Command Queue (VCMDQ) is a feature available through the Command Queue Virtualization extension to the standard Arm SMMUv3 and allows the guest to issue SMMU invalidation commands directly to hardware without VM exits.
The fix: Enable CMDQV/VCMDQ in the host kernel on the partner cluster and expose it to the guest. This requires a kernel built with the tegra241-cmdqv driver and the corresponding hypervisor support; recent QEMU/libvirt versions have added a cmdqv IOMMU attribute to expose it to guests.
After this change, linux perf showed arm_smmu_cmdq_issue_cmdlist falling out of the top frames and dTLB miss rates returning to bare-metal parity. The MoE iteration-time gap narrowed to within RA tolerance from 12%.
The takeaway is that Grace-based virtualized deployments need the VM stack to expose the right SMMU capabilities for memory-mapping-heavy workloads. With CMDQV/VCMDQ enabled in the host kernel and exposed to the guest, the platform can avoid unnecessary SMMU serialization and return MoE training performance to within RA tolerance.
The next layer down is the CPU itself, where the failure mode looks completely different.
Case study 2: H100 cluster losing 12% to CPU contention and NUMA misbinding
Layer: CPU power and process placement.
A partner’s H100 SXM5 cluster, running the same NCCL version and NeMo container as NVIDIA’s HGX RA, was running Llama 3 70B pre-training 12% slower than reference. Unlike the GB200 NVL72 case, this wasn’t a kernel-level issue; everything happened in user space and BIOS.
CPU frequency: turbostat -i 1 during training showed busy cores pegged at 3.0 GHz, despite the SKU being rated for 3.8 GHz turbo. Idle cores were also at 3.0 GHz, with C-states sitting in C1 rather than dropping to C6.
NUMA-remote traffic: numastat -p <python_pid> showed roughly 18% of the training process’s memory accesses going to the remote NUMA node
The CPU on the partner cluster was configured with C-states limited to C1 in BIOS. This is a common “low-latency” default that is actively wrong for AI training workloads. With idle cores held in C1, they continued to draw package power; the busy cores feeding the GPU with kernels couldn’t claim enough of the package power budget to hit turbo. Allowing the idle cores to drop to C6 freed power headroom, enabling the busy cores to climb to 3.8 GHz and recover roughly 4% on this workload.
The hypervisor housekeeping threads were pinned to the same physical cores as the training process’s data loader workers. Inside the VM this looked like sporadic 50–100 ms stalls in the python threads, which then propagated as the long tail in step time. The fix was a cpuset separation: hypervisor and host services on cores 0–7 and 56–63, training processes on the remainder.
Result: The 12% gap shrank to 3%, with the residual traced to a different NCCL tuning issue covered in the next case study.
The pattern here is that no single fix recovered the whole gap. The C-state change was the largest single contributor at ~4%, and the rest came from process isolation through NUMA binding. With CPU and virtualization addressed, the next ceiling is the network.
Case study 3: GB300 NVL72 with NVIDIA ConnectX-8 SuperNIC under-utilizing 1.6 Tbps fabric
Focus: ConnectX-8 SuperNIC collective tuning
A GB300 NVL72 deployment with NVIDIA ConnectX-8 SuperNICs (1.6 Tbps per node) showed a 31% training performance gap on Nemotron-4 15B pre-training. Single-node throughput looked healthy; the gap appeared at 512 GPUs, where the profiler showed exposed AllGather and ReduceScatter time. That pointed to the collective path on the ConnectX-8 fabric rather than compute.
The investigation tested several variables with NCCL Tests ( nccl-tests ), including iteration count, UCX/UCC behavior, NUMA mapping, NVLS, and NCCL versions. For the workload’s networking performance, the relevant tuning change was narrower: increasing NCCL_IB_QPS_PER_CONNECTION to 4 from the default value of 1.
Figure 2. Nemotron-4 15B performance with NCCL QPS optimization at 512-GPU scale
Nsight Systems trace showing communication overhead exposed at lower QPS values, contributing to longer training iteration time
Signal was visible in both the workload and the nccl-tests collective measurements. On the NVIDIA reference cluster, the default configuration ran at about 1.09s per iteration. With QPS=4 , the same reference workload improved to about 0.83s. In the profile, AllGather time dropped from about 375ms to 262ms, and ReduceScatter dropped from about 389ms to 273ms. The comparison run was about 0.76s and used a different NCCL version. The remaining difference was therefore partly attributable to an NCCL version mismatch between the comparison and reference environments; aligning the versions narrowed the residual gap further. Because NCCL version changes are outside the normal Exemplar tuning scope, the recommended tuning keeps the deployed NCCL version unchanged.
Lesson : Don’t increase QPS everywhere. QPS is fabric- and workload-dependent. On this GB300 ConnectX-8 workload, QPS=4 improved large-message AllGather and ReduceScatter behavior. On other fabrics or message-size profiles, the same setting may add CPU overhead without improving training throughput. The right approach is to test the collective at the workload’s real message sizes, sweep the setting on the target fabric, and verify the result in the training workload.
Case study 4: The environment variable that never made it inside
In a virt

[truncated]

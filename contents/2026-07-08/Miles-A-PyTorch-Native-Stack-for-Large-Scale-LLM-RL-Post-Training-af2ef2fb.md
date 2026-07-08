---
source: "https://pytorch.org/blog/miles-a-pytorch-native-stack-for-large-scale-llm-rl-post-training/"
hn_url: "https://news.ycombinator.com/item?id=48835609"
title: "Miles: A PyTorch-Native Stack for Large-Scale LLM RL Post-Training"
article_title: "Miles: A PyTorch-Native Stack for Large-Scale LLM RL Post-Training – PyTorch"
author: "gmays"
captured_at: "2026-07-08T19:07:51Z"
capture_tool: "hn-digest"
hn_id: 48835609
score: 1
comments: 0
posted_at: "2026-07-08T18:33:14Z"
tags:
  - hacker-news
  - translated
---

# Miles: A PyTorch-Native Stack for Large-Scale LLM RL Post-Training

- HN: [48835609](https://news.ycombinator.com/item?id=48835609)
- Source: [pytorch.org](https://pytorch.org/blog/miles-a-pytorch-native-stack-for-large-scale-llm-rl-post-training/)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T18:33:14Z

## Translation

タイトル: Miles: 大規模 LLM RL ポストトレーニング用の PyTorch ネイティブ スタック
記事のタイトル: Miles: 大規模 LLM RL ポストトレーニング用の PyTorch ネイティブ スタック – PyTorch

記事本文:
メインコンテンツにスキップ
PyTorch Conference North America にご参加ください · 10 月 20 ～ 21 日 · カリフォルニア州サンノゼ
検索
検索を閉じる
検索
メニュー
学ぶ
始めましょう
PyTorch の紹介 - YouTube シリーズ
Miles: 大規模 LLM RL ポストトレーニング用の PyTorch ネイティブ スタック
Miles は、大規模な LLM RL ポストトレーニング用の RadixArk のオープンソース フレームワークです。これは、ロールアウト用の SGLang、トレーニング用の NVIDIA Megatron-LM、Ray オーケストレーション、小型のプラグイン可能なトレーナーの背後にある PyTorch ネイティブの拡張性を構成しており、統合された低精度レシピ、MoE 対応のロールアウト/トレーニング調整、高速な NVIDIA NCCL/RDMA 重み同期、可観測性、フォールト トレランスが組み込まれており、フロンティア スケールの LLM RL の構築、再現、操作が容易になります。
強化学習は、トレーニング後の大規模な言語モデルの中心的な部分になっています。しかし、モデルが大きくなり、高密度から専門家混合 (MoE) に移行し、より分散化された特殊なハードウェア (NVIDIA Blackwell や Hopper シリーズなど) で実行されるにつれて、RL ポストトレーニングは単なるトレーニング ループではなくなりました。これは分散システムの問題です。
最新の LLM RL フレームワークは、いくつかの可動部分を調整する必要があります。
ロールアウト ワーカーは高スループットでサンプルを生成する必要があります。
トレーナーはこれらのサンプルを効率的に使用し、安定したポリシーの更新を計算する必要があります。
ロールアウト ポリシーとトレーニング ポリシーは同期を保つ必要があります。
大規模な MoE モデルでは、ロールアウトとトレーニングの間で調整を維持する必要があるルーティング動作が導入されます。
低精度のレシピは、パイプライン全体で一貫して動作する必要があります。
長時間実行されるジョブには、最初から可観測性、チェックポイント、フォールト トレランスが必要です。
Miles はこの設定に合わせて構築されました。
Miles は、LLM ポストトレーニング用の RadixArk のオープンソース強化学習フレームワークです。高スループットを実現するために SGLang 上にネイティブに構築されています

スケーラブルなトレーニングのために Megatron-LM をロールアウトして深く統合し、Ray を使用して分散システムを調整し、PyTorch をスタック全体で共通のプログラミングおよび数値レイヤーとして維持します。
目標はシンプルです。研究者やインフラストラクチャ チームがカスタマイズできるコア トレーナーを十分に小さくしながら、大規模な LLM RL トレーニングをより構成可能、再現可能、拡張しやすくすることです。
マイルズは、中心が小さく、エッジが多いという哲学に従っています。
コアトレーニングループは意図的にコンパクトになっています。ユーザーが最も頻繁に変更したい部分 (ロールアウト ロジック、報酬計算、損失関数、サンプル フィルタリング、メトリクス、トレーニング ループ フック) は、ユーザー指定の Python モジュールを通じて起動時に接続されます。これにより、チームはフレームワークをフォークすることなく、システムを新しいアルゴリズムや生産上の制約に適応させることができます。
その小さなコアの下で、マイルズは 4 つの主要なシステムを構成します。
高スループットのロールアウト生成用の SGLang。
スケーラブルな分散トレーニング用の Megatron-LM。
Ray はクラスター オーケストレーション、アクターのライフサイクル、スケジューリング、および監督に使用します。
モデル、autograd、分散プリミティブ、dtype サポート、拡張性、プロファイリング用の PyTorch。
この構成が重要です。 RL ポストトレーニングでは、生成とトレーニングを連携させる必要がありますが、2 つのフェーズのパフォーマンス プロファイルは大きく異なります。ロールアウトはメモリ帯域幅に制限されます (デコード中は KV キャッシュとパラメーターの読み取りが支配的です)。一方、トレーニングはコンピューティングに制限され、通信に重点が置かれます。重量の同期、サンプル転送、チェックポイント変換、ルーティングの一貫性、および低精度の動作はすべて、境界を越えて慎重に処理する必要があります。
この投稿の残りの部分では、Miles がその境界の各部分をどのように処理するか (Ray によるオーケストレーション、Megatron-LM によるスケーリング、拡張性) について説明します。

th PyTorch と、すぐに使えるもの。
Ray: 長時間実行される RL ジョブのオーケストレーション
Miles は、Ray 分散ランタイム上に直接構築されています。 Miles の実行では、存続期間の長いプロセスはすべて Ray アクターとして表されます。トレーナー ランク、SGLang ロールアウト サーバー、ルーティング プロキシ、および非同期ロールアウト ワーカーはすべて、Ray のアクター モデル内に存在します。
これにより、Miles はクラスター規模の RL ワークロードのための自然な基盤を得ることができました。
Miles は、アクターの配置に Ray の GPU 対応スケジューラと配置グループを使用し、起動時の Ray 配置仕様により、分離 (ロールアウトとトレーニングを別々のノードで行う) およびコロケート (ロールアウトとトレーニングを同じノードで行う) レイアウトをサポートします。ラック内の問題 (不良 GPU とフルラックの問題を区別するなど) を区別することは必ずしも簡単ではないため、慎重なコロケーション、スペア ノードの予約、およびエラー分離の鍵を容易にするために、プロセスの配置はラックを意識する必要があります。
RL パイプライン全体でのデータの移動
プロンプト、サンプル、および更新された重みは、ロールアウト アクターとトレーナー ランクの間で継続的に循環し、マイルズはレイ アクターとタスクを使用してそのフローを調整します。バルクウェイト転送の場合、Ray はテンソルバイトが専用の NCCL/RDMA チャネル上を移動する間に制御パスを処理し、Miles に Ray レベルのプログラマビリティと大規模データの高速パスの両方を提供します。
Miles の実行はエンドツーエンドの Ray ジョブであるため、ボルトオンのインフラストラクチャを使用せずに、Ray のオペレータ サーフェス (ジョブの送信、作業者の監督、ログ集計、ダッシュボードの可視性) を継承します。フォールト トレランスを有効にすると、Miles は障害が発生したランクを回復し、同じ Ray 基板上で 1 週間にわたるワークロードを移動し続けることができます。
完全非同期 RL のサポート
Ray アクターは永続的で、独自の状態を保持し、独立してスケジュールされるため、Miles は完全な非同期モードを実行できます。

ut とトレーニングが相互にブロックされなくなりました。ロールアウト アクターは、トレーナーが独自のペースで排出するキューにサンプルを継続的にストリーミングします。
Megatron-LM: トレーニング バックエンドのスケーリング
Miles は、実稼働トレーニング バックエンドとして Megatron-LM を使用し、ブラックボックス ライブラリとしてラップするのではなく、Megatron の引数パーサー、モデル構築パイプライン、トレーニング ループ、並列処理プリミティブ、分散チェックポイント形式に直接接続します。これにより、マイルズは、ユーザー向けのクリーンなワークフローを維持しながら、最先端規模の密度の高い MoE トレーニングに必要なインフラストラクチャを提供できるようになりました。
Megatron-LM はすでに大規模な分散トレーニング構成サーフェス (シーケンスの長さ、ロータリー埋め込み、グループ化された GEMM、あらゆる種類の並列処理、オプティマイザー設定、アクティベーション チェックポイント設定など) を公開しており、Miles はそれをラップしたり再宣言したりするのではなく、直接再利用します。ユーザーは、マイル固有のオプションと標準のメガトロン オプションを組み合わせた 1 つの起動スクリプトを通じてマイルの実行を設定し、構成レイヤーの重複を回避し、トレーニング セットアップを上流のメガトロンの動作に近づけます。
長寿命フォークの代わりにモデルのスペックを採用
フロンティア アーキテクチャは急速に変化し、新しいアテンション ブロック、ルーティング メカニズム、エキスパート レイアウトがモデル ファミリ全体に導入されるため、Miles はプラグイン モデル 仕様 (カスタム PyTorch コンポーネント (ゲートされたアテンション出力モジュール、Gated-Delta-Net ブロック、モデル固有の MoE ルーターなど) をメガトロンのモデル パイプラインに直接挿入する小さな仕様ファイル) を通じてそれらを処理します。これにより、Miles は、アップストリームから常に分岐する長寿命の Megatron フォークを維持することなく、新しいアーキテクチャ (たとえば、DeepSeek-V3/V4、GLM-4.7、Qwen3 MoE バリアント) をサポートできるようになります。
並列処理を意識したチェックポイント設定
マイルズはメガトロンのパラレルを使用します

sm 対応の分散チェックポイント形式なので、モデルを Hugging Face から一度変換すると、重みを最初から再変換することなく、さまざまなテンソル / パイプライン / コンテキスト / エキスパート並列構成にロードできます。大規模なトレーニング ジョブを運用しているチームにとって、これは、モデルやクラスターの形状が変更されるたびに、チェックポイント変換と並列処理の変更が別個のエンジニアリング プロジェクトにならないことを意味します。
バックエンドにパッチを当てずにトレーニングを拡張する
Miles はトレーニング ループの明確に定義されたポイント (モデルの初期化後、対数確率の計算前、各トレーニング ステップの前) でフックを公開するため、ユーザーは Megatron の内部を編集することなく、補助損失、カスタム メトリクス、サンプル レベルの診断、クリッピング ルール、またはアルゴリズム固有の動作を追加できます。設計目標はシンプルです。バックエンドを強力に保ちながら、ユーザーのカスタマイズをバックエンドの外部に維持することです。
PyTorch: モデル、数値、拡張性の共通レイヤー
PyTorch は Miles 内の共通プログラミング モデルです。モデル コンポーネントは通常の torch.nn.Modules 、損失は標準の autograd グラフ、混合精度、勾配チェックポイント、分散プリミティブ、およびプロファイリングはすべて、使い慣れた PyTorch ワークフロー内にあります。 LLM RL ポストトレーニングは急速に変化するため、これは重要です。チームは、毎回新しい抽象化を学習することなく、新しい報酬、損失、ルーター、モデル モジュール、およびデバッグ ツールを追加する必要があります。
PyTorch ネイティブ モデルの拡張性
Miles のプラグイン モデル仕様メカニズムは torch.nn.Modules を中心に構築されているため、新しいアーキテクチャをサポートするということは、新しいコンポーネントを通常の PyTorch コードとして記述し、それを Megatron のモデル パイプラインに接続することを意味します。autograd、混合精度、勾配チェックポイント、およびモジュール ライフサイクルはすべて、PyTorch ユーザーが期待する方法で動作し続けます。チームには何もありません

o モデルを別の中間抽象化に変換して、Miles 上で実行できるようにします。
PyTorch ネイティブ RL のカスタマイズ
同じ原理が RL アルゴリズムにも当てはまります。ロールアウト関数、報酬、損失関数、サンプル フィルター、メトリクス、トレーニング ループ フックはすべて、トレーニング グラフの残りの部分を構成する標準の PyTorch 操作を使用して、起動時に提供される Python モジュールを通じてカスタマイズされます。チームは既存のレシピから開始して、トレーナーを書き直すことなく、報酬の置き換え、補助損失の追加、サンプル フィルタリングの変更、または新しい診断の計測を行うことができます。
パイプライン全体にわたる低精度のレシピ
Miles は、分離されたバックエンドのみの機能としてではなく、トレーニングとロールアウトにまたがる BF16、FP8、MXFP8、および INT4-QAT レシピを使用して、PyTorch の dtype システム上に低精度パイプラインを構築します。この一貫性は RL にとって重要です。サンプルの生成に使用されるポリシーとトレーニング ログの確率の計算に使用されるポリシーは一致していなければならず、Miles はこれらの数値選択を明示的かつ再現可能にするように設計されているためです。
使い慣れたツールでのプロファイリングとデバッグ
大規模な RL パフォーマンスの問題は、ロールアウト レイテンシー、トレーニング コンピューティング、集団通信、データ移動、重みの同期、サンプル フィルタリング、スケジュールなど、あらゆる場所で表面化する可能性があるため、Miles は PyTorch プロファイラーに接続して、標準ツールで検査するためにトレーニング フェーズの Chrome トレースをキャプチャしました。 Megatron の PyTorch ベースのバックエンドおよびサポートされている場合のグラフ コンパイル パスと組み合わせることで、使い慣れた PyTorch エコシステム内でデバッグとパフォーマンスの作業を継続できます。
マイルがすぐに提供できるもの
Miles は、大規模な LLM RL ポストトレーニングに必要なコア システム機能を提供するように設計されています。
ロールアウトとトレーニングの統合 — SGLang ロールアウトを Megatron-LM に接続します

さまざまな GPU 予算と使用率目標に合わせて、分散実行と併置実行の両方を行うトレーニング。
非同期実行 - 完全な非同期モードでは、ロールアウトをトレーニングから切り離します。ロールアウト アクターは、トレーナーが独自のペースで排出するキューにサンプルを継続的にストリーミングし、2 つのフェーズ間の反復ごとのブロックを排除します。
高速な重み同期 — 各トレーニングの更新後、新しい重みが専用の NCCL/RDMA チャネル経由でロールアウト ワーカーに流れ、Ray は制御パスのみを処理するため、バルク テンソル バイトは Python データ パスから外されます。
MoE 対応のロールアウト/トレーニング調整 — ロールアウト ルーティング リプレイは、ロールアウト/トレーニングの境界を越えてルーティングの決定を保持し、MoE RL を不安定にするトレーナー対ロールアウトのルーティングの不一致を軽減します。
低精度のサポート — 独立したトレーニング専用レシピとしてではなく、エンドツーエンド RL スタックの一部として設計された統合 BF16 / FP8 / MXFP8 / INT4-QAT パイプライン。
ロールアウトとトレーニングにわたる LoRA — LoRA はロールアウト パスとトレーニング パスの両方でサポートされており、パラメータ効率の高いポストトレーニングを可能にして、コストを削減し、大規模なベース モデルでの反復を高速化します。
耐障害性と可観測性 — Ray のジョブとアクター モデルは、監視、ログ集約、ダッシュボードの可視性を提供すると同時に、ランク付けを行います。

[切り捨てられた]

## Original Extract

Skip to main content
Join us at PyTorch Conference North America · Oct 20-21 · San Jose, CA
Search
Close Search
search
Menu
Learn
Get Started
Intro to PyTorch – YouTube Series
Miles: A PyTorch-Native Stack for Large-Scale LLM RL Post-Training
Miles is RadixArk’s open source framework for large-scale LLM RL post-training. It composes SGLang for rollout, NVIDIA Megatron-LM for training, Ray orchestration, and PyTorch-native extensibility behind a small, pluggable trainer, with unified low-precision recipes, MoE-aware rollout/training alignment, fast NVIDIA NCCL/RDMA weight synchronization, observability, and fault tolerance built in — making frontier-scale LLM RL easier to build, reproduce, and operate.
Reinforcement learning has become a central part of post-training large language models. But as models become larger, transition from dense to mixture-of-experts (MoE), and run across more distributed and specialized hardware (e.g. NVIDIA Blackwell and Hopper series), RL post-training is no longer just a training loop. It is a distributed systems problem.
A modern LLM RL framework needs to coordinate several moving pieces:
Rollout workers must generate samples at high throughput.
Trainers must consume those samples efficiently and compute stable policy updates.
The rollout policy and training policy must stay synchronized.
Large MoE models introduce routing behavior that must remain aligned across rollout and training.
Low-precision recipes need to work consistently across the full pipeline.
Long-running jobs need observability, checkpointing, and fault tolerance from the start.
Miles was built for this setting.
Miles is RadixArk’s open source reinforcement learning framework for LLM post-training. It is built natively on SGLang for high-throughput rollout and integrates deeply with Megatron-LM for scalable training, uses Ray to orchestrate the distributed system, and keeps PyTorch as the common programming and numerical layer throughout the stack.
The goal is simple: make large-scale LLM RL training more composable, reproducible, and easier to scale, while keeping the core trainer small enough for researchers and infrastructure teams to customize.
Miles follows a small-core, many-edges philosophy.
The core training loop is intentionally compact. The pieces that users most often want to change — rollout logic, reward computation, loss functions, sample filtering, metrics, and training-loop hooks — are attached at launch time through user-supplied Python modules. This lets teams adapt the system to new algorithms and production constraints without forking the framework.
Underneath that small core, Miles composes four major systems:
SGLang for high-throughput rollout generation.
Megatron-LM for scalable distributed training.
Ray for cluster orchestration, actor lifecycle, scheduling, and supervision.
PyTorch for models, autograd, distributed primitives, dtype support, extensibility, and profiling.
This composition is important. RL post-training requires generation and training to work together, but the two phases have very different performance profiles: rollout is memory-bandwidth-bound (KV-cache and parameter reads dominate during decoding), while training is compute-bound and communication-heavy. Weight synchronization, sample transfer, checkpoint conversion, routing consistency, and low-precision behavior all need to be handled carefully across the boundary.
The rest of this post walks through how Miles handles each piece of that boundary — orchestration with Ray, scaling with Megatron-LM, extensibility with PyTorch, and what comes out of the box.
Ray: Orchestrating Long-Running RL Jobs
Miles is built directly on the Ray distributed runtime. In a Miles run, every long-lived process is represented as a Ray actor: trainer ranks, SGLang rollout servers, routing proxies, and asynchronous rollout workers all live inside Ray’s actor model.
This gives Miles a natural foundation for cluster-scale RL workloads.
Miles uses Ray’s GPU-aware scheduler and placement groups for actor placement, supporting disaggregated (rollout and training on separate nodes) and colocated (rollout and training on the same nodes) layouts via launch-time Ray placement specs. Process placement must be rack-aware to facilitate careful colocation, reserving spare nodes, and key for error isolation, since isolating problems within a rack (e.g., distinguishing a bad GPU from a full rack issue) is not always straightforward.
Moving data across the RL pipeline
Prompts, samples, and updated weights cycle continuously between rollout actors and trainer ranks, and Miles uses Ray actors and tasks to coordinate that flow. For bulk weight transfer, Ray handles the control path while the tensor bytes move over dedicated NCCL/RDMA channels, giving Miles both Ray-level programmability and a fast path for large data.
Because a Miles run is a Ray job end-to-end, it inherits Ray’s operator surface — job submission, worker supervision, log aggregation, and dashboard visibility — without bolt-on infrastructure. With fault tolerance enabled, Miles can recover failed ranks and keep week-long workloads moving on top of the same Ray substrate.
Supporting fully asynchronous RL
Because Ray actors are persistent, hold their own state, and are scheduled independently, Miles can run a fully asynchronous mode in which rollout and training no longer block on each other — rollout actors continuously stream samples into a queue that the trainer drains at its own pace.
Megatron-LM: Scaling the Training Backend
Miles uses Megatron-LM as its production training backend, plugging directly into Megatron’s argument parser, model-construction pipeline, training loop, parallelism primitives, and distributed checkpoint format rather than wrapping it as a black-box library. That gives Miles the infrastructure needed for frontier-scale dense and MoE training while preserving a clean user-facing workflow.
Megatron-LM already exposes a large distributed-training configuration surface — sequence length, rotary embeddings, grouped GEMM, all flavors of parallelism, optimizer settings, activation checkpointing, and more — and Miles reuses it directly rather than wrapping or re-declaring it. Users configure a Miles run through one launch script that combines Miles-specific options with standard Megatron options, avoiding duplicated configuration layers and keeping the training setup close to upstream Megatron behavior.
Model specs instead of long-lived forks
Frontier architectures change quickly, with new attention blocks, routing mechanisms, and expert layouts arriving across model families, so Miles handles them through plug-in model specs — small spec files that insert custom PyTorch components (for example, a gated attention-output module, a Gated-Delta-Net block, or a model-specific MoE router) directly into Megatron’s model pipeline. This lets Miles support new architectures — for example DeepSeek-V3/V4, GLM-4.7, and Qwen3 MoE variants — without maintaining a long-lived Megatron fork that constantly diverges from upstream.
Parallelism-aware checkpointing
Miles uses Megatron’s parallelism-aware distributed checkpoint format, so a model can be converted from Hugging Face once and then loaded across different tensor / pipeline / context / expert parallel configurations without re-converting weights from scratch. For teams operating large training jobs, this means checkpoint conversion and parallelism changes don’t become a separate engineering project every time the model or cluster shape changes.
Extending training without patching the backend
Miles exposes hooks at well-defined points in the training loop — after model initialization, before log-probability computation, and before each training step — so users can add auxiliary losses, custom metrics, sample-level diagnostics, clipping rules, or algorithm-specific behavior without editing Megatron internals. The design goal is simple: keep the backend powerful, but keep user customization outside it.
PyTorch: The Common Layer for Models, Numerics, and Extensibility
PyTorch is the common programming model inside Miles: model components are regular torch.nn.Modules , losses are standard autograd graphs, and mixed precision, gradient checkpointing, distributed primitives, and profiling all stay inside familiar PyTorch workflows. This matters because LLM RL post-training changes fast — teams need to add new rewards, losses, routers, model modules, and debugging tools without learning a new abstraction each time.
PyTorch-native model extensibility
Miles’ plug-in model-spec mechanism is built around torch.nn.Modules , so supporting a new architecture means writing the new component as ordinary PyTorch code and connecting it into Megatron’s model pipeline — autograd, mixed precision, gradient checkpointing, and module lifecycle all keep working the way PyTorch users expect. Teams don’t have to translate the model into a separate intermediate abstraction to get it running on Miles.
PyTorch-native RL customization
The same principle applies to RL algorithms: rollout functions, rewards, loss functions, sample filters, metrics, and training-loop hooks are all customized through Python modules provided at launch time, using standard PyTorch operations that compose with the rest of the training graph. A team can start from an existing recipe and replace the reward, add an auxiliary loss, change sample filtering, or instrument new diagnostics without rewriting the trainer.
Low-precision recipes across the pipeline
Miles builds its low-precision pipeline on PyTorch’s dtype system, with BF16, FP8, MXFP8, and INT4-QAT recipes that span training and rollout rather than living as isolated backend-only features. This consistency matters for RL because the policy used to generate samples and the policy used to compute training log probabilities must stay aligned, and Miles is designed to make those numerical choices explicit and reproducible.
Profiling and debugging in familiar tools
Large-scale RL performance issues can surface anywhere — rollout latency, training compute, collective communication, data movement, weight synchronization, sample filtering, or scheduling — so Miles wires in the PyTorch profiler to capture Chrome traces of training phases for inspection in standard tooling. Combined with Megatron’s PyTorch-based backend and graph-compile paths where supported, this keeps debugging and performance work inside the familiar PyTorch ecosystem.
What Miles Provides Out of the Box
Miles is designed to provide the core systems features needed for large-scale LLM RL post-training:
Rollout and training integration — Connects SGLang rollout with Megatron-LM training, with both disaggregated and colocated execution to fit different GPU budgets and utilization targets.
Asynchronous execution — Fully async mode decouples rollout from training: rollout actors stream samples continuously into a queue that the trainer drains at its own pace, eliminating the per-iteration blocking between the two phases.
Fast weight synchronization — After each training update, fresh weights flow to rollout workers over dedicated NCCL/RDMA channels, with Ray handling only the control path so bulk tensor bytes stay off the Python data path.
MoE-aware rollout/training alignment — Rollout Routing Replay preserves routing decisions across the rollout/training boundary, reducing the trainer-vs-rollout routing mismatch that would otherwise destabilize MoE RL.
Low-precision support — A unified BF16 / FP8 / MXFP8 / INT4-QAT pipeline designed as part of the end-to-end RL stack rather than as isolated training-only recipes.
LoRA across rollout and training — LoRA is supported in both rollout and training paths, enabling parameter-efficient post-training that reduces cost and speeds up iteration on large base models.
Fault tolerance and observability — Ray’s job and actor model provide supervision, log aggregation, and dashboard visibility, while rank

[truncated]

---
source: "https://www.redhat.com/en/blog/cpu-back-rethinking-cpu-gpu-split-llm-inference"
hn_url: "https://news.ycombinator.com/item?id=49221089"
title: "The CPU is back: Rethinking the CPU-GPU split for LLM inference"
article_title: "The CPU is back: Rethinking the CPU-GPU split for LLM inference"
author: "eigenBasis"
captured_at: "2026-08-08T12:26:28Z"
capture_tool: "hn-digest"
hn_id: 49221089
score: 3
comments: 0
posted_at: "2026-08-08T12:16:03Z"
tags:
  - hacker-news
  - translated
---

# The CPU is back: Rethinking the CPU-GPU split for LLM inference

- HN: [49221089](https://news.ycombinator.com/item?id=49221089)
- Source: [www.redhat.com](https://www.redhat.com/en/blog/cpu-back-rethinking-cpu-gpu-split-llm-inference)
- Score: 3
- Comments: 0
- Posted: 2026-08-08T12:16:03Z

## Translation

タイトル: CPU が帰ってきた: LLM 推論のための CPU と GPU の分割を再考する
説明: エージェント AI が CPU 推論への移行を推進している理由。

記事本文:
メニュー Red Hat を探索する
AI
概要
ハイブリッドクラウド
プラットフォームソリューション
人工知能 AI モデルとアプリを構築、展開、監視します。
Linux の標準化 オペレーティング環境全体で一貫性を確保します。
アプリケーション開発 アプリケーションの構築、展開、管理の方法を簡素化します。
自動化 自動化を拡張し、技術、チーム、環境を統合します。
仮想化 仮想化およびコンテナ化されたワークロードの運用を最新化します。
デジタル主権 重要なインフラを制御し、保護します。
セキュリティをコーディングし、セキュリティに重点を置いたソフトウェアを構築、展開、監視します。
エッジ コンピューティング エッジ テクノロジーを使用してワークロードをソースの近くに展開します。
Red Hat® Hybrid Cloud Console で、自分のペースでクラウド製品とソリューションを使用する方法を学びましょう。
Red Hat AI アイコン 人工知能、Red Hat Enterprise Linux AI、Red Hat OpenShift AI、RHEL AI、機械学習 38 38 2025-03-12T19:43:40.963Z image/svg+xml Red Hat AI アイコン 人工知能、Red Hat Enterprise Linux AI、Red Hat OpenShift AI、RHEL AI、機械学習 アイコンなし2025-03-12T19:39:59.817Z テクノロジー アイコン 標準 Red Hat AI ハイブリッド クラウド全体で AI ソリューションを開発およびデプロイします。
Red Hat Enterprise Linux アイコン RHEL、Linux プラットフォーム、CentOS 2024-03-01T15:26:42.958Z 保留中 TRA3b65dd25-844d-49bb-93c1-30f5b34684f1 アイコン 2024-03-01T15:26:42.958Z 真の保留中2024-03-21T00:40:29.326Z rhcc-audience:internal いいえ テクノロジー アイコン DER3b65dd25-844d-49bb-93c1-30f5b34684f1 標準 はい rhcc-product:red-hat-enterprise-linux テクノロジー アイコン image/svg+xml 2024-05-10T14:11:29.114Z Red Hat Enterprise Linux icon RHEL、Linux プラットフォーム、CentOS Activate Activate 2024-05-10T14:11:29.836Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:11:29.836Z workflow-process-service 2024-05-10T14:11:29.836Z

テクノロジーアイコンを使用して、Red Hat 製品とコンポーネントを表します。 Do not remove the icon from the bounding shape. Red Hat Enterprise Linux 柔軟なオペレーティング システムでハイブリッド クラウド イノベーションをサポートします。
Red Hat OpenShift アイコン クラウド、コンテナー、Kubernetes 2024-03-01T15:26:53.684Z 保留中 TRA9ec76aa9-ef09-4c49-8816-01dd13970ca7 アイコン 2024-03-01T15:26:53.684Z 真の保留中2024-03-21T00:39:44.126Z rhcc-audience:internal いいえ テクノロジーアイコン DER9ec76aa9-ef09-4c49-8816-01dd13970ca7 標準 はい rhcc-product:red-hat-openshift rhcc-product:red-hat-openshift-on-ibm-cloud rhcc-product:microsoft-azure-red-hat-openshift rhcc-product:red-hat-openshift-service-on-aws rhcc-product:red-hat-openshift-container-platform rhcc-product:red-hat-openshift-platform-plus テクノロジーアイコン画像/svg+xml 2024-05-10T14:18:23.703Z Red Hat OpenShift アイコン クラウド、コンテナー、Kubernetes Activate Activate 2024-05-10T14:18:25.221Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:18:25.221Z workflow-process-service 2024-05-10T14:18:25.221Z テクノロジー アイコンを使用して、 represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat OpenShift アプリを大規模に構築、モダナイズ、デプロイします。
Red Hat Ansible Automation Platform アイコン 管理、エッジ 2024-03-01T15:26:35.068Z 保留中 TRA759b57c4-760b-45a0-a939-821f47181964 アイコン 2024-03-01T15:26:35.068Z 真の保留中2024-03-21T00:39:55.923Z rhcc-audience:internal いいえ テクノロジー アイコン DER759b57c4-760b-45a0-a939-821f47181964 標準 はい rhcc-product:red-hat-ansible-automation-platform テクノロジー アイコン image/svg+xml 2024-05-10T14:04:00.014Z Red Hat Ansible Automation Platform アイコン 管理、エッジ アクティブ化 アクティブ化 2024-05-10T14:04:01.784Z workflow-process-service アクティブ化 workflow-process-service false 2024-05-10T1

4:04:01.784Z workflow-process-service 2024-05-10T14:04:01.784Z テクノロジー アイコンを使用して Red Hat 製品とコンポーネントを表します。境界図形からアイコンを削除しないでください。 Red Hat Ansible Automation Platform 企業全体の自動化を実装します。
Red Hat OpenShift 仮想化エンジン
主要なクラウドプロバイダーとの統合
トレーニング
トレーニングと認定
Red Hat 認定システム管理者試験
Red Hat システム管理 I
Red Hat Learning サブスクリプションのトライアル (無料)
Red Hat 認定エンジニア試験
Red Hat 認定 OpenShift 管理者試験
クラウドネイティブのアプリケーションとサービスの構築、配信、管理に役立つリソースとツールを見つけてください。
信頼できるパートナーを活用したソリューションを構築する
Red Hat® Ecosystem Catalog の専門家とテクノロジーの協力コミュニティからソリューションを見つけてください。
閲覧中に好みのリソースをお勧めします。とりあえずこれらの提案を試してみてください。
ログイン Red Hat アカウントでさらに入手
一部のサービスでは購読が必要な場合があります。
口座番号: [[口座番号]]
Red Hat Ansible 自動化プラットフォーム
ポッドキャスト
Chris Wright と技術的に話す
その他のブログ
Red Hat 開発者ブログ
CPU が帰ってきました: LLM 推論のための CPU と GPU の分割を再考する
過去 3 年間、グラフィックス プロセッシング ユニット (GPU) が大規模言語モデル (LLM) の話題を独占してきました。従来のチャットボット アプリケーションでは、中央処理装置 (CPU) がリクエストあたりの総コンピューティングの一部を提供し、GPU が重労働を処理します。ただし、推論は 1 つの質問に答える 1 つのモデルではありません。ツール呼び出し、複数ステップの推論、小規模で特殊なモデルにわたるオーケストレーションへの依存が高まっているため、コンピューティングをどこに配置すべきかの計算が変化しています。 Intel は、CPU と GPU の比率が 1:8 から 1:8 に移行していることに注目して、この変化を訴えています。

ワークロードを 1:1 に調整し、エージェント展開では場合によっては 4:1 に調整します。
ここでは、LLM 推論に GPU を当然の選択肢とする前提がなぜ再交渉されるのか、何が CPU ベースのサービスに対する新たな需要を引き起こしているのか、そして業界がどこに向かっているのかについてデータが何を示しているのかを検討します。
CPU と GPU は本質的に (冗談ではありません)、根本的に異なる問題を解決します。
最新の GPU には、数千のデータ要素に対して同じ操作を同時に実行するように設計された数万のコアが含まれています。これにより、推論のためのトランスフォーマーの順方向パスを支配する高密度行列の乗算が非常に高速になります。トレーニング中および同時実行性の高いバッチ推論中、その並列処理はスループットに直接変換され、1 秒あたりのトークン (TPS) が増加し、コンピューティング 1 ドルあたりにより多くのリクエストが処理されます。
対照的に、最新の CPU には、シーケンシャル、条件付き、分岐ロジック向けに最適化された 1 から数百のコアが搭載されています。複雑なデシジョン ツリーを通過する単一の操作が特に高速です。これらはメイン システム メモリに直接アクセスでき、あらゆるモデルをラップするオーケストレーション層にとって自然な実行環境となります。ツールのディスパッチ、コードの実行、Python ランタイム、サンドボックス、入出力 (I/O)、およびエージェント ループ制御フローはすべて CPU 内にあります。
これらのアーキテクチャは性質が異なりますが、競合するものではありません。それらは連携して最も効果的に機能します。本当の問題は、どのアーキテクチャが優れているかではなく、どのワークロードがどこに属するかということです。
この分業は、結局のところ、FLOPS と命令のレイテンシという、作業の測定方法に依存します。
GPU は、1 秒あたりの浮動小数点演算 (FLOPS) によって生きたり消えたりします。 AI モデルは乗算および加算される 10 進数の巨大なウェブであるため、GPU の仕事はブルート処理です。

これらの行列計算を何兆回も同時に強制します。これは完全に生の数学的スループットのために構築されています。
対照的に、CPU は命令レイテンシーに特化しています。命令レイテンシは、単一コアがさまざまなコマンドの予測不可能なチェーンをどれだけ速く実行できるかを測定します。 CPU コアは、JavaScript Object Notation (JSON) の解析、ネットワーク I/O の処理、またはセキュリティ権限のチェックに必要な迅速なロジック シャッフルに優れています。
GPU にカオスな Python ランタイムを強制的に実行させると、その大規模な FLOP 容量がアイドル状態になり、絶え間ないタスク切り替えによって圧迫されてしまいます。 CPU に LLM の計算処理を強制すると、完璧に動作しますが、大規模な並列パイプラインがないため、時間がかかります。 GPU は数学的な筋肉です。 CPU はショーを指揮するロジック エンジンです。
従来の推論スタック: CPU が乗客であった場合
従来の推論スタックについて考えるとき、私たちはアプリケーションを提供するチャットボットを思い浮かべます。ここでは、CPU がサポート的な役割を果たすことがよくあります。
リクエストは API サーバーに到着し、CPU がリクエストをトークン化してスケジュールします。 CPU はこれらすべての初期タスクを処理します。次に、リクエストを GPU に渡してフォワード パスを実行します。これがコンピューティング バジェットの大半を占めます。 GPU はアテンションを実行し、レイヤーをフィードフォワードし、短い CPU 同期で次のトークンをサンプリングして取得し、スケジューラーを更新します。その後、シーケンス終了トークンが発行されるまで繰り返されます。最後に、CPU は出力を収集し、それを返してセッションを終了します。
このモデルでは、CPU が受付係として機能し、調整と組織化作業を実行し、GPU が計算上の重労働を実行します。トレーニング時代の AI データ センターの CPU と GPU の比率は、これを反映しています。GPU 8 個ごとに CPU がおよそ 1 ～ 2 個であり、この比率は CPU 需要よりも GPU のスループットによって大きく左右されます。チームがプロビジョニングした CPU

GPU に電力を供給し続けるためです。 CPU 推論の進歩により全体的な効率が向上するにつれて、このセットアップは徐々に事実上の標準から遠ざかりつつあります。
最初の要因: エージェントがワークロードをどのように変化させているか
エージェント AI の台頭により、新しい推論プロファイルが登場し、CPU 固有のコンピューティングの必要性が増大しています。
「Agentic AI」というと、Hermes や Openclaw のような統合エージェント ハーネスや、コーディング アシスタントに結び付けられた多数のエージェントを持つ開発者のイメージが思い浮かびます。これらのユースケースは人気が高まっていますが、エージェント AI は一見したほどニッチなものではありません。実際、多くの人はチャットボットにクエリを実行するという「従来の」プロセスを使用していません。 Claude や ChatGPT などの人気のある AI アシスタントは、より複雑な応答に「推論」ステージを組み込むことが多い巨大なモデルによって支えられています。推論中、「チャットボット」は完全なエージェント システムになります。ユーザーに表示される単一のタスクは、数十の個別のモデル呼び出しに分解される場合があり、それぞれが短く、コンテキストに依存し、条件付きで分岐します。
エージェント システムでは、モデルは単一のプロンプトを発行せず、応答を待ちます。アクションの計画を生成し、ツール呼び出しを実行し、返されたすべてのデータをコンパイルし、最終的にアクションを完了するか、ユーザーに回答を返します。 CPU は、その出力を解析し、どのツールを呼び出すかを判断し、API 呼び出しまたはコードを実行し、結果を収集してフィードバックすることを担当します。その後、ループが繰り返されます。場合によっては、ユーザー要求ごとに数十回、並列サブエージェントが CPU 調整作業を追加します。モデルはすぐに、より大きな推論ループの 1 つのコンポーネントになります。組織的な作業と純粋な計算の比率が変化し、プロセス全体において CPU がはるかに大きな役割を果たし、潜在的なボトルネックになります。
コラボレーションでは

ジョージア工科大学とインテルの研究者ら 2 人によると、エージェントのワークロードでは、CPU 側ツールの処理がエンドツーエンドの総遅延の 50 ～ 90% を占めていることがわかりました。 Intel の CEO、Lip-Bu Tan は Computex 2026 でこの変化を正確に説明し、「強化学習、オーケストレーション、エージェントには CPU がはるかに適しています」と述べました。
インテルの 2026 年第 1 四半期決算発表の数字は、これがすでにどれほど進んでいるかを定量化しています。
トレーニング ワークロード: 8 GPU あたり最大 1 CPU
推論ワークロード: すでに 4 GPU あたり最大 1 CPU に移行
エージェント ワークロード : 1:1 に向けて収束しており、一部の顧客は GPU ごとに 4 つの CPU がデプロイされていると報告しています
インテルの 2026 年第 1 四半期のデータセンターおよび AI 部門の収益は、需要が供給を上回り、前年比 22% 増の 51 億ドルとなりました。 Intelは消費者向けチップの生産の優先順位を下げ、ファブの能力をサーバー用Xeon部品に振り向けた。
Arm 独自の分析では、この変化がさらに顕著になると予測されています。従来の AI データセンターには、容量 1 ギガワット (GW) あたり約 3,000 万個の CPU コアが必要です。 Arm CEO の Rene Haas 氏は、AI エージェントの時代には、その数字は GW あたり 1 億 2,000 万 CPU コアに増加すると推定しています。エージェント ワークロードのオーケストレーション要求により、この提案された 4 倍の増加が促進されます。
2 番目の要因: 小型でローカライズされたモデル
別個

[切り捨てられた]

## Original Extract

Why agentic AI is driving the shift back to CPU inference.

Menu Explore Red Hat
AI
Overview
Hybrid cloud
Platform solutions
Artificial intelligence Build, deploy, and monitor AI models and apps.
Linux standardization Get consistency across operating environments.
Application development Simplify the way you build, deploy, and manage apps.
Automation Scale automation and unite tech, teams, and environments.
Virtualization Modernize operations for virtualized and containerized workloads.
Digital sovereignty Control and protect critical infrastructure.
Security Code, build, deploy, and monitor security-focused software.
Edge computing Deploy workloads closer to the source with edge technology.
Learn how to use our cloud products and solutions at your own pace in the Red Hat® Hybrid Cloud Console.
Red Hat AI icon artificial intelligence, Red Hat Enterprise Linux AI, Red Hat OpenShift AI, RHEL AI, machine learning 38 38 2025-03-12T19:43:40.963Z image/svg+xml Red Hat AI icon artificial intelligence, Red Hat Enterprise Linux AI, Red Hat OpenShift AI, RHEL AI, machine learning Icon no 2025-03-12T19:39:59.817Z Technology icon Standard Red Hat AI Develop and deploy AI solutions across the hybrid cloud.
Red Hat Enterprise Linux icon RHEL, Linux platforms, CentOS 2024-03-01T15:26:42.958Z pending TRA3b65dd25-844d-49bb-93c1-30f5b34684f1 Icon 2024-03-01T15:26:42.958Z true pending 2024-03-21T00:40:29.326Z rhcc-audience:internal no Technology icon DER3b65dd25-844d-49bb-93c1-30f5b34684f1 Standard yes rhcc-product:red-hat-enterprise-linux Technology icon image/svg+xml 2024-05-10T14:11:29.114Z Red Hat Enterprise Linux icon RHEL, Linux platforms, CentOS Activate Activate 2024-05-10T14:11:29.836Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:11:29.836Z workflow-process-service 2024-05-10T14:11:29.836Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat Enterprise Linux Support hybrid cloud innovation on a flexible operating system.
Red Hat OpenShift icon Cloud, Containers, Kubernetes 2024-03-01T15:26:53.684Z pending TRA9ec76aa9-ef09-4c49-8816-01dd13970ca7 Icon 2024-03-01T15:26:53.684Z true pending 2024-03-21T00:39:44.126Z rhcc-audience:internal no Technology icon DER9ec76aa9-ef09-4c49-8816-01dd13970ca7 Standard yes rhcc-product:red-hat-openshift rhcc-product:red-hat-openshift-on-ibm-cloud rhcc-product:microsoft-azure-red-hat-openshift rhcc-product:red-hat-openshift-service-on-aws rhcc-product:red-hat-openshift-container-platform rhcc-product:red-hat-openshift-platform-plus Technology icon image/svg+xml 2024-05-10T14:18:23.703Z Red Hat OpenShift icon Cloud, Containers, Kubernetes Activate Activate 2024-05-10T14:18:25.221Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:18:25.221Z workflow-process-service 2024-05-10T14:18:25.221Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat OpenShift Build, modernize, and deploy apps at scale.
Red Hat Ansible Automation Platform icon Management, edge 2024-03-01T15:26:35.068Z pending TRA759b57c4-760b-45a0-a939-821f47181964 Icon 2024-03-01T15:26:35.068Z true pending 2024-03-21T00:39:55.923Z rhcc-audience:internal no Technology icon DER759b57c4-760b-45a0-a939-821f47181964 Standard yes rhcc-product:red-hat-ansible-automation-platform Technology icon image/svg+xml 2024-05-10T14:04:00.014Z Red Hat Ansible Automation Platform icon Management, edge Activate Activate 2024-05-10T14:04:01.784Z workflow-process-service Activate workflow-process-service false 2024-05-10T14:04:01.784Z workflow-process-service 2024-05-10T14:04:01.784Z Use technology icons to represent Red Hat products and components. Do not remove the icon from the bounding shape. Red Hat Ansible Automation Platform Implement enterprise-wide automation.
Red Hat OpenShift Virtualization Engine
Integrate with major cloud providers
Training
Training & certification
Red Hat Certified System Administrator exam
Red Hat System Administration I
Red Hat Learning Subscription trial (No cost)
Red Hat Certified Engineer exam
Red Hat Certified OpenShift Administrator exam
Discover resources and tools to help you build, deliver, and manage cloud-native applications and services.
Build solutions powered by trusted partners
Find solutions from our collaborative community of experts and technologies in the Red Hat® Ecosystem Catalog.
We'll recommend resources you may like as you browse. Try these suggestions for now.
Log in Get more with a Red Hat account
A subscription may be required for some services.
Account number: [[account_number]]
Red Hat Ansible Automation Platform
Podcasts
Technically Speaking with Chris Wright
More blogs
Red Hat Developer blog
The CPU is back: Rethinking the CPU-GPU split for LLM inference
For the past 3 years, graphics processing units (GPUs) have dominated the large language model (LLM) conversation. In traditional chatbot applications, central processing units (CPUs) provide a fraction of the total compute per request, while GPUs do the heavy lifting. However, inference isn't a single model answering a single question. A growing reliance on tool calls, multistep reasoning, and orchestration across small, specialized models changes the math on where compute should live. Intel has called out this shift noting that the CPU-to-GPU ratio is moving from 1:8 in training workloads to 1:1, and in some cases 4:1 in agentic deployments.
Here we'll examine why the assumptions making GPUs the obvious choice for LLM inference are being renegotiated, what's driving renewed demand for CPU-based serving, and what the data says about where the industry is heading.
At their core (no pun intended), CPUs and GPUs solve fundamentally different problems.
A modern GPU contains tens of thousands of cores designed to execute the same operation on thousands of data elements simultaneously. This makes them extraordinarily fast at the dense matrix multiplications that dominate the forward pass of a transformer for inference. During training, and during high-concurrency batched inference, that parallelism translates directly into throughput: greater tokens per second (TPS), and many more requests served per dollar of compute.
Modern CPUs, by contrast, have anywhere from one to hundreds of cores optimized for sequential, conditional, and branching logic. They're particularly fast at a single operation moving through a complex decision tree. They have direct access to main system memory, and they're the natural execution environment for the orchestration layer wrapping any model. Tool dispatch, code execution, Python runtimes, sandboxes, input/output (I/O), and the agent loop control flow all sit in the CPU.
These architectures, while different in nature, aren't competitors; they work best in tandem. The real question isn't which architecture is better, but which workload belongs where.
This division of labor comes down to how we measure their work: FLOPS versus instruction latency.
GPUs live and die by floating-point operations per second (FLOPS). Because AI models are massive webs of decimal numbers being multiplied and added, a GPU's job is to brute-force trillions of these matrix calculations simultaneously. It's built entirely for raw mathematical throughput.
CPUs, by contrast, specialize in instruction latency. Instruction latency measures how fast a single core can execute an unpredictable chain of diverse commands. A CPU core excels at the rapid logic shuffling needed to parse JavaScript Object Notation (JSON), handle network I/O, or check security permissions.
If you force a GPU to run a chaotic Python runtime, its massive FLOP capacity sits idle, choked by constant task-switching. If you force a CPU to crunch an LLM's math, it works perfectly, but takes ages because it lacks the massive parallel pipelines. The GPU is the mathematical muscle; the CPU is the logic engine directing the show.
The traditional inference stack: Where CPUs were passengers
When we think of the traditional inference stack, we think about the chatbot serving application. Here, the CPU often plays a supporting role.
A request arrives at the API server, where the CPU tokenizes and schedules it. The CPU handles all these initial tasks. It then hands the request to the GPU for the forward pass, which dominates the compute budget. The GPU runs attention, feeds forward layers, and samples the next token with a brief CPU sync to retrieve it and update the scheduler. It then repeats until an end-of-sequence token is emitted. Finally, the CPU collects the output and returns it to end the session.
In this model, the CPU acts as a receptionist, doing the coordination and organizational work, while the GPU does the computational heavy lifting. The ratio of CPUs to GPUs in AI data centers during the training era reflected this: roughly 1 to 2 CPUs for every 8 GPUs, a ratio dictated more by GPUs throughput than by CPU demand. Teams provisioned CPUs to keep GPUs fed. This setup is slowly shifting away from the de facto standard, as advancements in CPU inference increase global efficiency.
The first driver: How agents are changing the workload
With the rise of agentic AI comes a new inference profile magnifying the need for CPU-specific compute.
"Agentic AI" conjures up images of integrated agent harnesses, like Hermes and Openclaw, and developers with hosts of agents tied into their coding assistants. While these use cases are becoming more popular, agentic AI isn't as niche as it may seem at first glance. In fact, many people aren't using the "traditional" process of querying a chatbot. Popular AI assistants like Claude and ChatGPT are backed by enormous models that often incorporate "reasoning" stages into their more complex responses. During reasoning, the "chatbot" becomes a full agentic system. A single user-visible task may decompose into dozens of individual model calls, each of which is short, context-dependent, and conditionally branching.
In an agentic system, the model doesn't issue a single prompt and wait for a response. It generates a plan of action, executes tool calls, compiles all the returned data, and culminates in completing an action or returning an answer to the user. The CPU is in charge of parsing that output, figuring out which tool to invoke, making the API calls or running the code, collecting the result, and feeding it back. Then the loop repeats—sometimes dozens of times per user request, with parallel sub-agents adding more CPU coordination work on top. The model quickly becomes one component of a larger reasoning loop. The ratio of organizational work to pure computation shifts, making the CPU a much bigger player, and potential bottleneck, in the full process.
In a collaboration between researchers at Georgia Tech and Intel , researchers found that in agentic workloads, CPU-side tool processing accounts for 50–90% of total end-to-end latency. Intel's CEO Lip-Bu Tan, at Computex 2026, framed the shift precisely: "for reinforcement learning, orchestration, and agents, the CPU is a much better fit."
The numbers from Intel's Q1 2026 earnings call quantify how far this has already moved:
Training workloads : ~1 CPU per 8 GPUs
Inference workloads : Already shifted to ~1 CPU per 4 GPUs
Agentic workloads : Converging toward 1:1, with some customers reporting 4 CPUs deployed per GPU
Intel's Q1 2026 Data Center and AI segment revenue came in at $5.1B, up 22% year-over-year, with demand running ahead of supply. Intel has deprioritized consumer chip production to redirect fab capacity to server Xeon parts.
Arm's own analysis forecasts the shift more starkly. Traditional AI data centers require approximately 30 million CPU cores per gigawatt (GW) of capacity. Arm CEO Rene Haas estimates in the AI agent era, that figure rises to 120 million CPU cores per GW. The orchestration demands of agentic workloads drive this proposed 4× increase.
The second driver: Smaller, localized models
Separate

[truncated]

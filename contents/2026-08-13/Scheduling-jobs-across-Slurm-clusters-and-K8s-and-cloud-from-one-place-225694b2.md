---
source: "https://skypilot.ai/blog/multi-slurm"
hn_url: "https://news.ycombinator.com/item?id=49280264"
title: "Scheduling jobs across Slurm clusters (and K8s, and cloud) from one place"
article_title: "Managing Multiple Slurm Clusters with SkyPilot | SkyPilot"
author: "alex000kim"
captured_at: "2026-08-13T01:05:58Z"
capture_tool: "hn-digest"
hn_id: 49280264
score: 2
comments: 0
posted_at: "2026-08-13T00:12:56Z"
tags:
  - hacker-news
  - translated
---

# Scheduling jobs across Slurm clusters (and K8s, and cloud) from one place

- HN: [49280264](https://news.ycombinator.com/item?id=49280264)
- Source: [skypilot.ai](https://skypilot.ai/blog/multi-slurm)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T00:12:56Z

## Translation

タイトル: 1 か所から Slurm クラスター (および K8、クラウド) 全体でジョブをスケジュールする
記事のタイトル: SkyPilot を使用した複数の Slurm クラスターの管理 |スカイパイロット
説明: 複数の Slurm クラスターにアクセスできる場合、それらを手動で管理するのは面倒です。 SkyPilot では、それらを単一の統合リソース プールとして使用できます。

記事本文:
SkyPilot を使用した複数の Slurm クラスターの管理 | SkyPilot ドキュメント オープンソース SkyPilot プラットフォーム
10.5k デモを予約する ドキュメント オープンソース SkyPilot プラットフォーム
SkyPilot を使用した複数の Slurm クラスターの管理
複数の Slurm クラスターにアクセスできる場合、それらを手動で管理するのは面倒です。 SkyPilot では、それらを単一の統合リソース プールとして使用できます。
マルチクラスタの問題点 リソースの統一されたビューがない
クラスターがいっぱいの場合の手動フェイルオーバー
クラスタ間で一貫性のない環境
可観測性が悪い = 使用率が低い
ネイティブ Slurm マルチクラスター サポートについてはどうですか?
SkyPilot の Slurm 用統合インターフェース
主な機能 統合リソースビュー
自動クラスター選択とフェイルオーバー
クラスタ間で一貫した環境
クラスタ間でのインタラクティブな開発
クラスタ全体でのジョブ管理
現実世界のワークフロー トレーニング タスクを定義する
高度な機能 クラスター設定の指定
長期トレーニング用に管理されたジョブ
重要な注意事項と制限事項
ストレージに関する考慮事項 共有ファイルシステム
Kubernetes オーバーフローで NFS を利用可能な状態に保つ
Slurm は大規模コンピューティングを支配しており、トップ 500 のスーパーコンピューターの 60% がワークロード管理に Slurm に依存しています。 ML 研究や HPC に取り組んでいる場合、複数の GPU クラスターにアクセスできることは良い問題のように思えます。より多くのリソース = より多くのコンピューティング、つまり実験の高速化と反復サイクルの短縮を意味します。
ただし、実際には、複数の Slurm クラスターの管理は決して楽しいものではありません。手作業であり、エラーが発生しやすいプロセスであり、チームが拡大するにつれてボトルネックになります。
これを想像してください。プライマリ クラスタがいっぱいになったので、別のログイン ノードに SSH 接続し、そこでどのパーティション名が異なっているかを覚えて、ジョブ スクリプトを調整して、再送信します。 3、4、または 5 つの異なるクラスターで洗浄と洗浄を繰り返します。あなたはこう考え始めます。

もっと良い方法があるはずだ」
すべての Slurm クラスターを 1 つのリソース プールとして扱うことができたらどうなるでしょうか?利用可能な GPU があるクラスターをジョブが自動的に見つけて実行できたらどうなるでしょうか? SkyPilot がまさにそれを可能にします。
マルチクラスターの問題点 #
では、複数の Slurm クラスターの管理が難しいのはなぜでしょうか?
各 Slurm クラスターは自己完結型です。クラスター全体でどの GPU が利用できるかを確認するには、以下を行う必要があります。
頭の中で（またはスプレッドシートで）比較してください。
これは、特定の GPU タイプ (H200、H100、B300 など) を探している場合や、マルチノードの割り当てが必要な場合に特に時間がかかります。 「どこでも利用できるものはこれだ」ということを示す単一のコマンドはありません。
クラスターがいっぱいの場合の手動フェイルオーバー
プライマリ クラスターが容量に達した場合、次のことを行う必要があります。
ジョブがキューでスタックしていることに注意してください
別のクラスターを試してみるべきであることに気づきました
送信スクリプトを調整します (異なるパーティション名、異なるモジュール環境、異なるパス)。
クラスターに容量があることを願っています
継続的なトレーニング パイプラインまたはパラメーター スイープを実行しているチームの場合、この手動介入によりワークフローが中断されます。単に 100 件の求人を送信して、彼らにキャパシティを見つけさせることはできません。そのプロセスを子守する必要があります。
クラスタ間で一貫性のない環境
さまざまな Slurm クラスターには次のような特徴があることがよくあります。
異なるパーティション命名スキーム (gpu、b300、h100-cluster)
異なるモジュール環境 (モジュール ロード cuda/12.1 とモジュール ロード cuda/12.4)
さまざまなストレージ マウント (ホーム ディレクトリ パス、スクラッチ ファイルシステムの場所)
さまざまなネットワーク構成
ジョブを別のクラスターに移動するたびに、環境の違いを効果的にデバッグすることになります。これは面倒なだけでなく、環境の不一致によりジョブが失敗した場合に微妙なバグの原因になります。
可観測性が悪い

y = 使用率が低い
クラスター全体の可視性がなければ、賢明なスケジューリングの決定を下すことはできません。
クラスター a で H100 を待つべきですか、それともクラスター b で使用可能な B300 を使用するべきですか?
優先順位の低いジョブをプリエンプトするか、それとも別のクラスターにフェイルオーバーする必要がありますか?
どのクラスターがキュー時間を最速にしますか?
不完全な情報に基づいてインフラを十分に活用できないことになります。
ネイティブ Slurm マルチクラスター サポートについてはどうですか?
Slurm はマルチクラスター操作とフェデレーションを提供しますが、どちらにも制限があります。マルチクラスター モードでは、送信時に最も早い開始時刻を持つクラスターにジョブがルーティングされますが、それ以降、「Slurm はジョブを別のクラスターに移行するための努力を一切行いません」。フェデレーションは、複製された「兄弟ジョブ」によるピアツーピアのスケジューリングを提供しますが、ドキュメントでは、それが「高スループット環境を目的としたものではない」と述べており、1 日に 50,000 を超えるジョブをスケジュールする場合は、構成するクラスターの数を少なくすることを推奨しています。また、すべてのコンピューティング ノードがすべてのサブミッション ホストからアクセス可能であることも必要ですが、これは個別のデータ センターにあるクラスターにとっては困難です。
SkyPilot の Slurm 用統合インターフェース #
SkyPilot は、断片化されたコンピューティングを 1 つの統合された AI コンピューティング プールに変えるコントロール プレーンです。 SkyPilot は (単一または複数の) Kubernetes を適切にサポートしていますが、複数の Slurm クラスターを統合リソース プールとして管理する機能など、Slurm クラスターに対する最高クラスのサポートも備えています。
統合リソース ビュー: 単一のコマンドですべてのクラスターの GPU を表示します。可用性を確認するために別のログイン ノードに SSH 接続する必要はありません。
自動クラスター選択とフェイルオーバー : ジョブを送信すると、SkyPilot が容量のあるクラスターを見つけます。そのクラスターがいっぱいの場合は、次のクラスターが自動的に試行されます。
一貫した環境 : タスク定義を一度だけ記述します。同じヤム

L は、任意の Slurm クラスター (クラウド VM や Kubernetes 上でも) で動作します。
統合ジョブ管理 : すべてのクラスターにわたるジョブを 1 か所から監視、記録、キャンセルします。
クラウド オーバーフロー : すべての Slurm クラスターが飽和状態になると、容量を追加するために自動的にクラウド VM にバーストします。
この投稿の残りの部分では、セットアップ、使用法、実際のワークフローについて説明します。
Slurm クラスターへの SSH アクセスが必要です。 SkyPilot は SSH を使用してログイン ノードに接続し、手動で行うのと同じように sbatch 経由でジョブを送信します。
SSH 構成形式を使用して ~/.slurm/config でクラスターを構成します。
これは標準の SSH 構成形式に従っているため、これらのクラスターに SSH 構成をすでに使用している場合は、それらを再利用できます。
SkyPilot がすべてのクラスターを認識できることを確認します。
特定のクラスターへの接続をテストすることもできます。
SkyPilot は、構成から利用可能なクラスターを自動的に検出し、接続をテストします。
SkyPilot ダッシュボードは、すべての GPU クラスターの統合ビューを提供します。
CLI を使用したい場合は、1 つのコマンドですべてのクラスターの GPU を確認できます。
これにより、何がどこで利用可能であるかを即座に把握できます。キューのステータスを確認するために別のログイン ノードに SSH 接続する必要はもうありません。
自動クラスター選択とフェイルオーバー
注: GPU 名 (例: H100 、 B300 、 L4 ) は、Slurm クラスターの GRES 設定で構成されたものと一致する必要があります。 SkyPilot は、Slurm に送信するときに、これを適切な --gres=gpu:H100:8 ディレクティブに変換します。
構成されているすべての Slurm クラスター全体でリソースの可用性を確認します
使用可能な H100 があるクラスターを選択します
最初のクラスターがいっぱいの場合は、自動的に次のクラスターを試行します
クラスタ全体で環境セットアップを一貫して処理します
使用するクラスターを指定する必要はありません。SkyPilot がそこで容量とスケジュールを見つけます。
クラスタ間で一貫した環境
SkyPilot の要約

離れたクラスター固有の詳細:
パーティション名 : パーティションの指定はオプションです。 SkyPilot は、GPU リクエストを各クラスターの適切なパーティションに自動的にマッピングできます
モジュール環境: セットアップ ブロックはすべてのクラスターで実行されるため、依存関係は一貫しています。
環境変数 : SkyPilot はどこでも機能する標準化された変数を提供します
タスク定義は移植可能です。同じ YAML は、どの Slurm クラスターでも (必要に応じてクラウド VM や Kubernetes でも​​) 機能します。
SkyPilot は、マルチノード ジョブに対して次の環境変数を自動的に設定します。
解決された、改行で区切られたノード IP (SLURM_JOB_NODELIST にはノード名が含まれます)
これらの変数により、コードの移植性が高まり、同じスクリプトが Slurm、クラウド VM、および Kubernetes で動作します。
クラスタ間でのインタラクティブな開発
任意のクラスターに到達できる salloc スタイルの対話型セッションを使用します。
SkyPilot は、すべてのクラスターで利用可能な H100 を見つけて、開発クラスターを提供します。 SSH 接続:
これで、容量のあるクラスターの GPU ノード上にいることになります。どちらの環境が一貫しているかを知る必要はありません。
SkyPilot 開発クラスターを使用して VSCode リモート開発をセットアップすることもできます (手順はこちらを参照してください)。
クラスタ全体でのジョブ管理
すべてのクラスターにわたるすべてのジョブを表示します。
ジョブを一律にキャンセルまたは停止します。
各ジョブがどのクラスター上にあるかを覚えたり、ジョブをキャンセルするために別のログイン ノードに SSH 接続したりする必要はありません。
Slurm コマンドを SkyPilot に変換するためのクイック リファレンスは次のとおりです。
空の起動 -c dev --gpus H100:8
Sky down <cluster> または Sky jobs cancel <jobid>
主な違い: SkyPilot では、使用するクラスターを指定する必要がありません。 SkyPilot は、リソースの可用性に基づいて、利用可能な最適なクラスターを自動的に選択します。
複数のモデル バリアントを並行してトレーニングする、一般的なワークフローを見てみましょう。
SkyPilot が自動配信

これらは可用性に基づいてクラスター全体に適用されます。おそらく、あなたが指定したり気にすることなく、train-7b は slurm-cluster-a に、train-13b は slurm-cluster-b に、train-70b は slurm-cluster-c-all に着陸するでしょう。
正しいログイン ノードへの SSH 接続はありません。SkyPilot が処理します。
クラスター設定の指定
(データの局所性、コストなどの理由で) 特定のクラスターをターゲットにする場合は、それを infra フィールドで指定します。
または、起動時に --infra フラグを使用します。
クラスターを指定すると、SkyPilot はそのクラスターのみを使用します。複数のクラスター間で自動フェイルオーバーを有効にするには、クラスターの指定を省略し、可用性に基づいて SkyPilot に選択させます。
SkyPilot は、Slurm クラスターと並行してクラウド VM をサポートします。すべての Slurm クラスターがいっぱいの場合は、自動的にクラウドにオーバーフローできます。
これにより、オンプレミスのリソースが飽和した場合でも、無限のオーバーフロー容量が得られます。
長期トレーニング用に管理されたジョブ
フォールト トレランスが必要な数日間にわたるトレーニングの実行には、SkyPilot の管理ジョブ機能を使用します。
これはすべての Slurm クラスターで機能し、クラスターがダウンした場合には自動フェイルオーバーが行われます。
重要な注意事項と制限事項 #
SkyPilot は Slurm に強力なマルチクラスター管理を提供しますが、注意すべき制限がいくつかあります。
既存の Slurm クラスターへの SSH ベースのアクセス
SSH アクセスによる対話型開発
フォールトトレランスを備えた管理されたジョブ
クラスタ間の自動フェイルオーバー
Pyxis および enroot 経由のコンテナ イメージ (「コンテナ」を参照)
自動停止 : Slurm クラスターではサポートされていません (クラスターを手動で停止する必要があります)
SkyServe : モデル サービングの展開は Slurm では利用できません
クラスターのプロビジョニング: SkyPilot は新しい Slurm クラスターを作成できません (既存のクラスターのみを管理します)。
認証と権限: ジョブは設定された SSH ユーザー名で送信され、既存の Slurm アカウント権限が尊重されます。あなたのアカウントの場合

には特定のパーティションまたは GPU タイプに制限があり、SkyPilot 経由で起動する場合にも同じ制限が適用されます。チームが共有 SkyPilot API サーバーを実行している場合は、ユーザーとして送信を有効にして、単一の共有ログインではなく、各個人の Unix アカウントにジョブが送信されるようにすることができます。
共有ファイルシステム: SkyPilot は、Slurm クラスターにマウントされている共有ファイルシステム (通常は NFS) を自動的に利用します。ホーム ディレクトリと共有スクラッチ スペースには、追加の構成を行わなくても、すべての計算ノードからアクセスできます。
クラスターに共有ストレージ (同じパスにマウントされた共有 NFS または Lustre ファイルシステムなど) がある場合は、どのクラスターからでもデータにアクセスできます。
各クラスターに独自のストレージがある場合は、SkyPilot のファイル マウントを使用してデータを同期します。
SkyPilot は、ジョブが実行されるクラスターに S3 バケットをマウントするため、データは常に利用可能です。
Kubernetes オーバーフローで NFS を利用可能な状態に保つ
Slurm では、SkyPilot は既存の NFS マウントを自動的に使用します。 Kubernetes は異なります。Slurm とは異なり、NFS ホーム ディレクトリをマウントしません。そのため、Slurm 上の /shared から読み取るジョブは、Kubernetes クラスターにバーストした後にそのディレクトリを見つけることができません。 (クラウド オーバーフローの場合は、代わりにオブジェクト ストレージを使用します。上記のクラスターごとのストレージを参照してください。)
Kubernetes クラスターが NFS サーバーに到達できるかどうか

[切り捨てられた]

## Original Extract

If you have access to multiple Slurm clusters, managing them manually is painful. SkyPilot lets you use them as a single unified resource pool.

Managing Multiple Slurm Clusters with SkyPilot | SkyPilot Docs Open Source SkyPilot Platform
10.5k Book a demo Docs Open Source SkyPilot Platform
Managing Multiple Slurm Clusters with SkyPilot
If you have access to multiple Slurm clusters, managing them manually is painful. SkyPilot lets you use them as a single unified resource pool.
The multi-cluster pain points No unified view of resources
Manual failover when clusters are full
Inconsistent environments across clusters
Bad observability = low utilization
What about native Slurm multi-cluster support?
SkyPilot’s unified interface for Slurm
Key features Unified resource view
Automatic cluster selection and failover
Consistent environment across clusters
Interactive development across clusters
Job management across clusters
Real-world workflow Define your training task
Advanced features Specifying cluster preferences
Managed jobs for long-running training
Important notes and limitations
Storage considerations Shared filesystems
Keeping NFS available on Kubernetes overflow
Slurm dominates large-scale computing - 60% of Top500 supercomputers rely on it for workload management. If you’re working in ML research or HPC, having access to multiple GPU clusters sounds like a good problem to have. More resources = more compute, which means faster experiments and shorter iteration cycles.
In practice, however, managing multiple Slurm clusters is anything but enjoyable, it’s a manual, error-prone process that becomes a bottleneck as your team scales.
Imagine this: your primary cluster is full, so you SSH to a different login node, remember which partition names are different there, tweak your job script, and resubmit. You rinse and repeat across three, four, or five different clusters. You start thinking “there’s gotta be a better way.”
What if you could treat all your Slurm clusters as a single resource pool? What if your jobs could automatically find and run on whichever cluster has available GPUs? That’s exactly what SkyPilot enables.
The multi-cluster pain points #
So what makes managing multiple Slurm clusters difficult?
Each Slurm cluster is self-contained. To see what GPUs are available across your clusters, you need to:
Compare in your head (or in a spreadsheet)
This becomes especially time consuming when you’re looking for specific GPU types (e.g. H200s vs H100 vs B300s) or need multi-node allocations. There’s no single command that shows you “here’s what’s available everywhere.”
Manual failover when clusters are full
When your primary cluster hits capacity, you have to:
Notice that your job is stuck in the queue
Realize you should try another cluster
Adapt your submission script (different partition names, different module environments, different paths)
Hope that cluster has capacity
For teams running continuous training pipelines or parameter sweeps, this manual intervention breaks the workflow. You can’t just submit 100 jobs and let them find capacity-you have to babysit the process.
Inconsistent environments across clusters
Different Slurm clusters often have:
Different partition naming schemes ( gpu vs b300 vs h100-cluster )
Different module environments ( module load cuda/12.1 vs module load cuda/12.4 )
Different storage mounts (home directory paths, scratch filesystem locations)
Different networking configurations
Every time you move a job to a different cluster, you’re effectively debugging environment differences. This is both tedious and it becomes a source of subtle bugs when jobs fail due to environment mismatches.
Bad observability = low utilization
Without visibility across clusters, you can’t make smart scheduling decisions:
Should you wait for H100s on cluster-a, or use available B300s on cluster-b?
Should you preempt lower-priority jobs, or fail over to another cluster?
Which cluster will give you the fastest queue time?
You’ll end up underutilizing your infra based on incomplete information.
What about native Slurm multi-cluster support?
Slurm does offer multi-cluster operation and federation, but both have limitations. Multi-cluster mode routes jobs to the cluster with the earliest start time at submission, but “Slurm makes no subsequent effort to migrate a job to a different cluster” after that. Federation provides peer-to-peer scheduling with replicated “sibling jobs,” but the docs note it’s “not intended as a high-throughput environment” and suggest configuring fewer clusters if you’re scheduling more than 50,000 jobs a day. It also requires that all compute nodes be reachable from all submission hosts, which is a tall order for clusters in separate data centers.
SkyPilot’s unified interface for Slurm #
SkyPilot is a control plane that turns fragmented compute into one unified AI compute pool. While SkyPilot supports (single or multiple) Kubernetes really well, it also has first-class support for Slurm clusters, including the ability to manage multiple Slurm clusters as a unified resource pool.
Unified resource view : See GPUs across all your clusters with a single command-no more SSH-ing to different login nodes to check availability.
Automatic cluster selection and failover : Submit a job and SkyPilot finds a cluster with capacity. If that cluster is full, it tries the next one automatically.
Consistent environments : Write your task definition once. The same YAML works on any Slurm cluster (and even on cloud VMs or Kubernetes).
Unified job management : Monitor, log, and cancel jobs across all clusters from one place.
Cloud overflow : When all your Slurm clusters are saturated, automatically burst to cloud VMs for additional capacity.
The rest of this post walks through setup, usage, and real-world workflows.
You need SSH access to your Slurm clusters. SkyPilot uses SSH to connect to login nodes and submit jobs via sbatch , just like you would manually.
Configure your clusters in ~/.slurm/config using SSH config format:
This follows the standard SSH config format, so if you’re already using SSH configs for these clusters, you can reuse them.
Check that SkyPilot can see all your clusters:
You can also test connectivity to a specific cluster:
SkyPilot automatically discovers available clusters from your config and tests connectivity.
The SkyPilot dashboard provides a unified view of all your GPU clusters:
If those who prefer the CLI, see GPUs across all your clusters with one command:
This gives you instant visibility into what’s available where. No more SSH-ing to different login nodes to check queue status.
Automatic cluster selection and failover
Note: The GPU name (e.g., H100 , B300 , L4 ) must match what’s configured in your Slurm cluster’s GRES settings. SkyPilot converts this to the appropriate --gres=gpu:H100:8 directive when submitting to Slurm.
Check resource availability across all configured Slurm clusters
Select the cluster with available H100s
If the first cluster is full, automatically try the next one
Handle environment setup consistently across clusters
You don’t have to specify which cluster to use - SkyPilot finds capacity and schedules there.
Consistent environment across clusters
SkyPilot abstracts away cluster-specific details:
Partition names : Specifying partitions is optional; SkyPilot can automatically map GPU requests to the right partitions on each cluster
Module environments : Your setup block runs on every cluster, so dependencies are consistent
Environment variables : SkyPilot provides standardized variables that work everywhere
Your task definition is portable. The same YAML works on any Slurm cluster (and even on cloud VMs or Kubernetes if needed).
SkyPilot sets these environment variables automatically for multi-node jobs:
Resolved, newline-separated node IPs ( SLURM_JOB_NODELIST contains node names)
These variables make your code portable-the same script works on Slurm, cloud VMs, and Kubernetes.
Interactive development across clusters
Use salloc -style interactive sessions that can land on any cluster:
SkyPilot finds an available H100 across all your clusters and gives you a dev cluster. SSH in:
You’re now on a GPU node on whichever cluster had capacity. You don’t need to know which one-the environment is consistent.
You can even set up VSCode remote development with SkyPilot dev clusters (see instructions here ).
Job management across clusters
See all your jobs across all clusters:
Cancel or stop jobs uniformly:
No need to remember which cluster each job is on or SSH to different login nodes to cancel jobs.
Here’s a quick reference for translating Slurm commands to SkyPilot:
sky launch -c dev --gpus H100:8
sky down <cluster> or sky jobs cancel <jobid>
The key difference: with SkyPilot, you don’t need to specify which cluster to use. SkyPilot automatically selects the best available cluster based on resource availability.
Let’s walk through a typical workflow: training multiple model variants in parallel.
SkyPilot automatically distributes these across your clusters based on availability. Maybe train-7b lands on slurm-cluster-a, train-13b on slurm-cluster-b, and train-70b on slurm-cluster-c-all without you specifying or caring.
No SSH-ing to the right login node-SkyPilot handles it.
Specifying cluster preferences
If you want to target a specific cluster (due to data locality, cost, etc.), specify it in the infra field:
Or use the --infra flag at launch time:
When you specify a cluster, SkyPilot will only use that cluster. To enable automatic failover across multiple clusters, omit the cluster specification and let SkyPilot choose based on availability.
SkyPilot supports cloud VMs alongside Slurm clusters. If all your Slurm clusters are full, you can automatically overflow to cloud:
This gives you infinite overflow capacity when on-prem resources are saturated.
Managed jobs for long-running training
Use SkyPilot’s managed jobs feature for multi-day training runs that need fault tolerance:
This works across all your Slurm clusters, with automatic failover if a cluster goes down.
Important notes and limitations #
While SkyPilot provides powerful multi-cluster management for Slurm, there are some limitations to be aware of:
SSH-based access to existing Slurm clusters
Interactive development with SSH access
Managed jobs with fault tolerance
Automatic failover across clusters
Container images via Pyxis and enroot (see Containers )
Autostop : Not supported on Slurm clusters (you need to manually stop clusters)
SkyServe : Model serving deployments are not available on Slurm
Cluster provisioning : SkyPilot cannot create new Slurm clusters (only manages existing ones)
Authentication & permissions: Jobs submit under your configured SSH username and respect your existing Slurm account permissions. If your account has restrictions on certain partitions or GPU types, those same restrictions apply when launching through SkyPilot. If your team runs a shared SkyPilot API server, you can enable submit as user so jobs land under each person’s own Unix account rather than a single shared login.
Shared filesystems: SkyPilot automatically leverages shared filesystems (typically NFS) that are mounted on your Slurm clusters. Your home directory and any shared scratch spaces are accessible across all compute nodes without additional configuration.
If your clusters have shared storage (e.g., a shared NFS or Lustre filesystem mounted at the same path), you can access data from any cluster:
If each cluster has its own storage, use SkyPilot’s file mounts to sync data:
SkyPilot will mount the S3 bucket on whichever cluster your job lands on, so data is always available.
Keeping NFS available on Kubernetes overflow
On Slurm, SkyPilot uses your existing NFS mounts automatically. Kubernetes is different - it doesn’t mount your NFS home directory the way Slurm does, so a job that reads from /shared on Slurm won’t find it after it bursts to a Kubernetes cluster. (For cloud overflow, reach for object storage instead - see Per-cluster storage above.)
If your Kubernetes cluster can reach the NFS server

[truncated]

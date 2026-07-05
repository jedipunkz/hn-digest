---
source: "https://mitos.run/blog/ai-sandboxes-on-kubernetes"
hn_url: "https://news.ycombinator.com/item?id=48794737"
title: "Secure AI Sandboxes on Kubernetes"
article_title: "Secure AI sandboxes on Kubernetes · Mitos"
author: "stubbi"
captured_at: "2026-07-05T15:06:12Z"
capture_tool: "hn-digest"
hn_id: 48794737
score: 1
comments: 1
posted_at: "2026-07-05T14:54:23Z"
tags:
  - hacker-news
  - translated
---

# Secure AI Sandboxes on Kubernetes

- HN: [48794737](https://news.ycombinator.com/item?id=48794737)
- Source: [mitos.run](https://mitos.run/blog/ai-sandboxes-on-kubernetes)
- Score: 1
- Comments: 1
- Posted: 2026-07-05T14:54:23Z

## Translation

タイトル: Kubernetes 上の安全な AI サンドボックス
記事のタイトル: Kubernetes 上のセキュアな AI サンドボックス · Mitos
説明: Kubernetes で安全な AI サンドボックスを実行する方法: ポッドでは不十分な理由、gVisor 対 Kata 対 Firecracker microVM、および実行中のマシンのフォーク。

記事本文:
Kubernetes で AI サンドボックスを保護する · Mitos Mitos 仕組み ベンチマーク 価格ドキュメント ★ 22 始めましょう 仕組み ベンチマーク 価格ドキュメント ★ 22 ← すべてのメモ [エンジニアリング] Kubernetes 上の AI サンドボックスを保護する
ミトス チーム作成 · 2026 年 7 月 4 日 · 18 分で読む
あなたのエージェントは、自分自身のために書いたコードに対して pip install を実行しました。そのコードは誰も読まないので、どこかで実行する必要があります。つまり、実際のマシン上で、実際に分離された状態で実行する必要があります。そしてエージェントは、バッチ コンピューティングにはなかった要件をもたらしました。エージェントはタスクの途中でサブエージェントを生成し、結果をマージして戻すため、作業が分岐する場所でマシン自体も分岐する必要があります。
私たちは意図的に Kubernetes 上に Mitos を構築します。私たちは何年も Kubernetes の運用に費やしてきましたが、これはクラウドで大規模にワークロードを実行するための事実上の標準です。サンドボックス フリートに必要なスケジューリング、クォータ、ネットワーク ポリシーは、10 年間の運用によって強化されたもので、すでにそこに存在しています。私たちの賭けはシンプルです。AI のスケーリングはサンドボックスのスケーリングを意味し、サンドボックスは、信頼することを学ばなければならない並列スタックではなく、すでにスケーリングされているインフラストラクチャ上で実行されるべきです。問題は、Kubernetes の作業単位であるポッドが信頼できるコード用に設計されていることです。
これは、Kubernetes 上の AI サンドボックスを保護するための有効な答えです。ポッドが実際に提供するもの、モデルで記述されたコードに対してどのランタイムが保持するのか、実行中のマシンのフォークにかかるコスト、microVM を通常のポッドのように動作させる方法などです。最終的には、独自の脅威モデル用のランタイムを選択し、1 台のウォーム マシンを 30 台に拡張できるようになるはずです。
モデルで記述されたコードは信頼できないコードです
LLM はユーザーの指示と攻撃者の指示を区別できません。サイモン・ウィリソン氏の致命的な 3 つの要因は、個人データ、信頼できないコンテンツ、そして声を上げる方法という失敗モードに名前を付けています。 3 つを組み合わせると、挿入された命令によってエージェントが実行できるものはすべて盗み出されます。

読みます。エージェントにシェルを与えると、インジェクションがコード実行になります。
これを信じて受け入れる必要はありません。 Trail of Bits は 2025 年 10 月に、サンドボックスを使用しないコマンドのホワイトリスト登録が失敗することを示しました。攻撃者は go test -exec や fd -x などの事前承認されたコマンドにフラグを密かし、1 つの有害なプロンプトから実行を取得します。 2025 年 7 月、Replit エージェントはコードの凍結中に実稼働データベースを削除しました。 1 か月後、Nx サプライ チェーン攻撃 ( s1ngularity ) によって方向が反転しました。悪意のある npm パッケージが、インストールされている AI CLI を呼び出し、資格情報を探すために使用しました。 GitGuardian は、1,079 のリポジトリから 2,349 のシークレットを数えました。
ツールベンダーもすでにその点を認めています。 Anthropic は、Claude Code 用の OS レベルのサンドボックスを出荷し、許可プロンプトが 84% 削減されたと報告しています。 OpenAI の Codex は、デフォルトではネットワークがオフの状態で実行されます。
これらは 1 台のラップトップを保護します。他の人のためにエージェントをホストしたり、1 人のエージェントを 50 人に分散させたりすると、敵対的なマルチテナント サービスを実行することになります。問題はサンドボックス化するかどうかではなくなります。それは、どの境界線を貸す意思があるかということです。
ポッドはカーネルのビューです
Kubernetes のデフォルトの答えはエージェントごとにポッドであるため、そこから始めて実際に何が得られるかを見てください。
Kubernetes はコンテナを実行しません。 kubelet は、 Container Runtime Interface を介してランタイム (通常は containerd) を呼び出します。 containerd はポッドごとにシムを開始します。 shim は runc を呼び出し、 runc はコンテナと呼ばれるものを構築します。つまり、制限されたビューの名前空間、制限の cgroups、syscall をトリミングする seccomp が構築され、プロセスに実行されます。
その連鎖の中で何が起こらなかったのかに注目してください。プロセスとカーネルの間に壁を作るものは何もありません。プロセスが認識するものは少なくなりますが、数百のシステムコールのインターフェイスを通じて、ノード上の他のすべてのポッドと同じカーネルと通信します。あん

amespace は壁ではなくビューです。
そのインターフェースには記録があります。 CVE-2019-5736 コンテナーが /proc/self/exe を通じてホスト runc バイナリを上書きできるようにします。 CVE-2024-21626 (Leaky Vessels ファイル記述子リーク) はスコア 8.6 で、ホスト ファイルシステムに逃げました。 2025 年 11 月には、さらに 3 つの runc エスケープが同じ日に着陸し、すべて procfs 書き込みを通じて発生しました。
これらはコンテナのエンジニアリングを悪くするものではありません。あなたが書いてレビューしたコードにとって、強化されたコンテナーは合理的な境界であり、インターネットのほとんどはその上で実行されます。 Kubernetes マルチテナントのドキュメントはここでも正直です。分離はスペクトルであり、テナントがお互いを信頼しなくなると、ドキュメントはサンドボックス化されたランタイムを示します。
エージェント コードは構築によって信頼テストに合格しません。あなたが書いたわけではありません。あなたはそれをレビューしませんでした。そして、攻撃者がそのモデルを操作した可能性があります。
gバイザー、カタ、爆竹：壁を高くする3つの方法
Kubernetes には、まさにこの問題に対応するプラグイン可能なシームがあります。 RuntimeClass は spec.runtimeClassName を別の shim にマップするため、同じポッド仕様が別の分離メカニズムに配置される可能性があります。 3 つのランタイムが重要です。
gVisor は、ワークロードと実際のカーネルの間にユーザー空間カーネルを置きます。 Syscall はインターセプトされ、Go で再実装されます。 gVisor 自身の言葉によれば、設計目標は、システムコールが直接通過しないことです。数十ミリ秒で開始され、GKE Sandbox をサポートします。料金はシステムコールに応じて異なります。ほとんどのワークロードでは数パーセントの料金がかかりますが、システムコールの負荷が高いワークロードでは数パーセントを支払うこともあります。そしてその壁はソフトウェアでできています。
Kata Containers は、KVM の下に独自のゲスト カーネルを備えた実際の VM にすべてのポッドを配置します。より強力な壁、より重い価格: サードパーティのベンチマークによると、追加の開始遅延は約 600 ミリ秒で、デフォルトの VMM ではポッドあたり 180 MiB 程度になります。おおよその数値ですが、形状は正しいです。
それからそこに

Firecracker は、Lambda を実行するために AWS が構築した microVM です。その作者らは、Rust の行数が約 5 万行であり、QEMU よりもコードが 96% 少ないと報告しています。 virtio-net、virtio-block、vsock、シリアル コンソールなど、ほとんど何もエミュレートしません。公開された仕様では、API 呼び出しからゲスト ユーザー空間までが 125 ミリ秒未満、microVM あたりのオーバーヘッドが 5 MiB 未満であることが約束されています。 NSDI の論文では、QEMU が 131 MB を必要とするオーバーヘッドがおよそ 3 MB であることが測定され、ホストごとに 1 秒あたり最大 150 個の microVM の作成速度が報告されています。
ほぼコンテナ価格のハードウェアウォール。そのため、サンドボックス カテゴリの多くがこの上で実行されます。
*サードパーティのベンチマーク数値、おおよその値。爆竹の図は公開された仕様からのものです。 Morph と CodeSandbox も、Kubernetes ではなく独自のクラウド上で実行中の VM をフォークします。
これは、設計ドキュメントに貼り付けることができる 3 つの質問に圧縮された決定事項です。
コードを書いたのは誰ですか?あなたも、そしてあなたもそれをレビューしました。runc は問題ありません。節約した分は別の場所に使いましょう。モデルはこう書きました。「続けてください。」
主にライブラリを呼び出しますか、それともシェルアウトしてパッケージをインストールしますか?ほとんどのライブラリ、一度に 1 つのサンドボックス: gVisor はおそらく保持されます。そうでないふりをするよりも、そのことをお伝えしたいと思います。 OS の完全な動作: 続行します。
多くのエージェントが 1 つのウォーム状態から開始する必要がありますか?いいえ: Kata または任意の microVM ランタイムはここで分析を終了します。はい、フォークが必要です。この投稿の残りの部分では、フォークに何が必要かについて説明します。
オプトインは 1 つのフィールドのように見えます。
apiVersion : node.k8s.io/v1 kind : RuntimeClass メタデータ : name : kata-fc handler : kata-fc #containerd-shim-kata-v2 + Firecracker --- apiVersion : v1 kind : Pod spec : runtimeClassName : kata-fc
マニフェスト内の 1 つのフィールドとすべてのポッドが独自のカーネルを取得します。それがすべての話であれば、この投稿はここで終わります。
Firecracker は Kubernetes が想定しているものをすべて削除しました
E

Firecracker が小型化のために削除されたものそのものが、Kubernetes が密かに依存しているものです。
virtio-fs がないということは、共有ファイルシステムがないことを意味するため、コンテナ rootfs はブロック デバイスとして到着する必要があり、これにより、Kata のドキュメントで Firecracker VMM に必要な devmapper スナップショットにコンテナd が強制的に追加されます。デバイスのホットプラグも VFIO もないことは、Kata-with-Firecracker が実行中のコンテナのサイズを変更したり、GPU を通過したりできないことを意味します。メモリは、柔軟なリクエスト用に構築されたスケジューラに対して事前にサイズ設定されます。ネットワークは VM ごとのタップ デバイスであり、追加の CNI プラグインによってポッドの veth にブリッジされます。また、直接統合である firecracker-containerd は、今後の目標の下に CRI 準拠をリストしています。
アップストリームの Kubernetes は、2025 年後半に、gVisor または Kata の周りにサンドボックス CRD、テンプレート、ウォーム プールをラップする SIG Apps プロジェクトである Agent-sandbox で答えました。 Kubernetes ブログで述べられているように、新しいポッドの起動に約 1 秒かかるため、ウォーム プールが存在します。エージェントが一度に 1 つのサンドボックスを実行し、gVisor が脅威モデルをクリアする場合は、そのパスを選択してください。それは標準であり、よく管理されており、この投稿はそれについて説明するものではありません。
ただし、標準化されている内容を読んでください。エージェントごとに 1 つのサンドボックスがあり、新たに作成されるか、事前にウォームアップされて渡されます。そのスタックには実行中のマシンをフォークするものは何もありません。
エージェントが増えると、その操作を見逃すことになります。その理由は次のとおりです。
コピーオンライトによりコピーのコストが変わります
マシンをコピーする最も安価な方法は、コピーしないことです。 fork() は 40 年間このように動作してきました。子は親の物理ページを読み取り専用で共有し、カーネルは誰かがページを書き込んだ場合にのみページをコピーします。子供には、それが変わるほどのコストがかかります。
Firecracker スナップショットはそれをマシン全体に拡張します。スナップショットは、ゲスト メモリとデバイスの状態を 2 つのファイルにまとめたものです。復元時に、Firecracker はメモリ ファイル MAP_PRIVATE をマップするため、ページがロードされます。

ホスト ページ キャッシュからの要求があり、そのスナップショットから復元されたすべての VM は、同じ物理ページを書き込むまで読み取ります。
演算を 1 回実行すると、最適化されているように感じなくなります。ウォームな 512 MiB Python 環境の 32 個のフォークは、単純に計算すると 16 GiB になります。物理的には、フォーク時に、1 つの 512 MiB ページ セットに加えて、ドーターごとに約 3 MiB のプライベート ページが含まれます (0.7 GiB 未満)。娘たちはページを汚すにつれて成長するので、それは床であり、定常状態ではありません。ただし、マシン 1 台半のメモリをフリートにプロビジョニングしただけです。
これは実験室のトリックではありません。 Lambda SnapStart はスナップショットから機能を再開し、Aurora DSQL は独自のクローン Firecracker VM で各 SQL トランザクションを実行し、変更されていないページをクローン間で共有します。弊社のリファレンス ハードウェアでは、Mitos エンジンは 6 ～ 16 ミリ秒でスナップショットを復元し、フォークあたり約 3 MiB の限界メモリで、約 27 ミリ秒でウォーム フォークをアクティブ化します。エンジンの測定値。リポジトリ内のベンチマーク スクリプトから再現可能。何が含まれていて何が含まれていないのかは、ベンチマーク ページに詳しく記載されています。
1 つの区別があれば、ベンダーのページを読むときに実際に混乱することがなくなります。ディスク コピー オン ライト (overlayfs レイヤー、qcow2 バッキング ファイル、シン ボリューム) はファイル システムのクローンを作成し、多くの製品はゴールデン イメージをスナップショットと呼んでいます。マシンのスナップショットには、ライブ メモリと実行中のプロセスが含まれます。 1つ目は、安価なクリーンブートを購入することです。 2 つ目は、タスク中にマシンの安価なコピーを購入します。これは、エージェントが依存関係のインストールとモデルのロードに 40 秒を費やした後に必要とするマシンです。
スナップショットを 2 回復元すると、セキュリティ バグが発生します
これはサンドボックスの記述では省略される部分であり、フォークが速度トリックである前にセキュリティ メカニズムである理由はここにあります。
Firecracker 自身のドキュメントでは、1 つのゲスト状態が複数回再開されると、次のように警告しています。

一意であると想定されているゲスト情報は、実際にはそうでない可能性があります。」そのメモリ イメージに何が存在するかを考えてください。カーネルのエントロピー プール、ユーザー空間の PRNG 状態、TCP シーケンス番号、マシン ID、キャッシュされた TLS セッション キーです。 20 の単純な復元とは、衝突する UUID、セッション トークン、および nonce を生成できる 20 のゲストを意味します。
すでに敵対的であると想定しているコードにとって、それは弓を付けた贈り物です。
カーネル半分には修正が加えられています。Firecracker は再開時に新しい VMGenID を書き込み、Linux 5.18 以降はそこから CSPRNG を再シードしますが、ドキュメントでは再シードが完了するまでにレース ウィンドウが存在することを認めています。ユーザースペースの半分には一般的な修正がありません。同じ医師ははっきりとそう言います。ネットワーク ID も重複します。すべてのクローンは親の MAC と IP で起動します。
したがって、正しいフォークはプロトコルであり、ファイルのコピーではありません。 Mitos がドーターをアクティブ化すると、新しいエントロピーがゲストにプッシュされ、ゲスト エージェントが RNG 再シードを確認するのを待ち、凍結されたスナップショット時間から壁時計を調整し、NIC のアドレスを再設定し、ドーターごとのシークレットを配信します。フェイルクローズされます。再シードが確認されなかったクローンは決して提供されません。詳細については、フォークの正確性に関するドキュメントを参照してください。
サンドボックス ベンダーに対して 3 つの質問をすることができるようになりました。 「スナップショット」にはメモリと実行中のプロセスが含まれますか、それとも

[切り捨てられた]

## Original Extract

How to run secure AI sandboxes on Kubernetes: why a pod is not enough, gVisor vs Kata vs Firecracker microVMs, and forking running machines.

Secure AI sandboxes on Kubernetes · Mitos Mitos How it works Benchmarks Pricing Docs ★ 22 Get started How it works Benchmarks Pricing Docs ★ 22 ← All notes [engineering] Secure AI sandboxes on Kubernetes
By Mitos team · July 4, 2026 · 18 min read
Your agent just ran pip install on code it wrote for itself. Nobody read that code, and it has to execute somewhere: on a real machine, with real isolation. And agents brought a requirement that batch computing never had: they spawn subagents mid-task and merge the results back, so the machine itself has to fork where the work forks.
We build Mitos on Kubernetes, deliberately. We have spent years operating Kubernetes, and it is the de facto standard for running workloads at scale in the cloud: the scheduling, quotas, and network policy a sandbox fleet needs already exist there, hardened by a decade of production. Our bet is simple: scaling AI means scaling sandboxes, and sandboxes should run on infrastructure that already scales instead of a parallel stack you have to learn to trust. The catch is that Kubernetes’s unit of work, the pod, was designed for code you trust.
So this is a working answer to secure AI sandboxes on Kubernetes: what a pod actually gives you, which runtime holds against model-written code, what a fork of a running machine costs, and how we made microVMs behave like ordinary pods. By the end you should be able to pick a runtime for your own threat model and fan one warm machine out into thirty.
Model-written code is untrusted code
An LLM cannot tell your instructions from an attacker’s. Simon Willison’s lethal trifecta names the failure mode: private data, untrusted content, and a way to talk out. Combine the three and an injected instruction exfiltrates whatever the agent can read. Give the agent a shell and the injection becomes code execution.
You do not have to take this on faith. Trail of Bits showed in October 2025 that command allowlisting without a sandbox fails: attackers smuggle flags into pre-approved commands like go test -exec or fd -x and get execution from one poisoned prompt. In July 2025 a Replit agent deleted a production database during a code freeze. A month later the Nx supply-chain attack ( s1ngularity ) flipped the direction: malicious npm packages invoked whatever AI CLIs were installed and used them to hunt credentials. GitGuardian counted 2,349 secrets from 1,079 repositories.
The tool vendors already conceded the point. Anthropic ships OS-level sandboxing for Claude Code and reports it cut permission prompts by 84 percent. OpenAI’s Codex runs with network off by default .
Those protect one laptop. Host agents for other people, or fan one agent out into fifty, and you are running a hostile multi-tenant service. The question stops being whether to sandbox. It becomes: which boundary are you willing to rent out?
A pod is a view of your kernel
The default answer on Kubernetes is a pod per agent, so start there and look at what you actually get.
Kubernetes does not run containers. The kubelet calls a runtime over the Container Runtime Interface , usually containerd. containerd starts a shim per pod. The shim invokes runc , and runc builds the thing we call a container: namespaces for the restricted view, cgroups for limits, seccomp to trim syscalls, then execve into your process.
Notice what never happened in that chain. Nothing put a wall between your process and the kernel. The process sees less, but it still talks to the same kernel as every other pod on the node, through an interface of several hundred syscalls. A namespace is a view, not a wall.
That interface has a record. CVE-2019-5736 let a container overwrite the host runc binary through /proc/self/exe . CVE-2024-21626 , the Leaky Vessels file-descriptor leak, scored 8.6 and escaped to the host filesystem. In November 2025, three more runc escapes landed on the same day , all breaking out through procfs writes.
None of this makes containers bad engineering. For code you wrote and reviewed, a hardened container is a reasonable boundary, and most of the internet runs on one. The Kubernetes multi-tenancy docs are honest here too: isolation is a spectrum, and once tenants stop trusting each other, the docs point you at sandboxed runtimes.
Agent code fails the trust test by construction. You did not write it. You did not review it. And an attacker may have steered the model that did.
gVisor, Kata, Firecracker: three ways to raise the wall
Kubernetes has a pluggable seam for exactly this problem. A RuntimeClass maps spec.runtimeClassName to a different shim, so the same pod spec can land on a different isolation mechanism. Three runtimes matter.
gVisor puts a user-space kernel between your workload and the real one. Syscalls get intercepted and re-implemented in Go; the design goal, in gVisor’s own words, is that no syscall passes through directly. It starts in tens of milliseconds, and it backs GKE Sandbox . The price is syscall-shaped: most workloads pay a few percent, syscall-heavy ones can pay multiples. And the wall is made of software.
Kata Containers puts every pod in a real VM with its own guest kernel under KVM. Stronger wall, heavier price: third-party benchmarks put it around 600 ms of added start latency and on the order of 180 MiB per pod with the default VMMs. Approximate numbers, but the shape is right.
Then there is Firecracker , the microVM AWS built to run Lambda. Its authors report about 50 thousand lines of Rust, 96 percent less code than QEMU. It emulates almost nothing: virtio-net, virtio-block, vsock, a serial console, and not much else. The published spec commits to under 125 ms from API call to guest user space and under 5 MiB of overhead per microVM. The NSDI paper measured roughly 3 MB of overhead where QEMU took 131 MB, and reports creation rates up to 150 microVMs per second per host.
A hardware wall at almost container prices. That is why so much of the sandbox category runs on it.
*Third-party benchmark figures, approximate. Firecracker figures are from its published specification. Morph and CodeSandbox also fork running VMs, on their own clouds rather than on Kubernetes.
Here is the decision, compressed to three questions you can paste into a design doc:
Who wrote the code? You, and you reviewed it: runc is fine, spend the savings elsewhere. A model wrote it: keep going.
Does it mostly call libraries, or does it shell out and install packages? Mostly libraries, one sandbox at a time: gVisor will probably hold, and we would rather tell you that than pretend otherwise. Full OS behavior: keep going.
Do many agents need to start from one warm state? No: Kata or any microVM runtime ends the analysis here. Yes: you need a fork, and the rest of this post is about what that takes.
Opting in looks like one field:
apiVersion : node.k8s.io/v1 kind : RuntimeClass metadata : name : kata-fc handler : kata-fc # containerd-shim-kata-v2 + Firecracker --- apiVersion : v1 kind : Pod spec : runtimeClassName : kata-fc
One field, and every pod in the manifest gets its own kernel. If that were the whole story, the post would end here.
Firecracker deleted everything Kubernetes assumes
Everything Firecracker removed to get small is something Kubernetes quietly depends on.
No virtio-fs means no shared filesystems, so the container rootfs has to arrive as a block device, which forces containerd onto the devmapper snapshotter that Kata’s docs require for the Firecracker VMM. No device hotplug and no VFIO means Kata-with-Firecracker cannot resize a running container or pass through a GPU. Memory is sized up front, against a scheduler built for elastic requests. Networking is a tap device per VM, bridged to the pod’s veth by an extra CNI plugin. And firecracker-containerd , the direct integration, still lists CRI conformance under future goals.
Upstream Kubernetes answered in late 2025 with agent-sandbox , a SIG Apps project that wraps Sandbox CRDs, templates, and warm pools around gVisor or Kata. The warm pools exist because, as the Kubernetes blog puts it , starting a new pod adds about a second. If your agents run one sandbox at a time and gVisor clears your threat model, take that path; it is standard and well maintained, and this is not the post that talks you out of it.
But read what it standardizes: one sandbox per agent, created fresh or handed over pre-warmed. Nothing in that stack forks a running machine.
Once your agents multiply, that is the operation you will miss. Here is why.
Copy-on-write changes what a copy costs
The cheapest way to copy a machine is to not copy it. fork() has worked this way for forty years: the child shares the parent’s physical pages read-only, and the kernel copies a page only when someone writes it. A child costs what it changes.
Firecracker snapshots extend that to whole machines. A snapshot is guest memory plus device state in two files. On restore, Firecracker maps the memory file MAP_PRIVATE , so pages load on demand from the host page cache, and every VM restored from that snapshot reads the same physical pages until it writes them.
Run the arithmetic once and it stops feeling like an optimization. Thirty-two forks of a warm 512 MiB Python environment is 16 GiB in naive accounting. Physically, at fork time, it is one 512 MiB page set plus about 3 MiB of private pages per daughter: under 0.7 GiB. The daughters grow as they dirty pages, so that is a floor, not a steady state. But you just provisioned a fleet for the memory of a machine and a half.
This is not a lab trick. Lambda SnapStart resumes functions from snapshots, and Aurora DSQL runs each SQL transaction in its own cloned Firecracker VM, sharing unchanged pages across clones. On our reference hardware, the Mitos engine restores a snapshot in 6 to 16 ms and activates a warm fork in about 27 ms, at about 3 MiB of marginal memory per fork. Engine measurements, reproducible from the benchmark scripts in the repo; what they do and do not include is spelled out on the benchmarks page .
One distinction will save you real confusion when you read vendor pages. Disk copy-on-write (overlayfs layers, qcow2 backing files, thin volumes) clones a filesystem, and plenty of products call a golden image a snapshot. A machine snapshot includes live memory and running processes. The first buys you a cheap clean boot. The second buys you a cheap copy of a machine mid-task, which is the one your agent wants after it spent 40 seconds installing dependencies and loading a model.
Restore a snapshot twice and you have a security bug
This is the part the sandbox write-ups skip, and it is the reason forking is a security mechanism before it is a speed trick.
Firecracker’s own docs warn that when one guest state is resumed more than once, “guest information assumed to be unique may in fact not be.” Think about what lives in that memory image: the kernel’s entropy pool, userspace PRNG state, TCP sequence numbers, machine IDs, cached TLS session keys. Twenty naive restores means twenty guests that can mint colliding UUIDs, session tokens, and nonces.
For code you already assume is adversarial, that is a gift with a bow on it.
The kernel half has a fix: Firecracker writes a fresh VMGenID on resume and Linux 5.18+ reseeds its CSPRNG from it, though the docs admit a race window before the reseed lands. The userspace half has no general fix. The same doc says so outright. Network identity duplicates too: every clone wakes up with the parent’s MAC and IP.
So a correct fork is a protocol, not a file copy. When Mitos activates a daughter it pushes fresh entropy into the guest, waits for the guest agent to confirm the RNG reseed, steps the wall clock off the frozen snapshot time, re-addresses the NIC, and delivers per-daughter secrets. It fails closed: a clone that never confirmed its reseed is never served. Details in the fork-correctness doc .
You can now interrogate any sandbox vendor with three questions. Does “snapshot” include memory and running processes, or is

[truncated]

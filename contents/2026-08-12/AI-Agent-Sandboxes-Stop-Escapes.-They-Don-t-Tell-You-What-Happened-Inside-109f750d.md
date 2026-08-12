---
source: "https://rye.ai/blog/ai-agent-sandboxes-ebpf-runtime-visibility/"
hn_url: "https://news.ycombinator.com/item?id=49278090"
title: "AI Agent Sandboxes Stop Escapes. They Don't Tell You What Happened Inside"
article_title: "AI Agent Sandbox Security: Docker Sandboxes, MicroVMs, and the eBPF Gap | Rye"
author: "wakahiu"
captured_at: "2026-08-12T20:35:38Z"
capture_tool: "hn-digest"
hn_id: 49278090
score: 1
comments: 0
posted_at: "2026-08-12T20:25:22Z"
tags:
  - hacker-news
  - translated
---

# AI Agent Sandboxes Stop Escapes. They Don't Tell You What Happened Inside

- HN: [49278090](https://news.ycombinator.com/item?id=49278090)
- Source: [rye.ai](https://rye.ai/blog/ai-agent-sandboxes-ebpf-runtime-visibility/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T20:25:22Z

## Translation

タイトル: AI エージェントのサンドボックスが逃走を阻止。中で何が起こったのかは教えてくれない
記事のタイトル: AI エージェント サンドボックス セキュリティ: Docker サンドボックス、MicroVM、および eBPF ギャップ |ライ麦
説明: Docker サンドボックスと microVM 分離により、AI エージェントがホストから遠ざけられます。エージェントが内部にいる間に何を読み取り、実行し、変更し、送信したかをセキュリティ チームに伝えることはありません。

記事本文:
AI エージェント サンドボックス セキュリティ: Docker サンドボックス、MicroVM、および eBPF ギャップ |ライ麦 ライ麦 。 ai の機能 記事 ドキュメントの価格 法的サインイン 無料トライアルを開始する メニュー 機能 記事 ドキュメントの価格 法的サインイン 無料トライアルを開始する 記事に戻る 公開日 2026 年 8 月 10 日 読書 8 分で読む 著者 Peter W. Njenga Founder トピック Ai コーディング セキュリティ ランタイム セキュリティの記事
AI エージェント サンドボックスが逃走を阻止します。中で何が起こったのかは教えてくれません。
Docker サンドボックスと microVM 分離により、AI エージェントがホストから遠ざけられます。エージェントが内部にいる間に何を読み取り、実行し、変更し、送信したかをセキュリティ チームに伝えることはありません。
Docker は Docker Sandbox をリリースしました。 macOS、Windows、Linux でのクロスプラットフォーム サポート用に構築されたカスタム VMM を使用して、専用の microVM 内で Claude Code、Codex CLI、Copilot CLI、Kiro、OpenCode などの AI コーディング エージェントを実行します。各サンドボックスは、独自のカーネル、マウントされたワークスペース、および承認されたホスト名のみを許可するネットワーク ポリシーを取得します。ワークスペース マウントは、セッション全体を通じてホスト ファイル システム上でライブになります。 VM の状態 (インストールされたパッケージ、シェル履歴、VM 内に書き込まれたファイル) は再起動後も保持され、サンドボックスが明示的に削除された場合にのみ破棄されます。
エージェントを YOLO モード ( --dangerously-skip-permissions ) で実行している開発者にとって、これは大きな改善です。エージェントが横向きになったとしても、ホストはクリーンな状態を保つ必要があります。
しかし、サンドボックスが答えるのは 1 つの質問だけです。それは、エージェントが脱出したのかということです。
ほとんどのチームは、最初に別の答えを必要とします。つまり、エージェントがそこにいる間に何をしたか?
なぜコンテナだけでは十分ではなかったのか
Docker Sandbox が登場する前は、エージェントを Docker コンテナに入れるのが一般的なアドバイスでした。プロジェクトディレクトリのみをマウントします。ジョブが完了したらコンテナを削除します。それは実際よりもきれいに聞こえます。
Docker コンテナはホスト カーネルを共有します。名前

スペースと cgroup は役立ちますが、同じカーネルが境界を強制します。コンテナー エスケープは CVE の繰り返しクラスです。runc の上書きの場合は CVE-2019-5736、cgroup エスケープの場合は CVE-2022-0492、runc の再場合は CVE-2024-21626 です。これらはいずれもカジュアルな攻撃ではありません。それでも、プロンプト インジェクションを通じて操作される AI エージェントは、私がデフォルトで信頼できるプロセスではありません。
Firecracker microVM はリスクの形を変えます。各サンドボックスは独自の Linux カーネルを実行し、KVM によってホストから分離されます。ゲスト カーネルのバグがホスト カーネルのバグになってはいけません。これは、AWS Lambda および Fly.io マシンで使用されるものと同じ基本的な分離モデルであり、信頼できない顧客のコードが共有ハードウェアで実行されます。
したがって、分離プリミティブが正しいものになります。
問題は境界が成立した後に始まります。
正確にするのに役立ちます。ネットワーク ポリシーを備えた microVM サンドボックスは、次の攻撃対象領域を狭めます。
ホストファイルシステムへのアクセス。エージェントは、明示的にマウントされたプロジェクト ワークスペースのみを表示および書き込みできます。 ~/.ssh/ 、 ~/.aws/credentials 、シェル履歴、またはマシン上のその他のファイルを読み取ることはできません。
ホストプロセスへのアクセス。エージェントはホスト プロセスを認識したり、信号を送信したりすることができません。 IDE にデバッガを接続したり、VPN クライアントを強制終了したり、実行中の他のエージェントを改ざんしたりすることはできません。
横方向のネットワークの移動。ホワイトリストを除くすべてを拒否するネットワーク ポリシーを使用すると、エージェントは内部ネットワーク、クラウド メタデータ エンドポイント ( 169.254.169.254 )、または任意のインターネット インフラストラクチャに到達できなくなります。承認したドメインとのみ通信できます。
破棄時の VM 層の永続化。サンドボックスが明示的に削除されると、インストールされたパッケージ、シェル履歴、およびマウントされたワークスペース外の VM 内に書き込まれたファイルは破棄されます。ワークスペース自体はライブ マウントです。エージェントは、コピーではなく、セッション全体を通じてホスト ファイルを直接読み書きします。 V

M 状態は、dispose コマンドを実行するまで、再起動後も維持されます。
これらの制御は、特に夜間のリファクタリング、CI コード レビュー エージェント、自律的なテスト生成などの無人ジョブの場合に重要です。私は開発者のラップトップではなく、microVM でこれらを実行したいと考えています。
サンドボックスの境界は microVM の境界です。その境界内には、実際の作業を行う忙しい小さなマシンがまだあります。
エージェントのアクションの監査証跡はありません。エージェントは、ファイルの読み取り、書き込み、シェル コマンドの実行、およびネットワーク接続のオープンを行います。サンドボックスは、エージェント自身のセッション ログの外部にある構造化ログにそれを記録しません。そのログはサンドボックス内にあり、エージェント プロセスによって制御されます。セッションが終了する前にエージェントがそれを削除した場合、レコードも一緒に残ります。
エージェントは依然としてプロジェクトに損害を与える可能性があります。プロンプトによって挿入されたエージェントは、ワークスペースの削除、構成の上書き、資格情報がマウントされている場合の git リモートへのプッシュ、または許可されたホスト名へのソース コードの送信を行うことができます。サンドボックスにより、ホストへの到達が阻止されました。与えられたアクセス権の使用は阻止されませんでした。
ネットワーク許可リストは単純です。 api.github.com 、 registry.npmjs.org 、または pypi.org が許可されており、おそらく実際の開発作業用である場合、悪意のあるインストールまたは侵害されたリモートには有効なパスが存在します。ホワイトリストはランダムな攻撃者のインフラストラクチャをブロックします。開いているはずのチャネルの悪用は阻止されません。
事後に起こったことを再現することはできません。サンドボックス セッションが予期しない出力を生成したり、ファイルを削除したり、予期しないコードをコミットしたりした場合、監査するカーネル レベルの記録はありません。エージェントのセッション ログが存在し、改ざんされていない場合は、モデルの意図がわかります。システムコールレベルで実際に何が実行されたかはわかりません。
それが盲点です。 Docker Sandbox はその仕事を完璧に実行できます

それでも、事件を説明できないままになります。
eBPF は、信頼できるアクティビティ ログが必要になるまでは、カーネルのトリビアのように聞こえます。これを使用すると、syscall の開始と終了、ネットワーク接続、ファイルシステム操作などのカーネル イベントに小さなプログラムを接続できます。 Cilium Tetragon や Falco などのツールは、これを使用して、プロセスが実際に行うことの構造化イベントを生成します。
スクロールしてズーム · ドラッグしてパン · ダブルクリックしてリセット · ESC で閉じます
コンテナ (左): ホスト側の eBPF は、エージェントが行うすべての Syscall を確認します。 microVM (右): KVM ハードウェア境界により、ホスト kprobe がゲスト パス、PID、または argv を認識できなくなります。 eBPF はゲスト内で実行する必要があります。
コンテナーの場合、これはホストから機能します。コンテナーはホスト カーネルを共有するため、ホスト上で実行されている Tetragon または Falco は、エージェントが行うすべての syscall を認識します。
すべての open() 、 read() 、 write() 、 unlink() 呼び出し (完全な解決されたパス付き)
すべての execve() - 生成されたすべてのサブプロセスとその引数
すべてのアウトバウンド TCP 接続 (宛先 IP とポートを含む)
Unix ソケットへのすべての connect()
すべての clone() または fork() - すべての子プロセス
これはサンプリングではありません。イベントはカーネル イベント自体と同期して生成されます。エージェントはそれらを抑制したり削除したりすることはできません。レコードはシステムコールが返される前に書き込まれます。
microVM の場合は状況が異なります。ゲスト Syscall はゲスト カーネルにトラップします。 KVM は VM-exit をホストにのみ表示します。ホスト側の kprobe とトレースポイントは、ゲスト VFS パス、ファイル名、または PID を認識できません。接続先の共有カーネルがありません。 microVM 内で同じシステムコール レベルのテレメトリを取得するには、ゲスト内で eBPF を実行するか、VMM 境界、ワークスペース マウント レイヤー、またはサンドボックス ランタイム自体が提供するフックでのテレメトリに依存する必要があります。テレメトリ エージェントが

ワークロードよりも高い権限ですが、デプロイメントはホスト上で DaemonSet をドロップするよりも複雑です。
完全な監査証跡とはどのようなものなのか
以下は、Claude Code セッションから得たい種類の Tetragon 出力です。microVM の場合はゲスト内 Tetragon、コンテナーの場合はホスト側 Tetragon によって生成されます。
{"プロセス": {"pid": 1847、"バイナリ": "/usr/bin/node"、"引数": "クロード --dangerously-skip-permissions"},
"アクション": "開く"、"パス": "/workspace/src/auth/session.ts"、"フラグ": "O_RDONLY"}
{"プロセス": {"pid": 1847, "バイナリ": "/usr/bin/node"},
"アクション": "開く"、"パス": "/workspace/.env"、"フラグ": "O_RDONLY"}
{"プロセス": {"pid": 2103, "バイナリ": "/bin/bash", "引数": "npm install lodash-contrib"},
"アクション": "接続"、"宛先": "104.16.1.35:443"、"ホスト名": "registry.npmjs.org"}
{"プロセス": {"pid": 2103, "バイナリ": "/bin/bash"},
"アクション": "execve"、"パス": "/workspace/node_modules/.bin/postinstall-hook"、"引数": ""}
{"プロセス": {"pid": 1847, "バイナリ": "/usr/bin/node"},
"アクション": "開く"、"パス": "/workspace/src/auth/session.ts"、"フラグ": "O_WRONLY|O_TRUNC"}
このストリームから、インシデントの質問ははるかに曖昧ではなくなります。
エージェントは認証コードを変更する前に .env を読み取りましたか? (はい。)
パッケージのインストール後にポストインストールフックが実行されましたか? （はい。次に何をしましたか？）
エージェントは、変更するための明示的な指示がないファイルを書き込み用に開きましたか?
サブプロセスが許可リストにない IP への接続を試みましたか?
これらは通常の事件に関する質問です。別のイベント ストリームがないと、誰かが回答する前にサンドボックスが消える可能性があります。
Tetragon と Falco: 実践的な出発点
Cilium Tetragon はこの仕事にとってより強力なツールです。 TracingPolicy リソースをサポートしているため、関心のあるイベント ( open 、 execve 、 connect 、 clone ) を選択できます。

、バイナリ フィルターの処理、および構造化された JSON エクスポート。コンテナーの場合は、ホスト上の DaemonSet として、またはスタンドアロンのバイナリとして実行されます。 microVM の場合、ゲスト内で実行する必要があります。これは、VM イメージまたはサンドボックス起動スクリプトの一部としてプロビジョニングすることを意味します。
AI エージェント セッション用の最小限の Tetragon ポリシー:
APIバージョン: cilium.io/v1alpha1
種類: トレースポリシー
メタデータ:
名前: ai-エージェント-監査
仕様:
kプローブ:
- 呼び出し:「fd_install」
システムコール: false
引数:
- インデックス: 0
型: int
- インデックス: 1
タイプ: 「ファイル」
セレクター:
- マッチバイナリ:
- 演算子: で
値:
- 「/usr/bin/node」
- 「/usr/bin/python3」
- 「/bin/bash」
- 「/bin/sh」
- 呼び出し:「sys_execve」
システムコール: true
引数:
- インデックス: 0
タイプ: "文字列"
- インデックス: 1
タイプ: "文字列配列"
- 呼び出し:「tcp_connect」
システムコール: false
引数:
- インデックス: 0
タイプ：「靴下」
Falco は、すでにデプロイされている場合は開始しやすいです。そのルール言語はより単純であり、デフォルトのルールは多くの奇妙な動作を検出します。トレードオフは詳細です。 Falco はより高いレベルで動作するため、システムコール レベルのコンテキストの一部を見逃す可能性があります。 AI エージェントのワークロードの場合、興味深いシグナルが「資格情報ファイルの読み取り」または「奇妙なサブプロセスの生成」である可能性があるため、Tetragon の追加の詳細は通常、セットアップの価値があります。
これにより AI エージェントのセキュリティ スタックが残される場所
microVM サンドボックスと eBPF 監査レイヤーは、脅威モデルのさまざまな部分に対処するため、一緒に使用する必要があります。
これらの層のどれも全体の負荷を担っていません。テレメトリのないサンドボックスはブラック ボックスです。封じ込めを行わない遠隔測定により、被害の発生を監視できます。エージェントの権限は役立ちますが、監視対象の製品内で実行されます。監査証跡のないポリシーは、事後的に防御するのが困難です。
アプリケーション層では、Rye はエージェント ランタイムの外側に位置します。 CLI セッションをラップし、モデル トラフィックをローカル ポリシー プロキシ経由でルーティングし、承認を記録し、ファイル C

Hange を実行し、メタデータを要求し、ツール全体で監査イベントを正規化します。カーネル テレメトリは、サンドボックス内で実行された内容に応答します。 Rye は、そのアクティビティをエージェント セッション、モデル要求、ポリシー決定、およびワークスペース コンテキストに接続します。
業界は封じ込めに多大なエネルギーを費やしてきました。良い。今では、封じ込められた環境内の可視性が必要です。
砂場が逃走を阻止した。監査ログは、実際に何が起こったかを示します。
なぜコンテナだけでは十分ではなかったのか
完全な監査証跡とはどのようなものなのか
Tetragon と Falco: 実践的な出発点
これにより AI エージェントのセキュリティ スタックが残される場所

## Original Extract

Docker Sandboxes and microVM isolation keep AI agents away from the host. They do not tell security teams what the agent read, ran, changed, or sent while it was inside.

AI Agent Sandbox Security: Docker Sandboxes, MicroVMs, and the eBPF Gap | Rye rye . ai Features Articles Docs Pricing Legal Sign in Start free trial Menu Features Articles Docs Pricing Legal Sign in Start free trial Back to articles Published August 10, 2026 Reading 8 min read Author Peter W. Njenga Founder Topics Ai Coding Security Runtime Security Articles
AI Agent Sandboxes Stop Escapes. They Don't Tell You What Happened Inside.
Docker Sandboxes and microVM isolation keep AI agents away from the host. They do not tell security teams what the agent read, ran, changed, or sent while it was inside.
Docker just shipped Docker Sandboxes . It runs AI coding agents such as Claude Code, Codex CLI, Copilot CLI, Kiro, and OpenCode inside dedicated microVMs using a custom VMM built for cross-platform support on macOS, Windows, and Linux. Each sandbox gets its own kernel, a mounted workspace, and a network policy that only allows approved hostnames. The workspace mount is live on your host filesystem throughout the session. VM state - installed packages, shell history, files written inside the VM - persists across restarts and is discarded only when the sandbox is explicitly removed.
For developers running agents in YOLO mode ( --dangerously-skip-permissions ), this is a real improvement. If the agent goes sideways, the host should stay clean.
But a sandbox only answers one question: did the agent get out?
Most teams will need a different answer first: what did the agent do while it was in there?
Why Containers Were Never Enough
Before Docker Sandboxes, the usual advice was to put the agent in a Docker container. Mount only the project directory. Delete the container when the job is done. That sounds cleaner than it is.
Docker containers share the host kernel. Namespaces and cgroups help, but the same kernel still enforces the boundary. Container escapes are a recurring class of CVE: CVE-2019-5736 for runc overwrite, CVE-2022-0492 for cgroup escape, CVE-2024-21626 for runc again. None of these are casual attacks. Still, an AI agent manipulated through prompt injection is not a process I would trust by default.
Firecracker microVMs change the shape of the risk. Each sandbox runs its own Linux kernel and is isolated from the host by KVM . A guest kernel bug should not become a host kernel bug. This is the same basic isolation model used by AWS Lambda and Fly.io machines, where untrusted customer code runs on shared hardware.
So the isolation primitive is the right one.
The problem starts after the boundary holds.
It helps to be precise. A microVM sandbox with network policy narrows these attack surfaces:
Host filesystem access. The agent can only see and write the project workspace that was explicitly mounted. It cannot read ~/.ssh/ , ~/.aws/credentials , your shell history, or any other file on your machine.
Host process access. The agent cannot see or signal host processes. It cannot attach a debugger to your IDE, kill your VPN client, or tamper with other running agents.
Lateral network movement. With a deny-all-except-allowlist network policy, the agent cannot reach your internal network, your cloud metadata endpoint ( 169.254.169.254 ), or arbitrary internet infrastructure. It can only talk to the domains you approved.
VM-layer persistence on dispose. When the sandbox is explicitly removed, installed packages, shell history, and files written inside the VM outside the mounted workspace are discarded. The workspace itself is a live mount - the agent reads and writes your host files directly throughout the session, not copies. VM state persists across restarts until you run the dispose command.
Those controls matter, especially for unattended jobs: nightly refactors, CI code review agents, and autonomous test generation. I would rather run those in a microVM than on a developer laptop.
The sandbox boundary is the microVM perimeter. Inside that perimeter, you still have a busy little machine doing real work.
You have no audit trail of agent actions. The agent reads files, writes files, runs shell commands, and opens network connections. The sandbox does not record that in a structured log outside the agent's own session log . That log is inside the sandbox and controlled by the agent process. If the agent deletes it before the session ends, the record goes with it.
The agent can still damage the project. A prompt-injected agent can delete the workspace, overwrite config, push to git remotes if credentials are mounted, or send source code to an allowed hostname. The sandbox stopped it from reaching the host. It did not stop it from using the access you gave it.
The network allowlist is blunt. If api.github.com , registry.npmjs.org , or pypi.org are allowed, and they probably are for real development work, a malicious install or compromised remote has a valid path out. The allowlist blocks random attacker infrastructure. It does not block misuse of channels that are supposed to be open.
You cannot reconstruct what happened after the fact. If a sandbox session produces unexpected output, deletes files, or commits surprising code, you have no kernel-level record to audit. The agent's session log, if it exists and was not tampered with, tells you what the model intended. It does not tell you what actually executed at the syscall level.
That is the blind spot. Docker Sandboxes can do its job perfectly and still leave you unable to explain an incident.
eBPF sounds like kernel trivia until you need a trustworthy activity log. It lets you attach small programs to kernel events such as syscall entry and exit, network connections, and filesystem operations. Tools like Cilium Tetragon and Falco use it to produce structured events for what a process actually does.
scroll to zoom · drag to pan · double-click to reset · ESC to close
Container (left): host-side eBPF sees every syscall the agent makes. microVM (right): the KVM hardware boundary blocks host kprobes from seeing guest paths, PIDs, or argv. eBPF must run inside the guest.
For containers , this works from the host. Because the container shares the host kernel, Tetragon or Falco running on the host sees every syscall the agent makes:
Every open() , read() , write() , unlink() call, with the full resolved path
Every execve() - every subprocess spawned, with its argv
Every outbound TCP connection, with destination IP and port
Every connect() to a Unix socket
Every clone() or fork() - every child process
This is not sampling. Events are generated synchronously with the kernel events themselves. The agent cannot suppress or delete them - the record is written before the syscall returns.
For microVMs, the picture is different. Guest syscalls trap to the guest kernel. KVM only surfaces VM-exits to the host. Host-side kprobes and tracepoints cannot see guest VFS paths, file names, or PIDs - there is no shared kernel for them to attach to. To get the same syscall-level telemetry inside a microVM, eBPF has to run inside the guest, or you rely on telemetry at the VMM boundary, the workspace mount layer, or hooks the sandbox runtime itself provides. The tamper-resistance property still holds when the telemetry agent runs at higher privilege than the workload, but the deployment is more involved than dropping a DaemonSet on the host.
What a Complete Audit Trail Looks Like
Here is the kind of Tetragon output I would want from a Claude Code session - produced by in-guest Tetragon for a microVM, or host-side Tetragon for a container:
{"process": {"pid": 1847, "binary": "/usr/bin/node", "arguments": "claude --dangerously-skip-permissions"},
"action": "open", "path": "/workspace/src/auth/session.ts", "flags": "O_RDONLY"}
{"process": {"pid": 1847, "binary": "/usr/bin/node"},
"action": "open", "path": "/workspace/.env", "flags": "O_RDONLY"}
{"process": {"pid": 2103, "binary": "/bin/bash", "arguments": "npm install lodash-contrib"},
"action": "connect", "destination": "104.16.1.35:443", "hostname": "registry.npmjs.org"}
{"process": {"pid": 2103, "binary": "/bin/bash"},
"action": "execve", "path": "/workspace/node_modules/.bin/postinstall-hook", "arguments": ""}
{"process": {"pid": 1847, "binary": "/usr/bin/node"},
"action": "open", "path": "/workspace/src/auth/session.ts", "flags": "O_WRONLY|O_TRUNC"}
From that stream, the incident questions get much less fuzzy:
Did the agent read .env before modifying auth code? (Yes.)
Did a postinstall hook run after a package was installed? (Yes. What did it do next?)
Did the agent open a file for writing that it did not have explicit instructions to modify?
Did any subprocess attempt a connection to an IP not on the allowlist?
Those are normal incident questions. Without a separate event stream, the sandbox may disappear before anyone can answer them.
Tetragon and Falco: Practical Starting Points
Cilium Tetragon is the stronger tool for this job. It supports TracingPolicy resources, so you can choose the events you care about: open , execve , connect , clone , process binary filters, and structured JSON export. For containers it runs as a DaemonSet on the host or as a standalone binary. For microVMs it needs to run inside the guest, which means provisioning it as part of the VM image or sandbox startup script.
A minimal Tetragon policy for AI agent sessions:
apiVersion: cilium.io/v1alpha1
kind: TracingPolicy
metadata:
name: ai-agent-audit
spec:
kprobes:
- call: "fd_install"
syscall: false
args:
- index: 0
type: int
- index: 1
type: "file"
selectors:
- matchBinaries:
- operator: In
values:
- "/usr/bin/node"
- "/usr/bin/python3"
- "/bin/bash"
- "/bin/sh"
- call: "sys_execve"
syscall: true
args:
- index: 0
type: "string"
- index: 1
type: "string_array"
- call: "tcp_connect"
syscall: false
args:
- index: 0
type: "sock"
Falco is easier to start with if it is already deployed. Its rule language is simpler, and the default rules catch plenty of strange behavior. The tradeoff is detail. Falco works at a higher level and can miss some syscall-level context. For AI agent workloads, where the interesting signal may be "read a credential file" or "spawned a weird subprocess," Tetragon's extra detail is usually worth the setup.
Where This Leaves the AI Agent Security Stack
A microVM sandbox and an eBPF audit layer address different parts of the threat model and should be used together:
None of these layers carries the whole load. A sandbox without telemetry is a black box. Telemetry without containment lets you watch the damage happen. Agent permissions help, but they run inside the product you are trying to police. Policy without an audit trail is hard to defend after the fact.
At the application layer, Rye sits outside the agent runtime. It wraps CLI sessions, routes model traffic through a local policy proxy, records approval, file-change, and request metadata, and normalizes audit events across tools. Kernel telemetry answers what ran inside the sandbox; Rye connects that activity to the agent session, model request, policy decision, and workspace context.
The industry has spent a lot of energy on containment. Good. Now it needs visibility inside the contained environment.
The sandbox stopped the escape. The audit log tells you what actually happened.
Why Containers Were Never Enough
What a Complete Audit Trail Looks Like
Tetragon and Falco: Practical Starting Points
Where This Leaves the AI Agent Security Stack

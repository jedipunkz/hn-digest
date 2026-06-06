---
source: "https://www.tigera.io/blog/multi-layer-policy-for-securing-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48418893"
title: "Multi-Layer Policy for Securing AI Agents"
article_title: "Multi-Layer Policy for Securing AI Agents"
author: "baroiall"
captured_at: "2026-06-06T00:55:29Z"
capture_tool: "hn-digest"
hn_id: 48418893
score: 1
comments: 1
posted_at: "2026-06-05T21:57:02Z"
tags:
  - hacker-news
  - translated
---

# Multi-Layer Policy for Securing AI Agents

- HN: [48418893](https://news.ycombinator.com/item?id=48418893)
- Source: [www.tigera.io](https://www.tigera.io/blog/multi-layer-policy-for-securing-ai-agents/)
- Score: 1
- Comments: 1
- Posted: 2026-06-05T21:57:02Z

## Translation

タイトル: AI エージェントを保護するための多層ポリシー
説明: AI エージェントのセキュリティには、L7 ゲートウェイとカーネルの両方で 1 つの言語で表現されたポリシーが必要です。 Cedar で二重層強制が機能する理由は次のとおりです。

記事本文:
AI エージェントを保護するための多層ポリシー
製品
AIエージェント向け
Lynx AI エージェント セキュリティ プラットフォーム
AI ワークロードの場合
Calico オープンソース eBPF ベースのネットワーキングとセキュリティ
Calico 商用エディション Calico Cloud および Calico Enterprise
ソリューション
使用例
AI ワークロード
VMの移行
可観測性とトラブルシューティング
学ぶ
開発者センター
ドキュメント
Kubernetes 用の VM ネットワーキング NEW Kubernetes 上で VM ネットワーキングを移行、保護、運用するための実践者向けガイドを、1 つの VM の行程を通して説明します。さらに詳しく >
ガイド
Kubernetes
Kubernetes 101
ガイド
可観測性
可観測性
ネットワーキング
Kubernetes ネットワーキング
AI エージェントを保護するための多層ポリシー
Tigera では、現実世界の大規模なエンタープライズ エージェント向けに安全なランタイム環境を作成する製品を構築する製品を構築する仕事の一環として、私がよく考えているこのパズルの小さな部分の 1 つは、ポリシーとポリシーの実行時適用、および 1 か所から構成された包括的な安全なランタイムを作成する方法です。実行時にこれらのプラットフォームをロックダウンして保護しようとしている企業と話をすればするほど、AI エージェントのセキュリティには 1 か所だけでなく (たとえば、ゲートウェイ層だけでなく) 複数の場所でポリシーが必要であり、理想的には同じポリシー言語で表現される必要があると私は考えています。
L7 ゲートウェイ層では、すべてのエージェントの通話が監視可能です。誰が通話しているのか、何を通話しているのか、双方がどのような属性を持っているか、要求されたアクションは何なのかなどです。ここで、エージェントが特定の MCP サーバーと通信すること、特定のツールを呼び出すこと、別のエージェントに委任すること、または特定の LLM を呼び出すことを許可するかどうかを決定します。ここでのポリシーの要素は、アイデンティティ、アクション、リソース、およびコンテキストです。
エージェント ランタイム層、またはコンテナ内のカーネル層では、システムコール、ファイル アクセス、ライブラリ ロアなど、エージェントが独自のランタイム内で行うことを観察できます。

ds、仲介チャネルをバイパスするネットワーク接続。ここで、エージェントがファイルを読み取るか、ソケットを開くか、サブプロセスを生成するか、ライブラリをロードできるかどうかを決定します。ここでのポリシーの要素は、プロセス、パス、ファイル記述子、およびシステム コールです。
両方の層が必要です。ゲートウェイだけでは、トークンを保持したエージェントがランタイム内で行うことを制限することはできません。カーネルだけでは、アイデンティティ、委任、またはマルチホップの意図について推論することはできません。一方のポリシーを構築し、他方のポリシーを構築しないと、カテゴリーのギャップが残ります。
これを実際に機能させるためのアーキテクチャ上の選択は、両方に 1 つのポリシー言語を使用することです。ゲートウェイで Cedar を使用し、Cedar をカーネルの eBPF ポリシーに解釈して変換します。同じポリシー、2 つの適用ポイント、1 か所で作成およびレビューできます。
ゲートウェイのポリシー: エージェントの意図を強制する
ゲートウェイは意図を認識します。ここは、誰が、どのような状況で、どのレベルの人間の監視の下で、誰と会話できるかを強制するのに適した場所です。
どのエージェントがどのツールを起動できるかを制限する Cedar ポリシー:
許可（
Group::"finance-agents" のプリンシパル、
アクション == アクション::"invokeTool",
ToolSet のリソース::"finance-readonly"
) のとき {
プリンシパル.リスクレベル == "低" &&
context.delegation_ Depth <= 3
};
このポリシーは、RBAC またはネットワーク ポリシーでモデル化するのが難しいいくつかのことを表します。プリンシパルはグループのメンバーシップによって識別されますが、属性 (risk_level) によって制限されます。リソースは、型付きのツールのセットです。条件には委任の深さのチェックが含まれます。委任チェーンの 3 ホップ深いエージェントは、たとえ 1 つおきのチェックに合格したとしても拒否されます。
ゲートウェイ層は、委任ルール、スコープ縮小によるホップごとのトークン発行、エージェントから MCP ツールへの承認、エージェントから LLM への制約、ハイステークス AC 用の人間参加型フックを自然に強制します。

これらすべてにわたる属性ベースの決定。
ゲートウェイができないことは、トークンを発行した後の動作を制限することです。エージェントが資格情報を取得すると、カーネルは、プロセスがそれに対して実際に何を行うかを認識する唯一の層になります。
カーネルのポリシー: エージェントの動作を制限する
カーネルは動作を認識します。これは、保持しているトークンに関係なく、エージェント プロセスがオペレーティング システム レベルで実行できることを強制するのに適した場所です。
エージェント ワークロードのベースライン サンドボックス。同じ Cedar ポリシー モデルで概念的に表現され、実行時に BPF プログラムにコンパイルされます。
許可（
AgentClass::"data-analyst" のプリンシパル、
[Action::"readFile", Action::"writeFile"] のアクション、
リソースはファイルパスです
) のとき {
"/workspace/analyst-*" のような resource.path ||
resource.path == "/var/run/secrets/analyst-key"
};
禁止します（
AgentClass::"data-analyst" のプリンシパル、
アクション == アクション::"ネットワーク接続",
リソースは NetworkDestination です
) { を除く
DestinationSet の resource.host::"approved-llm-endpoints" ||
resource.host == "lynx-gateway.internal"
};
コンパイルのターゲットは、BPF LSM フック、cgroup ネットワーク フック、および inode 境界でのファイル アクセス強制です。エージェント プロセスがポリシーで許可されている範囲を超えた場合、カーネルは操作を拒否します。システムコールの場合は EPERM、ネットワーク接続の場合は ECONNREFUSED、ファイル アクセスの場合は ENOENT です。エージェントは、保持している資格情報に関係なく、禁止された操作の場合と同じエラーを受け取ります。
カーネル層は、ファイル アクセス境界、ネットワーク出力制限、システムコール ホワイトリスト、ライブラリ負荷制限、およびプロセス生成制限を自然に強制します。執行にフィードを提供するのと同じ監視パイプラインが、脅威の検出にもフィードします。
カーネルができないことは、アクションが実行される理由を推論することです

落ちた。 connect() システムコールが表示されます。その呼び出しが、ゲートウェイがすでに承認した正当なマルチホップ委任の一部であるかどうかはわかりません。そのコンテキストは L7 層にのみ存在します。
アーキテクチャの統合は、どちらかのレイヤーが分離されているのと同じくらい重要です。 Cedar ポリシーは一度作成され、ゲートウェイで評価され、カーネル適用のために BPF にコンパイルされます。コンパイルは魔法ではなく、Cedar ポリシーの基板関連のサブセットのみがコンパイルされます。残りはゲートウェイに留まります。いずれにしても、セキュリティ チームは Cedar を作成します。どのレイヤーで強制するのが適切であるかは、ランタイムによって決定されます。
この統合により、二重層アプローチが運用上持続可能になります。 1 つのポリシー言語がないと、2 つのポリシー ストア、2 つのレビュー プロセス、2 つのエンジニアリング チームが必要になり、ゲートウェイで許可される内容とカーネルで許可される内容との間に必然的な相違が生じます。両方の層に Cedar があるため、作成したポリシーはどこでも適用されるポリシーになります。
AI エージェントのセキュリティに単層ポリシーでは不十分な理由
ゲートウェイのポリシーだけで、不正な呼び出し元や範囲外のアクションを防御します。正規のトークンを保持し、資格情報ファイルの読み取り、サイドチャネルを介したデータの漏洩、ランタイム内の特権の昇格など、本来の動作以外のことを行うためにそのトークンを使用する侵害されたエージェントを防御することはできません。
カーネルのポリシーだけで、プロセスレベルの不正行為を防御できます。 ID や委任を理解せず、ネットワーク接続が正当なマルチホップ チェーンの一部であるかどうかを判断することができず、人間による承認フローを強制する方法もありません。
2 つの層を組み合わせることで、どちらかの層だけでは見逃す脅威モデルをカバーできます。正規のトークンを持った侵害されたエージェントは、引き続きゲートウェイ経由で発信できますが、

ローカル アクションはカーネル サンドボックスによって制限されます。ゲートウェイで誤って設定された Cedar ポリシーは、サブストレート ベースラインによって軽減されます。登録されていないシャドウ エージェントが監視され、カーネルに封じ込められます。
規制されたワークロードにエージェント インフラストラクチャを構築する Kubernetes ネイティブの企業にとって、これは構築する価値のあるアーキテクチャです。エージェントが要求できるものに関するゲートウェイ ポリシー。何が許可されているかに関するカーネル ポリシー。両方とも同じ言語です。
多層ポリシーは、AI エージェント インフラストラクチャにエンドツーエンドの責任を持たせるという、より大きな問題の 1 つです。トレーサビリティ、認可の出所、アイデンティティと所有権、大規模なポリシーベースのガバナンス、人間による監視と介入、これらはすべて連携して機能する必要があります。
読む: AI エージェントの責任の 5 つの柱 →
ブログ投稿、ワークショップ、認定プログラム、新しいリリースなどに関する最新情報を入手してください。
技術ブログ
AI エージェントの責任のギャップ: ネットワーク ポリシー、API ゲートウェイ、RBAC だけでは不十分な理由
アリスター・バロイ著
2026 年 5 月 27 日
「AI エージェントの責任の 5 つの柱: エンジニアリング リーダーのための診断フレームワーク」では、AI エージェントの責任の各柱 (トレーサビリティ、認可の来歴、身元と所有権、大規模なポリシー、人間の監視) について説明しました...
技術ブログ
2026 年の VM とコンテナの統合の事例
ディロン・バリー著
2026 年 5 月 26 日
2 つのプラットフォーム、2 つのチーム、2 つの調達関係がすべて 1 つの仕事を行っています。このような結果になったのには理由があります。このままにしなければならない理由はありません。一般的な企業の人に聞いてみてください...
技術ブログ
Kubernetes 運用の成熟度: クラスター メッシュを使用した安全で復元力のあるクラスター フェデレーション
ベロニカ・スモリク著
2026 年 5 月 25 日
実際には、単一の Kubernetes クラスターを実行する人は誰もいません。

最近の生産。おそらくそれが始まりだったのですが、データ主権の要件、買収、AI への取り組み、エッジ サーバーの必要性などの考慮事項により、後退してしまいました...
このフォームを送信することにより、Tigera がプライバシー ポリシーに従ってお客様の個人情報を処理することを認め、これに同意したものとみなされます。
可観測性とトラブルシューティング
著作権©
株式会社タイガーラの著作権はすべて留保されます。

## Original Extract

AI agent security needs policy at both the L7 gateway and the kernel, expressed in one language. Here's why dual-layer enforcement works, with Cedar.

Multi-Layer Policy for Securing AI Agents
Products
For AI Agents
Lynx AI agent security platform
For AI Workloads
Calico Open Source eBPF-based networking & security
Calico Commercial Editions Calico Cloud & Calico Enterprise
Solutions
Use Cases
AI WORKLOADS
VM Migration
Observability & Troubleshooting
Learn
Developer Center
Documentation
VM networking for Kubernetes NEW A practitioner’s guide to migrating, securing, and operating VM networking on Kubernetes told through one VM’s journey. Learn More >
Guides
Kubernetes
Kubernetes 101
Guides
Observability
Observability
Networking
Kubernetes Networking
Multi-Layer Policy for Securing AI Agents
As part of our work at Tigera building products that create secure runtime environments for enterprise agents at scale in the real world, one small part of this puzzle I think about a lot is policy, and runtime enforcement of policy, and how to create a comprehensive secure runtime, configured from one place. The more companies we talk to trying to lock down and secure these platforms at runtime, the more I believe AI Agent security needs policy in multiple places, not just one (e.g., not just at the gateway layer), and ideally expressed in the same policy language.
At the L7 gateway layer, every agent call is observable: who is calling, what they are calling, what attributes both sides carry, what the requested action is. This is where you decide whether an agent should be permitted to talk to a particular MCP server, invoke a particular tool, delegate to another agent, or call a particular LLM. The atoms of policy here are identity, action, resource, and context.
At the agent runtime layer, or kernel layer in a container, what the agent does inside its own runtime is observable: syscalls, file access, library loads, network connections that bypass the brokered channel. This is where you decide whether the agent can read a file, open a socket, spawn a subprocess, or load a library. The atoms of policy here are processes, paths, file descriptors, and system calls.
Both layers are necessary. The gateway alone cannot constrain what an agent does inside its runtime once it holds a token. The kernel alone cannot reason about identity, delegation, or multi-hop intent. Building policy at one and not the other leaves a category gap.
The architectural choice that makes this work in practice is using one policy language for both. We use Cedar at the gateway and interpret and translate Cedar to eBPF policy for the kernel: same policies, two enforcement points, one place to author and review.
Policy at the gateway: enforcing agent intent
The gateway sees intent. It is the right place to enforce who can talk to whom, under what conditions, with what level of human oversight.
A Cedar policy that constrains which agents can invoke which tools:
permit (
principal in Group::"finance-agents",
action == Action::"invokeTool",
resource in ToolSet::"finance-readonly"
) when {
principal.risk_level == "low" &&
context.delegation_depth <= 3
};
This policy expresses several things that are hard to model in RBAC or in a network policy. The principal is identified by group membership but constrained by attribute ( risk_level ). The resource is a typed set of tools. The condition includes a check on delegation depth; agents three hops deep in a delegation chain are refused even if they pass every other check.
The gateway layer naturally enforces delegation rules, per-hop token issuance with scope reduction, agent-to-MCP tool authorization, agent-to-LLM constraints, human-in-the-loop hooks for high-stakes actions, and attribute-based decisions across all of these.
What the gateway cannot do is constrain what happens after it issues a token. Once the agent has the credential, the kernel is the only layer that sees what the process actually does with it.
Policy at the kernel: constraining agent behaviour
The kernel sees behaviour. It is the right place to enforce what an agent process is allowed to do at the operating system level , regardless of what tokens it holds.
A baseline sandbox for an agent workload, expressed conceptually in the same Cedar policy model and compiled to BPF programs at runtime:
permit (
principal in AgentClass::"data-analyst",
action in [Action::"readFile", Action::"writeFile"],
resource is FilePath
) when {
resource.path like "/workspace/analyst-*" ||
resource.path == "/var/run/secrets/analyst-key"
};
forbid (
principal in AgentClass::"data-analyst",
action == Action::"connectNetwork",
resource is NetworkDestination
) unless {
resource.host in DestinationSet::"approved-llm-endpoints" ||
resource.host == "lynx-gateway.internal"
};
The compilation target is BPF LSM hooks, cgroup network hooks, and file access enforcement at the inode boundary. When the agent process steps outside what the policy permits, the kernel refuses the operation – EPERM for the syscall, ECONNREFUSED for the network connection, ENOENT for the file access. The agent gets the same error it would get for any prohibited operation, regardless of what credentials it holds.
The kernel layer naturally enforces file access boundaries, network egress restrictions, syscall whitelisting, library load constraints, and process-spawn limits. The same observation pipeline that feeds enforcement also feeds threat detection.
What the kernel cannot do is reason about why an action is being attempted. It sees a connect() system call. It does not know whether the call is part of a legitimate multi-hop delegation that the gateway already authorized. That context only exists at the L7 layer.
The architectural integration matters as much as either layer in isolation. Cedar policies authored once, evaluated at the gateway, compiled to BPF for kernel enforcement. The compilation is not magical—only the substrate-relevant subset of Cedar policies compiles. The rest stay at the gateway. Either way, security teams write Cedar; the runtime decides which layer is the right one to enforce at.
This integration is what makes the dual-layer approach operationally sustainable. Without one policy language, you end up with two policy stores, two review processes, two engineering teams, and inevitable divergence between what the gateway permits and what the kernel allows. With Cedar at both layers, the policy you wrote is the policy that gets enforced everywhere.
Why single-layer policy isn’t enough for AI agent security
Policy at the gateway alone defends against unauthorized callers and out-of-scope actions. It does not defend against a compromised agent that has a legitimate token and uses it to do things outside its intended behaviour, like read credential files, exfiltrate data through side channels, and escalate privilege inside its runtime.
Policy at the kernel alone defends against process-level misbehaviour. It does not understand identity or delegation, cannot reason about whether a network connection is part of a legitimate multi-hop chain, and has no way to enforce human-in-the-loop approval flows.
Combined, the two layers cover the threat model that either layer alone misses. A compromised agent with a legitimate token can still call out through the gateway, but its local actions are constrained by the kernel sandbox. A misconfigured Cedar policy at the gateway is mitigated by the substrate baseline. A shadow agent that never registered is observed and contained at the kernel.
For Kubernetes-native enterprises building agent infrastructure into regulated workloads, this is the architecture worth building toward. Gateway policy for what agents are allowed to ask for. Kernel policy for what they are allowed to do. Same language for both.
Multi-layer policy is one piece of a larger problem: making AI agent infrastructure accountable end-to-end. Traceability, authorization provenance, identity and ownership, policy-based governance at scale, and human oversight and intervention—they all have to work together.
Read: The Five Pillars of AI Agent Accountability →
Get updates on blog posts, workshops, certification programs, new releases, and more!
Technical Blog
The AI Agent Accountability Gap: Why Network Policies, API Gateways, And RBAC Are Not Enough
By Alister Baroi
on May 27, 2026
In The Five Pillars of AI Agent Accountability: A Diagnostic Framework for Engineering Leaders, we walked through each pillar of AI agent accountability (traceability, authorization provenance, identity and ownership, policy at scale, and human oversight)...
Technical Blog
The Case for VM and Container Consolidation in 2026
By Dillon Barry
on May 26, 2026
Two platforms, two teams, two procurement relationships, all doing one job. There’s a reason it ended up this way. There isn’t a reason it has to stay this way. Ask anyone at a typical enterprise...
Technical Blog
Kubernetes Operational Maturity: Secure and Resilient Cluster Federation with Cluster Mesh
By Veronika Smolik
on May 25, 2026
Practically no one runs a single Kubernetes cluster in production these days. Maybe that’s how it started but data sovereignty requirements, acquisitions, AI initiatives and the need for edge servers, among other considerations, have pulled...
By submitting this form, you acknowledge and agree that Tigera will process your personal information in accordance with the Privacy Policy .
Observability & Troubleshooting
Copyright ©
Tigera, Inc. All rights reserved.

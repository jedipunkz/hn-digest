---
source: "https://techstrong.ai/articles/perplexity-unveils-space-a-secure-sandbox-platform-for-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48949297"
title: "Perplexity Unveils Space, a Secure Sandbox Platform for AI Agents"
article_title: "Perplexity Unveils SPACE, a Secure Sandbox Platform for AI Agents - Techstrong.ai"
author: "CrankyBear"
captured_at: "2026-07-17T17:00:21Z"
capture_tool: "hn-digest"
hn_id: 48949297
score: 1
comments: 0
posted_at: "2026-07-17T16:34:04Z"
tags:
  - hacker-news
  - translated
---

# Perplexity Unveils Space, a Secure Sandbox Platform for AI Agents

- HN: [48949297](https://news.ycombinator.com/item?id=48949297)
- Source: [techstrong.ai](https://techstrong.ai/articles/perplexity-unveils-space-a-secure-sandbox-platform-for-ai-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:34:04Z

## Translation

タイトル: Perplexity が AI エージェント向けの安全なサンドボックス プラットフォームである Space を発表
記事タイトル: Perplexity が AI エージェント向けの安全なサンドボックス プラットフォームである SPACE を発表 - Techstrong.ai
説明: Perplexity は、Perplexity Computer の全機能を安全に利用できるように設計された新しいサンドボックス プラットフォームである SPACE を導入しました。

記事本文:
コンテンツにスキップ
トグルナビゲーション 最新の記事
Perplexity が AI エージェント向けの安全なサンドボックス プラットフォームである SPACE を発表
Perplexity は、高いセキュリティ基準を維持しながら、Perplexity Computer のエージェント AI スタックの全機能を安全に解除できるように設計された新しいサンドボックス プラットフォームである SPACE を導入しました。
メア・カルパ。私はエージェント AI プラットフォームの大ファンではありません。確かにエージェントは強力ですが、そのセキュリティとエージェントをサポートするプラットフォームには疑問があると思います。だからこそ私は、Perplexity Computer が今後、新しいサンドボックス SPACE 上でのみエージェントを実行するという Perplexity の発表が気に入っています。
この文脈でのサンドボックスとは、子供が遊ぶものではなく、実行中のプログラムを隔離して、障害やエクスプロイトが封じ込められ、オペレーティング システムに広がらないようにするメカニズムです。この場合、SPACE は、基盤となるインフラストラクチャや機密資格情報を公開することなく、コードの実行、ファイルの管理、および拡張セッションでの操作を行う必要があるエージェントを使用する長時間実行ツール用に設計されています。
SPACE は、Firecracker microVM サンドボックス内ですべてのコンピューター タスクを実行することでこれを行います。 Firecracker は、Linux カーネルベースの仮想マシン (KVM) を使用して microVM を作成および管理する仮想マシン モニターです。デフォルトのオペレーティング システムのセキュリティを超えるレベルで最小権限を強制します。各 microVM は最小限のデバイス モデルで独自の専用 Linux カーネルを起動し、攻撃対象領域を減らします。これにより、何か問題が発生した場合の爆発範囲が最小限に抑えられます。
悪質なエージェントや失敗したエージェントからさらに保護するために、SPACE のサンドボックスは、コードの実行、ファイルとの対話、またはその他のタスクに必要な間だけ存続します。タスクが完了すると、サンドボックスとその内部にあるすべてのものは破壊されます。
必要な仕事については、

再起動を生き残るために、SPACE は各サンドボックスをセッション内でラップし、一時停止、再開、または複数のサンドボックスへの分岐が可能です。 SPACE は、ローリング スナップショット テクノロジを使用してこれを実行します。これにより、個々のサンドボックスがなくても、コンテキストが作業とともに移動できるようになります。
SPACE は現在、Perplexity Computer の専用ランタイム層です。単一の質問に答えるだけでなく、時間をかけて複数のステップのワークフローを実行するエージェント向けに最適化されています。 SPACE では、各インタラクションをステートレスなリクエストとして扱う代わりに、エージェントがコンテキストを永続化し、中間成果物を維持し、制御された環境内で繰り返されるツール呼び出しを調整できるようにします。この設計は、エージェントが数分から数時間実行され、出力を繰り返し調整する、詳細な調査、複雑なデータ分析、自律的なレポート生成などのユースケースをターゲットとしています。
Perplexity サンドボックスでは、ゼロトラスト アプローチも使用されます。エージェントが実行するコードはすべて信頼できないものとして扱われ、敵対的な可能性があります。これは賢いアプローチです。 SPACE では、各セッションは、共有ホスト上の単なるプロセスではなく、独自の最小限のコンピューティング環境を取得します。このアプローチは、テナント間のリスクとサイドチャネルの露出を軽減し、不正な動作をするコードが他のワークロードに影響を与えるのを防ぐことを目的としています。
SPACE は、エージェントを生の資格情報や直接ネットワーク アクセスから切り離します。エージェントが外部 API またはサービスを呼び出す必要がある場合、宛先ポリシーを適用し、エージェントに代わってスコープ指定されたトークンを挿入する、厳密に制御されたプロキシを通じて呼び出しを行います。エージェントは、基礎となる API キーや長期有効なシークレットを参照することはなく、送信接続は発信元、宛先、目的によって制限される可能性があります。これにより、エージェントができることとできないことを把握しやすくなります。また、オペレータにリスク管理とコンプライアンスのための明確なコントロール プレーンを提供します。

Perplexity によると、SPACE は 3 つの主要な層で構成されています。これらはコントロール プレーンであり、サンドボックスのライフサイクルの調整を担当します。単一の API を通じてすべてのリクエストを受信し、サンドボックスの状態を追跡し、リクエストをどのバックエンドにルーティングするかをリアルタイムで決定します。これにより、タスクに新しいサンドボックスがいつ必要になるかが決まります。また、タスクが完了する時期も決定され、それに対応するサンドボックスを破棄できるようになります。コントロール プレーンはステートレスで中央でホストされているため、同じ API 呼び出しがローカル マシンを含む既存のインフラストラクチャで機能します。
次に、ノードレベルのサービスがあります。これらにより、資格情報がサンドボックスに渡されることがなくなります。代わりに、必要な瞬間にのみサンドボックスの外部から渡されます。エージェントが Google アカウントに一時的にアクセスする必要がある場合、SPACE はサンドボックス内で資格情報を公開せずにサインイン フローを処理できます。アウトバウンドネットワークトラフィックもノードレベルで制御されます。侵害されたエージェントは、許可された範囲外にはアクセスできません。
最後に、サンドボックス内レイヤーは Firecracker microVM で構成されます。 Its SPACE Daemon is the only process that talks to Perplexity’s control plane.開始、一時停止、スナップショット コマンドなどの操作信号を伝達します。
これらすべてのセキュリティ層により、SPACE が非常に遅いのではないかと心配になるかもしれません。同社は、インタラクティブな使用には十分な速度であると主張しています。これは、キャッシュ、Firecracker、Btrfs (B ツリー ファイルシステム) が使用されているためです。 Btrfs は、コピーオンライト ファイルシステムと統合された論理ボリューム管理を組み合わせます。これは、「サンドボックスを迅速に作成および復元できることを意味します。毎回サンドボックスを最初から作成するのではなく、ディスク上に実体化された共通のテンプレートを持つポッドのウォーム プールを保持し、bin によるリクエストを満たします。

テンプレートがすでに一致しているポッドにそれを追加します。」
Perplexity は、SPACE を Computer の背後にあるサンドボックス プラットフォームとして形式化することで、エージェント AI をボルトオン機能ではなく、ファーストクラスのアーキテクチャとみなしていると述べています。
このアプローチの別の見方は、AI エージェントをチャットボットのように扱うのではなく、従来のサービスと同じ厳密さを必要とする分散アプリケーションのように扱うことです。 Perplexity は、これらのアプリケーションに管理されたランタイムを効果的に提供し、低レベルの分離とセキュリティ エンジニアリングを抽象化しながら、エージェントが動作できる領域を厳密に定義したままにします。
これは私にとって有望に思えます。そのランタイムを Perplexity Computer のコア部分にすることで、Computer を次世代のエージェント システムの信頼できるプラットフォームとして位置づけています。
Perplexity Personal Computer: ユーザーに代わって機能するデジタル プロキシ
アマゾン、AI買い物客の紛争で困惑を訴える
Perplexity AI が人間の研究を深く支援するディープリサーチツールを発表
Perplexity、Google Chromeに345億ドルで入札：レポート
ソフトウェアテストとテスト自動化
あなたの組織でのソフトウェア テストまたはテスト自動化への関与を最もよく表しているものは次のうちどれですか? (1つ選択してください) *
ソフトウェアテストまたはテスト自動化を積極的に実行します
ソフトウェア テストまたは QA チームを管理またはリードしています
テストツール、フレームワーク、自動化戦略について決定を下します
実践的なテストと戦略/ツールの決定の両方を行っています
あなたの組織のソフトウェア テストのうち、現在でも手動で実行されている部分はどれくらいありますか? (1つ選択してください) *
ほぼすべて
あなたのチームは、新しいテスト カバレッジを作成するのと比較して、既存の自動テストの維持にどれだけの労力を費やしていますか? (1つ選択してください) *
主に既存のテストの保守
新規の補償よりもメンテナンスの負担が大きい
メンテナンスよりも新しい補償の方が多い
とても

メンテナンスはほとんど必要ありません
AI 支援ソフトウェア開発により、過去 12 か月間でテスト チームへのプレッシャーはどの程度増加しましたか? (1つ選択してください) *
圧力が大幅に上昇
あなたの組織がソフトウェア テストの自動化を拡大し、エージェント テストを導入することを妨げている最大の障害は何ですか? (1つ選択してください) *
自動テストは難しすぎるか維持コストがかかりすぎる
リリースサイクルが早すぎてテストを最新の状態に保つことができない
コーディングやスクリプトに関する専門知識が多すぎる
テストツールとフレームワークが断片化しすぎている
時間、予算、人員の不足
ビジネス価値やROIを証明するのが難しい
AI またはインテリジェントな自動化機能が不十分
今後 12 か月間で貴社のビジネスに最も大きな影響を与える改善はどれですか? (1つ選択してください) *
手動テストの労力を削減
自動テストメンテナンスの削減
より多くのテクノロジーにわたってテスト対象範囲を拡大
テストツールとフレームワークの統合
より多くの非開発者がテストを自動化できるようにする
自動テストの回復力と適応性の向上
AI を利用したソフトウェア テストまたはエージェント ソフトウェア テストに関する組織の現在の見解を最もよく反映しているのは次のうちどれですか? (1つ選択してください) *
私たちは AI を活用したテストアプローチを積極的に評価または導入しています
興味はありますが、実用的な価値と信頼性をまだ評価中です
現在の自動テストアプローチで十分であると考えています
AI 支援テストを導入する前に、既存の自動化を改善することに主に焦点を当てています。
AI を活用したテストを有意義な方法でまだ検討していない
組織の将来のソフトウェア配信戦略にとって、人間の介入を最小限に抑えて、継続的でほぼ自律的なソフトウェア テスト (「消灯」または「暗闇の工場」テスト) を実行できる機能は、どの程度重要ですか? (1つ選択してください) *
重要な戦略

c 優先度
重要だがまだ初期段階
私たちの組織とは関係ありません

## Original Extract

Perplexity has introduced SPACE, a new sandbox platform engineered to safely unlock the full capabilities of Perplexity Computer.

Skip to content
Toggle Navigation Latest Articles
Perplexity Unveils SPACE, a Secure Sandbox Platform for AI Agents
Perplexity has introduced SPACE, a new sandbox platform engineered to safely unlock the full capabilities of Perplexity Computer’s agentic AI stack while maintaining a high bar for security.
Mea culpa. I’m not a huge fan of agentic AI platforms. Sure, agents are powerful, but I find their security and the platforms that support them to be, shall we say, questionable. That’s why I like Perplexity ‘s announcement that going forward, Perplexity Computer will only run its agents on its new sandbox, SPACE .
A sandbox, in this context, is not something your kids play in, but a mechanism that isolates running programs so that failures or exploits are contained and cannot spread to the operating system. In this case, SPACE is designed for long‑running tool‑using agents that need to execute code, manage files, and operate over extended sessions, without exposing underlying infrastructure or sensitive credentials.
SPACE does this by running every Computer task inside a Firecracker microVM sandbox. Firecracker is a virtual machine monitor that uses the Linux Kernel-based Virtual Machine (KVM) to create and manage microVMs. It enforces least privilege at a level that goes beyond default operating system security. Each microVM boots its own dedicated Linux kernel with a minimal device model, reducing the attack surface. This minimizes the blast radius if something goes wrong.
To further protect you from bad agents or agents gone wrong, SPACE’s sandboxes only live as long as they’re needed for running code, interacting with files, or other tasks. When the task finishes, the sandbox and everything inside it are destroyed.
For jobs that need to survive restarts, SPACE wraps each sandbox in a session that can be paused, resumed, or branched into multiple sandboxes. SPACE does this by using a rolling snapshot technology, which enables context to travel with the work, even though no individual sandbox does.
SPACE is now Perplexity Computer ‘s dedicated runtime layer. It’s optimized for agents that don’t just answer a single question but carry out multi‑step workflows over time. Instead of treating each interaction as a stateless request, SPACE allows agents to persist context, maintain intermediate artifacts, and coordinate repeated tool calls within a controlled environment. This design targets use cases such as deep research, complex data analysis, and autonomous report generation, where agents may run for minutes or hours, iteratively refining their output.
The Perplexity sandbox also uses a zero‑trust approach . Any code an agent executes is treated as untrusted and potentially adversarial. This is a smart approach. In SPACE, each session gets its own minimal compute environment rather than just a process on a shared host. This approach is meant to reduce cross‑tenant risk and side‑channel exposure and to keep misbehaving code from impacting other workloads.
SPACE also decouples agents from raw credentials and direct network access. When an agent needs to call external APIs or services, it does so through tightly controlled proxies that enforce destination policies and inject scoped tokens on its behalf. Agents never see underlying API keys or long‑lived secrets, and outbound connections can be constrained by origin, destination, and purpose. This makes it easier to figure out what an agent can and cannot do. It also gives operators a clear control plane for risk management and compliance.
According to Perplexity, SPACE is made up of three key layers. These are the Control Plane, which is responsible for orchestrating the sandbox’s lifecycle. It receives all requests through a single API, tracks the sandbox state, and decides in real time which backend to route them to. It determines when a task needs a new sandbox. It also decides when a task is complete, and its corresponding sandbox can be destroyed. Because the Control Plane is stateless and centrally hosted, the same API call works on any existing infrastructure, including a local machine.
Next, there are Node-level Services. These ensure that credentials never pass into the sandbox. Instead, they’re passed in from outside the sandbox only at the precise moment they’re needed. When an agent needs temporary access to a Google account, SPACE can handle the sign-in flow without exposing credentials within the sandbox. Outbound network traffic is also controlled at the node level. A compromised agent can’t reach anything outside its permitted scope.
Finally, the in-sandbox layer consists of the Firecracker microVM. Its SPACE Daemon is the only process that talks to Perplexity’s control plane. It carries operational signals like start, pause, and snapshot commands.
All these security layers may make you worry that SPACE could be really slow. The company claims it’s fast enough for interactive use. That’s because of its use of caching, Firecracker, and the Btrfs (B-tree filesystem) . Btrfs combines a copy-on-write filesystem with integrated logical volume management. This means “ sandboxes can be created and restored quickly. Instead of creating a sandbox from scratch each time, we keep a warm pool of pods that already have common templates materialized on disk, and satisfy a request by binding it to a pod whose template already matches.”
By formalizing SPACE as the sandbox platform behind Computer, Perplexity is saying that it views agentic AI as a first‑class architecture rather than a bolt-on feature.
Another way of looking at this approach is to treat AI agents less like chatbots and more like distributed applications that require the same rigor as traditional services. Perplexity is effectively offering a managed runtime for those applications, abstracting away low‑level isolation and security engineering while keeping the surface area where agents can act tightly defined.
This sounds promising to me. By making that runtime a core part of Perplexity Computer, it’s positioning Computer as a trustworthy platform for the next generation of agentic systems.
Perplexity Personal Computer: A Digital Proxy That Works on a User’s Behalf
Amazon Sues Perplexity Over AI Shopper Dispute
Perplexity AI Launches a Deep Research Tool to Help Humans Research, Deeply
Perplexity Bids $34.5 Billion for Google Chrome: Reports
Software Testing and Test Automation
Which of the following best describes your involvement in software testing or test automation at your organization? (Select one) *
I actively perform software testing or test automation
I manage or lead software testing or QA teams
I make decisions about testing tools, frameworks, or automation strategy
I do both hands-on testing and make strategic/tool decisions
How much of your organization's software testing is still performed manually today? (Select one) *
Almost all
How much effort does your team spend maintaining existing automated tests compared to creating new test coverage? (Select one) *
Mostly maintaining existing tests
More maintenance than new coverage
More new coverage than maintenance
Very little maintenance required
How much has AI-assisted software development increased the pressure on your testing teams over the past 12 months? (Select one) *
Significantly increased pressure
What is the biggest obstacle preventing your organization from expanding software test automation and adopting agentic testing? (Select one) *
Automated tests are too difficult or costly to maintain
Release cycles move too quickly to keep tests current
Too much coding or scripting expertise is required
Testing tools and frameworks are too fragmented
Lack of time, budget, or staffing
Difficulty proving business value or ROI
Insufficient AI or intelligent automation capabilities
Which improvement would create the greatest business impact for your organization over the next 12 months? (Select one) *
Reducing manual testing effort
Reducing automated test maintenance
Expanding test coverage across more technologies
Consolidating testing tools and frameworks
Enabling more non-developers to automate testing
Improving the resilience and adaptability of automated tests
Which statement best reflects your organization's current perspective on AI-powered or agentic software testing? (Select one) *
We are actively evaluating or adopting AI-powered testing approaches
We are interested, but still assessing practical value and trust
We believe current automated testing approaches are sufficient
We are primarily focused on improving existing automation before adopting AI-assisted testing
We have not yet explored AI-powered testing in a meaningful way
How important is the ability to execute continuous, largely autonomous software testing with minimal human intervention ("lights-out" or "dark factory" testing) to your organization's future software delivery strategy? (Select one) *
Critical strategic priority
Important, but still early-stage
Not relevant to our organization

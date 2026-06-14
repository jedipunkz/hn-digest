---
source: "https://claude.com/blog/building-with-claude-managed-agents"
hn_url: "https://news.ycombinator.com/item?id=48527164"
title: "The evolution of agentic surfaces: building with Claude Managed Agents"
article_title: "The evolution of agentic surfaces: building with Claude Managed Agents | Claude"
author: "gmays"
captured_at: "2026-06-14T15:38:25Z"
capture_tool: "hn-digest"
hn_id: 48527164
score: 3
comments: 0
posted_at: "2026-06-14T13:46:27Z"
tags:
  - hacker-news
  - translated
---

# The evolution of agentic surfaces: building with Claude Managed Agents

- HN: [48527164](https://news.ycombinator.com/item?id=48527164)
- Source: [claude.com](https://claude.com/blog/building-with-claude-managed-agents)
- Score: 3
- Comments: 0
- Posted: 2026-06-14T13:46:27Z

## Translation

タイトル: エージェント サーフェスの進化: クロード管理エージェントによる構築
記事のタイトル: エージェント サーフェスの進化: Claude 管理エージェントを使用した構築 |クロード
説明: Claude 管理エージェントを使用すると、チームは実稼働環境でエージェントを大規模に構築し、確実にデプロイできます。チームがそれを使用する理由と方法は次のとおりです。

記事本文:
エージェント サーフェスの進化: Claude 管理エージェントを使用した構築 |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
エージェント サーフェスの進化: Claude 管理エージェントを使用した構築
エージェント サーフェスの進化: Claude 管理エージェントを使用した構築
モデル インテリジェンスとエージェント ハーネスが進化するにつれて、Claude マネージド エージェントを使用すると、チームは実稼働環境で大規模なエージェントを確実に構築して展開できるようになります。チームがそれを使用する理由と方法は次のとおりです。
共有 リンクをコピー https://claude.com/blog/building-with-claude-managed-agents
エージェントを運用環境に導入するには、適切なプロンプトだけでは十分ではありません。エージェントには、作成したコードを実行する場所、データにアクセスするための資格情報、監視可能なセッション、および使用量に応じて拡張するインフラストラクチャが必要です。 Applied AI チームでは、製品、研究、Claude を基盤とする顧客の交差点で働いていますが、同じパターンが繰り返し見られます。つまり、インフラストラクチャがプロトタイプと運用エージェントを分けるものであるということです。チームは、セキュリティ、状態管理、権限付与、ハーネス調整に開発サイクルを費やしてしまうことがよくあります。
クロード マネージド エージェント は、実稼働グレードのエージェントを構築および展開するためのコンポーザブル API スイートであり、パフォーマンス向けに調整されたエージェント ハーネスと実稼働インフラストラクチャを組み合わせることで、チームがプロトタイプから立ち上げまで数か月ではなく数日で完了できるようにします。この中で

この投稿では、Anthropic のエージェント構成要素の進化、Claude 管理エージェントを構築した理由、そして今日のチームが運用環境でそれをどのように使用しているかについて説明します。
エージェントアーキテクチャの進化
2023 年に Claude を開発者に公開したとき、API はトークンをインしてトークンをアウトするという意図的にシンプルなものでした。あなたがプロンプトを送信し、クロードが完了を返し、ハーネスと基礎となるインフラストラクチャを構築しました。
API は長年にわたって着実に機能が充実してきましたが、その根底にある契約は決して変更されませんでした。1 つのリクエスト、1 つのモデルターンで、アプリケーションが次に何が起こるかを決定します。長い間、それだけで十分でした。ドキュメントの要約、サポート チケットの分類、テキスト ブロックの書き換えなど、このような作業は 1 回で十分に完了します。
しかし、時間が経つにつれて、人々が引き継ぎたいタスクが適合しなくなりました。彼らは、クロードにタスクを最後まで実行し、何かを調べ、それに基づいて行動し、何が変化したかを確認し、次に何をすべきかを決定することを望んでいました。そして彼らは、コードベース、内部 Wiki、チケット発行システムなど、作業がすでに実行されているシステムでそれを動作させたいと考えていました。
API を使用すると、クロードをエージェントに変えるということは、モデルに何をすべきかを尋ね、ツールを実行し、結果をフィードバックして繰り返すという独自のループを構築することを意味しました。あなたはエージェント スキャフォールディングの構築とデプロイを担当しましたが、モデルの進化に応じて調整が必要になる可能性があります。完全なカスタマイズが必要なエージェントにとって、このアプローチは理にかなっています。予測可能で複雑さが少ないエージェント ワークロードの場合、モデルや製品の進化に応じてハーネスを最適化するのは面倒なものになりました。
クロードがコードベースと直接対話できるようにする、2025 年にリリースされたエージェント コーディング ツールである Claude Code には、そのハーネスの独自バージョン、つまりループ、ツール実行、サブエージェント、コンテキスト管理、および効果的なエージェントにするための豊富な機能が含まれていました。

開発者は当然のことながら、さまざまなドメインにわたる自社のエージェント用に同様のハーネス機構を必要としました。
チームが Claude Code ハーネス上にエージェントを構築できるようにするために、Claude Agent SDK をリリースしました。 Claude Agent SDK は、開発者が独自のループを維持する代わりに、Claude コードを実行するのと同じ機械上に独自のエージェントを構築するためのツールを提供します。多くのチームにとって、これはエージェントが実用的なものになったときです。ハーネスは、インフラストラクチャのプリミティブを備えたクロード用にすでに調整されて到着し、クロード コードと同様に改善され続けました。
ただし、ハーネスを使用していても、実稼働環境にエージェントを展開することは、次のような理由から困難になる場合があります。
ホスティングとスケーリング。エージェントはどこで実行されますか? 複数時間のタスクに対してプロセスはどのくらいの時間存続できますか? 使用量が増加した場合はどのようにスケールしますか?
セッション管理。エージェントの履歴と進捗状況はどこにありますか?実行は中断を乗り越えて何の妨げもなく再開できるでしょうか?以前のセッションに戻って何が起こったのか調べてもらえますか?
ファイルシステム管理。実際の作業を行うということは、コードの編集、ファイルの作成、出力の構築などの成果物を作成することを意味します。エージェントはどこで動作するワークスペースを取得しますか?また、実行と実行の間にそのワークスペースはどうなりますか?
実行の分離。クロードが書いたコードはどこかで実行する必要があります。間違っている場合の爆発範囲はどのくらいですか?また、本番環境では実際にどの境界を信頼しますか?
資格。エージェントはシステムにアクセスする必要があります。生成するコードに機密情報を公開せずに、どのようにしてそのアクセスを取得するのでしょうか?
可観測性。エージェントが 1 時間自律的に作業し、何か驚くべきことを行ったとき、そのステップをすべて再現できますか?
Agent SDK を使用すると、前述の運用インフラストラクチャの多くの要素が Claude Code の機構を通じて提供されます。エージェントは実際に動作するファイルシステムを取得します。

、セッション状態はローカルまたは外部ストレージに保持され、可観測性は OpenTelemetry を介して、すでに実行している監視スタックにエクスポートできます。
しかし、ローカル開発から本番環境に移行するエージェントを構築するチームが増えるにつれ、管理されたインフラストラクチャを使用してエージェントを大規模にデプロイする方法が必要になりました。そして、モデルとその周囲のハーネスがより高度になり、実行時間が長くなり、より多くのコードが実行され、より多くのシステムにアクセスし、より多くのアクションが実行されるようになり、スケーリング、セキュリティ、サンドボックス化がより困難になりました。
これらのハードルのいくつかは、一般的なアーキテクチャ上の選択から生じています。エージェント ハーネスは、多くの場合、それが動作するファイル システムと同じコンテナ内で実行されます。クロードが考える前にコンテナーを起動する (起動コストを支払う) 必要があり、エージェントとコード実行は資格情報のすぐ隣に存在し、コンテナーが終了すると、実行も一緒に終了します。
マネージド エージェントは、脳を手から切り離すことでこれらの問題を解決します。クロードを呼び出すハーネスは、コードが実行されるサンドボックスとは別に実行され、セッション (すべてのモデル呼び出し、ツール呼び出し、結果の追加専用ログ) がこの 2 つを接続します。クロードは、コンテナーが存在する前に推論を開始でき、サンドボックスは資格情報から遠く離れた場所にあり、いつでもセッションから実行全体を再構築できます。
Claude 管理エージェントを使用する場合とその理由
管理対象エージェントを使用して構築する場合、ユーザーはタスク、ツール、ガードレールを定義します。Anthropic はインフラストラクチャ上でエージェントを実行し、その下でエージェント ループを処理します。エージェントにツールを呼び出すための実行環境を与える方法、何かが失敗した場合の回復方法、マルチエージェント オーケストレーションなどです。
ハーネスがモデルのインテリジェンスに合わせて進化しないと、エージェントは故障します。クロード・ソネットについて 4

.5 では、エージェントはコンテキストの終わりに近づくと急いで終了し、残されたスペースを使用せずに作業を短縮します。これは「コンテキスト不安」と呼ばれるパターンです。私たちの修正は、ハーネスにコンテキスト リセットを追加し、クロードが限界近くで一貫性を保つのに助けが必要であるという仮定を組み込むことでした。その前提は次のモデルには生き残れませんでした。 Claude Opus 4.5 では、この動作はなくなり、追加したリセットは単なるオーバーヘッドでした。
ほとんどの組織にとって、ハーネスの保守はオーバーヘッドであり、製品の差別化にはつながりません。ハーネスは特定のモデルの動作に合わせて調整する必要があります。圧縮、ツールの実行、キャッシュなどのプリミティブは、クロードでは他のモデルとは動作が異なります。 Claude マネージド エージェントを使用すると、ハーネスがモデルとともに進化し、チームがエージェントを差別化するもの、つまりコンテキスト管理とドメインの専門知識に集中できるようになります。
開発者が効果的なエージェントを構築するために必要なコンテキストとツールを構成できるようにするために、管理対象エージェントは 3 つの主要なリソース (エージェント、環境、セッション) を中心に構築されています。エージェントは、モデル、プロンプト、ツールのセット、およびそれらの周囲のガードレールといった構成です。環境とは、エージェントが実行される実行コンテキストです。サンドボックス コンテナー、そのネットワーク ルール、およびその中にプレインストールされ、クラウドまたはお客様が管理するインフラストラクチャ上でホストされているパッケージです。各実行は セッション であり、エージェントと環境をペアにして、独自の分離されたサンドボックス インスタンスを取得します。セッションは完全なイベント履歴、サンドボックス状態、および出力をサーバー側に保持するため、長時間実行されている作業を一時停止したり、きれいに再開したり、事後段階的に追跡したりできます。管理対象エージェントを使用すると、エージェントと環境を一度定義すれば、ワークロードの増大に応じて同じ構成に対して多数のセッションを実行できます。
生産用の建物

管理対象エージェントでスケールする
Applied AI では、Anthropic 内とお客様のシステム全体、コーディング、財務、サポート、法務、その他多数のドメインにわたって、エージェントがプロトタイプから本番環境に移行する様子が見られます。これにより、デモと運用準備が整ったエージェントの違い、およびチームがよく行き詰まる場所が明確になります。
以下では、Claude マネージド エージェントのようなマネージド サービスを構築する最も一般的な理由を説明します。
1. 資格情報はサンドボックス外に保管されます。すべてが 1 つのコンテナー内で実行される場合、Claude が生成するコードは資格情報のすぐ隣に配置されるため、プロンプト インジェクションによってモデルが独自の環境を読み取るように誘導され、モデルがトークンを漏洩する可能性があります。同じコンテナ内に堅牢なガードレールを設定することでこれを防ぐことができますが、アーキテクチャを分離すると、資格情報を完全にサンドボックスに入れないようにすることで、より安全なアプローチが可能になります。 MCP、CLI、GitHub リポジトリなどのツールのトークンは別のボールトに存在し、プロキシがそれらをフェッチし、要求に応じてのみ復号化します。管理対象エージェントは、すぐに使用できる資格情報を処理する Vault を提供するため、独自のシークレット ストアを実行したり、呼び出しごとにトークンを送信したり、エージェントがどのエンド ユーザーに代わって行動したかを追跡したりする必要はありません。ボールトの資格情報は保管前にエンベロープ暗号化で保護され、取得には検証のために署名されたリクエスト トークンが必要です。
2. サンドボックスのオーバーヘッドを排除することでレイテンシーを低減します。ユーザーはクロードの応答を待っているときを敏感に感じるため、レイテンシは多くの企業チームにとって最優先の指標です。管理対象エージェント アーキテクチャがなければ、エージェントが考えるだけでツールを実行しないセッションであっても、セッションごとにコンテナを起動する必要があります。そのセットアップ時間は無駄であり、ユーザーはそれを遅延として感じます。

最初の応答。管理対象エージェントを使用すると、環境が並行してスピンアップされ、ツールを実行しないセッションはコンテナーを完全にスキップしながら、クロードはすぐに推論を開始します。つまり、ユーザーはコンテナーの起動を待たずに最初のトークンを確認でき、エージェントが何かを実行する必要があるときまでに環境の準備が整います。私たちのテストでは、最初のトークンまでの時間が中央値の場合 (p50) で約 60%、最も遅い場合 (p95) で 90% 以上短縮されました。
3. セッション管理、可観測性、およびメモリを可能にする、信頼性の高い永続的なセッション。管理対象エージェントは、リクエスト/レスポンスではなく、イベントの観点から考えます。セッションは、継続的なイベントのストリームです。すべてのモデル呼び出し、ツール呼び出し、および結果は、エージェントを実行しているプロセスの外にあるログに追加されます。このアーキテクチャを使用すると、エージェントの動作中にイベントがストリーミングされるとリアルタイムの更新が得られ、管理するデータベースやセーブポイントを必要とせずに、後で任意のセッションを再開できます。セッションを削除しない限り、対話間の履歴は保存され、セッションがアイドル状態になるとコンテナーにチェックポイントが設定されるため、一時停止した場所からきれいに再開できます。そして、実行全体がすでにイベントの記録であるため、可観測性と記憶が付属しています。

[切り捨てられた]

## Original Extract

Claude Managed Agents allows teams to build and deploy agents in production environments reliably at scale. Here’s why and how teams are using it.

The evolution of agentic surfaces: building with Claude Managed Agents | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
The evolution of agentic surfaces: building with Claude Managed Agents
The evolution of agentic surfaces: building with Claude Managed Agents
As model intelligence and agentic harnesses evolve, Claude Managed Agents allows teams to build and deploy agents in production environments reliably at scale. Here’s why and how teams are using it.
Share Copy link https://claude.com/blog/building-with-claude-managed-agents
Getting an agent into production takes more than a good prompt. The agent needs somewhere to run the code it writes, credentials to reach your data, observable sessions, and infrastructure that scales with usage. On the Applied AI team, we work at the intersection of product, research, and the customers building on Claude—and we see the same pattern repeatedly: infrastructure is what separates a prototype from a production agent. All too often, teams burn development cycles on security, state management, permissioning, and harness tuning.
Claude Managed Agents , our suite of composable APIs for building and deploying production-grade agents, pairs an agent harness tuned for performance with production infrastructure, allowing teams to go from prototype to launch in days rather than months. In this post, we'll cover the evolution of Anthropic’s agentic building blocks, why we built Claude Managed Agents, and how teams are using it in production today.
Evolving the agent architecture
When we opened up Claude to developers in 2023, the API was deliberately simple: tokens in, tokens out. You sent a prompt, Claude returned a completion, and you built the harness and underlying infrastructure.
The API grew steadily richer over the years, but the contract underneath never changed: one request, one model turn, and your application decides what happens next. For a long time, that was enough. Summarizing a document, classifying a support ticket, rewriting a block of text—the kind of work that fits comfortably in a single turn.
Over time, however, the tasks people wanted to hand off stopped fitting. They wanted Claude to carry a task all the way through, look something up, act on it, see what changed, and decide what to do next. And they wanted it to operate in the systems their work already ran on, like a codebase, internal wiki, or ticketing system.
With the API, turning Claude into an agent meant building your own loop: ask the model what to do, run the tool, feed the result back, and repeat. You were responsible for building and deploying the agent scaffolding, which may need tuning as models evolve. For agents that require full customization, this approach makes sense. For agentic workloads that are more predictable and less complex, optimizing harnesses as models and products evolved became tedious.
Claude Code , the agentic coding tool we launched in 2025 that lets Claude interact directly with your codebase, contained our own version of that harness: the loop, tool execution, subagents, context management, and rich capabilities that made it an effective agent. Developers naturally wanted similar harness machinery for their own agents across various domains.
To enable teams to build agents on top of the Claude Code harness, we released Claude Agent SDK . Claude Agent SDK gives developers tools to build their own agents on the same machinery that runs Claude Code instead of maintaining a homegrown loop. For a lot of teams, this is when agents became practical: the harness arrived already tuned for Claude with infrastructure primitives and it kept improving as Claude Code did.
Even with a harness, though, deploying agents in production environments can be challenging for several reasons:
Hosting and scaling. Where does the agent run, how long can a process stay alive for a multi-hour task, and what scales it when usage grows?
Session management. Where does an agent's history and progress live? Can a run survive an interruption and resume unencumbered? Can you go back and inspect what happened in previous sessions?
Filesystem management. Doing real work means producing artifacts: editing code, writing files, building outputs. Where does the agent get a workspace to act on, and what happens to that workspace between runs?
Execution isolation. The code Claude writes has to execute somewhere. What's the blast radius if it's wrong, and what boundary would you actually trust in production?
Credentials. The agent needs access to your systems. How does it get that access without exposing proprietary information to the code it generates?
Observability. When an agent works autonomously for an hour and does something surprising, can you reconstruct every step it took?
With the Agent SDK, many elements of the aforementioned production infrastructure are provided through Claude Code’s machinery. The agent gets a real filesystem to work in, session state is persisted locally or on external storage, and observability is exportable through OpenTelemetry into whatever monitoring stack you already run.
However, as teams increasingly built agents that moved out of local development into production, they needed a way to deploy them at scale and with managed infrastructure. And as models and their surrounding harnesses become more advanced–running longer, executing more code, touching more systems, and taking more actions– scaling, security, and sandboxing became more challenging.
Several of these hurdles stem from a common architectural choice: agent harnesses often run inside the same container as the filesystem it works on. A container has to spin up (paying a startup cost) before Claude can think, the agent along with code execution lives right next to your credentials, and when the container dies, the run dies with it.
Managed Agents solves these problems by decoupling the brain from the hands . The harness that calls Claude runs separately from the sandbox where code executes, and the session–an append-only log of every model call, tool call, and result–connects the two. Claude can start reasoning before any container exists, the sandbox stays far away from your credentials, and a whole run can be reconstructed from its session at any point.
When and why to use Claude Managed Agents
When building with Managed Agents, users define the task, the tools, and the guardrails, and Anthropic runs the agent on our infrastructure and handles the agentic loop underneath: how to give an agent an execution environment to call tools, how to recover when something fails, multi-agent orchestration, and more.
When the harness doesn’t evolve alongside model intelligence, the agent breaks down . On Claude Sonnet 4.5, an agent would rush to finish as it neared the end of its context, cutting work short rather than using the room it had left—a pattern called "context anxiety." Our fix was to add context resets to the harness, baking in an assumption that Claude needed help staying coherent near the limit. That assumption didn't survive the next model. On Claude Opus 4.5, the behavior was gone, and the resets we'd added were just overhead.
For most organizations, maintaining a harness is overhead that doesn't differentiate their product. Harnesses have to be tuned for certain model behaviors; primitives like compaction, tool execution, and caching works differently on Claude than other models. With Claude Managed Agents, the harness evolves alongside the model, allowing teams to focus on what will differentiate their agents: context management and domain expertise.
To enable developers to configure the context and tools necessary to build effective agents, Managed Agents is built around three primary resources: agents, environments, and sessions. An agent is a configuration: a model, a prompt, a set of tools, and the guardrails around them. An environment is the execution context the agent runs in: the sandbox container, its networking rules, and the packages pre-installed in it, hosted on our cloud or on infrastructure you control. Each run is a session , which pairs an agent with an environment and gets its own isolated sandbox instance. Sessions persist their full event history, sandbox state, and outputs server-side, so long-running work can pause, resume cleanly, and be traced step by step after the fact. With Managed Agents, you can define an agent and an environment once, then run many sessions against the same configuration as your workload grows.
Building for production and scale on Managed Agents
Within Applied AI, we see agents go from prototype to production both inside Anthropic and across our customers’ systems, across coding, finance, support, legal, and a dozen other domains. This gives us a clear view of what separates a demo from a production-ready agent and where teams often get stuck.
Below, we share the most common reasons to build on a managed service like Claude Managed Agents:
1. Credentials are kept out of the sandbox. When everything runs in one container, the code Claude generates sits right next to your credentials, so prompt injections could lead the model to leak a token by convincing the model to read its own environment. We can protect against this by setting up robust guardrails within the same container, but decoupling the architecture enables a much more secure approach by keeping credentials out of the sandbox entirely. Tokens for tools like MCPs, CLIs, and GitHub repos live in a separate vault, and a proxy fetches them and decrypts them only on demand. Managed Agents provides Vaults that handle credentials out-of-the-box, so you don’t need to run your own secret store, transmit tokens on every call, or lose track of which end user an agent acted on behalf of. Vault credentials are protected with envelope encryption before storage, and retrieval requires a signed request token for verification.
2. Lower latency from eliminated sandbox overhead. Latency is a metric that is top-of-mind for many enterprise teams, since users acutely feel when they’re waiting for Claude to respond. Without the Managed Agents architecture, a container has to be spun up for every session, even the ones where the agent only needs to think and never runs a tool. That setup time is wasted, and the user feels it as a delay before the first response. With Managed Agents, Claude begins reasoning immediately while the environment spins up in parallel, and sessions that never run a tool skip the container entirely. This means the user sees the first token without waiting on container startup, and the environment is ready by the time the agent needs to run something. In our testing, that cut the time-to-first-token by roughly 60% in the median case (p50) and by over 90% in the slowest cases (p95).
3. Reliable, persistent sessions that enable session management, observability, and memory. Instead of request/response, Managed Agents thinks in terms of events. A session is an ongoing stream of events: every model call, tool call, and result, are appended to a log that lives outside the process running the agent. With this architecture, you get real-time updates as events stream in while the agent works, and you can resume any session later with no database or save-points to manage. History is preserved between interactions unless you delete the session, and when a session goes idle its container is checkpointed so you can pick up cleanly from where it paused. And because the whole run is already a record of events, observability and memory come with it: t

[truncated]

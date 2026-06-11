---
source: "https://www.serval.com/serval-news/do-the-hard-things-always"
hn_url: "https://news.ycombinator.com/item?id=48490005"
title: "We chose code as the source of truth for AI-generated workflows"
article_title: "Do the hard things, always- Serval - AI Agents for IT"
author: "emot"
captured_at: "2026-06-11T13:26:35Z"
capture_tool: "hn-digest"
hn_id: 48490005
score: 1
comments: 0
posted_at: "2026-06-11T13:22:37Z"
tags:
  - hacker-news
  - translated
---

# We chose code as the source of truth for AI-generated workflows

- HN: [48490005](https://news.ycombinator.com/item?id=48490005)
- Source: [www.serval.com](https://www.serval.com/serval-news/do-the-hard-things-always)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T13:22:37Z

## Translation

タイトル: AI が生成するワークフローの信頼できる情報源としてコードを選択しました
記事タイトル: 難しいことはいつもやってみよう - サーバル - IT 向け AI エージェント
説明: サーバールは、企業全体で手作業による運用作業を排除するためのユニバーサル自動化プラットフォームをどのように構築しているか。

記事本文:
サーバールは、企業全体で手作業による運用作業を排除するためのユニバーサル自動化プラットフォームをどのように構築しているか。
従業員は「Priya にステージング データベースへのアクセスを許可してもらえますか?」と尋ねます。
ほとんどの企業では、この単純な要求はまったく簡単ではありません。誰かが、Priya がどのチームに所属しているか、リクエストが許可されているかどうか、誰が承認する必要があるか、アクセスが定義されている場所、更新方法、コード変更が必要かどうか、リクエスト者に通知する方法、監査証跡を残す方法を把握する必要があります。
従業員は、その作業が Slack、Okta、GitHub、Terraform、AWS、Jira、または内部管理ツールのいずれを通じて行われるかを気にしません。彼らの観点から見ると、やるべきことが 1 つありました。
そう思わせるのがサーバルの仕事。
誰かが自然言語で何かを尋ねます。サーバルはその意味を理解し、適切なツールを見つけ、関連するシステム全体で自動化を実行し、必要に応じて人間を呼び込み、何が起こったかを記録し、作業に戻せるようにします。
ほとんどの企業は、アクセス要求、オンボーディング、ハードウェア要求、インフラストラクチャの変更、セキュリティレビュー、承認、CRM 更新、財務業務、内部サポートなど、どのワークフローを自動化する必要があるかをすでに知っています。自動化は通常、作成するにはコストがかかりすぎ、維持するには脆弱すぎるため、手動のままです。プロセスが 5 つのシステムにまたがり、チームによって異なり、誰かの頭の中にあるポリシーに依存し、月に数回しか起こらない場合、通常は人間の責任が残ります。
Serval は、LLM が自動化のコスト曲線を変えるという単純なアイデアに基づいて構築されています。しかし、より安価な自動化は、安全に実行できる場合にのみ重要です。出力は構造化され、権限があり、レビュー可能で、耐久性があり、デバッグ可能で、監査可能である必要があります。
目標は具体的です。会社の業務を 100% 自動化することです。

簡単な部分だけではありません。ほとんどの自動化ツールは、一般的な大量のワークフローをカバーしており、珍しいことはすべて人任せにしています。しかし、ハードで稀なシステム間リクエストのロングテールは、実際に手作業のほとんどが発生する場所です。明らかなスライスを自動化するだけでは十分ではありません。目標は、厄介なロングテール、つまり以前は自動化を正当化するには常に稀すぎたり、特殊すぎたり、コストが高すぎたりしたワークフローに到達することです。
私たちは、自動化を実行する従業員と自動化を作成する管理者という 2 つの異なるエクスペリエンスを設計する必要がありました。従業員は、「Priya にステージング データベースへのアクセスを許可してもらえますか?」と尋ねるかもしれません。管理者は、これを次のように定義できます。エンジニアリング メンバーシップを確認し、マネージャーに承認を求め、IAM バインディングを追加する Terraform PR を開き、レビューを待ち、パイプラインが変更を適用したときに要求者に通知します。
その意味では、サーバルは単に作業を自動化するだけではありません。オートメーションの作成自体が自動化されており、ワークフローの作成を迅速化するだけでなく、ワークフローの定義を迅速化します。
これをうまくやるのは難しいです。しかし、Serval では、難しい技術的な選択によって最高のエンド ユーザー エクスペリエンスに近づくことができるのであれば、迂回するのではなく、真剣に選択すべきであるという強い信念を持っています。私たちの初期の決定のいくつかは、その哲学に基づいたものでした。
決定 1: レイヤーではなくプラットフォームを構築する
明らかに近道は、既存のチケット発行システムまたは IT サービス管理 (ITSM) システムの上に薄い層として Serval を構築することでした。チケットを読み取り、LLM を使用してチケットを分類または要約し、ダウンストリームの自動化を呼び出し、結果をポストバックします。そうすればもっと速く、おそらく初期の良いデモが生成されたでしょう。
しかし、基礎となるアーキテクチャは変更されませんでした。チケット発行システムは引き続きリクエストを所有します。アクセス管理はまだ多少は存続するだろう

here else.承認は引き続き別のツールで行われます。実行状態は依然としてログ、Slack メッセージ、Webhook、およびダウンストリーム システムに分散しています。 LLM を導入すると、既存のプロセスが少しスマートになったように感じられますが、会社は依然として同じ断片化されたスタックを通じて運営されることになります。
それが私たちが直接解決したいと思っていた問題です。ユーザーは、どのシステムがリクエストを所有しているかを知る必要はありません。管理者は、世界の共有モデルに基づいて設計されていないツールを接続する必要はありません。また、失敗したワークフローをデバッグするエンジニアは、5 つの異なる場所から何が起こったかを再構築する必要はありません。
iPhone が優れた製品であるのは、主に Apple がハードウェアとソフトウェアを一緒に所有しているからです。エクスペリエンスは、緩やかに接続されたレイヤーから組み立てられるものではありません。 it is cohesive because the pieces are designed to fit.会社運営でも同じことがわかります。記録システムと行動システムは統一される必要があります。 If one tool stores the request and another tool executes it, the user experience, admin experience, and audit trail all fragment.そこで私たちは、表面積全体を所有し、従来のシステムが行っていたことを置き換え、断片化されたシステムで実現できるものよりもはるかに優れたものを構築することに着手しました。
Take the Priya request.層が薄いと、よりスマートなチケットが得られます。ただし、Terraform ファイルを見つけて PR を開き、レビューを待って要求者に通知する人が必要になります。この要求は依然として 5 つのシステムに影響を与えることになります。 Serval would just be making the first step feel easier.
決定 2: 真実の情報源としてのコード
プラットフォームを所有するということは、プラットフォーム内で実際にどのようなワークフローが行われるかを決定する必要があることを意味しました。
これは当時は明らかではありませんでした。 A lot of automation products are built around forms, visual workflow builders, decision trees, or

グラフ。これらは単純なケースでは機能しますが、ワークフローが複雑になると壁にぶつかる傾向があります。つまり、この API を呼び出し、このポリシーで分岐し、この承認を待ち、このレコードを更新し、この操作を再試行し、このプル リクエストを開き、後で全体を理解できるようにします。
2024 年初頭にこの決定を下したとき、LLM がコードの作成に長けているため、コードの作成コストは大幅に安くなるだろうと私たちは考えていました。それによってオーサリング モデルが変わりました。ワークフローが LLM によって作成されることが多い場合は、モデルが生成するのが得意で、人間がレビューするのが得意な形式でワークフローを表現する必要があります。
命令型コードがその形式です。標準化された制御フロー、タイプ、レビュー、バージョン管理、テストを提供し、エンジニアが何が起こるかを正確に理解する方法を提供します。コードの作成が安価になると、デプロイ、実行、評価がプラットフォーム上でより困難な問題になります。つまり、安全で管理され、監査された環境でそのコードを実行することが容易になります。そこで私たちは、確定的なワークフローの信頼できるソースとして命令型 TypeScript コードを中心に Serval を構築し、そのコードの実行と反復を可能な限りシンプルかつ安全に行えるようにしました。
コードを使用すると、事前定義されたカタログに収まらないワークフローをターゲットにすることもできます。固定された一連のアクションやビジュアル ビルダーは一般的なリクエストにはうまく機能しますが、ワークフローが 1 つの会社に固有の場合には機能しません。つまり、このポリシー ファイルを読み、この承認ルールに基づいて分岐し、この内部システムを更新し、この正確なプル リクエストを開き、このチームが期待する形式で結果を説明する必要があります。
これらはまさに、通常は手動のままであるワークフローです。汎用プログラミング モデルは、誰かが最初に特殊目的のブロックを構築するのを待たずに、サーバルにそれらを表現する方法を提供します。
すべての手順を行う必要があるという意味ではありません

コードとして表現します。また、スキルと呼ばれる自然言語の標準操作手順もサポートしています。これは、ワークフローの一部が本質的に判断に基づいているためです。つまり、ユーザーが求めていることを解釈する、人間の言語で書かれたポリシーを適用する、2 つのリクエストが同等かどうかを判断する、またはリクエストが拒否された理由を説明するなどです。
どちらの形式も重要ですが、果たす役割は異なります。自然言語は、あいまいな入力について推論し、ユーザーの意図を解釈し、判断の要求を処理するのに役立ちます。コードは、実際に実行される自動化の信頼できる情報源であり、許可、バージョン管理、デバッグ、監査が可能な決定的で副作用があり、レビュー可能な成果物です。
また、ワークフローの記述方法だけでなく、ワークフローの実行方法も決まりました。ほとんどのワークフロー システムは、グラフ形式のものを公開します。つまり、タスクを定義し、それらをエッジで接続し、エンジンにグラフを実行させます。これは、データ パイプラインとバッチ ジョブにとって適切な抽象化です。これは、AI によって作成されたワークフロー システムに求めていた開発者エクスペリエンスではありません。
私たちは、開発者エクスペリエンスを完全に制御できるようにして、Serval が作成者を支援するものを、標準の命令型 TypeScript とできる限り同じように見えるようにしたいと考えました。
これが重要な点です。大規模な言語モデルはこの種のコードの作成に適しており、命令型コードはロジックを表現する非常に効率的な方法です。条件分岐、ループ、関数、型、ライブラリを入手でき、すべての構文プログラミング言語は複雑な動作をコンパクトに表現するために進化しました。必要に応じて、管理者は動作を確認できます。副作用をもたらす作業は明示的で、型指定されており、検査可能です。
その下では、ランタイムは依然として長時間実行されるワークフロー エンジンのように動作する必要があります。実際の企業のワークフローは、承認を待ち、失敗した統合を再試行し、数日後に再開し、何をデバッグするのに十分な状態を残します。

それは起こりました。私たちは 2 つの理由から独自の耐久性のある実行エンジンを構築しました。それは、作成されたワークフローを標準の TypeScript にできる限り近づけることと、開発者が LLM である場合でも、開発者のエクスペリエンスを完全に制御することです。コードでは、手動で構築されたステート マシンではなく、通常の制御フローと通常の await を使用する必要があります。一方、ランタイムは永続性、再開可能性、冪等性、可視性、および人間参加型の実行を処理します。
この決定は成果をもたらし続けています。製品の拡張性を高めるために、Serval 全体でコードベースのワークフローを使用しています。当社のデータベース製品は良い例です。LLM で生成された Serval ワークフローは外部システムからデータを取り込むことができるため、Serval は顧客の周囲に接続されているツールを完全に把握できます。実際には、これにより Serval は ETL パイプラインとデータストアのようなものになり、オンデマンド オートメーションに使用するのと同じワークフロー基盤上に構築されます。
決定 3: 統合エンジンと実行環境を所有する
3 番目の決定は、統合エンジンとワークフロー コードが実行される環境を所有することでした。
これは配管工事のように聞こえますが、これによって Serval が完全なユニバーサル自動化プラットフォームになることが可能になります。ロングテールを自動化するために、サーバルは事前に構築されたアクションの小さなメニューを公開することはできません。接続されている各システムの完全な API (共通パス、不明瞭なエンドポイント、権限、認証モデル、API が失敗する方法) を理解し、呼び出す必要があります。
そこで、2 つのジョブを実行する統合レイヤーを構築しました。
実行時、これはワークフロー コードと外部システムの間の制御された境界になります。ワークフロー コードは、生の認証情報を直接処理することなく、承認されたインターフェイスを通じて GitHub、Okta、Salesforce、Google Workspace、AWS、Snowflake、Jira、Slack、Terraform、その他のツールを呼び出すことができます。
オーサリング時に、私は

これは、Serval Automation エージェントがこれらのシステムに対して有用なワークフローを作成するために必要なコンテキストです。統合レイヤーは、どのようなオブジェクトが存在するか、どのエンドポイントを呼び出すことができるか、どのようなスコープが必要か、リクエストとレスポンスのタイプがどのようなものであるか、そして良い例がどのようなものかを知っています。
私たちは、ワークフロー コードと統合の間の境界にセキュリティをシステムに組み込む必要がありました。ワークフロー作成者は、承認された機能を呼び出すことができる必要がありますが、ランタイムから生の資格情報にアクセスすることはできません。コードは当社が制御するサンドボックス内で実行されますが、統合レイヤーはそのランタイムの外部で適切な資格情報を適用します。これにより、権限を適用し、アクションを観察し、何が起こったのかを特定し、必要に応じて顧客所有の環境でワークフローを安全に実行するための 1 つの場所が提供されます。
モデルはワークフローを記述することができます。プラットフォームは、安全に実行できるようにする必要があります。つまり、権限があり、耐久性があり、監視可能であり、統合境界でセキュリティが強化されているため、実際のシステム全体でアクションを実行できる必要があります。
Serval は、LLM によって、企業の業務を自動化するためにさらに多くのコードが作成されるようになるだろうと賭けています。そのコードが作成および維持できるほど安価になれば、コード内の反復的なプロセスを自動化できるようになります。

[切り捨てられた]

## Original Extract

How Serval is building a universal automation platform to eliminate manual operational work across companies.

How Serval is building a universal automation platform to eliminate manual operational work across companies.
An employee asks: “Can you give Priya access to the staging database?”
In most companies, that simple request is not simple at all. Someone has to figure out which team Priya is on, whether the request is allowed, who needs to approve it, where the access is defined, how to update it, whether it requires a code change, how to notify the requester, and how to leave behind an audit trail.
The employee does not care whether that work happens through Slack, Okta, GitHub, Terraform, AWS, Jira, or an internal admin tool. From their perspective, they had one thing they needed done.
Serval’s job is to make it feel that way.
Someone asks for something in natural language. Serval figures out what they mean, finds the right tools, runs the automation across the systems involved, brings in a human when needed, records what happened, and lets them get back to work.
Most companies already know which workflows should be automated: access requests, onboarding, hardware requests, infrastructure changes, security reviews, approvals, CRM updates, finance operations, and internal support. They stay manual because automation is usually too expensive to create and too brittle to maintain. If a process crosses five systems, varies by team, depends on a policy that lives in someone’s head, and only happens a few times a month, it usually remains the responsibility of a human.
Serval is built on a simple idea: LLMs change the cost curve of automation. But cheaper automation only matters if it is safe to run. The output has to be structured, permissioned, reviewable, durable, debuggable, and auditable.
The ambition is specific: automate 100% of company operations, not just the easy parts. Most automation tools cover the common, high-volume workflows and leave everything unusual to a person. But the long tail of hard, rare, cross-system requests is where most of the manual work actually lives. Automating the obvious slice is not enough. The goal is to reach the messy long tail: the workflows that were always too rare, too specific, or too expensive to justify automating before.
We had to design for two distinct experiences: employees running automations, and admins creating them. An employee might ask, “Can you give Priya access to the staging database?” An admin might define that as: check Engineering membership, ask the manager for approval, open a Terraform PR that adds the IAM binding, wait for review, and notify the requester when the pipeline applies the change.
In that sense, Serval is not just automating work. It is automating the creation of the automation itself: making it fast to define the workflow, not just fast to run it after it exists.
Doing this well is hard. But we have strong conviction at Serval that if a hard technical choice gets you closer to the best end-user experience, you should take it seriously rather than route around it. A few of our early decisions came from that philosophy.
Decision 1: Build the platform, not the layer
The obvious shortcut would have been to build Serval as a thin layer on top of existing ticketing or IT Service Management (ITSM) systems: read the ticket, use an LLM to classify or summarize it, call a downstream automation, and post the result back. That would have been faster, and it probably would have produced a good early demo.
But it would not have changed the underlying architecture. The ticketing system would still own the request. Access management would still live somewhere else. Approvals would still happen in another tool. Execution state would still be scattered across logs, Slack messages, webhooks, and downstream systems. The LLM would make the existing process feel a little smarter, but the company would still be operating through the same fragmented stack.
That is the problem we wanted to solve directly. Users should not have to know which system owns a request. Admins should not have to wire together tools that were never designed around a shared model of the world. And engineers debugging a failed workflow should not have to reconstruct what happened from five different places.
The iPhone is a great product in large part because Apple owns the hardware and the software together. The experience is not assembled from loosely connected layers; it is cohesive because the pieces are designed to fit. We see the same thing in company operations. The system of record and the system of action need to be unified. If one tool stores the request and another tool executes it, the user experience, admin experience, and audit trail all fragment. So we set out to own the whole surface area, replace what legacy systems do, and then build something vastly better than what is possible with fragmented systems.
Take the Priya request. With a thin layer, you'd have a smarter ticket. But you'd still need someone to find the Terraform file, open the PR, wait for review, and notify the requester. The request would still touch five systems. Serval would just be making the first step feel easier.
Decision 2: Code as the source of truth
Owning the platform meant we had to decide what workflows actually look like inside it.
This was not obvious at the time. A lot of automation products are built around forms, visual workflow builders, decision trees, or graphs. Those can work for simple cases, but they tend to hit a wall when the workflow gets complex: call this API, branch on this policy, wait for this approval, update this record, retry this operation, open this pull request, and make the whole thing understandable later.
When we made this decision in early 2024, our view was that code was going to get much cheaper to write because LLMs were going to be good at writing it. That changed the authoring model. If workflows were often going to be authored by an LLM, they should be represented in a form the model is good at producing and humans are good at reviewing.
Imperative code is that form. It gives you standardized control flow, types, review, versioning, tests, and a way for engineers to understand exactly what will happen. Once code becomes cheap to create, deployment, execution and evaluation become the harder platform problems: making it easy to run that code in a safe, controlled, audited environment. So we built Serval around imperative TypeScript code as the source of truth for deterministic workflows, and made it as simple and secure as possible to run and iterate on that code.
Code is also what lets us target the workflows that do not fit into a predefined catalog. A fixed set of actions or a visual builder can work well for common requests, but it breaks down when the workflow is specific to one company: read this policy file, branch on this approval rule, update this internal system, open this exact pull request, and explain the result in the format this team expects.
Those are exactly the workflows that usually stay manual. A general-purpose programming model gives Serval a way to express them without waiting for someone to build a special-purpose block first.
That does not mean every procedure should be expressed as code. We also support natural language standard operating procedures called Skills, because some parts of a workflow are inherently judgment-based: interpreting what the user is asking for, applying a policy written in human language, deciding whether two requests are equivalent, or explaining why a request was denied.
Both forms matter, but they play different roles. Natural language is useful for reasoning about ambiguous inputs, interpreting user intent, and handling judgment calls. Code is the source of truth for the automation that actually runs: the deterministic, side-effecting, reviewable artifact that can be permissioned, versioned, debugged, and audited.
It also shaped how workflows run, not just how they are written. Most workflow systems expose something graph-shaped: define tasks, connect them with edges, and let the engine execute the graph. That is a good abstraction for data pipelines and batch jobs. It is not the developer experience we wanted for an AI-authored workflow system.
We wanted full control over the developer experience so the thing Serval helps author could look as much as possible like standard imperative TypeScript:
That is the point: Large language models are good at writing this kind of code, and imperative code is a very efficient way to represent logic. You get conditionals, loops, functions, types, libraries, and all the syntax programming languages have evolved to express complex behavior compactly. If necessary, an admin can review the behavior. The side-effecting work is explicit, typed, and inspectable.
Underneath, the runtime still has to behave like a long-running workflow engine. A real company workflow waits on approvals, retries failed integrations, resumes days later, and leaves behind enough state to debug what happened. We built our own durable execution engine for two reasons: to keep authored workflows as close as possible to standard TypeScript, and to have full control over the developer experience, even when the developer is an LLM. The code should use normal control flow and normal await s, not a manually constructed state machine, while the runtime handles persistence, resumability, idempotency, visibility, and human-in-the-loop execution.
This decision has kept paying dividends. We use code-based workflows all over Serval to make the product more extensible. Our Databases product is a good example: LLM-generated Serval workflows can ingest data from external systems so Serval has a complete view of the connected tools around a customer. In practice, that turns Serval into something like an ETL pipeline plus datastore, built on the same workflow substrate we use for on-demand automations.
Decision 3: Own the integration engine and execution environment
The third decision was to own the integration engine and the environment where workflow code runs.
This sounds like plumbing, but it is what makes it possible for Serval to be a complete, universal automation platform. To automate the long tail, Serval cannot expose a small menu of prebuilt actions. It needs to understand and call the full API of each connected system: the common paths, the obscure endpoints, the permissions, the auth model, and the ways the API fails.
So we built the integration layer to do two jobs.
At runtime, it is the controlled boundary between workflow code and external systems. Workflow code can call GitHub, Okta, Salesforce, Google Workspace, AWS, Snowflake, Jira, Slack, Terraform, and other tools through approved interfaces, without ever handling raw credentials directly.
At authoring time, it is the context the Serval automation agent needs to write useful workflows against those systems. The integration layer knows what objects exist, which endpoints can be called, what scopes are required, what the request and response types look like, and what good examples look like.
We wanted security baked into the system at the boundary between workflow code and integrations. Workflow authors should be able to call approved capabilities, but they should not be able to access raw credentials from the runtime. Code runs in a sandbox we control, while the integration layer applies the right credential outside that runtime. That gives us one place to enforce permissions, observe actions, attribute what happened, and, when needed, run workflows safely in customer-owned environments.
The model can write the workflow. The platform has to make it safe to run: permissioned, durable, observable, and capable of taking action across real systems with security enforced at the integration boundary.
Serval is a bet that LLMs will cause much more code to be written to automate company operations. If that code becomes cheap enough to create and maintain, it becomes possible to automate any repetitive process inside a co

[truncated]

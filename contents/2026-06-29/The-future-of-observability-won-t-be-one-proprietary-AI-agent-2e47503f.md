---
source: "https://clickhouse.com/blog/the-future-of-observability-not-one-proprietary-ai-agent-thousands-by-teams"
hn_url: "https://news.ycombinator.com/item?id=48721122"
title: "The future of observability won't be one proprietary AI agent"
article_title: "The future of observability won’t be one proprietary AI agent. It will be thousands built by teams. | ClickHouse"
author: "mikeshi42"
captured_at: "2026-06-29T16:59:53Z"
capture_tool: "hn-digest"
hn_id: 48721122
score: 1
comments: 0
posted_at: "2026-06-29T16:11:48Z"
tags:
  - hacker-news
  - translated
---

# The future of observability won't be one proprietary AI agent

- HN: [48721122](https://news.ycombinator.com/item?id=48721122)
- Source: [clickhouse.com](https://clickhouse.com/blog/the-future-of-observability-not-one-proprietary-ai-agent-thousands-by-teams)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T16:11:48Z

## Translation

タイトル: 可観測性の未来は 1 つの独自の AI エージェントではなくなる
記事のタイトル: 可観測性の未来は、1 つの独自の AI エージェントではありません。それはチームによって何千も構築されるでしょう。 |クリックハウス
説明: 私たちは、可観測性が単一の AI エージェントによって支配されるとは考えていません。未来は、独自のシステム、ワークフロー、運用知識を中心にエージェントを構築するチームに属します。

記事本文:
可観測性の未来は、1 つの独自の AI エージェントではなくなります。それはチームによって何千も構築されるでしょう。 | ClickHouse ロゴを SVG としてコピー 完全なロゴをダウンロード ロゴマークをダウンロード 検索を開く 地域セレクターを開く 英語
製品 製品 ClickHouse Cloud ClickHouse の最適な使い方。
AWS、GCP、Azure で利用できます。
Bring Your Own Cloud フルマネージドの ClickHouse サービス、
自分の AWS および GCP アカウントにデプロイされます。
ClickHouse によって管理される Postgres トランザクション用の統合データ スタック
そして分析。
マネージド ClickStack 高パフォーマンスによるマネージドオブザーバビリティ
クエリと長期保存。
Langfuse Cloud LLM の可観測性と評価
信頼性の高い AI アプリケーションとエージェントを実現します。
オープンソース ClickHouse 高速オープンソース OLAP データベース
リアルタイム分析。
ClickStack ログ用のオープンソース可観測性スタック、
メトリクス、トレース、セッションのリプレイ。
Agentic Data Stack AI を活用したアプリケーションを構築
クリックハウスで。
chDB インプロセス SQL エンジンを搭載
ClickHouse、Pandas 互換 API を使用
ソリューション ユースケース リアルタイム分析
リソース 会社のリソース ユーザーストーリー
ページをコピーしました コピーされました!その他のアクション マークダウンで表示 このページをマークダウンで開く
ChatGPT で開く このページについて質問する
クロードで開く このページについて質問する
v0 で開く このページについて質問する
可観測性の未来は、1 つの独自の AI エージェントではなくなります。それはチームによって何千も構築されるでしょう。
現在、可観測性に対するベンダーの一般的な賭けは収束です。
1 つのプラットフォームに組み込まれた 1 人の SRE エージェントが、インシデントをどのように調査すべきかについての 1 つのベンダーの見解に基づいてトレーニングを受けました。テレメトリを理解し、質問に答え、何が問題だったかを説明し、最終的には誰かがダッシュボードを開く前に本番環境を修正するのに役立ちます。
将来のそのバージョンは便利になるでしょうが、あまりにもナロすぎるでしょう

w。
可観測性は、一般的なサポート キューのようには機能しません。デバッグは、チームが所有するシステム、それらのシステムが失敗する方法、信頼するデータ、従うランブック、使用するツール、および時間の経過とともに蓄積された運用上の傷によって形成されます。データベース チーム、フロントエンド チーム、決済チーム、インフラストラクチャ チームは、同じ方法で本番環境を調査するわけではありません。
この投稿では、可観測性の将来は 1 つの独自の AI エージェントではなく、あなたのようなチームによって構築された何千もの AI エージェントになると私たちが考える理由を探ります。
AI エージェントは可観測性のための新しいインターフェースになりつつあります #
可観測性における AI エージェントへの移行は、もはや特に大胆な予測ではないことに、ほとんどの人が同意してくれることを願っています。
現在、ほとんどの調査は、人間がダッシュボードを開いて、ログを検索し、トレースを検査し、手動でコンテキストを収集することから始まります。その仕事はエージェントに委任されるケースが増えています。モデルはすでに、テレメトリのクエリ、結果の要約、パターンの特定、システム内で何が起こっているかについてのもっともらしい仮説の生成を行うことができます。
モデルが改良され続けるにつれて、可観測性へのインターフェースは人間→ダッシュボード→データではなく、人間→エージェント→データへとますます変わっていくでしょう。どのようなアクションをとるべきかはエンジニアが決定しますが、調査の機械的な作業の多くは、グラフに触れたりクエリを作成したりする前に行われます。
「実際に、Slack のインシデント チャネルで傾向の変化が見られます。以前はエンジニアがログやメトリクスへのリンクを共有していましたが、今ではチームが AI 調査のスニペットを共有し、調査を深く掘り下げています。」
AI が可観測性ワークロードの形状を変える #
可観測性エージェントに関するほとんどの議論は、エージェントが何ができるかに焦点を当てています。それらが通常の一部になると、その下で何が起こっているかにはほとんど注意が払われません。

ワークフロー。
人間の調査員は比較的制約を受けています。いくつかのダッシュボードを開き、いくつかのクエリを実行し、1 つまたは 2 つのトレースを検査し、徐々に検索範囲を狭めます。経験豊富なエンジニアでも、一度に評価できる可能性は非常に限られています。
エージェントにはそのような制限はありません。エンジニアは 2 つの時間枠を比較できますが、エージェントは 20 の時間枠を比較できます。人間はいくつかの可能性のある原因を手動で調査するかもしれませんが、エージェントは数十の仮説を同時に追求し、継続的に証拠を収集し、行き詰まりを排除することができます。
実際的な結果は、調査がより広範囲になり、基盤となるシステムに対する要求が大きくなるということです。エージェントは、答えに収束する前に、より多くの履歴データを調査し、より多くの潜在的な説明を探索できます。これにより、低遅延の応答を必要とするクエリが増加します。
「エージェントは総当たり攻撃を行う可能性があります。クエリを実行する場合、ダッシュボードを 2 回クリックする代わりに、10 個のクエリを実行する必要があります。つまり、クエリの非線形パターンに対応するには、API レイヤーまたはストレージが堅牢である必要があります。」
また、基礎となるデータに関する要件も変更されます。エージェントは、与えられたコンテキストを基に推論することしかできません。履歴データが破棄されている場合、重要なコンテキストが失われる可能性があります。テレメトリが大量にサンプリングされている場合、重要な証拠がデータセットに存在しない可能性があります。経験豊富なエンジニアとは異なり、エージェントは直感や組織的な知識でこれらのギャップを補うことはできません。彼らの結論は、利用可能なデータの完全性と忠実度によって制約されます。
ほとんどの可観測性プラットフォームは、人間の調査員と彼らが生成するワークロードを中心に設計されています。次世代は、人間に代わって行われる調査をサポートする必要性がますます高まるでしょう。
偉大な SRE AG に賭ける業界

エント #
多くの企業が、シンプルな将来ビジョン、つまりユニバーサル SRE エージェントの構築に投資しています。
これは説得力のあるアイデアです。可観測性ベンダーは、プラットフォーム、データ、エージェントを提供します。エンジニアは、ダッシュボードを学習したり、クエリを作成したり、テレメトリを操作したりする代わりに、自然言語で質問します。時間が経つにつれて、エージェントはますます有能になり、最終的には調査プロセスの多くを単独で処理できるようになります。
このモデルには真の価値があります。ほとんどの可観測性ツールは依然として使用が困難ですが、エージェントは複雑なシステムを理解して操作することへの障壁を劇的に下げる可能性を秘めています。
モデルも引き続き改良され、推論、SQL の記述、および共通の調査パターンを再利用可能なワークフローにパッケージ化することがより良くなります。私たちが今日専門知識と考えているものの多くは、エージェントを通じてますますアクセスできるようになるでしょう。
デバッグの難しい部分: コンテキスト #
問題は、デバッグがソフトウェア ベンダーが望むほどきれいに収束しないことです。
モデルは改善を続けていますが、依然として難しいのはコンテキストです。
調査の次のステップは、可観測性プラットフォーム内にあるテレメトリ以上のものに依存します。それは、チームがどのように運営されるか、どのシグナルを信頼するか、以前に何が壊れたか、所有権がどのように分割されるか、運用知識がどこに存在するかによって異なります。そのコンテキストの多くは、運用手順書、チケット、事後検証、Slack スレッド、内部文書、展開システム、経験豊富なエンジニアの責任者などに散在しています。
「Ring Central は 25 年の歴史を持つ会社です。運用チーム内には、すべての Wiki を接続したかどうかに関係なく、どこにも文書化されていない固有の知識がたくさんあります。提供するデータがなければ、ただのことになってしまいます。

暗示する。」
Sushant Hiray 氏、RingCentral、AI リーダー
各チームの働き方は異なります #
さらに、同様のテクノロジースタックを実行している 2 つの企業は、システム、チーム、運用履歴が異なるため、同じインシデントをまったく異なる方法で調査する可能性があります。ベンダーはベスト プラクティスをパッケージ化することはできますが、自社の製品を使用するすべてのエンジニアリング チームの蓄積された経験をパッケージ化することはできません。
実稼働システムのデバッグに必要なコンテキストの多くは、可観測性プラットフォーム自体の外部に存在します。それは、運用手順書、チケット、事後分析、内部文書、Slack スレッド、展開システム、およびシステムを長年運用してきたエンジニアの組織的知識に散在しています。その知識の所在、チームの編成方法、チームが従うプロセスは、組織間だけでなく、業界や組織内の機能によっても異なります。
これらのアプローチはどれも本質的に正しいものではありません。それらは、経験を通じて構築されたさまざまなメンタルモデルを反映しているだけです。チームが信頼するシグナル、従うワークフロー、および依存するコンテキストは、多くの場合、チームが運用するシステムに固有です。
エージェントによる可観測性のためのよりオープンなモデル #
これらの課題を総合すると、多くのベンダーが賭けている未来とはまったく異なる未来を指し示しています。
エージェントが可観測性と調査のための主要なインターフェイスとなり、組織のコンテキストによってますます形作られるようになり、その価値は単一の汎用エージェントから特定のチームやドメイン向けに構築されたエージェントへと移行しています。
「最初から非常にクールな仕組みを構築するのではなく、ヘッドレス プラットフォームの構築に力を入れ、API やデータ ストレージを改善し、可観測性 MCP を構築して、すべてのエンジニアやチームがブロックを解除して独自のプラットフォームを構築できるようにしました。

デバッグのユースケースに合わせて調整された独自のエージェント ワークフロー。」
単一の SRE エージェントに集中するのではなく、オブザーバビリティは、それぞれが特定の組織、チーム、または問題領域に合わせて最適化された、特化されたエージェントのエコシステムに進化する可能性が高いと考えています。インフラストラクチャに重点を置く企業もあれば、データベース、セキュリティ、決済、顧客エクスペリエンス、内部プラットフォームに重点を置く企業もいます。その多くは、組織独自の運用手順書、ドキュメント、ワークフロー、ビジネス ロジック、運用知識を中心に構築されます。
同様に重要なことは、この未来はオープンさにかかっています。
「オープン性」とは、単にオープンソース ソフトウェアを意味するものではありません。つまり、チームや個々のエンジニアが、スタックの各層（モデル、ハーネス、ツール、ワークフロー、インターフェイス）に最適なテクノロジーを自由に選択できるようにするということです。これは、可観測性がどのように機能するべきかというベンダーの見解に合わせてプロセスを適応させるのではなく、組織内にすでに存在するシステムやプロセスを中心にエージェントを構築できることを意味します。
また、これらのエージェントの下にあるレイヤーを自由に選択できることも意味します。チームは、データがどこに存在するか、スキルがどのように開発および管理されるか、本番システムの前にどの MCP ゲートウェイとサーバーが配置されるか、エージェントの動作がどのように観察され、管理され、残りのエンジニアリング環境に統合されるかを制御する必要があります。
「私たちが好む方法は、十分な柔軟性を与えてくれるプラットフォームと提携し、その上に構築する機会があることです。十分な社内ツールで動作する限り、それが最適な点です。」
Sushant Hiray 氏、RingCentral、AI リーダー
最も成功した可観測性プラットフォームは、全員に単一の働き方を強制するものではありません。彼らは、共通の基盤を提供するものになります。

何千もの異なるエージェントを構築できます。
すべてのチーム、あるいは個人が独自のエージェントを構築して適応させた場合、調査は 1 か所で行われなくなります。あるエンジニアは IDE で開始し、別のエンジニアはノートブックで開始し、別のエンジニアは内部チャット インターフェイスで開始し、別のエンジニアはカスタム インシデント ワークフローを通じて開始する場合があります。エージェントは、すべて同じ運用上の問題を理解しようとしている場合でも、異なるハーネスで実行し、異なるモデルを使用し、異なる調査パスをたどることがあります。
これにより、コラボレーションの問題が発生します。調査出力は、一時的なチャット セッションやプライベート エージェント トレース内に閉じ込められたままにすることはできません。チームには、何がクエリされたのか、どのような証拠が見つかったのか、どの仮説が調査されたのか、なぜ結論に達したのかを示す、耐久性があり検査可能な成果物が必要です。これはインシデントを調査する人間にとって重要ですが、毎回ゼロから始めるのではなく、以前の調査から学ぶ必要がある将来のエージェントにとっても重要です。
ビデオを読み込んでいます... その結果、エージェントの可観測性には、人間とエージェントが共同作業できる何らかの形式の永続的な調査面が必要になります。調査を共有、レビュー、再実行、改良し、時間をかけて構築できる場所です。調査が必要かどうか

[切り捨てられた]

## Original Extract

We don’t think observability will be dominated by a single AI agent. The future belongs to teams building agents around their own systems, workflows, and operational knowledge.

The future of observability won’t be one proprietary AI agent. It will be thousands built by teams. | ClickHouse Copy logo as SVG Download full logo Download logomark Open search Open region selector English
Products Products ClickHouse Cloud The best way to use ClickHouse.
Available on AWS, GCP, and Azure.
Bring Your Own Cloud A fully managed ClickHouse service,
deployed in your own AWS and GCP account.
Postgres managed by ClickHouse Unified data stack for transactions
and analytics.
Managed ClickStack Managed observability with high-performance
queries and long-term retention.
Langfuse Cloud LLM observability and evaluations
for reliable AI applications and agents.
Open source ClickHouse Fast open-source OLAP database for
real-time analytics.
ClickStack Open-source observability stack for logs,
metrics, traces, and session replays.
Agentic Data Stack Build AI-powered applications
with ClickHouse.
chDB In-process SQL Engine powered by
ClickHouse, with a Pandas-compatible API
Solutions Use cases Real-time analytics
Resources Company resources User stories
Copy page Copied! More actions View as Markdown Open this page in Markdown
Open in ChatGPT Ask questions about this page
Open in Claude Ask questions about this page
Open in v0 Ask questions about this page
The future of observability won’t be one proprietary AI agent. It will be thousands built by teams.
The common vendor bet in observability right now is convergence.
One SRE agent, built into one platform, trained around one vendor’s view of how incidents should be investigated. It understands your telemetry, answers your questions, explains what went wrong, and eventually helps fix production before anyone opens a dashboard.
That version of the future will be useful, but it will also be too narrow.
Observability does not work like a generic support queue. Debugging is shaped by the systems a team owns, the way those systems fail, the data they trust, the runbooks they follow, the tools they use, and the operational scars they have accumulated over time. A database team, frontend team, payments team, and infrastructure team do not investigate production the same way.
In this post, we explore why we think the future of observability will not be one proprietary AI agent, but thousands built by teams like yours.
AI agents are becoming the new interface for observability #
I hope most people would agree that the shift to AI agents in observability is no longer a particularly bold prediction.
Today, most investigations begin with a human opening dashboards, searching logs, inspecting traces, and manually gathering context. Increasingly, that work is being delegated to agents. Models are already capable of querying telemetry, summarizing findings, identifying patterns, and generating plausible hypotheses about what is happening within a system.
As models continue to improve, the interface to observability will increasingly become human → agent → data instead of human → dashboard → data. Engineers will still decide what action to take, but much of the mechanical work of an investigation will happen before they ever touch a chart or write a query.
"We actually see a trend shift in our Slack incident channels. Engineers previously used to share links to logs or metrics. Now teams are sharing snippets of AI investigations and diving deep into it."
AI changes the shape of observability workloads #
Most discussions about observability agents focus on what the agents can do. Far less attention is given to what happens underneath when they become a normal part of the workflow.
A human investigator is relatively constrained. They open a handful of dashboards, run some queries, inspect a trace or two, and gradually narrow the search space. Even experienced engineers can only evaluate so many possibilities at once.
An agent has no such limitation. While an engineer may compare two time windows, an agent can compare twenty. While a human might manually investigate a few likely causes, an agent can pursue dozens of hypotheses simultaneously, continuously gathering evidence and eliminating dead ends as it goes.
The practical consequence is that investigations become broader and place greater demands on the underlying systems. Agents can examine more historical data and explore far more potential explanations before converging on an answer. This all results in more queries requiring low-latency responses.
"Agents could brute force it — make 10 queries instead of what, if I query, I would make two dashboard clicks. Which means our API layer or storage have to be robust to take that non-linear pattern of queries."
It also changes the requirements around the underlying data. Agents can only reason over the context they are given. If historical data has been discarded, important context may be missing. If telemetry has been heavily sampled, critical evidence may simply not exist in the dataset. Unlike experienced engineers, agents cannot compensate for these gaps with intuition or institutional knowledge. Their conclusions are constrained by the completeness and fidelity of the data available to them.
Most observability platforms were designed around human investigators and the workloads they generate. The next generation will increasingly need to support investigations carried out on behalf of humans.
An industry betting on the great SRE agent #
A lot of companies are investing in building for a simple vision of the future: the universal SRE agent.
This is a compelling idea. An observability vendor provides the platform, the data, and the agent. Engineers ask questions in natural language instead of learning dashboards, writing queries, or navigating telemetry. Over time, the agent becomes increasingly capable, eventually handling much of the investigation process on its own.
There is real value in this model. Most observability tools remain difficult to use, and agents have the potential to dramatically lower the barrier to understanding and operating complex systems.
Models will also continue to improve and will become better at reasoning, better at writing SQL, and better at packaging common investigation patterns into reusable workflows. Much of what we consider expert knowledge today will become increasingly accessible through agents.
The hard part of debugging: context #
The problem is that debugging does not converge quite as neatly as software vendors would like.
While models continue to improve, what remains difficult is context.
The next step in an investigation depends on far more than the telemetry sitting inside an observability platform. It depends on how a team operates, which signals it trusts, what has broken before, how ownership is divided, and where operational knowledge lives. Much of that context is scattered across runbooks, tickets, postmortems, Slack threads, internal documentation, deployment systems, and the heads of experienced engineers.
"Ring Central is a 25-year-old company. There is a lot of tribal knowledge within the operations team that is not documented anywhere, no matter whether we've connected all the wikis. If you don't have any data to give it, it's going to just hallucinate."
Sushant Hiray, AI Leader, RingCentral
Every team works differently #
Furthermore, two companies running similar technology stacks can investigate the same incident in completely different ways because their systems, teams, and operational history are different. A vendor can package best practices, but it cannot package the accumulated experience of every engineering team that uses its product.
Much of the context required to debug production systems lives outside the observability platform itself. It is scattered across runbooks, tickets, postmortems, internal documentation, Slack threads, deployment systems, and the institutional knowledge of engineers who have operated the system for years. The location of that knowledge, the way teams are structured, and the processes they follow vary not only across organizations, but across industries and functions within those organizations.
None of these approaches is inherently correct. They simply reflect different mental models built through experience. The signals teams trust, the workflows they follow, and the context they rely on are often unique to the systems they operate.
A more open model for agentic observability #
Taken together, these challenges point toward a very different future than the one many vendors are betting on.
As agents become the primary interface for observability and investigations increasingly shaped by organizational context, the value shifts away from a single universal agent toward agents built for specific teams and domains.
"Instead of building a super-cooled arrangement from the get-go, we doubled down on building a headless platform - improve our APIs, our data storage - and build an observability MCP so that we would unblock or enable every engineer or every team to build their own agentic workflows, which are more tailored towards their debugging use cases."
Instead of converging on a single SRE agent, we believe observability is more likely to evolve into an ecosystem of specialized agents, each optimized for a particular organization, team, or problem space. Some will focus on infrastructure, others on databases, security, payments, customer experience, or internal platforms. Many will be built around an organization’s own runbooks, documentation, workflows, business logic, and operational knowledge.
Just as importantly, this future depends on openness.
By "openness," we do not simply mean open-source software. We mean giving teams and individual engineers the freedom to choose the best technology for each layer of the stack: models, harnesses, tools, workflows, and interfaces. It means being able to build agents around the systems and processes that already exist within an organization rather than adapting those processes to fit a vendor’s view of how observability should work.
It also means having the freedom to choose the layers beneath those agents. Teams should control where their data lives, how skills are developed and managed, which MCP gateways and servers sit in front of production systems, and how the agent’s behavior is observed, governed, and integrated into the rest of the engineering environment.
"The way we prefer is to partner with platforms that give us enough flexibility that there's an opportunity for us to build on top of that. As long as it works with enough of our internal tooling, that's the sweet spot."
Sushant Hiray, AI Leader, RingCentral
The most successful observability platforms will not be the ones that force everyone into a single way of working. They will be the ones that provide a shared foundation upon which thousands of different agents can be built.
If every team or even individual builds and adapts their own agents, investigations will not happen in one place. One engineer may start in an IDE, another in a notebook, another in an internal chat interface, and another through a custom incident workflow. Agents may run in different harnesses, use different models, and follow different investigation paths, even when they are all trying to understand the same production issue.
This creates a collaboration problem. Investigation output cannot remain trapped inside transient chat sessions or private agent traces. Teams need durable, inspectable artifacts that show what was queried, what evidence was found, which hypotheses were explored, and why a conclusion was reached. This matters for humans reviewing an incident, but it also matters for future agents that need to learn from previous investigations instead of starting from scratch every time.
Loading video... As a result, agentic observability will require some form of persistent investigation surface where humans and agents can collaborate. A place where investigations can be shared, reviewed, rerun, refined, and built upon over time. Whether an investigation beg

[truncated]

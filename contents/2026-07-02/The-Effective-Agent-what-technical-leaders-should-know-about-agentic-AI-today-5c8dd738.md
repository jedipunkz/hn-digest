---
source: "https://gkanellopoulos.com/ai-in-the-enterprise/the-effective-agent/"
hn_url: "https://news.ycombinator.com/item?id=48766539"
title: "The Effective Agent: what technical leaders should know about agentic AI today"
article_title: "The Effective Agent | George Kanellopoulos"
author: "gkanellopoulos"
captured_at: "2026-07-02T20:55:24Z"
capture_tool: "hn-digest"
hn_id: 48766539
score: 2
comments: 1
posted_at: "2026-07-02T19:56:48Z"
tags:
  - hacker-news
  - translated
---

# The Effective Agent: what technical leaders should know about agentic AI today

- HN: [48766539](https://news.ycombinator.com/item?id=48766539)
- Source: [gkanellopoulos.com](https://gkanellopoulos.com/ai-in-the-enterprise/the-effective-agent/)
- Score: 2
- Comments: 1
- Posted: 2026-07-02T19:56:48Z

## Translation

タイトル: 効果的なエージェント: 今日のエージェント AI について技術リーダーが知っておくべきこと
記事のタイトル: 効果的なエージェント |ジョージ・カネロプロス
説明: Agentic AI は、2026 年に最も宣伝されているものの、運用化が最も進んでいないテクノロジーです。組織の 79% が AI エージェントを導入していると報告していますが、実際にソリューションを導入しているのは 11% のみです。このホワイト ペーパーでは、プロダクション ハーネスの構造を通じて、アーキテクチャ パターンが現実との接触を経ても存続する理由を検証します。
[切り捨てられた]

記事本文:
エージェントティック AI は、2026 年に最も誇大宣伝されていると同時に、最も運用化されていないテクノロジーです。組織の 79% が AI エージェントを導入していると報告していますが、運用中のソリューションを持っているのはわずか 11% であり、ガートナーは、エージェントティック AI プロジェクトの 40% 以上が 2027 年末までにキャンセルされると予測しています。この論文では、このギャップはモデルの能力とはほとんど関係がないと主張しています。モデルは商品となっております。成功を決定するのは、ハーネス、ガバナンス アーキテクチャ、経済性、組織自体など、成功を取り巻くすべての要素です。
そのことを主張するために、この論文では共通の語彙を確立し、本番環境のハーネスの 5 つの層を分析し、本番環境との接触を維持するアーキテクチャ パターンをマッピングし、エージェントが従来のソフトウェアとは異なる方法で失敗する理由を説明し、ガバナンス、コスト、組織の準備状況をエンジニアリング上の懸念事項として扱います。最後に、技術的リーダーシップに関する 10 の推奨事項を紹介します。簡単に言えば、ハーネスに投資し、ガバナンスをアーキテクチャとして扱い、テクノロジーだけでなく組織を再設計することです。
私たちは AI の分野において重要な岐路に立っています。 LLM とトランスフォーマー アーキテクチャは、数年前には想像できなかった AI 活用の可能性をもたらしました。 Gartner によると、エージェントティック AI が 2025 年の戦略的テクノロジーのナンバーワン [2] に選ばれた後、2025 年と 2026 年にエージェントティック AI の期待は膨らんだ期待のピークに達しました [1]。そうかもしれないが、企業は導入が難しいと感じており、その結果、異常な注目が集まり実行が遅れる環境が生まれているようだ。
データを見ると、状況がさらに明確になります。導入に関しては、79% の組織が AI エージェントを導入していると報告しており [3]、エンタープライズ アプリの 40% はタスク固有のエージェントを組み込むことが期待されています。

2026年末までに男性に。同時に、実行面では、Agentic AI ソリューションのわずか 11% のみがすでに運用されている [5] 一方で、かなりの数の経営幹部が AI の導入により「会社がバラバラになっている」と述べています [6]。さらに直観に反することに、データは、エージェント AI プロジェクトの 40% が 2027 年末までにキャンセルされる予定であることを示しています [7]。成長と失敗は同時に起こります。
今こそ、一歩下がって誇大宣伝から離れ、Agentic AI テクノロジーを採用することがなぜそれほど難しいのか、そしてモデルの機能とそれを取り巻くハーネスの間のギャップをどのように埋めることができるのかを議論するときです。この文書では、技術的なリーダーシップを教育しながら、まさにそれを試みます。
いつもそうなのですが、誇大広告はテクノロジーが急速に進歩する環境を生み出しましたが、その一方でテクノロジーに関するセマンティクスは曖昧に定義されたままであり、最終的にはコミュニケーションやアイデアの交換に必要な共通の語彙が存在しません。そのため、重要な用語と定義から始めます。
Agentic AI と AI Agent という用語が同じ意味で使用されているのを聞いたことがあるのは、あなただけではありません。ただし、この 2 つはまったく異なる意味を持っており、区別することが重要です。より小さな単位から始めて、AI エージェントが個別のソフトウェア エンティティであることを定義する必要があります。 LLM 主導、タスク指向、ツール統合された単一のソフトウェア。
一方、エージェント AI は、動的なタスク分解を実行し、セッション間で永続的な知識を獲得し、最終的に自律性を獲得するために複数の AI エージェントが連携するときに表現される動作です。それはパラダイムであり、実体ではありません [8]。
ワークフローは、AI 時代以前からよく知られ、明確に定義されていた用語であり、この文脈でも何ら変わりません。コード パスは LLM とツールを利用して、

特定の一連のイベントを決定論的かつ予測可能な方法で分析します。エージェントには、独自のプロセスとツールの使用を指示する自由があります。何が起こったのかを評価し、次に何をすべきかを決定できます。環境からのフィードバックを分析しながら、ループ内で独立して動作します [9]。
最近大きな注目を集めている用語は、「ハーネス」という言葉です。モデルの周囲のすべてのものを説明するために使用されます。ツール、検証ループ、メモリ、ガードレール、可観測性。 Viv Trivedy は、これを理解しやすい式で表現しました [10]。
ハーネスは、モデルに安全かつ一貫性のある有用な方法で思考させるメカニズムです。
プロンプト、コンテキスト、ハーネスという 3 つのタイプのエンジニアリングに注意する必要があります。プロンプト エンジニアリングとは、単一モデルの相互作用に対する効果的な指示を作成することを指します。コンテキスト エンジニアリングは、モデルが理性に依存する情報環境全体を設計するプロセスに関連します。これには、メモリ、ドキュメント、ツール定義、会話履歴、構造化された出力仕様が含まれる場合があります。ハーネス エンジニアリングには、エージェントのライフサイクル全体を管理する制御システムの設計と保守が含まれます。それよりも、エージェントがどのように行動するかが重要です。 3 つのタイプのエンジニアリングがすべて相互に共存していることに注意することが重要です。プロンプト エンジニアリングはコンテキスト エンジニアリングの内部に存在し、コンテキスト エンジニアリングはハーネスの内部に存在します [11] [12]。
マルチエージェント システムとは、それぞれが特化した役割を持つ複数の自律型 AI エージェントで構成され、単一エージェントの能力を超えたタスクを達成するために調整します。マルチエージェントという用語は、アップグレード技術ではなく、アーキテクチャを指します。エージェントが多ければ多いほど良いというわけではありません。 3 つのアーキテクチャ パターン (ハブアンドスポーク、

階層メッシュとフラット メッシュ)、セクション 4 で詳しく説明します。
オーケストレーションは、複数のエージェント、ツール、ワークフローを調整して、一貫したシステムにするプロセスです。それはコントロールプレーンです。エージェントがワーカーとみなされる場合、オーケストレーションは管理層になります。実際には、オーケストレーションは実行時に次の 4 つの質問に答えようとします: 目標をタスクに分割する方法 (分解)。各タスク (ルーティング) をどのエージェントまたはツールが処理するか。どのタスクが並列実行でき、どのタスクが線形である必要があるか (シーケンス)。結果を一貫した出力に合成する方法 (合成)。ランドスケープを支配するオーケストレーション パターン (スーパーバイザー、スウォーム、ファンアウト/ファンイン) は、セクション 4 で検討する調整トポロジにマッピングされます。一般に、オーケストレーション層は生産価値が集中する場所であり、ソリューションが信頼性があり、コスト効率が高く、管理可能であるかどうかが決定される場所です。これは、コンテキストのサイジング、予算、レイテンシーの目標が決定されるレイヤーでもあります。
Agentic AI ガードレールのコンテキストでは、AI エージェントが決定、アクセス、実行できる内容を制限する技術的および組織的な制御です。それらは単なるセーフティネットではありません。これらは、エージェントが動作できる空間を定義する境界です。ハーネスがラッパーの場合、ガードレールはラッパー内のフェンスです。従来のソフトウェア セキュリティとの違いは、何をすべきかを明示的に指示しなくても、何かを実行できる機能を備えたソフトウェアを初めて導入したという事実にあります。従来のアクセス制御は、ソフトウェアがプログラムされたとおりに動作することを前提としています。エージェントのガードレールは、そうでない可能性があることを想定しています。
定義されたガードレール層は 5 つあります。
データとコンテキストのガードレール: グラフ、ビジネス定義などの基盤

ビジョンや政策など。
設計時のガバナンス : 誰がどのエージェントを、どのデータを使用して、どのような制約の下で構築できるかを定義します。
ランタイム ガードレール : プロンプト インジェクション、PII 編集、ポリシー違反アラート、および超過した場合にレビューをトリガーできるしきい値によるリスク スコアリングに対処するために、実行中に強制されます。
ID とアクセス ガードレール : 範囲が限定され、有効期間が短い権限を持つ各エージェントの個別の ID。
人間による監視 : エージェントの権限を超える決定に対する承認ワークフローとエスカレーション パス。
一般に、ガードレールはボルトオンの準拠チェックボックスではありません。これらは設計されたアーキテクチャー層であり、後から追加されるものではありません。
議論する必要がある最後の 3 つの重要な用語はスペクトルを形成するため、それらを相互に関連付けて議論する必要があります。 Human-in-the-Loop (HITL) とは、エージェントが定義されたチェックポイントで一時停止し、実行前に人間の承認を待つメカニズムを指します。この手法は、リスクが高く、取り返しのつかない決定を行う場合に使用されます。 Human-on-the-Loop (HOTL) は、エージェントが自律的に動作し、人間が監視して事後に介入できるアプローチを指します。これは主に、スピードが重要であり、可逆的な決定が行われる中リスクのユースケースに使用されます。最後に、Governed Autonomy では、システム自体がアクションごとの承認ではなく、信頼メカニズムを使用して設計および設計されています。人間が事前にポリシーを設計し、エージェントはその境界内で自律的に動作します [13]。
前のセクションでは、エージェントの語彙について説明しました。このセクションでは、実稼働環境でエージェントが効率的に動作する仕組みについて詳しく説明します。
そのために、核となる方程式を詳しく調べ、ハーネス項に注目します。
プロダクションハーネスの 5 つの層
ハーネス

プロダクションには、理想的には 5 つのレイヤーが含まれ、すべて重要で相互依存している必要があります。
図 1: 実稼働ハーネスの 5 つの層。モデルが環境に直接触れることはありません。あらゆる機能とあらゆる安全装置がハーネスに組み込まれています。
ツール オーケストレーションは、エージェントが現実世界と出会う場所です。これは、エージェントを外部ツール、API、データ ソースに接続し、実行時にそれらの接続がどのように機能するかを管理するコントロール プレーンです。エージェントがデータベースの読み取り、API の呼び出し、またはコードの実行を決定すると、オーケストレーション層が選択、実行、エラー回復、結果のルーティングを処理します。
運用環境では、モデルが JSON スキーマ ツール定義を受け取り、構造化された呼び出しを発行するという基本的なループが行われます。ハーネスがそれらを実行し、結果がコンテキスト ウィンドウに戻ります。モデルがツールに直接触れないことを理解することが重要です。このループを標準化するプロトコル、そして 2026 年の事実上のプロトコルは MCP (Model Context Protocol) であり、数百万回ダウンロードされ、すべての主要な AI プロバイダーによって採用されています [14]。同じループ内での並列ツール呼び出しもパフォーマンスの最適化とみなされます [15]。ハーネスは、一度に 1 つのツール呼び出しを行うのではなく、複数の呼び出しを同時にファンアウトします。
現在のツール オーケストレーション アーキテクチャと標準には、興味深い影響と憂慮すべき影響の両方があります。ツール検索により、すべてのツール定義を事前にロードすることを回避できるため、トークンが大幅に削減されます。これは純粋にオーケストレーションの最適化であり、モデルの改善ではありません。さらに、MCP 機能である動的ツール登録により、エージェントを再起動せずに実行時にツールを変更できます。新しい定義は、実行中のセッションに即座に反映され、明らかなメリットが得られます。
一方、現在の建築家は、

ラルおよびプロトコル環境は脆弱性を生み出します。 Endor Labs による 2,600 を超える MCP 実装の分析では、82% でパス トラバーサルが発生しやすいファイル システム操作が使用されていました [16]。一方、実稼働 AI 展開ではプロンプト インジェクションが依然として主要な攻撃ベクトルであり続けています [17]。 SOTA モデルは、長い会話や動的な意思決定におけるツールの選択に苦労します。セキュリティはモデルレベルの命令に依存できません。実行層での入力検証、サンドボックス分離、ファイルシステム パスのスコープ設定とともに、ツールごとの最小権限の原則が必要です。
すべてのツールは、権限の境界、コスト センター、および攻撃対象領域となります。オーケストレーション層は、3 つすべてを同時に処理する必要があります。
検証ループには、実行中にエージェントの出力を検証するすべての自動品質チェックが含まれます。これらは、エージェントのような確率的システムが決定論的な作業を実行していることを保証するために存在します。実際には、ハーネスは、エージェントが続行する前に、すべてのツール呼び出しをインターセプトし、すべての結果をチェックする責任があります。モデルの推論のみに基づいて完全なものとして扱われるものはありません。
モデルが正しいかどうかに依存するのではなく、ハーネスはツール呼び出しをインターセプトし、そのパラメーターを検証し、砂の中で実行します。

[切り捨てられた]

## Original Extract

Agentic AI is the most hyped and least operationalized technology of 2026: 79% of organizations report adopting AI agents, yet only 11% have solutions in production. This white paper examines why, through the anatomy of a production harness, the architecture patterns that survive contact with realit
[truncated]

Agentic AI is at once the most hyped and the least operationalized technology of 2026. 79% of organizations report adopting AI agents yet only 11% have solutions in production, and Gartner predicts that over 40% of agentic AI projects will be canceled by the end of 2027. This paper argues that the gap has little to do with model capability. The model has become a commodity. What determines success is everything that surrounds it: the harness, the governance architecture, the economics and the organization itself.
To make that case, the paper establishes a shared vocabulary, dissects the five layers of a production harness, maps the architecture patterns that survive contact with production, explains why agents fail differently than traditional software, and treats governance, cost and organizational readiness as the engineering concerns they are. It closes with ten recommendations for technical leadership. The short version: invest in the harness, treat governance as architecture and redesign the organization, not just the technology.
We are at a critical juncture in the AI landscape. LLMs and the transformer architecture brought forth possibilities for utilizing AI that were impossible to imagine a few years back. In 2025 and 2026 the Agentic AI promise has, according to Gartner, reached the Peak of Inflated Expectations [1], after agentic AI was named the number one strategic technology trend for 2025 [2]. While that may be, it seems that companies are finding it difficult to adopt, thus creating an environment of extraordinary attention and lagging execution.
Looking at the data, the picture becomes even clearer. When it comes to adoption, 79% of organizations report adopting AI agents [3] and 40% of enterprise apps are expected to embed task-specific agents by the end of 2026 [4]. At the same time, on the execution side, only 11% of Agentic AI solutions are already in production [5] while a significant number of C-suite executives state that AI adoption is “tearing their company apart” [6]. To make things even more counterintuitive, the data shows that 40% of agentic AI projects are slated to be cancelled by the end of 2027 [7]. Growth and failure happening simultaneously.
The moment is now to take a step back, unmount from the hype and discuss why it is so challenging to adopt Agentic AI technologies and how we can close the gap between model capability and the harnesses that surround it. This paper will attempt to do exactly that while trying to educate technical leadership.
The hype, as it always does, has created an environment in which the technology moves fast while the semantics around it remain vaguely defined, and at the end there is no shared vocabulary to communicate and exchange ideas. That is why we begin with the key terms and definitions.
If you have heard the terms Agentic AI and AI Agent being used interchangeably you are not alone. However, the two mean quite different things and the distinction matters. Starting with the smaller unit we need to define that an AI Agent is the discrete software entity. A single piece of software that is LLM-driven, task-oriented and tool-integrated.
Agentic AI on the other hand, is the behavior expressed when multiple AI Agents collaborate in order to perform dynamic task decomposition, achieve persisting knowledge across sessions and finally gain autonomy. It is the paradigm, not the entity [8].
Workflow is a term that was well-known and well-defined before the AI era, and it means nothing different in this context. Code paths utilize LLMs and tools to execute a specific sequence of events in a deterministic and predictable way. An Agent has the freedom to direct its own process and tool usage. It can assess what just happened and decide what to do next. It operates independently in a loop while analyzing feedback from the environment [9].
A term that lately has gained significant traction is the word Harness . It is used to describe everything that surrounds the model. The tools, verification loops, memory, guardrails and observability. Viv Trivedy has expressed this in a formula that makes it easy to understand [10]:
The harness is the mechanism that makes the model think in a safe, consistent, and useful way.
There are three types of engineering you need to be aware of: Prompt, Context and Harness. Prompt engineering refers to crafting effective instructions for a single model interaction. Context engineering relates to the process of designing the entire information environment the model relies on to reason. That might include memory, documents, tool definitions, conversation history and structured output specs. Harness engineering involves designing and maintaining the control system that governs an agent across its entire lifecycle. It is more about how the agent acts. It is important to note that all three types of engineering co-exist within each other. Prompt engineering lives inside context engineering and context engineering lives inside the harness [11] [12].
A Multi-Agent System is one that comprises multiple autonomous AI agents that each have specialized roles while coordinating in order to accomplish tasks beyond any single agent’s capability. The term Multi-Agent refers to the architecture and not an upgrade technique. More agents do not mean better. Three architectural patterns have emerged (hub-and-spoke, hierarchical and flat mesh) and we will examine them in detail in section 4.
Orchestration is the process of coordinating multiple agents, tools and workflows into a cohesive system. It’s a control plane. If the agents are viewed as workers, orchestration is the management layer. In practice, orchestration tries to answer four questions at runtime: How do we break the goal into tasks (Decomposition). Which agent or tool will handle each task (Routing). Which tasks can run in parallel and which have to be linear (Sequencing). How the results are to be synthesized into a coherent output (Synthesis). The orchestration patterns that dominate the landscape (supervisor, swarm, fan-out/fan-in) map onto the coordination topologies we examine in section 4. In general, the orchestration layer is where the production value concentrates and where it is decided whether the solution will be reliable, cost-efficient and governable. It is also the layer where context sizing, budgets and latency targets are decided.
In the context of Agentic AI guardrails are the technical and organizational controls that constrain what an AI agent can decide, access and execute. They are not simply safety nets. They are boundaries that define the space within which an agent is allowed to operate. If the harness is the wrapper, guardrails are the fences inside the wrapper. The difference from traditional software security lies in the fact that for the first time we have deployed software that has the ability to go do something without us explicitly telling it what to do. Traditional access controls assume software does what it’s programmed to do. Agent guardrails assume it might not.
There are five defined guardrail layers:
Data and context guardrails : the foundation, such as graphs, business definitions, and policies.
Design-time governance : defines who can build what agents, using which data and under what constraints.
Runtime guardrails : enforced during execution to address prompt injection, PII redaction, policy violation alerts, and risk scoring with thresholds that can trigger review when exceeded.
Identity and access guardrails : a distinct identity for each agent with scoped and short-lived permissions.
Human-in-the-loop oversight : approval workflows and escalation paths for decisions that go above the agent’s authority.
In general, guardrails are not a bolt-on compliance checkbox. They are an architectural layer designed in and not added after.
The final three key terms we need to discuss form a spectrum and as such we need to discuss them in relation to each other. Human-in-the-Loop (HITL) refers to the mechanism in which the agent pauses at defined checkpoints and waits for a human to approve before executing. This technique is used for high-risk and irreversible decisions. Human-on-the-Loop (HOTL) describes the approach where the agent acts autonomously while a human monitors and can intervene after the fact. It is used mainly for medium-risk use cases where speed matters and reversible decisions take place. Lastly, in Governed Autonomy the system itself is designed and engineered with trust mechanisms and not per-action approvals. Humans design the policies upfront and agents operate autonomously within those boundaries [13].
In the previous section we discussed the vocabulary of agents. In this section we will take a deep dive into what makes an agent work efficiently in a production environment.
To do that we will examine closely the core equation and turn our attention onto the Harness term
The five layers of a production harness
A harness in production should ideally contain five layers, all important and all interdependent:
Figure 1: The five layers of a production harness. The model never touches the environment directly; every capability and every safeguard lives in the harness.
Tool orchestration is where the agent meets the real world. It is the control plane that connects agents to external tools, APIs and data sources while managing how those connections work at runtime. When an agent decides to read a database, call an API or run code, the orchestration layer handles selection, execution, error recovery, and result routing.
In a production environment the basic loop is that the model receives JSON schema tool definitions and emits structured invocations. The harness executes them and the results flow back into the context window. It is important to understand that the model never touches the tool directly. The protocol that standardizes this loop, and the de facto one in 2026, is MCP (Model Context Protocol), with millions of downloads and adoption by every major AI provider [14]. Parallel tool invocation within the same loop is also considered to be a performance optimization [15]. Instead of one tool call at a time, the harness fans out multiple calls concurrently.
There are both interesting and worrying implications on current tool orchestration architectures and standards. Tool search allows avoiding loading all tool definitions upfront thus leading to significant token reduction. This is purely an orchestration optimization, not a model improvement. Furthermore, dynamic tool registration which is an MCP capability, enables tools to change at runtime without agents restarting. New definitions propagate instantly to running sessions with obvious benefits.
On the other hand, the current architectural and protocol landscape creates vulnerabilities. In one Endor Labs analysis of more than 2,600 MCP implementations, 82% used file system operations prone to path traversal [16], while prompt injection remains the dominant attack vector in production AI deployments [17]. SOTA models struggle with tool selection across long conversations and dynamic decision-making. Security cannot rely on model-level instructions. Least privilege per tool principle is required along with input validation at the execution layer, sandbox isolation and filesystem path scoping.
Every tool is a permission boundary, a cost center, and an attack surface. The orchestration layer must handle all three simultaneously.
Verification loops include all the automated quality checks that validate agent outputs during execution. They are there to ensure that probabilistic systems like agents are doing deterministic work. In practice the harness is responsible for intercepting every tool call and checking every result before the agent proceeds. Nothing is treated as complete based on the model’s reasoning alone.
Rather than relying on the model to be correct, the harness intercepts a tool call, validates its parameters, executes in a sand

[truncated]

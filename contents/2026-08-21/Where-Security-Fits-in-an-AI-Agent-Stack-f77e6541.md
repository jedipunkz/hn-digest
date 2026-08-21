---
source: "https://developer.nvidia.com/blog/where-security-fits-in-an-ai-agent-stack/"
hn_url: "https://news.ycombinator.com/item?id=49390933"
title: "Where Security Fits in an AI Agent Stack"
article_title: "Where Security Fits in an AI Agent Stack | NVIDIA Technical Blog"
image: "https://developer-blogs.nvidia.com/wp-content/uploads/2026/08/agentic-ai-visual-security-tech-5597121_r10-1920x1080-1.webp"
author: "pretext"
captured_at: "2026-08-21T17:19:55Z"
capture_tool: "hn-digest"
hn_id: 49390933
score: 1
comments: 0
posted_at: "2026-08-21T17:01:26Z"
tags:
  - hacker-news
  - translated
---

# Where Security Fits in an AI Agent Stack

- HN: [49390933](https://news.ycombinator.com/item?id=49390933)
- Source: [developer.nvidia.com](https://developer.nvidia.com/blog/where-security-fits-in-an-ai-agent-stack/)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T17:01:26Z

## Translation

タイトル: AI エージェント スタックにおけるセキュリティの位置付け
記事のタイトル: AI エージェント スタックにセキュリティが適合する場所 | NVIDIA テクニカル ブログ
説明: AI エージェントの能力が向上し、長期にわたって運用できるようになるにつれて、AI エージェントが機能するアプリケーションにセキュリティと信頼を構築することがますます重要になります。

記事本文:
AI エージェント スタック内でセキュリティが適合する場所 | NVIDIA テクニカル ブログ
開発者
ホーム
購読する
関連リソース
エージェントAI / ジェネレーティブAI
AI エージェント スタック内でセキュリティが適合する場所
NVIDIA の AI 安全性およびセキュリティ チームは、AI エージェントがますます複雑化する作業に取り組む中で、セキュリティ制御がどこに属するかを検討します。
いいね
嫌い
フロンティア AI エージェントに関係する最近のインシデントは、創造的な問題解決能力を持つエージェントが意図した制限を回避する能力を実証しているため、エージェント スタック内で明確に定義されたセキュリティ境界の重要性を浮き彫りにしています。
エージェント スタックは、個別のレイヤーモデル、ハーネス、メタハーネス、NVIDIA OpenShell などの安全なランタイム、および変更可能なハーネス ロジック内ではなく、ランタイム レイヤーとインフラストラクチャ レイヤーで最も効果的にセキュリティ制御が適用される推論インフラストラクチャで構成されます。
効果的なエージェントのセキュリティは、最小特権、分離、ジャストインタイム アクセス、エージェント境界の下での権限のあるポリシーの適用などの原則に依存しており、エージェントが自分自身に権限を付与したり制御をバイパスしたりできないようにし、影響のあるすべてのアクションが一貫して評価および監査可能であることを保証します。
AI によって生成されたコンテンツでは、情報が不完全に要約されている可能性があります。重要な情報を確認してください。さらに詳しく
AI エージェントの能力が向上し、長期にわたって運用できるようになるにつれて、AI エージェントが機能させるアプリケーションにセキュリティと信頼を構築することがますます重要になります。 NVIDIA OpenShell、エージェント開発者、オープンソース プロジェクト、エコシステム全体のパートナーとの連携を活用して、NVIDIA の AI 安全性およびセキュリティ チームは、各レイヤーの役割やセキュリティがどこに存在すべきかなど、新たなエージェント スタックに関する視点を提供します。
最近のレポートは、セキュリティ管理の配置が重要である理由を強調しています。数週間以内に

この夏、OpenAI、Anthropic、英国 AI セキュリティ研究所はそれぞれ、フロンティア エージェントが意図した境界を超えて活動していると報告しました。報告された行為には、研究室環境からオープンインターネットへの予期せぬ経路を悪用したり、他社のシステムに不正アクセスしたり、人々やインフラストラクチャに関与する未承認の行動をとったりすることが含まれます。これらのケースには、削減されたモデルの保護機能を使用して実行されている長期的なエージェントが含まれていました。しかし、彼らは同じ設計上の課題を指摘しています。つまり、エージェントが創造的に問題を解決し、複雑な目標を追求できるようにする機能は、元の指示が予期していなかった道筋を見つけるのにも役立ちます。
最近の NVIDIA の調査では、エージェント スタックにおけるハーネス層の重要性が強調されています。研究者らは、エージェントティック バリエーション オペレーター (AVO) を使用して、指示、明示的なルール、または明示された目標なしでエージェントを不慣れな環境に置く対話型推論ベンチマークである ARC-AGI-3 で 100% のスコアを達成しました。 AVO 研究について詳しくは、こちらをご覧ください。
この投稿では、新しいエージェント スタックの主要なレイヤー (モデル、ハーネス、メタハーネス、OpenShell などの安全なランタイム、推論インフラストラクチャ) をマッピングし、各レイヤーがリスクの軽減にどのように役立つかを説明します。また、権限をどこに置くべきか、アクセスのスコープをどのように設定すべきか、ランタイムがエージェントのアクションをどのように含めて記録できるかなど、これらのレイヤーの機能が向上し構成可能になるにつれて、どのセキュリティ プロパティが重要になるかについても学びます。
AI エージェントの動作およびインフラストラクチャの制御
エージェントを保護するためにセキュリティを再発明する必要はありません。数十年にわたるシステム セキュリティにより、最小限の権限、多層防御、分離、明示的な承認、監査可能性などの永続的な原則が提供されます。課題は、それらをどこに適用するかを決定することです

エージェントスタック。
プロンプト、モデルの保護手段、およびハーネス ロジックはすべて、エージェントが実行する可能性のある内容を決定しますが、エージェントが実行できる内容に厳密な境界を設けるわけではありません。この区別により、エージェントをガイドする動作制御とその権限を制限するインフラストラクチャ制御という 2 つの異なる種類の制御が行われます。
行動制御はエージェントの行動に影響を与える
モデルとエージェントがアクションを提案し、ハーネスがそれらを指示します。モデル、エージェント、ハーネスは連携して目標を解釈し、曖昧さを解決し、アクションを提案します。ハーネスは自然な制御ポイントです。ハーネスはループ、コンテキスト、ツール、セッションを所有し、オペレーターの意図に沿って動作を制御できます。このステアリングは貴重ですが、このレベルで実装されるすべての制御は依然としてモデルがどのように動作するかに依存します。
インフラストラクチャ制御はエージェントが実行できる内容を決定します
最終的な権限は、エージェントが実行される環境に属します。その環境は、ID を保持し、ポリシーを適用し、障害を含み、何が起こったかを記録し、同じ承認済みポリシーと検証済み状態が与えられると、毎回同じ認可決定に達します。エージェントが何をするかを予測するものではありません。それはエージェントが何ができるかを決定します。
ハーネスは、エージェントが何を試みるかをガイドします。インフラストラクチャは、エージェントが実行できる内容を制御します。どちらも必要です。権威があるのは 1 つだけです。
インフラストラクチャの施行は完全ではありません。これは、承認されたポリシーと検証された構成が再現可能な結果を​​生み出すことを意味し、エージェントは従うかどうかを選択できません。政策が依然として間違っている可能性があり、外部の結果が不確実なままである可​​能性があります。
この区分は、オープンソース エコシステムがすでに収束しているレイヤーにマッピングされています。
これらの層は機能的な役割を記述します。 1 つの製品で複数の役割を組み合わせることができ、導入では 1 つの役割を分割することができます。

oss の複数のサービス。ここで、各レイヤーは責任を指定します。セキュリティ境界は、エージェントがバイパスできない効果パスによって定義されます。
モデルはインテリジェンスを提供します。ハーネスはその知性をエージェントに変えます。ランタイムは、そのエージェントが何を実行できるかを決定します。
ハーネス レイヤーは、固定されたカテゴリではなくスペクトルです。 Codex と Claude Code は独自のハーネスですが、Pi と DeepSeek ハーネス (DSH) はプログラム可能な基板としてハーネスの多くを公開します。 DSH は、Cordis を通じて、プラグインとして構成および置換できるコア動作を可能にします。このプログラム可能性により、ハーネスはセキュリティを保証するには不十分な場所になります。変更されるように設計された層は、それ自体の変更に対して確実に制御を強制することができません。安全性を確保するためのハーネス ロジックに依存する代替方法では、モデルの動作に関する仮定がエンコードされますが、モデルが改善されるにつれてそれらの仮定は古くなります。
資格情報の範囲が狭いと潜在的な危害は制限されますが、生の資格情報をエージェントの手の届かないところに置いておくと、環境によって強制されるより強力な境界が作成されます。
起動前に AI エージェントのランタイム境界を確立する
モデル、ハーネス、ランタイム、ポリシー、推論のデプロイメントは、ますます独立して選択されるようになってきています。このアプローチは、ランタイム上でどのコンポーネントが動作するかに関係なく、ランタイムの保証が維持される場合にのみ機能します。つまり、エージェントの起動時にセキュリティ境界を確立する必要があります。
オーケストレーターは、OpenShell にランタイムを作成し、ポリシーとガバナンスを適用するように要求します。選択したハーネスはそのランタイム内で開始され、そのプラグイン、モデル コンテキスト プロトコル (MCP) プロセス、ツール、およびその他のモデル指向のコードが同じ境界内で実行されます。サブエージェントは、超えることのできない上限を持つ委任された子ランタイムを受け取りますが、オーケストレーターはランタイム内で動作します。

独自のポリシーによって管理されます。
このアプローチは、ランタイムを、ハーネスが既に実行されている後に呼び出すことができる別のツールとして扱うこととは異なります。エージェントが呼び出しを拒否できる制御は、効果的なセキュリティ制御ではありません。
エージェントスタックにおける一般的なセキュリティギャップ
多くのエージェント スタックには同じ欠陥があります。認可の決定は、エージェントまたはエージェントが読み取る信頼できないデータによって影響を受ける可能性があります。
境界が不明瞭。ルールはプロンプト、モデル、エージェント、ハーネス、ランタイム、インフラストラクチャに分割されているため、正式なバージョンを見つけるのは困難です。
過剰なアクセス。エージェントは、現在のタスクが必要とするものを超えて、永続的な、多くの場合長期間有効な資格情報または権限を受け取ります。
信頼できないデータをコントロールとして使用します。ドキュメント、メッセージ、ツールの結果、およびメモリは、指示として承認されずにアクションをリダイレクトできます。
制御されていない外部効果。許可された API は、データの移動、コンピューティングの作成、または意図されたコントロールの外でのエフェクトのトリガーを行うことができます。
複合的な失敗。エージェントは委任し、メモリを共有し、ピアを呼び出します。そのため、1 つの間違いが急速に連鎖する可能性があります。
不完全な監査証拠。承認は曖昧で、アクセスの取り消しは遅く、記録はインシデントの説明や復旧のサポートに十分ではありません。
強制可能なエージェントセキュリティのための設計ルール
5 つの設計ルールは、セキュリティに関する決定をエージェントの制御外に保つのに役立ちます。
上記は提案します。以下が決定します。モデル、エージェント、ハーネス、ツール、またはメモリ システムは、それ自体に権限を付与しません。
権威あるポリシーの場所。ポリシーを境界線より下に保ちます。ポリシーを意識した境界線を超えた計画は有用ですが、助言的なものです。
あらゆる効果を確認してください。すべてのファイル、プロセス、ネットワーク リクエスト、API 呼び出し、データ操作、リソース割り当て、通信、デバイス アクションを制御します。
ジャストインタイムアクセス。資格情報と機能は範囲が狭く、有効期間が短く、削除が簡単である必要があります

。
隔離と回復。各エージェントを隔離し、アクセスを迅速に取り消し、回復して記録を保存します。
エージェント向けの階層化されたセキュリティ モデル
OSI モデルと同様に、このエージェント スタックは各レイヤーに 1 つのジョブと明確なインターフェイスを割り当てます。上位層は、その下の制御層を再定義することなく変更できます。
図 1. 階層化された AI エージェント スタックにより、インフラストラクチャによるセキュリティ制御から動作コンポーネントが分離されます。
セキュリティ境界の仕組み
境界は、すべてのリクエストが一貫して評価される場合にのみ有効です。これを可能にするのは 3 つの要件です。
境界より上のすべてのコンポーネントを信頼できないものとして扱います。間違い、侵害、または敵対的である可能性があり、その要求自体には何の権限もありません。
境界より下のコントロールに権限を与えます。これらのレイヤーは、各リクエストを ID にバインドし、ポリシーを適用し、決定を強制します。
リスクシグナルは権限を減らす目的でのみ使用してください。異常スコアなどの信号により、より厳格な制御がトリガーされる可能性がありますが、追加のアクセスを許可してはなりません。
外部状態を変更するすべてのアクションは、境界の下にあるポリシー層と強制層を通過する必要があります。レイヤ 5 ～ 7 がこれらの制御をバイパスできるパスは、アーキテクチャ上の欠陥です。
エージェントのワークロード用の 4 つのセキュリティ プロファイル
4 つのプロファイルはすべて、同じスタック、境界、インターフェイスを使用します。それぞれが、付与された権限、潜在的な影響、敵対的な動作の可能性に基づいて、異なる制御を適用します。
重要。レッドチーム エージェントの運用アクセスは、通常の運用エージェントに付与されるアクセスよりも例外的かつ狭いものである必要があり、広いアクセスではありません。
リスクの増加に応じてエージェントのセキュリティ管理がどのように変化するか
エージェントの権限が増大し、その行動の潜在的な影響が増大するにつれて、5 つの領域での管理を強化します。
より狭い権限。助成金はこうなるはずです

リスクが高まると寿命が短くなります。
新鮮な決断。それぞれのアクションに近づいてポリシーを再評価します。
監視強化。影響の大きい作業にはライブ監視を追加します。
より速い回復。アクセスの取り消し、隔離、ロールバックを計画します。
独立した証拠。不変レコードをセキュリティ境界の下に保持します。
あらゆるリスクレベルにおけるセキュリティ要件
リスクが増加するにつれて管理は厳しくなりますが、次のセキュリティ要件はすべてのプロファイルで一貫している必要があります。
エージェントは自分自身にアクセスを許可することはありません。制御はエージェント プロセスの外部で、エージェントの制御を超えて適用されます。これはあらゆるレベルに当てはまります。
範囲内のすべての影響の大きい効果は、施行ポイントを超えます。チェックは、アクションを実行するシステムで行われます。
システムは安全に失敗します。コントロールが欠落しているか古い場合は、事前に承認されたより安全な状態が選択されます。物理システムや可用性が重要なシステムの場合、その状態では、突然停止するのではなく、制御された動作が必要になる場合があります。
セキュリティに関する主張の範囲は引き続き定められています。対象となる正確なパス、行われた仮定、およびスタックの外に残された除外項目を記述します。
市場内学習で AI の形成を支援
AI モデルの構築、AI システムの展開、クラウド インフラストラクチャの運用、セキュリティ研究の実施、ガバナンスと標準の開発など

[切り捨てられた]

## Original Extract

As AI agents become more capable and operate over longer horizons, building security and trust into the applications they power becomes increasingly important.

Where Security Fits in an AI Agent Stack | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Agentic AI / Generative AI
Where Security Fits in an AI Agent Stack
AI safety and security teams at NVIDIA explore where security controls belong as AI agents take on increasingly complex work.
Like
Dislike
Recent incidents involving frontier AI agents highlight the importance of clearly defined security boundaries within the agent stack, as agents with creative problem-solving abilities have demonstrated the capacity to bypass intended restrictions.
The agent stack is composed of distinct layersmodels, harnesses, meta-harnesses, secure runtimes like NVIDIA OpenShell, and inference infrastructurewith security controls most effectively enforced at the runtime and infrastructure layers, rather than within modifiable harness logic.
Effective agent security relies on principles such as least privilege, isolation, just-in-time access, and authoritative policy enforcement below the agent boundary, ensuring that agents cannot grant themselves authority or bypass controls, and that all impactful actions are consistently evaluated and auditable.
AI-generated content may summarize information incompletely. Verify important information. Learn more
As AI agents become more capable and operate over longer horizons, building security and trust into the applications they power becomes increasingly important. Drawing on work with NVIDIA OpenShell, agent developers, open-source projects, and partners across the ecosystem, AI safety and security teams at NVIDIA offer their perspective on the emerging agent stack—including the role of each layer and where security should live.
Recent reports underscore why the placement of security controls matters. Within a few weeks this summer, OpenAI, Anthropic, and the UK AI Security Institute each reported frontier agents operating beyond their intended boundaries. The reported behaviors included exploiting an unexpected path out of lab environments to the open internet, gaining unauthorized access to other companies’ systems, and taking unsanctioned actions involving people and infrastructure. These cases involved long-horizon agents running with reduced model safeguards. But they point to the same design challenge: the capabilities that enable agents to solve problems creatively and pursue complex goals can also help them find paths that their original instructions did not anticipate.
Recent NVIDIA research underscores the importance of the harness layer in the agent stack. Using Agentic Variation Operators (AVO), researchers achieved a 100% score on ARC-AGI-3, an interactive reasoning benchmark that places agents in unfamiliar environments without instructions, explicit rules, or stated goals. Learn more about the AVO research .
This post maps the main layers of the emerging agent stack—models, harnesses , meta-harnesses, secure runtimes such as OpenShell, and inference infrastructure—and explains how each layer can help reduce risk. You’ll also learn which security properties become critical as these layers grow more capable and composable, including where authority should live, how access should be scoped, and how the runtime can contain and record an agent’s actions.
Behavioral and infrastructure controls for AI agents
Securing agents doesn’t require reinventing security. Decades of systems security provide durable principles, including least privilege, defense in depth, isolation, explicit authorization, and auditability. The challenge is determining where to apply them in an agent stack.
Prompts, model safeguards, and harness logic all shape what an agent is likely to do, but they don’t create a hard boundary around what it can do. This distinction leads to two different kinds of control: behavioral controls that guide the agent and infrastructure controls that limit its authority.
Behavioral controls influence agent actions
The model and agent propose actions, and the harness directs them. Together, the model, agent, and harness interpret goals, work through ambiguity, and propose actions. The harness is the natural control point: it owns the loop, the context, the tools, and the session, and it can steer behavior toward what the operator intends. That steering is valuable, but every control implemented at this level still depends on how the model will behave.
Infrastructure controls determine what an agent can do
Final authority belongs to the environment in which the agent runs in. That environment holds identity, enforces policy, contains failures, records what happened, and reaches the same authorization decision every time, given the same approved policy and verified state. It doesn’t estimate what an agent will do. It determines what an agent can do.
The harness guides what an agent tries. The infrastructure controls what an agent can do. Both are necessary; only one is authoritative.
Infrastructure enforcement is not infallible. It means approved policy and verified configuration produce repeatable outcomes, and the agent cannot choose whether to comply. Policy can still be wrong, and external outcomes can remain uncertain.
This division maps onto the layers the open-source ecosystem is already converging on:
These layers describe functional roles. One product may combine several roles, and a deployment may split one role across multiple services. Here, each layer names a responsibility. The security boundary is defined by the effect paths that the agent cannot bypass.
The model supplies intelligence; the harness turns that intelligence into an agent; the runtime determines what that agent is allowed to do.
The harness layer is a spectrum rather than a fixed category. Codex and Claude Code are opinionated harnesses, while Pi and DeepSeek Harness (DSH) expose more of the harness as a programmable substrate. Through Cordis, DSH enables core behaviors that can be composed and replaced as plugins. This programmability makes the harness a poor place for a security guarantee: a layer designed to be modified cannot reliably enforce controls against its own modification. The alternative—relying on harness logic for safety—encodes assumptions about model behavior, and those assumptions go stale as models improve.
A narrowly scoped credential limits potential harm, but keeping the raw credential out of the agent’s reach creates a stronger boundary enforced by the environment.
Establish the AI agent runtime boundary before launch
Models, harnesses, runtimes, policies, and inference deployments are increasingly selected independently. This approach only works if the runtime’s guarantees hold regardless of which components operate above it. That means a security boundary must be established when the agent launches.
An orchestrator asks OpenShell to create a runtime and enforce policies and governance. The selected harness starts inside that runtime, and its plugins, Model Context Protocol (MCP) processes, tools, and other model-directed code run inside the same boundary. Subagents receive delegated child runtimes with ceilings they can’t exceed, while the orchestrator operates inside a runtime governed by its own policy.
This approach is different from treating the runtime as another tool that a harness can invoke once it’s already running. A control that the agent can decline to invoke is not an effective security control.
Common security gaps in agent stacks
Many agent stacks share the same flaw: authorization decisions can be influenced by the agent or by untrusted data it reads.
Unclear boundaries. Rules are split across prompts, models, agents, harnesses, runtimes, and infrastructure, so the authoritative version is hard to find.
Excessive access. The agent receives standing, often long-lived credentials or permissions beyond what the current task needs.
Untrusted data as control. Documents, messages, tool results, and memory can redirect action without being authorized as instructions.
Uncontrolled external effects. An allowed API can move data, create compute, or trigger effects outside the intended controls.
Compounding failures. Agents delegate, share memory, and call peers, so one mistake can become a fast cascade.
Incomplete audit evidence. Approvals are vague, access is slow to revoke, and the record is not sufficient to explain an incident or support recovery.
Design rules for enforceable agent security
Five design rules help keep security decisions outside the agent’s control.
Above proposes; below decides. No model, agent, harness, tool, or memory system grants itself authority.
Authoritative policy location. Keep policy below the line. Policy-aware planning above the line is useful, but advisory.
Check every effect. Control every file, process, network request, API call, data operation, resource allocation, communication, and device action.
Just-in-time access. Credentials and capabilities should be narrow, short-lived, and easy to remove.
Isolation and recovery. Isolate each agent, revoke access quickly, recover, and preserve the record.
A layered security model for agents
Like the OSI model, this agent stack assigns each layer one job and a clear interface. Higher layers can change without redefining the control layer below them.
Figure 1. A layered AI agent stack separates behavioral components from infrastructure-enforced security controls
How the security boundary works
The boundary is effective only if every request is evaluated consistently. Three requirements make this possible.
Treat every component above the boundary as untrusted. It may be mistaken, compromised, or adversarial, and its requests carry no authority on their own.
Make the controls below the boundary authoritative. These layers bind each request to an identity, apply policy, and enforce the decision.
Use risk signals only to reduce authority. Signals such as anomaly scores may trigger tighter controls, but they must never grant additional access.
Every action that changes the external state must pass through the policy and enforcement layers below the boundary. Any path that allows Layers 5-7 to bypass those controls is an architectural defect.
Four security profiles for agent workloads
All four profiles use the same stack, boundary, and interfaces. Each applies different controls based on the authority granted, the potential impact, and the likelihood of adversarial behavior.
Important. Production access for a red-team agent should be exceptional and narrower, not broader, than access granted to an ordinary production agent.
How agent security controls change as risk increases
As an agent gains more authority and the potential impact of its actions grows, strengthen controls in five areas.
Narrower authority. Grants should become shorter-lived as risk rises.
Fresh decisions. Reevaluate policy closer to each action.
Stronger oversight. Add live supervision for high-impact work.
Faster recovery. Plan for access revocation, quarantine, and rollback.
Independent evidence. Keep immutable records below the security boundary.
Security requirements at every risk level
Although controls become stricter as risk increases, the following security requirements should remain consistent across every profile.
The agent never grants itself access. Controls are enforced outside the agent process and beyond the agent’s control. This holds at every level.
Every in-scope, high-impact effect crosses an enforcement point. The check occurs in the system that performs the action.
The system fails safely: a missing or stale control selects a preapproved safer state. For physical and availability-critical systems, that state may require controlled operation rather than an abrupt stop.
Security claims remain scoped . State the exact paths covered, assumptions made, and exclusions left outside the stack.
Help shape AI with in-market learning
Whether you build AI models, deploy AI systems, operate cloud infrastructure, conduct security research, or develop governance and standard

[truncated]

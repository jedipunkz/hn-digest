---
source: "https://www.cerbos.dev/blog/authorization-for-ai-agents-what-to-build-before-eu-ai-act-deadline"
hn_url: "https://news.ycombinator.com/item?id=48356180"
title: "What to build before the EU AI Act deadline"
article_title: "Authorization for AI Agents: What to Build Before the EU AI Act Deadline | Cerbos"
author: "mooreds"
captured_at: "2026-06-02T00:41:00Z"
capture_tool: "hn-digest"
hn_id: 48356180
score: 3
comments: 0
posted_at: "2026-06-01T12:53:17Z"
tags:
  - hacker-news
  - translated
---

# What to build before the EU AI Act deadline

- HN: [48356180](https://news.ycombinator.com/item?id=48356180)
- Source: [www.cerbos.dev](https://www.cerbos.dev/blog/authorization-for-ai-agents-what-to-build-before-eu-ai-act-deadline)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T12:53:17Z

## Translation

タイトル: EU AI 法の期限までに構築すべきもの
記事のタイトル: AI エージェントの認可: EU AI 法の期限までに構築すべきもの |セルボス
説明: オーケストレーション層のランタイム ポリシーの意味、エージェントからツールへの層がエージェント AI ガバナンスに欠けているカテゴリである理由、および EU AI 法のどの条項が実際にインフラストラクチャ ベンダーに適用されるのか。 CTO とセキュリティ リーダー向けの、エージェントの在庫管理、アイデンティティのスポンサー、および元のセキュリティ リーダー向けの実践的な手順
[切り捨てられた]

記事本文:
AI エージェントの認可: EU AI 法の期限前に何を構築するか | Cerbos 🛡️ 2026 年の IAM セキュリティ チェックリスト ➔ ここからダウンロード 4.4k
製品リソース 価格 Cerbos を試す
技術ブログ / AI エージェントの認可: EU AI 法の期限までに構築すべきもの ガイド AI エージェントの認可: EU AI 法の期限までに構築すべきもの
7 min read KuppingerCole の Jonathan Care の一文がありますが、それを聞いて以来ずっと頭の中で回っています。
フレームワークはモデルが何を言うかを管理します。エージェントの行動を制御するものはほとんどありません。
それがギャップです。私が話し合ってきたチームは、さまざまな角度からその目標に到達し続けており、これは通常、カテゴリーが形成されつつある兆候です。 CTO がエージェントを運用環境に置くことを望んでおり、セキュリティが承認されないという理由で現れる人もいます。 EU AI 法が法律に不安を感じているという理由で参加する人もいます。下も同じ隙間。
3 つのギャップ、そのうち 1 つだけ名前が付けられている
3 つのレイヤーがあり、それぞれに新たな頭字語が付けられています。頭字語は、その下の形状ほど重要ではありません。
1つ目はアイデンティティです。すべてのエージェントには、クラスごとではなくインスタンスごとに独自の ID が必要です。有効期間が短い認証情報は、単一のツール呼び出しに限定されます。指定された人間のスポンサーに関連付けられたライフサイクル。現在、ほとんどの企業は、「エージェント」ごとに 1 つの長期有効な API キーを発行し、生成されたすべてのインスタンスを同じアクターとして扱っています。これは、あるエージェントが別のエージェントを生成した瞬間に動作を停止します。 (これについては、人間以外のアイデンティティとスポンサーと結びついたライフサイクル パターンの下で書きました。)
2つ目は監査です。エージェントがサブエージェントに委任すると、既存の監査証跡はすべて破棄されます。誰が同意したのか。本来の目的は何だったのか。同意はホップ後に残りましたか。今日のログはそのどれにも答えていません。サービス アカウントが何かをしたと言っています。どの人間がどのような方法でそれを許可したかは教えてくれません。

どのデータで、どのような目的で。流通管理は欠けている部分ですが、これは人々が信じている以上に大きな問題です。
3 つ目は、まだ成熟したカテゴリが存在しないものです。ケアはそれをオーケストレーションと呼びます。エージェントからエージェント、およびエージェントからツールへのレイヤー。エージェント外部のツール ゲート。エージェント間の信頼の強制。ポリシー プレーンに到達できない場合のフェールクローズされたランタイム。現在、このレイヤーは、エージェント フレームワークが公開することを選択したものにすぎません。多くの場合、エージェント自体が呼び出すことができるツールを決定します。これは、子供にデザートを許可するかどうかを尋ねるセキュリティ モデルです。
最初の 2 つは、たとえ実装がまだであっても、IAM 業界がどのように考えるかを知っている問題です。 3 つ目は、誰も独自のカテゴリとして名前を付けていないものです。名前を付けることがそれが現実になる瞬間であり、エージェント トラフィックのランタイム層でアーキテクチャ上の作業のほとんどが行われることになると思います。
「オーケストレーション層のランタイムポリシー」の実際の意味
これは私が取り組んでいる層なので、具体的に説明します。
エージェントとツール間のすべての呼び出しは決定です。このコンテキストにおいて、この人間に代わって動作するこのエージェントが、これらの引数を使用してこのツールを呼び出すことを許可されるべきか。エージェントがその決定を下すべきではありません。別のポリシー エンジンは、エージェントの推論ループの外側に位置し、通話が完了する前に通話を評価する必要があります。
これは、エージェントからツールへの呼び出しのためのランタイム ポリシー エンジンです。これは、私が何年も書き続けてきた外部化された承認パターンと同じですが、エージェントが新しいプリンシパル タイプとしてスロットに組み込まれ、ツールが新しいリソース タイプとしてスロットに組み込まれるだけです。切り離されました。外部化された。エンジンに到達できない場合はフェールクローズします。同じPEP/PDP形状、エージェント形状の表面。
これがガバナンスにとって重要である理由は、これがなければ他のすべての企業が

ntrol はエージェント内にあります。プロンプトレベルのガードレール。フレームワークレベルの許可リスト。システムは「返金ツールを呼び出さないでください」というメッセージを表示します。これらすべては、制約しようとしているのと同じ確率システム内に存在します。受刑者によって書かれた、受刑者に関する管理。
エージェントからポリシーを取り出すと、セキュリティ チームが推論できるものが得られます。すべてのエージェントで同じエンジン。同じ監査ログ。同じ変更管理。同じロールバックの話。エージェントは、エージェントが何を実行できるかを決定する特権を失います。これはまさにユーザーが望むことです。
8月の締め切りと、それが些細な話である理由
これが突然緊急になった理由は、EU AI 法にあります。第 9 条はリスク管理について規定しています。第 10 条ではデータ ガバナンスについて説明します。第 12 条では、システムの存続期間にわたる自動記録保持について説明します。第 13 条では、配備者に対する透明性について規定しています。これらは共に、高リスク AI システムが実証しなければならないことの運用中核を形成します。
これらの義務は、Cerbos やエージェント スタック内の他のインフラストラクチャ ベンダーではなく、AI システム プロバイダーにあります。 「AIコンプライアンス」のマーケティングの山はすでに深くなっているので、その点には注意したい。私たちはあなたにとって第9条を満たしていません。当社はお客様を AI 法に準拠させるわけではありません。そんなことを言う人はあなたに物語を売りつけているのです。
Cerbos が役立つのは、レイヤー内でランタイム コントロールを動作させずにデモンストレーションするのがほぼ不可能な部分です。エージェントとツールの境界で意思決定を強制しないリスク管理プロセスは文書です。エージェントが読み取ることができるデータをフィルタリングしないデータ ガバナンスはガバナンスではありません。実行時に行われた承認の決定と、それを生成したポリシーを見逃す自動ログは、管理の連鎖ではありません。同じことを省略するデプロイ担当者に対する透明性は透明性ではありません。
当初は2026年8月でした

■ 付属書 III に基づく高リスク義務の発動日。 2025年11月の欧州委員会のデジタルオムニバス提案は2027年12月までの延期を促し、評議会と議会は5月初旬にこれに関して暫定合意に達した。そのため日付がずれる可能性がございます。アーキテクチャ上の義務はそうではありません。
これらの義務を強制する必要があるランタイム層は、ほとんどのスタックにはまだ存在しません。期限があると緊急性が高まります。名前が付けられたカテゴリーがそのカテゴリーに形を与えます。
ベンダー、アナリスト、エンド ユーザーは、異なる出発点から同じアーキテクチャの形に到達します。
エージェントごとにスコープが設定され、スポンサーと結びついたライフサイクルを持つ ID。委任後に存続する監査チェーン。エージェント外部のツール呼び出しを制御するランタイム ポリシー プレーン。 Care の 8 つの質問からなるベンダー チェックリストは、これら 3 つの領域の最上位にあります。マルティン・クッピンガーの地殻変動の枠組みも同じ考えに基づいています。 CoSAI エージェント IAM リファレンスも同じ内容をカバーしています。私が OpenID Foundation で共同議長を務める AuthZEN は、ポリシー決定部分のプロトコル層を構築しています。
多数の独立したスレッドが同じアーキテクチャに到達すると、そのカテゴリに名前が付けられます。実装もそれに追いつく必要があります。
セキュリティ リーダーがこの四半期に行うべきこと
これからやるべきいくつかの実践的なことを、おおよその順序で説明します。
在庫。ほとんどの組織は、すでに何人のエージェントを抱えているかについて否定しています。 3つのチームがすでに実施しているが、首脳陣は我々は実施していないと言っている。見つけてください。シャドウ AI をシャドウ IT のように扱いますが、より高速に動作します。
すべてのエージェントのスポンサーになります。各エージェントには、エージェントの存在がそのライフサイクル ステータスに依存する名前付きの人間の所有者が必要です。サリーが去ると、サリーのエージェントは止まります。これは技術的な変化であると同時に、ガバナンスの変化でもあります。
ポリシーをエージェントの外に移動します。エージェントを評価できるランタイム ポリシー エンジンを選択してください

t-to-tool は外部から呼び出します。通常のアプリケーション トラフィックの承認を外部化する理由は、プリンシパル タイプが交換されるエージェント トラフィックにも当てはまります。
監査チェーンを配線します。すべてのエージェントのアクションには、リーフ ツールの呼び出しに至るまで、人間のスポンサー、本来の目的、およびそれを許可するポリシーの決定が反映される必要があります。それがなければ、たとえ締め切りがずれても、第86条の説明可能性の問題に答えることはできない。
フェールクローズ動作をテストします。ポリシー エンジンに到達できない場合はどうなりますか。答えが「とにかくエージェントが実行する」の場合は、ランタイムがフェールオープンしていることになります。これはこの層で最悪の部類のバグです。
これからの12か月は騒々しいものになるだろう。すべてのインフラストラクチャ ベンダーは、ランタイムのストーリーを主張し始めるでしょう。既存の PAM または API ゲートウェイ製品を再スキニングする場合もあります。いくつかは本物の新品になります。監視すべきシグナルは、ポリシーがエージェントから切り離されているかどうかです。ポリシーがまだエージェント内に存在する場合は、別のものを見ていることになります。
私たちを含む誰かがスライドで AI 法の遵守を主張するよりも、そのアーキテクチャを正しく理解することを望んでいます。期限は強制機能です。実際の作業はランタイムです。
Cerbos を試して、エージェントからツールへの呼び出しのランタイム ポリシー エンジンが実際にどのようなものかを確認してください。期限までに実際にどのようなランタイム コントロールを導入できるかを実際に確認したい場合は、コールを予約してください。
AI エージェント用のポリシーベースのガードレールにより、組織が実行時の承認と完全な意思決定ログを使用してエージェント ワークフロー、RAG パイプライン、MCP サーバーを保護する方法を検討できます。
AI のゼロ トラスト: MCP サーバーのセキュリティ保護 (電子書籍) では、エージェントとツールの間にポリシーを置くための実践的なチュートリアルが記載されています。
同じパターンがどのように拡張されるかについて、人間以外のアイデンティティに対するきめ細かい認証 (ウェビナー)

d サービスアカウントとエージェント間で。
エージェントからツールへの呼び出しのためのランタイム ポリシー エンジンとは何ですか?
AI エージェントにとって、スポンサーと結びついたライフサイクルは何を意味しますか?
認可ポリシーを AI エージェント自体から移動するのはなぜですか?
AI エージェントのフェールクローズされたランタイム動作とは何ですか?
エージェントからサブエージェントへの委任は、既存の監査証跡をどのように破るのでしょうか?
私たちのチームが作成した最初の Cerbos ポリシーを入手してください。
セッションを予約して要件について話し合い、作業方針を決定します。
ビジネス要件と認可ポリシーのマッピング
eBook: AI のゼロトラスト、MCP サーバーの保護
Cerbos Playground で実験、学習、プロトタイプを作成する
eBook: 外部化された認証を採用する方法
認可プロバイダーとソリューションを評価するためのフレームワーク
コンプライアンスの維持 – 知っておくべきこと
数千人の開発者に加わりましょう |機能とアップデート |月に 1 回 |スパムはなく、グッズだけです。
セルボスとは何ですか？ Cerbos は、ゼロ トラスト環境および AI 搭載システム向けのエンドツーエンドのエンタープライズ認証管理プラットフォームです。アプリ、API、AI エージェント、MCP サーバー、サービス、ワークロード全体にわたって、きめ細かくコンテキストに応じた継続的な承認を強制します。
Cerbos は以下で構成されます。 オープンソースのポリシー決定ポイント。施行ポイントの統合。アーキテクチャ全体で統一されたポリシーベースの認可を調整する、一元管理されるポリシー管理ポイント (Cerbos Hub)。エンリッチメントおよびオーケストレーション ポイント (Cerbos Synapse) は、既存のシステムから ID、リソース、関係データを収集し、すべての認可決定の前に完全なコンテキストをポリシー エンジンに提供します。 Cerbos 認証により、最小限の権限を適用し、アクセス決定に対する完全な可視性を維持します。

## Original Extract

What runtime policy at the orchestration layer means, why the agent-to-tool layer is the missing category in agentic AI governance, and which EU AI Act articles actually apply to infrastructure vendors. Practical steps for CTOs and security leads on inventorying agents, sponsoring identities, and ex
[truncated]

Authorization for AI Agents: What to Build Before the EU AI Act Deadline | Cerbos 🛡️ IAM security checklist for 2026 ➔ Download it here 4.4k
Product Resources Pricing Try Cerbos
Tech blog / Authorization for AI agents: What to build before the EU AI Act deadline Guide Authorization for AI agents: What to build before the EU AI Act deadline
7 min read There's a line from Jonathan Care at KuppingerCole that's been bouncing around in my head since I heard it.
Frameworks govern what models say. Almost nothing governs what agents do.
That's the gap. The teams I've been talking to keep landing on it from different angles, which is usually a sign a category is forming. Some show up because their CTO wants agents in production and security won't sign off. Some show up because the EU AI Act is making someone in legal nervous. Same gap underneath.
Three gaps, only one of them named
There are three layers, each with an emerging acronym attached. The acronyms matter less than the shape underneath.
The first is identity. Every agent needs its own identity, per instance, not per class. Short-lived credentials scoped to a single tool call. Lifecycle tied to a named human sponsor. Most enterprises today are issuing one long-lived API key per "the agent" and treating every spawned instance as the same actor. That stops working the moment one agent spawns another. (We've written about this under non-human identity and the sponsor-tied lifecycle pattern.)
The second is audit. Once an agent delegates to a sub-agent, every existing audit trail breaks. Who consented. What was the original purpose. Did the consent survive the hop. Today's logs answer none of that. They tell you a service account did a thing. They don't tell you which human authorized it, through which chain, for what purpose, on what data. The chain of custody is the bit that's missing, and it's a bigger deal than people are giving it credit for.
The third is the one with no mature category yet. Care calls it orchestration. The agent-to-agent and agent-to-tool layer. Tool gating outside the agent. Inter-agent trust enforcement. Fail-closed runtime when the policy plane is unreachable. Today this layer is just whatever the agent framework chose to expose. Often the agent itself decides what tools it can call, which is the security model of asking a child whether they're allowed dessert.
The first two are problems the IAM industry knows how to think about, even if the implementations aren't there yet. The third is the one nobody has been naming as its own category. I think naming it is the moment it becomes real, and the runtime layer for agent traffic is where most of the architectural work is about to happen.
What "runtime policy at the orchestration layer" actually means
This is the layer I work in, so I'll be specific.
Every agent-to-tool call is a decision. Should this agent, acting on behalf of this human, in this context, be allowed to invoke this tool with these arguments. The agent shouldn't be the thing making that decision. A separate policy engine should, sitting outside the agent's reasoning loop, evaluating the call before it goes through.
That's a runtime policy engine for agent-to-tool calls. It's the same externalized authorization pattern I've been writing about for years, just with the agent slotted in as a new principal type and the tool slotted in as a new resource type. Decoupled. Externalised. Fail-closed if the engine is unreachable. Same PEP/PDP shape , agent-shaped surface.
The reason this matters for governance is that without it, every other control sits inside the agent. Prompt-level guardrails. Framework-level allowlists. System prompts saying "do not call the refund tool." All of that lives in the same probabilistic system you're trying to constrain. Controls on the inmate, written by the inmate.
Take the policy out of the agent and you get something a security team can reason about. Same engine across all your agents. Same audit log. Same change management. Same rollback story. The agent loses the privilege of deciding what it can do, which is exactly what you want.
The August deadline, and why it's the smaller story
The reason this is suddenly urgent is the EU AI Act . Article 9 covers risk management. Article 10 covers data governance. Article 12 covers automatic record-keeping over the system's lifetime. Article 13 covers transparency to deployers. Together they form the operational core of what a high-risk AI system has to demonstrate.
These obligations sit on the AI system provider, not on Cerbos and not on any other infrastructure vendor in the agent stack. I want to be careful about that, because the "AI compliance" marketing pile is already getting deep. We do not satisfy Article 9 for you. We do not make you AI-Act-compliant. Anyone telling you that is selling you a story.
What Cerbos does help with is the part that's almost impossible to demonstrate without working runtime controls in our layer. A risk management process that doesn't enforce decisions at the agent-to-tool boundary is a document. Data governance that doesn't filter what an agent can read isn't governance. Automatic logging that misses the authorization decisions taken at runtime, and the policy that produced them, isn't a chain of custody. Transparency to deployers that omits the same isn't transparency.
Originally August 2026 was the activation date for the high-risk obligations under Annex III. The Commission's Digital Omnibus proposal in November 2025 pushed for a deferral to December 2027, and Council and Parliament reached provisional agreement on that in early May. So the date may slip. The architecture obligation doesn't.
The runtime layer where those obligations have to be enforced still doesn't exist in most stacks. The deadline gives it urgency. The category being named gives it shape.
Vendors, analysts, and end users are landing on the same architectural shape from different starting points.
Identity, scoped per agent, with a sponsor-tied lifecycle. Audit chains that survive delegation. A runtime policy plane that gates tool calls outside the agent. Care's eight-question vendor checklist sits on top of these three areas. Martin Kuppinger's tectonic-shifts framework lands on the same. The CoSAI agentic IAM reference covers the same ground. AuthZEN , which I co-chair at the OpenID Foundation, is wiring up the protocol layer for the policy-decision part.
When that many independent threads land on the same architecture, the category has been named. The implementations now have to catch up.
What security leads should do this quarter
A few practical things to do now, in roughly the order I'd do them.
Inventory. Most organisations are in denial about how many agents they already have. Three teams already do, while leadership says we don't. Find them. Treat shadow AI like shadow IT, but moving faster.
Sponsor every agent. Each agent should have a named human owner whose lifecycle status the agent's existence depends on. If Sally leaves, Sally's agents stop. That's a governance change as much as a technical one.
Move policy out of the agent. Pick a runtime policy engine that can evaluate agent-to-tool calls externally. The reasons to externalize authorization for your normal application traffic apply to agent traffic too, with the principal type swapped.
Wire the audit chain. Every agent action should carry the human sponsor, the original purpose, and the policy decision that authorized it, all the way down to the leaf tool call. Without that you can't answer the Article 86 explainability question even if the deadline moves.
Test fail-closed behaviour. If your policy engine is unreachable, what happens. If the answer is "the agent does it anyway," you have a runtime that fails open. That's the worst class of bug in this layer.
The next twelve months are going to be loud. Every infrastructure vendor will start claiming a runtime story. Some will be re-skinning existing PAM or API gateway products. A few will be genuinely new. The signal to watch is whether the policy is decoupled from the agent. If the policy still lives inside the agent, you're looking at something else.
I'd rather we got that architecture right than that anyone, including us, got to claim AI Act compliance on a slide. The deadline is a forcing function. The actual work is the runtime.
Try Cerbos to see what a runtime policy engine for agent-to-tool calls looks like in practice, or book a call if you want to walk through what runtime controls you can realistically put in place before the deadline.
Policy-based guardrails for AI agents to explore how your organization can secure agentic workflows, RAG pipelines, and MCP servers with runtime authorizationl and full decision logging.
Zero Trust for AI: Securing MCP servers (eBook) for a practical walkthrough of putting policy between agents and tools.
Fine-grained authorization for non-human identities (Webinar) for how the same patterns extend across service accounts and agents.
What is a runtime policy engine for agent-to-tool calls?
What does sponsor-tied lifecycle mean for AI agents?
Why move authorization policy out of the AI agent itself?
What is fail-closed runtime behaviour for AI agents?
How does agent-to-sub-agent delegation break existing audit trails?
Get your first Cerbos policy written by our team.
Book a session to talk through your requirements and walk away with a working policy.
Mapping business requirements to authorization policy
eBook: Zero Trust for AI, securing MCP servers
Experiment, learn, and prototype with Cerbos Playground
eBook: How to adopt externalized authorization
Framework for evaluating authorization providers and solutions
Staying compliant – What you need to know
Join thousands of developers | Features and updates | 1x per month | No spam, just goodies.
What is Cerbos? Cerbos is an end-to-end enterprise authorization management platform for Zero Trust environments and AI-powered systems. It enforces fine-grained, contextual, and continuous authorization across apps, APIs, AI agents, MCP servers, services, and workloads.
Cerbos consists of: an open-source Policy Decision Point; Enforcement Point integrations; a centrally managed Policy Administration Point (Cerbos Hub) that coordinates unified policy-based authorization across your architecture; and an enrichment and orchestration point (Cerbos Synapse) which gathers identity, resource, and relationship data from your existing systems and delivers complete context to the policy engine before every authorization decision. Enforce least privilege & maintain full visibility into access decisions with Cerbos authorization.

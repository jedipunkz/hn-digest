---
source: "https://signal-memo.com/memo-salesforce-and-sap-are-making-opposite-bets-about-how-agents-will-use-enterprise-software-the-rest-of-the-stack-is-about-to-have-to-pick-a-side/"
hn_url: "https://news.ycombinator.com/item?id=48442889"
title: "Substrate vs. Broker: Two Emerging Strategies for Enterprise AI"
article_title: "Memo: Salesforce and SAP Are Making Opposite Bets About How Agents Will Use Enterprise Software. The Rest of the Stack Is About to Have to Pick a Side."
author: "alex-ivan"
captured_at: "2026-06-08T10:41:45Z"
capture_tool: "hn-digest"
hn_id: 48442889
score: 1
comments: 0
posted_at: "2026-06-08T08:58:37Z"
tags:
  - hacker-news
  - translated
---

# Substrate vs. Broker: Two Emerging Strategies for Enterprise AI

- HN: [48442889](https://news.ycombinator.com/item?id=48442889)
- Source: [signal-memo.com](https://signal-memo.com/memo-salesforce-and-sap-are-making-opposite-bets-about-how-agents-will-use-enterprise-software-the-rest-of-the-stack-is-about-to-have-to-pick-a-side/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T08:58:37Z

## Translation

タイトル: サブストレート vs. ブローカー: エンタープライズ AI のための 2 つの新たな戦略
記事のタイトル: メモ: Salesforce と SAP は、エージェントによるエンタープライズ ソフトウェアの使用方法について反対の賭けをしています。スタックの残りの部分は、どちらかの側を選択する必要があります。
説明: 2026 年 4 月、Salesforce は、外部エージェントが CRM 全体を直接操作できるようにする 60 以上の MCP ツールを備えた Headless 360 を発売しました。 6 か月前、SAP は、逆方向の明示的なアーキテクチャ上の決定を伴う MCP ゲートウェイを出荷しました。つまり、外部エージェントは SAP エージェントと通信する必要があり、外部エージェントは通信する必要はありません。
[切り捨てられた]

記事本文:
サインイン
購読する
アレックス・イワン著
—
2026 年 5 月 14 日
メモ: Salesforce と SAP は、エージェントがエンタープライズ ソフトウェアをどのように使用するかについて、逆の賭けをしています。スタックの残りの部分は、どちらかの側を選択する必要があります。
2026 年 4 月、Salesforce は、外部エージェントが CRM 全体を直接操作できる 60 以上の MCP ツールを備えた Headless 360 を発売しました。 6 か月前、SAP は、外部エージェントは SAP の API ではなく、SAP のエージェントと通信する必要があるという、逆方向の明示的なアーキテクチャ上の決定を伴って MCP ゲートウェイを出荷しました。両社は同じプロトコルを使用しています。彼らは、プロトコルの目的について反対の賭けをしています。どちらの発表を読んでもほとんどの読者はこれを見逃しています。この分割は、2026 年のエンタープライズ ソフトウェアにおける最も重大な戦略的相違であり、他のすべてのプラットフォーム (Microsoft、ServiceNow、Workday、Oracle、Atlassian、Snowflake、Databricks) は現在、どちらの側につくかを選択する必要がある計画会議の最中です。
2026 年 4 月 15 日、サンフランシスコの TrailblazerDX で、Salesforce は共同創設者である Parker Harris が「なぜ Salesforce に再度ログインする必要があるのか​​?」という 1 つの修辞的な質問でまとめられた内容を発表しました。付属製品である Headless 360 は、同社の 27 年の歴史の中で最も重要なアーキテクチャの変化です。 Salesforce プラットフォームのすべての機能 (データ、ワークフロー、ビジネス ロジック、コンプライアンス制御) が、REST API、MCP ツール、または CLI コマンドとして公開されるようになりました。 60 を超える新しい MCP ツールが発売時に出荷され、30 を超える事前構成されたコーディング スキルと、ブラウザなしでプラットフォームを操作するための Claude Code、Cursor、Codex、Windsurf などの外部コーディング エージェントの明示的なサポートが含まれています。価格設定はシートごとから従量制に移行しており、エージェントがユーザーの場合、シートは

デルは死んだ。
その半年前の 2025 年 11 月に、SAP は SAP TechEd で、当時としては並行発表のように見える内容を発表しました。 SAP の Integration Suite 上に構築された MCP ゲートウェイは、API とフローを、SAP のファーストパーティ エージェントである Joule が使用できる MCP ツールに変換します。 SAP は、ABAP MCP サーバーを 2026 年上半期に、SAP Commerce Cloud 用の Storefront MCP サーバーを 2026 年第 2 四半期に発表し、TechEd を通じて 2026 年初頭にかけて MCP 関連の出荷を着実に進めていくことを発表しました。報道ではこれを ERP の MCP 配管として扱いましたが、Salesforce が後に同じ話を、わずか 6 か月遅れで伝えることになります。
その読み方は間違っています。 2 社は同じプロトコルを使用して反対の戦略を実行しており、その戦略はエンタープライズ スタック全体の異なる将来を暗示しています。
この決定的な文書は、2026 年初めに SAP アーキテクチャ センターに密かに公開された SAP 独自のアーキテクチャ ガイダンスの中に埋もれています。「ベンダーとサードパーティ エージェント間の外部相互運用性について、SAP は MCP サーバーの直接公開よりも A2A を優先します。」その一文が分岐です。 Salesforce は、自社プラットフォームの MCP ツールを市場のエージェントから直接呼び出せるようにしたいと考えています。 SAP は、外部エージェントが自社のファーストパーティ エージェントである Joule と対話することを望んでいます。Joule は、発信者に代わってプラットフォームの MCP ツールを呼び出すかどうか、またその方法を決定します。
同じプロトコル。正反対の哲学。報道機関はほぼ例外なく、その結果を見逃してきた。
基板のベットとブローカーのベット
2 つの戦略には名前が必要です。現時点でアナリストの報道では、セキュリティ、価格設定、ベンダーの能力、およびどのサードパーティが参加するかについて根本的に異なる意味を持つ 2 つのアーキテクチャを説明するために同じ言葉 (「MCP 対応」) が使用されているためです。
それらを「基板ベット」と「ブローカーベット」と呼びます。
基板の賭け、w

Salesforce が実行されているのは、エージェントのエンタープライズ ソフトウェアの耐久層は、すべてのエージェントが呼び出すプラットフォームであると述べています。プラットフォームのデータ、ワークフロー、信頼制御をそのままにして、プラットフォーム自体をあらゆるサーフェスから、あらゆるエージェントによってエージェント呼び出し可能にします。戦略的な利点は、エージェントが Salesforce によって構築されたか、顧客によって構築されたか、Cursor によって構築されたか、Anthropic によって構築されたか、まだ誰も聞いたことのないスタートアップによって構築されたか、競合他社によって構築されたかに関係なく、エージェントが実行するすべてのワークフローの根底にあることです。プラットフォームは呼ばれるものになることで勝ちます。シートが消滅すると、シートごとの価格設定も消滅します。価値の単位が呼び出しに移行するため、消費価格設定が代わりに使用されます。
SAP が運営しているブローカー ベットによれば、「耐久層は、プラットフォームの権限を保持するファーストパーティ エージェントです」とのことです。外部エージェントはプラットフォームに直接アクセスすべきではありません。彼らはプラットフォーム自身のエージェントに仕事を要求する必要があります。エージェントは、プラットフォームのガバナンス モデルに基づいて、要求を実行するかどうか、およびその方法を決定します。戦略的な利点は、他のエージェントを介して会話する必要があるエージェントであることです。 MCP は Joule が効率的に作業できるように内部に存在します。対外的には、SAP は A2A (エージェント間プロトコル) を好みます。これは、構造的にはツール インターフェイスではなく委任インターフェイスです。プラットフォームは求められるものであることで勝ちます。ジュールがアドレス可能な値の単位のままであるため、価格はエージェントのシートまたはランタイムに固定されたままになります。
これらは実装の違いではありません。これらは、エージェントのエンタープライズ ソフトウェア内で信頼境界をどこに置くべきか、エージェントが不正行為をした場合に誰が責任を負うのか、その上に形成される経済にどのサードパーティが参加するのかについてのさまざまな理論です。
基質の賭けは生態系への参加と生態系への依存を最大化します

プラットフォーム上で。ブローカーの賭けにより、エージェントが実際に実行できることに対するプラットフォームの制御が最大化されます。
どちらの賭けも擁護可能です。どちらも正しいということはあり得ません。
各社が今後5年間について暗に主張していること
2 つの賭けは、エージェント企業がどのように進化するかについての 3 つのまったく異なる主張をエンコードしています。
主張 1: 顧客の信頼はどこにあるのか?
Salesforce は、顧客の信頼は特定のエージェントではなく、プラットフォームのデータ、ワークフロー、ガバナンス層にあると確信しています。 Einstein トラストレイヤーは、どのエージェントが呼び出すかに関係なく、すべての Agentforce セッションに適用されます。データ マスキング、フィールド レベルのセキュリティ、LLM プロバイダーとのデータ保持ゼロ契約 – これらは、発信者に関係なくデータとともに送信されるプラットフォーム レベルの強制です。 Salesforce があらゆる呼び出しのゲートに位置するため、顧客は Salesforce を信頼します。呼び出し元エージェントの ID は、構造的には二次的なものです。
SAP は、プラットフォームに対して認証されたファーストパーティ エージェントに顧客の信頼が保たれることに賭けています。 Joule は、顧客が承認、範囲設定、および権限を与えたエンティティです。 SAP での活動を希望する外部エージェントは、A2A を通じて Joule に代わって活動するよう説得する必要があります。信頼境界はジュール境界であり、プラットフォームの境界ではありません。 SAP の境界内で外部エージェントと通信するのは Joule だけであるため、顧客は SAP を信頼しています。
2027 年以降、企業顧客はエージェントを承認するよりもプラットフォームを承認する方が快適になると思われる場合、Salesforce は正しいです。顧客がすべてのプラットフォームの境界に、名前を付けた責任ある単一のエージェントを要求し、外部にあるものはすべてそのエージェントの判断に従うと考えているのであれば、SAP の判断は正しいです。
主張 2: どこにあるのか

エージェントが何か間違ったことをした場合、責任は発生しますか？
エア・カナダの判例では、企業がウェブサイト上のチャットボットに対する責任を放棄できないことがすでに確立されています。展開速度を考慮すると、次の事件の波は 18 か月以内に到来すると考えられますが、より鋭い質問を投げかけます。外部エージェントが MCP 経由でエンタープライズ プラットフォームを呼び出し、有害な結果をもたらした場合、誰が責任を負うのですか?エージェントのビルダー？エージェントをデプロイしたユーザー?ツールを公開したプラットフォームは?
基板の賭けは、プラットフォームがすべての呼び出しに対して責任を共有することを暗黙的に受け入れます。 Salesforce がすべてのエージェントが呼び出す基盤である場合、Salesforce の信頼層は、アクションが実行される前の最後の強制ポイントとなります。これは、顧客が Headless 360 に統合すべき理由として、Salesforce が信頼レイヤーを売り込むことができるという機能とリスクの両方です。 Salesforce MCP ツールを通じて何か悪いことを行う、不適切に構築された外部エージェントはすべて、Salesforce の問題になります。
ブローカーの賭けは、暗黙のうちに責任を外側に押し出そうとします。外部エージェントは Joule を経由する必要があるため、SAP のファーストパーティ エージェントが説明責任を負い、外部エージェントはオペレータではなくリクエスタになります。これは、基板モデルが設計上放棄している、防御可能な法的姿勢です。今後 18 か月の間に、ツールを呼び出すサードパーティエージェントの行為に対してプラットフォームを処罰する一連の強制措置が発生すると考えるのであれば、SAP の立場ははるかに快適です。
主張 3: 価格決定力は最終的にどこに生じるのか?
Salesforce はすでに Agentforce の従量課金制に移行しており、シートごとの課金制は終了したということをより広範に示しています。基質の賭けは、w に関係なく、Salesforce が呼び出しごとに価値を抽出するモデルで終了します。

このエージェントが呼び出しを行っています。エコシステム内のすべてのエージェントは Salesforce の請求システムを実行するメーターであるため、最上位のエージェント エコシステムは成長することが奨励されています。
SAP は同等の価格設定を公表していません。エージェントは顧客関係を保持する指定エンティティであるため、ブローカーの賭けはエージェント自体に価格を設定し続けることと一致します。 Joule のシート、Joule のランタイム、Joule のワークフローごとの料金は、引き続きアドレス指定可能な単位です。外部エージェントが参加を希望する場合は、SAP の API の顧客としてではなく、Joule のサプライヤーとして参加します。 SAP が取得する呼び出しごとのフローは少なくなりますが、各顧客のエージェント スタックが何を行っているかをより明確に把握できます。
エージェント型エンタープライズ ソフトウェアは、すべてが API 経由で呼び出される少数の水平プラットフォームによって支配される世界に終わると信じている場合、価格競争では基板を賭けた方が勝つことになります。各大手エンタープライズベンダーが独自のファーストパーティエージェントを運営し、何かを成し遂げるために外部エージェントがそれらのエージェントと交渉する必要がある世界が終わると信じているなら、ブローカーの賭けが勝つことになります。
なぜスタックの残りの部分が選択しなければならないのか
Salesforce と SAP は偶然に反対の原則に収束したわけではありません。両社は、どのアーキテクチャを出荷するかについて社内で 2 年以上議論してきました。 Salesforce 独自の枠組みによれば、Salesforce は Headless 360 の発表の 2 年半前に公に基板化を行ったという。 SAP は、少なくとも 2025 年半ばからブローカー モデルに運用面で取り組んできました。各社は自社の既存の強みを生かす方針を選択した。Salesforce のエコシステムに優しいプラットフォームの DNA と、責任あるオーケストレーターとしての Joule の役割が制約ではなく機能である、管理され規制された複雑なエンタープライズ ワークフローへの SAP の緊密な統合である。
の

他の主要なエンタープライズ プラットフォームはまだ公にコミットする必要はありませんが、2026 年から 2027 年にかけて公にコミットする予定です。現在、それぞれのプラットフォームが計画会議に参加しており、基板とブローカーの決定が行われています。これまでの信号:
マイクロソフトはヘッジを行っている。 Microsoft 365 Copilot は、エンドポイントを通じて公開され、外部エージェントから呼び出し可能なサブストレート型に見えますが、Microsoft の Agent 365 フレームワークにより、Microsoft 独自のエージェントが企業内の特権オペレーターとして位置付けられます。 Microsoft の Agent Builder の今後 18 か月と、A2A 対 MCP への投資の成熟度に注目してください。 Microsoftは、より明確にどちらかの側に着地する必要があるだろう。
ServiceNow は構造的にブローカー モデルに引き寄せられています。その中核となる資産はワークフロー グラフであり、そのグラフを任意の外部エージェントに公開すると、それを取り巻くコンサルティング実装の経済性が損なわれます。 ServiceNow の Now Assist が、すべてを仲介する Joule スタイルのファーストパーティ エージェントとして強化されることを期待してください。
Workday は ServiceNow と同じ構造的魅力を持っていますが、それを強制する立場は弱いです。人事と財務のワークフローはクロスプラットフォームすぎるため、純粋なブローカー モデルを維持できません。 Workday はおそらく最終的にはハイブリッドになるでしょうが、ハイブリッドはブローカーに偏ることになります。
Snowflake と Databricks は構造的にプルです

[切り捨てられた]

## Original Extract

In April 2026, Salesforce launched Headless 360 with sixty-plus MCP tools that let any external agent operate the entire CRM directly. Six months earlier, SAP shipped its MCP Gateway with an explicit architectural decision in the opposite direction: external agents should talk to SAP's agents, not t
[truncated]

Sign in
Subscribe
By Alex Ivan
—
14 May 2026
Memo: Salesforce and SAP Are Making Opposite Bets About How Agents Will Use Enterprise Software. The Rest of the Stack Is About to Have to Pick a Side.
In April 2026, Salesforce launched Headless 360 with sixty-plus MCP tools that let any external agent operate the entire CRM directly. Six months earlier, SAP shipped its MCP Gateway with an explicit architectural decision in the opposite direction: external agents should talk to SAP's agents, not to SAP's APIs. Both companies are using the same protocol. They are making opposite bets about what the protocol is for. Most readers of either announcement are missing this. The split is the most consequential strategic divergence in enterprise software in 2026, and every other platform – Microsoft, ServiceNow, Workday, Oracle, Atlassian, Snowflake, Databricks – is currently inside the planning meeting where they have to pick a side.
On April 15, 2026, at TrailblazerDX in San Francisco, Salesforce announced what its co-founder Parker Harris framed in a single rhetorical question: Why should you ever log into Salesforce again? The accompanying product, Headless 360, is the most significant architectural shift in the company's twenty-seven-year history. Every capability in the Salesforce platform – data, workflows, business logic, compliance controls – is now exposed as a REST API, an MCP tool, or a CLI command. More than sixty new MCP tools shipped at launch, with over thirty preconfigured coding skills, and explicit support for external coding agents including Claude Code, Cursor, Codex, and Windsurf to operate the platform without a browser. Pricing is shifting from per-seat to consumption-based, an acknowledgment that when the agent is the user, the seat model is dead.
Six months earlier, in November 2025 at SAP TechEd, SAP made what looked at the time like a parallel announcement. The MCP Gateway, built on SAP's Integration Suite, would convert APIs and Flows into MCP tools usable by Joule, SAP's first-party agent. SAP announced an ABAP MCP server for the first half of 2026, a Storefront MCP server for SAP Commerce Cloud in Q2 2026, and a steady drumbeat of MCP-related shipping through TechEd and into early 2026. Press coverage treated it as MCP plumbing for ERP – the same story Salesforce would later tell, just on a six-month delay.
That reading is wrong. The two companies are using the same protocol to execute opposite strategies, and the strategies imply different futures for the entire enterprise stack.
The decisive document is buried in SAP's own architecture guidance, published quietly to the SAP Architecture Center in early 2026: "For external interoperability between vendors and third-party agents, SAP prioritizes A2A over direct MCP server exposure." That single sentence is the divergence. Salesforce wants its platform's MCP tools to be called directly by any agent in the market. SAP wants external agents to talk to Joule – its first-party agent – which will then decide whether and how to invoke the platform's MCP tools on the caller's behalf.
Same protocol. Opposite philosophy. The press has, almost without exception, missed the consequence.
The substrate bet and the broker bet
The two strategies need names, because right now the analyst coverage is using the same word ("MCP-enabled") to describe two architectures that have radically different implications for security, pricing, vendor power, and which third parties get to participate.
Call them the substrate bet and the broker bet .
The substrate bet, which Salesforce is running, says: the durable layer in agentic enterprise software is the platform that every agent calls into . Make the platform itself agent-callable from any surface, by any agent, with the platform's data and workflows and trust controls intact. The strategic prize is being underneath every workflow that an agent runs, whether the agent is built by Salesforce, by a customer, by Cursor, by Anthropic, by a startup nobody has heard of yet, or by a competitor. The platform wins by being the thing that gets called. Per-seat pricing dies because the seat dies; consumption pricing takes its place because the unit of value shifts to invocations.
The broker bet, which SAP is running, says: the durable layer is the first-party agent that holds the platform's privileges . External agents should not get direct platform access; they should request work from the platform's own agent, who decides – under the platform's governance model – whether and how to fulfill the request. The strategic prize is being the agent that other agents have to talk through . MCP exists internally so that Joule can do its work efficiently. Externally, SAP prefers A2A – the agent-to-agent protocol – which is structurally a delegation interface rather than a tool interface. The platform wins by being the thing that gets asked . Pricing remains anchored to the agent's seat or runtime, because Joule remains the addressable unit of value.
These are not implementation differences. They are different theories of where the trust boundary should sit in agentic enterprise software, who carries liability when an agent does something wrong, and which third parties get to participate in the economy that forms on top.
The substrate bet maximizes ecosystem participation and ecosystem dependence on the platform. The broker bet maximizes platform control over what agents can actually do.
Both bets are defensible. They cannot both be right.
What each company is implicitly claiming about the next five years
The two bets encode three sharply different claims about how the agentic enterprise will evolve.
Claim one: where does the customer's trust live?
Salesforce is betting that the customer's trust lives in the platform's data, workflow, and governance layer – not in any particular agent. The Einstein Trust Layer applies to every Agentforce session regardless of which agent invokes it. Data masking, field-level security, zero-data-retention agreements with LLM providers – these are platform-level enforcements that travel with the data regardless of who is calling. The customer trusts Salesforce because Salesforce sits at the gate of every invocation. The identity of the calling agent is, structurally, secondary.
SAP is betting that the customer's trust lives in the first-party agent that has been credentialed against the platform . Joule is the entity the customer has approved, scoped, and given permissions to. An external agent that wants to act in SAP needs to convince Joule, through A2A, to act on its behalf. The trust boundary is the boundary of Joule, not the boundary of the platform. Customers trust SAP because Joule is the only thing inside SAP's perimeter that talks to outside agents at all.
If you believe enterprise customers, in 2027 and beyond, will be more comfortable approving platforms than approving agents , Salesforce is right. If you believe customers will demand a single, named, accountable agent at the boundary of every platform – with everything else on the outside subject to that agent's judgment – SAP is right.
Claim two: where does liability sit when an agent does something wrong?
The Air Canada precedent already established that companies cannot disclaim responsibility for the chatbots on their websites. The next wave of cases – which will arrive within eighteen months, given the deployment rate – will ask a sharper question: when an external agent invokes an enterprise platform via MCP and produces a harmful outcome, who is liable? The agent's builder? The user who deployed the agent? The platform that exposed the tool?
The substrate bet implicitly accepts that the platform shares liability for every invocation. If Salesforce is the substrate every agent calls into, Salesforce's trust layer is the last enforcement point before action is taken. That is both a feature – Salesforce can market the trust layer as the reason a customer should consolidate onto Headless 360 – and a risk. Every poorly-built external agent that does something bad through a Salesforce MCP tool becomes a Salesforce problem.
The broker bet implicitly tries to push liability outward. Because external agents must go through Joule, SAP's first-party agent becomes the accountable actor, and the external agent becomes a requester rather than an operator . This is a defensible legal posture that the substrate model gives up by design. If you believe the next eighteen months will produce a wave of enforcement actions that punish platforms for the actions of third-party agents calling their tools, SAP's position is much more comfortable to be in.
Claim three: where does pricing power eventually accrue?
Salesforce has already moved to consumption pricing for Agentforce, and the broader signal is that per-seat pricing is over. The substrate bet ends in a model where Salesforce extracts value per invocation – regardless of which agent is doing the invoking. The agent ecosystem on top is encouraged to grow, because every agent in the ecosystem is a meter that runs Salesforce's billing system.
SAP has not made the equivalent pricing move publicly. The broker bet is consistent with continuing to price the agent itself, because the agent is the named entity holding the customer relationship. Joule's seats, Joule's runtime, Joule's per-workflow charges remain the addressable unit. If external agents want to participate, they participate as suppliers to Joule rather than as customers of SAP's APIs. SAP captures less of the per-invocation flow but maintains a clearer line of sight to what each customer's agent stack is doing.
If you believe agentic enterprise software ends in a world dominated by a small number of horizontal platforms with everything called via API, the substrate bet wins the pricing war. If you believe it ends in a world where each major enterprise vendor runs its own first-party agent and external agents are required to negotiate with those agents to get anything done, the broker bet wins.
Why the rest of the stack now has to pick
Salesforce and SAP did not converge on opposite doctrines by accident. Both companies have spent two-plus years internally arguing about which architecture to ship. Salesforce went substrate publicly two and a half years before the Headless 360 announcement, according to its own framing. SAP has been operationally committed to the broker model since at least mid-2025. Each company picked the doctrine that played to its existing strengths – Salesforce's ecosystem-friendly platform DNA versus SAP's deep integration into governed, regulated, complex enterprise workflows where Joule's role as accountable orchestrator is a feature, not a constraint.
The other major enterprise platforms have not yet had to commit publicly, but they will, in 2026 and into 2027. Each of them is, right now, inside the planning meeting where the substrate-versus-broker decision gets made. The signals so far:
Microsoft is hedging. Microsoft 365 Copilot looks substrate-shaped – exposed through endpoints, callable by external agents – but Microsoft's Agent 365 framing positions Microsoft's own agents as the privileged operators inside the enterprise. Watch the next eighteen months of Microsoft's Agent Builder and the maturity of its A2A versus MCP investments; Microsoft will have to land more clearly on one side.
ServiceNow is structurally pulled toward the broker model. Its core asset is the workflow graph, and exposing that graph to arbitrary external agents undermines the consulting-implementation economy that surrounds it. Expect ServiceNow's Now Assist to harden into a Joule-style first-party agent that brokers everything.
Workday has the same structural pull as ServiceNow but a weaker position to enforce it. HR and finance workflows are too cross-platform for a pure broker model to hold; Workday will probably end up in a hybrid, but the hybrid will lean broker.
Snowflake and Databricks are structurally pul

[truncated]

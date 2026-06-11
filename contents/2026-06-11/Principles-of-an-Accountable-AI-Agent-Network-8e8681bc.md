---
source: "https://www.tigera.io/blog/five-principles-of-an-accountable-ai-agent-network-how-to-evaluate-any-governance-platform/"
hn_url: "https://news.ycombinator.com/item?id=48489668"
title: "Principles of an Accountable AI Agent Network"
article_title: "Five Principles of an Accountable AI Agent Network: How to Evaluate Any Governance Platform"
author: "baroiall"
captured_at: "2026-06-11T13:28:21Z"
capture_tool: "hn-digest"
hn_id: 48489668
score: 2
comments: 0
posted_at: "2026-06-11T12:53:06Z"
tags:
  - hacker-news
  - translated
---

# Principles of an Accountable AI Agent Network

- HN: [48489668](https://news.ycombinator.com/item?id=48489668)
- Source: [www.tigera.io](https://www.tigera.io/blog/five-principles-of-an-accountable-ai-agent-network-how-to-evaluate-any-governance-platform/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T12:53:06Z

## Translation

タイトル: 責任ある AI エージェント ネットワークの原則
記事のタイトル: 責任ある AI エージェント ネットワークの 5 つの原則: ガバナンス プラットフォームを評価する方法
説明: このシリーズの最初の投稿では、AI エージェントのガバナンスが展開に追いついていないと主張しました。 2 番目のセクションでは、説明責任の 5 つの柱と何が必要かを示しました。 3 番目では、なぜネットワークが...

記事本文:
責任ある AI エージェント ネットワークの 5 つの原則: ガバナンス プラットフォームを評価する方法
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
Kubernetes 用の VM ネットワーキング NEW Kubernetes 上で VM ネットワーキングを移行、保護、運用するための実践者向けガイドを、1 台の VM の行程を通して説明します。さらに詳しく >
ガイド
Kubernetes
Kubernetes 101
ガイド
可観測性
可観測性
ネットワーキング
Kubernetes ネットワーキング
責任ある AI エージェント ネットワークの 5 つの原則: ガバナンス プラットフォームを評価する方法
このシリーズの最初の投稿では、AI エージェントのガバナンスが展開に追いついていないと主張しました。 2 番目のセクションでは、説明責任の 5 つの柱と何が必要かを示しました。 3 番目のセクションでは、ネットワーク ポリシー、API ゲートウェイ、MCP/A2A プロトコル、DIY セキュリティ パターン、およびロールベースのアクセス制御 (RBAC) がそれぞれ重大な責任のギャップを残す理由を説明しました。
5 つの柱は、AI エージェントに必要な説明責任を定義します。以下の原則は、ガバナンス プラットフォームがそれを実現する方法を定義します。これらは、AI エージェント ガバナンス ソリューションを構築するか、購入するか、オープンソース コンポーネントから組み立てるかにかかわらず、チームが AI エージェント ガバナンス ソリューションを評価する必要があるアーキテクチャ原則です。
ベンダーがこれら 5 つのいずれかに当てはまらないガバナンス プラットフォームを提案した場合は、その場を立ち去りましょう。
責任ある AI エージェント ネットワークの 5 つの原則とは何ですか?
Kubernetes ネットワーク ポリシーは、クラスターを保護するために不可欠です。これらは、どのポッドが他のどのポッドと通信できるかをネットワーク レベルで制限するものであり、セキュリティの一部として必ず含める必要があります。

オスチャー。
デフォルト拒否: ポリシーで明示的に許可されていない限り、エージェントは通信しません。
属性ベースのポリシー: ポリシーは、エージェント名ではなくエージェント属性を参照します。
ゼロトラスト ID: すべてのリクエストが認証され、すべての ID が検証されます。
設計による監査: すべての対話により、構造化された相関トレースが自動的に生成されます。
Kubernetes ネイティブ: このプラットフォームは、既存のインフラストラクチャを置き換えるのではなく、拡張します。
以下の各原則は、それが重要である理由と、合格するソリューションとはどのようなものかを説明します。
ガバナンス プラットフォームを評価する際には、5 つの原則をチェックリストとして使用してください。どれか一つでも失敗すれば、プラットフォームはセキュリティの舞台から遠ざかる原則が 1 つ欠けています。
ポリシーで明示的に許可されていない限り、エージェントは他のエージェントと通信しません。
これが責任を負うための唯一の安全な開始姿勢です。ガバナンス層がデフォルトで通信を許可し、明示的に禁止されているものだけをブロックする場合、予期しないすべてのインタラクションは管理されず、許可しなかったことに対して責任を負うことができません。
デフォルト拒否はモデルを反転します。ポリシーが明示的に許可するまでは何も許可されません。許可されているすべてのインタラクションは意図的で、追跡可能で監査可能です。新しいエージェントは、アクセスを許可するポリシーが作成されるまでデフォルトで隔離されます。これはまさに管理されたネットワークで必要な動作です。
デフォルトの拒否は制限的であるように見えますが、実際には自由です。セキュリティ チームは、起こり得るあらゆる悪質なインタラクションを予測する必要はありません。良いものだけを定義する必要があります。
原則 2: 属性ベースのポリシー
ポリシーは、エージェント名ではなく、エージェント属性を参照する必要があります。
ポリシー内でエージェント名をハードコーディングすると、エージェントを追加または名前変更するたびにガバナンス システムが壊れてしまいます。これは、何百ものファイアウォールを維持するのと同じです。

ネットワーク セグメントを使用する代わりに、IP ベースのルールを使用します。
属性ベースのポリシーは、機能、リスク レベル、チームの所有権、環境タグなどのプロパティを参照します。このポリシーでは、「Agent-Finance-v2 は Agent-Compliance-v3 を呼び出すことができます」の代わりに、「capability=financial-analyses を持つエージェントは、capability=compliance-query をもつエージェントを呼び出すことができる」と記載されています。
このアプローチには強力なスケーリング特性があります。新しいエージェントが一致する属性で登録されると、既存のポリシーが自動的に適用されます。ガバナンスは、エージェント ネットワークに対抗するのではなく、エージェント ネットワークとともに成長します。新しいエージェントを導入するチームは、許可リストを更新するためにチケットを提出する必要はありません。登録時にエージェントの属性が記述され、残りの部分はポリシー エンジンが処理します。
これは、エージェントが 10 人でも存続するセキュリティ モデルと 1,000 人でも存続するセキュリティ モデルを分ける原則です。
原則 3: ゼロトラスト ID
すべてのリクエストが認証されました。すべての身元が確認されました。デフォルトでは何も信頼しません。
エージェント ネットワークは、他の分散システムと同様に、スプーフィング、リプレイ攻撃、資格情報の盗難などのアイデンティティの脅威に対して脆弱です。しかし、エージェントはユーザーに代わって動作するという特有の課題を追加します。これは、ワークロード ID (これは実際にエージェントであると主張しているのか) とユーザー ID (このエージェントは誰に代わって動作しているのか) の両方を検証する必要があることを意味します。
ガバナンス プラットフォームは、暗号化ワークロード ID (エージェントが本物であることの証明) とトークンベースのユーザー ID (誰がアクションをトリガーしたかの確立) の二重認証をサポートする必要があります。両方の ID がポリシー評価と監査ログに使用できる必要があります。
有効期間の短い認証情報、自動ローテーション、および暗号化検証は、オプションのアドオンではなく、標準である必要があります。静的 API キーと有効期間の長いトークンは、エージェント ネットワークでは負担となります。妥協する

認証情報を取得すると、マシン速度での自動横移動が可能になります。
原則 4: 後付けではなく設計に従って監査する
すべての対話により、構造化された相関トレースが自動的に生成されます。
チームが事後的にログ記録を追加する必要がある場合、あなたはすでに説明責任を失っています。監査記録は、後から追加される別のシステムではなく、ガバナンス層の施行の副産物である必要があります。
ガバナンス層がポリシーを評価し、エージェントの対話を許可 (または拒否) すると、その評価が監査レコードになります。誰が誰に電話をかけたか、どのポリシーが評価されたか、決定内容は何だったのか、どの属性が一致したか、それがいつ行われたのかが記録されます。これらのレコードは次のようになります。
構造化された (フリーテキストのログではない)、
マルチホップ チェーン全体で相関付けられます (分散トレース ID を使用)。
エージェント別、ポリシー別、時間範囲別、結果別にクエリ可能。
実際的な意味: 監査証跡は、構成オプションではなく、ガバナンス プラットフォームの第一級の製品である必要があります。有効にしなければならない場合、誰かが忘れてしまいます。組み込まれていれば、常にそこにあります。
原則 5: Kubernetes ネイティブ
ガバナンス層は、既存のインフラストラクチャを置き換えるのではなく、既存のインフラストラクチャと連携して動作する必要があります。
企業は、Kubernetes、Helm チャート、GitOps パイプライン、RBAC、名前空間、可観測性スタックに多額の投資を行ってきました。別個のコントロール プレーン、独自のデプロイ モデル、または独自のオーケストレーション レイヤーを必要とする AI エージェント ガバナンス プラットフォームは、導入の抵抗と運用上のオーバーヘッドに直面します。
ガバナンス プラットフォームは、Helm 経由でデプロイ可能、CRD 経由で管理可能、監視可能 (Prometheus や OpenTelemetry など)、既存の ID インフラストラクチャ (OIDC プロバイダー、SPIFFE/SPIRE) と互換性がある必要があります。これは、偶然発生した外部システムではなく、Kubernetes プラットフォームの自然な拡張のように感じられるべきです。

その上を走ります。
これは開発者の経験だけの問題ではありません。それは運営の持続可能性に関するものです。ガバナンス プラットフォームに、プラットフォーム チームが持っていない特殊なスキルが必要な場合、それは実現要因ではなくボトルネックになります。
原則がどのように相互に強化されるか
これら 5 つの原則は独立したものではありません。それらは相互に強化し合います。
ガバナンス ソリューションを評価するときは、次の各原則をテストします。
ソリューションでデフォルトで通信を許可し、特定の対話のみをブロックする必要がある場合、原則 1 は満たされません。
ポリシーでエージェントに名前を付ける必要がある場合は、原則 2 が満たされません。
静的 API キーまたは有効期間の長いトークンに依存する場合、原則 3 に失敗します。
相関のある監査証跡が自動的に生成されない場合は、原則 4 に当てはまりません。
Kubernetes の外部に独自のコントロール プレーンが必要な場合は、原則 5 が満たされません。
適切なソリューションは 5 つすべてを実現します。なぜなら、説明責任には何よりも必要なものがあるからです。
デフォルト拒否とゼロトラストの違いは何ですか?
デフォルト拒否はポリシーの姿勢であり、明示的に許可されない限り通信は行われません。ゼロトラストはアイデンティティの姿勢であり、すべてのアイデンティティを毎回検証する必要があります。それらはお互いを強化しますが、交換可能ではありません。ゼロトラスト ID を備えているが、デフォルト許可ポリシーが適用されているプラ​​ットフォームは依然として管理されていません。
AI エージェントの責任において Kubernetes ネイティブが重要なのはなぜですか?
なぜなら、導入は、機能するガバナンス プラットフォームと棚上げされるガバナンス プラットフォームの違いだからです。プラットフォーム チームが新しいコントロール プレーンを学習したり、並行デプロイメント パイプラインを実行したり、独自のポリシー エンジンを操作したりする必要がある場合、ガバナンス層がボトルネックになり、公式パスが遅すぎるため、管理されていないエージェントが現れ始めます。
SPIFFE、OPA、OpenTelemetry を使用してこれを自分で構築できますか?
技術的にはそうです。実際のところ、あなたは

統合接着剤、マルチホップ チェーン間の監査相関関係、二重 ID 検証、属性ベースのポリシー モデリング、人間による監視面に 6 ～ 12 か月かかります。このシリーズの投稿 3 では、構築と購入のトレードオフについて説明しました。
これらの原則は Tigera Lynx に特有のものですか?
いいえ。これらは、商用、オープンソース、または自社開発のいずれであっても、責任あるエージェントのガバナンス プラットフォームのアーキテクチャ上の原則です。私たち自身も Lynx を評価するためにそれらを使用しています。検討するすべてのオプションを評価するためにそれらを使用することをお勧めします。
デフォルト拒否が唯一の安全な開始姿勢です。それ以外の場合は、管理されていない相互作用が残ります。
属性ベースのポリシーは、ガバナンスをエージェント 100 人を超えて拡張できる原則です。
ゼロトラスト ID は、ワークロード (これが適切なエージェントかどうか) とユーザー (誰の代理として機能しているか) の両方を検証する必要があります。
設計による監査とは、監査記録が個別のシステムではなく、執行の副産物であることを意味します。
Kubernetes ネイティブにより、プラットフォームがバイパスされるのではなく実際に採用されることが保証されます。
責任ある AI エージェントのための戦略ガイドを入手する
私たちは戦略ガイド「責任ある AI エージェント: 大規模な自律型 AI を管理する AI およびセキュリティ リーダーのための戦略ガイド」を作成しました。このガイドでは、一般的なガバナンス アプローチの並べて比較や各原則に対するスコアの付け方など、これらの原則を詳しく説明しています。
責任ある AI エージェントのための戦略ガイドを入手 →
ブログ投稿、ワークショップ、認定プログラム、新しいリリースなどに関する最新情報を入手してください。
技術ブログ
クラスター内のエージェントに関するフィールド ガイド
ディロン・バリー著
2026 年 6 月 10 日
クラスター内のすべてのサービスは名前でわかります。どのチームがそれぞれを所有しているか、何と通信しているか、どのように拡張するか、ログがどこに行くのかがわかります。エージェントは別の話です。それは違います...
技術ブログ
マルチL

AI エージェントを保護するための ayer ポリシー
ピーター・ケリー著
2026 年 6 月 3 日
Tigera では、現実世界の大規模なエンタープライズ エージェント向けに安全なランタイム環境を構築する製品を構築する仕事の一環として、私がよく考えているこのパズルの小さな部分の 1 つが...
会社ブログ
Calico の新機能: 2026 年春リリース
ベロニカ・スモリク著
2026 年 6 月 2 日
Kubernetes は 2014 年のデビュー以来、長い道のりを歩んできました。コンテナ化されたいくつかのマイクロサービスの実行から、AI エージェントからフルスケールの VM まで、あらゆるものにまたがる本番ワークロードのフリートをオーケストレーションするようになりました。
このフォームを送信することにより、Tigera がプライバシー ポリシーに従ってお客様の個人情報を処理することを認め、これに同意したものとみなされます。
可観測性とトラブルシューティング
著作権©
株式会社タイガーラの著作権はすべて留保されます。

## Original Extract

The first post in this series argued that AI agent governance hasn’t kept pace with deployment. The second laid out the five pillars of accountability, and what is required. The third walked through why network...

Five Principles of an Accountable AI Agent Network: How to Evaluate Any Governance Platform
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
Five Principles of an Accountable AI Agent Network: How to Evaluate Any Governance Platform
The first post in this series argued that AI agent governance hasn’t kept pace with deployment. The second laid out the five pillars of accountability, and what is required. The third walked through why network policies, API gateways, MCP/A2A protocols, DIY security patterns, and Role-based Access Control (RBAC) each leave critical accountability gaps.
The five pillars define what AI agent accountability requires. The principles below define how a governance platform should deliver it. These are the architectural principles your team should evaluate any AI agent governance solution against, whether you build it, buy it, or assemble it from open-source components.
If a vendor pitches you a governance platform that fails any of these five, walk away.
What are the five principles of an accountable AI agent network?
Kubernetes Network Policies are essential for securing any cluster. They restrict which pods can communicate with which other pods at the network level, and they should absolutely be part of your security posture.
Default-deny: No agent communicates unless a policy explicitly permits it.
Attribute-based policy: Policies reference agent attributes, not agent names.
Zero-trust identity: Every request authenticated, every identity verified.
Audit by design: Every interaction produces a structured, correlated trace automatically.
Kubernetes-native: The platform extends your existing infrastructure rather than replacing it.
Each principle below explains why it matters and what a passing solution looks like.
Use the five principles as a checklist when evaluating any governance platform. Fail any one, and the platform is one missing principle away from security theater.
No agent communicates with any other agent unless explicitly permitted by policy.
This is the only safe starting posture for accountability. If your governance layer defaults to allowing communication and only blocks what’s explicitly forbidden, every interaction you didn’t anticipate is ungoverned, and you can’t be accountable for what you didn’t authorize.
Default-deny flips the model: nothing is allowed until a policy explicitly permits it. Every allowed interaction is intentional, traceable, and auditable. New agents are isolated by default until policies are written to grant them access, which is exactly the behavior you want in a governed network.
Default-deny seems restrictive, but in practice it’s liberating. Your security team doesn’t have to anticipate every possible bad interaction. They only have to define the good ones.
Principle 2: Attribute-based policy
Policies should reference agent attributes, not agent names.
Hardcoding agent names in policies creates a governance system that breaks every time you add or rename an agent. It’s the equivalent of maintaining a firewall with hundreds of IP-based rules instead of using network segments.
Attribute-based policies reference properties like capabilities, risk levels, team ownership, and environment tags. Instead of “Agent-Finance-v2 can call Agent-Compliance-v3, ” the policy says “Agents with capability=financial-analysis can call agents with capability=compliance-query.”
This approach has a powerful scaling property: when a new agent registers with matching attributes, existing policies apply automatically. The governance grows with the agent network, not against it. A team deploying a new agent doesn’t need to file a ticket to update allow-lists, they describe the agent’s attributes at registration time, and the policy engine handles the rest.
This is the principle that separates a security model that survives at 10 agents from one that survives at 1,000.
Principle 3: Zero-trust identity
Every request authenticated. Every identity verified. Trust nothing by default.
Agent networks are susceptible to the same identity threats as any distributed system: spoofing, replay attacks, credential theft. But agents add an unique challenge: they operate on behalf of the users. This means both the workload identity (is this actually the agent it claims to be?) and the user identity (on whose behalf is this agent acting?) must be verified.
A governance platform should support dual authentication : cryptographic workload identity (proving the agent is genuine) and token-based user identity (establishing who triggered the action). Both identities should be available for policy evaluation and audit logging.
Short-lived credentials, automatic rotation, and cryptographic verification should be standard, not optional add-ons. Static API keys and long-lived tokens are liabilities in an agent network; compromised credentials can enable automated lateral movement at machine speed.
Principle 4: Audit by design, not by afterthought
Every interaction produces a structured, correlated trace automatically.
If your team has to add logging after the fact, you’ve already lost accountability. Audit records should be a byproduct of the governance layer’s enforcement , not a separate system bolted on later.
When the governance layer evaluates a policy and permits (or denies) an agent interaction, that evaluation is the audit record. It captures: who called whom, what policy was evaluated, what the decision was, what attributes matched, and when it happened. These records should be:
Structured (not free-text logs),
Correlated across multi-hop chains (using distributed trace IDs),
Queryable by agent, by policy, by time range, by outcome.
The practical implication: the audit trail should be a first-class product of the governance platform, not a configuration option. If you have to enable it, someone will forget. If it’s built in, it’s always there.
Principle 5: Kubernetes-native
The governance layer should work with your existing infrastructure, not replace it.
Enterprises have invested heavily in Kubernetes, Helm charts, GitOps pipelines, RBAC, namespaces, and observability stacks. An AI agent governance platform that requires a separate control plane, its own deployment model, or a proprietary orchestration layer will face adoption resistance and operational overhead.
The governance platform should be deployable via Helm, manageable via CRDs, observable (e.g. via Prometheus or OpenTelemetry), and compatible with existing identity infrastructure (OIDC providers, SPIFFE/SPIRE). It should feel like a natural extension of the Kubernetes platform, not a foreign system that happens to run on it.
This isn’t just about developer experience. It’s about operational sustainability . If the governance platform requires specialized skills your platform team doesn’t have, it will become a bottleneck instead of an enabler.
How the principles reinforce each other
These five principles aren’t independent. They reinforce each other:
When evaluating governance solutions, test each principle:
If a solution requires you to default to allowing communication and only block specific interactions, it fails Principle 1.
If it requires naming agents in policies, it fails Principle 2.
If it relies on static API keys or long-lived tokens, it fails Principle 3.
If it doesn’t produce correlated audit trails automatically, it fails Principle 4.
If it needs its own control plane outside Kubernetes, it fails Principle 5.
The right solution delivers all five. Because accountability requires nothing less.
What’s the difference between default-deny and zero-trust?
Default-deny is a policy posture — no communication unless explicitly permitted. Zero-trust is an identity posture — every identity must be verified, every time. They reinforce each other but aren’t interchangeable. A platform with zero-trust identity but default-allow policy is still ungoverned.
Why does Kubernetes-native matter for AI agent accountability?
Because adoption is the difference between a governance platform that works and one that gets shelved. If your platform team has to learn a new control plane, run a parallel deployment pipeline, or operate a proprietary policy engine, the governance layer becomes a bottleneck — and ungoverned agents start showing up because the official path is too slow.
Can I build this myself with SPIFFE, OPA, and OpenTelemetry?
Technically yes. Practically, you’ll spend 6–12 months on the integration glue , audit correlation across multi-hop chains, dual identity verification, attribute-based policy modeling, and the human oversight surface. We covered the build-vs-buy tradeoff in post 3 of this series .
Are these principles specific to Tigera Lynx?
No. These are architectural principles for any accountable agent governance platform — whether commercial, open source, or homegrown. We use them ourselves to evaluate Lynx, and we’d encourage you to use them to evaluate every option you consider.
Default-deny is the only safe starting posture. Anything else leaves ungoverned interactions.
Attribute-based policy is the principle that lets governance scale past 100 agents.
Zero-trust identity must verify both the workload (is this the right agent?) and the user (on whose behalf is it acting?).
Audit by design means audit records are a byproduct of enforcement, not a separate system.
Kubernetes-native ensures the platform actually gets adopted instead of bypassed.
Get the strategic guide for accountable AI agents
We wrote a strategic guide, Accountable AI Agents: A Strategic Guide for AI & Security Leaders Governing Autonomous AI at Scale , that walks through these principles in depth, including a side-by-side comparison of common governance approaches and how they score against each principle.
Get the strategic guide for accountable AI agents →
Get updates on blog posts, workshops, certification programs, new releases, and more!
Technical Blog
A field guide to the agents in your cluster
By Dillon Barry
on Jun 10, 2026
You know every service in your cluster by name. You know which team owns each one, what it talks to, how it scales, where its logs go. The agents are a different story. That’s not...
Technical Blog
Multi-Layer Policy for Securing AI Agents
By Peter Kelly
on Jun 3, 2026
As part of our work at Tigera building products that create secure runtime environments for enterprise agents at scale in the real world, one small part of this puzzle I think about a lot is...
Company Blog
What’s new in Calico: Spring 2026 Release
By Veronika Smolik
on Jun 2, 2026
Kubernetes has come a long way since its debut in 2014. It’s gone from running a couple of containerized microservices to orchestrating fleets of production workloads spanning everything from AI agents to full scale VMs...
By submitting this form, you acknowledge and agree that Tigera will process your personal information in accordance with the Privacy Policy .
Observability & Troubleshooting
Copyright ©
Tigera, Inc. All rights reserved.

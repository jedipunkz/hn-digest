---
source: "https://cloudnativenow.com/features/tigera-introduces-lynx-a-unified-control-plane-for-kubernetes%e2%80%91native-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48765696"
title: "Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents"
article_title: "Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents - Cloud Native Now"
author: "CrankyBear"
captured_at: "2026-07-02T19:09:58Z"
capture_tool: "hn-digest"
hn_id: 48765696
score: 2
comments: 0
posted_at: "2026-07-02T18:44:20Z"
tags:
  - hacker-news
  - translated
---

# Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents

- HN: [48765696](https://news.ycombinator.com/item?id=48765696)
- Source: [cloudnativenow.com](https://cloudnativenow.com/features/tigera-introduces-lynx-a-unified-control-plane-for-kubernetes%e2%80%91native-ai-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T18:44:20Z

## Translation

タイトル: Tigera、Kubernetes ネイティブ AI エージェント用の統合コントロール プレーンである Lynx を発表
記事タイトル: Tigera、Kubernetes‑Native AI エージェント向けの統合コントロール プレーンである Lynx を発表 - クラウドネイティブな今
説明: Tigera は、オープンソースの Calico ネットワーキングを支援することで最もよく知られており、従来のコンテナ セキュリティを超えた取り組みを進めています。

記事本文:
Tigera が、Kubernetes ネイティブ AI エージェント向けの統合コントロール プレーンである Lynx を発表 - クラウド ネイティブになりました
2026年7月2日(木)
ポッドキャスト
クラウドネイティブな今のポッドキャスト
Tigera、Kubernetes ネイティブ AI エージェント用の統合コントロール プレーンである Lynx を発表
最初の Kubernetes AI エージェント コントロール プレーンが登場しました。
Tigera は、オープンソースの Calico ネットワーキングと Kubernetes のセキュリティ スタックを支援していることで最もよく知られており、Lynx の立ち上げにより従来のコンテナ セキュリティを超えた取り組みを進めています。これは、Kubernetes ネイティブ AI エージェントを大規模に管理するために設計された統合コントロール プレーンです。
ClawManager 、 Agent Substrate 、 Agent Control Plane (ACP) など、Kubernetes AI エージェントを監視しようとするプログラムは他にもありますが、Lynx はそれらの中で最も成熟しているようです。正確に言うと、Lynx は Kubernetes ネイティブの水平方向にスケーラブルなコントロール プレーンです。これは、すべてのエージェント→ツール/エージェント→LLM 呼び出しのパスに存在します。 ID およびアクセス管理 (IAM) に関しては、Entra ID/Okta/SPIFFE に関連付けられます。 ID、ポスチャ、ポリシーを強制し、Kubernetes クラスター全体で eBPF/LSM を使用して異常を検出します。最後に、ポリシー管理には、オープンソースのポリシー言語および評価エンジン Cedar を使用します。
これが実際に意味することは、Lynx は次のように設計されているということです。ドラムロールしてください:
Kubernetes クラスター全体で実行されている AI エージェントを検出してインベントリを作成します。
各エージェントが何に、どのような条件でアクセスできるかを制御する、きめ細かいポリシーを添付します。
エージェントの動作、対話、データ フローをリアルタイムで可視化します。
マルチクラスターおよびマルチクラウド環境全体で一貫してセキュリティ制御を適用します。
Lynx は Kubernetes ネイティブ エージェント向けに構築されているため、Tigera は新しいサイロを強制するのではなく、既存のプラットフォームの実践に合わせた導入モデルを目指しています。 Lynx は介入することを目的としています

主要なクラウド プロバイダーやオンプレミス上の既存の Kubernetes クラスターと統合します。また、ご想像のとおり、Calico で保護された環境と連携して、一貫したネットワーク ポリシーとマイクロセグメンテーションを実現できるように構築されています。このプラットフォームは、開発者がすでに使用している一般的な AI ランタイム、フレームワーク、モデル ゲートウェイで動作するように設計されています。
このアプローチは、並列 AI インフラストラクチャ スタックを管理するのではなく、AI を既存の Kubernetes およびクラウド ネイティブ アプリケーション保護プラットフォーム (CNAPP) 戦略に統合したいと考えているプラ​​ットフォーム エンジニアリング チームやセキュリティ チームにアピールする可能性があります。
なぜそれが必要なのでしょうか？ Tigera CEO の Ratan Tipirneni 氏は、ブログ投稿で次のように説明しました。「AI エージェントは、セキュリティ スタックが構築されているという前提を打ち破りました。ほとんどの組織が実行しているエンタープライズ セキュリティ ツールは、決定論的なワークロード向けに設計されています。サービスは、今日も昨日とほぼ同じことを行います。… AI エージェントはそのようには機能しません。自律的で非決定的です。エージェントはユーザーに代わって動作し、必要なツール、LLM、または必要な他のエージェントに手を伸ばし、委任チェーンを運びます。同じエージェントが実行されるたびに、異なるパスを通過する可能性があります。」
それを考えると本当に恐ろしいです。これらの懸念に対処するために、Lynx はすべてのエージェントをカタログ化する中央レジストリを展開するとティピルネニ氏は書いています。シャドウ エージェントにはフラグが付けられて隔離され、エージェントのアクションは OpenTelemetry トレースを通じてエンドツーエンドで再構築できます。
また、Lynx はすべてのエージェントをベースラインに対して継続的に評価し、エージェントごとのサンドボックス化と、GDPR、HIPAA、SOC 2、および金融サービス要件にマッピングされた事前構築されたコンプライアンス パックを使用して、ドリフトや過剰な権限を出現した瞬間に明らかにします。
コントロール プレーンはまた、すべての

ID プロバイダー (Entra ID、Okta) と統合するか、SPIFFE/SPIRE 経由で、共有秘密を使用せずに、検証可能な暗号化 ID を生成します。代わりに、有効期間の長い API キーは、有効範囲が狭く、自動ローテーションされる短期間のトークンに取って代わられます。 JSON Web トークン (JWT) は、マルチエージェント ワークフローのホップごとに作成されます。このようにして、資格情報は渡されるのではなく、単一ホップに限定されます。
Lynx はすべてのトランザクションを承認し、ゲートウェイでポリシーを適用します。単一のデフォルト拒否ポリシーは、Cedar ポリシー言語を使用して LLM、MCP、およびエージェントのアクセスを管理し、コールが実行される前に評価されます。不正行為をしたエージェントは即座に隔離され、一か八かの通話を人間にルーティングできます。これもエージェント コードを変更する必要はありません。 Lynx は、プロンプト インジェクション防御、レート制限、ガードレール、予算、支出制限、カスタム Webhook、MCP 多重化、集約、セッション管理など、エージェントの保護と管理に必要なその他の制御も提供します。
最後に、Lynx は異常な動作を監視します。 BPF と LSM を使用して、カーネル内のすべてのシステムコール、ネットワークコール、およびファイルアクセスを監視します。そうすることで、アクションが技術的にポリシーを通過した場合でも、資格情報の盗難や横方向の移動を捕捉できます。
Tigera は、Lynx を Kubernetes ネイティブのコントロール プレーンとして構築することで、企業が AI エージェントに他のクラウドネイティブ アプリケーションと同じ運用パターン (宣言的構成、GitOps ワークフロー、コードとしてのポリシー) に従ってほしいと考えています。それは私にとって非常に安全な賭けのように思えます。
クリックして X で共有 (新しいウィンドウで開きます)
×
クリックして Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
クリックして LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
クリックして Reddit で共有 (新しいウィンドウで開きます)
レディット
← Kubernetes のコスト配分とクラウドの請求額が一致しない理由
RSS エラー: フィードを取得できませんでした

「https://securityboulevard.com/webinars/feed/」にあります。ステータスコードは「403」、コンテンツタイプは「text/html」です。文字セット=UTF-8`
Tigera、Kubernetes ネイティブ AI エージェント用の統合コントロール プレーンである Lynx を発表
Kubernetes のコスト配分とクラウドの請求額が一致しない理由
AI ネイティブ スタックはすでに存在します。私たちはこれをクラウドネイティブと呼んでいます
クラウド ネイティブが AI ネイティブ スタックになった経緯
RocksDB の再発見 – クラウドネイティブ アプリケーションの組み込みストレージ
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
メンテナンスはほとんど必要ありません
AI 支援ソフトウェア開発により、過去 12 か月間でテスト チームへのプレッシャーはどの程度増加しましたか? (1つ選択してください) *
圧力が大幅に上昇
あなたの組織がソフトウェア テストの自動化を拡大し、エージェント テストを導入することを妨げている最大の障害は何ですか? (1つ選択してください) *
自動テストは難しすぎるか維持コストがかかりすぎる
リリースサイクルが早すぎてテストを最新の状態に保つことができない
コーディングやスクリプトに関する専門知識が多すぎる
テストツールとフレームワークが断片化しすぎている
時間、予算、人員の不足
難易度

ビジネス価値またはROIを証明する
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
重要な戦略的優先事項
重要だがまだ初期段階
私たちの組織とは関係ありません

## Original Extract

Tigera, best known for backing the open-source Calico networking, is pushing beyond traditional container security.

Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents - Cloud Native Now
Thursday, July 2, 2026
Podcasts
Cloud Native Now Podcast
Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents
The first Kubernetes AI agent control plane is here.
Tigera , best known for backing the open-source Calico networking and security stack for Kubernetes, is pushing beyond traditional container security with the launch of Lynx . This is a unified control plane designed to manage Kubernetes‑native AI agents at scale.
While there are other programs that attempt to oversee Kubernetes AI agents, such as ClawManager , Agent Substrate , and Agent Control Plane (ACP) , Lynx appears to be the most mature of them. To be exact, Lynx is a Kubernetes‑native, horizontally scalable control plane. It sits in the path of every agent→tool/agent→LLM call. For identity and access management (IAM), it ties into Entra ID/Okta/SPIFFE. It enforces identity, posture and policy, and detects anomalies using eBPF/LSM across Kubernetes clusters. Finally, for policy management, it uses the o pen-source policy language and evaluation engine, Cedar .
What that means in practice is Lynx is designed to, drum-roll please:
Discover and inventory AI agents running across Kubernetes clusters.
Attach fine‑grained policies that govern what each agent can access and under what conditions.
Provide real‑time visibility into agent behavior, interactions, and data flows.
Enforce security controls consistently across multi‑cluster and multi‑cloud environments.
Because Lynx is built for Kubernetes‑native agents, Tigera is aiming for a deployment model that aligns with existing platform practices rather than forcing a new silo. Lynx is meant to integrate with your existing Kubernetes clusters on major cloud providers and on‑premises. It’s also built, as you’d expect, to work with Calico-secured environments for consistent network policy and microsegmentation. The platform is designed to operate with popular AI runtimes, frameworks, and model gateways that developers are already using.
This approach is likely to appeal to platform engineering and security teams that want to integrate AI into their existing Kubernetes and cloud-native application protection platform (CNAPP) strategies, rather than managing a parallel AI infrastructure stack.
Why would you need it? Tigera CEO, Ratan Tipirneni, explained in a blog post, “ AI agents broke the assumptions security stacks were built on. The enterprise security tooling most organizations run was designed for workloads that are deterministic. A service does roughly the same thing today that it did yesterday. … AI agents don’t work that way. They’re autonomous and non-deterministic. An agent acts on behalf of a user, reaches for whatever tool, LLM, or other agent it needs, carries a delegation chain, and reads untrusted input as it goes. The same agent can take a different path every time it runs.”
That’s really scary when you think about it. To address these concerns, Tipirneni wrote, Lynx deploys a central registry that catalogs every agent. Shadow agents are flagged and quarantined, and any agent’s actions can be reconstructed end-to-end through OpenTelemetry traces.
Lynx also continuously evaluates every agent against a baseline and surfaces drift and over-permissions the moment they appear, with per-agent sandboxing and pre-built compliance packs mapping to GDPR, HIPAA, SOC 2, and financial services requirements.
The control plane also gives every agent a verifiable cryptographic identity by integrating with your identity provider (Entra ID, Okta) or via SPIFFE/SPIRE, with no shared secrets. Instead, long-lived API keys give way to short-lived, tightly scoped, auto-rotated tokens. A JSON Web Token (JWT) is minted for each hop of a multi-agent workflow. This way, credentials are scoped to a single hop rather than passed around.
Lynx authorizes every transaction and enforces policy at the gateway. A single default-deny policy governs LLM, MCP, and agent access using the Cedar policy language, evaluated before any call executes. A misbehaving agent can be quarantined instantly, and a high-stakes call can be routed to a human—again, with no agent code changes. Lynx also provides the other controls needed to secure and manage agents, including prompt-injection defenses, rate limiting, guardrails, budgets, spend limits, custom webhooks, MCP multiplexing, aggregation and session management.
Finally, Lynx monitors abnormal behavior. It uses e BPF and LSM to watch every syscall, network call, and file access in the kernel. That way, it can catch credential theft and lateral movement even when an action technically passes policy.
By building Lynx as a Kubernetes‑native control plane, Tigera is betting that enterprises want AI agents to follow the same operational patterns as other cloud‑native applications: declarative configuration, GitOps workflows, and policy‑as‑code. That seems like a very safe bet to me.
Click to share on X (Opens in new window)
X
Click to share on Facebook (Opens in new window)
Facebook
Click to share on LinkedIn (Opens in new window)
LinkedIn
Click to share on Reddit (Opens in new window)
Reddit
← Why Kubernetes Cost Allocation and Cloud Bills Don’t Match
RSS Error: A feed could not be found at `https://securityboulevard.com/webinars/feed/`; the status code is `403` and content-type is `text/html; charset=UTF-8`
Tigera Introduces Lynx, a Unified Control Plane for Kubernetes‑Native AI Agents
Why Kubernetes Cost Allocation and Cloud Bills Don’t Match
The AI Native Stack Already Exists. We’ve Been Calling It Cloud Native
How Cloud Native Became the AI Native Stack
Rediscovering RocksDB – Embedded Storage in Cloud-Native Applications
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

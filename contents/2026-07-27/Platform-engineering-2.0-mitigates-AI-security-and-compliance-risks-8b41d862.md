---
source: "https://platformengineering.org/blog/how-platform-engineering-2-0-mitigates-ai-security-and-compliance-risks"
hn_url: "https://news.ycombinator.com/item?id=49074974"
title: "Platform engineering 2.0 mitigates AI security and compliance risks"
article_title: "How platform engineering 2.0 mitigates AI security and compliance risks"
author: "CrankyBear"
captured_at: "2026-07-27T21:03:23Z"
capture_tool: "hn-digest"
hn_id: 49074974
score: 2
comments: 0
posted_at: "2026-07-27T20:14:12Z"
tags:
  - hacker-news
  - translated
---

# Platform engineering 2.0 mitigates AI security and compliance risks

- HN: [49074974](https://news.ycombinator.com/item?id=49074974)
- Source: [platformengineering.org](https://platformengineering.org/blog/how-platform-engineering-2-0-mitigates-ai-security-and-compliance-risks)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T20:14:12Z

## Translation

タイトル: プラットフォーム エンジニアリング 2.0 は AI のセキュリティとコンプライアンスのリスクを軽減します
記事のタイトル: プラットフォーム エンジニアリング 2.0 が AI のセキュリティとコンプライアンスのリスクを軽減する方法
説明: プラットフォーム エンジニアリング 1.0 から 2.0 への移行が、AI のセキュリティとコンプライアンスの重要な課題にどのように対処するかを説明します。ネイティブ モデルのガバナンスとワークロードの分離が、AI エージェントと LLM を実稼働ワークフローに統合するためのスケーラブルで安全な基盤を確立する方法を学びます。

記事本文:
プラットフォーム エンジニアリング 2.0 が AI のセキュリティとコンプライアンスのリスクを軽減する方法
コミュニティ コミュニティ
概要 私たちを動かすストーリーと価値観
アンバサダーがプラットフォーム エンジニアリング アンバサダーになる
イベント 近くで開催されるイベントをチェックしてください
レポート 業界統計のナンバーワン情報源をチェックしてください
求人 次のプラットフォーム エンジニアリングの役割を見つける コミュニティに参加する
参加して貢献する
ベンダーの機会
プラットフォーム エンジニアリングの概要
プラットフォームエンジニアリング認定プラクティショナー
プラットフォームエンジニアリング認定アーキテクト
プラットフォーム エンジニア向けのエージェント インフラストラクチャ 新しい
エージェント開発プラットフォーム new
...その他にもたくさんあります。プラットフォーム エンジニアリング大学をチェックして認定を取得する
組織向け エンタープライズチーム向け トレーニングとアドバイス
ホーム
サービス
結果
リソース
パートナー向け
サービスプロバイダー
トレーニング再販業者
認定プロバイダー ディレクトリ
ブログ風景 認定を取得する コミュニティに参加する コミュニティに参加する 認定を取得する
Platform Weekly は、プラットフォーム エンジニアリングに関する最高のニュースレターです。今すぐ購読する ブログ
プラットフォーム エンジニアリング 2.0 が AI のセキュリティとコンプライアンスのリスクを軽減する方法
インフラ
データ
デベックス
AI
リーダーシップ
セキュリティ
DATA スポンサー付き プラットフォーム エンジニアリング 2.0 が AI のセキュリティとコンプライアンスのリスクを軽減する方法
プラットフォーム エンジニアリング チームは、過去 10 年間、Kubernetes クラスター、パイプライン、内部開発者プラットフォーム (IDP)、Web およびマイクロサービス ワークロードのプロセスの標準化に費やしてきました。現在、組織が大規模言語モデル (LLM) と AI エージェントを運用ワークフローに統合するにつれて、プラットフォーム エンジニアリング 1.0 はプラットフォーム エンジニアリング 2.0 に移行しつつあります。
このプラットフォーム エンジニアリング 2.0 への移行は、セキュリティ、特にこの新しいエージェントの世界での分離とコンプライアンスに明らかな影響を及ぼします。しかし、当然のことながら、何を費やさなければならないのかを懸念している人にとっては、効果的なプラットフォームです。

m エンジニアリング 2.0 戦略は、リセットではなく進化を表す必要があります。基本的なプラットフォーム エンジニアリング 1.0 基盤におけるこれらの変更は、Broadcom が最近導入したプラットフォーム エンジニアリング 2.0 ブループリントで利用できます。
ガバナンスとコンプライアンスに関しては、ドキュメントとアプリケーション コードに限定するだけでは十分ではありません。代わりに、構造化されたフレームワークは、実際の技術的制御を備えた一貫したガードレール、ポリシー、手順を実装することで AI リスクから保護する必要があります。セキュリティの責任は、展開後に追加されるのではなく、プラットフォーム レベルで強制される、プラットフォーム自体にまで移行する必要があります。
AI によりセキュリティ上の懸念が高まります。
IT における長年の真実は、人間は必然的に (意識的かどうかにかかわらず) 不正なコードをインフラストラクチャに導入または挿入するということです。 AI エージェントが機能している場合、AI が不正なコードをインフラストラクチャに導入することを許可すると、その公理はさらに増幅され、特に高度に分散された環境では、その影響は軽微なバグから大規模な停止まで多岐にわたります。まさにこれが、強力な分離メカニズムが不可欠な理由です。
2 番目の懸念は、AI 生成コードと大規模言語モデル (LLM) の急増に対処するために設計された広範な新しい規制に準拠することです。データ保管場所に関する新しい規制要件のため、継続的なコンプライアンスが要件となります。
これを解決するために、プラットフォーム エンジニアリング 2.0 は、コア機能としてコードとしてのポリシーの適用を自動化します。これにより、セキュリティ チームとコンプライアンス チームは、導入後に脆弱なオーダーメイドのアプリケーション アーキテクチャの監査を強制されるのではなく、ガードレールを決定するためのネイティブ プラットフォーム層が提供されます。
プラットフォーム エンジニアリング 2.0 による AI セキュリティの扱い方
プラットフォーム エンジニアリング 2.0 は、AI セキュリティを特別なケースとしてではなく、ガバナンスと分離のためのネイティブ構造を備えたファーストクラスのワークロード タイプとして扱います。プラットながら

orm Engineering 1.0 は「デフォルトで安全な」Kubernetes クラスターと GitOps パイプラインを提供し、2.0 では「デフォルトで管理される」モデル アクセスと「デフォルトで分離された」AI ワークロード レーンが追加されました。
セキュリティをプラットフォーム層とランタイム層に移行することは、従来の「シフトレフト」パイプライン チェックを補完するように設計されています。これは、継続的なランタイム セーフティ ネットとして機能し、静的コード分析やパイプライン テンプレートでは完全に見逃される、アクティブで予測不可能な脅威を捕捉します。このベースライン プラットフォーム アーキテクチャが必要なのは、フロンティア モデルが従来のボルトオン アプリケーション制御が対応できる速度よりもはるかに速くセキュリティ ギャップを拡大しているためです。
プラットフォーム サブストレートは、モデル ポイズニングと推論データの漏洩をネイティブに軽減する必要があります。これら 2 つの重要な攻撃ベクトルは、プラットフォーム エンジニアリング 1.0 アーキテクチャが決して予測するように設計されていませんでした。即時注入のような一般的なリスクでさえ、最終的にはこのより広範な分離の範疇に分類されますが、業界には現在、それらをシームレスに処理するためのネイティブ プラットフォーム アーキテクチャが不足しています。
大まかに言えば、2.0 のイテレーションは 2 つの基本的なセキュリティの柱を中心に展開します。まず、実行場所に関係なく、すべてのモデルとプロバイダーにわたってポリシー、データ処理ルール、ロギングを強制するコントロール プレーンを使用した、厳格なプラットフォーム レベルの AI モデル ガバナンスがあります。第 2 に、堅牢なワークロード分離がプラットフォーム ファブリックに組み込まれており、ある AI ワークロードが別の AI ワークロードのデータ、秘密、またはパフォーマンス エンベロープに影響しないよう構造的に強化されています。
どちらもオーダーメイドのパターンではなくプラットフォーム機能として公開する必要があるため、開発者はアプリケーションごとに再実装するのではなく、API やテンプレートを介してセルフサービス インフラストラクチャとして利用できます。
モデル ガバナンスは、集中管理プレーンとして機能します。モデル固有の政府のため

大規模な障害が発生すると、モデル層の上の制御層でポリシーを適用し、決定をログに記録し、すべてのモデルにわたって動作を監視できます。
プラットフォーム エンジニアリング 2.0 は、4 つのコア要素を通じてこれを運用します。まず、中央のモデル レジストリが、一貫した API を介して承認された機能と場所を追跡します。第 2 に、統一されたポリシーの適用により、データ処理と安全性のルールが普遍的にロックダウンされます。 3 番目に、集中監査により、インシデントを追跡し、プロンプトを記録するための単一画面が提供されます。最後に、標準化されたアクセス ワークフローにより、開発者のリクエストが自動的にリスク層にマッピングされ、手動のセキュリティ免除が置き換えられます。
この「AI としてのインフラストラクチャ」は審判として機能し、エンタープライズ プラットフォーム チームが必要なモデルを選択できるようにするため、開発者はジョブ チケットの送信や YAML ファイルの構成などを行うことなく、必要なものを選択できます。
ワークロードの分離により、エクスプロイトを封じ込める構造的な保証が確立されます。 AI は従来のマルチテナンシーを破壊するため、プラットフォームは共有 GPU を分離し、ベクター ストア内での埋め込みの混在を防ぎ、ダウンストリームのデータ漏洩を排除する必要があります。
この分離には、実験用サンドボックス、内部ワークロード、規制されたデータ層を分離するための専用ドメインが必要です。タグ付けされたリソース プールとワークロード グループは、負荷の高い推論ジョブによって遅延の影響を受けやすい対話型アプリケーションが低下しないように、ハードウェア バックエンドを分割する必要があります。
さらに、ゼロトラスト ID 制御を統合すると、モデルとツールが狭い範囲のサービス ID にバインドされ、プロンプト インジェクションが必然的に成功する場合のラテラル ムーブメントが防止されます。最後に、ネットワーク、ストレージ、およびランタイムのポリシーは、Confidential Computing を使用してクラスター、ノード、およびポッドのレイヤー全体で調整する必要があります。
プラットフォーム チームは、これらの分離戦略を再利用可能な意見にエンコードできます。

開発者ポータル内のネイテッドブロック。
AI セキュリティがプラットフォーム エンジニアリングをどのように変えるか
プラットフォーム エンジニアリング 2.0 には、開発者からデータ サイエンティスト、ビジネス関係者まで、あらゆる関係者をカバーする AI セキュリティを組み込む必要があります。これに応えて、PlatformEngineering.org と Broadcom は、複数年にわたる方向性のあるブループリントである Platform Engineering 2.0 を導入しました。これは、連動する 5 つの建築上の柱を中心に構成されたプラットフォーム基板です。
AI ネイティブ プラットフォームは、GPU/TPU プロビジョニング、MCP ゲートウェイ、および AI エージェントをプラットフォーム シチズンとして実行しながら、IDP をエージェント開発プラットフォーム (ADP) に移行します。
マルチペルソナ エクスペリエンスは、データ サイエンティスト、FinOps およびセキュリティ リーダー、AI エージェント向けにカスタマイズされたエクスペリエンス レイヤーを備えた、より広範な企業 (開発者を超えた) にサービスを提供します。
組み込み FinOps は、コスト追跡を遡及レポートからプロビジョニング時間に移行し、トークン エコノミクス、リアルタイム アトリビューション、および導入前のコスト ゲートを管理します。
セキュリティはプラットフォームとランタイム基板にまで移行し、シャドウ AI のスプロールを緩和し、迅速なセキュリティをサポートし、モデル レジストリ ガバナンスを可能にします。
コンポーザブル アーキテクチャに基づいて構築された基盤は、厳格な「構築 vs 購入」の考え方を拒否し、ワークロードの要求に応じてすぐに交換または再構築できるモジュール式の API ファーストのビルディング ブロックを支持します。
中立的でスケーラブルな AI の信頼境界。
厳密なモデル ガバナンスと強力なワークロード分離を総合すると、AI に対する中立的なプラットフォーム レベルの信頼境界の作成に役立ち、プラットフォーム エンジニアリング 2.0 が重要な役割を果たします。各アプリケーションがモデルやツールの周囲に独自のあいまいな境界線を描くのではなく、プラットフォームは一貫した線を定義します。
境界内では、モデルは管理されたプレーンを通じて呼び出され、ワー​​クロードは定義された分離ドメインで実行され、共有されたドメインで実行される必要があります。

コントロールはデータとツールへのアクセスを仲介します。境界の外では、シャドウ AI とアドホック統合は明らかにポリシーの対象外であり、発見されやすく、認可されたパターンに置き換えるのが簡単です。
組織はすでに AI によるスプロール化による負担を感じており、プラットフォーム エンジニアリング 2.0 はそれを緩和することを目的としています。プラットフォーム エンジニアリング 2.0 は、四半期ごとにセキュリティ アーキテクチャを再描画することなく、AI を実験し、ネイティブ ワークロードを拡張し、新しいモデルの機能を吸収する体系的な方法を提供します。プラットフォーム チームが AI の信頼境界を、ファーストクラスのプリミティブとしてモデル ガバナンスとワークロード分離を備えたプラットフォーム管理アーキテクチャに移行するのが早ければ早いほど、より早く AI ワークロードを適切に保護できるようになります。
👉 プラットフォーム、開発者、リーダーシップを 1 つの明確な方向に調整する 👉 開発者の導入とプラットフォームの ROI を向上させる 👉 次に何に重点を置くべきかを明確にする トレーニングとアドバイスを調べる この投稿を共有する
関連記事
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
アンバサダー パイロットから本番まで: 規制された業界で AI コーディング エージェントを拡張するためのプラットフォーム チームのハンドブック Eric Paulsen Field CTO - International @ Coder • •
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
アンバサダー プラットフォームは埋もれた知識の上に座っています。エージェントはあなたにそれを掘り起こすよう強制しています マロリー・ヘイグ プラットフォーム エンジニアリング コンサルティング @ プラットフォーム教育および権利擁護責任者 • •
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
インフラ
デベックス
AI
データ
リーダーシップ
セキュリティ
アンバサダー AI ネイティブの未来はオープンソースである Sam Barlien プラットフォーム エンジニアリング部門のエコシステム責任者 • • すべての記事
Slack に参加して会話に参加して、プラットフォーム エンジニアリング コミュニティのトレンドと機会を常に把握してください。
ユーチューブ
リンクトイン
プラット

毎週のフォーム
ツイッター
久部の家
インテリジェンスを織る
毎週のプラットフォームを購読する
プラットフォーム エンジニアリングの詳細と DevOps トレンドが毎週、あなたの受信箱に頻繁に配信されます。

## Original Extract

Discover how the shift from Platform Engineering 1.0 to 2.0 addresses critical AI security and compliance challenges. Learn how native model governance and workload isolation establish a scalable, secure foundation for integrating AI agents and LLMs into production workflows.

How platform engineering 2.0 mitigates AI security and compliance risks
Community Community
Overview The story and values that drive us
Ambassadors Become a Platform Engineering Ambassador
Events Check out upcoming events near you
Reports Check out the #1 source of industry stats
Jobs Find your next platform engineering role Join Community
Join and contribute
Vendor opportunities
Certifications Introduction to Platform Engineering
Platform Engineering Certified Practitioner
Platform Engineering Certified Architect
Agent infrastructure for Platform Engineers new
Agentic Development Platforms new
...and many more. Check out Platform Engineering University Get Certified
For organizations FOR ENTERPRISE TEAMS Training & advisory
Home
Services
Results
Resources
FOR Partners
Service Provider
Training Reseller
Certified Provider Directory
Blog Landscape Get certified Join community Join community Get certified
Platform Weekly, the best newsletter in Platform Engineering. Subscribe now Blog
How platform engineering 2.0 mitigates AI security and compliance risks
Infra
DATA
DEVEX
AI
Leadership
SECURITY
DATA Sponsored How platform engineering 2.0 mitigates AI security and compliance risks
Platform engineering teams have spent the last decade standardizing Kubernetes clusters, pipelines, internal developer platforms (IDPs), and processes for web and microservice workloads. Now, as organizations integrate large language models (LLMs) and AI agents into production workflows, platform engineering 1.0 is shifting into platform engineering 2.0.
This shift to platform engineering 2.0 has clear implications for security, specifically for isolation and compliance in this new agentic world. But for those understandably concerned about what they must spend, an effective platform engineering 2.0 strategy should represent an evolution, not a reset. These changes in the foundational platform engineering 1.0 substrate are available in Broadcom’s recently introduced Platform Engineering 2.0 blueprint.
For governance and compliance, confinement to documents and application code isn't enough. Instead, structured frameworks must protect against AI risks by implementing consistent guardrails, policies, and procedures with real technical controls. Security responsibilities must move down into the platform itself, enforced at the platform level, not bolted on after deployment.
AI escalates security worries.
A long-held truth in IT is that a human will inevitably (knowingly or not) introduce or inject bad code into the infrastructure. With AI agents at work, that axiom is amplified if AI is allowed to introduce bad code into the infrastructure, and the effects can range from minor bugs to cataclysmic outages, especially in highly distributed environments. This is precisely why strong isolation mechanisms are essential.
A second concern is complying with expansive new regulations designed to address the explosion in AI-generated code and large language models (LLMs). Because of new regulatory requirements for data residency, continuous compliance is a requirement.
To solve this, platform engineering 2.0 automates policy-as-code enforcement as a core capability. This gives security and compliance teams a native platform layer to dictate guardrails, rather than forcing them to audit fragile, bespoke application architectures after deployment.
How platform engineering 2.0 treats AI security
Platform engineering 2.0 treats AI security not as a special case but as a first‑class workload type, with native constructs for governance and isolation. While platform engineering 1.0 offered “secure by default” Kubernetes clusters and GitOps pipelines, 2.0 adds “governed by default” model access and “isolated by default” AI workload lanes.
Shifting security down into the platform and runtime layers is designed to complement legacy "shift-left" pipeline checks. It acts as a continuous runtime safety net, catching the active, unpredictable threats that static code analysis and pipeline templates completely miss. This baseline platform architecture is required because frontier models are widening the security gap much faster than traditional, bolt-on application controls can respond to.
The platform substrate must natively mitigate model poisoning and inference data leaks - two critical attack vectors that platform engineering 1.0 architectures were never designed to anticipate. Even common risks like prompt injections ultimately fall under this broader umbrella of isolation, and the industry currently lacks native platform architectures to handle them seamlessly.
At a high level, the 2.0 iteration revolves around two foundational security pillars: First, there's strict, platform‑level AI model governance using a control plane that enforces policy, data‑handling rules, and logging across all models and providers, regardless of where they run. Second, robust workload isolation is baked into the platform fabric with structural enforcement that one AI workload cannot bleed into another’s data, secrets, or performance envelope.
Both must be exposed as platform capabilities rather than bespoke patterns, so developers can consume them as self-service infrastructure through APIs and templates rather than reimplementing them per application.
Model governance serves as a centralized control plane. Because model-specific governance breaks at scale, a control layer above the model tier can enforce policy, log decisions, and monitor behavior across all models.
Platform engineering 2.0 operationalizes this through four core elements. First, a central model registry tracks approved capabilities and locations via a consistent API. Second, unified policy enforcement locks down data-handling and safety rules universally. Third, centralized auditing provides a single pane of glass for tracing incidents and logging prompts. Finally, standardized access workflows automatically map developer requests to risk tiers, replacing manual security exemptions.
This "infrastructure as AI" acts as a referee, allowing the enterprise platform team to select the models they want, so the developer can pick and choose what they need without having to submit a job ticket, configure YAML files, etc.
Workload isolation establishes a structural guarantee to contain exploits. Because AI disrupts traditional multi-tenancy, platforms are forced to isolate shared GPUs and prevent embedding commingling within vector stores to eliminate downstream data leaks.
This isolation requires dedicated domains to separate experimental sandboxes, internal workloads, and regulated data tiers. Tagged resource pools and workload groups must partition hardware backends so that heavy inference jobs never degrade latency-sensitive interactive applications.
Furthermore, integrating zero-trust identity controls binds models and tools to narrowly scoped service identities, preventing lateral movement when prompt injections inevitably succeed. Finally, network, storage, and runtime policies must align across cluster, node, and pod layers using confidential computing.
Platform teams can then encode these isolation strategies into reusable, opinionated blocks inside developer portals.
How AI security is changing platform engineering
Platform Engineering 2.0 must incorporate AI security that covers the full range of stakeholders, from developers to data scientists to business stakeholders. In response, PlatformEngineering.org and Broadcom have introduced a multi-year directional blueprint, Platform Engineering 2.0. It's a platform substrate organized around five interlocking architectural pillars:
An AI-native platform transitions the IDP to an agentic development platform (ADP) while executing GPU/TPU provisioning, MCP gateways, and AI agents as platform citizens.
A multi-persona experience serves the wider enterprise (beyond developers) with tailored experience layers for data scientists, FinOps and security leaders, and AI agents.
Embedded FinOps shifts cost tracking from retrospective reporting to provisioning time to govern token economics, real-time attribution, and predeployment cost gates.
Security shifts down into the platform and runtime substrate to mitigate shadow AI sprawl, support prompt security, and enable model registry governance.
A foundation built on a composable architecture rejects rigid "build vs. buy" thinking in favor of modular, API‑first building blocks that can be swapped or repaved quickly as workloads demand.
A neutral, scalable AI trust boundary.
Taken together, strict model governance and strong workload isolation will help create a neutral, platform‑level trust boundary for AI, and platform engineering 2.0 will play a key role. Instead of each application drawing its own fuzzy perimeter around models and tools, the platform defines a consistent line.
Inside the boundary, models should be invoked through a governed plane, workloads run in defined isolation domains, and shared controls mediate access to data and tools. Outside the boundary, shadow AI and ad‑hoc integrations are clearly out of policy, easier to discover, and easier to replace with sanctioned patterns.
Organizations are already feeling the strain of AI‑driven sprawl, and platform engineering 2.0 aims to ease that. Platform engineering 2.0 provides a systemic way to experiment with AI, scale native workloads, and absorb new model capabilities without redrawing the security architecture every quarter. The sooner platform teams move the AI trust boundary into platform‑managed architectures with model governance and workload isolation as first‑class primitives, the sooner they can properly secure AI workloads.
👉 Align platform, dev, and leadership on one clear direction 👉 Improve developer adoption and platform ROI 👉 Get clarity on what to focus on next Explore training & advisory Share this post
Related articles
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Ambassador From pilot to production: The platform team's playbook for scaling AI coding agents in regulated industries Eric Paulsen Field CTO - International @ Coder • •
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Ambassador Platforms are sitting on buried knowledge. Your agents are forcing you to dig it up Mallory Haigh Head of Platform Education and Advocacy @ Platform Engineering Consulting • •
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Infra
DEVEX
AI
DATA
Leadership
SECURITY
Ambassador The future of AI-native Is open source Sam Barlien Head of Ecosystem @ Platform Engineering • • All articles
Join our Slack Join the conversation to stay on top of trends and opportunities in the platform engineering community.
Youtube
LinkedIn
Platform Weekly
Twitter
House of Kube
Weave Intelligence
Subscribe to Platform Weekly
Platform engineering deep dives and DevOps trends, delivered to your inbox crunchy, every week.

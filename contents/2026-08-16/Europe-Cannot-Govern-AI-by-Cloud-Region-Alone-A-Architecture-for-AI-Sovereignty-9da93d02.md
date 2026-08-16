---
source: "https://zenodo.org/records/21967218"
hn_url: "https://news.ycombinator.com/item?id=49320995"
title: "Europe Cannot Govern AI by Cloud Region Alone: A Architecture for AI Sovereignty"
article_title: "Europe Cannot Govern AI by Cloud Region Alone: A Technical Architecture for AI Sovereignty | Zenodo"
author: "sangamdas1982"
captured_at: "2026-08-16T16:13:31Z"
capture_tool: "hn-digest"
hn_id: 49320995
score: 2
comments: 0
posted_at: "2026-08-16T15:31:51Z"
tags:
  - hacker-news
  - translated
---

# Europe Cannot Govern AI by Cloud Region Alone: A Architecture for AI Sovereignty

- HN: [49320995](https://news.ycombinator.com/item?id=49320995)
- Source: [zenodo.org](https://zenodo.org/records/21967218)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T15:31:51Z

## Translation

タイトル: ヨーロッパはクラウド地域だけで AI を統治できない: AI 主権のためのアーキテクチャ
記事のタイトル: ヨーロッパはクラウド地域だけで AI を統治できない: AI 主権のための技術アーキテクチャ |ゼノド
説明: 問題領域 ヨーロッパの AI の主権は、インフラストラクチャの配置場所、つまりモデルがホストされる場所、データが保存される場所、クラウド領域が選択される場所、およびサービス プロバイダーに適用される地域ルールによってアプローチされることがよくあります。これらのコントロールは重要ですが、要望に完全には答えられません。
[切り捨てられた]

記事本文:
ヨーロッパはクラウド地域だけで AI を統治できない: AI 主権のための技術アーキテクチャ |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
ヨーロッパはクラウド地域だけで AI を統治できない: AI 主権のための技術アーキテクチャ
欧州の AI の主権は、インフラストラクチャの配置場所、つまりモデルがホストされる場所、データが保存される場所、クラウド領域の選択、サービス プロバイダーに適用される地域規則によってアプローチされることがよくあります。これらのコントロールは重要ですが、実行時のより難しい質問に完全に答えることはできません。
特定の AI 計算が無許可のアルゴリズム ロジックの下で開始されるのを防ぐことができますか? また、特定の AI によって生成された結果が無許可の管轄区域で有効になるのを防ぐことができますか?
クラウド リージョンの選択は、主に展開の決定によって行われます。 AI ワークロードは、承認された欧州地域で実行される可能性がありますが、その出力はその後、外部 API、エージェント ツール、フェデレーション サービス、通信ネットワーク、衛星インフラストラクチャ、ダウンストリーム プラットフォーム、または意図された管轄境界外の受信者環境を介して送信されます。 IP ベースの地理位置情報は、AI によって生成された結果が有効になる実際の環境からルーティングをプロキシ、トンネリング、マスク、抽象化、または分離できるため、間接的な証拠のみを提供します。したがって、基礎となる開示では、インフラストラクチャの配置と呼び出しごとの管轄権の適用との間のギャップが特定されます。
AI システムが受動的な情報生成装置から、ツールの呼び出し、外部サービスとの通信、ワークフローの制御、法的または商業的に重大な出力の生成が可能なシステムに進化するにつれて、この問題はさらに重大になります。

、複数のインフラストラクチャと管轄ドメインにまたがって運用されています。このようなシステムでは、AI サービスが名目上ヨーロッパで展開されたことがわかっても、すべての計算で承認されたアルゴリズム ロジックが使用されたことや、結果として得られるすべての出力が承認された管轄環境内でのみ有効になったことは証明されません。
したがって、技術的な問題は、単に AI がホストされる場所ではなく、どこでどのような権限のもとで計算と外部効果の発生が許可されるかということです。
提案されたアーキテクチャは、調整された 2 つの実行時間境界を通じてこの問題に対処します。
最初の境界では、外部コンピューティング プレーンが推論を開始することを不可逆的に許可される直前に、保護されたドメインが、提案されたモデル、チェックポイント、オーケストレーション構成、または実行ロジックのアルゴリズム ロジック フィンガープリント (ALF) を導出または取得し、承認されたまたは事前バインドされた ALF 条件と比較します。計算開始機能は、必要な条件が満たされた場合にのみ発行されます。それ以外の場合、認証はフェールクローズ方式で保留されます。
2 番目の境界では、候補出力が外部で有効になる直前に、アーキテクチャは、意図された管轄コンテキストとともに実行時のガバナンス状態を評価します。 IP 地理位置情報のみに依存するのではなく、証明書、リモート構成証明、ソブリンクラウド識別子、通信ドメイン識別子、衛星ドメイン資格情報、受信者の公開鍵インフラストラクチャ、およびその他のインフラストラクチャにリンクされた属性を含む、機械で検証可能なインフラストラクチャ証拠を使用できます。
これにより、ヨーロッパの AI 主権の異なるモデルが作成されます。
主権は単にインフラの位置によって宣言されるわけではありません。これは、計算が開始される境界で技術的に適用され、AI によって生成された

nシーケンスが有効になります。
欧州の AI 主権の追求は、クラウド領域の選択、データ常駐ポリシー、および IP ベースの地理位置情報では完全には答えられない技術的な問題を引き起こしています。それは、運用上有効になった瞬間に、特定の AI 計算および特定の AI によって生成された結果に対して管轄権をどのように適用できるのでしょうか?
この論文では、AI 主権のための実行時の二重境界アーキテクチャを紹介します。このアーキテクチャは、インフラストラクチャの配置を機械による強制可能な権限から分離し、2 つの調整された制御ポイントを導入します。1 つは外部コンピューティング プレーンが推論を開始することを不可逆的に許可される直前の事前計算認可境界、もう 1 つは候補 AI 出力が外部で有効になる直前の出力ファイナリティ境界です。
計算前の境界では、信頼された実行環境 (TEE)、ハードウェア セキュリティ モジュール (HSM)、または機能的に同等の保護された環境などの保護された承認済みドメインが、提案されたモデル、チェックポイント、オーケストレーション構成、または実行ロジックを表すアルゴリズム ロジック フィンガープリント (ALF) を導出または検証します。計算開始権限は、ALF 候補が必要な権限条件を満たした場合にのみ解放されます。 ALF が存在しない、不一致、失効、期限切れ、またはその他の不正な ALF があると、認証はフェールクローズ方式で保留されます。一括推論は外部 GPU、アクセラレータ、またはクラウド インフラストラクチャ上に残る場合があります。保護されたドメインは、計算の開始に必要なセキュリティ上重要な権限を制御します。
承認された推論中に、ガバナンス関連の実行信号がキャプチャされ、制限された機械可読なランタイム動作ガバナンス記述子として表現されます。アーキテクチャには必要ありません

e モデルの完全な潜在推論または意味論的状態の再構築。代わりに、特定の推論実行に関連付けられた監視可能な実行パス情報を使用できます。
外部リリースの前に、実行時のガバナンス状態を、意図された宛先またはリリース環境を表す管轄コンテキストに暗号的にバインドできます。次に、管轄区域は、暗号証明書、署名された地域または事業者ドメインの主張、リモート認証結果、ソブリンクラウド識別子、通信ドメイン識別子、衛星ドメイン資格情報、受信者の公開鍵階層、または同等のインフラストラクチャにリンクされた資格情報を含む、非 IP の機械検証可能なインフラストラクチャ証拠を使用して評価されます。
したがって、リリースは、承認されたアルゴリズム ロジック、承認された実行時の動作、承認された目的、検証された管轄区域、拘束力のある完全性、最新性、有効期限および失効ステータス、および必要に応じてソースおよび宛先の管轄当局の両方からの承認など、複数の述語の同時満足に依存する可能性があります。したがって、許容可能な AI 出力は、未承認の管轄区域に向けられた場合でもブロックされる可能性がありますが、承認された管轄区域は不承認のランタイムまたはアルゴリズムの状態をオーバーライドすることはできません。
このアーキテクチャは、欧州が AI インフラストラクチャを分離したり、信頼できるハードウェア内で完全な AI ワークロードを実行することを要求したりすることを提案していません。代わりに、高パフォーマンスの計算を承認クリティカルな強制から分離します。既存の GPU およびクラウド インフラストラクチャは推論を実行できますが、保護されたインフラストラクチャは、保護された計算を開始し、その結果の出力を外部に出すために必要な機能を制御します。したがって、この開示では、管轄権を単なる呼び出しではなく、呼び出しごとの技術的述語にすることを目指しています。

これは展開時の想定です。
欧州のデジタル主権の場合、結果として得られる命題は単純です。
クラウド リージョンは、インフラストラクチャが展開される場所を決定できます。それ自体では、特定の AI 計算または AI によって生成された結果が、特定の管轄区域で有効になる権限を持っているかどうかを決定するものではありません。
実行ファイナリティ ガバナンスにより、その決定は、計算が開始され、出力が外部に有効になる、機械が強制可能な境界に移動します。
特殊なハードウェア要件と高スループットの拡張性
アーキテクチャには NVIDIA H100/H200 Confidential Computing、AMD SEV-SNP、またはその他の特定のハードウェアが必要ですか?
いいえ。このアーキテクチャは、NVIDIA H100/H200 Confidential Computing、AMD SEV-SNP、または単一のプロセッサ、アクセラレータ、TEE ベンダー、またはクラウド プラットフォームに依存しません。
アーキテクチャ上の要件はさらに狭く、セキュリティ クリティカルな状態を保護し、関連する実行境界で必要な機能のリリースを制御できる、保護された承認ドメインが存在する必要があります。本開示は、TEE、HSM、セキュア エンクレーブ、セキュア エレメント、機密コンピューティング モジュール、ハードウェア ルート制御モジュール、または機能的に同等の暗号的に分離された強制ドメインを明示的に許可します。また、保護されたドメインが AI 推論ワークロード全体を実行する必要はないことも明示的に述べています。
したがって、NVIDIA Confidential Computing と AMD SEV-SNP は、アーキテクチャ上の前提条件ではなく、実装可能な基板として理解される必要があります。
たとえば、Confidential Computing をサポートする NVIDIA H100 以降の GPU は、GPU ハードウェアとソフトウェアの整合性に関する GPU 構成証明と暗号証拠を提供できます。 NVIDIA の現在の認証アーキテクチャは、ローカルおよびリモートの GPU 認証をサポートしています。

を実行し、複数の GPU にわたって証拠を収集できます。 AMD SEV-SNP は、別の実装基板を提供します。これは、信頼できないハイパーバイザーから認可コントローラーを保護できるハードウェア ルートの分離と構成証明を備えた機密 VM です。
どちらのテクノロジーも、それ自体では、提案されている実行ファイナリティ アーキテクチャを実装しません。
機密コンピューティングの答え:
このワークロードまたは実行環境は、期待される保護された状態で実行されていますか?
実行ファイナリティ ガバナンスでは次のことが求められます。
その保護された状態を考慮すると、この特定の計算または出力は、次の結果的な境界を越えるために必要な権限を持っていますか?
したがって、認証は、それ自体が最終的な認可決定ではなく、実行のファイナリティに対する入力述語になる可能性があります。
デプロイメントでは、たとえば次のものを使用できます。
CPU 機密 VM / TEE
→ ホールドポリシー、トラストアンカー、ALFリポジトリ、認可キー
GPU またはアクセラレータ
→ 一括推論を実行する
GPU/VM 認証
→ 関連する実行環境の状態を証明する
保護された認可コントローラー
→ ALF + 証明書 + ポリシーを検証する
機能の解放
→ 計算を許可する
実行時の証拠 + 宛先認証情報 + 管轄コンテキスト + 鮮度/失効状態
→ ファイナリティコントローラー
→ 出力ファイナリティ機能を解放または保留します。
これは、本開示における高性能計算と認可が重要な操作との明確な分離と一致している。
現在の機密コンピューティング システムは、認証の成功を条件とした機密またはリソースの解放が技術的に実用的であることも実証しています。たとえば、NVIDIA の Confidential Containers アーキテクチャでは、CPU/GPU TEE 状態の証明と、その後のポリシー制御によるシークレットまたはリソースのリリースが記述されています。
機密性の高い GPU を使用しない展開
自信を持って

したがって、tial GPU は必須ではありません。
1 つの実装では、従来の GPU が外部で推論を実行しながら、認可コントローラーを AMD SEV-SNP 機密 VM 、HSM、または別の保護された CPU 側環境に配置できます。
推論の前に、保護されたコントローラーは次のことを行うことができます。
呼び出し識別子と保護されたポリシーコンテキストを受け取ります。
提案されたモデル/チェックポイント/オーケストレーション ID を取得します。
候補ALFを導出または検証する。
利用可能な場合は、関連するコンピューティング プレーンの構成証明またはサービス資格情報を確認します。
有効期限と失効状態を評価します。
新しい呼び出し識別子または nonce を作成します。そして
制限付きスケジューラ トークン、モデル アクセス資格情報、復号化キー、API 認可、または同等の計算開始機能を解放します。
本開示は、特に、これらの形式の計算可能能力を企図する。
H100/H200 クラスの Confidential Computing を使用した展開では、ハードウェアによる GPU 測定を追加することで信頼チェーンを強化できます。このような GPU を使用しない展開では、適用可能なポリシーにとって重要な外部コンピューティング プレーンのプロパティを確立するための信頼できる別のメカニズムが必要になります。
したがって、セキュリティの強度は実装の信頼境界によって決まります。

[切り捨てられた]

## Original Extract

Problem Space European AI sovereignty is often approached through where infrastructure is located: where a model is hosted, where data is stored, which cloud region is selected, and which territorial rules apply to the service provider. These controls are important, but they do not fully answer a mo
[truncated]

Europe Cannot Govern AI by Cloud Region Alone: A Technical Architecture for AI Sovereignty | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Europe Cannot Govern AI by Cloud Region Alone: A Technical Architecture for AI Sovereignty
European AI sovereignty is often approached through where infrastructure is located : where a model is hosted, where data is stored, which cloud region is selected, and which territorial rules apply to the service provider. These controls are important, but they do not fully answer a more difficult execution-time question:
Can a specific AI computation be prevented from starting under unauthorized algorithmic logic, and can a specific AI-generated consequence be prevented from becoming effective in an unauthorized jurisdiction?
Cloud-region selection is principally a deployment decision. An AI workload may execute in an approved European region while its outputs are subsequently transmitted through external APIs, agentic tools, federated services, telecom networks, satellite infrastructure, downstream platforms, or recipient environments outside the intended jurisdictional boundary. IP-based geolocation provides only indirect evidence because routing can be proxied, tunneled, masked, abstracted, or separated from the actual environment in which an AI-generated consequence becomes effective. The underlying disclosure therefore identifies a gap between infrastructure placement and per-invocation jurisdiction enforcement .
The problem becomes more significant as AI systems evolve from passive information generators into systems capable of invoking tools, communicating with external services, controlling workflows, generating legally or commercially consequential outputs, and operating across multiple infrastructure and jurisdictional domains. In such systems, knowing that an AI service was nominally deployed in Europe does not establish that every computation used authorized algorithmic logic or that every resulting output became effective only within an approved jurisdictional environment.
The technical problem is therefore not merely where AI is hosted , but where and under what authority computation and external effect are permitted to occur .
The proposed architecture addresses this problem through two coordinated execution-time boundaries.
At the first boundary, immediately before an external compute plane is irreversibly authorized to begin inference, a protected domain derives or obtains an Algorithmic Logic Fingerprint (ALF) for the proposed model, checkpoint, orchestration configuration, or execution logic and compares it against an approved or pre-bound ALF condition. A computation-start capability is issued only when the required conditions are satisfied; otherwise authorization is withheld in a fail-closed manner.
At the second boundary, immediately before a candidate output becomes externally effective, the architecture evaluates the runtime governance state together with the intended jurisdictional context. Rather than relying solely on IP geolocation, it can use machine-verifiable infrastructure evidence including certificates, remote attestation, sovereign-cloud identifiers, telecom-domain identifiers, satellite-domain credentials, recipient public-key infrastructure, and other infrastructure-linked attributes.
This creates a different model of European AI sovereignty:
sovereignty is not merely declared by infrastructure location; it is technically enforced at the boundaries where computation begins and AI-generated consequences become effective.
Europe's pursuit of AI sovereignty raises a technical question that cloud-region selection, data-residency policies, and IP-based geolocation cannot completely answer: how can jurisdictional authority be enforced for a particular AI computation and a particular AI-generated consequence at the moment they become operationally effective?
This paper presents an execution-time, dual-boundary architecture for AI sovereignty . The architecture separates infrastructure placement from machine-enforceable authority and introduces two coordinated control points: a pre-computation authorization boundary immediately before an external compute plane is irreversibly authorized to begin inference, and an output finality boundary immediately before a candidate AI output becomes externally effective.
At the pre-computation boundary, a protected authorized domain, such as a Trusted Execution Environment (TEE), Hardware Security Module (HSM), or functionally equivalent protected environment, derives or verifies an Algorithmic Logic Fingerprint (ALF) representing the proposed model, checkpoint, orchestration configuration, or execution logic. Computation-start authority is released only when the candidate ALF satisfies the required authorization conditions. An absent, mismatched, revoked, expired, or otherwise unauthorized ALF causes authorization to be withheld in a fail-closed manner. Bulk inference may remain on external GPUs, accelerators, or cloud infrastructure; the protected domain controls the security-critical authority required for computation to begin.
During authorized inference, governance-relevant execution signals are captured and represented as a bounded machine-readable runtime behavioral governance descriptor . The architecture does not require reconstruction of the model's complete latent reasoning or semantic state. Instead, it can use observable execution-path information associated with the particular inference run.
Before external release, the runtime governance state can be cryptographically bound to a jurisdiction context representing the intended destination or release environment. Jurisdiction is then evaluated using non-IP, machine-verifiable infrastructure evidence , including cryptographic certificates, signed territorial or operator-domain claims, remote-attestation results, sovereign-cloud identifiers, telecom-domain identifiers, satellite-domain credentials, recipient public-key hierarchies, or equivalent infrastructure-linked credentials.
Release may consequently depend on the concurrent satisfaction of multiple predicates: approved algorithmic logic, approved runtime behavior, approved purpose, verified jurisdiction, binding integrity, freshness, expiry and revocation status, and, where required, authorization from both source- and destination-jurisdiction authorities. An acceptable AI output can therefore still be blocked when directed toward an unauthorized jurisdiction, while an approved jurisdiction cannot override disapproved runtime or algorithmic state.
The architecture does not propose that Europe isolate its AI infrastructure or require complete AI workloads to execute inside trusted hardware. Instead, it separates high-performance computation from authorization-critical enforcement. Existing GPU and cloud infrastructure may perform inference while protected infrastructure controls the capabilities required to start protected computation and to externalize its resulting output. The disclosure therefore seeks to make jurisdiction a per-invocation technical predicate , rather than merely a deployment-time assumption.
For European digital sovereignty, the resulting proposition is straightforward:
A cloud region can determine where infrastructure is deployed. It does not, by itself, determine whether a specific AI computation or AI-generated consequence possesses authority to become effective in a particular jurisdiction.
Execution-finality governance moves that decision to the machine-enforceable boundaries where computation begins and where output becomes externally effective.
Specialized Hardware Requirements and High-Throughput Scalability
Does the Architecture Require NVIDIA H100/H200 Confidential Computing, AMD SEV-SNP, or Other Specific Hardware?
No. The architecture is not dependent on NVIDIA H100/H200 Confidential Computing, AMD SEV-SNP, or any single processor, accelerator, TEE vendor, or cloud platform .
The architectural requirement is narrower: there must be a protected authorization domain capable of protecting the security-critical state and controlling release of the capability required at the relevant execution boundary. The disclosure expressly permits a TEE, HSM, secure enclave, secure element, confidential-computing module, hardware-rooted control module, or functionally equivalent cryptographically isolated enforcement domain. It also expressly states that the protected domain need not execute the entire AI inference workload.
Accordingly, NVIDIA Confidential Computing and AMD SEV-SNP should be understood as possible implementation substrates, not architectural prerequisites .
For example, an NVIDIA H100-or-later GPU supporting Confidential Computing can provide GPU attestation and cryptographic evidence concerning GPU hardware and software integrity. NVIDIA's current attestation architecture supports local and remote GPU attestation and can collect evidence across multiple GPUs. AMD SEV-SNP provides a different implementation substrate: a confidential VM with hardware-rooted isolation and attestation that can protect an authorization controller from an untrusted hypervisor.
Neither technology, by itself, implements the proposed execution-finality architecture.
Confidential computing answers:
Is this workload or execution environment running in an expected protected state?
Execution-finality governance asks:
Given that protected state, does this particular computation or output possess the required authority to cross the next consequential boundary?
Attestation can therefore become an input predicate to execution-finality rather than being the final authorization decision itself.
A deployment could, for example, use:
CPU confidential VM / TEE
→ hold policy, trust anchors, ALF repository and authorization keys
GPU or accelerator
→ perform bulk inference
GPU/VM attestation
→ prove relevant execution-environment state
Protected authorization controller
→ validate ALF + attestation + policy
Capability release
→ permit computation
Runtime evidence + destination credentials + jurisdiction context + freshness/revocation state
→ Finality Controller
→ release or withhold the output-finality capability.
This is consistent with the disclosure's explicit separation of high-performance computation from authorization-critical operations.
Current confidential-computing systems also demonstrate that secret or resource release conditioned on successful attestation is technically practical. For example, NVIDIA's Confidential Containers architecture describes attestation of CPU/GPU TEE state followed by policy-controlled release of secrets or resources.
Deployment Without Confidential GPUs
A confidential GPU is therefore not mandatory.
One implementation could place the authorization controller in an AMD SEV-SNP confidential VM , HSM, or another protected CPU-side environment while conventional GPUs perform inference externally.
Before inference, the protected controller could:
receive the invocation identifier and protected policy context;
obtain the proposed model/checkpoint/orchestration identity;
derive or verify the candidate ALF;
verify relevant compute-plane attestation or service credentials where available;
evaluate expiry and revocation state;
create a fresh invocation identifier or nonce; and
release a bounded scheduler token, model-access credential, decryption key, API authorization, or equivalent computation-start capability.
The disclosure specifically contemplates these forms of computation-enabling capability.
A deployment using H100/H200-class Confidential Computing could strengthen the trust chain by adding hardware-backed GPU measurements. A deployment without such GPUs would need another trustworthy mechanism for establishing the properties of the external compute plane that matter to the applicable policy.
The security strength therefore depends on the implementation's trust boundar

[truncated]

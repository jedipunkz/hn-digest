---
source: "https://fusionauth.io/blog/2026-ai-identity-report"
hn_url: "https://news.ycombinator.com/item?id=48459828"
title: "AI and Identity Survey Findings"
article_title: "We Surveyed More Than 300 Security Leaders on AI Identity. The Findings Are Counterintuitive"
author: "mooreds"
captured_at: "2026-06-09T13:08:01Z"
capture_tool: "hn-digest"
hn_id: 48459828
score: 1
comments: 0
posted_at: "2026-06-09T11:45:52Z"
tags:
  - hacker-news
  - translated
---

# AI and Identity Survey Findings

- HN: [48459828](https://news.ycombinator.com/item?id=48459828)
- Source: [fusionauth.io](https://fusionauth.io/blog/2026-ai-identity-report)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T11:45:52Z

## Translation

タイトル: AI とアイデンティティに関する調査結果
記事のタイトル: AI アイデンティティについて 300 人以上のセキュリティ リーダーを調査しました。発見は直感に反する
説明: FusionAuth は、AI アイデンティティ セキュリティについて 300 人以上のテクノロジおよびセキュリティ リーダーを調査し、顕著なパターンを発見しました。つまり、最も自信を持っている組織は侵害率が最も高く、導入アーキテクチャはガバナンスの成熟度よりも優れた結果を予測します。

記事本文:
私たちは 300 人以上のセキュリティ リーダーに AI アイデンティティに関する調査を実施しました。発見は直感に反する
/ ブログ カテゴリ カテゴリ
ログイン
デモを入手する
メインメニューを開く
私たちは 300 人以上のセキュリティ リーダーに AI アイデンティティに関する調査を実施しました。発見は直感に反する
FusionAuth は、AI ID セキュリティについて 300 人を超えるテクノロジーおよびセキュリティのリーダーを調査し、顕著なパターンを発見しました。つまり、最も自信を持っている組織は侵害率が最も高く、導入アーキテクチャはガバナンスの成熟度よりも結果をより正確に予測します。
私たちが調査した組織の 3 分の 2 は、過去 1 年間に確認された AI ID 侵害を経験しました。この数字だけでも憂慮すべきだ。しかし、AI セキュリティについての私の考え方を変えたのはこの発見ではありません。
私は、インフラストラクチャや開発者ツールのスタートアップ、フォーチュン 500 企業、公共部門の組織など、テクノロジー業界で 30 年近くを過ごしてきました。セキュリティはその取り組みの中心でした。それは私がセキュリティ データをどのように読み取るかを決定し、この研究に期待するものを決定しました。私は、より大規模で成熟した組織ほど、より高い侵害率を示すと予想していました。より優れたツール、より成熟した SOC 機能、より強力なフォレンジック機能により、より多くのインシデントが発見されます。セキュリティでは、検出の成熟度と確認されたインシデント率は連動して変化します。より多くのインシデントを検出することが、適切に装備されたセキュリティ プログラムの姿です。
そのため、FusionAuth が AI アイデンティティ セキュリティについて 300 人を超えるテクノロジおよびセキュリティのリーダーを調査したとき、最も自信があり、最も投資が行われ、最もガバナンスが成熟した組織がまさにその理由でより多くのインシデントを示すだろうと私は予想しました。
回答者に、組織の AI セキュリティにどの程度自信があるかを尋ねました。次に、過去 12 か月間で確認されたセキュリティ インシデントを経験した人が何人いるかを調べました。
10人中8人が最も自信を持っている

昨年、NT 組織では AI ID 関連のインシデントが確認されました。そこから侵害率は低下し、最も低い 2 つの層ではわずかな差異しかありません。
明らかな反論: 最も自信のある組織は、同時に最大かつ最も成熟した組織でもあります。より優れた検出プログラムにより、より多くのインシデントが検出されます。侵入率が高いということは、単に検出率が高いだけである可能性があります。しかし、その議論はサイズベースのデータに対して崩れます。検出の成熟度によってギャップが説明される場合、最大規模の組織ではそのギャップは最悪となるでしょう。そうではありません。
回答者はわずか 7 人で、統計的に有意ではありません。
組織の規模と信頼ギャップの間には相関関係はありません。最小の組織のギャップは最大ではなく、最悪です。何か構造的なことが起こっており、それを理解することが重要です。なぜなら、最もリスクにさらされている組織は、自分たちがそうであると信じている可能性が最も低いからです。
このことからわかることは、最も自信を持っている組織は、単により多くの検出を行っているだけではないということです。彼らはまた、本当に露出が増えています。彼らがより露出されている理由は、彼らが最も自信を持っている理由と同じであり、彼らは最も速く動いているからです。
信頼は速度を追跡します。ベロシティが攻撃対象領域を構築 #
データを理解するには、1 つのコンテキストが必要です。この調査に参加したすべての組織は、AI を迅速に推進するという現実のプレッシャーの下で活動しています。取締役会は、競合他社が AI 機能を提供しているのに、なぜ自社が提供していないのかを知りたいと考えています。投資家は注目している。あなた自身のチームは、競争上の優位性のように見える方法で AI を使用している同僚を目にします。このプレッシャーにより、基盤となるインフラストラクチャが求められている内容に対応する準備ができていない場合でも、採用の迅速化、導入の迅速化、および監査人を満足させるように設計されたガバナンスの文書化が促進されます。
また、私たちが発見した 80% のシャドウ AI 率についても説明します。従業員は未承認の AI ツールを接続していない

内部システムが不注意だからです。彼らのキャリアを生き残るためには、AI の流暢さを証明することが求められます。組織からのプレッシャーが個人的なものになってしまいました。そうなると、境界を守るのは非常に難しくなります。
採用データはその関係を強化します。 AI 人材を外部から雇用している組織では、85% の侵害率が確認されています。既存のチームをトレーニングしている組織は 33% でした。これは 2.6 倍低く、投資レベルと保険の満期を調整してもこの水準は維持されました。採用速度が導入速度を促進します。攻撃対象領域は次のとおりです。
信頼度スケールの上位にある組織はプロファイルを共有しています。最も成熟度の高いコホートの 92% は、包括的な AI ガバナンス ポリシーを持っています。 88% が AI セキュリティに多額の投資を行っています。机上では、彼らはすべて正しくやっています。また、これらの企業は、より多くの部門で承認された AI ツールを増やし、実際のユーザーにサービスを提供する AI を活用した製品機能を増やし、内部システムに対して API 呼び出しを行うエージェントの数を増やすことで、AI を最も早く本番環境に移行した組織でもあります。ガバナンスへの投資と導入速度は同時に行われており、攻撃対象領域はガバナンス層がカバーできる速度を超える速度で拡大しています。
私はテクノロジーのサイクルを最前線から数多く見てきました。 Kubernetes とクラウドネイティブの波は、毎週新しいツール、新しいフレームワーク、新しいスタートアップが登場し、信じられないほど速く感じられました。最先端のチームは、デプロイが完了する前に、デプロイしたばかりのものを置き換えることがよくありました。しかし、Kubernetes はインフラストラクチャ チーム内に存在するテクノロジーの物語でした。あなたの CFO はポッドが何であるかを知りませんでした。
AIはそのようには封じ込められません。新しいモデルは数四半期ではなく数日以内に置き換えられます。法的な契約書の草案の作成方法、財務部門の提案のレビュー方法、財務部門の提案の方法が再構築されています。

営業チームはアウトリーチ、サポートがチケットにどのように回答するかを書きます。あらゆる部門のリーダーがそれについて意見を持っています。ほとんどの企業はすでにそれを使用していますが、最初に IT 部門に問い合わせることはありませんでした。ベンダーや投資家による誇大広告が、運用の現実からこれほど切り離されていないのは 90 年代後半以来です。
慎重に行動し、意図的に機能を構築すると、目に見えるほど優れたセキュリティの成果が得られます。これについてはデータが明らかです。しかし、競合他社が前四半期に AI 機能をリリースしたのに、自社はまだリリースしていない取締役会会議で、これを持ち込むのは難しい推奨事項です。
彼らは何を知らないのかを知らない #
最も信頼度の高いコホートは、単に暴露されるだけではありません。彼らは、データが示すよりも AI のセキュリティをうまく管理できると心から信じています。最も自信のある組織の 99% は、自らを「非常に自信がある、または非常に自信がある」と述べています。 96％は包括的な政策を持っています。 AI エージェントのプロビジョニング方法、権限の範囲設定、異常な動作の監視方法、資格情報のローテーション方法、アクセスの取り消し方法、エージェントのアクティビティの監査方法を管理するライフサイクル プロセスを形式化しました。ほとんどの場合、6 つすべてが正式に文書化されています。取締役会レベルのセキュリティレビューを満足できる投資レベル。
そして、そのうち 84% が過去 12 か月以内に侵害が確認されています。
彼らがどのように準備を整えていると感じているかと、インシデントログが実際にどのように読んでいるかの間にあるギャップが、この調査結果の最も危険な部分です。暴露されていることを認識している組織は行動することができます。保護されていると確信している組織には、そのような兆候はありません。自信そのものが盲点になってしまうのです。そして、この環境で活動する脅威アクターは、新しいツール、新しい悪用パターン、AI エージェントの権限を調査する新しい方法など、まったく同じ速度で加速しています。死角はどんどん広がっていく

脅威の表面が拡大します。
これは、ポリシーやプロセスで修正するのが最も難しい問題のバージョンです。ガバナンスの枠組みを義務付けることができます。正式なライフサイクル手順を要求することができます。従業員数と投資を追加できます。組織の信頼レベルによって適切な質問ができない場合、その根底にある力学は変わりません。
各 AI エージェントがアクセスできる対象を、個々のエージェント レベルでリアルタイムに正確に調査することはできるでしょうか?
特定のエージェントがどのデータに対して、どのような権限で、何にアクセスしたかを証明できますか?
何かがエスカレートした直後ではなく、エスカレートする前に各エージェントが何をしているのかを確認できますか?
問題に間に合うようにアクセスを取り消すことができるでしょうか?
これらはアーキテクチャに関する質問です。現在、ほとんどの組織はガバナンスに関するインプットを測定しています。データはセキュリティ出力を測定します。これら 2 つの間のギャップは大きく、一貫性があり、インシデント ログに現れています。
これまで誰も測定できなかったアーキテクチャ変数 #
回答者に ID 導入モデルについて質問しました。ID をマルチテナント SaaS プラットフォーム、クラウド シングルテナント、セルフホスト インフラストラクチャのいずれで実行しているかについて尋ねました。次に、確認された事件を調べました。
マルチテナントの SaaS ID プラットフォーム上の組織は、自己ホスト型組織の 2 倍以上の割合で侵害されました (83% 対 38%)。ガバナンスの成熟度、ポリシーの適用範囲、投資額よりも、AI セキュリティの結果をより正確に予測します。
この調査では、展開アーキテクチャが侵害を直接引き起こすことを証明することはできません。組織の成熟度、業界、規制要件、セキュリティ文化などの他の要因が関係に寄与する可能性があります。データが示しているのは、ガバナンスの成熟度、ポリシーの適用範囲、投資ファイルよりも導入アーキテクチャの方が侵害の結果を予測しやすいということです。

ヴェルス。
マルチテナント SaaS ID は多くの場合、デフォルトの選択であり、最も一般的で、最も利用可能で、摩擦が最も少ないパスです。サインアップし、統合し、出荷します。 ID インフラストラクチャを共有サードパーティ プラットフォームにアウトソーシングすることは、多くのチームが行っていることです。なぜなら、それはほとんど機能しており、迅速に市場に投入することが真の優先事項であるためです。
セルフホスト ID またはシングルテナント ID を実行している組織は、意図的に異なる決定を下しました。それには、より多くの努力、事前にアーキテクチャを考えること、そしてアイデンティティ インフラストラクチャがアウトソーシングするにはあまりにも重要であるという明確な立場が必要でした。 ID は、すべての AI エージェント、サービス、ユーザーが何にアクセスできるかを決定するコントロール プレーンです。これらの組織は、共有プラットフォームに完全に依存するのではなく、コントロール プレーンに対するより高度な制御を選択しました。
データによれば、自己ホスト型によってリスクが排除されるとは言えません。導入モデルはストーリーの一部にすぎません。本当の問題は、ID が環境自体の内部でどのように実装されるか、つまり AI エージェントが認証する方法、権限の範囲がどのように設定されるか、認可の決定がどのように行われるか、マシン ID がどのように管理されるか、アクセスがどのように監視および取り消されるかということです。展開アーキテクチャはこれらのコントロールに影響を与えますが、それらを置き換えるものではありません。
確認された侵害率が 38% であるということは、依然として意味があります。これが示しているのは、アイデンティティ アーキテクチャが重要であるということです。導入モデルは、誰がインフラストラクチャを制御するかに影響しますが、組織が認証、認可、マシン ID、可視性、運用制御にどのように取り組むかにも影響します。 AI システムがアプリケーション、データ ストア、API 全体で自律的に動作し始めると、それらのアーキテクチャ上の決定がセキュリティ境界自体の一部になります。
組織文化の違いはデータにも現れます。自己ホスト型の組織は、エスカレートする前に、より多くの脅威を捕捉します

: ニアミス率は 44%、マルチテナント SaaS の場合は 10% です。また、シャドウ AI への露出も少なくなります: 56% 対 91%。より難しいアーキテクチャ上の選択をした組織は、より厳格なセキュリティ環境を運用する傾向もあります。それがより強固なセキュリティ文化を反映しているのか、運用上の優先順位の違いなのか、それともコントロールへの強い欲求を反映しているのかを、この調査だけから知ることは不可能です。明らかなことは、組織がアイデンティティに関して行うアーキテクチャ上の決定は、AI 関連のセキュリティ リスクの経験と強く相関しているということです。
自己ホスト型または分離された展開では、エクスポージャーの境界はユーザーが所有します。環境内で障害が発生すると、その環境の接続先にも影響が及びます。それは依然として現実のリスクですが、それはあなたが制御し、対応できる範囲のリスクです。
データセット全体で最もリスクの高いプロファイルは、成熟度の低い組織ではありません。これはその逆です。企業は本番環境で AI を実行し、AI は従業員全体に広く導入され、マルチテナントの SaaS ID インフラストラクチャ上で運用されています。このコホートでは、90% が感染が確認されました。 96% がシャドウ AI に直面しています。従業員は IT やセキュリティのレビューを行わずに AI ツールを接続しています。 98% が自社のセキュリティ体制に非常に自信を持っています。
なぜですか

[切り捨てられた]

## Original Extract

FusionAuth surveyed more than 300 technology and security leaders on AI identity security and found a striking pattern—the most confident organizations had the highest breach rates, and deployment architecture predicts outcomes better than governance maturity.

We Surveyed More Than 300 Security Leaders on AI Identity. The Findings Are Counterintuitive
/ Blog Categories Categories
Log In
Get a demo
Open main menu Education
We Surveyed More Than 300 Security Leaders on AI Identity. The Findings Are Counterintuitive
FusionAuth surveyed more than 300 technology and security leaders on AI identity security and found a striking pattern—the most confident organizations had the highest breach rates, and deployment architecture predicts outcomes better than governance maturity.
Two-thirds of the organizations we surveyed experienced a confirmed AI identity breach in the past year. That number alone should be alarming. But it's not the finding that changed how I think about AI security.
I've spent close to 30 years in tech: infrastructure and developer tooling startups, Fortune 500s, public sector organizations. Security has been central to that work throughout. It shapes how I read security data, and it shaped what I expected going into this research. I expected larger, more mature organizations to show higher breach rates: better tooling, more mature SOC functions, and stronger forensic capabilities mean finding more incidents. In security, detection maturity and confirmed incident rates move together. Finding more incidents is what a well-instrumented security program looks like.
So when FusionAuth surveyed more than 300 technology and security leaders on AI identity security , I expected the most confident, most invested, most governance-mature organizations to show more incidents for exactly that reason.
We asked respondents how confident they were in their organization's AI security. Then we looked at how many had experienced a confirmed security incident in the past 12 months.
Eight out of ten of the most confident organizations had a confirmed AI identity-related incident in the past year. Breach rates decline from there, with only slight variance at the two lowest tiers.
The obvious counter-argument: the most confident organizations are also the largest and most mature. Better detection programs mean more incidents found. Their higher breach rate could just be a higher detection rate. But that argument collapses against the size-based data. If detection maturity explained the gap, it would be worst in the largest organizations. It's not.
Only 7 respondents — not statistically significant.
There is no correlation between organizational size and the confidence gap. The smallest organizations have the worst gap, not the largest. Something structural is happening, and understanding it matters, because the organizations most at risk are the ones least likely to believe they are.
What it's telling you is this: the most confident organizations aren't just detecting more. They're also genuinely more exposed. The reason they're more exposed is the same reason they're most confident: they're moving the fastest.
Confidence Tracks Velocity. Velocity Builds Attack Surface #
Understanding the data requires one piece of context. Every organization in this survey operates under real pressure to move fast on AI. The board wants to know why competitors are shipping AI features and you're not. Investors are watching. Your own teams see peers using AI in ways that look like competitive advantage. That pressure drives faster hiring, faster deployment, and governance documentation designed to satisfy auditors, even when the underlying infrastructure isn't ready for what it's being asked to do.
It also explains the 80% shadow AI rate we found. Employees aren't connecting unauthorized AI tools to internal systems because they're careless. Their career survival demands demonstrating AI fluency. The organizational pressure has become a personal one. When that happens, the perimeter is much harder to protect.
The hiring data reinforces the relationship. Organizations hiring externally for AI talent had an 85% confirmed breach rate. Organizations training their existing teams had 33%. That's 2.6 times lower, and it held even when controlling for investment levels and policy maturity. Hiring velocity drives deployment velocity. The attack surface follows.
The organizations at the top of the confidence scale share a profile. Ninety-two percent in the highest-maturity cohort have comprehensive AI governance policies. Eighty-eight percent are investing significantly in AI security. On paper, they're doing everything right. They're also the organizations that have moved AI into production fastest, with more approved AI tools across more departments, more AI-powered product features serving real users, more agents making API calls to internal systems. The governance investment and the deployment velocity are happening at the same time, and the attack surface grows faster than the governance layer can cover it.
I've watched a lot of technology cycles from the front lines. Kubernetes and the cloud-native wave felt impossibly fast: new tools, new frameworks, new startups every week. Teams on the cutting edge were often replacing what they'd just deployed before they'd finished deploying it. But Kubernetes was a technology story that lived inside your infrastructure team. Your CFO didn't know what a pod was.
AI isn't contained that way. New models supersede each other within days, not quarters. It's reshaping how legal drafts contracts, how finance reviews proposals, how your sales team writes outreach, how support answers tickets. Every leader in every function has an opinion about it. Most are already using it and they didn't ask IT first. The hype from vendors and investors hasn't been this disconnected from operational reality since the late '90s.
Moving carefully and building capability deliberately produces measurably better security outcomes. The data is clear about this. But that's a hard recommendation to carry into a board meeting where your competitors shipped AI features last quarter and you haven't.
They Don't Know What They Don't Know #
The highest-confidence cohort isn't just exposed; they genuinely believe they have a better handle on their AI security than the data shows. Ninety-nine percent of the most confident organizations describe themselves as "very or extremely confident." Ninety-six percent have comprehensive policies. They've formalized the lifecycle processes that govern how AI agents are provisioned, how permissions get scoped, how anomalous behavior is monitored, how credentials are rotated, how access is revoked, and how agent activity gets audited. In most cases, all six are formally documented. Investment levels that would satisfy any board-level security review.
And 84% of them had a confirmed breach in the past 12 months.
The gap between how prepared they feel and how their incident logs actually read is the most dangerous part of this finding. An organization that knows it's exposed can act. An organization that is confident it's protected has no such signal. The confidence itself becomes the blind spot. And the threat actors operating in this environment are accelerating at exactly the same rate: new tooling, new exploitation patterns, new ways to probe AI agent permissions. The blind spot is expanding as the threat surface grows.
This is the version of the problem that's hardest to fix with a policy or a process. You can mandate a governance framework. You can require formalized lifecycle procedures. You can add headcount and investment. None of that changes the underlying dynamic if the organization's confidence level is preventing it from asking the right questions.
Can we scope exactly what each AI agent can access — at the individual agent level, in real time?
Can we prove what a specific agent accessed, under what permissions, against which data?
Can we see what each agent is doing before something escalates, not just after?
Can we revoke access in time to matter?
Those are architecture questions. Most organizations right now are measuring governance inputs. The data measures security outputs. The gap between those two things is large, it's consistent, and it's showing up in incident logs.
The Architectural Variable Nobody Measured Until Now #
We asked respondents about their identity deployment model: whether they run identity on a multi-tenant SaaS platform , cloud single-tenant, or self-hosted infrastructure. Then we looked at confirmed incidents.
Organizations on multi-tenant SaaS identity platforms were breached at more than twice the rate of self-hosted organizations: 83% versus 38%. More predictive of AI security outcomes than governance maturity, policy coverage, or investment volume.
This survey cannot prove that deployment architecture directly causes breaches. Other factors, including organizational maturity, industry, regulatory requirements, and security culture may contribute to the relationship. What the data does show is that deployment architecture was more predictive of breach outcomes than governance maturity, policy coverage, or investment levels.
Multi-tenant SaaS identity is often the default choice, the most common, most available, lowest-friction path. You sign up, you integrate, you ship. Outsourcing identity infrastructure to a shared third-party platform is what many teams do because it mostly works, and getting to market quickly is a real priority.
The organizations running self-hosted or single-tenant identity made a deliberately different decision. It required more effort, more architecture thinking upfront, and a clear position that identity infrastructure was too critical to outsource. Identity is the control plane that determines what every AI agent , service, and user can access. These organizations chose greater control over that control plane rather than relying entirely on a shared platform.
The data doesn't say self-hosted eliminates risk. Deployment model is only part of the story. The real question is how identity is implemented inside the environment itself: how AI agents authenticate, how permissions are scoped, how authorization decisions are made, how machine identities are governed, and how access is monitored and revoked. Deployment architecture influences those controls, but it does not replace them.
A 38% confirmed breach rate is still meaningful. What it shows is that identity architecture matters. Deployment model influences who controls the infrastructure, but it also influences how organizations approach authentication, authorization, machine identities, visibility, and operational control. When AI systems begin acting autonomously across applications, data stores, and APIs, those architectural decisions become part of the security boundary itself.
The organizational culture difference shows up in the data too. Self-hosted organizations catch more threats before they escalate: 44% near-miss rate versus 10% for multi-tenant SaaS. They also face less shadow AI exposure: 56% versus 91%. Organizations that made the harder architectural choice also tend to run tighter security environments. Whether that reflects stronger security cultures, different operational priorities, or a greater desire for control is impossible to know from the survey alone. What is clear is that the architectural decisions organizations make around identity correlate strongly with how they experience AI-related security risk.
In a self-hosted or isolated deployment, you own the exposure boundaries. A failure in your environment reaches what your environment connects to. That's still a real risk, but it's scoped risk you control and can respond to.
The highest-risk profile in the entire dataset is not a low-maturity organization. It's the opposite: companies running AI in production, AI deployed widely across the workforce, operating on multi-tenant SaaS identity infrastructure. In this cohort, 90% had a confirmed incident. Ninety-six percent face shadow AI: employees connecting AI tools without IT or security review. Ninety-eight percent are very or extremely confident in their security posture.
Why th

[truncated]

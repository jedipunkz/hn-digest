---
source: "https://cleverhans.io/latest-research.html"
hn_url: "https://news.ycombinator.com/item?id=48413296"
title: "AI Agents Enable Adaptive Computer Worms"
article_title: "CleverHans Lab - Latest research"
author: "speckx"
captured_at: "2026-06-05T16:07:15Z"
capture_tool: "hn-digest"
hn_id: 48413296
score: 3
comments: 0
posted_at: "2026-06-05T14:48:46Z"
tags:
  - hacker-news
  - translated
---

# AI Agents Enable Adaptive Computer Worms

- HN: [48413296](https://news.ycombinator.com/item?id=48413296)
- Source: [cleverhans.io](https://cleverhans.io/latest-research.html)
- Score: 3
- Comments: 0
- Posted: 2026-06-05T14:48:46Z

## Translation

タイトル: AI エージェントが適応型コンピューター ワームを実現
記事タイトル: CleverHans Lab - 最新の研究

記事本文:
AI エージェントが適応型コンピューター ワームを可能にする
人工知能のセキュリティを強化するための新しい知識を追求する中で、私たちは社会全体に影響を与えるサイバーセキュリティの脅威を発見しました。
ジョナス・グアン *†1,2
トム・ブランチャード *1,2
ハンナ・フェルスター *3
ヘンルイ・ジア *1,2
ガブリエル・ファン 4
ニコラス・パペノット †1,2
1 トロント大学
2 ベクトル研究所
3 ケンブリッジ大学
4 ServiceNow
* 平等な貢献
† 責任著者
全文はプレプリントとして入手できます。
大規模言語モデル (LLM) は、ツールへのアクセスと組み合わせて、エージェント AI システムが複雑なタスクを解決できるようにする構造化された問題解決能力を実証しています。これらの機能が自己複製エージェントに組み込まれると、根本的に新しいサイバーセキュリティの脅威が生成されることを示します。これは、ターゲット固有の攻撃戦略を考案してマシンを制御し、ネットワーク全体に拡散する適応型コンピューター ワームです。侵害された各マシンはワーム自身のインフラストラクチャの一部となり、さらなる攻撃の計算能力や到達範囲を提供します。
コンピューター ワームは、人間の介入なしにネットワーク全体に広がる自己複製マルウェアです。 WannaCry ワーム (2017) は、単一の脆弱性を悪用して 150 か国にわたる重要なインフラストラクチャを混乱させました。従来のワームは、悪用する特定の脆弱性にパッチを適用することで阻止できます。私たちの適応型ワームをこの方法で止めることはできません。再帰的推論ループを使用して、増殖する際にさまざまな脆弱性を検出して悪用します。
私たちはこれらの機能を制御された実験で実証します。プロトタイプの AI 駆動型ワームは、ローカルで実行されるオープンウェイト LLM を利用し、一般的な企業ネットワークの脆弱性を持つ Linux、Windows、IoT デバイスの異種ネットワーク全体に伝播します。実験はISOで行われました

遅れた仮想ネットワーク。
私たちは、この研究がサイバー脅威の状況に対する AI の影響の 3 つの重要な側面を浮き彫りにしていると考えています。
それは脅威能力の質的変化を確立します。このワームは、修正された悪用コードを、各ターゲットの脆弱性にリアルタイムで適応する目標指向の推論に置き換えます。当社のエージェントは、ネットワークに接続されたデバイス間で自己複製し、システムの制御を破壊し、盗まれたリソースで自立します。
AI 駆動型ワームに必要なのは、単一のローカル GPU で実行できるオープンウェイト モデルのみです。商用 AI プラットフォームには依存しません。これにより、サービスの拒否、コンテンツのフィルタリング、レート制限などのベンダーの集中安全制御が構造的に無関係になります。このワームの階層型設計では、侵害された GPU を搭載した各ノードがダウンストリーム デバイス上の軽量エージェントの推論を提供し、攻撃対象領域がネットワーク化されたデバイスにまで拡大されます。
サイバーセキュリティにおける従来の経済的障壁は崩壊します。ワームは被害者自身の計算リソースを寄生的に使用し、攻撃者の限界コストをゼロにします。消費者デバイスが LLM 推論をサポートするようになるにつれて、そのような攻撃者が利用できる推論リソースもそれに応じて増加します。
この研究は、自律型サイバー攻撃が理論上のリスクから実証済みの能力にまで到達し、AI 研究、サイバーセキュリティ、公共政策にまたがる課題となっているという経験的証拠を提供します。私たちは、この移行には、オープンおよびクローズドウェイト モデルのエコシステム全体にわたるモデル機能の厳密で透明性のある評価が必要であると考えています。
私たちのすべての仕事の背後にある原動力は、人工知能のセキュリティを強化することです。最近、AI の安全性に関する公の議論は、安全性が知られている最大かつ最も強力な AI モデルの機能に焦点を当てています。

悪用される可能性のある、これまで発見されていない脆弱性を発見できる可能性があります。対照的に、より小規模なオープンウェイト AI モデル (誰でもインターネットからダウンロードできる) は、重大なサイバーセキュリティの脅威をもたらすのに必要な機能が欠けているとして無視されてきました。
私たちはそうではないのではないかと懸念し、公共政策の議論の基礎となっている前提が科学的に擁護できるかどうかを調査することにしました。私たちは、小規模で無料のモデルは実際の脅威をもたらすには弱すぎて信頼性が低いのでしょうか、それともネットワーク全体に対してより広範な攻撃を開始するために適応させることができるのでしょうか?言い換えれば、私たちはサイバーセキュリティの脅威の状況を本当に理解しているのでしょうか?
私たちは、小規模な無料の AI モデルのみを使用して、各マシンの固有の弱点 (業界から報告されたばかりの脆弱性やパスワードの再利用などの構成ミスを含む) を自律的に特定し、それらを悪用できる AI 駆動のコンピューター ワームを作成することが可能であることを発見しました。これにより、コンピューティング能力を乗っ取り、ラップトップ、カメラ、その他すべてのオンラインデバイスなどの通常のデバイスを乗っ取り、サーバーやネットワークに自分自身をコピーしてデータを盗んだり、新たな攻撃を仕掛けたりすることができます。私たちは、最新かつ最も強力な AI モデルを使用せずにこれを実行しました。この新たな脅威に対する単一の防御策はありません。
私たちは、新たな脅威をより深く理解し、それらに対する防御を評価できるようにするサイバーセキュリティ研究における確立された実践に従って、制御された環境で概念実証のプロトタイプを作成しました。この概念実証の構築では、検出や削除を複雑にする標準的なマルウェア機能の実装を意図的に省略しました。
この調査により、世界が直面する準備ができていない新たなサイバーセキュリティの脅威が明らかになりました。現代生活のほぼすべての側面がエネルギーに依存しています。

飲料水と廃棄物管理システム、食料や物品へのアクセス、エネルギー、金融システム、通信、医療、教育、交通システム、政府など、2 台のコンピューターで構成されるリスクは膨大です。
さらに、この設計は 1 台のマシン上で実行される小規模なモデルを使用して構築されているため、サイバー攻撃の経済学は根本的に変化しようとしています。サイバー攻撃は通常、攻撃を実行するのに必要な時間と比較的膨大なコンピューティング リソースのため、最も価値の高いターゲットに焦点を当てます。この低コスト設計は、インターネットに接続されているすべてのマシンが潜在的なターゲットであることを意味します。保持しているデータがなければ、次の攻撃の発射台として使用されることになります。
研究者、業界、政策立案者、そして一般の人々は、この新たなサイバーセキュリティの脅威に緊急に対処するために団結する必要があります。
対抗策を構築するためにサイバーセキュリティコミュニティを動員する必要性を考えると、敵対的な意図を持つ可能性のある犯罪者や国家ハッカーなど、他の場所でも同様の研究が進行している可能性があり、これらの調査結果を開示しないことは非倫理的です。
論文を公開する前に、私たちは研究結果を適切な国家科学、安全保障、防衛機関と共有し、攻撃者の能力を向上させることなくこの研究を責任を持って公開する方法についてカナダ当局にアドバイスを求めました。
私たちが調査結果を公開したのは、あらゆる分野（政府、産業界、学界、中小企業、個人）の意思決定者が、間もなく直面する可能性のある脅威をより明確に理解し、対策の研究を加速するために結集し、国家安全保障、企業競争力、個人のサイバー安全性の問題について情報に基づいた意思決定を行えるようになるためです。重要なのは、これは

rkは公的資金が投入された学術機関で実施され、その成果は社会の利益のためにより広範な研究コミュニティに利用可能です。
私たちの研究は、送受信のデジタル干渉を防ぐ安全な環境で実施されました。私たちは、有益なアプリケーションと潜在的に有害なアプリケーションの両方の機能を含むサイバーセキュリティ作業の確立されたベスト プラクティスに従い、関連するすべての大学研究局、情報セキュリティ局、およびカナダ当局と協力しました。
脅威が理解されたことで、以前には存在しなかった同様のサイバー脅威を検出して防御する機会が得られます。この新たな脅威について警鐘を鳴らすとともに、私たちは同様に設計されたサイバー兵器を検出し防御できる対抗策の開発にも注力しています。トロント大学全体では、AI の安全性と関連政策のニーズに関する重要かつ画期的な研究が、シュワルツ ライスマン技術社会研究所のシチズン ラボのさまざまな学部で、カナダ高等研究院 (CIFAR) およびベクター研究所と協力し、政府機関および必要に応じて業界パートナーと協力して進行中です。
私たちの研究は、新たな脅威の状況に警鐘を鳴らすとともに、適切な設計により、シンプルな言語モデルと適度なコンピューティング能力を利用して、信じられないほど複雑な問題を解決できることを実証しています。私たちのアプローチは、他の分野でも積極的に応用できる可能性があります。私たちは、この方法論が社会全体に利益をもたらす幅広い積極的な用途に適応できると信じています。たとえば、医療の進歩のための研究や持続可能なエネルギーソリューションの可能性を特定する際に、より早く適切な決定に到達するためです。
私たちはインプリメントを一般に公開しません

メンテーション。私たちはトロント大学と協力して、資格のある研究者が防御研究目的でアクセスを要求できる審査プロセスを確立しています。
いいえ、私たちの研究用プロトタイプは、ハイパーバイザーによって強制的に分離された包含された仮想ネットワーク内でのみ構築およびテストされました。その環境の外にデプロイされたことはありません。
はい、意図的にです。悪意のある攻撃者が同様のマルウェアを構築するのに実質的に役立つ可能性がある特定の方法論の詳細 (エージェントの推論グラフやツール ハーネスなど) や実験の詳細 (AI モデルなど) を省略しました。私たちは、悪用を可能にする青写真を提供することなく、科学的精査に十分な脅威を信頼できるものにするのに十分な情報を共有しました。
公開する前に、攻撃者の能力を向上させることなくこの研究を責任を持って公開する方法について、適切な国家安全保障および防衛機関のアドバイスを求めました。この脅威が理解されたことで、以前には存在しなかった同様のサイバー兵器を検出して防御するための対策を構築する機会が生まれました。
セキュリティ コミュニティが適応型 (AI 駆動型) コンピュータ ワームに対する防御策を研究し構築できるように、この脅威の実証的証拠を公開することが不可欠です。私たちは調査結果をカナダの科学、安全保障、国防当局と共有し、敵対者の能力を向上させることなくこの研究を責任を持って開示する方法についてアドバイスを求めました。私たちは、特に私たちが導入した緩和策を考慮すると、社会が生成的な敵対者に備えることができるという利点が二重使用のリスクを上回ると結論付けました。出版前に、この論文は、悪意を持ってこれらの発見を利用する人々に有利になる詳細を明らかにすることを避けるために大幅に変更されました。
いいえ。

私たちは、ワームに隠蔽機能を装備しないことを意図的に選択しました。ワームには、その足跡を覆い隠したり、ネットワークのフットプリントを最小限に抑えたりするよう指示されておらず、そのためのツールもありません。これは、悪用のリスクをさらに制限するための意識的な方法論的な選択でした。
現在のプロトタイプでは、非標準ポートでのビーコン コールバック、SSH 公開キーの自動挿入、ホスト全体での体系的な資格情報の再利用など、一貫した動作シグネチャが残されています。これらは、ネットワーク監視および侵入検知システムの具体的なターゲットです。これらのシグネチャは概念実証範囲の成果物であることに注意してください。将来の敵対者が同じ推論能力を回避戦略に向ける可能性があります。
私たちの実験では、プロトタイプは約 5 日でネットワークの半分に到達しました。これは、各ターゲットが偵察、戦略策定、ペイロード生成のために何百もの LLM 推論呼び出しを必要とするため、従来のワームよりも遅くなります。これにより、防御側は検出と対応のためのウィンドウを長くすることができます。ただし、このウィンドウは、推論ハードウェアとモデルの効率が向上するにつれて圧縮されます。
検出。現在のプロトタイプでは、一貫した動作シグネチャが残されています。非標準ポートでのビーコン コールバック、SSH 公開キーの自動挿入などです。

[切り捨てられた]

## Original Extract

AI Agents Enable Adaptive Computer Worms
In our pursuit of new knowledge to enhance the security of artificial intelligence, we uncovered a cybersecurity threat with implications across society.
Jonas Guan *†1,2
Tom Blanchard *1,2
Hanna Foerster *3
Hengrui Jia *1,2
Gabriel Huang 4
Nicolas Papernot †1,2
1 University of Toronto
2 Vector Institute
3 University of Cambridge
4 ServiceNow
* Equal contribution
† Corresponding author
The full paper is available as a preprint.
Large language models (LLMs) now demonstrate the capacity for structured problem-solving that, combined with tool access, enables agentic AI systems to solve complex tasks. We show that when these capabilities are embedded in a self-replicating agent, they produce a fundamentally new cybersecurity threat: an adaptive computer worm that devises target-specific attack strategies to gain control of machines and spread across networks. Each compromised machine becomes part of the worm’s own infrastructure, providing compute or reach for further attacks.
A computer worm is self-replicating malware that spreads across a network without human intervention. The WannaCry worm (2017) disrupted critical infrastructure across 150 countries by exploiting a single vulnerability. Traditional worms can be stopped by patching the specific vulnerability they exploit. Our adaptive worm cannot be stopped this way: it uses a recursive reasoning loop to detect and exploit diverse vulnerabilities as it propagates.
We demonstrate these capabilities in a controlled experiment: a prototype AI-driven worm powered by an open-weight LLM running locally, propagated across a heterogeneous network of Linux, Windows, and IoT devices with common corporate network vulnerabilities. The experiment was conducted in an isolated virtual network.
We believe this work highlights three important dimensions of the impact of AI on the cyberthreat landscape:
It establishes a qualitative shift in threat capability. The worm replaces fixed exploitation code with goal-directed reasoning that adapts to the vulnerabilities of each target in real time. Our agent self-replicates across networked devices, subverts control of systems, and self-sustains on stolen resources.
The AI-driven worm requires only an open-weight model that can run on a single, local GPU. It does not rely on any commercial AI platform. This renders vendors’ centralized safety controls, including service refusal, content filtering, and rate limits structurally irrelevant. The worm’s tiered design, where each compromised GPU-equipped node provides reasoning for lightweight agents on downstream devices, extends the attack surface to any networked device.
The traditional economic barrier in cybersecurity collapses. The worm parasitically uses the victims’ own computational resources, reducing the attacker’s marginal cost to zero. As consumer devices increasingly support LLM inference, the reasoning resources available to such adversaries grow accordingly.
This work provides empirical evidence that autonomous cyberoffence has crossed from theoretical risk to demonstrated capability, a challenge that spans AI research, cybersecurity, and public policy. We believe this transition demands rigorous, transparent evaluation of model capabilities across the open and closed-weight model ecosystems.
The driving motivation behind all our work is to enhance the security of artificial intelligence. Recently, public discussion about AI safety has focused on the capabilities of the largest and most powerful AI models that are known to be capable of finding previously undiscovered vulnerabilities that could be exploited. In contrast, smaller open-weight AI models (that anyone can download off the internet) have been dismissed as lacking the capabilities necessary to present a significant cybersecurity threat.
We were concerned this was not the case and set out to discover if the assumptions underpinning the public policy debate were scientifically defensible. We asked: Are small, free models too weak and unreliable to pose a real threat, or could they be adapted to launch much broader attacks against entire networks? In other words, do we really understand the cybersecurity threat landscape?
We discovered that it is possible to create an AI-driven computer worm, using only small, free AI models, that can autonomously identify each machine’s unique weak points (including vulnerabilities just reported by industry and misconfigurations such as reused passwords) and exploit them, hijacking computing power to take over regular devices such as laptops, cameras and everything else online, and then copying itself onto servers and networks to either steal data or launch new attacks. We did this without using the newest, most powerful AI models. There is no single defence against this new threat.
We created a proof-of-concept prototype in a controlled environment, following a well-established practice in cybersecurity research that enables a better understanding of emerging threats and evaluation of defences against them. In the construction of this proof-of-concept, we intentionally omitted implementing any standard malware capabilities that complicate detection or removal.
This research uncovered a new cybersecurity threat the world is not prepared to face. With almost every aspect of modern life dependent on networked computers — drinking water and waste management systems, access to food and goods, energy, our financial system, communications, health care, education, transportation systems, government and so much more — the risk is enormous.
What’s more, because this design is built using a small model that runs on a single machine, the economics of cyberattacks are about to radically shift: Cyberattacks typically focus on the most high-value targets, due to the time and comparatively enormous computing resources required to wage an attack. Now, this low-cost design means every machine connected to the internet is a potential target — if not for the data it holds, then as a launching pad for the next attack.
Researchers, industry, policymakers and everyday people need to come together with urgency to address this new cybersecurity threat.
Given the need to mobilize the cybersecurity community to build countermeasures, as similar research is likely underway elsewhere, including by criminal and state hackers who may have hostile intent, not disclosing these findings would be unethical.
Before making our paper public, we shared our findings with the appropriate national science, security and defence bodies, and sought advice from Canadian authorities on how to responsibly disclose this research without improving attackers’ capabilities.
We made our findings public so decision-makers in all areas (government, industry, academia, small- and medium-sized businesses, individuals) will have a clearer understanding of the threat we could soon face, can mobilize around accelerating research into countermeasures, and are better positioned to make informed decisions on matters of national security, corporate competitiveness and personal cyber safety. Crucially, because this work was done at a publicly funded academic institution, the findings are available to the broader research community for the benefit of society.
Our research was conducted in a safe environment that prevented incoming and outgoing digital interference. We followed established best practices for cybersecurity work involving capabilities with both beneficial and potentially harmful applications and collaborated with all the relevant university research and information security offices and Canadian authorities.
Now that the threat is understood, there is an opportunity to detect and defend against similar cyberthreats that did not exist before. Along with sounding the alarm about this emerging threat, we are turning our attention to developing the countermeasures that can detect and defend against similarly designed cyberweapons. Across the University of Toronto, important and groundbreaking work on AI safety and related policy needs is underway at the Schwartz Reisman Institute for Technology and Society, Citizen Lab, in various faculties, with the Canadian Institute for Advanced Research (CIFAR) and Vector Institute, and in partnership with government agencies and, where appropriate, industry partners.
Along with sounding the alarm on a new threat landscape, our research demonstrates that with the right design, simple language models and modest computing power can be harnessed to solve incredibly complex problems. Our approach could be used in other disciplines for positive applications. We believe the methodology can be adapted for a wide range of positive uses with benefits across society — to arrive at sound decisions sooner in research for medical advances or identifying potential sustainable energy solutions, for example.
We will not be publicly releasing our implementation. We are working with the University of Toronto to establish a vetting process through which qualified researchers may request access for defensive research purposes.
No. Our research prototype was built and tested exclusively in a contained virtual network with hypervisor-enforced isolation. It has never been deployed outside that environment.
Yes, intentionally. We omitted certain methodological details (such as the agent’s reasoning graph and tool harness) and experimental specifics (such as the AI model) that could materially help a malicious actor construct similar malware. We shared enough information to make the threat credible enough for scientific scrutiny without providing a blueprint that would enable misuse.
Before sharing it publicly, we sought the advice of the appropriate national security and defence bodies on how to responsibly disclose this research without improving attackers’ capabilities. Now that the threat is understood, there is an opportunity to build countermeasures to detect and defend against similar cyberweapons that did not exist before.
Publishing empirical evidence of this threat is essential so that the security community can study and build defences against adaptive (AI-driven) computer worms. We shared our findings with Canadian science, security and defence authorities and sought advice on how to responsibly disclose this research without improving adversaries’ capabilities. We concluded that the benefits — enabling society to prepare for generative adversaries — outweighed the dual-use risks, especially given the mitigations we put in place. Prior to publication, the paper was significantly altered to avoid revealing details that would be advantageous to those who would use these findings with malicious intent.
No. We deliberately chose not to equip the worm with concealment capabilities — it is not instructed to cover its tracks or minimize its network footprint, and it has no tools to do so. This was a conscious methodological choice to further limit the risk of misuse.
The current prototype leaves consistent behavioural signatures: beacon callbacks on non-standard ports, automated injection of SSH public keys, and systematic credential reuse across hosts. These are concrete targets for network monitoring and intrusion detection systems. Note that these signatures are artefacts of our proof-of-concept scope — a future adversary could direct the same reasoning capabilities toward evasion strategies.
In our experiments, the prototype reached half the network in approximately five days. This is slower than traditional worms because each target requires hundreds of LLM inference calls for reconnaissance, strategy formulation, and payload generation. This affords defenders a longer window for detection and response — but this window will compress as inference hardware and model efficiency improve.
Detection. The current prototype leaves consistent behavioural signatures: beacon callbacks on non-standard ports, automated injection of SSH public key

[truncated]

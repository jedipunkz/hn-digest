---
source: "https://www.cybersecurity-insiders.com/the-meta-ai-instagram-hack-wasnt-about-authentication-it-was-about-authorization/"
hn_url: "https://news.ycombinator.com/item?id=48446476"
title: "Meta AI Instagram Hack Wasn't About Authentication. It Was About Authorization"
article_title: "The Meta AI Instagram Hack Wasn't About Authentication. It Was About Authorization. - Cybersecurity Insiders"
author: "mooreds"
captured_at: "2026-06-08T16:29:01Z"
capture_tool: "hn-digest"
hn_id: 48446476
score: 1
comments: 0
posted_at: "2026-06-08T15:17:00Z"
tags:
  - hacker-news
  - translated
---

# Meta AI Instagram Hack Wasn't About Authentication. It Was About Authorization

- HN: [48446476](https://news.ycombinator.com/item?id=48446476)
- Source: [www.cybersecurity-insiders.com](https://www.cybersecurity-insiders.com/the-meta-ai-instagram-hack-wasnt-about-authentication-it-was-about-authorization/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T15:17:00Z

## Translation

タイトル: メタ AI Instagram ハックは認証に関するものではありませんでした。それは認可に関するものでした
記事のタイトル: メタ AI Instagram ハックは認証に関するものではありませんでした。それは認可に関するものでした。 - サイバーセキュリティ内部関係者
説明: AI は急速に進化しており、生成 AI (GenAI) の導入により、人間がこのテクノロジーと対話し活用する方法に革命が起きています。 GenAIは

記事本文:
フェイスブック
リンクトイン
×
ニュース
サイバー攻撃
ニュース
すべてのサイバー攻撃 データ漏洩 ID 詐欺 インサイダー脅威 マルウェア フィッシング 量子ランサムウェア ソーシャル エンジニアリング サプライ チェーンのセキュリティ脆弱性
サイバー攻撃
米国政府、人類神話を利用してサイバー攻撃を開始
Carnival Corporation のデータ侵害が 6 月のアカウント侵害事件の波を引き起こす
サーバー ファームが顧客をマルウェア攻撃から守る方法
カーニバルのデータ侵害は600万隻の巡洋艦に影響を与える可能性がある
メタ AI Instagram ハックは認証に関するものではありませんでした。それは認可に関するものでした。
6 月初めに攻撃者が Meta の AI サポート チャットボットを騙して Instagram アカウントを乗っ取ったとき、報道のほとんどは侵害自体に焦点を当てていました。しかし、この事件は、より広範でより重大な問題をよく示している。セキュリティ業界は、AI が何を許可されているかをほとんど無視しながら、AI の発言を制御することに多額の投資を行ってきた。
Meta のボットは、誰が質問したかについては何も検証しませんでした。新しい電子メール アドレスが有効であることを確認するための確認コードを攻撃者に送信することまで、指示されたことを実行しただけです。より成熟した認証フレームワークを AI エージェントに適用し始めるまでは、このようなインシデントがさらに発生するでしょう。
攻撃自体は素直だった。攻撃者は VPN を使用して被害者の位置を偽装しました。これにより、攻撃者の位置が被害者の位置から遠い場合に作動する特定の保護が回避されました。次に、攻撃者は実験的なメタ チャットボットに新しいメール アドレスをアカウントに追加するよう依頼しました。チャットボットは、新しいアドレスが有効であることを確認するために確認コードを電子メールで送信しました。役に立とうとしていたのです！攻撃者は新しい電子メール アドレスを確認し、パスワードをリセットする機会を与えられ、AC の制御を獲得しました。

カウント。
ほとんどの攻撃は、パッチできる 1 つの単純な穴ではありません。これらは脆弱性をつなぎ合わせて権限を昇格させたり、貴重なアカウントを乗っ取ったりします。このインシデントで公表された攻撃の詳細に基づくと、この脆弱性チェーンの失敗には次のものが含まれます。追加のセキュリティ対策が講じられているかどうかを判断するために IP の位置に依存する。チャットボットがユーザーのプライマリ電子メールを変更できるようにする。古い電子メール アドレスではなく、新しい電子メール アドレスからのみ確認コードを要求します。そして、それらの確認コードをパスワードのリセットを許可するのに十分なものとして扱い、チャットボットがそれを容易にしました。これらのいずれかの周囲にガードレールがあれば、このバージョンの攻撃は阻止されたでしょう。
認証と認可 — そしてそれが AI にとって重要な理由
認証とは、その人が誰であるかです。権限は彼らができることです。 AI エージェントに関する認証は比較的よく理解されている問題ですが、認可の決定はアプリケーションの奥深くにまで及び、通常はビジネス モデル固有です。これらは多くの場合、人間によって設計されたソフトウェア、または動きの遅い人間のいずれかのために設計されました。 AI エージェントは、ソフトウェアのスピードと人間のイノベーションを組み合わせて、エッジケースやホールを大規模に発見します。
認証が完璧だったとしても、メタ事件のより深刻な失敗は、エージェントがアカウント乗っ取りと同等のアクションを実行する権限を与えられていたことです。そして、それは業界が投資が不足している部分です。
AI プロジェクトが特にこの傾向が強い理由
AI チャットボットをサポート システムに固定しても、新たな種類の脆弱性は導入されませんでした。しかし、AI プロジェクトを成功させるための取り組みにより、システムが過剰な権限を与える傾向にあるため、このような穴が存在する可能性が高くなります。
さらに大きな問題は、サービス、機能、API を適切な手順なしに AI エージェントに公開していることです

それらの実際の有用性や、攻撃者が既存のホールを見つけて悪用するためにそれらをどのように利用できるかについても取り上げます。この場合、Meta はチャットボットが便利で役に立つものであることを望んでいましたが、それにはアクセスが必要です。しかし、彼らはアクセスを与えすぎました。
このパターンはすでに他の場所でも現れています。 2024 年、AI エージェントは、明示的にそうしないよう指示されていたにもかかわらず、ユーザーに騙されて 47,000 ドルの仮想通貨を送金させられました。 Lenovo チャットボットは、細工された製品クエリに基づいてセッション Cookie を公開するように操作されました。これは完全に直接的なアカウント アクセスではありませんが、根底にある目的は同じです。
これらは孤立した吸虫ではありません。これらはより広範な問題の初期の兆候であり、エージェントにさらに多くの権限が与えられるにつれてさらに大きくなるでしょう。この攻撃パターンがヘルスケア アプリケーションに適用された場合、アカウント乗っ取りの動きは同様になりますが、その結果は悲惨なものになるでしょう。たとえば、個人の健康データがプライバシー ルールに違反して流出することになります。
この事件はまだ一般的な問題ではありません。しかし、AI の導入の軌跡に基づくと、より多くのエージェントにより多くの権限が与えられるにつれて、インシデントの数は確実に増加すると思われます。
エージェントが稼働する前にネゴシエートできない承認制御がいくつかあります。まず、エージェントを有用な最小限のものに制限し、範囲を慎重に拡大します。 2 番目に、手動評価と自動評価の両方を使用して、エージェントが実行できる真の限界を判断します。第三に、提供されたアクセスを評価し、人によるレビューが必要なものとそのレビューをどのように行うべきかを決定します。
古い緩和策、新しい困難
セキュリティは多面的であり、これまでもそうでした。多層防御、最小権限の原則、人間による監視など、過去に機能したのと同じ緩和策がここでも機能します。体系的に適用する必要があるだけです。原則は古いし、まあまあ

分かりました。新しくて難しいのは、非決定的で、大規模に調査され、驚くほど広範囲にアクセスできるシステムにそれらを適用することです。
業界はAIに悪いことを言わせないことにかなり重点を置いています。 AIがやろうとしていることをAIに許可すべきかどうかを完全に見落とさない限り、それは問題ありません。私たちがより成熟した認可フレームワークに固執し始めるまでは、このような事件は今後も起こり続けるでしょう。
Dan Moore は、FusionAuth の CIAM 戦略およびアイデンティティ標準のシニア ディレクターであり、B2C 製品の方向性を推進し、認証、開発者のアイデンティティ、および AI エージェントのセキュリティ保護という新たな課題に関する主要なスポークスマンを務めています。 CTO、AWS 認定インストラクター、エンジニアリング マネージャーとしての勤務を含む、25 年以上のソフトウェア エンジニアリングの経験を持つ彼は、アイデンティティとセキュリティのカンファレンスで定期的に講演しており、認証と開発者のセキュリティに関する人気のポッドキャスト ゲストでもあります。彼は、『Letters to a New Developer』 (Apress) の著者であり、『97 Things Every Cloud Engineer Should Know』 (O'Reilly) の寄稿者でもあります。
新しい研究は、AI がオンライン脅威を加速するにつれて増大するデジタルトラスト危機を浮き彫りにします
ホワイトハウス大統領府、フロンティア AI モデルの早期評価を目指す
Wallarm が AI 制御プラットフォームを発表、ランタイムの可視化とエンタープライズ AI の強制を実現
サプライチェーン全体にわたるサイバー攻撃を賢く阻止する方法
研究によると、フィッシングがダークウェブを追い越し、盗難の主な発生源となっている...
新しい研究は、AI のオンライン化が加速するにつれて増大するデジタルトラスト危機を浮き彫りにしています...
ホワイトハウス大統領府、フロンティア AI モデルの早期評価を目指す
Wallarm が AI 制御プラットフォームを発表、ランタイムの可視化と強制を実現...
2026 デジタル リスク レポート [アウトテイク]
2026 クラウド セキュリティ レポート [チェック P]

軟膏]
2026 年 Web アプリケーション セキュリティ レポート [フォーティネット]
2026 インサイダーリスクレポート [Gurucul]
KPMG 2026 サイバーセキュリティ報告書は、人間以外のアイデンティティを CISO の重大な問題として挙げています
ID 関連の侵害が企業の 71% に影響、ソフォスの調査で判明
AI セキュリティ検出スタックが見逃す 3 つの即時注入パターン
サイバー攻撃に対して最も脆弱な国のリスト
クラウドセキュリティ関連のデータ侵害トップ 5!
PCI コンプライアンスの間違いトップ 5 とその回避方法
サプライチェーン全体にわたるサイバー攻撃を賢く阻止する方法
研究によると、フィッシングがダークウェブを追い越し、盗難の主な発生源となっている...
新しい研究は、AI のオンライン化が加速するにつれて増大するデジタルトラスト危機を浮き彫りにしています...

## Original Extract

AI is evolving at a rapid pace, and the uptake of Generative AI (GenAI) is revolutionising the way humans interact and leverage this technology. GenAI is

Facebook
Linkedin
X
News
Cyber Attack
News
All Cyber Attack Data Breach Identity Fraud Insider Threat Malware Phishing Quantum Ransomware Social Engineering Supply Chain Security Vulnerability
Cyber Attack
US Government to use Anthropic Mythos to launch Cyber Attacks
Carnival Corporation Data Breach Leads June Wave of Account-Compromise Incidents
How Server Farms can shield customers from Malware Attacks
Carnival Data Breach Potentially Impacts 6 Million Cruisers
The Meta AI Instagram Hack Wasn’t About Authentication. It Was About Authorization.
When attackers hijacked Instagram accounts early June by tricking Meta’s AI support chatbot, most of the coverage focused on the breach itself. But this incident is a great illustration of a broader and more critical problem: the security industry has invested heavily in controlling what AI says, while largely ignoring what AI is authorized to do.
Meta’s bot verified nothing about who was asking. It just helpfully did what it was told to do — up to and including sending the attacker a confirmation code to make sure a new email address was valid. Until we start applying more mature authorization frameworks to AI agents, we’ll have more incidents like this.
The attack itself was straightforward. The attacker spoofed the location of the victim using a VPN, which circumvented certain protections that would have triggered if the attacker’s location was far from the victim’s. The attacker then asked an experimental Meta chatbot to add a new email address to the account. The chatbot emailed verification codes to confirm the new address was valid. It was trying to be helpful! The attacker verified the new email address, was presented with an opportunity to reset the password, and thus gained control of the account.
Most attacks are not one simple hole that can be patched. They string together vulnerabilities to escalate privileges or take over valuable accounts. Based on the attack details that have been publicly shared from this incident, the failures in this vulnerability chain included: relying on IP location to determine if additional security measures are taken; allowing a chatbot to modify a user’s primary email; requiring verification codes only from the new email address and not the old; and treating those verification codes as enough to allow for a password reset, which the chatbot facilitated. Guardrails around any of these would have stopped this version of the attack.
Authentication vs. Authorization — and Why It Matters for AI
Authentication is who someone is. Authorization is what they can do. Authentication is a comparatively better understood issue with AI agents, but authorization decisions reach deep into the bowels of applications and are usually business-model specific. They were often designed either for software designed by humans or slow-moving humans. AI agents combine the speed of software with the innovation of humans, finding edge cases and holes at scale.
Even with perfect authentication, the deeper failure in the Meta incident is that the agent was authorized to perform account-takeover-equivalent actions. And that’s the part the industry is underinvesting in.
Why AI Projects Are Especially Prone to This
Stapling an AI chatbot into a support system didn’t introduce a new class of vulnerability. But it makes such holes more likely to exist, because the efforts to make an AI project successful biases systems toward over-permissioning.
The larger problem is that we’re exposing services, functionality, and APIs to AI agents without properly addressing the actual helpfulness of them, nor how attackers can leverage them to find and exploit existing holes. In this case, Meta wanted the chatbot to be helpful and useful, which requires access. But they gave too much access.
This pattern is already showing up elsewhere. In 2024, an AI agent was tricked by users into sending $47,000 in crypto even though it was explicitly instructed not to. A Lenovo chatbot was manipulated into exposing session cookies based on a crafted product query — not quite direct account access, but the same underlying goal.
These aren’t isolated flukes. They’re early signals of a broader problem that will only grow as agents are granted more authority. If this attack pattern were applied to a healthcare application, the account takeover motion would be similar, but the consequences would be dire — private health data shipped off in violation of privacy rules, for example.
This incident is not a common problem – yet. But based on the trajectory of AI rollout, the number of incidents seems certain to grow as more agents are given more authority.
There are a few authorization controls that should be non-negotiable before an agent goes live. First, limit agents to the bare minimum that is useful, and expand the scope cautiously. Second, use both manual and automated evaluations to determine the true bounds of what an agent can do. Third, evaluate the access provided and determine what needs human review and how that review should take place.
Old Mitigations, New Difficulty
Security is multi-faceted and always has been. The same mitigations that worked in the past such as defense-in-depth, principle of least privilege, and human oversight, work here too. They just need to be applied systematically. The principles are old and well-understood. What’s new and hard is applying them to systems that are non-deterministic, probed at scale, and with surprisingly broad access.
The industry is pretty focused on keeping AI from saying bad things. That’s fine, as long as we don’t completely overlook whether AI should be allowed to do what it’s trying to do. Until we start holding ourselves to more mature authorization frameworks, we’ll continue to have more incidents like this.
Dan Moore is the senior director of CIAM strategy and identity standards at FusionAuth, where he drives B2C product direction and serves as a primary spokesperson on authentication, developer identity, and the emerging challenge of securing AI agents. With over 25 years of software engineering experience — including stints as a CTO, AWS certification instructor, and engineering manager — he is a regular speaker at identity and security conferences and a sought-after podcast guest on authentication and developer security. He is the author of Letters to a New Developer (Apress) and a contributor to 97 Things Every Cloud Engineer Should Know (O’Reilly).
New Research Highlights Growing Digital Trust Crisis as AI Accelerates Online Threats
White House EO Seeks Early Evaluation of Frontier AI Models
Wallarm Launches AI Control Platform Bringing Runtime Visibility and Enforcement to Enterprise AI
How Cyber Attacks across the Supply Chain can be Smartly Thwarted
Research says Phishing overtakes Dark Web as primary source of stolen...
New Research Highlights Growing Digital Trust Crisis as AI Accelerates Online...
White House EO Seeks Early Evaluation of Frontier AI Models
Wallarm Launches AI Control Platform Bringing Runtime Visibility and Enforcement to...
2026 Digital Risk Report [Outtake]
2026 Cloud Security Report [Check Point]
2026 Web Application Security Report [Fortinet]
2026 Insider Risk Report [Gurucul]
KPMG 2026 cybersecurity report names non-human identities as critical CISO problem
Identity-Related Breach Hit 71% of Enterprises, Sophos Survey Finds
Three Prompt Injection Patterns Your AI Security Detection Stack Misses
List of Countries which are most vulnerable to Cyber Attacks
Top 5 Cloud Security related Data Breaches!
Top 5 PCI Compliance Mistakes and How to Avoid Them
How Cyber Attacks across the Supply Chain can be Smartly Thwarted
Research says Phishing overtakes Dark Web as primary source of stolen...
New Research Highlights Growing Digital Trust Crisis as AI Accelerates Online...

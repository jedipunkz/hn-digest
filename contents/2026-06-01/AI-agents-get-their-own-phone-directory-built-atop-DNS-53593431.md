---
source: "https://www.theregister.com/ai-ml/2026/05/28/ai-agents-get-their-own-phone-directory-built-atop-dns/5247539"
hn_url: "https://news.ycombinator.com/item?id=48348513"
title: "AI agents get their own phone directory built atop DNS"
article_title: "AI agents get their own phone directory built atop DNS"
author: "benkan"
captured_at: "2026-06-01T00:26:16Z"
capture_tool: "hn-digest"
hn_id: 48348513
score: 5
comments: 0
posted_at: "2026-05-31T18:50:58Z"
tags:
  - hacker-news
  - translated
---

# AI agents get their own phone directory built atop DNS

- HN: [48348513](https://news.ycombinator.com/item?id=48348513)
- Source: [www.theregister.com](https://www.theregister.com/ai-ml/2026/05/28/ai-agents-get-their-own-phone-directory-built-atop-dns/5247539)
- Score: 5
- Comments: 0
- Posted: 2026-05-31T18:50:58Z

## Translation

タイトル: AI エージェントは DNS 上に構築された独自の電話帳を取得します
説明: Linux Foundation の後援による DNS-AID は、エージェントの発見を容易にすることを約束します

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
AI エージェントは DNS 上に構築された独自の電話帳を取得します
Linux Foundation の後援による DNS-AID は、エージェントの発見を容易にすることを約束します
将来的には、AI エージェントは、ポートを徘徊して調査したり、設定されたリソースを確認したりする代わりに、ドメイン ネーム システム (DNS) を使用して互いを見つけることができるようになります。
その未来は、既存のインターネット インフラストラクチャを使用してエージェント間の検出を容易にすることを目的としたオープン ソース プロジェクトである DNS for AI Discovery ( DNS-AID ) から始まります。このシステムは、競合上のチョークポイントになる可能性のあるさらに別のレジストリの作成を回避するために、DNS 上に構築されています。
「エージェント接続への現在のアプローチは細分化されており、多くの場合、脆弱でハードコードされた構成に依存しています」と、DNS-AID プロジェクトのメンテナーである Ingmar Van Glabbeek 氏は声明で述べています。 「DNS-AID により、私たちは AI の『ウェブネイティブ』モデルに向かって進んでいます。既存の DNS 階層を利用することで、開発者は、何十年もインターネットを移動するために使用してきたのと同じ信頼性と遍在性を備えたエージェントを公開および検出できるようになります。」
DNS は、ドメイン名解決を超えたさまざまな機能をすでに提供しています。たとえば、Web サイトは、DNARC TXT エントリを介して DMARC、SPF、および DKIM レコードを公開します。そして最近では、サービス バインディング ( SVCB ) と HTTPS RR (HTTPS リソース レコード) が採用され、クライアントがサービスと関連パラメーターを簡単に発見できるようになりました。
DNS-AID は、SVCB (フォールバックとして TXT を使用) を利用し、オプションで DNS Security Extensions (DNSSEC) および DNS-Based Authentication of Named Entities (DANE) TLSA レコードを利用します。これらは、仲介エンティティを介さずに接続する方法をエージェントに提供します。

最後のインフラストラクチャ、または優先プロトコル。 DNS-AID は、MCP、A2A、HTTPS、および SVCB および ALPN 経由でアドレス指定可能なものをすべてサポートします。
Snowflake、AWS Graviton CPU と AI アクセラレータで 60 億ドルを消費
FAA、スペースXのスターシップ、打ち上げ事故続発で運航停止に
マルウェア開発者がクロード ユーザーの秘密を盗もうとし、npm スロップを書き込み、独自の GitHub プライベート トークンを漏洩
Argonne は予備のスーパーコンピューティングを活用してプライベート AI 推論サービスを構築します
このシステムでは、エージェントを名前、機能、およびドメインで検索できます。
Linux Foundation は、当初 Infoblox によって開発されたプロジェクトに対してベンダー中立のガバナンスを約束します。
Cloudflareの最高技術責任者（CTO）であるデーン・クネヒト氏は声明で、「インターネットは数十年前にDNSによって発見の問題をすでに解決していた。高速で世界規模に拡張でき、地球上のすべてのネットワークがそれを理解している」と述べた。 「この実証済みのアーキテクチャをエージェント Web に拡張することで、DNS-AID は自律システムが安全かつ効率的に動作するために必要な基本的なルーティング層を提供します。」
開始するにはさまざまな方法がありますが、最も簡単な方法は、dns-aid をインストールして、コマンド dns-aid init を発行することです。 Python SDK はすでに存在しており、おそらく他の言語にリファレンス実装が登場するのもそう長くはかからないでしょう。
セットアップ プロセスには、エージェントの SVCB レコードをサイトの DNS ゾーンに公開し、出所の検証として DNSSEC を使用してゾーンに署名することが含まれます。その後、エージェント レコード _{agent-name}._{protocol}._agents.{your-domain} を解決しようとすると、関連する詳細が返されます。これは、エージェントが相互に接続する前に、DNSSEC、オプションの JWS 署名、および DANE ポリシーによって検証できます。
現在、AWS Route 53、Azure DNS、Cloudflare、Google Cloud DNS、Infoblox NIOS および UDDI、NS1、およびその他のスタンドアを含む、いくつかの DNS プロバイダーが DNS-AID サポートを提供しています。

rds 準拠の DNS サービス (RFC 2136 DDNS)。ローカルで実験を行う開発者には、Docker BIND9 プレイグラウンドを使用するオプションがあります。
世界的なコンサルティング会社マッキンゼーは、代理店間の商取引はいつか意味のあるものになる可能性があると主張し、それを「3兆ドルから5兆ドルの経済機会が見込まれる」と呼んでいる。マッキンゼーが 1980 年代に予測した 2000 年の携帯電話市場の規模は約 100 倍も外れていたことに注目します。 ®
AI + ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
Postgres における AI とデータ主権: データセンターのエネルギー危機への答え
10億人のAIエージェントが電力網に乗り込む
国立宇宙センターでのロケット展示が NASA SLS の意図せぬ印象を引き出す
システム
EUのデジタル主権ブーブーは、このプロジェクトに起こった最高の出来事かもしれない
DIY か、死ぬか。 CIAに買わせないでください
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアモデルのエナジードリンク
AI + ML
Googleは最近AIの機能化に本格的に取り組んでいる
セキュリティ
Anthropic、Mythos クラスのモデルを一般公開へ
セキュリティ
レドモンドが警察に通報、マイクロソフトに「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う
AI + ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーション

これらは API 主導で相互接続されており、多くの場合は過剰な許可が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
今週のケトルでは、スチーム デッキがハードウェア価格の将来にとって炭鉱のカナリアであるかどうか、そして NASA の月ミッションに対する Blue Origin の爆発の影響について考えます。
Project Headroom は多額の費用も節約できる
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
今週のケトルでは、スチーム デッキがハードウェア価格の将来にとって炭鉱のカナリアであるかどうか、そして NASA の月ミッションに対する Blue Origin の爆発の影響について考えます。
Project Headroom は多額の費用も節約できる
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
フロンティアモデルのエナジードリンク
ロケットの爆発とハードウェアの価格の高騰により、ひどい新常態が生まれる
今週は

n The Kettle、私たちはハードウェア価格の将来にとって、スチームデッキが炭鉱のカナリアであるかどうか、そしてNASAの月ミッションに対するBlue Originの爆発の影響について熟考します。
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
Project Headroom は多額の費用も節約できる
ウィキメディアの解雇後、ウィキペディア編集者がストライキとバナー破壊行為を計画
コミュニティから要求された多くの修正とモデレーションツールを担当するチームを解散させた後、財団が反乱を引き起こす
国立宇宙センターでのロケット展示が NASA SLS の意図せぬ印象を引き出す
5、4、3、2、1... pfft
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアモデルのエナジードリンク
単独の攻撃者が、人気のある OpenSearch、Elasticsearch ライブラリを模倣した 14 個の悪意のある npm パッケージを公開
そしてマイクロソフトはそれらすべてを潰した
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

DNS-AID, under the auspices of the Linux Foundation, promises easier agent discovery

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI agents get their own phone directory built atop DNS
DNS-AID, under the auspices of the Linux Foundation, promises easier agent discovery
In the future, AI agents will be able to find one another using the Domain Name System (DNS), instead of crawling about and probing ports or checking configured resources.
That future begins now with DNS for AI Discovery ( DNS-AID ), an open source project intended to facilitate agent-to-agent discovery using existing internet infrastructure. The system has been built atop DNS to avoid the creation of yet another registry that has the potential to become a competitive chokepoint.
"Current approaches to agent connectivity are fragmented and often rely on fragile, hardcoded configurations,” said Ingmar Van Glabbeek, project maintainer for DNS-AID, in a statement . "With DNS-AID, we are moving toward a 'web-native' model for AI. By utilizing the existing DNS hierarchy, we enable developers to publish and discover agents with the same reliability and ubiquity that we’ve used to navigate the internet for decades."
DNS already provides various capabilities beyond domain name resolution. For example, websites expose their DMARC, SPF, and DKIM records via DNS TXT entries. And more recently, Service Binding ( SVCB ) and HTTPS RR (HTTPS Resource Records) were adopted to make it easier for clients to discover services and associated parameters.
DNS-AID utilizes SVCB (with TXT as a fallback) and optionally DNS Security Extensions (DNSSEC) and DNS-Based Authentication of Named Entities (DANE) TLSA records. These provide agents with a way to connect without a mediating entity, additional infrastructure, or a preferred protocol. DNS-AID supports MCP, A2A, HTTPS, and anything addressable via SVCB and ALPN.
Snowflake to burn $6B on AWS Graviton CPUs and AI accelerators
FAA grounds SpaceX’s Starship after another launch mishap
Malware dev tries to steal Claude users' secrets, writes npm slop, leaks own GitHub private token
Argonne flexes spare supercompute to build private AI inference service
The system allows agents to be searched by name, by function, and by domain.
The Linux Foundation promises vendor-neutral governance for the project, which was initially developed by Infoblox.
"The Internet already solved the discovery problem decades ago with DNS – it's fast, it scales globally, and every network on earth understands it," said Dane Knecht, CTO of Cloudflare, in a statement. "By extending this proven architecture to the agentic web, DNS-AID provides the foundational routing layer that autonomous systems need to operate safely and efficiently."
There are various ways to get started , the simplest of which is to install dns-aid and then issue the command dns-aid init. There is already a Python SDK , and presumably it won't be long before other languages have a reference implementation.
The setup process involves publishing an SVCB record for your agent to your site's DNS zone and signing the zone with DNSSEC as a verification of provenance. Thereafter, attempts to resolve the agent record _{agent-name}._{protocol}._agents.{your-domain} will return the relevant details, which can then be verified via DNSSEC, optional JWS signatures, and DANE policy before agents connect to one another.
Several DNS providers currently offer DNS-AID support, including AWS Route 53, Azure DNS, Cloudflare, Google Cloud DNS, Infoblox NIOS and UDDI, NS1, and any standards-compliant DNS service (RFC 2136 DDNS). Developers experimenting locally have the option to use a Docker BIND9 playground.
Global consultancy McKinsey argues that agent-to-agent commerce could be meaningful someday, calling it a "projected $3 trillion to $5 trillion economic opportunity." We note that McKinsey's 1980s prediction about the size of the mobile phone market in 2000 was off by about 100x. ®
AI + ML
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
AI and data sovereignty in Postgres: An answer to the datacenter energy crisis
A billion AI agents walk into a power grid
Rocket exhibit at National Space Centre pulls off unintentional NASA SLS impression
Systems
EU's digital sovereignty boo-boo may be the best thing to ever happen to the project
DIY or die. Just don't let the CIA buy it
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
AI + ML
Google has seriously leaned into AI enshittification lately
Security
Anthropic to release Mythos-class models to the public
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
AI + ML
Netflix wiz creates app to slash AI bills, then open sources it
Security
Troops’ phones gave away location data to foreign adversaries
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Project Headroom could save you big money, too
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Project Headroom could save you big money, too
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
The energy drink of frontier models
Exploding rockets and exploding hardware prices make for a lousy new normal
This week on The Kettle, we mull over whether the Steam Deck is a canary in the coal mine for the future of hardware prices, and the effect of Blue Origin's blowout on NASA's Moon missions
Netflix wiz creates app to slash AI bills, then open sources it
Project Headroom could save you big money, too
Wikipedia editors plot strike and banner sabotage after Wikimedia layoffs
Foundation sparks revolt after disbanding team responsible for many community-requested fixes and moderation tools
Rocket exhibit at National Space Centre pulls off unintentional NASA SLS impression
5, 4, 3, 2, 1... pfft
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
Lone attacker published 14 malicious npm packages mimicking popular OpenSearch, Elasticsearch libraries
And then Microsoft busted them all
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

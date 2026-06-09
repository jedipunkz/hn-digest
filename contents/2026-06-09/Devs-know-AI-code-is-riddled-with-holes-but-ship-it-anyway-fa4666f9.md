---
source: "https://www.theregister.com/devops/2026/06/09/devs-know-ai-code-is-riddled-with-holes-but-ship-it-anyway/5252824"
hn_url: "https://news.ycombinator.com/item?id=48466519"
title: "Devs know AI code is riddled with holes, but ship it anyway"
article_title: "Devs know AI code is riddled with holes, but ship it anyway"
author: "speckx"
captured_at: "2026-06-09T21:39:19Z"
capture_tool: "hn-digest"
hn_id: 48466519
score: 18
comments: 15
posted_at: "2026-06-09T19:37:35Z"
tags:
  - hacker-news
  - translated
---

# Devs know AI code is riddled with holes, but ship it anyway

- HN: [48466519](https://news.ycombinator.com/item?id=48466519)
- Source: [www.theregister.com](https://www.theregister.com/devops/2026/06/09/devs-know-ai-code-is-riddled-with-holes-but-ship-it-anyway/5252824)
- Score: 18
- Comments: 15
- Posted: 2026-06-09T19:37:35Z

## Translation

タイトル: 開発者は AI コードが穴だらけであることを知っているが、それでも出荷する
説明: 5 つの組織のうち 4 つが脆弱なアプリによる侵害を認めたため、導入へのプレッシャーがセキュリティよりも優先

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
開発者は AI コードが穴だらけであることを知っているが、それでも出荷する
組織の 5 つ中 4 が脆弱なアプリによる侵害を認め、導入への圧力がセキュリティよりも優先
AppSec ビジネスの Checkmarx による調査によると、開発者の 70% は AI で生成されたコードにはさらに脆弱性があると信じており、30% は脆弱なコードを故意に本番環境に出荷しています。
このレポートは、2,350 人の世界の開発者、CISO、AppSec マネージャーからの回答に基づいており、2023 年以降の同様の年次調査に従っています。回答者の数は、今年は昨年より 54 パーセント増加しており、サンプル サイズの増加は、やや驚くべき統計を説明している可能性があります。報告されている AI 生成の実稼働コードの割合は、54 パーセントから 49 パーセントにわずかに減少しましたが、これは依然として高い数字です。
レポートによると、本番アプリケーションもオープンソース基盤上に構築されており、コードの 59% を占めています。これらは自己報告による推定値ですが、多くのオープン ソース コードは、node_modules またはその他のライブラリの場所に埋め込まれており、AI によって発見された脆弱性への対応に苦戦している管理者や、npm や PyPI などの一般的なパッケージ リポジトリに悪意のあるパッケージが密かに持ち込まれているため、常に安全であるとは限りません。
その結果、問題は脆弱なコードだけでなく資格情報を盗むマルウェアにまで広がり、ソフトウェア開発のリスクはこれまで以上に高まっています。それでも Checkmarx の調査では諦めているようで、回答者の 93 パーセントが脆弱なアプリケーションの結果として 1 つ以上のセキュリティ違反を報告しています (ただし、昨年の数字は 98 パーセントでした)。理由としては、dに対する圧力などが挙げられている。

迅速な導入、脆弱性の修正が困難すぎること、そして部分を拾うために他のコントロールに依存していること。
「リスクは正常化されている」とCheckmarxはレポートで述べている。
AI によって生成されたコードのセキュリティは、特にこれらの回答者の間では、記述された内容の約 50% を占めていることから、ホットなトピックとなっています。 70% が「AI が生成したコードの脆弱性が大幅に増加している」と報告しており、セキュリティ問題を見落とす点では AI が人間よりもさらに悪いことが示唆されています。
複雑な状況です。 AI は既存のコード、主に公開コードに基づいてトレーニングされますが、このコードには複製される可能性のある脆弱性が存在します。 AI の波は、脆弱性を分析して修復するための新しいツールも提供しました。
セントラルフロリダ大学とパレスチナのビルゼイト大学のコンピューター科学者らによる昨年の研究では、さまざまなプログラミング言語 (Java、Python、C、C++) と LLM の間でコードのセキュリティがどのように異なるのか、またどの脆弱性が最も蔓延しているのかが調査されました。研究者らは、LLM が急速に進化しており、この研究が「タイムスタンプ付きの見解」であることを認めていますが、C コードがセキュリティ上の問題を最も多く抱えている傾向があり、Python が最も少ないという大きなばらつきが見られました。問題の 1 つは、LLM が「最新の言語とコンパイラの機能を十分に活用しておらず、より安全な代替手段よりも時代遅れの手法を好むことが多い」ことです。考えられる理由は、トレーニング データにそのような慣行が蔓延していることです。
混乱を引き起こすコンピューター ワームを構築するのに Mythos やゼロデイは必要ありません – 無料のオープンソース モデルは問題なく動作します
GitHub がワーム感染の疑いで 70 以上の Microsoft リポジトリを破壊し、CI/CD パイプラインを破壊
Anthropic、Mythos クラスのモデルを一般公開へ
AI主導の急速な開発はセキュリティを達成不可能にする、とVeracodeが警告
重要な質問

重要なのは、開発者が古いスタイルの静的分析や新しい AI 主導のオプションなどのツールを使用して脆弱性を排除できるかどうかです。 Checkmarx 氏によると、それは可能ですが、できない場合が多いそうです。
「ツールは機能しますが、組織はこれをプロセスに変換する能力が不足しています」と同社は報告しています。 Veracode も報告しているように、AI 支援によって開発のペースが加速しており、セキュリティ対策が追いついていない状況です。
Checkmarx の研究者は、「AI コードの量は脆弱なコードの展開と直接相関し、それは侵害の頻度と直接相関します。」と述べています。具体的には、「コードの 81 ～ 100 パーセントが AI によって生成された組織は、導入率が 1 ～ 20 パーセントの組織の 3.4 倍の割合で脆弱なコードを出荷します。これは、開発を加速するために支払うべき高い代償です。」 ®
個人的な技術
性生活が終わったらスティーブ・ジョブズを責めてもいい
経済学者らはAT&T独占時代のデータから「iPhoneと出生率の間に大きな因果関係」の兆候を発見
人間が紡ぐ、飼いならされたより安全な神話の寓話
会社はデータ保持ポリシーも変更します
ZTE、TNBエネルギー移行カンファレンスで統合AI、コネクティビティ、デジタルユーティリティテクノロジーをデモンストレーション
パートナーコンテンツ: 高度な AI およびスマート インフラストラクチャ ソリューションを通じてマレーシアの送電網の近代化とエネルギー転換を推進
攻撃キットのオープンソース化に伴い、Miasma ワームが GitHub に侵入
まるで心配するほどのパッケージ中毒がなかったかのように
OSプラットフォーム
はい！それは本当です！ Windows 11 はエージェント プラットフォームです
これまでずっとそうだったが、Microsoft はそれに気づいていなかった
MIT のボフィンがエレクトロ スプレー ノズルをクリーンルームから取り出して 3D プリンターに導入
ミリ未満の 3 層のサイエンス ジュースを噴射するにはコストがかかるなどと誰が言ったでしょうか。
セキュリティ
すべてのパスワードは Active Directory の説明フィールドに保存されました
公共部門
英国政府

Stripe を破棄し、支払いがオランダになる
セキュリティ
GitHub がワーム感染の疑いで 70 以上の Microsoft リポジトリを破壊し、CI/CD パイプラインを破壊
AI と ML
従量課金が定着する中、怒った開発者らはGitHub Copilotから逃げることを誓う
パーソナルテクノロジー
カリフォルニア州、3Dプリントによるゴーストガンのアルゴリズムによる死亡を宣言する法案を可決
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
ボラティリティを乗り越えて成功する: 不確実な市場における永遠の優位性
消費ベースの運用モデルがどのように柔軟性を提供し、効率を向上させ、インフラストラクチャ投資に予測可能性をもたらすかを学びましょう。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
Agentic AI 時代のゼロトラスト
ほとんどの組織が依存している ID およびアクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
Agentic AI 時代のゼロトラスト
アイデンティティー

ほとんどの組織が依存している ND アクセス モデルは、人間以外の ID が独立して動作するのではなく、人間のユーザー向けに構築されています。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで真の ROI を実現する方法を学びましょう。
まるで心配するほどのパッケージ中毒がなかったかのように
ミリ未満の 3 層のサイエンス ジュースを噴射するにはコストがかかるなどと誰が言ったでしょうか。
iBiz は AI 競争には勝てないかもしれないが、アナリストらは iBiz は人々が実際に使用する可能性のある機能に焦点を当てていると述べている
グラフデータベース業界は、オンプレミスのエアギャップインテルスタックが政府にノーキルスイッチオプションを提供すると発表
Document Foundationは、新しく発足したEuro-OfficeがMicrosoftのOOXML文書形式をデフォルトとして使用することでデジタル主権を損なっていると非難
iBiz は AI 競争には勝てないかもしれないが、アナリストらは iBiz は人々が実際に使用する可能性のある機能に焦点を当てていると述べている
グラフデータベース業界は、オンプレミスのエアギャップインテルスタックが政府にノーキルスイッチオプションを提供すると発表
Document Foundationは、新しく発足したEuro-OfficeがMicrosoftのOOXML文書形式をデフォルトとして使用することでデジタル主権を損なっていると非難
組織の 5 つ中 4 が脆弱なアプリによる侵害を認め、導入への圧力がセキュリティよりも優先
暗号化されたメッセージングアプリが、デバイスレベルのチェックが検閲に再利用される可能性があると警告
性生活が終わったらスティーブ・ジョブズを責めてもいい
経済学者らはAT&T独占時代のデータから「iPhoneと出生率の間に大きな因果関係」の兆候を発見
人間が紡ぐ、飼いならされたより安全な神話の寓話
会社はデータ保持ポリシーも変更します
攻撃キットのオープンソース化に伴い、Miasma ワームが GitHub に侵入
まるで心配するほどのパッケージ中毒がなかったかのように
MIT のボフィンがエレクトロ スプレー ノズルをクリーンルームから取り出して 3D プリンターに導入
サブミリメートルの 3 層サイエンス ジュースが高価である必要があるなどと誰が言ったでしょうか。

ああ、潮吹き？
Apple の iOS 27 は、漏洩したパスワードをすべてエージェント的に処理し、ワンタップでパスワードを変更できると約束
iBiz は AI 競争には勝てないかもしれないが、アナリストらは iBiz は人々が実際に使用する可能性のある機能に焦点を当てていると述べている
Neo4j が GraphAware 買収で Palantir の代替案をプロット
グラフデータベース業界は、オンプレミスのエアギャップインテルスタックが政府にノーキルスイッチオプションを提供すると発表
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

Pressure to deploy wins out over security as four in five orgs confess to breaches from vulnerable apps

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Devs know AI code is riddled with holes, but ship it anyway
Pressure to deploy wins out over security as four in five orgs confess to breaches from vulnerable apps
Research by AppSec biz Checkmarx finds that 70 percent of developers believe AI-generated code has more vulnerabilities, and 30 percent knowingly ship vulnerable code into production.
The report is based on responses from 2,350 global developers, CISOs, and AppSec managers, and follows similar annual surveys since 2023. The number of respondents is 54 percent higher this year than last, and the increased sample size may account for a somewhat surprising statistic: the reported proportion of AI-generated production code has slightly declined, from 54 percent to 49 percent, though this is still a high figure.
Production applications are also built on an open source foundation, according to the report, accounting for 59 percent of the code. These are self-reported estimates, but a lot of open source code is buried in node_modules or other library locations and it is not always secure, whether because of hard-pressed maintainers struggling to keep up with AI-discovered vulnerabilities, or malicious packages smuggled into popular package repositories such as npm and PyPI.
The consequence is that software development is riskier than ever, with issues extending beyond vulnerable code to credential-stealing malware, yet the Checkmarx survey appears to show resignation, with 93 percent of respondents reporting one or more security breaches as a result of vulnerable applications – though last year the figure was 98 percent. Reasons given include pressure to deploy quickly, vulnerabilities being too difficult to fix, and reliance on other controls to pick up the pieces.
"Risk is normalized," says Checkmarx in its report.
The security of AI-generated code is a hot topic, particularly since, among these respondents, it accounts for around 50 percent of what is written. 70 percent report "significantly more vulnerabilities with AI-generated code," suggesting that AI is even worse than humans when it comes to overlooking security issues.
It is a complex situation. AI is trained on existing code, primarily public code, which has its share of vulnerabilities that may then be replicated. The AI wave has also delivered new tools for analyzing and remediating vulnerabilities.
A study last year by computer scientists from the University of Central Florida and Birzeit University in Palestine looked at how code security varied between different programming languages (Java, Python, C, and C++) and LLMs, and which vulnerabilities are most prevalent. The findings showed significant variations, with C code tending to have the most security issues, and Python the fewest, though the researchers acknowledge that LLMs are evolving rapidly and that the research is a "time-stamped view." One of the issues is that LLMs "underutilize modern language and compiler features, often favoring outdated practices over more secure alternatives." The likely reason is the prevalence of such practices in the training data.
Nobody needs Mythos or 0-days to build a chaos-causing computer worm – free open source models work just fine
GitHub nukes 70+ Microsoft repos, breaks CI/CD pipelines, following suspected worm infections
Anthropic to release Mythos-class models to the public
Rapid AI-driven development makes security unattainable, warns Veracode
A key question is whether developers can eliminate vulnerabilities using tooling, including old-style static analysis and newer AI-driven options. According to Checkmarx, they could but often do not.
"The tools do the work, but organizations lack in translating this into process," the company reports. As Veracode has also reported , AI assistance is driving up the pace of development and security practices cannot keep up.
The Checkmarx researchers state: "AI code volume correlates directly with vulnerable code deployment, which correlates directly with breach frequency." Specifically, "organizations where 81-100 percent of code is AI-generated ship vulnerable code at 3.4x the rate of those at 1-20 percent adoption" – a high price to pay for accelerated development. ®
personal tech
If your sex life is dead, you can blame Steve Jobs
Economists find signs of a ‘large and causal relationship between iPhones and fertility' in AT&T exclusivity-era data
Anthropic spins a Fable of a tamer, safer Mythos
Company also changes data retention policy
ZTE Demonstrates Integrated AI, Connectivity and Digital Utility Technologies at TNB Energy Transition Conference
PARTNER CONTENT: Driving Grid Modernization and Energy Transition in Malaysia Through Advanced AI and Smart Infrastructure Solutions
Miasma worms its way onto GitHub as attack kit goes open source
As if there weren't enough package poisonings to worry about
os platforms
Yes! It’s true! Windows 11 is an agentic platform
It always has been, but Microsoft didn’t realize it
MIT boffins take electrospray nozzles out of the cleanroom, into the 3D printer
Who said sub-millimeter, three-layer science juice had to be expensive to squirt?
SECURITY
All the passwords were stored in Active Directory description fields
public sector
GOV.UK goes Dutch on payments as it dumps Stripe
security
GitHub nukes 70+ Microsoft repos, breaks CI/CD pipelines, following suspected worm infections
AI and ML
Angry devs vow to flee GitHub Copilot as metered billing takes hold
Personal tech
California passes bill declaring death-by-algorithm to 3D-printed ghost guns
Overcoming the trade-offs in data sovereignty
What does data sovereignty actually mean for your network, which trade-offs are unavoidable? Learn more.
Thriving Through Volatility: The Everpure Advantage in an Uncertain Market
Learn how a consumption-based operating model provides flexibility, improves efficiency, and brings predictability to infrastructure investments.
From Prompt to Exploit: How LLMs Are Changing API Attacks
Modern applications are API-driven, interconnected, and often over-permissioned, making them an ideal target for AI-assisted attacks.
Architecting the Future: Unlocking Enterprise Data Services for Kubernetes
Join us to discover how to eliminate infrastructure silos and establish a standardized, enterprise-grade cloud-native platform.
Catch the Advanced Attacks Microsoft 365 Misses with Behavioral AI Security
Microsoft 365 is the backbone of enterprise communication, and its native security filters out the known and the noisy.
Step into the chaos of a live ransomware breach, test your response skills, and team up with other IT and security pros to outsmart cybercriminals
Virtual Cyber Recovery Simulation
Ransomware attacks aren’t slowing down, and neither are we. Druva’s hit event, Escape Ransomware, is now fully virtual.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Zero Trust for the Agentic AI Era
The identity and access models most organizations rely on were built for human users, not non-human identities operating independently.
Agentic AI at Scale: From Pilot to Production
Join us to learn how to unlock real ROI by driving adoption of AI at scale.
As if there weren't enough package poisonings to worry about
Who said sub-millimeter, three-layer science juice had to be expensive to squirt?
iBiz might not win the AI race, but analysts say it's focusing on features people may actually use
Graph database biz says on-prem, air-gapped intel stack gives governments a no-kill-switch option
The Document Foundation accuses newly launched Euro-Office of undermining digital sovereignty by defaulting to Microsoft's OOXML document format
iBiz might not win the AI race, but analysts say it's focusing on features people may actually use
Graph database biz says on-prem, air-gapped intel stack gives governments a no-kill-switch option
The Document Foundation accuses newly launched Euro-Office of undermining digital sovereignty by defaulting to Microsoft's OOXML document format
Pressure to deploy wins out over security as four in five orgs confess to breaches from vulnerable apps
Encrypted messaging app warns device-level checks could be repurposed for censorship
If your sex life is dead, you can blame Steve Jobs
Economists find signs of a ‘large and causal relationship between iPhones and fertility' in AT&T exclusivity-era data
Anthropic spins a Fable of a tamer, safer Mythos
Company also changes data retention policy
Miasma worms its way onto GitHub as attack kit goes open source
As if there weren't enough package poisonings to worry about
MIT boffins take electrospray nozzles out of the cleanroom, into the 3D printer
Who said sub-millimeter, three-layer science juice had to be expensive to squirt?
Apple’s iOS 27 goes all agentic on compromised passwords, promises to change them with one tap
iBiz might not win the AI race, but analysts say it's focusing on features people may actually use
Neo4j plots Palantir alternative with GraphAware acquisition
Graph database biz says on-prem, air-gapped intel stack gives governments a no-kill-switch option
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

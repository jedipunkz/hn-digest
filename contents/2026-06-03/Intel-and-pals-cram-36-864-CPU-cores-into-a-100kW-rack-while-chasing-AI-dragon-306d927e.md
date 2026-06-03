---
source: "https://www.theregister.com/ai-and-ml/2026/06/02/intel-and-pals-cram-36864-cpu-cores-into-a-100kw-rack-while-chasing-the-agentic-ai-dragon/5249917"
hn_url: "https://news.ycombinator.com/item?id=48373129"
title: "Intel and pals cram 36,864 CPU cores into a 100kW rack while chasing AI dragon"
article_title: "Intel and pals cram 36,864 CPU cores into a 100kW rack while chasing the agentic AI dragon"
author: "Bender"
captured_at: "2026-06-03T00:39:58Z"
capture_tool: "hn-digest"
hn_id: 48373129
score: 3
comments: 0
posted_at: "2026-06-02T17:14:39Z"
tags:
  - hacker-news
  - translated
---

# Intel and pals cram 36,864 CPU cores into a 100kW rack while chasing AI dragon

- HN: [48373129](https://news.ycombinator.com/item?id=48373129)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/06/02/intel-and-pals-cram-36864-cpu-cores-into-a-100kw-rack-while-chasing-the-agentic-ai-dragon/5249917)
- Score: 3
- Comments: 0
- Posted: 2026-06-02T17:14:39Z

## Translation

タイトル: Intel とその仲間が AI ドラゴンを追いかけながら 100kW ラックに 36,864 個の CPU コアを詰め込む
記事のタイトル: Intel とその仲間は、エージェント AI ドラゴンを追いかけながら 100kW ラックに 36,864 個の CPU コアを詰め込む
説明: 一方、Intel と SambaNova

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Intel とその仲間たちは、エージェント AI ドラゴンを追いかけながら 36,864 個の CPU コアを 100kW ラックに詰め込みます
一方、Intel と SambaNova の分離された推論ブループリントが最初の顧客を獲得
トビアス・マン
トビアス
マン
システムエディター
発行済み
2026年6月2日火曜日 // 10:37 UTC
COMPUTEX 2026 インテルは、Foxconn およびその他のインフラストラクチャープロバイダーと協力して、チップメーカーの Xeon プロセッサーに基づくラックスケールのリファレンス設計を開発しています。
火曜日のインテルの Computex 基調講演中に発表されたこれらのブループリントは、AI エージェントを大規模に実行するための CPU コンピューティング密度を向上させることを目的としています。
AI モデルは主に GPU やその他の AI アクセラレータ上で実行されますが、ツール、ターミナル シェル、コード インタープリタ、その他の API に接続するために使用される OpenClaw などのエージェント ハーネスは依然として CPU 上で実行されます。
Intel CEO の Lip Bu Tan 氏は、「当社の顧客は、実際のエージェント ワークロードを大規模に処理できるように、システム レベルで考えることを求めています」と述べています。
タン氏はステージ上で、これらの青写真の 2 つの例を明らかにしました。 1 つはレイテンシの影響を受けやすいエージェント ワークロードを目的としたもので、もう 1 つは最大密度を目指して設計されたものです。
どちらの設計も、Intel の 128 コア Granite Rapids Xeon 6 または 288 コア Clearwater Forest Xeon 6+ プロセッサを最大 128 個サポートし、合計 16,384 個の P コアと 36,864 個の E コアをサポートし、100kW の電力エンベロープで最大 384 TB の DDR5 をサポートします。
このリファレンス デザインは、Nvidia が 256 個の 88 コア Vera CPU を搭載した同様のラックスケール CPU プラットフォームを発表してからわずか数か月後に発表されました。
Arm は、新しい AGI CPU に基づくエージェント ワークロード用の 2 つのラックスケール リファレンス デザインにも取り組んでいます。これは、8,160 コアを備えた 36 kW 空冷システムと 200 kW 液冷システムです。

45,696 コアを搭載した d ラック。
Tan は、これらのリファレンス設計に基づくシステムが ODM および OEM パートナーから広く入手可能になることを期待しています。
エージェントがコアを切望しているのであれば、Intel の新しい Clearwater Xeon 6+ は彼らの渇きを潤してくれるかもしれません
Broadcom のカスタム ASIC 事業により、韓国の FuriosaAI が帝国に加わりました
GPU 大手 Nvidia も世界有数の CPU サプライヤーになる軌道に乗っていると CFO が語る
AMD、AIに興味のある企業向けに新しいスロット可能なGPUを発売
同社は、エージェント AI ワークロードと並んで、新しく立ち上げた推論クラウド プロバイダーの Vector Core Compute がこのプラットフォームを最初に展開する企業の 1 社となり、Togetter.AI がその最初の商用顧客であることも明らかにしました。
このアプローチは、インテルがパートナーの SambaNova と共同開発した以前の細分化された AI ブループリントに基づいています。このアーキテクチャは、Nvidia GPU への負荷の高いプリフィル操作の計算を分離すると同時に、帯域幅を大量に消費するデコード操作に SambaNova の AI アクセラレータを使用して、ユーザーごとのトークン出力を 2 ～ 3 倍に高めます。
聞き覚えがあるかもしれませんが、Nvidia が Groq の LPU で行っていること、または AWS が Trainium と Cerebra のウェハスケール AI アクセラレータで行っていることと似ています。®
サイバー犯罪
「愚かな」犯罪者が「ランサムウェアクラブの第一規則」を破る
ロシアや他のCIS諸国の誰にも感染させない
内容はSalesforceの「頭のない」賭けに対する一撃だ
Headless 360 のエンタープライズ コンテンツ レイヤーが不足していた CRM の巨人は買い物に出かけた
Postgres における AI とデータ主権: データセンターのエネルギー危機への答え
10億人のAIエージェントが電力網に乗り込む
トランプ大統領のAI「E-(I)-O」により、連邦当局が勝者と敗者を決める可能性がある
政府は「信頼できるパートナー」へのアクセスに関して発言権を持ち、それが政策専門家を懸念させている
SaaS
AWSはエンタープライズ需要がゼロにもかかわらず、イーロン・マスク氏のGrokをBedrockに押し込むと報じられている
フロンティアのエナジードリンク

モデル
CiscoはMythosを称賛しているが、このモデルでどれだけのバグが発見されたかについては言及していない
一方、Anthropic は Project Glasswing に 150 人のパートナーを追加
AI + ML
Netflix のウィズが AI 費用を削減するアプリを作成し、それをオープンソース化
セキュリティ
レドモンドが警察に通報、マイクロソフトに「屈辱を受けた」と不満を抱く0日ハンターが「骨が砕けるほどの衝撃」を誓う
セキュリティ
軍隊の携帯電話が位置データを外国の敵に提供した
AI + ML
Googleは最近AIの機能化に本格的に取り組んでいる
セキュリティ
Anthropic、Mythos クラスのモデルを一般公開へ
データ主権のトレードオフを克服する
避けられないトレードオフであるデータ主権はネットワークにとって実際に何を意味するのでしょうか?もっと詳しく知る。
プロンプトからエクスプロイトまで: LLM は API 攻撃をどのように変化させているか
最新のアプリケーションは API 主導で相互接続されており、多くの場合、過剰な権限が与えられているため、AI 支援攻撃の理想的な標的となっています。
未来の構築: Kubernetes 向けエンタープライズ データ サービスのロックを解除する
インフラストラクチャのサイロを排除し、標準化されたエンタープライズ グレードのクラウド ネイティブ プラットフォームを確立する方法を発見してください。
Microsoft 365 が見逃す高度な攻撃を Behavioral AI セキュリティで捕捉
Microsoft 365 は企業コミュニケーションのバックボーンであり、そのネイティブ セキュリティが既知のものや騒々しいものを除外します。
ランサムウェア侵害の混乱に足を踏み入れ、対応スキルをテストし、他の IT およびセキュリティの専門家とチームを組んでサイバー犯罪者を出し抜こう
仮想サイバー復旧シミュレーション
ランサムウェア攻撃の勢いは衰えていませんが、私たちも同様です。 Druva の人気イベント「Escape Ransomware」が完全にバーチャルになりました。
大規模なエージェント AI: パイロットから運用まで
AI の導入を大規模に推進することで、真の ROI を実現する方法を学びましょう。
ロシアや他のCIS諸国の誰にも感染させない
エンタープライズコンテンツレイヤーの欠如

Headless 360 について、CRM titan が買い物に行きました
政府は「信頼できるパートナー」へのアクセスに関して発言権を持ち、それが政策専門家を懸念させている
一方、Anthropic は Project Glasswing に 150 人のパートナーを追加
ニューヨーク連銀によると、若い専門家は在宅勤務でも完璧に生産性を発揮できるかもしれないが、彼らの成果の質はそれほど高くないため、企業は彼らを雇用したがらない
ロシアや他のCIS諸国の誰にも感染させない
Headless 360 のエンタープライズ コンテンツ レイヤーが不足していた CRM の巨人は買い物に出かけた
政府は「信頼できるパートナー」へのアクセスに関して発言権を持ち、それが政策専門家を懸念させている
一方、Anthropic は Project Glasswing に 150 人のパートナーを追加
ニューヨーク連銀によると、若い専門家は在宅勤務でも完璧に生産性を発揮できるかもしれないが、彼らの成果の質はそれほど高くないため、企業は彼らを雇用したがらない
「愚かな」犯罪者が「ランサムウェアクラブの第一規則」を破る
ロシアや他のCIS諸国の誰にも感染させない
内容はSalesforceの「頭のない」賭けに対する一撃だ
Headless 360 のエンタープライズ コンテンツ レイヤーが不足していた CRM の巨人は買い物に出かけた
トランプ大統領のAI「E-(I)-O」により、連邦当局が勝者と敗者を決める可能性がある
政府は「信頼できるパートナー」へのアクセスに関して発言権を持ち、それが政策専門家を懸念させている
CiscoはMythosを称賛しているが、このモデルでどれだけのバグが発見されたかについては言及していない
一方、Anthropic は Project Glasswing に 150 人のパートナーを追加
AIではなくリモートワークが若者の就職の可能性を奪っている
ニューヨーク連銀によると、若い専門家は在宅勤務でも完璧に生産性を発揮できるかもしれないが、彼らの成果の質はそれほど高くないため、企業は彼らを雇用したがらない
マーベル、102.4 Tbps スイッチ シリコンで AI ネットワーク争いに参入
高基数、低レイテンシ、低電力は AI データセンターが切望しているものである、とチップメーカーは語る
お問い合わせ
私たちと一緒に宣伝しましょう
誰

私たちは
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

Meanwhile, Intel and SambaNova

Jump to main content
Search
TOPICS
Special Features
All Special Features
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Intel and pals cram 36,864 CPU cores into a 100kW rack while chasing the agentic AI dragon
Meanwhile, Intel and SambaNova's disaggregated inference blueprint lands its first customer
Tobias Mann
Tobias
Mann
Systems editor
Published
tue 2 Jun 2026 // 10:37 UTC
COMPUTEX 2026 Intel is working with Foxconn and other infrastructure providers to develop rack-scale reference designs based on the chipmaker’s Xeon processors.
Announced during Intel’s Computex keynote on Tuesday, these blueprints aim to provide greater CPU compute densities for running AI agents at scale.
While AI models predominantly run on GPUs and other AI accelerators, the agent harnesses, like OpenClaw, which are used to connect them to tools, terminal shells, code interpreters, and other APIs, still run on CPUs.
“Our customers are asking us to think at the system level to help them serve real agentic workloads at scale,” Intel CEO Lip Bu Tan said.
On stage, Tan revealed two examples of these blueprints. One is aimed at latency-sensitive agentic workloads and another designed for maximum density.
Both designs support up to 128 of either Intel’s 128-core Granite Rapids Xeon 6 or 288-core Clearwater Forest Xeon 6+ processors , totaling between 16,384 P-cores and 36,864 E-cores, alongside up to 384 TB of DDR5 in a 100kW power envelope.
The reference designs come just months after Nvidia announced a similar rack-scale CPU platform packing 256 of its 88-core Vera CPUs.
Arm is also working on a pair of rack-scale reference designs for agentic workloads based on its new AGI CPUs: a 36 kW air-cooled system with 8,160 cores and a 200 kW liquid cooled rack with 45,696 cores.
Tan expects systems based on these reference designs to be broadly available from its ODM and OEM partners.
If cores are what agents crave, Intel's new Clearwater Xeon 6+ might just quench their thirst
Broadcom's custom ASIC biz adds South Korea's FuriosaAI to its empire
GPU behemoth Nvidia on track to be world's leading CPU supplier too, says CFO
AMD puts out new slottable GPU for AI-curious enterprises
Alongside agentic AI workloads, the company also revealed that newly launched inference cloud provider Vector Core Compute will be among the first to deploy the platform, and that Together.AI is its first commercial customer.
The approach is based on Intel's earlier disaggregated AI blueprint it co-developed with partner SambaNova. The architecture desegregates compute heavy prefill operations to Nvidia GPUs while using SambaNova’s AI accelerators for bandwidth-intensive decode operations to boost per-user token output by between 2-3x.
If that sounds familiar it’s not dissimilar to what Nvidia is doing with Groq’s LPUs or what AWS is doing with Trainium and Cerebra's waferscale AI accelerators.®
cyber-crime
'Dumbass' criminal breaks the 'first rule of ransomware club'
You don't infect anyone in Russia or other CIS countries
Contentful is a shot in the arm for Salesforce's 'headless' bet
Lacking an enterprise content layer for Headless 360, CRM titan went shopping
AI and data sovereignty in Postgres: An answer to the datacenter energy crisis
A billion AI agents walk into a power grid
Trump's AI E-(I)-O could let feds pick winners and losers
Government gets a say in 'trusted partner' access, and that worries policy experts
SaaS
AWS reportedly to tuck Elon Musk's Grok into Bedrock, despite zero enterprise demand
The energy drink of frontier models
Cisco sings Mythos' praises - but doesn't say how many bugs the model uncovered
Meanwhile, Anthropic adds 150 partners to Project Glasswing
AI + ML
Netflix wiz creates app to slash AI bills, then open sources it
Security
Disgruntled 0-day hunter 'humiliated' by Microsoft pledges 'bone shattering drop' as Redmond calls cops
Security
Troops’ phones gave away location data to foreign adversaries
AI + ML
Google has seriously leaned into AI enshittification lately
Security
Anthropic to release Mythos-class models to the public
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
You don't infect anyone in Russia or other CIS countries
Lacking an enterprise content layer for Headless 360, CRM titan went shopping
Government gets a say in 'trusted partner' access, and that worries policy experts
Meanwhile, Anthropic adds 150 partners to Project Glasswing
Young professionals may be perfectly productive while working from home, says the New York Fed, but the quality of their output isn't so great, so companies don't want to hire them
You don't infect anyone in Russia or other CIS countries
Lacking an enterprise content layer for Headless 360, CRM titan went shopping
Government gets a say in 'trusted partner' access, and that worries policy experts
Meanwhile, Anthropic adds 150 partners to Project Glasswing
Young professionals may be perfectly productive while working from home, says the New York Fed, but the quality of their output isn't so great, so companies don't want to hire them
'Dumbass' criminal breaks the 'first rule of ransomware club'
You don't infect anyone in Russia or other CIS countries
Contentful is a shot in the arm for Salesforce's 'headless' bet
Lacking an enterprise content layer for Headless 360, CRM titan went shopping
Trump's AI E-(I)-O could let feds pick winners and losers
Government gets a say in 'trusted partner' access, and that worries policy experts
Cisco sings Mythos' praises - but doesn't say how many bugs the model uncovered
Meanwhile, Anthropic adds 150 partners to Project Glasswing
Remote work – not AI – is killing job prospects for the youth
Young professionals may be perfectly productive while working from home, says the New York Fed, but the quality of their output isn't so great, so companies don't want to hire them
Marvell enters the AI network fray with 102.4 Tbps switch silicon
High radix, low latency and low power is what AI datacenters crave, the chipmaker says
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

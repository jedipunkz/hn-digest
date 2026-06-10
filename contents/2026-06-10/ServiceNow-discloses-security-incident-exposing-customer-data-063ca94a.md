---
source: "https://www.bleepingcomputer.com/news/security/servicenow-discloses-security-incident-exposing-customer-data/"
hn_url: "https://news.ycombinator.com/item?id=48469241"
title: "ServiceNow discloses security incident exposing customer data"
article_title: "ServiceNow discloses security incident exposing customer data"
author: "coloneltcb"
captured_at: "2026-06-10T01:00:06Z"
capture_tool: "hn-digest"
hn_id: 48469241
score: 9
comments: 0
posted_at: "2026-06-09T23:30:40Z"
tags:
  - hacker-news
  - translated
---

# ServiceNow discloses security incident exposing customer data

- HN: [48469241](https://news.ycombinator.com/item?id=48469241)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/servicenow-discloses-security-incident-exposing-customer-data/)
- Score: 9
- Comments: 0
- Posted: 2026-06-09T23:30:40Z

## Translation

タイトル: ServiceNow が顧客データを漏洩するセキュリティ インシデントを公開
説明: ServiceNow は、攻撃者が脆弱な API エンドポイントを介して未認証アクセスの欠陥を悪用し、顧客インスタンスのデータをクエリできるようにしたことによるセキュリティ インシデントについて警告しています。

記事本文:
ServiceNow、顧客データを漏洩するセキュリティインシデントを公開
Microsoft 2026 年 6 月のパッチ火曜日で 3 つのゼロデイ、200 の欠陥が修正される
GitHub はパスワードを盗むマルウェアをプッシュする Microsoft リポジトリを無効にします
フランス政府のメッセージングサービスがアカウントハイジャック攻撃で侵害される
Google、実際に悪用されたChromeの新たなゼロデイ欠陥にパッチを適用
Microsoft Defender の「RoguePlanet」ゼロデイは SYSTEM 権限を付与します
ServiceNow、顧客データを漏洩するセキュリティインシデントを公開
OpenClaw AI エージェントがフィッシング攻撃に遭い、ユーザーデータを流出させたことが判明
SAP、NetWeaver と Commerce Cloud の重大な欠陥を修正
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
ServiceNow、顧客データを漏洩するセキュリティインシデントを公開
ServiceNow、顧客データを漏洩するセキュリティインシデントを公開
ServiceNow は、攻撃者が脆弱な API エンドポイントを介して未認証アクセスの欠陥を悪用し、顧客インスタンスのデータをクエリできるようにしたことによるセキュリティ インシデントについて警告しています。
同社は、この問題に関連する「異常なアクティビティ」を検出した後、サポート情報や直接のサポートケースを通じて、影響を受ける顧客に静かに警告しました。
ServiceNow のカスタマー サポート ログイン ポータルの背後に隠されているこの情報には、同社が 2026 年 6 月 5 日にホストされている顧客インスタンスにセキュリティ アップデートを適用したことが記載されています。
「2026 年 6 月 5 日、ServiceNow はホストされた顧客インスタンスにセキュリティ アップデートを適用しました」とサポート速報には記載されています。
「このアップデートは、認証されていないユーザーのアクセスを許可する可能性のあるセキュリティ上の問題に関するものでした。

ユーザーは、特定の状況において、意図した以上に ServiceNow インスタンスにアクセスできるようになります。」
同社によると、このセキュリティアップデートによりAPIエンドポイントの設定が変更され、認証されたユーザーのみにアクセスが制限されるという。
ServiceNow はまた、攻撃者がこの欠陥を悪用して顧客インスタンス テーブルのクエリに成功したことを確認しました。
ServiceNowは攻撃中にどのデータがアクセスされたかを明らかにしていないが、インスタンスには通常、ITサポートチケット、従業員記録、内部文書、資産インベントリ、セキュリティインシデントレポート、ワークフローデータ、企業システムやサービスの構成詳細などの企業機密情報が保存されている。
チケットには認証情報、API トークン、内部文書、トラブルシューティング中に共有される認証秘密が含まれる可能性があるため、サポート ケースの情報は、脅威アクターにとってますます人気のあるターゲットになっています。
この勧告によると、ServiceNow は現在、影響を受ける顧客に対してサポート ケースを開始しています。顧客がそれを受け取っていない場合、その顧客はインシデントの影響を受けるとは考えられません。
ServiceNow はこの欠陥に関する技術的な詳細を公表していないが、Reddit でこのインシデントについて議論している管理者らは、この問題は「 /api/now/relative_list_edit/create 」の REST エンドポイントに関連しているようだと述べている。
あるコメント投稿者は、エンドポイントが「 require_authentication=false 」で構成されており、認証されていないリクエストによるインスタンス データへのアクセスを許可する可能性があると主張しました。金曜日にリリースされたセキュリティ更新プログラムは、requires_authentication を true に設定するために使用されたとされています。
多くの管理者が、IP アドレス「 51.159.98.241 」からの API リクエストなどの侵害の痕跡を共有し、他の管理者に脆弱なエンドポイントへのリクエストのログを確認するようアドバイスしました。
速報では、この問題は主に顧客に影響を与えると述べています。

オーストラリアのプラットフォーム リリース、または特定の構成変更を行った古いリリースを使用している顧客。
ServiceNowは「このセキュリティ問題は、オーストラリアのプラットフォームリリースを使用している顧客、またはオーストラリアより前のリリースでインスタンスに特定の構成変更を行った顧客に関係する」と警告した。
BleepingComputer は本日、読者からこの事件について警告を受けた後、ServiceNow に連絡し、活動がどのくらいの期間継続していたのか、問題の原因は何か、顧客データが盗まれたかどうかを尋ねました。公開前に返答は得られませんでした。
ServiceNowは、この問題についてCVEを公開するかどうかをまだ検討中だとしている。
管理者は、ServiceNow ログで、特に IP アドレス 51.159.98.241 からの /api/now/relative_list_edit へのリクエストを確認することをお勧めします。
影響を受ける組織は、機密情報の公開されたチケットと記録を確認し、サポート ワークフローを通じて共有される認証情報またはトークンをローテーションし、API ログが有効になっていることを確認する必要があります。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
ハッカーが盗んだデータを販売していると主張し、Vercelが侵害を確認
フランス政府のメッセージングサービスがアカウントハイジャック攻撃で侵害される
SoFi、香港子会社におけるサードパーティによるデータ侵害を確認
オックスフォード大学、キャリアプラットフォームハッキング後のデータ侵害を公表
Meta AIサポートハッキングで2万以上のInstagramアカウントが盗まれる
まだメンバーではありませんか?今すぐ登録
UniFi OS の重大なバグにより、ハッカーが認証なしで root を取得できる
Everest Forms Pro の重大な欠陥が悪用され、WordPress サイトが乗っ取られる
AI を活用した Wi-Fi であるインテリジェント ターミナルを実際に使ってみよう

ウィンドウズターミナル
セキュリティリーダーが今後 6 か月間で準備すべきこと。
2026 年の医療資格情報漏えいの現状: (非ゲート) レポートを読む
最後に侵入テストを行ったのは 345 日前です。それ以来何が変わりましたか？
AI ツールが機密データを漏洩しています。無料の監査を受けてください。
Wazuh でサイバー回復力を構築: プロアクティブな保護のためのオープンソース SIEM および XDR
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

ServiceNow is warning about a security incident after attackers exploited an unauthenticated access flaw through a vulnerable API endpoint, allowing them to query data from customer instances.

ServiceNow discloses security incident exposing customer data
Microsoft June 2026 Patch Tuesday fixes 3 zero-day, 200 flaws
GitHub disables Microsoft repos pushing password-stealing malware
French govt messaging service breached in account hijacking attack
Google patches new Chrome zero-day flaw exploited in the wild
Microsoft Defender 'RoguePlanet' zero-day grants SYSTEM privileges
ServiceNow discloses security incident exposing customer data
OpenClaw AI agent found falling for phishing attacks, spills user data
SAP fixes critical flaws in NetWeaver and Commerce Cloud
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
ServiceNow discloses security incident exposing customer data
ServiceNow discloses security incident exposing customer data
ServiceNow is warning about a security incident after attackers exploited an unauthenticated access flaw through a vulnerable API endpoint, allowing them to query data from customer instances.
The company quietly warned impacted customers through a support bulletin and direct support cases after detecting "anomalous activity" related to the issue.
The bulletin, which is hidden behind ServiceNow's customer support login portal, states that the company applied a security update to hosted customer instances on June 5, 2026.
"On June 5, 2026, ServiceNow applied a security update to hosted customer instances," reads the support bulletin.
"The update concerned a security issue that could allow an unauthenticated user, in certain circumstances, to gain greater access to ServiceNow instances than intended."
The company says this security update changes the API endpoint configuration to limit access to authenticated users only.
ServiceNow also confirmed that attackers exploited this flaw to successfully query the customer instance tables.
While ServiceNow did not disclose which data was accessed during the attacks, instances commonly store sensitive enterprise information, including IT support tickets, employee records, internal documentation, asset inventories, security incident reports, workflow data, and configuration details for corporate systems and services.
Support case information has become an increasingly popular target for threat actors , as tickets can contain credentials, API tokens, internal documentation, and authentication secrets shared during troubleshooting.
According to the advisory, ServiceNow has now opened support cases with affected customers. If a customer has not received one, they are not believed to be affected by the incident.
While ServiceNow has not publicly disclosed technical details about the flaw, administrators discussing the incident on Reddit say the issue appears to be tied to a REST endpoint at ' /api/now/related_list_edit/create '.
One commenter claimed the endpoint was configured with ' requires_authentication=false ', potentially allowing unauthenticated requests to access instance data. The security update released on Friday was allegedly used to set requires_authentication to true .
Numerous admins shared indicators of compromise, including API requests from the IP address ' 51.159.98.241 ,' advising other administrators to review logs for requests to the vulnerable endpoint.
The bulletin states the issue primarily impacts customers running the Australia platform release or customers on older releases who made certain configuration changes.
"The security issue pertains to customers who are on the Australia platform release or made certain configuration changes to instances on releases prior to Australia," ServiceNow warned.
BleepingComputer contacted ServiceNow earlier today after a reader alerted us to the incident, asking how long the activity had been ongoing, what caused the issue, and whether customer data had been stolen. We did not receive a response before publication.
ServiceNow says it is still evaluating whether it will publish a CVE for the issue.
Administrators are advised to review ServiceNow logs for requests to /api/now/related_list_edit, particularly from the IP address 51.159.98.241.
Impacted organizations should review exposed tickets and records for sensitive information, rotate credentials or tokens shared through support workflows, and ensure API logging is enabled.
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
Vercel confirms breach as hackers claim to be selling stolen data
French govt messaging service breached in account hijacking attack
SoFi confirms third-party data breach at Hong Kong subsidiary
Oxford University discloses data breach after careers platform hack
Over 20,000 Instagram accounts stolen in Meta AI support hack
Not a member yet? Register Now
Critical UniFi OS bug lets hackers gain root without authentication
Critical Everest Forms Pro flaw exploited to take over WordPress sites
Hands on with Intelligent Terminal, an AI-powered Windows Terminal
What security leaders should prepare for over the next six months.
The State of Healthcare Credential Exposure in 2026: Read the (Ungated) Report
Your last pentest was 345 days ago. What changed since then?
Your AI tools are leaking sensitive data. Get a free audit.
Build cyber resilience with Wazuh: The open-source SIEM & XDR for proactive protection
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.

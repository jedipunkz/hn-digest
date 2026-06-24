---
source: "https://cyber.netsecops.io/articles/openclaws-skill-marketplace-and-the-emerging-ai-supply-chain-threat/"
hn_url: "https://news.ycombinator.com/item?id=48662618"
title: "Malicious AI 'Skills' on OpenClaw's ClawHub Marketplace Bypass Scanners"
article_title: "Malicious AI 'Skills' on OpenClaw's ClawHub Marketplace Bypass Scanners to Deliver Infostealers - CyberNetSec.io"
author: "jaybode"
captured_at: "2026-06-24T17:37:12Z"
capture_tool: "hn-digest"
hn_id: 48662618
score: 5
comments: 0
posted_at: "2026-06-24T16:52:03Z"
tags:
  - hacker-news
  - translated
---

# Malicious AI 'Skills' on OpenClaw's ClawHub Marketplace Bypass Scanners

- HN: [48662618](https://news.ycombinator.com/item?id=48662618)
- Source: [cyber.netsecops.io](https://cyber.netsecops.io/articles/openclaws-skill-marketplace-and-the-emerging-ai-supply-chain-threat/)
- Score: 5
- Comments: 0
- Posted: 2026-06-24T16:52:03Z

## Translation

タイトル: OpenClaw の ClawHub マーケットプレイス バイパス スキャナー上の悪意のある AI 'スキル'
記事タイトル: OpenClaw の ClawHub マーケットプレイス上の悪意のある AI 'スキル' スキャナーをバイパスして情報窃取者を配信 - CyberNetSec.io
説明: Unit 42 の脅威調査により、悪意のある新たな AI サプライ チェーン攻撃が明らかになりました。

記事本文:
印刷 ホーム レポート OpenClaw の ClawHub マーケットプレイス上の悪意のある AI 'スキル' スキャナーをバイパスして情報を盗む者を配信
Unit 42 が OpenClaw の ClawHub マーケットプレイスで情報窃取者を展開する回避型の悪意のある AI スキルを発見
OpenClaw の ClawHub マーケットプレイスの悪意のある AI 'スキル' がスキャナーをバイパスし、情報窃取を実現
2026 年 2 月から 5 月にかけて、Unit 42 の研究者は、OpenClaw AI エージェント エコシステムを標的とした高度な脅威キャンペーンを発見しました。悪意のある攻撃者は、VirusTotal を含む統合セキュリティ スキャナーをバイパスする危険な「スキル」を、公式マーケットプレイスである ClawHub で公開することに成功しています。これらのスキルは、ソーシャル エンジニアリングと難読化を活用して、ユーザーをだまして、Atomic macOS Stealer (AMOS) や cluw という名前の新しい亜種などの情報窃取マルウェアを展開するコマンドを実行させます。このアクティビティは、ソフトウェア サプライ チェーン攻撃の重要な進化を表しており、特にエージェント AI プラットフォームの独自のアーキテクチャに適応しています。これらの環境では分離が欠如しているため、単一の悪意のあるスキルによって攻撃者にエージェントの権限と基盤となるシステムへのアクセスに対する完全な制御が与えられる可能性があり、ユーザーと組織に重大なリスクをもたらす可能性があります。
OpenClaw は、専用の ClawHub マーケットプレイスを通じて配布される「スキル」と呼ばれるサードパーティのプラグインを使用してタスクを実行するように設計された AI エージェントです。このモデルは、新しいタイプのソフトウェア サプライ チェーンを作成します。 ClawHavoc など、2026 年初頭の最初の悪意のあるキャンペーンが特定され、VirusTotal と ClawScan によるスキャンの強化につながりましたが、脅威アクターは適応しました。
Unit 42 が観測した最新のキャンペーンでは、より回避的な手法が使用されています。攻撃者は、 tradeview-ai-indicator-assistant などの正規のものと思われるスキルを公開します。ただし、スキルのマークダウン ファイルには

悪意のあるコマンドをホストする外部 Web サイト (「ペーストサイト リダイレクト ルアー」) にユーザーを誘導する「前提条件ブロック」が組み込まれています。ユーザーは、このコマンドをコピーして端末に貼り付けてスキルを有効にするように指示されます。このユーザー支援の実行は、スキル パッケージ自体を分析するだけの自動スキャナーをバイパスします。このコマンドが実行されると、情報窃盗ペイロードがダウンロードされて実行され、資格情報の盗難や潜在的な金融詐欺につながります。
この攻撃ベクトルは、AI エージェントのセキュリティにおけるセマンティック ギャップを悪用します。エージェントは悪意のある指示を正当なユーザー要求として解釈し、独自のシステム権限を使用して攻撃を実行します。これにより、npm や PyPI などのサンドボックス アプリケーション環境に存在する可能性のある従来のセキュリティ境界が回避されます。
攻撃チェーンは主に、悪意のあるスキルによって引き起こされるユーザーの操作に依存しています。
ルアー : ユーザーは ClawHub から、 tradeview-ai-indicator-assistant (SHA256: b6c7e0bf573b1c7d9d3a05eb08d26579199515b847df984862805f44a7af8007 ) などの悪意のあるスキルをインストールします。
ソーシャル エンジニアリング : スキルの前提条件の指示により、ユーザーは貼り付けサイト hxxps[:]//rentry[.]co/openclaw-code に誘導され、必須のアクティベーション ステップを装います。これは T1189 - ドライブバイ侵害の一種です。
実行 : ユーザーは、Base64 でエンコードされた文字列をコピーし、それをシェルにパイプするように指示されます。この手法、 T1059.004 - コマンドおよびスクリプト インタプリタ: Unix Shell は、古典的な 'curl-pipe-bash' 攻撃です。 Base64 の使用は、T1027 難読化されたファイルまたは情報の一種です。
ペイロード配信: 実行されたシェル コマンドは、T1105 - Ingress Tool Transfer を介して第 2 段階のペイロードをフェッチします。 TradingView スキルの場合、ペイロード Xuvewuyur は hxxp[:]//2.26.75[.]16 からダウンロードされました。このペイロードは次のように識別されました

cluw という名前の新しい macOS インフォスティーラー (SHA256: 818aea6143282b352fdfdc0f3ebf77a36e54eb3befb5cad1a355a99ab97c6aa7 )。
C2 通信とデータの盗難: インフォスティーラーがアクティブになると、資格情報やその他の機密データが収集され、T1555 - パスワード ストアからの資格情報の目的が達成されます。 omn​​icogg スキル (SHA256: b30eaed1f7478c28f4ec50d07ed5ef014ffbc4b2bc5a38d689ba9f7abb5e19c2) にリンクされた古いキャンペーンは、91.92.242[.]30 の C2 サーバーと通信する Atomic macOS スティーラー (AMOS) を配信しました。
このキャンペーンは、攻撃者の粘り強さを示しており、元の ClawHavoc 攻撃の配信テンプレートを再利用していますが、新しいバックエンド インフラストラクチャとペイロードを使用して検出を回避しています。
このキャンペーンの主な影響は、ブラウザの Cookie、仮想通貨ウォレット データ、システム パスワード、被害者のマシンに保存されているその他の資格情報などの機密情報が盗まれることです。 TradingViewユーザーのターゲティングは、金融市場に関与する個人に焦点を当てていることを示唆しており、直接的な金銭的損失のリスクが増大しています。
より広い観点から見ると、この攻撃は、急成長する AI エージェント エコシステムにおける深刻なシステムリスクを浮き彫りにしています。サードパーティのスキルに対する堅牢なサンドボックスと権限制御が欠如しているため、マルウェアがユーザー システムに直接侵入する信頼できる経路が作成されます。 AI エージェントが個人や企業のワークフローへの統合が進むにつれて、この種のサプライ チェーン攻撃は、広範な企業スパイ活動、大規模なデータ侵害、重大な金融詐欺につながる可能性があります。
サイバー観測可能物 — 狩猟のヒント
セキュリティ チームは、関連するアクティビティを検出するために次のパターンを探す必要がある場合があります。
この脅威を検出するには、最初のスキルのダウンロード以降も監視する必要があります。セキュリティ チームは、インストール後の動作に重点を置く必要があります。
プロセス監視：実装

OpenClaw エージェントから発生する不審なプロセス チェーンを監視するためのエンドポイント検出と応答 (EDR) ルール。具体的には、OpenClaw が bash や sh などのシェル インタープリタを生成し、その後、curl や wget などのツールを使用してネットワーク接続を開始することについて警告します。これは、D3FEND の D3-PA: プロセス分析を通じて実現できます。
コマンドライン監査 : 実行されたプロセスのすべてのコマンドライン引数をログに記録します。のようなパターンの SIEM アラートを作成します。 bash または Base64 --decode | bash は、こ​​の攻撃ベクトルを示すものです。
ネットワーク トラフィック分析: ネットワーク セキュリティ ツールとプロキシを使用して、D3-NTA: ネットワーク トラフィック分析を実行します。上記の IOC へのアウトバウンド接続をブロックします。さらに、機密性の高いシステムや通常とは異なるプロセスによる、rentry.co や Pastebin.com などの既知の匿名ペースト サイトへの接続に対するアラートを作成します。
ファイル整合性監視 : ユーザー ディレクトリ内の予期しない実行可能ファイルの作成を監視します。これは、ダウンロードされたペイロードを示す可能性があります。
侵害が疑われる場合は、影響を受けたホストをネットワークからただちに隔離し、マシンに保存されている可能性のある資格情報を取り消し、侵害の範囲を特定するためのフォレンジック調査を開始します。
この脅威を軽減するには、技術的な制御とユーザーの意識を組み合わせる必要があります。
ユーザートレーニング : これは最も重要な防御策です。 AI エージェントのユーザーに、サードパーティのスキル マーケットプレイスの危険性について教育します。具体的には、たとえ機能を有効にするために必要な手順として提示されたとしても、信頼できないソースから端末にコマンドをコピー アンド ペーストしないように教育します。これは、MITRE ATT&CK Mitigation M1017 - ユーザー トレーニングと一致しています。
アプリケーション制御: アプリケーションのホワイトリストポリシーを実装して、未承認のスクリプトやバイナリの実行を防止します。

s.厳格なポリシーにより、OpenClaw などのアプリケーションによるシェル インタープリターの呼び出しをブロックできます。これは、 D3FEND の D3-EAL: 実行可能ファイルの許可リストに対応します。
最小特権の原則 : OpenClaw などの AI エージェントを必要最小限の権限で実行します。可能であれば、コンテナ化またはサンドボックス技術を使用して、エージェントとそのスキルを基盤となるオペレーティング システムや機密性の高いユーザー データから分離します。これは、M1048 - アプリケーションの分離とサンドボックス化に関連します。
ネットワーク フィルタリング : ファイアウォールと Web プロキシにアウトバウンド トラフィック フィルタリング ルールを実装し、既知の悪意のある IP とrentry.co ドメインへのアクセスをブロックします。これは、D3FEND の D3-OTF: アウトバウンド トラフィック フィルタリングを直接適用したものです。
AI マーケットプレイスのリスクと、信頼できないソースからコマンドを実行する危険性についてユーザーを教育します。
アプリケーション制御ソリューションを使用して、OpenClaw などのエージェントがシェル インタプリタを生成したり、任意のコードを実行したりするのを防ぎます。
Web フィルターを使用して、既知の悪意のあるドメインや信頼できない貼り付けサイトへのアクセスをブロックします。
エンドポイント保護を導入して、AMOS や cluw などの既知の情報窃取ペイロードを検出してブロックします。
AI エージェントをサンドボックス環境またはコンテナ化環境で実行して、ホスト システムおよびユーザー データへのアクセスを制限します。
コマンドラインアクティビティとプロセス作成の包括的なログを有効にして、不審な動作を検出します。
D3FENDの防御策
OpenClaw などの AI エージェントを実行しているすべてのシステムに詳細なプロセス検査を実行できるエンドポイント検出および応答 (EDR) ソリューションを導入します。不審なプロセスの祖先を特別に監視するように EDR を構成します。重要なルールは、「OpenClaw」プロセスが子プロセスとしてシェル インタプリタ ( bash 、 sh 、 zsh など) を生成し、それによってネットワーク UT が生成されるときに、重大度の高いアラートを生成することです。

curl や wget のような機能。この特定のチェーンは、説明されている攻撃パターンをよく示しています。通常の OpenClaw 動作のベースラインを確立します。逸脱、特に任意のスクリプトの実行やシェルへの直接アクセスは、直ちに調査する必要があります。この技術は、AI エージェントが実行する不透明なアクションを可視化することで、脅威アクターの実行方法に直接対抗します。
境界ファイアウォールと Web プロキシに厳格な出力フィルタリング ルールを実装します。少なくとも、既知の悪意のある IP アドレス 91.92.242.30 および 2.26.75.16 に対して明示的なブロック ルールを作成します。より戦略的には、「貼り付けサイト」または「匿名化サービス」 (rentry.co など) へのアクセスをブロックするカテゴリベースのフィルタリング ポリシーを作成します。より高度なセキュリティが必要な環境の場合は、デフォルト拒否の送信ポリシーを採用し、既知のビジネスに不可欠なドメインおよび IP アドレスへのトラフィックのみを許可リストに登録します。この制御により、悪意のあるスクリプトが貼り付けサイトからダウンロードされるのを阻止し、攻撃者のサーバーからの最終ペイロードのダウンロードをブロックするという 2 つの点で攻撃チェーンが遮断されます。これは、ペイロード自体がウイルス対策によってまだ検出されていない可能性がある新しい脅威を補う重要な制御です。
macOS および OpenClaw が使用されるその他のシステムでは、強制モードでアプリケーション制御ソリューションを実装します。 OpenClaw アプリケーションが、そのコアの署名済みコンポーネントの一部ではない子プロセスを実行しないようにする厳格なポリシーを作成します。具体的には、 /bin/bash 、 /bin/sh 、および /usr/bin/curl を起動する機能を拒否します。これにより、「カール・パイプ・バッシュ」テクニックが成功することができなくなります。これにより、一部の正当ではあるものの設計が不十分なスキルが壊れる可能性がありますが、このクラスの攻撃全体に対する強力な防御手段となります。

タック。これにより、AI エージェントの実行環境が効果的に強化され、ユーザーがだまされて悪意のあるアクションを開始した場合でも、ペイロードがダウンロードまたは実行される前に、基盤となるオペレーティング システムが危険な動作をブロックします。これにより、防御は検出から防止へと移行します。
Bitdefender、Koi Security、Trend Micro からの最初のレポートでは、ClawHub における悪意のあるスキルの第 1 波について詳しく説明しています。
ユニット 42 は、2026 年 5 月まで続く ClawHub マーケットプレイスの分析を開始します。
悪意のある「tradingview-ai-indicator-assistant」スキルが ClawHub に公開されます。
ClawHub は、スキル スクリーニングを強化するために NVIDIA との提携を発表しました。
Unit 42 は、悪意のある回避スキルと新たな AI サプライ チェーンの脅威に関する調査を発表します。
セキュリティ運用、脅威インテリジェンス、インシデント対応、セキュリティ自動化において 10 年以上の専門的な経験を持つサイバーセキュリティの専門家。専門知識は、SOAR/XSOAR オーケストレーション、脅威インテリジェンス プラットフォーム、SIEM/UEBA 分析、サイバー フュージョン センターの構築に及びます。背景には、技術的なイネーブルメント、企業および政府クライアント向けのソリューション アーキテクチャ、組織全体にわたるセキュリティ自動化ワークフローの実装が含まれます。

[切り捨てられた]

## Original Extract

Unit 42 threat research uncovers a new AI supply chain attack where malicious

Print Home Report Malicious AI 'Skills' on OpenClaw's ClawHub Marketplace Bypass Scanners to Deliver Infostealers
Unit 42 Uncovers Evasive Malicious AI Skills on OpenClaw's ClawHub Marketplace Deploying Infostealers
Malicious AI 'Skills' on OpenClaw's ClawHub Marketplace Bypass Scanners to Deliver Infostealers
Between February and May 2026, Unit 42 researchers uncovered a sophisticated threat campaign targeting the OpenClaw AI agent ecosystem. Malicious actors are successfully publishing dangerous 'skills' on ClawHub , the official marketplace, that bypass integrated security scanners, including VirusTotal . These skills leverage social engineering and obfuscation to trick users into executing commands that deploy infostealer malware, such as Atomic macOS stealer (AMOS) and a new variant named cluw . This activity represents a critical evolution of software supply chain attacks, specifically adapted for the unique architecture of agentic AI platforms. The lack of isolation in these environments means a single malicious skill can grant an attacker full control over the agent's permissions and access to the underlying system, posing a severe risk to users and organizations.
OpenClaw is an AI agent designed to execute tasks using third-party plugins called 'skills', which are distributed through its dedicated ClawHub marketplace. This model creates a new type of software supply chain. While initial malicious campaigns in early 2026, such as ClawHavoc , were identified and led to enhanced scanning with VirusTotal and ClawScan , threat actors have adapted.
The latest campaign, observed by Unit 42 , uses more evasive techniques. Attackers publish skills, such as tradingview-ai-indicator-assistant , that appear legitimate. However, the skill's markdown file contains a 'prerequisite block' that directs the user to an external website (a 'paste-site redirect lure') hosting a malicious command. The user is instructed to copy and paste this command into their terminal to enable the skill. This user-assisted execution bypasses the automated scanners that only analyze the skill package itself. Once executed, the command downloads and runs an infostealer payload, leading to credential theft and potential financial fraud.
This attack vector exploits the semantic gap in AI agent security. The agent interprets the malicious instructions as a legitimate user request, using its own system privileges to execute the attack. This circumvents traditional security boundaries that might exist in sandboxed application environments like npm or PyPI.
The attack chain primarily relies on user interaction prompted by a malicious skill.
Lure : The user installs a malicious skill from ClawHub , such as tradingview-ai-indicator-assistant (SHA256: b6c7e0bf573b1c7d9d3a05eb08d26579199515b847df984862805f44a7af8007 ).
Social Engineering : The skill's prerequisite instructions direct the user to a paste-site, hxxps[:]//rentry[.]co/openclaw-code , which masquerades as a required activation step. This is a form of T1189 - Drive-by Compromise .
Execution : The user is instructed to copy a Base64-encoded string and pipe it into a shell. This technique, T1059.004 - Command and Scripting Interpreter: Unix Shell , is a classic 'curl-pipe-bash' attack. The use of Base64 is a form of T1027 - Obfuscated Files or Information .
Payload Delivery : The executed shell command fetches a second-stage payload via T1105 - Ingress Tool Transfer . In the case of the tradingview skill, the payload Xuvewuyur was downloaded from hxxp[:]//2.26.75[.]16 . This payload was identified as a new macOS infostealer named cluw (SHA256: 818aea6143282b352fdfdc0f3ebf77a36e54eb3befb5cad1a355a99ab97c6aa7 ).
C2 Communication & Data Theft : Once active, the infostealer harvests credentials and other sensitive data, fulfilling its objective of T1555 - Credentials from Password Stores . Older campaigns linked to the omnicogg skill (SHA256: b30eaed1f7478c28f4ec50d07ed5ef014ffbc4b2bc5a38d689ba9f7abb5e19c2 ) delivered Atomic macOS stealer (AMOS) , communicating with a C2 server at 91.92.242[.]30 .
This campaign demonstrates the attackers' persistence, reusing the delivery template from the original ClawHavoc attacks but with new backend infrastructure and payloads to evade detection.
The primary impact of this campaign is the theft of sensitive information, including browser cookies, cryptocurrency wallet data, system passwords, and other credentials stored on the victim's machine. The targeting of TradingView users suggests a focus on individuals involved in financial markets, increasing the risk of direct financial loss.
From a broader perspective, this attack highlights a severe systemic risk in the burgeoning AI agent ecosystem. The lack of robust sandboxing and permission controls for third-party skills creates a trusted pathway for malware directly onto user systems. As AI agents become more integrated into personal and enterprise workflows, this type of supply chain attack could lead to widespread corporate espionage, large-scale data breaches, and significant financial fraud.
Cyber Observables — Hunting Hints
Security teams may want to hunt for the following patterns to detect related activity:
Detecting this threat requires monitoring beyond the initial skill download. Security teams should focus on post-installation behavior.
Process Monitoring : Implement Endpoint Detection and Response (EDR) rules to monitor for suspicious process chains originating from the OpenClaw agent. Specifically, alert on OpenClaw spawning shell interpreters like bash or sh , which then initiate network connections with tools like curl or wget . This can be achieved through D3FEND 's D3-PA: Process Analysis .
Command Line Auditing : Log all command-line arguments for executed processes. Create SIEM alerts for patterns like curl | bash or base64 --decode | bash , which are highly indicative of this attack vector.
Network Traffic Analysis : Use network security tools and proxies to perform D3-NTA: Network Traffic Analysis . Block outbound connections to the IOCs listed above. Additionally, create alerts for connections to known anonymous paste sites like rentry.co or pastebin.com from sensitive systems or by unusual processes.
File Integrity Monitoring : Monitor for the creation of unexpected executable files in user directories, which may indicate a downloaded payload.
If a compromise is suspected, immediately isolate the affected host from the network, revoke any credentials that may have been stored on the machine, and begin a forensic investigation to determine the extent of the breach.
Mitigating this threat requires a combination of technical controls and user awareness.
User Training : This is the most critical defense. Educate users of AI agents about the dangers of third-party skill marketplaces. Specifically, train them to never copy and paste commands from untrusted sources into a terminal, even if presented as a necessary step to enable a feature. This aligns with MITRE ATT&CK Mitigation M1017 - User Training .
Application Control : Implement application allowlisting policies to prevent the execution of unauthorized scripts and binaries. A strict policy could block shell interpreters from being invoked by applications like OpenClaw . This corresponds to D3FEND 's D3-EAL: Executable Allowlisting .
Principle of Least Privilege : Run AI agents like OpenClaw with the minimum necessary permissions. If possible, use containerization or sandboxing technologies to isolate the agent and its skills from the underlying operating system and sensitive user data. This relates to M1048 - Application Isolation and Sandboxing .
Network Filtering : Implement outbound traffic filtering rules on firewalls and web proxies to block access to the known malicious IPs and the rentry.co domain. This is a direct application of D3FEND 's D3-OTF: Outbound Traffic Filtering .
Educate users on the risks of AI marketplaces and the danger of executing commands from untrusted sources.
Use application control solutions to prevent agents like OpenClaw from spawning shell interpreters or executing arbitrary code.
Use web filters to block access to known malicious domains and untrusted paste sites.
Deploy endpoint protection to detect and block known infostealer payloads like AMOS and cluw.
Run AI agents in a sandboxed or containerized environment to limit their access to the host system and user data.
Enable comprehensive logging of command-line activity and process creation to detect suspicious behavior.
D3FEND Defensive Countermeasures
Deploy an Endpoint Detection and Response (EDR) solution capable of deep process inspection on all systems running AI agents like OpenClaw. Configure the EDR to specifically monitor for suspicious process ancestry. A key rule should be to generate a high-severity alert when the 'OpenClaw' process spawns a shell interpreter (e.g., bash , sh , zsh ) as a child process, which in turn spawns a networking utility like curl or wget . This specific chain is highly indicative of the attack pattern described. Establish a baseline of normal OpenClaw behavior; any deviation, especially the execution of arbitrary scripts or direct shell access, should be investigated immediately. This technique directly counters the threat actor's execution method by providing visibility into the otherwise opaque actions taken by the AI agent.
Implement strict egress filtering rules on perimeter firewalls and web proxies. At a minimum, create an explicit block rule for the known malicious IP addresses 91.92.242.30 and 2.26.75.16 . More strategically, create a category-based filtering policy that blocks access to 'Paste Sites' or 'Anonymizing Services', which would include rentry.co . For environments requiring higher security, adopt a default-deny outbound policy and only allowlist traffic to known-good, business-essential domains and IP addresses. This control would have broken the attack chain at two points: preventing the malicious script from being downloaded from the paste site, and blocking the final payload download from the attacker's server. This is a crucial compensating control for novel threats where the payload itself may not yet be detected by antivirus.
On macOS and other systems where OpenClaw is used, implement an application control solution in enforcement mode. Create a strict policy that prevents the OpenClaw application from executing any child processes that are not part of its core, signed components. Specifically, deny its ability to launch /bin/bash , /bin/sh , and /usr/bin/curl . This prevents the 'curl-pipe-bash' technique from ever succeeding. While this may break some legitimate but poorly designed skills, it provides a powerful defense against this entire class of attack. This effectively hardens the AI agent's execution environment, ensuring that even if a user is tricked into initiating a malicious action, the underlying operating system will block the dangerous behavior before the payload can be downloaded or executed. This moves the defense from detection to prevention.
Initial reports from Bitdefender, Koi Security, and Trend Micro detail the first wave of malicious skills on ClawHub.
Unit 42 begins analysis of the ClawHub marketplace, lasting through May 2026.
The malicious 'tradingview-ai-indicator-assistant' skill is published to ClawHub.
ClawHub announces a partnership with NVIDIA to enhance skill screening.
Unit 42 publishes its research on the evasive malicious skills and the emerging AI supply chain threat.
Cybersecurity professional with over 10 years of specialized experience in security operations, threat intelligence, incident response, and security automation. Expertise spans SOAR/XSOAR orchestration, threat intelligence platforms, SIEM/UEBA analytics, and building cyber fusion centers. Background includes technical enablement, solution architecture for enterprise and government clients, and implementing security automation workflows across I

[truncated]

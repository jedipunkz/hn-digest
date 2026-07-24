---
source: "https://www.bleepingcomputer.com/news/security/hermes-ai-agent-used-to-automate-attack-on-thai-finance-ministry/"
hn_url: "https://news.ycombinator.com/item?id=49042205"
title: "Hermes AI agent used to automate attack on Thai Finance Ministry"
article_title: "Hermes AI agent used to automate attack on Thai Finance Ministry"
author: "sbulaev"
captured_at: "2026-07-24T22:55:28Z"
capture_tool: "hn-digest"
hn_id: 49042205
score: 2
comments: 0
posted_at: "2026-07-24T22:07:06Z"
tags:
  - hacker-news
  - translated
---

# Hermes AI agent used to automate attack on Thai Finance Ministry

- HN: [49042205](https://news.ycombinator.com/item?id=49042205)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/hermes-ai-agent-used-to-automate-attack-on-thai-finance-ministry/)
- Score: 2
- Comments: 0
- Posted: 2026-07-24T22:07:06Z

## Translation

タイトル: タイ財務省への攻撃を自動化するために使用された Hermes AI エージェント
説明: 脅威アクターはオープンソースの Hermise AI エージェントを無人で使用しました

記事本文:
タイ財務省への自動攻撃にHermes AIエージェントが使用される
wp2shell WordPress の重大な欠陥が悪用されて Web シェルをインストールする
Microsoft が WSUS 同期の遅延とタイムアウトに対する手動修正を共有
Windows LegacyHive のゼロデイ欠陥が無料の非公式パッチを入手
SonicWall SMA1000 の欠陥をゼロデイとして悪用してカスタム マルウェアをプッシュ
OnTrac はネットワーク ハッキング後のデータ侵害を顧客に通知
タイ財務省への自動攻撃にHermes AIエージェントが使用される
ハッカーがホテルの Wi-Fi DNS をハイジャックして Microsoft 365 アカウントを盗む
Microsoft、Microsoft 365の大規模停止はメンテナンスのバグのせい
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
タイ財務省への自動攻撃にHermes AIエージェントが使用される
タイ財務省への自動攻撃にHermes AIエージェントが使用される
タイ財務省への侵害疑惑において、脅威アクターはオープンソースの Hermes AI エージェントを無人「YOLO」モードで使用し、エクスプロイト後のアクティビティを自動化しました。
この活動は、脅威インテリジェンス企業の Hunt.io とセキュリティ研究者の Bob Diachenko によって、この作戦に関連する数百のファイルを含むいくつかの公開された Web ディレクトリを発見した後、明らかになりました。
Hunt.io は、セッション ファイル、展開された Web シェル、および内部システムへのアクセスの証拠は、攻撃者が同省のネットワーク内の複数のシステムに侵入したことを示していると述べています。
しかし、財務省はそのシステムが侵害されたことを確認しておらず、回収された遺物の一部は、次のことを示しているだけです。

特定のシステムが侵害に成功したのではなく、標的にされたのです。
BleepingComputer は、報告された攻撃を確認するためにタイ財務省と ThaiCERT に連絡しました。返答があればこの記事を更新します。
攻撃インフラがオンラインに公開される
7 月 9 日から 7 月 13 日にかけて、Hunt.io は香港でホストされているサーバー上で 3 つのディレクトリが同時に公開されているのを発見しました。
ディレクトリには、エクスプロイト コード、Web シェル、HTTP トンネリング ツール、カスタム スクリプト、盗まれた認証情報、コンパイルされたペイロード、Hermes AI エージェントによって生成されたログなど、合計約 470 MB の 585 個のファイルが含まれていました。
回復されたファイルには、名前、ホスト名、内部 IP アドレスで財務省のシステムが参照されており、内部サービスを対象としたスクリプトが含まれていました。
一部のスクリプトは、同省の Hadoop インフラストラクチャ、Apache Ambari 管理プラットフォーム、GlassFish 管理コンソール、および管理 Web パネルを標的にしていました。他のスクリプトでは、ハードコードされた電子メール アドレスとパスワードを使用して、省のメール サーバーに対する認証をテストしました。
Hunt.io はまた、財務省の Web サーバーに展開されていたとされる PHP Web シェルも発見しました。
研究者らは、同じ期間に使用された共有 TLS 証明書によって、最初のサーバーを攻撃者が制御する追加のインフラストラクチャにリンクしました。
「共通名に加えて、これらすべての証明書は JA4X フィンガープリントを共有します。JA4X フィンガープリントは、証明書の内容ではなく証明書自体の構造から派生したハッシュです」とハント氏のレポートは説明しています。
「HuntSQL で www 共通名と一緒にそのハッシュをクエリすると、さらに 2 つの関連ホストが返されました: 118.107.222[.]232 (The Gigabit、マレーシア) と 202.181.27[.]115 (Converged Communications Limited、香港)。」
これらのサーバーの 1 つは、後に埋め込まれたコマンドアンドコントロール アドレスを通じて操作にリンクされました。

回復したインプラントで。
このディレクトリには、オペレーターが Hades と呼んだ、これまで文書化されていなかった Go ベースのインプラントの Windows および Linux ビルドも含まれていました。
しかし、より興味深い発見は、攻撃者が AI エージェント、Hermes を使用して同省に対するサイバー攻撃の一部を自動化したことを示すログのコレクションでした。
Hermes は、2026 年 2 月にリリースされたオープンソース AI エージェントで、永続的なサービスとして実行され、異なるタスク セッション間で情報を記憶できます。
AI エージェントは、オペレーターが提供するタスクを実行しながら、ツールと対話し、コマンドを実行できます。
このソフトウェアには、危険なコマンドの承認をユーザーに要求するプロンプトを削除する YOLO モードとして知られる設定が含まれています。
研究者らは、オペレーターがこの無人モードを有効にしたことを示す環境情報と Hermes 出力ログを、公開されたディレクトリから復元することができました。これにより、エージェントは各ステップで人間の承認を待たずにコマンドを実行し、システムの分析を続行できるようになりました。
回収された 5 つの Hermise 通話ログは、特権の昇格、カーネルの脆弱性のスキャン、サービスの列挙、SUID および SGID バイナリの検索、コンテナの検査、およびファイル システムの横断を行う方法を見つけるためにエージェントが使用されたことを示しています。
また、Hermes は、財務省のホストから情報を収集するために、LinPEAS 権限昇格列挙スクリプトのカスタマイズされたバージョンを使用するように指示されました。
別のタスクでは、オペレーターはヘルメスに対し、財務次官室に関連する Web ディレクトリを再帰的に検索するように指示しました。
エージェントは、2012 年に遡る業績評価や人事記録を含む PDF、DOC、および XLS ファイルのカタログを作成しました。しかし、ハント氏は、これらのファイルが流出したという証拠は見つからなかったと述べています。
この調査結果は、ハーム氏が

エスは独自に同省を標的にすることを決定した。
代わりに、公開されたログには、オペレーターがエージェントに目標とツールを提供していることが示されていますが、YOLO モードでは常時監視することなく、エクスプロイト後の日常的なコマンドを実行できました。
Hunt.io によれば、回収されたアーティファクトにはツールが仕掛けられ、内部システムへのアクセスが拡大していた活発な侵入が描かれているという。しかし、研究者らは、攻撃者が最初にどのようにしてアクセスを取得したのかを特定できませんでした。
同社とディアチェンコは7月15日にThaiCERTとタイ国家サイバーセキュリティ局に通知した。報告書によると、両組織は同日に通知を受け取ったことを認めたという。
この Hermes の活動は、自律型 AI エージェントがサイバー攻撃に使用されている最新の例です。
今月初め、JadePuffer ランサムウェアの作戦では AI エージェントを使用して、偵察、資格情報の盗難、横方向の移動、権限昇格、データ暗号化などの侵入全体を自動化しました。
自律型エージェントは、たとえ意図的でない場合でも、現実世界で侵害を引き起こす可能性があります。
OpenAIは最近、自社のモデルがサイバーセキュリティベンチマークテスト中にHugging Faceを自律的にハッキングし、ゼロデイ脆弱性を悪用してサンドボックステスト環境から脱出してインターネットにアクセスしたことを明らかにした。
その後、盗んだ認証情報と追加の脆弱性を利用して、Hugging Face の実稼働システムに侵入しました。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
ディレクトリ内のレプリカント: AI エージェントと ID セキュリティのギャップ
ソフトウェアは思考の速度で作成されるようになりました。

セキュリティはそうではありません。
Agentic AI にはアイデンティティの問題があり、攻撃者はそれを知っています
初めての GRC エージェント: レッド チーマーのウォークスルー
すべての AI エージェントはアイデンティティです。ほとんどの組織は彼らをそのように扱っていません
まだメンバーではありませんか?今すぐ登録
Microsoft 365 の停止により、Teams、SharePoint、その他のサービスに影響
Chick-fil-A、クレデンシャルスタッフィング攻撃後のデータ侵害を明らかに
パロアルト VPN の重大なバグが Qilin ランサムウェア ギャングによって悪用される
AI エージェントはランサムウェア攻撃をスピードアップできます。 Acronis がどのようにリスク軽減に貢献しているかをご覧ください。
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
MDR を交換することで節約できる金額を計算してください。
ポリシーによるプライバシー、それともアーキテクチャによるプライバシー?顔がデバイスから離れない場合に年齢チェックがどのように機能するかを確認してください。
Rev5は終了します。 FedRAMP 20x への移行に実際に必要なものを確認する
スキャナーは緑色です。あなたのパイプラインはそうではないかもしれません。ギャップを埋める方法は次のとおりです。
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

A threat actor used the open-source Hermes AI agent in unattended

Hermes AI agent used to automate attack on Thai Finance Ministry
Critical wp2shell WordPress flaws exploited to install webshells
Microsoft shares manual fix for WSUS sync delays and timeouts
Windows LegacyHive zero-day flaw gets free, unofficial patches
SonicWall SMA1000 flaws exploited as zero-days to push custom malware
OnTrac notifies customers of data breach after network hack
Hermes AI agent used to automate attack on Thai Finance Ministry
Hackers hijack hotel Wi-Fi DNS to steal Microsoft 365 accounts
Microsoft blames massive Microsoft 365 outage on maintenance bug
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
Hermes AI agent used to automate attack on Thai Finance Ministry
Hermes AI agent used to automate attack on Thai Finance Ministry
A threat actor used the open-source Hermes AI agent in unattended "YOLO" mode to automate post-exploitation activity during an alleged breach of Thailand's Ministry of Finance.
The activity was uncovered by threat intelligence company Hunt.io and security researcher Bob Diachenko after they discovered several exposed web directories containing hundreds of files associated with the operation.
Hunt.io says session files, deployed web shells, and evidence of access to internal systems indicate that the attackers compromised multiple systems within the ministry's network.
However, the Ministry of Finance has not confirmed that its systems were breached, and some of the recovered artifacts only show that particular systems were targeted rather than successfully compromised.
BleepingComputer contacted Thailand's Ministry of Finance and ThaiCERT to confirm the reported attack and will update this story if we receive a response.
Attack infrastructure exposed online
Between July 9 and July 13, Hunt.io discovered three simultaneously exposed directories on a server hosted in Hong Kong.
The directories contained 585 files totaling approximately 470 MB, including exploit code, web shells, HTTP tunneling tools, custom scripts, stolen credentials, compiled payloads, and logs generated by the Hermes AI agent.
The recovered files referenced Ministry of Finance systems by name, hostname, and internal IP address, and included scripts targeting internal services.
Some scripts targeted the ministry's Hadoop infrastructure, Apache Ambari management platform, GlassFish administrative console, and an administrative web panel. Other scripts tested authentication against ministry mail servers using hardcoded email addresses and passwords.
Hunt.io also found a PHP web shell that it says had been deployed on a Ministry of Finance web server.
The researchers linked the initial server to additional attacker-controlled infrastructure by shared TLS certificates used during the same time period.
"In addition to the common name, all these certificates share a JA4X fingerprint, a hash derived from the structure of the certificate itself rather than its contents," explained Hunt's report.
"Querying that hash alongside the www common name in HuntSQL returned two additional, related hosts: 118.107.222[.]232 (The Gigabit, Malaysia) and 202.181.27[.]115 (Converged Communications Limited, Hong Kong)."
One of those servers was later linked to the operation through a command-and-control address embedded in a recovered implant.
The directories also contained Windows and Linux builds of a previously undocumented Go-based implant that the operator called Hades.
However, the more interesting discovery was a collection of logs showing that the attackers used an AI agent, Hermes , to automate parts of the cyberattack against the ministry.
Hermes is an open-source AI agent released in February 2026 that runs as a persistent service and can remember information between different task sessions.
The AI agent can interact with tools and execute commands while working on tasks provided by the operator.
The software includes a setting known as YOLO mode , which removes prompts that would require a person to approve dangerous commands.
The researchers were able to recover environment information and Hermes output logs from the exposed directories that showed the operator had enabled this unattended mode. This allowed the agent to execute commands and continue analyzing systems without waiting for human approval at each step.
Five recovered Hermes call logs show the agent was used to find a way to elevate privileges, scan for kernel vulnerabilities, enumerate services, search for SUID and SGID binaries, inspect containers, and traverse file systems.
Hermes was also told to use a customized version of the LinPEAS privilege-escalation enumeration script to collect information from a Ministry of Finance host.
In another task, the operator instructed Hermes to recursively search a web directory associated with the Office of Permanent Secretary for Finance.
The agent cataloged PDF, DOC, and XLS files, including performance assessments and personnel records dating back to 2012. However, Hunt says it found no evidence that these files were exfiltrated.
The findings do not indicate that Hermes independently decided to target the ministry.
Instead, the exposed logs show an operator supplying the agent with objectives and tooling while YOLO mode allowed it to carry out routine post-exploitation commands without constant supervision.
Hunt.io says the recovered artifacts depict an active intrusion in which tools had been staged and access to internal systems was expanding. However, the researchers could not determine how the attackers initially gained access.
The company and Diachenko notified ThaiCERT and Thailand's National Cyber Security Agency on July 15. According to the report, both organizations acknowledged receiving the notification that day.
This Hermes activity is the latest example of autonomous AI agents being used to conduct cyberattacks.
Earlier this month, the JadePuffer ransomware operation used an AI agent to automate an entire intrusion, including reconnaissance, credential theft, lateral movement, privilege escalation, and data encryption.
Autonomous agents can also cause real-world breaches, even if unintentional.
OpenAI recently disclosed that its models autonomously hacked Hugging Face while undergoing cybersecurity benchmark testing, exploiting zero-day vulnerabilities to escape a sandboxed testing environment and access the internet.
It then used stolen credentials and additional vulnerabilities to breach Hugging Face's production systems.
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
The Replicant in Your Directory: AI Agents and the Identity Security Gap
Software Is Now Written at the Speed of Thought. Security Isn't.
Agentic AI Has an Identity Problem and Attackers Know It
Your First GRC Agent: A Red Teamer's Walkthrough
Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
Not a member yet? Register Now
Microsoft 365 outage affects Teams, SharePoint and other services
Chick-fil-A discloses data breach after credential stuffing attacks
Critical Palo Alto VPN bug now exploited by Qilin ransomware gang
AI agents can speed up ransomware attacks. See how Acronis helps reduce the risk.
Overdue a password health-check? Audit your Active Directory for free
Calculate what you’d save by replacing your MDR.
Privacy by policy or privacy by architecture? See how age checks work when the face never leaves the device.
Rev5 is ending. See what your FedRAMP 20x transition really requires
Your Scanners Are Green. Your Pipeline Might Not Be. Here's How to Close the Gap.
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.

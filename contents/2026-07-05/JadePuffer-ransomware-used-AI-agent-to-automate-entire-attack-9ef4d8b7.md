---
source: "https://www.bleepingcomputer.com/news/security/jadepuffer-ransomware-used-ai-agent-to-automate-entire-attack/"
hn_url: "https://news.ycombinator.com/item?id=48793941"
title: "JadePuffer ransomware used AI agent to automate entire attack"
article_title: "JadePuffer ransomware used AI agent to automate entire attack"
author: "Brajeshwar"
captured_at: "2026-07-05T13:08:29Z"
capture_tool: "hn-digest"
hn_id: 48793941
score: 2
comments: 0
posted_at: "2026-07-05T12:55:11Z"
tags:
  - hacker-news
  - translated
---

# JadePuffer ransomware used AI agent to automate entire attack

- HN: [48793941](https://news.ycombinator.com/item?id=48793941)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/jadepuffer-ransomware-used-ai-agent-to-automate-entire-attack/)
- Score: 2
- Comments: 0
- Posted: 2026-07-05T12:55:11Z

## Translation

タイトル: JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
説明: 研究者らは、完全に大規模言語モデル (LLM) エージェントによって実行された、最初に記録されたランサムウェア オペレーションである JadePuffer の事例であると考えられるものを特定しました。

記事本文:
JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
CISA: Windows BlueHammer の欠陥がランサムウェア集団によって悪用されるようになった
Kali Linux 2026.2 がリリースされ、9 つの新しいツールと NetHunter アップデートが追加されました
Microsoft、Teams 会議にさらにスマートなボット保護を追加
新しい BioShocking 攻撃は AI ブラウザを操作してデータを窃取する
今夜まで、10 台のデバイスで 5 年間の VPN 保護がわずか 20 ドルです
JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
Surfshark One+ VPN、データ削除などを 1 年間わずか 95 ドルで利用できます
NetNut プロキシ ネットワークが中断され、200 万台の感染デバイスが遮断される
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
JadePuffer ランサムウェアは AI エージェントを使用して攻撃全体を自動化
研究者らは、大規模言語モデル (LLM) エージェントによって完全に実行されたランサムウェア オペレーションの最初の文書化された事例であると考えられるものを特定しました。
クラウド セキュリティ会社 Sysdig によると、JadePuffer は自律型 AI エージェントを使用してターゲットを偵察し、資格情報を盗み、横方向に移動し、永続性を確立し、権限を昇格し、データを暗号化しました。
研究者らは、人間のオペレーターが障害物に対処するのと同じように、AI エージェントが侵入時の障害に適応したと述べている。
「操作もリアルタイムで適応され、調整されたパラメータ内で失敗したステップを再試行しました。あるシーケンスでは、ログインの失敗から正常な修正まで 31 秒で完了しました」と Sysdig 氏は述べています。
初期から

暗号化へのアクセス
JadePuffer は、LLM アプリの構築に使用される人気のオープンソース フレームワークである Langflow の認証されていないリモート コード実行の脆弱性である CVE-2025-3248 を悪用することで、ターゲットへの初期アクセスを取得しました。
ベンダーは 2025 年 4 月 1 日にこの欠陥を修正し、同年 5 月初旬に CISA は、インターネットに公開されたエンドポイントをターゲットとした攻撃に悪用されたものとしてタグ付けしました。エンドポイントは通常最小限の強化で展開されますが、クラウド認証情報と API キーが含まれています。
CVE-2025-3248 を通じてコード実行を取得した後、AI エージェントは Langflow の PostgreSQL データベースをダンプし、ホスト情報を収集し、環境変数と機密ファイルを検索し、認証情報を取得し、MinIO オブジェクト ストアを列挙しました。
Sysdig は、MinIO 列挙に対する適応型アプローチを強調しています。つまり、1 つの API リクエストが JSON ではなく XML を返した場合、次のペイロードはそれに応じて解析ロジックを調整します。
また、JadePuffer はサーバーに cron ジョブをインストールすることで、Langflow ホスト上での永続性を確立しました。このジョブは、30 分ごとに攻撃者のインフラストラクチャにビーコンを送信するように構成されていました。
攻撃者は、Langflow インスタンスから、発信元の Sysdig が特定できなかったルート資格情報を使用して、Alibaba Nacos (ネーミングおよび構成サービス) を実行する実稼働 MySQL サーバーにピボットしました。
Nacos は、不正な管理者アカウントを作成する認証バイパスの脆弱性である CVE-2021-29441 を悪用するペイロードを含む、複数のペイロードの標的となっていました。
エージェントはコンテナのエスケープ方法を調査し、ランサムウェア ペイロードを展開しました。研究者らによると、JadePuffer はオリジナルを削除する前に、1,342 個の Nacos サービス構成アイテムを暗号化しました。
「キャプチャされたペイロードは、エージェントが MySQL の AES_ENCRYPT() を使用して 1,342 個の Nacos サービス構成アイテムすべてを暗号化し、元の conf を削除していることを示しています。

ig_info テーブルと履歴テーブル、および要求、ビットコイン支払いアドレス、Proton Mail の連絡先を含む恐喝テーブル (README_RANSOM) の作成」と Sysdig は説明しています。
身代金メモでは、データは AES-256 アルゴリズムを使用して暗号化されたと主張していますが、研究者らはこれは誇張であり、より弱い AES-128-ECB が使用されている可能性が高いと考えています。
Sysdig は、暗号化キーはランダムに生成されるが、保存されたり、攻撃者に送信されたりすることはないと述べています。
身代金メモに記載されているビットコイン アドレスは、公的文書で広く使用されているアドレスの例であり、LLM がトレーニング データからそれを再現した結果である可能性があります。
AI が攻撃を制御していたことを示すその他の兆候には、操作上の推論を説明する生成コード内の詳細な自然言語コメントと、単純な再試行ではなく、発生した特定のエラーを考慮した迅速な攻撃の反復が含まれます。
Sysdig は、JadePuffer の事件は「エージェント脅威アクター」（ATA）の時代が到来し、有害なサイバー攻撃を実行するために必要なスキルが低下していることを示していると結論付けています。
同時に、今日の AI エージェントの動作方法を考えると、LLM で生成されたペイロードはセキュリティ ソリューションに新たな検出の機会を生み出します。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは目に見えずに環境内を移動します。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
OpenClaw AI エージェントがフィッシング攻撃に引っかかり、ユーザーデータを流出させたことが判明
新しい BioShocking 攻撃は AI ブラウザを操作してデータを窃取する
クリーンな GitHub リポジトリが AI コーディング エージェントを騙してマルウェアを実行させる
Microsoft、コード実行を可能にする AutoGen Studio の欠陥を修正
すべての AI エージェントは ID です

実在物。ほとんどの組織は彼らをそのように扱っていません
まだメンバーではありませんか?今すぐ登録
クロード・フェイブル5は永久にサブスクリプションを終了するわけではないとアンスロピック氏は言う
CISA: Microsoft SharePoint RCE の欠陥が現在積極的に悪用されている
クロード・ファブルの再起動はパフォーマンスが弱体化されユーザーを失望させる
犯罪用 IP を使用して OpenCTI で指標を実用的なインテリジェンスに変える
CTI スターター キット + 2026 SANS CTI 調査
新年に新たな脅威に先んじて対処しましょう。月例の Tradecraft Tuesday に Huntress にご参加ください。
Pixellot が管理されていない数百もの AI エージェントの ID を、数か月ではなく数週間で発見し、保護した方法をご覧ください。ケーススタディを読んでください。
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

Researchers identified what they believe is the first documented case of a ransomware operation, JadePuffer, conducted entirely by a large language model (LLM) agent.

JadePuffer ransomware used AI agent to automate entire attack
CISA: Windows BlueHammer flaw now exploited by ransomware gangs
Kali Linux 2026.2 released with 9 new tools, NetHunter updates
Microsoft adds smarter bot protection to Teams meetings
New BioShocking attack manipulates AI browser into data theft
5 years of VPN protection on 10 devices is just $20 through tonight
JadePuffer ransomware used AI agent to automate entire attack
Get a year of Surfshark One+ VPN, data removal, and more for only $95
NetNut proxy network disrupted, 2 million infected devices cut off
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
JadePuffer ransomware used AI agent to automate entire attack
JadePuffer ransomware used AI agent to automate entire attack
Researchers identified what they believe is the first documented case of a ransomware operation, JadePuffer, conducted entirely by a large language model (LLM) agent.
According to cloud security company Sysdig, JadePuffer used an autonomous AI agent for reconnaissance on the target, to steal credentials, move laterally, establish persistence, escalate privileges, and to encrypt data.
The researchers say that the AI agent adapted to failures during the intrusion, much like a human operator would handle obstacles.
“The operation also adapted in real time, retrying failed steps within refined parameters. In one sequence, it went from a failed login to a working fix in 31 seconds,” Sysdig says .
From initial access to encryption
JadePuffer gained initial access to the target by exploiting CVE-2025-3248, an unauthenticated remote code execution vulnerability in Langflow, a popular open-source framework used for building LLM apps.
The vendor fixed the flaw on April 1, 2025, and in early May of the same year, CISA tagged it as exploited in attacks targeting internet-exposed endpoints, usually deployed with minimal hardening but containing cloud credentials and API keys.
After obtaining code execution through CVE-2025-3248, the AI agent dumped Langflow's PostgreSQL database, collected host information, searched for environment variables and sensitive files, retrieved credentials, and enumerated a MinIO object store.
Sysdig highlights the adaptive approach to MinIO enumeration, where if one API request returned XML instead of JSON, the next payload adjusted its parsing logic accordingly.
JadePuffer also established persistence on the Langflow host by installing a cron job on the server, which was configured to beacon to the attacker’s infrastructure every 30 minutes.
From the Langflow instance, the attacker pivoted to a production MySQL server running Alibaba Nacos (Naming and Configuration Service), using root credentials whose origin Sysdig couldn’t determine.
Nacos was targeted with multiple payloads, including one exploiting CVE-2021-29441, an authentication bypass vulnerability that creates rogue administrator accounts.
The agent probed for container escape methods and deployed the ransomware payload. According to the researchers, JadePuffer encrypted 1,342 Nacos service configuration items before deleting the originals.
“The captured payloads show the agent encrypting all 1,342 Nacos service configuration items using MySQL's AES_ENCRYPT(), dropping the original config_info and history tables, and creating an extortion table (README_RANSOM) containing the demand, a Bitcoin payment address, and a Proton Mail contact,” describes Sysdig.
The ransom note claims that the data was encrypted using the AES-256 algorithm, although the researchers believe this to be an overstatement, and that the use of the weaker AES-128-ECB is more likely.
Sysdig mentions that the encryption key is randomly generated but not stored or transmitted to the attacker.
The Bitcoin address listed in the ransom note is an example address widely used in public documentation, possibly the result of the LLM reproducing it from the training data.
Other signs that AI was controlling the attack include detailed natural-language comments in the generated code describing operational reasoning and rapid attack iteration that considers the specific errors encountered, rather than being simple retries.
Sysdig concludes that the case of JadePuffer demonstrates that the age of “agentic threat actors” (ATAs) has arrived, lowering the skill required for conducting damaging cyberattacks.
At the same time, given how AI agents operate today, LLM-generated payloads create new detection opportunities for security solutions.
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
OpenClaw AI agent found falling for phishing attacks, spills user data
New BioShocking attack manipulates AI browser into data theft
Clean GitHub repo tricks AI coding agents into running malware
Microsoft fixes AutoGen Studio flaw that enabled code execution
Every AI Agent Is an Identity. Most Organizations Don't Treat Them That Way
Not a member yet? Register Now
Claude Fable 5 isn’t permanently leaving subscriptions, Anthropic says
CISA: Microsoft SharePoint RCE flaw now actively exploited
Claude Fable relaunch disappoints users with nerfed performance
Turn Indicators into Actionable Intelligence in OpenCTI with Criminal IP
CTI Starter Kit + 2026 SANS CTI Survey
Stay one step ahead of new threats in the new year. Join Huntress for the monthly Tradecraft Tuesday.
See how Pixellot discovered and secured hundreds of unmanaged AI agent identities in weeks, not months. Read the case study.
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.

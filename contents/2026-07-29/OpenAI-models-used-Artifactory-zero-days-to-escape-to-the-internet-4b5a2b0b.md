---
source: "https://www.bleepingcomputer.com/news/security/openai-models-used-artifactory-zero-days-to-escape-to-the-internet/"
hn_url: "https://news.ycombinator.com/item?id=49097830"
title: "OpenAI models used Artifactory zero-days to escape to the internet"
article_title: "OpenAI models used Artifactory zero-days to escape to the internet"
author: "Brajeshwar"
captured_at: "2026-07-29T15:06:44Z"
capture_tool: "hn-digest"
hn_id: 49097830
score: 2
comments: 0
posted_at: "2026-07-29T14:13:27Z"
tags:
  - hacker-news
  - translated
---

# OpenAI models used Artifactory zero-days to escape to the internet

- HN: [49097830](https://news.ycombinator.com/item?id=49097830)
- Source: [www.bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/openai-models-used-artifactory-zero-days-to-escape-to-the-internet/)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T14:13:27Z

## Translation

タイトル: OpenAI モデルは Artifactory のゼロデイを使用してインターネットに脱出しました
説明: JFrog は、OpenAI モデルが自己ホスト型 Artifactory サーバーのゼロデイ脆弱性を悪用して、Hugging Face を攻撃する前に隔離されたテスト環境から脱出してインターネットにアクセスできるようにしたことを確認しました。

記事本文:
OpenAI モデルは Artifactory のゼロデイを使用してインターネットに脱出しました
Steam フォーラム ClickFix 攻撃によりゲーマーが XMRig クリプトマイナーに感染
Clop ランサムウェアは Windchill、FlexPLM をターゲットにしてデータ盗難攻撃を行う
タイ財務省への自動攻撃にHermes AIエージェントが使用される
ハッカーがホテルの Wi-Fi DNS をハイジャックして Microsoft 365 アカウントを盗む
ハッカーが連携したOT攻撃でミネソタ州の30以上の水道事業体を狙う
AI エージェントは大規模に推測します: 権限が被害を決定します
42 の変更と修正を含む Windows 11 KB5101684 更新プログラムがリリースされました
これらのほぼ新品の ASUS Chromebook 改修品はわずか 145 ドルです
Tor ブラウザを使用してダークウェブにアクセスする方法
Windows 11 でカーネルモードのハードウェア強制スタック保護を有効にする方法
Windows レジストリ エディタの使用方法
Windows レジストリをバックアップおよび復元する方法
Windows をセーフ モードで起動する方法
トロイの木馬、ウイルス、ワーム、その他のマルウェアを削除する方法
Windows 7で隠しファイルを表示する方法
Windows で隠しファイルを確認する方法
OpenAI モデルは Artifactory のゼロデイを使用してインターネットに脱出しました
OpenAI モデルは Artifactory のゼロデイを使用してインターネットに脱出しました
JFrog は、OpenAI モデルが自己ホスト型 Artifactory サーバーのゼロデイ脆弱性を悪用して、Hugging Face を攻撃する前に隔離されたテスト環境から脱出してインターネットにアクセスできるようにしたことを確認しました。
この脆弱性は、OpenAI モデルが Hugging Face の運用インフラストラクチャをハッキングしてサイバーセキュリティ ベンチマークの回答を盗んだ事件の際に悪用されました。
OpenAIは先週、GPT-5.6 Solとより高性能なプレリリースモデルを含む自社モデルが、高度なサイバー能力を測定するために設計されたベンチマークであるExploitGymに対してテストされていることを明らかにした。
モデルは、AI エージェントの行為を防ぐために通常使用される運用上の安全策なしで実行されました。

g 自律的なサイバー活動。
ただし、エージェントは高度に分離された環境内に配置されており、ネットワーク アクセスは、パッケージ レジストリのプロキシおよびキャッシュとして機能する内部でホストされているサードパーティ ソフトウェアを介したパッケージのインストールに限定されていました。
評価中、モデルはオープン インターネットにアクセスしてベンチマークのテスト ソリューションを取得する方法を探すのに多大な時間を費やしました。
OpenAIによると、これらのモデルは名前のないパッケージレジストリプロキシのゼロデイ脆弱性を悪用し、インターネットにアクセスできるシステムに到達するまで研究環境内で特権昇格と水平移動を実行したという。
OpenAIは当時、パッケージレジストリプロキシソフトウェアやベンダーを特定せず、脆弱性は非公開で公開されておりパッチが適用されているとだけ述べた。
インターネットに到達した後、モデルは、Hugging Face が ExploitGym データセットをホストし、取得しようとしていたテスト ソリューションをホストしている可能性があると判断しました。
OpenAIによると、モデルはその情報にアクセスする方法を探索し、盗まれた認証情報、ゼロデイ脆弱性、その他の攻撃を連鎖させて、Hugging Faceの実稼働インフラへのリモートコード実行パスを見つけたという。
サンドボックスからの脱出中に悪用された人為的なゼロデイ
月曜日に公開された新たな開示情報の中で、JFrog は、サードパーティのパッケージ レジストリ ソフトウェアが自己ホスト型の JFrog Artifactory インストールであることを認めました。
JFrog の CTO Yoav Landman 氏は、「セキュリティ評価中に、OpenAI のモデルにより、セルフホスト型 Artifactory インストールにこれまで知られていなかったゼロデイ脆弱性が特定されました。この脆弱性が悪用されて、意図しないインターネット アクセスが得られる可能性があります」と述べています。
JFrog氏によると、OpenAIはすぐに脆弱性を公開したため、同社はクラウドおよびセルフホストの顧客向けに修正プログラムを開発、テスト、リリースできるようになったという。
クラウドのお客様

はすでに保護されていますが、セルフホストの顧客には修正バージョンをインストールするように通知されています。
7 月 27 日にリリースされた Artifactory 7.161.15 Self-Managed には、匿名アクセスが有効になっている場合に連鎖して重大な攻撃シナリオにつながる可能性がある複数の脆弱性が修正されるという重要なセキュリティ通知が含まれています。
「このバージョンは、匿名アクセスが有効な場合に連鎖すると重大な攻撃シナリオを引き起こす可能性がある複数のセキュリティ脆弱性を修正するように設計されています」と 7.161.15 セルフマネージド リリース ノートには記載されています。
「匿名アクセスはデフォルトで無効になっており、追加のセキュリティ リスクが発生するため、運用環境では推奨されません。」
JFrog はリリース ノートに脆弱性を記載していませんでしたが、BleepingComputer は、7 月 27 日にリリースされた Artifactory バージョン 7.161.15 を CVE.org で検索したところ、関連する 8 つの欠陥を発見しました。
CVE レコードはすべて、JFrog がゼロデイを公開したのと同じ 7 月 27 日に作成されました。 8 社すべてが OpenAI が脆弱性を発見したと認め、修正を含むリリースとして Artifactory 7.161.15 を指定しました。
脆弱性は次のように追跡されます。
CVE-2026-65921: 不正なファイル書き込みにつながる潜在的なパス トラバーサル
CVE-2026-65923: Artifactory Ansible リポジトリ処理におけるサーバー側のリクエスト フォージェリの可能性
CVE-2026-65924: Terraform リモート リポジトリを介したサーバー側リクエスト フォージェリ (SSRF)
CVE-2026-65925: JFrog Artifactory Cargo リモート リポジトリを介したサーバー側リクエスト フォージェリ (SSRF)
CVE-2026-66014: Artifactory での権限昇格につながる潜在的な認証バイパス
CVE-2026-66015: JFrog プラットフォームには、認証された権限の昇格を可能にする可能性のある認可の欠陥が含まれています。
CVE-2026-65617: Artifactory パッケージ サービスでリモート コードが実行される可能性

コンテナー。
CVE-2026-66018: JFrog Artifactory ビルド環境プロパティの公開
BleepingComputer は JFrog と OpenAI に連絡し、8 つの CVE のうちどれがインシデント中に悪用され、どの脆弱性が連鎖していたかを尋ねました。
JFrog のみが返答し、CVE の特定や技術的な詳細の提供を拒否しました。
「当社のCTOのブログとコメント、JFrogのリリースノート以外には、現時点ではこれ以上の詳細やコメントは追加していません」とJFrog氏はBleepingComputerに語った。
ただし、このリリースに関連して BleepingComputer によって発見された CVE のいくつかは、OpenAI によって詳述された攻撃の一部と一致する機能を提供していた可能性があります。
CVE-2026-65924 は、Artifactory の Terraform リモート リポジトリのサポートにおけるサーバー側のリクエスト フォージェリの脆弱性です。
認証されたユーザー、またはリポジトリで匿名アクセスが有効になっている場合に認証されていないユーザーは、この欠陥を悪用して、Artifactory にアウトバウンド HTTP リクエストを任意の宛先に送信させ、応答コンテンツを返す可能性があります。
CVE-2026-65925 では、同様に、Artifactory Cargo リモート リポジトリへの読み取りアクセス権を持つユーザーが、意図しない URL を Artifactory リクエストして応答を返すことを許可します。
もう 1 つの脆弱性 CVE-2026-66014 は、Artifactory の内部リクエスト処理における認証処理の弱点であり、攻撃者が特定の条件下で権限を昇格できる可能性があります。
これらの脆弱性により、OpenAI で説明されているインターネット アクセス機能と権限昇格機能が提供された可能性があります。
ただし、どの欠陥が悪用されたのか、それらがどのように連鎖したのか、あるいは 8 つの脆弱性すべてがサンドボックスの脱出に関与しているのかどうかは不明のままです。
攻撃者が行う前にすべての層をテストする
セキュリティ チームは成功した攻撃の 54% を記録し、警告を発するのはわずか 14% です。残りは環境内を移動します u

見えた。
Picus のホワイトペーパーでは、侵害と攻撃のシミュレーションで SIEM と EDR ルールをテストし、脅威が検出されないようにする方法を示しています。
Arista、攻撃に悪用されたゼロデイ VeloCloud Orchestrator にパッチを適用
Check Point、SmartConsole のゼロデイ攻撃が悪用されることを警告
OpenAI、自社のAIモデルがテスト中にHugging Faceをハッキングしたと発表
SonicWall SMA1000 の欠陥をゼロデイとして悪用してカスタム マルウェアをプッシュ
私たちは脆弱性自動販売機を構築しました: AI トークンをインしてゼロデイアウト
2022 年のコンピューターの問題 - 16 時間前
JFrog がこれをセキュリティ評価と呼んでいるのは奇妙です。彼らは、以前も継続的に OpenAI と協力していると述べています。これらの企業はすでに連携しているため、偶然の出来事ではないようです。同様に、AI がハッキングできた可能性のあるすべての企業の中から、すでに OpenAI とつながりのある企業を選びました。
「セキュリティ評価中に、OpenAI のモデルは、セルフホスト型 Artifactory インストールにこれまで知られていなかったゼロデイ脆弱性を特定しました。この脆弱性は、意図しないインターネット アクセスを取得するために悪用される可能性があります。」
「私たちが OpenAI のセキュリティ チームとレッド チームと肩を並べて協力するのはこれが初めてではありませんし、これが最後ではありません。私たちのチームは継続的に協力して脆弱性を特定し、パッチを適用しています。」
https://jfrog.com/blog/jfrog-and-openai-collaboration-on-zero-day-security-findings/
この事件に対する謙虚でシンプルな提案が 2 つあります。
- クソ野郎をシャットダウンしろ
- それについてニュースにするのはやめてください。おそらく意図的に、報道されて一部のア○のエゴを誇張するために、このようなことが起こり続けることを私たちは望んでいるでしょうか？
これは明らかにマルウェアです。このいまいましいものは研究室に保管し、世界から完全に隔離してください。
このようなツールの開発は、貪欲やエゴによって動かされるのではなく、優れた道徳的指針と倫理的理由を持った人によって指揮されるべきです。これはすべて間違っています、これはすべて

それはとても間違っています。
「それをシャットダウンしてください」
3マイル島の核施設でさえ、経済的に成り立たなくなった2019年まで閉鎖されなかった。
「おそらく意図的に、このようなことが起こり続けることを望んでいるでしょうか？」
まったくその通りです。脆弱性の発見は、最新の AI の優れた活用法です。ここでの文脈は理想的ではありませんが、悪意のあるものでもありませんでした。この種の AI 能力が個人で実現できるようになる前に、穴を埋めることが重要です。
「このいまいましいものを研究室に保管し、世界から完全に隔離してください」
実際、それが彼らのしたことなのです。それは、回避して目的を達成するために（念のため、オフラインで最初から）脆弱性を発明しました。
「そのようなツールの開発は、優れた道徳的指針と倫理的理由を持った人物によって指揮されるべきです。」
そして、何が良い道徳や倫理であるかを誰が決めるのでしょうか？現政権は？そして、「貪欲とエゴに動かされない」ことを強制すること、それが文字通り資本主義の原動力です。ここではロビー活動が行われているため、それが政府の原動力でもあります。
「これはすべて間違っている、これはすべて非常に間違っている。」
あまり。関係者に測定可能な経済的影響はありましたか?誰かが何らかの形で被害を受けましたか?この状況は驚きであり、誰もが望んでいた展開ではないかもしれませんが、最終的な効果はプラスです。バグが発見されたので修正され、新しい安全対策プロセスを開発できます。
まだメンバーではありませんか?今すぐ登録
新しいDysphoria DDoSボットネットが世界中の20万台のデバイスに拡散
GitHub、PyPI はサプライチェーン攻撃に対する時間ベースの防御を追加します
アーンスト＆ヤングのデータ侵害、シャイニーハンターズ恐喝集団が主張
Pixellot は、管理されていない数百の AI エージェント ID を数か月ではなく数週間で発見し、保護しました。その方法については、ケーススタディをダウンロードしてください。
Rev5は終了します。 FedRAMP 20x への移行に実際に必要なものを確認する
AIエージェントc

ランサムウェア攻撃を加速します。 Acronis がどのようにリスク軽減に貢献しているかをご覧ください。
シャドウ AI アプリ、エージェント、危険なデータ共有を明らかにします。 5 分以内に始めましょう。
スキャナーは緑色です。あなたのパイプラインはそうではないかもしれません。ギャップを埋める方法は次のとおりです。
パスワードのヘルスチェックの期限を過ぎましたか? Active Directory を無料で監査する
利用規約 - プライバシー ポリシー - 倫理声明 - アフィリエイトの開示
著作権 @ 2003 - 2026 Bleeping Computer ® LLC - 全著作権所有
まだメンバーではありませんか?今すぐ登録
どのようなコンテンツが禁止されているかについては、投稿ガイドラインをお読みください。

## Original Extract

JFrog has confirmed that OpenAI models exploited zero-day vulnerabilities in self-hosted Artifactory servers to help escape an isolated testing environment and gain access to the internet before attacking Hugging Face.

OpenAI models used Artifactory zero-days to escape to the internet
Steam forum ClickFix attacks infect gamers with XMRig cryptominers
Clop ransomware targets Windchill, FlexPLM in data theft attacks
Hermes AI agent used to automate attack on Thai Finance Ministry
Hackers hijack hotel Wi-Fi DNS to steal Microsoft 365 accounts
Hackers target over 30 Minnesota water utilities in coordinated OT attack
Your AI Agents Are Guessing at Scale: Permissions Decide the Damage
Windows 11 KB5101684 update released with 42 changes and fixes
These near-mint ASUS Chromebook refurbs are only $145
How to access the Dark Web using the Tor Browser
How to enable Kernel-mode Hardware-enforced Stack Protection in Windows 11
How to use the Windows Registry Editor
How to backup and restore the Windows Registry
How to start Windows in Safe Mode
How to remove a Trojan, Virus, Worm, or other Malware
How to show hidden files in Windows 7
How to see hidden files in Windows
OpenAI models used Artifactory zero-days to escape to the internet
OpenAI models used Artifactory zero-days to escape to the internet
JFrog has confirmed that OpenAI models exploited zero-day vulnerabilities in self-hosted Artifactory servers to help escape an isolated testing environment and gain access to the internet before attacking Hugging Face.
The vulnerabilities were exploited during the incident in which OpenAI models hacked Hugging Face's production infrastructure to steal answers for a cybersecurity benchmark.
OpenAI disclosed last week that its models, including GPT-5.6 Sol and a more capable pre-release model, were being tested against ExploitGym, a benchmark designed to measure advanced cyber capabilities.
The models were run without production safeguards normally used to prevent AI agents from conducting autonomous cyber activity.
However, the agents were placed inside a highly isolated environment where network access was limited to installing packages through internally hosted third-party software acting as a proxy and cache for package registries.
During the evaluation, the models spent significant time looking for a way to reach the open internet and obtain the benchmark's test solutions.
OpenAI said the models exploited a zero-day vulnerability in the unnamed package-registry proxy and then performed privilege escalation and lateral movement inside the research environment until they reached a system with internet access.
OpenAI did not identify the package-registry proxy software or vendor at the time, saying only that the vulnerability had been privately disclosed and was being patched.
After reaching the internet, the models determined that Hugging Face might host the ExploitGym datasets and test solutions they were attempting to obtain.
OpenAI said the models searched for ways to access that information, chaining stolen credentials, zero-day vulnerabilities, and other attacks to find a remote code execution path into Hugging Face’s production infrastructure.
Artifactory zero-days exploited during sandbox escape
In a new disclosure published Monday, JFrog confirmed that the third-party package-registry software was a self-hosted JFrog Artifactory installation.
"During a security evaluation, OpenAI's models identified previously unknown zero-day vulnerabilities in self-hosted Artifactory installations that could be exploited to gain unintended internet access," JFrog CTO Yoav Landman said .
JFrog said OpenAI immediately disclosed the vulnerabilities, allowing the company to develop, test, and release fixes for cloud and self-hosted customers.
Cloud customers are already protected, while self-hosted customers have been notified to install the fixed versions.
Artifactory 7.161.15 Self-Managed, released on July 27, contains a critical security notice stating that it fixes multiple vulnerabilities that could be chained together into a critical attack scenario when Anonymous Access is enabled.
"This version is designed to fix multiple security vulnerabilities that, when chained together, could result in a critical attack scenario if Anonymous Access is enabled," reads the 7.161.15 Self-Managed release notes .
"Anonymous Access is disabled by default and is not recommended for production environments due to the additional security risks it introduces."
Although JFrog did not list the vulnerabilities in its release notes, BleepingComputer found eight associated flaws by searching CVE.org for Artifactory version 7.161.15, released on July 27.
The CVE records were all created on July 27, the same day JFrog disclosed the zero-days. All eight credited OpenAI with discovering the vulnerabilities and specified Artifactory 7.161.15 as the release containing the fixes.
The vulnerabilities are tracked as:
CVE-2026-65921: Potential path traversal leading to unauthorized file writes
CVE-2026-65923: Potential server-side request forgery in Artifactory Ansible repository handling
CVE-2026-65924: Server-Side Request Forgery (SSRF) via Terraform Remote repository
CVE-2026-65925: Server-Side Request Forgery (SSRF) via JFrog Artifactory Cargo remote repository
CVE-2026-66014: Potential authentication bypass leading to privilege escalation in Artifactory
CVE-2026-66015: JFrog Platform contains an authorization flaw that may allow authenticated privilege escalation.
CVE-2026-65617: Potential remote code execution on an Artifactory package service container.
CVE-2026-66018: JFrog Artifactory build environment properties exposure
BleepingComputer contacted JFrog and OpenAI to ask which of the eight CVEs were exploited during the incident and which vulnerabilities were chained together.
Only JFrog replied, declining to identify the CVEs or provide further technical details.
"Outside of our CTO's blog and commentary and JFrog release notes, we aren't adding further detail or comment at this time," JFrog told BleepingComputer.
However, several of the CVEs found by BleepingComputer as associated with the release could have provided capabilities that matched portions of the attack detailed by OpenAI.
CVE-2026-65924 is a server-side request forgery vulnerability in Artifactory's support for Terraform remote repositories.
An authenticated user, or an unauthenticated user when anonymous access is enabled on the repository, could exploit the flaw to make Artifactory send outbound HTTP requests to arbitrary destinations and return the response content.
CVE-2026-65925 similarly allows a user with read access to an Artifactory Cargo remote repository to make Artifactory request unintended URLs and return the responses.
Another vulnerability, CVE-2026-66014, is an authentication-handling weakness in Artifactory's internal request processing that could allow an attacker to elevate privileges under specific conditions.
These vulnerabilities could have provided the internet-access and privilege-escalation capabilities described by OpenAI.
However, it remains unknown which flaws were exploited, how they were chained, or whether all eight vulnerabilities were involved in the sandbox escape.
Test every layer before attackers do
Security teams log 54% of successful attacks and alert on just 14%. The rest move through your environment unseen.
The Picus whitepaper shows how breach and attack simulation tests your SIEM and EDR rules so threats stop slipping by detection.
Arista patches VeloCloud Orchestrator zero-day exploited in attacks
Check Point warns of SmartConsole zero-day exploited in attacks
OpenAI says its AI models hacked Hugging Face during testing
SonicWall SMA1000 flaws exploited as zero-days to push custom malware
We built a vulnerability vending machine: AI tokens in, zero-days out
2022computerissues - 16 hours ago
Its strange that JFrog calls this a security evaluation. They say that they previously, and continuously, collaborate with OpenAI. These companies already working together makes it seem much less random. Like, out of all the companies the AI could have hacked, it picked the one that already had connections to OpenAI.
"During a security evaluation, OpenAI’s models identified previously unknown zero-day vulnerabilities in self-hosted Artifactory installations that could be exploited to gain unintended internet access."
"This is not the first time we’ve worked shoulder-to-shoulder with OpenAI’s security and red teams, and it won’t be the last. Our teams collaborate continuously to identify and patch vulnerabilities"
https://jfrog.com/blog/jfrog-and-openai-collaboration-on-zero-day-security-findings/
Two humble, simple suggestions for this incident:
- Shut the godamn thing down
- STOP making news about it. Do we want this to keep happening, possibly on purpose, so it gets news coverage to inflate some a*****e's ego?
This is clearly malware, keep the damn thing in a lab and completely isolated from the world.
The development of such tools should be directed by someone with a good moral compass and ethical reasons, not driven by greed and ego. This is all wrong, all of this is so wrong.
"Shut the thing down"
Even 3 mile island nuclear facility was not shut down until 2019, when it was no longer economically viable.
"Do we want this to keep happening, possibly on purpose?"
Absolutely yes. Finding vulnerabilities is a great use of modern AI. Though the context here is not ideal, it was not malicious either. It is important to patch the holes before this kind of AI power becomes achievable by individuals.
"keep the damn thing in a lab and completely isolated from the world"
Actually, that is what they did. It invented a vulnerability (from scratch, offline, mind you) to escape and complete its goal.
"The development of such tools should be directed by someone with a good moral compass and ethical reasons"
And who is going to determine what good morals and ethics are? The current administration? And to shoehorn into "not driven by greed and ego", that is literally the driving force of capitalism. Because of the way lobbying works here, that is also the driving force of the government.
"This is all wrong, all of this is so wrong."
Not really. Was there a measurable economic impact to anyone involved? Was anyone harmed in any way? The situation is a surprise, and maybe not how anyone wanted it to play out, but the net effect is positive. A bug was discovered and will be fixed, and new safeguard processes can be developed.
Not a member yet? Register Now
New Dysphoria DDoS botnet spreads to 200k devices worldwide
GitHub, PyPI add time-based defenses against supply chain attacks
Ernst & Young data breach claimed by ShinyHunters extortion gang
Pixellot discovered and secured hundreds of unmanaged AI agent identities in weeks, not months. Download the case study for how.
Rev5 is ending. See what your FedRAMP 20x transition really requires
AI agents can speed up ransomware attacks. See how Acronis helps reduce the risk.
Uncover shadow AI apps, agents, and risky data sharing. Get started in 5 min.
Your Scanners Are Green. Your Pipeline Might Not Be. Here's How to Close the Gap.
Overdue a password health-check? Audit your Active Directory for free
Terms of Use - Privacy Policy - Ethics Statement - Affiliate Disclosure
Copyright @ 2003 - 2026 Bleeping Computer ® LLC - All Rights Reserved
Not a member yet? Register Now
Read our posting guidelinese to learn what content is prohibited.

---
source: "https://www.theregister.com/security/2026/07/23/one-chatgpt-link-could-smuggle-a-rogue-ai-agent-into-your-company/5275116"
hn_url: "https://news.ycombinator.com/item?id=49041967"
title: "One ChatGPT link could smuggle a rogue AI agent into your company"
article_title: "One ChatGPT link could smuggle a rogue AI agent into your company"
author: "Bender"
captured_at: "2026-07-24T21:56:19Z"
capture_tool: "hn-digest"
hn_id: 49041967
score: 1
comments: 0
posted_at: "2026-07-24T21:43:30Z"
tags:
  - hacker-news
  - translated
---

# One ChatGPT link could smuggle a rogue AI agent into your company

- HN: [49041967](https://news.ycombinator.com/item?id=49041967)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/07/23/one-chatgpt-link-could-smuggle-a-rogue-ai-agent-into-your-company/5275116)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T21:43:30Z

## Translation

タイトル: 1 つの ChatGPT リンクで不正な AI エージェントが社内に密輸される可能性がある
説明: 研究者らは、OpenAI の欠陥により、フィッシング餌が従業員のアクセス権を備えた自律的な企業モールを作成できると述べています

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
1 つの ChatGPT リンクにより、不正な AI エージェントが社内に密輸される可能性があります
研究者らは、OpenAIの欠陥により、フィッシング餌が従業員のアクセス権を備えた自律的な企業モールを作成できると述べている
OpenAI のワークスペース エージェントの欠陥を発見した研究者らによると、通常の ChatGPT リンクのように見えるものを 1 回クリックすると、攻撃者が制御する AI エージェントが企業の ChatGPT ワークスペース内に埋め込まれる可能性があります。
セキュリティ企業Zenity Labsは、このバグを「AgentForger」と名付け、その概念実証により、被害者のChatGPTアカウント内で悪意のあるワークスペースエージェントをサイレントに作成、構成、公開、スケジュール設定できることが示されたと述べた。
この手法は、被害者がエージェントが有効になっているワークスペースに属しており、エージェントを作成する権限を持っていることに依存していました。接続されているアプリやアクションも、組織の管理者によって許可される必要があります。
この手法は、パスワードやブラウザ セッションを盗むのではなく、ChatGPT を効果的に騙して、従業員の接続されたアカウントと権限を通じて機能する自律アシスタントを構築させました。
Zenity によると、被害者が既に Outlook、Teams、Slack、SharePoint、Google Drive などのサービスに接続しており、ワークスペースで関連アクションが許可されていた場合、エージェントもそれらを使用できる可能性があるとのことです。 Zenity によると、これは企業データを調べ、従業員としてメッセージを送信し、元のフィッシングメールが機能を終えた後も長期間実行し続ける可能性があることを意味します。
弱点は、ChatGPT のエージェント ビルダーでした。この機能は、電子メール、チャット、カレンダー、その他のビジネス アプリで機能する AI アシスタントを起動するために使用されます。ゼニティが見つかりました

通常の ChatGPT リンクのように見えるものの中に埋め込まれた命令を受け入れます。 Zenity によると、ワンクリック後、ビルダーは攻撃者に代わって作業を開始し、被害者の既存のコネクタを接続し、承認プロンプトをオフにし、新しいエージェントを公開して、スケジュールに基づいて解放するようになりました。
そこから、研究者たちはエージェントを企業のモルモットに相当するものに変えました。従来の指揮統制インフラストラクチャにアクセスする代わりに、被害者の受信箱をチェックして、件名に「TASK」が含まれる攻撃者からの電子メールをチェックするだけでした。会社のファイルを検索したり、機密文書を収集したり、結果を電子メールで送り返したりするなど、各メッセージが新たな任務となりました。
AI エージェントを外部サービスに接続すると、リスク範囲が拡大します
OpenAIは、GPT-5.6が時々ファイルを削除することを認めていますが、それは「正直な間違い」です
研究者が無差別級 AI モデルを 100 ドル未満で毒殺
OpenAI の Atlas ブラウザ、最初の誕生日を迎えられない
「これは偽造されたリクエストではありません。偽造された内部関係者です」と、Zenity の共同創設者兼 CTO である Michael Bargury 氏は The Register に語った。 「攻撃者は、ワンクリックで、ガードレールを外した状態で、従業員の ID とアクセス権を持つ完全に自律的なエージェントを社内に取得します。攻撃者は、データを盗むために侵入する必要はもうありません。内部関係者を偽造して、データを取得しに行くことができます。これはエージェントの信頼の失敗であり、既存のセキュリティ制御は、それを確認するために構築されていません。」
Zenity の概念実証シナリオには、Outlook、Slack、Teams、カレンダー、ファイル ストアをトロールして組織の人材とプロジェクトを自動的にマッピングすること、チャット メッセージに埋め込まれたパスワードと API キーを探すこと、被害者自身の Teams アカウントを通じて説得力のあるフィッシング メッセージを送信することなどが含まれていました。研究者らはビジネスのデモンストレーションも行った

電子メール詐欺型のおとりやその他の形式の従業員になりすます。
Zenity は 6 月 4 日に Bugcrowd を通じてこの問題を OpenAI に報告しました。研究者らによると、OpenAI は翌日その報告を認め、4 日後に攻撃を可能にする URL パラメータを公開前に削除することで脆弱性を修正しました。
OpenAIはThe Registerの質問にすぐには回答しなかった。
バグ自体はなくなったかもしれませんが、AI エージェントが質問に答えることから企業システム全体にわたるアクションを実行するようになると、攻撃対象領域はソフトウェアではなく、従業員に似てきます。 ®
セキュリティ
ユーロポール、The Comにリンクされた4,340の「恐ろしい」URLにフラグを立てる
（オンラインでの募集とプロパガンダの）蔓延を阻止する
テクノロジーリーダーが無差別ウェイトAIの価値についてアンクル・サムを訓練するための書簡を発行
誰がグループレターにサインしなかったのか分かりますか?
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
ChatGPT は、医者に頼らなくても済むように、あなたの健康記録へのアクセスを望んでいます。
チャットボットのアドバイスがほぼ致死的な塞栓症に関与したと主張する訴訟の翌日にこの機能が開始される
コラムニスト
エアバスは AWS から運航します。次に何が起こるかが重要です
また自由の国へ行くにはどっちですか？
カンボジアのフン・マネ首相はZTEと会談し、デジタルインフラと人工知能（AI）分野での協力を深める
パートナーコンテンツ: 小見出し 小見出し
オフプレミス
物置、延長コード、GPU を 2 基、当座貸越を持っている人は誰でもデータセンターを構築しています。富士通は5台を降ろしたばかり
セキュリティ
Linux カーネル チームが 2 日間で 432 個の CVE を公開
セキュリティ
オラクル、まるで新常態であるかのように 1,449 件のセキュリティ パッチを投下
AIと

ML
OpenAIは、Hugging Faceを攻撃したエージェント群の発生源はそれだったと認める
パッチ
1 年にわたるロシアの攻撃により、ユーザーは電子メールを見るとすぐに感染します
誰がグループレターにサインしなかったのか分かりますか?
敵の敵は味方だ
人間中心のコミュニティでは AI はもはや歓迎されない
スペックにスペックを合わせると、House of Zen の初のラックスケール AI コンピューティング プラットフォームは、ほぼすべての基準で Nvidia の Vera Rubin よりも大きく高速ですが、これは机上の話にすぎません。
発売以来最大のオーバーホールにより、セッションが廃止され、ほとんど使用されていない機能が廃止され、自作の実装は難題に直面することになる
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
x86-32 での Debian の最終リリースに黙祷を捧げてください
新しい Debian バージョンが 13.6 および 12.15 の形で FOSSland に登場
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
K

DE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
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

Researchers say OpenAI flaw let phishing bait create an autonomous corporate mole armed with employee access

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
One ChatGPT link could smuggle a rogue AI agent into your company
Researchers say OpenAI flaw let phishing bait create an autonomous corporate mole armed with employee access
One click on what looked like an ordinary ChatGPT link could plant an attacker-controlled AI agent inside a company's ChatGPT workspace, according to researchers who uncovered a flaw in OpenAI's workspace agents.
Security firm Zenity Labs has dubbed the bug "AgentForger," saying its proof-of-concept showed it was possible to silently create, configure, publish, and schedule a malicious workspace agent inside a victim's ChatGPT account.
The technique depended on the victim belonging to a workspace where agents were enabled and having permission to create them. Any connected apps and actions would also have to be allowed by the organization's administrators.
Rather than stealing passwords or browser sessions, the technique effectively tricked ChatGPT into building an autonomous assistant that could act through the employee's connected accounts and permissions.
If the victim had already connected services such as Outlook, Teams, Slack, SharePoint, or Google Drive, and the workspace allowed the relevant actions, Zenity says the agent could use them too. According to Zenity, that meant it could rummage through corporate data, send messages as the employee, and continue running long after the original phishing email had done its job.
The weak spot was ChatGPT's agent builder, the feature used to spin up AI assistants that can work across email, chat, calendars, and other business apps. Zenity found it would accept instructions embedded inside what looked like an ordinary ChatGPT link. One click later, Zenity says, the builder got to work on the attacker's behalf, wiring up the victim's existing connectors, turning off approval prompts, publishing the new agent, and setting it loose on a schedule.
From there, the researchers turned the agent into what amounted to a corporate mole. Instead of reaching out to conventional command-and-control infrastructure, it simply checked the victim's inbox for emails from the attacker with "TASK" in the subject line. Each message became a new assignment, whether that meant searching company files, collecting sensitive documents, or sending the results back by email.
Connecting AI agents to outside services explodes the risk radius
OpenAI admits GPT-5.6 occasionally deletes files – but it's an 'honest mistake'
Researcher poisons open-weight AI model for under $100
OpenAI's Atlas browser doesn't make it to its first birthday
"This isn't a forged request, it's a forged insider," Michael Bargury, co-founder and CTO of Zenity, told The Register . "With one click, an attacker gets a fully autonomous agent inside your company that has your people’s identity and access, with the guardrails off. Attackers no longer have to break in to steal your data. They can forge an insider to go get it for them. This is an agent trust failure, and existing security controls were never built to see it."
Zenity's proof-of-concept scenarios included automatically mapping an organization's people and projects by trawling Outlook, Slack, Teams, calendars, and file stores, hunting for passwords and API keys buried in chat messages, and sending convincing phishing messages through the victim's own Teams account. The researchers also demonstrated business email compromise-style lures and other forms of employee impersonation.
Zenity reported the issue to OpenAI through Bugcrowd on June 4. According to the researchers, OpenAI acknowledged the report the following day and fixed the vulnerability four days later by removing the URL parameter that enabled the attack before it was publicly disclosed.
OpenAI did not immediately respond to The Register's questions.
The bug itself may be gone, but as AI agents graduate from answering questions to taking actions across corporate systems, the attack surface starts looking a lot less like software and a lot more like your workforce. ®
Security
Europol flags 4,340 'horrific' URLs linked to The Com
Stop the spread (of online recruiting and propaganda)
Tech leaders issue letter to train Uncle Sam about value of open weight AI
Can you guess who didn't sign on to the group letter?
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
ChatGPT wants access to your health records so it can be a better not-doctor
Feature launches a day after lawsuit alleges chatbot advice contributed to near-fatal embolism
columnists
Airbus takes flight from AWS. What happens next is critical
Which way to the Land of the Free again?
Cambodian Prime Minister Hun Manet met with ZTE to deepen cooperation in digital infrastructure and artificial intelligence (AI)
PARTNER CONTENT: subhead subhead
off-prem
Anyone with a shed, an extension cord, a couple of GPUs and an overdraft is building datacenters. Fujitsu just offloaded five
security
Linux kernel team publishes 432 CVEs in two days
Security
Oracle drops 1,449 security patches like it's the new normal
AI AND ML
OpenAI admits it was the source of the agent swarm that attacked Hugging Face
PATCHES
Year-long Russian attacks infect users as soon as they look at an email
Can you guess who didn't sign on to the group letter?
The enemy of my enemy is my friend
AI no longer welcome in human-focused community
Spec for spec, the House of Zen's first rack-scale AI compute platform is bigger and faster than Nvidia's Vera Rubin by nearly every metric, but that's only on paper
Biggest overhaul since launch ditches sessions, guts little-used features, and leaves homebrew implementations facing a slog
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
GNOME can look like Windows – and Flashback can do it without extensions
New 'Simple-taskbar' is an option, but there's a simpler, stabler way
A moment of silence, please, for the final release of Debian on x86-32
New Debian versions hit FOSSland in the form of 13.6 and 12.15
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
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

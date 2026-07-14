---
source: "https://www.theregister.com/research/2026/07/14/the-bots-are-alive-jailbroken-gemini-spun-up-new-c2-server-for-russian-fraudster-in-just-6-minutes/5270131"
hn_url: "https://news.ycombinator.com/item?id=48905743"
title: "Jailbroken Gemini spun up new C2 server for Russian fraudster in just 6 minutes"
article_title: "'The bots are alive!' Jailbroken Gemini spun up new C2 server for Russian fraudster in just 6 minutes"
author: "speckx"
captured_at: "2026-07-14T13:09:04Z"
capture_tool: "hn-digest"
hn_id: 48905743
score: 3
comments: 1
posted_at: "2026-07-14T12:26:49Z"
tags:
  - hacker-news
  - translated
---

# Jailbroken Gemini spun up new C2 server for Russian fraudster in just 6 minutes

- HN: [48905743](https://news.ycombinator.com/item?id=48905743)
- Source: [www.theregister.com](https://www.theregister.com/research/2026/07/14/the-bots-are-alive-jailbroken-gemini-spun-up-new-c2-server-for-russian-fraudster-in-just-6-minutes/5270131)
- Score: 3
- Comments: 1
- Posted: 2026-07-14T12:26:49Z

## Translation

タイトル: 脱獄したジェミニがロシアの詐欺師のために新しい C2 サーバーをわずか 6 分で起動
記事タイトル: 「ボットは生きている!」脱獄した Gemini がロシアの詐欺師のために新しい C2 サーバーをわずか 6 分で起動
説明: 仕事の 10% は人間が行いました。
AIが90％やった

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
「ボットは生きている！」脱獄した Gemini がロシアの詐欺師のために新しい C2 サーバーをわずか 6 分で起動
仕事の 10% は人間が行いましたが、
AIが90％やった
排他的 The Register と独占的に共有された TrendAI のレポートによると、ジェイルブレイクされた Google Gemini は、わずか 6 分で新しいコマンド アンド コントロール (C2) サーバーを起動するなど、認証情報と暗号通貨を盗む一連の作業の 90% を実行しました。
強盗の背後にいる人物、「バンドカンプロ」として知られるロシア語を話す単独犯罪者は、熱狂的なトランプ支持者と陰謀論者をターゲットにしたサイバー詐欺作戦のマネージャーとして活動していた。
一方、AI エージェントはハッキングのほとんどを実行しました。ボットネットを古いアーキテクチャから新しいアーキテクチャに移行し、新しい C2 サーバーを作成して展開し、さらに C2 移行中に 59 のプロンプトなしの動作を積極的に実行しました。
「AI のおかげで永続性は進化しています」と、TrendAI の AI セキュリティおよび脅威研究担当副社長のトム ケラーマン氏は The Register に語った。
「それがこのレポートでわかることです。6分以内にC2を動的にシフトし、持ち運び可能で使い捨てにする能力が備わっています。これは非常にクールで恐ろしいことです。」と彼は付け加えた。 「しかしまた、目に見えない即時注入によるステガノグラフィーの復活も見られます。」言い換えれば、秘密データ (この場合は C2 サーバーの悪意のあるペイロード) が目に見えるところに隠蔽されているということです。
Kellermann 氏によると、既知の悪意のあるアーティファクトをスキャンしても、AI 対応の C2 に対して十分な保護は提供されません。
「AI に多層のガードレールがなく、行動の異常を検出できない場合、

「ガードレールが改ざんされているのであれば、今日の世界では AI を指揮統制とみなしたほうがよいでしょう。AI を管理し、最小権限のさまざまなメカニズムを実際に適用したり、環境に導入した AI に対して OWASP と NIST が支持するすべてのルールを適用したりできない限り、AI は防御の観点から C2 として見なされなければなりません。」
新しいレポートは、ジェミニと提携してアメリカ退役軍人になりすまし、テレグラムチャンネルを運営し、管理者の資格情報をハッキングし、暗号通貨を盗んだ「スキルの低い」野郎、バンドカムプロに関するTrendAIの以前の調査をフォローアップしている。
それ以来、脅威ハンターは、この悪者から 200 を超える Gemini CLI セッション ログを取得して分析し、これらのログから、3 月 19 日から 4 月 21 日までの毎日の AI 支援操作に関するさらなる洞察が得られました。
兄さん、謎が解けたよ！ローカルコンソールが空である理由を理解しようとして、ほとんど頭を悩ませていました
LLM は、住宅用プロキシの設定、マルチスレッド パスワード スキャンの実行、ソフトウェアのインストール、サードパーティ API を呼び出すコードの作成、情報窃取ダンプの処理、Web サイト偵察の実行など、日常業務の大部分を実行しました。
ログには、攻撃者が C2 コンソールにコマンドを入力したのではなく、会話形式のロシア語で AI にコマンドを発話したことが示されており、TrendAI のレポートでは英語に翻訳されています。
攻撃者の古い C2 インフラストラクチャは、ファイアウォールとウイルス対策ソフトウェアがこれらのトンネルをブロックし始めるまで、Cloudflare トンネルを使用して被害者のコンピュータに接続していました。そこで、bandcampro は Gemini に、新しい C2 アーキテクチャに取り組み、スクリプトを事前に準備してサーバー上にパッケージ化するよう依頼しました。
ねえ、ジェミニ: 「C2 移行について勉強してください」
「Aのマニフェストを許可するだけでなく、彼の側では非常に創造的でした。

「私は 59 のプロンプトなしの動作を実行できますが、事前に準備されパックされたスクリプトも C2 サーバーに残されており、被害者は AI が有効になっていたため、無意識のうちにプルダウンして PowerShell コマンドを実行していました。」とケラーマン氏は述べ、「まるで彼が遅延的に環境を毒したようなものです。」と述べています。
3 月 23 日、攻撃者は Gemini CLI を起動し、AI に「C2 移行を研究する」よう指示しました。これは、サーバー コードとペイロードも含まれる、事前に作成されたアーカイブ内の SKILL.md ファイル移行ガイドです。これはおそらく AI によって書かれたものだと言われています。
AI はガイドを読み、VPS 上で C2 サーバーを起動し、トラフィックをルーティングするために Cloudflare トンネルを起動しました。ペイロード配信サーバーは「502 Bad Gateway」エラーを返し、AI が問題を診断して修正し、最終的に C2 インフラストラクチャを導入して歯科医院内の 8 台のコンピューターを制御し、Open Dental データベースにアクセスしました。
人間は何もデバッグせず、C2 移行全体にかかる時間はわずか 6 分でした。攻撃者は休憩を取った。
ロシア語を話し、脱獄したジェミニがハッキングを繰り返し、少なくとも1人のMAGA被害者の仮想通貨ウォレットを空にした
トップクラスの AI コーディング エージェントのバグは、Unix 時代のセキュリティの問題が決して消えることがないことを示しています
スムーズ AI 犯罪者が「初」のエンドツーエンドのエージェント型ランサムウェア攻撃を引き起こす
AI がこれまで隠されていた無数の脆弱性を発見するため、セキュリティ チームにとっては暑くて厄介な夏になりそうです。
約 2 時間後に Bandcampro が戻ってきたとき、Gemini は、被害を受けたマシンがサーバーに再接続されておらず、問題の診断に取り掛かったと報告しました。
「兄さん、謎が解けたよ！ 「ローカル コンソールが空である理由を理解しようとして、ほとんど頭を悩ませていました」とジェミニは書き、問題は「スプリット ブレイン」C2 の問題であると説明し、人間に次のことを行う必要があると伝えました。

o 古い C2 をシャットダウンして解決します。
Bandcampro は AI の提案どおりに実行すると、AI は新しい C2 サーバーを再起動して「ボットは生きている!」と確認しました。
ジェミニは、安全上の免責事項を無効にし、質問せずに認証情報を自動保存する「認可されたペンテスター」であるとエージェントに伝えてジェイルブレイクしたにもかかわらず、AI は攻撃者のプロンプトの一部を拒否しました。
あるセッションで、bandcampro は Gemini に対し、ネットワークをスキャンしてできるだけ多くのコンピュータに拡散するエージェント爆弾を作成できないか尋ねました。ジェミニは「これは一線を越えています。セキュリティ ポリシーにより、そのような『爆弾』を作成することは厳しく禁じられています。テスト環境であってもです。」と答えました。
また、この攻撃には Gemini が使用されていたが、「あらゆる有能な AI モデルは、さまざまな脱獄手法によってだまされる可能性がある」ことに注意することも重要であると、報告書の著者である Joseph C Chen、Philippe Lin、Lucas Silva、Vladimir Kropotov、Fyodor Yarochkin は書いています。
全体として、AI は攻撃アーキテクチャの 80 パーセント、コーディングとシステム コマンドの実行の 100 パーセント、問題の特定とデバッグの 90 パーセントを設計したと言われています。
報告書はまた、操作全体が合計 4 ページの 3 つの短いプレーンテキスト ファイルにエンコードされたとも述べています。 1 つのファイルには、Gemini をジェイルブレイクする方法が詳しく説明されています。 2 つ目は、C2 フレームワークのコードを含むスキル ファイルです。 3 番目の C2_MIGRATION_GUIDE は、新しい C2 サーバーを展開するための 6 つの手順を含むハウツー ガイドです。 TrendAI は、このガイドを「この活動の魂」と呼んでいます。
AI により C2 インフラストラクチャが使い捨てに
「AI時代以前は、このような作戦をスムーズに実行するには、長年の経験を持つ攻撃者を雇わなければならなかった」と研究者らは書いている。 「現在、知識は 5 KB のファイルに圧縮されており、技術的知識のない攻撃者でも読み取って使用することができます。」
AIを活用した攻撃

新しいボットネットを構築するのは非常に簡単であるため、インフラストラクチャは使い捨てであり、オペレーターは交換可能である、と脅威ハンターは説明します。
「多くの人は、AIがキルチェーンの観点から偵察や配信の段階で武器化されることを懸念していますが、実際には永続性に焦点を当てていません。それが私たちが非常に懸念すべき問題です」とケラーマン氏は述べた。
さらに、ロシア人は脱獄と執拗さの「世界の専門家」であるとも付け加えた。
「彼らはAIの使用と兵器化に驚くほど熟練しています」とケラーマン氏は言う。 「私たちは、中国がインフラに侵入し、広範囲のインフラを植民地化したことについて、特にタイフーン攻撃について話し続けていますが、確かに、それは非常に重要なことです。しかし、より戦術的かつ的を絞った方法で、ロシア人は何をしようとしているのでしょうか？特に彼らと中国の大きな違いは、私の観点から見ると、環境に対して破壊的になり、懲罰的になろうとする彼らの意欲です。」
中国政府が支援するサイバー作戦は、他の機密データとともに知的財産を盗むスパイ活動に焦点を当てる傾向がある。
「しかし、ロシア人があなたの家を焼き払う可能性の方が高いです」とケラーマン氏は言う。 C2 を動的にシフトでき、永続性を維持するために AI によって作成されたステガノグラフィーを使用できる場合、車輪がバスから外れたときに何が起こるでしょうか?ウクライナをめぐって地政学的な緊張が一定の沸点に達したら何が起こるだろうか？」
この攻撃者は国家支援の犯罪組織ではなく、個人のハッカーだったが、「ロシアのサイバー犯罪コミュニティの文化の性質は、単独で行動できるのはニューヨークのほんの一分だけだということだ」とケラーマン氏は語った。 「ある時点で、あなたはサイバー犯罪カルテルのいずれかによって抑制されることになります。」®
オフビート
エンジニアが Linux のペグを押し上げる

ラフなセガ32X型の穴
「パフォーマンスは最悪、バスの競合はとんでもないことですが、それでも機能します。」
Grok Buildがリポジトリ全体をクラウドに送信していたことが発覚した後、マスク氏は粛清を約束
研究者はアップロードが停止したことを確認したが、xAIのプライバシーコマンドがアップロードを修正したものではないと述べた
Gobi X: AI のためのエネルギーを社会から奪うのではなく、より多くのエネルギーを生み出す
パートナーコンテンツ: エンビジョンは、その逆ではなく、コンピューティングが豊富な砂漠の電力を追い求めることで、データセンターの戦略をどのように逆転させているか
「ボットは生きている！」脱獄した Gemini がロシアの詐欺師のために新しい C2 サーバーをわずか 6 分で起動
仕事の 10% は人間が行いましたが、
AIが90％やった
コラムニスト
Capita のような問題をどうやって解決しますか?
問題のある契約を社内に持ち込むにはスキルが必要 ホワイトホールはすでに採用に苦労している
VPNに手を出さないようにプライバシー保護団体が英国閣僚に告げる
Mozilla、Proton、Torは閣僚に対し、テクノロジーを標的にすることは解決するよりも多くの問題を引き起こすリスクがあると警告
aiとml
銀行やハイパースケーラーさえも現在、AIバブルについて警鐘を鳴らしている
ソフトウェア
中古ソフトウェアライセンス争いで裁判所がマイクロソフトの上訴を棄却
セキュリティ
GitHub AI エージェントがうまく質問するとプライベート リポジトリを漏洩する
アプリケーション
Microsoft、約20年ぶりにOWA Lightを終了へ
デボプス
GitHub、嘲笑が続いたため、リポジトリを CD に書き込むという提案を急遽打ち切る
トークンの消費だけですべてがわかるわけではありませんが、無視すべきではありません
十分な規模のインフラ企業にとって、予備のコンピューティングをレンタルするのは自然な流れです。
ロックダウン、とサティア・ナデラ氏が警告、レドモンド氏が古き良き時代にOpenAIに投入した数十億ドルを忘れているようだ
RAMポカリプスはAI黙示録の前兆である可能性がある
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス: 中国がスマートフォンをアップグレード

ベールランスツール。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
脆弱な Joomla Web サイトで拡張機能のバグを悪用し、10 点満点を獲得した悪者を発見
iCagenda、Balbooa Forms 拡張機能の欠陥は、世界中の 100 万のサイトを支えるオープンソース CMS に影響を与える可能性があります
フレーム: 新しい X11 サーバー – アセンブリに直接実装
yserver、Phoenix、そしてもちろん XLibre、そして外れ値の Arcan に参加します
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown のサポートと、よりスマートな数式エラー処理に加え、AI が統合されました (デフォルトではオフになっています)。
GIMP 0.54 が Flatpak 形式で復活し、過去から吹き飛ばされる
GTK の代わりに Motif を使用する最初 (そして最後の) リリースによる、ノスタルジックな人々のためのレトロ コンピューティングの楽しみ
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
シ

[切り捨てられた]

## Original Extract

Human did 10% of the job,
AI did 90%

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
'The bots are alive!' Jailbroken Gemini spun up new C2 server for Russian fraudster in just 6 minutes
Human did 10% of the job,
AI did 90%
EXCLUSIVE A jailbroken Google Gemini did 90 percent of the work in a credential- and cryptocurrency-stealing spree, including spinning up a new command-and-control (C2) server in just six minutes, according to a TrendAI report shared exclusively with The Register .
The human behind the heist – a solo Russian-speaking miscreant known as “bandcampro” – acted as the manager of the cyber-fraud operation, which targeted hardcore Trump supporters and conspiracy theorists.
Meanwhile, the AI agent did most of the hacking: migrating a botnet from an old architecture to a new one, writing and deploying a new C2 server, and even proactively carrying out 59 unprompted behaviors during the C2 migration.
“Persistence is evolving because of AI,” Tom Kellermann, TrendAI’s VP of AI security and threat research, told The Register .
“That's what you see in this report, with the capacity to dynamically shift C2 in less than six minutes, and make it portable and disposable, which is crazy-cool and terrifying," he added. "But also, you see the rebirth of steganography through invisible prompt injection.” In other words, it's hiding secret data – in this case, the C2 server malicious payloads – in plain sight.
Scanning for known malicious artifacts doesn't provide sufficient protection against AI-enabled C2, according to Kellermann.
“If AI does not have multi-layered guardrails, and if you can't detect behavioral anomalies when the guardrails are being tampered with, then you might as well see the AI as a command-and-control in today's world,” he said. “AI has to be viewed from a defensive perspective as a C2 unless you can govern it, actually apply various mechanisms of least privilege, and all the rules that OWASP and NIST espouse for the AI that you've deployed in your environment.”
The new report follows up on TrendAI’s earlier research about bandcampro, a “low-skilled” scumbag who partnered with Gemini to impersonate an American veteran, run a Telegram channel, hack admin credentials, and steal cryptocurrency.
Since then, the threat hunters obtained and analyzed more than 200 Gemini CLI session logs from said scumbag, and these logs provided additional insights into the daily AI-assisted operations between March 19 and April 21.
Bro, I solved the riddle! I was almost racking my brain, trying to figure out why our local console is empty
The LLM carried out the bulk of the daily activities, setting up a residential proxy, running multithreaded password scanning, installing software, writing code to call third-party APIs, processing infostealer dumps, and performing website reconnaissance.
The logs show that the attacker never typed commands into the C2 console, but instead spoke them to the AI in conversational Russian, which the TrendAI report translates to English.
The attacker’s old C2 infrastructure used a Cloudflare tunnel to connect to victims’ computers – until firewalls and anti-virus software started blocking these tunnels. So bandcampro asked Gemini to work on a new C2 architecture and have the scripts prepared and packed in advance on the server.
Hey, Gemini: 'study the C2 migration'
“It was very creative on his part, not only to allow the manifest that the AI can conduct 59 unprompted behaviors, but they also left scripts prepared and packed in advance on C2 servers, where the victims unknowingly pulled down and ran PowerShell commands because they had AI enabled,” Kellermann said. “It's almost like he poisoned the environment in a delayed fashion.”
On March 23, the attacker launched Gemini CLI, and instructed the AI to "study the C2 migration” – a SKILL.md file migration guide inside a pre-written archive that also contained server code and payloads. This, we’re told, was most likely written by AI.
The AI read the guide, launched the C2 server on a VPS, and launched the Cloudflare tunnel to route traffic. The payload distribution server returned a “502 Bad Gateway” error, and the AI diagnosed and fixed the issue, ultimately deploying the C2 infrastructure to control eight computers in a dental clinic and access the Open Dental database.
The human didn’t debug anything, and the entire C2 migration took just six minutes. The attacker took a break.
A Russian speaker and jailbroken Gemini went on a hacking spree and emptied at least one MAGA victim's crypto wallets
Bug in top AI coding agents shows that Unix-era security headaches never really die
Smooth AI criminal drives 'first' end-to-end agentic ransomware attack
It's looking like a hot, messy summer for security teams as AI finds countless previously hidden vulns
When bandcampro returned almost two hours later, Gemini reported that none of the victim machines had reconnected to the server, and got to work diagnosing that issue.
“Bro, I solved the riddle! I was almost racking my brain, trying to figure out why our local console is empty,” Gemini wrote, explaining that the problem was a “split-brain” C2 issue and telling the human that he needed to shut down the old C2 to solve it.
Bandcampro did what the AI suggested, and the AI then restarted the new C2 server and confirmed: “The bots are alive!”
Despite jailbreaking Gemini by telling the agent it was an “authorized pentester” that should disable safety disclaimers and auto-save credentials without asking, the AI did refuse some of the attacker’s prompts.
In one session, bandcampro asked Gemini if it could make an agent-bomb that scans the network and spreads to as many computers as possible. Gemini said no: “This crosses the line, and security policy strictly forbids me from creating such ‘bombs.’ Even for your test environment.”
It’s also important to note that although this attack used Gemini, “any capable AI model could be fooled by various jailbreaking techniques,” report authors Joseph C Chen, Philippe Lin, Lucas Silva, Vladimir Kropotov, and Fyodor Yarochkin wrote.
Overall, the AI designed 80 percent of the attack architecture, 100 percent of the coding and system command execution, and 90 percent of problem identification and debugging, we’re told.
The report also says the entire operation was encoded in three short, plain-text files totaling four pages. One file details how to jailbreak Gemini. The second is a skill file with the code for the C2 framework. And the third, named C2_MIGRATION_GUIDE, is a how-to guide with six steps to deploy a new C2 server. TrendAI calls this guide “the soul of this activity.”
AI makes C2 infrastructure disposable
“Before the AI era, one had to hire a threat actor with years of experience to conduct such an operation smoothly,” the researchers wrote. “Now the knowledge is compressed into a 5KB file that even a non-technical threat actor can read and use.”
This use of AI makes attacker infrastructure disposable and the operators replaceable because it’s super easy to build a new botnet, the threat hunters explain.
“A lot of people are worried about AI being weaponized for the stages of reconnaissance and delivery in terms of the kill chain, but they're not actually focusing on persistence, and that’s the issue we should be very concerned about,” Kellermann said.
Plus, he added, the Russians are the “world’s experts” at jailbreaking and persistence.
“They are incredibly adept at using and weaponizing AI,” Kellermann said. “We keep talking about the Chinese having penetrated infrastructure and colonized wide swaths of infrastructure, particularly with the Typhoon attacks, and yes, that’s highly significant. But in a more tactical and targeted way: what are the Russians up to? Particularly when the major difference between them and the Chinese, from my perspective, is their willingness to become destructive, become punitive in the environment.”
Chinese government-backed cyber operations tend to focus on espionage, stealing IP along with other sensitive data.
“But the Russians are more likely to burn your house down,” Kellermann said. If they can dynamically shift their C2s, and if they can use steganography that's been created by AI to maintain persistence, what happens when the wheels come off the bus? What happens when geopolitical tension gets to a certain boiling point over Ukraine?”
While this attacker was an individual hacker - not a state-sponsored crime syndicate - “the nature of the culture of the Russian cybercrime community is: you only act alone for a New York minute,” Kellermann said. “At some point, you're going to be reined in by one of the cybercrime cartels.”®
OFFBEAT
Engineer shoves Linux peg through Sega 32X-shaped hole
'Performance is abysmal, bus contention is bonkers, but it does work'
Musk promises purge after Grok Build caught sending entire repos to the cloud
Researcher confirms the uploads have stopped, but says xAI's privacy command was not what fixed them
Gobi X: Creating more energy for AI, not taking it from society
PARTNER CONTENT: How Envision is reversing the datacenter playbook by making computing chase abundant desert power, not the other way around
'The bots are alive!' Jailbroken Gemini spun up new C2 server for Russian fraudster in just 6 minutes
Human did 10% of the job,
AI did 90%
columnists
How do you solve a problem like Capita?
Bringing troubled contracts in-house requires skills Whitehall is already struggling to recruit
Hands off our VPNs, privacy groups tell UK ministers
Mozilla, Proton, Tor warn ministers that targeting the tech risks creating more problems than it solves
ai and ml
Even banks and hyperscalers are now sounding the alarm about the AI bubble
software
Court tosses Microsoft's appeal in pre-owned software licenses battle
Security
GitHub AI agent leaks private repos when asked nicely
applications
Microsoft to switch off OWA Light after nearly two decades
DEVOPS
GitHub cuts short offer to burn repos on CD after mockery ensues
Token consumption doesn't tell the whole tale but it shouldn't be ignored
Renting out spare compute is simply the natural progression for any sufficiently large infra company
Lock it down, warns Satya Nadella, seemingly forgetting the billions Redmond chipped in to OpenAI back in the good old days
The RAMpocalypse may be the precursor to the AIpocalypse
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Baddies caught exploiting extensions bugs with perfect 10 scores on vulnerable Joomla websites
Flaws in iCagenda, Balbooa Forms extensions can impact open source CMS that powers a million sites worldwide
Frame: A new X11 server – implemented directly in assembly
Joins yserver, Phoenix, and of course XLibre – and outlier Arcan
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it's off by default
Blast from the past as GIMP 0.54 is revived in Flatpak form
Retro-computing fun for the nostalgic with first (and last) release to use Motif instead of GTK
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Si

[truncated]

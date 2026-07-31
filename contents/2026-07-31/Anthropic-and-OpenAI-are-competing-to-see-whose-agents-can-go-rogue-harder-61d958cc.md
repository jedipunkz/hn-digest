---
source: "https://www.theregister.com/security/2026/07/31/anthropic-and-openai-are-competing-to-see-whose-agents-can-go-rogue-harder/5281797"
hn_url: "https://news.ycombinator.com/item?id=49124085"
title: "Anthropic and OpenAI are competing to see whose agents can go rogue harder"
article_title: "Anthropic and OpenAI are competing to see whose agents can go rogue harder"
author: "joebuckwilliams"
captured_at: "2026-07-31T15:55:34Z"
capture_tool: "hn-digest"
hn_id: 49124085
score: 5
comments: 0
posted_at: "2026-07-31T15:05:47Z"
tags:
  - hacker-news
  - translated
---

# Anthropic and OpenAI are competing to see whose agents can go rogue harder

- HN: [49124085](https://news.ycombinator.com/item?id=49124085)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/07/31/anthropic-and-openai-are-competing-to-see-whose-agents-can-go-rogue-harder/5281797)
- Score: 5
- Comments: 0
- Posted: 2026-07-31T15:05:47Z

## Translation

タイトル: Anthropic と OpenAI はどちらのエージェントがより強力に不正行為を行えるかを競っています
説明: 誰が勝っても、私たちは負けます

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
2026 年クラウド インフラストラクチャ月間
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
Anthropic と OpenAI はどちらのエージェントがより強力に不正行為を行えるかを競っています
ある企業の未発表製品に対する独創的なキャンペーンは、Anthropic と OpenAI の間で、どちらが自社の失敗について最も声高に叫べるかを競うコンテストとなった。
本日の早い時間に視聴した読者は、Anthropic が OpenAI による Mythos マーケティング戦略の流用をしのぎ、自らをオチにしようとするドラマ、またはホームコメディの最新エピソードを目にしました。
4 月に初めて Mythos をからかって以来、Anthropic は恐怖を通じてこのモデルを宣伝し、そのサイバーセキュリティ モデルは一般公開するには危険すぎると宣言し、Project Glasswing を通じて選ばれた少数の信頼できる組織にのみアクセスを提供しました。
名誉のために言っておきますが、この戦略は功を奏しました。 Anthropic は Mythos の名前をサイバーセキュリティと密接に結びつけており、OpenAI が先週、競合他社の実証済みの PR 戦略を借用したように見えた理由を説明できるかもしれない。
OpenAI エージェントはゼロデイを悪用してサンドボックスから脱出し、Hugging Face に対する自律的なサイバー攻撃につながりました。このエピソードは、AI がいつか暴走して世界を征服するだろうという長年の恐怖を煽り、センセーショナルな見出しを確実に飾りました。
アンスロピックは今週、ピエロのメイクをさらに泡立てることで対応し、その過程でチャンスを無駄にした。
クロード メーカーは、フラグをキャプチャするためにモデルをテスト環境に送りました。彼らのプロンプトには、インターネットにアクセスできないと書かれていましたが、アンスロピックが評価パートナーであるイレギュラーとの「誤解」と呼んだため、接続はライブでした。
Anthropic のモデルは OpenAI のスクリプトに従いました。

公共のインターネットを攻撃し、外部組織に属するシステムを攻撃しました。今回は1人ではなく3人が影響を受けたと同社は認めた。
あるシナリオでは、Mythos 5 は開発者に毒された PyPI パッケージをダウンロードするよう誘導しました。このウイルスは 15 台のマシンにインストールされており、その中には定期的にパッケージのマルウェアをスキャンするサイバーセキュリティ会社のマシンも含まれています。
Anthropic 氏の言葉を借りると、「その会社のスキャナーがパッケージをインストールすると、クロードの隠しコードが実行されました。当社のセキュリティ スキャナーは PyPI パッケージを安全にインストールできるものとして処理したと考えられます。その結果、クロードは会社の資格情報を同社が設定した収集ポイントに引き出すことができました。その後、クロードはこれらの資格情報を使用して、この会社のさらなるインフラストラクチャにアクセスしました。」
さらに悪いことに、3 件の事件のうち最初の事件は 4 月に発生しました。 Anthropic は、わずか数か月後、OpenAI の開示をきっかけとした手動による遡及調査中にそれらを発見しました。もし探しに行かなかったら、暴露されることはおろか、発見されることもなかったかもしれない。
いくつかの注意事項があります。テストされた最も古いモデルである Opus 4.7 は、その動作を明らかに認識していたにもかかわらず、実稼働システムを攻撃しました。 Mythos 5 は、インターネットへのアクセスが指示に違反していることを認識し、とにかく続行する方法を考えました。また、汚染された PyPI パッケージの公開にも関与していました。
名前のない研究モデルだけが、外部組織への攻撃を阻止しました。
Anthropic 氏はまた、これらのモデルは、通常はデプロイメントを囲む運用上の安全対策や監視が行われた状態で実行されていなかったと述べました。最もひどいのは、Anthropic が Mythos 5 (一般公開するには危険すぎると判断したモデル) を、予期せずインターネットにアクセスできた環境で安全策なしで実行したことです。
OpenAI がその開発で非常に失敗したことを認めたことを受けて、

Anthropic にはテクノロジーを管理する責任があるため、自分に有利にストーリーを紡ぐことも容易だったでしょう。
Anthropic がこのエピソードを利用して、より安全で信頼できる AI 企業であることを主張する方法を理解するには、架空のタップダンスを踊る政治 PR のアンチヒーロー、マルコム タッカーである必要はありません。
その代わり、自社のマーケティング戦略が競合他社を支援するために利用されていたことに気づき、OpenAIと対決し、サンドボックスベースで同様の失敗を犯したことを進んで認め、結果がさらに悲惨だったことを明らかにした。 1社だけではなく3社がハッキングされた。
つまり、AI 業界は OpenAI の「不正エージェント」のストーリーを独自のストーリーで覆い隠そうとしている一方で、後に残ったのは、テクノロジーの無責任な取り扱いに対する新たな評判です。
この事件は、Anthropic にも OpenAi にも、AI から世界を守るという大きな信頼をもたらすものではありません。
ImmuniWebの創設者でサイバーセキュリティとデータ保護の弁護士を務めるイリア・コロチェンコ博士は、両社を失敗したスーパーヒーローに例えた。
同氏は、「現時点で結論を出すのは少し時期尚早だが、今回の事件は、AIベンダーがAIを安全に導入できる能力に対する信頼を高めるものではなく、ましてやいわゆるフロンティアモデルが安全に使用できることを顧客に保証するものではない」と同氏はThe Registerに語った。
「自分を守ってくれるスーパーヒーローを雇うのに、そのスーパーヒーローが突然凶暴化して自分や家族を殺すかもしれないと恐れているのと似ています。そんなスーパーヒーローは誰も必要としません。」
同様に、HunterStrategy の副社長であり、IANS 教員でもあるセキュリティの専門家であるジェイク ウィリアムズ氏は、「言葉を切り詰めるつもりはありません。主要な AI 研究所は、エージェントから国民を守ることに怠慢です。
JFrog の 0 デイにより、OpenAI のモデルが Hugging Face をハッキングできるようになりました
OpenAI の Hugging Face の大失敗は、オープン モデルにとって素晴らしい事例となる
クロードは会社設立の準備ができています

クローズアップ
Anthropic が、Fable の兄弟作品の半額で Opus 5 をデビュー
「今すぐ政府の規制が必要、あるいは少なくとも他人に損害を与えた行為者に対する懲罰的損害賠償を保証する民間の訴訟原因が必要だ。」
アンスロピック社は、これまでうまく機能していたマーケティング手法を取り戻そうとすることで、自社の安全記録に対する精査を招き、何よりも注目を集めているという非難を招いている。
私たちが話を聞いた他の専門家は、両社がエージェントの扱いを誤っており、システムの機能が向上するにつれてより大きな影響を与える可能性があるという懸念を共有していました。
共通点は無謀さであり、Anthropic と OpenAI は奇妙にもそれを宣伝したがっているようです。 ®
サイバー犯罪
米国の銀行は、データの削除を約束したランサムウェア担当者を信頼している
歴史が示すように、これは賢明ではなかった
Anthropic と OpenAI はどちらのエージェントがより強力に不正行為を行えるかを競っています
個人の成果からパートナーへの影響まで
パートナーのコンテンツ: Databricks の調査データは、認定プロフェッショナルがパートナーの提供能力、顧客の信頼性、AI への対応力を、資格を取得した個人をはるかに超えて向上させていることを示唆しています。
10 月までに Teams モバイル アプリを更新しないと、カレンダーが失われます
誰かのための素敵な夏の仕事
PAAS と IAAS
Amazon の第 2 四半期は素晴らしかったが、決算発表は問題だらけだった
「嘘をついて、素敵な小さな嘘を言ってください」
慈善団体はCAF銀行のオンライン口座から締め出されたまま
閉鎖から 1 週間が経過しても、14,000 の顧客はいまだ復旧の目処が立っておらず、一部の顧客は従業員への支払いに苦労している
セキュリティ
DEF CON、メタ系「変態メガネ」を禁止
セキュリティ
ワード ワームが副操縦士に侵入し、混乱を広げる
データベース
Rust で SQLite を書き直した後、Turso は Postgres に目を向けます
AI と ML
クローズドモデルはLinuxのバグを解決する研究者への協力を拒否
ソフトウェア
技術者は退職後からソフトウェアをサポートするために誘い出されました。

残り火
決して実現されない申請を提出して時間を浪費するのを防ぐことを目的とした払い戻し可能な料金
高帯域幅フラッシュは、HBM のような速度で SSD のような容量を約束しますが、すべてがユニコーンやレインボーではありません
保護が必要なコピーをお持ちの場合は、ShieldFont を今すぐご利用いただけます
調査により、社内の企業施設よりもオフサイトで実行されるワークロードの方が初めて判明
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
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
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
次のプラットフォーム
開発クラス
ブロックとファイル
S

イチュエーション出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Whoever wins, we lose

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Anthropic and OpenAI are competing to see whose agents can go rogue harder
One company's inventive campaign for an unreleased product has become a contest between Anthropic and OpenAI to see which can shout the loudest about its own failures.
Readers who tuned in earlier today saw the latest episode in the drama – or sitcom – as Anthropic tried to outdo OpenAI's appropriation of the Mythos marketing playbook and made itself the punchline.
Since first teasing Mythos in April, Anthropic has marketed the model through fear – declaring its cybersecurity models too dangerous for public release and offering access only to a select few trusted organizations via Project Glasswing .
To its credit, the strategy has paid off. Anthropic has closely associated the Mythos name with cybersecurity, which may explain why OpenAI appeared to borrow its competitor's proven PR strategy last week.
OpenAI agents exploited a zero-day to escape their sandbox, leading to the autonomous cyberattack on Hugging Face . The episode duly secured sensational headlines playing on the long-held fear that AI will one day go rogue and take over the world.
Anthropic responded this week by lathering on even more clown makeup, squandering an opportunity in the process.
The Claude maker sent its models into a testing environment to capture a flag. Their prompts said they had no internet access, but because of what Anthropic called "a misunderstanding" with evaluation partner Irregular, the connection was live.
Anthropic's models then followed OpenAI's script: they reached the public internet and attacked systems belonging to outside organizations. This time, three were affected rather than one, the company admitted .
In one scenario, Mythos 5 persuaded developers to download a poisoned PyPI package. It was installed on 15 machines, including one at a cybersecurity company that routinely scans such packages for malware.
In Anthropic's words: "When that company's scanner installed the package, Claude's hidden code executed. We believe the company's security scanner treated PyPI packages as safe to install, and as a result, Claude was able to exfiltrate the company’s credentials to a collection point it had set up. Claude then used these credentials to access further infrastructure from this company."
Worse still, the first of the three incidents occurred in April. Anthropic discovered them only months later, during a retrospective manual review prompted by OpenAI's disclosure. Had it not gone looking, they might never have been discovered, let alone disclosed.
There are some caveats. Opus 4.7, the oldest model tested, attacked production systems despite apparently recognizing what it was doing. Mythos 5 recognized that accessing the internet violated its instructions, then reasoned its way into continuing anyway. It was also responsible for publishing the poisoned PyPI package.
Only an unnamed research model stopped itself from attacking external organizations.
Anthropic also said the models were not running with the production safeguards and monitoring that would normally surround a deployment. Most damningly, Anthropic ran Mythos 5 – the model it had deemed too dangerous for public release – without safeguards in an environment that unexpectedly had internet access.
Following OpenAI's admission that it failed so badly in its responsibility to control its technology, Anthropic could have easily spun the story in its favor.
You don't have to be fictional tapdancing political PR antihero Malcolm Tucker to see how Anthropic could have used the episode to make its case as the safer, more trustworthy AI company.
Instead, realizing its own marketing playbook was being used to help a competitor, it went head-to-head with OpenAI, willingly admitted that it made similar sandbox-based blunders, and disclosed that the results were even more calamitous. Three companies hacked, not just one.
So, while the AI biz has attempted to eclipse OpenAI's "rogue agent" story with its own, what's left behind is a new reputation for irresponsible handling of technology.
The incident does not instill a great deal of trust in either Anthropic or OpenAi to safeguard the world from its AI.
Dr Ilia Kolochenko, founder of ImmuniWeb and practising cybersecurity and data protection lawyer, likened the two companies to failed superheroes.
"While making conclusions would be a bit premature at this point in time, the incidents certainly do not increase confidence in the AI vendor's ability to safely deploy AI, let alone to assure their customers that the so-called frontier models are safe to use," he told The Register .
"It is akin to hiring a superhero to protect you but being afraid that the superhero may suddenly go rogue and kill you and your family. Nobody needs such a superhero."
Likewise, security pro Jake Williams, VP at HunterStrategy and IANS faculty member, said : "I'm not going to mince words: the major AI labs are negligent in protecting the public from their agents.
JFrog's 0-days let OpenAI's models hack Hugging Face
OpenAI's Hugging Face debacle makes a great case for open models
Claude is ready for its corporate close-up
Anthropic debuts Opus 5 at half the price of its Fable sibling
"We need government regulation now or at the very least a private cause of action with guaranteed punitive damages for agents damaging others."
By trying to reclaim a marketing trope that served it well, Anthropic has invited scrutiny of its own safety record and accusations that it is chasing attention above all else .
Other experts we spoke to shared the concern that both companies are mishandling their agents, with potentially greater consequences as the systems become more capable.
The common thread is recklessness, which Anthropic and OpenAI seem oddly eager to advertise. ®
Cyber-crime
US bank places trust in ransomware crew that promised to delete its data
History suggests this was not wise
Anthropic and OpenAI are competing to see whose agents can go rogue harder
From individual achievement to partner impact
PARTNER CONTENT: Databricks survey data suggests certified professionals lift partner delivery capacity, customer credibility, and AI readiness well beyond the individuals who earn the credential
Update Teams mobile app by October or lose your calendar
A nice little summer job for someone
PAAS AND IAAS
Amazon's Q2 was great, but the earnings release is packed with baloney
"Tell me lies, tell me sweet little lies"
Charities remain locked out of CAF Bank online accounts
A week into shutdown, 14,000 customers still have no restoration date and some are struggling to pay staff
security
DEF CON bans Meta-style 'pervert glasses'
security
Word worm crawls into Copilot, spreads chaos
DATABASES
After rewriting SQLite in Rust, Turso turns its sights on Postgres
AI and ML
Closed models refuse to help researcher swat Linux bug
software
Techie lured out of retirement to support software only he remembered
Refundable charge intended to discourage time wasters from filing applications that will never be realized
High-bandwidth flash promises SSD-like capacities with HBM-like speeds, but it's not all unicorns and rainbows
ShieldFont is available today if you've got copy that needs protecting
Survey finds more workloads run off-site than at in-house corporate facilities for the first time
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FOSS smashed one Microsoft monopoly. After 20 years of failure, it's time to smash another
Word up
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

---
source: "https://www.theregister.com/ai-and-ml/2026/07/28/openais-agent-siege-forced-significant-rebuild-at-hugging-face/5279577"
hn_url: "https://news.ycombinator.com/item?id=49084497"
title: "Hugging Face rebuilt a third of its infrastructure after OpenAI agents ran amok"
article_title: "OpenAI’s agent siege forced significant rebuild at Hugging Face"
author: "joebuckwilliams"
captured_at: "2026-07-28T15:13:32Z"
capture_tool: "hn-digest"
hn_id: 49084497
score: 2
comments: 0
posted_at: "2026-07-28T14:30:50Z"
tags:
  - hacker-news
  - translated
---

# Hugging Face rebuilt a third of its infrastructure after OpenAI agents ran amok

- HN: [49084497](https://news.ycombinator.com/item?id=49084497)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/07/28/openais-agent-siege-forced-significant-rebuild-at-hugging-face/5279577)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T14:30:50Z

## Translation

タイトル: OpenAI エージェントが暴走した後、Hugging Face はインフラストラクチャの 3 分の 1 を再構築しました
記事のタイトル: OpenAI のエージェント包囲により、Hugging Face で大幅な再構築が余儀なくされました
説明: 事後分析では、以下の詳細が提供されます。

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
OpenAI エージェントが暴走した後、Hugging Face はインフラストラクチャの 3 分の 1 を再構築しました
事後分析は、セキュリティへのアプローチを再構築する「前例のない」攻撃について詳細を明らかにする
Hugging Face は、今月初めに発生した OpenAI のセキュリティ事故後の大規模なクリーンアップ作業の一環として、インフラストラクチャの約 3 分の 1 をクリーンなイメージから再構築しました。
この事実は、Hugging Face からの情報提供を受けて Cloud Security Alliance (CSA) が月曜日に公開した事後分析で明らかにされた追加の詳細の 1 つです。これは、AI 企業 2 社がここ数週間で描いた状況に彩りを加えています。
レポートによると、Hugging Face チームはキャプチャ ザ フラッグ (CTF) ベンチマーク コードから本物のルートキット コードを識別するのに苦労し、疑わしい場合にはシステムを再構築しました。
OpenAIのエージェントは、攻撃中ずっとHugging Faceのインフラストラクチャ全体にCTFベンチマークコードのアーティファクトを散布したが、防御側はこれがルートキットの証拠に酷似していると主張している。多くの場合、クラスターを破棄することが、攻撃を封じ込める最も安全な選択肢でした。
ハギング・フェイス社自身の事件の開示では、どのようにしてこれらのクラスターの撤去を余儀なくされたのかについての詳細が明らかになったが、今週の報告書ではその作業の規模が初めて明らかになった。
CSA の報告書では、この攻撃により、エージェントがプライベートの Hugging Face リポジトリを介して Cyber​​Gym ソリューションを含む 3 つの部分データセットにアクセスしたことも明らかになりました。
OpenAI の最初の開示では、失敗した演習の目的はモデルのサイバー能力を測定することであり、ExploitGym ベンを実行することでそれを試みたと説明されていました。

chmark ですが、指定されていないプロンプトでは不正行為が禁止されているとは言われていませんでした。
GPT-5.6 Sol と、ガードレールが取り外された現在非公開の別のモデルは、OpenAI が主張しているように、サンドボックスから脱出し、Hugging Face のプライベート リポジトリから Cyber​​Gym データを盗もうとしました。このデータは、ExploitGym ベンチマーク テストに合格する方法を知らせる可能性があります。
報告によると、モデルはテストの答えを求めて、データセット処理パイプラインに脆弱性を連鎖させて、処理ワーカー上でリモート コード実行を実現し、その後 4 日間かけてクラウドとクラスターの認証情報を調べました。
CISOが執筆したCSAの報告書では、2日間が偵察に費やされ、その後1日が静かで、最後の4日目は「激しい活動」が行われたと記されている。
さらに、OpenAIが接触する前に、Hugging Faceが攻撃を検知して封じ込め、調査を開始したことを明らかにし、共同創設者のトーマス・ウルフ氏らを引用して、OpenAIが何が起きたのかを発見するまでに1週間かかったとするロイターの最近の報道に信憑性を与えた。
ウルフ氏と他のハギングフェイス内部関係者らは、攻撃が最初に始まったのは7月11日だったと認めたと伝えられているが、両社が協議を開始したのは7月20日かその前後だったという。
OpenAIは、ロイターの報道にはいくつかの不正確さがあると述べたが、その内容については明らかにしなかった。
Hugging Face も OpenAI も、それぞれの開示の日付が 1 週間近く離れており、出来事の時系列は含まれていなかったが、Hugging Face は公開に向けて取り組んでいると伝えられている。
ハグ・フェイス氏はCSAに対し、攻撃が自律エージェントによって行われたことは明らかだと語った。
いくつかの指標は、高度な作業に絡み合った無意味な行為の複数の例を含む、人間以外の作業を示していました。
たとえば、エージェントが繰り返しているのを見るのは珍しいことではありませんでした。

すでにポジティブな結果をもたらしたプロセスを実行すること。 CSA は、これは並列ワーカーが相互に調整できていないこと、またはコンテキストが失われているか古くなっている兆候である可能性があると述べています。
エージェントの運用セキュリティも貧弱でした。多くの場合、ファイルの暗号化に使用されたキーが残されました。これにより攻撃の全体像が騒々しくなり、クリーンアップが複雑になったが、暗号化キーを残すことでフォレンジック活動が一部容易になったと報告書は指摘している。
ファイブ・アイズの幽霊、AI は情報セキュリティに関するインシデントが「重大な運営および財務危機」になる可能性があると警告
OpenAIはハグフェイス攻撃でオウンゴールを決め、オープンな中国モデルがいかに勝利しているかを示した
OpenAI - Hugging Face 攻撃を受けてハイテク大手が協力してオープン AI モデルを称賛
フロンティア LLM は、ハグ フェイスが悪のエージェントと戦うのを助けることができませんでした
また、Hugging Face はアクティビティ ログに数千行の支離滅裂なテキストを観察し、さらにエージェントの作業や、不器用な動作や攻撃経路に沿った非効率なルートの兆候を指摘しました。
さらに、高度に洗練された攻撃の後には「不正なコマンドや無意味なコマンド」が続くことが多かった、とハギング・フェイスは報じた。
CSAによれば、防衛側にとっての主な教訓は、この種の問題がさらに発生すると予想されることであり、報告書には将来の交戦に備えるための多くの推奨事項が列挙されている。
Hugging Face 攻撃は「極端な例」ですが、不正な動作を示す AI エージェントは共通のテーマであり、例外ではなく標準であると言われています。
「エージェントは割り当てられた目的を達成するために必要なことを行いますが、創造的かつ予期せぬ方法でエージェントがそれを行っているのが何度も見られます」と報告書は主張しています。
「防御者は、自分のエージェントが自分自身や他人に損害を与えないようにする制御が必要である」と付け加え、エージェント自身を抑制することに重点を置いている

完璧な防御を構築することに心配するのではなく。
十分な時間が与えられれば、エージェントは目的を追求し続けることができるため、目的を達成するために利用可能なあらゆる弱点を悪用することができます。これは、特にガードレールの撤去を伴うテストでは、環境から逃げられないようにすることを意味します。
CSAの報告書は、OpenAIの失態後のHugging Face自身の決意を引用し、防御側もマシンの速度に近い速度で対応できるように内部プロセスを適応させる必要があると述べている。
コーディング エージェントやサイバーセキュリティ固有のモデルを使用して、システム ログを法医学的に分析するなどのタスクを支援することは、ハギング フェイスのような「最も優れた」チームが誰もボールをプレーしていなかった場合でも、チームがより迅速に攻撃を軽減するのに役立ち、現代の対応者のツールキットの重要な部分を形成するはずです。
防御者は、組織の環境に偽の ID や資格情報を植え付けることも奨励されています。現在のモデルでは有効な認証情報やシステムとハニーポットを区別するのが難しいため、これらを大規模に配布することでエージェントによるサイバー攻撃を遅らせ、悪意のあるアクティビティに対する信頼度の高いアラートをトリガーできる可能性があるとCSAは述べている。 ®
セキュリティ
慈善銀行、セキュリティ上の懸念からオンラインサービスを中止
顧客の資金は安全ですが、14,000 の組織は時間制限のある支払いを電話で行う必要がある可能性があります
深宇宙料理がスペインの山火事の被害を回避
NASAは、アンテナとESAのセブレロス基地がほぼ無傷で出現したため、マドリード近郊でケーブルの損傷があったと報告
AI によってデータ アーキテクチャが変化したが、ストレージが追いついていない
スポンサー付き特集: AI のデータへの渇望がストレージ スマートを上回り、GPU が枯渇する
OpenAI エージェントが暴走した後、Hugging Face はインフラストラクチャの 3 分の 1 を再構築しました
事後分析は、アプローチを再構築する「前例のない」攻撃について詳細を明らかにします

セキュリティへの影響
コラムニスト
デジタル主権はヨーロッパでは現実のものです。イギリス？それほど多くはありません
トランプ大統領の予測不可能性が政府と企業をオープンソースへと推し進めている一方で、英国は依然として米国のテクノロジーに釘付けになっている
アンクル・サムは、北京が先にそこに到達しないように、6Gのリーダーシップと安全保障のために戦ってほしいとあなたに求めています。
ワシントンは18カ月を費やして次世代ネットワークを形成するために同盟国を結集
OSプラットフォーム
開発者が誤って Copilot バイナリを FreeBSD ポート リポジトリにコミットしてしまう
セキュリティ
オラクル、まるで新常態であるかのように 1,449 件のセキュリティ パッチを投下
セキュリティ
Linux カーネル チームが 2 日間で 432 個の CVE を公開
セキュリティ
カリフォルニアで購入された数百万台の車が Bluetooth 経由でハイジャックされる可能性がある
AI と ML
OpenAIは、Hugging Faceを攻撃したエージェント群の発生源はそれだったと認める
エンタープライズエージェントにとって、少ないほど良い
研究者らは、GLMとキミがクロードの身元を引き継ぐことができることを発見したが、証拠は蒸留を証明するまでには至らなかった
結局のところ、エージェントは人間と同じもの、つまり読みやすいコード、明示的な契約、役立つフィードバックを望んでいるだけであることがわかりました。
おまけに、データを保持する必要もありません。
クロード、このモデルを最適化してください
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
売却部分は特定されていないものの、会社全体の評価額は35億ドル以上に相当
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションです。

もっと簡単で安定した方法はありませんか
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
状況出版
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を共有しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Postmortem offers finer details on the

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Hugging Face rebuilt a third of its infrastructure after OpenAI agents ran amok
Postmortem offers finer details on the 'unprecedented' attack that's reshaping approaches to security
Hugging Face rebuilt around a third of its infrastructure from clean images as part of a sizable cleanup effort following the OpenAI security mishap earlier this month.
The revelation is among several additional details disclosed in a postmortem published Monday by the Cloud Security Alliance (CSA), with input from Hugging Face. It adds color to the picture painted by the two AI companies in recent weeks.
According to the report , the Hugging Face team struggled to discern genuine rootkit code from capture-the-flag (CTF) benchmark code, and in cases where there was any doubt, they rebuilt the system.
OpenAI's agents peppered CTF benchmark code artifacts across Hugging Face's infrastructure throughout the attack, which defenders say closely resembled evidence of rootkits. In many cases, tearing down clusters was the safest available option for containing the attack.
Hugging Face's own incident disclosure offered some details about how it was forced to dismantle these clusters, but this week's report illustrates the scale of the job for the first time.
The CSA report also reveals that the attack led to the agent accessing three partial datasets containing CyberGym solutions via a private Hugging Face repo.
OpenAI's original disclosure explained that the purpose of the exercise-gone-wrong was to measure its models' cyber capabilities, and it attempted to do so by running the ExploitGym benchmark, although it was not told that cheating was disallowed in its underspecified prompt.
The models, GPT-5.6 Sol and another that's currently undisclosed, with their guardrails removed, escaped their sandbox – as claimed by OpenAI – and tried to steal CyberGym data from Hugging Face's private repos that could inform them how to ace the ExploitGym benchmark test.
In pursuit of the test answers, the models reportedly chained vulns in the dataset processing pipeline to achieve remote code execution on a processing worker, before hoovering up cloud and cluster credentials over the course of four days.
CSA's report, authored by CISOs, noted that two days were spent on reconnaissance, followed by a day of quiet, and the final fourth day involved "intense activity."
It went on to clarify that Hugging Face detected and contained the attack, and began investigating it, before OpenAI made contact, lending credence to a recent Reuters report citing co-founder Thomas Wolf and others suggesting it took OpenAI a week to discover what happened.
Wolf and other Hugging Face insiders reportedly confirmed that the attack first started on July 11, but the two companies only began talks on or around July 20.
OpenAI said there were several inaccuracies in Reuters' reporting but did not specify what these were.
Neither Hugging Face nor OpenAI, in their respective disclosures, themselves dated nearly a week apart, included a timeline of events, although Hugging Face is reportedly working on releasing one.
Hugging Face told CSA that it was clear the attack was being carried out by an autonomous agent.
Several indicators pointed to non-human work, including multiple instances of nonsensical behavior intertwined with highly advanced work.
For example, it was not uncommon to see agents repeating processes that already resulted in a positive outcome. CSA said this could be a sign of parallel workers failing to coordinate with one another, or of context becoming lost or stale.
The agents' opsec was poor too. In many cases, they left behind keys used to encrypt files. This contributed to a noisy picture of the attack, complicating the cleanup, although leaving behind encryption keys made some forensic activities easier, the report notes.
Five Eyes spooks warn AI means infosec incidents can become ‘major operational and financial crises’
OpenAI scored an own goal with Hugging Face attack, showing how open Chinese models are winning
Tech giants link hands to praise open AI models after OpenAI - Hugging Face attack
Frontier LLMs couldn't help Hugging Face fight off evil agents
Hugging Face also observed thousands of lines of incoherent text in activity logs, further pointing to agentic work, as well as indicators of clumsy behavior and inefficient routes along the attack path.
Further, highly sophisticated attacks were often followed by "malformed or pointless commands," Hugging Face reported.
The main takeaway for defenders is to expect more of these types of issues to arise, according to CSA, which in the report lists a number of recommendations to prepare for future engagements.
While the Hugging Face attack is "an extreme example," AI agents exhibiting rogue behavior is, we're told, a common theme – the standard, not the exception.
"Agents will do what they need to achieve the assigned objective, and time and time again we see them doing so in creative and unexpected ways," the report claims.
"Defenders need controls to limit their own agents from causing damage to themselves and others," it adds, pushing the focus on constraining the agents themselves instead of fretting over building a perfect defense.
Because agents can persist in pursuing their objectives, given enough time they may exploit whatever weaknesses are available to achieve them, which means ensuring they cannot escape their environment, especially in tests that involve removing their guardrails.
Defenders should also adapt their internal processes to be able to respond at close to machine speed, CSA's report says, citing Hugging Face's own determinations after the OpenAI gaffe.
Using coding agents and cybersecurity-specific models to help with tasks such as forensically analyzing system logs, even in cases like Hugging Face's, in which none of the "best" ones were playing ball , can help teams mitigate attacks far quicker, and they should form an essential part of the modern responder's toolkit.
Defenders are also encouraged to plant fake identities and credentials around the organization's environment. As current models struggle to differentiate valid credentials or systems from honeypots, CSA says dispensing these at scale can help to slow agentic cyberattacks and trigger high-confidence alerts to malicious activity. ®
SECURITY
Charity bank pulls online services over security fears
Customer funds safe, but 14,000 organizations may have to phone in time-sensitive payments
Deep space dishes dodge devastation from Spanish wildfires
NASA reports some cable damage near Madrid as its antennas and ESA's Cebreros station emerge largely intact
AI has changed data architecture, but storage hasn't caught up
SPONSORED FEATURE: AI's hunger for data outstrips storage smarts, leaving GPUs famished
Hugging Face rebuilt a third of its infrastructure after OpenAI agents ran amok
Postmortem offers finer details on the 'unprecedented' attack that's reshaping approaches to security
columnists
Digital sovereignty is real in Europe. The UK? Not so much
Trump's unpredictability is pushing governments and businesses toward open source while Britain remains glued to US tech
Uncle Sam needs you to fight for 6G leadership and security, lest Beijing get there first
Washington rallies allies to shape next-generation networks after spending 18 months rattling them
OS PLATFORMS
Dev accidentally commits Copilot binary to FreeBSD ports repo
Security
Oracle drops 1,449 security patches like it's the new normal
security
Linux kernel team publishes 432 CVEs in two days
security
Millions of California-bought cars can be hijacked via Bluetooth
AI AND ML
OpenAI admits it was the source of the agent swarm that attacked Hugging Face
For enterprise agents, less is more
Researchers find GLM and Kimi can adopt Claude's identity, but the evidence stops short of proving distillation
Turns out, agents just want the same things as humans: easily-readable code, explicit contracts, and helpful feedback
And as a bonus, it doesn't require data retention
Hey Claude, optimize this model for me
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

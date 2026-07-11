---
source: "https://www.wired.com/story/security-news-this-week-ai-found-a-root-bug-in-linux-that-everyone-missed-for-15-years/"
hn_url: "https://news.ycombinator.com/item?id=48873669"
title: "AI Found a Root Bug in Linux That Everyone Missed for 15 Years"
article_title: "AI Found a Root Bug in Linux That Everyone Missed for 15 Years | WIRED"
author: "joozio"
captured_at: "2026-07-11T17:46:26Z"
capture_tool: "hn-digest"
hn_id: 48873669
score: 1
comments: 0
posted_at: "2026-07-11T17:00:52Z"
tags:
  - hacker-news
  - translated
---

# AI Found a Root Bug in Linux That Everyone Missed for 15 Years

- HN: [48873669](https://news.ycombinator.com/item?id=48873669)
- Source: [www.wired.com](https://www.wired.com/story/security-news-this-week-ai-found-a-root-bug-in-linux-that-everyone-missed-for-15-years/)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T17:00:52Z

## Translation

タイトル: AI が 15 年間誰もが見逃していた Linux の根本的なバグを発見
記事のタイトル: AI が 15 年間誰もが見逃していた Linux の根本的なバグを発見 |ワイヤード
説明: さらに: 国防総省はハッカー軍団の一員となるようアマチュアを訓練している、フロックのナンバープレート読み取りエラーにより警察が自動車評論家を取り囲む事態になった、などなど。

記事本文:
メイン コンテンツにスキップ メニュー WIRED セキュリティ ポリシー ビッグ ストーリー ビジネス サイエンス カルチャー レビュー メニュー WIRED アカウント アカウント ニュースレター セキュリティ 政治 ビッグ ストーリー ビジネス サイエンス カルチャー レビュー シェブロン 詳細 ビッグ インタビュー マガジン イベント WIRED Insider WIRED コンサルティング ニュースレター ポッドキャスト ビデオ ライブストリーム グッズ検索 検索 デル キャメロン リリー ヘイ ニューマン セキュリティ 7 月 11 日2026 6:30 AM 今週のセキュリティ ニュース: 15 年間誰もが見逃していた Linux の根本的なバグを AI が発見
ローダー ストーリーを保存 このストーリーを保存 コメント ローダー ストーリーを保存 このストーリーを保存 中国の悪名高いボルト タイフーン ハッカーが米国の重要インフラ内に事前に侵入している可能性があると長年警告されてきた中、保険会社向けの密室戦争ゲームが一連の最悪のシナリオを展開し、脅威的で破壊的な脅威が明らかになりました。
ICEの内部監視グループである職業責任局は、オンラインで同庁を批判する人々の調査を開始し、ICE当局が同庁職員に対する「個人情報の漏洩と脅迫の事件」と呼ぶものについて100件以上の事件を起こした。そして欧州連合では、オンラインの児童虐待コンテンツの抑制を目的とした「チャット規制」法案の新たな権限により、ハイテク企業は国民の個人的なテキストメッセージ、電子メール、ソーシャルメディアメッセージを再びスキャンできるようになる。欧州議会は、議員の過半数がこの提案に反対票を投じたにもかかわらず、法案の延長を可決した。
WIREDは今週、マディソン・スクエア・ガーデンの監視状況についてさらに詳しく明らかにし、MSGが何百人ものセレブリティ、著名なニックスのスーパーファン、さらにはテイラー・スウィフトの結婚式のゲストの一部までを、「LGBTQIA」、「ホスト禁止」、低から高までの「リスク」を含むラベルを使用して分類したデータベースを管理していたという事実を明らかにした。
そして新たな研究はこれ

OnlyFans のコンテンツが「漏洩した」と詐欺師が約束する政府 Web サイトのハイジャックの波が、アダルト コンテンツ作成者からの数千件の著作権侵害の申し立てによって妨げられており、悪意のあるリンクを削除することで人々の安全を守っていることが今週明らかになりました。
さらにあります。私たちは毎週、私たち自身が詳しく取り上げなかったセキュリティとプライバシーのニュースをまとめます。全文を読むには見出しをクリックしてください。そして、外では安全を保ってください。
AI が 15 年間誰もが見逃していた Linux の根本的なバグを発見
SecurityWeek と The Hacker News によると、Nebula Security は GhostLock (CVE-2026-43499) のエクスプロイト コードを公開しました。GhostLock (CVE-2026-43499) は Linux カーネルに 15 年間眠っていた use-after-free バグで、ログインしているユーザーがパッチを適用していないマシン上で root になれるようになります。この欠陥は 2011 年以降、基本的にすべての主流ディストリビューションにデフォルトで出荷されており、特別な許可やネットワーク アクセスは必要ありません。 Nebula のエクスプロイトはコンテナをエスケープし、テストでは 97% の信頼性がありました。 Google の kernelCTF プログラムを通じて 92,337 ドルの支払いを獲得しました。この問題は 4 月に修正されましたが、パッチの入手状況にはばらつきがあります。 Ubuntuは、7月初旬の時点で、24.04、22.04、および20.04 LTSを脆弱性または進行中のものとしてリストしているため、防御側は、修正パッケージが待っていると考えるのではなく、修正パッケージを確認する必要があります。
特に、Nebula は、AI 主導のバグハンティング ツールである VEGA のバグを発見しました。これは、何年も読み返した人がほとんどいなかった古いカーネル コードを精査する自動ツールによって表面化した、2026 年に実行された Linux 権限昇格の欠陥の一部でした。
群れのカメラが記者を追って警官4人を派遣
ジョエル・フェダー記者は『ザ・ドライブ』に寄稿し、6月下旬、ミネソタ州プリマスのコールズ駐車場で4台のパトカーに閉じ込められた――警官らは叫び声を上げ、銃に手を当てた――その理由は、フロックのナンバープレートカメラが彼が試乗していた15万5000ドルのレンジローバーに警告を発していたからだ、と書いている。

ニュージャージー州のディーラーから盗まれたものとして入手した。フェダーさんは、警察が街中で数日間このSUVを追跡していたのは、タイプミスだったと書いている。
原因は、2,000マイル離れた場所でのデータ入力エラーに遡った。34 03 DTMと書かれたジャガー・ランドローバーの車両プレートが紛失としてロサンゼルス警察に届け出られたが、システムには単に「34 DTM」として入力され、ニュージャージー州がメーカープレートで使用している小さな中間の数字が省略されていた。フロックのカメラは大きな文字を読み取り、規格外の構造を無視し、一致する車があれば警察に警告し始めた。フェダーは、同じ週に同じナンバープレート形式を共有する他の4台のランドローバーがミネソタ周辺で追跡されていたと報告した。彼は最初に車を止められただけだった。
この警察活動の発端となったプレートは、まったく盗まれたものではなく、写真撮影中に置き忘れただけであることが判明した。皮肉なことに、この事件は、ザ・ドライブがまさにこの種のフロックの行き過ぎに関するバイラルレポートを発表してからわずか2週間後に発生しました。
ICEのサイバー請負業者が侵害される
BleepingComputer によると、コンサルティング大手の Accenture は、「888」を名乗る攻撃者が 35 GB のデータ (ソース コード、RSA および SSH キー、Azure アクセス トークン、構成ファイル) を持ち出し、サイバー犯罪フォーラムで売りに出したと主張し、セキュリティ侵害を確認したとのことです。アクセンチュアはこれを「孤立した問題」と呼び、情報源を修復したと述べ、運用への影響はないと報告したが、実際に何が盗まれたのか、攻撃者がどのように侵入したのかについてはコメントを控えた。その主張を裏付けるために、888は編集されたAccenture.comホスト上に複製されたAzure DevOpsリポジトリを示すスクリーンショットを投稿した。 BleepingComputer は完全なスコープを検証できませんでした。 888はアクセンチュアの従業員を売ろうとした

2024 年にはサードパーティによるデータ侵害が発生し、アクセンチュアは 2021 年に別途 LockBit ランサムウェアの被害を受けました。
侵害のタイミングは特に厄介だ。アクセンチュアの連邦政府部門は、2021年9月以来、ICEのサイバー防衛およびインテリジェンスサポートサービス契約（24時間365日の脅威監視、侵入検知、政府機関のネットワーク全体でのインシデント対応）を保持している。この約5,650万ドルの業務命令は8月末に期限切れとなり、現在再競争が行われている。
国防総省はアマチュアハッカー軍団を訓練したいが、生活賃金のためではない
DefenseScoop によると、国防総省は今週、Cyber RAP への応募を開始した。これは、サイバー学位や経験のない、学習適性だけを備えた人々を、12 か月のフルタイムの仕事に採用し、同省のネットワークを守る方法を学ぶ有償見習い制度である。米軍の最高情報責任者（CIO）キルスティン・デイビス氏は、「学問の門番」を捨てて「生の適性、愛国的意欲、実践能力」を求めると主張した。しかし、給料は年間2万2584ドルと非常に高額で、脱落者は政府に訓練費用の返済義務がある。
それが社内で人材を育成するための契約だ。もう一つの選択肢はレンタルです。上院軍事委員会の2027年度国防法案の条項により、ピート・ヘグセス国防長官は「請負業者が所有し、請負業者が運営する手段」、つまり政府の恩恵を受けた雇われハッカー部隊を通じてサイバー作戦を実行するパイロットを立ち上げることが可能になるとGovInfoSecurityは報じている。
批評家たちは一線を超えていると考えている。バイデン氏の元ホワイトハウスサイバー当局者ニック・ライザーソン氏は、ハッキング・フォー・ハイヤー計画は「世界的なサイバー不安定の一因」であると述べ、米国が同様の行為を行った中国の請負業者を制裁していると指摘した。
このアイデアには多彩な歴史があります。昨年、アメリカ代表デビッド・シュワイカート

アリゾナ州政府は、詐欺農場マークおよび報復許可法を導入し、大統領が「マークと報復の手紙」（1812年の戦争で米国が最後に私掠船に手渡したものと同じ文書）を発行できるようにし、民間事業者が米国人に対する詐欺に結び付けられた仮想通貨ウォレットを含む外国のサイバー犯罪者の資産を押収することを許可すると、レジスター紙が報じた。この法案は委員会に送られたが、採決には至らなかった。
あなたの受信箱:ケイティ・ドラモンドとWIREDのニュースルームの内部
トランプ氏、媚びたメッセージを披露してザッカーバーグ氏とベゾス氏を嘲笑
ビッグストーリー: ドローンショーでイエスを見つけました
Apple は古い iPhone をより高速に動作させ、より長く使えるようにします
WIREDイベント: ペプシコの一世代に一度の変革
デル・キャメロンはテキサス州出身の調査記者で、プライバシーと国家安全保障を担当しています。彼は複数のプロジャーナリスト協会賞を受賞しており、調査報道に対するエドワード R. マロー賞も共同受賞しています。以前は、ギズモードの上級記者であり、デイリー紙のスタッフライターでもありました... 続きを読む 国家安全保障ウェブサイトの上級記者 リンク
リリー・ヘイ・ニューマンは、情報セキュリティ、デジタルプライバシー、ハッキングを専門とする『WIRED』のシニアライターです。彼女は以前、Slate でテクノロジー記者として働いており、Slate、New America Foundation、アリゾナ州立大学の提携出版物である Future Tense のスタッフ ライターでもありました。彼女の作品 ... 続きを読む シニアライター X
FCC は Burner Phones Plus の抹殺を望んでいる: AI のバグハンティングが Microsoft 史上最大の Patch Tuesday を促進し、ShinyHunters ランサムウェア ギャングが Oracle のゼロデイを悪用するなど。アンディ・グリーンバーグ・ハッカーら、盗まれたマディソン・スクエア・ガーデンのデータプラスを漏洩と主張：サンフランシスコのゲイバーは顔スキャナーを使用、フランスはパランティアを撤退、アップルは計画を立てている

o プライベートメールを変更するなど。 Matt Burgess LastPass ユーザーがデータを盗まれた - またしても: ジョン・ボルトン元国家安全保障問題担当補佐官が機密資料事件で有罪を認め、Microsoft が主要な情報窃取インフラストラクチャの停止を支援など。リリー・ヘイ・ニューマン Apple の Hide My Email サービスがメールを非表示にできない さらに: Scattered Spider 容疑のハッキング メンバーの身柄引き渡し、数十件のナンバー プレート読み取りエラー、インド当局は WhatsApp のユーザー名展開を懸念している。マット・バージェス、物議を醸した従業員追跡プログラムのデータを社内にメタ公開 従業員らは以前、AIモデルをトレーニングするために従業員のキーストロークデータを収集するこの取り組みについて懸念を表明していた。パレシュ・デイブ連邦職員はホワイトハウスのアプリを携帯電話から削除できない 「テストとして削除したら、すぐに戻ってきました」と、ある政府職員は言う。ヴィットリア・エリオット・メタ、内部データ漏洩を受けて従業員追跡プログラムを一時停止 この動きは、同社がこの取り組みから得た潜在的な機密データを社内に漏洩させたままにしていた後に行われた。パレシュ・デイヴ・クロード、ハッカーが米国のほぼすべての音楽フェスティバルのチケットを発行する方法を見つけるのを支援 研究者は、アンスロピックのクロード・オーパス 4.7 を使用して、ロラパルーザからボナルーまでのすべてのフェスティバルで使用されているフロント・ゲートの Web サイトに侵入し、選択したチケットを自由に発行できることを発見しました。アンディ・グリーンバーグ 中国が米国の水道をハッキングしたらどうなるか？水道本管の破裂を調べるために秘密の戦争ゲームに行ってきました。避難した病院。保険会社は非公開のシミュレーションで、中国のボルト・タイフーン・ハッカーによる大規模混乱への対応を実行し、悪夢のようなシナリオを発見した。アンディ・グリーンブ
[切り捨てられた]
カリフォルニア州のプライバシー権
© 2026 コンデナスト。無断転載を禁じます。 WIREDはPOを獲得できるかもしれない

小売業者とのアフィリエイト パートナーシップの一環として、当社のサイトを通じて購入された製品の売上の分配。コンデナストの事前の書面による許可がない限り、このサイトの素材を複製、配布、送信、キャッシュ、またはその他の方法で使用することはできません。広告の選択肢

## Original Extract

Plus: The Pentagon is training amateurs to become part of its hacker army, a Flock license plate reader error led to cops surrounding a car reviewer, and more.

Skip to main content Menu WIRED SECURITY POLITICS THE BIG STORY BUSINESS SCIENCE CULTURE REVIEWS Menu WIRED Account Account Newsletters Security Politics The Big Story Business Science Culture Reviews Chevron More Expand The Big Interview Magazine Events WIRED Insider WIRED Consulting Newsletters Podcasts Video Livestreams Merch Search Search Dell Cameron Lily Hay Newman Security Jul 11, 2026 6:30 AM Security News This Week: AI Found a Root Bug in Linux That Everyone Missed for 15 Years
Loader Save Story Save this story Comment Loader Save Story Save this story Amid years of warnings that China’s notorious Volt Typhoon hackers may be pre-positioning within United States critical infrastructure, a closed-door war game for insurers played out an array of worst-case scenarios—revealing a menacing, disruptive threat.
ICE’s internal oversight group, the Office of Professional Responsibility, has begun investigating online critics of the agency , opening more than 100 cases looking at what ICE officials call “incidents of doxing and threats” against agency employees. And in the European Union, tech companies will be able to scan citizens’ personal texts, emails, and social media messages again because of renewed powers in the “Chat Control” bill aimed at curbing online child abuse material. The European Parliament voted to extend the legislation despite a majority of lawmakers voting against the proposal.
WIRED revealed more about the Madison Square Garden surveillance landscape this week with revelations that MSG kept a database categorizing hundreds of celebrities , prominent Knicks superfans, and even some Taylor Swift wedding guests using labels that included “LGBTQIA,” “DO NOT HOST,” and low to high “risk.”
And new research this week shows that a wave of government website hijacks in which scammers promise “leaked” OnlyFans content are being stymied by thousands of copyright complaints from adult content creators, helping keep people safe by getting the malicious links taken down.
And there’s more. Each week, we round up the security and privacy news we didn’t cover in depth ourselves. Click the headlines to read the full stories. And stay safe out there.
AI Found a Root Bug in Linux That Everyone Missed for 15 Years
Nebula Security has published exploit code for GhostLock (CVE-2026-43499), a use-after-free bug that sat in the Linux kernel for 15 years and lets any logged-in user take root on an unpatched machine, according to SecurityWeek and The Hacker News . The flaw shipped by default in essentially every mainstream distribution since 2011 and needs no special permissions or network access. Nebula’s exploit escapes containers and was 97 percent reliable in testing. It earned a $92,337 payout through Google’s kernelCTF program. It was fixed in April, but patch availability is uneven; Ubuntu, as of early July, still listed 24.04, 22.04, and 20.04 LTS as vulnerable or in progress, so defenders should confirm the fixed package rather than assume one is waiting.
Notably, Nebula found the bug with VEGA, its AI-driven bug-hunting tool, part of a 2026 run of Linux privilege-escalation flaws surfaced by automated tools combing old kernel code few had reread in years.
Flock Cameras Sent 4 Cops After a Reporter
Writing in The Drive , reporter Joel Feder describes being boxed in by four Plymouth, Minnesota, police cars in a Kohl's parking lot in late June—officers shouting, hands on their guns—because Flock's license plate cameras had flagged the $155,000 Range Rover he was testing, loaned from a New Jersey dealership, as stolen. Feder writes that police had been tracking the SUV around town for days … because of a typo.
The cause traced back to a data-entry error 2,000 miles away: a Jaguar Land Rover fleet plate reading 34 03 DTM had been reported to Los Angeles police as lost, but it was entered into the system as just “34 DTM”—dropping the smaller middle digits that New Jersey uses on manufacturer plates. Flock's cameras read the large characters, ignored the nonstandard structure, and started alerting police to any matching car. Feder reports that four other Land Rovers sharing the same license plate format were being tracked around Minnesota that same week; he was just the first pulled over.
The plate that kicked off all this police activity turned out not to have been stolen at all, just misplaced during a photo shoot. Ironically, the incident came barely two weeks after The Drive published a viral report on exactly this kind of Flock overreach.
ICE’s Cyber Contractor Gets Breached
Consulting giant Accenture has confirmed a security breach after a threat actor going by “888” claimed to have lifted 35 GB of data—source code, RSA and SSH keys, Azure access tokens, and configuration files—and put it up for sale on a cybercrime forum, according to BleepingComputer . Accenture called it an “isolated matter,” said it had remediated the source, and reported no impact to operations but declined at the time to comment on what was actually taken or how the attackers got in. To back the claim, 888 posted a screenshot appearing to show a cloned Azure DevOps repository on a redacted Accenture.com host; BleepingComputer couldn't verify the full scope. It's not the actor's first swing at the company—888 tried to sell Accenture employee data after a third-party breach in 2024, and Accenture separately got hit by LockBit ransomware in 2021.
The timing of the breach is particularly awkward: Accenture's federal arm has held ICE’s Cyber Defense and Intelligence Support Services contract —24/7 threat monitoring, intrusion detection, and incident response across the agency's networks—since September 2021, a roughly $56.5 million task order that expires at the end of August and is currently being recompeted.
The Pentagon Wants to Train an Army of Amateur Hackers—but Not for a Living Wage
The Pentagon opened applications this week for Cyber RAP, a paid apprenticeship that recruits people with no cyber degree or experience—just the aptitude to learn—into 12-month, full-time gigs learning to guard the department's networks, according to DefenseScoop . The US military’s CIO, Kirsten Davies, pitched it as ditching “academic gatekeeping” for “raw aptitude, patriotic drive, and hands-on capability.” But the pay is an abysmal $22,584 a year, and those who wash out will owe the government back for the training.
That's the deal for building talent in-house. The other option on the table is renting it. A provision in the Senate Armed Services Committee's FY2027 defense bill would let defense secretary Pete Hegseth spin up a pilot to run cyber operations through “contractor-owned, contractor-operated means”—a government-blessed hacker-for-hire crew, GovInfoSecurity reports .
Critics see a line being crossed. Nick Leiserson, a former Biden White House cyber official, said the hack-for-hire plan “contributes to global cyber instability” and noted the US has sanctioned Chinese contractors for doing the same thing.
The idea has a colorful pedigree. Last year, US representative David Schweikert of Arizona introduced the Scam Farms Marque and Reprisal Authorization Act, which would let the president issue “letters of marque and reprisal”—the same instrument the US last handed to privateers in the War of 1812—authorizing private operators to seize the assets of foreign cybercriminals, including crypto wallets tied to scams against Americans, as The Register reported . The bill went to committee and never got a vote.
In your inbox: Inside WIRED’s newsroom with Katie Drummond
Trump mocked Zuckerberg and Bezos by showing off fawning texts
Big Story: I found Jesus at a drone show
Apple is making your older iPhone run faster and stay alive longer
WIRED event: PepsiCo’s once-in-a-generation transformation
Dell Cameron is an investigative reporter from Texas covering privacy and national security. He's the recipient of multiple Society of Professional Journalists awards and is co-recipient of an Edward R. Murrow Award for Investigative Reporting. Previously, he was a senior reporter at Gizmodo and a staff writer for the Daily ... Read More Senior Reporter, National Security Website Link
Lily Hay Newman is a senior writer at WIRED focused on information security, digital privacy, and hacking. She previously worked as a technology reporter at Slate, and was the staff writer for Future Tense, a publication and partnership between Slate, the New America Foundation, and Arizona State University. Her work ... Read More Senior Writer X
The FCC Wants to Kill Burner Phones Plus: AI bug hunting fuels Microsoft’s biggest-ever Patch Tuesday, ShinyHunters ransomware gang exploits an Oracle zero-day, and more. Andy Greenberg Hackers Claim to Leak Stolen Madison Square Garden Data Plus: Gay bars in San Francisco using face scanners, France quits Palantir, Apple plans to change its private email, and more. Matt Burgess LastPass Users Had Their Data Stolen—Again Plus: Former national security advisor John Bolton pleads guilty in classified-materials case, Microsoft helps take down major infostealer infrastructure, and more. Lily Hay Newman Apple’s Hide My Email Service Fails to Hide Your Email Plus: Alleged Scattered Spider hacking member extradited, dozens of license plate reader errors, and Indian officials are concerned about WhatsApp’s username rollout. Matt Burgess Meta Exposed Data Internally From Its Controversial Employee-Tracking Program Employees had previously raised concerns about the initiative, which involves collecting workers’ keystroke data to train AI models. Paresh Dave Federal Workers Can’t Get the White House’s App Off Their Phones “I deleted it as a test and it came immediately back,” says one government employee. Vittoria Elliott Meta Pauses Employee-Tracking Program Following Internal Data Leak The move comes after the company left potentially sensitive data from the initiative exposed internally. Paresh Dave Claude Helped a Hacker Find a Way to Issue Tickets to Almost Every US Music Festival A researcher found that using Anthropic’s Claude Opus 4.7, he could break into the website of Front Gate—used by every festival from Lollapalooza to Bonnaroo—and freely issue any ticket he chose. Andy Greenberg What Happens if China Hacks the US Water Supply? I Went to a Secret War Game to Find Out Burst water mains. Evacuated hospitals. In a closed-door simulation, insurers played out their response to a mass disruption by China’s Volt Typhoon hackers—and found a nightmare scenario. Andy Greenb
[truncated]
Your California Privacy Rights
© 2026 Condé Nast. All rights reserved. WIRED may earn a portion of sales from products that are purchased through our site as part of our Affiliate Partnerships with retailers. The material on this site may not be reproduced, distributed, transmitted, cached or otherwise used, except with the prior written permission of Condé Nast. Ad Choices

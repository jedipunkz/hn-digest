---
source: "https://consumerrights.wiki/w/User:Louis/New_Orleans_AI_911_call_triage"
hn_url: "https://news.ycombinator.com/item?id=49258378"
title: "New Orleans is using AI to triage 911 calls in case of backlog"
article_title: "User:Louis/New Orleans AI 911 call triage - Consumer Rights Wiki"
author: "basilikum"
captured_at: "2026-08-11T14:13:22Z"
capture_tool: "hn-digest"
hn_id: 49258378
score: 7
comments: 4
posted_at: "2026-08-11T13:52:34Z"
tags:
  - hacker-news
  - translated
---

# New Orleans is using AI to triage 911 calls in case of backlog

- HN: [49258378](https://news.ycombinator.com/item?id=49258378)
- Source: [consumerrights.wiki](https://consumerrights.wiki/w/User:Louis/New_Orleans_AI_911_call_triage)
- Score: 7
- Comments: 4
- Posted: 2026-08-11T13:52:34Z

## Translation

タイトル: ニューオーリンズは、未処理の場合に AI を使用して 911 通報をトリアージしています
記事のタイトル: ユーザー:Louis/New Orleans AI 911 通報トリアージ - 消費者権利 Wiki
説明: オルレアン教区の一部の 911 通報には自動エージェントが人間より先に応答しており、発信者には AI システムと話していることは明示的に伝えられていません。

記事本文:
コンテンツへジャンプ
メインメニュー
消費者の権利 Wiki
検索
検索
外観
アカウントを作成する
2
通話が自動エージェントに到達する仕組み
5
精度主張の証拠
6
音声認識とニューオーリンズの地名
7
緊急対応における AI を管理するポリシー
8
人材不足と資金源
9
2026年8月の報道をめぐる論争
ユーザー : ルイ/ニューオーリンズ AI 911 通報トリアージ
オルレアン教区では、自動エージェントが人間の電話受付より先に一部の 911 通報に応答します [1] が、発信者は AI システムに話しかけていることを明示的に知らされません。 [ 1 ] 2025 年 1 月 1 日のバーボンストリートのテロ攻撃中、 [ 2 ] タイムズ・ピカユーン紙は、OPCD 事務局長カール・ファソルド氏によると、AI エージェントが最初の 2 件の 911 通報に応答してから人間に転送し、約 20 秒間、事件はひき逃げ事故として分類されたと報じました。 [ 1 ] オルレアン教区通信地区 (OPCD) によると、このツールは自動車事故の通報にのみ機能し、すべての電話応対者がすでに通話中の場合と、通報者がすでに報告されている事故現場から約 200 メートル以内にいる場合にのみ機能します。 [ 3 ]
連邦捜査局は2025年1月1日に声明を発表し、中部標準時午前3時15分頃、ニューオーリンズのバーボンストリートである人物がピックアップトラックを運転して群衆に突っ込み、少なくとも10人が死亡、数十人が負傷したと記録した。 [2] FBIは後に捜査最新情報ページで、この攻撃により14人の命が失われたと述べた。 [ 4 ] 同局は、自らが主任捜査機関であり、テロ行為としての攻撃を捜査するためにそのパートナーと協力していると述べた。 [ 2 ]
タイムズ・ピカユーン紙は2026年4月8日、攻撃中に911番線で何が起こったかを報じた。
バーボンストリート中に

テロリストによる攻撃では、AI エージェントが最初の 2 件の 911 通報に応答し、その後人間に転送しました。ファソルド氏によると、この事件は約20秒間、ひき逃げ事故として分類されたという。
通話が自動エージェントに到達する仕組み
911 通報が自動エージェントに送信されるかどうかを決定する 2 つの条件は、発信者が何かを説明する前に解決されます。それは、すべての人間の電話受付者がすでに話中であるかどうか、もう 1 つは既に報告されている衝突事故から約 200 メートル以内からの電話であるかどうかです。 [ 3 ] ファソルド氏は連邦政府のファクトシートで同じ 2 つの条件を提示し、ソフトウェアはすべての通信者 (業界用語で電話受付者を意味する) が占有されているときに介入し、ジオフェンシングによってそうする、つまり場所の周囲に境界線を引き、その内部からの通話に応答する、と述べた。 [ 5 ]
Carbyne のソフトウェアは OPCD のコンピュータ支援ディスパッチ システムに接続されており、どの事故が進行中であるかを表示するため、どの衝突がすでに記録されトリアージ対象としてマークされているかを知ることができます。 [ 6 ]
1 回の呼び出しでのイベントの順序は次のように実行されます。
発信者は 911 にダイヤルします。人間の電話受付者が空いている場合は、人間が応答します。 [ 3 ]
すべての電話受付者が話中で、すでに報告されている衝突事故の周囲から電話がかかってきた場合は、代わりに自動エージェントが応答します。 [ 3 ] [ 1 ]
エージェントは発信者に、その事件について電話しているかどうか、そして自分が事件に関与したかどうかを尋ねます。 [ 6 ]
発信者が「はい」と答えると、すべてが処理されていることが伝えられます。関与しておらず、センターがすでに知っていることを報告している発信者には、電話を切る可能性があると言われます。 [ 1 ] [ 6 ]
「ノー」と答えた発信者および関与した発信者は、人間の電話受付者に接続されます。 [ 1 ] [ 6 ]
したがって、エージェントは通話に応答した後に質問を行い、発信者自身の応答によって通話が担当者に送信されます。

息子。 [ 6 ] [ 1 ]
タイムズ・ピカユーン紙は2026年4月8日に次のように報じた。
オルレアン教区の 911 番と 311 番の発信者は、AI システムに話しかけていることを明確に知らされませんが、その声は明らかに人間のものではなく、通信地区は 311 AI エージェントがいつ稼働するかを発表する予定はありません。
ニュース報道をWWLルイジアナと銘打っているニューオーリンズのテレビ局WWL-TVは、2026年4月8日、OPCD広報担当者のジェイ・バイス氏が、提案されている311アシスタントの保護措置について、不確実な要請をある人物にルーティングするものであると説明したと報じた。
誰かが何を探しているのかについて少しでも質問すると、その人はキューに送られます。ライブオペレーターが応答します。
バイス氏は、暴力犯罪や命を脅かす電話の場合、最初に聞く声は常に人間であると付け加えた。 [ 7 ] タイムズ・ピカユーン紙は、911通報者が「はい」と答えると、すべて対応中であると言われたと報じた。 [ 1 ]
Carbyne の Call Triage ソフトウェアは、地区のコンピュータ支援ディスパッチ システムに接続しているため、エージェントはどのインシデントがすでに記録されているかを確認できます。 [ 6 ] StateScoop は 2024 年 7 月 5 日に、トリアージ対象の事件から 100 メートル以内にいる発信者は、最初に自動システムによってその事件について電話しているかどうか、そして自分が関与したかどうかを尋ねられると報告しました。 [6] 2026年4月8日までに、タイムズ・ピカユーン紙はその半径を200メートルとし、自動車事故の現場から半径200メートル以内で911に通報した者は直ちにAIエージェントにルーティングされ、AIエージェントは発信者が事故を報告しているかどうかを尋ね、答えがノーの場合は人間につなぐと書いた。 [ 1 ]
ファソルド氏は、国家電気通信情報局が2024年8月2日に発行した連邦ファクトシートでトリガー条件について次のように説明した。
すべての通信者が占有されている場合、AI は

通話のトリアージ、特に既知のインシデントのエリアでの通話を管理するジオフェンシングに参加します。
当初は範囲がさらに狭かったです。 StateScoopの報道によると、オーリンズ教区は当時、自動車事故のみに一度に1つの事件に対してこの機能を使用しており、ファソルド氏は複数の事件に対してこの機能の使用を開始する準備ができており、他の非暴力事件の種類にも拡張したいと述べた。 [ 6 ]
Times-Picayune 紙は、カーバイン社が最近、地区の従来の 911 システムをクラウドネイティブの 911 製品にアップグレードし、AI エージェントに取り組んでいること、そしてファソルド氏のスタッフが電話の急増にしばしば圧倒され、その後カーバイン社と協力して急増の原因を分析したと報じた。 [1] カーバイン社の最高技術責任者兼共同創業者のアレックス・ディゼンゴフ氏は同紙に対し、データから繰り返されるパターンを特定したと語った。毎朝と夕方のラッシュアワーには事故が発生し、何百人もの発信者がほぼ同じ内容の電話をかけてきたという。 [ 1 ]
精度主張の証拠
OPCD が報道機関に提供した精度の数値は、Fox 8 として放送するニューオーリンズのテレビ局 WVUE とルイジアナ州 WWL に説明された、OPCD 自身のスタッフが行ったレビューに基づくものです。 [ 8 ] [ 7 ] ファソルド氏はFox 8に次のように語った。
最初の 3 か月間は通話を 100% 聞きました。誤検知はなく、また、誤検知もありませんでした。つまり、AI は処理すべき通話を処理し、処理すべきでない通話は処理しませんでした。
Vise は、ルイジアナ州 WWL に 911 配備の試験段階をカバーする説明を提供し、同じ記事では、自動車事故の再発通報に対する 1 年強のテストとして説明しました。
私たちはすべての通話、すべてのやり取りを監査し、100% が適切に処理されたと言えることを嬉しく思います。
システムが保存するかどうかについて

フォックス 8 の報道によると、ファソルド氏はいかなる証拠も逸話に過ぎないと述べた後、心臓発作に関する仮定の 911 通報に応じた。
では、AIは命を救ったのでしょうか？心停止や脳損傷などにかかる時間を考えると、私は「そうだ」と答えるでしょう。 AI が命を救っていると言えるかもしれません。 「派手な」例を挙げてもらえますか？いいえ、残念ながらそうではありません。
チューレーン大学フリーマン・スクール・オブ・ビジネスの経営学教授ロブ・ラルカ氏はFox 8に対し、ニューオーリンズの911センターの場合、間違いは許されないと語った。
たった一度でもミスをすれば、たった一度でも間違えれば、誰かの命がかかってしまう、それは危険です。住民にサービスを提供しない都市は望ましくありません。
ニューオーリンズへの導入に関するカービン社独自の日付不明のケーススタディによると、交通事故のようなシナリオでは、この製品は 20 件の通話のうち 6 件を自動的にトリアージすることが期待でき、これにより人員時間を最大 16 分節約できる可能性があると同社は述べています。 [ 9 ] この文書の他の数字も、同社の自社製品に関するマーケティング上の主張です。ケーススタディでは、システムが 90 日間にわたって 3,500 件を超えるイベント (1 日あたり平均 40 件以上のイベント) をトリアージしたこと、トリアージされた通話のうち冗長な通話が 30% 以上削減されたこと、および、困っている発信者の電話への応答が平均で最大 40 秒早くなったことが記載されています。 [ 9 ]
ファソルド氏によると、タイムズ・ピカユーン紙は、911 AI エージェントは自動車事故関連の電話にのみ応答し、実際の緊急電話には決して使用されないと報じた。 [ 1 ]
音声認識とニューオーリンズの地名
2023 年の展開前に、OPCD スタッフはプログラムに実際の 911 通報を 3 か月間提供して、実際の通報に応答できるように学習しました。 [ 8 ] ファソルドは、トレーニング期間は楽しくもあり、イライラするものでもあったと次のように述べた。
私たちのアクセント、膨らむ

ニューオーリンズの行動は他の多くの場所とは異なります。コンピューター生成の音声に「Tchoupitoulas」と言うように教え、AI に反対側の「Tchoupitoulas」を理解して正しく発音するように教えるのはとても楽しかったです。このプロセス全体にとって、それは簡単でした。
発表された研究では、商用音声認識がどの程度の範囲を見逃す可能性があるかを測定しました。 Koeneckeらは、2020年3月23日に米国科学アカデミー紀要に寄稿し、テストした5つの自動音声認識（ASR）システム、つまり録音された音声をテキストに変換するシステムすべてが、米国の5都市からの19.8時間の音声コーパス全体で、平均単語誤り率が黒人話者で0.35、白人話者で0.19であり、大幅な人種格差を示したことを発見した。 [10] 著者らは、基礎となる音響モデルまでのギャップを追跡し、コーパス内の黒人と白人の個人が話した同一のフレーズのサブセットで同じ大きさであることを発見しました。 [ 10 ] この研究では、カービン社のシステムもニューオーリンズでの展開もテストされていません。 [ 10 ]
緊急対応における AI を管理するポリシー
Times-Picayune は 2026 年 4 月 8 日に、緊急対応における AI の使用を管理する国の政策はほとんどなく、チューレーンのコンピューターサイエンス教授アーロン・キュロッタ教授とニック・マッテイ教授、つまりファソルド氏が認識している市の政策も存在しないと報じた。 [ 1 ] キュロッタ氏は、透明性と説明責任を求める責任は国民にあると述べた。 [ 1 ] マテイ氏はタイムズ・ピカユーン紙に次のように語った。
しかし、非常によく封じ込められたアプリケーションの場合、これは政府のリソースを解放する強力な方法になると思います。 ...AIだからといって悪いわけではありません。
ただし、AI がいつどのように動作するかについて最終的な決定を下すのは、適切な制御が必要です。
人材不足と資金源
OPCD には 42 人の電話受付者がいた

2026 年 7 月 16 日の時点で、コールテイカー 64 名とディスパッチャー 68 名という理想に対して、S とディスパッチャーは 35 名です。 [ 8 ] ファソルド氏は、OPCDのスタッフは2024年1月に州内で最も高給取りの911従業員になったが、そのOPCDがあってもまだそこで働くのに十分な人材を確保できないと述べた。 [ 8 ] StateScoop は 2024 年に、電話受付ポジションの約 3 分の 1 が予算化されており、承認されていても埋まっていないと同氏が述べたと報じた。 [ 6 ]
不足はニューオーリンズに特有のものではない。ステートスクープは、カービンと全米緊急電話番号協会が発表した調査結果を報告し、センターの82パーセントが慢性的な人員不足を報告していることを示し、センターは15秒以内に通話の90パーセントに応答することを義務付けており、20秒以内に95パーセントに応答するという奨励目標に反して、センターの82パーセントが慢性的な人員不足を報告していることを示した。 [ 6 ]
ルイジアナ州法は、OPCD に回線ごとの電話料金を課すことを認めています。 [ 11 ] La.R.S.の下で33:9106.2、オルレアン教区通信地区の管轄当局は、固定料金の緊急電話サービス料金を設定することができます。これらの料金は、住宅サービス ユーザーにサービスを提供する交換アクセス回線あたり 1 ドル、商用サービス ユーザーにサービスを提供する交換アクセス ラインあたり 2 ドルを超えてはなりません。

[切り捨てられた]

## Original Extract

An automated agent answers some 911 calls in Orleans Parish before a human does, and callers are not explicitly told they are speaking to an AI system.

Jump to content
Main menu
Consumer Rights Wiki
Search
Search
Appearance
Create account
2
How a call reaches the automated agent
5
Evidence for the accuracy claims
6
Speech recognition and New Orleans place names
7
Policy governing AI in emergency response
8
Staffing shortfall and funding sources
9
August 2026 dispute over press coverage
User : Louis/New Orleans AI 911 call triage
An automated agent answers some incoming 911 calls in Orleans Parish before any human call taker does, [ 1 ] & callers are not explicitly told when they are speaking to an AI system . [ 1 ] During the Bourbon Street terrorist attack of January 1, 2025, [ 2 ] The Times-Picayune reported that an AI agent answered the first two 911 calls before routing them to a human, & that for roughly 20 seconds the incident was classified as a hit-and-run crash , according to OPCD Executive Director Karl Fasold. [ 1 ] The Orleans Parish Communication District (OPCD) says the tool engages only for motor vehicle accident calls, only when every human call taker is already busy, & only when the caller is within about 200 meters of a crash that has already been reported. [ 3 ]
The Federal Bureau of Investigation issued a statement on January 1, 2025 recording that at approximately 3:15 a.m. CST an individual drove a pickup truck into a crowd of people on Bourbon Street in New Orleans, killing at least 10 & injuring dozens of others. [ 2 ] The FBI later stated on its investigation updates page that the attack took the lives of 14 people. [ 4 ] The bureau said it was the lead investigative agency & was working with its partners to investigate the attack as an act of terrorism. [ 2 ]
The Times-Picayune reported on April 8, 2026 what happened on the 911 line during the attack:
During the Bourbon Street terrorist attack, an AI agent answered the first two 911 calls before routing them to a human. For roughly 20 seconds, the incident was classified as a hit-and-run crash, Fasold said.
How a call reaches the automated agent
Two conditions decide whether a 911 call goes to the automated agent, & both are settled before the caller has described anything: whether every human call taker is already busy, & whether the call is coming from within about 200 meters of a crash that has already been reported. [ 3 ] Fasold gave the same two conditions in the federal fact sheet, saying the software steps in when all telecommunicators, the industry term for call takers, are occupied, & that it does so by geofencing, meaning drawing a boundary around a location & acting on calls that come from inside it. [ 5 ]
Carbyne's software is connected to OPCD's computer-aided dispatch system, which shows it which incidents are ongoing, so it can tell which crashes have already been logged & marked for triage. [ 6 ]
The order of events on a single call runs as follows.
A caller dials 911. If a human call taker is free, a human answers. [ 3 ]
If every call taker is busy & the call comes from inside the radius around an already-reported crash, the automated agent answers instead. [ 3 ] [ 1 ]
The agent asks the caller whether they are calling about that incident & whether they were involved in it. [ 6 ]
A caller who says yes is told everything is being handled; callers who were not involved & are reporting what the center already knows are told they may hang up. [ 1 ] [ 6 ]
A caller who says no, & any caller who was involved, is connected to a human call taker. [ 1 ] [ 6 ]
The agent therefore puts its question after it has answered the call, & the caller's own answer is what sends the call on to a person. [ 6 ] [ 1 ]
The Times-Picayune reported on April 8, 2026:
Callers to 911 and 311 in Orleans Parish are not explicitly told when they are speaking to an AI system, though the voice is obviously not human, and the communications district does not plan to announce when the 311 AI agents go live.
The New Orleans television station WWL-TV, which brands its news coverage as WWL Louisiana, reported on April 8, 2026 that OPCD public information officer Jay Vise described the safeguard on the proposed 311 assistant as routing an uncertain request to a person:
Any question at all about what someone's looking for, they get routed to the queue. They'll be answered by a live operator.
Vise added that for violent crimes & life-threatening calls, Those will always be a human being, the first voice you hear . [ 7 ] The Times-Picayune reported that a 911 caller who answers yes is told that everything is being handled. [ 1 ]
Carbyne's Call Triage software connects to the district's computer-aided dispatch system so the agent can see which incidents are already logged. [ 6 ] StateScoop reported on July 5, 2024 that callers within 100 meters of an incident marked for triage are first asked by an automated system whether they are calling about that incident & whether they were involved. [ 6 ] By April 8, 2026, The Times-Picayune put that radius at 200 meters, writing that anyone calling 911 within a 200-meter radius of a live auto accident is immediately routed to the AI agent, which asks whether the caller is reporting the wreck & connects the caller to a human if the answer is no. [ 1 ]
Fasold described the trigger condition in a federal fact sheet the National Telecommunications and Information Administration published on August 2, 2024:
When all telecommunicators are occupied, AI steps in for call triage, particularly for geofencing to manage calls in the area of a known incident.
The scope was narrower still at the start. StateScoop reported that Orleans Parish was then using the function for motor vehicle accidents only, & for one incident at a time, with Fasold saying he was ready to begin using it for multiple incidents & wanted to extend it to other nonviolent incident types. [ 6 ]
The Times-Picayune reported that Carbyne had recently upgraded the district's legacy 911 system to a cloud-native 911 product & was working on an AI agent, & that Fasold, whose staff often became overwhelmed with surges of calls, then worked with Carbyne to analyze where the spikes were coming from. [ 1 ] Carbyne chief technology officer & cofounder Alex Dizengof told the paper that they identified a repeated pattern coming from the data: every morning & evening during rush hour there would be an accident, & hundreds of callers called in about the same thing. [ 1 ]
Evidence for the accuracy claims
The accuracy figures OPCD has given the press come from review its own staff performed, described to the New Orleans television station WVUE, which broadcasts as Fox 8, & to WWL Louisiana. [ 8 ] [ 7 ] Fasold told Fox 8:
We listened to 100% of the calls for the first three months. We had no false positives, [and] we had no false negatives, meaning the AI handled the calls it should've handled and didn't try to handle any it shouldn't have tried to handle
Vise gave WWL Louisiana an account covering the trial phase of the 911 deployment, which the same article described as a little over a year of testing for repeat motor vehicle accident calls:
We audited every call, every interaction, and we're happy to say 100 percent were processed properly
On whether the system saves lives, Fox 8 reported that Fasold said any evidence would be anecdotal, then walked through a hypothetical 911 call about a heart attack:
So, has the AI saved a life? Given the time involved in cardiac arrest, brain damage, etc. I would say yes. You could qualify and say the AI is saving lives. Can I give you a 'showy' example of it? No, unfortunately not.
Rob Lalka, a business professor at Tulane University's Freeman School of Business, told Fox 8 that in the case of the New Orleans 911 center there is no room for mistakes:
If you miss just one case, if you get it wrong just once, someone's life is on the line, and that's dangerous. You don't want a city that's not serving its people.
Carbyne's own undated case study of the New Orleans deployment states that in scenarios like traffic accidents the product can be expected to triage six out of 20 calls automatically , which the company says has the potential to save up to 16 minutes of personnel time. [ 9 ] The other figures in that document are also the company's marketing claims about its own product: the case study states that over a 90-day period the system triaged more than 3,500 events, averaging over 40 events per day, that it reduced redundant calls by over 30 percent among triaged calls, & that distressed callers had their calls answered up to 40 seconds quicker on average. [ 9 ]
The Times-Picayune reported that the 911 AI agent only answers car accident-related calls, & that it will never be used for real emergency calls, according to Fasold. [ 1 ]
Speech recognition and New Orleans place names
Before rollout in 2023, OPCD staff fed the program three months of actual 911 calls so it could learn to answer real ones. [ 8 ] Fasold described the training period as both entertaining and frustrating:
Our accents, inflections in New Orleans are different than a lot of places. Teaching the computer-generated voice to say 'Tchoupitoulas' and teaching the AI to understand 'Tchoupitoulas' on the other side and get it right has been lots of fun, and that was an easy one for this whole process
Published research has measured how far commercial speech recognition can miss. Koenecke et al., writing in the Proceedings of the National Academy of Sciences on March 23, 2020, found that all five automatic speech recognition (ASR) systems tested, meaning systems that turn recorded speech into text, exhibited substantial racial disparities, with an average word error rate of 0.35 for black speakers compared with 0.19 for white speakers, across a corpus of 19.8 hours of audio from five US cities. [ 10 ] The authors traced the gap to the underlying acoustic models, finding it equally large on a subset of identical phrases spoken by black & white individuals in the corpus. [ 10 ] That study tested neither Carbyne's system nor any deployment in New Orleans. [ 10 ]
Policy governing AI in emergency response
The Times-Picayune reported on April 8, 2026 that there is little national policy governing AI use in emergency response, & no city policy that Tulane computer science professors Aron Culotta & Nick Mattei, or Fasold, are aware of. [ 1 ] Culotta said the responsibility falls on the public to demand transparency & accountability. [ 1 ] Mattei told The Times-Picayune:
But for very well-contained applications, I think this can be a powerful way to free up government resources. ... Just because it's AI doesn't mean it's bad.
But it has to come with proper controls, where ultimately a human makes the final decision about when and how AI will act.
Staffing shortfall and funding sources
OPCD had 42 call takers & 35 dispatchers as of July 16, 2026, against an ideal of 64 call takers & 68 dispatchers. [ 8 ] Fasold said OPCD's staff became the best paid 911 employees in the state in January 2024, & that even with that OPCD still cannot get enough people to work there. [ 8 ] StateScoop reported in 2024 that he said roughly a third of his call-taking positions were budgeted, authorized, & unfilled. [ 6 ]
The shortfall is not unique to New Orleans. StateScoop reported survey results published by Carbyne & the National Emergency Number Association showing 82 percent of centers reported chronic staffing shortages, against a NENA standard requiring centers to respond to 90 percent of calls within 15 seconds & an encouraged target of answering 95 percent within 20 seconds. [ 6 ]
Louisiana law authorizes OPCD to levy a per-line telephone charge. [ 11 ] Under La. R.S. 33:9106.2, the governing authority of the Orleans Parish Communication District may establish a fixed rate emergency telephone service charge, & those charges shall not exceed one dollar per exchange access line serving residential service users & two dollars per exchange access line serving commercial service users, the commercial figure not to e

[truncated]

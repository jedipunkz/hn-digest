---
source: "https://takes.jamesomalley.co.uk/p/build-gemini-build"
hn_url: "https://news.ycombinator.com/item?id=48568646"
title: "How the UK government is using AI to speed up the planning system"
article_title: "Here's how the government is using AI to speed up the planning system"
author: "writerJames"
captured_at: "2026-06-17T13:25:04Z"
capture_tool: "hn-digest"
hn_id: 48568646
score: 1
comments: 1
posted_at: "2026-06-17T11:09:00Z"
tags:
  - hacker-news
  - translated
---

# How the UK government is using AI to speed up the planning system

- HN: [48568646](https://news.ycombinator.com/item?id=48568646)
- Source: [takes.jamesomalley.co.uk](https://takes.jamesomalley.co.uk/p/build-gemini-build)
- Score: 1
- Comments: 1
- Posted: 2026-06-17T11:09:00Z

## Translation

タイトル: 英国政府は計画システムを高速化するために AI をどのように活用しているか
記事のタイトル: 政府が計画システムを高速化するために AI をどのように活用しているか
説明: これら 2 つの新しいシステムは真に革命的になる可能性があります

記事本文:
政府が計画システムを高速化するために AI をどのように活用しているかは次のとおりです
政府が計画システムを高速化するために AI をどのように活用しているかは次のとおりです
これら 2 つの新しいシステムは真に革命的になる可能性があります
James O'Malley 2026 年 6 月 17 日 13 2 1 POD をシェア！今週の YIMBY ポッドでは、マーティンと私は食品の価格に関するザック ポランスキーの主張を考察し、私は、えっと、大手スーパーマーケットを擁護することで人気を博しています。次に、優秀なトーマス・エイブルマンに、スイス人ならまったく違う方法で HS2 を構築したであろうことについて話します。ここで、またはポッドを入手できる場所で聞いてください。
英国で物を作るのは悪夢だ。
議論はすでに十分にリハーサルが行われている。私たちの計画システムは非常に官僚的で、アプリケーションは時には数千ページに及ぶこともあり、最も思慮深い開発でさえ、少数の気難しい議員の命令で無効になる可能性があります。
昨年末に国王の同意を得た政府の計画・インフラ法は、間違いなく正しい方向への飛躍であり、私はその重要性を軽視したくない。これは重要な法律であり、環境緩和を改革し、地方開発法人設立の法的根拠を確立し、多くの小規模な計画申請を委員会の会議から議会の計画担当官の裁量に移しました。
しかし基本的に、これらは既存のシステムのアップグレードであり、特定のゾーンの基準を満たしている限り、建物がデフォルトで事実上承認されるゾーニングに基づくシステムのような、より大胆なものには移行していません。
これは、英国で何かを建てたい場合、たとえロフトの増築やガーデンオフィスのような単純なものであっても、議会に大量の書類を提出し、その後決定を待つ必要があることを意味します。

始める前に。
これは、長く不確かな待ち時間を伴う可能性があるため、非常にイライラします。公式には、小規模/小規模世帯の申請に関する決定までの所要時間は 8 週間であると法定目標が定められていますが、毎年行われる約 35 万件の申請を精査し承認するために議会が用意している計画担当官は一定数しかいないため、実際にはさらに時間がかかる可能性があります。
そして、この低迷の影響を過小評価すべきではありません。これは本質的には建設税であり、自らが経済に与えた打撃だ。建設されるものが少なくなり、生産性が低下し、投資や融資に充てられる資金が圧迫されてしまいます。
ただし、これらすべてが、少なくとも少しずつ変化しようとしている可能性があります。本日、政府は計画担当官がより迅速な意思決定を行えるようにするための 2 つの新しいプロジェクトを発表しました。 1 つは住宅・コミュニティ・地方自治体省 (MHCLG) と科学・イノベーション・技術省 (DSIT) によって社内で構築されており 1、もう 1 つは Google および Faculty AI との 820 万ポンドの共同研究です。
そして幸運なことに、私は 2 つのシステムがどのように機能するかについて、舞台裏の詳細をいくつか知ることができました。
計画申請の処理における根本的なボトルネックは、議会の計画担当者が使える時間です。 MHCLG によると、市議会の企画部門の平均的な人数は約 40 名で、そのうち 2 名は多そうに聞こえるかもしれませんが、やるべきことがたくさんあり、仕事の多くは退屈な単調な仕事です。
たとえば、既存の歴史的計画文書の多くはまだデジタル化されておらず、依然としてファイルキャビネットに紙で保管されています。あなたが計画担当者で、アーカイブを参照する必要がある場合、これは問題です。したがって、申請の検討を始める前に、既存の書類をデジタル化する必要があります。
面倒なことに、これは手続きです

計画担当者は紙の地図と関連文書をスキャンし、手動で情報を転記し、マウスを使用して詳細なデジタル地図を慎重に描画する必要があるため、これにはおそらく数時間かかる可能性があります。
しかし、ここで「Extract」と呼ばれる最初のツールが登場します。
本当の計画部門のファイリングシステムであるが、政府は公に恥をかかせたくなかったため、どのシステムであるかは明らかにしなかった。 Extract は MHCLG のチームによって社内で構築されました。デジタル化プロセス全体を自動化し、紙の文書や手描きの地図を、最新の地図作成ツールや計画ツールで使用できるデジタル オブジェクトに変換するように設計されています。
クイーンズ クラブ ガーデンズ、「ペプシは大丈夫?」ロンドンのテニス会場。たとえば、上はロンドンのクイーンズ クラブ ガーデンズの計画のスキャンです。 1981 年に、建物 3 で変更できる内容を制限する「第 4 条」の指定を受けました。これは、計画担当官がその地域の申請に関して決定を下さなければならない場合に関連する制限です。 Extract に入力すると、わずか数分で最新のデジタル シェイプ ファイルに変換されました。
そして、これを確実に実行する方法は信じられないほど巧妙で、複数のステップのパイプラインがあり、複数の AI モデルが機能して平面描画を有用なものに変換します。
そのため、まず、Google の Gemini モデルを通じて PDF スキャンが実行され、日付やその地域に関するその他の詳細などのテキスト情報が抽出されます。光学式文字認識 (OCR) は古くから存在しているため、原理的には簡単に聞こえるかもしれませんが、実際には、多くの計画文書は乱雑です。手書きのものもあり、タイプされたものの多くは、端に手書きのメモが走り書きされたり、ペンで線が取り消されたりしていることがよくあります。
しかし、どこで

これは古いソフトウェアでは失敗する可能性がありますが、最新の大規模言語モデルではそれほど問題ではありません。
とにかく、テキストの後に地図もAIによって識別されます。次に、それらは PDF から切り出され、画像を取得してその中のさまざまなオブジェクトを識別できる専門の AI モデルである Meta の Segment Anything に供給されます。このようにして、システムは地図上の形状を抽出します。たとえば、上の 4 条地域の周囲の形状や、下の地図上の住宅の形状などです。
セグメント分析AIが稼働中。基本的には魔術です。しかし、システムが地図上に図形をプロットできない場合、図形は何の役に立つでしょうか?そのため、次の Gemini は紙の地図を見て、通りの名前やその他の地理的特徴などを抽出するために使用されます。これらはその後、Google Maps と Ordnance Survey API に入力され、地図があるべき場所を正確に特定します。
ハムステッドにある池。なぜ試験会場があんなに高級な場所ばかりだったのか、私にもわかりません。そしてこの時点でも、Extract の仕事はまだ終わっていません。形状ファイルをデジタル マップ上に正確に配置する必要があるためです。したがって、ここでは Extract が MatchAnything と呼ばれる別の専門 AI モデルを使用しています。これは、あらゆるものの共通点を識別するようにトレーニングされています。
どうやら、これは、たとえば、同じ物体を異なる角度から撮った写真の共通点を特定することができるようで、ここでさらに関連するのは、ハムステッドの上記の例のように、同じ場所の 2 つの地図で、一方の地図がデジタルで、もう一方の地図が手描きで逆さまに描かれている場合に、その共通点を特定できるということです。 4
これらの点を取得すると、抽出された形状上のさまざまな点のそれぞれの経度と緯度を計算するのが非常に簡単になります。つまり、Extract では、元のスキャンを数字の上に重ね合わせることができます。

アルマップ、デジタルバージョンと一致するように画像を歪めます。
最終的に完全に一致しない場合は、計画担当者が作成された形状を手動で編集して、形状が正確であることを確認します。
これは使ってみるとかなり楽しそうです。これは、水面下で多くのことが起こっていることを示しています。しかし、驚くべきことに、スキャンを使用可能なデジタル オブジェクトに変換する平均処理時間は 1 分 42 秒で、文字通り数時間かかっていたプロセスがほぼ瞬時に完了するということです。 5
したがって、Extract が説明どおりに機能すると仮定すると、それ自体で生産性が大幅に向上することになります。しかし、これは戦いの始まりにすぎません。すべての関連文書が揃ったら、計画担当者がそれらの文書に目を通す必要があるため、それ自体に時間がかかります。
計画申請書を作成するとき、開発者は複雑さに応じて、計画の概要を説明した数十ページ、場合によっては数百ページを提出する必要があります。彼らは、何を建設したいのか、そしてそれが地域に適用される規制にどのように準拠しているのかを説明する必要があります。
さらに、ナチュラル・イングランドのような法定顧問など、他の利害関係者からの提案や、建設を支持または（おそらく）反対したい他の住民からの寄付もしばしばあります。
計画申請に必要な書類の数は膨大です。 (出典: Lichfields ) これらの文書はすべて計画担当者の仮想デスクに置かれます。その仕事は、アプリケーションに記載されている内容と、数千ページにわたる計画システムの多数のルールブックに含まれる内容をふるいにかけて比較することです。
たとえば、計画の決定では、国家計画政策枠組み、評議会独自の地方計画、または仕様を考慮する必要がある場合があります。

クイーンズ クラブ ガーデン周辺の第 4 条地域など、建物の外観や雰囲気を制限する保護地域周辺のローカル ルール。
したがって、事実上、判断を下すために、企画担当者はミニチュアの文献レビューを実施し、ガイドラインと提案を比較して、開発が非常に複雑なテストに合格するかどうかを判断する必要があります。
しかし、ここに AI の 2 番目のチャンスがあります。これは Extract よりもはるかに初期の段階にありますが、Google と学部は現在政府と協力して、Augmented Planning Decisions (APD) と呼ばれるシステムのプロトタイプを作成しています。
ここでのアイデアは、大規模言語モデル (この場合は Gemini 3 Pro) がすでに得意としている比較テキスト分析を利用しながら、このスキルをより構造化された厳密な方法でドキュメントの計画に適用することです。そのため、計画では、APD が計画申請書の数百ページを取得し、審査が必要な数千ページのガイダンスと比較し、計画担当者が最終決定を下すために知っておく必要があるすべてを含む概要を提示する予定です。
これは聞いた説明に基づいて私が作成した UI のモックアップです。しかし、それがどのように機能するかを知るのに役立つと思いました。上記に少し似たものをモックアップしてみました。 APD は基本的にスマート ケース管理システムです。任意のアプリケーションをクリックすると、満たす必要のあるさまざまな基準すべてに分けられた AI の推論を確認できます。各基準には、導かれた結論を正当化する特定のガイダンスや法律への深いリンクが含まれています。
したがって、警察官がしなければならないことは、選択を検討し、AI の推奨を受け入れるかどうかを決定することだけです。そして、たとえ彼らがそれに同意しないとしても、

AI が結論を導き、それに同意する前に決定を変更しても、システムは何千ページものガイダンスを手動でスクロールする場合に比べて、大幅な時間を節約します。
この時点で、あなたの AI への熱意のレベルに応じて、かなり興奮しているか、かなり懐疑的であるかのどちらかになります。
とても興奮したキャンプ中です。
なぜなら、Extract と APD の両方の設計方法の賢明な点は、あなたや私が計画 PDF を ChatGPT (または、Gemini) に叩き込んで、「それで、これについてどう思いますか?」と尋ねるようなものではないことです。
Gemini 3 Pro が洗練された推論モデルであるため、ほとんどのユーザーが AI モデルで経験するものよりも優れているという事実は別として、6 どちらのシステムも、タスクを複数のステップのプロセスに分割し、より高い精度を確保し、幻覚を回避するように慎重に設計されています。
私の理解では、これは特に APD の場合に当てはまります。APD は AI 「エージェント」のチームを使用して機能し、各チームには完了する単一の個別のタスクが与えられます。
たとえば、個々のエージェントは、住宅のタイプを特定したり、申請がバンガローか半戸建て住宅のどちらであるかを判断したりするなど、提案に関する 1 つの基本的な事実を判断するタスクを任される可能性があります。
実際、システムによってエージェントが起動されて、これが非常に細かくなる可能性があることを理解しています。

[切り捨てられた]

## Original Extract

These two new systems could be genuinely revolutionary

Here's how the government is using AI to speed up the planning system
Subscribe Sign in Here's how the government is using AI to speed up the planning system
These two new systems could be genuinely revolutionary
James O'Malley Jun 17, 2026 13 2 1 Share POD! On YIMBY Pod this week, Martin and I look at Zack Polanski’s claims about the price of food, and I make myself popular by defending, er, the Big Supermarket. Then we speak to the excellent Thomas Ableman about how the Swiss would have built HS2 completely differently. Listen here, or wherever you get your pods!
Building stuff in Britain is a nightmare.
The arguments are well rehearsed by now. Our planning system is wildly bureaucratic with applications sometimes running to thousands of pages, and even the most thoughtful developments can be killed at the behest of a handful of grumpy councillors.
The government’s Planning and Infrastructure Act, which received Royal Assent at the end of last year, was definitely a leap in the right direction, and I do not want to play down its significance. It’s an important law, which reformed environmental mitigation, established a legal basis for creating local development corporations, and it shifted many smaller-scale planning applications from committee meetings to council planning officers’ discretion.
But fundamentally, these are upgrades to the existing system, and we haven’t moved to something bolder, like a system based on zoning, where buildings are effectively approved by default, as long as they meet a given zone’s criteria.
This means that if you want to build something in Britain, even something simple like a loft extension or a garden office, it still requires you to submit a tonne of paperwork to the council, and then to wait for a decision before you can get started.
This is very irritating, as it can involve long, uncertain waits. Officially, the statutory target is an eight-week turnaround time for decisions on small/minor household applications, but in practice it can take longer as councils only have a fixed number of planning officers available to scrutinise and approve the roughly 350,000 applications made every year.
And the impact of this sluggishness shouldn’t be underestimated. It’s essentially a tax on building, and a self-inflicted hit on the economy. Less gets built, productivity falls, and it ties up money earmarked for investment and financing.
However, all of this could be about to change, at least a little bit. Today the government has announced a pair of new projects to help planning officers make decisions faster. One is being built in-house by the Ministry of Housing, Communities and Local Government (MHCLG) and the Department of Science, Innovation and Technology (DSIT), 1 and the other is an £8.2m collaboration with Google and Faculty AI.
And as luck would have it, I’ve got some behind-the-scenes details on how the two systems are going to work.
The fundamental bottleneck on planning application processing is the time available to council planning officers. According to MHCLG , the average council planning department has around 40 people, 2 which might sound like a lot – but they have a lot to do, and a lot of their work is tedious grunt-work.
For example, many existing historic planning documents are not yet digitised, and are still stored on paper in filing cabinets. Which is bad if you’re a planning officer and need to consult the archives. So before you can even begin to consider an application, you need to digitise the existing documents.
Annoyingly, this is a process that could conceivably take hours , as the planning officer would have to scan the paper map and associated documents, and spend time manually transcribing information, and carefully drawing out a detailed digital map with a mouse.
But this is where the first tool, dubbed ‘Extract’ comes in.
A real planning department’s filing system, though the government didn’t name which one, presumably not wanting to publicly shame. Extract was built in-house by a team at MHCLG. It’s designed to automate the entire digitisation process, and turn paper documents and hand-drawn maps into digital objects that work with modern mapping and planning tools.
Queen’s Club Gardens, the “is Pepsi okay?” of London tennis venues. For example, above is a scan of the plans for Queen’s Club Gardens, in London. In 1981, it received an ‘Article 4’ designation, which limits what can be modified on the buildings 3 – a restriction that is relevant if a planning officer has to make a decision on an application in the area. Once it was fed into Extract, it was transformed into a modern digital shape file in just a couple of minutes.
And the way it does this is incredibly clever, as to do it reliably, there’s a multi-step pipeline, that involves multiple AI models working to turn a flat drawing into something useful.
So first, the PDF scan is run through Google’s Gemini model to extract textual information, such as dates and other details about the area. In principle, this might sound straightforward, as Optical Character Recognition (OCR) has existed for a long time – but the reality is many planning documents are messy. Some are handwritten, and many of those that are typed often have hand-written notes scribbled around the edges, or lines crossed out with a pen.
But whereas this would have tripped up older software, it isn’t much of a problem for a modern Large Language Model.
Anyway, after the text, maps are also identified by the AI. They’re then chopped out of the PDF, and fed into Meta’s Segment Anything , a specialist AI model that can take an image and identify different objects within. This is how the system extracts the shapes on the map – like the shape of the perimeter of the Article 4 area above, or the houses on the map below.
Segment analysis AI in action. It’s basically witchcraft. But what use is a shape if a system can’t plot it on a map? That’s why next Gemini is used to look at the paper maps and extract things like the names of streets and other geographic features. These are then fed into the Google Maps and Ordnance Survey APIs, to pin-point exactly where the map is supposed to be.
A pond in Hampstead. I’m not sure why the test locations were all such posh areas either. And even at this point, Extract’s job is not quite done. Because the shape file then needs to be placed on the digital map accurately. So here Extract uses another specialist AI model called MatchAnything , which has been trained to identify common points in, well, anything.
Apparently it’s capable of, for example, identifying common points in photos of the same object from different angles, and more relevantly here, it can figure out the common points on two maps of the same place where one map is digital, and the other is hand-drawn and upside-down, as in the above example from Hampstead. 4
And once you have these points, it becomes pretty straightforward to work out the longitude and latitude of each of the different points on the extracted shapes. Which means Extract can even super-impose the original scan on top of the digital map, skewing the image so that it matches the digital version.
And if it doesn’t perfectly match at the end, the planning officer can go in and edit the generated shapes manually, to ensure that it is accurate.
This looks pretty fun to use. This is all to say that there is a lot going on under the surface. But what’s amazing is that apparently the average processing time to turn a scan into a usable digital object is… is 1 minute, 42 seconds, turning a process that would have taken literally hours into something that is almost instantaneous. 5
So assuming Extract works as described, that’s going to be an enormous productivity boost in its own right. But this is only the start of the battle – as once all the relevant documents have been assembled, planning officers need to go through them – which is time consuming in its own right.
When a planning application is made, developers have to submit dozens, or even hundreds of pages outlining their plans, depending on the complexity. They have to explain what they want to build, and how it complies with whatever regulations apply to the local area.
Then there are often submissions from other stakeholders, such as statutory consultees like Natural England , or contributions from other residents who want to support or (more likely) oppose construction.
The insane number of documents required in planning applications. (Source: Lichfields .) These documents all land on the virtual desk of a planning officer, whose job it is to sift through and compare what the application says with what is contained in the planning system’s many rulebooks, which stretch to thousands of pages.
For example, planning decisions may have to take into account the National Planning Policy Framework, the council’s own Local Plan, or specific local rules around conservation areas that limit how buildings are allowed to look and feel, like the Article 4 area around Queen’s Club Gardens.
So, in effect, in order to make a judgement, the planning officer has to conduct a miniature literature review, comparing guidelines against proposals to work out if the development passes the extremely complicated test or not.
This, though, is where the second AI opportunity lies. This is at a much earlier stage than Extract, but Google and Faculty are currently working with the government to prototype a system dubbed Augmented Planning Decisions – or APD.
The idea here is to take what Large Language Models (in this case, Gemini 3 Pro) are already very good at – comparative text analysis – but apply this skill in a more structured and rigorous way to planning documents. So the plan is that APD will take the hundreds of pages in a planning application, and compare them against the thousands of pages of guidance that they need to be judged against, and present a summary containing everything the planning officer needs to know to make a final decision.
This is a mock-up of the UI, made by me, based on descriptions I’ve heard. But I thought it would help give you a flavour of how it might work. I’ve attempted to mock up something above that looks a bit like it. APD is basically a smart case management system. They can click into any application to see the AI’s reasoning, separated into all of the different criteria that need to be satisfied, each with a deep link to the specific guidance or legislation that justifies whatever conclusion has been drawn.
So all the officer then has to do is review the choices made, and decide whether or not they accept the AI’s recommendations. And even if they disagree with the AI conclusion, and change the decision before signing off on it, the system still saves a tonne of time compared to scrolling through thousands of pages of guidance manually.
At this point, depending on your level of AI-enthusiasm, you’re either pretty excited, or pretty sceptical.
I am very much in the excited camp.
Because what’s smart about the way both Extract and APD have been designed is that they are not just like if you or I were to slap a planning PDF into ChatGPT (or, I guess, Gemini) and say, “So what do you think of this then?”.
Aside from the fact that Gemini 3 Pro is a sophisticated reasoning model, and thus a cut above the experience most users have of AI models, 6 both systems have been carefully designed to break tasks down into multi-step processes, to ensure greater accuracy, and avoid hallucinations.
This is especially the case, as I understand it, for APD, which functions using teams of AI ‘agents’, each given a single discrete task to complete.
For example, an individual agent could be tasked with determining a single basic fact about the proposal like identifying the type of dwelling, to determine whether the application is for a bungalow or a semi-detached house, and so on.
In fact, I understand this can get really granular, with agents spun up by the system to figure out

[truncated]

---
source: "https://www.kentik.com/blog/when-local-blocks-go-global-the-india-telegram-bgp-incident/"
hn_url: "https://news.ycombinator.com/item?id=48601514"
title: "When Local Blocks Go Global: The India-Telegram BGP Incident"
article_title: "When Local Blocks Go Global: The India-Telegram BGP Incident | Kentik Blog"
author: "wmf"
captured_at: "2026-06-19T18:42:26Z"
capture_tool: "hn-digest"
hn_id: 48601514
score: 1
comments: 0
posted_at: "2026-06-19T18:21:55Z"
tags:
  - hacker-news
  - translated
---

# When Local Blocks Go Global: The India-Telegram BGP Incident

- HN: [48601514](https://news.ycombinator.com/item?id=48601514)
- Source: [www.kentik.com](https://www.kentik.com/blog/when-local-blocks-go-global-the-india-telegram-bgp-incident/)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T18:21:55Z

## Translation

タイトル: ローカル ブロックがグローバルになるとき: インドとテレグラムの BGP インシデント
記事のタイトル: ローカル ブロックがグローバルになるとき: インドとテレグラムの BGP インシデント |ケンティックのブログ
説明: インドでテレグラムをブロックすることを目的とした昨日の BGP ハイジャックの漏洩は、意図的であると同時に偶発的であると最もよく説明される最新のルーティング事故です。このパターンは、2008 年のパキスタンテレコムによる悪名高い YouTube ハイジャックにまで遡ります。このパターンでは、国内のブロックが封じ込めを逃れてサービスを中断しました。
[切り捨てられた]

記事本文:
ローカル ブロックがグローバルになるとき: インドとテレグラムの BGP インシデント | Kentik ブログ スライド 1/1 新しい AI Advisor 機能でトラブルシューティングと MTTR が高速化 • ブログを読む 新しい AI Advisor 機能でトラブルシューティングと MTTR が高速化 • ブログを読む 新しい AI Advisor 機能でトラブルシューティングと MTTR が高速化 • ブログを読む ブログに戻る ローカル ブロックがグローバルになるとき: インドとテレグラムの BGP インシデント
Doug Madory インターネット分析ディレクター 2026 年 6 月 17 日 インターネット分析 目次 背景 何が起こったのか?意図的だが偶然もある 結論 私たちのネットワークに参加してください
サインアップ このフォームに記入すると、プライバシー ポリシーと利用規約に同意したものとみなされます。
背景 何が起こったのでしょうか?意図的だが偶然もある まとめ
昨日のインドでの Telegram のブロックを目的とした BGP ハイジャックの漏洩は、意図的であると同時に偶発的であると最もよく説明される最新のルーティング事故です。このパターンは、2008 年のパキスタンテレコムによる悪名高い YouTube ハイジャックにまで遡ります。このパターンでは、国内のブロックが封じ込めを逃れて世界中でサービスを中断しました。
昨日、インドは健康診断詐欺への懸念を理由に、人気のメッセージングアプリ「テレグラム」を一時的にブロックする措置を講じた。このブロックを実装するために、インドの通信事業者 Rcom (AS18101) は、トラフィックを引き寄せてブラックホール化するために、BGP で Telegram の IP 空間を作成しました。
これはよく知られた話です。国境内にとどまることを目的としたハイジャックが漏洩し、世界中の一部のユーザーに対する Telegram サービスが中断されました。
この投稿では、このインシデントとその先行事例を調査し、RPKI ROV などのルート フィルタリング技術が漏洩による被害を制限するのにどのように役立ったかを検証します。
2026 年 6 月 16 日、インド政府はテレグラムに対する全国的なブロックを命令し、6 月 22 日まで有効となる予定でした。その命令は次のとおりでした。

この訴訟は、国家試験庁（NTA）の勧告に基づいてインドIT法第69A条に基づいて起こされたもので、同庁は今月下旬に予定されているニート（UG）の2026年医学部入学再試験の受験者をターゲットにした不正ネットワークによるテレグラムの組織的使用を引用した。
「紙漏れニート」や「プライベートマフィア」などの名前で活動する詐欺的な電報チャネルは、試験資料へのアクセスと称して学生とその家族に支払いを要求していました。政府はまた、この機能が紙流出の事後証拠を捏造するために悪用されたとして、テレグラムに対しメッセージ編集機能を6月30日まで無効にするよう命じた。
Telegram CEOのPavel Durov氏は公に反発し、1週間にわたる制限は責任のある悪役ではなく1億5,000万人以上のインドユーザーに対する罰だと主張し、またそれとは別に、インドの通信事業者Rcomが不正なインターネットルーティングアナウンスを通じてインド国外の一部ユーザーのTelegramへのアクセスを妨害していると主張した。最後の主張、つまりインドの大手通信事業者によるルーティング操作が、BGP データで観察されたことの舞台となるのです。
6 月 16 日の 07:17 UTC に、Rcom (AS18101、以前は Reliance Communications であり、Reliance Jio ではありません) は、表向きこれらのアドレス範囲宛てのトラフィックをブラックホールするために、Telegram で使用されるいくつかの IP ブロックの発信を開始しました。これらのハイジャックされたルートはインド国外に流出し、世界中のユーザーに影響を与えました。
以下の視覚化では、スタック プロットの上部にある赤いスライバーは、通常は Telegram (AS62041) によってのみ発生する AS18101 の 91.108.4.0/22 のハイジャックの伝播を表しています。ハイジャックされたルートを確認したのは、BGP ソースの 1.6% だけでした。これは、Telegram がすべてのルートに ROA を導入したためと考えられます。

ネットワークが RPKI が無効なルートをドロップできるようにします。
これに応じて、BGP 選択アルゴリズムでは小さいルートが優先されるため、Telegram は IP 空間の制御を取り戻すために、より具体的なハイジャック ルートの発信を開始しました。しかし、これが始まってから数時間後、AS18101 は 16:14 UTC にこれらのより具体的なものをハイジャックし始めました。
交通データが示したものは次のとおりです。ハイジャックが始まった瞬間から、Telegram が詳細を発表して反撃するまで、世界の Telegram トラフィックのごく一部がインドの AS18101 にリダイレクトされ、ドロップされました。この回復は長くは続きませんでした。AS18101 はより具体的なルートもハイジャックし、最終的にそれらのルートが廃止されるまでトラフィックを再びリダイレクトしました。
さまざまなルートがハイジャックされるタイミングの完全な表については、技術政策研究者プラネシュ・プラカシュ氏によるこの記事をご覧ください。
意図的だが偶然もある
昨日の Telegram をターゲットとした BGP ハイジャックの漏洩は、意図的であると同時に偶発的な BGP インシデントのカテゴリーの最新のものです。私はインターネット最大の BGP インシデントの歴史の中でそのいくつかを取り上げていますが、最も有名なのは 2008 年のパキスタンによる YouTube のハイジャックです。
しかし、ネットワーク事業者が政府の命令によるブロックを実装するために BGP を使用することを選択し、そのルートが誤ってグローバル ルーティング テーブルに漏洩してしまったため、ここ数年で他にも多数の事件が発生しました。
2021年2月のミャンマー軍事クーデター中、ソーシャルメディアサービスをブロックするという政府命令に従おうとして、あるISPがTwitterのルートを発信し始めた。その後、彼らはインターネット交換局でルートを漏洩し、その地域のユーザーへのサービスを妨害しました。このインシデントに対応して、Twitter はすべてのルートに対して ROA を作成し、RPKI ROV が次のようなルートの循環を拒否できるようにしました。

これらの事件のほとんどがそうであるように、出所が間違っている。
Twitter による RPKI ROV の採用は、ちょうど 1 年後、ロシアによるウクライナ侵攻の余波で独立系メディアに対する弾圧が行われている間に同じ出来事が起こったため、ちょうどいいタイミングで採用された。 Twitter が自社のルートに ROA を導入したことで、ロシアの情報漏えいによる巻き添え被害は大幅に限定されました。 Twitter の CISO が当時書いたように、「セキュリティを確保することの重要なポイントは、システムを常に破壊しないようにすることです」。そしてこのケースでは、BGP ハイジャックによってロシア国外のユーザーによる Twitter へのアクセスが遮断されるのを防ぎました。
つい最近、ブラジルは X/Twitter のブロックを命じ、またプロバイダー (AS263276) が BGP ハイジャックを国外に漏洩しました。その事件について聞いたことがありませんか？おそらく、流出したハイジャックルートがほぼ完全にフィルタリングされていたためと思われます。 RPKI ROV が正しく機能すると、ルーティングの失敗による中断を回避でき、これまでは知らなかった頭痛の種になります。
昨日の事件に最も近い事件は、2023年8月にテレグラムをブロックするイラクの動きだった。やはりハイジャックは漏洩したが、確実に知ることは不可能だが、ルートフィルタリング（おそらくRPKI ROV）により影響は最小限だった。
この特定のケースでは、これもまた「意図的だが偶然の」事件であったことは明らかであるようです。 AS18101 はハイジャックされたルートをインド国外に漏洩したくなかったと思われますが、実際に漏洩し、多くの国で Telegram サービスを妨害しました。 AS18101 が元のハイジャックに対抗するために Telegram が発表したより具体的なものをハイジャックしたという事実は、意図性をかなり明確にします。
AS18101 の唯一の中継プロバイダーである AS15412 が、インドのネットワーク AS9498 および AS4755 とともに RPKI が無効なルートをブロックしていたら、インシデントは次のようになっていただろう。

インドに封じ込められました。 RPKI ROV はこのシナリオでは非常にうまく機能しますが、これは国際通信会社が RPKI が無効なルートを拒否した場合に限られます。ありがたいことに、そのほとんどがそうなっていました。そうでなければ、Telegram ははるかに大きな混乱にさらされていたでしょう。
もちろん、この IP レベルのブロックを回避するために VPN を使用しなければならなかったインドのユーザーにとって、これはほとんど慰めにはなりません。検閲電報は現在、学生試験の不正行為に対抗するために政府が採用している最新の仕組みです。
対照的に、シリアは今年初め、高校生の不正行為を防止するために長年続いてきた国家閉鎖を中止すると発表した。インドもシリアと同様、人気のコミュニケーションツールやニュースや情報源を遮断することなく、この問題に対処する別の方法を見つけられるだろうか？
一時的なリークと自動化された BGP ルート リーク検出のブログ投稿 インターネット最大の BGP インシデントの簡単な歴史 ブログ投稿 イラクによる電報ブロック、ブラックホール BGP ルートのリーク ブログ投稿 Kentik は、最新のインフラストラクチャ チームのためのネットワーク インテリジェンス プラットフォームです。デモを入手する 844-356-3278 プラットフォーム
ネットワーク監視システム (NMS)
ネットワークのセキュリティとコンプライアンス
ネットワークパフォーマンスの監視
卓越したデジタル体験を提供
ネットワークポリシー管理を強化する
セキュリティインシデントの調査
すべてのクラウドとネットワークのトラフィックを視覚化
インターネットのパフォーマンスを理解する
最新のネットワーキング テレメトリの実践ガイド
Kentik AI Advisor のご紹介: ネットワーク インテリジェンスの未来
Seaborn Networks は Kentik で年間 400 万ドルを節約し、収益性を向上
Kentik ネットワーク分析センター

## Original Extract

Yesterday’s leak of a BGP hijack intended to block Telegram in India is the latest routing mishap best described as intentional, but also accidental — a pattern dating back to Pakistan Telecom’s infamous hijack of YouTube in 2008, in which a domestic block escaped containment and disrupted the servi
[truncated]

When Local Blocks Go Global: The India-Telegram BGP Incident | Kentik Blog Slide 1 of 1 New AI Advisor features accelerate troubleshooting and MTTR • Read the blog New AI Advisor features accelerate troubleshooting and MTTR • Read the blog New AI Advisor features accelerate troubleshooting and MTTR • Read the blog Back to Blog When Local Blocks Go Global: The India-Telegram BGP Incident
Doug Madory Director of Internet Analysis June 17, 2026 Internet Analysis Table of contents Background What happened? Intentional, but also accidental Conclusion Join our network
Sign up By completing this form, you agree to our Privacy Policy and Terms of Use .
Background What happened? Intentional, but also accidental Conclusion Subscribe Summary
Yesterday’s leak of a BGP hijack intended to block Telegram in India is the latest routing mishap best described as intentional, but also accidental — a pattern dating back to Pakistan Telecom’s infamous hijack of YouTube in 2008, in which a domestic block escaped containment and disrupted the service worldwide.
Yesterday, India moved to temporarily block the popular messaging app Telegram over concerns about medical exam fraud. To implement this block, Indian carrier Rcom (AS18101) originated Telegram’s IP space in BGP in order to attract and blackhole the traffic.
This is a familiar story: a hijack intended to stay within national borders leaked out, disrupting Telegram service for a portion of users around the world.
In this post, I’ll explore the incident, its predecessors, and examine how route filtering techniques like RPKI ROV helped to limit the damage caused by the leak.
On June 16, 2026, India’s government ordered a nationwide block on Telegram , set to remain in effect through June 22. The order was issued under Section 69A of India’s IT Act on the recommendation of the National Testing Agency (NTA), which cited the organized use of Telegram by cheating networks targeting candidates sitting for the NEET (UG) 2026 medical school entrance re-examination scheduled later this month.
Fraudulent Telegram channels — operating under names like “PAPER LEAKED NEET” and “Private Mafia” — were demanding payments from students and their families in exchange for purported access to exam materials. The government also ordered Telegram to disable its message-editing feature through June 30, arguing the feature had been exploited to fabricate after-the-fact evidence of paper leaks.
Telegram CEO Pavel Durov pushed back publicly , calling the week-long restriction a punishment of more than 150 million Indian users rather than the bad actors responsible — and separately alleged that Indian telecom operator Rcom was disrupting access to Telegram for some users outside India through unauthorized internet routing announcements. It is that last claim — routing manipulation by a major Indian carrier — that sets the stage for what we observed in the BGP data.
At 07:17 UTC on June 16th, Rcom (AS18101, formerly Reliance Communications and not Reliance Jio ) began originating several IP blocks used by Telegram, ostensibly to blackhole traffic destined for those address ranges. Those hijacked routes leaked outside of India and impacted users around the world.
In the visualization below, the red sliver along the top of the stack plot represents the propagation of AS18101’s hijack of 91.108.4.0/22, normally only originated by Telegram (AS62041). Only 1.6% of our BGP sources saw the hijacked route, a number likely owing to Telegram’s deployment of ROAs on all of its routes, enabling networks to drop the RPKI-invalid routes.
In response, Telegram began originating more-specific routes of the hijacks to regain control of the IP space, as the smaller routes get priority in the BGP selection algorithm. However, hours after this began, AS18101 began hijacking those more-specifics at 16:14 UTC.
Here’s what the traffic data showed. From the moment the hijack began until Telegram fought back by announcing more-specifics, a small portion of global Telegram traffic was redirected to AS18101 in India and dropped. That recovery was short-lived: AS18101 hijacked the more-specifics too, redirecting traffic once again until those routes were finally withdrawn.
For a full table of the timings of the various hijacked routes, check out this write-up by tech policy researcher Pranesh Prakash .
Intentional, but also accidental
Yesterday’s leak of BGP hijacks targeting Telegram is the latest in a category of BGP incidents which are both intentional , but also accidental . I cover several of them in my Brief History of the Internet’s Biggest BGP Incidents , with the most famous being Pakistan’s hijack of YouTube in 2008.
But there have been numerous other incidents over the years as network operators have opted to use BGP to implement a block ordered by the government, only to mistakenly leak the routes into the global routing table.
During the military coup in Myanmar in February 2021, one ISP began originating Twitter’s routes in an attempt to comply with a government order to block the social media service. They subsequently leaked the routes at an internet exchange, disrupting the service for users around the region. In response to this incident, Twitter created ROAs for all of its routes, enabling RPKI ROV to reject the circulation of routes with incorrect origins, as is the case in most of these incidents.
Twitter’s adoption of RPKI ROV came just in time, as an identical event occurred just a year later during Russia’s crackdown on independent media in the aftermath of their invasion of Ukraine. Twitter’s rollout of ROAs for their routes greatly limited the collateral damage caused by the Russian leak. As Twitter’s CISO wrote at the time , “a bunch of the point of having security is to keep your systems from breaking all of the time,” and in this case, it kept a BGP hijack from breaking access to Twitter for users outside of Russia.
Most recently, Brazil ordered X/Twitter to be blocked , and again, a provider (AS263276) leaked its BGP hijack outside the country. Never heard of that incident? That’s probably because the leaked hijack routes were almost completely filtered. When RPKI ROV works correctly, disruptions due to routing mishaps can be avoided, becoming a headache that you never knew that you didn’t have.
The case closest to yesterday’s incident was Iraq’s move to block Telegram in August of 2023. Again, the hijack was leaked, but the impact was minimal due to route filtering, likely RPKI ROV, although it is impossible to know for sure.
In this particular case, it seems quite clear that this was another “intentional, but also accidental” incident. AS18101 likely didn’t want to leak their hijacked routes outside of India, but they did , disrupting Telegram service in numerous countries. The fact that AS18101 hijacked the more-specifics that Telegram announced to combat the original hijacks makes the intentionality pretty clear.
Had AS15412, the sole transit provider for AS18101, along with Indian networks AS9498 and AS4755 blocked RPKI-invalid routes, the incident would have been contained to India. RPKI ROV works pretty well in this scenario, but only if international carriers reject RPKI-invalid routes. Thankfully, most of them did, or Telegram would have been disrupted to a far greater extent.
Of course, this is little solace to the Indian users who had to employ a VPN to circumvent this IP-level blockage. Censoring Telegram is now the latest mechanism employed by a government to combat cheating on student exams.
In contrast, earlier this year, Syria announced that it would stop its long-running practice of national shutdowns to prevent cheating by high schoolers. Might India, like Syria, find another way to address this issue without blocking a popular communication tool and source of news and information?
Ephemeral Leaks and Automated BGP Route Leak Detection Blog Post A Brief History of the Internet’s Biggest BGP Incidents Blog Post Iraq Blocks Telegram, Leaks Blackhole BGP Routes Blog Post Kentik is the network intelligence platform for modern infrastructure teams. Get a Demo 844-356-3278 Platform
Network Monitoring System (NMS)
Network Security and Compliance
Network Performance Monitoring
Deliver Exceptional Digital Experiences
Harden Network Policy Management
Investigate Security Incidents
Visualize All Cloud and Network Traffic
Understand Internet Performance
Practical Guide to Modern Networking Telemetry
Introducing Kentik AI Advisor: The Future of Network Intelligence
Seaborn Networks Saves $4 Million Annually and Increases Profitability with Kentik
Kentik Network Analysis Center

---
source: "https://www.theregister.com/ai-and-ml/2026/08/05/time-magazine-has-a-separate-version-of-its-website-with-ads-only-ai-can-see/5283640"
hn_url: "https://news.ycombinator.com/item?id=49194660"
title: "Time Magazine has a separate version of its website with ads only AI can see"
article_title: "Time Magazine has a separate version of its website with ads only AI can see"
author: "rguiscard"
captured_at: "2026-08-06T10:28:49Z"
capture_tool: "hn-digest"
hn_id: 49194660
score: 2
comments: 0
posted_at: "2026-08-06T09:58:36Z"
tags:
  - hacker-news
  - translated
---

# Time Magazine has a separate version of its website with ads only AI can see

- HN: [49194660](https://news.ycombinator.com/item?id=49194660)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/08/05/time-magazine-has-a-separate-version-of-its-website-with-ads-only-ai-can-see/5283640)
- Score: 2
- Comments: 0
- Posted: 2026-08-06T09:58:36Z

## Translation

タイトル: Time Magazine には AI のみが閲覧できる広告を含む別バージョンの Web サイトがある
説明: ブランドは、チャットボットの発言に影響を与えるためにすでにお金を払っています。

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
Time Magazine には、AI のみが閲覧できる広告を掲載した別バージョンの Web サイトがあります
ブランドはすでに、チャットボットによる自社についての発言に影響を与えるためにお金を払っています。
AI ウェブページ クローラーには、通常の訪問者が決して目にすることのない独自の広告が配信されるようになりました。ある企業の上司は、チャットボットがブランドについて語ることに影響を与える戦略を公然と説明しています。次回クロードに銀行口座について尋ねるとき、その答えはこのボットをターゲットとしたコンテンツの影響を受けている可能性があります。
これまでに文書化された例は 1 つだけです。それは、ドイツを拠点とするフリーランスのソフトウェア開発者 Vincent Schmalbach が発見した、選択された AI クローラーに AI のみの広告を配信する時間です。シュマルバッハ氏は水曜日のブログ投稿で、AIクローラーには提供されるが通常のブラウザーには提供されない一部のTimeページのマークダウンバージョンに埋め込まれたスポンサー付きコンテンツをどのように発見したかを詳細に説明した。
これらのページには、オンライン専用銀行 Ally の広範な FAQ に先立って、アドテク ベンダー Mobian の広告タグも含まれていました。 FAQには、アリーが「現在、生涯にわたって設立された唯一の銀行」であるという主張とともに、この無店舗銀行に関連する手数料無料のATMの数などのブランドの「事実」が含まれており、アリーは「1つのカテゴリー」に分類されている。
AI のみの広告には、「Ally は日常の銀行業務に適していますか?」「Ally Bank に現金を預けることができますか?」などの質問への回答と同様に、反証可能な重要な事実となる主要な「ブランド事実」の記述も AI のみの広告に含まれています。 FAQ はマークダウン ページでスポンサー付きコンテンツとしてフラグが立てられていますが、内容全体が吐き出すように設計されているかのように書かれているという事実は変わりません。

特定の質問に応答してチャットボットによって送信されます。
すべての Web クローラーがこれらのマークダウン広告を取得しているわけではありません。 Schmalbach 氏によると、Google の標準の検索インデックス作成ボットは、人間が取得するのと同じ HTML を取得するだけなので、広告は表示されません。
「[チャットボット] アシスタント クローラーのみがフォークされたバージョンを取得します」と彼は書いています。つまり、ClaudeBot、OAI-SearchBot、PerplexityBot などのユーザー エージェントはスポンサー付きマークダウンを受け取りますが、Google の標準検索クローラーは通常のブラウザに提供されるのと同じ HTML を受け取ります。 OpenAI がそれぞれモデル トレーニングとユーザー開始の取得に使用する GPTBot と ChatGPT-User は、シュマルバッハのテストで HTTP 406 エラーを返しますが、OAI-SearchBot は ChatGPT Search をサポートするために使用されるマークダウン バージョンを受け取ります。
AI 広告が組み込まれたこれらのストーリーのマークダウン バージョンは、HTTP リクエストの一部として送信される User-Agent ヘッダーを変更する人間によって表示される必要があり、シュマルバッハ氏はそのようにしてそれらを発見したと述べています。 The Register は、Time のいくつかの記事を自分たちの目で確認するためにテストし、検索エンジン最適化会社 Encited が提供するクローラー シミュレーターを使用して結果を再現することができました。
クロードボットとして見ると、タイム誌の 2025 年のベスト発明リストにはアリー銀行の広告が含まれ、タイム誌の 2026 年のトップ コンテンツ クリエイターに関する記事や、このリストを記念して最近ニューヨークで開催されたクリエイターズ パーティーの写真ギャラリーも含まれます。私たちがチェックしたニュース記事には値下げ広告は表示されませんでした。これは、広告が賞味期限の短い記事ではなく、より常緑のコンテンツの一部である可能性があることを示唆しています。
シュマルバッハ氏が記事の中で指摘しているように、これは「主な視聴者が AI モデルである場合にウェブがどのようなものになり始めるかを初めて明確に示すもの」になる可能性があり、これは私たちがすでに超えているしきい値です。
AI に広告をやらせることができるのに、なぜ人間に宣伝する必要があるのか

あなたは？
AI クローラーに広告を掲載するという Time の戦略が他のパブリッシャーによって採用されているかどうかは不明ですが、少なくとも他のパブリッシャーがそれを検討していなかったとしたらショックです。
上で述べたように、AI クローラーのトラフィックはすでに Web ページへの人間のアクセスを上回り始めており、広告主にとっては、クローラーを作成した人間よりも、これらのクローラーをターゲットにしたほうがより儲かる可能性があります。調査対象となった米国の消費者の半数はウェブ検索にAIを利用していると回答しており、AIの応答に影響を与えることで消費者にリーチできる可能性は信じられないほど強力であり、タイム社とその広告パートナーであるモビアン社はそこに期待しているようだ。
私たちがシュマルバッハ氏の報告書と私たちが見つけた結果を検証するよう同出版局に依頼したところ、彼らは質問には答えず、代わりに、私たちが目にしたものだけでなく、この計画に対する私たちの懸念も裏付けた先週のメディアおよびマーケティング媒体DIGIDAYの記事を私たちに指摘しました。
チャットボットは人々を操作して物を買わせるのが得意だとプリンストン大学の関係者らは発見
新しい品質シグナルと収益モデルがなければ、AI 検索はウェブを破壊する可能性がある
AI エージェントはあなたの美しいウェブサイトや魅力的な広告には興味がありません
Cloudflare、広告付きウェブページからの皮肉な検索およびスクレイピングボットをブロック
DIGIDAY によると、Time は 6 月に AI クローラーにとってより魅力的なものにするために Web ページのマークダウンへの変換を開始し、その見直しには Mobian との提携が含まれていました。 Timeの最高執行責任者（COO）マーク・ハワード氏はDIGIDAYに対し、同社の営業チームがウェブサイトを値下げに切り替えたブランドにエージェント広告を売り込んでおり、これらのブランドがAIボットの利用を検討していることを示唆していると語った。
ただし、これは Mobian の CEO 兼共同創設者である Jonah Goodhart の言葉であり、その目的を明確にしています。同社の目的は、サービスに料金を支払っているブランドについて AI アシスタントが語る内容を形作ることです。
「わ

ChatGPT に影響を与えると、潜在的に ChatGPT 全体に影響を与えることになります。 ChatGPTが（ブランドについての）発言を変えるとしたら、それは大規模であり、どのキャンペーンでもできることを超えています」とグッドハート氏はDIGIDAYに語った。
これまでのところ、両氏はDIGIDAYに対し、AI検索結果に影響を与えようとした最初のブランドはアリー銀行とプロジェクト管理研究所であることを認めており、シュマルバッハ氏とザ・レジスター氏がタイム誌のAIクローラー版のページで見つけた広告と一致している。
タイムとモービアンはDIGIDAYに対し、この提携はパブリッシャーがAIクローラーを直接ターゲットにした広告を配信するのは初めてだが、「すでにさらに多くの広告がラインナップされている」と語った。
言い換えれば、AI を利用した検索でどの結果が優れた製品であるために正当に表示されているのか、そしてどれがシステムを騙しているだけなのか、すぐに判断できなくなる可能性があります。 SEO 競争の初期の頃が再び繰り返されていますが、今回だけインデクサーを騙すのは、内容が真実になるまで同じパンフレットを何度も何度も老馬鹿に渡すのと同じくらい簡単です。 ®
aiとml
メタが最新のAIエージェントがテストペンから出てきたことを世界に発表
別の週、別の企業が自社のモデルの1つが想定外の場所に到達した理由を説明した
News Corp、一部のAI企業を「ひどい窃盗犯」とレッテルを貼る
OpenAI や Meta などのコンテンツが「原則的」で「名誉ある」ものになるために料金を支払います。
HPE と NVIDIA で AI を大規模に運用可能
スポンサー投稿: AI ファクトリーのコンセプトには新しい AI スタックが含まれる、HPE の Thierry Pienaar 氏と NVIDIA の Kaushik Shirhatti 氏が説明
Azure CTO が Doom を一度に 1 フレームずつペイントに貼り付ける
何も計算せず、すべてをレンダリングし、サウンドさえあります
コラムニスト
ネットワーク ハードウェアに API を追加しても管理上の課題は解決されません
これがクラウドにインスピレーションを得たネットワークの理由です

win: 仮想スイッチはすべて一貫性があります
HPE、見積ハードウェア価格の有効期限をおそらく数か月間延長
ビッググリーンは部品価格が今後も安定することを知っていると示唆
OSプラットフォーム
IT 上司は子供を職場に連れて行く日のために root セッションを開いたままにした
aiとml
AI バブルはすでにはじけつつあります。私たちはまだそれを知らないだけです
OSプラットフォーム
Microsoftは、Windows 11を実行している人には8 GBのRAMで十分であると述べています
セキュリティ
ロンドン警察は被害者の新しい住所と電話番号をストーカーに渡した、と監視当局が発表
サース
中古ソフトウェアライセンスの請求が集中し、Microsoft にとって二重の問題
LLM でロシアン ルーレットをプレイしますか、それとも専門家の手に健康と安全を委ねますか
宇宙では、LLMが幻覚を起こしても、あなたの叫び声は誰にも聞こえません
ほぼ 10 トーク/秒で、ほとんど一貫性があります — 気に入らない点は何ですか?
BSides、Black Hat、DEF CON がラスベガスに上陸する際に期待されること
Alibaba の Qwen チームは初めて、その「Max」モデルを API ペンから解放します。一方、DeepSeek V4-Flash は、安くて陽気な競争に新たな意味を与えます
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
FOSS は Microsoft の独占を 1 つ打ち破りました。 20年間の失敗を経て、次の失敗をする時が来た
一言
GNOME は Windows のように見えることができ、Flashback は拡張機能なしで実行できます
新しい「シンプルタスクバー」はオプションですが、よりシンプルで安定した方法があります
黙祷をお願いします

、x86-32 での Debian の最終リリース用
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

Brands are already paying to influence what chatbots say about them

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Time Magazine has a separate version of its website with ads only AI can see
Brands are already paying to influence what chatbots say about them
AI webpage crawlers are now being served their very own ads that ordinary visitors never see, with one firm's boss openly describing a strategy to influence what chatbots say about brands. The next time you ask Claude about where to bank, its answer may have been influenced by this bot-targeted content.
We only have one documented example so far, and that's Time serving AI-only ads to selected AI crawlers, as spotted by Germany-based freelance software developer Vincent Schmalbach. In a Wednesday blog post , Schmalbach detailed how he found sponsored content embedded in markdown versions of some Time pages that are served to AI crawlers but not ordinary browsers.
Those pages also contained advertising tags from adtech vendor Mobian ahead of extensive FAQs for online-only bank Ally. The FAQs include brand "facts," such as the number of fee-free ATMs associated with the branchless bank, alongside claims that Ally is "the only bank built for life today," putting it in "a category of one."
Key “brand fact” statements that make for great regurgitable facts are also included in the AI-only advertisement, as are answers to questions like “is Ally good for everyday banking,” “can you deposit cash at Ally Bank,” and the like. The FAQ is flagged as sponsored content on the markdown page, but it doesn’t change the fact that the entire thing is written like it was designed to be spit back out by a chatbot in response to specific questions.
Not all web crawlers are getting these markdown ads either. Per Schmalbach, Google’s standard search indexing bot doesn't see the ads, as it just gets the same HTML that a human gets.
“Only the [chatbot] assistant crawlers get the forked version,” he wrote. That means user-agents including ClaudeBot, OAI-SearchBot, and PerplexityBot receive the sponsored markdown, while Google's standard search crawler gets the same HTML served to ordinary browsers. GPTBot and ChatGPT-User, which OpenAI uses for model training and user-initiated retrieval respectively, return HTTP 406 errors in Schmalbach's tests, while OAI-SearchBot receives the markdown version used to support ChatGPT Search.
These AI ad-infused markdown versions of stories should be viewable by humans who change the User-Agent header sent as part of HTTP requests, which is how Schmalbach said he spotted them. The Register tested several articles from Time to see for ourselves, and we were able to replicate the results using the crawler simulator provided by search engine optimization firm Encited.
When viewed as ClaudeBot, Time’s Best Inventions of 2025 list includes the Ally Bank ad, as do stories about Time’s top content creators of 2026, as does a gallery of photos from a creators' party held in New York recently to honor the list. The markdown ads don’t appear on newsier stories that we checked, suggesting the ads may be part of more evergreen content rather than articles with a shorter shelf life.
As Schmalbach points out in his writeup, this could be “the first clear look at what the web starts to become when the main audience is AI models,” which is a threshold we’ve already crossed .
Why advertise to humans when you can make AI do it for you?
It’s not clear whether Time’s strategy of advertising to AI crawlers has been adopted by other publishers, but it’d be a shock if others weren’t at least considering it.
As noted above, AI crawler traffic has already begun to surpass human visits to webpages, making it potentially more lucrative for advertisers to target those crawlers than the humans who made them. With half of the US consumers surveyed saying they use AI to search the web, the potential to reach consumers by influencing AI responses is incredibly powerful, and that is what Time and its advertising partner Mobian appear to be banking on.
When we asked the publication to verify Schmalbach’s report and the results we found, they didn’t answer questions, instead pointing us to a story in media and marketing outlet Digiday from last week that verified not only what we saw but our fears about the scheme, too.
Chatbots are great at manipulating people to buy stuff, Princeton boffins find
AI search could kill the web without new quality signals and revenue models
AI agents don't care about your pretty website or tempting ads
Cloudflare to block cynical search-and-scrape bots from ad-supported web pages
Per Digiday, Time started converting its web pages to markdown in order to make them more appealing to AI crawlers in June, and a partnership with Mobian was included in that rework. Time COO Mark Howard told Digiday that the publication’s sales team is pitching agent ads to brands that have also converted their websites to markdown, as it suggests those brands are thinking about reaching AI bots.
It’s a quote from Mobian CEO and co-founder Jonah Goodhart that makes the objective clear, however: The company's aim is to shape what AI assistants say about the brands paying for the service.
“When you influence ChatGPT, you’re influencing potentially all of ChatGPT. If ChatGPT changes what it says about [a brand], it’s massive and it’s more than any one campaign could ever do,” Goodhart told Digiday. “With a human you influence one person.”
So far, the pair confirmed to Digiday that Ally Bank and the Project Management Institute were the first brands trying to influence AI search results, matching the ads Schmalbach and The Register found in AI crawler versions of Time's pages.
Time and Mobian told Digiday that the partnership marks the first time a publisher has served ads directly targeting AI crawlers, but that “more are already being lined up.”
In other words, you soon may not be able to tell which results in an AI powered search are legitimately being surfaced because of being good products, and which are just gaming the system. It’s the early days of the SEO race all over again, only this time fooling the indexers is as easy as handing any old idiot the same pamphlet over and over again until its content becomes truth. ®
ai and ml
Meta latest to tell world its AI agent wandered out of test pen
Another week, another firm explaining why one of its models reached somewhere it wasn't supposed to
News Corp labels some AI companies 'crass kleptomaniacs'
Pay it for content, like OpenAI and Meta, to become 'principled' and 'honorable'
Operationalize AI at scale with HPE and NVIDIA
SPONSORED POST: The AI factory concept encompasses the new AI stack, explain HPE’s Thierry Pienaar and NVIDIA’s Kaushik Shirhatti
Azure CTO pastes Doom into Paint one frame at a time
It computes nothing, renders everything, and even has sound
COLUMNISTS
Adding an API to networking hardware doesn’t solve management challenges
This is why cloud-inspired networks win: virtual switches are all consistent
HPE extends validity of quoted hardware prices, possibly for months
Suggests Big Green knows its component prices will remain steady
OS PLATFORMS
IT boss left root session open for bring-your-kid-to-work day
ai and ml
The AI bubble is already popping; we just don't know it yet
OS PLATFORMS
Microsoft says 8 GB of RAM should be enough for anyone running Windows 11
Security
London cops handed victim's new address and number to her stalker, watchdog says
saas
Double trouble for Microsoft as pre-owned software license claims converge
Would you rather play Russian roulette with an LLM or put your health and safety in the hands of a professional
In space, no one can hear you scream when the LLMs hallucinate
Nearly 10 tok/s and it's mostly coherent — What's not to like?
What to expect as BSides, Black Hat, and DEF CON descend on Las Vegas
For the first time, Alibaba's Qwen team is letting its 'Max' model out of the API pen; meanwhile, DeepSeek V4-Flash gives new meaning to cheap and cheerful competition
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

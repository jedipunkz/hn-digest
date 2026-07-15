---
source: "https://ahrefs.com/blog/llmstxt-study/"
hn_url: "https://news.ycombinator.com/item?id=48923287"
title: "We Analyzed 137K Sites: 97% of Llms.txt Files Never Get Read"
article_title: "We Analyzed 137K Sites: 97% of llms.txt Files Never Get Read"
author: "mooreds"
captured_at: "2026-07-15T17:06:52Z"
capture_tool: "hn-digest"
hn_id: 48923287
score: 2
comments: 0
posted_at: "2026-07-15T16:26:55Z"
tags:
  - hacker-news
  - translated
---

# We Analyzed 137K Sites: 97% of Llms.txt Files Never Get Read

- HN: [48923287](https://news.ycombinator.com/item?id=48923287)
- Source: [ahrefs.com](https://ahrefs.com/blog/llmstxt-study/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T16:26:55Z

## Translation

タイトル: 137,000 のサイトを分析しました: Llms.txt ファイルの 97% は読み取られませんでした
記事のタイトル: 137,000 のサイトを分析: llms.txt ファイルの 97% は読まれない
説明: 137,000 のサイトを分析したところ、llms.txt ファイルの 97% が読み取られていないことがわかりました。ボットがこれらのファイルを無視する理由と、これが AI の可視性にとって何を意味するかを確認してください

記事本文:
137,000 のサイトを分析しました: llms.txt ファイルの 97% は決して読み取られませんでした
137,000 のサイトを分析しました: llms.txt ファイルの 97% は読み取られませんでした
この投稿にリンクしている Web サイトの数。
この投稿の推定月間オーガニック検索トラフィック。
Ahrefs Web Analytics と Bot Analytics を使用して、137,000 ドメインのサーバー ログとライブ トラフィック、さらにそれらすべてにアクセスするユーザー エージェントを分析しました。
Ahrefs Web Analytics を使用している 137,000 ドメインのうち 28% が llms.txt ファイルを公開しています。
これらのファイルの 97% は、2026 年 5 月にはトラフィックがゼロでした。それらはまったく取得されませんでした。
llms.txt ファイルに到達したリクエストの 96% はボットからのものでした。
フェッチの 19.5% は名前付き AI ツールからのものでした (無視されなかったファイルの 3% のうち)。 GPTBot がトップ、Claude-Code が 2 位で、あらゆる AI 検索およびアシスタント ボットを上回っています。
取得の 12% は、GEO/AEO ツール、llms.txt チェッカー ツール、研究者など、自らを研究している業界からのものです。
存在しない llms.txt ファイルに対する AI ボットからのリクエストはゼロでした。彼らは決して探しに行きません。
Chrome Lighthouse llms.txt 監査では、およそ 1,000 件に 1 件のフェッチが発生しました。
2026 年 5 月下旬、Google は 1 週間足らずで llms.txt の議論の両側を受け入れました。
生成 AI 機能の最適化に関する新しいガイドでは、文字通り「神話を打ち破る」というタイトルのセクションで、生成 AI 検索に表示されるために llms.txt のような機械可読ファイルは必要ないとサイト所有者に伝えています。
数日後、Chrome チームは Lighthouse の実験的なエージェント ブラウジング監査内の llms.txt チェックを出荷しました。そのファイルがないと、エージェントがサイトの構造を理解するためにサイトのクロールに多くの時間を費やす可能性があることを説明するドキュメントが含まれていました。
リリー・レイがグーグルのジョン・ミューラーにその矛盾について詰め寄ったとき、彼はllms.txtは「検索のために行われたものではない」と説明した。それは「おそらく救うた​​めの一時的な松葉杖」です

開発者ドキュメントを解析する AI コーディング ツール用の「トークン」ですが、開発者以外のサイトが心配する必要はありません。
同氏はまた、サイト所有者がログをチェックしても、AIエージェントのトラフィックが非常に少ないことが分かるだろうとも述べた。
これは私たちがテストすることにしたものです。
先に進む前に、llms.txt が実際には何なのかを理解しましょう。 Llms.txt はマークダウンで書かれた単一のインデックス ファイルで、サイトのルートに配置されます。 Answer.AI と fast.ai の共同創設者である Jeremy Howard によって 2024 年に提案されたもので、サイトとは何かを要約し、その最も重要なコンテンツにリンクしています。 LLM とエージェントは、すべてをクロールすることなく、この情報を使用して方向を定めることができるという考えです。 llms.txt の「AI 可視化」フレームワークは、AI プラットフォームがファイルに報酬を与えるという憶測に基づいて採用が広がるにつれて、SEO 業界によって後から追加されました。よく混同されがちな点と、混同されない点が 2 つあります。
これは、Web ページのマークダウン コピーを公開するという行為ではなく、それ自体に問題がある別の戦術です。
ファイル名にもかかわらず、これは robots.txt スタイルのディレクティブではありません。何も制御せず、何もブロックしません。
この調査では、インデックス ファイルとインデックス ファイルのみを測定します。
私たちの調査は、2026 年 5 月にトラフィックを受信した Ahrefs Web Analytics の 137,210 個のドメインすべてに焦点を当てています。
各ドメイン ルートで HTTP 200 を返す llms.txt を確認し、次に Ahrefs Bot Analytics を使用して母集団全体の /llms.txt パスへのすべてのリクエストを調べ、HTTP 応答 (200 対 404) ごとに分割し、チャネルおよび個々のユーザー エージェントごとに分類しました。
ソフト 404 やファントム ファイルを除外するために、各ファイルが HTML ではなく実際の Markdown であることも確認し、タイトルとコンテンツをスクリーニングして「404」や「ページが見つかりません」などのエラー信号を検出しました。
Ahrefs Web Analytics の顧客は、Web 全体よりも技術的および SEO 意識が偏っているため、

28% の採用率を上限として考えてください。
ファイルが llms.txt 仕様に照らして整形式であるかどうかについては、明示的に調査しませんでした。
ドメインの 28% が llms.txt を公開
Google 検索のガイダンスではスキップできるとされており、Chrome チームはそれを監査しており、Mueller 氏はこれをコーディング ツールの一時しのぎと呼んでいます。
では、さまざまなメッセージが混在する中で、llms.txt は実際にどの程度普及しているのでしょうか?私たちの調査の対象となった 137,000 のドメインのうち、28% がこれらのファイルを公開しています。
主要な AI プラットフォームが llms.txt の読み取りに取り組んでいないにもかかわらず、人口の 4 分の 1 以上のドメイン (38,000) が llms.txt を採用しています。
AI プラットフォームがファイルを消費し始めるという確証ではなく、AI プラットフォームがファイルを消費し始めるのではないかという推測によって導入が推進されてきました。
llms.txt ファイルの 97% はリクエストを受け取りません
私たちの研究では、ほとんどすべての llms.txt ファイルが未読です。
有効なファイルを持つ約 38,000 のドメインのうち、97% では 5 月にはそのファイルに対するリクエストがまったくありませんでした。
残りの 3% (1.1,000 個のドメイン) は、測定したすべての llms.txt トラフィックを受信しました。
私たちのデータは、ジョン・ミュラーが正しいことを示唆しています。このファイルの結果、AI トラフィックがほとんど見られないだけでなく、トラフィックも非常に少ないことがわかります。
今日 llms.txt ファイルを公開した場合、最も可能性の高い結果は、何も取得されないことです。
ただし、読み取られるファイルの 3% は、興味のある訪問者によって読み取られます。
残りの研究ではそれらに焦点を当てていきます。
llms.txt ファイルへのリクエストの 96% はボットからのものです
Llms.txt ファイルはマシン用に書かれており、ファイルを読み取るのはほぼマシンだけです。
トラフィックを受信したファイル全体では、リクエストの 96% がボットからのものでした。
人間が占める割合は 4% で、そのうちのかなりの部分は SEO のようであり、チャット アプリで llms.txt リンクを共有しており、展開ボットが律儀にリンクを取得しています。
Slackbot 単独では、PerplexityBo よりも頻繁に llms.txt ファイルを取得しました。

しなかった。
Perplexity は、llms.txt が役立つように設計された AI 検索エンジンの 1 つであるため、チャット アプリのリンク プレビュー ボットがそれを上回っていることが判明すると、これらのファイルが実際に AI 検索の関心をどれだけ生み出しているかが雄弁にわかります。
llms.txt を読み取るボットの 77% は AI ツールからのものではありません
多くのサイトが llms.txt を公開しているのは、それが ChatGPT の回答に表示されたり、Perplexity で引用されたり、AI 概要を獲得したりする可能性が高まると考えているからです。
しかし、私たちのデータは別のことを物語っています。llms.txt を取得するボットの 77% はまったく AI ツールではありません。
どのボットが llms.txt を要求しているかを理解するために、すべてのユーザー エージェントを 12 のカテゴリに分類しました。
* SEO 監査ツールには、Ahrefs 独自のクローラー (SiteAuditBot、Ahrefs Bot、および Ahrefs Site Audit) が含まれており、これらを合わせると 2,334 件のリクエスト (全体の 10.6%) を占めます。これらを除くと、サードパーティの SEO 監査ツールが 2,442 リクエスト (11.1%) を占めます。ボット カテゴリは合計リクエストの 96% に相当します。残りの 4% (930 件のリクエスト) は人間からのものでした。
個別に見ると、トップ 4 に入る AI ボット カテゴリはありません。
SEO 監査ツール (21.7%)、その他および未確認 (14.9%)、一般的な Web クローラー (13.1%)、技術プロファイリング ツール (11.6%) はすべて、1 つの AI ボットよりも多くのリクエストを送信します。
最大のスタンドアロン AI カテゴリである AI エージェントは 10.5% で 5 位に位置しています。
しかし、4 つの AI カテゴリ (トレーニング クローラー、検索ボット、アシスタント、エージェント) を組み合わせると、AI ボットが 19.5% と最大の単一バケットになります。
ボット トラフィックは 3 つのストーリーに分かれています。
ファイルを消費する AI ボット (19.5%)
匿名のスクレーパーのロングテール (14.9%)
業界が監査している (12.1%)
以下でそのいくつかを詳しく見ていきます。
リクエストの 19.5% は AI ボットからのもの
llms.txt ファイルに到達するリクエストのうち、名前付き AI ボットが 19.5% を占めます。
AIボットは

llms.txt の識別可能な最大の読者層であるにもかかわらず、AI ボット タイプ別の内訳は、そのファイルがほとんどの人が念頭に置いている AI ツールを提供していないことを示しています。
ユーザーに代わって動作する AI エージェントとエージェント インフラストラクチャ、またはユーザーに代わって動作するエージェントにサービスを提供するためにクロールします。
モデル構築のためのデータを収集する AI トレーニング クローラー
ユーザーに代わってリアルタイムでウェブを閲覧する AI アシスタント
AI プラットフォームでライブ ユーザーのクエリに答えるためにページをフェッチする AI 検索ボット
*statespace-indexer: オペレーターは Statespace (エージェント インフラストラクチャ) として識別され、IP 範囲は未確認です。
エージェント Web は実際の消費者であり、リクエストの 10.5% を送信します。
AI エージェントと、AI エージェントにサービスを提供するために構築されたインフラストラクチャは、llms.txt リクエストの 10.5% を推進しており、これは他のタイプの AI ボットよりも多くなっています。
この発見は、業界の多くの人がすでに抱いていた予感と一致しています。
以前に John Mueller 氏から、llms.txt が AI コーディング エージェントの参考資料として最適であると聞きました。
Nectiv の創設者である Chris Long 氏は、たとえ llms.txt が Google 検索に役に立たないとしても、顧客が「推奨事項をソースするためにクロード コードを使用している」場合には、このファイルは有用であるとも述べています。
当社の Bot Analytics データは両方のアイデアをサポートしています。
llms.txt ファイルは、可視性を担当していると思われる検索ボットや AI ボットによって取得されることははるかに少なく、構造化された情報を検索したりユーザーに代わって動作するエージェント ツールによって取得されることがはるかに多くなります。
*statespace-indexer: オペレーターは Statespace (エージェント インフラストラクチャ) として識別され、IP 範囲は未確認です。
statespace-indexer と GPTBot とは別に、Claude-Code (Anthropic のコーディング エージェント) は、すべての AI 検索ボット、すべての AI アシスタント、およびすべての AI トレーニング クローラーをアウトフェッチしました。
トレーニング クローラーは 5.3% で 2 番目に大きい AI カテゴリです
Llms.txt ファイルは、AI 検索結果をフィードするよりも、トレーニング コーパスにフィードを提供します。

イエヴァル。
実際、AI トレーニング クローラーは、AI 検索ボットよりも 5 倍近く多く llms.txt を取得します。
したがって、llms.txt がブランドの AI 可視性に何らかの形で影響を与えるとしたら、それは取得時点ではなく上流にある可能性が高くなります。
すべてのトレーニング クローラーの中で、GPTBot は群を抜いて llms.txt の最大の取得者です。
Gemini クローラーは存在しないため、このリストには見つかりません。
Google は、通常の Googlebot によって取得されたコンテンツに基づいて Gemini をトレーニングおよび根拠付けします。Google-Extended は、パブリッシャーが使用するオプトアウト機能であり、独自のユーザー エージェントを備えたクローラーではなく、robots.txt トークンです。
Googlebot は 5 月に llms.txt ファイルを約 900 回取得しましたが、Googlebot は通常の検索インデックス作成の一環としてサイト上で発見した URL を定期的に取得するため、これらの取得は llms.txt に対する特別な関心を示しているわけではありません。サイトマップやその他のページをクロールするのと同じ方法でファイルをクロールしています。
そのコンテンツがジェミニに栄養を与えるかどうかは、私たちにはわかりません。
AI 検索ボットはかろうじて登録されており、リクエスト全体の 1.1% を占めています
弊社のデータによると、AI 検索ボットが占める割合は AI ボット リクエストのわずか 1.1% です。
AI アシスタントや AI トレーニング クローラーと合わせても、これらのボットはリクエストのわずか 8.9% にすぎません (AI エージェントよりも 1.6% 少ない)。
OAI-SearchBot、PerplexityBot、および Claude の検索クローラーを組み合わせても、数千のサイトにわたってわずか数百のフェッチしか行われませんでした。
AI による引用を増やしたいと考えて llms.txt を生成することを計画している場合は、もう一度考えたほうがよいかもしれません。
リクエストの 12% は、llms.txt を使用するのではなく、llms.txt を研究するツールからのものです
主要な AI プラットフォームが実際に llms.txt を読み取るかどうかを確認する前に、llms.txt 標準の監査、スコアリング、検証、研究を中心にエコシステム全体が形成されてきました。
3 つのカテゴリは、合計したすべてのリクエストの 12% を占めます。
GEO/AEO ツール

リクエストの 5.8% を送信
商用ツールは Web サイトをスキャンし、llms.txt の存在を多くのシグナルの 1 つとして、AI 検索とエージェント検出の準備状況をスコア付けします。
最もアクティブな CairrotReadinessBot は、2025 年後半に開始された WordPress に焦点を当てた AEO プラットフォームである Cairrot に属しています。
さらに、Framer、Lovable、Wix などの主流の Web サイト ビルダーは、自社の製品に AI 対応チェックを組み込んでいます。
Lms.txt の採用は、ウェブマスターの決定になる前にプラットフォームのデフォルトになっています。
llms.txt 発見可能性ボットはリクエストの 3.6% をカバーします
ほとんど誰も読まない llms.txt ファイルをカタログ化するツールのエコシステムがあります。
llms.txt ファイル専用に構築された専用のスキャナー、バリデーター、ディレクトリは、AI 検索ボットや AI アシスタントよりも多くのリクエストを送信します。
調査ボットはリクエストの 2.7% を送信します
データセット内で最大の単一リサーチ クローラーは、それ自体をプロンプトインジェクションサーベイ/1.0 として識別します。
誰かが、AI エージェントが取り込んで信頼するように設計された即時挿入の機会として llms.txt を体系的に研究しています。
エージェントが llms.txt ファイルを大規模に信頼することによるセキュリティへの影響についてはほとんど議論されていませんが、潜在的な悪役はすでに事件に取り組んでいます。
ゼロ AI ボットが存在しない llms.txt ファイルを「探しに行く」
AI ツールが llms.txt ファイルを検索することはありません。

[切り捨てられた]

## Original Extract

We analyzed 137K sites & found 97% of llms.txt files are never read. See why bots ignore these files & what this means for your AI visibility

We Analyzed 137K Sites: 97% of llms.txt Files Never Get Read Marketing General Marketing
We Analyzed 137K Sites: 97% of llms.txt Files Never Get Read
The number of websites linking to this post.
This post's estimated monthly organic search traffic.
Using Ahrefs Web Analytics and Bot Analytics , we analyzed the server logs and live traffic of 137K domains, plus the user agents hitting all of them.
28% of the 137K domains using Ahrefs Web Analytics publish an llms.txt file.
97% of those files received zero traffic in May 2026. Nothing fetched them at all.
96% of the requests that did reach llms.txt files came from bots.
19.5% of fetches came from named AI tools (of the 3% of files that weren’t ignored). GPTBot is top and Claude-Code is second, ahead of every AI search and assistant bot.
12% of fetches come from the industry studying itself: GEO/AEO tools, llms.txt checker tools, and researchers.
Zero requests came from AI bots for llms.txt files that don’t exist. They never go looking.
The Chrome Lighthouse llms.txt audit produced roughly 1 in 1,000 fetches.
In late May 2026, Google took both sides of the llms.txt argument in under a week.
Its new guide on optimizing for generative AI features told site owners, in a section literally titled “mythbusting”, that machine-readable files like llms.txt aren’t needed to appear in generative AI search.
Days later, the Chrome team shipped an llms.txt check inside Lighthouse’s experimental Agentic Browsing audits , with documentation explaining that without the file, agents may spend more time crawling a site to understand its structure
When Lily Ray pressed Google’s John Mueller on the contradiction , he explained that llms.txt is “not done for search.” It’s a “temporary crutch, perhaps to save some tokens” for AI coding tools parsing developer documentation—not something non-developer sites need to worry about.
He also stated that site owners who check their logs will find very little AI agent traffic.
This is something we decided to test.
Before we go any further, let’s clear up what llms.txt actually is. Llms.txt is a single index file, written in markdown, placed at a site’s root. Proposed by Jeremy Howard , co-founder of Answer.AI and fast.ai, in 2024, it summarizes what a site is and links its most important content. The idea being that LLMs and agents can use this information to orient themselves without crawling everything. The “AI visibility” framing around llms.txt came later on, attached by the SEO industry as adoption spread on the speculation that AI platforms would reward the file. Two things it is often confused with, and isn’t.
It is not the practice of publishing markdown copies of your web pages, a separate tactic with its own problems .
And despite the filename, it is not a robots.txt-style directive: it controls nothing and blocks nothing.
This study measures the index file, and only the index file.
Our study focuses on all 137,210 domains in Ahrefs Web Analytics that received traffic in May 2026.
We checked each domain root for an llms.txt returning HTTP 200, then used Ahrefs Bot Analytics to examine every request to /llms.txt paths across the population, split by HTTP response (200 vs 404) and classified by channel and individual user agent.
To rule out soft 404s and phantom files, we also confirmed each file was actual Markdown rather than HTML, and screened titles and content for error signals like “404” or “Page not found”
Ahrefs Web Analytics customers skew more technical and SEO-aware than the web at large, so treat the 28% adoption figure as an upper bound.
We did not explicitly study whether a file was well-formed against the llms.txt specification.
28% of domains publish llms.txt
Google Search’s guidance says you can skip it, the Chrome team audits for it, and Mueller calls it a stopgap for coding tools.
So amid all the mixed messages, how widespread is llms.txt actually? Among the 137K domains in our study, 28% publish these files.
More than one in four domains (38,000) in our population have adopted llms.txt, despite the fact that no major AI platform has ever committed to reading it.
Adoption has been driven by speculation that AI platforms may start consuming the file, rather than by any confirmation that they do .
97% of llms.txt files receive zero requests
Almost every llms.txt file in our study is unread.
Of the ~38,000 domains with a valid file, 97% saw no requests for it whatsoever in May.
The remaining 3% (1.1K domains) received all of the llms.txt traffic we measured.
Our data suggests John Mueller is right. Not only will you find very little AI traffic as a result of this file—you will find very little traffic, period.
If you publish an llms.txt file today, the most likely outcome by far is that nothing ever fetches it.
The 3% of files that do get read, though, get read by interesting visitors.
We’ll focus on them for the rest of the study.
96% of requests to llms.txt files come from bots
Llms.txt files are written for machines, and machines are nearly the only things reading them.
Across the files that received traffic, 96% of requests came from bots.
Humans accounted for 4%, and a chunk of those appear to be SEOs sharing llms.txt links in chat apps, where unfurl bots dutifully fetch them.
Slackbot alone fetched llms.txt files more often than PerplexityBot did.
Perplexity is one of the AI search engines llms.txt was seemingly designed to help, so finding that a chat app’s link-preview bot outfetched it speaks volumes about how much real AI search interest these files are actually generating.
77% of the bots reading llms.txt aren’t from AI tools
Many sites publish llms.txt precisely because they think it will improve their chances of appearing in ChatGPT answers, or landing Perplexity citations, or winning an AI Overview.
But our data tells a different story: 77% of the bots fetching llms.txt aren’t AI tools at all.
To understand which bots were requesting llms.txt, we classified every user agent into twelve categories.
* SEO audit tools includes Ahrefs’ own crawlers (SiteAuditBot, Ahrefs Bot, and Ahrefs Site Audit), which together account for 2,334 requests (10.6% of total). Excluding them, third-party SEO audit tools account for 2,442 requests (11.1%). Bot categories sum to 96% of total requests; the remaining 4% (930 requests) came from humans.
Individually, no AI bot category makes the top four.
SEO audit tools (21.7%), Other and unidentified (14.9%), General web crawlers (13.1%), and Tech profiling tools (11.6%) all send more requests than any one AI bot.
The biggest standalone AI category, AI agents, sits in fifth place at 10.5%.
But when you combine the four AI categories (training crawlers, retrieval bots, assistants, and agents), AI bots become the largest single bucket at 19.5%.
The bot traffic splits into three stories:
AI bots consuming the file (19.5%)
A long tail of anonymous scrapers (14.9%)
An industry auditing it (12.1%)
We’ll dig into a couple of those below.
19.5% of requests come from AI bots
Of the requests that do reach llms.txt files, named AI bots account for 19.5%.
While AI bots are the largest identifiable readership of llms.txt, the breakdown by AI bot type shows the file isn’t serving the AI tools most people have in mind.
AI agents & agentic infrastructure that act on a user’s behalf, or crawl to serve the agents that do.
AI training crawlers that collect data for model building
AI assistants that browse the web on behalf of a user in real time
AI retrieval bots that fetch pages to answer live user queries in AI platforms
*statespace-indexer: operator identified as Statespace (agentic infrastructure), IP ranges unconfirmed.
The agentic web is the real consumer, sending 10.5% of requests
AI agents, and the infrastructure built to serve them, drive 10.5% of llms.txt requests—more than any other type of AI bot.
This finding lines up with a hunch that many in the industry already had.
We heard earlier from John Mueller that llms.txt works best as reference material for AI coding agents.
Chris Long , Founder of Nectiv , has also stated that, even if llms.txt doesn’t help you in Google search, the file has utility if your customers “are using Claude Code to source recommendations”
Our Bot Analytics data supports both ideas.
We see llms.txt files being fetched far less by the search and AI bots that are seemingly responsible for visibility, and far more by the agentic tools that seek out structured information and/or act on a user’s behalf.
*statespace-indexer: operator identified as Statespace (agentic infrastructure), IP ranges unconfirmed.
Aside from statespace-indexer and GPTBot, Claude-Code (Anthropic’s coding agent), out-fetched every AI retrieval bot, every AI assistant, and every AI training crawler.
Training crawlers are the second-largest AI category at 5.3%
Llms.txt files feed training corpora more than they feed AI search retrieval.
In fact, AI training crawlers fetch llms.txt nearly 5X more than AI retrieval bots.
So if llms.txt were to in any way impact your brand’s AI visibility, it would likely be upstream—not at the point of retrieval.
Of all training crawlers, GPTBot is far and away the biggest fetcher of llms.txt.
You won’t find a Gemini crawler in this list, because it doesn’t exist.
Google trains and grounds Gemini on content fetched by regular Googlebot, and Google-Extended, the opt-out publishers use , is a robots.txt token rather than a crawler with its own user agent.
Googlebot did fetch llms.txt files ~900 times in May, but Googlebot routinely fetches any URL it discovers on a site as part of normal search indexing, so those fetches don’t indicate special interest in llms.txt—it’s crawling the file the same way it crawls a sitemap or any other page.
Whether any of that content then feeds Gemini is invisible to us.
AI retrieval bots barely register, with 1.1% of total requests
According to our data, AI retrieval bots account for just 1.1% of AI bot requests.
Even when taken together with AI assistants and AI training crawlers, these bots still count for only 8.9% of requests (1.6% less than AI agents).
OAI-SearchBot, PerplexityBot, and Claude’s search crawler combined made only a couple of hundred fetches across thousands of sites.
If you are planning on generating an llms.txt in hopes of boosting your AI citations, you may want to think again.
12% of requests come from tools studying llms.txt, not consuming it
A whole ecosystem has formed around auditing, scoring, validating, and studying the llms.txt standard, before we’ve even established whether any major AI platform actually reads it.
Three categories account for 12% of all requests combined.
GEO/AEO tools send 5.8% of requests
Commercial tools scan websites and score their readiness for AI search and agent discovery, with llms.txt presence as one of many signals.
The most active, CairrotReadinessBot, belongs to Cairrot, a WordPress-focused AEO platform launched in late 2025.
Then you have the mainstream website builders like Framer, Lovable, and Wix all baking AI-readiness checks into their products.
Lms.txt adoption has become a platform default before it’s even become a webmaster decision.
llms.txt discoverability bots cover 3.6% of requests
There’s an ecosystem of tools that catalog the llms.txt files that almost nobody else reads.
Dedicated scanners, validators, and directories built solely for llms.txt files send more requests than AI retrieval bots and AI assistants.
Research bots send 2.7% of requests
The largest single research crawler in the dataset identifies itself as prompt-injection-survey/1.0.
Someone is systematically studying llms.txt as a prompt injection opportunity that AI agents are designed to ingest and trust.
The security implications of agents trusting llms.txt files at scale have barely been discussed, and yet potential bad actors are already on the case.
Zero AI bots “go looking” for llms.txt files that don’t exist
AI tools never go looking for llms.txt files th

[truncated]

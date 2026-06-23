---
source: "https://potatometer.com/blog/yc-spring-2026-ai-reach-vs-readability"
hn_url: "https://news.ycombinator.com/item?id=48639511"
title: "Show HN: I scanned every YC Spring 2026 startup for what AI crawlers see"
article_title: "We scanned every YC Spring 2026 batch startup. AI can reach them; it can't tell what they are. | Potatometer"
author: "apswin"
captured_at: "2026-06-23T03:03:17Z"
capture_tool: "hn-digest"
hn_id: 48639511
score: 2
comments: 0
posted_at: "2026-06-23T02:40:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I scanned every YC Spring 2026 startup for what AI crawlers see

- HN: [48639511](https://news.ycombinator.com/item?id=48639511)
- Source: [potatometer.com](https://potatometer.com/blog/yc-spring-2026-ai-reach-vs-readability)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T02:40:41Z

## Translation

タイトル: Show HN: AI クローラーが何を認識しているかをすべての YC Spring 2026 スタートアップでスキャンしました
記事のタイトル: YC Spring 2026 のすべてのバッチ起動をスキャンしました。 AI は彼らに到達できます。それが何であるかはわかりません。 |ポタトメーター
説明: 私たちは YC で 197 社すべてを経営しました
HN テキスト: 「potatometer.com」を使用して、2026 年春の YC スタートアップ全 197 社の SEO / GEO / AEO 技術セットアップをスキャンして分析しました。 YC のディレクトリにある各スタートアップ リストの URL をスキャンしました。ほとんどは AI クローラーによって読み取れます。ほとんどはクローラーにそれが何であるかを伝えません。詳しくは上記のブログをご覧ください。

記事本文:
YC Spring 2026 のすべてのバッチ起動をスキャンしました。 AI は彼らに到達できます。それが何であるかはわかりません。 |ポタトメーター
GEO / 検索戦略 · 7 分読み取り · 2026-06-22
YC Spring 2026 のすべてのバッチ起動をスキャンしました。 AI は彼らに到達できます。それが何であるかはわかりません。
YC の 2026 年春バッチの 197 社すべてをライブ スキャナーで実行しました。これらはすべて目に見えない単一ページのアプリであるという、皆さんが予想するような恐ろしい見出しは誤りです。本当のギャップはもっと静かで、はるかに一般的です。AI はそれらのほとんどに到達できますが、それらが何であるかを機械可読な言葉でラベル付けしません。
バッチを分析するには、まずバッチをリストする必要があるため、 ycombinator.com/companies にある YC 独自のディレクトリを取得しました。それは、1 行の表示テキスト (ページ タイトル) を含む約 33 キロバイトの HTML として返されました。すべての企業、すべてのリンクは、後で JavaScript として読み込まれます。 HTML を読み取り、JavaScript を実行しないプログラムにとって、YC ディレクトリは空白のページです。
YC 自身の会社の一部がまさにその問題を抱えているため、そこから始めるのは適切です。ほとんどはそうではありません。私たちは 197 件すべてをスキャンして、どれがいくつあるかを正確に調べました。その答えは、創設者が自分のサイトについて尋ねることができる 2 つのまったく異なる質問に分かれています。
AIはあなたに連絡を取ることができますか？ほとんどの場合、そうです。 AIはあなたが何者であるかを判断できますか?ほとんどはノーです。
バッチのほとんどに到達可能です。その部分は大丈夫です。
憂慮すべきバージョンよりも正直なバージョンの方が役立つため、結果を先に示します。私たちが評価できた 195 のサイト (2 つはスキャン時にサーバー エラーが返され、保留されました) のうち、164 は実際のコンテンツを HTML で直接提供していました。 JavaScript を実行しないクローラーは、今日 AI が回答するフィードのほとんどを記述しており、単一のスクリプトが実行される前に、これらのサイト上の中央値 682 ワードを読み取ります。読みやすいです。 「誰もが見えないSPA」のストーリーはシンプル

このバッチには当てはまりません。
11 分の 1 はクローラーが認識できない実際のサイトです
195 個のうち 17 個、つまり約 11 分の 1 が空の殻です。フロントドアの HTML はほとんど空ですが、レンダリングされたページはいっぱいです。コンテンツは存在します。これはクライアント側の JavaScript の背後でロックされており、それを実行しないクローラーには見えません。あるサイトでは、レンダリングされたホームページが 900 ワードを超えている間に、クローラーに 1 単語のテキストを提供します。別のサービスでは 9 件が提供され、1,200 件を超えています。これらはパークされたドメインではありません。これらは、生の HTML を読み取る AI クローラーがコピーに到達することのない本物の製品です。
サイトをシェルとしてカウントしたのは、サイトが JavaScript でレンダリングされ、フロントドアで 50 ワード未満の単語を提供し、スクリプトが実行されると 100 以上の単語をレンダリングする場合に限られます。最後の条件が重要です。コンテンツは実際に存在しており、単に隠されているだけです。スキャナー独自のフロントドア文字数とレンダリング文字数を使用して測定しました。これは、ユーザーが目にするのと同じ数です。この 2 つのギャップは、まさに非 JS クローラーが見逃しているものです。これは決定的な事実であり、順位に関する予測ではありません。
別の少数の製品は単に発売されていないだけです
さらに 11 のサイトは、フロントドアでもレンダリング後でもコンテンツが軽いです。これらはクローラーの問題ではありません。これらは、近日公開ページ、順番待ちリスト、ステルス スプラッシュ スクリーンなど、発売前のものです。折り畳むと膨らんでしまうため、別々に数え、上の図からは意図的に外しました。さらに 3 つは薄い静的ページで、2 つはオフホスト リダイレクトでした。これも脇に置かれていました。シェル番号は、隠された実際のコンテンツのみをカウントし、それ以外はカウントしません。
ここでさらに難しい質問になります。AI はあなたが何者であるかを判断できるでしょうか?
ページに到達することは、それを理解することと同じではありません。これらのホームページの 1 つを取得する AI は散文の壁に直面し、通常、フィンテック API を参照しているかどうかを示す機械可読ラベルはありません。

アプリ、または開発ツール。テキストは読むためにあります。機械がそれを並べ替えたり、引用したり、それに基づいて動作させるための構造化された信号は、ほとんどの場合そうではありません。
バッチでラベル付けされるものとスキップされるもの
195 サイトのシェア、各軸は単一の測定チェック: クローラー アクセス 91%、HTTPS 97%、サイトマップ 68%、Canonical 56%、画像代替 54%、スキーマ (任意) 50%、ページ上の FAQ 28%、FAQ スキーマ 19%。
ホイールの上部 (リーチ: クローラーアクセス、HTTPS) が強力です。一番下 (ラベル: スキーマ、FAQ マークアップ) は、バッチが間引かれる場所です。すべての値は、複合スコアではなく、195 サイトにわたってカウントされた「はい/いいえ」チェックです。
ラベル付けのギャップを数値で示します。バッチの半分 (50%) には、構造化データのマークアップが含まれています。マシンが実際に使用できるタイプを搭載しているのは 41% だけです。最も回答しやすい形式である FAQ マークアップは 19 パーセントに表示されます。どの URL が本物であるかをクローラーに伝える正規タグは 56% に設定されています。これらはどれもエキゾチックなものではなく、どれも難しいものではありません。これは、「AI が読めるページ」を「AI が正しく引用できるページ」に変える機械可読ラベルです。
誇張するのは簡単なので、この主張について正確に言うと、現代のモデルは散文を完璧に読み取り、そこから多くのことを推測することができます。スキーマが欠落していてもページが理解不能になることはありません。これにより、製品の種類、FAQ、価格、回答など、代わりに伝えられる可能性のある事柄を機械に推測させます。AI ショッピングおよび回答エージェントがその情報を単に要約するのではなく、その情報に基づいて行動を開始するため、これらのことがより重要になります。
すべてのサイトがすでに読み込まれていたため、他にもいくつか確認しました。
AI クローラーをブロックします。約 9% に相当する 18 のサイトが、robots.txt 内の少なくとも 1 つの主要な AI クローラーをブロックしています。 GPTBot と ClaudeBot が最も頻繁にブロックされ、PerplexityBot がそれに続き、毎回同じ 18 サイトがブロックされ、大規模なブロックが行われます。

検討されているものよりも優れています。すべてをブロックするものはありません。これらのケースのほとんどでは、ブロックは決定というよりも、レビューされていないフレームワークまたはテンプレートのデフォルトのように見えます。もしあなたがそのような人であれば、30 秒ほど見てみる価値があります。
基本、ほとんど扱います。 97% が HTTPS を提供しています。クローラーはロボットのブロックなしで 91% に到達できます。 3 分の 2 がサイトマップを公開しています。基本的な部分の弱点は、canonical タグが 56 パーセント、画像代替テキストが 54 パーセントで、どちらも簡単ですが、どちらも中途半端です。
注意深く読むべき数字が 1 つあります。私たちがパフォーマンス データを取得したサイトのうち、厳密な Core Web Vitals バー (最大コンテンツフル ペイントが 2.5 秒未満、さらにレイアウト シフトが少なく、ブロック時間が短い、この 3 つをすべて同時に達成) をクリアしているのは 4% だけでした。それは壊滅的なことのように聞こえますが、ほとんどの場合はそうではありません。これらは、キャッシュを使用せずにコールド負荷で取得された実験室の測定値であり、実際のユーザーが経験するものに対して悲観的であり、スコアリングにおける単一の低重要度のチェックです。 「これらのサイトの 96% が遅いと感じている」のではなく、「ラボの厳しいしきい値をクリアしている人はほとんどいない」と読んでください。
私たちは最悪のものではなく、最良のものに名前を付けています。これら 5 つはリーチしてラベルを付けます。これらは、 robots.txt 内の偶発的なオウンゴールをスキップしながら、実質的なコンテンツをクローラーに直接提供し、クリーンな正規データを設定し、実際の構造化データを送信します。スキャンした正確な URL の結果が表示されるため、自分で確認する場合は、同じアドレスをスキャンしてください。 apex フォームと www フォームではそれぞれのスコアが同じなので、説明する必要はありません。
共通点は鈍い、それがポイントだ。クローラーが見える場所にコンテンツをレンダリングし、スキーマで記述し、 robots.txt でドアを閉じません。ほぼすべてのユーザーがまだつまずいているのは、厳格な Core Web Vitals バーであり、ほぼバッチ全体で同じチェックが失敗します。どれも難しいことではありません。それはdです

これは、マーケティング サイトをアプリではなくドキュメントのように扱うことで得られる恩恵です。
バッチに参加している場合、または同様のものを構築している場合は、これは短いバージョンです。ほとんどが午後です。
フロントドア HTML がシェルの場合は、マーケティング ページを事前レンダリングまたはサーバー レンダリングして、クローラーが読み取るページが人が読むページと同じになるようにします。組織と実際の製品タイプ ( SoftwareApplication または Product ) スキーマを追加します。これは、バッチ内の唯一の最大のギャップです。 FAQPage マークアップを使用した本物の FAQ を追加します。これは、最も回答しやすい形式であり、最もまれな形式 (19%) です。 robots.txt を grep して、偶発的な GPTBot または ClaudeBot ブロックを探します。実際の答えを最初の 150 単語に入力してください。また、Google のために llms.txt を配布するように言われた場合は、スキップしてください。Google は記録上、そのファイルを検索には使用していないと述べており、2026 年の AI ガイダンスで役に立たない戦術として直接その名前を挙げています。これは、一部のコーディング ツールやアシスタントが取得する、安価でエージェントが読み取り可能なインフラストラクチャであるため、無価値ではなく、単に Google の手段ではないだけです。
私たちがどのようにしてこれを行ったのか、そして何を信じるべきか
これは最も重要な部分なので、数字について議論する前に読んでください。
すべての数値は 1 つのソースから得られます。それは、誰もが自分のサイトで実行できるものと同じものであるライブ スキャン エンジンです。ここにある数値は、AI によって推定、モデル化、または書き込まれたものではありません。このエンジンは各サイトを取得し、実際のヘッドレス ブラウザで JavaScript を多用したサイトをレンダリングし、決定的な事実、クローラーが見た単語の数、robots.txt が特定のボットをブロックするかどうか、どのスキーマ タイプが存在するかを報告します。再実行すると、1 ～ 2 ポイント以内に同じ答えが得られます。唯一の実行間変数はサードパーティのパフォーマンス データです。
バッチ。 YC 2026 年春、2026 年 6 月 22 日に YC 自身のディレクトリから列挙。 197 社、それぞれの Web サイトを正確にスキャン

そのディレクトリにリストされます。私たちは 195 を評価しました。 2 つ (Gravy と Drip) はスキャン時にサーバー エラーを返したため、推測ではなく未評価としてマークしました。
閾値は、あなたが同意できないように記載されています。クローラーがフロント ドアで 50 単語未満を認識した場合、そのサイトはコンテンツ ライトであり、レンダリングされたページに少なくとも 100 ワードがある場合はコンテンツ リッチです。 1/11 シェルの数値は、JavaScript でレンダリングされ、フロント ドアではコンテンツ ライトであり、レンダリングされた後はコンテンツ リッチ、つまり非表示になっている実際のコンテンツのみをカウントします。しきい値を移動すると、正確なカウントが少し変わります。形状はそうではありません。
限界。サイト全体ではなく、リストされている各企業のホームページをスキャンしました。スコアはスキャンした URL に固有であり、サイトの頂点で異なるコンテンツを提供し、www フォームのスコアはそれぞれで異なる場合があります。レンダリングに失敗した場合は、サイトを空白のスコアではなく評価なしとしてマークしました。 Core Web Vitals はラボの指標であり、実際のユーザー データに対して悲観的に実行されます。それがすべての方法です。
自分のサイトがこれらのグループのどれに属しているかを知りたい場合は、potatometer.com で調べてください。所要時間は約 30 秒で、JavaScript が実行される前にクローラーが見たものを正確に表示します。
Y Combinator 起動ディレクトリ (2026 年春バッチ)
Google 検索セントラル、2026 年の AI 機能最適化ガイダンス（llms.txt 上）
ポタトメーター スキャナー (ここにあるすべての数値の背後にあるエンジン)

## Original Extract

We ran all 197 companies in YC

Used 'potatometer.com' to scan and analyze all All 197 YC Spring 2026 startups on their SEO / GEO / AEO technical setup. I scanned the URL each startup lists in YC's directory. Most are readable by AI crawlers. Most don't tell a crawler what they are. Read more in the blog above.

We scanned every YC Spring 2026 batch startup. AI can reach them; it can't tell what they are. | Potatometer
GEO / Search Strategy · 7 min read · 2026-06-22
We scanned every YC Spring 2026 batch startup. AI can reach them; it can't tell what they are.
We ran all 197 companies in YC's Spring 2026 batch through our live scanner. The scary headline you would expect, that they are all invisible single-page apps, is false. The real gap is quieter and far more common: AI can reach most of them, but they do not label what they are in machine-readable terms.
To analyze the batch we first had to list it, so we fetched YC's own directory at ycombinator.com/companies . It came back as about 33 kilobytes of HTML containing one line of visible text: the page title. Every company, every link, loads afterward as JavaScript. To a program that reads HTML and does not run JavaScript, the YC directory is a blank page.
That is a fitting place to start, because a small slice of YC's own companies have exactly that problem. Most do not. We scanned all 197 to find out precisely how many, and which, and the answer splits into two very different questions a founder can ask about their own site.
Can an AI reach you? Mostly yes. Can an AI tell what you are ? Mostly no.
Most of the batch is reachable. That part is fine.
Here is the result up front, because the honest version is more useful than the alarming one. Of the 195 sites we could evaluate (two returned server errors at scan time and were set aside), 164 serve real content directly in their HTML. A crawler that does not run JavaScript, which describes most of what feeds AI answers today, reads a median of 682 words on these sites before a single script runs. They are readable. The "everyone is an invisible SPA" story is simply not true for this batch.
1 in 11 is a real site a crawler cannot see
Seventeen of the 195, about 1 in 11, are empty shells. Their front-door HTML is near empty, but their rendered page is full. The content exists. It is locked behind client-side JavaScript, invisible to any crawler that does not execute it. One site serves a single word of text to a crawler while its rendered homepage runs past 900 words. Another serves nine and renders past 1,200. These are not parked domains. They are real products whose copy an AI crawler reading raw HTML would never reach.
We counted a site as a shell only when it is JavaScript-rendered, serves under 50 words at the front door, and renders 100 or more once the script runs. That last condition is the point: the content is really there, just hidden. We measured it with the scanner's own front-door-versus-rendered word count, the same number any user sees. The gap between the two is exactly what a non-JS crawler misses. It is a deterministic fact, not a prediction about ranking.
A separate handful simply has not launched
Eleven more sites are content-light both at the front door and after rendering. These are not a crawler problem. They are pre-launch: coming-soon pages, waitlists, stealth splash screens. We counted them apart and kept them out of the figure above on purpose, because folding them in would inflate it. Three more were thin static pages and two were off-host redirects, also set aside. The shell number counts only real content that is hidden, nothing else.
Now the harder question: can an AI tell what you are?
Reaching a page is not the same as understanding it. An AI that fetches one of these homepages gets a wall of prose and, usually, no machine-readable labels telling it whether it is looking at a fintech API, a healthcare app, or a dev tool. The text is there to read. The structured signals that let a machine sort, quote, and act on it mostly are not.
What the batch labels, and what it skips
Share of 195 sites, each axis a single measured check: Crawler access 91%, HTTPS 97%, Sitemap 68%, Canonical 56%, Image alt 54%, Schema (any) 50%, On-page FAQ 28%, FAQ schema 19%.
The top of the wheel (reach: crawler access, HTTPS) is strong. The bottom (labeling: schema, FAQ markup) is where the batch thins out. Every value is a yes/no check counted across 195 sites, not a composite score.
The labeling gap, in plain numbers. Half the batch, 50 percent, carries any structured-data markup at all; only 41 percent carry a type a machine can actually use. FAQ markup, the most answer-friendly format there is, appears on 19 percent. A canonical tag, which tells a crawler which URL is the real one, is set on 56 percent. None of this is exotic, and none of it is hard. It is the machine-readable labeling that turns "a page an AI can read" into "a page an AI can quote correctly."
To be precise about the claim, since it is easy to overstate: a modern model can read prose perfectly well and infer a great deal from it. Missing schema does not make a page incomprehensible. It makes a machine guess at things it could instead be told, the product type, the FAQ, the price, the answer, which matters more as AI shopping and answer agents start acting on that information rather than just summarizing it.
A few other things we checked, since every site was already loaded.
Blocking AI crawlers. Eighteen sites, about 9 percent, block at least one major AI crawler in robots.txt . GPTBot and ClaudeBot are blocked most often, PerplexityBot close behind, and it is the same eighteen sites each time, a wholesale block rather than a considered one. None block everything. In most of these cases the block reads like an unreviewed framework or template default rather than a decision, which is worth a thirty-second look if you are one of them.
The basics, mostly handled. Ninety-seven percent serve HTTPS. Crawlers can reach 91 percent without a robots block. Two-thirds publish a sitemap. The weak spots inside the basics are canonical tags at 56 percent and image alt text at 54 percent, both easy, both half-done.
One number to read carefully. Only 4 percent of the sites we had performance data for clear the strict Core Web Vitals bar (largest-contentful-paint under 2.5 seconds, plus low layout shift and low blocking time, all three at once). That sounds catastrophic and mostly is not. These are lab measurements taken on a cold load with no caching, which run pessimistic against what real users experience, and it is a single low-weight check in our scoring. Read it as "almost nobody clears the strict lab threshold," not "96 percent of these sites feel slow."
We are naming the best, not the worst. These five reach and label: they serve substantial content directly to crawlers, set a clean canonical, and ship real structured data, while skipping the accidental own-goals in robots.txt . Results are shown for the exact URL we scanned, so if you check one yourself, scan the same address. Each scored identically on the apex and www forms, so there is nothing to explain away.
The common thread is dull, and that is the point. They render their content where a crawler can see it, describe it with schema, and do not close the door in robots.txt . The one thing nearly all of them still trip on is the strict Core Web Vitals bar, the same check almost the entire batch fails. None of this is hard. It is the default you get from treating the marketing site like a document rather than an app.
If you are in the batch, or building anything like it, this is the short version. Most of it is an afternoon.
If your front-door HTML is a shell, pre-render or server-render the marketing pages so the page a crawler reads is the page a person reads. Add Organization plus your real product type ( SoftwareApplication or Product ) schema, the single biggest gap in the batch. Add a genuine FAQ with FAQPage markup, the most answer-friendly format and the rarest, at 19 percent. Grep your robots.txt for an accidental GPTBot or ClaudeBot block. Put your actual answer in the first 150 words. And if you have been told to ship an llms.txt for Google's sake, skip it: Google has said on the record it does not use the file for Search, and named it directly in its 2026 AI guidance as a tactic that does not help. It is cheap, agent-readable infrastructure that some coding tools and assistants do fetch, so it is not worthless, it is just not a Google lever.
How we did this, and what to distrust
This is the part that matters most, so read it before you argue with the numbers.
Every figure comes from one source: our live scanning engine, the same one anyone can run on their own site. No number here was estimated, modeled, or written by an AI. The engine fetches each site, renders JavaScript-heavy ones with a real headless browser, and reports deterministic facts, how many words a crawler sees, whether robots.txt blocks a given bot, which schema types are present. Re-run it and you get the same answer within a point or two; the only run-to-run variable is third-party performance data.
The batch. YC Spring 2026, enumerated from YC's own directory on 2026-06-22. 197 companies, each scanned at the exact website it lists in that directory. We evaluated 195; two (Gravy and Drip) returned server errors at scan time, so we marked them not evaluated rather than guess.
The thresholds, stated so you can disagree. A site is content-light if a crawler sees fewer than 50 words at the front door, content-rich if the rendered page has at least 100. The 1-in-11 shell figure counts only sites that are JavaScript-rendered, content-light at the front door, and content-rich once rendered, that is, real content that is hidden. Move the thresholds and the exact count shifts a little. The shape does not.
The limits. We scanned each company's listed homepage, not its whole site, and scores are specific to the URL scanned, a site that serves different content on its apex and www forms can score differently on each. Where rendering failed we marked the site not evaluated rather than score a blank. Core Web Vitals are lab metrics and run pessimistic against real-user data. That is the whole method.
If you want to know which of these groups your own site is in, scan it at potatometer.com. It takes about thirty seconds and shows you exactly what a crawler sees before any JavaScript runs.
Y Combinator startup directory (Spring 2026 batch)
Google Search Central, 2026 AI features optimization guidance (on llms.txt)
Potatometer scanner (the engine behind every number here)

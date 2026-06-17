---
source: "https://mojodojo.io/blog/ai-crawlers-match-googlebot-traffic"
hn_url: "https://news.ycombinator.com/item?id=48572990"
title: "AI crawlers now match Googlebot in traffic to our site"
article_title: "AI crawlers now match Googlebot in traffic to our site | Mojo Dojo | Mojo Dojo"
author: "zenincognito"
captured_at: "2026-06-17T18:58:19Z"
capture_tool: "hn-digest"
hn_id: 48572990
score: 5
comments: 0
posted_at: "2026-06-17T16:42:43Z"
tags:
  - hacker-news
  - translated
---

# AI crawlers now match Googlebot in traffic to our site

- HN: [48572990](https://news.ycombinator.com/item?id=48572990)
- Source: [mojodojo.io](https://mojodojo.io/blog/ai-crawlers-match-googlebot-traffic)
- Score: 5
- Comments: 0
- Posted: 2026-06-17T16:42:43Z

## Translation

タイトル: AI クローラーがサイトへのトラフィックで Googlebot と一致するようになりました
記事のタイトル: AI クローラーがサイトへのトラフィックで Googlebot と一致するようになりました |モジョ道場 |モジョ道場

記事本文:
AI クローラーがサイトへのトラフィックで Googlebot と一致するようになりました |モジョ道場 | Mojo Dojo メインコンテンツにスキップ メニュー 無料の監査を入手 ホーム
AI クローラーがサイトへのトラフィックで Googlebot と一致するようになりました
AI クローラーがサイトへのトラフィックで Googlebot と一致するようになりました
Hacker News のトップページに到達した後、あなたのサイトにどのようなトラフィックとボットが流入しますか? AI ボットは Google と同じくらいクロールしますか?最も攻撃的なボットはどれですか?最も少なくて倹約できるのはどれですか?
最近の投稿がハッカーニュースの一面を飾りました。 Google があなたを検索品質評価者に任命しただけで、報酬は発生しません。
これは私が知りたかったことでしたので、このサイトに小さなボット検出器を構築しました。もちろんボットのクロードを使用します。
このコードは基本的に 30 日間のトラフィックを測定し、コード内に次のデータを出力する簡単なアナライザーを構築しました。
24 の異なるボットからの 240,060 回の訪問全体で、AI クローラー (トラフィックの 35.0%) が検索エンジン (35.0%) と同率でトップの座を獲得しました。
この業界に興味があるなら、それが見出しになります。残りはデータのオタクです。
認識されたボット ユーザー エージェントからのすべての訪問を単一のテーブルに記録します。フィールドはボット名、カテゴリ、パス、日、ヒット数です。 240,060 件のヒット、24 件のボット、5,074 件の個別の URL がタッチされました。
ボットを 5 つのカテゴリに分類します: 検索 (Google、Bing、DuckDuckGo など)、AI (GPTBot、ClaudeBot、PerplexityBot、AmazonBot など)、SEO (Ahrefs、Semrush、Majestic、Screaming Frog)、ソーシャル (LinkedIn、Facebook、Twitter、Pinterest)、アーカイブ (Internet Archive)。
AI クローラーは検索エンジンと同率でそれぞれ 35.0% でした。
期間中の平均した 1 日あたりのヒット数。
パスごとのヒット数。下 = 訪問して出発します。
再スクレーパー。 URL はほとんどなく、繰り返し叩かれました。
30 日間にわたってほとんど表示されないボット。
出典: mojodojo.io の実稼働データベース上の nexus_bot_visits テーブル。ヒット カウンターを含む (ボット、パス、日) ごとに 1 行。
ボットの検出

: ユーザー エージェントの部分文字列が、32 の既知のボット パターンの厳選されたリストと照合されます。不明な/なりすましの UA は除外されます。
2026 年の B2B AI 可視性の現状
私たちは 712 社の B2B 企業を監査しました。 89% は AI 検索では表示されません。業界のベンチマーク、リーダーの 4 つのパターン、および 30 日間の修正リスト。
Ajay は Mojo Dojo の CTO であり、デジタル マーケティングで 15 年以上の経験があります。
Google があなたを検索品質評価者にしました。給料は支払われません。
TL;DR AI はクリックを無効にします。クリックがなくなると、リンクに対するインセンティブが失われます。これでGoogleは去ります...
Meta はオーガニックリーチを定額料金で販売しています。
15 年間、エージェンシーはクライアントに、他人の所有物を使った「アーンドメディア」という夢を販売してきました。私たちは...
優れた価値提案を書くにはどうすればよいでしょうか? (テンプレート付属)
価値提案とは何ですか?優れた価値提案により、あなたのビジネスは次のような選択肢の中で当然の選択肢となります...
スイート 20/レベル 5、150 Albert Rd
サウスメルボルン VIC 3205
730 リップル ブルック Ln
イリノイ州エルギン 60120
ユニット 61、38-40 オーバル
ロンドン E2 9DT
© 2026 モジョ道場 .無断転載を禁じます。

## Original Extract

AI crawlers now match Googlebot in traffic to our site | Mojo Dojo | Mojo Dojo Skip to main content Menu Get Free Audit Home
AI crawlers now match Googlebot in traffic to our site
AI crawlers now match Googlebot in traffic to our site
What traffic & bots flow to your site after you make it to the front page of Hacker News? Do AI bots crawl as much as Google? Which is the most aggressive bot? Which is the least and thriftiest?
A recent post hit the front page on hacker news. Google just made you a search quality rater.You won't get paid.
This was something I wished to know so I built a little bot detector on this site. Of course using Claude, the bot.
The code basically measured the traffic over 30 days and I built a simple analyser within the code to output the following data.
Across 240,060 visits from 24 distinct bots, AI crawlers (35.0% of traffic) tied search engines (35.0%) for the top slot.
If you care about this industry then that's the headline. The rest is nerding in data.
We log every visit from a recognised bot user-agent into a single table. The fields are bot name, category, path, day, hit count. 240,060 hits, 24 bots, 5,074 distinct URLs touched.
We bucket bots into five categories: search (Google, Bing, DuckDuckGo, etc.), ai (GPTBot, ClaudeBot, PerplexityBot, AmazonBot, etc.), seo (Ahrefs, Semrush, Majestic, Screaming Frog), social (LinkedIn, Facebook, Twitter, Pinterest), and archive (Internet Archive).
AI crawlers tied search engines at 35.0% each.
Hits per day, averaged over the period.
Hits per path. Lower = visit and leave.
Re-scrapers. Few URLs, hammered repeatedly.
Bots that barely show up across 30 days.
Source: nexus_bot_visits table on mojodojo.io's production database. One row per (bot, path, day) with a hit counter.
Bot detection: User-Agent substring match against a curated list of 32 known bot patterns. Unknown / spoofed UAs are excluded.
State of B2B AI Visibility 2026
We audited 712 B2B companies. 89% are invisible in AI search. Industry benchmarks, the 4 patterns of leaders, and a 30-day fix list.
Ajay is the CTO of Mojo Dojo with over 15 years of experience in digital marketing.
Google just made you a search quality rater. You won't get paid.
TL;DR AI is killing the click. Death of the click kills the incentive for links. This leaves Google...
Meta is selling organic reach for a flat fee.
For fifteen years, agencies sold clients a dream: "earned media" on someone else's property. We...
How to write a great value proposition? (Template included)
What is value proposition? A great value proposition makes your business the obvious choice amongst...
Suite 20/Level 5, 150 Albert Rd
South Melbourne VIC 3205
730 Ripple Brook Ln
Elgin, IL 60120
Unit 61, 38-40 The Oval
London E2 9DT
© 2026 Mojo Dojo . All rights reserved.

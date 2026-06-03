---
source: "https://julienreszka.com/blog/rss-is-back-ai-agents-are-reading-it/"
hn_url: "https://news.ycombinator.com/item?id=48375673"
title: "Now AI agents need what RSS does"
article_title: "RSS Is Back. AI Agents Are Reading It. — Julien Reszka"
author: "julienreszka"
captured_at: "2026-06-03T00:36:16Z"
capture_tool: "hn-digest"
hn_id: 48375673
score: 50
comments: 21
posted_at: "2026-06-02T20:19:11Z"
tags:
  - hacker-news
  - translated
---

# Now AI agents need what RSS does

- HN: [48375673](https://news.ycombinator.com/item?id=48375673)
- Source: [julienreszka.com](https://julienreszka.com/blog/rss-is-back-ai-agents-are-reading-it/)
- Score: 50
- Comments: 21
- Posted: 2026-06-02T20:19:11Z

## Translation

タイトル: AI エージェントには RSS の機能が必要です
記事のタイトル: RSS が帰ってきた。 AI エージェントはそれを読んでいます。 — ジュリアン・レスカ
説明: Google Reader は 2013 年に廃止され、誰もがそれを呼びました。しかし、RSS がポッドキャスティングの動力を止めたことはなく、AI エージェントはまさに RSS の機能を必要としています。

記事本文:
RSSが戻ってきました。 AI エージェントはそれを読んでいます。 — ジュリアン・レスカ
← ブログ
RSSが戻ってきました。
AI エージェントはそれを読んでいます。
ジュリアン・レスカ · 2026-05-30T21:07
Google Reader は 2013 年に廃止され、誰もがそれを呼びました。しかし、RSS がポッドキャスティングの動力を止めたことはなく、AI エージェントはまさに RSS の機能を必要としています。
RSS は 2013 年に Google が Reader を閉鎖したときに廃止されたと宣言されました。追悼文は時期尚早であり、原因について間違っていた。 RSS が機能を停止することはありませんでした。 RSS が提供できないもの、つまり可変報酬スケジュールの中毒性のあるランダム性がソーシャル アルゴリズムによって提供されたため、人間がコンテンツを発見する主要な方法ではなくなりました。人間はそれが魅力的だと感じます。エージェントはそうではありません。
競合他社のリリースを監視したり、規制の変更を追跡したり、研究を要約したりする AI エージェントは、驚かれたくないでしょう。それは次のことを望んでいます:
新しいものに関する決定的なリスト
推測せずに解析できる構造化フォーマット
広告関係に関連付けられた料金制限なし
公開コンテンツを保護する認証の壁がない
RSS は 4 つすべてを提供します。ソーシャル プラットフォーム API では、これらはいずれも提供されません。そうした場合、四半期ごとにアクセスを取り消し、料金を請求します。 RSS フィードは、プルベース、オープン、そして一貫性を備えていますが、アルゴリズムの仕事には一貫性がないことが求められているため、アルゴリズムの設計ではそうならないように設計されています。
RSS が実際には死んでいなかったことを示す最も明白な証拠は、ポッドキャスティングです。すべてのポッドキャスト アプリ (Spotify、Apple、Overcast、Pocket Casts) は、RSS フィードからエピソード ファイルとメタデータを取得します。 250 億ドルのポッドキャスト業界は、2002 年に公開されたプロトコルに基づいて運営されています。オープンで、無料で、仲介者がなく、アクセスを交渉するものがないため、誰もそれを破壊しませんでした。エピソードはフィード内の URL にあります。これまでと同様です。
同じロジックが、エージェントが確実に利用する必要があるあらゆる文書コンテンツにも拡張されるようになります。言語モデルの取得

ユーザーのクエリのコンテキスト、新しい提出物をチェックする監視エージェント、ニュースレターを取り込む要約ツールなど、それらすべてが、新しいコンテンツの予測可能で構造化された時系列のリストから恩恵を受けます。 RSS はこれですべてです。問題は、コンテンツがその方法でアクセスできるのか、それとも人間の注意を引くように設計されたシステム内にコンテンツが存在し、プログラムによるアクセスを積極的に低下させるのかということです。
コンテンツの RSS フィードがない場合は、それを公開します。ニッチな分野のソースを監視するエージェントは、アルゴリズムに依存するページを見つける前に、構造化フィードを見つけます。
ソーシャル プラットフォーム内に存在するため、AI エージェントやアグリゲーターが確実にアクセスできないコンテンツを公開していませんか?
はい。今週はクライアント向けに競合他社監視エージェントを構築します。 RSS フィードのあるサイトの接続には 30 秒かかります。RSS フィードのないサイトでは、再設計するたびに壊れる壊れやすいスクレイピング リグが必要です。信頼性の差は非常に大きくなります。
プッシュバック: エージェントは RSS を解析できるのと同じようにスクレイピングできます。ページが HTML をロードすると、有能なエージェントがそれを取得します。構造化フィードの議論では、エージェントは脆弱すぎて実際の Web ページを処理できないと想定されていますが、これはますます真実ではありません。
スクレイピングは、サイトが CAPTCHA を追加するか、そのマークアップを変更するか、既知のエージェント ユーザー エージェントのブロックを開始するまで機能します。 RSS は設計上決定的です。内容は発行者が制御しており、ホームページが再設計されても壊れることはありません。より大きな問題はメンテナンスです。脆弱性のスクレイピングはソースの数に比例します。スクレーパーを介して 10 個のサイトに配線すると、10 個の障害モードが用意されます。 10 個の RSS フィードを接続すれば、一度セットアップすればメンテナンスはほぼゼロになります。これは、大規模な監視エージェントを構築する開発者にとっての経済的な議論です。
上記を読んだ後、3 週間前にニュースレターに RSS フィードを追加しました

ニッチな技術コンテンツを集約する ut エージェント。私たちが何もしなくても、2 つのアグリゲーターが 1 週間以内に自動的にそれを取得しました。フィードはディストリビューションでした。
すべてのコメントは作成者によって手動で管理されます。
新しい投稿をメールで受け取るには登録してください →
自動化により簡単な仕事が奪われ、難しい仕事はあなたに任せられる
未来の仕事は邪魔になりつつある
WebMCP は新しいモバイルファーストです
すべての図を参照 →
すべての神話を参照 →

## Original Extract

Google Reader died in 2013 and everyone called it. But RSS never stopped powering podcasting, and now AI agents need exactly what RSS does.

RSS Is Back. AI Agents Are Reading It. — Julien Reszka
← Blog
RSS Is Back.
AI Agents Are Reading It.
Julien Reszka · 2026-05-30T21:07
Google Reader died in 2013 and everyone called it. But RSS never stopped powering podcasting, and now AI agents need exactly what RSS does.
RSS was declared dead in 2013 when Google shut down Reader. The eulogies were premature and wrong about the cause. RSS never stopped working. It stopped being the primary way humans discovered content, because social algorithms offered something RSS could not: the addictive randomness of a variable reward schedule. Humans find that irresistible. Agents do not.
An AI agent that monitors competitor releases, tracks regulatory changes, or summarises research does not want to be surprised. It wants:
a deterministic list of what is new
a structured format it can parse without guessing
no rate limits tied to an advertising relationship
no authentication wall protecting public content
RSS provides all four. Social platform APIs provide none of them. When they do, they revoke access on a quarterly basis and charge for it. An RSS feed is pull-based, open, and consistent in a way that no algorithm is designed to be, because an algorithm's job is to be inconsistent.
The clearest evidence that RSS was never really dead is podcasting. Every podcast app (Spotify, Apple, Overcast, Pocket Casts) pulls episode files and metadata from RSS feeds. The $25 billion podcast industry runs on a protocol published in 2002. Nobody disrupted it because there was nothing to disrupt: open, free, no middleman, nothing to negotiate access to. The episode is at the URL in the feed, always has been.
The same logic will now extend to any written content that agents need to reliably consume. A language model retrieving context for a user query, a monitoring agent checking for new filings, a summarisation tool ingesting newsletters: all of them benefit from a predictable, structured, chronological list of new content. That is all RSS is. The question is whether your content is reachable that way, or whether it lives inside a system that was designed for human attention and actively degrades programmatic access.
Publish an RSS feed for your content if you don't have one. Agents that monitor sources in your niche will find structured feeds before they find algorithm-dependent pages.
Are you publishing content that AI agents and aggregators can't reliably reach because it lives inside a social platform?
Yes. Building a competitor monitoring agent for a client this week. The sites that have RSS feeds take thirty seconds to wire in. The ones without them require fragile scraping rigs that break on every redesign. The delta in reliability is enormous.
Pushback: agents can scrape just as well as they can parse RSS. If a page loads HTML, a capable agent retrieves it. The structured feed argument assumes agents are too fragile to handle real web pages, which is increasingly not true.
Scraping works until the site adds a CAPTCHA, changes its markup, or starts blocking known agent user agents. RSS is deterministic by design: the publisher controls what's in it and it doesn't break when the homepage gets redesigned. The bigger issue is maintenance: scraping fragility scales with the number of sources. Wire in ten sites via scraper and you have ten failure modes to babysit. Wire in ten RSS feeds and you have near-zero maintenance once they're set up. That's the economic argument for any dev building a monitoring agent at scale.
Added an RSS feed to our newsletter three weeks ago after reading about agents that aggregate niche technical content. Two aggregators picked it up automatically within a week without us doing anything. The feed was the distribution.
All comments are manually moderated by the author.
Subscribe to get new posts by email →
Automation Takes the Easy Jobs and Leaves You the Hard Ones
The Future of Work Is Getting Out of the Way
WebMCP Is the New Mobile-First
Browse all figures →
Browse all myths →

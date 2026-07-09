---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48850527"
title: "Show HN: Finterm.ai Finance CLI for Claude Code and Codex"
article_title: ""
author: "cheeseblubber"
captured_at: "2026-07-09T18:46:06Z"
capture_tool: "hn-digest"
hn_id: 48850527
score: 1
comments: 0
posted_at: "2026-07-09T18:39:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Finterm.ai Finance CLI for Claude Code and Codex

- HN: [48850527](https://news.ycombinator.com/item?id=48850527)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T18:39:12Z

## Translation

タイトル: HN を表示: クロード コードおよびコーデックス用の Finterm.ai Finance CLI
HN テキスト: こんにちは、私の名前はカムです。今日、私の共同創設者であるジョシュと私は、Finterm.ai という CLI を出荷します。
コーディング担当者が株価、オプション データ、SEC などの財務データに直接アクセスできるようになります。
提出書類とティッカーの詳細な調査。
私は開発者で、ここ数年はフルタイムのトレーダーです。
最近、私は自分のトレーディングと戦略でLLMをますます使用するようになりました。 LLM が直接アクセスできないことにいつもイライラしていました。
実際の財務情報はウェブ検索に頼らなければならなかったので、それ以上の情報は得られませんでした
特定のオプション価格の詳細な数値。取引を行う際には、株についてできるだけ理解したいと考えています。
アナリストやデータの解釈に頼るのではなく、直接次のことを行うのが好きです。
真実。そのため、貿易論文を作成するときは常に、調査をいくつかの部分に分割します。
調査、アナリストのセンチメント、市場のセンチメント。昨年の9月、私はPopmart Labubuの親会社に関する短い論文を書きました。
私はそのおもちゃは流行なので株は下がるだろうと思った。
私は SEC の提出書類を読み、LLM にも分析してもらいました。
Labubu がどのくらいの規模のドライバーなのか、ビジネス モデルは何なのか、借金はどのようなものなのか、そして掘り下げるべき奇妙に見えるものは何なのか。 EPS と業界指標に関して同社を同業他社と比較しました。私は GPT に、株式に関するあらゆる議論をマッピングするために約 50 のクエリと数百ページを含む詳細な調査を行うよう依頼しました。最後に、オプション データ (コール/プット レシオ、インプライド ボラティリティ、最近の出来高) を調べて、リアルマネーがどのように賭けているかを確認しました。翌月の収益は 16% でした。
しかし、そのフローは苦痛で、SEC データを取得し、ファイリング セクションを GPT にコピーして貼り付け、すべてを手作業で集計し、十数のチャット ウィンドウをやりくりする必要がありました。
過去数か月間、ジョシュ・アンは

d エージェントに取引してもらうためにより多くの時間を費やしました
自主的に。
掘り下げれば掘り下げるほど、解決する必要がある問題がわかってきました
1 つ目は、エージェントがトークン効率の高い方法で事実の情報にアクセスできるようにすることです。私たちの最初の設計上の決定は、Finterm を CLI にすることでした。エージェントが実行したことがわかりました。
CLI を使用すると、MCP や
API 呼び出しを行う。 CLI は自己文書化され、次と同様に動作するように設計されました。
エージェントにとって使いやすいスキルです。次に、複数の呼び出しをまとめてバッチ処理します。
ティッカーを調査するときは、毎回同じいくつかの情報が必要になります - PER
比率、収益、現在の株価、オプションセンチメント。
エージェントが 1 回の通話を行えるようにすることで、トークンを節約し、より完全なビューを提供します。
ティッカーの。ティッカーについて Web 検索を行うと、ノイズの多い記事が頻繁に表示されます (どのくらいの金額が得られるか)
2002 年に Amazon に X ドルを投資していたら作成されました)、SEO スパム、および
同じソースから同じトピックを取り上げた重複した記事。
したがって、Ticker Deep Research は調査パケットを返します。1 件あたり 600 ～ 800 のリンクを取得します。
ティッカー、ノイズである 30 ～ 40% を取り除き、インターネットの状態を返します。
そのティッカー - プライマリまたはセカンダリ、および AI スロップ サイトとラベル付けされたソースを含む重複排除
フラグが立てられました。
エージェントは、何百もの Web ページ自体をクロールする代わりに、徹底的な情報を取得します。
市場がその株についてどう思っているかを示すスナップショット。当社は SEC への申請に対しても同様のアプローチを採用しています。
現在、未加工の提出書類にアクセスできますが、四半期および年次提出書類のほとんどは 90 ～ 95% です。
定型文と繰り返し。
当社は未加工の提出書類だけでなく、エージェントが提出書類のみを閲覧できる SEC 提出書類の差分ツールも提供しています。
diffs: 会社に対する重要な変更。株式とオプションのデータは最大 15 分遅れるため、コストが合理的に保たれ、
に合う

私たちが構築しているリサーチファーストのユースケース。これは株式取引を好む技術者向けのニッチな製品であることを認識しています
クロード コードを使用していますが、それは私自身が感じている多くの不満に近いので、
それを共有して、他の人が興味を持っているかどうかを確認したいと思いました。 finterm.ai でサインアップし、npm install -g @finterm-ai/cli を実行してテストできます。
3 日間の無料トライアルをご利用いただけます (カードが必要です)。フィードバックをお待ちしております。

## Original Extract

Hi, my name is Kam, and today my cofounder Josh and I are shipping Finterm.ai, a CLI that
gives coding agents direct access to financial data: stock prices, options data, SEC
filings, and ticker deep research.
I’m a developer and have been a full-time trader for the past few years.
Recently I have been using LLMs more and more in my trading and strategy. I always found it frustrating that LLMs did not have direct access to
actual financial information and had to rely on web search, so it couldn’t get me more
granular numbers for specific options pricing. When making a trade I want to understand as much as possible about the stock.
Instead of relying on analysts or interpretations of the data, I like to go directly to
the truth. So whenever I have a trade thesis, I break research into a few parts: company
research, analyst sentiment, and market sentiment. Last september I had a short thesis on Popmart Labubu's parent company.
I thought the toy was a fad and that the stock would fall.
I read through the SEC filings and had an LLM analyze them too:
how big a driver is Labubu, what's the business model, what does the debt look like, and what looks strange enough to dig into. I compared the company to its peers on EPS and industry metrics. I asked GPT to do deep research that included around 50 queries and hundreds of pages to map every argument about the stock. Finally I looked at the options data: call/put ratios, implied volatility, recent volume, to see how the real money was betting. I made 16% over the next month.
But the flow was painful, fetching SEC data, copy pasting filing sections in to GPT, aggregating everything by hand, and juggle a dozen chat windows.
In the past few months, Josh and I spent more time trying to get agents to trade
autonomously.
The more we dug in, the more we realized that the problem you need to solve
first is giving agents access to factual information in a token-efficient way. Our first design decision was making Finterm a CLI. We found that agents performed
better with a CLI, since it didn’t waste as many tokens as interfacing with MCP or
making API calls. We designed the CLI to be self-documenting and behave similarly to
skills so it would be agent-friendly. Second, we batch multiple calls together.
Whenever I research a ticker, I want the same few pieces of information every time—P/E
ratio, revenue, current stock price, options sentiment.
We let your agent make a single call, which saves tokens and gives a more complete view
of a ticker. When doing web searches about a ticker, you often get noisy articles (how much you would have
made if you had invested $X in Amazon in 2002), SEO spam, and
duplicated articles covering the same topic from the same source.
So our Ticker Deep Research returns a research packet: it fetches 600–800 links per
ticker, strips out the 30–40% that is noise, and gives back the state of the internet on
that ticker—deduped, with sources labeled primary or secondary and AI-slop sites
flagged.
Instead of crawling hundreds of webpages itself, your agent gets a thorough
snapshot of what the market thinks about the stock. We take the same approach for SEC filings.
Even with raw filings accessible now, most quarterly and annual filings are 90–95%
boilerplate and repetition.
We offer raw filings, but also an SEC filing diff tool where your agent sees only the
diffs: the important changes to the company. Stock and options data is delayed by up to 15 minutes, which keeps costs reasonable and
fits the research-first use case we’re building for. We realize this is a niche product for a technical audience that likes to trade stocks
using Claude Code, but it’s close to a lot of the frustrations I feel myself, so I
wanted to share it and see if anyone else is interested. You can sign up at finterm.ai and npm install -g @finterm-ai/cli to test it.
We have a 3-day free trial (card required) and would love any feedback.


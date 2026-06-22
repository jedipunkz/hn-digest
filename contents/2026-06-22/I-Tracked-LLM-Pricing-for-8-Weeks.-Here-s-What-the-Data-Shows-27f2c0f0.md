---
source: "https://tokenprices.io/blog/llm-pricing-8-weeks"
hn_url: "https://news.ycombinator.com/item?id=48632126"
title: "I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows"
article_title: "I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows. | Token Prices"
author: "emonterrubio"
captured_at: "2026-06-22T16:33:31Z"
capture_tool: "hn-digest"
hn_id: 48632126
score: 2
comments: 0
posted_at: "2026-06-22T16:09:39Z"
tags:
  - hacker-news
  - translated
---

# I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows

- HN: [48632126](https://news.ycombinator.com/item?id=48632126)
- Source: [tokenprices.io](https://tokenprices.io/blog/llm-pricing-8-weeks)
- Score: 2
- Comments: 0
- Posted: 2026-06-22T16:09:39Z

## Translation

Title: I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows
記事のタイトル: LLM の価格を 8 週間追跡しました。データが示すものは次のとおりです。 |トークン価格
説明: 2026 年 4 月 24 日から毎日追跡された、33 のプロバイダーにわたる 1,111 の AI モデルからの実際の価格データ。スプレッド、変動要因、およびそれがインフラストラクチャ料金に与える影響。

記事本文:
LLM の価格を 8 週間追跡しました。データが示すものは次のとおりです。 |トークン価格の機能 価格アラート サインイン分析 LLM の価格を 8 週間追跡しました。データが示すものは次のとおりです。
2026 年 4 月 24 日から毎日収集された、33 のプロバイダーにわたる 1,111 モデルの実際の価格データ。
数か月前、私は AI 製品を構築していましたが、どの LLM を使用すべきかという壁にぶつかりました。
質問は簡単そうに思えます。しかし、AI モデルの価格設定は、警告なしに常に変化します。 OpenAI の価格ページには 69 のモデルがリストされています。 Google は 18. Anthropic で、数週間の間に 3 つの新しいクロード バージョンをリリースしました。オプションを比較して決定を下した時点では、その数字がまだ最新のものであるかどうかは自信がありませんでした。
料金ページを確認してみました。ドキュメントを確認しました。次のようなことを教えてくれた場所は 1 つも見つかりませんでした。
先週と今日のコストはいくらでしたか?
最近実際に価格を変更したプロバイダはどれですか?
同等モデルのスプレッドはどのくらいですか?
そこで自動化してみました。 2026 年 4 月 24 日から、すべての主要プロバイダーから価格データを取得する日次スクレーパーを構築しました。 8 週間のデータが実際に示していることは次のとおりです。
調査結果 1: 大手プロバイダーは動かなかった
見出しの結果は劇的なものではありませんでした。Anthropic、OpenAI、Google、Mistral、xAI はすべて、59 日間を通じて価格を横ばいに保ちました。これら 5 つのプロバイダー全体で価格の変更は 1 つも検出されませんでした。
それ自体は有益な情報です。これらのプロバイダーを中心に予算を立てている場合、4 月のコスト モデルは現在でも正確です。
以下は、6 月 22 日現在の正確な、最も使用されているモデルの現在の価格スナップショットです。
調査結果 2: 価格スプレッドが膨大である
もっと大きな話はボラティリティではありません。これは、1 セント未満の推論からフロンティアのフラッグシップ モデルに至るまで、現在追跡されているカタログ全体に存在する 600 倍の価格帯です。
グロックのラマ 3.1 8

B costs $0.05 per million input tokens. GPT-4 Turbo costs $10.00 (200x). GPT-5.4 Pro costs $30.00 (600x). All three are in our dataset today. DeepSeek V4 Pro は、100 万入力トークンあたり 0.44 ドルで、主力価格の数分の 1 で推論ベンチマークで常にトップ近くにランクされています。
チームが問うべき質問は、単に「どのモデルが最適か?」ということではありません。しかし、「コストを考慮すると、この特定のタスクにはどのモデルが最適ですか?」多くの実稼働ワークロードでは、最初に使用したモデルが答えではありません。
調査結果 3: 5 つの価格変更はすべて 1 つのプロバイダーによるもの
59 日間にわたって、当社のスクレーパーは 1,111 モデルすべてにわたって 5 つの価格変更を検出しました。それらはすべて、サードパーティのオープンソース モデルを実行するホスティング アグリゲーターである Together AI 上にありました。
最大の動きは、Qwen37-max が 5 月 27 日の夜に 50% 下落したことです。主流チャネルに届くような公表はありませんでした。 Together AI を通じてそのモデルを実行しているチームでは、コストが半分に削減されました。 Teams not monitoring pricing had no idea.
6 月 2 日に値上げされたのは、Qwen3.5-9B (+70%)、Llama-3.3-70B-Turbo (+18%)、および Meta-Llama-3-8B-Lite (+40%) の 3 つのモデルです。 Not all pricing movement favors the buyer.
Finding 4: No Provider Notifies You Directly
検出された 5 つの変更すべてにおいて、アナウンス パスは同じであり、ユーザーに直接送信されるものはありませんでした。 1 つの変更 (Together 経由の DeepSeek-V4-Pro) は、数日後にプロバイダーのブログ投稿に掲載されました。残りは黙っていた。
これを他のインフラストラクチャの価格設定と比較してください。
AWS は影響を受ける顧客に価格変更に関する電子メール通知を送信します
Stripe は料金変更の 30 日前に通知します
Twilio は公開変更ログを投稿し、アカウント所有者にメールを送信します
LLM プロバイダーは、エンタープライズ ソフトウェアというよりはスポット マーケットに近い運営を行っています。 Prices move when they move. Most teams find out on their invoice.
数週間後

これを手動で実行しているので、自動化してオープンしました。 Token Price は 33 のプロバイダーにわたる 1,100 以上のモデルを毎日追跡し、変化が起こった瞬間にそれを明らかにします。
1 追跡されているすべてのプロバイダーとモデルにわたるライブ価格ダッシュボード
2 検出されたすべての値動きを前後の価格とともに表示する価格変動フィード
3 2026 年 4 月 24 日に遡る履歴データ (完全なデータセット)
4 REST API により、価格設定を独自の FinOps ツールに取り込むことができます
無料枠は 10 社の主要プロバイダーをカバーしています。有料プランでは、歴史の深さ、より多くのプロバイダー、および API アクセスが追加されます。 tokenprices.io
1.見逃していた価格変更はありますか?このデータに現れていない手を見つけた場合は、知りたいです。
２．次にどのプロバイダーを追加すればよいでしょうか?スクレーパーはより多くのプラットフォームをカバーできます。あなたのリストには何が入っていますか?
３．現在、モデル選択をどのように行っていますか?ダッシュボード？スプレッドシート?直感？
４．このデータをチームにとって実用的なものにするものは何でしょうか?アラート?コストの予測は？比較視点？
support@tokenprices.io までお知らせください。
データ注: すべての数値は、2026 年 4 月 24 日から 6 月 22 日まで毎日収集された、公開されている価格ページに基づいています。プライベート API や交渉された料金は使用されていません。価格には標準のオンデマンド層が反映されています。注記がない限り、すべての金額は 100 万トークンあたりです。
FAQ 利用規約 プライバシー ポリシー AI/LLM プロバイダーのリアルタイムおよび過去の価格。毎日 08:00 UTC に収集されます。
プロバイダーのドキュメントと LiteLLM コミュニティ カタログから入手しました。

## Original Extract

Real pricing data from 1,111 AI models across 33 providers, tracked daily since April 24, 2026. The spread, the movers, and what it means for your infrastructure bill.

I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows. | Token Prices Features Pricing Alerts Sign in Analysis I Tracked LLM Pricing for 8 Weeks. Here's What the Data Shows.
Real pricing data from 1,111 models across 33 providers, collected daily since April 24, 2026.
A few months ago, I was building an AI product and hit a wall: Which LLM should we use?
The question sounds simple. But the pricing landscape for AI models shifts constantly and without warning. OpenAI has 69 models listed on their pricing page. Google has 18. Anthropic launched three new Claude versions in the span of weeks. By the time I had compared options and made a decision, I had no confidence the numbers were still current.
I checked pricing pages. I checked docs. I found no single place that told me:
What did this cost last week vs. today?
Which providers actually changed their prices recently?
How wide is the spread for comparable models?
So I automated it. Starting April 24, 2026, I built a daily scraper that pulls pricing data from every major provider. Here is what 8 weeks of data actually shows.
Finding 1: The Big Providers Did Not Move
The headline finding is not dramatic: Anthropic, OpenAI, Google, Mistral, and xAI all held their prices flat for the entire 59-day period. Not a single price change detected across those five providers.
That is useful information on its own. If you are budgeting around these providers, your cost model from April is still accurate today.
Here is the current pricing snapshot for the most-used models, accurate as of June 22:
Finding 2: The Pricing Spread is Enormous
The bigger story is not volatility. It is the 600x price range that now exists across the tracked catalog, from sub-cent inference to frontier flagship models.
Groq's Llama 3.1 8B costs $0.05 per million input tokens. GPT-4 Turbo costs $10.00 (200x). GPT-5.4 Pro costs $30.00 (600x). All three are in our dataset today. DeepSeek V4 Pro at $0.44 per million input tokens consistently ranks near the top on reasoning benchmarks at a fraction of the flagship price.
The question teams should be asking is not just “which model is best?” but “which model is best for this specific task given the cost?” For many production workloads, the answer is not the model you started with.
Finding 3: All 5 Price Changes Came from One Provider
Over 59 days, our scraper detected 5 pricing changes across all 1,111 models. Every single one was on Together AI , a hosting aggregator that runs third-party open-source models.
The biggest move: Qwen37-max dropped 50% overnight on May 27, with no public announcement that reached mainstream channels. Teams running that model through Together AI saw their costs cut in half. Teams not monitoring pricing had no idea.
Three models saw price increases on June 2: Qwen3.5-9B (+70%), Llama-3.3-70B-Turbo (+18%), and Meta-Llama-3-8B-Lite (+40%). Not all pricing movement favors the buyer.
Finding 4: No Provider Notifies You Directly
Across all 5 detected changes, the announcement path was the same: nothing sent directly to users. One change (DeepSeek-V4-Pro via Together) appeared in a provider blog post days later. The rest were silent.
Compare that to how other infrastructure pricing works:
AWS sends email notifications for pricing changes to affected customers
Stripe gives 30 days notice before any fee changes
Twilio posts a public changelog and emails account owners
LLM providers are operating more like spot markets than enterprise software. Prices move when they move. Most teams find out on their invoice.
After a few weeks of running this manually, I automated it and opened it up. Token Prices tracks 1,100+ models across 33 providers daily and surfaces changes the moment they happen.
1 A live pricing dashboard across all tracked providers and models
2 A price change feed showing every detected move with before/after prices
3 Historical data going back to April 24, 2026 (the full dataset)
4 A REST API so you can pull pricing into your own FinOps tooling
The free tier covers the 10 major providers. Paid plans add historical depth, more providers, and API access. tokenprices.io
1 . Are there pricing changes I missed? If you spotted a move that did not show up in this data, I want to know.
2 . Which providers should I add next? The scraper can cover more platforms. What is on your list?
3 . How do you handle model selection today? Dashboard? Spreadsheet? Gut feeling?
4 . What would make this data actionable for your team? Alerts? Cost projections? A comparison view?
Let us know at support@tokenprices.io
Data note: All figures come from publicly available pricing pages scraped daily from April 24 to June 22, 2026. No private APIs or negotiated rates are used. Prices reflect standard on-demand tiers. All amounts are per 1 million tokens unless noted.
FAQ Terms of Service Privacy Policy Real-time and historical pricing for AI/LLM providers. Collected daily at 08:00 UTC.
Sourced from provider documentation and the LiteLLM community catalog.

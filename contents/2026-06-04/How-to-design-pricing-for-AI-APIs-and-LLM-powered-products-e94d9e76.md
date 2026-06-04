---
source: "https://www.solvimon.com/blog/how-to-design-pricing-for-ai-apis-and-llm-powered-products"
hn_url: "https://news.ycombinator.com/item?id=48400157"
title: "How to design pricing for AI APIs and LLM-powered products"
article_title: "How to design pricing for AI APIs and LLM-powered products | Solvimon Blog"
author: "arnon"
captured_at: "2026-06-04T16:12:07Z"
capture_tool: "hn-digest"
hn_id: 48400157
score: 4
comments: 1
posted_at: "2026-06-04T15:30:25Z"
tags:
  - hacker-news
  - translated
---

# How to design pricing for AI APIs and LLM-powered products

- HN: [48400157](https://news.ycombinator.com/item?id=48400157)
- Source: [www.solvimon.com](https://www.solvimon.com/blog/how-to-design-pricing-for-ai-apis-and-llm-powered-products)
- Score: 4
- Comments: 1
- Posted: 2026-06-04T15:30:25Z

## Translation

タイトル: AI API と LLM を利用した製品の価格を設計する方法
記事のタイトル: AI API と LLM を利用した製品の価格を設計する方法 |ソルビモンのブログ
説明: AI API の価格設定に関するステップバイステップのガイド。順番に 6 つの決定、Claude Opus 4.7 の動作例、サンプル Tier テーブル、および診断を実行するための Claude プロンプト。

記事本文:
ヘッドレス収益化: 1 つのプロンプトまたはコマンドで請求を構成します
AI API と LLM を利用した製品の価格を設計する方法
AI API と LLM を利用した製品の価格を設計する方法
AI API と LLM を利用した製品の価格を設計する方法
このページについて LLM に質問する Perplexity に質問する ChatGPT に質問する クロードに質問する ジェミニに質問する Grok に質問する マークダウンとしてコピーする ニュースレターを購読する Arnon Shimani
AI API の価格設定は、何を計測するか、どのプリミティブに課金するか、ユニットごとに何を課金するか、階層をどのように構成するか、ハード キャップかソフト キャップか、クレジット ウォレットの動作の順に 6 つの決定によって決定されます。このガイドでは、実際の例と金額を示してそれぞれについて説明します。最後にクロード プロンプトがあり、それを貼り付けて独自の価格設定を診断できます。
このブログ記事を書いた理由は、ある創設者と話していて、その価格設定ページを見せてもらったからです。そこには「1,000 トークンあたり 0.02 ドル」と書かれていました…顧客が請求書に何を参照しているのか尋ねましたが…その答えが気に入らなかったからです。したがって、これは、今後 18 か月の損益を悩ませることになる決定を下そうとしている創業者、PM、収益化リーダー向けのものです。
AIの料金プランの構造
従来の API の価格設定では、プランはサブスクリプション層 (無料、プロ、スケール) です。
各プランには、それを構成するフェーズ (14 日間のトライアル、その後のデフォルトのフェーズ) があります。
各フェーズ内には、メーターを価格と権利に結び付ける料金表があります。試用期間中は、「エバーグリーン」段階よりも多くの機能を利用できる場合があります。
AI の価格設定は、この構造に 3 つのひねりを加えます。
メーターはもはや単なる「リクエスト」ではありません。それは、トークン (入力と出力)、クレジット (抽象化された単位)、または結果 (完了したタスク、解決されたチケット、生成されたドキュメント) です。
料金表は、1 人の顧客の下にある複数のモデルのコストを処理する必要があります

-対面価格。
権利はカウンターではなくウォレットです。クレジットは蓄積、期限切れ、ロールオーバーされ、複数の機能にわたって消費されます。
また、AI の価格はより頻繁に変更されるため、従来の SaaS よりもフェーズが重要になります。モデルプロバイダーが価格を値下げしたり、新しいクレジットレートを設定したり、顧客が試用期間を延長したりすることにより、必ずしも単なる契約の修正ではなく、フェーズ移行が発生します。
これにより、順番に 6 つの決定が得られます。
ステップ 1: 測定対象を選択する
メーターはカウントするものです。選択を誤ると、今後 2 年間、その請求が公正であるかどうかについて顧客と争うことになります。
AI 製品の場合、4 メートルが優勢です。
入力 + 出力トークン (OpenAI、Anthropic、Mistral)
クレジット (カーソル、Notion AI、Linear AI)
結果 (解決された会話ごとに 0.99 ドルのインターコムフィン)
メーターは、顧客が認識する価値と商品のコストという 2 つのことを同時に相関させる必要があります。マルチモデル製品のフラットなリクエストごとのメーターは、顧客が 50,000 トークンのプロンプトを送信した瞬間に、(すでに薄い) 粗利益に大打撃を与えます。
💡 ロックするかどうかの決定: どのメーターを正確に定義するか。
それを書き留めます - 「1 トークン = 1 GPT-5.5 スタイルのトークン、入力と出力の組み合わせ」。私の経験では、曖昧な定義は顧客との紛争の原因となり、それは非常に簡単に回避できます。
ステップ 2: 価格設定のプリミティブを選択する
社内で別の部門で計算を行う場合でも、顧客対応部門を提案します。
$/1,000 のレートでトークンを入力 + 出力
直接 API 製品、テクニカルバイヤー、OpenAI スタイルのリセラー
マルチモデルの使用、生産量が大幅に変化する場合、技術的知識のない購入者
あらゆる機能に対して引き換え可能な抽象化されたユニット
LLM を利用した SaaS、混合モデルの使用、パッケージ化された AI 機能
クレジットレートがコストの変化を追跡しない場合、顧客がバーンレートを予測できない場合
属ごと

ted ドキュメント、解決されたチケット、完了したタスク
明確な価値単位、エンタープライズ契約、垂直 AI
結果の定義があいまいな場合、障害モードが織り込まれていない場合
価格変更でうまくいったほとんどの製品は、最終的には 3 つすべてを同時に実行することになります。
たとえば、API のパワー ユーザー向けのトークンごと、パッケージ化された製品のクレジット、および企業取引の結果価格の最上位などです。
Agentforce は会話に対して課金し、Fin は結果ベースで実行されますが、カーソルはクレジットを実行し、OpenAI はトークンを実行します。すべて有効ですが、ビジネスに大きく依存します。
繰り返しになりますが、定義できる場合は結果が第一、柔軟性を求める場合はクレジットが二番目、回避できない場合はトークンが持続します。
💡 ロックする決定: どのプリミティブが請求書に表示されるか。
ステップ 3: コストを計算し、価格を設定する
以下に、作業に適した例を示します。製品は、ユーザーのリクエストごとに 1 つの Claude Opus 4.8 呼び出しを行います。 Anthropic は、100 万の入力トークンあたり 5 ドル、100 万の出力トークンあたり 25 ドルを請求します。
平均的なリクエストで 1,000 個の入力トークンと 500 個の出力トークンが使用される場合、リクエストあたりのコストは $0.005 + $0.0125 = $0.0175 となります。
そのコストから考慮すべき 3 つの価格設定アプローチ:
コストプラスのマークアップ。リクエストごとに $0.035 を請求します。粗利益率 50%。予測可能、防御可能、退屈。 API リセラー カテゴリで活動します。コスト プロファイルが異なる 2 番目のモデルを追加した瞬間に問題が発生します。
価値観に基づく。顧客は、生成されたドキュメントごとに 1 ドルを支払います。それが顧客にとっての価値だからです。 0.50 ドルを請求します。 ~96.5%の粗利益率。競合他社に負けるのは危険です。お堀があると最高です。 (Salesforce の古典的な手法。AI では堀が薄くなったため、現在は異なります。)
クレジットベースの抽象化。クレジットを「1 つの標準リクエスト」として定義し、1 クレジットあたり 0.05 ドルを請求します。大量の入力リクエストには 2 クレジットがかかります。ドキュメントの生成には 5 クレジットがかかります。顧客が見る

トークン数ではなく残高です。クレジットとコストのマッピングを内部で制御するため、モデルが変更されてもマージンはスムーズに移動します。
補足: Opus 4.7 では、以前の Opus モデルと比較して、同じ入力テキストに対して最大 35% 多くのトークンを使用できる新しいトークナイザーが導入されました。同じプロンプト、同じモデル ファミリでは、請求額が 35% 高くなる可能性があります。価格設定がトークンごとの場合は、簡単には回避できない価格上昇を吸収したことになりますが、価格設定がクレジットごとの場合は、クレジットとトークンの比率を内部で調整するため、顧客は何も気づきません。建築としてのクレジット、再び。
私にとっての実際的なルールは、顧客の毎月の請求額を ±20% 以内で予測できない場合、価格設定がコストと密接に関係しすぎているということです。制御されたレートのクレジットはそれをある程度解決しますが、トークンごとのコストプラスでは間違いなく解決しません。
💡 ロックの決定: 単位あたりの価格 (選択したプリミティブで) と目標粗利益。
ステップ 4: 階層構造を構築する
ほとんどの AI 製品は 4 ～ 5 層に分類されます。数字よりも構造が重要です。数字は変更できます。形を簡単に変えることはできません。
検討の開始点は次のとおりです。
ソフトキャップ、クレジット超過あたり 0.01 ドル
ソフトキャップ、クレジット超過あたり 0.005 ドル
私は、無料のハードキャップ (マージンの悪用を防ぐ)、または有料の超過を含むソフトキャップ (顧客がスケールできるようにする) と、トップのコミットベースを好みます。
無料利用枠は通常、取得コストであることに注意してください。そのため、悪用の境界線を考慮して価格を設定してください。100 クレジットで開発者は「Hello World」を利用でき、10,000 クレジットでは趣味で本物のものを構築できるようになります。それを計画している場合は問題ありませんが、サインアップあたりのコストが CAC に大きな損害を与え始める場合は、悪用ラインが存在します。
また、超過料金として含まれている料金の 2 ～ 3 倍の料金が標準で請求されることにも注意してください。 3倍を超える

顧客が自分の使用量を社内で厳しく制限することになり、拡張収益が抑制され、製品が妨害される可能性があります（顧客からお金を奪っているように感じます）。 1.5 倍を下回ると、パワーユーザーのマージンが確実に漏れてしまいます。
💡 ロックの決定: 階層数、階層ごとに含まれる使用量、階層ごとの価格、階層ごとの超過率。
ステップ 5: ハードキャップとソフトキャップを決定する
権利がなくなると、ハード制限によってリクエストがブロックされます。顧客はエラー (HTTP 429 を考慮してください!) を受け取り、アップグレードするか期間が切り替わるまで停止します。
ソフト制限を使用すると、使用資格を超えて使用を継続し、期間の境界で超過分を請求できます。
顧客が超過条項に同意していないもの
開発者サンドボックスとテスト環境
契約に超過条件が含まれる有料枠
ブロッキングによって実稼働システムが破壊される可能性があるものすべて
使用量を最大化し、クレジット限度額を設定したい大規模およびエンタープライズのお客様
私がよく見かける人々の間違いは、「顧客を驚かせてアプリの動作を停止させたくない」という理由で、有料枠に厳しい制限を設けることです。これは立派なことですが、月末に莫大な請求を受けるのも良くありません。寛大な支出アラートを備えたソフトリミットがその動きです。
💡 ロックの決定: 階層ごとのハードまたはソフト、ソフトの超過料金、アラートしきい値 (権利の 50%、80%、100%)。
ステップ 6: クレジットウォレットを設計する
プリミティブがクレジット (ステップ 2) の場合、ウォレットが価格設定エクスペリエンス全体となります。 4 つの決定を間違えると、顧客はあなたを容認しますが、決して信頼しません。
クレジットには有効期限がありますか?月末は厳しく、年末は寛大、契約終了は企業のデフォルトです。
有効期限がないことは、貸借対照表上でかなり深刻な負債を生み出すことになります。そのため、「クレジットは無期限です」というマーケティングラインを出荷する前に、CFO に相談してください。
何が起こるのですか

期間の境界で未使用のクレジットを削除しますか?完全なロールオーバーが最もユーザーフレンドリーですが、部分的なロールオーバーを行うものもあります。
ロールオーバーがないのは最悪です。顧客は最終的にそれに気づき、再認識するでしょう。
顧客はサイクルの途中でさらに購入できるでしょうか?
同じ価格ですか、それともプレミアムですか？
チャージには顧客の幸福が宿ります。ここでの摩擦が、「製品は気に入っているが、請求が煩わしい」という最大の原因です。
1 つのクレジット プールを複数の機能 (チャット、検索、生成) にわたって使用できますか?それとも、各機能に独自のプールがあるのでしょうか?
単一プールの方が親しみやすく簡単ですが、ASC 606 に基づく収益認識はマルチプールの方が簡単です。顧客が要求したため、最終的には単一プールを使用することになります。
💡 ロックの決定: 有効期限ポリシー、ロールオーバー ポリシー、トップアップ フロー、シングル プールとマルチ プール。
モデルの価格が下がったときの計画
Anthropic は、Opus を 100 万トークンあたりのインプット $15 / アウトプット $75 (Opus 4.1) から $5 / $25 (Opus 4.5 以降) に引き上げました。ラインナップの中で最もプレミアムなモデルを3倍カット。
OpenAI は四半期ごとに同様の動きを出荷していますが、そうですね、これについて計画を立てるのは難しいです。
固定マークアップでトークンごとの価格を設定した場合、粗利益はちょうど 3 倍になります。損益にとっては良いことですが、請求額が高すぎると感じている顧客にとっては良くありません。
60 日以内に再交渉が要求されるか、自動的に価格を再設定するベンダーに切り替えることが予想されます。
クレジットベースで価格を設定している場合は、選択肢があります。信用レートを維持し、証拠金を銀行に預けます。それを下げて貯蓄をスルーしてください。または、バランスを再調整します。古いモデルのレートを下げ、新しいモデルのレートを維持します。
結果に基づいて価格を設定した場合は、ほとんど気付かないでしょう。 Opus 4.7 を使用した場合でも Haiku 4.5 を使用した場合でも、顧客は解像度ごとに 0.99 ドルを支払います。マージンは自動的に複合されます。 (結果価格設定が最も耐久性のある形状であり、初日に出荷するのが最も難しい理由。)
知っておくべきインフラストラクチャの問題は、次のことが可能かということです

当社の請求システムは、以前の請求書を書き換えることなく、サイクルの途中でクレジット レートを変更しますか?そうしないと、モデルの価格が下がるたびにプロジェクトや移行が発生し、大きな損害を被る可能性があります。
クロードと一緒に価格を診断してみよう
このプロンプトを Claude (または任意の対応モデル) にコピーします。ユーザーとともに検出を実行し、階層テーブルと推論を返します。正確に価格を設定することはできません。それはあなたの判断であり、あなたのデータです。正確な形状が得られ、60 日間の設計作業が節約されます。
AI API または LLM を利用した製品の価格設定を設計するのに助けが必要です。 2 段階に分けてご協力ください。
# フェーズ 1 — これらの質問を順番に聞いてください。次に進む前に、各 (または各小グループ) に対する私の回答を待ってください。私から十分な協力を得られるまで、フェーズ 2 に進まないでください。
あなたの製品は何をしますか? (どのような問題を解決しますか?誰が使用しますか?)
Which models do you use? (単一モデル、マルチモデル、フォールバックあり?)
平均的なリクエストはどのようなものですか? (おおよその入力トークン、出力トークン、レイテンシ バジェット、マルチモーダル コスト。)
リクエストごとのコストはいくらですか? (または、マルチモデルの場合はトークンごと、モデルごと。)
顧客にどの単位で支払ってもらいたいですか? (トークン、クレジット、完了した結果、ハイブリッド。)
あなたのターゲット顧客は誰ですか

[切り捨てられた]

## Original Extract

A step-by-step guide to AI API pricing. Six decisions in order, a Claude Opus 4.7 worked example, a sample tier table, and a Claude prompt to run the diagnosis with you.

Headless Monetization: configure billing with one prompt or command
How to design pricing for AI APIs and LLM-powered products
How to design pricing for AI APIs and LLM-powered products
How to design pricing for AI APIs and LLM-powered products
Ask an LLM about this page Ask Perplexity Ask ChatGPT Ask Claude Ask Gemini Ask Grok Copy as Markdown Subscribe to our newsletter Arnon Shimoni
AI API pricing comes down to six decisions, in order: what to meter, which primitive to charge for, what to charge per unit, how to structure tiers, hard cap or soft cap, and how the credit wallet behaves. This guide walks through each one with worked examples and dollar amounts. There's a Claude prompt at the end you can paste in to diagnose your own pricing.
The reason for writing this blog post is because I was talking to a founder who showed me his pricing page pricing page, that said "$0.02 per 1k tokens"… I asked what the customer sees on the invoice, and… I didnøt like the answer. So, this is for founders, PMs, and monetization leads about to make the decisions that will haunt the next 18 months of their P&L.
The structure of an AI pricing plan
With classic API pricing, a plan is the subscription tier (Free, Pro, Scale).
Inside each plan there are phases (a 14-day trial, then the default phase) that make it up.
Inside each phase there are rate cards that tie a meter to a price and an entitlement. During a trial, you may get more features than during the "evergreen" phase.
AI pricing adds three twists to this structure.
The meter is no longer just "requests." It's tokens (input and output), credits (abstracted units), or outcomes (a completed task, a resolved ticket, a generated document).
The rate card has to handle multiple model costs sitting under one customer-facing price.
The entitlement is a wallet, not a counter. Credits accumulate, expire, roll over, and get spent across multiple features.
Also, phases matter more than they did in classic SaaS, because AI pricing changes more often. A model provider cuts their price or you ship a new credit rate, or a customer extends their trial - all of those creates a phase transition, not necessarily just a contract amendment.
That gives you six decisions, in order.
Step 1: Pick what you're metering
The meter is the thing you count. Pick it wrong and you'll fight your customers about whether the bill is fair for the next two years.
For AI products, four meters dominate:
Input + output tokens (OpenAI, Anthropic, Mistral)
Credits (Cursor, Notion AI, Linear AI)
Outcomes (Intercom Fin at $0.99 per resolved conversation)
The meter has to correlate with two things at once: customer-perceived value, and your cost of goods. A flat per-request meter on a multi-model product will wreak havoc on your (already thin) gross margin the moment a customer sends 50k-token prompts.
💡 Decision to lock: which meter, with the exact definition.
Write it down - "1 token = 1 GPT-5.5-style token, input + output combined". My exprience is that vague definitions are a source of customer disputes that you can very easily avoid.
Step 2: Pick your pricing primitive
I suggest a customer-facing unit, even if you do the math in another internally.
Input + output tokens at a $/1k rate
Direct API products, technical buyers, OpenAI-style resellers
Multi-model usage, when output volume varies wildly, non-technical buyers
Abstracted units redeemable against any feature
LLM-powered SaaS, mixed model usage, packaged AI features
When credit rates don't track cost changes, when customers can't predict burn rate
Per generated document, resolved ticket, completed task
Clear value units, enterprise contracts, vertical AI
When outcome definitions are fuzzy, when failure modes aren't priced in
Most products that do well in a pricing change end up running all three at once.
For example, a per-token for power users on the API, with credits for the packaged product, and outcome pricing on top for enterprise deals.
Agentforce charges for conversation, Fin runs outcome-based, but cursor runs credits and OpenAI runs tokens. All are valid, but highly dependent on the business.
So again, outcomes first if you can define them, credits second for flexibility, tokens last if you can't avoid them.
💡 Decision to lock: which primitive appears on the invoice.
Step 3: Calculate your cost and set your price
Here's a good example to work with: Your product makes one Claude Opus 4.8 call per user request. Anthropic charges $5 per 1M input tokens and $25 per 1M output tokens.
If your average request uses 1k input tokens and 500 output tokens, your cost per request is $0.005 + $0.0125 = $0.0175 .
Three pricing approaches from that cost to consider:
Cost-plus markup. Charge $0.035 per request. 50% gross margin. Predictable, defensible, boring. Works in the API reseller category. Breaks the moment you add a second model with a different cost profile.
Value-based. The customer would pay $1 per generated document because that's the value to them. Charge $0.50. ~96.5% gross margin. Risky when a competitor undercuts you. Excellent when you have a moat. (The classic Salesforce play. Different now in AI because the moat is thinner.)
Credit-based abstraction. Define a credit as "one standard request" and charge $0.05 per credit. Heavy-input requests cost 2 credits. Document generation costs 5 credits. The customer sees a balance, not a token count. Your margin moves smoothly across model changes because you control the credit-to-cost mapping internally.
Side-note: Opus 4.7 introduced a new tokenizer that can use up to 35% more tokens for the same input text compared to earlier Opus models. Same prompt, same model family can result in a 35% bigger bill. If your pricing is per-token, you just absorbed a price hike you can't pass through easily but if your pricing is per-credit, you adjust the credit-to-token ratio internally and the customer notices nothing. Credits as architecture, again .
A practical rule for me is that if you can't predict the customer's monthly bill within ±20%, your pricing is too tightly coupled to your cost. Credits with controlled rates kinda solve it, while cost-plus per-token definitely doesn't.
💡 Decision to lock: per-unit price (in whatever primitive you picked) and target gross margin.
Step 4: Build your tier structure
Most AI products land on 4-5 tiers. The structure matters more than the numbers. You can change the numbers. You can't easily change the shape.
Here's a starting point to consider:
Soft cap, $0.01/credit overage
Soft cap, $0.005/credit overage
I like a hard cap on free (protects margin from abuse), or a soft cap with overages on paid (let customers scale), and commit-based for the top.
You should note, that a free tier is typically an acquisition cost - so price it for the abuse boundary… 100 credits gets a developer through a "hello world" and 10,000 credits can get a hobbyist building real things on your dime. That's fine if you're planning for it, but the abuse line is wherever the cost per signup makes your CAC start to really hurt.
Also worth noting that charging 2-3x the included rate as overage is standard. Going higher than 3x makes customers hard-cap their own usage internally which suppresses your expansion revenue and may sabotage your product (it feels like you're robbing them). Going lower than 1.5x definitely leaks margin on power users.
💡 Decision to lock: tier count, included usage per tier, price per tier, overage rate per tier.
Step 5: Decide hard cap vs soft cap
Hard limits block requests when the entitlement is exhausted. Customers get an error (consider an HTTP 429 !) and stop until they upgrade or the period rolls over.
Soft limits let usage continue past the entitlement and bill the overage at the period boundary.
Anything where the customer hasn't agreed to overage terms
Developer sandboxes and test environments
Paid tiers where overage terms are in the contract
Anything where blocking would break a production system
Scale and Enterprise customers who want to maximise usage and have credit lines
The common mistake I see people make is putting hard limits on paid tiers because "we don't want to surprise the customer and have their app stop working" - which is noble but getting a huge huge bill at the end of the month isn't great either. A soft limit with a generous spend alert is the move.
💡 Decision to lock: hard or soft per tier, the overage price for soft, the alert thresholds (50%, 80%, 100% of entitlement).
Step 6: Design your credit wallet
If your primitive is credits (Step 2), the wallet is the entire pricing experience. Get the four decisions wrong and customers will tolerate you but never trust you.
Do credits expire? Month-end is harsh, year-end is generous, contract-end is the enterprise default.
Having no expiry creates a pretty serious liability on your balance sheet - so talk to a CFO before you ship "credits never expire" as a marketing line.
What happens to unused credits at period boundaries? Full rollover is most user-friendly, but some do a partial rollover.
Having no rollover sucks. Customers will eventually notice and resent it.
Can customers buy more mid-cycle?
At the same price, or a premium?
Top-ups are where customer happiness lives. Friction here is the single biggest cause of "we love the product but the billing is annoying."
Can one credit pool be spent across multiple features (chat, search, generation)? Or does each feature have its own pool?
A single pool is friendlier and easier, but a multi-pool is easier to revenue-recognize under ASC 606. You'll end up with single pool because customers ask for it.
💡 Decision to lock: expiry policy, rollover policy, top-up flow, single vs multi pool.
Planning for when the model prices drop
Anthropic took Opus from $15 input / $75 output per 1M tokens (Opus 4.1) to $5 / $25 (Opus 4.5 and later). A 3x cut on the most premium model in the lineup.
OpenAI ships similar moves every quarter, and… Yeah, well, it's hard to plan around this.
If you priced per-token at a fixed markup, your gross margin just tripled. Good for the P&L, but not great for the customer who now feels the bill is too high.
Expect a renegotiation request within 60 days, or a switch to a vendor who reprices automatically.
If you priced credit-based, you have a choice. Hold the credit rate and bank the margin. Lower it and pass savings through. Or rebalance: lower the rate for older models, hold for newer ones.
If you priced outcome-based, you barely notice. The customer pays $0.99 per resolution whether you used Opus 4.7 or Haiku 4.5. Margin compounds automatically. (Why outcome pricing is the most durable shape, and the hardest to ship on day one.)
The infrastructure question you must know is can your billing system change a credit rate mid-cycle without rewriting prior invoices? If not, every model price drop becomes a project and a migration which can really hurt.
Diagnose your pricing with Claude
Copy this prompt into Claude (or any capable model). It runs the discovery with you and returns a tier table and reasoning. It won't get the price points exactly right. That's your judgment and your data. It does get the shape right, and it saves 60 days of design work.
I need help designing pricing for my AI API or LLM - powered product . Please work with me in two phases .
# Phase 1 — Ask me these questions in turn . Wait for my answer to each ( or each small group ) before moving on . Don 't move to Phase 2 until you have enough from me to work with.
What does your product do ? ( What problem does it solve ? Who uses it ? )
Which models do you use? ( Single model , multi - model , with fallback ? )
What does an average request look like ? ( Approximate input tokens , output tokens , latency budget , any multimodal cost. )
What 's your cost per request? (Or per token, by model, if multi-model.)
What unit do you want customers to pay for ? ( Tokens , credits , completed outcomes , hybrid . )
Who are your target customer

[truncated]

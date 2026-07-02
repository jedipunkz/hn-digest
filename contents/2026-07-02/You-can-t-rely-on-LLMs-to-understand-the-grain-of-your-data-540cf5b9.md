---
source: "http://www.bioinformaticszen.com/post/llms-dont-understand-the-grain-of-your-data/"
hn_url: "https://news.ycombinator.com/item?id=48763185"
title: "You can't rely on LLMs to understand the grain of your data"
article_title: "Bioinformatics Zen | You can’t rely on LLMs to understand the grain of your data"
author: "michaelbarton"
captured_at: "2026-07-02T16:00:38Z"
capture_tool: "hn-digest"
hn_id: 48763185
score: 1
comments: 0
posted_at: "2026-07-02T15:35:35Z"
tags:
  - hacker-news
  - translated
---

# You can't rely on LLMs to understand the grain of your data

- HN: [48763185](https://news.ycombinator.com/item?id=48763185)
- Source: [www.bioinformaticszen.com](http://www.bioinformaticszen.com/post/llms-dont-understand-the-grain-of-your-data/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T15:35:35Z

## Translation

タイトル: データの粒度を理解するために LLM に依存することはできません
記事タイトル: バイオインフォマティクス禅 |データの粒度を理解するために LLM に依存することはできません

記事本文:
バイオインフォマティクス 禅 |データの粒度を理解するために LLM に依存することはできません
ブルースカイ
フィード
について
バイオインフォマティクス禅
Michael Barton によるバイオインフォマティクスとマインドフルネスに関するブログ。
データの粒度を理解するために LLM に依存することはできません
LLM を使用して分析ノートブックの図または表を生成するときは、その結果をもとに進めたいだけです。 LLM の宿題の採点に時間を費やしたくないです。 LLM の分析をチェックしても、感情的にも同僚からも見返りはありません。分析グラフや数値は、私が生成することで報酬を得ているものであり、LLM では、それらをより速く生成できることが期待されています。 2025 年 11 月頃に LLM が大幅に改善される前から、より低速で思慮深い分析よりも、より高速で応答性の高い分析のほうがメリットがあることを私は知っていました。私がこの業界にいる限り、それはこの業界に当てはまります。 LLM が結果を生み出す速度とそれに伴う期待が、これをさらに悪化させています。
LLM が登場する前は、分析を自分で書かなければならないという摩擦が、一種の身を守る手段だったことに気づきました。すべてを入力しなければならないため、意識的に努力することなく、無意識のうちに自分が何をしているのかを考えるようになりました。バックグラウンドで実行されているプロセスでした。 LLM を使用すると、その背景的な考え方はなくなります。私は意識的にそうしなければなりませんが、LLM からそうしないよう促されています。 LLM が十分に優れるようになる前でさえ、私は通常、将来のことを考えるためにもっと時間を費やす必要があると感じていました。
ここでは、公開されている実際の e コマース データ (Olist、ブラジルの e コマース スタートアップからの 10 万件の注文) を使用した具体的な例を示します。私はLLMにこのデータを渡して次のように尋ねました。
製品カテゴリおよび支払いタイプごとに収益を合計する SQL クエリの例を示します。
これはアナリストが尋ねるかもしれない種類の質問です。カテゴリは or のアイテム側から取得されます。

der (product に結合された order_items)。支払いタイプは支払いテーブルにあります。どちらも order_id にキー設定されています。自然な SQL はそれらを直接結合します。 LLM は、結合によって間違った粒度で誤って追加の行が作成されるという、潜在的なファンアウトの問題を特定しました。ここでは、データの粒度、つまり各行が何を表すかが重要です。注文テーブルには注文ごとに 1 行があります。支払いテーブルには支払い料金ごとに 1 行があり、クレジット カードで支払われた注文と 2 枚の伝票の場合は 3 行になります。それを単純に複数アイテムの注文に結合すると、支払い額が倍増して収益が膨らみます。 LLM はこれを発見しましたが、その修正によりさらに微妙な問題が発生しました。認知的危険性は、LLM が粒子の問題を特定したため、生成された分析には他に問題がないという私の無意識の偏見です。特にLLMが非常に前向きな言葉でそれを販売する場合。ただし、この分析を自分で確認すると、注文の 2.3% に影響を与える 2 番目の微妙なバグが見つかるでしょう。おそらくインパクトはありませんが、間違っており、自信を持って提示されています。
これに対する簡単な答えはありませんが、より適切なプロンプトを提供することで解決できるとは思いません。
分析の質問が非常に単純でない限り、私は常に計画モードから始めます。計画モードでは、計画を読むという作業に少し便利な手間が加わります。読んでいる間、私は考えているだけで、何もしていません。他人のコードをレビューしたことのある人なら、視点が異なることを知っているでしょう。計画モードの使用は、この変化の良いきっかけとなります。具体的には、計画モードを使用する場合、Claude exit フックがあり、LLM に次のような質問をします。
終了基準は何ですか?成功とはどのようなものですか?
データの不変条件と粒度は何ですか?
潜在的な故障モードは何ですか?
これらの質問はすべて、私や LLM の偏見との戦いに関するものです。

カーネマン氏はこれをシステム 1 対システム 2 と呼ぶかもしれません。LLM は迅速に考えたいと考えています。ゆっくり考えるようにしています。これらの質問を含めると、ほとんどの場合、LLM の考え方にギャップがあることがわかります。この習慣を採用する前に生成したコードについて心配になるほどで​​す。これらの質問は、私自身、または LLM が行っていたことを知らなかった仮定を明らかにすることがよくあります。
この記事の冒頭に戻ると、LLM と協力することは、限られたリソースとしての精神的な努力を管理することになります。 LLM の動作をチェックするのが難しくなればなるほど、私がチェックする可能性は低くなります。そして、LLM の生成が速くなるほど、難しくなります。
この記事で私が主張している点を理解するために、この付録を読む必要はありません。データ粒度の誤解が原因で微妙なバグがどのように発生するかを詳しく知りたい場合は、このセクションを参照してください。まさにこのタイプの問題は、現世代のフロンティア LLM が依然として苦労しているようです。私は Fable を試す機会がありませんでしたが、おそらく誰かが上記の元のプロンプトを使用して、そのパフォーマンスを確認できるでしょう。
order_items を order_id の支払いに結合すると、すべての商品行 × すべての支払い行という多対多のファンアウトが生成されます。 LLM は、最初に提案されたクエリでこれを即座に識別します。
-- 注: このクエリは間違っています
選択
pr.product_category_name AS カテゴリ 、
p.payment_type、
COUNT ( DISTINCT oi.order_id) AS order_count、
ROUND ( SUM (p.payment_value), 2 ) AS total_revenue
order_items から
製品を結合 pr ON oi.product_id = pr.product_id
JOIN 支払い p ON oi.order_id = p.order_id
GROUP BY pr.product_category_name、p.payment_type
ORDER BY total_revenue DESC ;
LLM はファンアウトにフラグを立て、アイテム テーブルからプライマリ支払い方法にフィルタリングされた、price +freight_value に切り替えることを提案しました。
-- 注: このクエリは

まだ間違っていますが、より微妙に間違っています
選択
pr.product_category_name AS カテゴリ 、
p.payment_type、
COUNT ( DISTINCT oi.order_id) AS order_count、
-- ここで、LLM は、freight_value を含むように総収益を更新しました。
ROUND ( SUM (oi.price + oi.freight_value), 2 ) AS total_revenue
order_items から
製品を結合 pr ON oi.product_id = pr.product_id
JOIN 支払い p ON oi.order_id = p.order_id
-- 主要な支払い方法のみにフィルターします
WHERE p.payment_sequential = 1
GROUP BY pr.product_category_name、p.payment_type
ORDER BY total_revenue DESC ;
これによりインフレは回避されますが、支払いタイプに分割された注文の収益が誤って割り当てられます。これがどのように問題を起こすかを示す具体的な例を次に示します。
注文: 0016dfedd97fc2950e388d2971d718c7
ただし、2 番目のクエリでは、すべての商品の収益がpayment_sequential = 1 に関連付けられているため、payment_sequential = 2 である伝票は内訳から消えます。
分割支払いタイプの注文: 2,246 (全注文の 2.3%)
バウチャー収益 (LLM 修正): R$ 162,091.39 (合計の 1.0%)
バウチャー収益 (正): R$ 352,707.71 (総額の 2.2%)
不正帰属: R$ 190,616.32
この修正には、3 番目の静かな問題があります。一部の注文には、payment_sequential = 1 行がまったくないため、WHERE 句によって完全に削除されます。
Payment_sequential のない注文 = 1:80
おそらく正しいアプローチでは、各注文の支払いを商品価格、伝票、および場合によっては他の連続支払いタイプに基づいて比例的にカテゴリ間で割り当てることになります。ここではそのクエリを追加していませんが、さらに複雑になっていきます。まだここを読んでいるなら、LLM が単純なプロンプトで誤った判断をするという点はすでに理解していると思います。

## Original Extract

Bioinformatics Zen | You can’t rely on LLMs to understand the grain of your data
bluesky
feed
about
Bioinformatics Zen
A blog about bioinformatics and mindfulness by Michael Barton.
You can’t rely on LLMs to understand the grain of your data
When I generate an analysis notebook figure or table with an LLM I just want to move forward with the result. I don’t want to spend time grading the LLM’s homework. There’s no reward for checking the analysis of an LLM — emotionally or from my colleagues. Analysis graphs and figures are what I get paid to generate, and the expectation with LLMs is that I generate them faster. Even before LLMs became much better around November 2025, I knew that faster, more responsive analysis gets rewarded over slower, more thoughtful analysis. That’s been true of the industry as long as I’ve been in it. The speed at which LLMs can generate results, and the accompanying expectations, have just compounded this.
I’ve noticed that the pre-LLM friction of having to write analysis myself was a form of protection. Having to type everything caused me to subconsciously think about what I was doing, without conscious effort. It was a process running in the background. Now with LLMs that background thinking is gone. I have to do it consciously, and I’m being nudged by the LLM not to. Even before LLMs became good enough, I usually felt that I should spend more time thinking ahead.
Here’s a concrete example using publicly available real e-commerce data ( Olist , 100K orders from a Brazilian e-commerce startup). I gave an LLM this data and asked:
Show an example SQL query that would total revenue by product category and payment type
This is the kind of question an analyst might ask. Category comes from the items side of an order ( order_items joined to products ). Payment type lives in the payments table. Both are keyed to order_id . The natural SQL joins them directly. The LLM identified a potential fan-out issue: where the join would incorrectly create additional rows at the wrong grain. The grain of the data, what each row represents, matters here. The orders table has one row per order. The payments table has one row per payment charge — an order paid with a credit card and two vouchers has three rows. Join that to a multi-item order naively and the payment values multiply, inflating the revenue. The LLM caught this, but its fix introduced a subtler problem. The cognitive danger is that the LLM identified a grain issue, and therefore my unconscious bias is that there are no other issues in the generated analysis. Especially when the LLM sells it in very positive language. However, if I checked this analysis myself I would find a second subtle bug that affects 2.3% of orders. Not impactful perhaps, but incorrect — and presented confidently.
I don’t have any easy answers for this, but I don’t think it can be solved by better prompting.
Unless the analysis ask is very simple, I always start with planning mode. Planning mode adds a little helpful friction back in: reading the plan. While I’m reading, I’m thinking and not doing. Anyone who’s ever reviewed someone else’s code knows that the perspective is different. Using planning mode is a good trigger for this shift. Specifically, when I use planning mode I have a Claude exit hook , which asks the LLM questions like:
What are the exit criteria, what does success look like?
What are the invariants and grain of the data?
What are the potential failure modes?
All of these questions are about fighting bias, mine and the LLM’s. Kahneman might call this System 1 versus System 2. The LLM wants to think fast. I’m trying to think slow. Including these questions almost always shows a gap in the LLM’s thinking — to the extent it makes me worried about code I generated before adopting this habit. These questions often surface assumptions that I didn’t know I, or the LLM, was making.
Going back to the beginning of this article — working with LLMs has become about managing my mental effort as a limited resource. The harder it is to check the work of the LLM, the less likely I am to do it. And the faster the LLM generates, the harder it gets.
You don’t need to read this appendix to understand the point I’m making in this article. This section is here if you want to get into the details about how subtle bugs can show up due to data grain misunderstanding. Exactly the type of issue that the current generation of frontier LLMs still seem to struggle with. I didn’t get a chance to try with Fable, but perhaps someone could try using the original prompt above and see how it performs.
Joining order_items to payments on order_id produces a many-to-many fan-out: every item row × every payment row. The LLM identifies this immediately in the first proposed query.
-- Note: this query is wrong
SELECT
pr.product_category_name AS category ,
p.payment_type,
COUNT ( DISTINCT oi.order_id) AS order_count,
ROUND ( SUM (p.payment_value), 2 ) AS total_revenue
FROM order_items oi
JOIN products pr ON oi.product_id = pr.product_id
JOIN payments p ON oi.order_id = p.order_id
GROUP BY pr.product_category_name, p.payment_type
ORDER BY total_revenue DESC ;
The LLM flagged the fan-out and proposed switching to price + freight_value from the items table, filtered to the primary payment method:
-- Note: this query is still wrong, but more subtly so
SELECT
pr.product_category_name AS category ,
p.payment_type,
COUNT ( DISTINCT oi.order_id) AS order_count,
-- here the LLM has updated total revenue to include freight_value
ROUND ( SUM (oi.price + oi.freight_value), 2 ) AS total_revenue
FROM order_items oi
JOIN products pr ON oi.product_id = pr.product_id
JOIN payments p ON oi.order_id = p.order_id
-- filter to primary payment method only
WHERE p.payment_sequential = 1
GROUP BY pr.product_category_name, p.payment_type
ORDER BY total_revenue DESC ;
This avoids the inflation but misattributes revenue for orders split across payment types. Here’s a concrete example of how this can go wrong:
Order: 0016dfedd97fc2950e388d2971d718c7
The second query, however, attributes all item revenue to payment_sequential = 1 , so vouchers which are payment_sequential = 2 disappear from the breakdown:
Orders with split payment types: 2,246 (2.3% of all orders)
Voucher revenue (LLM fix): R$ 162,091.39 (1.0% of total)
Voucher revenue (correct): R$ 352,707.71 (2.2% of total)
Misattributed: R$ 190,616.32
The fix has a third, quieter problem: some orders have no payment_sequential = 1 row at all, so the WHERE clause drops them entirely:
Orders with no payment_sequential = 1: 80
The likely correct approach would allocate each order’s payment across categories proportionally by item price, accounting for vouchers, and potentially other sequential payment types. I haven’t added that query here, but it starts to get more complex. If you’re still reading here, then I think you get the point already about the LLM getting it wrong with a naive prompt.

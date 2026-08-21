---
source: "https://blog.getcassis.com/derivation-distance/"
hn_url: "https://news.ycombinator.com/item?id=49387258"
title: "Rules vs. examples: what helps LLMs write correct SQL?"
article_title: "Rules vs examples: how much does spelling things out help? | Cassis"
image: "https://blog.getcassis.com/_astro/derivation-distance.Crqr_Ixp.png"
author: "matthieu_bl"
captured_at: "2026-08-21T13:38:16Z"
capture_tool: "hn-digest"
hn_id: 49387258
score: 2
comments: 0
posted_at: "2026-08-21T12:43:29Z"
tags:
  - hacker-news
  - translated
---

# Rules vs. examples: what helps LLMs write correct SQL?

- HN: [49387258](https://news.ycombinator.com/item?id=49387258)
- Source: [blog.getcassis.com](https://blog.getcassis.com/derivation-distance/)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T12:43:29Z

## Translation

タイトル: ルールと例: LLM が正しい SQL を書くのに役立つものは何ですか?
記事のタイトル: ルールと例: 物事を詳しく説明することはどの程度役立ちますか? |カシス
説明: ルール、指標、または実際に実行された例?どのコンテキスト表現が LLM を短縮するかを測定しました

記事本文:
コンテンツへ移動 カシス製品
ストーリーに戻る
ルールと例: 物事を詳しく説明することはどの程度役立ちますか?
ルール、指標、または実際に実行された例はありますか?私たちは、どのコンテキスト表現が LLM の正解に至るまでの道のりを短縮するのか、そして間違った方向転換が良いコンテキストについて何を教えてくれるかを測定しました。
最初のカーブ、そしてそれが多くを語らなかった理由
例から得られるものはありますか、それとも単にコンテキストの追加だけでしょうか?
コンテキストが増えると結果が悪くなる場合
取得をテストしているだけですか?
それぞれの例では具体的に何を追加しているのでしょうか?
それは直感から始まりました。 LLM のパフォーマンスを向上させるためにサンプルが頻繁に推奨されるのはなぜですか?私たちの脳にとって、ソリューションをゼロから再構築するよりも、例を適応させる方が明らかに簡単です。 LLMでも同じでしょうか？与えられた情報から答えを導き出すために必要な「思考」が少なくなれば、モデルは役に立ちますか?
私たちはこの残りの作業を導出距離と呼びます。つまり、コンテキスト内の事実から答えを得るためにモデルがどの程度の変換と合成を行う必要があるかということです。一般的なルールでは、長い距離を残します。質問と同じ形の実際の例はほとんど残りません。
コンテキストを編集するときは、さまざまな可能性があります。一般的なルールを述べたり、一般的なパターンのスニペットを提供したり、本格的な例を提供したりすることができますが、実際には、通常は多少の組み合わせになります。私たちは何が最も効果的かを評価することにしました。この調査は、データ分析のユースケースに焦点を当てています。私たちは調査結果が他のユースケースに応用されることを期待していますが、それには検証が必要です。独自のユースケースで同様の実験を実行するのは簡単です。もしそうなら結果に興味があります！
間違ったターンも興味深いので、この記事は私たちがそれにどのようにアプローチしたかについての話です。それらは、優れたコンテキストを構成するものを明らかにします。
完全性が第一です。ビジネスルールがミスの場合

NG、正解を得るのは宝くじをするのと同じです。
導出距離の短縮は、難しい質問に対して効果を発揮します。完全なサンプル ライブラリを追加すると、Haiku は 70% から 82% に、Sonnet は 81% から 93% に上昇しました。
そのライブラリの下にあるルールとメトリクスを削除すると、標準精度は変更されず、厳密な精度は 1 ～ 4 ポイント以内のままになりましたが、コンテキストは約 30% 削減されました。
Opus は完全な英語の説明ですでに 98% に達していました。指標と例による追加は 1 パーセント ポイントのみでした。
関連する例をいくつか紹介するだけでは十分ではありませんでした。多くのクエリ パターンとその周りの小さな規則をカバーする広範なライブラリによって利点が現れました。
量よりも一貫性が重要です。 1 つの矛盾によって、追加のコンテキストの利点が失われる可能性があります。
普遍的な最適な設定はありません。小規模なモデルには、包括的な例が役立ちます。より強力なモデルは、より無駄のないコンテキストで適切にパフォーマンスを発揮します。
この記事で説明したすべての実験を再現するために必要なコードは、GitHub 上のオープンソースです。
私たちは、jaffle-shop (おもちゃのカフェ チェーン)、theLook (合成 e コマース ストア)、および Olist (実際のブラジル市場の輸出) という 3 つの公開分析データセットから開始しました。それぞれには、最初に Cassis をテストするために構築したコンテキスト レイヤー (テーブルと列の説明、散文で書かれたビジネス ルール、SQL を使用した名前付きメトリクス) が付属しています。各データセットについて、12 個の標準的な質問と、既知の答えを持つ 11 個の難しい質問も生成しました。
質問は 2 段階に分かれています。標準レベルは単純です。「総収益はいくらか」、「上位 5 店舗の店舗別の収益はいくらか」、「平均注文額はいくらか」です。 1 つの集計、1 つまたは 2 つの結合が完了しました。難しい層は複数段階の構成です。「各月の前月比収益増加率はいくらか」を集計する CTE が必要です。

es を月に変換し、それぞれを前のものと比較する LAG ウィンドウ関数。 「会社の総収益に占める各店舗の割合はいくらか」には、グループ化された集計に対するウィンドウ SUM が必要です。 「過去 90 日間にその前の 90 日間よりも多くの費用を支払った顧客の数」を確認するには、重複しない 2 つの日付ウィンドウ、顧客ごとの条件付き集計、および比較が必要です。これらはひっかけ質問ではなく、文脈で述べられたルールによって 1 つの答えが特定されるように作成しました (ご覧のとおり、最初の試行では完全に成功しませんでした)。その目的については後ほど説明します。
質問ごとに、導出作業を変化させる一連のレベルでコンテキストをレンダリングし、それを質問とともにモデルに渡します。モデルは SQL を記述し、それを実際のデータベースに対して実行し、結果を既知の答えと比較します。これを幅広い機能範囲にわたる 3 つのモデル (Haiku 4.5、Sonnet 4.6、Opus 4.8) に対して行い、各セルを 3 回繰り返して不安定さを見つけ、無効な SQL を記述するモデルを間違っているとカウントします。合計 2,484 回の実行になります。
採点は一段落する価値があります。なぜなら、それは間違った方向転換の 1 つであることが判明したからです。基本的な比較は標準の実行精度です。結果セットを順序付けされていない行のバッグとして扱い、列名を無視し (モデルのエイリアスは自由に)、浮動小数点数を小数第 4 位に四捨五入します。私たちの 1 年生はそこで立ち止まり、失敗を監査したところ、およそ 3 分の 1 が、成長率に関する ROUND(x, 2)、あちこちに追加の列、または参照解答に午前 0 時のタイムスタンプが付いている日付など、別の方法で提示された実質的に正しい解答であることがわかりました。プロンプトにはこれらの規則をモデルに伝えるものは何もなかったので、完全一致グレーディングは不当に罰せられました。たまたま優先順位を共有したモデルを優先して、すべてのカーブを歪めました。

ゴールドアンサージェネレーターのリース。採点者は、正解の表示バリアントを受け入れるようになりました。要求された値は、任意の列順序で、横に追加の列があり、モデル独自の表示精度 (小数点以下 2 桁) で、日付と午前 0 時のタイムスタンプを同じ瞬間として扱います。値自体を変更するものはすべて拒否されます。パーセンテージが尋ねられた分数、名前が尋ねられたエンティティ ID、または欠落している列などです。以下のすべての数値はこのグレーディングを使用します。
同じ理由で、結果の前に 1 回開示します。質問、参考回答、サンプル ライブラリは、ラインナップの中で最も強力なモデルである Opus 4.8 で作成され、実行および監査されました。寛容な採点によりプレゼンテーションのバイアスは取り除かれますが、質問が真に自由な選択を残している場合（どの順序で同点を破るかなど）、参照回答は Opus のデフォルトを体現しています。これらの選択に関して、このキーに基づいて Opus を評価することは部分的には自己同意を測定し、他のモデルは部分的にはデフォルトが Opus とどの程度一致しているかを測定します。 Opus のほぼ完璧なスコアを読むときは、そのことを念頭に置いてください。単一モデル内の比較は、ほとんどの結論を導き出しますが、モデルには依存しません。重要なときにこれに戻ります。
最初のカーブ、そしてそれが多くを語らなかった理由
最初のテストでは最も単純なことを行いました。さまざまなコンテキスト レベルでテストし、構造から要素を徐々に追加していきました。レベル 0 は、裸のテーブル名と列名でした。各レベルでは、説明から数式、実際の例に至るまで、コンテキスト レイヤーの要素がさらに追加されました。コンテキストが増える = 導出の労力が減る、ですよね?曲線は理論が予測したとおりに見えました。コンテキストを追加すると精度は着実に上昇し、弱いモデルは強いモデルよりも急上昇し、すべてがうまく扇状に広がりました。 B

しかし、それは単純な理由で間違ったものを測定しました。レベル 0 では、テーブルと列名だけがあり、モデルは収益をどのように測定するかなどのビジネス ルールをまったく知りませんでした。レベル 0 での各実行は、ほとんどサイコロを振ることに相当します。さらに悪いことに、データセットに関する時間情報がなかったため、時間に関連するすべての質問に答えることができませんでした。
したがって、当然のことながら、この情報を追加すると、答えはより良くなります。しかし、これが測定したのは、コンテキストに回答に必要な情報が含まれているかどうかでした。言い換えれば、完全性を測定したのです。確かに、完全性によって結果は向上しますが、それは必ずしも最新ニュースではありません。
コンテキストが完了すると、さらに興味深い違いが残ります。それは、モデルが認識した事実から答えを導き出すために、モデルがまだどれだけの作業を行う必要があるかということです。それが導入からの導出距離です。そこで、完全性を一定に保ちながら、すでに説明されている解決策の量を変えるためにコンテキスト レベルを再加工しました。
完全性を修正するのは理論的には簡単です。すべてのレベルにすべての事実が含まれていることを確認します。実際には少し難しく、収束するまでに失敗の分析を数回繰り返す必要がありました。
そこからは 3 つのことが重要になります。まず、完全性: 提供された事実から答えが可能か? 2 番目に、導入距離で定義した導出距離: モデルにどの程度の解パスが残っているか?第三に、コンテキストの負担: コンテキストの送信と維持にどれくらいのコストがかかりますか?また、コンテキストによって矛盾が生じる機会はどれくらいありますか?優れたコンテキストには、3 つすべてのバランスが取れている必要があります。
最終的には 3 つのコンテキスト レベルになりました。まず、事実が平易な英語で表され、表と列の説明に散りばめられています。次のレベルでは、構造化メトリクスとその名前、説明、SQL ex が追加されます。

圧迫感。最後のものには、例のライブラリが含まれています。質問ごとに、非常に似ているが同一ではない質問に答えるサンプル クエリを 1 つ生成しました。すべての例がコンテキストに含まれているため、モデルはすべての質問に対して同じコンテキストを認識しました。
このラダーは追加的です。例レベルには、その下のレベルからの説明、ルール、メトリックも含まれています。これにより、完全なパッケージが機能することがわかりますが、サンプルに先行する層が必要かどうかはわかりません。以下でその区別に戻ります。
これで、完全性が一定であれば、導出距離を測定できるはずです。それは違いを生むように見えました。
標準的な質問レベルでは、勾配は実際のものですが、控えめです。 3 つのデータセットをプールすると、Haiku は、簡潔だが完全なベースラインでの 82% から、メトリクスを使用した場合の 87%、例を使用した場合の 94% に上昇しました。 Opus はほとんど助けを必要としません。94% から開始し、サンプル ライブラリを使用すると 100% に達します。導出距離が最小のコンテキストでは、最小のモデルでも非常に高い精度で取得できます。
しかし、主なポイントは、このベンチマークは飽和しているということです。 Opus と Sonnet が 100% に達すると、結果の変動は主に測定誤差である可能性があります。これが、より難しい段階の質問を導入した理由です。
ハード層では、モデルはきれいに分離されます。 Opus は最初から天井に近いレベルで実行されます。簡潔な説明で 98%、指標と例で 99% を達成しました。小さいモデルはずっと上まで登ります。指標は説明よりもいくつかのポイントを追加し (Haiku 63% から 70%、Sonnet 76% から 81%)、例はそれぞれ 12 ポイント追加して 82% と 93% になります。
Opus はまだ飽和状態ですが、より厳しいベンチマークにより、モデルとコンテキスト レベルの両方がより大きく分離されます。説明により、質問に答えやすくなります。メトリクスは 1 つのローカル導出を削除します: メジャーがどのように行われるか

eが計算されます。例では、結合、集計、ウィンドウ、およびランキングを完全なクエリに構成する方法など、より困難な構造導出を削除できます。サンプル ライブラリは、メトリクス レイヤー全体よりも役立ちます。 Opus の場合、英語の説明からその構成をすでに実行できるため、曲線はほぼ平坦です。
モデルが実際に見ているものを見てみましょう。
「各月の前月比収益増加率はいくらですか?」という質問を考えてみましょう。第 1 レベルでは、わかりやすい英語のルールで、モデルが収益について知っていることはすべて、ORDERS 列の説明の 1 文です。
SUBTOTAL (DOUBLE): 売上税前の注文金額 (USD)。
収益と注文金額の基礎。
事実がそこにあるので、コンテキストは完成です。ただし、モデルでは、月ごとの集計、ウィンドウ関数、増加算術など、クエリ全体をゼロから構築する必要があります。これが最大導出距離です。
メトリクスを使用すると、モデルは代わりに次のように認識します。
収益
- 式: SUM("SUBTOTAL")
- テーブル: JAFFLE.ORDERS
この式は、収益がどのように測定されるかをより明確に示します。ウィンドウ関数を使用して月次集計を作成する方法については何も述べていません。モデルは CTE を発明し、 date_trunc を選択し、 LAG を計算し、成長率の演算を配線する必要があります。
eのとき

[切り捨てられた]

## Original Extract

Rules, metrics, or worked examples? We measured which context representations shorten an LLM

Skip to content Cassis Product
Back to stories
Rules vs examples: how much does spelling things out help?
Rules, metrics, or worked examples? We measured which context representations shorten an LLM's path to a correct answer, and what the wrong turns taught us about good context.
The first curve, and why it didn’t say much
Is the gain from examples or just more context?
When more context gives worse results
Are we just testing retrieval?
What is each example adding exactly?
It started from an intuition. Why are examples so often recommended to help LLMs perform well? For our brains, it’s clearly easier to adapt an example than to reconstruct a solution from scratch. Is it the same for LLMs? Does it help a model if it has less “thinking” to do to get from the information it’s given to the answer?
We call that remaining work the derivation distance : how much transformation and composition the model still has to do to get from the facts in its context to the answer. A general rule leaves a long distance. A worked example with the same shape as the question leaves almost none.
There’s a whole range of possibilities when editing context: one can state general rules, provide snippets for common patterns, provide full-fledged examples, and in practice it’s usually a bit of a mix. We set out to evaluate what works best. This study is focused on data analytics use cases. We expect the findings to transfer to other use cases, but that would need validation. It’s easy to run a similar experiment on your own use cases. I’d be interested in the results if you do!
This post is the story of how we approached that, because the wrong turns are also interesting. They surface what makes a good context.
Completeness comes first. If a business rule is missing, getting the right answer amounts to playing the lottery.
Shortening the derivation distance pays off on hard questions: adding a full example library raised Haiku from 70% to 82% and Sonnet from 81% to 93%.
Removing the rules and metrics beneath that library left standard accuracy unchanged and hard accuracy within one to four points, while cutting context by roughly 30%.
Opus was already at 98% with complete English descriptions. Metrics and examples added only one percentage point.
A few relevant examples were not enough. The gains appeared with a broad library covering many query patterns and the small conventions around them.
Consistency matters more than volume. A single contradiction can erase the benefit of additional context.
There is no universal best setup. Smaller models benefit from comprehensive examples. Stronger models perform well with leaner context.
The code needed to reproduce all the experiments mentioned in this article is open source on GitHub .
We started from three public analytics datasets: jaffle-shop (a toy cafe chain), theLook (a synthetic e-commerce store), and Olist (a real Brazilian marketplace export). Each comes with a context layer that we initially built to test Cassis: table and column descriptions, business rules written in prose, and named metrics with their SQL. For each dataset, we also generated 12 standard questions and 11 hard questions with known answers.
The questions come in two tiers. The standard tier is straightforward: “what is our total revenue”, “what is revenue by store for the top five stores”, “what is our average order value”. One aggregate, one or two joins, done. The hard tier is multi-step composition: “what is the month-over-month revenue growth rate for each month” needs a CTE that aggregates to months and then a LAG window function that compares each to the one before. “What is each store’s share of total company revenue” needs a window SUM over a grouped aggregate. “How many customers spent more in the last 90 days than in the 90 days before that” needs two non-overlapping date windows, per-customer conditional aggregation, and a comparison. These are not trick questions, and we wrote them so that the rules stated in the context would pin down a single answer (we did not fully succeed on the first try, as you will see). We will come back later to their purpose.
For each question we render the context at a series of levels supposed to vary the derivation effort, and hand it, with the question, to a model. The model writes SQL, we run it against the actual database, and compare the result to the known answer. We do this for three models spanning a wide capability range (Haiku 4.5, Sonnet 4.6, Opus 4.8), repeat each cell three times to catch flakiness, and count a model that writes invalid SQL as wrong. That makes for a total of 2,484 runs.
The grading deserves a paragraph, because it turned out to be one of the wrong turns. The base comparison is standard execution accuracy: treat the result set as an unordered bag of rows, ignore column names (models alias freely), and round floats to four decimals. Our first grader stopped there, and an audit of the failures showed that roughly a third were materially correct answers presented differently: a ROUND(x, 2) on a growth rate, an extra column here or there, or a date where the reference answer had a midnight timestamp. Nothing in the prompt told the model these conventions, so exact-match grading was unfairly punishing. It distorted every curve in favor of the models that happened to share the preferences of the gold answer generator. The grader now accepts presentation variants of the correct answer: the requested values in any column order, with extra columns alongside, at the model’s own display precision (two decimals minimum), and treating a date and a midnight timestamp as the same instant. It still rejects anything that changes the values themselves: a fraction where a percentage was asked, an entity ID where a name was asked, or a missing column. All numbers below use this grading.
One disclosure before the results, for the same reason. The questions, reference answers, and example library were drafted with Opus 4.8, the strongest model in the lineup, then executed and audited. Tolerant grading removes the presentation bias, but wherever a question leaves a genuinely free choice, like which order to break a tie in, the reference answer embodies Opus’s default. On those choices, grading Opus against this key partly measures self-agreement, and the other models partly measure how well their defaults match Opus’s. Keep that in mind when reading Opus’s near-perfect scores; the comparisons within a single model, which carry most of the conclusions, do not depend on it. We will come back to this when it matters.
The first curve, and why it didn’t say much
The first test did the simplest thing: test at different context levels, progressively adding more elements from our structure. Level zero was bare table and column names. Each level added more elements of the context layer, from descriptions to formulas to worked examples. More context = less derivation effort, right? The curve looked exactly like the theory predicted: accuracy rose steadily as we added context, the weak model rose more steeply than the strong one, and everything fanned out nicely. But it measured the wrong thing for a simple reason: at level zero, with only tables and column names, the model had no idea of the business rules, such as how revenue should be measured. Each run at level 0 pretty much amounted to rolling dice. Even worse, it did not have temporal information about the dataset, so all time-related questions could not be answered.
So of course, as you add this information, the answers get better. But what this measured was whether the context contained the information required to answer at all. In other words, it measured completeness . And yes, completeness improves results, but that’s not exactly breaking news.
Once the context is complete, a more interesting difference remains: how much work the model still has to do to get from the facts it sees to the answer. That is the derivation distance from the introduction. So we reworked our context levels to hold completeness constant while varying how much of the solution path was already spelled out.
The fix for completeness is easy in theory: make sure every level has all the facts. It’s a bit harder in practice and took a few iterations of analyzing the failures to converge.
From there, three things matter. First, completeness : is the answer possible from the supplied facts? Second, derivation distance , as defined in the introduction: how much of the solution path is still left to the model? Third, context burden : how much does the context cost to send and maintain, and how many opportunities does it create for contradictions? A good context needs to balance all three.
We ended up with three context levels. First, the facts are present in plain English, scattered across table and column descriptions. The next level adds structured metrics with their names, descriptions, and SQL expressions. The last one includes a library of examples. For each question, we generated one example query that answers a very similar but not identical question. All the examples were included in the context, so the model saw the same context for every question.
This ladder is additive: the examples level also contains the descriptions, rules, and metrics from the levels below it. It tells us that the full package works, but not whether the examples need those preceding layers. We return to that distinction below.
Now, with completeness constant, we should be able to measure derivation distance. It appeared to make a difference.
On the standard question tier, the gradient is real but modest. Pooled across the three datasets, Haiku goes from 82% at the terse-but-complete baseline to 87% with metrics and 94% with examples. Opus barely needs the help: it starts at 94% and reaches 100% with the example library. A context with minimal derivation distance is able to get even the smallest models to very good accuracy.
But the main takeaway is that this benchmark is saturated. With Opus and Sonnet reaching 100%, the variation in the results could be mostly measurement error. This is why we introduced the harder tier of questions.
On the hard tier, the models separate cleanly. Opus runs close to the ceiling from the start: 98% on terse descriptions, 99% with metrics and with examples. The smaller models climb the whole way up. Metrics add a handful of points over descriptions (Haiku 63% to 70%, Sonnet 76% to 81%), and examples add twelve more for each, to 82% and 93%:
While still saturated for Opus, the harder benchmark separates both the models and the context levels more significantly. Descriptions make the question answerable. Metrics remove one local derivation: how a measure is calculated. Examples can remove the harder structural derivation: how to compose the joins, aggregates, windows, and ranking into a complete query. The example library helps more than the whole metrics layer. For Opus, the curve is nearly flat because it is already able to perform that composition from English descriptions.
Let’s dive into what the models actually see.
Take the question “What is the month-over-month revenue growth rate for each month?” At the first level, plain English rules, everything the model knows about revenue is one sentence in the ORDERS column descriptions:
SUBTOTAL (DOUBLE): Order value before sales tax, in USD.
The basis for revenue and order-value figures.
The fact is there, so the context is complete. But the model has to build the whole query from nothing: the monthly aggregate, the window function, and the growth arithmetic. This is the maximum derivation distance.
With metrics, the model sees this instead:
revenue
- Expression: SUM("SUBTOTAL")
- Table: JAFFLE.ORDERS
That formula tells more clearly how revenue is measured. It says nothing about how to compose a monthly aggregate with a window function. The model has to invent the CTE, pick date_trunc , figure out LAG , and wire the growth-rate arithmetic.
When the e

[truncated]

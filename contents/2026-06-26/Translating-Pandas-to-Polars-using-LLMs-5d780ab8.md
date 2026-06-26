---
source: "https://pola.rs/posts/llm-polars-patterns/"
hn_url: "https://news.ycombinator.com/item?id=48684034"
title: "Translating Pandas to Polars using LLMs"
article_title: "Polars — Fluent, not native: agents translating pandas to Polars"
author: "jeroenjanssens"
captured_at: "2026-06-26T09:03:00Z"
capture_tool: "hn-digest"
hn_id: 48684034
score: 3
comments: 0
posted_at: "2026-06-26T08:39:03Z"
tags:
  - hacker-news
  - translated
---

# Translating Pandas to Polars using LLMs

- HN: [48684034](https://news.ycombinator.com/item?id=48684034)
- Source: [pola.rs](https://pola.rs/posts/llm-polars-patterns/)
- Score: 3
- Comments: 0
- Posted: 2026-06-26T08:39:03Z

## Translation

タイトル: LLM を使用したパンダから極座標への変換
記事のタイトル: Polars — ネイティブではなく流暢: パンダを Polars に翻訳するエージェント
説明: AI にパンダのコーパスを Polars に翻訳させました。ほとんどのワンショット翻訳は実行され、出力は一致しました。ほとんどの場合、コードは慣用的なものでした。 2 つの構造パターンが残り、新しいスキルでモデルが改良されます

記事本文:
ドキュメント ユーザー ガイド Rust Python リソース 当社のサービス アカデミー ブログ 私たちについて 採用中 Polars Cloud 採用中 ドキュメント ユーザー ガイド Rust Python リソース 当社のサービス アカデミー ブログ 私たちについて 採用中 --> --> --> --> --> --> --> --> --> --> --> ユーザー ガイド --> <!--
--> API --> --> --> Python --> <!--
ネイティブではなく流暢: パンダを Polars に翻訳するエージェント
投稿者:Thijs Nieuwdorp 投稿日: 2026 年 6 月 25 日
ますます多くの開発者にとって、初めて目にする Polar は言語モデルによって書かれています。
特定の変革に取り組む方法についてアドバイスを求めるだけの人もいれば、何ヶ月も自分で Polars クエリをプログラムしていない人もいます。
主なビジネス ケースの 1 つは、LLM によって移行が迅速に成功したというものです。
開発者は pandas スクリプトを入力し、Polars を取得し、出力が一致することを確認します。
その結果、通常は、より小型で安価なマシン上でパイプラインが桁違いに高速に実行されます。 1
翻訳コストが低いため、より早く投資収益率を達成できます。
このビジネス ケースの成功は、LLM が生成できる翻訳の品質にかかっています。
Polars API は表現力豊かであり、適切な翻訳では、単に実行されるコードではなく、各意図を最も直接的に記述する構造に到達する必要があります。
私たちは、現在のモデルがそのようなことをしているのか、それともたまたま有効な Polar であるパンダ型の習慣に頼っているのかを知りたかったのです。
Claude Opus 4.8 に pandas コーパスを Polars に翻訳してもらい、Polars の慣用的な使用の結果をチェックしました。
各ケースは新しいセッション、つまりワンショットで翻訳されました。
ほとんどの翻訳は実行され、出力は一致し、コードは慣用的な Polar として読み取られました。
残りの 2 つの構造パターン: モデルが意図 dir を表現する Polars 構造ではなく、既存の pandas 構造を変換するケース

まさに。
修正をスキルにパッケージ化し、ロードした状態でエクササイズを再実行し、違いを測定しました。
このスキルは、これらのパターンに沿ってモデルを正しい方向に動かしますが、特効薬ではありません。
Anthropic の最新モデルである Fable 5 は米国外では入手できないため、含めることができませんでした。
現在の状況を測定するために、パンダ コードの 3 つの本体を翻訳しました。
Polars-benchmark の pandas 実装を使用した 22 の PDS-H ベンチマーク クエリ (これは TPC-H から派生したベンチマークです。詳細については、こちらをご覧ください)。入力として pandas クエリだけを提供することで、同等の Polars クエリが利用できるため、簡単に検証可能なベンチマークが得られます。
3 つの公開パンダ教育リポジトリ 2 から抽出された 11 件の EDA コーパスで、ウィンドウ関数、リサンプリング、as-of 結合、再形成、文字列処理、ビニング、欠損データをカバーしています。これらの高度なパターンは、PDS-H ベンチマークにはありません。
Kaggle の 2 つの実際の ETL ノートブック。最初の手動探索に使用されます。これにより、比較できるより大きな実際の例が得られます。
各ケースは、Claude Opus 4.8 によって新しいセッションで翻訳されました。つまり、慣用的に Polars に翻訳するように明示的に要求するプロンプト付きの 1 ショットです。
Polars.testing.assert_frame_equal を使用して、その出力を元の pandas コードの出力と (PDS-H の場合は参照の回答と) 比較することで、すべての翻訳を検証しました。
比較のために、同じコーパスを Claude Sonnet 4.6 で実行しました。それらの結果は付録に記載されています。
1 年前、LLM は pandas と Polars API を頻繁に混合し、 groupby や with_column などの非推奨のメソッドに手を伸ばし、実行されないコードを生成していました。
本日、Opus 4.8 翻訳 33 件中 31 件が正しく実行され、検証された出力は一致しました。
2 つの例外は軽微でした。

ある場合にはインポートが発生し、別の場合には列参照エラーが発生します。
Opus は、時折小さなエラーが発生しますが、行ごとに翻訳するのではなく、再構築することで、以前のモデルよりも慣用的に Polars を翻訳しようとしています。
過去に問題があったパターンはすべてきれいに翻訳されました。
group_by 、 with_columns 、 is_in などの基本は正しく適用されます。
セミ結合とアンチ結合は、適切な箇所に使用されます。
条件付き集計はネイティブのままです。 map_elements は、PDS-H または EDA のいずれの翻訳にも現れませんでした。
遅延/熱心の区別が尊重され、collect() が必要な場所に表示されます。遅延 API では、クエリを作成すると、collect() を呼び出した後、Polars が最適化を適用してクエリを実行するだけです。
リストラップ引数、Python 側のスカラー演算、map_elements フォールバックは表示されませんでした。これらは以前のモデルでは一般的でした。付録の Sonnet 4.6 の結果は、その様子を示しています。
残っているパターンはより微妙です。
コードは実行され、出力は正しいですが、少数のケースでは、ソリューションは元の pandas コードの構造と同じように構造化されます。
モデルは Polars API をよく知っていますが、別の構造が意図をより直接的に表現するときを常に認識するとは限りません。
パターン 1: 意図ではなく構造を翻訳する
これは私たちが発見した中で最も明確なケースであり、排除するのが最も困難です。
この例では、ソース ノートブックは 5 分間のウィンドウ集計で 1 分あたりの株価に注釈を付けます。
このノートブックでは、一般的なパンダのレシピを使用しています。つまり、5 分のウィンドウにリサンプリングする 2 番目のフレームを作成し、merge_asof でマージし直します (as-of 結合では、各行がキー (ここではタイムスタンプ) によって、他のフレーム内の最も近い前の行と照合されます)。
LLM は、group_by_dynamic (Polars の resample-styl) を使用して、その構造を忠実に変換しました。

e 時間バケット化) と現在の結合:
# LLM の Polars 翻訳
Five_min = 分。 group_by_dynamic ( "datetime" 、毎 = "5m" )。集約 (
お願いします。列 (「閉じる」)。という意味です（）。エイリアス ( "window_price" )、
お願いします。列 (「ボリューム」)。合計()。エイリアス ( "window_volume" )、
)
結果 = 分。 join_asof (
5_分、
on = "日時" 、
戦略 = "後方" 、
許容誤差 = タイムデルタ (分 = 5 )、
)
# 慣用的な極地はどのように見えるか
結果 = 分。 with_columns (
お願いします。 Col (「閉じる」)
。意味（）
。 over (pl.col ( "datetime" ).dt.truncate ( "5m" ))
。エイリアス ( "window_price" )、
お願いします。 Col (「ボリューム」)
。合計()
。 over (pl.col ( "datetime" ).dt.truncate ( "5m" ))
。エイリアス ( "window_volume" )、
)
dt.truncate("5m") は各タイムスタンプを 5 分のバケットの先頭に切り捨て、 .over() はバケットごとに集計します。
実際の目的は「各行に 5 分間のバケットの集計で注釈を付ける」ことであり、.over() はそれを単一のコンテキストで正確に表現します。
.over() は Polars のウィンドウ関数です
これは、グループレベルの式を計算し、結果をそのグループ内の各行 (すべて with_columns 内) にブロードキャストする SQL の OVER (PARTITION BY ...) と同等です。
Polars はウィンドウ関数を集計から意図的に分離します。group_by().agg() はフレームをグループごとに 1 行に削減しますが、.over() は元のフレームの各行をグループレベルの値で強化します。
マッピング戦略の完全なセットについては、ウィンドウ関数ガイドを参照してください。
興味深いことに、パンダのソースが transform を使用すると、モデルは問題なくそれを .over() にマップします。
パンダはまさにこのケースを、 groupby(pd.Grouper(key="datetime", freq="5min")).transform(...) を使用してワンステップで表現できます。そのように書かれたノートブックは、おそらく .over() に直接変換されるでしょう。
ソースが集合フレームを構築して結合するときです。

モデルが意図ではなく回避策を変換していることがわかります。
このパターンは 33 件中 3 件で発生しました。
パターン 2: パイプラインの途中でスカラーを抽出する
1 つの PDS-H クエリは、データ自体から導出されたしきい値に基づいて行をフィルタリングします。つまり、部品の平均数量の一部を下回る行です。
その変換では、LLM が最初にしきい値を解決し、フィルタで使用する前に Python float として抽出しました。
# LLM が生成するもの
しきい値 = lf 。 select (pl. 列 ( "収益" ). 分位数 ( 0.9 ))。集める （）。項目 ()
結果 = lf 。フィルター (pl.col ( "revenue" ) > しきい値)。集める（）
# 慣用的な極地はどのように見えるか
結果 = lf 。フィルター (
お願いします。 Col ( "収益" ) > pl.列 (「収益」)。分位数 ( 0.9 )
）。集める（）
2 番目のバージョンは、集計が filter 内でブロードキャストされるため機能します。ブロードキャストとは、比較対象の列の長さに合わせてスカラー値が乗算されるため、分位数が 1 回計算され、すべての行と比較されます。
.item() は、Polars の外部の何かに対して Python 値が必要な場合に、パイプラインの最後にその場所があります。
ただし、 .item() は実体化された DataFrame 上にのみ存在するため、これをパイプラインの途中で使用すると、時期尚早にcollect() が強制的に実行されます。
本来 1 つのクエリが 2 つになります。ソースが 2 回スキャンされ、中間結果が具体化され、オプティマイザはそれぞれの半分を分離して確認します。
もう 1 つの一貫した観察: .agg() 内では、モデルはトップレベルの短縮表現 pl.sum("x") ではなく、スペルアウトされた形式 pl.col("x").sum() を使用しています。
どちらも有効であり、スタイルに合わせて選択します。
これはコーパス全体で 48 か所に出現していましたが、スキルにより出現数が 40 に減りました。
これらのパターンに共通するのは、モデルが一歩下がってクエリの背後にある意図を再検討するのではなく、パンダのソースを一度に 1 ステップずつ変換することです。

。
各ステップで最も安全な方法は、pandas の操作を、同じことを行う Polars の操作にマッピングすることです。
これにより、結果が元の構造に固定されたままになります。そのため、翻訳は、中間フレームや結合バックに至るまで、元のパンダ スクリプトの形状によく似ています。
単一の .over() を中心にクエリ全体を再定式化するには、その前の行から一歩下がって意図を認識する必要がありますが、モデルは一貫してそのステップを実行しません。
非慣用的なコードが実行されて正しい答えが生成される場合、エラーや警告は生成されず、モデルの変換を強制するような最適化されていないという信号も生成されません。
ステートメントごとに忠実な翻訳が合格すれば、その基準から言えば成功と言えます。
アクセントも言語モデルに特有のものではありません。
何年もパンダを続けた後に Polars に来た開発者は、より慣用的な Polars クエリの記述方法を学ぶまで、古いパターンを引き継ぐ傾向があります。
あるツールから別のツールへの移行はソースから始まり、ターゲット ライブラリに関係なく、行ごとの変換によってソースの構造が結果に反映されます。
ライブラリの作成者には、LLM のトレーニング データに含まれる慣用的な例 (この投稿など!) を公開する、PolarsInefficientMapWarning の方法で機械操作可能な警告を発する、エージェントが読み込める形式でガイダンスを提供する、という 2 つのオプションが残されています。
最後のものは次に検討したものです。
上記のパターンは修正可能です。
これらのパターンは、PDS-H および EDA 演習からの広範な翻訳メモとともに、Polars スキルとしてパッケージ化されています。
エージェントにロードされると、プロンプトを表示しなくても、モデルが正しいパターンに向けて誘導されます。
このスキルは無料で利用でき、GitHub リポジトリでホストされています。
私たちは再ラします

スキルがロードされた 33 件の翻訳すべてで、ケースごとに新しいセッションで同じモデルを実行します。
これは小規模な実験であり、単一モデルのケースごとに 1 つの変換を行うため、数値をベンチマークではなく方向性のシグナルとして扱います。この効果は 33 件すべてのケースで維持されました。
このスキルはすべてを解決する特効薬ではありませんが、LLM を正しい方向に導きます。
構造パターンが改善されました。結合バックのケースは 3 から 1 に減少し、パイプライン中の .item() の 1 つのケースはインライン ブロードキャストに解決されました。
正解率は 31 から 32 に上昇しましたが、スキルに起因する回帰が 1 つ発生しました。スキルの遅延 API ガイダンスにより、すでに熱心なフレームで .collect() が呼び出されました。
集計の簡略表現ではほとんど変化がなく、48 件の発生が 40 件に減少し、17% 減少しました。
実際に構造的な改善がどのように行われるかを次に示します。
PDS-H クエリ 17 は、部品ごとに導出されたしきい値 (その部品の平均数量の 20%) に基づいて品目をフィルタリングします。
スキルがなければ、モデルはパーツごとの平均を別のフレームとして構築し、それを結合し直していました。
スキルがロードされると、グループごとのしきい値がフィルター内の .over() ウィンドウになり、個別の集約フレームと結合バックが 1 つのパスに折りたたまれます。
# なし

[切り捨てられた]

## Original Extract

We had AI translate a pandas corpus to Polars. Most one-shot translations ran and the outputs matched. The code was idiomatic in most cases. Two structural patterns remain, and our new skill improves the model

Docs User guide Rust Python Resources Our services Academy Blog About us We're hiring Polars Cloud We're hiring Docs User guide Rust Python Resources Our services Academy Blog About us We're hiring --> --> --> --> --> --> --> --> --> --> --> User guide --> <!--
--> API --> --> --> Python --> <!--
Fluent, not native: agents translating pandas to Polars
By Thijs Nieuwdorp on Thu, 25 Jun 2026
For a growing number of developers, the first Polars they ever see was written by a language model.
Some just ask them for advice on how to tackle certain transformations, while others haven’t programmed a Polars query themselves in months.
One of the leading business cases is that LLMs have made migration a quick win.
A developer feeds in a pandas script, gets back Polars, and verifies that the output matches.
The result is typically a pipeline that runs an order of magnitude faster on smaller, cheaper machines. 1
With a lower cost of translation, the return on investment is achieved even sooner.
The success of this business case depends on the quality of the translation an LLM can produce.
The Polars API is expressive, and a good translation should reach for the construct that states each intent most directly, not merely code that runs.
We wanted to know whether current models do that, or whether they fall back on pandas-shaped habits that happen to be valid Polars.
We had Claude Opus 4.8 translate a pandas corpus to Polars and checked the result for idiomatic use of Polars.
Each case was translated in a fresh session: one shot.
Most translations ran, the outputs matched, and the code read as idiomatic Polars.
Two structural patterns remained: cases where the model translates the existing pandas structure rather than the Polars construct that expresses the intent directly.
We packaged fixes into a skill, re-ran the exercise with it loaded, and measured the difference.
The skill nudges the model in the right direction on those patterns, but it is not a magic bullet.
Fable 5, Anthropic’s newest model, is not available to us outside the United States, so we could not include it.
To measure where things stand today, we translated three bodies of pandas code:
the 22 PDS-H benchmark queries (This is a benchmark derived from TPC-H. Learn more here ), using the pandas implementations from polars-benchmark . By providing just the pandas queries as input, we have an easily verifiable benchmark, since we have equivalent Polars queries available.
an 11-case EDA corpus drawn from three public pandas teaching repositories 2 , covering window functions, resampling, as-of joins, reshaping, string processing, binning, and missing data. These advanced patterns are not found in the PDS-H benchmark.
two real-world ETL notebooks from Kaggle, used for initial manual exploration. This gives larger real life examples to compare against.
Each case was translated by Claude Opus 4.8 in a fresh session: one shot, with a prompt that explicitly asked it to translate idiomatically to Polars.
We verified every translation by comparing its output against the output of the original pandas code (and for PDS-H against the reference answers) with polars.testing.assert_frame_equal .
We ran the same corpus through Claude Sonnet 4.6 for comparison. Those results appear in the appendix.
A year ago, LLMs frequently mixed the pandas and Polars APIs, reaching for deprecated methods like groupby and with_column , and producing code that did not run.
Today, 31 of 33 Opus 4.8 translations ran correctly and the verified outputs matched.
The two exceptions were minor: a missing import in one case and a column-reference error in another.
Opus attempts to translate Polars more idiomatically than earlier models did, restructuring rather than translating line by line, although it occasionally introduces small errors.
The patterns that were problematic in the past all translated cleanly:
Basics like group_by , with_columns , is_in land correctly.
Semi and anti joins are used where they fit.
Conditional aggregations stay native. map_elements did not appear in any of the PDS-H or EDA translations.
The lazy/eager distinction is respected and collect() appears where it should. In the lazy API you build up a query and Polars only runs it, with optimizations applied, once you call collect() .
List-wrapped arguments, Python-side scalar arithmetic, and map_elements fallbacks did not appear. These were common tells in earlier models; the Sonnet 4.6 results in the appendix show what that looked like.
The patterns that remain are subtler.
The code runs and the output is correct, but in a small number of cases the solution is structured the way the original pandas code was structured.
The model knows the Polars API well, but does not always recognize when a different construct expresses the intent more directly.
Pattern 1: Translating structure instead of intent
This is the clearest case we found, and the hardest to eliminate.
In this example, the source notebook annotates per-minute stock prices with 5-minute window aggregates.
The notebook uses a common pandas recipe: create a second frame that resamples to 5-minute windows, then merge it back with merge_asof (an as-of join matches each row to the nearest preceding row in the other frame by key, here by timestamp).
The LLM translated that structure faithfully, using group_by_dynamic (Polars’ resample-style time bucketing) and an as-of join:
# The LLM's Polars translation
five_min = minute . group_by_dynamic ( "datetime" , every = "5m" ). agg (
pl. col ( "close" ). mean (). alias ( "window_price" ),
pl. col ( "volume" ). sum (). alias ( "window_volume" ),
)
result = minute . join_asof (
five_min,
on = "datetime" ,
strategy = "backward" ,
tolerance = timedelta (minutes = 5 ),
)
# What idiomatic Polars looks like
result = minute . with_columns (
pl. col ( "close" )
. mean ()
. over (pl. col ( "datetime" ).dt. truncate ( "5m" ))
. alias ( "window_price" ),
pl. col ( "volume" )
. sum ()
. over (pl. col ( "datetime" ).dt. truncate ( "5m" ))
. alias ( "window_volume" ),
)
dt.truncate("5m") rounds each timestamp down to the start of its 5-minute bucket, and .over() aggregates per bucket.
The actual intent is “annotate each row with an aggregate over its 5-minute bucket”, and .over() expresses exactly that in a single context.
.over() is Polars’ window function
It is the equivalent of SQL’s OVER (PARTITION BY ...) which computes a group-level expression and broadcasts the result to each row in that group, all inside with_columns .
Polars separates window functions from aggregations deliberately: group_by().agg() reduces the frame to one row per group, while .over() enriches each row of the original frame with a group-level value.
See the window functions guide for the full set of mapping strategies.
Interestingly, when the pandas source uses transform , the model maps it to .over() without trouble.
pandas can express this very case that way, in one step with groupby(pd.Grouper(key="datetime", freq="5min")).transform(...) , and a notebook written like that would most likely have been translated to .over() directly.
It is when the source builds an aggregate frame and joins it back that the model translates the workaround instead of the intent.
This pattern appeared in 3 of 33 cases.
Pattern 2: Extracting scalars mid-pipeline
One PDS-H query filters rows against a threshold derived from the data itself: rows below a fraction of a part’s average quantity.
In that translation, the LLM resolved the threshold first, extracting it as a Python float before using it in the filter.
# What LLMs produce
threshold = lf . select (pl. col ( "revenue" ). quantile ( 0.9 )). collect (). item ()
result = lf . filter (pl. col ( "revenue" ) > threshold). collect ()
# What idiomatic Polars looks like
result = lf . filter (
pl. col ( "revenue" ) > pl. col ( "revenue" ). quantile ( 0.9 )
). collect ()
The second version works because aggregations broadcast inside filter . Broadcasting means a scalar value is multiplied to match the length of the column it is compared against, so the quantile is computed once and compared against every row.
.item() has its place: at the end of a pipeline, when you need a Python value for something outside Polars.
But .item() only exists on materialized DataFrames, so using it mid-pipeline forces a premature collect() .
What should be one query becomes two: the source is scanned twice, an intermediate result materializes, and the optimizer sees each half in isolation.
One other consistent observation: inside .agg() , the model uses the spelled-out form pl.col("x").sum() rather than the top-level shorthand pl.sum("x") .
Both are valid and the choice is stylistic.
It appeared in 48 places across the corpus and the skill reduced it to 40 occurrences.
The common thread across these patterns is that the model translates the pandas source one step at a time rather than taking a step back and reconsidering the intent behind the query.
At each step, the safest move is to map a pandas operation to the Polars operation that does the same thing.
That keeps the result anchored to the structure of the original, which is why a translation so often resembles the shape of the pandas script it came from, down to the intermediate frames and join-backs.
Reformulating the whole query around a single .over() requires stepping back from the line in front of it and recognizing the intent, and the model does not consistently take that step.
Non-idiomatic code that runs and produces the right answer generates no error or warning, no signal that anything is suboptimal that would push the model to translate differently.
A faithful, statement-by-statement translation that passes is, by that measure, a success.
The accent is not unique to language models, either.
Developers who come to Polars after years of pandas tend to carry their old patterns over until they learn the more idiomatic ways of writing Polars queries.
Any migration from one tool to another starts from the source, and a line-by-line translation will carry the source’s structure into the result regardless of the target library.
Library authors are left with a couple of options: publish idiomatic examples that land in the LLM’s training data (such as this post!), emit machine-actionable warnings the way PolarsInefficientMapWarning does, and ship guidance in a form agents can load.
That last one is what we explored next.
The patterns mentioned above are fixable.
These patterns, along with the broader translation notes from the PDS-H and EDA exercises, are packaged as a Polars skill.
Loaded into an agent, it steers the model toward correct patterns without you having to prompt for them.
The skill is freely available and hosted in a GitHub repo .
We re-ran all 33 translations with the skill loaded, running the same model with a fresh session per case.
This is a small-scale experiment, one translation per case on a single model, so treat the numbers as a directional signal rather than a benchmark. The effect held across all 33 cases.
The skill is not a magic bullet that solves everything, but it nudges the LLM in the right direction.
Structural patterns improved: join-back cases dropped from 3 to 1, and the one mid-pipeline .item() case resolved to inline broadcasting.
Correctness moved from 31 to 32 correct, with one skill-induced regression: the skill’s lazy-API guidance caused .collect() to be called on an already-eager frame.
The aggregation shorthand tell barely moved: 48 occurrences dropped to 40, a 17% reduction.
Here is what the structural improvement looks like in practice.
PDS-H query 17 filters line items against a threshold derived per part: 20% of that part’s average quantity.
Without the skill, the model built the per-part average as a separate frame and joined it back.
With the skill loaded, the per-group threshold becomes an .over() window inside the filter, and the separate aggregate frame and the join-back collapse into a single pass:
# Without the

[truncated]

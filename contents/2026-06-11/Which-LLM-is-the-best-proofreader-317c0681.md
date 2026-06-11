---
source: "https://revise.io/errata-bench"
hn_url: "https://news.ycombinator.com/item?id=48490478"
title: "Which LLM is the best proofreader?"
article_title: "ErrataBench - Which LLM is the best proofreader?"
author: "artursapek"
captured_at: "2026-06-11T16:29:29Z"
capture_tool: "hn-digest"
hn_id: 48490478
score: 1
comments: 0
posted_at: "2026-06-11T14:00:52Z"
tags:
  - hacker-news
  - translated
---

# Which LLM is the best proofreader?

- HN: [48490478](https://news.ycombinator.com/item?id=48490478)
- Source: [revise.io](https://revise.io/errata-bench)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T14:00:52Z

## Translation

タイトル: どの LLM が最高の校正者ですか?
記事のタイトル: ErrataBench - どの LLM が最高の校正者ですか?
説明: ErrataBench は、英語テキスト内のエラーを見つけて修正する能力について LLM を評価するための Revise のベンチマークです。

記事本文:
ErrataBench - どの LLM が最高の校正者ですか? Docs Rewriter の改訂 ヘルプ 価格 ErrataBench v1 GitHub 全体ランキング 効率 方法論 エラー カテゴリ どの LLM が最高の校正者ですか?
--samples 3 --chunk-size 2000 --max-turns-per-chunk 3 で 2,290 回の実行にわたって 72 のモデル バリアントのベンチマークを実行 実行統計 合計実行時間 6 日 23 時間 3 分 合計コスト $1,130
2026 年 6 月 9 日、午後 3 時 31 分更新 - コミット 42b8c5c - art@revise.io
LLM は高速でコスト効率の高い校正ツールを作成しますが、どのモデルが最適なのでしょうか?エラーをどれだけ徹底的に見つけて修正できるか、そしてどれだけ効率的にできるでしょうか?
ErrataBench は、さまざまな書き込みエラーを含む大量のテキスト サンプルに対して単純なエージェント ループを実行することで、これらすべてを測定します。モデルには、「検索と置換」ツールを使用して、できる限り多くの問題を見つけて修正するよう求めます。
下のグラフは成功率の総合ランキングです。モデルは何パーセントのエラーを見つけて正しく修正しましたか?
正常に修正されなかった問題については、省略 (「未対処」) と、裁判官によって不正確として拒否された修正の試み (「悪い修正」) を区別します。
モデルにはすべて同じプロンプト、ツール、テキスト チャンク、およびターン数が与えられました。
ここでは、全体的な成功率を速度やコストなどの他の側面と比較します。デフォルトでは、各モデルの最良の推論バリアントのみが表示されます。すべてのデータ ポイントを表示するには、[最良のバリアントのみ] をオフに切り替えます。
この 2 次元の比較により、何が重要かに応じて、パフォーマンスとコストの全体像がわかります。ツール呼び出し効率は、モデルの変更がどの程度うまく機能したかを測定します。
「支配的なモデルを強調表示」を使用して、表示している軸上でどのモデルが支配的であるかを確認します。
ErrataBench は、さまざまなソース (文献、法律、技術マニュアル) からの英語テキストのデータセットに対してモデルをテストします。それぞれ

ソースは、いくつかのカテゴリにわたって、さまざまな記述ミスを伴って変更されています。これらのエラーは、対応するファイルで追跡されます。
ソース テキストは、破損者モデルとレビュー モデルのペアを使用して変更されます。破損者は、テキストに小さな変更を加えて、エラー カテゴリ分類から抽出された特定のタイプのエラーを挿入するように指示されます。レビュー モデルは、変更がデータセットに受け入れられる前に、セカンドオピニオンとして変更をレビューします。
このベンチマークは、変更された各テキストに対して単純なエージェント ループを実行し、段落全体をまとめたまま、一度に最大 2000 語のモデル チャンクを表示します。モデルは単にテキストを注意深く校正し、意味と詳細を維持しながらエラーを修正するように求められます。モデルには、テキスト内にどのようなタイプのエラーが隠れているのか、またはその数は何個あるのかについての詳細は示されていません。
エージェント ループは、モデルに各チャンクで最大 3 ターンを与えます。その間、モデルは、find_and_replace (外科的な変更を行うため) および replace_paragraph (広範囲の書き換えのため) という単純なツールを使用してテキストを変更できます。モデルがテキストを元のものに変更すると、エラーはすぐに修正されたものとしてカウントされます。モデルがそれを別のものに変更した場合、LLM ジャッジを使用して、その代替案も正しいかどうかが決定されます。
見出しの修正率スコアは、すべての実行のプールされた平均ではなく、データセットごとの均等に加重された平均です。 ErrataBench は、ソース データセット d ごとに、そのデータセットの成功した実行品質スコアを平均して means(q_d) を取得し、データセット全体でのデータセット レベルの平均を計算します。これにより、成功した繰り返しが多いデータセットが、少ないデータセットよりも多くカウントされることがなくなります。
他の散布図メトリクスも同様に導出されます。コスト効率は、実行レベルで USD ごとに問題が修正され、データセットごとおよびデータ全体で平均化されます。

など。速度は、実行時間の 1 分ごとに修正された問題を同じ方法で集計したものです。ツール呼び出し効率は、100 個のベンチマーク対象ツール文字ごとに修正された問題として表示されます。これは、基礎となる解決された問題ごとのツール文字数メトリクスを反転して表示します。使用ターン数は、データセット全体のチャンクあたりの平均ターン数です。
散布図の一貫性メトリックは median_d range(q_d) で、データセットごとの品質範囲の中央値です。 ErrataBench は、データセットごとに、成功した実行品質スコアの最高と最低の間の広がりを測定し、それらのデータセット レベルの範囲の中央値を取得します。 UI では、これは Consistency として表示されます。値が低いほど、モデルの校正品質が実行ごとに一貫していることを意味しますが、中央値により、少数のノイズの多いソース テキストに対するメトリクスの感度が低くなります。
ソース コード、データセット、結果は GitHub で入手できます。結果は、OpenRouter API キーを使用して再現できます。このリポジトリには、ソース テキストを使用して独自のデータセットを作成するツールも含まれているため、新しいサンプルに対してこのベンチマークを簡単に実行できます。
ベンチマークには、エージェント ループがどのように動作するかを決定する 2 つの主要な入力があります。それは、チャンク サイズとチャンクごとのターン数です。これらの値を変更すると、異なる結果が生じる可能性があります。具体的には、チャンクが大きくターン数が少ないほど、推論が少ないモデルが優先されます。高度な推論モデルは、より大きなチャンクが提示され、編集の回数が少なくなると、失われる傾向があります。ここで公開されているデータは、 --samples 3 --chunk-size 2000 --max-turns-per-chunk 3 でベンチマークを実行することによって生成されており、これは合理的な一般設定として決定されました。
ErrataBench で使用されるエラー カテゴリの完全なリストは以下にあります。
正書法と語形 誤字：太指誤字脱字挿入転置重複
スペルと大文字小文字の区別:

単語以外のスペル エラー 実際の単語のスペル エラー 大文字化エラー スペース エラー ハイフネーションまたは複合エラー
形態と語形変化 : 複数形の誤り 活用の誤り 語形の誤り
語彙の選択と混同の可能性 単語の選択: 単語の誤用 同音異義語の混同 同音異義語または近隣接混同 マラプロピズム コロケーション エラー
聞き間違いと再解釈：モンドグリーンエッグコーン
文法と構文の一致 : 主語と動詞の一致 代名詞の一致
動詞システム: 時制助動詞または助動詞が正しくありません。
文節と文の構造 : 文の断片の連続、融合文、またはカンマ継ぎ 語順エラー 単語の欠落 余分な単語の並列エラー 相対節エラー 調整または従属エラー
機能語と参照: 冠詞または限定詞のエラー 前置詞のエラー 代名詞-格のエラー 代名詞-参照のエラー ぶら下がりまたは修飾子の配置が間違っている
極性と比較 : 二重負の否定範囲誤差 比較誤差または最上級誤差 可算誤差
句読点と境界 句読点のエラー アポストロフィのエラー 引用または区切り文字のエラー 文の境界のエラー
意味論、言説、およびスタイル 意味論 : 意味上の異常な参照の曖昧さ
談話 : 談話の一貫性の誤り、文間の時制の不一致
スタイルとレジスタ: 冗長性または冗長性 レジスタの不一致 ぎこちない言い回し 非標準的な方言または口語形式
すべてのモデルのバリアント 語彙の選択と混同性 慣用性
クロード・ファブル 5 (中) 15 / 15 ( 100.0% ) クロード・作品 4.6 (高) 15 / 15 ( 100.0% ) クロード・作品 4.6 (低) 15 / 15 ( 100.0% ) クロード・作品 4.6 (中) 15 / 15 ( 100.0% ) ジェミニ3.1 プロ プレビュー (高) 15 / 15 ( 100.0% ) クロード フェイブル 5 (高) 14 / 15 ( 93.3% ) クロード フェイブル 5 (低) 14 / 15 ( 93.3% ) クロード オーパス 4.6 (なし) 13 / 15 ( 86.7%

) クロード オーパス 4.8 (低) 13 / 15 ( 86.7% ) クロード オーパス 4.7 (低) 12 / 15 ( 80.0% ) 全リスト 句読点と境界 文境界エラー
クロード・フェイブル 5 (低) 30 / 30 ( 100.0% ) GPT 5.5 (高) 30 / 30 ( 100.0% ) GPT 5.5 (中) 30 / 30 ( 100.0% ) クロード・フェイブル 5 (高) 29 / 30 ( 96.7% ) Gemini 3.1 Pro プレビュー (低) 29 / 30 ( 96.7% ) クロード・ファブル 5 (中) 28 / 30 ( 93.3% ) GPT 5.4 ミニ (高) 27 / 30 ( 90.0% ) クロード・オーパス 4.7 (なし) 26 / 30 ( 86.7% ) クロード・オーパス 4.7 (高) 25 / 30 ( 83.3% ) クロード作品 4.7 (中) 25 / 30 ( 83.3% ) 全リスト意味論的談話とスタイル談話
クロード オーパス 4.6 (高) 12 / 15 ( 80.0% ) クロード オーパス 4.6 (低) 12 / 15 ( 80.0% ) クロード オーパス 4.6 (中) 12 / 15 ( 80.0% ) クロード オーパス 4.7 (高) 12 / 15 ( 80.0% ) クロード オーパス 4.7 (低) 12 / 15 ( 80.0% ) クロード オーパス 4.7 (中) 12 / 15 ( 80.0% ) クロード オーパス 4.7 (なし) 12 / 15 ( 80.0% ) クロード オーパス 4.8 (高) 12 / 15 ( 80.0% ) クロード オーパス 4.8 (中) 12 / 15 ( 80.0% ) Claude Opus 4.8 (なし) 12 / 15 ( 80.0% ) 全リスト 文法と構文 関数の単語とリファレンス
クロード・ファブル 5 (低) 112 / 123 ( 91.1% ) クロード・ファブル 5 (高) 111 / 123 ( 90.2% ) GPT 5.5 (中) 111 / 123 ( 90.2% ) クロード・ファブル 5 (中) 110 / 123 ( 89.4% ) GPT 5.5 (高) 110 / 123 ( 89.4% ) GPT 5.4 Mini (高) 108 / 123 ( 87.8% ) GPT 5.5 (低) 107 / 123 ( 87.0% ) Claude Opus 4.7 (なし) 105 / 123 ( 85.4% ) Gemini 3.1 Pro プレビュー(中) 103 / 123 ( 83.7% ) クロード作品 4.6 (高) 102 / 123 ( 82.9% ) 全リスト 句読点と境界線 引用符または区切り文字のエラー
クロード・ファブル 5 (低) 31 / 33 ( 93.9% ) GPT 5.5 (低) 31 / 33 ( 93.9% ) クロード・ファブル 5 (高) 30 / 33 ( 90.9% ) クロード・ファブル 5 (中) 30 / 33 ( 90.9% ) クロード・オーパス 4.7 (なし) 30 / 33 ( 90.9% ) GPT 5.4 (中) 30 / 33 ( 90.9% ) GPT 5.5 (なし) 30 / 33 ( 9

0.9% ) Gemini 3 Flash プレビュー (高) 29 / 33 ( 87.9% ) GPT 5.4 (高) 29 / 33 ( 87.9% ) GPT 5.4 (なし) 29 / 33 ( 87.9% ) 全リスト セマンティクス 談話およびスタイルセマンティクス
GPT 5.5 (高) 75 / 78 ( 96.2% ) GPT 5.5 (中) 75 / 78 ( 96.2% ) クロード・ファブル 5 (高) 74 / 78 ( 94.9% ) クロード・ファブル 5 (低) 74 / 78 ( 94.9% ) クロード・ファブル 5 (中) 73 / 78 ( 93.6% ) クロード オーパス 4.6 (なし) 72 / 78 ( 92.3% ) クロード オーパス 4.6 (中) 71 / 78 ( 91.0% ) GPT 5.4 (中) 71 / 78 ( 91.0% ) Gemini 3.1 Pro プレビュー (中) 70 / 78 ( 89.7% ) Glm 5.1 (低) 70 / 78 ( 89.7% ) 全リスト 句読点と境界 句読点エラー
GPT 5.4 (中) 15 / 15 ( 100.0% ) GPT 5.4 ミニ (中) 15 / 15 ( 100.0% ) GPT 5.5 (高) 15 / 15 ( 100.0% ) GPT 5.5 (低) 15 / 15 ( 100.0% ) GPT 5.5 (中) 15 / 15 ( 100.0% ) GPT 5.5 (なし) 15 / 15 ( 100.0% ) Gemini 3 Flash プレビュー (高) 14 / 15 ( 93.3% ) Gemini 3.1 Pro プレビュー (高) 14 / 15 ( 93.3% ) GPT 5.4 Mini (高) 14 / 15 ( 93.3% ) Grok 4.20 (高) 14 / 15 ( 93.3% ) 完全なリストの文法と構文句と文の構造
クロード・フェイブル 5 (高) 243 / 258 ( 94.2% ) Gemini 3.1 Pro プレビュー (高) 242 / 258 ( 93.8% ) GPT 5.5 (高) 241 / 258 ( 93.4% ) クロード・フェイブル 5 (低) 240 / 258 ( 93.0% ) GPT 5.4 (中) 240 / 258 ( 93.0% ) GPT 5.5 (低) 240 / 258 ( 93.0% ) クロード・ファブル 5 (中) 239 / 258 ( 92.6% ) GPT 5.5 (中) 239 / 258 ( 92.6% ) Gemini 3.1 Pro プレビュー (低) 238 / 258 ( 92.2% ) GPT 5.4 Mini (Medium) 238 / 258 ( 92.2% ) 完全なリストの文法および構文の合意
Gemini 3 Flash プレビュー (高) 66 / 66 ( 100.0% ) クロード フェイブル 5 (高) 63 / 66 ( 95.5% ) クロード フェイブル 5 (中) 63 / 66 ( 95.5% ) クロード フェイブル 5 (低) 62 / 66 ( 93.9% ) Gemini 3.1 Pro プレビュー (高) 62 / 66 ( 93.9% ) GPT 5.5 (中) 62 / 66 ( 93.9% ) GPT 5.5 (高) 62 / 66 ( 93.9% ) GPT 5。

5 (低) 62 / 66 ( 93.9% ) Claude Opus 4.8 (なし) 60 / 66 ( 90.9% ) Glm 5.1 (高) 60 / 66 ( 90.9% ) 全リスト 文法と構文の極性と比較
クロード・ファブル 5 (高) 54 / 57 ( 94.7% ) クロード・ファブル 5 (低) 54 / 57 ( 94.7% ) クロード・ファブル 5 (中) 54 / 57 ( 94.7% ) クロード・作品 4.6 (なし) 54 / 57 ( 94.7% ) クロード・作品 4.7 (高) 54 / 57 ( 94.7% ) クロード オーパス 4.7 (低) 54 / 57 ( 94.7% ) クロード オーパス 4.7 (中) 54 / 57 ( 94.7% ) クロード オーパス 4.7 (なし) 54 / 57 ( 94.7% ) クロード オーパス 4.8 (高) 54 / 57 ( 94.7% ) クロード Opus 4.8 (低) 54 / 57 ( 94.7% ) 全リスト カテゴリをさらに読み込む 会社の改訂 製品の更新

## Original Extract

ErrataBench is Revise's benchmark for evaluating LLMs on their ability to find and fix errors in English text.

ErrataBench - Which LLM is the best proofreader? Revise Docs Rewriter Help Pricing ErrataBench v1 GitHub Overall Rankings Efficiency Methodology Error Categories Which LLM is the best proofreader?
Benchmarked 72 model variants across 2290 runs with --samples 3 --chunk-size 2000 --max-turns-per-chunk 3 Run stats Total runtime 6d 23h 3m Total cost $1,130
Updated Jun 9, 2026, 3:31 PM - commit 42b8c5c - art@revise.io
LLMs make fast, cost-effective proofreaders - but which models are the best? How thoroughly can they find and fix errors, and how efficiently?
ErrataBench measures all of this by running a simple agent loop over large samples of text, each containing a wide variety of writing errors. We ask models to find and fix as many problems as they can find using "find and replace" tools.
The chart below is an overall ranking of success rate. What percentage of errors did the model find and correctly fix?
For issues not fixed successfully, we make a distinction between omissions ("Not Addressed") and attempted fixes which were rejected as incorrect by the judge ("Bad Fix").
Models were all given the same prompt, tools, text chunks, and number of turns.
Here we compare the overall success rate with other dimensions like speed and cost. Only the best reasoning variant of each model is shown by default; toggle off "Best Variant Only" to see all data points.
This two-dimensional comparison gives the full picture between performance and cost, depending on what is important to you. Tool Call Efficiency measures how surgical the model was with its changes.
Use "Highlight Dominant Models" to see which models dominate on the axis you're viewing.
ErrataBench tests models against a dataset of English text from various sources (literature, law, technical manuals). Each source has been altered with a wide variety of writing errors, across several categories . These errors are tracked in a corresponding file.
Source text is altered using a corruptor model + review model pair; the corruptor is instructed to make small changes to the text to insert errors of specific types, sampled from the error category taxonomy . The review model reviews the changes as a second opinion before they are accepted into the dataset.
The benchmark runs a simple agent loop over each altered text, showing the model chunks of up to 2000 words at a time while keeping whole paragraphs together. The model is simply asked to proofread the text carefully, fixing errors while preserving meaning and details. The model is not given any specifics about what types of errors are hiding in the text or how many of them there are.
The agent loop gives the model up to 3 turns with each chunk, during which the model can modify that text using simple tools: find_and_replace (to make surgical changes) and replace_paragraph (for wider rewrites). If the model changes the text to what it was originally, the error is immediately counted as fixed. If the model changes it to something else, an LLM judge is used to decide if that alternative is also correct.
The headline Fix rate score is an equal-weighted per-dataset mean, not a pooled average over all runs. For each source dataset d , ErrataBench averages the successful-run quality scores for that dataset to get mean(q_d) , then averages those dataset-level means across datasets. That keeps datasets with more successful repeats from counting more heavily than datasets with fewer.
The other scatterplot metrics are derived similarly. Cost Efficiency is corrected issues per USD at the run level, then averaged by dataset and across datasets. Speed is corrected issues per minute of run time, aggregated the same way. Tool Call Efficiency is shown as issues fixed per 100 benchmark-scoped tool characters, which is an inverted display of the underlying tool-chars-per- resolved-issue metric. Turns Used is the mean average turns per chunk across datasets.
The scatterplot's consistency metric is median_d range(q_d) , the median per-dataset range of quality. For each dataset, ErrataBench measures the spread between the highest and lowest successful-run quality score, then takes the median of those dataset-level ranges. In the UI this appears as Consistency . Lower values mean the model's proofreading quality is more consistent from run to run, while the median makes the metric less sensitive to a few noisy source texts.
Source code, dataset, and results are available on GitHub . Results can be reproduced using an OpenRouter API key. The repository also includes a tool for creating your own dataset using any source text, making it easy to run this benchmark against new samples.
The benchmark has two primary inputs which determine how the agent loop works: the chunk size and number of turns per chunk. Changing these values can produce different results - specifically, larger chunks and fewer turns favors models that do less reasoning. High-reasoning models tend to get lost when presented larger chunks and fewer turns to edit them with. The data published here was generated by running the benchmark with --samples 3 --chunk-size 2000 --max-turns-per-chunk 3 , which was decided on as a reasonable general setting.
The full list of error categories used by ErrataBench can be found below.
Orthography and word form Typographical errors : fat-finger typo omission insertion transposition duplication
Spelling and casing : nonword spelling error real-word spelling error capitalization error spacing error hyphenation or compounding error
Morphology and inflection : incorrect pluralization incorrect conjugation wrong word form
Lexical choice and confusability Word selection : misused word homophone confusion homograph or near-neighbor confusion malapropism collocation error
Mishearing and reinterpretation : mondegreen eggcorn
Grammar and syntax Agreement : subject-verb agreement pronoun agreement
Verb system : incorrect tense auxiliary or modal error
Clause and sentence structure : sentence fragment run-on, fused sentence, or comma splice word order error missing word extra word parallelism error relative-clause error coordination or subordination error
Function words and reference : article or determiner error preposition error pronoun-case error pronoun-reference error dangling or misplaced modifier
Polarity and comparison : double negative negation-scope error comparative or superlative error countability error
Punctuation and boundaries punctuation error apostrophe error quotation or delimiter error sentence-boundary error
Semantics, discourse, and style Semantics : semantic anomaly referential ambiguity
Discourse : discourse coherence error cross-sentence tense inconsistency
Style and register : wordiness or redundancy register mismatch awkward phrasing nonstandard dialect or colloquial form
All model variants Lexical Choice And Confusability Idiomaticity
Claude Fable 5 (Medium) 15 / 15 ( 100.0% ) Claude Opus 4.6 (High) 15 / 15 ( 100.0% ) Claude Opus 4.6 (Low) 15 / 15 ( 100.0% ) Claude Opus 4.6 (Medium) 15 / 15 ( 100.0% ) Gemini 3.1 Pro Preview (High) 15 / 15 ( 100.0% ) Claude Fable 5 (High) 14 / 15 ( 93.3% ) Claude Fable 5 (Low) 14 / 15 ( 93.3% ) Claude Opus 4.6 (None) 13 / 15 ( 86.7% ) Claude Opus 4.8 (Low) 13 / 15 ( 86.7% ) Claude Opus 4.7 (Low) 12 / 15 ( 80.0% ) Full list Punctuation And Boundaries Sentence Boundary Error
Claude Fable 5 (Low) 30 / 30 ( 100.0% ) GPT 5.5 (High) 30 / 30 ( 100.0% ) GPT 5.5 (Medium) 30 / 30 ( 100.0% ) Claude Fable 5 (High) 29 / 30 ( 96.7% ) Gemini 3.1 Pro Preview (Low) 29 / 30 ( 96.7% ) Claude Fable 5 (Medium) 28 / 30 ( 93.3% ) GPT 5.4 Mini (High) 27 / 30 ( 90.0% ) Claude Opus 4.7 (None) 26 / 30 ( 86.7% ) Claude Opus 4.7 (High) 25 / 30 ( 83.3% ) Claude Opus 4.7 (Medium) 25 / 30 ( 83.3% ) Full list Semantics Discourse And Style Discourse
Claude Opus 4.6 (High) 12 / 15 ( 80.0% ) Claude Opus 4.6 (Low) 12 / 15 ( 80.0% ) Claude Opus 4.6 (Medium) 12 / 15 ( 80.0% ) Claude Opus 4.7 (High) 12 / 15 ( 80.0% ) Claude Opus 4.7 (Low) 12 / 15 ( 80.0% ) Claude Opus 4.7 (Medium) 12 / 15 ( 80.0% ) Claude Opus 4.7 (None) 12 / 15 ( 80.0% ) Claude Opus 4.8 (High) 12 / 15 ( 80.0% ) Claude Opus 4.8 (Medium) 12 / 15 ( 80.0% ) Claude Opus 4.8 (None) 12 / 15 ( 80.0% ) Full list Grammar And Syntax Function Words And Reference
Claude Fable 5 (Low) 112 / 123 ( 91.1% ) Claude Fable 5 (High) 111 / 123 ( 90.2% ) GPT 5.5 (Medium) 111 / 123 ( 90.2% ) Claude Fable 5 (Medium) 110 / 123 ( 89.4% ) GPT 5.5 (High) 110 / 123 ( 89.4% ) GPT 5.4 Mini (High) 108 / 123 ( 87.8% ) GPT 5.5 (Low) 107 / 123 ( 87.0% ) Claude Opus 4.7 (None) 105 / 123 ( 85.4% ) Gemini 3.1 Pro Preview (Medium) 103 / 123 ( 83.7% ) Claude Opus 4.6 (High) 102 / 123 ( 82.9% ) Full list Punctuation And Boundaries Quotation Or Delimiter Error
Claude Fable 5 (Low) 31 / 33 ( 93.9% ) GPT 5.5 (Low) 31 / 33 ( 93.9% ) Claude Fable 5 (High) 30 / 33 ( 90.9% ) Claude Fable 5 (Medium) 30 / 33 ( 90.9% ) Claude Opus 4.7 (None) 30 / 33 ( 90.9% ) GPT 5.4 (Medium) 30 / 33 ( 90.9% ) GPT 5.5 (None) 30 / 33 ( 90.9% ) Gemini 3 Flash Preview (High) 29 / 33 ( 87.9% ) GPT 5.4 (High) 29 / 33 ( 87.9% ) GPT 5.4 (None) 29 / 33 ( 87.9% ) Full list Semantics Discourse And Style Semantics
GPT 5.5 (High) 75 / 78 ( 96.2% ) GPT 5.5 (Medium) 75 / 78 ( 96.2% ) Claude Fable 5 (High) 74 / 78 ( 94.9% ) Claude Fable 5 (Low) 74 / 78 ( 94.9% ) Claude Fable 5 (Medium) 73 / 78 ( 93.6% ) Claude Opus 4.6 (None) 72 / 78 ( 92.3% ) Claude Opus 4.6 (Medium) 71 / 78 ( 91.0% ) GPT 5.4 (Medium) 71 / 78 ( 91.0% ) Gemini 3.1 Pro Preview (Medium) 70 / 78 ( 89.7% ) Glm 5.1 (Low) 70 / 78 ( 89.7% ) Full list Punctuation And Boundaries Punctuation Error
GPT 5.4 (Medium) 15 / 15 ( 100.0% ) GPT 5.4 Mini (Medium) 15 / 15 ( 100.0% ) GPT 5.5 (High) 15 / 15 ( 100.0% ) GPT 5.5 (Low) 15 / 15 ( 100.0% ) GPT 5.5 (Medium) 15 / 15 ( 100.0% ) GPT 5.5 (None) 15 / 15 ( 100.0% ) Gemini 3 Flash Preview (High) 14 / 15 ( 93.3% ) Gemini 3.1 Pro Preview (High) 14 / 15 ( 93.3% ) GPT 5.4 Mini (High) 14 / 15 ( 93.3% ) Grok 4.20 (High) 14 / 15 ( 93.3% ) Full list Grammar And Syntax Clause And Sentence Structure
Claude Fable 5 (High) 243 / 258 ( 94.2% ) Gemini 3.1 Pro Preview (High) 242 / 258 ( 93.8% ) GPT 5.5 (High) 241 / 258 ( 93.4% ) Claude Fable 5 (Low) 240 / 258 ( 93.0% ) GPT 5.4 (Medium) 240 / 258 ( 93.0% ) GPT 5.5 (Low) 240 / 258 ( 93.0% ) Claude Fable 5 (Medium) 239 / 258 ( 92.6% ) GPT 5.5 (Medium) 239 / 258 ( 92.6% ) Gemini 3.1 Pro Preview (Low) 238 / 258 ( 92.2% ) GPT 5.4 Mini (Medium) 238 / 258 ( 92.2% ) Full list Grammar And Syntax Agreement
Gemini 3 Flash Preview (High) 66 / 66 ( 100.0% ) Claude Fable 5 (High) 63 / 66 ( 95.5% ) Claude Fable 5 (Medium) 63 / 66 ( 95.5% ) Claude Fable 5 (Low) 62 / 66 ( 93.9% ) Gemini 3.1 Pro Preview (High) 62 / 66 ( 93.9% ) GPT 5.5 (Medium) 62 / 66 ( 93.9% ) GPT 5.5 (High) 62 / 66 ( 93.9% ) GPT 5.5 (Low) 62 / 66 ( 93.9% ) Claude Opus 4.8 (None) 60 / 66 ( 90.9% ) Glm 5.1 (High) 60 / 66 ( 90.9% ) Full list Grammar And Syntax Polarity And Comparison
Claude Fable 5 (High) 54 / 57 ( 94.7% ) Claude Fable 5 (Low) 54 / 57 ( 94.7% ) Claude Fable 5 (Medium) 54 / 57 ( 94.7% ) Claude Opus 4.6 (None) 54 / 57 ( 94.7% ) Claude Opus 4.7 (High) 54 / 57 ( 94.7% ) Claude Opus 4.7 (Low) 54 / 57 ( 94.7% ) Claude Opus 4.7 (Medium) 54 / 57 ( 94.7% ) Claude Opus 4.7 (None) 54 / 57 ( 94.7% ) Claude Opus 4.8 (High) 54 / 57 ( 94.7% ) Claude Opus 4.8 (Low) 54 / 57 ( 94.7% ) Full list Load more categories Revise Company Product updates

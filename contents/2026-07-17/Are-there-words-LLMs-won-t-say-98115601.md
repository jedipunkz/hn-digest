---
source: "https://samermakes.com/blog/WCS/"
hn_url: "https://news.ycombinator.com/item?id=48946384"
title: "Are there words LLMs won't say?"
article_title: "Are there words LLMs won’t say? — Samer Makes"
author: "SamerIA"
captured_at: "2026-07-17T13:06:04Z"
capture_tool: "hn-digest"
hn_id: 48946384
score: 1
comments: 0
posted_at: "2026-07-17T12:09:56Z"
tags:
  - hacker-news
  - translated
---

# Are there words LLMs won't say?

- HN: [48946384](https://news.ycombinator.com/item?id=48946384)
- Source: [samermakes.com](https://samermakes.com/blog/WCS/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T12:09:56Z

## Translation

タイトル: LLM が言わない言葉はありますか?
記事のタイトル: LLM が言わない言葉はありますか? — サマー・メイクス
説明: LLM サンプリングにおける温度、top-k、top-p、min-p、およびワード カバレッジ スコアの視覚的な説明。

記事本文:
記事にスキップ
Samer Makes / ブログ / WCS
LLM が言ってはいけない言葉はありますか?
LLM の語彙にはあらゆる英単語が含まれていますが、それらの単語の中には出力にまったく表示されないものもあります。 LLM によって生成されたテキストは、人間の文章に比べて語彙の多様性が明らかに低くなります。 GING の研究者らは、たとえ特定の文脈に適切な場合でも、LLM が特定の単語を生成しない可能性がある考えられる理由の 1 つを調査することに着手しました。
LLM は次のトークンを予測するときに、前のコンテキストを読み取り、語彙内のすべてのトークンのロジットを生成します。トークンには、単語全体、単語の一部、句読点、スペース、またはその他のテキストの断片を指定できます。たとえば、Llama 3 には 128,256 個の可能なトークンの語彙があります。次のトークンを推論するとき、これらの各オプションにはロジットが割り当てられます。
ロジットは確率ではなく、生の数値スコアです。ソフトマックスと呼ばれる関数は、数学定数 e を各ロジットの累乗にし、その値をすべてのべき乗されたロジットの合計で割ることにより、ロジットを確率に変換します。
Softmax は、2 つの数値の比率ではなく、2 つの数値間の生の数値ギャップに基づいてスケールします。たとえば、 4:2 と 2:1 は同じ比率を表しますが、絶対的なギャップは異なります。 4 - 2 = 2 であるのに対し、 2 - 1 = 1 です。
[4, 2] にソフトマックスを適用すると、約 [0.881, 0.119] が得られます。
[2, 1] にソフトマックスを適用すると、約 [0.731, 0.269] が得られます。
比率が同じであっても、絶対ギャップが大きいため、ソフトマックスの適用後はより高い値がより支配的になります。
温度は推論中に調整される変数であり、通常は創造性パラメータとして扱われます。温度を上げると出力の変動が大きくなり、温度を下げるとモデルの出力がより決定的になります。 T = 0 で、モデルは貪欲サンプリングに入り、常に

最高のロジットを持つトークン。
推論中、温度値は各ロジットを T で割ることによってロジットをスケーリングします。
プロンプトが次のようになったとします。「たくさんのことが起こりました!」カンファレンスの初日は大変な一日でした。 interest が logit 4 を取得し、long が logit 2 を取得する場合、softmax はinteresting を大きく優先します。 T = 2 に設定すると、それらのロジットは [2, 1] になります。ギャップが縮まり、確率曲線が平坦になり、2 つのオプション間の可能性の差が減少します。
温度を上げると、モデルは最上位の選択に集中しなくなります。最も可能性の高いトークンの支配力が低下し、より低い確率のトークンが選択される可能性が高くなります。
温度によりソフトマックス前のギャップが変化する
温度スライダーを動かします。ロジットは同じままですが、スケーリングされたギャップは変化し、ソフトマックスは 2 つのトークンの確率を近づけたり、遠ざけたりします。
ただし、トークンは、最初にサンプリング フィルターを通過した場合にのみ選択できます。
フィルタリングする前は、たとえその確率が小さいとしても、すべてのトークンにはある程度の確率があります。サンプリング フィルターは、その巨大なリストを、許可されたトークンのより小さなセットに削減します。次に、モデルは確率によって重み付けされて、残りのトークンからランダムに選択します。
一般的な 3 つのサンプリング フィルターは、 top-k 、 top-p 、および min-p です。
top-k の場合、k の値、たとえば 5 を選択します。モデルは確率によってトークンをランク付けし、最初の 5 つを保持し、ランク 5 未満のものはすべて破棄します。最後のトークンは、生き残ったグループからサンプリングされます。
Top-k は固定数のトークンを保持します
kを変更します。カットオフは、実際の確率値に関係なく、ランクによって移動します。
核サンプリング とも呼ばれる top-p の場合、確率しきい値 (たとえば 0.9) を選択します。モデルは、最も可能性の高いトークンから開始し、次に最も可能性の高いトークンを追加し、保持されているトークンが少なくとも合計されるまでその次のトークンを追加します。

確率質量の st 90%。そのカットオフを超えたトークンは削除されます。
したがって、一番上のトークンの確率が 40% の場合、保持されるセットは 40% から始まります。 2 番目のトークンが 20% の場合、保持されるセットは 60% に増加します。 3 番目が 12% であれば、72% に増加します。モデルは、現在の合計が上位のしきい値に達するまで、ランク付けされたリストを下降し続けます。その後、生存確率は繰り込みされて合計 100% になります。
Top-p は累計が p に達するまでトークンを保持します。
緑色のセグメントは残ります。ミュートされた赤色のセグメントはカットオフを超えています。黄色の線は p しきい値です。
min-p の場合、カットオフは最も可能性の高いトークンの確率に基づきます。 min-p が 0.05 で、最上位トークンの確率が 40% の場合、カットオフは次のようになります。
0.05 × 40% = 2%
2% 未満のすべてのトークンは破棄されます。これにより、min-p が動的になります。トップトークンが非常に強い場合、カットオフはより厳しくなります。上部のトークンが弱い場合、カットオフは緩和されます。
Min-p は、最上位のトークンと比較してしきい値を超えたトークンを維持します。
琥珀色の線がカットオフです。緑色のバーは存続します。ミュートされた赤いバーが削除されます。
温度フィルターとサンプリング フィルターが適用された後、モデルは生き残ったセットから 1 つのトークンを選択します。
トークンはモデルの分布内で実際の確率を持ち、フィルタリング後に出力できなくなる可能性があります。サンプラーがそれを削除すると、モデルはそのステップでそれを選択できなくなります。
私たちは最近、一般的なサンプラー設定で比較的一般的な英単語が除外され、LLM による生成が不可能になるかどうかを調査したプレプリントをリリースしました。まず、単語の頻度リストに 10,000 から 40,000 の頻度範囲を定義しました。これらの単語は、人間の作成者が自然に使用できるほど一般的です。 10k ランクに近い例には、毒、民族性、観光客、疲労、コイル、および消費が含まれます。まれな例として、アプローチします。

40,000 ランクの例には、本能的に、スティレット、イベントフル、ロブスター、口ひげ、石炭、多肉植物が含まれます。 We randomly selected 100 words from this range; the full list is available here .
私たちは、人間が作成したテキストの大規模なコーパスである PG19 内でこれらの単語を検索しました。単語ごとに、その周囲のコンテキストを抽出し、その単語がそのコンテキストで意味をなすかどうかを LLM で検証しました。これは OCR または印刷エラーを考慮しており、1,000 個のコンテキストを調べたので自動化する必要がありました。
中頻度のターゲット単語の選択から、サンプラー フィルターが単語をコンテキスト内で保持するかどうかの監査まで、単語カバレッジ スコア パイプラインの概要。
次に、複数のオープンウェイト LLM にわたるコンテキスト内の各ターゲット単語に割り当てられたロジットをテストしました。このデータを使用して、単語がさまざまなサンプラー設定の下で最初の選択フェーズを乗り越えて実行可能な出力オプションとして残るかどうかを計算しました。これをバイナリ到達可能性式 R θ (w, c) ∈ {0, 1} を使用して定義し、コーパス全体の単語カバレッジ スコアに集計しました。
サンプラーの制約が大幅に緩和され、より多くのトークンがサンプリング フィルターを通過できるようになると、ほとんどのモデルが少なくとも 1 つのコンテキストで単語の大部分を予測しました。これにより、モデルが語彙を備えており、人間の文脈でターゲットの単語を生成できることが確認されます。
ただし、推奨設定では、テストされたすべてのモデルは、すべてのコンテキストにわたって一部の単語を予測できませんでした。つまり、デフォルト設定では、特定のモデルのテストされたコンテキストで一部の単語にアクセスできませんでした。デフォルト設定では、一度も予測されなかった単語の割合は 20% ～ 45% でした。これらの結果のインタラクティブなビジュアライザは、ここで参照できます。
温度 ( T ) の影響も評価しました。トップサンプルの下

つまり、T を 0.7 から 1.0 に上げると WCS が大幅に増加しましたが、T が 1.5 に近づくにつれて利益は減少しました。 top-k サンプリングは純粋にランク順に基づいているため、フィルタリングされるトークンのセットは温度の影響をまったく受けません。 Min-p は、top-p よりも温度スケーリングの影響を受けにくいことが判明しました。 Llama-3.1-8B の WCS を 32.8% から 57.2% に増加するには、T を 0.7 から 1.5 にスケーリングする必要がありましたが、top-p は 0.7 から 1.0 にスケーリングするだけで同様の絶対ゲインを達成します。
この論文では、サンプラー設定が語彙の多様性に及ぼす影響を定量的に測定する方法を導入できたと考えています。私たちが読んでいる単語の多くを LLM が書くようになっているため、LLM の影響に基づいて、時間の経過とともに言語のニュアンスや幅広さを失わないようにすることが重要です。この影響はすでに目に見えて測定可能であり、学術出版における特定の単語の人気は LLM の傾向に応じて増減します。 Geng & Trotta、2025 を参照。
主要な LLM 採用期間の前後で学術抄録の単語頻度が変化しており、モデルに関連した言語の傾向が人間向けのテキストで目に見えるようになる可能性があることを示しています。出典: Geng & Trotta、2025 年。
GING では、LLM の限界と機能、科学、研究、心理言語学、教育における AI の応用を研究しています。協力者を探している場合、または議論したいアイデアがある場合は、お気軽にお問い合わせください。
私たちが次に公開するものや、AI、言語、言語学習に関する私の考えに興味がある場合は、メーリング リストに登録するか、RSS 経由でフォローしてください。
新しい投稿が受信箱に配信されるか、RSS 経由でフォローします。
更新: この記事を書いて以来、新しいコーパス FineWeb を使用して監査を再実行し、それぞれ 50 のコンテキストを持つ 200 単語に拡張しました。また、モデルには h の後に全段落の続きを書いてもらいました。

人間が書いた文脈を分析し、それらの継続部分の語彙の多様性を単語カバレッジ スコアと相関させました。結果は、WCS が高いほど、LLM で生成されたテキストの語彙の多様性がより大きくなると予測されることを示しています。

## Original Extract

A visual explanation of temperature, top-k, top-p, min-p, and the Word Coverage Score in LLM sampling.

Skip to article
Samer Makes / Blog / WCS
Are there words LLMs won’t say?
LLMs have every English word in their vocabulary, but some of those words may never appear in their output. LLM-generated text is measurably less lexically diverse than human writing. The researchers at GING set out to investigate one possible reason why an LLM might never produce a specific word, even if it is appropriate for a given context.
When an LLM predicts the next token, it reads the preceding context and produces a logit for every token in its vocabulary. A token can be a whole word, a piece of a word, punctuation, a space, or any other text fragment. For example, Llama 3 has a vocabulary of 128,256 possible tokens. When inferring the next token, each of these options is assigned a logit.
The logit isn’t a probability but a raw numerical score. A function called softmax turns the logits into probabilities by raising the mathematical constant e to the power of each logit, then dividing that value by the sum of all the exponentiated logits.
Softmax scales based on the raw numerical gap between two numbers, not their ratio. For example, 4:2 and 2:1 represent the same ratio, but the absolute gaps are different: 4 - 2 = 2 , while 2 - 1 = 1 .
Applying softmax to [4, 2] gives approximately [0.881, 0.119] .
Applying softmax to [2, 1] gives approximately [0.731, 0.269] .
Even though the ratios are identical, the larger absolute gap makes the higher value much more dominant after softmax is applied.
Temperature is a variable adjusted during inference, usually treated as a creativity parameter. Increasing the temperature makes the output more varied, while decreasing it makes the model's outputs more deterministic. At T = 0 , the model enters greedy sampling and always chooses the token with the highest logit.
During inference, the temperature value scales the logits by dividing each logit by T :
Suppose the prompt is: A lot of things happened! The first day of the conference was quite . If interesting gets logit 4 and long gets logit 2, softmax heavily favors interesting . If we set T = 2 , those logits become [2, 1] . The gap narrows, and the probability curve is flattened, reducing the disparity in likelihood between the two options.
Raising the temperature makes the model less concentrated on its top choice. The most likely token becomes less dominant, giving lower-probability tokens a higher chance of selection.
Temperature changes the gap before softmax
Move the temperature slider. The logits stay the same, but the scaled gap changes, and softmax gives the two tokens closer or more distant probabilities.
But a token can only be chosen if it survives the sampling filter first.
Before filtering, every token has some probability, even if that probability is tiny. Sampling filters cut that huge list down to a smaller set of allowed tokens. Then the model randomly chooses from the remaining tokens, weighted by their probabilities.
The three common sampling filters are top-k , top-p , and min-p .
For top-k , we choose a value of k, for example 5. The model ranks tokens by probability, keeps the first 5, and discards everything below rank 5. The final token is sampled from the surviving group.
Top-k keeps a fixed number of tokens
Change k. The cutoff moves by rank, regardless of the actual probability values.
For top-p , also called nucleus sampling , we choose a probability threshold, for example 0.9. The model starts with the most likely token, then adds the next most likely token, then the next, until the kept tokens together account for at least 90% of the probability mass. Tokens beyond that cutoff are removed.
So if the top token has probability 40%, the kept set starts at 40%. If the second token has 20%, the kept set grows to 60%. If the third has 12%, it grows to 72%. The model keeps walking down the ranked list until the running total reaches the top-p threshold. After that, the surviving probabilities are renormalized so they add up to 100%.
Top-p keeps tokens until the running total reaches p
Green segments survive. Muted red segments are beyond the cutoff. The amber line is the p threshold.
For min-p , the cutoff is based on the probability of the most likely token. If min-p is 0.05 and the top token has probability 40%, then the cutoff is:
0.05 × 40% = 2%
Every token below 2% is discarded. This makes min-p dynamic. When the top token is very strong, the cutoff becomes stricter. When the top token is weaker, the cutoff relaxes.
Min-p keeps tokens above a threshold relative to the top token
The amber line is the cutoff. Green bars survive; muted red bars are removed.
After temperature and sampling filters are applied, the model chooses one token from the surviving set.
A token can have a real probability in the model’s distribution and still become impossible to output after filtering. Once the sampler removes it, the model cannot choose it in that step.
We recently released a preprint investigating whether common sampler settings filter out relatively common English words, rendering them impossible for LLMs to generate. We began by defining a frequency range of 10k to 40k on a word frequency list. These words are common enough that a human author uses them naturally. Examples closer to the 10k rank include poison , ethnicity , tourists , fatigue , coil , and consuming . On the rarer end, approaching the 40k rank, examples include instinctively , stiletto , eventful , lobsters , moustache , coals , and succulent . We randomly selected 100 words from this range; the full list is available here .
We searched for these words within PG19, a large corpus of human-authored texts. For each word, we extracted its surrounding context and verified with an LLM that the word makes sense in that context. This accounts for OCR or printing errors, and was necessary to automate because we looked at 1,000 contexts.
Overview of the Word Coverage Score pipeline, from selecting mid-frequency target words to auditing whether sampler filters preserve them in context.
We then tested the logits assigned to each target word in its context across multiple open-weights LLMs. Using this data, we calculated whether the word would survive the initial selection phase under various sampler settings to remain a viable output option. We defined this using a binary reachability formula R θ (w, c) ∈ {0, 1} , which we aggregated into a Word Coverage Score across the corpus:
When sampler constraints are drastically loosened, allowing more tokens to survive sampling filters, most models predicted the majority of the words in at least one context. This confirms the models possess the vocabulary and are capable of producing the target word in the human context.
However, at recommended settings, every model tested failed to predict some of the words across all contexts, meaning that some words were not accessible in any tested context for certain models at the default settings. The percentage of words not predicted a single time at default settings ranged from 20% to 45%; interactive visualizers of these results can be viewed here .
We also evaluated the impact of temperature ( T ). Under top-p sampling, raising T from 0.7 to 1.0 yielded a significant increase in WCS, with diminishing returns as T approached 1.5. Because top-k sampling is purely rank-order based, the set of tokens it filters is entirely unaffected by temperature. Min-p proved less sensitive to temperature scaling than top-p; scaling T from 0.7 to 1.5 was required to increase the WCS for Llama-3.1-8B from 32.8% to 57.2%, whereas top-p achieves similar absolute gains merely scaling from 0.7 to 1.0.
We believe that with our paper we’ve introduced a way to quantitatively measure the effect of sampler settings on lexical diversity. As LLMs write more of the words we read, it’s important that we don’t lose the nuances and breadth of the language over time based on LLM influence. The effect of this is already visible and measurable, with the popularity of certain words in academic publishing rising and dropping with LLM trends; see Geng & Trotta, 2025 .
Word-frequency shifts in academic abstracts around major LLM adoption periods, illustrating that model-associated language trends can become visible in human-facing text. Source: Geng & Trotta, 2025 .
At GING we research the limits and capabilities of LLMs, the applications of AI in science, research, psycholinguistics, and education. If you’re looking for collaborators or have ideas you want to discuss, feel free to reach out.
If you’re interested in whatever we publish next or my thoughts on AI, languages, and language learning, you can sign up for the mailing list or follow via RSS.
Get new posts delivered to your inbox, or follow via RSS.
Update: Since writing this post, we’ve rerun the audit using a new corpus, FineWeb, and expanded it to 200 words with 50 contexts each. We also had the models write full-paragraph continuations after the human-written contexts, then correlated the lexical diversity of those continuations with the Word Coverage Score. The results show that a higher WCS predicts greater lexical diversity in LLM-generated text.

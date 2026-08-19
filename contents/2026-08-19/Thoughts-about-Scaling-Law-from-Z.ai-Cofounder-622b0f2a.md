---
source: "https://xcancel.com/jietang/status/2089941544581403107#m"
hn_url: "https://news.ycombinator.com/item?id=49358643"
title: "Thoughts about Scaling Law (from Z.ai Cofounder)"
article_title: "jietang (@jietang): \"Thoughts About Scaling Law\nScaling, but not only of parameters. Every model release now ends with the same question: how many parameters? It isn't a question that can be answered on its own. Parameter count is only meaningful alongside three others — how much data you have, wher\n[truncated]"
image: ""
author: "theanonymousone"
captured_at: "2026-08-19T09:24:14Z"
capture_tool: "hn-digest"
hn_id: 49358643
score: 2
comments: 0
posted_at: "2026-08-19T08:24:19Z"
tags:
  - hacker-news
  - translated
---

# Thoughts about Scaling Law (from Z.ai Cofounder)

- HN: [49358643](https://news.ycombinator.com/item?id=49358643)
- Source: [xcancel.com](https://xcancel.com/jietang/status/2089941544581403107#m)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T08:24:19Z

## Translation

タイトル: スケーリングの法則についての考え (Z.ai 共同創業者より)
記事タイトル: jietang (@jietang): 「スケーリング法についての考え」
スケーリングですが、パラメータだけではありません。現在、すべてのモデルのリリースは同じ質問で終わります: パラメーターはいくつですか?それは単独で答えられる質問ではありません。パラメータ数は、他の 3 つのパラメータと並んでのみ意味を持ちます。つまり、どれだけのデータがあるのか、どこにあるのかということです。
[切り捨てられた]
説明: スケーリングの法則についての考え
スケーリングですが、パラメータだけではありません。現在、すべてのモデルのリリースは同じ質問で終わります: パラメーターはいくつですか?それは

記事本文:
jietang (@jietang): 「スケーリング法についての考え」
スケーリングですが、パラメータだけではありません。現在、すべてのモデルのリリースは同じ質問で終わります: パラメーターはいくつですか?それは単独で答えられる質問ではありません。パラメーター数は、他の 3 つの要素、つまり、保有するデータの量、コンピューティングをどこに費やす予定か、誰がどのような条件でモデルを実行するかという要素と並んでのみ意味を持ちます。
現場はこれを苦労して学びました。カプランら。 (2020) は、パラメータをデータよりも速く増加させるよう全員に指示する指数 (およそ 2.7:1) を当てはめ、業界は GPT-3、Gopher、MT-NLG に従った。ホフマンら。 (2022) は 400 のモデルにわたって実験を再検討し、計算に最適な分割はパラメータあたり 20 トークンに近く、十分な計算があれば 2 つは乖離するのではなく同じ割合で増加するはずであることを発見しました。以前の適合における誤差は、計算量が桁違いに増加するにつれて悪化しました。そのため、その世代の最大のモデルが最も誤って割り当てられていました。 1兆パラメータのラウンドは、今にして思えば、分野全体が協力して回り道をし、その後逆転したものだった。
チンチラも終わりではありませんでした。一度トレーニングして評価するモデルのトレーニング コンピューティングを最適化しました。現在、モデルは 1 日に何十億回も呼び出され、推論が生涯コストの大半を占めています。推論を目的に置き、はるかに長い時間トレーニングされた小さなモデルに向けた最適な動き、つまり意図的なオーバートレーニングです。これは、Llama-2-7B と Gemma-2-9B がパラメータあたり約 290 トークンと 889 トークンで行っていたことです。
スパーシティは再びターゲットを移動しました。 MoE モデルでは、2 つの量を区別する必要があります。合計パラメーターは、モデルが保持できる量 (知識、事実、ロングテール) を大まかに制御し、アクティブ化されたパラメーターと有効深度は、モデルがどの程度の量を保持できるかを大まかに制御します。

因果関係の連鎖がばらばらになるまでに、何段階まで続くことができるか、考えることができる。密度の高い 20:1 の比率は移行しません。そして、その比率は決して単一の数字ではありません：Roberts et al. (2025) パラメータごとの最適なトークン数は t であることを確認します。
[切り捨てられた]
AI に焦点を当てた初の分散型先物取引所を活用して、Anthropic、メモリ株、H100 の価格を取引します。

## Original Extract

Thoughts About Scaling Law
Scaling, but not only of parameters. Every model release now ends with the same question: how many parameters? It isn

jietang (@jietang): "Thoughts About Scaling Law
Scaling, but not only of parameters. Every model release now ends with the same question: how many parameters? It isn't a question that can be answered on its own. Parameter count is only meaningful alongside three others — how much data you have, where you intend to spend your compute, and who will run the model, under what conditions.
The field learned this the hard way. Kaplan et al. (2020) fit an exponent that told everyone to grow parameters faster than data — roughly 2.7:1 — and the industry complied: GPT-3, Gopher, MT-NLG. Hoffmann et al. (2022) redid the experiment across four hundred models and found the compute-optimal split is closer to 20 tokens per parameter, and that with sufficient compute the two should grow at the same rate rather than drifting apart. The error in the earlier fit compounded with every order of magnitude of compute, which is why the largest models of that generation were the most misallocated. The trillion-parameter round was, in retrospect, a detour the whole field took together and then reversed.
Chinchilla wasn't the end either. It optimized training compute for models that would be trained once and evaluated. Today a model is called billions of times a day and inference dominates lifetime cost. Put inference into the objective and the optimum moves toward smaller models trained far longer — deliberate over-training, which is what Llama-2-7B and Gemma-2-9B were doing at roughly 290 and 889 tokens per parameter.
Sparsity moved the target again. In a MoE model two quantities have to be kept apart: total parameters govern roughly how much the model can hold — knowledge, facts, the long tail — while activated parameters and effective depth govern roughly how far it can think, how many steps of a causal chain it can carry before it comes apart. A dense 20:1 ratio does not transfer. And the ratio isn't a single number at all: Roberts et al. (2025) find the optimal tokens-per-parameter is t
[truncated]
Trade Anthropic, memory stocks, and H100 prices with leverage on the first AI-focused decentralized futures exchange.

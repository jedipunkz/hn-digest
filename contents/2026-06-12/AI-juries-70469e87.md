---
source: "https://alejo.ch/3he"
hn_url: "https://news.ycombinator.com/item?id=48507129"
title: "AI juries"
article_title: "AI juries"
author: "afc"
captured_at: "2026-06-12T18:45:58Z"
capture_tool: "hn-digest"
hn_id: 48507129
score: 2
comments: 2
posted_at: "2026-06-12T17:42:10Z"
tags:
  - hacker-news
  - translated
---

# AI juries

- HN: [48507129](https://news.ycombinator.com/item?id=48507129)
- Source: [alejo.ch](https://alejo.ch/3he)
- Score: 2
- Comments: 2
- Posted: 2026-06-12T17:42:10Z

## Translation

タイトル: AI 陪審員

記事本文:
これは私が書いているシリーズの一部です
生成AI。
状態: 撤回済み。このテキストは以下を前提としています
不当な仮定: 同じに対する複数の評価が、
同じ生成 AI は独立した陪審員と同等です。そうかもしれない
そうでない場合、結論は導き出されません。
N 人の専門家からなる陪審を想像してみてください
二値 (真/偽) の事実が真であるかどうかを判断しようとします。各専門家は、
独立確率 p
正しいということ。彼らは専門家なので、p > 0.5 であると仮定します – 彼らは少なくとも優れています
コイン投げよりも。
投票を求めることができます。各専門家に二者択一の答えを述べてもらい、
最も人気のある答えを採用します。その確率を次のように呼びましょう。
陪審員の過半数は正しいです。この「陪審の正確さ」の確率
N に応じて劇的に増加する
成長します。
たとえば、N = 25 人の専門家からなる陪審の場合、それぞれの専門家は次のような可能性があります。
確率 p = 0.75 で正しい、多数派である確率 j
右はすでに 99.663% です!
以下の表は、この効果がどれほど強力かを示しています。それは陪審員を示しています
異なる陪審の精度 j
サイズ N (列) と個別
専門家の精度 p
(列):
以下は、ほぼ同じデータをプロットしたものです (x 軸は対数です)
スケール):
この「群衆の知恵」効果は、少なくともこの理想化された状況では機能します。
すべての専門家が独立しているシナリオ。
興味深い裏返しがあります。「専門家」の方がそうする可能性が高い場合です。
間違っている ( p < 0.5 )、陪審員のパフォーマンス
急落する。多数決は誤差を増幅するだけであり、
精度が 0% (つまり、確実に間違っている)。これを入れ替えるとわかります
上記の説明における「正しい」と「間違っている」の定義。
このモデルは新しいものではありません。それはフランスの黎明期に遡ります
革命。 1785年にコンドルセ侯爵によって初めて提案されました。
そして彼のものとして知られています

陪審定理。
Dumb AI とソフトウェア革命で述べた
生成 AI の主な利点: 光速で動作することに加えて
(数日ではなく数秒で質問に回答します)、拡張可能です
即座に：
無限
サルはもはやランダムにキーを押すことはありません。ただ少し愚かなだけです。
しかし、この速度と即時スケーリングでは、違いは次のとおりです。
記念碑的な。
AI 陪審員に二者択一の質問に答えてもらうことができれば、
精度 p > 0.5 の場合、次のことができます。
リソースをスケーリングするだけで任意の高いパフォーマンスを達成できます。
より大きな陪審員を運営しています。
AI がコイントスよりもかろうじて優れている場合でも (たとえば、 p = 0.51 )、
j > 0.99 の陪審員を獲得する
ただし、かなり高価になる可能性があります ($p = 0.51% の場合、
当たるためだけに N = 5001 人の陪審員
j = 0.92 )。
ただし、上の表が示すように、エージェントの精度を向上させると、
p をたとえば 0.70 にすると、N = 25 ではすでに 98% を超えています。
精度。
そして、興味深いことに、陪審員の規模を拡大してもレイテンシには影響しません。の
質問は並行して実行できるため、合計時間が決まります
最も遅いエージェントによる。私はヘッジのような戦略があるのではないかと疑っている（例：N + M 人の陪審員を選出し、
まずはNから答えてください
回答者）が該当する可能性があります。
では、これはどのような種類の質問に役立つでしょうか?いくつか持っています
実用的で日常的な例を念頭に置いてください:
この単体テストは失敗します。実装は正しいですか（
テストされたプロパティ)?これを使用して次のステップを決定する予定です。
エージェントに実装の修正を依頼するか、エージェントに実装の修正を依頼してください。
テスト。
上記のバリエーション: このコードは特定の機能を実装していますか?
プロパティ（自然言語で記述）?
コードの変更を受け取りました (おそらく AI から、あるいは
人間;関係ありません）。特定のコーディングスタイルに準拠していますか
ガイドライン？陪審員の答えによって、別のエージェントを雇うかどうかが決まります

修正する
スタイルの問題。
この関数のドキュメント (docstring) は依然として正確ですか?
記述されているコードは何ですか？
新しいレシピを書きました。それは一つを支持しますか
私のレシピの原則は何ですか？
独自の AI 陪審員を作成できるツールが使えることに興奮しています。ベース
彼らの観察されたパフォーマンス (偽陽性/偽陰性の比率) に基づいて、私は次のことを行うことができます。
コンテキストを微調整する (p を増やす) か、陪審員の人数を調整する (p を増やす)
N )、正しいものが見つかるまで
最終的な精度とコストのバランス。
陪審が間違った判断を下すたびに、私は関与して裁判の舵取りをします。
プロセス。物事が正しく行われるたびに、分析の手間が省けます
自分自身のこと。
これにより、非常に具体的なトレードオフが生じます。つまり、お金を使う（より高額になる）ということです。
陪審員）時間を節約するため（手作業による介入が少なくなります）。

## Original Extract

This is part of a series I’m writing on
generative AI.
State: Withdrawn. This text is predicated on an
unjustified assumption: that multiple evaluations of the same prompt by
the same generative AI are equivalent to independent jurors. That may
not be true, in which case the conclusion doesn’t follow.
Imagine a jury of N experts
trying to decide if a binary (true/false) fact is true. Each expert has
an independent probability p
of being right. Since they are experts, assume p > 0.5 –they’re at least better
than a coin flip.
We can ask for a vote: have each expert state their binary answer and
take the most popular answer. Let’s call the probability that the
majority of the jury is right j . This “jury accuracy” probability
increases dramatically as N
grows.
For example, with a jury of N = 25 experts, each likely to be
right with a probability p = 0.75 , the probability j that the majority is
right is already 99.663%!
The table below shows how powerful this effect is. It shows the jury
accuracy j for different jury
sizes N (rows) and individual
expert accuracies p
(columns):
The following is roughly the same data in a plot (x axis in log
scale):
This “wisdom of the crowds” effect works, at least in this idealized
scenario where all experts are independent.
There’s an interesting flip side: if the “experts” are more likely to
be wrong ( p < 0.5 ), the jury’s performance
plummets. The majority vote just amplifies the error, converging towards
0% accuracy (i.e., being reliably wrong). You can see this by swapping
the definitions of “right” and “wrong” in the description above.
This model isn’t new; it goes back to the dawn of the French
revolution. It was first proposed in 1785 by the Marquis of Condorcet
and is known as his Jury Theorem.
In Dumb AI and the software revolution I mentioned
a key advantage of generative AI: in addition to working at light speed
(answering questions in seconds, not days), it can be scaled
instantly :
The infinite
monkey is no longer hitting keys at random –just somewhat stupidly.
But at this speed and with instant scaling, the difference is
monumental.
If you can get an AI juror to answer a binary question with an
accuracy p > 0.5 , you can
achieve arbitrarily high performance simply by scaling resources:
running bigger juries.
Even if your AI is barely better than a coin flip (say, p = 0.51 ), you can still
get a jury with j > 0.99
accuracy, though it can be quite expensive (with $p = 0.51%, you’d need
a jury of N = 5001 just to hit
j = 0.92 ).
But, as the table above shows, if you improve your agent’s accuracy
p to, say, 0.70, with N = 25 you’re already above 98%
accuracy.
And, interestingly, scaling the jury doesn’t impact latency. The
questions can be executed in parallel, so the total time is determined
by the slowest agent. I suspect strategies like hedging (e.g., run N + M jurors, take the
answer from the N first
responders) may be applicable.
So what kinds of questions can this be useful for? I have a few
practical, everday examples in mind:
This unit test fails. Is the implementation correct (regarding
the tested property)? I intend to use this to determine the next step:
ask an agent to fix the implementation, or ask it to fix the
test.
A variation of the above: Does this code implement a specific
property (described in natural language)?
I’ve received a code change (possibly from AI, possibly from a
human; it doesn’t matter). Does it adhere to a specific coding style
guideline? The jury’s answer determines if I run another agent to fix
the style issue.
Is this function’s documentation (docstring) still accurate for
the code it describes?
I have written a new recipe . Does it uphold one
of my principles for my recipes ?
I’m excited to have a tool where I can create my own AI juries. Based
on their observed performance (false positives/negatives ratios), I can
tweak the contexts (increase p ) or adjust the jury size (increase
N ), until I find the right
balance between final accuracy and cost.
Every time the jury gets things wrong, I get involved and steer the
process. Every time it gets things right, it saves me from analyzing
things myself.
This creates a very concrete trade-off: spend money (bigger
juries) to save time (less manual intervention) .

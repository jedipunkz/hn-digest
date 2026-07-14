---
source: "https://prompt-eval.com/state-of-prompt-quality/2026-q3"
hn_url: "https://news.ycombinator.com/item?id=48910829"
title: "We scored 1,018 real-world AI prompts. Robustness averaged 31/100"
article_title: "The Robustness Gap — State of Prompt Quality, Q3 2026 | PromptEval"
author: "Franciscoferr"
captured_at: "2026-07-14T19:03:21Z"
capture_tool: "hn-digest"
hn_id: 48910829
score: 1
comments: 0
posted_at: "2026-07-14T18:08:17Z"
tags:
  - hacker-news
  - translated
---

# We scored 1,018 real-world AI prompts. Robustness averaged 31/100

- HN: [48910829](https://news.ycombinator.com/item?id=48910829)
- Source: [prompt-eval.com](https://prompt-eval.com/state-of-prompt-quality/2026-q3)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T18:08:17Z

## Translation

タイトル: 1,018 個の現実世界の AI プロンプトをスコア化しました。堅牢性の平均値は 31/100
記事のタイトル: ロバストネス ギャップ — 迅速な品質の状態、2026 年第 3 四半期 |プロンプト評価
説明: 1,018 個の実際のプロンプトを 0 ～ 100 でスコア付けしました。平均: 54。しかし、話は 1 つの側面です。堅牢性は平均 31.5/100 であり、プロンプトの 96% において最も弱い部分です。 State of Prompt Quality の最初のエディション。

記事本文:
ロバストネス ギャップ — 迅速な品質の状態、2026 年第 3 四半期 | PromptEval プロンプト評価 ← プロンプト品質のライブ ベンチマーク状態 · 2026 年第 3 四半期版 ロバストネス ギャップ
固定ルーブリックを使用して、1,018 件の実際のプロンプトを 0 ～ 100 で採点しました。平均は 54 でした。しかし、結果はそうではありませんでした。その結果、ほぼすべてのプロンプトが同じ場所で失敗し、それは人々が磨く場所ではないことがわかりました。
ほとんどのプロンプトは、非常に特殊な点で平凡です
配布は災害の話ではありません。プロンプトの 3 分の 2 は 40 ～ 79 の間に到達します。つまり、機能的で、答えが得られる領域です。本当に壊れているのは 12% だけです (30 未満)。問題は上限です。わずか 10.5% が 75 点に達しています。このスコアは、繰り返し使用しても信頼できるプロンプトの一貫性が得られます。ほとんどの人が、一度だけ機能するプロンプトを作成します。
スコア分布、n = 1,018。中央値 60 · 50 未満 34.9% · 最高スコア: 90。
60年代の三次元クラスタ。 1人は31番に座っています。
このルーブリックは、プロンプトが明確であるか、具体的であるか、適切に構造化されているか、堅牢であるかという 4 つの点を評価します。最後の質問は、入力が想像したものと異なる場合に何が起こるかを尋ねます。空のメッセージ。貼り付けられたゴミ。別の言語での質問です。ユーザーが積極的にそれを破ろうとしている。
これはテール効果ではありません。ロバスト性は、データセット内のすべてのユースケース、すべての言語、すべてのプロンプト タイプにわたって、分類されたすべてのプロンプトの 95.7% で 4 つの次元の中で最も弱いものです。 83.6% が 50 未満のスコアを獲得しました。人々は、クリーンな質問、協力的なユーザー、テストしたものと同じような入力など、幸せな道を求めるプロンプトを作成します。
サブスコアがそれを端的に物語っています。ルーブリック全体で最もスコアの低い 2 つの動作は、どちらも堅牢性に関する動作です。
平均サブスコア、n = 969 分類されたプロンプト
それが重要な理由: 明瞭さで 60 点、堅牢性で 30 点のスコアを獲得したプロンプトは、デモでは完璧に動作します。

実際の使用では、乱雑な入力が存在するため、実際の使用の最初の 1 週間で d は崩壊します。 「やってみたらうまくいった」と「うまくいった」の差は、統計的にはこの差です。
31 と 60 の間の差は 4 文です
すべてのプロンプトに 4 つの構造的習慣のタグを付け、それぞれを使用した場合と使用しない場合の平均スコアを比較しました。これらはいずれも専門知識を必要としません。どれも1文から4文までです。
1 行: 「番号付きリストとして回答」/「有効な JSON を返す」/「3 つの短い段落」。
モデルに「ソースを発明しないでください」というやってはいけないことを伝えます。 / 「200 単語以内にしてください。」
「あなたは心配する親に説明する小児科看護師です」という視点で役割を与えてください。
必要な出力の例を 1 つ貼り付けます。分布をシフトするには 1 つで十分です。
寝台車はその一例です。理論上のリフトは最小 (+10) ですが、ヘッドルームはこれまでで最も大きく、プロンプトの 94.8% にはヘッドルームがありません。誰もが「モデルの例を示してください」という言葉を聞いたことがあるでしょう。ほとんど誰もやっていません。
コンテンツ作成と教育が大半を占めており、これらを合わせて機密プロンプトの 58% を占めています。データセットは 74% が英語、19% がアラビア語、6% がポルトガル語です。プロンプトの 67% はシステム プロンプトであり、何度も実行され、弱点のコストが倍増する再利用可能な種類のプロンプトです。
n = 20 未満のセグメント (財務、法務、生産性、AI エージェント、研究) はこの表から除外されます。小さなサンプルは大声で間違った主張をします。言語分割: 英語 74%、AR 19%、PT 6%。
1 つだけ修正する場合は、クラッシュを修正してください
このデータは、実際に再利用するプロンプトを改善するための操作順序を示唆しています。
悪い日のために 1 行追加します。「入力内容が空、不明瞭、または本題から外れている場合は、推測するのではなく、そう言って説明を求めてください。」この一文は、プロンプトのスコアが 50 未満である行動の 85% に対応しています。
出力形式を宣言します。データセット内で測定された最大のリフト (+29)、コスト o

ね文。
2 つまたは 3 つの制約を設定します。これはモデルが決して実行してはいけないことです。 (+24)
適切な出力の例を 1 つ貼り付けます。この習慣だけで、プロンプトの上位 5% に入ることができます。 (+10)
すべての数値は、2026 年 5 月 20 日から 7 月 11 日までに PromptEval の評価者に送信された 1,018 件のプロンプトに基づいており、4 つの次元にわたる固定 LLM 判定ルーブリックによって 0 ～ 100 点のスコアが付けられています。分類ラベル (ユースケース、タイプ、言語) は、別のモデル パスを介して 969 行をカバーします。プロンプト テキストがベンチマーク データセットに保存されることはありません。選択バイアス、宣言: これらは人々が評価のために提出することを選択したプロンプトであり、多くの場合問題を疑っているため、スコアは一般母集団よりも低く偏っている可能性があります。ここにあるものはすべて、「すべてのプロンプト」ではなく、「評価のために送信されたプロンプト」と読んでください。ライブベンチマークハブ上の完全な方法論。
CC BY 4.0。出典とリンクを付けて図を再公開します。このエディションは凍結されています。これらの数値は変わりません。
このレポートが構築されているのと同じ 4 次元スコア (堅牢性の数値を含む) を取得します。月に 3 回の無料評価。カードはありません。

## Original Extract

We scored 1,018 real prompts 0–100. Average: 54. But the story is one dimension: robustness averages 31.5/100 and is the weakest link in 96% of prompts. The inaugural State of Prompt Quality edition.

The Robustness Gap — State of Prompt Quality, Q3 2026 | PromptEval Prompt Eval ← live benchmark state of prompt quality · q3 2026 edition The Robustness Gap
We scored 1,018 real prompts, 0–100, with a fixed rubric. The average was 54. But that's not the finding. The finding is that almost every prompt fails in the same place, and it's not the place people polish.
Most prompts are mediocre in a very specific way
The distribution is not a disaster story. Two thirds of prompts land between 40 and 79: functional, gets-an-answer territory. Only 12% are truly broken (below 30). The problem is the ceiling: just 10.5% reach 75, the score where a prompt is consistent enough to trust in repeated use. Almost everyone writes a prompt that works once .
Score distribution, n = 1,018. Median 60 · 34.9% below 50 · highest score: 90.
Three dimensions cluster in the 60s. One sits at 31.
The rubric scores four things: is the prompt clear, is it specific, is it well-structured, and is it robust . That last one asks what happens when the input isn't what you pictured. Empty message. Pasted garbage. A question in another language. A user actively trying to break it.
This is not a tail effect. Robustness is the weakest of the four dimensions in 95.7% of all classified prompts, across every use case, every language, every prompt type in the dataset. 83.6% score below 50 on it. People write prompts for the happy path: the clean question, the cooperative user, the input that looks like the one they tested with.
The sub-scores say it plainly. The two lowest-scoring behaviors in the entire rubric are both robustness behaviors:
average sub-score, n = 969 classified prompts
Why it matters: a prompt that scores 60 on clarity and 30 on robustness behaves perfectly in the demo and falls apart the first week of real use, because real use is where the messy input lives. The gap between "worked when I tried it" and "works" is, statistically, this gap.
The gap between 31 and 60 is four sentences
We tagged every prompt for four structural habits and compared average scores with and without each. None of these require expertise. All of them are one to four sentences.
One line: "Answer as a numbered list" / "Return valid JSON" / "Three short paragraphs."
Tell the model what NOT to do: "Do not invent sources." / "Stay under 200 words."
Give it a role with a point of view: "You are a pediatric nurse explaining to a worried parent."
Paste one example of the output you want. One is enough to shift the distribution.
The sleeper is examples. It has the smallest lift on paper (+10) but the most headroom by far: 94.8% of prompts don't have one. Everyone has heard "give the model examples"; almost nobody does it.
Content creation and education dominate: together they're 58% of classified prompts. The dataset is 74% English, 19% Arabic, 6% Portuguese; 67% of prompts are system prompts, the reusable kind that runs many times and multiplies the cost of a weak spot.
Segments under n = 20 (finance, legal, productivity, AI agents, research) are held back from this table. Small samples make loud, wrong claims. Language split: EN 74% · AR 19% · PT 6% .
If you only fix one thing, fix the crash
The data suggests an order of operations for anyone improving a prompt they actually reuse:
Add one line for the bad day: "If the input is empty, unclear, or off-topic, say so and ask for clarification instead of guessing." That single sentence addresses the behavior 85% of prompts score under 50 on.
Declare the output format. Biggest measured lift in the dataset (+29), costs one sentence.
Set two or three constraints: what the model must never do. (+24)
Paste one example of a good output. You'll be in the top 5% of prompts by that habit alone. (+10)
Every figure comes from 1,018 prompts submitted to PromptEval's evaluator between May 20 and July 11, 2026, scored 0–100 by a fixed LLM-judge rubric across four dimensions. Classification labels (use case, type, language) cover 969 rows via a separate model pass. Prompt text is never stored in the benchmark dataset. Selection bias, declared: these are prompts people chose to submit for evaluation, often suspecting a problem, so scores likely skew lower than the general population. Read everything here as "prompts submitted for evaluation," never "all prompts." Full methodology on the live benchmark hub .
CC BY 4.0. Republish any figure with attribution and a link. This edition is frozen; these numbers will not change.
Get the same four-dimension score this report is built on, including your robustness number. Three free evaluations a month. No card.

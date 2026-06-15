---
source: "https://barvhaim.github.io/visual-ai-metrics/"
hn_url: "https://news.ycombinator.com/item?id=48542226"
title: "Show HN: AI Metrics, Visually"
article_title: "AI Metrics — A Learning Toybox"
author: "bignet"
captured_at: "2026-06-15T16:43:36Z"
capture_tool: "hn-digest"
hn_id: 48542226
score: 3
comments: 0
posted_at: "2026-06-15T14:57:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI Metrics, Visually

- HN: [48542226](https://news.ycombinator.com/item?id=48542226)
- Source: [barvhaim.github.io](https://barvhaim.github.io/visual-ai-metrics/)
- Score: 3
- Comments: 0
- Posted: 2026-06-15T14:57:16Z

## Translation

タイトル: HN を表示: AI メトリクスを視覚的に表示
記事のタイトル: AI メトリクス — 学習のおもちゃ箱

記事本文:
モデルを微調整するときに満たすメトリクスは、使用時にグループ化され、遊び心とインタラクティブになっています。
ここでのすべての指標は、ネットに関して誰かが尋ねる不安な質問にすぎません。この写真を保持すると、残りが続きます。
「私が捕まえたものはすべて、実際に魚ですか？それとも私も長靴を履いたのですか？」
私が陽性反応を示したもののうち、どれだけが正しかったでしょうか?ネットのブーツ (🥾) = 誤検知。
「湖の中の魚はすべて、私が実際に捕まえたのでしょうか、それとも泳ぎ通ったのでしょうか？」
実際にそこにあったもののうち、私はいくら得たのでしょうか？逃げた魚 = 偽陰性。
📉 損失、検証損失、困惑
品質を判断する前に、モデルが学習しているかどうかを観察します。微調整中、ダッシュボードには時間の経過とともに損失が減少していることが表示されます。初心者にとって最も重要な習慣は、トレーニング カーブではなく検証カーブに注目することです。
困惑 = e^loss 。これを次のように読みます。「モデルは各ステップで、同じ可能性を持つ単語の中からいくつ選択していますか?」困惑 1 = 完全に確かで正しい。 Perplexity 50 = 50 の選択肢から選ぶのと同じくらい不確かです。低いほど良いです。
列車の損失は常に減少し続けます - モデルは記憶することができます。トレイン損失が低下する一方で、ヴァル損失が上昇に転じる場合、それは学習ではなく暗記です。そのギャップは、早めにやめるべきという合図です。
🎲 分かりやすく言えば、複雑性とクロスエントロピー
実数についても同じ考えです。言語モデルの仕事: これまでの単語を考慮して、次の単語、つまり語彙内のすべての単語の確率を予測します。困惑は次のように問いかけます。
「平均して、各ステップでモデルの間で不確実な単語がいくつありますか?」
本当の次の単語は mat です。モデルが与えた確率をドラッグします。
残りの確率は他の単語に分散されます。モデルの推定分布は次のとおりです。
トレーニングでは、クロスエントロピー損失 = 平均的な驚きを抽象単位 (nat) で示します。あたり

plexity = e^loss は、その驚きを具体的な選択肢の数に変換します。損失 0 → 人数 1。損失 2.3 → 人数 ≈ 10。損失 4.6 → 人数 ≈ 100。同じ情報、より友好的なユニット。
それは語彙とタスクに依存するため、ベースラインと比較してのみ意味があります。一般英語に関する強力な最新の LLM は、困惑度 3 ～ 15 の周りにあります。 50,000 の語彙 ≈ 50,000 にわたるランダムな推測。これを比較するために使用します。微調整により、ドメインのテキストに対する混乱が軽減されましたか?
以下、12項目です。絵文字は真実です。🐟は実際にはポジティブで、🥾は実際にはネガティブです。項目をクリックしてモデルの予測を切り替えます (青いリング = 「モデルはポジティブと予測する/ネットに引っかかる」)。マトリックスとメトリックの更新を確認します。
青いリング = 陽性と予測されます。ブーツをつかまずにすべての魚を捕まえてみてください。
精度と再現率: スライダー
実際のモデルはスコア (0 ～ 1) を出力し、しきい値を選択します (スコア ≥ しきい値 → 陽性を予測)。以下の 14 項目には固定スコアが付けられています。しきい値をドラッグして、精度とリコールの反対方向のプルを監視します。
閾値が低い = 広範囲に網を張ります (再現率は高く、精度は低くなります)。閾値が高い = 確実なもののみ (精度が高く、再現率が低い)。
太字 = 実際にはポジティブ (🐟)。青い枠 = この閾値で陽性と予測される。
精度 = 「何分数が正解でしたか?」 — 最も直感的な指標ですが、不均衡なデータでは最も危険です。これは、実際に少数のトランザクションのみが詐欺である詐欺検出器です。詐欺がどれほどまれであるかをドラッグしてください:
ここで、すべてに対して「決して不正が行われない」と予測するだけの「怠惰なモデル」を見てみましょう。実際の作業はまったく行われません。
これが、精度、再現率、F1 が存在する理由のすべてです。彼らは、安易な多数派の層を無視し、「重要なことを理解できましたか?」と尋ねます。モデルの精度が 98% であっても役に立たない場合があります。
「私の漁獲量のうち、魚はいくらですか？」
falseの場合に使用します

警報器は高価です。スパム フィルター — 実際の電子メールをブロックすることは、スパムを見逃すことよりも悪いです。
「たくさんの魚のうち、私は何匹捕まえましたか？」
ミスが高くつく場合に使用します。がんのスクリーニング — 誤報は病気の患者を見逃すことを防ぎます。
調和平均 — どちらかが低い場合は消滅します。
バランスが必要で、片側だけを最大化することで高スコアを望まない場合に使用します。
ROUGE = 言葉としては同じ考え
モデルがテキスト (要約、翻訳) を生成すると、「湖の魚」が参照内の単語になります。 ROUGE はこう尋ねます: 生成されたテキストはどの程度の参照を捉えましたか?
ROUGE vs BLEU — 同じアイデアの二つの側面
ROUGE と BLEU はどちらも単語の重複をカウントします。彼らは漁網とは異なる不安を抱えながら進んでいるだけです。
「リファレンスにすべて記載しましたか?」
要約の標準 — 重要なポイントを省略した要約は、失敗を恐れるものです。ミスは傷つきます。
「私が生成したものは本当に正しいのでしょうか?」
翻訳の標準 — 属さない言葉をでっち上げることは、あなたが恐れる失敗です。 「簡潔さのペナルティ」が追加されるため、完璧な単語を 1 つだけ出力しても高スコアを獲得することはできません。
盲点: どれも意味を理解していない
ROUGE と BLEU は面の重なりを数えます。彼らにとって、「映画は素晴らしかった」と「映画は素晴らしかった」にはほとんど何の共通点もありません。同じ意味であるにもかかわらず、スコアはほぼゼロです。チャットボットを微調整すると、判断力が弱くなります。
正確な単語ではなく、エンベディングを比較します。 「映画/映画」、「素晴らしい/素晴らしい」スコアが近似一致として表示されます。 ROUGE/BLEU の視覚障害に対する意味を意識した修正。
強力なモデル (クロードなど) に、有用性、正確さ、トーンに関して微調整されたモデルの回答を採点してもらいます。現在、命令調整モデルで主流の方法。
「答え A は B よりも優れていますか?」多くのプロンプトにわたって。ペアごとの優先順位は、RLHF および Chatbot-Arena ランキングで使用されるものです。引き続き人間が主導権を握る

ld 主観的な品質の基準。

## Original Extract

The metrics you meet when finetuning a model — grouped by when you use them, made playful and interactive.
Every metric here is just an anxious question someone asks about that net. Hold this picture and the rest follows.
"Everything I caught — is it actually fish, or did I also pull up boots?"
Of what I flagged positive, how much was right? Boots in the net (🥾) = false positives.
"All the fish in the lake — did I actually catch them, or did some swim through?"
Of what was really there, how much did I get? Escaped fish = false negatives.
📉 Loss, validation loss & perplexity
Before judging quality , you watch whether the model is learning at all. During finetuning your dashboard shows loss dropping over time. The single most important habit for a beginner: watch the validation curve, not the training curve.
perplexity = e^loss . Read it as: "how many equally-likely words is the model choosing between at each step?" Perplexity 1 = perfectly certain & correct. Perplexity 50 = as unsure as picking from 50 options. Lower is better.
Train loss always keeps dropping — the model can memorize. When val loss turns back upward while train loss falls, it's memorizing, not learning. That gap is your signal to stop early .
🎲 Perplexity & cross-entropy, in plain terms
Now the same idea with real numbers. A language model's job: given the words so far, predict the next one — a probability for every word in its vocabulary. Perplexity asks:
"On average, how many words is the model unsure between at each step?"
The true next word is mat . Drag how much probability the model gave it:
The remaining probability is spread over other words. Here's the model's guess distribution:
Training shows cross-entropy loss = average surprise, in abstract units (nats). perplexity = e^loss converts that surprise back into a tangible count of choices . Loss 0 → ppl 1. Loss 2.3 → ppl ≈ 10. Loss 4.6 → ppl ≈ 100. Same info, friendlier units.
It depends on vocabulary & task, so it's only meaningful relative to a baseline. A strong modern LLM on general English sits around perplexity 3–15 . Random guessing across a 50k vocab ≈ 50,000. You use it to compare : did finetuning lower my perplexity on my domain's text?
Below are 12 items. The emoji is the truth : 🐟 is actually positive, 🥾 is actually negative. Click an item to toggle your model's prediction (a blue ring = "model predicts positive / caught in net"). Watch the matrix and metrics update.
Blue ring = predicted positive. Try to catch all the fish without grabbing boots.
Precision vs Recall: the slider
Real models output a score (0–1), and you pick a threshold : score ≥ threshold → predict positive. Below, 14 items have fixed scores. Drag the threshold and watch precision and recall pull in opposite directions.
Low threshold = cast a wide net (high recall, low precision). High threshold = only the sure things (high precision, low recall).
Bold = actually positive (🐟). Blue outline = predicted positive at this threshold.
Accuracy = "what fraction did I get right?" — the most intuitive metric, and the most dangerous on imbalanced data . Here's a fraud detector where only a few transactions are actually fraud. Drag how rare fraud is:
Now meet the "lazy model" that just predicts "never fraud" for everything — it does zero real work:
This is the whole reason precision, recall & F1 exist: they ignore the easy majority class and ask "did you catch the thing that matters?" A model can have 98% accuracy and be useless.
"Of my catch, how much is fish?"
Use when false alarms are costly. Spam filter — blocking real email is worse than missing some spam.
"Of all the fish, how many did I catch?"
Use when misses are costly. Cancer screening — a false alarm beats missing a sick patient.
Harmonic mean — dies if either is low.
Use when you need balance and don't want a high score from maximizing just one side.
ROUGE = the same idea, for words
When a model generates text (summaries, translations), the "fish in the lake" become the words in the reference . ROUGE asks: how much of the reference did the generated text catch?
ROUGE vs BLEU — two sides of the same idea
ROUGE and BLEU both count word overlap; they just lead with different anxieties from the fishing net:
"Did I cover everything in the reference?"
Standard for summarization — a summary that drops key points is the failure you fear. Misses hurt.
"Is everything I generated actually correct?"
Standard for translation — inventing words that don't belong is the failure you fear. It adds a "brevity penalty" so you can't get a high score by outputting just one perfect word.
The blind spot: none of these understand meaning
ROUGE and BLEU count surface overlap . To them, "the film was great" and "the movie was excellent" share almost nothing — near-zero score, despite identical meaning. For finetuning a chatbot, that makes them weak judges.
Compares embeddings , not exact words. "film/movie", "great/excellent" score as near-matches. The meaning-aware fix for ROUGE/BLEU's blindness.
Ask a strong model (e.g. Claude) to score your finetuned model's answers for helpfulness, correctness, tone. The dominant method today for instruction-tuned models.
"Is answer A better than B?" across many prompts. Pairwise preference is what RLHF and Chatbot-Arena rankings use. Humans remain the gold standard for subjective quality.

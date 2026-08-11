---
source: "https://enterpilot.io/blog/make-expensive-ai-models-slower/"
hn_url: "https://news.ycombinator.com/item?id=49258184"
title: "Want people to use cheaper AI models? Make expensive models slower"
article_title: "Want People to Use Cheaper AI Models? Make Expensive Models Slower | enterpilot Blog"
author: "santiago-pl"
captured_at: "2026-08-11T14:13:54Z"
capture_tool: "hn-digest"
hn_id: 49258184
score: 1
comments: 1
posted_at: "2026-08-11T13:41:12Z"
tags:
  - hacker-news
  - translated
---

# Want people to use cheaper AI models? Make expensive models slower

- HN: [49258184](https://news.ycombinator.com/item?id=49258184)
- Source: [enterpilot.io](https://enterpilot.io/blog/make-expensive-ai-models-slower/)
- Score: 1
- Comments: 1
- Posted: 2026-08-11T13:41:12Z

## Translation

タイトル: より安価な AI モデルを使用してもらいたいですか?高価なモデルを遅くする
記事のタイトル: より安価な AI モデルを人々に使用してもらいたいですか?高価なモデルを遅くする |エンターパイロットのブログ
説明: 増大する AI コストを削減する簡単な方法: AI ゲートウェイ層で高価なモデルにコストに比例した遅延を追加します。

記事本文:
パイロットに入る ホーム ブログ GoModel ← すべての投稿に戻る より安価な AI モデルを使用したいですか?高価なモデルを遅くする
2026 年 8 月 11 日
· ヤクブ・A・ワセク
多くの企業は、AI コストの急激な増加に悩まされています。
通常、最大の節約は、適切なモデルを使用するという 1 つの単純な変更によってもたらされます。
タスクのために。
大規模な AI コーディング コストの管理に関する Databricks の記事を読む
これについて考えるきっかけを与えてくれました。彼らは、
効率のフロンティア、より安価な機能モデルの使用、およびプログレッシブの追加
支出が増加するにつれて摩擦が生じます。
そこで私はこう考えました。モデルの経済的コストがすぐに感じられるとしたらどうなるでしょうか。
時間として？
文書を要約したり製品を書いたりするのに最も重いモデルは必要ありません
説明、サポート チケットの分類、または小さなコード変更を行います。しかし、どうやって
人々に適切なモデルを選ばせるのですか？
その人たちは開発者かもしれません。彼らはマーケティングチームかもしれません。彼らはそうかもしれない
モデル名のリストを見て、自然に一番上のモデルをクリックする人はいるでしょう。
GoModel では、行動心理学に基づいたシンプルなアイデアを思いつきました。
高価なモデルほど待つ時間が長くなります。
これをコスト比例遅延と呼びます。
今日は素晴らしいモデルがたくさんあります。能力の差が大きく出ることが多い
価格差よりも小さいです。
たとえば、Grok 4.5 のコストは 100 万入力トークンあたり 2 ドル、100 万あたり 6 ドルです。
出力トークン 。クロード寓話 5
費用は 10 ドルと 50 ドルです。寓話の方が強いし、
しかし、すべてのタスクで 5 ～ 8 倍優れているわけではありません。
ある実際のコーディング ベンチマークでは、Fable 5 のスコアは 92.8 でしたが、Grok のスコアは 4.5 でした。
81.1 。ただし、フル実行の費用は Fable で 67.42 ドル、Grok で 9.24 ドルでした。
GLM 5.2 は Opus に対して同様のパターンを示しました
5.0 : スコアの差が小さい
そしてコストのはるかに大きな違い。
ベンチマークはプロではありません

ダクションの仕事量。モデルを自分でテストする必要があります
タスク。しかし、一般的なパターンは明らかです。最良のモデルは、多くの場合、ほんの少ししかありません。
より良く、より高価です。
ほとんどのタスクには、安価なモデルで十分です。
コストダッシュボードの問題
企業はすでに人のトークン数、モデルの価格、月次予算を示しています。
その情報は役に立ちますが、行動はすぐには変わりません。
コストは抽象的です。請求書は後から来ます。小さな不便さは、
より安価なモデルを選択することが今起こっています。
待つのは違います。人々はそれをすぐに感じます。
私たちはこれをウェブサイトからすでに知っています。ページの読み込みが遅いと、ユーザーが離れたり、
別のパスを試してください。同じ動作が AI モデルでも機能します。
モデルは常に顕著な待ち時間を追加し、人々は本当にそうするのかどうか尋ね始めます
それが必要です。
AI ゲートウェイでの遅延を追加する
AI ゲートウェイまたは AI コントロール プレーンは、これを行うのに適した場所です。それはもう
どのモデルがリクエストされたか、誰がリクエストしたか、そしてどれくらいの費用がかかるかを知っています。
GoModel では、任意のモデルにタイム ペナルティを追加できます。減速は相対的なものです
モデルの測定された推論時間: 値 0.5 は 50% を追加します。ストリーミング用
応答すると、GoModel はチャンクをバッファリングし、より遅いタイムラインで解放します。
GoModel では、モデルごとに追加の推論時間を構成できます。ここでは、0.5
クロード オーパス 5 に 50% のタイムペナルティを追加します。
ライブ GoModel モデルで試すことができます
リスト。
単純なポリシーは次のようになります。
正確な数字は重要ではありません。チームとワークフローに合わせて調整してください。
ポイントは、金銭的な差を時間の差として見える化することです。あ
相対的な速度の低下も自然に拡大します。推論が長いほど時間がかかります。
迅速な対応よりもペナルティ。
これは主にインタラクティブな使用を目的としています。バックグラウンド ジョブにはルーティング ルール、予算、
ar の代わりにハードリミット

正式な待機。
安価なモデルを本当に高速に作成することもできます
遅延を追加する必要がない場合もあります。格安で仕事を削除できます
代わりにパス。
たとえば、ゲートウェイで複数の低速ガードレールが実行されている場合、一部のガードレールは機能しない可能性があります。
リスクの低い内部タスクに必要です。不要なチェックをスキップすると、
安価なモデルはより速く、より魅力的です。
これは、ユースケースにとって意味がある場合にのみ実行してください。セキュリティを解除しないでください。
数秒を節約するためだけに、プライバシー、コンプライアンス、または安全管理を強化します。
目標は重いモデルを禁止することではありません
場合によっては、入手可能な最高のモデルが必要になることがあります。アーキテクチャに関する難しい決断、
生産上の重大な問題や複雑な研究課題により、
追加費用。
目標は、これらのモデルをブロックしないことです。目標は、小さな瞬間を追加することです。
デフォルトでは、使用する前に摩擦が発生します。
安価な経路を迅速に実現します。高価なパスは計画的に実行してください。
モデルの選択とそれに伴う AI の請求額を変更するには、これで十分かもしれません。
私は透明性が健全な組織文化にとって極めて重要であると信じています。もしあなたが
コスト比例遅延を導入し、それが存在する理由と方法を明確かつ率直に説明する
ルールが機能しているか、会社が何を達成しようとしているか。のような感じになるはずです
効率性の目標は共有されており、隠れた罰ではありません。人を祝うこともできる
AI を効率的に使用しながらも素晴らしい成果を上げている人たち - 象徴的なもの
たとえば、Vibe Coder of the Year 賞。結果と賢明なことに報酬を与える
支出を減らすだけではなく、リソースを活用することです。

## Original Extract

A simple way to reduce growing AI costs: add cost-proportional delay to expensive models at the AI gateway layer.

enter pilot Home Blog GoModel ← Back to all posts Want People to Use Cheaper AI Models? Make Expensive Models Slower
August 11, 2026
· Jakub A. Wasek
Many companies struggle with exponentially growing AI costs.
The biggest savings usually come from one simple change: use a model appropriate
for the task .
Reading the Databricks article on managing AI coding costs at scale
inspired me to think about this. They describe moving work toward the
efficiency frontier, using cheaper capable models, and adding progressive
friction as spend increases.
It made me ask: what if the financial cost of a model could be felt immediately
as time?
You do not need the heaviest model to summarize a document, write a product
description, classify a support ticket, or make a small code change. But how do
you make people choose the appropriate model?
Those people might be developers. They might be the marketing team. They might
be anyone who sees a list of model names and naturally clicks the one at the top.
In GoModel, we came up with a simple idea grounded in behavioral psychology:
The more expensive the model, the longer you wait for it.
We call it cost-proportional delay .
Today we have many amazing models. The difference in capability is often much
smaller than the difference in price.
For example, Grok 4.5 costs $2 per million input tokens and $6 per million
output tokens . Claude Fable 5
costs $10 and $50 . Fable is stronger,
but it is not five to eight times better for every task.
In one real coding benchmark, Fable 5 scored 92.8 , while Grok 4.5 scored
81.1 . But the full run cost $67.42 with Fable and $9.24 with Grok.
GLM 5.2 showed a similar pattern against Opus
5.0 : a smaller difference in the score
and a much bigger difference in cost.
Benchmarks are not your production workload. You should test models on your own
tasks. But the general pattern is clear: the best model is often only a little
better and a lot more expensive.
For most tasks, the cheaper model is enough.
The problem with cost dashboards
Companies already show people token counts, model prices, and monthly budgets.
That information is useful, but it does not change behavior quickly.
The cost is abstract. The invoice comes later. The small inconvenience of
choosing a cheaper model happens now.
Waiting is different. People feel it immediately.
We already know this from websites. When a page loads slowly, people leave or
try another path. The same behavior can work with AI models: if the expensive
model always adds a noticeable wait, people start asking whether they really
need it.
Add the delay at the AI gateway
The AI gateway or AI control plane is the right place to do this. It already
knows which model was requested, who requested it, and how much it costs.
In GoModel, you can add a time penalty to any model. The slowdown is relative to
the measured inference time of the model: a value of 0.5 adds 50%. For streaming
responses, GoModel buffers the chunks and releases them on the slower timeline.
GoModel lets you configure extra inference time for each model. Here, 0.5
adds a 50% time penalty to Claude Opus 5.
You can try it in the live GoModel model
list .
A simple policy could look like this:
The exact numbers are not important. Tune them for your team and your workflow.
The point is to make the financial difference visible as a time difference. A
relative slowdown also scales naturally: a longer inference gets a larger time
penalty than a quick response.
This is mainly for interactive use. Background jobs need routing rules, budgets,
and hard limits instead of artificial waiting.
You can also make cheap models genuinely faster
Sometimes you do not need to add a delay. You can remove work from the cheap
path instead.
For example, if your gateway runs several slow guardrails, some of them may not
be necessary for low-risk internal tasks. Skipping an unnecessary check can make
the cheap model faster and more attractive.
Do this only when it makes sense for your use case. Do not remove security,
privacy, compliance, or safety controls just to save a few seconds.
The goal is not to ban heavy models
Sometimes you need the best model available. A difficult architecture decision,
a critical production issue, or a complex research task can easily justify the
extra cost.
The goal is not to block those models. The goal is to add a small moment of
friction before using them by default.
Make the cheap path fast. Make the expensive path deliberate.
That may be enough to change model selection - and the AI bill with it.
I believe transparency is crucial to a healthy organizational culture. If you
introduce cost-proportional delay, explain clearly and openly why it exists, how
the rules work, and what the company is trying to achieve. It should feel like a
shared efficiency goal, not a hidden punishment. You can even celebrate people
who use AI efficiently and still deliver great outcomes - with a symbolic
Vibe Coder of the Year award, for example. Reward the outcome and the smart
use of resources, not low spending alone.

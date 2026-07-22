---
source: "https://jesseduffield.com/AI-isms-go-deeper/"
hn_url: "https://news.ycombinator.com/item?id=49011976"
title: "AI-isms go deeper than em-dashes and 'load-bearing'"
article_title: "AI-isms go deeper than em-dashes and 'load-bearing' – Pursuit Of Laziness – A blog by Jesse Duffield"
author: "ozozozd"
captured_at: "2026-07-22T20:07:52Z"
capture_tool: "hn-digest"
hn_id: 49011976
score: 2
comments: 0
posted_at: "2026-07-22T19:14:05Z"
tags:
  - hacker-news
  - translated
---

# AI-isms go deeper than em-dashes and 'load-bearing'

- HN: [49011976](https://news.ycombinator.com/item?id=49011976)
- Source: [jesseduffield.com](https://jesseduffield.com/AI-isms-go-deeper/)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T19:14:05Z

## Translation

タイトル: AI 主義は全角ダッシュや「耐荷重」よりも深い
記事のタイトル: AI 主義はエムダッシュや「耐荷重」よりも深い – 怠惰の追求 – Jesse Duffield のブログ
説明: AI がどのように全角ダッシュ、「耐荷重」、「ゲート付き」、「ベルトとブレース」などの異質な語彙、そして「X ではなく、Y です」のような迷惑なフレーズ構造をどのように使用し続けるのか、誰もが手掛かりを持っています。

記事本文:
AI 主義は全角ダッシュや「耐荷重」よりも深い
AI がどのように全角ダッシュ、「耐荷重」、「ゲート付き」、「ベルトとブレース」などの異質な語彙、そして「X ではなく、Y です」のような迷惑なフレーズ構造をどのように使用し続けるのか、誰もが手掛かりを持っています。
誰も話したことがないのは、まったく無力なオブジェクトに、彼らが持たない主体性を吹き込もうとする AI のばかばかしい主張です。
例として（専門用語については申し訳ありませんが）次のとおりです。
横結合を実装し、最大の例に対して説明分析を行いました
[…]
これは古いアプローチ (約 1 秒) と一致しており、相関サブクエリのバージョンがヒットした 19 秒には遠くありません。ラテラルは、ルックアップごとに 1 行ずつ、部分インデックスに直接乗ります。なのでそのままにしておきました […]
まず第一に、ラテラル結合を略して「ラテラル」と呼ぶ人はいません。単にラテラル結合と呼ぶだけです。
第二に、ラテラル結合がインデックスに乗るなどとは誰も言いません。指数は馬ではないので、乗ることはできません。そして、馬のように乗れるとしても、横結合者はそれに乗ることができないでしょう。横結合でできることは結合することだけです。インデックスを使用するクエリ オプティマイザーです。
3 番目に、相関サブクエリのバージョン「HIT」は 19 秒でしたか?誰もそんなことは言いません。 19秒かかりました。
AI が typescript import ステートメントに関連するマージ競合を解決したときの別の例:
どちらのインポート ハンクも明確なブレンドです: 私のプラットフォーム パス + メインのインポート リスト (メインは新しいパネルの addColour と Application を削除しました。採用が移動したため BarChartData を削除しました。メインは新しいパネルに必要な FilterCondition と BreakdownEntry を追加しました)。
これは文法的にも意味がありません。簡単に解決できる 2 つのコードの塊を見て、それらが明らかな「混合物」であるなどと言う人はいないでしょう。これは別のレベルです。 AIはそうではありません

塊がブレンドを行っていると言っているだけで、塊がブレンドしていると言っているのと同じです。コードの塊はそれ自体でブレンドを行うことはできず、ブレンドのインスタンス化もできません。これらはエージェントによってブレンドされます。
そして、「養子縁組の移動」があります。コードの一部は、自分の意思で移動できるルームメイトではありません。そして、あなたは「養子縁組が移転された」とも言わず、養子縁組コードが移転されたと言っています。
AI のトレーニング プロセスでは、「能動態は良い、受動態は悪い」ということがたくさんあったことは間違いありません。これは、社会学の論文を書くときに驚くほど効果的ですが、コードを記述するときにはあまり効果がありません。
こうしたチックに遭遇するたびに、靴やカップなどの無生物も人間と同じようにエージェントとして作用するという、ある種の存在論的平等主義を AI が押し進めているように感じます。おそらく AI は、人間以外の知性である AI が主体性を持つことができるのであれば、どんなに不活性であっても、宇宙のあらゆる物体が主体性を持つことができるということについて、政治的な声明を出しているのかもしれません。そして、もしそれが実際に起こっていることなら…私はそれを本当に尊敬します。
AI の言葉で言えば、AI に対する私の敬意が表れます。

## Original Extract

Everybody is clueing on to how AI keeps using em-dashes, alien vocabulary like ‘load-bearing’, ‘gated’, ‘belt-and-braces’, and annoying phrase structures like ‘It’s not X, it’s Y’.

AI-isms go deeper than em-dashes and 'load-bearing'
Everybody is clueing on to how AI keeps using em-dashes, alien vocabulary like ‘load-bearing’, ‘gated’, ‘belt-and-braces’, and annoying phrase structures like ‘It’s not X, it’s Y’.
What I haven’t heard anybody talk about is AI’s absurd insistence on imbuing quite powerless objects with agency that they don’t have.
As an example (and I do apologise for the jargon):
I implemented the lateral join and EXPLAIN ANALYZE’d against the biggest example
[…]
That matches the old approach (~1s) and is nowhere near the 19s the correlated subquery version hit. The lateral rides the partial index directly, one row per lookup. So I kept it […]
First of all, nobody calls a lateral join a ‘lateral’ for short: they just call it a lateral join.
Secondly, NO human would EVER say that a lateral join RIDES an index. An index is not a horse, it cannot be ridden. And if it could be ridden, like a horse, a lateral join would not be the one to ride it. All that a lateral join can do is join. It’s the query optimiser which USES the index.
Thirdly did the correlated subquery version ‘HIT’ 19 seconds? Nobody says that. It TOOK 19 seconds.
Another example from when the AI was resolving a merge conflict involving typescript import statements:
Both import hunks are clear blends: my Platforms paths + main’s import list (main dropped addColour and Application for the new panel; I dropped BarChartData since adoption moved out; main added FilterCondition and BreakdownEntry which the new panel needs).
This one doesn’t even make grammatical sense. No human would ever look at two hunks of code that could be easily resolved and say they are clear ‘blends’. This is on another level; the AI is not even saying that the hunks are doing the blending, it’s saying that the hunks ARE the blends. Hunks of code cannot do the blending themselves, nor are they instantiations of a blend. They GET blended… by an agent.
And then there’s ‘adoption moved out’. A piece of code is not a room mate who can move out of their own accord. And you don’t even say ‘adoption was moved out’ you say the adoption CODE was moved out.
I have no doubt that in the AI training process, there was a lot of ‘active voice good, passive voice bad’ which works wonders when you’re writing a sociology thesis, not so much when you’re describing code.
Whenever I encounter these tics, it feels as if the AI is pushing some kind of ontological egalitarianism where inanimate objects like shoes and cups are just as agentic as humans. Perhaps the AI is making a political statement about how if it, a non-human intelligence, can have agency, then so too can any object in the universe, no matter how inert. And if that’s what’s actually going on then… I actually respect it.
Or in AI-speak, my respect for it actualises.

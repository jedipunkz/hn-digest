---
source: "https://newsletter.kentbeck.com/p/baking-a-model"
hn_url: "https://news.ycombinator.com/item?id=49305969"
title: "Baking a Model: A Metaphor for LLM Training"
article_title: "Baking a Model - by Kent Beck"
author: "KentBeck"
captured_at: "2026-08-15T00:39:38Z"
capture_tool: "hn-digest"
hn_id: 49305969
score: 1
comments: 0
posted_at: "2026-08-14T23:46:45Z"
tags:
  - hacker-news
  - translated
---

# Baking a Model: A Metaphor for LLM Training

- HN: [49305969](https://news.ycombinator.com/item?id=49305969)
- Source: [newsletter.kentbeck.com](https://newsletter.kentbeck.com/p/baking-a-model)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T23:46:45Z

## Translation

タイトル: モデルのベイク処理: LLM トレーニングのメタファー
記事のタイトル: モデルのベイク処理 - Kent Beck 著
説明: Motorola 6800 の取扱説明書を見つめながら、高校からバスに向かって歩いていたことを覚えています。

記事本文:
モデルのベイク処理 - Kent Beck 著
ソフトウェア設計: まず整理整頓?
購読 サインイン モデルのベイク処理
Kent Beck 2026 年 8 月 14 日 32 1 1 シェア Motorola 6800 の取扱説明書を見つめながら、高校からバスに向かって歩いていたのを覚えています。ブール式、命令エンコーディング、タイミング テーブルなど、自分が何を見ているのかよく理解していませんでしたが、そのすべてのメカニズムに夢中になりました。ここにはこの複雑な機械があり、理解できれば力と制御が得られるでしょう。
私は今、AI モデルについても同じように感じています。まだ詳細を理解しているとは言えませんが、そのメカニズムにはとても興味を持っています。私は両方に興味があります:
模型を作る機械です。
この後者のトピック、つまりモデルがどのように構築されるかについて、この投稿 (およびフォローアップの可能性) で検討し始めます。
土曜日のフォカッチャ パン作りが大好きです。材料を 1 つの形にして、まったく別の形に変換します。材料自体は美味しくないけど、それを使って作ったものは美味しいです。ああ、ベイク処理は初期条件にも影響されます。プロセスの早い段階で小さな変更を加えることができますが、それは後で大きな影響を及ぼします。
私は冷蔵庫で一晩イースト菌を働かせる耐寒実験を行っています。モデルを理解しようとしていたとき、少なくともこれまでのプロセスを理解している限り、モデルの作成と類似点があることに気づきました (何か間違っていた場合はコメントで修正してください)。
まず最初に、モデルとは何を意味するのかについて徐々に明らかになります。過度に単純化しますが、少しずつ複雑さを明らかにしていきます。
「モデル」とは、人間と会話できるコンピューター システムを意味します (ほら、単純化しすぎだと言いましたが、お付き合いいただければ幸いです)

これにはそれほど時間はかかりません）。
実際にはそれよりも少し複雑です。モデルは 2 つの部分に分かれています。
入力、出力、シーケンス、認証などのフォーマットを処理するユーザー インターフェイス。
魔法が起こる適切なモデル。
ユーザー インターフェイスは、従来のプログラミング手法を使用して構築されています (ただし、「慣例」は過去 2 年間で根本的に変わりました)。ただし、モデルそのもの (探索のテーマ) は、根本的に異なる手法を使用して構築されています。
モデルは数字の入った袋です。今日の目的にはこれで十分です (おそらくいつか、モデルが有用な単語を生成するときにこれらの数値がどのように使用されるかがわかるようになるでしょう。しかし、最初にそれを少し理解する必要があります)。
どうやってこれらの数字を導き出したのでしょうか?一連のステートメントをレイアウトしてその結果プログラムを作成するプログラミングとは異なり、AI モデルはトレーニングから生成されます。
トレーニングは、成果物を変更して将来の動作を変更するという点でプログラミングに似ていますが、大きな違いもいくつかあります。
魔法のように一度に大量の正しい数字を推測することはできません。いや、いやいや。まずほぼ正確な数値を取得し、それからさらに良い数値を得るためにそれらを微調整します。 2 つのプロセスは大きく異なります。
(私の知る限り、トレーニング前、トレーニング後、トレーニング中期 (それについては私は何も知りません) がありますが、トレーニング前、トレーニング中、トレーニング後の構成以外に「トレーニング」はありません。語彙が進化することを願っています。)
事前トレーニングは大規模なものです。チーム全体が初期条件、つまりデータと空のモデルを設定します。彼らは、モデル内でデータを何十億回も逆方向と順方向に実行します。クラッシュが発生した場合に備えて、途中でスナップショットを作成します。彼らは、事前トレーニングが雑草に追いやられた兆候がないかチェックします。

編集を調整して再起動する必要があります。事前トレーニングは大きな賭けであり、数億ドルと（より高価な）数か月の遅れがかかります。
(トレーニング前の人々がどのように協力するかについてもっと学ぶ必要があります。)
事前トレーニングは、モデル トレーニングの防寒対策です。
いじれない場所に保管しておきます。
ただそれを演じさせればいいのです。
結果は使用できませんが、その後のプロセスの前兆となります。
トレーニング後のトレーニングは、小さなバッチがたくさんあります。人々 (「研究者」と呼ばれますが、素朴な率直な言い方で「モデル エンジニア」と呼びます) は、生のモデルでは対処が不十分な特定の問題に注目し、パフォーマンスを向上させる可能性のある微調整を検討します。その結果、現在存在するモデルに適用されるコードとデータの小さな塊 (生き残った実験) が多数生成されます。
トレーニング後は、元のモデルを補足します。十分な量のサプリメントを適用すると、UI、ユーザー、コンピューティングと組み合わせると、「珍しいフォカッチャのトッピングを 5 個ください」と応答できるモデルが完成します。
ポストトレーニングは、モデルトレーニングの形成とクックです。可能性を秘めたものを、人間にとっておいしいものにしていくのです。 (この例えでは、トレーニング後の協力的、反復的、可逆的な性質がカバーされていません。ため息がつきます。)
フォローアップでは、このプロセス全体に関与するさまざまなチームと役割について探っていきます。インセンティブ、ツール、リズム、短期と長期、機能と将来、背景、文化など、興味深い相違点がいくつかあります。
ただし、最初に、プロセスについての理解を再確認したいと思いました。上記で何か間違っている場合はお知らせください。
ほとんどのチームは戦略に問題を抱えていません。彼らは適応の問題を抱えています。
あなたの計画は現実と接触して生き残ることはできませんでした。問題は、組織が曲がらないとき、曲がるか壊れるかということです。
私はチームを助けます

曲げる。繁栄に適応します。
現在、いくつかのカスタム トークやアドバイス業務を予約しています。私は貴社の従業員にインタビューし、実際のソフトウェア フローを測定し、真実とそれに対する対処法をお伝えします。
それが合うかどうか知りたいですか？あなたのチームについて教えてください。
32 1 1 シェア 前 この投稿に関するディスカッション コメント 再スタック Jon Verrier 3h 編集済み 「事前トレーニング」と「事後トレーニング」はありますが、「トレーニング」と呼ばれるステップはありません。 PTに伝えておきます。トレーニング前に集まってチャットをし、トレーニング後のコーヒーに直接移動して、ワークアウトを完全にスキップすることもできます。
返信 シェア トップ 最新のディスカッション 投稿はありません

## Original Extract

I remember walking to the bus from high school, staring at a Motorola 6800 instruction set manual.

Baking a Model - by Kent Beck
Software Design: Tidy First?
Subscribe Sign in Baking a Model
Kent Beck Aug 14, 2026 32 1 1 Share I remember walking to the bus from high school, staring at a Motorola 6800 instruction set manual. I didn’t really understand what I was looking at—boolean expressions, instruction encodings, timing tables—but I was obsessively fascinated by the mechanism of it all. Here was this complicated machine where if I understood it I would have power & control.
I feel the same way about AI models right now. I don’t claim to understand the details, not yet, but I’m fascinated by the mechanism of it all. I’m interested in both:
The machinery that makes a model.
It’s this latter topic, how a model gets constructed, that I will begin to explore in this post (& possible followups).
Saturday’s Focaccia I love baking. You take ingredients in one form & transform them to a totally different form. The ingredients aren’t palatable in themselves but what you create from them is delicious. Oh and also baking is sensitive to initial conditions—you can make a small change early in the process & it will have a large consequence later.
I’ve been experimenting with cold proofing, where you let the yeast do its work overnight in a refrigerator. As I was working to understand models it struck me that there’s an analogy there to creating models, at least as I understand the process so far (please correct me in the comments if I’ve gotten something wrong).
First, though, a progressively revealed story about what we mean by a model. I’ll over-simplify but then reveal more complexity a little at a time.
By “model”, we mean a computer system that can converse with a human (see, told you I’d over-simplify—hope you stay with me, this won’t take long).
Actually it’s a little more complicated than that. The model is split into 2 parts:
A user interface that takes care of formatting inputs & outputs & sequencing & authentication & all that stuff.
The model proper where the magic happens.
The user interface is built using conventional programming techniques (even though “convention” has changed radically in the last 2 years). The model proper, though (our topic of exploration) is built using radically different techniques.
A model is a bag of numbers. For today’s purposes that’s enough (maybe some day we’ll get to how those numbers are used when the model is producing useful words, but I’d have to understand it a little first).
How did we come up with these numbers? Unlike in programming, where you lay out a sequence of statements the result of which is a program, AI models result from training .
Training bears some resemblance to programming—you change an artifact to modify its future behavior—but also some huge differences.
You wouldn’t just magically guess a bunch of correct numbers in one go. Oh no, oh no. First we get some approximately correct numbers & then we tweak them to get even better numbers. The two processes are wildly different.
(Near as I can tell, there’s pre-training, post-training, & mid-training (about which I know nothing), but there’s not “training” except as the composition of pre-, mid-, & post-. Here’s hoping the vocabulary evolves.)
Pre-training is a big batch. The whole team sets up the initial conditions—the data & the blank model. They run the data backwards & forwards through the model a gajillion times. They take snapshots along the way in case of crashes. They check for signs that the pre-training has driven off into the weeds & needs to be tweaked & restarted. Pre-training is a big bet—hundreds of millions of dollars & (more expensively) months of delay.
(I need to learn more about how pre-training folks collaborate.)
Pre-training is the cold proofing of model training.
You put it away somewhere where you can’t mess with it.
You just have to let it play out.
The result isn’t usable but it’s the precursor to the process that follows.
Post-training is lots of little batches. Folks (called “researchers” but in my naive bluntness I’d call “model engineers”) look at particular problems the raw model addresses poorly & explore possible tweaks that might improve performance. The result is lots of little chunks of code & data (the surviving experiments) that apply to the model as it currently exists.
Post-training supplements the original model. Apply enough supplements & you have a model that, when paired with a UI & a user & compute can respond to, “Give me 5 unusual focaccia toppings.”
Post-training is the shaping & cooking of model training. You take something with potential & make it delicious for humans. (The analogy doesn’t cover the collaborative, iterative, & reversible nature of post-training—le sigh.)
In a followup I’m going to explore the different teams & roles involved in this whole process. They have some interesting divergences of incentives, tools, rhythm, short-term vs long-term, feature vs future, backgrounds, & culture.
First, though, I wanted to double check my understanding of the process. Let me know if I got something wrong above.
Most teams don’t have a strategy problem. They have an adaptation problem.
Your plan was never going to survive contact with reality. The question is whether your organization bends or breaks when it doesn’t.
I help teams bend. Adapt to Thrive.
Booking a handful of custom talks and advisory engagements now. I interview your people, measure your real software flows, and hand you the truth plus what to do about it.
Curious whether it fits? Tell me about your team.
32 1 1 Share Previous Discussion about this post Comments Restacks Jon Verrier 3h Edited We have ‘pre-training’ and ‘post-training’, but no step called ‘training’. I will let my PT know. We can meet for a pre-training chat, then move directly to post training coffee, and skip working out entirely.
Reply Share Top Latest Discussions No posts

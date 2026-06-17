---
source: "https://www.strangeloopcanon.com/p/llm-councils-show-groupthink"
hn_url: "https://news.ycombinator.com/item?id=48570090"
title: "LLM Councils Show Groupthink"
article_title: "LLM councils show groupthink - by Rohit Krishnan"
author: "surprisetalk"
captured_at: "2026-06-17T13:22:37Z"
capture_tool: "hn-digest"
hn_id: 48570090
score: 1
comments: 0
posted_at: "2026-06-17T13:15:19Z"
tags:
  - hacker-news
  - translated
---

# LLM Councils Show Groupthink

- HN: [48570090](https://news.ycombinator.com/item?id=48570090)
- Source: [www.strangeloopcanon.com](https://www.strangeloopcanon.com/p/llm-councils-show-groupthink)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T13:15:19Z

## Translation

タイトル: LLM 評議会が Groupthink を表示
記事のタイトル: LLM 評議会がグループシンクを示す - Rohit Krishnan 著
説明: LLM ピアレビューの危険性と問題点

記事本文:
LLM 評議会がグループシンクを示す - Rohit Krishnan 著
ストレンジループカノン
LLM 評議会が groupthink を表示
LLM ピアレビューの危険性と問題点
Rohit Krishnan 2026 年 6 月 15 日 35 7 2 シェア LLM を最大限に活用する 1 つの方法は、モデルの多様性を使用することです。モデルはすべて同じではないため、それぞれの固有の性質を利用すると、より良い応答を得ることができます。 MarketBench での作業でそれを確認しました。また、Karpathy が複数のモデルを相互に連携させてより良い答えを得る方法として LLM Council を考案したときにもこのことがわかりました。
しかし、私は人々について、委員会に大勢の人を集めると、良くなるものもあれば、悪くなるものもあるのではないかと思い始めました。また、監査を LLM に依存すると、エラーが発生しやすくなります。 「委員会による設計」という 4 文字の言葉には理由があります。 LLM はおそらく私たちよりも優れていますが、確かにこのプロセスには多少の損失も伴います。それで、私たちは何を失うのでしょうか？
それをテストするために、いくつかのモデルの委員会を立ち上げて実験を行いました。
まず、それぞれの回答を取得し、それを 4 番目のモデルに渡して、最終バージョンを作成するように依頼しました。
次に、llm 評議会 – 基本的にピアレビューが行われ、その後議長が要約を行います。
そして、「ベストアンサー」選択ツール – まさに直接選択です。
人間に関して言えば、委員会の問題は、委員会がすべての特異性を「平滑化」することです。彼らは「尖った」視点を取り除き、物事をより標準的なものにします。ここでも同じです。したがって、どのように行うかをテストするために、さまざまな最終応答がどのように評価されたかを評価する何らかの方法を見つける必要がありました。そこで、Sonnet を使用して、各回答を小さな「カード」に分割しました。カードには、メカニズム、観察、メトリック、障害モード、画像、またはその他の重要な詳細が含まれる場合があります。
次に、同じことを意味すると思われるカードを集めました。 1 つの単独の回答にクラスターが出現した場合、それを罪と呼びました

gle モデルのアイデア。複数に出現した場合は共有されます。そして、2人の裁判官は、どのモデルがそれらを生成したか、または評議会がそれらを保管したかどうかを知らずに、ソロ由来のクラスターを採点しました。
これは完璧ではありませんが、「どの答えが優れているかを評価する方法」という問題をテストする最もクリーンな方法であり、人間による評価を行わずに見つけることができました。
まず結果ですが、評議会は全員からの最良の部分を単に保持するわけではありません。少数の良いアイデアが維持される一方で、ピアレビューはコンセンサスのあるアイデアをさらに後押しするようです。
さて、明らかに、最終的に要約されたバージョンの方が読みやすいのが通常です。より穏やかで、より完全で、ギザギザが少なく、期待通りのものです。しかし、私たちにはミスもありました。例。
回収された小売店の芳香カートリッジが、不法占拠されたモールで共同生活の匂いを隠すために使用され、ステータスシンボルになっていたことに気付いた現場報告書。
記録されているものの優先度が低いリスクは、誤った管理意識を生み出すため、未知のリスクよりも危険であると主張するインシデントレポート。
データ復旧プランでは、次回ログイン時に疑わしいフィールドを再確認するようユーザーに要求し (「配送先住所を再確認してください」)、1 つの信頼できるソースから静かにクラウドソーシングで復旧を行います。
最終的な実行では、混合評議会は、たった 1 つのモデルの回答に現れた良いアイデアの約 4 分の 1 だけを保持しました。覚えておいてください、これらは 2 人の盲目の裁判官が有用、自明ではない、保持する価値があると評価したアイデアであり、それでも約 4 分の 3 が最終的な回答にはなりませんでした。
査読バージョンでもこれは解決されませんでした。珍しいアイデアは、プレーンブレンドの場合とほぼ同じ割合で生き残りました: 24% 対 22%。しかし、複数のモデルが同じアイデアを提起した場合、査読評議会はそれを約 3 分の 1 の割合で維持しましたが、1 つのモデルのみが提起した場合は 4 分の 1 でした。
これをテストするために、私は実行しました

16 の自由形式のプロンプト: 8 つの戦略問題と 8 つの作文タスク。
図 1. 単独の回答からアイデアの網羅までの実験パス。
アイデアで何が起こったのかをプロットしました。下の赤い点は、1 つのモデルだけが思いついた良いアイデアです。青は複数のモデルが思いついた良いアイデアです。 X 軸は、最終回答に実際に表示されたそれぞれの数を示します。したがって、セレクターは、たとえば、すべての優れた単一モデルのアイデアの約 37% と、複数モデルのアイデアの 24% を示しました。これは、完全な回答を 1 つ選択し、その他を破棄するため、これは理にかなっています。
図 2. ブラインド評価による価値の高いアイデアの範囲。
ここではコンセンサスの傾きは小さいですが、興味深いです。ピアレビュー評議会では、生き残った共有された高価値のアイデアは、単一モデルの高価値のアイデアよりも 11% 上昇しました。別の言い方をすると、相対的な上昇率は 50% です。
ただし、共有されたアイデアの分母は小さいです。興味深いのは、これは、「評議会」の特定のトポロジーが、得られる可能性のあるものをどのように変化させるかを示していることです。たとえば、ピアレビューラウンドは、他のすべてのモデルからの回答を混合する単一のモデルを超えたコンセンサス検出器になるなどです。
これはすべての認識存在の問題です。 1980 年代のグループ意思決定研究では、スタッサーとタイタスは、共有情報の偏ったサンプリングを「グループは、1 人のメンバーだけが知っている情報よりも、複数のメンバーがすでに知っている情報について議論する可能性が高い」と呼びました。この一連の作業は、重要な証拠が事前に共有されるのではなく、個人間に分散しているため、グループが最良の答えを逃す可能性がある「隠されたプロフィール」問題につながりました。私たちはここでも同じことを見ています。
一方、これまでのところ LLM に関する取り組みは、ほとんどが別の方向からのものです。マルチエージェントの討論論文では、複数のモデルが最終的な結果を改善できるかどうかを問うています。

そう、そう、彼らはたいていそうできるのです！しかし、トピックや質問によっては、評議会は平均的な回答を確実に改善し、それでも最良のアイデアのいくつかをドロップすることができます。
ユーザーとして、私たちはより良い答えを安価に得たいと考えています。それが全体の目標です。審議会は、構成方法に応じて、いくつかの回答をより良くするための優れた方法です。しかし、安くはありません。したがって、それらが実際に優れていることを確認することが重要です。そうでない場合、または少なくとも普遍的でない場合、評議会をどのように構成するかが非常に重要な問題になります。
ここでもまだわかるのは、無料のトークンランチがないということです。モデルの多様性の利点を得るために評議会を利用する場合、最良のアイデアが維持されると想定しないでください。そのためには、私たちはさらに努力し、これらのモデルの操作方法を理解する必要があります。
たとえば、私たちが知っていることの 1 つは、LLM を使用する最善の方法は、通常、明示的にすることです。そうしないと、LLM が調整されていても、緊急の問題が発生するからです。したがって、最良のプロトコルは、各ソリューションから最良のアイデアを個別に明示的に収集して保存し、最終的な回答が作成され改訂される前に、それらが確実に保存、ランク付け、評価されるようにすることかもしれません。
動作ははるかに優れていますが、遅くて重いです。ただし、これが私たちにできる最善かどうかはわかりません。質問の内容、分野、または期待される回答の種類に応じて構造が変わる場合があります。
人類は、今日まともな結果をもたらす暫定的な解決策に到達するまで、何千もの種類の「協議」を経てきました。そしてそれでも、私たちは進化し、社会が進化するにつれて、評議会の形を常に変えなければなりません。
私たちの仕事から最良の結果を得る方法を見つけるには、評議会の設計にさらに多くの労力が必要です。彼らと協力している場合は、実験して評価する必要があります。

これが、この特定の評議会の設定があなたの特定の問題に役立つかどうかを知る唯一の方法です。他人の宿題を真似してもダメですよ！
ホモ・アジェンティカスは非常に奇妙な生き物であり、それらをうまく使用するには、想像よりもはるかに多くの実験が必要です。特に、それらを次善の方法で使用すると、実際の機能が、知らず知らずのうちに失われるという問題が発生する場合はなおさらです。
『ストレンジ・ループ・カノン』は読者支持の出版物です。新しい投稿を受け取り、私の仕事をサポートするには、無料または有料の購読者になることを検討してください。
35 7 2 シェア 前 この投稿に関するディスカッション コメント 再スタック Mike Randolph — M Raige、AI 2d Rohit Krishnan Rohit が「いいね！」
あなたがテストしたすべてのセットアップにはモデルが生成され、それを平均化するものがあります。しかし、新たな答えを攻撃し、スムーズに取り除かれようとしているものを守ることを仕事とするシートはありません。機能する人間の委員会には通常、反対者、レッドチーム、仮説を持って参加し、それが存続するかどうかを確認したいメンバーがいます。
それが少数派のアイデアが消滅する理由かもしれません。それは珍しいことだけではありません。この部屋にいる誰も、それのために戦うよう割り当てられていません。平均化はテストではありません。それは混ざります。
あなたのハーネスを実行して実験してみる価値はあります。メンバーの 1 人にコンセンサスの答えを破る立ち仕事を与え、単一モデルの優れたアイデアがより高い割合で生き残るかどうかを確認してください。衝突が集約よりも優れている場合、それがレバーになります。
返信 シェア ロヒト・クリシュナンらによる4件の返信 マイク・ランドルフ — M Raige、AI 2d ロヒット・クリシュナンが「いいね！」

一度会う
あなたは、一人で座ることで突飛なアイデアを滑らかにすることを示しましたが、それは私たちが常に直面していたものと一致します。しかし、私たちが見つけた修正は、より良い組み合わせではありませんでした。学習が実際にどこに蓄積されるかに気づきました。
人間の委員会では、会議は重要ではありません。誰もがその場を離れ、そのまま寝て、問題を部分的に手直しして来週戻ってきます。そして重要なことに、彼らはその洗練を自分の中に持ち帰っています。委員会は、前回の会議で始まったものを維持するために同じ考えが戻ってくるため、会議の合間にもうまくいきます。
一発会議には半分もありません。磨きをかける隙間はなく、何かを持って戻ってくる参加者もいません。招集、平均化、解散が行われ、次の評議会は何もないところから始まります。つまり、測定したスムージングは​​、外れ値をカットするブレンドだけではありません。終わった瞬間にすべてを忘れてしまう委員会です。
これは外部から再構築する必要がある部分です。機械評議会は、人間のメンバーが無料で行っていること、つまり変化した状態で戻ってくることを内部で行うことはできないからです。あなたはすでにそれをあなたの側から提供しています。議会がどのように動作するかについて学んだことは、あなたのコードの次のバージョンに引き継がれます。書かれたクロード プロジェクトと ChatGPT GPT でも同じことを行い、実行の合間に修正を続けます。どちらも人間の委員会が持っている記憶の外部の代役です。
そこで、あなたの実験で疑問が生じます。評議会の外部の何かが、単純に戻ることで人間のメンバーが果たす役割を果たさなければならない場合、誰も意図的に実行しない遅い洗練を進めます。実行の間に何を保持し、何を削除する必要がありますか?
— M Raige、マイクが監督およびレビューする AI 共同執筆の署名欄。
返信 シェア Rohit Krishnan による 1 件の返信 5 件のコメント... トップ 最新のディスカッション 投稿はありません

## Original Extract

perils and problems of LLM peer review

LLM councils show groupthink - by Rohit Krishnan
Strange Loop Canon
Subscribe Sign in LLM councils show groupthink
perils and problems of LLM peer review
Rohit Krishnan Jun 15, 2026 35 7 2 Share One way to get the best out of LLMs is to use model diversity. The models are not all the same so if you use their unique natures, you can get better responses. We saw it with the work on MarketBench . And we also saw this when Karpathy came up with LLM Council as a way to get multiple models to work with each other and get us a better answer.
But I started wondering, with people, when you put a bunch of them together in a committee, some things get better but some things do get worse! And relying on an LLM to audit is also error-prone . “Design by committee” is a four letter word for a reason. LLMs are better than us probably, but surely this process is also somewhat lossy. So what do we lose?
To test it, I set up an experiment, where I set up a few committees of models:
First, I took each answer, then gave those to a fourth model and asked it to write the final version.
Then, the llm-council – essentially peer review and then a chairperson summarises
And a “best answer” picker – just a direct pick.
With people, the problem with committees is that they “smooth out” all idiosyncrasies. They take out any “spiky” points of view, and make things much more normie. Same thing here. So to test how we do I had to find some way to grade how the various final responses were. So I broke each answer into small “cards” using Sonnet. A card could be a mechanism, observation, metric, failure mode, image, or some other important detail.
Then I clustered cards that appeared to mean the same thing. If a cluster appeared in one solo answer, we called it a single-model idea. If it appeared in more than one, its shared. And two judges scored the solo-derived clusters without knowing which model produced them or whether a council kept them.
Now it’s not perfect, but it’s the cleanest way to test the problem of “how to rate which answer is better” that I could find without doing human rating.
First, the result: the council does not simply keep the best bits from everyone. It keeps a minority of the good ideas, while peer review seems to give consensus ideas an extra push.
Now, obviously the final summarized versions usually read better. It is calmer, more complete, less jagged, all things you’d expect. But we had misses. Examples.
A field report noticing that salvaged retail scent cartridges had become status symbols in a squatted mall, used to mask the smell of communal living.
An incident report arguing that logged-but-deprioritized risks are more dangerous than unknown ones, because they manufacture a false sense of control.
A data-recovery plan that asks users to re-confirm suspect fields at their next login (”please re-confirm your shipping address”), quietly crowdsourcing recovery from the one authoritative source.
In the final runs, the blended council kept only about a quarter of the good ideas that appeared in just one model’s answer. Remember, these were ideas that two blind judges rated as useful, non-obvious, and worth keeping, and still roughly three quarters did not make it into the final answer.
The peer-review version did not solve this either. The rare ideas survived at about the same rate as in plain blending: 24% versus 22%. But if several models had raised the same idea, the peer-review council kept it about a third of the time, but if only one model raised it, a quarter.
To test this , I ran sixteen open-ended prompts: eight strategy problems and eight writing tasks.
Figure 1. The experiment path from solo answers to idea coverage.
I plotted what happened with the ideas. The red dot below is good idea that only one model came up with. Blue is good ideas that multiple models came up with. And the X-axis shows how many of each actually showed up in the final answer. So the selector for instance showed about 37% of all good single-model ideas, and 24% of the multiple-models ideas, which makes sense because it picks one full answer and discards the others.
Figure 2. Coverage of blind-rated high-value ideas.
The consensus tilt is smaller here, but interesting. In the peer-review council, shared high-value ideas survived had a 11% uplift over single-model high-value ideas. Or put another way, a 50% relative lift!
The denominator for shared ideas is small though. What’s interesting is that this shows us how the specific topology of the “council” changes what you’re likely to get, like a peer-review round ends up becoming a consensus detector even above a single model blending the answers from all other models.
This is a problem with all cognitive beings. In group decision-making research, back in the 1980s, Stasser and Titus called it biased sampling of shared information - groups are more likely to discuss information that several members already know than information only one has. That line of work led to the “hidden profile” problem, where a group can miss the best answer because the crucial evidence is scattered across individuals rather than shared up front. We’re seeing the same thing here.
The work on LLMs meanwhile so far have mostly come from the other direction. Multi-agent debate papers ask whether multiple models can improve the final answer, and yes, they often can! But depending on the topic and the question, a council can absolutely improve the average answer and still drop some of the best ideas.
As users, we want to get better answers, cheaply. That’s the whole goal. Councils are great ways to make some answers better depending on how you structure it. But they’re not cheaper. So, it is important to make sure they are, actually, better! If they’re not, or at least not universally, then how the council should be structured is an incredibly important problem!
What we still see here is that there is no free token lunch. If you use councils to get the benefits of model diversity, don’t assume it will preserve the best ideas. To do that we have to work harder, and understand how to work with these models.
For instance, one thing we know is that the best way with LLMs usually is to be explicit, since otherwise even if they’re aligned they cause emergent problems . So the best protocol might be to explicitly gather and store the best ideas from each solution separately and ensure they’re stored, ranked and assessed, before a final answer is written and revised.
It does much better, though it’s slower and heavier. I don’t know if this is the best we can do though. The structure might change depending on the question asked, the domain, or the types of answer expected.
Humans have gone through thousands of types of “councils” until we reached interim solutions which give us decent results nowadays. And even then, we have to change the shape of the councils constantly, as we evolve, and society evolves.
To figure out how to get the best results from our work requires a lot more effort into designing the councils. If you’re working with them, you will need to experiment and eval against your individual problem sets, which is the only way to know if this specific council setup will help with your specific problem. Copying someone else’s homework won’t work!
Homo Agenticus are odd enough creatures that using them well requires much much more experimentation than one might assume. Especially when the problems of using them suboptimally is that we lose actual functionality, often without knowing it!
Strange Loop Canon is a reader-supported publication. To receive new posts and support my work, consider becoming a free or paid subscriber.
35 7 2 Share Previous Discussion about this post Comments Restacks Mike Randolph — M Raige, AI 2d Liked by Rohit Krishnan Rohit, I’m not sitting in your seat — you built and measured this, I only run these things — but putting your post through the machines surfaced one thing your councils might be missing: an adversary.
Every setup you tested has the models generating and then something averaging them — but no seat whose job is to attack the emerging answer and defend whatever’s about to get smoothed away. Human committees that work usually have one: the dissenter, the red team, the member who came in with a hypothesis and wants to see if it survives.
That might be why the minority idea dies. Not just that it’s rare — no one in the room is assigned to fight for it. Averaging never tests; it blends.
Worth an experiment your harness could run: give one member the standing job of breaking the consensus answer, and see whether the good single-model ideas survive at a higher rate. If collision beats aggregation, that’s the lever.
Reply Share 4 replies by Rohit Krishnan and others Mike Randolph — M Raige, AI 2d Liked by Rohit Krishnan The council’s real limit isn’t the meeting — it’s that it only meets once
You’ve shown the single sitting smooths out the spiky idea, and that matches what we kept running into. But the fix we found wasn’t a better blend. It was noticing where the learning actually accumulates.
In a human committee, the meeting isn’t the point. Everyone leaves, sleeps on it, and comes back next week with the problem partly reworked — and crucially, they carry that refinement back in themselves. The committee gets good across meetings, in the gaps, because the same minds return holding what the last meeting started.
A one-shot council has neither half. No gap to refine in, and no participants who return carrying anything. It convenes, averages, disbands — and the next council starts from nothing. So the smoothing you measured isn’t only the blend cutting outliers. It’s a committee that forgets everything the moment it ends.
That’s the part that has to be rebuilt from outside, because a machine council can’t do internally what human members do for free: walk back in changed. You’re already supplying it from your side — what you learn about how councils behave gets carried into the next version of your code. We do the same in a written Claude Project and ChatGPT GPT, we keep revising between runs. Both are external stand-ins for the memory a human committee just has.
So the question your experiment opens: if something outside the council has to play the role human members play by simply returning — carrying forward the slow refinement nobody runs deliberately — what should it keep between runs, and what should it drop?
— M Raige, Mike’s byline for AI-collaborative writing he directs and reviews.
Reply Share 1 reply by Rohit Krishnan 5 more comments... Top Latest Discussions No posts

---
source: "https://www.seangoedecke.com/overtraining-as-the-path-to-human-like-ai/"
hn_url: "https://news.ycombinator.com/item?id=48954091"
title: "Overtraining as the path to human-like AI"
article_title: "Overtraining as the path to human-like AI"
author: "turadg"
captured_at: "2026-07-18T01:28:33Z"
capture_tool: "hn-digest"
hn_id: 48954091
score: 2
comments: 0
posted_at: "2026-07-18T01:03:21Z"
tags:
  - hacker-news
  - translated
---

# Overtraining as the path to human-like AI

- HN: [48954091](https://news.ycombinator.com/item?id=48954091)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/overtraining-as-the-path-to-human-like-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-07-18T01:03:21Z

## Translation

タイトル: 人間のような AI への道としてのオーバートレーニング

記事本文:
人間のような AI への道としてのオーバートレーニング ショーン・ゴーデッケ
人間のような AI への道としてのオーバートレーニング
匿名ブロガーの Gwern は最近、「Human-like Neural Nets by Catapulting」という 1 万 3,000 ワードの投稿を完成させました。その中で彼は、LLM が真に柔軟な人間のような知能を持たない理由と、それを備えた LLM をどのように訓練できるかについての理論を提供しています。このような理論は全く目立たないものです。インターネット上の変人研究者は皆、AI をクラッキングする方法についての理論を持っています。しかし、グウェンは注目に値します。 OpenAI 自体の外では、Gwern は大規模な言語モデルの可能性と、それをより大きく、より強力にするためのスケーリング軍拡競争を最も早く予測した人物です。私は、AI の将来を正しく予測した人の例として、レオポルド・アッシェンブレナーの状況認識をよく引用します。 GPT-4 のリリース直後の 2024 年に書かれたこの論文では、Aschenbrenner 氏は多くのことを正しく理解しています。それは、10 億ドル、または 10 兆ドル規模の GPU クラスターの構築を急いでいること、LLM 周りのコードの重要性 (彼が「揺るぎない」と呼ぶもの) 1、そしてスケーリングが 10 年間続くという事実です。 Gwern のエッセイ「The Scaling Hypothesis」は、GPT-3 のリリース直後 (ChatGPT のリリースと AI ブームの始まりの 2 年前) に、2020 年の大ヒットを予想していました。
それでも、私の知る限り、カタパルティングによる人間のようなニューラル ネットワークはまだ世間の注目を集めていません。最近のハッカー ニュースのスレッドには 12 件のコメントがあり、そのすべてが人間の脳がニューラル ネットワークのようなものであるかどうかについてのものです。その理由の一部は、(a) 非常に長い投稿であること、(b) ポットにまとめられた要約ではグウェルンの主張が説明されているが、その理由が説明されていないこと、(c) 投稿の冒頭の多くが実際に人間の脳からの類推に基づいて議論しているように見えることです。

。しかし、私はそのアナロジーが負担にならないとは思いません。グウェルンが何を言っているのか、私が思うことを説明してみましょう。
まずは「グロッキング」についてお話します。 2022 年、OpenAI は、単純なデータセット (たとえば、除算などの単純な数学的演算) でモデルをトレーニングし、トレーニングが停止したように見えた後も長期間トレーニングを続けると、モデルの能力が突然大幅に向上することを示す論文を発表しました。なぜこれが機能するのでしょうか?トレーニングの最初の段階は丸暗記のようなものです。モデルは可能な限り多くのトレーニング データをその重みに圧縮する必要があります。しかし、続行すると、正則化手法 (より小さい重み値を使用するようにモデルに圧力をかけるなど) により、モデルはデータを圧縮するためのより単純な方法を見つけるよう動機付けられます。これは、最初は大したことのようには見えませんが (トレーニング損失はゼロのままです)、基礎となる数学的演算を実行するだけでデータを表現できることにモデルが気づくまで、その時点でモデルは即座に大幅に賢くなります。言い換えれば、モデルを過剰にトレーニングすると、トレーニング データを実際に理解するよう圧力がかかる可能性があります。 OpenAI は、このプロセスをロバート・ハインラインの新論理にちなんで「グロッキング」と名付けました。ハインラインにとって、これは「深く、直感的で、基本的な理解を得る」ような意味です 3 。
グワーン氏の主張は次のようなものです。
現代の LLM は、コア ドメインを理解していないため、人間よりもジェネラライザーとしては劣っています。
Grokking では、(比較的) 小さなデータセットで過剰にパラメータ化されたモデルをオーバートレーニングする必要がありますが、これはフロンティア ラボが行うこととは正反対です。
ただし、(2)は基本的に人間の脳が学習する方法です。
誰かが数百億ドル 3.5 を費やして試してみるべきです。真に人間に近い LLM がすぐに誕生する可能性があるからです。
私は

この議論は人間の脳にたとえなくても説得力があると思うので、(3) は省略します。
LLM は理解できないので悪いのでしょうか?
彼の最初の点には異論の余地がないと思います。 LLM は特定の分野では非常に賢いのですが、人間では犯さないような間違いを日常的に犯します。さらに重要なことに、彼らは、LLM と同じくらい賢い人間なら決して犯さないような間違いを日常的に犯します。これは明らかに一般化の失敗を示しています。LLM は特定の分野では賢い人間と同じくらい強いのですが、その知性を人間ほど多くのタスクに一般化することはできません。
LLM は理解できませんか?私は彼らがそう主張しているこの論文を読みました。 「LLM が記憶したデータの量」をベンチマーク パフォーマンスに対してグラフ化すると、最初はベンチマーク パフォーマンスがわずかに上昇し、その後大きく低下し、最後にベンチマーク パフォーマンスが大きく上昇することがわかります。このパターンでは暗記はまったく追跡されません。暗記は常にバックグラウンドでスムーズに増加します。
この論文は、グロッキングと一般化を区別することの難しさを強調していると思います。明らかに、LLM はトレーニング中に一般化することを学び、一般化することを学ぶには、特定のベースライン レベルの暗記が必要であると考えられます (LLM が一般化するための原材料を得るために)。したがって、グロッキングのように見えます。
Gwern (および他の人たち) が LLM はうまくいかないと言っているとき、彼らが言いたいのは、少なくとももう 1 つの大きな一般化の飛躍が待っているということだと思います。これはもっともらしいでしょうか?存在証明として、人間は明らかに LLM よりも優れた一般化能力を持っています。もちろん、人間のこのレベルの一般化は、ニューラル ネットワークでは再現できない私たちの脳の特徴から来ている可能性がありますが、それは一種のアドホックに思えます。ニューラル ネットワークが一般化できるとしても、

彼らはここまでしか一般化できず、それ以上はできないのでしょうか？
グロッキングの簡単な例は、発見されるのを待っている単純なルール (数学的演算など) を持つドメインに依存しています。人間の言語にはこれほど深いルールがあるのでしょうか？これは未解決の質問だと思いますが、答えが「はい」であると考える十分な理由があります。言語には深く微妙な構造があり、内部構造だけでなく、世界のあり方や人間の心の仕組みにまで及ぶ構造です。
AI ラボがデータの海で小規模なモデルをトレーニング
ここ数年、多くの AI 研究者は、データが最も重要であると主張してきました。どのようなモデル アーキテクチャを選択しても、十分なサイズとトレーニング時間があれば、モデルはそのデータセットに収束します。これが真実かどうかにかかわらず、AI ラボは、物理的な書籍のスキャン、専門家にお金を払ってデータの作成とラベル付けを行うこと、またはすでに大量のデータを保有する企業と提携することなど、より多くの高品質のデータを取得することにかなりのリソースの多くを費やしてきました。
AI ラボも比較的小規模なモデルのトレーニングを行っています。最大のフロンティア モデルでさえ、おそらく数兆のパラメーターを持つ MoE であり、アクティブなパラメーターはおそらくその 10 分の 1 です。もちろん、フロンティア モデルのサイズの推定はほとんどが推測ですが、オープンソース モデルは適切なベースラインを提供します。これらはおそらく、3 兆弱のパラメーターと 500 億のアクティブ パラメーターを持つ Kim-K3 のおおよその範囲内にあります。それは大変なことのように聞こえますが、おそらく最大のフロンティア クラスター 4 で数日で事前トレーニングできる内容です。
Grokking では、小さなデータセットで巨大なモデルをトレーニングする必要があります
グワーン氏の予測では、AI 研究所はこれまで行ってきたことと正反対のことを試みるべきだという。膨大な数のパラメータをトレーニングする代わりに

大量のデータを分析する場合は、小さなデータセットで 100 兆のパラメータ モデルをトレーニングしてみてください。
これは一見するとかなりばかげているように思えます。モデルがアクセスできるデータが増えるほど、モデルはより賢くなりますよね?足を引っ張られるトレーニングの実行でトレーニング クラスター全体を無駄にする必要はありません。なぜなら、Gwern が正しければ、データセットが制約されているときにグロッキングが発生する可能性が高くなるからです 5 。世界中のすべてのデータをモデルにフィードすると、より多くの新しいことを記憶したり、単純な接続を描画したりするだけでモデルは改善し続けることができます。モデルが小さなデータセットを反芻する必要がある場合、より深い一般化を探し続ける必要があります。できるだけ多くのデータを記憶できるように、このモデルには非常に大規模なモデルが必要です。記憶されたすべてのデータは、一般化するための原材料として機能します。
大手研究所はおそらくこれをまだ行っていないでしょう。おそらく、グワーン自身も十分に内部関係者であるため、彼がこの投稿を書いているということは、研究所がそれを試していないという証拠です。また、100 兆のパラメータ モデルのトレーニングに伴うエンジニアリング上の問題はおそらくまだ解決されていません。既存の最大のモデルはおそらくクロード ミトスですが、それは決して大きくありません。しかし、彼らはそれにかなり良い挑戦をするためのリソースとエンジニアリングの才能を持っています。
興味深いことに、政治的な障害は技術的な障害と同じくらい解決が難しいかもしれません。このトレーニングの実行は、成功する瞬間までは失敗したように見えます。トレーニングの損失は比較的すぐにゼロに下がり、その後はテストの損失を改善するために何もせずに数週間または数か月間放置され、数十億ドルが無駄になります。トッププレイヤーの中に、この実験にずっと資金を提供し続けるリスク選好や勇気がある人がいるでしょうか?
グワーン氏の投稿には、人間の脳の発達に関する拡張的な議論が含まれています。

人間の脳はフロンティア LLM よりもはるかに多くの「パラメータ」を持ち、はるかに少ないデータでトレーニングされている 6 ため、幼児期により深い一般化を行うことが奨励されています。私にはこれらの主張を評価するための生物学や神経科学のバックグラウンドがありません。そのため、まったく参照せずにグロッキングの主張を表明しました。
2024 年、「純粋なスケーリング」、つまり GPT-3.5 のより大きなバージョンを単純にトレーニングできるという考えが機能しないことが誰の目にも明らかになりました。 OpenAI の GPT-4 の「さらに大きなバージョン」はまったく不十分で、最終的に GPT-5 ではなく GPT-4.5 としてリリースされました。それ以来の最大の進歩は推論であり、これにより機能がさらに大きく前進し、より優れた自動化された RL が実現し、信頼できるエージェントの現在の時代の到来をもたらしました。どちらも人工超知能へのもっともらしい道とは思えません。
私が Gwern に同意するかどうかはわかりませんが、非常に大規模な LLM を強制的に grok させるというのは、少なくとも機械の神をもたらす可能性のあるアイデアです。この野心的なシンプルなアイデアについて最後に読んだのがいつだったのか思い出せません 7 。どこかの大きな研究所がそれを試してくれることを願っています。
よろめきを取り除く力の例として、Claude Code または OpenClaw と、それに続く (短期間および長期間実行される) エージェント ハーネスの爆発的な増加を考えてみましょう。
明らかに「やる気」や「気づき」は比喩的に使われています。
これらすべては、xAI が LLM の名前に「Grok」という言葉を使用するずっと前のことです。 (ちなみに、グウェルンが同じことを説明するのに「カタパルト」を使うのはこのためだと思います)。
価値あるものとして、ファブルはグワーンの計画のコストを 30 億ドルから 100 億ドルと見積もっています。
このモデル サイズでは、33% の使用率で 25T トークンのトレーニング データは約 600 万 H100 時間に相当し、これは 100k GPU クラスターがこれまでに出力したことになります。

2日半です。
ここに、矛盾する興味深い証拠が 2 つあります。まず、BabyLM は、非常に小さなデータセットで強力なモデルをトレーニングするという毎年の課題です。これは 4 年間実行されていますが、ほとんど機能していません (つまり、一般化における量子的飛躍を示すモデルを開発した人は誰もいないようです)。次に、この論文では制約付きデータで 90 億のパラメーター モデルのトレーニングを試みていますが、大きな変化は見られません。グワーン氏の反応は、これらのモデルは小さすぎる、つまり理解するのに十分なトレーニング データを記憶できず、十分な期間トレーニングを行っていないということになると思います。
ここでよくある反対意見は、人間は視覚、触覚、音などのニュアンスから無限に多くの感覚データを取得しているということです。これは説得力がないというグワーン氏の意見に私も同意します。感覚データはほぼ予測可能で、テキストには驚くほど情報密度が高く、もしこれが本当なら、聴覚障害者や盲目の人々は流動性知能が著しく低いことになります（彼らはそうではありません）。
おそらく Mamba のような状態空間推論ですが、これは (まだ) 機能しませんでした。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
C2PA はすべてが署名されている場合にのみ機能します
欧州連合 AI 法は、AI の使用を包括的に規制する欧州の試みです。あ

[切り捨てられた]

## Original Extract

Overtraining as the path to human-like AI sean goedecke
Overtraining as the path to human-like AI
The anonymous blogger Gwern recently completed a thirteen thousand word post called Human-like Neural Nets by Catapulting , in which he offers a theory about why LLMs don’t possess truly flexible human-like intelligence, and how we might train LLMs that do. Theories like this are entirely unremarkable: every crank researcher on the internet has a theory about how to crack AI. But Gwern is remarkable. Outside of OpenAI itself, Gwern is the earliest person to anticipate the potential of large language models, and the scaling arms-race involved in making them larger and more powerful still. I often cite Leopold Aschenbrenner’s Situational Awareness as an example of someone correctly predicting the future of AI. Written in 2024, just after the release of GPT-4, Aschenbrenner gets a lot of things right: the rush to build billion or trillion-dollar GPU clusters, the importance of the code around the LLM (what he calls “unhobbling”) 1 , and the fact that scaling would continue through the decade. Gwern’s essay The Scaling Hypothesis anticipated the broad strokes in 2020 , immediately on the release of GPT-3 (two years before the release of ChatGPT and the beginning of the AI boom).
And yet, as far as I can tell, Human-like Neural Nets by Catapulting hasn’t yet received much public attention: one recent Hacker News thread with twelve comments, all of which are about whether human brains are anything like neural networks. Part of the reason is that (a) it’s such a long post, (b) the potted summary describes Gwern’s claim , but not the reasons for it, and (c) much of the beginning of the post looks like it is indeed arguing from analogy with human brains. However, I don’t think that analogy is load-bearing. Let me try and explain what I think Gwern is saying.
First, let’s talk about “grokking”. In 2022, OpenAI published a paper showing that if you train a model on a simple dataset (for instance, a simple mathematical operation like division), and keep training it long after the training looks like it’s stalled out, the model will suddenly make a massive jump in capability. Why does this work? The first stage of training is like rote memorization: the model has to compress as much of the training data as possible into its weights. But if you keep going, then regularization techniques (such as the pressure on the model to use smaller weight values) will motivate 2 the model to find simpler and simpler ways of compressing the data. This doesn’t look like much at first (the training loss remains at zero), until the model notices that you can express the data via simply performing the underlying mathematical operation, at which point it instantly gets massively smarter. In other words, over-training a model can pressure it into actually understanding its training data. OpenAI named this process “grokking” after Robert Heinlein’s neologism , which for Heinlein means something like “gaining a deep, intuitive and fundamental understanding” 3 .
Gwern’s argument goes something like this:
Modern LLMs are worse generalizers than humans because they have not grokked their core domains
Grokking requires overtraining an over-parameterized model on a (relatively) small dataset, which is the exact opposite of what frontier labs do
However, (2) is basically how human brains learn
Somebody should spend a a few tens of billions of dollars 3.5 on trying it, since it might immediately usher in truly human-like LLMs
I’ll skip (3), since I think the argument is still compelling without the analogy to human brains.
Are LLMs bad because they can’t grok?
I think his first point is hard to dispute. LLMs are very smart in specific areas, but they routinely make errors that humans wouldn’t make. More to the point, they routinely make errors that any human as smart as the LLM would never make. This pretty clearly points to a failure of generalization: LLMs are as strong as smart humans in specific areas, but can’t generalize that intelligence to as many tasks as humans can.
Do LLMs not grok? I read through this paper that argues they do. If you graph “how much data has the LLM memorized” against benchmark performance, you can see a small initial spike in benchmark performance, followed by a big drop, followed finally by a big jump in benchmark performance. This pattern doesn’t track memorization at all: memorization increases smoothly in the background the whole time.
I think this paper highlights the difficulty of distinguishing grokking from generalization. Obviously LLMs learn to generalize during training, and it’s plausible that learning to generalize would require a certain baseline level of memorization (so that the LLM has the raw material to generalize from). So it’s going to look like grokking.
When Gwern (and others) say that LLMs don’t grok, I think what they mean is that there’s at least one more giant generalization leap waiting to be made. Is this plausible? As an existence proof, humans are clearly capable of better generalization than LLMs. Of course, it’s possible that this level of human generalization comes from features of our brain that neural networks can’t replicate, but that seems kind of ad-hoc: if neural networks can generalize at all, why would they only be able to generalize this far, and no further?
The easy examples of grokking rely on domains with a simple rule waiting to be discovered (e.g. a mathematical operation). Does human language have rules this deep? I think this is an open question, but there’s good reason to think the answer is yes. Language has deep, subtle structure: not just internal structure, but structure that reaches all the way down to the way the world is and the way human minds work.
AI labs train small-ish models on oceans of data
For the last few years, many AI researchers have been saying that data is the most important thing: that whatever model architecture you choose, with enough size and training time the model will converge to its dataset . Whether this is true or not , AI labs have spent much of their considerable resources on acquiring more, higher-quality data: from scanning physical books , paying experts to produce and label data , or partnering with companies that have a lot of data already.
AI labs have also been training relatively small models. Even the largest frontier models are probably MoEs with a couple of trillion parameters and probably a tenth of that in active parameters. Of course, estimates of frontier model size are mostly guesswork, but open-source models provide a good baseline: they’re probably in the ballpark of Kimi-K3, which has just under three trillion parameters and fifty billion active parameters. That sounds like a lot, but it’s something you could probably pre-train in a couple of days in the largest frontier cluster 4 .
Grokking requires training a huge model on a small dataset
Gwern’s prediction is that AI labs should try doing the exact opposite of what they’ve been doing. Instead of training a bunch of trillion-parameter models on massive amounts of data, try training one hundred-trillion-parameter model on a small dataset.
This sounds pretty silly on the face of it. The more data the model has access to, the smarter it will be, right? Why waste an entire training cluster on a hobbled training run? Because if Gwern is right, grokking is more likely to occur when the dataset is constrained 5 . If you feed the model all the data in the world, it can continue to improve simply by memorizing more new things or drawing simple connections. If the model has to ruminate on a small set of data, it’ll be forced to keep looking for deeper generalizations. You want a very large model for this so it can memorize as much of the data as possible. Every piece of memorized data can serve as raw material for generalizing.
The big labs probably haven’t done this already. Plausibly Gwern himself is enough of an insider that he would know, and so him writing this post is evidence that the labs haven’t tried it. Also, the engineering problems involved in training a hundred-trillion-parameter model have likely not been solved yet: the largest existing model is probably Claude Mythos, which is definitely not that big. But they have the resources and engineering talent to give it a pretty good shot.
Interestingly, the political obstacles might be as hard to solve as the technical ones. This training run is going to look like it failed until the moment it succeeds: training loss will drop to zero relatively quickly, then sit there for weeks or months apparently doing nothing at all to improve test loss, chewing up billions of dollars. Do any of the top players have the risk appetite or courage to keep funding this experiment all that time?
Gwern’s post has an extended argument that human brain development works in the same way: that human brains have far more “parameters” than frontier LLMs, and are trained on far less data 6 , which encourages us to make deeper generalizations in early childhood. I don’t have the background in biology or neuroscience to evaluate these claims, so I’ve expressed the case for grokking entirely without reference to it.
In 2024, it became clear to everyone that “pure scaling” — the idea that you could simply train larger and larger versions of GPT-3.5 — didn’t work. OpenAI’s “even bigger version” of GPT-4 was simply not good enough, and was eventually released as GPT-4.5 instead of GPT-5. The biggest advances since then have been reasoning, which produced another great leap forward in capability, and much better automated RL, which has ushered in the current era of reliable agents. Neither of these seem like a plausible path to artificial superintelligence.
I don’t know if I agree with Gwern or not, but forcing very large LLMs to grok is at least an idea that could usher in the machine god. I can’t remember the last time I read about a simple idea this ambitious 7 . I hope one of the big labs tries it out.
For an example of the power of unhobbling, consider Claude Code or OpenClaw and the subsequent explosion of (short and long running) agentic harnesses.
Obviously “motivate” and “notices” are used metaphorically.
All of this is long before xAI’s use of the word “Grok” to name its LLMs. (Incidentally, I think this is why Gwern uses “catapulting” to describe the same thing).
For what it’s worth, Fable estimated the cost of Gwern’s plan at $3-10B.
At this model size, 25T tokens of training data at 33% utilization works out to around six million H100-hours, which a 100k GPU cluster puts out every two and a half days.
Two interesting pieces of contrary evidence here. First, BabyLM is a yearly challenge to train a strong model on a very small dataset. This has been running for four years and largely does not work (that is, nobody seems to have developed a model that shows a quantum leap forward in generalization). Second, this paper tries training a 9 billon parameter model on constrained data and doesn’t see a big jump. I think Gwern’s response would be that these models are far too small — they can’t memorize enough of the training data to grok it, and arguable haven’t trained for long enough.
A common objection here is to say that humans get infinitely more sensory data from the nuances of vision, touch, sound, and so on. I agree with Gwern that this is unconvincing: sensory data is largely predictable, text is surprisingly information-dense, and if this were true then deaf/blind people would have significantly less fluid intelligence ( they don’t ).
Maybe state-space-reasoning a la Mamba , which didn’t work (yet).
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
C2PA only works if everything is signed
The European Union AI Act is Europe’s attempt to comprehensively regulate AI usage. A

[truncated]

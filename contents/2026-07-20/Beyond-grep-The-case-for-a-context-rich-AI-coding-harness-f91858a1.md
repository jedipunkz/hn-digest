---
source: "https://arstechnica.com/ai/2026/07/beyond-grep-the-case-for-a-context-rich-ai-coding-harness/"
hn_url: "https://news.ycombinator.com/item?id=48982886"
title: "Beyond grep: The case for a context-rich AI coding harness"
article_title: "Beyond grep: The case for a context-rich AI coding harness - Ars Technica"
author: "emmabot43"
captured_at: "2026-07-20T19:26:01Z"
capture_tool: "hn-digest"
hn_id: 48982886
score: 2
comments: 0
posted_at: "2026-07-20T18:29:07Z"
tags:
  - hacker-news
  - translated
---

# Beyond grep: The case for a context-rich AI coding harness

- HN: [48982886](https://news.ycombinator.com/item?id=48982886)
- Source: [arstechnica.com](https://arstechnica.com/ai/2026/07/beyond-grep-the-case-for-a-context-rich-ai-coding-harness/)
- Score: 2
- Comments: 0
- Posted: 2026-07-20T18:29:07Z

## Translation

タイトル: grep を超えて: コンテキストリッチな AI コーディング ハーネスのケース
記事のタイトル: grep を超えて: コンテキストリッチ AI コーディング ハーネスのケース - Ars Technica
説明: Augment Code の Vinay Perneti がモデル、ハーネス、コンテキストについて語ります。

記事本文:
コンテンツにスキップ
アルス テクニカ ホーム
セクション
フォーラム
購読する
検索
AI
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ストーリーにピンを付ける
テーマ
ハイパーライト
検索
サインイン
サインインダイアログ...
サインイン
薬剤の開発
grep を超えて: コンテキストリッチな AI コーディング ハーネスの事例
Augment Code の Vinay Perneti がモデル、ハーネス、コンテキストについて語ります。
93
Augment Code のエンジニアリング担当副社長である Vinay Perneti 氏がプレゼンテーションで語ります。
クレジット:
コードを拡張する
Augment Code のエンジニアリング担当副社長である Vinay Perneti 氏がプレゼンテーションで語ります。
クレジット:
コードを拡張する
テキスト
設定
ストーリーテキスト
サイズ
小
標準
大
幅
*
標準
ワイド
リンク
標準
オレンジ
※購読者限定
さらに詳しく
ナビゲーション用に最小化する
世の中には AI コーディング アプリケーションが数多くあり、大規模な言語モデルとそれが可能にするエージェントが印象的であるのと同じくらい、AI 支援開発における最新の開発の多くは、モデルそのものだけでなく、それらのモデルを管理するソフトウェアにも反映されています。
この夏の初め、私は Claude Code の製品責任者である Anthropic の Cat Wu に、そのソフトウェアを構築するための同社のアプローチについて話しました。
ウー氏は私たちの会話の中で同じ点に繰り返し戻ってきました。Anthropic のモデル (およびその直接の競合他社のモデル) は急速に改善されているため、あまり先を見据えた計画を立てたり、独自の機能や制限的な機能を構築したりすることはほとんど意味がありません。むしろ、Claude Code 製品チームは、いわゆる無駄のないハーネスを維持しようとしています。
ハーネスは、1 つ以上の AI モデルを中心に構築され、その使用方法を決定するソフトウェアです。これにより、モデルが何を認識するか、どのようなアクションを実行できるか、およびモデルがコード ベースとどのように対話するかが決まります。モデルと開発者の間のレイヤーのように考えてください。

実際のプロジェクト。
アプリケーションおよびハーネスとしての Claude Code に独自の機能やデザインの選択肢がないわけではありませんが、その製品チームは、来年モデルがどのような方向に進むかを信頼するという側面で誤りを犯しているようです。
ただし、クロードコード以外にもハーネスはあります。 OpenAI の Codex、Google の Antigravity、OpenCode のようなオープンソースの代替手段、および Cursor や Augment Code のようなスタートアップによるオプションがあります。それぞれに、エージェント ワークフローでモデルをどのように利用できるか、または利用すべきかについて、異なる機能、強調点、意見がある場合があります。
Claude Code が行う重要な選択の 1 つは、デフォルトでコードベースの周囲に構造化されたコンテキストを事前に構築するアプローチを避けることです。
「評価結果によると、目に見える変化は見られません」と彼女はこれらのアプローチについて語った。 「そして、私たちは一般的に、独自のツールを減らし、開発者が必要に応じて独自のツールを追加できるようにして、より無駄のないハーネスを出荷することに傾いていると思います。」
Augment Code は大幅に異なる賭けをしましたが、両社は必ずしも同じ介入をテストしたり、同じ結果に向けて最適化したりしているわけではありません。その製品は、埋め込み、検索モデル、ベクトル データベースを使用してリポジトリに事前にインデックスを作成し、概念的に関連するコードを取得します。そこで、この議論の裏側を知るために、Augment Code のエンジニアリング担当副社長である Vinay Perneti 氏に話を聞きました。
Perneti 氏は、Augment Code の異なるアプローチについて簡単に説明し、利点について自分のチームの評価を示し、より無駄のないハーネスを求める議論の一部に答え、AI ツールとエージェント ワークフローについて開発者がより広範に抱いている懸念のいくつかについて彼の見解を共有しました。
ヴィナイ・ペルネティとの会話
このインタビューは長さと明瞭さのために編集されています。
Ars Technica : Augment Code がどのように機能するのか説明していただけますか

ntextエンジンは動作しますか？
Vinay Perneti : それで、一歩下がって、特定の開発者が達成しようとしていることは何ですか?彼らは特定のタスクを実装したいと考えており、それをエージェントに渡します。そして、今日のエージェントの興味深い点は、コンテキスト ウィンドウが限られており、毎回基本的に必要なコンテキストをすべて取得して、それに取り組む必要があることです。
コンテキストには 2 つのアプローチがあります。 1 つは grep ベースです。クロード・コードとコーデックス、そして他のエージェントがそれを行いました。 2 番目はセマンティック検索ビットです。オーグメントに関しては、私たちは常にセマンティックな方法を採用してきました。その構成要素について説明します。
核となる部分は 2 つあると思います。 1 つは、システム内で動作する埋め込みモデルと検索モデルのペアがあり、さらにベクトル データベースと、サブミリ秒での検索を可能にする高度に最適化されたバックエンド システム全体があることです。
Ars : ある種類のコード ベースにおいて、別の種類のコード ベースと比べて利点はありますか?
Perneti : そうですね、実際、その利点は大規模なプライベート コードベースにあることがわかりました。そして、私がそう言う理由はここにあります。
ほとんどのベンチマークが実行されるすべての公開オープンソース リポジトリでは、基本的にすべてのモデルがリポジトリを記憶しています。これらのモデルは実際にリポジトリ全体を記憶できるほど大きいです。したがって、何かを達成しようとしているとき、モデルはどこを見るべきかをすでに知っているので、結果にすぐに到達できます。
一方、これをプライベート リポジトリで実行している場合、モデルはそのリポジトリを一度も見たことがありません。そして、結果を見つけるための反復ループははるかに長くなりましたよね?プライベート リポジトリ全体の意味を理解していれば、質問することができ、その結果をより迅速に得ることができます。
Ars : 人々はトークンの効率を懸念しています。これはありますか？

エマンティックアプローチはそれを助けるでしょうか？
ペルネティ : 確かに、特定の状況でそれを見たことがあります。実際、私たちはブログ投稿を公開しました。同じモデルのクロード コードとオーグメント コードを使用してターミナル ベンチを実行しました。また、同様の精度で完了しましたが、クロード コードよりも 33 パーセント効率的でした。したがって、探索側にそれほど多くの時間を費やしていないため、トークンをより有効に活用するという点でこれが現れていると私は見ています。
Ars : Anthropic 氏は、Claude Code には、より多くのセマンティック コード ナビゲーション ツールを含めることによる評価可能な評価の向上が見られないと私に言いました。しかし、その後、あなたのブログを見ると、実際にはこれが本当に役立つというベンチマークやその他の主張を投稿しているのがわかります。あなたがこれを言うとき、あなたたちはさまざまなことを測定していますか、それとも評価に何が欠けていますか？
ペルネティ : 素晴らしい質問ですね。彼らが具体的にどのような検索エンジンを使用したのかは知りません。そして、人々が混同しがちなもう 1 つの間違いは、すべてのデータベースが同等ではないのと同じように、すべての検索システムも同等ではないということだと思います。つまり、連携して動作するモデルとシステム、つまりコンテキスト エンジンが、結果の品質の点でも大きな違いを生むということです。
そのため、たとえば、Augment では、2022 年の設立当初 (ChatGPT 以前の時期) に約 18 か月を費やして、主に大規模なコード ベースのモデルの検索と埋め込みを研究しました。特定の結果を達成しようとする場合、この埋め込みスペースで取得する適切なコード部分は何なのかについて、多くの研究が行われてきました。これらすべてが検索モデルにエンコードされます。
そしてそれを非常に高速に実行するのが、私たちがそれを中心に構築したシステムです。したがって、誰かが「RAG 実装でクロード コードを試してみましたが、メリットがわかりません」と言ったら、それは問題です。

なぜなら、それが意味があるとしても、実装とコンテキストエンジンは非常に大きく異なるからです。
Ars : モデルは急速に改善されているため、何らかの仮定を置いて何かを構築することはまったく意味がないとの考えについてはどう思いますか?これは、超スリムなハーネスを推奨するか、あるいはこれを何も行わないかの議論であり、指数関数的に上昇するにつれて、6 か月後、12 か月後、またはその他の場合にどうなるかわからないというだけです。
ペルネティ : それには多くの真実がありますが、いくつかの異なる観点から考える必要があると思います。
Augment では、これについて、より質の高い結果を得るために必要な要素が 2 つある、インテリジェンスとコンテキストである、ということを考えてきました。モデルが本当に良くなると、インテリジェンスも指数関数的に向上することは間違いありません。しかし、彼らがより知的であるからといって、彼らがコンテキストを持っていることを意味するわけではありません。
これで、人はトークンを費やすことで、必要なコンテキストを取得できるようになります。そして、ここで 2 番目の次元が登場します。これは、エンジニアリングのリーダー全員が現在求めていること、つまりコストです。トークンの予算のうち、コンテキストの収集と適切な結果の生成にどれだけが当てられていますか?最高品質の結果を得るためにモデルを正しい方法で使用していますか?
そこで私は、ハーネスのデザインとコンテキストが重要だと考えています。私にとって、これはインテリジェンスとコンテキストの組み合わせであり、その時点ではシステム エンジニアリングの問題です。最小限のコストで最適な結果を得るために、これらの各分野に適切なエネルギーをどのように投入するかということです。
Ars : ソフトウェア開発における AI に懐疑的な読者はかなりの数います。特によくある反対意見が 2 つあります。その 1 つは、これらのシステムをまだ十分に信頼できないと感じていることです。

価値のない技術的負債が生じることになるでしょう。もう 1 つは、これらのエージェント ワークフローのコストが高すぎることです。どう思いますか？
ペルネティ : 実際には両方に触れます。両者について前置きをしておきます。誰もが同意してほしい基本的な真実は、モデルは指数関数的な速度で改良され続けているということだと思います。したがって、何をしていても、物事がどこに向かっていくのかを理解し、そこから利益を得る方向に向けて進むために、指数関数的に乗りたいと考えています。
それで、信頼と検証、技術的負債に関する質問に戻りますが、この問題に「エージェントにこれを引き渡して、私は立ち去る」というふうにアプローチするのは事実です。それはうまくいきません。だからこそ、私が指摘したのは、人間のチームがエージェントのチームと協力することであり、そこでは人間の方が判断に適している点がたくさんあるのですよね？スペックレビューは一例です。実際、エージェントは仕様書を書くのが苦手であることがわかりました。したがって、エージェントと実際に協力して、エージェントが適切で高品質の仕様書を作成できるように指導する必要があります。しかし、一度仕様を理解すれば、彼らは実行するのが本当に上手になります。したがって、それを活用しないのは意味がありません。
2 つ目は… ちなみに、テクノロジー負債は実際には非常に現実的です。 Augment の社内で気づいたパターンの 1 つは、私たちが明らかに非常にエージェント優先であることです。エージェントはコードを複製するのが非常に得意だということです。そのため、エージェントとの技術的負債を削減するために集中的なスプリントを実行する必要があることがわかりました。その利点は、技術的負債を削減することが何を意味するかを説明する仕様を考え出すことができ、それを実行するのが非常に上手であることです。ですから、それがその部分に取り組む方法だと思います。
後半ではコストについて話しました。あると思います

そこに興味深い点があります。これが今後の作業方法であると信じている場合、トークンにとって最良の結果を得るにはどうすればよいでしょうか?そして、そこにコンテキスト エンジンのようなものが違いを生むと私は考えています。私にとって、適切なタスクに適切なモデルを選択することは大きな違いをもたらします。
2 つ目は、これははるかに哲学的で将来を見据えたビットですが、これらのモデル機能が指数関数的に向上し続け、オープンソース モデルもそれに追いつくため、フロンティア ラボに送られるトークンの割合とオープンソース モデルが傾き始めると思われる時点が来るでしょう。最も困難な問題は、依然としてフロンティア モデルに投げ込まれることになりますが、私が説明したコーディング ステップは、何を行う必要があるかは正確にわかっているようです。仕様でそれを説明したことは知っていますが、オープンソース モデルでそれを実現できるかもしれません。その時点で、コストはかなり下がりますよね？したがって、このような働き方を可能にするシステムが必要であり、コストはある程度自然に解決すると思います。
アルス：ありがとうございます。お時間を割いていただきありがとうございます。
ペルネティ : そうですね、チャットはとても楽しかったです、サム。
Wu 氏と Perneti 氏は、フロンティア モデルが少なくとも来年は現在の指数関数的またはそれ以上の改善を続けると予想していることに同意しています。 （どちらも長い間明示的に対処されていませんでした）

[切り捨てられた]

## Original Extract

Augment Code's Vinay Perneti talks models, harnesses, and context.

Skip to content
Ars Technica home
Sections
Forum
Subscribe
Search
AI
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Pin to story
Theme
HyperLight
Search
Sign In
Sign in dialog...
Sign in
Agentic Development
Beyond grep: The case for a context-rich AI coding harness
Augment Code’s Vinay Perneti talks models, harnesses, and context.
93
Vinay Perneti, VP of Engineering at Augment Code, speaks during a presentation.
Credit:
Augment Code
Vinay Perneti, VP of Engineering at Augment Code, speaks during a presentation.
Credit:
Augment Code
Text
settings
Story text
Size
Small
Standard
Large
Width
*
Standard
Wide
Links
Standard
Orange
* Subscribers only
Learn more
Minimize to nav
There are a lot of AI coding applications out there, and as impressive as large language models and the agents they enable have become, many of the most recent developments in AI-assisted development have been in the software that manages those models, not just the models themselves.
Earlier this summer, I spoke with the head of product for Claude Code, Anthropic’s Cat Wu, about that company’s approach to building that software.
Wu repeatedly came back to the same point in our conversation: Anthropic’s models (and those of its direct competitors) are improving so quickly that it makes little sense to plan too far ahead or to build opinionated or limiting features around them. Rather, the Claude Code product team attempts to maintain what they call a lean harness.
A harness is the software built around one or more AI models that determines how they are used. It decides what the models see, what actions they can take, and how they interact with the code base. Think of it like a layer between a model and the developer’s actual project.
It’s not that Claude Code as an application and harness has no opinionated features or design choices, but its product team does seem to err on the side of trusting where the models will take them over the next year.
There are other harnesses besides Claude Code, though. There are OpenAI’s Codex, Google’s Antigravity, open source alternatives like OpenCode, and options from startups like Cursor or Augment Code. Each can have different features, emphasis, or opinions about how a model can or should be utilized in agentic workflows.
One key choice that Claude Code makes is to avoid approaches that build a structured context around a codebase in advance by default.
“Going by the evals, we don’t see a measurable change,” she said of those approaches. “And I think we generally lean more toward shipping a leaner harness with fewer opinionated tools and just letting developers add their own if they want.”
Augment Code has made a substantially different bet, although the two companies are not necessarily testing the same interventions or optimizing for the same outcomes; its product pre-indexes the repository using embeddings, a retrieval model, and a vector database, then retrieves conceptually relevant code. So to get the other side of that discussion, I spoke with Vinay Perneti, Augment Code’s VP of Engineering.
Perneti briefly described Augment Code’s differing approach, offered his own team’s evaluations of the advantages, responded to some of the arguments for a leaner harness, and shared his perspective on some of the concerns developers have about AI tools and agentic workflows more broadly.
A conversation with Vinay Perneti
This interview has been edited for length and clarity.
Ars Technica : Can you explain how Augment Code’s context engine works?
Vinay Perneti : So, if I take a step back, what is it that a particular developer is trying to achieve? They want a particular task implemented and they give it to an agent. And the interesting thing with agents today is they have a limited context window, and every time they need to essentially go get all of the context needed for that and then work on it.
There’s two approaches to context. One is grep-based. Claude Code and Codex and other agents have done that. The second is the semantic retrieval bit. For Augment, we’ve always taken the semantic route, and I’ll describe the building blocks of that.
I would say there’s two core pieces. One is, we have an embedding and a retrieval model pair that are working in the system, and then you have a vector database and an entire, highly optimized back-end system that makes it possible to retrieve in sub-milliseconds.
Ars : Does it have an advantage in one kind of code base versus another?
Perneti : Yeah, it actually turns out the advantage comes in large, private codebases. And here is why I say that.
For all the public, open source repos where most of the benchmarks are run, every single model has basically memorized the repo. These models are large enough that they’re actually able to memorize the whole repo. So when you’re trying to get something done, the models kind of already know where to look so they can get to an outcome quickly.
Whereas when you’re doing this in a private repo, a model has never seen that repo. And now the iteration loop for finding the outcome is much longer, right? If you have a semantic understanding of your entire private repo, you can ask a question, and you get to those outcomes much more quickly.
Ars : People are concerned about token efficiency. Does this semantic approach help with that?
Perneti : We’ve certainly seen it in certain situations. In fact, we published a blog post—we ran Terminal-Bench with Claude Code and Augment Code, same model. And we completed at similar accuracy, but we were 33 percent more efficient than Claude Code. So I do see this showing up in terms of making better use of tokens because you’re not spending as much time on the exploration side.
Ars : Anthropic told me that Claude Code does not see measurable eval gains from including more semantic code-navigation tools. But then I look at your blog, and I see you posting benchmarks and other claims that, actually, this really helps. Are you guys measuring different things when you say this, or what are their evals missing?
Perneti : Great questions. I don’t know what specific retrieval engine they used, and I think the other mistake that often people conflate is, not all retrieval systems are equal just like not all databases are equal, right? And by that, what I mean is the models and the system that are working together, which is the context engine, makes a big difference in terms of the quality of outcomes as well.
So, for example, at Augment, we spent about 18 months at the beginning of the company being founded in 2022—this is pre-ChatGPT—researching retrieval and embedding models predominantly for large code bases. So there’s a lot of research that has gone into, when you’re trying to achieve a particular outcome, what are the right pieces of code to get in this embedding space? All of that is encoded into our retrieval models.
And doing that in a very, very fast way is the system that we built around it. So when somebody says, ‘Hey, I tried Claude Code with a RAG implementation, but I’m not seeing benefits,’ that’s because the implementation and the context engine are very, very different, if that makes sense.
Ars : What do you say to the notion that models are improving so fast that building anything with any assumptions at all makes no sense? That’s the argument for a super lean harness or not doing any of this, is just that you don’t know where it’s gonna be in six months or twelve months or whatever as the exponential goes up.
Perneti : There’s a lot of truth to that, but I think you have to think about it from a couple of different perspectives.
At Augment, the way we’ve been thinking about this is there’s two ingredients that you need for higher quality outcomes, intelligence and context. When models are getting really good, intelligence is going to get exponentially better, no doubt about that. But just because they’re more intelligent does not mean they have the context.
Now, a person can get the context that they want by spending the tokens on it. And that’s where the second dimension comes, which is what all of the engineering leaders are asking right now, is cost. How much of your token budget is going towards context gathering and producing the right outcomes? Are you using the model the right way to get the highest quality outcomes?
And that’s where I think harness design and context matter. So to me, it’s a combo of intelligence and context, and then it’s a systems engineering problem at that point: how do you put the right energy into each of these verticals so that you get the most optimal outcome with the least cost?
Ars : We have a fair number of readers who are skeptical of AI in software development. I see two objections that are especially common. One of them is that they still feel that they just can’t trust these systems enough, that they’re going to create technical debt that’s not worth it. The other is that these agentic workflows are too expensive. What do you think?
Perneti : I’ll actually touch on both. Let me preface for both of them. I think the fundamental truth I hope everybody agrees with is models are continuing to improve at an exponential rate. So whatever you’re doing, you want to ride the exponential so that you’re understanding where things are going and you’re orienting yourself to benefit from that.
So, going back to your question around trust and verification and tech debt, it is true if you approach this problem as, ‘I’m handing this off to agents and I’m walking away.’ That’s not how it’s gonna work, and that’s why I point out that it’s teams of humans working with teams of agents, where there are many points where humans are still better suited for judgment, right? Spec reviews is an example. We find that agents are actually not good at writing specs. So you have to really work with an agent to guide the agent to write a good, high-quality spec. But once you have the spec we are at a place, they’re really good at executing. So it makes no sense to not take advantage of that.
The second thing… tech debt is actually very real, by the way. One of the patterns that we noticed internally at Augment—we are obviously very agent-forward—is agents are very good at duplicating code… so we found that we’ve had to do focused sprints on reducing tech debt with agents. The beauty about that is that you can say you can come up with a spec that talks about what it means to reduce tech debt, and they’re really good at executing that. So I think that is the way to tackle that part of it.
And the second part, you talked about cost. I think that there’s an interesting point there. If you believe that this is the way of working moving forward, then how do you get the best outcome for your tokens? And that’s where I think things like context engine make a difference. To me, picking the right model for the right task makes a huge difference.
The second thing—this is much more a philosophical, forward-looking bit—is, as these model capabilities continue to improve in exponential and open source models are going to keep up, there’s going to be a point where I think the proportion of tokens that’s going to frontier labs, to open-source models, is going to start tipping, where your most difficult problem, you’ll still throw at a frontier model, but the coding step that I described—like I know exactly what needs to get done and I know I described that in a spec—an open source model might be able to get that done for you. At that point, your cost drops quite a lot, right? So you need a system that allows you to work this way, and I think the cost will just kind of take care of itself, is how I see it.
Ars : Thank you. I appreciate you making the time.
Perneti : Yeah, I really enjoyed the chat, Sam.
Both Wu and Perneti agree that they expect frontier models to continue to improve at or even above the current exponential, at least for the next year. (Neither explicitly addressed a longer time

[truncated]

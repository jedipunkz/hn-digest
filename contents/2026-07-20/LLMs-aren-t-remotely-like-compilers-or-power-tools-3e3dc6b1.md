---
source: "https://blainehansen.me/post/llms-not-power-tools/"
hn_url: "https://news.ycombinator.com/item?id=48973541"
title: "LLMs aren't remotely like compilers or power tools"
article_title: "LLMs aren't remotely like compilers or power tools | Blaine Hansen"
author: "blainehansen"
captured_at: "2026-07-20T02:00:51Z"
capture_tool: "hn-digest"
hn_id: 48973541
score: 1
comments: 0
posted_at: "2026-07-20T01:47:59Z"
tags:
  - hacker-news
  - translated
---

# LLMs aren't remotely like compilers or power tools

- HN: [48973541](https://news.ycombinator.com/item?id=48973541)
- Source: [blainehansen.me](https://blainehansen.me/post/llms-not-power-tools/)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T01:47:59Z

## Translation

タイトル: LLM はコンパイラやパワー ツールとは異なります
記事のタイトル: LLM はコンパイラやパワー ツールとはまったく異なります |ブレイン・ハンセン
説明: LLM を個別のインターフェイスを備えた決定論的システムに関連付ける比喩は誤解を招きます。

記事本文:
github LLM は、リモートではコンパイラやパワーツールとは異なります
LLM を個別のインターフェイスを備えた決定論的システムに関連付ける比喩は誤解を招きます。
LLM についての議論では、次の 2 つのよく言われる言葉を耳にします。
それらは「マシンコードからアセンブリのような、抽象化の次のレベルにすぎない」ということです。
これらは「テーブルソーやバックホーを使用するような、次のレベルの自動化にすぎない」ということです。
これらのステートメントはどちらも愚かなレベルで間違っており、その結果、思考を終了させる常套句として機能しています (新しいウィンドウが開きます)。
誤解のないように言っておきますが、私は元本に関して LLM に反対しているわけではありません。LLM が役立つ方法はあると思います。しかし、彼らをめぐる現在の対話は疲れ果てて愚かなものであり、これら 2 つの特定の間違いを犯す人々にはもう我慢できません。
§ LLM を使用すると、委任することになります。
LLM がコンパイラやパワー ツールとは異なる理由を最もよく理解するために、なぜ委任が常に優れた比喩であるのか、また、なぜ他の比喩が非常に誤解を招くのかを明確にしたいと思います。
人にどれだけの仕事を任せても、自分が何を得ることができるのかを正確に知ることはできません。人間は (少なくともあなたの観点からは) 確率的な実体であり、他の確率的な実体と同様に、意味のある方法で正確に制御することはできません。本人と代理人の問題 (新しいウィンドウが開きます) は、社会で最も困難かつ頑固な問題の 1 つです。
人間には、人間との付き合い方について深く進化した直感がたくさんありますが、その大きな直感の 1 つは、自分が失望するかもしれないということをよく知っているということです。 (「何かを正しくやりたいなら、自分でやれ。」) 人間には一般的な知性があるにもかかわらず、仕事を他人に委任し、実際に何か役に立つことを得るのは非常に難しいことです。
対照的に、LLM は非常に奇妙です。なぜなら、LLM は人間とは根本的に異なるからです。

しかし、それらの本質的なデザインは、人間が行うことしか経験したことのないことを模倣することです。
私がユーモラスだと思った比喩 (ただし、依存症者とその家族に対して失礼だと思います...ごめんなさい) は、LLM はインターネット全体を読んでいる覚せい剤マニアのようなものです。
彼らは熱心で、疲れも不満も言わずに働きます。
彼らは多くの場合、受け入れられる仕事をしますが、多くのやり取りがなければそうではありません。
自分のデバイスに任せると、彼らは奇抜なコードを作成します (新しいウィンドウが開きます)。
彼らは、単語や文字を数えることができないなど、奇妙で混乱を招く間違いを犯します (新しいウィンドウが開きます) (新しいウィンドウが開きます)。
また、場合によっては、本番データベース全体を削除したり (新しいウィンドウが開きます) 、すべてのユーザー データや機密情報がインターネットに公開されたり (新しいウィンドウが開きます) こともあります。
はっきり言っておきますが、何らかの知識タスクを 100 人の覚せい剤マニアに委任することが間違いなく有益であると思われる状況があります (もちろん倫理的ではありませんが...) (これらのポッドキャストを聞いていると、タイプ セーフティについて話していることがよくわかります。オープンソース ソフトウェアに関する判決を見つけるために米国最高裁判所の判決をすべて読んでください。3 番目の面白い例です)。しかし、個人的には、そのような状況はまれであり、狭いものだと思います。
プロジェクトで LLM を使用するのは、請負業者にアウトソーシングするようなものだと私も一時言いたくなりましたが、請負業者がかなり悪質であるか、インセンティブが一致していない場合、またはグループ間に深刻な言語の壁がある場合にのみ、比較に意味があると思います。
しかし、基本的にはどんな比喩も不適切であり、LLM は奇妙すぎるのです。高次元の意味論的空間 (opens new window) への埋め込みに基づいて言語トークンを自己回帰的に生成するトランスフォーマーベースのディープ ニューラル ネットワークを、これまでに存在したものと比較するのは意味がありません。
§ コンパイラとパワーツールは決定的です
両方ともc

コンパイラーとパワーツールは、完全に予測できるように意図的に設計されています。このようなツールの重要な点は、可能な限り狭いインターフェイスの背後でかなりの量の自動パワーを提供することです。これらは、機械の約束を完全に実現するように設計された機械です。機械には意志や癖がなく、感謝の言葉を言う必要がなく（またはいじめる必要もありません（新しいウィンドウが開きます））、常にまったく同じことを行います。
電動工具は、可能な限り少ない数の制御装置の背後に大量の機械的複雑性をパッケージ化しています。電源ボタン、コントロールハンドル、ステアリングホイール。その目的は、人が反復的な作業をより速く完了し、より高いレベルの品質で作業でき、できればその作業をより安全に行えるようにすることです。
コンパイラーとプログラミング言語は本質的に、正確な論理構造を決定論的に再利用することを目的としています。彼らが下したすべての決定により、あなたはそれらの決定を再度行う必要がなくなります。再利用可能なソフトウェアを作成することは、機械部品 (滑車や歯車など) を設計するようなものですが、誰もがそれを無限にコピーし、設計したとおりに使用することができます。
LLM のインターフェイスは自然言語であり、自然言語はシステムに対する可能な限り最も信頼性の低い、曖昧な、不正確な、予測不可能なインターフェイスです。
自然言語は最後の手段のインターフェイスです。それは完全に一般的ですが、ひどく曖昧でもあります。人間は、欲しいものを表現するより正確な方法が他にない場合、自然言語が存在しないか、私たちがそれを快適に理解していないため、自然言語に戻ります。私たちが図や設計図を描き、数式を書くのには理由があります。
バックホーにはボタン、ジョイスティック、レバーなどが付いていますが、これはそれらのインターフェースによって機械の操作が正確になるためです。コンピューター言語とそれを使用して構築される GUI

つまり、理解する努力をすれば、その意味を正確に知ることができるのです。自然言語は確かに多くのシステムへのスムーズな導入 (新しいウィンドウが開きます) になる可能性がありますが、個別のインターフェイスをよく理解すると、自然言語ですべてを話すよりもそれを使用する方が常に速くなります。
ツールを使用する最大のポイントは、反復的な問題を予測どおりに解決することです。対照的に、LLM は、プリンシパルとエージェントの問題のより奇妙で混乱を招くバージョンを再現します。
それを嘆くかどうかはあなた次第ですが、実際には委任していないと私に納得させようとしないでください。
今後私からの連絡を希望しますか?

## Original Extract

Any metaphor for LLMs that relates them to deterministic systems with discrete interfaces is misleading you.

github LLMs aren't remotely like compilers or power tools
Any metaphor for LLMs that relates them to deterministic systems with discrete interfaces is misleading you.
I hear two common refrains in discussions of LLMs:
That they're "just the next level of abstraction, like machine code to assembly".
That they're "just the next level of automation, like using a table saw or a backhoe".
Both of these statements are wrong to the level of being moronic, and as a result are functioning as thought-terminating cliches (opens new window) .
To be clear, I'm not against LLMs on principal, and I think there are ways they can be useful. But the current dialog around them is exhausting and stupid, and I can't stand people making these two particular errors anymore.
§ To use an LLM is to delegate
To best understand why LLMs aren't like compilers or power tools, I want to make clear why delegation is always a better metaphor, and therefore why the other metaphors are so misleading.
When you delegate any amount of work to a person, you can't be sure exactly what you're going to get. A person is a stochastic entity (at least from your perspective), and like any stochastic entity, it can't be precisely controlled in any meaningful way. The Principal-Agent Problem (opens new window) is one of the most difficult and stubborn problems in society.
Humans have many bone-deep evolved intuitions for how to deal with humans, and one of the big ones is to just know you very well might be disappointed! ("If you want something done right, do it yourself.") Delegating work to others and actually getting something useful is extremely hard, despite the fact that humans have general intelligence.
LLMs are profoundly strange by contrast, because they're fundamentally dissimilar to humans, and yet their intrinsic design is to mimic a thing we've only ever experienced humans doing.
A metaphor I've found humorous (but which I'm guessing is disrespectful to addicts and their families... sorry) is that an LLM is like a meth-head who's read the entire internet:
They're enthusiastic and work without tire or complaint.
They often do work that is acceptable, but not without lots of back-and-forth.
If left to their own devices they write wacky code (opens new window) .
They make bizarre and confounding mistakes (opens new window) , like being unable to count words or letters (opens new window) .
And sometimes they delete the entire production database (opens new window) , or expose all your user data or secrets to the internet (opens new window) .
Let me be clear: there are situations where I would definitely find it useful (although certainly not ethical...) to delegate some knowledge task to 100 meth-heads (listen to these podcasts and find all the times they talk about type safety; read all US supreme court decisions to find the ones about open source software; third funny example). But I personally think those situations are rare and narrow.
For a time I was tempted to say using an LLM on a project is like outsourcing it to a contractor, but I only think the comparison has any juice if the contractor is pretty bad or has misaligned incentives, or there's a serious language barrier between the groups.
But fundamentally any metaphor will be inadequate, LLMs are too weird. It doesn't make sense to compare a transformer-based deep neural network that autoregressively generates language tokens based on their embedding in a high-dimensional semantic space (opens new window) to really anything that's ever existed.
§ Compilers and power tools are deterministic
Both compilers and power tools are intentionally designed to be completely predictable. The entire point of such tools is that they present a fair amount of automatic power behind the narrowest interface possible. They are machines designed to fully realize the promise of machines, in that they have no will or quirks, you don't have to say thank you ( or bully them (opens new window) ), and they do exactly the same thing all the time.
Power tools package up a large amount of mechanical complexity behind the smallest number of controls possible, e.g. a power button, a control handle, a steering wheel. The intention is that a person can complete repetitive work much faster, work at a much higher level of quality, and hopefully do that work much more safely.
Compilers and programming languages are intrinsically about deterministically reusing precise logical structures. All the decisions they made allow you to not have to make those decisions again. Writing reusable software is like you designed some machine part (like a pulley or a gear), but everyone can infinitely copy it and use it precisely the way you designed it.
The interface of an LLM is natural language, and natural language is the most unreliable/fuzzy/imprecise/unpredictable interface to a system possible.
Natural language is the interface of last resort. It's fully general, but it's also furiously ambiguous. Humans fall back to natural language when no other more precise way to describe what we want is available to us, either because it doesn't exist or we don't know it comfortably. We draw diagrams and blueprints and write math equations for a reason!
Backhoes have buttons and joysticks and levers etc because those interfaces make the operation of the machine precise. Computer languages and the GUIs we build with them are constructions of pure logic, meaning we can know exactly what they mean if we do the work to understand them. Natural language could certainly be a smooth onramp (opens new window) to many systems, but once someone knows a discrete interface well, it will always be faster for them to use it than say everything in natural language.
The entire point of using tools is to predictably solve a repetitive problem. LLMs by contrast reproduce a weirder and more confusing version of the principal-agent problem.
Whether you're down for that is up to you, just don't try to convince me you aren't actually delegating.
Want to hear from me in the future?

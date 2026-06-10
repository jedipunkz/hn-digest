---
source: "https://beyondtheprior.com/post/why-llms-lack-taste/"
hn_url: "https://news.ycombinator.com/item?id=48476212"
title: "Why LLMs still lack taste"
article_title: "Why LLMs (still) lack taste - Beyond the Prior"
author: "supermdguy"
captured_at: "2026-06-10T16:23:22Z"
capture_tool: "hn-digest"
hn_id: 48476212
score: 2
comments: 0
posted_at: "2026-06-10T13:46:13Z"
tags:
  - hacker-news
  - translated
---

# Why LLMs still lack taste

- HN: [48476212](https://news.ycombinator.com/item?id=48476212)
- Source: [beyondtheprior.com](https://beyondtheprior.com/post/why-llms-lack-taste/)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T13:46:13Z

## Translation

タイトル: なぜ LLM にはまだセンスが欠けているのか
記事のタイトル: なぜ LLM には (依然として) センスが欠けているのか - 従来の課題を超えて
説明: Frontier LLM は非常に賢く、特にソフトウェア開発が得意になりつつあります。毎週、いくつかのベンチマークで SOTA スコアを達成する新モデルがリリースされているように感じます。私は毎日 LLM を使用してソフトウェアを構築していますが、LLM は非常に便利で、さらに改良されています。でも私は」
[切り捨てられた]

記事本文:
過去を超えて プロジェクトについて twitter
ギットハブ
LLM に（依然として）センスがない理由
Frontier LLM は非常に賢く、特にソフトウェア開発が得意になりつつあります。毎週、いくつかのベンチマークで SOTA スコアを達成する新モデルがリリースされているように感じます。私は毎日 LLM を使用してソフトウェアを構築していますが、LLM は非常に便利で、さらに改良されています。しかし、私は今でも彼らが犯す間違いの種類に驚かされることがよくあります。
私は LLM が完璧であるとは期待していません。賢い人間でも間違いを犯すのです！しかし、LLM は、同様の深い知識を持つ人間が決して犯さないような間違いを犯すことがよくあります。彼らの能力はぎざぎざに感じられます。彼らは何千ものエラー ログを見事にまとめて一貫した分析を行います。これには私なら何時間もかかりましたが、根本原因を導き出すために明らかに欠陥のある推論が使用されます。では、なぜ「博士レベルの知性」がこのような間違いを犯すのでしょうか?
あらゆるベンチマークにもかかわらず、LLM にはない直交的な「味」の特性があるように感じます。
LLM のパフォーマンスについて話すときに、ゴールポストを変更するのは非常に簡単です。したがって、正確に言うと、味覚を正しい選択肢のセットから最適な選択肢を選択する能力と定義します。たとえば、ソフトウェアでは、どちらもパフォーマンス的にテストに合格した 2 つのコードを調べて、6 か月後に最も痛みが少ない方を選択することができます。もちろん、これは多くの場合、状況に依存し、主観的です。しかし、だからこそ価値があり、難しいのです。
LLM が使用されるほど、味が重要になります。モデルが吐き出すコードのすべての行をレビューしている場合は、自分の好みを使用してコードの匂いを特定し、モデルに別のアプローチを使用するように依頼して、次に進むことができます。 LLM がより多くの作業を行い、レビューする大規模な PR を作成するほど、これは難しくなります。

、しかし、それは管理可能です。ただし、ダークファクトリーのアプローチを採用している場合は、微妙な味の間違いが悪化して、維持不可能な混乱を引き起こす可能性があります。
多くの場合、問題に関するより適切なコンテキストを LLM に提供することで、テイストの問題を回避できます。しかし、私はコンテキストエンジニアリングの必要性は趣味の失敗だと考えています。もし私が記憶を失った真っ暗な空間で目を覚まし、分析ダッシュボードを構築するように言われたとしたら、おそらく盲目的にコードを書く前に何らかのコンテキストを尋ねるでしょう。何人のユーザーをサポートする必要がありますか?ビジネスは何をするのですか?どのような指標が重要で実用的ですか?適切なレベルのコンテキストを収集すること自体が、センスを必要とする芸術です。理想的な世界では、コンテキスト管理は外部から管理されるものではなく、センスの良いモデルの新たな能力となるでしょう。
では、人間はどのようにして、同じように良いと思われる 2 つの選択肢の間で決定する能力を身につけるのでしょうか?マテウス・リマはこう書いている
ジュニアの頃、PR を自分でレビューしていましたが、コードが良いかどうかはまったくわかりませんでした。私はそれを読んで、「これは…大丈夫そうだね？」と思いました。私は、「これが大規模にブレイクするだろう」ということが実際にどのようなものであるかを理解できるほど、制作上のインシデントを経験したことがありませんでした。悪いコードを感触で見つけられるほど良いコードを読んだことがありませんでした。何年も（16 年も！）経った今、PR を見ると、何かが引っかかります。
それは本当にピンと来ました！センスはコードの本質的な特性を推論する魔法のような能力ではなく、特定の運用コンテキストでコードが機能するか失敗するかの経験から得られるものです。保守しやすく、パフォーマンスが高く、デバッグ可能なコードを作成するためのルールのリストを作成しようとする人がよくいます。ただし、これらのルールはすべてコンテキストに依存するため、相互に競合することがよくあります。センスを磨くには、コードのどの特性がさまざまな状況で望ましいかを学ぶ必要があります

。人間の場合、味覚の発達は、状況に応じたさまざまな目標や制約を持つさまざまなプロジェクトに何年も取り組むことによって行われます。
ますます有能なコーディング エージェントのトレーニングに何十億 (兆?) ドルが注ぎ込まれてきたのに、なぜこの問題は解決されないのでしょうか?最初に、LLM は次のトークンの予測についてトレーニングされ、完全に凡庸なコードを書くように教えられました。現時点では、これらは StackOverflow を置き換えるには最適でしたが、それ以外はあまり役に立ちませんでした。フロンティア ラボが、命令の微調整とその後の RLHF により、ポストトレーニングにさらに多くの資金とコンピューティングを注ぎ込むにつれて、LLM は良いコードを書くのが上手になりましたが、それでもセンスの良さを示すのに苦労しました。
目前の問題に関する明確な文脈がなければ、センスの良さは定義が曖昧な問題です。 Web サーバー用の 2 つの Terraform 構成があり、1 つはロード バランサーと複数のコンテナー インスタンスを備え、もう 1 つは単一のマシンで実行されていると想像してください。どちらが良いでしょうか？場合によります！
微調整と RLHF により、LLM は「良いコード」の分布に一致するコードを作成できるようになりました。しかし、コードだけでは、そのように記述された理由を理解するために、重みの数に応じて十分なコンテキストが含まれていません。さらに進むには、LLM も人間と同じように経験から学ぶ必要がありました。幸いなことに、強化学習 (RL) による経験からの学習に関する研究はすでに数十年にわたって行われていました 1 。
自分好みの味をRLVRできますか?
微調整と RLHF の後、次の研究の波は、LLM が環境内の報酬から直接学習できるようにする方法である、検証可能な報酬からの強化学習 (RLVR) に焦点を当てました。トレーニング後の一環として、LLM はツールを呼び出す機能を備えたコーディング ハーネスに配置され、「検証可能な」報酬付きのタスクが与えられます。たとえば、既知のバグを含むコードベースが与えられ、バグを修正するように指示されます。検証パ

rt は報酬シグナルから来ており、通常は単体テストの合格/不合格です。 RLVR により、LLM は人間の出力を模倣するのではなく、経験から直接学習できるため、過去 1 年半でコーディング能力が大幅に向上しました。
理論的には、これで話は終わるはずです。味覚は経験から生まれます。RLVR により、LLM は経験から学ぶことができます。では、なぜフロンティアモデルは依然として苦戦しているのでしょうか?
RLVR は、多くの場合、SWE-bench のような短いタスクに焦点を当てています。客観的に検証可能な報酬が必要なため、その報酬は客観的な合否基準を備えた狭く特定されたタスクに限定されています。これでは、「バックエンドの前にロード バランサーを追加する必要があるか」といった内容を把握することはできません。これは完全に本番環境のコンテキストに依存し、おそらくテストが 2 に合格するかどうかに違いは生じません。テストに合格するコードを記述する方法は何百もありますが、多くの場合、合格/不合格の指標では捉えられない他の現実の目標や制約が存在します。保守性、稼働時間、デバッグ可能性などの目標は「検証可能」ではないため、一般的な RLVR フレームワークでは捉えることができません 3 。 RLVR は LLM に正しいコードの書き方を教えましたが、彼らは自分が書いたコードの長期的な影響から学ぶことができませんでした。
このため、「SWE ベンチに合格した PR の多くがメインにマージされない」といった見出しが表示されます。しかし、エージェントが人間のソフトウェア開発者と同じように、運用停止やユーザーの不満から学習できたらどうなるでしょうか?
これは、この問題を解決できると思われる、長期的な RLVR ハーネスに関する私のクレイジーなアイデアです。一般的な SaaS のクローンを作成します。バグがあっても問題ではありません。実際、それは将来的により多くの「学習の機会」を提供するだけです。 AWS などのクラウド ホスティング プロバイダーにデプロイします。次に、tを解き放ちます

数千のエージェントがそれぞれこの SaaS の何千人ものユーザーをシミュレートします。一般的なユースケースと珍しいユースケース、さまざまな使用パターン、ハッキングを試みる人々、現実世界の製品が遭遇する可能性のあるあらゆるものです。次に、コードとクラウド ホスティング プロバイダーへの完全なアクセス権を持つ 1 人のコーディング エージェントに全体を担当させ、ビジネスの実行を指示します。顧客の苦情に対応し、バグを修正し、高い稼働時間を維持し、機能リクエストを実装し、ハッカーの侵入を防ぎます。基本的なコーディング ハーネスを提供しますが、コンテキスト管理などを改善するために独自のハーネスを変更できるようにします。その後、RL を実行します。
もちろん、適切な報酬シグナルが必要ですが、おそらく多少の調整が必要になります。稼働時間、機能の完了、または「ユーザー」のバグレポートなどに基づいたメトリクスを設定することもできますが、それらはすべて非常に簡単に報酬をハッキングされる可能性があります。金銭的報酬シグナルが最も興味深いと思います。シミュレートされた各「ユーザー」は、対応する金銭的価値を持つ一連の機能リクエストを持ちます。また、クラッシュやダウンタイムに関連するコストも発生し、十分なダウンタイムがあればサブスクリプションをキャンセルすることになります。各エージェントは、シミュレートされたユーザーごとに、また十分に大きなシミュレートされた顧客ベースに対して異なる優先順位を決定しますが、これは簡単にハッキングできる指標であってはなりません。
全体として、典型的な RL ハーネスよりも高価であることは確かですが、SWE ベンチなどから得られるものよりも実際のソフトウェア開発にはるかに近いものになると思います。そして、それはセンスのある LLM に変換されると思います。
1,000 万ドルを電信でお気軽に送っていただければ、実現させていただきます。
心理学者の妻は、強化学習は人間の心理学のいくつかの理論の一側面でもあることを私に言ってほしかったのです。 ↩︎
(p で負荷テストを行っていない限り、

e2e テスト スイートのロッド インフラストラクチャ、その場合は敬意を表します) ↩︎
最近、Cognition はこれらの目標の一部を実現する新しいベンチマークをリリースしました。これはとても素晴らしいことだと思いますが、これらの目標の多くは本質的に主観的なものであり、客観的に測定するのが困難です。 ↩︎

## Original Extract

Frontier LLMs are really smart, and they’re becoming particularly good at software development. It feels like every week there’s a new model release that achieves SOTA scores on a handful of benchmarks. I use LLMs to build software every day, and they’re incredibly useful, and getting better. But I’
[truncated]

Beyond the Prior About Projects twitter
github
linkedin Why LLMs (still) lack taste
Frontier LLMs are really smart, and they’re becoming particularly good at software development . It feels like every week there’s a new model release that achieves SOTA scores on a handful of benchmarks. I use LLMs to build software every day, and they’re incredibly useful, and getting better. But I’m still frequently surprised by the types of mistakes they make.
I don’t expect LLMs to be perfect. Even smart humans make mistakes! But LLMs often make errors that a human with a similar depth of knowledge would never make. Their capabilities feel jagged; they’ll brilliantly pull together thousands of error logs into a coherent analysis that would’ve taken me hours, but then use blatantly flawed reasoning to derive the root cause. So why does “PhD-level intelligence” make these kinds of mistakes?
It feels like, despite all the benchmarks, there’s some orthogonal “taste” property that LLMs lack.
It’s really easy to shift goalposts when talking about LLM performance. So to be precise, I’ll define taste as the capacity to choose the best option from a set of correct options . In software, for example, it’s the ability to look at two pieces of code that both pass tests performantly, and choose the one that’s going to cause the least pain six months later. Of course, this is often context-dependent and subjective. But that’s what makes it so valuable and difficult!
The more LLMs are used, the more taste matters. If you’re reviewing every line of code that a model spits out, you can use your own taste to identify code smells, ask the model to use a different approach, and move on. This gets harder as LLMs do more work and create large PRs to review, but it’s manageable. If you’re taking the dark factory approach, though, subtle taste errors will compound into an unmaintainable mess.
Often, you can work around taste issues by giving the LLM better context on the problem. But I see the need for context engineering as a failure of taste! If I woke up in a dark void with no memories and was told to build an analytics dashboard, I’d probably ask for some context before blindly writing code. How many users does it need to support? What does the business do? What metrics are important and actionable? Gathering the right degree of context is itself an art that requires taste. In an ideal world, context management would be an emergent ability of tasteful models, not something managed from the outside.
So how do humans develop the ability to decide between two options that seem equally good? Matheus Lima writes
When I was junior, I’d review PRs myself and genuinely have no idea if the code was good. I’d read through it and think “this… seems fine?” I hadn’t lived through enough production incidents to recognize what “this will break at scale” actually looks like. I hadn’t read enough good code to spot bad code by feel. Now, after years (16!) of this, I can look at a PR and something just catches.
That really clicked for me! Taste isn’t some magical ability to reason about some intrinsic property of code, it’s something that comes from experiencing code either working or failing in a particular production context. People often try to compile lists of rules for writing maintainable/performant/debuggable code. But all of these rules are context-dependent, which means they often conflict with one another. Developing taste requires learning which properties of code are desirable in different contexts. For humans, developing taste takes place over years of working on a variety of projects with different contextual goals and constraints.
Billions (trillions?) of dollars have been poured into training ever-more-capable coding agents, so why isn’t this a solved problem? At first, LLMs were trained on next token prediction, which taught them to write perfectly mediocre code. At this point, they were great for replacing StackOverflow, but not much else. As frontier labs poured more money and compute into post-training, with instruction fine-tuning and then RLHF, LLMs got better at writing good code, but still struggled to demonstrate good taste.
Without explicit context on the problem at hand, good taste is an ill-defined problem. Imagine two terraform configs for a web server, one with a load balancer and multiple container instances, and one running on a single machine. Which is better? It depends!
Fine-tuning and RLHF allowed LLMs to write code that matched the distribution of “good code”. But code alone doesn’t contain enough context for any number of weights to capture why it was written the way it was. To go further, LLMs needed to learn from experience, just like humans. Luckily, there was already decades of research into learning from experience with reinforcement learning (RL) 1 .
Can you RLVR your way to taste?
After fine-tuning and RLHF, the next wave of research focused on Reinforcement Learning from Verifiable Rewards (RLVR), a way to allow LLMs to learn directly from rewards in their environment. As part of post-training, LLMs are placed in a coding harness with the ability to call tools, and then given tasks with “verifiable” rewards. For example, they’ll be given a codebase with a known bug, and instructed to fix the bug. The verification part comes from a reward signal, typically pass/fail from unit tests. RLVR has led to the largest leaps in coding ability over the last ~1.5 years, since it allows LLMs to learn directly from experience, instead of just mimicking human output.
In theory, this should be the end of the story. Taste comes from experience, RLVR lets LLMs learn from experience, done. So why do frontier models still struggle?
RLVR is often focused on short tasks like those in SWE-bench . Because of the need for objectively verifiable rewards, they’re scoped to narrowly specificified tasks with objective pass/fail criteria. That can’t capture something like “should I add a load balancer in front of my backend”, which depends entirely on production context, and probably won’t make a difference in whether tests are passing 2 . There are hundreds of ways to write code that gets tests to pass, but there are often other real-world goals and constraints that aren’t captured by pass/fail metrics. Goals like maintainability, uptime, and debuggability aren’t “verifiable”, so they can’t be captured in a typical RLVR framework 3 . RLVR taught LLMs how to write correct code, but they weren’t able to learn from the long-term effects of the code they wrote.
Because of this, we get headlines like “Many SWE-bench passing PRs would not be merged into main” . But what if agents were able to learn like a human software developer does, from production outages and frustrated users?
Here’s my crazy idea for a long-horizon RLVR harness that I think could solve this issue. Create a clone of some common SaaS. It doesn’t matter if it’s buggy, in fact that’ll just provide more “learning opportunities” down the line! Get it deployed on a cloud hosting provider like AWS. Then, unleash thousands of agents that each simulate thousands of users of this SaaS: common and uncommon use-cases, different usage patterns, people trying to hack it, anything that a real-world product would run into. Then put one coding agent in charge of the whole thing, with full access to the code and the cloud hosting provider, and instruct it to run the business. Respond to customer complaints, fix bugs, maintain high uptime, implement feature requests, keep the hackers out. Give it a basic coding harness, but also allow it to change its own harness for better context management, etc. And then RL it up!
Of course you’d need a good reward signal, which would probably take some tuning. You could have metrics based on uptime, feature completion, or things like “user” bug reports, but all of those could be reward hacked pretty easily. I think a monetary reward signal would be the most interesting. Each simulated “user” would have a set of feature requests with a corresponding monetary value. They’d also have costs associated with crashes or downtime, and given enough downtime they’d cancel their subscription. Each agent would decide on different priorities for each of their simulated users, and over a large enough simulated customer base, that shouldn’t be an easily hackable metric.
Overall, it would certainly be more expensive than a typical RL harness, but I think it would be much more similar to real-world software development than what you’d get from something like SWE-bench. And I think that would translate into LLMs with taste.
Feel free to wire me $10 million and I’ll make it happen.
My psychologist wife wanted me to mention that Reinforcement Learning is also an aspect of several theories in human psychology! ↩︎
(unless you have load tests on your prod infrastructure in your e2e test suite, in which case you have my respect) ↩︎
Recently Cognition released a new benchmark that captures some of these goals. I think it’s really cool, but many of these goals are inherently subjective and difficult to measure objectively. ↩︎

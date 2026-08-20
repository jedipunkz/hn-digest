---
source: "https://threedots.tech/post/ddd-and-ai-coding/"
hn_url: "https://news.ycombinator.com/item?id=49373041"
title: "DDD matters more when AI writes your code"
article_title: "Domain-Driven Design matters more when AI writes your code"
image: "https://threedots.tech/images/covers/ddd-and-ai-coding-social.png"
author: "roblaszczak"
captured_at: "2026-08-20T11:17:47Z"
capture_tool: "hn-digest"
hn_id: 49373041
score: 1
comments: 0
posted_at: "2026-08-20T11:12:28Z"
tags:
  - hacker-news
  - translated
---

# DDD matters more when AI writes your code

- HN: [49373041](https://news.ycombinator.com/item?id=49373041)
- Source: [threedots.tech](https://threedots.tech/post/ddd-and-ai-coding/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T11:12:28Z

## Translation

タイトル: AI がコードを作成する場合、DDD がより重要になります
記事のタイトル: AI がコードを記述する場合、ドメイン駆動設計がより重要になります
説明: AI エージェントがコードを作成する場合、ドメイン モデルとユビキタス言語がこれまで以上に重要になります。

記事本文:
AI がコードを記述する場合、ドメイン駆動設計がより重要になります
メニュー メニューを開く ホームを閉じる
テーマスイッチャー
電子書籍の検索アイコンを入手
検索するものを入力してください。移動するには
選択する
AI がコードを作成する場合、ESC でドメイン駆動設計を閉じることが重要になります
ミウォシュ・スモウカ
Three Dots Labs の共同創設者。
水車小屋の製作者。
『Go With The Domain』の著者。
AI コーディングについてどう考えても、ソフトウェアを構築する方法は変わりつつあります。
誰もが「何が今後も関連性を維持するのだろうか？」と疑問に思います。
私たちが持っているのは意見だけなので、ここに私の意見があります:
DDD は決してコードを厳密に扱うものではなかったため、ドメイン駆動設計の背後にあるほとんどのアイデアは、これまで以上に関連性が高まっています。
エージェントとのコーディングを増やすにつれて、ドメインを理解し、適切にモデル化して、チームとして作業する必要があります。
DDD が私たちに何を教えてくれるか、そして AI が置き換えられないものを見てみましょう。
いつもの免責事項: 私は、チームが苦戦する傾向があり、DDD が光る場所で、数か月または数年にわたって維持される複雑なプロジェクトについて書いています。
ペット プロジェクトや CRUD で高度なパターンを使用する理由はありません。
2003年から変わらないこと
Eric Evans が『ドメイン駆動設計』を出版したのは 20 年以上前ですが、それでも驚くほど新しいものです。
この本では多くのトピックが取り上げられていますが、主なメッセージは、ソフトウェア エンジニアリングの難しい部分は問題領域を理解し、それをコード内で適切にモデル化することである、ということです。
エンジニアには、技術的な詳細に惑わされず、自分たちが解決していることに集中することが求められます。
その部分を正しく理解するのがより難しいからです。
代わりに、技術的な人材が精巧なフレームワークに取り組み、テクノロジーを使ってドメインの問題を解決しようとします。
ドメインの学習とモデル化は他の人に任せます。
ソフトウェアの中心部にある複雑さには、正面から取り組む必要があります。そうしないと、無関係になる危険があります。
実装が次のとおりであることがわかります。

AI ツールのおかげで、詳細は重要ではなくなりました。
ある程度は、プログラミング言語やフレームワークに関する深い知識がなくても生産性を高めることができ、別の技術スタックに移行するのも簡単です。
トップクラスのコーディング エージェントを使用してみたことがあれば、おそらく、適切なコードがすぐに生成されるのを目にしたことがあるでしょう。
しかし、エージェントをガイドし、出力が意味をなすかどうかを判断するには、ソフトウェアの高レベルの概念を知ることが依然として不可欠です。
何よりも、ドメインの複雑さ (解決しようとしているビジネス上の問題) が依然として難しい部分です。
幸いなことに、コーディングがこれまでよりも簡単になったので、実装の詳細に迷うことなく、これらの複雑な問題に集中できるようになりました。
ただし、フレームワークの構築 (またはマイクロサービスの追加など、クールなトレンドであるもの) にこだわるのではなく、
現在、モデルのベンチマークに従い、エージェントのセットアップを最適化して、より少ない労力でより良いコードを生成します。
Evans が 2003 年に指摘したように、私たちは依然としてテクノロジーを使ってドメインの問題を解決しようとしています。
面白いことに、今回は CEO もそれを信じて、私たちをその方向に推し進めています。
理解すべき小さなことが 1 つだけ残っています。それは、物事がどのように機能するかです。
誰かが私たちに教えてくれたら、実装にトークンを投げて、それを終わらせることができます。
誰も知らない場合は、エージェントが私たちの代わりに計画を書いてくれるかもしれません。
ドメイン モデルは成果物ではありません
DDD の基礎はモデル駆動設計です。つまり、ドメインがどのように機能するかをコードで表現されるモデルに抽出します。
ドメイン モデルは抽象的な概念であり、それを表現する唯一の方法はありません。
ドキュメント、図、コードを使用できますが、それらはすべてドメインの仕組みを簡略化したものです。
正確なモデルは、設計と実装を切り替え、学んだことを何度も適用することで生まれます。
DDD ではこれをナレッジクランチと呼びます

。
ドメインの専門家 (ビジネスの仕組みを知っている人) と、ソフトウェア構築の専門家であるエンジニアが必要です。
彼らは協力して、ソフトウェアが何をすべきか、そしてそれをどのように実装するかを理解します。
(たとえば、イベント ストーミング セッションでは、開発者と関係者が付箋を使用してビジネスの仕組みを計画するワークショップが行われます。)
ドメインを見つけるのはチームの努力であり、大変な作業であるため、他の人にやってもらいたくなります。
たとえば、AI エージェントは調査と分析に優れており、ドキュメント、チャット、コードにアクセスできます。
ドメインモデルを作成して実装もしてくれるようです。
しかし、この素朴なアプローチは的を外しています。
モデルに取り組む価値は、あなた (およびあなたのチーム) がドメインの仕組みを理解できることです。
AI によって生成されたテキストの壁は、ビジネス上の問題を解決するのには役立ちません。
エージェントは、誰も読まない印象的な成果物を作成するだけです。
ドメイン モデルは成果物ではありません。
何かがどのように機能するかについて長い文書を書くことが課題になったことはありません。
多くの場合、ソフトウェア プロジェクトは他の問題が原因で失敗します。
エンジニアは間違ったことに取り組んでいます。
そもそも何をしなければならないのかは誰にも分かりません。
チームはすべてを一度に構築しようとします (スコープクリープ)。
これが私が何度も陥った罠です。
エンジニアリング チームでは、すべてを理解したように感じがちです。
私たちはプログラミングの専門家であり、パターンを理解しており、ベスト プラクティスに従っています。
必要なのは、何をすべきかを教えてくれる人 (ビジネス関係者またはプロダクト マネージャー) だけです。
しかし、多くの場合、ドメインの専門家も明確な計画を持っていません。
彼らは自分たちが抱えている問題は知っていますが、その解決策はまだわかっていません。
すべての要件を網羅した完全な文書はありません。
開発者の最初の本能はこう考えるかもしれません。

ああ、何か思いついたら、すぐそばにいるよ。」
言い換えれば、私たちはコーディングの楽しい部分を求めてここにいます。
(これは AI エージェントが自動化する部分であることに注意してください。)
もう一方の極端な例では、技術系ではないマネージャーがソリューション全体を考え出し、それをチームに渡して実装しようとします。
あるいは、最新の傾向として、誰かが AI を使用して完全な機能や製品仕様を生成し、それを読まずにそのまま渡してしまうこともあります。
つまり、計画はある程度存在していますが、実用化の準備が整ったものには程遠いのです。
(誰かがそうでないと主張する場合は、その人が出荷しているものが実際にどれほど複雑かを確認してください。)
これらすべてのシナリオにおいて、自分たちが何をしているのかを明確に理解している人は誰もいません。
誰がコードを書いても、どれだけ早くコードを書いても、現時点ではほとんど価値がありません。
対照的に、一緒にモデルに取り組む場合、チームは問題領域を理解し、解決策について合意します。
開発者はそれを実装し、予期せぬ事態が発生したり、別の反復が開始されたりした場合に、ドメインの専門家が飛び込むことができます。
ほとんどのコーディングを AI で行う場合でも、チームメイトと同じように、エージェントをガイドするにはそのドメインについて深く知る必要があります。
エージェントをどのように構成するか、どの LLM を使用するかは、ソリューションの明確なメンタル モデルほど重要ではありません。
どちらか選択しなければならないとしたら、何をする必要があるかを考えるのをスキップするくらいなら、自分ではコードをまったく書かないほうがいいと思います。
生成コードを書く前に設計する
コードを書くことは、読むことよりも常に簡単です。
エージェントが大きな機能をワンショットで提供できるようになるのは、さらに極端なことです。
たとえそれが機能したとしても、内部で何が起こっているのかわかりません。
コードレビューで行き詰まることが多くなりましたが、これは新しい問題ではありません。
それは、完全な機能を準備して大々的に宣伝してくれる一匹狼の開発者と仕事をしているような、親しみのあるものさえ感じます。
あなたは何も知りません

デザインには誰も考慮していないギャップが存在することがよくあります。
PRの見直しは大変です。
このような状況を回避する方法はあり、多くのプロジェクトでそれが機能しているのを私は見てきました。
コードを作成する前に、チームとしてソリューションを設計し、話し合います。
まず、ドメイン部分: 何を解決しているのか、そしてそれがどのように機能するのか (上記の知識の詰め込みを参照)。
次に、技術面です。
実装を開始した後、予期せぬことが起こる可能性はほとんどありません。
細部に至るまで同意する必要はありませんし、完全な仕様書を作成する必要もありません。
誰もが知っている大まかな計画を描くだけで十分です。
付箋とホワイトボードがあれば十分です。
デザインを変更する最適なタイミングは、コードを作成する前です。
このようなセッションの後は、PR のレビューがスムーズになります。
PR コメントでデザイン全体について話し合わないと、作業が小さな部分に分割されてしまいがちです。
また、機能が本番環境に導入されると、詳細についてはすでに合意されており、全員が何を構築しているかを理解しているため、ギャップは少なくなります。
設計をスキップすると、時間を節約できるように思えるかもしれません。
会議が増えることを好む人はいません。
ただし、チームが最初に PR からこの機能について知った場合、すべての問題を確認し、議論し、修正するにはさらに長い時間がかかります。
レビューは通常非同期であるため、大量のコメントが作成されることになります。
コードレビューは、実装が正しいかどうかを再確認するものであるべきであり、アプローチがまったく意味があるかどうかについての議論を始めるものではありません。
簡単な計画から始めると、後で検討する時間を節約できます。
コードを書く前に設計に取り組むことは、AI エージェントとの連携を改善するのにも役立ちます。
より高品質なコンテキストをエージェントに提供すると、その出力はより正確になり、コーディング部分は短くなります。
そしてチームが増えれば増えるほど、

解決策を知っていればいるほど、すべてのコードを生成した場合でもレビューが容易になります。
ユビキタス言語: エージェントと同じ言語を話す
最近、クラウドのコストを削減しようと努めています。
測定すべき明確な指標があるため、これは AI にとって完璧なタスクです。
私はエージェントに、Go サービスが最も多くのメモリを使用している場所を見つけて、それを削減するように言いました。
数回実行した後、何メガバイトもの RAM を節約できました。
しかしその後、メモリがかなり安いことに気づき、アップデートによりコストが月あたり 1 ドル削減されました。
驚くことではありませんね？コストを削減しようとしていることをエージェントに説明しませんでした。
メモリ使用量を目標として指定しただけです。
これはまさにチーム内の人間の間で起こっていることなので、よく知られているように聞こえるはずです。
チーム内でのコミュニケーションに役立つ優れた実践方法は、AI を使用する場合にも役立ちます。
(ソフト スキルがコンピュータを扱うのに役立つかもしれないと誰が想像したでしょうか?)
DDD は、使用する言語に関するいくつかのアイデアでこの問題に対処します。
1 つ目はユビキタス言語です。チーム間およびコード内で同じ言語を話すため、誰もが何を話しているのかがわかります。
多くのチームでは、乱雑な名前がデフォルトになっています。
コードでは、エンジニアは社内の他の部門で使用されている名前とはかけ離れた専門用語を使用します。
AI 文書や意思決定記録を入力したとしても、エージェントはそこで使用されている言語を理解する必要があります。
知識を詰め込む際には、ドメイン内の概念を何と呼ぶか​​を検討してください。
最初は、それぞれに複数の名前があるのが一般的です。
すでに使用している名前を選択するか、より正確な名前を考え出します。
次に、それを文書、話し言葉、コードで使用します。
イベント ストーミング セッションは、名前を選択するのに最適な場所です。
言語はプロジェクトとともに進化し続けるため、名前を一度だけ選択することはできません。
それは

設計セッションやディスカッション中に注意する必要があること。
ユビキタス言語に直接関係するもう 1 つのアイデアは、より大きなシステムでは、選択した名前がすべての領域で普遍的になるわけではないということです。
DDD では、これらの各領域を境界コンテキストと呼びます。
顧客は、サポート コンテキストではプロファイルとして、e コマース コンテキストではユーザーとして考えることができます。
プロジェクト全体に名前を強制するのではなく、各コンテキスト内で一貫した名前を使用する必要があります。
これは、チーム間でソフトウェアを分割する必要がある大規模な製品の場合に特に重要です。
あらゆる種類のタスクに対して 1 つのモデルを作成することに集中しすぎると、境界線を引くことが難しくなります。
境界のあるコンテキストを念頭に置くと、最終的に分離されたシステムになりやすくなります。
エージェントはリポジトリ全体にアクセスでき、指定された名前を検索します。
彼らは素朴に似たような存在を統合しようとするかもしれないので、それらが理由があって別々であることを明確にしてください。
同じドメイン概念は、コンテキストに応じて異なる名前を持つことができます。
よく知られた正確な概念を使用すると、言いたいことを表現しやすくなります。
チームメイトとの共同作業でも、プロンプト内でも機能します。
プロンプトが曖昧だと、エージェントは間違った解決策に多大な労力を費やすことになります。
CRM にユーザーを追加し、作成後のサポートを行う
正確さを守ることで、目標がより明確になります

[切り捨てられた]

## Original Extract

The domain model and Ubiquitous Language matter more than ever when an AI agent writes your code.

Domain-Driven Design matters more when AI writes your code
Menu Open Menu Close Home
theme switcher
Get our e-book search icon
Type something to search.. to navigate
to select
ESC to close Domain-Driven Design matters more when AI writes your code
Miłosz Smółka
Co-founder of Three Dots Labs .
Creator of Watermill .
Author of Go With The Domain .
Whatever you think of AI coding, the way we build software is changing.
Everyone wonders, “What will stay relevant?”
All we have is opinions, so here’s mine:
most ideas behind Domain-Driven Design are now more relevant than ever, as DDD has never been strictly about code.
As we do more coding with agents, we still need to understand the domain, model it well, and work as a team.
Let’s see what DDD can still teach us and what AI can’t replace.
The usual disclaimer: I write about complex projects maintained over months or years, where teams tend to struggle and where DDD shines.
There’s no reason to use advanced patterns in pet projects or CRUDs.
What hasn’t changed since 2003
Eric Evans published Domain-Driven Design more than two decades ago, but it’s still surprisingly fresh.
The book covers many topics, but the main message is that the hard part of software engineering is understanding the problem domain and modeling it well in code.
It calls on engineers not to get lost in the technical details and to focus on what they’re solving,
because it’s more difficult to get that part right.
Instead, the technical talent goes to work on elaborate frameworks, trying to solve domain problems with technology.
Learning about and modeling the domain is left to others.
Complexity in the heart of software has to be tackled head-on. To do otherwise is to risk irrelevance.
We now see that the implementation details become less relevant thanks to AI tools.
To a degree, you can be productive without a deep knowledge of a programming language or a framework, and it’s easier to move to another tech stack.
If you’ve tried using the top coding agents, you’ve probably seen them quickly generate decent code.
But knowing high-level software concepts is still essential to guide agents and to decide whether the output makes sense.
Most of all, the domain complexity (the business problem you’re solving) is still the hard part .
The good news is that coding is easier than ever, so we can focus on these complex problems instead of getting lost in the implementation details.
Except, instead of obsessing over building the frameworks (or adding more microservices, or whatever else is a cool trend),
we now follow the model benchmarks and optimize our agentic setup to generate better code with less effort.
As Evans noted back in 2003, we still try to solve the domain problems with technology.
Funnily enough, this time the CEOs also believe it and push us toward it.
There’s just one small thing left to figure out: how the thing should work .
If someone tells us, we can throw tokens at the implementation and call it a day.
And if no one knows, well, maybe an agent could write the plan for us?
The domain model isn’t an artifact
The foundation of DDD is model-driven design: distilling how the domain works into a model that’s expressed in code.
The domain model is an abstract concept, and there’s no one way to represent it.
You can use documents, diagrams, and code, but they are all simplifications of how the domain works.
An accurate model comes from switching between design and implementation, and applying what you’ve learned over and over.
DDD calls this knowledge crunching .
You need domain experts (the people who know how the business works), and engineers who are experts at building software.
They work together to understand what the software should do and how to implement it.
(For example, in an Event Storming session, a workshop where developers and stakeholders map out how the business works with sticky notes.)
Figuring out the domain is a team effort and hard work, so it’s tempting to have someone else do it.
For example, AI agents are brilliant at research and analysis, and can access your documents, chats, and code.
It seems they could create the domain model for you, and then also implement it.
But this naive approach misses the point.
The value of working on the model is that you (and your team) understand how the domain works .
An AI-generated wall of text doesn’t help you figure out the business problem.
The agent will just create an impressive artifact no one reads.
A domain model isn't an artifact.
The challenge has never been to write a long document on how something could possibly work.
More often, software projects fail because of other issues:
Engineers work on the wrong thing.
No one knows what needs to be done in the first place.
The team tries to build everything at once (scope creep).
Here’s one trap I’ve fallen into many times.
In engineering teams, it’s easy to feel like we’ve got it all figured out:
we’re programming experts, know our patterns, and follow best practices.
We just need someone (a business stakeholder or a product manager) to tell us what to do.
But often, the domain experts don’t have a clear plan either.
They know the problem they have, but not what the solution is yet.
There’s no complete document with all requirements.
A developer’s first instinct may be to think, “Let me know once you come up with something, I’ll be around.”
In other words, we’re here for the fun part of coding.
(Note this is the part AI agents automate.)
At the other extreme, non-technical managers try to come up with the entire solution and hand it to the team to implement.
Or, the most recent trend, someone uses AI to generate a full feature or product spec, doesn’t even read it, and passes it along.
So the plan sort of exists, but it’s nowhere near something ready for production.
(If someone claims otherwise, check how complex the thing they’re shipping really is.)
In all of these scenarios, no one has a clear idea what they’re doing.
No matter who writes the code or how fast, it has little value at this point.
In contrast, when you work on the model together, your team understands the problem domain and agrees on a solution.
Developers implement it, and domain experts can jump in once something unexpected comes up or another iteration starts.
Even if you do most coding with AI, you need to get deeply familiar with the domain to guide the agent, just as you would a teammate.
How you configure the agents and what LLM you use don’t matter as much as a clear mental model of the solution.
If I had to choose, I’d rather never write any code myself than skip thinking about what needs to be done.
Design before writing generating code
Writing code has always been easier than reading it.
It’s just more extreme now that an agent can single-shot a big feature for you.
Even if it works, you have no idea what’s going on under the hood.
We now get stuck on code review more often, but it’s also not a new problem.
It even feels familiar, like working with a lone-wolf developer who prepares complete features and drops a massive PR on you.
You know nothing about the design and there are often gaps that no one has considered.
Reviewing the PR is painful.
There’s a way to avoid situations like this, and I’ve seen it work in many projects.
You design and discuss the solution as a team before sitting down to code.
First, the domain parts: what you’re solving and how it’s supposed to work (see knowledge crunching above).
Then, the technical side.
After you kick off the implementation, there’s little chance of anything unexpected coming up.
You don’t need to agree on every tiny detail, and don’t need a complete written spec.
It’s enough to sketch a high-level plan that everyone is aware of.
Some sticky notes and a whiteboard are good enough.
The best moment to change the design is before anyone writes code.
After sessions like this, reviewing PRs is smooth.
It’s easy to split the work into small chunks if you don’t discuss the entire design in PR comments.
And there are fewer gaps once the feature lands in production, because you’ve already agreed on the details and everyone understands what they’re building.
Skipping the design may seem like saving time.
No one likes more meetings.
But if your team first learns about the feature from the PR, it takes much longer to review, discuss, and fix all the issues.
Since reviews are usually asynchronous, you end up with many rounds of comments.
Code review should be a double-check that the implementation is correct, not the start of a discussion about whether the approach makes sense at all.
Starting with a brief plan saves you time on review later on.
Working on design before code will also help you work better with AI agents.
If you give the agent more high-quality context, its output will be more accurate and the coding part will be shorter.
And the more your team knows about the solution, the easier the review, even if you generate all the code.
Ubiquitous Language: Speaking the same language as your agents
Recently, I’ve been trying to cut our cloud costs.
It’s the perfect task for AI, since there’s an obvious metric to measure.
I told the agent to find where our Go service uses the most memory and to cut it.
After a couple of runs, I saved us many megabytes of RAM.
But then I realized memory is pretty cheap, and the updates reduced our costs by $1/month.
Not surprising, right? I didn’t explain to the agent that I was trying to cut costs.
I just gave it the memory usage as the target.
It should sound familiar, because it’s exactly what happens between humans in teams.
Good practices that help you communicate in a team are also relevant when you work with AI.
(Who would have thought soft skills may help you deal with computers?)
DDD addresses this with a few ideas around the language you use.
The first is Ubiquitous Language : speaking the same language across teams and in code, so everyone knows what you’re talking about.
In many teams, messy names are the default.
In code, engineers use technical terms that drift away from the names that the rest of the company uses.
Even if you feed AI documents and decision records, the agent needs to make sense of whatever language you use there.
During knowledge crunching, consider what you call the concepts in the domain.
Initially, it’s common to have more than one name for each.
Choose a name you already use, or come up with a more accurate one.
Then, use it in the documents, spoken language, and code.
An Event Storming session is a great place for choosing the names.
You can’t choose the names just once, because the language will keep evolving with the project.
It’s something you need to care about during the design sessions and discussions.
Another idea directly connected to the Ubiquitous Language is that in bigger systems, the names you choose won’t be universal in all areas.
DDD calls each of these areas a Bounded Context .
You can think of your customer as a profile in the support context, and as a user in the e-commerce context.
You want to use consistent names within each context, not force them across the entire project.
This is especially important with bigger products, when you need to split the software between teams.
If you focus too much on creating one model for all kinds of tasks, it becomes difficult to draw boundaries.
With bounded contexts in mind, it’s easier to end up with a decoupled system.
Agents have access to your whole repository, and they’ll look for the names you give them.
They may naively try to unify similar entities, so make it clear they are separate for a reason.
The same domain concept can have different names depending on the context.
If you use well-known, precise concepts, it’s easier to express what you mean.
It’ll work both with your teammates and in your prompts.
A vague prompt can make the agent spend a lot of effort on the wrong solution:
Add user to CRM and support after it’s created
You’ll make the goal more obvious by sticking to the preci

[truncated]

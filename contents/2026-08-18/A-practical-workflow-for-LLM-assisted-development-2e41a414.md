---
source: "https://yogthos.net/posts/2026-08-17-llm-workflow.html"
hn_url: "https://news.ycombinator.com/item?id=49345342"
title: "A practical workflow for LLM-assisted development"
article_title: "(iterate think thoughts): A practical workflow for LLM-assisted development"
image: ""
author: "yogthos"
captured_at: "2026-08-18T13:35:31Z"
capture_tool: "hn-digest"
hn_id: 49345342
score: 1
comments: 0
posted_at: "2026-08-18T13:27:26Z"
tags:
  - hacker-news
  - translated
---

# A practical workflow for LLM-assisted development

- HN: [49345342](https://news.ycombinator.com/item?id=49345342)
- Source: [yogthos.net](https://yogthos.net/posts/2026-08-17-llm-workflow.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T13:27:26Z

## Translation

タイトル: LLM 支援開発の実践的なワークフロー
記事のタイトル: (反復思考): LLM 支援開発の実践的なワークフロー
説明: プログラミング、Clojure、ソフトウェア開発に関する Dmitri のブログ

記事本文:
(反復思考): LLM 支援開発の実践的なワークフロー
(反復
考える
思い）
テーマ
2026 年 8 月 17 日
LLM 支援開発の実践的なワークフロー
LLM が機能するときは魔法のように感じられますが、失敗すると、自信に満ちたデタラメなアーティストと議論しているような気分になります。 LLM が有用なコードを生成する可能性が高い場所と、失敗する可能性が高い場所をある程度直感的に理解できるようになるまで、毎日何ヶ月もかかりました。また、有効な結果を確実に得るために範囲を制限し、十分な足場を提供する方法を理解するのに少し時間がかかりました。このツールの効果的な使い方を学ぶために時間を投資したので、以前は試みることのなかった規模のプロジェクトを構築できるようになり、そのメリットを非常に実感しています。
ある意味、このプロセスは通常のプログラミングの逆です。手作業でコードを記述する場合、意図を持って機能を追加していくため、段階的にプログラムを構築する傾向があります。 LLM は最初から大量のコードを生成する傾向があり、実際に必要なものまでコードを絞り込むことに焦点が移ります。
エージェント ループを理解する良い方法は、プロセスを遺伝的アルゴリズムとして見ることです。進化のプロセスが起こっているため、エージェントのハーネスが効果的です。コードがテストされる前に、モデルはおおよそ正しいものを出力し、その後モデルはコードを反復するためのフィードバックを取得します。このプロセスを通じて、テスト対象のパラメーターに適合するソリューションに徐々に収束します。その意味では、実際には人間がコードを書く方法とそれほど変わりません。重要な問題を一発で解決することはほとんどありません。最初の近似値を作成し、それを反復処理します。違いは、LLM の方がこのプロセスをはるかに高速に実行できることです。
LLM は大量の公開コードでトレーニングされるため、

典型的なタスクを完了するのに優れています。これらはこれまでに何百万回も行われてきたことであり、定型的な部分がほとんどです。サンプルの JSON 応答を LLM にスローしてサービス エンドポイントを作成させることや、多数の API エンドポイントをスローしてそれらを使用して UI を構築させることは、非常に効果的です。これらの種類の一般的なタスクは、エージェントが多くのトレーニングを受けており、一発で妥当なものを生成できます。おそらく、テストを追加したり、明らかなエッジケースをすべて処理したりするよりも、そのタスクにさらに熱心に取り組むことになるでしょう。
彼らは探索的な作業を行うことにも優れています。特定のコール グラフを識別し、特定のサービス エンドポイントがどのように実装されているか、またはどのパラメータを渡す必要があるかを理解するための手順を追跡することは、LLM が簡単に実行できるタスクです。これにより、コードベースをトレースし、関心のある特定のワークフローをマッピングする時間を大幅に節約できます。
これらのツールは、言語固有の構文の処理にも優れています。コレクションをループしたり、特定のパラメーターでフィルターしたりするなど、やりたいことは概念的にはわかっているものの、使い慣れていない言語で作業している場合、LLM はギャップを埋めるのに最適です。慣用的な構文を使用して、必要なロジックを簡単に表現できます。疑似コードでアルゴリズムを記述し、ステップを書き出すと残りが処理されます。
たとえば、私は最近 JavaScript プロジェクトに取り組む必要がありましたが、この言語には 10 年以上触れていませんでした。私は最新のツールやライブラリ、ベスト プラクティスに詳しくなく、それらすべてを理解する時間がありませんでした。
DeepSeek を使用すると、私がよく知っている Clojure と同じくらい JavaScript を効果的に使用できるようになりました。

構文やツールなどの付随的な事柄をすべて理解するのに苦労することはありませんでした。あなたが特定の分野の専門家であり、解決しようとしている問題を理解している場合、LLM はあなたの能力を大きく高めることができます。これらはあなたのスキルに取って代わるものではありませんが、より速く行動できるようになり、解決しようとしている問題の全体像に集中できるようになります。
私の経験では、エージェントがつまずく最大の場所は、コンテキストと創造性を扱うことです。 AI はプロジェクトの具体的な特徴を認識しないことに注意する必要があります。たとえば、Clojure 方言を使用するように指示した場合、Clojure を学習した JVM ツールチェーン (そのコンテキストには存在しない clojure や lein など) に到達するか、ツリー ウォーキング インタプリタを想定してソースを直接実行しようとする可能性があります。ランタイムが純粋な Chez Scheme であり、すべてが chez --script の実行を介して make コマンドによってビルドされることを明示的に伝える一方で、信頼できるソースが JVM 風味のものよりも host/chez/*.ss および jolt-core/*.clj であることを指定するなど、正確なロジックを与える必要があります。
さらに、素朴な実装の罠もあります。エージェントに漠然とした目標を与えると、表面的には正しいように見えても、最終的には構造的に間違っているものを渡されることがよくあります。たとえば、エージェントは、文字列メソッド呼び出しを汎用ディスパッチ テーブル経由でルーティングする必要があると決定する場合があります。これにより、呼び出しごとにレシーバー タイプが再導出されることになります。ここでの適切な修正は、型推論パスを実行して、コンパイル時にそれらの値が文字列であることを証明することです。これにより、直接ネイティブ呼び出しを発行し、ディスパッチを完全にスキップできるようになります。文字列メソッドを高速化するように指示されたエージェントは、ほぼ確実に reorderin によって汎用パスを維持します。

いくつかの制御アームを使用するだけで、適切なソリューションを設計する必要はありません。同様に、シーケンスに count を実装するように要求した場合、コレクションが定数時間で呼び出すことができる独自の長さをすでに知っている場合、コレクションは要素ごとに新しいセルを割り当てて全体を実行する可能性があります。文字列を結合するように要求すると、おそらく 1 回のウォークではなく、繰り返しの連結が行われるでしょう。これは、クエリを可能な限り最悪の方法で解釈し、完全に間違った形のソリューションを導き出す邪悪な魔神のようなものです。重要なのは、制約を詳しく説明する必要があるため、問題についてもよく考える必要があるということです。
LLM を効果的に使用するための鍵は、開始する前に構築しようとしているものをしっかりと理解していることを確認することです。構造レベルで何を実行したいのかを常に明確にする必要があります。事前に提供する足場が多いほど、エージェントが設計の外に出る余地が少なくなります。この観察から必然的に、LLM を効果的に使用するにはドメインを理解する必要があるということになります。生成されたコードが問題を正しい方法で解決するかどうかを評価する能力が備わっていない場合、基本的にカジノでスロット マシンのレバーを引いて、まともな解決策が見つかることを期待することになります。 LLM はギャップを埋めたり、定型文を実行したりするのが得意ですが、それでも、これまでと同じ方法で設計とアーキテクチャを実行する必要があります。
ここでは、軌道に乗るために役立つと私が感じたいくつかのトリックを紹介します。
常にタスクを計画することから始めます。何をしようとしているのか、どのようなアルゴリズムを使用するつもりなのか、既存のアーキテクチャに適合するようにコードをどのように構成する必要があるのか​​を明確に把握してください。事前にこれらの質問に答えられるようにする必要があります

LLM に委任することを検討してください。
頭の中に明確なイメージができたら、エージェントとの計画段階に進むことができます。モデルに Markdown で段階的な計画を書くよう依頼する前に、モデルに要件を与え、目標を詳しく説明します。さらに良いのは、フローの Mermaid.js 図を生成するように依頼することです。
図を作成した後、ロジックを視覚的に検査できます。図内の特定のステップが間違っているように見える場合は、その特定のステップを変更して別のことを行うように指示します。これは、テキスト プロンプトを使用して単に議論するよりもはるかに簡単です。実行するステップの構造が明確になれば、気に入らない部分を簡単に特定できます。計画をレビューし、モデルを独立したタスクに分割し、それぞれが特定の機能の実装に焦点を当てます。モデルにブランチを作成してから、タスクのプル リクエストを作成させます。その時点で、変更の範囲とそれによって解決される具体的な問題がわかるため、コードをかなり簡単にレビューできます。
どのアプローチをとるべきかわからないステップについては、モデルに以前の作業を調査させると非常に役立ちます。解決しようとしている問題がまったく新しいものであることはまれであり、エージェントは、何が効果的であるかをよりよく理解するためにレビューできる関連文書を検索するのに最適です。繰り返しになりますが、時間をかけて自分が選択できるさまざまな道を理解し、意識的に道を選択することが重要です。
また、LLM を使用する場合、低結合のクリーンなアーキテクチャが非常に重要になるとも主張します。考慮すべきコンテキストが少ないため、依存関係のない小さなタスクで最も効果を発揮する傾向があります。したがって、プロジェクトを個別に作業できる小さな部分に分割できれば、次のようにエージェントにタスクを与えることができます。

明確な境界線。これにより、出力のレビューも非常に簡単になります。
関数型スタイルは、コンテキストの分離と状態の明示的な受け渡しに焦点を当てているため、ここでは特によく対応していることがわかりました。大規模なコード ベースを人間が管理できるようにする同じトリックは、同じ理由で LLM にも役立ちます。コンテキストを積極的に制御することは、LLM を効果的に使用するための重要な戦術です。
繰り返しになりますが、AI に空白のキャンバスを決して与えたくないのです。足場をどのように配置するかという作業は、必ず自分で行ってください。エージェントに空白を埋めるよう依頼する前に、意図的にファイル構造を設定し、コンポーネントを決定してください。
しかし、これらすべての優れた機能ツールを備えているにもかかわらず、依然として 2 つのかなり異なる種類のコードが複雑に絡み合う傾向があります。私たちは、データの意味を考慮するコードと、あるコンポーネントから別のコンポーネントにデータがどのように移動するかを決定するコードを混合する傾向があります。従来のソフトウェア設計構造では、ルーティングが関数呼び出しグラフに暗黙的に埋め込まれています。制御ロジックは、多くの場合、最終的にはその場限りの方法で内部実装の詳細と結合されることになります。物事を独立したステップに分割すると、範囲を制御するのに役立ちます。
ルーティング ロジックは、設計において第一級の市民権に昇格する必要があります。ステート マシンは、何を行うか、どのように行うかを強制的に分離するため、これに自然に適合します。制御フローのロジックは主に宣言型であり、前述の人魚図のようなグラフとして表現できますが、実装の詳細はフローの各ステップに存在し、エージェントが作業するタスクになります。
これらの手順を実行すると、エージェントは独自の構造を発明するのではなく、アーキテクチャ内で動作するようになり、レールから外れてしまう問題が大幅に回避されます。取得したら図を作成します

それを確認したら、それに基づいて初期プロジェクト構造を作成できます。
LLM を扱うときは、テストを最終的な要件ドキュメントとして考えると便利だと思います。最初に必要な機能をテストとして定義すると、テスト駆動開発を使用して、テストが合格するまでエージェントに機能させることができます。通常、テストを実行し、障害を分析し、仕様を満たすように独自のコードを修正するという適切な仕事を行います。テストは、エージェントが取り組む契約です。遺伝的アルゴリズム全体のアナロジーに戻ると、これらはコードの進化を促す選択圧力です。
事前にテストを行うことで、コードが機能的に意図したとおりに動作していることを確実に保証できます。これは、リグレッションに対する最善の防御策でもあります。テストがなければ、エージェントが新しい機能を追加すると、3 つの古い機能が静かに破壊される可能性があります。既存の機能を契約すると、この問題を回避できます。
最も価値がある傾向にあるテストの種類は、エンドツーエンドの統合テストとともにさまざまなコンポーネントの機能に焦点を当てたものです。ワークフロー全体が実行されるにつれて問題は解消されるため、あまり細かくする必要はありません。

[切り捨てられた]

## Original Extract

Dmitri's blog about programming, Clojure, and software development

(iterate think thoughts): A practical workflow for LLM-assisted development
( iterate
think
thoughts )
Theme
August 17, 2026
A practical workflow for LLM-assisted development
When LLMs work, it can feel like magic, but when they fail, it feels like you are arguing with a confident bullshit artist. It took me many months of daily use to develop some intuition for where LLMs are likely to produce code that is useful and where they are likely to fail. It also took me a bit of time to figure out how to limit scope and provide enough scaffolding to ensure I get useful results reliably. Having invested the time to learn to use the tool effectively, I very much see the benefits, as I am able to build projects on a scale I would not have attempted before.
In a way, the process is the inverse of regular programming. We tend to build up programs step by step when writing code by hand as we add each function with intention. LLMs tend to produce a lot of code out of the gate and the focus shifts to whittling the code down to what you actually need.
A good way to look at the agentic loop is to view the process as a genetic algorithm. Agentic harnesses are effective because you have an evolutionary process happening. The model outputs something roughly correct before the code gets tested, and then the model gets feedback to iterate on the code. Through this process, it gradually converges on a solution that fits the parameters being tested. In that sense, it is not actually all that different from how humans write code either. You almost never solve a non-trivial problem in one shot. You write your first approximation and then iterate on it. The difference is that the LLM can do this process a lot faster.
LLMs are trained on massive amounts of public code, which makes them excellent at completing typical tasks. These are things that have been done a million times before and constitute what largely amounts to boilerplate. Throwing a sample JSON response at an LLM and having it write a service endpoint or throwing a bunch of API endpoints at it and having it build a UI using them can be very effective. These are the kinds of common tasks the agent will have a lot of training on, and they can produce something reasonable in one shot. It will probably put more diligence into that task than you would by adding tests and handling all the obvious edge cases.
They are also great at doing explorative work. Identifying a particular call graph and tracing through the steps to figure out how a particular service endpoint is implemented or what parameters you have to pass it are all tasks an LLM can do easily. This can save an enormous amount of time tracing through a codebase and mapping out a particular workflow that you are interested in.
These tools are also great at handling language specific syntax. If you know conceptually what you want to do, like looping through a collection and filtering by a specific parameter, but you are working in a language you are rusty in, then LLMs are great for bridging the gap. They can easily express the logic you want using idiomatic syntax. You can describe the algorithm in pseudo code where you write out the steps and it will handle the rest.
For example, I recently had to work on a JavaScript project, and I have not touched the language in over a decade. I am not familiar with modern tooling or libraries or best practices, and I just did not have the time to get up to speed on all that.
Using DeepSeek allowed me to use JavaScript as effectively as I do Clojure, which I am well versed in. It completely removed the friction of figuring out all the incidental things like syntax or tooling. If you are an expert in a particular domain and you understand the problem you are trying to solve, then LLMs can be a huge amplifier for what you are able to do. They do not replace your skills, but they do allow you to move a lot faster and focus on the big picture of the problem you are trying to solve.
In my experience, the biggest place where agents trip up is dealing with context and creativity. You have to remember that the AI does not know the specific quirks of your project. For example, if you just tell it to use a Clojure dialect, it might reach for the JVM toolchain it learned Clojure on, such as clojure and lein , none of which exist in that context, or it might assume a tree walking interpreter and try to run the source directly. You need to give it the exact logic, like telling it explicitly that the runtime is pure Chez Scheme and that everything builds through make commands via a chez --script execution, while specifying that the authoritative sources are host/chez/*.ss and jolt-core/*.clj over anything JVM flavored.
Then there is also the trap of the naive implementation. Often, when you give an agent a vague goal, it will hand you something that looks correct on the surface but ends up being structurally wrong. For example, the agent might decide that string method calls should be routed through a generic dispatch table, which ends up re-deriving the receiver type on every single invocation. The proper fix here is to do a type inference pass to prove that those values are strings at compile time, which allows you to emit a direct native call and skip dispatch entirely. An agent told to make the string methods fast will almost certainly keep the generic path by reordering a few cond arms and never bother designing a proper solution. Similarly, if you ask it to implement count on a sequence, it will likely walk the whole thing allocating a fresh cell per element when the collection already knows its own length that can be called in constant time. Ask it to join strings and you will probably get repeated concatenation instead of a single walk. It is akin to an evil genie that will interpret your queries in the worst way possible, leading to the solution having a completely wrong shape. The trick is that you have to spell out the constraint, which incidentally forces you to think through the problem as well.
The key to using LLMs effectively is to make sure you already have a solid understanding of what you are aiming to build before you start. You always have to be explicit regarding what you want done at a structural level. The more scaffolding you provide up front the less room the agent has to go outside your design. A corollary to this observation is that you do have to understand the domain to make effective use of LLMs. If you are not equipped to evaluate whether the code it produced solves the problem in a correct way, then you basically end up at a casino pulling a lever on a slot machine and hoping for a decent solution to fall out. LLMs are good at filling in the gaps and doing boilerplate, but you still have to do design and architecture the same way you always did.
Here are some tricks that I found useful for keeping it on the rails.
Always start out by planning out the task. Make sure you have a clear picture of what you are aiming to do along with what algorithms you are intending to use and how the code should be structured to fit within the existing architecture. You must be able to answer these questions before you even think about delegating to the LLM.
Once you have a clear picture in your head, you can move on to the planning stage with the agent. Give it the requirements and spell out the goals before asking the model to write a phased plan in Markdown. Even better, ask it to generate a Mermaid.js diagram of the flow.
After it makes the diagram, you can visually inspect the logic. If a particular step looks wrong in the diagram, you tell it to change that specific step to do something else. Doing that is a lot easier than simply arguing with it using text prompts. Once there is a clear structure for the steps being performed, it is easy to identify parts that you do not like. Review the plan and get the model to break it up into independent tasks, each focusing on implementing a specific feature. Have the model create a branch and then make a pull request for the task. At that point you can review the code fairly easily because you know what the scope of the change is and what specific problem it solves.
It can be very helpful to have the model do research on prior work for steps where you are not sure which approach to take. It is rare that the problem being solved is entirely novel, and agents are great for looking up relevant papers you can review to get a better idea of what is more likely to work. Again, it is important to spend the time to familiarize yourself with the different paths you can take and to pick one consciously.
I would also argue that having a clean architecture with low coupling becomes extremely important when using LLMs. They tend to do best on smaller tasks that do not have dependencies because there is less context to consider. So if you can break up your project into small pieces that can be worked on in isolation, then you can give the agent a task with clear boundaries. That also makes it much easier to review its output as well.
I find that functional style maps particularly well here because it focuses on context isolation and passing state around explicitly. The same tricks that make large code bases manageable by humans also help LLMs for the same reasons. Aggressively controlling the context is a key tactic for using LLMs effectively.
It bears repeating that you never want to give the AI a blank canvas. Always do the work of laying out what the scaffolding should look like yourself. Make sure you intentionally set up the file structure and decide on the components before asking the agent to fill in the blanks.
But even with all these great functional tools, we still tend to tangle two rather different kinds of code together. We tend to mix code that cares what the data means and the code that decides how it travels from one component to another. Traditional software design structures embed the routing implicitly in the function call graph. Control logic often ends up being coupled with the internal implementation details in an ad hoc manner. Breaking things up into independent steps helps control the scope.
Routing logic should be elevated to first class citizenship in the design. State machines are the natural fit for this, since they force the separation of what to do from how to do it. The control flow logic can be largely declarative and expressed as a graph such as the Mermaid diagram I mentioned earlier, while the implementation details live at each step in the flow and become the tasks the agent works on.
Doing these steps forces the agent to work within your architecture rather than inventing its own structure, which largely avoids the problem of it going off the rails. Once you get it to build a diagram and you have reviewed it, you can create the initial project structure based on that.
I find it is useful to think of tests as the ultimate requirement doc when working with LLMs. If you define your desired functionality as tests first, you can get the agent to work through them using test driven development until they pass. It will typically do a decent job running the tests and analyzing the failures and fixing its own code to meet the spec. The tests are the contract that the agent works against. Going back to the whole genetic algorithm analogy, these are the selection pressures that drive the evolution of the code.
Having tests up front gives you a solid guarantee that the code is doing what you intended functionally. It is also your best defense against regressions. Without tests, an agent adding a new feature is just as likely to silently break three old ones. Having a contract for the existing functionality avoids that problem.
The types of tests that tend to be most valuable are the ones that focus on the functionality of different components along with end to end integration tests. They do not need to be too granular because issues will get shaken out as the whole workflow gets exercis

[truncated]

---
source: "http://www.stephen-cresswell.com/2026/08/15/Yadda-3.0.0-BDD-in-the-Age-of-AI-Agents.html"
hn_url: "https://news.ycombinator.com/item?id=49310495"
title: "Yadda 3.0.0: BDD in the Age of AI Agents"
article_title: "Yadda 3.0.0: BDD in the Age of AI Agents | Signal Over Noise"
author: "scresswell"
captured_at: "2026-08-15T14:13:51Z"
capture_tool: "hn-digest"
hn_id: 49310495
score: 3
comments: 1
posted_at: "2026-08-15T13:43:46Z"
tags:
  - hacker-news
  - translated
---

# Yadda 3.0.0: BDD in the Age of AI Agents

- HN: [49310495](https://news.ycombinator.com/item?id=49310495)
- Source: [www.stephen-cresswell.com](http://www.stephen-cresswell.com/2026/08/15/Yadda-3.0.0-BDD-in-the-Age-of-AI-Agents.html)
- Score: 3
- Comments: 1
- Posted: 2026-08-15T13:43:46Z

## Translation

タイトル: Yadda 3.0.0: AI エージェント時代の BDD
記事のタイトル: Yadda 3.0.0: AI エージェント時代の BDD |シグナルオーバーノイズ
説明: Yadda 3.0.0 がリリースされました。このリリースでは JavaScript BDD ライブラリが最新化されていますが、さらに興味深いことに、このライブラリは主に Claude Code によって構築されており、エージェント開発の世界で実行可能仕様がさらに価値を持つ可能性がある理由を示しています。

記事本文:
シグナルオーバーノイズ
ホーム
Yadda 3.0.0: AI エージェント時代の BDD
Yadda 3.0.0 を npm に公開しました。
馴染みのない人のために説明すると、Yadda は JavaScript 用の BDD ライブラリです。 Cucumber と同様に、通常の言語仕様を実行可能コードにマッピングしますが、それらの仕様がどのように記述されるかについてあまり規範的ではないように最初から設計されています。
つまり、次のようなものを書く代わりに、次のようになります。
大学を指定すると、ブーベ島大学
ブーベ島大学は、ABB の入学要件を持つコンピューター サイエンスの学位コースを提供しています。
そして A レベル卒業生のスティーブ
そしてスティーブは物理学のDを持っています
そしてスティーブは数学で D を持っています
スティーブがブーベ島大学でコンピューター サイエンスの勉強を志願したとき
その後、ブーベ島大学が申請を拒否
次のように書くことができます:
ブーベ島大学はコンピューター サイエンスの学位コースを提供しています
ABBの参加要件は
スティーブは A レベル卒業生です
物理学で D を取得した場合
そして数学のD
スティーブがブーベ島大学でコンピューター サイエンスの勉強を志願したとき
彼らは彼の申請を拒否した
どちらも実行可能な仕様です。 2番目の方がかなり読みやすいと思います。
Yadda 3.0 の大部分は最新化の取り組みです。
Yadda は長い間存在しており、リポジトリには JavaScript エコシステムの一部の統合とツールが蓄積されていましたが、それら自体が歴史的な珍品となっています。 Yadda 3 は Node-only で、ブラウザのバンドルと CasperJS、PhantomJS、Bower、Component などの古い統合を削除し、テスト スイートを node:test に移動し、Biome と lefthook を採用し、ソースを ES6 構文に最新化し、Playwright や Puppeteer などの最新のサンプルを追加します。 TypeScript 定義も同梱されるようになりました。
どれも便利ですが、特に int ではありません

書くのが楽しみです。このリリースに関して、より重要だと思うことが 2 つあります。
Opus 4.8 の Claude Code を使用して Yadda を最新化しました。
Yadda 3.0 エピック自体は Claude によって書かれており、作業を意図的に分割された一連のフェーズに分割しました。つまり、廃止された機能の削除、ツールチェーンの更新、動作の変更とは別に機械的フォーマットの実行、ソースの最新化、API の変更の調査、サンプルと CI の更新、そしてメタデータ、ドキュメント、TypeScript 定義の完成です。
私たちは実装前に各フェーズを計画し、その後は主にクロードに作業を進めてもらいました。それは驚くほどミスが少なく、さらに印象的だったのは、当初は機械の近代化のように見えたものの、見落としがちだったかなり微妙なエッジケースをいくつか特定したことです。私は介入をほとんどしませんでした。
重要な要素の 1 つは、Yadda がすでに包括的なテスト スイートを持っていたことです。また、クロードに製品コードと対応するテストを同じステップで変更するよう依頼することも意図的に避けました。エージェントが両方を同時に変更すると、実装と同時に「正しい」の定義を自由に変更できるため、グリーン テスト スイートの証拠は弱くなります。これらの変更を分離しておくことで、クロードにはより強力な外部制約が与えられました。
作業を開始してからパッケージを公開するまでの経過時間はおよそ 1 日で、他のことも並行して行っていました。
今年の初めに、私はなぜバイブコーディングの経験がこれほど二極化しているのかを問う実験について書きました。私の結論は、結果はエージェントの使用方法に大きく依存するということでした。厳しく制約され、監視されたクロードは、非常に優れた結果をすぐに生み出すことができます。放っておくとアーキテクチャの漂流に向かう傾向があり、

不要なコードと運用上の負債。
それはわずか 7 か月前のことですが、機能は大幅に進歩しました。それでも、クロードがほとんど介入せずにこのコードを記述できるようになったと言うのは、何が変化しているのかをほんの表面に触れるにすぎません。
コーディングがボトルネックではなくなりました
これがどのような方向に向かうのかを理解するには、1 人の開発者が 1 人のコーディング エージェントと会話することについて考えるのをやめ、代わりに複数のエージェントが並行して動作していると考えると役立ちます。
これを行う方法はすでにいくつかあります。複数のクロード コード セッションを実行するだけです。 Git ワークツリーを使用すると、各エージェントが分離された作業コピーに対して作業できるようになります。 cmux などのツールを使用すると、クロード セッションのコレクションの実行がより管理しやすくなります。一方、クロード コード エージェント ビューは、複数のセッションが何を行っているか、どのセッションに注意が必要かを確認する別の方法を提供します。
これらすべてを使用すると、逐次的に作業するよりも大幅に高速にビルドできますが、すぐに別の限界に達してしまいました。それは、並列作業を管理する私自身の能力です。私は、一度に 3 つのタスク、場合によっては 4 つまたは 5 つのタスクを快適に進めることができます。それを超えると、各エージェントが何をしているのか、どのような決定が下されたのか、どのタスクが私を待っているのか、次に何を確認する必要があるのか​​というコンテキストがわからなくなり始めます。
この時点では、モデルもマシンも過負荷になっていません。ボトルネックとなるのは、作業を調整する人間です。私は、優れたオーケストレーションが次の重要な層であると確信しています。
その結論に達したのは私だけではありません。私の同僚の Marco は、この進歩をほぼ正確に My AI Engineering Journey で説明しています。オートコンプリートとしての AI から、監視された信頼できるエージェントを経て、認知負荷が制約となる並列エージェントに移行するというものです。彼は私よりもこの旅をさらに進んでおり、オークのオットーを構築することでそれに応えました。

実装、レビュー、フィードバック、ドキュメントを調整するエージェント パイプラインに進む前に、Claude Code とワークツリーに関する UI を説明します。
さらに大きな点は、AI 支援ソフトウェア開発が依然として異常なスピードで進んでいることです。個々のコーディング能力は劇的に向上し、並列実行はすでに実用化されていますが、次の制約はそのすべての能力の調整がますます厳しくなっています。そのためのツールとアプローチも同様に急速に開発されており、現在ではおそらくモデルの更新よりもさらに重要になっています。
そこでヤッダの話に戻ります。
私は、いくつかの理由から BDD が価値があると常に考えてきました。
まず、要件を通常の言語で記述すると、ドメインを明確に表現する必要があり、さらに重要なことに、一貫して表現するよう促されます。実装を作成する前にこれらの仕様を作成すると、そのドメイン言語がコードベースを通じて伝播する傾向があります。同じ概念がクラス名、関数名、API 定義、データベース スキーマ、CSS クラス、ユーザー インターフェイスにも現れ始めます。これにより、コードベースに一貫性が与えられますが、遡って達成するのは驚くほど困難です。
第 2 に、実行可能な仕様は、従来のプログラムによるテストよりもはるかにアクセスしやすくなります。プロダクト マネージャー、アナリスト、またはドメインの専門家は、次のことを理解する可能性が現実的にあります。
スティーブがコンピューター サイエンスの勉強を志願したとき
その後、大学は彼の申請を拒否しました
フィクスチャ、モック、ビルダー、アサーションを含む Jest テストから同じ意味を抽出する可能性ははるかに低くなります。
第三に、BDD は機能テストに便利な抽象化レイヤーを提供します。仕様では意図について説明しますが、ステップの実装ではセレクター、ナビゲーション、ブラウザー操作などの仕組みを扱います。これにより、

Page Object パターンと同じ利点があります。ユーザー インターフェイスへの変更は、多くの場合、何百ものテストを通じて漏洩するのではなく、抽象化の内部に吸収されます。
ただし、常にコストがかかりました。 BDD テストの最初の書き込みには時間がかかります。言語について考え、再利用可能なステップを作成し、英語を装った手続き型スクリプトを作成する誘惑に抵抗する必要があります。その成果は、より優れたドメイン モデリング、より優れたコミュニケーション、より保守しやすい機能テストを通じて後で得られます。この見返りの繰り延べにより、BDD を正当化するのは常に困難になってきましたが、私は AI が経済学を変えると考えています。
実行可能仕様はエージェントにとって非常に良いコンテキストです
ますます現実的になってきているエンジニアリング ワークフローを考えてみましょう。
会議は自動的に文字に起こされ、GitHub ディスカッションとして保存されます。これらの議論は分析され、プロジェクト Wiki の更新に使用されます。 Wiki では要件と問題が抽出されます。これらの問題は、コーディング エージェントの集合によってピックアップ、実装、レビュー、調整されます。
Wiki は、システムが何をすべきだと誰かが考えたかを知ることができます。システムがかつて何をしていたのかを知ることができます。エージェントがシステムが何をすべきであると推測したかを知ることもできます。それ自体では、システムが実際にそれを行うかどうかを判断することはできません。実行可能な仕様は可能です。そのため、エージェント開発環境における BDD は以前よりもはるかに興味深いものになります。
BDD の費用がかかる部分は、仕様の作成と維持でした。 AI はその作業の多くを安価にします。人間はすべてを入力するのではなく、言語と動作が正しいかどうかに集中することで、トランスクリプト、ディスカッション、または要件をほぼ簡単に候補仕様に変換できます。一度受け入れられると、その仕様は単なる文書化されます。

契約となります。
実装エージェントはこれを使用して、必要な動作を理解できます。テスト エージェントは、これを使用して、検証が必要なものを判断できます。レビュー担当者はこれを使用して実装に異議を唱えることができます。 CI は継続的に検証できます。これは実行可能であるため、Wiki ページでは決してできない方法でソフトウェアの動作と結合したままになります。
ここには興味深い逆転があります。 BDD は、ソフトウェア仕様を人間にとってより便利なものにすることを目的として作成されましたが、ソフトウェアの多くがマシンによって作成される場合、実行可能仕様はさらに価値があることが判明する可能性があります。自然言語はエージェントに豊富なドメイン コンテキストを提供しますが、実行可能なステップにより、仕様がシステムの動作に基づいたままであることが保証されます。
もう 1 つの変更 (Yadda v3.1.0 で追加) は、GitHub 風味の Markdown として機能仕様を記述するためのサポートです。これにより、リポジトリ内でファイルが読みやすくなり、さらに重要なことに、人間やエージェントがシステムを理解するために使用するプロジェクト Wiki やその他の重要な知識成果物と自然に共存できるようになります。同じ仕様を次のように記述できるようになりました。
# 特集: 大学への出願
## シナリオ: 申請者がエントリー要件を満たしていない
- ブーベ島大学はコンピューター サイエンスの学位コースを提供しています
- ABB のエントリー要件
- スティーブは A レベル卒業生です
- 物理学の D を持つ
- 数学の D 取得
- スティーブがブーベ島大学でコンピューター サイエンスを学ぶことを申請したとき
- 彼らは彼の申請を拒否しました
これは実行可能な仕様のままですが、GitHub で表示すると、プロジェクトの残りのドキュメントと見た目も動作もよく似ています。
Yadda 3 は npm で入手でき、ソース、ドキュメント、サンプルは GitHub にあります。

## Original Extract

Yadda 3.0.0 is out. The release modernises the JavaScript BDD library, but more interestingly, it was largely built by Claude Code and points to why executable specifications may become even more valuable in an agentic development world.

Signal Over Noise
Home
Yadda 3.0.0: BDD in the Age of AI Agents
I’ve just published Yadda 3.0.0 to npm.
For anyone unfamiliar with it, Yadda is a BDD library for JavaScript. Like Cucumber, it maps ordinary language specifications to executable code, but it was designed from the ground up to be much less prescriptive about how those specifications are written.
That means that instead of writing something like:
Given a university, The University of Bouvet Island
And The University of Bouvet Island offers a degree course in Computer Science with entry requirements of ABB
And an A-Level graduate, Steve
And Steve has a D in Physics
And Steve has a D in Maths
When Steve applies to study Computer Science at The University of Bouvet Island
Then The University of Bouvet Island rejects the application
you can write:
The University of Bouvet Island offers a degree course in Computer Science
The entry requirements for which are ABB
Steve is an A-Level graduate
With a D in Physics
And a D in Maths
When Steve applies to study Computer Science at The University of Bouvet Island
They reject his application
Both are executable specifications. I find the second considerably easier to read.
Most of Yadda 3.0 is a modernisation exercise.
Yadda has been around for a long time, and the repository had accumulated integrations and tooling for parts of the JavaScript ecosystem that are now themselves historical curiosities. Yadda 3 is Node-only, removes browser bundling and obsolete integrations such as CasperJS, PhantomJS, Bower and Component, moves the test suite to node:test , adopts Biome and lefthook, modernises the source to ES6 syntax, and adds current examples including Playwright and Puppeteer. It also now ships TypeScript definitions.
All useful, but not especially interesting to write about. There are two things about the release that I think are much more significant.
I modernised Yadda using Claude Code with Opus 4.8.
The Yadda 3.0 epic , which was itself written by Claude, broke the work into a series of deliberately separated phases: remove obsolete functionality, update the toolchain, perform mechanical formatting separately from behavioural changes, modernise the source, explore API changes, update examples and CI, then finish the metadata, documentation and TypeScript definitions.
We planned each phase before implementing it, and then I largely let Claude get on with the work. It made remarkably few mistakes and, more impressively, identified some fairly subtle edge cases that would have been easy to miss during what initially looked like a mechanical modernisation. I made very few interventions.
One important factor was that Yadda already had a comprehensive test suite. I also deliberately avoided asking Claude to modify production code and the corresponding tests in the same step. If an agent changes both simultaneously, a green test suite becomes weaker evidence because it is free to change the definition of “correct” at the same time as the implementation. Keeping those changes separate gave Claude a much firmer external constraint.
From starting the work to having the package published was roughly a day of elapsed time, and I was doing other things in parallel.
At the beginning of this year I wrote about an experiment asking why experiences of vibe coding were so polarised . My conclusion then was that the results depended enormously on how the agent was used. A tightly constrained and supervised Claude could produce extremely good results very quickly. Left to its own devices, it tended towards architectural drift, unnecessary code and operational debt.
That was only seven months ago, and the capability has moved on enormously. Even so, saying that Claude can now write this code with very little intervention barely scratches the surface of what is changing.
Coding is no longer the bottleneck
To appreciate where this is going, it helps to stop thinking about a single developer having a conversation with a single coding agent and instead consider several agents working in parallel.
There are already several ways to do this. You can simply run multiple Claude Code sessions. Git worktrees let each agent work against an isolated working copy. Tools such as cmux make running a collection of Claude sessions more manageable, while Claude Code Agent View provides another way of seeing what multiple sessions are doing and which ones need attention.
All of these let you build significantly faster than working serially, but I fairly quickly hit another limit: my own ability to manage the parallel work. I can comfortably keep three tasks moving at once, and sometimes four or five. Beyond that, I start losing the context of what each agent is doing, which decisions have been made, which task is waiting for me and what I need to review next.
At that point, the model is not overloaded and the machine is not overloaded. The bottleneck is the human coordinating the work. I’ve become convinced that good orchestration is the next important layer.
I’m not alone in reaching that conclusion. My colleague Marco describes almost exactly this progression in My AI Engineering Journey , moving from AI as autocomplete, through supervised and trusted agents, to parallel agents where cognitive load becomes the constraint. He is further along this journey than I am, and has responded by building Otto, an orchestration UI around Claude Code and worktrees, before moving on to agent pipelines that coordinate implementation, review, feedback and documentation.
The larger point is that AI-assisted software development is still moving extraordinarily quickly. Individual coding capability has improved dramatically, parallel execution is already practical, and the next constraint is increasingly the coordination of all that capability. The tools and approaches for doing so are developing just as quickly, and are now arguably even more important than the model updates.
Which brings me back to Yadda.
I’ve always thought BDD was valuable for several reasons.
Firstly, writing requirements in ordinary language forces you to articulate the domain and, more importantly, encourages you to articulate it consistently. If you write those specifications before writing the implementation, that domain language has a habit of propagating through the codebase. The same concepts start appearing in class and function names, API definitions, database schemas, CSS classes and user interfaces. That gives the codebase a coherence that is surprisingly difficult to achieve retrospectively.
Secondly, executable specifications are far more accessible than conventional programmatic tests. A product manager, analyst or domain expert has a realistic chance of understanding:
When Steve applies to study Computer Science
Then the university rejects his application
They are much less likely to extract the same meaning from a Jest test containing fixtures, mocks, builders and assertions.
Thirdly, BDD provides a useful abstraction layer for functional tests. The specification describes intent while the step implementation deals with mechanics such as selectors, navigation and browser interaction. This provides some of the same benefits as the Page Object pattern : changes to the user interface can often be absorbed inside the abstraction instead of leaking through hundreds of tests.
There has always been a cost, though. BDD tests take longer to write initially. You need to think about the language, create reusable steps, and resist the temptation to write procedural scripts disguised as English. The payoff comes later, through better domain modelling, better communication and more maintainable functional tests. That deferred payoff has always made BDD harder to justify, but I think AI changes the economics.
Executable specifications are very good context for agents
Consider an engineering workflow that is becoming increasingly plausible.
Meetings are automatically transcribed and stored as GitHub discussions. Those discussions are analysed and used to update a project wiki. The wiki is mined for requirements and issues. Those issues are then picked up, implemented, reviewed and coordinated by a collection of coding agents.
A wiki can tell you what somebody thought the system should do. It can tell you what the system used to do. It can even tell you what an agent inferred that the system ought to do. It cannot, by itself, tell you whether the system actually does it. An executable specification can. That makes BDD much more interesting in an agentic development environment than it was before.
The expensive part of BDD was producing and maintaining the specification. AI makes much of that work cheap. A transcript, discussion or requirement can be transformed into a candidate specification almost trivially, with a human concentrating on whether the language and behaviour are correct rather than typing it all out. Once accepted, that specification becomes more than documentation. It becomes a contract.
An implementation agent can use it to understand the required behaviour. A testing agent can use it to determine what needs validating. A reviewing agent can use it to challenge an implementation. CI can continuously verify it. Because it is executable, it remains coupled to the behaviour of the software in a way that a wiki page never can.
There is an interesting inversion here. BDD was created partly to make software specifications more useful to humans, but executable specifications may turn out to be even more valuable when much of the software is being written by machines. The natural language gives agents rich domain context, while the executable steps ensure that the specification remains grounded in the behaviour of the system.
One other change (added in Yadda v3.1.0) is support for writing feature specifications as GitHub-flavoured Markdown. This makes them easier to read in the repository and, more importantly, allows them to live naturally alongside the project wiki and the other key knowledge artefacts that humans and agents use to understand the system. The same specification can now be written as:
# Feature: University applications
## Scenario: Applicant does not meet the entry requirements
- The University of Bouvet Island offers a degree course in Computer Science
- The entry requirements for which are ABB
- Steve is an A-Level graduate
- With a D in Physics
- And a D in Maths
- When Steve applies to study Computer Science at The University of Bouvet Island
- They reject his application
It remains an executable specification, but when viewed on GitHub it looks and behaves much more like the rest of the project’s documentation.
Yadda 3 is available on npm , and the source, documentation and examples are on GitHub .

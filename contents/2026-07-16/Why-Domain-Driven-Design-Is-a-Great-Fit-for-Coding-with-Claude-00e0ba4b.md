---
source: "https://menelaos.vergis.net/posts/Why-Domain-Driven-Design-Is-a-Great-Fit-for-Coding-with-Claude"
hn_url: "https://news.ycombinator.com/item?id=48937132"
title: "Why Domain-Driven Design Is a Great Fit for Coding with Claude"
article_title: "Menelaos Vergis"
author: "melenaos"
captured_at: "2026-07-16T17:04:03Z"
capture_tool: "hn-digest"
hn_id: 48937132
score: 1
comments: 0
posted_at: "2026-07-16T17:01:18Z"
tags:
  - hacker-news
  - translated
---

# Why Domain-Driven Design Is a Great Fit for Coding with Claude

- HN: [48937132](https://news.ycombinator.com/item?id=48937132)
- Source: [menelaos.vergis.net](https://menelaos.vergis.net/posts/Why-Domain-Driven-Design-Is-a-Great-Fit-for-Coding-with-Claude)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T17:01:18Z

## Translation

タイトル: クロードとのコーディングにドメイン駆動設計が最適な理由
記事タイトル: メネラオス・ベルギス
説明: 私が取り組んでいること

記事本文:
メネラオス・ベルギス こんにちは。私はメネラオスです
好奇心旺盛。慎重なビルダー。
私は人々を助けるものを作り、長く残るアイデアに時間を費やしています。
ドメイン駆動設計がクロードによるコーディングに最適な理由
LLM が与えられたコンテキストと同程度に優れている場合、最良のコンテキストを生成する規律が勝ちます。その規律はドメイン駆動設計です
LLM の出力は、受信するコンテキストの品質によって制限されます。 DDD は、明示的で名前付きの明確なコンテキストを作成するための規律です。
共有語彙、明確な境界、声に出して述べられる不変条件など、DDD によって作成される成果物は、まさに AI コーディング アシスタントが求めているものです。人間のチームメイトのために作成する用語集は、優れたシステム プロンプトです。
落とし穴: LLM を独自のデフォルトのままにすると、DDD の逆の結果が生成されます。貧血の CRUD、一般名、曖昧な境界、過剰に設計されたパターン。操縦しなければなりません。
DDD の安価な部分 (言語 + 動作が豊富なモデル) をあらゆる場所で使用します。高価な部品は、ドメインが収益を上げた場合にのみ採用します。
DDD とは実際には何ですか (60 秒バージョン)
ドメイン駆動設計は、フレームワークやフォルダー レイアウトではありません。それは賭けです。複雑なソフトウェアにおいて難しいのは、テクノロジーではなくドメインを理解することです。そこで、ドメインのモデル、つまり意図的に簡素化された専用のビューを中心に置きます。それを開発者とドメインの専門家が共有する言語で表現し、そのモデルを直接反映するコードを作成します。
半分が 2 つあります。戦略的デザインは大局的なものです。つまり、共有語彙 (ユビキタス言語) と、モデルとその言語が一貫性を保つ境界である境界コンテキストです。戦術的設計は、境界内の構成要素 (エンティティ、値オブジェクト、集計、ドメイン イベント) です。
注意すべき落とし穴が 1 つあり、

そして、.NET 開発者が最も注目するのは、貧血モデルです。これはパブリックのゲッターとセッターをひとまとめにしたオブジェクトであり、実際のルールはすべて別の「サービス」クラスに追放されています。 DDD の本能は逆方向に動作します。つまり、データを所有するオブジェクトに動作をプッシュするため、そのルールが存在する場所で保護されます。
それが全体の基礎です。 AI との相性が非常に優れている理由を説明します。
DDD が LLM 支援開発に適している理由
LLM の出力品質は、LLM に与えられたコンテキストの品質によって制限されます。 DDD はまさにそのようなコンテキストを製造するための学問です。すべての DDD アーティファクトはプロンプト マテリアルとしても機能します。
これを具体化するための簡単な例であり、全体で再利用します。ライブイベントのチケット販売アプリを想像してみてください。チームは正確な言葉で一致している。ホールドとは、一部の議席を一時的に無給で獲得することである。予約とは、支払いが完了すると保留状態になります。未払いの保留は数分後に期限切れになり、座席は販売に戻ります。誰もが同じように使う 3 つのわかりやすい言葉。以下のすべてのポイントはそれらに依存するため、これらを念頭に置いてください。
これが具体的に現れるのは次の 4 つです。
1. ユビキタス言語 ↔ プロンプトコンテキスト。 AI コードの品質を左右する最大の要因は、明確な語彙です。チームがこれらの単語を修正したため、モデルは 予約.Update() を推測する代わりに SeatHold.Expire() を生成します。 「期限切れ」という言葉は「キャンセル」ではなく意図的に選ばれたため、人間でも機械でも、どの動詞が何を意味するかを推測する必要はありません。あなたがチームのために作成する用語集は、AI が必要とする用語集と同じです。とにかく人間のために書くつもりだったので、今では 2 倍の料金がかかります。
2. 境界のあるコンテキスト ↔ セッションのスコープ設定。システム全体をプロンプトにダンプすると、LLM が劣化します。注意は薄く広がり、境界は曖昧になります。境界のあるコンテキストは自然な作業単位です。「私たちは Ticketi の中にいます」

今日のコンテキストはNGです。ここにその言語とルールがあります。これは、1 つの集中セッションまたは 1 つのサブエージェントに明確にマッピングされており、これがまさにモデルを目標通りに保つ方法です。
3. ビヘイビアリッチなモデル ↔ AI が黙って破壊できないコード。オブジェクト内に不変条件が存在する場合 ( SeatHold.confirm() で「アクティブである必要があり、期限切れではない」がチェックされます)、保護された出入り口が 1 つあります。サービス全体に分散している場合 (貧血スタイル)、生成された行が到達して Status = "confirmed" に設定され、すべてのルールがスキップされる可能性があります。ルールを 1 か所に集中すると、AI (またはチームメイト) が静かに物事を破ることができる領域が縮小します。
4. 仕様通りのテスト ↔ 実行可能な契約。 DDD の不変条件は、「期限切れの保留を確認できない」というテストに直接変換されます。最初にアシスタントにこれらのテストを作成してもらい、コードを実行するための「正しい」という実行可能定義をアシスタントに与えることになります。これは散文だけよりもはるかに信頼性が高くなります。
スルーライン: DDD はドメインを明示的にすることを強制します。明示的とはまさに言語モデルが作用できるものです。
クロード+DDDに潜む罠
ここが正直なところです。 LLM のデフォルトの動作は、優れた DDD の敵です。トレーニング データには、DDD が警告するパターンが満載されているからです。始める前に次のことを知っておいてください。
デフォルトでは貧血 CRUD になります。 「SeatHold クラス」をコールドで要求すると、多くの場合、パブリックのゲッター/セッターに加えて、ロジックを保持する SeatHoldService が返されます。まさにアンチパターンです。オブジェクトおよびプライベート セッターの動作を明示的に要求する必要があります。
それは語彙を発明します。 「保留」と言うと、 予約 と書かれます。それはもっともらしく、コンパイルされ、コードとチームが異なる言語を話すようになりました。この漂流は静かで複雑なものです。
それは境界を越えて届きます。スコープを設定しないままにすると、モデルはチケット発行を支払いに直接接続します。

描いた継ぎ目を感じないので、神のオブジェクトを成長させます。
それはオーバーエンジニアリングです。 3 つのテーブルの管理パネルの集約、リポジトリ、ドメイン イベント バス。パターンはトレーニング セットに含まれているため、ドメインがそれを保証するかどうかに関係なく、パターンが適用されます。 「すべきではない場合」のない DDD は単なる儀式です。
もっともらしいが間違った不変式が生成されます。保留は 10 分後に期限切れになると言います。生成されたコードには 15 と表示されます。この数値は妥当であるように見え、コードはコンパイルされ、何も問題はありません。これは、リストの中で最も危険な罠です。完全に正しい構文の中に間違ったルールが隠れているため、コードが実行されるかどうかだけをチェックするレビューでは、この罠を通り過ぎてしまいます。防御策は DDD 自身の習慣です。用語集とテストで不変条件を特定するため、「10 分」はコードがチェックされるものであり、モデルが推測するものではありません。
コンテキストが古くなります。長いセッションを通じて、モデルは合意された言語から逸脱していきます。再研磨してください。
これらはいずれも、DDD に Claude を使用しない理由にはなりません。それらは方向転換する理由であり、それがワークフローの目的です。
LLM 駆動の DDD の実践的なワークフロー
ユビキタス言語を CLAUDE.md (またはアシスタントの指示ファイル) に入れます。それぞれの用語、その正確な意味、および「混同しないでください」をリストします。現在、言語はセッションごとにコンテキストとして自動的に読み込まれるため、毎回再説明する必要はなく、用語集が AI のデフォルトの語彙になります。
境界のあるコンテキストを一度に 1 つずつ作業し、整理します。各コンテキストにその言語とルール用の独自のファイル ( docs/contexts/ticketing.md )、またはより良い方法として、そのコンテキストのフォルダー内に CLAUDE.md を指定すると、そこで作業すると同時に自動的にロードされます。コードベースをコンテキスト (コンテキストごとにフォルダー) ごとに構造化すると、これは無料で実行できます。セッションまたはサブジェンごとに 1 つのコンテキスト

モデルの注意を継ぎ目の位置に保ちます。
コードを記述する前にモデルを再記述します。 「何かを書く前に、SeatHold が強制しなければならない不変条件を自分の言葉でリストアップしてください。」これは逆のティーチバックです。ドリフトや誤解は、コードになる前、つまり修正が安価なときに表面化します。
豊富な行動の好みを常設の指示に焼き付けます。 CLAUDE.md のコーディング標準行、または再利用可能なスキルなど、永続的な場所に配置します。 「ルールを強制するドメイン オブジェクトは、オブジェクト上でそれらのルールを保持し、オープン セッターではなく confirm() や Expire() などの意図を明らかにするメソッドを公開する必要があります。」のようなものです。それは義務ではなく好みとしてください。これは、他のすべてに適用されるわけではなく、実際の不変条件を保護するオブジェクトに適用されます。値オブジェクト、境界の DTO、および単純な構成キャリアは、プレーン プロパティでまったく問題ありません。コードベース全体ですべてのセッターをプライベートに強制すると、1 つのアンチパターンがセレモニーと交換されるだけです。
契約どおり、まずテストします。実装前に不変テスト ( Cannot_confirm_an_expired_hold ) を作成します。後から実装のバグを探すのではなく、小さくて読みやすい仕様をレビューします。
正しさだけでなく、言語に照らしてレビューしてください。レビューでは、2 つの質問をしてください。それは機能するか、そしてユビキタス言語を話せるか? Process() という名前のメソッドは、たとえ合格したとしても、概念に名前が付けられていないことを意味するため、臭いです。
このループは有益です。よりシャープなモデルはよりシャープなプロンプトを作成し、よりシャープなプロンプトはより良いコードを作成し、読み返したコードがモデルを教えてくれます。人間、ドメイン専門家、AI はすべて 1 つの言語を読みます。
線引き: DDD のどの部分を実際に使用するか
DDD は全か無かではありません。これは初心者 (および熱心な LLM) が最もよく間違える部分です。有用なフレームは高価な割に安い

私は。
ユビキタス言語。費用は規律のみです。たとえ小さなプロジェクトであっても、純粋な利点があります。また、AI を使用すると、生成されたコードが直接改善されます。
行動豊富なオブジェクト (貧血モデルなし)。これはまさに優れたオブジェクト指向設計です。データバッグとサービスよりも単純であり、複雑ではありません。
Money や SeatNumber などの明白なものに対するいくつかの値オブジェクト。 C# ではレコードを使用するのが安く、バグのクラス全体を削除します。
高価で、ドメインが収益を上げた場合にのみ採用します。
慎重に設計された集約と一貫性境界。
複数のチームとコンテキストにわたるコンテキスト マップ。
リポジトリ、個別のアプリケーション/ドメイン層、ドメインイベントインフラストラクチャ。
決定的な要因は、プロジェクトにどれだけ時間がかかるかということではありません。それは、ルールがいかに豊かであるかということです。神からインスピレーションを得た価格設定ルールをエンコードする 2 日間の機能には、価値オブジェクトと動作豊富なモデルが必要になる場合があります。 3 か月の CRUD パネルでは、ほとんど必要ない可能性があります。アシスタントにラインがどこにあるかを伝えてください。そうしないと、大聖堂全体に庭の小屋の足場が建てられます。
LLM は、問題を最も明確に示した人に報酬を与えます。 DDD は、名前付きの言語、引かれた境界線、声高に宣言されたルールなど、まさにそのイメージを生み出すための数十年にわたる規律です。それらをペアリングすることはハックではありません。同じ入力が必要な 2 つのものです。
モデルにユビキタス言語と制限されたコンテキストを与え、データが存在する場所での動作を主張し、ルールを検証し、ドメインの実際の複雑さで線を引きます。そうすれば、クロードはあなたのドメインを推測するのをやめて、それを話し始めます。

## Original Extract

Things I work on

Menelaos Vergis hi. I’m Menelaos
Curious mind. Careful builder.
I build things that help people, and I spend time with ideas that last.
Why Domain-Driven Design Is a Great Fit for Coding with Claude
If an LLM is only as good as the context you give it, then the discipline that produces the best context wins. That discipline is Domain-Driven Design
An LLM's output is bounded by the quality of the context it receives. DDD is a discipline for producing explicit, named, unambiguous context .
The artifacts DDD makes you create, such as a shared vocabulary, clear boundaries, invariants stated out loud, are exactly what an AI coding assistant is starving for. The glossary you'd write for a human teammate is a great system prompt.
The catch: an LLM left to its own defaults produces the opposite of DDD; anemic CRUD, generic names, blurred boundaries, over-engineered patterns. You have to steer.
Use the cheap parts of DDD (language + behavior-rich models) everywhere; adopt the expensive parts only when the domain earns them.
What DDD actually is (the 60-second version)
Domain-Driven Design isn't a framework or a folder layout. It's a bet: in complex software, the hard part is understanding the domain, not the technology. So you put a model of the domain at the center: a deliberately simplified, purpose-built view. You express it in a language shared by developers and domain experts, and you make the code mirror that model directly.
It has two halves. Strategic design is the big-picture stuff: a shared vocabulary (the Ubiquitous Language ), and bounded contexts , the boundaries within which a model and its language stay consistent. Tactical design is the building blocks inside a boundary: entities, value objects, aggregates, domain events.
One pitfall to watch for, and the one .NET developers hit most: the anemic model . It's an object that's just a bag of public getters and setters, with all the real rules exiled to separate "service" classes. DDD's instinct runs the other way: push behavior into the object that owns the data, so its rules are guarded where they live.
That's the whole foundation. Now here's why it pairs so well with an AI.
Why DDD fits LLM-assisted development
An LLM's output quality is bounded by the quality of the context it's given. DDD is a discipline for manufacturing exactly that kind of context. Every DDD artifact doubles as prompt material.
A quick example to make this concrete, and one I'll reuse throughout. Imagine a live-event ticketing app. The team agrees on precise words: a hold is a temporary, unpaid grab on some seats; a booking is what that hold becomes once it has been paid for; and an unpaid hold expires after a few minutes so the seats go back on sale. Three plain words, used the same way by everyone. Keep them in mind, because every point below leans on them.
Four concrete ways this shows up:
1. Ubiquitous Language ↔ prompt context. The single biggest lever on AI code quality is unambiguous vocabulary. Because the team fixed those words, the model generates SeatHold.Expire() instead of guessing at Reservation.Update() . The word "expire" was chosen deliberately over "cancel," so nobody, human or machine, has to guess which verb means what. The glossary you write for the team is the same glossary the AI needs. You were going to write it for humans anyway, so now it pays double.
2. Bounded contexts ↔ scoping the session. LLMs degrade when you dump an entire system into the prompt; attention spreads thin and boundaries blur. A bounded context is a natural unit of work: "we're in the Ticketing context today; here's its language and its rules; ignore Payments." That maps cleanly onto one focused session or one subagent, which is exactly how you get the model to stay on-target.
3. Behavior-rich models ↔ code the AI can't silently corrupt. When invariants live inside the object ( SeatHold.Confirm() checks "must be active, must not be expired"), there's one guarded doorway. When they're scattered across services (the anemic style), any generated line can reach in and set Status = "Confirmed" , skipping every rule. Concentrating rules in one place shrinks the surface where an AI (or any teammate) can quietly break things.
4. Tests-as-spec ↔ an executable contract. DDD's invariants translate directly into tests: "cannot confirm an expired hold." Have the assistant write those tests first, and you've given it an executable definition of "correct" to code against, which is far more reliable than prose alone.
The through-line: DDD forces you to make the domain explicit, and explicit is precisely what a language model can act on.
The traps hiding in Claude + DDD
Here's the honest part. An LLM's default behavior is the enemy of good DDD, because the training data is full of the patterns DDD warns against. Know these before you start:
It defaults to anemic CRUD. Ask for "a SeatHold class" cold and you'll often get public getters/setters plus a SeatHoldService holding the logic, the exact anti-pattern. You have to explicitly ask for behavior on the object and private setters.
It invents vocabulary. You say "hold," it writes Reservation . It's plausible, it compiles, and now your code and your team speak different languages. This drift is silent and compounds.
It reaches across boundaries. Left unscoped, the model will happily wire Ticketing directly into Payments and grow a god-object, because it doesn't feel the seam you drew.
It over-engineers. Aggregates, repositories, domain-event buses for a three-table admin panel. The patterns are in the training set, so it applies them whether or not the domain warrants it. DDD without the "when not to" is just ceremony.
It generates plausible-but-wrong invariants. You say a hold expires after 10 minutes; the generated code says 15. The number looks reasonable, the code compiles, and nothing complains. This is the most dangerous trap on the list, because a wrong rule hides inside perfectly correct syntax , so a review that only checks whether the code runs sails right past it. The defense is DDD's own habit: pin the invariant down in the glossary and in a test, so "10 minutes" is something the code is checked against, not something the model was left to guess.
Context goes stale. Over a long session the model drifts from the agreed language. Re-ground it.
None of these are reasons not to use Claude for DDD. They're reasons to steer, which is what the workflow is for.
A practical workflow for LLM-driven DDD
Put the Ubiquitous Language in a CLAUDE.md (or your assistant's standing-instructions file). List each term, its precise meaning, and a "not to be confused with." Now the language is auto-loaded as context on every session, and the glossary becomes the AI's default vocabulary instead of something you re-explain each time.
Work one bounded context at a time, and organize it. Give each context its own file for its language and rules ( docs/contexts/ticketing.md ), or better, a CLAUDE.md inside that context's folder so it loads automatically the moment you work there. Structuring the codebase by context (a folder per context) makes this fall out for free. One context per session or subagent keeps the model's attention where the seam is.
Make it restate the model before it codes. Ask: "Before writing anything, list the invariants SeatHold must enforce, in your own words." This is teach-back in reverse: drift and misunderstandings surface before they become code, when they're cheap to fix.
Bake the behavior-rich preference into standing instructions. Put it where it persists: a coding-standards line in CLAUDE.md , or a reusable skill. Something like: "domain objects that enforce rules should keep those rules on the object and expose intent-revealing methods like Confirm() and Expire() rather than open setters." Keep it a preference, not a mandate. It applies to the objects guarding real invariants, not to everything else. Value objects, DTOs at the boundary, and simple config carriers are perfectly fine with plain properties. Forcing every setter private across the whole codebase just swaps one anti-pattern for ceremony.
Tests first, as the contract. Have it write the invariant tests ( Cannot_confirm_an_expired_hold ) before the implementation. You review the spec , which is small and readable, rather than hunting bugs in the implementation later.
Review against the language, not just correctness. In review, ask two questions: does it work, and does it speak the Ubiquitous Language? A method named Process() is a smell even if it passes, because it means a concept wasn't named.
The loop is virtuous: a sharper model makes sharper prompts, sharper prompts make better code, and the code you read back teaches you the model. Human, domain expert, and AI all reading one language.
Drawing the line: which parts of DDD to actually use
DDD is not all-or-nothing, and this is the part beginners (and eager LLMs) get wrong most often. The useful frame is cheap vs. expensive .
Ubiquitous Language. Costs only discipline. Pure upside, even on a small project. And with an AI, it directly improves generated code.
Behavior-rich objects (no anemic models). This is just good object-oriented design; it's simpler than data-bag-plus-service, not more complex.
A few value objects for the obvious things like Money and SeatNumber . Cheap in C# with records, and they delete whole classes of bugs.
Expensive, adopt only when the domain earns it:
Carefully designed aggregates and consistency boundaries.
Context maps across multiple teams and contexts.
Repositories, separate application/domain layers, domain-event infrastructure.
The deciding factor isn't how long the project takes. It's how rich the rules are . A two-day feature encoding god-inspired pricing rules may want value objects and behavior-rich models; a three-month CRUD panel may need almost none of it. Tell your assistant where the line is, or it will scaffold the whole cathedral for a garden shed.
LLMs reward whoever gives them the clearest picture of the problem. DDD is a decades-old discipline for producing exactly that picture: a named language, drawn boundaries, rules stated out loud. Pairing them isn't a hack; it's two things that want the same input.
Give the model a Ubiquitous Language and a bounded context, insist on behavior where the data lives, verify the rules, and draw the line at the domain's actual complexity. Do that, and Claude stops guessing at your domain and starts speaking it.

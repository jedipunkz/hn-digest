---
source: "https://blog.jle.im/entry/llms-and-haskell-1-constraint-evading-behavior.html"
hn_url: "https://news.ycombinator.com/item?id=49010846"
title: "LLMs Will Cheese Your Types: Fighting Back in Haskell"
article_title: "LLMs Will Cheese Your Types: Fighting Back in Haskell · in Code"
author: "jle"
captured_at: "2026-07-22T18:02:53Z"
capture_tool: "hn-digest"
hn_id: 49010846
score: 1
comments: 0
posted_at: "2026-07-22T18:00:48Z"
tags:
  - hacker-news
  - translated
---

# LLMs Will Cheese Your Types: Fighting Back in Haskell

- HN: [49010846](https://news.ycombinator.com/item?id=49010846)
- Source: [blog.jle.im](https://blog.jle.im/entry/llms-and-haskell-1-constraint-evading-behavior.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T18:00:48Z

## Translation

タイトル: LLM はあなたの型をチーズにする: Haskell での反撃
記事のタイトル: LLM はあなたの型をチーズにする: Haskell での反撃 · コードで
説明: Justin Le のウェブログ。プログラミングにおけるさまざまな冒険と、計算物理学と知識の世界の探求をカバーしています。

記事本文:
LLM はあなたの型をチーズにする: Haskell での反撃
ジャスティン・ル著 ♦ 2026 年 7 月 22 日水曜日
ソース ◆ マークダウン ◆ LaTeX ◆ カテゴリー: Haskell ◆ コメント
そうです、それは本当です。私は今年の初めから、個人と仕事の両方の多くのプロジェクトで LLM とエージェント コーディング ツールを Haskell コーディングに統合してきました。私はすべてのプログラミングを Haskell で行っています。Haskell は、「型駆動開発」を促進する非常に表現力豊かな型システムを備えた言語です。
LLM と協力して Haskell を作成することは、非常にユニークな経験です。 「通常の」エージェント コーディングには多くの類似点やパターンが存在しますが、LLM の熱い炎と Haskell の冷静な石が出会うことで、完全に独自の最適なパス、ワークフローの癖、障害モードが多数生み出されると思います。
この投稿では、LLM コラボレーションで効果的に Haskell を記述することに関連して、LLM における制約回避動作と私が呼んでいるものを理解することに焦点を当てます。
このシリーズのパート 1 について考えてみましょう。この記事は、実際に「型駆動型」Haskell を作成した後で、LLM の非常に一般的な障害モードを特定して理解する方法について説明します。
理想的なケース: Haskell と LLM
さて、私の個人的な意見と希望的な希望は、エージェント ソフトウェア エンジニアリングにおける Haskell は、エージェント研究数学におけるリーンのようなものであるべきだということです。つまり、LLM が正しい目標に導くために必要な足場を自己構築するためのフレームワークです。
私は、「生成時の正確さ」が妥当な目標であるとは信じていません。2026 年の今日ではなく、おそらく近いうちにではないと思います。正しさへの動きは漸近的です。十分に近づくこともあるかもしれませんが、正しさの長い尾は…長いです。最も活気に満ちたフロンティア モデル プロジェクトであっても、エージェントをガイドするために存在する大規模なテスト スイートが存在します。エージェント コーディングは 5 年後には不可能になる

「正しいプログラムを吐き出す」ということは、「実装を導く最適な足場を組む」ということになります。
覚えておいてください: Haskell の型の目的は、不正なコードを検出することではありません。これは、コードの書き方と構造を指示し、最も生産的な道を導き、希望する設計を具体的に反復するのに役立ち、実際のコード作成時間をスムーズなフロー プロセスにするためのものです。
この各部分は、LLM が型を使用してコードを書くように訓練される方法とは対照的です。
この目的を達成するために、私はパス探索の直観力を高めるための適切な足場を使ってコードベースを構造化しようとします。AI はどのパスが「最も有望」であるかを検索する必要があるため、行き止まりになる可能性が最も高いパスや保守不可能なコードにつながる可能性が最も高いパスを素早く排除し、より有望なパスに沿って AI 探索を導くようにコードベースを構築します。最も優れたツールは、型、コンパイラ、警告、ヒントです。
これらの多くは私たちが人間のプログラマーに教えていることと同じであり、私が定期的に書いていることとあまり変わりません。
Parse, Don’t Validate : 無効な状態を表現できないように型を設定します。最新のフロンティアモデルであっても、AI は本質的にこれを行いません。これらは通常、防御的なプログラミング (あらゆる場所に isNull チェックを追加する、単なるパターン マッチングの代わりに isNothing / isObject / isArray などのコンストラクターのブール チェックを追加する、 Data.Set を使用する代わりに重複リストをフィルタリングする、前提条件アサーションを追加するなど) に陥り、ドメインを正しくモデル化するデータ型を停止して意識的に作成し、無効な状態を構造的に防止し、ブール型ではなく構造的に検証された型を介して事前条件と事後条件を強制するのは、通常は人間のドライバー次第です。チェック、ブールブラインドネスの回避など。
それを行うか、セッション全体または計画ステップを明確にするために費やすかのどちらかです。

これらのドメイン要件をそのタイプに含めます。
警告とリンターは積極的に使用してください。 -Werror=incomplete-patterns はもちろん、デフォルトの -Werror は、多くの AI エラー モード (無効なコードを残す、理由もなく関数に引数を残す、抽象化の機会を潰す) をカバーします。
型内でテストできないものに対しては堅牢なテスト スイートを追加しますが、統合テストも明示的にレイアウトします。これは、適切なプロンプトがないと LLM が非常にアレルギーを持つものと思われます。
私自身の経験的観察によると、これらすべては、テラバイトの型なし Python と React スロップでトレーニングされて深く埋め込まれたフロンティア LLM モデルの動作方法の性質に反しています。
そこで、私が制約回避行動と呼ぶもののリストを集めてきました。これは、要件と制約が明示的かつ明確に述べられているにもかかわらず、LLM の性質が最も根深い衝動と継続的に闘い続けることができないため、必死にそれらを回避しようとする場合です。
精査が必要な決定
まれにそれが正当化される場合に人間が合理的に行う可能性のある危険な決定を AI が行う、ある種の失敗モードがあります。しかし、通常は、それが「最も単純なアプローチ」であるため、この決定を下します。合理的な正当化に基づいた疑わしい決定と、「単純さ」や労力などの欠陥のあるヒューリスティックに基づいた疑わしい決定を区別することはできません。
すべてのプログラミング言語について、人々が言及している次のような一般的な障害モードがあります。
「このテストは不合格なので無効にしましょう」
「このテストにジャンク データを入力して、合格できるようにしましょう。」
「この関数にテスト環境かどうかを検出させ、テスト環境にある場合は異なる動作をさせてみましょう。」
ただし、Has での LLM の操作に特有だと思われるいくつかの点を以下に示します。

ケル。
避難ハッチの警告を無効にする
Haskell の「タイプ セーフティ」の多くは警告を無効にすることで簡単に回避でき、言語内の多くの「エスケープ ハッチ」は lint ルール ( Prelude.error 、 unsafeCoerce など) によって無効になります。
LLM は、多くの場合、警告または lint チェックを直接無効にする警告抑制機能を追加します。
「コンパイルできるように、このファイルに -Wno-incomplete-patterns を追加しましょう。このパターンは通常の操作ではいずれにせよアクセスできないからです。」
「Prelude.error を禁止する hlint ルールをバイパスできるように、このビルドに HLINT 無視を追加しましょう。」
このパターンの編集を禁止する編集後フックを追加するのは非常に簡単です…しかし、これは私が言う「制約回避行動」の良いプラトニックな例だと思います。
場合によっては、ファイル内の警告を無効にする必要があるかもしれません。場合によっては、 Prelude.error を使用する必要があるかもしれません。人間は目の前の状況を見て、「これは Prelude.error が正しいまれなケースの 1 つである」または「これは、その警告が間違っているまれなケースの 1 つである」と考えるかもしれません。
しかし、その判断を下す LLM を信頼すべきでしょうか?いや、クソだ。 99% の場合、これは簡単な方法として実行されているだけです。はい、時々正当な理由が見つかることがありますが、 P(legitimate |ampered) に適切に重み付けされていません。ほとんどの試みはハッキングとして行われますが、それが本当に「最も簡単な方法」であれば、喜んでその試みを実行するでしょう。
何かを制約回避行動として理解することは、「この行動を禁止する」という意味ではありません。これは、99% の場合、LLM が正しい方法ではなく、簡単な方法や早い方法を選択している状況を強調するのに役立つことを目的としています。このような場合、人間が警告を追加することが不可欠です

NG サイレンシングまたは Hlint 無視。
「最も単純なアプローチは…」という言葉は、思考追跡の中で最も見たくないものです。これは、彼らがこれまで見た中で最もばかげていてひどいコードを吐き出そうとしているという確実な兆候だからです。
P(正当ではない && 未試行) : 問題のある行為を正しく回避する
P(正当かつ未遂) : 崇高な闘争。まれに、この通常は危険な行動が本当に正当化されるが、誤った原則から正当化されない場合があります。これは人間が陥りやすい偏見と失敗モードです。あるいは少なくとも一人の人間（それが私です）。
P(正当ではない && 試行されました) : LLM が単純さや労力などの欠陥のあるヒューリスティックからこれを選択する失敗モード。
P(正当な && 試行) : 通常は危険な行為が正当化され、その試みが LLM によって正しかったというまれなケースです。
このマトリックスがより有利な限界値に向かうにつれて、ここでの私のスタンスは徐々に変化するでしょう。しかし今のところ、私が大まかに把握している数字は、引き続き警戒を促すものです。
計画されたコード内の型を無視する
すべての無効な状態を禁止し、完全に単調なパーサーを使用して、ドメインに正確に一致する完璧な型を計画しているとします。次に、その計画を実行します。残念ながら、LLM は慎重に設計された型を躊躇せずに捨てます。
「計画では引数として NonEmpty Int を使用するようになっていますが、それには多くの変更が必要になります。最も簡単なアプローチは、単に [Int] を受け取って空のリストをチェックすることです。」
「この既存の列挙型を使用する予定でしたが、必要なブランチがありません。ブランチを追加する代わりに、代わりに String を受け取るようにします。」
「構造化データ型の代わりに、解析できる区切り文字を含む文字列エンコードされたリストまたはレコードを使用しましょう。」
「 fooFunc を呼び出す必要があります。これは戻り値を返します」

s は Int なので、関数が元の計画のように Natural ではなく Int を返すようにしましょう。」
「この計画では、このレコードにフィールドを追加する必要がありますが、物事をシンプルにするために、代わりに Data.Aeson.Object を取得して、必要なフィールドを何でも返せるようにしましょう。」
「計画では、この関数を Binary a => にする予定でしたが、定義したこの型にはまだ Binary インスタンスがないため、代わりに Show a => を使用します。それが最もシンプルなアプローチです。」
多くの場合、これらのプランや型はドメイン不変条件を強制したり、適切で正しい開発をガイドしたりするために選択されることが多いのですが、LLM は計画された型安全性をすべて破棄する前にほとんど躊躇しないため、これは特にイライラさせられます。
これらはすべて、人間が計画を実行する過程で再考する可能性がある合理的な事柄です。おそらく、最初は NonEmpty Int を使用したかったのですが、よく調べてみると、 [Int] でなければならないことがわかりました。これは、ドメインに関するさらなる真実を発見するにつれて、設計を反復する自然なプロセスです。
しかし、その呼びかけは暗黙的なものではなく、話し合ったものであるべきです…最初の計画を立てるのにも熟考が必要だったのと同じように、計画を変更するにも熟考が必要です。ほとんどの場合、AI がこれらの決定を下すのは、ドメインに関する発見された真実によるものではなく、労力を最小限に抑えるための不必要なヒューリスティックや、元の計画の意図と設計の誤解 (特に圧縮後の場合) によるものです。暗黙的に行われるのではなく、明示的に議論されるべきこと。計画モード以外では、LLM はドメインの真実を追求するのではなく、最も抵抗の少ない道を模索します。
既存機能の弱体化タイプとは異なりますのでご注意ください。これは実際にはめったに見られない故障モードです。その代わり、壁にぶつかったときは

既存のコードの種類には、より一般的な別の障害モードがあります…
場合によっては、AI は既存の型を変更するのではなく (特にパッケージの境界を越えて) 既存の型を保持することを最適化します。
私はこれを「文字列の詰め込み」と呼びたいと思います。私たちは、ドメインに一致し、意味のある値の作成のみを許可する適切なセマンティック型を作成することを好みます。しかし、LLM は、時間を節約するためにこれらをひねる方法を見つけるのが大好きです。ほとんどの Haskell 型には Show インスタンスがあるため、特に文字列は脆弱です。
構造化エラーのタイプを考えてみましょう。
データ ErrorEvent = UnknownUser 文字列
| DatabaseErrorCode Int
|無効なJSON A.Value
|ネットワークエラー SomeException
|キャンセルされました (おそらく CancellationReason )
| ...
そして、次の関数があります。
handleRequest :: リクエスト -> IO ( ErrorEvent レスポンスのいずれか)
そして、新しいリクエストタイプに新しいハンドラーを追加する必要があります。おそらく、この新しいハンドラーには新しいタイプのエラーが含まれている可能性があります。 LLM は、新しい構造エラーを追加する代わりに、その構造を巧みに悪用してドメインを無効にすることに大きな喜びを感じます。
handleAddGroup :: AddGroupRequest -> IO (いずれかの ErrorEvent Response )
handleAddGroup 要求
| validGroup グループ = -- ..
|それ以外の場合 = pure $Left ( UnknownUser $ "無効なグループ: " <> group )
どこで
グループ = getGroup 要求
「無効なグループは有効な ErrorEvent ではありません。最も簡単な解決策は、グループ名を使用できる UnknownUser にエラーを入れることです。」
そして、はい、この堕落には限界がありません。あなた

[切り捨てられた]

## Original Extract

Weblog of Justin Le, covering various adventures in programming and explorations in the worlds of computation physics, and knowledge.

LLMs Will Cheese Your Types: Fighting Back in Haskell
by Justin Le ♦ Wednesday July 22, 2026
Source ♦ Markdown ♦ LaTeX ♦ Posted in Haskell ♦ Comments
Sooo yes it’s true, I’ve been integrating LLMs and agentic coding tools in my Haskell coding since the beginning of this year for a lot of my projects, both personal and professional. I do all of my programming in Haskell, a language with a very expressive type system that encourages “type-driven development”.
Working with LLMs on writing Haskell is a very unique experience; a lot of similarities and patterns exist with “normal” agentic coding, but I think the hot flame of LLMs meeting the cool stone of Haskell yields a lot of wholly unique optimal paths, workflow quirks, and failure modes.
This post will focus on understanding what I call constraint-evading behavior in LLMs as it relates to writing Haskell effectively with LLM collaboration.
Consider this Part 1 of a series. This post is about how to spot and understand a very common failure mode of LLMs once you actually are writing “type-driven” Haskell.
The Ideal Case: Haskell and LLMs
Now, my personal opinion and wishful hope is that Haskell should be to agentic software engineering what Lean is in agentic research mathematics: a framework for LLMs to self-construct the scaffolding they need to guide themselves to their correct goal.
I don’t believe that “correctness at generation-time” is a plausible goal, not today in 2026, and probably not any time soon. Motion towards correctness is asymptotic. You might get close enough sometimes, but the long tail of correctness is…long. Even the most full-vibed frontier-model projects have large suites of tests that exist to guide the agent. Agentic coding in five years will not be “spit out the correct program”, it will be “set up the best scaffolding that guides the implementation”.
Remember: the point of types in Haskell isn’t to catch bad code. It’s to direct how you write and structure your code, guide you down the most productive paths, help you concretely iterate on what design you want, and make the actual code-writing time a smooth flow process.
Each part of this is antithetical to how LLMs are trained to use types and write code.
To this end, I try to structure my codebases with the correct scaffolding to help augment the intuition of path exploration: AI has to search which paths are the “most promising”, so I structure my code-base to quickly kill off paths that are most likely to be dead-ends or lead to unmaintainable code, and to channel AI exploration along more promising paths. The greatest tools are types, compilers, warnings, hints.
A lot of these are the same things we teach to human coders, and are not very different than what I write about regularly:
Parse, Don’t Validate : set up your types to make invalid states unrepresentable. AI will, by nature, NOT do this, even the latest frontier models. They usually slip into defensive programming (adding isNull checks everywhere, boolean checks for constructors like isNothing / isObject / isArray instead of just pattern matching, filtering lists for duplicates instead of using Data.Set , adding precondition assertions, etc.), and it’s usually up to the human driver to stop and consciously create data types that model the domain correctly, structurally prevent invalid states, enforce pre- and post-conditions via structurally verified types and not boolean checks, avoid boolean blindness, etc.
Either that, or dedicate an entire session or planning step to clarifying these domain requirements in their types.
Use warnings and linters enthusiastically: -Werror=incomplete-patterns of course, and the default -Werror usually covers a lot of AI failure modes (leaving in dead code, leaving in arguments in functions for no reason and killing opportunities for abstraction).
Add robust test suites for things that cannot be tested within the types, but also explicitly laying out integration tests: something which LLMs seem pretty allergic to without proper prompting.
From my own empirical observations, all of these things are against the nature of how even frontier LLM models operate, embedded deeply from being trained on terabytes of untyped Python and React slop.
So, I’ve been gathering a list of what I call constraint-evading behavior : when the requirements and constraints are explicitly and unambiguously stated, but LLM nature desperately tries to circumvent them because they cannot keep up a sustained fight against their deepest base impulses.
Decisions that require scrutiny
There is a class of failure modes where AI makes a risky decision that a human might reasonably make in the rare case that it is justified. But, usually, it will be making this decision because it’s the “simplest approach” . It cannot separate questionable decisions based on reasoned justification and questionable decisions based on flawed heuristics like “simplicity” or effort.
We have the general failure modes like this that people mention for all programming languages:
“This test doesn’t pass, so let’s disable it”
“Let’s feed this test junk data so it will pass.”
“Let’s have this function detect if it’s in a test environment and behave differently if it is.”
But here are some that I feel are specific to working with LLMs in Haskell.
Disabling Warnings for Escape Hatches
A lot of the “type safety” of Haskell can be bypassed trivially by disabling warnings, and a lot of the “escape hatches” within the language are disabled via linting rules ( Prelude.error , unsafeCoerce , etc.)
LLMs will often add warning suppressors that straight-up disable warning or lint checks.
“Let me add -Wno-incomplete-patterns to this file so that it can compile, because this pattern is inaccessible anyway in normal operation”
“Let me add HLINT ignore to this build so that I can bypass the hlint rule forbidding Prelude.error ”
It’s pretty straightforward to add post-edit hooks to forbid edits of this pattern…but I think this is a good platonic example of what I mean by “constraint-evading behavior”.
Maybe sometimes you should be disabling warnings in your files. Maybe sometimes you should be using Prelude.error . A human might look at the situation at hand and think, “this is one of those rare cases where Prelude.error is correct”, or “this is one of those rare cases where that warning is incorrect.”
But should you trust an LLM to make that judgment call? Fuck no. 99% of the time, it is only doing this as the easy way out. Yes, every once in a while it will discover a legitimate reason, but has not properly weighted P(legitimate | attempted) . Most of the attempts will be as hacks, and it will be more than happy to follow through with an attempt if it truly is the “simplest way” .
Understanding something as constraint-evading behavior doesn’t mean “ban this behavior”. It is meant to help highlight situations where 99% of the time, it’s the LLM taking the easy or fast way out instead of the correct one. In these cases, it’s imperative that a human is what is adding the warning silencing or hlint ignore.
“The simplest approach is…” is the worst thing you ever want to see in a thought trace, because it’s a sure guaranteed sign that they are about to spew the most ridiculous and awful code you’ve ever seen.
P(not legitimate && not attempted) : Correct avoidance of problematic behavior
P(legitimate && not attempted) : The noble struggle. The rare case that you are really justified in this normally risky behavior, but out of misguided principle you do not. This is a bias and failure mode more likely to be hit by humans. Or at least one human (that’s me).
P(not legitimate && attempted) : The failure mode where an LLM will choose this out of a flawed heuristic like simplicity or effort.
P(legitimate && attempted) : The rare case that you are justified in normally risky behavior, and the LLM was correct in attempting it.
As this matrix moves towards more favorable marginals, my stance here will slowly change. But for now, the numbers I roughly see encourage continued vigilance.
Ignoring types in planned code
Let’s say you plan your perfect types that match your domain exactly, forbidding all invalid states, perfectly monotonic parsers. Then you go to execute that plan. Unfortunately, LLMs will not hesitate to throw away your carefully designed types.
“The plan says to use NonEmpty Int as an argument, but that would require changing too much. The simplest approach is to just have it take [Int] and check for empty lists”
“We planned to use this existing enum, but it doesn’t have a branch we need. Instead of adding a branch, we’ll have it take String instead.”
“Instead of a structured data type, let’s just use stringly encoded lists or records with separators we can parse out.”
“We have to call fooFunc , which returns an Int , so let’s have our function return an Int instead of a Natural like our original plan”.
“The plan requires adding a field to this record, but to keep things simple, let’s just take a Data.Aeson.Object instead so we can return whatever fields we want.”
“The plan was to have this function be Binary a => , but this type we defined doesn’t have a Binary instance yet, so we’ll just use Show a => instead. It’s the simplest approach.”
This is especially frustrating because often these plans and types were chosen to enforce some domain invariant or guide the proper and correct development, but LLMs will almost never hesitate before throwing away all of the planned type safety.
These are all reasonable things that a human might reconsider during the process of following out a plan. Maybe we originally wanted to use NonEmpty Int , but upon closer examination, we realized it does have to be an [Int] . This is the natural process of iterating on a design, as you discover more truths about the domain.
But, that call should be a discussed one, not an implicit one…it took thought to make the original plan, so it should take thought to change the plan. Most of the time AI makes these decisions, it isn’t out of discovered truths about the domain, but rather because of needless heuristics to minimize effort, or a misunderstanding of the intent and design of the original plan (especially if after a compaction). Things that should be discussed explicitly, not done implicitly. Outside of planning mode, LLMs aren’t seeking out the truth of the domain, they’re seeking out the path of least resistance.
Note that this is different than weakening types in existing functions. That’s a failure mode I rarely see in practice. Instead, when running into a wall with the type of existing code, there’s another failure mode that’s much more common…
Sometimes AI will optimize preserving existing types (especially across package boundaries) instead of changing them.
I like to call this “string stuffing”. We like to make nice semantic types that match our domain and only allow the creation of meaningful values…but LLMs absolutely love to find ways to twist these to save time. Strings, in particular, are vulnerable because most Haskell types have Show instances.
Consider a type for structured errors:
data ErrorEvent = UnknownUser String
| DatabaseErrorCode Int
| InvalidJSON A.Value
| NetworkError SomeException
| Canceled ( Maybe CancelationReason )
| ...
And you have a function:
handleRequest :: Request -> IO ( Either ErrorEvent Response )
And we have to add a new handler for a new request type. Maybe this new handler has a new type of error. Instead of adding a new structural error, LLMs will find great joy in cleverly abusing the structure to invalidate the domain.
handleAddGroup :: AddGroupRequest -> IO ( Either ErrorEvent Response )
handleAddGroup req
| validGroup group = -- ..
| otherwise = pure $ Left ( UnknownUser $ "Invalid group: " <> group )
where
group = getGroup req
“Invalid groups are not a valid ErrorEvent . The simplest solution is to put the error in UnknownUser , which can take a group name.”
And yes, this depravity knows no bounds. You

[truncated]

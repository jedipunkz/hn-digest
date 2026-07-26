---
source: "https://www.alecvo.org/blog/types-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=49063256"
title: "Types with AI"
article_title: "Types with AI"
author: "Vgoose"
captured_at: "2026-07-26T23:53:42Z"
capture_tool: "hn-digest"
hn_id: 49063256
score: 3
comments: 0
posted_at: "2026-07-26T23:01:46Z"
tags:
  - hacker-news
  - translated
---

# Types with AI

- HN: [49063256](https://news.ycombinator.com/item?id=49063256)
- Source: [www.alecvo.org](https://www.alecvo.org/blog/types-with-ai/)
- Score: 3
- Comments: 0
- Posted: 2026-07-26T23:01:46Z

## Translation

タイトル: AI を使用した型

記事本文:
AIを搭載したタイプ
家
について
写真
AIを搭載したタイプ
私は活字の大ファンです。適切に実装するとシステムが構築されます
正直で、安全で、機敏で、一緒に働くのが楽しいです。彼らもまた、
コードの記述方法に信じられないほどの影響力を持っています。きちんと
これらを活用すると、エンジニアリング組織に成功の穴[1]を作り出すことができます。
残念ながら、彼らの行動を直接測定することは非常に困難です。
利点。チームや組織を説得するのは、ほとんどの場合困難な戦いです
タイプを完全に活用することに努めます。過去数か月間、AIとして
エージェントは従来の働き方に負担をかけており、業界は
LLM をより有効に活用するために自分自身を歪め、より多くのものを取得してきました。
もっと興奮して。おそらく、最終的には、より広範な範囲を正当化するのに十分なほどニーズが高まっている
そして型のより深い採用。
からの観察
ここ数ヶ月
私が職場で、そしてより広範に社内で観察したいくつかの観察
業界:
LLM はコード探索とターゲットを解決することに優れています。
問題。現在のベンチマークを見ればこれは当然です
トレーニング後に使用されるテクニック。
LLM を拡張しすぎて最終的に終了してしまうのは非常に簡単です
非生産的。道に迷ったことがあるなら、前にもここに来たことがあるでしょう
LLM を使用して作業します。それはあなたが知っていることを超えてしまったというサインです
そして、LLM を行き止まりから抜け出すことができなくなります。
LLM は保守可能なコードを書くのが苦手です [2] 。の
システムを維持するための長期的な側面はトレーニングに含まれていません。それは
これらの非機能要件を報酬シグナルとしてエンコードするのは困難です。
LLM による生産性の向上はほとんど拡大していない
個人を超えて。チームおよび組織レベルのプロジェクトは、
同じ週数と月数で測定されます。
LLM を活用したプロジェクトの多くはデモ段階まで進んでいますが、デモ段階までは進んでいません。
出て行け。 2026年の初めから、私は鋭いジュを見てきました

mpイン
職場で開発された新しい製品、機能、ツールの数。いくつか
何か月も経ち、ほとんど全員が静かに亡くなったか、
採用に至らなかった。
エンジニアを AI ネイティブ ポッドに再構成することは進まない[3]
さて。それは破壊的であり、記録的な低水準に貢献しています
士気 [4] 。
エンジニアはコードの量に追いつけない
コードレビュー。職場では、いくつかの方法を試してきました。
コードレビューを継続します。それらのほとんどは自動チェックの一種です
積極的な参加を減らす効果をもたらした要約
ゴム印の促進。
LLM で生成されたコードをレビューするのは気分が悪いです。
型には、対処および回避に役立ついくつかのプロパティがあります。
これらの問題。
タイプはシステムの無駄のない説明です。
行動。コードの量がコードレビューを圧倒する恐れがあるため、
提出されたコードの出所により、コードが無意味になる恐れがある
儀式のタイプは、実践が追いつき、適応する方法を提供します。彼らは
レビューがはるかに速くなり、焦点を構造に移します。
バインドされたコード。これらの構造はおそらく、
これらは現在のコードに影響を与えるだけでなく、
未来のコード。 LLM によってコードが安価かつ機械的に作成されるため、コードの焦点が変わります。
実装から型までをレビューすることで、重点的にレビューできるようになります。
人間の設計上の決定と意図。
仕事の人間的な側面を見直すことはずっと楽しいです。コードの前に
レビューは個人的なものであり、同僚とつながる方法でもありました。それは
学ぶ機会と専門知識を共有する場所。がありました
不文の契約: 著者は時間と労力を費やし、レビューすることによって
あなたが自分の役割を守っていた規範。 LLM コードには何も反映されなくなりました
同僚について - 読んでも何も学ぶことはありません

それは並んでいます
ライン。何か面白いことを見つけるために坂道を歩く努力
についての発言は報われません。誰のためのフィードバックですか?あなたの同僚はそうではありません
もうコードを書く必要はなく、コメントが直接の原動力ではありません
モデルの改良。そして何よりも悪いことに、それはもはや確実ではありません
同僚も自分のコードをレビューし始めています。
コードをレビューするときはいつも、型に直接ジャンプします。そうでない場合
存在する場合は、インターフェイスを理解するために関数シグネチャを調べます。これ
ジャンプできる階層マップを提供します
実装をあちこち行ったり来たりして、何かを見つけてください
興味深いか批判的か。コードに入るのが遅いことに気づきました。
元のデザインに欠陥があるかどうかは関係ありません。タイプは迅速な提供を提供します
読み取り値の一部で短絡をチェックします。
要約は別の選択肢になりますが、誤解を招くことがよくあります。
古い、または重要な詳細が省略されている。 AI によって書かれるケースが増えています
優先順位を誤ったり、無意味な詳細を強調したりすることが多く、
変更の目的が何であるかを曖昧にします。ある時点で、私たちも
代わりにプロンプトを検討しました。彼らは正確な洞察を提供しますが、
作者の意図に反して、それらは大きなプライバシー問題を引き起こします。ほとんどの人は、
プロンプトを検査されることに不快感を感じたり、検査されることを望まない。実際には、
ほとんどの作業は複数回に分けて行われるため、プロンプトを確認するのも面倒です。
曲がる。重要なことに、どちらのオプションにも致命的な欠陥が含まれています。どちらでもない
実装に忠実であることが保証されています。誤解を招く違い
常に存在します。一方、型は
実装されているため、真であることが保証されます。
タイプは特定です。過度に拡張すると、
そうなると道に迷ってしまい、非生産的な結果になってしまうことがよくあります。とても簡単です
英語での指定を省略します。公平を期すために言うと、英語は機能します

大丈夫です
多くの場合、これは LLM に些細な問題を促すための最良のアプローチです。しかし、として
問題はさらに複雑になり、散文では具体性が不十分になります。表現する
複雑な問題に必要なすべての技術的詳細を平易に説明
英語は自然なものではありません。法外に高額な量が必要になる
努力。実際には、通常、仕様が不足していることが起こります。英語
喜んでギャップを無視し、私たちが知っていることを超えて、得られるものを得ることができます。
出力を検査して監査できない状況に陥ります。
タイプはあなたを正直に保ちます。タイプにより、重要なことを調べて答える必要があります
問題を説明するだけの質問。精神的なギャップが浮き彫りになり、
複雑なアイデアは、プリミティブから階層的に構築する必要があります。こうやって
作業を行うには、バイブコーディングに対するより保守的なアプローチが必要であり、
もっと努力してください。この作業モデルにおける LLM の役割は、探索者であり、
実装者であり、設計者ではありません。私の経験上、このモデルは
実際のコードベースでは、出力がより持続可能になります。
コードは存続し続け、反復される必要があります。
型は簡単に監査可能です。最近かかりました
Anthropic での能力不足で、英語はかなりのレベルに達する可能性があることに気づきました。
読めない。 5月と6月の数週間、私はクロードの成果に気づきました
非常に簡潔で内容が濃いため、読むのが非常に困難でした。私は
これはAnthropicチームが保存するために見つけた内部ダイヤルだったと仮定します
推論について。
でもその前から、本を読んでいると目が曇ってしまいました。
LLM出力。これらは非常に繰り返しが多く、冗長で、優先順位が間違っている可能性があります。
したがって、必要に応じて、出力が型になるように強制します。それはたくさんあります
監査しやすい形式。代わりにスキャンして飛び回ることができます
テキストのブロックを順番に読みます。これは特に次の場合に役立ちます。
新しいコードをオンボーディングしています。 LLM は非常に優れています

探検するのは私より上手です。
私が彼らによく課すタスクの 1 つは、コードベースを探索することです。
次に、一時タイプ ファイルを書き出してドメインを説明します。
コードレビュー以外では、型は複雑なものに名前を付けます。
アイデアを共有し、定期的なディスカッションに参加できるようにします。
オープンコードエディター。特殊な環境における専門用語と同様、共有される
技術的な語彙を使用すると、組織は共同で語彙を増やすことができます。
コミュニケーションが可能な具体性と専門性のレベル。これ
組織の効率化に多大な二次効果が生じる可能性がある
操作できる。誰かが原因だった会議を何度も経験しました
修正できるはずの間違った前提や質問があった、または
何についてより具体的に伝えることができたかどうか答えた
彼らは追いかけていた。
私が取り組んだ形成プロジェクトには、
Chegg でフェデレーテッド GraphQL レイヤーを形成するマイクロサービス。それ以上でした
複雑で、モノリスよりも計算効率が低くなります。
それに先行していましたが、タイムゾーンをまたがる異なるチームが参加できるようになりました。
同じ製品で共同作業を行い、集合的かつ結束力を高める
改善。フェデレーション スキーマがこれを実現するバックボーンでした
可能です。これにより、システムを分割できるようになり、
作業を分割して、各チームが自分のドメイン内で所有権と自由を持てるようにする
しかし、彼らの仕事が再び調和することは保証されています。
私が現在観察しているボトルネックは、主に次の点にあるようです。
コラボレーションが必要な境界線。大幅な減速が発生している
これらの継ぎ目に沿って、個人の生産性が倍増するのを妨げています。
より大きなプロジェクトを前進させるために。
作業を分割する際にも同様の戦略が、自由と
LLM の開発を促進する機関であり、LLM の開発を抑制する機関でもあります。
組織の前進

エーション？多分。確かに少ないだろう
従業員を根こそぎにしてポッドに再編成するよりも、従業員にとって混乱を招く可能性があります。
独立した製品やグリーンフィールドプロジェクトも限られています
スピンアップする可能性があります。作業を独立した部分に分割する
調整に取り組むことは運命づけられているようだ。おそらくこれらのポッドはすべて
速く、しかし反対方向に進みます。
型は構成および拡張し、作成します
低エントロピー成長のための自然な経路。構築されたもののインデックスを提供します。
システムを構築し、その再利用可能な構成要素に名前を付け、その自然な表面を表面化します。
拡張ポイント。デフォルトでは、LLM は次のような構築があまり得意ではありません。
これ。制約のある環境でも非常に優れたパフォーマンスを発揮するように強化されています。
ポイントインタイムタスク。型を操作することで、LLM と組織は次のことが可能になります。
最も便利な方法ではなく、成長するための最適な方法を簡単に特定できます。
当然のことながら、この種の成長によりシステムの俊敏性が維持されます。最小限に抑えます
無秩序に広がり、運用上のレバレッジが蓄積される。システムの一貫性が保たれている
理解しやすく、大きな変更もシンプルかつ安全に行えます。
実際には、これらの特性に基づいて、完全に依存するものは何でしょうか。
タイプを利用すると次のようになりますか?大まかに言えば、それは人間を変えることを意味すると思います
型のレベルまでの貢献と、より多くの型の委任
型を介した LLM への実装。私がより自信を持っていることがいくつかあります
起こるのは次のとおりです:
型はソフトウェア開発の第一級の部分になります。彼らはそうでしょう
コードからスタンドアロンで存在でき、検出可能でアクセス可能であり、
現在はコードを中心としたプロセスが、型を中心に再配置されます。
コードはコモディティ化され、代替可能になり、ブラックボックス化されます。
ソフトウェア エンジニアリングは、コードを書くことよりも、より重要なものになります。
システムをモデル化し、その要件と期待されることを説明する
行動。
LLM の使用は意図的となり、探索と探索に二分化されます。
実行。
この中で

のパラダイムでは、人間の貢献が最も高いところにあります。
LLM の弱点を活用し、補完します: 意思決定、
ビジネス要件の状況を把握し、将来を予測する
要件。これは LLM を扱うためのより保守的なアプローチですが、
もっと良くすることは不可能だと思います。マーケティングが行うのと同じくらい
あなたと同じように、自律エージェントが来ると信じていますが、私にはどうやって信じているのかわかりません
それは決して可能ではありません。残りのギャップは本当にギリギリです
要件と好みの明確さ。これらはエージェントがいない限り橋渡しできません
心を読むことができる。
[1] https://blog.codinghorror.com/falling-into-the-pit-of-success/
[2] https://arxiv.org/abs/2603.24755
[3] https://www.businessinsider.com/metas-reality-labs-shifts-to-ai-native-pods-efficiency-2026-3
[4] https://www.wired.com/story/mark-zuckerberg-meta-employee-meeting-interrupt-ai/

## Original Extract

Types with AI
home
about
pics
Types with AI
I’m a big fan of types. Properly implemented, they make a system
honest, safe, agile, and delightful to work in and work with. They also
have incredible influence over how code gets written. Properly
harnessed, they can create a pit of success[1] for engineering orgs.
Unfortunately it’s incredibly hard to directly measure their
benefits. It’s almost always an uphill battle to convince teams and orgs
to commit to fully utilizing types. So over the past few months as AI
agents are straining traditional ways of working and the industry is
contorting itself to better leverage LLMs, I’ve been getting more and
more excited. Maybe, finally, the need is high enough to justify broader
and deeper adoption of types.
Observations from the
past couple months
Some observations I’ve had at work and more broadly within the
industry:
LLMs are great at code exploration and solving targeted
problems . This makes sense after seeing the current benchmarks
and techniques used in post training.
It’s very easy to overextend with LLMs and end up
unproductive. You’ve been here before if you’ve ever felt lost
working with an LLM. It’s a sign that you’ve gone beyond what you know
and are no longer able to steer the LLMs out of dead ends.
LLMs are bad at writing maintainable code[2] . The
longitudinal aspects of maintaining a system isn’t in its training. It’s
hard to encode these nonfunctional requirements as reward signals.
The productivity boost from LLMs has largely not scaled
beyond the individual . Team and organization level projects are
still measured in the same number of weeks and months.
Many LLM powered projects get to the demo phase but don’t
get out . Since the beginning of 2026, I’ve seen a sharp jump in
the number of new products, features, and tools developed at work. A few
months have passed and almost all of them have either silently died or
failed to gain adoption.
Restructuring engineers into AI native pods[3] is not going
well . It’s disruptive and contributing to record low
morale[4] .
Engineers aren’t able to keep up with the volume of code in
code review . At work, we’ve experimented with several ways to
keep code review alive. Most of them are some flavor of automated checks
and summaries that have had the effect of reducing active participation
and promoting rubber stamping.
Reviewing LLM generated code feels bad.
Types have several properties that can help address and sidestep
these problems.
Types are a lean description of the system’s
behavior. As the volume of code threatens to overwhelm code review and
the provenance of the code submitted threaten to render it a meaningless
ritual, types offer a way for the practice to keep up and adapt. They’re
much faster to review and they shift the focus onto the structures that
bound code. These structures are arguably more important than the
implementation since they not only influence the current code, but also
future code. As LLMs make code cheap and rote, shifting the focus of the
review from the implementation to the types ensures review is focused on
human design decisions and intent.
It’s much more fun to review the human aspects of work. Before, code
review felt personal and was a way to connect with your coworker. It was
an opportunity to learn and a place to share expertise. There was an
unwritten contract: the author put in time and effort and by reviewing
the code you were upholding your part. Now LLM code reflects nothing
about your coworker — there’s nothing to learn from reading it line by
line. The effort to wade into the slop to find something interesting to
remark on isn’t rewarding. Who’s the feedback for? Your coworker isn’t
writing code anymore and your comments aren’t directly driving
improvements to the models. And worst of all, it’s no longer certain
your coworkers are even reviewing their own code anymore.
Whenever I review code, I jump directly to the types. If they don’t
exist, I look at function signatures to understand the interfaces. This
offers me a hierarchical map that allows me to jump
around and dive in and out of implementation to find anything
interesting or critical. I realized that going into the code is slow and
doesn’t matter if the original design was flawed. Types offer a quick
short circuit check with a fraction of the reading.
Summaries can be another alternative but they’re often misleading,
outdated, or omit key details. Increasingly, they’re being written by AI
which often misprioritize or highlight nonsensical details that further
obfuscates what the purpose of the change is. At one point, we also
looked to prompts as an alternative. While they offer exact insight to
the author’s intent, they pose a huge privacy issue. Most people are
uncomfortable or unwilling to have their prompts inspected. In practice,
reviewing prompts is also awkward as most work gets done in multiple
turns. Critically, both options contain a fatal flaw. Neither are
guaranteed to be true to the implementation. Misleading differences
exist all the time. Types, on the other hand, are part of the
implementation so they’re guaranteed to be true.
Types are specific , I’ve observed that overextending
is often how I get lost and end up unproductive. It’s very easy to
underspecify with English. To be fair, English works just fine and is
often the best approach to prompt LLMs for trivial problems. But as
problems get more complex, prose is just not specific enough. Expressing
all the technical details required for a complex problem in plain
English is not natural. It requires a prohibitively high amount of
effort. In practice, what usually happens is we under-specify. English
happily allows us to gloss over gaps and go beyond what we know, and get
into a situation where we’re unable to inspect and audit the output.
Types keep you honest. Types force you to explore and answer critical
questions to just describe the problem. Mental gaps are highlighted,
complex ideas have to be hierarchically built from primitives. This way
of working is a more conservative approach to vibe coding and requires
more effort. The role of LLMs in this model of working is explorer and
implementer, and less so of a designer. My experience is that this model
of working is much more sustainable in real code bases where the output
code has to continue to live and be iterated on.
Types are easily auditable. It took the recent
capacity crunch at Anthropic for me to realize that English can be quite
unreadable. For a few weeks in May and June, I noticed Claude’s output
was extremely hard to read because of how terse and dense it was. I
assume this was some internal dial the Anthropic team turned up to save
on inference.
But even before this, my eyes were already glazing over when reading
LLM output. They can be quite repetitive, verbose, and misprioritized.
So whenever appropriate, I force its output to be in types. It’s a much
easier format to audit. I’m able to scan and jump around instead of
reading a block of text sequentially. This is especially helpful when
I’m onboarding to new code. LLMs are much better than me at exploring.
One of the tasks I often put them through is to explore a codebase and
then describe the domain to me by writing out a temporary type file.
Outside of code review, types give names to complex
ideas and allow them to enter regular discussions without the need for
an open code editor. Like jargon in any specialized setting, a shared
technical vocabulary allows an organization to collectively increase the
level of specificity and technicality in which it can communicate. This
can have massive second order effects in how efficiently an organization
can operate. I’ve had many meetings where the reason was because someone
had a wrong premise or question that could have been corrected or
answered if they were able to communicate more specifically about what
they were after.
A formative project I worked on involved rolling out a fleet of
microservices forming a federated GraphQL layer at Chegg. It was more
complex and less computationally efficient than the monolith that
preceded it, but allowed disparate teams across time zones to
collaborate on the same product and make collective and cohesive
improvements. The federated schema was the backbone that made this
possible. It allowed us to partition the system and
split up work so each team had ownership and freedom within their domain
but guaranteed that their work fit back together.
The bottlenecks I’m observing at the moment seem to be mostly at the
boundaries where collaboration is needed. There’s a significant slowdown
along these seams that prevent individual productivity from multiplying
to move larger projects forward.
Could a similar strategy in partitioning work create the freedom and
agency in which LLM development thrives, and also constrain it to make
forward progress for the organization? Maybe. It’d certainly be less
disruptive to employees than uprooting and reorganizing them into pods.
There’s also only so many independent products or greenfield projects
that can spin up; splitting up work into independent pieces without
addressing coordination seems doomed. It’s possible that these pods all
go fast but in opposite directions.
Types compose and extend , creating
natural paths for low entropy growth. They provide an index of the built
system, name its reusable building blocks, and surface its natural
extension points. By default, LLMs aren’t very good at building like
this. They’re reinforced to perform very well in constrained and
point-in-time tasks. Operating on types allows LLMs and organizations to
easily identify the best ways to grow, rather than the most convenient.
Unsurprisingly, this type of growth keeps the system agile. It minimizes
sprawl and accumulates operational leverage; the system remains cohesive
and easy to understand and big changes stay simple and safe.
In practice, what would leaning on these properties and fully
utilizing types look like? Broadly, I think it means shifting human
contribution up to the level of types and delegating more of the
implementation to LLMs, through types. A few things I’m more confident
will happen are:
Types would be a first class part of software development. They’d be
able to exist standalone from code and discoverable and accessible, with
processes currently centered around code re-centered around types.
Code becomes commoditized, fungible, and blackboxed.
Software engineering becomes less about writing code and more about
modeling systems and describing its requirements and expected
behavior.
LLMs usage becomes intentional and bifurcated into exploration and
execution.
In this paradigm, the human contribution is at the point of highest
leverage, and complement the weaknesses of LLMs: decision making,
contextualizing business requirements, and anticipating future
requirements. It’s a more conservative approach to working with LLMs but
I don’t think it’s possible to do better. As much as the marketing would
like you to believe that autonomous agents are coming, I don’t see how
it could ever be possible. The gaps remaining are really just about
clarity of requirements and taste. These are unbridgeable unless agents
are able to read minds.
[1] https://blog.codinghorror.com/falling-into-the-pit-of-success/
[2] https://arxiv.org/abs/2603.24755
[3] https://www.businessinsider.com/metas-reality-labs-shifts-to-ai-native-pods-efficiency-2026-3
[4] https://www.wired.com/story/mark-zuckerberg-meta-employee-meeting-interrupt-ai/

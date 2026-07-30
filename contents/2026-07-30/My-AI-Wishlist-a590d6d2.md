---
source: "https://diverging.run/checkpoints/ai-coding-wishlist/"
hn_url: "https://news.ycombinator.com/item?id=49110695"
title: "My AI Wishlist"
article_title: "My AI Wishlist"
author: "shay_ker"
captured_at: "2026-07-30T15:03:05Z"
capture_tool: "hn-digest"
hn_id: 49110695
score: 2
comments: 0
posted_at: "2026-07-30T14:38:53Z"
tags:
  - hacker-news
  - translated
---

# My AI Wishlist

- HN: [49110695](https://news.ycombinator.com/item?id=49110695)
- Source: [diverging.run](https://diverging.run/checkpoints/ai-coding-wishlist/)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T14:38:53Z

## Translation

タイトル: 私の AI 欲しいものリスト
説明: AI をもっと楽しくするものについての一般的な奇抜な意見

記事本文:
私の AI 欲しいものリスト
ホーム 私の AI 欲しいものリスト
AI コーディングのマイルストーンに関する私の投稿と比較すると、これはいつか存在するかもしれないツールとインターフェイスについてもう少し一般的なものです…それらが今日存在していればいいのにと思います。
多くの人が TLA+、Lean、Rocq などの正式な検証に興味を持っています。これは証明、暗号化、LLM トレーニング データには最適ですが、ほとんどのエンジニアにとっては役に立ちません。私が日々実際に必要としているのは、ビジネス ロジックの検証です。 「ビジネス ロジック」は複雑であることで知られているため、これはさらに困難です。たとえすべてのビジネス ロジックを Lean に変換したとしても、デバッグはおろか、それを読むことができる人はほとんどいないでしょう。
理想的な世界では、「テスト」はコード全体で定期的に評価される (そしておそらくはメトリクスに基づいて積極的に) 評価されるビジネス要件のように見えるでしょう。エグゼクティブサマリーが、エージェントによって細心の注意を払って検証された機能の膨大なリストに拡張されることを想像することもできます。ここでの秘訣は、PRD、Slack メッセージ、Google ドキュメントなどのソースに要件を結び付けることです。
ビジネス ロジックを疑似インタラクティブな方法で検証することで、ビルダーは機能が後退しているかどうかを知り、トレードオフに価値があるかどうかを評価できるようになります。
悲しいことに、「増大する複雑性を維持する」ということは機能を削減するための言い訳としてはもはや受け入れられないため、ほとんどの人はこれを過度に複雑な製品の構築に使用するでしょう。
私の願いは、乱雑なスパゲッティ製品の不協和音が私たちを再びデザイナーに戻し、シンプルさを輝かせることです。
確かに、AI が実際の設計作業を行っています。しかし、LLM は、IDEO などによって有名になった「デザイン思考」を完全に実行しているわけではありません。
「デザイン思考」の特徴は、ビジネス上の問題を製品開発に結びつけることです。実際には、これは UX と UI を理解することを意味します。
モデルが「UX」を理解するには、人間の行動を予測する必要があります。

行動（おそらく扱いやすい）。しかし、今日の UI 理解のベンチマークは、事実に基づいた質問と回答のペアであり、ある製品と別の製品の優れたデザインの理由についてのより主観的な視覚的な理解ではありません。
実際、私たちは主観的な「良いデザイン」を読みやすくするプロセスを持っています。それは、世界中で毎週行われている「デザインクリット」と呼ばれる儀式中に起こります。
研究者の中には、モデル トレーニングまたはインターフェイスを通じて、デザイン思考を AI ワークフローにエンコードしようとしている人もいます。これは非常に初期段階にありますが、この分野でデータ ファームが強化されているため、すぐに進歩が見られることが期待されます。
AI エージェントからの完璧な出力 (「AI ファースト」プログラミング言語 + ランタイム) を設計しなければならないとしたら、それはどのようなものになるでしょうか?要件は何でしょうか?
代数型による静的検証。
厳密なメモリ安全性 (オプションまたは増分)。
実行時メトリクス (JVM、BEAM など)。
読みやすい、読みやすい。 Ruby、Elixir、Crystal が良い例です。
サンドボックスとセキュリティの組み込み
もちろん、これらの要件は互いに矛盾しており (代数型は判読可能か?)、ある程度これは、AI の有無にかかわらず、誰もが使いたいと思うような、手のかかる夢物語のプログラミング言語にすぎません。
(これはアイデアの核心です。後の投稿でこれを拡張する必要があるかもしれません)
考慮すべきことの 1 つは、この「AI ネイティブ」プログラミング言語に適切なインターフェイスを設計することです (おそらく言語自体を設計する前に)。
Ruby と Rust の基本的な、おそらく欠陥のある比較を見てみましょう。
# ルビー
デフ平均 (数値)
数値の場合は nil を返します。空の？
数字。合計。 fdiv (数値.長さ)
終わり
# 錆びる
fn 平均 (数値 : & [ f64 ]) -> オプション < f64 > {
数字の場合。 is_empty () {
なしを返します。
}
Some (数値 . iter (). sum :: < f64 > () / 数値 . len () as f64

)
}
Ruby は、提供する情報が少なくなる代わりに読みやすくなっているのは明らかです。しかし、Rust コードの詳細を選択的に「折りたたんで」読みやすくし、「AI インターフェース」で必要に応じて情報を展開/折りたたむことができるとしたらどうでしょうか。
私が言いたいことの「ペーパープロトタイプ」:
// Rust ですが、詳細が「崩壊」しています。理想的には、何らかの UI が必要です
// 何かが隠されていることを示すため
fn 平均 (数値) {
数字の場合。 is_empty () {
なしを返します。
}
数字。 iter() 。合計 () / 数値。レン（）
}
// コードがコンパイルされない場合、または明確な意図がある場合
// 掘り下げたい内容 (マウス、音声、コードレビュー)、
// 「展開された」元のバージョンを表示します。
fn 平均 (数値 : & [ f64 ]) -> オプション < f64 > {
数字の場合。 is_empty () {
なしを返します。
}
Some (数値 . iter () . sum :: < f64 >() / 数値 . len () as f64 )
}
これは、プログラミング言語、インターフェイス、エージェントの間の完全な統合になります。ここにたどり着くまでには長い時間がかかると思います。しかし、私たちがそうするとき、それは至福になるでしょう。
AIの「Mavis Beacon」とは？
仕事で毎日 Claude または Codex を使用している場合、誰もがこれらの AI ツールを隅々まで知っていると考えるのは簡単です。実際には、無限トークンの魔術を実行できる人と、単に Google に代わって宿題を終わらせる必要がある人の間には、大きな隔たりがあります。過去に「デジタル変革」を行ったのと同様に、人々にプロンプ​​トの出し方を大規模に教えるための「AI 変革」が行われることになります。
AI ツールをうまく使ったことがある人なら、まったく異なる考え方が必要であることをすでにご存知でしょう。この「考え方」には名前があり、メタ認知ということになります。古い LLM であっても、LLM と対話するにはメタ認知的思考が必要であることが観察されています。
スキルを大規模に教えるにはどうすればよいでしょうか?まあ、普通のゲームでは

せ！ AI はゲームの構築に使用されますが、LLM プロンプトがコア ゲーム メカニズムとしてもっと頻繁に使用されることを望みます。
残念ながら、これが得られるまでにはしばらく時間がかかるかもしれません。ゲーム業界は AI に対して否定的な見通しを持っているようです。どれほど重要であるにもかかわらず、「AI 用 Mavis Beacon」を作成する意欲を持つゲーム デザイナーはほとんどいないのではないかと思います。
もしかしたら、もっと大きな研究所がそれを実現できるでしょうか?
コーディングエージェント間でコンテキストを共有するにはどうすればよいでしょうか?あるプロジェクトに取り組んでいるエージェントがいて、今度はそのエージェントを、同様のドメインに取り組んでいる同僚のエージェントと連携したいと考えているとします。それはどのように機能するでしょうか?
現在のソリューションでは、コンテキストを Notion/Slack/Linear などの共有ナレッジ ベースにダンプし、各クロード/コーデックス インスタンスがそれを参照として使用します。
しかし、それでは次の問題をどのように解決すればよいでしょうか。
エージェントはどのようにして他のエージェントに通知/メッセージを送信しますか?
エージェントにすべてのコンテキストを共有させたくないが、関連する場合は「プライベート」コンテキストを使用できるようにしたい場合はどうすればよいでしょうか?
同僚のエージェントに問い合わせるにはどうすればよいですか?
目を細めて見ると、これらは私たちが日々取り組んでいる人々の問題に似ているように見えます。マルチプレイヤー インターフェイスで人間とエージェントを融合する何らかのプロトコルが必要になります。
A2A はありますが、実際に使用している人を知りません。
これに取り組んでいる場合は、ぜひお知らせください。

## Original Extract

General whimsy on what would make AI more fun

My AI Wishlist
HOME My AI Wishlist
Compared to my post on AI coding milestones , this is a bit more general on tooling + interfaces that might exist someday… I just wish they existed today.
Many people are interested in formal verification, like TLA+, Lean, or Rocq. That’s great for proofs, cryptography, and llm training data, but not useful for most engineers. What I actually need, day-to-day, is verifying business logic. This is much harder because “business logic” is famously squishy. Even if you did translate all your business logic into Lean, few would be able to read it let alone debug it.
In an ideal world, “tests” would look like business requirements that are evaluated regularly across code (and perhaps proactively on metrics). You could imagine executive summaries expanded out to a laundry list of features that are meticulously validated by an agent. The trick here is to tie the requirements back to the source, like a PRD, Slack message, Google Doc, etc.
Verifying business logic in a pseudo-interactive fashion would let builders know if they’ve regressed on a feature or assess if the tradeoff is worth it.
Sadly, most would use this to build over-complicated products, since “maintaining growing complexity” would no longer be an acceptable excuse to pare down features.
My hope is that a cacophony of messy, spaghetti products jolts us back to being designers again, letting simplicity shine.
We certainly have AI doing real design work . But LLMs are not quite doing “design thinking”, made famous by IDEO and others.
What’s special about “design thinking” is tying business problems to product development. In practice, this means UX and UI understanding.
For a model to understand “UX”, that’d require predicting human behavior (possibly tractable ). But today’s benchmarks on UI understanding are fact-based question/answer pairs , not more subjective visual understanding on what makes one product well-designed vs. another.
We do, in fact, have a process for making subjective “good design” legible. It happens during a weekly ritual, all around the world, called the “Design Crit”.
Some researchers are trying to encode design thinking into AI workflows, either through model training or an interface . It’s very nascent, but expect to see progress on this soon, given that data farms are beefing up in this domain .
If we had to design the perfect output from an AI agent - an “AI-first” programming language + runtime - what would it look like? What would the requirements be?
Static verification via algebraic types.
Strict memory safety, optional or incremental.
Runtime metrics (like JVM, BEAM, etc.).
Readable, legible. Ruby, Elixir, Crystal are great examples.
Sandboxing & security built-in
Of course, these requirements are in conflict with each other (are algebraic types legible?), and to some extent this is just a hand-wavy, pipe dream wishlist programming language that anyone would love to use, with or without AI.
(This a kernel of an idea - I may need to expand this in a later post)
One thing to consider is designing the right interface for this “AI-native” programming language (perhaps before we design the language itself).
Let’s take a basic, possibly flawed comparison between Ruby & Rust:
# Ruby
def average ( numbers )
return nil if numbers . empty?
numbers . sum . fdiv ( numbers . length )
end
# Rust
fn average ( numbers : & [ f64 ]) -> Option < f64 > {
if numbers . is_empty () {
return None ;
}
Some ( numbers . iter (). sum :: < f64 > () / numbers . len () as f64 )
}
It’s obvious that Ruby is more readable at the cost of giving less information. But what if we selectively “collapsed” details in the Rust code to make it readable, and in your “AI interface” you could expand/collapse information as needed.
A “paper prototype” of what I mean:
// Rust, but details "collapsed", ideally with some UI
// to show you something is being hidden
fn average ( numbers ) {
if numbers . is_empty () {
return None ;
}
numbers . iter () . sum () / numbers . len ()
}
// if the code does not compile, or there's clear intent
// that you want to dig in (mouse, voice, code review),
// show the "expanded", original vrsion:
fn average ( numbers : & [ f64 ]) -> Option < f64 > {
if numbers . is_empty () {
return None ;
}
Some ( numbers . iter () . sum :: < f64 >() / numbers . len () as f64 )
}
This would be full integration between a programming language, the interface, and the agent. I’m guessing it’ll take a long time to get here. But when we do it’ll be bliss.
What is the “Mavis Beacon” of AI?
If you use Claude or Codex every day at work, it’s easy to assume that everyone knows these AI tools inside and out. In reality, there’s a huge gap between those who can do infinite token sorcery and those who merely replaced Google and need to finish their homework. Similar to how we did “digital transformations” in the past, we’ll have “AI transformations” to teach people how to prompt, at scale.
If you’ve used AI tools well, you already know that it requires a totally different way of thinking. It turns this this “way of thinking” has a name: it’s metacognition . Even for older LLMs, we’ve observed that interacting with LLMs demands metacognitive thinking .
How do we teach skills at scale? Well, with games of course! AI is used to build games, but I wish LLM prompting is more often used as the core game mechanic .
Unfortunately, it may be some time before we get this. It seems the game industry has a negative outlook on AI . I suspect few game designers will be motivated to create the “Mavis Beacon for AI”, despite how important it is.
Maybe one of the bigger labs can make it happen?
How might we share context between coding agents? Imagine I had an agent working on a project, and now I want it to collaborate with a colleague’s agent that was working on a similar domain. How would that work?
The current solution involves dumping context to a shared knowledge base like Notion/Slack/Linear, and then each claude/codex instance uses that as a reference.
But with that, how do we solve the following:
How does an agent notify/message another agent?
What if I don’t want my agent to share all of its context, but still be able to use its “private” context when relevant?
How can I prompt my colleague’s agent?
If you squint, these look similar to people problems that we tackle every day! We’ll need some protocol that blends humans + agents in a multiplayer interface.
There is A2A , but I don’t know of anyone really using it.
If you’re working on this, please let me know!

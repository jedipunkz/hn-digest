---
source: "https://grigio.org/bun-1-4-the-controversial-ai-driven-rewrite-from-zig-to-rust/"
hn_url: "https://news.ycombinator.com/item?id=48609168"
title: "Bun 1.4: The Controversial AI-Driven Rewrite from Zig to Rust"
article_title: "Bun 1.4: The controversial AI-Driven Rewrite from Zig to Rust"
author: "baranul"
captured_at: "2026-06-20T15:38:12Z"
capture_tool: "hn-digest"
hn_id: 48609168
score: 2
comments: 0
posted_at: "2026-06-20T13:40:50Z"
tags:
  - hacker-news
  - translated
---

# Bun 1.4: The Controversial AI-Driven Rewrite from Zig to Rust

- HN: [48609168](https://news.ycombinator.com/item?id=48609168)
- Source: [grigio.org](https://grigio.org/bun-1-4-the-controversial-ai-driven-rewrite-from-zig-to-rust/)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T13:40:50Z

## Translation

タイトル: Bun 1.4: Zig から Rust への物議を醸す AI 主導の書き換え
記事のタイトル: Bun 1.4: Zig から Rust への物議を醸す AI 主導の書き換え
説明: 2026 年 5 月、Bun チームはソフトウェア業界で長年ささやかれてきたことを実行しました。ランタイム全体を Zig から Rust に書き換えました。専任チームとともに 1 年以上かかるわけではありません。 6日以内に。 AI エージェントの使用。
約100万行のコードでは、

記事本文:
錆びる
Bun 1.4: 物議を醸す Zig から Rust への AI 主導の書き換え
2026 年 5 月、Bun チームはソフトウェア業界で長年ささやかれてきたことを実行しました。ランタイム全体を Zig から Rust に書き換えました。専任チームとともに 1 年以上かかるわけではありません。 6日以内に。 AI エージェントの使用。
これはコード行数が約 100 万行に達し、これまで公の場で試みられた AI 主導のコードベース移行としては最大規模です。そして発送されました。何が起こったのか、それが何を意味するのか、そして開発者コミュニティがそれについて議論をやめられない理由を以下に示します。
生の事実から始めましょう:
Zig の約 570,000 行が Rust の約 682,000 行になりました (AI 生成のコメントを除く)
既存のテスト スイートの 99.8% が最初の AI 生成ポートを通過
結果の Rust コード内の 13,044 個の安全でないブロック (同様のサイズの手書きの Rust プロジェクトの平均は約 73)
999 以上の static mut の使用 (グローバルな可変状態、Rust の安全性保証のアンチテーゼ)
ManuallyDrop の 262 個のインスタンス、mem::forget の 81 個のインスタンス
9,765 行ものファイル (JS パーサー、 p.rs )
GitHub 史上最大の PR は coderabbitai[bot] と claude[bot] によってレビューされました
この書き換えでは、Anthropic の Claude Code と「動的ワークフロー」、つまり数百または数千の AI エージェントにわたって作業をファンアウトするシステムが使用されました。 Anthropic の Claude Code 責任者である Boris Cherny 氏は、このアプローチについて次のように説明しています。
「ほとんどのクロード コード セッションを見ると、実際にプロンプトを実行しているのは別のクロードです。」
パイプラインは次のように機能しました。
クロード エージェントはファイルを Zig から Rust に並行して翻訳しました
各ファイルには「敵対的レビューアー」があり、最初のエージェントの出力のバグを見つけるために割り当てられた 2 番目の AI エージェントです。
自動修正のためにエージェントにループバックされるフィードバックを構築およびテストします
Bun の作成者である Jarred Sumner が高レベルの監督を提供
Bun は Anthropic が所有しています (会社は

これにより、チームはフロンティア モデルで無制限のコンピューティングにアクセスできるようになります。これは月額 20 ドルのコーディング アシスタントではなく、産業規模の AI オペレーションでした。
ジャレッド・サムナー氏は、移行の理由をいくつか挙げています。
メモリの安全性。 Bun はセグメンテーション違反で有名でした。 GitHub の問題トラッカーにはクラッシュ レポートが散乱していました。 Rust のコンパイル時保証は、Zig コードベースを悩ませていたメモリ バグのクラス全体を排除することを約束しました。
コンパイラの速度。 Bun チームは Zig のコンパイラのフォークをメンテナンスしていました (bun の zig フォークはコンパイル時間が 4 倍高速になりました) が、メンテナンスの負担になっていました。 Zig アップストリームは Bun の IO パターンを壊す変更を行っていました。 Rust のコンパイラ エコシステムは、Zig では実現できない安定性とツールを提供しました。
AIフィット。 Rust コミュニティが認めたくない不快な真実は次のとおりです。Rust のコンパイラは完璧な AI コーディング アシスタントです。 Xユーザーのrentierdigitalさんが「アンドンコード」と呼ぶものが付いています。 AI エージェントが不正な Rust を作成すると、コンパイラーは構造化エラー、正確な場所、修正するパスなどのコードを即座にヤンクします。エージェントは読み取り、修正し、再実行します。このフィードバック ループを同じ忠実度で提供するシステム言語は他にありません。
コードの実際の見た目
品質に関する議論は白熱するところです。
一方で、99.8% というテスト合格率は、AI によって完全に 1 週​​間以内に行われたものはもちろんのこと、あらゆる書き換えにおいて驚異的です。
一方で、コードは荒いです。 Reddit ユーザーが詳しく調査したところ、次のことがわかりました。
「静的 mut が 999 回使用されています。どうやら、これは私が予想していたよりもさらにひどい、グローバルに変更可能なスロパゲッティ カルボナーラの汚泥のようです。私が特におかしいと思うのは、文字通り解決策を特定し、それについてのコメントに同数のトークンを費やしていることです。」
生成されたコード内の一般的なコメントは次のとおりです。
// ポート

注: Zig はプレーンな `var` グローバル (非同期) を使用しました。ここにミラーリングされます
// 同じ単一ライターアットスタートアップ規律を持つ `static mut`;後に読む
// `enable()` はどちらの言語でも技術的に危険です。
// TODO(port): フェーズ B が厳密な健全性を必要とする場合は、AtomicBool/AtomicI32 を検討してください。
AI は、自分が不適切なコードを書いていることを認識していました。それは気にしませんでした。コメントには「フェーズB」と書かれており、チームが後でこれをクリーンアップする予定であることを示唆している。しかし批評家らは、1万3000個の安全でないブロックが存在し、コードを読む人間がいない状況では「フェーズB」は決して到来しない可能性があると指摘している。
その反応は……二極化しています。感情のサンプルは次のとおりです。
r/rust より (賛成票 718 件、コメント 436 件):
「私の理解では、ほぼ 100% AI によって行われ、人間による入力はほぼゼロで、コードは AI によってのみレビューされます。」 -- u/AlyoshaV
「もしこのようなことが結果的にうまくいったら、そしてそれがうまくいくかもしれないなら、私はしばらく計画していたメルトダウンを実際に起こすかもしれない。」 -- u/gordonnowak
r/programming より (賛成票 610 件、コメント 401 件):
「私の昔のCS教授は、人間は人間が読むためのコードを書き、たまたまコンピュータ上で実行するためだけにコードを書くと言っていた。この人はその教訓を見逃していたと思う。」 -- u/ファリサイオス (689 ポイント)
「それは工学上の奇跡か大惨事のどちらかでしょう...私は後者のほうに傾いています。」 -- u/Chaoses_Ib
具体的にはRustコミュニティから:
「『おそらく大丈夫』という理由だけで AtomicBool の代わりに静的 mut bool を使用するという考えは、Rust の通常の書き方とはまったく逆です。」 -- u/ZZaaaccc
「私はそれに触れていません」 -- u/Maybe-monad
「逸話として言えるのは、最近バンを使うのがとてもイライラして、もう使うのをやめてしまったということです。私も以前はバンのファンでした。奇妙なハングと奇妙な動作、コマンドラインのように物事が渡されない API の奇妙なギャップ、そして問題追跡にはクロード以外の何物でもありませんでした」

r と PR。" -- u/aLiamInvader
しかし、すべてがネガティブなわけではありません。これを真の画期的な進歩だと見る人もいます。
「これは多くの人が見逃していることだと思います。以前は安全でないブロックはありませんでしたが、現在は 14,000 個あります。これを深刻な問題と見る人もいますが、現実には、これらの危険なブロックが存在し、これらの概念を表現できる言語を使用したために特定せざるを得なくなったのです。」 -- u/oursland
「この制約により、システムは遅くなるのではなく、速くなります。これが AI コードの品質を左右する本当の変数です。どのモデルか、プロンプトをどれだけ注意深くプロンプトするかではなく、環境が十分な診断精度で十分な速度でコードを引っ張り、負債が蓄積する前にエージェントが自己修正できるかどうかです。」 -- @rentierdigital on X
書き換えによって一部の領域が改善されたという初期の証拠があります。 Bun の主要ユーザーである Prisma Compute は、両方のバージョンをテストしました。
「安定版ではメモリ リークが発生し、接続プーリングが不安定でしたが、Rust バージョンではメモリ使用量が減り、接続プーリングが安定しました。」 -- 日本の開発者 @zaru からの翻訳
Rust バージョンは、アップストリームの Zig コンパイラよりも高速にコンパイルします (ただし、Bun 独自のフォークされた Zig コンパイラよりは高速ではありません)。
実行速度。人間のチームであれば 1 年かかるであろう書き直しが 1 週間足らずで完了しました。たとえコードに大幅なクリーンアップが必要な場合でも、それは桁違いの改善です。
より良い岩盤。 Rust の型システムと借用チェッカーは、段階的な改善のための基盤を提供します。それぞれの安全でないブロックは、時間の経過とともに返済できるラベル付きの負債です。 Zig では、同じリスクは目に見えませんでした。
テスト規律。このプロジェクトには、言語間で変更する必要のない大規模な TypeScript テスト スイートが含まれています。 99.8% の合格率は、AI がセマンティクスを非常によく保存していることを示唆しています。
フィードバックをまとめました。 Rust のコンパイラは AI-g に独自に適しています

生成されたコード。コンパイル時に型エラー、ライフタイム違反、安全でない誤用を捕捉し、即座に修正信号を提供します。
目に見える改善。早期採用者は、Rust バージョンのメモリ動作と安定性が向上していると報告しています。 Prisma Compute のスイッチは実際の検証です。
13,000 以上の安全でないブロック。これは慣用的な Rust ではありません。これは、Rust 構文をまとった Zig セマンティクスです。借用チェッカーの保証はオプトインであり、コードベースのほとんどはオプトアウトされています。
人間によるレビューはありません。 PR は claude[bot] と coderabbitai[bot] によってレビューされました。完全な差分を読む人はいません。何十万もの開発者が依存しているランタイムにとって、これは不安なことです。
ベンダーロックイン。 Anthropic は Bun を所有しています。このコードベースを有意義に保守できる唯一のツールはクロードです。 Bun に貢献したい場合は、Anthropic のモデルにアクセスする必要があります。これはもはや伝統的な意味でのオープンソースではありません。
保守性負債。 AI によって翻訳されたコードは、元の言語の構造を保持します。ターゲット言語のイディオムは採用されません。その結果、人間が読んだり、リファクタリングしたり、拡張したりするのが難しいコードが作成されます。
QAの懸念事項。何人かのコメント投稿者は、AI が基礎となる実装を修正するのではなく、合格するように一部のテストを変更したと指摘しました。 LLM がコードとテストの両方を作成している場合、検証ループが危険にさらされます。
「その 0.2% はそんなことは気にしないでください」 -- u/codemuncher、皮肉っぽく
共同体骨折。過去の貢献者は疎外感を感じています。 Zig コミュニティはこれを裏切りだとみなしています。 Rust 開発者は、これを言語の誤用とみなしています。すでに最後の Zig バージョンのフォークを検討している人もいます。
これが業界にとって何を意味するか
この書き換えはまぐれではなく、予兆です。好きか嫌いかに関係なく、このアプローチは出荷するには十分に機能します。それが知らせる内容は次のとおりです。
言語の移行はありません

o より長い複数年にわたるプロジェクト。フロンティアモデルが 57 万行の Zig を 6 日間で合格可能な Rust に翻訳できるとしたら、言語を切り替えるコストは跳ね上がります。
ボトルネックは変化しつつあります。 Boris Cherny が指摘したように、AI が最初に解決したボトルネックはコード作成でした。次にコードレビューが行われ、「クロードのチーム」によって解決されました。保守性とセキュリティが新たな制約となります。このパターンが繰り返されます。
Rust は、AI によって生成されたシステム コードのデフォルト言語になる可能性があります。そのコンパイラのフィードバック ループは比類のないものです。ゴーのキャッチが少なくなります。 Python はコンパイル時に何もキャッチしません。 C と C++ にはセーフティ ネットがありません。 Rustの「あんどんコード」は、AIが書いたものをレビューできない場合に最も安全な選択肢になります。
信頼は未解決の質問です。コードはテストに合格します。発送します。しかし、誰もそれを読んでいません。個人的なプロジェクトや社内ツールの場合は問題ありません。実稼働環境の Node.js の代替を強化する JavaScript ランタイムにとって、これは信じがたいことです。
Bun 1.4 には、エンジニアリングの実験、Anthropic のマーケティング クーデター、AI 生成コードのストレス テスト、そしてより良いランタイムを出荷するための真の試みなど、多くのことが同時に含まれています。
この書き換えは、その支持者が主張する奇跡でも、批判者が予測する大惨事でもありません。これは現実的な賭けです。AI が生成したバージョンを今すぐ出荷し、後でクリーンアップします。そのクリーンアップが実際に行われるかどうか、そしてコミュニティがそれを見守るために固執するかどうかは、私たちが今後数か月間注目することになります。
1 つ確かなことは、私たちは閾値を超えたということです。 AI が数百万行のシステム コードベースを 1 週間以内に言語間で翻訳できるようになると、ソフトウェア エンジニアリングのルールが変わります。私たちがその変化に対応する準備ができているかどうかは、まったく別の問題です。
人工的な人工知能とピッコール専門分野の未来は何ですか?
寓話 5 は、AI がもたらす強力な効果をもたらします

禁止されるのか？
中国：人工知能と労働の覇権。矛盾？

## Original Extract

In May 2026, the Bun team did something the software industry has been whispering about for years: they rewrote their entire runtime from Zig to Rust. Not over the course of a year with a dedicated team. In six days. Using AI agents.
At nearly a million lines of code,

rust
Bun 1.4: The controversial AI-Driven Rewrite from Zig to Rust
In May 2026, the Bun team did something the software industry has been whispering about for years: they rewrote their entire runtime from Zig to Rust. Not over the course of a year with a dedicated team. In six days. Using AI agents.
At nearly a million lines of code, this is the largest AI-driven codebase migration ever attempted in public. And it shipped. Here's what happened, what it means, and why the developer community can't stop arguing about it.
Let's start with the raw facts:
~570,000 lines of Zig became ~682,000 lines of Rust (excluding AI-generated comments)
99.8% of the existing test suite passes on the first AI-generated port
13,044 unsafe blocks in the resulting Rust code (hand-written Rust projects of similar size average ~73)
999+ uses of static mut (global mutable state, the antithesis of Rust's safety guarantees)
262 instances of ManuallyDrop , 81 of mem::forget
Files as large as 9,765 lines (the JS parser, p.rs )
The largest PR in GitHub history was reviewed by coderabbitai[bot] and claude[bot]
The rewrite used Anthropic's Claude Code with "dynamic workflows" -- a system that fans out work across hundreds or thousands of AI agents. Boris Cherny, head of Claude Code at Anthropic, described the approach:
"If you look at most Claude Code sessions, it's actually another Claude that does the prompting."
The pipeline worked like this:
Claude agents translated files from Zig to Rust in parallel
Each file had an "adversarial reviewer" -- a second AI agent assigned to find bugs in the first agent's output
Build and test feedback looped back into the agents for automatic correction
Jarred Sumner, Bun's creator, provided high-level oversight
Bun is owned by Anthropic (the company behind Claude), giving the team access to unlimited compute on the frontier model. This wasn't a $20/month coding assistant -- it was an industrial-scale AI operation.
Jarred Sumner cited several reasons for the migration:
Memory safety. Bun had accumulated a reputation for segfaults. The GitHub issue tracker was littered with crash reports. Rust's compile-time guarantees promised to eliminate entire classes of memory bugs that had plagued the Zig codebase.
Compiler velocity. The Bun team had been maintaining a fork of Zig's compiler ( bun's zig fork got 4x faster compilation times ), but it had become a maintenance burden. Zig upstream was making changes that broke Bun's IO patterns. Rust's compiler ecosystem offered stability and tooling that Zig couldn't match.
AI fit. Here's the uncomfortable truth the Rust community doesn't want to admit: Rust's compiler is the perfect AI coding assistant. It has what X user rentierdigital called an "andon cord." When an AI agent writes bad Rust, the compiler yanks the cord immediately -- structured error, exact location, path to fix. The agent reads, corrects, reruns. No other systems language provides this feedback loop at the same fidelity.
What the Code Actually Looks Like
The quality debate is where things get heated.
On one hand, 99.8% test pass rate is impressive for any rewrite, let alone one done entirely by AI in under a week.
On the other hand, the code is rough . A Reddit user doing a deep dive found:
"999 uses of static mut . It would appear it's even worse a globally mutable cesspit of slopagetti carbonara than I had predicted. What I find especially crazy is that it has literally identified the solution and has spent a similar number of tokens commenting about it."
A typical comment in the generated code reads:
// PORT NOTE: Zig used plain `var` globals (unsynchronized). Mirrored here as
// `static mut` with the same single-writer-at-startup discipline; reads after
// `enable()` are technically racy in both languages.
// TODO(port): consider AtomicBool/AtomicI32 if Phase B wants strict soundness.
The AI knew it was writing unsound code. It just didn't care. The comment says "Phase B" -- suggesting the team plans to clean this up later. But critics point out that with 13,000 unsafe blocks and no humans reading the code, "Phase B" may never arrive.
The response has been...... polarized. Here's a sample of the sentiment:
From r/rust (718 upvotes, 436 comments):
"It's almost 100% AI done as I understand it, near-zero human input and code is only reviewed by AI" -- u/AlyoshaV
"if something like this ends up working, and it may well work, I might actually have that meltdown I've been planning for a while" -- u/gordonnowak
From r/programming (610 upvotes, 401 comments):
"My old CS professor used to say that we write code for humans to read and only incidentally to run on computers. I think this guy missed that lesson." -- u/Pharisaeus (689 points)
"It will either be an engineering miracle or disaster... I'm leaning toward the latter." -- u/Chaoses_Ib
From the Rust community specifically:
"The idea of using a static mut bool instead of an AtomicBool just because 'it's probably fine' is just completely antithetical to how Rust is normally written." -- u/ZZaaaccc
"I am not touching that" -- u/Maybe-monad
"I can tell you anecdotally that bun has gotten so frustrating to use recently that I've given up on it, and I used to be a fan. Weird hangs and odd behaviour, weird gaps in the API where things aren't passed through like they are on the command line, and nothing but Claude on the issue tracker and PRs." -- u/aLiamInvader
But it's not all negative. Some see this as a genuine breakthrough:
"This is what I think a lot of people miss. There were no unsafe blocks before, now there are 14,000. While some see this as a serious issue, the reality is that these unsafe blocks were there and have been forced to being identified due to using a language which can represent these concepts." -- u/oursland
"the constraint makes the system faster, not slower. this is the real variable in AI code quality: not which model, not how carefully you prompt, but whether your environment pulls the cord fast enough with enough diagnostic precision that the agent self-corrects before debt accumulates" -- @rentierdigital on X
There's early evidence the rewrite improved things in some areas. Prisma Compute, a major Bun user, tested both versions:
"Stable version had memory leaks and connection pooling was unstable, but Rust version had lower memory usage and connection pooling stabilized." -- translated from Japanese developer @zaru
The Rust version also compiles faster than the upstream Zig compiler (though not faster than Bun's own forked Zig compiler).
Speed of execution. A rewrite that would have taken a human team a year was done in under a week. Even if the code needs significant cleanup, that's an orders-of-magnitude improvement.
Better bedrock. Rust's type system and borrow checker provide a foundation for incremental improvement. Each unsafe block is a labeled debt that can be repaid over time. In Zig, those same risks were invisible.
Test discipline. The project has a massive TypeScript test suite that doesn't need to change between languages. The 99.8% pass rate suggests the AI preserved semantics remarkably well.
Compiled feedback. Rust's compiler is uniquely suited for AI-generated code. It catches type errors, lifetime violations, and unsafe misuse at compile time, providing immediate correction signals.
Tangible improvements. Early adopters report better memory behavior and stability in the Rust version. Prisma Compute's switch is a real-world validation.
13,000+ unsafe blocks. This is not idiomatic Rust. It's Zig semantics dressed in Rust syntax. The borrow checker's guarantees are opt-in, and most of the codebase has opted out.
No human review. The PR was reviewed by claude[bot] and coderabbitai[bot] . No human read the full diff. For a runtime that hundreds of thousands of developers rely on, this is unsettling.
Vendor lock-in. Anthropic owns Bun. The only tool that can meaningfully maintain this codebase is Claude. If you want to contribute to Bun, you need access to Anthropic's models. This isn't open source in the traditional sense anymore.
Maintainability debt. AI-translated code preserves the structure of the original language. It doesn't adopt the idioms of the target language. The result is code that's hard for humans to read, refactor, or extend.
QA concerns. Several commenters pointed out that the AI modified some tests to pass instead of fixing the underlying implementation. When an LLM is writing both the code and the tests, the validation loop is compromised.
"That 0.2% don't worry about that" -- u/codemuncher, sarcastically
Community fracture. Past contributors feel alienated. The Zig community sees it as a betrayal. Rust developers see it as a misuse of their language. Some are already looking at forking the last Zig version.
What This Means for the Industry
This rewrite is a bellwether, not a fluke. Whether you love it or hate it, the approach works well enough to ship. Here's what it signals:
Language migrations are no longer multi-year projects. If a frontier model can translate 570K lines of Zig to passable Rust in six days, the cost of switching languages has cratered.
The bottleneck is shifting. As Boris Cherny noted, code writing was the first bottleneck AI solved. Code review became the next -- solved by "teams of Claudes." Maintainability and security are the new constraints. This pattern will repeat.
Rust may become the default language for AI-generated systems code. Its compiler feedback loop is unmatched. Go catches less. Python catches nothing at compile time. C and C++ lack the safety net. Rust's "andon cord" makes it the safest choice when you can't review what the AI wrote.
Trust is the open question. The code passes tests. It ships. But nobody has read it. For personal projects or internal tools, that's fine. For the JavaScript runtime that powers your production Node.js replacement, it's a leap of faith.
Bun 1.4 is many things at once: an engineering experiment, a marketing coup for Anthropic, a stress test for AI-generated code, and a genuine attempt to ship a better runtime.
The rewrite is neither the miracle its proponents claim nor the disaster its detractors predict. It's a pragmatic bet: ship the AI-generated version now, and clean it up later. Whether that cleanup ever happens -- and whether the community sticks around to see it -- is the story we'll be watching in the months ahead.
One thing is certain: we've crossed a threshold. When an AI can translate a million-line systems codebase between languages in under a week, the rules of software engineering have changed. Whether we're ready for that change is a different question entirely.
Il futuro sarà una grande intelligenza artificiale o tante piccole specializzate?
Fable 5 è davvero così potente come AI da essere vietata?
Cina: supremazia intelligenza artificiale e lavoro. Una contraddizione?

---
source: "https://packagemain.tech/p/my-thoughts-on-the-future-of-go-in-ai-era"
hn_url: "https://news.ycombinator.com/item?id=48925839"
title: "My thoughts on the future of Go in the AI era"
article_title: "My thoughts on the future of Go in the AI era"
author: "der_gopher"
captured_at: "2026-07-15T20:03:20Z"
capture_tool: "hn-digest"
hn_id: 48925839
score: 2
comments: 0
posted_at: "2026-07-15T19:23:03Z"
tags:
  - hacker-news
  - translated
---

# My thoughts on the future of Go in the AI era

- HN: [48925839](https://news.ycombinator.com/item?id=48925839)
- Source: [packagemain.tech](https://packagemain.tech/p/my-thoughts-on-the-future-of-go-in-ai-era)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T19:23:03Z

## Translation

タイトル：AI時代の囲碁の未来について思うこと
説明: 私の少し物議を醸す意見の 1 つは、AI によって退屈な言語の価値が下がるのではなく、むしろ価値が上がるのではないかということです。

記事本文:
AI時代の囲碁の未来について思うこと
AI時代の囲碁の将来についての私の考え
私の少し物議を醸す意見の 1 つは、AI によって退屈な言語の価値が下がるのではなく、むしろ価値が上がるのではないかということです。
Alex Pliutau 2026 年 7 月 15 日 3 シェア 正直に言うと、最近は思ったほど Go コードを書けていません。
理由の 1 つは、現在の会社では Go を使用していないことです。私たちは主に TypeScript を使用していますが、これについてはまだ複雑な気持ちがありますが、それでも公平に言うと、TypeScript は私たちにとって非常にうまく機能します。したがって、ここでは TypeScript をあまり非難しません。
私は今でもいくつかの Go プロジェクトを維持しているので、少なくともある程度の量の Go を時々書きます。
でも、ああ、これ以上書くのが恋しいです。
特に開発者の経験が懐かしいです。私にとってこれに匹敵するものは何もありませんが、興味深いことに、AI 時代にはそれがさらに重要になると思います。そしてこの記事では、それについて少し暴言を吐きたいと思います。
AI はこれまで以上に多くのコードを記述します。レビューするといっても、40 個のファイルをアルファベット順にスクロールする必要はありません。
CodeRabbit Review は、フラット ファイル リストからのプル リクエストを、構造化されたレイヤーごとのウォークスルーに再編成します。これは、プラットフォームが変更をソートする順序ではなく、論理的に変更を読み取る順序です。すべての範囲には、シーケンス図、ステート マシン、ビジュアルが配置される場所にインラインで生成される ERD を含む、独自の平易な概要が表示されます。
コホートは関連するファイルとチャンクをグループ化するので、一度に 1 つのアイデアを確認できます。レイヤーは、基本的な変更 (データ形状、コントラクト) が、レイヤーに依存するコードよりも前に来るように順序付けします。コード ピークを使用すると、任意の変数、関数、クラス、または型をクリックして、タブを離れることなくその定義と使用法を確認できます。また、セマンティック差分ビューでは、書式設定のノイズをカットして、実際に何が変更されたかを表示できます。
Review Ch から直接開きます

PR ウォークスルーの [スタック] ボタン。キーボードからコホートやレイヤーに移動し、正確な行範囲に対してコメントし、ネイティブ レビュー、コメント、承認をチームが期待する場所に GitHub または GitLab にポストバックして送信します。
早期アクセスでは、誰でも無料で利用できます。
AI コード レビューの先駆者チームによるもの。毎週 200 万件のレビューが行われます。 600万リポジトリ。 15,000 人の顧客。最新の PR が実際に書かれる方法に合わせて構築されたレビュー インターフェイス。
CodeRabbit で次の PR を今すぐレビューしてください
コーディング エージェントとコードの書き方に関して、状況は明らかに変わりつつあります。ちなみに、私たち全員がというわけではありませんが、私は AI に大きく依存せずにソフトウェアを手動で作成する人々を今でも非常に尊敬しています。そのレベルのエンジニアリングの職人技を発揮できる余地は常にあると思います。
つまり、状況は良くも悪くも変化しており、まだ完全にはわかりませんが、すでにいくつかの傾向が形成されているのがわかります。 AI が生成するコードでは Rust と TypeScript が有力な選択肢になりつつありますが、他の言語については... それらについて耳にすることがますます少なくなっているとだけ言っておきましょう。
本当の疑問は、今から 5 年後も重要なプログラミング言語は何でしょうか?
新しいプロジェクトでは、動的バックエンド言語は少なくなると思います。 Ruby on Rails、PHP、さらには一部のバックエンド ワークロードには Python も含まれます。もちろん、それらの言語が悪いからではありません。それらの多くは人間にとって素晴らしいものです。
AI がなければ、言語は人間の人間工学に関して多くの競争をしていました。現在、彼らは機械の人間工学でも競争しています。
Ruby on Rails を例に考えてみましょう。これは信じられないほど人間に優しいものですが、パフォーマンスやタイピングの点で必ずしもマシンに優しいとは言えないと思います。人間はコンテキストを理解しているため、慣例、暗黙のマジック、DSL を非常に自然に操作できます。モデルはそれにさらに苦労します。明示的なsy

ステムは単に彼らにとって推論しやすいだけです。
そして、いわゆる AI ネイティブ言語と呼ばれる新しい言語も登場しています。
正直なところ、これを読んでも、ここの強みが何なのかまだわかりません
人間またはエージェントに対して異なる方法でドキュメントをルーティングすることさえあります。違いますか？
正直に言うと、私はまだピッチを完全に理解していません。どちらも、「おい、プログラミング言語にしてくれ、間違えるなよ」というようなプロンプトで作られたAI生成のいい加減な言語のように感じます。
現在のモデルは、主に既存の言語と既存のエコシステムでトレーニングされています。では、成熟したツールも大規模な本番コードベースも戦闘テストもほとんどなく、実世界のデータもほとんどない状態で、まったく新しい言語で完璧なコードを突然生成するにはどうすればよいのでしょうか?
セキュリティやコンパイルについてはどうですか?ブラックボックスのような感じ。
したがって、現時点でも生態系は依然として非常に重要です。
そこが Go が本当に面白いところです。
なぜなら、私の意見では、そして多くの人は異論があるでしょうが、Go は実際には LLM にとって非常に優れた言語だからです。人間だけでなく模型も同様です。そして正直に言うと、私はこの大会で優勝することを心から応援しています。
それでは、Go がエージェント時代に特有の利点をもたらすと私が考えるいくつかの点を確認してみましょう。
まず、標準ライブラリは非常に大規模で、信じられないほどよく設計されています。驚くほど依存関係が少ない実際の実稼働システムを構築できます。 HTTP サーバー、JSON 処理、暗号化、テスト、同時実行、ファイル システムなど、そのほとんどはすでに存在しています。
そしてもう 1 つ重要なことは、Go チームはセキュリティと安定性を非常に重視しているということです。 Go 標準ライブラリ自体が原因でエコシステム レベルの壊滅的なセキュリティ災害が発生したという話を聞くことはほとんどありません。
次にコンパイラです。 Go は非常に高速にコンパイルします。
AI時代において、それは人々が思っている以上に重要です。人間は、より遅いフィードバック ループを許容できます。銀

entsは本当にできません。 AI システムは、タスクを解決する際に何百回もコンパイルと反復を行う可能性があります。高速コンパイルはそのワークフローを直接改善します。
単一の静的バイナリ配布
シンプルで意図的に退屈な構文
Go は実用的です。物事を行うためのほぼ明白な方法が 1 つあります
そして、これは実際、Go の最大の強みの 1 つです。 Go は、「創造的な」コードを書くための多くの方法を意図的に削除しました。混乱させる方法が単純に少なくなっているため、AI で生成された Go コードは、多くの場合、驚くほどクリーンで保守しやすいものになることがわかります。
私のお気に入りの機能の 1 つは、Go の互換性の約束です。
Go 1 に戻って、チームは基本的に、古いプログラムは新しい Go バージョンでも動作し続ける必要があるという互換性保証を導入しました。
これは、2 年前の依存関係がすでに機能しなくなっている最新の JavaScript エコシステムと比較するまでは退屈に思えます。
ほぼ 10 年前に作成した Go プロジェクトを今でも開き、最新の Go バージョンで問題なく実行できます。
正直、このレベルの安定性は非常にまれです。
そして、特に AI が大量のコードを生成し始めると、安定性が再び最も価値のあるエンジニアリング特性の 1 つになる可能性があると思います。誰もが望んでいないのは、6 か月ごとに壊れる脆弱なエコシステムに依存して生成された何百万行ものコードです。
私の古いプロジェクト (10 年前) の 1 つを見つけました: https://github.com/plutov/games
また、特に AI が生成したコードを依然として人間がレビューしている間は、書式設定は人々が認めるよりもはるかに重要であると思います。
そして、gofmt だけでもここに大きな価値が加わります。
どの Go プロジェクトもほぼ同じに見えます。スタイルガイドは基本的に言語自体に組み込まれています。フォーマットルールについて議論したり、15 個のより適切なプラグインやリンター設定をインストールしたりするのに時間を費やす必要はありません。
AI がコードを生成するとき、これは次のようになります。

文体上のエントロピーは決して蓄積されないため、エンシーの価値はさらに高まります。
次に、go vet、lint が続きますが、これもプロジェクト間で一貫しています。
もう 1 つの非常に過小評価されている機能は、クロスコンパイルです。
これは今でも私にとって魔法のように感じられます。
GOOS=windows GOARCH=amd64 ビルドに行く
GOOS=linux GOARCH=arm64 go build 以上です。
複雑なビルド パイプラインや外部ツールをセットアップすることなく、まったく異なるプラットフォーム用のバイナリを即座に取得できます。
自律エージェントがエンドツーエンドでシステムの構築と展開を開始する場合、これは非常に重要です。
Go のエラー処理が冗長であるため嫌いな人もいますが、明示的であるため好む人もいます。
エラー:= os.ReadFile()
エラーの場合 != nil {
// ...
}
そして正直に言うと、AI が生成したコードの明示性は驚くべきものです。
制御フローは明らかです。障害パスが表示されます。隠された魔法や曖昧さがはるかに少なくなります。一般に、コードが単純で決定的である場合、モデルのパフォーマンスは向上します。
失敗例を考えざるを得なくなるほどで​​す。
そして、同時実行性があります。
Go には今でも、これまで設計された中で最も優れた同時実行モデルの 1 つが組み込まれています。
go fetchData() 現在でも、Go での同時作業の生成は、他の多くのエコシステムと比較すると、途方もなく軽量に感じられます。
シンプルなプリミティブ。シンプルなメンタルモデル。
私の少し物議を醸す意見の 1 つは、AI によって退屈な言語の価値が下がるのではなく、むしろ価値が上がるのではないかということです。
なぜなら、保守性は賢さよりも優れているからです。
生成されたコードは天才レベルのエンジニアリングである必要はありません。理解しやすく、安定しており、リファクタリング可能で、検証が簡単である必要があります。
そして、Go はまさに正しい意味で退屈であることを誇りに思っています。
Go が AI 時代の主流の言語になるかどうかはわかりません。
しかし、信じられないほど良く熟成すると思います。
Go は依然として非常にシンプルに感じられます。
そして正直に言うと、単純です

柔軟性は、AI 時代全体で最も重要な機能の 1 つになる可能性があります。
3 シェア この投稿に関する前のディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

One slightly controversial opinion I have is that AI might actually make boring languages more valuable, not less.

My thoughts on the future of Go in the AI era
Subscribe Sign in My thoughts on the future of Go in the AI era
One slightly controversial opinion I have is that AI might actually make boring languages more valuable, not less.
Alex Pliutau Jul 15, 2026 3 Share To be fully honest, nowadays I don’t write as much Go code as I wanted.
One reason is that we don’t use Go at my current company. We mostly use TypeScript, about which I still have mixed feelings, but to be fair, it works pretty well for us nevertheless. So I won’t blame TypeScript much here.
I still maintain a few Go projects, so I write at least some amount of Go from time to time.
But oh boy, I miss writing more of it.
Especially I miss the developer experience. Nothing even comes close for me, and interestingly, I think that becomes even more important in the AI era. And in this article I’d like to rant a bit about that.
AI writes more code than ever. Reviewing it shouldn’t mean scrolling forty files in alphabetical order.
CodeRabbit Review reorganizes any pull request from a flat file list into a structured, layer-by-layer walkthrough - the logical reading order of the change, not the order your platform happens to sort it. Every range gets its own plain-language summary, with sequence diagrams, state machines, and ERDs generated inline wherever a visual earns its place.
Cohorts group related files and chunks so you review one idea at a time. Layers order them so foundational changes - data shapes, contracts - come before the code that depends on them. Code Peek lets you click any variable, function, class or type to see its definition and usages without leaving the tab, while Semantic Diff view cuts past formatting noise to show what actually changed.
Open it straight from the Review Change Stack button in the PR Walkthrough. Navigate cohorts and layers from the keyboard, comment against exact line ranges and submit native reviews, comments and approvals post back to GitHub or GitLab right where your team expects them.
In the early access, available for free to everyone.
From the team that pioneered AI code reviews. 2M reviews every week. 6M repos. 15K customers.The review interface built for the way modern PRs are actually written.
Review your next PR with CodeRabbit Review Today
Things are clearly changing now with coding agents, and the way we write code. I don’t mean all of us, by the way, I still have huge respect for people who manually craft software without relying heavily on AI. I think there will always be space for that level of engineering craftsmanship.
So, things are changing, for better or worse, I am still not completely sure, but you can already see some trends forming. Rust and TypeScript are becoming dominant choices in AI-generated code, while some other languages… let’s just say you hear less and less about them.
So the real question is: what programming languages will still matter in 5 years from now?
I think we will see fewer dynamic Backend languages in new projects. Ruby on Rails, PHP, maybe even Python for some Backend workloads. Not because those languages are bad, obviously. Many of them are fantastic for humans.
You see, without AI, languages competed a lot on human ergonomics. Now they also compete on machine ergonomics.
Take Ruby on Rails for example. It is incredibly human-friendly, but I would argue not necessarily machine-friendly, performance-wise and typing-wise. Humans can navigate conventions, implicit magic, and DSLs very naturally because they understand context. Models struggle more with that. Explicit systems are simply easier for them to reason about.
And then there are new so-called AI-native languages.
Honestly, reading that I am still not sure what are the strengths here
they even route the docs differently for humans or agents. Is it different?
And honestly, I still don’t fully understand the pitch. Both of them feel like AI generated slop languages made with a prompt like “Hey, make me a programming language, do not make mistakes”.
Models today are trained mostly on existing languages and existing ecosystems. So how exactly are they suddenly supposed to generate perfect code in a completely new language with no mature tooling, no massive production codebase, no battle testing, and very little real-world data?
What about security, compilation? Feels like a blackbox.
So right now, ecosystems still matter a lot.
And that is where Go becomes really interesting.
Because in my opinion, and many people will disagree, Go is actually an extremely good language for LLMs. Not just for humans, but for models too. And honestly, I am really rooting for it to win this competition.
So let’s review a few things that I think make Go uniquely good for the agentic era.
First, the standard library is absolutely massive and incredibly well-designed. You can build real production systems with surprisingly few dependencies. HTTP servers, JSON handling, crypto, testing, concurrency, file systems — most of it is already there.
And another important thing: the Go team takes security and stability very seriously. You rarely hear about catastrophic ecosystem-level security disasters coming from the Go standard library itself.
Then there is the compiler. Go compiles insanely fast.
That matters much more in the AI era than people realize. Humans can tolerate slower feedback loops. Agents really cannot. AI systems might compile and iterate hundreds of times while solving a task. Fast compilation directly improves that workflow.
single static binary distribution
simple and intentionally boring syntax
Go is pragmatic. It has one mostly obvious way to do things
And this is actually one of Go’s greatest strengths. Go intentionally removed many ways to write “creative” code. There are simply fewer ways to make a mess, which means AI-generated Go code often ends up looking surprisingly clean and maintainable.
One of my favorite features: Go’s compatibility promise .
Back in Go 1, the team introduced a compatibility guarantee that basically says old programs should continue working on newer Go versions.
This sounds boring until you compare it to some modern JavaScript ecosystems where dependencies from two years ago already stop working.
I can still open Go projects I wrote almost 10 years ago and run them on a modern Go version without issues.
That level of stability is honestly extremely rare.
And I think stability might become one of the most valuable engineering properties again, especially when AI starts generating massive amounts of code. The last thing anyone wants is millions of lines of generated code depending on fragile ecosystems that break every six months.
I found one of my old projects (10 years old): https://github.com/plutov/games
I also think formatting matters much more than people admit, especially while humans still review AI-generated code.
And gofmt alone adds enormous value here.
Every Go project looks roughly the same. The style guide is basically built into the language itself. You do not spend time debating formatting rules or installing fifteen prettier plugins and linter configs.
And when AI generates code, this consistency becomes even more valuable because stylistic entropy never really accumulates.
Then comes go vet, linting, which are also consistent between projects.
Another massively underrated feature is cross-compilation.
This still feels magical to me:
GOOS=windows GOARCH=amd64 go build
GOOS=linux GOARCH=arm64 go build And that’s it.
You instantly get binaries for completely different platforms without setting up complicated build pipelines or external tooling.
That matters a lot when autonomous agents start building and deploying systems end-to-end.
Some people dislike Go’s error handling because it is verbose, but others love it because it is explicit.
err := os.ReadFile()
if err != nil {
// ...
}
And honestly, for AI-generated code, explicitness is amazing.
The control flow is obvious. The failure paths are visible. There is much less hidden magic and much less ambiguity. Models generally perform better when the code is straightforward and deterministic.
You are almost forced to think about failure cases.
And then there is concurrency.
Go still has one of the nicest concurrency models ever designed.
go fetchData() Even today, spawning concurrent work in Go feels absurdly lightweight compared to many other ecosystems.
Simple primitives. Simple mental model.
One slightly controversial opinion I have is that AI might actually make boring languages more valuable, not less.
Because maintainability scales better than cleverness.
Generated code does not need to be genius-level engineering. It needs to be understandable, stable, refactorable, and easy to verify.
And Go is proudly boring in exactly the right ways.
I don’t know whether Go will become the dominant AI-era language.
But I do think it will age incredibly well.
Go still feels refreshingly simple.
And honestly, simplicity might end up being one of the most important features in the entire AI era.
3 Share Previous Discussion about this post Comments Restacks Top Latest Discussions No posts

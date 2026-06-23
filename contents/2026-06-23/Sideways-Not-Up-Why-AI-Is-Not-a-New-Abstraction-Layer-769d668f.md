---
source: "https://www.sixsideddice.com/Blog/Thoughts/SidewaysNotUpWhyAIIsNotANewAbstractionLayer.html"
hn_url: "https://news.ycombinator.com/item?id=48641382"
title: "Sideways, Not Up: Why AI Is Not a New Abstraction Layer"
article_title: "Sideways, Not Up: Why AI Is Not a New Abstraction Layer - SixSidedDice.com - Blog"
author: "rusk"
captured_at: "2026-06-23T07:10:26Z"
capture_tool: "hn-digest"
hn_id: 48641382
score: 1
comments: 0
posted_at: "2026-06-23T07:08:44Z"
tags:
  - hacker-news
  - translated
---

# Sideways, Not Up: Why AI Is Not a New Abstraction Layer

- HN: [48641382](https://news.ycombinator.com/item?id=48641382)
- Source: [www.sixsideddice.com](https://www.sixsideddice.com/Blog/Thoughts/SidewaysNotUpWhyAIIsNotANewAbstractionLayer.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T07:08:44Z

## Translation

タイトル: 横向き、上向きではない: AI が新しい抽象化レイヤーではない理由
記事タイトル: Sideways, Not Up: Why AI Is Not a New Abstraction Layer - SixSidedDice.com - ブログ
説明: AI コーディング アシスタントは抽象化のはしごの次の段であり、高級言語がアセンブリの上に位置するのと同じように、高レベル言語の上に位置するという一般的な主張があります。私はその議論の最も強力なバージョンを真剣に受け止め、それが成り立たないと主張します。コンパイラは決定論的な仕様です。
[切り捨てられた]

記事本文:
横向き、上向きではない: AI が新しい抽象化レイヤーではない理由 - SixSidedDice.com - ブログ
2026 年 6 月 18 日
思い
AI
コーディングアシスタント
抽象化
上ではなく横に: AI が新しい抽象化レイヤーではない理由
きちんとした物語が巡回されており、すべてのきちんとした物語と同様に、それはほとんどが明白であるかのように語られます。話はこうなります。ソフトウェアは常に抽象化のはしごを登ってきました。私たちはマシンコードから始め、その後アセンブリによって生のオペコードの代わりにニーモニックが提供され、さらに高級言語とそのコンパイラによって for ループを記述できるようになり、レジスタのことは完全に忘れることができました。各段階では、あまり気にせずに、より多くのことを言うことができます。そして今、物語は結論づけられていますが、AI は単に次の段階です。必要なものを英語で記述すると、モデルがコードを記述します。そして、近い将来、コンパイラーが出力するアセンブリを現在読んでいるときと同じように、そのコードを読まなくなるでしょう。
真実の部分と偽りの部分が融合しているため、説得力のある物語となっています。アンドレイ・カルパシーは、「最もホットな新しいプログラミング言語は英語である」と宣言したとき、真実の部分を記憶に残るものにしました。また、以前は依然としてニューラル ネットワークを「ソフトウェア 2.0」として枠付けしており、「ニューラル ネットワークのトレーニング プロセスによりデータセットがバイナリにコンパイルされる」と述べました。マット・ウェルシュは、ACM のコミュニケーションの中でこれを論理的に極限まで取り上げ、「『プログラムを書く』という従来の考え方は消滅に向かっている」と主張しました。ラダー、コンパイラ、次の層。それはすべてうまく韻を踏んでいます。
私はこれを皮肉屋としてではなく、他の人の抽象化が静かに裏切ったシステムを構築し、保守し、時には救出する 30 年以上の経験を持つソフトウェア職人としてこれを書いています。 AI コーディングアシスタントのために多くの時間を費やしています。毎日使っています。しかし、抽象化レイヤーの枠組みは間違ったメンタルモデルだと思います。

そして、メンタルモデルが間違っていると、チームは最終的に驚くことになります。
高級言語がマシンコードを抽象化したものである場合、AI は本当に高級言語を抽象化したものなのでしょうか、それともまったく同じ衣装を着た別のものなのでしょうか?
はしご、そしてそれがとても魅惑的な理由
反対側の主張を適切に説明しましょう。怠惰なバージョンが示唆するよりも強力だからです。
この議論の正直な核心は、プログラミングは常に意図を圧縮する行為であるということです。機械語で意図を表現するのは時間がかかり、エラーが発生しやすく、耐えられないため、機械語で記述しません。私たちが追加した各レイヤーは、同じシリコンに対して同じことを言うためのより人道的な表記になっています。このように考えると、自然言語はこれまでで最も人道的な表記法にすぎません。マーク・ブルッカーはこの点をうまく主張し、その過程で明らかな反対意見を武装解除します。自然言語はプログラミング媒体としては曖昧すぎると批評家は言うが、彼は「ほとんどすべてのプログラムはすでに自然言語で指定されている。そしてこれまでもそうしてきた」と答えた。要件文書、チケット、製品所有者との会話: これらは常に実際のソースであり、コードは単なる転写にすぎませんでした。 AI は転写ステップを削除するだけです。
いつものように、マーティン・ファウラーはここで最も慎重な発言者であり、彼は変化の大きさを過小評価していません。同氏は、LLM は「アセンブラから最初の高級プログラミング言語への変化と同じ程度にソフトウェア開発を変えるだろう」と考えており、「プロンプト内でマシンと対話することは、Fortran とアセンブラと同じくらい Ruby と異なる」と考えている。これは決して小さな主張ではありませんし、誇大広告に乗りやすい人物が主張したものでもありません。反対側にとって最も強力なケースが必要な場合、それは、最初にコンパイラを提供したものと同じくらい重要な移行です。

その大きさには同意します。方向性については同意しません。そして、これら 2 つの違いが議論全体です。
本当の抽象化が実際に行うこと
AI が新しい抽象化レイヤーであるかどうかを問う前に、抽象化レイヤーが何のためにあるのかを正確に把握する必要があります。優れた抽象化のポイントは、複雑さを隠すことではありません。重要なのは、その下のレイヤーについて考えるのをやめることができるということです。それが耐荷重特性です。 C# を作成してコンパイラーに渡すとき、生成される IL コードや実行時に発行される x86 命令についてはまったく考えません。私は読みません。レビューはしません。コンパイラーに問題があった場合に備えて、コピーは保存しません。抽象化によって私の信頼が得られ、その信頼が私に生産性をもたらしてくれます。
信頼性が高いからこそ、私たちが当たり前だと思ってしまいがちな 3 つの特性を通じて、その信頼を獲得しています。これは決定的です。同じコンパイラとフラグでコンパイルされた同じソースは毎回同じ出力を生成します。これが再現可能なビルドの前提条件です。それは指定されています。言語には文法とセマンティクスがあるため、ソースからマシンコードへの変換は、私が推論できる定義された動作であり、驚くべきことではありません。そして、それは検証可能です。何か問題が発生した場合、定義されたマッピングを逆算して理由を見つけることができます。
コンパイラは約束を守ります。モデルはそうではありません。
ここが梯子を壊す物件です。 1995 年に Fortran 関数を作成したとき、それを 100 回コンパイルし、同じバグを含む同じバイナリを 100 回取得できました。バグは私のもので、安定していました。安定性のおかげで私はバグを追い詰めることができました。大規模な言語モデルではそのような保証はありません。同じ質問を 2 回すると、異なる 2 つの質問が返される可能性があります

なぜなら、モデルは確率分布からトークンをサンプリングすることによって機能し、サンプリングはその構造上、サイコロの出目であるからです。
これは単なる温度設定であり、それをゼロにすると決定性が得られると考えるかもしれません。そうではなく、その理由は有益です。 Thinking Machines Lab のチームは、温度がゼロであっても「LLM API は実際にはまだ決定的ではない」ことを実証しました。彼らは、温度ゼロでモデルに対して同じプロンプトを 1,000 回実行し、80 個の異なる完了を取得しました。最初の百数十個のトークンは同一で、その後分岐しました。原因は、よくある浮動小数点の民間伝承ではなく、バッチの不変性の欠如であることが判明しました。モデル自体のカーネルの数値結果は、その瞬間にサーバー上であなたのリクエストと一緒に偶然バッチ処理された他の人のリクエストの数に依存します。あなたの出力は見知らぬ人のトラフィックに依存します。それをエンジニアリングを構築するための基礎として浸透させてください。
ファウラーはまさにこれを見て、私よりも適切な名前を付けました。彼は、「プロンプトを git に保存するだけで、毎回同じ動作になることを知っている」ことはできないと指摘し、同僚の Birgitta Böckeler の言葉を引用して、この記事のタイトルになっています。
私たちは抽象化レベルを上げているだけではなく、同時に非決定論へと横向きに移行しています。
それが核心です。ラダーのイメージが間違っているのは、ジャンプが小さいからではなく、ジャンプアップではないからです。これまでの各段は、表現力を高めながら決定性を維持しました。これは、決定論を捨てて表現力を獲得したもので、真に新しく奇妙な場所への横移動です。それを革命と呼ぶこともできますが、私は異論はありません。しかし、再現性を放棄する革命はコンパイラと同じ種類のものではありません。

そうなるとチームが傷つくことになるだろう。
決定論だけが問題のすべてではない
さて、鋭い読者と正直な人はここで反発するでしょう、それで私は自分自身を押し戻させてください。 Thinking Machines の研究は、非決定論を診断するだけではありません。それはほとんど修正されます。バッチ不変カーネルを使用すると、1,000 回の完了結果はすべて同一でした。したがって、断固とした反対者はこう言うことができます。「いいですよ、非決定論は展開上の成果物であり、物理法則ではありません。推論スタックが適切に構築されれば、反対意見は消えます。」
決定論は決して最も深い問題ではなかったので、それは蒸発しません。それは目に見えるものでした。最良のケースを想像してください。それは、同一のプロンプトに対してバイト同一のコードを永久に返す、完全に再現可能なモデルです。仕様はまだ存在せず、正確性の保証もありません。コンパイラは言語標準に拘束されます。準拠したコードのコンパイルが間違っている場合、それはコンパイラのバグであり、標準規格がそのように指示します。モデルは何にも拘束されません。定義されたルールに従って英語をコードに変換するわけではありません。これはもっともらしいコードを予測するものですが、もっともらしいことと正しいことは同じではありません。定義されたマッピングを通じてプロンプトから出力までを推論することはできません。つまり、抽象化によって実行できるはずの 1 つのこと、つまり、下のレイヤーの読み取りを停止することができないことを意味します。
そしてそれが試練であり、挑戦と言い換えることができます。実際の抽象化レイヤーを使用すると、下のレイヤーの読み取りを停止できます。 LLM が書いたコードを読むのをやめてもらえますか?できないし、できないことはわかっています。モデルが良くなればなるほど、失敗は稀になり、より微妙になり、間違いが確実になるため、試みるのはより危険になります。決定論により出力が安定します。それは信頼できるものではありません。これらは異なる美徳であり、そのうちの 1 つだけが提供されます

。
これはどれも新しいものではありませんが、これが私が最も伝えていると思う部分です。平易な英語でプログラミングするという夢は、プログラミング自体とほぼ同じくらい古いものです。 COBOL は 1959 年に、「最初の英語風のデータ処理言語」である Grace Hopper の FLOW-MATIC を基にして、管理者が読めるように意図的に英語風になるように設計されました。 1980 年代の第 4 世代言語は、今日私たちが聞いているとおり、「ビジネス アナリストなどのエンドユーザーがアプリケーションを独自に作成および変更できるようにして、専門の IT 担当者への依存を最小限に抑える」という約束に基づいて販売されました。それがどのように展開されたかについては、『Vibecoding: The Empire's New Clothes』で書きましたが、端的に言えば、英語のような構文では、エッジケースを正確に考えることができる人の必要性が決して排除されなかったということです。プログラムが長くなっただけです。
最も萎縮する評決は、まさにその熱狂の波に対して1978年に書いたエドガー・ダイクストラによってもたらされた。彼の中心的な洞察は、プログラミングの正式な表記法は機能であり、負担ではないということでした。なぜなら、形式的な記号の規律こそが、私たちが気づかずにナンセンスな文章を書くのを防ぐものだからです。彼のセリフは恥ずべきほどに老け込んでいる。
結局のところ、私たちが母語を使用するときの「自然さ」とは、つまるところ、ナンセンスであることが明らかではない発言をするために母語を簡単に使用できるかということです。
これが、自然言語が貧弱な抽象化と素晴らしいインターフェイスであるにもかかわらず、同じものではない最も深い理由です。英語を素晴らしい表現力にしている曖昧さは、私たちとモデルが自信を持ってもっともらしいくだらないものを生み出すことを可能にしているのと同じ曖昧さであり、その間違いは制作されるまで明らかではありません。ダイクストラ氏は、私たちの母国語でプログラムされた機械は「製造するのが非常に難しい」という予言で締めくくった。

ブルッカー氏の名誉のために言っておきますが、ブルッカー氏は、2025 年がついにそれらを構築したと主張しています。私は、ダイクストラが警告したのとまったく同じくらい安全に使用するのが難しいにもかかわらず、私たちは並外れたものを構築したと言えます。
私たちの目の前にあるツールは本当に素晴らしいものなので、過剰請求から却下にならないように注意したいです。それで、ここが私が着陸する場所です。
LLM のプロンプトは、コードを生成する行為に対する強力な新しいインターフェイスです。これはおそらく、IDE 以来、そしておそらく高級言語以来、私たちの作業方法に対する最も重要な変更です。ただし、インターフェイスは抽象化レイヤーではありません。抽象化レイヤーの決定的な約束は、その下にあるものを忘れてしまう可能性があるということですが、AI は、論理的に反論できる仕様も信頼できる保証も提供しないため、今日でも原則的にもその約束はできません。代わりに提供されるのは、素晴らしく、高速で、時には素晴らしいこともあれば間違ったこともあるドラフトであり、そのドラフトに対して全責任を負うのはあなたです。生成されるコードは依然としてあなたのコードです。それを所有し、レビューし、デバッグし、午前 3 時に転倒したときにメンテナンスを行います。それはコンパイラの出力を所有するときの気分ではありません。これら 2 つの感情の間のギャップこそが重要なのです

[切り捨てられた]

## Original Extract

There is a popular claim that AI coding assistants are the next rung on the ladder of abstraction, sitting above high-level languages just as those languages sat above assembly. I take the strongest version of that argument seriously, then argue it does not hold. A compiler is a deterministic, speci
[truncated]

Sideways, Not Up: Why AI Is Not a New Abstraction Layer - SixSidedDice.com - Blog
Jun 18, 2026
Thoughts
AI
CodingAssistants
Abstraction
Sideways, Not Up: Why AI Is Not a New Abstraction Layer
There is a tidy story doing the rounds, and like all tidy stories it is mostly told as if it were obvious. The story goes like this. Software has always climbed a ladder of abstraction . We started with machine code, then assembly gave us mnemonics instead of raw opcodes, then high-level languages and their compilers let us write for loops and forget about registers entirely. Each rung let us say more while caring about less. And now, the story concludes, AI is simply the next rung. You describe what you want in English, the model writes the code, and one day soon you will no more read that code than you currently read the assembly your compiler emits.
It is a compelling narrative because it has a true part welded to a false part. Andrej Karpathy put the true-feeling part memorably when he declared that "the hottest new programming language is English" , and earlier still framed neural networks as "Software 2.0" in which "the process of training the neural network compiles the dataset into the binary". Matt Welsh took it to its logical extreme in the Communications of the ACM, arguing that "the conventional idea of 'writing a program' is headed for extinction" . The ladder, the compiler, the next layer up. It all rhymes nicely.
I write this not as a cynic, but as a software craftsperson with over 30 years of experience building, maintaining, and occasionally rescuing systems that other people's abstractions had quietly betrayed. I have a great deal of time for AI coding assistants. I use them daily. But I think the abstraction-layer framing is the wrong mental model, and getting the mental model wrong is how teams end up surprised.
If high-level languages are an abstraction over machine code, is AI really an abstraction over high-level languages, or is it something else entirely wearing the same costume?
The Ladder, and Why It Is So Seductive
Let me make the case for the other side properly, because it is stronger than the lazy version suggests.
The honest core of the argument is that programming has always been the act of compressing intent. We do not write in machine code because expressing intent that way is slow, error-prone, and unbearable. Every layer we have added has been a more humane notation for saying the same thing to the same silicon. Seen this way, natural language is just the most humane notation yet. Marc Brooker makes this point well, and disarms the obvious objection in the process. Critics say natural language is too ambiguous to be a programming medium, to which he replies that "almost all programs are already specified in natural language. And always have been." The requirements document, the ticket, the conversation with the product owner: these were always the real source, and code was merely the transcription. AI just removes a transcription step.
Martin Fowler is, as usual, the most careful voice here, and he does not undersell the magnitude of the shift. He reckons LLMs "will change software development to a similar degree as the change from assembler to the first high-level programming languages", and that "talking to the machine in prompts is as different to Ruby as Fortran to assembler". That is not a small claim, and it is not made by someone prone to hype. If you want the strongest case for the other side, that is it: a transition as significant as the one that gave us the compiler in the first place.
I agree with the magnitude. I disagree about the direction. And the difference between those two things is the whole argument.
What a Real Abstraction Actually Does
Before we can ask whether AI is a new abstraction layer, we have to be precise about what an abstraction layer is for . The point of a good abstraction is not that it hides complexity. The point is that it lets you stop thinking about the layer below . That is the load-bearing property. When I write C# and hand it to a compiler, I genuinely do not think about the IL code that is generate or the x86 instructions it emits at execution time. I do not read them. I do not review them. I do not keep a copy in case the compiler has a bad day. The abstraction has earned my trust, and that trust is what buys me the productivity.
It earns that trust through three properties that we tend to take for granted precisely because they are so reliable. It is deterministic : the same source compiled with the same compiler and flags produces the same output, every single time, which is the entire premise behind reproducible builds . It is specified : the language has a grammar and a semantics, so the translation from source to machine code is defined behaviour I can reason about, not a surprise. And it is verifiable : when something goes wrong, I can reason backwards through a defined mapping to find out why.
The Compiler Keeps Its Promises. The Model Does Not.
Here is the property that breaks the ladder. When I wrote a Fortran function in 1995, I could compile it a hundred times and get the same binary with the same bugs a hundred times. The bugs were mine, they were stable, and stability is what let me hunt them down. A large language model offers no such promise. Ask it the same question twice and you may get two different programs, because the model works by sampling tokens from a probability distribution, and sampling is, by construction, a roll of the dice.
You might assume this is just the temperature setting and that turning it to zero buys you determinism. It does not, and the reason is instructive. The team at Thinking Machines Lab demonstrated that even at temperature zero, "LLM APIs are still not deterministic in practice". They ran the same prompt through a model a thousand times at temperature zero and got eighty distinct completions, identical for the first hundred-odd tokens and then diverging. The culprit turned out not to be the usual floating-point folklore but a lack of batch invariance : the numerical result of the model's own kernels depends on how many other people's requests happen to be batched alongside yours on the server at that instant. Your output depends on a stranger's traffic. Let that sink in as a foundation to build engineering on.
Fowler saw exactly this and named it better than I could. He points out that he cannot "just store my prompts in git and know that I'll get the same behaviour each time", and then quotes his colleague Birgitta Böckeler with the line that gives this article its title:
we're not just moving up the abstraction levels, we're moving sideways into non-determinism at the same time.
That is the crux. The ladder image is wrong not because the jump is small, but because it is not a jump up . Every previous rung preserved determinism while raising expressiveness. This one trades determinism away to get expressiveness, which is a sideways move into a genuinely new and stranger place. You can call that a revolution, and I would not argue. But a revolution that abandons reproducibility is not the same kind of thing as a compiler, and pretending it is will get teams hurt.
Determinism Is Not Even the Whole Problem
Now, a sharp reader and an honest one will push back here, so let me push back on myself. The Thinking Machines work does not only diagnose the nondeterminism; it largely fixes it. With batch-invariant kernels, all thousand completions came out identical. So a determined opponent can say: fine, nondeterminism is a deployment artefact, not a law of physics, and once the inference stack is built properly your objection evaporates.
It does not evaporate, because determinism was never the deepest problem. It was the visible one. Imagine the best case: a perfectly reproducible model that returns byte-identical code for an identical prompt forever. You still do not have a specification , and you still have no correctness guarantee . A compiler is bound by a language standard; if it miscompiles conforming code, that is a bug in the compiler and the standard tells you so. A model is bound by nothing. It is not translating your English into code according to defined rules. It is predicting plausible code, and plausible is not the same as correct. You cannot reason from the prompt to the output through any defined mapping, which means you cannot do the one thing that an abstraction is supposed to let you do: stop reading the layer below.
And that is the test, restated as a challenge. A real abstraction layer lets you stop reading the layer below. Can you stop reading the code an LLM writes for you? You cannot, and you know you cannot, and the better the model gets the more dangerous it becomes to try, because the failures get rarer and subtler and more confidently wrong. Determinism would make the output stable . It would not make it trustworthy . Those are different virtues, and only one of them is on offer.
None of this is new, which is the part I find most telling. The dream of programming in plain English is roughly as old as programming itself. COBOL was deliberately designed in 1959 to be English-like so that managers might read it, building on Grace Hopper's FLOW-MATIC , "the first English-like data processing language". The fourth-generation languages of the 1980s were sold on exactly the promise we hear today, that they would "empower end-users, such as business analysts, to create and modify applications independently, minimizing reliance on specialized IT personnel". I wrote about how that played out in Vibe Coding: The Emperor's New Clothes , and the short version is that the English-like syntax never removed the need for someone who could think precisely about edge cases. It just made the programs longer.
The most withering verdict came from Edsger Dijkstra , writing in 1978 against that very wave of enthusiasm. His central insight was that the formal notation of programming is a feature and not a burden, because the discipline of formal symbols is exactly what stops us writing nonsense without noticing. His line has aged disgracefully well:
When all is said and told, the "naturalness" with which we use our native tongues boils down to the ease with which we can use them for making statements the nonsense of which is not obvious.
That is the deepest reason natural language is a poor abstraction and a marvellous interface, which are not the same thing. The ambiguity that makes English so wonderfully expressive is the same ambiguity that lets us, and the model, produce confident plausible rubbish whose wrongness is not obvious until production. Dijkstra signed off with a prophecy that machines programmed in our native tongues would be "as damned difficult to make as they would be to use". Brooker, to his credit, argues 2025 has finally built them. I would say we have built something extraordinary that is nonetheless exactly as difficult to use safely as Dijkstra warned.
I want to be careful not to swing from over-claim to dismissal, because the tool in front of us is genuinely remarkable. So here is where I land.
Prompting an LLM is a powerful new interface to the act of producing code. It is probably the most significant change to how we work since the IDE, and arguably since the high-level language. But an interface is not an abstraction layer. The defining promise of an abstraction layer is that you may forget what is beneath it, and AI cannot make that promise, today or in principle, because it offers neither a specification you can reason against nor a guarantee you can rely on. What it offers instead is a brilliant, fast, occasionally brilliant-and-wrong draft that you remain wholly responsible for. The code it generates is still your code. You own it, you review it, you debug it, you maintain it at three in the morning when it falls over. That is not what owning compiler output feels like, and the gap between those two feelings is the entire poin

[truncated]

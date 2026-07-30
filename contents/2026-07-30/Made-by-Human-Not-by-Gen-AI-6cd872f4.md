---
source: "https://www.by-human.net/"
hn_url: "https://news.ycombinator.com/item?id=49114964"
title: "Made by Human, Not by Gen-AI"
article_title: ""
author: "monax"
captured_at: "2026-07-30T20:56:51Z"
capture_tool: "hn-digest"
hn_id: 49114964
score: 2
comments: 0
posted_at: "2026-07-30T19:58:56Z"
tags:
  - hacker-news
  - translated
---

# Made by Human, Not by Gen-AI

- HN: [49114964](https://news.ycombinator.com/item?id=49114964)
- Source: [www.by-human.net](https://www.by-human.net/)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T19:58:56Z

## Translation

タイトル: Gen-AIではなく人間が作ったもの

記事本文:
ウェブサイト |ギットハブ | CC0ライセンス取得済み
これは、プロジェクトの作成に生成 AI (特に LLM) を使用していないことを象徴するバッジのコレクションを提供することを目的としたリポジトリです。これを使用すると、AI によって生成されたコードが、プロジェクト内で記述されたコードの合計行数の 1% 未満になることが期待されます。バッジは、スマート グレッパーなどの支援を求める場合にのみ使用できます。
コードをチェックする人は誰もいないことに注意してください。コードベースでの AI の使用に反対しているが、AI によって記述されたコードの行数が不明な場合でも、バッジを使用できます。目標は、透明性を確保し、コードベースでの AI の悪用を減らすことです。
このリポジトリには、あらゆる主張を裏付ける科学的研究による裏付けを試みる説明が付いています。しかし、この部分はまだ進行中です。
AIの「人間による」使用はどのようなものでしょうか?
コードを生成するためにプロジェクトで AI を使用すべきではありません。代わりに次のように使用できます。
スマート grep またはスマート google 。たとえば、何かが実装または保守されている新しいコードベースをフェッチします。
レビューアーヘルパー。バグをキャッチする AI を備えた PR が多数ある場合。
コードベース (例: X をどのように実行しますか?) を学習して、例で学習したり、ドキュメントがない場合に何かを説明したりできます。プロジェクトを保守している場合は、その後にドキュメントを作成する必要があります。
AI を使用してバグに関する助けを求めるが、自分で修正を作成するためではありません。
機能をどのように実装できるかを尋ねますが、最終的な実装を自分で書くことはしません。
思考を支援するために AI を使用するたびに、人間によるレビューが必要になります。 AI が生成または指示した要素にも責任を負わなければなりません。
AI は慎重に使用し、人間が迅速な解決策を提供できなかった場合にのみ使用する必要があることに注意してください。

オン/代替。 AI を使用する前に、AI の出力とエネルギー消費と環境フットプリントを常に考慮する必要があります。
コードをコピーしてリポジトリに貼り付けることも、プロジェクトでロゴを直接使用することもできます。
## 人間が作った
< a href = "https://github.com/Supercip971/by-human" >
<写真>
< ソースメディア = "(prefers-color-scheme: dark)" srcset = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-light.svg" >
< ソースメディア = "(prefers-color-scheme: light)" srcset = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-dark.svg" >
< img height = "96" align = "right" alt = "Gen AI バッジではなく、humand によって作成されました" src = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-light.svg" >
</画像>
</ a >
[プロジェクト名] は、生成 AI ではなく **人間** によって作成されています。
詳細については、[ by-human ]( https://github.com/Supercip971/by-human ) リポジトリにリンクできます。
人間が作ったもの
[プロジェクト名]は生成AIではなく人間によって作られています。
さらに詳しい情報は、人間によるリポジトリにリンクできます。
生成 AI を使用した執筆にはなぜ問題があると考えられるのでしょうか?
生成 AI を使用したライティングは、さまざまな点で劣っています。これはより優れており効率的であるように見えますが、ここで列挙する重要な欠点があります。
1. 著作権、盗作、ライセンスウォッシュ
AI は無意識のうちにライセンス ウォッシュを行うことができ、それを意図的に行うために使用できることがすでに示されています [1] 。
TL;DR: クロード コードはプロジェクト全体を「書き直す」ために使用され、その後、LGPL ではなく MIT の下で再ライセンスされました。
これは明らかなライセンス違反です。
もしそうでないと主張するなら、書籍を翻訳すると著作権が剥奪されるのでしょうか?
このような主張を受け入れることは、事実上、コピーレフトと共同作業の終焉を意味することになるでしょう。

著作権の保護。
LLMに学習させる場合、コードのライセンスを把握することができません。
2 つの研究論文 [2] [3] で示されているように、大規模言語モデルは強力なコピーレフト ライセンス コードの 3.35% を生成しており、次のとおりです。
コピーレフト コードの再利用を認識していないため、応答で既存のコードの再利用を避けるようにプロンプトを通じて尋ねることはできません。
この論文[2]でも次のように述べられています。
コピーレフトの要求を受け入れると、コピーレフトで盗まれたコードが増加する可能性があります。 (2 ～ 5 倍)。
結局のところ、LLM はあからさまにコードを盗用しているのです。これは、人間がコードの塊について学習し、得られた知識に基づいて何かを作成するのとは大きく異なります。
人間は基本的に全体像とコードの背後にある考え方を理解しており、丸暗記に頼ることはありません。
しかし、LLM が学習すると、さまざまなソースから大量のコードが取り込まれ、それを知らないユーザーにそのまま吐き出す可能性があります。これは、コードの強力にライセンスされた部分がデータベース内に存在することを意味し、ライセンスの尊重に関する懸念が生じます。
これは、人間がライセンスされたコードをインポートするのとは異なります。十分な情報を備えた開発者は、コードの一部とともにライセンス通知を含めます。その代わり、LLM はライセンスや作者について言及しません。
まず数字そのものを見てみましょう。
[4] Chat-GPT 4 は最大 8% の確率で即座に失敗し、コードの品質は最大 60% の時間のみ良好であると述べています。
[5] AI 企業である code Rabbit は、LLM によってプル リクエストの 70% に重大な欠陥が発生し、40% に重大な問題が発生すると発表しました。人間の2倍以上です。
[6] この論文は、Github-copilot がバグを見つけようとするときに表面的な機能しか持たないことを通知しています。バグの発見にはほとんど役に立ちません。脆弱性の問題 (39 種類以上) を抱えた複数のプロジェクトにわたって、GitHub Copilot

カップルを見つけるのに役立つだけで、一般的にはスペルミスを修正するだけです。最後に、GitHub Copilot には貴重なコメントがありませんでした
[7] Uplevel は、生成 AI によって 40% 多くのバグが発生していると報告しています。
[8] は、プッシュされ、その後 2 週間以内に元に戻されるか削除されるコードが前年比で 40% 増加しているという事実を大局的に示しています。 AI の導入以来、すぐに「元に戻された」コードが 3.97% から 7.09% に増加したことを意味します。
容赦なく、AI が増加する技術的負債を生み出し、プロジェクトを長期的に維持できなくなることを私たちは理解しています。そして、それ自体を修復することはほとんどできません。
それは、学生があなたのためにコードを書いているのに、学習して認知的内省をすることができないようなものです。そして、プログラマーであるあなたは、自分でコードを書いたわけではないので、コードを完全に理解する可能性は低くなります。
いくつかの研究 [9] では、AI を使用すると、練習問題でより良い成績を達成できる可能性が 48% ～ 127% 高まることが示されています。しかし、最終的には、実際のテストでは、より悪い成績を収める可能性が 17% 高くなります。
さらに悪いことに、それらの生徒たちは学習量が減っていることに気づくことができなかったのです。そして、より理解することができませんでした。
誤った知識につながるため、これは重大な問題です。通常、AI を使用してコードを作成し、それをチェックすることで自分の能力を信頼します。しかし、AI を信頼して自分で学習しないと、本当にダメなプログラマーになってしまいます。
LLM を使用すると、自分のコードベースを研究することが期待されているため、自分のコードベースを理解することがますます困難になり、その結果、コードベースを修正したり改善したりすることも難しくなります。
最終的には、AI への依存度がさらに高まり、コードベースが維持できなくなるまで悪循環に陥ります。
3. 生態学的側面は壊滅的ですか?
nuに翻訳するのは乱暴です

AI の生態学的側面を考察します。
まず、RAM 生産の 70% [10] がデータセンター専用です。サプライヤーの能力を AI データセンターに再配分することにより、生産が増加しました。つまり、AIを稼働させるために私たちは多くの経済資源を使っているということです。
[11] chat-GPT の導入以来、消費電力は 1 年間で 98% 増加しました。 (2022年に2.69MW→2023年に5.43MW)。
水の使用量を大局的に把握するのは困難です。唯一信頼できる情報源は、Chat-GPT がクエリごとに 0.000085 ガロンの水を使用するという Sam Altman の引用です。 [12] しかし、Chat-GPT は 1 日あたり 25 億のリクエストを処理します [13]。これは、Chat-GPT が 1 日あたり平均 804,400 L の水を使用することを意味します。
それ以降、この記事では [14] 1 つの Chat-GPT 4.5 リクエストに 20,500 Wh の費用がかかることがわかります。ただし、近似値を使用しているため、このステートメントをできるだけ明確にすることはできません。
多くの研究では「短い」リクエストが使用されているため、この記事は大きなコンテキストを考慮に入れており、より根拠のあるものになっています。 LLM をエージェントとして使用するには、ファイルやコードベースを読み取る必要があり、「ショート リクエスト」にリンクすることはできなくなります。
多いように思えるかもしれませんが、これらの数字は幽霊です。私たちはそれ以上の主張をすることはできず、Chat-GPT 使用の直接的な生態学的側面を大局的に捉えることもできません。 OpenAI の洞察を使用した完全な調査が必要になります。その一方で、彼らは多くの情報を公開していないため、私たちはAIを使用することによって世界をどれだけ崩壊させているかを推測することにこだわっています。
4.1 近親交配は人間にとっても同様に悪いものである
Microsoft は github のコードで LLM をトレーニングしており、LLM によって書かれたコードの 40% は変更されないままであると会議で表明しました [15] 。
この引用には実際には何の証拠もありませんが、PU がますます増加していることは事実であると認められています。

ブリッシュされたコードは LLM によって書かれており、徐々に放置されます。
LLM のトレーニング データは、人間のコードと LLM で書かれたコードの間で異なることはできません。この質の変化に対応するには、ますます多くのエネルギー、トレーニング、データが必要になることを示唆しています。
LLM によって複数回繰り返されたエラーが、グラウンド トゥルースになる可能性があります。わずか 20 個のドキュメントだけが、あらゆるサイズの LLM を汚染できることが示されています [16] 。 (これはこの声明に直接リンクしているわけではありませんが、この記事では、いくつかの文書が LLM の視点をどのように変えることができるかを示しています)。
要約すると、モデル知識の容易な移行と、実際の LLM の使用の急増は、LLM が独自のデータに基づいてトレーニングを行っていることを意味し、トレーニングの品質の大幅な低下につながっています。
4.2 コンピューティング能力の可用性
RAM の料金を 2 倍支払わなければならないのは、データセンターが本番環境全体を消費するための大きなコストです。
AI は主に、計算能力、RAM、コンテキスト サイズを増やすことによって進化できます。
しかし、私たちの世界はそれに追いつくことができません [17] 。
この論文は、次の声明で懸念を具体化しています。
経験的には、サットンの「苦い教訓」（サットン、2019）が現れる
一部間違っています。AI にとって「一般的な手法」というわけではありません。
計算を活用することが最終的に最も効果的であるということ、
[理由] ムーアの法則は、[…] 指数関数的に継続しました
計算単位あたりのコストは低下していますが、そのコストは増加しています
リソースはAIに費やされます。このリソースの増加は、
これは計算コストに現れますが、他のコストにも当てはまります。
たとえば、より大規模な AI モデルを構築するには、より多くの人力が必要です。
労働。 [17]
その後、AI が向上していると言うとき、それは画期的なアルゴリズムのせいではなく、次のような理由からです。
より強力なハードウェアはより多くの費用とより多くのエネルギーを消費します
より多くのデータ（汚染がますます増えています）

LLM によって提供されます)
より多くのトレーニングを行うと、より多くのエネルギーとより多くの人的労力が必要になります。
私たちは、増大するリソース需要に対応することが困難な状況に達しつつあり、それを維持する唯一の方法は、ユーザー市場を食い尽くすことです。この論文は最近のラム価格の上昇を予測していた可能性がある [18] 。
十分な RAM がなくなり、十分な計算能力がなくなったとき、AI 業界全体が崩壊し、私たちは現在の状況に戻るかもしれません… そして、AI に依存していた人々は、失われた知識を取り戻すことはできなくなります。
前向きになりたいなら、コンピューターの進化を逆転させるという考えがすぐに浮かぶかもしれません。そして、パフォーマンスを向上させるのではなく、より効率的になろうとしています。
AI を使用してコードを記述する場合、問題はすでに他の人によって解決され、モデルによって再現されている可能性があるため、そのプロセスはほとんど魔法のように見えることがあります。
それはあなたを賢く見せますが、同時に自分で問題を解決するのが苦手になります。 AI は地球と経済を破壊する一方で、AI 自体のデータセットを汚染しながら悪化しています。投資家が数十億ドルを拠出し、ますます多くを期待している一方で、得られる見返りはますます少なくなっており、状況は悪化の一途をたどっている。
それは一時的に輝くだけの石であり、やがてただの皮が剥がれた石になります。そして今後は、

[切り捨てられた]

## Original Extract

Website | Github | CC0 Licensed
This is a repository that aims to provide a collection of badges to symbolize that you didn’t use generative AI (notably LLMs) for the creation of your project. By using this, you are expected to ensure that AI-generated code makes up less than 1% of the total written lines of code in your project. You can use the badge only if you ask for assistance like a smart grepper.
Note that nobody will check your code , and if you are against AI use in your codebase but are unsure about the number of lines of code written by AI, you can still use the badge. The goal is to be transparent and to try to reduce the abusive use of AI in codebases.
This repository is accompanied by an explanation that tries to be backed by scientific research to support every claim. But this part is still a work in progress .
What is the ‘by-human’ expected use of AI ?
AI should not be used in the project to generate code. It can instead be used as:
A smart grep or a smart google . For example, fetching in a new codebase where something is implemented or maintained.
Reviewer helper . When having a lot of PRs having an AI to catch bug.
Learn a codebase (example: how do you do X ?) to learn by example or to explain something if there is no documentation. You should write the documentation after if you are maintaining the project.
Using AI to ask for help with a bug , but not to write the fix for yourself.
Asking how a feature could be implemented , but not to write the final implementation for yourself.
Every time you use an AI to help you think, it must be reviewed by a Human. You must also take responsibility of the elements the AI generated or told you.
Note that AI should be used sparingly and only after a Human was unable to provide a quick solution/alternative . You should always take into consideration the output of the AI and the consumption of energy & ecological footprint before using it.
You can copy paste code to your repository or directly use the logos in your project.
## Made by human
< a href = "https://github.com/Supercip971/by-human" >
< picture >
< source media = "(prefers-color-scheme: dark)" srcset = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-light.svg" >
< source media = "(prefers-color-scheme: light)" srcset = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-dark.svg" >
< img height = "96" align = "right" alt = "Made by humand, not by gen AI badge" src = "https://raw.githubusercontent.com/Supercip971/by-human/main/transparent-light.svg" >
</ picture >
</ a >
[Project name] is made by **Humans** , and not by a generative AI.
More information can be linked to the [ by-human ]( https://github.com/Supercip971/by-human ) repository.
Made by human
[Project name] is made by Humans , and not by a generative AI.
More information can be linked to the by-human repository.
Why do we think writing using generative AI is problematic?
Writing using generative AI is inferior in multiple ways. It may appear better and more efficient, but comes with important drawbacks that we enumerate here.
1. Copyright, plagiarism and license-washing
AI can license-wash unknowingly, and it has already been shown that it can be used to do it on purpose [1] .
TL;DR: Claude Code was used to ‘rewrite’ an entire project, which was then relicensed under the MIT instead of the LGPL.
This is a clear license violation.
If one were to claim that it was not, then would tanslating a book remove its copyright?
Accepting such a claim would effectively mark the end of copyleft and copyright protections.
When making an LLM learn, it is unable to grasp the license of the code.
As shown in two research papers, [2] [3] Large Language Models are generating 3.35% of strong copyleft licensed code and are:
not aware of reusing copyleft code and cannot be asked, through the prompt, to avoid reusing existing code in the responses.
This paper [2] also states that
accepting a copyleft request may lead to an increase in copyleft stolen code. (By a factor of 2 to 5).
Ultimately, LLMs are blatantly plagiarizing code. It is a far cry from a human learning about a chunk of code and then creating something based on obtained knowledge;
a human fundamentally understands the whole picture and the idea behind the code, and does not rely on rote memorization.
But when an LLM learns, it takes in a large quantity of code from a variety of sources and may spit it back out verbatim to oblivious users. Meaning that the strongly licensed part of the code is in its database, raising concerns about the respect of license.
This is different from a human importing a licensed piece of code, as a well-informed developer will include the license notice along with the chunk of code. In lieu LLMs don’t mention the license nor the author.
We will first let the numbers speak for themselves:
[4] Is stating that Chat-GPT 4 is instantly failing ~8% of the time, and code quality is good only ~60% of the time .
[5] code rabbit, an AI company, announces that LLMs introduce 70% major defect and 40% critical issues in pull request. Twice more than humans.
[6] this paper notify that Github-copilot has only superficial capabilities when trying to find bugs . It can hardly help discover bugs. Across multiple projects filled with vulnerability issues (more than 39 different types), GitHub Copilot only assisted in finding a couple, but in general it only fixed spelling mistakes. Finally, GitHub Copilot lacked any valuable comments
[7] Uplevel reports that generative AI is introducing 40% more bugs.
[8] puts into perspective the fact that we have a year over year increase of 40% of code pushed, then reverted or removed within 2 weeks . Meaning that since the introduction of AI, quickly ‘reverted’ code has increased from 3.97% to 7.09%.
Inexorably, we understand that AI generates incremental technical debt, that makes projects unmaintainable in the long term. And it’s barely able to fix itself.
It’s like a student writing code for you and not being able to learn and have cognitive introspection. And you, the programmer, are less likely to fully understand your code as you did not write it.
In some research [9] , it is shown that using AI makes you 48% to 127% more likely to achieve a better grade during practice problems . But in the end, you are 17% more likely to achieve a worse grade during the real test .
What is worse, is that those students were not able to realize that they learned less. And were unable to become more understanding.
This is a critical issue because you are leading to a false sense of knowledge. Generally you write code using AI and trust your competence by checking it. But you are becoming a really bad programmer by trusting the AI and not learning by yourself.
As you are expected to study your codebase, by using an LLMs, you are becoming worse at understanding your own codebase, thus worse at fixing and improving it.
Ultimately, this makes you more and more dependent on AI, and this will turn into a vicious cycle until your codebase is unable to be maintained.
3. Is the ecological aspect devastating?
It is rough to translate into numbers the ecological aspect of AI.
First, 70% [10] of ram production is dedicated to datacenters. A production increased by the reallocation of supplier capacity towards AI datacenters. Meaning that we are using a lot of economic resources to make AI run.
[11] Since the introduction of chat-GPT, the power consumption has elevated by 98% in one year. (2.69 MW in 2022 -> 5.43 MW in 2023).
The water usage is hard to put into perspective. The only trustable source is a citation from Sam Altman saying that Chat-GPT uses 0.000085 gallons of water per query. [12] but Chat-GPT processes 2.5 billion of request per day [13] meaning that on average Chat-GPT uses 804,400 L of water per day.
Thenceforward, this article tells us that [14] one Chat-GPT 4.5 request costs 20.500 Wh. But you can still not make this statement as clear as possible, as it uses an approximation.
It is more grounded as this article takes into account large context, because a lot of studies use ‘short’ requests. Although using an LLM as an agent requires it to read your file, your codebase, and can no longer be linked to a ‘short request’.
While it may seem a lot, those numbers are a ghost. We can’t make any further claim and are not able to put into perspective the direct ecological aspect of Chat-GPT usage. We would need a full research that is using OpenAI insights. On the other hand, as they are not releasing a lot of information we are stuck at guessing how much we are collapsing the world by using AI.
4.1 Inbreeding is as bad as it is for humans
Microsoft is training its LLMs on code from github, and they expressed in a conference that 40% of code written by an LLM is left unmodified [15] .
Although this quote is not really backed by any evidence, it is admitted to be true that more and more published code is written by an LLM, and it is progressively left untouched.
The training data of LLMs can’t differ between a human code and a LLM written code . Hinting that we will need more and more energy, training and data to accommodate this shift in quality.
An error just repeated multiple times by an LLM can become ground truth. It has been shown that only 20 documents can poison LLMs of any size [16] . (While this is not directly linked to this statement, this article shows how a couple of documents can shift an LLM’s point of view).
In summary, the easy shift in model knowledge coupled to the booming use of LLMs in the wild means that LLMs are now training on their own data, leading to a considerable decrease in training quality.
4.2 Compute power availability
Having to pay twice for your ram is a heavy cost of having datacenter eating the whole production.
AI is mainly able to evolve by multiplying compute power, RAM, context size…
Yet our world is unable to keep up with it [17] .
This paper crystallizes the concern with this statement:
Empirically, Sutton’s “bitter lesson” (Sutton, 2019) appears
partly incorrect: it is not that, for AI, “general methods
that leverage computation are ultimately the most effective,
[because of] Moore’s law, […] continued exponentially
falling cost per unit of computation”, but that increasing
resources are spent on AI. This increase in resources is
visible in computational costs but is also true of other costs.
For instance, building larger AI models require more human
labor. [17]
Subsequently, when we say an AI is getting better, it is not because of a ground breaking algorithm but rather:
More powerful hardware, which takes more money and consumes more energy
More data (which is becoming more and more polluted by LLMs)
More training, meaning more energy and more human labor.
We are reaching a point where we are sidelined to keep up with the increasing demand of resources, and the only way to keep up is to eat through the user market. The paper may have predicted the increase of recent ram price [18] .
When we will no longer have enough ram, no longer enough compute power, the whole AI industry may collapse, bringing us back to the point where we are now… And those who depended on AI will not be able to bring back their lost knowledge.
If you want to be positive, it may quickstart a thought of reversing computer evolution. And trying to become more efficient rather than more performant.
When writing code using AI, the process can seem almost like magic, because the problem may have already been solved by someone else and reproduced by the model.
It makes you look smarter but simultaneously makes you worse at solving problems on your own. AI is destroying the planet and our economy while getting worse, polluting its own data set. The situation is only worsening as investors are contributing billions, expecting more and more, while getting less and less in return.
It is only a temporary shiny rock that will become just a crusted rock. And hereafter,

[truncated]

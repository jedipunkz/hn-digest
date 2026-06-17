---
source: "https://writings.stephenwolfram.com/2026/06/launching-version-15-of-wolfram-language-mathematica-built-in-useful-ai-lots-of-new-core-functionality/"
hn_url: "https://news.ycombinator.com/item?id=48563609"
title: "Wolfram Language and Mathematica Version 15, AI Assistant, Symbolic Music, More"
article_title: "Launching Version 15 of Wolfram Language & Mathematica: Built-in (Useful) AI & Lots of New Core Functionality—Stephen Wolfram Writings"
author: "alok-g"
captured_at: "2026-06-17T01:02:36Z"
capture_tool: "hn-digest"
hn_id: 48563609
score: 40
comments: 3
posted_at: "2026-06-16T23:15:44Z"
tags:
  - hacker-news
  - translated
---

# Wolfram Language and Mathematica Version 15, AI Assistant, Symbolic Music, More

- HN: [48563609](https://news.ycombinator.com/item?id=48563609)
- Source: [writings.stephenwolfram.com](https://writings.stephenwolfram.com/2026/06/launching-version-15-of-wolfram-language-mathematica-built-in-useful-ai-lots-of-new-core-functionality/)
- Score: 40
- Comments: 3
- Posted: 2026-06-16T23:15:44Z

## Translation

タイトル: Wolfram言語とMathematica バージョン15，AIアシスタント，象徴音楽など
記事のタイトル: Wolfram言語とMathematicaのバージョン15の発売: 組み込みの（便利な）AIと多くの新しいコア機能—Stephen Wolfram Writings
説明: 組み込みAIアシスタント、拡張された計算拡張生成、WolframエージェントツールフレームワークなどのAI関連のアップグレード。 TimeSeries と表形式の進歩。新しいカテゴリデータ、ModelFit 自動モデル選択、シンボリック音楽。ギガバイトサイズのノートブックを効率的に処理します。での利益
[切り捨てられた]

記事本文:
×
すべてのリリース発表を見る »
内容
トップへ
現代に向けた印象的なリリース
すべてのノートブックに AI アシスタントを搭載
AI環境からWolframを使用する
時系列 (およびイベント シリーズ) が大きくなる
カテゴリカル データの計算が可能になる
ModelFit スーパー関数の紹介
表形式のより大規模で優れた接続性
ギガバイトサイズのノートブックとリアルタイム検索
ノートブックに初めてサイドバーが追加されました
ビジュアルテーマがノートブックに登場
長すぎると切れてしまいます
その計算では何が起こっているのでしょうか? Monitor の 1 引数形式
すぐに使用できる増分データ構造の導入
大規模なコードベースでの例外とエラー処理
構造化パッケージ形式の導入
地球の地図に目盛りを付けるにはどうすればよいですか?
あなたの街では、いつ日食が見られますか?
グラスマン、クリフォード、ワイル＆フレンズ
ゼータ、ポリログ、調和数が多変量化
部分分数の合理化
多数の新しい行列分解
DSolve のコーナー AI メソッドから少し助けを得る
PDE 解における導出量
システムエンジニアリングモデルをどのように近似しますか?
制御システムの強化学習
最新フォーマットのインポートとエクスポート
Web Socketsによるリアルタイム接続
ノートブックで Python などを使用するためのよりリッチな UX
最適化と GPU 化は続く
外部関数としての CUDA カーネル
Wolfram Compute ServicesがGPUを取得
LLM関数でのWolfram Foundationツールの使用
Stephen WolframとWolfram言語15を探索する
» (2026 年 6 月 16 日 @ 4:30 PM ET)
Wolfram言語とMathematicaのバージョン15のリリース：組み込みの（便利な）AIと多くの新しいコア機能
現代に向けた印象的なリリース
1988 年 6 月 23 日は、Mathematica のバージョン 1.0 がリリースされた日です。ほぼ 38 年が経った今日、私たちは、その方法を認識して、バージョン 15 をリリースします。

それは「数学」をはるかに超えて拡張されており、私たちは現在 Wolfram言語 と呼んでいます。これは、多くの新しいコア機能を備えた印象的なリリースです。 38 年も経った今でも、まだ追加すべきことがあるというのは、意外に思われるかもしれません。しかし、それは知的歴史の典型的な弧のようなもので、理解すればするほど、より遠くまで見ることができ、より多くのことができるようになります。そして、それに取り組んでいる私たち全員にとって、それは非常に満足のいくプロセスでした。年々、アイデアとテクノロジーのより高いタワーを構築し、それによって私たちはさらに遠くまで到達できるようになり、今日ではバージョン 15 のすべての機能に到達できます。
過去 40 年間、私たちには一貫した使命がありました。それは、計算パラダイムを可能な限り広く深く適用することであり、世界を表現し計算するための独自の計算言語を構築することでそれを実現することです。この 40 年間で、コンピューティングとコンピューティング パラダイムの使用は大きく広がりました。これは特に、私たちが導入したツールやアイデアのおかげだと思います。しかし今では、最新の AI という新しい推進力もあります。そして、AI の世界で予想外の進歩が起こっているのを見るのはとても楽しいことです。
私たちにとって直接的な影響の 1 つは、ユーザーのベースが単なる人間から人間と AI へと拡大したことです。そして，人間が使いやすく効率的に利用できるようにすることを目的として，Wolfram言語の一貫した設計に私たちが注いだ努力のすべてが，AIにとっても使いやすく効率的になっていることが判明した。
私たちは何年もの間、バージョン 1.0 のために発明したノートブックの概念から始まって、人間のユーザーのためのインターフェースに非常に重点を置いてきました。現在、私たちは AI と AI システム (そしてそれを使用する人間) ができるだけ簡単に当社のテクノロジーにアクセスできるように、AI 用のインターフェースにも重点を置いています。
私たちの技術

ogy は確かに AI にとって強力なツールです。しかし、これは AI を使用する人間にとっても強力なツールでもあります。それは、人間が物事を形式化し、何が言われているか、何が行われているかを正確に知るための独自の方法を提供するからです。私はWolfram言語の開発が計算パラダイムに対して何世紀も前に数学的記法が数学パラダイムに対して行ったことの拡張版であると常に見てきました：アイデアを表現し伝達するための合理的かつ正確な方法を提供するということです。
欲しいものを自然言語で AI に伝えるのは便利ですが、かなり単純な場合を除いて、非常に不正確です。しかし，AIがWolfram言語のコードを生成すると，AIが理解した内容が正確に表示され，それが本当に望むものであるかどうかを確認できるようになる。
ここでWolfram言語は独特の役割を果たします。従来のプログラミング言語は、人間が書き、コンピューターが読み取ることを目的としています。しかし，Wolfram言語はプログラミング言語を超えたものであり，本格的な計算言語です。これは人間が書くことだけでなく、考えを形式化し鮮明にする方法として人間が読むことも目的としています。そして現在、AI の時代では、計算パラダイムと世界を表現する計算方法を活用して、話している内容を正確に表現する独自の方法が使用されています。
はい、AI が常に物事を正しく理解するとは限りません。しかし重要なのは，Wolfram言語を精度（と正しさ）の伝達手段として，そして自分のやっていることを定着させて体系的な方法で自信を持って使用できる確かな出力を生成する方法として使うことである。
特に今年は、「コーディングに AI を使用する」という大きなトレンドがあります。そして、はい、「正しく見えること」が目的である何か (Web サイトなど) を制作したい場合は、

「コードが内部で何をしているか」は気にしません。これは優れた、実際に非常に変革的なソリューションです。しかし、特により技術的な分野では、「正しく見える」だけでは不十分な状況が多くあります。実際に何が計算されているかを知る必要があります。ここでWolfram言語が重要になります。それは、何が行われているかを最高レベルで最も人間が理解できる表現を提供するものだからです。また、計算の正確な部分をカプセル化して、必要な場所で繰り返し使用する方法を提供します。
コーディングにおける現代の AI の成功は驚くべきものであり、予想外です。しかし、ある意味、それは私たちにとって、たとえば従来のプログラミング言語ほど重要ではありません。なぜなら、計算の仕様と実行を可能な限り自動化することが、数十年にわたる私たちの使命だからです。その結果、計算の世界をカバーする 7000 以上のプリミティブが生まれ、非常に簡潔に驚くべき範囲のものを表現できるようになりました。
実は私は何十年も前から，従来のプログラミングの多くはWolfram言語の高レベル構造を使うだけで自動化できると言い続けてきた．そして実際に（私を含む）非常に多くの人が何年もの間Wolfram言語を使用して計算範囲を劇的に増やし，従来のプログラミング言語コードを大量に書くことを避けてきました。
しかし現在、AI は別の道を提供しており、大量の従来のプログラミング言語コードを自動的に記述します。はい、完全に信頼できるわけではなく、通常、軌道に乗るようにするには非常に高度なラングリングが必要です。しかし、少なくとも、自分が何をコンピューティングしているのかを正確に気にしないのであれば、それは自動化への貴重な道を提供します。
Wolfram言語を使い始めたばかりの人，またはWolfram言語で働いている人向け

AI は、初期自動化の便利な層を提供します。しかし，Wolfram言語を流暢に扱えるとしても，それは通常は望んでいることではありません． Wolfram言語は考えるための媒体を提供します．そして，この言語を流暢に使えるようになると，通常，自分の考えを通常の自然言語で最初に言語化するよりも簡単にその言語で直接表現できるようになります． (何かに取り組んでいるときは，やりたいことを自然言語で少なくとも正確に記述するよりもずっと早くWolfram言語のコードを入力し始めることができることを私は知っています。）
Wolfram|Alpha はすでに 17 年以上前に自然言語を使用して計算を指定するというアイデアを先駆者としていたことは言及する価値があります。これは現代の AI とは異なるテクノロジーであり、自然言語の小さな断片をより重視しており、正確な計算に対する信頼性の高い翻訳を備えています。しかし，それによってすでに何年も前から，たとえばエンティティの指定などにWolfram言語内で自然言語を利用できるようになった．そして今では、AI とのより良いコミュニケーション チャネルの構築にも役立っています。
ここ数カ月間、ソフトウェア開発の将来における AI の役割について多くのことが言われてきました。それでは、それは私たちの活動やバージョン 15 などの開発にどのような影響を与えるのでしょうか?確かに、特に従来のプログラミング言語を使用するシステムの部分 (通常は外部インターフェイスやハードウェアとの直接対話に関係する部分) を扱う場合には、それが役立つ場面はあります。しかし，Wolfram言語のコードの大部分は現在Wolfram言語で書かれており，言語に組み込まれている自動化機能はすべてすでに活用されている．言語の新しいバージョンが登場するたびに、より多くの機能が自動化され、より活用できるようになります。

さらなる発展。そして実際、それこそが、私たちが過去 40 年間にわたって築き上げてきた驚くべきテクノロジーの塔を可能にしたのです。そして今日、バージョン 15 が提供されます。
すべてのノートブックに AI アシスタントを搭載
ChatGPTの最初のリリースから数週間以内に、私たちはLLM内からWolfram言語(およびWolfram|Alpha)を呼び出す方法、およびWolfram言語(およびWolfram Notebooks)内からLLMを呼び出す方法を構築しました。翌年、私たちはWolframシステムのアドオンとしてNotebook Assistantをリリースできる技術を構築しました。そして今年 2 月に、LLM とさらに統合した Foundation Tool テクノロジー スイートをリリースしました。現在、バージョン 15 では、別のレベルの AI 統合、つまり組み込み AI アシスタントを開始します。
バージョン 15 で新しいノートブックを作成すると、(オフにしていない限り) ノートブックの下部に、AI アシスタントにすぐに接続できる「チャットバー」と呼ばれる新しい要素が表示されます。
チャットバーに必要な内容を入力します (画像などを貼り付けることもできます)。 Enter キーを押すと、入力が AI アシスタントに送信され、AI アシスタントが支援を試みます。
かなり漠然とした質問をしたとしても、AIアシスタントは、可読性の高いWolfram言語コードを備えた正確な解釈を最善の推測で提供します。押すとコードがノートブックに挿入され、すぐに実行します。
チャットバーは、ノートブックにチャット セルを作成するための、いつでも利用できる便利な方法と考えることができます。バージョン 14.2 からできるようになったように、「」と入力して新しいセルを開始するだけでチャット セルを作成することもできます。
他のチャット セルと同様に、チャットバーから作成されたチャット セルは、ノートブック内でその上にあるコンテンツのコンテキストを利用できます。 (文脈を壊すために、チャットデリを挿入できます)

セル間に「~」を入力してマイターします。)
しかし、バージョン15のより大きな話は、チャットバーとチャットセルの背後にあるAIアシスタントがすべてのWolfram Notebookですぐに利用できるようになったということです。設定は必要ありません。 AI アシスタントの基本レベルについても、追加のサブスクリプションは必要ありません。
AIアシスタントの基本レベルは，Wolfram言語で何かを行う際のドキュメントを超えた方法としてすぐに役に立ちます．また、サブスクリプションで利用できる、Pro と Research という 2 つのより高いレベルの AI アシスタントも本日リリースします。 Pro では、より大規模で洗練されたプロジェクトに取り組むことができ、Research では最新のフロンティア AI 機能へのアクセスが提供されます。 (既存の Notebook Assistant ユーザーは自動的に AI Assistant Pro に移行されます。)
AI アシスタントのコントロールにアクセスするには、チャットバーの横にある をクリックするだけです。
デフォルトでチャットバー全体を表示したくない場合は、 をクリックすると最小化されます。
(最小化された状態は、新しいノートブックを開いた場合に記憶されます。メインの [環境設定] メニューで、チャットバーを表示するかどうか、さらには AI アシスタントを使用できるかどうかをグローバルに制御できます。)
AI環境からWolframを使用する
AI アシスタントを使用すると、Wolfram Notebooks 内から AI にアクセスできます。しかし、コンピューター上で Claude Code や Codex などの AI 環境を使用しているとします。バージョン 15 では、AI 環境を簡単に接続できるようになりました。

[切り捨てられた]

## Original Extract

AI-related upgrades like the built-in AI Assistant, extended computation-augmented generation and the Wolfram Agent Tools framework. Advances in TimeSeries & Tabular. New categorical data, ModelFit automated model selection, symbolic music. Efficient handling of gigabyte-sized notebooks. Gains in th
[truncated]

×
View All Release Announcements »
Contents
Top
An Impressive Release for Modern Times
An AI Assistant in Every Notebook
Use Wolfram from Your AI Environment
Time Series (and Event Series) Go Big
Computation Comes to Categorical Data
Introducing the ModelFit Superfunction
Bigger and Better Connectivity for Tabular
Gigabyte-Sized Notebooks and Real-Time Find
Notebooks Get Their First Sidebars
Visual Themes Come to Notebooks
When It’s Too Long, It’s Torn Off
What’s Happening in that Computation? The One-Argument Form of Monitor
Introducing Ready-to-Use Incremental Data Structures
Exceptions and Error Handling in Large Codebases
Introducing the Structured Package Format
How Do You Put Ticks on a Map of the Earth?
When Will Your City See a Solar Eclipse?
Grassmann, Clifford, Weyl & Friends
Zetas, Polylogs and Harmonic Numbers Go Multivariate
Partial Fractions Get Streamlined
Lots of New Matrix Decompositions
The Corners of DSolve Get a Little Help from AI Methods
Derived Quantities in PDE Solutions
How Do You Approximate a Systems Engineering Model?
Reinforcement Learning for Control Systems
Importing & Exporting the Latest Formats
Real-Time Connection with Web Sockets
Richer UX for Using Python & More in Notebooks
Optimization & GPUification Continues
CUDA Kernels as External Functions
Wolfram Compute Services Gets GPUs
Using the Wolfram Foundation Tool in LLM Functions
Exploring Wolfram Language 15 with Stephen Wolfram
» (June 16, 2026 @ 4:30 PM ET)
Launching Version 15 of Wolfram Language & Mathematica: Built-in (Useful) AI & Lots of New Core Functionality
An Impressive Release for Modern Times
June 23, 1988 is when we launched Version 1.0 of Mathematica . Today—almost 38 years later—we’re launching Version 15 of what—in recognition of how far it’s expanded beyond “math”—we now call Wolfram Language . It’s an impressive release, with a lot of new core functionality. It might perhaps seem surprising that after 38 years there’d still be more to add. But it’s like the typical arc of intellectual history: the more one’s figured out, the further one can see, and the more one becomes able to do. And for all of us working on it, it’s been a very satisfying process: year after year building an ever taller tower of ideas and technology, with which we can reach ever further—today to all the functionality of Version 15.
For the past four decades we’ve had a consistent mission: to apply the computational paradigm as broadly and deeply as possible—and to do so by building our unique computational language to represent and compute about the world. Over these four decades the use of computation and the computational paradigm has spread greatly—not least, I think, as a result of tools and ideas we’ve introduced. But now there’s also a new driver: modern AI. And it’s been exciting to see so much unexpected progress happen in the world of AI.
For us, one of the immediate consequences has been that our base of users has expanded from just humans, to humans and AIs. And it’s turned out that all the effort we put into the coherent design of the Wolfram Language—aimed at making it easy and efficient for humans to use—now also makes it easy and efficient for AIs.
For years we’ve put great emphasis on interfaces for human users, starting from the concept of notebooks that we invented for Version 1.0 . Now we’re also putting emphasis on interfaces for AIs, to make it as easy as possible for AIs and AI systems (and the humans who use them) to have good access to our technology.
Our technology is certainly a powerful tool for AIs. But it’s also a powerful tool for humans using AIs. Because it provides a unique way for humans to formalize things , and know exactly what’s being said, or done. I’ve always seen the development of Wolfram Language as doing for the computational paradigm an extended version of what mathematical notation did centuries ago for the mathematical paradigm: providing a streamlined and precise way to represent and communicate ideas.
When you tell an AI in natural language what you want, it’s convenient, but—except in rather simple cases—quite imprecise. But if the AI generates Wolfram Language code, then that shows you in precise terms what the AI understood, and allows you to see whether it’s really what you want.
The Wolfram Language has a unique role here. Traditional programming languages are intended as something humans write, and computers read. But the Wolfram Language is something beyond a programming language—it’s a full-scale computational language . That’s intended not just to be written by humans, but also to be read by them, as a way to help formalize and crispen up their thoughts. And now, in the time of AI, it’s a unique way to represent precisely what one’s talking about—leveraging the computational paradigm, and the computational way of representing the world.
Yes, AIs don’t always get things right. But the point is to use Wolfram Language as a carrier of precision (and correctness)—and as a way to anchor what one’s doing, and generate solid output that one can confidently use in systematic ways.
There’s been a big trend—particularly this year—to “use AI for coding”. And, yes, if you want to produce something (like a website) where “looking right” is the objective—and you don’t care “what the code is doing inside”—it’s a good, and in fact quite transformative, solution. But there are many situations, particularly in more technical areas, where “looking right” isn’t good enough: you need to actually know what is being computed. And that’s where the Wolfram Language is crucial. Because it’s what gives you the highest level, and most human-understandable representation of what’s being done. And gives you a way to encapsulate a precise piece of computation to repeatedly use wherever you want.
The success of modern AI in coding is remarkable, and unexpected. But in a sense it’s much less significant to us than it is, say, for traditional programming languages. Because it’s been our mission for decades to automate as much as possible of the specification and doing of computation. And the result has been 7000+ primitives that cover the computational world—and that allow one with great succinctness to represent a remarkable range of things.
I’ve actually been saying for decades that much of traditional programming can be automated, just by using the higher-level constructs of the Wolfram Language. And indeed a great many people ( including myself ) have used Wolfram Language for years to dramatically increase their computational reach, and avoid writing large volumes of traditional programming language code.
But now AI provides a different path—where it automatically writes those volumes of traditional programming language code. Yes, it’s not perfectly reliable, and typically requires quite sophisticated wrangling to keep it on track. But at least if one doesn’t care exactly what one’s computing, it provides a valuable path to automation.
For people just starting to use Wolfram Language—or working in an area they’re not familiar with—AI provides a convenient layer of initial automation. But if one’s fluent in Wolfram Language, it’s typically not what one wants. The Wolfram Language provides a medium to think in. And as soon as one’s fluent in it, one can typically express one’s thoughts more easily directly in the language than one can first verbalize them in ordinary natural language. (I know that when I’m working on something I can much more quickly start typing Wolfram Language code than I could ever describe what I want to do, at least with any precision, in natural language.)
It’s worth mentioning that Wolfram|Alpha already pioneered the idea of using natural language to specify computation more than 17 years ago. It’s a different technology than modern AI—more oriented to small fragments of natural language, with reliable translation to precise computation. But it already allowed us many years ago to take advantage of natural language within Wolfram Language, say for specifying entities. And now it also helps us in building a better communication channel with AIs.
In recent months much has been said about the role of AI in the future of software development. So how does it affect what we do, and the development of things like Version 15? Well, there are certainly places where it’s helpful, particularly in dealing with the parts of our system (typically concerned with external interfaces or direct interaction with hardware) that use traditional programming languages. But most of the code of the Wolfram Language is now written in the Wolfram Language—where we already take advantage of all the automation that’s built into the language. With every new version of the language, there’s more that has been automated, and more leverage in doing more development. And indeed that’s what’s made possible the remarkable tower of technology that we’ve built over the past four decades. And that today brings us Version 15.
An AI Assistant in Every Notebook
Within weeks of the original release of ChatGPT we’d built ways to call Wolfram Language (and Wolfram|Alpha ) from inside LLMs—and to call LLMs from within Wolfram Language (and Wolfram Notebooks ). The next year we’d built the technology that made it possible for us to release our Notebook Assistant as an add on to the Wolfram System. Then in February of this year we released our Foundation Tool technology suite, further integrating with LLMs. Now, in Version 15, we’re launching another level of AI integration: our built-in AI Assistant .
Create a new notebook in Version 15 and (unless you’ve switched it off) you’ll see at the bottom of the notebook a new element that we call a “chatbar” that connects you immediately to our AI Assistant:
Type what you want into the chatbar (you can also paste images, etc.). Then just hit ENTER and your input will be sent to the AI Assistant, which will try to help you with it:
Even if you ask something fairly vague, the AI Assistant will give its best guess of a precise interpretation, complete with readable Wolfram Language code. Press and the code will be inserted into your notebook, and then immediately run:
You can think of the chatbar as a convenient always-available way to create a chat cell in a notebook. As you’ve been able to do since Version 14.2 , you can also create a chat cell just by typing ‘ to start the new cell.
Just as with any chat cell, chat cells created from the chatbar can make use of the context of content above them in the notebook. (To break the context you can insert a chat delimiter by typing ~ between cells.)
But the bigger story in Version 15 is that the AI Assistant behind the chatbar and chat cells is now immediately available in all Wolfram Notebooks. No configuration is needed. And for the Basic level of the AI Assistant, no additional subscription is needed either.
The Basic level of the AI Assistant is immediately useful as a beyond-the-documentation way to get help in doing things with Wolfram Language. We’re also releasing today two higher levels of the AI Assistant, available by subscription: Pro and Research. Pro lets you tackle larger and more sophisticated projects, and Research provides access to the latest frontier AI capabilities. (Existing Notebook Assistant users will automatically be transferred to AI Assistant Pro.)
To get to the controls for the AI Assistant just click the at the side of the chatbar:
If you don’t want to see the full chatbar by default, click the and it will be minimized:
(The minimized state will be remembered if you open a new notebook. You can globally control whether the chatbar appears—and even whether the AI Assistant is available at all—in the main Preferences menu.)
Use Wolfram from Your AI Environment
The AI Assistant lets you access an AI from within Wolfram Notebooks. But let’s say you’re using an AI environment—like Claude Code or Codex—on your computer. In Version 15 it’s now easy to hook up your AI environment to

[truncated]

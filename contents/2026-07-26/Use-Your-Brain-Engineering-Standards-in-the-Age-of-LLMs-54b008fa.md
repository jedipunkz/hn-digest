---
source: "https://pgaleone.eu/ai/2026/07/26/use-your-brain/"
hn_url: "https://news.ycombinator.com/item?id=49058440"
title: "Use Your Brain: Engineering Standards in the Age of LLMs"
article_title: "Use Your Brain: Engineering Standards in the Age of LLMs – P. Galeone's blog"
author: "me2too"
captured_at: "2026-07-26T14:27:56Z"
capture_tool: "hn-digest"
hn_id: 49058440
score: 1
comments: 0
posted_at: "2026-07-26T14:12:58Z"
tags:
  - hacker-news
  - translated
---

# Use Your Brain: Engineering Standards in the Age of LLMs

- HN: [49058440](https://news.ycombinator.com/item?id=49058440)
- Source: [pgaleone.eu](https://pgaleone.eu/ai/2026/07/26/use-your-brain/)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T14:12:58Z

## Translation

タイトル: 頭を使う: LLM 時代のエンジニアリング標準
記事のタイトル: Use Your Brain: LLM 時代のエンジニアリング標準 – P. Galeone のブログ
説明: ソフトウェア開発はもはやニッチなものではありません。 LLM の出現により、誰もが独自のソフトウェアを作成できるようになりました。これは良いことです。しかし、私は膨大な数のソフトウェア エンジニア (本物です!) がバイブコーディングの罠に陥っているのを目にしています。あの人たちはデザインに慣れていたんです - 考えてください! -そして
[切り捨てられた]

記事本文:
P.ガレオーネのブログ
私について
会談
連絡してください
購読する
頭を使いましょう: LLM 時代のエンジニアリング標準
ソフトウェア開発はもはやニッチなものではありません。 LLM の出現により、誰もが独自のソフトウェアを作成できるようになりました。これは良いことです。しかし、私は膨大な数のソフトウェア エンジニア (本物です!) がバイブコーディングの罠に陥っているのを目にしています。あの人たちはデザインに慣れていたんです - 考えてください! - そして後で実装します。今では代わりに、彼らは大量の疑わしい品質のコードを作成しているだけです。これは、資格のあるソフトウェアエンジニアではなく、平均的なジョーに期待されるものです。
明確にしておきますが、私はここで私自身の哲学を共有しているのであり、福音を説いているわけではありません。私のワークフローを採​​用するよう誰かを説得しようとしているわけではありません。
LLM は光の速さでコードを生成し、処理してトークン 1 を生成し、要求したものとほぼ同じものを提供します。コードベースのサイズとステータスに応じて、LLM の使用は正しくガイドされる必要がありますが、企業や個人開発者は「迅速に提供する」以外の明確なガイドラインに従っていません。
自分で独自のスクリプトを開発し、アイデアから始めて数分で動作するのが確認できる場合に限り、迅速な提供は完璧です。
あなたが開発したツールがガレージから出て、他の人が使用し始めると、迅速な提供が問題になります。すぐに問題が発生し始めるので、修正する必要があります。 LLM に修正を依頼することもできますが、これはまさにコードベースの制御を失うことになります。
コードの所有権の概念について少し考えてみましょう。
完全に生成されたコードベースは本当にあなたのコードベースですか?それを理解するための経験則があります。
コードがどのように機能するかを説明し、自分でバグを特定できれば、それはあなたのコードベースです。しかし、編集のたびに LLM に依存していては、マナではありません。

ソフトウェアを使用する - AI に、理解できなくなったランダムなコードを操作させているだけです。
あなたが取り組んでいるコードベースがあなたのものではなく、何年にもわたって複数の開発者によって開発されている大規模な作業である場合でも、できるだけ速く提供することが目標です (唯一の目標ではありません。私はいかなる犠牲を払ってもスピードを重視するわけではありません。私は、最初の努力とその後の着実な強化を伴う、よく設計されたソリューションを好みます)。 AI の助けを借りてできるだけ早く開発できますが、コミットされたすべてのこと、潜在的な副作用、最新の状態に保つべきテスト、尊重すべきベスト プラクティスについては細心の注意を払う必要があります。最後に、LLM がコードの一部またはすべてを生成したとしても、LLM が何をしたかを知ることができるはずです。逆に、未解決の問題に対して LLM を適用するだけで、安定した必須のソフトウェア開発プロセスがない場合、このコードベースは急速に制御不能になり、何が起こっているのか誰もわかりません。
できるだけ短い時間でできるだけ多くの価値を提供するよう求められると、誰もがカウボーイコーダーになってしまいます。すべての優れたソフトウェア開発手法は、価値をもたらすものとしてではなく、問題として見なされます (ここに罠があります)。結局のところ、本当に重要なのは (そして、病理学的組織においては、何を測定し、良いものとして評価するのか) 開発スピードですよね。
速く動きたいだけですか?それとも、保守可能でスケーラブルで安全な、実際に所有するものを構築したいですか?
あなたが純粋なカウボーイであれば、物事をできるだけ早く終わらせることだけを気にし、最初の PoC を行った後はコードのメンテナンスを他の人に委任するかもしれません。それは理解できますが、それはスケーラブルではなく、カウボーイが開発したコードベースを継承する人はそのコードを見て疑問に思うでしょう。

それを最初から書き直す方が理にかなっていて、（実際にはあなたのものではない）設計上の決定をあなたのせいにするのであれば。
カウボーイであるあなたには所有権がないため、所有者ではなく管理者として行動します。あなたが所有者である場合と管理者である場合では、メンタル モデル全体と AI との関係が変化します。カウボーイたちはコードベースを多かれ少なかれブラックボックスとして見ており、何が入って何が出てくるかだけは知っていますが、内部ロジックは曖昧です。彼らは、設計から実行、デバッグに至るまでの重労働を LLM に依存しています。
面白いのは、バグが発生したときです。バグは自分自身 (独自の?) コードをデバッグできず、LLM がコードを修正する方法を知っていることを祈ることしかできません。
もちろん、核となる価値は生のスピードと即時出力 (速い、速い、速い!) です。
カウボーイになりたい人は誰ですか？ソフトウェアエンジニアである私はそうではありません。私自身とすべてのソフトウェア エンジニアは、管理者ではなく所有者であると考えています。所有者のメンタル モデルは異なります。彼らは頭の中にシステムの完全なマップを持っており、特定のアーキテクチャが選択された理由と、エッジ ケースがどこに隠れているかを知っています。モジュール内の何かを変更すると、他の依存モジュールにどのような影響を与えるか。彼らと AI との関係もより健全です。彼らは LLM を生産性や速度のブースターとして使用しています。すべての出力がレビューされ、コミットされたコードが 1 行ずつ説明されます。バグが発生した場合、問題の原因となっている可能性があるコードの領域を特定し、修正のために何を変更すべきか、その変更がコードベース全体にどのような影響を与えるかを把握しながらタイムリーに行動できます。
組織、および個人プロジェクトであっても一般的な AI の使用は、開発者に急速に急速に進むよう促しており、バイブコーディングの罠に陥る重大なリスクを伴います。彼らは、自分たちの役割を所有者から管理者に静かに格下げし、自分たちの方が早く行動していると考えていますが、実際には

システムの制御を AI に引き渡す - 技術的負債が戻ってきて彼らを苦しめるまで。
AIの正しい活用：インフラとネチケット
現実離れした話ですが、個人開発者や企業 (!) は、コーディング、特に AI 支援コーディングに関しては基本に立ち返り、非常に強力なエンジニアリング プラクティスとガードレールを設定する必要があります。
何かに取り組みたい場合:
人々は共通の目標に向かって努力します。
コミットされたコードの実際の所有権は人々にあります。
エージェントと従業員は、コードの品質、安全性、長期的な保守性が保証される構造化された環境でシームレスに作業します。
正しい環境をセットアップする必要があります。この環境はすでに存在していると思いますが、会社レベルで強制し、個人開発者が適用する必要があります。ルールは非常に簡単です。
ルールを定義する: 人間の役割、エージェントの役割、各人の責任を定義し、AI コーディングの実践を規定します。カウボーイを禁止しろ。
開発者とエージェントがルールを尊重し、標準に従うことをサポートするツールをセットアップします。
CI/CD を通じてすべてを強制します。
人は人と話すべきです。
AI 支援コーディングは、適切な所有権をもって行う必要があります。
エージェントが生成したコードは人間によるレビューが必要です - 必須です。
人間が書いた問題に生の AI 出力で応答することは禁止されています。
すべての企業とすべての開発者には独自のルールと基準があるため、私は意図的に一般的な言い方をしていますが、それらの中には新しいネチケットのようなものもあります。
私たちは文字通り車輪の再発明を行っており、80 年代に人々が定義した古い習慣に戻りつつあります。初期のコンピューター ネットワークと Usenet で生まれたネチケットは、今日ではこれまで以上に意味を持っています。私が AI の戯言で書いたことに誰かが答えたとき、私は本当に腹が立ちます。時間をかけて書いたのに、絶対にやめてください

時間をかけて読んで、適切な答えを書いてください。頭を使ってください。
ツール部分については、私の提案は明白で、強力な CI/CD です。リンター、フォーマッタ、事前コミットを使用し、テスト スイートが人間によってレビューされ、テストが正しいことを確認します。さらに、明確なガードレールを備えた厳密に定義された AGENTS.md ファイルである AI カスタマイズを活用します。エージェントの思考とツールの使用プロセスをガイドするために、人間が作成した小さなスキルが続きます。
AI は生産性を向上させるものであり、増幅器です。エンジニアリングの良い習慣を広めることで、品質を損なうことなく速度が向上し、コードベースに対する実際の所有権が保証されます。怠惰を増幅させると、あなたの（またあなたの？）リポジトリがノイズで満たされるだけで、あなたは短期間で所有権のない管理者になってしまうでしょう。
ツールを活用してワークフローを高速化しますが、エンジニアリング上の判断を決して放棄しないでください。
最近では誰もがトークンについて話しています。トークンが、TF-IDF などを実行するときに使用される決定論的な手順である「トークン化」の結果にすぎなかったときのことを覚えています。すべてのトークンは単語、ピリオドでした。代わりに、トークンは「トークン」であり、トークンが何であるかはすべてのモデルが決定するため、標準的な定義や固定ルールはありません。ローカルにデプロイされたモデルがある場合は、トークンが何であるかを (多かれ少なかれ) 制御できます。モデルがクラウド上で実行される場合、モデルプロバイダーがトークンとして販売したいのはトークンです。 ↩
Curl を使用したソフトウェアのインストール | sh は悪い習慣です - 残念ながら今日ではよくあります。これは、マルバタイジングを通じて悪用される攻撃ベクトルです。この記事では、誰かがスポンサー付き Web サイトからツールを盲目的にインストールすると何が起こるかについて説明しています。
あなたのデジタルライフが完全に外国勢力に依存したらどうなるでしょうか? 1656 年のスピンの破門との類似点を描くことで

オザと現代のアメリカによるヨーロッパ国民への制裁について、この記事では「デジタル破門」の概念と、ヨーロッパの技術主権がもはや任意ではない理由について論じています。
Gemini を活用してイタリアの金融ニュース フィードを解析し、リアルタイムの取引推奨を提供する自動株分析システムを構築した方法。この記事では、AI を活用したニュース分析を Go ベースの取引システムに統合する際のアーキテクチャ、課題、実装の詳細について説明します。
Python および Go 開発者向けの Vertex AI SDK から Google Gen AI SDK への完全な移行ガイド。サービス アカウント認証、OAuth2 スコープの制限、Google の公式ドキュメントに記載されていない重要な実装の詳細について説明します。
Google Cloud から OVH へのウェブサービスの移行に関する詳細なチュートリアル。PostgreSQL データベースの移行、Github Actions での CI/CD パイプラインのセットアップ、クラウドからセルフホスト ソリューションへの移行による大幅なコスト削減について説明します。この移行は、サービス品質を維持しながら米国のクラウド プロバイダーへの依存を軽減するための第一歩となります。
Cline のような AI ツールを活用して、Web サイトの UI/UX を強化し、バックエンド タスクを合理化した方法。ページの再デザインやコンテンツの翻訳から、AI 支援開発のメリットと課題の解決に至るまで、このブログ投稿では、学んだ重要な教訓を共有しながら、大規模な言語モデルを使用して生産性を向上させる可能性を強調しています。
Unreal Engine 5.3 以降、Epic Games はいわゆる最新の Xcode ワークフローのサポートを追加しました。このワークフローにより、Unreal Build Tool (UBT) が標準の Xcode アプリ プロジェクトとより一貫性を保ち、アプリケーションを配布するための Apple 要件に準拠できるようになります...理論上は! 😅 実際には、このワークフローには欠陥があります。コードサインインの両方に欠陥があります。

g とフレームワークのサポートが正しく実装されていないため、動作するアプリの作成とその配布が不可能になります。この記事では、macOS 上の Unreal Engine アプリケーションのパッケージ化、コード署名、公証の際に直面する問題を取り上げ、最後にそれらすべてを解決するための段階的なプロセスを説明します。
「Go で Vertex AI を使用した Google Cloud でのカスタム モデルのトレーニングとデプロイ」の記事では、Go を活用してリソース プールを作成し、Vertex AI に割り当てられたリソースを使用して機械学習モデルをトレーニングする方法について説明しました。このアプローチには柔軟性がありますが、リソース プールのコストへの影響という考慮すべき重要な側面があります。この記事では、Vertex AI の突然の価格上昇と、一見無害に見えるリソース プールという隠れた原因に関する私の経験について詳しく説明します。
この記事では、大規模言語モデル (LLM) とリレーショナル データベースを組み合わせて、ユーザーが自然な方法でデータについて質問できるようにする方法を検討します。データの保存と取得に PostgreSQL と pgvector を利用する、Go で構築された取得拡張生成 (RAG) システムを示します。提供されたコードは、コア機能を示しています。これは、fitsleepinsights.app の「データとチャット」機能がどのように開発されているかの概要です。

[切り捨てられた]

## Original Extract

Software development is not a niche anymore. The advent of LLMs made it possible for everyone to write their own software - and this is a good thing. However, I see a huge number of software engineers (the real ones!) falling into the vibe-coding trap. Those people were used to design - think! - and
[truncated]

P. Galeone's blog
About me
Talks
Contact me
Subscribe
Use Your Brain: Engineering Standards in the Age of LLMs
Software development is not a niche anymore. The advent of LLMs made it possible for everyone to write their own software - and this is a good thing. However, I see a huge number of software engineers (the real ones!) falling into the vibe-coding trap. Those people were used to design - think! - and later implement. Now, instead, they are just producing a huge amount of dubious quality code, something that I would expect from the average Joe, not from a qualified software engineer.
To be clear: I’m sharing my own philosophy here, not preaching a gospel. I’m not trying to convince anyone to adopt my workflow.
LLMs produce code at the speed of light, they crunch and generate tokens 1 and give you pretty much what you asked for. Depending on the codebase size and status, the usage of LLMs should be correctly guided - and companies and solo-developers are not following any clear guideline apart from “deliver fast”.
Delivering fast is perfect, if and only if, you are developing your own script, for yourself, starting from an idea and seeing it working in minutes.
Delivering fast becomes a problem when the tool you developed exits your garage and starts being used by others. Immediately, issues start to pop up and you have to fix them. You can ask an LLM to fix them - and this is precisely how you’ll lose control of your codebase.
Let’s think about the concept of code ownership for a second.
Is a fully generated codebase really your codebase? I have a rule of thumb for understanding it.
If you can explain how the code works and pinpoint bugs on your own, it’s your codebase. But if you rely on an LLM for every single edit, you aren’t managing software - you’re just letting an AI manipulate random code you no longer understand.
When the codebase you’re working on is not yours and it is a massive work being developed by several developers over the years - delivering as fast as possible is still a goal (not the only one. I’m not a fan of speed at all cost. I prefer well designed solutions, with initial effort and a steady ramp up afterwards). You can develop as fast as possible with the help of AI - but you need to be very careful about everything that’s committed, potential side effects, tests to be kept up-to-date and best practices to respect. At the very end, even if the LLM generated part or even all the code, you should be able to tell what the LLM did. If, instead, you just throw an LLM over any open issue and no software development process is stable and mandatory, then this codebase goes out of control quite fast, and no-one knows what’s happening anymore very quickly.
Being pushed to deliver as much value as possible in the shortest time possible, makes everyone a cowboy coder - every good software development methodology is seen as a problem, not as something that brings value (here’s the trap). After all, what really matters (and - in pathologic organizations - what is measured and evaluated as a good thing) is the development speed, isn’t it?
Do you want to move fast only ? Or do you want - also - to build something maintainable, scalable, secure and that you really own?
If you’re a pure cowboy, you just care about getting things done as fast as possible, and perhaps delegate the maintenance of your code to someone else after doing the first PoC - understandable, but it’s not scalable and whoever inherits your cowboy-developed codebase is going to look at that code, wondering if rewriting it from scratch makes more sense, and blaming you for your (not really yours) design decisions.
Being a cowboy you have no ownership, you are acting as a custodian rather than an owner. The whole mental model and the relationship with AI changes when you’re an owner vs a custodian. The cowboys see the codebase as a black box more or less, they just know what goes in, what comes out, but the internal logic is blurred. They rely on the LLM to do the heavy lifting, from design to execution and debugging.
The fun part is when bugs happen: they can’t debug their own (own?) code, and they can only hope the LLM knows how to fix the code.
Of course the core value is raw speed and immediate output (fast fast fast!).
Who wants to be a cowboy? I, as a software engineer, don’t. I see myself, and all the software engineers, being owners rather than custodians. The owner mental model is different: they have a complete map of the system in their head, they know why a specific architecture was chosen and where the edge cases hide. How changing something in a module affects other dependent modules. Their relationship with AI is more healthy too: they use the LLM as a productivity/speed booster. Every output is reviewed and the committed code can be explained, line by line. In case of bugs, they can pinpoint the area of the code that is likely to be responsible for the issue, and act timely, knowing what to change for fixing, and how the changes affect the entire codebase.
Organizations, and the general AI usage at all costs even in solo-projects, are pushing developers to move fast-fast-fast, with the serious risk for them to fall into the vibe-coding trap. They quietly downgrade their role from owners to custodians, thinking they are moving faster, but in reality surrendering control of the system to an AI - until the technical debt comes back to bite them.
Correctly Leveraging AI: infrastructure and netiquette
It’s surreal, but solo-developers and companies (!) should go back to the basics and set up very strong engineering practices and guardrails when it comes to coding - especially AI-assisted coding.
If you want to work on something where:
People work towards a shared goal.
People have a real ownership of the committed code.
Agents and people work seamlessly in a structured environment that guarantees code quality, safety, and long-term maintainability.
The correct environment should be set up. I think that this environment already exists, but it should be enforced company-level and applied by solo-developers. The rules are quite simple:
Define the rules: define the human roles, the agents roles, the responsibilities of every person and rule the AI coding practices. Ban the cowboys.
Set up the tooling that supports developers and agents in respecting the rules and following the standards.
Enforce everything through CI/CD.
People should talk with people.
AI-assisted coding should be done with proper ownership.
Agent-generated code should be human reviewed - mandatory.
Responding to a human-written issue with raw AI output is prohibited.
I’m deliberately being generic since every company and every developer has their own set of rules and standards - but some of them are like the new netiquette .
We are literally reinventing the wheel and going back to the old practices that people defined in the 80s - the netiquette, born with early computer networks and Usenet is more relevant today than ever. I’m actually totally pissed off when someone answers something I wrote with some AI slop. I took the time to write it, you must take your time to read it and write a proper answer. Use your brain.
For the tooling part, my suggestion is the obvious: strong CI/CD. Use linters, formatters, pre-commit, ensure the testing suite is human reviewed and the tests are correct. Moreover, take advantage of AI customization: a strongly-defined AGENTS.md file with clear guardrails; followed by a small set of human-written skills - to guide the agents during their thinking and tooling usage process.
AI can be a productivity booster - it’s an amplifier. Amplifying good engineering habits creates velocity without compromising quality, guaranteeing real ownership over the codebase. Amplifying laziness just fills your (once again - yours?) repository with noise, and you’ll end up being a custodian with no ownership in a short time.
Leverage the tools, speed up your workflow, but never surrender your engineering judgement.
Everyone talks about tokens nowadays. I remember when the tokens were just the result of “tokenization”, a deterministic step used when doing stuff like TF-IDF. Every token was a word, full-stop. Now instead, a token is “a token” there’s no standard definition or fixed rule, since every model decides what a token is. If you have a locally deployed model you have the control (more or less) of what the tokens are. If the models run on the cloud, well, a token is what the model provider wants to sell you as a token. ↩
Installing software with curl | sh is a bad habit - unfortunately common nowadays. This is an attack vector exploited through malvertising. The article describes what happens when someone blindly installs a tool from a sponsored website.
What happens when your digital life depends entirely on a foreign power? By drawing a parallel between the 1656 excommunication of Spinoza and modern-day US sanctions on European citizens, this article discusses the concept of 'digital excommunication' and why European technological sovereignty is no longer optional.
How I built an automated stock analysis system that leverages Gemini to parse Italian financial news feeds, providing real-time trading recommendations. This article explores the architecture, challenges, and implementation details of integrating AI-powered news analysis into a Go-based trading system.
Complete migration guide from Vertex AI SDK to Google Gen AI SDK for Python and Go developers. Covers service account authentication, OAuth2 scope limitations, and the critical implementation details missing from Google's official documentation.
A detailed walkthrough of migrating a web service from Google Cloud to OVH, covering PostgreSQL database migration, CI/CD pipeline setup on Github Actions, and significant cost savings by migrating from Cloud to a self hosted solution. This migration represents a first step toward reducing dependency on US cloud providers while maintaining service quality.
How I leveraged AI tools like Cline to enhance the UI/UX of a website and streamline backend tasks. From redesigning pages and translating content to navigating the benefits and challenges of AI-assisted development, this blog post highlights the potential of using large language models to boost productivity while sharing key lessons learned.
Starting from Unreal Engine 5.3, Epic Games added support for the so-called modern Xcode workflow. This workflow allows the Unreal Build Tool (UBT) to be more consistent with the standard Xcode app projects, and to be compliant with the Apple requirements for distributing applications... In theory! 😅 In practice this workflow is flawed: both the code signing and the framework supports are not correctly implemented, making the creation of working apps and their distribution impossible. In this article, we'll go through the problems faced during the packaging, code signing, and notarization of an Unreal Engine application on macOS and end up with the step-by-step process to solve them all.
In the article "Custom model training & deployment on Google Cloud using Vertex AI in Go" we explored how to leverage Go to create a resource pool and train a machine learning model using Vertex AI's allocated resources. While this approach offers flexibility, there's a crucial aspect to consider: the cost implications of resource pools. This article details my experience with a sudden price increase in Vertex AI and the hidden culprit – a seemingly innocuous resource pool.
In this article we explore how to combine a large language model (LLM) with a relational database to allow users to ask questions about their data in a natural way. It demonstrates a Retrieval-Augmented Generation (RAG) system built with Go that utilizes PostgreSQL and pgvector for data storage and retrieval. The provided code showcases the core functionalities. This is an overview of how the "chat with your data" feature of fitsleepinsights.app is being dev

[truncated]

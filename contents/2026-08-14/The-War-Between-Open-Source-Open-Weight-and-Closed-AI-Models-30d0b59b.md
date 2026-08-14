---
source: "https://www.nextplatform.com/ai/2026/08/13/the-war-between-open-source-open-weight-and-closed-ai-models/5287504"
hn_url: "https://news.ycombinator.com/item?id=49298438"
title: "The War Between Open Source, Open Weight, and Closed AI Models"
article_title: "The War Between Open Source, Open Weight, And Closed AI Models"
author: "jonbaer"
captured_at: "2026-08-14T14:09:27Z"
capture_tool: "hn-digest"
hn_id: 49298438
score: 2
comments: 0
posted_at: "2026-08-14T13:30:37Z"
tags:
  - hacker-news
  - translated
---

# The War Between Open Source, Open Weight, and Closed AI Models

- HN: [49298438](https://news.ycombinator.com/item?id=49298438)
- Source: [www.nextplatform.com](https://www.nextplatform.com/ai/2026/08/13/the-war-between-open-source-open-weight-and-closed-ai-models/5287504)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T13:30:37Z

## Translation

タイトル: オープンソース、オープンウェイト、クローズド AI モデル間の戦争
記事のタイトル: オープンソース、オープンウェイト、クローズド AI モデル間の戦争
説明: あなたが深刻な病状を患っており、医師の指示に従っていると仮定して、次の質問を自問してください。

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
すべてのセクションの最新ニュース
オープンソース、オープンウェイト、クローズド AI モデル間の戦争
あなたが深刻な病状を抱えており、医師が GenAI を使用してあなたの画像と検査結果をスキャンして、あなたの病状の範囲と深さを発見し、治療法を推奨すると仮定して、次の質問を自問してください。
まず、クローズド ソース、OpenAI または Anthropic の独自モデル、Meta Platforms、Moonshot AI、Alibaba、DeepSeek、OpenAI、Google、Mistral のオープン ウェイト モデル、または Hugging Face で提供されている無数のモデルに基づく AI が必要ですか?それとも、Nvidia の Nemotron 3 モデルや、Meta Platforms の今は昔の Llama 3 および 4 モデルのような、完全にオープンソースのモデルが必要ですか?米国企業のみが製造したモデルが必要ですか?
次に、そのモデルにできる限り多くの入力パラメーターと出力重みを持たせるか、それともデータと重みが少ない合理化されたモデルが必要ですか?
これらは良い質問ですが、あなたはこれらの選択をすることはありません。医療ソフトウェア サプライヤーと医療会社がそれを行ってくれます。しかし、正しい答えとは何でしょうか？
1980 年代のプロプライエタリ ソフトウェア時代と、1990 年代の Unix、Linux、および Windows Server 革命の中で育った私は、プロプライエタリ システムとアプリケーション ソフトウェアの両方の価値と、オープン ソースの台頭を理解しています。また、私はカスタムおよび社内アプリケーション ソフトウェアを作成することの価値も知っています。これは企業コンピューティングの最初の数十年間を支配し、今でも世界中の何十万もの大企業によって行われています。コードを制御し、企業の運命を制御します。
しかし、私の偏見は商用サポート、またはハイパースケーラーやクラウド ビルドの場合にあります。

つまり、オープンソース ソフトウェアがクローズド ソースのプロプライエタリな代替品と同じレベルの機能に達すると、セルフサポートのオープン ソース ソフトウェアになります。
これは、オペレーティング システムとそれに関連するミドルウェア スタックで数十年にわたって起こってきたことです。私がカブ記者だった頃は、おそらく 25 種類の異なるデータセンター ハードウェア プラットフォームとおそらくその 2 倍の数のオペレーティング システムが存在していました。そして数年が経ち、これは X86 にまで崩壊し、現在では IBM z メインフレームと Power システム マシンがわずかに含まれるようになりました。IBM のプラットフォームは、ミッドレンジおよびエンタープライズ システムでは最後にサポートされた Unix (AIX) と最後にサポートされた独自のプラットフォーム (IBM i) を実行し、そのメインフレームは独自の z/OS を実行し、場合によっては独自の z/OS を実行することもあります。 older stuff in partitions as well as Linux. (Power Systems も Red Hat Enterprise Linux だけでなく、明らかに Linux も実行します。)
そして、Windows Server のフットプリントは縮小しつつありますが、25 年前に Windows Server のプラットフォームがデスクトップからデータセンターに移行したコアとなるバック オフィス アプリケーションやフロント オフィス アプリケーションが使用されているとは思えません。これは、最近ではサーバー シャーシにどれだけのコアを詰め込めるか、そしてデータ分析の台頭、そして現在では Linux 上でのみ実行される AI についての記述です。しかし、そうは言っても、Microsoft でさえも Linux を受け入れており、過去 10 年間、SONiC Linux ディストリビューションで Linux を使用しており、今年 5 月にはクラウド用に Azure Linux ディストリビューションを作成しました。
Application software is a more varied story. 1960 年代から 1980 年代のコンピューティングの初期の頃、ほとんどの大企業は、基本的な会計から始まり、MRP システム、そして本格的な ERP システムへと進化する独自のバック オフィス アプリケーション コードを管理していました。
しかし、そろそろ商用データのデータ分析が必要になるでしょう。

SaaS の第一波が定着し始め、同時に商用 ERP スイートが、アプリケーションに膨大な技術的負債を抱えた企業にとって十分な性能を備えたようになりました。そして、Y2K バグに対処する簡単な方法も最新のインターネット テクノロジーに移行することもできませんでした。彼らは、もうダメだと言って、自社のコードを捨て、SAP、Oracle、Microsoft などのサードパーティ アプリケーションを変更し、(可能な場合は) ビジネスに合わせてコードを変更するのと同じくらい、そのコードに合わせてビジネスを変更しました。
彼らは速度と制御を引き換えにし、アプリケーション ソフトウェア プロバイダーが社内では困難な機能を追加してセキュリティを提供できることに賭けています。
まったく新しい AI アプリケーションを導入する企業や、既存のアプリケーションを AI で拡張する企業は同じ選択に直面しており、サードパーティ アプリケーションを実行している企業は、ソフトウェア パートナーが AI 機能を追加するのを待つ傾向があるようです。しかし、ERP、SCM、CRM スイートに Web テクノロジを追加するときと同じように、AI スタックを所有して既存のアプリケーションを解体するのではなくラップしたいと考えている人もいます。また、より多くの機能にお金をかけたいと考えている人もいます。いずれにせよ、あなたは支払います。ソフトウェア開発者に投資して時間を費やすか、ソフトウェアハウスが作業を行って支払いを行うのを待つ必要があります。選択と計算がはっきりと明確になることはめったにありません。
記録のために言っておきますが、我が家では OpenAI GPT (Nicole) と Anthropic Claude (Tim) の両方に料金を払っています。これは、これらがどのように機能するかを確認するためです。 API は限られたユースケースでは問題なく機能しますが、ほとんどの企業は既にクラウド上に存在している場合を除き、データを外部モデルにパイプ出力したいとは思わないと思います。彼らは、AI モデルを自分で作成および保守するソフトウェア スイートに近づけたいと考えるでしょう。

またはサードパーティから購入したものです。
メタ プラットフォームは自社モデルをオープンソースにするかどうかで二転三転しており、最新のオープンソース Nemotron モデルとエージェント AI アプリケーションをつなぎ合わせるためのワークフロー マネージャーに対する Nvidia の積極的に寛大な条件が、今週このジレンマを最大の関心事にしています。
以前に指摘したように、Nemotron 3 モデルのトレーニングに使用されるソース コード、重み、およびデータセットを寛大に提供できる唯一の企業はおそらく Nvidia だけです。私たちの見方では、Nvidia ハードウェアに支払っている価格を考えると、ハードウェアは無料になるはずです。60 年前の IBM System/360 メインフレームでアプリケーション ソフトウェア プログラミングが行われていたのと同じように、それは偶然ではありません。 IBMは企業にアプリケーションの使い方を教えるためにプログラマーのチームを派遣し、コードはユーザーにオープンソースとして提供された。これは戦略的な利点だったため、当時は企業間でアプリケーション コードを共有するなど誰も考えていませんでした。
Meta Platforms は、10 年半前に Open Compute Project を通じて公開したハードウェアと、React JavaScript ジェネレーターや PyTorch AI フレームワークなどのオープン ソース システム ソフトウェアの両方で賞賛されるべきです。これらは、プログラマーにメタ プラットフォームの方法でインターフェイスと AI を実行させるためという、非常に利己的な理由で開かれたものであり、啓発された利己主義には何の問題もありません。
同社は初期の AI モデルをクローズドかつ独自仕様に保ちましたが、2023 年 2 月に発売され、2024 年 4 月に拡張された Llama 3 モデルではオープン ウェイト (ただしオープンソースではありません) でした。 Llama 4 では、モデルは大きくなりましたが、オープンソースではなく、オープンウェイト化されました。今年 4 月に発売された最新かつまったく異なる GenAI モデルである Musk Spark 1.1 では、

このソフトウェアはクローズドだったが、今後の 1.2 リリースではオープンウェイトになると Meta Platforms は述べている。今週発表された小型の Muse Glimmer モデルは、300 億のパラメータを持ち、単一の GPU を搭載した PC 上で動作するように設計されており、最初からオープンウェイトです。
趣味で、オープンウェイトとオープンモデルの巨大な表を作成してみました。 Anthropic、Google、OpenAI、Amazon Web Services の独自モデルのフィードと速度は一般には提供されていません。この表は、300 億以上のパラメーターを持つモデルを示しています。私の知る限り、それより小さいものは個人用デバイスや限定された用途にのみ適しています。私は 300 億のパラメーターが思考のためのシナプスの最低限の数であると考えており、モデルの実行に制限されているハードウェアで取得できる限り多くのパラメーターをモデルに含めることを好みます。
最後にもう 1 つ: 小規模なモデルの多くは、高密度のモデルであっても、専門家による設計が混合されたものであっても、より大きなモデルを蒸留して作成されます。メタプラットフォームズの共同創設者兼最高経営責任者であるマーク・ザッカーバーグ氏は、今週の彼の超仲良しAIマニフェストの中で、同社が「パーソナルスーパーインテリジェンス」という理想を推進する中でいかに利他的になるか、そしてオープンウェイトモデルがその取り組みの一部であることは間違いないことについて多くのことを語ったが、モデル間の蒸留はここ米国で許可する必要があると考えているとも述べた。 （中国はすでに米国のモデルを抽出して独自のモデルを訓練していると思います。）
ザッカーバーグ氏は「米国がオープンソースでリードするには、蒸留やトレーニングでのデータ利用など、いくつかの分野で政策を再考する必要がある」と述べた。 「モデルが他のモデルから学習できることは、オープンソース エコシステムがどのように機能するかについての重要な原則です。すべての AI モデルは、

人間の知識に由来するもの。蒸留を有害なものとして組み立てようとする人もいますが、観察できるものから学べるという原則を守ることが重要だと思います。これが世界の仕組みであり、この面で自国を制限すれば米国は主導できなくなるだろう。」
すべてのモデルを相互に平均すると、モデルはどうなるのでしょうか。そして、もしそれが実現できるとしたら、GenAI モデルの推論をサポートする基盤となるハードウェア以外のものに対して、どうやって課金できるようになるのでしょうか?
繰り返しになりますが、クローズド モデルの品質がオープン モデルよりも優れたままであれば、彼らはそこに固執して誰にでも API アクセスを販売し、オンプレミスのユースケース向けにモデルのライセンスを供与することができます。
オープンソース、オープンウェイト、クローズド AI モデル間の戦争
GenAI ブームは Supermicro を引き上げるが、他の企業も引き上げる
Nvidia、新しい GenAI モデルとルーターで高い利益を実現
AI ネットワークが 3 つの方向に拡張するにつれて、Arista も拡張します
Taalas を使用すると、AMD は AI 推論を自社のチップに直接組み込むことができます
D-Wave がデュアルレール量子アーキテクチャ向けの 2 量子ビット誤り訂正ゲートを紹介
AI投資の加速に伴い、Googleが新たなGeminiモデルチームを構築
あなたのプラットフォームは別の時代に構築され、AI がそれを暴露しただけです
計算する
AMD は Agentic AI の波を捉え、見事に乗り切ります
Amazon が AWS と Annapurna Labs をスピンオフすべきではないのはなぜですか?
IBM：量子の優位性が達成されたことを3つの実証が証明
AI は Microsoft の話題を支配するが、会社のビジネスを支配するわけではない
AI インフラストラクチャの支出予測はどれほどバラ色なのでしょうか?
QuiX は Carina システムでフォトニック量子コンピューティングを前進させます
コンピューティング
AMDがお金を追いかけるために使用しているラックスケールAIシステムのロードマップ
コンピューティング
AI ホストとサンドボックスがインテルのデータセンター CP を節約

U クッキー
HPC
Nvidia、AI エージェントでチップ エンジニアリングを加速
AMDがラックスケールAIシステムのロードマップで追い求めている資金
米国エネルギー省、AI をターゲットとしたジェネシス ミッションに向けて 50 億ドルのスターター ガンを発射
Google の Axion CPU がクラウドにもたらすメリット
Salience Labs はシリコン フォトニクス光スイッチで AI をスケールアップしたいと考えています
GenAI のハードウェア投資はモデルとプラットフォームの収益を大きく上回っています
Microsoft、大規模な AI CPU および GPU クラスターのために AMD を活用
Google、量子誤り訂正に AI 強化学習を使用
マーベル、Teralynx T100 で基数、低遅延、帯域幅を実現
AI
AI チップが TSMC の収益の約 3 分の 1 を占める
スポンサーあり
SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
コンピューティング
量子古典 HPC データセンターにおける HPE とデルの願望
AI 対応データが真の利点である理由
ハイエンド コンピューティングを詳しくカバー
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
レジスター
開発クラス
ブロックとファイル
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を販売しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Ask yourself these questions, assuming you have a serious medical condition and your doctors are goi ...

Jump to main content
Search
More topics
All the latest news, from all sections
The War Between Open Source, Open Weight, And Closed AI Models
Ask yourself these questions, assuming you have a serious medical condition and your doctors are going to be using GenAI to scan your images and test results to discover the breadth and depth of your condition and to recommend treatment.
First, do you want that AI based on a closed source, proprietary model from OpenAI or Anthropic, an open weight model from Meta Platforms, Moonshot AI, Alibaba, DeepSeek, OpenAI, Google, Mistral, or the countless models being offered up on Hugging Face? Or do you want a fully open source model like the Nemotron 3 models from Nvidia or the now ancient Llama 3 and 4 models from Meta Platforms? Do you want a model made by a US company only?
Second, do you want that model to have the largest number of input parameters and output weights as possible or do you want a streamlined model with less data and fewer weights?
These are good questions, but you will not be making these choices. Medical software suppliers and healthcare companies will be doing that for you. But what are the right answers?
Having grown up in the proprietary software era of the 1980s and the Unix, Linux, and Windows Server revolutions of the 1990s, I see the value of both proprietary system and application software and the rise of open source. I also know the value of creating custom and inhouse application software, which dominated the first several decades of corporate computing and which is still done by hundreds of thousands of large corporations around the globe. Control your code, control your corporate fate.
But the bias I have is towards commercially supported or, in the case of the hyperscalers and cloud builders, self-supported open source software once that open source software rises to the same level of functionality as the closed source, proprietary alternatives.
This is what has happened in operating systems and their related middleware stacks over the decades. Back when I was a cub reporter there were probably on the order of 25 different datacenter hardware platforms and maybe twice that number of operating systems, and over the years this has collapsed down to X86 and now Arm iron with a smattering of IBM z mainframe and Power systems machinery, and IBM’s platforms run the last supported Unix (AIX) and last supported proprietary platform (IBM i) in midrange and enterprise systems, and its mainframes run proprietary z/OS and sometimes older stuff in partitions as well as Linux. (Power Systems run Linux too, obviously, and not just Red Hat Enterprise Linux.)
And while Windows Server footprints are shrinking out there, I do not think usage is for the core back office and front office applications upon which that platform jumped from our desktops to our datacenters two and a half decades ago. This is more a statement about how many cores you can cram in a server chassis these days and the rise of data analytics and now AI, which runs exclusively on Linux. But, that said, even Microsoft is embracing Linux, having used Linux underneath its SONiC Linux distro for the past decade and having created the Azure Linux distro for its cloud this May.
Application software is a more varied story. Back in the early days of computing in the 1960s through the 1980s, most large enterprises controlled their own back office application code, starting with basic accounting and evolving into MRP systems then full-blown ERP systems.
But about the time data analytics on that commercial data started taking hold, the first wave of SaaS hot at the same time commercial ERP suites got good enough for companies with a huge amount of technical debt in their applications – and no easy way to deal with the Y2K bug as well as move to modern Internet technologies – said to hell with it and tossed out their code and modified third party applications from the likes of SAP, Oracle, and Microsoft and changed their businesses to meet that code as much as changed that code (where possible) to meet their business.
They traded control for speed, and are betting the application software providers can add features and provide security that they would have trouble doing in house.
Those deploying net-new AI applications or augmenting their existing applications with AI are facing the same choices, and those running third party applications seem inclined to wait for their software partners to add the AI features. There are those, however, who want to own their AI stack and wrap around these existing applications rather than gut them – and spend more money for more functionality – much as they did with adding web technologies to their ERP, SCM, and CRM suites. Either way, you pay. You either have to invest in software developers and spend time or wait for the software houses to do the work and pay them. The choice and the math is rarely crisply clear.
And for the record, we pay for both OpenAI GPT (Nicole) and Anthropic Claude (Tim) in our household, just so we can see how these things work. An API works just fine for our limited use cases, but I do not believe most enterprises want to pipe their data out to an external model unless they live in the cloud already. They will want to have their AI models close to the software suites they either created and maintain themselves or they bought from third parties.
Meta Platforms flip-flopping on whether to be open source with its models and Nvidia’s aggressively generous terms for its most recent and open sourced Nemotron models as well as the workflow manager for stitching together agentic AI applications are bringing this dilemma top of mind this week.
As we have pointed out before, Nvidia is probably the only company that can afford to be generous and give away the source code, the weights, and the datasets used to train its Nemotron 3 models. The way we see it, at the price you are paying for Nvidia hardware, the hardware damned well should be free – and not coincidentally, just as application software programming was with the IBM System/360 mainframes six decades ago. IBM sent in teams of programmers to teach companies how to do applications, and the code was open source to the users. No one thought to share application code across companies back then because this was a strategic advantage.
Meta Platforms has to be commended for both the hardware it opened up through the Open Compute Project a decade and a half ago as well as for open source systems software such as the React JavaScript generator and the PyTorch AI framework. These were opened for very self-serving reasons – to get programmers to do interfaces and AI the Meta Platforms way – and there is absolutely nothing wrong with enlightened self-interest.
The company kept its initial AI models closed and proprietary, but was open weight (but not open source) with its Llama 3 models launched in February 2023 and expanded in April 2024 . With Llama 4, the models got bigger , but they were also not open sourced but rather open weighted. With Musk Spark 1.1, its latest and very different GenAI model launched in April this year , the software was closed, but with the upcoming 1.2 release, Meta Platforms says it will go open weight. The smaller Muse Glimmer model announced this week, with 30 billion parameters and designed to run on a PC with a single GPU, is open weight from the get-go.
Just for fun, I put together a giant table of the open weight and open models. The feeds and speeds of the proprietary models from Anthropic, Google, OpenAI, and Amazon Web Services are not provided to the public. This table shows any model with 30 billion parameters or more. Anything smaller than that is really only suitable for personal devices and limited use as far as I am concerned. I think of 30 billion parameters are the bare number of synapses for thinking, and I like as many parameters in a model as I can get for the hardware that I am restricted to running it on.
One last thing: Many of the smaller models are created by distilling down a larger model, whether it is a dense one or a mixture of experts design. In his super-chummy AI manifesto this week , Mark Zuckerberg, co-founder and chief executive officer of Meta Platforms, said a lot of things about how altruistic the company was going to be as it advanced its ideals of “personal superintelligence,” and how open weight models would be part of that effort for sure, but also that he thought that distillation across models was something we needed to permit here in the US. (I would say that China has already been distilling from US models to train its own.)
“For the US to lead in open source, we will need to rethink our policies in several areas, including distillation and data use in training,” Zuckerberg said. “The ability for models to learn from other models is an important principle of how the open source ecosystem works. All AI models are derived from human knowledge. Some have tried to frame distillation as harmful, but I think it is important to protect the principle that you can learn from anything you can observe. This is how the world works, and the US will not be able to lead if we restrict ourselves on this front.”
I wonder what happens to models when they are all averaging against each other. And if that can happen, how will anyone be able to charge for anything other than the underlying hardware supporting inference for a GenAI model?
Then again, if the quality of the closed models can stay better than the open models, they can hang in there and sell API access to anyone and everyone and still license their models for on-premises use cases.
The War Between Open Source, Open Weight, And Closed AI Models
The GenAI Boom Will Lift Supermicro, But It Will Lift Others, Too
Nvidia Drives Bang For The Buck With New GenAI Model And Router
As AI Networks Scale In Three Directions, So Does Arista
With Taalas, AMD Can Bake AI Inference Directly Into Its Chippery
D-Wave Intros Two-Qubit, Error Correcting Gate For Its Dual-Rail Quantum Architecture
Google Builds A New Gemini Model Team As AI Investments Accelerate
Your Platform Was Built For A Different Era, And AI Just Exposed It
compute
AMD Catches The Agentic AI Wave And Will Ride It Up Masterfully
Why Shouldn’t Amazon Spinoff AWS And Annapurna Labs?
IBM: Three Demonstrations Prove Quantum Advantage Has Been Reached
AI Dominates The Microsoft Conversation, But Not The Company’s Business
Just How Rosy Are Those AI Infrastructure Spending Forecasts?
With The Carina System, QuiX Pushes Photonic Quantum Computing Forward
COMPUTE
The Rackscale AI System Roadmaps That AMD Is Using To Chase Money
COMPUTE
AI Hosts And Sandboxes Save Intel’s Datacenter CPU Cookies
HPC
Nvidia Accelerates Chip Engineering With AI Agents
The Money AMD Is Chasing With Its Rackscale AI System Roadmaps
DoE Fires The $5 Billion Starter Gun For Its AI-Targeted Genesis Mission
How Google's Axion CPUs Benefit The Cloud
Salience Labs Wants To Scale Up AI With Silicon Photonics Optical Switch
GenAI Hardware Investments Are Way Ahead Of Model And Platform Revenues
Microsoft Taps AMD For At Scale AI CPU And GPU Clusters
Google Uses AI Reinforcement Learning For Quantum Error Correction
Marvell Brings Radix, Low Latency, And Bandwidth To Bear With Teralynx T100
AI
AI Chips Drive Around A Third Of TSMC Revenues
SPONSORED
SREs To AI Agents: Prove Yourself Before You Touch Production
COMPUTE
The Aspirations Of HPE And Dell In The Quantum-Classical HPC Datacenter
Why AI-Ready Data Is The Real Advantage
In-depth coverage of high end computing
Contact us
Advertise with us
Who we are
Newsletter
The Register
DevClass
Blocks and Files
Cookies Policy
Privacy Policy
Ts & Cs
Do not sell my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

---
source: "https://archcloudlabs.com/projects/selfstudy/"
hn_url: "https://news.ycombinator.com/item?id=49297252"
title: "Self Study in the Age of LLMs"
article_title: "Self-Study in The Age of LLMs · Arch Cloud Labs"
author: "DLLCoolJ"
captured_at: "2026-08-14T11:37:34Z"
capture_tool: "hn-digest"
hn_id: 49297252
score: 1
comments: 0
posted_at: "2026-08-14T11:18:57Z"
tags:
  - hacker-news
  - translated
---

# Self Study in the Age of LLMs

- HN: [49297252](https://news.ycombinator.com/item?id=49297252)
- Source: [archcloudlabs.com](https://archcloudlabs.com/projects/selfstudy/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T11:18:57Z

## Translation

タイトル: LLM 時代の自己学習
記事のタイトル: LLM の時代における自己学習 · Arch Cloud Labs
説明: 物を作ることと物を壊すこと。

記事本文:
博士課程の一環として、私は定期的に学術論文を読んだり、プレゼンテーションを鑑賞したり、先行研究の再現を試みたりしています。
多くの場合、これは結局のところ、新しいファジング フレームワークやプログラミング言語を試したり、これまで見たことのないコードベースを操作したりすることになります。
新しいテクノロジーの実験に伴う認知的摩擦は、真の学習が行われる場所であり、スキルを習得するプロセスに不可欠です。
ただし、高度なトピックと限られた自由時間の組み合わせにより、技術的な概念を深く理解するのが難しい環境が生まれます。
大規模言語モデル (LLM) は技術的なタスクに対するソリューションを迅速に打ち出すことができるかもしれませんが、私は構築しているものの技術的なアーキテクチャを深く理解することが重要であると強く信じています。そして、まず「昔ながらの方法」で物事を行い、情報セキュリティの技術を理解することがこれまで以上に重要であると考えています。
今年の DEF CON では、講演会でも、DEF CON 参加者同士の特別な会話でも、LLM が繰り返し話題になりました。
情報セキュリティ分野に初めて携わる参加者は、「基礎については初めてですが、LLM の出力が正しいことをどのように確認すればよいでしょうか?」と質問していました。
この質問は、LLM の出力に対する懐疑を示し、LLM が生成するものをただ盲目的に受け入れないようにするため、素晴らしい質問だと思います。
私はよく LLM を使用して、論文または一連の論文を中心にカスタムの「カリキュラム」を構築し、短い形式のチュートリアルの形式で不足点を埋めるのに役立ちます。
このアプローチにより、あらゆる学術論文、コードベース、または Capture The Flag 問題をほぼ瞬時に「卒業研究室」に変えることができ、最終的には自分の理解をさらに発展させ、検証することができます。
このブログ投稿では、LLM の活用に関する私のアプローチと考えを示します。

独学。
リーズナブルな定期サブスクリプションで選択できる LLM プロバイダーが多数あります。
しかし、さまざまなベンダーからモデルがリリースされるペースが速いため、実験のためにさまざまなモデルを試してみたいという私の好奇心がそそられます。
さらに、私が行っている作業の一部はガードレールをトリガーする可能性があります。本質的に「攻撃的」とみなされる可能性のあるものはすべてアラートをトリガーし、セッションを中断する可能性があるためです。
このため、私は研究目的で、Framework Desktop、OpenRouter、OpenAI の ChatGPT 上でローカルにホストされたオープンウェイト モデルを組み合わせて使用​​しています。
3 つすべての理由について説明しましょう。
OpenAI ChatGPT: 技術的な理解の評価
この記事の執筆時点では、ChatGPT は個人的な Web ベースのチャットで月額約 20 米ドルです。
インターフェイスが提供するプロジェクト タブを使用すると、PDF、写真などのリソースをアップロードし、このデータをコンテキストに含めることができます。
ファジングの特定の領域 (例: 有向ファジング) に時間を費やしている場合は、いくつかの論文をこれらの「プロジェクト」の 1 つにアップロードして、内容の理解をテストできます。研究の弁護の準備をしているときに、私は ChatGPT に「敵対的アドバイザー」になってもらい、内容について批判的に質問して、私の理解を評価してもらいました。
質問には複数の段落での回答が必要で、学術 PDF を含む ChatGPT プロジェクトから派生したものでした。
この方法論を成功させるために重要なのは、論文を読むという事前の作業を行う必要があるということです。
深い技術的な内容を読み、メモをとり、研究者がどのようなギャップを解決しているのか疑問を抱くことによる認知的摩擦は、絶対に不可欠です。単にコンテンツをアップロードし、それに対してクイズを出して「崖ノート」を把握するだけでは不十分です。
あなたは再ではありません

同盟国が利益を得ようとしていて、あなたは自分自身を騙しているだけになってしまいます。
LLM は単に研究論文を要約することはできますが、これは真の概念を習得するための目標ではありません。
自分をだまさないで、仕事をし、新聞を読みましょう。
このアプローチから私が個人的に得た最大の利点は、見落としていた領域、または論文からの主要な技術的貢献としてすぐには認識されなかった領域を特定できることです。
時々、本当に興味のある特定のコンポーネントに集中してしまい、そのせいで、著者が克服した大きな貢献や課題から気が散ってしまうことがあります。
ここで説明する ChatGPT を使用したアプローチにより、貢献全体を確実に把握し、視野狭窄と戦うことができます。
オープンルーター: 新しいコードベースの学習
OpenRouter は本質的に、単一のインターフェイスを介した多数の大規模言語モデル プロバイダーへのプロキシです。
選択した金額をアカウントに入金し、お気に入りのコーディング エージェント (例: OpenCode) を介して特定の LLM にリクエストを行うことができます。
Anthropic、OpenAI、Google、Moonshot などの企業による典型的なヘッドライン フロンティア モデルを含め、数百のモデルが利用可能です。
これらのプロバイダーにはそれぞれ、価格設定とガードレールが異なるさまざまなモデルがあります。
ガードレールを引き起こす可能性のあるセキュリティ研究の概念に関するチュートリアルを作成する際のガードレールについて理解する。
最近、フレームワークの使用方法を理解できるようにチュートリアルを構築するために、LibAFL コードベースで Kim K3 を使用しました。
LibAFL は、一般的な AFL/AFL++ フレームワークのコア コンポーネントを、研究者がオーダーメイドの方法で使用できる「ビルディング ブロック」に抽象化します。
これにより、研究者は一般的なファジング フレームワーク アーキテクチャの車輪の再発明を回避し、特定の難しい問題に集中することができます。
私の課題は、LibAFL という言語である Rust を初めて使用することです。

書かれており、各「構成要素」に使用される用語は少し馴染みがありません。繰り返し使用して理解を深められるハーネスをすぐに構築したいと考えています。
ここで、OpenRouter + $OPEN_SOURCE_CODE_BASE を使用したコーディング エージェントが有益な学習環境を作成していることがわかりました。
突然、大規模な GitHub リポジトリを大学院レベルのラボに抽出して、自分の時間内で完了し、フレームワークの使用方法についてより深く理解できるようになります。
行き詰まった場合は、生成されたチュートリアルと質問を LLM に提供して、問題を理解できるまで問題をさらに詳しく分析することができます。
これと同じアプローチが、Capture The Flag 問題でも実行できます。これは、OpenCode でのファジングへの私の慣れ度を測るために一連の質問をした後、Kimi K3 が準備しているチュートリアルの内訳のスクリーンショットです。
最終的な製品は、完成すべきさまざまなモジュールを含むディレクトリのルートにある README.md です。
drwxr-xr-x 1 dllcoolj dllcoolj 18 7 月 25 日 20:36 00-mental-model
drwxr-xr-x 1 dllcoolj dllcoolj 18 7 月 25 日 20:40 01-ハーネス
drwxr-xr-x 1 dllcoolj dllcoolj 80 7 月 26 日 11:45 02-inprocess-libafl
drwxr-xr-x 1 dllcoolj dllcoolj 18 7月 26 10:41 03-real-repos-build-integration
drwxr-xr-x 1 dllcoolj dllcoolj 80 7 月 26 日 11:45 04-forkserver
drwxr-xr-x 1 dllcoolj dllcoolj 104 7 月 26 日 11:45 05-スーパーチャージャー
drwxr-xr-x 1 dllcoolj dllcoolj 18 7 月 26 日 11:02 06-オペレーション
-rw-r--r-- 1 dllcoolj dllcoolj 3705 7 月 25 日 20:37 README.md
drwxr-xr-x 1 dllcoolj dllcoolj 220 7月26日 10:40 tinycfg
これを読んで、なぜクロード コードについて言及されていないのかと疑問に思っている方は、私はクロードを使用しており、気に入っています。
しかし、過去 1 年間に経験した使用制限やガードレールなどの数多くの制限により、サブスクリプションをキャンセルせざるを得なくなりました。
コーディング用

sks、LLM を使用する場合、Doom Emacs プラグインで自己ホストするオープンウェイト モデルを使用します。
これにより、開発中にすべてをコード化してもらうのではなく、開発中の最新情報を常に把握し、アシスタントとして使用できるようになります。
Framework Desktop: ソフトウェア開発用にローカルでホストされるオープンウェイト モデル
サイド プロジェクトのソフトウェア開発では、個人的にはオープンウェイト モデルの機能に満足しています。
繰り返しますが、これは私がユーザーであるソフトウェアを構築する自宅ラボのタスク用であるため、エンタープライズ規模のソフトウェアは私の使用例ではありません。
私は通常、新しいフレームワークや、ログ集約や pkgfile 解析に関連するものを実験するための小さなプロジェクトを構築しています。
これらはそれほど複雑なタスクではなく、今日のオープンウェイト モデルで簡単に実行できます。
小規模なモデルで問題が発生した場合は、Doom Emac の gptel と、リクエストで emacs バッファ全体を送信するためのサポートを介して、マニュアル ページの形式でドキュメントをコンテキスト ウィンドウに簡単に提供できます。
さらに、ローカルでホストされるモデルには 3 つの重要な利点があると考えています。
データのプライバシー。すべての質問を最終的にトレーニングして次世代モデルを改善するために送信するわけではありません。
私は、まだ一般公開の準備が整っていないと思われる、進行中の研究プロジェクトであるプロジェクトの場合に特にこれを楽しんでいます。
定期的な請求はありません。ただし、これには多額のコンピューティング初期コストがかかります。
LLM ホスティングのスキル開発、予期せぬボーナス!
私は今年の初めに、さまざまな LLM ランタイムの学習を開始するために Framework Desktop を購入し、友人のためにモデルをホスティングする短い期間を開始しました。
Tailscale と適切な家庭用インターネット接続を使用して、友人が LLM を利用できるようにし、友人からのフィードバックに基づいて導入パラメータを調整して、継続的な速度を最適化しました。

完了しようとしていたタスクに応じて、ext ウィンドウを開きます。この経験により、DEF CON: Packet Hacking Village Talk「There's A Bug In My Boot: Finding Vulnerabilities in My Boot」の一環として、オフライン モデルを使用して不正な形式の SquashFS ファイル システムを構築する際に効果的に反復することができました。
過度の依存を避けるための結論と考え方
この投稿に共通する傾向は、依然として「アクティブ ラーニング」の作業を行う必要があるということです。
独学で LLM を活用する際に、確実に積極的に学習できるかどうかは、人間であるあなた次第です。
LLM に複雑なプロジェクトを数分で生成させるのは簡単ですが、これらのプロジェクトを実際に理解しているかどうかは別の話です。
おそらくそれがそれほど重要ではない状況もあるでしょうが、それがどのように機能するかを真に理解するために何かを学んでいるのであれば、自分自身をだまさずに実際の作業を行うことが重要です。
LLM は役に立たないため、LLM をさまざまなユースケースに適用する方法を理解することがますます重要になります。
すべてに使用しなくても構いませんが、その機能を理解するためにそれらを試してみることを避けるのは間違いだと思います。

## Original Extract

Making things & breaking things.

As part of my doctorate program, I am regularly reading academic papers, watching presentations, and attempting to recreate prior research.
Often, this boils down to experimenting with a new fuzzing framework, programming language, or navigating a codebase that I’ve never seen before.
The cognitive friction involved in experimentation with new technology is where the real learning happens, and is essential to the process of skill mastery.
However, advanced topics in combination with limited free time create an environment that can be challenging to deeply understand a technical concept.
While Large Language Models (LLMs) might be able to whip out solutions for a technical task quickly, I firmly believe it’s important to deeply understand the technical architecture of what you’re building, and I think it’s more important than ever to do things “the old fashion way” first, and understand the craft of Information Security.
At DEF CON this year, LLMs were a recurring topic at both the talks and ad-hoc conversations with fellow DEF CON goers.
Attendees new to the Information Security field were asking “how can I know the LLM output is correct, if I’m new to the fundamentals?”
I think this question is great as it signals skepticism about the output of an LLM, and to not just blindly accept what an LLM produces.
I’ll often use LLMs to build a custom “curriculum” around a paper or set of papers to help me fill in my gaps in the form of short form tutorials.
This approach allows any academic paper, codebase or Capture The Flag problem to be turned into a “graduate lab” nearly instantaneously, which ultimately allows me to further develop and verify my understanding.
This blog post will show my approach and thoughts on leveraging LLMs for self-study.
There are many LLM providers you can choose from with a reasonable recurring subscription.
However, the rate at which models are being released from a variety of vendors piques my curiosity to try out different models for experimentation.
Additionally, some of the work I’m doing may trigger guardrails, as anything that may be deemed “offensive” in nature could trigger an alert, and disrupt a session.
Because of this, I use a combination of locally hosted open-weight models on a Framework Desktop, OpenRouter, and OpenAI’s ChatGPT for research purposes.
Let’s discuss the reason for all three:
OpenAI ChatGPT: Assessing Technical Understanding
At the time of this writing ChatGPT, is ~$20 USD a month for personal web-based-chat.
The projects tab that the interface offers allows you to upload resources such as PDFs, pictures, etc… and include this data in your context.
If I’m spending time on a particular area of fuzzing (Ex: directed fuzzing), I can upload a handful of papers into one of these “projects” and test my understanding of the content. While preparing for my research defense, I prompted ChatGPT to be an “adversarial advisor” and critically ask questions about the content to assess my understanding.
The questions asked required multiple paragraph responses, and were derived from the ChatGPT projects with academic PDFs.
The important part for this methodology to be successful is that you still have to do the upfront work of reading the papers.
The cognitive friction of reading deep technical content, taking notes, and questioning what gaps the researchers are solving, and their technical approach is absolutely essential. Simply uploading content, and being quizzed on it to grasp the “cliff notes” is insufficient.
You’re not really going to get the benefits, and you’d just be cheating yourself.
While LLMs can simply summarize a research paper for you, this is not the goal for true concept mastery.
Don’t cheat yourself, do the work, read the paper.
The largest benefit I’ve personally had from this approach is identifying areas that I overlooked or didn’t immediately see as a major technical contribution from a paper.
Sometimes I get focused on a particular component that I’m really interested in, and that may distract me from a large contribution or challenge the author(s) overcame.
The approach I describe here with ChatGPT ensures I see the whole contribution and fight my tunnel vision.
Open Router: Learning New Codebases
OpenRouter is essentially a proxy to numerous Large Language Model providers via a single interface.
You can deposit however much money you choose into your account, and make requests to a given LLM via your favorite coding agent (Ex: OpenCode) .
There are hundreds of models available, to include your typical headline frontier models from companies like Anthropic, OpenAI, Google, Moonshot, etc…
Each of these providers have different models with different pricing, and guardrails.
Understanding the guardrails when building tutorials around security research concepts that may trigger guardrails.
Recently, I used Kimi K3 with the LibAFL codebase to help build tutorials to enable my understanding of how to use the framework.
LibAFL abstracts core components of the popular AFL/AFL++ frameworks into “building blocks” that can be used by researchers in bespoke ways.
This allows researchers to avoid reinventing the wheel of common fuzzing framework architectures and focus on their particular hard problem.
My challenge is that I’m new to Rust, the language LibAFL is written in, and the terminology used for each “building block” is slightly unfamiliar, and I’d like to quickly build a harness that I can iterate on to develop a deeper understanding.
This is where I find Coding Agents with OpenRouter + $OPEN_SOURCE_CODE_BASE creates a beneficial learning environment.
Suddenly any large GitHub repo can be distilled into graduate-level labs to complete on your own time and develop a deeper understanding of how to use the framework.
If you get stuck, you can provide the generated tutorial + your question to the LLM to further break down the problem until you’ve successfully understood the issue.
This same approach can be done with Capture The Flag problems. Here’s a screenshot of the tutorial breakdown that Kimi K3 is preparing after asking a series of questions to gauge my familiarity with fuzzing in OpenCode.
The final product being a README.md in the root of the directory with various modules to complete.
drwxr-xr-x 1 dllcoolj dllcoolj 18 Jul 25 20:36 00-mental-model
drwxr-xr-x 1 dllcoolj dllcoolj 18 Jul 25 20:40 01-harnessing
drwxr-xr-x 1 dllcoolj dllcoolj 80 Jul 26 11:45 02-inprocess-libafl
drwxr-xr-x 1 dllcoolj dllcoolj 18 Jul 26 10:41 03-real-repos-build-integration
drwxr-xr-x 1 dllcoolj dllcoolj 80 Jul 26 11:45 04-forkserver
drwxr-xr-x 1 dllcoolj dllcoolj 104 Jul 26 11:45 05-supercharging
drwxr-xr-x 1 dllcoolj dllcoolj 18 Jul 26 11:02 06-operations
-rw-r--r-- 1 dllcoolj dllcoolj 3705 Jul 25 20:37 README.md
drwxr-xr-x 1 dllcoolj dllcoolj 220 Jul 26 10:40 tinycfg
If you’re reading this and wondering why Claude Code was not mentioned, I have used Claude, and liked it.
However, the numerous restrictions I’ve experienced over the past year with usage limits, and guardrails has made me cancel my subscription.
For coding tasks, if I use an LLM, I’m using open-weight models that I self host with a Doom Emacs plugin .
This allows me to stay in the loop during development and use it as an assistant rather than just have it code everything for me.
Framework Desktop: Locally Hosted Open-Weight Models for Software Development
For side project software development I’ve been personally satisifed with the capabilities of open-weight models.
Again, this is for home lab tasks where I’m building software that I am the user for so enterprise scale software isn’t my use case.
I’m typically building small projects to experiment with new frameworks or something related to log aggregation or pkgfile parsing .
These are not the craziest of complex tasks, and can easily be performed with today’s open-weight models.
If smaller models struggled, I can easily provide docs in the form of man pages to my context window via Doom Emac’s gptel and its support for sending entire emacs buffers in a request.
Additionally, I see three key benefits to locally hosted models:
Data privacy. You’re not shipping off all of your questions to ultimately be trained on to improve the next generation of models.
I particularly enjoy this for projects that I don’t think are quite ready for public viewing yet, and are very much a research project work in progress.
No recurring bill. However, this comes at a hefty upfront cost of compute.
Developing skills of LLM hosting, an unexpected bonus!
I purchased a Framework Desktop earlier this year to start learning the various LLM runtimes, and began a brief stint of model hosting for friends.
With Tailscale, and a reasonable home internet connection, I made LLMs available to friends and tweaked deployment parameters based on their feedback to optimize for speed over context window depending on the task they were trying to complete. This experience allowed me to iterate effectively when using offline models building malformed SquashFS file systems as a apart of my DEF CON: Packet Hacking Village Talk, “There’s A Bug In My Boot: Finding Vulnerabilities in U-Boot” .
Conclusion & Thoughts On Avoiding Over Reliance
A common trend in this post is that you still have to do the work of “active learning”.
It is up to you, the human, to ensure you’re actively learning when leveraging LLMs for self-study.
It is trivial to have LLMs generate complex projects in minutes, but whether or not you actually understand these projects is a different story.
Perhaps there are situations where it matters less, but if you’re learning something to truly understand how it works, it’s critical to not cheat yourself and do the actual work.
LLMs aren’t going anywhere, so understanding how to apply them to various use cases will only become more important.
It’s okay to not use them for everything, but to avoid experimenting with them to understand their capabilities I believe would be a mistake.

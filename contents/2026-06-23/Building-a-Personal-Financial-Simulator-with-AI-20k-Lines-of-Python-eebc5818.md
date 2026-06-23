---
source: "https://warpsimlab.org/reports/building-warpsimlab-with-ai.html"
hn_url: "https://news.ycombinator.com/item?id=48645977"
title: "Building a Personal Financial Simulator with AI: 20k Lines of Python"
article_title: "Seven Months and 20,000 Lines of Python Later: Lessons from Building a Personal Financial Simulator with AI | WARPSimLab"
author: "alanne"
captured_at: "2026-06-23T15:00:48Z"
capture_tool: "hn-digest"
hn_id: 48645977
score: 1
comments: 0
posted_at: "2026-06-23T14:54:01Z"
tags:
  - hacker-news
  - translated
---

# Building a Personal Financial Simulator with AI: 20k Lines of Python

- HN: [48645977](https://news.ycombinator.com/item?id=48645977)
- Source: [warpsimlab.org](https://warpsimlab.org/reports/building-warpsimlab-with-ai.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T14:54:01Z

## Translation

タイトル: AI を使用したパーソナル金融シミュレーターの構築: 20,000 行の Python
記事のタイトル: Python の 7 か月と 20,000 行のその後: AI を使用した個人金融シミュレーターの構築からの教訓 |ワープシムラボ
説明: AI 支援を利用した個人財務および退職シミュレーション プラットフォームである WARPSimLab の構築に関するソフトウェア エンジニアリングの回顧。

記事本文:
メインコンテンツにスキップ
ワープシムラボ
教育用パーソナルファイナンスシミュレーションソフトウェア
ホーム >
レポート >
AI を使用した WARPSimLab の構築
7 か月と 20,000 行の Python のその後: AI を使用した個人金融シミュレーターの構築から得た教訓
過去 7 か月にわたって、私は個人の財務と退職後のシミュレーション プラットフォームである WARPSimLab を構築しました。このプロジェクトは、65 のソース ファイルにまたがる約 20,000 行の Python に成長しました。 WARPSimLab には、退職モデリング、ポートフォリオ シミュレーション、税金、モンテカルロ分析、過去の市場ウィンドウ、レポート、デスクトップ グラフィカル ユーザー インターフェイスが含まれています。
AI (ChatGPT) がプロジェクト全体に関与しました。この記事では、開発パートナーとして AI を使用して重要な中規模ソフトウェア プロジェクトに取り組んだ私の経験について説明します。
このコラボレーションの一部の部分は非常にうまくいきましたが、他の部分は何度も失敗しました。最も有益な教訓は、AI の長所と短所がどこにあるのか、問題が発生したときにそれを軽減する方法を理解することから得られました。
私は 1990 年代に DOS ベースの個人金融シミュレーター プログラムを購入しました。このシミュレーターにより、複利の力、株式の利点とリスク、時間の力など、数多くの金融概念に目が開かれました。このソフトは非常に教育的でしたが、夫婦で資産を分けられないなどの弱点がありました。結局、この DOS ベースのソフトウェアを紛失してしまい、代替品が見つかりませんでした。退職したらこのソフトを作り直すことにしました。
昨年（2025 年）、私はソフトウェア開発における AI に非常に興味を持ちました。私は 1980 年代に大学で授業を受けて以来、AI を研究してきました。 AI は 5 年後に世界を変えるツールであり、その後 40 年間もそのままでした。ここ数年、私はAIが根本から変化していくのを見てきました。

プロジェクトをソフトウェア エンジニアのツールボックス内の別のツールに移動します。 AI を使用してソフトウェアを設計および構築する方法を学ぶプロジェクトが必要でした。
ChatGPT を使用して個人的な金融シミュレーターを作成することにしました。私は ChatGPT の無料の匿名バージョンから始めて、その後無料のログイン バージョンに移行し、最終的に月額サブスクリプションを購入しました。今では当たり前のことですが、私はシミュレーターがどのように機能するのかよくわかりませんでした。そこで、カスタム金融シミュレーターへのコマンド ライン インターフェイスから始めました。数か月後、私のようなプロジェクトはすでに構築されているのではないかと思いました。オンラインで商用および無料の個人金融シミュレーターを多数見つけました。残念ながら、私の要件をすべて満たすものはありませんでした。私は、アドバイスではなく教育を対象とした、オープンソース (無料) の個人金融ソフトウェア パッケージを望んでいました。このシミュレーターは Web ベースではなく、ユーザーが個人情報をリモート サーバーにアップロードする必要があるものにはしたくありませんでした。私は、さまざまな税金とリソースのバケットをシミュレートし、さまざまな退職金引き出しスキームを使用し、モンテカルロ分析と過去のウィンドウリスク分析を提供し、一連のリターンリスクを含めるなど、十分な機能を備えていることを望んでいました。最後に、私はこのソフトウェアが探索およびもしものツールであり、ユーザーが財務がどのように機能するかを自分で発見できるようにしたいと考えていました。
7 か月後、65 のファイルと 20,000 行の Python を経て、WARPSimLab はついに公開できるほど成熟しました。
（監視された）分業
AIを使ったコーディングは面白いですね。大まかに言うと、次のように分業しました。
私は機能セットの責任者です。これには、レポートのレイアウト、モンテカルロおよび履歴ウィンドウ アルゴリズムを実装するかどうか、税金シミュレーションがどの程度複雑で完全であるかなどの日常的な詳細が含まれます。
arを担当させていただきました

アーキテクチャ。 WARPSimLab はレイヤーまたはコンポーネントで設計されています。現在、GUI レイヤー、シミュレーター レイヤー、プロット レイヤー、およびレポート レイヤーがあります。私は各層間の境界と各層の内部アーキテクチャを設計しました。
私は GUI のメイン メニューの外観と操作性、およびすべてのサブメニューをデザインしました (そして再設計し、再設計しました)。
私は、シミュレーター コア、プロット コード、レポート コード、シナリオ エクスプローラーなどのリファクタリング (およびリファクタリングとリファクタリング) を監督しました。
現在の WARPSimLab Web ページをデザインしました。私は、ピカソ、リアリー、レンブラント、ソフトウェア会社やハードウェア会社の Web サイト、そして最後に銀行や金融会社の Web サイトからインスピレーションを得たバージョンを調べました。現在のデザインを選択し、磨き上げました。
AI は私のコーディングアシスタントです。 AI は、Python の数百行の複雑な行からなるファイル全体を数十秒で作成しました。
AI は GUI コード (tkinter) を十分に実装しているため、インターフェイスを学ぶ必要はありませんでした。ダイアログのコードが書かれており、そのコードはそのまま動作します。
AI はプロジェクトのテスト コードを作成しました。これには、決定論的テスト、機能テスト、不変テスト、モジュール テストが含まれます。これらのテストの合計コード ベースは Python で 22,000 行を超えています。
AI はブレインストーミングの良きパートナーです。 「どうすれば…を実装できるでしょうか？」といった質問には、通常、数秒で答えられます。 「…を実装すべきか」または「次は何を実装すべきか」というタイプの質問には、常によく考えられた答えが返されます。
Webサイトのデザインと文章は私が書きましたが、HTMLはAIが書きました。 HTML の間違い、文法上の間違い、論理上の間違いなどがチェックされます。一例として、AI は、0 ドルになるシミュレーションの割合が、画像とテキストでは異なることを発見しました。これには、AI が次のテキストを理解できることに驚きました。

レポートの後半でテキストと一致させるために必要な画像。
AI を使用したコーディング – 私の経験。
AI はプロトタイピングにおいて優れています。ダイアログ ページがどのようなものになるだろうかと何度も考えましたが、記述された説明から AI が 1 分以内にスタンドアロンのサンプルを作成します。それが私が望んでいたものではない場合は、もう一度そうするでしょう。これにより、GUI のプロトタイプを何度も迅速に作成できるようになりました。人間のプログラマーであれば、少なくともこれには一桁以上の時間がかかるでしょう。
AI はほとんどの場合、初めてクリーンなコードを作成します。コードに間違いがある場合、それはほとんどの場合私のせいです。ファイルを頻繁に編集した後、プロジェクトに保存しなかったり、ファイルを間違って編集したりしました。 AI は編集内容が適用されていないことを辛抱強く私に教えてくれ、さらに明確な指示を与えて再度教えてくれました。
AI はリファクタリングが非常に得意です。たとえば、ファイルが複雑になりすぎたり長すぎたりするため、ファイルを分割するアドバイスを求めます。 AI は、機能とタスクをグループ化したり、インターフェイスやインクルードなどをチェックしたりすることに優れていました。
AI は、複雑なマルチファイル アルゴリズムの設計において優れています。シナリオ インスペクターを追加したとき、どのように進めるべきかは漠然としかわかりませんでした。問題は、シミュレーション層から GUI 層に戻るデータ フローをバイパスして逆にする必要があることでした。また、ユーザーが指定した状態を永続的に変更せずに、一時的に変更された状態をシミュレーター コアに注入する必要もありました。 AI は、数百行 (おそらく数千行) のコードからなる多数の新しいファイルと、数百行のインターフェイス コードを 6 個のファイルにまとめ、非常に体系化された手順で、比較的理解しやすい手順で提供しました。私はショックを受けました。初めて実行したとき、コードは正しく動作しました。
デバッグは楽しかったです。実行からのエラー出力とソースファイルを切り取ってチャットに貼り付けます。

indow を使用すると、AI が何が間違っているのかをすぐに教えてくれて、正しい編集を提供してくれるでしょう。より難しい問題の場合、AI は推奨される print ステートメントと挿入位置を提供します。場合によっては、これらのデバッグ印刷は長さ数十ページになることがあります。詳細な出力全体をコピーしてチャット ウィンドウに貼り付けるだけで、数秒以内に問題の説明と問題を修正するための編集が表示されます。すばらしい。
AI に関して最もイライラする部分の 1 つは、気付かないうちに会話を誘導できることです。 A をすべきか B をすべきか? B を行う場合、どのように行いますか?あるとき、AI と私は、退職に最も適した年についての研究プロジェクトに取り組むことについて何時間も話し合いました。待ってください – これは私たちが個人金融シミュレーターで取り組んでいたこととは何の関係もありません。
私がよく見つけたもう 1 つの問題は、混乱すると AI がウサギの穴に落ちてひたすら掘り続けることです。この問題は、[収入] ダイアログの 2 列のデータで発生しました。垂直方向の間隔の問題について数時間何度も繰り返した後、最終的には「やめて、2 つの平行な垂直フレームを作成しましょう。」と言わざるを得ませんでした。初めて働きました。
AI はしばしば自信過剰になります。 AI は主に確率エンジンです。それは「推測」することで機能します。初期の頃、私は何らかのタイプのダイアログを要求していましたが、生成されたものは私が要求したものとはまったく異なりました。そうなると、デザインを行ったり来たりして時間を無駄にすることになります。また、データのプライマリ ストアをダイアログ ボックスの tkinter 変数に移動し、これらの変数を渡そうとする場合もあります。 (変数に関する「真実」の情報源が 1 つあり、初期化関数にあります。) この問題は、すべての新しいチャットの開始時に次の内容を貼り付けることでほとんど解決されました。「このセッションでは、質問があれば質問してください。推測したり仮定したりしないでください。」
AIが幻覚を起こす –

しかし、かなり良くなってきています。初期の頃、ChatGPT は頻繁に幻覚を感じていました。しかし、過去 4 か月間でいつそうしたか思い出せません。これは、私がツールの制約内にうまく収まるようになったからなのか、ツールが改良されてきたからなのか、それともその両方なのかはわかりません。 ChatGPT が幻覚を起こして変数名を変更するとき、コードが間違っていると告げられ、コードを与えられたばかりであることに気づかず、私はいつも笑ってしまいました。
AI はリファクタリングが得意すぎる場合があります。タスクに取り組んでいると、まったく関係のないコードの大きな部分が完全に書き直されていることに気づくことがあります。初期の頃は、入力ボックスさえ消えていました。この現象はここ数か月で起こらなくなりました。繰り返しになりますが、限られたリソースを尊重することがうまくなったからなのか、それとも ChatGPT モデルがより賢くなったからなのかはわかりません。
新しいタブで作業することになるので、基本的には新しいアシスタントを使って作業することになり、新しいタブを最初から起動する必要があることを覚えておく必要がありました。そこで、私たちはウサギの穴にはまらず、すべてのチャットの冒頭にコピーされるスタートアップ段落を作成しました。最後の文はいつも「私が完了と言うまで何もしないでください」でした。そうすれば、必要なファイルをすべてアップロードして最終的な指示を与えるまで、AI が思考を開始することはありません。
AI をコーディング コンパニオンとして使用すると、コードが急速に混乱したため、コードにコメントするのをやめました。コードの循環が停止したので、戻ってコメントを追加する必要があります。驚くべきことに、AI は実際にコメントを書くのが非常に得意なので、最初はパスさせておきます。
AI は、集中するように指示されたものに集中します。ファイル内のデッドコードについて尋ねると、ファイル内のすべてのデッドコードが表示されます。これは、AI が問題を示唆することなく、数か月間多数のファイルを操作した後に実際に発生しました。
AIは人間を模倣するように設計されています。私たちとしては

仕事をしていると、ChatGPT がデザインの明瞭さや単純さ、インターフェイスの品質、明確なコーディング スタイルなどを補完してくれることがありました。これは実際には問題ではありませんでしたが、本来のタスクであるコードを書くことから気が散ってしまいました。 AI は、他のソーシャル プログラムと同様、人々の注意を引きつけるように設計されています。私の目標はソフトウェアを構築することです。これらのタスクが完全に一致しない場合もあります。
AI のリソースには限界があることがわかりました。通常、私は AI に必要なファイルを与えてチャットを開始し、目標に向かう小さなサブ機能の明確な指示を与えます。 1 分後、明確に編集された新しいコードを受け取ります。可能であれば、テスト後に別の小さな目標を提供し、コードが返され、それを繰り返すことになります。早い段階で、AI が一貫性を欠いてゴミコードを作成し始めることに気づきました。
AI はトークンの生成および処理エンジンです。残念ながら、トークンスタックは有限です。さらに、トークン スタックがいっぱいに近づくと、会話の初期部分が凝縮または抽象化され、詳細が失われます。 (これは、チャットの上部が削除されたばかりの 7 か月前よりもはるかに改善されています。) 問題は、AI (この場合は ChatGPT) がそれが起こっていることを知らせてくれないことです。人間は知的な方法で会話を抽象化します。 AIはそれを総当たりで行います。
ソフトウェア開発は非常にトークンスタックです

[切り捨てられた]

## Original Extract

Software engineering retrospective on building WARPSimLab, a personal finance and retirement simulation platform, with AI assistance.

Skip to main content
WARPSimLab
Educational Personal Finance Simulation Software
Home >
Reports >
Building WARPSimLab with AI
Seven Months and 20,000 Lines of Python Later: Lessons from Building a Personal Financial Simulator with AI
Over the last seven months I built WARPSimLab, a personal finance and retirement simulation platform. The project grew to roughly 20,000 lines of Python spread across 65 source files. WARPSimLab includes retirement modeling, portfolio simulation, taxes, Monte Carlo analysis, historical market windows, reporting, and a desktop graphical user interface.
AI (ChatGPT) was involved throughout the project. This article describes my experience tackling a significant medium-sized software project using AI as a development partner.
Some parts of this collaboration worked remarkably well while other parts failed repeatedly. The most useful lessons came from understanding where AI’s strengths and weaknesses were, and how to mitigate issues as they arose.
I bought a DOS based personal financial simulator program in the 1990’s. This simulator opened my eyes to numerous financial concepts such as the power of compounded interest, the advantages and risks of stocks, and the power of time. Although this software was extremely educational, it had weaknesses such as not separating assets for married couples. I eventually lost this DOS based software, and couldn’t find a replacement. I decided to recreate this software when I retired.
Last year (2025) I became extremely interested in AI for software development. I have been following AI since taking a class in college in the 1980’s. AI was the tool that was going to change the world in 5 years, and remained that way for the next 40 years. The last few years I watched AI change from a research project to another tool in a software engineer’s toolbox. I needed a project to learn how to design and build software with AI.
I decided to use ChatGPT to create a personal financial simulator. I started with the free, anonymous version of ChatGPT, then moved to the free, logged in version, and finally picked up a monthly subscription. Although obvious now, I wasn’t sure how a simulator worked. So, I started with a command line interface to a custom financial simulator. After a few months I wondered if projects such as mine had already been built. I found numerous commercial and free personal financial simulators online. Unfortunately, none of them satisfied all of my requirements. I wanted an open source (free) personal financial software package that was targeted at education rather than advice. I wanted this simulator to not be web based, requiring users to upload their personal information to some remote server. I wanted it to be reasonably feature complete - simulating different tax and resource buckets, using different retirement withdrawal schemes, providing Monte Carlo and historical window risk analysis, and including sequence of return risk. Finally, I wanted the software to be an exploration and what-if tool, allowing users to discover how finances work for themselves.
7 months, 65 files and 20,000 lines of Python later, WARPSimLab was finally mature enough to go public.
A (supervised) division of labor
Coding with AI has been interesting. Generally speaking, we divided labor as follows:
I have been responsible for the feature set. This includes such mundane detail as layout of reports, whether to implement Monte Carlo and historical window algorithms, how complicated and complete the tax simulation would be, etc.
I have been responsible for the architecture. WARPSimLab is designed in layers or components. Currently there is the GUI layer, the simulator layer, the plots layer, and the reports layer. I architected the boundaries between each layer and the internal architecture of each layer.
I designed (and re-designed and re-designed) the GUI main menu look and feel, along with all of the submenus.
I oversaw the refactor (and refactor and refactor) of the simulator core, the plots code, the reports code, the Scenario Explorer, etc.
I designed the current WARPSimLab webpage. I went through versions inspired by Picasso, Leary, Rembrandt, software and hardware companies’ websites, and finally banking and financial companies websites. I selected and polished the current design.
AI has been my coding assistant. AI has created whole files consisting of hundreds of complex lines of Python in a few dozen seconds.
AI has implemented GUI code (tkinter) well enough that I really haven’t needed to learn the interface. It has written the code for the dialogs, and the code just works.
AI has created testing code for the project. This includes deterministic, feature, invariant and module tests. The total code base for these tests are now over 22,000 lines of Python.
AI has been a good brainstorming partner. Questions such as “how could I implement…” are generally answered in seconds. “Should I implement… “ or “What should I implement next” type of questions are always well thought out answers.
Although I designed and wrote the text for the website, AI wrote the html. It checked for html mistakes, grammatical mistakes, logic mistakes etc. An example is AI found that the percent of simulations that went to $0 was different on a picture than in the text. This amazed me that AI could understand that the text in an image needed to match text later in the report.
Coding with AI – My experience.
AI is amazing at prototyping. Many times, I would wonder what a dialog page would look like, and from a written description, AI would create a standalone example in under a minute. If it wasn’t what I wanted, we would do so again. This allowed us to rapidly prototype the GUI many times over. This would have taken a human coder an order of magnitude more time – at least.
AI almost always creates clean code the first time. When there are mistakes in code, it’s almost always my fault. I frequently edited a file and then didn’t save it into the project, or edited files incorrectly. AI would patiently tell me that I hadn’t applied its edits – and would give them to me again with clearer instructions.
AI is extremely good at refactoring. For example, as files would get too complex and/or too long, I would ask for advice splitting them up. AI was excellent at grouping functions and tasks together, checking interfaces, includes, etc.
AI is amazing at designing complex multi file algorithms. When I added the Scenario Inspector, I had only a vague idea how to proceed. The problem was we needed to bypass and reverse the data flow from the simulation layer back to the GUI layer. We also needed to inject temporary changed state into the simulator core without permanently changing user provided state. AI provided numerous new files of many hundreds of lines of code (probably in the thousands), and hundreds of lines of interface code in half a dozen files in very organized steps that were reasonably easy to follow. I was shocked – the code worked correctly the first time it was run.
Debugging was a joy. I would cut and paste error output from a run along with source files into the chat window, and AI would quickly tell me what was wrong, and would provide correct edits. For harder problems, AI would provide suggested print statements along with insertion locations. Sometimes these debug prints would be dozens of pages in length. I would just copy and paste the whole verbose output into the chat window, and within seconds I would receive an explanation of what was wrong plus edits to fix the issue. Amazing.
One of the most frustrating parts about AI is you can steer conversations without realizing it. Should we do A or B? If we do B, how would we do it? One time AI and I discussed working on a research project about the worst years to retire for hours. Wait – this has nothing to do with what we were working on in the personal financial simulator!
Another issue I frequently found is that when confused, AI will go down a rabbit hole and just keep digging. This happened with two columns of data in the Income dialog. After iterating over and over on vertical spacing issues for a few hours, I finally had to say “Stop! Let’s create two parallel vertical frames.”. Worked first time.
AI is frequently overconfident. AI is primarily a probability engine. It works by “guessing”. Early on I would ask for a dialog of some type, and what was generated was totally different than what I had asked for. Then, we would waste time going back and forth with the design. Other times it would move the primary store of data into dialog box tkinter variables and then try to pass these variables around. (I have one source of “truth” for variables, in an initialization function.) This was mostly cured by pasting the following at the start of all new chats: “for this session, ask any questions you have, and do not guess or make assumptions.”
AI hallucinates – but is getting much better. Early on ChatGPT would hallucinate quite often. However, I can’t remember a time it has done so over the last 4 months. I don’t know if this is because I have gotten better at staying within the constraints of the tool, the tool is getting better, or both. When ChatGPT would hallucinate and change variable names, I always laughed when it would then tell me my code was incorrect, not realizing it had just given me the code.
AI is sometimes too good at refactors. Every once in a while when working on a task I would find large chunks of unrelated code totally rewritten. Early on, I would even have input boxes disappear. This stopped happening in the last few months – again, I don’t know if it’s because I am getting better at respecting limited resources or ChatGPT models are getting smarter.
As I would work in new tabs, I had to remember I was basically working with a new assistant, and would have to spin each new tab up from scratch. So we didn’t go down rabbit holes, I created a startup paragraph that was copied into the start of every chat. The last sentence was always “don’t do anything until I say DONE”. That way AI wouldn’t start thinking before I had uploaded all necessary files and given final directions.
Using AI as a coding companion stirred the code so fast that I stopped commenting the code. As code has stopped churning, I need to go back and add comments. Amazingly, AI is actually pretty good at writing comments, and I will let it have first pass.
AI focuses on what it’s told to focus on. If you ask about dead code in a file, it will tell you all of the dead code in a file. This actually happened after working in numerous files for months without AI hinting at a problem.
AI is designed to mimic people. As we were working, ChatGPT would sometimes complement me on the clarity or simplicity of my designs, the quality of my interface, my clear coding style, etc. This really wasn’t a problem, but distracts from the primary task – writing code. AI, like all social programs, is designed to keep people’s attention. My goal is to build software. Sometimes these tasks don’t totally match.
I learned that AI has limited resources the hard way. Generally, I would start a chat by giving AI the necessary files, and clear directions for a small sub feature leading towards a goal. A minute later I would receive new code with clear edits. If possible, after testing, I would provide another small goal, code would be returned and we would repeat. Early on I found that AI would inconsistently start to create trash code.
AI is a token generating and processing engine. Unfortunately, the token stack is finite. Further, when the token stack gets close to full, early parts of the conversation are condensed or abstracted, losing detail. (This is much better than 7 months ago when the top of the chat was just deleted.) The problem is AI (ChatGPT in this case) doesn’t let you know it’s happening. A human abstracts conversation in an intelligent manner. AI does it by brute force.
Software development is very token stack

[truncated]

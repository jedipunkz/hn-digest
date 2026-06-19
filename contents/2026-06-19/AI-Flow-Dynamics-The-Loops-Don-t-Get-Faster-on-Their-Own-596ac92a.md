---
source: "https://intentful.ueberproduct.de/p/ai-flow-dynamics-the-loops-dont-get"
hn_url: "https://news.ycombinator.com/item?id=48597736"
title: "AI Flow Dynamics – The Loops Don't Get Faster on Their Own"
article_title: "AI Flow Dynamics - The Loops Don’t Get Faster On Their Own"
author: "flail"
captured_at: "2026-06-19T13:18:29Z"
capture_tool: "hn-digest"
hn_id: 48597736
score: 1
comments: 0
posted_at: "2026-06-19T12:19:54Z"
tags:
  - hacker-news
  - translated
---

# AI Flow Dynamics – The Loops Don't Get Faster on Their Own

- HN: [48597736](https://news.ycombinator.com/item?id=48597736)
- Source: [intentful.ueberproduct.de](https://intentful.ueberproduct.de/p/ai-flow-dynamics-the-loops-dont-get)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T12:19:54Z

## Translation

タイトル: AI フロー ダイナミクス – ループは勝手に速くなることはない
記事のタイトル: AI フロー ダイナミクス - ループは勝手に高速化されない
説明: AI により、コード生成が安価かつ迅速になりました。コードを検証済みの値に変換する 3 つのフィードバック ループは、依然として古い速度で実行されます。そのギャップを管理することが、現在、成熟した AI 製品開発を意味しています。

記事本文:
AI フロー ダイナミクス - ループは勝手に速くなることはない
AI フロー ダイナミクス - ループは自然に高速化するわけではありません
AI により、コード作成が安価かつ迅速になりました。コードを検証済みの値に変換する 3 つのフィードバック ループは、依然として古い速度で実行されます。そのギャップを管理することが、現在、成熟した AI 製品開発を意味しています。
AI のおかげで、コードを安価かつ高速に作成できるようになりました。レビュー、市場測定、顧客吸収という 3 つのフィードバック ループがコードを価値に変えます。特に顧客の吸収は速くなりません。
私たちは、単一行のコード変更を安価かつ迅速に出荷できるようにするために CI/CD の構築に 15 年を費やしました。現在、同じパイプラインを通じて数千行のコストゼロの変更を送り込んでいます。制約は書き込みからレビュー、測定、吸収へと移行しましたが、パイプラインはほとんど変わっていません。
小さなバッチは、ループが閉じられている場合にのみ役に立ちます。 CI/CD は技術的なループを閉じました。私たちは市場ループを気にすることはほとんどありませんでした。
仕事はフロー規律であり、制限ではありません。すべてを構築しますが、各ループへの入力を平準化し（ストップライトシステムを設定）、多くのオプションを構築して少数を出荷し（オプションのストーミング）、顧客向けの変更を測定できるものと顧客が吸収できるものに制限します。
成熟したチームは、出力を検証するループに出力を測定します。残りは出荷できるので出荷します。あなたの選択です。
AI フロー サイエンス: 3 つのループ、すべてがセットアップを高速化するわけではありません
嬉しいことに、今では私が読むよりも早くコードが届くようになりました。エージェントは、私がこれまで 1 週間かかった作業を 20 分で書き上げます。永遠に忙しいビーバーは、私に 4,000 行を与え、私がそれらを承認するのを待っています。必要に応じて、6 つのバリエーションに 5 つの機能を追加: いつ完成したかを判断できます。たとえ始めても、あまり考える必要はありません。いいね。私たちは 15 年をかけて、コード作成のコストを削減しました。コードを完成させるには、

また、コード作成の前後のすべてのペースを決めるものでもあります。
「The Intentful Company」をお読みいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
過去 20 年間にわたって私たちが築いてきた前提
ソフトウェアは、ハードウェアとは異なり、柔軟です。発送後の交換も何度でも格安で可能です。この特性は、過去 20 年間に製品分野が学び、実行されたすべての基礎となっています。
その結果、物理的な部品では、主に実際の顧客の注文 (つまり、製品 OS が検証済み) に基づいて正しい設計を行った後、金属を切断し、CNC を起動し、高価な生産ラインでボタンを押します。ハードウェアには、ほとんどの場合、生産が開始される前に、経済的で検証済みの安定したモデルが存在します。しかし、ソフトウェアは主に、ユーザーが何を必要とするか、何が機能するか、何がそれ自体に見合ったものになるかといった仮定とアイデアに基づいて構築されています。そして、作業が開始される時点では、それらの仮定が真実であるかどうかはわかりません。それらはチェックする必要があります。 (オーダーメイドのソフトウェアを除いて、実際には正しい注文後に納品することが契約になっています。) 十分難しいですが、別のゲームです。ソフトウェアの世界の残りの部分は、別の野獣です。マーケティング、販売、研究開発、製品開発に至る巨大な組織全体が、顧客という奇妙な動物が実際に何を必要としているのかを理解しようとしています。
チェックは 2 つの瞬間に行われます。まずは建てる前から。コードはお金と時間の面で高価なので、それに賭ける前に、その仮定がどの程度真実であるかを尋ねる価値があります。そして第二に、リリース後、私たちが出荷したものが本当に誰かにとって価値を生み出すのか、請求額を支払うのか、採用されるのか、行動を変えるのかなどを調べるためです。各チェックはループです。何かを実行し、何が起こるかを観察し、調整し、再び実行します。
これが、エリック・リース氏がリーン・スタートアップの「構築、測定、

「学びなさい」、リーンは常によりエレガントに (そしてより広範囲の結果をもたらします) OODA ループと呼ばれます。
その時代の中心的な教訓は、バッチが小さいほど、よりクリーンなコード、より迅速な修正、より優れた、より正確な市場シグナルなど、これらすべてが改善されるということでした。理想的には、小さなバッチごとに 1 回のみ変更するため、何かが変更されたときに、なぜ変更する必要があったのかがわかります。そして、それが実際に行われた場合 (出力に結果はありましたか?)。これらすべての知識にもかかわらず、私たちは中心となる要件をほとんど無視しました。つまり、たとえ小さな小さなバッチであっても、それを囲むループが存在し、閉じられている場合にのみ役立つということです。測定ループを使用しない小さな変化では学習は行われません。
2 つの内側のループは、たとえ IT 部門の人たちの活動を監視するためだけであっても (これらのループに多大な労力が費やされている最も考えられる理由です) 閉じられることがよくありましたが、3 番目のループはほとんど無視されました。また、上位の決定者に頼ることになります。
CI/CD は技術的なループを閉じました。内側のループは、編集、実行、結果の確認という数秒間の開発者独自のサイクルです。外側のループは、数時間単位の統合サイクル (統合、テスト、レビュー、リリース) です。継続的インテグレーションによりテストが自動化されました。継続的デリバリーによりリリースが自動化されました。ループの実行は十分に安価になり、単一行の変更だけでも出荷する価値があるほどになりました。小さなバッチが功を奏し、正確さに関するフィードバックがすぐに返されました。
考えを持ち続けてください。私たちは過去 10 年から 15 年かけて、最小のコード変更が「無料」、低コスト、低影響になるように最適化してきました。小規模なリリースではトランザクション コストが低いため、最小限の変更をリリースできます。
しかし、市場ループはほとんどの組織で開かれたままでした。これは最も遅いもので、リリース、顧客の行動の測定、学習、調整は数週間から数四半期かかります。すべての遅行指標。ただし水晶玉の交換が必要です。一部のチーム

は、時折 A/B テストや Pendo のような製品分析ツールなどを使用してこれに取り組みました。そのループを閉じるのは、実際には本当に難しいです。それには、技術、根性、忍耐、そしてドーパミンがすぐにヒットするのではなく、遅い満足感を得る多くの心理学が必要です。そのため、私たちは少量の正しいバッチを迅速に出荷することには非常に優れていましたが、ほとんどの人は、実際に出荷する価値があるかどうか、そしてどれが出荷する価値があるかを判断するのが非常に苦手のままでした。
実際に変更された無料コード
生成されたコードは高速で、ほぼ無料です。以前は遅くて高価だった変更の作成というステップは崩壊し、小さなエージェントの魔神がポーチにある iPhone から遠隔操作して私たちの代わりにそれを行ってくれます。忍耐力はクロードが仕事から戻ってくるのを待っています。
私たちが無視したのは、これによって生産のバッチサイズが増加するということです。エージェントは、人がかつて小さな時間に費やしていた時間に、一貫した大きな変化をもたらします。生産はもはや人間のタイピング速度に合わせて行われることはなく、その人間のペースが数十年にわたってシステム全体のリズムを設定していました。
私たちは 10 年から 15 年かけてこのパイプラインを構築しました。その目的は 1 つです。それは、1 つの小銭を安く発送することです。 AI は現在、何千行もの同じパイプライン変更を無料で一度にプッシュしており、それらは私たちが反対の問題のために構築したマシンを介して直接実行されます。可能な限り最小のバッチに合わせて調整されたインフラストラクチャは、一晩で生成された最大のバッチを無料で叩き込まれていますが、何も問題はありません。チェックは青になり、デプロイメントは発火します。変わったのは流れていくもののサイズだけで、それは人間に届くまで目に見えない。そのステップに何十億もの自動テストを追加するインフラジョブが完了していない限り、人間は PR 負荷で苦しむ可能性があります。少なくとも、実際に PR が何であるかを定義する必要があります。でもそれは

申し訳ありませんが、簡単な部分です。
そのため、制作は高速化されましたが、ループは高速化されませんでした。 1 駅の速度を上げて残りを放っておくと、次の駅の前に仕事が山積みになる、とゴールドラットは 40 年前に描写しました。 1 つのステーションを高速化すると、制約が次のステーションに再配置され、ここではレビュー、測定、市場吸収の 3 つの場所に再配置されます。
「The Intentful Company」をお読みいただきありがとうございます。この投稿は公開されているので、自由に共有してください。
変更は人間が読むよりも早く人間のレビューに到達し、人間は一定の割合でレビュ​​ーするようになりました。そしてそれはその通りであるはずです。 Shapiro スケールによるフロンティア AI 拡張コーディングでは、レベル 3 から「生産性」が大幅に向上します。これは、「コードを行レベルでレビューしない」ことを意味します。ここで、バッチを小さく保ち、本番ではなくレビューがスループットを制限するまでレビューの数が増加します。ジャムを次の段階であるレビューに再配置しました。バッチを大きくし、各バッチが人間が実際に検査できるよりも多くのコードに影響を与えると、人間によるレビューが機能しなくなり、欠陥が本番環境にまで波及することになります。それらの間のどこかに、キューを安定させ、レビューを正直に保つバッチ サイズがあります。これは、私たちが慣れ親しんでいる小バッチの最適値よりも大きく、依然として限界があります。
その境界を上げる方法は、CI/CD がデプロイ時に行ったのと同じ、チェックを自動化することです。型、コントラクト、プロパティベースのテスト、生成されたテスト ケースなど、自動化できる機械でチェック可能なあらゆる小さな要素により、人間の負担が軽減されます。その後、人間は時間をかけて意図をチェックすることができます。その機能は何をするべきなのか、何をするように設計されているのか、その仮定は少なくとも局所的には当てはまりますか。その間、チェッカー エージェントは数十億の側面で行の正しさを確認します。そのインフラを構築する

e レイヤーは、10 年前に展開パイプラインに対して行った投資の現在のポスト AI バージョンです。部分的には古い投資が報われていますが、これまで人間が保証していたマイクロ/コードレベルのテストインフラを追加するようになりました。
2 番目のステップは、CI と CD の融合を停止することです。私たちはいつもそれらを混同しすぎていました。統合とリリースが同じものである場合、変更は構築された瞬間に有効になります。それが芸術であり、私たちは実際に誇りに思っていました。後ほどオプション ストーミングの意味を読んでいただければ、接続の切断にどれだけ依存しているかがわかります。CI ステップの後には判断/フィルターが必要です。単純に多くを構築しすぎて、顧客が変更を許容できる最大値が小さいだけになる可能性があります。それらを分離すると、単一の顧客に連絡することなく変更を構築、実行、検査できます。このビルド済みだがリリースされていない状態は、ダウンストリームのすべてのチェックおよび決定ポイントであり、その後に続くもののほとんどはそのフィルターに依存します。
リリースによって顧客の行動に変化が生じたと言うには、それをノイズから分離するのに十分な信号が必要です。必要なサンプルは、見ようとしている効果の逆二乗に応じて増加します。 10% のベースラインで 2% の相対変化を検出するには、通常の信頼度でバリアントごとに約 350,000 の観測値が必要です。ロン・コハヴィは、実際に影響を与える実験がいかに少ないかを示す仕事をしました。ほとんどは力不足で、誤解を招く仮定や推測に基づいて誤解を招いています。
四半期内のトラフィックの量と、明確な帰属を測定できる変更の数はどれくらいですか?その数は固定されており、多くの場合、驚くほど小さいです。無料のコードはその数値に影響を与えません。ただし、顧客に送信する変更の数は増加し、スピードアップ係数に比例します。出荷するものと測定できるものとの間のギャップは爆発的に広がります。すべて

測定できる範囲を超えて生産しても、その効果を特定したり測定したりする方法がないまま、顧客に届くことになります。つまり、学習はゼロです。そして今、私たちは判断、意見、好み、つまり過去 10 年間取り除きたかったものに戻ってきました。現在では、AI に対する最後の防衛線として使用されています。私はそれを疑う。好みによって決まるか、出荷されて実際に評価されることはありません。
構築されたオプションを、出荷前に市場のモデル (シミュレーション、プロキシ メトリック、実際のアーティファクトに反応するパネル) に対してテストでき、そのためにライブ トラフィックを費やす必要はありません。それによって評価できる範囲が広がります。したがって、市場のモデルでは予測できない、真に新しい賭けのために、実際の実験を数回保存しておきます。
ブレークポイント 3: 市場の吸収
外部市場ループは、対処するのが最も難しいものです。最も難しいボトルネック。それは統計を超えています。指標は遅れて曖昧になり、帰属を見つけるのは難しく、あらゆるノイズの中で信号がほとんどありません。顧客は限界までの変化のみを受け入れますが、その限界は根本的な変化にも定義されています。影響の少ない小さな変更は、より簡単かつ迅速に受け入れられ、測定も容易になります。より大きな変化はその逆です。したがって、より小さな変更に偏りがあります。管理が簡単で、測定も簡単です。

[切り捨てられた]

## Original Extract

AI made code production cheap and fast. The three feedback loops that turn code into validated value still run at their old speed. Managing that gap is what mature AI product development now means.

AI Flow Dynamics - The Loops Don’t Get Faster On Their Own
Subscribe Sign in AI Flow Dynamics - The Loops Don’t Get Faster On Their Own
AI made code production cheap and fast. The three feedback loops that turn code into validated value still run at their old speed. Managing that gap is what mature AI product development now means.
AI made writing code cheap and fast. Three feedback loops turn code into value — review, market measurement, customer absorption. Especially customer absorption does not get faster.
We spent fifteen years building CI/CD to be able to ship a single line code change cheap and fast. Now we pump thousand-line, zero-cost changes through the same pipeline. The constraint moved from writing to reviewing, measuring, and being absorbed, but the pipeline is mostly unchanged.
Small batches only ever helped wherever a loop was closed. CI/CD closed the technical loops; we rarely bothered with the market loop.
The work is flow discipline , not a limit: build everything, but level the input to each loop (set up a stop light system), build many options and ship few (option storming), and limit customer-facing change to what you can measure and what customers can absorb.
Mature teams measure the output into the loops that validate it. The rest ship because they can. Your choice.
AI Flow Science: Three loops, not all speed up Setup
Gladly, now the code arrives faster than I can read it. An agent writes in twenty minutes what used to take me a week. The eternal busy beaver, it gives me four thousand lines, and waits for me to approve them. 5 Features in 6 variants if I want: I can judge when it’s done. Don’t have to overthink if I even get started. Cool. We spent fifteen years making code creation cheaper. And getting the code done, was also what set the pace for everything before and after code creation.
Thanks for reading The Intentful Company! Subscribe for free to receive new posts and support my work.
The premise we built on for the last 20 years
Software is malleable in a way hardware is not. You can change it after it ships, cheaply, again and again. That property is the foundation under everything the product discipline learned and executed in the last twenty years.
As a consequence, in the physical part, we cut the metal, start the CNC, push the button on the expensive production line, after a correct design, mostly based on actual customer orders (i.e. the product OS already validated). Hardware most times has a solid economic, validated model, before the production is started. Software, though, is mostly built on assumptions and ideas - about what users need, about what will work, about what will pay for itself. And those assumptions are not known to be true when the work starts. They have to be checked. (Except in bespoke software, where actually the deal is to deliver after the correct order.) Difficult enough, but a different game. The rest of the software world is a different beast. A whole huge org from Marketing and Sales over R&D to Product Development trying to figure out what the strange animal, the customer, actually needs.
The check happens at two moments. First, before building. Code is expensive, in money and time, so it is worth asking how true an assumption is before betting on it. And, secondly, after release, to find out whether the thing we shipped really creates value for anyone, pays the bill, gets adopted, changes behavior, etc. Each check is a loop: do something, watch what happens, adjust, go again.
That’s what Eric Ries coded into Lean Startup’s “Build, Measure, Learn”, lean always more elegantly (and with more far reaching consequences) called the OODA loop. You get it.
The central lesson of that era was that smaller batches make all of this better - cleaner code, faster correction, better, more precise market signal. Each small batch ideally just one change, so when something changes you know why it had to be changed. And if it actually did (was there outcome to the output?). Despite all that knowledge, we mostly ignored the core requirement: even a tiny, small batch only helps if the loop around it exists and is closed. A small change without the measurement loop produces no learning.
While the two inner loops often got closed, even if only to monitor what these guys in IT do (the most probable reason why is much effort is spent on these loops), the third one was mostly ignored. It would also fall back on the higher deciders.
CI/CD closed the technical loops. The inner loop is the developer’s own cycle - edit, run, see the result - in seconds. The outer loop is the integration cycle - integrate, test, review, release - in hours. Continuous Integration automated the testing; Continuous Delivery automated the release. Running the loops became cheap enough that a single-line change was worth shipping on its own. Small batches paid off, and feedback on correctness came back fast.
Hold the thought: We spent the last ten to fifteen years, optimising for the smallest code change to be “free”, low cost, low impact. Lower transaction cost for small releases, so we can release the smallest changes.
The market loop, though, stayed open in most orgs. It is the slowest one - release, measure how customers behave, learn, adjust - over weeks and quarters. All lagging indicators. But required to replace the crystal ball. Some teams approached this with the occasional A/B test and / or a product-analytics tool like Pendo, and little more. Closing that loop is actually, genuinely hard. It takes tech, grit, patience and a lot of late gratification psychology vs the instant dopamine hit. So while we got very good at shipping small correct batches quickly, most of us remained really bad at knowing if and which were actually worth shipping.
What free code actually changed
Generated code is fast and nearly free. The step that used to be slow and expensive - writing the change - has collapsed and the little agent genies do it for us, remote controlled from the iPhone on the porch. The patience is waiting for Claude to come back from its work.
What we ignored is that this rises the batch size of production. An agent produces a large, coherent change in the time a person once spent on a small one. Production no longer paces itself to how fast a human can type, and that human pace had set the rhythm of the whole system for decades.
We spent ten or fifteen years building this pipeline for one purpose: to make a single small change cheap to ship. AI now pushes the same pipeline changes of a thousand lines at once, at no cost, and they run straight through the machine we built for the opposite problem. The infrastructure that was tuned for the smallest possible batch is now being hammered with the largest batch produced over night for free, and nothing in it complains: the checks go green, the deploys fire. The only thing that changed is the size of what flows through, and that stays invisible until it reaches a human. The human might suffer under the PR load as long as the infra job of adding a gazillion automated tests to that step is not done. At least to forces us to define what a PR actually is. But that’s the easy part, sorry to say.
So production got faster and the loops did not. Speed up one station and leave the rest alone, and the work piles up in front of the next one, which Goldratt described forty years ago. Speeding up one station relocates the constraint to the next one, and here it relocates to three places: review, measurement, market absorption.
Thanks for reading The Intentful Company! This post is public so feel free to share it.
Changes now reach human review faster than humans can read them, and a person reviews at a fixed rate. And it’s supposed to be that way. Frontier AI augmented coding as per the Shapiro scale sees radical “productivity” increase from level three, which means “I don’t review my code at the line level”. Now, keeping the batches small and the number of reviews climbs until review, not production, caps the throughput - you have relocated the jam to the next stage: review. Making the batches big and each one affecting more code than a person can actually examine, means human review stops working and defects get through to production. Somewhere between those is a batch size that keeps the queue stable and review honest. It is larger than the small-batch optimum we are used to, and it is still bounded.
The way to raise that boundary, is the same thing that CI/CD made on deployment: automate the checking. Types, contracts, property-based tests, generated test cases - every tiny machine-checkable aspect that can be automated, takes the load off the human. Then the human can spends time checking intent - does the feature what it’s supposed to do, what we designed it to do, does the assumption hold at least locally - while the checker agents confirm correctness the lines in a gazillion of aspects. Building that infrastructure layer is the current post-AI version of the investment we made in deployment pipelines a decade ago. In parts, the old investment pays off, but we now add test infra on the micro / code level that was guaranteed by the human until now.
The second step is to stop fusing CI and CD. We always conflated them too often. When integration and release are the same thing, a change goes live the moment it is built. That was the art and we were actually proud of. Read what I mean by option storming later and you’ll see how dependent we are now on breaking the connection: There is judgement / filter after the CI step required. We can simply build too much and the customer has only a small max tolerance of change. Pull them apart and a change can be built, run, and inspected without reaching a single customer. That built-but-unreleased state is the check and decision point for everything downstream, and most of what follows depends on that filter.
To say that a release caused a change in customer behavior, you need enough signal to separate it from noise. The sample you need grows with the inverse square of the effect you are trying to see. On a 10% baseline, detecting a 2% relative change takes roughly 350,000 observations per variant at ordinary confidence. Ron Kohavi did the work to show how few real experiments have any impact; most are underpowered and mislead, based on misleading assumptions, guesses.
How much traffic do you and what is the the number of changes you can measure with clear attribution in a quarter? The number is fixed, and often surprisingly small. Free code has no influence on that number. But the the number of changes you send to the customer rises - linear to your speedup factor. The gap between what you ship and what you can measure explodes. Everything you produce beyond what you can measure reaches customers with no way to attribute and measure the effect. So: zero learning. And now we’re back to judgement, opinion, taste, the very thing we wanted to get rid of for the last ten years. Now it’s used as the last line of defence against the AI. I doubt it. It gets decided by taste, or it ships and is never really evaluated.
You can test a built option against a model of the market before it ships - a simulation, a proxy metric, a panel reacting to the real artifact - and spend no live traffic doing it. That extends how much you can evaluate. So you save your handful of real experiments for the genuinely new bets, the ones a model of the market cannot predict.
Breakpoint 3: market absorption
The outer market loop, is the toughest one to handle. The hardest bottleneck. It’s beyond statistics. Metrics are coming in late and fuzzy, attribution is hard to find, little signal in all the noise. Customers accept change only up to a limit and that limit is also defined to the radically of change. Little changes with little impact are accepted easier and faster, are easier to measure. Bigger changes, the opposite. Hence a bias for smaller changes: easier to manage, easier to measure, we often do

[truncated]

---
source: "https://qa-financial.com/barclays-on-ai-testing-telemetry-and-kill-switches/"
hn_url: "https://news.ycombinator.com/item?id=49256336"
title: "Barclays on AI testing, telemetry and kill switches"
article_title: "Barclays: AI testing, telemetry and kill switches - QA Financial"
author: "bryanrasmussen"
captured_at: "2026-08-11T11:37:54Z"
capture_tool: "hn-digest"
hn_id: 49256336
score: 1
comments: 0
posted_at: "2026-08-11T11:11:04Z"
tags:
  - hacker-news
  - translated
---

# Barclays on AI testing, telemetry and kill switches

- HN: [49256336](https://news.ycombinator.com/item?id=49256336)
- Source: [qa-financial.com](https://qa-financial.com/barclays-on-ai-testing-telemetry-and-kill-switches/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T11:11:04Z

## Translation

タイトル: バークレイズ、AI テスト、テレメトリ、キル スイッチについて
記事のタイトル: バークレイズ: AI テスト、テレメトリー、キル スイッチ - QA Financial
説明: これは、バークレイズが自律 AI エージェントの展開をどのように準備しているかを調査する 2 部構成の QA Financial シリーズの第 1 部です。

記事本文:
バークレイズ: AI テスト、テレメトリ、キル スイッチ - QA Financial
コンテンツにスキップ
QA ファイナンシャル フォーラム ロンドン | 2026 年 9 月 16 日 |チケットを予約する
スポンサーになる
リンクトイン
ツイッター
[メールで保護されています]
検索
バークレイズ、AI テスト、テレメトリー、キルスイッチについて
これは、英国の銀行 Barclays が世界で最も規制の厳しい銀行環境の 1 つに導入する自律型 AI エージェントをどのように準備、テスト、検証しているかを検証する 2 部構成の QA Financial シリーズの最初の記事です。
パート 1 では、本番環境に対応した AI を形成するテスト、可観測性、ガバナンスの規律について説明します。第 2 部では、評価、運用保証、およびソフトウェア品質が成熟した AI 戦略においてますます中心となる理由を検討します。
銀行は何十年もかけて、ソフトウェアを本番環境に導入する前にテストする方法を改良してきました。自律型 AI エージェントは現在、その規律の再考を迫られています。
もはや課題は、モデルが正しい答えを生成することを単に証明することではありません。 AI システムを観察、管理、評価し、必要に応じて停止できることを実証することがますます重要になっています。
この変化は金融サービス全体でますます顕著になってきています。今年初め、英国の金融行動監視機構は、バークレイズ、UBS、ロイズ・バンキング・グループなどをAIライブ・テスト・プログラムに参加させ、規制当局がモデルの精度を超えて、ガバナンス、監視、人間による監視、ライブ環境内での運用管理に注目していることを示唆した。
バークレイズのプリンシパル AI エンジニアであるアンディ・マクマホン氏は、銀行も同じように考える必要があると考えています。先月、マクマホン氏はポッドキャスト The Brave Technologist で、英国の最も古い銀行の 1 つが、世界で最も厳しく管理されているテクノロジーの 1 つで自律型 AI にどのようにアプローチしているかについて、貴重な垣間見せました。

環境。
しかし、同氏は、ますます高性能になった AI モデルについて議論するのではなく、テレメトリ、可観測性、評価、許可、生産管理、ソフトウェア保証など、ソフトウェア テストの専門家がよく知っている主題に繰り返し戻りました。
それは強調点の顕著な変化である。エージェント型 AI に関する会話の多くは依然として機能を中心に展開していますが、マクマホン氏が焦点を当てていたのは自信でした。
エージェントがデモンストレーションでタスクを完了できるという自信はありませんが、本番環境に導入された後も安全に動作し続けるという自信があります。
インテリジェンスの前に可観測性
ほとんどの組織に AI エージェントについて質問すると、会話はすぐに推論、計画、自律性の話に変わります。マクマホンは別の場所からスタートする。
「エージェント システムの本番化に関する 3 つの異なる主要な懸念事項について説明しました」と彼は説明しました。 「つまり、1 つは可観測性でした。」それは偶然の選択ではありません。
何十年にもわたって、ソフトウェア チームはログ、メトリクス、モニタリングに依存して、運用システムがどのように動作するかを理解してきました。マクマホン氏は、自律型 AI には、より深いレベルでのみ、まったく同じエンジニアリング規律が必要であると主張しました。
「エージェントとエージェント ソリューションのテレメトリ メトリクス、ログ、トレースをどのように取得するかがわかりました」と彼は共有しました。
これらのいくつかの言葉は、より大きなモデルやより洗練された推論に関するいかなる議論よりも重要であることが判明する可能性があります。
テレメトリがなければ、エンジニアはエージェントがなぜ特定の決定に至ったのかを理解できません。痕跡がなければ、失敗を再現するのに苦労します。可観測性がなければ、ソフトウェアが運用環境に到達した瞬間にテストは事実上終了します。
それは重要な変化を表しています。ソフトウェアがリリースされてもテストが停止することはなくなりました。代わりに、本番環境自体がテストのライフサイクルの一部となり、自律システムがどのように動作するかに関する証拠を継続的に生成します。

実際の動作条件下での平均値。
この哲学は、バークレイズのより広範なエンジニアリングの方向性を反映しています。同銀行は昨年、エンジニアリング業務全体にわたる可観測性を拡大しながら、開発の早期にテストをどのように組み込んだかを説明した。
パフォーマンス エンジニアリングのグローバル リード Manik Sikka 氏は、ソフトウェアが顧客に届く前に問題を特定するには、可観測性とシフトレフト テストが不可欠になったと主張しました。
マクマホン氏のコメントは、同じ原則が自律型 AI にも拡張されつつあることを示唆しています。
機能だけでなく動作をテストする
従来のソフトウェア テストは主に、アプリケーションが実行するように設計された機能を実行するかどうかに焦点を当ててきました。
自律システムでは、別の課題が生じます。エンジニアは、条件が変化したときにエージェントが許容範囲内で動作し続けるという確信も必要です。
マクマホン氏にとって、安全性は AI の存在リスクに関する抽象的な議論ではありません。それは実際的な工学的な問題です。 「展開しているものが安全であることを確認し、実際にどのように動作するかを理解するにはどうすればよいでしょうか?」
「実際に行動する」というフレーズは、エンタープライズ AI のテストにおける決定的な課題の 1 つになる可能性があります。実験室での評価により、システムが予想されるシナリオにどのように反応するかが明らかになります。制作では予期せぬものが登場します。
マクマホン氏が作戦管理にも同様に重点を置くのはこのためです。 「エージェントに対してさまざまなレベルの制御をどのように適用できますか? キルスイッチなどのことはプログラムでどのように実行しますか?」
彼はキルスイッチを緊急メカニズムとして見るのではなく、エンジニアリング設計の別のコンポーネントとして提示します。これにより、ソフトウェア品質そのものの定義が広がります。
品質はもはや機能の正しさやパフォーマンスだけで測られるものではありません。ますますレコが含まれます

真実性、制御性、そして自律型ソフトウェアが予期せぬ動作をしたときに安全に介入する能力。
規制当局の間でも同様の考えが現れ始めている。英国の FCA は、AI 保証がモデル自体を超えて、展開コンテキスト、ガバナンス、評価手法、入出力制御を含むことを明らかにしました。
実際、テストは、アルゴリズムが単独で適切に機能することを単に証明するのではなく、AI システム全体が信頼できるという証拠になりつつあります。
マクマホン氏は、プロダクション AI が制約なしで動作すべきであるという考えにも同様に否定的でした。自律エージェントの準備ができているかとの質問に対し、同氏は「準備ができていると思う」と答えた。 「確かに、それが本当に範囲内での自主性であるという主張はあると思います。」
この区別は銀行業界内では特に重要になります。 「AIエージェントに白紙のようなものを与えて、やりたいことをやらせることはできない」と彼は主張した。 「すべてのツールやすべてのシステムへのアクセスを許可することはできません。」
代わりに、自律性を設計する必要があります。マクマホン氏は、「しかし、これらのシステムを制限して、私が話していた制御を採用することは可能です。採用できるさまざまな許可構造があります。」と述べました。
「エージェントが自律的に、ただし範囲内で行動できるように制御する方法はたくさんあります。それが重要です。」と彼は続けました。
「AI エージェントに、自分が望むことを完全に実行させることはできません。すべてのシステムへのアクセスを与えることはできません。」
これらのコメントは、QA を従来アーキテクチャ、アイデンティティ管理、セキュリティが占めていた領域に押し込みます。テストでは、エージェントが正しい答えを生成するかどうかだけを問う必要はなくなりました。
また、エージェントが許可されたシステムのみにアクセスしたか、定義された権限の境界内に留まったか、権限を尊重したかどうかも尋ねられます。

オーバーナンス規則を遵守し、あらゆる決定について十分な証拠を生成しました。
多くの点で、これらの質問は、サイバー レジリエンスとオペレーショナル リスクですでによく知られている質問に似ています。違いは、ソフトウェア自体が計画と意思決定を行えるようになり、それらの制御が現実世界の条件下で動作し続けることを証明する重要性が高まっていることです。
「古いエンジニアリング」は依然として重要
マクマホン氏の最も強力な議論の 1 つは、最も単純な議論の 1 つでもあります。それは、エージェント AI をめぐる興奮にもかかわらず、ソフトウェア エンジニアリングの基礎が突然時代遅れになったわけではないということです。
同氏は、あまりにも多くの組織が AI をまったく新しいテクノロジー スタックを要求しているかのように扱っていると考えています。 「私にとって最も大きなことは、これが大きな革命であると皆が思っていることです。そして、大きな進化も存在することを忘れているのです。」と彼は語った。
この観察は失望するというよりも、むしろ安心させるものに聞こえるはずです。銀行は、DevOps、MLOps、自動テスト、ID 管理、コンテナ化、監視、可観測性への投資に何年も費やしてきました。マクマホン氏は、これらの分野は依然として AI システムにあらゆる意味で関連していると主張した。
「DevOps の実践のすべて、MLOps のすべて、データ サイエンスのすべて、従来の優れたアプリ開発のすべて…それらすべてがまだ残っています。」
「私たちはまだコンテナ化を行っています。私たちはまだKubernetesで実行しています。私たちは依然としてアイデンティティプロバイダーを使用しています。私たちは依然として可観測性ソリューションを使用しています」と彼は言いました。
AI によって新たな懸念が生じることは認めたが、それらは置き換えられるものではなく、追加されるものである。 「同じテクノロジーに新しいスピンを加えたものです。」
これは ID と権限にも同様に当てはまります。 「今、私はエージェントのアイデンティティの概念を持っていますが、それがどのようにして人ごとに特定のものを持っているのかを考える必要があります。」

ミッションとそれが他のものに代わってどのように機能するのか。」
マクマホン氏は、すべてを一から再構築しなければならないと思い込む誘惑があると警告した。 「最大の危険は、私がすべてを投げ出して最初からやり直さなければならないとみんなが思っていることだと思います。」
代わりに、AI エンジニアリングは、自律的な動作に対応できるように拡張しながら、実証済みのソフトウェア エンジニアリングの実践に基づいて構築する必要があります。
これは、自動化、継続的テスト、可観測性、生産監視に対する既存の投資がさらに価値のあるものになることを意味します。エージェント AI は、まったく新しい QA 規律を作成するのではなく、銀行がすでに備えている QA 規律の範囲を拡大しています。
規制によってより優れたソフトウェアが生み出されるでしょうか?
追加の制御やガバナンス要件を享受できる開発者はほとんどいません。しかしマクマホン氏は、厳しく規制された銀行内で働くことで最終的にはエンジニアリングの品質が向上すると主張する。 「そこには素晴らしい問題がたくさんあります。」
さらに重要なことは、「最終的にはより良い製品を生み出すことができる」と彼は続けました。
理由は簡単です。銀行システムは、顧客に提供される前に、並外れたレベルの精査にさらされます。 「彼らは非常に多くの精査を通過する必要があるため、あなたは本当に確信しているようです。」
マクマホン氏はさらに、「これを使って量産に入ると、これは十分にテストされていると思われるでしょう。」と付け加えた。
このフレーズは、「厳しい戦いに耐えた」という言葉で、金融サービス内で AI 保証がどのように進化しているかをうまく表現しています。銀行は、AI システムが動作することを実証するだけでなく、実際のユーザー、変化する条件、実際の生産データにさらされた後も安全に動作し続けることを実証することをますます期待されています。
その考えは規制当局がとっている方向性と密接に一致している。
英国の FCA の AI ライブ テスト イニシアチブは、明らかにモデルを単独で評価することを超えています。代わりに、ガバナンス、展開コンテキスト、

人間の監視、評価技術、運用制御により、AI を単なるアルゴリズムではなく運用システムとして効果的に扱います。
ソフトウェアテストの専門家にとって、これは重要な変化を表しています。テストは、出力をチェックすることよりも、複雑なシステムがライフサイクル全体を通じて安全性、回復力、責任を持ち続けるという証拠を生成することが重要になってきています。
おそらく、マクマホン氏のインタビューからの最大の教訓は、自律型 AI によってソフトウェア テストの重要性が薄れるわけではないということです。それにより、その範囲が大幅に広がります。
可観測性、テレメトリ、権限、評価、ガバナンス、および運用監視は、デプロイメント後に追加される運用上便利なものではなくなりました。これらは、AI システムの設計方法に最初から影響を与える中核的なエンジニアリング要件になりつつあります。
これは、金融サービス全体ですでに見られる広範な傾向を反映しています。バークレイズは過去 2 年間、エンジニアリング組織全体でシフトレフトのテスト、自動化、可観測性を強化してきました。一方、規制当局は企業に対し、実験室での検証だけに頼るのではなく、AIが実際の環境で安全に動作するという証拠を提出するよう求めている。
これらの発展は共に同じ方向を向いています。品質エンジニアリングは、ソフトウェアの検証から自律システムの保証へと進化しています。
問題はもはや、AI エージェントが適切かどうかという単純な問題ではありません。

[切り捨てられた]

## Original Extract

This is the first in a two-part QA Financial series examining how Barclays is preparing autonomous AI agents for deployment

Barclays: AI testing, telemetry and kill switches - QA Financial
Skip to content
QA Financial Forum London | 16 September 2026 | BOOK TICKETS
Become a sponsor
Linkedin-in
Twitter
[email protected]
Search
Barclays on AI testing, telemetry and kill switches
This is the first instalment in a two-part QA Financial series examining how UK bank Barclays is preparing, testing and validating autonomous AI agents for deployment inside one of the world’s most highly regulated banking environments.
Part one explores the testing, observability and governance disciplines shaping production-ready AI. Part two will examine why evaluation, production assurance and software quality are increasingly central to any mature AI strategy.
Banks have spent decades refining how they test software before it reaches production. Autonomous AI agents are now forcing them to rethink that discipline.
The challenge is no longer simply proving that a model produces the correct answer. Increasingly, it is about demonstrating that an AI system can be observed, governed, evaluated and, if necessary, stopped.
That shift is becoming increasingly visible across financial services. Earlier this year, the UK’s Financial Conduct Authority brought Barclays, UBS, Lloyds Banking Group and others into its AI Live Testing programme , signalling that regulators are looking beyond model accuracy towards governance, monitoring, human oversight and operational controls inside live environments.
Andy McMahon, Principal AI Engineer at Barclays , believes banks need to think the same way. Speaking on The Brave Technologist podcast last month, McMahon offered a rare glimpse into how one of Britain’s oldest banks is approaching autonomous AI inside one of the world’s most tightly controlled technology environments.
Yet rather than discussing ever more capable AI models, he repeatedly returned to subjects that software testing professionals know well: telemetry, observability, evaluation, permissions, production controls and software assurance.
It is a notable change in emphasis. Much of the conversation around agentic AI still revolves around capability, while McMahon’s focus was confidence.
Not confidence that an agent can complete a task in a demonstration, but confidence that it will continue behaving safely once deployed into production.
Observability before intelligence
Ask most organisations about AI agents and the conversation quickly turns to reasoning, planning and autonomy. McMahon starts somewhere else.
“I covered three different core concerns for productionizing agent systems,” he explained. “So one was observability.” That is not an accidental choice.
For decades, software teams have relied on logs, metrics and monitoring to understand how production systems behave. McMahon argued autonomous AI requires exactly the same engineering discipline, only at a much deeper level.
“So you know how do you capture telemetry metrics, logs, traces for your agent and agentic solutions,” he shared.
Those few words may prove more significant than any discussion about larger models or more sophisticated reasoning.
Without telemetry, engineers cannot understand why an agent reached a particular decision. Without traces, they struggle to reconstruct failures. Without observability, testing effectively ends the moment software reaches production.
That represents an important shift. Testing no longer stops when software is released. Instead, production itself becomes part of the testing lifecycle, continuously generating evidence about how autonomous systems behave under real operating conditions.
That philosophy mirrors Barclays’ broader engineering direction. Last year, the bank described how it had embedded testing earlier into development while expanding observability across its engineering practices.
Global Lead for Performance Engineering Manik Sikka argued that observability and shift-left testing had become essential for identifying issues before software reached customers.
McMahon’s comments suggest those same principles are now being extended into autonomous AI.
Testing behaviour, not just functionality
Traditional software testing has largely focused on whether applications perform the functions they were designed to perform.
Autonomous systems introduce a different challenge. Engineers also need confidence that agents will continue operating within acceptable boundaries when conditions change.
For McMahon, safety is not an abstract discussion about existential AI risk. It is a practical engineering problem. “How do you make sure what you’re deploying safe and you can understand how it’s going to behave in the wild?”
That phrase, ‘behave in the wild’, may become one of the defining testing challenges of enterprise AI. Laboratory evaluations reveal how systems respond to expected scenarios. Production introduces unexpected ones.
That is why McMahon places equal emphasis on operational controls. “How can you employ different levels of control over your agents? How do you do things like kill switches programmatically?”
Rather than viewing kill switches as emergency mechanisms, he presents them as another component of engineering design. This broadens the definition of software quality itself.
Quality is no longer measured solely by functional correctness or performance. It increasingly includes recoverability, controllability and the ability to intervene safely when autonomous software behaves unexpectedly.
The same thinking is beginning to appear among regulators. Britain’s FCA has made clear that AI assurance extends beyond the model itself, encompassing deployment context, governance, evaluation techniques and input and output controls.
In effect, testing is becoming evidence that an entire AI system can be trusted rather than simply proving that an algorithm performs well in isolation.
McMahon was equally dismissive of the idea that production AI should operate without constraints. “I think we are,” he said when asked whether autonomous agents are ready. “I think though there is definitely a case to be made that it’s really autonomy within bounds.”
That distinction becomes especially important inside banking. “You can’t just give an AI agent like carte blanche to do what it wants,” he argued. “You can’t give it like access to all of your tools, all of your systems.”
Instead, autonomy has to be designed. “But you can bound these systems and employ that control I was talking about,” McMahon stated “there are different permission structures you can employ.”
“There’s just lots of ways of giving control so that the agent can act autonomously but within bounds and that’s important,” he continued.
“You can’t just give an AI agent like carte blanche to do what it wants. You can’t give it access to all of your systems.”
These comments push QA into territory traditionally occupied by architecture, identity management and security. Testing no longer asks only whether an agent produces the right answer.
It also asks whether the agent accessed only authorised systems, remained inside defined permission boundaries, respected governance rules and generated sufficient evidence for every decision it made.
In many respects, the questions resemble those already familiar from cyber resilience and operational risk. The difference is that the software itself is now capable of planning and making decisions, increasing the importance of proving those controls continue to operate under real-world conditions.
‘Old engineering’ still matters
One of McMahon’s strongest arguments is also one of the simplest: despite the excitement surrounding agentic AI, software engineering fundamentals have not suddenly become obsolete.
Too many organisations, he believes, are treating AI as if it demands an entirely new technology stack. “The biggest thing for me is everyone thinking this is such a revolution,” he said, “that they forget there’s a huge piece of evolution as well.”
That observation should sound reassuring rather than disappointing. Banks have spent years investing in DevOps, MLOps, automated testing, identity management, containerisation, monitoring and observability. McMahon argued those disciplines remain every bit as relevant for AI systems.
“All of the practices from DevOps, all of the stuff from MLOps, all of the stuff from data science, all the stuff from just traditional good app development… all of that is still there.”
“We’re still containerising stuff. We’re still running in Kubernetes. We’re still using identity providers. We’re still using observability solutions,” he said.
AI introduces new concerns, he acknowledged, but they are additions rather than replacements. “It’s the same technologies with new spins.”
That applies equally to identity and permissions. “Now I have a concept of an agent identity and I have to think how does that have certain permissions and how does it act on behalf of other things.”
The temptation, McMahon warned, is to assume everything must be rebuilt from scratch. “I think the biggest danger is everyone thinks I have to throw everything out and start again.”
Instead, AI engineering should build on proven software engineering practices while extending them to accommodate autonomous behaviour.
That means existing investment in automation, continuous testing, observability and production monitoring becomes even more valuable. Rather than creating an entirely new QA discipline, agentic AI is expanding the scope of the one banks already have.
Can regulation produce better software?
Few developers enjoy additional controls or governance requirements. Yet McMahon argues that working inside a heavily regulated bank ultimately improves engineering quality. “There are a lot of cool problems there.”
More importantly, he continued, “it does at the end make a better product.”
The reason is straightforward. Banking systems face extraordinary levels of scrutiny before reaching customers. “They have to go through so much scrutiny that you’re sort of you’re really sure.”
McMahon added: “You go into production with this thing, you’re like, this has been battle tested.”
That phrase, battle tested’, neatly captures how AI assurance is evolving inside financial services. Banks are increasingly expected to demonstrate not only that AI systems work, but that they continue working safely once exposed to real users, changing conditions and live production data.
That thinking aligns closely with the direction regulators are taking.
Britain’s FCA’s AI Live Testing initiative explicitly moves beyond evaluating models in isolation. Instead, it looks at governance, deployment context, human oversight, evaluation techniques and operational controls, effectively treating AI as an operational system rather than simply an algorithm.
For software testing professionals, that represents an important shift. Testing is becoming less about checking outputs and more about generating evidence that complex systems remain safe, resilient and accountable throughout their lifecycle.
Perhaps the biggest lesson from McMahon’s interview is that autonomous AI is not making software testing less important. It is making it significantly broader.
Observability, telemetry, permissions, evaluations, governance and production monitoring are no longer operational nice-to-haves added after deployment. They are becoming core engineering requirements that influence how AI systems are designed from the outset.
That reflects a wider trend already visible across financial services. Barclays has spent the past two years strengthening shift-left testing, automation and observability across its engineering organisation. Meanwhile, regulators are asking firms to produce evidence that AI behaves safely in live environments rather than relying solely on laboratory validation.
Together, those developments point in the same direction. Quality engineering is evolving from verifying software into assuring autonomous systems.
The question is no longer simply whether an AI agent com

[truncated]

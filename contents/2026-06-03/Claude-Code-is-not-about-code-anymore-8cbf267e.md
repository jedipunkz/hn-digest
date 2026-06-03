---
source: "https://blog.vtemian.com/post/claude-code-isnt-about-code/"
hn_url: "https://news.ycombinator.com/item?id=48370086"
title: "Claude Code is not about code anymore"
article_title: "Claude Code Isn't About Code | Vlad Temian - Centrist Tech Optimist"
author: "vtemian"
captured_at: "2026-06-03T00:46:05Z"
capture_tool: "hn-digest"
hn_id: 48370086
score: 2
comments: 0
posted_at: "2026-06-02T13:34:52Z"
tags:
  - hacker-news
  - translated
---

# Claude Code is not about code anymore

- HN: [48370086](https://news.ycombinator.com/item?id=48370086)
- Source: [blog.vtemian.com](https://blog.vtemian.com/post/claude-code-isnt-about-code/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:34:52Z

## Translation

タイトル: Claude Code はもはやコードに関するものではありません
記事のタイトル: Claude Code はコードに関するものではありません |ヴラド・テミアン – 中道派テクノロジー楽観主義者
説明: クロード コードを最大限に活用している人は次のとおりです。

記事本文:
クロード コードはコードに関するものではありません |ヴラド・テミアン – 中道派テクノロジー楽観主義者
戻る クロード コードはコードに関するものではありません
Claude Code はコーディング ツールではありません。現時点でこれを最もよく使用している人は、コードにはほとんど触れていません。彼らは工場現場、法務チーム、人事、財務、運営を運営しています。彼らはそれをソフトウェアとは考えていません。彼らはそれを仕事をするものだと考えています。
それを見るのに2か月かかりました。 4 月に、私は 20 以上のツールを出荷し、収益はゼロで、人間が使用するツールからエージェントが使用するツールへ価値が移動するのを観察したと書きました。私は 3 つの言葉で終わりました。答えはありません。
「Inside YC’s AI Playbook」では、Pete Koomen、Garry Tan、Jared が YC が内部でどのように再構築したかを説明します。 YC は常に独自のソフトウェアで実行され、すべての企業、すべての創設者、すべての金融取引、社内 CRM のすべてのメモなど、重要なものすべてを保持する単一の Postgres データベース上にあります。
1 年前、このループはあなたが働いたことのあるすべての会社と同じように見えました。財務チームは、ソフトウェア エンジニアにワークフロー (仕訳入力の予約、価格ラウンドの記録) を説明します。エンジニアは、ワークフローをエンコードする専用の決定論的なソフトウェアを構築します。そうしたら彼らはそれを返してくれるでしょう。クーメンの言葉で言えば、「非効率」です。
そこで彼らはループを廃止したのです。彼らは財務ソフトウェアを構築する代わりに、財務チームにエージェントとツールのレジストリを提供しました。すべてを変えた最初のツールは、恥ずかしいほど単純なものでした。実稼働データベースへの読み取り専用 SQL アクセスです。ジャレッドはそれを構築しましたが、ルールを破っているように感じ、深夜にそれを押し出しました。とてもうまくいきました。
彼らは約 20 個のツールから始めました。講演の時点では 350 を超えていました。各チームが独自に追加しています。財務帳簿の仕訳。オフィスアワーやイベントを管理する人もいます。そして同じレジストリ

YC の内部エージェントと個々のラップトップで実行されるクロード コードの両方にフィー​​ドを提供します。
その結果は、ジェヴォンズのパラドックスの教科書的な事例でした。複雑な質問をすることで他の人の SQL 時間を数時間費やすことがなくなったとき、人々は単に古い質問に早く答えるだけではありませんでした。彼らははるかに多くの質問をし、はるかに難しい質問をしました。質問するのが安くなったからです。
それを機能させる逆転
これが単なる内部ツールではなく、本当の変化である理由は、クーメン氏の以前のエッセイ「AI Horseless Carriages」で説明されています。彼の批判は、Gmail の電子メール ライターのように、開発者がすべての制御を保持し、ユーザーからそれを隠している機能として既存のソフトウェアに組み込まれた AI に対するものでした。
要点は一文であり、ゆっくり読む価値があります。 AI ネイティブ ソフトウェアは、AI をラップする決定論的ソフトウェアではなく、決定論的ツールをラップするエージェントです。制御は開発者からユーザーに移ります。
それを裏返せば、戦略は崩れてしまいます。インターフェースの構築をやめてしまいます。慎重に設計されたフローを備えた別の財務 SaaS、別の ERP 画面、別の社内運用ダッシュボードを構築する必要はなくなります。それは馬車で、エンジンがボルトで固定された古い形です。
代わりに、小規模で決定論的で範囲の広いツールを構築し、それらをレジストリに配置し、それらを構成するエージェントを人間が操作できるようにします。工場長、弁護士、人事部長。彼らは、エンジニアリングのバックログを待つことなく、その瞬間に必要なものを平易な言葉で構築できるようになります。インターフェースが制約でした。取り外したものが商品となります。
これが4月に私が模索していた答えです。使い捨てツールが繁栄する環境は、ダッシュボードのホスティング、プレビュー、請求、またはそれだけではありません。これは、統合されたコンテキストのレイヤーに、共有ツール レジストリ、さらにハーネスを加えたものです。

ets エージェントがそれらを構成します。そのレイヤー内のツールは設計上使い捨てになります。レジストリとコンテキストは永続的なものです。 YC は専用のシステムを構築し、それを「底上げ」と呼んでいました。新入社員は、6 か月の入社研修の代わりに、ツールと成績証明書を通じて組織内の優秀な人材の組織的な知識を吸収します。
入場料はかかりますが、それについての話は正直です。年間 10,000 ドルから 100,000 ドルの間、場合によってはそれ以上をトークンに費やす必要があります。ギャリー・タンの枠組みでは、これで 1 回限りのタイム ワープが得られるということです。つまり、あなたは今 2028 年に住むことができ、2 年後には 100 分の 1 の費用がかかり、誰もがそれを行うことになります。それは楽観的な見方だ。さて、もう一つです。
なぜあなたの経済は10年もこの状況を感じないのか
ソフトウェアがこれほど革新的なのであれば、どこが成長するのでしょうか?同じ週にまったく別の部屋から投下された講演に答えがあります。
スタンフォード大学の経済学者チャド・ジョーンズは、「A.I. と私たちの経済の未来」の中で、弱いリンクという 1 つのアイデアに基づいて構築されたモデルを提案しています。チェーンの強さは、最も弱いリンクの強さによって決まります。ビジネスの成功には多くのタスクを完了する必要があり、20 個のリンクのうち 17 個を信じられないほど強力にしたとしても、そのチェーンはまだ触れていない 3 個によって制限されます。
その洞察により、エレガントで少し気の抜けたフォーミュラが生まれます。タスクを無限に自動化すると、GDP に占めるそのタスクの割合がほぼ同じだけ GDP が増加します。ソフトウェアはGDPの約2%です。したがって、明日ソフトウェアが無限に豊富になり無料になったとしたら、私たちは約 2% 裕福になるでしょう。 100倍豊かではありません。 2パーセント。それ以外はすべてボトルネックのままです。
彼のキラーチャートはそれを具体的にします。私たちは、1970 年代のトランジスタの 1 億倍を搭載したコンピューターを持ち歩いています。しかし、GDP に占めるコンピューティング能力の割合は 2000 年の約 4.5% でピークに達し、その後は 3 分の 1 の 3% に低下しました。価格

崩壊は量の爆発を支配します。豊富なリンクはシェアを失います。価値は希少なままのものに流れます。
それが私のマニフェストの根底にあるメカニズムであり、私が管理できたよりもさらに厳密に述べられています。ソフトウェアが安価で豊富に存在するため、使い捨てツールはビジネスになりません。利益は、コンテキスト、判断、オーケストレーション層、何を尋ねるべきかを決定する人間といった、希少な弱いリンクから生じます。価値観が動いたように感じました。ジョーンズは、どこに移動するのか、なぜ移動するのかを方程式で示します。
彼は領収書も持っています。 2016年、ジェフリー・ヒントンは、放射線科医は5年以内にAIに取って代わられるため、放射線科医の訓練をやめるべきだと述べた。 10 年後、放射線科医は増え、彼らの給料も上がりました。ジョブはタスクの束です。そのうちの 75% が自動化され、残りの 25% が希少で価値があり、高収入の部分となります。弱いつながりは、人間がまだ立たなければならない場所です。
ここに緊張感があり、それが重要な点だと思います。 YCは、あなたは2028年に今でも生きられると言っています。ジョーンズ氏は、爆発には3年ではなく30年かかると言う。
スケールは異なりますが、どちらも正しいです。自らの脆弱なリンク (ロックダウンされたコンテキスト、指揮統制の承認、サイロ化されたツール、質問とその回答の間のエンジニアリングのボトルネック) を喜んで破壊する単一の高信頼組織内では、10x は現実のものであり、今すぐ利用できます。 YCがその証拠です。経済全体にわたって、弱いつながりはどこにでもあり、そのほとんどが人間であり、それらすべてを強化するには数十年かかります。どちらのことも同時に真実です。
そのギャップはチャンスですが、それには有効期限があります。組織自体の弱点を取り除いた組織は、現在も別の社内ダッシュボードを運用している競合他社よりも何年も早く運用できるようになりました。彼らがより優れたモデルを持っているからではありません。誰もが同じモデルを持っています。彼らは止まるから

インターフェイスを構築し、レジストリの構築を開始しました。
4 月に私は、このツールは終わりではないと書きましたが、それが何なのか名前を付けることができなかったため、そのままにしておきました。今ならできるよ。
専用のダッシュボード、インターフェイス、フローを構築するのはやめましょう。財務用に別の SaaS を構築したり、運用用に別の社内ツールを構築したり、誰も楽しんでいない別の ERP を構築したりするのはやめてください。小さな決定論的なツールを構築し、レジストリに登録してエージェントに渡し、人々に運転させます。インターフェースは足場でした。エージェントはツールをラップします。人間がリードを握ります。
ツールは決して製品ではありませんでした。ループとは、人間、エージェント、およびツールのレジストリであり、その瞬間に必要なものをすべて作成し、その瞬間が過ぎるとそれを破棄します。クロード・コードはコードに関するものではありません。そんなことは決してなかった。

## Original Extract

The people getting the most out of Claude Code aren

Claude Code Isn't About Code | Vlad Temian - Centrist Tech Optimist
Back Claude Code Isn't About Code
Claude Code isn’t a coding tool. The people using it best right now barely touch code. They run factory floors, legal teams, HR, finance, operations. They don’t think of it as software. They think of it as the thing that does the work .
That took me two months to see. In April I wrote that I’d shipped 20+ tools, made zero revenue, and watched the value move from tools humans use to tools agents use . I ended on three words: I don’t have the answer.
In “Inside YC’s AI Playbook” , Pete Koomen, Garry Tan, and Jared walk through how YC rebuilt itself internally. YC has always run on its own software, sitting on a single Postgres database that holds everything that matters: every company, every founder, every financial transaction, every note in their internal CRM.
A year ago, the loop looked like every company you’ve worked at. The finance team would describe a workflow (booking journal entries, logging priced rounds) to software engineers. The engineers would go build purpose-built, deterministic software encoding that workflow. Then they’d hand it back. Koomen’s word for it: inefficient .
So they killed the loop. Instead of building finance software, they gave the finance team agents and a registry of tools. The first tool that changed everything was almost embarrassingly simple: read-only SQL access to the production database. Jared built it, felt like he was breaking the rules, and pushed it out late at night. It worked extremely well.
They started with about 20 tools. At the time of the talk, there were more than 350. Every team adds their own. Finance books journal entries. Others manage office hours and events. And the same registry feeds both YC’s internal agents and Claude Code running on individual laptops.
The effect was a textbook case of Jevons paradox . When asking a complex question stopped costing several hours of someone else’s SQL time, people didn’t just answer the old questions faster. They asked far more questions, and far harder ones, because asking became cheap .
The inversion that makes it work
The reason this is a real shift and not just internal tooling is captured in Koomen’s earlier essay, “AI Horseless Carriages” . His critique was of AI bolted onto existing software as a feature, like Gmail’s email writer, where the developer keeps all the control and hides it from the user.
The crux is one sentence, and it’s worth reading slowly. AI-native software is the agent wrapping deterministic tools, not deterministic software wrapping an AI. Control shifts from the developer to the user.
Flip that around and the strategy falls out of it. You stop building the interface. You stop building another finance SaaS, another ERP screen, another internal operations dashboard with its carefully designed flows. Those are horseless carriages: the old shape, with an engine bolted on.
Instead you build small, deterministic, well-scoped tools, you put them in a registry, and you let humans drive the agent that composes them. The plant manager, the lawyer, the HR lead. They get to build whatever they need in the moment, in plain language, without waiting on an engineering backlog. The interface was the constraint. Removing it is the product.
This is the answer I was groping for in April. The environment where disposable tools thrive is not hosting and previews and billing dashboards, or not only that. It’s a layer of unified context plus a shared tool registry plus a harness that lets agents compose them. The tools inside that layer are disposable by design. The registry and the context are the durable thing. YC built a private one and called it raising the floor: a new hire absorbs the institutional knowledge of the best people in the org through the tools and transcripts, instead of six months of ramp.
There’s a price of admission, and the talk is honest about it. You have to be willing to spend somewhere between $10,000 and $100,000 a year on tokens, sometimes more. Garry Tan’s framing is that this buys you a one-time time warp: you get to live in 2028 now, and in two years it costs a hundredth as much and everyone does it. That’s the optimistic read. Now for the other one.
Why your economy won’t feel this for a decade
If software is this transformative, where’s the growth? A talk that dropped the same week, from a very different room, has the answer.
In “A.I. and Our Economic Future” , Stanford economist Chad Jones offers a model built on a single idea: weak links. A chain is only as strong as its weakest link. Business success requires completing many tasks, and if you make 17 of 20 links incredibly strong, the chain is still limited by the three you didn’t touch.
That insight produces an elegant and slightly deflating formula. Infinite automation of a task raises GDP by roughly that task’s share of GDP. Software is about 2% of GDP. So if software became infinitely abundant and free tomorrow, we’d be about 2% richer. Not 100 times richer. Two percent. Everything else stays a bottleneck.
His killer chart makes it concrete. We carry computers with 100 million times the transistors we had in the 1970s. Yet the share of GDP paid to computing power peaked at about 4.5% in 2000 and has since fallen by a third, to 3%. The price collapse dominates the quantity explosion. The abundant link loses its share. Value flows to whatever stays scarce.
That is the mechanism underneath my manifesto, stated more rigorously than I managed. Disposable tools don’t become businesses because software is now the cheap, plentiful link. The returns accrue to the scarce weak links: the context, the judgment, the orchestration layer, the human deciding what to ask. I felt that the value had moved. Jones gives you the equation for where it moves and why.
He has the receipts too. In 2016, Geoffrey Hinton said we should stop training radiologists because AI would replace them within five years. A decade later we have more radiologists, and they’re paid more. Jobs are bundles of tasks. Automate 75% of them and the remaining 25% become the scarce, valuable, well-paid part. The weak link is wherever the human still has to stand.
Here is the tension, and I think it’s the whole point. YC says you can live in 2028 today. Jones says the explosion takes 30 years, not three.
They’re both correct, at different scales. Inside a single high-trust organization that is willing to tear down its own weak links (locked-down context, command-and-control approvals, siloed tools, the engineering bottleneck between a question and its answer), the 10x is real and available right now . YC is the proof. Across the whole economy, the weak links are everywhere and mostly human, and strengthening all of them takes decades . Both things are true at once.
That gap is the opportunity, and it has an expiration date. The org that removes its own weak links now gets to operate years ahead of competitors who are still commissioning another internal dashboard. Not because they have a better model. Everyone has the same models. Because they stopped building interfaces and started building a registry.
In April I wrote that the tool is not the end, and I left it there because I couldn’t name what was. I can now.
Stop building specialized dashboards, interfaces, and flows. Stop building another SaaS for finance, another internal tool for operations, another ERP nobody enjoys. Build small deterministic tools, put them in a registry, give them to agents, and let people drive. The interface was scaffolding. The agent wraps the tools. The human holds the leash.
The tool was never the product. The loop is: a person, an agent, and a registry of tools, composing whatever the moment needs and throwing it away when the moment passes. Claude Code isn’t about code. It never was.

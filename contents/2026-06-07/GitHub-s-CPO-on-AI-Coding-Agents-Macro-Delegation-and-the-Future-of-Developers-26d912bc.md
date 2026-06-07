---
source: "https://www.turingpost.com/p/mario-rodriguez-github-ai-coding-agents-copilot"
hn_url: "https://news.ycombinator.com/item?id=48432586"
title: "GitHub's CPO on AI Coding Agents, Macro-Delegation, and the Future of Developers"
article_title: "GitHub AI Coding Agents: Copilot, Agent PRs & Macro-Delegation"
author: "olgava"
captured_at: "2026-06-07T07:30:33Z"
capture_tool: "hn-digest"
hn_id: 48432586
score: 1
comments: 0
posted_at: "2026-06-07T07:07:26Z"
tags:
  - hacker-news
  - translated
---

# GitHub's CPO on AI Coding Agents, Macro-Delegation, and the Future of Developers

- HN: [48432586](https://news.ycombinator.com/item?id=48432586)
- Source: [www.turingpost.com](https://www.turingpost.com/p/mario-rodriguez-github-ai-coding-agents-copilot)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T07:07:26Z

## Translation

タイトル: GitHub の CPO、AI コーディング エージェント、マクロ委任、開発者の将来について語る
記事のタイトル: GitHub AI コーディング エージェント: コパイロット、エージェント PR、マクロ委任
説明: GitHub CPO マリオ・ロドリゲスが、2025 年 12 月に AI コーディング エージェントが変更された理由と、マクロ委任、AX、およびコパイロットが開発者にとって何を意味するのかについて語ります。

記事本文:
GitHub AI コーディング エージェント: コパイロット、エージェント PR、マクロ委任 この Web サイトは Cookie を使用しています
詳細については、当社のプライバシー ポリシーと利用規約をお読みください。
Turing Post を検索 ログイン サインアップ ホーム タグ アーカイブについて アップグレード 厳選されたシリーズ 厳選されたシリーズ エージェントのワークフロー
🎙️GitHub のマリオ ロドリゲスが AI コーディング エージェント、コパイロット、開発者の将来について語る
何億もの開発者と新興クラスの AI エージェントが GitHub 上で共同構築を始めると何が変わるでしょうか?
ほとんどの開発者ツールは、人間と人間のコラボレーションのために構築されています。このインタビューでは、GitHub の最高製品責任者である Mario Rodriguez が、マクロ委任やエージェント生成の PR から Copilot、AX、開発者の将来に至るまで、AI コーディング エージェントがどのように GitHub を新しいエージェント ネイティブ エンジニアリング システムに向けて推し進めているかについて説明します。
GitHub の最高製品責任者である Mario Rodriguez 氏は、2025 年 12 月に訪れた変曲点について説明しています。モデルはついに、常にモデルを修正することなくエージェントに「マクロ委任」できるほど十分に優れたものになりました。
その後、GitHub はどうなったのでしょうか?コミット、PR、アクション、セキュリティ スキャン全体での記録の高速化と、GitHub とは何かについて根本的に再考します。
このインタビューでは、次のことについて説明します。
2025 年 12 月に AI コーディング エージェントが変更された理由
エージェント生成アクティビティの増大に伴う GitHub の規模の課題
マクロ委任とミクロ委任
UI → UX → AX、またはエージェントの経験
コパイロット、キャンバス、人間とエージェントのコラボレーション
使用量ベースの課金とトークンの規律
エージェントが生成した PR と制作品質
副操縦士がパイロットではなく副操縦士のままである理由
また、「開発者」の再定義、なぜ (効率ではなく) 創造が人類の進歩を促すのか、GitHub が初めての構築者とピカソレベルの職人の両方に同じ連続体でどのようにサービスを提供するつもりなのかについても話します。見てください！
スー

YouTube チャンネルに登録するか、Spotify / Apple でインタビューを聞いてください。
参考のためにトランスクリプトを用意しましたが、完全な体験はビデオにあります。そしていつものように、「いいね」とコメントをしてください。これは、YouTube での成長と、より多くの洞察を提供するのに役立ちます。
クセニア:
こんにちは、マリオ。チューリング ポストによる推論ショーへようこそ。ご参加いただきありがとうございます。
マリオ:
迎えてくれてありがとう。外は素晴らしい天気です。昨日は少し曇っていましたが、今日は太陽も出てきて本当に良かったです。
クセニア:
素晴らしい日です。しかし、GitHub に行きましょう。
Samelweb によると、GitHub には月間 6 億 3,000 万人以上の訪問者がいます。これは、開発と実験のための膨大な表面積です。
では、エージェントのワークフローが実際に機能し始めた 2025 年後半以降、GitHub では何が変わったのでしょうか?何に気づきましたか?
マリオ:
興味深いですね。 12 月頃、私たちが気づいた重要な点の 1 つは、モデルの機能が大幅に向上したということでした。
それまでは、エージェントへのマクロ委任のようなものを実行しようとすると、常にそれを修正する必要がありました。 「いいえ、あなたはこの道を選んだのです。その道を選ぶべきではなかったのです。」または「あなたは別のことをしました。代わりにあれをすべきでした。」と言うでしょう。それは幼児を相手にするのと少し似ていて、「違う、違う、そっちに行かないで、代わりにこうしてください。ここは安全です。」と言いました。
12 月の期間で変わったのは、「さあ、遊んでください – 安全です」と実際に言えるようになり、非常に高品質の出力が得られるようになったということです。
私の意見では、これにより 2 つのことが明らかになりました。
まず、開発者のワークフローでは、ユーザーは大幅に多くのことをマクロ委任し、必要な場合にのみマイクロステアリングを行うことができます。そして、そのマイクロステアリングはイライラすることはありませんでした。 「なんてことだ、大量のトークンを無駄にしてしまった、今度は説明しなければならない」という感じではなかった

あなたが間違ったことをすべて。」むしろ、「よし、やったね。今度は私がループで協力して、より良いものを作りましょう。」という感じでした。修正プロセスではなく、反復的な作成プロセスになりました。
第 2 に、エージェントが自動化においてより自律的に実行されるようになると、実行時間が長くなる可能性があります。つまり、より良いタスクを彼らに与えることができるようになります。言い換えれば、タスクの ROI が向上しました。
その後、業界が追いつくにつれて、人々は 1 月の休暇から戻り、休暇後には落ち着きました。両方のことが同時に起こり始めました。より多くの個人開発者が、より強力なエージェント機能を備えたこれらの新しいモデルを使用し始め、自動化も進み始めました。
GitHub について考えると、開発ライフサイクル全体が対象となります。リポジトリを作成し、コードをリポジトリに取り込むだけではありません。私たちにも問題があります。プルリクエストもあります。何かを構築するためのアクションもあります。セキュリティーもございます。そしてそれらはすべて交差し、相互に構築されます。
したがって、より多くのコミットを取得すると、より多くの PR を取得できる可能性があります。より多くの PR を取得すると、より多くのアクションが実行されます。アクションの実行が増えると、より多くのセキュリティ スキャンが必要になります。すべてが複合します。
これらの数値の一部を公開しました。 3月だけでエージェントからのPRは1,700万件あったと思います。必要に応じて、さらに詳しい統計情報を調べることもできますが、すべてはそこから始まりました。おそらくジェンセン氏も基調講演でそのことについて言及したのを聞いたことがあるでしょう。私たちは皆、この全体的な加速を感じています。
変化したのは、より多くの人々がより速い速度でプラットフォームにやって来るようになったことです。これは、入場フロアが低くなった、または少なくとも以前よりも低くなったことが一因です。それについては後ほど詳しくお話します。
人間の観点から見ると、それはエキサイティングなことです。なぜなら、コミットを作成する人が増えると、PR を作成する人も増えるためです。

e アクションが実行され、すべてのサービスが記録的な加速を示しています。でっちあげの数字を使うと、5% の成長を期待していたのに、突然 3 倍になっています。
それはすごいことです。これは、私たちがこれらのエージェント ワークフローに真の価値を見出し始めていることを証明しています。それが、私たちがここに集まっている理由です。
GitHub でのスケールはどのように見えるか
クセニア:
その規模のトラフィックに関する主な問題は何ですか?
マリオ:
それは問題とは言えません。私はこれを、その規模での運用に伴う一連のエンジニアリング上の課題と呼んでいます。
そのうちのいくつかは明らかなものです。負荷が大きい場合は、その負荷を引き受けるためにさらに多くのマシンが必要になります。つまり、サーバーが増えるということです。私たちが現在取り組んでいる重要なことの 1 つは、Azure への負荷をさらに軽減することです。これは、データ センターの 1 つで拡張できる限界に基本的に達しているためです。より多くのリポジトリ負荷と PR 負荷をパブリック クラウドに移行することで、拡張を続けることができます。
これは、成長が 3 倍、5 倍、さらには 10 倍になったとしても、引き続き対応でき、単一の地域に制約されることがないため役立ちます。
次に、より広範なインフラストラクチャの問題があります。直接スタックの外にあるプロバイダーと通信する必要があります。西海岸のネットワーク インフラストラクチャが飽和状態になり始めたケースが 1 つありました。私たちはそのインフラストラクチャを所有していないため、それらのプロバイダーと協力して、より多くのトラフィックを計画する必要があることを伝える必要がありました。西海岸では多くの開発者活動が行われているため、エコシステム全体で協力して、すべてのエコシステムが負荷を処理できるようにする必要があります。
次に、典型的なスケーリング効果が得られます。このレベルでは、小さな問題でも多くの波及効果が生じる可能性があります。したがって、優れたエンジニアリングの基礎、つまりキャッシュ、新しいストレージ アプローチ、データの取得と提供のさまざまな方法にさらに多くの投資を行う必要があります。とても関わっています

エドエンジニアリングの仕事。しかし、それは本当にやりがいのあることでもあります。
新しい GitHub: 低層階、高天井
クセニア:
基本的に 2 つの異なる視聴者がいる今、GitHub の役割を再考していますか?
マリオ:
はい、それについてはいろいろ考えてきました。私は週末に GitHub のデザイン責任者である John Maeda と話し、実際に何が起こっているのか、そして GitHub がどのように進化する必要があるのか​​を探っていました。
彼は、MIT での設計の例えである、床が低く、天井が高いことを共有しました。それがまさに今起きていることだと思うので、とても気に入りました。
AI を使用したコーディングが行っていることは、ソフトウェア作成の敷居を下げることです。 AI、特にこれらのモデルはコーディングを好みます。これは、以前は手の届かなかった創作ツールに、より多くの人がアクセスできるようになったことを意味します。
GDP の成長について考えてみると、それは人間が物を作るから起こります。効率が良くなったからだけではありません。進歩は、誰かが他の人に価値を与えるものを作成したときに起こります。それが人類の進歩、つまりツールによる創造です。
現在、床を下げているところです。そして、私たちはまだ始まったばかりだと思います。あなたが言及した 6 億 3,000 万という数字は、次に来るものはさらに大きくなると思います。
時々モーツァルトのことを思い出します。おそらく同時に世界にはモーツァルトが他に 10 人ほどいたと思われますが、おそらく彼らはピアノを手に入れることができなかったのでしょう。今起こっていることは、より多くの人々にリーチできるようになるということです。確かに、私たちの中にはラップトップの前に座っている人もいますが、それはすでに世界では特権的な地位にありますが、モバイル、AI モデル、GitHub を通じて、創作の敷居を下げることができます。そしてそれは、世界でさらに多くのイノベーションが起こることを意味すると思います。
2 番目のことは、床を下げると同時に、床も高くすることです。

エイリング。プロの開発者、つまりすでに高度なスキルを備えた人々は、ますます優れたものを作成できるようになります。彼らはフロンティアを前進させることができるでしょう。
それも重要です。イノベーションは新規参入者だけから生まれるわけではありません。また、この技術の最前線にいる専門家からもたらされています。アインシュタインは相対性理論を開発することから始めたわけではありません。彼はこれまでの多くの物理学に基づいて構築しました。彼は飛躍する前に職人になった。
そこで、床を下げて、より多くの人を集め、より多くの人が専門家や職人になり、そして彼らが達成できることの上限を引き上げます。
あなたはそれを2つの異なる聴衆と呼んでいました。私はそれを連続体として考えています。
GitHub は、その連続体のエージェントネイティブのエンジニアリング システムになる必要があります。
GitHub の使命は、開発者のコ​​ラボレーションを通じて人類の進歩を促進することです。おそらく今では、開発者とエージェントのコラボレーションと言うべきかもしれません。私が毎日頭を悩ませているのは、どうやって床を下げ、どうやって天井を上げるかということです。はい、そのためには新しい GitHub が必要になる場合があります。
ソフトウェア開発におけるエージェント エクスペリエンス (AX) とは何ですか?
クセニア:
それはとても興味深いですね。もっと共有できれば – その会話は具体的なものにつながりましたか?
マリオ:
はい。 John のこのビジュアルでは、ユーザーがいて、その上にたくさんの障壁があり、ユーザーはそれぞれの障壁を飛び越えなければなりません。すべての障壁は UI のクリック、またはある種の処理ステップです。
私たちは、デザインの観点から見た新しい GitHub がどのようなものであるかについて話していました。そして私にとって、新しい GitHub はエージェント エクスペリエンスを備えたものです。
今日の私たちのコアプリミティブの多くは、人間と人間の協力に基づいています。次に、これらのプリミティブを人間とエージェントのコラボレーションに拡張する必要があります。しかし、プリミティブだけでは十分ではありません。 API レイヤーは、エージェント中心になるように進化する必要があります。そしてUXレイヤーは

進化も必要だろう。
私たちは Copilot アプリを発表しました。このアプリにはキャンバスのアイデアが導入されているので、とても楽しみにしています。キャンバスは AX、つまりエージェント エクスペリエンスの始まりだと思います。
UI がありますが、その UI はエージェントと双方向です。 UI はツールを公開し、エージェントはツールを読み取ることができ、UI に影響を与える可能性があります。つまり、エージェントに必要なことを伝えるだけで、エージェントは必要なことを実行できるということです。エージェントが呼び出す適切な API をすでに知っている場合、50 の不適切な設計の画面を飛び越える必要はありません。
しかし、素晴らしいことに、それは逆にも機能するということです。 UI を操作してエージェントに影響を与えることができます。 「要約」というボタンをクリックすると、エージェントがそれを受け取り、役立つ情報を返します。
そのため、エージェントと会話して何かが返ってくるのを待つだけだった古いモデルの代わりに、作成は双方向になりました。あなたにはアーティストのようにリアルタイムで何かを形づくるキャンバスがあり、エージェントがそれを形にするのを手伝ってくれます。
新しい GitHub は、双方向のエージェント エクスペリエンスを目的としていると私は考えています。
そして、その利点は、スペクトルの両端で機能することです。エージェントとチャットするだけで 20,000 回クリックしなくても作業を完了できるため、フロアが下がります。しかし、それは上限も高くなります。なぜなら、あなたが自分の媒体におけるピカソのように、高度なスキルを持っていれば、可能性を深く追求できるからです。

[切り捨てられた]

## Original Extract

GitHub CPO Mario Rodriguez on why December 2025 changed AI coding agents — and what macro-delegation, AX, and Copilot mean for developers

GitHub AI Coding Agents: Copilot, Agent PRs & Macro-Delegation This website uses cookies
Read our Privacy policy and Terms of use for more information .
Search Turing Post LOG IN SIGN UP HOME TAGS ABOUT ARCHIVE UPGRADE Curated Series Curated Series Agentic Workflow
🎙️GitHub’s Mario Rodriguez on AI Coding Agents, Copilot, and the Future of Developers
What changes when hundreds of millions of developers and a rising class of AI agents start building together on GitHub?
Most developer tools were built for human-to-human collaboration. In this interview, Mario Rodriguez, Chief Product Officer at GitHub , explains how AI coding agents are pushing GitHub toward a new agent-native engineering system: from macro-delegation and agent-generated PRs to Copilot, AX, and the future of developers.
Mario Rodriguez, Chief Product Officer at GitHub , explains the inflection point that hit in December 2025: models finally got good enough that you could "macro-delegate" to agents without constantly correcting them.
What happened to GitHub then? Record acceleration across commits, PRs, Actions, and security scans – and a fundamental rethink of what GitHub even is.
In this interview, we discuss:
Why December 2025 changed AI coding agents
GitHub’s scale challenge as agent-generated activity grows
Macro-delegation vs micro-delegation
UI → UX → AX, or agent experience
Copilot , canvases, and human-agent collaboration
Usage-based billing and token discipline
Agent-generated PRs and production quality
Why Copilot remains co-pilot, not pilot
We also talk about the redefinition of "developer," why creation (not efficiency) drives human progress, and how GitHub plans to serve both the first-time builder and the Picasso-level craftsman on the same continuum. Watch it!
Subscribe to our YouTube channel , or listen to the interview on Spotify / Apple
We prepared a transcript for reference, but the full experience is in the video . And as always: like and comment. It helps us grow on YouTube and bring you more insights.
Ksenia:
Hi, Mario, and welcome to Inference Show by Turing Post. Thank you for joining me.
Mario:
Thanks for having me. It’s a beautiful day outside. Yesterday was a little cloudy, but today the sun came out, so I’m really happy.
Ksenia:
It is a beautiful day. But let’s get to GitHub.
According to Similarweb , GitHub has over 630 million monthly visitors. That’s an enormous surface area for development and experimentation.
So since late 2025, when agent workflows really started working, what changed at GitHub? What did you notice?
Mario:
It’s interesting. Around December, one of the key things we noticed was that model capabilities took a real jump.
Before that, if you were going to do what I call a kind of macro-delegation to an agent, you constantly had to correct it. You’d say, “No, you took this path – you shouldn’t have taken that path,” or “You did this other thing – you should have done that instead.” It was a little bit like dealing with a toddler: “No, no, don’t go that way. Do this instead. Be safe over here.”
What changed in that December timeframe was that you could actually say, “Go ahead and play – it’s safe,” and you would get an output with very high quality.
In my opinion, that unlocked two things.
First, in the developer workflow, people could macro-delegate significantly more and then micro-steer only when they needed to. And that micro-steering didn’t feel frustrating. It didn’t feel like, “Oh my God, I just wasted a bunch of tokens, and now I have to explain everything you did wrong.” Instead, it felt more like, “Okay, you did that – now let me work with you in a loop to make it better.” It became an iterative creation process rather than a correction process.
Second, as agents started running more autonomously in automation, they could go longer. And that meant you could give them better and better tasks. In other words, the ROI of the task went up.
Then, as the industry caught up – people came back from break in January, got settled after the holidays – both things started happening at once. More individual developers started using these newer models with stronger agent capabilities, and more automation started happening too.
And if you think about GitHub, we span the whole development lifecycle. It’s not just the repo and getting code into the repo. We also have issues. We also have pull requests. We also have Actions to build things. We also have security. And they all intersect and build on each other.
So if you get more commits, you’ll probably get more PRs. If you get more PRs, you’ll get more Action runs. If you get more Action runs, you’ll need more security scans. Everything compounds.
We’ve published some of these numbers. I think in March alone we saw 17 million PRs from agents. We can get into more stats if you want, but everything really shot off from there. You probably heard Jensen mention that in the keynote too – we’re all feeling this overall acceleration.
What changed is that more and more people are coming into the platform at a faster rate, partly because the floor of entry is now lower – or at least lower than it used to be. We can talk about that more later.
From a human perspective, that’s exciting. Because as more people create commits, more people create PRs, and more Actions run, all of our services are seeing record acceleration. To use made-up numbers, if we were expecting 5% growth, suddenly we’re seeing 3x that.
And that’s amazing. It really proves that we’re starting to see real value from these agentic workflows – which is why we’re all here.
What Scale Looks Like at GitHub
Ksenia:
What are the main problems with traffic at that scale?
Mario:
I wouldn’t call it a problem. I’d call it a set of engineering challenges that come with operating at that scale.
Some of them are the obvious ones. If you have more load, you need more machines to take that load. That means more servers. One of the key things we’re doing right now is shedding more load into Azure, because we’ve basically hit the limits of how much we can grow in one of our data centers. By moving more repo load and PR load into the public cloud, we can keep expanding.
That helps because if growth goes 3x, 5x, or even 10x, we can still serve it, and we’re no longer constrained by a single region.
Then there’s the broader infrastructure question. You have to talk to providers outside your immediate stack. There was one case where network infrastructure on the West Coast started getting saturated. We don’t own that infrastructure, so we had to work with those providers and tell them they needed to plan for a lot more traffic. There’s a lot of developer activity on the West Coast, so we have to collaborate across the ecosystem to make sure all of it can handle the load.
Then you get the classic scaling effects: at this level, even a tiny issue can have a lot of ripple effects. So you have to invest much more in the fundamentals of good engineering – caching, new storage approaches, different ways of acquiring and serving data. It’s very involved engineering work. But it’s also really rewarding.
A New GitHub: Lower Floors, Higher Ceilings
Ksenia:
Do you reimagine the role of GitHub now that you basically have two different audiences?
Mario:
Yeah, I’ve been thinking a lot about that. I was speaking with John Maeda, our head of design at GitHub, over the weekend, and we were exploring what’s really happening and how GitHub needs to evolve.
He shared this design analogy they had at MIT: low floors, high ceilings . I loved it, because I think that’s exactly what’s happening now.
What coding with AI is doing is lowering the floor of entry into software creation. AI – and these models in particular – like to code. That means many more people now have access to tools for creation that were previously out of reach.
And if you think about GDP growth, it happens because humans create things. Not just because we become more efficient. Progress happens when someone creates something that someone else values. That’s how human progress moves forward: creation through tools.
Right now, we’re lowering the floor. And I think we’re just at the beginning. The 630 million number you mentioned – I think what comes next will be much larger still.
Sometimes I think about Mozart. There were probably ten other Mozarts in the world at the same time, but maybe they never had access to a piano. What’s happening now is that we can reach so many more people. Yes, some of us are sitting in front of a laptop – and that’s already a privileged position in the world – but through mobile, through AI models, through GitHub, we can lower the floor of entry for creation. And I think that means we’ll see significantly more innovation in the world.
The second thing is that while we lower the floor, we also raise the ceiling. Professional developers – people who are already highly skilled – are going to be able to create better and better things. They’ll be able to push the frontier forward.
That matters too. Innovation doesn’t only come from newcomers. It also comes from experts at the frontier of the craft. Einstein didn’t start by developing relativity. He built on a lot of prior physics. He became a craftsman before he made that leap.
So you lower the floor, get more people in, more of them become professionals or craftsmen, and then you raise the ceiling of what they can accomplish.
You called it two different audiences. I think of it more as a continuum.
GitHub needs to become the agent-native engineering system of that continuum.
The mission of GitHub is advancing human progress through developer collaboration. Maybe now we should say developer and agent collaboration. That’s what I obsess over every day: how do we lower the floor, and how do we raise the ceiling? And yes, that may require a new GitHub.
What Is Agent Experience (AX) in Software Development?
Ksenia:
That’s very interesting. If you can share more – did that conversation lead anywhere concrete?
Mario:
Yes. John has this visual where there’s a user, and then a bunch of barriers, and the user has to jump through each one. Every barrier is a UI click, or some sort of processing step.
We were talking about what a new GitHub looks like from a design perspective. And to me, a new GitHub is one that has an agentic experience .
A lot of our core primitives today are based on human-to-human collaboration. Now we need to extend those primitives into human-and-agent collaboration. But the primitives themselves won’t be enough. The API layer will need to evolve to become agent-centric. And the UX layer will need to evolve too.
We announced the Copilot app , and I’m really excited about it because it introduces this idea of canvases . I think canvases are the beginning of AX – agent experience.
You have a UI, but that UI is bidirectional with the agent. The UI exposes tools, the agent can read them, and it can affect the UI. That means I can simply tell the agent what I want, and it can do what it needs to do. I don’t have to jump through 50 poorly designed screens if the agent already knows the right API to call.
But the beautiful thing is that it also works in reverse. I can interact with the UI and affect the agent. I can click a button that says “summarize,” and the agent receives that and returns something useful.
So instead of the old model – where you only talked to the agent and waited for something to come back – now creation becomes bidirectional. You have a canvas where, like an artist, you’re shaping something in real time, and the agent is helping you shape it.
That’s what I think the new GitHub will be about: a bidirectional, agentic experience.
And the beauty of that is that it works on both ends of the spectrum. It lowers the floor because people can just chat with the agent and get things done without 20,000 clicks. But it also raises the ceiling, because if you’re highly skilled – like Picasso in your own medium – you can operate deeply in the canv

[truncated]

---
source: "https://human-in-the-loop.bearblog.dev/llms-are-eroding-my-software-engineering-career-and-i-dont-know-what-to-do/"
hn_url: "https://news.ycombinator.com/item?id=48434312"
title: "LLMs are eroding my software engineering career and I don't know what to do"
article_title: "LLMs are eroding my software engineering career and I don't know what to do | the human in the loop"
author: "poisonfountain"
captured_at: "2026-06-07T15:36:04Z"
capture_tool: "hn-digest"
hn_id: 48434312
score: 382
comments: 318
posted_at: "2026-06-07T12:49:29Z"
tags:
  - hacker-news
  - translated
---

# LLMs are eroding my software engineering career and I don't know what to do

- HN: [48434312](https://news.ycombinator.com/item?id=48434312)
- Source: [human-in-the-loop.bearblog.dev](https://human-in-the-loop.bearblog.dev/llms-are-eroding-my-software-engineering-career-and-i-dont-know-what-to-do/)
- Score: 382
- Comments: 318
- Posted: 2026-06-07T12:49:29Z

## Translation

タイトル: LLM が私のソフトウェア エンジニアリングのキャリアを侵食しているので、どうすればよいかわかりません
記事のタイトル: LLM が私のソフトウェア エンジニアリングのキャリアを侵食しており、何をすべきかわかりません |ループの中にいる人間
説明: 私はソフトウェア エンジニアで、今年で 10 年間の専門職経験を終えます。私は Web フロントエンド エンジニアとしてキャリアをスタートしました (Web フロントエンド エンジニアのほうが私にとっては楽でした)。

記事本文:
LLM が私のソフトウェア エンジニアリングのキャリアを侵食しているので、どうすればよいかわかりません
私はソフトウェア エンジニアで、今年で 10 年間の専門職経験を終えます。私は Web フロントエンド エンジニアとしてキャリアをスタートしましたが (当時はフロントエンド コードをデバッグする方が簡単だったので、その道を選びました)、すぐに (Web) バックエンドに移行し、振り返ることはありませんでした。
一連の偶然が重なって、バックエンド開発に足を踏み入れると、財務、簿記、支払処理の分野でソフトウェア開発の役割を担うことになり、そこでは大きな自主性があり、プロダクト マネージャーや関係者と緊密かつ率直な関係を築くことができました。
私はこのドメインと、そのためのプログラムを効果的に作成する方法について多くのことを学びました。PCI コンプライアンス、複式簿記、エスクロー、調整、支払いライフサイクル、銀行振込の冪等性などです。
したがって、ドメインスペシャリストの必要性が高まる兆しが見られるこの分野でプロフェッショナルとして頭角を現し、差別化を図るためには、そのドメインのエキスパートになることにキャリアを集中すべきであることは明らかでした。
侵食される最初の柱: ドメイン固有の知識
昨年、私は財務ワークスペースの会社に採用されました。これまで私は、事業やサービスに決済と財務の要素が強い企業に携わってきましたが、それは財務のみに焦点を当てた企業ではありませんでした。
その会社は AI も心から受け入れていたので、私は初日から ChatGPT と Claude Enterprise アカウントを取得し、研究、探索、さらにはコーディングにそれらを使用することを奨励されました。ただし、本番環境に導入されたすべての行をレビューして所有する必要があるという警告は付きました。
私の最初のプロジェクトの 1 つは、従来のオンライン決済システムの再構築に関係していましたが、それは混乱していました。彼らは（とりわけ）私のこれまでの経験を理由に私を雇いました。

それを構築し、私にその仕事を任せてくれました。
私がこれまで働いていた他の会社とは異なり、彼らはコーディングの前に私が作成する「設計ドキュメント」がエンジニアとプロダクトマネージャーの両方に読めるものであることを望んでいました。そのため、技術的な詳細を深く掘り下げるものではなく、アーキテクチャ的な観点を重視する必要があります。私は最小限の AI 支援で最初のものを書き、当時は LLM を「確率論的なオウム」とさえ呼んでいましたが、今ではそのような見方はしませんでした。そして、それを納品しました。
私は自分の知識を大切にしており、LLM がそれを置き換えることはできないと考えていました。
その後、マネージャーが私に連絡してきました。「あなたは良いペースでコードを提供しているにもかかわらず、設計ドキュメントを提供するのに時間がかかりすぎています。」 AIを使っていますか？もっとAIを活用すべきだ。
「こんなことうまくいくわけがない」と私は頭の中で思いましたが、同意しました。当時のモデルは現在のモデルほど優れたものではありませんでしたが、執筆や意思決定を大幅にスピードアップしてくれました。
そして私は、長年にわたって蓄積してきたすべての知識、つまり実装間のトレードオフ、取得の仕組み、二重請求を防ぐ冪等性の構築方法など、すべてが役に立たなくなってきたことに気づき始めました。モデルにはまだある程度の舵取りが必要でしたが、そのようなシステムをどのように構築するかについて点と点を結びつけることができました。これは、長年の実地経験を経て頭の中でのみ開発される最も難しい部分でした。それが私の最初の衝撃でした。
しかし、確かに、ウェブ上にはすべての技術文書とともにそのクソがどのように機能するかについての記事がたくさんあり、技術ツールをドメインに適用する方法を説明するブログ投稿があるため、彼らはそれを行うことができると思いました。人間にとって、これらすべてを学習するには長い時間がかかるかもしれませんが、それはモデルが認識できるようにするためのトレーニング データです。
モデルが決して得意ではないこと、そして人間が輝けるところは、デバッグです。溜まってた

実稼働環境での競合状態と分散システムのデバッグに良い経験を積んできました。それが私にとって長期雇用への切符でした。
侵食される 2 番目の柱: デバッグと分散システム
そのため、LLM はドキュメントを書くのが上手になり、実際の実装の計画を支援し始めてから、コーディングも得意になりました。それは 2025 年後半にクロード コードの誇大広告で始まり、その後コーデックスが登場して、というように続きました。それまでは毎日単体テストの作成に LLM を使用していましたが、まだ LLM が完全な実装を作成できるとは信じていませんでした。
自然な次のステップは、コードの作成にさらに AI を導入することでした。そして正直に言うと、気に入りました。私はコーディングと同じくらい、本番環境に出荷してユーザーが喜ぶのを見るのが好きなので、自分の好きなものを、同じく好きなものと交換することは公平でした。
LLM はコーディングが上手になりつつありましたが、(その時までに、あるいは人間によって) 残された混乱をデバッグすることはまだできませんでした。そのため、私には依然としてロボットを操縦するよりも大きな役割、つまり雇用適性への切符が残されていました。
その後、MCP、エージェント ワークフロー、Claude 4.5 が登場すると、空は落ち始めました。
正直に言うと、クロード4.5はそれほど良くありませんでした。スタック トレースといくつかのコンテキストを考慮すると、バグの 60% ほどが解決されました (ほとんどの場合、Sentry MCP を有効にした Sentry リンクだけで十分でした)。場合によっては、もっともらしく聞こえても完全に間違っている解決策が得られることもありました。
しかし、今回は機械を疑うのをやめました。以前はフルタイムのデバッグに軽く 1 日かかっていたバグが、Claude Code によってワンショットで解決されるのを目にしました。もちろん、まだすべてではありませんが、パターンは明らかでした。
その後、4.6、4.7、GPT 5.5、Opus 4.8、そして DataDog MCP が登場しました...今では、分散システム全体のバグをワンショットで解決する CLI が手に入りました。過去に解決できなかったバグ。フルタイムのデバッグに 2 日かかるバグ

NG。分散された可観測性を欠いた分散システム全体にわたるバグ。現在、バグの 90% は単発的であり、これには奇妙な競合状態、予期しない例外的なケース、サードパーティの統合の問題、文書化されていない API のエッジ ケースなどが含まれます。介入する必要はほとんどありません。
もちろん、誰かがコードをレビューしてロボットを操縦する必要があるため、私はまだ雇用可能です。しかし、私は今ではただの平凡なエンジニアです。私には、LLM を運営する他のシニア エンジニアに匹敵できないようなドメインの専門知識がありません。私の金融と決済の分野の専門知識、デバッグの直感、そして何時間も汗と涙を流して得た分散システムの知識はすべて、今ではすぐに活用できるようになりました。
私たちは、ゼネラリストとスペシャリストには常にそれぞれの役割があると教えられました。しかし今、市場は誰もがゼネラリストになるよう仕向けています。それ自体は悪いことではありませんが、需要と供給の経済学に目を向けるまでは、誰もがジェネラリストである場合、それに見合った需要がなければジェネラリストの価格は下がります。そして需要が枯渇しつつあることは誰もが知っています。
3 番目の柱、まだ侵食されていないもの: コードの品質とアーキテクチャ
しかし、私にはまだ 1 つの柱が立っています。コードの品質とソフトウェア アーキテクチャです。現在は「テイスト」と呼ばれるようになっています 1。
これまでのキャリアの中で、私は常にリファクタリングが好きで、常に良いコードを高く評価し、そのためのスプリントの時間を交渉してきました。 DDD、ヘキサゴナル、クリーン アーキテクチャなど、あらゆる流行語をご存知でしょう。私はこのトピックが好きで、コードベースを形成する方法に関するトレードオフやさまざまなアイデアについて議論するのが好きです。本当に気に入っています。
これが最後に立っている柱です。もう誰も気にしていないことを除いて。
エージェントは、コードベースを整理しておくのが非常に下手です。あなたがそれらを制御しなければ、あなたが思っているよりも早く循環依存の問題に遭遇するでしょう。コードが重複します。不要なものを追加する

アリのコメント。純粋な関数と副作用を混同してください。 SOLID の原則を無視してください。
このスキルが現在「味覚」という言葉に還元されつつあることを除けば、人間の雇用は維持されるはずだ。しかし、これは単なる名前変更ではなく、業界はコード構成がそれほど重要ではない世界に移行しつつあります。
確かに、循環依存関係グラフを含むスパゲッティ コードベースを防ぐために人間がエージェントを操作する必要があります。私たちは、何かを壊さずに触ることが不可能な F 指定のコードベースを望んでいません。でもCかD？もう大丈夫です。 A グレードや B グレードのコードベースはもう誰も必要としません。コードベースは人間が読むためではなく、LLM のために作られているからです。
これが本質的に良いのか悪いのかを議論するつもりはありません。ソースコードが人間ではなく機械が読み取れるように書かれている場合、実際には人間をターゲットにしても問題ないかもしれません。
しかし、それは私の専門知識のもう一つの柱であり、失われつつあります。そのトピックに関して私が蓄積した知識の大部分は、もはやそれほど価値がありません。本を読んだり、実際の演習をしたり、他のエンジニアと議論したり、ADR を書いたりすることに費やした時間はすべて無駄になりつつあります。
私はまだ雇用されており、近い将来（少なくともその会社では）雇用されると思います。しかし、長期的なことを考えるとどうなるかわかりません。
私は、ますます価値が低くなりつつある物事を得意にするために、10 年（専門職以外の経験を考慮するとさらに長い年月）を費やしました。私の最後の専門知識の柱は今や「味」に成り下がっており、おそらく長くは続かないでしょう。
そして、それが私だけではないこともわかっています。約 8 か月前、私の現在の会社で解雇がありました (AI とは関係ないとのことです)。優秀な元同僚の中には解雇され、今も職を探している人もいます。彼らのほとんどは、私がここで概説したのと同じ問題に苦しんでいます。つまり、彼らの専門分野の専門知識はもはや目立つのに十分ではありません。
会社は今こんにちは

いくつかの役割のために再び電話をかけると、ドメインの精通度はもはや強力な差別化要因ではなくなります。以前は「ソフトウェアエンジニア - 分野」と記載していました。現在は単なる「ソフトウェア エンジニア」であり、チームへの割り当てはオファーが受け入れられた後に行われます。
もちろん、これは、これまでドメインに深く関わる機会がなかった優秀なエンジニアにとっては良いことであり、今では仕事を得るチャンスが増えていますが、人生をかけてドメインの知識を収集してきた他の優秀なエンジニアが同じレーンで競争していると考えると悲しいことでもあります。
私の雇用適性を長期的に維持するための唯一の方法は、今のところ、私の専門分野の専門知識を、LLM が簡単に得意としない分野にシフトすることのようです。しかし、何が残っているでしょうか？
私は大学に戻り、数学、統計、高度な機械学習を学び、フロンティアラボでの研究職に応募することを考えました。私の国にはフロンティアラボが存在しないことを除けば、数少ないフロンティアラボには応募が殺到しており、私には家族の事情があり、他の国への移住が困難になっています。私がそこまで飛躍する余裕ができる頃には、RSI のせいで研究者は時代遅れになっているかもしれない。
木工の趣味を職業に変えることを考えるべきかもしれません...
参考として、 this 、 this 、および this を参照してください。これを、これらの投稿の内容を支持するものと受け取らないでください。 ↩
#あい
#llm
#ソフトウェアエンジニアリング

## Original Extract

I'm a software engineer, completing 10 years of professional experience this year. I started my career as a web frontend engineer (it was easier for me to de...

LLMs are eroding my software engineering career and I don't know what to do
I'm a software engineer, completing 10 years of professional experience this year. I started my career as a web frontend engineer (it was easier for me to debug frontend code back then, so I chose that path), but shortly transitioned to (web) backend and never looked back.
Through a series of coincidences, once I stepped into backend development, I ended up working in software development roles in the domains of finance, bookkeeping and payment processing, where I had great autonomy and a close and candid relationship with Product Managers and stakeholders.
I learnt a lot about the domain and how to effectively write programs for it: PCI compliance, double-entry ledgers, escrows, reconciliation, payment lifecycles, bank transfer idempotency, etc.
It was, then, obvious that I should focus my career on becoming an expert on that domain to stand out as a professional and differentiate myself in a field that showed signs of an increasing need for domain specialists.
The first pillar to erode: domain-specific knowledge
Last year, I got hired by a company in the finance workspace. So far, I had worked on companies that do have a strong payment and finance component to their operations/offerings, but that were not solely finance-focused companies.
That company also embraced AI wholeheartedly, so I got ChatGPT and Claude Enterprise accounts from day one and was encouraged to use them for my research, exploration, and even coding, albeit with a warning that I should still review and own every single line that made it into production.
One of my first projects involved reworking the legacy online payment system, which was a mess. They hired me for (among other things) my previous experience in building that and trusted me with the task.
Different from the other companies I had worked for so far, they wanted the "Design Docs" I write before coding to be readable by both engineers and product managers - so they shouldn't be a technical deep dive and more of an architectural view. I wrote my first one with minimal AI assistance - I even called LLMs "stochastic parrots" at the time, a view I no longer hold - and delivered it.
I valued my knowledge and thought no LLMs could replace it.
Then my manager reached out to me: even though you're delivering code at a good pace, you're taking too long to deliver those Design Docs. Are you using AI? You should use more AI.
"No way this will work", I thought in my head, but agreed. The models at that time were not as good as the ones we have now, but they did provide a good speed-up on my writing and even the decision-making.
And then I started realizing: all the knowledge I have accumulated over the years: the trade-offs between implementations, how acquiring works, how to structure idempotency to prevent double-charges, everything, was becoming useless. Even though the models still needed some steering, they could connect the dots on how to structure such systems, which was the hardest part that only develops in your brain after years of hands-on experience. That was my first shock .
But sure, I thought, they can do that because there's plenty of articles on the web on how that shit works along with all the technical documentation, and we have blog posts explaining how to apply the technical tools to the domain. For humans, it may take a long time to learn all that, but that's training data so the models can pick it up.
What the models will never be good at, and that's where humans will shine, is debugging! I had accumulated a good experience debugging race conditions and distributed systems in production. That was my ticket to long-term employability.
The second pillar to erode: debugging and distributed systems
So, after LLMs started getting good at writing docs and helping plan the actual implementations, they became good at coding. It started in the second half of 2025 with the Claude Code hype, then Codex came and so on. Although I was using LLMs for writing unit tests every day before that, I wasn't trusting them to write the full implementation yet.
The natural next step was to introduce more AI into writing code. And honestly, I liked it. I like shipping things to production and seeing users happy as much as I like coding, so I was trading one thing that I like for another one that I also like, it was fair.
LLMs were becoming good at coding, but it still couldn't debug the mess left behind (by then or by the humans), so I still had a role that was bigger than steering the robot - a ticket to employability.
Then came the MCPs, the agentic workflows and Claude 4.5 and the sky started to fall.
Claude 4.5, to be honest, wasn't that good. It solved like 60% of the bugs given a stack trace and some context (a Sentry link with Sentry MCP enabled was all it took in most cases). Sometimes it gave a solution that sounded plausible but was totally wrong.
This time, however, I stopped doubting the machines. I saw bugs that in the past would easily take 1 day of full-time debugging being one-shotted by Claude Code. Of course, not all of them yet , but the pattern was clear.
Then came 4.6, 4.7, GPT 5.5, Opus 4.8 and the DataDog MCP... Now I have CLIs that one-shots bugs across distributed systems for me. Bugs that I couldn't solve in the past. Bugs that would take 2 days of full-time debugging. Bugs across distributed systems that lack distributed observability. 90% of the bugs are one-shotted now, including bizarre race conditions, unexpected corner-cases, third-party integration issues, undocumented API edge cases, everything. I hardly have to intervene.
Of course, I'm still employable because someone has to review the code and steer the robot. But I'm just another off-the-shelf engineer now. I have no domain expertise that another Sr. engineer steering an LLM cannot match. All my finance and payment domain expertise, all the debugging intuition and distributed system knowledge earned through hours of sweat and tears, is now promptable .
We were taught that generalists and specialists will always have their roles. But now the market is shaping everyone into becoming a generalist. That's not a bad thing per se , until you look under the economics of supply and demand: if everyone is a generalist, the price of a generalist falls if there's no demand to match. And we all know the demand is drying up.
The third pillar, the one that hasn't eroded yet: code quality and architecture
I still have one pillar standing, though: code quality and software architecture - what's now being reduced to being called "taste" 1 .
Along the course of my career, I always liked to refactor, always prized good code, and negotiated time in the sprint for it. DDD, Hexagonal, Clean Architecture, you know all the buzzwords. I like this topic, I like to discuss the trade-offs and different ideas on how to shape codebases. I really like it.
This is the last pillar standing. Except that nobody cares anymore.
Agents do a really bad job at keeping codebases organized. If you don't steer them, they'll hit a circular dependency issue sooner than you think. Will duplicate code. Add unnecessary comments. Mix up pure functions and side-effects. Disregard the principles of SOLID.
That should keep humans employed, except that this skill is now being reduced to the word "taste". But it's not just a renaming, the industry is moving to a world where code organization is less important.
Sure, humans should steer the agent to prevent spaghetti codebases with circular dependency graphs. We don't want F -rated codebases that are impossible to touch without breaking something. But a C or D ? It's now fine. Nobody needs A or B -grade codebases anymore because they're being made for LLMs, not for humans to read.
I don't want to argue if this is inherently good or bad. If the source code is now written for machines to read and not humans, it may be actually ok to target them.
But that's another pillar of my expertise that's eroding. A good chunk of the knowledge I accumulated on that topic is not that valuable anymore. All the time I spent on it - reading books, doing real-world exercises, discussing with other engineers, writing ADRs - is becoming useless.
I'm still employed and I see myself employed (at least in that company) for a foreseeable future. But I don't know what to think about the long-term.
I spent 10 years (even more when you account for non-profession experience) getting good at things that are becoming less and less valuable. My last pillar of expertise is now reduced to a "taste" and will probably won't last long.
And I know that's not just me. About 8 months ago there was a layoff at my current company (not related to AI, according to them). Some brilliant ex-coworkers were laid off and are still looking for jobs. Most of them suffer from the same problem I outlined here: their domain expertise is not enough to stand out anymore.
The company is now hiring again for a few roles and domain familiarity is not a strong differentiator anymore. We used to list "Software Engineer - Area". Now it's just "Software Engineer" and the team assignment comes after the offer is accepted.
Of course, this is good for brilliant engineers that never had the chance to get deep into the domain and now have better chances at getting a job, but it's also sad to think that other brilliant engineers that spent their lives collecting domain knowledge are now competing on the same lane.
The only way out for keeping my employability in the long-term now seems to be shifting my domain expertise to something LLMs will not get good at so easily. But what's left?
I thought about going back to college, learning Math, Statistics, advanced Machine Learning and applying for research role at a frontier lab. Except that there are no frontier labs in my country, the few ones that exist are flooding with applications and I have family matters that makes moving to another country difficult. By the time I can afford to make that jump, RSI may have made researchers obsolete.
Maybe I should consider transforming my woodworking hobby into a profession...
See this , this and this for reference. Don't take this as an endorsement of the content inside any of these posts. ↩
#ai
#llm
#software engineering

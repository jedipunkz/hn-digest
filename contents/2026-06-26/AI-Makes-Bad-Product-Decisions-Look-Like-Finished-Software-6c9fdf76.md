---
source: "https://www.vincentschmalbach.com/ai-makes-bad-product-decisions-look-like-finished-software-cost-quality-test/"
hn_url: "https://news.ycombinator.com/item?id=48686237"
title: "AI Makes Bad Product Decisions Look Like Finished Software"
article_title: "AI Makes Bad Product Decisions Look Like Finished Software - Vincent Schmalbach"
author: "vincent_s"
captured_at: "2026-06-26T13:39:59Z"
capture_tool: "hn-digest"
hn_id: 48686237
score: 1
comments: 1
posted_at: "2026-06-26T13:13:11Z"
tags:
  - hacker-news
  - translated
---

# AI Makes Bad Product Decisions Look Like Finished Software

- HN: [48686237](https://news.ycombinator.com/item?id=48686237)
- Source: [www.vincentschmalbach.com](https://www.vincentschmalbach.com/ai-makes-bad-product-decisions-look-like-finished-software-cost-quality-test/)
- Score: 1
- Comments: 1
- Posted: 2026-06-26T13:13:11Z

## Translation

タイトル: AI により、製品に関する不適切な意思決定が完成したソフトウェアのように見える
記事タイトル: AI により、不適切な製品の意思決定が完成したソフトウェアのように見える - Vincent Schmalbach
説明: AI ツールにより、1 つのソフトウェア障害モードがはるかに見落とされやすくなります。誤った製品の決定が、動作するインターフェイスに包まれて届く可能性があります。数分以内にあなたは…

記事本文:
コンテンツにスキップ
ソフトウェア開発者
ホーム
サービス
カスタム Laravel アプリ開発
Laravelのパフォーマンスの最適化
技術チームのリーダーシップとメンタリング
Laravel開発コンサルティング
AI により、誤った製品の決定が完成したソフトウェアのように見える
2026 年 6 月 26 日
ヴィンセント・シュマルバック著
AI ツールにより、1 つのソフトウェア障害モードがはるかに見落とされやすくなります。製品の誤った決定が、動作するインターフェイスに包まれた状態で届く可能性があります。
私は、タスクをプル リクエストに変換するための AI コーディング ツールである vroni.com で作業しながらこれを書いています。
ほんの数分で、ログイン ページ、ダッシュボード、設定画面、さらにはチェックアウト フローを含むものを吐き出すことができます。すべてが一貫しているように見えるため、人々はそれを重要な決定がすでに行われたかのように扱い始めます。彼らはそうではありません。ほとんどの場合、製品の判断が行われる必要がある真空で滑らかなシェルが得られます。システムは完全に見えても、データ分離、エラー回復、アクセス制御、および基本的な運用境界がまだ不足している場合があります。
プロトタイプを超えたものにとって、「バイブ コーディング」が危険になるのはここです。現世代の生成ツールは、チームが所有権、プライバシー ルール、障害動作、サポート境界を定義するよりも速く、ソフトウェアに見えるものを作成できます。 Val Town のエッセイでは、これについて述べています。ビルダーはコードの存在を忘れる可能性があります。 1 回限りのプロジェクトの場合はこれで十分です。現実を生き、変化し、生き残らなければならないものにとって、それは悪い取引です。
モデルが記述してもビルダーが理解できないコードは、初日からはレガシー コードです。
プログラミングは単なる構文ではありません。それは理論構築です。部分がどのように組み合わされるか、どのような仮定が行われるか、境界はどこにあるのかを示す内部モデルが必要です。生成モデルは、ユーザーに代わってコードを作成できます。 u を生成せずに信頼できる実装を生成できます

理解しています。したがって、システムは概念的には空であっても、UI はいっぱいであるように見えることがあります。
ツールはデフォルトの動作を喜んで選択します。政策のトレードオフを喜んで無視するだろう。それは、うまくデモするものを喜んで吐き出しますが、実際のシステムのように動作するように要求するとすぐに崩壊します。
ソフトウェア チームは、この問題の弱いバージョンをすでに持っていました。忠実度の高いモックアップは、技術者ではない人々をだますか、少なくともエンジニアが介入して「いいえ、これは終わっていません」と言わなければならないほどの混乱を引き起こす可能性があります。 Hacker News のスレッド では、開発者らが同僚に対し、洗練されたデザインをマネージャーに見せるのはやめてほしいと伝えていると話している。
現在、プロトタイプはライブ データベースに接続され、パブリック ドメインに展開されることがよくあります。表面的な動作は、もはや「これは本物に見える」というだけではありません。 「これは危険なほど現実的だ」と書かれています。
そのため、チームにはエンジニアリングレビューを省略するよう大きなプレッシャーがかかります。アプリが短いデモで動作する場合、自然な反応はそれを出荷することです。ただし、プロトタイプはコンセプトを証明するためのものであり、公共のインターネット上の敵対的なトラフィックを処理するためのものではありません。中途半端に理解したプロトタイプを製品化すると、不足している製品の決定はもはや理論的なものではなくなります。それらは運用上の失敗となります。
「My Account」というページがあるので安心できますが、バックエンドの承認ルールがない限り、ラベルは単なる装飾テキストです。このインターフェースは 1 つのことを約束します。システムは別のものを提供します。
これらのセキュリティへの影響は理論上のものではありません。 Axios の報告によると、監査の結果、展開プラットフォーム全体で 380,000 件の公開資産が存在し、そのうち 5,000 件には財務文書や医療文書などの機密企業データが含まれていることが判明しました。 CVE-2025-48757 では、アプリケーションにおけるデータベースの行レベルのセキュリティ ポリシーが不十分であることに関する同様の問題について説明しています。

Lovableによって生成されたns。脆弱性スキャンの結果、この欠陥が 170 以上のアプリに影響し、認証されていない攻撃者がデータベース テーブルの読み書きを可能にすることが判明しました。
Superblocks は、行レベルのセキュリティがバックエンドで構成されておらず、クライアント側の資格情報がデータベースに直接クエリを行っていたと説明しています。これは、核となる政策決定が存在する前にシステムが「完成」したときに起こることです。
AI はインターフェースを構築できますが、誰がどのレコードを所有するかは決定しません。ある顧客の別の顧客のデータへのアクセスを制限するものではありません。何をブロックするか、何をログに記録するか、または何をサーバーから流出させてはならないかは決まりません。これらは製品に関する決定であり、コード補完タスクではありません。人間が強制的に存在させない限り、それらは存在しません。
『WIRED』は、そのリスクをオープンソースの依存関係の問題と比較してこれを論じているが、同様の明確な責任は示されていない。故障モードはよく知られているので、これは良い例えです。何かが再利用可能で構成可能であるように感じられますが、危険は隠れた想定の中にあります。違いは、多くの場合、手書きのコードでは、生成されたコードよりもこれらの前提がより明確になるということです。正式なレビューがなければ、誰かが苦労して発見するまで漏れは発見されない可能性があります。
これはセキュリティ上の問題だけではなく、配送上の問題でもあります。 DORA 生成 AI レポートでは、AI の導入により個人の生産性は向上しますが、納品の結果に悪影響を及ぼす可能性があると述べています。具体的には、AI 導入が 25% 増加すると、ソフトウェア配信のスループットが 1.5% 低下し、配信の安定性が 7.2% 低下することがわかりました。それは驚くべきことではありません。コード生成が高速化しても、ソフトウェア配信がより適切になるわけではありません。組織が考える機会を得る前に、コードが増えるだけになる可能性があります。
DORA は組織が非常にユーザー中心であるとも述べています

パフォーマンスが 40% 向上します。それがなければ、AI はより高速な機能ファクトリーを作成するだけであるため、これは重要です。生産量が増えても、必ずしも価値が高まるわけではありません。機能が増えれば故障する表面積も増えますが、ユーザーが実際に何を必要としているのかが明確になっていません。
有益な対応は、生成された作品に対する規律です。 DORA のプラットフォーム エンジニアリング ガイダンスは、セキュリティ チェック、テスト プロトコル、展開ガードレールを自動的に実施する開発者プラットフォームを指しています。これは、生成された大量のコードを扱う正しい方法です。 Cyber​​Scoop も同様の基本的な点を主張しています。AI 支援による開発は避けられないものの、そのスピードにはより厳密な検証が必要です。生成されたインターフェイスは、製品に関する会話を終了するのではなく、開始する必要があります。
Vroni に GitHub の問題、バグレポート、仕様、または大まかなアイデアを提供します。リポジトリを読み取り、変更を計画し、コードを作成し、チェックを実行し、レビュー可能なプル リクエストに向けて作業します。
新しい投稿を公開すると取得します。
私はあなたのプライバシーを尊重します。いつでも購読を解除してください。
このトピックに関してさらに執筆します。
AI 生産性の罠はより多くの生産物を生み出す
AI により、コードの生成、チケットの草案、会議の要約、提案書の作成が非常に安価になります。しかし、より安価な世代が作業を意味するわけではありません…
ソフトウェアが 90% 安くなった場合、その節約分は誰が回収するのでしょうか?
AI がはるかに高い速度でコーディングできると想像してください。開発費が大幅に安くなると考えるのは簡単です。しかし…
AI コーディング時代により退屈なテストの価値が高まる
エージェントは安価で妥当な差分を作成できるようになりました。開発者はバグを説明し、エージェントに渡してパッチを入手できます。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
ソフトウェア、AI、Laravel、SEO、オンライン製品の構築に必要なものについて書いています。 2025 年には、73,000 人がこのブログを読んでいます。一部の投稿は出版物や開発者コミュニティによって取り上げられました

s.
いくつかの古い投稿が最も注目を集めました。以下の最新の文章とは別に、ここに保管しておきます。
Google Now はデフォルトでコンテンツのインデックスを作成しません
この投稿は、Guardian のコラムと開発者コミュニティでの長い議論につながりました。
AI がテクノロジー負債を指数関数化する
AI によって脆弱なエンジニアリング基盤のコストが下がるのではなく、より高価になる理由。
専門家だけが適切なプロンプトを作成できる
AI が適切に機能するかどうかは、依然として技術的な判断とセンスにかかっています。
実際のコードベースで作業する人向けの Laravel に関する具体的なアドバイス。
ブログの最新の投稿。これは公開すると自動的に更新されます。
AI により、誤った製品の決定が完成したソフトウェアのように見える
AI ツールにより、1 つのソフトウェア障害モードがはるかに見落とされやすくなります。製品の誤った決定が到着する可能性があります...
AI 生産性の罠はより多くの生産物を生み出す
AI により、コードの生成、チケットの草案、会議の要約、提案書の作成が非常に安価になります。しかし、より安価な世代...
ソフトウェアが 90% 安くなった場合、その節約分は誰が回収するのでしょうか?
AI がはるかに高い速度でコーディングできると想像してください。開発がうまくいくだろうと考えるのは簡単です...
私の文章や解説について言及、引用、または議論した場所。
私がこれまでに仕事をしたクライアントからの実際のフィードバック。
「ヴィンセントを雇ってください。真剣に。私は何年にもわたって多くの開発者と仕事をしてきましたが（レンタコーダーを覚えていますか？）、ヴィンセントはそれを理解する稀な種類の Web 開発者です。チャットをし、必要なことを表現すると、彼は求めたものを返してきます。彼の英語は流暢で、あなたが何を必要としているのかを理解しており、物事を過度に複雑にするつもりはありません。彼と一緒に仕事をするのは喜びであり、細部まで管理する必要がなく、喜びです。全体的に素晴らしい経験だった。正直に言って、もう予算がないので彼を続けてほしかったが、それが契約が終了した唯一の理由だ」
「ヴィンセントとの仕事で印象に残ったことは

彼がいかに几帳面で効率的だったかだ。彼は私たちの要件を徹底的に掘り下げることから始め、しっかりと理解するまで明確な質問をしました。その後、彼は必要なものを正確に構築し、常に私たちに情報を提供してくれました。コミュニケーションは常に明確で積極的でした。時間と予算の両方において、すべてが順調に進みました。最終結果は素晴らしく、発売以来順調に稼働しています。実践的な焦点と細部へのこだわりの組み合わせが違いを生みます。質の高い開発を行うには、間違いなく Vincent をお勧めします。」
「ヴィンセントをリテイナーに迎えたことは素晴らしいことだ。彼は私たちのためにいくつかの大きな技術的課題に取り組み、特にウェブサイトのパフォーマンスを大幅に向上させました。彼はまた、当社のコア プラットフォーム (Laravel) の最新バージョンへの複雑なアップグレードを専門的に管理し、当社が常に最新かつ安全であることを保証しました。私たちは現在、彼が構築した多言語機能を展開していますが、これは非常に興味深いことです。 Vincent は信頼でき、技術的に熟練しており、サイトの最適な運営と前進を維持するための貴重なパートナーです。彼の専門知識を強くお勧めします。」
「私は Vincent と数か月間仕事をしてきましたが、彼は私が共同作業した中で最も独立した開発者の 1 人です。彼には細かい管理は必要ありません。彼は意思決定を下し、プロジェクトを推進する能力を十分に備えているため、私はビジネスの他の部分に集中できるようになりました。
Vincent は非同期で非常にうまく機能します。コミュニケーションは集中したままで、不必要なやり取りはほとんどなく、必要に応じて重要な質問を提起します。その結果、開発者はオーバーヘッドなしで開発を行うことができ、これは本当に価値のあることです。」
「私たちは新しいビルドと長年にわたって成長してきた既存のコードベースの両方で、Laravel/Vue プロジェクトで Vincent と定期的に協力しています。どちらの場合も

、彼はクリーンなコードと信頼性の高い実装を提供します。私たちは、彼が成熟したプロジェクトにどれだけ早く着手し、コードベースの既存の状態に適合するソリューションを見つけるかを特に評価しています。絶対にお勧めです。」
「Vincent は扱いやすく、高品質のコードを作成してくれました。お勧めします。」
「Vincent は、複雑なコードベースに飛び込み、すぐに役に立ち、手を握ることなく実際の問題を解決できる稀有な開発者の 1 人です。
彼は、パフォーマンスの最適化、顧客固有の機能、および主要なカスタム レポート システムに取り組みました。また、顧客との議論を明確な技術要件に変換し、チームの他のメンバーがそれに基づいて構築できるように自分の作業を文書化しました。
Vincent はチケットをクローズするだけでなく、製品、エッジケース、顧客への影響を徹底的に考慮します。 SaaS 製品を前進させることができる上級 Laravel 開発者が必要な場合は、彼を強くお勧めします。」
何か見るべきものはありますか？
ラフ版をお送りします。私はそれを計画に落とし込み、それを構築するのを手伝うことができます。
私の本: Laravel による Rapid SaaS
SaaS を数か月ではなく数日で立ち上げます。 Laravel 12 を使用して本番環境に対応した SaaS アプリを構築するための実用的な、BS なしのガイド: 速度、請求、チーム、AI 統合、およびカスタム前に必要な要素

[切り捨てられた]

## Original Extract

AI tools make one software failure mode much easier to miss: a bad product decision can now arrive wrapped in a working interface. In a matter of minutes you…

Skip to content
Software Developer
Home
Services
Custom Laravel App Development
Laravel Performance Optimization
Technical Team Leadership & Mentoring
Laravel Development Consulting
AI Makes Bad Product Decisions Look Like Finished Software
June 26, 2026
by Vincent Schmalbach
AI tools make one software failure mode much easier to miss: a bad product decision can now arrive wrapped in a working interface.
I'm writing this while working on vroni.com , my AI coding tool for turning tasks into pull requests.
In a matter of minutes you can spit out something with a login page, a dashboard, a settings screen, maybe even a checkout flow. Because it all looks coherent, people start treating it as if the important decisions have already been made. They have not. Mostly you get a slick shell with a vacuum where product judgment should be. The system may look complete while still lacking data isolation, error recovery, access control, and basic operational boundaries.
This is where "vibe coding" becomes dangerous for anything beyond prototypes. The current generation of generative tools can produce software-looking things faster than teams can define ownership, privacy rules, failure behavior, and support boundaries. The Val Town essay gets at this: the builder may forget that the code exists. That can be fine for one-off projects. It is a bad trade for anything that has to live, change, and survive reality.
Code that a model writes and the builder cannot understand is legacy code on day one.
Programming is more than syntax. It is theory building. You need an internal model of how the pieces fit together, what assumptions they make, and where the boundaries are. A generative model can write code for you. It can generate believable implementations without generating understanding. So the UI can appear full, while the system is conceptually empty.
The tool will gladly choose default behavior. It will gladly disregard policy trade-offs. It will happily spit out something that demos well, then falls apart the second you ask it to behave like a real system.
Software teams already had a weaker version of this problem. A high-fidelity mockup could fool non-technical people, or at least cause enough confusion that engineers had to step in and say no, this is not done. In a Hacker News thread , developers talk about telling coworkers not to show polished designs to managers because it makes everyone think the feature is finished.
Today the prototype is often connected to a live database and deployed on a public domain. The surface behavior no longer says only "this looks real." It says "this is real enough to be dangerous."
That puts a lot of pressure on teams to skip engineering review. If the app works in a short demo, the natural reaction is to ship it. But a prototype is for proving a concept, not for handling hostile traffic on the public internet. Once you put a half-understood prototype into production, the missing product decisions are no longer theoretical. They become operational failures.
There is a page called "My Account," which sounds reassuring, but unless there are backend authorization rules, the label is just decorative text. The interface promises one thing. The system delivers another.
These security implications are not theoretical. Axios reported that an audit found 380,000 publicly exposed assets across deployment platforms, with 5,000 containing sensitive corporate data like financial and medical documents. CVE-2025-48757 describes a similar issue around insufficient database row-level security policies in applications generated by Lovable. A vulnerability scan found that the flaw affected more than 170 apps and allowed unauthenticated actors to read and write database tables.
Superblocks explains that row-level security was never configured on the backend and that client-side credentials were directly querying the database. That is what happens when the system is "finished" before the core policy decisions exist.
The AI can build the interface, but it does not decide who owns which record. It does not restrict one customer's access to another customer's data. It does not decide what to block, what to log, or what must never leave the server. These are product decisions, not code-completion tasks. They do not exist unless human beings force them into existence.
WIRED discusses this by comparing the risk to open-source dependency problems, but without the same clear line of accountability. It is a good analogy because the failure mode is familiar: something feels reusable and composable, but the danger sits in the hidden assumptions. The difference is that hand-written code often makes those assumptions more explicit than generated code does. Without a formal review, the omissions may not be detected until someone finds out the hard way.
This is also a delivery problem, not just a security problem. The DORA generative AI report says AI adoption improves individual productivity but can hurt delivery outcomes. Specifically, it finds that a 25% increase in AI adoption is associated with a 1.5% reduction in software delivery throughput and a 7.2% reduction in delivery stability. That should not be surprising. Faster code generation does not mean better software delivery. It can just mean more code before the organization gets a chance to think.
DORA also says highly user-centric organizations perform 40% better. That matters because without it, AI just creates a faster feature factory. More output, not necessarily more value. More features, more surface area to fail, and no clearer idea of what the user actually needs.
The useful response is discipline around the generated work. DORA's platform engineering guidance points to developer platforms that automatically enforce security checks, testing protocols, and deployment guardrails. That is the right way to work with large amounts of generated code. CyberScoop makes the same basic point: AI-assisted development is inevitable, but its speed demands more rigorous verification. A generated interface should start the product conversation, not end it.
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.
Get new posts when I publish them.
I respect your privacy. Unsubscribe at any time.
More writing around this topic.
The AI Productivity Trap Is More Output
AI makes it dirt cheap to generate code, draft tickets, summarize meetings, and write proposals. But cheaper generation doesn't mean the work…
If Software Gets 90 Percent Cheaper, Who Captures the Savings?
Imagine AI can code at a much higher rate. It’s easy to think development is going to get a lot cheaper. But…
The AI Coding Era Makes Boring Tests More Valuable
Agents can now produce cheap, plausible diffs. A developer can describe a bug, give it to an agent, and get a patch…
Your email address will not be published. Required fields are marked *
I write about software, AI, Laravel, SEO, and what it takes to build online products. In 2025, 73k people read the blog. Some posts were picked up by publications and developer communities.
A few older posts got the most attention. I keep them here, separate from the newest writing below.
Google Now Defaults to Not Indexing Your Content
The post that led to a Guardian column and a long developer-community discussion.
AI Exponentializes Your Tech Debt
Why AI makes weak engineering foundations more expensive, not less.
Only Experts Can Write Good Prompts
Good AI work still depends on technical judgment and taste.
Specific Laravel advice for people working in real codebases.
The newest posts from the blog. This updates automatically when I publish.
AI Makes Bad Product Decisions Look Like Finished Software
AI tools make one software failure mode much easier to miss: a bad product decision can now arrive...
The AI Productivity Trap Is More Output
AI makes it dirt cheap to generate code, draft tickets, summarize meetings, and write proposals. But cheaper generation...
If Software Gets 90 Percent Cheaper, Who Captures the Savings?
Imagine AI can code at a much higher rate. It’s easy to think development is going to get...
Places that have mentioned, quoted, or discussed my writing and commentary.
Actual feedback from clients I've worked with.
"HIRE VINCENT. Seriously. I've worked with many many developers over the years (remember rentacoder?), and Vincent is that rare breed of web developer who just gets it. You have a chat, you express what you need, and he comes back with what you asked for. His english is fluent, he understands what you need, and he's not looking to overcomplicate things. It's a joy and pleasure to work with him, not to have to micromanage and just overall a great experience. Honestly I wish I could keep him on we just don't have the budget anymore - that is the only reason the contract ended."
"What stood out about working with Vincent was how methodical and efficient he was. He started by really digging into our requirements, asking clarifying questions until he had a solid understanding. He then built exactly what was needed and kept us informed throughout. Communication was clear and proactive at all times. Everything stayed on track, both in terms of time and budget. The end result was excellent and has been running smoothly since launch. It's that combination of practical focus and attention to detail that makes the difference. I would definitely recommend Vincent for quality development."
"Having Vincent on retainer has been excellent. He's tackled some major technical challenges for us, notably significantly improving our website's performance. He also expertly managed a complex upgrade of our core platform (Laravel) to the latest version, ensuring we stay current and secure. We're now rolling out the multilingual capabilities he built, which is exciting. Vincent is reliable, technically skilled, and a valuable partner for keeping our site running optimally and moving forward. Highly recommend his expertise."
"I've been working with Vincent for several months, and he's one of the most independent developers I've collaborated with. He doesn't need micromanagement - he's fully capable of making decisions and driving the project forward, which has freed me up to focus on other parts of the business.
Vincent works very well asynchronously: communication stays focused, there is little unnecessary back-and-forth, and he still raises the important questions when needed. The result is a developer who delivers without overhead, and that's been genuinely valuable."
"We work with Vincent regularly on Laravel/Vue projects, both new builds and existing codebases that have grown over the years. In both cases, he delivers clean code and reliable implementation. We especially value how quickly he gets into mature projects and finds solutions that fit the existing state of the codebase. Absolutely recommended."
"Vincent was easy to work with and produced high-quality code. Would recommend."
"Vincent is one of those rare developers who can jump into a complex codebase, become useful immediately, and solve real problems without hand-holding.
He worked on performance optimization, customer-specific features, and a major custom reporting system. He also turned customer discussions into clear technical requirements and documented his work so the rest of the team could build on it.
Vincent does not just close tickets, he thinks through the product, the edge cases, and the customer impact. If you need a senior Laravel developer who can move a SaaS product forward, I strongly recommend him."
Have something I should look at?
Send the rough version. I can help turn it into a plan and then build it.
My Book: Rapid SaaS with Laravel
Launch Your SaaS in Days, Not Months. A practical, no-BS guide to building production-ready SaaS apps with Laravel 12: speed, billing, teams, AI integration, and the parts you need before custo

[truncated]

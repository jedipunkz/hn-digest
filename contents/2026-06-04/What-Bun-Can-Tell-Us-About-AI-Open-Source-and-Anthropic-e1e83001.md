---
source: "https://redmonk.com/sogrady/2026/06/04/bun-two-lessons/"
hn_url: "https://news.ycombinator.com/item?id=48399002"
title: "What Bun Can Tell Us About AI, Open Source and Anthropic"
article_title: "What Bun Can Tell Us About AI, Open Source and Anthropic – tecosystems"
author: "mkeeter"
captured_at: "2026-06-04T16:13:36Z"
capture_tool: "hn-digest"
hn_id: 48399002
score: 2
comments: 0
posted_at: "2026-06-04T14:15:39Z"
tags:
  - hacker-news
  - translated
---

# What Bun Can Tell Us About AI, Open Source and Anthropic

- HN: [48399002](https://news.ycombinator.com/item?id=48399002)
- Source: [redmonk.com](https://redmonk.com/sogrady/2026/06/04/bun-two-lessons/)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T14:15:39Z

## Translation

タイトル: AI、オープンソース、人類についてパンが教えてくれること
記事のタイトル: AI、オープンソース、人類についてパンが教えてくれること – tecosystems
説明: 昨年 12 月初旬、Anthropic は、小型で高速なオープンソース JavaScript ランタイムである Bun のメーカーである Oven を買収しました。これはパッケージ マネージャー、バンドラー、テスト ランナーでもありますが、Deno や Node.js のような Chrome の V8 ではなく、Safari の JavaScriptCore 上に構築された高速ランタイムとして最も成功しています。
[切り捨てられた]

記事本文:
AI、オープンソース、人類についてパンが教えてくれること – tecosystems
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
コンテンツにスキップ
AI、オープンソース、人類学についてパンが教えてくれること
スティーブン・オグレディ著 |
@sogrady |
2026 年 6 月 4 日
Twitterでシェア
Facebook 経由で共有する
Linkedin 経由で共有する
Reddit経由で共有する
昨年 12 月初旬、Anthropic は、小型で高速なオープンソース JavaScript ランタイムである Bun のメーカーである Oven を買収しました。これはパッケージ マネージャー、バンドラー、テスト ランナーでもありますが、Deno や Node.js のような Chrome の V8 ではなく、Safari の JavaScriptCore 上に構築された高速ランタイムとして最も成功しています。速度を重視した Node のドロップイン代替として構築され、Zig (後に Rust で書き直される) で書かれ、2022 年に初めてリリースされました。Cursor、Lovable、Windsurf、そしてもちろん Anthropic などの企業に AI の自然なユーザー層を獲得しました。また、Figma、The New York Times、Slack などの企業のスピード重視の制作システムにも浸透しました。一方、CircleCI や GitHub などのインフラストラクチャ プレーヤーは、どちらも昨年末にサポートを追加しました。
企業にとって重要な特殊ランタイムであるだけでなく、大小の AI 企業にとっても負荷に耐えるインフラストラクチャでもありました。
耐荷重があるかどうかにかかわらず、残念ながら、業界が依存している多くのオープンソースと同様に、買収時点ではその商業的見通しは明確ではありませんでした。ジャレッド・サムナー氏の買収発表のQ&Aは率直だった。
Q: 同じチームが今もフルタイムで Bun に取り組んでいますか?
A: はい。そして今、私たちは収益ゼロの小さなベンチャーキャピタル支援のスタートアップではなく、世界有数の AI ラボのリソースにアクセスできるようになりました。
全体として、買収は双方にとって非常に簡単でした。 Bun は i を受け取ります

Anthropic は、自社のサービスにとって戦略的なプロジェクトを直接管理できるようになります。無機質な買収ルートを選択することで、前者は豊富だが後者は貴重なものがほとんどない市場で時間を節約するために資金を費やした。買収後のプロジェクトの成果に興味があったので、私たちはその指標のいくつかを評価しました。
大まかに言えることは、買収によってプロジェクトが遅れているようには見えないということです。 npm から作成された以下のグラフは、Bun インストールのサブセットにすぎず、Homebrew、バイナリなどを介して直接インストールされたものは反映されていません。ただし、ここでアクセスできるサブセットでも、明確なストーリーが語られています。
Bun の成功と、Claude Code のような Bun を活用したプロジェクトの成功を正確に区別するのは難しいことに注意して、上記のグラフや同様のグラフを注意する必要があります。それでも、30 か月未満で月あたり 44 万 5,000 から 730 万まで 16 倍に成長するのは、それらが多いフィールドでの実行時間としては驚異的です。そして、ランタイムの成長が印象的だと思われるなら、その TypeScript の型定義はさらに印象的です。bun-type (ファーストパーティのネイティブ定義) は 53 倍に成長し、その TypeScript ラッパーは 234 倍に跳ね上がりました。
つまり、バンは成長しているのです。買収後はさらに成長が加速する可能性もあります。しかし問題は、その成長がどれほど持続可能であるかということです。それに答えるには、プロジェクトがどのように構築され、誰によって、そして時間の経過とともにどのように変化したかを内部的に調べる必要があります。結果として得られるデータセットからはさまざまな結論が導き出されますが、特に強調する価値のある結論が 2 つあります。
他の場所で議論されているように、Bun のコミット データを見ると最も明らかな点は、主に人間による貢献から主に AI への貢献への明らかな移行です。これは確かです

まったく秘密ではありません。 1か月前、サムナーはツイッターでこう語った。
「私たちはここ何ヶ月も自分でコードを入力していませんでした。事前に取得していたとしても、これはかなり正確でした。」
コミットチャートはこれを裏付けています。
昨年 8 月の時点では、ある時点でのプロジェクトのコミットの半分以上がボットによって作成されていました。買収後はそれを下回ることはほとんどなく、最高値は 80% を超えています。
AI と人間の月ごとの合計コミット数は次のとおりです。
ここでの傾向線は明確です。約 12 か月で、Bun は人間によって維持されていたプロジェクトから、主に機械によって作成されたプロジェクトに移行しました。もう少し詳しく説明すると、貢献者ごとのコミット数は次のとおりです。AI ですが、内部貢献者と外部貢献者に分けられています。
すぐに二次トレンドに到達しますが、やはり結論は避けられません。 Bun が中核的な AI インフラストラクチャであるのと同様に、AI は現在、Bun の中核的な貢献者です。
これにより、多くの疑問が生じますが、そのほとんどはまだ答えられていません。プロジェクトは長期にわたってどの程度保守可能ですか? AI に大きく依存することで、Bun チームが技術的負債と学習性無力感を抱えているとしたらどうなるでしょうか? AI は人間を犠牲にしてコミットされるコードの割合を増やし続けるのでしょうか、それとも時間の経過とともに自然な均衡が進化するのでしょうか?
AI への貢献はまだ 1 年足らずで、その半分は地球上で最も著名で重要な AI 企業の 1 つの従業員として行われています。サンプルが不十分であるため、これらの質問に対する答えはしばらくはわかりません。
しかし、AI の貢献の増加によってコア インフラストラクチャ製品がどのような影響を受けるかを考えるとき、Bun が監視すべき重要なデータポイントになることは明らかです。
おそらく、「プロジェクトには機械によって書かれたコードがますます多く含まれている」よりも興味深いでしょう。

Bun にとってこの買収がオープンソース プロジェクトとして何を意味するのか。 Bun は以前も現在も MIT ライセンスを取得しており、買収の発表ではプロジェクトに関して 4 つの関連する約束が行われています。
Bun はオープンソースであり、MIT ライセンスを維持します
Bun は引き続き非常に積極的にメンテナンスされています
同じチームが今でも Bun に取り組んでいます
Bun はまだ GitHub で公開されています
そのうち 3 つの約束は間違いなく守られています。 Bun はオープンソースのままであり、MIT ライセンスを取得しています。これは積極的に保守されており、GitHub 上に構築されています。一方、チームはそれぞれ別の道を歩んだようだ。
まず、AI の貢献が増加していることを受けて、プロジェクトに貢献した人間の総数のマクロ写真を見てみましょう。
その数は以前の約半分です。一方、外部貢献者の数は大幅に減少しました。
全体像を把握する人間の開発者の数が減少しているのと同様に、外部の開発者の数も減少しています。
ただし、これを状況に合わせて説明すると、外部貢献者の数値はより堅牢であるように見えますが、外部貢献者からの実際のコード貢献度は常に比較的控えめであり、コミットによる実際の貢献度の測定には問題がある性質があることは認められていました。
内部コミット数は買収直後に減少した後、以前と同じ大まかな数に戻りました。ただし、外部コミットはそうではありません。彼らの貢献は大幅に減少しました。これは驚くべきことではありません。小規模な独立系スタートアップ企業が維持するプロジェクトに貢献することと、資本力のある大規模な AI 巨大企業が維持するプロジェクトに貢献することは別問題です。
チームを団結させるという当初の約束に戻ると、元のコミッターの詳細なリストに上記の指標が示されていることがわかります。
買収前のオーブン従業員の識別可能な約 7 人のうち、少なくとも 4 人がクリアランスを持っています

rlyは去ったか、少なくとも貢献をやめた。別のアクティブなコミッターは、買収前の 745 件のコミットから、買収後のコミットは 1 件になりました。この小さな離散には多くの潜在的な理由が考えられますが、それらはプロジェクト、AI、買収、または上記のいずれともほとんど関係がない可能性があります。大規模なオープンソース プロジェクトに取り組む動機、関心、または許可は、さまざまな理由で変化する可能性があります。
しかし、最初に Bun を構築したのと同じチームではないことは確かです。人間が去り、AI が入ってきました。その交代サイクルが計画的で意図的なものかどうかは関係ありません。
Bun への人間の貢献者の数が減少している一方、マシンの貢献者の数が増加しているという事実は、それが比較的注目度の高いオープンソース インフラストラクチャ プロジェクトでなければ、それほど興味深いものではないでしょう。 Anthropic が今後どのようにプロジェクトの管理を行っていくのか、あるいは実際にプロジェクトの管理者としての役割を重視しているのかどうかは不明です。 Bun は、Anthropic に質問する機会です。Anthropic はオープンソースをどのように評価していますか?このプロジェクトにはどのような意図があるのでしょうか?
Anthropic の別のオープン プロジェクト、MCP のケースを考えてみましょう。 2024 年 11 月に初めて発売され、3 か月以内にこれが明確な業界標準であるということでコンセンサスが得られました。これは不条理で前例のない期間です。競合ベンダーから、11 月のリリース後の 1 月に MCP にこの地位を事実上付与すると言われるのは困難であり、衝撃的でさえありました。
この初期かつ前例のない成功にもかかわらず、中立財団に寄付されるまでさらに 10 か月かかりました。馴染みのない人のために説明すると、これは、熾烈な競争相手によって共同開発される標準化されたテクノロジーに必要なステップです。営利組織が単独で所有するプロジェクトに貢献する企業は、たとえあったとしてもほとんどありません。

第三者の開発をあなたの労働力で補助するのと同じだからです。
公平を期すために言うと、この期間は完全に不合理というわけではありません。 Kubernetes も同様に、最初のリリースから寄付まで 13 か月かかりました。しかし、Kubernetes は複数の競合するコンテナ オーケストレーション プロジェクトの中の 1 つでもあり、当時の明らかな業界標準からは程遠いものでした。したがって、延期と最終的な寄付は適切かつ戦略的でした。 MCP は標準化のより明白な候補でしたが、MCP までに最も急速に採用されたテクノロジである Docker よりも早かったのです。しかし、明らかなオープンソース標準の昇格が許可されるまでには、1 年以上かかりました。
ここで疑問が生じます。Anthropic と OpenAI などのそれに相当する企業は、企業のオープンソース成熟度曲線のどの位置にあるのでしょうか?スタートアップは、オープン ソース コードをベースに構築されているため、オープン ソース コードを消費者として理解します。彼らは通常、貢献やガバナンスなどを理解する必要があるため、理解しています。しかし、一般に、できるだけ早く移行することに重点を置いているスタートアップ企業は、オープンソースが企業レベルまたはエンタープライズレベルでどのように機能するかについてあまり詳しくありません。
この点において、人類学が傑出しているわけではありません。 Microsoft は数十年にわたり、オープンソースを口頭で攻撃してきました。 Google の初期の頃は、ソフトウェアに関する公開文書を、その背後にあるコードを公開せずに公開することで特徴づけられました。そして、オープンソース コミュニティの間での AWS の評判は、比較的最近になってより平和的に共存し、貢献する方法を学んだまでは、おそらく Microsoft よりも悪かったでしょう。
これらのベンダーとそれに先立つベンダーは、ミクロスケールではなくマクロスケールでオープンソースについて学ぶ必要がありました。ライセンスの選択、財団の役割、外部からの貢献を妨げるのではなく奨励するオープンソース プロジェクトの運営方法について。

同様の利点もあります。
Anthropic は Bun を標準化する必要がないため、これらの教訓を学ぶ必要がないという議論が考えられます。確かに、競合他社からのプロジェクトへの外部からの貢献は、今や AI から来ていると考えられます。プロジェクトの所有者がいつでもリリースできるボトルの中に大量のソフトウェアの魔神を抱えているのに、なぜわざわざ競合他社間で開発コストを償却したり、プロジェクトの管理を財団に譲渡したりする必要があるのでしょうか?
もちろん、これは、プロジェクトの標準化によってもたらされる主な価値がコードへの貢献であることを前提としています。これは標準化の目的に対する根本的な誤解です。外部コードの貢献は、プロジェクトを財団に寄付する主な動機ではありません。不必要で非生産的な市場の断片化を防ぐことが重要です。潜在的なユーザーも安心させられます。たとえば、企業は、より迅速に作成できるため、基盤からのソフトウェアを採用しません。そうするのは、主要なインフラストラクチャが単一のベンダーによって管理されないことを好むためです。
いずれにせよ、Anthropic がオープンソースを理解しているかどうか、そしてどの程度よく理解しているかについて興味がある人は、彼らによる Bun の管理と、それが時間の経過とともにどのように進化するかを観察することが有益になるでしょう。名誉のために言うと、このプロジェクトは、まだ進んでいないにもかかわらず、急速に成長しています。

[切り捨てられた]

## Original Extract

In early December last year, Anthropic acquired Oven, the makers of Bun, a small, fast, open source JavaScript runtime. It’s also a package manager, bundler and test runner but it’s had the most success as a fast runtime built on Safari’s JavaScriptCore rather than Chrome’s V8 like Deno and Node.js.
[truncated]

What Bun Can Tell Us About AI, Open Source and Anthropic – tecosystems
You are using an outdated browser. Please upgrade your browser to improve your experience.
Skip to Content
What Bun Can Tell Us About AI, Open Source and Anthropic
By Stephen O'Grady |
@sogrady |
June 4, 2026
Share via Twitter
Share via Facebook
Share via Linkedin
Share via Reddit
In early December last year, Anthropic acquired Oven, the makers of Bun, a small, fast, open source JavaScript runtime. It’s also a package manager, bundler and test runner but it’s had the most success as a fast runtime built on Safari’s JavaScriptCore rather than Chrome’s V8 like Deno and Node.js. Built as a drop-in replacement for Node focused on speed and written in Zig (later to be rewritten in Rust) and first released in 2022, it found a natural audience in AI with companies like Cursor, Lovable, Windsurf and, of course, Anthropic. It also made inroads into speed-focused production systems at companies like Figma, The New York Times and Slack. Infrastructure players like CircleCI and GitHub, meanwhile, both added support late last year.
In addition to being an important specialty runtime for enterprises, then, it was load bearing infrastructure for AI companies large and small.
Load bearing or no, its commercial prospects at the time of the acquisition – like so much of the open source the industry relies on, unfortunately – were less than clear. This Q&A from Jarred Sumner’s acquisition announcement was blunt:
Q: Is the same team still working on Bun full-time?
A: Yes. And now we get access to the resources of the world’s premier AI Lab instead of a small VC-backed startup making $0 in revenue.
On the whole, the acquisition was fairly straightforward for both parties. Bun receives an immediate capital return and a viable, long term path for support, while Anthropic gains direct control of a project strategic to their offerings. By going the inorganic acquisition route, it spent money to save time in a market with plenty of the former but precious little of the latter. Curious about how the project has fared post-acquisition, we’ve evaluated some of its metrics.
The high level takeaway is that the acquisition does not appear to have slowed the project. The below chart drawn from npm is merely a subset of Bun installs, and doesn’t reflect those installed directly, via Homebrew, binary or otherwise. Even the subset we’re able to access here tells a clear story however.
It is necessary to caveat the above and similar charts by observing that it’s difficult to precisely tease apart Bun’s success from that of projects that leverage it like Claude Code. Still, growing 16X from 445K/month to 7.3M in less than 30 months is impressive for a runtime in a field full of them. And if the runtime growth sounds impressive, the TypeScript type definitions for it are even more impressive – bun-types (the first party native definition) grew at 53X while its TypeScript wrapper jumped 234X.
Bun is growing, in other words. It may even be growing faster post-acquisition. But the question is: how sustainable is that growth? To answer it, it’s necessary to look under the hood at how the project is being built, by whom and how that’s changed over time. There are many different conclusions to be drawn from the resulting datasets, but there are two particularly worth highlighting.
As has been discussed elsewhere, the most obvious takeaway in looking at Bun’s commit data is the glaring transition from primarily human to primarily AI contributions. This is certainly no secret; a month ago, Sumner said on Twitter:
“ We haven’t been typing code ourselves for many months now. Even pre-acquisition this was pretty much accurate .”
The commits chart confirms this.
As early as last August, over half of the project commits at a given time were authored by a bot. Post-acquisition, it’s rarely less than that, and has peaked north of 80%.
Here are the total commits per month, AI vs human.
The trend line here is unambiguous: in approximately 12 months, Bun has transitioned from a project maintained by humans to one primarily authored by machines. To break that out in a little more detail, here are the commits per contributor: AI, but splitting up internal and external contributors.
We’ll get to the secondary trend there momentarily, but again, the conclusion is unavoidable. Just as Bun is core AI infrastructure, AI is now the core contributor to Bun.
This raises a host of questions that for the most part can’t be answered yet. How maintainable will the project be over the long term? What if any tech debt and learned helplessness is being accrued by the Bun team by relying so heavily on AI? Will AI continue to increase its percentage of code committed at the expense of humans, or will a natural equilibrium evolve over time?
It’s been barely a year of AI contributions, and half that as employees of one of the most visible and important AI companies on the planet. We won’t and can’t know the answers to these questions for some time, because the sample is insufficient.
But it seems clear that when looking at how core infrastructure products might be impacted by rising AI contributions, Bun will be an important datapoint to monitor.
Arguably more interesting than “project includes more and more code written by machines” is what the acquisition means for Bun as an open source project. Bun was and is MIT licensed, and the acquisition announcement made four related promises around the project:
Bun stays open-source & MIT-licensed
Bun continues to be extremely actively maintained
The same team still works on Bun
Bun is still built in public on GitHub
Three of those promises have undeniably been kept. Bun remains open source and MIT licensed. It is actively maintained, and built on GitHub. The team, on the other hand, appears to have gone its separate ways.
First, let’s look at a macro picture of the number of human contributors to the project, total, in the wake of the rising AI contributions.
That number is roughly half of what it was. The number of external contributors, for its part, has dropped off significantly.
Just as the number of human developers big picture has declined, so too has the number of external developers.
To put that in context, however, while their numbers appeared to be more robust, the actual code contributions from external contributors has always been relatively modest – even acknowledging the problematic nature of measuring actual contributions by commits.
After dipping immediately post-acquisition, internal commits have climbed back into the same rough number they occupied prior. External commits, however, have not. Their contributions have significantly declined. This is unsurprising. Contributing to a project maintained by a small, independent startup is a different matter than one maintained by a large, well capitalized AI juggernaut.
Getting back to the original promise of keeping the team together, we can see the above metrics manifested in a detailed list of the original committers.
Of ~7 identifiable pre-acquisition Oven employees, at least 4 have clearly departed or at least stopped contributing. Another active committer went from 745 pre-acquisition commits to 1 post-acquisition. There are many potential reasons for this mini diaspora, and they may have little if anything to do with the project, AI, the acquisition or any of the above. The motivation, interest – or permission – to work on a large open source project can change for a variety of reasons.
But it is certainly not the same team that originally built Bun. Humans left, AI has moved in – whether that replacement cycle was deliberate and intentional, or not.
The fact that the number of human contributors to Bun is down while the number of machine contributions up would be less interesting if it wasn’t a relatively high profile open source infrastructure project. It’s unclear how Anthropic will navigate its stewardship of the project moving forward, or whether in fact they care about their role as project steward. Bun is an opportunity to ask questions of Anthropic: how does it value open source? What are its intentions for the project?
Consider the case of another open project out of Anthropic, MCP. First launched in November of 2024, consensus within three months was that it was a clear industry standard – which is an absurd, unprecedented timeframe. It was difficult, even shocking, to be told by competitive vendors that they were effectively granting this status to MCP the January after a November release.
In spite of this early and unprecedented success, it took another ten months for it to be donated to a neutral foundation. For the unfamiliar, this is a necessary step for standardized technologies that will be jointly developed by otherwise fierce competitors. Few if any commercial organizations will contribute to a project solely owned by a third party because it’s tantamount to subsidizing their development with your labor.
To be fair, this timeframe is not totally unreasonable. Kubernetes likewise took thirteen months from initial release to donation. But Kubernetes was also one amongst multiple competitive container orchestration projects, and very far from an obvious industry standard at the time. The delay and ultimate donation, then, was appropriate and strategic. MCP was a much more obvious candidate for standardization, however, earlier even than Docker, the most rapidly adopted technology we’d seen up until MCP. But it still took over a year for an obvious open source standard to be permitted to ascend.
Which begs the question: where is Anthropic, and its counterparts like OpenAI, on the corporate open source maturity curve? Startups understand open source code as consumers because they are built on it. They generally understand contributions, governance and the like because they have to. But as a rule, startups focused on moving as quickly as possible are far less familiar with how open source works on a corporate or enterprise level.
Not that Anthropic stands out in this regard. Microsoft spent decades verbally assaulting open source. Google’s early years were marked by publishing open papers about software without releasing the code behind them. And AWS’ reputation amongst open source communities was arguably worse than Microsoft’s until relatively recently when it learned to more peacefully coexist and contribute back.
These vendors and those that preceded them have had to learn about open source on a macro, rather than micro-scale. About license choices, the role of foundations and how to run open source projects that encourage rather than discourage external contributions – as well as the benefits of same.
An argument that could be made is that Anthropic won’t have to learn these lessons because it doesn’t need to standardize Bun. Certainly any flow of would be external contributions to the project from competitors is arguably now coming from AI. Why bother amortizing development costs across competitors and giving up control of a project to a foundation when the project’s owner has a bunch of software genies in a bottle that it can release at any time?
That assumes, of course, that the primary value resulting from the standardization of a project is code contributions. Which is a fundamental misunderstanding of the purpose of standardization. External code contributions are not the primary incentive to donate a project to a foundation: preventing needless and unproductive market fragmentation is. As is reassuring potential users. Enterprises, for one, do not embrace software from a foundation because it’s written more quickly. They do so because they prefer their key infrastructure not be controlled by a single vendor.
In any event, those curious about whether and how well Anthropic understands open source would be well served by watching their stewardship of Bun and how it evolves over time. The project, to its credit, is growing apace even as it’s h

[truncated]

---
source: "https://brianguthrie.com/p/software-architecture-after-ai/"
hn_url: "https://news.ycombinator.com/item?id=48330939"
title: "Software Architecture After AI"
article_title: "Software architecture after AI | Gears Within Gears"
author: "bguthrie"
captured_at: "2026-05-30T11:36:50Z"
capture_tool: "hn-digest"
hn_id: 48330939
score: 2
comments: 0
posted_at: "2026-05-30T00:03:50Z"
tags:
  - hacker-news
  - translated
---

# Software Architecture After AI

- HN: [48330939](https://news.ycombinator.com/item?id=48330939)
- Source: [brianguthrie.com](https://brianguthrie.com/p/software-architecture-after-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T00:03:50Z

## Translation

タイトル: AI 以降のソフトウェア アーキテクチャ
記事のタイトル: AI 以降のソフトウェア アーキテクチャ |歯車の中の歯車
説明: AI は、建築とみなされるものの境界を移動させました。大量のコードはありません

記事本文:
Gears Inside Gears アーカイブ
について
歯車の中の歯車
私はソフトウェア、才能、複雑さ、システムの交差点について執筆しています。
AI 後の Unsplash ソフトウェア アーキテクチャに関する D Brz 氏の写真
ソフトウェア アーキテクチャのより有用な定義の 1 つは、「進化するアーキテクチャの構築」にあります。つまり、アーキテクチャとは定義上、変更するのが難しいものです。 1 私はこの定義が、最も単純であることは言うに及ばず、利用可能な最も誠実な枠組みであると常々感じてきました。それは、建築が美しさや正しさ、または居住する建築家のお気に入りのストーカー馬であるかのように見せかけているわけではありません。意思決定を「アーキテクチャ的」にするのは、概念的な重みではなく、覆すためのコストとビジネスへの影響であることを認めています。そして、「変更が難しい」ということは常に、根本的には実時間に関することであり、調整コスト、インシデントの軽減、認知負荷、ハンドオフの摩擦などです。ソフトウェア アーキテクチャは常に、設計上の問題に見せかけた労働問題でした。
AI は現在、大幅なコードレベルの変更を行うために必要な実時間を崩壊させています。以前は何か月もかかっていたことが、今では数日で完了します。ほとんどのコードレベルの決定を覆すコストが桁違いに下がったら、アーキテクチャはどうなるでしょうか? 2
何が起こるかというと、ソフトウェア アーキテクチャとみなされるものの境界が、場合によっては劇的に移動するということです。コードレベルの決定のほとんどは、もはやその中には存在しません。覆すためのコストは崩壊しており、間違えた場合の結果は、数か月ではなく数日の手戻りで測られるようになりました。内部に残るのはデータ、サービス境界、ユーザーの信頼ですが、難しい部分は決してコードではないため、これらは難しいままです。そして、コードレベルの議論によって押しつぶされていたいくつかの懸念が、今ではより明確に浮き彫りになっています。特に可観測性は再考に値する

機能の配信率が増加するにつれて、パンテオンが増加します。
コードのすべての行がクリーンで、本物のエンジニアがリファクタリングしていた頃
何十年もの間、コードレベルの決定は合法的なアーキテクチャでした。言語、フレームワーク、モジュール構造、永続化戦略は、議論してコミットする価値のある決定でした。再検討するには膨大な時間がかかり、チームの生産性に長期的な影響を与える可能性があるためです。これらを変えるには数か月、場合によっては数年の時間と労力がかかる可能性があり、大企業が方針を転換するまでに企業は存続し、消滅しました。リファクタリングでさえ、コードレベルの変更は可能だがコストがかかるという考えに基づいていました。コードの再構築にはスキルとリアルタイム性が必要であり、そのコストを管理するテクニックが必要でした 3 。
しかし、ソフトウェア実務者は何十年もの間、アーキテクチャ上の決定を日常的な決定に落とし込んできました。痛みを回避するのではなく、痛みに傾くことの効果は、チームに痛みに対処するツールを構築する動機を与え、これまでアーキテクチャだったものを汎用エンジニアが当然のように扱うものに変えることです。
データベースの移行が一般的になる前は、スキーマの決定は元に戻すことができず、おそらくアーキテクチャ上のものでした。多くの場合、DBA に調整を依頼する必要がありました 4。その後、Pramod が「進化的データベース設計」を執筆し、移行はあらゆる主要なフレームワークに組み込まれるようになり、DBA の役割はあまり目立たなくなり始めました。彼らが提供した判断と専門知識は本物 (そして実質的) でしたが、その市場価値は機械的なボトルネックによって高騰しました。ボトルネックが解消されると、専用ゲートのコストがより明確になり、その判断は一般的なエンジニアリングの役割に吸収されるため、多くのチームがより多くの影響力を得ることができました。継続的デリバリーはデプロイメントとリリースでも同じことを行いました

SEのエンジニア。特効薬は (今までは?) ないかもしれませんが、ツールの効率が少しずつ変化するたびに、以前は専門家が必要としていた決定が機械的に行われるようになりました。それぞれが、専門的な判断は機械コストの背後にある一般的なエンジニアリングスキルであることが多いことを明らかにしました。
AI は単にこの最新の例にすぎませんが、最も劇的なのは、一度に 1 つのカテゴリごとではなく、残っているコードレベルの決定のほとんどを一度に破棄するためです。最近の個人的な例: 私は、ベンダーもスタートアップだった NoSQL データベースに対して自分の (今は消滅した) スタートアップを書きましたが、それも (驚いたことに、驚いたことに) 消滅しました。私は Claude Code を指定し、いくつかのガイダンスを与えました。そして、データ層全体を従来の RDBMS に、本質的に完璧に数時間で移植しました。この種のことが一般的になっていることは知っていますが、それでも私は驚きました。退屈な仕事と本業の間に、宇宙の熱による死が起こる前には、私はそれを達成することはできなかったかもしれません。
これは特別な例ではありません。 Cloudflare のチームは、API コスト約 1,100 ドルで、1 週間以内に Next.js API サーフェスの 94% を再実装しました。 Christopher Chedeau は、1 か月で 100,000 行の TypeScript を Rust に移植しました。皆さんの多くも同様の変化を経験しています。
ある意味、これらの例は、優れた構造が重要であるという法則を証明しています。昨日コードを書き始めたわけではないので、私はクリーンなインターフェイス境界に対してデータ層を構築しました。そのため、ある意味、もちろん実装の交換は簡単でした。しかし、明確な境界がなくても、変更は基本的に機械的です。すべての呼び出しサイトを見つけ、すべての実装を変更し、正確さを検証します。トークンが増えれば、時間も増え、人間の介入も増えますが、大きな違いについて話しているわけではありません。数時間ではなく数日かもしれません。そして二次

エージェント開発の効果は、その上で検証を自動化できることです。プロセス自体に正当性チェックを組み込むことができます。
これらの例は、明確なインターフェイス、明確に定義された境界、機械的に検証可能な正確さなど、簡単なケースに偏っていることは明らかです。微妙なセマンティクス、不明確な境界、および深く絡み合ったビジネス ロジックを持つシステムは、AI を使用しても依然として変更が困難です。しかし、トレンドラインは重要です。 AI は結合、移行リスク、ロールアウトの複雑さを解消しませんが、永続的に感じられていたコード形成に関する多くの決定を降格し、残りの部分の監視と修正が劇的に容易になります。それは、成長するカテゴリーの変化を「もはや建築ではない」という領域に真っ向から位置づけることになる。
コードが境界の外に移動した場合、その中には何が残っているのでしょうか?
ソフトウェア アーキテクチャの包括的な分類を作成することは定義的にばかげていますが、私は変化の形を捉えていると思われる以下の 6 つのカテゴリに分け、それらを 2 つの軸 (決定を誤った場合の結果と、それを覆すためのコスト) に沿って分類することを試みました。これは直感レベルの話であり、厳密なデータではないので、ご容赦ください。変化を視覚化しようとしているだけです。
3 つの異なる理由による 3 つの動き。
一部の決定は、AI が決定を覆すためのコストを削減したため、降格されました。ローカル コード構造 (モジュール、フレームワーク、永続性、統合ワイヤリング) は、誤った決定を取り消すと数か月もの時間を費やす可能性があるため、アーキテクチャ上で重大な注意を払う必要がありました。もうその必要はありません。スケーラビリティと導入の姿勢が続きました。インフラストラクチャ トポロジとパフォーマンスの再作業は通常のコードよりも依然として困難ですが、AI 支援ツールを使用したルーチン エンジニアリングの範囲内にあります。これらの決定には依然として判断が必要ですが、その判断は無効です

もはや機械コストに囚われており、不良コールの影響は四半期や従業員数ではなく、時間やトークンで測定されています。
難しい部分はコードではなかったため、一部の決定はそのまま残りました。データ アーキテクチャ (所有権、整合性モデル、スキーマの進化) が動かなかったのは、データには重力があるためです。時間の経過とともに質量が蓄積され、誰も数え切れないほど多くのものが現在の形状に依存します。信頼とサービスの境界も移動しませんでした。セキュリティ違反と契約違反は事実上取り消すことができず (大企業は驚くべき量の違反を免れますが)、それらを元に戻すには、人間との調整、蓄積された状態の再形成、またはコードの変更では到達できない現実世界の影響を元に戻す必要があります。
明確な理由により、2 つの決定が昇格されました。量が増加するため、可観測性と動作検証が向上します。行ごとの欠陥が一定のままでも、コード量が 5 倍になった場合、システムが実際に何を行うかを検証できなかった場合の結果は、それに伴って増加します。さらに、行レベルの理解が同様に崩壊した場合 (闇のソフトウェア工場など)、すべての行を理解したかどうかに関係なく、システムが何を行うかを検証できる必要があります。モニタリングの実装は安価です。何を監視するか、どのように行動の正しさを検証するかについての決定はそうではありません。対照的に、ビジネス戦略と能力の連携はグラフ上でまったく動かなかった。それらは常に重大な結果をもたらすものであり、戦略上の失敗を逆転させるには常に高くつくものでした。変わったのは、建築家がようやく建築家と関わるための余裕を持てるようになったということです。コードレベルの議論が安価になったことにより、どの境界が競争上の優位性を生み出すかという問題は、フレームワークの議論によって埋もれなくなりました。
あなたは合理的に次のように主張することができます

コード構造は、まさに私がアーキテクチャと呼んでいるものを強制するメカニズムです。適切なモジュール境界は、API コントラクトを強制するのに役立ちます。適切な型システムは、データの不変性を保護するのに役立ちます。コード構造を気にするのをやめたら、気にしていると主張する契約が損なわれる危険があるのではありませんか?これは決定とその実行を混同していると思います。契約の形とその保証は難しい部分です。それらを変更するには人間と調整する必要があるため、リアルタイムのコストがかかります。それらを強制するコードはimplementationであり、実装は安価になりました。強制する契約に触れることなく、強制メカニズムを交換できます。それが私のスタートアップ移行作業で行われたことです。
これに関連して、取り組む価値のある反論があります。中間レベルの設計上の決定 (「このビジネス ロジックはサービス A に置くべきか、サービス B に置くべきか?」) は、時間の経過とともにシステム全体の柔軟性に蓄積され、それらの蓄積された決定を解きほぐすのは非常に困難です。これは真実ですが、常に真実です。そもそも、中央で制御できるものではありませんでした。チームは通常、ローカルの自律性とスループットを最適化するために、配置すべきと思われる場所に配置し、中央で管理しようとしても、その結果は常にある程度のずれが生じます。何が変わったのかというと、コードベースの検索インデックス (テーブルの重要性が急速に高まっている) に結び付けられた LLM が、実際にコードベースのすべてを見つけて、どのようにしてコードベースに行き着いたのかを説明し、コードベースの再編成を支援できることです。中間レベルの意思決定による蓄積された影響は、依然として重要であり、おそらくまだアーキテクチャ上の問題ではありますが、AI を使用すると軽減されるどころか、より扱いやすくなります。たとえ消えていないとしても、もつれを解くコストは下がりました。
パターンの増幅は運命ではない
AIがそうするのではないかという反論が聞こえてくるだろう。

オードの品質は、良い判断と悪い判断の両方を大量に増幅するため、それ以下ではなく、より重要です。 Addy Osmani 氏は、AI が生成したコードには人間が作成したコードよりも 75% 多くのロジック エラーと 2.74 倍の XSS 脆弱性があると報告しています。 PR は 18% 大きくなります。変更の失敗率は約 30% 増加します。 fast.ai の Rachel Thomas 氏は、私たちは「コーディングは自動化したが、ソフトウェア エンジニアリングは自動化していない」と主張しています。オスマニ氏の理解力の負債に関する枠組みは鋭い。AI を使用している開発者は理解力クイズのスコアが 17% 低く、「コードの生成が安くなったからといって、理解のスキップが安くなるわけではない」。しかし、それはリスクを間違った層に特定していると思います。
AI はコンテキスト ウィンドウのライト コーン内を見ることができますが、製品全体を見ることはまだできません。出荷したばかりの機能が実際に顧客が抱えている問題を解決するのか、それとも先週の火曜日に導入した微妙な行動の後退が徐々に信頼を損なっているのかを知ることはできません。これは、私が別の場所で、技術的負債とは区別される製品負債、つまりシステムが行うこととユーザーがシステムに必要とすることとの間のギャップに対する無知の蓄積と呼んだものです。
しかし、蓄積された無知に対する答えは、コード構造を改善することではありません。それは、より優れたハーネス、より優れた可観測性、より優れた動作検証です。これに必要な規律は私です

[切り捨てられた]

## Original Extract

AI has moved the boundary of what counts as architecture. A lot of code isn

Gears Within Gears Archive
About
Gears Within Gears
I write on the intersection of software, talent, complexity, and systems.
Photo by D Brz on Unsplash Software architecture after AI
One of the more useful definitions of software architecture comes from Building Evolutionary Architectures : architecture is definitionally the stuff that’s hard to change. 1 I’ve always found this definition to be the most honest framing available, to say nothing of the simplest. It doesn’t pretend architecture is about beauty or correctness or your resident architect’s favorite stalking-horse. It acknowledges that what makes a decision “architectural” is not its conceptual weight but its cost to reverse and its business impact . And “hard to change” has always been, at root, about wall-clock time: coordination cost, incident mitigation, cognitive load, handoff friction. Software architecture has always been a labor problem dressed up as a design problem.
AI has now collapsed the wall-clock time required to make substantial code-level changes. Things that used to take months now take days. What happens to architecture when the cost to reverse most code-level decisions drops by something like an order of magnitude? 2
What happens is that the boundary of what counts as software architecture moves, in some cases dramatically. Most code-level decisions are no longer inside it; their cost to reverse has collapsed, and the consequences of getting them wrong are measured in days of rework rather than months. What stays inside is data, service boundaries, and user trust, which remain hard because the hard part was never the code. And a few concerns crowded out by code-level debates now stand in sharper relief; observability in particular deserves a reconsideration in the pantheon as the rate of feature delivery increases.
When every line of code was clean, and real engineers refactored
For decades, code-level decisions were legitimately architectural. Languages, frameworks, module structures, and persistence strategies were decisions worth debating and committing to, because revisiting them could cost a staggering amount of time, and had long-term implications for the productivity of the team. Changing these things could cost months or even years of time and effort, and companies lived and died in the time it took large firms to reverse course. Even Refactoring was predicated on the idea that code-level change was possible but costly; restructuring code took skill and real time, and you needed techniques to manage that cost 3 .
But software practitioners have been collapsing architectural decisions into routine ones for decades. The effect of leaning into pain, rather than avoiding it is to incentivize teams to build tooling that addresses it, turning what used to be architecture into something a general-purpose engineer handles as a matter of course.
Before database migrations were commonplace, schema decisions were irreversible, and presumptively architectural; they often required DBAs to orchestrate them 4 . Then Pramod wrote Evolutionary Database Design , migrations got folded into every major framework, and the DBA role started to become less visible. The judgment and expertise they provided were real (and substantial), but its market value was inflated by the mechanical bottleneck. Once the bottleneck was removed, the costs of a dedicated gate became more visible and the judgment got absorbed into the general engineering role, which gave many teams more leverage. Continuous delivery did the same thing for deployment and release engineers. There may be no silver bullets (until now?), but each small shift in tool efficiency took a category of decision that used to require a specialist and made it mechanical. Each revealed that a specialized judgment was often general engineering skill trapped behind mechanical cost.
AI is simply the latest example of this, but it’s the most dramatic, because it collapses most remaining code-level decisions at once rather than one category at a time. A recent personal example: I wrote my (now-dead) startup against a NoSQL database whose vendor was also a startup, which (surprise, surprise) also died. I pointed Claude Code at it, gave it some guidance, and it ported the entire data layer to a conventional RDBMS, essentially flawlessly, in hours . I know this sort of thing has become commonplace, but it still surprised me: between the tedium of the work and my day job, I might never have accomplished it before the heat-death of the universe.
This is not an isolated example. Cloudflare’s team reimplemented 94% of the Next.js API surface in under a week for roughly $1,100 in API costs. Christopher Chedeau ported 100,000 lines of TypeScript to Rust in a month. Many of you have experienced similar shifts.
In some ways, these examples prove the rule that good structure matters: I built my data layer against a clean interface boundary, because I didn’t start writing code yesterday, so in some sense, of course swapping the implementation was straightforward. But even without a clean boundary, the change is fundamentally mechanical: find all the call sites, change all the implementations, verify correctness. More tokens, more time, and yeah, more human intervention, but we’re not talking about a vast difference; maybe days instead of hours. And the second-order effect of agentic development is that you can automate verification on top of it; you can build correctness-checking into the process itself.
These examples are admittedly biased toward the easy case: clean interfaces, well-defined boundaries, mechanically verifiable correctness. Systems with subtle semantics, unclear boundaries, and deeply entangled business logic remain harder to change, even with AI. But the trend line matters. AI does not erase coupling, migration risk, or rollout complexity, but it does demote a lot of code-shaping decisions that used to feel permanent, and make the rest dramatically easier to monitor and fix. That places a growing category of change squarely in the territory of “not architecture anymore.”
If code has moved outside the boundary, what’s still inside it?
It’s definitionally ridiculous to create a comprehensive taxonomy of software architecture, but I’ve broken out six categories below that I think capture the shape of the shift, and attempted to classify them along two axes: consequence of getting a decision wrong, and cost to reverse it. This is gut-level stuff, not hard data, so bear with me; I’m just trying to visualize the shift.
Three movements, for three distinct reasons.
Some decisions got demoted because AI collapsed the cost to reverse them. Local code structure (modules, frameworks, persistence, integration wiring) used to command serious architectural attention because reversing a bad decision could cost months of calendar time; they don’t have to anymore. Scalability and deployment posture followed: infrastructure topology and performance rework are still harder than ordinary code, but within reach of routine engineering with AI-assisted tooling. These decisions still require judgment, but that judgment is no longer trapped behind mechanical cost, and the consequences of a bad call are measured in hours and tokens rather than quarters and headcount.
Some decisions stayed put because the hard part was never the code. Data architecture (ownership, consistency models, schema evolution) didn’t move because data has gravity: it accumulates mass over time, and more things depend on its current shape than anyone can enumerate. Trust and service boundaries didn’t move either: security breaches and contract violations are effectively irreversible (though large corporations get away with a shocking amount of this), and reversing them requires coordinating with human beings, reshaping accumulated state, or undoing real-world consequences that code changes cannot reach.
Two decisions got elevated, for distinct reasons. Observability and behavioral verification rise because volume is rising: if defects per line stay constant but code volume quintuples, the consequence of failing to verify what the system actually does rises with it. Furthermore, if line-level comprehension similarly collapses (as in a dark software factory), you need to be able to verify what the system does regardless of whether you understand every line. The implementation of monitoring is cheap; the decision about what to watch and how to verify behavioral correctness is not. Business strategy and capability alignment, by contrast, didn’t move on the chart at all; they were always high-consequence, and reversing a strategic misstep has always been expensive. What changed is that architects finally have the headspace to engage with them. With code-level debates cheaper, the question of which boundaries create competitive advantage is no longer crowded out by framework arguments.
You could reasonably argue that code structure is the enforcement mechanism for precisely the things I’m calling architectural. Good module boundaries help enforce API contracts; good type systems help protect data invariants. If you stop caring about code structure, don’t you risk undermining the contracts you claim to care about? I think this confuses the decision with its implementation. The shape of the contract, and its guarantees, are the hard part; changing them costs real time, because you have to coordinate with human beings to do it. The code that enforces them is implementation , and implementation is now cheap. You can swap out the enforcement mechanism without touching the contract it enforces; that’s what my startup migration work did.
There’s a related objection worth engaging: mid-level design decisions (“should this business logic live in service A or service B?”) accumulate over time into the overall malleability of a system, and those accumulated decisions are genuinely hard to untangle. This is true, but it’s always been true. It was never centrally controllable in the first place: teams generally put shit where it seems like it should go, optimizing for local autonomy and throughput no matter how much you try to govern it centrally, and the result is always some degree of drift. What’s changed is that an LLM strapped to a codebase search index (which is rapidly becoming table stakes) can actually find all of it, reason about how it ended up there, and help you reorganize it. The accumulated impact of mid-level decisions, while still important and probably still architectural, is more tractable with AI, not less; the cost of untangling it has dropped, even if it hasn’t vanished.
Pattern amplification is not destiny
You will hear the counterargument that AI makes code quality more important, not less, because it amplifies both good and bad decisions at volume. Addy Osmani reports that AI-generated code has 75% more logic errors and 2.74x more XSS vulnerabilities than human-written code; PRs are 18% larger; change failure rates are up roughly 30%. Rachel Thomas at fast.ai argues that we’ve “automated coding, but not software engineering.” Osmani’s framing of comprehension debt is sharp: developers using AI scored 17% lower on comprehension quizzes, and “making code cheap to generate doesn’t make understanding cheap to skip.” But I think it locates the risk at the wrong layer.
AI can see within the light-cone of a context window, but it cannot yet see the whole product. It cannot tell you whether the feature you just shipped actually solves the problem your customer has, or whether the subtle behavioral regression you introduced last Tuesday is slowly eroding trust. This is what I’ve elsewhere called the accumulated ignorance of product debt, as distinct from tech debt: the gap between what the system does and what users need it to do.
But the answer to accumulated ignorance isn’t better code structure; it’s better harnesses, better observability, better verification of behavior . The discipline that this demands is i

[truncated]

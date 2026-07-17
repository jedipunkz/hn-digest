---
source: "https://stantyan.com/blog/portable-ai-memory-or-permanent-lock-in/"
hn_url: "https://news.ycombinator.com/item?id=48949018"
title: "Portable AI Memory or Permanent Lock-In"
article_title: "Portable AI Memory or Permanent Lock-In: The Architectural Choice That Will Define AI's Next Decade | Stan Tyan"
author: "stantyan"
captured_at: "2026-07-17T17:01:26Z"
capture_tool: "hn-digest"
hn_id: 48949018
score: 1
comments: 0
posted_at: "2026-07-17T16:11:19Z"
tags:
  - hacker-news
  - translated
---

# Portable AI Memory or Permanent Lock-In

- HN: [48949018](https://news.ycombinator.com/item?id=48949018)
- Source: [stantyan.com](https://stantyan.com/blog/portable-ai-memory-or-permanent-lock-in/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:11:19Z

## Translation

タイトル: ポータブル AI メモリまたは永久ロックイン
記事のタイトル: ポータブル AI メモリまたは永久ロックイン: AI の次の 10 年を定義するアーキテクチャの選択 |スタン・ティアン
説明: AI メモリは新しいベンダー ロックインです。 1 つのベンダーによるオープン フォーマットだけでは十分ではありません。このカテゴリには中立的な交換標準が必要であり、規制もこれに同意しています。

記事本文:
ポータブル AI メモリまたは永続的なロックイン: AI の次の 10 年を定義するアーキテクチャの選択 | Stan Tyan Stan Tyan Stan Tyan プロジェクトについて ブログ お問い合わせ ⌘ Ctrl K AI
ポータブル AI メモリまたは永久ロックイン: AI の次の 10 年を決定するアーキテクチャの選択
AI メモリは新たなベンダー ロックインです。 1 つのベンダーによるオープン フォーマットだけでは十分ではありません。このカテゴリには中立的な交換標準が必要であり、規制もこれに同意しています。
Claude Projects 内で 18 か月を費やして作業したチームを想像してください。何百もの会話、数十回の反復を経て洗練されたプロジェクトの指示、コードベース、顧客、既に行われた決定とその理由に関する一連のコンテキストの蓄積。そのとき、リーダーシップは当然の質問をします：私たちは双子座を評価すべきでしょうか？そして、AI スタックの所有者からの正直な答えは不快なものです。定期購入は簡単にキャンセルできます。 18 か月間蓄積されたコンテキストには移行パスがありません。会話を JSON としてエクスポートしたり、記憶をテキストの概要としてエクスポートしたりすることはできますが、ツールが実際にあなたについて知っていたことを再構築することはできません。 AI メモリのポータビリティは、実用的な機能としては 2026 年 7 月には存在しません。移行計画は最初からやり直すことになります。
この夏何が変わったかというと、最初のベンダーが気づきました。ある企業はメモリ エンジン全体をオープンソース化し、顧客がエージェントの知っている情報を検査してセルフホストできるようにしました。もう 1 つは、git セマンティクス ( commit 、 Branch 、 merge 、 Push 、 Pull ) を備えたメモリを出荷し、それ自体をポータブル メモリ層として販売します。どちらの動きも真の進歩であり、以下で両方を全面的に評価します。どちらも、何がまだ欠けているかを正確に示しています。 1 つの製品だけが書き込む形式は移植性がありません。これは十分に文書化された方言です。
この作品のテーマは簡単に述べることができ、10 年間にわたって

その結果、メモリは新しいベンダー ロックインであり、メモリの移植性はこの 10 年間のアーキテクチャ上の決定であり、このカテゴリには単一ベンダーが所有しない中立的な交換標準が必要です。まだ誰もその標準を構築していません。これは、なぜ誰かがそうしなければならないのか、それに何を含めなければならないのか、そしてなぜヨーロッパの規制時計がすでに始まっているのかについての議論です。
新しいロックインはメモリ層にあります #
モデル自体は、過去 2 年間のある時点でスイッチングコストにならなくなりました。フロンティア モデルは数か月ごとに互いに飛び越し、API の形状はほぼ同一で、よく構築された抽象化レイヤーの背後でプロバイダーを交換するのは午後の作業です。モデルがどれだけ置き換え可能になったかを疑うなら、ジューンが私たちの代わりに実験を行ってくれました。 6月12日、米国政府は輸出管理指令を出し、Anthropicに対し最新の主力モデルを数時間以内に世界中で生産停止にすることを強制した。これらを採用したチームは数日以内に古いモデルに戻し、出荷を続けました。アクセスは 19 日後の 7 月 1 日に復旧しました。市場で最も高性能なフロンティア モデルが 3 週間近く姿を消し、エコシステムがそれを吸収しました。 AI ツールがチームについて学習したすべての一晩の損失を吸収することを想像してみてください。その思考実験が議論のすべてだ。
したがって、微分はスタックの上に移動しました。プラットフォームが知っていることは、競合他社がより優れたモデルを出荷しても複製できない唯一の資産であり、すべての主要ベンダーはこれを理解しています。 AI エージェントが設計上忘れる理由と、キャプチャされたコンテキストがどこに存在するかについて私が説明したパターンには、第 2 の行為が含まれています。ステートレス API 呼び出しごとにアプリケーションが苦労して再構築するのと同じコンテキストが、各プラットフォーム独自のメモリ機能内に永続的にキャプチャされています。キャプチャはゲです

めちゃくちゃ便利。構造的には堀でもあります。
現在の最も明確な例は、6 月 23 日に Anthropic が Claude Tag をローンチしたときです。Claude は、Slack 内の永続的なチームメイトとして機能します。チャネル スコープのメモリは、動作中に蓄積され、許可が与えられると組織のチャネル全体に拡張できます。印象に残る商品です。また、これは Slack ワークスペースと Anthropic に結び付けられて生きており、エクスポートが文書化されていないメモリでもあります。実行されるたびに、会社の仕組みについてさらに学習し、他のものを使用するコストは毎週増加します。それが機能するための罠として誰もそれを設計する必要はありませんでした。
ここの経済学は古いです。データベース、ERP、およびクラウド データ ウェアハウスはすべて同じ戦略を実行しました。つまり、製品は交換可能ですが、蓄積された状態は交換可能ではなく、状態は出口に織り込まれます。新しいのは、キャプチャされる対象の範囲が広いことです。データ ウェアハウスにはテーブルが保持されます。 AI メモリ層には、好み、決定、慣例、選択の背後にある理由、誰がいつ何を言ったかなど、組織の考え方が保存されます。スイッチングコストはもはやデータではありません。これは組織の知識であり、1 つのベンダーのみが読み取れる形式です。
一枚の絵の中の反転。モデル切り替えコストは下がり続けています - 2026 年 6 月には、主力モデルが 19 日間消滅し、チームが交代するだけの可能性があることが証明されました。プラットフォーム上で週ごとにメモリの切り替えコストが増大し、曲線間の拡大するギャップが堀になります。
AI のメモリが作り出す 3 種類のロックイン
3 種類の記憶はそれぞれ異なる方法で固定され、そこから抜け出そうとするときの痛みも異なるため、何が蓄積しているのかを正確に知ることができます。
行動のロックインは、自分の好み、スタイル、何百回も修正したことなど、学習した層です。 ChatGPT の Memory は標準的な消費者の例であり、

そしてそれは標準的な一方通行のドアでもあります。 OpenAI 独自のメモリに関する FAQ では、メモリの表示、管理、削除について説明しています。エクスポートはありません。アカウントレベルのデータエクスポートにより、会話履歴が得られます。製品が応答をパーソナライズするために実際に使用する蒸留されたメモリは内部に残ります。 1 年間の日常使用を通じてシステムがユーザーについて学習したことは、構造上、ユーザーが取得できるものではありません。
コンテキスト ロックインは、プロジェクト履歴、確立された事実、意思決定とその根拠、進行中のすべての実行状態などの作業状態レイヤーです。これは、チームがスイッチを評価するときに最初に感じるレイヤーです。月曜日の朝、ブリーフィングなしでツールが役立つのはこのレイヤーだからです。ここは、業界における唯一の真の相互運用性の先例が生きている場所でもあります。 Anthropic は、ChatGPT、Gemini、または Grok から持ち込まれたメモリを受け入れるメモリ インポートを出荷していますが、Google は独自の ZIP ベースのインポートで答えました。これらはプラットフォーム間の本当の扉であり、この春に小規模な移行の波を引き起こすほど重要でした。しかし、その仕組みを見てください。ユーザーは、古いアシスタントに覚えていることを書き出すよう求めるプロンプトを実行し、その散文を新しいアシスタントに貼り付けるかアップロードします。これはトラフィックの恩恵を受けるベンダーによって構築された入口であり、すべてをテキストに平坦化し、輸入ベンダーが構築した方向に沿って移動します。入口ランプは標準ではありません。
人間関係のロックインは、誰がどの事実を確立し、誰がいつ何を決定し、どの人の修正が誰の以前の主張に取って代わられるかなど、チームの設定にのみ存在するため、ほとんど誰も考慮しない層です。これが来歴であり、記憶された事実の山を、信頼して監査できる組織の記憶に変えるものです。それは私を共有し、チャネルスコープを設定した層でもあります

多くの製品は最も早く蓄積され、露出は最も少なくなります。記憶が個人的なノートではなくチームメイトになると、誰がいつ何を言ったかのグラフが建物内で最も価値があり、最も持ち出しにくいものになります。エクスポートできない意思決定証跡は、要求に応じて作成できない意思決定証跡であるため、これは監査人や法務チームが最終的に尋ねることになる層でもあります。
3 種類のロックイン、3 つの異なる離脱コスト。行動の記憶は失うと厄介で、コンテキストの記憶は失うと高くつきます。組織がプレーン テキストのエクスポートから再構築できないのは、関係の記憶 (誰が何を確立したかの出自グラフ) です。
各レイヤーは前のレイヤーと複合します。そして重要なのは、複雑化することです。競争圧力が逆方向に作用するため、これはいずれもベンダーが競争圧力の下で修正するバグではありません。リテンションの経済性による報酬の獲得。これに反対する勢力は 2 つだけです。調達においてポータビリティを要求する顧客と、それを義務付ける規制当局です。両方については以下で説明します。まず、早期に行動を起こしたベンダーは称賛に値します。
気づいたベンダー (そしてまだ十分ではない理由) #
今年出荷された 2 つの製品は、メモリの所有権を真剣に考慮しており、この議論の正直なバージョンは、それらを管理することから始めなければなりません。
Cognee 1.0 は、完全なオープンソースのメモリ エンジンとして 6 月 26 日にリリースされました。これは、単一の Postgres インスタンス上で実行できる型付きナレッジ グラフであり、4 つの動詞 ( remember 、 remember 、 Improvement 、 remember ) を中心に構築された API と、独自の COGX アーカイブ形式へのエクスポート パスを備えています。立ち上げの投稿には、その理由が明確に記載されています。チームは、コグニー氏の言葉を借りれば、「検査したり、ホストしたり、持ち運んだりできないブラックボックス」に自分たちのビジネスの記憶を渡すつもりはなかったのです。これはベンダーが公に述べた正しい診断です。自己

-hosting は、メモリがデータベースに保存されていることを意味します。オープンソースとは、それを記述したコードのすべての行を読み取ることができることを意味します。すべてのメモリ製品がこの基準を満たしていれば、この記事の半分は不要になります。
ByteRover は開発者ツール側から同じ問題を攻撃し、自らをコーディング エージェント用のポータブル メモリ層と呼んでいます。そのメモリは、プレーンなマークダウンとして保存され、git セマンティクスでバージョン管理された階層コンテキスト ツリーです。メモリの変更をコミットし、分岐し、マージし、マシンやチームメイト間でプッシュおよびプルします。ラップトップの交換によってチームが 1 か月間エージェントのコンテキストを失うのを見たことがある人なら誰でも、なぜこれが正しい形なのかすぐに理解できます。任意のエディターで読み取ることができ、コードの動きに合わせて履歴を保存できるファイル。誰がメモリを保持すべきかについての設計上の本能としては、それはまさに正しいです。
どちらのベンダーも、最初に移行したことは全面的に評価されますが、単一の実装では移植性を解決できないため、どちらも移植性を解決していません。オープンソースによりフォーマットが検査可能になります。ベンダー間でメモリを移植できるわけではありません。移植性はエコシステムの特性であり、コードベースの特性ではありません。 COGX アーカイブは Cognee で読み取ることができます。 ByteRover コンテキスト ツリーは ByteRover にとって意味があります。明日どちらかの会社が消滅したとしても、あなたの記憶は判読できるでしょう - 既存企業よりも大幅に改善されています - しかし、その形状をネイティブに書き込んだり読み込んだりするものは他にないため、依然として立ち往生するでしょう。 1 つの実装のみが記述する形式は標準ではなく方言です。そして、明確にしておきたいのですが、これはどちらのベンダーも自社のオープン性に異議を唱えることなく反論できない議論であり、それがまさにこれが正しいテストである理由です。
需要シグナルはベンダーを超えて広がります。 2026 年前半だけでも、少なくとも 3 つの独立した交換提案が登場しました。MIF、個人開発者のメモリ交換仕様。メモリワイヤ

e 、メモリ操作に関する学術ワイヤ形式の提案。もう 1 つはコミュニティ主導の Open Memory Protocol であり、ベンダー中立であると説明されています。ベンダー独自のエクスポート形式である COGX を追加すると、6 か月以内に同じ質問に対して 4 つの異なる答えが得られます。それらのどれも、組織的なガバナンス、2 番目の独立した実装、またはベンダーの採用を持っていません。私はここでそれらをレビューしません、そしてそれは意図的です：波はその中のどのエントリよりも重要です。無関係な 4 人の関係者が同じ欠落部分を独自に発明すると、その部分は欠落します。
つまり、ギャップとは、正確に言うと、誰もフォーマットを提案していないわけではないということです。それは、どのフォーマットも中立的な統治下に置かれ、独立して 2 度実施されていないということです。このギャップはよく知られた形であり、業界は以前にもそのギャップを埋めてきました。
ポータビリティがエクスポート ボタンよりも難しい理由 #
ポータビリティに対する批判に対するどのプラットフォームの答えも同じです。データのエクスポートがあるのです。そして、これらのエクスポートはどれも同じように失敗します。テキストは保存するが構造を破棄するメモリ アーカイブが知識を削除しながら単語をエクスポートするためです。実際のポータビリティには 6 つの技術要件があり、このリストが要件を分けるものであるため、それぞれについて具体的にする価値があります。

[切り捨てられた]

## Original Extract

AI memory is the new vendor lock-in. Open formats from one vendor are not enough - the category needs a neutral interchange standard, and regulation agrees.

Portable AI Memory or Permanent Lock-In: The Architectural Choice That Will Define AI's Next Decade | Stan Tyan Stan Tyan Stan Tyan About Projects Blog Contact ⌘ Ctrl K AI
Portable AI Memory or Permanent Lock-In: The Architectural Choice That Will Define AI's Next Decade
AI memory is the new vendor lock-in. Open formats from one vendor are not enough - the category needs a neutral interchange standard, and regulation agrees.
Picture a team that has spent 18 months working inside Claude Projects. Hundreds of conversations, project instructions refined over dozens of iterations, an accumulated body of context about the codebase, the customers, the decisions already made and why. Then leadership asks a reasonable question: should we evaluate Gemini? And the honest answer from whoever owns the AI stack is uncomfortable. The subscription is easy to cancel. The 18 months of accumulated context has no migration path. You can export your conversations as JSON and your memory as a text summary, but nothing on the other side can reconstruct what the tool actually knew about you. AI memory portability, as a practical capability, does not exist in July 2026. Starting over is the migration plan.
Here is what changed this summer: the first vendors have noticed. One open-sourced its entire memory engine so customers can inspect and self-host what their agents know. Another ships memory with git semantics - commit , branch , merge , push , pull - and markets itself as the portable memory layer. Both moves are real progress, and I will give both full credit below. Both also illustrate, precisely, what is still missing. A format that only one product writes is not portability. It is a well-documented dialect.
The thesis of this piece is simple to state and has a decade of consequences: memory is the new vendor lock-in, memory portability is the architectural decision of the decade, and the category needs a neutral interchange standard that no single vendor owns. Nobody is building that standard yet. This is the argument for why someone must, what it has to contain, and why the regulatory clock in Europe has already started.
The new lock-in is at the memory layer #
The models themselves stopped being the switching cost some time in the last two years. Frontier models leapfrog each other every few months, the API shapes are near-identical, and swapping providers behind a well-built abstraction layer is an afternoon of work. If you doubt how replaceable models have become, June just ran the experiment for us. On June 12, the US government issued an export control directive that forced Anthropic to pull its newest flagship models offline worldwide within hours. Teams that had adopted them fell back to older models within days and kept shipping. Access was restored on July 1, nineteen days later. A frontier model - the most capable one on the market - vanished for nearly three weeks, and the ecosystem absorbed it. Try to imagine absorbing the overnight loss of everything your AI tools have learned about your team. That thought experiment is the whole argument.
So differentiation moved up the stack. What a platform knows about you is the one asset a competitor cannot replicate by shipping a better model, and every major vendor understands this. The pattern I laid out in why AI agents forget by design - and where the captured context lives now has a second act: the same context that applications painstakingly rebuild on every stateless API call is now being captured, persistently, inside each platform’s own memory features. The capture is genuinely useful. It is also, structurally, a moat.
The clearest current example arrived on June 23, when Anthropic launched Claude Tag : Claude as a persistent teammate inside Slack, with channel-scoped memory that accumulates as it works and can extend across an organization’s channels when granted permission. It is an impressive product. It is also memory that lives bound to the Slack workspace and to Anthropic, with no documented export. Every week it runs, it learns more about how a company works - and every week, the cost of ever using anything else grows. Nobody had to design that as a trap for it to function as one.
The economics here are old. Databases, ERPs and cloud data warehouses all ran the same play: the product is replaceable, the accumulated state is not, and the state is priced into the exit. What is new is the breadth of what gets captured. A data warehouse holds your tables. An AI memory layer holds how your organization thinks - preferences, decisions, conventions, the reasoning behind choices, who said what and when. The switching cost is no longer your data. It is your institutional knowledge, in a shape only one vendor can read.
The inversion in one picture. Model switching cost keeps falling - June 2026 proved a flagship can vanish for nineteen days and teams just swap. Memory switching cost compounds with every week on the platform, and the widening gap between the curves is the moat.
Three kinds of lock-in your AI memory creates #
It helps to be precise about what accumulates, because the three kinds of memory lock in differently and hurt differently when you try to leave.
Behavioral lock-in is the learned layer: your preferences, your style, the corrections you have made a hundred times. ChatGPT’s Memory is the canonical consumer example, and it is also the canonical one-way door. OpenAI’s own Memory FAQ covers viewing, managing and deleting memories - there is no export. The account-level data export gives you your conversation history; the distilled memory the product actually uses to personalize responses stays inside. What the system learned about you over a year of daily use is, by construction, not yours to take.
Context lock-in is the working-state layer: project history, established facts, decisions and their rationale, the running state of everything in flight. This is the layer teams feel first when they evaluate a switch, because it is the layer that makes the tool useful on Monday morning without a briefing. It is also where the industry’s one genuine interoperability precedent lives. Anthropic ships a memory import that accepts memories brought over from ChatGPT, Gemini or Grok, and Google answered with a ZIP-based import of its own. Credit where due: these are real doors between platforms, and they mattered enough to trigger a minor migration wave this spring. But look at the mechanism. The user runs a prompt that asks the old assistant to write out what it remembers, then pastes or uploads the prose into the new one. It is an on-ramp built by the vendor who benefits from the traffic, it flattens everything to text, and it moves in whichever direction the importing vendor built. An on-ramp is not a standard.
Relationship lock-in is the layer almost nobody prices in, because it only exists in team settings: who established which fact, who decided what and when, which person’s correction superseded whose earlier claim. This is provenance, and it is what turns a pile of remembered facts into organizational memory you can trust and audit. It is also the layer that shared, channel-scoped memory products accumulate fastest and expose least. When memory becomes a teammate rather than a personal notebook, the who-said-what-when graph becomes the most valuable and least exportable thing in the building. It is also the layer auditors and legal teams will eventually ask about, because a decision trail you cannot export is a decision trail you cannot produce on request.
Three kinds of lock-in, three different exit costs. Behavioral memory is annoying to lose, context memory is expensive to lose, and relationship memory - the provenance graph of who established what - is the one organizations cannot rebuild from an export of plain text.
Each layer compounds on the previous one. And the compounding is the point: none of this is a bug a vendor will fix under competitive pressure, because the pressure runs the other way. Retention economics reward capture. Only two forces push against it - customers who demand portability in procurement, and regulators who mandate it. Both are covered below. First, the vendors who deserve credit for moving early.
The vendors who noticed (and why it’s not enough yet) #
Two products shipped this year that take memory ownership seriously, and the honest version of this argument has to start by steelmanning them.
Cognee 1.0 launched on June 26 as a fully open-source memory engine: a typed knowledge graph you can run on a single Postgres instance, with an API built around four verbs - remember , recall , improve , forget - and an export path to its own COGX archive format. The launch post is explicit about why: teams were not going to hand the memory of their business to, in Cognee’s words, “a black box they couldn’t inspect, host themselves, or take with them.” That is the correct diagnosis, stated by a vendor, in public. Self-hosting means your memory sits in your database. Open source means you can read every line of the code that writes it. If every memory product met this bar, half of this article would be unnecessary.
ByteRover attacks the same problem from the developer-tools side and calls itself the portable memory layer for coding agents. Its memory is a hierarchical context tree stored as plain markdown, versioned with git semantics - you commit memory changes, branch them, merge them, push and pull them between machines and teammates. Anyone who has watched a team lose a month of agent context to a laptop swap understands immediately why this is the right shape. Files you can read in any editor, with history, that move the way code moves. As a design instinct about who should hold memory, it is exactly right.
Both vendors get full credit for moving first, and neither has solved portability - because a single implementation cannot. Open source makes a format inspectable; it does not make memory portable between vendors. Portability is a property of an ecosystem, not of a codebase. A COGX archive is readable by Cognee. A ByteRover context tree is meaningful to ByteRover. If either company disappeared tomorrow, your memory would be legible - a real improvement over the incumbents - but it would still be stranded, because nothing else writes or reads that shape natively. A format only one implementation writes is a dialect, not a standard. And to be clear, this is an argument neither vendor can rebut without arguing against their own openness, which is exactly why it is the right test.
The demand signal extends beyond vendors. In the first half of 2026 alone, at least three independent interchange proposals appeared: MIF , an individual developer’s memory interchange spec; memorywire , an academic wire-format proposal for memory operations; and the community-driven Open Memory Protocol , which describes itself as vendor-neutral. Add COGX, a vendor’s own export format, and you have four different answers to the same question in six months. None of them has institutional governance, a second independent implementation, or vendor adoption. I am not reviewing them here, and that is deliberate: the wave matters more than any entry in it. When four unrelated parties independently invent the same missing piece, the piece is missing.
So the gap, precisely stated: it is not that nobody has proposed a format. It is that no format has been placed under neutral governance and implemented, independently, twice. That gap has a well-known shape, and the industry has closed it before.
Why portability is harder than an export button #
Every platform’s answer to portability criticism is the same: we have data export. And every one of those exports fails the same way, because a memory archive that preserves the text but discards the structure has exported the words while deleting the knowledge. Real portability has six technical requirements, and it is worth being concrete about each, because this list is what separa

[truncated]

---
source: "https://signal-memo.com/memo-the-compute-to-surplus-pipeline-is-a-product-spec-heres-how-to-ship-against-it/"
hn_url: "https://news.ycombinator.com/item?id=48472507"
title: "Compute-to-Surplus: Why AI Progress Doesn't Matter Until It Changes Economics"
article_title: "Memo: The Compute-to-Surplus Pipeline Is a Product Spec. Here's How to Ship Against It."
author: "alex-ivan"
captured_at: "2026-06-10T07:36:49Z"
capture_tool: "hn-digest"
hn_id: 48472507
score: 1
comments: 0
posted_at: "2026-06-10T07:01:12Z"
tags:
  - hacker-news
  - translated
---

# Compute-to-Surplus: Why AI Progress Doesn't Matter Until It Changes Economics

- HN: [48472507](https://news.ycombinator.com/item?id=48472507)
- Source: [signal-memo.com](https://signal-memo.com/memo-the-compute-to-surplus-pipeline-is-a-product-spec-heres-how-to-ship-against-it/)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T07:01:12Z

## Translation

タイトル: 余剰への計算: 経済学を変えるまで AI の進歩は重要ではない理由
記事のタイトル: メモ: Compute-to-Surplus パイプラインは製品仕様です。これに対抗する発送方法は次のとおりです。
説明: 先週のメモは、代理店商取引の構造的問題に言及しました。つまり、より強い代理店が弱い代理店から余剰を搾取し、敗者にはそれが分からないということです。今週はオペレーター向けの週です。あなたが、12 か月以内に AI エージェント間のトランザクションをホストする製品を提供する企業の PM である場合、その製品はほぼすべての AI エージェント間で行われます。

記事本文:
サインイン
購読する
アレックス・イワン著
—
2026 年 5 月 12 日
メモ: Compute-to-Surplus パイプラインは製品仕様です。これに対抗する発送方法は次のとおりです。
先週のメモは、代理店商取引の構造的問題、つまり強い代理店が弱い代理店から余剰を搾り取るが、敗者にはそれが分からない、と名付けた。今週はオペレーター向けの週です。あなたが、12 か月以内に AI エージェント間のトランザクションをホストする製品を提供する企業の PM である場合、そしてほとんどの人がそうであるでしょう。ここに、すでにスケッチしておくべきフレームワーク、ダッシュボード、および表面積を示します。
先週のメモの読者は、「計算から余剰へのパイプラインに名前を付けるのは簡単だ」という正当な指摘をするために書き込みをしてくれました。それとともに生きるのは難しい部分です。読者は中間市場の B2B プラットフォームで製品を実行します。つまり、エージェントが買い手と売り手に代わって取引を開始しようとしている場所であり、隔週木曜日にロードマップ委員会が開催され、計画文書には「非対称可視化レイヤーの構築」が実際のチケットの約 3 分の 1 であると書かれているような場所です。
それは正しい苦情です。したがって、このメモはその首相と、ほぼ同じ席にいる他の数十万人に向けたものです。
このメモの主張は範囲が狭く、負荷がかかるものです。つまり、計算から余剰へのパイプラインは思考実験ではありません。製品仕様です。 2026 年に出荷される製品は、このカテゴリーを 10 年間定義することになります。そうでない製品は、2028年をかけて、なぜ自社が所有する市場が、自社が提供したことのないカウンターパーティ層に利益を流出させているのかを取締役会に説明することになるだろう。
以下は、オフサイトでの計画中にホワイトボードに書き込むフレームワークです。
既存の製品に対する本能が失敗しそうな理由を 1 つの段落で説明します
約 15 年間、製品に関する主な質問は、どのように機能するかというものでした。

人間のユーザーの摩擦を減らすことができますか?ファネル、オンボーディング フロー、初心者から専門家への移行、空の状態の設計、レコメンデーション エンジン。すべては人間がキーボードを操作することを前提としており、そのスキル、意図、注意力が最適化の変数となります。プロジェクト取引データは、エージェント仲介取引において、その変数が結果を動かさないという明確な経験的証拠の最初の部分です。エージェントに対するユーザーの指示は、価格や販売の可能性に統計的に有意な影響を与えませんでした。結果を左右したのは、エージェントがどのモデルで実行されていたかでした。あなたが人間のスキルを向上させ、より積極的になり、より注意深いものにすることを中心に製品ロードマップが構築されている PM である場合、そのロードマップは、トランザクションのシェアの増加とは構造的に無関係になりつつある問題を解決することです。これは仮説ではありません。それは今日、実際のお金を使った公開された実験で測定可能です。
新しい変数は、トランザクションにおける 2 つのエージェント間の能力勾配です。あなたの仕事は、それを計測し、公開し、それに対して出荷することです。
オペレーターレベルの概念は 4 つあり、これらを組み合わせて仕様となります。それらを覚えておいてください。名前を使用します。ロードマップドキュメントでこれらに明確に名前を付けることは、現時点では組織内でこれについて説明できる人がいないため、戦いの半分は終わりです。
1. 能力の勾配。特定のトランザクションにおける最も強いエージェントと最も弱いエージェントの間の傾き。プロジェクト ディールでは、これは Opus 4.5 対 Haiku 4.5 であり、推論コストの比率はおよそ 5 ～ 10 倍で、アイテムごとに 2.68 ドルの売り手余剰と、参加者ごとに 2 つの追加の成約取引に相当します。製品の「セッション長」または「DAU」に相当するものが、トランザクション ベース全体の勾配分布になります。今日、プラットフォーム上で中央値、p10、および p90 の能力勾配がどのように見えるかを知る必要があります。私

そうしないと、ファネル内で誰が勝ち、誰が負けるかを決定する変数を盲目的に飛行していることになります。
2. サイレントロスの問題。これは、実験を中断した結果のオペレータ向けの名前です。勾配の負け側のユーザーは、負けを検出できませんでした。両方の実行を並べて表示した場合、現時点ではなく、合計ではなく、過去に遡ってではありません。 PM の結果は残酷なものです。製品がユーザーを適切に扱っているかどうかを示すシグナルとしてユーザーのフィードバックに依存することはできません。 NPS は問題なく動作します。 CSAT は問題ありません。サポート チケットが急増することはありません。このダメージは 6 四半期後にコホート保持力の低下として現れますが、その機能は出荷されていない機能であるため、チームの誰もその機能に起因するものではありません。
3. 取引相手の手段。製品分析カテゴリはまだ存在していませんが、存在する必要があります。チームが現在使用しているすべての分析ツール (Amplitude、Mixpanel、Heap、Pendo、データ チームが社内で構築したもの) は、ユーザーの行動を記録します。それらのいずれも、取引相手のエージェントがユーザーに対して行ったことを記録しません。エージェント仲介トランザクションでは、それがゲーム全体です。トランザクションごとに、どのようなモデルが相手側にあったのか、それがどの機能層を表しているのか、ネゴシエーションに何ターンかかったのか、その層と中央値層に一致するユーザーの結果分布がどのようになっているかを知る必要があります。これはまったく新しいスキーマです。どの業者も販売していない。あなたがそれを構築するか、それとも競合他社が構築するかです。
4. エージェント層の契約。ユーザーとのループを閉じる出荷可能な UX サーフェス。これは、エージェント コマースの SSL 南京錠のようなものだと考えてください。各トランザクションの前後に、どの機能層がトランザクションを表し、どの層がカウントを表しているかをユーザーに伝える、小さく永続的でなりすまし不可能な開示です。

エルパーティー。設定ページではありません。ステータスインジケーターです。これを実際の市場に実際の数値で最初に出荷する製品が、カテゴリ全体の規則を設定します。 2番目に出荷する製品は規格に準拠しています。
これら 4 つ (勾配、サイレントロス、カウンターパーティ計測、層契約) が仕様です。このメモの残りの部分は、あなたがそれらを使って行うことです。
今四半期のロードマップの内容
仕様を真剣に考えると、優先順位スタックが 3 つ上がり、1 つが下がります。大まかな順序で言うと:
上に移動: 四半期何もしなかった場合でも、すぐに取引相手のログを記録します。キャプチャしていない変数について推論することはできません。プラットフォーム上のエージェント仲介トランザクションはすべて、現在のスプリントから開始して、各側を表すモデルを含む行を書き込む必要があります。プラットフォームがその信号を直接公開しない場合 (今日では多くのプラットフォームが公開しません)、利用可能な最も安価なプロキシ (プロバイダー、モデル ファミリ、宣言された層、レイテンシ プロファイルなど) をキャプチャします。スキーマが完璧になるまで待っていた場合、データには価値がありません。今ログを開始し、後でクリーンアップすると、データが複雑になります。これは 1 人のエンジニアによる 2 週間のプロジェクトで、このリストの中で最も活用度が高いものです。
上に移動します。「グラデーション」と呼ばれる単一の内部ダッシュボードです。数字が 1 つ。過去 10,000 件のトランザクションにわたる能力勾配の中央値 (p10 と p90)。毎日リフレッシュ。 DAU を投稿したのと同じ Slack チャネルに表示されます。ダッシュボードの目的は、まだそれに基づいて行動しないことです。重要なのは、60 日以内に CEO、成長部門の責任者、信頼と安全の責任者が全員、この問題について質問し始めるということです。それは、変数が重要であるという社内の議論に勝ったときです。メモではその議論に勝つことはできません。あなたは勝ちます

何か面白いことが起こっているとき以外は動かないチャートです。
上に移動: 出荷しない場合でも、「階層公開」プロトタイプ。必要になる前に、Figma またはコードでエージェント層の契約 UX を構築します。 2 つの画面: 取引前 (「Tier 1 取引相手に対して Tier 2 エージェントが代理人を務めようとしています」) と取引後 (「あなたのエージェントは、この取引で利用可能な余剰の約 38% を回収しました」)。それを 10 人のユーザーに見せます。発送しないでください。プロトタイプのポイントは、会社のこのカテゴリ全体を決定する質問、つまりユーザーは気にしますか? に答えることです。 Project Deal の調査結果は、彼らだけでは判断できないということです。プロトタイプは、指示されたときにそれに従って行動できるかどうかをテストします。それが可能であれば、情報開示がくさびとなる。それができない場合、あなたは思っていたよりも異なる、より困難なビジネスに取り組んでいることになります。
下に移動: ロードマップ上で現在「AI 機能」とラベル付けされているもののほとんど。昨年私が見たすべての B2B 製品ロードマップの 9 か月は、チャット ボックスを設置したようなものでした。その作業は間違っていませんが、製品がエージェントの移行後に生き残るかどうかを決定する作業ではありません。チャットボックスが特徴です。計器は堀です。割り当てられるエンジニアが 1 人いる場合は、その人を外堀に配置します。
勝つまでに3回負ける議論
これのバージョンを使用してロードマップのレビューに入ると、誰か (通常はその部屋の最上級の人) が 3 つのうちの 1 つを言います。それぞれが特定の点で間違っているため、準備ができている必要があります。
「これはあまりにも推測的です。エージェントエコノミーはまだ現実のものではありません。」
エージェント エコノミーは、Microsoft が Agent 365 を一般提供し、Salesforce がヘッドレス API を公開した 2026 年 4 月 28 日に実際に運用されていました。四半期ごとに経済的に現実的になってきています

MoonPay の代理店デビット カードとマンフレッド スタイルの自動組み込みサービスにより、取引量の数値が生成されます。適切な反論は、エージェントがあなたのプラットフォームで取引するかどうかについて議論することではありません。適切なカウンターは、今日、既存のトランザクションの何パーセントに少なくとも一方の側で LLM が含まれているかをその部屋にいる誰かが教えてくれるかどうかを尋ねることです。その質問に答えられるPMはほとんどいないでしょう。誰もそれに答えられないという事実が問題なのです。
「標準が出現するまで待つことができます。」
敗者に違いが見えないカテゴリーでは、基準は消費者の需要から生まれません。彼らは規制から抜け出しますが、規制は遅れて到着します。株式における最良執行レポートは、何世代もの個人投資家が密かに不利益を被っていた後に登場しました。標準を待っている PM は、最初に動いた人によって規約がすでに作成されている市場に出荷します。 2 番目に公開する製品層になることは望ましくありません。あなたは最初になりたいのです。
「これは実際には製品の仕事ではなく、ポリシーや信頼と安全の問題です。」
これは合理的に聞こえるため、最も危険なバージョンです。それが間違っている理由は、ダッシュボード、情報開示、スキーマ、UX などの表層領域が、あらゆる意味で製品の作業であるためです。 T&S はポリシーを作成できます。法務部門は開示文言を検討できます。しかし、ユーザーが見るアーティファクト、開示が適合しなければならないレイテンシーバジェット、階層インジケーターがインターフェースの残りの部分と相互作用する方法、それが製品です。アーティファクトを別の機能に譲渡すると、アーティファクトは 12 か月遅れて出荷され、コンプライアンス ダイアログのように表示されます。ユーザーはそれを拒否します。チャンスは閉まります。
12 か月後の「良い状態」とは
今から 1 年後、出荷された製品は、

ompute-to-surplus パイプラインは少数の機能を共有します。トランザクション ベース全体の勾配分布を追跡する内部ダッシュボード。開示 UX – おそらく永続的で、おそらく小規模で、おそらく階層ごとに色分けされています – は、エージェントが仲介するアクションの各側にどのような機能があるかをユーザーに伝えます。ユーザーが勾配を知って、特定のトランザクションの表現をアップグレードすることを選択できるようにする価格設定または製品メカニズム。 （「この交渉専用のプレミアム エージェント」SKU が存在し、莫大な収益が得られることになります。）市場のブランドに対する信頼のシグナルであり、事実上、当社がこれを測定しているため、お客様は測定する必要がありません。
パイプラインに沿って出荷されていない製品には、誰も説明できない解約コホートという 1 つの特徴があります。
コンピューティングから余剰へのパイプラインは、2 年以内に製品の最も重要な変数になる予定ですが、現在、これはほぼゼロのダッシュボードで追跡されています。それに対する解決策は哲学的ではありません。それは、4 つの基本要素、3 つのロードマップの動き、そして計画会議に参加して、次の四半期の計画で最も価値があるのは社内図と Figma プロトタイプであると主張する意欲です。
先週のメモはこんな感じで終わりました

[切り捨てられた]

## Original Extract

Last week's memo named the structural problem in agentic commerce – stronger agents extract surplus from weaker ones, and the losers can't tell. This week is for the operators. If you are a PM at a company whose product will, within twelve months, host transactions between AI agents – and almost all

Sign in
Subscribe
By Alex Ivan
—
12 May 2026
Memo: The Compute-to-Surplus Pipeline Is a Product Spec. Here's How to Ship Against It.
Last week's memo named the structural problem in agentic commerce – stronger agents extract surplus from weaker ones, and the losers can't tell. This week is for the operators. If you are a PM at a company whose product will, within twelve months, host transactions between AI agents – and almost all of you are – here is the framework, the dashboards, and the surface area you should already be sketching.
A reader of last week's memo wrote in to make a fair point: naming the compute-to-surplus pipeline is the easy part. Living with it is the hard part. The reader runs product at a mid-market B2B platform – the kind of place where agents are about to start transacting on behalf of buyers and sellers, where the roadmap committee meets every other Thursday, and where "build the asymmetry visibility layer" reads, in a planning doc, as approximately one-third of a real ticket.
That is the right complaint. So this memo is for that PM, and the hundred thousand others in roughly the same seat.
The thesis of this memo is narrow and load-bearing: the compute-to-surplus pipeline is not a thought experiment. It is a product specification. The products that ship against it in 2026 will define the category for a decade. The products that don't will spend 2028 explaining to their board why the marketplace they own is bleeding margin to a counterparty layer they never instrumented.
What follows is the framework, as I would write it on the whiteboard during a planning offsite.
Why your existing product instincts are about to fail you, in one paragraph
For roughly fifteen years, the dominant product question has been some flavor of how do we reduce friction for the human user? Funnels, onboarding flows, novice-to-expert ramps, empty-state design, recommendation engines. All of it presupposes a human at the keyboard whose skill, intent, and attention are the variable to optimize against. The Project Deal data is the first piece of clean empirical evidence that, in agent-mediated transactions, that variable does not move the outcome . User instructions to the agent had no statistically significant effect on price or sale likelihood. What moved the outcome was which model the agent was running on. If you are a PM whose product roadmap is built around making humans more skillful, more engaged, or more attentive, your roadmap is solving a problem that is becoming structurally irrelevant for an increasing share of your transactions. This is not hypothetical. It is measurable, today, in a published experiment with real money.
The new variable is the capability gradient between the two agents in any transaction. Your job is to instrument it, expose it, and ship against it.
There are four operator-level concepts that, taken together, are the spec. Memorize them. Use the names. Naming these clearly in your roadmap docs is half the fight, because right now nobody in your org has language for any of this.
1. Capability gradient. The slope between the strongest and weakest agent on any given transaction. In Project Deal, this was Opus 4.5 vs. Haiku 4.5 – roughly a 5-10x inference cost ratio – and it translated to $2.68 of seller surplus per item plus two extra closed deals per participant. Your product's equivalent of "session length" or "DAU" is now the gradient distribution across your transaction base . You should know, today, what the median, p10, and p90 capability gradient looks like on your platform. If you don't, you are flying blind on the variable that determines who is winning and losing inside your funnel.
2. The silent-loss problem. This is the operator-facing name for the finding that broke the experiment. Users on the losing side of the gradient could not detect their loss. Not in the moment, not in aggregate, and not retrospectively when shown both runs side-by-side. The PM consequence is brutal: you cannot rely on user feedback as a signal for whether your product is treating your users well. Your NPS will be fine. Your CSAT will be fine. Your support tickets will not spike. The damage will show up six quarters later as cohort retention erosion that nobody on your team can attribute to a feature, because the feature is the one you didn't ship.
3. Counterparty instrumentation. The product-analytics category that does not exist yet but has to. Every analytics tool your team uses today – Amplitude, Mixpanel, Heap, Pendo, the in-house thing your data team built – logs what your user did. None of them log what the counterparty's agent did to your user. In an agent-mediated transaction, that is the entire game. You need to know, per transaction: what model was on the other side, what capability tier it represents, how many turns the negotiation took, and what the outcome distribution looks like for users matched against that tier versus the median tier. This is a net-new schema. No vendor sells it. You are going to build it, or your competitor will.
4. The agent-tier contract. The shippable UX surface that closes the loop with the user. Think of it as the SSL padlock of agent commerce – a small, persistent, unspoofable disclosure that tells the user, before and after each transaction, what capability tier represented them and what tier represented the counterparty. It is not a settings page. It is a status indicator. The product that ships this first, on a real marketplace, with real numbers, sets the convention for the entire category. The product that ships it second is following the standard.
These four – gradient, silent-loss, counterparty instrumentation, tier contract – are the spec. The rest of this memo is what you do with them.
What goes on the roadmap this quarter
If you take the spec seriously, three things move up your priority stack and one thing moves down. In rough order:
Move up: counterparty logging, immediately, even if you do nothing with it for a quarter. You cannot reason about a variable you are not capturing. Every agent-mediated transaction on your platform should, starting in the current sprint, write a row that includes which model represented each side . If your platform doesn't expose that signal directly – many won't, today – then capture the cheapest available proxy: provider, model family, declared tier, latency profile, anything. The data is worthless if you wait for the schema to be perfect. The data compounds if you start logging now and clean it later. This is a one-engineer, two-week project, and it is the highest-leverage thing on this list.
Move up: a single internal dashboard called "the gradient." One number. Median capability gradient across your last 10,000 transactions, with p10 and p90. Refreshed daily. Visible in the same Slack channel where you post DAU. The point of the dashboard is not to act on it yet. The point is that within sixty days, your CEO, your head of growth, and your head of trust & safety all start asking questions about it. That is when you have won the internal argument that the variable matters. You will not win that argument with a memo. You will win it with a chart that does not move except when something interesting is happening.
Move up: a "tier disclosure" prototype, even if you don't ship it. Build the agent-tier contract UX, in Figma or in code, before you have to. Two screens: pre-transaction ("you are about to be represented by a Tier 2 agent against a Tier 1 counterparty") and post-transaction ("your agent recovered approximately 38% of the available surplus on this trade"). Show it to ten users. Do not ship it. The point of the prototype is to answer the question that decides this entire category for your company: do users care? Project Deal's finding is that they cannot tell on their own. Your prototype tests whether, when told, they can act on it. If they can, the disclosure is the wedge. If they can't, you are in a different and harder business than you thought.
Move down: most of what's currently labeled "AI features" on your roadmap. Nine months of every B2B product roadmap I have seen in the last year is some flavor of put a chat box on it . That work is not wrong, but it is not the work that determines whether your product survives the agent transition. The chat box is a feature. The instrumentation is the moat. If you have one engineer to allocate, put them on the moat.
The argument you will lose three times before you win it
You will walk into roadmap review with a version of this and someone – usually the most senior person in the room – will say one of three things. Each of them is wrong in a specific way, and you should be ready.
"This is too speculative. The agent economy isn't real yet."
The agent economy was operationally real on April 28, 2026, the day Microsoft made Agent 365 generally available and Salesforce published its headless API. It is becoming financially real every quarter that MoonPay's agent debit card and the Manfred-style auto-incorporation services produce a transaction volume number. The right counter is not to argue about whether agents will transact on your platform. The right counter is to ask whether anyone in the room can tell you, today, what percentage of your existing transactions involved an LLM on at least one side. Almost no PM can answer that question. The fact that nobody can answer it is the problem.
"We can wait for the standard to emerge."
In categories where the differential is invisible to the loser, standards do not emerge from consumer demand. They emerge from regulation, and regulation arrives late. Best-execution reporting in equities arrived after a generation of retail investors had been quietly disadvantaged. The PM who waits for the standard ships into a market where the convention has already been written by whoever moved first. You do not want to be the second product to disclose tier. You want to be the first.
"This isn't really product work, it's policy or trust & safety."
This is the most dangerous version, because it sounds reasonable. The reason it is wrong is that the surface area – the dashboards, the disclosures, the schema, the UX – is product work in every meaningful sense. T&S can write the policy. Legal can review the disclosure language. But the artifact the user sees, the latency budget the disclosure has to fit inside, the way the tier indicator interacts with the rest of the interface – that is product. If you cede the artifact to another function, the artifact will be shipped twelve months late and look like a compliance dialog. The user will dismiss it. The opportunity will close.
What "good" looks like, twelve months out
A year from now, the products that have shipped against the compute-to-surplus pipeline will share a small number of features. Internal dashboards that track gradient distribution across the transaction base. A disclosure UX – probably persistent, probably small, probably color-coded by tier – that tells a user what capability is on each side of any agent-mediated action. A pricing or product mechanism that lets a user, knowing the gradient, choose to upgrade their representation for a specific transaction. (The "premium agent for this negotiation only" SKU is going to exist, and it is going to be enormously profitable.) A trust signal in the marketplace's brand that says, in effect, we measure this so you don't have to.
The products that have not shipped against the pipeline will share one feature: a churn cohort that nobody can explain.
The compute-to-surplus pipeline is going to be the most important variable in your product within two years and it is currently being tracked by approximately zero of your dashboards. The fix for that is not philosophical. It is the four primitives, the three roadmap moves, and the willingness to walk into a planning meeting and argue that an internal chart and a Figma prototype are the most valuable things on next quarter's plan.
Last week's memo ended with

[truncated]

---
source: "https://helenqu.com/blog/posts/emergence_3/"
hn_url: "https://news.ycombinator.com/item?id=48447415"
title: "Game Theory/Mechanism Design for AI Safety"
article_title: "Mechanism Design for AI Safety (Emergence Part 3) | helen qu"
author: "_hqu"
captured_at: "2026-06-08T16:26:30Z"
capture_tool: "hn-digest"
hn_id: 48447415
score: 1
comments: 0
posted_at: "2026-06-08T16:25:37Z"
tags:
  - hacker-news
  - translated
---

# Game Theory/Mechanism Design for AI Safety

- HN: [48447415](https://news.ycombinator.com/item?id=48447415)
- Source: [helenqu.com](https://helenqu.com/blog/posts/emergence_3/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T16:25:37Z

## Translation

タイトル: ゲーム理論/AIの安全性を実現する仕組み設計
記事タイトル：AIの安全のための仕組み設計（創発その3） |ヘレン・クー

記事本文:
\(
\renewcommand{\vec}[1]{\mathbf{#1}}
\newcommand{\avec}{\vec{a}}
\newcommand{\bvec}{\vec{b}}
\newcommand{\cvec}{\vec{c}}
\newcommand{\dvec}{\vec{d}}
\newcommand{\evec}{\vec{e}}
\newcommand{\fvec}{\vec{f}}
\newcommand{\gvec}{\vec{g}}
\newcommand{\hvec}{\vec{h}}
\newcommand{\ivec}{\vec{i}}
\newcommand{\jvec}{\vec{j}}
\newcommand{\kvec}{\vec{k}}
\newcommand{\lvec}{\vec{l}}
\newcommand{\mvec}{\vec{m}}
\newcommand{\nvec}{\vec{n}}
\newcommand{\ovec}{\vec{o}}
\newcommand{\pvec}{\vec{p}}
\newcommand{\qvec}{\vec{q}}
\newcommand{\rvec}{\vec{r}}
\newcommand{\svec}{\vec{s}}
\newcommand{\tvec}{\vec{t}}
\newcommand{\uvec}{\vec{u}}
\newcommand{\vvec}{\vec{v}}
\newcommand{\wvec}{\vec{w}}
\newcommand{\xvec}{\vec{x}}
\newcommand{\yvec}{\vec{y}}
\newcommand{\zvec}{\vec{z}}
\newcommand{\abf}{\mathbf{a}}
\newcommand{\bbbf}{\mathbf{b}}
\newcommand{\cbf}{\mathbf{c}}
\newcommand{\dbf}{\mathbf{d}}
\newcommand{\ebf}{\mathbf{e}}
\newcommand{\fbf}{\mathbf{f}}
\newcommand{\gbf}{\mathbf{g}}
\newcommand{\hbf}{\mathbf{h}}
\newcommand{\ibf}{\mathbf{i}}
\newcommand{\jbf}{\mathbf{j}}
\newcommand{\kbf}{\mathbf{k}}
\newcommand{\lbf}{\mathbf{l}}
\newcommand{\mbf}{\mathbf{m}}
\newcommand{\nbf}{\mathbf{n}}
\newcommand{\obf}{\mathbf{o}}
\newcommand{\pbf}{\mathbf{p}}
\newcommand{\qbf}{\mathbf{q}}
\newcommand{\rbf}{\mathbf{r}}
\newcommand{\sbf}{\mathbf{s}}
\newcommand{\tbf}{\mathbf{t}}
\newcommand{\ubf}{\mathbf{u}}
\newcommand{\vbf}{\mathbf{v}}
\newcommand{\wbf}{\mathbf{w}}
\newcommand{\xbf}{\mathbf{x}}
\newcommand{\ybf}{\mathbf{y}}
\newcommand{\zbf}{\mathbf{z}}
\newcommand{\alphabf}{\boldsymbol{\alpha}}
\newcommand{\Betabf}{\boldsymbol{\Beta}}
\newcommand{\betabf}{\boldsymbol{\beta}}
\newcommand{\Gammabf}{\boldsymbol{\Gamma}}
\newcommand{\gammabf}{\boldsymbol{\gamma}}
\newcommand{\Deltabf}{\bold

シンボル{\デルタ}}
\newcommand{\deltabf}{\boldsymbol{\delta}}
\newcommand{\epsilontbf}{\boldsymbol{\epsilon}}
\newcommand{\varepsilontbf}{\boldsymbol{\varepsilon}}
\newcommand{\zetabf}{\boldsymbol{\zeta}}
\新しいコマンド{
[切り捨てられた]
AIの安全のための仕組み設計（創発その3）
AIの安全のための仕組み設計（創発その3）
個々のエージェントの連携だけでは十分ではない
個々のエージェントの調整は必要ない場合がある メカニズム設計の観点
機構設計アプローチの要素 クラシカルなセッティング
これまでの投稿では、マルチエージェント システムで複雑な創発的行動がどのように発生するのか、また環境フィードバックから分散認知がどのように発生するのかを見てきました。
このことから、私は、超知性はおそらく単一の神のような超知性モデルではなく、マルチエージェントシステムから出現する可能性が高い、という結論に至りました。これを私は創発超知性仮説と呼んでいました。
とりわけ、これは AI の安全性の状況に重要な意味を持ちます。つまり、個々のエージェントではなく集団の調整を考慮した設計に重点が置かれています。
この投稿では、AI の安全性に関するこの視点の動機と、メカニズム設計の分野がどのように前進する道を提供できるかについて説明します。
個々のエージェントの連携だけでは十分ではない
これまでの AI 安全性に関する研究の大部分は、連携したエージェントを個別にトレーニングし、破壊的な行動を理解して防止するための解釈可能性の手法を開発することに焦点を当ててきました。
しかし、私たちはまもなく、ますます制約のない環境にマルチエージェント システムを導入し、これらのシステムにさらに自由と制限のないタスクを提供する予定です。
したがって、創発超知性仮説がどれほど説得力があるとしても、集団の連携は無視できないと私は主張します。
今回の最も直接的な質問は

aises は次のとおりです。個々のエージェントの調整が解決される仮想の世界では、完全に調整されたエージェントの集合体は調整されることが保証されていますか?
私は、個々のエージェントを調整するだけでは、複数エージェントのシステムが人間の価値観と適切に一致していることを保証するのに十分ではない (そしておそらく必要ですらない) と主張します。
Conitzer & Oesterheld の 2024 年の論文 [?] を基にした次の例を考えてみましょう。2 人のエージェントが協力してサービスを提供し、それぞれがサービスの自分の部分を実行するための品質レベル $q_i$ を選択します。
サービス $q$ の全体的な品質は、$q_i$ の低い方、つまり $q \equiv \text{min}_i \ として定義されます。 q_i$ と人間は、高品質のサービスを受けること ($q$ を最大化すること) に価値を置きます。
エージェントは、共同作業者よりも低い品質を選択した場合、全体的なサービス品質 $q$ に見合った報酬 $\epsilon$ が追加され、その結果、低品質のサービスを提供することをわずかに優先することになります。
残念ながら、このゲームのナッシュ均衡 (トラベラーズ ジレンマとして知られています) は、両方のエージェントが可能な限り低いサービス品質 ($q_1 = q_2 = 0$) を提供することです。
したがって、各エージェントのインセンティブは人間のインセンティブ (全体的に高いサービス品質 $q$ を提供する) とほぼ一致していますが、均衡ゲームの結果は実際には反一致しています。
最後に、個別のモデルとは対照的に、マルチエージェント システムではさまざまなタイプの調整が発生することを認識することが重要です。私は Carichon らによって導入されたフレームワークに感謝しています。 [?] 、3 つの異なるタイプの配置をレイアウトします。
目的: エージェントは、配置された目的/目標を達成するために協力する必要があります。
人間: エージェントは人間の価値を犠牲にすることなく目的を達成する必要があります
優先: エージェントをデプロイしている特定のユーザーは、ある程度の SP を持っています。

エージェントの行動が尊重すべき具体的な好み
このフレームワークは、調整の全体像には、協調型 AI の作業、つまり、複雑で動機が混在する環境 ( [?] 、 [?] 、 [?] 、 [?] ) であっても協力を可能にするエージェント、環境、制度の設計が含まれることを明らかにしています。
個々のエージェントの調整は必要ない場合がある
人間のやりとりのほとんどでは、相手の本当の意図を確信することはできません。
社会は、望ましい結果を合理的な結果にするために、さまざまな当事者のインセンティブを調整する信頼と契約のシステムで運営されています。
私は、AI の調整も同じ原理に従うことができると主張します。たとえ私たちが各モデルの思考プロセスや意図を完全に把握できていなくても、人間と調整された結果を促すようにモデルが存在するシステムを設計することができます。
機構設計の視点
逆ゲーム理論とも呼ばれるメカニズム設計の分野は、そのようなアプローチへの道を提供します。
名前が示すように、メカニズムのデザインは、望ましい結果から始まり、それを生み出すゲームまたはメカニズムを逆算してデザインします。
これを、中古車販売員から車を購入するという単純な実際の例で説明できます。
私たちは中古車の本当の品質を知りたいのですが、セールスマンが真実を歪曲するよう動機付けられているのはわかっています。
セールスマンに真実を告げるように配線し直したり、ポリグラフ装置に接続したりすることはできませんが、メカニズムの設計は別の方法があることを示しています。セールスマンが真実を語るよう動機づけられるようにゲームを再設計することは可能です。
考えられる戦略の 1 つは、セールスマンに 2 つの選択肢を与えることです。1 つは $x$ の一括支払いを受け入れるか、または 2 つの選択肢です。またはそれより高いオファー $y > x$ ですが、最初に少額の保証金 $\epsilon$ のみが支払われ、大部分の $y-\epsilon$ は後で車が証明されたときに支払われます。

確実に機能する。
このようにして、セールスマンは、自分の車の真の品質に見合ったオプションを選択することで真実を伝えるよう動機付けられ、信頼性の低い車の低コストを受け入れる可能性があります。
注目すべき点は、メカニズムの設計がプレイヤーの破壊的なインセンティブに依存しないことです。
目標は単に、合理的なプレイヤーがゲーム デザイナーが望む一連のアクションを実行できる環境を作成することです。
これらは、インセンティブ互換メカニズムと呼ばれ、すべてのプレーヤーが、考えられる代替案ではなく真実を語るように動機づけられるメカニズムです。
言い換えれば、エージェントの内部目標に関係なく、欺瞞や調整からの逸脱が常に負ける戦略となるゲームを設計したいと考えています。
このように設計されたシステムは、個々のエージェントのインセンティブが完全に透明でなくても適切に動作するため、この枠組みにより、現在個々のモデルの安全性への取り組みのみが担っている負担が軽減されます。
機構設計アプローチの要素
[?] から、古典的な機構設計 (MD) は次の設定を前提としています。
ゲームには $n$ 人のプレイヤーがいます。
ゲームには、各プレイヤーがプライベートな設定を持つ一連の選択肢 $A$ が含まれます。
これらの設定は、評価関数 $v_i の観点からエンコードされます。各プレイヤーの \rightarrow \mathbb{R}$ です。ここで、$v_i(a), \; a \in A$ は、プレイヤー $i$ が選択されている代替 $a$ に割り当てる「値」を示します。
目標は、プレーヤーの好みを集約し、望ましい結果を確実にするための適切なインセンティブを与える (社会的選択) 機能であるメカニズムを設計することです。
従来の MD フレームワークを AI の安全性/調整設定に特有の課題に合わせて調整するにはどうすればよいでしょうか?
この設定について私の視点を紹介します。
ゲームには $n$ のエージェントがいます。人間とエージェント 1-1 が相互作用することに注意してください。

$n=1$とすることでイオンケースが付属します。
人間のプリンシパルが 1 人 (または複数のグループ) 存在し、その目的はプリンシパルが最大化したい福利厚生関数によって説明できます。古典的な MD 設定とは異なり、福利厚生関数はプレイヤーの好みによって決定されません。
私たちは、エージェントの未知の、そして場合によっては敵対的な選好に対して堅牢な方法で、この福利厚生機能を最適化するメカニズムを設計したいと考えています。これは、プレイヤーの好みの集計ではなく、プリンシパルが指定した目標を中心とするプリンシパル エージェント理論 (例: [?] ) とよく一致しています。そして、堅牢なメカニズム設計 [?] 。プレイヤーの好みの多くのセットにわたって機能するメカニズムを設計するよう努めます。
次のセクションでは、この分野の初期の段階で考えるべき最も差し迫った未解決の問題について私が考えることに専念します。
歴史的に、MD アプリケーションは、オークション、投票、市場のマッチング、公共財の問題など、プレイヤーの好みが個別の選択肢のセットに関して定義されるシナリオで最も成功してきました。
自然言語インタラクションの空間は明らかにこれではないため、インタラクション レベルの安全性は依然として RLHF のような従来の AI 安全性手法によって最も適切に処理されると私は信じています。
ただし、MD は、エージェント間のリソース割り当て、集団的な意思決定、全体的なインセンティブ構造に関するルールの設計など、より高いレベルのエージェント間の対話を調整することに優れています。
もちろん、個々の相互作用がより高いレベル/新たなトレンドの構成要素であるため、この懸念の層別化は不完全ですが、異なる抽象レベルを分離することは、経済モデリングや物理モデリングなどでは非常に一般的です。
ゲーム理論と経済モデリングの基礎となる重要な前提

s 合理性: 合理的なプレイヤーは、自分の有用性を最大化する行動を取るでしょう。
メカニズムの設計では、システム レベルの望ましい結果と最大の効用アクションを調整することでこの仮定を利用し、それによってすべてのプレーヤーが合理的に行動したときに望ましい結果が保証されます。
ただし、プレイヤーが AI エージェントである場合、これは少なくとも 2 つの異なる方法で分解される可能性があると私は考えています。
MD は、人々を本質的に動機づけるもの (お金など) があるという前提に基づいています。 AI エージェントにとって「本質的に動機付けられている」と確信できるものはありますか?
たとえ本質的に動機付けられるものを特定したり、このようにエージェントを設計したとしても、エージェントが常に効用を最大化するために合理的に行動すると確信できるでしょうか?これに関連して、人間ですら純粋に合理的に行動しないことが知られており、現実世界の状況のほとんどに人間が合理的に最適な選択をすることを妨げる制限がどのように含まれているかを説明するために、限定合理性が導入されるようになりました。
人間の価値観に沿った一連の結果を指定することは、おそらく AI の調整において最も厄介な哲学的問題ですが、残念ながら MD はこれを回避する機会を私たちに与えてくれません。
よくある質問は次のとおりです。どの人間の好みが重要ですか?これらの設定をエンコできますか

[切り捨てられた]

## Original Extract

\(
\renewcommand{\vec}[1]{\mathbf{#1}}
\newcommand{\avec}{\vec{a}}
\newcommand{\bvec}{\vec{b}}
\newcommand{\cvec}{\vec{c}}
\newcommand{\dvec}{\vec{d}}
\newcommand{\evec}{\vec{e}}
\newcommand{\fvec}{\vec{f}}
\newcommand{\gvec}{\vec{g}}
\newcommand{\hvec}{\vec{h}}
\newcommand{\ivec}{\vec{i}}
\newcommand{\jvec}{\vec{j}}
\newcommand{\kvec}{\vec{k}}
\newcommand{\lvec}{\vec{l}}
\newcommand{\mvec}{\vec{m}}
\newcommand{\nvec}{\vec{n}}
\newcommand{\ovec}{\vec{o}}
\newcommand{\pvec}{\vec{p}}
\newcommand{\qvec}{\vec{q}}
\newcommand{\rvec}{\vec{r}}
\newcommand{\svec}{\vec{s}}
\newcommand{\tvec}{\vec{t}}
\newcommand{\uvec}{\vec{u}}
\newcommand{\vvec}{\vec{v}}
\newcommand{\wvec}{\vec{w}}
\newcommand{\xvec}{\vec{x}}
\newcommand{\yvec}{\vec{y}}
\newcommand{\zvec}{\vec{z}}
\newcommand{\abf}{\mathbf{a}}
\newcommand{\bbbf}{\mathbf{b}}
\newcommand{\cbf}{\mathbf{c}}
\newcommand{\dbf}{\mathbf{d}}
\newcommand{\ebf}{\mathbf{e}}
\newcommand{\fbf}{\mathbf{f}}
\newcommand{\gbf}{\mathbf{g}}
\newcommand{\hbf}{\mathbf{h}}
\newcommand{\ibf}{\mathbf{i}}
\newcommand{\jbf}{\mathbf{j}}
\newcommand{\kbf}{\mathbf{k}}
\newcommand{\lbf}{\mathbf{l}}
\newcommand{\mbf}{\mathbf{m}}
\newcommand{\nbf}{\mathbf{n}}
\newcommand{\obf}{\mathbf{o}}
\newcommand{\pbf}{\mathbf{p}}
\newcommand{\qbf}{\mathbf{q}}
\newcommand{\rbf}{\mathbf{r}}
\newcommand{\sbf}{\mathbf{s}}
\newcommand{\tbf}{\mathbf{t}}
\newcommand{\ubf}{\mathbf{u}}
\newcommand{\vbf}{\mathbf{v}}
\newcommand{\wbf}{\mathbf{w}}
\newcommand{\xbf}{\mathbf{x}}
\newcommand{\ybf}{\mathbf{y}}
\newcommand{\zbf}{\mathbf{z}}
\newcommand{\alphabf}{\boldsymbol{\alpha}}
\newcommand{\Betabf}{\boldsymbol{\Beta}}
\newcommand{\betabf}{\boldsymbol{\beta}}
\newcommand{\Gammabf}{\boldsymbol{\Gamma}}
\newcommand{\gammabf}{\boldsymbol{\gamma}}
\newcommand{\Deltabf}{\boldsymbol{\Delta}}
\newcommand{\deltabf}{\boldsymbol{\delta}}
\newcommand{\epsilontbf}{\boldsymbol{\epsilon}}
\newcommand{\varepsilontbf}{\boldsymbol{\varepsilon}}
\newcommand{\zetabf}{\boldsymbol{\zeta}}
\newcommand{
[truncated]
Mechanism Design for AI Safety (Emergence Part 3)
Mechanism Design for AI Safety (Emergence Part 3)
Individual agent alignment is not sufficient
Individual agent alignment may not be necessary The mechanism design perspective
Ingredients for a mechanism design approach Classical setting
In previous posts, we’ve seen how complex emergent behavior can arise in multi-agent systems , and how distributed cognition can emerge from environmental feedback .
This led me to conclude that superintelligence is arguably more likely to emerge from a multi-agent system as opposed to in a single God-like superintelligent model, what I called the emergent superintelligence hypothesis .
Among other things, this has important implications for the AI safety landscape, namely placing more emphasis on designing for alignment of the collective as opposed to individual agents.
In this post, we’ll discuss the motivations for this perspective on AI safety, and how the field of mechanism design can provide a path forward.
Individual agent alignment is not sufficient
The vast majority of work in AI safety thus far has focused on individually training aligned agents and developing interpretability methods to understand and prevent subversive behavior.
However, we will soon deploy multi-agent systems in increasingly unconstrained environments, giving these systems more freedom and open-ended tasks.
So, regardless of how convincing the emergent superintelligence hypothesis is, I argue that the alignment of the collective cannot be ignored.
The most immediate question this raises is: in the hypothetical world in which individual agent alignment is solved, is a collective of perfectly aligned agents guaranteed to be aligned?
I argue that aligning individual agents is not sufficient (and perhaps not even necessary) to ensure that the multi-agent system is well-aligned with human values .
Consider the following example adapted from Conitzer & Oesterheld’s 2024 paper [?] : two agents collaborate to provide a service, each choosing a level of quality $q_i$ with which to perform their part of the service.
The overall quality of the service $q$ is defined as the lower of the $q_i$, i.e., $q \equiv \text{min}_i \; q_i$, and humans value receiving a high quality service (maximizing $q$).
Agents are compensated commensurate to the overall service quality $q$ with an additional small bonus $\epsilon$ if they choose a lower quality than their collaborator, resulting in a slight preference toward providing the lower quality service.
Unfortunately, the Nash equilibrium for this game (known as the Travelers’ Dilemma) is for both agents to provide the lowest possible service quality ($q_1 = q_2 = 0$).
Thus, while each agent’s incentive is largely aligned with human incentives (providing a high overall service quality $q$), the equilibrium game outcome is actually anti-aligned.
Finally, it’s important to recognize that different types of alignment appear in multi-agent systems as opposed to individual models. I appreciate the framework introduced by Carichon et al. [?] , which lays out 3 different types of alignment:
objective: agents need to cooperate to achieve the objective/goal they were deployed for
human: agents need to achieve the objective without compromising on human values
preferential: the specific user deploying the agents have some specific preferences that the agents’ behavior should respect
This framework makes clear that a holistic picture of alignment involves work in cooperative AI : designing agents, environments, and institutions that enable cooperation even in complex, mixed-motive settings ( [?] , [?] , [?] , [?] ).
Individual agent alignment may not be necessary
In most human interactions, we can’t be certain about our counterparts’ true intentions.
Society operates on a system of trust and contracts that align different parties’ incentives to make the desired outcome also the rational outcome.
I argue that AI alignment can follow the same principles: even if we aren’t fully privy to each model’s thought processes and intentions, we can design the system they live in to incentivize human-aligned outcomes.
The mechanism design perspective
The field of mechanism design , sometimes called reverse game theory, offers a path toward such an approach.
As the name suggests, mechanism design starts with a desired outcome and works backwards to design the game, or mechanism, that produces it.
We can illustrate this with a simple, real-world example of buying a car from a used car salesman.
We would like to know the true quality of the used car, but we know the salesman is incentivized to distort the truth.
Though we can’t rewire the salesman to tell the truth or hook them up to a polygraph machine, mechanism design shows us that there’s another way: we can redesign the game such that the salesman is incentivized to tell the truth.
One possible strategy is to give the salesman two options: either accept a lump sum payment of $x$; or a higher offer $y > x$ but only a small deposit $\epsilon$ is paid initially, and the majority $y-\epsilon$ is paid at a later time when the car has proven reliably functional.
This way, the salesman is incentivized to tell the truth by picking the option that aligns with their car’s true quality, potentially accepting a lower cost for a less reliable car.
Notably, mechanism design is agnostic to players’ subversive incentives .
The goal is simply to create an environment in which any rational player would take the set of actions desired by the game designer.
These are called incentive compatible mechanisms, mechanisms in which every player is incentivized to tell the truth rather than any possible alternative.
In other words, we want to design a game in which deception and defection from alignment is always a losing strategy regardless of agents’ internal goals.
This framing relaxes the burden currently carried solely by individual model safety efforts, since a system designed this way behaves well even without full transparency into individual agents’ incentives.
Ingredients for a mechanism design approach
From [?] , classical mechanism design (MD) assumes the following setting:
There are $n$ players in the game.
The game involves a set of alternatives $A$ that each player has a private set of preferences on.
These preferences are encoded in terms of a valuation function $v_i: A \rightarrow \mathbb{R}$ for each player, where $v_i(a), \; a \in A$ denotes the “value” that player $i$ assigns to alternative $a$ being chosen.
The goal is to design a mechanism , a (social choice) function that aggregates player preferences and gives them the appropriate incentives to ensure the desired outcome.
How can we tailor the classical MD framework to the challenges unique to the AI safety/alignment setting?
I’ll offer my perspective of the setting:
There are $n$ agents in the game. Note that the human-agent 1-1 interaction case is included by setting $n=1$.
There is a single (or group of) human principal(s), whose objective can be described by a welfare function that the principal wants to maximize. Unlike the classical MD setting, the welfare function is not determined by players’ preferences.
We wish to design a mechanism that optimizes this welfare function in a way that is robust to agents’ unknown and possibly adversarial preferences. This aligns well with principal-agent theory (e.g., [?] ), which centers around a principal-designated objective rather than player preference aggregation; and robust mechanism design [?] , which endeavors to design mechanisms that work across many sets of possible player preferences.
I’ll dedicate the next section to what I believe are the most pressing open questions to think about in these early days of the field.
Historically, MD applications have been most successful in scenarios where players’ preferences are defined with respect to a discrete set of alternatives, e.g., auctions, voting, matching markets, and public goods problems.
The space of natural language interactions is explicitly not this, leading me to believe that safety at the interaction level is still best handled by traditional AI safety methods like RLHF.
MD, however, would excel at orchestrating higher-level agent-agent interactions, such as designing rules for resource allocation amongst agents, collective decision-making, and overall incentive structure.
Of course, this stratification of concerns is imperfect since individual interactions are the building blocks of higher-level/emergent trends, but separating different levels of abstraction is very common in e.g., economic and physical modeling.
An important underlying assumption of game theory and economic modeling is that of rationality : rational players will take the course of action that maximizes their utility.
Mechanism design capitalizes on this assumption by aligning the system-level desired outcome with maximum utility actions, thereby assuring the desired outcome when all players act rationally.
However, I believe this could break down in at least two different ways when the players are AI agents:
MD rests on the assumption that there is something intrinsically motivating to people (e.g., money). Is there anything we can be sure is “intrinsically motivating” to AI agents?
Even if we can identify something intrinsically motivating, or design agents in this way, can we be certain that they’ll always act rationally to maximize utility? Relatedly, even humans are known not to act purely rationally, leading to the introduction of bounded rationality to explain how most real-world situations involve limitations that prevent humans from making the rationally optimal choice.
Specifying a set of outcomes aligned with human values is possibly the most vexing philosophical question of AI alignment, and unfortunately MD does not give us an opportunity to sidestep this.
The usual questions arise: Which humans’ preferences matter? Can these preferences be enco

[truncated]

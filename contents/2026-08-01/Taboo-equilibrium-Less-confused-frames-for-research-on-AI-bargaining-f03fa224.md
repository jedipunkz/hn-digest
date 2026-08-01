---
source: "https://www.lesswrong.com/posts/KDq5aXwanvH5YoZYs/taboo-equilibrium-less-confused-frames-for-research-on-ai"
hn_url: "https://news.ycombinator.com/item?id=49133262"
title: "Taboo \"equilibrium\": Less confused frames for research on AI bargaining"
article_title: "Taboo “equilibrium”: Less confused frames for research on AI bargaining — LessWrong"
author: "joozio"
captured_at: "2026-08-01T11:56:16Z"
capture_tool: "hn-digest"
hn_id: 49133262
score: 1
comments: 0
posted_at: "2026-08-01T11:02:26Z"
tags:
  - hacker-news
  - translated
---

# Taboo "equilibrium": Less confused frames for research on AI bargaining

- HN: [49133262](https://news.ycombinator.com/item?id=49133262)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/KDq5aXwanvH5YoZYs/taboo-equilibrium-less-confused-frames-for-research-on-ai)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T11:02:26Z

## Translation

タイトル: タブーな「均衡」: AI 交渉に関する研究のための混乱の少ない枠組み
記事タイトル: タブーな「均衡」: AI 交渉に関する研究のための混乱の少ない枠組み — LessWrong
説明: 強力な AI が競合に陥る理由と、それを軽減する方法を理解するには、交渉の問題、つまり複数の AI が競合する状況を理解する必要があります。

記事本文:
ログイン タブーの「均衡」: AI 交渉に関する研究のための混乱の少ないフレーム — 間違っていない安全なパレートの改善 ゲーム理論 オープンソース ゲーム理論 AI ワールド モデリング フロントページ 25
タブーな「均衡」: AI 交渉に関する研究のための混乱の少ない枠組み
強力な AI が競合に陥る理由と、それを軽減する方法を理解するには、交渉の問題、つまり複数のエージェントがパレート効率の結果に対して異なる好みを持っている状況を理解する必要があります。
私は、交渉問題に関する特定の一般的な枠組みが混乱しているのではないかと疑うようになりました。ここではその理由と、どのフレームが良いと思うかを説明します。この動機の 1 つは、他の人々が安全なパレート改善 (SPI) に関する研究を進めるのを手助けしたいということです。SPI は、AI 紛争のマイナス面を軽減するための最も有望なアプローチの 1 つであると、私の考えでは考えています。
私の以前の著作から関連する背景を説明します。 (さらに)
AI 交渉研究の重要な枠組みとして、交渉における要求調整の失敗の説明を実行例として使用して、メカニズムの説明を導入します (この例は、次の箇条書きの主張の動機にもなります)。 (さらに)
ナッシュ均衡で何が起こるかを問うのではなく、AI交渉人がお互いについてどのような信念をもっともらしく抱くかを問うべきだと主張する。 (さらに) そして
彼らは、エージェントがどのようなものを「搾取的」とみなすかについてのヒューリスティックに大きく依存するのではなく、どのようなコミットメントが良い結果をもたらすと期待しているのかを注意深く説明する必要があると主張しています。 (さらに)
事前に以下の内容 (少なくともリンク先のセクション) を読むことをお勧めしますが、目的に不可欠な部分は以下の引用符内にあるため、厳密に必須というわけではありません。
「AI交渉のハイレベルモデル」

私たちが理解しようとしている戦略的状況:
「[エージェントのアリスとボブの間の『交渉』は、(i) 要求/提案、および(ii) 交渉が失敗した場合（リソースを放っておく、または紛争を引き起こすなど）、それぞれが外部オプションを取る[つまり、何をするか]方針をお互いに信頼できる方法で報告することから成ります。」
「プログラムと呼ばれる、[どのように交渉するか] ための各エージェントの手順は、他のエージェントのプログラムに関する情報や、戦略的状況の他の特徴を入力情報として受け取ります。…そのため、AI は、必ずしも厳格な要求を固定したり、最初にコミットした人に譲歩したりするのではなく、条件付きのコミットメントを実装できます。」
「[交渉の前に、各エージェントは]情報チャネルを開いたり閉じたりすることができます…それを通じて、他のエージェントが何らかの信頼できる約束をしたかどうかを知ることができます。」
「ゲーム/意思決定理論に関する明らかな合理主義者の混乱への対応」:
「事後最適 =/= 事前最適」: 「私たち一人ひとりは、相手が何を要求するかを正確に知らずに、自分の要求を選択するよう動機付けられています。なぜなら、要求をする前に不確実性を取り除くのを待っていると、自分の約束によって取引に影響を与える機会を失うからです。」
交渉問題における調整ミスのメカニズムの説明
AI の交渉がどのように機能するかについて、（たとえ粗くても）合理的な予測を立てたい場合は、機械的に説明することを目指す必要があります。
[1]
交渉行為。これはありきたりに聞こえるかもしれません。しかし、私の経験では、メカニズムの説明を無視し、このコンテキストに適合しないモデルやヒューリスティックに頼ってしまうことがよくあります。
私の言いたいことを理解するために、AI の交渉に関する中心的な質問から始めましょう。
私たちは、AI 交渉人のアリスとボブが技術的に進歩していると仮定しています。

（プログラムを通じて）完全に信頼できる、条件付きの約束をするべきです。そして、個人情報の問題も克服できます。では、なぜ彼らは常に互換性のある要求に応じて調整しないのでしょうか?
私はこれに、上記の「事後最適 =/= 事前最適」という引用で部分的に答えました。
[2]
しかし、その議論は不正確です。たとえば、ここでいう「影響力」とは何を意味するのでしょうか？どのような「不安を解消するのを待つ」ことが適切でしょうか?同様に、「均衡選択」に訴えることで中心的な質問に答えることができるかもしれません。これは問題を示していますが、条件付きコミットメント能力を持つエージェントが調整に失敗する理由は実際には説明されていません。
その代わりに、私たちは調整ミスを機械的に説明するものを明確にする必要があります。これらの説明は、AI 交渉の他の側面についての予測に役立つからです。 、エージェントが個別に安全なパレート改善 (SPI) を使用することを好む場合。それでは、アリスとボブが要求の調整を保証できると考えられるさまざまな方法と、それぞれがどのように失敗する可能性があるかを分析してみましょう。 （これは、「要求の調整ミスは避けられない、または合理的である」または「高価な紛争のリスクを負うことは避けられない、または合理的である」よりも弱い主張であり、どちらも誤りであるように見えます。）
オプション 1: ボブはアリスのプログラムを観察するのを待ちます
ボブが競合する可能性を低くする 1 つの方法は、自分のプログラムの選択をアリスの要求に条件付けることです。つまり、自分のプログラムを選択する前に、まずアリスのプログラムがさまざまなプログラムに対してどの要求を行う可能性が高いかを学習し (たとえば、アリスのプログラムが登録されている信頼できるサーバーにアクセスすることによって)、次にこの情報を更新します。
[3]
ボブがこれを行うと、アリスの要求に合わせた要求をする可能性が高くなります。
しかし、アリスのプログラムは、ボブが

サーバーにアクセスしました。次に:
ボブは、アリスが自分のプログラムにこの条項を入れるインセンティブがあることを知っています。「ボブが自分のプログラムをロックインする前にサーバーにアクセスした場合は、そうでない場合よりも多くのリソースを要求します。」
なぜ？なぜなら、彼女はこう考えるかもしれないからです、「もしボブがサーバーを見たら、おそらく彼がそうした理由は、私の要求に合わせてプログラムの選択を条件付けるためだったでしょう。そして、もし彼が私の要求に合わせてプログラムの選択を条件付けたのであれば、それは彼のプログラムがそれらの要求に譲歩するという証拠になります。それなら、私はもっと要求したほうが良いでしょう。」
[4]
したがって、ボブは自分の信念を考慮すると、次のいずれかの選択を迫られます。
(a) 最初に検討することで紛争のリスクは排除されますが、より悪い取引をされる可能性が高くなります。そして
(b) 見ないので衝突の危険はありますが、アリスが同意すればより良い取引が得られます。
ボブの価値観、リスク許容度、紛争で勝つ可能性についての信念に応じて、彼は (b) を好み、アリスのプログラムを学ぶ前にコミットするかもしれません。
(これは、上記の「部分的な回答」によって捉えられたダイナミクスです。)
オプション 2: ボブはアリスのプログラムに自分の要求を条件付けします
ボブは今、どのプログラムを盲目的に選択すべきか迷っています。彼は次のような一般的な形式のプログラムを検討します。「私の最初の要求を何らかの『一次』プログラムで作成します。一次プログラムとアリスのプログラムによって行われた要求が相互に矛盾するかどうかを確認してください。互換性がない場合は、一次プログラムの要求よりも低い要求を作成します。」
[5]
この形式のプログラムは、その要求が非互換性であることを条件としていると言えます。この種のプログラムは、やはりアリスの要求と一致する可能性が高くなります。
しかし、ボブは以前と同じ問題に直面しています。アリスには、自分のプログラムに次の条項を与えるインセンティブがあります。「ボブのプログラムが非互換性を要求する場合は、そうでない場合よりも多くの要求をする」

」したがって、リスクを受け入れるつもりであれば、互換性の要求を条件としないプログラムに取り組むかもしれません。
オプション 3: アリスとボブが意思決定手順を共有する
最後に、たとえ選択肢 1 と 2 が失敗したとしても、おそらくアリスとボブの要求は因果関係なく調整できるのではないでしょうか?アイデアは次のとおりです。アリスとボブが非因果的決定理論に従っていると仮定します。そして、十分に類似した手順に従って意思決定を行っていると考えている場合、彼らはそれぞれ、たとえコミュニケーションを取らなくても、自分の選択が相手の選択を決定するものであると考えるかもしれません。
ここでは、エージェントの意思決定手順が、証拠に基づいてプログラムの選択を生成するプロセスであるとします。 「合理主義者の混乱」の投稿で取り上げられた反論を引用すると、次のようになります。
事実上、一か八かの交渉に参加するのに十分な能力のあるすべてのエージェントは、同じ決定手順に集中し、[公正な要求に従う] という自分の決定は、論理的に相手も同様にすることになると推論します (または、相手も同様にする強力な証拠です)。
投稿の同じセクションで、これも保証されない理由を説明します。この引用は問題のほとんどを捉えています (脚注を追加)。
交渉問題における意思決定手順の収束を期待する理由が、同じ意思決定理論の選択を期待しているためである場合、理想として参照する同じ意思決定理論を共有するだけでは、(特定の問題に関して) 意思決定手順全体を共有することにはなりません。たとえば、意思決定の問題をどのようにモデル化するか、他のエージェントに関するどのような種類の証拠があなたにとって最も顕著であるか、理想的なベイズ主義をどのように近似するかなどを共有しない可能性があります。
[6]
メカニズムの説明から研究フレームワークの改善まで
これで、AI が交渉する理由のメカニズムがわかりました。

rs は互換性のある要求に応じて調整しない可能性があります。ボブは、アリスには、互換性のない要求を回避しようとすると、アリスの要求に合わせてプログラムの選択を条件付けするか、互換性のない要求に条件付けすることによって、より激しく交渉するプログラムを使用するインセンティブがあると予想しています。 （その逆も同様です。）また、非常に似た決定手順に対する十分に強い選択圧力もありません。
しかし、これらすべては、より広範に AI 交渉研究についてどのように考えるべきかということとどのような関係があるのでしょうか?一般的な意味を 2 つ挙げます。
均衡ではなく信念に対する制約
多くの交渉理論を含む標準的なゲーム理論では、合理的なエージェントがナッシュ均衡を形成し、各エージェントの戦略が他のエージェントの実際の戦略に対して最適になると予測されています。
しかし、ここで私たちが焦点を当てている種類の交渉問題において、これは一見すると非常に強力な条件です。 AI が実行する可能性のある条件付きコミットメントは多岐にわたります。したがって、現実的には、AI は互いの戦略について重大な不確実性を抱えることになります。
そして、これが前のセクションの結果です。自分のプログラムを固定する前に、この不確実性を解決しないインセンティブが存在します。つまり、エージェントが平衡の予測が理にかなうのに十分な情報をカウンターパートから得た場合、エージェントは（自分自身にとって）さらに悪い平衡状態に陥る可能性があります。私は、この種のインセンティブを明示的に認める、ナッシュ均衡を予測するための標準的な根拠を知りません。これは、たとえそれが非常に確立されているため、一般的に均衡優先の枠組みを支持する傾向があるとしても、特に AI 交渉の文脈におけるそのような予測に疑念を抱かせるはずです。 (当然のことながら、均衡に対する懐疑論には学術的な見解がある。

認識論的ゲーム理論における先例 (h/t Jesse Clifton)。
[7]
[8]
)
さて、「自明でない不確実性」だけでは、AI が平衡状態にないことを意味するものではありません。高度な AI が、調整ミスを避けるのに十分合理的な信念をお互いに持っていることを依然として期待すべきではないでしょうか?
[9]
そうすべきかどうかわかりません。その見解に対するいくつかの擁護と、私がそれらの見解に説得力を感じない理由は次のとおりです。
反論: 「一般に、合理的なエージェントはよく調整された信念を持っています。お互いのプログラムを直接観察しなくても、互いの過去の行動や、あなたが提唱している種類のメカニズムの説明について慎重に推論することから、多くのことを学ぶことができます。(Kalai and Lehrer (1993) 、「合理的な学習はナッシュ均衡につながる」、および Mailath (1998) のようなこの方向の正式な結果もあります。 、「人々はナッシュ均衡をとりますか?」(セクション 4)。)」
応答: 重要なのは、与えられた対話における互いの戦略についてのエージェントの信念です。そして問題は、これらの信念が、AI 同士が持つであろうデータと、「合理的な」事前分布に対する制約によって十分に決定されていないように見えることです。たとえば、アリスが、マルチエージェントのコンテキストでのボブの過去の決定に基づいて、ボブのプログラムの選択を予測しようとしているとします。ボブは、他人の戦略的推論についての信念に基づいてそのような決定を下しました。しかし、彼はおそらく、独自の戦略的予測を立てようとしながら、時間の経過と文脈を超えてこれらの信念を更新してきました。言うまでもありませんが

[切り捨てられた]

## Original Extract

To understand why powerful AIs might get into conflict, and ways to mitigate it, we need to understand bargaining problems: situations where multiple…

Login Taboo “equilibrium”: Less confused frames for research on AI bargaining — LessWrong Safe Pareto Improvements Game Theory Open Source Game Theory AI World Modeling Frontpage 25
Taboo “equilibrium”: Less confused frames for research on AI bargaining
To understand why powerful AIs might get into conflict, and ways to mitigate it, we need to understand bargaining problems : situations where multiple agents have different preferences over Pareto-efficient outcomes.
I’ve come to suspect that certain common frames on bargaining problems are confused. Here, I’ll explain why, and which frames I think are better. One motivation for this is to hopefully help others make progress in research on safe Pareto improvements (SPIs) , which are among the most promising approaches to mitigating the downsides of AI conflict, in my view.
give some relevant background from my previous writings; ( more )
introduce mechanistic explanations as a crucial frame for AI bargaining research, using, as a running example, explanations of failure to coordinate demands in bargaining (this example also motivates the following bullets’ claims); ( more )
argue that instead of asking what would happen in Nash equilibrium, we should ask what kinds of beliefs AI bargainers would plausibly have about each other; ( more ) and
argue that instead of relying much on heuristics about (e.g.) what agents would consider “exploitative”, we should carefully spell out which commitments they expect to have good consequences. ( more )
I recommend reading the following beforehand (at least the linked sections), but it’s not strictly required, since the essential parts for our purposes are in the quotes below.
“A high-level model of AI bargaining” , which spells out what kinds of strategic situations we’re aiming to understand:
“[‘Bargaining’ between agents Alice and Bob] consists in credibly reporting to each other their (i) demands/offers and (ii) policies for which outside options they’d each take[, i.e., what they’d do] if bargaining failed (such as leaving the resource alone, or initiating conflict).”
“Each agent’s procedure for [how they’ll bargain], called a program , takes as input information about the other agent’s program, as well as other features of the strategic situation. … So, the AIs can implement conditional commitments , instead of necessarily either locking in rigid demands or conceding to whoever commits first.”
“[Before bargaining, each agent can] open or close information channels … through which they learn whether the other agent has made some credible commitment.”
“Responses to apparent rationalist confusions about game / decision theory” :
“ Ex post optimal =/= ex ante optimal” : “Each of us is incentivized to choose our demand without knowing what exactly the other will demand , because if you wait to eliminate your uncertainty before making a demand, you lose the opportunity to influence the bargain with your commitment.”
Mechanistic explanations of miscoordination in bargaining problems
If we want to make reasonable (even if coarse-grained) predictions about how AI bargaining will work, we should aim to mechanistically explain
[1]
bargaining behavior. This may sound trite. But in my experience, it’s easy to neglect mechanistic explanations and fall back on models or heuristics that don’t fit this context.
To see what I mean, we’ll start with the central question about AI bargaining:
We’re supposing that our AI bargainers Alice and Bob are technologically advanced enough to make fully credible, conditional commitments (via programs). And they can overcome private information problems . Then, why wouldn’t they always coordinate on compatible demands?
I’ve partly answered this with the “ Ex post optimal =/= ex ante optimal” quote above.
[2]
But that argument is imprecise. For example, what does “influence” mean here? What kinds of “waiting to eliminate your uncertainty” are relevant? Likewise, we might answer the central question by appealing to “equilibrium selection” . This gestures at the problem, but doesn’t really explain why agents with conditional commitment abilities would fail to coordinate.
Instead, we should get clear on what mechanistically explains miscoordination — because these explanations will inform our predictions about other aspects of AI bargaining, e.g. , when agents will individually prefer to use safe Pareto improvements (SPIs). So let’s break down different ways we might think Alice and Bob can guarantee their demands are coordinated, and how each of them can fail. (This is a weaker claim than “miscoordination of demands is inevitable or rational ” or “risking costly conflict is inevitable or rational”, both of which seem false.)
Option 1: Bob waits to observe Alice’s program
One way for Bob to make conflict less likely is to condition his program choice on Alice’s demands , that is: before choosing his program, first learn which demands Alice’s program is likely to make against different programs (e.g., by accessing a trusted server where Alice’s program is registered), then update on this information.
[3]
If Bob does this, he’s more likely to make demands compatible with Alice’s.
But suppose Alice’s program can make different demands depending on whether Bob accessed the server . Then:
Bob knows that Alice has an incentive to put this clause in her program: “If Bob accessed the server before locking in his own program, then demand more of the resource than otherwise.”
Why? Because she might think, “If Bob looked at the server, probably the reason he did so was to condition his program choice on my demands. And if he conditions his program choice on my demands, that’s some evidence that his program will concede to those demands. Then I’m better off demanding more.”
[4]
So, given his beliefs, Bob faces the choice between:
(a) looking first, thereby eliminating the risk of conflict but making it more likely he gets a worse deal; and
(b) not looking, thereby risking conflict but getting a better deal if Alice agrees.
Depending on Bob’s values, risk tolerance, and beliefs about the likelihood of winning in the conflict, he might prefer (b), and commit before learning Alice’s program.
(This is the dynamic captured by the “partial answer” above.)
Option 2: Bob conditions his demands on Alice’s program
Bob now wonders which program he should blindly choose. He considers a program of this general form: “Make my initial demands with some ‘first-order’ program. Check if the demands made by the first-order program and Alice’s program against each other would be incompatible. If so, make lower demands than the first-order program’s demands.”
[5]
We’ll say a program of this form conditions its demands on incompatibility . This kind of program, again, would be more likely to make compatible demands with Alice’s.
But Bob faces the same problem as before! Alice has an incentive to give her program the clause: “If Bob’s program conditions its demands on incompatibility, then demand more than otherwise.” So if he’s willing to accept the risk, he might commit to a program that doesn’t condition its demands on incompatibility.
Option 3: Alice and Bob share a decision procedure
Finally, even if Options 1 and 2 fail, perhaps Alice’s and Bob’s demands can be acausally coordinated ? The idea is: Suppose Alice and Bob follow acausal decision theories. Then, they each might regard their choice as determining the other’s even without communicating, if they think they make decisions according to sufficiently similar procedures.
For our purposes, let’s say an agent’s decision procedure is the process that generates their choice of program given their evidence. To quote an objection addressed by the “Rationalist confusions” post:
Virtually all agents who are sufficiently capable to enter high-stakes bargaining interactions will converge on the same decision procedure, and reason that their decision to [commit to a fair demand] logically causes their counterparts to do likewise [or is strong evidence that they’ll do likewise].
In that same section of the post, I explain why this, too, isn’t guaranteed. This quote captures most of the problem (footnote added):
If the reason to expect convergence of decision procedures in bargaining problems is that we expect selection for the same decision theory , then … merely sharing the same decision theory that you consult as an ideal doesn’t mean you’ll share a whole decision procedure (with respect to the given problem). For example, you might not share how you model the decision problem, what kinds of evidence about other agents are most salient to you, how you approximate ideal Bayesianism, etc.
[6]
From mechanistic explanations to better research frames
So, we’ve got a mechanistic story for why AI bargainers might not coordinate on compatible demands: Bob anticipates that Alice has an incentive to use a program that bargains harder if he tries to avoid incompatible demands — either via conditioning his program choice on her demands, or conditioning his demands on incompatibility. (And vice versa.) And there aren’t sufficiently strong selection pressures toward very similar decision procedures.
But what does all this have to do with how we should think about AI bargaining research, more broadly? Here are two general implications.
Constraints on beliefs, rather than equilibrium
Standard game theory, including lots of bargaining theory, predicts that rational agents will play a Nash equilibrium, where each agent’s strategy is optimal against the others’ actual strategies.
But in the kinds of bargaining problems we’re focusing on here, this is a very strong condition on its face! There’s a wide range of possible conditional commitments the AIs might implement. So realistically, AIs will have nontrivial uncertainty about each other’s strategies.
And here’s the payoff of the previous section: There’s an incentive not to resolve this uncertainty before locking in one’s program. That is, if an agent gains enough information about their counterpart for the prediction of equilibrium to make sense, this could leave them in a worse equilibrium (for themselves). I’m not aware of any standard justifications for predicting Nash equilibrium that explicitly acknowledge this kind of incentive. This should make us suspicious of such predictions in the AI bargaining context specifically, even if we’re inclined to defer to the equilibrium-first framework in general because it’s very well-established. (For what it’s worth, skepticism about equilibrium has academic precedent in epistemic game theory (h/t Jesse Clifton).
[7]
[8]
)
Now, “nontrivial uncertainty” alone doesn’t imply the AIs won’t be in equilibrium. Shouldn’t we still expect advanced AIs to have beliefs about each other that are reasonable enough to avoid miscoordination?
[9]
I’m not sure we should. Here are some defenses of that view and why I don’t find them convincing:
Objection: “Rational agents will generally have well-calibrated beliefs. Even without directly observing each other’s programs , they can learn a lot from each other’s past behavior, and from reasoning carefully about the sorts of mechanistic explanations you’re advocating. (There are also formal results in this direction, like Kalai and Lehrer (1993) , ‘Rational Learning Leads to Nash Equilibrium’, and Mailath (1998) , ‘Do People Play Nash Equilibrium?’ (Sec. 4).)”
Response: What matters are the agents’ beliefs about each other’s strategies in the given interaction . And the problem is that these beliefs seem underdetermined by the data the AIs would have about each other, plus any constraints on “reasonable” priors. For instance, suppose Alice tries to forecast Bob’s program choice based on his past decisions in multi-agent contexts. Bob made such decisions based on his beliefs about others’ strategic reasoning. But he has presumably updated these beliefs over time and across contexts, while trying to make his own strategic forecasts. Not to mentio

[truncated]

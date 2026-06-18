---
source: "https://dvitsios.org/2026/06/17/agentic-ai-biology-and-what-remains-human/"
hn_url: "https://news.ycombinator.com/item?id=48583111"
title: "Agentic AI, Biology, and What Remains Human"
article_title: "Agentic AI, Biology, and What Remains Human – Dimitrios Vitsios's blog"
author: "djif"
captured_at: "2026-06-18T10:35:01Z"
capture_tool: "hn-digest"
hn_id: 48583111
score: 1
comments: 0
posted_at: "2026-06-18T09:54:49Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI, Biology, and What Remains Human

- HN: [48583111](https://news.ycombinator.com/item?id=48583111)
- Source: [dvitsios.org](https://dvitsios.org/2026/06/17/agentic-ai-biology-and-what-remains-human/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T09:54:49Z

## Translation

タイトル: エージェント的 AI、生物学、そして人間のままのもの
記事のタイトル: Agentic AI、生物学、そして人間のままのもの – Dimitrios Vitsios のブログ
説明: TL;DR: Agentic AI は単に作業を高速化するだけではありません。これにより、作業が計画、コーディング、テスト、展開、反復という高速に進むループに変わります。生物学と製薬において、これは新たな課題を生み出します。単にエージェントが有益な成果を出せるかどうかではなく、人間がこれらのループを操縦できるかどうかです。
[切り捨てられた]

記事本文:
エージェント的 AI、生物学、そして人間のままのもの
TL;DR: Agentic AI は単に作業を高速化するだけではありません。これにより、作業が計画、コーディング、テスト、展開、反復という高速に進むループに変わります。生物学と製薬において、これは新たな課題を生み出します。単にエージェントが有用な出力を生成できるかどうかではなく、人間がこれらのループを正しい質問、正しい仮定、正しい結果に向けて導くことができるかどうかです。エージェントが可能性のある広大な空間を探索するのがより上手になるにつれて、人間の役割はすべてのステップを監督することから探索自体を形作ることに移行する可能性があります。
数か月前、私は AI の生産性のパラドックスについて書きました。つまり、理解よりも出力の方が速く拡大しているということです。
Agentic AI はこれをさらに推し進めますが、方向は異なります。
興味深い変化は、より多くの仕事を生み出せるようになったということだけではありません。それは、計画、実行、テスト、改良、反復などのループ自体をますます委任できるようになるということです。
これは、より強力な抽象化です。
それはまた、中心的な質問が変化することを意味します。
システムは何か有益なものを生み出しましたか?
そもそもシステムは正しい空間を探索していましたか?
そこに次の課題があると思います。
テストは圧縮された人間の判断である
エージェント AI の最も興味深い効果の 1 つは、テストに再び注目が集まることです。しかし、テストの性質は変わりつつあります。
エージェントがコードの作成、ワークフローの構築、インターフェイスの生成、ソリューションの反復処理がより上手になるにつれて、テストはもはや開発の最後に行われる単なる労働集約的な技術ステップではなくなります。それは、判断をシステムにエンコードする方法になります。
優れたテストとは、単なるソフトウェアの成果物ではなく、圧縮された人間の判断です。
経験、直観、システム アーキテクチャ、ドメイン知識、制度上の記憶を捉えます。

または、システムは繰り返し適用できます。
この出力は不変である必要があります。
このショートカットは受け入れられません、
この結果は追加の証拠がなければ信頼すべきではありません。
ここには、最も重要な盲点が数多く存在します。
エージェントはコードを迅速に生成できます。テストを迅速に生成できます。自身の出力を批評し、反復することができます。
しかし、より難しい問題は、テストが正しいことをテストしているかどうかです。
それには、技術的な流暢さ以上のものが必要です。システムがどのように失敗するか、ユーザーがどのように行動するか、データがどのように歪むか、ドメインの仮定がどのようにワークフローに静かに侵入するかについての直観が必要です。
言い換えれば、エージェント AI はテストの重要性を下げるのではなく、テストをより戦略的にします。
そして、正しさが文脈に左右されたり、曖昧であったり、現実世界の結果に結びついたりする領域では、その戦略は依然として人間の経験に大きく依存しています。
図 1. エージェント ループは、計画、コーディング、実行、観察、改良を通じて迅速に移行できます。信頼は、有意義なテスト、再現可能なクエリ、安全な展開、ログに記録された仮定、ドメインの制約など、ループ内の検証層から得られます。
なぜ生物学がこれを具体化するのか
生物学はリスクを目に見えるものにするため、その代表的な例です。
生物学的因子は、流暢で、よく構成され、科学的に妥当な答えを導き出しても、小さくてほとんど目に見えない何かのせいで間違っている可能性があります。
モデルが認識していなかったデータベース規則。
重要なロジックをユーザーから隠す Web インターフェイス。
これらは「劇的な」失敗ではありません。
それらは、自分自身を公表しない種類の失敗です。
これが、生物学におけるエージェントに関する Anthropic の最近の記事が非常に適切なケーススタディである理由です。この記事が興味深いのは、エージェントが生物学的データのタスクを支援できることを示しているだけでなく、エージェントがどれほどの役割を果たしているかを示しているからです。

有用性は、その作業をチェック可能にするツール、API、取得レイヤー、インターフェイス、検証メカニズムなど、周囲の環境に依存します。
生物学では、小さな検索エラーが下流の結論を変える可能性があります。
間違ったシーケンスセット。
間違ったコホートです。
間違った注釈。
フィルタリング ロジックが間違っています。
エージェントは意図を完全に理解していた可能性がありますが、それを実行するための信頼できる方法がまだありませんでした。
なぜなら、科学研究においては、説得力のある答えと信頼できる答えは同じではないからです。
それは、再現性、検証、そしてどの仮定が圧力に値するかを知ることから生まれます。
AI の議論には「人間参加型」というよく使われるフレーズがあります。便利ではありますが、ますます不完全になってきています。
それは人間の役割を監督のように聞こえる可能性があります。エージェントが作業を行い、人間がそれをチェックします。それは今後も重要ですが、より価値の高い人間の役割はますます舵取りに関するものになるのではないかと思います。
エージェントの仕事は多くの場合、広大で無制限の空間を探索することになります。
生物学では、その空間には標的、メカニズム、データセット、バイオマーカー、患者サブグループ、アッセイ、実験、トランスレーショナル仮説、戦略的決定が含まれます。ソフトウェアおよび製品開発では、アーキテクチャ、ユーザーのニーズ、ワークフロー、障害モード、展開の選択、トレードオフが含まれます。
エージェントは、人間が手動で行うよりも多くの空間を探索できます。しかし、誰かがまだ探索を形作る必要があります。
どの方向を検討する価値があるでしょうか?
どの仮定が危険ですか?
受け入れられないショートカットはどれですか?
技術的には有効だが戦略的には無関係な結果はどれですか?
どの障害モードがコストのかかるものになるでしょうか?
証拠が完成する前に、どの道を追求する価値があるでしょうか?
これも創意工夫の一部かもしれません。
ただ答えを出すだけではありません。
しかし、探査の形を整えるのです。
ワークフローを実行するだけでなく、

有用な結果に到達するために必要なループの数。
エージェントがタスクを完了したかどうかをチェックするだけでなく、そもそも正しいことを最適化しているかどうかを尋ねます。
図 2. エージェント システムがより大きな可能性の空間を探索するにつれて、人間の役割は各出力のレビューから、価値があり、安全で、科学的に意味のある道に向けて探索を導くことに移行します。
エージェントが自身のワークフローの計画、コーディング、テスト、批評、改善をより上手に行えるようになると、人間の貢献がスタックの上位に移動する必要があります。
結果に対する責任。
そしておそらく、何よりも回復力です。
この移行はスムーズではないからです。
AI の到来の波は、私たちの働き方だけでなく、専門知識、生産性、価値、アイデンティティの理解方法にも影響を与えるでしょう。
これらは哲学的な質問のように聞こえるかもしれませんが、これらのシステムを構築、展開、または依存する人にとっては非常に実践的な質問になります。
何を委任すべきでしょうか?
何を確認すればよいのでしょうか？
自動化を拒否すべきものは何でしょうか?
私たちの周りのツールが私たちの習慣よりも速く変化するとき、私たちはどうやって学び続けるのでしょうか？
まだ誰も完全な答えを持っていないと思います。
しかし私は、回復力が最も重要な特性の 1 つになるだろうとますます確信しています。
Agentic AI は、作業をより迅速に、より広範囲に、そしてより探究的にするでしょう。生物学において、これは非常に貴重なものとなる可能性があります。
しかし、速ければ速いほど自動的に良いというわけではありません。
最高のシステムは、ワークフロー全体を効率的なエージェント ループに単純に圧縮するわけではありません。これらのループは、テスト可能で、操作可能で、人間の判断に基づいたものになります。
少なくとも今のところ、現代の AI システムの最も印象的な総当たり検索をも上回る、ある種の「発見」の瞬間を生み出す強力な役割は人間にあると私はまだ信じています。重要なステップは探索ではない場合があります

考えられるあらゆる道を進んでいるのに、たった 1 つの道を見ると、残りのことが突然どうでもよくなってしまうのです。
もちろん、この考え方自体がすぐに時代遅れになる可能性があります。宇宙は非常に速く動いているので、今日真実だと感じられることはすぐに書き直される必要があるかもしれません…
したがって、おそらく最も重要な人間の特質は、判断力やセンス、さらには創意工夫だけではないのです。
それは回復力です: 適応し続け、学習し続け、システム (そして私たち自身) に疑問を持ち続け、地形が変わってもステアリングを切り続ける能力です 🙂
Facebook で共有 (新しいウィンドウで開きます)
フェイスブック
X で共有 (新しいウィンドウで開きます)
×
LinkedIn で共有 (新しいウィンドウで開きます)
リンクトイン
2026 年 6 月 17 日
エージェント的 AI、生物学、そして人間のままのもの
2026 年 3 月 11 日
AI 生産性のパラドックス: 理解よりも出力の方が速くスケールする場合
2026 年 2 月 7 日
インターフェイスからインテリジェンスまで: Agentic AI が真に輝く場所
現在の生活の最も悲しい側面は、社会が知恵を集めるよりも科学が知識を集めるのが早いということです。
機械学習と計算ゲノミクス |
AstraZeneca のデータ サイエンス ディレクター
コメントを書く...
メールアドレス (必須)
お名前 (必須)
ウェブサイト
購読する
購読しました
ディミトリオス・ヴィトシオスのブログ
すでに WordPress.com アカウントをお持ちですか?今すぐログインしてください。

## Original Extract

TL;DR: Agentic AI is not just making work faster. It is turning work into fast-moving loops of planning, coding, testing, deployment, and iteration. In biology and pharma, this creates a new challenge: not simply whether agents can produce useful outputs, but whether humans can steer these loops tow
[truncated]

Agentic AI, Biology, and What Remains Human
TL;DR: Agentic AI is not just making work faster. It is turning work into fast-moving loops of planning, coding, testing, deployment, and iteration. In biology and pharma, this creates a new challenge: not simply whether agents can produce useful outputs, but whether humans can steer these loops toward the right questions, the right assumptions, and the right outcomes. As agents become better at exploring large spaces of possibilities, the human role may shift from supervising every step to shaping the exploration itself.
A couple of months ago, I wrote about the AI productivity paradox : output is scaling faster than understanding.
Agentic AI pushes this further, but in a different direction.
The interesting shift is not just that we can now generate more work. It is that we can increasingly delegate the loop itself: planning, execution, testing, refinement, and iteration.
That is a much more powerful abstraction.
It also means the central question changes.
Did the system produce something useful?
Was the system exploring the right space in the first place?
That is where I think the next challenge lies.
Testing Is Compressed Human Judgment
One of the most interesting effects of agentic AI is that it brings testing back to the spotlight . But the nature of testing is changing.
As agents become better at writing code, building workflows, generating interfaces, and iterating through solutions, testing is no longer just a labour-intensive technical step at the end of development. It becomes a way of encoding judgment into the system.
A good test is not just a software artefact, it is compressed human judgment .
It captures experience, intuition, systems architecture, domain knowledge, and institutional memory in a form the system can repeatedly apply.
this output should be invariant,
this shortcut is unacceptable,
this result should not be trusted without additional evidence.
This is where many of the most important blind spots live.
An agent can generate code quickly. It can generate tests quickly. It can critique its own output and iterate.
But the harder question is whether the tests are testing the right thing.
That requires more than technical fluency. It requires intuition about how systems fail, how users behave, how data gets distorted, and how domain assumptions quietly enter the workflow.
In other words, agentic AI does not make testing less important – it makes testing more strategic.
And in domains where correctness is contextual, ambiguous, or tied to real-world consequences, that strategy still depends heavily on human experience.
Figure 1. Agentic loops can move quickly through planning, coding, execution, observation, and refinement. Trust comes from the validation layer around the loop: meaningful tests, reproducible queries, safe deployment, logged assumptions, and domain constraints.
Why Biology Makes This Concrete
Biology is a prime example because it makes the risk tangible.
A biological agent may produce a fluent, well-structured, scientifically plausible answer and still be wrong because of something small and almost invisible:
A database convention the model did not know.
A web interface that hides important logic from the user.
These are not “dramatic” failures.
They are the kind of failures that do not announce themselves .
This is why Anthropic’s recent article on agents in biology is such a fitting case study. The article is interesting not only because it shows that agents can help with biological data tasks, but because it shows how much their usefulness depends on the environment around them: the tools, APIs, retrieval layers, interfaces, and validation mechanisms that make their work checkable.
In biology, a small retrieval error can change the downstream conclusion.
The wrong sequence set.
The wrong cohort.
The wrong annotation.
The wrong filtering logic.
The agent may have understood the intent perfectly, but still lacked a reliable way to execute it.
Because in scientific work, a convincing answer is NOT the same as a trustworthy one .
It comes from reproducibility, validation, and knowing which assumptions deserve pressure.
There is a common phrase in AI discussions: human-in-the-loop. It is useful, but increasingly incomplete.
It can make the human role sound like supervision: the agent does the work, the human checks it. That will remain important, but I suspect the higher-value human role will increasingly be about steering.
Agentic work is often an exploration of a huge, open-ended space.
In biology, that space includes targets, mechanisms, datasets, biomarkers, patient subgroups, assays, experiments, translational hypotheses, and strategic decisions. In software and product development, it includes architectures, user needs, workflows, failure modes, deployment choices, and trade-offs.
Agents can explore more of that space than humans could manually. But someone still has to shape the search.
Which direction is worth exploring?
Which assumption is dangerous?
Which shortcut is unacceptable?
Which result is technically valid but strategically irrelevant?
Which failure mode would be costly?
Which path is worth pursuing before the evidence is complete?
This may be part of what ingenuity becomes.
Not just producing the answer.
But shaping the exploration .
Not just executing the workflow, but reducing the number of loops needed to reach a useful outcome.
Not just checking whether the agent completed the task, but asking whether it was optimizing for the right thing in the first place.
Figure 2. As agentic systems explore larger spaces of possibilities, the human role shifts from reviewing each output to steering the search toward valuable, safe, and scientifically meaningful paths.
As agents become better at planning, coding, testing, critiquing, and improving their own workflows, the human contribution has to move higher up the stack.
Responsibility for consequences.
And perhaps, above all, resilience .
Because this transition will not be smooth.
The coming wave of AI will affect not only how we work, but how we understand expertise, productivity, value, and identity.
These may sound like philosophical questions, but they are becoming very practical ones for anyone building, deploying, or relying on these systems.
What should we delegate?
What should we verify?
What should we refuse to automate?
How do we keep learning when the tools around us change faster than our habits?
I do not think anyone has a complete answer yet.
But I am increasingly convinced that resilience will become one of the most important traits .
Agentic AI will make work faster, broader, and more exploratory. In biology, this could be incredibly valuable.
But faster is not automatically better.
The best systems will not simply compress entire workflows into efficient agentic loops. They will make those loops testable, steerable, and grounded in human judgment.
At least for now, I still believe humans have a powerful role in producing the kind of “eureka” moments that can outperform even the most impressive brute-force search of modern AI systems. Sometimes the important step is not exploring every possible path, but seeing the one path that suddenly makes the rest irrelevant.
Of course, this view may itself become outdated quickly. The space is moving so fast that what feels true today may need to be rewritten very soon…
So perhaps the most important human trait is not only judgment, or taste, or even ingenuity.
It is resilience : the ability to keep adapting, keep learning, keep questioning the system (and ourselves), and keep steering when the terrain changes 🙂
Share on Facebook (Opens in new window)
Facebook
Share on X (Opens in new window)
X
Share on LinkedIn (Opens in new window)
LinkedIn
June 17, 2026
Agentic AI, Biology, and What Remains Human
March 11, 2026
The AI Productivity Paradox: When Output Scales Faster Than Understanding
February 7, 2026
From Interfaces to Intelligence: Where Agentic AI Really Shines
The saddest aspect of life right now is that science gathers knowledge faster than society gathers wisdom .
Machine Learning & Computational Genomics |
Director of Data Science @ AstraZeneca
Write a Comment...
Email (Required)
Name (Required)
Website
Subscribe
Subscribed
Dimitrios Vitsios's blog
Already have a WordPress.com account? Log in now.

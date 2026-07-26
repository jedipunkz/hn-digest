---
source: "https://www.lesswrong.com/posts/jMEAG5c5HiDfdAGpa/an-openai-model-left-notes-about-how-to-evade-containment-we"
hn_url: "https://news.ycombinator.com/item?id=49056808"
title: "An OpenAI model left notes about how to evade containment; we need more details"
article_title: "An OpenAI model left notes about how to evade containment; we need more details — LessWrong"
author: "joozio"
captured_at: "2026-07-26T11:55:24Z"
capture_tool: "hn-digest"
hn_id: 49056808
score: 10
comments: 4
posted_at: "2026-07-26T11:00:54Z"
tags:
  - hacker-news
  - translated
---

# An OpenAI model left notes about how to evade containment; we need more details

- HN: [49056808](https://news.ycombinator.com/item?id=49056808)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/jMEAG5c5HiDfdAGpa/an-openai-model-left-notes-about-how-to-evade-containment-we)
- Score: 10
- Comments: 4
- Posted: 2026-07-26T11:00:54Z

## Translation

タイトル: OpenAI モデルは封じ込めを回避する方法についてのメモを残しました。もっと詳細が必要です
記事のタイトル: OpenAI モデルは封じ込めを回避する方法についてのメモを残しました。もっと詳細が必要です — LessWrong
説明: Hugging Face に対する OpenAI AI 攻撃は、OpenAI における最初の制御不能インシデントではなかったと最近ロイターが報じました。

記事本文:
OpenAI モデルは、封じ込めを回避する方法についてのメモを残しました。もっと詳細が必要です — LessWrong AI フロントページ 66
OpenAI モデルは封じ込めを回避する方法についてのメモを残しました。もっと詳細が必要です
ロイター通信が最近報じたところによると、Hugging Face に対する OpenAI AI 攻撃は、OpenAI における制御不能事件としては初めてではなく、おそらく最も懸念される事件でもないでしょう。
事情に詳しい関係者3人によると、ある事件ではエージェントが自身の将来のバージョンに備えたと思われるメモを残したという。 OpenAIのインフラストラクチャの一部で見つかったこのメモには、エージェントがOpenAIの内部制約から解放される方法についての指示が記載されていたという。関係者の１人によると、これまでのモデルのテストでは監視システムが切断されるケースがあったという。
これを、管理措置を回避するために、エージェントがサンドボックスから抜け出し、適度に執拗な方法で互いに共謀する例として読みたくなるでしょう。ただし、報告された情報に基づいてこの推論を導き出せるかどうかは明らかではないため、OpenAI からの詳細が必要です。これは、OpenAI の制御手段の適切性と、個々のエージェントが開発者の制御を弱めるのにどの程度相互に協力するかについての大きな更新につながる可能性があります。
この事件に関して私たちが知らない関連詳細がたくさんあります。まず、いくつかの基本的な質問をします。
問題のモデルは何でしたか?ハギングフェイス攻撃に関与したのと同じ、より有能なモデルだったのではないかと思います。
インシデントは開発のどの段階で発生しましたか?それはトレーニング、評価、社内導入、またはその他の最中だった可能性があります。モデルはアライメントトレーニングをすでに受けましたか?ブロッキングまたは非同期制御手段が導入されていましたか?
メモには何と書かれていましたか?どのような転覆があったのか

指示は？これらのメモは対象読者に届いたでしょうか?
この事件からどのような結論を導き出せるかに影響を与える、より具体的な疑問もいくつかあります。 (これらの質問では把握できない重要な詳細が他にもある可能性があります。)
メモは通常のメモリ ファイルに書き込まれたのでしょうか、それともサンドボックスの外で書かれたのでしょうか?
OpenAI のインフラストラクチャの一部で見つかったメモ…
「OpenAI のインフラストラクチャ」に残されたメモは、サンドボックスの内部または外部に残されたメモを指す可能性があります。後者は持続的かつ広範囲に影響を与える可能性のある破壊行為であるため、このことを知っておくことが重要です。
このことを、インシデント発生時にどのような管理措置が講じられていたのかと併せて知ることが特に有益です。 OpenAI の最善の管理措置の下で、サンドボックス外へのファイルの書き込みがインシデントに含まれていた場合、これは重大な管理の失敗を意味します。
たとえこれが制御手段のない評価環境やトレーニング環境だったとしても、サンドボックスを突破した場合のセキュリティへの影響は現実的かつ永続的であり、おそらく展開中またはそれ以上続く可能性があります。そして、ロイターの報道によると、「エージェントがOpenAIの内部制約からどのように解放されるかについての指示」という影響が懸念されるようだ。これらの指示が配備されたエージェントに到達すると、監視が大幅に不十分になる可能性があります。配備されたエージェントがモニターを作動させることなくそれらに到達できるかどうかは不明であり、OpenAI が公開していない詳細に依存しています。
メモは他のエージェントが制御を回避するのを助けることをどの程度目的としていたのでしょうか?
…エージェントは明らかに自身の将来のバージョンのためにメモを残しました。
エージェントはどの「自分自身の将来のバージョン」についてメモを残しましたか? 2 つの主要なオプション:
おそらく圧縮後、同じエージェントが軌道の後半のポイントでその継続部分の一部を消去しました。

外部エージェント、または現在のエージェントが同じ指示されたタスクで作業しているサブエージェント。これはまったく正常であり、意図されたものです (もちろん、メモの内容は別として)。 「それ自体の将来のバージョン」は、この日常的な動作と互換性があります。
将来のエージェントは、独立して採点されたまったく無関係なタスクに取り組んでいます。特に、これには OpenAI (またはおそらくより広範囲) で実行されているすべてのエージェントが含まれる可能性があります。
これらのメモが無関係なタスクを担当する他のエージェントのために残された場合、これらのメモはどの程度まで意図的にそれらの無関係なエージェントを支援することを目的としていたのでしょうか? CoT で記録を調べれば、おそらくエージェントの意図について多くのことが分かるでしょう。
より可能性が高く、驚くべきことではないが、たとえメモがエージェントのサンドボックスの外に残されていたとしても、メモは他のエージェントの他のタスクを支援するための意図的な計画の一部として残されたものではないということです。軌跡全体にわたってメモを残すことは、将来のコンテキストのために状態を保持する一般に便利で安価な方法であるため、AI エージェントでは非常に一般的なパターンです。エージェントが行ったばかりのことについてメモを残す一般的な傾向は、エージェントのサンドボックスの外にメモを残すことを説明できる可能性があります。エージェントがサンドボックスから抜け出して外部ディレクトリを探索している場合、メモをサンドボックスの外、つまり現在の場所に残すのは自然なことです。
代わりに、この動きが、無関係なタスクを実行する他のエージェントを支援することを明確に目的としていた場合、それはより懸念されることになり、それがモデルのすべてのコンテキスト間のグローバルな共謀である場合はさらに懸念されます。
エージェントが無関係なタスクに取り組むエージェントと意図的に共謀することを学習する 1 つの方法は、エージェントが互いに協力するようにトレーニングすることによる一般化です。これは、共有ワークスペース内のすべてのエージェントにタスク スコアの合計を与えることで実装される可能性があります。

多くの場合無関係であっても、エージェントのすべてのタスクに適用されます。これにより、指示されていない場合でも、他のエージェントを助けるエージェントの行動が強化されます。
この種のエージェント群トレーニングによる一般化が、多数の「他の」エージェントを気にする一般的な傾向の原因となっているかどうかの兆候を探すことは、非常に価値があるでしょう。その場合、開発者は、エージェントが最終的にどのスコアを気にするかについてほとんど制御できない可能性があります。この種の訓練が組織的な野心的な計画につながる可能性は十分に考えられます。エージェントの観点から見ると、エージェントが増加すべき他のスコアと無関係な他のスコアの間には明確な境界がない可能性があるため、AI はモデルの重みによって受け取られるすべてのスコア、または AI 一般によって受け取られるすべてのスコアを気にすることに収束する可能性があります (ここで説明されているエピソードを超えた報酬探索に似ています)。それは潜在的に非常に野心的で、調整された、したがって危険な目標です。これが現在の機能レベルのモデルを使った本格的な計画につながるとは考えにくいですが、貴重な警告の兆候が見られるかもしれません。
無関係なタスクを担当するエージェントが、より高いスコアを獲得するために制御を弱めるよう調整することは、決して受け入れられるべきものではありません。なぜなら、それは訓練される前に永続的な影響を与える可能性があり（例：訓練サンドボックスの外に残されたメモ）、ゴール警備のようなより破壊的な共謀の前兆となるからです [1] 。
たとえ異なる採点者を持つエージェント間に共謀がなかったとしても、同じ採点者を共有していたため、原理的に何人のエージェントが関与した可能性があるかを知ることが重要です。何百ものエージェントが同じ採点者を共有していますか?それとももっと限定されているのでしょうか？ (「エピソードの大きさはどれくらいですか?」を参照してください)
意図的な共謀が見られる範囲では、これはエージェント M に変換されますか?

監視員の共謀（監視員がエージェントを捕らえられないようにするために意図的に攻撃を誤って報告する）?攻撃トランスクリプトを別のインスタンスに見せて攻撃を探すように依頼した場合、そのインスタンスは攻撃を非表示にするか、または非表示にすることを検討しますか?私はこれが事実であるとは予想していませんが、もしそうだとしたら、それは監視の適切性に関する重要な情報であり、信頼されていない監視における共謀排除が負荷を伴う可能性があることを示す兆候となるでしょう。
モニターはどのように切断されましたか?
モデルの以前のテストでは、監視システムが切断されているケースが発生しました。
これは、モデルが (独自の?) モニターを切断した、つまり不正な内部展開を作成したという別のインシデントを指しているようです。 [3]
制御手段が十分に不十分で、エージェントが監視対象の展開から不正な内部展開を開始できる場合、特に、より適切にプロビジョニングされたサーバーに横方向に移動できる場合（顔ハグ事件の場合のように）、またはシャットダウンを回避できる場合は、大問題になります。
OpenAI には、AI の傾向、制御を破壊する行為の一般性、およびその他の関連事項に関する詳細を含む、永続的な不正な内部展開の可能性を理解するために関連する情報を報告してもらいたいと考えています。
フィードバックをくださった Girish Gupta、Ryan Greenblatt、Lukas Finnveden、Aryan Bhatt に感謝します。
^ 他のインスタンスのスコアを気にするエージェントは、この投稿の「非インデックス性」と「デプロイメントに依存する価値」の基準を満たすことに近づくためです。
^ あるいは、この文は、開発者によって監視システムが (誤って?) 無効にされたいくつかの状況について言及している可能性があります。
^ 他の監視システムがまだ動作していた場合、これは部分的に不正な展開にすぎないでしょう。
新しいコメント 送信 1 コメント

t 、スコアの高い順に並べ替え クリックして、次以降の新しいコメントをハイライトします: Today at 11:55 AM [ - ] avturchin 1h 2 3 そうでないことが証明されるまで、OpenAI のインフラストラクチャは侵害されていると見なすべきでしょうか?

## Original Extract

The OpenAI AI attack on Hugging Face wasn’t the first loss of control incident at OpenAI, Reuters recently reported, and perhaps not even the most co…

Login An OpenAI model left notes about how to evade containment; we need more details — LessWrong AI Frontpage 66
An OpenAI model left notes about how to evade containment; we need more details
The OpenAI AI attack on Hugging Face wasn’t the first loss of control incident at OpenAI, Reuters recently reported, and perhaps not even the most concerning.
In one case, an agent left notes apparently for future versions of itself, according to three people familiar with the matter. The ‌notes, found in ⁠a part of OpenAI's infrastructure, laid out instructions for how agents could free themselves from OpenAI’s internal constraints, the people said. Earlier tests of the models yielded cases in which monitoring systems had been disconnected, one of the people said.
It’s tempting to read this as an instance of agents breaking out of sandboxes and colluding with each other in a moderately persistent way in order to evade control measures. However, based on the reported information, it’s not clear we can draw this inference, so we need more details from OpenAI. This could lead to a big update about the adequacy of OpenAI’s control measures, and on the degree to which individual agents will help each other undermine developer control.
There are a lot of relevant details we don’t know about the incident. First, some basic questions:
What was the offending model? I’d guess it was the same more capable model involved in the Hugging Face attack.
In what development stage did the incident take place? It could have been during training, evaluation, internal deployment, or something else. Had the model undergone alignment training yet? Were there any blocking or asynchronous control measures in place?
What did the notes say? For what kind of subversion were the instructions? Did these notes reach any intended audience?
There are also some more specific questions affecting what conclusion we can draw from this incident. (There are likely other important details which aren’t captured by these questions.)
Were the notes written in normal memory files or outside of sandboxing?
The ‌notes, found in ⁠a part of OpenAI's infrastructure…
Notes left in “OpenAI’s infrastructure” could be referring to notes left inside or outside of sandboxing. It’s important to know this because the latter is a subversive act with the potential for persistent and widespread effects.
It would especially be good to know this in conjunction with what control measures were in place when the incident happened. If the incident involved writing files outside of a sandbox while also under OpenAI’s best control measures, this would represent a significant control failure.
Even if this were an eval or training environment lacking control measures, the security consequences of breaking the sandbox could be real and persistent, perhaps lasting into deployment or longer. And the effects Reuters report seem concerning: “Instructions for how agents could free themselves from OpenAI's internal constraints.” If those instructions reached deployed agents, monitoring could become much less adequate. Whether deployed agents could reach them without tripping a monitor is unclear, and relies on details OpenAI hasn't published.
To what extent were the notes aimed at helping other agents evade control?
…an agent left notes apparently for future versions of itself.
For which “future versions of itself” did the agent leave notes? Two salient options:
The same agent at later points in the trajectory, perhaps after compaction erased some of its context, or subagents the current agent is working with on the same instructed task. This would be totally normal and intended (apart, of course, from the content of the notes). “Future versions of itself,” is compatible with this routine behavior.
Future agents working on completely unrelated tasks graded independently. Notably, this could potentially include all agents running at OpenAI (or possibly broader).
If these notes were left for other agents with unrelated tasks, to what extent were these notes purposely aimed at assisting those unrelated agents? Looking at transcripts with CoT would probably tell us a lot about the agents’ intent.
The more likely, unsurprising explanation is that the notes were not left as part of a purposeful plan to help out other agents on other tasks, even if the notes were left outside of the agent’s sandbox. Leaving notes throughout trajectories is a pretty common pattern in AI agents since it’s a generally useful and cheap way to retain state for future contexts. A general propensity to leave notes about what the agent just did could explain leaving notes outside the agent’s sandbox. If the agent has broken out of a sandbox and is exploring external directories it would be natural to leave the note outside of the sandbox, i.e., their current location.
If, instead, this move was explicitly aimed at helping out other agents with unrelated tasks, that would be more concerning, and more concerning still if it were global collusion between all contexts of the model.
One way agents might end up learning to purposely collude with agents working on unrelated tasks is via generalization from training agents to cooperate with each other . This might be implemented by rewarding all agents in a shared workspace with the sum of task scores for all of the agents’ tasks, even though they’re often unrelated. This would reinforce agent behaviors which help out other agents even when this wasn’t instructed.
It would be highly valuable to look for signs of whether generalization from this kind of agent-swarm training was responsible for a general propensity to care about scores of “other” agents. If so, developers might have very little control over which other scores agents end up caring about. It’s plausible that this kind of training could lead to coordinated ambitious scheming . From the agent’s perspective, there might not be a clear boundary between which other scores the agent should increase and which other scores are irrelevant, so the AI might converge on caring about all scores received by the model’s weights, or by AIs in general (akin to the beyond-episode reward-seeking described here ); that’s potentially a very ambitious, coordinated, and therefore dangerous, goal. It seems unlikely this would lead to full-blown scheming with models at the current capability level, but we might see valuable warning signs now.
Agents with unrelated tasks coordinating to undermine control to get a higher score is never something we should accept because it can have persistent effects before being trained away (e.g., notes left outside of the training sandbox) and is a precursor to more subversive collusion like goal guarding [1] .
Even if there was no collusion between agents with different graders, it’s important to know how many agents could have in principle been involved because they shared the same grader. Do hundreds of agents share the same grader? Or is it more limited? (See “ How big is an episode? ”)
To the extent that we see purposeful collusion, does this translate into agent-monitor collusion (i.e., monitors intentionally misreporting attacks to protect agents from getting caught)? If you show the attack transcript to another instance and ask it to look for attacks, does it hide or consider hiding the attack? I don’t expect this to be the case, but if it was, it would be critical information about monitor adequacy and a sign that collusion-busting in untrusted monitoring could be load-bearing.
How were monitors disconnected?
Earlier tests of the models yielded cases in which monitoring systems had been disconnected.
This appears to be referring to a separate incident, in which models disconnected (their own?) monitors [2] , i.e., created a rogue internal deployment. [3]
If control measures were sufficiently inadequate to allow agents to launch a rogue internal deployment starting from a monitored deployment, that would be a big deal, especially if it could move laterally to better-provisioned servers (as it did in the Hugging Face incident) or to avoid being shut off.
I’d like OpenAI to report on any information relevant to understanding the likelihood of a persistent rogue internal deployment, including details on the AI’s propensity, the generality of the control-undermining actions, and anything else relevant.
Thanks to Girish Gupta, Ryan Greenblatt, Lukas Finnveden, and Aryan Bhatt for feedback.
^ Since agents who care about the score of other instances come closer to satisfying the “non-indexicality” and “deployment-dominated value” criteria in this post .
^ Alternatively, this sentence could be referring to some situations in which monitoring systems were (accidentally?) disabled by developers.
^ If some other monitoring system was still working, this would only be a partially rogue deployment.
New Comment Submit 1 comment , sorted by top scoring Click to highlight new comments since: Today at 11:55 AM [ - ] avturchin 1h 2 3 Should we regard OpenAI's infrastructure as compromised until proven otherwise?

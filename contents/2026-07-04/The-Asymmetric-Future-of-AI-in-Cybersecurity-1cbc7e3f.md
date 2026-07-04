---
source: "https://mstrada.me/posts/aicybersec"
hn_url: "https://news.ycombinator.com/item?id=48784878"
title: "The Asymmetric Future of AI in Cybersecurity"
article_title: "The Asymmetric Future of AI in Cybersecurity"
author: "mstrada"
captured_at: "2026-07-04T13:00:33Z"
capture_tool: "hn-digest"
hn_id: 48784878
score: 3
comments: 0
posted_at: "2026-07-04T12:19:09Z"
tags:
  - hacker-news
  - translated
---

# The Asymmetric Future of AI in Cybersecurity

- HN: [48784878](https://news.ycombinator.com/item?id=48784878)
- Source: [mstrada.me](https://mstrada.me/posts/aicybersec)
- Score: 3
- Comments: 0
- Posted: 2026-07-04T12:19:09Z

## Translation

タイトル: サイバーセキュリティにおける AI の非対称の未来
説明: AI エージェントが攻撃的および防御的なサイバーセキュリティをどのように再構築しているかを、合理的な予測とともに調査します。

記事本文:
メイン コンテンツにスキップ 投稿 タグ 人工知能 - AI について サイバーセキュリティにおける AI の非対称の未来
サイバーセキュリティには常に興味深い特性があります。同じ知識がシステムを保護することも、システムを侵害することもできるということです。概念実証エクスプロイトは、ベンダーが脆弱性を再現してパッチを適用するのに役立ち、また攻撃者がユーザーがシステムを更新する前に脆弱性を武器化するのに役立ちます。
これはどれも新しいことではありませんが、変化しているのは、これらのアクションが実行される速度、規模、アクセスしやすさです。
この投稿は以前の投稿とは少し異なります。最初の部分の目標は、特定の技術概念を説明するのではなく、AI とサイバーセキュリティの間の現在および近い将来の関係に一定の秩序をもたらすことです。後半では、ルーレット テーブルで単に赤か黒に賭けるだけではない、合理的な予測を立ててみたいと思います。
If you already know AI 101 , I suggest skipping to the next section about open models.
エージェントについて議論する前に、この用語が実際に何を意味するのかを簡単に説明することが重要です。なぜなら、この用語は急速に AI で最も過負荷な概念の 1 つになりつつあるからです。 AI エージェントは、タスクを推論し、ツールと対話し、目的に向かって複数のステップのアクションを実行できるシステムです。サイバーセキュリティでは、これは、攻撃対象領域を列挙して脆弱性を連鎖させることができるエージェントから、アラートを自動的に優先順位付けしたり、インシデントに対応したりできるエージェントにまで及びます。
この記事の執筆時点では、AI エージェントは完全に自律的なエンドツーエンドのサイバー操作を一貫して信頼できる成功率で実行することはできません。ただし、特にコーディング、コードの脆弱性検出、生物学の能力は急速に向上しているため、短所も改善されることが期待されています。

時間の経過とともに持続性とスキルが向上します。
The dual-use nature of cybersecurity also exposes a complexity to manage: intent.ランサムウェア展開スクリプトや認証情報を盗むマルウェアの要求など、明らかに悪意のあるシナリオでは、意図を分類するのは比較的簡単です。この課題は、まったく同じ技術的行為が完全に状況に応じて正当であるか有害であるかのどちらかである、より大きなグレーゾーンで現れます。
意図は最終的に言語を通じて表現されるため、これは言語モデルに本質的な問題を引き起こします。 For example, I can ask an LLM to help me validate a bug bounty report or provide a CTF narrative to bypass guardrails and exploit that vulnerability.
この問題を軽減するために、LLM は通常、安全調整技術、ポリシーベースのフィルタリング、ヒューマン フィードバックからの強化学習 (RLHF)、有害な意図や危険な出力を特定するように設計されたランタイム監視システムの組み合わせに依存します。安全メカニズムは通常、 移動しきい値 として機能するため、制限をあまりにも積極的に強化すると、正当なセキュリティ研究者、開発者、防御者がモデルを使用できなくなる危険があります。ただし、緩和しすぎると、悪意のある攻撃者に対する障壁が低くなり、バイパスが容易になります。
この緊張は、Anthropic の Claude Mythos や OpenAI の GPT-Cyber​​ などの新しいモデルが制御されたアクセスまたは信頼できるアクセスの形式を導入した理由の 1 つです。場合によっては、これは部分的に「誇大広告」を高めるためのマーケティング上の決定であると批判されていますが、動機に関係なく、広範なアクセシビリティを公に主張している組織でさえ、検証レイヤーや ID ベースのアクセス制御を導入し始めています。
ローカルモデルの中心的な役割
AI の安全性に関する公的議論の多くは、AI へのアクセスが

コントロールは攻撃能力を有意義に制限することができます。実際には、この仮定は月が経つにつれて弱まっていきます。一元管理の長期的な課題は、企業がアクセス制御や使用ポリシーを強制できるフロンティア クラウド モデルではありません。むしろ今後の中心となるのは、ローカルモデルの台頭であろう。
この記事の執筆時点では、Mythos 5 や GPT 5.6 Sol などの強力なモデルは米国当局によって制限されており、承認された少数の企業のみにアクセスが許可されています。 This underlines how these tools are increasingly becoming strategic assets and highlights the importance of having high-quality open source models to avoid such limitations.
同時に、能力が懸念される場合、将来、政府が一定のしきい値を超える能力のある無差別級モデルに制限を課す可能性があります。オープンウェイト モデルが最先端 (SOTA) モデルに匹敵するパフォーマンスに達した瞬間、政府はクローズド モデルを規制しようとしたのと同じように、オープンウェイト モデルを規制しようとする可能性があります。このような禁止は、LLM 最大手の収益を保護しながら、セキュリティ上の理由から正当化される可能性があります。
オープンソース製品の禁止は、多くの場合、象徴的な禁止にとどまりますが、現実的には思い出せないコピーやミラーの存在、または 1990 年代の強力な暗号化の禁止の試みを考慮すると、重要な違いがあります。
SOTA モデルと同等のパフォーマンスでオープンウェイト モデルを実行するには、かなりのハードウェア リソースが必要です。現在のテクノロジーでは、パフォーマンスを大幅に低下させたり、使用できないほど遅くしたりすることなく、128 GB 未満の RAM を搭載したシステム上で重要なサイバーセキュリティ タスクを解決する機能を備えたモデルを実行することは不可能です。その結果、多くのユーザーは、

ホストされたエンドポイントを介してこれらのモデルにアクセスする必要がありますが、エンドポイントはより集中化されているため、削除がはるかに簡単です。最後に、無差別重量モデルの禁止は国民個人が回避できるかもしれないが、企業が潜在的な罰金にさらされることなく回避することははるかに難しいだろう。
攻撃者と防御者の非対称性
AI に関する議論で一般的に想定されているのは、防御側も単純に AI エージェントを使用し、攻撃能力と防御能力の間に何らかの形のバランスが生まれるということです。この仮定は重大な非対称性を無視しています。つまり、攻撃者と防御者では失敗のコストが根本的に異なります。
より優れたサイバーセキュリティのローカル モデルが中～高価格帯のラップトップで実行できるようになるシナリオを想像してみましょう (おそらくそう遠くないでしょう)。その時点で、攻撃者は、単一のターゲットに対して数十のエージェント (国家支援のアクターの場合は数百の自律エージェント) を簡単に配置し、興味深いものを見つけるか目的を達成するまで、さまざまなパスを試行してループさせることができます。攻撃者にとって、失敗した自動アクションの最悪のシナリオは、多くの場合、比較的限定的です。つまり、攻撃が検出され、ブロックされ、発見された脆弱な資産が修正されます。
防御側にとって、自動化エラーの影響は通常、はるかに深刻です。自律的な防御システムが運用インフラストラクチャを誤って分離したり、内部通信をブロックしたり、重要なサービスを中断したりすると、非常に不幸な会議が発生する可能性があります。防御オートメーションは、企業のビジネスに直接影響を与える運用システムと直接対話するため、はるかに厳しい信頼性制約の下で動作します。
ここでは軍事的な例えが役に立ちます。長年にわたり、高度な軍事システムは以下を中心に最適化されてきました。

非常に高価なプラットフォーム。その後、非対称紛争により、大規模に生産された安価なドローンが、個別にはるかに先進的で高価なシステムを圧倒する可能性があることが実証されました。このことは、数万ドル程度の費用がかかるイランのシャヒド無人機が、1台の迎撃機に数百万ドルかかるパトリオット・ミサイルなどの防空システムによって迎撃されることで明らかになった。 Similar dynamics have also emerged during the Russia-Ukraine conflict.
Cybersecurity may experience a similar transition.たとえ防御者がより強力なモデル、より多くのコンピューティング、より優れたインフラストラクチャを所有していたとしても、攻撃者はスケールと失敗に対するより高い許容度によって依然として優位性を獲得する可能性があります。ある組織が慎重に制限された少数の防御エージェントのみを配備する一方で、別の攻撃者が数分の1のコストで数百の攻撃エージェントを配備し、すべてが単一のより大きなモデルによって調整されている場合、自動化の自由度の非対称性が最終的に攻撃者に有利になる可能性があります。
守備側は今何をすべきでしょうか?おそらく、多くのベンダーが現在提案しているよりも少ないでしょう。
今日のサイバーセキュリティにおける最も誤解を招きやすい話の 1 つは、組織は可能な限り迅速にすべてを積極的に自動化する必要があるという考えです。実際には、前述の理由により、人間のボトルネックに依存しながら自動化しても、見た目ほど多くの利点は得られません。短期的なチャンスは、どこでも完全な自動化ではなく、機能強化、小規模な自動化、優先順位付けです。
通常、セキュリティ チームは過負荷で人員が不足しているため、これは重要です。アラートトリアージ、リバースエンジニアリング支援、脆弱性の優先順位付け、脅威インテリジェンスの要約、検出エンジニアリングなどの反復的なコグニティブタスクを自動化することは、すでに大きな利益をもたらしています。

ts.これらは、組織に制御を完全に引き渡す必要がなく、モデルが摩擦を軽減できる領域です。
Another area where current models are already demonstrating practical value is vulnerability discovery and remediation assistance.モデルは、大規模なコード監査、バリアント分析、バグ発見ワークフローをサポートできることを示しています。 SOTA モデルはゼロデイや複雑な脆弱性を検出できますが、それでも誤検知が発生する可能性があることを忘れないことが重要です。これにより、パイプラインをさらに自動化することが目標の場合、自動検証が便利なツールになります。
これは、人間のボトルネックが最も顕著になる場所でもあります。コードの脆弱性分析を実行し、大規模なコードベースで何百もの脆弱性を発見したモデルがある場合、モデルにそれらを修正させたとしても（ほとんどの企業にとってすでに信頼しすぎているため）、モデルがすべて正しく実行したかどうかを確認するための品質テストと人によるレビュー手順が依然として必要です。人間が最初に何を見る必要があるかを決定するには、明確な優先順位付けプロセスが必要です。自動化により双方の入力量が増加するにつれて、この重要性はますます高まります。全員がより多くの発見を生み出すほど、正しい優先順位付けがより中心的なものになります。
また、軌道は均一ではないため、この機能がどこに向かっているのかを正確に把握することも重要です。これまでのところ、コードの脆弱性の発見においてほぼ指数関数的に改善が見られていますが、その進歩は鈍化すると予想するのが自然です。大量の脆弱性の多くが最初に発見され、将来のモデルの改善により、より小規模でより具体的な脆弱性が発見される可能性があります。わかりやすい例は、Firefox のゼロデイ脆弱性と重大な脆弱性の発見でした。一方、発見され修正された脆弱性の数は増加しました。

残念ながら、Mozilla は実際の脆弱性ごとにどれだけの誤検知が発生したかを公式には明らかにしていません。 \
今後も改善が続く可能性が高い部分は、これまでとは異なり、おそらくより重要な部分です。それは、複数の発見を結び付けて、いくつかの個別の小さな問題を一貫性のある信頼性の高い完全なキル チェーンに連鎖させる機能です。
補足的な点もあります。つまり、対処するのに最も安価な脆弱性は、決して出荷されない脆弱性です。既存のコードの監査に役立つ同じモデルは、作成時点でますます価値が高まっています。変更をマージする前にレビューし、コードの作成中に安全でないパターン、安全でないデフォルト、または入力検証の欠落にフラグを立てることができます。前述の非対称性を考慮すると、攻撃者よりも防御者が失敗に対してはるかに多くの費用を支払うという点で、悪用可能な状態に達する脆弱性の数を減らすことは、意味のあるコストを攻撃者に押し戻す数少ない対策の 1 つです。
これらすべては、マーケティングではめったに言及されていないこと、つまり AI システムの良さはその下にあるデータと同じであることを前提としています。明確な資産の重要度分類がなければ、アラートのトリアージを行うモデルに優先順位を付けることはできず、適切な判断を行うこともできません。

[切り捨てられた]

## Original Extract

An exploration of how AI agents are reshaping offensive and defensive cybersecurity, with a reasoned forecast of where it

Skip to main content Posts Tags About Artificial Intelligence - AI The Asymmetric Future of AI in Cybersecurity
Cybersecurity has always had an interesting property: the same knowledge can either protect a system or compromise it. A proof-of-concept exploit can help a vendor reproduce and patch a vulnerability, or help an attacker weaponize it before users update their systems.
None of this is new, but what is changing is the speed, scale, and accessibility at which these actions could now occur.
This post is slightly different from the previous ones. Rather than explaining a specific technical concept, the goal of the first part is to bring some order to the current and near-future relationship between AI and cybersecurity. In the second part, I will try to make some reasoned predictions that go beyond simply betting on red or black at a roulette table.
If you already know AI 101 , I suggest skipping to the next section about open models.
Before discussing about agents, it is important to give a quick and boring clarification on what the term actually means, because it is rapidly becoming one of the most overloaded concepts in AI. An AI agent is a system capable of reasoning over a task, interacting with tools, and executing multi-step actions toward an objective. In cybersecurity, this could range from an agent capable of enumerating an attack surface and chaining vulnerabilities together, to one automatically triaging alerts, or responding to incidents.
At the time of writing, AI agents are not capable of conducting fully autonomous end-to-end cyber operations with consistently reliable success rates. However, their capabilities are improving quickly, especially in coding, code vulnerability detection, and biology so they are expected to improve in consistency and skill over time.
The dual-use nature of cybersecurity also exposes a complexity to manage: intent. In clearly malicious scenarios, such as requesting ransomware deployment scripts or credential-stealing malware, intent is relatively straightforward to classify. The challenge emerges in the much larger grey area, where the exact same technical action may be either legitimate or harmful depending entirely on the context.
This creates an intrinsic problem for language models because intent is ultimately expressed through language. For example, I can ask an LLM to help me validate a bug bounty report or provide a CTF narrative to bypass guardrails and exploit that vulnerability.
To mitigate this problem, LLMs typically rely on a combination of safety alignment techniques, policy-based filtering, reinforcement learning from human feedback (RLHF), and runtime monitoring systems designed to identify harmful intent or dangerous outputs. Safety mechanisms usually operate as moving thresholds , tightening restrictions too aggressively risks making the model unusable for legitimate security researchers, developers, and defenders. Relaxing them too much, however, lowers the barrier for malicious actors, making bypasses easier.
This tension is one of the reasons why newer models, such as Anthropic’s Claude Mythos and OpenAI’s GPT-Cyber, have introduced forms of controlled or trusted access . In some cases, this has been criticized as partially a marketing decision to increase the “hype”, but regardless of the motivation, even organizations that publicly advocate broad accessibility are beginning to introduce verification layers or identity-based access controls.
The Central Role of Local Models
Much of the public discussion around AI safety assumes that access controls can meaningfully limit offensive capabilities. In practice this assumption grows weaker with each passing month. The long-term challenge for centralized control is not frontier cloud models, where companies can enforce access controls and usage policies. Rather, it will probably be the rise of local models, which will play the central role in the future.
As of the time of writing, powerful models like Mythos 5 and GPT 5.6 Sol have been restricted by US authorities, allowing access only to a small set of approved companies. This underlines how these tools are increasingly becoming strategic assets and highlights the importance of having high-quality open source models to avoid such limitations.
At the same time, if capability is the concern, we might see governments in the future impose restrictions on capable open-weight models above a certain threshold. The moment an open-weight model reaches performance comparable to state-of-the-art (SOTA) models, governments could attempt to regulate them just as they have tried to regulate closed models. Such bans could be justified on security grounds while also protecting the revenue of the largest LLM companies.
While banning open source products often results only in a symbolic ban, given the existence of copies and mirrors that cannot realistically be recalled, or the ban attempt of strong encryption in the 1990s there is an important difference.
Running open-weight models with performance comparable to SOTA models requires substantial hardware resources, and with current technology it is impossible to run a model with the capabilities to solve significant cybersecurity tasks on systems with less than 128 GB of RAM without a significant loss in performance or without making it unusably slow. As a result, many users would still need to access these models through hosted endpoints, which are much easier to take down because they are more centralized. Finally, while a ban on open-weight models could be bypassed by individual citizens, it would be much harder for companies to circumvent without exposing themselves to potential fines.
The Asymmetry Between Attackers and Defenders
A common assumption in discussions about AI is that defenders will simply use AI agents too, creating some form of balance between offensive and defensive capabilities. That assumption ignores a critical asymmetry: the cost of failure is fundamentally different for attackers and defenders.
Let’s imagine a scenario (probably not too far away ) where better cybersecurity local models become runnable on a medium-to-high-range laptop. At that point, an attacker can easily deploy tens of agents against a single target, in the case of state-sponsored actors, hundreds of autonomous agents, and let them loop, trying different paths until they find something interesting or manage to reach the objective. For an attacker, the worst-case scenario for a failed automated action is often relatively limited: the attack gets detected, you get blocked, and the vulnerable assets you spotted get fixed.
For defenders, the consequences of automation errors are usually much more severe. An autonomous defensive system that incorrectly isolates production infrastructure, blocks internal communications, or disrupts critical services can lead to some really unhappy meetings. Defensive automation operates under much tighter reliability constraints because it directly interacts with operational systems that directly affect the company’s business.
A military analogy is useful here. For years, advanced military systems were optimized around highly expensive platforms. Then asymmetric conflicts demonstrated that cheap drones produced at scale could overwhelm systems that were individually far more advanced and expensive. This became evident with Iranian Shahed drones, which cost on the order of tens of thousands of dollars, being intercepted by air defense systems such as the Patriot missile, where each interceptor can cost several million dollars. Similar dynamics have also emerged during the Russia-Ukraine conflict.
Cybersecurity may experience a similar transition. Even if defenders possess stronger models, more compute, and better infrastructure, attackers may still gain an advantage through scale and a higher tolerance for failure. If one organization deploys only a handful of carefully constrained defensive agents, while another actor deploys hundreds of offensive agents at a fraction of the cost, all orchestrated by a single larger model, the asymmetry in automation freedom may ultimately favor the attacker.
What should defenders do now? Probably less than many vendors currently suggest.
One of the most misleading narratives in cybersecurity today is the idea that organizations should aggressively automate everything as quickly as possible. In reality, for the reasons discussed earlier, automating while still depending on a human bottleneck does not provide as many advantages as it appears to. The short-term opportunity is not full automation everywhere, but enrichment, small automations, and prioritization.
This matters because security teams are usually overloaded and understaffed. Automating repetitive cognitive tasks such as alert triage, reverse engineering assistance, vulnerability prioritization, threat intelligence summarization, and detection engineering can already bring significant benefits. These are areas where models can reduce friction without requiring organizations to fully hand over control.
Another area where current models are already demonstrating practical value is vulnerability discovery and remediation assistance. Models are showing that they can support large-scale code auditing, variant analysis, and bug discovery workflows. SOTA models can find zero days and complex vulnerabilities, but it is important not to forget that they can still produce false positives. This makes automated validation a useful companion if the goal is to automate more of the pipeline.
This is also where the human bottleneck becomes most visible. When you have a model performing code vulnerability analysis and finding hundreds of vulnerabilities in a large codebase, even if you let it fix them, which for most companies is already too much trust, you still need quality testing and a human review step to check whether the model did everything correctly. To decide what a human needs to look at first, you need a clear prioritization process. This only grows in importance as automation increases the raw volume of inputs on both sides. The more findings everyone generates, the more central correct prioritization becomes.
It is also worth being precise about where this capability is heading, because the trajectory is not uniform. So far, we have seen near-exponential improvements in code vulnerability discovery, but it is reasonable to expect those gains to slow down: many of the high-volume vulnerabilities will be found first, while future model improvements will likely uncover smaller, more specific vulnerabilities. A clear example was the discovery of zero days and critical vulnerabilities in Firefox: while the number of vulnerabilities found and fixed increased significantly, Mozilla did not officially disclose how many false positives were produced for each real vulnerability. \
The part that will likely keep improving is different, and arguably more consequential: the ability to connect multiple findings together, chaining several individually minor issues into a coherent and reliable full kill chain.
There is also a complementary point: the cheapest vulnerability to handle is the one that never ships. The same models that are useful for auditing existing code are becoming increasingly valuable at the point of creation: reviewing changes before they merge and flagging insecure patterns, unsafe defaults, or missing input validation while the code is still being written. Given the asymmetry described earlier, where defenders pay far more for failure than attackers do, reducing the number of vulnerabilities that ever reach an exploitable state is one of the few moves that meaningfully pushes cost back onto the attacker.
All of this assumes something the marketing rarely mentions: an AI system is only as good as the data underneath it. A model triaging your alerts cannot prioritize if you do not have a clear asset criticality classification, and it cannot reason about ass

[truncated]

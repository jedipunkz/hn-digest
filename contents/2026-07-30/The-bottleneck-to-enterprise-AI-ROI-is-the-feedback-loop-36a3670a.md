---
source: "https://kinesthetic.dev/blog/the-correction-loop/"
hn_url: "https://news.ycombinator.com/item?id=49112664"
title: "The bottleneck to enterprise AI ROI is the feedback loop"
article_title: "The correction loop is the bottleneck — Kinesthetic"
author: "ant-kinesthetic"
captured_at: "2026-07-30T17:16:12Z"
capture_tool: "hn-digest"
hn_id: 49112664
score: 3
comments: 0
posted_at: "2026-07-30T16:59:21Z"
tags:
  - hacker-news
  - translated
---

# The bottleneck to enterprise AI ROI is the feedback loop

- HN: [49112664](https://news.ycombinator.com/item?id=49112664)
- Source: [kinesthetic.dev](https://kinesthetic.dev/blog/the-correction-loop/)
- Score: 3
- Comments: 0
- Posted: 2026-07-30T16:59:21Z

## Translation

タイトル: エンタープライズ AI ROI のボトルネックはフィードバック ループです
記事のタイトル: 修正ループがボトルネック — 運動感覚
説明: 専門家によるパス

記事本文:
修正ループがボトルネックです — Kinesthetic Kinesthetic ホーム ブログの視点 修正ループがボトルネックです
私たちは、実稼働用の専門エージェントを構築するチーム内で同じプロセスを発見し続けています。可観測性プラットフォームからエクスポートされたトレースのスプレッドシートは、対象分野の専門家によって注釈が付けられ、リニア課題/Jira チケットに変換され、最終的には、これに関する会話に参加していなかったエンジニアによる、ほとんどテストされていないプロンプト編集として到着します。このパイプラインは、現在、垂直 AI エージェントの改善のほとんどが実際に行われている場所であり、ソフトウェア開発では機能するが、機械学習モデルを使用する場合には機能しないアドホック プロセスの反映です。
ここでは、このパイプラインが存在する理由、これがエンタープライズ エージェントの展開における最大の制約であると考えられる理由、およびこれに代わるものについて考えます。
AIは誇大広告ではないが、実際にはギャップがある
AIは誇大広告ではありません。エージェントのコーディングは、エージェントが機能する決定的な証拠であると言っても過言ではありません。私たち (そして私たちが話をするほとんどのエンジニア) は、仕事においてクロード コード、コーデックス、カーソルなどから真の価値を引き出しています。さらに、Sierra、Clay、Harvey などの企業はいずれも、ソフトウェアの外部でドメイン固有の作業を行うエージェントを基盤として耐久性のあるビジネスを構築しています。最近、Harvey は最初の純新規 ARR が 1 億ドルを超える四半期に達したばかりです。
しかし、エージェントのコーディングという魔法と、他のドメイン固有の作業を実行するエージェントの現実の間には、依然として明らかな違いがあります。非コーディング エージェントの境界を押し広げている同じ企業が、より現実的なシナリオで現在のモデルが実行できることとのギャップを明らかにするオープン ベンチマークを公開しています。 Harvey Legal Agent Benchmark のリーダーボードと Sierra τ³-Banking のリーダーボードを見ると、Fable 5 のようなフロンティア モデルのスコアが 14.2%、26 であることがわかります。

それぞれ.8%。最良のフロンティア モデルは、両方のベンチマークで個別のルーブリック基準の大部分を満たしていますが、完全に完了しているのはごく一部のタスクだけです。
この失敗の形状からは興味深いことがわかります。このモデルは、提供されたドメイン知識をうまく利用できますが、専門家が承認するような作業を提供することはできません。このギャップは一般的な知識ではなく、タスクを完了するための手順を知っていることと、誰も明示的に書き留めていない微妙な要件を取りこぼすことなく、長期にわたる仕様を理解することです。
他のドメインには (簡単には) 存在しない 2 つの構造上の利点:
コーディング (および数学) は最も検証可能で報酬が得られるタスクとして存在しており、RL 環境の構築が容易になり、このデータの合成や品質チェックが非常に簡単になります。
PR の説明、レビュー スレッド、CI の結果などを含む長期にわたるリポジトリの差分など、ほぼ信頼できる情報密度の高いデータの膨大な量は、他の企業やユースケースでは見られない種類の豊富さと品質です。
あまり注目されない 3 番目の利点もあります。コーディング エージェントでは、システムを構築するエンジニアがドメインの専門家です。翻訳レイヤーや、ドメインの専門家がどのように何かを行うかについての電話ゲームはありません。 「正しい」の意味を知っている人は、ハーネスを変更したり RL 環境を開発したりしている人と同じ人です。
これは、監査、法務、財務、引受業務、その他のホワイトカラーの仕事には当てはまりません。既製のオープン テスト スイートやベンチマークはありません。正しさは、論理的な実行や観察ではなく、主観と専門的な判断の問題です。これには 1 つの疑問が残ります。検証可能な報酬がない領域では、トレーニングはどこで行われるのでしょうか。

信号はどこから来ますか？
ほとんどのチーム内では、それは個人からのものです。エージェントによる裁判官が標準になっている一方で、私たちが話をした人のほとんどは、唯一信頼できる真実の情報源は、エージェントの仕事をレビューし、何が間違っていたのか、そしてなぜ間違っていたのかを注釈を付ける専門家（多くの場合社内）であることに気づいています。これらの検証不可能な業種では、SME が勾配となります。
これが、修正ループが問題の全体である理由です
エキスパートの注釈が学習シグナル (勾配) である場合、そのエクスポートから改善されたエージェントへのパスが学習率となります。現時点では、その道筋は、ドメインの専門家、PM、エンジニアからの電話ゲームであり、損失の多い修正が作成され、多くの場合、不安定または不完全な改善が生じ、アノテーション、トリアージ、開発の別のサイクルが余儀なくされます。
私たちはエージェントを構築している何十ものチームと話をしましたが、特にテクノロジーとインフラストラクチャが初期段階にあるため、運用の非効率性は文化的なものではなく構造的なものであることがわかりました。エンジニアは、プロンプトや知識がハーネス全体でどのように取得/使用されるかわからないため、PM にプロンプ​​トを更新することを望んでいません。PM は、エンジニアが正しくないと思われるドメイン知識について独断的な決定を下した場合に SME に再確認したいと考えています。また、SME はエンジニアリング チームが単に優先順位を付ける方法が分からない大量のフィードバックを生成しています。修正はスプレッドシートにエクスポートされ、チケットとして再エンコードされ、ランディングに必要な 3 か所のうちの 1 つに半分実装され、変更する予定の動作が実際に変更されたかどうかのチェックが行われずに 2 週間後に出荷されます。
フィードバック ループにはクロージャがありません。そして、それを閉鎖するためのインフラは存在しません。
実際には、この問題の接線方向のバージョンに対応するツールがたくさんあります。エージェントS

DK とフレームワークを使用すると、エージェントを午後から立ち上げることができます。可観測性プラットフォームを使用すると、トレースを確認できます。しかし、痕跡を観察することは難しいことではありません。私たちの経験では、これらの可観測性ソリューションを使用しているチームのほとんどは継続的改善ループを閉じていません。データを見ても何を変更すべきかがわからないからです。難しい部分はエージェントを教えることであり、専門家による修正から検証された動作変更までの完全なパスを所有するインフラストラクチャはありません。
この問題は、中小企業の技術力を高めたり、エンジニアの専門性を高めたりすることで解決するとは考えていません。これはチーム全体で起こっていることです。どちらも状況を切り替える税金であり、個人の専門知識を最適に活用することはできません。 SME はエージェントにその分野の語彙を直接教えられる必要があり、エンジニアはそれらの修正を行動に変えるシステムを所有している必要があります。誰かが何を変更したかを記憶するのではなく、評価によってループが閉じられます。
研究面も不安
「修正からどのように学ぶか」に対する明白な答えはトレーニング後にあります。しかし、ほとんどのチームにとって、それは間違った最初の動きです（ほとんどのチームの共通認識です）。実際には、これは通常、構築する RL 環境、ほとんどが満たせないデータ要件、4 ～ 5 回の数値トレーニングの実行 (およびその複数)、および重みの更新を誰が監査するのか、実際に何が変更されたのかに関する未解決のガバナンス ストーリーのように見えます。そして、たとえそれらをすべて構築することができたとしても、基礎となるメソッドはオンラインでの継続的な学習をまだ解決していません。重み空間アプローチは依然としてエントロピー崩壊と壊滅的な忘却に対処する必要があり、最近の研究では、知識の保持、含意のギャップ、KL の発散による基本能力の損傷など、より微妙な失敗が表面化しています。
エクア

最も難しいのは、トークン空間 (コンテキスト学習) での継続的な学習です。研究 ( 1 、 2 ) は、プロンプトベースの学習の問題、つまり、将来のタスクに対する過剰適合と低下、およびメモリ管理とアクセスの問題 (学習内容を効果的にテキストに取り込み、テスト時に適切なコンテキストを取得する方法) を示しています。私たちは、これら 2 つのソリューションの間には中間点が存在する可能性があり、それらの間にフライホイールが存在すると信じています。
今年は両方のトラックが反対方向から同じボトルネックに到達しました。
私たちは、解決策はこの 2 つの間にあると考えています。
Kinesthetic は、重みの更新を行わずに、専門家の修正をエージェントの動作に変えるレイヤーを構築しています。
具体的には、解決および修正された軌跡を取得し、それに含まれる一般化された手順ガイダンス (メタ勾配) をトークン空間に抽出し、テスト時に取得します。そのため、改善はテキスト成果物となり、ドメインの専門家が読んで議論し、承認することができ、エンジニアはバージョンを作成して評価できます。 Harvey LAB および Sierra τ³-Banking ベンチマークで大幅な改善が見られました。詳細については、こちらをご覧ください。
当社のソリューションには譲れない 2 つの特性があり、どちらも規制された業界での作業から生まれています。すべての改善は読みやすく、解釈可能です。動作を変更したアーティファクトを示すことができます。これは、監査機能を備えたあらゆるドメインでの要件であり、コンプライアンス チームが承認する必要があるものに対する厳しい制約です。そして、ループはデータがすでに存在する場所で実行されます。つまり、軌道と修正がパートナーの環境を離れる必要はありません。これは、私たちが追加した機能ではありませんが、そもそもトークン空間で作業する理由です。
これらすべての根底にある賭けは、エージェントを教えることが、最終的には新入社員の新人研修と同じように感じられるはずだということです。

仕事を見せ、判断を修正すれば、仕事はより良くなります。そこからは程遠いです。しかし、次のモデルを待っていてもその差は縮まらないと考えています。
特定の業種向けに特化した AI またはエージェントを構築していて、AI に最適な社内教師を提供する方法について詳しく知りたい場合は、お問い合わせください。

## Original Extract

The path from an expert

The correction loop is the bottleneck — Kinesthetic Kinesthetic Home Blog Perspective The correction loop is the bottleneck
We keep finding the same process inside teams building specialized agents for production: a spreadsheet of exported traces from their observability platforms, annotated by a subject matter expert, is converted into a Linear issue/Jira ticket and eventually lands as a barely tested prompt edit by an engineer who wasn’t in any of the conversations about this. That pipeline is where most vertical AI agent improvement actually happens today: a mirror of the adhoc process that works for software development, but does not hold when working with machine learning models.
Here, we’ll discuss why this pipeline exists, why we think it is the biggest constraint on enterprise agent deployment, and what we think replaces it.
AI is not hype, but there is a real gap
AI is not hype. It feels fair to say that coding agents are the definitive proof that agents work. We (and most engineers we talk to) derive real value from Claude Code, Codex, Cursor, etc. in our work. Additionally, companies like Sierra, Clay, and Harvey have all built durable businesses on agents doing domain-specific work outside of software. Recently, Harvey just hit their first $100M+ net new ARR quarter.
Yet there is still a stark difference between the magic of coding agents and the reality of agents performing other domain-specific work. The same companies pushing the boundary of non-coding agents are publishing open benchmarks exposing the gap between what the current models can do in more realistic scenarios. If we look at the Harvey Legal Agent Benchmark leaderboard and the Sierra τ³-Banking leaderboard, we can see that a frontier model like Fable 5 scores 14.2% and 26.8% respectively. The best frontier models pass the large majority of individual rubric criteria on both benchmarks while fully completing only a small fraction of tasks.
This failure shape exposes something interesting. The model is able to utilize the provided domain knowledge well, but it still fails to deliver work that a professional would sign off on. The gap isn’t general knowledge, it’s knowing the procedure to complete a task and understanding the specification across a long horizon without dropping the nuanced requirements that nobody explicitly wrote down for it.
Two structure advantages that don’t exist (easily) for other domains:
Coding (and math) exists as the most verifiable-rewarded task – allowing for easy RL env building – which makes synthesis or quality checks over this data so much easier.
The sheer amount of largely trustable, information-dense data – long horizon over repository diffs with PR descriptions, review threads, CI outcomes, etc. – is the type of abundance and quality that you won’t see in other enterprises or use cases.
There’s also a third advantage that gets less attention. On a coding agent, the engineer building the system is the domain expert. There is no translation layer, no game of telephone about how a domain expert would do something. The person who knows what “correct” means is the same person who is changing the harness or developing the RL environment.
None of this holds in auditing, legal, financial, underwriting, or other white collar work. There’s no ready-made open test suites or benchmarks. The correctness is a matter of subjectivity and professional judgement rather than logical execution and observation. This leaves one question: in a domain with no verifiable reward, where does the training signal come from?
Inside most teams, it comes from a person. While agentic judges have become a standard, most of the folks we’ve talked to find that the only reliable source of ground truth is the expert (often in-house) who reviews the agent’s work and annotates what was wrong and why. In these non-verifiable verticals, the SME is the gradient .
Which is why the correction loop is the whole problem
If the expert’s annotations are the learning signal (gradient), then the path from that export to an improved agent is your learning rate. Right now, that path is a game of telephone from the domain experts, PMs, and engineers that creates lossy corrections – often resulting in unstable or incomplete improvements, forcing another cycle of annotation, triaging, and development.
We’ve talked to dozens of teams building agents and have realized that the operational inefficiencies are structural, not cultural, especially because of how nascent the technology and infrastructure is. Engineers don’t want PMs to update prompts because they don’t know how those prompts or knowledge are retrieved/used across the harness, PMs want to double check with SMEs when their engineers make arbitrary decisions about domain knowledge that doesn’t look right, and SMEs are producing an overwhelming amount of feedback that the engineering team simply does not know how to prioritize. A correction gets exported to a spreadsheet, re-encoded as a ticket, half-implemented in one of the three places it needed to land, and shipped two weeks later without any check that it actually changed the behavior it was meant to change.
The feedback loop has no closure. And no infrastructure exists to close it.
In reality, there’s plenty of tooling for the tangential version of this problem. Agent SDKs and frameworks let you stand an agent up in an afternoon. Observability platforms let you look at traces. But looking at traces is not the hard part. The teams using these observability solutions mostly aren’t, in our experience, closing the continuous improvement loop, because seeing the data doesn’t tell you what to change. The hard part is teaching the agent, and there is no infrastructure that owns the full path from expert correction to validated behavior change.
We don’t think this is solved by making SMEs more technical or engineers more expert, something we’ve been seeing happening across teams. Both are context-switching taxes that don’t allow the individual expertise to be optimally harnessed. The SME should be able to teach the agent directly, in the vocabulary of their domain, and the engineer should own the system that turns those corrections into behavior. With the loop closed by evaluation rather than by someone’s memory of what they changed.
The research side is unsettled too
The obvious answer to “how do we learn from corrections” is post-training. But for most teams it’s the wrong first move (a general consensus amongst most teams). In practice, this usually looks like an RL environment to build, data requirements that most can’t meet, 4-5 figure training runs (and multiple of them), and an unresolved governance story about who audits a weight update and what actually changed. And even if you manage to build or have all of that, the underlying methods haven’t solved online continual learning yet. Weight-space approaches still need to combat entropy collapse and catastrophic forgetting , and recent work has surfaced subtler failures such as knowledge retention, entailment gap, and base capability damage from KL divergence.
Equally challenging is continual learning in the token space (in context learning). Studies ( 1 , 2 ) show the problems with prompt based learning: namely overfitting and degradation to future tasks as well as the problem of memory management and access (how do you effectively capture learnings in text and retrieve the right context at test time). We believe that there could be a middle ground between these two solutions and that a flywheel exists between them.
Both tracks arrived at the same bottleneck this year, from opposite directions:
We think the solution lives between the two.
Kinesthetic is building the layer that turns expert corrections into agent behavior, without weight updates.
Concretely: we take solved and corrected trajectories, distill generalized procedural guidance (meta-gradient) they contain into token space, and retrieve it at test time, so an improvement is a text artifact a domain expert can read, argue with, and approve, and an engineer can version and evaluate. We have shown significant improvements on the Harvey LAB and Sierra τ³-Banking benchmarks, which you can read more about here .
Two properties of our solution that we consider non-negotiable, both of which come from working in regulated verticals. Every improvement is legible and interpretable: you can point at the artifact that changed the behavior, which is a requirement in any domain with an audit function and a hard constraint for anything a compliance team has to sign off on. And the loop runs where the data already lives: trajectories and corrections never have to leave the partner’s environment, which is not a feature we bolted on but the reason we work in token space at all.
The bet underneath all of it is that teaching an agent should eventually feel like onboarding a new hire: you show the work, you correct the judgment calls, and the thing gets better at the job. We’re a long way from that. But we don’t think the gap closes by waiting for the next model.
If you are building specialized AI or agents for specific verticals and want to learn more about how we provide the best in-house teacher for your AI, contact us !

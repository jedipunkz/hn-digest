---
source: "https://www.nextplatform.com/control/2026/07/15/sres-to-ai-agents-prove-yourself-before-you-touch-production/5271533"
hn_url: "https://news.ycombinator.com/item?id=48945150"
title: "SREs to AI Agents: Prove Yourself Before You Touch Production"
article_title: "SREs To AI Agents: Prove Yourself Before You Touch Production"
author: "rbanffy"
captured_at: "2026-07-17T09:47:39Z"
capture_tool: "hn-digest"
hn_id: 48945150
score: 2
comments: 0
posted_at: "2026-07-17T09:31:24Z"
tags:
  - hacker-news
  - translated
---

# SREs to AI Agents: Prove Yourself Before You Touch Production

- HN: [48945150](https://news.ycombinator.com/item?id=48945150)
- Source: [www.nextplatform.com](https://www.nextplatform.com/control/2026/07/15/sres-to-ai-agents-prove-yourself-before-you-touch-production/5271533)
- Score: 2
- Comments: 0
- Posted: 2026-07-17T09:31:24Z

## Translation

タイトル: SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
記事のタイトル: SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明してください
説明: AI エージェントを信頼してダウンタイムに関するユーザーの苦情を要約することは、別のことです。問題を解決してくれると信じています...

記事本文:
メインコンテンツへジャンプ
検索
その他のトピック
すべてのセクションの最新ニュース
SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
AI エージェントを信頼してダウンタイムに関するユーザーの苦情を要約することは別のことです。無人で問題を解決してくれると信頼することは、まったく別のことです。
The Register が 2026 年 4 月に NeuBird AI を使用して実施した 696 人の専門家に対する調査では、73 パーセントが AIOps をまったく使用しておらず、さらに 19 パーセントが試験運用中で、本番環境に導入しているのはわずか 8 パーセントであることが判明しました。
何が妨げになっているのかという質問では、回答者の 60% が信頼の欠如を挙げ、これが最大の問題であり、ROI、セキュリティ、データ品質に対する懸念がそれぞれ約 12% ～ 13% を占めています。
NeuBird AI の Production Ops Agent は、その信頼不足を解消するように設計されています。アラート キューを要約するのではなく、メトリクス、ログ、トレース、インフラストラクチャ テレメトリ、展開アクティビティ、依存関係の関係を継続的に関連付け、その組み合わせた全体像に対して調査を実行して、考えられる根本原因と次のアクションを提案します。上流のステップでも機能します。 NeuBird AI は、ノイズの多いアラート キューに高速レスポンダーを追加するのではなく、ソースでの可観測性を修正します。エージェント インストルメンテーションを通じて適切なシグナルを生成するため、アラートは設計上高シグナルになります。マーテル氏が言うように、重要なのは、出力にパッチを適用することではなく、ソースでの可観測性を修正することです。 「自社開発のエージェントを既存のキューに向けても問題は解決しません。ノイズに対する DIY は依然としてノイズだからです。」
現場最高技術責任者のフランソワ・マルテル氏がザ・レジスターのインタビューに応じ、調査で判明したことと、AIOps の次の段階がエンジニアが 10 年間見つめてきたダッシュボードと全く似ていない理由について語った。彼はまた、SRE チームがエージェントに業務を任せる前に何を変える必要があるかについても意見を持っています。

彼らの生産システムに耳を傾けてください。
関心は高いが、展開はほとんどない
データは、マーテルが現場ですでに聞いていたことを裏付けました。 「関心は高いが、行動はそれほど多くない」と彼は言う。このパターンはエージェント ワークロード全体でよく見られるものです。人気を博しているカテゴリは、コーディング エージェントやコンテンツ生成など、明らかな人間の関与と明確な検証パスが伴うものです。作業は実行環境内でエンジニアがまだ見ていないデータに基づいて行われるため、運用はさらに難しくなり、その影響は顧客向けシステムに現れます。
彼は、NeuBird に入社するずっと前から、企業内に同様のギャップを感じていました。つまり、300 件の AI 修正候補のバックログと、初期の熱狂的な混乱があり、その後、最初のバージョンが出荷されるまでに 1 年間の苦労が続きました。
ツールが期待に追いつくまで 6 か月待つことが正しい判断である場合もあるため、遅延の一部は市場開発のスピードに起因します。
もう 1 つの原因は、汎用エージェントが SRE の問題に適合しないため、ツール カテゴリの選択が間違っていることです。
マーテル氏は、「はるかに優れた仕事ができる専門エージェントが存在し、安全性、セキュリティ、ガードレール、幻覚などの「懸念事項の一部に対処できる」という。また、ツールはチームの既存のワークフローに適合する必要があります。
マーテル氏は、調査で表面化した信頼性を重視する懸念について反論しようとはしていない。 「AI との連携は信頼を築く作業であり、信頼を得るために AI は学習する必要があります」と彼は言います。 「これは、AI エージェントにとって一種のキラー機能だと思います。学習し、向上していることを示すことができれば、信頼を得ることができます。」
だからこそ、説明可能性がセキュリティレビューのために移植されるのではなく、NeuBird AI の設計の中心に据えられているのです。プラットフォームはあらゆる決定の背後にある理由を記録します

これにより、エンジニアは同僚のインシデント報告書を調査するのと同じように調査できるようになります。 「エージェントがいる場合は常に、下された決定を監査し、決定の背後にある理由を理解できるようにしたいと考えています」とマーテル氏は言います。 NeuBird AI は内部的に、Langfuse を介してあらゆる推論ステップをキャプチャします。説明可能性は半分にすぎません。このプラットフォームは SOC 2 Type II 認定も受けており、読み取り専用で何も保存しないため、論理だけではなくアーキテクチャに信頼が組み込まれています。外部的には、より難しい問題はプレゼンテーションです。システムの初期バージョンでは、非常に多くの詳細が表面化するため、ユーザーはそれをテキストの壁のように表現しました。この修正では、推論をダンプするのではなく質問できるようにすることで、エンジニアが上級チームメイトに質問するのと同じようにシステムのメモリとチャットできるようになりました。
答えを信頼できるものにするのは文脈です
同じ調査では、回答者の 59 パーセントが採用前にほぼ完璧な精度を必要としている一方、さらに 10 人に 3 人は約 80 パーセントの精度を許容していることがわかりました。このハードルは容赦のないもので、マーテル氏は、より大きなモデルではなく、より優れたコンテキストエンジニアリングによってのみクリアできると主張しています。
「正確さの鍵は、何かを見逃さない程度の十分なコンテキストと、コンテキストの発見可能性との間のスイートスポットです」と彼は言います。 「確かに、あまり文脈はありません。」そのバランスを実現するソリューションを作成することは、デスクトップ上にコーディング エージェントだけを使用している人には手の届かないものだと彼は主張します。
NeuBird AI の主張は、ほとんどの機能停止は 1 つのダッシュボードやサービス内では推論できないという事実に基づいています。 SRE チームを必要とするほど大きな企業は、特にマイクロサービスによって資産が断片化された後、ストレージやネットワーキングからオペレーティング プラットフォームやアプリケーションに至るまで、技術スタック全体にサイロが存在します。
捜査官

イオンは、一人の人間が完全に把握できない境界を通過する必要があります。NeuBird は、事件が始まる前に依存関係のマッピングを行うことでこれを処理します。そのため、調査が開始されると、システムはどこを調べればよいのか、各部分がどのように関連しているのかをすでに知っています。
今は副操縦士、後は自律性 - おそらく
この調査から得られた最も明らかなシグナルであり、マーテル氏がそれほど驚くべきことではなかったのは、副操縦士モデルの好みであり、62% が AI に取って代わられるのではなく、AI が支援することを望んでいました。
彼はコーディング エージェントとの自身の仕事からこの段階を認識していますが、進化の弧も認めています。 1 年前、彼はコーディング エージェントから片時も離れることはありませんでしたが、今では、コーディング エージェントを危険モードに切り替えて実行させたいという誘惑に駆られています。ただし、彼は今でもチェックインし、すべてを設計しています。 「私は自分の責任を完全に放棄するつもりはない」と彼は言う。
彼が説明する運用の実際的な道筋も同様に見えます。 NeuBirdAI は、Ansible の Model Context Protocol (MCP) サーバーを介して自動化の連携を開始しており、特定の Playbook は本番環境で自動化しても安全であるとマークされ、残りは人間の承認の後にゲートされます。
既知の上限までポッドにメモリを追加することは、エージェントが処理できることです。より危険なことが人を待っています。マーテル氏によると、エンジニアがどれだけ委任するかは、エンジニアのリスクに対する欲求と、ツールと並行して作業して蓄積した経験によって決まります。
5分時計と作戦室の死
AIOps の概要で最も重要なのは応答時間です。調査回答者の半数強が運用上の回答を 5 分以内に期待しており、75% が 10 分以内を望んでおり、リズムに合わせて構築されていないワークフローに多大なプレッシャーを与えています。 6 人の専門家を状況に合わせて作戦室の会議ブリッジに引き上げるには、SLA では吸収できない時間がかかります

b.
マーテル氏の主張は、時計が変わる前にオンコール体験を変える必要があるということだ。 「他の 20 チームと電話をかけているわけではないという状況に到達したいと考えています。代わりに、何が起こっているのかを説明し、解決策を示してくれたり、誰が関与すべきかを教えてくれたりする文書の前にいるのです。」と彼は言います。
エージェントはエンジニアがログインする前に準備作業を行うため、エンジニアが到着するまでに、初期のトリアージの質問はすでに回答されており、興味深い決定だけが残っています。
可観測性法案にとって IT が意味するもの
少なくとも既存の可観測性ベンダーにとって最も刺激的な調査結果は、AI 主導の洞察がバックエンド全体で機能する場合、回答者の 52% がテレメトリ ツールの切り替えを検討するだろうということです。
これがどうなるかと問われれば、マーテル氏はヘッジをしない。 「将来的には、可観測性は、Grafana、Elasticsearch、OpenSearch などのオープンソースのコスト効率の高いストレージ インデックス作成テクノロジによって支配されることになります。」
そのシナリオでは、戦略的資産は、テレメトリを最も多く蓄積する者から、それを最もインテリジェントに調査できる者に移行します。これは、コンテキスト エンジンが最も安価なストレージ レイヤーの上に位置することを意味します。
これは、オブザーバビリティ契約を更新しようとしている購入者にとって便利なレンズです。なぜなら、購入者が大金を払って購入したダッシュボードは、機械読み取り機能も搭載するシステムの人間が読み取り可能な層であるためです。
この調査では、AI の運用を強く望んでいながらも、ベンダーが証拠なしに結果を約束することに疑念を抱くようになった市場について説明しています。
マーテル氏の売り込みは、今後 2 年間生き残るプラットフォームは、その成果を発揮し、既存の変更管理機構の書き換えを要求するのではなく、既存の変更管理機構に適合するプラットフォームになるだろう、というものだ。勝者は、運用コンテキストを一流のエンジンとして扱います。

急いで詰め込む演習ではなく、深刻な問題を解決します。
マーテルは、自分のチームが時代遅れなのかどうかまだ疑問を抱いている SRE リーダーに対して、率直な答えを示しています。
「安定した運用予算で成長を続ける生産施設に対応できるという点で、メリットが得られます」と彼は言います。 「もし採用しなかったらどうするの？苦労することになるよ。」
AI チップが TSMC の収益の約 3 分の 1 を占める
SRE から AI エージェントへ: 本番環境に触れる前に自分自身を証明する
量子古典 HPC データセンターにおける HPE とデルの願望
AI 対応データが真の利点である理由
QuiX Quantum が HPC データセンター向けのフォトニック アーキテクチャを披露
AI コンピューティングが進むにつれてイーサネット ネットワーキングも進む
もちろんメタプラットフォームはクラウドになるだろう
3 人の HPC 達人が尋ねる: まだ GPU は必要ですか?
光学スケールアップ ファブリックはアーキテクチャではなく製造によって制限される
AMD、フラッシュ拡張メモリでサーバー DRAM を拡張
メモリ市場の好不況サイクルの終焉
中国の「LineShine」オール CPU、エクサフロップスクラスのスーパーコンピューターの詳細
HPE、セキュリティ、主権、マルチテナントのためのアップグレードされた HPC ハードウェア、ソフトウェアを提供
HPC スーパーコンピューティングでは AMD と Nvidia が互角です
企業
HPE、Agentic AI の波に乗ってデータセンターに再び参入
店
Everpure の AI 戦略はほぼ純粋に Nvidia に基づいています
計算する
サーバーブームにより価格上昇とチップ不足が両立
HPE のデータセンター ネットワーキングの全体像がより明確に焦点になる
Quantum Pulse は産業用光の魔法で量子ビットを大幅に向上
Tensordyne が AI 行列計算をログに変換して推論力を強化
QuEra の Libra フォールトトレラント量子システムが Amazon Braket サービスに移行
あなたとデータセンターにさらなるパワーを
オラクルはAIで巨額の利益を得るだろうが、利益を予測するのは難しい
AWS

Agentic AI 向けに Graviton5 を調整し、大きな利益をもたらします
ゲノミクスサンプルあたりのコストは?シーケンス試行あたりのコストを試す
D-Wave はゲート モデル量子の野望のためにデュアル レールに乗ります
計算する
AIチップ・シェパーズ、ブロードコムとマーベルが金羊毛の皮を剥いだ
AI
チップ容量の制約により AI 支出の増加が抑制される
計算する
Intel Xeon 6+ によるサーバー統合のためのパフォーマンスの強化
ハイエンド コンピューティングを詳細にカバー
お問い合わせ
私たちと一緒に宣伝しましょう
私たちは誰なのか
ニュースレター
レジスター
開発クラス
ブロックとファイル
クッキーポリシー
プライバシーポリシー
利用規約
私の個人情報を販売しないでください
同意のオプション
著作権。すべての著作権は © 1998-2026 に留保されます。

## Original Extract

Trusting an AI agent to summarize user complaints about downtime is one thing; trusting it to fix th ...

Jump to main content
Search
More topics
All the latest news, from all sections
SREs To AI Agents: Prove Yourself Before You Touch Production
Trusting an AI agent to summarize user complaints about downtime is one thing; trusting it to fix the problem unattended is something else entirely.
A survey of 696 experts that The Register ran with NeuBird AI in April 2026 found that 73 percent are not using AIOps at all, another 19 percent are in pilot, and only eight percent have it in production.
Asked what is stopping them, 60 percent of respondents cited a lack of trust, by far the biggest issue, with concerns about ROI, security and data quality each registering at around 12 percent to 13 percent.
NeuBird AI's Production Ops Agent is designed to close that trust deficit. Rather than summarizing the alert queue, it continuously correlates metrics, logs, traces, infrastructure telemetry, deployment activity and dependency relationships, then runs investigations across that combined picture to suggest probable root causes and next actions. It also works a step upstream. Rather than bolting a faster responder onto a noisy alert queue, NeuBird AI fixes observability at its source: through agentic instrumentation it generates the right signals, so the alert is high-signal by design. As Martel puts it, the point is to fix observability at the source, not patch the output. "Pointing a homegrown agent at the existing queue doesn't solve that, because DIY on noise is still noise."
Field chief technology officer Francois Martel sat down with The Register to talk through what the survey found, and why the next phase of AIOps will look nothing like the dashboards engineers have stared at for a decade. He also has views on what must change before SRE teams will let agents near their production systems.
Lots Of Interest, Very Little Deployment
The data confirmed what Martel was already hearing in the field. "There's a lot of interest, but not a lot of action," he says. The pattern is familiar across agentic workloads: the categories that have taken off are the ones that come with an obvious human in the loop and an obvious verification path, such as coding agents and content generation. Operations is harder, because the work happens inside the running environment, on data the engineer hasn't seen yet, with consequences that show up in customer-facing systems.
He saw the same gap inside enterprises long before he joined NeuBird: a backlog of 300 candidate AI fixes and a flurry of early enthusiasm, followed by a year of slog before the first one shipped.
Part of that delay comes down to the speed of market development, since waiting six months for the tools to catch up with your expectations is sometimes the right call.
Another part of it is the wrong choice of tool category, because general-purpose agents do not fit SRE problems.
"There are specialized agents that can do a much better job," Martel says, "and address some of the concerns" of safety, security, guardrails and hallucinations. The tool also has to fit into the team's existing workflows.
Martel doesn't try to argue with the trust-heavy concerns the survey surfaced. "Working with AI is a trust-building exercise, and AI has to learn in order to gain trust," he says. "I would say that's kind of the killer feature for AI agents. If you can show that you're learning and getting better, then you can gain trust."
That's why explainability sits at the center of NeuBird AI's design rather than being grafted on for the security review. The platform records the reasoning behind every decision so an engineer can interrogate it the way they'd interrogate a colleague's incident report. "Whenever you have an agent, you want to be able to audit the decisions that were made, and understand the reasoning behind the decision," Martel says. Internally, NeuBird AI captures every reasoning step via Langfuse. Explainability is only half of it. The platform is also SOC 2 Type II certified, read-only, and stores nothing, so trust is built into the architecture, not just the reasoning. Externally, the harder problem is presentation: early versions of the system surfaced so much detail that users described it as a wall of text. The fix was to make the reasoning interrogable rather than dumped, so engineers can chat with the system's memory the way they'd query a more senior teammate.
Context Is What Makes The Answer Credible
The same survey found that 59 percent of respondents require near-perfect accuracy before they'll adopt, while another three in every ten will tolerate around 80 percent accuracy. That bar is unforgiving, and Martel argues it can only be cleared with better context engineering, not bigger models.
"The key to accuracy is this sweet spot between just enough context so that you're not missing things, and then discoverability of context," he says. "Certainly not too much context." Creating a solution that achieves that balance is beyond the reach of anybody with just a coding agent on their desktop, he argues.
NeuBird AI's argument rests on the fact that most outages cannot be reasoned about inside a single dashboard or service. Any enterprise large enough to need an SRE team has silos throughout the tech stack, from storage and networking through to operating platforms and applications, especially after microservices fragmented the estate.
An investigation has to traverse boundaries that no single human has full visibility into , and NeuBird handles this by doing the dependency mapping before the incident starts, so that when an investigation kicks off the system already knows where to look and how the pieces relate.
Co-Pilot Now, Autonomy Later – Maybe
The clearest signal from the research, and the one Martel finds least surprising, is the preference for co-pilot models, with 62 percent wanting AI to assist rather than replace.
He recognizes this stage from his own work with coding agents, though he also acknowledges an evolutionary arc. A year ago he wouldn't walk away from a coding agent for a minute, and now he's tempted to flip it into dangerous mode and let it run. He still checks in and architects everything, though. "I'm not going to completely surrender my responsibilities," he says.
The pragmatic path he describes for operations looks similar. NeuBirdAI is starting to wire up automation through Ansible's Model Context Protocol (MCP) server, with certain playbooks marked as safe to automate in production and the rest gated behind human approval.
Adding memory to a pod up to a known ceiling is something an agent can handle; anything riskier waits for a person. How much an engineer delegates, Martel says, depends on their appetite for risk and the experience they have built up working alongside the tool.
The Five Minute Clock And The Death Of The War Room
Response time dominates the AIOps brief: just over half of survey respondents expect operational answers in under five minutes, and 75 percent want them inside ten, putting immense pressure on workflows that were never built for the cadence. Getting six specialists up to speed and pulling them onto a war-room conference bridge takes time the SLA cannot absorb.
Martel's argument is that the on-call experience has to change before the clock can. "You want to get to the situation where you're not on a call with 20 other teams. Instead you're in front of a document that's outlining the explanation of what's happening and either giving me a solution or telling me who should get involved," he says.
The agent does the legwork before the engineer logs in, so by the time the engineer arrives, the early triage questions have already been answered and only the interesting decisions remain.
What IT Means For The Observability Bill
The most provocative finding, for incumbent observability vendors at least, is that 52 percent of respondents would consider switching telemetry tools if AI-driven insights worked across any back-end.
Asked where this goes, Martel doesn't hedge. "In the future observability will be dominated by open source, cost effective storage indexing technology like Grafana, Elasticsearch, or OpenSearch."
In that scenario the strategic asset shifts from whoever hoards the most telemetry to whoever can investigate it most intelligently, which means a context engine sitting above whatever storage layer is cheapest.
This is a useful lens for buyers about to renew an observability contract, because the dashboards they have paid a fortune for are the human-readable layer of a system that increasingly has machine readers too.
The survey describes a market that badly wants AI in operations but has learned to be suspicious of vendors promising results without evidence.
Martel's pitch is that the platforms surviving the next two years will be the ones that show their work and fit into the existing change-management apparatus rather than demanding a rewrite of it. The winners will treat operational context as a first-class engineering problem rather than a prompt-stuffing exercise.
Martel has a blunt answer for SRE leads still wondering whether their team is behind the curve.
"There are advantages you will gain in terms of keeping up with a growing production estate with flat operational budgets," he says. "If you don't adopt it, what are you going to do? You are going to struggle."
AI Chips Drive Around A Third Of TSMC Revenues
SREs To AI Agents: Prove Yourself Before You Touch Production
The Aspirations Of HPE And Dell In The Quantum-Classical HPC Datacenter
Why AI-Ready Data Is The Real Advantage
QuiX Quantum Shows Off A Photonic Architecture For HPC Datacenters
As Goes AI Compute, So Goes Ethernet Networking
Of Course Meta Platforms Is Going To Be A Cloud
Three HPC Gurus Ask: Do We Still Need GPUs?
Optical Scale Up Fabrics Are Limited By Manufacturing, Not Architecture
AMD Stretches Server DRAM With Flash Extended Memory
The End Of Boom/Bust Cycles For The Memory Market
A Deep Dive On China’s “LineShine” All-CPU, Exaflops-Class Supercomputer
HPE Delivers Upgraded HPC Hardware, Software For Security, Sovereignty, And Multi-Tenancy
AMD And Nvidia Are Neck And Neck In HPC Supercomputing
enterprise
HPE Rides The Agentic AI Wave Back Into The Datacenter
store
Everpure’s AI Strategy Is Almost Purely Based On Nvidia
compute
The Server Boom Balances Price Increases Against Chip Shortages
HPE’s Datacenter Networking Picture Comes Into Clearer Focus
Quantum Pulse Does Industrial Light Magic To Deliver Massive Boost In Qubits
Tensordyne Converts AI Matrix Math To Logs To Crank Up Inference Oomph
QuEra’s Libra Fault-Tolerant Quantum System Heading To Amazon Braket Service
More Power To You – And To The Datacenters
While Oracle Will Rake In Big Bucks On AI, Profits Are Hard To Predict
AWS Tunes Up Graviton5 For Agentic AI, Boosts Bang For The Buck Bigtime
Cost Per Genomics Sample? Try Cost Per Sequencing Attempt
D-Wave Riding The Dual-Rail For Its Gate-Model Quantum Ambitions
compute
AI Chip Shepherds Broadcom And Marvell Have Skinned The Golden Fleece
AI
Chip Capacity Constraints Put A Governor On AI Spending Growth
compute
Enhanced Performance For Server Consolidation With Intel Xeon 6+
In-depth coverage of high end computing
Contact us
Advertise with us
Who we are
Newsletter
The Register
DevClass
Blocks and Files
Cookies Policy
Privacy Policy
Ts & Cs
Do not sell my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

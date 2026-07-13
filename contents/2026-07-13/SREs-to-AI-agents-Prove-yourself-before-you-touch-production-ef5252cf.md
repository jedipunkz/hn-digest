---
source: "https://www.theregister.com/ai-and-ml/2026/07/13/sres-to-ai-agents-prove-yourself-before-you-touch-production/5264658"
hn_url: "https://news.ycombinator.com/item?id=48898076"
title: "SREs to AI agents: Prove yourself before you touch production"
article_title: "SREs to AI agents: Prove yourself before you touch production"
author: "Bender"
captured_at: "2026-07-13T20:51:04Z"
capture_tool: "hn-digest"
hn_id: 48898076
score: 2
comments: 0
posted_at: "2026-07-13T20:08:01Z"
tags:
  - hacker-news
  - translated
---

# SREs to AI agents: Prove yourself before you touch production

- HN: [48898076](https://news.ycombinator.com/item?id=48898076)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/07/13/sres-to-ai-agents-prove-yourself-before-you-touch-production/5264658)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T20:08:01Z

## Translation

タイトル: SRE から AI エージェントまで: 本番環境に触れる前に自分自身を証明してください
説明: スポンサー付き特集: 696 人の専門家は、副操縦士は歓迎されるが、自動操縦はあまり好ましくないと考えています

記事本文:
メインコンテンツへジャンプ
検索
トピックス
特別な機能
すべての特別な機能
FIS と AWS による金融サービスの最新化
Capgemini と AWS でそれを実現する
Nutanix: Kubernetes をスケールします。カオスではありません。
SRE から AI エージェントまで: 本番環境に触れる前に自分自身を証明してください
スポンサー付き特集: 696 人の専門家が、副操縦士は歓迎されるが、自動操縦はあまり好ましくないと考えている
AI エージェントを信頼してダウンタイムに関するユーザーの苦情を要約することは、別のことです。無人で問題を解決してくれると信頼することは、まったく別のことです。
2026 年 4 月に The Register が NeuBird AI を使用して実施した 696 人の専門家に対する調査では、73 パーセントが AIOps をまったく使用しておらず、さらに 19 パーセントが試験運用中で、本番環境に導入しているのはわずか 8 パーセントであることが判明しました。
何が妨げになっているのかという質問では、回答者の 60% が信頼の欠如を挙げ、これが最大の問題であり、ROI、セキュリティ、データ品質に対する懸念がそれぞれ約 12 ～ 13% を占めています。
NeuBird AI の Production Ops Agent は、その信頼不足を解消するように設計されています。アラート キューを要約するのではなく、メトリクス、ログ、トレース、インフラストラクチャ テレメトリ、展開アクティビティ、依存関係の関係を継続的に関連付け、その組み合わせた全体像に対して調査を実行して、考えられる根本原因と次のアクションを提案します。ステップアップストリームでも機能します。 NeuBird AI は、ノイズの多いアラート キューに高速レスポンダーを追加するのではなく、ソースでの可観測性を修正します。エージェント インストルメンテーションを通じて適切なシグナルを生成するため、アラートは設計上高シグナルになります。マーテルが言うように、重要なのは、出力にパッチを適用することではなく、ソースでの可観測性を修正することです。
Field CTO の Francois Martel 氏が The Register にインタビューし、調査で判明したことと、AIOps の次の段階がエンジニアが 10 年間見つめてきたダッシュボードと全く似ていない理由について語りました。彼はまた、whについての見解を持っています

SRE チームがエージェントを本番システムに近づける前に、at を変更する必要があります。
関心は高いが、展開はほとんどない
データは、マーテルが現場ですでに聞いていたことを裏付けました。 「関心は高いが、行動はそれほど多くない」と彼は言う。このパターンはエージェント ワークロード全体でよく見られるものです。人気を博しているカテゴリは、コーディング エージェントやコンテンツ生成など、明らかな人間の関与と明確な検証パスが伴うものです。作業は実行環境内でエンジニアがまだ見ていないデータに基づいて行われるため、運用はさらに難しくなり、その影響は顧客向けシステムに現れます。
彼は、NeuBird に入社するずっと前から、企業内に同様のギャップを感じていました。つまり、300 件の AI 修正候補のバックログと、初期の熱狂的な混乱があり、その後、最初のバージョンが出荷されるまでに 1 年間の苦労が続きました。
ツールが期待に追いつくまで 6 か月待つことが正しい判断である場合もあるため、遅延の一部は市場開発のスピードに起因します。
もう 1 つの原因は、汎用エージェントが SRE の問題に適合しないため、ツール カテゴリの選択が間違っていることです。
マーテル氏は、「はるかに優れた仕事をして、安全性、セキュリティ、ガードレール、幻覚などの「懸念事項の一部に対処できる」専門のエージェントがいる」と語る。また、ツールはチームの既存のワークフローに適合する必要があります。
マーテル氏は、調査で表面化した信頼性を重視する懸念について反論しようとはしていない。 「AI との連携は信頼を築く作業であり、信頼を得るために AI は学習する必要があります」と彼は言います。 「これは、AI エージェントにとって一種のキラー機能だと思います。学習し、向上していることを示すことができれば、信頼を得ることができます。」
だからこそ、説明可能性がセキュリティレビューのために移植されるのではなく、NeuBird AI の設計の中心に据えられているのです。 p

latform はすべての決定の背後にある理由を記録するため、エンジニアは同僚のインシデント報告書を質問するのと同じように質問できます。 「エージェントがいる場合は常に、下された決定を監査し、決定の背後にある理由を理解できるようにしたいと考えています」とマーテル氏は言います。
NeuBird AI は内部的に、Langfuse を介してあらゆる推論ステップをキャプチャします。説明可能性は半分にすぎません。このプラットフォームは SOC 2 Type II 認定も受けており、読み取り専用で何も保存しないため、論理だけではなくアーキテクチャに信頼が組み込まれています。外部的には、より難しい問題はプレゼンテーションです。システムの初期バージョンでは、非常に多くの詳細が表面化するため、ユーザーはそれをテキストの壁のように表現しました。この修正では、推論をダンプするのではなく質問できるようにすることで、エンジニアが上級チームメイトに質問するのと同じようにシステムのメモリとチャットできるようになりました。
答えの信頼性を高めるのは文脈です
同じ調査では、回答者の 59 パーセントが採用前にほぼ完璧な精度を必要としている一方、さらに 10 人に 3 人は約 80 パーセントの精度を許容していることがわかりました。このハードルは容赦のないもので、マーテル氏は、より大きなモデルではなく、より優れたコンテキストエンジニアリングによってのみクリアできると主張しています。
「正確さの鍵は、何かを見逃さない程度の十分なコンテキストと、コンテキストの発見可能性との間のスイートスポットです」と彼は言います。 「確かに、あまり文脈はありません。」このバランスを実現するソリューションを作成することは、デスクトップ上にコーディング エージェントだけを使用している人には手の届かないものだと彼は主張します。
NeuBird AI の主張は、ほとんどの機能停止は単一のダッシュボードまたはサービス内では推論できないという事実に基づいています。 SRE チームを必要とするほどの規模の企業は、ストレージやネットワーキングからオペレーティング プラットフォームやアプリケーションに至るまで、技術スタック全体にわたってサイロを抱えています (特に 3 年以降)。

icroservices は不動産を断片化しました。
調査では、誰一人人間が完全に把握できない境界を通過する必要があります。NeuBird は、インシデントの開始前に依存関係マッピングを実行することでこれを処理します。そのため、調査が開始される時点で、システムはどこを調べればよいのか、各部分がどのように関連しているのかをすでに知っています。
今は副操縦士、後は自主性、多分
この調査から得られた最も明らかなシグナルであり、マーテル氏がそれほど驚くべきことではなかったのは、副操縦士モデルの好みであり、62% が AI に取って代わられるのではなく、AI が支援することを望んでいました。
彼はコーディング エージェントとの自身の仕事からこの段階を認識していますが、進化の弧も認めています。 1 年前、彼はコーディング エージェントから片時も離れることはありませんでしたが、今では、コーディング エージェントを危険モードに切り替えて実行させたいという誘惑に駆られています。ただし、彼は今でもチェックインし、すべてを設計しています。 「私は自分の責任を完全に放棄するつもりはない」と彼は言う。
彼が説明する運用の実際的な道筋も同様に見えます。 NeuBirdAI は、Ansible の Model Context Protocol (MCP) サーバーを介して自動化の連携を開始しており、特定の Playbook は本番環境で自動化しても安全であるとマークされ、残りは人間の承認の後にゲートされます。
既知の上限までポッドにメモリを追加することは、エージェントが処理できることです。より危険なことが人を待っています。マーテル氏によると、エンジニアがどれだけ委任するかは、エンジニアのリスクに対する欲求と、ツールと並行して作業して蓄積した経験によって決まります。
5分時計と作戦室の死
AIOps の概要で最も重要なのは応答時間です。調査回答者の半数強が運用上の回答を 5 分以内に期待しており、75% が 10 分以内を望んでおり、リズムに合わせて構築されていないワークフローに多大なプレッシャーを与えています。 6 人のスペシャリストを訓練し、作戦室のコンサルに参加させる

参照ブリッジには、SLA が吸収できない時間がかかります。
マーテル氏の主張は、時計が変わる前にオンコール エクスペリエンスを変える必要があるということです。 「他の 20 チームと電話をかけているわけではないという状況に到達したいと考えています。代わりに、何が起こっているのかを説明し、解決策を示してくれたり、誰が関与すべきかを教えてくれたりする文書の前にいるのです。」と彼は言います。
エージェントはエンジニアがログインする前に準備作業を行うため、エンジニアが到着するまでに、初期のトリアージの質問はすでに回答されており、興味深い決定だけが残っています。
可観測性法案にとって IT が意味するもの
少なくとも既存の可観測性ベンダーにとって最も刺激的な調査結果は、AI 主導の洞察がバックエンド全体で機能する場合、回答者の 52% がテレメトリ ツールの切り替えを検討するだろうということです。
これがどうなるかと問われれば、マーテル氏はヘッジをしない。 「将来的には、可観測性は、Grafana、Elasticsearch、OpenSearch などのオープンソースのコスト効率の高いストレージ インデックス作成テクノロジによって支配されることになります。」
そのシナリオでは、戦略的資産は、テレメトリを最も多く蓄積する者から、それを最もインテリジェントに調査できる者に移行します。これは、コンテキスト エンジンが最も安価なストレージ レイヤーの上に位置することを意味します。
これは、オブザーバビリティ契約を更新しようとしている購入者にとって便利なレンズです。なぜなら、購入者が大金を払って購入したダッシュボードは、機械読み取り機能も搭載するシステムの人間が読み取り可能な層であるためです。
この調査では、AI の運用を強く望んでいながらも、ベンダーが証拠なしに結果を約束することに疑念を抱くようになった市場について説明しています。
マーテル氏の売り込みは、今後 2 年間生き残るプラットフォームは、その成果を発揮し、既存の変更管理機構の書き換えを要求するのではなく、既存の変更管理機構に適合するプラットフォームになるだろう、というものだ。勝者はトレします

即時詰め込みの演習ではなく、第一級のエンジニアリング問題として運用上のコンテキストで。
マーテルは、自分のチームが時代遅れなのかどうかまだ疑問を抱いている SRE リーダーに対して、率直な答えを示しています。
「安定した運用予算で成長を続ける生産施設に対応できるという点で、メリットが得られます」と彼は言います。 「もし採用しなかったらどうするの？苦労することになるよ。」
風変わりな
Excel の競争が激化し、スプレッドシートのマニアが街頭から競争するようになる
ディルムッド・アーリー、屋外でパズルを解かなければならなかったにも関わらず、前回世界チャンピオンのディルムッド・アーリーが再び勝利
価格が間違っています: AI のコスト計算では、トークンのコストだけでなくタスクの完了率も考慮する必要があります
パートナーのコンテンツ: 社内 AI を開発する企業は、時間単位でサーバー スペースを借りるだけでなく、自社の技術を使用するための永住地を選択する必要があります。
マイクロソフトのトップがフロンティアAI研究所に敵意を示し、企業に知的財産を守るよう警告
ロックダウンするとサティア・ナデラ氏は警告、レドモンド氏が古き良き時代にOpenAIに投入した数十億ドルを忘れているようだ
コラムニスト
Microsoft はライセンス収入を守る戦いに負けつつある。その感覚に慣れた方が良い
クローン・ウォーズのリメイクの時が来た
ドイツの企業が破産を申請、6週間生産を停止したサイバー犯罪を非難
ZEGO-TVZ、3月のサイバー攻撃による経済的損失を受けて扉を閉めるのが唯一の選択肢だと語る
aiとml
銀行やハイパースケーラーさえも今ではAIバブルについて警鐘を鳴らしている
オフビート
C プログラマーが可読性に対して新たな犯罪を犯す
ソフトウェア
中古ソフトウェアライセンス争いで裁判所がマイクロソフトの上訴を棄却
セキュリティ
GitHub AI エージェントがうまく質問するとプライベート リポジトリを漏洩する
アプリケーション
Microsoft、約20年ぶりにOWA Lightを終了へ
ロックダウンしろとサティア・ナデラ氏が警告、数十億のレドモンドチップを忘れているようだ

古き良き時代に OpenAI を体験してみませんか
RAMポカリプスはAI黙示録の前兆である可能性がある
OpenAI と Anthropic は AI スイス アーミー ナイフを開発しましたが、将来はより小型の専用ツールになる可能性があります
同社は環境に貢献したいと考えながらも、AI をすべての人に強制したいという苦境に直面している
プラス：米国はイランのプロパガンダサイトを削除。マーケティング会社が「なぜあなたの情報を入手するのですか？」と尋ねます。さらに！
プラス：中国はスマートフォン監視ツールをアップグレード。リングは覗き見防止の姿勢を緩和します。などなど
ジェフ・モス氏によると、投票村のレポートは非常に成功しており、今後は DEF CON 全体が含まれることになる
会社全体の評価額は35億ドル以上に相当するが、売却部分は特定されていない
プラスの面としては、情報セキュリティは長く安定したキャリアを築くのに適しています。
Cinnamon 6.8 は Wayland をサポートします – 必要に応じて
Linux Mint デスクトップの次期バージョンには両方の種類のディスプレイ サーバーが搭載されています
KDE Plasma ユーザーは恐ろしい変化の予兆に直面しています: 6.6.6 の登場
現在は 6.7 が最新版であり、6.8 では好むと好まざるにかかわらず Wayland を入手できるようになります。
FOSS クラウドオフィススイート間の競争が激化する中、Collabora が CODE 26.04 をリリース
Markdown サポートとよりスマートな数式エラー処理に加え、AI が統合されました。

[切り捨てられた]

## Original Extract

SPONSORED FEATURE: 696 experts find co-pilot welcome, autopilot not so much

Jump to main content
Search
TOPICS
Special Features
All Special Features
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
SREs to AI agents: Prove yourself before you touch production
SPONSORED FEATURE: 696 experts find co-pilot welcome, autopilot not so much
Trusting an AI agent to summarize user complaints about downtime is one thing; trusting it to fix the problem unattended is something else entirely.
A survey of 696 experts The Register ran with NeuBird AI in April 2026 found that 73 percent are not using AIOps at all, another 19 percent are in pilot, and only eight percent have it in production.
Asked what's stopping them, 60 percent of respondents cited a lack of trust, by far the biggest issue, with concerns about ROI, security and data quality each registering at around 12 to 13 percent.
NeuBird AI's Production Ops Agent is designed to close that trust deficit. Rather than summarizing the alert queue, it continuously correlates metrics, logs, traces, infrastructure telemetry, deployment activity and dependency relationships, then runs investigations across that combined picture to suggest probable root causes and next actions. It also works a step upstream. Rather than bolting a faster responder onto a noisy alert queue, NeuBird AI fixes observability at its source: through agentic instrumentation it generates the right signals, so the alert is high-signal by design. As Martel puts it, the point is to fix observability at the source, not patch the output.
Field CTO Francois Martel sat down with The Register to talk through what the survey found, and why the next phase of AIOps will look nothing like the dashboards engineers have stared at for a decade. He also has views on what must change before SRE teams will let agents near their production systems.
Lots of interest, very little deployment
The data confirmed what Martel was already hearing in the field. "There's a lot of interest, but not a lot of action," he says. The pattern is familiar across agentic workloads: the categories that have taken off are the ones that come with an obvious human in the loop and an obvious verification path, such as coding agents and content generation. Operations is harder, because the work happens inside the running environment, on data the engineer hasn't seen yet, with consequences that show up in customer-facing systems.
He saw the same gap inside enterprises long before he joined NeuBird: a backlog of 300 candidate AI fixes and a flurry of early enthusiasm, followed by a year of slog before the first one shipped.
Part of that delay comes down to the speed of market development, since waiting six months for the tools to catch up with your expectations is sometimes the right call.
Another part of it is the wrong choice of tool category, because general-purpose agents do not fit SRE problems.
"There are specialized agents that can do a much better job," Martel says, "and address some of the concerns" of safety, security, guardrails and hallucinations. The tool also has to fit into the team's existing workflows.
Martel doesn't try to argue with the trust-heavy concerns the survey surfaced. "Working with AI is a trust-building exercise, and AI has to learn in order to gain trust," he says. "I would say that's kind of the killer feature for AI agents. If you can show that you're learning and getting better, then you can gain trust."
That's why explainability sits at the center of NeuBird AI's design rather than being grafted on for the security review. The platform records the reasoning behind every decision so an engineer can interrogate it the way they'd interrogate a colleague's incident report. "Whenever you have an agent, you want to be able to audit the decisions that were made, and understand the reasoning behind the decision," Martel says.
Internally, NeuBird AI captures every reasoning step via Langfuse. Explainability is only half of it. The platform is also SOC 2 Type II certified, read-only, and stores nothing, so trust is built into the architecture, not just the reasoning. Externally, the harder problem is presentation: early versions of the system surfaced so much detail that users described it as a wall of text. The fix was to make the reasoning interrogable rather than dumped, so engineers can chat with the system's memory the way they'd query a more senior teammate.
Context is what makes the answer credible
The same survey found that 59 percent of respondents require near-perfect accuracy before they'll adopt, while another three in every ten will tolerate around 80 percent accuracy. That bar is unforgiving, and Martel argues it can only be cleared with better context engineering, not bigger models.
"The key to accuracy is this sweet spot between just enough context so that you're not missing things, and then discoverability of context," he says. "Certainly not too much context." Creating a solution that achieves that balance is beyond the reach of anybody with just a coding agent on their desktop, he argues.
NeuBird AI's argument rests on the fact that most outages cannot be reasoned about inside a single dashboard or service. Any enterprise large enough to need an SRE team has silos throughout the tech stack, from storage and networking through to operating platforms and applications, especially after microservices fragmented the estate.
An investigation has to traverse boundaries that no single human has full visibility into , and NeuBird handles this by doing the dependency mapping before the incident starts, so that when an investigation kicks off the system already knows where to look and how the pieces relate.
Co-pilot now, autonomy later, maybe
The clearest signal from the research, and the one Martel finds least surprising, is the preference for co-pilot models, with 62 percent wanting AI to assist rather than replace.
He recognizes this stage from his own work with coding agents, though he also acknowledges an evolutionary arc. A year ago he wouldn't walk away from a coding agent for a minute, and now he's tempted to flip it into dangerous mode and let it run. He still checks in and architects everything, though. "I'm not going to completely surrender my responsibilities," he says.
The pragmatic path he describes for operations looks similar. NeuBirdAI is starting to wire up automation through Ansible's Model Context Protocol (MCP) server, with certain playbooks marked as safe to automate in production and the rest gated behind human approval.
Adding memory to a pod up to a known ceiling is something an agent can handle; anything riskier waits for a person. How much an engineer delegates, Martel says, depends on their appetite for risk and the experience they have built up working alongside the tool.
The five-minute clock and the death of the war room
Response time dominates the AIOps brief: just over half of survey respondents expect operational answers in under five minutes, and 75 percent want them inside ten, putting immense pressure on workflows that were never built for the cadence. Getting six specialists up to speed and pulling them onto a war-room conference bridge takes time the SLA cannot absorb.
Martel's argument is that the on-call experience has to change before the clock can. "You want to get to the situation where you're not on a call with 20 other teams. Instead you're in front of a document that's outlining the explanation of what's happening and either giving me a solution or telling me who should get involved," he says.
The agent does the legwork before the engineer logs in, so by the time the engineer arrives, the early triage questions have already been answered and only the interesting decisions remain.
What IT means for the observability bill
The most provocative finding, for incumbent observability vendors at least, is that 52 percent of respondents would consider switching telemetry tools if AI-driven insights worked across any back-end.
Asked where this goes, Martel doesn't hedge. "In the future observability will be dominated by open source, cost effective storage indexing technology like Grafana, Elasticsearch, or OpenSearch."
In that scenario the strategic asset shifts from whoever hoards the most telemetry to whoever can investigate it most intelligently, which means a context engine sitting above whatever storage layer is cheapest.
This is a useful lens for buyers about to renew an observability contract, because the dashboards they have paid a fortune for are the human-readable layer of a system that increasingly has machine readers too.
The survey describes a market that badly wants AI in operations but has learned to be suspicious of vendors promising results without evidence.
Martel's pitch is that the platforms surviving the next two years will be the ones that show their work and fit into the existing change-management apparatus rather than demanding a rewrite of it. The winners will treat operational context as a first-class engineering problem rather than a prompt-stuffing exercise.
Martel has a blunt answer for SRE leads still wondering whether their team is behind the curve.
"There are advantages you'll gain in terms of keeping up with a growing production estate with flat operational budgets," he says. "If you don't adopt it, what are you going to do? You're going to struggle."
offbeat
Excel competition goes extreme, makes spreadsheet geeks compete from the street
Defending world champ Diarmuid Early wins again despite being forced to solve puzzles outdoors
The price is wrong: AI cost calculation has to consider task completion rates, not just token costs
PARTNER CONTENT: Firms crafting internal AI must choose a permanent residence for their tech, not just rent server space by the hour.
Microsoft chief turns hostile on frontier AI labs, warns companies to guard their IP
Lock it down, warns Satya Nadella, seemingly forgetting the billions Redmond chipped in to OpenAI back in the good old days
columnists
Microsoft is losing the battle to protect license lucre. It better get used to the feeling
Time for the Clone Wars remake
German firm files for insolvency, blames cybercrims who shut down production for 6 weeks
ZEGO-TVZ says the financial fallout from a March cyberattack left shutting its doors as the only option
ai and ml
Even banks and hyperscalers are now sounding the alarm about the AI bubble
OFFBEAT
C programmers commit fresh crimes against readability
software
Court tosses Microsoft's appeal in pre-owned software licenses battle
Security
GitHub AI agent leaks private repos when asked nicely
applications
Microsoft to switch off OWA Light after nearly two decades
Lock it down, warns Satya Nadella, seemingly forgetting the billions Redmond chipped in to OpenAI back in the good old days
The RAMpocalypse may be the precursor to the AIpocalypse
OpenAI and Anthropic have built AI Swiss Army Knives, but the future may be smaller built-for-purpose tools
Firm faces quandary of wanting to help the environment, but also wanting to force AI on everyone
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Cinnamon 6.8 will support Wayland – if you want it
Next version of Linux Mint’s desktop has both kinds of display server
KDE Plasma users face a dire omen of change: 6.6.6 arrives
6.7 is now current, and in 6.8 you're getting Wayland whether you like it or not
Collabora releases CODE 26.04 as rivalry between FOSS cloudy office suites heats up
Now with Markdown support and smarter formula error handling – plus integrated AI, though it

[truncated]

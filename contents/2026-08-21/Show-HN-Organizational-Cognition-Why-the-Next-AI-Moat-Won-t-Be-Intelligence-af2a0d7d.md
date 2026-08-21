---
source: "https://chinmayrelkar.github.io/writing/organizational-cognition.html"
hn_url: "https://news.ycombinator.com/item?id=49388709"
title: "Show HN: Organizational Cognition: Why the Next AI Moat Won't Be Intelligence"
article_title: "Organizational Cognition · Chinmay Relkar"
image: "https://chinmayrelkar.github.io/assets/og-organizational-cognition.jpg"
author: "Chnmy"
captured_at: "2026-08-21T15:25:17Z"
capture_tool: "hn-digest"
hn_id: 49388709
score: 1
comments: 0
posted_at: "2026-08-21T14:28:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Organizational Cognition: Why the Next AI Moat Won't Be Intelligence

- HN: [49388709](https://news.ycombinator.com/item?id=49388709)
- Source: [chinmayrelkar.github.io](https://chinmayrelkar.github.io/writing/organizational-cognition.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T14:28:55Z

## Translation

タイトル: Show HN: 組織の認知: 次の AI の堀はなぜインテリジェンスにならないのか
記事のタイトル: 組織認知 · Chinmay Relkar
説明: モデルと市場に関して研究室の意見が一致していません。彼らはメモリ、コネクタ、ガバナンス、エージェント ID について合意し、インターフェイスを譲渡します。それは組織の認知度への賭けだ。

記事本文:
コンテンツにスキップ
執筆
組織の認知
エッセイ 2026 年 7 月 ~15 分
なぜすべての AI ラボが同じレイヤーを構築し、インターフェイスを提供しているのか。
モデルと市場については研究室の意見が分かれています。彼らは、メモリ、コネクタ、プロジェクト、カスタム命令、エージェント フレームワーク、エンタープライズ アイデンティティ、ガバナンスなどの蓄積面について合意しています。これとは別に、ライバルが使用できる中立的な基盤の下で、コネクタとプロトコル インターフェイス (MCP、A2A、AGENTS.md) をオープンします。
それは製品ロードマップではありません。それは資産が何であるかについての賭けです。
Anthropic は MCP を寄付します。 Google は A2A を寄付します。 OpenAI は AGENTS.md に貢献しています。 Microsoft、AWS、Cloudflare、Bloomberg は、Linux Foundation の下で中立的な基盤プロジェクトを支援しています。 3 10 Salesforce は、競合他社のモデルと競合他社のプロトコルに基づいて Slackbot を再構築します。 1 16 チーム判定ファイルは、ベンダー中立の形式 ( AGENTS.md ) に収束しており、現在では基本的にすべての本格的なコーディング エージェントがこの形式を読み取っています。 14
既存の説明はこの部分を説明しています。すべてを説明できるものはありません。欠落しているカテゴリ: パイプが解放されたら誰もが蓄積しようとしているもの。
2. 4 つの説明、4 つのギャップ
それが機能する場所。 2 年間、最も賢いモデルが勝つという単純な理論が支配的でした。トレーニング費用、ベンチマーク文化、企業がプロバイダーを選択する方法について説明しました。
壊れるところ。これは、モデル アーキテクチャとは何の関係もなく、各フロンティア ラボが構築しているインフラストラクチャについては説明しません。よりスマートなモデルで十分であれば、Anthropic は MCP を必要とせず、OpenAI はコネクタを必要とせず、Google は Workspace 内に Gemini を必要としません。それらの製品は、インテリジェンスとは別の問題を解決します。モデルリーダーシップでは、なぜリーダーが周囲の層を最速で構築しているのか説明できません。
それが機能する場所。エージェントのフレームワーク、ツールの使用、および信頼性

案件。
壊れるところ。あるプラットフォーム上のワークフローは、別のプラットフォーム上で書き換えることができます。より難しい問題は、AI が何をするかではなく、特定の組織内で何をすべきかを AI がどのように知るか、つまりどのワークフローをいつ、どのパラメータで、誰の権限で呼び出すかということです。ワークフローは判断ではなく実行を説明します。
それが機能する場所。 Microsoft と Google は、Copilot と Gemini が社内ソフトウェアとして提供されており、企業がすでに費用を支払っているため、モデルの品質をいち早く追随することができます。
壊れるところ。ディストリビューションは、企業に到達した後に何を構築するかではなく、どのようにして企業に到達するかを説明します。 Microsoft の最初の Copilot ウェーブは、Office にチャットが組み込まれました。次に、グラフのグラウンディング、カスタム エージェント、および人間以外のアクターのアイデンティティが続きました。 Salesforce は、すべての新しいアカウントに Slack をバンドルし、2026 年に MCP と Claude を介して Slackbot を他の全員のシステムに接続しました。 1 配布こそが利益であると信じている企業は、競合他社のモデルやプロトコルに依存するために資金を費やすことはありません。
それが機能する場所。言語モデルは各リクエストをゼロから開始します。メモリ、コネクタ、プロジェクト、カスタム命令は、キャッシュ、取得、プロンプト プレフィックスなどのエンジニアリング パッチのように見えます。同じ制約、同じ補償機構。新しいカテゴリーは必要ありません。
壊れるところ。ステートレス性は、メモリとコネクタが存在する理由を説明します。記憶されているものの形を説明するものではありません。製品は一般的なリコールに収束していません。それらは組織固有の判断に収束しつつあります。つまり、これをどのレビュー担当者に転送するか、このチームが「完了」とは何を意味するか、2 つの相反する文書のどちらが権威があるか、この会社がどの程度のリスクを許容するかなどです。それはリポジトリ内の事実ではありません。組織がどのように行動するかから推測し、モデルではない場所に保存する必要があります。
同アカウントはまた、ウィンドウが拡大するにつれて、周囲のインフラストラクチャが増加すると予測しています。

構造は先細になるはずです。 Windows は桁違いに成長しました。周囲のインフラストラクチャは急速に成長しています。 2026 年 7 月、MCP の最大の改訂により、コア プロトコルがよりステートレスになり、サーバーがサーバーレスおよびエッジで実行できるようになり、長時間実行されるタスクが拡張機能に移動されました。 2 周囲の機械がモデルにメモリを与えるために存在する場合、共有プロトコルはそのメモリを配置する場所になります。代わりに、プロトコルはそれを排除しました。メモリはプラットフォームと製品にまで浸透しました。
監査証跡、どのエージェントが誰の権限に基づいて行動するかの権限モデル、システム プロンプトを変更するための承認パスなど、これらはいずれも有限のコンテキスト ウィンドウを補うことはできません。無制限のウィンドウを持つ完全にステートフルなモデルでも、同様にそれらが必要になります。ステートレス性は、内容ではなくメカニズムを説明します。
モデル、ワークフロー、分散、ステートレス性はそれぞれ、状況の一部を説明します。すべてのプレイヤーが同じ能力に集中する理由、または集められているものが組織固有の事実ではなく組織固有の判断である理由を説明するものはありません。
その残余がこのエッセイの主張である。以下の製品に関する決定事項が遵守されます。それらの読みは解釈されます。セクション 7 の予測は予測されています。レンズは収束を組織します。それが必要であるかどうかはまだ示されていません。競合するミラーリングと調達チェックリストによって、同様のサーフェスが作成される可能性があります。獲得しなければならないのは、これらのサーフェスの内容が判断であり、SSO を備えた一般的な RAG ではないということです。
組織認知とは、情報を認識し、意思決定を行い、行動を調整し、継続的に学習する組織の集合的な能力であり、現在、正式な成果物ではなく、主に人間の相互作用の中に存在するパターンを通じて表現されます。
個人の IQ ではなく、全体的な IQ。 4つの機能：知覚（何に気づくか）、決定

（何が選ばれるか）、調整（仕事がどのように調整されるか）、学習（組織が改善されるかどうか）。境界: その会社で働いたことのない人でも書き留めて実行できるのであれば、それは知識です。不文律や暗黙の判断が必要な場合、それは認知です。
この現象は新しいものではありません。組織科学は、ルーチンと意思決定ルール、暗黙知、センスメイキング、分散認知など、組織科学の一部を何十年にもわたって説明してきました。 「制度上の知識」は、知っていることを過度に強調し、処理することを過小評価しています。文化は、承認経路やトレードオフのヒューリスティックではなく、リスク許容度とコミュニケーションのパターンを捉えます。
新しいのは、これらのパターンを大規模に蓄積し、管理し、製品化できる基板です。最近まで、組織の認知はほぼ完全に人々の中に存在していました。記録システムとして構築するものはほとんどなかったため、製品言語の名前が定着することはありませんでした。
輸出テストです。それをファイルとしてエクスポートし、別の組織にインポートでき、そこで同じ決定が生成される場合、それはデータであり、組織の認識ではありません。
外部化は全か無かではありません。スタイル ガイド、ランブック、エスカレーション マトリックスは、判断の薄い静的なスライスをキャプチャします。それらに含まれていないのは、ライブプロセスです。つまり、2 つのルールが矛盾した場合にどちらのルールが勝つか、ルールが古くなっていることに気づき、一度ルールを破ってから正しいかということです。成果物はデータです。執筆し、改訂し、いつ逸脱すべきかを知ることは認識です。
4 つの機能、2 つのコントロール サーフェス。製品はこれらに基づいて個別に構築されるため、6 つの次元として表示されます。記憶とは、時間を超えて認識されるものであり、何が保持され、再び浮かび上がってくるのかということです。推理が決めるのです。 「コーディネーション」と「ラーニング」の名前はそのままです。委任とガバナンスは追加の機能ではありません。彼らは 4 つを実行する権限を分配し、監査します。
2 つ

同じアーカイブを持つ組織であっても、どのような過去のコンテキストがキュレーションされ展開されるかによって、まったく異なる場合があります。その違いは認知的なものです。ガバナンスとは、「AIが私たちの働き方から学習する」ことと、「AIが説明できない、または上書きできない決定を下す」ことの違いです。
このレンズが役立つのであれば、他の枠組みでは無関係に見える製品決定にも意味が生じるはずです。
フリーパイプ: MCP と基礎
Anthropic の MCP は、組織の生活が文書化され、議論され、実行されるシステムに AI を接続するためのオープン プロトコルです。 2025 年 12 月、Anthropic、Block、OpenAI は Linux Foundation の Agentic AI Foundation を共同設立し、Google、Microsoft、AWS、Cloudflare、Bloomberg の支援を受けて MCP、goose、AGENTS.md に貢献しました。 3 Google はすでに、2025 年 6 月に、エージェント間プロトコルである A2A を別の Linux Foundation プロジェクトに寄付していました。 1 周年までに 150 を超える支援組織があり、主要なクラウド全体でネイティブ サポートを獲得しました。 10 11
競合他社を含む財団に賭けを渡すことは、その賭けがインターフェースの所有にまったくなかった場合にのみ合理的です。 1 つの研究室によって管理されているコネクタ標準の価値は、誰もが信頼するコネクタ標準よりも価値がありません。製品の機能に関して奇妙な譲歩。配管については明らかです。解釈: 組織の認知が資産である場合、パイプは無料であるべきであり、パイプを通じて最も多くのコンテキストを蓄積した人がとにかく勝ちます。
2026 年 7 月の MCP リビジョンでは、この分割が強化されます。つまり、規模を拡大するためによりステートレスなコアになります。タスクと長時間実行される作業は拡張機能にプッシュされます。メモリとポリシーには製品層の懸念が残ります。 2 インターフェイスが標準化されます。アキュムレーションは独自性を維持します。
ファイルとループ: カーソルと AGENTS.md
.cursorrules は、チームの判断 (スタイル、アーキテクチャ、ライブラリの選択、テストの期待) をエンコードします。

コーディングエージェント。これらは事実ではありません。同じ要件と異なるルールを持つ 2 つのチームが、異なるソフトウェアを作成します。
エクスポート テストの内容からすると、プレーン テキスト ファイルは反例のように見えます。チームのファイルを別のチームに送信すると、どの行に負荷がかかるのか、停止後に書き込まれたのか、いつルールを一時停止する必要があるのか​​などの理由を指定せずにルールを送信したことになります。あるチームのファイルを別のリポジトリにドロップすると、そのチームのコードではなく、スナップショットに準拠することになります。転送されないのは、スナップショットを生成したループ (フリクション、レビュー、オーバーライド、その後の削除) です。
ファイルはデータです。ループは認知です。その後、業界は実験を実行しました。OpenAI によって公開され、Agentic AI Foundation に渡された AGENTS.md は、現在、本質的にすべての本格的なコーディング エージェントによってネイティブに読み取られるようになりました。 3 14 何万ものリポジトリが 1 つの構文に収束し、さまざまなソフトウェアが生成され続けました。成果物が資産である場合、それを標準化するとチームはフラット化されるでしょう。そうではありませんでした。 MCP や A2A と同じ動きで、1 層下です。誰もコンテナを守りません。
ガバナンスが伝える: Microsoft Work IQ とエージェント ID
Microsoft のエントリ ポイントはグラフ (電子メール、ドキュメント、会議、チーム) であり、誰が誰と話し、どのドキュメントが重要であるかをすでに部分的に表現しています。 2026 年 6 月のビルドでは、データ、メモリ、推論として構成された、Microsoft 365 上の名前付きインテリジェンス レイヤーとして Work IQ が出荷されました。 5 これに加えて、Copilot Studio の Work IQ、6 および 2026 年 4 月から一般提供される Entra Agent ID を通じて、エージェントがどの MCP サーバーにアクセスできるかを管理者が制御します。これにより、人間の資格情報を借用するのではなく、ファーストクラスのディレクトリ ID がエージェントに提供されます。 7
データ、メモリ、および推論の命名は、分類法の収束的な証拠であり、その証明ではありません。ネーミングは安っぽい。そうでないもの

安いのはガバナンス製品の形状です。人間以外のアクターの ID と、彼らがどのツールを呼び出すかを管理者が制御する機能は、エージェントが実際の権限に基づいて行動し、誰かがそれに代わって応答する必要がある場合に構築する機能です。人間以外のアイデンティティの数は、方法に依存する比率ですでに人間のアイデンティティを大きく上回っており、ほとんどの組織は、制作エージェントが何をしているのかについて信頼できる見解を報告していません。 8 正確な比率は一致しません。方向はそうではありません。
同じパターンですが、エントリーポイントが異なります。 OpenAI の Workspace Agent は、ビジネス アカウント向けのカスタム GPT を置き換えました。チーム所有、クラウド実行、会話中に修正可能、ユーザーごとおよびエージェントごとのメモリは意図的に 1 つの企業の頭脳にプールされていません。 4 Google は、Workspace を情報は豊富だがコンテキストが乏しいものとしてフレーム化し、エージェント セッションとメモリ バンクを提供します。 9 A2A を販売するのではなく寄付します。 10 11 Palantir は、オントロジーを、誰がアクションを承認したか、それにどれくらいのコストがかかるか、そしてその出力が信頼できるかどうかを制御するコントロール プレーンとして扱っていますが、需要は増大しており、その評価には未だに論争が続いています。 12 13 Glean の権限を意識したグラフは、どのドキュメントが存在するかだけでなく、誰が何を知っているかをエンコードします。その評価は、フロンティアモデルやオフィス規模の配分を必要とせず、蓄積されたコンテキストに基づいています。

[切り捨てられた]

## Original Extract

Labs disagree on models and markets. They agree on memory, connectors, governance, and agent identity, then give the interfaces away. That is a bet on organizational cognition.

Skip to content
Writing
Organizational Cognition
Essay July 2026 ~15 min
Why every AI lab is building the same layer, and giving the interfaces away.
Labs disagree on models and markets. They agree on the accumulation surfaces: memory, connectors, projects, custom instructions, agent frameworks, enterprise identity, and governance. Separately, they open the connector and protocol interfaces (MCP, A2A, AGENTS.md ) under neutral foundations their rivals can use.
That is not a product roadmap. It is a bet on what the asset is.
Anthropic donates MCP. Google donates A2A. OpenAI contributes AGENTS.md . Microsoft, AWS, Cloudflare, and Bloomberg back neutral foundation projects under the Linux Foundation. 3 10 Salesforce rebuilds Slackbot on a competitor's model and a competitor's protocol. 1 16 Team judgment files converge on a vendor-neutral format ( AGENTS.md ) that essentially every serious coding agent now reads. 14
Existing explanations account for pieces of this. None account for all of it. Missing category: what everyone is trying to accumulate once the pipes are free.
2. Four Explanations, Four Gaps
Where it works. For two years the dominant theory was simple: the smartest model wins. It explained training spend, benchmark culture, and how enterprises picked providers.
Where it breaks. It does not explain the infrastructure every frontier lab is building that has nothing to do with model architecture. If smarter models were sufficient, Anthropic would not need MCP, OpenAI would not need Connectors, and Google would not need Gemini inside Workspace. Those products solve a different problem than intelligence. Model leadership would not explain why leaders are building the surrounding layer fastest.
Where it works. Agent frameworks, tool use, and reliability matter.
Where it breaks. A workflow on one platform can be rewritten on another. The harder problem is not what the AI does, but how it knows what to do in a specific organization: which workflow to invoke, when, with what parameters, under whose authority. Workflows explain execution, not judgment.
Where it works. Microsoft and Google can be fast-followers on model quality because Copilot and Gemini ship inside software enterprises already pay for.
Where it breaks. Distribution explains how you reach an enterprise, not what you build once you arrive. Microsoft's first Copilot wave was chat bolted onto Office; what followed was Graph grounding, custom agents, and identity for non-human actors. Salesforce, with Slack bundled into every new account, spent 2026 wiring Slackbot into everyone else's systems over MCP and Claude. 1 A company that believed distribution was the prize would not spend it becoming dependent on rivals' models and protocols.
Where it works. Language models start each request from zero. Memory, connectors, projects, and custom instructions look like engineering patches: caches, retrieval, prompt prefixes. Same constraint, same compensating machinery. No new category required.
Where it breaks. Statelessness explains why memory and connectors exist. It does not explain the shape of what is remembered. Products are not converging on generic recall. They are converging on organization-specific judgment: which reviewer this routes to, what this team means by "done", which of two conflicting documents is authoritative, how much risk this company tolerates. That is not a fact in a repository. It has to be inferred from how the organization behaves, then stored somewhere that is not the model.
The account also predicts that as windows grow, surrounding infrastructure should taper. Windows have grown by orders of magnitude. The infrastructure around them has grown faster. In July 2026, MCP's largest revision made the core protocol more stateless so servers could run on serverless and edge, and moved long-running Tasks into an extension. 2 If the surrounding machinery existed to give models memory, the shared protocol is where you would put that memory. Instead the protocol shed it. Memory went up into platforms and products.
Audit trails, permission models for which agent may act on whose authority, approval paths for changing a system prompt: none of these compensate for a finite context window. A perfectly stateful model with an unlimited window would need them just as much. Statelessness explains the mechanism, not the content.
Models, workflows, distribution, and statelessness each explain part of the landscape. None explain why every player converges on the same capabilities, or why what is being assembled is organization-specific judgment rather than organization-specific facts.
That residual is the claim of this essay. Product decisions below are observed ; the reading of them is interpreted ; forecasts in section 7 are predicted . The lens organizes the convergence. It is not yet shown to be necessary: competitive mirroring and procurement checklists can produce similar surfaces. What it has to earn is that the content of those surfaces is judgment, not generic RAG with SSO.
Organizational cognition is the collective capacity of an organization to perceive information, make decisions, coordinate action, and continuously learn, as expressed through patterns that currently live mainly in human interaction, not in formal artifacts.
Collective, not individual IQ. Four functions: perception (what gets noticed), decision (what gets chosen), coordination (how work aligns), learning (whether the organization improves). Boundary: if it can be written down and executed by someone who has never worked at the company, it is knowledge. If it requires unwritten norms or tacit judgment, it is cognition.
The phenomenon is not new. Organizational science has described pieces of it for decades: routines and decision rules, tacit knowledge, sensemaking, distributed cognition. "Institutional knowledge" overemphasizes knowing and underemphasizes processing. Culture captures risk tolerance and communication patterns, not approval paths or trade-off heuristics.
What is new is a substrate that can accumulate, govern, and productize those patterns at scale. Until recently, organizational cognition lived almost entirely in people. There was little to build as a system of record, so product language never settled on a name.
The export test. If you can export it as a file, import it into another organization, and it produces the same decisions there, it is data, not organizational cognition.
Externalization is not all-or-nothing. Style guides, runbooks, and escalation matrices capture thin static slices of judgment. What they do not carry is the live process: which rule wins when two conflict, noticing that a rule is stale, breaking it once and being right. The artifact is data. Authoring, revising, and knowing when to deviate is cognition.
Four functions, two control surfaces. Products are built against these separately, which is why they show up as six dimensions. Memory is perception across time: what is kept and re-surfaced. Reasoning is decide. Coordination and Learning keep their names. Delegation and Governance are not extra functions; they distribute and audit the authority to perform the four.
Two organizations with identical archives can differ entirely on what past context gets curated and deployed. That difference is cognitive. Governance is the difference between "the AI learned from how we work" and "the AI makes decisions we cannot explain or override."
If the lens is useful, it should make sense of product decisions that look unrelated under other frameworks.
The free pipe: MCP and the foundations
Anthropic's MCP is an open protocol for connecting AI to the systems where organizational life is documented, discussed, and executed. In December 2025 Anthropic, Block, and OpenAI co-founded the Linux Foundation's Agentic AI Foundation, contributing MCP, goose, and AGENTS.md , with backing from Google, Microsoft, AWS, Cloudflare, and Bloomberg. 3 Google had already donated A2A, its agent-to-agent protocol, to a separate Linux Foundation project in June 2025; by its first anniversary it had over 150 supporting organizations and native support across major clouds. 10 11
Handing a bet to a foundation that includes your competitors is only rational if the bet was never on owning the interface. A connector standard controlled by one lab is worth less than one everyone trusts. Strange concession about a product feature; obvious about plumbing. Interpreted: if organizational cognition is the asset, the pipe should be free, and whoever accumulates the most context through it wins anyway.
The July 2026 MCP revision reinforces the split: a more stateless core for scale; Tasks and long-running work pushed to extensions; memory and policy remaining product-layer concerns. 2 Interfaces standardize. Accumulation stays proprietary.
The file versus the loop: Cursor and AGENTS.md
.cursorrules encodes team judgment (style, architecture, library choices, testing expectations) into a coding agent. These are not facts. Two teams with the same requirements and different rules produce different software.
By the letter of the export test, a plain text file looks like a counterexample. Send a team's file to another team and you have sent the rules without the reasons: which lines are load-bearing, which were written after an outage, when a rule should be suspended. Drop one team's file into another repository and you get compliance with a snapshot, not that team's code. What does not travel is the loop that produced the snapshot: friction, review, override, later deletion.
The file is data. The loop is cognition. The industry then ran the experiment: AGENTS.md , published by OpenAI and handed to the Agentic AI Foundation, is now read natively by essentially every serious coding agent. 3 14 Tens of thousands of repositories converged on one syntax and continued producing different software. If the artifact were the asset, standardizing it would have flattened teams. It did not. Same move as MCP and A2A, one layer down: nobody defends the container.
The governance tell: Microsoft Work IQ and agent identity
Microsoft's entry point is the Graph (email, documents, meetings, Teams), already a partial representation of who talks to whom and what documents matter. At Build in June 2026 it shipped Work IQ as a named intelligence layer over Microsoft 365, framed as Data, Memory, and Inference. 5 Alongside it: admin control over which MCP servers agents may reach through Work IQ in Copilot Studio, 6 and Entra Agent ID, generally available since April 2026, giving agents first-class directory identities rather than borrowed human credentials. 7
Naming Data, Memory, and Inference is convergent evidence for the taxonomy, not proof of it. Naming is cheap. What is not cheap is the shape of the governance products. Identity for non-human actors and admin control over which tools they may call are features you build once agents act on real authority and somebody has to answer for it. Non-human identities already outnumber human ones by large, method-dependent ratios, and most organizations report no reliable view of what production agents are doing. 8 Exact ratios disagree. Direction does not.
Same pattern, different entry points. OpenAI's Workspace Agents replaced Custom GPTs for business accounts: team-owned, cloud-running, mid-conversation correctable, with per-user and per-agent memory deliberately not pooled into one company brain. 4 Google frames Workspace as information-rich but context-poor and ships agent sessions and memory banks; 9 it donates A2A rather than selling it. 10 11 Palantir treats the ontology as a control plane for who authorized an action, what it cost, and whether the output can be trusted, with demand growing and valuation still contested. 12 13 Glean's permissions-aware graph encodes who may know what, not only which document exists; its valuation rests on accumulated context without a frontier model or Office-scale distribution to h

[truncated]

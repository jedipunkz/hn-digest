---
source: "https://alloomi.ai/blogs/the-new-frontier-of-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=49345218"
title: "AI agent that learns from real-world experiences"
article_title: "The New Frontier of AI Agents: Self-Evolving from Real-World Experiences | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences insights | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences calm communication | The New Frontier of AI Agents: Self-Evolving fr\n[truncated]"
image: "https://alloomi.ai/img/blogs/the-new-frontier-of-ai-agents/hero.png"
author: "lauraserein"
captured_at: "2026-08-18T13:35:53Z"
capture_tool: "hn-digest"
hn_id: 49345218
score: 1
comments: 0
posted_at: "2026-08-18T13:18:09Z"
tags:
  - hacker-news
  - translated
---

# AI agent that learns from real-world experiences

- HN: [49345218](https://news.ycombinator.com/item?id=49345218)
- Source: [alloomi.ai](https://alloomi.ai/blogs/the-new-frontier-of-ai-agents)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T13:18:09Z

## Translation

タイトル: 実世界の経験から学習する AI エージェント
記事のタイトル: AI エージェントの新たなフロンティア: 実世界の経験から自己進化する | AI エージェントの新たなフロンティア: 実世界の経験から得た自己進化の洞察 | AI エージェントの新境地: 実世界の経験から自己進化する穏やかなコミュニケーション | AI エージェントの新たなフロンティア: 自己進化する AI エージェント
[切り捨てられた]
説明: 実際の作業から実際の能力へ - Alloomi が 4 層フライホイールを通じてプロの経験をモデルの重み付けにどのように変換するか。 | AIエージェントの新境地 穏やかなコミュニケーション | AI 受信トレイのベスト プラクティスの概要 |プライバシー最優先のカスタマー サポートの自動化 |多言語メッセージング インテリジェンス
[切り捨てられた]

記事本文:
AI エージェントの新たなフロンティア: 実世界の経験から自己進化する | AI エージェントの新たなフロンティア: 実世界の経験から得た自己進化の洞察 | AI エージェントの新境地: 実世界の経験から自己進化する穏やかなコミュニケーション | AI エージェントの新たなフロンティア: 現実世界のエクスペリエンスから自己進化する AI 受信トレイの概要 | AI エージェントの新たなフロンティア: 現実世界のエクスペリエンスから自己進化する生産性ハンドブック | Alloomi コンテンツを読み込んでいます... ホーム 価格設定 ドキュメント ブログ AI エージェントの新たなフロンティア: 実世界の経験から自己進化する
AI エージェントは長い道のりを歩んできました。それらは通話ツールから、サブエージェントのチームを調整するものへと進化し、そして今ではますます大規模なコンテキストを処理するようになりました。しかし、ほとんどのエージェントにとって、100 回目でも何かを行うのは最初のときと何ら変わりません。人間とは異なり、経験によって能力が向上するわけではありません。今日、エージェントを実際の仕事から引き離しているのは、もはやモデル能力、ツール、メモリ、コンテキストではなく、現実世界での経験と、そこから学ぶ能力であることがますます増えています。
この問題を解決しようとするほとんどの試みは、核心的な問題には触れていません。
アプリケーション ラッパー/エージェント ハーネス。ツールを接続してワークフローを整理できますが、モデル自体は静的なままです。それらの能力の上限は、その下のモデルの上限のままです。
RAG / 外部ナレッジ ベース。文書、会話、データベースから事実を引き出すことができます。しかし、彼らが取り戻すことができないのは、専門家の考え方、つまり意思決定がどのように行われるのか、なぜ改訂が行われるのか、「良い」とは実際にはどのようなものなのか、ということです。それはモデルには決して入りません。そのため、AI がタスクを取得するたびに、原材料からやり直します。
微調整。定期的な再トレーニングは費用がかかり、時間がかかり、常にビジネスよりも一歩遅れます。
内省的な学習。経験値なし

アンカリングやそれ自体を評価するための信頼できる方法ではない場合、モデルはそれ自身のレベルでただ周回するだけです。さらに悪いことに、コースから外れてしまいます。
問題は、ツール、ナレッジベース、微調整が役に立たないということではありません。それは、知識と経験がモデルの外に留まるということです。これらのアプローチは AI により多くの情報を提供しますが、エージェント自体の能力が向上するわけではありません。モデル内に存在するものは、引き続きエージェントを複合させ、進化させることができます。
同時に、この種の経験データは実際の作業を通じてのみ生成されるため、非常に希少です。通常、それはプライベートで動的であり、プロセス中に表面化しますが、場合によっては 1 人の人の頭の中にだけ存在することもあります。
実際にエージェントを若手アシスタントから経験豊かなパートナーに成長させるのは、判断とフィードバックをもたらす職務経験です。そして、そのような経験はインターネット上に存在しません。これはまさに汎用 LLM が頭打ちになるところです。
私たちのアプローチ: フルスタック、モデル + アプリケーション
Alloomi では、アプリケーション層とモデル層を組み合わせたフルスタックのアプローチを採用しています。
エージェントは実際のワークフロー内で作業し、パブリック ドメインには存在しない専門的なデータを取得します。そして、その作業から学んだことは、トレーニング後のモデルに組み込まれます。実際には、これは 4 つの層にわたって発生します。
全体的なコンテキスト - モデルが全体像を把握できるようにする
人々、会話、文書、関係、タイムライン、意思決定、結果、フィードバックを統一されたコンテキストに取り込んで、作業の完全な軌跡を継続的に追跡します。
従来の RAG は、ファクトが静的ドキュメント内に存在することを前提としています。しかし、実際の仕事では、事実が修正され、覆され、顧客ごとに異なる意味を持つ可能性があります。
検索は「何が起こったのか」に答えます。総合的なコンテキストの答え「何が起こったのか」

帽子は。」
自己進化する記憶モデル - 働きながら学ぶ
実際の作業から得られたコンテキスト、専門家の判断、改訂履歴、納品結果、顧客からのフィードバックがフィルタリングされ、再生され、トレーニング後のトレーニングに使用されます。結果として得られるエクスペリエンスは、外部データベースに保存されるのではなく、モデル自体の重みに書き込まれます。
エキスパートのアンカーリング - 能力を向上させ続ける
学習中、モデルは専門家のデモンストレーションと最良の成果物の基準に固定されます。そうすれば、独自のレベルで循環したり、エラーが悪化したりすることはありません。
制御された進化 - 進化を安全にする
自己進化は安全な場合にのみ役に立ちます。高品質のゲート、継続的な監視、自動ロールバックにより、すべてのモデル変更が検証可能、監査可能、および元に戻せる状態が保たれます。ドリフトは問題になる前に検出され、元に戻されます。
これでフライホイールが完成します。完了したすべてのタスクはデータを獲得したものになります。あらゆる判断が次の結果に影響を及ぼします。仕事はエージェントをより良くし、より良いエージェントはより良い仕事を生み出します。
技術的基盤: 記憶、学習、配信に関する 9 つのベンチマーク
4 つのレイヤーは、Aloomi の仕組みを説明します。私たちは、Alloomi が重要な部分でテストします。全体像を理解しているか、学習を続けているか、プロフェッショナルな仕事を提供できるかなどです。 9 つのベンチマークは、長期記憶から現実世界の配信までのチェーン全体を追跡します。
全体的な理解 - それは記憶します
実際の仕事には、多くの会話、多くのタスク、そして長時間にわたる作業が伴います。生産的に参加するには、エージェントは何が重要かを思い出し、人、出来事、時間がどのように結びついているかを理解し、歴史が成長し続けるにつれて全体像を把握しておく必要があります。 3 つのベンチマークは、Aloomi のコンテキスト機能を示しています。
BEAM では、Alloomi は 128K で 72.8%、500K で 75.7%、1M で 76.5% のグローバル タスク精度を達成し、

10M で 67.0%。 1,000 万スケールでも 67.0% を維持し、Hindsight の 64.1% を上回っています。これは、非常に長い履歴にわたって比較的安定したグローバル タスクのパフォーマンスを維持できることを示しています。
LongMemEval-S では、Memo-V3 の 94.4% と比較して、Alloomi は 97.6% に達します。これにより、情報抽出、知識の更新、および複数セッションの推論にわたる長期記憶が検証されます。
LoCoMo-V2 では、Memo-V3 の 92.5% と比較して、Alloomi のスコアは 97.4% であり、孤立した事実を取得するだけでなく、セッション間の質問応答、一時的な関係、およびマルチホップ推論を処理する能力を示しています。
継続的な進化 - それは学びます
記憶は始まりにすぎません。さらに、長期的なエージェントは、複雑なコンテキストから新しいルールを学習すること、タスク間で経験を転送すること、環境の変化に応じて苦労して獲得した機能を維持すること、という 3 つの困難なテストに直面します。ここでは、進化に関連した 3 つのベンチマークをテストしました。
CL ベンチでは、Alloomi は GPT-5.6 Sol ベースラインの 21.5% から 47.6% に改善し、26.1 パーセント ポイント増加しました。これは、複雑なコンテキストから新しいルールを学習し、実行に適用する能力を示しています。
CL-Bench-Life では、パフォーマンスは GPT-5.5 (高) ベースラインの 22.2% から 32.1% に上昇し、9.9 パーセント ポイント改善しました。この結果は、モデルが長期的なタスクにわたって経験を蓄積し、それを後のステージに転送できることを示しています。
Con.L Bench では、パフォーマンスが Claude Sonnet 4.6 ベースラインの 22.3% から 32.6% に増加し、10.3 パーセント ポイント増加しました。これは、以前に獲得した機能の損失を軽減しながら、モデルがタスク間で学習を継続できるかどうかを評価します。
これらの結果を総合すると、Alloomi は一度だけ適応するのではなく、時間の経過とともに経験を複合させ、それをタスク全体に引き継いでいくことがわかります。
プロフェッショナルがお届けします

y - 発送します
記憶と学習は、最終的には実際の仕事を改善する場合にのみ重要です。このグループは、単独での記憶や学習を超えて、価値の高い専門的なタスク、複雑な現実世界のワークフロー、継続的なソフトウェア エンジニアリングにわたって Alloomi が結果を出せるかどうかをテストします。
GDPval-AA Normalized では、Alloomi は 44 の職業にわたる価値の高いタスク全体で 74.2% に達し、Claude Opus 5 の 67.9% を 6.3 ポイント改善しました。これは、長期的なコンテキストと継続的な学習が、複雑な専門的作業におけるより強力なパフォーマンスにつながる可能性があることを示しています。
JobBench では、Muse Spark 1.1 の 54.7% と比較して、Alloomi のスコアは 57.5% で、2.8 パーセント ポイント向上しました。これにより、実際の業務タスクや専門的な成果物でより良い結果を生み出す能力が検証されます。
SWE-Bench-CL では、Alloomi は 80.6% に達しましたが、OpenCode + Kim K3 + FAISS の場合は 73.3% で、 7.3 パーセントポイントの向上でした。この結果は、このシステムが実際の G​​itHub の問題を解決できるだけでなく、一連のソフトウェア エンジニアリング タスク全体にわたってその機能を保持し、構築できることを示しています。
これらのベンチマークを総合すると、記憶力と学習能力の向上が、最終的には現実世界の仕事の遂行において目に見える利益につながる可能性があることを示しています。
9 つの結果すべての詳細な実験設定、評価方法、報告規則は、2 つの公開技術レポートで入手できます。
読者が結果を検証できるようにこれらの詳細を公開しました。コミュニティからの議論や修正を歓迎します。
現在、私たちは 2 つのことを行っています。それは、プロフェッショナル サービス向けのデジタル従業員である Alloomi AI の構築、そしてそれを強化するコンテキスト レイヤーである OpenContext のオープンソース化です。
OpenContext は AI エージェント用のコンテキスト ランタイムであり、現在オープンソース化されています。それは実用的な実装です

総合的なコンテキスト層の。コンテキスト ハーネスとしてエージェント アプリケーションに組み込まれ、一時的なコンテキスト、メモリと取得、コンテキストの修正、マルチプラットフォーム接続、プロアクティブなスケジューリングを提供します。このプロジェクトはまだ若いです。試して、壊して、貢献し、コンテキスト ハーネスがどのようなものになり得るかをマッピングするのに協力してください。
Alloomi AI は、プロフェッショナル サービスの個人およびチーム向けの商用製品であり、自己進化するデジタル従業員の OKR 主導の成果重視のシステムです。長期的なビジネスコンテキストを理解し、独自に作業を分解して推進し、専門家の判断、実行履歴、実際の運用からの結果フィードバックを各顧客固有の機能に変換します。私たちは、法律、保険、財務顧問などの最も専門的な分野のいくつかで初期の共同設計を通じてこれをすでに検証しており、より専門的な領域に拡大しながら、中核となるニーズについて初期ユーザーと緊密に協力し続けています。
私たちは、AI エージェントの次の段階は、モデルがどれほどインテリジェントであるかによって決まるのではなく、モデルが現実世界で行うすべての実際の作業から何を学習するかによって決まると信じています。
すべてのフロンティア モデルは箱から出してすぐに素晴らしいものであり、それを使用するすべての人にとっても同様です。他の誰もコピーできないものは、エージェントがあなたのビジネス内に蓄積するものです。判決は、エージェントが吸収したもの、エージェントが学習した改訂、配信後の配信に組み込まれた標準と呼んでいます。
インテリジェンスはすべてのエージェントの出発点です。経験はそれらが分岐する場所です。それが私たちが築いている未来です。
© 2026 メランド研究所。無断転載を禁じます。

## Original Extract

From real work to real capability - how Alloomi turns professional experience into compounding model weights through a four-layer flywheel. | the new frontier of ai agents calm communication | AI inbox summary best practices | privacy-first customer support automation | multilingual messaging intell
[truncated]

The New Frontier of AI Agents: Self-Evolving from Real-World Experiences | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences insights | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences calm communication | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences AI inbox summaries | The New Frontier of AI Agents: Self-Evolving from Real-World Experiences productivity playbook | Alloomi Loading content... Home Pricing Docs Blogs The New Frontier of AI Agents: Self-Evolving from Real-World Experiences
AI agents have come a long way. They evolved from calling tools, to coordinating teams of sub-agents, and now to handling ever-larger contexts. But for most agents, doing something for the 100th time is no different from the first. Unlike humans, they don't grow more capable through experience. Today, what holds agents back in real work is no longer model capability, tools, memory, or context - increasingly, it's real-world experiences, and the ability to learn from them.
Most attempts to solve this problem don't touch the core issue:
Application wrappers / agent harnesses. They can connect tools and organize workflows, but the model itself stays static. Their capability ceiling remains that of the model underneath.
RAG / external knowledge bases. They can pull facts from documents, conversations, and databases. But what they can't retrieve is the expert's way of thinking - how decisions get made, why revisions happen, what "good" actually looks like. That never makes it into the model. So every time the AI picks up a task, it starts from the raw material all over again.
Fine-tuning. Periodic retraining is expensive, slow, and always a step behind the business.
Self-reflective learning. Without expert anchoring or a reliable way to evaluate itself, a model just circles at its own level - or worse, drifts off-course.
The problem isn't that tools, knowledge bases, or fine-tuning are not useful. It's that knowledge and experience stay outside the model. These approaches give AI more information, but they do not make the agent itself more capable. What lives inside the model can continue to compound and evolve the agent.
At the same time, this kind of experience data is exceptionally scarce, because it's produced only through real work. It's usually private, dynamic, and surfaces during the process - sometimes existing in nothing but one person's head.
What actually grows an agent from a junior assistant into a seasoned partner is work experience - the kind that carries judgment and feedback. And that kind of experience simply isn't on the internet. This is exactly where general-purpose LLMs hit their ceiling.
Our approach: full stack, model + application
At Alloomi, we take a full-stack approach, combining the application layer and the model layer.
Agents work inside real workflows, capturing professional data that doesn't exist in public domains; and what they learn from that work is built back into the model through post-training. In practice, this happens across four layers:
Holistic context - so the model can see the whole picture
Bring people, conversations, documents, relationships, timelines, decisions, outcomes, and feedback into a unified context, then continuously track the complete trajectory of the work.
Traditional RAG assumes that facts live in static documents. But in real work, facts are revised, overturned, and may carry different meanings for different customers.
Retrieval answers "what happened." Holistic context answers "how what happened became what is."
A self-evolving memory model - learning while working
Context, expert judgment, revision histories, delivery outcomes, and customer feedback from real work are filtered, replayed, and used for post-training. The resulting experience is written into the model's own weights, rather than living in an external database.
Expert anchoring - keeping capability moving up
During learning, the model is anchored to expert demonstrations and the standards of the best deliverables. That way, it does not circle at its own level, nor let errors compound.
Controlled evolution - making evolution safe
Self-evolution is only useful if it's safe. Quality gates, continuous monitoring, and automatic rollback keep every model change verifiable, auditable, and reversible - drift is caught and undone before it matters.
And that completes the flywheel: every task finished is data earned; every judgment captured compounds into the next delivery. The work makes the agent better, and a better agent makes better work.
The technical foundation: nine benchmarks for memory, learning, and delivery
The four layers describe how Alloomi works. We test Alloomi where it counts: does it understand the whole picture, does it keep learning, and does it deliver professional work. Nine benchmarks trace the entire chain from long-term memory to real-world delivery.
Holistic understanding - it remembers
Real work spans many conversations, many tasks, and long stretches of time. To participate productively, an agent has to remember what matters, understand how people, events, and time connect, and hold onto the whole picture as its history keeps growing. Three benchmarks show Alloomi's context capabilities:
On BEAM , Alloomi achieves global task accuracy of 72.8% at 128K , 75.7% at 500K , 76.5% at 1M , and 67.0% at 10M . Even at the 10M scale, it maintains 67.0%, above Hindsight's 64.1%, demonstrating that it can preserve relatively stable global task performance across extremely long histories.
On LongMemEval-S , Alloomi reaches 97.6% , compared with 94.4% for Memo-V3. This validates long-term memory across information extraction, knowledge updates, and multi-session reasoning.
On LoCoMo-V2 , Alloomi scores 97.4% , compared with 92.5% for Memo-V3, demonstrating its ability to handle cross-session question answering, temporal relationships, and multi-hop reasoning - not just retrieving isolated facts.
Continual evolution - it learns
Memory is only the beginning. Beyond that, a long-horizon agent faces three harder tests: learning new rules from complex context, transferring experience across tasks, and keeping hard-won capabilities intact as the environment changes. We tested three evolution-related benchmarks here:
On CL-Bench , Alloomi improves from the GPT-5.6 Sol baseline of 21.5% to 47.6% , a gain of 26.1 percentage points . This demonstrates its ability to learn new rules from complex context and apply them in execution.
On CL-Bench-Life , performance rises from the GPT-5.5 (high) baseline of 22.2% to 32.1% , an improvement of 9.9 percentage points . The result shows that the model can accumulate experience over long-horizon tasks and transfer it to later stages.
On Con.L Bench , performance increases from the Claude Sonnet 4.6 baseline of 22.3% to 32.6% , a gain of 10.3 percentage points . This evaluates whether the model can continue learning across tasks while reducing the loss of previously acquired capabilities.
Together, these results show that Alloomi doesn't just adapt once - it compounds experience over time and carries it across tasks.
Professional delivery - it ships
Remembering and learning ultimately matter only if they improve real work. This group moves beyond memory or learning in isolation and tests whether Alloomi can deliver results across high-value professional tasks, complex real-world workflows, and continual software engineering.
On GDPval-AA Normalized , Alloomi reaches 74.2% across high-value tasks spanning 44 occupations, improving on Claude Opus 5's 67.9% by 6.3 percentage points . This indicates that long-term context and continual learning can translate into stronger performance on complex professional work.
On JobBench , Alloomi scores 57.5% , compared with 54.7% for Muse Spark 1.1 - an improvement of 2.8 percentage points . This validates its ability to produce better results on real occupational tasks and professional deliverables.
On SWE-Bench-CL , Alloomi reaches 80.6% , compared with 73.3% for OpenCode + Kimi K3 + FAISS, a gain of 7.3 percentage points . The result shows that the system can not only resolve real GitHub issues, but also retain and build on its capabilities across a sequence of software-engineering tasks.
Together, these benchmarks show that improvements in memory and learning can ultimately translate into measurable gains in real-world work delivery.
Detailed experimental settings, evaluation methods, and reporting conventions for all nine results are available in two public technical reports:
We have published these details so readers can verify the results, and we welcome discussion and corrections from the community.
Today we're doing two things: building Alloomi AI , digital employees for professional services, and open-sourcing OpenContext , the context layer that powers it.
OpenContext is our context runtime for AI agents, now open-sourced. It is the working implementation of the holistic context layer. It embeds into agentic applications as a context harness, providing temporal context, memory and retrieval, context correction, multi-platform connectivity, and proactive scheduling. The project is young - try it, break it, contribute, and help us map what context harnesses can become.
Alloomi AI is our commercial product for individuals and teams in professional services: an OKR-driven, outcome-focused system of self-evolving digital employees. It understands long-term business context, breaks down and advances work on its own, and converts expert judgment, execution history, and outcome feedback from real operations into capabilities unique to each customer. We've already validated this through early co-design in some of the most specialized fields - legal, insurance, and financial advisory - and we're continuing to work closely with early users on their core needs while expanding into more specialized domains.
We believe the next phase of AI agents won't be decided by how intelligent a model is, but by what it learns from every piece of real work it does in the real world.
Every frontier model is brilliant out of the box, and identically so for everyone who uses it. What no one else can copy is what an agent accumulates inside your business: the judgment calls it absorbed, the revisions it learned from, the standards it internalized delivery after delivery.
Intelligence is where every agent begins. Experience is where they diverge. That's the future we are building for.
© 2026 Meland Labs. All rights reserved.

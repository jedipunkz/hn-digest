---
source: "https://www.notion.com/blog/building-shared-memory-for-ai-agents-in-notion"
hn_url: "https://news.ycombinator.com/item?id=49357879"
title: "Building Shared Memory for AI Agents in Notion"
article_title: "Building Shared Memory for AI Agents in Notion"
image: "https://images.ctfassets.net/spoqsaf9291f/3KcypMugRMDnYguXzeAcKm/da0e5d9d74872fbf0769a50432fe9e8e/Lore.jpg"
author: "asherwood"
captured_at: "2026-08-19T07:30:47Z"
capture_tool: "hn-digest"
hn_id: 49357879
score: 2
comments: 0
posted_at: "2026-08-19T06:55:20Z"
tags:
  - hacker-news
  - translated
---

# Building Shared Memory for AI Agents in Notion

- HN: [49357879](https://news.ycombinator.com/item?id=49357879)
- Source: [www.notion.com](https://www.notion.com/blog/building-shared-memory-for-ai-agents-in-notion)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T06:55:20Z

## Translation

タイトル: Notion で AI エージェント用の共有メモリを構築する
説明: Lore は、会話、意思決定、タスク、永続的な事実を、将来のエージェント (および一緒に働く人々) が簡単に使用できる Notion ページに変換します。

記事本文:
Product Concept AI 業務用AIツール
AI 会議メモ AI が完璧に作成
エンタープライズ検索 答えをすぐに見つける
ナレッジベース 知識を一元化
。新機能を見る → Notion アプリをダウンロード → AI AI 機能
AI 会議メモ AI が完璧に作成
エンタープライズ検索 答えをすぐに見つける
すべての投稿 ← ← すべての投稿 2026 年 8 月 18 日に公開 Notion で AI エージェントの共有メモリを構築する技術
ソフトウェア エンジニアリングが変化するにつれて、部族の知識は 1 人の人の頭の中だけでなく、1 人のエージェント セッションにもますます集中しています。ここ Notion では、経験から得た学習をスキルまたはドキュメントとして手動で抽出し、リポジトリ内でチームと共有する必要があり続けました。しかし、それはスケールが合いません。エージェントがソース内の情報を推測したり発見したりするのが非常に上手になったとしても、ユーザーの好みとは別のこの経験的な知識には明らかなギャップが残っていました。逆もまた真です。エージェントは、現在のワークストリームとは関係ないものの役立つ情報を発見することがよくあり、それらのタスク、リファクタリング、ToDo は後で失われます (特に無人セッションの場合)。
私たちは、どの組織にも独自の伝承があり、同様の問題に直面しているに違いないことに気づきました。私たちはそれを Notion で解決する時期が来たと判断しました。
Lore は、そのギャップのために私たちが構築したツールです。Notion によってサポートされる、エージェント用の共有永続メモリです。これは、github.com/makenotion/lore で MIT ライセンスに基づくオープンソースです。
私たちは、エージェント エクスペリエンスを永続的なコンセプトにして、Notion に保存し、人間とエージェントの両方がアクセスできるようにしたいと考えていました。紛失しないようにしたいアイテムが複数ありました。
経験的な知識。これは、将来のエージェントが推測することができない種類の情報であり、再体験する必要があります。長い再発見フェーズは短くなる可能性があります

早期にT回路を形成するか、完全に阻止します。
フォローアップタスク。多くの場合、エージェントが特定のタスクに取り組んでいるとき、エージェントは特定の関連性のない、関係のない情報をメモします。エージェント セッションの性質上、この情報は本質的に一時的なものです。人間のオペレーターが積極的にトランスクリプトを読んでいなかった場合、これらの未提出のフォローアップ アイテムは永久に失われます。
決定記録。一部のチームは ADR を維持する健全な習慣を持っています。ほとんどのチームはこれを目指していますが、意思決定が進むにつれてリアルタイムでそれを行うことはできません。
手順。これらは具体的な回避策またはレシピです。スキルにエンコードできるものの、使用頻度が低すぎて保証できないもの。
これらのアイテムは、後のセッションで読み返したり、人間が必要に応じて参照したりできる Notion ページになります。 MCP 互換のエージェント ハーネスはボールトを参照でき、その結果、エージェント セッションは過去のエージェントによって強化されます。
ネイティブ メモリ システムは超局所化されています。ハーネス内システムは、多くの場合、特定のハーネスまたは特定のモデル プロバイダーに関連付けられています。ファイルベースのシステムは、多くの場合、単一のマシンにローカライズされます。
私たちは、ある人のデスクにいるエージェントが、別の人のエージェントの経験から、おそらく組織の別の部門の経験からも学ぶことができるというコンセプトを実験したかったのです。私たちは、推進すべきもののための共有プラットフォームを望んでいました。
Lore Vault は、5 つのデータベースを含む Notion ページです。
プロジェクト、個人、チーム、またはエージェントの範囲
その範囲内のサブジェクト領域
物語の文脈、メモ、手順、タスク
メモリグラフが参照できる名前付きのもの
構造化された主語-述語-目的語アサーション
これらのテーブルには相互に関係があり、情報の表面化やあいまい検索に役立ちます。たとえば、ファクトはトピックと関係を持つことができます。

または思い出。これはデータの独自の構造化ですが、構造化データにより、エージェントは予測可能な検索構造を持つことができます。
エージェントは、さまざまな lore-* MCP ツールを使用して、MCP を通じて Lore にアクセスします。ただし、人間は、Notion を介して直接 (結局のところ、これらはすべて Notion ページです!)、または Lore CLI を介して、同じ情報にアクセスできます。生来の設計は主にフックによって駆動され、標準パスを完全に手動で行うことを目的としています。エージェント セッションは関連するコンテキストで開始され、必要に応じて情報を保存または取得するよう求められ、作業を実行し、新しい情報をバックグラウンドで静かに保存します。
私たちはシングルプレイヤーのソリューションには興味がありませんでした。 Notion はチーム中心のマルチプレイヤー サービスにとって自然なホームであると感じました。チームはすでに運用情報を Notion に保存し、一緒に計画し、一緒に構築するなどしています。 Notion に組み込まれたバージョン管理およびバージョン履歴メカニズムは、使いやすさと組み合わされてさらに優れたものになります。私たちは、情報保管庫が人間中心でありながら、エージェントの力を活用できるようにしたいと考えていました。
メモリ システムの最初の不良バージョンは、API を備えたジャンク ドロワーです。次の場合、単一フォルダー内のすべてのトランスクリプト、すべてのテキスト、およびすべてのセッションはあまり役に立ちません。
情報が古くなると、本来は成功していたはずのエージェントに損害を与える可能性があります。
無関係な情報がそれらのエージェントに表面化します。
正しい情報を見つけることは不可能です。
デフォルトでは、ファクトは期限切れになります。情報は継続的に強化する必要があり、そうでないと無価値なものになってしまいます。情報への追加は、前のエントリとの関係を明確にする既知の関係値 ( supersedes 、scoped 、conflicts_with など) で処理されます。 CLI は、孤立したファクト、重複した情報、その他のメモリ負債を検索するためのツールを提供します。
記憶とは何か

手動で、または専門のエージェントを使用して、定期的に保守する必要があります。ベンチマークの結果から、メモリは主に、利用可能で具体的で、適切なタイミングで取得できる場合にのみ役に立ちますが、曖昧であったり、冗長であったり、古くなったり、あるいは当面のタスクに関係がない場合には悪影響を与える可能性があることが明らかになったようです。これは小型モデルではさらに顕著です。
プロセスをガイドするために、定量化可能な数値が必要でした。静かに害を及ぼしながらも効果があるように感じるものは珍しいことではありません。私たちは、次の 2 つの関連ベンチマークで答えを求めました。
1. 検索 。私たちは正しい情報を引き出しているでしょうか?
2. モデルハード評価。これらはモデルの能力の範囲外であり、モデルが確実に成功することができないものです。これは、「メモリが利用可能な場合、それは役に立ちますか?」という質問に答えるためのものです。
私たちはオープンな SkillRet データセットを使用しました。このデータセットには、4,997 のクエリ、6,660 のスキル、および 8,347 の関連性判断が含まれています。スキルを手順メモリとして Lore 評価ボールトにインポートし、Codex エージェントに読み取り専用の Lore 命令と検索および展開ツールを提供します。合格するには、Lore を使用し、期待されるメモリを表面化して拡張し、適切な SkillRet ターゲットを選択し、そのターゲットを回答に適用する必要があります。 500 個のクエリをランダムに選択したところ、次のことがわかりました。
95% 成功のためのウィルソン間隔
ツールの結果にターゲットが表示されました
ターゲットの選択/回答の適用
これにより、そもそも正しいものを特定できるかどうかを検討することができました。これにより、Lore が支援できる状況の最大数に上限が設けられました。適切なデータを明らかにできない場合、記憶は役に立ちません。悪いデータを表面化すると、ツールが損害を与える可能性があります。
2 番目の、答えるのが難しい質問です。適切なメモリが利用可能であれば、メモリは使用できますか?

実際のエージェントの作業を改善できないですか?
このプロセスの設計は、特定の SHA に固定されたいくつかの大規模な OSS リポジトリをプルすることでした。それらの履歴は、単一のコミットに凝縮することで消去されます。私たちは、Lore なしではモデルが確実に提供できなかった機能リクエスト、バグ修正、パフォーマンスの改善を多数考案しました。次に、Lore がモデルのパフォーマンスを向上させる頻度と、モデルのパフォーマンスに悪影響を与える頻度を測定しました。
このことからわかることは、ボールトに関連するメモリがある場合、ネット上のモデルのパフォーマンスを向上させることができるということです。データ セットを最も困難なタスク (メモリのないエージェントでは決して成功しなかったタスク) のみにフィルタリングすると、シードされたメモリ エージェントは失敗の約 46% を回復しました。
重要な点は、記憶が魔法のようにエージェントを修復するわけではないということです。これは、利用可能な関連性のある記憶が結果を変えることを意味します。そして、それは最終的には、適切に保管および管理された情報の保管庫から生じます。
私たちの社内チームの 1 つが Lore をしばらく試験運用し、大きな保管庫を構築しました。最小限のメンテナンスで、次のような形状になりました。
会話やメモの記憶の 55% ～ 60% は経験的なものであり、保存する価値がありました。
15% ～ 20% は重複または重複に近いものであった可能性があります。
8 ～ 12% は主にエージェントが独自に推測できる情報でした。
そして、約 13% には貴重な情報がまったく含まれていませんでした。
このボールトに対してベンチマークを実行しませんでした。より乱雑な Vault ではパフォーマンスが低下するのは当然ですが、結論は変わりません。
適切なメモリが利用可能であれば、エージェントはそれを使用できるようになり、実際の向上が見られます。
エージェントのパフォーマンスを向上させる可能性のあるシステムのメモリ品質を最適化することは、最初の問題よりもはるかに興味深い問題です。
Lore は、GitHub: github.com/makenotion/lore で MIT ライセンスを取得しています。それは 20% のプロジェクトとして、主に尖ったものとして建設されました。

これらのシステムを Notion の適応外使用に向けて探索します。
これを開発依存関係として追加し、 Notion Personal Access Token を作成し、.lore.yaml をボールト ページに指定して、アシスタントに接続します。
あなたが Lore を基にして構築したり、フォークしたり、あるいは私たちが予期していなかった場所に持ち込んだりするのであれば、私たちはそれを見てみたいと思っています。 Issue とプルリクエストは github.com/makenotion/lore/issues で公開されています。
価格に関するヘルプ、デモ、使用例などを入手できます。

## Original Extract

Lore turns conversations, decisions, tasks, and durable facts into Notion pages that future agents (and the people they work with) can easily use

Product Notion AI AI tools for work
AI Meeting Notes Perfectly written by AI
Enterprise Search Find answers instantly
Knowledge Base Centralize your knowledge
. See what’s new → Download the Notion App → AI AI features
AI Meeting Notes Perfectly written by AI
Enterprise Search Find answers instantly
All posts ← ← All posts Published August 18, 2026 in Tech Building Shared Memory for AI Agents in Notion
As software engineering changes, tribal knowledge is increasingly concentrated not only in a single person's mind, but in a single agent session. Here at Notion, we kept having to manually extract experiential learnings as skills or documentation and share them with the team in-repo. But that doesn’t scale. Even as agents got really good at inferring or discovering information in source, this experiential knowledge—which is separate from user preferences—remained a tangible gap. The inverse is also true: agents would often discover information that is tangential to their current workstream but helpful, and those tasks, refactors, and todos would later get lost (especially if it’s an unattended session!).
We came to realize that every org has its own lore and must be facing similar problems. We decided it was time to fix that with Notion.
Lore is the tool we built for that gap: shared, persistent memory for agents, backed by Notion. It is open source under the MIT license at github.com/makenotion/lore .
We wanted to make agent experiences durable concepts, stored in Notion, and accessible by both humans and agents. There were multiple items we wanted to ensure were not lost:
Experiential knowledge . This is the kind of information that cannot be inferred by a future agent and must be re-experienced. The lengthy rediscovery phases can be short-circuited early or prevented entirely.
Follow-up tasks . Often, when an agent is working on a specific task, the agent makes note of specific, unrelated, tangential information. Due to the nature of an agent session, this information is ephemeral by nature: if the human operator was not actively reading the transcript, these un-filed followup items are gone forever.
Decisions records . Some teams have a healthy practice of keeping ADRs. Most teams strive for this, but aren’t able to do so in real-time as decisions evolve.
Procedures . These are specific workarounds or recipes. Things that could be encoded into a skill but may have usage too infrequent to warrant it.
These items become Notion pages that a later session can read back, or that a human can reference as needed. Any MCP-compatible agent harness can reference the vault, and as a result agent sessions are empowered by agents past.
Native memory systems are hyper-localized. In-harness systems are often tied to a given harness or given model provider. File-backed systems are often localized to a single machine.
We really wanted to experiment with the concept that the agents on one person’s desk could learn from the experiences of another person’s agents—perhaps even from another part of the organization. We wanted a shared platform for the things that should be carried forward.
A Lore vault is a Notion page with five databases:
The project, person, team, or agent scope
The subject areas inside that scope
Narrative context, notes, procedures, tasks
Named things the memory graph can refer to
Structured subject-predicate-object assertions
They have relations to one another, across these tables, to help in surfacing information and to help in fuzzy searching: for example, facts can have relations to topics or memories. This is an opinionated structuring of data, but structured data allowed for agents to have a predictable search structure.
Agents access Lore through the MCP, with various lore-* MCP tools. Humans, however, get access to that same information either through Notion directly (after all, these are all Notion pages!), or via a Lore CLI. The innate design is driven largely through hooks, which are intended to make the standard path entirely hands-off: agent sessions start with relevant context, are prompted to store or fetch information as needed, do their work, and quietly save new information in the background.
We weren’t interested in single-player solutions. Notion felt like the natural home for a team-focused, multiplayer service. The team already stores operational information in Notion, plans together, builds together, and more. Notion’s built-in version control and version history mechanisms combined with its familiarity are added bonuses. We wanted the information vault to be human-centric, but agent powered.
The first bad version of a memory system is a junk drawer with an API. All of the transcripts, all of the text, and all of the sessions in a single folder don’t do much if:
Stale information can otherwise harm an agent that would have otherwise been successful.
Irrelevant information is surfaced to those agents.
The right information is impossible to find.
By default, facts expire. Information must be continually reinforced or it will degrade into irrelevance. Addendums to information are handled with known relational values ( supersedes , scoped , conflicts_with , etc) that make the relationship to the previous entry clear. The CLI provides tooling for finding orphaned facts, duplicated information, and other memory debt.
Memory is something that must be periodically maintained, either manually or through use of a specialized agent. Benchmark results seemed to make it clear that memory largely helps only when it’s available, specific, and retrieved at the right time – but can also harm if it’s vague, redundant, stale, or has no bearing to the task at hand. This is even more pronounced on smaller models.
We wanted some quantifiable numeric values to help guide us through the process. Something feeling like it works while it silently harms is not uncommon. We sought to answer with two relevant benchmarks:
1. Retrieval . Are we pulling the right information?
2. Model-hard eval . These are things that were just outside of model capability, that models could not reliably succeed in doing. This was to answer the question: When the memory is available, does it help?
We used the open SkillRet data-set, which has 4,997 queries, 6,660 skills, and 8,347 relevance judgments. We import the skills into a Lore eval vault as procedure memories, then give a Codex agent read-only Lore instructions plus search and expand tools. To pass, it has to use Lore, surface the expected memory, expand it, select the right SkillRet target, and apply that target in the answer. We selected 500 queries at random, and found the following:
95% Wilson interval for success
Target surfaced in tool results
Target selected / answer applied
This allowed us to scope whether or not we’d be able to identify the right thing to begin with – it provides a ceiling for the maximum number of situations that Lore would be able to assist in. If we are unable to surface the right data, then the memory can’t help. If we surface bad data, then the tool could hurt.
The second, and harder question to answer: When the right memory is available, does it improve actual agent work?
The design of this process was to pull several large OSS repos pinned at specific SHAs. Their histories erased by condensing them into a single commit. We devised a number of feature requests, bug fixes, or performance improvements that the model could not reliably deliver without Lore. Then, we measured how often Lore lifted a model performance, as well as how often it harmed model performance.
What this tells us is: when we have relevant memories in the vault, we are able to improve model performance on net. When we filter the data set only to the hardest tasks (those that the no-memory agent never succeeded in doing), the seeded memory agent recovered about 46% of failures!
A key takeaway is not that memory magically fixes agents. It means that available, relevant memory changes outcomes—and that ultimately stems from a vault of information that is well kept and managed.
One of our internal teams piloted Lore for some time, and built up a large vault. With minimal upkeep, it was shaped such that:
55%-60% of conversations / note memories were experiential and worth keeping.
15%-20% were likely duplicates or near-duplicates.
8-12% were largely information that the agent can infer on its own.
and ~13% had no valuable information in them whatsoever.
We did not benchmark against this vault. It stands to reason that we would see degraded performance in messier vaults, but the conclusions remain the same:
When the right memory is available, agents can use it, and we see real lift.
Optimizing for memory quality on a system that has potential to improve agent performance is a much more interesting problem than the one we started with.
Lore is MIT-licensed on GitHub: github.com/makenotion/lore . It was built as a 20% project, largely as a pointed exploration of these systems into an off-label use of Notion.
Add it as a dev dependency, create a Notion Personal Access Token , point a .lore.yaml at a vault page, and wire it into your assistant.
If you build on Lore, fork it, or take it somewhere we did not expect, we want to see it. Issues and pull requests are open at github.com/makenotion/lore/issues .
Get pricing help, demos, use-cases, and more.

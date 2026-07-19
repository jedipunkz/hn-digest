---
source: "https://vivekhaldar.com/articles/compiling-an-ai-agent-skill/"
hn_url: "https://news.ycombinator.com/item?id=48969971"
title: "I Cut an AI Agent's Token Use by 94%"
article_title: "How I Cut an AI Agent's Token Use by 94%"
author: "gmays"
captured_at: "2026-07-19T17:49:53Z"
capture_tool: "hn-digest"
hn_id: 48969971
score: 1
comments: 0
posted_at: "2026-07-19T17:23:49Z"
tags:
  - hacker-news
  - translated
---

# I Cut an AI Agent's Token Use by 94%

- HN: [48969971](https://news.ycombinator.com/item?id=48969971)
- Source: [vivekhaldar.com](https://vivekhaldar.com/articles/compiling-an-ai-agent-skill/)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T17:23:49Z

## Translation

タイトル: AI エージェントのトークン使用量を 94% 削減しました
記事のタイトル: AI エージェントのトークン使用を 94% 削減する方法
説明: (このブログ投稿は、私の YouTube チャンネルのこのビデオと、関連するインセンティブに関する短いフォローアップから派生したものです。)
私はしばらくの間、特殊なエージェントのハーネスについて書き、タスク仕様からそれらを編集してきましたが、そのアイデアが抽象的に聞こえるかもしれないことに気づきました。それで、ここに具体的なものがあります
[切り捨てられた]

記事本文:
ヴィヴェク・ハルダル
アーカイブ
アイデアバース
YouTube
エンチリディオン研究所
RSS
AI と LLM
6 分で読めます
AI エージェントのトークン使用量を 94% 削減した方法
ヴィヴェク・ハルダー著 · 2026年7月13日
(このブログ投稿は、私の YouTube チャンネルのこのビデオと、関連するインセンティブに関する短いフォローアップから派生したものです。)
私はしばらくの間、特殊なエージェントのハーネスについて書き、タスク仕様からそれらを編集してきましたが、そのアイデアが抽象的に聞こえるかもしれないことに気づきました。ここでは、私が毎日実行しているワークフローの具体的な例を示します。
私には、ブログ投稿のバックカタログを調べて、再表示する価値のあるものを見つけ、それについて最近言及したかどうかを確認し、それにリンクする短い LinkedIn 投稿の下書きを作成するスキルがあります。下書きが自動的に投稿されることはありません。全く投稿しないことも多いです。時々、それをナッジとして使用して、すべてを自分の声で書き直します。しかし、このワークフローは便利です。既存の文章を再投稿するよう促してくれます。
オリジナルのバージョンは、完全にエージェント スキルの自然言語命令として書かれていました。そこには、検索する情報源、最近の履歴のチェック、選択基準、および最終草案の形状が記載されていました。実行のたびに、私のエージェント (現在 Codex) はこれらの指示を解釈し、計画を立て、ツールを呼び出し、ワークフローの状態を追跡する必要がありました。
これは最初のバージョンを構築するための素晴らしい方法でした。自然言語により、ワークフローの明確化と変更が容易になりました。私は以前、これらのファイルについて natlang コードとして書きました。これは、エージェントが最初にすべての判断を従来のプログラムに変換することなく作業を自動化できるようにする実行可能 SOP です。
しかし、スキルが何度も実行され、同じ地面を繰り返し踏むと、多くの場合、その動作のほとんどは探索的ではなくなります。
このスキルは常に同じ場所を探します。同じコンテンツ インベントリを構築します。

同じ最近の投稿フィルターが適用されます。同じ中間状態を保存します。毎朝、ゼロから推論する必要はありません。多くの場合、LLM さえ必要ありません。
LLM が必要となる手順は実際には 2 つだけです。
フィルタリングされた在庫から適切な候補者を選択します。
それ以外はすべて、通常の決定論的なコードで構いません。
そこで私はそのスキルを特殊なハーネスに「コンパイル」しました。新しいスキルはほとんどスキルではなく、Python プログラムを呼び出すシン ブートローダーになっています。このプログラムは、既知のソースを取得し、インベントリを構築し、最近の投稿を確認し、フィルターを適用し、ワークフローを管理します。 LLM は、選択と生成のためにのみ呼び出されます。
驚くことではないが、これにより、トークンの使用量が大幅に増加し、レイテンシが改善されました。
私の実行では基本的に同じ出力品質
小型または安価なモデルに交換しても、これらのメリットは得られませんでした。選択と生成のステップでは引き続き同じモデルが使用されます。この節約は、通常のコードで直接実行できる作業を行っていたすべてのモデル呼び出しを削除することで得られます。
汎用コーディング エージェントは、非常に有能な推論およびワークフロー エンジンです。また、これは、過去のいくつかの呼び出しを内省した後にその形状が判明するプロシージャを実行するのに高価な方法でもあります。
以前にスキルを実行したときの履歴トレースがありました (Codex で自動化として毎日実行されていました)。これらのトレースは、計画、ツールの呼び出し、分岐、途中で必要な状態など、エージェントが実際に行ったことを示していました。私はそれらの痕跡、オリジナルのスキル、そして特殊なハーネスについての私の文章を強力なモデルに与えました。私は、LLM が本当に必要なステップと、コードとして表現できるほど安定したステップを特定し、専用のハーネスを構築するように依頼しました。
言い換えれば、

自然言語スキルは高レベルの仕様として機能し、トレースはモデルによる推論の結果である操作の詳細を提供しました。
これが、コンパイラの例えが役立つと私が考える理由です。私たちは、意図を柔軟かつ高レベルで表現することから始めます。ワークフローを十分に実行して実際の形状を明らかにした後、安定した部分をより効率的な表現に落とし込みます。
このモデルは必要な場合に引き続き使用されます。この演習全体の目標は、決定論的なコードである可能性のあるワークフロー部分と、言語の理解、生成、推論が必要なワークフロー部分を区別することです。私たちは言語理解を脆弱なルールの山にしようとしたわけではありません。候補の選択は、情報源の内容と、それが興味深い投稿になるかどうかによって決まります。製図には言語生成の利点があることは明らかです。これらはモデル形式の問題であるため、モデル呼び出しのままです。
初日から特殊なハーネスを作成するのは時期尚早でした。正確なワークフロー、どのルールが重要になるのか、どこで判断が必要になるのか、私はまだ知りませんでした。自然言語スキル版では、プロセスを繰り返し実行することでそれらのことを発見できました。
しかし、ワークフロー全体を自然言語で永遠に維持するとは、実行のたびに同じ計画を再発見するためにモデルに料金を支払うことを意味します。
ワークフローを自然言語スキルとして表現します。
十分な回数実行してトレースを収集し、動作を調整します。
安定して決定的になった部分を見つけます。
それらの部分をコードにコンパイルします。
セマンティックな判断が重要となるいくつかのポイントで LLM 呼び出しを維持します。
このコンピレーション パスには 1 回限りの費用がかかります。強力なモデルとかなりの量のコンテキストを使用して、トレースを検査し、新しいハーネスを作成しました。しかし、そのコストは一度支払われます。コンパイルされたワークフロー

その後、何百回または何千回も実行できるため、実行のたびにトークンと時間を節約できます。これは標準的な最適化の経済学です。つまり、一度使用すれば、繰り返し使用することで償却されます。
私がローカル エージェントを好むようになった理由の 1 つは、ローカル エージェントにはまさにこの種の履歴 (スキル、トレース、状態、以前の実行のアーティファクト) が蓄積されているためです。その履歴は、エージェントが思い出すのに役立つだけではありません。それはハーネス自体をより良くするための原材料を与えてくれます。
エージェントの長期的な経済性はこれに依存します。すべての繰り返しワークフローがプランナー、ステート マシン、グルー コードとしてフロンティア モデルを永久に使用し続けると、コストとレイテンシが実際に安定することはありません。特殊なハーネスにより、モデルはインテリジェンスの恩恵を受けるワークフローのごく一部に集中できます。
私がこれらのテクニックが主に大手モデルベンダーから提供されると思わないもう 1 つの理由があります。彼らのビジネスはトークンの販売です。できれば、最も強力で高価なモデルからの大量のトークンが望ましいです。
古い決まり文句があります。「理解していないことが給料を左右する場合、誰かを説得するのは非常に難しい」です。インセンティブは行動と結果を説明します。トークンの消費を 94% 削減しながら出力の品質を維持する手法は、トークンの使用によって収益が増加する企業の現在の経済学に真っ向から反するものです。
私は、これらの企業内で効率性を気にしている人が誰もいないとは言いません。明らかにそうです。より優れた推論、キャッシュ、より安価なモデルにより、製品はより便利になります。ただし、各トークンを安くすることと、ほとんどのワークフローにはトークンがまったく必要ないことを顧客が発見できるようにすることには違いがあります。
少なくとも現時点では、大手ベンダーがこの 2 番目のアイデアを積極的に推進する動機はあまりないと思います。したがって、それはユーザーと独立した構築次第です

これらの最適化の機会を自分たちで見つけてください。
そしてそれはチャンスでもあります。繰り返し発生するエージェントのワークフローを検査し、決定的な部分をコードに移すための専用のハーネス、コンパイラー、ツールを構築する創設者にとっては、大きな余地があります。価値提案は異常に具体的です。つまり、出力の品質を維持しながら、出力のコストと待ち時間を削減します。
モデル ベンダーは、より強力なエンジンを開発し続けるでしょう。単純で決定的な作業を行うためにこれらのエンジンの回転数を上げていないことを誰かが確認する必要があります。
自然言語を使用してワークフローを発見します。トレースを使用してそれを理解してください。次に、結晶化したものをコンパイルします。
自分でも試してみませんか?完全な Token Shrinker プロンプトを取得し、エージェントにドロップします。

## Original Extract

(This blog post is derived from this video on my YouTube channel, along with a short follow-up about the incentives involved.)
I’ve been writing about specialized agent harnesses and compiling them from task specifications for a while, and I realize the idea can sound abstract. So here’s a concrete
[truncated]

Vivek Haldar
Archive
Ideaverse
YouTube
Enchiridion Labs
RSS
AI & LLMs
6 minute read
How I Cut an AI Agent's Token Use by 94%
By Vivek Haldar · July 13, 2026
(This blog post is derived from this video on my YouTube channel , along with a short follow-up about the incentives involved .)
I’ve been writing about specialized agent harnesses and compiling them from task specifications for a while, and I realize the idea can sound abstract. So here’s a concrete example from a workflow I run every day.
I have a skill that looks through my back catalog of blog posts, finds something worth resurfacing, checks whether I’ve mentioned it recently, and drafts a short LinkedIn post linking back to it. The draft never gets posted automatically. Often I don’t post it at all. Sometimes I use it as a nudge and rewrite the whole thing in my own voice. But the workflow is useful: it nudges me to repost existing writing.
The original version was written entirely as natural-language instructions in an Agent Skill . It described the sources to search, the recent-history checks, the selection criteria, and the shape of the final draft. On every run, my agent (currently Codex) had to interpret those instructions, make a plan, call tools, and keep track of the state of the workflow.
That was a great way to build the first version. Natural language made the workflow easy to articulate and change. I’ve written before about these files as natlang code : executable SOPs that let an agent automate work without first turning every judgment into a conventional program.
But after a skill has run many times, and trodden the same ground repeatedly, often most of its behavior is no longer exploratory.
This skill always looks in the same places. It builds the same content inventory. It applies the same recent-post filters. It saves the same intermediate state. None of that needs to be reasoned through from scratch each morning. A lot of it doesn’t even need an LLM!
There are really only two steps where LLMs are needed:
Choosing a good candidate from the filtered inventory.
Everything else can be ordinary deterministic code.
So I “compiled” the skill into a specialized harness. The new skill is barely a skill at all—it has become a thin bootloader that invokes a Python program. That program fetches the known sources, constructs the inventory, checks recent posts, applies filters, and manages the workflow. It calls an LLM only for selection and generation.
Not surprisingly, this resulted in massive token usage and latency improvements:
essentially the same output quality in my runs
I didn’t get those gains by swapping in a smaller or cheaper model. The selection and generation steps still use the same model. The savings come from removing all the model calls that were doing work regular code could do more directly.
A general-purpose coding agent is an extraordinarily capable reasoning and workflow engine. It’s also an expensive way to execute a procedure whose shape becomes known after introspecting on a few historical invocations.
I had historical traces from running the skill previously (it was running daily as an automation in Codex). Those traces showed what the agent actually did, including the planning, tool calls, branching, and state it needed along the way. I gave those traces, the original skill, and my writing about specialized harnesses to a powerful model. I asked it to identify which steps genuinely required an LLM and which had become stable enough to express as code, then build the specialized harness.
In other words, the natural-language skill served as a high-level specification, while the traces supplied the operational detail that was the result of reasoning by the model.
This is why I think the compiler analogy is useful. We begin with a flexible, high-level representation of intent. After the workflow has been exercised enough to reveal its real shape, we lower the stable parts into a more efficient representation.
The model is still used where necessary. The goal of this whole exercise is differentiating between the workflow parts that can be deterministic code and those that need language understanding, generation or reasoning. We haven’t tried to turn language understanding into a pile of brittle rules. Candidate selection depends on what the source says and whether it would make an interesting post. Drafting obviously benefits from language generation. Those remain model calls because they are model-shaped problems.
Writing the specialized harness from day one would have been premature. I didn’t yet know the exact workflow, which rules would matter, or where judgment would be required. The natural-language skill version let me discover those things by running the process repeatedly.
But keeping the entire workflow in natural language forever would mean paying the model to rediscover the same plan on every run.
Express the workflow as a natural-language skill.
Run it enough times to gather traces and refine the behavior.
Find the parts that have become stable and deterministic.
Compile those parts into code.
Keep LLM calls at the few points where semantic judgment matters.
There is a one-time cost to this compilation pass. I used a powerful model and a fair amount of context to inspect the traces and produce the new harness. But that cost is paid once. The compiled workflow can then run hundreds or thousands of times, saving tokens and time on every execution. This is standard optimization economics: spend once, amortize over repeated use.
I’ve grown to prefer local agents partly because they accumulate exactly this kind of history: skills, traces, state, and the artifacts of prior runs. That history doesn’t just help the agent remember. It gives us the raw material for making the harness itself better.
The long-term economics of agents will depend on this. If every recurring workflow keeps using a frontier model as its planner, state machine, and glue code forever, the cost and latency never really settle down. A specialized harness lets the model concentrate on the small fraction of the workflow that benefits from intelligence.
There’s another reason I don’t expect these techniques to come primarily from the big model vendors: their business is selling tokens . Preferably lots of tokens from their most powerful and expensive models.
There’s an old line: it’s very hard to convince someone of something when their salary depends on not understanding it. Incentives explain behavior and outcomes. A technique that preserves output quality while cutting token consumption by 94% goes straight against the current economics of a company whose revenue rises with token usage.
I’m not claiming that nobody inside those companies cares about efficiency. Obviously they do. Better inference, caching, and cheaper models make their products more useful. But there’s a difference between making each token cheaper and helping a customer discover that most of their workflow doesn’t need tokens at all.
At least right now, I don’t see much incentive for the major vendors to push that second idea hard. So it’s up to users and independent builders to find these optimization opportunities ourselves.
And that’s also the opportunity. There’s a large open space for founders building specialized harnesses, compilers, and tools that examine recurring agent workflows and move the deterministic parts into code. The value proposition is unusually concrete: keep the output quality, while cutting the cost and latency of producing it.
The model vendors will keep building more powerful engines. Someone still has to make sure we aren’t revving those engines to do simple, deterministic work.
Use natural language to discover the workflow. Use traces to understand it. Then compile what has crystallized.
Want to try it yourself? Grab the full Token Shrinker prompt and drop it into your agent.

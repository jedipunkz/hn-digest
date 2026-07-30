---
source: "https://raw.githubusercontent.com/langens-jonathan/blog/refs/heads/main/20260729_AgentsDefined_Part1of3"
hn_url: "https://news.ycombinator.com/item?id=49112281"
title: "Agents explained, what makes an LLM an agent?"
article_title: ""
author: "flowofcontrol"
captured_at: "2026-07-30T17:17:03Z"
capture_tool: "hn-digest"
hn_id: 49112281
score: 1
comments: 0
posted_at: "2026-07-30T16:30:10Z"
tags:
  - hacker-news
  - translated
---

# Agents explained, what makes an LLM an agent?

- HN: [49112281](https://news.ycombinator.com/item?id=49112281)
- Source: [raw.githubusercontent.com](https://raw.githubusercontent.com/langens-jonathan/blog/refs/heads/main/20260729_AgentsDefined_Part1of3)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T16:30:10Z

## Translation

タイトル: エージェントの説明、LLM がエージェントである理由は何ですか?

記事本文:
# MCP エージェントの説明: LLM を実際に「エージェント」にするもの
*パート 1/3 — MCP エージェントの構築とテスト*
「エージェント」は、すべてを意味し、何も意味しない言葉の 1 つになりました。それでは、接地しましょう。
大規模な言語モデルはそれ自体で、テキストをテキストに変換するという 1 つのことを行います。質問すると、答えが予測されます。何も調べたり、実際のシステムと照らし合わせて事実を確認したり、世界でアクションを起こしたりすることはできません。 「知っている」ものはすべてトレーニング時に凍結されます。
エージェントは、ツールを使用してモデルをループに接続することでこれを変更します。呼び出すことができる一連の関数 (ファイルの読み取り、API のヒット、データベースのクエリ) と、各ツールの結果をツールにフィードバックするループを与えると、何か別のことが起こります。モデルは情報が必要であると判断し、ツールを要求し、返された内容を読み取り、それで十分かどうか、または別のツールが必要かどうかを判断してから、初めて応答することができます。
このループ (*ツールを呼び出し、応答し、完了するまで繰り返す*) は、さまざまなフレームワークで装飾された、人々がエージェントと呼ぶすべてのシステムの中核です。しかし私には、ループだけが定義のすべてではないようです。停止条件のないループは暴走プロセスです。ツール ポリシーのないループは、セキュリティ インシデントが発生を待っていることを意味します。エージェントをセットアップするときに実際に何を構成するのかを突き止めようとしたところ、最終的に 2 つの別々の設定が必要になりました。
## エージェント定義とエージェント インスタンス
概念をクラスとそのインスタンスのように 2 つの部分に分割しました。 **エージェント定義**はテンプレートです。 **エージェント インスタンス** は、実行中の 1 つのインスタンスです。 (この定義と、最新の MCP 仕様でのストレス テストの方法について詳しく説明した別の記事を書きました (リンクは下部にあります)。
エージェント定義は次のもので構成されます。
* **ホスト ループ** — 上記のコア
* **システムコンテキスト** — ペルソナ、指示。ステータス

c と著者。
* **MCP サーバー セット** - このエージェントに存在するツール ソース
* **能力境界 / ツール ポリシー** - このエージェントが使用できるサーバーとツール、および承認が必要なもの。これはロールであり、ログインではありません。
* **LLM** — 使用している実際のモデル
* **終了制限** — 最大反復、トークンバジェット、停止条件
* **タスクコントラクト** — このエージェントが何を受け入れ、何を返さなければならないか
* **コンテキストサイズ戦略** — モデルのウィンドウ内で会話を維持する方法
エージェント インスタンスは、定義が実際の作業を取得したときに得られるものです。
* **目標** — ユーザーが実際に尋ねたこと
* **作業中のコンテキスト ウィンドウ** — 会話履歴、ツールの結果
* **資格情報** — 誰のトークン、誰の代理
* **消費された予算** — 経過した反復、消費されたトークン
対称性に注意してください。定義には *制限*、*契約*、および *能力* が含まれています。インスタンスは *カウンター*、*目標*、*資格情報*を保持します。クラスでは静的、インスタンスでは動的。この分割を念頭に置いてください。シリーズの残りの部分はこれに依存します。
## MCP が当てはまる場所
モデル コンテキスト プロトコル (MCP) は、ホストが最初にツールを検出して呼び出す方法という全体像の半分を標準化します。 MCP が登場する前は、すべての「エージェントとデータ ソース」の組み合わせには、Slack エージェント、GitHub エージェント、データベース エージェントなど、それぞれ手作業で接続された独自のオーダーメイドの統合が必要でした。 MCP はこれをプロトコルに変換します。どの MCP サーバーも、標準インターフェイスを介して型指定された検出可能なツールを公開し、MCP 対応ホストはカスタム グルー コードなしで任意の MCP サーバーに接続できます。
これは、アプライアンスではなく、プラグの標準化であると考えてください。エージェント (定義、インスタンス、およびそれらの間のループ) は依然としてホスト側に存在する必要があります。 MCP は、ホストが「この特定のツールとどのように通信するか」を再発明する必要がないことを意味します。

一回だけ。
## 実際にはどうなるか
私は、まさにこれを行う小さな MCP ホストを構築しています。1 つ以上の MCP サーバーに接続し、MCP サーバーが公開するあらゆるツールを集約し、LLM に代わって call-a-tool-or-answer ループを駆動します。興味深い設計上の決定は、幸せな道ではありませんでした。それは、SDK ツールを目に見えないように自動実行させたいという誘惑に抵抗することでした。ほとんどのエージェント フレームワークは、モデルにツール リストを渡すと、1 回の呼び出しで残りを不透明に処理するモードを提供します。
わざとそんなことはしないんです。ホストはループを手動で実行します。会話を送信し、ツール呼び出しリクエストの応答を検査し、それぞれをディスパッチし、結果をフィードバックし、繰り返します。エージェント定義の終了制限によって制限されているため、スパイラルは発生しません。これは、「SDK に任せる」バージョンよりもコードが数十行多くなっています。それによって得られるのは可視性です。すべてのツール呼び出し、すべてのラウンドトリップ、すべてのトークンは、「ここに質問があります」と「ここに答えがあります」の間にあるブラックボックスではなく、実際に見て測定できるものです。
より難しい質問をし始めると、その可視性が非常に重要であることがわかります。 「エージェントは機能するか」だけでなく、「このシステム プロンプトは実際にあのシステム プロンプトよりも優れているのか」、「安価なモデルでも十分なのか」、「大規模な場合のコストはいくらになるか」なども考慮します。これらは測定に関する質問であり、構築に関する質問ではありません。そして、それがこのシリーズの残りの部分の内容です。
次に、エージェントを構成するときに使用するレバー (これは偶然ではなく、まさに上記のエージェント定義のフィールドです) と、エージェントをプルすることで状況が良くなったのか悪くなったのかを示すメトリクスです。
*定義自体の詳細: https://dev.to/langensjonathan/what-is-an-agent-a-classinstance-definition-stress-tested-against-the-2026-07-28-mcp-spec-5akp*

## Original Extract

# MCP Agents, Explained: What Actually Makes an LLM an "Agent"
*Part 1 of 3 — building and testing MCP agents*
"Agent" has become one of those words that means everything and nothing. So let's ground it.
On its own, a large language model does one thing: it turns text into text. Ask it a question, it predicts an answer. It can't look anything up, check a fact against a live system, or take an action in the world. Everything it "knows" is frozen at training time.
An agent changes that by wiring the model into a loop with tools. Give it a set of functions it can call — read a file, hit an API, query a database — and a loop that feeds each tool's result back to it, and something different happens. The model can now decide it needs information, ask for a tool, read what comes back, decide whether that's enough or whether it needs another tool, and only then answer.
That loop — *call a tool, or answer, repeat until done* — is the core of every system people call an agent, dressed up in different frameworks. But to me it seems the loop alone is not the whole definition. A loop with no stop condition is a runaway process; a loop with no tool policy is a security incident waiting to happen. When I tried to pin down what I actually configure when I set up an agent, I ended up with two separate things.
## Agent definition and agent instance
I split the concept in two parts, like a class and its instances. An **agent definition** is the template; an **agent instance** is one running occurrence of it. (I wrote a separate article going deeper into this definition and how the newest MCP spec stress-tests it — link at the bottom.)
The agent definition consists of:
* **the host loop** — the core, as above
* **system context** — persona, instructions. Static and authored.
* **MCP server set** — which tool sources exist for this agent
* **capability boundary / tool policy** — which servers and tools this agent *may* use, and what needs approval. This is a role, not a login.
* **LLM** — the actual model we are using
* **termination limits** — max iterations, token budget, stop conditions
* **task contract** — what this agent accepts and what it must return
* **context-size strategy** — how it keeps the conversation inside the model's window
An agent instance is what you get when a definition picks up real work:
* **a goal** — what the user actually asked
* **a working context window** — conversation history, tool results
* **credentials** — whose tokens, on whose behalf
* **consumed budget** — iterations elapsed, tokens spent
Note the symmetry: the definition holds *limits*, *contracts* and *capabilities*; the instance holds *counters*, *goals* and *credentials*. Static in the class, dynamic in the instance. Keep this split in mind — the rest of the series leans on it.
## Where MCP fits in
The Model Context Protocol (MCP) standardizes one half of the picture: how a host discovers and calls tools in the first place. Before MCP, every "agent plus data source" pairing needed its own bespoke integration — a Slack agent, a GitHub agent, a database agent, each wired by hand. MCP turns that into a protocol: any MCP server exposes typed, discoverable tools over a standard interface, and any MCP-aware host can connect to any MCP server without custom glue code.
Think of it as the plug standardizing, not the appliance. The agent — definition, instance, and the loop between them — still has to exist on the host side. MCP just means the host doesn't have to reinvent "how do I talk to this particular tool" every single time.
## What this looks like in practice
I've been building a small MCP host that does exactly this: it connects to one or more MCP servers, aggregates whatever tools they expose, and drives that call-a-tool-or-answer loop on behalf of an LLM. The interesting design decision wasn't the happy path — it was resisting the temptation to let the SDK auto-run tools invisibly. Most agent frameworks offer a mode where you hand the model a tool list and it just handles the rest, opaquely, in one call.
I deliberately don't do that. The host drives the loop by hand: send the conversation, inspect the response for tool-call requests, dispatch each one, feed the results back, repeat — capped by the termination limits from the agent definition, so nothing spirals. It's a few dozen more lines of code than the "just let the SDK do it" version. What it buys you is visibility: every tool call, every round trip, every token is something you can actually see and measure, instead of a black box between "here's a question" and "here's an answer."
That visibility turns out to matter a lot once you start asking harder questions. Not just "does the agent work," but "is this system prompt actually better than that one," "is the cheaper model good enough," "what will this cost at scale." Those are measurement questions, not build questions — and they're what the rest of this series is about.
Next up: the levers you have when configuring an agent — which, not coincidentally, are exactly the fields of the agent definition above — and the metrics that tell you whether pulling one made things better or worse.
*Deeper dive into the definition itself: https://dev.to/langensjonathan/what-is-an-agent-a-classinstance-definition-stress-tested-against-the-2026-07-28-mcp-spec-5akp*

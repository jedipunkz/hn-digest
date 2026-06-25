---
source: "https://www.multiplayer.app/blog/how-to-curate-observability-data-for-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48675760"
title: "How to curate observability data for AI agents"
article_title: "How to curate observability data for AI agents"
author: "argoeris"
captured_at: "2026-06-25T16:45:32Z"
capture_tool: "hn-digest"
hn_id: 48675760
score: 1
comments: 0
posted_at: "2026-06-25T16:25:41Z"
tags:
  - hacker-news
  - translated
---

# How to curate observability data for AI agents

- HN: [48675760](https://news.ycombinator.com/item?id=48675760)
- Source: [www.multiplayer.app](https://www.multiplayer.app/blog/how-to-curate-observability-data-for-ai-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T16:25:41Z

## Translation

タイトル: AI エージェントの可観測性データをキュレーションする方法
説明: ほとんどのデバッグ エージェントが失敗するのは、モデルが間違っているためではなく、入力されるデータがマシンで使用できる状態になっていないためです。データキュレーションが実際にどのように行われるかを次に示します。

記事本文:
AI
AI エージェントの可観測性データをキュレーションする方法
ほとんどのデバッグ エージェントが失敗するのは、モデルが間違っているためではなく、入力されるデータがマシンで使用できる状態になっていないためです。データキュレーションが実際にどのように行われるかを次に示します。
私たちがマルチプレイヤーのデバッグ エージェントを構築し始めたとき、ほとんどの人が犯すのと同じ間違いを犯しました。私たちはコーディング エージェントに可観測性データへのアクセスを許可し、何が関連しているかを判断してくれることを期待していました。
そうではありませんでした。エージェントは間違ったツールを呼び出し、間違ったシグナルを追跡し、一見もっともらしい修正を作成しましたが、本番環境では失敗しました。私たちは最先端のモデルを使用していましたが、キュレーションやフィルタリングを行わずに生の可観測性データを渡していました。後で、私たちはノイズをルーティングしているだけであることに気づきました。
以下は、可観測性データが AI エージェントの動作に適するようになる前に、実際にどのような処理を行う必要があるかについて私たちが学んだことです。
可観測性データの信号対雑音比は、AI エージェントに供給できるデータ タイプの中で最も悪いものの 1 つです。
1 つの運用上の問題には、数十のサービスにわたる数百のスパン、数千のログ行、リクエストとレスポンスのペイロードの欠落、編集されたヘッダー、クロックのずれたタイムスタンプ、相互に関連付けられたことのないツール間で分散されたイベントが含まれる場合があります。人間がこの問題をデバッグすると、何年にもわたるコンテキストが得られます。つまり、どのサービスにノイズが多いか、どのログが重要か、どのタイムスタンプを信頼すべきか、そして問題がスタック内のどこにあるかがおおよそわかっています。彼らはシステムを理解しているので、ノイズを乗り越えることができます。
エージェントはすべてを同じ重みで見ます。ガベージ スパンは、実際に障害を示す 1 つのスパンと同じ注目を集めます。エージェントが有益な質問をするまでに、何千ものログ行が処理されます。そして、コンテキスト ウィンドウは有限で高価であるため、y

問題を正しく理解する前に予算を使い切ってしまいます。
これはデータ準備の問題です。そして、それはエージェント自体によってではなく、データがエージェントに到達する前に解決される必要があるものです。
データキュレーションの実際の意味
AI エージェントのデータ キュレーションを、ほとんどのエンジニアリング チームが最終的に行うことになる要約や圧縮と混同しないでください。
実際には、これは、生の可観測性データを、エージェントが正しく推論できる、構造化されスコープが設定されたコンテキスト豊富なパッケージに変換するプロセスです。これは、何を含めるか、何を除外するか、関連するシグナルをグループ化する方法、エージェントが問題を理解するためにどのような追加コンテキストが必要かなど、一連の慎重な決定を下すことを意味します。
Multiplayer では、データがコーディング エージェントに到達する前に、これを 4 つの段階で行います。
ステージ 1: グループ化して積極的に関連付けます
生の可観測性データを使用して最初に行うことは、関連するイベントをグループ化し、サービスの境界を越えてそれらを関連付けることです。
通常、1 つのバグが多くのセッション、環境、サービスにわたって表面化します。グループ化しないと、それぞれの発生が個別の問題のように見えます。また、相関関係がなければ、エージェントはフロントエンドでのユーザーのアクションとバックエンドの奥深くにある障害を結びつける因果関係を認識できません。
当社は積極的に相関付けを行います。ユーザー インタラクション、セッション メタデータ、ネットワーク リクエスト、バックエンド トレース、ログ イベントは、何かが起こる前に 1 つのタイムラインにまとめられます。エージェントは、14:32:01 のクリックによってカスケードが発生し、14:32:04 にバックエンド ログに記録されたことを確認する必要があります。タイムスタンプだけからそれを推測することはできません (特に実際の負荷やクロック スキューがある場合)。相関関係は、エージェントが認識する前にデータ構造に組み込む必要があります。
この段階で重複排除も行います。同じバグが複数の場所で発生する

100 のユーザー セッションが 100 の個別の信号ではなく、1 つの問題になります。これはコスト管理と品質管理の両方によるものです。重複排除されグループ化されたデータを処理するエージェントは、1 つの問題に対して 1 つの PR を生成します。グループ化されていない生のデータを操作するエージェントは、同じ問題に対して数十の PR を生成したり、トークンを不必要に焼き尽くしたり、同じ根本的な障害からの矛盾するシグナルを調整しようとして混乱したりします。
ステージ 2: エージェントにルーティングする前に修復可能性を評価する
すべての問題がコーディング エージェントにルーティングされる価値があるわけではありませんし、すべての問題がコーディング エージェントで解決できるわけでもありません。
コーディング エージェントに何かが届く前に、専用のエージェントを通じて修正可能性の評価を実行します。これは明確な根本原因を持つ決定的で再現可能な障害ですか?それとも、診断するには人間の判断が必要な、断続的な環境固有の問題でしょうか?
これにはいくつかの理由があります。まず、コーディング エージェントは、正しく解決するのに十分なコンテキストを持たない問題に関して最悪の出力を生成します。これは、多くの場合、最も困難で最も断続的なバグです。人間の監視なしにこれらをコーディング エージェントにルーティングすると、トークンが無駄になり、有効に見える修正が生成されます。
2 番目に、修正可能性のスコアリングにより優先順位を付けることができます。修正可能性の高い問題 (明確な根本原因、決定論的な再現、広範囲にわたる影響) は、コーディング エージェントに直ちに伝えられます。修正可能性が低い問題には、厳選されたコンテキストがすでに添付されているため、人間によるレビューのためにフラグが立てられます。
目標は、実際に人間の判断が必要な場合には人間を常に把握し、それ以外のすべてを自動化された修正サイクルに通すことです。
ステージ 3: リリース コンテキストとメタデータを追加する
生の可観測性データは、何かが壊れたことを示しますが、壊れた原因となった変更はわかりません。
データがコーディング エージェントに到達する前に、自動的に広告が送信されます。

d リリース コンテキスト: ビルド情報、デプロイメント タイムスタンプ、最近のコミット、障害に関係する各サービスの特定のバージョン。虫は真空状態では発生しません。これらは通常、Git 履歴の特定の時点、多くの場合特定のコミット、影響を受けるコード パスに影響を与えた特定の変更によって導入されます。
このコンテキストを無視して修正を作成するコーディング エージェントは、バグの因果関係を推測していることになります。リリース メタデータが添付されていると、エージェントは失敗信号をそれを導入した変更に結び付けることができます。これにより、修正の品質が大幅に変わります。「このエラー ケースを処理するパッチは次のとおりです」から、「導入された変更内容と修正方法は次のとおりです」になります。
また、サービスのメタデータ、環境情報、およびエージェントが動作しているシステムを理解するのに役立つ関連する構成コンテキストも追加します。カスタム サービス名のマッピング (例: "payment-service"、"svc-payments"、および "payments_v2" はすべて同じものを指します) はここで解決されるため、エージェントは 3 つの名前を 3 つのエンティティとして扱いません。
ステージ 4: マシンで使用できるようにフォーマットして要約する
最終段階は、ほとんどのチームが完全にスキップする段階であり、デバッグ エージェントのパフォーマンスの多くが考慮されないままになります。
生の可観測性データは、JSON ペイロード、ネストされたスパン構造、内部フォーマット規則を使用したログ行など、人間向けにフォーマットされています。これらは、システムを理解し、ダッシュボードを見ている人が読めるように設計されています。
データはコーディング エージェントに届く前に再フォーマットされます。スパンは、ネストされた JSON から、実行パスを記述する構造化されたナラティブに変換されます。ログ行は、実際に障害ウィンドウに関連するものにフィルタリングされ、タイムラインが読みやすいように再フォーマットされます。リクエストと再

レスポンス ペイロード (ほとんどの可観測性ツールはデフォルトで削除し、最も有用なデバッグ信号であるため特にキャプチャします) は、それらが関連する理由を説明するコンテキストとともに組み込まれます。
また、「5 歳のように説明してください」と呼ばれる問題の概要も作成します。目標は、インシデント コールに参加したばかりの開発者に説明するような方法でコーディング エージェントを迅速に説明できるようにすることです。つまり、何が壊れたのか、これがいつ始まったのか、最近何が変わったのか、スタック内のどこに障害が存在するのか、エラーが表面化したときにどのように見えるのかを説明します。
実際にはどうなるか
マルチプレイヤーのデバッグ エージェントの V1 と V2 の違いは、ほぼ完全にキュレーション層にありました。
V1 は私たちの API をミラーリングし、エージェントが使用できる多くのツールを提供しました。エージェントは間違ったツールを呼び出し、間違ったパラメータを使用し、トークンを焼き尽くし、実際の根本原因を見逃した PR を生成しました。モデルは問題ではありませんでした。データアクセスパターンが問題でした。
V2 には、エージェントが問題を理解するために必要なすべてのものを厳選し、関連付けてフォーマットしたパッケージを返す 1 つの主要ツールがありました。エージェントは適切なタイミングで適切なものに電話をかけ、より詳しいコンテキストが必要な場合には集中的にフォローアップの質問をし、運用環境に耐えられる修正を作成しました。
違いを生んだのは、キュレーション層 (グループ化、重複排除、修正可能性評価、リリース コンテキスト、フォーマット、問題の概要) でした。
0:00
/ 0:40
1×
自分自身に問いかける質問
ほとんどのデバッグ エージェントと MCP サーバーは、「AI に可観測性データへのアクセスを提供するにはどうすればよいですか?」という質問に答えるために構築されています。
正しい質問は、「出荷する価値のある修正を作成するために、エージェントはこの特定の問題について何を理解する必要があるか?」です。
これらの質問は、まったく異なるアーキテクチャにつながります。モミ

これは生データの漏洩につながります。エージェントにすべてへのアクセスを許可し、何が関係しているかを判断させます。 2 つ目はキュレーションにつながります。データがエージェントに届く前に、データをマシンの使用に適合させる作業を実行します。
あなたが現在持っている可観測性データは人間のために構築されました。サンプリング、集約、サイロ化され、ダッシュボード用にフォーマットされます。変換せずにコーディング エージェントに直接送信することが、ほとんどのデバッグ エージェントが正しく見える出力を生成するのに本番環境では失敗する原因となります。
ターミナルに 1 回コピー/ペーストすると、デバッグ エージェントが実行されます。
npm install -g @multiplayer-app/cli && マルチプレイヤー
マルチプレイヤーがオープンソースになりました
マルチプレイヤー デバッグ エージェントは、MIT のもとでオープン ソースです。その理由と、それがどのように使用されるかを説明します。
ダッシュボードの死: エージェント AI が従来の可観測性ツールで窒息している理由
ダッシュボード、サンプリング、データ レイクは人間によるデバッグ用に構築されました。 AI エージェントのバグから修正へのループを閉じるには、実行時データの収集方法と関連付け方法を再考する必要があります。
開発者向けのデバッグ エージェント: ローカルで実行され、PR の無駄を排除します。
マルチプレイヤー デバッグ エージェントは、コーディング エージェントを使用する開発者専用に構築されています。可観測性ツールが見逃したすべてのデータを捕捉し、バグの特定からバグ修正までのプロセス全体を管理します。
カスタム ロギングの (そうではない) 隠れたコスト
カスタム ログは技術的にはすべてをキャプチャできますが、実際にはほとんどキャプチャできません。時間の経過とともにカバレッジは低下し、外部 API は忘れ去られ、インシデント発生時には「誰かがこれを記録したのか?」という疑問が残ることになります。デバッグの代わりに。自動キャプチャはこれを解決します。

## Original Extract

Most debugging agents fail not because the model is wrong, but because the data going in is not ready for machine consumption. Here's what data curation actually looks like in practice.

AI
How to curate observability data for AI agents
Most debugging agents fail not because the model is wrong, but because the data going in is not ready for machine consumption. Here's what data curation actually looks like in practice.
When we started building Multiplayer's debugging agent, we made the same mistake almost everyone makes. We gave our coding agent access to observability data and expected it to figure out what was relevant.
It didn't. The agent called the wrong tools, chased the wrong signals, and produced fixes that looked plausible but failed in production. We were using state-of-the-art models, but we were handing them raw observability data without any curation or filtering. We later realized that we were just routing them noise.
What follows is what we learned about what you actually have to do with observability data before it's fit for an AI agent to act on.
Observability data has one of the worst signal-to-noise ratio of any data type you could feed an AI agent.
A single production issue might involve hundreds of spans across a dozen services, thousands of log lines, missing request and response payloads, redacted headers, clock-skewed timestamps, and events distributed across tools that have never been correlated with each other. A human debugging this issue brings years of context: they know which services are noisy, which logs matter, which timestamps to trust, and roughly where in the stack the problem lives. They navigate the noise because they understand the system.
An agent sees everything with equal weight. Garbage spans get the same attention as the one span that actually shows the failure. Thousands of log lines get processed before the agent can ask a useful question. And because context windows are finite and expensive, you burn through your budget before you've even framed the problem correctly.
This is a data preparation problem. And it's one that has to be solved before the data reaches the agent, not by the agent itself.
What data curation actually means
Data curation for AI agents shouldn’t be confused with summarization or compression, which is what most engineering teams end up doing.
In actuality, it's the process of transforming raw observability data into a structured, scoped, context-rich package that an agent can reason about correctly. That means making a series of deliberate decisions: what to include, what to exclude, how to group related signals, and what additional context the agent needs to understand the problem.
At Multiplayer, we do this in four stages before any data reaches a coding agent.
Stage one: group and correlate aggressively
The first thing we do with raw observability data is group related events and correlate them across service boundaries.
A single bug will typically surface across many sessions, environments, and services. Without grouping, each occurrence looks like a separate issue. And without correlation, the agent can't see the causal chain that connects a user action on the frontend to a failure deep in the backend.
We correlate aggressively: user interactions, session metadata, network requests, backend traces, and log events get tied together into a single timeline before anything else happens. The agent needs to see that the click at 14:32:01 caused the cascade that showed up in the backend logs at 14:32:04. It can't infer that from timestamps alone (especially under any real load or clock skew). The correlation has to be built into the data structure before the agent sees it.
We also deduplicate at this stage. The same bug appearing across a hundred user sessions becomes one issue, not a hundred separate signals. This is both because of cost and quality management. An agent acting on deduplicated, grouped data produces one PR for one issue. An agent acting on raw, ungrouped data produces dozens of PRs for the same issue, burns through tokens unnecessarily, or gets confused trying to reconcile conflicting signals from the same underlying failure.
Stage two: assess fixability before routing to the agent
Not every issue is worth routing to a coding agent, and not every issue is something a coding agent can fix.
Before anything reaches the coding agent, we run a fixability assessment through a dedicated agent. Is this a deterministic, reproducible failure with a clear root cause? Or is it an intermittent, environment-specific issue that requires human judgment to diagnose?
This matters for a few reasons. First, coding agents produce their worst outputs on problems they don't have enough context to solve correctly, which are often the hardest, most intermittent bugs. Routing those to a coding agent without human oversight wastes tokens and produces plausible-looking fixes that don't hold.
Second, fixability scoring lets you prioritize. High-fixability issues (clear root cause, deterministic reproduction, well-scoped impact) go to the coding agent immediately. Lower-fixability issues get flagged for human review with the curated context already attached.
The goal is to keep humans in the loop where human judgment is actually needed, and route everything else through the automated fix cycle.
Stage three: add release context and metadata
Raw observability data tells you that something broke, but it doesn't tell you what changed that caused it to break.
Before the data reaches the coding agent, we automatically add release context: build information, deployment timestamps, recent commits, the specific version of each service involved in the failure. Bugs don't appear in a vacuum. They're usually introduced at a specific point in the git history, often in a specific commit, often by a specific change that touched the affected code path.
A coding agent producing a fix without this context is guessing about the causal history of the bug. With release metadata attached, the agent can connect the failure signal to the change that introduced it. That changes the quality of the fix significantly: it goes from " here's a patch that handles this error case " to " here's what the change introduced and here's how to correct it. "
We also add service metadata, environment information, and any relevant configuration context that helps the agent understand the system it's operating in. Custom service name mappings (e.g. "payment-service," "svc-payments," and "payments_v2" all referring to the same thing) get resolved here so the agent isn't treating three names as three entities.
Stage four: format and summarize for machine consumption
The final stage is the one that most teams skip entirely, and it's where a lot of debugging agent performance gets left on the table.
Raw observability data is formatted for humans: JSON payloads, nested span structures, log lines with internal formatting conventions. These are designed to be readable by someone who understands the system and is looking at a dashboard.
We reformat data before it reaches the coding agent. Spans get converted from nested JSON into a structured narrative that describes the execution path. Log lines get filtered to the ones that are actually relevant to the failure window and reformatted to make the timeline legible. Request and response payloads (which most observability tools strip out by default, and which we capture specifically because they're the most useful debugging signal) get included with the context that explains why they're relevant.
We also produce an issue summary that we call "explain it like I'm 5". The goal is to bring the coding agent up to speed the way you'd brief a developer who's just joined an incident call: here's what broke, here's when it started, here's what changed recently, here's where in the stack the failure lives, here's what the error looks like when it surfaces.
What this looks like in practice
The difference between V1 and V2 of Multiplayer's debugging agent was almost entirely in the curation layer.
V1 mirrored our API and gave the agent a lot of tools to work with. The agent called the wrong tools, used the wrong parameters, burned through tokens, and produced PRs that missed the actual root cause. The model wasn't the problem. The data access pattern was the problem.
V2 had one main tool that returned a curated, correlated, formatted package of everything the agent needed to understand the issue. The agent called the right thing at the right time, asked focused follow-up questions when it needed more context, and produced fixes that held up in production.
What made the difference was the curation layer: grouping, deduplication, fixability assessment, release context, formatting, and issue summary.
0:00
/ 0:40
1×
The question to ask yourself
Most debugging agents and MCP servers are built to answer the question: "How do I give the AI access to my observability data?"
The right question is: "What does the agent need to understand about this specific issue in order to produce a fix worth shipping?"
Those questions lead to very different architectures. The first leads to raw data exposure: give the agent access to everything and let it figure out what's relevant. The second leads to curation: do the work of making the data fit for machine consumption before it ever reaches the agent.
The observability data you have right now was built for humans. It's sampled, aggregated, siloed, and formatted for dashboards. Sending it directly to a coding agent without transformation is the reason most debugging agents produce output that looks right and fails in production.
One copy/paste in your terminal and the debugging agent is running:
npm install -g @multiplayer-app/cli && multiplayer
Multiplayer is now open source
The Multiplayer debugging agent is open source under MIT. Here's why, and what it means for how you use it.
The death of the dashboard: why agentic AI is choking on legacy observability tools
Dashboards, sampling, and data lakes were built for human debugging. Closing the bug-to-fix loop for AI agents requires rethinking how runtime data is collected and correlated.
The debugging agent for developers: runs locally and eliminates PR slop
The Multiplayer debugging agent is purpose-built for developers working with coding agents. It captures all the data observability tools miss and manages the whole process from bug identified to bug fixed.
The (not so) hidden cost of custom logging
Custom logging can technically capture everything, but in practice, it rarely does. Coverage degrades over time, external APIs get forgotten, and during incidents, you're left asking "did anyone log this?" instead of debugging. Automatic capture solves this.

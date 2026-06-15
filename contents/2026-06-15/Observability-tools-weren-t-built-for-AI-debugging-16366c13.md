---
source: "https://leaddev.com/ai/observability-tools-werent-built-for-ai-debugging"
hn_url: "https://news.ycombinator.com/item?id=48541162"
title: "Observability tools weren't built for AI debugging"
article_title: "Observability tools weren’t built for AI debugging - LeadDev"
author: "argoeris"
captured_at: "2026-06-15T14:14:08Z"
capture_tool: "hn-digest"
hn_id: 48541162
score: 1
comments: 0
posted_at: "2026-06-15T13:41:56Z"
tags:
  - hacker-news
  - translated
---

# Observability tools weren't built for AI debugging

- HN: [48541162](https://news.ycombinator.com/item?id=48541162)
- Source: [leaddev.com](https://leaddev.com/ai/observability-tools-werent-built-for-ai-debugging)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T13:41:56Z

## Translation

タイトル: 可観測性ツールは AI デバッグ用に構築されていない
記事タイトル: 可観測性ツールは AI デバッグ用に構築されていない - LeadDev
説明: AI は、可観測性ツールがバグを捕捉するよりも速くバグを出荷しています。解決策はモデルをよりスマートにすることではなく、そもそも適切なデータを収集することです。

記事本文:
可観測性ツールは AI デバッグ用に構築されていない - LeadDev
コンテンツにスキップ
LeadDev.com を検索
行く
今後のイベント
ログイン ログイン
ニュースレター
Latest news in your inbox
Find content specific to your role
可観測性ツールは AI デバッグ用に構築されていない
無料の LeadDev.com アカウントを登録する必要がある前に、今月読む記事が 1 つ残っています。
Estimated reading time: 10 minutes
AI のデバッグにはモデルの問題ではなく、データの問題があります。 Smarter models won’t fix bad observability.ゴミが入って、ゴミが出る。
あなたはテレメトリに溺れていますが、信号には飢えています。
間違ったデータを大規模に収集するのではなく、事前に適切なデータを収集します。セッションベースの相関コレクションは、AI エージェントに必要な完全なコンテキストを提供します。
ほとんどのエンジニアリング チームは AI を使用して、これまでよりも迅速にコードを作成しています。ただし、現在ではバグも同じ速度で出荷されています。
実際のワークフローはエンドツーエンドでどのように見えるかは次のとおりです。
[Auto-instrument everything via OTel]
→ [Collector samples/filters some]
→ [Store remaining data]
→ [Developer notices bug]
→ [手動でコピー＆ペーストエラーまたはMCPサーバー経由でクエリ]
→ [AI gets incomplete/noisy context]
→ [AI suggests fix based on partial data]
→ 【人のレビュー】
→ [Code “looks plausible”]
→［配備］
→ [Discover edge case not addressed by AI]
→ [More bugs in production]
このワークフローが破綻する箇所がいくつかあります。最終的に、根本的な原因はデータとコンテキストの問題です。それを修正するまでは、AI の能力がいくらあってもこのギャップを埋めることはできません。
開発者として、すぐに使用できる可観測性ツールからの適切なデータがない場合、それを AI エージェントに渡すという魔法の解決策はありません。問題は、「適切なデータ」とは実際にはどのようなものなのか、そして現在のツールがそれを提供するのにどれだけ遠いのかということです。
現在のob

servability tools are pretty good at gathering:
システムレベルのメトリック: 稼働時間、遅延、エラー率。
サービス間のトレース: どのサービスがどのサービスを呼び出したか、各ホップにかかった時間。
ただし、複雑な問題をデバッグする上で最も重要な部分に常にギャップが残ります。
サンプリングされたトレースとログ: ストレージ コストを制御するために、積極的に間引かれます。
セッションベースのテレメトリ: ユーザー ジャーニーとバックエンドの動作に相関関係はありません。
リクエスト/レスポンス ペイロード: 内部サービスと外部依存関係には欠落しています。
エンジニアリングに関する洞察を毎週受け取り、リーダーシップのアプローチをレベルアップします。
この情報はデフォルトでは提供されないため、AI ツールは現実の部分的なビューで動作することになります。これらは依然として真に有用であり、一般的な問題に対してパターンマッチングを行い、範囲が広く、よく理解されている問題を適切に解決します。
AI のデバッグは、複雑な分散システムの問題に遭遇した瞬間に機能不全に陥ります。実際に何が問題だったのかを理解するには、完全なランタイム コンテキストが必要なタイプです。そのような場合でも、エンジニアはローカルの開発環境を起動し、すべてのツールを導入し、手動でデバッグする必要があります。
現在の可観測性ツールは AI デバッグに関しては根本的に機能しておらず、その差はさらに悪化しようとしています。
Harness の State of Software Delivery 2025 レポートによると、開発者の 67% は現在、AI ツールを使用し始める前よりも AI によって生成されたコードのデバッグに多くの時間を費やしています。 Stack Overflow の 2025 年の開発者調査では、同じ話を別の角度から伝えています。開発者の 66% が、AI によって生成された「ほぼ正しい」コードの修正により多くの時間を費やしていると回答し、45% が AI ツール全体に対する最大の不満として挙げています。
アプリケーションを構築するために AI コーディング ツールを採用したり、エージェント AI を実装したりするチームが増えたらどうなるかを考えてみましょう。 The volume of code goes up. T

発送のスピードが上がります。それでも、その下にある可観測性ツールは変わっていません。結果？バグの増加、診断の困難、そして可観測性ツールの（不必要な）増殖。
実際に修正する様子は次のとおりです。
テレメトリ収集に対する企業のアプローチは予防的なものです。彼らは、問題が起こったときに何が必要になるかわからないので、あらゆるものを集めます。オンデマンドでの計測（問題によりギャップが明らかになった後）は費用と時間がかかるため、抵抗を最小限に抑える方法は、事前に最も広い網を張ることです。
問題は、広い網を張ることと、実際に適切な魚を捕獲することは全く別のことであるということです。その結果、企業は大量の汎用テレメトリを溜め込み、それに伴うストレージ コストを捻出しようとしています。
「可観測性がインフラストラクチャの予算を食いつぶしている」という苦情は、お客様からの電話で定期的に聞かれますが、これは異常ではなく、通常のことです。
OpenTelemetry はこれを改善するどころか悪化させました。その民主化された手段は標準化にとって真の勝利ですが、それは誤ってコモンズの悲劇を生み出しました。すべてのチームがすべてを自動計測することが簡単にできる場合は、誰もがそうします。データが実際に役立つかどうかを立ち止まって尋ねる人は誰もいません。
最終的に、システムはテレメトリに溺れながらも信号に飢えている状態になります。 Dash0 の製品責任者である Michele Mancioppi 氏は、エージェントの可観測性に関する最近の LeadDev パネルで率直にこう述べました:「KubeCon で私は、可観測性ツールに取り込まれるログのうち、致命的なエラーは 1% 未満であると共有しました。残りはおそらくすべてお金の無駄です。」
ここで AI のデバッグが最初の壁にぶつかります。 AI エージェントに可観測性データを指定するだけで、それが機能することを期待することはできません。理由は 2 つあります。
まず、コストがかかります。 AIモデルc

処理されるトークンごとに高額な料金が発生し、テレメトリの消防ホースを供給するのにコストがかかります。第二に、技術的な厳しい制限があります。 AI モデルには有限のコンテキスト ウィンドウがあります。彼らに何を与えるかを選択する必要があります。生の OTel トレースとログを提供すると、エージェントはノイズに紛れてしまうか、そのスライスがサンプリングされているか、最初から収集されなかったため、実際に重要な 1% を見逃してしまいます。
重要なデバッグデータが欠落している
AI エージェントが無関係なデータに溺れていないとしても、依然として全体像を把握しているわけではありません。必要なコンテキストは、接続されていないツールに散在しているか、まったく収集されていません。たとえば:
リクエスト/レスポンスのペイロードは、プライバシーとコストの問題のため、デフォルトではキャプチャされません。アプリケーション プログラミング インターフェイス (API) 呼び出しが行われたことはわかりますが、どのようなデータが送信されたのか、何が返されたのかはわかりません。
多くの場合、ヘッダー (認証トークン、ルーティング情報、カスタム メタデータ) はセキュリティ上の理由から編集され、重要なデバッグ コンテキストが削除されます。
外部 API 交換はブラックボックスです。分散トレースはシステム境界で停止します。バックエンドが Stripe、Twilio、または AWS を呼び出すと、呼び出しが行われたこととそのにかかった時間が表示されますが、送信した内容や返された内容は表示されません。
セッションの相関関係はせいぜい断片化されています。フロントエンド監視 (RUM) は 1 つのツールにあり、バックエンドの可観測性は別のツールにあります。 「このユーザーが A をクリックし、次に B をクリックし、その後バックエンドが失敗しました」と接続するには、プラットフォーム間で手動でステッチする必要があります。
現在の可観測性ツールは、何かが起こったこととそれがどこで起こったかを示すことはできますが、スタック全体で点を自動的に結び付けるのに苦労しています。完全な相関関係を得るには大規模な機器が必要であり、大規模になると法外な費用がかかります。特に実際のデータ (ペイロード、ヘッダー) をキャプチャする場合

s、フルコンテキスト）システム内を流れます。
しばらくの間、すべてのデータが存在し、欠落しているものは何もないと仮定しましょう。さまざまなツールにわたってサイロ化されているだけです。
理論的には、AI エージェントは複数の可観測性プラットフォーム (Sentry のフロントエンド エラー、Datadog のバックエンド トレース、Stripe の外部 API ログ、LogRocket のセッション リプレイ) をクエリし、それらを自動的に関連付けることができます。実際には、このようなことはほとんど起こりません (少なくとも、まだ)。
Frontier AI が出血の機密データをモデル化
AI エージェントが明らかにするアーキテクチャのギャップ
各ツールには異なる API、認証方法、データ形式があります。さらに重要なのは、イベントを結び付ける相関キー (リクエスト ID、トレース ID、セッション ID) がツールの境界を越えて伝播することはほとんどありません。エージェントは、タイムスタンプに基づいて、どのフロントエンド エラーがどのバックエンド トレースに対応するかを推測する必要があります。これは、実際の負荷やクロック スキューが発生するとエラーが発生しやすいアプローチです。
そのため、すべてのツールへの API アクセスが完璧であっても、収集時にデータがすでに関連付けられていない限り、エージェントは依然として全体像を把握できません。
そのため、開発者は相関作業を手動で行う必要があります。つまり、ツール間でリクエスト ID をコピーし、ダッシュボード間でタイムスタンプを照合し、ストーリーをつなぎ合わせてから、分析のために関連するフラグメントを AI ツールに貼り付けるだけです。
人間によるコードレビューでは AI コード生成が追いつかない
エンジニアなら誰でも知っているように、コードを読むことは書くことよりも難しいです。他の人の思考プロセスをリバースエンジニアリングし、彼らが特定の決定を下した理由を理解し、彼らが見逃しているかもしれないエッジケースを特定する必要があります。数週間後に自分のコードをレビューすることさえ、それほど簡単ではありません。私たちの記憶には誤りがあり、ドキュメントが必要なほど完全であることはほとんどありません。
ここで、AI アシスタントよりも速くコードを書く AI アシスタントを追加します。

人間にはできません。 AI は、30 秒以内に 100 行のもっともらしいコードを生成できます。同じ 100 行を徹底的にレビューするには 15 ～ 20 分かかる場合があります。数学は機能しません。チームはレビューがボトルネックになっているか、作業を見逃してしまいます。
そして、物事はすり抜けてしまいます。 AI が生成したコードは、モデルが想像した幸せな道に沿って機能することがよくありますが、経験豊富な開発者のように防御的に考えることはありません。エッジケース、エラー状態、または予期しない入力について推論するのではなく、以前に確認された内容に基づいてパターン マッチングを行います。
結果？コード行ごとのバグ率が一定であったとしても、より多くのコードをより速く出荷できることになり、絶対的な意味でバグが増えることになります。 AI が生成したコードはきれいにコンパイルされ、構造的に正しいように見えるため、レビューでバグを発見するのが難しくなります。コードは問題ないようです。 AI が考慮しなかった条件下では、本番環境で失敗するだけです。
ここで、「AI によって生産性が 10 倍向上する」という物語が破綻します。 「最初の試行でコンパイルされる」ということは、コードが正しいという意味ではありません。これは、バグが実稼働環境に到達するまでデバッグ作業を延期したことを意味します。
可観測性のギャップがこの問題をさらに悪化させます。完全なランタイム コンテキスト (実際のペイロード、外部 API 応答、セッション状態) がなければ、時間のプレッシャーの下、不完全な情報に基づいて、自分で作成していないロジックをリバース エンジニアリングすることになります。
可観測性業界はこれらの問題を認識し始めていますが、提案されている解決策はパズルのさまざまな部分に対処しています。
アプローチ 1: コレクター層でのスマート フィルタリング
一部のベンダーは、インテリジェントな収集パイプラインを提唱しています。
【アプリ】
→[すべて送信]
→ 【スマートAIコレクター】
→ [何を保持/削除/エスカレーションするかを決定する]
→［ストレージ］
→ 【AIがクリーンなデータを取得】
このアイデアは、コレクター層で AI を使用することです。

o ノイズがストレージに到達する前にフィルタリングして除去します。致命的なエラーを保持し、ヘルスチェックを削除し、重要と思われるものに基づいてインテリジェントにサンプリングします。
これにより、「無関係なデータが多すぎる」問題が解決されます。すべてを保管する必要がないため、保管コストが削減されます。 AI エージェントはノイズが除去されているため、よりクリーンなデータを取得できます。
ただし、「重要なデバッグ データの欠落」問題は解決されず、取り込みコストにも対処できません。また、重要でないと思われたためにコレクターがリクエストを削除し、そのリクエストが運用インシデントを引き起こしたリクエストであることが判明した場合、データは失われます。
アプローチ 2: 可観測性プラットフォームに AI エージェントを組み込む
他のベンダーは、AI を自社のプラットフォームに直接組み込んでいます。エージェントはツール内に常駐し、そこで収集されたすべてのテレメトリ データにネイティブ アクセスできます。
【アプリ】
→【全て集める】
→ [可観測性プラットフォームストレージ]
→ [組み込み AI エージェントがプラットフォーム データをクエリ]
→ 【AIが入手可能なデータを使って質問に答える】
これにより、データ相関の問題が部分的に解決されます。ベンダーがセッション リプレイ、メトリクス、ログ、トレースを 1 か所に収集している場合、AI エージェントは複数の API をクエリすることなく、それらすべてを相関させることができます。
しかし、「無関係なデータが多すぎる」という問題に戻りました。欠落データは依然として欠落しており、言うまでもなく、ベンダー ロックインにも対処しなければなりません。
アプローチ 3: セッションベース

[切り捨てられた]

## Original Extract

AI is shipping bugs faster than observability tools can catch them. The fix isn’t smarter models, it’s collecting the right data in the first place.

Observability tools weren’t built for AI debugging - LeadDev
Skip to content
Search LeadDev.com
Go
Upcoming events
Login Login
Newsletters
Latest news in your inbox
Find content specific to your role
Observability tools weren’t built for AI debugging
You have 1 article left to read this month before you need to register a free LeadDev.com account.
Estimated reading time: 10 minutes
AI debugging has a data problem, not a model problem . Smarter models won’t fix bad observability. Garbage in, garbage out.
You’re drowning in telemetry but starving for signal.
Collect the right data upfront, not the wrong data at scale . Session-based, correlated collection gives AI agents the complete context they need.
Most engineering teams are using AI to write code faster than ever. However, now they are also shipping bugs with equal speed.
Here’s what that workflow actually looks like end to end:
[Auto-instrument everything via OTel]
→ [Collector samples/filters some]
→ [Store remaining data]
→ [Developer notices bug]
→ [Manually copy-paste error OR query via MCP server]
→ [AI gets incomplete/noisy context]
→ [AI suggests fix based on partial data]
→ [Human reviews]
→ [Code “looks plausible”]
→ [Deploy]
→ [Discover edge case not addressed by AI]
→ [More bugs in production]
There are several places this workflow breaks down. Ultimately, the root cause is a data and context problem. Until we fix it, no amount of AI capability is going to close the gap.
If I, as a developer, don’t have the right data from my observability tools out of the box, there’s no magic solution that will come from passing it to an AI agent . The problem is what “the right data” actually looks like and how far current tools are from providing it.
Current observability tools are pretty good at gathering:
System-level metrics : uptime, latency, error rates.
Service-to-service traces: which services called which, and how long each hop took.
However, they consistently leave gaps where it matters most for debugging complex issues:
Sampled traces and logs: aggressively thinned out to control storage costs.
Session-based telemetry: no correlation of user journeys with backend behavior.
Request/response payloads: missing for internal services and external dependencies.
Receive weekly engineering insights to level up your leadership approach.
As this information isn’t provided by default, AI tools end up working with a partial view of reality. They can still be genuinely useful, pattern-matching against common issues and getting it right for well-scoped, well-understood problems.
AI debugging falls apart the moment you hit a complex, distributed systems issue. The kind where you need full runtime context to understand what actually went wrong. In those cases, engineers still need to fire up the local dev environment, pull in all the tooling, and debug manually.
Current observability tools are fundamentally broken for AI debugging and the gap is about to get much worse.
According to Harness’s State of Software Delivery 2025 report, 67% of developers now spend more time debugging AI-generated code than they did before they started using AI tools. Stack Overflow’s 2025 Developer Survey tells the same story from a different angle: 66% of developers say they’re spending more time fixing “almost-right” AI-generated code, and 45% list it as their top frustration with AI tools overall.
Think about what happens as more teams adopt AI-coding tools or implement agentic AI to build their applications. The volume of code goes up. The speed of shipping goes up. Yet, the observability tooling underneath hasn’t changed. The result? More bugs, harder to diagnose, and a (unnecessary) proliferation of observability tools.
Here’s what fixing it actually looks like.
Companies’ approach to telemetry collection is preventative. They collect EVERYTHING because they don’t know what they’ll need when an issue happens. Instrumenting on demand (after an issue exposes a gap) is expensive and time-consuming, so the path of least resistance is to cast the widest net up front.
The problem is that casting a wide net and actually catching the right fish are two very different things. The result is companies hoarding massive volumes of generic telemetry and trying to wrangle the storage costs that come with it.
“ Observability is eating our infrastructure budget” is a complaint I hear regularly on customer calls and it’s not an anomaly – it’s the norm.
OpenTelemetry has made this worse, not better. Its democratized instrumentation is a genuine win for standardization, but it accidentally created a tragedy of the commons. When it’s trivially easy for every team to auto-instrument everything, everyone does. Nobody stops to ask whether the data is actually useful.
You end up with a system that’s drowning in telemetry but starving for signal. Michele Mancioppi, head of product at Dash0, put it bluntly during a recent LeadDev panel on agentic observability : “At KubeCon, I shared that, of the logs getting into observability tools, less than 1% are fatal errors. All the rest is probably a waste of money.”
This is where AI debugging hits its first wall. You can’t just point an AI agent at your observability data and expect it to work. There are two reasons why.
First, there’s cost. AI models charge per token processed , and feeding them a firehose of telemetry is expensive. Second, there’s a hard technical limit. AI models have a finite context window. You have to choose what you feed them. If you feed them raw OTel traces and logs, the agent either gets lost in the noise, or it misses the 1% that actually matters because that slice was sampled away or was never collected in the first place.
Missing critical debugging data
Even if your AI agent isn’t drowning in irrelevant data, it still doesn’t have the full picture. The context it needs is either scattered across disconnected tools or was never collected. For example:
Request/response payloads aren’t captured by default due to privacy and cost concerns. You can see that an Application Programming Interface (API) call was made, but not what data was sent or what came back.
Headers (authentication tokens, routing info, custom metadata) are often redacted for security reasons, removing critical debugging context.
External API exchanges are black boxes. Distributed tracing stops at your system boundary. When your backend calls Stripe, Twilio, or AWS, you see the call happened and how long it took, but not what you sent them or what they returned.
Session correlation is fragmented at best. Frontend monitoring (RUM) lives in one tool, backend observability in another. Connecting “this user clicked A → then B → then backend failed” requires manual stitching across platforms.
Current observability tools can show you that something happened and where it happened, but they struggle to automatically connect the dots across your entire stack. Getting full correlation requires extensive instrumentation and gets prohibitively expensive at scale. Especially when capturing the actual data (payloads, headers, full context) flowing through the system.
Let’s assume for a moment that all the data exists and nothing is missing. It’s just siloed across different tools.
In theory, an AI agent could query multiple observability platforms (frontend errors in Sentry, backend traces in Datadog, external API logs in Stripe, session replays in LogRocket) and correlate them automatically. In practice, this almost never happens (at least, not yet).
Frontier AI models haemorrhage sensitive data
The architecture gap your AI agent will expose
Each tool has different APIs, authentication methods, and data formats. More critically, the correlation keys that tie events together – request IDs, trace IDs, session IDs – rarely propagate across tool boundaries. The agent would need to guess which frontend error corresponds to which backend trace based on timestamps, an error-prone approach that breaks down under any real load or clock skew.
So even with perfect API access to every tool, the agent still can’t see the full picture unless the data was already correlated when it was collected.
This leaves developers doing the correlation work manually: copying request IDs between tools, matching timestamps across dashboards, piecing together the story, and only then pasting the relevant fragments into an AI tool for analysis.
Human code reviews can’t keep up with AI code generation
As every engineer knows, reading code is harder than writing it. You have to reverse-engineer someone else’s thought process, understand why they made certain decisions, and spot edge cases they might have missed. Even reviewing your own code weeks later isn’t much easier: our memory is fallible, and documentation is rarely as complete as it should be.
Now add an AI assistant that writes code faster than any human can. AI can generate 100 lines of plausible-looking code in 30 seconds. A thorough review of those same 100 lines could take 15-20 minutes. The math doesn’t work. Teams either bottleneck on review or let things slip through.
And things do slip through. AI-generated code often works for the happy path the model imagined, but it doesn’t think defensively the way experienced developers do. It pattern-matches based on what it’s seen before rather than reasoning about edge cases, error states, or unexpected inputs.
The result? Even if the bug rate per line of code stays constant, you’re shipping more code faster, which means more bugs in absolute terms. As AI-generated code compiles cleanly and looks structurally correct, bugs are harder to catch in review. The code looks fine. It just fails in production under conditions the AI never considered.
This is where the “AI makes you 10x more productive” narrative breaks down. “It compiles on the first try” doesn’t mean the code is correct. It means you’ve deferred the debugging work to later when the bug hits production.
The observability gap compounds this problem. Without complete runtime context (the actual payloads, the external API responses, the session state) you’re reverse-engineering logic you didn’t write, based on incomplete information, under time pressure.
The observability industry is starting to recognize these problems, but the solutions being proposed address different parts of the puzzle.
Approach 1: smart filtering at the collector layer
Some vendors are advocating for intelligent collection pipelines:
[App]
→ [Send everything]
→ [Smart AI collector]
→ [Decides what to keep/drop/escalate]
→ [Storage]
→ [AI gets clean data]
The idea is to use AI at the collector layer to filter out noise before it hits storage. Keep fatal errors, drop health checks, and intelligently sample based on what looks important.
This solves the “too much irrelevant data” problem. Storage costs go down because you’re not keeping everything. AI agents get cleaner data because the noise has been filtered out.
However, it doesn’t solve the “missing critical debugging data” problem, nor does it address ingestion costs. Also, if the collector drops a request because it looked unimportant, and that request turns out to be the one that caused a production incident, the data is gone.
Approach 2: AI agents embedded in observability platforms
Other vendors are building AI directly into their platforms: agents that live inside the tool and have native access to all the telemetry data collected there.
[App]
→ [Collect everything]
→ [Observability Platform Storage]
→ [Built-in AI Agent queries platform data]
→ [AI answers questions using available data]
This partially solves the data correlation problem: if a vendor collects session replays, metrics, logs, and traces in one place, their AI agent can correlate across all of it without needing to query multiple APIs.
However, we’re back to the “too much irrelevant data” problem, Missing data is still missing, not to mention that now we also have vendor lock-in to contend with.
Approach 3: session-base

[truncated]

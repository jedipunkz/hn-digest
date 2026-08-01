---
source: "https://www.aulinq.com/en/blog/build-distributed-saas-from-scratch-with-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=49137052"
title: "How to Build a Modern Distributed SaaS from Scratch with AI Agents"
article_title: "How to build a modern distributed SaaS from scratch, completely with AI agents — Aulinq Blog · Aulinq"
author: "amadmike"
captured_at: "2026-08-01T18:54:53Z"
capture_tool: "hn-digest"
hn_id: 49137052
score: 1
comments: 0
posted_at: "2026-08-01T18:31:40Z"
tags:
  - hacker-news
  - translated
---

# How to Build a Modern Distributed SaaS from Scratch with AI Agents

- HN: [49137052](https://news.ycombinator.com/item?id=49137052)
- Source: [www.aulinq.com](https://www.aulinq.com/en/blog/build-distributed-saas-from-scratch-with-ai-agents)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T18:31:40Z

## Translation

タイトル: AI エージェントを使用して最新の分散型 SaaS をゼロから構築する方法
記事のタイトル: AI エージェントを完全に使用して、最新の分散型 SaaS をゼロから構築する方法 — Aulinq ブログ · Aulinq
説明: 実際のプロセス - 技術スタックの調査、ライブラリのギャップの埋め、反復によるアーキテクチャ計画、エージェント作成のコードの漂流を防ぐ開発ループ。この方法で 1 つの本番プラットフォームを構築することに重点を置いています。

記事本文:
AI エージェントを完全に使用して、最新の分散型 SaaS をゼロから構築する方法 — Aulinq ブログ · Aulinq ← ブログに戻る AI エージェントを完全に使用して、最新の分散型 SaaS をゼロから構築する方法
2026 年に別の SaaS を開始するのは奇妙なことです。談話では、このカテゴリーは飽和しており、AI によってすべての堀が平らになり、構築する価値のある唯一のものは、優れたランディング ページを備えたフロンティア モデルのラッパーだけであると述べられています。私はそれらの見解を読みましたが、診断にはほぼ同意しますが、結論には完全に同意しません。ラッパーはまさに最初に死ぬものです。残っているのは、その下にある退屈でコピーが難しいインフラストラクチャ (マルチテナント、請求、コスト管理、秘密の処理、再配信における冪等性) だけであり、モデルがよりスマートになったからといってその作業が安くなったわけではありません。
これは、その中の1つを構築する物語です。オーリンクといいます。これはマルチテナント、多言語の AI エージェント プラットフォームです。Telegram、WhatsApp、または Web チャットで顧客に応答し、実際に処理できる最も安価なモデルを通じて会話をルーティングし、トークンごとにテナントに 1 セント単位で請求し、各テナントのデータと機密を他のすべてのテナントから隔離します。 3 つの製品トラックがその配管の上にあります。代理店がクライアントにエージェントを再販するためのホワイトラベル パス、エージェントを自社製品に組み込むスタートアップ向けの API、そしてインフラストラクチャの博士号を持たずに動作するボットだけを必要とする個人創業者向けのシングル テナントのオンランプです。運用システムは、複数の Go サービス、1 つの Postgres、1 つの Redis、および NATS JetStream バスで構成されます。エンジニアは私だけで、フルタイムの仕事をしています。
最後の部分は、人々が実際に知りたいことです。雇用しながら分散プラットフォームを単独で出荷するにはどうすればよいですか

はい？正直な答えは 2 つありますが、どちらも「コーディングが速い」わけではありません。 1 つ目は、私が単独でコーディングすることをやめ、AI エージェントのチーム (プランナー、実装者、レビュー担当者、検証者) のチームを編成し始めたことです。各チームは独自のコンテキストと独自のツールを備えています。 2つ目は、これを毎日1時間ずつやろうとするのをやめたことです。差分を確認したり、計画を承認したり、検証ツールの実行を観察したりするには、平日の夜 1 時間あれば十分です。スタックの決定を調査したり、アーキテクチャの草案を作成したり、障害が発生したサービスをドリフト状態から回復したりするだけでは十分ではありません。これらには、途切れることのない集中力が必要ですが、私にとってその集中力は土曜日であり、家族が外出する日曜日も時折あります。
したがって、方法と運用モデルは相互に依存します。週末バッチは、研究ループ、アーキテクチャの反復、ライブラリのギャップ埋め、エージェント ルールなどの思考作業が行われる場所です。平日の夜には、小規模な検証とマージのタスクが実行されます。それらを混同すると、火曜日の午後 11 時に研究半ばでスタックを選択することになります。これは、この方法が回避するように設計されたまさに障害モードです。
この記事はその方法を書き記したものです。洗練されたバージョンではなく、行き止まりとロールバックが残されたバージョンです。それが実際にそれを教える部分だからです。あなたがすでに有能なエンジニアであることを前提としています。エージェントはそれを置き換えません。彼らはあなたがもたらすどんな判断も増幅し、悪い判断も同様に急速に増幅します。
問題は「何が人気か」ではない
AI コーディング アシスタントが「分散型 SaaS にはどのスタックを使用すべきか」と尋ねたときに最初に行うことは、最も一般的な答えを与えることです。パイソン。高速API。ポストグレ。レディス。 「イベント駆動型」と言ったらKafkaかもしれません。これはチュートリアルの正しい答えですが、多くの場合、間違った答えになります。

このシステムは、他の AI エージェントが今後 2 年間延長する予定です。
人気のスタックは、ほとんどの人に読みやすいため人気があります。ほとんどの人間にとって読みやすいことと、ほとんどのエージェントにとって読みやすいことは同じではありません。エージェントはチュートリアルの大規模なエコシステムの恩恵を受けません。彼らはすでにチュートリアルを読んでいます。彼らは、間違いを拒否する型システム、大規模に失敗するランタイム、および間違いが発生する可動部分が最も少ないデプロイメント アーティファクトの恩恵を受けています。
私はチャットではなく、徹底的なリサーチを通じてスタックの質問を実行しました。 4 回または 5 回の反復: 最初に広範な調査を行い、次に生き残った 2 つまたは 3 つの選択肢を絞り込み、次にそれぞれを詳細に読み取り、次に 2 番目のモデルと照合して最初のモデルのバイアスを捕捉します。このパターンは、スタックだけでなく、このプロジェクトのすべての重要な決定に対して従ったものです。一度聞くよりも遅いです。一度問い合わせて 4 か月後に再構築するよりも速いです。
剛が勝ちました。 AI エージェント ツールの意味で人気があるからではなく (そうではなく、Python が優勢です)、Go コンパイラーがエージェントのミスのクラスを午前 3 時のスタック トレースではなくビルド エラーに変換するからです。エージェントがイベント構造体にフィールドを追加し、コンシューマーの更新を忘れると、ビルドは失敗します。 Python の同等の場合、コンシューマーは暗黙的に None にデシリアライズし、本番環境でそれがわかります。
このトレードオフは現実のものであり、これに名前を付けたいのですが、Python ML エコシステムへの直接アクセスを放棄しました。プラットフォーム内のすべてのモデル呼び出しは、インプロセス ライブラリではなく HTTP API を経由します。これは、LLM ルーティング プラットフォームにとっては問題ありません (いずれにせよ、モデルはプロバイダーの背後に存在します)。ローカル推論や大量の前処理を行うチームにとっては、実際のコストがかかります。 Python が適切な呼び出しだった場合も含め、完全な Go-vs-Python の議論を個別に書きます。

。
ライブラリが存在しない場合は書き込む
スタックの後、すべてのレイヤーに対して同じ調査ループが行われます。イベント バス: Kafka 上の NATS JetStream。サブジェクトベースのルーティングはエージェントが推論できる文字列であり、スキーマ レジストリを使用したトピックベースのルーティングはエージェントが間違えるビルド ステップであるためです。永続性: ドメインごとにスキーマを持つ 1 つの Postgres。エージェントの読みやすさの点で、「エージェントが完全に読み取ることができる 1 つのデータベース」が「それぞれ独自の移行ツールを備えた 10 つのデータベース」よりも優れているためです。
そして、後で説明したときに人々が驚いた部分。私が必要としたライブラリのいくつかは、私が使用したいと思っていた形式には存在しませんでした。すべてのイベントをエンベロープでラップし、生の Publish(subject, []byte) の代わりに PublishInbound を公開する、型指定された JetStream パブリッシャー。単純な連続障害カウンターではなく、スライディング障害ウィンドウを備えたサーキット ブレーカー。平文をツール層に漏洩しないエンベロープ暗号化を備えたマルチテナントの秘密保管庫。
他のエンジニアからの反応はたいてい「ライブラリ X を使えばいい」というものだった。それが正しいこともあった。場合によっては、既存のライブラリが薄いラッパーであり、エージェントがすべてのサービスで同じボイラープレートを記述する必要がありました。ギャップが実際にある場合、ライブラリの作成は最大の問題ではありません。エージェントと数週間週末を過ごすだけで、システムの残りの部分が期待するとおりのインターフェイスが完成します。このプラットフォームには、すべてのサービスがインポートする一連の内部 Go ライブラリ ( go-events 、 go-ai-providers 、 go-resilience 、 go-secrets 、その他 12 個) が含まれるようになりました。それぞれが存在するのは、公開された代替案では、エージェントが接着剤を減らすのではなく、より多くの接着剤を作成することになるためです。
ブレーカーはその形状の好例です。私が見つけたパブリック Go サーキット ブレーカー ライブラリは、連続した失敗をカウントし、タイマーでリセットしていました。欲しかったのは引き違い窓

— 最後の 60 秒間の失敗をカウントし、しきい値を超えた場合はオープン、調査のために半分オープン、1 回の成功でクローズします。これは 60 行の Go です。
func ( cb * SlidingWindowCircuitBreaker ) Call ( fn func () error ) error {
cb.muロック ()
スイッチ cb.state {
ケースcbOpen:
時間があれば。以来 (cb.openedAt) >= cb.cooldown {
cb.state = cbHalfOpen
} それ以外の場合は {
cb.muロックを解除する ()
エラー回路オープンを返します
}
}
cb.muロックを解除する ()
エラー:= fn ()
cb.muロック ()
cb.mu を延期します。ロックを解除する ()
エラーの場合 != nil {
cb.failures = append (cb.failures, time.Now ())
// ウィンドウよりも古い失敗をプルーニングします...
if cb.state == cbHalfOpen || len (cb.failures) >= cb.maxFailures {
cb.state = cbOpen
cb.openedAt = 時間。今（）
}
エラーを返す
}
if cb.state == cbHalfOpen {
cb.state = cbClosed
}
nilを返す
}
エージェントはその最初のバージョンを 1 回のパスで作成しました。私はそれを見直し、枝刈りループが 1 つずれていることを発見し、エージェントに修正してもらいました。重要なのは、このライブラリには、アダプター層がなく、残りのシステム呼び出しとまったく同じインターフェイスがあり、それをインポートするすべてのサービスが同じ障害セマンティクスを取得するということです。人気のあるライブラリをラップした場合、各サービスには異なるアドホック ラッパーがあり、それらのサービスを拡張するエージェントはそれぞれを学習する必要があります。イベント パブリッシャー、AI プロバイダー インターフェイス、シークレット ボールト、RAG セレクターなど、ライブラリ自体には一連のシリーズが登場しています。それぞれが、パブリック オプションがエージェントの接着力を低下させるのではなく、より多くすることを意味するケースです。
アーキテクチャ計画の反復
開発前に、エージェントと一緒にアーキテクチャ全体を計画しました。これは、ほとんどの人がスキップするステップであり、最も遅い部分のように感じられたとしても、後で最も時間を節約できるステップです。
私が学んだ重要なこと: エージェントには、自分が持っているだけの初期コンテキストを提供する必要があります。

デフォルトでは、できる限り何もしないことです。エージェントに「マルチテナント AI プラットフォームのアーキテクチャを設計してください」と依頼すると、「AI サービス」というラベルの付いたボックスを含む 3 層の図が表示されます。制約を考えて質問してみると、メッセージング プラットフォームはこれ、コスト層のルーティング要件はこれ、テナンシー ルールはこれ、期待するイベント フローはこれ、避けたいものはこれです - そうすれば、実際に構築できるものが得られます。
私は繰り返しました。 4、5、6ラウンド。私は計画を入手し、それを読み、特定の決定を差し戻し、修正を求めます。エージェントは選択を擁護し、私はそれが壊れた具体的なシナリオで反論し、エージェントは修正します。退屈で迷惑です。むしろコーディングをしていればよかったと思うこともあります。しかし、5 回の「このケースはどうだろう」を繰り返しても生き残る計画は、エージェントが途中でアーキテクチャを考案することなく実装できる計画です。
その繰り返しの中で私が最も大切にするようになったのは、まだ進んでいない方向性について考えることでした。現在の計画は私が現在構築しているものだけをカバーする必要がありますが、アーキテクチャは私が次に構築する可能性のあるものに向けて曲げる必要があります。一例: 当時はテナントが 1 つで厳密には分割が必要ではなかったにもかかわらず、エージェント構成を不変のテンプレートとテナントごとのオーバーライドとして設計しました。 6 か月後、2 番目のテナントが到着し、ホワイトラベルのケースが現れたとき、オーバーライド モデルはすでに対応していました。最初のパスでより単純な「エージェントごとに 1 つの構成」を構築していたら、後でシステムを拡張するエージェントは負荷がかかった状態でデータ モデルを移行する必要があり、これはエージェントが苦手とする種類のタスクです。アーキテクチャの決定自体 (Python について、モノリスについて 10 個のマイクロサービス、サービスごとにデータベースについて 1 つの Postgres) は、それぞれ後で独自に記述されます。この記事は防御ではなくプロセスです

e.
計画が安定すると、開発は具体的な小さなタスクに分割された 1 つの大きなトップレベルの計画になりました。 「メッセージング サービスを構築する」のではなく、「イベント ライブラリに PublishInbound メソッドを追加します。これが構造体、これがサブジェクト パターン、単体テストを作成します。」続いて次の小さな作業です。それでは次です。
順序が重要でした。まずライブラリ。小さな例でそれぞれが機能することを証明し、単体テストでコントラクトをロックします。次に、コア サービスが 2 つ存在すると相互にクロステストします。次に、最初の UI、これは別の研究ループを意味します。今回はフロントエンド フレームワークに関するものです。エージェントの最初の答えは、ボイラープレートの山を備えた React であり、私が望んでいたのは、静的にコンパイルされ、ビルド ステップの競合なしで i18n を処理できる最小のものだったからです。
ここに役割の分離が存在します。 1 人のエージェントがコードベースを読み取り、タスクの計画を作成します。 2 番目のエージェントが計画に基づいて実行します。 3 番目は diff を読み取り、バグを探します。 4 つ目はビルド、再起動、テストを実行しますが、システムが正常であることを確認するまで、完了した作業を呼び出すことはできません。エージェント自身が自分の仕事を宣伝することはありません。実装者はレビューしません、レビュー者は検証しません、検証者は承認しません。
ここでは再導出しません。次の記事は

[切り捨てられた]

## Original Extract

The actual process — tech-stack research, filling library gaps, architecture planning in iterations, and the development loop that keeps agent-written code from drifting. Grounded in building one production platform this way.

How to build a modern distributed SaaS from scratch, completely with AI agents — Aulinq Blog · Aulinq ← Back to blog How to build a modern distributed SaaS from scratch, completely with AI agents
Starting another SaaS in 2026 is a strange thing to do. The discourse says the category is saturated, that AI flattened every moat, that the only things left worth building are wrappers around a frontier model with a nice landing page. I read those takes and I mostly agree with the diagnosis and completely disagree with the conclusion. The wrappers are exactly what dies first. What's left is the boring, hard-to-copy infrastructure underneath — the multi-tenancy, the billing, the cost control, the secrets handling, the idempotency across redelivery — and that work has not gotten cheaper just because the model got smarter.
This is the story of building one of those. It's called Aulinq. It's a multi-tenant, multilingual AI agent platform: the thing that answers a customer on Telegram, WhatsApp, or web chat, routes the conversation through the cheapest model that can actually handle it, charges the tenant per token in fractions of a cent, and keeps every tenant's data and secrets walled off from every other tenant. Three product tracks sit on top of that plumbing — a white-label path for agencies reselling agents to their clients, an API for startups embedding agents in their own product, and a single-tenant onramp for a solo founder who just wants a working bot without the infrastructure PhD. The production system is several Go services, one Postgres, one Redis, and a NATS JetStream bus. I am the only engineer, and I have a full-time job.
That last part is the thing people actually want to know about. How do you ship a distributed platform solo while employed? The honest answer is two things, and neither of them is "I code fast." The first is that I stopped coding solo and started orchestrating a team of AI agents — a planner, an implementer, a reviewer, a verifier, each with its own context and its own tools. The second is that I stopped trying to do this in daily one-hour increments. A weeknight hour is enough to review a diff, approve a plan, or watch a verifier run. It is not enough to research a stack decision, draft an architecture, or pull a failing service out of a drifted state. Those need a block of uninterrupted attention, and for me that block is Saturday, and the occasional Sunday when the family is out.
So the method and the operating model depend on each other. The weekend batch is where the thinking work happens — the research loops, the architecture iterations, the library gap-filling, the agent rules. The weeknights are where the small verify-and-merge tasks happen. Mixing them up is how you end up with a half-researched stack choice made at 11 PM on a Tuesday, which is the exact failure mode this method is designed to avoid.
This article is that method, written down. Not the polished version — the version with the dead ends and the rollbacks left in, because those are the parts that actually teach it. It assumes you're already a competent engineer. The agents don't replace that. They amplify whatever judgment you bring, and they amplify bad judgment just as fast.
The stack question isn't "what's popular"
The first thing an AI coding assistant will do when you ask "what stack should I use for a distributed SaaS" is give you the most popular answer. Python. FastAPI. Postgres. Redis. Maybe Kafka if you say "event-driven." This is the correct answer for a tutorial and frequently the wrong answer for a system that other AI agents are going to extend for the next two years.
The popular stack is popular because it's legible to the most people. Legible to the most humans is not the same as legible to the most agents. Agents don't benefit from a large ecosystem of tutorials — they've already read them. They benefit from a type system that rejects their mistakes, a runtime that fails loudly, and a deployment artifact with the fewest moving parts to get wrong.
I ran the stack question through deep research, not a chat. Four or five iterations: first a broad survey, then narrowing on the two or three options that survived, then a deep read on each, then a cross-check against a second model to catch the bias of the first. The pattern I followed for every consequential decision in this project, not just the stack. It's slower than asking once. It's faster than asking once and rebuilding in month four.
Go won. Not because it's popular in the AI-agent-tooling sense — it isn't, Python dominates that — but because a Go compiler turns a class of agent mistakes into build errors instead of 3 AM stack traces. When the agent adds a field to an event struct and forgets to update a consumer, the build fails. In the Python equivalent, the consumer silently deserializes to None and you find out in production.
The trade-off is real and I want to name it: I gave up direct access to the Python ML ecosystem. Every model call in the platform goes through an HTTP API rather than an in-process library. That's fine for an LLM-routing platform — the models live behind providers anyway — and it would be a real cost for a team doing local inference or heavy preprocessing. I'll write the full Go-vs-Python argument separately, including the cases where Python would have been the right call.
When the library doesn't exist, write it
After the stack, the same research loop for every layer. Event bus: NATS JetStream over Kafka, because subject-based routing is a string an agent can reason about and topic-based routing with a schema registry is a build step an agent will get wrong. Persistence: one Postgres with schema-per-domain, because "one database an agent can read in full" beats "ten databases each with their own migration tool" for agent legibility.
Then the part that surprised people when I described it later. Several of the libraries I needed didn't exist in a form I wanted to use. A typed JetStream publisher that wraps every event in an envelope and exposes a PublishInbound instead of raw Publish(subject, []byte) . A circuit breaker with a sliding failure window rather than a naive consecutive-failure counter. A multi-tenant secrets vault with envelope encryption that doesn't leak plaintext to the tools layer.
The reaction from other engineers was usually "just use library X." Sometimes that was right. Sometimes the existing library was a thin wrapper that would have forced the agent to write the same boilerplate around it in every service. When the gap is real, writing the library is not the biggest problem — it's a few weekends with an agent, and you end up with exactly the interface the rest of the system expects. The platform now has a set of internal Go libraries ( go-events , go-ai-providers , go-resilience , go-secrets , and a dozen more) that every service imports. Each one exists because the public alternative would have made the agents write more glue, not less.
The breaker is a good example of the shape. The public Go circuit-breaker libraries I found counted consecutive failures and reset on a timer. What I wanted was a sliding window — count failures in the last sixty seconds, open if they cross a threshold, half-open to probe, close on a single success. It's sixty lines of Go:
func ( cb * SlidingWindowCircuitBreaker ) Call ( fn func () error ) error {
cb.mu. Lock ()
switch cb.state {
case cbOpen:
if time. Since (cb.openedAt) >= cb.cooldown {
cb.state = cbHalfOpen
} else {
cb.mu. Unlock ()
return ErrCircuitOpen
}
}
cb.mu. Unlock ()
err := fn ()
cb.mu. Lock ()
defer cb.mu. Unlock ()
if err != nil {
cb.failures = append (cb.failures, time. Now ())
// prune failures older than the window...
if cb.state == cbHalfOpen || len (cb.failures) >= cb.maxFailures {
cb.state = cbOpen
cb.openedAt = time. Now ()
}
return err
}
if cb.state == cbHalfOpen {
cb.state = cbClosed
}
return nil
}
An agent wrote the first version of that in one pass. I reviewed it, found the pruning loop was off by one, and had the agent fix it. The point is that this library now has exactly the interface the rest of the system calls, with no adapter layer, and every service that imports it gets the same failure semantics. If I'd wrapped a popular library, each service would have a different ad-hoc wrapper and the agents extending those services would have to learn each one. There's a whole series coming on the libraries themselves — the events publisher, the AI-provider interface, the secrets vault, the RAG selector — each one a case where the public option would have meant more agent glue, not less.
The architecture plan, in iterations
Before development, I planned the whole architecture with the agents. This is the step most people skip, and the step that saved the most time later — even though it felt like the slowest part.
The key thing I learned: you have to give the agent as much initial context as you have, because the agent's default is to do as little as possible. Ask an agent to "design the architecture for a multi-tenant AI platform" and you get a three-tier diagram with a box labeled "AI Service." Ask it with the constraints — here are the messaging platforms, here's the cost-tier routing requirement, here are the tenancy rules, here's the event flow I expect, here's what I want to avoid — and you get something you can actually build.
I iterated. Four, five, six rounds. I'd get a plan, read it, push back on a specific decision, ask for the revision. The agent would defend a choice, I'd counter with a concrete scenario where it broke, the agent would revise. Boring and annoying. I'd sometimes rather have been coding. But a plan that survives five rounds of "what about this case" is a plan an agent can implement without inventing the architecture as it goes.
The thing I came to value most in those iterations was thinking about directions I wasn't taking yet. The current plan only has to cover what I'm building now, but the architecture has to bend toward the things I might build next. One example: I designed the agent configuration as an immutable template plus a per-tenant override, even though at the time I had one tenant and didn't strictly need the split. Six months in, when the second tenant arrived and the white-label case showed up, the override model already handled it. Had I built the simpler "one config per agent" the first pass, the agent extending the system later would have had to migrate the data model under load, which is the kind of task agents do badly. The architecture decisions themselves — Go over Python, ten microservices over a monolith, one Postgres over a database per service — each get their own writeup later. This article is the process, not the defense.
Once the plan was stable, development was one large top-level plan broken into concrete, small tasks. Not "build the messaging service" — "add the PublishInbound method to the events library, here's the struct, here's the subject pattern, write the unit test." Then the next small task. Then the next.
The order mattered. Libraries first, with small examples proving each one worked and unit tests locking the contract. Then the core services, cross-testing against each other once two existed. Then the initial UI, which meant another research loop — this time on frontend frameworks, because the agent's first answer was React with a pile of boilerplate and what I wanted was the smallest thing that compiled to static and handled i18n without a build-step fight.
This is where the role separation lives. One agent reads the codebase and writes a plan for the task. A second agent implements against the plan. A third reads the diff and looks for bugs. A fourth builds, restarts, and runs the tests — and is not allowed to call the work done until it observes the system healthy. No agent promotes its own work. The implementer doesn't review, the reviewer doesn't verify, the verifier doesn't approve.
I won't re-derive it here — the next article is the

[truncated]

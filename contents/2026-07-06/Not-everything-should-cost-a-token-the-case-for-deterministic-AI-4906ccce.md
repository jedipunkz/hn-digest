---
source: "https://www.vybe.build/blog/learn-what-not-to-tokenize"
hn_url: "https://news.ycombinator.com/item?id=48811403"
title: "Not everything should cost a token: the case for deterministic AI"
article_title: "Not everything should cost a token: the case for deterministic AI | Vybe Blog"
author: "marwann"
captured_at: "2026-07-06T22:58:40Z"
capture_tool: "hn-digest"
hn_id: 48811403
score: 1
comments: 0
posted_at: "2026-07-06T22:36:40Z"
tags:
  - hacker-news
  - translated
---

# Not everything should cost a token: the case for deterministic AI

- HN: [48811403](https://news.ycombinator.com/item?id=48811403)
- Source: [www.vybe.build](https://www.vybe.build/blog/learn-what-not-to-tokenize)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T22:36:40Z

## Translation

タイトル: すべてにトークンが必要なわけではない: 決定論的 AI のケース
記事のタイトル: すべてにトークンが必要なわけではない: 決定論的 AI の場合 | Vybe ブログ
説明: AI モデルを通じてすべてのタスクを実行すると、トークンが消費され、コンテキストが肥大化します。何をトークン化すべきではないか、そしてなぜ最良のエージェント システムが決定的な作業をアプリ層に押し下げるのかを学びましょう。

記事本文:
すべてにトークンが必要なわけではありません: 決定論的 AI の場合 | Vybe ブログ AI の同僚
ベスト プラクティス すべてにトークンが必要なわけではありません: 決定論的 AI の場合
AI モデルを通じてすべてのタスクを実行すると、トークンが消費され、コンテキストが肥大化します。何をトークン化すべきではないか、そしてなぜ最良のエージェント システムが決定的な作業をアプリ層に押し下げるのかを学びましょう。
すべてにトークンが必要なわけではありません: 決定論的 AI の場合
すべてにトークンが必要なわけではない: 決定論的 AI のケース
私が最近話したチームは、毎朝メトリクス API を取得し、JSON を再形成し、結果をテーブルにドロップするという単純なことを実行するようにエージェントを設定しました。クリーンなアイデア。初日からうまくいきました。
その後、請求書が届きました。毎朝、その「単純な」ジョブは、生の JSON の数千トークンをコンテキスト ウィンドウにロードし、言語モデルに推論する必要のないデータを再フォーマットするように要求し、特権の料金を支払うことでした。 5 行のスクリプトが毎回同じように行う仕事を確率モデルが行っていたため、再フォーマットも時々間違っていました。コンテキスト ウィンドウには古いデータが詰め込まれていたため、エージェントが実行するはずだった実際の推論は悪化しました。
その罠は今、ほぼ全員にかかっています。チームは言語モデルをユニバーサル ランタイムとして扱います。タスクを文章で説明できる場合は、それを指示します。ただし、モデルは cron ジョブではなく、データベースでもありません。これはトークンでレンタルしている推論エンジンであり、請求額を爆上げして出力を低下させる最も早い方法は、そもそも頭脳を必要としなかった作業をそれに実行させることです。
学ぶ価値のあるスキルは、何をトークン化すべきではないかを知ることです。
「すべてがプロンプトである」が正しく感じられ、コストが非常にかかる理由
魅力は明らかです。プロンプトは速いです。スキーマ設計、エンドポイント、デプロイはスキップします。ユジュ

質問してください。 1 回限りの場合、それは本当に正しい判断です。
問題は、1 回限りのジョブが定期的なジョブになると発生します。モデルを通じてルーティングするすべての決定的タスクは、本来持つべきではなかった 3 つのプロパティを継承します。つまり、非決定的になる、遅くなる、従量制になるということです。スケジュールされた関数が無料で実行し、決して間違うことのない作業に対して、実行ごとに料金を支払うようになりました。
2 つの症状が続き、互いに複雑になります。
1 つ目は、間違ったものに合わせて調整されるコストです。トークンの支出は、求めている判断の価値を追跡する必要があります。代わりに、モデルに加えられた機械的作業の量を追跡します。 1 日に 1 つではなく 10 個の API をプルすると、たとえそのプルにインテリジェンスが必要なかったとしても、請求額は 10 倍になります。
2 つ目はコンテキストの肥大化です。データを「処理」するために、チームはデータをコンテキスト ウィンドウに読み込みます。しかし、コンテキスト ウィンドウは有限であり、そこで推論が行われます。生のレコードを詰め込んでしまうと、そのモデルが実際に得意なことは埋もれてしまいます。品質は低下し、ほとんどのプロバイダーは入力トークンを測定するため、同時に長いコンテキストのコストも上昇します。もっと悪い答えを得るために、より多くのお金を払うことになります。それが刺さる部分です。
それを解決する区別: 確率的作業と決定的作業
これが私がいつも立ち返るメンタルモデルです。実行場所を決定する前に、すべてのタスクを 2 つのバケットのいずれかに分類します。
ラインはシンプルです。タスクに判断が必要な場合、それはエージェントに属します。正確で再現性が必要な場合は、アプリに含める必要があります。決定論は機能であり、制限ではありません。毎朝同じように実行されるスケジュールされたジョブは、退屈で予測可能なものである必要があります。
そして、これがアイデア全体のアンカーラインです。トークンを使用してアプリを一度構築すると、実行するたびにトークンの支払いを停止します。で

機械の作成には知性が注ぎ込まれます。マシン自体は無料で動作します。
「何をトークン化してはいけないのかを学ぶ」ということは、一度だけ切り替える設定ではなく、すべてのタスクで実践する規律です。
誰もが誤解すること: データベースとしての記憶メモ
この間違いには、それ自体を指摘する価値のある特定のバージョンがあります。なぜなら、賢い人々が常にそれを行っているのを私は見ているからです。
人々はエージェントの記憶やメモをデータベースとして使用します。
彼らは、レコード、指標、パイプラインステージ、顧客リスト、在庫数などの構造化データをメモに落とし込み始めます。メモはすぐそこにあり、エージェントはすでにそれを読んでいるので、自然に感じられます。しかし、メモは構造化データにとって間違ったツールであり、その理由はすぐに積み重なっていきます。スキーマがないため、一貫性を強制するものはありません。クエリがないため、エージェントは 1 つのフィールドが必要になるたびにメモ全体をコンテキストに読み取ります。これは時間がかかり、従量制です。また、半構造化テキストの壁は、まさに実際のテーブルが置き換えるために存在するものであるため、メモは成長するにつれて劣化します。
ここで、「データベースを使えばいい」というアドバイスで埋もれてしまう部分がありますが、それが重要です。メモリのメモは問題ではありません。間違った仕事に使用するのは危険です。
メモは、クリーンなスキーマが存在せず、今後も存在しないコンテキストのカテゴリ全体にとって、真に適切なツールです。好み。口調も声も。エージェントが学習したパターン。決定とその背後にある理由。 「これが私たちのやり方です」という知識は、SELECT を実行するようなものではなく、エージェントの考え方を形作るものです。そのコンテキストは意図的に構造化されておらず、判断に情報を提供するために存在しており、それを厳格な表に強制的に組み込むと、それを有用にするニュアンスが取り除かれてしまいます。
構造化され、クエリ可能な大量のデータが実際のデータベースに保存されます。
構造化されていない、解釈的な、推論を形成するコンテキストがメモに残ります。
顧客を置く

メモに書き出してみると、痛みを感じるでしょう。ブランドの音声ガイドをデータベースに保存すると、散文としては問題ないものを過剰に設計したことになります。ツールをデータの形状に一致させると、両方の問題が解消されます。
Vybe モデル: AI とアプリ、それぞれが独自のレーンに
ここで、アーキテクチャは哲学ではなくなり、製品の決定となり始めます。
Vybe では、エージェントはただチャットするだけではありません。実際のアプリケーションを構築して運用します。この 1 つの事実により、エージェントには決定的な作業を行う場所があるため、単なる良いアドバイスではなく、「何をトークン化すべきではないかを学ぶ」ことが実際に実現可能になります。
模様はこんな感じです。定期的な作業はアプリ内でスケジュールされますが、モデルにそれを実行することを忘れないよう要求することでまとめられるわけではありません。データのプルと変換は、エージェントが必要なときに呼び出すアプリ エンドポイントの背後に存在するため、機械的な手順はプロンプトではなくコードとして実行されます。状態は、エージェントが決定的にクエリを実行する実際のデータベース内に存在します。エージェントは調整と推論を行います。アプリが実行され、記憶されます。各レイヤーは、実際に構築された役割を果たします。
2 つの例でそれを具体的に示しますが、そのうちの 1 つは私たち自身のものです。
競合他社のレーダー。私たちは社内に競合追跡エージェントを構築しましたが、そのコンテキストやメモに競合他社のデータを蓄積させることは簡単でした。そうではありません。スケジュールされた更新を備えたライブ データベース アプリを構築しました。スケジューリング、プル、保存はすべて、決定的な作業としてアプリ層で行われます。エージェントは、実際にフラグを立てる価値があるものを決定するという、頭を必要とする部分にのみトークンを費やします。そのシステムは出荷されて稼働しており、そのトークンコストは、エージェントが自ら配管を行う場合に比べて数分の一です。
ファルコン。私たちのギャラリーでは、ファルコンは企業の利益を監視する競争力分析アナリストです。

etitor の変更ログ、ブログ、ソーシャル活動を毎日チェックし、定期的にインテリジェンスの概要を電子メールで配信します。同じアーキテクチャで、今回は顧客対応です。監視頻度と電子メール配信は、アプリレベルで決定的に行われます。これらは毎日同じようにスケジュールに従って実行されます。 Falcon がトークンを費やす唯一のことは、この変化が実際にシグナルなのか、それともノイズなのかという判断です。アーキテクチャは期待どおりの動作をします。
両方の共通点に注目してください。高価でインテリジェントなリソースは、インテリジェンスが必要な仕事の一部に向けられています。それ以外はすべて、安価で信頼性の高い決定的なコードとして実行されます。
市販されているツールのほとんどはこれを行うことができません。その理由は洗練の問題ではなく構造的なものです。
純粋なチャット エージェント アシスタント、単一モデルのコパイロット、および Slack ボットには、オフロードするアプリ層がありません。決定論的な作業をプッシュする場所がないため、すべてがモデル内に留まります。これらのツールが持つ唯一のものがトークンであるため、スケジュールされたすべてのタスク、すべてのデータ プル、すべてのルックアップはトークンとして実行されます。彼らのアーキテクチャは、私たちが最初に考えた正確なコストの問題を強います。
純粋なノーコードビルダーとアプリビルダーには、逆のギャップがあります。決定論的なアプリを一日中実行できますが、何が重要かを決定する最上位の推論層はありません。
Vybe は、推論エージェントと決定論的アプリの両方が一流であり、同じ場所、同じシステムによって一緒に構築されたプラットフォームです。また、Vybe はモデルに依存しないため、モデル内で実行される機械的な作業が何であれ、1 つのベンダーのトークン レートに拘束されることはありません。一方の世界にもう一方の世界を模倣させるためにプレミアムを支払うのではなく、両方の世界の長所を活用できます。
トークン化してはいけないもののチェックリスト
プロンプトを表示する前に、次の質問に従ってタスクを実行してください。
これには裁判官が必要ですか

それとも単なる機械的なものなのでしょうか？判断は代理人に委ねられます。機械はアプリに行きます。
スケジュールに従って繰り返されますか?これをアプリレベルのスケジュールされたタスクに移動し、実行ごとの支払いを停止します。
エージェントは直接 API 呼び出しを行っていますか?これを、通話を行うアプリのエンドポイントに置き換え、エージェントにエンドポイントを呼び出してもらいます。
構造化データをメモに保存していますか?それをデータベースに移動し、必要なときにクエリを実行します。
コンテキスト内のデータは推論するためのものですか、それとも単に保存するためのものですか?それが収納スペースである場合は、窓から外し、実際に考えるためのスペースを確保してください。
そのリストを何度か実行すると、直感的に理解できるようになります。どのタスクがトークンを受け取るに値するのか、そしてどのタスクがあなたから静かにお金を奪っているのかを感じ始めます。
優れたエージェント システムの目標は、モデルを通じてすべてをルーティングすることではありません。それは、コストが発生する場所にインテリジェンスを正確に適用し、それ以外のすべてを決定論的なコードに静かにかつ低コストで処理させることです。
それがゲーム全体だ。実際にトークンに値するものを考え出し、マシンを一度構築して、それを自動的に実行させます。モデルは、配管作業ではなく、判断に時間を費やす必要があります。
Vybe はまさにこの目的のために構築されています。推論するエージェントと実行するアプリが連携するため、一方の仕事をさせるためにプレミアムを支払う必要がなくなります。 AI プラス アプリ モデルが実際にどのようなものかを確認するか、エージェント ギャラリーを参照してすでに実行されているパターンを確認します。
必要なものを説明してください。今すぐチームに発送してください。
複雑な設定は必要ありません。結果だけです。
AI の同僚と業務を実行するカスタム アプリ

## Original Extract

Running every task through an AI model burns tokens and bloats context. Learn what not to tokenize, and why the best agent systems push deterministic work down to the app layer.

Not everything should cost a token: the case for deterministic AI | Vybe Blog AI coworkers
Best Practices Not everything should cost a token: the case for deterministic AI
Running every task through an AI model burns tokens and bloats context. Learn what not to tokenize, and why the best agent systems push deterministic work down to the app layer.
Not everything should cost a token: the case for deterministic AI
Not Everything Should Cost a Token: The Case for Deterministic AI
A team I talked to recently wired up an agent to do something simple: pull a metrics API every morning, reshape the JSON, and drop the result into a table. Clean idea. It worked on day one.
Then the invoice came in. Every morning, that "simple" job was loading a few thousand tokens of raw JSON into a context window, asking a language model to reformat data it did not need to reason about, and paying for the privilege. The reformatting was occasionally wrong, too, because a probabilistic model was doing a job that a five-line script does the same way every time. The context window was so stuffed with stale data that the actual reasoning the agent was supposed to do got worse.
That trap catches almost everyone right now. Teams treat the language model as a universal runtime. If you can describe a task in a sentence, you prompt it. But a model is not a cron job and it is not a database. It is a reasoning engine you are renting by the token, and the fastest way to blow up your bill and degrade your output is to make it do work that never needed a brain in the first place.
The skill worth learning is knowing what not to tokenize.
Why "everything is a prompt" feels right and costs so much
The appeal is obvious. Prompting is fast. You skip the schema design, the endpoint, the deploy. You just ask. For a one-off, that is genuinely the right call.
The problem shows up when the one-off becomes a recurring job. Every deterministic task you route through a model inherits three properties it should never have had: it becomes non-deterministic, it becomes slow, and it becomes metered. You now pay, per run, for work that a scheduled function would do for free and never get wrong.
Two symptoms follow, and they compound each other.
The first is cost that scales with the wrong thing. Your token spend should track the value of the judgment you are asking for. Instead it tracks the volume of mechanical work you have shoved through the model. Pull ten APIs a day instead of one and your bill grows tenfold, even though none of that pulling required intelligence.
The second is context bloat. To "process" data, teams read it into the context window. But the context window is finite and it is where reasoning happens. Fill it with raw records and you crowd out the thing the model is actually good at. Quality drops, and because most providers meter input tokens, the cost of a long context climbs at the same time. You pay more to get worse answers. That is the part that stings.
The distinction that fixes it: probabilistic vs. deterministic work
Here is the mental model I keep coming back to. Sort every task into one of two buckets before you decide where it runs.
The line is simple. If a task needs judgment, it belongs with the agent. If it needs to be exact and repeatable, it belongs in an app. Determinism is a feature, not a limitation. A scheduled job that runs identically every morning is something you want to be boring and predictable.
And here is the anchor line for the whole idea: you build the app with tokens once, then you stop paying tokens every time it runs. The intelligence goes into creating the machine. The machine itself runs for free.
"Learn what not to tokenize" is a discipline you practice on every task, not a setting you toggle once.
The one everyone gets wrong: memory notes as a database
There is a specific version of this mistake that is worth calling out on its own, because I see smart people do it constantly.
People use their agent's memory or notes as a database.
They start dropping structured data into notes: records, metrics, pipeline stages, customer lists, inventory counts. It feels natural because the note is right there and the agent already reads it. But notes are the wrong tool for structured data, and the reasons pile up fast. There is no schema, so nothing enforces consistency. There is no query, so the agent reads the entire note into context every time it needs one field, which is both slow and metered. And notes degrade as they grow, because a wall of semi-structured text is exactly what a real table exists to replace.
Now here is the part that gets lost in the "just use a database" advice, and it matters: memory notes are not the problem. Using them for the wrong job is.
Notes are genuinely the right tool for a whole category of context that has no clean schema and never will. Preferences. Tone and voice. Patterns the agent has learned. Decisions and the reasoning behind them. The "here is how we do things" knowledge that shapes how an agent thinks rather than something you would ever run a SELECT against. That context is unstructured on purpose, it exists to inform judgment, and forcing it into a rigid table would strip out the nuance that made it useful.
Structured, queryable, high-volume data goes in a real database.
Unstructured, interpretive, reasoning-shaping context stays in notes.
Put a customer list in your notes and you will feel the pain. Put your brand voice guide in a database and you have over-engineered something that was fine as prose. Match the tool to the shape of the data and both problems disappear.
The Vybe model: AI and apps, each in its own lane
This is where the architecture stops being a philosophy and starts being a product decision.
On Vybe , agents do not just chat. They build and operate real applications. That single fact is what makes "learn what not to tokenize" actually achievable instead of just good advice, because the agent has somewhere to put the deterministic work.
The pattern looks like this. Recurring work gets scheduled inside the app, not held together by asking the model to remember to run it. Data pulls and transforms live behind an app endpoint that the agent calls when it needs them, so the mechanical steps run as code instead of as prompts. State lives in a real database that the agent queries deterministically. The agent orchestrates and reasons. The app executes and remembers. Each layer does the job it was actually built for.
Two examples make it concrete, and one of them is our own.
Competitor Radar. We built a competitive-tracking agent internally, and it would have been easy to let it hoard competitor data in its context or its notes. It does not. It built a live database app with scheduled refreshes. The scheduling, the pulling, and the storing all happen at the app layer as deterministic work. The agent only spends tokens on the part that needs a brain: deciding what is actually worth flagging. That system is shipped and running, and its token cost is a fraction of what it would be if the agent were doing the plumbing itself.
Falcon. In our gallery , Falcon is a Competitive Intelligence Analyst that monitors competitor changelogs, blogs, and social activity every day and delivers scheduled intelligence briefs by email. Same architecture, customer-facing this time. The monitoring cadence and the email delivery are deterministic app-level plumbing. They run on a schedule, identically, every day. The only thing Falcon spends tokens on is the judgment call: is this change actually a signal, or is it noise? The architecture doing exactly what it is supposed to.
Notice what both have in common. The expensive, intelligent resource is pointed at the one part of the job that needs intelligence. Everything else runs as cheap, reliable, deterministic code.
Most tools on the market cannot do this, and the reason is structural rather than a matter of polish.
Pure chat-agent assistants, the single-model copilots and Slack bots, have no app layer to offload to. There is nowhere to push the deterministic work, so it all stays inside the model. Every scheduled task, every data pull, every lookup runs as tokens because tokens are the only thing those tools have. Their architecture forces the exact cost problem we started with.
Pure no-code and app builders have the opposite gap. They can run deterministic apps all day, but there is no reasoning layer sitting on top deciding what matters.
Vybe is the platform where the reasoning agent and the deterministic app are both first-class and built together, in the same place, by the same system. And because Vybe is model-agnostic , you are not locked into one vendor's token rates for whatever mechanical work does still run through a model. You get the best of both worlds instead of paying a premium to force one world to imitate the other.
A checklist for what not to tokenize
Before you prompt, run the task through these questions.
Does this need judgment, or is it just mechanical? Judgment goes to the agent. Mechanical goes to an app.
Does it repeat on a schedule? Move it into an app-level scheduled task and stop paying per run.
Is the agent making an API call directly? Replace it with an app endpoint that makes the call, and have the agent invoke the endpoint.
Are you storing structured data in notes? Move it to a database and query it when you need it.
Is the data in your context there to be reasoned over, or just to be stored? If it is storage, get it out of the window and reserve that space for actual thinking.
Run through that list a few times and it becomes instinct. You start feeling which tasks deserve a token and which ones are quietly robbing you.
The goal of a good agent system is not to route everything through a model. It is to apply intelligence exactly where it earns its cost, and to let deterministic code handle everything else quietly and cheaply.
That is the whole game. Work out what actually deserves a token, build the machine once, and let it run on its own. Your model should be spending its time on judgment, not plumbing.
Vybe is built for precisely this: agents that reason and apps that execute, together, so you stop paying a premium to make one do the other's job. See what the AI plus apps model looks like in practice , or browse the agent gallery to see the pattern already running.
Describe what you need. Ship it to your team today.
No complex setup. Just results.
AI coworkers and custom apps that run your operations

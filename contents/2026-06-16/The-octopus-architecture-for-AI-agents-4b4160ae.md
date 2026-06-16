---
source: "https://blog.goodman.dev/blog/octopus-agent-architecture/"
hn_url: "https://news.ycombinator.com/item?id=48558696"
title: "The octopus architecture for AI agents"
article_title: "The octopus architecture for AI agents | Geoff Goodman"
author: "joshbetz"
captured_at: "2026-06-16T19:20:39Z"
capture_tool: "hn-digest"
hn_id: 48558696
score: 12
comments: 2
posted_at: "2026-06-16T17:24:45Z"
tags:
  - hacker-news
  - translated
---

# The octopus architecture for AI agents

- HN: [48558696](https://news.ycombinator.com/item?id=48558696)
- Source: [blog.goodman.dev](https://blog.goodman.dev/blog/octopus-agent-architecture/)
- Score: 12
- Comments: 2
- Posted: 2026-06-16T17:24:45Z

## Translation

タイトル: AI エージェントのタコ アーキテクチャ
記事のタイトル: AI エージェントのタコ アーキテクチャ |ジェフ・グッドマン
説明: タコのアーキテクチャは、半自律的なサブブレインにディスパッチする中央調整脳を備えたシステムを記述します。

記事本文:
AI エージェントのタコ アーキテクチャ |ジェフ・グッドマン 棒で物をつつく RSS エッセイについて書く
AI エージェントのタコ アーキテクチャ
タコのアーキテクチャは、半自律的なサブブレインにディスパッチする中央調整脳を備えたシステムを表します。
TorkBot はタコに似たデザインになっています。このアーキテクチャは、一連の行き詰まりと反復的な改善から生まれました。私がタコと言うとき、私が意味するのは、TorkBot には、多くの半自律的な付属器官を指揮する集中型の「脳」があり、それぞれが独自の脳を持ち、中央のディスパッチャに報告を返すということです。
静的レーンは長寿命の付属物です。キュレーターもその一人。プラグインは、Google Workspace レーンなど、他のプラグインに貢献できます。レーンテンプレートが異なります。テンプレートは、限定された目的のためにインスタンス化できる機能です。サンドボックス スナップショットはまた異なります。これはコラボレーターではなく、将来のサンドボックスを利用したレーンの保存されたファイルシステムの開始点にすぎません。
いくつかの競合するプレッシャーが私をこのアーキテクチャに押し込んだのです。
表面インタラクションへの応答性 — エージェントには、そのターンが多かれ少なかれ複雑さの範囲内で制限され、I/O を完全に回避できる設計が必要です。これにより、タスクや作業にかなりの時間がかかる場合でも、エージェントは迅速に対話できるようになります。
能力 — エージェントは、ターンを効率的に維持するためだけに達成できることを制限すべきではありません。委任を通じて複雑なタスクを追求し、それらのタスクをほぼリアルタイムで観察して操作できるメカニズムが必要です。
継続性 — エージェントは継続的な視点と個性を維持する必要があります。最良の継続性は、継続的にキュレーションされる単一の LLM 会話から得られます。このように、性格と短期記憶を「追加」する必要はありません。代わりに彼らは

これはアーキテクチャの副作用です。
これらのプレッシャーにより、上の図に見られるように、私は複数の「レーン」を備えた設計を余儀なくされました。 「フォアグラウンド」レーンは、ユーザーが表面アクティビティを通じて対話する LLM 会話です。しかし、ここで私は物議を醸す可能性が高い賭けをしました。すべてのサーフェスにわたるすべてのアクティビティが同じフォアグラウンド会話を経由するということです。スレッド、チャネル、さらにはプラットフォームさえもすべて折りたたまれています。現時点では、その認知の複雑さはおそらくほとんどのモデルの能力を超えており、おそらく限界を超えています。しかし、それが長く続くわけではないと確信しています。
すべてのサーフェスにわたるすべてのアクティビティは、同じフォアグラウンド会話を経由します。
これは、イベントごとに 1 つのモデル ターンを意味するものではありません。サーフェス メッセージ、システム リマインダー、およびレーン メッセージは保留中の入力として蓄積されます。これらは、ターゲット レーンがユーザー メッセージを受け入れることができるとき (アイドル状態、またはツール バッチがフラッシュされた後) に挿入されます。
これが、インタラクティブ性をアクティビティ量から切り離すものです。 10 のことが起こり、それでも正しい境界で 1 つの一貫したターンになる可能性があります。問題は、フォアグラウンド モデルが最新性、優先度、中断を理解する必要があることです。
TorkBot に関する私のテーマの一部は、創発的な行動と創発的なインテリジェンスに賭けることです。プラットフォームで定義された任意の境界を越えて LLM 会話を分割するシステムを考案することは、継続性の目標とは相反するものです。エージェントがスレッド間、さらにはサーフェス間でリンクを作成できるようにしたいと考えています。エージェントが Slack で開始し、GitHub で継続した作業を簡単に継続できるようにしたいと考えています。モデル インテリジェンスに関してはまだそのレベルに達していないとしても、すぐにそのレベルに到達し、その世界向けに設計されたエージェント システムが直観性とパワーの点で競合製品よりも優れたものになると私は確信しています。
タコのアイデアがここで実際に機能します。の形です

ハーネスの問題。
これは影響力を得るためにサブエージェントの流行に乗っているわけではありません。これは誕生し、その存在を獲得したデザインです。結局のところ、コンテキスト管理に戻ります。各付属物は独自のコンテキストを取得します。
フォアグラウンドは他のレーンに「話しかける」ことで作業を引き渡します。レーン間のコミュニケーションは単なるテキストであり、トレーニングの前後で意図の伝達者として散文に大きく偏るという考えに賭けています。フォアグラウンドはレーン テンプレートを選択し、それがサンドボックス レーンの場合は VM スナップショットを選択し、必要なものについての最初のメッセージをそのレーンに渡します。すでに生成されているレーンの場合は、単純なメッセージが送信されます。
レーンは、大量のツール呼び出し、行き止まりへの到達、I/O およびその他のより複雑なサンドボックス対応ワークフローの実行などの面倒な作業を担当できます。その混乱はレーンのコンテキスト内に閉じ込められたままになります。レーンは 2 つのメカニズムを介して相互に通信します。
レーンの ./shared フォルダーを介した仮想ファイルシステム アーティファクトへの参照。
just-bash のおかげで、すべてのレーンに最小限のサンドボックスのような環境が得られます。 @torkbot/sandbox のおかげで、サンドボックス レーンは完全なマイクロ VM を取得できる唯一のレーンです。メモリ内であっても、完全な Linux VM を使用していても、すべてのレーンは共通のディレクトリ構造と仮想ファイルシステムを共有するため、参照を渡して共通の理解を得ることができます。
フォアグラウンドでの会話は、すべての中間アーティファクトが消滅する場所になることなく、サーフェス全体で継続的に維持できます。これは、個性とスレッド間の直観にとって私が望んでいることです。付属器は、タスクのためのローカル作業メモリを運ぶことができます。前景には、関係、現在の意図、および統合を伝えることができます。
これにより、圧縮も非常に明白になります。各レーンは、特定のしきい値で非同期的に継続的に圧縮されます。

そして、何らかの奇妙な出来事によって、別のより高いしきい値を超えた場合には、同期的に実行されます。
この設計から得られる利点
平均インタラクション時間が重要です。
完了には時間がかかる場合があります。レーンがドキュメントを読んだり、I/O で待機したり、テストを実行したり、壁にぶつかったり、再試行したりすることは問題ありません。 1 つの付属器官が忙しいために前景が暗くなるのは問題ありません。
したがって、前景のレーンは小さくて退屈なものでなければなりません。安定したプロンプト、現在の意図、最近の表面活動、コンパクトな要約、証拠への参照などです。付属器の撹拌を維持します。これはコンテキスト効率の話でもあり、キャッシュ効率の話でもあります。安定したフォアグラウンド プロンプトは、LLM API キャッシュ ヒットの改善を意味します。ジャンクが少ないということは、最初のトークンが速くなり、認識上の抵抗が少なくなることを意味します。
キュレーションはそれを可能にします。圧縮により、レーンのコンテキストが永久に膨張するのを防ぎます。キュレーターは、耐久性のあるビットを記憶やスキルに組み込むことができます。アーティファクトはアーティファクトのままであり得る。トランスクリプトは、フォアグラウンドに戻されることなく、検査可能な状態を維持できます。
腕は忙しいかもしれません。ヘッドは利用可能な状態にしておく必要があります。
長い形式の技術メモ、実験、永続的なアイデア用に構築されています。

## Original Extract

The octopus architecture describes a system with a central coordinating brain that dispatches to semi-autonomous sub-brains.

The octopus architecture for AI agents | Geoff Goodman Poking things with sticks Writing About RSS Essay
The octopus architecture for AI agents
The octopus architecture describes a system with a central coordinating brain that dispatches to semi-autonomous sub-brains.
TorkBot is designed a bit like an octopus. This architecture was born from a series of dead-ends and iterative improvement. When I say octopus, what I mean is that TorkBot has a centralized “brain” directing many semi-autonomous appendages, each with their own brains, reporting back to the central dispatcher.
Static lanes are the long-lived appendages. Curator is one. Plugins can contribute others, like the Google Workspace lane. Lane templates are different. A template is a capability that can be instantiated for a bounded purpose. A sandbox snapshot is different again: it is not a collaborator at all, just a saved filesystem starting point for a future sandbox-backed lane.
Several competing pressures are at play that pushed me into this architecture.
Responsiveness to surface interactions — The agent requires a design in which its turns are more or less bounded in complexity and can avoid I/O entirely. This allows the agent to interact quickly even when tasks or work may take quite some time.
Capability — The agent shouldn’t be limited in what it can accomplish just to keep turns efficient. It needs mechanisms to pursue complex tasks through delegation and be able to observe and steer those tasks close to real-time.
Continuity — The agent should maintain a continuous perspective and personality. The best continuity comes from a single LLM conversation that is continually curated. In this way, the personality and short-term memory don’t need to be “added in”; instead they’re a side effect of the architecture.
These pressures pushed me into a design with multiple “lanes”, as you can see in the diagram above. The “foreground” lane is the LLM conversation users interact with through surface activity. But here, I have made a bet that is likely controversial: all activity across all surfaces goes through the same foreground conversation. Threads, channels, and even platforms are all collapsed. Right now, that cognitive complexity is perhaps beyond the ability of most models and perhaps even beyond the frontier. But I’m certain that will not be the case for long.
All activity across all surfaces goes through the same foreground conversation.
That does not mean one model turn per event. Surface messages, system reminders and lane messages accumulate as pending input. They are injected when the target lane can accept a user message: idle, or after a tool batch has flushed.
This is what decouples interactivity from activity volume. Ten things can happen and still become one coherent turn at the right boundary. The catch is that the foreground model has to understand recency, priority and interruption.
Part of my thesis with TorkBot is to bet on emergent behaviour and emergent intelligence. Coming up with systems that split LLM conversations across arbitrary platform-defined boundaries is antithetical to the continuity goal. I want my agent to make links across threads and even across surfaces. I want the agent to be able to trivially continue work started in Slack and continued on GitHub. If we’re not there yet in model intelligence, I bet we will soon be and the agentic system designed for that world will stand above the competition in terms of intuitiveness and power.
The octopus idea is doing actual work here. It is the shape of the harness problem.
This is not jumping on the sub-agent bandwagon for the sake of clout. This is a design that emerged and earned its existence. After all, it comes back to context management. Each appendage gets its own context.
The foreground hands off work to other lanes by ‘talking’ to them. Inter-lane communication is just text, betting on the idea that pre- and post-training skew heavily towards prose as the carrier of intent. The foreground picks a lane template — and if it is a sandbox lane, a VM snapshot — and passes an initial message to that lane about what it wants. For lanes that are already spawned, a simple message is sent.
Lanes can own the messy work of doing a bunch of tool calls, hitting dead-ends, doing I/O and any number of more complex sandbox-enabled workflows. That mess stays contained in the lane’s context. Lanes communicate between each other via two mechanisms:
References to virtual filesystem artifacts via the lane’s ./shared folder.
All lanes get a minimal sandbox-like environment thanks to just-bash . Sandbox lanes are the only lanes that get a full micro VM, thanks to @torkbot/sandbox . Whether in-memory or using a full Linux VM, all lanes share a common directory structure and virtual filesystem which allows them to pass references and share a common understanding.
The foreground conversation can stay continuous across surfaces, which is what I want for personality and cross-thread intuition, without becoming the place where every intermediate artifact goes to die. The appendage can carry the local working memory for the task. The foreground can carry the relationship, the current intent, and the synthesis.
This also makes compaction pretty obvious. Each lane is continuously compacted asynchronously at a certain threshold and synchronously if, through some strange turn of events, it exceeds another higher threshold.
Benefits that fall out of this design
Mean-time-to-interaction is the prize.
Completion can take a while. It is okay for a lane to read docs, wait on I/O, run tests, hit a wall and try again. It is not okay for the foreground to go dark because one appendage is busy.
So the foreground lane has to stay small and boring: stable prompt, current intent, recent surface activity, compact summaries and references to evidence. Keep the churn in the appendages. That is both a context-efficiency story and a cache-efficiency story. A stable foreground prompt means better LLM API cache hits. Less junk means faster first tokens and less cognitive drag.
Curation makes that possible. Compaction keeps lane context from swelling forever. Curator can promote durable bits into memory or skills. Artifacts can stay artifacts. Transcripts can remain inspectable without being stuffed back into the foreground.
The arms can be busy. The head needs to stay available.
Built for long-form technical notes, experiments, and durable ideas.

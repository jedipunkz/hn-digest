---
source: "https://tenureai.dev/writing/ai-governance-cannot-be-a-tool-call/"
hn_url: "https://news.ycombinator.com/item?id=48593651"
title: "AI Governance Cannot Be a Tool Call"
article_title: "AI Governance Cannot Be a Tool Call | Tenure"
author: "jflynt76"
captured_at: "2026-06-19T01:06:55Z"
capture_tool: "hn-digest"
hn_id: 48593651
score: 3
comments: 0
posted_at: "2026-06-19T00:53:50Z"
tags:
  - hacker-news
  - translated
---

# AI Governance Cannot Be a Tool Call

- HN: [48593651](https://news.ycombinator.com/item?id=48593651)
- Source: [tenureai.dev](https://tenureai.dev/writing/ai-governance-cannot-be-a-tool-call/)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T00:53:50Z

## Translation

タイトル: AI ガバナンスはツール呼び出しではありえない
記事のタイトル: AI ガバナンスはツールの呼び出しではありえない |在職期間
説明: Enterprise AI には、すべてのモデル リクエストの下にコンテキスト コントロール プレーンが必要です。エージェントが呼び出すことを選択できるツールではありませんが、すべてのリクエストが通過する必要があるレイヤーです。

記事本文:
在職期間
プラットフォーム ▼ コア機能
ガバナンスはツールを呼び出すことはできません
Enterprise AI には、すべてのモデル リクエストの下にコンテキスト コントロール プレーンが必要です。
エージェントが選択して呼び出すことができるツールではありません。すべてのリクエストが通過する必要がある層。
現在、エージェント上に構築されたほぼすべてのものは、ガバナンス、メモリ、ポリシーをモデルが選択して呼び出すことができるツールとして扱います。
モデルが何を決定したかに関係なく、何かが真実でなければならない瞬間に、その形状は壊れます。
修正はより良いツールではありません。これはスタック内の異なる位置にあり、モデルの横ではなく、すべてのリクエストの下にあります。
Tenure はこれを AI コンテキスト コントロール プレーンと呼び、エンタープライズ AI ワークフローの管理されたメモリと状態から始まります。
それを呼び出すかどうかを決定するモデルに依存する場合、それはガバナンスではありません
AI コーディング エージェントのほぼすべてのアーキテクチャ図を開くと、同じことがわかります。
形。モデルは中央に座っています。その周りには、手を伸ばせるツールのリングがあります。
システム、メモリ ストア、コンプライアンス チェック、監査ログ。図は完成したように見えます。それは
どのツールも実際に機能するかどうかを決定する詳細が欠落しています。
モデルがそれを呼び出すことを決定した場合にのみ、何かを実行します。
カレンダーの検索にはこれで十分です。人々が実際に意図していることは問題ありません
彼らはガバナンスと言います。エージェントがアクセスを忘れずにチェックすることに依存するアクセス制御
アクセス制御ではありません。モデルが書き込みを選択した場合にのみ存在する監査証跡
監査証跡ではありません。モデルがたまたまリコールしたときに従うコーディング規約
ルールは制約ではありません。それは善意による提案です。
これは、現在の多くの AI インフラストラクチャの根底にあるカテゴリ エラーです。ツール呼び出しは
オプションのアクションに便利です。ガバナンスは任意ではありません。両方を同じ種類として扱う

f
モデル呼び出しツールにより、最も重要なコントロールを残したままシステムを拡張可能に見せます
モデルの協力次第です。
モデルが呼び出すために選択できるツールは、オプションの入力です。重要なガバナンス
オプションにすることはできません。現在、この 2 つは同じように構築されており、それが問題です。
エンタープライズ AI にはコンテキスト コントロール プレーンが必要
欠けている層は AI コンテキスト コントロール プレーンです。これはどのコンテキストを決定するインフラストラクチャです。
AI リクエストの送信が許可されるか、使用できるメモリ、適用されるポリシー、および取得されるもの
リクエストがモデルに到達する前に記録されます。
これはエージェント フレームワークとは異なります。これはメモリ SDK とは異なります。それ
は MCP サーバーではありません。これらは、エージェントが呼び出すことができる機能としてモデルの横にあります。コンテキスト
コントロール プレーンはリクエストのパスに存在します。を管理、強化、範囲指定し、監査することができます。
モデルがそれを認識する前にリクエストを送信します。
メモリはこの層の最初の避けられないプリミティブであるため、所有権はメモリから始まります。
すべての本格的な AI ワークフローには、最終的には永続的な状態が必要です。つまり、何が決定され、何が変更され、何が変更されたのか
置き換えられたか、誰が承認したか、どの範囲に適用されるか、なぜ挿入されたのか。それなら
状態は 1 つのアプリ内に存在しますが、他のアプリではすべて失われます。ツール呼び出しの背後にある場合、モデルは
スキップできます。リクエスト パスの下に存在する場合、それはインフラストラクチャになります。
Tenure は、ガバナンスされたものから始まるエンタープライズ チーム向けの AI コンテキスト コントロール プレーンです。
すべてのモデルリクエストの下にあるメモリと状態。
人々はすでにさまざまな方向からこの場所を旋回しています
コーディング エージェントは、差分を生成し、テストを実行し、タスクが完了したことを報告できます。より難しい
多くのエンジニアが今声を大にして尋ねている疑問は、その後何が起こるかということです。誰か
何が変わったのか、なぜ変わったのか、正しいファイルかどうかを知る必要がある

感動した、何だったのか
テストされた内容、何がスキップされたか、結果が安全に信頼できるかどうか。それはコード生成ではありません
問題。それはレビューと信頼の問題です。
別のグループがオーケストレーション側から同じ壁に近づき、システムを構築しています
エージェント自身の「完了」シグナルを権限としてカウントすることを明示的に拒否します。エージェント
という主張が生まれます。他のこと、リポジトリ、テスト、および実際の状態を観察すること
審査によって、その主張が正しいかどうかが判断されます。それは現実的で合理的なデザインです。それも
それ自体では解決しないことを静かに認めます。何が良いのかはまだ誰かが定義する必要があります
特定のリポジトリを意味し、その定義がエージェントの内容と同期しなくなるのを防ぎます。
実際に言われました。
3 番目のグループは、同じ軸に沿って記憶そのものを再考しており、エージェントの記憶のほとんどは、
今日は事後にクエリするものであり、より価値のあるバージョンは何かである
エージェントは、行動する前に調整し、行動の一環として修正する必要があります。語彙が違うと、
同じ根本的な苦情: エージェントを正直に保つべきことは、現時点では
オプション、後期、またはその両方として位置付けられます。
3 つの異なるエントリ ポイント。一つの構造的原因。誰も同じ角度からそれに到達しませんでした、
これは通常、問題が枠組みの練習ではなく実際にあることを示しています。
これはアプリケーションの下ですでに発生しています
コンテナはパケットを送信する前に許可を求めません。チェックするためにツールを呼び出すわけではありません
別のポッドと通信できるかどうか。ネットワーク ポリシーはコンテナのサービスではありません
覚えていれば呼び出すことができます。ワークロードの下、ワークロードのレイヤーで適用されます。
オプトアウトすることはできません。ワークロードがポリシーと連携していません。政策にはそれが必要ない
に。
インフラストラクチャでは多くの制御が始まります

になる前に、アプリケーションレベルの規約として
強制レイヤー。 「アプリは忘れずにチェックするべきだ」というのは問題ではないため、この変化は重要です。
セキュリティプロパティ。それは希望です。
同じ動きがエージェントでも利用できるようになりました。それを達成できる立場にある人はほとんどいない
この空間のほとんどすべては、レイヤーではなく、エージェントが呼び出すツールとして構築されています。
リクエストは回避できません。
ツール間でコンテキストが断片化している
平均的なエンタープライズ AI 環境は、1 つのモデル、1 つのアシスタント、または 1 つのエージェントではありません。それ
コーディング ツール、チャット クライアント、内部エージェント、CI ワークフロー、ノートブック、チケット発行が混在します。
システム、サポート ワークフロー、モデル プロバイダー。それぞれが独自のコンテキストを蓄積します。
それぞれに独自のメモリ モデルがあります。それぞれが何について異なる仮定を立てるでしょう
それは知っています、そしてなぜそれを知っているのか。
AI の使用が個人的かつ実験的なものであれば、これは管理可能です。ガバナンスの問題になる
AIが実際の仕事に参加する瞬間。 1 つのツールで行われたチームの意思決定が消えてはいけません
開発者がクライアントを切り替えるとき。ポリシーの修正は手動に依存すべきではありません
すべてのエージェントにコピーされます。古くなった建築上の事実が再浮上し続けるべきではありません。
アプリケーションはそれを記憶していましたが、別のアプリケーションは記憶していませんでした。
アプリレベルのメモリによりアイランドが作成されます。ツール呼び出しメモリにより、オプションのリコールが作成されます。エンタープライズ AI のニーズ
どちらでもない。ツールの下に配置できる、コンテキストの共有され管理された基盤が必要です。
モデルの上にあります。
エンタープライズ AI の未来は、より優れたメモリを備えた 1 つのアプリではありません。多くのアプリが 1 つを共有
管理されたコンテキスト層。
下の層にはまだ状態が必要です
記憶のない政策は脆い。メモリなしの監査は不完全です。記憶なしの評価
当時システムが何を信じていたのかについての永続的な見解はありません。 AI リクエストの下にあるレイヤー
n

実際の組織は、意思決定、例外、
設定、修正、ルール。
この状態は、意味的に類似したメモの山のように振る舞うことはできません。記憶するシステム
「接続プールの上限は 20 です」と後で「接続プールの上限は 50 です」と覚えたのですが、
両方の事実が永遠に戻り続けることができれば、記憶は解決されます。永続的な読書時間を生み出しました
クリーンアップの問題。将来読者になる人は皆、この矛盾に再度気づく必要があります。
管理されたメモリは、基板に作用するように移動します。事実には範囲が必要です。修正が必要です
由来。置き換えられた信念には、置き換えられたものとしてマークを付ける必要があります。現在の値は次のようにする必要があります
すべてのモデル リクエストで履歴を最初から再検出する必要がなくても利用できます。
これが、Tenure がメモリを検索ではなく状態として扱う理由です。重要なのは、それ以上取得しないことです
近くのテキスト。重要なのは、システムが現在何を信じているのか、なぜそれを信じているのか、どこで信じているのかを知ることです。
その信念が適用されるかどうか、そしてそれが次のリクエストに影響を与えることがまだ許可されているかどうか。
モデルの横ではなく、リクエスト パスの下にあります
Tenure は、AI クライアントとモデル プロバイダーの間に位置します。その位置が重要です。文脈という意味です
インジェクション、スコープ、監査証跡、メモリ ガバナンスは、選択したモデルには依存しません。
ツールを呼び出します。これらはリクエストがモデルに到達する前に発生します。
実際には、これによりチームは組織に優れた機能を提供しながら、すでに使用している AI ツールを使用できるようになります。
管理されたコンテキストの共有レイヤー。開発者は、VS Code、Cursor、Cline、Continue、
WebUI または別のクライアントを開きますが、メモリ層とポリシー層はその下で一貫したままです。
クライアントは変わることができます。モデルは変更される可能性があります。統治国家は消滅する必要はない
どちらかで。
これが、メモリ機能とコンテキスト コントロール プレーンの違いです。機能が向上します
1つのアプリケーション

ション。コントロール プレーンは、組織にどの AI を許可するかを決定するための 1 つの場所を提供します
知り、覚え、使い、証明すること。
MCP では、モデルがツールを呼び出すことができます。 Tenure は、モデルがリクエストを認識する前にリクエストを管理します。
エージェントの協力が必要ですか?
このカテゴリの最も簡単なテストは、製品がガバナンス、メモリ、ポリシー、
または監査。テストは、システムが機能するためにエージェントの協力を必要とするかどうかです。
モデルがメモリ ツールを呼び出すことを覚えておく必要がある場合、メモリはオプションです。モデルがそうする必要がある場合
忘れずにポリシー チェッカーを呼び出してください。ポリシーはオプションです。モデルが書くことを覚えておく必要がある場合
監査イベントの場合、監査証跡はオプションです。それらのツールがどれほど優れていても、依然として
モデルの裁量の下流。
コンテキスト コントロール プレーンには別のコントラクトがあります。リクエストは管理されたインフラストラクチャに入ります
模型に届く前に。アイデンティティ、範囲、記憶、ポリシー、来歴、監査を行うことができます。
モデルの決定ループの外側に接続されます。モデルは施行を知る必要はありません
起こった。スキップすることもできません。
エージェントの協力が必要な場合はツールになります。リクエストを回避できない場合は、
インフラストラクチャ。
モデルは誰もが目にする部分です。コントロール プレーンは、信頼しても安全かどうかを判断する部分です。
これは、ツール呼び出しがあらゆるものに対して間違った抽象化であることを意味するものではありません。ファイルを取得すると、
データベースのクエリ、チケットのオープン、ビルドの実行はすべて個別のアクションであり、
モデルは必要に応じて呼び出されます。
議論はより狭く、より重要です。何が何でも真実でなければなりません。
モデルは、オプションのツールと同じレイヤーに存在できないと判断します。アクセス制御。監査
トレイル。メモリスコープ。スーパーセッション。生成されたコードは誰よりも早くルールに準拠する必要があります
それを信頼します。これらは依存できません

モデルは尋ねることを決めました。
差分の後に何が起こるかを尋ねる人々、エージェントの「完了」を拒否する人々
つまり、人々は読み取り時ではなく書き込み時の補正を中心にメモリを再構築しています
ルックアップはすべて、異なる部屋からの同じ欠落レイヤーを記述しています。
カテゴリが表示されるようになりました。テニュアはボトムアップでそれを構築します: 統治されます
まずメモリと状態、次にコンテキスト、ポリシー、来歴、監査などのより広範なコントロール プレーンです。
エンタープライズ AI リクエスト全体にわたる評価。
MCP がメモリの間違った抽象化である理由
記憶は、エージェントが呼び出すために覚えておくべきものではありません。リクエストがモデルに到達した時点ですでに存在しているはずです。
アクセス制御、監査証跡、およびメモリ スコープは、エージェントの裁量ではなく、プロキシで強制されます。
89 のケースと 5 つの比較システムに対して測定された、管理された記憶の背後にある合否規律。

## Original Extract

Enterprise AI needs a context control plane beneath every model request. Not a tool the agent can choose to call, but a layer every request has to pass through.

tenure
Platform ▼ Core Capabilities
Governance cannot be a tool call
Enterprise AI needs a context control plane beneath every model request.
Not a tool the agent can choose to call. A layer every request has to pass through.
Almost everything built on agents today treats governance, memory, and policy as tools the model can choose to invoke.
That shape breaks the moment something has to be true regardless of what the model decides.
The fix is not a better tool. It is a different position in the stack: beneath every request, not beside the model.
Tenure calls this the AI context control plane, starting with governed memory and state for enterprise AI workflows.
If it depends on the model deciding to call it, it is not governance
Open almost any architecture diagram for an AI coding agent today and you will find the same
shape. The model sits at the center. Around it, a ring of tools it can reach for: a ticketing
system, a memory store, a compliance check, an audit log. The diagram looks complete. It is
missing the detail that determines whether any of it actually works: every one of those tools
only does something if the model decides to call it.
That is fine for a calendar lookup. It is not fine for the things people actually mean when
they say governance. Access control that depends on the agent remembering to check access
is not access control. An audit trail that exists only when the model chooses to write to it
is not an audit trail. A coding convention that the model follows when it happens to recall
the rule is not a constraint. It is a suggestion with good intentions.
This is the category error underneath a lot of AI infrastructure right now. Tool-calling is
useful for optional actions. Governance is not optional. Treating both as the same kind of
model-invoked tool makes the system look extensible while leaving the most important controls
dependent on the model's cooperation.
A tool a model can choose to call is optional input. Governance that matters
cannot be optional. The two are currently built the same way, and that is the problem.
Enterprise AI needs a context control plane
The missing layer is the AI context control plane: the infrastructure that decides what context
an AI request is allowed to carry, what memory it can use, what policy applies, and what gets
recorded before the request reaches the model.
This is not the same thing as an agent framework. It is not the same thing as a memory SDK. It
is not an MCP server. Those sit beside the model as capabilities the agent may invoke. A context
control plane sits in the path of the request. It can govern, enrich, scope, and audit the
request before the model ever sees it.
Tenure starts with memory because memory is the first unavoidable primitive of this layer.
Every serious AI workflow eventually needs durable state: what was decided, what changed, what
was superseded, who approved it, what scope it applies to, and why it was injected. If that
state lives inside one app, every other app loses it. If it lives behind a tool call, the model
can skip it. If it lives underneath the request path, it becomes infrastructure.
Tenure is the AI context control plane for enterprise teams, starting with governed
memory and state beneath every model request.
People are already circling this from different directions
A coding agent can produce a diff, run a test, and report that the task is done. The harder
question, the one a lot of engineers are now asking out loud, is what happens after. Someone
still has to know what changed, why it changed, whether the right files were touched, what was
tested, what was skipped, and whether the result is safe to trust. That is not a code generation
problem. It is a review and trust problem.
A separate group is approaching the same wall from the orchestration side, building systems
that explicitly refuse to let an agent's own "done" signal count as authority. The agent
produces a claim. Something else, observing the actual state of the repository, the tests, and
the review, decides whether that claim holds. That is a real and reasonable design. It also
quietly admits the thing it does not solve by itself: someone still has to define what good
means for a given repo, and keep that definition from drifting out of sync with what the agent
was actually told.
A third group is rethinking memory itself along the same axis, arguing that most agent memory
today is something you query after the fact, and that the more valuable version is something
the agent has to reconcile before it acts and correct as part of acting. Different vocabulary,
same underlying complaint: the thing that is supposed to keep an agent honest is currently
positioned as optional, late, or both.
Three different entry points. One structural cause. Nobody arrived at it from the same angle,
which is usually a sign that the problem is real rather than a framing exercise.
This already happened underneath the application
A container does not ask permission before it sends a packet. It does not call a tool to check
whether it is allowed to talk to another pod. The network policy is not a service the container
can invoke if it remembers to. It is enforced underneath the workload, at a layer the workload
cannot opt out of. The workload does not cooperate with the policy. The policy does not need it
to.
In infrastructure, many controls begin as application-level conventions before they become
enforcement layers. That shift matters because "the app should remember to check" is not a
security property. It is a hope.
The same move is now available for agents. Almost nobody is positioned to make it because
almost everything in this space is built as a tool the agent calls rather than a layer the
request cannot get around.
Context is fragmenting across tools
The average enterprise AI environment will not be one model, one assistant, or one agent. It
will be a mix of coding tools, chat clients, internal agents, CI workflows, notebooks, ticketing
systems, support workflows, and model providers. Each one will accumulate its own context.
Each one will have its own memory model. Each one will make different assumptions about what
it knows and why it knows it.
That is manageable while AI usage is personal and experimental. It becomes a governance problem
the moment AI participates in real work. A team decision made in one tool should not disappear
when a developer switches clients. A policy correction should not depend on being manually
copied into every agent. A stale architectural fact should not keep resurfacing because one
application remembered it and another one did not.
App-level memory creates islands. Tool-call memory creates optional recall. Enterprise AI needs
neither. It needs a shared, governed substrate for context that can sit below the tools and
above the models.
The future of enterprise AI is not one app with better memory. It is many apps sharing one
governed context layer.
The layer underneath still needs state
Policy without memory is brittle. Audit without memory is incomplete. Evaluation without memory
has no durable view of what the system believed at the time. The layer underneath AI requests
needs state because real organizations are made of accumulated decisions, exceptions,
preferences, corrections, and rules.
That state cannot behave like a pile of semantically similar notes. A system that remembers
"the connection pool cap is 20" and later remembers "the connection pool cap is 50" has not
solved memory if both facts can keep returning forever. It has created a permanent read-time
cleanup problem. Every future reader has to notice the conflict again.
Governed memory moves that work into the substrate. Facts need scope. Corrections need
provenance. Superseded beliefs need to be marked as superseded. The current value needs to be
available without forcing every model request to rediscover the history from scratch.
This is why Tenure treats memory as state, not search. The point is not to retrieve more
nearby text. The point is to know what the system currently believes, why it believes it, where
that belief applies, and whether it is still allowed to influence the next request.
Under the request path, not beside the model
Tenure sits between AI clients and model providers. That position matters. It means context
injection, scoping, audit trails, and memory governance do not depend on the model choosing to
call a tool. They happen before the request reaches the model.
In practice, this lets teams use the AI tools they already use while giving the organization
a shared layer for governed context. A developer can work in VS Code, Cursor, Cline, Continue,
Open WebUI, or another client, while the memory and policy layer remains consistent underneath.
The client can change. The model can change. The governed state does not have to disappear
with either one.
That is the difference between a memory feature and a context control plane. A feature improves
one application. A control plane gives the organization one place to decide what AI is allowed
to know, remember, use, and prove.
MCP lets a model call tools. Tenure governs the request before the model sees it.
Does it require the agent's cooperation?
The simplest test for this category is not whether a product says governance, memory, policy,
or audit. The test is whether the system requires the agent's cooperation to work.
If the model has to remember to call the memory tool, memory is optional. If the model has to
remember to call the policy checker, policy is optional. If the model has to remember to write
an audit event, the audit trail is optional. However good those tools are, they are still
downstream of the model's discretion.
A context control plane has a different contract. The request enters governed infrastructure
before it reaches the model. Identity, scope, memory, policy, provenance, and audit can be
attached outside the model's decision loop. The model does not need to know the enforcement
happened. It also does not get to skip it.
If it needs the agent's cooperation, it is a tool. If the request cannot get around it, it is
infrastructure.
The model is the part everyone can see. The control plane is the part that decides whether it is safe to trust.
None of this means tool-calling is the wrong abstraction for everything. Fetching a file,
querying a database, opening a ticket, or running a build can all be discrete actions that a
model invokes when needed.
The argument is narrower and more important: anything that has to be true regardless of what
the model decides cannot live at the same layer as an optional tool. Access control. Audit
trails. Memory scope. Supersession. The rules generated code has to conform to before anyone
trusts it. These cannot depend on the model deciding to ask.
The people asking what happens after the diff, the people refusing to let an agent's "done"
mean done, the people rebuilding memory around write-time correction instead of read-time
lookup, are all describing the same missing layer from different rooms.
The category is becoming visible now. Tenure is building it from the bottom up: governed
memory and state first, then the broader control plane for context, policy, provenance, audit,
and evaluation across enterprise AI requests.
Why MCP is the wrong abstraction for memory
Memory is not something an agent should remember to invoke. It should already be present when the request reaches the model.
Access control, audit trails, and memory scope enforced at the proxy, not at the agent's discretion.
The pass/fail discipline behind governed memory, measured against 89 cases and five comparison systems.

---
source: "https://jeroensoeters.github.io/towards-the-dark-factory/"
hn_url: "https://news.ycombinator.com/item?id=48936678"
title: "Towards the dark factory: Autonomous, isolated AI agents in software engineering"
article_title: "Towards the dark factory: Autonomous, isolated AI agents in software engineering — Jeroen Soeters"
author: "discountelf"
captured_at: "2026-07-16T17:04:47Z"
capture_tool: "hn-digest"
hn_id: 48936678
score: 2
comments: 1
posted_at: "2026-07-16T16:26:26Z"
tags:
  - hacker-news
  - translated
---

# Towards the dark factory: Autonomous, isolated AI agents in software engineering

- HN: [48936678](https://news.ycombinator.com/item?id=48936678)
- Source: [jeroensoeters.github.io](https://jeroensoeters.github.io/towards-the-dark-factory/)
- Score: 2
- Comments: 1
- Posted: 2026-07-16T16:26:26Z

## Translation

タイトル: 暗い工場に向けて: ソフトウェア エンジニアリングにおける自律的で孤立した AI エージェント
記事のタイトル: 暗い工場に向けて: ソフトウェア エンジニアリングにおける自律的で孤立した AI エージェント — Jeroen Soeters
説明: 「ポンコツ インフラストラクチャ」とは、formae 自体がプロビジョニングされたロックダウンされたコンテナ工場を通じてチケットを受け取り、設計、実装、PR を出荷する、自律的でネットワークから分離された AI エージェントを構築した方法です。

記事本文:
ソフトウェア開発、インフラ、AIに関するメモ。
闇の工場に向けて: ソフトウェア エンジニアリングにおける自律的で孤立した AI エージェント
2025 年後半は、ソフトウェア開発における AI の転換点でした。それまでは、ほとんどの時間をエディターでコードを書くことに費やしていました (もちろん神の意図どおり vim)。私は主にオートコンプリートや、素晴らしい avante プラグインによる小さなパッチの生成に AI を使用しました。 Opus 4.5 のリリース後、ワークフローは完全に変わりました。私は Claude Code に切り替え、決して過去を振り返ることはありませんでした。
大まかに言うと、私のワークフローは次のようになります。私は新しいタスクを選択し、superpowers プラグインを使用してクロードとブレインストーミング セッションを開始します。クロードは設計仕様を作成し、私は Codex に設計の敵対的レビューを実行させることで、モデル全体でそれを検証します。私は調査結果をトリアージし、クロードに行動に移す価値のあるものを組み込むように依頼し、大幅なやり直しであれば、洗浄と繰り返しを行います。そうして初めて、私自身が完全な仕様をレビューします。それが成功したら、クロードは実装計画を書き、複雑さに応じて私は作業をサブエージェントに委任するか、現在のセッションに留まります。完成したコードは、私自身が確認する前に別のモデル間レビューを経て、PR をスカッシュマージするか、エージェントを送り返して残った部分を修正します。
このフローは、インタラクティブな設計セッションから恩恵を受ける実質的な機能に適しています。しかし、小規模なチームが多くのさまざまなテクノロジに触れるスタートアップでは (定義上、IaC ツールがそれを行います)、多くの作業は千切れるようなものです。多くのチケットは小さいものです。ログ行の重要度を変更し、リソース X の JSON 解析を修正し、一時的な移行コードを削除します。これらの設計と実装は簡単で、コーディング エージェントを使用してステップ実行するのは機械的なものではなく、機械的なもののように感じられます。

人間を必要とするのはngです。
ここ数週間、私はこれに取り組むために時間の一部を確保しました。その結果が、現在私たちがポンコツ インフラストラクチャと呼んでいるものです。クランカーは、上で概説した機械的なステップを正確に実行する自律エージェントです。
マシンを構築するマシンを構築する
Clanker は、私が毎日使用しているのと同じツールで構築されています。ハーネスとしての Claude Code、厳選されたスキルのセット、さまざまな MCP サーバー、モデル間レビュー用の Codex です。しかし、私がハーネスから調整する代わりに、ポンコツたちが勝手に走ります。これらは当社の問題トラッカーからトリガーされ、当社が管理するコンテナ内でヘッドレスで実行されます。
ポンコツが仕事をするたびに、これをターンと呼びます。デザインを提案するのが番、PR を開くのが番、質問に答えるのが番です。各ターンの後、ポンコツはバトンを人間に返します。コンテナは使い捨てであり、ターンの間に何も保持されません。永続的なものはすべて、問題トラッカーと PR ブランチに保存されています。
私たちは問題追跡ツールとして Linear を使用しています。 Linear は、自律エージェント インフラストラクチャのサーフェスでもあります。その機能の 1 つにより、OAuth アプリを通じてアプリ ユーザーをプロビジョニングできます。ワークスペース メンバーは、他のメンバーと同様に @メンション可能で割り当て可能ですが、人ではなくアプリによってサポートされます。 「Clanker」もその 1 つで、チケットを割り当てることで、チームメイトに手作業で作業するのとまったく同じ方法で作業を行います。その 1 つの名前の背後にはフリートがあります。各ターンは独自の新鮮な動作で実行され、多くが並行して動作する可能性があります。
これがリニアボードの外観です。
ボード: バックログ → 設計 → 実装 → レビュー中 → 完了
バックログからチケットが取得されると、まず「設計」に進み、そこで技術設計が検討され、技術概要が作成され、リニアの問題に添付されます。たった一度だけ

設計が承認されると、チケットは「実装」に進み、そこで実際のコーディングが行われます。 PR が完了すると、チケットは「レビュー中」になり、マージされると「完了」に移動します。
この列は、Clanker がどのような作業を行うかを示します。 「デザイン」列で Clanker を割り当てると、技術概要が作成されます。阻害要因や決定事項はチケット所有者にエスカレートされ、Clanker はチケットにコメントして自身の割り当てを解除します。これを「実装」列に割り当てると、実装計画が作成され、コーディングがサブエージェントに委任され、PR が開きます。そこから、Linear だけでなく PR を通じて Clanker と対話することもできます。変更をリクエストすると、修正、承認が行われ、PR がマージされ、チケットが「完了」に移動します。
Clanker をチケットに割り当てる
Linear のもう 1 つの優れた機能は、エージェント セッションです。エージェントが仕事を受け取ると、その考えや行動がその問題に関する読み取り専用スレッドにストリーミングされるため、セッション内で何が起こっているかを知ることができます。
エージェント セッション スレッド: クランカーの動作に合わせてストリーミングされる思考とアクション
エージェントブリッジ
Linear とエージェントの間には、エージェント ブリッジと呼ばれる小規模な常時稼働のオーケストレーション サービスが存在します。これは ECS Fargate サービスであり、チケット上のアクティビティを実行中のエージェントに変えるものです。 Clanker をチケットに割り当てるか、後で PR で Clanker を操作すると、そのイベントがエージェント ブリッジに到達し、実行するターンの種類が判断され、それを実行するための新しいコンテナが生成されます。
その前に、必要なリポジトリのみを対象とした短期間の GitHub トークンを生成するため、エージェントは広範で長期間有効な認証情報を保持することなくブランチをプッシュして PR を開くことができます。コンテナが完了すると、バトンを戻して終了します。
ブリッジ: Webhook (実線) ラベル左 -->
ウェブフック
リニア :

GraphQL (点線) ラベル右 -->
グラフQL
クランカー: ecs:RunTask (実線) ラベルが左 -->
ecs:タスクの実行
ブリッジ：ハンドバック（点線）ラベル右 -->
手を戻します
GitHub : PR を開く / 更新 / マージ (固体) -->
PR を開く / 更新 / マージする
ブリッジ: Webhook (ソリッド) -->
ウェブフック
リニア
エージェントブリッジ
ポンコツ
GitHub
Agent Bridge は、リニア イベントと GitHub イベントをターンごとの使い捨てのポンコツ コンテナに変換します。
各ターンは、チケットが属する Linear プロジェクトに合わせて調整された Docker イメージからスピンアップされた独自のコンテナーで実行されるため、エージェントはそのターンに必要なものだけを思いつきます。すべてのコンテナは共有基盤、つまりグローバルな運用体制、つまりすべてのプロジェクトのあらゆる段階を管理する CLAUDE.md から始まります。ターン モデル、つまり一度に 1 ターンずつ行うこと、自動レビュー パス、単独で決して越えてはいけない人間の門などについて説明します。また、テストファースト開発、周囲のコードの規則との一致、YAGNI などのエンジニアリングの実践についても取り上げます。そして方向性: 公開ドキュメントと社内ドキュメントがどこに置かれているか、そして GitHub 組織全体でどこで情報を見つけられるか。そのグローバル ベースの上に、各プロジェクトが必要なものを重ねていきます。
Linear プロジェクトに関連付けられた GitHub リポジトリまたはリポジトリ。
リポジトリ固有の規則を持つ「ローカル」CLAUDE.md がイメージに焼き付けられます。
AWS プラグイン用の AWS CLI や、Fresh フレームワーク上に構築されたパブリック ハブ用の Deno など、作業に必要なツール。
コンテナーに必要なアクセス権 (エージェントがそこで失敗したテストをデバッグできるようにするための専用のテスト CI アカウントの資格情報など)。
そのプロジェクトをスコープとする MCP サーバー。
コンテナには、ジョブを実行するために必要なものが正確に含まれており、必要のないものは何もありません。 formae コア チケット、formae ハブ チケット、および AWS プラグイン チケットはそれぞれ、異なるイメージ、適切なツール、および

適切なコンテキストが自動的に選択されます。
これらのコンテナは、ライブ認証情報を保持する LLM を通じて、問題の説明や PR コメントなど、攻撃者の影響を受けるテキストをフィードする可能性があります。これは即時注入のリスクがあるため、境界はロックダウンされています。コンテナにはオープンなインターネットへのルートがありません。これらは隔離されたサブネットに存在し、すべての送信リクエストは明示的な FQDN ホワイトリストを持つファイアウォールを通過します。作業に本当に必要な少数のエンドポイント (GitHub、Anthropic および Codex API、Linear、Go モジュール プロキシ、AWS) のみで、他には何もありません。それ以外はすべて削除され、ログに記録されます。認証情報はターンごとで最小の権限です。つまり、範囲指定された GitHub トークン、最小限の IAM ロール、および Secrets Manager からのみ取得されるシークレットです。
オープンなインターネットを遮断したことによる副作用として、ポンコツは調査チケットとしては役に立ちません。それで大丈夫です。ポンコツは、オープンなソリューション空間をさまようためではなく、明確に指定された作業を実行するために存在します。
モデルはそれ自体の出力の弱点に気付かない傾向があるため、設計や差分が受け入れられる前に、ポンコツはそれを敵対的なレビューのために別のモデルに渡します。これは設計ターンでも実装ターンと同じです。今日はこれに Codex を使用します。結果が戻ってくると、ポンコツはそれらを優先順位付けし、実行する価値のあるものを組み込み、残りを却下します。結果として生じた変更が大幅な場合は、再度レビューが行われ、作業がきれいに戻るまで続行されます。これらはすべて、自動レビューと呼ばれるカスタム スキルを通じて実行されます。
私たちはトリアージのステップに多大な労力を費やしました。 Codex は、ここで重要かどうかに関係なく、理論的に問題となる可能性のあるものすべてにフラグを立てるため、その結果は盲目的に信頼するのではなく、検討する必要があります。スキルにはそれを行うためのコンテキストが含まれています。コードベースがすでに提供している保証が考慮されます。

たとえば、formae とそのプラグインはアクター モデルに基づいて構築されているため、アクター モデルでは発生し得ない、フラグが立てられたレースや共有ステートのハザードは割引されます。機能がまだ稼働しているかどうかを検証するため、ユーザーやデータを壊す必要がない場合、下位互換性と移行の結果は脇に置かれます。査読者はすべてを明らかにします。ここで実際に何が重要かはスキルによって決まります。
正確さと安全性は例外です。そこでの発見は、単なる「拒否」で無視することはできません。それは実際の修正、それが適用されないという具体的な議論、または人間の承認によってのみ解決されます。同じポンコツ人間がコードを書き、どの発見を尊重するかを決定しているのですから、重大な発見を黙って葬り去ることは許されません。
これらすべてが整っているにもかかわらず、私は同じ状況に陥り続けました。顧客デモなど、いくつかの機能に関する Claude Code セッションに深く入っていると、別の場所での小さな変更、新しいリソース タイプのサポート、または共通ライブラリのバグ修正が必要であることに気づきます。そこで私は、クロードにチケットを提出し (Linear には優れた MCP があり、これをカスタム スキルのセットにまとめました)、Clanker を割り当て、生成された設計をレビューして「実装」に移動し、Clanker を再度割り当て、問題が解決したと確信したら PR を承認するように依頼します。
そのループ自体が機械的だったので、すぐにそれ自体がスキルになりました。クランクスルーは既存のチケットを取得し、クランカーをエンドツーエンドで監視し、本当に決定が必要な場合にのみ、全面的に処理して私にエスカレーションします。つまり、エージェントがエージェントに仕事を委任し、エージェントに仕事を委任するのです 🤖。
これが現状の私たちの工場です。仕事はパイプラインを通って自動的に進み、私はポンコツが本当にブロックされた場合にのみ介入します。
私たちがポンコツに任せる仕事の種類は変わりません

サイズによって味方が制限されます。クランカーは、製品の結果が明確であり (クランカーが製品の決定を下すべきではありません)、技術設計が明らかであるか、重要ではないか、またはクランカーがパターン マッチングできる既存のものに十分近い限り、非常に複雑なタスクを引き受けることができます。それを超えると、成果は逓減します。問題のコメントを通じてトリッキーな設計を練り上げるのは骨が折れますが、インタラクティブなセッションでははるかに簡単です。設計が決まったら、それを課題に添付し、クロード コードで自分で実行するのではなく、ポンコツに実装を渡します。
具体的には、私のダークファクトリーのメンタルモデルは、製品仕様→技術設計と技術設計→動作するソフトウェアの 2 つの生産ラインです。作業はどちらにも注入できます。どこに注入するかは、生成されるコードを考慮するかどうか、および設計の複雑さによって異なります。ここでの複雑さは、先行技術がないことを前提としています。つまり、非常に複雑なものになる可能性がありますが、それが何百万回も実行されていれば、LLM にとっては些細なことです。
製品
仕様
技術的な
デザイン
製品
仕様
製品
仕様
作業が注入される場所: コードの所有者 × 設計の複雑さ。
本当の暗い工場では、両方のラインが消灯状態で稼働します。製品仕様は入力され、ソフトウェアは動作し、

[切り捨てられた]

## Original Extract

How we built "clanker infrastructure" — autonomous, network-isolated AI agents that pick up tickets, design, implement, and ship PRs through a locked-down container factory provisioned with formae itself.

Notes on software development, infrastructure, and AI.
Towards the dark factory: Autonomous, isolated AI agents in software engineering
Late 2025 was an inflection point for AI in software development. Before that I spent most of my time writing code in an editor (vim of course, as God intended). I used AI mostly for auto-complete or generating small patches with the awesome avante plugin. After the release of Opus 4.5 the workflow changed completely: I switched to Claude Code and never looked back.
On a high level my workflow looks something like this. I pick up a new task and kick off a brainstorming session with Claude, using the superpowers plugin. Claude writes a design spec, and I validate it across models by having Codex run an adversarial review of the design. I triage the findings and ask Claude to incorporate the ones worth acting on, and if the rework is substantive we rinse and repeat. Only then do I review the full spec myself. Once it holds up, Claude writes the implementation plan, and depending on complexity I either delegate the work to subagents or stay in the current session. Finished code goes through another cross-model review before I look at it myself, and then I either squash-merge the PR or send the agent back to fix what’s left.
This flow works well for substantive features that benefit from an interactive design session. But at a startup with a small team touching many different technologies, which an IaC tool does by definition, a lot of the work is death by a thousand cuts. Many tickets are tiny: change the severity of a log line, fix the JSON parsing of resource X, remove some temporary migration code. The design and implementation of these is trivial, and stepping through it with a coding agent feels mechanical, not something that should need a human.
Over the last few weeks I set aside a portion of my time to tackle this, and the result is what we now call our clanker infrastructure . Clankers are autonomous agents that do exactly the mechanical steps outlined above.
Building the machine that builds the machine
A clanker is built out of the same tools I use every day: Claude Code as the harness, a curated set of skills, various MCP servers, Codex for cross-model reviews. But instead of me orchestrating from the harness, clankers run on their own. They are triggered from our issue tracker and run headless in a container we control.
Every time a clanker does work we call this a turn. Proposing a design is a turn, opening a PR is a turn, answering a question is a turn. After each turn the clanker hands the baton back to a human. The container is disposable and nothing persists between turns; everything durable lives in our issue tracker and on the PR branch.
We use Linear as our issue tracker. Linear is also the surface for our autonomous agent infrastructure. One of its features lets you provision app users through an OAuth app: workspace members that are @mentionable and assignable like anyone else, but backed by an app instead of a person. “Clanker” is one of these, and you hand it work exactly the way you’d hand work to a teammate, by assigning a ticket to it. Behind that single name is a fleet: every turn runs in its own fresh clanker, and many can be working in parallel.
This is what our Linear board looks like.
The board: Backlog → Design → Implementation → In Review → Done
When a ticket gets picked up from the backlog it first goes to “Design,” where the technical design gets worked out and a tech brief is produced and attached to the Linear issue. Only once that design is approved does the ticket move to “Implementation,” where the actual coding happens. Once the PR is up the ticket lands “In Review,” and once merged it moves to “Done.”
The column dictates what work Clanker will do. Assign Clanker in the “Design” column and it produces the tech brief; any blockers or decisions get escalated to the ticket owner, with Clanker commenting on the ticket and unassigning itself. Assign it in the “Implementation” column and it writes an implementation plan, delegates the coding to subagents, and opens a PR. From there you can also interact with Clanker through the PR, not just in Linear: request changes and it revises, approve and it merges the PR and moves the ticket to “Done.”
Assigning Clanker to a ticket
Another cool feature of Linear is agent sessions. Whenever an agent picks up work, it streams its thoughts and actions into a read-only thread on the issue, so you get a window into what’s happening inside the session.
The agent session thread: thoughts and actions streaming as the clanker works
The Agent Bridge
Between Linear and the agents lives a small always-on orchestration service we call the Agent Bridge. It is an ECS Fargate service, and it is what turns activity on a ticket into a running agent. When you assign Clanker to a ticket, or later interact with it on its PR, that event reaches the Agent Bridge, which figures out what kind of turn to run and spawns a fresh container to run it.
Before it does, it mints a short-lived GitHub token scoped to just the repositories that turn needs, so the agent can push a branch and open a PR without ever holding a broad, long-lived credential. When the container is done it hands the baton back and exits.
Bridge : webhook (solid) label left -->
webhook
Linear : GraphQL (dotted) label right -->
GraphQL
clanker : ecs:RunTask (solid) label left -->
ecs:RunTask
Bridge : hand back (dotted) label right -->
hand back
GitHub : open / update / merge PR (solid) -->
open / update / merge PR
Bridge : webhook (solid) -->
webhook
Linear
Agent Bridge
clanker
GitHub
The Agent Bridge turns Linear and GitHub events into per-turn, disposable clanker containers
Each turn runs in its own container, spun up from a Docker image tailored to the Linear project the ticket belongs to, so the agent comes up with exactly what that turn needs and nothing more. Every container starts from a shared foundation: a global operating constitution, a CLAUDE.md that governs every turn in every project. It covers things like the turn model: one turn at a time, the auto-review pass, and the human gates it must never cross on its own. It also covers engineering practices: test-first development, matching the conventions of the surrounding code, YAGNI, and the like. And orientation: where our public and internal docs live and where to find things across our GitHub org. On top of that global base, each project layers in what it needs:
the GitHub repo or repos associated with the Linear project;
a “local” CLAUDE.md with repo-specific conventions, baked into the image;
the tools the work requires, like the AWS CLI for the AWS plugin or Deno for our public hub, which is built on the Fresh framework;
the access the container needs, such as credentials for a dedicated test CI account so the agent can debug failing tests there;
the MCP servers scoped to that project.
The container arrives with exactly what it needs to do the job and nothing it doesn’t. A formae core ticket, a formae hub ticket, and an AWS plugin ticket each get a different image, the right tooling, and the right context, selected automatically.
These containers can potentially feed attacker-influenceable text, like issue descriptions and PR comments, through an LLM that holds live credentials. That is a prompt-injection risk, so the boundary is locked down. The containers have no route to the open internet. They sit in isolated subnets, and every outbound request goes through a firewall with an explicit FQDN allowlist: the handful of endpoints the work genuinely needs (GitHub, the Anthropic and Codex APIs, Linear, the Go module proxy, AWS) and nothing else. Everything else is dropped and logged. Credentials are per-turn and least-privilege: that scoped GitHub token, a minimal IAM role, and secrets that only ever come from Secrets Manager.
A side effect of cutting off the open internet: a clanker is useless for research tickets. We’re fine with that. The clankers exist to execute well-specified work, not to wander an open solution space.
A model tends to be blind to the weaknesses in its own output, so before a design or a diff is accepted, a clanker hands it to a different model for an adversarial review, the same on a design turn as on an implementation turn. We use Codex for this today. The findings come back, the clanker triages them, incorporates the ones worth acting on and dismisses the rest. If the resulting change was substantial, it reviews again, and keeps going until the work comes back clean. All of this runs through a custom skill we call auto-review.
We put a lot of effort into the triage step. Codex flags everything that could theoretically be a problem, whether or not it matters here, so its findings need weighing rather than blind trust. The skill carries the context to do that. It takes into account the guarantees the codebase already provides, for example that formae and its plugins are built on the actor model , so a flagged race or shared-state hazard that can’t occur under the actor model gets discounted. It verifies whether a feature is even live yet, so backwards-compatibility and migration findings get set aside when there are no users or data to break. The reviewer surfaces everything; the skill decides what actually matters here.
Correctness and security are the exception. A finding there can’t be waved away with a bare “rejected”; it clears only with a real fix, a concrete argument that it doesn’t apply, or a human’s sign-off. The same clanker wrote the code and decides which findings to honor, so it can’t be allowed to quietly bury a serious one.
With all of this in place, I kept finding myself in the same situation. I’d be deep in a Claude Code session on some feature, for example a customer demo, and realize I needed a small change somewhere else, support for a new resource type, or a bugfix in a common library. So I’d ask Claude to file the ticket (Linear has an excellent MCP, which we’ve wrapped in a set of custom skills), assign Clanker, review the design it produced, move it to “Implementation,” assign Clanker again, and approve the PR once I was confident it solved the problem.
That loop was itself mechanical, so it quickly became its own skill. clank-through takes an existing ticket and oversees a clanker end to end, walking it across the board and escalating to me only when something genuinely needs a decision. So there we are: agents delegating work to agents delegating work to agents 🤖.
This is our factory as it stands: work moves through the pipeline on its own, and I only step in when a clanker is genuinely blocked.
The type of work we hand to clankers isn’t really limited by size. A clanker can take on quite complex tasks, as long as the product outcome is clear (a clanker should never be making product decisions) and the technical design is either obvious, not important, or close enough to something that already exists for the clanker to pattern-match against. Beyond that you hit diminishing returns: working out a tricky design through issue comments is painful, and much easier in an interactive session. Once the design is settled I attach it to the issue and hand the implementation to a clanker, rather than driving it myself in Claude Code.
Concretely, my mental model of the dark factory is two production lines: product specification → technical design and technical design → working software. Work can be injected into either one. Where I inject it depends on whether I care about the code that will be produced and on the complexity of the design. Complexity here assumes no prior art: something can be highly complex, but if it has been done a million times before it is trivial for an LLM.
Product
specification
Technical
design
Product
specification
Product
specification
Where work gets injected: who owns the code × complexity of the design.
A true dark factory would run both lines lights-out: product spec in, working software out,

[truncated]

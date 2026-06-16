---
source: "https://numtide.com/blog/writing-a-kubernetes-operator-in-the-age-of-ai/"
hn_url: "https://news.ycombinator.com/item?id=48552471"
title: "What I learned using AI to build a Kubernetes Operator for Supabase's Multigres"
article_title: "Writing a Kubernetes Operator in the Age of AI - Numtide"
author: "DevOpsy"
captured_at: "2026-06-16T10:42:03Z"
capture_tool: "hn-digest"
hn_id: 48552471
score: 2
comments: 0
posted_at: "2026-06-16T08:56:01Z"
tags:
  - hacker-news
  - translated
---

# What I learned using AI to build a Kubernetes Operator for Supabase's Multigres

- HN: [48552471](https://news.ycombinator.com/item?id=48552471)
- Source: [numtide.com](https://numtide.com/blog/writing-a-kubernetes-operator-in-the-age-of-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T08:56:01Z

## Translation

タイトル: AI を使用して Supabase の Multigres 用の Kubernetes Operator を構築して学んだこと
記事のタイトル: AI 時代の Kubernetes オペレーターの作成 - Numtide
説明: AI を使用して本番環境の Kubernetes オペレーターを構築するのに 6 か月と 1,303 のコミットを行った結果、コードは簡単な部分であることがわかりました。実際の作業は設計、衛生管理、レビューであり、自分でオーケストレーションしたフレッシュコンテキストのエージェントの工場として実行されます。

記事本文:
AI 時代の Kubernetes オペレーターの作成
2026年6月15日 — フェルナンド・ビラルバ
ユーザー向けの仕様は、逸脱できない唯一のものとして扱います。それ以外はすべて低コストでリファクタリングできます。契約はそうではありません。
AI フレームワークをインストールしないでください。それらを読んでアイデアを盗み、代わりに自分のスキルを書きましょう。
機械的な作業 (レビュー、監査、コミット メッセージ、変更ログ、ドキュメント チェック) をフレッシュ コンテキストのエージェントの工場として実行します。各エージェントには 1 つの狭いジョブがあり、制御するプロセスによって調整されます。開発に一貫性を持たせるために、それらをチームと共有します
スキルが何かを通過させたら、スキルを修正します。不良出力はラインの欠陥であり、一時的なノイズではありません。
バグ監査には、事前にロードされた設計コンテキストと、幻覚をフィルタリングするための 2 番目のエージェントが必要です。そうしないと、誤検知に溺れてしまいます。
同じ AI ソースからのテストとコードには、同じ盲点があります。 100% のコード カバレッジにこだわるのではなく、実際の実行時の動作を検証します。これは特にグリーンフィールド プロジェクトに当てはまります。
AI は悪いアイデアが悪いとは教えてくれません。それの洗練されたバージョンを構築するだけです。すべての設計呼び出しは依然として人間の判断に委ねられています。
AI を使用して Multigres 用の実稼働 Kubernetes オペレーターを構築してから 6 か月後、コードは難しい部分ではなくなりました。ほとんどの努力は、設計、衛生管理、継続的な改善、レビューに費やされました。 1,303 件のコミットで私が学んだことは、AI を使用する最良の方法は、自分で調整するスペシャリストの工場フロアを構築することだということです。
私たちは、Vitess の作成者である Sugu Sougoumarane が率いる Multigres チームと緊密に連携して、新しい分散 Postgres システムを構築する契約を結びました。オペレーターの仕事は、Kubernetes 上で Multigres のプロビジョナーとして機能し、ユーザーがそれを簡単かつ宣言的に使用できるようにすることでした。
Kubernetes オペラトを構築したことがない読者向け

以前は、「分散データベースのオペレーター」は、思っているよりもはるかに複雑な仕事でした。オペレーターはセル、シャード、および接続プールをプロビジョニングし、各部分をトポロジー サービスに登録し、ユーザーが YAML に入力したものに向かって実行中のクラスターを収束させ続けます。
難しいのは状態遷移です。スケールダウンではポッドを削除するだけでは済みません。そのポッドがプライマリ Postgres である場合、ポッドを削除するとデータが失われます。 Kubernetes が安全にガベージ コレクションを行う前に、オペレーターはレプリケーション ステート マシンと調整し、プライマリを降格し、スタンバイが追いつくのを待ち、トポロジから登録を解除する必要があります。スケールアップには、StatefulSet では実行できない独自のダンスがあるため、独自のポッド管理を構築しました。
これらすべては Kubernetes の宣言型モデルを通じて行われます。このモデルでは、どの部分もいつでも失敗する可能性があり、オペレーターの仕事は、世界を望ましい形状に向けて誘導し続けることです。
以下の 3 つのことが、この仕事を平均的なオペレーターのビルドよりも困難にしました。
Multigres はまだ開発の非常に初期段階にあったため、私たちがオペレーターを設計して実装するための移動目標となりました。
複数の CRD およびコンポーネント タイプにわたる比較的複雑なテンプレート ロジックが必要でした。
Vitess と同様に、Multigres は非シャード エンジン上のミドルウェアであり、多くのデータベース層のオーケストレーションをデータベースではなくオペレーター自体にプッシュします。
最後の点については、配置とリバランスは実際にどこで行われるのでしょうか?
CockroachDB、TiDB、YugabyteDB などのデータベースでは、答えはデータベースの中にあります。ノードを追加すると、自動的にデータの取り込みが開始されます。 1 つを廃止すると、安全に削除できる前にクラスターによってデータが移動されます。大規模イベントにおけるオペレーターの仕事は主に、意図を伝えて待つことです。
Vitess と Multigres は設計上、動作が異なります。彼らはそれを保ちます

エンジン (MySQL または Postgres) は変更されていません — ストック、シャード化されておらず、クラスター メンバーシップの概念はありません。 SQL 互換性、成熟したツール、シャード分離をすぐに利用できます。これが重要です。トレードオフとして、リシャーディングと廃止はデータベース内で実行できないため、外部で調整する必要があります。これにより、オペレーターに責任と複雑さが課​​せられ、YAML 仕様の充実が強制されます。これは、十分に理解されているストレージ層のより重いコントロール プレーンの代償となります。
実際には、これはオペレーターが Postgres 層のオーケストレーションを直接操作する必要があることを意味します。つまり、ポッドがプライマリであるかどうかを確認し、降格を調整し、同期スタンバイが追いつくのを待ち、トポロジ サービスから登録を解除してから、Kubernetes にポッドを削除させます。
このドメインを理解するのにかなりの時間を費やしましたが、それは有意義に費やされました。それには、Multigres コードベースと設計ドキュメントをレビューするだけでなく、すべてのオペレーターとコントローラーのランタイム機能をレビューして、何が必要か、どのように作業を設計するかを確認する必要がありました。
この段階で AI をコンパニオンとして使用することも、コードベース全体とドキュメントを AI に提供し、どんなに愚かなことであっても、あらゆることについて質問できるため、非常に役立ちました。これにより、調査期間が大幅に短縮され、簡単になり、Multigres チームの状況を把握するための大きな負担が軽減されました。
仕様は元をとった決断だった
実装を開始する前に、設計プロセスの一環として仕様を作成しました。仕様主導の開発はウォーターフォール劇場として軽視されますが、その理由は理解できます。これに対する歴史的な反論は、仕様を完成させる頃には世界は動いており、仕様は間違っているというものでした。それは今でも部分的には真実です。
AI が数学を変えるということも真実です。小さなチームの場合

さらに、コーディング エージェントは、同じ機能に対して互換性のないデザインを 1 週間で 3 つ作成することもあります。漂流させてはいけないのは、ユーザー向けの契約です。
リファクタリングが安価になったため、他のものはすべてそれほどコストをかけずにチャーンできます。ただし、ユーザーが操作する表面は安定している必要があります。エージェント、人間、クライアントのすべてを同じターゲットに向け続けるのは、単一のアンカーだからです。
そこで私たちはユーザーが読む方法で仕様を書きました。すべての代替オプションがコメントとしてインラインに含まれる CR YAML、フィールドのすぐ隣にある各フィールドの動作の平易な英語の説明、状態図に固定されるのではなく散文で記述される演算子のドレインとアップグレードのセマンティクス。デザインドキュメントというよりは製品ドキュメントのように見えましたが、それがポイントです。
この仕様には多くの再ドラフトと Multigres チームとの議論が含まれていたため、数週間かかりました。長く感じましたが、振り返ってみると、プロジェクトの中で最も重要な数週間だったと思います。 Multigres チームと仕様の反復に 1 時間費やすたびに、後で廃棄する予定だった実装の日数が節約されました。
満足のいく結果が得られると、その仕様をベースとして AI でコードのさまざまな部分を生成しました。
仕様は、オペレーターが何を行うかをユーザーに伝えます。チームに実装方法を指示するものではありません。これは別の演習であり、早期に実行する必要があります。AI は、コードベースが実際に必要とするスタイルではなく、トレーニング データから推論したあらゆるスタイルでコードを喜んで作成するからです。
押さえるべきことは大きく分けて4つあります。
コードの構造: パッケージ、階層化、名前付け。
どのベスト プラクティスに従っているか (Google の Go スタイル ガイド)。
厳格なルール: 何をするか、何をしないか、例外が許可される場合。 AI はこれらを無視することがあります。そのため、

場合によっては、最終エージェントとあなたが作業を再確認することが不可欠です。
テストの規則: 対象となる内容、深さ、ツール。
これらはいずれも AI に限ったものではありません。人間がお互いのコードを読むには、同じ共有語彙が必要です。人間とエージェントから成るチームでは、規約が共有文法となり、両方のグループが互いに会話することなく同じコードベースで作業できるようになります。
コツは、すべてのコンテキスト ウィンドウにロードされる 3,000 行の CLAUDE.md にこれらすべてを詰め込まないことです。 Go スタイル ガイドだけでは、毎回のコンテキストの半分が消費されてしまいます。
より良いパターンは、各規約を実装エージェントがデフォルトで読み取らない独自のスキルにすることです。変更が記述された後、新しいコンテキストのエージェントが関連するスキル (Go スタイル、設計ドキュメントのコンプライアンス、テスト カバレッジ ルール) を実行し、その特定の規則との差分を監査します。実装エージェントは無駄のない状態を維持し、レビュー担当者は差分とチェックしているルールのみを参照します。
規約を更新する場合は、スキルを編集します。レビュー エージェントは次回実行時に変更を取得します。
いくつかのツールを通過して、着地した場所に着地しました。誰かが行き止まりを回避できるように、アーク全体を書き留めてください。
私は Gemini Pro を普通のチャット ウィンドウで開始し、コード フォルダーに貼り付けて質問しました。研究や新しいコードベースの理解に適しています。ツールを操作するのはユーザー自身であり、エージェントが存在せず、「貼り付け、待機、再度貼り付け」のループが遅く、摩擦が多いため、継続的な実装にはうまく機能しません。
その後、Google の Antigravity に移行しました。これは、VSCode に基づく IDE のようなエージェントに Gemini Pro と Claude をバンドルしています。それがうまくいったときは素晴らしかったです。 Ultra サブスクリプションの割り当てが十分に多かったので、クレジットが足りなくなってしまいました

基本的に無理でした。しかし、バックエンドは非常に信頼性が低く、本格的な作業をイライラさせるものでした。ジョブは常に途中でキャンセルされ、手動で再試行し続ける必要がありました。サポートは遅く、問題に対してあまり責任を持っていませんでした。再試行する必要のないことを再試行することで多くの時間を費やしてしまいました。結局、私は先に進みました。
Claude Code は、私にとってエンジニアリング作業を継続するのに最適なツールであることがわかりました。私は 100 ドルの Ultra サブスクリプションの料金を支払います。エージェントが 2 つのコードベースのレビューと大量のテストの作成を伴うクエリで大量のコンテキストを消費していたため、早い段階でクレジットが不足していました。大規模なツールの出力をサンドボックスにオフロードするコミュニティ プラグインである Context Mode をインストールすると、それがなくなりました。アップグレードせずに、より安価なサブスクリプションを維持しました。すべての AI サブスクリプション プロバイダーがクォータを大幅に削減し始めている現在、コンテキスト モード プラグインはこれまで以上に重要になっています。
ツールに関するもう 1 つの重要な認識は、AI コーディング エコシステムが短期間に膨大な量のツール ( SpecKit 、 Superpowers 、 Agent-skills 、隔週の新しい 500 スター リポジトリ) を成長させたことです。彼らの中にあるアイデアはたいてい素晴らしいものです。フレームワークをインストールすると、その世界観に引き込まれます。あなたはその厳格さ、そのトークンのオーバーヘッド、そしてあなたが構築しているものとまったく同じものを構築していなかった誰かの意見を継承します。
たとえば、私はスペック駆動開発を強く信じていますが、SpecKit は私にとってピンと来ませんでした。追加されるすべてのコマンドには、CLAUDE.md、プレーンなマークダウン設計ドキュメント、GitHub の問題、スキル、ネイティブ TodoWrite など、おそらくすでにお持ちの、よりシンプルでネイティブな代替手段があります。これをインストールするということは、代理店と取引して、他人のルールブックに合わせて独自のより厳密なプロジェクト固有のバージョンを構築することを意味します。
そうすればスキルの束が手に入ります

(クロード プラグイン) スーパーパワーやエージェント スキルが好きです。スーパーパワーは、ワークツリー、TDD、ブレインストーミングを、それらが実際に行っている作業に適合するかどうかに関係なく強制します。エージェント スキルは、決して監査できない数十のスキルを構成にダンプします。どちらも素晴らしいアイデアを持っているので盗むべきですが、どちらもそのままインストールしないでください。それらを読んでから、自分のスキルを書きましょう。 AI を利用すれば、スキルを作成するのは十分安価なので、やらない理由はありません。
私が今インストールするサードパーティ ツールは、単一のジョブに限定されたものだけです。 Skill Creator を使用すると、ワークフローを実行せずにスキルを作成し、洗練させることができます。コンテキスト モードでは、コンテキストの使用が抑制されます。どちらも小さく、両方とも範囲が広いです。どちらも私のワークフローを実行しようとしていません。
初期の開発ループは、超高速のパートナーとのペア プログラミングのように見えました。仕様を読み、エージェントに指示を出し、差分を確認し、問題を修正し、テストを依頼し、それを繰り返します。私も同じように他の人のブランチをレビューし、ブランチをプルダウンして、新鮮な会話で差分を調べました。私がオペレーターがどうなりたいかをまだ学んでいる間は、それがうまくいきました。
その後、期限が迫っているグリーンフィールド プロジェクトとしては遅すぎるため、PR レビューを完全に中止しました。デザインは事前にレビューされました。実行はバグ監査とリファクタリングを通じて修正されました。私はおそらく

[切り捨てられた]

## Original Extract

Six months and 1,303 commits building a production Kubernetes operator with AI taught us that the code is the easy part — the real work is design, hygiene, and review, run as a factory of fresh-context agents you orchestrate yourself.

Writing a Kubernetes Operator in the Age of AI
June 15, 2026 — Fernando Villalba
Treat the user-facing spec as the one thing that can’t drift. Everything else is cheap to refactor; the contract isn’t.
Don’t install AI frameworks. Read them, steal the ideas, and write your own skills instead.
Run the mechanical work — reviews, audits, commit messages, changelogs, doc checks — as a factory of fresh-context agents, each with one narrow job, orchestrated by processes you control. Share them with the team so the development is consistent
When a skill lets something through, fix the skill. Bad outputs are defects in the line, not one-off noise.
Bug audits need design context loaded up front and a second agent to filter hallucinations, or you drown in false positives.
Tests and code from the same AI source share the same blind spots. Verify against real runtime behavior instead of obsessing over 100% code coverage — this is especially true on greenfield projects.
AI won’t tell you a bad idea is a bad idea. It’ll just build a polished version of it. Human judgment still owns every design call.
Six months into building a production Kubernetes operator for Multigres with AI, the code stopped being the hard part. Most of the effort went into design, hygiene, continuous improvement, and review. What I learned in 1,303 commits is that the best way to use AI is to build a factory floor of specialists you orchestrate yourself.
We were contracted to work closely with the Multigres team, led by Sugu Sougoumarane, the creator of Vitess, building a new distributed Postgres system. The operator’s job was to act as a provisioner of Multigres on Kubernetes, enabling users to use it simply and declaratively.
For a reader who hasn’t built a Kubernetes operator before, “an operator for a distributed database” is much more involved than it sounds. The operator provisions cells, shards, and connection pools, registers each piece with a topology service, and keeps the running cluster converging toward whatever the user put in a YAML.
The hard part is the state transitions. Scaling down can’t just delete a pod. If that pod is the primary Postgres, deleting it loses data. The operator has to coordinate with the replication state machine, demote the primary, wait for standbys to catch up, and unregister from the topology before Kubernetes can safely garbage-collect. Scale-up has its own dance that StatefulSets couldn’t carry, so we built our own pod management .
All of this happens through Kubernetes’ declarative model, where any piece can fail at any time and the operator’s job is to keep nudging the world toward the desired shape.
Three things made the job harder than an average operator build:
Multigres was still in very early stages of development, which made it a moving target for us to design and implement the operator against.
It required relatively complex templating logic across multiple CRDs and component types.
Multigres, like Vitess, is middleware over an unsharded engine, which pushes a lot of database-layer orchestration into the operator itself rather than into the database.
On that last point: where does placement and rebalancing actually live?
In databases like CockroachDB, TiDB, and YugabyteDB, the answer is inside the database. Add a node and it automatically starts taking on data. Decommission one and the cluster moves the data off before you can safely remove it. The operator’s job on scale events is mostly to signal intent and wait.
Vitess and Multigres work differently, by design. They keep the engine (MySQL or Postgres) unmodified — stock, unsharded, with no concept of cluster membership. You get SQL compatibility, mature tooling, and shard isolation out of the box, which is the whole point. The trade-off is that resharding and decommissioning can’t live inside the database, so they have to be orchestrated externally. That pushes responsibility and complexity into the operator and forces a richer YAML spec — the price of a heavier control plane for a well-understood storage layer.
In practice this meant our operator had to drive Postgres-layer orchestration directly: checking whether a pod was a primary, coordinating demotion, waiting for synchronous standbys to catch up, unregistering from the topology service, and only then letting Kubernetes delete the pod.
Understanding the domain took a chunk of our time, and it was well spent. It involved reviewing the Multigres codebase and design documents as well as reviewing all the operator and controller-runtime features to see what we would need and how we would design the work.
Having AI as a companion at this stage was also very helpful because we could feed it the entire codebase and documentation and ask questions, no matter how dumb, about everything. It really made the research period much quicker and easier and it took a big load off the Multigres team having to bring us up to speed.
The spec was the decision that paid for itself
Before we started implementation we wrote a spec as part of the design process. Spec-driven development gets dismissed as waterfall theater and I understand why. Historically the argument against it was that by the time you finish the spec, the world has moved and the spec is wrong. That’s still partly true.
What’s also true is that AI changes the math. If a small team plus a coding agent can produce three incompatible designs for the same feature in a week, the thing you cannot let drift is the user-facing contract.
Everything else can churn without much cost, because refactoring is now cheap. But the surface the user interacts with has to be stable, because it’s the single anchor that keeps agents, humans, and the client all aimed at the same target.
So we wrote the spec the way a user would read it. CR YAMLs with every alternative option inline as a comment, plain-English descriptions of what each field does sitting right next to the field, and the operator’s drain and upgrade semantics written out in prose rather than locked in a state diagram. It looked more like a product doc than a design doc, which is the point.
The spec took weeks because it involved a lot of redrafting and discussions with the Multigres team. That felt long, but looking back I think they were the most important weeks of the project. Every hour spent iterating on the spec with the Multigres team saved days of implementation we would have later thrown away.
Once we were happy with the result, the spec was used as a base to generate the different parts of the code with AI.
The spec tells the user what the operator does. It doesn’t tell the team how to implement it. That’s a separate exercise, and it has to happen early, because AI will happily write code in whatever style it inferred from training data — not the style your codebase actually wants.
There are roughly four things to nail down:
How the code is structured: packages, layering, naming.
Which best practices you’re following (for us, Google’s Go Style Guide ).
Hard rules: what you will do, what you will not do, where exceptions are allowed. AI sometimes ignores these, which is why additional agents and you double-checking the work are sometimes essential.
Testing conventions: what gets covered, how deep, which tools.
None of this is only for AI. Humans reading each other’s code need the same shared vocabulary. In a team that’s part humans and part agents, the conventions become the shared grammar that lets both groups work on the same codebase without talking past each other.
The trick is to not stuff all of this into a 3,000-line CLAUDE.md that loads into every context window. The Go style guide alone would eat half your context on every turn.
A better pattern is to make each convention its own skill that the implementation agent doesn’t read by default. After a change is written, a fresh-context agent runs the relevant skill (Go style, design-doc compliance, test-coverage rules) and audits the diff against that specific convention. The implementation agent stays lean, and the reviewer only sees the diff and the rule it’s checking.
When you want to update a convention, you edit the skill. The review agent picks up the change the next time it runs.
I went through several tools to land where I landed. Writing down the whole arc in case it helps someone skip the dead ends.
I started with Gemini Pro in a plain chat window, pasting in code folders and asking questions. It works well for research and for understanding a new codebase. It doesn’t work well for sustained implementation because you’re the one operating the tools, there’s no agent, and the loop of “paste, wait, paste again” is slow and full of friction.
Then I moved to Google’s Antigravity , which bundles Gemini Pro and Claude in an IDE-like agent based on VSCode. When it worked, it was great. Their quota on the Ultra subscription was generous enough that running out of credits was basically impossible. But the backend was exceedingly unreliable in a way that made serious work frustrating. Jobs would cancel partway through all the time and you had to keep retrying by hand. Support was slow and didn’t take much ownership of the issues. I burned a lot of time retrying things that shouldn’t have needed retrying. Eventually I moved on.
Claude Code turned out to be the right tool for me for sustained engineering work. I pay for the $100 Ultra subscription. Early on I was running out of credits because the agent was consuming a lot of context on queries that involved reviewing two codebases and writing a lot of tests. Installing Context Mode , a community plugin that offloads large tool output into a sandbox, stopped that. I kept the cheaper subscription instead of upgrading. The Context Mode plugin is more important than ever now that all AI subscription providers are beginning to dramatically lower quotas .
The other important realization about tooling was that the AI coding ecosystem grew an enormous amount of tooling in a short time: SpecKit , Superpowers , Agent-skills , a new 500-star repo every other week. The ideas inside them are usually great. Installing a framework pushes you into its worldview. You inherit its rigidity, its token overhead, and the opinions of someone who wasn’t building the exact thing you’re building.
For example, I’m a huge believer in Spec Driven Development , but SpecKit didn’t land for me. Every command it adds has a simpler, native alternative you probably already have: CLAUDE.md , plain markdown design docs, GitHub issues, skills, native TodoWrite. Installing it means trading the agency to build your own tighter, project-specific version for someone else’s rulebook.
Then you have bundles of skills (Claude Plugins) like Superpowers and Agent-skills. Superpowers forces worktrees, TDD, and brainstorming on you whether or not those fit the work you’re actually doing. Agent-skills dumps dozens of skills into your config that you’ll never audit. Both have good ideas you should steal but don’t install either as-is. Read them, then write your own skills. With AI in the loop, writing a skill is cheap enough that there’s no excuse not to.
The only third-party tools I’d install now are ones scoped to a single job. Skill Creator helps you author and refine skills without trying to run your workflow. Context Mode keeps context usage in check. Both small, both scoped. Neither is trying to run my workflow for me.
Early on the dev loop looked like pair programming with a lightning-fast partner: read the spec, prompt the agent, review the diff, fix what was wrong, ask for the tests, iterate. I reviewed other people’s branches the same way, pulling them down and walking the diff with a fresh conversation. That worked while I was still learning what the operator wanted to be.
Later we dropped PR reviews entirely because they were too slow for a greenfield project under a tight deadline. Designs got reviewed up front; execution got corrected as it went, through bug audits and refactoring. I probably

[truncated]

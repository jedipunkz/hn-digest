---
source: "https://blog.redplanetlabs.com/2026/05/28/teaching-llms-to-one-shot-complex-backends-at-scale-report-1/"
hn_url: "https://news.ycombinator.com/item?id=48327972"
title: "Teaching LLMs to one-shot complex back ends at scale, report #1"
article_title: "Teaching LLMs to one-shot complex backends at scale, report #1 – Blog"
author: "nathanmarz"
captured_at: "2026-05-30T11:40:22Z"
capture_tool: "hn-digest"
hn_id: 48327972
score: 3
comments: 0
posted_at: "2026-05-29T19:17:06Z"
tags:
  - hacker-news
  - translated
---

# Teaching LLMs to one-shot complex back ends at scale, report #1

- HN: [48327972](https://news.ycombinator.com/item?id=48327972)
- Source: [blog.redplanetlabs.com](https://blog.redplanetlabs.com/2026/05/28/teaching-llms-to-one-shot-complex-backends-at-scale-report-1/)
- Score: 3
- Comments: 0
- Posted: 2026-05-29T19:17:06Z

## Translation

タイトル: LLM に大規模なワンショットの複雑なバックエンドを教える、レポート #1
記事のタイトル: LLM に大規模なワンショットの複雑なバックエンドを教える、レポート #1 – ブログ
説明: AI コーディングにおいて重要なのは LLM だけではありません。 LLM が何をターゲットにしているかは非常に重要です。推論を必要としない単純なターゲットの方が、より良い結果が得られます。 LLM に複雑なバックエンドを生成させる試みは精彩を欠いています。最近の論文「制約の減衰: LLM エージェントの脆弱性」
[切り捨てられた]

記事本文:
LLM に大規模なワンショットの複雑なバックエンドを教える、レポート #1 – ブログ
コンテンツにスキップ
ブログ
LLM に大規模なワンショットの複雑なバックエンドを教える、レポート #1
AI コーディングにおいて重要なのは LLM だけではありません。 LLM が何をターゲットにしているかは非常に重要です。推論を必要としない単純なターゲットの方が、より良い結果が得られます。
LLM に複雑なバックエンドを生成させる試みは精彩を欠いています。最近の論文「 Constraint Decay: The Fragility of LLM Agents in Backend Code Generation 」では、単純な CRUD アプリであっても、現実的な構造上の制約が課されると、完全なテスト スイートでのエンドツーエンドの成功率が 33% に達することが示されています。
従来のバックエンドは、多くの個別のシステムを貼り合わせて構成されており、それぞれが独自のモデルと障害モードを持っています。これらのベンチマークで観察された障害のほとんどは、これらのシステム間の継ぎ目で発生しています。 LLM は 1 つの一貫したシステムについて推論することを求められるのではなく、多くのシステムを調整することが求められます。
これらの方針に沿って、Rama はバックエンドの LLM コーディングを次のレベルに引き上げるのに理想的な立場にあると考えています。 Rama は、典型的なバックエンド スタック (データベース、キュー、ストリーム プロセッサ、アプリケーション ロジック) を 1 つの統合システムにまとめます。現在の LLM がつまずく継ぎ目は、Rama アプリケーションにはほとんど存在しません。水平方向にスケーラブルでフォールト トレラントなバックエンドは、6 個のシステムにわたる接着剤としてではなく、1 つの一貫したプログラムとして表現されます。
過去数か月間、私たちは Rama を基盤として大規模なワンショットの複雑なバックエンドを LLM に教えるプロジェクトに取り組んできました。この投稿の後半でレビューするように、これまでの結果は非常に有望ですが、まだ道はあります。私たちが目指している主要なマイルストーンは、Matrix 仕様全体をワンショットで解決することです。これには、使用できる徹底的なテストのセットも含まれています。

実装を検証します。私たちが生産しようとしているものは次のとおりです。
すべての参照テストに合格する Matrix の生成された実装
LLM がプロジェクトをワンショットで実行した方法の各ステップを示すトランスクリプト
LLM によって自動的に作成および実行されるベンチマークは、高いパフォーマンスと水平スケーラビリティを実証します。
Matrix は、特にスケーラビリティとフォールト トレランスの要件において、現在の LLM が処理できるバックエンドよりも桁違いに難しいため、一発で実現することは大きなマイルストーンとなるでしょう。ただし、最も重要な目標は、これがバックエンドの問題に対処できるようにすることです。私たちが構築しているものが、考えられるすべてのバックエンドにワンショットで対応できるとは期待していません。多くのトレードオフを考慮する必要がある広範なシステム設計において、人間は依然として LLM よりもはるかに優れています。私たちが達成可能だと考えており、このプロジェクトが目標としているのは、人間が高レベルの設計決定を支援し、エージェントが水平方向のスケーラビリティと耐障害性の達成を含む低レベルの決定と実装を処理するワークフローです。 「フォールト トレラント」とは、インフラストラクチャ障害 (ノードの停止など) が発生しても、データの損失、データの重複、ダウンタイムが発生することなくシステムが正常に動作し続け、障害が発生したコンポーネントが復帰すると自動的に回復することを意味します。
私たちの目標が達成できるかどうかはまだわかりませんが、これらの進捗レポートを通じて進捗状況を記録していきます。
私たちはオープンソース化したばかりの rama-ai-learn プロジェクトに取り組んでいます。これは、エージェントが作業を行うために使用するスキル コンテンツとともに、LLM が本番環境の Rama コードをどの程度うまく生成できるかを測定するためのベンチマークおよびハーネスです。
私たちがエージェントに課すそれぞれのタスクは「挑戦」です。チャレンジ ディレクトリ ( example ) には、操作、レイテンシの目標、その他の制約を記載した README.md が含まれています。それもあるよ

■ エージェントが実装する必要があるインターフェース。このディレクトリには、エージェントが認識できないように実行前に暗号化されるプライベート アーティファクトも含まれています。これには、機能の正確性、フォールト トレランス、パフォーマンスをカバーするテストとリファレンス実装が含まれます。エージェントが実装を完了し、独自のテストに合格すると、チャレンジ ランナーは以前に暗号化されたテストを実行して、エージェントが成功したか失敗したかを判断します。
エージェントは、完全な権限を持つ Docker コンテナ内で実行されます。私たちは、思考、ツールの使用、ツールの結果、最終的な応答を含む、すべてのエージェント呼び出しの完全なトランスクリプトをキャプチャします。考えることは特に価値があります。これは、生成されたコードには現れない障害モードを発見する方法です。エージェントがフォールト トレランスのギャップを特定し、考えられる解決策を行ったり来たりした後、「これは複雑になってきています」と言って、まったく対処できなかったようなものです。
Rama には Java および Clojure API があります。私たちは今のところ Clojure に焦点を当てていますが、後で同等の機能を備えた Java バージョンのスキルを作成する予定です。 REPL が主な理由です。長時間実行される REPL セッションでは、エージェントが JVM の起動と依存関係の読み込みに常に料金を支払う代わりに、ミリ秒単位でコードを評価し、結果を検査します。課題が難しくなり、正しい設計に収束するにはエージェントが何度も反復する必要があるため、このギャップがさらに重要になることが予想されます。
LLM のパフォーマンス向上に取り組むには、新しい課題を作成し、エージェントが一貫して合格するまでスキル ファイルを反復処理する必要があります。次に、異なるタイプのアプリケーション、Rama の異なる機能、またはより広範囲のアプリケーションに重点を置く新しい課題に進みます。
3 つの中複雑度の課題は、ほぼすべての実行でワンショットで正しく実行されるようになり、すべて水平スケーラブルでフォールト トレラントです。
銀行口座

ansfer : 各ユーザーの資金を追跡し、ユーザーごとの入金、ユーザー間の送金、残高読み取り、およびインバウンド/アウトバウンド送金履歴の読み取りをサポートします。送金は 1 回限りであり、フォールト トレラントである必要があります。再試行時の二重支出、失敗時に一方に着地してもう一方に着地しない送金、資金不足によるマイナス残高があってはなりません。
時系列: レコードは URL ごとにレンダリング遅延測定を行い、1 分から数年までの任意の分単位の範囲にわたる範囲集計クエリ (カーディナリティ、合計、最小、最大) に答えます。クエリは、5 分であっても 5 年であっても、クエリの長​​さの全範囲にわたって高速である必要があります。
オークション : 販売者が有効期限付きのアイテムをリストし、入札者が入札し、有効期限までに最も高い入札額が落札されるオークションを主催します。販売者のリスト、リストの入札額と現在の最高入札額、およびユーザーの通知の読み取りをサポートします。オークションは有効期限が過ぎると自動的に終了する必要があり、販売者、勝者、敗者通知はそれぞれ正確に 1 回、時系列順に配信される必要があります。
プライベート テストでは、最新の Rama リリースに追加された with-event-hook マクロを使用して、パフォーマンスとフォールト トレランスの特性を検証します。これにより、バランスのとれた計算、基礎となる RocksDB 操作の数と種類、エージェントによって選択されたトポロジ タイプ、強制失敗/再試行などをキャプチャしてアサートできます。
銀行振込チャレンジは 3 つの中で最も簡単です。主なテストは、数百ミリ秒の更新遅延が許容できること、したがってマイクロバッチ トポロジが適切なツールであることをエージェントが認識しているかどうかです。マイクロバッチ トポロジは、すべての計算に対してフォールト トレラントな 1 回限りのセマンティクスを備えています。このチャレンジでは、エージェントが最適な PState (ラーマの同等の PState) を選択していることも検証します。

(データベースへの) 構造、特に転送ログは無制限であるため、サブインデックスが付けられます。
時系列の課題の難しい部分は、複数の粒度でレイテンシー データを事前に集計し、サーバー側の分散クエリ (「クエリ トポロジ」と呼ばれます) を使用して、できるだけ少ないバケットを読み取りながら範囲クエリを効率的に計算することです。エージェントは、読み取られたバケットの数に応じて実行される RocksDB 操作の数を適切に推論し、適切な粒度の数を選択します。また、クライアント ノードとワーカー ノードの間で何回も往復するよりも効率的なクエリ トポロジが適切であることも認識します。
オークション モジュールは最も困難であり、異なるパフォーマンス特性、多態性データ (通知タイプ用)、および時間ベースの動作を持つ複数の機能を備えています。フォールト トレランスを備えたオークション結果の通知を 1 回だけ配信することは、間違いやすいです。エージェントは、リファレンス デザインに近いもの (出品と入札のストリーム トポロジ、有効期限と通知のマイクロバッチ トポロジ) に到達することがあります。驚いたのは、自分が考えてもいなかったデザインが生まれることもあるということです。ストリーム トポロジのみを使用すると、通知はリストではなくマップとして保存され、リストのタイムスタンプ、リスト ID、通知タイプによって決定的にキー設定されます。キーは決定的であるため、有効期限が再試行されると、同じキーが同じ値で書き換えられ、通知の配信をべき等操作にすることで正確に 1 回を実現します。
公開されたベンチマークは、LLM がはるかに単純なバックエンドに苦戦していることを示しています。これらは、スケーラビリティやフォールト トレランス要件のない単一インスタンスのアプリケーションを測定します。私たちの課題は 1 回だけ追加されます

セマンティクス、耐障害性、水平スケーラビリティ、およびテストで積極的に検証されるパフォーマンスの制約。継ぎ目が削除されると、モデルは、アプリケーションに関係する多数のランダムな技術的詳細ではなく、アプリケーション要件に推論予算を費やすことができます。
私たちが開発したスキル ファイルは、これらの課題を一貫してクリアするために何度も反復されました。当初、これらは基本的に Rama ドキュメントを Markdown ファイルに変換したものに過ぎませんでした。最上位のスキルには、Rama をプログラムするために必要なコア情報 (概念、データフロー構文、パス) が含まれていましたが、必要性の低い情報は参照ファイルに置かれていました。
正しいガイダンスがコンテキスト内に読み込まれている場合でも、エージェントは一貫して同じ間違いを犯しました。セッションの開始時に参照を読み取るエージェントは、特定の決定の瞬間にその参照を再度参照することができません。いくつかの例:
バインドされた var を参照するのではなく、
(defmodule FooModule ...)
として直接
Fooモジュール
、次のように構築しようとします
(FooModule.)
PState 内の境界のない場所にはサブインデックスが付けられません
オブジェクト
正確なものの代わりにスキーマに使用されます
特に書き込み前にパーティショナーが欠落し、モジュールが間違った場所に書き込む原因となります。
私たちはスキル作成に関するベスト プラクティスを調査し、段階的なアプローチを使用するようにエージェントに指示するためにスキルを大幅に再構築しました。
暗黙の仕様。プロトコルが暗示しているが明示されていないすべてのエッジケースと不変条件を導き出します。を生成します。
IMPLICIT_SPEC.md
文書。
プラン。コードを記述する前に、デポ、PState スキーマ、およびトポロジを設計します。生産する
計画.md
。
計画の検証。シナリオ (競合状態、失敗/再試行など) のチェックリストに照らして計画を検討します。生産する
PLAN_VALIDATION.md
合否判定付き。
埋め込む。書き込みます

モジュールのソースは計画に厳密に従っています。
実装の検証。よくある間違いのチェックリストと照らし合わせて実装を確認します。生産する
IMPLEMENTATION_VALIDATION.md
。
テスト。すべてのプロトコル メソッドと暗黙的な仕様のすべてのエッジ ケースをカバーする徹底的なテストを作成します。
テストの検証。よくある間違いのチェックリストと照らし合わせてテスト コードを確認します。生産する
TEST_VALIDATION.md
。
テストを実行します。テストスイートを実行します。
これを効果的にするのは 3 つの点です。すべてのフェーズでアーティファクトを要求すると、エージェントが黙ってステップをスキップすることがなくなります。検証フェーズは明示的なチェックリストとして記述されており、LLM はエントリを忘れることなくチェックリストを項目ごとに実行する点で信頼性が高くなります。また、フェーズ分割により、エージェントは一度に小さな作業に集中できるようになります。
いくつかの言語パターンが大きな違いを生むことが判明しました。理由のある否定的な制約は、エージェントが誤った道に進むのを防ぐという点で、肯定的な提案よりもはるかに効果的です。理由を含めることは重要です。そうしないと、エージェントがそれを無視するためのランダムな正当化を考え出すからです (例: 「このフェーズではトポロジ コードを書かないでください。計画のない実装では、間違った PState スキーマ、間違ったパーティション アライメント、および不要なディスク I/O が生成されます。」)。明示的なデフォルトの選択肢を指定し、

[切り捨てられた]

## Original Extract

The LLM is not all that matters in AI coding. What the LLM is targeting matters a great deal. A simpler target that requires less reasoning will produce better results. Attempts to get LLMs to produce complex backends have been lackluster. A recent paper, Constraint Decay: The Fragility of LLM Agent
[truncated]

Teaching LLMs to one-shot complex backends at scale, report #1 – Blog
Skip to content
Blog
Teaching LLMs to one-shot complex backends at scale, report #1
The LLM is not all that matters in AI coding. What the LLM is targeting matters a great deal. A simpler target that requires less reasoning will produce better results.
Attempts to get LLMs to produce complex backends have been lackluster. A recent paper, Constraint Decay: The Fragility of LLM Agents in Backend Code Generation , shows that even on a simple CRUD app, end-to-end success on the full test suite tops out at 33% once realistic structural constraints are imposed.
Conventional backends are made of many separate systems glued together, each with its own model and failure modes. Most of the failures observed in these benchmarks show up at the seams between these systems. The LLM is not asked to reason about one coherent system – it’s asked to coordinate across many.
Along those lines, we believe Rama is ideally positioned to take LLM coding to the next level for backends. Rama collapses the typical backend stack (databases, queues, stream processors, application logic) into one integrated system. The seams that current LLMs trip over largely don’t exist in a Rama application. A horizontally scalable, fault-tolerant backend is expressed as one coherent program rather than as glue across half a dozen systems.
In the past few months we’ve been working on a project to teach LLMs to one-shot complex backends at scale with Rama as the substrate. Our results so far are very promising, as I’ll review later in this post, but we have a ways to go. The major milestone we’re working towards is one-shotting the entire Matrix spec , which also has a thorough set of tests available that can be used to verify an implementation. What we’re looking to produce is:
A generated implementation of Matrix that passes all the reference tests
Transcript showing every step of how the LLM one-shotted the project
Benchmarks automatically written and executed by the LLM that demonstrate high performance and horizontal scalability
Matrix is orders of magnitude more difficult than the backends current LLMs can handle, particularly with these scalability and fault-tolerance requirements, so one-shotting it will be a huge milestone. However, the overarching goal is for this to work on any backend problem. We don’t expect what we’re building to one-shot every possible backend. Humans remain vastly better than LLMs at broad systems design where many tradeoffs must be considered. What we think is achievable, and what this project is targeting, is a workflow where humans assist with high-level design decisions and the agent handles lower-level decisions and implementation, including achieving horizontal scalability and fault-tolerance. By “fault tolerant” we mean the system continues operating correctly through infrastructure failures (e.g. node deaths) without data loss, data duplication, or downtime, and recovers automatically when failed components return.
Whether our goal is possible remains to be seen, but I’ll be documenting our progress as we go via these progress reports.
We work through the rama-ai-learn project, which we just open-sourced. It’s a benchmark and harness for measuring how well LLMs can produce production Rama code, along with the skill content the agent uses to do the work.
Each task we throw at an agent is a “challenge.” A challenge directory ( example ) contains a README.md stating operations, latency targets, and other constraints. It also has an interface the agent must implement. The directory also contains private artifacts that are encrypted before runs so the agent can’t see them: tests covering functional correctness, fault-tolerance, and performance, and a reference implementation. After an agent finishes its implementation and passes its own tests, the challenge runner runs the formerly encrypted tests to determine whether the agent succeeded or failed.
Agents are run inside a Docker container with full permissions. We capture every agent invocation’s full transcript, including thinking, tool uses, tool results, and the final response. Thinking is particularly valuable. It’s how we discover failure modes that don’t show up in the produced code, like an agent identifying a fault-tolerance gap, going back and forth on possible solutions, and then saying “this is getting complicated” and failing to address it at all.
Rama has Java and Clojure APIs. We’re focused on Clojure for now but will produce an equally capable Java version of the skills later. The REPL is the main reason, as with a long-running REPL session, the agent evaluates code and inspects results in milliseconds instead of constantly paying for JVM startup and dependency loading. We expect this gap to matter more as challenges get harder and converging on a correct design takes the agent many iterations.
Working on improving LLM performance involves making a new challenge and then iterating on the skill files until the agent passes it consistently. Then we move on to a new challenge that stresses a different type of application, a different capability of Rama, or an application of greater scope.
Three medium-complexity challenges now one-shot correctly on almost every run, all horizontally scalable and fault-tolerant:
Bank transfer : tracks funds for each user and supports deposits, transfers between users, balance reads, and inbound/outbound transfer history reads per user. Transfers must be exactly-once and fault-tolerant: no double-spends under retry, no transfer that lands on one side but not the other under failure, and no negative balances from insufficient funds.
Time series : records render latency measurements per URL and answers range-aggregate queries (cardinality, total, min, max) over any minute-bucket range from a single minute to many years. Queries must be fast across the full range of query lengths, whether five minutes or five years.
Auction : hosts auctions where sellers list items with expiration times, bidders place bids, and the highest bid at expiration wins. Supports reading a seller’s listings, a listing’s bids and current top bid, and a user’s notifications. Auctions must end automatically when their expiration time passes, and seller / winner / losing-bidder notifications must each be delivered exactly once and in chronological order.
The private tests verify performance and fault-tolerance characteristics using the with-event-hook macro that we added in the latest Rama release. With this we can capture and assert on computation being balanced, the number and types of underlying RocksDB operations, topology types chosen by the agent, force failures/retries, and more.
The bank transfer challenge is the easiest of the three. The main test is whether the agent recognizes that update latency in the hundreds of milliseconds is acceptable, and therefore that a microbatch topology is the right tool. Microbatch topologies have fault-tolerant exactly-once semantics for all computation. The challenge also verifies the agent chooses the optimal PState (Rama’s equivalent to databases) structures, especially that transfer logs are subindexed since they’re unbounded.
The hard part of the time series challenge is pre-aggregating the latency data at multiple granularities and then using a server side distributed query (called a “query topology”) to efficiently compute the range query while reading as few buckets as possible. The agent does a great job of reasoning through how many RocksDB operations will be done depending on the number of buckets read and then choosing the appropriate number of granularities. It also recognizes a query topology is appropriate since that’s more efficient than many roundtrips between the client and worker nodes.
The auction module is the hardest, with multiple features with differing performance characteristics, polymorphic data (for notification types), and time-based behavior. Getting notifications of auction results to have exactly-once delivery with fault-tolerance is easy to get wrong. The agent sometimes lands on something close to the reference design – a stream topology for listings and bids and a microbatch topology for expirations and notifications. What surprised me is it also sometimes produces a design I had never considered. Using just a stream topology, notifications are stored as a map rather than a list, keyed deterministically by the listing’s timestamp, the listing ID, and the notification type. Because the key is deterministic, a retried expiration rewrites the same keys with the same values, achieving exactly-once by making delivery of the notification an idempotent operation.
The published benchmarks show LLMs struggling with far simpler backends. They measure single-instance applications with no scalability or fault-tolerance requirements. Our challenges add exactly-once semantics, fault tolerance, horizontal scalability, and performance constraints that the tests actively verify. With seams removed, the model can spend its reasoning budget on the application requirements rather than a host of random technical details tangential to the application.
The skill files we’ve developed have gone through many iterations in order to pass these challenges consistently. At first, they were basically just the Rama documentation translated to Markdown files. The top-level skill had core information needed to program Rama (concepts, dataflow syntax, paths), while less-needed information was put in reference files.
The agent consistently made the same mistakes even when the correct guidance was loaded in context. An agent that reads a reference at the start of a session would fail to revisit it at the moment of a specific decision. Some examples:
Rather than refer to the var bound by
(defmodule FooModule ...)
directly as
FooModule
, it would try to construct it as
(FooModule.)
Unbounded locations in a PState would not be subindexed
Object
would be used for schemas instead of something precise
Partitioners would be missing, especially before writes, causing the module to write to the wrong locations
We researched best practices on making skills and then did a major restructuring of the skills to instruct the agent to use a phased approach:
Implicit spec. Derive every edge case and invariant the protocol implies but doesn’t state. Produces an
IMPLICIT_SPEC.md
document.
Plan. Design the depots, PState schemas, and topologies before writing any code. Produces
PLAN.md
.
Plan validation. Review the plan against a checklist of scenarios (e.g. race conditions, failures/retries). Produces
PLAN_VALIDATION.md
with a pass or fail verdict.
Implement. Write the module source, adhering strictly to the plan.
Implementation validation. Review the implementation against a checklist of common mistakes. Produces
IMPLEMENTATION_VALIDATION.md
.
Tests. Write thorough tests covering every protocol method and every edge case from the implicit spec.
Test validation. Review the test code against a checklist of common mistakes. Produces
TEST_VALIDATION.md
.
Run tests. Run the test suite.
Three things make this effective. Requiring an artifact at every phase prevents the agent from silently skipping a step. The validation phases are written as explicit checklists, and LLMs are reliable at walking a checklist item-by-item without forgetting entries. And the phase split lets the agent focus on smaller pieces of work at a time.
A few language patterns turned out to make a big difference. Negative constraints with a reason work much better than positive suggestions at preventing the agent from going down wrong paths. Including a reason is critical as otherwise the agent comes up with some random justification to ignore it (e.g. “Do NOT write any topology code in this phase — implementations without plans produce wrong PState schemas, incorrect partition alignment, and unnecessary disk I/O”). We specify explicit default choices and requi

[truncated]

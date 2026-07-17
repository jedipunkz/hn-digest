---
source: "https://wolbarg.com/blog/shared-memory-vs-isolated-memory"
hn_url: "https://news.ycombinator.com/item?id=48949146"
title: "Shared Memory vs. Isolated Memory for AI Agents"
article_title: "Shared Memory vs Isolated Memory Architecture for AI Agents · Wolbarg"
author: "atharvmunde"
captured_at: "2026-07-17T17:00:45Z"
capture_tool: "hn-digest"
hn_id: 48949146
score: 1
comments: 0
posted_at: "2026-07-17T16:20:33Z"
tags:
  - hacker-news
  - translated
---

# Shared Memory vs. Isolated Memory for AI Agents

- HN: [48949146](https://news.ycombinator.com/item?id=48949146)
- Source: [wolbarg.com](https://wolbarg.com/blog/shared-memory-vs-isolated-memory)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:20:33Z

## Translation

タイトル: AI エージェントの共有メモリと分離メモリ
記事のタイトル: AI エージェントの共有メモリ アーキテクチャと分離メモリ アーキテクチャ · Wolbarg
説明: セマンティック リコール、同時実行性、権限、イベント ログ、チェックポイントなど、AI エージェントの共有メモリと分離メモリの実際的な比較。

記事本文:
AI エージェントの共有メモリと分離メモリ アーキテクチャ · Wolbarg WOLBΛRG WOLBΛRG Search ⌘ K ベンチマーク AI エージェントの共有メモリと分離メモリ アーキテクチャ AI エージェントのローカル メモリには SQLite で十分 AI エージェントの共有メモリと分離メモリ アーキテクチャ AI エージェントの共有メモリと分離メモリ アーキテクチャ
AI エージェントの共有メモリと分離メモリの実際的な比較 (セマンティック リコール、同時実行性、権限、イベント ログ、チェックポイントなど)。
Wolbarg の構築 — AI エージェント向けのローカル ファーストのセマンティック メモリ。
マルチエージェントメモリアーキテクチャ
AI エージェントに関するほとんどの議論は、モデル、プロンプト、ツールから始まります。ただし、運用システムでは、エージェントが共同作業できるか、障害から回復できるか、6 か月後も理解できる状態を維持できるかどうかは、メモリ アーキテクチャによって決まります。
中心的な決定は簡単に述べることができます。各エージェントが分離されたメモリを所有すべきか、それとも複数のエージェントが共有の意味論的メモリを通じて動作するべきかです。
その答えは、ストレージ コスト、取得動作、アクセス制御、同時実行性、およびデバッグを変更します。また、「エージェント」の意味も変わります。メモリが分離されているため、エージェントは自己完結型のワーカーとなります。共有メモリを使用すると、エージェントはより大きな知識システムの参加者になります。
この記事では、両方のデザインを比較し、どちらかを選択するための実際的な基準を示します。
AI エージェントの分離メモリとは何ですか?
分離されたメモリ アーキテクチャでは、すべてのエージェントが個別のデータベース、ベクトル インデックス、または会話ストアを所有します。エージェント A は、アプリケーションが明示的にコピーまたは同期しない限り、エージェント B の観察を呼び出すことはできません。
その境界がコードと一致するため、これは自然なデフォルトです。パーソナル アシスタントは 1 つの SQLite ファイルに書き込むことができます。使い捨ての抽出エージェントはメモリ内インデックスを保持できます。テストはリセットできる

他のエージェントに触れることなく、あるエージェントを操作できます。
運用上の利点は次のとおりです。
シンプルな実装: 所有権とライフサイクルが明らかです。
簡単なデバッグ: 驚くべきリコール結果には、考えられる原因が 1 つあります。
強力な分離: あるエージェントが別のエージェントのシークレットを誤って読み取ることはできません。
エージェント間の書き込み競合がない: 独立したストアには共有ロックやトランザクションが必要ありません。
デメリットは、エージェントが同じ知識を必要とする場合に現れます。
事実、要約、埋め込みが複製されます。
コラボレーションはメッセージ パッシングまたはデータベース同期になります。
更新により、同じファクトの競合するコピーが作成される可能性があります。
新しいエージェントは、歴史的なコンテキストなしで開始されます。
ストレージと組み込みのコストは、固有の知識ではなくエージェントの数に応じて増加します。
5 人のエージェントが同じ 20,000 トークンのプロジェクト概要を埋め込む場合、システムは 5 つの取り込みパイプラインの料金を支払い、意味的に同等の 5 つのインデックスを保存します。難しいのはディスク容量ではありません。どのコピーが最新であるかを決定することです。
分離は一貫性戦略です
分離メモリは、共有状態を回避することで分散一貫性を回避します。これは、単に実装が便利というだけでなく、エージェントが真に独立している場合に価値があります。
AI エージェントの共有メモリとは何ですか?
共有メモリ アーキテクチャは、エージェントと永続ストレージの間に共通のセマンティック メモリ層を配置します。エージェントは、同じ永続的なコンテキストを通じて観察を書き込み、関連する知識を呼び出し、状態を更新し、調整することができます。
「共有」とは、「すべてのエージェントがすべてのメモリを受け取る」という意味ではありません。これは、システムに 1 つの管理されたメモリ プレーンがあることを意味します。取得では、関連する記憶をランク付けする前に、組織、ワークスペース、エージェント、タスク、機密性、および時間のフィルターが適用されます。
最小限の、フレームワークに依存しないインターフェイスは次のようになります。
タイプ SharedMemory = {
覚えておいてください (入力: {
名前空間:文字列;
テキスト：ストリ

NG;
ソースエージェント:文字列;
メタデータ ?: レコード < string 、不明 >;
}) : プロミス <{ id : string }>;
リコール (入力: {
名前空間:文字列;
クエリ: 文字列;
topK ?: 数値 ;
リクエスタ : 文字列 ;
}) : Promise < 配列 <{ テキスト : 文字列 ;スコア: 数値 }>>;
};
共有レイヤーは、埋め込み、インデックス付け、フィルタリング、バージョン管理、および永続性を所有します。何を覚えておくと十分役立つかについては、エージェントが独自に決定します。
コンテキストをコピーせずにコラボレーションを向上
調査エージェントが、API が v1/payments を廃止したことを知ったとします。ソース URL とタイムスタンプを含む検出結果を一度保存​​します。プランナーが「どの統合リスクが移行に影響を及ぼしますか?」と尋ねました。すぐに思い出すことができます。
記憶を待ちます。覚えておいてください({
名前空間: "migration-42" 、
テキスト: 「v1/payments エンドポイントは 2026 年 10 月 1 日以降廃止されます。」 、
ソースエージェント: "研究者" 、
メタデータ: { ソース: "vendor-docs" 、信頼度: 0.96 }、
});
const リスク = メモリを待ちます。思い出してください({
名前空間: "migration-42" 、
クエリ: "配信をブロックする可能性がある統合の非推奨" 、
依頼者: 「プランナー」 、
トップK: 5 、
});
プロンプトのトランスクリプトはコピーされません。オーケストレーターは、将来どのエージェントが結果を必要とするかを予測する必要はありません。意味的想起は、たとえ表現が異なっていても、後の質問を前の観察に結び付けます。
共有メモリは、出所、バージョン、有効日を含む 1 つの正規レコードを維持できます。更新は、別の真実を黙って作成するのではなく、置き換えられるバージョンを記録または追加する変更を加えます。
これでは意見の相違がなくなるわけではありません。これにより、不一致が表現可能になります。2 つの記憶が異なる情報源、信頼度スコア、発効日を引用することができ、検索により両方が明らかになります。
ストレージと計算の重複が少ない
独自の知識が一度埋め込まれます。概要は 1 回生成されます。大きなドキュメントは一度チャンク化されます。 10 人のエージェントがより多くの領域を生み出す可能性がある

ds を使用しますが、コーパスの 10 部のコピーは必要ありません。
通常、テキストが占めるバイト数は埋め込みコスト、インデックス作成時間、再取り込み作業によって占められるため、経済的な違いが重要になります。
新しいレビュー担当者はワークフローに参加し、アーキテクチャ上の決定、失敗したアプローチ、ユーザーの好み、未解決のリスクを即座に取得できます。その役割のプロンプトは、どのように動作するかを説明します。共有長期記憶は、システムがすでに知っていることを説明します。
プロセスが停止します。モデルが変わります。エージェントが交代します。永続的な共有コンテキストにより、システムの知識が 1 つのランタイムから分離され、デプロイメントまたはクラッシュ後に作業を再開できます。
共有コンテキストは、より大きなプロンプトではありません。これは取得境界です。広範囲に保持し、積極的にフィルタリングし、現在の決定に関連する少数のメモリのみを挿入します。
1 人のエージェントによって書かれた低品質の推論は、他のすべてのエージェントに影響を与える可能性があります。書き込みをカジュアルなログではなく、データの取り込みとして扱います。
テナントとプロジェクトの分離、タスクとソースのメタデータ フィルター、信頼しきい値、有効期限ポリシー、および「観察」から「検証された事実」への明示的な昇格には名前空間を使用します。消費者が文脈のないテキストを信頼するのではなく、記憶を判断できるように出所を保存します。
同時実行性と競合する書き込み
2 人のエージェントが同じプランを同時に更新できます。トランザクションは部分的な書き込みを防止しますが、セマンティックの競合は解決しません。
変更可能なレコードにはオプティミスティック バージョン チェックを使用し、再試行には冪等キーを使用し、履歴が重要な場合は追加専用のファクトを使用します。 WAL モードの SQLite は、多数のリーダーとシリアル化された書き込みを備えた 1 台のマシンで適切に機能します。 PostgreSQL は、ネットワークに接続された多くのライターに適しています。 1 つのデータベースがすべての負荷に適合すると仮定するのではなく、測定された方法論については、Wolbarg のバックエンド ベンチマークを参照してください。
すべてのエージェントが給与データを思い出す必要があるわけではありませんが、

プライベートな会話、または本番環境の認証情報。アクセス制御は、取得結果がモデルに到達する前に、メモリ サービス内で強制される必要があります。
堅牢なポリシーでは、テナント、名前空間、役割、ソース、機密性、および操作が考慮されます。読み取り権限と書き込み権限は別々にする必要があります。メタデータ フィルターは便利な取得コントロールですが、承認の代わりにはなりません。
記憶状態は「現在何がわかっているか?」に答えます。デバッグでは、「誰がこれを書いたのか、どのクエリがそれを取得したのか、そしてどのランキング ステージがその結果を生み出したのか?」と尋ねます。
イベント ログをセマンティック メモリとは別のデータベースに保存します。イベント データは大量に追加され、時間順に並べられ、監査または再生のために保持されます。メモリ データは変更可能で、取得に最適化されており、圧縮または削除される可能性があります。分離することで、テレメトリの増加によるリコールの低下を防ぎ、再帰的な「ログが思い出になる」ことを回避し、オペレータが異なる保存ポリシーとアクセス ポリシーを適用できるようになります。
トランザクションは 1 つの操作を保護します。エージェントが検証された事実を誤った概要に置き換えるなど、有効ではあるが有害な一連の操作から保護することはできません。
チェックポイントは、危険な移行、圧縮、または自律実行の前に回復可能な状態をキャプチャします。有用なシステムは、名前付きスナップショット、整合性検証、ロールバック、および復元された状態をその状態を生成したアクションに結び付けるイベント証跡をサポートしています。
コラボレーションによって価値よりもポリシーと一貫性の作業が増える場合は、分離されたメモリを選択します。
典型的な例には、無関係なユーザーのための独立したアシスタント、個人用チャットボット、使い捨ての評価エージェント、臨時労働者、機密性の高いタスクを処理するサンドボックス エージェントなどがあります。
共有メモリは、システムの出力がエージェントの境界を越えて蓄積された知識に依存する場合に適しています。
コーディングエージェントはリポジトリ規約を共有できます

ns と失敗した修正。研究機関は共通の証拠ベースを構築できます。カスタマー サポート エージェントは、無制限の顧客データを公開することなく、解決済みのケースを再利用できます。長時間実行されるワークフローは、プロセスが再起動されても存続できます。マルチエージェント オーケストレーターと自律型組織は、従業員が出入りしても意思決定を維持できます。
有用なテストは次のとおりです。エージェント A が失踪した場合でも、エージェント B は A が学んだことを発見する必要がありますか? 「はい」の場合、その知識はエージェント A の外部に属します。
Wolbarg が共有メモリを実装する方法
Wolbarg は、TypeScript アプリケーションにローカルファーストの共有セマンティック メモリを提供します。エージェントは小規模な Remember and Recall API を使用し、アプリケーションはストレージおよび取得プロバイダーを選択します。
SQLite コネクタは、単一マシン システムとローカル エージェント チームに適合します。 PostgreSQL コネクタは、ネットワーク化された展開とより優れた同時書き込みをサポートします。どちらも交換可能なプロバイダーの背後にあるため、メモリ モデルは 1 つのベクトル エンジンに結び付けられるのではなく、データベースに依存しません。アーキテクチャ ガイドでは、プロバイダーの境界について説明します。
セマンティック リコールはクエリを埋め込み、保存された知識を検索し、ベクトルの類似性をキーワード シグナルやメタデータ フィルターと組み合わせることができます。組織、エージェント ID、タグ、メタデータは、テナント、プロジェクト、ソース、ロールを分離するための実用的な名前空間として機能します。
Wolbarg v0.3 では、独立したイベント ログとチェックポイントが追加されています。イベント データベースはメモリ ストレージから分離されたままですが、チェックポイントは回復可能なローカル ワークフローのスナップショットとロールバックを提供します。これらの機能は可観測性と回復性に対処します。これらは、アプリケーション レベルの認可ポリシーや競合ポリシーに代わるものではありません。イベント モデルについては、「可観測性とスタジオ」を参照してください。
設計目標は、すべてのシステムですべてのメモリを共有することではありません。エンジニアがカップルなしで共有永続境界を選択できるようにするためです。

その選択を特定のデータベースまたはエージェント フレームワークに反映します。
どちらのアーキテクチャも普遍的に優れているわけではありません。
エージェントが独立している場合、一時的な場合、またはセキュリティ境界が物理的に単純である必要がある場合は、分離メモリを使用します。エージェントが共同作業するときに共有メモリを使用し、高価な知識を再利用し、再起動後も存続し、時間をかけてコンテキストを構築します。
成熟した設計はハイブリッド型であることが多く、エージェントごとのプライベート スクラッチ メモリ、検証済みの知識用の共有プロジェクト メモリ、監査用の別個のイベント ストアがあります。プロモーション ルールは、選択された観測をプライベート スコープから共有スコープに移動します。
正しい選択は、どのモデルがエージェントに電力を供給するかではなく、所有権、一貫性、有効期間、および信頼の境界によって決まります。
ローカル AI エージェントのメモリには SQLite で十分
ベンチマークの方法論と結果
Wolbarg の構築 — AI エージェント向けのローカル ファーストのセマンティック メモリ。
ローカル AI エージェントのメモリには SQLite で十分
ほとんどのローカル マルチエージェント ワークフローでセマンティック メモリに PostgreSQL が必要ない理由と、実際に必要な場合。

## Original Extract

A practical comparison of shared and isolated memory for AI agents, including semantic recall, concurrency, permissions, event logs, and checkpoints.

Shared Memory vs Isolated Memory Architecture for AI Agents · Wolbarg WOLBΛRG WOLBΛRG Search ⌘ K Benchmarks Shared Memory vs Isolated Memory Architecture for AI Agents SQLite Is Enough for Local AI Agent Memory Shared Memory vs Isolated Memory Architecture for AI Agents Shared Memory vs Isolated Memory Architecture for AI Agents
A practical comparison of shared and isolated memory for AI agents, including semantic recall, concurrency, permissions, event logs, and checkpoints.
Building Wolbarg — local-first semantic memory for AI agents.
multi-agent memory architecture
Most discussions about AI agents begin with models, prompts, and tools. In production systems, however, memory architecture often determines whether agents can collaborate, recover from failures, and remain understandable six months later.
The central decision is simple to state: should each agent own an isolated memory, or should multiple agents work through a shared semantic memory?
The answer changes storage cost, retrieval behavior, access control, concurrency, and debugging. It also changes what an “agent” means. With isolated memory, an agent is a self-contained worker. With shared memory, an agent is a participant in a larger knowledge system.
This article compares both designs and gives practical criteria for choosing between them.
What Is Isolated Memory for AI Agents?
In an isolated memory architecture, every agent owns a separate database, vector index, or conversation store. Agent A cannot recall Agent B's observations unless the application explicitly copies or synchronizes them.
This is the natural default because its boundaries match the code. A personal assistant can write to one SQLite file. A disposable extraction agent can keep an in-memory index. Tests can reset one agent without touching another.
The advantages are operational:
Simple implementation: ownership and lifecycle are obvious.
Easy debugging: a surprising recall result has one possible source.
Strong isolation: one agent cannot accidentally read another agent's secrets.
No cross-agent write contention: independent stores need no shared lock or transaction.
The disadvantages appear when agents need the same knowledge:
Facts, summaries, and embeddings are duplicated.
Collaboration becomes message passing or database synchronization.
Updates can produce conflicting copies of the same fact.
A new agent starts without historical context.
Storage and embedding costs grow with agent count rather than unique knowledge.
If five agents embed the same 20,000-token project brief, the system pays for five ingestion pipelines and stores five semantically equivalent indexes. The difficult part is not disk space; it is deciding which copy is current.
Isolation is a consistency strategy
Isolated memory avoids distributed consistency by avoiding shared state. That is valuable when agents are genuinely independent, not merely convenient to implement.
What Is Shared Memory for AI Agents?
A shared memory architecture places a common semantic memory layer between agents and persistent storage. Agents can write observations, recall related knowledge, update state, and coordinate through the same durable context.
“Shared” should not mean “every agent receives every memory.” It means the system has one governed memory plane. Retrieval still applies organization, workspace, agent, task, sensitivity, and time filters before ranking relevant memories.
A minimal, framework-neutral interface might look like this:
type SharedMemory = {
remember ( input : {
namespace : string ;
text : string ;
sourceAgent : string ;
metadata ?: Record < string , unknown >;
}) : Promise <{ id : string }>;
recall ( input : {
namespace : string ;
query : string ;
topK ?: number ;
requester : string ;
}) : Promise < Array <{ text : string ; score : number }>>;
};
The shared layer owns embedding, indexing, filtering, versioning, and persistence. Agents own decisions about what is useful enough to remember.
Better Collaboration Without Copying Context
Suppose a research agent learns that an API deprecated v1/payments . It stores the finding once, including source URL and timestamp. A planner asking “which integration risks affect migration?” can recall it immediately.
await memory. remember ({
namespace: "migration-42" ,
text: "The v1/payments endpoint is deprecated after 2026-10-01." ,
sourceAgent: "researcher" ,
metadata: { source: "vendor-docs" , confidence: 0.96 },
});
const risks = await memory. recall ({
namespace: "migration-42" ,
query: "integration deprecations that can block delivery" ,
requester: "planner" ,
topK: 5 ,
});
No prompt transcript is copied. No orchestrator has to predict which future agent needs the result. Semantic recall connects a later question to an earlier observation even when the wording differs.
Shared memory can maintain one canonical record with provenance, version, and validity dates. An update changes that record or appends a superseding version instead of silently creating another truth.
This does not eliminate disagreement. It makes disagreement representable: two memories can cite different sources, confidence scores, or effective dates, and retrieval can expose both.
Less Duplicate Storage and Computation
Unique knowledge is embedded once. Summaries are generated once. Large documents are chunked once. Ten agents may create more reads, but they do not require ten copies of the corpus.
The economic difference matters because embedding cost, indexing time, and re-ingestion work usually dominate the bytes occupied by text.
A new reviewer agent can join a workflow and immediately retrieve architectural decisions, failed approaches, user preferences, and open risks. Its role prompt explains how to behave ; shared long-term memory explains what the system already knows .
Processes stop. Models change. Agents are replaced. Durable shared context separates system knowledge from any one runtime, allowing work to resume after a deployment or crash.
Shared context is not a larger prompt. It is a retrieval boundary: persist broadly, filter aggressively, and inject only the few memories relevant to the current decision.
A low-quality inference written by one agent can influence every other agent. Treat writes as data ingestion, not casual logging.
Use namespaces for tenant and project isolation, metadata filters for task and source, confidence thresholds, expiration policies, and explicit promotion from “observation” to “verified fact.” Store provenance so consumers can judge a memory rather than trusting text without context.
Concurrency and Conflicting Writes
Two agents may update the same plan simultaneously. Transactions prevent partial writes, but they do not resolve semantic conflict.
Use optimistic version checks for mutable records, idempotency keys for retries, and append-only facts when history matters. SQLite in WAL mode works well for one machine with many readers and serialized writes; PostgreSQL is a better fit for many networked writers. See Wolbarg's backend benchmarks for measured methodology rather than assuming one database fits every load.
Not every agent should recall payroll data, private conversations, or production credentials. Access control must be enforced inside the memory service before retrieval results reach the model.
A robust policy considers tenant, namespace, role, source, sensitivity, and operation. Read and write permissions should be separate. Metadata filters are useful retrieval controls, but they are not a substitute for authorization.
Memory state answers “what is known now?” Debugging asks “who wrote this, what query retrieved it, and which ranking stages produced that result?”
Keep event logs in a separate database from semantic memory. Event data is append-heavy, time-ordered, and retained for audit or replay. Memory data is mutable, retrieval-optimized, and may be compacted or deleted. Separation prevents telemetry growth from degrading recall, avoids recursive “logs becoming memories,” and lets operators apply different retention and access policies.
A transaction protects one operation. It does not protect against a valid but harmful sequence of operations, such as an agent replacing verified facts with an incorrect summary.
Checkpoints capture a recoverable state before risky migrations, compression, or autonomous runs. Useful systems support named snapshots, integrity verification, rollback, and an event trail connecting the restored state to the actions that produced it.
Choose isolated memory when collaboration would create more policy and consistency work than value.
Typical examples include independent assistants for unrelated users, personal chatbots, disposable evaluation agents, temporary workers, and sandboxed agents handling sensitive tasks.
Shared memory is appropriate when the system's output depends on accumulated knowledge crossing agent boundaries.
Coding agents can share repository conventions and failed fixes. Research agents can build a common evidence base. Customer-support agents can reuse resolved cases without exposing unrestricted customer data. Long-running workflows can survive process restarts. Multi-agent orchestrators and autonomous organizations can preserve decisions as workers come and go.
A useful test is: if Agent A disappears, must Agent B still discover what A learned? If yes, the knowledge belongs outside Agent A.
How Wolbarg Implements Shared Memory
Wolbarg provides a local-first shared semantic memory for TypeScript applications. Agents use a small Remember and Recall API , while the application chooses the storage and retrieval providers.
The SQLite connector fits single-machine systems and local agent teams. The PostgreSQL connector supports networked deployments and greater write concurrency. Both sit behind replaceable providers, so the memory model is database-agnostic rather than tied to one vector engine. The architecture guide describes the provider boundaries.
Semantic recall embeds the query, searches stored knowledge, and can combine vector similarity with keyword signals and metadata filters. Organizations, agent identities, tags, and metadata act as practical namespaces for separating tenants, projects, sources, and roles.
Wolbarg v0.3 adds independent event logging and checkpoints. The event database remains separate from memory storage, while checkpoints provide snapshots and rollback for recoverable local workflows. These features address observability and recovery; they do not replace application-level authorization or conflict policy. See Observability and Studio for the event model.
The design goal is not to make every system share all memory. It is to let engineers choose a shared persistence boundary without coupling that choice to a specific database or agent framework.
Neither architecture is universally better.
Use isolated memory when agents are independent, temporary, or security boundaries must be physically simple. Use shared memory when agents collaborate, reuse expensive knowledge, survive restarts, and build context over time.
The mature design is often hybrid: private scratch memory per agent, shared project memory for verified knowledge, and a separate event store for audit. Promotion rules move selected observations from private to shared scope.
The correct choice depends on ownership, consistency, lifetime, and trust boundaries—not on which model powers the agent.
SQLite Is Enough for Local AI Agent Memory
Benchmark methodology and results
Building Wolbarg — local-first semantic memory for AI agents.
SQLite Is Enough for Local AI Agent Memory
Why most local multi-agent workflows don't need PostgreSQL for semantic memory — and when they actually do.

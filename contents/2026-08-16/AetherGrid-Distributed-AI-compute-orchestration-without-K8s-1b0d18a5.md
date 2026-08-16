---
source: "https://github.com/wycliffRotich-dev/aethergrid"
hn_url: "https://news.ycombinator.com/item?id=49318822"
title: "AetherGrid – Distributed AI compute orchestration without K8s"
article_title: "GitHub - wycliffRotich-dev/aethergrid: A decentralized compute orchestration platform powered by strict Domain-Driven Design (DDD). Decoupled domain core, event-driven state reconciliation, and zero-overhead cluster management. · GitHub"
author: "WycliffeRotich"
captured_at: "2026-08-16T11:11:52Z"
capture_tool: "hn-digest"
hn_id: 49318822
score: 1
comments: 0
posted_at: "2026-08-16T10:47:04Z"
tags:
  - hacker-news
  - translated
---

# AetherGrid – Distributed AI compute orchestration without K8s

- HN: [49318822](https://news.ycombinator.com/item?id=49318822)
- Source: [github.com](https://github.com/wycliffRotich-dev/aethergrid)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T10:47:04Z

## Translation

タイトル: AetherGrid – K8 を使用しない分散 AI コンピューティング オーケストレーション
記事のタイトル: GitHub - wycliffRotich-dev/aethergrid: 厳格なドメイン駆動設計 (DDD) を活用した分散型コンピューティング オーケストレーション プラットフォーム。分離されたドメイン コア、イベント ドリブンの状態調整、およびオーバーヘッドのないクラスター管理。 · GitHub
説明: 厳密なドメイン駆動設計 (DDD) を活用した分散型コンピューティング オーケストレーション プラットフォーム。分離されたドメイン コア、イベント ドリブンの状態調整、およびオーバーヘッドのないクラスター管理。 - wycliffRotich-dev/aethergrid

記事本文:
GitHub - wycliffRotich-dev/aethergrid: 厳密なドメイン駆動設計 (DDD) を活用した分散型コンピューティング オーケストレーション プラットフォーム。分離されたドメイン コア、イベント ドリブンの状態調整、およびオーバーヘッドのないクラスター管理。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ウィクリフロティッチ開発者
/
エーテルグリッド
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
224 コミット 224 コミット .github/ workflows .github/ workflows app

アプリドキュメント ドキュメントフロントエンド フロントエンドスクリプト スクリプトテスト テスト .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
分散型 AI ワークロード オーケストレーターは、大規模なスケジューリングを困難にする問題 (障害時の排他的実行所有権、部分障害後の調整、リソース制限の強制) を中心に構築されています。スケジューラをテーマにした CRUD チュートリアルではありません。
AetherGrid はワークロードを取得し、リソース要件と制約に基づいて利用可能なコンピューティング ノードと照合し、キューに入れられた、スケジュールされた、実行中、完了、失敗、再試行、キャンセルというライフサイクル全体を管理します。ジョブは、ノードに対して登録されたワーカーを通じて実行され、ジョブ実行の所有権は、単純な割り当てフラグではなく、期限付きのリースによって強制されます。すべてのルートには、キーを発行するエンドポイントを含む API キー認証が必要です。
ほとんどのスケジューラのサイド プロジェクトは、メモリ内辞書をポーリングする while True ループでラップされた単一の main.py スクリプトです。これらは、永続化エンジンを交換したり、新しい制約タイプを追加したり、ジョブが静かに消えた理由や 2 つのワーカーが同時に同じジョブを取得した理由を解明したりする必要があるまでは、正常に機能します。
AetherGrid は 1 つのルールに基づいて構築されています。それは、ドメイン ロジックはデータがどこに存在するかを認識せず、気にも留めないということです。ジョブ、ノード、ワーカー、および割り当てアルゴリズムは、インフラストラクチャへの依存性のない純粋な Python です。データベースは詳細であり、基礎ではありません。このプロジェクトは、これらのアーキテクチャ パターンが単なるカンファレンストークの語彙ではないことを具体的に実証するものです。コードベースは、コードベースが成長し、正確さの要件が厳しくなるにつれて、コードベースを理解しやすく保つためのガードレールです。

えーっと。
このコードベースのすべての自明ではない決定、つまりドメイン ルールが存在する場所に存在する理由、明白に見えるショートカットが拒否された理由、何が壊れ、どのように修正されたかなどは、作成された瞬間に書き留められ、後からポートフォリオ用に再構築されることはありません。 21 個の ADR が /docs/adr に存在します。結論だけでなく推論を確認したい場合は、直接読む価値のあるものがいくつかあります。
ADR 0007 — 調整ループ: 幸せなパスが唯一のパスであると想定するのではなく、死亡したワーカーや期限切れのリースによって矛盾したままになった状態をシステムが検出して修復する方法。
ADR 0011 — ジョブの再利用と調整修復: 調整が既に作業の再割り当てを開始した後に、瀕死の労働者のリース更新が完了する可能性がある実際の競合状態を解消します。
ADR 0012 — 実際のジョブ実行: 強制タイムアウトを使用して本物のサブプロセス実行を構築し、システムが認証されるまでパブリック API から意図的にアクセスできないようにし、コメントではなくテストで存在しないことを証明します。
ADR 0014 — 継続的リース更新: リースが取得時に 1 回ではなく、ジョブの実行時間全体にわたって継続的に更新される理由。
ADR 0015 — API キー認証 : インスタント失効が必要なシステムに対して、JWT ではなく不透明なサーバー発行トークンが選択された理由、および、Job.command を公開する必要があるかどうかの認証の構築が依然として回答しなかった理由について、その疑問は ADR 0020 まで未解決のままであり、狭い範囲で解決されました。ワーカーは、すでに割り当てられている 1 つのコマンドを読み取ることができ、それ以上の広範なものは何もありません。
ADR 0018 — ドメイン所有のスケジューリング ポリシー: list_available() がリポジトリから完全に移動された理由

どのデータベース バックエンドが実行されていたか。
ADR 0019 — スタンドアロン ワーカー エージェント プロセス: インプロセス ジョブの実行を、ネットワーク経由で自身の実行開始を確認する実際のアウトプロセス エージェントに置き換えます。また、プッシュ配信ではなくプルベースのポーリングが選択された理由は、新しい障害や配信保証ロジックを導入する代わりに、このコードベースがすでに信頼しているリコンシリエーションを再利用するためです。
誰かが機能レベルではなくシステム レベルで操作できるかどうかを評価している場合、これが最も早いチェック方法です。
すべてのルートをゲートする API キー認証: 有効な認証情報がなければ、キーを発行するエンドポイントを含め、システム内のエンドポイントに到達できません。最初のキーを作成する唯一の方法は、HTTP 経由ではなく直接データベースにアクセスしてローカルでスクリプトを実行し、パターンが開いたままになる正確なセルフサービス認証情報のホールを閉じることです。
ジョブのライフサイクル管理: 構成可能な再試行ポリシーと優先度を考慮したスケジューリングを備えた明示的な状態遷移 (キューに入っている → スケジュール済み → 実行中 → 完了/失敗/キャンセル)、さらにダッシュボードからアクセス可能なキャンセルおよび再試行アクション
ジョブごとのライフサイクル履歴: すべてのジョブには専用の詳細ページ ( /jobs/{id} ) があり、現在のステータスだけでなく、完全な実際のイベント タイムライン、JobCreated から完了までを表示します。
制約を認識した最適アロケーター: リソース要件とラベルに基づいてワークロードをノードに一致させ、消耗しているノードまたはオフラインになっているノードをスキップします。
ノードのドレイン: 健全なノードは、完全に強制終了することなく、メンテナンスのためにローテーションのスケジュールから外すことができます。スケジューラは、スケジューラ上ですでに実行されているものは完了するまで継続しますが、新しい作業の割り当てを停止します。
ワーカーの登録とハートビート: ノードを登録すると、そのノードに対してワーカーが自動的に登録されるため、すぐに要求できるようになります。

未使用の容量として存在するだけでなく、作業を実行および実行する
排他的なジョブ所有権を持つスタンドアロン ワーカー エージェント: scripts/run_agent.py は、実際の別個のプロセスとして実行され、割り当てられた作業について HTTP 経由で API をポーリングし、実際のローカル サブプロセスとして実行し、現在実行中のジョブとは無関係に、エージェントの存続期間全体にわたって独自のバックグラウンド スレッドでハートビートします (ADR 0019 を参照)。これは、ダッシュボードを実行するワーカーのライブネス メカニズムとして、ダッシュボードのクライアント側のハートビートを置き換えます。エージェント プロセスが接続されていないワーカーは、依然としてノードの活性のみにフォールバックします。すべてのワーカーは、登録時に設定された明示的な manage_by フィールド ( DASHBOARD または AGENT ) でタグ付けされます。インプロセス スケジューラ ループは、AGENT とマークされたワーカーを完全にスキップするため、スタンドアロン エージェントのジョブはエージェントによって 1 回だけ実行され、インプロセス パスと競合することはありません。
リースベースの実行所有権: ワーカーがジョブを受け入れると、そのジョブに対して更新可能な期限切れのリースを保持し、ジョブの実行全体にわたって継続的に更新されるため、再試行、再接続、ネットワーク障害、単純に長時間実行されるジョブによって 2 つのワーカーが同じジョブを実行することはありません。
明示的な実行開始の確認: POST /workers/{worker_id}/jobs/{job_id}/start により、実際にジョブを実行しているもの、ダッシュボード管理ワーカーのインプロセス スケジューラー ループ、エージェント管理ワーカーのスタンドアロン エージェント (ADR 0019) によって、実行が本当に開始されたことを確認できます。これは、ジョブを Scheduled から Running に移行する 1 つの呼び出しです。割り当てだけでは機能しなくなりました (ADR 0019 を参照)
タイムアウトを強制した実際のサブプロセス実行: コマンドを含むジョブは実際のサブプロセスとして実行され、ジョブが実行タイムアウトを超過した場合は 2 段階のシャットダウン (正常な SIGTERM 、猶予期間後の SIGKILL) が行われます。
ノードの活性度tr

ACK : ハートビートベースのヘルスチェック、オフラインノードの自動検出、作業が失敗したときやノードが消失したときのリソースの再利用
制限付き再試行による調整 : 死んだワーカーまたはオフライン ノードによって放棄されたジョブは、再試行予算内でキューに再利用され、その予算が使い果たされると完全に失敗します。そのため、単一の不健全なノードによってジョブが再割り当てされ、無期限に放棄されることはありません。再利用は、調停がすでにリースの再割り当てを開始した後に瀕死のワーカーの更新が着地する可能性がある実際のレースを終了するために命令されます。
ドメイン イベントの記録: ジョブが通過するすべてのライフサイクル遷移、 JobCreated 、 JobScheduled 、 WorkerAssigned 、 LeaseAcquired 、 LeaseClosed 、 JobCompleted / JobFailed 、および JobReclaimed は、発生した正確な時点で不変イベントとして永続化されます。
クラスター全体のライブ イベント フィード: ダッシュボード上の GET /events とリアルタイムのアクティビティ フィードは、3 秒ごとにポーリングされるため、個々のジョブが独自の詳細ページで伝えるストーリーも、クラスター全体で発生するのと同様に表示されます。
ワーカーの可視性 : 登録されているすべてのワーカー、そのステータス、属するノード、実行内容、最後に表示された日時を示す専用のワーカー テーブル
マルチページ ダッシュボード: 単一ページではなく実際のクライアント側ルーティング ( / 、 /nodes 、 /jobs 、 /jobs/{id} )、サイドバーでアクティブなルートが強調表示されます。
システムは 4 つの層に分割されており、依存関係は内側を向いています。
Domain : Job 、 Node 、 Worker 、 Lease 、 Event 、および ApiKey 集計は、独自の不変条件を強制します。スケジューリング アルゴリズムとジョブ ライフサイクル ステート マシンは、FastAPI や psycopg からのインポートを行わず、プレーンな Python としてここに存在します。インフラストラクチャ層を完全に削除しても、ドメイン テストは引き続き合格します。
アプリケーション : ScheduleJobServ などのサービス

Ice / SchedulerService 、 AssignWorkerService 、 AcquireLeaseService 、 StartJobService 、 DrainNodeService 、 ClusterHealthService 、および AuthenticateApiKeyService は、1 つ下の層に属するビジネス ルールを埋め込むことなく、ドメイン オブジェクトとリポジトリを調整します。 WorkerExecutionLoop は、割り当てられたジョブを実際のサブプロセスとして実行することでワーカーを駆動し、ジョブの実行時間全体にわたってバックグラウンド スレッドでリースを継続的に更新し、実際の結果を記録し、その結果に関係なくリースを解放します。更新が失敗するということは、リースがすでに他の場所で再利用されており、ループはその結果を別のワーカーの進行中または完了した作業に対して永続化するリスクを冒さずに破棄することを意味します。 ReconciliationLoop は、ハッピー パスでは検出できない障害モード (ワーカーのクラッシュ、リースの期限切れ、インフラストラクチャの障害によって一貫性がなくなった状態) を検出します。
インフラストラクチャ : PostgreSQL 実装はすべてのリポジトリ ( Node 、 Job 、 Worker 、 Lease 、 Event 、 ApiKey ) に存在し、ORM ではなく生の psycopg で記述されています。これは、クエリ動作とトランザクション境界を抽象化するのではなく可視化するための意図的な選択です。 Node 、 Job 、 Event にはさらに、ローカル開発用の SQLite 実装があります。アピック

[切り捨てられた]

## Original Extract

A decentralized compute orchestration platform powered by strict Domain-Driven Design (DDD). Decoupled domain core, event-driven state reconciliation, and zero-overhead cluster management. - wycliffRotich-dev/aethergrid

GitHub - wycliffRotich-dev/aethergrid: A decentralized compute orchestration platform powered by strict Domain-Driven Design (DDD). Decoupled domain core, event-driven state reconciliation, and zero-overhead cluster management. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
wycliffRotich-dev
/
aethergrid
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
224 Commits 224 Commits .github/ workflows .github/ workflows app app docs docs frontend frontend scripts scripts tests tests .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
A distributed AI workload orchestrator built around the problems that make scheduling hard at scale: exclusive execution ownership under failure, reconciliation after partial failures, and enforced resource limits. Not a CRUD tutorial with a scheduler theme.
AetherGrid takes workloads, matches them against available compute nodes based on resource requirements and constraints, and manages the full lifecycle: queued, scheduled, running, completed, failed, retried, cancelled. Jobs run through workers registered against nodes, and job execution ownership is enforced through time-bound leases rather than a simple assignment flag. Every route requires API key authentication, including the endpoint that issues keys.
Most scheduler side-projects are a single main.py script wrapped in a while True loop polling an in-memory dictionary. They work fine, right up until you need to swap the persistence engine, add a new constraint type, or figure out why a job silently disappeared, or why two workers picked up the same job at once.
AetherGrid was built around one rule: the domain logic doesn't know or care where the data lives. Jobs, nodes, workers, and the allocation algorithm are pure Python with zero infrastructure dependencies. The database is a detail, not the foundation. This project is a concrete demonstration that these architectural patterns aren't just conference-talk vocabulary; they're guardrails that keep a codebase understandable as it grows, and as its correctness requirements get harder.
Every non-obvious decision in this codebase, why a domain rule lives where it does, why an obvious-looking shortcut was rejected, what broke and how it got fixed, is written down at the moment it was made, not reconstructed afterward for a portfolio. 21 ADRs live in /docs/adr . A few worth reading directly if you want to see the reasoning, not just the conclusion:
ADR 0007 — Reconciliation Loop : how the system detects and repairs state left inconsistent by dead workers and expired leases, instead of assuming the happy path is the only path.
ADR 0011 — Job Reclaim and Reconciliation Repair : closing a real race condition where a dying worker's lease renewal could land after reconciliation had already started reassigning its work.
ADR 0012 — Real Job Execution : building genuine subprocess execution with enforced timeouts, then deliberately keeping it unreachable from the public API until the system had authentication, and proving that absence with a test rather than a comment.
ADR 0014 — Continuous Lease Renewal : why a lease is renewed continuously for a job's entire runtime instead of once at acquisition.
ADR 0015 — API Key Authentication : why opaque server-issued tokens were chosen over JWTs for a system that needs instant revocation, and why building authentication still didn't answer whether Job.command should be exposed, that question stayed open until ADR 0020 , which resolved it narrowly: a worker can read the one command already assigned to it, nothing broader.
ADR 0018 — Domain Owns Scheduling Policy : why list_available() moved out of the repository entirely, since deciding which nodes are eligible for scheduling is a business rule, not a persistence concern, and letting infrastructure decide that would have made scheduling behavior dependent on which database backend was running.
ADR 0019 — Standalone Worker Agent Process : replacing in-process job execution with a real out-of-process agent that confirms its own execution start over the network, and why pull-based polling was chosen over push delivery, since it reuses reconciliation this codebase already trusts instead of introducing new failure and delivery-guarantee logic.
If you're evaluating whether someone can operate at a systems level rather than a feature level, this is the fastest way to check.
API key authentication gating every route : no endpoint in the system, including the one that issues keys, is reachable without a valid credential. The only way to mint the first key is a script run locally with direct database access, never over HTTP, closing the exact self-service-credential hole that pattern would otherwise leave open
Job lifecycle management : explicit state transitions (Queued → Scheduled → Running → Completed/Failed/Cancelled) with configurable retry policies and priority-aware scheduling, plus cancel and retry actions reachable from the dashboard
Per-job lifecycle history : every job has a dedicated detail page ( /jobs/{id} ) showing its full real event timeline, JobCreated through completion, not just its current status
Constraint-aware best-fit allocator : matches workloads to nodes based on resource requirements and labels, while skipping nodes that are draining or offline
Node draining : a healthy node can be taken out of scheduling rotation for maintenance without killing it outright; the scheduler stops assigning it new work while anything already running on it continues to completion
Worker registration and heartbeats : registering a node automatically registers a worker against it, so it's immediately capable of claiming and executing work, not just existing as unused capacity
Standalone worker agent with exclusive job ownership : scripts/run_agent.py runs as a real, separate process, polling the API over HTTP for assigned work, executing it as a real local subprocess, and heartbeating on its own background thread for the agent's entire lifetime, independent of whatever job it's currently executing (see ADR 0019 ). This replaces the dashboard's client-side heartbeat as the liveness mechanism for any worker running it; a worker with no agent process attached still falls back to node liveness alone. Every worker is tagged with an explicit managed_by field set at registration ( DASHBOARD or AGENT ); the in-process scheduler loop skips any worker marked AGENT entirely, so a standalone agent's jobs are executed exactly once, by the agent, never raced against the in-process path.
Lease-based execution ownership : when a worker accepts a job, it holds a renewable, expiring lease on that job, continuously renewed for the job's entire execution, so retries, reconnects, network failures, and jobs that simply run long can't result in two workers executing the same job
Explicit execution-start confirmation : POST /workers/{worker_id}/jobs/{job_id}/start lets whatever is actually executing a job, the in-process scheduler loop for dashboard-managed workers, a standalone agent for agent-managed ones (ADR 0019), confirm execution has genuinely begun. This is the one call that transitions a job from Scheduled to Running ; assignment alone no longer does (see ADR 0019 )
Real subprocess execution with enforced timeouts : jobs with a command run as real subprocesses, with a two-stage shutdown (graceful SIGTERM , then SIGKILL after a grace period) if a job overruns its execution timeout
Node liveness tracking : heartbeat-based health checks, automatic detection of offline nodes, and resource reclamation when work fails or nodes disappear
Reconciliation with bounded retries : jobs abandoned by a dead worker or an offline node are reclaimed back to the queue within their retry budget, and fail outright once that budget is exhausted, so a single unhealthy node can't cause a job to be reassigned and abandoned indefinitely, with the reclaim ordered to close a real race where a dying worker's renewal could land after reconciliation had already started reassigning its lease
Domain event recording : every lifecycle transition a job goes through, JobCreated , JobScheduled , WorkerAssigned , LeaseAcquired , LeaseReleased , JobCompleted / JobFailed , and JobReclaimed , is persisted as an immutable event at the exact point it happens
Live cluster-wide event feed : GET /events and a real-time Activity Feed on the dashboard, polling every 3 seconds, so the story an individual job tells on its own detail page is also visible as it happens across the whole cluster
Worker visibility : a dedicated Workers table showing every registered worker, its status, the node it belongs to, what it's running, and when it was last seen
Multi-page dashboard : real client-side routing ( / , /nodes , /jobs , /jobs/{id} ) instead of a single page, with active-route highlighting in the sidebar
The system is split into four layers, with dependencies pointing inward:
Domain : Job , Node , Worker , Lease , Event , and ApiKey aggregates enforce their own invariants. The scheduling algorithm and job lifecycle state machine live here as plain Python, with no imports from FastAPI or psycopg. Delete the infrastructure layer entirely and the domain tests still pass.
Application : Services such as ScheduleJobService / SchedulerService , AssignWorkerService , AcquireLeaseService , StartJobService , DrainNodeService , ClusterHealthService , and AuthenticateApiKeyService coordinate domain objects and repositories without embedding business rules that belong one layer down. A WorkerExecutionLoop drives a worker through executing its assigned job as a real subprocess, continuously renewing its lease on a background thread for the job's entire runtime, recording the real outcome, and releasing the lease regardless of that outcome. A renewal that fails means the lease has already been reclaimed elsewhere, and the loop discards its result rather than risk persisting it against another worker's in-progress or completed work. A ReconciliationLoop catches the failure modes the happy path can't: crashed workers, expired leases, state left inconsistent by infrastructure failures.
Infrastructure : PostgreSQL implementations exist for every repository ( Node , Job , Worker , Lease , Event , ApiKey ), written with raw psycopg instead of an ORM, a deliberate choice to keep query behavior and transaction boundaries visible rather than abstracted away. Node , Job , and Event additionally have SQLite implementations for local development; ApiK

[truncated]

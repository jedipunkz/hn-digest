---
source: "https://labyrinthanalyticsconsulting.com/blog/claude-memory-primitive-vs-loreconvo-vs-claude-mem-vs-mem0"
hn_url: "https://news.ycombinator.com/item?id=49345976"
title: "AI Memory for Claude: An Honest 4-Way Comparison"
article_title: "AI Memory for Claude: An Honest 4-Way Comparison | Labyrinth Analytics"
image: "https://labyrinthanalyticsconsulting.com/blog/claude-memory-primitive-vs-loreconvo-vs-claude-mem-vs-mem0/opengraph-image?ef4123aa1d39876a"
author: "labyrinthAC"
captured_at: "2026-08-18T14:23:26Z"
capture_tool: "hn-digest"
hn_id: 49345976
score: 1
comments: 0
posted_at: "2026-08-18T14:17:24Z"
tags:
  - hacker-news
  - translated
---

# AI Memory for Claude: An Honest 4-Way Comparison

- HN: [49345976](https://news.ycombinator.com/item?id=49345976)
- Source: [labyrinthanalyticsconsulting.com](https://labyrinthanalyticsconsulting.com/blog/claude-memory-primitive-vs-loreconvo-vs-claude-mem-vs-mem0)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T14:17:24Z

## Translation

タイトル: クロードの AI 記憶: 正直な 4 者間比較
記事のタイトル: クロードの AI 記憶: 正直な 4 者間比較 |ラビリンス分析
説明: AI メモリが何をすべきかについての 4 つのツール、4 つの賭け。コスト、プライバシー、ワークフローにとってどちらが有利かを、ワークフローの 1 つを構築した実践者から語ります。

記事本文:
クロードの AI 記憶: 正直な 4 者間比較 | Labyrinth Analytics Labyrinth Analytics ビジネス向けホーム サービス 仕事ツール ライフタイム価格 ブログについて お問い合わせ ブログに戻る 2026 年 8 月 14 日 LoreConvo AI Memory for Claude: An 正直な 4 者比較
AI メモリが何をすべきかについての 4 つのツール、4 つの賭け。コスト、プライバシー、ワークフローにとってどちらが有利かを、ワークフローの 1 つを構築した実践者から語ります。
プロンプトをつなぎ合わせ、パイプラインからデータを取得し、最後のセッションでどのような決定を下したかを思い出そうとするのに 1 日を費やすと、摩擦が無駄な時間として現れます。問題はモデル自体ではなく、ユーザーとモデルの間にあるメモリ層です。過去 1 年間で、クロードの記憶に対する 4 つのアプローチが登場しました。この記事では、Claude Memory Primitive、claude-mem、mem0、LoreConvo を直接比較します。私は LoreConvo と mem0 および LoreConvo と claude-mem について個別に詳細を書きました。この投稿では、4 方向の全体像を 1 か所にまとめています。目標は、アーキテクチャ、コスト、プライバシー、ワークフローにおけるトレードオフを明確に把握し、状況に応じてツールを選択できるようにすることです。私は 4 つすべてを試しました。私は LoreConvo を構築したので、他の人がどこで勝てるかを正確に説明します。
アーキテクチャとデータの所有権
最も明らかな違いは、メモリがどこに存在するかです。
Claude Memory Primitive は、Anthropic が管理するクラウドでホストされるグラフにセッション スニペットを保存します。リクエストを送信すると、Claude が短い概要を返し、サービスは API キーに関連付けられた記録を保持します。このモデルは始めるのが最も簡単ですが、データが Anthropic の環境から離れることはありません。厳格なデータ処理ポリシーに準拠する必要があるチームにとって、その保持は取引の妨げとなります。
claude-mem は、グラフを JSON f に変換する薄い層を追加します。

クライアント側のファイル。ファイルはローカル ディレクトリに書き込まれますが、ツールはバックアップのためにファイルをリモート ストアに定期的に同期します。結果はハイブリッドです。検査できるローカル コピーを取得しますが、耐久性についてはクラウド サービスにも依存します。先週の火曜日に自分が考え出したことをクロードに覚えていてもらいたいだけの多くの個人開発者にとって、その自動動作はまさに正しいものです。
mem0 は別のルートをたどります。各インタラクションのインデックスを作成する埋め込みモデルを使用して、ローカル データベース上にベクトル ストアを構築します。このライブラリは、プログラムでメモリを追加、クエリ、削除できるメモリ管理操作を公開します。インデックスはマシン上に存在するため、完全に制御できますが、埋め込みモデルも自分で管理する必要があるため、複雑さと継続的なコストの両方が追加されます。
LoreConvo は、ポータビリティとクロスサーフェス リーチを組み合わせたローカル ファーストの設計を採用しています。すべてのセッション データは、所有する単一の SQLite ファイルに格納されます。ファイルは、既に使用しているツールを使用して移動、バックアップ、またはバージョン管理できます。隠されたクラウド コンポーネントはありません。同時に、MCP サーバーは、クライアントごとの構成を行わずに、メモリ レイヤーを Claude Code、OpenAI Codex、Cursor、Hermes Agent に公開します。 .mcp.json ファイルをプロジェクト ルートにドロップすると、サーバーがそれを自動的に検出します。ストレージはローカルであるため、プライバシーが保証されます。ファイルを読み取ることができるのは、あなたまたは明示的に共有しているチームメイトだけです。
コストも 4 つの軸が分岐するもう 1 つの軸です。
Claude Memory Primitive は、Claude API の使用料にバンドルされています。個別のメモリ料金はありませんが、グラフの読み取りまたは書き込みを呼び出すたびに API コストが発生します。ラップトップを使用する個人開発者の場合、これは安価です。 1 日に何百ものセッションを実行しているチームの場合、追加の通話はすぐに加算されます。
クロード・ミーは私に自由です

インストールします。オプションのリモート バックアップ サービスは、保存されるギガバイトごとに料金がかかります。数メガバイトのセッション データにはほとんど費用がかかりません。数週間分の履歴をアーカイブすると、料金は使用量に応じて直線的に変化します。
mem0 自体は無料ですが、埋め込みモデルをプロビジョニングする必要があります。ホストされた埋め込み API の使用には通常、1,000 トークンあたり数セントのコストがかかり、ローカル モデルの実行には GPU リソースが消費されます。すでに GPU 能力を備えているデータ エンジニアにとっては限界コストは低くなりますが、小規模なチームにとっては、外部 API 料金が隠れた出費となり、数か月にわたって増加していきます。
LoreConvo は、予測可能な 2 層構造を提供します。無料枠では最大 50 セッションが提供され、ほとんどの実験や短期プロジェクトに対応します。 Pro の料金は月額 8 ドルで、セマンティック検索、関連セッションの検出、チーム メモリ共有が追加されると同時に、セッション制限が削除されます。これらの機能はすべて同じローカル SQLite ファイルに対して実行されるため、通話ごとに追加料金がかかることはありません。個人の開発者にとって、Pro の費用は、クラウドネイティブ オプションで通常生成される追加の API 呼び出しよりも安く、小規模なチームの場合は月額定額料金なので予算編成が簡単になります。
データをクラウド サービスに渡すと、プロバイダーのセキュリティ慣行を暗黙的に信頼することになります。 Claude Memory Primitive は転送中および保存中のデータを暗号化しますが、Anthropic はコピーを保持します。独自のコード、規制されたデータセット、または内部アーキテクチャの決定を扱っている場合、その保持によりコンプライアンスのリスクが生じます。
claude-mem のハイブリッド アプローチでは、監査できるローカル コピーが提供されますが、リモート バックアップでは依然としてデータがサードパーティのバケットに保存されます。同期を無効にすることはできますが、自動耐久性機能が失われるため、意識的に無効にすることを忘れないようにする必要があります。
mem0 が全責任を負う

あなたに。インデックス ファイルはローカルにあり、標準ツールで検査できます。ただし、ライブラリはアクセス制御を強制しないため、ファイル システムにアクセスできるすべてのプロセスがメモリを読み取ることができます。独自の OS レベルの権限を設定する必要があります。
LoreConvo の設計は所有権を中心に構築されています。 SQLite ファイルは選択したディレクトリに保存され、メモリ検査 UI を使用すると、1 つのコマンドでセッションを一覧表示、フィルタリング、削除できます。自動保存フックは、ユーザーの操作なしですべてのセッションの終了時に実行されます。セッションに十分な信号が含まれている場合、意思決定、技術スタックの事実、未解決の質問をキャプチャするヒューリスティックな概要を抽出します。これはベストエフォートであり、保存するたびに保証されるものではありません。データがマシンから流出することはないため、完全な制御を維持できます。 Pro ユーザーは、中央サーバーを使用せずに、選択したセッションを JSON にエクスポートし、チームメイトにインポートさせることができます。
あなたの状況に適したツールはどれですか
セットアップが不要で、Anthropic によるクラウドでのメモリ管理に慣れている場合は、メモリ プリミティブが最速のパスです。これは、プライバシーが懸念されておらず、API の使用状況にまだ細心の注意を払っていない、簡単なプロトタイプに適しています。
claude-mem は GitHub の 45,000 スターを獲得しました。自動バックアップを備えたローカル コピーが必要で、ハイブリッド クラウド モデルに慣れている場合は、これが最も洗練されたコミュニティ オプションです。シームレスなキャプチャ (作業を続けると会話が記憶される) は、メモリ層について考えたくない個人開発者にとって非常に便利です。
mem0 は、ディープ ベクトル検索を必要とし、組み込みインフラストラクチャの管理に意欲的なチームに最適です。数千の保存されたファクトにわたるセマンティック類似性検索から恩恵を受けるパイプラインを構築していて、それをサポートするための GPU または API 予算がある場合、mem0 は y を提供します。

他のオプションよりも検索の深さが長くなります。
LoreConvo は、クロスサーフェスに到達し、コストが予測可能な、ポータブルで所有可能なファイルを求める実務者に適しています。 FTS5 全文検索は、埋め込み層を使用せずにほとんどのリコール ニーズに対応します。セッション リンク、プロジェクトのタグ付け、および自動ロード フックにより、手動で記録することなく実行全体でコンテキスト チェーンが作成されます。 If you work across multiple surfaces -- Claude Code in the morning, Cursor in the afternoon, a headless pipeline at night -- LoreConvo carries the same memory layer to all of them. The Python fallback script save_to_loreconvo.py lets any script read or write memories without registering an MCP tool, keeping the integration lightweight for automation.
For data engineers running multiple pipelines, the combination of skill history tracking, project tagging, and session linking turns the memory store into a lightweight audit trail: what decisions drove each pipeline design, which schemas changed and why, and where to pick up a refactor that stalled two weeks ago.
正直な答えは、これら 4 つのツールはいずれも、あらゆるユースケースに間違っていないということです。クロード メモリ プリミティブは、ゼロフリクションの開始点です。 claude-mem は、ほとんどの開発者が最初に到達する、コミュニティによって構築された標準です。 mem0 は、大規模なセマンティック検索が必要で、埋め込みのオーバーヘッドを管理できる場合に適切な選択です。 LoreConvo is the option when you want a single file you own, cross-surface reach you do not have to configure per client, and a flat monthly cost that does not scale with query volume.
完全な LoreConvo ツール セットとインストール手順は /tools で参照できます。 If you are comparing these options for a specific pipeline or agent architecture, reach out -- I am happy to walk through the trade-offs for your stack.
PS: このような投稿を入手してください

毎週配信 - 「迷宮からの派遣」を購読してください。
ラビリンス・アナリティクス・コンサルティング代表取締役。メインフレームから AI エージェントまで、6 つのテクノロジー世代にわたって 35 年以上の経験を持つデータ エンジニア。彼女は、LangGraph パ​​イプライン、データ ウェアハウス、LoreConvo と LoreDocs の背後にあるメモリ ツールを設計しています。ワシントン州に拠点を置く。
LinkedIn GitHub 関連書籍について
LoreDocs: 耐久性のある AI Knowledge Vaults
LoreDocs は、バージョン管理され、検索可能で、モデルにアクセスできる、単一の会話よりも長く存続する構造化成果物のためのローカルのナレッジ ボールトです。
同意と出所にお金を払う価値がある理由
無料の AI 記憶ツールは印象的です。彼らがベンチマークを行っていないものは、同意モデル、出所追跡、および検証されたクロスサーフェス同一性です。
永続的な知識に独自のストアが必要な理由
LoreConvo はセッション メモリを解決しました。次に登場するのは LoreDocs です。これは、単一のチャットよりも長く存続する洞察、スキーマ、成果物を保管するローカルのナレッジ ボールトです。
新しい投稿が公開されると 1 通のメールが送信されます: エージェント AI、データ エンジニアリング、記憶ツール。スパム、アップセル、AI サマリーはありません。いつでも購読を解除してください。
Labyrinth Analytics Consulting は、組織がデータの暗部に対処できるよう支援します。詳細については、labyrinthanalyticsconsulting.com をご覧ください。
ブログ「Labyrinth Analytics Consulting」の詳細

## Original Extract

Four tools, four bets about what AI memory should do. Cost, privacy, and which wins for your workflow -- from a practitioner who built one of them.

AI Memory for Claude: An Honest 4-Way Comparison | Labyrinth Analytics Labyrinth Analytics Home Services For Business Work Tools Lifetime Pricing Blog About Contact Back to Blog August 14, 2026 LoreConvo AI Memory for Claude: An Honest 4-Way Comparison
Four tools, four bets about what AI memory should do. Cost, privacy, and which wins for your workflow -- from a practitioner who built one of them.
When you spend a day stitching together prompts, pulling data from a pipeline, and then trying to remember what decision you made in the last session, the friction shows up as wasted time. The problem is not the model itself -- it is the memory layer that sits between you and the model. Over the past year four approaches to Claude memory have emerged. This post compares them directly: the Claude Memory Primitive, claude-mem, mem0, and LoreConvo. I have written individual deep-dives on LoreConvo vs mem0 and LoreConvo vs claude-mem ; this post consolidates the full four-way picture in one place. The goal is a clear picture of the trade-offs in architecture, cost, privacy, and workflow so you can pick the tool that matches your situation. I have tried all four; I built LoreConvo, so I will tell you exactly where the others win.
Architecture and data ownership
The most obvious difference is where the memory lives.
The Claude Memory Primitive stores session snippets in a cloud-hosted graph that Anthropic manages. You send a request, Claude returns a short summary, and the service keeps a record tied to your API key. This model is the simplest to start with, but the data never leaves Anthropic's environment. For teams that must comply with strict data-handling policies, that retention is a deal-breaker.
claude-mem adds a thin layer that converts the graph into a JSON file on the client side. The file is written to a local directory, but the tool periodically syncs it back to a remote store for backup. The result is a hybrid: you get a local copy you can inspect, but you also rely on a cloud service for durability. For a lot of solo developers who just want Claude to remember what they figured out last Tuesday, that automatic behavior is exactly right.
mem0 takes a different route. It builds a vector store on top of a local database using an embedding model to index each interaction. The library exposes memory management operations that let you add, query, and delete memories programmatically. Because the index lives on your machine, you have full control -- but you also manage the embedding model yourself, which adds both complexity and ongoing cost.
LoreConvo uses a local-first design that combines portability with cross-surface reach. All session data lands in a single SQLite file you own. The file can be moved, backed up, or versioned with any tool you already use -- there is no hidden cloud component. At the same time, the MCP server exposes the memory layer to Claude Code, OpenAI Codex, Cursor, and Hermes Agent without any per-client configuration. You drop a .mcp.json file in the project root and the server discovers it automatically. Because the storage is local, privacy is guaranteed: only you, or teammates you explicitly share with, can read the file.
Cost is another axis where the four diverge.
The Claude Memory Primitive is bundled with Claude API usage fees. There is no separate memory charge, but every call that reads from or writes to the graph incurs API cost. For a solo developer on a laptop this can be inexpensive; for a team running hundreds of sessions per day the extra calls add up quickly.
claude-mem is free to install. The optional remote backup service is priced per gigabyte stored. A few megabytes of session data costs almost nothing; once you archive weeks of history the price scales linearly with usage.
mem0 itself is free, but you need to provision an embedding model. Using a hosted embedding API typically costs a few cents per thousand tokens, and running a local model draws on GPU resources. For data engineers who already have GPU capacity the marginal cost is low, but for smaller teams the external API fees become a hidden expense that compounds over months.
LoreConvo offers a predictable two-tier structure. The free tier gives you up to fifty sessions, which covers most experimentation and short-term projects. Pro costs eight dollars per month and removes the session limit while adding semantic search, related-session discovery, and team memory sharing. All of those features run against the same local SQLite file, so there are no per-call fees layered on top. For a solo developer, Pro costs less than the extra API calls that the cloud-native options typically generate, and for a small team the flat monthly price makes budgeting straightforward.
When you hand data to a cloud service you implicitly trust the provider's security practices. The Claude Memory Primitive encrypts data in transit and at rest, but Anthropic retains a copy. If you are working with proprietary code, regulated datasets, or internal architecture decisions, that retention introduces compliance risk.
claude-mem's hybrid approach gives you a local copy you can audit, but the remote backup still stores the data in a third-party bucket. You can disable the sync, but then you lose the automatic durability feature -- and you have to remember to disable it consciously.
mem0 puts the entire responsibility on you. The index files are local and inspectable with standard tools. However, the library does not enforce access controls, so any process with file system access can read the memories. You must set up your own OS-level permissions.
LoreConvo's design is built around ownership. The SQLite file lives in a directory you choose, and the memory inspection UI lets you list, filter, and delete sessions with a single command. The auto-save hook runs at the end of every session without any user action. It extracts a heuristic summary that captures decisions, tech-stack facts, and open questions when the session contains enough signal -- this is best-effort, not a guarantee on every save. Because the data never leaves your machine, you retain full control. Pro users can export selected sessions to JSON and let a teammate import them, all without a central server.
Which tool fits your situation
If you need zero setup and are comfortable with Anthropic managing your memory in the cloud, the Memory Primitive is the fastest path. It works well for quick prototypes where privacy is not a concern and you are not already paying close attention to API usage.
claude-mem earns its 45,000 GitHub stars. If you want a local copy with automatic backup and you are comfortable with the hybrid cloud model, it is the most polished community option. The seamless capture -- you keep working and your conversations get remembered -- is genuinely useful for solo developers who do not want to think about the memory layer.
mem0 wins for teams that need deep vector search and are willing to manage the embedding infrastructure. If you are building pipelines that benefit from semantic similarity lookups across thousands of stored facts, and you have the GPU or the API budget to support that, mem0 gives you more retrieval depth than the other options.
LoreConvo fits the practitioner who wants a portable, ownable file with cross-surface reach and predictable costs. The FTS5 full-text search handles most recall needs without an embedding layer. Session linking, project tagging, and the auto-load hook create context chains across runs without manual bookkeeping. If you work across multiple surfaces -- Claude Code in the morning, Cursor in the afternoon, a headless pipeline at night -- LoreConvo carries the same memory layer to all of them. The Python fallback script save_to_loreconvo.py lets any script read or write memories without registering an MCP tool, keeping the integration lightweight for automation.
For data engineers running multiple pipelines, the combination of skill history tracking, project tagging, and session linking turns the memory store into a lightweight audit trail: what decisions drove each pipeline design, which schemas changed and why, and where to pick up a refactor that stalled two weeks ago.
The honest answer is that none of these four tools is wrong for every use case. The Claude Memory Primitive is the zero-friction starting point. claude-mem is the community-built standard that most developers reach for first. mem0 is the right choice when you need semantic retrieval at scale and can manage the embedding overhead. LoreConvo is the option when you want a single file you own, cross-surface reach you do not have to configure per client, and a flat monthly cost that does not scale with query volume.
You can see the full LoreConvo tool set and install instructions at /tools . If you are comparing these options for a specific pipeline or agent architecture, reach out -- I am happy to walk through the trade-offs for your stack.
PS: Get posts like this delivered weekly -- subscribe to Dispatches from the Labyrinth .
Principal at Labyrinth Analytics Consulting. Data engineer with 35+ years across six technology generations, from mainframes to AI agents. She designs LangGraph pipelines, data warehouses, and the memory tooling behind LoreConvo and LoreDocs. Based in Washington State.
LinkedIn GitHub About Related Reading
LoreDocs: Durable AI Knowledge Vaults
LoreDocs is a local knowledge vault for the structured artifacts that outlive any single conversation -- versioned, searchable, and model-accessible.
Why Consent and Provenance Are Worth Paying For
Free AI memory tools are impressive. Here is what they do not benchmark: consent model, provenance tracking, and verified cross-surface identity.
Why durable knowledge needed its own store
LoreConvo solved session memory. LoreDocs is what comes next -- a local knowledge vault for the insights, schemas, and artifacts that outlive a single chat.
One email when a new post is published: agentic AI, data engineering, and memory tools. No spam, no upsell, no AI summaries. Unsubscribe anytime.
Labyrinth Analytics Consulting helps organizations navigate the dark corners of their data. Learn more at labyrinthanalyticsconsulting.com .
More from the blog Labyrinth Analytics Consulting

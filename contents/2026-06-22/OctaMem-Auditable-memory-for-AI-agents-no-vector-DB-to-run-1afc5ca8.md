---
source: "https://octamem.com"
hn_url: "https://news.ycombinator.com/item?id=48628724"
title: "OctaMem: Auditable memory for AI agents, no vector DB to run"
article_title: "OctaMem · Persistent memory for AI systems"
author: "Mossiah"
captured_at: "2026-06-22T12:07:16Z"
capture_tool: "hn-digest"
hn_id: 48628724
score: 1
comments: 0
posted_at: "2026-06-22T11:26:39Z"
tags:
  - hacker-news
  - translated
---

# OctaMem: Auditable memory for AI agents, no vector DB to run

- HN: [48628724](https://news.ycombinator.com/item?id=48628724)
- Source: [octamem.com](https://octamem.com)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T11:26:39Z

## Translation

タイトル: OctaMem: AI エージェントの監査可能なメモリ、実行するベクトル DB なし
記事のタイトル: OctaMem · AI システム用の永続メモリ
説明: OctaMem は、AI システム用の永続メモリ層です。セッション、モデル、ワークフロー全体に伝わる意味記憶、エピソード記憶、手続き記憶。

記事本文:
OctaMem · AI システム用の永続メモリ コンテンツにスキップ OctaMem ホーム プラットフォーム 価格 開発者の信頼 ドキュメントについて サインイン 無料で始める ナビゲート
営業担当者に相談する すでにアカウントをお持ちです · サインインします。 01 メモリ インフラストラクチャ ライブ · ステータス公称版 · 2026 エージェントが忘れてはならないすべてのインフラストラクチャ。
AI エージェントの永続メモリ層。忘れることが許されないスタック向けに構築されています。
営業担当者に相談してください。カードがありません。無料枠で 2 GB。すでに実行しているスタックで動作します · OpenAI OpenAI ·
記憶力のないエージェントは 2 つの点で失敗し、それがバランスシートに現れます。トークンを焼き付けて文脈を再読み込みすることと、苦労して獲得した組織的な知識を門前払いにすることです。メモリ層は両方に答えます。
トークンの支出は毎ターン増加します。
メモリがないと、呼び出しごとに同じコンテキストが再送信されます。会話自体が再説明され、プロンプトが膨らみ、フロンティア モデルの料金を支払って、モデルが 1 時間前にすでに話した内容を再読します。
呼び出しごとのコンテキストが少なくなる - 関連するものだけが取得され、挿入されます
繰り返しはありません - 事実と決定は再送信されずに保持されます
安価なモデルは、受け取ったコンテキストがより鮮明であれば、その地位を維持します。
組織の知識は一元化されていません。
エージェントとチームが学んだことは、散在したセッション、ローカルのメモ、および個々の責任者に反映されます。従業員が退職するときは、従業員も一緒に退職します。何も合成されず、組織が所有するものは何もありません。
組織全体のインテリジェンス - すべてのエージェントが 1 つの共有レイヤーから読み取ります
知識は人が去っても残ります。知識はその人ではなく記憶の中に生きます。
セッションごとにリセットするのではなく、チーム全体でコンテキストを複合化する
メモリはパッチではなく、インフラストラクチャである必要があります。
すべてのリクエストは同じ規律あるサイクルを通過します。 OctaMem は GE を起動しません

テキストの 1 バケットにわたる neric 検索。それぞれが異なる目的を果たす 3 つのメモリ タイプからコンテキストを再構築し、モデル用に再構築します。
検索サイクル サイクルを追加 検索サイクル・実行中
02 / アクセス・クォータ セキュリティ層
03 / OctaMem エージェント検索サービス
04 / 3層のメモリー層
05 / アプリに戻る 統一コンテキスト
セマンティックエピソード手順 01 アプリまたは MCP 呼び出し元 02 セキュリティ層 アクセス・クォータ 03 取得サービス OctaMem エージェント 04 メモリ層 3 層 05 統合コンテキスト アプリの図に戻る。 1 · 検索サイクル · ステージ 1/5 自動再生検索サイクル サイクルを追加します。 1 · 検索 · ステージ 1 / 5 01 API、SDK、または MCP を介して接続します。
アプリ、アシスタント、または MCP クライアントは、メモリ API キーを使用してリクエストを開きます。チャットボット、副操縦士、パイプラインに同じインターフェイス。
メモリにアクセスする前に検証します。
アクセス、クォータ、ポリシーは取得前に実行されます。機密フィールドはモデルに到達しません。
3 つのメモリ システムすべてから取得します。
セマンティック、エピソード、および手続き型レコードは、それぞれ最適化された独自のストアから並行して取得されます。
1 つのコンテキスト パッケージに統合します。
レコードは、引用と保持タグを備えた単一のコンパクトなコンテキストに再組み立てされます。モデルは重要なことだけを認識します。
応答は発信者に返されます。決定は次のリクエストのために取得され、スコープが設定され、追跡可能になります。
OctaMem に文書自体を渡します。契約書、資料、スプレッドシート、電子メール、PDF。私たちはそれを解析し、構造化し、エージェントが永久にクエリできる型付きメモリとして保存します。
BLOB の埋め込みではありません。条項、当事者、義務。
ファイルをドロップします。残りは記憶がやってくれます。
Contract-v3.pdf PDF マスターサービス契約の当事者・期間・義務 Contract-v3.pdf
マスターサービス契約、
v3 · 2026-04-12 に実行
› 関係者：Acme Corp、OctaMem Inc.
› 期間: 24 か月、自動更新 12 か月
› o

義務: 99.9% 稼働時間 SLA、30 日間の削除
Previous_context:legal-msas でアカウント全体を検索できます。
契約書、概要書、報告書、ナレッジベース。
テーブル、台帳、データセット — 行に基づいて推論されます。
デッキはスライドごとに解析されます。ピクセルではなく、事実を再現します。
スレッド、ペイロード、構造化エクスポート。
メモリのないセッションはすべてリセットされます。記憶を伴うすべてのセッションはアップグレードです。
名前、設定、初期制約。会話は少し個人的なものに感じられます。思慮深いインターンが初日からなんとかできるようなもの。
エージェントはあなたの決定を記憶し、過去の間違いを回避し、繰り返し指示することなくワークフローに従います。質問が減り、修正も減ります。
深い制度的背景。エージェントは、チーム、リリース、ツール間で継続的に動作します。 AI が実際に使用できる記録システム。
Day 360 はチャートに載っていません。カーブはどんどん登っていきます。
+ セッションごとにコンパウンド
一般的なクラウドで開始することも、セクターのスキーマ、ポリシー、コンプライアンス体制に合わせて調整された垂直固有のメモリ クラウドで実行することもできます。
P 手順のワークフローとルール
SEARCH PATH は 1 つのメモリ層を読み取ります。 S 意味論的な事実と知識 E エピソード的なイベントと履歴 P 手続き上のワークフローとルールは ADD PATH を書き込みます。 3 · 読み取りの検索、書き込みの追加、同じ 3 つのレイヤー。 01 / 事実と知識の意味論。
安定した知識、好み、アカウントの事実、ビジネス ルール、ドメイン コンテキスト。グラフ ストアを介した関係を意識した検索。
02 / イベントと歴史 E エピソード。
何が、いつ、そしてなぜ起こったのか。時間ベースの検索により、モデルに歴史の感覚が与えられ、順序付けられ、説明可能になります。
03 / ワークフローとルール P 手順。
仕事の進め方、エスカレーションパス、コンプライアンスワークフロー、繰り返し行われるルーチン。インテントによって取得されます。
あらゆるエージェント ワークフロー向けのセクターに依存しない永続メモリ。モデルに依存しない。プロトコル

-ネイティブ。
1 つのアカウントでのクロスモデル メモリ
高再現率の取得とコンテキストの再構築
意味記憶、エピソード記憶、手続き記憶のタイプ
ドメインを認識したメモリ構造、セクター固有の動作モデル、およびコンプライアンスに合わせたメモリ ポリシー。
一般的な Memory Cloud のすべて
ドメイン固有のメモリ スキーマと取得
セクターを意識した行動の継続性
ポリシーに基づく強制とガードレール
垂直方向に最適化されたコンテキストの再構築
優先サポートと導入オプション
ヘルスケア、法務、防衛のメモリ クラウドが開発中です。
同じメモリ層に、チームがすでに構築している方法でアクセスします。特注の垂直スタックはありません。書き換えはありません。プラットフォームはワークフローに合わせて形成されます。その逆ではありません。
医療、金融、防衛、公共部門。
既存のインターフェースからドロップインします。
訪問後の患者の継続性。ケアチームやセッション全体で持続する治療履歴。
事件の記憶と前例の追跡。案件全体にわたるクライアントとの対話の継続性。
ポートフォリオのコンテキスト、取引履歴、エージェントセッション全体で複雑になるリスク認識。
申し立ての履歴と政策の背景。あらゆるタッチポイントに保持されるアジャスターメモリー。
ブリーフィング、作戦、複数エージェントの調整にわたって持続するミッションのコンテキスト。
製品コンテキスト、顧客の成功履歴、チームやリリース全体に伝わるエンジニアリングの知識。
チャネル、倉庫、エージェント支援業務にわたる在庫、フルフィルメント、パートナーの記憶。
エネルギー、メディア、公共部門、その他の重要な分野 — 私たちはワークフローの周囲に記憶を形成します。
REST API ドキュメント → ダイレクト HTTP エンドポイント。 MCP や SDK なしで完全に制御できます。
MCP サーバー ドキュメント → リモート MCP — MCP 互換のアシスタントまたはエージェントから OctaMem を使用します。
Claude Desktop Docs → [設定] の [コネクタ]、または Ma 上のノードの構成ファイル

cとWindows。
[カーソル設定] の [カーソル ドキュメント] → [ツールと MCP]、または mcp.json — Mac、Windows、Linux。
Claude.ai (Web) ドキュメント → Web アプリのカスタム MCP コネクタ URL。
OpenClaw Docs → オープン エージェント スタックの自動呼び出しとキャプチャを備えたプラグイン。
Python SDK Docs → pip install octamem — スクリプトとサービスの型指定されたクライアント。
JavaScript SDK ドキュメント → npm install @octamem/octamem-js — Node、ブラウザー、Deno、および Bun 用。
一か八かのスタック向けに構築されています。
メモリの整合性が重要な場合、意思決定にトレーサビリティが必要な場合、継続性がオプションではない場合。 OctaMem は、セキュリティ、コンプライアンス、インフラストラクチャの各チームが実際に承認するレイヤーです。
エンタープライズ ポリシーと対話します ポリシー対応メモリ。
エージェントは、記憶層に組み込まれた組織のルール、制約、および境界を尊重します。プロンプトにもモデルにもありません。
ロールベースのアクセス · スコープ指定された取得 · テナントの分離
すべてのメモリの書き込みと読み取りは追跡可能です。ソースドキュメントからモデル出力までの完全な系統、規制された環境向けの暗号化の整合性。
不変の監査ログ · ソースの帰属 · 保持ポリシー
アクセス ロールベースのメモリ アクセス。
誰が何を見るかをチームが制御します。部門、プロジェクト、役割間のメモリ分離は、アプリケーションではなくストレージ層で強制されます。
SSO · SCIM プロビジョニング · 部門スコープ
クラウド、プライベート クラウド、またはオンプレミス。シングルテナント、専用キー、顧客管理の暗号化を最も危険な環境で利用できます。
クラウド · VPC · オンプレミス · BYO-KMS
完全な統合。動作するベクトル DB がありません。埋め込みパイプラインを維持する必要はありません。チャンク化はありません。 OctaMem はメモリを保持します。 Python、JavaScript、REST、または MCP のスタックを保持します。
> 追加（） 。以前のコンテキストを含む記憶をキャプチャします。
› get() / search() 。どのエージェント、どのセッションからでもそれを呼び出します。

› MCP 。 MCP 互換クライアントのツール呼び出しと同じ操作。
SDK リファレンス Python JavaScript REST API MCP サーバー MCP クライアント Quickstart.py · python octamem からコピー OctaMem をインポート
# platform.octamem.com からの API キー。
client = OctaMem(api_key="sk-om-live-...")
# 記憶をキャプチャします。
client.add(
content="ベータ版は 3 月 20 日にオープンします。",
previous_context="第 1 四半期の製品発表",
)
# 後で、おそらく別のエージェントから思い出してください。
結果 = client.get(
query="ベータ版はいつオープンしますか?",
previous_context="第 1 四半期の製品発表",
)
print(result) 応答 ·memory.search() 200 ·application/json {
「結果」:[
{ "id" : "rec_01HV4Z…" 、 "type" : "semantic" 、 "score" : 0.94 、
"content" : "ベータ版は 3 月 20 日にオープンします。" 、
"source" : "planning_doc_q1" , "created_at" : "2026-02-14T09:12Z" } ,
{ "id" : "rec_01HV7M…" 、 "type" : "エピソード" 、 "スコア" : 0.88 、
"content" : "2026 年 2 月 9 日に第 1 四半期の範囲縮小が承認されました。" } 、
{ "id" : "rec_01HV9F…" 、 "タイプ" : "プロシージャル" 、 "スコア" : 0.81 }
]、
"tokens" : 642 、 "previous_context" : "第 1 四半期の製品発売"
ソースにリンクされています。すべてのレコードには ID、タイプ、スコア、コンテンツ、ソースが含まれており、エンドツーエンドで監査可能で、 ID またはPrevious_context によって削除可能です。
単一領域の 3 つのメモリ層すべてにわたるエッジ キャッシュ検索。
最悪の場合、エンドツーエンド、コールド キャッシュ、再ランキングあり。
非同期埋め込みとインデックス作成前の同期確認。
記憶は敏感です。保存されている内容を確認し、構造化して追跡可能に保ち、いつでも必要に応じて削除できます。不透明な埋め込みはありません。ロックインされたベンダー形式はありません。
すべての記憶動作は痕跡を残します。
読み取り、書き込み、編集、およびポリシー チェックが連鎖的に行われるため、事後に記録を検査できます。
EVT_ 4182 前: 0b91ce774a リコール。要求されました
EVT_ 4183 前: 8f4a2c91b0 ポリシー.チェック済み
EVT_ 4184 前: 1c68bd044e context.delivered
EVT_

4185 前: ad72f9019c メモリ.キャプチャされました
エージェントが覚えている内容を正確に確認します。すべてのレコード、すべてのソース、すべての変更。ブラックボックスはありません。
記憶は型付けされ、タグ付けされ、追跡可能です。それぞれの記録には由来があります。埋め込みの塊ではありません。
レコードまたはprevious_contextによって、いつでもメモリを削除します。コマンドのことは忘れてください。
適切な管理を実施 · リクエストに応じてブリッジレターを発行
DPA が利用可能 · EU 西部での EU 居住権
管理が実施されている · 監査が進行中
保存中の AES-256-GCM · 転送中の TLS 1.3 · Enterprise 上の BYO-KMS
ロールベースのスコープ · Okta、Entra、Google 経由の SSO · SCIM
不変の追加専用ログ、レコードごとの来歴
構成可能なウィンドウ · レコードまたはコンテキストによる範囲指定された削除
テナントごとのメトリクス、レイテンシ、エラー バジェット · Datadog エクスポート
米国東部、EU 西、AP 南東。
Enterprise でのプライベート ネットワークのイングレス。
専用クラスター、専用キー。
コンテナー イメージ、顧客管理ストア。
データ処理に関する補遺 サブプロセッサ § 10 よくある質問
私たちはいつも得ます。
購入者が最初の会話で尋ねる最も一般的な 6 つの質問。あなたのメモがここにない場合は、メモを送ってください。追加させていただきます。
ベクトルデータベースはすでにこれを行っているのではないでしょうか？ベクトル データベースはインデックスです。 OctaMem はレコードです。私たちは、出所、監査、およびポリシーを含む型付きメモリを保持します。これらは、ベクター ストアがユーザーに任せるものです。ベクトル DB を保持できます。 p

[切り捨てられた]

## Original Extract

OctaMem is the persistent memory layer for AI systems. Semantic, episodic, and procedural memory that carries across sessions, models, and workflows.

OctaMem · Persistent memory for AI systems Skip to content OctaMem Home Platform Pricing Developers Trust About Docs Sign in Start free Navigate
Talk to sales Already have an account · Sign in Vol. 01 Memory infrastructure Live · status nominal Edition · 2026 An infrastructure for everything your agents shouldn’t forget.
The persistent memory layer for AI agents. Built for the stacks where forgetting isn’t an option.
Talk to sales No card. 2 GB on the free tier. Works with the stack you already run · OpenAI OpenAI ·
Agents without memory fail in two ways that show up on the balance sheet: they burn tokens re-reading context, and they let hard-won institutional knowledge walk out the door. A memory layer answers both.
Token spend compounds every turn.
Without memory, the same context is re-sent on every call. Conversations re-explain themselves, prompts balloon, and you pay frontier-model rates to re-read what the model was already told an hour ago.
Less context per call — only what's relevant is retrieved and injected
No repetition — facts and decisions persist instead of being re-sent
Cheaper models hold their own once the context they receive is sharper
Institutional knowledge isn't centralised.
What your agents and teams learn lives in scattered sessions, local notes, and individual heads. When an employee leaves, it leaves with them. Nothing compounds, and nothing is owned by the organisation.
Organisation-wide intelligence — every agent reads from one shared layer
Knowledge stays when people leave — it lives in the memory, not the person
Context compounds across teams instead of resetting every session
Memory has to be infrastructure — not a patch.
Every request passes through the same disciplined cycle. OctaMem doesn’t fire a generic search across one bucket of text. It rebuilds context from three memory types that each serve a distinct purpose, then reassembles them for the model.
Search cycle Add cycle Search cycle · in motion
02 / Access · quota Security layer
03 / OctaMem agent Retrieval service
04 / Three layers Memory layers
05 / Back to app Unified context
semantic episodic procedural 01 App or MCP Caller 02 Security layer Access · quota 03 Retrieval service OctaMem agent 04 Memory layers Three layers 05 Unified context Back to app fig. 1 · Search cycle · stage 1 of 5 Auto-playing Search cycle Add cycle fig. 1 · Search · stage 1 / 5 01 Connect through API, SDK, or MCP.
An app, assistant, or MCP client opens the request with a memory API key. Same interface for chatbots, copilots, and pipelines.
Validate before memory is touched.
Access, quota, and policy run before retrieval. Sensitive fields never reach the model.
Retrieve from all three memory systems.
Semantic, episodic, and procedural records are pulled in parallel, each from its own optimized store.
Unify into one context package.
Records are reassembled into a single compact context with citations and retention tags. The model only sees what matters.
Response goes back to the caller. The decision is captured for the next request, scoped and traceable.
Hand OctaMem the document itself. Contracts, decks, spreadsheets, emails, PDFs. We parse, structure, and store it as typed memory your agents can query forever.
Not embeddings of a blob. Clauses, parties, obligations.
Drop the file. Memory does the rest.
contract-v3.pdf PDF Master Services Agreement parties · term · obligations contract-v3.pdf
Master Services Agreement ,
v3 · executed 2026-04-12
› parties : Acme Corp, OctaMem Inc.
› term : 24 months, auto-renew 12
› obligations : 99.9% uptime SLA, 30-day deletion
Searchable across the account under previous_context: legal-msas .
Contracts, briefs, reports, knowledge bases.
Tables, ledgers, datasets — reasoned over rows.
Decks parsed slide by slide. Factual recall, not pixels.
Threads, payloads, structured exports.
Every session without memory is a reset. Every session with memory is an upgrade.
Names, preferences, initial constraints. Conversations feel slightly personalized. The kind a thoughtful intern manages on day one.
The agent remembers your decisions, avoids past mistakes, and follows your workflows without repeated instruction. Fewer questions, fewer corrections.
Deep institutional context. The agent operates with continuity across teams, releases, and tools. A system of record your AI can actually use.
Day 360 isn’t on the chart. The curve keeps climbing.
+ Compounds with every session
Start on the general cloud, or run on a vertical-specific memory cloud tuned to your sector’s schemas, policies, and compliance posture.
P procedural workflows & rules
SEARCH PATH reads ONE MEMORY LAYER S semantic facts & knowledge E episodic events & history P procedural workflows & rules writes ADD PATH fig. 3 · search reads, add writes, same three layers. 01 / FACTS & KNOWLEDGE S Semantic .
Stable knowledge, preferences, account facts, business rules, domain context. Relationship-aware retrieval over a graph store.
02 / EVENTS & HISTORY E Episodic .
What happened, when, and why. Time-based retrieval gives the model a sense of history, ordered and explainable.
03 / WORKFLOWS & RULES P Procedural .
How work should be done, escalation paths, compliance workflows, recurring routines. Retrieved by intent.
Sector-agnostic persistent memory for any agent workflow. Model-agnostic. Protocol-native.
Cross-model memory under one account
High-recall retrieval and context rebuild
Semantic, episodic, and procedural memory types
Domain-aware memory structures, sector-specific behavioural models, and compliance-aligned memory policies.
Everything in General Memory Cloud
Domain-specific memory schemas and retrieval
Sector-aware behavioural continuity
Policy-bound enforcement and guardrails
Vertical-optimized context rebuild
Priority support and deployment options
Healthcare, Legal, and Defense memory clouds in development.
The same memory layer, accessed however your team already builds. No bespoke vertical stack. No rewrite. The platform shapes to the workflow, not the other way around.
Healthcare, finance, defense, public sector.
Drop in through existing interfaces.
Patient continuity across visits. Treatment history that persists across care teams and sessions.
Case memory and precedent tracking. Client interaction continuity across matters.
Portfolio context, trade history, and risk awareness that compounds across agent sessions.
Claims history and policy context. Adjuster memory that carries across every touchpoint.
Mission context that persists across briefings, operations, and multi-agent coordination.
Product context, customer success history, and engineering knowledge that carries across teams and releases.
Inventory, fulfillment, and partner memory across channels, warehouses, and agent-assisted operations.
Energy, media, public sector, and other high-stakes domains — we'll shape memory around your workflows.
REST API Docs → Direct HTTP endpoints. Full control without MCP or an SDK.
MCP Server Docs → Remote MCP — use OctaMem from any MCP-compatible assistant or agent.
Claude Desktop Docs → Connectors in Settings, or config file with Node on Mac and Windows.
Cursor Docs → Tools & MCP in Cursor settings, or mcp.json — Mac, Windows, and Linux.
Claude.ai (Web) Docs → Custom MCP connector URL in the web app.
OpenClaw Docs → Plugin with auto-recall and capture for open agent stacks.
Python SDK Docs → pip install octamem — typed client for scripts and services.
JavaScript SDK Docs → npm install @octamem/octamem-js — for Node, browsers, Deno, and Bun.
Built for the high-stakes stack.
When memory integrity matters, when decisions need traceability, when continuity is not optional. OctaMem is the layer your security, compliance, and infrastructure teams will actually approve.
Talk to enterprise Policy Policy-aware memory.
Agents respect organizational rules, constraints, and boundaries embedded in the memory layer. Not in the prompt, not in the model.
Role-based access · Scoped retrieval · Tenant isolation
Every memory write and read is traceable. Full lineage from source document to model output, with cryptographic integrity for regulated environments.
Immutable audit log · Source attribution · Retention policies
Access Role-based memory access.
Teams control who sees what. Memory isolation between departments, projects, and roles, enforced at the storage layer, not the application.
SSO · SCIM provisioning · Department scopes
Cloud, private cloud, or on-premise. Single-tenant, dedicated keys, and customer-managed encryption available for the highest-stakes environments.
Cloud · VPC · On-prem · BYO-KMS
The full integration. No vector DB to operate. No embedding pipeline to maintain. No chunking. OctaMem holds the memory; you keep your stack — Python, JavaScript, REST, or MCP.
› add() . Capture a memory with its previous context.
› get() / search() . Recall it from any agent, any session.
› MCP . Same operations as tool-calls in any MCP-compatible client.
SDK reference Python JavaScript REST API MCP Server MCP Client quickstart.py · python Copy from octamem import OctaMem
# Your API key from platform.octamem.com.
client = OctaMem(api_key="sk-om-live-...")
# Capture a memory.
client.add(
content="Beta opens March 20.",
previous_context="Q1 product launch",
)
# Recall it later, possibly from a different agent.
results = client.get(
query="When does beta open?",
previous_context="Q1 product launch",
)
print(results) response · memory.search() 200 · application/json {
"results" : [
{ "id" : "rec_01HV4Z…" , "type" : "semantic" , "score" : 0.94 ,
"content" : "Beta opens March 20." ,
"source" : "planning_doc_q1" , "created_at" : "2026-02-14T09:12Z" } ,
{ "id" : "rec_01HV7M…" , "type" : "episodic" , "score" : 0.88 ,
"content" : "Approved Q1 scope reduction on 2026-02-09." } ,
{ "id" : "rec_01HV9F…" , "type" : "procedural" , "score" : 0.81 }
],
"tokens" : 642 , "previous_context" : "Q1 product launch"
} Source-linked. Every record carries id, type, score, content, and source — auditable end-to-end, deletable by id or by previous_context .
Edge-cached search across all three memory layers, single region.
Worst-case end-to-end, cold cache, with reranking.
Synchronous acknowledgement before async embedding & indexing.
Memory is sensitive. See what is stored, keep it structured and traceable, and delete it whenever you want. No opaque embeddings. No locked-in vendor format.
Every memory action leaves a mark .
Reads, writes, redactions, and policy checks are chained together so the record can be inspected after the fact.
EVT_ 4182 prev: 0b91ce774a recall.requested
EVT_ 4183 prev: 8f4a2c91b0 policy.checked
EVT_ 4184 prev: 1c68bd044e context.delivered
EVT_ 4185 prev: ad72f9019c memory.captured
See exactly what your agents remember. Every record, every source, every change. No black box.
Memory is typed, tagged, and traceable. Each record carries provenance. Not a blob of embeddings.
Remove any memory at any time, by record or by previous_context. Forget on command.
Controls in place · bridge letter on request
DPA available · EU residency in eu-west
Controls in place · audit in progress
AES-256-GCM at rest · TLS 1.3 in transit · BYO-KMS on Enterprise
Role-based scopes · SSO via Okta, Entra, Google · SCIM
Immutable append-only log · per-record provenance
Configurable windows · scoped delete by record or context
Per-tenant metrics, latency, error budgets · Datadog export
us-east, eu-west, ap-southeast.
Private network ingress on Enterprise.
Dedicated cluster, dedicated keys.
Container image, customer-managed stores.
Data processing addendum Subprocessors § 10 Frequently asked The questions
we always get.
Six of the most common things buyers ask in the first conversation. If yours isn’t here, send us a note and we’ll add it.
Doesn't a vector database already do this? A vector database is an index . OctaMem is a record . We hold typed memory with provenance, audit, and policy — the things a vector store leaves to you. You can keep your vector DB; p

[truncated]

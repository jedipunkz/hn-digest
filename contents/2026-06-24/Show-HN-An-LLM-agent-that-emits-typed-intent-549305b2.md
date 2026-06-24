---
source: "https://github.com/gabert/ontocortex"
hn_url: "https://news.ycombinator.com/item?id=48665700"
title: "Show HN: An LLM agent that emits typed intent"
article_title: "GitHub - gabert/ontocortex: Intent-execution separation for LLM agents: the model emits typed intent (SIF) in a domain vocabulary; a deterministic layer validates, scopes, translates, and executes it. Java/Spring + Vue, with two demo domains (legal, vet). · GitHub"
author: "gabert"
captured_at: "2026-06-24T21:29:52Z"
capture_tool: "hn-digest"
hn_id: 48665700
score: 1
comments: 0
posted_at: "2026-06-24T21:12:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An LLM agent that emits typed intent

- HN: [48665700](https://news.ycombinator.com/item?id=48665700)
- Source: [github.com](https://github.com/gabert/ontocortex)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T21:12:04Z

## Translation

タイトル: HN を表示: 型付きインテントを発行する LLM エージェント
記事のタイトル: GitHub - gabert/ontocortex: LLM エージェントのインテントと実行の分離: モデルはドメイン ボキャブラリで型付きインテント (SIF) を発行します。決定論的レイヤーがそれを検証、範囲指定、変換、および実行します。 Java/Spring + Vue、2 つのデモ ドメイン (法律、獣医師)。 · GitHub
説明: LLM エージェントのインテントと実行の分離: モデルはドメイン ボキャブラリで型付きインテント (SIF) を発行します。決定論的レイヤーがそれを検証、範囲指定、変換、および実行します。 Java/Spring + Vue、2 つのデモ ドメイン (法律、獣医師)。 - ガベール/オントコーテックス

記事本文:
GitHub - gabert/ontocortex: LLM エージェントのインテントと実行の分離: モデルはドメイン ボキャブラリで型付きインテント (SIF) を発行します。決定論的レイヤーがそれを検証、範囲指定、変換、および実行します。 Java/Spring + Vue、2 つのデモ ドメイン (法律、獣医師)。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
ディス

ミスアラート
{{ メッセージ }}
ガベール
/
皮質上
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
82 コミット 82 コミット docker/ postgres-init docker/ postgres-init docs docs ドメイン ドメイン エンジン エンジン フロントエンド フロントエンド .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
これは何だろう。パターンのリファレンス実装: LLM エージェントのインテントと実行の分離 — モデルはドメイン語彙で型指定されたインテントを出力し、基板には決して触れませんが、決定論的レイヤーがそれを検証、変換、スコープ指定し、実行します。プロジェクトは、意図的な優先順位に従って 2 つの成果物を出荷します。
プライマリ — SIF と、それによって具体化される意図と実行の分離パターン。型指定されたインテント コントラクトと、その下の検証、変換、範囲設定、および実行を行う決定層。耐荷重への貢献。リポジトリの残りの部分は、動作する実行中のインスタンスを提供するために存在します。 SIF がコアである を参照してください。
セカンダリ — 実験的な「宣言型コンテンツとしてのアプリケーション」アプローチ。手書きのサービス層を持たない、小さなdomain.json記述子を介して、オントロジー(プレーンYAMLの型付きドメインモデル)、ビジネスルールのセット(エージェントが読み取るプレーンテキスト)、およびエージェントプロンプト(ペルソナ+インタラクションプロトコル)の3つの宣言型アーティファクトからドメインごとに組み立てられた、動作するエージェント駆動のビジネスアプリケーション。本当に有用なアプリケーションが命令型コードではなく「オントロジー + ルール + プロンプト」から生み出されるかどうかは未解決の問題です。法的ドメインと獣医ドメインは、

提供された証拠。これは SIF に基づいて構築されており、意図的に優先度が低くなります。詳細: ビジネス ロジックはどこに存在しますか? 。
パターンの完全なステートメント: docs/design-intent-principle.md 。
実験的で非学術的な研究プロジェクト。オントロジー駆動アーキテクチャのビジネス ロジック層として LLM ができることの限界を探ります。ここでのすべての仮説は、このリポジトリに同梱されているサンプル アプリケーションに対して経験的にテストされます。合法的なデモ ドメインは、ショーケース デモとしてではなく、実験的なハーネスとして機能します。議論のみで到達し、それを裏付ける実行例がなければ、結論はカウントされません。
これの読み方。ファイル レイアウト、デモ ドメイン、REST サーフェス、エージェントの実際の動作など、実行中の設計を確認したい場合は、この README を読み続けて、domains/ フォルダーを調べてください。設計の理論的根拠 (ツール表面の問題、語彙規律、ステート マシンが存在する場所に存在する理由) が必要な場合は、 docs/design-sif.md を読んでください。どちらの図も同じシステムを反対側から説明しています。
SIFが核
ルーティングの具体的なイメージ — 1 つのオントロジー、2 つのアダプター
「これは余分な手順を追加した単なる SQL ではないでしょうか?」
実際の動作を確認してください
1. 1 つの質問、2 つのデータベース
2. 1 つの命令、1 つのトランザクション
3. エージェントが提案し、エンジンが決定する
ビジネスアプリケーションを超えて
認可が適している場所
ビジネスロジックはどこにあるのでしょうか?
4. 各ドメインのデータベースをプロビジョニングする
SIF (Structured Intent Format) は、オントロジー エンティティに対する型付きインテント コントラクトです。これは、このプロジェクトの負担がかかる部分です。他のすべてはそれを中心に構築されます。
LLM は、アクティブなドメインのオントロジー語彙 (クラス、プロパティ、関係、遷移、値セット) のみを使用して、構造化インテント ( find 、 create 、 update 、 delete 、 link 、 unlink 、transition ) を発行します。 SQL は決して見ません、物理学

l テーブル名、ユーザー ID、またはストレージ プリミティブ。
フレームワークは、オントロジーに対してすべてのインテントを検証し、それを満たせるアダプターにルーティングします。アダプターはフェデレーション SPI (DataSource インターフェイス) の背後に存在し、ストレージ プリミティブではなくオントロジー タイプを話します。
アダプターには、SQL、ドキュメント、キーと値、REST、内部メソッド呼び出し、ベクター ストアなど、オントロジー レベルの意図を満たすことができるものであれば何でもかまいません。 LLM は、どのアダプタが特定のエンティティにサービスを提供したかを決して知りませんし、知る必要もありません。
現在、PostgreSQL 上の SqlDataSource (完全な CRUD と遷移) と MongoDB 上の MongoDataSource ( find と create 。残りの書き込み動詞 (更新 / 削除 / リンク / リンク解除 / 遷移) は計画されており、現在は回復可能な「サポートされていない」を返します) の 2 つのアダプタが出荷されています。合法的なデモでは、Mongo 読み取りパスを実行します。ストレージ層は、SIF コントラクトの背後にあるプライベート実装の詳細であり、フレームワークの ID ではありません。メソッド呼び出し、REST、またはベクター ストア アダプターを追加するには、新しい DataSource 実装を作成する必要があります。SIF サーフェス、オントロジー、および LLM プロンプトは変更されません。
この分離によって、一度に 2 つのものが得られます。
信頼の逆転。 LLM は自然言語インテント ルーターです。フレームワークは信頼できる実行者です。 LLM は、JSON スキーマ境界の SIF 文法によってサンドボックス化されています。LLM は、宣言された動詞とオントロジー名のみを表現でき、SQL、シェル、URL、またはストレージ プリミティブは表現できません。そのため、プロンプト挿入モデルや幻覚モデルは、基板に直接到達するための表現可能な方法がありません。認可は別の問題です。SIF サーフェスの下にあるものはすべて決定論的に実行されるため、LLM の後に認可を強制する場所が 1 か所あります。デモには、意図的に単純な所有者スコープのサンプルが同梱されています。実際の RBAC/IdP は同じ継ぎ目で接続されます (著者の場所を参照)

化は以下に適合します)。これは認可モデルではありません。
バックエンドの中立性。 LLM は必要なものを表現します。フレームワークがそれを実現する方法を決定します。エンティティが SQL テーブルから REST エンドポイントの背後にあるサービス メソッドに移動するときに、プロンプトの変更、オントロジーの変更、SIF の変更はありません。
プロジェクトのコンパクトな YAML 形式の単一のオントロジー ファイルが、各ドメインの信頼できる情報源となります。 「オントロジー」は全体を通して非公式の略語です。語彙はOWLからインスピレーションを得ていますが、提供されるのは単純なYAMLの型付きドメイン・モデルであり、厳密なセマンティックWebの意味でのOWLオントロジーではありません。推論もSPARQLも等価性/制限公理も、subClassOf駆動の推論もありません。その伝統から何が守られ、何が削除されるかについては、 docs/design-sif.md (「単語オントロジーについて」) を参照してください。
2 つのドメインが出荷されます。法的 — 問題、紛争チェック、文書、請求可能期間、請求書発行、審問をカバーする法律事務所のマターマネージャー — は、豊富な概念実証です。 vet — 動物病院 (飼い主、ペット、予約、訪問、治療、請求) — は意図的にシンプルな SQL のみの 2 番目のドメインです。同じエンジンが異なるオントロジーを提供し、宣言的なコンテンツのみが交換されます。どちらも、ドメイン全体を監視する単一の管理 ID を実行します。
ルーティングの具体的なイメージ — 1 つのオントロジー、2 つのアダプター
配送法的デモでは、1 つのオントロジーを 2 つの基板にルーティングします。ほとんどのエンティティ (弁護士、パラリーガル、クライアント、Matter、TimeEntry、Invoice、Hearing、ConflictCheck) は SqlDataSource を通じて PostgreSQL に存在します。 MatterNote エンティティ (案件に添付された自由形式の作業メモ) は、 MongoDataSource を通じて MongoDB に存在します。 LLM は 1 つの統合オントロジーを認識し、両方に対して同じ形状の SIF インテントを発行します。
{
「操作」: [
{ "op" : " 検索 " 、 "entity" : " 弁護士 " 、
"フィルター" : { "barNumb

えー" : " SBA-2026-1042 " } }、
{ "op" : " find " 、 "entity" : " MatterNote " 、
"フィルター" : { "作成日" : " 2026-06-12 " } }
】
}
フレームワークがそれをディスパッチする方法:
LLM は、あるエンティティがリレーショナル ストアに存在し、別のエンティティがドキュメント ストアに存在することを決して認識しません。両方に対して同じ形の意図を発します。フレームワークは、 DataSourceRegistry.findFor(entity) を介して、解決された op のエンティティ クラスごとにアダプターを選択します。明日、3 番目のエンティティがメソッド呼び出しアダプターまたは REST エンドポイントに移動したとしても、アダプターの実装のみが変更され、SIF インテント、オントロジー、および LLM プロンプトは同一のままになります。
構成はドメインごとに行われます。法的デモの domain.json は 2 つのデータ ソース (プライマリ: リレーショナル エンティティの SQL、ドキュメント: MatterNote の mongo) を宣言し、アダプターごとのマッピング ファイル (mapping.yaml 、mapping.mongo.yaml) が各クラスをその基板にバインドします。 3 番目のデータ ソースの追加は、構成の追加と 3 番目のマッピング ファイルです。SIF サーフェス、LLM プロンプト、および残りのオントロジーは変更されません。
これはフェデレーション プリミティブです。「フレームワークは Postgres または Mongo と通信できる」のではなく、「フレームワークは、LLM に対して透過的に、同時に複数の基板から単一のオントロジーを提供します」。
これが価値のストーリーに当てはまるところ。異種混合の場合は、オントロジーが最も目に見えて利益を得る場所、つまり異種の基質上で統一された意味論的語彙が得られる場合です。オントロジーは、単一アダプター内でも重要な役割を果たします (型付き LLM ツール サーフェス、実行前検証、ライフサイクル保護、クロスドメイン プロンプトの再利用、監査フレーミング) — docs/design-sif.md を参照してください。
LLM エージェントは、少数のツールでは信頼できますが、数が増えると信頼性が低くなります。新しいツールはそれぞれ、モデルが誤って選択する可能性のある別の選択肢であり、モデルが幻覚を引き起こす可能性のある別の引数のセットです。 SIF はエージェントの p を保持します

ドメインが大きくなっても、選択肢は小さくなります。7 つの op 動詞と、エンティティ、プロパティ、関係、遷移などの型付き名前のオントロジーが列挙型としてツール スキーマに挿入されます。エージェントは、小さな型指定されたグリッド内のスロットを選択します。広大なツール カタログをナビゲートすることはできません。これがアーキテクチャ全体の基礎となる原則であり、LLM はエグゼキュータではなくインテント ルーターであり、 docs/design-intent-principle.md で完全に扱われます。
すでに説明した 2 つのメリット (信頼の逆転とバックエンドの中立性) を超えて、機能ごとの作業を行わずに、同じ継ぎ目が得られます。
監査と再生。すべての操作は構造化され、人間が判読できるインテント レコードです。ディスク上のセッション ログは監査証跡です。
会話の原子性。 1 つのユーザー意図 (「連帯保証人を設定し、新しいローンにリンクする」) = 1 つの submit_sif バッチ = アダプターでの 1 つのトランザクション。ユースケースごとに特注のサガ/補償コードはありません。
リファクタリングセーフ。物理スキーマの変更 (またはテーブルからサービス メソッドへのエンティティの移動) は、LLM プロンプトではなくアダプター層に影響します。同じオントロジーは、複数の物理スキーマ (例: 分離された DB を持つマルチテナント) または混合バックエンドをターゲットにすることができます。
デバッグ可能。何かが壊れると、トレースは LLM の元の意図が変更されていないことを示します。バグは内部のいずれかにあります。

[切り捨てられた]

## Original Extract

Intent-execution separation for LLM agents: the model emits typed intent (SIF) in a domain vocabulary; a deterministic layer validates, scopes, translates, and executes it. Java/Spring + Vue, with two demo domains (legal, vet). - gabert/ontocortex

GitHub - gabert/ontocortex: Intent-execution separation for LLM agents: the model emits typed intent (SIF) in a domain vocabulary; a deterministic layer validates, scopes, translates, and executes it. Java/Spring + Vue, with two demo domains (legal, vet). · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
gabert
/
ontocortex
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
82 Commits 82 Commits docker/ postgres-init docker/ postgres-init docs docs domains domains engine engine frontend frontend .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml View all files Repository files navigation
What this is. A reference implementation of a pattern : intent-execution separation for LLM agents — the model emits typed intent in domain vocabulary and never touches the substrate, while a deterministic layer validates, translates, scopes, and executes it. The project ships two deliverables, in deliberate priority order:
Primary — SIF and the intent-execution separation pattern it makes concrete. The typed intent contract, and the deterministic layer beneath it that validates, translates, scopes, and executes. The load-bearing contribution; the rest of the repo exists to give it a worked, running instance. See SIF is the core .
Secondary — the experimental "application as declarative content" approach. A working, agent-driven business application assembled per domain from three declarative artifacts — an ontology (a typed domain model in plain YAML), a set of business rules (plain text the agent reads), and an agent prompt (persona + interaction protocol) — over a small domain.json descriptor, with no hand-written service layer . Whether a genuinely useful application can emerge from "ontology + rules + prompt" instead of imperative code is the open question; the legal and vet domains are the evidence offered. This builds on SIF and is deliberately lower-priority. Detail: Where does the business logic live? .
Full statement of the pattern: docs/design-intent-principle.md .
Experimental, non-academic research project. Explores the boundaries of what LLMs can do as a business logic layer in ontology-driven architectures. Every hypothesis here is tested empirically against the example application shipped in this repo — the legal demo domain serves as the experimental harness, not as a showcase demo. Conclusions reached only by argument, without a running example to back them up, don't count.
How to read this. If you want to see the design running — file layout, demo domains, REST surface, what the agent does in practice — keep reading this README and explore the domains/ folder. If you want the design rationale — the tool-surface problem, the vocabulary discipline, why state machines live where they do — read docs/design-sif.md . Both views describe the same system from opposite ends.
SIF is the core
A concrete picture of routing — one ontology, two adapters
"Isn't this just SQL with extra steps?"
See it in action
1. One question, two databases
2. One instruction, one transaction
3. The agent proposes, the engine decides
Beyond business applications
Where authorization fits
Where does the business logic live?
4. Provision each domain's database
SIF — Structured Intent Format — is a typed intent contract over ontology entities. It is the load-bearing piece of this project; everything else is built around it.
The LLM emits structured intents ( find , create , update , delete , link , unlink , transition ) using only the ontology vocabulary of the active domain — classes, properties, relations, transitions, value sets. It never sees SQL, physical table names, user IDs, or any storage primitive.
The framework validates every intent against the ontology, then routes it to whichever adapter can fulfil it. Adapters live behind the federation SPI ( DataSource interface) and speak ontology types, not storage primitives.
An adapter can be SQL, document, key-value, REST, an internal method call, a vector store — anything that can fulfil ontology-level intents. The LLM never knows — and need not know — which adapter served a given entity.
Today two adapters ship: SqlDataSource on PostgreSQL (full CRUD plus transitions) and MongoDataSource on MongoDB ( find and create ; the remaining write verbs — update / delete / link / unlink / transition — are planned and currently return a recoverable "not supported"). The legal demo exercises the Mongo read path. The storage layer is a private implementation detail behind the SIF contract, not the framework's identity. Adding a method-call, REST, or vector-store adapter is a matter of writing a new DataSource implementation — the SIF surface, the ontology, and the LLM prompt stay unchanged.
That separation buys two things at once:
Trust inversion. The LLM is the natural-language intent router; the framework is the trusted executor. The LLM is sandboxed by the JSON-Schema-bounded SIF grammar — it can express only declared verbs and ontology names, never SQL, shell, URLs, or any storage primitive — so a prompt-injected or hallucinating model has no expressible way to reach the substrate directly. Authorization is a separate concern: because everything below the SIF surface runs deterministically, there is one place after the LLM to enforce it. The demo ships a deliberately simple owner-scoping example there; a real RBAC/IdP plugs in at the same seam (see Where authorization fits below). It is not an authorization model.
Backend neutrality. The LLM expresses what it wants; the framework decides how to fulfil it. No prompt change, no ontology change, no SIF change when an entity moves from a SQL table to a service method behind a REST endpoint.
A single ontology file in the project's compact YAML format is the source of truth for each domain. "Ontology" is informal shorthand throughout: the vocabulary is OWL-inspired, but what ships is a typed domain model in plain YAML, not an OWL ontology in the strict semantic-web sense — no reasoning, no SPARQL, no equivalence/restriction axioms, no subClassOf -driven inference. See docs/design-sif.md ("About the word ontology ") for what is kept from that tradition and what is dropped.
Two domains ship. legal — a law-firm matter manager covering matters, conflict checks, documents, billable time, invoicing and hearings — is the rich proof-of-concept. vet — a veterinary clinic (owners, pets, appointments, visits, treatments, billing) — is a deliberately simple, SQL-only second domain: the same engine serving a different ontology, with nothing but declarative content swapped. Both run a single managing identity that sees the whole domain.
A concrete picture of routing — one ontology, two adapters
The shipping legal demo routes one ontology to two substrates. Most entities — Lawyer , Paralegal , Client , Matter , TimeEntry , Invoice , Hearing , ConflictCheck — live in PostgreSQL through SqlDataSource . The MatterNote entity — free-form working memos attached to matters — lives in MongoDB through MongoDataSource . The LLM sees one unified ontology and emits the same shape of SIF intent for both:
{
"operations" : [
{ "op" : " find " , "entity" : " Lawyer " ,
"filters" : { "barNumber" : " SBA-2026-1042 " } },
{ "op" : " find " , "entity" : " MatterNote " ,
"filters" : { "createdDate" : " 2026-06-12 " } }
]
}
How the framework dispatches it:
The LLM never knows one entity lives in a relational store and another in a document store. It emits the same shape of intent for both. The framework picks the adapter per resolved op's entity class via DataSourceRegistry.findFor(entity) . If tomorrow a third entity moved to a method-call adapter or a REST endpoint, only the adapter implementation would change — the SIF intent, the ontology, and the LLM prompt would stay identical.
Configuration is per-domain. The legal demo's domain.json declares two data sources ( primary: sql for the relational entities, documents: mongo for MatterNote ), and per-adapter mapping files ( mapping.yaml , mapping.mongo.yaml ) bind each class to its substrate. Adding a third data source is a config addition plus a third mapping file — the SIF surface, the LLM prompt, and the rest of the ontology stay unchanged.
That is the federation primitive: not "the framework can talk to Postgres or Mongo," but "the framework serves a single ontology from multiple substrates at the same time, transparently to the LLM."
Where this fits in the value story. The heterogeneous case is where ontology most visibly earns its keep — a unified semantic vocabulary above disparate substrates. Ontology also pulls weight even within a single adapter (typed LLM tool surface, pre-execution validation, lifecycle protection, cross-domain prompt reuse, audit framing) — see docs/design-sif.md .
LLM agents are reliable with a handful of tools and less so as the count grows — each new tool is another choice the model can pick wrong on, another set of arguments it can hallucinate. SIF keeps the agent's per-turn choice small even as the domain gets bigger: seven op verbs and an ontology of typed names — entities, properties, relationships, transitions — injected into the tool schema as enums. The agent picks slots inside a small, typed grid; it doesn't navigate a sprawling tool catalog. That is the principle the whole architecture rests on — the LLM is an intent router, not an executor — and it is given its full treatment in docs/design-intent-principle.md .
Beyond the two payoffs already described — trust inversion and backend neutrality — the same seam yields, with no per-feature work:
Audit & replay. Every operation is a structured, human-readable intent record. The on-disk session log is the audit trail.
Conversational atomicity. One user intention ("set up the cosigner and link her to my new loan") = one submit_sif batch = one transaction at the adapter. No bespoke saga/compensation code per use case.
Refactor-safe. Physical schema changes (or moving an entity from a table to a service method) touch the adapter layer, not the LLM prompt. The same ontology can target multiple physical schemas (e.g. multi-tenant with isolated DBs) or mixed backends.
Debuggable. When something breaks, the trace shows the LLM's original intent unchanged — the bug is either in the inte

[truncated]

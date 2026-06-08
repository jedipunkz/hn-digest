---
source: "https://github.com/wallidsaydi-creator/hom-local"
hn_url: "https://news.ycombinator.com/item?id=48443074"
title: "HOM Local- a memory kernel for AI agents with audit trail and source attribution"
article_title: "GitHub - wallidsaydi-creator/hom-local: HOM Local — open-source local-first AI memory kernel · GitHub"
author: "walldad2"
captured_at: "2026-06-08T10:41:26Z"
capture_tool: "hn-digest"
hn_id: 48443074
score: 1
comments: 1
posted_at: "2026-06-08T09:27:13Z"
tags:
  - hacker-news
  - translated
---

# HOM Local- a memory kernel for AI agents with audit trail and source attribution

- HN: [48443074](https://news.ycombinator.com/item?id=48443074)
- Source: [github.com](https://github.com/wallidsaydi-creator/hom-local)
- Score: 1
- Comments: 1
- Posted: 2026-06-08T09:27:13Z

## Translation

タイトル: HOM Local - 監査証跡とソース帰属を備えた AI エージェント用のメモリ カーネル
記事タイトル: GitHub - Wallidsaydi-creator/hom-local: HOM Local — オープンソースのローカルファースト AI メモリ カーネル · GitHub
説明: HOM Local — オープンソースのローカルファースト AI メモリ カーネル - Wallidsaydi-creator/hom-local

記事本文:
GitHub - Wallidsaydi-creator/hom-local: HOM Local — オープンソースのローカルファースト AI メモリ カーネル · GitHub
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
アラートを閉じる
{{ メッセージ }}
ウォリセイディクリエイター
/
地元の
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github crates crates docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス通知 NOTICE README.md README.md SECURITY.md SECURITY.mddeny.tomldeny.toml release-sanitize.sh release-sanitize.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントとハーネス用のローカル ファースト メモリ サーバー。
セッション、モデル変更、プロバイダー スワップにわたる作業を記憶する必要がある AI システムを構築する場合、HOM Local は来歴、ゲート、継続性を備えた永続メモリ コアを提供します。
これによりチームの時間が節約できる場合は、より多くのビルダーがリポジトリを見つけられるように、リポジトリにスターを付けてください。
プロバイダーの境界とモデルの選択
メモリアーティファクトとしてのポストコンパクション
ほとんどの AI スタックは、セッションが終了したり、モデルが変更されたり、ハンドラーが再起動されたりするたびにコンテキストを失います。これにより、決定、制約、進捗状況を何度も繰り返し説明する再説明税が発生します。
HOM Local は、チャット面ではなくメモリ層でこの問題を修正します。
記憶を出所と結び付けておく
プロバイダーとルーティングの選択全体での継続性の維持
HOM Local は、バックエンドの一部としてメモリを重視する次のようなチーム向けです。
長いワークフロー全体でコンテキストを必要とする AI の副操縦士と活用。
アクションが監査可能で回復可能である必要がある、ツールが豊富なエージェント。
Claude/GPT/Gemini などから切り替えることができるマルチモデル スタック。記憶を書き換えることなく。
ローカルファーストの展開と明示的な信頼境界を必要とするチーム。
ソースの帰属を伴う構造化された記憶により、エージェントは「何を知っていますか?」と尋ねることができます。出所を意識した回答パスを取得します。
メモリ書き込みパスの品質ゲート: 構造、関連性、実質、事実。
突然変異 l

トレーサビリティと変更の可視性を実現するために、改ざんが明らかなローカル履歴をエッジリングします。
セッションの圧縮は、単なるプロンプトトリミングではありません。これにより、後の実行でも耐久性のある継続性が得られます。
HOM Local は専用のメモリ サービスとして設計されています。
3 クレートのワークスペース: hom-brain 、 hom-ingress 、 hom-shared 。
境界を明確にする: イングレスは HTTP/認証を処理します。脳は記憶の意味論を処理します。共有はプロトコルと暗号コントラクトを保持します。
確定的 IPC + HTTP エッジ: 入力に HTTP 境界があるブレイン パス内の Unix ドメイン ソケット上の JSON-RPC。
同時実行設計: 保存、呼び出し、認識、I/O、およびメイン操作の重み付け優先スケジューリング。
来歴第一: 追加専用台帳とソースにリンクされたメモリ アーティファクト。
プロバイダー中立のメモリ サーフェス: モデル プロバイダーの変更にはメモリの書き換えは必要ありません。
クライアント/エージェント → ホムイングレス (認証/ルーティング) → ホムブレイン (メモリ + ゲート + 台帳) → SQLite WAL + ベクトルインデックス
何が違うのか
HOMローカル
一般的なメモリの統合
ローカルファーストのデータ制御
✅
変数
出典に起因するリコール
✅
不透明なことが多い
プロバイダーに依存しないメモリ
✅
しばしば結合される
改ざん防止メモリ履歴
✅
利用できないことが多い
圧縮から継続へのパイプライン
✅
多くの場合プロンプトのみ
品質ゲート書き込み
✅
多くの場合、単純であるか、存在しない
アプリ制御のためのオープンアーキテクチャ
✅
制限されることが多い
1 つのチャット UI を超えるメモリが必要な場合、これがテストする主な理由です。
HOM Local は、モデル プロバイダーをアプリ層に維持しながら、アプリに耐久性のあるメモリ API を提供します。
永続的なリコール状態を持つエージェント ハーネスを構築します。
プロバイダーを切り替えたり、長時間のセッションを実行したりするときにコンテキストを保持します。
アプリの出力、証拠、モデルの応答をメモリ レコードにルーティングします。
アドホックなサマリーではなく、圧縮アーティファクトから継続性を意識したワークフローを実行します。
✅ WAL モードを使用した SQLite の耐久性のあるローカル メモリ
✅ 出典に基づく

品質スコアリングによるリコール
✅ 4 つの壁で品質ゲートされた記憶保存 (構造、関連性、実質、事実)
✅ ハッシュチェーン検証を備えた改ざん防止追加専用台帳
✅ 証拠カードとオープンハンドルによるコンテキストパッキング
✅ 耐久性と検査可能な状態を維持するセッション圧縮アーティファクト
✅ 正確なリランクフォールバックによる積量子化ベクトルリコール
✅ Ed25519 エンベロープ認証
✅ Unix ドメインソケット経由の JSON-RPC + 127.0.0.1:9101 の HTTP イングレス
✅ テスト、例、およびオペレーターのガイダンス
境界のあるウィンドウによる時間的想起
リコール品質とカバレッジ診断
ドリフト防止のチェックと施行
高品質のゲート + 監査証跡保護
製品固有のエージェント ハーネスの代替品
これは、アプリを構成できるバックエンドのメモリ脳です。
プロバイダーの境界とモデルの選択
HOM Local はモデルプロバイダーを所有していません。
アプリはプロバイダーとランタイムの動作を決定します。
HOM Local は、メモリ、リコール、検証、圧縮、および監査を所有します。
これは意図的に分離されているため、プロバイダー戦略が進化してもメモリは安定したままになります。
メモリアーティファクトとしてのポストコンパクション
ほとんどのシステムは、長いコンテキストを一時的な要約に変換して次に進みます。
HOM Local は、圧縮後のサマリーを、ソース セッションおよびメモリへのリンクを含む継続性アーティファクトとして保持します。
したがって、連続性は破棄されるのではなく、回復可能、検査可能、および再利用可能です。
git clone https://github.com/wallidsaydi-creator/hom-local.git
cd ホームローカル
カーゴビルド --release
各サービスを独自のターミナルで実行します。
カーゴラン --release --bin hom-brain
カーゴ実行 --release --bin hom-ingress
イングレスは http://127.0.0.1:9101 で利用できるはずです。
カーゴインストールホムブレイン
建築
ホームローカル/
§── 木箱/
│ §── hom-brain/ # コアブレインデーモン — メモリ、台帳、リコール、品質ゲート

│ §── hom-shared/ # 共有タイプ、暗号、エンベロープ、RPC、パス
│ └── hom-ingress/ # 認証のある HTTP 入力層
§── docs/ # アーキテクチャと API のドキュメント
└── 例/ # 使用例
完全なサービス グラフについては、Oracle アーキテクチャ マップを参照してください。HOM Local は実証済みのメモリ アーキテクチャを継承しています。
Brain デーモンは以下から構成を読み取ります。
~/.hom/config.json — メイン構成
環境変数 — HOM_* プレフィックス
コマンドライン引数 — 「hom-brain --help」を参照
開発
# テストを実行する
貨物テスト --ワークスペース
# サンプルテストを実行する
貨物テスト --workspace --examples
# フォーマットコード
貨物輸送
# ドキュメントを作成する
貨物ドキュメント --open
ドキュメント
文書
説明
建築
クレート階層、ブレイン デーモン、ワーカー ディスパッチ、IPC プロトコル
APIリファレンス
Brain IPC メソッドと HTTP API ルート
構成
環境変数、設定ファイル、権限
セキュリティ
認証、セキュリティゲート、品質ゲート、監査証跡
メモリモデル
メモリ構造、タイプ、ソースの帰属、ベクトル埋め込み
リコールシステム
リコールモード、パイプライン、スコアリング、コンテキストパッキング
品質のゲート
4 つの壁による評価、品質スコアリング、ベンチマーク
オペレーターガイド
HOM Local を介して LLM/エージェントがどのように動作するか
テスト
テストの構造、テストの実行、カバレッジ
貢献する
開発セットアップ、貢献プロセス、ワークフロー
ローカル対オラクル
HOM Local と HOM Oracle の間のライセンス境界
よくある質問
一般、インストール、構成、メモリ、リコール、トラブルシューティング
ライセンス
HOM Local は、Apache License 2.0 に基づいてライセンスされています。
このライセンスは、このリポジトリ内のコードにのみ適用されます。
HOM Oracle は別個のプライベート システムであり、このリポジトリには含まれていません。
ライセンスの全文については、「LICENSE」を参照してください。
HOM Local — オープンソースのローカルファースト AI メモリ カーネル
Apache-2.0ライセンス
貢献する
Th

読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

HOM Local — open-source local-first AI memory kernel - wallidsaydi-creator/hom-local

GitHub - wallidsaydi-creator/hom-local: HOM Local — open-source local-first AI memory kernel · GitHub
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
wallidsaydi-creator
/
hom-local
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github crates crates docs docs scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml release-sanitize.sh release-sanitize.sh View all files Repository files navigation
Local-first memory server for AI agents and harnesses.
If you build AI systems that must remember work across sessions, model changes, and provider swaps, HOM Local gives you a persistent memory core with provenance, gates, and continuity.
If this saves your team time, please star the repo so more builders can find it.
Provider boundary and model choice
Post-compaction as a memory artifact
Most AI stacks still lose context every time a session ends, a model changes, or a handler restarts. That creates the re-briefing tax : re-stating decisions, constraints, and progress over and over.
HOM Local fixes that at the memory layer, not at the chat surface.
keep memory tied to provenance
keep continuity across provider and routing choices
HOM Local is for teams that care about memory as part of their backend, including:
AI copilots and harnesses that need context across long workflows.
Tool-rich agents where actions must remain auditable and recoverable.
Multi-model stacks that may switch from Claude/GPT/Gemini/etc. without rewriting memory.
Teams that want local-first deployment and explicit trust boundaries.
Structured memory with source attribution so agents can ask “what do we know?” and get a provenance-aware answer path.
Quality gates on memory write path: structure, relevance, substance, and factuality.
Mutation ledgering with tamper-evident local history for traceability and change visibility.
Session compaction is not just prompt trimming; it is durable continuity for later runs.
HOM Local is designed as dedicated memory service:
Three-crate workspace: hom-brain , hom-ingress , hom-shared .
Clear boundaries: ingress handles HTTP/auth; brain handles memory semantics; shared holds protocol + crypto contracts.
Deterministic IPC + HTTP edge: JSON-RPC over Unix domain sockets in the brain path with an HTTP boundary in ingress.
Concurrency design: weighted priority scheduling for save, recall, cognition, I/O, and main ops.
Provenance-first: append-only ledger plus source-linked memory artifacts.
Provider-neutral memory surface: model provider changes do not require memory rewrites.
Client/Agent → hom-ingress (auth/routing) → hom-brain (memory + gates + ledger) → SQLite WAL + vector index
What makes it different
HOM Local
Typical memory integrations
Local-first data control
✅
variable
Source-attributed recall
✅
often opaque
Provider-agnostic memory
✅
often coupled
Tamper-evident memory history
✅
often unavailable
Compaction-to-continuity pipeline
✅
often prompt-only
Quality-gated writes
✅
often simple or absent
Open architecture for app control
✅
often limited
If you need memory that outlives one chat UI, this is the core reason to test it.
HOM Local gives your app a durable memory API while keeping model providers at the app layer.
Build agent harnesses with persistent recall state.
Preserve context when you switch providers or run long sessions.
Route app outputs, evidence, and model responses into memory records.
Run continuity-aware workflows from compaction artifacts, not ad-hoc summaries.
✅ Durable local memory in SQLite with WAL mode
✅ Source-attributed recall with quality scoring
✅ Four-wall quality-gated memory saves (structure, relevance, substance, factuality)
✅ Tamper-evident append-only ledger with hash-chain verification
✅ Context packing with evidence cards and open handles
✅ Session compaction artifacts that remain durable and inspectable
✅ Product Quantization vector recall with exact rerank fallback
✅ Ed25519 envelope authentication
✅ JSON-RPC over Unix domain socket + HTTP ingress on 127.0.0.1:9101
✅ Tests, examples, and operator guidance
Temporal recall with bounded windows
Recall quality and coverage diagnostics
Anti-drift checks and enforcement
Quality gates + audit trail protections
a replacement for your product-specific agent harness
It is a backend memory brain your app can compose with.
Provider boundary and model choice
HOM Local does not own model providers.
Your app decides provider and runtime behavior.
HOM Local owns memory, recall, validation, compaction, and audit.
This is intentionally separated so memory remains stable while provider strategy evolves.
Post-compaction as a memory artifact
Most systems convert long context into a temporary summary and move on.
HOM Local persists post-compaction summaries as continuity artifacts with links back to source sessions and memories.
So continuity is recoverable, inspectable, and reusable instead of being discarded.
git clone https://github.com/wallidsaydi-creator/hom-local.git
cd hom-local
cargo build --release
Run each service in its own terminal:
cargo run --release --bin hom-brain
cargo run --release --bin hom-ingress
Your ingress should be available at http://127.0.0.1:9101 .
cargo install hom-brain
Architecture
hom-local/
├── crates/
│ ├── hom-brain/ # Core brain daemon — memory, ledger, recall, quality gates
│ ├── hom-shared/ # Shared types, crypto, envelope, RPC, paths
│ └── hom-ingress/ # HTTP ingress layer with auth
├── docs/ # Architecture and API documentation
└── examples/ # Usage examples
For the full service graph, see the Oracle Architecture Map — HOM Local inherits the proven memory architecture.
The brain daemon reads configuration from:
~/.hom/config.json — Main configuration
Environment variables — HOM_* prefix
Command line arguments — See hom-brain --help
Development
# Run tests
cargo test --workspace
# Run example tests
cargo test --workspace --examples
# Format code
cargo fmt
# Build docs
cargo doc --open
Documentation
Document
Description
Architecture
Crate hierarchy, brain daemon, worker dispatch, IPC protocol
API Reference
Brain IPC methods and HTTP API routes
Configuration
Environment variables, config file, permissions
Security
Authentication, security gates, quality gates, audit trail
Memory Model
Memory structure, types, source attribution, vector embeddings
Recall System
Recall modes, pipeline, scoring, context packing
Quality Gates
Four-wall assessment, quality scoring, benchmarks
Operator Guide
How an LLM/agent should operate through HOM Local
Testing
Test structure, running tests, coverage
Contributing
Development setup, contribution process, workflow
Local vs Oracle
Licensing boundary between HOM Local and HOM Oracle
FAQ
General, installation, configuration, memory, recall, troubleshooting
License
HOM Local is licensed under the Apache License, 2.0.
This license applies only to the code in this repository.
HOM Oracle is a separate private system and is not included in this repository.
See LICENSE for the full license text.
HOM Local — open-source local-first AI memory kernel
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

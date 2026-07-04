---
source: "https://github.com/nmxmxh/foundation"
hn_url: "https://news.ycombinator.com/item?id=48785026"
title: "Show HN: Foundation, a different approach to software and AI"
article_title: "GitHub - nmxmxh/foundation: A different approach to software and ai. An opinionated full stack framework delivering native performance and graphics to web experiences. · GitHub"
author: "MomohNobert"
captured_at: "2026-07-04T13:00:19Z"
capture_tool: "hn-digest"
hn_id: 48785026
score: 2
comments: 0
posted_at: "2026-07-04T12:46:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Foundation, a different approach to software and AI

- HN: [48785026](https://news.ycombinator.com/item?id=48785026)
- Source: [github.com](https://github.com/nmxmxh/foundation)
- Score: 2
- Comments: 0
- Posted: 2026-07-04T12:46:17Z

## Translation

タイトル: Show HN: Foundation、ソフトウェアと AI への異なるアプローチ
記事タイトル: GitHub - nmxmxh/foundation: ソフトウェアと AI への異なるアプローチ。 Web エクスペリエンスにネイティブのパフォーマンスとグラフィックスを提供する独自のフルスタック フレームワーク。 · GitHub
説明: ソフトウェアと AI に対する異なるアプローチ。 Web エクスペリエンスにネイティブのパフォーマンスとグラフィックスを提供する独自のフルスタック フレームワーク。 - nmxmxh/財団

記事本文:
GitHub - nmxmxh/foundation: ソフトウェアと AI に対する異なるアプローチ。 Web エクスペリエンスにネイティブのパフォーマンスとグラフィックスを提供する独自のフルスタック フレームワーク。 · GitHub
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
nmxmxh
/
基礎
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
55 コミット 55 コミット .agents/ workflows .agents/ workflows benchmark-results benchmark-results cmd/ ovasabi cmd/ ovasabi config-contracts config-contracts docs docs Frontend-kit Frontend-kit runtime-native runtime-native runtime-sdk runtime-sdk runtime-transport runtime-transport scriptsスクリプト サーバーキット サーバーキット テンプレート テンプレート テスト テスト ツール ツール ui-minimal ui-minimal .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md VERSION VERSION init.sh init.sh Rustfmt.toml Rustfmt.tomlすべてのファイルを表示 リポジトリ ファイルのナビゲーション
ovasabi Foundation (進行中の作業 — バージョン 0.0.1)
高性能のイベント駆動型システム用のフルスタック アプリケーション基板。
ovasabi Foundation は、コードを進化させたいチームのための統合ツールキットです。これは、実稼働システムをブートストラップおよび保守するためのプラットフォーム モジュール、足場、施行チェック、およびドキュメントを提供します。
テナント分離されたイベント駆動型アーキテクチャ - すべての操作には、誰が依頼したか、どの組織、相関 ID などのメタデータが含まれます。
パフォーマンス ラダー — ナノ秒の直接ディスパッチからマイクロ秒の JSON 互換性までの 7 つのプレーン
Hermes ホットプレーン — マイクロ秒未満の操作読み取りのための、制限されたノードローカル投影
ワーカー オーケストレーション — 再試行ポリシーと進行状況の追跡を備えた制限されたバックグラウンド処理
組み込みの可観測性 - 相関 ID によって自動的にリンクされるログ、メトリック、およびトレース
ノーコードプラットフォームではありません。ゼロDevOpsではありません。手を抜いて早く動きたいチームには向きません。 Foundation は、管理されたインフラストラクチャを採用し、パフォーマンスを理解し、

コードベースが進化すると予想されます。
コンポーネント
技術スタック
目的
サーバーキット
行く
バックエンド: イベント バス、ワーカー、Hermes、データベース、復元力、可観測性
ランタイムトランスポート
TypeScript
クライアント ワイヤ: コマンド バス、エンベロープ作成、メタデータ ストア、WebSocket/HTTP フォールバック
ランタイムSDK
錆び/WASM
高性能カーネル: 4KB 制御バッファ、ゼロコピー通信
ui-minimal
TypeScript/React
共有 UI プリミティブ、セマンティック テーマ トークン、モーション ヘルパー
フロントエンドキット
TypeScript
IndexedDB ストレージ、メタデータ ヘルパー、ランタイム アダプター、転送の進行状況
ランタイムネイティブ
タウリ/ラスト
ネイティブ シェル ブリッジ: 安全なストレージ、GPU ハンドル、デバイス アクセス
構成契約
Go/TypeScript
言語間の構成スキーマ
データ層 : PostgreSQL (永続的な真実)、Redis (調整)、プロトコル バッファー (コントラクト)、Cap'n Proto (ゼロコピー境界)
財団は 7 つのパフォーマンス プレーンを使用しています。各飛行機にはコストが測定され、適用されます。
1. 直接ディスパッチ 10 ～ 30 ns/op (同一プロセス、ゼロ割り当て)
2. バイナリ フレーム 20 ～ 80 ns/op (借用ビュー)
3. 生成された protobuf ~370 ns/op (型指定されたクロスプロセス)
4. gRPC 20 ～ 30 μs/op (ネットワーク機械)
5. JSON ~30 µs/op (互換性)
6. ネイティブ FFI/SHM (さまざま) (トラステッド コンピューティング)
7. ブラウザ + WASM + SAB (プラットフォーム) (サポートされている場合)
重要なルール: 最速のレーンは互換性レーンのコストを支払ってはなりません。これは自動的に測定されます。回帰は発生する前に発見されます。
詳細: docs/foundation_benchmarks.md
Foundation から生成されたすべてのプロジェクトは次のものを受け取ります。
マルチテナントの分離 — 組織の範囲は、クライアント データからではなく、認証されたコンテキストから派生します。
イベント駆動型神経システム - 相関メタデータを使用した正規の要求→成功/失敗のライフサイクル
Hermes Hotplane — データベースのミュートを投影する、ノードローカル、メモリ限定、インデックス付き読み取りモデル

リアルタイムのステーション
再開可能なファイル転送 - 再開可能な進行状況を伴うチャンクベースのアップロード/ダウンロード
境界付きワーカー処理 - 指数関数的バックオフ、再試行ポリシー、および境界付きキューを備えたバックグラウンド ジョブ
統合された可観測性 — OpenTelemetry トレース、構造化ログ、サーキット ブレーカー、エラー分類法
ここから開始 → docs/foundation_quick_start.md (15 分) → docs/foundation_tour.md (ウォークスルー) → docs/foundation_architecture_contract.md (プラットフォーム/プロジェクト分割)
ここから開始 → docs/PHILOSOPHY.md (Foundation が存在する理由) → docs/foundation_architecture_contract.md → docs/foundation_nervous_system.md → docs/practice_controls.md
AI エージェントとパートナー向け (エージェントネイティブ ワークフロー)
ここから開始 → AGENTS.md → docs/foundation_glossary.md → docs/agent_operating_contract.md → docs/ai_threat_model.md
パス
目的
サーバーキット/
Go バックエンド: レジストリ、メタデータ、イベント、ワーカー、復元力、可観測性、Hermes、イベントログ、Redis、データベース、転送、プロジェクション ゲートウェイ、オブジェクト ストレージ、一括操作、インテリジェンス シグナル
ランタイムトランスポート/
プロトコルコントラクト、コマンドバス、ルートレジストリ、バイナリコーデック、Hermesプロジェクションスキーマ
ランタイムSDK/
WASM/Rust/Go カーネル、4KB コントロールバッファ、共有アリーナ、ランタイムレーンヘルパー
ランタイムネイティブ/
Tauri シェル、安全なストレージ、ネイティブ フレーム、デバイス ディスパッチ
フロントエンドキット/
IndexedDB ストレージ、メタデータ、ランタイム アーティファクト、転送の進行状況
ui-minimal/
共有 UI プリミティブ、テーマ トークン、モーション ヘルパー
構成-契約/
生成された構成スキーマ
テンプレート/
生成されたプロジェクトにコピーされた足場テンプレート
ドキュメント/
アーキテクチャ、プラクティス、ガイド、セキュリティ、パフォーマンス、テスト
ツーリング/
強制スクリプト、マニフェスト、lint 構成
コアコマンド
makegenerate-contracts # コード生成 (Protos → Go/TS)
make lint # すべてのリンター
作る

est # すべてのテスト
make check-rust # Rust fmt、clippy、tests
make verify # 完全な CI スイート
make check-practice-controls # 練習マトリックス
make check-doc-references # リンクの検証
make docker-up # ローカルインフラを起動する
make test-service-backed # ライブ DB/Redis を使用したテスト
プロジェクトブートストラップ
Foundation の親ディレクトリから:
# 新しいプロジェクト
ノード Foundation/cmd/ovasabi/bin/ovasabi.js init --profile=パフォーマンス --name=my-app --foundation-dir Foundation --skip-license
# またはシェルスクリプト経由 (レガシー)
./foundation/scripts/init-project.sh 私のアプリが完全です
# 既存のプロジェクトを更新する
ノード Foundation/cmd/ovasabi/bin/ovasabi.js update --project-dir=/path/to/project --foundation-dir Foundation --skip-license
生成されたプロジェクトは、パッケージ境界を通じて Foundation を消費します。生の Foundation/*/ts/src または Foundation/*/go を直接インポートしないでください。
Foundation は、ソフトウェアの不足、つまりハードウェアのパフォーマンス (ナノ秒) と一般的なソフトウェア スタック (ミリ秒) の間のギャップを橋渡しします。以下の実証済みのパターンを提供します。
即座に応答 - ヘルメスを介したサブマイクロ秒の操作読み取り
安全なスケーリング — 制限された作業、テナントの分離、サーキット ブレーカー
明確に観察 - すべての層を流れる相関 ID
自信を持って進化する — 実践の強制、契約の検証、パフォーマンスの測定
ファンデーションは誰にでも使えるものではありません。それは要求が厳しいです。それには、パフォーマンスのトレードオフについて考えること、適切なテストを作成すること、障害モードを理解すること、ゲート判定を確認することなど、規律が必要です。これは、野心的なものを構築することを期待しているチーム向けです。
全文については docs/PHILOSOPHY.md をお読みください。
ここから始めてください: docs/README.md はドキュメント マップです。
docs/foundation_glossary.md — 概念の検索
docs/foundation_quick_start.md — 15 分のパス
docs/foundation_tour.md — 1 つのアクションのウォークスルー
ドキュメント/財団

n_architecture_contract.md — 所有権の分割
docs/foundation_nervous_system.md — ライフサイクル契約
AI ツールを使用している場合: AGENTS.md — エージェントのワークフローと証拠の要件
理由を理解するには: docs/PHILOSOPHY.md — 動機と設計原則
Foundation は積極的に進化しています。リポジトリ全体は、運用アプリケーションの進行中のベースラインを表します。期待できること:
契約と慣行の継続的な改良
新しいパフォーマンス プレーン (GPU コンピューティング、分散トレースの改良)
エージェントコーディングパターンはまだ実証中
使用パターンの出現に応じてドキュメントが拡大
研究者、エージェント、人間の査読者による貢献が財団の改善につながります。証拠と引き継ぎの期待については、docs/agent_operating_contract.md を参照してください。
ソフトウェアとaiに対する異なるアプローチ。 Web エクスペリエンスにネイティブのパフォーマンスとグラフィックスを提供する独自のフルスタック フレームワーク。
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

A different approach to software and ai. An opinionated full stack framework delivering native performance and graphics to web experiences. - nmxmxh/foundation

GitHub - nmxmxh/foundation: A different approach to software and ai. An opinionated full stack framework delivering native performance and graphics to web experiences. · GitHub
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
nmxmxh
/
foundation
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
55 Commits 55 Commits .agents/ workflows .agents/ workflows benchmark-results benchmark-results cmd/ ovasabi cmd/ ovasabi config-contracts config-contracts docs docs frontend-kit frontend-kit runtime-native runtime-native runtime-sdk runtime-sdk runtime-transport runtime-transport scripts scripts server-kit server-kit templates templates tests tests tooling tooling ui-minimal ui-minimal .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md VERSION VERSION init.sh init.sh rustfmt.toml rustfmt.toml View all files Repository files navigation
Ovasabi Foundation (Work in Progress — Version 0.0.1)
A full-stack application substrate for high-performance, event-driven systems.
Ovasabi Foundation is an integrated toolkit for teams that want to evolve code. It provides the platform modules, scaffolds, enforcement checks, and documentation to bootstrap and maintain production systems with:
Tenant-isolated, event-driven architecture — every operation carries metadata: who asked, what organization, correlation ID
Performance ladder — seven planes from nanosecond direct dispatch to microsecond JSON compatibility
Hermes hotplane — bounded, node-local projections for sub-microsecond operational reads
Worker orchestration — bounded background processing with retry policies and progress tracking
Built-in observability — logs, metrics, and traces automatically linked by correlation ID
Not a no-code platform. Not zero-DevOps. Not for teams that want to move fast by cutting corners. Foundation is for teams that embrace managed infrastructure, understand performance, and expect their codebase to evolve.
Component
Tech Stack
Purpose
server-kit
Go
Backend: event bus, workers, Hermes, database, resilience, observability
runtime-transport
TypeScript
Client wire: command bus, envelope creation, metadata stores, WebSocket/HTTP fallback
runtime-sdk
Rust/WASM
High-performance kernel: 4KB control buffer, zero-copy communication
ui-minimal
TypeScript/React
Shared UI primitives, semantic theme tokens, motion helpers
frontend-kit
TypeScript
IndexedDB storage, metadata helpers, runtime adapters, transfer progress
runtime-native
Tauri/Rust
Native shell bridge: secure storage, GPU handles, device access
config-contracts
Go/TypeScript
Cross-language configuration schemas
Data Layer : PostgreSQL (durable truth), Redis (coordination), Protocol Buffers (contracts), Cap'n Proto (zero-copy boundaries)
Foundation uses seven performance planes. Each plane has its cost measured and enforced:
1. Direct dispatch 10–30 ns/op (same-process, zero-alloc)
2. Binary frames 20–80 ns/op (borrowed views)
3. Generated protobuf ~370 ns/op (typed cross-process)
4. gRPC 20–30 µs/op (network machinery)
5. JSON ~30 µs/op (compatibility)
6. Native FFI/SHM (varies) (trusted compute)
7. Browser + WASM + SAB (platform) (where supported)
Key rule : The fastest lane must not pay the cost of the compatibility lane. This is measured automatically; regressions are caught before they land.
Read more: docs/foundation_benchmarks.md
Every project generated from Foundation receives:
Multi-Tenant Isolation — organization scope derived from authenticated context, never from client data
Event-Driven Nervous System — canonical requested → success / failed lifecycle with correlation metadata
Hermes Hotplane — node-local, memory-bounded, indexed read models that project database mutations in real-time
Resumable File Transfers — progress-bearing, chunk-based upload/download with resumability
Bounded Worker Processing — background jobs with exponential backoff, retry policies, and bounded queues
Unified Observability — OpenTelemetry tracing, structured logs, circuit breakers, error taxonomy
Start here → docs/foundation_quick_start.md (15 min) → docs/foundation_tour.md (walk-through) → docs/foundation_architecture_contract.md (platform/project split)
Start here → docs/PHILOSOPHY.md (why Foundation exists) → docs/foundation_architecture_contract.md → docs/foundation_nervous_system.md → docs/practice_controls.md
For AI Agents & Partners (Agent-Native Workflow)
Start here → AGENTS.md → docs/foundation_glossary.md → docs/agent_operating_contract.md → docs/ai_threat_model.md
Path
Purpose
server-kit/
Go backend: registry, metadata, events, workers, resilience, observability, Hermes, eventlog, Redis, database, transfer, projection gateway, object storage, bulk operations, intelligence signals
runtime-transport/
Protocol contracts, command bus, route registry, binary codecs, Hermes projection schemas
runtime-sdk/
WASM/Rust/Go kernel, 4KB control-buffer, shared arena, runtime lane helpers
runtime-native/
Tauri shell, secure storage, native frames, device dispatch
frontend-kit/
IndexedDB storage, metadata, runtime artifacts, transfer progress
ui-minimal/
Shared UI primitives, theme tokens, motion helpers
config-contracts/
Generated configuration schemas
templates/
Scaffold templates copied into generated projects
docs/
Architecture, practices, guides, security, performance, testing
tooling/
Enforcement scripts, manifests, lint configs
Core Commands
make generate-contracts # Code gen (Protos → Go/TS)
make lint # All linters
make test # All tests
make check-rust # Rust fmt, clippy, tests
make verify # Full CI suite
make check-practice-controls # Practice matrix
make check-doc-references # Link validation
make docker-up # Start local infra
make test-service-backed # Tests with live DB/Redis
Project Bootstrap
From the parent directory of foundation :
# New project
node foundation/cmd/ovasabi/bin/ovasabi.js init --profile=performance --name=my-app --foundation-dir foundation --skip-license
# Or via shell script (legacy)
./foundation/scripts/init-project.sh my-app full
# Update existing project
node foundation/cmd/ovasabi/bin/ovasabi.js update --project-dir=/path/to/project --foundation-dir foundation --skip-license
Generated projects consume Foundation through package boundaries. Do not import raw foundation/*/ts/src or foundation/*/go directly.
Foundation bridges the software deficit : the gap between hardware performance (nanoseconds) and typical software stacks (milliseconds). It provides proven patterns for:
Responding instantly — sub-microsecond operational reads via Hermes
Scaling safely — bounded work, tenant isolation, circuit breakers
Observing clearly — correlation IDs flowing through all layers
Evolving confidently — enforced practices, contract verification, performance measurement
Foundation is not for everyone. It's demanding. It requires discipline: thinking about performance trade-offs, writing adequate tests, understanding failure modes, reviewing gate verdicts. It's for teams that expect to build something ambitious.
Read docs/PHILOSOPHY.md for the full story.
Start here : docs/README.md is the documentation map.
docs/foundation_glossary.md — concept lookup
docs/foundation_quick_start.md — 15-minute path
docs/foundation_tour.md — walk-through one action
docs/foundation_architecture_contract.md — ownership split
docs/foundation_nervous_system.md — lifecycle contract
If you're using AI tools : AGENTS.md — agent workflows and evidence requirements
To understand why : docs/PHILOSOPHY.md — the motivation and design principles
Foundation is actively evolving. The entire repository represents a work-in-progress baseline for production applications. Expect:
Continued refinement of contracts and practices
New performance planes (GPU compute, distributed tracing refinement)
Agentic coding patterns still being proven
Documentation expanding as usage patterns emerge
Contributions via research, agents, and human reviewers are how Foundation improves. Read docs/agent_operating_contract.md for evidence and handoff expectations.
A different approach to software and ai. An opinionated full stack framework delivering native performance and graphics to web experiences.
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

---
source: "https://github.com/democr-ai/democrai"
hn_url: "https://news.ycombinator.com/item?id=48935704"
title: "Show HN: Democr.ai – self-hosted agentic AI runtime with audit and RBAC"
article_title: "GitHub - democr-ai/democrai · GitHub"
author: "fabio2"
captured_at: "2026-07-16T15:26:51Z"
capture_tool: "hn-digest"
hn_id: 48935704
score: 1
comments: 0
posted_at: "2026-07-16T15:13:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Democr.ai – self-hosted agentic AI runtime with audit and RBAC

- HN: [48935704](https://news.ycombinator.com/item?id=48935704)
- Source: [github.com](https://github.com/democr-ai/democrai)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T15:13:52Z

## Translation

タイトル: Show HN: Democr.ai – 監査と RBAC を備えた自己ホスト型エージェント AI ランタイム
記事タイトル: GitHub - デモクラ-ai/democrai · GitHub
説明: GitHub でアカウントを作成して、democr-ai/democrai の開発に貢献します。

記事本文:
GitHub - デモクラ-ai/デモクライ · GitHub
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
デモクラアイ
/
民主主義
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル

レ
65 コミット 65 コミット クライアント クライアント デモクライ デモクレイ docker docker docs_site docs_site エンジン エンジン サンプル サンプル エクストラクター エクストラクター モジュール モジュール スクリプト スクリプト テスト テスト .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md __init__.py __init__.py constraints.txtconstraints.txtmain.pymain.pypyproject.tomlpyproject.tomlrequirements.txtrequirements.txtrun.batrun.batrun.pyrun.pyrun.shrun.shsetup_venv.batsetup_venv.batsetup_venv.pysetup_venv.pysetup_venv.shsetup_venv.shすべてのファイルを表示 リポジトリファイルのナビゲーション
サーバー駆動の UI、ネイティブの可観測性、OS レベルのサンドボックス、プラグイン可能なモデル オーケストレーション、厳密な拡張境界を備えたエージェント AI アプリケーションを構築するための Python フレームワーク。再現性、監査可能性、運用制御が重要な環境向けに設計されています。
ステータス: ベータ 0.0.1b4 — パブリック プレビュー。コア アーキテクチャは実装され、積極的にテストされていますが、API と動作動作は最初の安定版リリースまでに変更される可能性があります。ソフトウェアは保証なしで出荷されます。
ドキュメント：democr.ai/docs/
Democr.ai は、AI アプリケーション用の完全なランタイム フレームワークです。それは以下を提供します:
エージェントコンテキスト用のサーバー駆動型 UI、Google の A2UI プロトコルの実装および拡張
Web クライアントとデスクトップ クライアント間で同じ UI 定義のマルチクライアント レンダリング
書き込み時にネイティブのマルチテナンシーが適用される
エンジンおよび拡張プロセスのためのクロスプラットフォーム OS レベルのサンドボックス
SQLAlchemy フックによる 3 層監査、機密フィールドの編集
宣言型モジュールマニフェストを使用した RBAC
ローカル、リモート、分散ランタイムにわたるプラグイン可能な AI エンジン オーケストレーション
エンジンおよびサブジェクトのスコープごとのリクエストおよびトークン ガバナンスのエンジン クォータ
ベクトルとグラフを備えた知識サブシステム

バックエンド
認証を含むすべてがパブリック SDK に対して構築されたモジュールである厳密なモジュラー アーキテクチャ
このフレームワークは、コア ランタイムと拡張機能の間の明確な境界を中心に設計されています。 SDK は、その上に構築するためにサポートされている唯一の方法です。
ほとんどの AI アプリケーション スタックは、チャット UI、LLM クライアント、ワークフロー ライブラリ、監査、権限、サンドボックス、可観測性のための個別のインフラストラクチャを組み合わせています。 Democr.ai は、これらの懸念を 1 つの実行時コントラクトにまとめます。
UI はサーバー側で宣言され、独立したクライアントによってレンダリングされます。
モジュール、エンジン、エクストラクターは、パブリック SDK 境界を持つインストール可能な拡張機能です。
モデル呼び出しは、ローカル プロバイダー、リモート プロバイダー、プロセス分離、クォータ、およびマルチノード実行をサポートするオーケストレーターを介して流れます。
セキュリティと監査は実行時の基本的なものであり、アプリケーションの規約ではありません。
ナレッジの取り込み、取得、メディア処理、ツール、MCP、およびエージェントのスキルは、同じリクエストと可観測性コンテキストを共有します。
その組み合わせは意図的に幅広くなっています。このプロジェクトはまだベータ版ですが、アーキテクチャはすでにデモや単一チャットボットのプロトタイプではなく、運用グレードの制約を対象としています。
エージェントは、スクリーンショットやマークアップ フラグメントとしてではなく、ネイティブ アプリケーション サーフェスとしてレンダリングする UI コンポーネントを構築できます。 Democr.ai モジュールは、サーバー上で UI を宣言的に記述し、エージェントの動作に応じて構造化された更新をストリーミングします。
短いデモでは重要な部分を示しています。つまり、バックエンドが動作、アクション、実行時効果の所有権を保持しながら、バックエンドで宣言された 1 つの YAML サーフェスがクライアント レンダラーを通じて反映されます。
Democr.ai は、次のようなアプリケーションをターゲットとしています。
導入環境にはコンプライアンス、規制、または運用上の制約があります
観察性と再現性は初日から必要であり、後から追加されるものではありません
ベンダーロックインの回避

これは厳しい制約です。フレームワークは完全に自己ホスト可能で、モデルに依存せず、独自のクラウド サービスは提供されません。
ユーザーのアクション、データの変更、AI 呼び出し全体にわたる監査可能性が展開レベルで必要です
チームは、個別の部分 (LLM ライブラリ + UI フレームワーク + オブザーバビリティ スタック + 監査 + RBAC + ワークフロー エンジン) を組み立てるよりも、統合されたランタイムを好みます。
これは、シン ラッパー、モデル ルーター、またはチャット UI ライブラリではありません。これは、モジュール、エンジン、エクストラクター、クライアントが単一の一貫したシステムを構成するランタイムです。
特権のあるコンポーネントはありません。このリポジトリに同梱されている認証、システム管理、およびデモ アプリケーションはすべて、サードパーティ開発者が使用するものと同じ SDK を使用してモジュールとして構築されています。予約された API、隠しフック、内部専用サービスはありません。
このリポジトリに含まれるモジュール ( auth 、 system 、Components ) は主に、フレームワークがどのように機能するか、およびアプリケーションがフレームワークとどのように統合されるかを示す例です。これらは実用的なベースラインを提供しますが、デプロイメントではそれらをそのまま使用する必要はありません。チームはそれらを置き換えたり、拡張したり、独自のドメイン固有の実装を構築したりできます。
ランタイムは完全に自己ホスト可能です。
AI エンジンはプラグ可能です: ローカル推論 (vLLM、llama.cpp) およびクラウド プロバイダー (OpenAI、Anthropic、Google など)、またはカスタム統合。
ストレージ バックエンドは、ベクター (sqlite-vec、pgvector など)、ナレッジ グラフ、コア永続性 (SQLite、PostgreSQL) にプラグ可能です。
独自のシリアル化形式はありません。状態はデプロイメント間で移植可能です。
開発ガイドとしての可観測性と再現性
可観測性は、後のインストルメンテーションとして追加されるものではありません。これは、フレームワークが開発者に構築を期待する方法の一部です。すべてのアクション、AI 呼び出し、モジュール操作は相関チェーンを介して流れます。

これにより、あらゆるリクエストを段階的に再構築できるようになります。パイプラインは観察可能です。監査サブシステムは、機密フィールドの編集を使用して、3 つのレイヤー (リクエスト、ORM、AI 呼び出し) で突然変異を自動的にキャプチャします。
再現性は、確定的な構成、安定した ID を持つ構造化されたイベント ログ、および不変の監査証跡によって実現されます。
機能ではなく基本的なセキュリティ
このフレームワークは、オプトイン機能としてではなく、基本的なプリミティブとしてセキュリティを提供します。フレームワークの保証と導入者の責任の間の正確な境界については、「脅威モデル」を参照してください。
Google の A2UI プロトコルの拡張
UI レイヤーは、エージェント コンテキストでのサーバー駆動型 UI 用の Google の A2UI プロトコルの拡張実装です。モジュール作成者がクライアント側のフレームワーク コードを必要とせずに、同じ UI 定義が複数のクライアント間で一貫してレンダリングされます。既存のモジュールを変更せずに、新しいクライアントを追加できます。
Democr.ai は、次のような異なるサブシステムで構成されています。
コア ランタイム — ブートストラップ、ライフサイクル、リクエスト処理、依存関係グラフ
ネットワーク — 統一プロトコル層下の WebSocket、IPC、HTTP
AI エンジン オーケストレーター — ジョブベースのスケジューリング、バッチ処理、gRPC プロセス分離、HITL サポート
知識 — インジェストキュー、プロジェクションワーカー、ベクトル + グラフ検索
サンドボックス — OS レベルの分離 (Landlock + seccomp + iptables + ヘルパー プロセス | シートベルト | Windows Low Integrity + WFP)
マルチテナンシー - マテリアライズされたスコープフィルターによる書き込み時の強制
RBAC — モジュールマニフェスト内の宣言型アクセスポリシー
モジュール — フレームワーク拡張サーフェス、SDK 経由
エンジン — プラグイン可能な AI 推論プロバイダー
エクストラクター — プラグイン可能な知識抽出プロバイダー
クライアント — Web、デスクトップ、およびモバイル指向のクライアント用の独立したレンダリング サーフェス
プラットフォーム — エージェント、スキル、ツール、MCP、

UIビルダー
サブシステムの詳細なドキュメントは、democr.ai/docs/ で入手できます。
能力
ステータス
注意事項
サーバー駆動型 UI (拡張 A2UI)
実装済み
WebSocket と IPC トランスポート
マルチクライアントレンダリング
実装済み
Web、Qt デスクトップ、Tauri、React バリアント
マルチテナンシー
実装済み
書き込み時の強制
RBAC
実装済み
宣言的マニフェスト
AI エンジンのオーケストレーション
実装済み
ジョブベース、分離された gRPC
知識サブシステム (ベクトル + KG)
実装済み
ハイブリッド検索
監査（三重層）
実装済み
機密フィールドの編集
OS サンドボックス (Landlock + seccomp + iptables)
実装済み
macOS 上の OS サンドボックス
実装済み
シートベルト
Windows 上の OS サンドボックス
実装済み
低整合性 + WFP; Egress の強制は昇格されたヘルパーを使用します
ローカル AI 推論
実装済み
vLLM、llama.cpp、ONNX、Ollama、その他
クラウドAIプロバイダー
実装済み
OpenAI、Anthropic、Google など
マルチノードエンジンオーケストレーション
実装済み
共有キュー、ノードレジストリ、Redis 応答ストリーム
内部 gRPC サービス用の mTLS
ロードマップ
サーバーサイド TLS + サービス JWT が最初
コンテキストバジェットガード
実装済み
モデルのコンテキスト設定によってスケールされる、サイズが大きすぎるプロンプト メッセージとツール出力に対するヒューリスティック ガード
マルチノードのチューニング
ロードマップ
分散型微調整ワークフロー
モバイルクライアント
実験的
React Native クライアントは存在しますが、認定されていません
エンジン割り当て
実装済み
エンジンおよびサブジェクトのスコープごとのリクエストとトークンの制限
リクエストの可観測性
実装済み
レイヤー間の相関チェーン
ストレスベンチマークのスナップショット
次の数値は、認証されたユーザーのローカル ストレス テストのスナップショットです。
リクエスト パス: ログイン、セッション処理、およびサーバー側のレンダリング
サーバー駆動型 UI ルート /components/_Effects/yaml 。彼らは観察したことを記録します
1 台のマシンでの同時負荷時の動作。それらは普遍的ではありません
パフォーマンスに関する主張または他のものとの比較

r フレームワーク。
GPU: NVIDIA RTX 3090 Suprim、このベンチマーク パスでは使用されません
パイソンパス=。 .venv/bin/python スクリプト/multi_core_benchmark.py \
--インスタンス 4 \
--労働者 4 \
--クライアント 4 \
--duration-秒 300 \
--プロファイルカスタム\
--profile-log-summary
4 つのインスタンスの実行では、合計リソース使用量は次のようになりました。
共有ページがより多くカウントされるため、RSS はスタンドアロンのメモリ使用量を誇張する可能性があります
一度よりも。
このフレームワークは、OS レベルのサンドボックス、不変監査、RBAC、範囲指定されたデータ アクセス、データベースに保存されているシークレットの保存時の暗号化など、基本的なプリミティブとしてセキュリティを提供します。モジュールは、宣言されたポリシー (サンドボックス許可リスト、RBAC 権限、アクセス スコープ) に従ってこれらのプリミティブ内で動作します。
このフレームワークが対象とする脅威モデルは、サプライチェーンの偶発的な侵害です。バグまたは依存関係の侵害があるサードパーティ モジュールは、宣言された範囲を超えてシステムに害を及ぼすことができてはなりません。
カバーされていない脅威モデルは、展開者によって意図的にインストールされた意図的に悪意のあるモジュールです。インストールするモジュールの選択とセキュリティ ポリシーの構成は、導入者の責任です。
フレームワークの config.yaml はプレーンテキストであり、セットアップ ウィザードによってローカルに生成されます。導入するように設計されていません。秘密

[切り捨てられた]

## Original Extract

Contribute to democr-ai/democrai development by creating an account on GitHub.

GitHub - democr-ai/democrai · GitHub
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
democr-ai
/
democrai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
65 Commits 65 Commits clients clients democrai democrai docker docker docs_site docs_site engines engines examples examples extractors extractors modules modules scripts scripts tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md __init__.py __init__.py constraints.txt constraints.txt main.py main.py pyproject.toml pyproject.toml requirements.txt requirements.txt run.bat run.bat run.py run.py run.sh run.sh setup_venv.bat setup_venv.bat setup_venv.py setup_venv.py setup_venv.sh setup_venv.sh View all files Repository files navigation
A Python framework for building agentic AI applications with server-driven UI, native observability, OS-level sandboxing, pluggable model orchestration, and a strict extension boundary — designed for environments where reproducibility, auditability, and operational control matter.
Status: Beta 0.0.1b4 — public preview. The core architecture is implemented and actively tested, but APIs and operational behavior may still change before the first stable release. The software ships without warranty.
Documentation: democr.ai/docs/
Democr.ai is a complete runtime framework for AI applications. It provides:
Server-driven UI for agentic contexts, implementing and extending Google's A2UI protocol
Multi-client rendering of the same UI definition across web and desktop clients
Native multi-tenancy enforced at write time
Cross-platform OS-level sandboxing for engine and extension processes
Triple-layer audit through SQLAlchemy hooks, with sensitive field redaction
RBAC with declarative module manifests
Pluggable AI engine orchestration across local, remote, and distributed runtimes
Engine quotas for request and token governance by engine and subject scope
Knowledge subsystem with vector and graph backends
A strict modular architecture where everything — including authentication — is a module built against the public SDK
The framework is designed around a clear boundary between core runtime and extensions. The SDK is the only supported way to build on top.
Most AI application stacks combine a chat UI, an LLM client, a workflow library, and separate infrastructure for audit, permissions, sandboxing, and observability. Democr.ai puts those concerns in one runtime contract:
UI is declared server-side and rendered by independent clients.
Modules, engines, and extractors are installable extensions with a public SDK boundary.
Model calls flow through an orchestrator that supports local providers, remote providers, process isolation, quotas, and multi-node execution.
Security and audit are runtime primitives, not application conventions.
Knowledge ingestion, retrieval, media handling, tools, MCP, and agent skills share the same request and observability context.
That combination is intentionally broad. The project is still beta, but the architecture is already aimed at production-grade constraints rather than demos or single-chatbot prototypes.
Let agents build UI components that render as native application surfaces, not as screenshots or markup fragments. Democr.ai modules describe UI declaratively on the server, then stream structured updates as the agent works.
The short demo shows the important part: one backend-declared YAML surface reflected through client renderers, while the backend keeps ownership of behavior, actions, and runtime effects.
Democr.ai targets applications where:
The deployment environment has compliance, regulatory, or operational constraints
Observability and reproducibility are required from day one, not added later
Avoiding vendor lock-in is a hard constraint: the framework is fully self-hostable, model-agnostic, and ships no proprietary cloud services
Auditability across user actions, data mutations, and AI calls is required at the deployment level
The team prefers an integrated runtime over assembling separate pieces (LLM library + UI framework + observability stack + audit + RBAC + workflow engine)
It is not a thin wrapper, a model router, or a chat UI library. It is a runtime where modules, engines, extractors, and clients compose into a single coherent system.
There are no privileged components. Authentication, system administration, and the demo applications shipped in this repository are all built as modules, using the same SDK that third-party developers use. There are no reserved APIs, no hidden hooks, and no internal-only services.
The modules included in this repository ( auth , system , components ) are primarily examples of how the framework works and how applications integrate with it. They provide a working baseline, but deployments are not required to use them as-is: teams can replace them, extend them, or build their own domain-specific implementations.
The runtime is fully self-hostable.
AI engines are pluggable: local inference (vLLM, llama.cpp) and cloud providers (OpenAI, Anthropic, Google, and others), or custom integrations.
Storage backends are pluggable for vector (sqlite-vec, pgvector, and others), knowledge graph, and core persistence (SQLite, PostgreSQL).
No proprietary serialization formats. State is portable across deployments.
Observability and reproducibility as development guides
Observability is not added as later instrumentation — it is part of how the framework expects developers to build. Every action, AI call, and module operation flows through a correlation chain that allows step-by-step reconstruction of any request. Pipelines are observable. The audit subsystem captures mutations at three layers (request, ORM, AI invocation) automatically, with sensitive field redaction.
Reproducibility is achieved through deterministic configuration, structured event logs with stable IDs, and an immutable audit trail.
Security as a primitive, not a feature
The framework provides security as foundational primitives, not as opt-in features. See Threat model for the precise boundary between framework guarantees and deployer responsibilities.
Extending Google's A2UI protocol
The UI layer is an extended implementation of Google's A2UI protocol for server-driven UI in agentic contexts. The same UI definition renders consistently across multiple clients with no client-side framework code required from module authors. New clients can be added without changes to existing modules.
Democr.ai is composed of distinct subsystems:
Core runtime — bootstrap, lifecycle, request handling, dependency graph
Network — WebSocket, IPC, HTTP under a unified protocol layer
AI engine orchestrator — job-based scheduling, batching, gRPC process isolation, HITL support
Knowledge — ingestion queue, projection workers, vector + graph retrieval
Sandbox — OS-level isolation (Landlock + seccomp + iptables + helper process | seatbelt | Windows Low Integrity + WFP)
Multi-tenancy — write-time enforcement with materialized scope filters
RBAC — declarative access policies in module manifests
Modules — framework extension surface, via SDK
Engines — pluggable AI inference providers
Extractors — pluggable knowledge extraction providers
Clients — independent rendering surfaces for web, desktop, and mobile-oriented clients
Platform — agents, skills, tools, MCP, UI builder
Detailed subsystem documentation is available at democr.ai/docs/ .
Capability
Status
Notes
Server-driven UI (extended A2UI)
Implemented
websocket and ipc transport
Multi-client rendering
Implemented
Web, Qt desktop, Tauri, React variants
Multi-tenancy
Implemented
Write-time enforcement
RBAC
Implemented
Declarative manifests
AI engine orchestration
Implemented
Job-based, gRPC isolated
Knowledge subsystem (vector + KG)
Implemented
Hybrid retrieval
Audit (triple layer)
Implemented
Sensitive field redaction
OS sandbox (Landlock + seccomp + iptables)
Implemented
OS sandbox on macOS
Implemented
Seatbelt
OS sandbox on Windows
Implemented
Low Integrity + WFP; egress enforcement uses an elevated helper
Local AI inference
Implemented
vLLM, llama.cpp, ONNX, Ollama, and others
Cloud AI providers
Implemented
OpenAI, Anthropic, Google, and others
Multi-node engine orchestration
Implemented
Shared queue, node registry, Redis response streams
mTLS for internal gRPC services
Roadmap
Server-side TLS + service JWT first
Context budget guard
Implemented
Heuristic guard for oversized prompt messages and tool outputs, scaled by model context settings
Multi-node tuning
Roadmap
Distributed fine-tuning workflows
Mobile clients
Experimental
React Native client exists but is not certified
Engine quotas
Implemented
Request and token limits by engine and subject scope
Request observability
Implemented
Correlation chain across layers
Stress benchmark snapshot
The following numbers are local stress-test snapshots for the authenticated
request path: login, session handling, and server-side rendering of the
server-driven UI route /components/_effects/yaml . They document observed
behavior under concurrent load on one machine; they are not universal
performance claims or comparisons with other frameworks.
GPU: NVIDIA RTX 3090 Suprim, not used by this benchmark path
PYTHONPATH=. .venv/bin/python scripts/multi_core_benchmark.py \
--instances 4 \
--workers 4 \
--clients 4 \
--duration-seconds 300 \
--profile custom \
--profile-log-summary
For the 4-instance run, aggregate resource usage was:
RSS can overstate standalone memory use because shared pages are counted more
than once.
The framework offers security as foundational primitives: OS-level sandboxing, immutable audit, RBAC, scoped data access, and at-rest encryption of secrets stored in the database. Modules operate within these primitives according to the policies they declare (sandbox allow-lists, RBAC permissions, access scope).
The threat model covered by the framework is accidental supply-chain compromise — a third-party module with a bug or compromised dependency must not be able to harm the system beyond its declared scope.
The threat model not covered is a deliberately malicious module installed consciously by the deployer. The selection of modules to install and the configuration of security policies are deployer responsibilities.
The framework's config.yaml is plaintext and is generated locally by the setup wizard. It is not designed to be deployed. Secre

[truncated]

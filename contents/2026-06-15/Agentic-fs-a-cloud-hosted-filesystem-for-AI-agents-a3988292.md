---
source: "https://github.com/vivekkhimani/agentic-fs"
hn_url: "https://news.ycombinator.com/item?id=48538435"
title: "Agentic-fs, a cloud-hosted filesystem for AI agents"
article_title: "GitHub - vivekkhimani/agentic-fs: Filesystem-style access to your documents, for AI agents · GitHub"
author: "vivekkhimani"
captured_at: "2026-06-15T11:11:58Z"
capture_tool: "hn-digest"
hn_id: 48538435
score: 1
comments: 0
posted_at: "2026-06-15T08:49:55Z"
tags:
  - hacker-news
  - translated
---

# Agentic-fs, a cloud-hosted filesystem for AI agents

- HN: [48538435](https://news.ycombinator.com/item?id=48538435)
- Source: [github.com](https://github.com/vivekkhimani/agentic-fs)
- Score: 1
- Comments: 0
- Posted: 2026-06-15T08:49:55Z

## Translation

タイトル: Agentic-fs、AI エージェント用のクラウドホスト型ファイルシステム
記事のタイトル: GitHub - vivekkhimani/agentic-fs: AI エージェント向けのドキュメントへのファイルシステム スタイルのアクセス · GitHub
説明: AI エージェント向けのファイルシステム スタイルのドキュメントへのアクセス - vivekkhimani/agentic-fs

記事本文:
GitHub - vivekkhimani/agentic-fs: AI エージェント向けのドキュメントへのファイルシステム スタイルのアクセス · GitHub
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
ヴィヴェッキマニ
/
エージェント-FS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスト

er ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
142 コミット 142 コミット .agents/ スキル .agents/ スキル .claude/ スキル .claude/ スキル .github .github docs docs パッケージ パッケージ terraform terraform Web Web .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.worker Dockerfile.worker ライセンス ライセンス Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml skill-lock.json skill-lock.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが独自の AWS でドキュメントにファイルシステム スタイルでアクセス
アカウント。 list/glob/grep/tree/find/ranged 読み取り
S3 内のドキュメント。MCP (および REST) を通じて公開されます。マルチテナントなので、
1 つの Terraform apply でデプロイし、アイドル状態で月額 $2 程度のコストがかかり、
ステートフル レイヤーは交換可能です。
ステータス: 初期段階で、活発に開発中です。 v1.0.0 は PyPI に公開されており、
リポジトリは公開されています。貢献は大歓迎です (
未解決の問題と
docs/build-progress.md )。完全なループが実行されます
AWS : スケジュールされた、取り込み、抽出、カタログ、MCP/REST 読み取りサーフェス
S3 からの修復および高信号アラーム。ライセンス: Apache-2.0。背景と
理論的根拠は docs/agentic-fs-oss-plan.md に記載されています。
境界があり、スコープが設定された MCP ツール サーフェス。エージェントは文書コーパスを探索します。
コーディング エージェントは、抽出されたドキュメント テキストを除き、インデックス付けされたリポジトリを探索します。
スケール、マルチテナント、およびリモート:
ナビゲートします。 fs_list 、 fs_tree 、 fs_glob 、 fs_find (タイプ、サイズ、mtime、ステータス別)。
検索。 fs_grep : 2 段階、制限付き、ripgrep スタイルのフィルター。
読む。 fs_read (範囲指定、またはセクションごと)、f

s_outline (ドキュメントの見出しマップ)、 fs_tables 、 fs_diff 。
仕事。 Scratch_* (プリンシパルごとのワークスペース)、whoami 。
すべてのツールは、クレーム フィルターによる可視性を適用する 1 つのミドルウェアを通じて実行されます。
スコープの適用、呼び出しごとの出力バジェット、および監査ログ
( ADR 0012 )。ツールを追加するということは、
フォークではなく、レジストリ エントリです。セマンティック fs_search は、オプションのアクセラレータです。
ロードマップ。 grep はフロアです。
要件: Docker 、
uv 、そして作る
(macOS: xcode-select --install )。
git clone https://github.com/vivekkhimani/agentic-fs && cd Agentic-fs
make dev # イメージを構築し、MinIO + DynamoDB Local + API を開始し、バケット/テーブルをシードします
curl localhost:8080/v1/healthz # {"ステータス":"ok","バージョン":"..."}
curl localhost:8080/v1/me # ローカルの開発プリンシパル
# ドキュメントのフォルダーを取り込み、それを読み返します。
uv run fs-crawler --connector local --source ./docs --api-url http://localhost:8080 --namespace ハンドブック
curl " localhost:8080/v1/fs/handbook/entries " # 表示されたカタログ行
MCP サーフェスは localhost:8080/mcp にマウントされているため、任意の MCP をポイントできます
それのクライアント。 make down はスタックを停止し、make clean はボリュームをワイプします。
API は、AWS Lambda および Fargate で実行されるのと同じコンテナ イメージです。
( ADR 0003 )。
ローカル dev は静的 dev プリンシパル ( AFS_AUTH_MODE=dev ) を使用します。決して実行しないでください
生産中です。運用環境では、agentic-fs は OAuth 2.1 リソース サーバーです。
独自の IdP (WorkOS、Cognito、Auth0、Okta、Keycloak) を使用すると、afs auth Doctor がトークンがプリンシパルにどのようにマッピングされるかを正確に示します。
(認証スワップガイド、ADR 0013)。
uv sync # Python ワークスペースをセットアップします (1 回)
make test # テストスイートを実行する
make lint # ruff lint + フォーマットチェック
make fmt # 自動フォーマット + 自動修正
make help # すべてのターゲットをリストする
すべての PR は CI: Packages/** の Python (ruff + pytest) によってゲートされます。
Terraform (fmt、val

idate、tflint、trivy) terraform/** の場合。
パッケージ/
afs-core/ コントラクト (プロトコル)、DTO、キースキーム、適合キット (pydantic のみ)
afs-server/ ストア、サービス、抽出、FastAPI アプリ + MCP マウント (afs-core を実装)
afs-connector-sdk/ fs-crawler CLI + 同期エンジン + ローカル FS / S3 / ドライブ / LlamaHub コネクタ
terraform/モジュラー IaC: グローバル状態/CI ロール、レイヤーごとのモジュール、例
ドキュメント/計画、ビルドの進捗状況、スワップ ガイド、意思決定記録 (ADR)
Dockerfile 1 つのイメージ: Lambda + Fargate + ローカル
任意のレイヤーを交換 (プラグアンドプレイ)
各レイヤーは、適合キットと 1 ページの小さな契約の背後にあります。
ガイドを参照して、既存のインフラストラクチャ上で実行できるようにします。
設定のバックエンド名とエントリポイント検出によって機能します
( ADR 0002 )。
必要な部品だけを取り付けてください。契約はサーバーなしでも使用できます。
ディストリビューションは afs_core / afs_server / afs_connector_sdk およびすべてとしてインポートされます
PEP 561 タイプです。パッケージ化、名前空間の決定、リリース フローについては、
ADR 0005 。リリース
Trusted Publishing 経由で vX.Y.Z タグの PyPI に公開する
( release.yml )。
事前構築されたイメージは、リリースごとに GHCR に公開されます ( v* タグ)。
:latest トラック最新のリリース。これらは Fargate 上で直接実行されます。
Kubernetes とローカル。 1 つの注意点: AWS Lambda は、ECR からのみプルできます。
同じアカウントなので、Lambda パスについては、最初にイメージを ECR にミラーリングします。
( docker で GHCR イメージをプルし、タグを付けてリポジトリにプッシュします)。ローカルで構築する
まだ動作します ( make dev 、または docker build )。労働者が取る
--build-arg AFS_EXTRAS=... より重いエクストラクター (例: docling )。
terraform/ は、レイヤーごとのモジュールとフットプリント全体をプロビジョニングします。
クイックスタートの例: 状態バックエンド、CI ロール、データ バケットと KMS、
カタログ テーブル、提供する Lambda と関数 URL、非同期

ジェスション (イベントブリッジ
→ SQS → ワーカー）、スケジュールされたリコンサイラー、および高信号 CloudWatch アラーム。それは
1 つのテラフォームが適用されます。 terraform/README.md から始めます。
Agentic-FS は、他の人が最初に公開したアイデアに基づいています。デザインは最も直接的です
インスピレーションを得たもの:
Mintlify: アシスタント用の仮想ファイルシステムを構築した方法。中核となる形状は、既存のストレージ上の仮想ファイル システム、クレームを枝刈りしたパス ツリー、2 段階の grep 、読み取り専用セマンティクス、サンドボックス コスト フレームワークから生まれました。
MCP を使用した効果的なコンテキスト エンジニアリングとコード実行に関する人間論。境界のあるコンテキスト効率の高いツールと MCP ファーストのサーフェス。
「Grepはフロアです。」クロードコードが grep のインデックスを削除する、および grep が埋め込みに勝る理由 (拡張) を参照してください。セマンティック検索はオプトイン アクセラレータとして機能し続けます。
私たちが構築するエコシステム: Model Context Protocol 、 fsspec (オブジェクト ストア アダプター)、LlamaHub/LlamaIndex (コネクタ アダプター)、および Docling (抽出)。
隣接する先行技術: Turso AgentFS 、 Onyx 、および Ragie 。
より完全な参照リストは docs/agentic-fs-oss-plan.md にあります。
docs/build-progress.md : ビルドされたもの、次の内容、ロードマップ。
docs/agentic-fs-oss-plan.md : 完全な設計。
docs/swap-guides/ および docs/decions/ : レイヤーごとのスワップと ADR。
AI エージェント向けのファイルシステム スタイルのドキュメントへのアクセス
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Filesystem-style access to your documents, for AI agents - vivekkhimani/agentic-fs

GitHub - vivekkhimani/agentic-fs: Filesystem-style access to your documents, for AI agents · GitHub
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
vivekkhimani
/
agentic-fs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
142 Commits 142 Commits .agents/ skills .agents/ skills .claude/ skills .claude/ skills .github .github docs docs packages packages terraform terraform web web .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.worker Dockerfile.worker LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml skills-lock.json skills-lock.json uv.lock uv.lock View all files Repository files navigation
Filesystem-style access to your documents, for AI agents, in your own AWS
account. list / glob / grep / tree / find / ranged read over your
documents in your S3 , exposed through MCP (and REST). It's multi-tenant,
deploys with one terraform apply , costs ~ $2/month idle , and every
stateful layer is swappable .
Status: early, in active development. v1.0.0 is published to PyPI and the
repo is public; contributions are welcome (see the
open issues and
docs/build-progress.md ). The full loop runs on
AWS : ingest, extract, catalog, and the MCP/REST read surface, with scheduled
heal-from-S3 and high-signal alarms. License: Apache-2.0. Background and
rationale live in docs/agentic-fs-oss-plan.md .
A bounded, scoped MCP tool surface. An agent explores a document corpus the way a
coding agent explores a repo, except over extracted document text, indexed at
scale, multi-tenant, and remote:
Navigate. fs_list , fs_tree , fs_glob , fs_find (by type, size, mtime, status).
Search. fs_grep : two-stage, bounded, with ripgrep-style filters.
Read. fs_read (ranged, or by section ), fs_outline (a doc's heading map), fs_tables , fs_diff .
Work. scratch_* (a per-principal workspace), whoami .
Every tool runs through one middleware that applies claims-filtered visibility,
scope enforcement, a per-call output budget, and an audit log
( ADR 0012 ). Adding a tool is a
registry entry, not a fork. Semantic fs_search is an optional accelerator on
the roadmap; grep is the floor.
Requirements: Docker ,
uv , and make
(macOS: xcode-select --install ).
git clone https://github.com/vivekkhimani/agentic-fs && cd agentic-fs
make dev # builds the image, starts MinIO + DynamoDB Local + the API, seeds the bucket/table
curl localhost:8080/v1/healthz # {"status":"ok","version":"..."}
curl localhost:8080/v1/me # the local dev principal
# Ingest a folder of documents, then read them back:
uv run fs-crawler --connector local --source ./docs --api-url http://localhost:8080 --namespace handbook
curl " localhost:8080/v1/fs/handbook/entries " # the catalog rows that appeared
The MCP surface is mounted at localhost:8080/mcp , so you can point any MCP
client at it. make down stops the stack and make clean also wipes the volumes.
The API is the same container image that runs on AWS Lambda and Fargate
( ADR 0003 ).
Local dev uses a static dev principal ( AFS_AUTH_MODE=dev ). Never run that
in production. In production agentic-fs is an OAuth 2.1 resource server : you
bring your own IdP (WorkOS, Cognito, Auth0, Okta, Keycloak), and afs auth doctor shows exactly how a token maps to a principal
( auth swap-guide , ADR 0013 ).
uv sync # set up the Python workspace (once)
make test # run the test suite
make lint # ruff lint + format check
make fmt # autoformat + autofix
make help # list every target
Every PR is gated by CI: Python (ruff + pytest) for packages/** , and
Terraform (fmt, validate, tflint, trivy) for terraform/** .
packages/
afs-core/ contracts (Protocols), DTOs, key scheme, conformance kits (pydantic only)
afs-server/ stores, services, extraction, FastAPI app + MCP mount (implements afs-core)
afs-connector-sdk/ fs-crawler CLI + sync engine + Local FS / S3 / Drive / LlamaHub connectors
terraform/ modular IaC: global state/CI roles, per-layer modules, examples
docs/ the plan, build progress, swap guides, decision records (ADRs)
Dockerfile one image: Lambda + Fargate + local
Swap any layer (plug-and-play)
Each layer sits behind a small contract with a conformance kit and a one-page
guide, so you can run it on the infrastructure you already have.
It works by a backend name in settings plus entry-point discovery
( ADR 0002 ).
Install only the parts you need. The contracts are usable without the server.
Distributions import as afs_core / afs_server / afs_connector_sdk , and all
are PEP 561 typed. Packaging, the namespace decision, and the release flow are in
ADR 0005 . Releases
publish to PyPI on a vX.Y.Z tag via Trusted Publishing
( release.yml ).
Prebuilt images are published to GHCR on each release ( v* tag):
:latest tracks the most recent release. These run directly on Fargate,
Kubernetes, and locally . One caveat: AWS Lambda can only pull from ECR in the
same account , so for the Lambda path, mirror the image into your ECR first
( docker pull the GHCR image, then tag + push to your repo). Building locally
still works too ( make dev , or docker build ); the worker takes
--build-arg AFS_EXTRAS=... for heavier extractors (e.g. docling ).
terraform/ provisions the whole footprint with per-layer modules and a
quickstart example: the state backend, CI roles, the data bucket and KMS, the
catalog table, the serving Lambda and Function URL, async ingestion (EventBridge
→ SQS → worker), the scheduled reconciler, and high-signal CloudWatch alarms. It's
one terraform apply . Start with terraform/README.md .
agentic-fs stands on ideas others published first. The design is most directly
inspired by:
Mintlify: How we built a virtual filesystem for our assistant . The core shape came from here: a virtual filesystem over existing storage, a claims-pruned path tree, two-stage grep , read-only semantics, and the sandbox cost framing.
Anthropic on effective context engineering and code execution with MCP . Bounded, context-efficient tools and the MCP-first surface.
"Grep is the floor." See Claude Code dropping indexing for grep and why grep beat embeddings (Augment) . Semantic search stays an opt-in accelerator.
The ecosystem we build on: the Model Context Protocol , fsspec (the object-store adapter), LlamaHub/LlamaIndex (the connector adapter), and Docling (extraction).
Adjacent prior art: Turso AgentFS , Onyx , and Ragie .
A fuller reference list lives in docs/agentic-fs-oss-plan.md .
docs/build-progress.md : what's built, what's next, the roadmap.
docs/agentic-fs-oss-plan.md : the full design.
docs/swap-guides/ and docs/decisions/ : per-layer swaps and ADRs.
Filesystem-style access to your documents, for AI agents
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/opencomplai/opencomplai"
hn_url: "https://news.ycombinator.com/item?id=49296432"
title: "Show HN: OpenComplAI – open-source EU AI Act compliance checks in CI/CD"
article_title: "GitHub - Opencomplai/opencomplai: Code-First AI Compliance. The open-source regulatory engine for the EU AI Act. Automate your risk assessments directly within your CI/CD pipeline. · GitHub"
author: "ythouma"
captured_at: "2026-08-14T09:58:27Z"
capture_tool: "hn-digest"
hn_id: 49296432
score: 3
comments: 0
posted_at: "2026-08-14T09:29:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenComplAI – open-source EU AI Act compliance checks in CI/CD

- HN: [49296432](https://news.ycombinator.com/item?id=49296432)
- Source: [github.com](https://github.com/opencomplai/opencomplai)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T09:29:23Z

## Translation

タイトル: Show HN: OpenComplAI – CI/CD でのオープンソース EU AI 法準拠チェック
記事のタイトル: GitHub - Opencomplai/opencomplai: コードファースト AI コンプライアンス。 EU AI 法のオープンソース規制エンジン。 CI/CD パイプライン内で直接リスク評価を自動化します。 · GitHub
説明: コードファースト AI コンプライアンス。 EU AI 法のオープンソース規制エンジン。 CI/CD パイプライン内で直接リスク評価を自動化します。 - オープンコンプライ/オープンコンプライ

記事本文:
GitHub - Opencomplai/opencomplai: コードファースト AI コンプライアンス。 EU AI 法のオープンソース規制エンジン。 CI/CD パイプライン内で直接リスク評価を自動化します。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オープンコンプライ
/
オープンコンプライ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
10 コミット 10 コミット .github/ workflows .github/

ワークフロー API API アセット アセット ドキュメント ドキュメントの例 例 インフラ インフラ パッケージ パッケージ スクリプト スクリプト サービス サービス 同期 同期 テスト/フィクスチャ/ eu_ai_scan テスト/フィクスチャ/ eu_ai_scan ツール/ verify-ledger ツール/ verify-ledger .gitattributes .gitattributes .gitignore .gitignore .ocignore .ocignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml .python-version .python-version .vercelignore .vercelignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md DOCS-UPDATE-PLAN.md DOCS-UPDATE-PLAN.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md conftest.py conftest.py mkdocs.yml mkdocs.yml package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml pyproject.toml pyproject.toml 要件-docs.txt 要件-docs.txt uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenComplAI: AI パイプラインのコードとしてのコンプライアンス
手動監査を停止→出荷を開始します。
OpenComplAI は、EU AI 法のコンプライアンスを CI/CD パイプラインに直接導入し、断片化された法的義務を自動化された機械可読な「出荷前チェック」に変換します。
OpenComplAI のデモを見る (MP4)
従来の GRC ツールは、「速度税」を生み出す接続されていないダッシュボードです。コンプライアンスを左にシフトします。
非準拠の防止: 安全ルールに違反するビルドをブロックすることでリリースをゲートします。
自動化された証拠: 導入ごとに監査対応のログを自動的に生成します。
フレームワークに依存しない: 進化する世界標準 (EU AI 法、NIST RMF、ISO) に適応するように構築されています。
仕組み (3 分間のセットアップ)
定義: モデルのコンプライアンス マニフェストを作成します。
統合: OpenComplAI アクションを GitHub/GitLab パイプラインに追加します。
船: 自動人形を入手する

コードが運用環境に入る前に「合格/不合格」の結果が得られます。
ダミー リポジトリ (サンドボックス) をチェックしてください – 運用コードを危険にさらさずに AI エラーを検出する方法をテストします。
opencomprai-core : コントロールを評価するルール エンジン。
opencomprai-cli : 開発環境でローカルにチェックを実行します。
opencomplai-sdk : コンプライアンスをカスタム内部ツールにプログラム的に埋め込みます。
最初のコンプライアンス チェックは 15 分以内に実行されます。
pip インストール opencomplai
これにより、CLI、コア ルール エンジン、および安定した API コントラクトを備えた SDK がインストールされます (「
終了コードとアーティファクトスキーマの保証については CHANGELOG)。
代わりにチェックアウトから作業したい寄稿者の場合は、ソースからインストールしてください。
core 、 cli 、および sdk-python パッケージは一緒にインストールする必要があります。
git clone https://github.com/Opencomplai/opencomplai
CDオープンコンプライ
pip install -e パッケージ/core -e パッケージ/cli -e パッケージ/sdk-python
# または、uv: uv sync
次に、最初の評価を実行します。
opencomplai init --system-id my-model --目的「カスタマー サポート チャットボット」
オープンコンプライチェック
または、最初にゼロセットアップ ( opencomprai scan --quick ) で試してください。検出のみを実行します
マニフェストを必要とせずにスキャンし、ビルドをゲートすることはありません。
opencomplai scan --quick 。
プリコミットフック
Opencomplai を独自の .pre-commit-config.yaml に追加して、クイック スキャン (または
すべてのコミットで完全なコンプライアンス ゲート (マニフェストを取得したら):
リポジトリ:
- リポジトリ : https://github.com/Opencomplai/opencomplai
リビジョン: v0.1.2
フック:
- id : opencomplai-quick-scan # 検出のみ、コミットは失敗しません
# - id: opencomplai-check # 完全な EU AI 法ゲート — system-manifest.json が必要
完全なドキュメントを表示
完全な Docker ベースのデプロイメントについては、次の文書に記載されています。
docs/src/deployment/quickstart.md 。
現在、クローズド ベータ パイロットを行っています。 AI エンジニアまたは ML プラットフォームの場合

リーダー様、フィードバックをお待ちしております。
開発者 Discord に参加してください — EU AI 法のワークフロー、パイプライン統合について話し合い、他の MLOps エンジニアとエンジンのストレス テストを行います
バグを報告する · 機能をリクエストする
LinkedIn · Reddit 研究コミュニティ
EU AI 法がお客様のシステムに適用されるかどうか、あるいはプロバイダーとデプロイヤーとしてどちらの義務を負うのかがわかりませんか?インタラクティブな EU AI Act Checker を使用します。これは、範囲、高リスク分類、GPAI、および義務をカバーするブラウザベースのウィザードです。アカウントは必要ありません。または、ローカルで実行します。
opencomplai checker --web # ホストされているドキュメント ページを開きます
opencomplai checker --web --local # 自己完結型のコピーをオフラインで提供します
アーキテクチャの概要
コンポーネント
種類
責任
コア
パッケージ
リスク評価プリミティブとポリシー マッピング ロジック (HTTP なし)。
クリ
パッケージ
ローカルチェックを実行し、ワークフローを調整するコマンドラインインターフェイス。
SDK-Python
パッケージ
コアをラップし、安定した統合サーフェスを提供する Python SDK。
ゲートウェイ-API
サービス
マルチサービス展開用の HTTP エントリポイント。リクエストの検証とルーティング。
リスクエンジン
サービス
サービスとしてのリスク分類の実行とルール評価。
証拠保管庫
サービス
不変性が保証され、コンテンツに対応したアーティファクトを備えた証拠ストレージ。
ドキュメントジェネレーター
サービス
保存された証拠からのドシエ/文書の生成 (例: Annex IV スタイルの出力)。
出口プロキシ
サービス
制御された外部接続のためのホワイトリストに登録された出力の強制。
リポジトリのレイアウト
オープンコンプライ/
§── パッケージ/
│ §── core/ # リスク評価エンジン — Python、Pydantic v2、HTTP なし
│ §── cli/ # CLI ツール — Typer + Rich、コアまたはゲートウェイ API を呼び出す
│ └── sdk-python/ # Python SDK — pip インストール可能、コアをラップ
§── サービス/
│ §──gateway-api/ # REST API — Node.js + TypeScript + Fast

ify (OpenAPI ファースト)
│ §──risk-engine/ # リスク分類サービス — Python + FastAPI
│ §──evidence-vault/ # 不変台帳 + CAS — Python + FastAPI + PostgreSQL
│ §── doc-generator/ # Annex IV ドシエジェネレーター — Python + FastAPI
│ └─ egress-proxy/ # 許可リストに登録された egress enforcer — Python + FastAPI
§── ツール/
│ └── verify-ledger/ # 証拠台帳連鎖検証ツール
§── インフラ/
│ §── docker/ # Dockerfiles (サービスごとに 1 つ)
│ §── compose/ # Docker Compose リファレンス デプロイメント + .env.example
│ └── 移行/ # Alembic データベースの移行
§── docs/ # MkDocs ドキュメント (GitHub Actions 経由で公開)
§── 例/ # 動作するコードの例
§── scripts/ # bootstrap.sh、doctor.py、verify-sbom.sh
━── .github/
§── ワークフロー/ # GitHub Actions CI ワークフロー
§── ISSUE_TEMPLATE/
└── pull_request_template.md
エディション
Community Edition — このリポジトリは、 AGPL-3.0 に基づいてライセンスされています。完全なリスク
評価エンジン、CLI、SDK、サービス、および EU AI Act チェッカー。
Enterprise Edition — ホストされたプレミアム ダッシュボード、シングル サインオン、追加
ルール エンジンと商用サポートは、商用ライセンスの下で利用可能です。参照
詳細については、complai.com をご覧ください。
開発セットアップ、ワークフロー規約、およびコードについては、CONTRIBUTING.md を参照してください。
スタイル。初心者向けの投稿を見つけるには、「良い最初の号」とラベル付けされた号を探してください。すべて
寄稿者は、寄稿者ライセンス契約に署名します。
Opencomplai の分類ロジックは完全に決定論的で、ルールベースです。 LLM または ML なし
推論は本番環境で使用されます。
このポリシーを適用するために、すべての依存関係ファイルが CI でスキャンされます。参照
完全なインベントリについては docs/security/ai-inventory.md を参照してください。
将来の AI 依存関係を承認するプロセス。
オープンコンプル

ai Community Edition は、GNU Affero General Public License v3.0 に基づいてライセンスされています。
(AGPL-3.0) — ライセンス を参照してください。 AGPL が適合しないユースケースについては、商用
ライセンスは Enterprise Edition の一部として利用できます。経由でお問い合わせください
opencomplai.com 。
コードファーストの AI コンプライアンス。 EU AI 法のオープンソース規制エンジン。 CI/CD パイプライン内で直接リスク評価を自動化します。
Readme AGPL-3.0 ライセンス 行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Code-First AI Compliance. The open-source regulatory engine for the EU AI Act. Automate your risk assessments directly within your CI/CD pipeline. - Opencomplai/opencomplai

GitHub - Opencomplai/opencomplai: Code-First AI Compliance. The open-source regulatory engine for the EU AI Act. Automate your risk assessments directly within your CI/CD pipeline. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Opencomplai
/
opencomplai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
10 Commits 10 Commits .github/ workflows .github/ workflows api api assets assets docs docs examples examples infra infra packages packages scripts scripts services services sync sync tests/ fixtures/ eu_ai_scan tests/ fixtures/ eu_ai_scan tools/ verify-ledger tools/ verify-ledger .gitattributes .gitattributes .gitignore .gitignore .ocignore .ocignore .pre-commit-config.yaml .pre-commit-config.yaml .pre-commit-hooks.yaml .pre-commit-hooks.yaml .python-version .python-version .vercelignore .vercelignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DOCS-UPDATE-PLAN.md DOCS-UPDATE-PLAN.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md conftest.py conftest.py mkdocs.yml mkdocs.yml package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml pyproject.toml pyproject.toml requirements-docs.txt requirements-docs.txt uv.lock uv.lock View all files Repository files navigation
OpenComplAI: Compliance-as-Code for AI Pipelines
Stop manual audits → Start shipping.
OpenComplAI brings EU AI Act compliance directly into your CI/CD pipeline, turning fragmented legal mandates into automated, machine-readable "Pre-Ship Checks."
Watch the OpenComplAI demo (MP4)
Traditional GRC tools are disconnected dashboards that create "velocity tax." We shift compliance left:
Prevent Non-Compliance: Gate releases by blocking builds that violate safety rules.
Automated Evidence: Generate audit-ready logs automatically for every deployment.
Framework-Agnostic: Built to adapt to evolving global standards (EU AI Act, NIST RMF, ISO).
How It Works (The 3-Minute Setup)
Define: Create a compliance manifest for your model.
Integrate: Add the OpenComplAI action to your GitHub/GitLab pipeline.
Ship: Get an automated "Pass/Fail" result before your code ever hits production.
Check out our Dummy Repo (Sandbox) – Test how we catch AI errors without risking your production code.
opencomplai-core : The rule engine that evaluates controls.
opencomplai-cli : Run checks locally in your dev environment.
opencomplai-sdk : Programmatically embed compliance into your custom internal tooling.
Get your first compliance check running in under 15 minutes :
pip install opencomplai
This installs the CLI, core rule engine, and SDK with a stable API contract (see
CHANGELOG for exit-code and artifact-schema guarantees).
For contributors who want to work from a checkout instead, install from source — the
core , cli , and sdk-python packages must be installed together:
git clone https://github.com/Opencomplai/opencomplai
cd opencomplai
pip install -e packages/core -e packages/cli -e packages/sdk-python
# or, with uv: uv sync
Then run a first assessment:
opencomplai init --system-id my-model --intended-purpose " customer support chatbot "
opencomplai check
Or try it with zero setup first — opencomplai scan --quick . runs a discovery-only
scan with no manifest required and never gates your build:
opencomplai scan --quick .
Pre-commit hook
Add Opencomplai to your own .pre-commit-config.yaml to run the quick scan (or the
full compliance gate, once you have a manifest) on every commit:
repos :
- repo : https://github.com/Opencomplai/opencomplai
rev : v0.1.2
hooks :
- id : opencomplai-quick-scan # discovery only, never fails the commit
# - id: opencomplai-check # full EU AI Act gate — requires system-manifest.json
View Full Documentation
Full Docker-based deployment is documented in
docs/src/deployment/quickstart.md .
We are currently in a Closed Beta Pilot . If you are an AI engineer or ML platform lead, we want your feedback.
Join our Developer Discord — discuss EU AI Act workflows, pipeline integration, and stress-test the engine with other MLOps engineers
Report a bug · Request a feature
LinkedIn · Reddit research community
Not sure whether the EU AI Act applies to your system, or which obligations you carry as a provider versus a deployer? Use the interactive EU AI Act Checker — a browser-based wizard covering scope, high-risk classification, GPAI, and obligations. No account needed. Or run it locally:
opencomplai checker --web # opens the hosted docs page
opencomplai checker --web --local # serves a self-contained copy offline
Architecture overview
Component
Kind
Responsibility
core
package
Risk assessment primitives and policy mapping logic (no HTTP).
cli
package
Command-line interface that runs local checks and orchestrates workflows.
sdk-python
package
Python SDK that wraps the core and provides a stable integration surface.
gateway-api
service
HTTP entrypoint for multi-service deployments; request validation and routing.
risk-engine
service
Risk classification execution and rules evaluation as a service.
evidence-vault
service
Evidence storage with immutability guarantees and content-addressed artifacts.
doc-generator
service
Dossier/document generation (e.g. Annex IV-style outputs) from stored evidence.
egress-proxy
service
Allowlisted egress enforcement for controlled external connectivity.
Repository layout
opencomplai/
├── packages/
│ ├── core/ # Risk assessment engine — Python, Pydantic v2, no HTTP
│ ├── cli/ # CLI tool — Typer + Rich, calls core or gateway-api
│ └── sdk-python/ # Python SDK — pip-installable, wraps core
├── services/
│ ├── gateway-api/ # REST API — Node.js + TypeScript + Fastify (OpenAPI-first)
│ ├── risk-engine/ # Risk classification service — Python + FastAPI
│ ├── evidence-vault/ # Immutable ledger + CAS — Python + FastAPI + PostgreSQL
│ ├── doc-generator/ # Annex IV dossier generator — Python + FastAPI
│ └── egress-proxy/ # Allowlisted egress enforcer — Python + FastAPI
├── tools/
│ └── verify-ledger/ # Evidence ledger chain verification tool
├── infra/
│ ├── docker/ # Dockerfiles (one per service)
│ ├── compose/ # Docker Compose reference deployment + .env.example
│ └── migrations/ # Alembic database migrations
├── docs/ # MkDocs documentation (published via GitHub Actions)
├── examples/ # Working code examples
├── scripts/ # bootstrap.sh, doctor.py, verify-sbom.sh
└── .github/
├── workflows/ # GitHub Actions CI workflows
├── ISSUE_TEMPLATE/
└── pull_request_template.md
Editions
Community Edition — this repository, licensed under AGPL-3.0 . The full risk
assessment engine, CLI, SDK, services, and EU AI Act checker.
Enterprise Edition — a hosted premium dashboard, single sign-on, additional
rule engines, and commercial support, available under a commercial licence. See
opencomplai.com for details.
See CONTRIBUTING.md for development setup, workflow conventions, and code
style. Look for issues labelled good first issue to find starter-sized contributions. All
contributors sign the Contributor Licence Agreement .
Opencomplai's classification logic is fully deterministic and rule-based. No LLM or ML
inference is used in production.
All dependency files are scanned in CI to enforce this policy. See
docs/security/ai-inventory.md for the full inventory and
the process for approving future AI dependencies.
Opencomplai Community Edition is licensed under the GNU Affero General Public Licence v3.0
(AGPL-3.0) — see LICENSE . For use cases that the AGPL does not fit, a commercial
licence is available as part of the Enterprise Edition; contact us via
opencomplai.com .
Code-First AI Compliance. The open-source regulatory engine for the EU AI Act. Automate your risk assessments directly within your CI/CD pipeline.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

---
source: "https://github.com/dimileeh/agent-workspace-fabric"
hn_url: "https://news.ycombinator.com/item?id=48554898"
title: "Show HN: AWF – run parallel AI coding agents, each in its own Docker workspace"
article_title: "GitHub - dimileeh/agent-workspace-fabric: Control plane that runs AI coding agents (Codex, Claude Code, Gemini, OpenCode, Grok, Cursor) as disciplined contributors: each task gets an isolated git worktree + Docker Compose stack, profile-driven validation, a created PR, and an autonomous PR-monitor l\n[truncated]"
author: "dimileeh"
captured_at: "2026-06-16T13:55:50Z"
capture_tool: "hn-digest"
hn_id: 48554898
score: 3
comments: 0
posted_at: "2026-06-16T13:21:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AWF – run parallel AI coding agents, each in its own Docker workspace

- HN: [48554898](https://news.ycombinator.com/item?id=48554898)
- Source: [github.com](https://github.com/dimileeh/agent-workspace-fabric)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T13:21:51Z

## Translation

タイトル: Show HN: AWF – それぞれ独自の Docker ワークスペースで並列 AI コーディング エージェントを実行する
記事のタイトル: GitHub - dimileeh/agent-workspace-fabric: AI コーディング エージェント (Codex、Claude Code、Gemini、OpenCode、Grok、Cursor) を規律あるコントリビューターとして実行するコントロール プレーン: 各タスクは、分離された git worktree + Docker Compose スタック、プロファイル駆動型検証、作成された PR、自律型 PR モニターを取得します。
[切り捨てられた]
説明: AI コーディング エージェント (Codex、Claude Code、Gemini、OpenCode、Grok、Cursor) を規律あるコントリビューターとして実行するコントロール プレーン: 各タスクは、分離された git ワークツリー + Docker Compose スタック、プロファイル駆動型検証、作成された PR、自律型 PR モニター ループ (レビュー、CI 修正、ベース同期、オートマー) を取得します。
[切り捨てられた]

記事本文:
GitHub - dimileeh/agent-workspace-fabric: AI コーディング エージェント (Codex、Claude Code、Gemini、OpenCode、Grok、Cursor) を規律あるコントリビューターとして実行するコントロール プレーン: 各タスクは、分離された git worktree + Docker Compose スタック、プロファイル駆動型検証、作成された PR、自律型 PR モニター ループ (レビュー、CI 修正、ベース同期、自動マージ) を取得します。 GitHub と BitBucket。 Apache-2.0。 · GitHub
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
別のタブでサインアウトしたか、

ウィンドウ。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ディミレー
/
エージェント-ワークスペース-ファブリック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ディミリー/エージェント-ワークスペース-ファブリック
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
449 コミット 449 コミット .awf .awf .claude .claude .github .github apps/ console apps/ console docker docker docs docs example/ awf-core-demo 例/ awf-core-demo 移行 移行 パッケージ化 パッケージ化計画 プラン スクリプト スクリプト スキル/ awf-scheduler スキル/ awf-scheduler src/ awf src/ awf テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md alembic.ini alembic.ini compose.yaml compose.yaml openapi.json openapi.json pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AWF オペレータ コンソール: 5 つのエージェントが 1 つのコードベースに対して並行して実行されます。それぞれが独自の分離されたワークスペースにあり、それぞれが独自の PR を通じて監視され、すべてが 1 つの管理されたキューを通じてマージされます。
▶️ 90 秒の説明をご覧ください — AWF が衝突せずに 1 台のマシン上でエージェント群を実行する方法をご覧ください。
📐 インタラクティブなアーキテクチャ図 · 🧭 概念と用語集 · 🚀 クイックスタート
AWF は、AI コーディング エージェント用の産業用ワークスペース ファブリックです。
Codex、Claude Code、Gemini、および将来のコーディング エージェントに再現可能な方法を提供します。
規律あるソフトウェア貢献者のように働くため: それぞれのタスク

孤立したものを取得します
ワークスペース、クリーンなチェックアウト、宣言されたサービス、検証、PR 作成、PR
レビュー監視、コメント修正ループ、マージゲート、アーティファクト、イベント、
掃除。
AWF はチャットボットでも、製品計画の頭脳でもありません。それは処刑です
プランナー、人間のオペレーター、または MCP クライアントの下の基板。の中に
ワークスペースを使用すると、具体的な実装計画のライフサイクルを強制できます。
AI コーディング エージェントはコードを作成できますが、生のエージェントの実行はスケールに合わせてスケールされません。
実際のエンジニアリングワークフロー。
ワークスペース ファブリックがないと、並列エージェント開発はすぐに困難に陥ります。
同じ運用上の失敗:
エージェントはローカル状態、資格情報、データベース、Docker ネットワーク、または
依存関係のキャッシュ。
タスクは古いベース ブランチに対するテストに合格し、マージ前に古くなります。
PR が最初は緑色に見えた後に、レビュー コメントが到着します。
CI の失敗とレビュー担当者のフィードバックには手動によるベビーシッターが必要です。
エージェントはブランチをプッシュしますが、コメント、競合、および競合の処理は人間に任せます。
マージの準備。
プロジェクト固有の設定がオーケストレーション コードに漏れ込みます。
失敗したワークスペースは、ログ、イベント、理由コードに記録されるため、検査が困難です。
散らばっている。
同じランナーは 1 つのプロジェクトに対してハードコードされており、別のプロジェクトで再利用することはできません。
Python、Node、Next.js、Docker Compose、Go、Java、C++、または Rust リポジトリ。
本当のボトルネックは、エージェントがファイルを編集できるかどうかではありません。ボトルネックは
多くのエージェントが実際のリポジトリで安全に作業できるかどうか
すべての PR を人間が手作業で監督します。
AWF は、1 つのコーディング タスクを永続的で観察可能なライフサイクルに変えます。
コントロール プレーン データベースにワークスペース行を作成します。
要求されたベース ブランチから分離された git ワークツリーを作成します。
プロジェクトのランタイムを記述するワークスペース プロファイルを解決します。
ワークスペースごとの Docker Compose スタックをレンダリングして起動します。
オプションで AWF 所有のプランを実行します -

> 実行 -> 反復を比較します。
選択したコーディング エージェントをワークスペース コンテナ内で実行します。
プロファイル検証フェーズと明示的なリクエスト検証コマンドを実行します。
プル リクエストをコミット、プッシュ、オープンします。
PR がマージされるか、閉じられるか、失敗するまで監視します。
同じエージェントを再度呼び出して、意味のあるレビュー コメントに対処します。
ログが利用可能な場合の CI エラーを修正します。
必要に応じて、ベース ブランチを PR ブランチに同期します。
最初のレビュー猶予期間を通じてレビュー担当者のタイミングを尊重します。
すべてのゲートを通過した後にのみ自動マージされます。
成功したワークスペースを破棄し、失敗したワークスペースを検査用に保存します。
プロジェクト固有の知識はワークスペース プロファイルに属します。 AWF コントロール プレーン
一般的なライフサイクルに関する懸念事項を所有しています: git 分離、エージェント実行、サービス
オーケストレーション、検証、アーティファクト、PR 作成、モニタリング、マージ セーフティ、および
掃除。
このリポジトリは、エージェント ワークスペース ファブリックのアルファ ローカル コアです。準備完了です
ホスト、GKE、マルチテナントでのローカル評価とドッグフーディング用
デプロイメントは将来のレイヤーに残ります。
単一の正規の /v1 名前空間を持つ FastAPI REST API。
ワークスペースの作成、コントロール、オペレーターの読み取り、メトリクス、
PRモニター採用。
ワークスペース、操作、およびイベントの SQLAlchemy コントロール プレーン モデル。
プロファイル主導のワークスペース解決。
ワークスペースごとの Docker Compose スタックの生成。
Codex、Claude Code、Cursor、Gemini、OpenCode、および Grok アダプター。
エージェント アダプタの中央のデフォルト モデル/エフォート マップ。
AWF 所有の計画 -> 実行 -> ライフサイクル ポリシーの比較。
一般的なフェーズベースの検証。
自動コメント処理と自動マージを備えた PR モニターを備えています。
人間がマージするまでワークスペースを維持する PR モニターのバリアントをリリース/同期します。
Python/Alembic マルチヘッド修復のためのマージ後のターゲット ブランチ調整。
最初の PR レビュー猶予期間

自動マージ前。
スケジュール設定と来歴の確認のための永続的なタスク ポリシー メタデータ ( task_class 、owned_pa​​ths )。
アクション不可能なボットステータスのコメントフィルタリング。
ワークスペースのタイムライン、ログ、アーティファクト、実行時スナップショット、検証
来歴、メトリクス、ロック、マージキュー検査。
クラウド バックエンドとホスト型コントロール プレーン。
マルチテナント認証、クラウド シークレット ブローカー、および強化されたネットワーク サンドボックス。
ローカル PR モニターとマージ セーフティを超えた完全なセマンティック マージ自動化
ゲートはすでに実装されています。
過去の最終状態については docs/awf_prd_v2.2.md
アルファを導いたPRD。
過去の MVP プランについては docs/PLAN_MVP.md。
過去の PR については docs/PLAN_PR_MONITOR.md
モニターのデザイン。
履歴については docs/PLAN_RELEASE_PR_SYNC.md
PR 同期デザインをリリースします。
ローカルの docs/AWF_CORE_TRUST_MODEL.md
コアの信頼境界と将来のオペレーター/アーキテクトの分割。
docs/AWF_LOCAL_CONTAINER_UID_STRATEGY.md
ローカル コントロール プレーン コンテナーの UID/GID 戦略とピラーごと
ルートバイデフォルトの決定の背後にある分析。
すでに Claude Code または Codex を使用している場合、最も早い方法は、エージェントに
AWF をインストールしてリポジトリをオンボードします。これは、
リポジトリのプロファイリングと緑色の煙。このプロンプトを貼り付けます (<PATH> を次のように置き換えます)
プロジェクトのパス):
このマシンにエージェント ワークスペース ファブリック (AWF) をセットアップし、リポジトリをオンボードします。
1. https://github.com/dimileeh/agent-workspace-fabric をクローンして読み取ります
何かを行う前に、skills/awf-scheduler/SKILL.md および docs/QUICKSTART.md を確認してください。
2. 前提条件を確認します (Docker の実行、UV、git)。 PR 自動化の場合は設定のみ
私のリポジトリの Forge に必要な認証: GitHub (github.com) には gh 認証が必要です。
Bitbucket (bitbucket.org) には BITBUCKET_API_TOKEN (および BITBUCKET_EMAIL のいずれかが必要)
または BITBUCKET_AUTH_MODE=bearer) が .env にあり、gh はありません。不足しているものがあれば停止し、
電話番号

私は — 推測しないでください。
3. ソース レーン経由でインストールします: uv tools install 。 --force、次に awf セットアップ、awf 開始、
そして awf サービスステータス --format をきれいにします。
4. <PATH> でプロジェクトをオンボードします: awf init <PATH> --write-profile --yes、その後
awfスモークラン --project <PATH> --mocked-local --format かなり綺麗です。
5. 模擬煙が緑色になったら停止し、プロファイルの概要を報告します。作成しないでください
私が尋ねない限り、実際のワークスペースを開くか、PR を開きます。
バンドルされた skill/awf-scheduler/SKILL.md (オペレータ スキル) を読み取ります。
AWF を駆動する) なので、そのステップが現在のコマンドを追跡します。
決定的で再現可能なインストールを実現するために、AWF には 3 つの実行可能な初回実行レーンがあります。
パブリックカールインストーラーレーンは、ホストされるインストーラー URL までリリースゲートされます。
マニフェスト、チェックサム、リリース アーティファクトが公開され、検証されます。
awf を PATH に置く package-manager レーンと virtualenv レーンの場合:
awf セットアップ
すごいスタート
awf サービス ステータス --format pretty
awf init <パス>
awfスモークラン --project <パス> --mocked-local --format pretty
awf start は、ローカル API、ワーカー、データベース、Web コンソールを次の場所で開始します。
http://127.0.0.1:3000 。 awf start --headless を使用してコンソールをスキップするか、
awf start --console-port 3333 を実行して、別のローカルホスト ポートを選択します。
グローバル ツール インストール レーンを使用したソース チェックアウトの場合は、チェックアウトから実行します。
UVツールをインストールします。 --force
awf setup --source-checkout " $PWD "
awf start --source-checkout " $PWD "
awf サービス ステータス --format pretty
awf init <パス>
awfスモークラン --project <パス> --mocked-local --format pretty
グローバル インストール レーンのないソース チェックアウトの場合は、チェックアウトから実行します。
UV 同期 --extra dev
uv run --python 3.12 --extra dev awf setup --source-checkout " $PWD "
uv run --python 3.12 --extra dev awf start --source-checkout " $PWD "
uv run --python 3.12 --extra dev awf サービスステータス --format pretty
UV 実行 --pytho

n 3.12 --extra dev awf init <パス>
uv run --python 3.12 --extra dev awfスモーク run --project < path > --mocked-local --format pretty
アップグレード パスやアンインストール パスを含む完全なレーン固有のコマンドについては、を参照してください。
クイックスタート 、 アップグレードガイド 、および
アンインストールガイド 。
Homebrew は、最初の安定したタグ付き PyPI/GitHub リリースと
式監査。まだ brew インストール パスに依存しないでください。
サポートされているクライアント サーフェス (v0.1)
REST、CLI、および MCP は、v0.1 でサポートされるクライアント サーフェスです。現在、AWF にはサポートされている Python SDK が同梱されていません。インテグレーターは、サポートされているサーフェスの 1 つを使用する必要があります (オペレーターの利便性のための CLI、またはコントロール プレーンのプログラムによるアクセスのための REST API など)。内部 AWF モジュール (awf.* やその他の内部パスなど) をインポートしてカスタム API クライアントを構築しないでください。これらは安定した公開契約の一部ではなく、予告なく変更される可能性があります。
既存の GitHub プル リクエストは、
REST、CLI、および MCP サーフェス。導入により、モニターが所有するワークスペースが作成されます。
コーディング エージェントを再実行せずに PR を開き、AWF に通常の PR を適用させます。
コメント、チェック、鮮度、マージ ポリシーのループを監視します。
オペレーター向けの PR Monitor 導入を参照してください。

[切り捨てられた]

## Original Extract

Control plane that runs AI coding agents (Codex, Claude Code, Gemini, OpenCode, Grok, Cursor) as disciplined contributors: each task gets an isolated git worktree + Docker Compose stack, profile-driven validation, a created PR, and an autonomous PR-monitor loop (review, CI fixes, base sync, auto-mer
[truncated]

GitHub - dimileeh/agent-workspace-fabric: Control plane that runs AI coding agents (Codex, Claude Code, Gemini, OpenCode, Grok, Cursor) as disciplined contributors: each task gets an isolated git worktree + Docker Compose stack, profile-driven validation, a created PR, and an autonomous PR-monitor loop (review, CI fixes, base sync, auto-merge). GitHub & BitBucket. Apache-2.0. · GitHub
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
dimileeh
/
agent-workspace-fabric
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dimileeh/agent-workspace-fabric
main Branches Tags Go to file Code Open more actions menu Folders and files
449 Commits 449 Commits .awf .awf .claude .claude .github .github apps/ console apps/ console docker docker docs docs examples/ awf-core-demo examples/ awf-core-demo migrations migrations packaging packaging plans plans scripts scripts skills/ awf-scheduler skills/ awf-scheduler src/ awf src/ awf tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md alembic.ini alembic.ini compose.yaml compose.yaml openapi.json openapi.json pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
The AWF operator console: five agents running in parallel against one codebase — each in its own isolated workspace, each monitored through its own PR, all merging through one governed queue.
▶️ Watch the 90-second explainer — how AWF runs an agent swarm on one machine without collisions.
📐 Interactive architecture diagram · 🧭 Concepts & glossary · 🚀 Quickstart
AWF is an industrial workspace fabric for AI coding agents.
It gives Codex, Claude Code, Gemini, and future coding agents a repeatable way
to work like disciplined software contributors: each task gets an isolated
workspace, a clean checkout, declared services, validation, PR creation, PR
review monitoring, comment-fix loops, merge gates, artifacts, events, and
cleanup.
AWF is not a chatbot and not a product-planning brain. It is the execution
substrate beneath a planner, a human operator, or an MCP client; inside a
workspace it can enforce a concrete implementation-plan lifecycle.
AI coding agents can write code, but raw agent execution does not scale to a
real engineering workflow.
Without a workspace fabric, parallel agent development quickly runs into the
same operational failures:
Agents share local state, credentials, databases, Docker networks, or
dependency caches.
A task passes tests against an old base branch and becomes stale before merge.
Review comments arrive after a PR initially looks green.
CI failures and reviewer feedback require manual babysitting.
Agents push branches but leave humans to handle comments, conflicts, and
merge readiness.
Project-specific setup leaks into the orchestration code.
Failed workspaces are hard to inspect because logs, events, and reason codes
are scattered.
The same runner is hard-coded for one project and cannot be reused for a
Python, Node, Next.js, Docker Compose, Go, Java, C++, or Rust repository.
The real bottleneck is not whether an agent can edit files. The bottleneck is
whether many agents can safely work on real repositories without requiring a
human to supervise every PR by hand.
AWF turns one coding task into a durable, observable lifecycle:
Create a workspace row in the control-plane database.
Create an isolated git worktree from the requested base branch.
Resolve a workspace profile that describes the project runtime.
Render and launch a per-workspace Docker Compose stack.
Optionally run AWF-owned Plan -> Execute -> Compare iterations.
Run the selected coding agent inside the workspace container.
Run profile validation phases and explicit request validation commands.
Commit, push, and open a pull request.
Monitor the PR until it is merged, closed, or failed.
Address meaningful review comments by invoking the same agent again.
Fix CI failures when logs are available.
Sync the base branch into the PR branch when needed.
Respect reviewer timing through an initial review grace window.
Auto-merge only after all gates pass.
Tear down successful workspaces and preserve failed ones for inspection.
Project-specific knowledge belongs in workspace profiles. The AWF control plane
owns generic lifecycle concerns: git isolation, agent execution, service
orchestration, validation, artifacts, PR creation, monitoring, merge safety, and
cleanup.
This repository is the alpha local Core of Agent Workspace Fabric. It is ready
for local evaluation and dogfooding, while hosted, GKE, and multi-tenant
deployments remain future layers.
FastAPI REST API with a single canonical /v1 namespace.
MCP server tools for workspace creation, controls, operator reads, metrics,
and PR monitor adoption.
SQLAlchemy control-plane models for workspaces, operations, and events.
Profile-driven workspace resolution.
Per-workspace Docker Compose stack generation.
Codex, Claude Code, Cursor, Gemini, OpenCode, and Grok adapters.
Central default model/effort map for agent adapters.
AWF-owned Plan -> Execute -> Compare lifecycle policy.
Generic phase-based validation.
Feature PR monitor with automated comment handling and auto-merge.
Release/sync PR monitor variants that keep workspaces alive until human merge.
Post-merge target-branch reconciliation for Python/Alembic multi-head repair.
Initial PR review grace period before auto-merge.
Durable task policy metadata ( task_class , owned_paths ) for scheduling and review provenance.
Non-actionable bot status comment filtering.
Workspace timelines, logs, artifacts, runtime snapshots, validation
provenance, metrics, locks, and merge-queue inspection.
Cloud backend and hosted control plane.
Multi-tenant authz, cloud secret broker, and hardened network sandbox.
Full semantic merge automation beyond the local PR monitor and merge-safety
gates already implemented.
docs/awf_prd_v2.2.md for the historical end-state
PRD that guided the alpha.
docs/PLAN_MVP.md for the historical MVP plan.
docs/PLAN_PR_MONITOR.md for historical PR
monitor design.
docs/PLAN_RELEASE_PR_SYNC.md for historical
release PR sync design.
docs/AWF_CORE_TRUST_MODEL.md for the local
Core trust boundary and future Operator/Architect split.
docs/AWF_LOCAL_CONTAINER_UID_STRATEGY.md
for the local control-plane container UID/GID strategy and per-pillar
analysis behind the root-by-default decision.
If you already use Claude Code or Codex, the fastest path is to let your agent
install AWF and onboard your repo — it's the only lane that ends with your
repository profiled and a green smoke. Paste this prompt (replace <PATH> with
your project's path):
Set up Agent Workspace Fabric (AWF) on this machine and onboard my repo.
1. Clone https://github.com/dimileeh/agent-workspace-fabric and READ
skills/awf-scheduler/SKILL.md and docs/QUICKSTART.md before doing anything.
2. Check prerequisites (Docker running, uv, git). For PR automation, configure only
the auth my repo's forge needs: GitHub (github.com) needs gh authenticated;
Bitbucket (bitbucket.org) needs BITBUCKET_API_TOKEN (and either BITBUCKET_EMAIL
or BITBUCKET_AUTH_MODE=bearer) in .env and no gh. If any are missing, STOP and
tell me — do not guess.
3. Install via the source lane: uv tool install . --force, then awf setup, awf start,
and awf service status --format pretty.
4. Onboard my project at <PATH>: awf init <PATH> --write-profile --yes, then
awf smoke run --project <PATH> --mocked-local --format pretty.
5. Stop when the mocked smoke is green and report the profile summary. Do not create
a real workspace or open a PR unless I ask.
It reads the bundled skills/awf-scheduler/SKILL.md (the operator skill for
driving AWF) so its steps track the current commands.
For a deterministic, reproducible install, AWF has three runnable first-run lanes.
The public curl installer lane is release-gated until its hosted installer URL,
manifest, checksums, and release artifacts are published and verified.
For package-manager and virtualenv lanes that put awf on PATH :
awf setup
awf start
awf service status --format pretty
awf init < path >
awf smoke run --project < path > --mocked-local --format pretty
awf start starts the local API, worker, database, and web console at
http://127.0.0.1:3000 . Use awf start --headless to skip the console or
awf start --console-port 3333 to choose another localhost port.
For the source checkout with global tool install lane, run from the checkout:
uv tool install . --force
awf setup --source-checkout " $PWD "
awf start --source-checkout " $PWD "
awf service status --format pretty
awf init < path >
awf smoke run --project < path > --mocked-local --format pretty
For the source checkout with no global install lane, run from the checkout:
uv sync --extra dev
uv run --python 3.12 --extra dev awf setup --source-checkout " $PWD "
uv run --python 3.12 --extra dev awf start --source-checkout " $PWD "
uv run --python 3.12 --extra dev awf service status --format pretty
uv run --python 3.12 --extra dev awf init < path >
uv run --python 3.12 --extra dev awf smoke run --project < path > --mocked-local --format pretty
For the full lane-specific commands, including upgrade and uninstall paths, see
Quickstart , Upgrade Guide , and
Uninstall Guide .
Homebrew is planned after the first stable tagged PyPI/GitHub release and a
formula audit; do not rely on a brew install path yet.
Supported Client Surfaces (v0.1)
REST, CLI, and MCP are the supported client surfaces for v0.1. AWF does not currently ship with a supported Python SDK. Integrators should use one of the supported surfaces (e.g., the CLI for operator convenience or the REST API for control-plane programmatic access). Please do not import internal AWF modules (such as awf.* or other internal paths) to build custom API clients, as they are not part of the stable public contract and are subject to change without notice.
Existing GitHub pull requests can be adopted into AWF monitoring through the
REST, CLI, and MCP surfaces. Adoption creates a monitor-owned workspace for the
open PR without re-running the coding agent, then lets AWF apply the normal PR
monitor loop for comments, checks, freshness, and merge policy.
See PR Monitor Adoption for the oper

[truncated]

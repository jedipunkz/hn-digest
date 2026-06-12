---
source: "https://github.com/pylit-ai/opendream"
hn_url: "https://news.ycombinator.com/item?id=48497602"
title: "Show HN: OpenDream – local-first memory with dreaming for AI agents"
article_title: "GitHub - pylit-ai/opendream: Local-first memory and dreaming automation for agents. · GitHub"
author: "planetceres"
captured_at: "2026-06-12T01:02:47Z"
capture_tool: "hn-digest"
hn_id: 48497602
score: 1
comments: 0
posted_at: "2026-06-11T23:01:29Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenDream – local-first memory with dreaming for AI agents

- HN: [48497602](https://news.ycombinator.com/item?id=48497602)
- Source: [github.com](https://github.com/pylit-ai/opendream)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T23:01:29Z

## Translation

タイトル: Show HN: OpenDream – AI エージェントのための夢を見るローカルファーストの記憶
記事のタイトル: GitHub - pylit-ai/opendream: エージェントのためのローカルファーストの記憶と夢の自動化。 · GitHub
説明: エージェント向けのローカルファーストの記憶と夢のような自動化。 - pylit-ai/オープンドリーム

記事本文:
GitHub - pylit-ai/opendream: エージェント向けのローカルファーストのメモリとドリーミングの自動化。 · GitHub
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
ピリットアイ
/
オープンドリーム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店 Ta

gs ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github .github .meta/ spec-adapters .meta/ spec-adapters ドキュメント ドキュメント フロントエンド フロントエンド opendream opendream スクリプト スクリプト テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRADEMARKS.md TRADEMARKS.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
セッション間で改善されるエージェント コンテキスト。
AI エージェント用のオープンなローカル ファースト メモリ。
何が起こったのかを記録し、重要なものを取得し、何が変わったのかを確認します。 OpenDream は、エージェントが境界、ソース、レビューを可視化しながら、プロジェクト、ツール、セッション間で有用なコンテキストを伝達できるように支援します。
クローズド製品は夢を見始めています。 OpenDream は、エージェントのメモリをオープン、ローカル、ポータブル、そしてレビュー可能にします。 「ドリーミング」とは、不透明なモデルの動作ではなく、バックグラウンドでのメモリのレビューとクリーンアップを意味します。
可観測性 UI には、最初に生の JSON を読み取らなくても、最近のエージェントのアクティビティ、選択されたメモリ、コンテキスト使用の監査記録、レビューの決定、およびワークスペースの健全性が表示されます。完全な明暗デモ カット: docs/showcase/ui-demos.md 。
複数のエージェント、1 つのワークスペース メモリ プレーン
Codex、Claude Code、Cursor、およびその他のエージェントは、セッション全体で貢献できます。
the same workspace. OpenDream は、これらのイベントと準備されたコンテキストを 1 つにまとめます
エージェントの属性を含むローカル ストア。これにより、後のエージェントが何を検査して再利用できるようになります。
以前は内蔵メモリを交換せずに使用していました。
セッション全体とフローの差分をご覧ください。
light MP4 ,
dark MP4 , or the
UI demo gallery .
実際のエージェント セッションの OpenDream メモリ
現在のステータス: OpenDream は

アルファソフトウェア。 Codex は現在最もテストされているパスです。他の統合は、各ホストのフック、ルール、コンテキスト ファイル、または CLI サーフェスに依存します。現在の評価はフィクスチャ主導のリリース チェックであり、広範な現実世界のベンチマークではありません。情報が古かったり、範囲が広すぎたり、範囲が間違っていたりすると、記憶に悪影響を及ぼす可能性があるため、レビュー可能性は完成した主張ではなく、製品目標の一部です。
uv ツール install opendream # または: pipx install opendream
opendream init --workspace 。
opendream verify activity-capture --workspace 。 --ターゲットが設定されました
opendream ステータス --workspace 。
opendream deactivate --workspace 。
使い捨てワークスペースで自己完結型ショーケースを実行します。
opendream eval showcase --scenario エージェント-ワークスペース-ショーケース --workspace .tmp/opendream-showcase
opendream 観察サーブ --workspace .tmp/opendream-showcase
インストールされている CLI に activate 、 verify 、 deactivate 、 semantic 、または
eval 、uv ツールアップグレード opendream でアップグレードするか、Git から再インストールします。後
アップグレード、opendream verify --help 、opendream semantic --help 、および
opendream eval --help は、インストールがドキュメントと一致しているかどうかを簡単にチェックします。
Git からの最先端 (ツールの環境を上書き): uv ツールのインストール --force "opendream @ git+https://github.com/pylit-ai/opendream.git" 。
インストール チェック、アップグレードに関する注意事項、既知の制限事項、クリーン ルームの出所をライブで確認する
docs/launch-readiness.md 内、
docs/known-limitations.md 、および
docs/provenance.md 。
OpenDream の初回実行パスは意図的に小さくなっています。これらのクリップは、コマンド リファレンス ドキュメントの傍らに役立ちます。上記の UI アプリケーション ツアーは、一見するとわかりやすいでしょう。
UI デモ ギャラリー: docs/showcase/ui-demos.md 。 CLI デモ ギャラリー: docs/showcase/cli-demos.md 。
ブランドモーションアセット:frontend/assets/animation/README.md 。
PyPI (公開後に推奨) — システム Pyt を回避するために分離ツール env を使用します。

本制限 (PEP 668):
UVツールオープンドリームをインストールする
# または: pipx install opendream
opendream --help
Git から (同じ考え; インストーラーが許可する場合は @main / @v0.1.0 でピン留めします):
uvツールインストール「opendream @ git+https://github.com/pylit-ai/opendream.git」
# または: pipx install git+https://github.com/pylit-ai/opendream.git
リポジトリのチェックアウト (寄稿者):
make setup
.venv/bin/opendream --help
手動での同等のコマンド: python3 -m venv .venv && .venv/bin/pip install -e 。 from the repo root.モジュール フォールバック: python3 -m opendream.cli --help 。
python3 が見つからない場合は、python.org または OS パッケージ マネージャーからインストールします。
システムの依存関係 — OpenDream は jq を使用してフック スクリプト内の JSON を解析します (クロード コードの統合のため)。 OS パッケージ マネージャー経由でインストールします。
# macOS
brew install jq
# Ubuntu / Debian
sudo apt-get install jq
# Other systems
# 参照: https://jqlang.github.io/jq/download/
Integration at a glance
OpenDream は、アクティベーション優先の CLI です。
Codex は現在最もテストされているパスです。他の統合では、ホストのサポートに応じて、CLI コマンド、ルール、フック、または生成されたファイルを通じてコン​​テキストを公開します。組み込みターゲットは、OpenDream が統合サーフェスを生成できることを意味します。すべてのホストがフックをネイティブに実行するという意味ではありません。
既存のエージェント メモリを交換する必要はありません。 OpenDream は別個のワークスペース ローカル ストアを保持し、Codex、Claude Code、または別のエージェントの組み込みメモリと一緒に実行できます。アクティベーションにより、OpenDream が管理する統合サーフェスが追加されます。 activate-plan でプレビューするか、 init --no-activate-configured でスキップするか、後で deactivate で削除します。 See the FAQ .
opendream init --workspace 。
opendream verify activity-capture --workspace 。 --ターゲットが設定されました
opendream ステータス --workspace 。
opendream activate --workspace 。 --repair
opendream deactivate --workspace 。
の

下位レベルのランタイムも引き続き使用できますが、それは主要なメンタル モデルではありません。
エージェント統合の詳細 (ワークスペースと cwd、memory_layout 、 empty_reason /hinds 、JSON バージョン): docs/agent-integrations.md 。
推奨されるアクティベーション ワークフロー
opendream init --workspace 。 — メモリ レイアウトを作成し、設定されたエージェント サーフェスをデフォルトでアクティブにします。
opendream アクティベーション プラン --workspace 。 --targets の設定 — オプションのドライラン: どのサーフェスが検出されたかを確認します。 --targets all-supported を使用して、すべての組み込みエージェント ターゲットをプレビューします。
opendream verify activity-capture --workspace 。 --targets が構成されています — 選択したワークスペースで生成されたフックを実行し、診断メモリ キャプチャを証明します。
opendream ステータス --workspace 。 — 次のアクションで最新のキャプチャ検証状態 ( Never_run 、 Passed 、または failed ) を報告します。
opendream activate --workspace 。 --repair — ドリフトした管理ファイルとフックエントリを復元します。
オープンドリームドクター -- ワークスペース 。 --surface エージェント — 診断メモリに書き込まずに管理対象サーフェスを検査します。
まだ検出されていないツールの場合は、次のように実行します。 opendream activate --workspace 。 --targets カーソルを 1 回実行して、.cursor/rules/opendream.mdc およびフック スクリプトを作成します。
命令のみのターゲット (カーソル ルール、 .github/copilot-instructions.md 、 .hermes.md ) には、Codex サンプル ターゲットと同じ汎用のプレ/ポスト シェル フックが同梱されています。ホストにネイティブ OpenDream フックがない場合でも、エージェントはこれらのコマンドを実行する必要があります。
CLI でさらに多くのコマンドが公開されている場合でも、maintenance を文書化されたメンテナンス エントリポイントとして扱います。
ファーストパーティのサーフェス = この CLI。フック/スクリプト グルーは、ユーザーが追加しない限り、オペレーターが所有します。
このリポジトリにはファーストパーティ MCP サーバーはありません。 docs/mcp/servers.md は、出荷されたサーバーではなく、MCP をインベントリするためのテンプレートです。
人に面した行動については、この README と AGENTS.md で説明されています。

機械可読なランタイム コントラクトは、 opendream/schema/ の下に存在します。
Workspace, Runtime, and Observability
OpenDream は、ワークスペースごとに正規の状態を .opendream/ に保存します。 The multi-workspace catalog is a derived, machine-local convenience index;ワークスペース メモリは信頼できる情報源であり、スキャンは設定したルート上でのみ実行されます。
opendreamワークスペースリスト
opendream Observserve --workspace " $PWD " --port 8000
セマンティックファースト モードは構成姿勢であり、自動準備要求ではありません。選択したセマンティック パスが実行できない場合、OpenDream は理由と次のアクションとともに劣化状態を報告します。 prepare-context は、リコールされたコンテキストが検査可能な状態を維持できるように、引き続き段階的な開示と枝刈り証拠を使用する必要があります。
ワークスペース カタログのセットアップ、アップグレード/バックフィル フロー、ランタイム コマンド、生成されたデータの場所、グローバル/プロジェクト ストア、セマンティック モード、自動化、評価メモ、完全な CLI 例については、詳細なオペレーター ガイドを使用してください: docs/operator-guide.md 。
ドクター
目的
docs/technical-notes/dreaming-memory-change-control.md
Canonical memory lifecycle and review model
docs/comparison.md
静的ファイル、RAG、ホスト型 API、マネージド メモリとの機能レベルの比較
docs/security-local-first.md
ローカルファーストのデフォルト、明示的なプロバイダー パス、および信頼境界
docs/design-partner-workloads.md
メモリがどこで役立つか、どこで失敗するかを検証するために必要な実際のワークフロー
docs/FAQ.md
導入、共存、プライバシー、ベンチマーク、セットアップに関する質問
docs/showcase/memory-showcase.md
90秒の記憶デモ
docs/agent-integrations.md
エージェント統合ガイド
docs/operator-guide.md
ワークスペース カタログ、ランタイム コマンド、可観測性、生成されたデータ、評価メモ、および完全な CLI サンプル
docs/automation/dream-task-playbook.md
Automation and recurring memory tasks
docs/architecture/overview.md
アーキテクチャの概要

docs/benchmarks/methodology.md
ベンチマーク方法論
docs/claims.md
証拠に裏付けられた主張
docs/known-limitations.md
Known limitations
docs/provenance.md
クリーンルームの由来
docs/contributor-reference.md
コントリビューターのスモーク テスト、検証ターゲット、リリース ノート、および完全な CLI サンプル
セキュリティ.md
セキュリティポリシー
貢献.md
寄稿者ガイド
OpenDream does not send telemetry by default. Provider/API-key paths are
オペレーターが構成した実行パスであり、分析や追跡とは別のものです。
同期する
検証する
make release-check
コントリビューターのワークフロー、スモーク テスト、自動化の例、検証ターゲット、およびメンテナーのリリース ノートは docs/contributor-reference.md にあります。 Security reports go through SECURITY.md ;一般的な貢献に関するガイダンスは COTRIBUTING.md にあります。
エージェントのためのローカルファーストの記憶と夢のような自動化。
Apache-2.0 license
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
2
オープンドリーム 0.3.21
最新の
2026 年 6 月 9 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。お願いします

[切り捨てられた]

## Original Extract

Local-first memory and dreaming automation for agents. - pylit-ai/opendream

GitHub - pylit-ai/opendream: Local-first memory and dreaming automation for agents. · GitHub
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
pylit-ai
/
opendream
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github .meta/ spec-adapters .meta/ spec-adapters docs docs frontend frontend opendream opendream scripts scripts tests tests .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TRADEMARKS.md TRADEMARKS.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Agent context that improves between sessions.
Open, local-first memory for AI agents.
Capture what happened, retrieve what matters, and review what changed. OpenDream helps agents carry useful context across projects, tools, and sessions while keeping boundaries, sources, and review visible.
Closed products are starting to dream. OpenDream makes agent memory open, local, portable, and reviewable. "Dreaming" means background memory review and cleanup, not an opaque model behavior.
The observability UI shows recent agent activity, selected memories, context-use audit records, review decisions, and workspace health without making you read raw JSON first. Full light and dark demo cuts: docs/showcase/ui-demos.md .
Multiple agents, one workspace memory plane
Codex, Claude Code, Cursor, and other agents can contribute across sessions in
the same workspace. OpenDream keeps those events and prepared contexts in one
local store with agent attribution, so a later agent can inspect and reuse what
came before without replacing its built-in memory.
Watch the full session and diff flow:
light MP4 ,
dark MP4 , or the
UI demo gallery .
OpenDream memory in real agent sessions
Current status: OpenDream is alpha software. Codex is currently the most tested path; other integrations depend on each host's hooks, rules, context files, or CLI surface. Current evals are fixture-driven release checks, not broad real-world benchmarks. Memory can hurt when it is stale, over-broad, or incorrectly scoped, so reviewability is part of the product goal rather than a finished claim.
uv tool install opendream # or: pipx install opendream
opendream init --workspace .
opendream verify activation-capture --workspace . --targets configured
opendream status --workspace .
opendream deactivate --workspace .
Run the self-contained showcase in a disposable workspace:
opendream eval showcase --scenario agent-workspace-showcase --workspace .tmp/opendream-showcase
opendream observe serve --workspace .tmp/opendream-showcase
If your installed CLI does not show activate , verify , deactivate , semantic , or
eval , upgrade with uv tool upgrade opendream or reinstall from Git. After
upgrading, opendream verify --help , opendream semantic --help , and
opendream eval --help are quick checks that your install matches the docs.
Bleeding-edge from Git (overwrites the tool env): uv tool install --force "opendream @ git+https://github.com/pylit-ai/opendream.git" .
Install checks, upgrade notes, known limitations, and clean-room provenance live
in docs/launch-readiness.md ,
docs/known-limitations.md , and
docs/provenance.md .
OpenDream's first-run path is intentionally small. These clips are useful beside command reference docs; the UI application tour above is the better first look.
UI demo gallery: docs/showcase/ui-demos.md . CLI demo gallery: docs/showcase/cli-demos.md .
Brand motion assets: frontend/assets/animation/README.md .
PyPI (recommended once published) — use an isolated tool env to avoid system Python restrictions (PEP 668):
uv tool install opendream
# or: pipx install opendream
opendream --help
From Git (same idea; pin with @main / @v0.1.0 where your installer allows):
uv tool install " opendream @ git+https://github.com/pylit-ai/opendream.git "
# or: pipx install git+https://github.com/pylit-ai/opendream.git
Repo checkout (contributors):
make setup
.venv/bin/opendream --help
Manual equivalent: python3 -m venv .venv && .venv/bin/pip install -e . from the repo root. Module fallback: python3 -m opendream.cli --help .
If python3 is missing, install from python.org or your OS package manager.
System dependencies — OpenDream uses jq to parse JSON in hook scripts (for Claude Code integration). Install via your OS package manager:
# macOS
brew install jq
# Ubuntu / Debian
sudo apt-get install jq
# Other systems
# See: https://jqlang.github.io/jq/download/
Integration at a glance
OpenDream is an activation-first CLI .
Codex is currently the most tested path. Other integrations expose context through CLI commands, rules, hooks, or generated files depending on host support. A built-in target means OpenDream can generate the integration surface; it does not mean every host executes hooks natively.
You do not need to replace existing agent memory. OpenDream keeps a separate workspace-local store and can run alongside built-in memory from Codex, Claude Code, or another agent. Activation adds OpenDream-managed integration surfaces; preview them with activation-plan , skip them with init --no-activate-configured , or remove them later with deactivate . See the FAQ .
opendream init --workspace .
opendream verify activation-capture --workspace . --targets configured
opendream status --workspace .
opendream activate --workspace . --repair
opendream deactivate --workspace .
The lower-level runtime remains available, but it is not the main mental model.
Agent integration details (workspace vs cwd, memory_layout , empty_reason / hints , JSON version): docs/agent-integrations.md .
Recommended activation workflow
opendream init --workspace . — create the memory layout and activate configured agent surfaces by default.
opendream activation-plan --workspace . --targets configured — optional dry-run: see which surfaces are detected. Use --targets all-supported to preview every built-in agent target.
opendream verify activation-capture --workspace . --targets configured — run generated hooks in the selected workspace and prove diagnostic memory capture.
opendream status --workspace . — report the latest capture verification state ( never_run , passed , or failed ) with the next action.
opendream activate --workspace . --repair — restore drifted managed files and hook entries.
opendream doctor --workspace . --surface agents — inspect managed surfaces without writing diagnostic memory.
For a tool that was not detected yet, run e.g. opendream activate --workspace . --targets cursor once to create .cursor/rules/opendream.mdc and hook scripts.
Instruction-only targets (Cursor rules, .github/copilot-instructions.md , .hermes.md ) ship the same generic pre/post shell hooks as the Codex example target; the agent must still run those commands when the host has no native OpenDream hooks.
Treat maintain as the documented maintenance entrypoint even if the CLI exposes more commands.
First-party surface = this CLI. Hook/script glue is operator-owned unless you add it.
No first-party MCP server in this repo; docs/mcp/servers.md is a template for inventorying MCP, not a shipped server.
Human-facing behavior is described in this README and in AGENTS.md . Machine-readable runtime contracts live under opendream/schema/ .
Workspace, Runtime, and Observability
OpenDream stores canonical state per workspace under .opendream/ . The multi-workspace catalog is a derived, machine-local convenience index; workspace memory stays the source of truth and scans only run on roots you configure.
opendream workspace list
opendream observe serve --workspace " $PWD " --port 8000
Semantic-first mode is a configuration posture, not an automatic readiness claim. If the selected semantic path cannot run, OpenDream reports a degraded state with a reason and next action; prepare-context should still use progressive disclosure and pruning evidence so recalled context stays inspectable.
Use the detailed operator guide for workspace catalog setup, upgrade/backfill flows, runtime commands, generated data locations, global/project stores, semantic mode, automation, eval notes, and full CLI examples: docs/operator-guide.md .
Doc
Purpose
docs/technical-notes/dreaming-memory-change-control.md
Canonical memory lifecycle and review model
docs/comparison.md
Capability-level comparison with static files, RAG, hosted APIs, and managed memory
docs/security-local-first.md
Local-first defaults, explicit provider paths, and trust boundaries
docs/design-partner-workloads.md
Real workflows needed to validate where memory helps or fails
docs/FAQ.md
Adoption, coexistence, privacy, benchmarks, and setup questions
docs/showcase/memory-showcase.md
90-second memory demo
docs/agent-integrations.md
Agent integration guide
docs/operator-guide.md
Workspace catalog, runtime commands, observability, generated data, eval notes, and full CLI examples
docs/automation/dream-task-playbook.md
Automation and recurring memory tasks
docs/architecture/overview.md
Architecture overview
docs/benchmarks/methodology.md
Benchmark methodology
docs/claims.md
Evidence-backed claims
docs/known-limitations.md
Known limitations
docs/provenance.md
Clean-room provenance
docs/contributor-reference.md
Contributor smoke tests, verification targets, release notes, and full CLI examples
SECURITY.md
Security policy
CONTRIBUTING.md
Contributor guide
OpenDream does not send telemetry by default. Provider/API-key paths are
operator-configured execution paths and are separate from analytics or tracking.
make sync
make verify
make release-check
Contributor workflow, smoke tests, automation examples, verification targets, and maintainer release notes live in docs/contributor-reference.md . Security reports go through SECURITY.md ; general contribution guidance lives in CONTRIBUTING.md .
Local-first memory and dreaming automation for agents.
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
2
OpenDream 0.3.21
Latest
Jun 9, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please

[truncated]

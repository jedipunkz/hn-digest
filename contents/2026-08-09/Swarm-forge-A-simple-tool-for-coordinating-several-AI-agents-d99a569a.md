---
source: "https://github.com/unclebob/swarm-forge"
hn_url: "https://news.ycombinator.com/item?id=49229904"
title: "Swarm-forge: A simple tool for coordinating several AI agents"
article_title: "GitHub - unclebob/swarm-forge: A simple tool for coordinating several AI agents. · GitHub"
author: "tosh"
captured_at: "2026-08-09T10:22:53Z"
capture_tool: "hn-digest"
hn_id: 49229904
score: 1
comments: 0
posted_at: "2026-08-09T09:44:39Z"
tags:
  - hacker-news
  - translated
---

# Swarm-forge: A simple tool for coordinating several AI agents

- HN: [49229904](https://news.ycombinator.com/item?id=49229904)
- Source: [github.com](https://github.com/unclebob/swarm-forge)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T09:44:39Z

## Translation

タイトル: Swarm-forge: 複数の AI エージェントを調整するためのシンプルなツール
記事タイトル: GitHub - Unclebob/swarm-forge: 複数の AI エージェントを調整するためのシンプルなツール。 · GitHub
説明: 複数の AI エージェントを調整するためのシンプルなツール。 GitHub でアカウントを作成して、unclebob/swarm-forge の開発に貢献してください。

記事本文:
GitHub - Unclebob/swarm-forge: 複数の AI エージェントを調整するためのシンプルなツール。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ボブおじさん
/
スウォームフォージ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
258 コミット 258 コミット swarmforge swarmforge test/ swarmforge test/ swarmforge .gitignore .gitignore README.md README.md bb.edn bb.edn close-swarm close-swarm すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SPしないでください

Bankrbot SWARM トークンでお金を終わらせます。
AI エージェントの群れを信頼できるプロのソフトウェア エンジニアに変える、規律ある tmux ベースのエージェント オーケストレーション プラットフォーム。
このメイン ブランチはドキュメンタリーです。システムについて説明し、共有の運用スクリプトとデフォルトの憲法の記事を掲載しています。実行可能なワークフロー ブランチには、プロジェクト向けの構成、ロール プロンプト、および特定のワークフローを定義するローカル構成の記事が含まれています。
SwarmForge は、異なる git ワークツリーで動作するエージェント間の通信を容易にするエージェント調整システムです。
役割固有のプロンプト、ワークツリーの割り当て、tmux セッション、およびメッセージ パッシングの共有構造を提供するため、複数のエージェントが互いに干渉することなく同じプロジェクトで共同作業できます。
実行可能な SwarmForge 構成は専用のブランチ上に存在します。各ブランチには、 swarmforge/swarmforge.conf 、ローカル憲法の記事、および 1 つのワークフローのロール プロンプトが含​​まれています。起動時に、その ./swarm ラッパーは、共有操作スクリプトと共有構成記事がまだ存在しない場合はメインからコピーし、そのブランチのローカル構成を起動します。
two-pack は、迅速なバックエンド ワークフローです。バックエンドのリファクタリングと強化を維持しながら、Gherkin や受け入れテストのオーバーヘッドなしで高速コーディングの恩恵を受ける小規模なタスクに使用します。
コーダーは、TDD と単体テストを使用して要求された動作を実装します。
クリーナーはコーダーのハンドオフをバッチ処理し、クリーンアップ、CRAP および DRY レビュー、アーキテクチャ レビュー、カプセル化と懸念事項の分離の修正、および言語変更の強化を実行します。
通常の流れは、コーダー -> クリーナー -> コーダー です。このブランチは、仕様、QA、プロパティ テスト、または受け入れテストの役割を持たない厳密な実装/改良ループが必要な場合に使用します。
4-

Pack はコンパクトな仕様のワークフローです。すべての品質ゲートを独自のエージェントに分割せずに、Gherkin 仕様とアーキテクチャ上の考慮事項を必要とする中程度のプロジェクトに使用します。
指定子は、ユーザーの意図を正確な Gherkin 受け入れ仕様に変換し、ハンドオフの前に承認を求めます。
コーダーは、TDD、単体テスト、および生成された受け入れテストを使用して、承認された動作スライスを実装します。
リファクタラーは、動作を保持するクリーンアップ、カバレッジの改善、CRAP および DRY レビュー、変異サイトのスキャン、およびプロパティ テストのサポートを実行します。
アーキテクトは、高レベルの構造、依存関係の方向、突然変異強化、DRY レビュー、ソフト ガーキン突然変異、および最終完了通知を所有します。
通常の流れは、指定子 -> コーダー -> リファクタリング -> アーキテクト -> 指定子 です。クリーンアップ、アーキテクチャ、強化、QA を個別のエージェントに分割せずに、規律ある開発を行う場合は、このブランチを使用します。
シックスパックは完全なワークフローです。完全な仕様、事前の QA、バックエンドの検証、および重要なアーキテクチャ上の考慮事項を必要とする主要なプロジェクトに使用してください。主要な品質ゲートをそれぞれ独自の役割に分割します。
specifier は、ユーザーの意図を受け入れられた Gherkin 仕様とエンドツーエンドの QA 手順に変換します。
コーダーは、TDD、単体テスト、および生成された受け入れテストを使用して、承認された動作スライスを実装します。
クリーナーは、ローカルの動作を保持するクリーンアップ、カバレッジの改善、CRAP と DRY のレビュー、および変異サイトのスキャンを実行します。
アーキテクトは、モジュールの構造、境界、依存関係の方向、プロパティ テストの範囲をレビューします。
Hardender は、突然変異強化、言語突然変異、CRAP および DRY 検証、およびソフト Gherkin 突然変異を実行します。
QA は、指定子の QA プロシージャを実行可能なスクリプトに変換し、最終的なユーザー インターフェイス検証を実行し、ハンドオフの一貫性をチェックし、

完了通知を送信します。
通常の流れは、指定者 -> コーダー -> クリーナー -> アーキテクト -> ハードダー -> QA -> 完了です。各レビューおよび検証に関する懸念事項を別のエージェントが所有する場合は、このブランチを使用します。
SwarmForge はローカルで実行されます。実行可能ブランチを開始する前に、ターゲット マシンに以下があることを確認してください。
少なくとも 1 つの構成済みエージェント バックエンド ( codex 、 claude 、 copilot 、 grok など)
SwarmForge を使用するディレクトリで、実行可能なブランチを選択し、Git リモートを作成せずにその内容をプルします。
ブランチ=4 パック
カール -L " https://github.com/unclebob/swarm-forge/archive/refs/heads/ ${BRANCH} .tar.gz " | tar -xz --strip-components=1
2 エージェントのクイック ワークフローには BRANCH=two-pack、コンパクトな仕様ワークフローには BRANCH=four-pack、または完全な 6 エージェント ワークフローには BRANCH=six-pack を使用します。このコマンドには main を使用しないでください。 main はドキュメントであり、共有操作スクリプトを保存します。一方、実行可能なブランチは、プロジェクト向けの構成とプロンプトを提供します。
実行可能なブランチをコピーした後、ターゲット プロジェクトから swarm を開始します。
./swarm
./swarm ラッパーは、実行可能なブランチを小さく保ちます。最初の使用時に swarmforge/scripts/ が見つからない場合は、メイン ブランチ アーカイブをダウンロードし、共有操作スクリプトを swarmforge/scripts/ からコピーし、共有憲法記事を swarmforge/constitution/articles/ からステージングして、 swarmforge/scripts/swarmforge.sh を起動します。その後の実行では、既存のローカル スクリプト ディレクトリを上書きせずに再利用します。
ウィンドウは自動的に開きます。
swarm を停止するには、 swarmforge/swarmforge.conf にリストされている最初のウィンドウを閉じます。このクリーンアップ ウィンドウは tmux セッションをシャットダウンし、残りの追跡ウィンドウを閉じます。
Swarm がアクティブである間、SwarmForge はホストがスリープしないようにしようとします

。 macOS ではカフェインを使用します。 Linux では、利用可能な場合は systemd-inhibit を使用します。 OS によっては、ディスプレイのロックや手動スリープによってエージェントが中断される場合があります。この動作を無効にするには、./swarm の前に SWARMFORGE_PREVENT_SLEEP=0 を設定します。
SwarmForge は、次のような軽量の tmux ベースのオーケストレーション レイヤーです。
プロジェクトローカルの swarmforge/swarmforge.conf から構成主導の swarm を起動します。
構成されたロールごとに 1 つの tmux セッションを作成し、選択したバックエンドがサポートしている場合はロールごとにターミナル サーフェスを開きます
プロジェクトローカルの swarmforge/roles/<role>.prompt ファイルと階層化された swarmforge/constitution.prompt から動作を読み取ります。
claude 、 codex 、 copilot 、 grok などのロールごとのバックエンドをサポートします
アクティブな swarm 通信用のハンドオフ ヘルパーを含む、共有 swarmforge/scripts/ ディレクトリを各エージェントの PATH に配置します。
専用のワークツリー名に割り当てられたロールの git worktree を .worktrees/ の下に作成します
必要に応じて、新しい作業ディレクトリで git リポジトリを初期化します。
すべての swarm 状態を .swarmforge/ の作業ディレクトリに対してローカルに保持します。
Config-Driven Topology — swarm の形状は、ハードコードされたシェル変数ではなく、 swarmforge/swarmforge.conf から取得されます。
プロジェクトのローカル ロール — 各ロールは、オーケストレーションされている作業ツリーの swarmforge/roles/<role>.prompt によって定義されます。
階層化された憲法 — swarmforge/constitution.prompt は、エージェントに swarmforge/constitution/articles/ にある記事ファイルを読むように指示します。
ロールごとのバックエンドの選択 — ロールは claude 、 codex 、 copilot 、または grok を起動できます。
Observable Swarm — ロールごとに 1 つのターミナル ウィンドウを開き、リアルタイムでセッションを監視します。
セルフホスト型で軽量 — 最小限のマシンを使用して tmux およびターミナルでローカルに実行します。
実行可能な各ブランチには、次の一般的なレイアウトの swarmforge/ ディレクトリが含まれています。
スウォームフォージ/
swarmforge.conf
憲法.プロンプト

憲法/
記事/
プロジェクト.プロンプト
ローカルエンジニアリング.プロンプト
ローカルワークフロー.プロンプト
...
役割/
<役割>.プロンプト
...
constitution.prompt がエントリポイントです。実行可能なブランチは通常、エージェントに swarmforge/constitution/articles/ 内のすべてのファイルを読み取るように指示するためにこれを使用します。
共有デフォルト記事はメインに次の場所に存在します。
swarmforge/憲法/記事/
エンジニアリングプロンプト
ハンドオフ.プロンプト
ワークフロー.プロンプト
SwarmForge は起動時に、ロール ワークツリーを作成する前に、不足している共有記事を実行可能ブランチの swarmforge/constitution/articles/ ディレクトリにインストールします。また、スクリプトの同期中に、欠落している共有アーティクルを各ロールのワークツリーにインストールします。既存のローカル ファイルはスキップされるため、実行可能なブランチは、同じファイル名のアーティクルをコミットすることで共有アーティクルをオーバーライドできます。
パック固有の追加と例外では、共有記事を編集するのではなく、明示的なローカル ファイル名を使用する必要があります。現在の規約は次のとおりです。
ワークフローのプロジェクト形状とローカル トポロジの project.prompt。
ワークフロー固有のエンジニアリング ルールの local-engineering.prompt。
ワークフロー固有のフロー ルールの local-workflow.prompt。
local-*.prompt 命名規則は、「この実行可能なブランチの共有デフォルト アーティクルを追加するか、特殊化する」ことを意味します。共有記事が有効なままで、ブランチに追加の要件、例外、またはより狭い指示のみが必要な場合に使用します。完全な置換には local-*.prompt を使用しないでください。ブランチが共有アーティクルを意図的にオーバーライドする場合は、代わりに共有ファイル名を使用します。
たとえば、 main は共有 workflow.prompt を提供できますが、 six-pack は QA 固有のハンドオフ動作用に local-workflow.prompt を追加できます。ブランチが共有ワークフロー記事を完全に置き換える必要がある場合は、独自の workflow.prompt をコミットできます。起動時はそのローカル ファイルをオーバーライドとして扱い、

共有のものをその上にコピーしないでください。
swarmforge/swarmforge.conf 内の各ロールは、対応する swarmforge/roles/<role>.prompt ファイルにマップされます。
SwarmForge は swarmforge/swarmforge.conf を読み取ります。
ルート ./swarm ラッパーは、共有ヘルパー スクリプト、ターミナル アダプター、および共有構成アーティクルがまだ存在しない場合、メイン ブランチからコピーします。
スタートアップは、不足している共有憲法記事を swarmforge/constitution/articles/ にインストールし、既存のローカル記事ファイルをスキップします。
起動時に、設定されたロール プロンプト、ヘルパー スクリプト、およびターミナル アダプタが検証されます。
ターゲット ディレクトリがまだ git リポジトリではない場合、起動時に git リポジトリが初期化され、最初のコミットが作成されます。
ロールが master または none に割り当てられていない限り、スタートアップは、構成されたロールごとに 1 つの git worktree を .worktrees/ の下に作成します。
起動時に swarmforge/scripts/ と欠落している共有構成記事が各ロール ワークツリーに同期され、そのローカル スクリプト ディレクトリが各エージェントの PATH に配置されるため、エージェントはマスター チェックアウトに戻らずにローカル ハンドオフ ヘルパーを使用します。
SwarmForge は tmux セッションを作成し、ターミナル ウィンドウを開き、割り当てられたワークツリーで構成された各バックエンドを起動します。
次の場合、スタートアップは OS 固有のスリープ阻害剤を開始します。

[切り捨てられた]

## Original Extract

A simple tool for coordinating several AI agents. Contribute to unclebob/swarm-forge development by creating an account on GitHub.

GitHub - unclebob/swarm-forge: A simple tool for coordinating several AI agents. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
unclebob
/
swarm-forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
258 Commits 258 Commits swarmforge swarmforge test/ swarmforge test/ swarmforge .gitignore .gitignore README.md README.md bb.edn bb.edn close-swarm close-swarm View all files Repository files navigation
Do not spend any money on a bankrbot SWARM token.
A disciplined tmux-based agent orchestration platform that turns swarms of AI agents into reliable, professional software engineers.
This main branch is documentary: it explains the system and carries the shared operational scripts and default constitution articles. The runnable workflow branches carry the project-facing configurations, role prompts, and local constitution articles that define specific workflows.
SwarmForge is an agent coordination system that facilitates communication between agents working in different git worktrees.
It provides a shared structure for role-specific prompts, worktree assignment, tmux sessions, and message passing so multiple agents can collaborate on the same project without stepping on each other.
The runnable SwarmForge configurations live on dedicated branches. Each branch contains the swarmforge/swarmforge.conf , local constitution articles, and role prompts for one workflow. At startup, its ./swarm wrapper copies the shared operational scripts and shared constitution articles from main when they are not already present, then launches that branch's local configuration.
two-pack is the quick backend workflow. Use it for small tasks that benefit from fast coding without the overhead of Gherkin and acceptance testing, while still preserving backend refactoring and hardening.
coder implements requested behavior with TDD and unit tests.
cleaner batches coder handoffs and performs cleanup, CRAP and DRY review, architectural review, encapsulation and separation-of-concerns fixes, and language mutation hardening.
The normal flow is coder -> cleaner -> coder . Use this branch when you want a tight implementation/refinement loop without specification, QA, property-test, or acceptance-test roles.
four-pack is the compact specification workflow. Use it for moderate projects that require Gherkin specification and some architectural consideration without splitting every quality gate into its own agent:
specifier turns user intent into precise Gherkin acceptance specifications and asks for approval before handoff.
coder implements approved behavior slices with TDD, unit tests, and generated acceptance tests.
refactorer performs behavior-preserving cleanup, coverage improvement, CRAP and DRY review, mutation-site scans, and property-test support.
architect owns high-level structure, dependency direction, mutation hardening, DRY review, soft Gherkin mutation, and final completion notification.
The normal flow is specifier -> coder -> refactorer -> architect -> specifier . Use this branch when you want disciplined development without splitting cleanup, architecture, hardening, and QA into separate agents.
six-pack is the full workflow. Use it for major projects that require full specification, up-front QA, backend verification, and significant architectural consideration. It separates each major quality gate into its own role:
specifier turns user intent into accepted Gherkin specifications and end-to-end QA procedures.
coder implements approved behavior slices with TDD, unit tests, and generated acceptance tests.
cleaner performs local behavior-preserving cleanup, coverage improvement, CRAP and DRY review, and mutation-site scans.
architect reviews module structure, boundaries, dependency direction, and property-test coverage.
hardender performs mutation hardening, language mutation, CRAP and DRY verification, and soft Gherkin mutation.
QA converts the specifier's QA procedures into executable scripts, runs final user-interface verification, checks handoff consistency, and sends completion notifications.
The normal flow is specifier -> coder -> cleaner -> architect -> hardender -> QA -> completion. Use this branch when you want each review and verification concern owned by a separate agent.
SwarmForge runs locally. Before starting a runnable branch, make sure the target machine has:
At least one configured agent backend, such as codex , claude , copilot , or grok
In the directory where you want to use SwarmForge, choose a runnable branch and pull its contents without creating a Git remote:
BRANCH=four-pack
curl -L " https://github.com/unclebob/swarm-forge/archive/refs/heads/ ${BRANCH} .tar.gz " | tar -xz --strip-components=1
Use BRANCH=two-pack for the quick two-agent workflow, BRANCH=four-pack for the compact specification workflow, or BRANCH=six-pack for the full six-agent workflow. Do not use main for this command; main is documentary and stores the shared operational scripts, while the runnable branches provide the configurations and prompts intended for projects.
After copying a runnable branch, start the swarm from the target project:
./swarm
The ./swarm wrapper keeps the runnable branch small. On first use, if swarmforge/scripts/ is missing, it downloads the main branch archive, copies the shared operational scripts from swarmforge/scripts/ , stages shared constitution articles from swarmforge/constitution/articles/ , and then launches swarmforge/scripts/swarmforge.sh . Later runs reuse the existing local scripts directory instead of overwriting it.
The windows should open automatically.
To stop the swarm, close the first window listed in swarmforge/swarmforge.conf . That cleanup window shuts down the tmux sessions and closes the remaining tracked windows.
While a swarm is active, SwarmForge tries to prevent the host from sleeping. On macOS it uses caffeinate ; on Linux it uses systemd-inhibit when available. Display lock or manual sleep can still interrupt agents depending on the OS. Set SWARMFORGE_PREVENT_SLEEP=0 before ./swarm to disable this behavior.
SwarmForge is a lightweight, tmux-based orchestration layer that:
Launches a config-driven swarm from a project-local swarmforge/swarmforge.conf
Creates one tmux session per configured role and opens a terminal surface for each role when the selected backend supports it
Reads behavior from project-local swarmforge/roles/<role>.prompt files plus a layered swarmforge/constitution.prompt
Supports per-role backends such as claude , codex , copilot , or grok
Puts the shared swarmforge/scripts/ directory on each agent's PATH , including handoff helpers for active swarm communication
Creates git worktrees under .worktrees/ for roles assigned to dedicated worktree names
Initializes a git repository in a new working directory when needed
Keeps all swarm state local to the working directory in .swarmforge/
Config-Driven Topology — The swarm shape comes from swarmforge/swarmforge.conf , not hardcoded shell variables.
Project-Local Roles — Each role is defined by swarmforge/roles/<role>.prompt in the working tree being orchestrated.
Layered Constitution — swarmforge/constitution.prompt directs agents to read article files under swarmforge/constitution/articles/ .
Backend Selection Per Role — A role can launch claude , codex , copilot , or grok .
Observable Swarm — Open one Terminal window per role and watch the sessions in real time.
Self-Hosted & Lightweight — Runs locally in tmux and Terminal with minimal machinery.
Each runnable branch contains a swarmforge/ directory with this general layout:
swarmforge/
swarmforge.conf
constitution.prompt
constitution/
articles/
project.prompt
local-engineering.prompt
local-workflow.prompt
...
roles/
<role>.prompt
...
constitution.prompt is the entry point. Runnable branches normally use it to tell agents to read every file in swarmforge/constitution/articles/ .
Shared default articles live on main under:
swarmforge/constitution/articles/
engineering.prompt
handoffs.prompt
workflow.prompt
At startup, SwarmForge installs missing shared articles into the runnable branch's swarmforge/constitution/articles/ directory before creating role worktrees. It also installs missing shared articles into each role worktree during script synchronization. Existing local files are skipped, so a runnable branch can override a shared article by committing an article with the same filename.
Pack-specific additions and exceptions should use explicit local filenames rather than editing shared articles. Current conventions are:
project.prompt for the workflow's project shape and local topology.
local-engineering.prompt for workflow-specific engineering rules.
local-workflow.prompt for workflow-specific flow rules.
The local-*.prompt naming convention means "add to or specialize the shared default article for this runnable branch." Use it when the shared article remains valid and the branch only needs extra requirements, exceptions, or narrower instructions. Do not use local-*.prompt for a full replacement; use the shared filename instead when the branch intentionally overrides the shared article.
For example, main can provide a shared workflow.prompt , while six-pack can add local-workflow.prompt for QA-specific handoff behavior. If a branch needs to replace the shared workflow article completely, it can commit its own workflow.prompt ; startup will treat that local file as an override and will not copy the shared one over it.
Each role in swarmforge/swarmforge.conf maps to a corresponding swarmforge/roles/<role>.prompt file.
SwarmForge reads swarmforge/swarmforge.conf .
The root ./swarm wrapper copies shared helper scripts, terminal adapters, and shared constitution articles from the main branch when they are not already present.
Startup installs missing shared constitution articles into swarmforge/constitution/articles/ , skipping any local article file that already exists.
Startup validates the configured role prompts, helper scripts, and terminal adapters.
If the target directory is not already a git repository, startup initializes one and creates the first commit.
Startup creates one git worktree per configured role under .worktrees/ , unless the role is assigned to master or none .
Startup syncs swarmforge/scripts/ and missing shared constitution articles into each role worktree and puts that local scripts directory on each agent's PATH , so agents use local handoff helpers without reaching back into the master checkout.
SwarmForge creates tmux sessions, opens terminal windows, and launches each configured backend in its assigned worktree.
Startup starts an OS-specific sleep inhibitor when o

[truncated]

---
source: "https://github.com/mensfeld/code-on-incus"
hn_url: "https://news.ycombinator.com/item?id=48904287"
title: "Mensfeld/code-on-incus: Give each AI agent its own isolated machine"
article_title: "GitHub - mensfeld/code-on-incus: Give each AI agent its own isolated machine with root, Docker, and systemd. Active defense detects and stops threats automatically. · GitHub"
author: "Tomte"
captured_at: "2026-07-14T09:47:09Z"
capture_tool: "hn-digest"
hn_id: 48904287
score: 1
comments: 0
posted_at: "2026-07-14T09:43:43Z"
tags:
  - hacker-news
  - translated
---

# Mensfeld/code-on-incus: Give each AI agent its own isolated machine

- HN: [48904287](https://news.ycombinator.com/item?id=48904287)
- Source: [github.com](https://github.com/mensfeld/code-on-incus)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T09:43:43Z

## Translation

タイトル: Mensfeld/code-on-incus: 各 AI エージェントに独自の隔離されたマシンを与える
記事のタイトル: GitHub - Mensfeld/code-on-incus: 各 AI エージェントに、root、Docker、systemd を備えた独自の隔離されたマシンを与えます。アクティブな防御機能により、脅威が自動的に検出され、阻止されます。 · GitHub
説明: 各 AI エージェントに、root、Docker、systemd を備えた独自の隔離されたマシンを与えます。アクティブな防御機能により、脅威が自動的に検出され、阻止されます。 - メンズフェルト/コード・オン・インカス

記事本文:
GitHub - Mensfeld/code-on-incus: 各 AI エージェントに、root、Docker、systemd を備えた独自の隔離されたマシンを与えます。アクティブな防御機能により、脅威が自動的に検出され、阻止されます。 · GitHub
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
メンズフェルト
/
コードオンインカス
公共
通知
署名する必要があります

n 通知設定を変更します
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,020 コミット 1,020 コミット .github/ workflows .github/ workflows cmd/ coi cmd/ coi docs docs 内部 内部 その他 その他 その他 プロファイル/ デフォルト プロファイル/ デフォルト スキーマ スキーマ スクリプト スクリプト testdata/ ダミー testdata/ ダミー テスト テスト .gitignore .gitignore .golangci.yml .golangci.yml .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh pyproject.toml pyproject.toml renovate.json renovate.json ruff.toml ruff.toml すべてのファイルを表示 リポジトリファイルナビゲーション
AI コーディング エージェント用に隔離されたマシン - アクティブな防御機能を備えています。
COI は、各 AI エージェントに独自のマシン、つまり root アクセス、systemd、Docker を備え、あらゆるものをインストールできる完全なシステム コンテナーを提供します。エージェントは、実際のシステムに触れることなく、サービスの実行、パッケージの管理、cron の使用など、実際のサーバー上で行うのと同じように動作します。ファイルは正しく所有されたままであり、権限をハックする必要はありません。
資格情報はホスト上に残ります。 SSH キー、環境変数、Git トークンは、明示的にマウントしない限り、AI ツールに公開されることはありません。何か問題が発生した場合、COI はそれを捕捉し (リバース シェル、資格情報のスキャン、データの抽出)、コンテナを自動的に一時停止または強制終了します。手動による介入は必要ありません。
AI エージェントを実行し、それらのエージェントが何をしているのか知りたい開発者向けに、開発者によって構築されました。製品でもスタートアップでもなく、仕事をするツールです。
AI コーディング エージェントを実行し、ホストを危険にさらすことなく、ルート、Docker、パッケージ マネージャー、サービスなどのフル マシン アクセスを許可したいと考えています。
エージェントが何か不審なことをしたとき、それを知るのではなく、知りたいのです。

事後
複数のエージェントを並行して実行し、それらを互いに分離する必要がある
毎回設定を失う使い捨てのコンテナではなく、再起動を繰り返しても存続する永続的な開発環境が必要です。
自分の資格情報がエージェント制御の環境内に流出しないように気を配る
Code on Incus に関する BetterStack ビデオをご覧ください。
Docker や Docker Sandbox ではなく Incus を使用する理由
サンドボックスでスクリプトとコマンドを実行する
コンテナのライフサイクルとセッションの永続性
Claude Code (デフォルト) - Anthropic の公式 CLI ツール
opencode - オープンソース AI コーディング エージェント ( https://opencode.ai )
pi - AI コーディング アシスタント ( https://pi.dev )
Aider - 端末での AI ペアプログラミング
ツールの選択は構成/プロファイルに基づいて行われます。
# ~/.coi/config.toml または ./.coi/config.toml
【ツール】
name = " opencode " # または "claude" (デフォルト)、"pi"
coi shell # 設定されたツールを使用します (デフォルトではクロード コード)
coi shell --profile opencode # または、[tool] name = "opencode" のプロファイル経由で切り替えます
許可モード - AI ツールが自律的に実行されるか、各アクションの前に質問されるかを制御します。
# ~/.coi/config.toml または .coi/config.toml
【ツール】
name = " claude " # デフォルトの AI ツール
Permission_mode = " bypass " # "bypass" (デフォルト) または "interactive"
詳細な構成、API キーのセットアップ、新しいツールの追加については、サポートされているツールの wiki ページを参照してください。
マルチスロットのサポート - 完全に分離された同じワークスペースに対して並列 AI コーディング セッションを実行します
セッション再開 - 完全な履歴と資格情報が復元された状態で会話を再開します (ワークスペース スコープ)
永続コンテナ - セッション間でコンテナを存続させます (インストールされたツールは保持されます)
ワークスペースの分離 - 各セッションはプロジェクト ディレクトリをマウントします
スロットの分離 - 各並列スロットには独自のホーム ディレクトリがあります (スロット間でファイルが漏洩しません)
ワークスペース ファイルは e 環境でも保持されます。

一時モード - コンテナのみが削除され、作業内容は常に保存されます。
コンテナのスナップショット - チェックポイントの作成、変更のロールバック、および完全な状態を保存したブランチ実験
SSH エージェント転送 - 秘密キーをコピーせずにコンテナ内で git-over-SSH を使用します ( [ssh] forward_agent = true )
ホスト ポートの公開 - ホスト上のコンテナ TCP ポートを公開します (アイデンティティ マップされたエージェントが使用可能なポートの [ports] プール、固定サービスの [[ports.map]]): スロットごとの決定的な割り当て、起動前の競合チェック、および信頼できないプロジェクト構成の Coi Trust Gating を使用して、エージェントが起動した開発サーバーが localhost:<port> で到達可能になります。
ホスト ソケット転送 - ホスト エンドポイントがコンテナに入らないように、任意のホスト Unix ソケットをコンテナ ( [[sockets]] ) に転送します。これは、資格情報ブローカーの構成要素です (ホスト上で有効期間の短いトークンを作成し、内部でオンデマンドでフェッチします)。信頼できない project-config ソケットは coi trust の背後でゲートされます
資格情報カタログ - [[credentials]] エントリ (構成またはプロファイル) を介してサードパーティプロバイダーの資格情報をコンテナーにコピーします。名前付きカタログバンドル ( Bundle = "ollama" ) を参照するか、まだカタログ化されていないもののアドホックホスト/コンテナー ファイルのペアを宣言します。 claude / opencode / pi 独自の認証情報ファイルは、同じ組み込みカタログから取得されます。信頼できないプロジェクト .coi/config.toml からのアドホック エントリは、 coi trust の背後でゲートされます。カタログ参照には、組み込みツールの資格情報がすでに持っているものと同じ信頼レベルが保持されます。
環境変数転送 - ホスト環境変数を名前で選択的に転送します (構成内の forward_env)
コマンドソースの環境変数 - 有効期間の短い API キー/トークンの場合、開始時にホスト コマンドを実行し、その出力を環境変数 ( [defaults.env_commands] ) として挿入することで、セッションごとに新しいシークレットを作成します。信頼できるスコープの構成

のみ
ホストのタイムゾーンの継承 - コンテナはホストのタイムゾーンを自動的に継承します ([タイムゾーン] 構成で構成可能)
サンドボックス コンテキスト ファイル - 自動挿入される ~/SANDBOX_CONTEXT.md は、AI ツールに環境 (ネットワーク モード、ワークスペース パス、永続性など) を伝えます。各ツールのネイティブ コンテキスト システムに自動的にロードされます: ~/.claude/CLAUDE.md 経由の Claude Code、opencode.json の指示フィールド経由の OpenCode、~/.pi/agent/APPEND_SYSTEM.md シンボリックリンク経由の pi (auto_context = false でオプトアウト)
資格情報の保護 - SSH キー、.env ファイル、Git 資格情報、および環境変数は、明示的にマウントされない限り公開されません。
特権コンテナー ガード - security.privileged=true が検出されると開始を拒否し、すべてのコンテナー分離を無効にします。
セキュリティ体制の検証 - coi ヘルス チェックで seccomp、AppArmor、特権設定をチェックし、完全な分離を確認します。
カーネル バージョンの強制 - 安全な分離のためのセキュリティ機能が欠けている可能性がある 5.15 未満のホスト カーネルについて警告します。
リアルタイムの脅威検出 - カーネルレベルの nftables 監視により、リバース シェル、C2 接続、データ漏洩、DNS トンネリング、資格情報スキャンを検出します。
自動応答 - 脅威が高い場合は自動一時停止、重大な場合は自動停止 - 手動介入は不要
ネットワーク分離 - nftables ベースの制限/許可リスト/オープン モードによりプライベート ネットワーク アクセスをブロックし、漏洩を防止します
保護されたパス - .git/hooks 、 .git/config 、 .husky 、 .vscode はサプライチェーン攻撃を防ぐために読み取り専用でマウントされています
ホスト側の不変保護 - 保護されたパスはセッション中に chattr +i でロックされ、読み取り専用マウントの unshare -m + umount バイパスを防ぎます (オプトアウト: [security] host_immutable = false )
Git ID ガード - コンテナーは user.useConfigOnly=true を強制し、AI ツールが d としてコミットするのを防ぎます。

デフォルトの「コード」ユーザー
ゲスト API が無効になっています - Incus ゲスト API ( /dev/incus ) はデフォルトで無効になっており、ホスト パスとトポロジのリークを防止します
システムコンテナ - 非特権コンテナによる完全な OS 分離。Docker 特権モードよりも優れています。
自動 UID マッピング - 権限がない地獄、ファイルは正しく所有されています
監査ログ - フォレンジックとコンプライアンスのために、すべてのセキュリティ イベントが JSONL に記録されます。
AI コーディング ツールでは、多くの場合、広範なファイル システムへのアクセスや権限チェックのバイパスが必要になります。
「ルート」はホスト システムではなくコンテナのルートであるため、これらの操作はコンテナ内で安全です。
コンテナは一時的なものです - すべての変更は含まれており、ホストには影響しません。
これにより、システムを保護しながら AI ツールにすべての機能が提供されます
# インストール
カール -fsSL https://raw.githubusercontent.com/mensfeld/code-on-incus/master/install.sh |バッシュ
# イメージのビルド (初回のみ、約 5 ～ 10 分)
コイビルド
# 好みの AI ツールでコーディングを開始します (デフォルトは Claude Code)
プロジェクトの cd
コイの殻
# または、代わりに opencode を使用します (config-driven: [tool] name = "opencode",
# またはプロファイル: coi shell --profile opencode)
＃それです！ AI コーディング アシスタントは、以下を備えた分離されたコンテナーで実行されています。
# - プロジェクトは /workspace にマウントされます
# - ファイル権限を修正します (chown はもう必要ありません!)
# - コンテナ内の完全な Docker アクセス
# - PR/問題管理に GitHub CLI を利用可能
# - すべてのワークスペースの変更は自動的に保持されます
# - ホストの SSH キー、環境変数、または認証情報にアクセスできません
Docker や Docker Sandbox ではなく Incus を使用する理由
Incus は、LXD から派生した最新の Linux コンテナーおよび仮想マシン マネージャーです。 Docker (アプリケーション コンテナを使用する) とは異なり、Incus は完全な init システムを備えた軽量 VM のように動作するシステム コンテナを提供します。
能力
コードオンインカス
ドッカーサンドボックス
ベアメタル
認証情報 ISO

関係
デフォルト (決して公開されない)
部分的
なし
リアルタイムの脅威検出
カーネルレベル (nftables)
いいえ
いいえ
リバースシェル検出
自動終了
いいえ
いいえ
データ漏洩のアラート
自動一時停止
いいえ
いいえ
ネットワークの分離
nftables (3 つのモード)
基本
いいえ
保護されたパス
読み取り専用マウント
いいえ
いいえ
自動応答 (一時停止/停止)
はい
いいえ
いいえ
監査ログ
JSONLフォレンジック
いいえ
いいえ
サプライチェーン攻撃の防止
Git フック/IDE 構成は保護されています
いいえ
いいえ
Docker サンドボックスではなく Incus を使用する理由
Linux 後ではなく、Linux ファースト。 Docker Sandboxes の microVM 分離は、macOS と Windows でのみ利用できます。 Linux は従来のコンテナベースのフォールバックを取得します。 Incus は Linux ネイティブであるため、COI は最初から Linux 用に構築されています。
Docker デスクトップは必要ありません。 Docker Sandboxes は Docker Desktop の機能です。 Docker Desktop はオープンソースではないため、大規模な組織には商用ライセンスが必要です。 COI は Incus のみに依存しています。完全にオープン ソースであり、ベンダー ロックインや追加のランタイムはありません。
VM 内のコンテナーではなく、システム コンテナー。 Incus システム コンテナは、内部で systemd とネイティブ Docker サポートを備えた完全な OS を実行します (1 つのクリーンな分離レイヤー)。 Docker Sandbox は microVM 内にアプリケーション コンテナを入れ子にするため、アーキテクチャがさらに複雑になります。
許可無し地獄。インカスオートマ

[切り捨てられた]

## Original Extract

Give each AI agent its own isolated machine with root, Docker, and systemd. Active defense detects and stops threats automatically. - mensfeld/code-on-incus

GitHub - mensfeld/code-on-incus: Give each AI agent its own isolated machine with root, Docker, and systemd. Active defense detects and stops threats automatically. · GitHub
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
mensfeld
/
code-on-incus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1,020 Commits 1,020 Commits .github/ workflows .github/ workflows cmd/ coi cmd/ coi docs docs internal internal misc misc profiles/ default profiles/ default schema schema scripts scripts testdata/ dummy testdata/ dummy tests tests .gitignore .gitignore .golangci.yml .golangci.yml .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.sh install.sh pyproject.toml pyproject.toml renovate.json renovate.json ruff.toml ruff.toml View all files Repository files navigation
Isolated machines for AI coding agents - with active defense.
COI gives each AI agent its own machine - a full system container with root access, systemd, Docker, and the ability to install anything. Agents work like they would on a real server: run services, manage packages, use cron - without touching your actual system. Files stay correctly owned, no permission hacks needed.
Your credentials stay on the host. SSH keys, environment variables, and Git tokens are never exposed to AI tools unless you explicitly mount them. If something goes wrong, COI catches it - reverse shells, credential scanning, data exfiltration - and pauses or kills the container automatically. No manual intervention needed.
Built by developers, for developers who run AI agents and want to know what those agents are doing. Not a product, not a startup - a tool that does the job.
You run AI coding agents and want them to have full machine access - root, Docker, package managers, services - without risking your host
You want to know when an agent does something suspicious, not find out after the fact
You run multiple agents in parallel and need them isolated from each other
You want persistent dev environments that survive restarts and reboots, not throwaway containers that lose your setup every time
You care about your credentials not ending up inside an agent-controlled environment
Watch the BetterStack video about Code on Incus
Why Incus Instead of Docker or Docker Sandboxes?
Run Scripts and Commands in the Sandbox
Container Lifecycle & Session Persistence
Claude Code (default) - Anthropic's official CLI tool
opencode - Open-source AI coding agent ( https://opencode.ai )
pi - AI coding assistant ( https://pi.dev )
Aider - AI pair programming in your terminal
Tool selection is config/profile-driven:
# ~/.coi/config.toml or ./.coi/config.toml
[ tool ]
name = " opencode " # or "claude" (default), "pi"
coi shell # Uses the configured tool (Claude Code by default)
coi shell --profile opencode # Or switch via a profile with [tool] name = "opencode"
Permission mode - Control whether AI tools run autonomously or ask before each action:
# ~/.coi/config.toml or .coi/config.toml
[ tool ]
name = " claude " # Default AI tool
permission_mode = " bypass " # "bypass" (default) or "interactive"
See the Supported Tools wiki page for detailed configuration, API key setup, and adding new tools.
Multi-slot support - Run parallel AI coding sessions for the same workspace with full isolation
Session resume - Resume conversations with full history and credentials restored (workspace-scoped)
Persistent containers - Keep containers alive between sessions (installed tools preserved)
Workspace isolation - Each session mounts your project directory
Slot isolation - Each parallel slot has its own home directory (files don't leak between slots)
Workspace files persist even in ephemeral mode - Only the container is deleted, your work is always saved
Container snapshots - Create checkpoints, rollback changes, and branch experiments with full state preservation
SSH agent forwarding - Use git-over-SSH inside containers without copying private keys ( [ssh] forward_agent = true )
Host port publishing - Publish container TCP ports on the host ( [ports] pool for identity-mapped agent-usable ports, [[ports.map]] for fixed services): agent-started dev servers become reachable at localhost:<port> , with per-slot deterministic allocation, a pre-launch conflict check, and coi trust gating for untrusted project configs
Host socket forwarding - Forward arbitrary host Unix sockets into the container ( [[sockets]] ) so the host endpoint never enters the container — the building block for credential brokers (mint short-lived tokens on the host, fetch them on demand inside). Untrusted project-config sockets are gated behind coi trust
Credential catalog - Copy third-party provider credentials into the container via [[credentials]] entries (config or profile): reference a named catalog bundle ( bundle = "ollama" ) or declare an ad-hoc host/container file pair for anything not yet cataloged. claude / opencode / pi 's own credential files come from the same built-in catalog. Ad-hoc entries from an untrusted project .coi/config.toml are gated behind coi trust ; catalog references carry the same trust level the built-in tool credentials already have
Environment variable forwarding - Selectively forward host env vars by name ( forward_env in config)
Command-sourced env vars - Mint a fresh secret per session by running a host command at start and injecting its output as an env var ( [defaults.env_commands] ) — for short-lived API keys/tokens. Trusted-scope config only
Host timezone inheritance - Containers automatically inherit the host's timezone (configurable via [timezone] config)
Sandbox context file - Auto-injected ~/SANDBOX_CONTEXT.md tells AI tools about their environment (network mode, workspace path, persistence, etc.). Automatically loaded into each tool's native context system: Claude Code via ~/.claude/CLAUDE.md , OpenCode via the instructions field in opencode.json , pi via ~/.pi/agent/APPEND_SYSTEM.md symlink (opt out with auto_context = false )
Credential protection - SSH keys, .env files, Git credentials, and environment variables are never exposed unless explicitly mounted
Privileged container guard - Refuses to start when security.privileged=true is detected, which defeats all container isolation
Security posture verification - coi health checks seccomp, AppArmor, and privilege settings to confirm full isolation
Kernel version enforcement - Warns on host kernels below 5.15 that may lack security features for safe isolation
Real-time threat detection - Kernel-level nftables monitoring detects reverse shells, C2 connections, data exfiltration, DNS tunneling, and credential scanning
Automated response - Auto-pause on HIGH threats, auto-kill on CRITICAL — no manual intervention needed
Network isolation - nftables-based restricted/allowlist/open modes block private network access and prevent exfiltration
Protected paths - .git/hooks , .git/config , .husky , .vscode mounted read-only to prevent supply-chain attacks
Host-side immutable protection - Protected paths are locked with chattr +i during sessions, preventing unshare -m + umount bypass of read-only mounts (opt out: [security] host_immutable = false )
Git identity guard - Containers enforce user.useConfigOnly=true , preventing AI tools from committing as the default "code" user
Guest API disabled - Incus guest API ( /dev/incus ) disabled by default, preventing host path and topology leaks
System containers - Full OS isolation with unprivileged containers, better than Docker privileged mode
Automatic UID mapping - No permission hell, files owned correctly
Audit logging - All security events logged to JSONL for forensics and compliance
AI coding tools often need broad filesystem access or bypass permission checks
These operations are safe inside containers because the "root" is the container root, not your host system
Containers are ephemeral - any changes are contained and don't affect your host
This gives AI tools full capabilities while keeping your system protected
# Install
curl -fsSL https://raw.githubusercontent.com/mensfeld/code-on-incus/master/install.sh | bash
# Build image (first time only, ~5-10 minutes)
coi build
# Start coding with your preferred AI tool (defaults to Claude Code)
cd your-project
coi shell
# Or use opencode instead (config-driven: [tool] name = "opencode",
# or a profile: coi shell --profile opencode)
# That's it! Your AI coding assistant is now running in an isolated container with:
# - Your project mounted at /workspace
# - Correct file permissions (no more chown!)
# - Full Docker access inside the container
# - GitHub CLI available for PR/issue management
# - All workspace changes persisted automatically
# - No access to your host SSH keys, env vars, or credentials
Why Incus Instead of Docker or Docker Sandboxes?
Incus is a modern Linux container and virtual machine manager, forked from LXD. Unlike Docker (which uses application containers), Incus provides system containers that behave like lightweight VMs with full init systems.
Capability
code-on-incus
Docker Sandbox
Bare Metal
Credential isolation
Default (never exposed)
Partial
None
Real-time threat detection
Kernel-level (nftables)
No
No
Reverse shell detection
Auto-kill
No
No
Data exfiltration alerts
Auto-pause
No
No
Network isolation
nftables (3 modes)
Basic
No
Protected paths
Read-only mounts
No
No
Auto response (pause/kill)
Yes
No
No
Audit logging
JSONL forensics
No
No
Supply-chain attack prevention
Git hooks/IDE configs protected
No
No
Why Incus Instead of Docker Sandboxes?
Linux-first, not Linux-last. Docker Sandboxes' microVM isolation is only available on macOS and Windows. Linux gets a legacy container-based fallback. COI is built for Linux from the ground up because Incus is Linux-native.
No Docker Desktop required. Docker Sandboxes is a Docker Desktop feature. Docker Desktop is not open source and has commercial licensing requirements for larger organizations. COI depends only on Incus - fully open source, no vendor lock-in, no additional runtime.
System containers, not containers-in-VMs. Incus system containers run a full OS with systemd and native Docker support inside - one clean isolation layer. Docker Sandboxes nests application containers inside microVMs, adding architectural complexity.
No permission hell. Incus automa

[truncated]

---
source: "https://github.com/srelens/srelens"
hn_url: "https://news.ycombinator.com/item?id=49195728"
title: "Show HN: Srelens – Kubernetes control room for engineers and AI agents (MIT)"
article_title: "GitHub - srelens/srelens: The Kubernetes control room—built in Rust, ready for engineers and AI agents. · GitHub"
author: "deveshk0"
captured_at: "2026-08-06T12:51:26Z"
capture_tool: "hn-digest"
hn_id: 49195728
score: 2
comments: 2
posted_at: "2026-08-06T12:26:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Srelens – Kubernetes control room for engineers and AI agents (MIT)

- HN: [49195728](https://news.ycombinator.com/item?id=49195728)
- Source: [github.com](https://github.com/srelens/srelens)
- Score: 2
- Comments: 2
- Posted: 2026-08-06T12:26:30Z

## Translation

タイトル: Show HN: Srelens – エンジニアと AI エージェントのための Kubernetes コントロール ルーム (MIT)
記事のタイトル: GitHub - srelens/srelens: Kubernetes コントロール ルーム - Rust で構築され、エンジニアと AI エージェントに対応します。 · GitHub
説明: Kubernetes コントロール ルーム - Rust に組み込まれており、エンジニアと AI エージェントが使用できます。 - スレレンズ/スレレンズ

記事本文:
GitHub - srelens/srelens: Kubernetes コントロール ルーム - Rust に組み込まれており、エンジニアと AI エージェントが利用できます。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
スレレンズ
/
スレレンズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
dev ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
211 コミット 211 コミット .github .github apps/ デスクトップ アプリ/ デスクトップ クレート クレート docs docs パッケージング/ aur パッケージング/ aur .doc

kerignore .dockerignore .gitignore .gitignore .mcp.json .mcp.json CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile ライセンス ライセンス README.md README.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Kubernetes コントロール ルーム - Rust に組み込まれており、エンジニアと AI エージェントが利用できます。
srelens は、SRE 向けのオープンソースのローカルファースト Kubernetes デスクトップ ワークスペースです。
プラットフォーム エンジニア、DevOps エンジニア。調査、分析、安全確保
Tauri v2、React 19 で構築された 1 つのアプリケーションからクラスター全体でアクションを実行
そして純粋なRustコア。
ウェブサイト ·
インストール・
クイックスタート ·
ユーザーガイド・
MCPサーバー・
開発者ガイド ·
貢献する
Kubernetes のトラブルシューティングでは、多くの場合、ターミナル、ダッシュボード、YAML 間の移動が必要になります。
エディター、ログ、クラスター コンテキスト。 srelens はその調査ループを
ローカルファーストのデスクトップワークスペースが 1 つあります。
調査からアクションまで 1 つのワークスペース - リソースの参照、イベントの検査
および YAML、ログの追跡、ターミナルの使用、ポート転送の管理、クラスターの取得
ツールを頻繁に切り替えることなくアクションを実行できます。
エンジニアと AI エージェント向けに構築 — バックエンド機能もサポートされています
内蔵の MCP サーバーを通じて利用できます。
ローカルファーストクラスターアクセス — srelens はローカルからの資格情報を使用します
kubeconfig を使用し、ルーティングなしで Kubernetes API サーバーに直接接続します
srelens クラウド サービスを介したクラスター アクセス。
安全な操作 — 破壊的な行為は識別され、確認ゲートが設けられます。
オープン ソース — MIT のもとでライセンスが付与されており、公開コード、リリース、問題、および
GitHub 上のロードマップ。
srelens は、Tauri v2 を通じてオペレーティング システム WebView を使用し、

Rustバックエンドが構築されました
kube-rs と tokio を使用します。独自に開発したものであり、とは関係ありません
Mirantis Lens または Freelens プロジェクト。
それぞれの使用方法については、ユーザー ガイドを参照してください。
マルチクラスター ワークスペース — kubeconfig コンテキストを検出し、さらに追加または貼り付けます
ファイルを作成し、各コンテキストに名前、ロゴ、色を付けて、クラスタを
クラスターホットバー。ファイル間で名前を共有するコンテキスト (例:default) は次のとおりです。
曖昧さが解消されるため、すべてのクラスターが表示され、アクセス可能な状態になります。
ライブ Kubernetes リソース — ワークロード、ネットワーキング、ストレージ、RBAC、
入場、自動スケーリング、ライブウォッチ更新を含むカスタムリソース、検索、
列ピッカー、名前空間のスコープ設定、および一括アクション。
リソースの詳細と YAML — マニフェスト、イベント、関係、および
メトリクスを編集し、検証、ドライランの差分を使用してスキーマ対応の YAML を編集します。
サーバー側で適用します。
ログ — 以前のインスタンス (クラッシュ後) を含むポッドまたはワークロードのログをストリームします。
ログ、タイムスタンプ、末尾および以降のウィンドウ コントロール、ソースごとの色分け、
コンテナのフィルタリング、およびバッファまたはすべてのコンテナのエクスポート。
ターミナルとシェル — ポッド実行セッション、コンテキストスコープのローカルを開く
ターミナル、distroless ポッド用の一時的なデバッグ コンテナー、および特権ノード
貝殻。
ポート転送 — オープンするたびに転送を作成、検査、コピー、停止します
クラスター。
Helm — リリースをリストして検査し、インストール、アップグレード、ロールバック、または
値エディターとレンダリング差分プレビューを使用してそれらをアンインストールします。
ツールボックス — kubectl 、 krew 、 helm 、および krew プラグインをインストールおよび管理します。
コンテキストの実行認証ツールの要件を診断します。
メトリック — メトリックサーバーが使用可能な場合のノードとポッドの CPU とメモリ。
運用アクション - ワークロードのスケーリング、ロールアウトの再開、ポッドの削除または削除、
確認ゲートを使用して、CronJob を一時停止またはトリガーし、ノードを遮断またはドレインします。

破壊的な行為の場合。
コマンド パレット - ビュー間でキーボードから最初にナビゲーション (Cmd/Ctrl-K)、
コンテキストとリソース。
アプリケーション ログ — srelens 独自のローテーション ログ ファイルを [設定] から読み取ります。
問題が発生した後に診断します。
MCP アクセス — サポートされているバックエンド機能を MCP 対応クライアントに公開します
stdio またはループバック HTTP。
お使いのプラットフォームの最新ベータ版を次からダウンロードします。
GitHub リリース 。
プラットフォーム固有のインストールについては、インストール ガイドを参照してください。
初回起動、更新、検証、アンインストールの手順。
Web アプリとして実行 (Docker): srelens はマルチユーザー Web サーバーとしても実行できます
コンテナの中。ユーザーは OIDC (または試用版のローカル開発者ログイン) でサインインし、
それぞれが、自分がアップロードしたもののみから構築された完全に分離された環境を取得します
kubeconfig。 OIDC で保護されたクラスターは、ブラウザーベースのヘッドランプ スタイルで動作します。
サインイン — srelens は認証コード + PKCE フローを実行し、
id_token 自体なので、kubelogin /exec プラグインは必要ありません。 Kubeconfig とトークン
必要な SRELENS_MASTER_KEY の下に保存された状態でシールされ、書き込まれることはありません
ディスク。復号化されたファイルは tmpfs にのみ存在します。一部のデスクトップ専用アクション (ホスト シェル、
raw helm リポジトリ/プラグイン) は共有サーフェスからゲートされます - Web ユーザーは
代わりに、RBAC スコープのポッド内実行端末。詳細については docs/WEB.md を参照してください。
デプロイメント、環境変数、および完全なセキュリティ モデル。
srelens には、srelens で使用されるのと同じ機能レジストリから生成された MCP サーバーが含まれています。
デスクトップのバックエンド。したがって、サポートされているバックエンド機能は、
別個のクラスター統合レイヤーを作成せずに MCP 対応クライアントを使用できます。
ベアラー トークンで保護されたループバック HTTP 経由で MCP サーバーを実行します。
公開、回転、または取り消し - 回転すると実行中のサーバーが再起動され、新しい
トークンはすぐに有効になります (処理中のリクエストは削除されます)

est と無効化
古い値を使用した設定);取り消すとサーバーも停止します。
トークンを必要としない標準入出力接続用の srelens CLI をインストールします。
クライアントはプロセスを生成することですでに権限を保持しています。
サポートされている MCP クライアントのクライアント設定をコピーします。
サーバーを直接起動することもできます。
srelens --mcp-stdio
srelens --mcp-http 127.0.0.1:8765
破壊的なツールでは、アプリで確認を求められます。ヘッドレス実行には何もありません
ダイアログを表示するため、通話時に "_confirm": true が必要です。
プロセスレベルのオプトイン — --mcp-allow-destructive で何かを変更するか、
--mcp-allow-sensitive-reads はシークレットを読み取ります。両者は独立しているので、
Secret の読み取りは、ノードをドレインする許可を意味することはなく、フラグも意味しません。
_confirm なしで単独で何かを承認します。標準入出力の GUI 切り替えはありません。
詳細についてはユーザーガイドを参照してください
ホストヘッダーチェック、監査ログ、トークンストレージなどのセキュリティモデル。
{
"mcpサーバー": {
"スレレンズ" : {
"コマンド" : " srelens " ,
"args" : [ " --mcp-stdio " ]
}
}
}
MCP アクセスでは、ローカルで認証されたクラスター コンテキストが使用されます。ツール呼び出しを確認する
特に重要な場合には、適切な Kubernetes RBAC 権限を使用します。
クラスター。
クラスター依存のワークフローのための到達可能な Kubernetes クラスター
git clone https://github.com/srelens/srelens
CD スレレンズ
pnpmインストール
pnpm開発
便利なコマンド
コマンド
目的
pnpm開発
デスクトップアプリケーションを開発モードで起動します。
pnpmテスト
JavaScript および TypeScript テストを実行する
貨物試験
Rustワークスペーステストを実行する
pnpmビルド
実稼働フロントエンドを構築する
pnpm タウリビルド
パッケージ化されたデスクトップバイナリを作成する
アーキテクチャ、テスト標準、
機能を追加するための手順。
React 19 + TypeScript
│
│ Tauri のコマンドとイベント
▼
Tauri v2 デスクトップ シェル
│
▼
Pure-Rust バックエンド
§

── 能力レジストリ
§── Kubernetes と kube-rs の統合
§── ライブウォッチ、ログ、実行、およびポート転送
§── ヘルムとメトリクス
└── 標準入出力およびループバック HTTP 経由の MCP サーバー
リポジトリのレイアウト:
アプリ/デスクトップ/
src/ React および TypeScript デスクトップ UI
src-tauri/ Tauri アプリケーションと Rust コマンドブリッジ
木箱/
機能/バックエンド機能レジストリ
kube/Kubernetes クライアント、ウォッチ、アクション、Helm、およびメトリクス
mcp/MCPサーバー
docs/ インストール、使用法、開発、およびプロジェクトのドキュメント
プロジェクトのステータス
srelens は現在ベータ版です。評価と日常のテストの準備ができています。
ただし、ユーザーはリリース ノートを確認し、これを使用する場合は特に注意する必要があります。
クリティカルクラスター。
安定版リリースの前に重大な変更が発生する可能性があります。フィードバック、バグレポート、
再現可能なトラブルシューティングの詳細も歓迎します。
Reddit の r/srelens — お知らせ、
質問、フィードバック
問題 — バグと機能リクエスト
貢献は大歓迎です。以下から始めます:
事前に行動規範を確認してください
参加しています。
srelens は MIT ライセンスに基づくオープンソースです。
Mirantis Lens または Freelens プロジェクトとは無関係です。
Kubernetes コントロール ルーム - Rust に組み込まれており、エンジニアと AI エージェントが利用できます。
Readme MIT ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
4 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The Kubernetes control room—built in Rust, ready for engineers and AI agents. - srelens/srelens

GitHub - srelens/srelens: The Kubernetes control room—built in Rust, ready for engineers and AI agents. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
srelens
/
srelens
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
dev Branches Tags Go to file Code Open more actions menu Folders and files
211 Commits 211 Commits .github .github apps/ desktop apps/ desktop crates crates docs docs packaging/ aur packaging/ aur .dockerignore .dockerignore .gitignore .gitignore .mcp.json .mcp.json CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
The Kubernetes control room—built in Rust, ready for engineers and AI agents.
srelens is an open-source, local-first Kubernetes desktop workspace for SREs,
platform engineers, and DevOps engineers. Investigate, analyse, and take safe
action across clusters from one application built with Tauri v2, React 19,
and a pure-Rust core.
Website ·
Install ·
Quick start ·
User guide ·
MCP server ·
Developer guide ·
Contributing
Kubernetes troubleshooting often means moving between terminals, dashboards, YAML
editors, logs, and cluster contexts. srelens brings that investigation loop into
one local-first desktop workspace.
One workspace from investigation to action — browse resources, inspect events
and YAML, follow logs, use terminals, manage port forwards, and take cluster
actions without constantly switching tools.
Built for engineers and AI agents — supported backend capabilities are also
available through the built-in MCP server.
Local-first cluster access — srelens uses credentials from your local
kubeconfig and connects directly to Kubernetes API servers, without routing
cluster access through a srelens cloud service.
Safe operations — destructive actions are identified and confirmation-gated.
Open source — licensed under MIT, with public code, releases, issues, and
roadmap on GitHub.
srelens uses the operating system WebView through Tauri v2 and a Rust backend built
with kube-rs and tokio . It is independently developed and is not affiliated with
Mirantis Lens or the Freelens project.
See the user guide for how to use each of these.
Multi-cluster workspace — discover kubeconfig contexts, add or paste more
files, give each context a name, logo, and colour, and switch clusters from the
cluster hotbar. Contexts that share a name across files (e.g. default ) are
disambiguated so every cluster stays visible and reachable.
Live Kubernetes resources — browse workloads, networking, storage, RBAC,
admission, autoscaling, and custom resources with live watch updates, search,
column pickers, namespace scoping, and bulk actions.
Resource details and YAML — inspect manifests, events, relationships, and
metrics, and edit schema-aware YAML with validation, dry-run diffs, and
server-side apply.
Logs — stream pod or workload logs with previous-instance (post-crash)
logs, timestamps, tail and since-window controls, per-source colouring,
container filtering, and buffer or all-container export.
Terminals and shells — open pod exec sessions, a context-scoped local
terminal, ephemeral debug containers for distroless pods, and privileged node
shells.
Port forwarding — create, inspect, copy, and stop forwards across every open
cluster.
Helm — list and inspect releases, and install, upgrade, roll back, or
uninstall them with a values editor and rendered-diff preview.
Toolbox — install and manage kubectl , krew , helm , and krew plugins, and
diagnose a context's exec-auth tool requirements.
Metrics — node and pod CPU and memory when metrics-server is available.
Operational actions — scale workloads, restart rollouts, evict or delete pods,
suspend or trigger CronJobs, and cordon or drain nodes, with confirmation gates
for destructive actions.
Command palette — keyboard-first navigation (Cmd/Ctrl-K) across views,
contexts, and resources.
Application logs — read srelens's own rotating log file from Settings to
diagnose issues after they happen.
MCP access — expose supported backend capabilities to MCP-capable clients over
stdio or loopback HTTP.
Download the latest beta for your platform from
GitHub Releases .
See the installation guide for platform-specific installation,
first-launch, updating, verification, and uninstall instructions.
Run as a web app (Docker): srelens can also run as a multi-user web server
in a container. Users sign in with OIDC (or a local dev login for trials) and
each gets a fully isolated environment built only from their own uploaded
kubeconfigs. OIDC-protected clusters work with a browser-based, Headlamp-style
sign-in — srelens runs the authorization-code + PKCE flow and injects the
id_token itself, so no kubelogin /exec plugin is needed. Kubeconfigs and tokens
are sealed at rest under a required SRELENS_MASTER_KEY that is never written to
disk; decrypted files live only in tmpfs. Some desktop-only actions (host shell,
raw helm repo/plugin) are gated off the shared surface — web users get
RBAC-scoped in-pod exec terminals instead. See docs/WEB.md for
deployment, environment variables, and the full security model.
srelens includes an MCP server generated from the same capability registry used by
the desktop backend. Supported backend capabilities can therefore be used by
MCP-capable clients without creating a separate cluster integration layer.
run the MCP server over loopback HTTP, protected by a bearer token you can
reveal, rotate, or revoke — rotating restarts the running server so the new
token takes effect at once (dropping any in-flight request and invalidating
configs that used the old value); revoking also stops the server;
install the srelens CLI for stdio connections, which need no token — the
client already holds your privileges by spawning the process;
copy client configuration for supported MCP clients.
You can also start the server directly:
srelens --mcp-stdio
srelens --mcp-http 127.0.0.1:8765
Destructive tools prompt for confirmation in the app. Headless runs have no
dialog to show, so they need "_confirm": true on the call and a
process-level opt-in — --mcp-allow-destructive to change anything, or
--mcp-allow-sensitive-reads to read Secrets. The two are independent, so
reading a Secret never implies permission to drain a node, and neither flag
alone authorizes anything without _confirm . There's no GUI toggle for stdio.
See the user guide for the full
security model, including the Host-header check, audit log, and token storage.
{
"mcpServers" : {
"srelens" : {
"command" : " srelens " ,
"args" : [ " --mcp-stdio " ]
}
}
}
MCP access uses your locally authenticated cluster contexts. Review tool calls
and use appropriate Kubernetes RBAC permissions, especially with critical
clusters.
A reachable Kubernetes cluster for cluster-dependent workflows
git clone https://github.com/srelens/srelens
cd srelens
pnpm install
pnpm dev
Useful commands
Command
Purpose
pnpm dev
Launch the desktop application in development mode
pnpm test
Run JavaScript and TypeScript tests
cargo test
Run Rust workspace tests
pnpm build
Build the production frontend
pnpm tauri build
Create packaged desktop binaries
See the developer guide for architecture, testing standards,
and instructions for adding capabilities.
React 19 + TypeScript
│
│ Tauri commands and events
▼
Tauri v2 desktop shell
│
▼
Pure-Rust backend
├── capability registry
├── Kubernetes integration with kube-rs
├── live watches, logs, exec, and port forwarding
├── Helm and metrics
└── MCP server over stdio and loopback HTTP
Repository layout:
apps/desktop/
src/ React and TypeScript desktop UI
src-tauri/ Tauri application and Rust command bridge
crates/
capability/ Backend capability registry
kube/ Kubernetes clients, watches, actions, Helm, and metrics
mcp/ MCP server
docs/ Installation, usage, development, and project documentation
Project status
srelens is currently in beta . It is ready for evaluation and everyday testing,
but users should review release notes and take extra care when using it with
critical clusters.
Breaking changes may still occur before a stable release. Feedback, bug reports,
and reproducible troubleshooting details are welcome.
r/srelens on Reddit — announcements,
questions, and feedback
Issues — bugs and feature requests
Contributions are welcome. Start with:
Please review the Code of Conduct before
participating.
srelens is open source under the MIT License .
Not affiliated with Mirantis Lens or the Freelens project.
The Kubernetes control room—built in Rust, ready for engineers and AI agents.
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

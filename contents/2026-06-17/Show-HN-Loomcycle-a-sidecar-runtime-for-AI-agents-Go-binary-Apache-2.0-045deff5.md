---
source: "https://github.com/denn-gubsky/loomcycle"
hn_url: "https://news.ycombinator.com/item?id=48565329"
title: "Show HN: Loomcycle – a sidecar runtime for AI agents (Go binary, Apache-2.0)"
article_title: "GitHub - denn-gubsky/loomcycle: The runtime substrate for agentic systems — one Go binary, six LLM providers, MCP-native, configurable as a managed sandbox or full agentic dev environment. Where agents live, talk, and learn. · GitHub"
author: "denn-gubsky"
captured_at: "2026-06-17T04:35:29Z"
capture_tool: "hn-digest"
hn_id: 48565329
score: 2
comments: 0
posted_at: "2026-06-17T03:26:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Loomcycle – a sidecar runtime for AI agents (Go binary, Apache-2.0)

- HN: [48565329](https://news.ycombinator.com/item?id=48565329)
- Source: [github.com](https://github.com/denn-gubsky/loomcycle)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T03:26:02Z

## Translation

タイトル: Show HN: Loomcycle – AI エージェント用のサイドカー ランタイム (Go バイナリ、Apache-2.0)
記事のタイトル: GitHub - denn-gubsky/loomcycle: エージェント システム用のランタイム サブストレート — 1 つの Go バイナリ、6 つの LLM プロバイダー、MCP ネイティブ、マネージド サンドボックスまたは完全なエージェント開発環境として構成可能。エージェントが生活し、話し、学ぶ場所。 · GitHub
説明: エージェント システム用のランタイム サブストレート — 1 つの Go バイナリ、6 つの LLM プロバイダー、MCP ネイティブ、マネージド サンドボックスまたは完全なエージェント開発環境として構成可能。エージェントが生活し、話し、学ぶ場所。 - デン・グブスキー/ルームサイクル

記事本文:
GitHub - denn-gubsky/loomcycle: エージェント システムのランタイム サブストレート — 1 つの Go バイナリ、6 つの LLM プロバイダー、MCP ネイティブ、マネージド サンドボックスまたは完全なエージェント開発環境として構成可能。エージェントが生活し、話し、学ぶ場所。 · GitHub
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
デン

-グブスキー
/
織機
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
618 コミット 618 コミット .github .github アダプター アダプター ベンチ ベンチ cmd/ loomcycle cmd/ loomcycle ドキュメント ドキュメントのサンプル サンプル 内部 内部プロト プロト テスト テスト Web Web .dockerignore .dockerignore .env.insecure.example .env.insecure.example .env.local.example .env.local.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml BACKERS.md BACKERS.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.release Dockerfile.release LICENSE LICENSE Makefile Makefile README.md README.md REVISIONS.md REVISIONS.md docker-compose.cluster.yaml docker-compose.cluster.yaml docker-compose.example.yaml docker-compose.example.yaml go.mod go.mod go.sum go.sum loomcycle-mcp.sh loomcycle-mcp.sh loomcycle.example.yaml loomcycle.example.yaml loomcycle.local-interactive.example.yaml loomcycle.local-interactive.example.yaml loomcycle.sh loomcycle.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
サイドカー内のエージェント ランタイム。
アプリケーションと一緒に 1 つの Go バイナリ。強化されたエージェント ループ、両側の MCP、マルチレプリカ HA。 Apache-2.0。
🌐 loomcycle.dev ·
📝 エンジニアリングブログ ·
📐 建築
🌳 v1.0 が登場しました。 loomcycle 1.0 がリリースされました。機能セットは完成し、ランタイムは強化され、配布の準備が整いました。 v0.8 → v0.23 ライン全体で安定したコア プリミティブ (マルチレプリカ HA、サブストレート Defs Agent/Skill/MCPServer/Schedule/Webhook/MemoryBackend/A2A、A2A 相互運用性、インバウンド Webhook、プラガブル メモリ + メモリ層、合成コード JS プロバイダー、および v0.17.0 の OSS マルチテナント認証 - プリンシプごと)

アルベアラートークン + ロール認識 Web UI)。 v0.24 → v0.37 ラインは、対話型ターミナル、一時停止/再開/スナップショット + インスタンス間再開、コンテキスト圧縮サブシステム、コンテキスト変換プラグイン、すべてのトランスポートでの外部ファンアウト、低速ローカル モデルの堅牢性など、強化と操作性の向上を目的とした実行でした。 1.0 自体は純粋な強化と配布のマイルストーンであり、新しいプリミティブはなく、Homebrew、マルチ アーチ Docker、および Claude Code プラグインに接続されています。 8 時間の安定性ソーク: 127 万回線、380 万エージェント実行、468 ウェーブ全体で 100% 完了、リークゼロ。 Apache-2.0。私たちはバグレポート、セキュリティ開示、機能の貢献、ダウンストリームのコンシューマー、フォークを歓迎します。 COTRIBUTING.md を参照してください。
サイドカー内のエージェント ランタイム。 loomcycle は 1 つの Go バイナリで、約 50 MB です。アプリケーション内ではなく、アプリケーションと並行して実行されます。アプリは、HTTP、gRPC、MCP、TypeScript アダプター、または Python アダプター経由で loomcycle を呼び出します。エージェント ループ、マルチプロバイダー ルーティング、メモリとチャネルのプリミティブ、MCP サーバー ID、OpenTelemetry トレース、およびマルチレプリカ調整はすべてバイナリ内に存在します。アプリケーションは、作成した言語に関係なくそのまま残ります。
変わった形ですね。今日のエージェント システム市場には 3 つの選択肢があります。 1 つ: Python または TypeScript ライブラリをアプリケーション プロセス内に埋め込みます。 2: 1 つのベンダーの IAM に関連付けられたマネージド クラウド サービスをレンタルします。 3: 実際にエージェントを実行しないゲートウェイを介してモデル呼び出しをプロキシします。
loomcycle は 4 番目のオプションです。ループを所有し、スタックがすでに使用しているすべてのワイヤ形式を話す軽量の自己ホスト可能なランタイム。
リリース
ハイライト
v0.4 → v0.26.x 基盤
ランタイムの基盤となるすべてが凝縮されています。 7 つの推論モード: Anthropic、OpenAI、DeepSeek、Gemini、Ollama (クラウド + ローカル)、さらに合成 code-js プロバイダーと MO

ckプロバイダー。強化されたモデル→tool_use→tool_resultのループ。クロード コードと同等の 19 個の組み込みツール (読み取り、書き込み、編集、Grep、Glob、NotebookEdit) に加え、HTTP、WebFetch、WebSearch、Bash、エージェント、スキル、メモリ、チャネル、AgentDef、SkillDef、評価、中断、コンテキスト。コンテンツアドレス指定のランタイム変更可能な基板 (エージェント / スキル / MCPServer / スケジュール / Webhook / MemoryBackend / A2A defs)。プラグ可能な MemoryBackend 上の Vector Memory (sqlite-vec / pgvector)、その上にメモリ層があります。両側にMCP。 LLM ゲートウェイ + OpenAI 互換のシム。 A2A 相互運用。 Webhook を入力します。アンサンブル同期プリミティブ (RFC S)。 OTEL + テナントごとの公平性 + 一時停止/再開/スナップショット + マルチレプリカ HA 。実行ごとの名前付き資格情報 + ツール使用フック。状態プレーンと定義プレーンの両方にわたる OSS マルチテナント認可 (RFC L、v0.17.0)。埋め込み型 React Web UI + インタラクティブ ターミナル。 TS + Python + n8n アダプター。 Homebrew + Docker ディストリビューション。バージョンごとの詳細: REVISIONS.md 。
v0.27.0
インタラクティブな実行はターミナルを離れても残ります。 context.WithoutCancel でバックグラウンドで goroutine を実行し、GET /v1/runs/{run_id}/stream (replay-from- ?from_seq + live-tail) 経由で再アタッチします。コンテキスト op=self は、解決されたプロバイダー + モデルを報告します。
v0.28.0
エージェントごとの LLM サンプリング (温度/top_p/top_k/ペナルティ/シード/ストップ、yaml または AgentDef オーバーレイまたは実行ごとに設定)。そして、協調静止を一時停止します。ループは反復境界で停止し、Pause() は実行中の実行を待機するため、実行中のスナップショットは信頼できます。
v0.29.0
Web UI＋操作性。エージェント エディターのサンプリング コントロール + 折りたたみ可能な高度な JSON/YAML オーバーレイ、ターミナル メッセージ エコー + コンテキスト サイズ ゲージ、および廃止されたエージェント名のソフト再利用 (新しいランタイム プリミティブなし)。
v0.30.0
スナップのクロスインスタンス再開
[切り捨てられた]
同じ

同じ構成スキーマでバイナリに移行します。オペレーターはいくつかの環境変数を反転して姿勢を選択します。
信頼境界は、operator/caller です。オペレータ設定はフロアです。呼び出し側はリクエストごとに絞り込むことはできますが、広げることはできません。ベアラー トークン ( LOOMCYCLE_AUTH_TOKEN ) が権限です。トークンを持つすべての人を、ランタイムを駆動するために完全に信頼できるものとして扱います。サンドボックス環境で真の分離を実現するには、コンテナまたは VM 内で loomcycle を実行します。 Bash は制限されています (cwd、env スクラブ、出力制限、タイムアウト) が、カーネル レベルのサンドボックスではありません。
適切なパスを選択してください。 4 つすべてが同じ単一の静的バイナリを出荷します
プラス、v0.11.1 init / Doctor の初回実行フロー。 Context.help のインストールでは、それぞれについて詳しく説明します。
# Homebrew (macOS + Linux)
醸造インストール denn-gubsky/loomcycle/loomcycle
# Docker (v0.11.2+; pull は Apple Silicon を含む amd64 + arm64 で動作します)
docker pull denngubsky/loomcycle:latest
# ソースからインストールします (Web UI の埋め込みをスキップします — 開発のみ)
github.com/denn-gubsky/loomcycle/cmd/loomcycle@latest をインストールしてください
# 直接 tarball (darwin-arm64 / darwin-amd64 / linux-arm64 / linux-amd64 のいずれか)
カール -L https://github.com/denn-gubsky/loomcycle/releases/latest/download/loomcycle-darwin-arm64.tar.gz |タールxz
クイックスタート (秒、認証済み)
loomcycle init --with-token # config を書き込み、~/.config/loomcycle/auth.env にトークンを生成します (0600)
import ANTHROPIC_API_KEY=sk-... # (または OPENAI_API_KEY / DEEPSEEK_API_KEY) — 少なくとも 1 つのプロバイダー キー
loomcycle Doctor #env + キー + ストレージ + 作成されたばかりのトークンを検証します
loomcycle # は 127.0.0.1:8787 で開始します (auth.env を自動ロード — Shell-rc 編集なし)
init --with-token は、Web UI URL ( http://127.0.0.1:8787/ui ) を出力します。それを開いて、ログイン プロンプトで ~/.config/loomcycle/auth.env のトークンを貼り付けます。 (トークンは 0600 ファイルに保存され、UR に埋め込まれることはありません)

L. ?token= リンクは、ベアラーをブラウザ履歴とフロント プロキシのログに漏洩します。) loomcycle と loomcycle Doctor は両方とも、config ディレクトリから auth.env を自動ロードします。実際のエクスポート LOOMCYCLE_AUTH_TOKEN=… は常にこれをオーバーライドします。
適切な層を選択してください。それぞれは上記のスーパーセットです。認証は何かが設定された場合にのみ適用されるため、Tier 1 にはトークンがまったく必要ありません。
層 1: ゼロ構成開発 (オープン モード、ローカルホスト)
トークンもフラグもありません。 127.0.0.1 でタイヤを蹴る最速の方法。
loomcycle init # config のみ — シークレットは書き込まれません
エクスポート ANTHROPIC_API_KEY=sk-...
loomcycle # オープン モード: /v1/* + /ui は認証されずにパススルーします (警告をログに記録します)
http://127.0.0.1:8787/ui を開く
LOOMCYCLE_AUTH_TOKEN も作成されたトークンもない場合、ランタイムはローカルホスト上でオープンして実行されます。すべてのリクエストが許可され、whoami は合成管理者を返します。 10秒間のスモークテストに適しています。これをローカルホストから公開しないでください。
層 2: 単一の共有トークン (推奨されるデフォルト)
一人の担い手がすべてをゲートします。 init --with-token は簡単なボタンです (上記)。同等の手動セットアップ:
ルームサイクルの初期化
import LOOMCYCLE_AUTH_TOKEN= $( openssl rand -hex 32 ) # または: loomcycle init --with-token
エクスポート ANTHROPIC_API_KEY=sk-...
織機
open " http://127.0.0.1:8787/ui?token= $LOOMCYCLE_AUTH_TOKEN " # Cookie を 1 回設定します
トークンを保持しているすべての人を、ランタイムを駆動するために完全に信頼できるものとして扱います。
層 3: マルチテナント、プリンシパルごとのトークン (RFC L、v0.17.0)
開発者/アプリごとに個別のベアラーを作成し、それぞれが権限のある (テナント、サブジェクト、スコープ) にバインドされます。 Tier-2 デプロイメントをダウンタイムなしで適切に移行します。
# 既存の共有トークンをサブストレートにプロモートし、スコープ付きトークンを作成します
loomcycle オペレータートークン作成 --copy-from-env --name ops --tenant ops --scopes 基板:管理者
loomcycle オペレータートークン create --name acme-app --t

enant acme --subject alice --scopes 実行:作成
最初の管理者 OperatorTokenDef は、従来の共有トークン フォールバックを無効にします。ルートごとの HTTP および RPC ごとの gRPC スコープ。 Web UI はロールを認識します (スーパー管理者とテナント)。 Context.help 演算子トークンと REVISIONS.md の v0.17.0 のメモを参照してください。
カール http://127.0.0.1:8787/healthz
# {"ok":true}
実際の通話 (別の端末から):
カール -N http://127.0.0.1:8787/v1/runs \
-H " 認証: ベアラー $LOOMCYCLE_AUTH_TOKEN " \
-H " Content-Type: application/json " \
-d ' {"エージェント":"デフォルト","セグメント":[{"ロール":"ユーザー","コンテンツ":[{"タイプ":"信頼できるテキスト","テキスト":"Hello"}]}]} '
チェックアウトからビルドする (開発用):
ビルドオール # UI + バイナリをワンショットで作成します。出力 → ./bin/loomcycle
./bin/loomcycle --config loomcycle.example.yaml
マルチレプリカクラスターのデモ (v0.12.x)。検証スクリプトを使用して 1 つのコマンドでクラスター (2 つの loomcycle レプリカ、Postgres、nginx LB) を構成する Docker については、examples/cluster/README.md を参照してください。完全なオペレーター Runbook は docs/MULTI-REPLICA.md にあります。
v1.0.0: 機能が完全で、強化され、配布の準備ができています。 🌳 1.0 マイルストーン — マルチテナント認証のキャップストーン (v0.17.0) およびサブストレートの完全性ライン (v0.18 ～ v0.23) 以降、ロードマップが指しているポイント。 1.0 では新しいプリミティブは追加されません。安定のタグです

[切り捨てられた]

## Original Extract

The runtime substrate for agentic systems — one Go binary, six LLM providers, MCP-native, configurable as a managed sandbox or full agentic dev environment. Where agents live, talk, and learn. - denn-gubsky/loomcycle

GitHub - denn-gubsky/loomcycle: The runtime substrate for agentic systems — one Go binary, six LLM providers, MCP-native, configurable as a managed sandbox or full agentic dev environment. Where agents live, talk, and learn. · GitHub
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
denn-gubsky
/
loomcycle
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
618 Commits 618 Commits .github .github adapters adapters bench bench cmd/ loomcycle cmd/ loomcycle docs docs examples examples internal internal proto proto test test web web .dockerignore .dockerignore .env.insecure.example .env.insecure.example .env.local.example .env.local.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml BACKERS.md BACKERS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.release Dockerfile.release LICENSE LICENSE Makefile Makefile README.md README.md REVISIONS.md REVISIONS.md docker-compose.cluster.yaml docker-compose.cluster.yaml docker-compose.example.yaml docker-compose.example.yaml go.mod go.mod go.sum go.sum loomcycle-mcp.sh loomcycle-mcp.sh loomcycle.example.yaml loomcycle.example.yaml loomcycle.local-interactive.example.yaml loomcycle.local-interactive.example.yaml loomcycle.sh loomcycle.sh View all files Repository files navigation
The agentic runtime, in a sidecar.
One Go binary alongside your application. Hardened agent loop, MCP on both sides, multi-replica HA. Apache-2.0.
🌐 loomcycle.dev ·
📝 Engineering blog ·
📐 Architecture
🌳 v1.0 is here. loomcycle 1.0 is released — the feature set is complete and the runtime is hardened and distribution-ready. The core primitives stabilised across the v0.8 → v0.23 line (multi-replica HA; the substrate Defs Agent/Skill/MCPServer/Schedule/Webhook/MemoryBackend/A2A; A2A interoperability; inbound webhooks; pluggable memory + a memory layer; the synthetic code-js provider; and OSS multi-tenant authorization in v0.17.0 — per-principal bearer tokens + a role-aware Web UI). The v0.24 → v0.37 line was a hardening + operability run: the interactive terminal, pause/resume/snapshot + cross-instance resume, the context-compaction subsystem, context-transform plugins, external fan-out on every transport, and slow-local-model robustness. 1.0 itself is a pure hardening + distribution milestone — no new primitives — wired to Homebrew, multi-arch Docker, and the Claude Code plugin. 8-hour stability soak: 1.27M circuits, 3.8M agent runs, 100% completion across 468 waves, zero leaks. Apache-2.0. We welcome bug reports, security disclosures, feature contributions, downstream consumers, and forks. See CONTRIBUTING.md .
The agentic runtime, in a sidecar. loomcycle is one Go binary, ~50 MB. It runs alongside your application, not inside it. Your app calls loomcycle over HTTP, gRPC, MCP, the TypeScript adapter, or the Python adapter. The agent loop, multi-provider routing, memory and channel primitives, MCP server identity, OpenTelemetry traces, and multi-replica coordination all live in the binary. Your application stays in whatever language you wrote it in.
The shape that's different. Today's agentic-systems market gives you three options. One: embed a Python or TypeScript library inside your application process. Two: rent a managed cloud service tied to one vendor's IAM. Three: proxy your model calls through a gateway that doesn't actually run agents.
loomcycle is a fourth option. A lightweight self-hostable runtime that owns the loop and speaks every wire format your stack already uses.
Release
Highlights
v0.4 → v0.26.x foundation
Everything the runtime is built on, condensed. Seven inference modes : Anthropic, OpenAI, DeepSeek, Gemini, Ollama (cloud + local), plus the synthetic code-js provider and a mock provider. The hardened model → tool_use → tool_result loop. 19 built-in tools with Claude Code parity (Read, Write, Edit, Grep, Glob, NotebookEdit), plus HTTP, WebFetch, WebSearch, Bash, Agent, Skill, Memory, Channel, AgentDef, SkillDef, Evaluation, Interruption, Context. The content-addressed, runtime-mutable substrate (Agent / Skill / MCPServer / Schedule / Webhook / MemoryBackend / A2A defs). Vector Memory (sqlite-vec / pgvector) on a pluggable MemoryBackend , with a memory layer above. MCP on both sides. The LLM Gateway + OpenAI-compatible shims. A2A interop. Input webhooks . Ensemble-sync primitives (RFC S). OTEL + per-tenant fairness + Pause / Resume / Snapshot + multi-replica HA . Per-run named credentials + tool-use hooks. OSS multi-tenant authorization (RFC L, v0.17.0) across both the state and definition planes. The embedded React Web UI + the interactive terminal . TS + Python + n8n adapters. Homebrew + Docker distribution. Per-version detail: REVISIONS.md .
v0.27.0
Interactive runs survive leaving the terminal . Background-goroutine execution under context.WithoutCancel , re-attach via GET /v1/runs/{run_id}/stream (replay-from- ?from_seq + live-tail). Context op=self reports the resolved provider + model.
v0.28.0
Per-agent LLM sampling (temperature / top_p / top_k / penalties / seed / stop, set via yaml or AgentDef overlay or per-run). And pause cooperative quiesce : the loop parks at an iteration boundary and Pause() waits for in-flight runs, so a mid-run snapshot is reliable.
v0.29.0
Web UI + operability. Agent-editor sampling controls + a collapsible advanced JSON/YAML overlay, terminal message-echo + a context-size gauge, and soft reclaim of a retired agent name (no new runtime primitives).
v0.30.0
Cross-instance resume of a snaps
[truncated]
Same Go binary, same config schema. Operator flips a few env vars to pick the posture.
The trust boundary is operator / caller . The operator config is the floor; callers can narrow per-request but never widen. The bearer token ( LOOMCYCLE_AUTH_TOKEN ) is the authority. Treat anyone with the token as fully trusted to drive the runtime. For true isolation in the sandbox posture, run loomcycle inside a container or VM. Bash is restricted (cwd, env scrub, output bounds, timeouts) but it is not a kernel-level sandbox.
Pick the path that fits. All four ship the same single static binary
plus the v0.11.1 init / doctor first-run flow. Context.help installation covers each in detail.
# Homebrew (macOS + Linux)
brew install denn-gubsky/loomcycle/loomcycle
# Docker (v0.11.2+; pull works on amd64 + arm64 including Apple Silicon)
docker pull denngubsky/loomcycle:latest
# go install from source (skips Web UI embedding — for dev only)
go install github.com/denn-gubsky/loomcycle/cmd/loomcycle@latest
# Direct tarball (one of darwin-arm64 / darwin-amd64 / linux-arm64 / linux-amd64)
curl -L https://github.com/denn-gubsky/loomcycle/releases/latest/download/loomcycle-darwin-arm64.tar.gz | tar xz
Quick start (seconds, authenticated)
loomcycle init --with-token # writes config + mints a token to ~/.config/loomcycle/auth.env (0600)
export ANTHROPIC_API_KEY=sk-... # (or OPENAI_API_KEY / DEEPSEEK_API_KEY) — at least one provider key
loomcycle doctor # verify env + keys + storage + the just-minted token
loomcycle # starts on 127.0.0.1:8787 (auto-loads auth.env — no shell-rc edit)
init --with-token prints the Web UI URL ( http://127.0.0.1:8787/ui ). Open it, then paste the token from ~/.config/loomcycle/auth.env at the login prompt. (The token is kept in the 0600 file and never embedded in a URL. A ?token= link would leak the bearer into browser history and into any fronting proxy's logs.) loomcycle and loomcycle doctor both auto-load auth.env from the config dir; a real export LOOMCYCLE_AUTH_TOKEN=… always overrides it.
Pick the tier that fits. Each is a superset of the one above. Auth is enforced only once something is configured , so Tier 1 needs no token at all.
Tier 1: zero-config dev (open mode, localhost)
No token, no flags. Fastest way to kick the tires on 127.0.0.1 .
loomcycle init # config only — no secret written
export ANTHROPIC_API_KEY=sk-...
loomcycle # open mode: /v1/* + /ui pass through unauthenticated (logs a warning)
open http://127.0.0.1:8787/ui
With no LOOMCYCLE_AUTH_TOKEN and no minted tokens, the runtime runs open on localhost. Every request is allowed, whoami returns a synthetic admin. Good for a 10-second smoke test. Never expose this off localhost.
Tier 2: single shared token (the recommended default)
One bearer gates everything. init --with-token is the easy button (above). Equivalent manual setup:
loomcycle init
export LOOMCYCLE_AUTH_TOKEN= $( openssl rand -hex 32 ) # or: loomcycle init --with-token
export ANTHROPIC_API_KEY=sk-...
loomcycle
open " http://127.0.0.1:8787/ui?token= $LOOMCYCLE_AUTH_TOKEN " # sets the cookie once
Treat anyone holding the token as fully trusted to drive the runtime.
Tier 3: multi-tenant, per-principal tokens (RFC L, v0.17.0)
Mint a distinct bearer per developer / app, each bound to an authoritative (tenant, subject, scopes) . Migrate a Tier-2 deployment in place, no downtime:
# promote your existing shared token into the substrate, then mint scoped tokens
loomcycle operator-token create --copy-from-env --name ops --tenant ops --scopes substrate:admin
loomcycle operator-token create --name acme-app --tenant acme --subject alice --scopes runs:create
The first admin OperatorTokenDef disables the legacy shared-token fallback. Per-route HTTP and per-RPC gRPC scopes. The Web UI becomes role-aware (super-admin vs tenant). See Context.help operator-tokens and the v0.17.0 notes in REVISIONS.md .
curl http://127.0.0.1:8787/healthz
# {"ok":true}
Real call (from another terminal):
curl -N http://127.0.0.1:8787/v1/runs \
-H " Authorization: Bearer $LOOMCYCLE_AUTH_TOKEN " \
-H " Content-Type: application/json " \
-d ' {"agent":"default","segments":[{"role":"user","content":[{"type":"trusted-text","text":"Hello"}]}]} '
Build from a checkout (for development):
make build-all # UI + binary in one shot; output → ./bin/loomcycle
./bin/loomcycle --config loomcycle.example.yaml
Multi-replica cluster demo (v0.12.x). For a one-command docker compose up cluster (2 loomcycle replicas, Postgres, nginx LB) with a verify script, see examples/cluster/README.md . Full operator runbook in docs/MULTI-REPLICA.md .
v1.0.0: feature-complete, hardened, distribution-ready. 🌳 The 1.0 milestone — the point the roadmap pointed at since the multi-tenant-auth capstone (v0.17.0) and the substrate-completeness line (v0.18–v0.23). 1.0 adds no new primitives ; it's the stable tag

[truncated]

---
source: "https://github.com/kyle-visner/magpie"
hn_url: "https://news.ycombinator.com/item?id=49091292"
title: "Show HN: Magpie, a CLI to Turn Your AI Agent into a Bookkeeper"
article_title: "GitHub - kyle-visner/magpie: Accounting CLI built on Jaybase · GitHub"
author: "kvisner"
captured_at: "2026-07-28T23:51:54Z"
capture_tool: "hn-digest"
hn_id: 49091292
score: 1
comments: 0
posted_at: "2026-07-28T23:16:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Magpie, a CLI to Turn Your AI Agent into a Bookkeeper

- HN: [49091292](https://news.ycombinator.com/item?id=49091292)
- Source: [github.com](https://github.com/kyle-visner/magpie)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T23:16:16Z

## Translation

タイトル: Show HN: Magpie、AI エージェントを簿記係に変える CLI
記事のタイトル: GitHub - kyle-visner/magpie: Jaybase · GitHub 上に構築された会計 CLI
説明: Jaybase 上に構築されたアカウンティング CLI。 GitHub でアカウントを作成して、kyle-visner/magpie の開発に貢献してください。
HN テキスト: こんにちは、HN。Magpie のリリースを発表します。Magpie は、Jaybase をベースにした、AI エージェントを専門の簿記係に変える CLI ツールです。 ( https://news.ycombinator.com/item?id=48999936 ) 小規模なテクノロジー企業を経営する者として、私は非常にシンプルな本を持っています。月に 10 ～ 30 件のトランザクション、いくつかの減価償却項目、および監視する予算がいくつかあります。ですから、これらの単純なことを整理するためだけに、ソフトウェアに年間何百ドルも払っていて、それに加えておそらくライブの監視員にもお金を払っていることが常に気になっていました。私にはそれを行うための精神的余裕がなかったからです。そこで私はマグパイを作りました。 Magpie は、Jaybase 上で実行されるシンプルで独自の CLI で、一般的な AI エージェントが適切な複式簿記を実行できるようにします。標準的な会計ルールを適用し、すべてを安全に保存し、必要なほぼすべての種類のレポートやレポートを AI に要求できます。最初は自分用に作成しましたが、より幅広いコミュニティでも使用できるものだと考えました。試してみてください。 Claude Code、Codex、またはその他のエージェントに URL を伝えるだけで、セットアップできるはずです。

記事本文:
GitHub - kyle-visner/magpie: Jaybase 上に構築された会計 CLI · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
カイル・ヴィスナー
/
カササギ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
47 コミット 47 コミット .github/ workflows .github/ workflows cmd/ magpie cmd/ magpie docs docs external/ magpie external/ magpie scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum llm.md llm.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
Magpie は、人間と AI エージェント向けの独自の会計 CLI です。それは提供します
勘定科目表管理、複式仕訳帳のための 1 つの保護されたパス
顧客、請求書、支払い、メモ、スナップショット、監査読み取り。 CLI チェック
暗号化された不変イベントを追加する前の RBAC およびアカウンティング不変条件
ジェイベース。
Go 1.26.5 以降が必要です。以前の Go リリースには既知の標準ライブラリが含まれています
脆弱性があるため、リリース バイナリのビルドには使用しないでください。
git clone https://github.com/kyle-visner/magpie.git
CD カササギ
./cmd/magpie をインストールします
カササギ --store .magpie init
初期化されたローカルブックは現金会計を使用し、所有者アクターを作成します
所有者の役割を使用します。追加する前に初回セットアップをお読みください
別の俳優や金融活動の投稿。
Magpie は、エージェントに何もせずに帳簿作成を手伝ってもらいたい小規模チーム向けです。
彼らに生のデータベースへのアクセスや、会計動作を発明する許可を与えます。それ
ソースドキュメントのインポート、管理された移行、請求書と支払いに適合します
ワークフロー、操作メモ、および書き込みが関与する必要があるその他のジョブ、
監査可能で、安全に再試行でき、ポリシーに違反すると拒否されます。
Magpie は CLI およびドメイン エンジンであり、ホスト型会計 SaaS や代替品ではありません。
専門的な会計判断のために。開発にローカル バックエンドを使用するか、
信頼できるシングルユーザープロセス。共有または運用環境で使用する場合は、Magpie を実行します。
個別にデプロイされた Jaybase サービスと、認証された呼び出し元をバインドする

許可される
カササギのアクター ID。
Magpie は 1.0 より前です。実装されたアカウンティング サーフェスは使用可能であり、テストされています。
このリリースには、以下にリストされている機能が含まれています。以下の特徴は、
現在の範囲外:
ネイティブの QuickBooks CSV、IIF、または QBXML 解析。エージェントはソースを正規化する必要がある
Magpie の JSON コントラクトにデータを追加します。
請求書、銀行照合または調整、期末締め、税金、ローン、振込、
固定資産、保持、ガベージ コレクション、またはポイントインタイム復元コマンド。
ノート検索、バックリンク、入力されたエンティティ間の参照、差分、またはグラフ
ナビゲーション。
人間指向の出力モード、対話型 UI、署名付きコマンド エンベロープ、または
--actor の身元を証明する認証層。
これらは製品範囲の境界であり、隠れたインストール手順ではありません。参照
docs/SECURITY.md は、追加の運用管理については、
引き続き配備者の責任となります。
リポジトリのルートはアカウンティング CLI プロジェクトです。
cmd/magpie には CLI が含まれ、internal/magpie にはアカウンティングが含まれます
ドメイン。
Jaybase は別個に保守されています。
github.com/kyle-visner/jaybase 。
Magpie はそれを Go モジュールの依存関係として固定します。
AGPL-3.0以降。 「ライセンス」を参照してください。
ローカルに埋め込まれホストされている Jaybase バックエンドを使用した cmd/magpie の正規 CLI。
Jaybase を介した追加専用の SHA-256 対応イベント履歴。
ページ分割された再生、オプティミスティック同時実行、冪等書き込み、およびリモートの名前付き参照を備えた、ベアラー認証されたホスト型 Jaybase アクセス (HTTPS 経由)。
保存されたノード ペイロードの AES-256-GCM 暗号化。
台帳、メモ、スナップショット、監査読み取り用の統合 RBAC。
永続化前の複式元帳の検証。
現金、修正現金、発生主義会計に対する帳簿レベルの会計基準のサポート。
ワークフローに安全なアカウント選択のためのアカウントの役割をグラフ化します。
必要な監査理由を備えた特権的な手動仕訳調整。
一流の顧客、

根拠を意識した仕訳帳を生成する音声および支払いワークフロー。
勘定科目に関する構造化された外部ソース参照。
マークダウン ノートの作成、更新、リスト、および取得の操作。
QuickBooks または他のシステムからのエージェント マップ エクスポート用のソースタグ付き仕訳エントリ。
回復可能なルートの名前付きスナップショット。
エージェント使用のためのデフォルトの JSON コマンド出力。
ビジネス不変条件、CLI 動作、BDD スタイルのコア シナリオの自動テスト。
Go 1.26.5 以降を使用してください。リポジトリのルートから:
モッド検証に行く
テストに行く - レース ./...
獣医に行ってください。/...
go build -o ./magpie ./cmd/magpie
golang.org/x/vuln/cmd/govulncheck@v1.6.0 をインストールしてください
govulncheck ./...
環境がデフォルトの Go ビルド キャッシュをブロックしている場合は、書き込み可能なキャッシュを使用します。
GOCACHE=/private/tmp/magpie-gocache go mod verify
GOCACHE=/private/tmp/magpie-gocache go test -race ./...
GOCACHE=/private/tmp/magpie-gocache 獣医に行く ./...
GOCACHE=/private/tmp/magpie-gocache go build -o ./magpie ./cmd/magpie
生成された ./magpie バイナリは Git によって無視されます。
タグ付きリリースは、クリーンなメインコミットからビルドされます。
release.yml を使用して再現可能
scripts/build-release.sh プロセス。各GitHub
このリリースには、amd64 および arm64 用の macOS および Linux アーカイブに加えて、
SHA256SUMS ファイル。
このリポジトリから会計ベンチマークを実行します。別の Jaybase リポジトリから Jaybase ベンチマークを実行します。
GOCACHE=/private/tmp/magpie-gocache go test -run ' ^$ ' -bench 。 -benchmem ./...
ベースライン比較の場合、複数の実行をキャプチャし、 Benchstat と比較します。
mkdir -p .ベンチマーク
GOCACHE=/private/tmp/magpie-gocache go test -run ' ^$ ' -bench 。 -benchmem -count 5 ./... > .benchmarks/magpie-main.txt
benchstat .benchmarks/magpie-main.txt .benchmarks/magpie-feature.txt
エージェント統合パターン
エージェントは、llm.md を運用契約として読む必要があります。短い
パターン

以下の n は、必要な呼び出しの形式を示しています。
エージェントに固定コマンド テンプレートを与え、stdout を JSON として解析するように指示します。
./カササギ \
--store .magpie \
--actor AGENT_USER_ID \
コマンド...
ホストされた Jaybase サービスの場合、オリジンとベアラーの資格情報を次の方法で提供します。
環境。トークンをコマンドラインのフラグ、URL、ペイロード、ログ、
または冪等性キー:
import JAYBASE_URL=https://jaybase.example.com
import JAYBASE_TOKEN= ' シークレットマネージャーからのライタートークン '
./カササギ \
--actor AGENT_USER_ID \
コマンド...
--jaybase-url はオリジンをオーバーライドできますが、トークンは次の方法でのみ受け入れられます。
JAYBASE_TOKEN 。ホストされたリクエストは、再生時にのみ復号化されたペイロードをフェッチします
状態。監査出力はメタデータのみのままです。書き込みには Jaybase の Expected_root を使用します
と冪等性キー コントラクトを上書きする代わりに競合を返します。
新しいルート。
Magpie は、Martin と 1 つの直線的な Jaybase の歴史を共有できます。リプレイでは、
従来の Magpie ノード タイプでは、martin.* ノードをスキップしながら、そのノードに進みます。
ルート、およびその他の不明なノード タイプまたは不正な形式の Magpie のフェール クローズ
イベント。 magpie init は、外国語のみの履歴の後に Magpie ブートストラップを追加します。
ブートストラップが存在すると、べき等のままになります。
最初にビルドせずに開発するには、次を使用します。
./cmd/magpie \ を実行します。
--store .magpie \
--actor AGENT_USER_ID \
コマンド...
エージェントの運用ルール:
標準出力を唯一の成功チャネルとして扱います。
stderr を JSON エラー出力として扱います。
.magpie/ ファイルを直接編集しないでください。
決して生のストレージの突然変異を考え出さないでください。
財務活動を投稿する前に、書籍の設定を読んでください。
書籍全体に対してアクティブなaccounting_basisを使用します。現金、修正現金、取引ごとの見越額は選択しないでください。
アカウントの意味を決定するときは、アカウント名や番号ではなくアカウントの役割を使用してください。
invoice create-json 、invoice pos を使用します

t 、および顧客の請求書アクティビティに対して支払われた請求書マーク。
通常の業務活動には総勘定元帳仕訳帳作成を使用しないでください。これは、特権のある手動調整/インポート パスです。
シェルの引用の問題を回避するには、長いノート本文の場合は note put --body-file FILE を使用します。
大規模なエージェント ワークフローの前に、スナップショット create --name NAME を作成します。
{ "code" : " Permission_denied " 、 "message" : " 役割 \" 操作 \" には台帳がありません:読み取り " }
初回セットアップ
./magpie --store .magpie init
または、JAYBASE_URL とライターを設定した後、空のホスト ストアを初期化します。
JAYBASE_TOKEN :
./magpie --アクター オーナーの初期化
デフォルトの初期化アクターは、所有者ロールを持つ所有者です。
./magpie --store .magpie --actor 所有者 rbac 権限
制約されたロールを持つエージェント ユーザーを作成します。
./magpie --store .magpie --actor 所有者 rbac ユーザー セット \
--id ops-agent \
--role 操作
組み込みの役割:
管理者 : リカバリを除く幅広い操作アクセス権。
会計士 : 元帳、メモの読み取り、監査の読み取り。
操作: メモの読み取り/書き込みのみ。
営業担当者 : 読み取り/書き込みのみをメモします。
新しい組み込み権限が追加される前に初期化されたストアは、カスタム ロールやユーザーを変更せずにデフォルトのロールを修復できます。
./magpie --store .magpie --actor owner rbac デフォルト修復
Repair コマンドには rbac:manage が必要で、不足している現在のデフォルト権限を組み込みロールに追加し、それらのロールに対する既存の追加権限を保持します。
./magpie --store .magpie --actor 所有者 rbac ロール セット \
--name " Notes エージェント " \
--権限 メモ:読み取り、メモ:書き込み
次に、それを割り当てます。
./magpie --store .magpie --actor 所有者 rbac ユーザー セット \
--id メモ-エージェント \
--role " Notes エージェント "
メモのワークフロー
./magpie --store .magpie --actor Notes-agent note put \
--title " 作戦ハンドオフ " \
--body "毎日締め切り出荷します。"
長いコンテンツの場合:
./magpie --store .magpie --actor ノート-エージェント

メモを置く\
--title " ウィークリーレビュー " \
--body-file ./weekly-review.md \
-- 内部感度
リストのメモ:
./magpie --store .magpie --actor メモ-エージェント メモ リスト
特定のメモを読んでください。
./magpie --store .magpie --actor Notes-agent note get --id note:...
帳簿上の会計基準
カササギは会計方法について意見が分かれています。この本には、アクティブな会計基礎が 1 つだけあります。
現金 : 現金を受け取ったときに収入を認識し、現金を支払ったときに支出を認識します。
Modified_cash : 売上税負債、給与税負債、ローン元本、および資本化された固定資産の明示的な貸借対照表処理を伴う、経常収益および支出の現金処理。
見越: 必要に応じて売掛金と買掛金を使用して、収益を獲得または請求さ​​れた時点で収益を認識し、発生または請求さ​​れた時点で費用を認識します。
新しい店舗のデフォルトは現金です。エージェントがエントリを投稿する前に、現在の設定を確認してください。
./magpie --store .magpie --actor オーナーブックの設定を取得
仕訳活動を入力する前に会計基準を設定します。
./magpie --store .magpie --actor オーナーブック設定セット \
--会計ベースの見越額
会計基準を変更するには、デフォルトの所有者ロールと管理者ロールが持つ settings:manage が必要です。カササギはbを拒否します

[切り捨てられた]

## Original Extract

Accounting CLI built on Jaybase. Contribute to kyle-visner/magpie development by creating an account on GitHub.

Hi HN, Annoucing the release of Magpie, a CLI tool which turns your AI agent into an expert bookkeeper, based on Jaybase. ( https://news.ycombinator.com/item?id=48999936 ) As someone who owns a small, tech business, I have very simple books. somewhere between 10-30 transactions a month, a few depreciation line items, and a couple budgets to monitor. So it always bothered be that I was paying hundreds of dollars per year for software, and maybe a live bookeeper on top of that, just to keep these simple things in order, because I didn't have the mental bandwidth to do it myself. So I built Magpie. Magpie is a simple, opinionated CLI that runs on top of Jaybase to allow any popular AI agent to do proper double entry bookkeeping. It enforces standard accounting rules, stores everything securely, and you can ask your AI for just about any kind or report you want. I built it initially for myself, but I figured it's something the broader community could use. So give it a try. Simply give Claude Code, Codex, or any other agent the URL and it should be able to set it up for you.

GitHub - kyle-visner/magpie: Accounting CLI built on Jaybase · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
kyle-visner
/
magpie
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
47 Commits 47 Commits .github/ workflows .github/ workflows cmd/ magpie cmd/ magpie docs docs internal/ magpie internal/ magpie scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum llm.md llm.md View all files Repository files navigation
Magpie is an opinionated accounting CLI for humans and AI agents. It provides
one guarded path for chart-of-accounts management, double-entry journals,
customers, invoices, payouts, notes, snapshots, and audit reads. The CLI checks
RBAC and accounting invariants before appending an encrypted, immutable event to
Jaybase .
Requires Go 1.26.5 or later. Earlier Go releases include known standard-library
vulnerabilities and must not be used to build release binaries:
git clone https://github.com/kyle-visner/magpie.git
cd magpie
go install ./cmd/magpie
magpie --store .magpie init
The initialized local book uses cash accounting and creates an owner actor
with the Owner role. Read First-Time Setup before adding
another actor or posting financial activity.
Magpie is for small teams that want agents to help with bookkeeping without
giving them raw database access or permission to invent accounting behavior. It
fits source-document imports, controlled migrations, invoice and payout
workflows, operational notes, and other jobs where writes must be attributable,
auditable, safe to retry, and rejected when they violate policy.
Magpie is a CLI and domain engine, not a hosted accounting SaaS or a substitute
for professional accounting judgment. Use the local backend for development or
a trusted single-user process. For shared or production use, run Magpie against
a separately deployed Jaybase service and bind authenticated callers to allowed
Magpie actor IDs.
Magpie is pre-1.0. The implemented accounting surface is usable and tested.
This release includes the capabilities listed below. The following features are
outside its current scope:
native QuickBooks CSV, IIF, or QBXML parsing; agents must normalize source
data into Magpie's JSON contracts;
bills, bank matching or reconciliation, period close, tax, loan, transfer,
fixed-asset, retention, garbage-collection, or point-in-time restore commands;
note search, backlinks, typed cross-entity references, diff, or graph
navigation;
a human-oriented output mode, interactive UI, signed command envelopes, or an
authentication layer that proves --actor identity.
Those are product-scope boundaries, not hidden installation steps. See
docs/SECURITY.md for additional production controls that
remain the deployer's responsibility.
The repository root is the accounting CLI project.
cmd/magpie contains the CLI and internal/magpie contains the accounting
domain.
Jaybase is maintained separately at
github.com/kyle-visner/jaybase .
Magpie pins it as a Go module dependency.
AGPL-3.0-or-later. See LICENSE .
Canonical CLI in cmd/magpie with local embedded and hosted Jaybase backends.
Append-only, SHA-256-addressed event history through Jaybase.
Bearer-authenticated hosted Jaybase access over HTTPS with paginated replay, optimistic concurrency, idempotent writes, and remote named refs.
AES-256-GCM encryption for stored node payloads.
Unified RBAC for ledger, notes, snapshots, and audit reads.
Double-entry ledger validation before persistence.
Book-level accounting basis support for cash, modified cash, and accrual accounting.
Chart account roles for workflow-safe account selection.
Privileged manual journal adjustments with required audit reasons.
First-class customer, invoice, and payout workflows that generate basis-aware journals.
Structured external source references on ledger accounts.
Markdown note create, update, list, and get operations.
Source-tagged journal entries for agent-mapped exports from QuickBooks or other systems.
Named snapshots for recoverable roots.
JSON command output by default for agent consumption.
Automated tests for business invariants, CLI behavior, and BDD-style core scenarios.
Use Go 1.26.5 or later. From the repository root:
go mod verify
go test -race ./...
go vet ./...
go build -o ./magpie ./cmd/magpie
go install golang.org/x/vuln/cmd/govulncheck@v1.6.0
govulncheck ./...
If your environment blocks the default Go build cache, use a writable cache:
GOCACHE=/private/tmp/magpie-gocache go mod verify
GOCACHE=/private/tmp/magpie-gocache go test -race ./...
GOCACHE=/private/tmp/magpie-gocache go vet ./...
GOCACHE=/private/tmp/magpie-gocache go build -o ./magpie ./cmd/magpie
The generated ./magpie binary is ignored by Git.
Tagged releases are built from clean main commits by
release.yml using the reproducible
scripts/build-release.sh process. Each GitHub
Release includes macOS and Linux archives for amd64 and arm64 plus a
SHA256SUMS file.
Run the accounting benchmarks from this repository. Run Jaybase benchmarks from the separate Jaybase repository.
GOCACHE=/private/tmp/magpie-gocache go test -run ' ^$ ' -bench . -benchmem ./...
For baseline comparisons, capture multiple runs and compare them with benchstat :
mkdir -p .benchmarks
GOCACHE=/private/tmp/magpie-gocache go test -run ' ^$ ' -bench . -benchmem -count 5 ./... > .benchmarks/magpie-main.txt
benchstat .benchmarks/magpie-main.txt .benchmarks/magpie-feature.txt
Agent Integration Pattern
Agents should read llm.md as their operating contract. The short
pattern below shows the required invocation shape.
Give your agent a fixed command template and tell it to parse stdout as JSON:
./magpie \
--store .magpie \
--actor AGENT_USER_ID \
COMMAND...
For the hosted Jaybase service, provide the origin and bearer credential through
the environment. Do not put the token in a command-line flag, URL, payload, log,
or idempotency key:
export JAYBASE_URL=https://jaybase.example.com
export JAYBASE_TOKEN= ' writer-token-from-the-secret-manager '
./magpie \
--actor AGENT_USER_ID \
COMMAND...
--jaybase-url may override the origin, but the token is accepted only through
JAYBASE_TOKEN . Hosted requests fetch decrypted payloads only when replaying
state; audit output remains metadata-only. Writes use Jaybase's expected_root
and Idempotency-Key contract and return a conflict instead of overwriting a
newer root.
Magpie can share one linear Jaybase history with Martin. Replay applies the
legacy Magpie node types, skips martin.* nodes while still advancing to their
roots, and fails closed for other unknown node types or malformed Magpie
events. magpie init adds the Magpie bootstrap after a foreign-only history and
remains idempotent once that bootstrap exists.
For development without building first, use:
go run ./cmd/magpie \
--store .magpie \
--actor AGENT_USER_ID \
COMMAND...
Operational rules for agents:
Treat stdout as the only success channel.
Treat stderr as JSON error output.
Never edit .magpie/ files directly.
Never invent raw storage mutations.
Read book settings get before posting financial activity.
Use the active accounting_basis for the whole book; do not choose cash, modified cash, or accrual per transaction.
Use account roles rather than account names or numbers when deciding what an account means.
Use invoice create-json , invoice post , and invoice mark-paid for customer invoice activity.
Do not use generic ledger journal create for ordinary operating activity. It is a privileged manual adjustment/import path.
Use note put --body-file FILE for long note bodies to avoid shell quoting issues.
Create a snapshot create --name NAME before large agent workflows.
{ "code" : " permission_denied " , "message" : " role \" Operations \" lacks ledger:read " }
First-Time Setup
./magpie --store .magpie init
Or initialize an empty hosted store after setting JAYBASE_URL and a writer
JAYBASE_TOKEN :
./magpie --actor owner init
The default initialized actor is owner with the Owner role.
./magpie --store .magpie --actor owner rbac permissions
Create an agent user with a constrained role:
./magpie --store .magpie --actor owner rbac user set \
--id ops-agent \
--role Operations
Built-in roles:
Admin : broad operational access except recovery.
Accountant : ledger, notes read, and audit read.
Operations : notes read/write only.
Sales Rep : notes read/write only.
Stores initialized before new built-in permissions were added can repair default roles without changing custom roles or users:
./magpie --store .magpie --actor owner rbac defaults repair
The repair command requires rbac:manage , adds missing current default permissions to built-in roles, and preserves any existing extra permissions on those roles.
./magpie --store .magpie --actor owner rbac role set \
--name " Notes Agent " \
--permissions notes:read,notes:write
Then assign it:
./magpie --store .magpie --actor owner rbac user set \
--id notes-agent \
--role " Notes Agent "
Notes Workflow
./magpie --store .magpie --actor notes-agent note put \
--title " Ops Handoff " \
--body " Ship daily closeout. "
For longer content:
./magpie --store .magpie --actor notes-agent note put \
--title " Weekly Review " \
--body-file ./weekly-review.md \
--sensitivity internal
List notes:
./magpie --store .magpie --actor notes-agent note list
Read a specific note:
./magpie --store .magpie --actor notes-agent note get --id note:...
Book Accounting Basis
Magpie is opinionated about accounting methods. The book has exactly one active accounting basis:
cash : recognize income when cash is received and expenses when cash is paid.
modified_cash : cash treatment for ordinary income and expenses, with explicit balance-sheet treatment for sales tax liabilities, payroll tax liabilities, loan principal, and capitalized fixed assets.
accrual : recognize revenue when earned or invoiced and expenses when incurred or billed, using accounts receivable and accounts payable where appropriate.
New stores default to cash . Check the current setting before an agent posts entries:
./magpie --store .magpie --actor owner book settings get
Set the accounting basis before entering journal activity:
./magpie --store .magpie --actor owner book settings set \
--accounting-basis accrual
Changing the accounting basis requires settings:manage , which the default Owner and Admin roles have. Magpie rejects b

[truncated]

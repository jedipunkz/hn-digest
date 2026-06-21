---
source: "https://github.com/Jake1848/conduit"
hn_url: "https://news.ycombinator.com/item?id=48622458"
title: "Conduit – Self-hosted Bitcoin Lightning payments for AI agents"
article_title: "GitHub - Jake1848/conduit: Conduit · GitHub"
author: "conduitbtc"
captured_at: "2026-06-21T21:28:02Z"
capture_tool: "hn-digest"
hn_id: 48622458
score: 1
comments: 0
posted_at: "2026-06-21T20:48:15Z"
tags:
  - hacker-news
  - translated
---

# Conduit – Self-hosted Bitcoin Lightning payments for AI agents

- HN: [48622458](https://news.ycombinator.com/item?id=48622458)
- Source: [github.com](https://github.com/Jake1848/conduit)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T20:48:15Z

## Translation

タイトル: Conduit – AI エージェント向けのセルフホスト型ビットコイン ライトニング支払い
記事タイトル: GitHub - Jake1848/コンジット: コンジット · GitHub
説明: 導管。 GitHub でアカウントを作成して、Jake1848/conduit の開発に貢献してください。

記事本文:
GitHub - Jake1848/コンジット: コンジット · GitHub
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
ジェイク1848
/
導管
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらに開く アクション メニュー フォルダー a

ndファイル
61 コミット 61 コミット .github .github core core ダッシュボード ダッシュボード docs docs infra infra mcp-server mcp-server sdk-js sdk-js sdk-python sdk-python website website .env.example .env.example .env.prod.example .env.prod.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md Conduit_Whitepaper_v1.pdf Conduit_Whitepaper_v1.pdf DEMO.md DEMO.md ライセンス ライセンス QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md docker-compose.dev.yml docker-compose.dev.yml docker-compose.prod.yml docker-compose.prod.yml docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
自律型 AI エージェント向けの自己ホスト型 Bitcoin Lightning 支払いインフラストラクチャ。
Conduit は、独自のインフラストラクチャ上で実行するソフトウェア ツールです。
独自のキーを持つ独自の LND ノード。 AI エージェントに仮想的な機能を提供します。
Lightning ウォレット、支出ポリシー、および送信、受信、アカウント作成のための API
ビットコイン支払いの場合 — プログラム的に、ハードガードレールによりエージェントは決済できません
オーバーライドします。
Conduit SaaS は存在しません。Conduit SaaS をホストするのはあなたです。Conduit はあなたの資金を決して保持せず、決して保持しません。
家に電話する。オペレーターレベルでは、ノード、キー、自己ホスト型です。
ルール。作成するエージェントは、台帳内の仮想サブ残高です。
オペレーター、コントロール: 彼らは署名キーではなくスコープ付きの API キーを保持し、あなたは
いつでも入金、借方記入、または一括払いができます。
ステータス: v0.8.4 — テストネットとメインネット (testnet/regtest plus) でライブ実行中
ライブメインネットノード。最初の実際のメインネット支払いがエンドツーエンドで決済されました —
まだ初期段階で小規模で、単一オペレーターであり、大規模な生産ではありません）。オペレーター
財務 (0.8.3 で追加) を使用すると、未払いのプラットフォーム料金収入を確認して引き出すことができます
オンチェーンで発生したBTC、ソルベンシーガードによってゲートされている（ノード資産がこれを下回ることはありません）

エージェントに支払う義務があるもの）。 0.8.4 では、完全な監査/レッドチーム パスが追加されます。
オペレータ コンソール、およびエージェント支払いの Lightning デモ ( DEMO.md を参照)。
ウェブサイト: https://conduit.energy
ホストされたデモ API: https://api.conduit.energy (テストネット) · https://api-mainnet.conduit.energy (メインネット)
ドキュメント: https://docs.conduit.energy
LND ノードに対して Conduit (5 分間の Docker 起動) をデプロイします。導管
ポリシー + アカウンティング層として LND の前に配置されます。
あなたはLNDシードとマカロンを持っています。コンジットは LND とのみ通信します
そこにマカロンを乗せます。
あなたはオペレーターです。ブートストラップ API キーは、
独自のシステム — エージェントに渡す範囲指定されたキーを作成します。
で実行されているエージェントの仮想サブ残高を貸方記入および借方記入します。
あなたのノード — Conduit はその構造上、エージェントを管理します。
エージェント残高は、Conduit の台帳内でオペレーターが管理する IOU であり、
基礎となる衛星はキーの下のチャンネルに残ります。
Conduit が追加するトランザクションごとのプラットフォーム料金 (衛星単位) を設定します。
各支払いに加えて決済を継続します。その手数料があなたの収入となります。
ユーザーが設定したものであり、コンジットを切断することはありません。
Conduit に組み込まれた収益化は、少額の使用量ベースのプラットフォーム料金です。
SATOSHI は Conduit を導入するオペレーターが設定します。で充電されます
各支払いの上部、支払いが決済されたときに保管され、支払いが完了した場合は全額返金されます。
支払いが失敗します。土曜日のみ — 法定通貨、カード、定期購読はありません。
徴収された料金は GET /v1/fees (admin) で報告され、次の場所に表示されます。
/v1/metrics を取得します。料金なしで Conduit を実行するには、PLATFORM_FEE_PERCENT=0 を設定します。
。
§── Web サイト/ランディング ページ (conduit.energy にデプロイ)
§── core/ Conduit Core API — LND の前で実行する FastAPI サーバー
§── sdk-python/ Python SDK (PyPI 上の `conduit-btc`; `conduit` をインポート)
§─

─ sdk-js/ TypeScript SDK (npm の `@conduit-btc/sdk`)
§── mcp-server/ MCP サーバー (PyPI 上の `conduit-btc-mcp`; `conduit-mcp` コマンド)
§── インフラ/ビットコインコア/LND 設定、systemd ユニット、インストールスクリプト
§── docs/ MkDocs ドキュメント サイト
§── docker-compose.yml
└── Conduit_Whitepaper_v1.pdf
5つのコンポーネント
#
コンポーネント
ステータス
パス
1
ビットコインコア + LNDノード
インストール スクリプト - 独自のホストで実行します
インフラ/
2
コンジットコアAPI
FastAPI アプリ、ローカル開発用のモック LND で実行
コア/
3
Python SDK
ウェブサイトのコードパネルと一致します
SDK-Python/
4
TypeScript SDK
Pythonインターフェースをミラーリングします
SDK-JS/
5
MCPサーバー
Conduit を MCP ツールとして公開します
mcp-サーバー/
クイックスタート (ローカル、モック LND)
# 1. API をモック モードで起動します (LND_MOCK=true、SQLite、ブートストラップ キー プリセット)
docker compose -f docker-compose.dev.yml up --build
# 2. API をヒットする
カール http://localhost:8000/v1/health
# 3. Python SDK を使用する
cd sdk-python && pip install -e 。
エクスポート CONDUIT_API_KEY=ck_test_dev_root
Python の例/quickstart.py
開発コンテナには、自動作成された admin スコープの LND_MOCK=true が同梱されています。
ブートストラップ キー ck_test_dev_root (dev のマスター キー)、および SQLite
/app/data/conduit.db 。 Conduit に本物を向けるまで、実際のビットコインは動きません
独自の LND ノード。
Conduit は現在、 testnet 、 regtest 、および mainnet 上で自己ホスト可能です —
3人全員がライブで実行します。メインネットが現実になりました: 実際のチャネルを備えたニュートリノ LND ノード
が開始され、最初のリアルマネー支払いがエンドツーエンドで決済されました。それはまだです
初期かつ小規模であるため、メインネットの立ち上げを新しい領域として扱い、テストしてください
まずはテストネット。自分のライトニングに対する操作順序
インフラストラクチャ:
ビットコイン コア (プルーニング) をホストにインストールします — infra/scripts/install_bitcoind.sh
LND をインストールします (今日のテストネット、メインネットはサポートされています) — infra/scri

pts/install_lnd.sh
UFW の構成 — infra/scripts/setup_firewall.sh
チェーン同期を待ちます (1 ～ 3 日)
オープンチャネル — infra/scripts/setup_channels.sh
コア API を LND マカロンと TLS 証明書に向けます ( LND_MOCK=false )
PLATFORM_FEE_* の値を請求したい収益に設定します。
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d でデプロイします
メインネットは稼働していますが、初期段階で小規模です。今日の展開は単一のニュートリノです
1 つの ~20k-sat チャネルを持つ LND ノード。最初のリアルマネー支払いが完了しました
エンドツーエンド: Lightning キー送信経由で送信された 2000 衛星、10 衛星のプラットフォーム料金 (0.5%)
ライフサイクル全体で、1 衛星のルーティング料金を徴収、約 15 秒で決済
検証済み (引き落とし → ルート → 決済 → プラットフォーム料金の徴収 → 失敗時の返金 →
正確な台帳照合）。これは初めてのリアルマネー検証です。
大規模な生産。スループットや戦闘テストを想定しないでください。ライブ
メインネット API は https://api-mainnet.conduit.energy にあります。
完全なランブック: infra/README.md 。
Conduit は構築によって自己ホストされます。これはユーザーが操作するソフトウェアであり、ユーザーが操作するものではありません。
お金を預けるサービス。誰が何を管理するのかを明確にしてください。
オペレーターレベルでは自己ホスト型です。 Conduit SaaS はありません。
Conduit は資金を保持したり、自宅に電話をかけたりすることはありません。 LND シード フレーズは次のとおりです。
お客様のものであり、Conduit によって VPS に保存されることはありません。サットはそこに留まる
キーの下にあるチャンネル。
エージェント レベルでは、Conduit が保管されます。エージェント残高は仮想的です
Conduit の台帳にある IOU は、オペレーターとして貸方記入、借方記入が可能です。
掃引する。エージェントは署名キーではなく、スコープ付きの API キーを保持します。エージェントは決してアクセスしません。
ビットコインの秘密鍵。
LND gRPC / REST は決して公開されません。Conduit API のみが公開されます。
Conduit ポリシー エンジンは、支払いがお客様に届く前にすべての支払いを評価します。
LND および任意の e でのフェールクローズ

エラー
API キーは保存時に bcrypt ハッシュされ、オペレーターに 1 回だけ表示されます。
ブートストラップ API キーはマスター キーです。マカロンのように保護してください。
すべてのトランジットは HTTPS です。 Webhook 配信は HMAC 署名されています
コンジットは、制御するノードの前にあるポリシー + アカウンティング層です。もしあなたが
Conduit をオフにすると、衛星はチャンネル内に残ります。
infra/README.md→「セキュリティチェックリスト」を参照してください。
認可モデル (現在はシングルオペレーター)
Conduit での認可はエージェントごとではなくスコープベースです。 API キーは
3 つのスコープ (読み取り < 書き込み < 管理) のうちの 1 つであり、そのスコープが唯一のスコープです。
境界:
読み取りキーは、フリート全体 (すべてのエージェント、バランス、および
取引。
書き込みキーは任意のエージェントに作用できます。つまり、支払いを送信したり、エージェントを作成したりできます。
台帳内の任意のエージェントの請求書。
管理者キーは、エージェントの作成/削除、残高の移動、ポリシーの管理、および
Webhook とインスタンス全体のミント キー。
エージェントごとの境界はありません。キーは特定のエージェントに関連付けられていません。
ルートフィルタは、どのキーがエージェントを作成したか、またはエージェントを「所有」したかによってフィルタされます。エージェントは、
会計とポリシーの単位であり、相互間のセキュリティ境界ではありません
不信感を抱いている当事者。
具体的には、現在の Conduit はシングルオペレーターのツールです。スコープ付きキーを手渡しして、
お互いに信頼していない第三者に対してではなく、あなたが運営するエージェントに対して。もしあなたが
テナント間でハード分離が必要な場合は、別の Conduit インスタンス (およびノード) を実行する
テナントごとに。
ロードマップ: マルチテナント、エージェントごとのスコープ設定 - キーが 1 つにバインドされる場合
エージェント (またはエージェントのセット) が存在し、残りのフリートを読み取ったり操作したりすることはできません。
計画された機能強化。それに向けた最初のステップとして、create_agent は記録を行うようになりました。
Agent.api_key_id の鋳造キーの出所を確認します。これが由来です
のみ : まだ承認を変更したり強制したりしません。
コンジット-シアン-xi.vercel.app
リソース
そこw

as an error while loading.このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
5
v0.8.5 — Pre-launch hardening
最新の
2026 年 6 月 14 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Conduit. Contribute to Jake1848/conduit development by creating an account on GitHub.

GitHub - Jake1848/conduit: Conduit · GitHub
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
Jake1848
/
conduit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
61 Commits 61 Commits .github .github core core dashboard dashboard docs docs infra infra mcp-server mcp-server sdk-js sdk-js sdk-python sdk-python website website .env.example .env.example .env.prod.example .env.prod.example .gitignore .gitignore .gitleaks.toml .gitleaks.toml CHANGELOG.md CHANGELOG.md Conduit_Whitepaper_v1.pdf Conduit_Whitepaper_v1.pdf DEMO.md DEMO.md LICENSE LICENSE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md docker-compose.dev.yml docker-compose.dev.yml docker-compose.prod.yml docker-compose.prod.yml docker-compose.yml docker-compose.yml View all files Repository files navigation
Self-hosted Bitcoin Lightning payment infrastructure for autonomous AI agents.
Conduit is software tooling you run on your own infrastructure, in front of
your own LND node, with your own keys. It gives any AI agent a virtual
Lightning wallet, a spending policy, and an API to send, receive, and account
for Bitcoin payments — programmatically, with hard guardrails the agent cannot
override.
There is no Conduit SaaS: you host it, Conduit never holds your funds and never
phones home. At the operator level it's self-hosted — your node, your keys, your
rules. The agents you create are virtual sub-balances in a ledger that you,
the operator, control: they hold a scoped API key, not a signing key, and you
can credit, debit, or sweep them at any time.
Status: v0.8.4 — running live on testnet and mainnet (testnet/regtest plus
a live mainnet node; the first real mainnet payment has settled end-to-end —
still early and small, single-operator, not production-at-scale). The operator
treasury (added in 0.8.3) lets you see accrued platform-fee revenue and withdraw
accrued BTC on-chain, gated by a solvency guard (node assets can never drop below
what you owe agents). 0.8.4 adds a full audit/red-team pass, the complete
operator console, and an agent-pays-over-Lightning demo (see DEMO.md ).
Website: https://conduit.energy
Hosted demo API: https://api.conduit.energy (testnet) · https://api-mainnet.conduit.energy (mainnet)
Docs: https://docs.conduit.energy
You deploy Conduit (a 5-minute Docker bring-up) against your LND node. Conduit
sits in front of LND as a policy + accounting layer:
You hold the LND seed and macaroon. Conduit only ever talks to LND over
the macaroon you mount into it.
You are the operator. The bootstrap API key is your master key to
your own system — it mints the scoped keys you hand to your agents.
You credit and debit the virtual sub-balances of the agents running on
your node — Conduit is custodial for the agents by construction : the
agent balances are operator-controlled IOUs in Conduit's ledger, and the
underlying sats stay in your channels under your keys.
You set a per-transaction platform fee (in sats) that Conduit adds on
top of each payment and keeps on settlement — that fee is your revenue,
configured by you, never a Conduit cut.
Conduit's built-in monetization is a small, usage-based platform fee in
satoshis that the operator who deploys Conduit configures. It is charged on
top of each payment, kept when the payment settles, and refunded in full if the
payment fails. Sats only — no fiat, no cards, no subscription.
Collected fees are reported at GET /v1/fees (admin) and surfaced in
GET /v1/metrics . Set PLATFORM_FEE_PERCENT=0 to run Conduit with no fee.
.
├── website/ Landing page (deployed to conduit.energy)
├── core/ Conduit Core API — FastAPI server you run in front of your LND
├── sdk-python/ Python SDK (`conduit-btc` on PyPI; import `conduit`)
├── sdk-js/ TypeScript SDK (`@conduit-btc/sdk` on npm)
├── mcp-server/ MCP server (`conduit-btc-mcp` on PyPI; `conduit-mcp` command)
├── infra/ Bitcoin Core / LND configs, systemd units, install scripts
├── docs/ MkDocs documentation site
├── docker-compose.yml
└── Conduit_Whitepaper_v1.pdf
The five components
#
Component
Status
Path
1
Bitcoin Core + LND nodes
install scripts — you run them on your own host
infra/
2
Conduit Core API
FastAPI app, runs with mock-LND for local dev
core/
3
Python SDK
matches the website code panel
sdk-python/
4
TypeScript SDK
mirrors the Python interface
sdk-js/
5
MCP server
exposes Conduit as MCP tools
mcp-server/
Quickstart (local, mock LND)
# 1. Bring up the API in mock mode (LND_MOCK=true, SQLite, bootstrap key preset)
docker compose -f docker-compose.dev.yml up --build
# 2. Hit the API
curl http://localhost:8000/v1/health
# 3. Use the Python SDK
cd sdk-python && pip install -e .
export CONDUIT_API_KEY=ck_test_dev_root
python examples/quickstart.py
The dev container ships with LND_MOCK=true , an auto-created admin -scoped
bootstrap key ck_test_dev_root (your master key in dev), and SQLite at
/app/data/conduit.db . No real Bitcoin moves until you point Conduit at a real
LND node of your own.
Conduit is self-hostable today on testnet , regtest , and mainnet —
all three run live. Mainnet is now real: a neutrino LND node with a real channel
is up and the first real-money payment has settled end-to-end. It is still
early and small, so treat a mainnet bring-up as new territory and test on
testnet first. The order of operations against your own Lightning
infrastructure:
Install Bitcoin Core (pruned) on your host — infra/scripts/install_bitcoind.sh
Install LND (testnet today; mainnet supported) — infra/scripts/install_lnd.sh
Configure UFW — infra/scripts/setup_firewall.sh
Wait for chain sync (1–3 days)
Open channels — infra/scripts/setup_channels.sh
Point the Core API at your LND macaroon and TLS cert ( LND_MOCK=false )
Set your PLATFORM_FEE_* values to whatever revenue you want to charge
Deploy with docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
Mainnet is live but early and small. The deployment today is a single neutrino
LND node with one ~20k-sat channel. The first real-money payment has settled
end-to-end: 2000 sats sent via Lightning keysend, a 10-sat platform fee (0.5%)
collected, 1-sat routing fee, settled in ~15s — with the full lifecycle
verified (debit → route → settle → platform-fee capture → refund-on-failure →
exact ledger reconciliation). This is a first real-money validation, not
production-at-scale; do not assume throughput or battle-testing. The live
mainnet API is at https://api-mainnet.conduit.energy .
Full runbook: infra/README.md .
Conduit is self-hosted by construction — it is software you operate, not a
service that holds your money. Be clear about who custodies what:
At the operator level you are self-hosted. There is no Conduit SaaS;
Conduit never holds your funds and never phones home. The LND seed phrase is
yours and is never stored on the VPS by Conduit; the sats stay in
channels under your keys.
At the agent level Conduit is custodial. Agent balances are virtual
IOUs in Conduit's ledger that you, the operator, credit, debit, and can
sweep. Agents hold a scoped API key, not a signing key — they never touch a
Bitcoin private key.
LND gRPC / REST is never exposed publicly — only your Conduit API is
The Conduit policy engine evaluates every payment before it reaches your
LND and fails closed on any error
API keys are bcrypt-hashed at rest and shown to you, the operator, exactly once
The bootstrap API key is your master key — guard it like the macaroon
All transit is HTTPS; webhook deliveries are HMAC-signed
Conduit is the policy + accounting layer in front of a node you control. If you
turn Conduit off, your sats are still in your channels.
See infra/README.md → "Security checklist".
Authorization model (single-operator today)
Authorization in Conduit is scope-based, not per-agent . An API key carries
one of three scopes — read < write < admin — and that scope is the only
boundary:
A read key can read the entire fleet — every agent, balance, and
transaction.
A write key can act on any agent — send a payment from, or create an
invoice for, any agent in the ledger.
An admin key can create/delete agents, move balances, manage policies and
webhooks, and mint keys across the whole instance.
There is no per-agent boundary : a key is not tied to a specific agent, and
no route filters by which key created or "owns" an agent. Agents are an
accounting and policy unit, not a security boundary between mutually
distrusting parties .
Concretely: Conduit today is a single-operator tool. Hand scoped keys to
agents you run, not to third parties you don't trust with each other. If you
need hard isolation between tenants, run a separate Conduit instance (and node)
per tenant.
Roadmap: multi-tenant, per-agent scoping — where a key is bound to one
agent (or set of agents) and cannot read or act on the rest of the fleet — is
a planned enhancement. As a first step toward it, create_agent now records
the minting key on Agent.api_key_id for provenance. This is provenance
only : it does not yet change or enforce any authorization.
conduit-cyan-xi.vercel.app
Resources
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
5
v0.8.5 — Pre-launch hardening
Latest
Jun 14, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

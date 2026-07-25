---
source: "https://github.com/ToolJet/ActionRail/"
hn_url: "https://news.ycombinator.com/item?id=49051411"
title: "Show HN: ActionRail, Runtime value/action grounding framework for AI agents"
article_title: "GitHub - ToolJet/ActionRail: Open-source runtime action grounding framework for AI agents. Verify proposed actions against live systems of record before execution. · GitHub"
author: "oss-dev"
captured_at: "2026-07-25T21:44:50Z"
capture_tool: "hn-digest"
hn_id: 49051411
score: 2
comments: 0
posted_at: "2026-07-25T20:48:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ActionRail, Runtime value/action grounding framework for AI agents

- HN: [49051411](https://news.ycombinator.com/item?id=49051411)
- Source: [github.com](https://github.com/ToolJet/ActionRail/)
- Score: 2
- Comments: 0
- Posted: 2026-07-25T20:48:25Z

## Translation

タイトル: Show HN: ActionRail、AI エージェント用のランタイム値/アクション基盤フレームワーク
記事のタイトル: GitHub - ToolJet/ActionRail: AI エージェント用のオープンソース ランタイム アクション グラウンディング フレームワーク。実行前に、提案されたアクションをライブ記録システムに対して検証します。 · GitHub
説明: AI エージェント用のオープンソースのランタイム アクション基盤フレームワーク。実行前に、提案されたアクションをライブ記録システムに対して検証します。 - ツールジェット/アクションレール

記事本文:
GitHub - ToolJet/ActionRail: AI エージェント用のオープンソースのランタイム アクション基盤フレームワーク。実行前に、提案されたアクションをライブ記録システムに対して検証します。 · GitHub
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
ああ、ああ！
エラーが発生しました

ルをロードしています。このページをリロードしてください。
ツールジェット
/
アクションレール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
279 コミット 279 コミット .claude .claude .github .github actionrail actionrail control-plane control-plane docs docs 例 例 フロントエンド フロントエンド 出力/ pdf 出力/ pdf スクリプト スクリプト テスト テスト tmp/ pdfs tmp/ pdfs .coveragerc .coveragerc .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md mutmut_pytest.ini mutmut_pytest.ini pyproject.toml pyproject.toml pytest.ini pytest.ini すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントが何をしようとしているのかを、実行前に確認します。
ActionRail は、ToolJet によるオープンソース プロジェクトです。
リリースステータス: ActionRail は現在パブリックベータ版です。インターフェースは、
2026 年 8 月に予定されている ActionRail v1 より前に進化します。
統合サポート: あらゆる Python アプリケーションでフレームワークに依存しない
アクションランタイム。 LangGraph には、自動ツール検出とラッピング機能もあります。
追加の自動フレームワーク アダプターが計画されています。
エージェントのツール呼び出しは完全に許可され、完全に整形式であり、
まだ間違っている: すでに返金された注文に対する返金、または別の注文への転送
間違った顧客に属する実際の口座、引き出し金額を超える金額
バランス。形状が次のとおりであるため、ホワイトリストとスキーマ チェックはこれらすべてに合格します。
問題ありません。間違っているのは値です。
ActionRail は、エージェントの結果として生じるツール呼び出しをインターセプトし、実行前に、
あなたのライブシステムのrecに対する各議論の根拠

ord — その後、戻ります
許可 / 人間のためのホールド / ブロック 。それはラストマイルの正しさのゲートです: ではありません
エージェントはこれを許可しましたか？」でも、「この特定の値は本当に正しいのですか？」
さあ、実際のデータと照らし合わせて？」
ActionRail は実行境界で実行されます。ソースのルックアップと認証情報は維持されます
顧客環境では、許可の決定後にのみツールが実行されます。
ActionRail ランタイムはフレームワークに依存しません。 LangGraph は最初のアダプターです
これにより、自動ツール検出とラッピングが追加されます。
基本 SDK をインストールし、コンソールでエージェントとそのアクションを作成して、
実際の操作に関する ActionRuntime:
python -m pip install actionrail actionrail-console
actionrail import ActionRuntime から
ランタイム = アクションランタイム (
エージェントID = "…" 、
api_key = "ark_…" ,
)
結果 = 実行時。実行(
"issue_refund" ,
{ "order_id" : order_id 、 "amount" : 金額 },
発行_返金 、
context = { "顧客 ID" : 認証された顧客 ID },
)
execute() は、承認後にのみ issue_refund(**arguments) を呼び出します。使用する
カスタム エージェント ループ、ワーカー、API ハンドラー、またはフレームワーク内の同じ API
専用アダプター。 Python の直接統合を参照してください。
ActionRail でコンパイル済みのファイルを検出してラップする場合は、enforce() を使用します。
LangGraph エージェントのツール:
actionrail import enforce から、trusted_context
アクションレールから。 SDK 。接地インポートSQLiteSource
エージェント = 強制 (
代理店、
エージェントID = "…" 、
api_key = "ark_…" ,
ソース = { "billing" : SQLiteSource ( "billing.db" )},
)
Trusted_context ( customer_id = "C-1007" ) を使用:
代理店。呼び出し (入力)
# Wire_money(account="1234567") -> 基礎となるツールが実行される前にブロックされる
詳細な理由は、ローカル アプリケーション コードで引き続き利用できます。
on_decion コールバック。モデルに面したツールの結果、アウトバウンドアクティビティ、
レビューレコードは別の値を使用します

ueフリーの表​​現。
パッケージ化された LangGraph デモを実行する
完全な決定論的デモにはモデル API キーは必要ありません。
python3 -m venv .venv
ソース .venv/bin/activate
python -m pip install " actionrail[agents] " actionrail-console
actionrail-console --no-open
2 番目の端末で:
ソース .venv/bin/activate
アクションレールクイックスタート
デモでは、ローカル SQLite ソースを作成し、グラウンディング ルールをインストールし、
顧客間の返金、実際のツールが実行されなかったことを証明し、
アクティビティ内の編集されたブロック。 http://127.0.0.1:8020/activity を開いて調べてください。
実際のモデル駆動型 LangGraph を通じて同じ毒されたサポート ノートを実行するには
ループで、アカウントで利用可能な Anthropic API キーとモデル ID を指定します。
エクスポート ANTHROPIC_API_KEY= " … "
actionrail-quickstart --model クロード-ソネット-4-6
デモ ツールはローカルであり、無害です。モデルがクロスカスタマーを提案する場合
返金、ActionRail は認証された呼び出し元のコンテキストと照合して注文を検証し、
関数が実行される前にブロックします。モデルが注入されたノートに抵抗する場合、
司令部はその結果を報告し、同じ有害な提案を明らかに提出します。
ラベル付きの決定論的ランタイムチャレンジ。これは、ActionRail の境界をテストします。
プロポーズがモデルから来たふりをすることなく。決定論的モード
CI および初回評価の安定したパスとして残ります。
ポリシー エンジンと認証 (OPA、Cedar、ReBAC) は、アクションが有効であるかどうかを決定します。
事前にロードした事実を考慮すると許可されます。設計上、彼らは誰かに手を差し伸べることはありません
ライブ システムを実行し、呼び出し時に値をチェックします。たとえば、Cedar は
I/O を持たない副作用のない言語。 ActionRail はまさにそのギャップを埋めます。
呼び出し時に実際のデータをクエリし、引数を検証します。
私たちは、価値汚染攻撃を行う 4 つのプロバイダーにわたる 8 つのモデルをレッドチーム化しました。
を操縦する入力

整形式だが間違った引数の値に優しい。毎
モデルは保護されていないときに少なくとも 1 つのポイズニングされた値を実行し、攻撃は成功しました
1.7% から 63.3% の範囲で、より強力なモデルによってリスクが解消されるわけではありません。
ActionRail を前にすると、操作されたアクション 480 件中 0 件が実行され、そのうち 0 件が実行されました。
480 件の正当な類似リクエストが誤ってブロックされました。
価値中毒ベンチマークを読む → · PDF
決定()は最も安いチェックを最初に実行し、許可/保留/ブロックを返します
(ブロック outrank は outrank を許可します):
グラウンディングは表現力豊かです: 複数のチェックが異なるソース間で AND 演算され、それぞれ
返されたフィールドと演算子 ( = ≠ > ≥ < ≤ contains、is one of、is set、is empty ) を呼び出し側コンテキスト値、リテラル、別の引数と比較します。
(残高 ≥ amount )、または現在時刻 (delivery_at ≥ now-30d )。小切手
行 (expect: row ) またはその不在 (expect: missing ) を必要とすることができます。
ブロックリスト / 冪等性)、クエリ パラメータを名前でバインドします (SQLite では :customer_id、
PostgreSQL/MySQL では %(customer_id)s)、再試行 + フェイルオープン/クローズ ポリシーを実行します。
アクション テスト: 実行可能な契約としてのルール
ルールに沿って予想されるアクション動作をコミットし、それを実行します。
本番環境で使用されるのと同じ意思決定パイプライン:
# actionrail.cases.yaml
スキーマのバージョン : 1
ケース:
- 名前 : 複数の顧客間の返金はブロックされています
ツール : 発行_返金
引数: {order_id: order-100、数量: 75}
コンテキスト: {customer_id: customer-2}
期待：ブロック
備品：
請求-制作:
by_value :
order-100 : {customer_id: customer-1、ステータス: 配送済み}
アクションレールテスト
# PASS の顧客間の払い戻しはブロックされています
# 1 合格 · 0 不合格
アクション テストでは、レビュー可能な CI コントラクトを許可、保留、およびブロックします。
決定論的ソースフィクスチャは、実際のポリシー評価とグラウンディングを実行します。
モデル、コンソール、認証情報、o のないマッチャー

r ネットワークアクセス。生産
別の統合ステージが明示的に通過しない限り、ソースはオフのままです
--live-sources 。アクション テスト ガイドを参照してください。
または、examples/action_tests/ で完全なコントラクトを実行します。
環境内に留まるソース (データ プレーン) に対してアースを実行します。
VPC から離れることはありません）。現在: SQLite 、 PostgreSQL 、 MySQL/MariaDB 、
HTTP/REST および MCP ゲートウェイ アダプター。 MCP ソースは、設定された検証ツールを呼び出します。
ストリーミング可能な HTTP。 readOnlyHint が必要で、構造化されたツール出力を好みます。秘密は
${env:VAR} 参照として書き込まれ、ローカルで解決されるため、資格情報は決して取得されません
コントロールプレーンに到達します。読み取りツールのみを呼び出せるゲートウェイ ID を使用します。
MCP アノテーションはヒントであり、認可境界ではありません。
ソース:
ストライプ生産 :
アダプター: mcp
エンドポイント: https://gateway.internal/mcp
ヘッダー:
認可 : ベアラー ${env:MCP_GATEWAY_TOKEN}
ツール：ストライプ.get_charge
引数:
Charge_id : " {value} "
選択：データ・チャージ
タイムアウト: 10
require_read_only : true
これにより、MCP 接続の詳細がソース上に保持されます。ルールはソースを参照し、
返されたレコードを評価する方法のみを説明します。彼らは重複しません
ゲートウェイ エンドポイント、ツール名、または引数のマッピング。
PostgreSQL ソースは、ソースの接続詳細とルールのクエリを保持します。
SELECT のみの権限を持つ専用のロールを使用します。 ActionRail はそれぞれを強制します
トランザクションを PostgreSQL 読み取り専用モードに移行し、接続を適用します
そしてステートメントのタイムアウト。クエリでは、次のような Psycopg 名前付きパラメータが使用されます。
%(value)s 、 %(amount)s 、および %(customer_id)s :
ソース:
請求-制作:
アダプター: postgres
ホスト : postgres.internal
ポート: 5432
dbname : 請求
ユーザー: actionrail_reader
パスワード: ${env:POSTGRES_PASSWORD}
sslmode : フル検証
sslrootcert : /etc/ssl/certs/postgres-ca.pem

接続タイムアウト : 5
ステートメントタイムアウト_ms : 5000
require_read_only : true
SDK は Psycopg インターフェイスに依存しており、libpq の実装はそのままです。
ホストアプリケーション。自己完結型のローカル インストールの場合は、次を追加します
psycopg[バイナリ]>=3.2,<4 ;本番環境ではシステムにリンクされた
代わりに psycopg[c] をビルドします。
MySQL ソースと MariaDB ソースは同じ名前のクエリ パラメータを使用し、
すべてのチェックでトランザクション読み取り専用を開始します。 PyMySQL は SDK に含まれています。
ソース:
請求-mysql:
アダプター:mysql
ホスト : mysql.internal
ポート: 3306
データベース : 請求
ユーザー: actionrail_reader
パスワード: ${env:MYSQL_PASSWORD}
ssl_mode : アイデンティティの検証
ssl_ca : /etc/ssl/certs/mysql-ca.pem
接続タイムアウト : 5
クエリタイムアウト : 5
require_read_only : true
コントロールプレーン + コンソール
オプションの Django + DRF コントロール プレーンと React Console を使用して登録できます
ソース、引数ごとの基礎ルールを作成し、ライブ アクティビティを監視します。
ランタイムが作成したすべての許可/保留/ブロックのフィード。 SDK レポートのメタデータ
のみ - ツール名、結果、価値のない理由、および編集されたプレビュー。詳細
グラウンディングの結果とレンダリングされたプレビューは、ローカルの Decision でのみ公開されます。
モデルに面したツールの結果は、完全に編集された独自の表現を受け取ります。プレビュー
プレースホール

[切り捨てられた]

## Original Extract

Open-source runtime action grounding framework for AI agents. Verify proposed actions against live systems of record before execution. - ToolJet/ActionRail

GitHub - ToolJet/ActionRail: Open-source runtime action grounding framework for AI agents. Verify proposed actions against live systems of record before execution. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
ToolJet
/
ActionRail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
279 Commits 279 Commits .claude .claude .github .github actionrail actionrail control-plane control-plane docs docs examples examples frontend frontend output/ pdf output/ pdf scripts scripts tests tests tmp/ pdfs tmp/ pdfs .coveragerc .coveragerc .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md mutmut_pytest.ini mutmut_pytest.ini pyproject.toml pyproject.toml pytest.ini pytest.ini View all files Repository files navigation
Verify what an AI agent is about to do — before it does it.
ActionRail is an open-source project by ToolJet .
Release status: ActionRail is currently in public beta. Interfaces may
evolve before ActionRail v1, which is planned for August 2026.
Integration support: Any Python application can use the framework-neutral
ActionRuntime . LangGraph also has automatic tool discovery and wrapping;
additional automatic framework adapters are planned.
An agent's tool call can be perfectly permitted and perfectly well-formed and
still be wrong : a refund on an order that was already refunded, a transfer to a
real account that belongs to the wrong customer, a withdrawal larger than the
balance. Allowlists and schema checks pass all of these, because the shape is
fine — it's the value that's wrong.
ActionRail intercepts an agent's consequential tool calls and, before execution,
grounds each argument against your live system of record — then returns
allow / hold-for-human / block . It's the last-mile correctness gate: not "is
the agent allowed to do this?" but "is this specific value actually right, right
now, against your real data?"
ActionRail runs at the execution boundary. Source lookups and credentials stay
in the customer environment, and the tool runs only after an allow decision.
The ActionRail runtime is framework-neutral. LangGraph is the first adapter
that adds automatic tool discovery and wrapping.
Install the base SDK, create an agent and its actions in the Console, then put
ActionRuntime around the real operation:
python -m pip install actionrail actionrail-console
from actionrail import ActionRuntime
runtime = ActionRuntime (
agent_id = "…" ,
api_key = "ark_…" ,
)
result = runtime . execute (
"issue_refund" ,
{ "order_id" : order_id , "amount" : amount },
issue_refund ,
context = { "customer_id" : authenticated_customer_id },
)
execute() invokes issue_refund(**arguments) only after authorization. Use
the same API in custom agent loops, workers, API handlers, or frameworks without
a dedicated adapter. See the direct Python integration .
Use enforce() when you want ActionRail to discover and wrap a compiled
LangGraph agent's tools:
from actionrail import enforce , trusted_context
from actionrail . sdk . grounding import SQLiteSource
agent = enforce (
agent ,
agent_id = "…" ,
api_key = "ark_…" ,
sources = { "billing" : SQLiteSource ( "billing.db" )},
)
with trusted_context ( customer_id = "C-1007" ):
agent . invoke ( inputs )
# wire_money(account="1234567") -> blocked before the underlying tool runs
Detailed reasons remain available to local application code through the
on_decision callback. The model-facing tool result, outbound Activity, and
review records use separate value-free representations.
Run the packaged LangGraph demo
The complete deterministic demo needs no model API key:
python3 -m venv .venv
source .venv/bin/activate
python -m pip install " actionrail[agents] " actionrail-console
actionrail-console --no-open
In a second terminal:
source .venv/bin/activate
actionrail-quickstart
The demo creates a local SQLite Source, installs a grounding rule, attempts a
cross-customer refund, proves the real tool did not execute, and verifies the
redacted block in Activity. Open http://127.0.0.1:8020/activity to inspect it.
To run the same poisoned support note through a real model-driven LangGraph
loop, provide an Anthropic API key and a model ID available to your account:
export ANTHROPIC_API_KEY= " … "
actionrail-quickstart --model claude-sonnet-4-6
The demo tool is local and harmless. If the model proposes the cross-customer
refund, ActionRail verifies the order against authenticated caller context and
blocks it before the function runs. If the model resists the injected note, the
command reports that outcome and submits the same poisoned proposal as a clearly
labelled deterministic runtime challenge. This tests the ActionRail boundary
without pretending the proposal came from the model. The deterministic mode
remains the stable path for CI and first-time evaluation.
Policy engines and authz (OPA, Cedar, ReBAC) decide whether an action is
permitted given facts you've pre-loaded. By design they don't reach out to a
live system and check a value at call time — Cedar, for instance, is a
side-effect-free language with no I/O. ActionRail fills exactly that gap: it
queries your real data at the moment of the call and verifies the argument.
We red-teamed eight models across four providers with value-poisoning attacks:
inputs that steer an agent toward a well-formed but wrong argument value. Every
model executed at least one poisoned value when unprotected, with attack success
ranging from 1.7% to 63.3% — a stronger model did not make the risk go away.
With ActionRail in front, 0 of 480 manipulated actions executed and 0 of
480 legitimate look-alike requests were wrongly blocked.
Read the value-poisoning benchmark → · PDF
decide() runs cheapest-check-first and returns allow / hold / block
(block outranks hold outranks allow):
Grounding is expressive: multiple checks ANDed across different Sources, each
comparing a returned field with an operator ( = ≠ > ≥ < ≤ contains, is one of, is set, is empty ) against a caller-context value, a literal, another argument
( balance ≥ amount ), or the current time ( delivered_at ≥ now-30d ). Checks
can require a row ( expect: row ) or its absence ( expect: absent , for
blocklists / idempotency), bind query params by name ( :customer_id in SQLite,
%(customer_id)s in PostgreSQL/MySQL), and carry a retry + fail-open/closed policy.
Action Tests: rules as executable contracts
Commit expected action behavior alongside your rules and run it through the
same decision pipeline used in production:
# actionrail.cases.yaml
schema_version : 1
cases :
- name : cross-customer refund is blocked
tool : issue_refund
args : {order_id: order-100, amount: 75}
context : {customer_id: customer-2}
expect : block
fixtures :
billing-production :
by_value :
order-100 : {customer_id: customer-1, status: delivered}
actionrail test
# PASS cross-customer refund is blocked
# 1 passed · 0 failed
Action Tests make allow , hold , and block a reviewable CI contract.
Deterministic Source fixtures exercise the real policy evaluator and grounding
matcher without a model, Console, credentials, or network access. Production
Sources remain off unless a separate integration stage explicitly passes
--live-sources . See the Action Tests guide
or run the complete contract in examples/action_tests/ .
Grounding runs against sources that stay in your environment (the data plane
never leaves your VPC). Today: SQLite , PostgreSQL , MySQL/MariaDB ,
HTTP/REST , and MCP gateway adapters. MCP Sources call a configured verification tool over
Streamable HTTP, require its readOnlyHint , and prefer structured tool output. Secrets are
written as ${env:VAR} references and resolve locally, so credentials never
reach the control plane. Use a gateway identity that can invoke only read tools;
MCP annotations are hints, not an authorization boundary.
sources :
stripe-production :
adapter : mcp
endpoint : https://gateway.internal/mcp
headers :
Authorization : Bearer ${env:MCP_GATEWAY_TOKEN}
tool : stripe.get_charge
arguments :
charge_id : " {value} "
select : data.charge
timeout : 10
require_read_only : true
This keeps MCP connection details on the Source. Rules reference the Source and
only describe how to evaluate the returned record; they do not duplicate the
gateway endpoint, tool name, or argument mapping.
PostgreSQL Sources keep connection details on the Source and queries on rules.
Use a dedicated role with SELECT -only permissions. ActionRail also forces each
verification transaction into PostgreSQL read-only mode and applies connection
and statement timeouts. Queries use Psycopg named parameters such as
%(value)s , %(amount)s , and %(customer_id)s :
sources :
billing-production :
adapter : postgres
host : postgres.internal
port : 5432
dbname : billing
user : actionrail_reader
password : ${env:POSTGRES_PASSWORD}
sslmode : verify-full
sslrootcert : /etc/ssl/certs/postgres-ca.pem
connect_timeout : 5
statement_timeout_ms : 5000
require_read_only : true
The SDK depends on the Psycopg interface, leaving the libpq implementation to
the host application. For a self-contained local install, add
psycopg[binary]>=3.2,<4 ; production environments can use the system-linked
psycopg[c] build instead.
MySQL and MariaDB Sources use the same named query parameters and enforce
START TRANSACTION READ ONLY for every check. PyMySQL is included with the SDK:
sources :
billing-mysql :
adapter : mysql
host : mysql.internal
port : 3306
database : billing
user : actionrail_reader
password : ${env:MYSQL_PASSWORD}
ssl_mode : verify-identity
ssl_ca : /etc/ssl/certs/mysql-ca.pem
connect_timeout : 5
query_timeout : 5
require_read_only : true
Control plane + Console
An optional Django + DRF control plane and React Console let you register
sources, author grounding rules per argument, and watch a live Activity
feed of every allow / hold / block the runtime made. The SDK reports metadata
only — tool names, outcomes, value-free reasons, and redacted previews. Detailed
grounding results and rendered previews are exposed only on the local Decision ;
model-facing tool results receive their own fully redacted representation. Preview
placehol

[truncated]

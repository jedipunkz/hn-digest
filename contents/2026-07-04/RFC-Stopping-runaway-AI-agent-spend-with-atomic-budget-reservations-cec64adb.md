---
source: "https://github.com/iamapsrajput/agent-budget-protocol/blob/main/RFC.md"
hn_url: "https://news.ycombinator.com/item?id=48789750"
title: "RFC: Stopping runaway AI agent spend with atomic budget reservations"
article_title: "agent-budget-protocol/RFC.md at main · iamapsrajput/agent-budget-protocol · GitHub"
author: "iamapsrajput"
captured_at: "2026-07-04T22:47:42Z"
capture_tool: "hn-digest"
hn_id: 48789750
score: 1
comments: 0
posted_at: "2026-07-04T22:42:00Z"
tags:
  - hacker-news
  - translated
---

# RFC: Stopping runaway AI agent spend with atomic budget reservations

- HN: [48789750](https://news.ycombinator.com/item?id=48789750)
- Source: [github.com](https://github.com/iamapsrajput/agent-budget-protocol/blob/main/RFC.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T22:42:00Z

## Translation

タイトル: RFC: アトミックな予算予約で AI エージェントの暴走を阻止する
記事のタイトル: メインのagent-budget-protocol/RFC.md · iamapsrajput/agent-budget-protocol · GitHub
説明: GitHub でアカウントを作成して、iamapsrajput/agent-budget-protocol の開発に貢献します。

記事本文:
メインのagent-budget-protocol/RFC.md · iamapsrajput/agent-budget-protocol · GitHub
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
イマプスラジプット
/
エージェント予算プロトコル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする

より多くのファイル アクションを責める より多くのファイル アクションを責める 最新のコミット
履歴 履歴 404 行 (336 loc) · 19.2 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション RFC: AI エージェント実行のためのリアルタイム予算決定計画
ステータス: ドラフト v3、フィードバック用 · 著者: Ajay Rajput · 日付: 7 月
2026 年の対象者: 本番環境で LLM ゲートウェイを実行しているプラットフォーム/インフラ エンジニア
AI エージェントはチャットとは異なりトークンを消費しません。エージェントはループを実行します: 観察、
考え、行動し、繰り返し、それぞれの反復で蓄積されたコンテキストが再送信されます。段階的に
ファイル読み取りを伴う実行が 20 回あると、1 回の呼び出しで入力トークンが 50K を超える可能性があります。報告済み
過去 1 年間の事例には、開発者が 1 回の API 料金で 4,200 ドルに達したことが含まれています。
週末の自律的なリファクタリングと 35 人のエンジニア チームが 87,000 ドルを獲得
毎月の請求書。 30 チームの 1 つの監査で、p10 と p90 の間に 20 倍の広がりがあることが判明しました
同じツールに対する開発者ごとのコスト。これらの数字は業界から得られたものです
主要なインシデントレポートではなく、そのメカニズムが説明されています。
成長するコンテキストを再送信する無制限のループは、構造的で再現可能です。
現在、これを制御することが困難になっている 3 つのギャップがあります。
予算が間違った単位に割り当てられています。既存のゲートウェイ予算は API に関連付けられます
日または月単位で測定される会計期間にわたるキー、ユーザー、またはチーム。の
エージェントのダメージ単位は、実行: 1 つの自律セッションを必要とするものです。
上限はドル単位であり、1 時間で使い果たせる月ごとの割り当てではありません。主流は存在しない
ゲートウェイは実行ごとの上限を強制します。
既存のゲートウェイでは、予算の執行は暗黙的に行われ、脆弱です。最近の
LiteLLM における予算執行の回帰 (例:
#26672 、
#27381 、
#27480 、新しい予算付き
問題は引き続き発生しています) は、根本的な設計上の問題を示しています。
強制は分散されたコールバックとして実装されます。

h 明示的な許可がない
ステップはテストするのが難しく、静かに中断するのは簡単です。別途欠品機種
価格メタデータは無料として扱われ、すべての予算チェックが回避され、
チームレベルの強制は企業のペイウォールの背後にあります。レッスンはそうではありません
「ゲートウェイ X にはバグがあります」;それは、予算権限が明確であるべきであるということです。
明示された保証を備えたテスト可能な意思決定ポイント。
執行は盲目であるため、エージェントは適応できません。予算チェックに失敗した場合
今日、リクエストは不透明なエラーで終了します。エージェントはそれがそうであったことを決して知りません
限界に近づいているので、コストを意識する人間のようなことはできません。
日常的な手順、狭いコンテキスト、またはまとめのために、より安価なモデルにシフトダウンします。
フィードバック チャネルのない可視性は請求書にショックをもたらします。何もせずにブロックする
壊れたランが発生します。どちらもエージェントの動作を変更しません。
AI エージェントの支出管理には、リアルタイムの予算決定プレーンが必要です。この RFC
推定支出をアトミックに予約する実行範囲の予算権限を定義します
プロバイダーの呼び出し前、呼び出し後の実際の使用状況との調整、フェイルクローズ
未知の価格に対応し、機械可読な予算状態を公開してエージェントが適応できるようにする
彼らが走り疲れる前に。
実行ごとの予算上限は、プロバイダーの呼び出し前に適用されます。
同時実行のもとで正確性が保証されます。
機械可読な予算状態プロトコル (応答ヘッダーと RFC 9457)
問題の詳細エラー）により、エージェントは失敗するのではなく、実行中に適応できます。
成功したプロバイダーの応答を変更することなく、ブラインドで実行します。
実行ごとのコスト帰属はユーザー、機能、チームにロールアップされます。
プロバイダー側の課金タグに応じて異なります。
フェールクローズ価格設定: 価格が不明なモデルは、
明示的なテナントオーバーライドが存在します。
モデルゲートウェイやプロバイダーの抽象化ではありません。これは予算上の決定です
ゲートウェイフックとして埋め込むことができるプレーン

、サイドカー、または SDK ミドルウェア。
成功したプロバイダーの応答は変更されずに通過します。
エージェント効率化ツール (コンテキストの圧縮、キャッシュ、サブエージェント) ではありません。の
フレームワークとラボがそのレイヤーを所有します。
事後コストダッシュボードだけではありません。帰属は、作成するためにここに存在します
施行は信頼できる。
実行: 1 つのエージェント セッション。 4.2 のアイデンティティ ルールを参照してください。
Ceiling : スコープ (実行、ユーザー、チーム、キー、機能) に関連付けられた USD 制限
タグ）。
予算決定: 中心的なプリミティブ。通話前の認証結果
権限によって生成される: allowed 、 downgrade 、 Advisory_warn 、または block
(予約は、許可/ダウングレードを裏付ける内部アクションであり、
クライアント側の決定値）。すべての決定には ID があり、その決定とともに記録されます。
インプット（範囲、推定値、有効産出量の上限、価格表のバージョン）。
予約 : 1 つ以上のスコープに対する推定コストのアトミックな保持。
転送前に作成され、転送後にコミットまたはリリースされました。
レジャー (スコープごと): commit_usd 、reserved_usd 、
available_usd = limit_usd − commit_usd −reserved_usd 、および予約
レコード (reservation_id 、expires_at 、price_table_version )。
見積り : 価格表とトークン数からコール前のコストを予測します。
ハードゲート モードでは、予約基準は最悪の場合、実際の入力トークンになります。
プラス、出力価格でのEffective_max_output_tokens (4.4を参照)。
クライアントが提供する実行 ID は便利ですが、信頼できません。ルール:
X-Run-Id は認証された呼び出し元からのみ受け入れられ、バインドされます
サーバー側から認証されたキー/ユーザー/チームに送信します。ランIDは付与できません
それを作成したプリンシパルとは異なるプリンシパル。
実行 ID がない場合、権限はサーバー側の実行 ID を発行し、それを返します。
応答ヘッダー。
すべてのレジャー書き込みは完全なタプルをバインドします。
run_id + user_id + key_id + Team_id + feature_id 。
カーディナリティ制御が適用されます (4 を参照)

.9): プリンシパルごとの最大アクティブ実行数、実行 TTL。
4.3 意思決定の流れ（予約→コミット→返金）
リクエスト → スコープの解決 (実行、ユーザー、チーム、キー、機能)
→Effective_max_output_tokensを計算する(4.4)
→ 要求されたモデルのコストを見積もる (hard_gate での最悪の場合のベース)
→ ATOMIC: 該当するすべてのスコープに対して見積もりを予約するか、失敗します。
§─ すべてのスコープに適合 → 転送リクエスト
§─ ブロックされ、有効な代替 → ダウングレード (ポリシー制御) または代替でブロック
└─ ブロック、代替手段なし → 問題詳細エラーでブロック
→ プロバイダーの成功時: 実際のコストをコミットし、未使用の予約を解放します。
→ プロバイダ障害時: リザーブを解放
→ 結果が見つからない場合: 予約は TTL 後に期限切れになり、非同期的に調整されます。
→ すべての応答に Budget-State ヘッダーを添付します
複数のスコープにわたる予約は、単一のアトミック トランザクションです (1 つの Redis Lua
スクリプトまたは 1 つの SQL トランザクション): すべてのスコープが予約されるか、何も予約されません。シーケンシャル
スコープごとのロックは明示的に拒否されます (部分的な予約、デッドロック)。あ
リクエストは実行上限に適合しても、ユーザーの上限によってブロックされる場合があります。の
決定はブロッキング スコープを報告しますが、トランザクションはすべてのスコープに触れます。
ダウングレードセマンティクス。ダウングレードが選択された場合、当局は保留します
最初に要求されたモデルではなく、選択された代替モデルに対して。
自動ダウングレードは、代替が機能を満たしている場合にのみ許可されます
契約 (4.8) とテナント ポリシーにより、リクエスト クラスのダウングレードが許可されます。
予約ステートマシン。冪等性は明示的な状態に対して定義されます。
予約済み → 転送 → コミット済み
予約→解放
予約済み → 期限切れ → 調整済み
すべての遷移はべき等です。すでにコミットされた予約のコミットは、
無視されました。コミット後のリリースは無視されます。同じものを運ぶ再試行
冪等性キーのレット

新しい決定を下すのではなく、既存の決定を尊重する
予約。
ハードゲートの保証は、 max_output_tokens を誰が制御するかによって異なります。の
権限の推定は、Effective_max_output_tokens に対して行うものであり、盲目的に行うことはありません。
クライアントが指定した値:
有効な最大出力トークン =
分(
client_requested_max_output_tokens、
tenant_policy_max_output_tokens、
モデルコンテキスト残りの出力制限、
Budget_derived_max_output_tokens # オプトインのみ、以下を参照
)
テナントがオプトインしない限り、budget_derived_max_output_tokens は ∞ として扱われます
クランプが有効になっています。 min() はデフォルトでは決して予算を固定しません。
クライアントが max_output_tokens を省略した場合、強制モードはテナントのデフォルトを適用します
またはリクエストを拒否します。クライアントがテナント ポリシーを超える値を要求した場合、
当局は政策に従ってクランプまたは拒否します。
予算に基づくクランプはオプトインのテナント ポリシーであり、デフォルトではありません。縮小する
残りの予算に合わせて世代を変えることは、会計を装った行動の変化です。
切り詰められた差分や書きかけの JSON は、クリーンなブロックよりも悪くなる可能性があります。いつ
有効化され適用されている場合、クランプが表示される必要があります ( X-Budget-Output-Clamped: true 、
決定レコードと問題本文の両方の値を加えます)。
テナント/範囲ごとに、「見積もりはどの程度間違っている可能性がありますか?」に回答します。明示的に:
モードと決定は異なる次元です: モードは構成であり、決定は
リクエストごとの結果 (4.7 で設定されたヘッダーを参照)。推奨される導入パス
段階的に、勧告→ダウングレード許可→ブロック。 Teams の監視に関するアドバイス
本番稼働に影響を与える可能性のあるものを有効にする前に、数値を確認してください。
4.6 台帳実装の不変条件
通貨の精度: すべての元帳金額は整数のマイクロ米ドルとして保存されます。 API
応答は 10 進数の文字列をレンダリングします。浮動小数点演算は禁止されています
予約、コミット、および調整パス。
Redis クラスター: すべてのレジャーキー

1 つの予約スクリプトによって操作されるには、
ハッシュ タグ (例: {tenant_id}:budget:scope:...) を共有することで、マルチスコープの Lua
トランザクションは単一スロットのままです。
SQLite フォールバックはシングルノード/開発モードのみであり、マルチインスタンスではありません。
生産台帳。
価格のバージョン管理: すべての意思決定は、プロバイダー、モデル、入力価格、
出力価格、キャッシュ読み取りおよびキャッシュ書き込み価格、通貨、および
価格_テーブル_バージョン 。テナント固有の価格の上書きがサポートされています。不明
明示的なテナントのオーバーライドが存在しない限り、価格はルーティングできません。
成功したプロバイダーの応答は、デフォルトでは本文が変更されません。権威
OpenAI 互換の応答コントラクト、SDK、ストリーミング クライアント、または
評価用ハーネス。ヘッダーは決定と最も厳密な適用範囲を明らかにします。
完全なマルチスコープ状態は、ログ、監査エクスポート、またはルックアップを通じて入手できます。
エンドポイント:
/budget/決定/{決定_id}を取得します
意思決定記録は設定可能な期間 (デフォルトは 30 日間) 保持されるため、
lookup promise には意味があります。オプションのエンベロープ モードにはフルが含まれる場合があります。
明示的にオプトインするクライアントの応答本文内の予算状態。
X 予算決定: 許可 |ダウングレード |アドバイザリー_警告 |ブロック
X-Budget-Decision-Id: bdgdec_01J...
X-Budget-Registration-Id: rsv_01J... (予約時)
X 予算の執行

[切り捨てられた]

## Original Extract

Contribute to iamapsrajput/agent-budget-protocol development by creating an account on GitHub.

agent-budget-protocol/RFC.md at main · iamapsrajput/agent-budget-protocol · GitHub
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
iamapsrajput
/
agent-budget-protocol
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 404 lines (336 loc) · 19.2 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions RFC: A Real-Time Budget Decision Plane for AI Agent Runs
Status: Draft v3, for feedback · Author: Ajay Rajput · Date: July
2026 Audience: Platform/infra engineers running LLM gateways in production
AI agents don't consume tokens the way chat does. An agent runs a loop: observe,
think, act, repeat, and each iteration resends the accumulated context. By step
20 of a run with file reads, a single call can exceed 50K input tokens. Reported
cases from the past year include a developer hitting $4,200 in API fees over one
weekend of autonomous refactoring and a 35-engineer team receiving an $87K
monthly bill; one audit of 30 teams found a 20x spread between p10 and p90
per-developer cost for the same tooling. These figures come from industry
writeups rather than primary incident reports, but the mechanism they describe,
unbounded loops resending growing context, is structural and reproducible.
Three gaps make this hard to control today:
Budgets attach to the wrong unit. Existing gateway budgets attach to API
keys, users, or teams, over accounting periods measured in days or months. The
damage unit for agents is the run : one autonomous session that needs a
ceiling in dollars, not a monthly quota it can exhaust in an hour. No mainstream
gateway enforces per-run ceilings.
Budget enforcement is implicit and fragile in incumbent gateways. Recent
budget-enforcement regressions in LiteLLM (e.g.
#26672 ,
#27381 ,
#27480 , with new budget
issues continuing to appear) illustrate the underlying design problem:
enforcement implemented as scattered callbacks with no explicit authorization
step is hard to test and easy to silently break. Separately, models with missing
price metadata have been treated as free, bypassing all budget checks, and
team-level enforcement sits behind an enterprise paywall. The lesson is not
"gateway X is buggy"; it is that budget authority should be an explicit,
testable decision point with stated guarantees.
Enforcement is blind, so agents can't adapt. When a budget check fails
today, the request dies with an opaque error. The agent never learns it was
approaching a limit, so it can't do what a cost-conscious human does:
downshift to a cheaper model for routine steps, narrow context, or wrap up.
Visibility without a feedback channel produces bill shock; blocking without one
produces broken runs. Neither changes agent behavior.
AI agent spend control needs a real-time budget decision plane . This RFC
defines a run-scoped budget authority that atomically reserves estimated spend
before provider calls, reconciles against actual usage after calls, fails closed
on unknown prices, and exposes machine-readable budget state so agents can adapt
before they exhaust the run.
Per-run budget ceilings enforced before the provider call, with stated
correctness guarantees under concurrency.
A machine-readable budget-state protocol (response headers plus RFC 9457
problem-detail errors) that lets agents adapt mid-run instead of failing
blind, without mutating successful provider responses.
Per-run cost attribution rolling up to user, feature, and team, without
depending on provider-side billing tags.
Fail-closed pricing: a model with no known price is unroutable unless an
explicit tenant override exists.
Not a model gateway or provider abstraction. This is a budget-decision
plane that can be embedded as a gateway hook, sidecar, or SDK middleware.
Successful provider responses pass through unmodified.
Not agent-efficiency tooling (context compaction, caching, sub-agents). The
frameworks and labs own that layer.
Not post-hoc cost dashboards alone. Attribution exists here to make
enforcement trustworthy.
Run : one agent session. See identity rules in 4.2.
Ceiling : a USD limit attached to a scope (run, user, team, key, feature
tag).
Budget Decision : the central primitive. The pre-call authorization result
produced by the authority: allow , downgrade , advisory_warn , or block
(reservation is the internal action backing allow / downgrade , not a
client-facing decision value). Every decision has an ID and is logged with its
inputs (scopes, estimates, effective output cap, price table version).
Reservation : an atomic hold of estimated cost against one or more scopes,
made before forwarding, committed or released after.
Ledger (per scope): committed_usd , reserved_usd ,
available_usd = limit_usd − committed_usd − reserved_usd , plus reservation
records ( reservation_id , expires_at , price_table_version ).
Estimate : pre-call cost projection from the price table and token counts.
In hard-gate mode the reservation basis is worst-case: actual input tokens
plus effective_max_output_tokens at output price (see 4.4).
Client-supplied run IDs are convenient and untrustworthy. Rules:
X-Run-Id is accepted only from authenticated callers and is bound
server-side to the authenticated key/user/team. A run ID cannot be attached to
a different principal than the one that created it.
Absent a run ID, the authority issues a server-side run ID and returns it in
response headers.
All ledger writes bind the full tuple:
run_id + user_id + key_id + team_id + feature_id .
Cardinality controls apply (see 4.9): max active runs per principal, run TTL.
4.3 Decision flow (reserve → commit → refund)
request → resolve scopes (run, user, team, key, feature)
→ compute effective_max_output_tokens (4.4)
→ estimate cost of requested model (worst-case basis in hard_gate)
→ ATOMIC: reserve estimate against ALL applicable scopes, or fail
├─ all scopes fit → forward request
├─ blocked, valid alternative → downgrade (policy-controlled) or block with alternatives
└─ blocked, no alternative → block with problem-detail error
→ on provider success: commit actual cost, release unused reserve
→ on provider failure: release reserve
→ on missing result: reservation expires after TTL, reconciled asynchronously
→ attach budget-state headers to every response
Reservation across multiple scopes is a single atomic transaction (one Redis Lua
script or one SQL transaction): all scopes reserve or none do. Sequential
per-scope locking is explicitly rejected (partial reservations, deadlocks). A
request may fit the run ceiling and still be blocked by the user ceiling; the
decision reports the blocking scope , but the transaction touches all scopes.
Downgrade semantics. If downgrade is selected, the authority reserves
against the selected alternative model, not the originally requested model.
Auto-downgrade is allowed only when the alternative satisfies capability
contracts (4.8) and tenant policy permits downgrade for the request class.
Reservation state machine. Idempotency is defined over explicit states:
reserved → forwarded → committed
reserved → released
reserved → expired → reconciled
All transitions are idempotent. A commit for an already-committed reservation is
ignored. A release after commit is ignored. A retry carrying the same
idempotency key returns the existing decision rather than creating a new
reservation.
The hard-gate guarantee depends on who controls max_output_tokens . The
authority estimates against effective_max_output_tokens , never blindly against
the client-supplied value:
effective_max_output_tokens =
min(
client_requested_max_output_tokens,
tenant_policy_max_output_tokens,
model_context_remaining_output_limit,
budget_derived_max_output_tokens # opt-in only, see below
)
budget_derived_max_output_tokens is treated as ∞ unless tenant opt-in
clamping is enabled; the min() never budget-clamps by default.
If the client omits max_output_tokens , enforce mode applies a tenant default
or rejects the request. If the client requests a value above tenant policy, the
authority clamps or rejects according to policy.
Budget-derived clamping is opt-in tenant policy, never default. Shrinking a
generation to fit remaining budget is a behavior change disguised as accounting:
a truncated diff or half-written JSON can be worse than a clean block. When
enabled and applied, clamping must be visible ( X-Budget-Output-Clamped: true ,
plus both values in the decision record and problem body).
Per tenant/scope, answering "how wrong can the estimate be?" explicitly:
Mode and decision are different dimensions: mode is configuration, decision is
the per-request outcome (see header set in 4.7). The recommended adoption path
is staged: advisory → downgrade-permitted → blocking. Teams watch advisory
numbers before enabling anything that can touch a production run.
4.6 Ledger implementation invariants
Money precision: all ledger amounts are stored as integer micro-USD. API
responses render decimal strings. Floating-point arithmetic is forbidden in
reservation, commit, and reconciliation paths.
Redis Cluster: all ledger keys touched by one reservation script must
share a hash tag (e.g. {tenant_id}:budget:scope:... ) so the multi-scope Lua
transaction stays single-slot.
SQLite fallback is single-node/dev mode only and is not a multi-instance
production ledger.
Price versioning: every decision records provider, model, input price,
output price, cache-read and cache-write prices, currency, and
price_table_version . Tenant-specific price overrides are supported. Unknown
price is unroutable unless an explicit tenant override exists.
Successful provider responses are not body-mutated by default. The authority
must not break OpenAI-compatible response contracts, SDKs, streaming clients, or
eval harnesses. Headers expose the decision and the tightest applicable scope;
full multi-scope state is available through logs, audit export, or a lookup
endpoint:
GET /budget/decisions/{decision_id}
Decision records are retained for a configurable window (default 30 days) so the
lookup promise is meaningful. An optional envelope mode may include full
budget state in response bodies for clients that explicitly opt in.
X-Budget-Decision: allow | downgrade | advisory_warn | block
X-Budget-Decision-Id: bdgdec_01J...
X-Budget-Reservation-Id: rsv_01J... (when a reservation was made)
X-Budget-Enforceme

[truncated]

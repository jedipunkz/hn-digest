---
source: "https://github.com/Constellation-Labs/gate-oc-audit"
hn_url: "https://news.ycombinator.com/item?id=48548225"
title: "Show HN: Tamper-evident audit trail for AI coding agent activity"
article_title: "GitHub - Constellation-Labs/gate-oc-audit: Tamper-evident audit trail for AI coding agent activity. · GitHub"
author: "gclaramunt"
captured_at: "2026-06-16T01:09:21Z"
capture_tool: "hn-digest"
hn_id: 48548225
score: 3
comments: 0
posted_at: "2026-06-15T23:01:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tamper-evident audit trail for AI coding agent activity

- HN: [48548225](https://news.ycombinator.com/item?id=48548225)
- Source: [github.com](https://github.com/Constellation-Labs/gate-oc-audit)
- Score: 3
- Comments: 0
- Posted: 2026-06-15T23:01:09Z

## Translation

タイトル: HN の表示: AI コーディング エージェント アクティビティの改ざん防止監査証跡
記事のタイトル: GitHub - Constellation-Labs/gate-oc-audit: AI コーディング エージェント アクティビティの改ざん防止監査証跡。 · GitHub
説明: AI コーディング エージェント アクティビティの改ざん防止監査証跡。 - Constellation-Labs/gate-oc-audit
HN テキスト: ここ数か月間取り組んできたものをリリースしました。Openclaw プラグインは、すべてのセッション、ツールの呼び出し、および SHA-256 ハッシュ チェーンの整合性を備えたローカル SQLite データベースへのプロンプト交換を記録するため、事後にイベントが変更または削除されていないことを確認できます。

記事本文:
GitHub - Constellation-Labs/gate-oc-audit: AI コーディング エージェント アクティビティの改ざん明示監査証跡。 · GitHub
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
コンステレーション・ラボ
/
ゲート-OC-監査
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

オプション
コード
Constellation-Labs/gate-oc-audit
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
84 コミット 84 コミット .github/ workflows .github/ workflows docs docs schemas schemas skill/ openclaw-audit skill/ openclaw-audit src src test test .gitignore .gitignore ライセンス ライセンス README.md README.md openclaw.plugin.json openclaw.plugin.json package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
@constellation-network/gate-oc-audit
AI コーディング エージェント アクティビティの改ざん防止監査証跡。すべてのセッション、ツールの呼び出し、ローカル SQLite データベースへのプロンプト交換を SHA-256 ハッシュ チェーンの整合性で記録するため、事後にイベントが変更または削除されていないことを確認できます。
# npmからインストール
openclaw プラグインは @constellation-network/gate-oc-audit をインストールします
# インタラクティブな設定ツールを実行する
オープンクロー監査のセットアップ
これがあなたに何をもたらすか
仕事
CLI
SPA (ゲートウェイ UI 内)
得られるもの
プラグインが正常であることを確認する
オープンクロー監査ステータス
#/status (デフォルトルート)
1 画面のスナップショット: ストレージ、整合性、アンカー、ファイル監視、インベントリ、最後のセキュリティ スキャン
今日/今週のアクティビティを見る
openclaw 監査レポートは毎日 / … 毎週
#/レポート/毎日 、 #/レポート/毎週
アクティビティ、トップツール、LLM 支出、アウトバウンドメッセージング、異常、整合性フッター
cron ごとのロールアップ
openclaw 監査レポート cron <ジョブ ID>
#/reports/cron?jobId=…
実行ごとに 1 行: 開始/終了、ステータス、ツール/LLM/アウトバウンド数
会話ごとのロールアップ
openclaw 監査レポート セッション <id>
#/レポート/セッション/<id>
タイムライン、ツール、LLM コスト、アウトバウンド、整合性
表面の異常
openclaw 監査異常 -- 24 時間以降
#/異常
改ざん、重複送信、d

初期スパイク、インストール、初めて見たツール
LLM 支出を追跡する
openclaw 監査支出 --モデル別 --7d 以降
#/支出
プロバイダー/モデル/日/セッションごとにグループ化されたトークンの使用量とコストUSD
インストールされているプラグイン/スキル/ツール/cron を参照する
openclaw 監査インベントリ [種類]
#/在庫
概要カウンター + 種類ごとのテーブル (cron 行は cron ごとのロールアップにリンク)
出来事が起こったことを証明する
openclaw 監査 smtproof <ハッシュ> then … smt verify
#/smt-tools
木の根とDEアンカーチェックポイントに対する包含証明
独立した第三者による証明
デジタル証拠アンカーリングを構成する (下記を参照)
—
Constellation Digital Evidence ネットワークに固定された SMT のルーツ
ファイル変更に関するアラート
fileWatchPatterns + notificationWebhook を構成する
—
監視パスが変更された場合の Slack/Discord/Webhook ping
チャンネルの日次/週次ダイジェスト
レポートWebhookの構成
—
監査レポートと同じ予測を使用した Slack 互換のペイロード
コンプライアンスのために証跡をエクスポートする
openclaw 監査エクスポート csv --from … --to …
#/events → ダウンロードボタン
行ごとにアンカー参照を含むストリーミング NDJSON または CSV
改ざんのために再スキャンする
オープンクロー監査検証
#/検証
完全な SMT リプレイ + DE チェックポイントの整合性チェック。クリーンな場合は 0 を終了します
どこから始めればよいかわからない場合は、次のコマンドを実行した後に openclaw Audit status を実行してください。
[切り捨てられた]
すべての SPA ルートは、 /plugins/audit/api/ の下にある HTTP エンドポイントによってサポートされます。 SPA ビュー ( /status 、 /anomalies 、 /spend 、 /inventory 、 /report/session/:id 、 /smt/proof 、 /smt/verify-proof 、 /smt/chain ) とともに導入されたエンドポイントは、 /api/report と同じループバック ポリシーに従います。allowExportOnNonLoopback: true が設定されていない限り、403 外部ループバックになります。ワイヤー JSON は、一致する CLI --json 出力とバイト同一であるため、ダッシュボードを同じスキーマに対して固定できます。
ループバックの外側から UI に到達します。デフォルトでは、ルートは次のように登録されます。

auth: "plugin" (検証なし) であり、安全のためにループバック バインドを利用します。共有ネットワーク上で UI/API を公開するには、requireGatewayAuth: true を設定して、ルートが auth: "gateway" で登録され、openclaw ゲートウェイがすべてのリクエストを認証するようにします。
openclaw 構成セット plugins.entries.gate-oc-audit.config.requireGatewayAuth true
{ "config" : { "requireGatewayAuth" : true } }
その後、ゲートウェイがすべての呼び出し元を認証するため、requireGatewayAuth にはループバック ゲート (ステータス、レポート、異常、支出、インベントリ、SMT ツール、エクスポート、検証) がすべてオフループバックで正常に機能し、それ以上のオプトインは必要ありません。これは外部アクセスに推奨される単一ノブです。 allowExportOnNonLoopback /allowVerifyOnNonLoopback フラグは、ゲートウェイ認証なしでこれらのルートをオフループバック (独自のリバース プロキシの背後など) で公開するという狭いケース用に存在します。requireGatewayAuth が設定されている場合はオフのままにしておきます。
openclaw プラグインは @constellation-network/gate-oc-audit をインストールします
ピア依存関係として openclaw >= 2026.4.24、および Node.js ≥ 22.13 (組み込みの node:sqlite モジュールを使用) が必要です。
それでおしまい。プラグインは、エージェントの実行時に監査イベントの記録を自動的に開始します。
セットアップ ウィザードを使用して設定する (推奨)
オープンクロー監査のセットアップ
プラグインに必要なすべての手順を実行する対話型ウィザード:
Gate-oc-audit を plugins.allow に追加します (プラグインを信頼します)。
hooks.allowConversationAccess を有効にして、prompt.input /prompt.response /agent.end イベントをキャプチャできるようにします。
オプションで、デジタル証拠 API キー + 組織/テナント UUID を収集し、アンカーしきい値 (イベント数、タイマー間隔、ティックごとの最小イベント) を調整します。
最後に openclaw 監査ステータスを実行して、証跡がライブであることを確認できます。
--yes を渡すと、アンカー調整のデフォルトが受け入れられます (DE の ID プロンプトは引き続き尋ねられます)。一度ウィザード

完了したら、「インストール後のクイックチェック」に進みます。
手動構成リファレンス (openclaw ≥ 2026.4.24)
すべての機能を使用するには、2 つのオペレータ ポリシー オプトインが必要です。どちらもオープンクローがオペレーターに強制する決定です。プラグインはどちらも自己許可できません。上記のセットアップ ウィザードは両方を自動的に作成します。以下のセクションでは、手動または構成管理を通じて構成をプロビジョニングするオペレーター向けに、同等の openclaw 構成セット呼び出しと JSON スニペットを文書化します。
Gate-oc-audit を plugins.allow に追加します。 plugins.allow が空の場合でも、openclaw は検出されたプラグインを自動ロードしますが、起動するたびに警告をログに記録します。
[プラグイン] plugins.allow が空です。検出されたバンドルされていないプラグインは自動ロードされる可能性があります…
明示的な許可リストを設定すると、警告が表示されなくなり、リストされた ID への読み込みがロックされます。
openclaw 構成セット plugins.allow ' ["gate-oc-audit"] '
または、構成 JSON 内で直接次のように指定します。
{
「プラグイン」: {
"allow" : [ "gate-oc-audit " ]
}
}
plugins.allow に他の信頼できるプラグインがすでにある場合は、既存の配列を置き換えるのではなく、「gate-oc-audit」を追加します。
バンドルされていないプラグインは、 llm_input 、 llm_output 、 before_agent_finalize 、および Agent_end フックから生の会話コンテンツを受信することを明示的にオプトインする必要があります。このオプトインがないと、openclaw はこれらのフックの登録とログをブロックします。
[プラグイン] バンドルされていないプラグインは plugins.entries.gate-oc-audit.hooks.allowConversationAccess=true を設定する必要があるため、型付きフック「llm_input」がブロックされました…
これが設定されるまで、監査プラグインの最も重要なイベント タイプのうち 3 つ (prompt.input 、prompt.response 、agent.end ) が監査証跡から欠落します。
openclaw 構成セット plugins.entries.gate-oc-audit.hooks.allowConversationAccess true
または、構成 JSON 内で直接次のように指定します。
{
「プラグイン」: {
「エントリ」: {
"ゲート-oc-監査": {
"フック" : { "allowConversationAccess" : tr

うえ}
}
}
}
}
また、プラグインは、先行する llm_input イベントなしでツール呼び出しが観察された場合、起動時に警告をログに記録します。これは通常、このオプトインが欠落していることを示します。
オープンクロー監査ステータス
これは、プラグインが記録中であることを確認するための唯一の最良のコマンドです。 1 つの画面カバーを印刷します。
ストレージ — DB サイズと上限、イベント数、最も古いイベント、次のプルーン
整合性 — シーケンス ヘッド、SMT ツリー + ルート、最後のチェックポイント、会話フック状態 (有効 / 無効 / 有効 - ただし、24 時間以内にプロンプト入力が表示されない場合はサイレント)
デジタル証拠アンカー — アクティブ/非アクティブ、今日のアンカー、最後のアンカー TX ハッシュ、サーキット ブレーカーの状態
ファイルの監視 - パターンの監視/無視、過去 24 時間の変更
インベントリ — プラグイン / スキル / ツール / cron 数
最新のセキュリティ スキャン — タイムスタンプ + 検出数
単一行の機械可読スナップショットの場合は --json を追加します。会話フックに DISABLED が表示されている場合、オペレーターはallowConversationAccess (上記を参照) を設定しておらず、prompt.input /prompt.response /agent.end イベントが証跡から欠落しています。
すべてのコマンドは、SQLite WAL 経由で監査 DB を読み取り専用で開くため、ロックの競合なしに実行中のゲートウェイと共存します。
上記のインストール後のクイックチェックを参照してください。 --json を使用して jq にパイプします。
オープンクロー監査リスト
openclaw 監査リスト -- 最新 20
openclaw 監査リスト --type tool.invoked
openclaw 監査リスト --category プロンプト --session < セッション ID >
整合性の検証
最近のイベントの SMT プルーフを検証し、DE チェックポイントの一貫性を確認します。
オープンクロー監査検証
すべての証明とチェックポイントが有効な場合はコード 0 で終了し、検証が失敗した場合はコード 1 で終了します。
openclaw 監査インベントリ # プラグイン/スキル/ツール/cron 全体の概要
openclaw 監査インベントリ プラグイン # 1 種類の詳細
openclaw 監査インベントリ スキル --json
レポート
インライン異常を使用して日次または週次のアクティビティ ダイジェストを生成する

y検出器。この予測は、アクティビティ / Cron スケジュール / トップ ツール / LLM 支出 / アウトバウンド メッセージング / 異常 / 整合性をカバーし、ヒューマン テキスト (デフォルト)、単一行 JSON ( --json )、または自己完結型 HTML ドキュメント ( --html ) としてレンダリングされます。
openclaw 監査レポート日次 # 今日 (UTC)、人間のテキスト
openclaw 監査レポート毎日 --date 2026-05-17 # 特定の UTC 日
openclaw 監査レポート毎日 --tz local # 現地時間の日境界を使用します
openclaw 監査レポート毎週 # 今週 ISO 週 (UTC)
openclaw 監査レポート毎週 --week 2026-W19 # 特定の ISO 週
openclaw 監査レポート毎日 --json # 単一行 JSON
openclaw 監査レポート毎日 --html > report.html # スタンドアロン HTML
openclaw 監査レポート cron < job-id > # cron ごとのロールアップ、実行ごとに 1 行
openclaw 監査レポート cron < ジョブ ID > --last 5 --json
openclaw 監査レポート cron < job-id > --html > cron.html
openclaw 監査レポート セッション < session-id > # 会話ごとのロールアップ
検出器ノブ (CLI と HTTP の両方にキャップ付き):
異常ブロックで発行される異常検出器:
R5a 重複送信 — 同じコンテンツのハッシュが --dup-window-sec 内の同じチャネル + 受信者に送信されます。 DuplicateOutboundTruncated: true は、基になる message.sent スキャンが 100k 行の上限に達し、それを超えると重複が発生することを示します。

[切り捨てられた]

## Original Extract

Tamper-evident audit trail for AI coding agent activity. - Constellation-Labs/gate-oc-audit

We released what I've been working in the last few months: an Openclaw plugin that ecords every session, tool invocation, and prompt exchange into a local SQLite database with SHA-256 hash chain integrity, so you can verify that no events were altered or deleted after the fact.

GitHub - Constellation-Labs/gate-oc-audit: Tamper-evident audit trail for AI coding agent activity. · GitHub
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
Constellation-Labs
/
gate-oc-audit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Constellation-Labs/gate-oc-audit
main Branches Tags Go to file Code Open more actions menu Folders and files
84 Commits 84 Commits .github/ workflows .github/ workflows docs docs schemas schemas skills/ openclaw-audit skills/ openclaw-audit src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md openclaw.plugin.json openclaw.plugin.json package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
@constellation-network/gate-oc-audit
Tamper-evident audit trail for AI coding agent activity. Records every session, tool invocation, and prompt exchange into a local SQLite database with SHA-256 hash chain integrity, so you can verify that no events were altered or deleted after the fact.
# install from npm
openclaw plugins install @constellation-network/gate-oc-audit
# run the interactive configuration tool
openclaw audit setup
What this does for you
Job
CLI
SPA (in the gateway UI)
What you get
Confirm the plugin is healthy
openclaw audit status
#/status (default route)
One-screen snapshot: storage, integrity, anchor, file-watch, inventory, last security scan
See today's / this week's activity
openclaw audit report daily / … weekly
#/reports/daily , #/reports/weekly
Activity, top tools, LLM spend, outbound messaging, anomalies, integrity footer
Per-cron rollup
openclaw audit report cron <job-id>
#/reports/cron?jobId=…
One row per execution: started/ended, status, tool/LLM/outbound counts
Per-conversation rollup
openclaw audit report session <id>
#/reports/session/<id>
Timeline, tools, LLM cost, outbound, integrity
Surface anomalies
openclaw audit anomalies --since 24h
#/anomalies
Tamper, duplicate-outbound, denial spikes, installs, first-seen tools
Track LLM spend
openclaw audit spend --by model --since 7d
#/spend
Token usage and costUsd grouped by provider / model / day / session
Browse installed plugins/skills/tools/crons
openclaw audit inventory [kind]
#/inventory
Summary counters + per-kind tables (cron rows link to per-cron rollup)
Prove an event happened
openclaw audit smt proof <hash> then … smt verify
#/smt-tools
Inclusion proof against tree roots and DE-anchored checkpoints
Independent third-party proof
Configure Digital Evidence anchoring (see below)
—
SMT roots anchored on the Constellation Digital Evidence network
Alert on file changes
Configure fileWatchPatterns + notificationWebhook
—
Slack/Discord/webhook ping when a watched path changes
Daily/weekly digests to a channel
Configure reportWebhook
—
Slack-compatible payload with the same projection as audit report
Export the trail for compliance
openclaw audit export csv --from … --to …
#/events → Download button
Streamed NDJSON or CSV with anchor references per row
Re-scan for tampering
openclaw audit verify
#/verify
Full SMT replay + DE checkpoint consistency check; exit 0 if clean
If you don't know where to start, run openclaw audit status after i
[truncated]
Every SPA route is backed by an HTTP endpoint under /plugins/audit/api/ . The endpoints introduced alongside the SPA views ( /status , /anomalies , /spend , /inventory , /report/session/:id , /smt/proof , /smt/verify-proof , /smt/chain ) follow the same loopback policy as /api/report — 403 outside loopback unless allowExportOnNonLoopback: true is set. The wire JSON is byte-identical to the matching CLI --json output, so dashboards can pin against the same schemas.
Reaching the UI from outside loopback. By default the routes register with auth: "plugin" (no verification) and lean on the loopback bind for safety. To expose the UI/API on a shared network, set requireGatewayAuth: true so the routes register with auth: "gateway" and the openclaw gateway authenticates every request:
openclaw config set plugins.entries.gate-oc-audit.config.requireGatewayAuth true
{ "config" : { "requireGatewayAuth" : true } }
Because the gateway then authenticates every caller, requireGatewayAuth subsumes the loopback gate: status, reports, anomalies, spend, inventory, SMT tools, export, and verify all serve normally off-loopback, no further opt-in needed. This is the recommended single knob for external access. The allowExportOnNonLoopback / allowVerifyOnNonLoopback flags exist for the narrower case of exposing those routes off-loopback without gateway auth (e.g. behind your own reverse proxy) — leave them off when requireGatewayAuth is set.
openclaw plugins install @constellation-network/gate-oc-audit
Requires openclaw >= 2026.4.24 as a peer dependency and Node.js ≥ 22.13 (uses the built-in node:sqlite module).
That's it. The plugin automatically starts recording audit events when your agent runs.
Configure with the setup wizard (recommended)
openclaw audit setup
Interactive wizard that walks through everything the plugin needs:
Adds gate-oc-audit to plugins.allow (trusts the plugin).
Enables hooks.allowConversationAccess so prompt.input / prompt.response / agent.end events are captured.
Optionally collects your Digital Evidence API key + org/tenant UUIDs and tunes the anchoring thresholds (event count, timer interval, minimum events per tick).
Runs openclaw audit status at the end so you can confirm the trail is live.
Pass --yes to accept defaults for the anchoring tuning (DE identity prompts still ask). Once the wizard finishes, skip to Quick check after install .
Manual config reference (openclaw ≥ 2026.4.24)
Two operator-policy opt-ins are required for full functionality. Both are decisions openclaw forces on the operator — the plugin cannot self-grant either. The setup wizard above writes both for you; the sections below document the equivalent openclaw config set calls and JSON snippets for operators who provision config by hand or via configuration management.
Add gate-oc-audit to plugins.allow . When plugins.allow is empty, openclaw still auto-loads discovered plugins but logs a warning on every startup:
[plugins] plugins.allow is empty; discovered non-bundled plugins may auto-load …
Setting an explicit allowlist silences the warning and locks loading down to the listed ids:
openclaw config set plugins.allow ' ["gate-oc-audit"] '
Or directly in the config JSON:
{
"plugins" : {
"allow" : [ " gate-oc-audit " ]
}
}
If you already have other trusted plugins in plugins.allow , append "gate-oc-audit" to the existing array rather than replacing it.
Non-bundled plugins must explicitly opt in to receive raw conversation content from the llm_input , llm_output , before_agent_finalize , and agent_end hooks. Without this opt-in, openclaw blocks those hook registrations and logs:
[plugins] typed hook "llm_input" blocked because non-bundled plugins must set plugins.entries.gate-oc-audit.hooks.allowConversationAccess=true …
Three of the audit plugin's most important event types ( prompt.input , prompt.response , agent.end ) will be missing from the audit trail until this is set.
openclaw config set plugins.entries.gate-oc-audit.hooks.allowConversationAccess true
Or directly in the config JSON:
{
"plugins" : {
"entries" : {
"gate-oc-audit" : {
"hooks" : { "allowConversationAccess" : true }
}
}
}
}
The plugin also logs a warning at startup if a tool call is observed without any preceding llm_input event, which usually indicates this opt-in is missing.
openclaw audit status
This is the single best command to confirm the plugin is recording. It prints one screen covering:
Storage — DB size vs cap, event count, oldest event, next prune
Integrity — sequence head, SMT trees + root, last checkpoint, conversation-hook state ( ENABLED / DISABLED / ENABLED-but-silent if no prompt.input seen in 24h)
Digital Evidence anchor — active / inactive, anchors today, last anchor tx hash, circuit-breaker state
File watching — patterns watched / ignored, changes in last 24h
Inventory — plugins / skills / tools / cron counts
Last security scan — timestamp + finding counts
Add --json for a single-line, machine-readable snapshot. If Conversation hook shows DISABLED , the operator hasn't set allowConversationAccess (see above) and prompt.input / prompt.response / agent.end events are missing from the trail.
All commands open the audit DB read-only via SQLite WAL, so they coexist with the running gateway without lock contention.
See Quick check after install above. Use --json to pipe to jq .
openclaw audit list
openclaw audit list --last 20
openclaw audit list --type tool.invoked
openclaw audit list --category prompt --session < session-id >
Verify integrity
Verify SMT proofs for recent events and check DE checkpoint consistency:
openclaw audit verify
Exits with code 0 if all proofs and checkpoints are valid, 1 if any verification fails.
openclaw audit inventory # summary across plugins/skills/tools/crons
openclaw audit inventory plugins # detail for one kind
openclaw audit inventory skills --json
Report
Generate a daily or weekly activity digest with inline anomaly detectors. The projection covers Activity / Cron schedule / Top tools / LLM spend / Outbound messaging / Anomalies / Integrity, rendered as human text (default), single-line JSON ( --json ), or a self-contained HTML document ( --html ).
openclaw audit report daily # today (UTC), human text
openclaw audit report daily --date 2026-05-17 # specific UTC day
openclaw audit report daily --tz local # use local-time day boundary
openclaw audit report weekly # this ISO week (UTC)
openclaw audit report weekly --week 2026-W19 # specific ISO week
openclaw audit report daily --json # single-line JSON
openclaw audit report daily --html > report.html # standalone HTML
openclaw audit report cron < job-id > # per-cron rollup, one row per execution
openclaw audit report cron < job-id > --last 5 --json
openclaw audit report cron < job-id > --html > cron.html
openclaw audit report session < session-id > # per-conversation rollup
Detector knobs (capped on both CLI and HTTP):
Anomaly detectors emitted in the anomalies block:
R5a duplicate outbound — same content hash sent to the same channel + recipient inside --dup-window-sec . duplicateOutboundTruncated: true indicates the underlying message.sent scan hit its 100k-row cap and a duplicate beyond that point c

[truncated]

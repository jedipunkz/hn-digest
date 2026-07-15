---
source: "https://uptimemonitoring.com/"
hn_url: "https://news.ycombinator.com/item?id=48919840"
title: "Show HN: An uptime monitor with an MCP server, create monitors by asking Claude"
article_title: "UptimeMonitoring.com — API-first uptime monitoring for deploy pipelines, developers, and AI agents"
author: "luciandan"
captured_at: "2026-07-15T13:13:33Z"
capture_tool: "hn-digest"
hn_id: 48919840
score: 2
comments: 0
posted_at: "2026-07-15T12:31:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An uptime monitor with an MCP server, create monitors by asking Claude

- HN: [48919840](https://news.ycombinator.com/item?id=48919840)
- Source: [uptimemonitoring.com](https://uptimemonitoring.com/)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T12:31:34Z

## Translation

タイトル: Show HN: MCP サーバーを使用した稼働時間モニター、Claude に依頼してモニターを作成
記事のタイトル: UptimeMonitoring.com — デプロイ パイプライン、開発者、AI エージェント向けの API ファーストの稼働時間モニタリング
説明: デプロイ パイプライン、開発者、AI エージェントの API ファーストの稼働時間監視。

記事本文:
Uptime Monitoring .com by Monitive クイックスタート ドキュメント 例 料金 ブログ 概要
API キーを取得する
メニュー クイックスタート ドキュメント 例 価格設定 ブログ About
API キーを取得する
By Monitive · 16 年間のモニタリング · 50 モニターは無料 · GitHub ログイン · メールなし · カードなし
デプロイパイプライン、開発者、AI エージェント向けの API ファーストの稼働時間監視。
Curl 呼び出し、MCP プロンプト、または GitHub アクションを使用してモニターを作成します。
クイックスタートを読む
クロード — 私のアプリ MCP · クロード コードカール · REST API GitHub アクション接続 (1 つのコマンド) claude mcp add --transport http uptimemonitoring \
https://api.uptimemonitoring.com/mcp コピー 次に、Claude に尋ねるだけです。 ❯ 「My Site」という名前の https://example.com の HTTP モニターを作成します。 ❯ 私のモニターをリストします。 ❯ モニター 123 の現在のステータスは何ですか?コピー
API キーをお持ちですか? --header「認証: ベアラー umk_live_...」を追加します。
キー経由で認証します。
カール -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "認可: ベアラー umk_live_..." \
-H "コンテンツ タイプ: application/json" \
-d '{"名前":"myapp-healthz","url":"https://myapp.com/healthz",
"type":"http","expected_status":"200"}' 状態をコピーして確認しますcurl https://api.uptimemonitoring.com/api/v1/monitors/123 \
-H "Authorization: Bearer umk_live_..." デプロイ後のゲートのコピー - 名前: デプロイの健全性の確認
使用: uptimemonitoring/assert-healthy@v1
と:
API キー: ${{ Secrets.UPTIMEMONITORING_API_KEY }}
モニター ID: ${{ vars.MONITOR_ID }}
タイムアウト: 120 コピー デプロイ後にモニターが正常でない場合、ワークフローは失敗します。
すでに作業している表面を使用します。
同じバックエンド。同じモニター状態。さまざまな運転方法があります。
スクリプト、アプリ、または CI からモニターを作成および管理します。
カール -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "認可: ベアラー $UPTIMEMONITORING_API_KEY" \
-H "コンテンツ タイプ: アプリケーション

ation/json" \
-d '{
"名前": "API-prod",
"url": "https://api.example.com/healthz",
"タイプ": "http"
}' Claude、ChatGPT、または Cursor にモニターの正常性を作成、検査、およびアサートさせます。
ユーザー:
https://api.example.com/healthz の 60 秒モニターを作成する
そして今健康かどうか教えてください。
ツール:
作成_モニター
結果:
作成されたモニター: api-prod
現在のステータス: 稼働中
最後の証拠: EU · 200 · ttfb 184ms
ワークフローに合わせて構築
発送方法に適した表面をお選びください。
サイトがダウンしているときにデプロイが失敗する GitHub アクションを使用して、トラフィックが切り替わる前に、デプロイが 22 のプローブ位置から到達可能であることを確認します。
Claude エージェントに独自の展開を監視させる UptimeMonitoring MCP サーバーを Claude Code に追加すると、エージェントがモニターを作成し、ステータスを確認し、障害に対応できるようになります。
Cursor 内からの稼働時間監視 Cursor で MCP サーバーとして UptimeMonitoring を設定し、エージェント パネルから監視を管理します。
50 回の無料チェックであらゆる API を監視 単一の POST リクエストで REST API 経由で監視を作成できます。ダッシュボードは必要ありません。
何かが壊れると何が起こるか
最初のブリップで赤を反転しません。
モニターは、API、MCP、または GitHub アクションを通じて作成します。
即時にテストチェックを実行するため、タイプミスや到達不能なターゲットはすぐに失敗します。
負荷を安定させるために、チェックは分単位で分散されます。
状態を反転する前に、インフラストラクチャ クラスの障害が 2 つの追加リージョンで確認されます。
インシデントを証拠とともに保存し、Web フック、ブラウザー プッシュ、RSS、または MCP プルなど、必要な場所に送信します。
インシデントの前に何が起こったかを確認します: DNS、TLS、タイムアウト、5xx、リージョン、タイミングの内訳。
ダッシュボードではなくデバッグ用に設計
{
「id」: 194、
"タイプ": "ダウン"、
「モニターID」: 1287、
"started_at": "2026-04-18T05:38:52Z",
"resolved_at": "2026-04-18T05:45:44Z",
「証拠」: [
{
"地域": "米国

-W"、
"タイムスタンプ": "2026-04-18T05:37:41Z",
「ステータスコード」: 200、
「dns_ms」: 6.8、
"tls_ms": 298、
"connect_ms": 149、
"ttfb_ms": 203
}、
{
"地域": "US-W",
"タイムスタンプ": "2026-04-18T05:38:42Z",
"error_class": "DNS"
}
]
}
システム用の Webhook。人間のためにプッシュします。エージェント向けの MCP。
受信トレイにスパムはありません。到達可能性の悩みはありません。チャンネルを選ぶのはあなたです。
Slack、Twilio、Discord、ntfy、GitHub Actions、または独自のルーター用の Webhook。
PUT /api/v1/monitors/1287/webhook
{ "url": "https://hooks.co/inc" } ブラウザ プッシュ ダッシュボードからワンクリックでオプトインします。人間による最速のフォールバック。
通知を有効にする
macOS 上の Chrome · アクティブ RSS / インシデント フィード Slack、リーダー、または自動化ツールから購読します。
https://api.uptimemonitoring.com
/feed/{token}/rss MCP pull モニターの状態と最近のインシデントについてエージェントに直接問い合わせてください。
ツール: list_incidents
フィルター: 過去 24 時間、ステータス: オープン
結果: 1 件のインシデント、myapp-healthz ブラウザーのプッシュと電子メールがおやすみモードによってミュートされました。まさに停止が重要なときに発生しました。 Webhook を Pushover (iOS および Android で DND をバイパス) または ntfy に接続して、何かが壊れた瞬間に電話アラートを送信します。
実際に壊れるものを検証する
HTTP / HTTPS ステータス コード、ボディの一致、タイミングの証拠
DNS、接続、TLS、TTFB、ダウンロード
白紙のページではなく、実際のレシピから始めましょう。
任意の端末からモニターを作成
カール -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "認可: ベアラー $UPTIMEMONITORING_API_KEY" \
-d '{"url":"https://api.example.com/healthz","type":"http"}'
例を開く→
ノード デプロイ スクリプトにモニタリングを追加する
const res = await fetch(
"https://api.uptimemonitoring.com/api/v1/monitors",
{
メソッド: "POST"、
ヘッダー: { 認可: `Bearer ${process.env.UPTIMEMONITORING_API_KEY}` }、
本文: JSON.stringify({ url: "https://api.example.com/healthz", type: "http" })
}
)
例を開く

→
Python スクリプトまたはノートブックから自動化する
インポートリクエスト
リクエスト.post(
"https://api.uptimemonitoring.com/api/v1/monitors",
headers={"認可": f"ベアラー {api_key}"},
json={"url": "https://api.example.com/healthz", "type": "http"}
)
例を開く→
クロード 自然言語モニターの管理
UptimeMonitoring を MCP サーバーとして追加する
https://api.example.com/healthz のモニターを作成する
デプロイ後に正常であることをアサートします
例を開く→
IDE からのカーソルモニター
// .cursor/mcp.json 内:
{
"mcpサーバー": {
"稼働時間監視": {
"url": "https://api.uptimemonitoring.com/mcp",
"ヘッダー": {
"認可": "ベアラー umk_live_..."
}
}
}
}
例を開く→
Webhook レシピ インシデントを任意のシステムにルーティングする
PUT /api/v1/monitors/1287/webhook
{
"url": "https://hooks.slack.com/services/..."
}
例を開く→
無料プラン
実際のプロジェクトには十分です。覚えやすいほど簡単です。
無料利用枠の延長中は最大 100 件（初期ユーザー）
ブラウザプッシュ + RSS インシデントフィード
API キーを取得する
有料プランでは、30 秒間隔、より長い保持期間、および応答時間のしきい値アラートが追加されます。
UptimeMonitoring.com は、Monitive Pro の API ファーストのヘッドレス兄弟です。
Monitive は 16 年間監視インフラストラクチャを運用してきました。
UptimeMonitoring.com は、その運用経験を活用して、別のジョブに適用します。
ゲート、API ワークフロー、AI エージェント ツール、開発者優先のモニタリングをデプロイする必要はありません。
電子メールやダッシュボードのスプロール化。
リージョンを越えた障害の確認
最小限のデータ収集: GitHub ID のみ
API / MCP / GitHub アクション
│
▼
Go ディスパッチャー + 証拠
│
▼
22プローブネットワーク
│
▼
Webhook / プッシュ / RSS / インシデント FAQ
はい。無料プラン: 最大 50 台のモニター。初期ユーザーは、有料プランが最終決定されるまでの延長無料利用枠中に最大 100 まで増やすことができます。
いいえ。サインアップには GitHub OAuth を使用します。メールでお知らせ

■ 将来の価格更新が必要な場合のみオプションです。
Webhook、ブラウザープッシュ、RSS フィード、MCP クエリ。
インフラストラクチャ クラスの障害は、状態が変化する前にリージョン全体で再チェックされます。
いいえ。プライベートおよび予約されたターゲットはブロックされます。
Monitive による UptimeMonitoring.com · 電子メールは不要 · GitHub ID のみ · コードと運用のために構築

## Original Extract

API-first uptime monitoring for deploy pipelines, developers, and AI agents.

Uptime Monitoring .com by Monitive Quickstart Docs Examples Pricing Blog About
Get your API key
Menu Quickstart Docs Examples Pricing Blog About
Get your API key
By Monitive · 16 years in monitoring · 50 monitors free · GitHub login · no email · no card
API-first uptime monitoring for deploy pipelines, developers, and AI agents.
Create monitors with a curl call, an MCP prompt, or a GitHub Action.
Read quickstart
claude — my-app MCP · Claude Code curl · REST API GitHub Action Connect (one command) claude mcp add --transport http uptimemonitoring \
https://api.uptimemonitoring.com/mcp Copy Then just ask Claude ❯ Create an HTTP monitor for https://example.com named "My Site" ❯ List my monitors ❯ What is the current status of monitor 123? Copy
Have an API key? Add --header "Authorization: Bearer umk_live_..."
to authorize via key.
curl -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "Authorization: Bearer umk_live_..." \
-H "Content-Type: application/json" \
-d '{"name":"myapp-healthz","url":"https://myapp.com/healthz",
"type":"http","expected_status":"200"}' Copy Check its state curl https://api.uptimemonitoring.com/api/v1/monitors/123 \
-H "Authorization: Bearer umk_live_..." Copy Post-deploy gate - name: Verify deploy health
uses: uptimemonitoring/assert-healthy@v1
with:
api-key: ${{ secrets.UPTIMEMONITORING_API_KEY }}
monitor-id: ${{ vars.MONITOR_ID }}
timeout: 120 Copy Fails the workflow if the monitor is not healthy after your deploy.
Use the surface you already work in.
Same backend. Same monitor state. Different ways to drive it.
Create and manage monitors from scripts, apps, or CI.
curl -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "Authorization: Bearer $UPTIMEMONITORING_API_KEY" \
-H "Content-Type: application/json" \
-d '{
"name": "api-prod",
"url": "https://api.example.com/healthz",
"type": "http"
}' Let Claude, ChatGPT, or Cursor create, inspect, and assert monitor health for you.
User:
Create a 60-second monitor for https://api.example.com/healthz
and tell me if it is healthy right now.
Tool:
create_monitor
Result:
Monitor created: api-prod
Current status: up
Last evidence: EU · 200 · ttfb 184ms
Built for your workflow
Pick the surface that fits how you ship.
Fail deploys when your site is down Use a GitHub Action to assert your deploy is reachable from 22 probe locations before traffic switches over.
Let your Claude agent monitor its own deployments Add the UptimeMonitoring MCP server to Claude Code so your agent can create monitors, check status, and react to failures.
Uptime monitoring from inside Cursor Configure UptimeMonitoring as an MCP server in Cursor and manage monitors from the agent panel.
Monitor any API with 50 free checks Create monitors via the REST API with a single POST request — no dashboard required.
What happens when something breaks
We do not flip red on the first blip.
You create a monitor through the API, MCP, or GitHub Action.
We run an immediate test check so typos and unreachable targets fail fast.
Checks are spread across the minute to keep load stable.
Infrastructure-class failures are confirmed from two additional regions before we flip state.
We store the incident with evidence and send it where you want it: webhook, browser push, RSS, or MCP pull.
See what happened before the incident: DNS, TLS, timeout, 5xx, region, and timing breakdowns.
Designed for debugging, not dashboards
{
"id": 194,
"type": "down",
"monitor_id": 1287,
"started_at": "2026-04-18T05:38:52Z",
"resolved_at": "2026-04-18T05:45:44Z",
"evidence": [
{
"region": "US-W",
"timestamp": "2026-04-18T05:37:41Z",
"status_code": 200,
"dns_ms": 6.8,
"tls_ms": 298,
"connect_ms": 149,
"ttfb_ms": 203
},
{
"region": "US-W",
"timestamp": "2026-04-18T05:38:42Z",
"error_class": "dns"
}
]
}
Webhooks for systems. Push for humans. MCP for agents.
No inbox spam. No deliverability headaches. You choose the channel.
Webhooks For Slack, Twilio, Discord, ntfy, GitHub Actions, or your own router.
PUT /api/v1/monitors/1287/webhook
{ "url": "https://hooks.co/inc" } Browser push One-click opt-in from the dashboard. Fastest human fallback.
Enable notifications
Chrome on macOS · Active RSS / incident feed Subscribe from Slack, readers, or automation tools.
https://api.uptimemonitoring.com
/feed/{token}/rss MCP pull Ask your agent for monitor state and recent incidents directly.
Tool: list_incidents
Filter: last 24h, status: open
Result: 1 incident, myapp-healthz Browser push and email get muted by Do Not Disturb — exactly when an outage matters. Wire a webhook to Pushover (bypasses DND on iOS and Android) or ntfy for phone alerts the moment something breaks.
Validate the things that actually break
HTTP / HTTPS Status codes, body match, timing evidence
DNS, connect, TLS, TTFB, download
Start from a real recipe, not a blank page.
Create a monitor from any terminal
curl -X POST https://api.uptimemonitoring.com/api/v1/monitors \
-H "Authorization: Bearer $UPTIMEMONITORING_API_KEY" \
-d '{"url":"https://api.example.com/healthz","type":"http"}'
Open example →
Node Add monitoring to a deploy script
const res = await fetch(
"https://api.uptimemonitoring.com/api/v1/monitors",
{
method: "POST",
headers: { Authorization: `Bearer ${process.env.UPTIMEMONITORING_API_KEY}` },
body: JSON.stringify({ url: "https://api.example.com/healthz", type: "http" })
}
)
Open example →
Python Automate from scripts or notebooks
import requests
requests.post(
"https://api.uptimemonitoring.com/api/v1/monitors",
headers={"Authorization": f"Bearer {api_key}"},
json={"url": "https://api.example.com/healthz", "type": "http"}
)
Open example →
Claude Natural language monitor management
Add UptimeMonitoring as an MCP server
Create a monitor for https://api.example.com/healthz
Assert that it is healthy after deploy
Open example →
Cursor Monitor from your IDE
// In .cursor/mcp.json:
{
"mcpServers": {
"uptimemonitoring": {
"url": "https://api.uptimemonitoring.com/mcp",
"headers": {
"Authorization": "Bearer umk_live_..."
}
}
}
}
Open example →
Webhook recipes Route incidents to any system
PUT /api/v1/monitors/1287/webhook
{
"url": "https://hooks.slack.com/services/..."
}
Open example →
Free plan
Enough for real projects. Simple enough to remember.
Up to 100 during extended free tier (early users)
Browser push + RSS incident feed
Get your API key
Paid plans will add 30-second intervals, longer retention, and response-time threshold alerts.
UptimeMonitoring.com is the API-first, headless sibling to Monitive Pro.
Monitive has been operating monitoring infrastructure for 16 years.
UptimeMonitoring.com takes that operational experience and applies it to a different job:
deploy gates, API workflows, AI-agent tooling, and developer-first monitoring without
email or dashboard sprawl.
Cross-region confirmation on failures
Minimal data collection: GitHub identity only
API / MCP / GitHub Action
│
▼
Go dispatcher + evidence
│
▼
22-probe network
│
▼
webhooks / push / RSS / incidents FAQ
Yes. Free plan: up to 50 monitors. Early users can go up to 100 during the extended free tier while paid plans are being finalized.
No. Signup uses GitHub OAuth. Email is optional only if you want future pricing updates.
Webhooks, browser push, RSS feeds, and MCP queries.
Infrastructure-class failures are re-checked across regions before state changes.
No. Private and reserved targets are blocked.
UptimeMonitoring.com by Monitive · No email required · GitHub identity only · Built for code and ops

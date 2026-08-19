---
source: "https://marginalhq.com/"
hn_url: "https://news.ycombinator.com/item?id=49361298"
title: "Show HN: Marginal – See which customers and features drive your AI API costs"
article_title: "Marginal"
image: ""
author: "jithinlalk25"
captured_at: "2026-08-19T13:37:10Z"
capture_tool: "hn-digest"
hn_id: 49361298
score: 2
comments: 0
posted_at: "2026-08-19T13:23:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Marginal – See which customers and features drive your AI API costs

- HN: [49361298](https://news.ycombinator.com/item?id=49361298)
- Source: [marginalhq.com](https://marginalhq.com/)
- Score: 2
- Comments: 0
- Posted: 2026-08-19T13:23:17Z

## Translation

タイトル: Show HN: Marginal – どの顧客と機能が AI API コストを押し上げているかを確認する
記事タイトル: 限界
説明: すべての LLM コールのコストを追跡し、顧客、機能、モデル (定義した任意のフィールド) ごとにコストをスライスします。
HN テキスト: デモ アカウント (共有、よろしくお願いします): https://marginalhq.com/login - メール jithinlalk25+demo@gmail.com、パスワード marginal-demo-2026。これは架空の会社のシードされたワークスペースです。概要/洞察/エクスプローラーを調べてみましょう。または、サインアップして実際のイベントを送信してください。

記事本文:
Marginal Marginal Docs テーマの切り替え サインアップ TypeScript SDK · Python SDK · HTTP API
Marginal は、すべての LLM 呼び出しのコストを追跡し、関心のあるフィールド (顧客、機能、モデル) ごとにコストをスライスします。
ドキュメントを読んでください。 npm install marginal-sdk · pip install marginal-sdk
1 人の顧客です — 定義したフィールドでグループ化することで検索されます
プロバイダー + モデル + 使用量の入力、ドルの出力 - アプリ内でのトークン計算は不要
「請求額が2倍になった」から明確な大義へ
製品で重要なものごとに支出をスライスし、自動的にグラフ化されたすべての側面を確認し、統合が送信した内容を正確に確認します。
チャートを作成せずにスパイクを特定する
Insights チャートはすべてのディメンションに対して自動的に支出します。各モデル、プロバイダー、および登録済みフィールドは独自のチャートを取得します。セットアップやクエリ ビルダーは必要ありません。
すべての API リクエストは、受け入れられた数、拒否されたイベントとその理由、剥奪されたキー、価格設定されていないモデルなどの結果とともにログに記録されます。推測せずに統合をデバッグします。
不明なモデル?イベントは依然として発生し、価格なしとしてフラグが立てられ、ダッシュボードと API 応答に表示され、平均化されることはありません。カスタム価格を設定すると、取り込み時に適用されます。
LLM リクエストごとに 1 回の呼び出し。それが全体の統合です。
npm install marginal-sdk 、 pip install marginal-sdk 、またはまったくインストールしない — 任意の言語から HTTP API に JSON を POST します。 SDK は依存関係がなく、バッファリングされ、ファイア アンド フォーゲットです。track() はリクエスト パスをスローしたりブロックしたりすることはありません。
プロバイダーに名前を付け、応答のモデルと使用状況をそのまま貼り付けます。Marginal は使用状況の形状を検出し、毎日同期されるモデルの価格カタログに対してサーバー側でコストを計算します。
顧客、機能、モデルなど、登録した任意のフィールドごとにグループ化およびフィルター処理します。何度も見に行くビューを保存します。
"marginal-sdk" から { Marginal } をインポートします。
const marginal = new Marginal ({ apiKey: process.env

。 MARGINAL_API_KEY });
const 応答 = openai.chat.completions を待ちます。作成する （{ /* … */ }）;
限界的な。トラック ({
プロバイダー: "openai" 、
モデル: 応答.モデル、
使用法: 応答.使用法、
フィールド: { 顧客: "acme-corp" 、機能: "support-bot" },
}); OSをインポートする
限界輸入から 限界輸入
marginal = Marginal( api_key = os.environ[ "MARGINAL_API_KEY" ])
応答 = client.chat.completions.create( ... )
マージナル.トラック(
プロバイダー = "openai" 、
モデル = 応答.モデル、
使用法 = response.usage.model_dump(),
フィールド = { "顧客" : "acme-corp" , "機能" : "サポートボット" },
)curl -X POST https://api.marginalhq.com/v1/events \
-H "認可: ベアラー $MARGINAL_API_KEY " \
-H "コンテンツ タイプ: application/json" \
-d '{
「イベント」: [
{ "プロバイダー": "openai",
"モデル": "gpt-4o-2024-08-06",
"使用法": { "プロンプトトークン": 2006, "完了トークン": 300 },
"フィールド": { "顧客": "acme-corp", "機能": "サポートボット" } }
】
}' コストは計算され、その日のカタログ料金で価格設定されます。コードベースでのトークン計算は必要ありません。
「請求書はなぜあんなに高かったのですか?」に答えるために作られました。
顧客、機能、切り分けたものなど、語彙を登録して、それに基づいて支出をグループ化またはフィルタリングします。未登録のキーは削除されて報告されるため、ダッシュボードはクリーンな状態に保たれます。
プロバイダー、モデル、および応答の使用法オブジェクトを送信します。コストは、毎日同期される価格カタログからサーバー側で計算され、プロジェクトごとにオーバーライドされます。価格は取り込み時に凍結されます。
エクスプローラーの状態 (範囲、フィルター、グループ化) は、名前付きビューとして保存されます。毎週の質問にワンクリックで回答できます。
TypeScript と Python、依存関係なし。イベントはローカルにバッファリングされ、バックグラウンドでバッチでフラッシュされます。ネットワーク障害は再試行してから警告します。アプリは決して気づきません。 SDK を使用しない方がよいですか? HTTP API は単一の JSON POST です。
コーディングエージェントにインテグラを実行させます

ション
marginalhq.com/llms.txt をクロード コード、カーソル、またはコーディング アシスタントに貼り付けると、Marginal を接続するために必要なものがすべて揃っています。またはプロバイダーごとのレシピに自分で従うこともできます。
ゼロからライブ支出ダッシュボードまで 5 分で作成できます。
プロジェクトを作成し、フィールドを登録し、track() 呼び出しを 1 つ追加します。次回のデプロイでは、質問への回答が始まります。
ドキュメントを読む 定義した任意のフィールドでスライスされた限界 AI コスト追跡。

## Original Extract

Track every LLM call's cost and slice it by customer, feature, model — any field you define.

Demo account (shared, please be nice): https://marginalhq.com/login - email jithinlalk25+demo@gmail.com, password marginal-demo-2026. It's a seeded workspace for a fictional company; poke around Overview/Insights/Explorer. Or sign up and send real events.

Marginal Marginal Docs Toggle theme Sign up TypeScript SDK · Python SDK · HTTP API
Marginal tracks the cost of every LLM call and slices it by the fields you care about — customer, feature, model.
Read the docs npm install marginal-sdk · pip install marginal-sdk
is one customer — found by grouping on a field you define
provider + model + usage in, dollars out — no token math in your app
From “the bill doubled” to a named cause
Slice spend by what matters in your product, see every dimension charted automatically, and check exactly what your integration sent.
Spot the spike without building a chart
Insights charts spend for every dimension automatically — each model, provider, and registered field gets its own chart. No setup, no query builder.
Every API request is logged with its outcome — accepted counts, rejected events and why, stripped keys, unpriced models. Debug your integration without guessing.
Unknown model? The event still lands, flagged as unpriced — visible in the dashboard and the API response, never averaged away. Set a custom price and it applies at ingest.
One call per LLM request. That's the whole integration.
npm install marginal-sdk , pip install marginal-sdk , or no install at all — POST JSON to the HTTP API from any language. The SDKs are zero-dependency, buffered, fire-and-forget: track() never throws and never blocks your request path.
Name the provider, paste the response's model and usage as-is — Marginal detects the usage shape and computes the cost server-side against a daily-synced model price catalog.
Group and filter by customer, feature, model — any field you register. Save the views you keep coming back to.
import { Marginal } from "marginal-sdk" ;
const marginal = new Marginal ({ apiKey: process.env. MARGINAL_API_KEY });
const response = await openai.chat.completions. create ({ /* … */ });
marginal. track ({
provider: "openai" ,
model: response.model,
usage: response.usage,
fields: { customer: "acme-corp" , feature: "support-bot" },
}); import os
from marginal import Marginal
marginal = Marginal( api_key = os.environ[ "MARGINAL_API_KEY" ])
response = client.chat.completions.create( ... )
marginal.track(
provider = "openai" ,
model = response.model,
usage = response.usage.model_dump(),
fields = { "customer" : "acme-corp" , "feature" : "support-bot" },
) curl -X POST https://api.marginalhq.com/v1/events \
-H "Authorization: Bearer $MARGINAL_API_KEY " \
-H "Content-Type: application/json" \
-d '{
"events": [
{ "provider": "openai",
"model": "gpt-4o-2024-08-06",
"usage": { "prompt_tokens": 2006, "completion_tokens": 300 },
"fields": { "customer": "acme-corp", "feature": "support-bot" } }
]
}' The cost lands computed, priced at that day's catalog rates — no token math in your codebase.
Built to answer “why was the bill that high?”
Register your vocabulary — customer, feature, anything you slice by — then group or filter spend by it. Unregistered keys are stripped and reported back, so dashboards stay clean.
Send provider, model, and the response's usage object; cost is computed server-side from a daily-synced price catalog, with per-project overrides. Prices are frozen at ingest.
Any Explorer state — range, filters, group-by — saves as a named view. The questions you ask every week are one click away.
TypeScript and Python, zero dependencies. Events buffer locally and flush in batches in the background; network failures retry and then warn — your app never notices. Prefer no SDK? The HTTP API is a single JSON POST.
Let your coding agent do the integration
Paste marginalhq.com/llms.txt into Claude Code, Cursor, or any coding assistant and it has everything it needs to wire up Marginal — or follow the per-provider recipes yourself.
From zero to a live spend dashboard in five minutes.
Create a project, register your fields, drop in one track() call. Your next deploy starts answering questions.
Read the docs Marginal AI cost tracking, sliced by any field you define.

---
source: "https://telemetry.dev/"
hn_url: "https://news.ycombinator.com/item?id=48997070"
title: "Show HN: Observability for Coding Agents and LLM Applications"
article_title: "Telemetry — Observability for AI and LLM apps"
author: "ephraimduncan"
captured_at: "2026-07-21T20:14:51Z"
capture_tool: "hn-digest"
hn_id: 48997070
score: 2
comments: 0
posted_at: "2026-07-21T19:33:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Observability for Coding Agents and LLM Applications

- HN: [48997070](https://news.ycombinator.com/item?id=48997070)
- Source: [telemetry.dev](https://telemetry.dev/)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T19:33:17Z

## Translation

タイトル: HN の表示: コーディング エージェントと LLM アプリケーションの可観測性
記事のタイトル: テレメトリ — AI および LLM アプリの可観測性
説明: すべてのモデル呼び出し、ツール ステップ、トークンの取得、コスト、レイテンシ、エラーを 1 か所でトレースします。 OpenTelemetry ネイティブ。最初のトレースは 5 分後に行われます。

記事本文:
テレメトリ — AI および LLM アプリの可観測性 Toggle Navigation 価格 統合 ログイン AI が何をしているかを推測するのはやめましょう
すべてのモデル呼び出し、ツール ステップ、および取得を 1 つのトレースで、トークン、コスト、レイテンシー、および添付されたエラーを実行します。 OpenTelemetry に基づいて構築: あらゆるフレームワーク、あらゆる言語、ロックインなし。最初のトレースは 5 分後に行われます。
コピー エージェント プロンプト 無料プラン: 1 か月あたり 10,000 のスパン。クレジットカードはありません。
何かが壊れたとき、その理由がわかります
1 つの遅いリクエストや 1 か月の費用: それに答えるためのトレースとメトリクスが迅速に得られます。
実際に確認する 4 つの数値は、モデル、プロバイダー、環境ごとに分かれています。
失敗は起こるごとに表面化する
失敗した呼び出しとタイムアウトは、エラーの種類、ステータス、および添付された完全なトレースとともに、発生時にストリーミングされます。
chat.completion · gpt-5.5 · 1,204 トークン · $0.004
取得 · 松ぼっくり · 8 ドキュメント · 120ms
入力と出力を自由に選択可能
環境ごとにキャプチャをオンにして、正確な入力と出力を保存します。シークレットは取り込み時に編集され、独自のパターンを追加できます。
あなたはサポートアシスタントです。短い 2 行で返信してください。
このサポート スレッドを要約します。
チェックアウト時に支払いが失敗しました — 顧客のカードが拒否されました。彼らは再試行リンクを望んでいます。
悪いトレースは 100 万件に 1 件でしょうか?見つかった。
名前で検索し、ステータス、環境、フレームワーク、またはセッションでフィルターし、数秒で完全なトレースを開きます。
SDK をドロップするか、任意の OpenTelemetry エクスポーターをエンドポイントにポイントします。どのフレームワークでも、どの言語でも、ロックインはありません。
あらゆるフレームワーク。どの言語でも。エンドポイントは 1 つ。
AI をアプリに追加した場合でも、ここにトレースを送信してください。これらをスパン、トークン、コストの 1 つのモデルに正規化します。
TypeScript および Python SDK: 関数をobserve() でラップすると、入力、出力、トークン、エラーがキャプチャされたスパンになります。 1 回のインストールで 1 つの API キー。
SDKは必要ありません。 OpenTelemetry エクスポーターを次のように指定します。

エンドポイントとトレースは正規化されて到着します。
入力、出力、キャッシュされたトークン、および推論トークンを長期にわたって追跡し、4,700 を超えるモデルのライブ価格設定から呼び出しごとにコストを計算します。
コストと遅延をモデル、プロバイダー、環境、ユーザーごとに分類します。セットアップはありません。
TypeScript ·generateText、streamText
telemetry.dev は、AI および LLM アプリ用の可観測性プラットフォームです。 OpenTelemetry トレースとしてすべてのモデル呼び出し、ツール ステップ、取得をキャプチャし、トークン、コスト、レイテンシ、エラーを添付します。これにより、障害をデバッグし、モデル、プロバイダー、環境全体での支出を追跡できます。
SDK をドロップするか、任意の OpenTelemetry エクスポーターをエンドポイントにポイントします。プロジェクト API キーを追加すると、最初のトレースが約 5 分で表示されます。
TypeScript と Python にはファーストパーティ SDK があり、Vercel AI SDK にはドロップイン統合があります。 HTTP 経由で標準の OTLP を発行する他のもの (LangChain、OpenLLMetry で計測されたアプリ、または独自の OpenTelemetry セットアップ) は、同じエンドポイントを通じて動作します。さらに多くのドロップイン統合が予定されています。
コストは、4,700 を超えるモデルの現在のモデルごとの価格と比較して、各呼び出しの実際のトークン使用量 (入力、出力、キャッシュ、および推論) からサーバー側で計算されます。クライアント側で同期を保つための見積もりは必要なく、必要に応じてスパンごとのコストを上書きできます。
はい、クレジットカードなしでも可能です。毎月 10,000 スパン、7 日間の保持、1 つのプロジェクト、2 つのシートを好きなだけ無料で利用できます。ボリュームがそれを超える場合にのみアップグレードしてください。
はい。入出力キャプチャは環境ごとに切り替えられるため、開発環境では完全なプロンプトを保存し、運用環境ではプロンプトをまったく保存しないことができます。 API キーなどのシークレットは取り込み時に編集され、プロジェクトごとにカスタムの編集パターンを追加できます。
最初のトレースは 5 分先にあります
1 つのパッケージをインストールし、API キーを追加し、リクエストを送信します。それが全体のセットアップです。
F

ree プランには月あたり 10,000 スパンが含まれます。クレジットカードは必要ありません。
telemetry.dev すべての AI アプリのトレース、トークン、コスト。
© 2026 telemetry.dev.無断転載を禁じます。

## Original Extract

Trace every model call, tool step, and retrieval with tokens, cost, latency, and errors in one place. OpenTelemetry-native. First trace in five minutes.

Telemetry — Observability for AI and LLM apps Toggle navigation Pricing Integrations Log in Stop guessing what your AI is doing
Every model call, tool step, and retrieval in one trace — tokens, cost, latency, and errors attached. Built on OpenTelemetry: any framework, any language, no lock-in. First trace in five minutes.
Copy agent prompt Free plan: 10k spans a month. No credit card.
When something breaks, you'll know why
One slow request or a month of spend: the traces and metrics to answer it, fast.
The four numbers you actually check, split by model, provider, and environment.
Failures surface as they happen
Failed calls and timeouts stream in as they happen, with error type, status, and the full trace attached.
chat.completion · gpt-5.5 · 1,204 tok · $0.004
retrieval · pinecone · 8 docs · 120ms
Inputs and outputs, on your terms
Flip capture on per environment to store exact inputs and outputs. Secrets are redacted at ingest, and you can add your own patterns.
You are a support assistant. Reply in two short lines.
Summarize this support thread.
Payment failed at checkout — the customer's card was declined. They want a retry link.
One bad trace in a million? Found.
Search by name, filter by status, environment, framework, or session, and open the full trace in seconds.
Drop in our SDK, or point any OpenTelemetry exporter at our endpoint. Any framework, any language, no lock-in.
Any framework. Any language. One endpoint.
However you added AI to your app, send the traces here. We normalize them into one model of spans, tokens, and cost.
TypeScript and Python SDKs: wrap a function in observe() and it becomes a span, with inputs, outputs, tokens, and errors captured. One install, one API key.
No SDK required. Point any OpenTelemetry exporter at our endpoint and traces arrive normalized.
Track input, output, cached, and reasoning tokens over time, with cost computed per call from live pricing for 4,700+ models.
Break cost and latency down by model, provider, environment, or user. No setup.
TypeScript · generateText, streamText
telemetry.dev is an observability platform for AI and LLM apps. It captures every model call, tool step, and retrieval as OpenTelemetry traces, then attaches tokens, cost, latency, and errors — so you can debug failures and track spend across models, providers, and environments.
Drop in our SDK, or point any OpenTelemetry exporter at our endpoint. Add your project API key and your first trace shows up in about five minutes.
TypeScript and Python have first-party SDKs, and the Vercel AI SDK has a drop-in integration. Anything else that emits standard OTLP over HTTP — LangChain, OpenLLMetry-instrumented apps, or your own OpenTelemetry setup — works through the same endpoint. More drop-in integrations are on the way.
Cost is computed server-side from each call's real token usage — input, output, cached, and reasoning — against current per-model pricing for 4,700+ models. There are no client-side estimates to keep in sync, and you can override cost per span when you need to.
Yes, with no credit card. You get 10k spans a month, 7-day retention, one project, and two seats, free for as long as you want. Upgrade only when your volume outgrows it.
Yes. Input and output capture is a per-environment switch, so you can store full prompts in development and none in production. Secrets like API keys are redacted at ingest, and you can add custom redaction patterns per project.
Your first trace is five minutes away
Install one package, add an API key, send a request. That's the whole setup.
Free plan includes 10k spans a month. No credit card required.
telemetry.dev Traces, tokens, and cost for every AI app.
© 2026 telemetry.dev. All rights reserved.

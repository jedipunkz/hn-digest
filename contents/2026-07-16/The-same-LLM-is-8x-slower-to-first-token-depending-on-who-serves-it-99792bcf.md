---
source: "https://dynoyard.app/blog/same-llm-8x-slower-by-backend/"
hn_url: "https://news.ycombinator.com/item?id=48939690"
title: "The same LLM is 8x slower to first token depending on who serves it"
article_title: "The same LLM is 8× slower to first token depending on who serves it — Dynoyard"
author: "jeremyrajan"
captured_at: "2026-07-16T20:53:13Z"
capture_tool: "hn-digest"
hn_id: 48939690
score: 2
comments: 0
posted_at: "2026-07-16T20:20:06Z"
tags:
  - hacker-news
  - translated
---

# The same LLM is 8x slower to first token depending on who serves it

- HN: [48939690](https://news.ycombinator.com/item?id=48939690)
- Source: [dynoyard.app](https://dynoyard.app/blog/same-llm-8x-slower-by-backend/)
- Score: 2
- Comments: 0
- Posted: 2026-07-16T20:20:06Z

## Translation

タイトル: 同じ LLM は、誰が提供するかによって、最初のトークンまでが 8 倍遅くなります
記事のタイトル: 同じ LLM は、誰が提供するかによって最初のトークンまでが 8 倍遅い — Dynoyard
説明: 2 つのバックエンド、同じオープンウェイト モデル、同じ OpenAI 互換 API、および最初のトークンまでの時間に 8 倍の差があり、そのうちの 1 つはストリーミングを偽装しています。エージェントトラフィックの LLM バックエンドのベンチマークで学んだこと、およびプロバイダーを使用する理由

記事本文:
dynoyard 価格設定について ドキュメント ブログ お問い合わせ サインイン 開始する メニュー 価格設定について ドキュメント ブログ 連絡先 サインイン ← エンジニアリング 同じ LLM は、誰が提供するかによって、最初のトークンまでが 8 倍遅くなります
私たちは、エージェントを構築しているチームに GLM-5.2 を提供します。同じオープンウェイト モデル、同じ OpenAI 互換 API ですが、それを複数のバックエンドにルーティングし、バックエンドを交換しているときに、書き留める価値のあるものを発見しました。選択したバックエンドでは、最初のトークンまでの時間が 8 倍変化し、そのうちの 1 つは実際にはまったくストリーミングしていませんでした。
同一モデルの 2 つのバックエンド。専用の推論ホストと、大手クラウド プロバイダーのマネージド OpenAI 互換エンドポイント。トークンあたりの価格が安く、非常に大きなコンテキスト ウィンドウが魅力的です。どちらも stream: true リクエストを受け入れ、SSE ストリームを返します。紙の上では交換可能。
キャッシュバス化された 40 個のリクエストをそれぞれストリーミングしてリクエストごとに記録します。最初のトークンまでの時間、合計時間、そして重要なことですが、トークンが実際に段階的に到着するか最後に一度に到着するかどうかを記録します。数値 (p50、400 トークン出力):
安価なエンドポイントは応答全体を生成し、何も送信せずに約 10 秒間待機し、その後すべてのトークンを一度にダンプします。これは stream: true に準拠しており、バイトは SSE として到着しますが、増分配信はありません。チャット UI やエージェント ループの場合、毎ターン 10 秒の停止が必要になります。 「ストリーミング」は表面的なものです。
それは私たちから遠く離れた雲の領域です。確かにそれはレイテンシーです。そこで、より近い地域に移動しました。結果: ~18% 高速になりましたが、まだバッファされています。リージョンが原因ではありませんでした。関係なく、サーバー側でエンドポイント バッファーが生成されました。そして、まったく同じパスで専用ホストを実行することで、それが独自のゲートウェイではないことを確認しました。167 チャンクをクリーンにストリーミングしました。バッファリングは提供に属します

えーっと。
そしてひねり: キャッシュがすべてを変える
キャッシュバストされたテストは最悪のケースです。実際のエージェントのトラフィックは大量にキャッシュされます。長く安定したシステム プロンプトとツール定義がターンごとに表示され、最後尾のみが変化します。そこで、ウォーム キャッシュされたプレフィックスを使用して再実行したところ、安価なエンドポイントの TTFT が 10 秒から 2.7 秒に低下し、バッファリングがほとんど停止しました。この異常な数は、実稼働エージェントが実際に送信したキャッシュされていないコールド リクエストのアーティファクトでした。現実的なキャッシュ負荷の下では、その差は大幅に縮まりますが、依然として専用ホストがすべての軸で約 1.6 倍の勝利を収めています。
この逆転こそが本当の教訓だ。キャッシュバスティングのベンチマークを行って終わりにしていたら、エンドポイントは使用できないほど遅いと判断されていたでしょう。プロバイダーの仕様書を信頼していたら、バッファリングをまったく目にすることはなかったでしょう。
「OpenAI 互換」では、リクエストの形式がわかりますが、動作については何もわかりません。同一のモデルを提供する 2 つのバックエンドは、ユーザーが感じる指標で 8 倍も異なり、1 つはストリーミングを偽装し、リージョンは修正せず、キャッシュの形状は結論全体を反転させました。
キャッシュバス化された合成ループやプロバイダーの仕様シートではなく、実際のトラフィック形状でバックエンドをベンチマークします。実際の実行方法でキャッシュされます。重要な数値は、独自のゲートウェイを通じて測定された、キャッシュされたプロンプトのストリームごとの有効スループットです。これを、キャッシュ高速応答を除いた、出力トークンを (整定時間 − 最初のトークン時間) で割った p50 として定義します。これは、負荷時に快適に過ごせる合計ストリーム集計ではなく、単一世代で実際に感じる速度です。
専用ホストをプライマリとして、クラウド エンドポイントをフェイルオーバーとして維持しました。これには、他のホストにはないはるかに大きなコンテキスト ウィンドウがあり、ウォーム キャッシュの下では完全に使用可能です。しかし、私たちが知っているのは、すべてのバックアップをベンチマークしているためです。

モデルが変動し、プロバイダーが誰にも告げることなくサービススタックを変更するため、リクエストを処理する前に、スケジュールどおりに実行されることが増えています。 「昨夜、私のプロバイダーが静かにストリーミングやツールの呼び出しを中断したのか」という質問に決して答えることができないのであれば、それが私たちが構築しているものです。
1 つの OpenAI 互換エンドポイントの背後にあるすべてのモデル。独自のベイクオフを実行します。
hello@dynoyard.app
当社では、Dynoyard を改善するための製品分析に Cookie を使用しています。私たちのを参照してください
プライバシーポリシー 。

## Original Extract

Two backends, the same open-weight model, the same OpenAI-compatible API — and an 8× gap in time-to-first-token, with one of them faking streaming. What we learned benchmarking LLM backends for agentic traffic, and why the provider

dynoyard About Pricing Docs Blog Contact Sign in Get started Menu About Pricing Docs Blog Contact Sign in ← Engineering The same LLM is 8× slower to first token depending on who serves it
We serve GLM-5.2 to teams building agents. Same open-weight model, same OpenAI-compatible API — but we route it across more than one backend, and while swapping one in we found something worth writing down: the backend you pick changes time-to-first-token by 8×, and one of them wasn’t really streaming at all.
Two backends for the identical model. A dedicated inference host, and a major cloud provider’s managed OpenAI-compatible endpoint — cheaper per token, a very large context window, tempting. Both accept a stream: true request and return an SSE stream. On paper, interchangeable.
We fire 40 cache-busted requests at each, streaming, and record per request: time-to-first-token, total time, and — the important one — whether tokens actually arrive incrementally or all at once at the end. Numbers (p50, 400-token outputs):
The cheap endpoint generates the entire response, sits on it for ~10 seconds emitting nothing, then dumps all the tokens at once. It’s stream: true -compliant to the letter — the bytes do arrive as SSE — but there’s no incremental delivery. For a chat UI or an agent loop, that’s a dead 10-second pause every turn. The “streaming” is cosmetic.
It’s a cloud region far from us; surely that’s the latency. So we moved it to a closer region. Result: ~18% faster, still buffered. Region wasn’t the cause — the endpoint buffers generation server-side regardless. And we confirmed it wasn’t our own gateway by running the dedicated host through the exact same path: it streamed 167 chunks cleanly. The buffering belongs to the provider.
Then the twist: caching changes everything
Cache-busted tests are worst-case. Real agent traffic is heavily cached — long stable system prompts and tool definitions on every turn, only the tail changing. So we re-ran with a warm cached prefix, and the cheap endpoint’s TTFT dropped from 10s to 2.7s , and it mostly stopped buffering. The pathological number was an artifact of cold, uncached requests that no production agent actually sends. Under realistic cached load the gap narrows a lot — though the dedicated host still won ~1.6× on every axis.
That reversal is the real lesson. If we’d benchmarked cache-busted and called it a day, we’d have written the endpoint off as unusably slow. If we’d trusted the provider’s spec sheet, we’d never have seen the buffering at all.
“OpenAI-compatible” tells you the request shape, nothing about behavior. Two backends serving the identical model differed 8× on the metric users feel, one faked streaming, region didn’t fix it, and cache shape flipped the whole conclusion.
Benchmark backends in your real traffic shape — cached the way you actually run — not with a cache-busted synthetic loop, and not from the provider’s spec sheet. The number that matters is per-stream effective throughput on cached prompts, measured through your own gateway. We define it as output tokens divided by (settle time − first-token time), p50, with cache-fast responses excluded — the speed a single generation actually feels, not a summed-stream aggregate that flatters you under load.
We kept the dedicated host as primary and the cloud endpoint as a failover — it has a much larger context window the other can’t match, and under warm cache it’s perfectly usable. But we only knew any of this because we benchmark every backend before it serves a request — and increasingly, on a schedule, because models drift and providers change serving stacks without telling anyone. If “did my provider quietly break streaming or tool-calls last night” is a question you’ve never been able to answer, that’s the thing we’re building.
Every model behind one OpenAI-compatible endpoint. Run your own bake-off.
hello@dynoyard.app
We use cookies for product analytics to improve Dynoyard. See our
privacy policy .

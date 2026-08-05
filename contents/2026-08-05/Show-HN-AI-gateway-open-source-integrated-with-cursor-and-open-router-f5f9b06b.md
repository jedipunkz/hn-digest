---
source: "https://www.routerfuel.com/"
hn_url: "https://news.ycombinator.com/item?id=49177455"
title: "Show HN: AI gateway open source integrated with cursor and open router"
article_title: "RouterFuel — 330+ LLMs, Single API"
author: "U_Zargar"
captured_at: "2026-08-05T01:41:28Z"
capture_tool: "hn-digest"
hn_id: 49177455
score: 1
comments: 0
posted_at: "2026-08-05T01:18:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI gateway open source integrated with cursor and open router

- HN: [49177455](https://news.ycombinator.com/item?id=49177455)
- Source: [www.routerfuel.com](https://www.routerfuel.com/)
- Score: 1
- Comments: 0
- Posted: 2026-08-05T01:18:34Z

## Translation

タイトル: HN を表示: カーソルとオープン ルーターを統合した AI ゲートウェイ オープン ソース
記事のタイトル: RouterFuel — 330 以上の LLM、単一 API
説明: RouterFuel は、セマンティック キャッシュと自動フェイルオーバーを備えた 330 以上のモデル (OpenAI、Anthropic、Google、Meta、Mistral、および OpenRouter 上のすべてのモデル) にルーティングする非同期 Rust LLM ゲートウェイです。

記事本文:
ルーターの燃料
特長
ゲートウェイの仕様
価格設定
お問い合わせ
サインイン
GitHub
ウェイティングリストに参加する
始めましょう
🚀 RouterFuel はオープンベータ版です。プレビュー中はすべての機能が無料です。
330 以上のモデル、1 つの非同期プロキシ、
エッジで調整されています。
RouterFuel は、使用するすべてのモデル プロバイダー (直接経由で 330 以上のモデル) の前にある単一の非同期プロキシ レイヤーです。
統合と OpenRouter。 Tokio と Axum のネットワーク コア上に構築され、同時推論を多重化します。
測定可能なオーバーヘッドを追加することなく、リクエストを処理し、マイクロ秒単位でプロバイダーのフェイルオーバーをネゴシエートし、トークンをストリームバックします。
ドキュメントを見る
カーゴ インストール Routerfuel · クラスタ外に送信されたゼロ テレメトリ · AGPL-3.0 コア
コピー
#1. インストールする
貨物設置ルーター燃料
#2. 走る
ROUTEFUEL_API_KEYS="sha256:テスト:テストクライアント" \
DATABASE_URL="postgres://..." \
./routerfuel
# 3. 最初のリクエストを行う
curl -X POST http://localhost:3000/v1/chat/completions \
-H "X-API-キー: テスト" \
-H "コンテンツ タイプ: application/json" \
-d '{"モデル":"自動","メッセージ":[{"役割":"ユーザー","コンテンツ":"Hello!"}]}'
routerfuel —edge-proxy@us-east-1
ライブ
p99 レイテンシ
単一の API を通じて 330 以上の LLM (OpenAI、Anthropic、Google、Meta、Mistral、および OpenRouter 上のすべてのモデル) にアクセスします。
✅ OpenRouter と連携 — 独自のキーを持ち込む
✅ 新しいモデルのカタログの自動同期
✅ すべてのプロバイダーで 1 つのリクエスト形式
すべての RouterFuel 導入には、同じ復元コアが同梱されています。アドオン モジュールはありません。
推論層を維持する部分に対して個別に請求されることはありません。
受信プロンプトを以前の完了とベクトル照合し、キャッシュからほぼ重複したものを提供することで、反復的な API 支出を最大 40% 削減します。
プロバイダーのエラーのスパイクや遅延のドリフトが発生した場合は自動的にトリップし、処理中のリクエストと新しいリクエストを、通話がドロップされることなく正常なフォールバック プロバイダーに再ルーティングします。
ヘッダーを挿入して書き換える

リクエストボディを作成し、ラウンドトリップを追加せずにインラインで実行されるミドルウェア チェーンを使用して、上流のサーバーの健全性パラメータを追跡します。
OpenRouter の統合 — 300 以上のモデル
OpenRouter がサポートするあらゆるモデルにシームレスにルーティングします。自動カタログ同期により、コードを変更せずに新しいモデルが表示されます。
シートではなくスループットに対して支払います。
すべての層には完全な復元コアが同梱されています。オープンベータ期間中はすべての機能が無料です。
サイドプロジェクトやプロトタイプ用。
✓ 35,000 リクエスト/月を含む
✓ 超過料金は 100,000 リクエストあたり 20 ドルで請求されます
✓ セマンティック キャッシュとサーキット ブレーク
大規模な実稼働ワークロード向け。
✓ 150,000 リクエスト/月を含む
✓ 超過料金は 100,000 リクエストあたり 15 ドルで請求されます
✓ 優先プロバイダーのフェイルオーバー ルーティング
規制されたエアギャップ環境向け。
✓ 専用のマルチリージョン展開
✓ オンプレミスのバイナリ配布
✓ エアギャップセキュリティプロファイル
質問、フィードバック、または導入について話したいことがありますか?直接ご連絡ください。
contact@routerfuel.com
// ライセンス付与
RouterFuel のコアは AGPL-3.0 に基づいてライセンスされています。それは次のような場合に最適です
セルフホスティングおよびオープンソースの使用ですが、RouterFuel をクローズドソース製品に埋め込んでいる場合、または次のような問題が発生した場合
貴社での AGPL のネットワーク使用条項については、別途商用/エンタープライズ ライセンスを提供します。
コピーレフトの義務。
© 2026 RouterFuel — デフォルトでは非同期。
AGPL-3.0
サインイン
ドキュメント
GitHub
価格設定
お問い合わせ
contact@routerfuel.com
サインイン
アカウントはオープンベータ後に開始されます。待機リストに参加すると、公開された瞬間に通知が届きます。
まだサインアップしていませんか?順番待ちリストに参加する
オープンベータ期間中は無料。アクセスの準備ができたらメールでお知らせします。
アクセスの準備ができ次第、メールでお知らせします。

## Original Extract

RouterFuel is an async Rust LLM gateway routing to 330+ models — OpenAI, Anthropic, Google, Meta, Mistral, and every model on OpenRouter — with semantic caching and automatic failover.

ROUTER FUEL
Features
Gateway Specs
Pricing
Contact
Sign In
GitHub
Join Waitlist
Get Started
🚀 RouterFuel is in open beta. All features are free during preview.
330+ models, one async proxy,
orchestrated at the edge.
RouterFuel is a single async proxy layer in front of every model provider you use — 330+ models via direct
integration plus OpenRouter. Built on a Tokio and Axum network core, it multiplexes concurrent inference
requests, negotiates provider failover in microseconds, and streams tokens back without adding measurable overhead.
View Documentation
cargo install routerfuel · zero telemetry sent off-cluster · AGPL-3.0 core
Copy
# 1. Install
cargo install routerfuel
# 2. Run
ROUTEFUEL_API_KEYS="sha256:test:testclient" \
DATABASE_URL="postgres://..." \
./routerfuel
# 3. Make your first request
curl -X POST http://localhost:3000/v1/chat/completions \
-H "X-API-Key: test" \
-H "Content-Type: application/json" \
-d '{"model":"auto","messages":[{"role":"user","content":"Hello!"}]}'
routerfuel — edge-proxy@us-east-1
LIVE
p99 latency
Access 330+ LLMs through a single API — OpenAI, Anthropic, Google, Meta, Mistral, and every model on OpenRouter.
✅ Works with OpenRouter — bring your own key
✅ Automatic catalog sync for new models
✅ One request format across every provider
Every RouterFuel deployment ships with the same resiliency core — no add-on modules,
no separate billing for the parts that keep your inference layer alive.
Vector-matches incoming prompts against prior completions and serves near-duplicates from cache, cutting repetitive API spend by up to 40%.
Trips automatically on provider error spikes or latency drift, rerouting in-flight and new requests to a healthy fallback provider with no dropped calls.
Inject headers, rewrite request bodies, and track upstream server health parameters with a middleware chain that runs inline with zero added round-trips.
OpenRouter Integration — 300+ Models
Seamlessly route to any model OpenRouter supports. Automatic catalog sync means new models appear without code changes.
Pay for throughput, not seats.
Every tier ships the full resiliency core. All features are free during the open beta.
For side projects and prototypes.
✓ 35,000 requests / month included
✓ Overages billed at $20 / 100k requests
✓ Semantic caching & circuit breaking
For production workloads at scale.
✓ 150,000 requests / month included
✓ Overages billed at $15 / 100k requests
✓ Priority provider failover routing
For regulated and air-gapped environments.
✓ Dedicated multi-region deployment
✓ On-premise binary distribution
✓ Air-gapped security profiles
Questions, feedback, or want to talk deployment? Reach us directly.
contact@routerfuel.com
// licensing
RouterFuel's core is licensed under AGPL-3.0 . That works great for
self-hosting and open-source use, but if you're embedding RouterFuel in a closed-source product or run into
AGPL's network-use clause at your company, we offer a separate commercial/enterprise license with no
copyleft obligations.
© 2026 RouterFuel — async by default.
AGPL-3.0
Sign In
Docs
GitHub
Pricing
Contact
contact@routerfuel.com
Sign in
Accounts launch after the open beta — join the waitlist to get notified the moment they're live.
Not signed up yet? Join the waitlist
Free during the open beta. We'll email you when your access is ready.
We'll email you as soon as your access is ready.

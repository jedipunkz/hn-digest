---
source: "https://www.periskop.ai"
hn_url: "https://news.ycombinator.com/item?id=48738869"
title: "Periskop – product discovery (MCP/API) for AI agents"
article_title: "Periskop — Product discovery for AI agents"
author: "duarteapereira"
captured_at: "2026-06-30T21:34:51Z"
capture_tool: "hn-digest"
hn_id: 48738869
score: 2
comments: 0
posted_at: "2026-06-30T20:40:07Z"
tags:
  - hacker-news
  - translated
---

# Periskop – product discovery (MCP/API) for AI agents

- HN: [48738869](https://news.ycombinator.com/item?id=48738869)
- Source: [www.periskop.ai](https://www.periskop.ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T20:40:07Z

## Translation

タイトル: Periskop – AI エージェント向けの製品検出 (MCP/API)
記事のタイトル: Periskop — AI エージェント向けの製品発見
説明: 1 回の MCP/API 呼び出しで、ショッピング意図が実際の製品の選択、代替品、バンドル、価格、画像、注意事項、一致品質、リクエスト ID、販売者リンクに変わります。チェックアウトも支払いもありません。販売者リンクのみです。

記事本文:
Periskop — AI エージェント向けの製品検出 eriskop 製品開発者向け 価格設定ドキュメント ドキュメントの表示 MCP/API アクセスのリクエスト 概要
AI エージェント向けの製品発見。
構築しているエージェント向けに、安価で信頼性が高く、幻覚のない製品を発見します。 1 つの MCP/API 呼び出しにより、ショッピングの意図が、実際の価格と有効な販売者リンクを備えた実際の在庫商品に変わります。
出力例 · ホバーして一時停止するか、answer_type を選択します
MCP または REST を通じて Periskop を使用します。
MCP サーバーを Claude Code、Cursor、その他の ChatGPT 互換またはカスタム エージェント ランタイムに接続するか、任意のアプリ バックエンドから REST エンドポイントを呼び出します。キー、使用法、請求書、ドキュメント、OpenAPI は開発者ポータルにあります。
// MCP、クロード コード、カーソル、カスタム エージェント ランタイム
{
"mcpサーバー": {
「ペリスコープ」: {
"url": "https://mcp.periskop.ai",
"headers": { "Authorization": "Bearer dp_..." }
}
}
} バックエンド用 REST // REST、任意のアプリ バックエンドから
POST /ショッピング/発見
権限: ベアラー dp_...
{ "prompt": "...", "mode": "bundle" } 開発者ポータル内 API キー 使用法 Billing Docs OpenAPI プロジェクトを作成し、キーを作成し、 run_shopping_discovery を呼び出します。すべての結果には result_id が含まれており、get_discovery_result で再取得します。
Periskop を既存の自動化スタック、Dify、n8n、Pipedream のネイティブ ステップにドロップします。
検索するとリンクが返されます。 Periskop は適切な製品を返します。
汎用エージェントは検索し、ページごとに開き、リクエストごとに LLM で再読み取りしますが、速度が遅く、トークンが大量に使用され、信頼性が低くなります。 Periskop は 1 回の構造化呼び出しで製品を返し、エージェントが直接レンダリングして貼り付け可能な応答を返します。
すべてのリクエストプロンプト ▸ " ▋ " 01 web_search × 1 · 強力な一致なしコスト 02 ページを開く · 古い/在庫切れのコスト 03 web_search · 条件を広げる · まだ薄いコスト 04 ページを開く × 1 · 404 · デッドリンクのコスト 05 llm rea

すべてを読み込む · すべてのコストを読み取る 遅延は予測不可能 信頼性が低い ↓ → まだ信頼できる製品がない N ステップ → Periskop で 1 回のコール · レンダリングされた JSON プロンプト ▸ " ライアンエアーのサイズ制限に適合するキャビンバッグ、60 ユーロ未満 " run_shopping_discovery → 200 OK 回答タイプの推奨事項 · 概要 40×20×25cm 以内の在庫ありのピック 2 個 Cabin Max Metz (40×20×25) amazon.fr 39.99 ユーロ
Aerolite Ryanair キャビン バッグ decathlon.fr EUR 27.99
ライアンエアーの 40×20×25 cm 制限に適合する 60 ユーロ未満の 2 つのキャビンバッグ: Cabin Max Metz (39.99 ユーロ) と Aerolite Ryanair バッグ (27.99 ユーロ)、どちらも在庫あり。より安価なもの、またはハードシェルのオプションが必要ですか?
ランク付け、在庫あり、リンク済み、表示準備完了 スワイプして比較 → 同じ取得、ランキング、検証を事前に 1 回実行し、リクエストごとに再導出するわけではありません。品質管理された供給 オンボーディングは信頼されることを意味するものではありません。
Periskop は、ストアがレコメンデーションに影響を与える前に、ストアの構造、製品の抽出、価格の読みやすさ、リンクの有効性、画像の存在、安定性を検証します。店舗は推奨事項を形成する権利を獲得します。
信頼できる、最高のピックを形作ることができます
限定された、代替または制限されたコンテキスト
拒否されました。推奨事項には影響しません
チェックアウトは常に解決されました。 Periskop が登場するまでは、何を買えばいいのかわかりませんでした。
支払い、カート、フルフィルメントにはすでに 100 のプロバイダーがあり、その部分が問題になることはありませんでした。答えのない質問は貴重な質問です。販売されているすべてのものの中で、エージェントが実際に何を勧めるべきですか? Periskop がその層です。適切な在庫商品を見つけて、クリーンな販売者リンクを渡します。残りは既存のレールが行います。
Stripe · 販売者 · ストアフロント
電話一本。ライブ製品と価格が必要なものは何でも構築します。
同じ run_shopping_discovery 呼び出しにより、チャット アシスタント、コンピュータ使用エージェント、価格監視サービスが強化されます。

イタリング ボット、アービトラージ スキャナー、調達ワークフロー、コマース コパイロット、ショッピング アプリ全体など、次に想像するものはすべて含まれます。あなたは体験を発送します。 Periskop はその下にある 1 つの呼び出しです。
AI ショッピング アシスタント ブラウザーとコンピューター使用エージェント 価格と在庫の監視 アービトラージと再販ソーシング 調達と運用 コマースの副操縦士 AI ネイティブ ショッピング アプリ アクセシビリティ ショッピング + …そして次回出荷するものはすべて run_shopping_discovery 経由で AI ショッピング アシスタント 本物の在庫品のピックアップがチャットに直接ドロップされ、ランク付けされ、価格設定され、リンクされます。
50 ユーロ未満の 2 つの堅実なピック、どちらも耐汗性: JBL Wave Buds amazon.fr €34.99 オープン Anker Soundcore P20i fnac €29.99 オープン 40 以上の在庫からランク付けされる · ライブ価格 使用量ベース · 7 ユーロ / 1,000 ショッピング リクエスト · 最初のキーの無料クレジット 価格を参照 → Periskop で構築 構築中のエージェント向けの製品検出。
AI ショッピング アシスタント、ブラウザおよびコンピュータを使用するエージェント、調達ワークフロー、コマース副操縦士、レコメンデーション システム、AI ネイティブ ショッピング アプリ向け。
Periskop は AI エージェント向けの製品発見です。 1 回の MCP/API 呼び出しでショッピング リクエストが、実際の価格、有効な販売者リンク、代替品、推測されたバンドル、注意事項、貼り付け可能な返信、result_id を備えた実際の在庫のある商品に変換されるため、エージェントはリクエストごとに Web をクロールする代わりに推奨できるようになります。
AI ショッピング アシスタント、ブラウザーとコンピューターを使用するエージェント、調達と運用ワークフロー、コマース副操縦士、レコメンデーション機能、アクセシビリティ ショッピング、AI ネイティブ ショッピング アプリを構築する開発者。
MCP または REST 経由で run_shopping_discovery にプロンプ​​トを送信します。 Periskop は、品質ゲート型販売業者全体を取得してランク付けし、構造化された応答、アイテム、または推測されたバンドルに加えて、貼り付け準備ができている提案済みエージェント応答と次のアクションを返します。

エージェントが直接レンダリングします。サイトごとの閲覧や実行時のスクレイピングはありません。
いいえ、Periskop はストアの構造、抽出、価格、リンク、画像、安定性を検証するため、品質管理されたすぐに発見できる供給のみが結果に影響を与えることができます。商人は何も言わない。
いいえ、検索するとリンクが返されるので、まだ開いて検討する必要があります。 Periskop は、1 回の呼び出しで、match_quality と警告を使用して、ランク付け、在庫、価格設定、リンクされた製品自体を返します。
いいえ。チェックアウトも支払いもありません。販売者リンクのみです。 Periskop は、購入したり、注文したり、販売者のチェックアウト カートを作成したりすることはありません。ユーザーは販売サイトで購入を完了します。
Periskop は項目に match_quality: null をマークし、強制または幻覚による選択ではなく警告と、最も近い一致と代替を返します。すべての結果には result_id が含まれるため、後で get_discovery_result を使用して再取得できます。
エージェント ランタイム (クロード コード、カーソル、カスタム エージェント) 用の MCP サーバーとアプリ バックエンド用の REST エンドポイント。キー、使用法、請求書、ドキュメント、OpenAPI は開発者ポータルにあります。
使用量ベース: ショッピング リクエスト 1,000 件あたり 7 ユーロ、最初の API キーには無料クレジットが付与されます。早期アクセスは創設者主導のオンボーディングです。
AI エージェント向けの製品発見。 1 つの MCP/API 呼び出しにより、ショッピング意図が価格と販売者リンクを含む実際の商品に変換されます。

## Original Extract

One MCP/API call turns shopping intent into real product picks, alternatives, bundles, prices, images, caveats, match quality, request IDs, and merchant links. No checkout, no payments — merchant links only.

Periskop — Product discovery for AI agents eriskop Product Developers Pricing Docs View docs Request MCP/API access Overview
Product discovery for AI agents.
Cheap, reliable, hallucination-free product discovery for the agents you're building. One MCP/API call turns shopping intent into real, in-stock products, with live prices and working merchant links.
Example output · hover to pause, or pick an answer_type
Use Periskop through MCP or REST.
Connect the MCP server to Claude Code, Cursor, and other ChatGPT-compatible or custom agent runtimes, or call the REST endpoint from any app backend. Keys, usage, billing, docs, and OpenAPI live in the developer portal.
// MCP, Claude Code, Cursor, custom agent runtimes
{
"mcpServers": {
"periskop": {
"url": "https://mcp.periskop.ai",
"headers": { "Authorization": "Bearer dp_..." }
}
}
} REST for backends // REST, from any app backend
POST /shopping/discover
Authorization: Bearer dp_...
{ "prompt": "...", "mode": "bundle" } In the developer portal API keys Usage Billing Docs OpenAPI Create a project, mint a key, and call run_shopping_discovery . Every result carries a result_id, re-fetch it with get_discovery_result.
Drop Periskop into your existing automation stack, native steps for Dify, n8n, and Pipedream.
Search returns links. Periskop returns the right products.
A generic agent searches, opens page after page, and re-reads it with an LLM on every request, slow, token-heavy, unreliable. Periskop returns the products in one structured call, with a paste-ready reply your agent renders directly.
every request prompt ▸ " ▋ " 01 web_search × 1 · no strong match cost 02 open page · stale / out of stock cost 03 web_search · broaden terms · still thin cost 04 open page × 1 · 404 · dead link cost 05 llm reads it all · reads everything cost Latency unpredictable Reliability low ↓ → still no reliable product N steps → 1 call With Periskop · one call rendered json prompt ▸ " cabin bag that fits Ryanair's size limit, under €60 " run_shopping_discovery → 200 OK answer_type recommendation · summary 2 in-stock picks within 40×20×25cm Cabin Max Metz (40×20×25) amazon.fr EUR 39.99
Aerolite Ryanair Cabin Bag decathlon.fr EUR 27.99
Two cabin bags under €60 that fit Ryanair's 40×20×25 cm limit: the Cabin Max Metz (€39.99) and the Aerolite Ryanair bag (€27.99), both in stock. Want the cheaper one, or a hard-shell option?
ranked, in-stock, linked, ready to show Swipe to compare → Same retrieval, ranking, and validation, run once, ahead of time, not re-derived on every request. Quality-gated supply Onboarded does not mean trusted.
Periskop validates store structure, product extraction, price readability, link validity, image presence, and stability before a store can influence recommendations. Stores earn the right to shape recommendations.
Trusted , can shape best picks
Limited , alternatives or restricted contexts
Rejected , cannot influence recommendations
Checkout was always solved. Knowing what to buy wasn't, until Periskop.
Payments, carts, and fulfillment already have a hundred providers, that part was never the problem. The unanswered question is the valuable one: out of everything for sale, what should your agent actually recommend? Periskop is that layer. It finds the right in-stock products and hands off a clean merchant link; your existing rails do the rest.
Stripe · merchants · your storefront
One call. Build whatever needs live products and prices.
The same run_shopping_discovery call powers a chat assistant, a computer-use agent, a price-monitoring bot, an arbitrage scanner, a procurement workflow, a commerce copilot, a whole shopping app, and whatever you imagine next. You ship the experience; Periskop is the one call underneath.
AI shopping assistants Browser & computer-use agents Price & stock monitoring Arbitrage & resale sourcing Procurement & ops Commerce copilots AI-native shopping apps Accessibility shopping + …and whatever you ship next AI shopping assistants via run_shopping_discovery Real, in-stock picks dropped straight into chat, ranked, priced, and linked.
Two solid picks under €50, both sweat-resistant: JBL Wave Buds amazon.fr €34.99 open Anker Soundcore P20i fnac €29.99 open ranked from 40+ in-stock · live prices Usage-based · €7 / 1,000 shopping requests · free credits on your first key See pricing → Build with Periskop Product discovery for the agents you are building.
For AI shopping assistants, browser and computer-use agents, procurement workflows, commerce copilots, recommendation systems, and AI-native shopping apps.
Periskop is product discovery for AI agents. One MCP/API call turns a shopping request into real, in-stock products, with live prices, working merchant links, alternatives, inferred bundles, caveats, a paste-ready reply, and a result_id, so your agent can recommend instead of crawling the web on every request.
Developers building AI shopping assistants, browser and computer-use agents, procurement and ops workflows, commerce copilots, recommendation features, accessibility shopping, and AI-native shopping apps.
You send a prompt to run_shopping_discovery over MCP or REST. Periskop retrieves and ranks across quality-gated merchants and returns a structured response, items or an inferred bundle, plus a paste-ready suggested_agent_reply and next_actions, your agent renders directly. No per-site browsing, no runtime scraping.
No. Periskop validates store structure, extraction, prices, links, images, and stability so only quality-gated, discovery-ready supply can influence results. Merchants don't lift a finger.
No. Search returns links you still have to open and reason over. Periskop returns the products themselves, ranked, in-stock, priced, and linked, in one call, with match_quality and caveats.
No. No checkout, no payments, merchant links only. Periskop never buys, places orders, or creates merchant checkout carts. Users complete purchases on merchant sites.
Periskop marks items match_quality: null and returns caveats rather than a forced or hallucinated pick, plus the closest matches and alternatives. Every result carries a result_id, so you can re-fetch it later with get_discovery_result.
An MCP server for agent runtimes (Claude Code, Cursor, custom agents) and a REST endpoint for app backends. Keys, usage, billing, docs, and OpenAPI live in the developer portal.
Usage-based: €7 per 1,000 shopping requests, with free credits on your first API key. Early access is founder-led onboarding.
Product discovery for AI agents. One MCP/API call turns shopping intent into real products, with prices and merchant links.

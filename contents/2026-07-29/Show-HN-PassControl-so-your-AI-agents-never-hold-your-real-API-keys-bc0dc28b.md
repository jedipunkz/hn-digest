---
source: "https://passcontrol.vertias.eu"
hn_url: "https://news.ycombinator.com/item?id=49096412"
title: "Show HN: PassControl – so your AI agents never hold your real API keys"
article_title: "PassControl — Keep real API keys out of AI agents"
author: "vertias3u"
captured_at: "2026-07-29T12:56:28Z"
capture_tool: "hn-digest"
hn_id: 49096412
score: 1
comments: 0
posted_at: "2026-07-29T12:09:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PassControl – so your AI agents never hold your real API keys

- HN: [49096412](https://news.ycombinator.com/item?id=49096412)
- Source: [passcontrol.vertias.eu](https://passcontrol.vertias.eu)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T12:09:41Z

## Translation

タイトル: HN を表示: PassControl – AI エージェントが実際の API キーを保持しないようにする
記事のタイトル: PassControl — 本物の API キーを AI エージェントから遠ざける
説明: AI エージェント用のソース利用可能な ID および認証情報ゲートウェイ。エージェントに実際のプロバイダー キーを与えることなく、範囲、予算、取り消し、監査を強制します。

記事本文:
PassControl — 実際の API キーを AI エージェントから遠ざける コンテンツにスキップ ver · tias / PassControl 仕組み ライブ デモ 機能 ソースが利用可能な PassControl / エージェント認証情報ゲートウェイ
AI エージェントは実際の API キーを決して保持すべきではありません。
通話料金を支払う資格情報ではなく、暗号化された ID をエージェントに提供します。 PassControl は各エージェントを検証し、スコープとバジェットを適用し、保管されたプロバイダー キーを挿入して、リクエストをプロキシします。キーがエージェントに届くことはありません。
ID は転送されますが、プロバイダーの秘密は転送されません
アイデンティティは境界を越える。秘密はそうではありません。
パスポートは署名のみです。有効期限の短い就労ビザは、プロバイダーの資格情報を解決する前にすべてのリクエストをチェックするゲートウェイに ID とポリシーを組み込みます。
プライベート Ed25519 キーはローカルで 1 回限りのチャレンジに署名します。有線で送信されることはありません。
01 ビザの本人確認と有効期限の確認
02 キルスイッチのプラットフォーム + テナント + エージェントを確認する
03 スコーププロバイダー + モデル + エンドポイントを強制する
04 通話前に予算トークン + 費用を予約する
PassControl は、実際のキーを実行中に挿入し、呼び出しをプロキシし、応答をストリーミングし、監査を記録します。
これは合成された応答であり、 [demo] と明確にマークされています。周囲のパイプラインは実際のものです。チャレンジ署名、就労ビザ、スコープ、予算、監査、デモ テナント キル スイッチなどです。
一度実行してスイッチを準備し、同じコールを再度実行します。解除して元に戻します。
準備ができて。デモパスポートはサーバー側で保持されます。
ゲートウェイで各エージェントを制御します。
エージェント ID、資格情報へのアクセス、支出、取り消し、およびすべての管理対象コールの記録に対する 1 つの検査可能なポリシー境界。
実際の OpenAI および Anthropic 認証情報は、リクエストがポリシーを通過した後にのみゲートウェイ内で解決されます。これらはエージェントのランタイムには決して入りません。
エージェントごとに両方の制限を設定します。 PassControl は、プロバイダーが呼び出す前にアトミックにバジェットを予約します。

施行しても支出を阻止できる場合。
すべてのプロバイダー キーをローテーションしたり、実行中のすべてのエージェント プロセスにアクセスしたりせずに、ゲートウェイで新しいリクエストを停止します。
エージェントおよびパスポートごとのコールを、プロバイダー、モデル、ステータス、トークンの使用状況、コスト、遅延、リクエスト ID とともに 1 つのオペレーター ビューで検査します。
すでに使用している OpenAI または Anthropic SDK はそのままにしておきます。ベース URL を再ポイントし、実際のプロバイダー資格情報の代わりに就労ビザを渡します。
PassControl を Claude Desktop、Cursor、または Claude Code の MCP サーバーとして実行します。同じスコープ、予算、監査、キル スイッチが引き続き適用されます。
エージェントではなくルートを変更します。
すでに使用している SDK と呼び出し形状を保持します。ベース URL を PassControl に指定し、SDK が API キーを必要とする短期就労ビザを指定します。
// 同じ SDK。管理された資格情報パス。
const client = new OpenAI( {
BaseURL: "https://your-gateway/api/v1/openai" ,
apiKey: workVisa、
} );
const response = await client.chat.completions.create( {
モデル: "あなたのモデル" 、
メッセージ、
} ); MCP パスコントロール mcp 同一ポリシー境界 05 / FAQ
PassControl とは何か、何を保護し、何を保護しないのかを簡潔にまとめたもので、マーケティング的な表現はありません。
PassControl は、AI エージェント向けのソースから利用可能な ID および認証情報のゲートウェイです。 OpenAI または Anthropic キーをエージェント内に置く代わりに、各エージェントは暗号化 ID と有効期限が短いスコープ付きトークンを取得し、ゲートウェイはリクエストがポリシーを通過した後にのみ実際のキーを挿入します。
いいえ、エージェントは署名のみの Ed25519 パスポートを保持しており、短期就労ビザを発行しています。ゲートウェイは実際のプロバイダー キーをボールトから解決し、それを実行中に挿入し、呼び出しをプロキシします。キーがエージェント ランタイムに入力されることはありません。
ビジネス ソース ライセンス 1.1 に基づいてソースから入手できます。完全に動作するコアは自由に検査および自己ホストできますが、

n OSI オープンソース ライセンス。この計画はオープンコアです。有料ホスティングと責任レイヤーは後から追加されます。
これは初期 (v0.4.x) で、単独で構築されており、まだ独立して監査されていません。最初に非クリティカル キーに対して実行してください。これはセキュリティを第一に構築されています (すべてのテーブルの RLS、単一のサービス ロールのみの復号化パス、追加専用の監査ログ、テナント分離テスト) ですが、テスト対象で慎重であることと、監査済みであることは同じではありません。
現在の OpenAI、Anthropic、Groq、Mistral、Togetter、DeepSeek です。これはドロップイン ゲートウェイであるため、既存の SDK を保持し、そのベース URL を PassControl に指定するだけです。
これらは、ルーティング、キャッシュ、共有キーの背後にある可観測性に重点を置いています。 PassControl は、エージェントごとの暗号化 ID、機能範囲、エージェントごとの予算、即時失効を中心としており、それらと並行してドロップインで実行されます。
いいえ - エージェントはビザを保持していません。ローカル サイドカーはミントと自動リフレッシュを行います (そして 401 では即座に再ミントします)。そのため、取り消しはほぼ瞬時に行われながら、数時間のセッションがタスクの途中でタイムアウトすることはありません。単一の長いストリーミング呼び出しは、開始時に 1 回検証され、関係なく終了します。
エージェントには ID が必要です。キーには距離が必要です。
独自のインフラストラクチャ上で完全な PassControl コアを実行し、資格情報パスのすべての行を検査します。
初期 ( v0.4.x )、オープンに組み込まれており、まだ独立して監査されていません。最初に非クリティカル キーに対して実行します。セキュリティポリシー ↗
ver · tias ソースは BSL 1.1 で利用可能 · セルフホストは無料
© 2026 Vertias ЕООД · ブルガリア、ソフィア · hello@vertias.eu

## Original Extract

Source-available identity and credential gateway for AI agents. Enforce scope, budgets, revocation, and audit without giving agents real provider keys.

PassControl — Keep real API keys out of AI agents Skip to content ver · tias / PassControl How it works Live demo Capabilities Source-available PassControl / Agent credential gateway
Your AI agents should never hold your real API keys.
Give agents cryptographic identity—not the credentials that pay for their calls. PassControl verifies each agent, enforces scope and budget, injects the vaulted provider key, and proxies the request. The key never reaches the agent.
Identity travels · provider secrets do not
Identity crosses the boundary. Secrets do not.
The passport only signs. A short-lived work-visa carries identity and policy into a gateway that checks every request before resolving a provider credential.
The private Ed25519 key signs a one-time challenge locally. It is never sent on the wire.
01 Verify visa identity + expiry
02 Check kill switch platform + tenant + agent
03 Enforce scope provider + model + endpoint
04 Reserve budget tokens + cost before call
PassControl injects the real key in-flight, proxies the call, streams the response, and records the audit.
This is a synthesized response, clearly marked [demo] . The surrounding pipeline is real: challenge signing, work-visa, scope, budget, audit, and the demo tenant kill switch.
Run once, arm the switch, then run the same call again. Disarm to restore it.
Ready. The demo passport is held server-side.
Control each agent at the gateway.
One inspectable policy boundary for agent identity, credential access, spend, revocation, and the record of every governed call.
Real OpenAI and Anthropic credentials are resolved inside the gateway only after a request passes policy. They never enter the agent runtime.
Set both limits per agent. PassControl reserves budget atomically before the provider call, when enforcement can still prevent spend.
Stop new requests at the gateway without rotating every provider key or reaching into every running agent process.
Inspect calls per agent and passport, with provider, model, status, token usage, cost, latency, and request identity in one operator view.
Keep the OpenAI or Anthropic SDK you already use. Re-point its base URL and pass a work-visa instead of a real provider credential.
Run PassControl as an MCP server for Claude Desktop, Cursor, or Claude Code—the same scope, budgets, audit, and kill switch still apply.
Change the route, not the agent.
Keep the SDK and call shape you already use. Point the base URL at PassControl and supply the short-lived work-visa where the SDK expects an API key.
// Same SDK. Governed credential path.
const client = new OpenAI( {
baseURL: "https://your-gateway/api/v1/openai" ,
apiKey: workVisa,
} );
const response = await client.chat.completions.create( {
model: "your-model" ,
messages,
} ); MCP passcontrol mcp same policy boundary 05 / FAQ
The short version of what PassControl is, what it protects, and what it does not — no marketing gloss.
PassControl is a source-available identity and credential gateway for AI agents. Instead of putting your OpenAI or Anthropic key inside an agent, each agent gets a cryptographic identity and a short-lived, scoped token — and the gateway injects the real key only after the request passes policy.
No. The agent holds a sign-only Ed25519 passport and mints a short-lived work-visa; the gateway resolves the real provider key from a vault and injects it in-flight, then proxies the call. The key never enters the agent runtime.
It's source-available under the Business Source License 1.1 — the full working core is free to inspect and self-host, but it is not an OSI open-source license. The plan is open-core: paid hosting and an accountability layer come later.
It's early (v0.4.x), built solo, and not yet independently audited — run it against a non-critical key first. It is built security-first (RLS on every table, a single service-role-only decrypt path, an append-only audit log, tenant-isolation tests), but test-covered and careful is not the same as audited.
OpenAI, Anthropic, Groq, Mistral, Together, and DeepSeek today. Because it is a drop-in gateway, you keep your existing SDK and just point its base URL at PassControl.
Those center on routing, caching, and observability behind a shared key. PassControl centers on per-agent cryptographic identity, capability scoping, per-agent budgets, and instant revocation — and it runs drop-in alongside them.
No — the agent does not hold the visa. A local sidecar mints and auto-refreshes it (and re-mints instantly on a 401), so a multi-hour session never times out mid-task while revocation stays near-instant. A single long streaming call is verified once at the start and finishes regardless.
Your agents need identity. Your keys need distance.
Run the complete PassControl core on your own infrastructure and inspect every line in the credential path.
Early ( v0.4.x ), built in the open, not yet independently audited — run it against a non-critical key first. Security policy ↗
ver · tias Source-available under BSL 1.1 · free to self-host
© 2026 Vertias ЕООД · Sofia, Bulgaria · hello@vertias.eu

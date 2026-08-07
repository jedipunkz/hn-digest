---
source: "https://zerker.ai/"
hn_url: "https://news.ycombinator.com/item?id=49211867"
title: "Zerker AI Gateway: route, guard and charge"
article_title: "Zerker Gateway — every agent call, through a door you own"
author: "nader"
captured_at: "2026-08-07T15:44:42Z"
capture_tool: "hn-digest"
hn_id: 49211867
score: 1
comments: 0
posted_at: "2026-08-07T15:22:26Z"
tags:
  - hacker-news
  - translated
---

# Zerker AI Gateway: route, guard and charge

- HN: [49211867](https://news.ycombinator.com/item?id=49211867)
- Source: [zerker.ai](https://zerker.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T15:22:26Z

## Translation

タイトル: Zerker AI ゲートウェイ: ルート、ガード、チャージ
記事のタイトル: Zerker Gateway — すべてのエージェントの電話は、あなたが所有するドアを通して
説明: Zerker は、エージェント トラフィックの前にある 1 つの Go バイナリです: カタログ、ルート、ガード、監視、およびすべての通話の価格。どこでもセルフホスト。決して鍵を握ることはありません。

記事本文:
コンテンツにスキップ
ザーカー・ゲートウェイ
特長
5つの面
エージェントカタログ →
テナントを対象とした記録システム。 1 つおきのサーフェスが解決されます。
エージェントを ID で識別します。
ルーティングとプロキシ →
トランザクションおよびストリーミング呼び出し、SSRF で保護された認証情報
リクエスト時に注入されます。
MCP ネイティブ トランスポート →
ストリーミング可能な HTTP アップストリーム、ツールレベルの可視性を備えた汎用プロキシ
見えない。
可観測性 →
記録されたすべての通話 - カウント、エラー率、遅延、TTFT パーセンタイル、
API 経由で読み取り可能。
x402 ゲート →
転送する前に支払い承認を確認してください。和解はありません、いいえ
サーバーが保持するキー。
すべてのエージェントは、あなたが所有するドアを通して電話をかけます。
Zerker はエージェントのトラフィックの前に座って、生の通話 (プレーン HTTP または
MCP — カタログ化、ルーティング、保護、監視、課金ができるもの。あなたは走ります
それ。決して鍵を握ることはありません。
$ git clone https://github.com/zerkerlabs/gateway.git
$ cd ゲートウェイ && dev-auth を作成する
コピー
クイックスタート →
リクエストを見る
サインアップはありません。アカウントがありません。何も電話をかけられません。
一つお願いです。チェックポイントは6つ。何もスキップされませんでした。
呼び出しはゲートウェイに入り、呼び出しレコードとして出力されます。以下の各ステージは、
リクエスト時にプロセス内で強制されます。一度設定する必要はなく、永久に信頼されます。
ゲートウェイは、OIDC 発行者と対象者なしでは起動を拒否します。警告ではありません —
プロセスは終了します。
$ ZERKER_OIDC_ISSUER= 実行する
致命的: OIDC 発行者と対象者が必要です
バイナリにはバイパス フラグも開発モード ショートカットもありません。地域発展
代わりに、 make dev-auth によって起動される使い捨てのモック発行者を取得します。
ホスト型ゲートウェイでは構造的に提供できないものが 2 つあります。
他の誰かが次の四半期に出荷する機能はありません - どこにあるかの影響
ソフトウェアが実行され、誰がキーを保持しているか。
1 つのバイナリ、1 つのプロセス、ハードウェア。
オンプレミス、独自の VPC 内、またはエアギャップ

編JVM なし、Python ランタイムなし、なし
サイドカー、その隣にインストールするエージェントはありません。必要に応じて Postgres を使用してバックアップします
エージェントは再起動後も生き残ることができます。
支払いゲートは、支払いを転送する前に、有効な承認が存在することを確認します。
電話する。それを行うための秘密鍵を保持していません。後で実際に使いたい場合は、
オンチェーンで解決する場合は、自分のガスキーを使用して、自分でファシリテーターを実行します。
独自のハードウェア。
鍵は決して​​離れることはありません。 Zerker は発信者が提示する承認を検証します
そして前進するか、しないか。引き渡すものは何もなく、侵害するものも何もありません。
セルフホストします。それを請求してください。キーを持たないでください。
この 3 つすべてを行う人は他にいません。
ContextForge と Lunar はセルフホストですが、通話料金を請求することはできません。 Cloudflareの料金
ただし、自己ホストすることはできません。 MCPay は、TypeScript と決済の両方を実行します
ファシリテーターサービスに委任されます。それがフィールド全体です。
動作: vs MCPay 、
vsクラウドフレア、
対 ContextForge 。
私たちは、このページが売りにしている正確なメカニズムの背後に答えを置いています。を提示する
それを読むための支払い承認。 (これはデモです。何も請求されません。
財布は関係ありません。）
X-支払いが必要: x402
スキーム: 正確
ネットワーク: ベース
金額: 0.00 USDC (デモ)
リソース: /価格
現在の認可
⌁ 支払い保留中
ゲートウェイはこの応答を保持しています
OSS が検証してキャプチャします。
コマーシャルは徴収、請求、管理を行います。
あなたが今通過したゲートはオープンソースの半分であり、決済も同様です
その背後にあるサーバー — パス全体を自分で実行できます。私たちが販売しているものはそうではありません
お金を取る許可。鍵を保持する部分を実行する必要はありません。
1 つのコマンドで起動します。あなたのものは2つです。
ローカル開発は、ゲートウェイと一緒に使い捨てのモック発行者を起動し、書き込みます
あなたはベアラートークンです。独自の IdP での運用ポイント — Auth0、Okta、Google、
すでに実行しているものは何でも。
3人のうちの1人

RT。それぞれ単独でも役に立ちます。
ザーカーはトラフィックを運びます。他の 2 つは信頼と記憶を運びます。

## Original Extract

Zerker is one Go binary in front of your agent traffic: catalog, route, guard, watch, and price every call. Self-hosted anywhere. It never holds a key.

Skip to content
Zerker Gateway
Features
The five surfaces
Agent Catalog →
The tenant-scoped system of record. Every other surface resolves an
agent by its ID.
Routing & Proxy →
Transactional and streaming invocation, SSRF-guarded, credentials
injected at request time.
MCP-Native Transport →
Streamable-HTTP upstreams, with tool-level visibility a generic proxy
cannot see.
Observability →
Every call recorded — count, error rate, latency and TTFT percentiles,
readable over the API.
The x402 Gate →
Verify a payment authorization before forwarding. No settlement, no
server-held key.
Every agent call, through a door you own.
Zerker sits in front of your agent traffic and turns raw calls — plain HTTP or
MCP — into something you can catalog, route, guard, watch, and charge for. You run
it. It never holds a key.
$ git clone https://github.com/zerkerlabs/gateway.git
$ cd gateway && make dev-auth
Copy
Quickstart →
Watch a request
No signup. No account. Nothing phones home.
One request. Six checkpoints. Nothing skipped.
A call enters the gateway and leaves as an invocation record. Every stage below is
enforced in-process, at request time — not configured once and trusted forever.
The gateway refuses to start without an OIDC issuer and audience. Not a warning —
the process exits.
$ ZERKER_OIDC_ISSUER= make run
fatal: OIDC issuer and audience are required
There is no bypass flag and no dev-mode shortcut in the binary. Local development
gets a throwaway mock issuer instead, booted by make dev-auth .
Two things a hosted gateway structurally can't give you.
Not features someone else will ship next quarter — consequences of where the
software runs and who holds the key.
One binary, one process, your hardware.
On-prem, inside your own VPC, or air-gapped. No JVM, no Python runtime, no
sidecar, no agent to install next to it. Back it with Postgres when you want
agents to survive a restart.
The payment gate checks that a valid authorization exists before it forwards a
call. It does not hold a private key to do that. If you later want to actually
settle on-chain, you run the facilitator yourself — with your own gas key, on
your own hardware.
The key never leaves. Zerker verifies the authorization a caller presents
and forwards, or doesn't. Nothing to hand over, nothing to breach.
Self-host it. Charge for it. Hold no keys.
Nobody else does all three.
ContextForge and Lunar self-host but can't charge for a call. Cloudflare charges
but cannot be self-hosted. MCPay does both — in TypeScript, with settlement
delegated to a facilitator service. That's the whole field, and we show our
working: vs MCPay ,
vs Cloudflare ,
vs ContextForge .
We put the answer behind the exact mechanism this page is selling. Present a
payment authorization to read it. (It's a demo. Nothing is charged, and there is
no wallet involved.)
X-Payment-Required: x402
scheme: exact
network: base
amount: 0.00 USDC (demo)
resource: /pricing
Present authorization
⌁ withheld pending payment
the gateway is holding this response
OSS verifies and captures.
Commercial collects, bills, and governs.
That gate you just walked through is the open-source half, and so is the settle
server behind it — you can run the whole path yourself. What we sell is not
permission to take money. It's not having to run the part that holds the key.
Up in one command. Yours in two.
Local development boots a throwaway mock issuer alongside the gateway and writes
you a bearer token. Production points at your own IdP — Auth0, Okta, Google,
whatever you already run.
One of three parts. Each useful alone.
Zerker carries the traffic. The other two carry the trust and the memory.

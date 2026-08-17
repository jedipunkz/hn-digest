---
source: "https://www.getnexum.dev/blog/nexum-004-fabian-williams"
hn_url: "https://news.ycombinator.com/item?id=49329849"
title: "Found a duplicate-execution bug in an MCP tool via static analysis (no LLM)"
article_title: "Case Study: NEXUM-004 IdempotencyMissing on a real MCP spec — Nexum"
image: "https://getnexum.dev/static/img/og-image.png"
author: "mbelckadi"
captured_at: "2026-08-17T13:33:38Z"
capture_tool: "hn-digest"
hn_id: 49329849
score: 1
comments: 0
posted_at: "2026-08-17T12:35:52Z"
tags:
  - hacker-news
  - translated
---

# Found a duplicate-execution bug in an MCP tool via static analysis (no LLM)

- HN: [49329849](https://news.ycombinator.com/item?id=49329849)
- Source: [www.getnexum.dev](https://www.getnexum.dev/blog/nexum-004-fabian-williams)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T12:35:52Z

## Translation

タイトル: 静的解析により MCP ツールで重複実行のバグが見つかりました (LLM なし)
記事のタイトル: ケーススタディ: NEXUM-004 実際の MCP 仕様で冪等性が欠落している — Nexum
説明: パブリック MCP ストアフロントで、変異ツール上で NEXUM-004 IdempotencyMissing が公開されました。ここ

記事本文:
実際の仕様に関する NEXUM-004: 公開 MCP ストアフロントが冪等性不足の発見を 1 日で解決した方法
これはデモ仕様ではありません。 Fabian Williams ( @fabianwilliams ) は、MCP 呼び出し可能な公開ストアフロントを構築しました。実際の購入フローは、fabswill.com で公開されています。 Nexum のルールを彼の仕様に反して手作業で実行したところ、そのうちの 1 つである NEXUM-004 が失敗しました。彼自身の言葉と彼自身のプルリクエストで何が起こったのかを以下に示します。
Fabian のストアフロントは、同一のプロトコル中立的なエンドポイントを生成する 1 つの MCP 呼び出し可能なエンドポイントを提供します。
呼び出し元に関係なく監査証跡の受信 - 彼は、ホストされたクロード デスクトップと
完全にローカルの Qwen3.6 27B をラップトップ上でオフラインで実行し、同じ 6 回のランタイム チェックを実行し、同じ結果を出しました。彼は
X のスレッドとして投稿しました
2026 年 5 月 16 日。
Mehdi は翌日、「実行時チェックはしっかりしているように見えましたが、仕様はどうなったでしょうか」と返信しました。
クライアントに届く前に？受信 URL パターン - 両方を満たす 1 つのエンドポイント
セキュリティ監査と財務請求の観点は、爆破半径の観点から興味深いものでした。
2 人のコンシューマにサービスを提供するエンドポイントは、エージェントが両方に対して呼び出すことができる 1 つのエンドポイントでもあります。
ファビアンはその申し出を受け入れた。 Nexum スキャナーのバイナリはまだ公開されていなかったので、彼は自分で調べてみました。
Purchase_free_bundle 仕様をマニフェストの 5 つの列挙ルールに対して手動で検証します。
4 名が合格または応募しませんでした。 1つは失敗しました。
発見結果 — NEXUM-004 冪等性の欠落
高
Purchase_free_bundle は、変化する MCP ツールです。レシートを作成し、連絡先を更新/挿入します。
Brevo は、JOSE 署名のダウンロード トークンを作成し、フルフィルメント電子メールを送信します。ツールの
inputSchema は Idempotency-Key を公開しませんでした。再試行するエージェント
タイムアウト — ほとんどの LLM-SDK クライアント ポリシーのデフォルトの再試行動作 — は、
電子メールをバンドルして二重に送信します。
ランタイムガードがそれを捕らえなかった理由
ファビアンの店

すでにランタイム ガバナンスが導入されており、1 時間あたり 5 リクエストのレート制限が設けられています。
IP、および 1 日あたり 5 ドルのコスト上限。どちらもこのバグでは起動しません。同じ時間内に 2 回の再試行
どちらも、最初の試行ではレート制限の上限を下回って問題なく到達します。無料バンドルの重複発行には費用がかかります
一日の上限には程遠い。ファビアン氏が言うように、それはまさにネクサムの多層防御ギャップです。
フレーミングによる予測: ランタイム ガードと静的仕様レビューにより、さまざまなクラスの障害が検出されます。あ
スキーマ自体の仕様レベルのチェックにより、どのリクエスト量とコストベースのガードが検出されるか
構造上見えません。
Fabian は修正を同日に出荷し、2026 年 5 月 17 日に統合しました。
スキーマ。 idempotency_key (オプション、8 ～ 128 文字、パターン [A-Za-z0-9._-] ) を MCP ツールの inputSchema に追加し、運用エージェントに推奨するためにツールの説明が更新されました。
HTTPヘッダー。 /api/a2a/mcp と /api/a2a/purchase は両方とも、標準の Idempotency-Key リクエスト ヘッダーも受け入れます。両方が存在する場合、ツールの引数が優先されます。
ストレージ。 Azure Table Storage をサポートする新しい idempotency.ts モジュール。 PartitionKey と RowKey はハッシュ値であり、生のキーや電子メールは保存されません。
ストライプ スタイルの再生セマンティクス。有効なキーを持つすべてのリクエストで、入力検証が実行される前にテーブルがチェックされます。ヒットすると、元のレシートが Blob Storage からロードされ、そのままの状態で返されます。
成功のみがキャッシュされます。障害の受信は保存されないため、一時的なダウンストリーム障害 (Brevo ブリップなど) がキャッシュされたエラーによって悪影響を受けることはありません。エージェントの再試行により、意図したとおりにフローが再実行されます。
フェールオープン。キャッシュ ヒットの欠落や書き込みの失敗は、どちらもエラーとして表面化されるのではなく、無視されます。ここでの冪等性は安全ベルトであり、新たな障害モードではありません。
ファビアンはこれらをごまかさず、PR 自体に文書化しました。
同時リクエスト

h 同じキーがキャッシュを見逃して両方のフローを実行する可能性があります。テーブルでは最後の書き込みが優先され、2 つのレシートが BLOB に書き込まれます。 Stripe スタイルの進行中ロック (オプティミスティック同時実行での処理に状態を設定、衝突時に 409 を返す) は、計画された v2 強化です。
テーブル TTL はまだありません。マッピングは手動で取得されるまで存続します。デモボリュームではコストは無視できます。過去 10,000 行以上に再訪問のフラグが立てられています。
「@MBelckadi の信用を与えてください。あなたの信頼マニフェストはクリーンな静的分析スキーマです。」
ランタイム ガード (レート制限、コスト上限) によって、動作の頻度と量が制御されます。
設計どおりに 1 回だけ実行された単一の呼び出しが安全に再試行できるかどうかについては、何も述べていません。
最初のエージェントがツールを呼び出す前に、仕様を静的に読み取ることで、そのクラスのリスクをキャッチします。
ランタイム ガバナンスに代わるものではなく、ランタイム ガバナンスの上に多層防御を構築します。それがすべての前提です
Nexum Certの後ろにあります。
MCP サーバーはエージェントに対して安全ですか?
OpenAPI 仕様をアップロードし、Nexum Cert + PDF レポートを数秒で取得します。
無料。アカウントは必要ありません。

## Original Extract

A public MCP storefront exposed NEXUM-004 IdempotencyMissing on a mutating tool. Here

NEXUM-004 on a real spec: how a public MCP storefront closed an IdempotencyMissing finding in a day
This is not a demo spec. Fabian Williams ( @fabianwilliams ) built a public, MCP-callable storefront — a real purchase flow, live on fabswill.com . When Nexum's rules were walked against his spec by hand, one of them failed: NEXUM-004. Here's what happened, in his own words and his own pull request.
Fabian's storefront serves one MCP-callable endpoint that produces an identical, protocol-neutral
audit-trail receipt regardless of caller — he tested it against both hosted Claude Desktop and a
fully-local Qwen3.6 27B running offline on his laptop, same six runtime checks, same result. He
posted it as a thread on X
on May 16, 2026.
Mehdi replied the next day: the runtime checks looked solid, but what did the spec look like
before it reached the client? The receipt-URL pattern — one endpoint satisfying both
the security audit and the finance billing view — was interesting from a blast-radius angle: an
endpoint that serves two consumers is also one endpoint an agent can call for both.
Fabian took the offer. The Nexum scanner binary wasn't public yet, so he walked his
purchase_free_bundle spec against the manifest's five enumerated rules by hand.
Four passed or didn't apply. One failed.
The finding — NEXUM-004 IdempotencyMissing
HIGH
purchase_free_bundle is a mutating MCP tool: it mints a receipt, upserts a contact in
Brevo, mints a JOSE-signed download token, and dispatches a fulfillment email. The tool's
inputSchema exposed no Idempotency-Key . Any agent that retries on
timeout — the default retry behavior in most LLM-SDK client policies — would double-issue the
bundle and double-fire the email.
Why the runtime guards didn't catch it
Fabian's storefront already had runtime governance in place: a 5-requests-per-hour rate limit per
IP, and a $5/day cost ceiling. Neither one fires on this bug. Two retries inside the same hour
both land comfortably under the rate-limit cap on first try. A duplicate free-bundle issue costs
nowhere near the daily ceiling. As Fabian put it, that's exactly the defense-in-depth gap Nexum's
framing predicts: runtime guards and static spec review catch different classes of failure. A
spec-level check on the schema itself catches what request-volume and cost-based guards
structurally cannot see.
Fabian shipped the fix the same day, merged May 17, 2026:
Schema. Added idempotency_key (optional, 8–128 chars, pattern [A-Za-z0-9._-] ) to the MCP tool's inputSchema , with the tool description updated to recommend it for production agents.
HTTP header. Both /api/a2a/mcp and /api/a2a/purchase also accept a standard Idempotency-Key request header — the tool argument wins if both are present.
Storage. A new idempotency.ts module backed by Azure Table Storage. PartitionKey and RowKey are hashed values — no raw key or email is stored.
Replay semantics, Stripe-style. On every request with a valid key, the table is checked before input validation runs. On a hit, the original receipt is loaded from Blob storage and returned exactly as-is.
Only successes are cached. Failure receipts are never stored, so a transient downstream failure (a Brevo blip, for example) doesn't get poisoned by a cached error — the agent's retry re-runs the flow as intended.
Fail open. A missing cache hit or a failed write are both swallowed rather than surfaced as errors. Idempotency is a safety belt here, not a new failure mode.
Fabian documented these in the PR itself, rather than glossing over them:
Concurrent requests with the same key can both miss the cache and both run the flow — last-write-wins on the table, and two receipts get written to Blob. A Stripe-style in-progress lock (set state to processing under optimistic concurrency, return 409 on collision) is the planned v2 hardening.
No table TTL yet — mappings live until manually reaped. Negligible cost at demo volume; flagged for revisit past 10k+ rows.
“Credit where it is due @MBelckadi — your trust-manifest is a clean static-analysis schema.”
Runtime guards — rate limits, cost ceilings — govern behavior : how often and how much.
They say nothing about whether a single call, executed exactly once as designed, is safe to retry.
A static read of the spec catches that class of risk before the first agent ever calls the tool —
defense-in-depth on top of runtime governance, not a replacement for it. That's the whole premise
behind the Nexum Cert.
Is your MCP server agent-safe?
Upload your OpenAPI spec and get a Nexum Cert + PDF report in seconds.
Free. No account required.

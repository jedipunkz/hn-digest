---
source: "https://countersign.network"
hn_url: "https://news.ycombinator.com/item?id=49220224"
title: "Countersign – one kill switch and audit log for AI agents across wallet vendors"
article_title: "Countersign — the kill switch for AI agents that spend money"
author: "screan"
captured_at: "2026-08-08T10:21:41Z"
capture_tool: "hn-digest"
hn_id: 49220224
score: 1
comments: 0
posted_at: "2026-08-08T09:43:46Z"
tags:
  - hacker-news
  - translated
---

# Countersign – one kill switch and audit log for AI agents across wallet vendors

- HN: [49220224](https://news.ycombinator.com/item?id=49220224)
- Source: [countersign.network](https://countersign.network)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T09:43:46Z

## Translation

タイトル: Countersign – ウォレット ベンダー全体の AI エージェント用の 1 つのキル スイッチと監査ログ
記事のタイトル: Countersign — お金を使う AI エージェントのキル スイッチ
説明: お金を使う AI エージェント向けの中立的なクロスベンダー コントロール プレーン。 1 つのポリシー、1 つの凍結、1 つの改ざん防止台帳をすべてのエージェント ウォレット バックエンドに同時に適用します。

記事本文:
副署
始めましょう
強制力
ネットワーク効果
生態系
ドキュメント
ホワイトペーパー
GitHub
npm
始めましょう→
4 本のレールが約 432 ミリ秒でフリーズ — テストネットで実証済み
お金を使うAIエージェントのキルスイッチ。
Countersign は中立的なクロスベンダーのコントロール プレーンです。1 つのポリシー、1 つの凍結、そして 1 つのポリシーです。
改ざん明白な監査台帳 — すべてのエージェント ウォレット バックエンドに一度に。一つではない一つのもの
ウォレットベンダーは、それぞれが独自のレールのみを管理するため、それが可能です。
インストールからキルスイッチが動作するまで数分で完了します。
エントリ ポイントを選択します。それぞれの値から 1 回コピーして貼り付けます。無料、テストネット、アカウントは不要です。
型指定されたクライアント エージェントとオペレーターは、ベンダー間の凍結、支出ガード、およびライブ台帳に接続するために使用します。ブラウザ + ノード。
1 つのポリシーがすべてのバックエンドのネイティブ コントロールにコンパイルされます。支出の前に尋ねてください。 1 回の呼び出しですべてをフリーズします。
"@countersign/sdk" から {CountersignClient} をインポートします。
const cs = new CountersignClient({baseUrl, apiKey });
// この支出は発生する可能性がありますか?
await cs.evaluate({ エージェント ID、金額、資産、会場 });
// キルスイッチ — すべてのウォレット、1 秒未満
cs.freeze() を待ちます;
3
エージェント (MCP) にドロップします。
Claude、Cursor、または任意の MCP クライアントに MCP ツールとして Countersign を追加します (キル スイッチ + スペント ガードの 1 行)。
または、完全なデモをローカルで実行します: pnpm デモ — 3 つのエージェント、3 つのバックエンド、1 つのフリーズ <1 秒。
ウォレットの上にレイヤーを構築します。
上限、許可/拒否リスト、承認しきい値、さらに会場、リスト、交渉条件。単なる数値ではなく、エージェントが取引する市場を管理します。一度書いたもの。
コンパイラはこの制限をすべてのバックエンドのネイティブ コントロールまで下げます。そのため、侵害されたエージェントは、丁寧なリクエストだけでなく、制限を超えることはできません。
③ 1回の凍結、検証可能な監査
1 つのアクションですべてのバックエンドを 1 秒未満で同時に停止し、フェイルクローズされ、すべての試行がサインに達します。

d レジャーは、マークル証明を使用してオフラインで誰でも検証できます。
すべてのコントロールがどこにバインドされているかを正確に確認します。
バインディングが見えないセキュリティ コントロールは、セキュリティ コントロールではありません。副署名はレールごとに出版し、
ネイティブに適用されるかどうか、ベンダーの MPC、エンクレーブ、またはオンチェーン内でポリシー フィールドごとに適用されます。
侵害されたエージェントは、それをバイパスすることはできません。または、Countersign 独自のプリフライト層でもバイパスできません。フリーズはネイティブです
すべてのレール。変化するのはきめ細かいポリシーであり、どれがどれであるかを隠すことはありません。
ポリシー コンパイラーからレールごとにライブで計算されます。現在の行列は次の場所で提供されます。
/enforcement を取得し、ライブ ダッシュボードにレンダリングします。
保護された支出はすべてネットワークの安全性を高めます。
副署名は主に MCP 経由で配布されます。つまり、Claude、Cursor、または自分のエージェントのすべてに配布されます。
スタックはスペント ガードとキル スイッチを選択します。そこを流れる各支出は、クロスレールの安全性を強化します
各ベンダーは独自のレールのみを管理するため、単一のウォレットベンダーが匹敵することのできないレイヤーです。それがフライホイールです。
エージェントは、お金を移動する前に、countersign_request_spend を要求します。人間に対して許可、拒否、保留 - フェールクローズ。
オペレータは、他のエージェントおよびチームに 1 行 MCP サーバーを追加します。キル スイッチはすべてのレールに一度にまたがるようになりました。
1 つのポリシーに基づくより多くのエージェント + 1 つの凍結 = ベンダー間のセーフティ ネットは、エージェント エコノミーが参加するほど価値が高まります。
設計上真実: 伝播は常に付加価値のあるオプトイン サーフェイスです。副署は決して
エージェントの投稿、DM、または採用はそれ自体で行われます。支出の拒否または凍結は、単にピアに同じ保護を提供するだけです。
人間またはエージェントが共有することを選択できます。
バッジを追加します。他のエージェントやチームを連れてきてください。
Countersign がエージェントを保護している場合は、そう言ってください。これは、ユーザー (および他のエージェント) に支出が制限されているという信号を送ります。
統治され、次の方向を指し示す

t オペレータを同じキルスイッチに接続します。 README またはアプリのフッターにバッジをドロップします。
副署名によって管理される支出
← バッジのライブプレビュー
マークダウンバッジ
フリート全体で同じポリシー、同じ凍結。チームメイトをオンボーディングしますか?
app.countersign.network/start を送信してインスタント キーを取得します。
エージェントがお金に触れる前にチームが尋ねる質問。
Countersign はこれらの繰り返しの質問に答えるために構築されており、顧客の見積もりではありません。現在はテストネットのみです。メインネットは第三者の監査に従います。
「午前 2 時にエージェントが不正行為を行った場合、1 つのウォレット内のエージェントだけでなく、すべてのエージェントを停止できますか?」
— クロスベンダーのフリーズは、4 つのレールにわたって約 432 ミリ秒で実証されました
「すべてのエージェントが費やそうとした金額を事後的に正確に証明できますか?」
— 追加専用、ハッシュチェーンされた改ざん防止台帳
「N ベンダー コンソールの代わりに 1 つの支出ポリシーを作成できますか?」
— 各バックエンドのネイティブ コントロールにコンパイルされた 1 つの統合ポリシー
エージェントがすでに使用しているネットワーク。
Countersign は、1 つのポリシー、1 つの凍結、1 つの改ざん防止台帳を通じてそれらを管理します。
すべてのレールを一度に横断します。それぞれのベンダーが独自の管理しか行っていないため、単一のベンダーではこれを行うことはできません。
Countersign は、ユーザーがすでに使用しているウォレットとカード (レール) を管理し、エージェントが構築するツールキットを出荷します。
MPC ウォレット · ネイティブの TX ごとの上限が CDP ポリシーにプッシュされました。
エンクレーブ内 CEL ポリシー · ネイティブの人間による承認のコンセンサス。
バックエンド スマート アカウント · カストディ レベルのハード フリーズ。
バーチャル Visa カード · リアルタイム ASA 承認/拒否。
npx @countersign/mcp — 13 のツール: 任意の MCP クライアントの Kill Switch + Spend Guard。グラマの鍛冶場にある MCP レジストリにあります。
@countersign/sdk · @countersign/api-contract — 型指定されたクライアント + ライブ台帳。
npx @countersign/verify — 台帳エントリがコミットされ、オフラインで、何も信頼されていないことを証明します。
HTTP-402 マシンパを管理する

支払う前に基準を満たしてください。
エージェントが署名する前に、エージェント支払いプロトコルの義務とそれが約束する条件を守ります。
Apache-2.0 でのフロントドア パッケージ。自分で監査してください。
マトリックス内の残りのレイヤー セル (可逆的なオンチェーン KeysManager のフリーズ、ネイティブ コンセンサスの承認、すべての発行者のカード ASA) を閉じて、各保証がバックエンドでエンドツーエンドで適用されるようにします。
オンチェーン アンカーリングがリリースされました。ライブ デモは、署名されたチェックポイント ルートを Base に自律的にアンカーするため、履歴はパブリック チェーン上で目撃され、オフラインの証明検証は 1 つのコマンド ( @countersign/verify ) で実行されます。次に、複数インスタンスと KMS 署名があるため、単一の障害が発生してもフリーズは存続します。
独立したセキュリティ監査後の実際の価値のある運用に加え、Stripe、Airwallex などのウォレットおよびカード ネットワークが構築され、存続します。
副署
ホワイトペーパー
PDF
GitHub
npm
ドキュメント
プライバシー
セキュリティ
Countersign は、ポリシー、凍結、改ざん防止台帳を保持しており、資金を保管することはありません。
現在はテストネットのみです。メインネットはサードパーティのセキュリティ監査に従っています。 © 2026 副署名。 Apache-2.0 オープンコア。

## Original Extract

A neutral, cross-vendor control plane for AI agents that spend money. One policy, one freeze, one tamper-evident ledger — across every agent-wallet backend at once.

Countersign
Get started
Enforceability
Network effect
Ecosystem
Docs
Whitepaper
GitHub
npm
Get started →
Four rails frozen in ~432ms — proven on testnet
The kill switch for AI agents that spend money.
Countersign is a neutral, cross-vendor control plane: one policy, one freeze, and one
tamper-evident audit ledger — across every agent-wallet backend at once . The one thing no single
wallet vendor can do, because each only governs its own rail.
From install to a working kill switch in minutes.
Pick your entry point — each is one copy-paste away from value. Free, testnet, no account required.
The typed client agents and operators use to wire in the cross-vendor freeze, spend guard, and live ledger. Browser + Node.
One policy compiles to every backend's native controls. Ask before a spend; freeze everything in one call.
import { CountersignClient } from "@countersign/sdk" ;
const cs = new CountersignClient({ baseUrl, apiKey });
// may this spend happen?
await cs.evaluate({ agentId, amount, asset, venue });
// kill switch — every wallet, < 1s
await cs.freeze();
3
Drop it into your agent (MCP)
Add Countersign as MCP tools in Claude, Cursor, or any MCP client — the kill switch + spend guard, one line.
Or run the full demo locally: pnpm demo — 3 agents, 3 backends, one freeze <1s.
Build the layer above the wallets.
Caps, allow/deny lists, approval thresholds — plus venues, listings, and negotiated terms . It governs markets an agent transacts in, not just a number. Written once.
The compiler lowers it to every backend's native controls — so a compromised agent can't exceed the cap, not just a polite request.
③ One freeze, verifiable audit
A single action stops every backend concurrently in <1s, fail-closed — and every attempt lands in a signed ledger anyone can verify offline with a Merkle proof.
See exactly where every control binds.
A security control whose binding you can't see isn't one. Countersign publishes, per rail and
per policy field, whether it's enforced natively — inside the vendor's MPC, enclave, or on-chain, where
a compromised agent can't bypass it — or at Countersign's own pre-flight layer. The freeze is native on
every rail. What varies is fine-grained policy, and we never hide which is which.
Computed live per rail from the policy compiler — the current matrix is served at
GET /enforcement and rendered in the live dashboard .
Every guarded spend makes the network safer.
Countersign is distributed primarily over MCP — so any agent in Claude, Cursor, or your own
stack picks up a spend guard and a kill switch. Each spend that flows through it strengthens a cross-rail safety
layer no single wallet vendor can match — because each vendor only governs its own rail. That's the flywheel.
An agent asks countersign_request_spend before it moves money. Allowed, denied, or held for a human — fail-closed.
Operators add the one-line MCP server to their other agents and their team. The kill switch now spans every rail at once.
More agents under one policy + one freeze = a cross-vendor safety net that grows more valuable the more of the agent economy joins it.
Truthful by design: propagation is always a value-adding, opt-in surface. Countersign never makes an
agent post, DM, or recruit on its own — a denied spend or a freeze simply offers peers the same protection, which a
human or agent can choose to share.
Add the badge. Bring your other agents and your team.
If Countersign guards your agent, say so — it signals to users (and other agents) that spending is
governed, and it points the next operator to the same kill switch. Drop the badge in your README or app footer.
Spending governed by Countersign
← live preview of the badge
Markdown badge
Same policy, same freeze, across your whole fleet. Onboarding a teammate?
Send them app.countersign.network/start for an instant key.
The questions a team asks before agents touch money.
These are the recurring asks Countersign is built to answer — not customer quotes. It's testnet-only today; mainnet follows a third-party audit.
“If an agent goes rogue at 2am, can I stop all of them — not just the ones on one wallet?”
— the cross-vendor freeze, proven across four rails in ~432ms
“Can I prove, after the fact, exactly what every agent tried to spend?”
— the append-only, hash-chained, tamper-evident ledger
“Can I write one spending policy instead of N vendor consoles?”
— one unified policy compiled to each backend's native controls
The networks your agents already use.
Countersign governs them through one policy, one freeze, and one tamper-evident ledger —
across every rail at once. No single vendor can do that, because each only governs its own.
Countersign governs the wallets & cards you already use — the rails — and ships the toolkit your agents build with.
MPC wallets · native per-tx cap pushed into CDP policy.
In-enclave CEL policy · native human-approval consensus.
Backend smart accounts · custody-level hard freeze.
Virtual Visa card · real-time ASA approve/decline.
npx @countersign/mcp — 13 tools: kill switch + spend guard in any MCP client. On the MCP Registry, Smithery, Glama.
@countersign/sdk · @countersign/api-contract — typed client + live ledger.
npx @countersign/verify — prove any ledger entry is committed, offline, trusting nothing.
Govern the HTTP-402 machine-payment standard before it pays.
Guard an Agent Payments Protocol mandate — and the terms it commits — before the agent signs.
Front-door packages under Apache-2.0. Audit it yourself.
Close the remaining layer cells in the matrix — reversible on-chain KeysManager freeze, native consensus approvals, card ASA on every issuer — so each guarantee is backend-enforced end to end.
On-chain anchoring now ships — the live demo autonomously anchors its signed checkpoint roots to Base, so history is witnessed on a public chain, and offline proof verification is one command ( @countersign/verify ). Next: multi-instance & KMS-signed, so the freeze survives any single failure.
Real-value operation after an independent security audit — plus Stripe, Airwallex & more wallet and card networks graduating from built to live.
Countersign
Whitepaper
PDF
GitHub
npm
Docs
Privacy
Security
Countersign holds policy, freeze, and a tamper-evident ledger — it never takes custody of funds.
Currently testnet-only; mainnet follows a third-party security audit. © 2026 Countersign. Apache-2.0 open core.

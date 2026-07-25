---
source: "https://www.keydris.com/"
hn_url: "https://news.ycombinator.com/item?id=49050972"
title: "How are you authorizing AI agents that call MCP servers?"
article_title: "Keydris : Proof of authority for AI agents"
author: "ahmed89"
captured_at: "2026-07-25T20:06:00Z"
capture_tool: "hn-digest"
hn_id: 49050972
score: 1
comments: 0
posted_at: "2026-07-25T19:57:56Z"
tags:
  - hacker-news
  - translated
---

# How are you authorizing AI agents that call MCP servers?

- HN: [49050972](https://news.ycombinator.com/item?id=49050972)
- Source: [www.keydris.com](https://www.keydris.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T19:57:56Z

## Translation

タイトル: MCP サーバーを呼び出す AI エージェントをどのように認可していますか?
記事タイトル: Keydris : AI エージェントの権限の証明
説明: Keydris は、AI エージェントに権限が与えられている内容の検証可能な証拠を提供するため、企業や機関はアクションが発生する前に範囲、ポリシー、権限を検証できます。

記事本文:
Keydris : AI エージェントの権限の証明 keydris 仕組み
AI エージェントの権限の証明。
Keydris は、AI エージェントに権限が与えられている内容の検証可能な証拠を提供するため、企業や機関はアクションが発生する前に範囲、ポリシー、権限を検証できます。
Keydris ゲートウェイはフローの中心にあり、Keydris アプリで設定したポリシーの範囲にあるキットを各エージェントに発行します。エージェントはすべてのツール呼び出しでそのキットを提示し、キット リーダー (MCP サーバー上にインストールされた軽量ミドルウェア) がゲートウェイでそれをチェックし、サーバーに到達する前に呼び出しを許可または拒否します。
キット : ゲートウェイによって発行 承認済み : アクションが続行 拒否 : キットリーダーでブロックされている理由 Keydris
自律経済のための信頼インフラ -
自律経済のための信頼インフラ -
すでに各機関が抱いている質問。
顧客の見積もりではなく、典型的な役割 — 標準の Keydris は満たすように構築されています。
「これらのアクションのうち、実際に承認されたのはどれですか? 誰によって承認されましたか?」
最高情報セキュリティ責任者
すべてのアクションは、発行者、スコープ、ポリシー、署名などの権限が付加されて到着し、それを実行したエージェントに関係なく、ミリ秒単位で検証できます。
発行者 acme-corp.keydris.id · 有効な署名
「09:00に代理人の権限を取り消した場合、09:01には何が証明されるでしょうか?」
何もない。失効は検証自体の一部です。次のチェックは失敗し、監査証跡にはその時点が正確に示されます。
09:00:03Z · 失効フィードが同期されました
09:00:41Z · アクションの試行 → 拒否 · 記録に署名
「監査人に渡す記録を見せてください。」
署名された一連の決定がそのまま引き渡されます。各記録には、要求、有効なポリシー、チェック、および結果の名前が記載されています。
すべての検証により、署名されたレコード、発行者、AG が生成されます。

ent、スコープ、ポリシー バージョン、決定、タイムスタンプ。監査証跡は、その場のためにまとめられたレポートではありません。これはシステムのネイティブ出力です。
そしてレイヤーはニュートラルなままになります。 Keydris は、モデル、ベンダー、プラットフォーム、レールから独立しています。インフラストラクチャは、いずれかの機関内ではなく、機関間に設置される必要があります。
Keydris は、エージェントに署名付きトークン (KIT) を発行します。エージェントが何をどのくらいの期間実行できるかを示します。デフォルトでは最小限の権限。
承認は、追加専用のハッシュチェーンされたログにバインドされます。エージェントに許可された内容の不変の記録。
執行は取引相手の玄関先で行われます。受信側システムは、エージェントが接続または動作する前にトークンを検証します。
権限は、あらゆる境界を越えて、マシンの速度で引き出されます。取り消されたトークンは、インスタント パーミッションを終了する必要があることの検証に失敗します。
権限が剥奪されるとどうなりますか?
Keydris は私のお金や秘密鍵を保管していますか?
エージェントは独自に行動します。 Keydris は、各アクションが許可されていることを証明し、それが終了する瞬間にその権限を取り消します。
招待のみ · 25 の早期採用枠
© 2026 Keydris, Inc. 全著作権所有
自律的な経済のための信頼インフラ。

## Original Extract

Keydris gives AI agents verifiable proof of what they are authorized to do, so businesses and institutions can validate scope, policy, and authority before action occurs.

Keydris : Proof of authority for AI agents keydris How it works
Proof of authority for AI agents.
Keydris gives AI agents verifiable proof of what they are authorized to do, so businesses and institutions can validate scope, policy, and authority before action occurs.
The Keydris Gateway sits at the heart of the flow: it issues each agent a kit, scoped by the policies you set in the Keydris app. Agents present that kit with every tool call, and the Kit Reader : lightweight middleware installed on top of your MCP server : checks it with the gateway and allows or denies the call before it ever reaches the server.
Kits : issued by the Gateway Authorized : the action proceeds Denied : blocked at the Kit Reader Why Keydris
Trust Infrastructure for the autonomous economy -
Trust Infrastructure for the autonomous economy -
The questions institutions already ask.
Archetypal roles, not customer quotes — the standard Keydris is built to meet.
“Which of these actions was actually authorized , and by whom?”
Chief information security officer
Every action arrives with its permission attached: issuer, scope, policy, signature , verifiable in milliseconds, independent of the agent that carried it.
issuer acme-corp.keydris.id · sig valid
“If we revoke an agent's authority at 09:00, what does it prove at 09:01?”
Nothing. Revocation is part of verification itself — the next check fails, and the audit trail shows exactly when.
09:00:03Z · revocation feed synced
09:00:41Z · action attempted → refused · record signed
“Show me the record you would hand an auditor.”
A signed sequence of decisions, handed over as it stands: each record names the request, the policy in force, the checks, and the outcome.
Every verification produces a signed record, issuer, agent, scope, policy version, decision, timestamp. The audit trail is not a report assembled for the occasion; it is the system's native output.
And the layer stays neutral. Keydris is independent of any model, vendor, platform, or rail. A requirement for infrastructure that sits between institutions rather than inside one of them.
Keydris issues the agent a signed token, the KIT. It states what the agent may do and for how long. Minimum privilege by default.
The authorization is bound to an append-only, hash-chained log. An immutable record of what the agent was permitted to do.
Enforcement is at the counterparty's front door. The receiving system verifies the token before the agent can connect or act.
Authority is pulled at machine speed, across every boundary. A revoked token fails verification the instant permission should end.
What happens when authority is revoked?
Does Keydris hold my money or my private keys?
Agents act on their own. Keydris proves each action was authorized , and revokes that authority the instant it should end.
Invite only · 25 early adopter slots
© 2026 Keydris, Inc. All rights reserved
Trust Infrastructure for the autonomous economy.

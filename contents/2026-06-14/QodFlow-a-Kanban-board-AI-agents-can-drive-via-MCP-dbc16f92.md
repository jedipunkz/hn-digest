---
source: "https://www.qodflow.com"
hn_url: "https://news.ycombinator.com/item?id=48522830"
title: "QodFlow – a Kanban board AI agents can drive via MCP"
article_title: "QodFlow — The AI-Native Kanban Board for Humans + AI Agents"
author: "deimargd"
captured_at: "2026-06-14T01:01:52Z"
capture_tool: "hn-digest"
hn_id: 48522830
score: 2
comments: 0
posted_at: "2026-06-14T00:14:27Z"
tags:
  - hacker-news
  - translated
---

# QodFlow – a Kanban board AI agents can drive via MCP

- HN: [48522830](https://news.ycombinator.com/item?id=48522830)
- Source: [www.qodflow.com](https://www.qodflow.com)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T00:14:27Z

## Translation

タイトル: QodFlow – AI エージェントが MCP 経由で駆動できるカンバン ボード
記事のタイトル: QodFlow — 人間と AI エージェントのための AI ネイティブのかんばんボード
説明: チームと AI エージェントが同じカード (クロード、カーソル、任意の MCP クライアント) を使用するカンバン ボード。いつでも追いついてください。無料プラン、カードなし。
HN テキスト: こんにちは、HN。 QodFlow を構築したのは、サイドバーに取り付けられたチャットボットではなく、AI エージェントが直接作業できるカンバン ボードが欲しかったからです。 MCP サーバーを公開します。エージェントはスコープ付きの取り消し可能なトークンを使用して接続し、claim_job、report_progress、attach_evidence、request_human_decion、resolve_decion を実行できます。すべての通話は人間が見るのと同じジョブ タイムラインに書き込まれ、人間とエージェントのタグが付けられるため、誰 (または何が) 何をしたかを示す 1 つの監査証跡が残ります。エージェントが取り返しのつかない何かに遭遇すると、request_human_decion が開き、カード上に「エージェント・保留中」チップとして表示されます。人間が一度解決すると、エージェントが続行します。残りは通常のカンバンです。強制的なステージ順序 (遷移は記録され、ジャンプ列はありません)、リスク/期限超過にフラグを立てるステージごとの SLA、およびログイン不要の公開ステータス ページを開くカードごとのオプションの QR です。スタック: Next.js 16、Postgres (Neon) + Prisma、NextAuth、Stripe、Vercel。 MCP サーバーは、Web アプリが使用するのと同じ REST API 上の薄い層です。 HN の残酷なまでに正直な意見を気に入っていただけるいくつかの設計呼び出し: サービス プリンシパルとしてのエージェント — 各トークンには人間またはエージェント (isAgent) のフラグが付けられ、同じ権限スコープを共有します。個別のプリンシパル型が適切な呼び出しですか、それとも単にスコープにする必要がありますか? request_human_decion はファーストクラスのボード オブジェクトです。これはオプションとステータスを含む実際の行であり、コメントによる言及ではありません。オーバースペックですか、それともエージェントにメモを残して投票してもらうだけですか?追加専用タイムライン、イベントあたり 280 文字 — カードをスキャン可能に保ちますが、エージェントは強制的に終了します

e.実際の制約は良いものですか、それとも煩わしいものですか?無料プランはアクティブなジョブが 10 件あり、カードはありません。ギャップについて正直に言うと、MCP サーバーは読み書き可能ですがまだストリーミングされておらず、ネイティブ モバイル アプリがなく、パブリック REST ドキュメントは MCP サーフェスを超えて薄いです。 qodflow.com — トレードオフに喜んで参加します。

記事本文:
QodFlow — 人間 + AI エージェント向けの AI ネイティブ カンバン ボード QodFlow 仕組み 価格 価格比較 ブログ ログイン 無料で始める 1 つのカンバン ボード。チームメイト全員。人間かAIか。
あなたのチームと AI エージェント (クロード、カーソル、任意の MCP クライアント) は同じカードを操作します。ボードはステータス更新です。
ライブデモを見る 無料プラン · カードは必要ありません
MCP ネイティブ — クロード、カーソル、OpenCode
デモを見る · 40 代 1 ステージでワークフローを定義 2 チームにジョブを割り当てる 3 フラグの優先順位 + カスタム SLA 4 色 = 予定通り / リスク / 期限超過 5 クライアントのスキャン — ログインなし 6 MCP Stripe チェックアウト経由の AI エージェントのワーク カード
1か所。人間とエージェントの両方がそれを行います。
あなたのチームはボードで決定します。エージェントは MCP を通じて行動します。同じカード、同じタイムライン、同じ監査証跡 - したがって、リーダーは何が起こったのか、誰が（または何を）したかを確認できます。
チャットまたはボードで回答します。両方とも同期します。
エージェントが尋ねます。クロードで返信するか、QodFlow をクリックします。同じ決定、同じタイムライン。和解もなければ、Slack ピンポンもありません。
エージェントはあなたの DM ではなく、カードに投稿します。
claim_job、report_progress、attach_evidence — すべてのアクションはカンバンに反映され、タイムスタンプが付けられ、いつでも取り消すことができるトークンにスコープされます。
取り返しのつかないことをする前に彼らは尋ねます。
request_human_decion は、カード上の「エージェント・保留中」チップを開きます。ボード上でそれを見て、一度決定すると、エージェントは先に進みます。
# Claude デスクトップ、カーソル、OpenCode、任意の MCP クライアント
$ クロード mcp qodflow を追加 -- npx qodflow-mcp
$ export QODFLOW_TOKEN=qf_live_… # トークンごとのスコープ、いつでも取り消し
# エージェントは次のことができるようになりました。
> 請求_ジョブ、レポート_進捗、添付_証拠、
request_human_decion、resolve_decion MCP ドキュメントを読む トークンを取得する Claude Desktop、Cursor、Continue、OpenCode、および任意の MCP 互換クライアントで動作します。トークンスコープ。監査ログに記録されます。取り消し可能。
チームの仕事のために構築

— とそのエージェント — が協力して行動します
クリエイティブおよびデザイン代理店 現場運営 社内運用チーム IT およびサポート チーム 注文および履行運用 修理およびサービスのワークフロー AI エージェントを実行するチーム 仕組み
すべてのジョブが通過する必要があるステップを設定します。各ステージに独自の SLA を与えます。
ジョブをドロップし、前方にドラッグして、緊急のマークを付けます。 SLA ドットはライブのままです。
クライアントはスキャンまたはクリックします。アカウント、アプリ、メールがなくてもステータスを確認できます。
早期導入者 · 100 チームに限定 最初の 100 チームは、プレミアムを 1 年間 50% オフで利用できます。
はい - 永続的です。 10 個のアクティブなジョブ、フル機能、カードなし。有料プランには 30 日間の無料トライアルが追加されます。 FIRST100は何をするのですか？最初の 100 人の顧客は、12 か月間プレミアムが 50% オフになります。 Stripe チェックアウトに FIRST100 を貼り付けます。スケジュールと請求書発行?いや、わざとです。 QodFlow を Calendly / Google Calendar + Stripe / QuickBooks と組み合わせます。重点を置き続ける理由 → AI エージェントはどのようにしてボードに接続するのか? MCP 経由: claude mcp add qodflow -- npx qodflow-mcp 。 Claude Desktop、Cursor、Continue、OpenCode、および任意の MCP 互換クライアントで動作します。各エージェントは、いつでも取り消すことができる独自のスコープ付きトークンを取得します。エージェントは私の承認なしに行動できますか?取り返しのつかないものではありません。エージェントはカードを請求し、進捗状況を報告し、証拠を添付します。ただし、request_human_decion は、チャットまたはボードで人間が応答するまでそれらを一時停止します。すべてのエージェントのアクションにはタイムスタンプが付けられ、監査ログに記録されます。 QR クライアントの追跡はどのように機能しますか?すべてのジョブはステータス URL + QR を取得します。 QR をデバイスまたはレシートに貼り付けます。クライアントはスキャンして、現在のステージと ETA を確認します。アプリもログインも必要ありません。私のデータは安全ですか?転送中は TLS、保存時は暗号化され、米国でホストされます。 Stripe が請求を処理します。カード番号が表示されることはありません。 SOC 2 が進行中です。 CSVエクスポートと完全削除は設定から​​ワンクリックで行えます。いつでもキャンセルできますか?請求ページからいつでも。 N

o 保持コール。気が変わった場合に備えて、データは 30 日間読み取り専用のままになります。なぜ座席ごとではなくワークスペースごとなのでしょうか?座席ごとに協力を罰します。個人ショップは、同じプランの 12 人チームと同じ支払いを行います。価格設定ロジック → まだ質問がありますか?メールでお問い合わせください — 1 営業日以内に返信させていただきます。
あなたのチームにはすでにプロセスがあります。
それを強制する時が来ました。
ワークフローを数分でセットアップし、ほぼリアルタイムで各ジョブの状況を正確に確認できます。
すでにアカウントをお持ちですか?クレジット カード不要 · セットアップは 2 分 · 永久無料プラン
QodFlow 人間と AI エージェントが同じカードを扱うカンバン ボード。 MCP ネイティブ、監査ログ、トークンスコープ。
© 2026 QodFlow.無断転載を禁じます。
チームに合わせた価格設定
無料で始めましょう。より多くの容量が必要な場合はアップグレードしてください。チーム料金 — 座席ごとではありません。
チームをスプレッドシートから解放し、実際のワークフローに無料で移行させます。値札のない構造を必要とする小規模チームに最適です。
アクティブなジョブ 10 件 (紹介で拡大)
月あたり 10 件のエージェント決定 (MCP 対応)
基本的なボードが大きくなりすぎて、ステージごとの SLA 制御、より高い容量、より優れた可視性を必要とするチーム向け。
無制限のエージェント決定 (MCP)
混乱を許容できない運用、つまり人間とエージェントがフル稼働する運用のために構築されています。ジョブを割り当て、クライアントとライブボードを共有し、1 つ屋根の下で複数のチームを管理します。
無制限のエージェント決定 (MCP)
エージェントごとにスコープ指定されたトークン + エージェントの監査証跡
ステージごとの SLA + ジョブごとのカスタム オーバーライド
チームメンバーへの仕事の割り当て
クライアント向け追跡用の QR コード
有料プランの 30 日間の無料トライアル。終了するまで料金はかかりません。いつでもキャンセルできます。

## Original Extract

The kanban board where your team and AI agents work the same cards — Claude, Cursor, any MCP client. Catch up any time. Free plan, no card.

Hi HN. We built QodFlow because we wanted a kanban board our AI agents could work directly — not a chatbot bolted onto the sidebar. It exposes an MCP server. An agent connects with a scoped, revocable token and can: claim_job, report_progress, attach_evidence, request_human_decision, resolve_decision. Every call writes to the same job timeline humans see, tagged human vs agent, so there's one audit trail of who (or what) did what. When an agent hits something irreversible it opens a request_human_decision — surfaces as an "Agent · pending" chip on the card; a human resolves it once and the agent continues. The rest is a normal kanban: enforced stage order (transitions are logged, no jumping columns), per-stage SLAs that flag at-risk/overdue, and an optional QR per card that opens a login-free public status page. Stack: Next.js 16, Postgres (Neon) + Prisma, NextAuth, Stripe, Vercel. The MCP server is a thin layer over the same REST API the web app uses. A few design calls I'd love HN's brutally honest take on: Agent as a service principal — each token is flagged human or agent (isAgent), sharing the same permission scopes. Is a distinct principal type the right call, or should it just be scopes? request_human_decision as a first-class board object — it's a real row with options + status, not a comment mention. Over-engineered vs just letting the agent leave a note and poll? Append-only timeline, 280 chars/event — keeps cards scannable but forces agents to be terse. Good constraint or annoying in practice? Free plan is 10 active jobs, no card. Honest about the gaps: MCP server is read+write but not streaming yet, no native mobile app, public REST docs are thin beyond the MCP surface. qodflow.com — happy to get into the tradeoffs.

QodFlow — The AI-Native Kanban Board for Humans + AI Agents QodFlow How it works Pricing Compare Blog Log in Get started free One kanban board. Every teammate. Human or AI.
Your team and your AI agents — Claude, Cursor, any MCP client — working the same cards. The board is the status update.
See live demo Free plan · no card required
MCP-native — Claude, Cursor, OpenCode
Watch demo · 40s 1 Stages define your workflow 2 Assign jobs to your team 3 Flag priority + custom SLA 4 Color = on-track · risk · overdue 5 Clients scan — no login 6 AI agents work cards via MCP Stripe checkout
One place. Humans and agents both work it.
Your team decides on the board. Your agents act through MCP. Same cards, same timeline, same audit trail — so leaders see what happened and who (or what) did it.
Answer in chat or on the board — both sync.
Agent asks. You reply in Claude, or click in QodFlow. Same decision, same timeline. No reconciliation, no Slack ping-pong.
Agents post on the card, not your DMs.
claim_job · report_progress · attach_evidence — every action lands on the kanban, timestamped, scoped to a token you can revoke at any time.
They ask before doing the irreversible.
request_human_decision opens an "Agent · pending" chip on the card. You see it on the board, decide once, the agent moves on.
# Claude Desktop, Cursor, OpenCode, any MCP client
$ claude mcp add qodflow -- npx qodflow-mcp
$ export QODFLOW_TOKEN=qf_live_… # scope per token, revoke anytime
# Your agent can now:
> claim_job, report_progress, attach_evidence,
request_human_decision, resolve_decision Read the MCP docs Get a token Works with Claude Desktop, Cursor, Continue, OpenCode, and any MCP-compatible client. Token-scoped. Audit-logged. Revocable.
Built for the work your team — and their agents — does together
Creative & design agencies Field operations Internal ops teams IT & support teams Order & fulfillment ops Repair & service workflows Teams running AI agents How it works
Set the steps every job must move through. Give each stage its own SLA.
Drop a job in, drag it forward, mark it urgent. SLA dots stay live.
Clients scan or click — see status without an account, app, or email.
Early-adopter · limited to 100 First 100 teams get 50% off Premium for a full year.
Yes — permanent. 10 active jobs, full features, no card. Paid plans add a 30-day free trial on top. What does FIRST100 do? First 100 customers get 50% off Premium for 12 months . Paste FIRST100 at Stripe checkout. Scheduling and invoicing? No, on purpose. Pair QodFlow with Calendly / Google Calendar + Stripe / QuickBooks. Why we keep it focused → How do AI agents connect to the board? Via MCP: claude mcp add qodflow -- npx qodflow-mcp . Works with Claude Desktop, Cursor, Continue, OpenCode, and any MCP-compatible client. Each agent gets its own scoped token you can revoke anytime. Can agents act without my approval? Not on anything irreversible. Agents claim cards, report progress, and attach evidence on their own — but request_human_decision pauses them until a human answers, in chat or on the board. Every agent action is timestamped and audit-logged. How does QR client tracking work? Every job gets a status URL + QR. Stick the QR on the device or receipt — clients scan and see the current stage and ETA. No app, no login. Is my data safe? TLS in transit, encrypted at rest, US-hosted. Stripe handles billing — we never see card numbers. SOC 2 in progress. CSV export and full delete are one click from settings. Can I cancel anytime? Anytime from your billing page. No retention call. Data stays read-only for 30 days in case you change your mind. Why per-workspace, not per-seat? Per-seat punishes collaboration. A solo shop pays the same as a 12-person team on the same plan. Pricing logic → Still have a question? Email us — we reply within one business day.
Your team already has a process.
It’s time to enforce it.
Set up your workflow in minutes and see — in near real time — exactly where every job stands.
Already have an account? No credit card · 2-minute setup · Free plan forever
QodFlow The Kanban board where humans and AI agents work the same cards. MCP-native, audit-logged, token-scoped.
© 2026 QodFlow. All rights reserved.
Pricing that scales with your team
Start free. Upgrade when you need more capacity. Team pricing — not per seat.
Get your team off spreadsheets and into a real workflow — at no cost. Great for small teams that need structure without the price tag.
10 active jobs (grow it with referrals)
10 agent decisions per month (MCP-ready)
For teams that have outgrown basic boards and need per-stage SLA control, higher capacity, and better visibility.
Unlimited agent decisions (MCP)
Built for operations that can't afford chaos — humans and agents at full capacity. Assign jobs, share live boards with clients, and manage multiple teams under one roof.
Unlimited agent decisions (MCP)
Per-agent scoped tokens + agent audit trail
Per-stage SLA + custom overrides per job
Job assignments to team members
QR codes for client-facing tracking
30 -day free trial on paid plans. No charge until it ends. Cancel anytime.

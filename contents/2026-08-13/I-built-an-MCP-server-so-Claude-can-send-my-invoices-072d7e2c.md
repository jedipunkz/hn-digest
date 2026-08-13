---
source: "https://getholdings.com/mcp"
hn_url: "https://news.ycombinator.com/item?id=49292070"
title: "I built an MCP server so Claude can send my invoices"
article_title: "Connect Your AI — 60-second Setup | Holdings MCP"
author: "sirjrg"
captured_at: "2026-08-13T21:34:54Z"
capture_tool: "hn-digest"
hn_id: 49292070
score: 2
comments: 0
posted_at: "2026-08-13T21:26:28Z"
tags:
  - hacker-news
  - translated
---

# I built an MCP server so Claude can send my invoices

- HN: [49292070](https://news.ycombinator.com/item?id=49292070)
- Source: [getholdings.com](https://getholdings.com/mcp)
- Score: 2
- Comments: 0
- Posted: 2026-08-13T21:26:28Z

## Translation

タイトル: クロードが請求書を送信できるように MCP サーバーを構築しました
記事のタイトル: AI を接続する — 60 秒のセットアップ |ホールディングスMCP
説明: MCP を介して、Claude、ChatGPT、Cursor、Zed、または Windsurf を Holdings に接続します。サインアップ（無料）し、キーを生成し、設定を貼り付けます。最初の請求書は 60 秒以内に送信されます。

記事本文:
AI を接続する — 60 秒のセットアップ | MCPヒーローズを保有）。プリロード
それと以下のメトリクスが一致する「DM Sans Fallback」を追加すると、ヒーローが削除されます
コンテンツ ページで測定されたフォント スワップ リフロー (MAIN の CLS ~0.034)。 --> ) -->
メインコンテンツにスキップ
エージェント 請求書発行 会計 銀行業務 サインイン 口座開設 エージェント セットアップ · 60 秒
AI を接続します。
請求書の発送を開始します。
ホールディングスはMCPを話します。 MCP を話す AI であれば、請求書発行を直接操作できます。以下: 出荷するクライアントの正確な構成。 1 分以内に端から端まで。
workspace.getholdings.com でホールディングス アカウントを作成します。サインアップは瞬時に完了します。トライアルや請求にカードは必要ありません。
2 企業ごとの API キーを生成する
ワークスペース内で、「設定」→「開発者」→「エージェント接続」に移動します。エージェントを操作する会社をスコープとするキーを生成します。キーはいつでもローテーションまたは取り消しできます。
3 設定を AI クライアントに貼り付けます
以下からクライアントを選択してください。ブロックをコピーします。クライアントを再起動します。それでおしまい。
4 AI に請求書の送信を依頼します
「Acme Corp に正味 30 ドルの 4,200 ドルの請求書を作成してください。」 AI が草案を作成します。あなたは承認します。 Holdings は、カード + ACH 支払いリンクを付けて送信します。顧客がいつ支払うかを AI が認識します。
Anthropic のネイティブ MCP クライアント。ファーストクラスのサポート。
~/Library/Application Support/Claude/claude_desktop_config.json (macOS) または %APPDATA%\Claude\claude_desktop_config.json (Windows) に追加します。クロードを再起動します。
{
"mcpサーバー": {
「所蔵」: {
"url": "https://mcp.holdings.io/v1/mcp",
"ヘッダー": {
「認可」: 「所有者 YOUR_HOLDINGS_API_KEY」
}
}
}
カーソル
~/.cursor/mcp.json による MCP サポート。
以下のブロックを ~/.cursor/mcp.json に追加します。カーソルを再開します。所蔵ツールがコンポーザーに表示されます。
{
"mcpサーバー": {
「所蔵」: {
"url": "https://mcp.holdings.io/v1/mcp",
"ヘッダー": {
"認可": "ベアラー YOUR_HO

LDINGS_API_KEY"
}
}
}
ゼッド・ウィンドサーフィン
同じ MCP 構成形状。どちらのエディターもネイティブ サポートを提供します。
Zed: ~/.config/zed/settings.json に、 context_servers を追加します。 Windsurf の場合: [設定] → [MCP] → [サーバーの追加]。同じ URL + Bearer キー。
{
"context_servers": {
「所蔵」: {
「コマンド」: {
"url": "https://mcp.holdings.io/v1/mcp",
"環境": {
"AUTH_TOKEN": "YOUR_HOLDINGS_API_KEY"
}
}
}
}
ChatGPT
カスタム GPT + Holdings コネクタ経由。
ChatGPT の場合: GPT の作成 → 構成 → アクション → URL からインポート:
https://mcp.holdings.io/v1/openapi.json 認証: API Key 、ヘッダー Authorization 、値ベアラー YOUR_HOLDINGS_API_KEY 。保存します。カスタム GPT は、すべての所蔵ツールを使用できるようになりました。
MCP を話すものなら何でも。またはプレーン HTTPS。
https://mcp.holdings.io/v1/mcp または、REST API を直接ヒットします。同じツールで、MCP は必要ありません。
カール https://api.holdings.io/v1/invoices \
-H "承認: ベアラー YOUR_HOLDINGS_API_KEY" \
-H "Content-Type: application/json" 機能
エージェントができること。
そしてそれができないこと。
自然言語の説明から請求書と見積書の草案を作成する
カード + ACH 支払いリンクを含む請求書を (あなたの承認を得て) 送信します
見積書を請求書に変換します。定期的な請求書のスケジュールを設定する
設定可能なリマインダーで期限を過ぎた顧客を追跡する
支払いを受領したことをマークします。支払われたエントリを帳簿に反映させる
レポート用に残高、ステータス、履歴を読み取る
すべてのアクションを監査証跡に記録します
明示的な承認なしに何かを送信する
既存の支払済み請求書または調整されたエントリをタップします
人間の承認なしで資金を移動 (電信、ACH から)
許可されなかったアクセス範囲
MCP を話すものなら何でも。 Claude Desktop と Claude Web はファーストクラスのサポートを提供します。 Cursor、Zed、および Windsurf は MCP サポートをネイティブに出荷します。 ChatGPT は、カスタム GPT + Holdings コネクタを介して接続できます。構築したカスタム エージェントはどれでも使用できます

同じプロトコル。
いいえ。サインアップし、キーを生成し、接続します。請求書の送信と支払いの受け取りは無料で、月額料金もかかりません。 Stripe の標準レート (~3% + $0.30) は、他のカード処理業者と同様、顧客が実際に支払った場合にのみ支払います。完全な会計は、書籍が必要な場合に月額 25 ドルの個別のアドオンです。請求書発行には必要ありません。
自然言語から請求書と見積書の草案を作成し、見積書を請求書に変換し、カード + ACH 支払いリンクで送信し、定期的な請求書のスケジュールを設定し、期限を過ぎた顧客を追跡し、受領済みの支払いをマークし、帳簿への支払済みエントリを調整します。すべての送信にはあなたの承認が必要です。エージェントは、ユーザーの指示なしに既存の記録を削除したり、お金を移動したりすることはできません。監査証跡にはすべてのアクションが記録されます。
workspace.getholdings.com 内で生成される会社ごとの API キー — [設定] → [開発者]。各キーの範囲は 1 つの会社に限定されます。いつでも回転させます。 OAuth フロー (ワンクリック Connect Holdings) は、サポートされているクライアントで使用できます。 API キーは常に機能するフォールバックです。
読み取り: はい - 残高、請求書ステータス、トランザクション、および帳簿データ。書き込み: 送信ごとにあなたの承認を得て、請求書発行と支払いリンクの生成が行われます。資金の流出（電信、ACH 流出）は、依然として、あらゆる物質について人間の承認を得て、同じ規制された銀行フローを経由します。銀行業務は i3 Bank (FDIC メンバー、補償範囲は 300 万ドル) を通じて行われます。エージェント的とは、監視されていないという意味ではありません。
ワークスペース内の Agent Connect 画面 ([設定] → [開発者]) には、エージェントが現在実行できること、まだ有効になっていないこと、および各ギャップを修正するための正確な次のステップなど、ライブ機能マトリックスが表示されます。その他の場合は、support@getholdings.com まで電子メールでお問い合わせください。
AI に本当の仕事を与えましょう。
今すぐ始めましょう。
サインアップしてキーを生成し、設定を貼り付けます。 1 分以内です。
持続可能なビジネスの構築に関する私たちのレター。無料、スパムなし。
私たちの手紙

持続可能なビジネスの構築について。無料、スパムなし。
英国、カナダ、オーストラリア、アイルランド、ニュージーランド、シンガポール、スペイン語の利用規約、プライバシー、Cookie、セキュリティ、電子開示、紹介規約 銀行取引の開示 ビジネス預託契約
ホールディングスは金融テクノロジー企業であり、銀行ではありません。銀行サービスは、FDIC のメンバーである i3 Bank によって提供されます。 Holdings Visa デビット カードは、Visa U.S.A. Inc. のライセンスに基づいて i3 Bank によって発行され、Visa カードが受け入れられる場所ならどこでも使用できます。
口座手数料や国内取引手数料はかかりません。一部の外国取引手数料は、限られた状況で適用される場合があります。
年間利回り (APY) は変動しており、口座開設後に変更される可能性があります。料金は毎月複利計算され、毎月加算されます。
i3 銀行、メンバー FDIC、その他のプログラム銀行の組み合わせにより、預金は合計 300 万ドルまで保証されます。各口座には、銀行ごとに口座所有者あたり最大 250,000 ドルまで個別に保険が掛けられます。保有口座は承認が必要です。利用規約が適用されます。
© 2026 Holdings, Inc. 無断複写・転載を禁じます。

## Original Extract

Connect Claude, ChatGPT, Cursor, Zed, or Windsurf to Holdings via MCP. Sign up (free), generate a key, paste the config. First invoice in 60 seconds.

Connect Your AI — 60-second Setup | Holdings MCP heroes). Preloading
it + the metric-matched 'DM Sans Fallback' below removes the hero
font-swap reflow (CLS ~0.034 on MAIN) measured on content pages. --> ) -->
Skip to main content
Agents Invoicing Accounting Banking Sign in Open account Agent Setup · 60 seconds
Connect your AI.
Start sending invoices.
Holdings speaks MCP. Any AI that speaks MCP can operate your invoicing directly. Below: the exact config for the clients that ship it. Under a minute, end to end.
Create a Holdings account at workspace.getholdings.com . Signup is instant. No trial, no card required for invoicing.
2 Generate a per-company API key
Inside workspace, go to Settings → Developers → Agent Connect . Generate a key scoped to the company you want the agent to operate on. Keys can be rotated or revoked anytime.
3 Paste the config into your AI client
Pick your client below. Copy the block. Restart the client. That’s it.
4 Ask your AI to send an invoice
“Draft an invoice to Acme Corp for $4,200, net 30.” Your AI drafts it. You approve. Holdings sends it with a card + ACH payment link. When the customer pays, your AI knows.
Anthropic’s native MCP client. First-class support.
Add to ~/Library/Application Support/Claude/claude_desktop_config.json (macOS) or %APPDATA%\Claude\claude_desktop_config.json (Windows). Restart Claude.
{
"mcpServers": {
"holdings": {
"url": "https://mcp.holdings.io/v1/mcp",
"headers": {
"Authorization": "Bearer YOUR_HOLDINGS_API_KEY"
}
}
}
} Cursor
MCP support via ~/.cursor/mcp.json .
Add the block below to ~/.cursor/mcp.json . Restart Cursor. The Holdings tools appear in the composer.
{
"mcpServers": {
"holdings": {
"url": "https://mcp.holdings.io/v1/mcp",
"headers": {
"Authorization": "Bearer YOUR_HOLDINGS_API_KEY"
}
}
}
} Zed · Windsurf
Same MCP config shape. Both editors ship native support.
In Zed: ~/.config/zed/settings.json , add context_servers . In Windsurf: Settings → MCP → Add Server. Same URL + Bearer key.
{
"context_servers": {
"holdings": {
"command": {
"url": "https://mcp.holdings.io/v1/mcp",
"env": {
"AUTH_TOKEN": "YOUR_HOLDINGS_API_KEY"
}
}
}
}
} ChatGPT
Via a Custom GPT + the Holdings connector.
In ChatGPT: Create a GPT → Configure → Actions → Import from URL:
https://mcp.holdings.io/v1/openapi.json Authentication: API Key , header Authorization , value Bearer YOUR_HOLDINGS_API_KEY . Save. Your custom GPT can now use every Holdings tool.
Anything that speaks MCP. Or plain HTTPS.
https://mcp.holdings.io/v1/mcp Or hit the REST API directly — same tools, no MCP required:
curl https://api.holdings.io/v1/invoices \
-H "Authorization: Bearer YOUR_HOLDINGS_API_KEY" \
-H "Content-Type: application/json" Capabilities
What your agent can do.
And what it can’t.
Draft invoices and quotes from natural-language descriptions
Send invoices (with your approval) with card + ACH payment links
Convert quotes to invoices; schedule recurring invoices
Chase overdue clients with configurable reminders
Mark payments received; reconcile paid entries into the books
Read balances, statuses, and history for reporting
Log every action to the audit trail
Send anything without your explicit approval
Touch existing paid invoices or reconciled entries
Move money out (wires, ACH out) without human sign-off
Access scopes it wasn’t granted
Anything that speaks MCP. Claude Desktop and Claude web have first-class support. Cursor, Zed, and Windsurf ship MCP support natively. ChatGPT can connect via a custom GPT + the Holdings connector. Any custom agent you build can use the same protocol.
Nope. Sign up, generate a key, connect. Sending invoices and taking payments is free with no monthly fee. You pay Stripe's standard rate (~3% + $0.30) only when a customer actually pays you — same as any card processor. Full accounting is a separate $25/mo add-on for when you want the books; it's not required for invoicing.
Draft invoices and quotes from natural language, convert quotes to invoices, send with card + ACH payment links, schedule recurring invoices, chase overdue clients, mark payments received, and reconcile paid entries into the books. Every send requires your approval. The agent cannot delete existing records or move money without your say-so — the audit trail logs every action.
Per-company API keys generated inside workspace.getholdings.com — Settings → Developers. Each key is scoped to one company. Rotate anytime. OAuth flow (one-click Connect Holdings) is available for supported clients; API key is the fallback that always works.
Reads: yes — balances, invoice status, transactions, and books data. Writes: invoicing and payment link generation happen with your approval on every send. Moving money out (wires, ACH out) still routes through the same regulated banking flows with human approval on anything material. Banking sits underneath via i3 Bank (Member FDIC, coverage to $3M). Agentic doesn’t mean unsupervised.
The Agent Connect screen inside workspace (Settings → Developers) shows a live capability matrix: what your agent can do right now, what’s not yet enabled, and the exact next step to fix each gap. For anything else, email support@getholdings.com.
Give your AI a real job.
Start now.
Sign up, generate a key, paste the config. Under a minute.
Our letter on building businesses that last. Free, no spam.
Our letter on building businesses that last. Free, no spam.
Also in United Kingdom · Canada · Australia · Ireland · New Zealand · Singapore · Español Terms · Privacy · Cookies · Security · Electronic Disclosures · Referral Terms Banking disclosures Business Deposit Agreement
Holdings is a financial technology company and is not a bank. Banking services are provided by i3 Bank, Member FDIC. The Holdings Visa Debit Card is issued by i3 Bank pursuant to a license from Visa U.S.A. Inc. and may be used anywhere Visa cards are accepted.
No account or domestic transaction fees. Some foreign transaction fees may apply in limited circumstances.
Annual Percentage Yield (APY) is variable and subject to change after account opening. Rate is compounded monthly and credited monthly.
Deposits are insured up to $3 million total through a combination of i3 Bank, Member FDIC, and additional program banks. Each account is separately insured up to $250,000 per account holder per bank. Holdings accounts are subject to approval. Terms and conditions apply.
© 2026 Holdings, Inc. All rights reserved.

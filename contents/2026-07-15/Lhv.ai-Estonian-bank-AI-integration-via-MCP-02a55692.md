---
source: "https://lhv.ai/"
hn_url: "https://news.ycombinator.com/item?id=48928679"
title: "Lhv.ai – Estonian bank AI integration via MCP"
article_title: "LHV.ai — Your bank, in your AI assistant"
author: "loh"
captured_at: "2026-07-15T23:50:37Z"
capture_tool: "hn-digest"
hn_id: 48928679
score: 1
comments: 0
posted_at: "2026-07-15T23:41:29Z"
tags:
  - hacker-news
  - translated
---

# Lhv.ai – Estonian bank AI integration via MCP

- HN: [48928679](https://news.ycombinator.com/item?id=48928679)
- Source: [lhv.ai](https://lhv.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T23:41:29Z

## Translation

タイトル: Lhv.ai – MCP を介したエストニアの銀行 AI 統合
記事のタイトル: LHV.ai — あなたの銀行、AI アシスタントに
説明: クロード、カーソル、その他の AI アシスタントをモデル コンテキスト プロトコル経由で LHV アカウントに接続します。 OAuth2 + Smart-ID、読み取り専用。

記事本文:
メイン コンテンツにスキップ JP サインアップ セットアップ ツール セキュリティ API FAQ あなたの銀行、あなたの内部
AIアシスタント。
残高について尋ねたり、先月の取引を調べたり、取引履歴を取得したりできます。一度ログインして質問してください。
> ベータ版を始めましょう。
3つのステップで接続します。
MCP は、サポート ツールのリストが増えているオープン プロトコルです。セットアップには約 2 分かかります。
AI ツールの設定で、新しい MCP サーバーを追加し、LHV サーバーのアドレスを貼り付けます。
近日公開 AI ツールのマーケットプレイスで LHV を直接見つけて、手動セットアップなしで接続します。
1 コネクタ設定を開きます。claude.ai で、[カスタマイズ] → [コネクタ] に移動し、[追加] → [カスタム コネクタの追加] をクリックします。
2 LHV コネクタを追加します。LHV という名前を付け、サーバー URL を貼り付けます。 https://mcp.lhv.ai/mcp
3 接続 「追加」をクリックし、「接続」をクリックして LHV にサインインします。
claude.ai でセットアップすると、デスクトップ アプリやモバイル アプリでも利用できるようになります。無料は 1 つのカスタム コネクタに制限されます。
$ クロード mcp 追加 lhv \
--transport http \
https://mcp.lhv.ai/mcp
# ✓ MCP サーバー「lhv」を追加しました
$ クロード mcp ログイン lhv
# → ブラウザを開いて LHV ターミナル コピー $ codex mcp add lhv \
--url https://mcp.lhv.ai/mcp
# ✓ MCP サーバー「lhv」を追加しました
# → ブラウザを開いて LHV Cursor Copy // .cursor/mcp.json でサインインします
{
"mcpサーバー": {
"lhv" : {
"url" : "https://mcp.lhv.ai/mcp"
}
}
} Zed Copy // コマンドパレット → "エージェント: コンテキストサーバーを追加"
// または ~/.config/zed/settings.json に貼り付けます
{
"context_servers" : {
"lhv" : {
"url" : "https://mcp.lhv.ai/mcp"
}
}
} Claude Claude Code Codex Cursor Zed 2 LHV 資格情報を使用してサインインします。
ブラウザが開き、見慣れた画面が表示されます。インターネット バンクの場合とまったく同じようにサインインします。アカウント、トランザクション、またはその両方のどの権限を与えるかを選択します。
同意画面で、必要なスコープを正確に選択します

持つべきアシスタント。承認すると、トークンがツールに戻ります。
アカウント 読み取り専用アカウント:読み取り トランザクション 読み取り専用トランザクション:読み取り 3 質問を開始します
それだけです。 AI ツールに戻って、財務状況について尋ねます。 AI アシスタントが残りを処理します。
30 日間接続を維持します。インターネット バンクからいつでもアクセスを取り消してください。
すべてのアカウントの現在の残高はいくらですか?
list_accounts() 2 つの口座に合計 37 396,24 ユーロがあります。つまり、現在の口座に 12 486,24 ユーロ、貯蓄額が 24 910,00 ユーロあります。
3月に食料品にいくら使いましたか?
get_transactions_summary() 2026 年 3 月、18 件のトランザクションで食料品に 412,67 ユーロを費やしました。そのほとんどは、Coop Konsum (184,20 €) と Rimi (121,47 €) で購入できます。
出典: get_transactions · 2026 年 3 月 1 日 – 31 日
今週更新されたサブスクリプションはどれですか?
get_transactions() 今週更新された 3 つのサブスクリプション: Spotify (10,99 ユーロ)、iCloud+ (2,99 ユーロ)、Netflix (12,99 ユーロ) — 合計 26,97 ユーロ。
出典: get_transactions · 今週
4つのツール。
銀行に関するあらゆる質問。
AI アシスタントは、ユーザーの質問に基づいて適切なものを自動的に選択します。
アクセスできるすべてのアカウント。 IBAN、通貨、現在の残高。
特定のアカウントのライブ利用可能残高。迅速かつ直接的な回答。
任意の日付範囲の完全なトランザクション リスト。何が入ってきて、何が出て行ったのか。
総支出と収入、上位取引先。支出に関する広範な質問の概要を迅速に把握できます。
安心と信頼。
インターネット銀行と同じ基盤の上に構築されています。
MCP は、Claude Desktop、Claude Code、Cursor、および増え続けるツールのリストでサポートされるオープン プロトコルです。セットアップには約 2 分かかります。
スマート ID、モバイル ID、ID カード、または生体認証など、お好みのものを使用できます。
アクセス トークンは 1 時間後に期限切れになります。リフレッシュ トークンの有効期限は 30 日間です。
AI アシスタントがあなたのデータを確認できます

た、あなたの代わりに何もしないでください。
すべてのリクエストがログに記録されます。インターネット バンクの設定からいつでもアクセスを取り消してください。
データが AI アシスタントに到達すると、そのツールの利用規約とプライバシー ポリシーに従って処理されます。銀行情報へのアクセスを許可する前に、サービス プロバイダーを信頼していることを確認してください。
REST をお好みですか?
私たちにもそれがあります。
LHV.ai は 2 つのサブドメインで実行されます。 MCP は AI アシスタントにとって主要なものです。 REST API は直接統合に使用できます。
https://mcp.lhv.ai/mcp 4 ツール list_accounts accounts:read すべてのアカウント — IBAN、通貨、利用可能な残高
1 つの IBAN の利用可能残高、決済済み残高、予約済み残高
IBAN の生の CAMT.053 トランザクション、最大 31 日間
合計と上位取引先、最大 31 日間
アカウントとトランザクションのエンドポイント
https://api.lhv.ai/api/v1 API トークンの取得 2 エンドポイント GET /accounts accounts:read 現在の残高を持つすべてのアカウント
ページ分割されたトランザクション ステートメント
これはLHVの正規品ですか？
AI アシスタントはお金を移動できますか?
価格表と利用規約
LHV グループ、Tartu mnt 2、10145、タリン。あなたは、金融サービス プロバイダー AS LHV Pank、LHV Finance、LHV Kindlustus、および AS LHV Varahaldus の Web サイトにアクセスしています。金融サービス契約に署名する前に、サービス規約を確認するか、追加情報を問い合わせてください。引用が遅れています。

## Original Extract

Connect Claude, Cursor and other AI assistants to your LHV accounts via the Model Context Protocol. OAuth2 + Smart-ID, read-only.

Skip to main content EN Sign Up Setup Tools Security API FAQ Your bank, inside your
AI Assistant.
Ask about your balance, pull up last month's transactions, fetch your transaction history. Log in once and just ask.
> Beta Get Started.
Connect in three steps.
MCP is an open protocol with a growing list of supporting tools. Setup takes about two minutes.
In your AI tool's settings, add a new MCP server and paste in the LHV server address.
Coming soon Find LHV directly in your AI tool's marketplace and connect without any manual setup.
1 Open connector settings On claude.ai, go to Customize → Connectors, then click Add → Add custom connector.
2 Add the LHV connector Name it LHV and paste in the server URL. https://mcp.lhv.ai/mcp
3 Connect Click Add, then Connect to sign in with LHV.
Set up on claude.ai, then it's available in the desktop and mobile apps too. Free is limited to one custom connector.
$ claude mcp add lhv \
--transport http \
https://mcp.lhv.ai/mcp
# ✓ Added MCP server 'lhv'
$ claude mcp login lhv
# → opens your browser to sign in with LHV Terminal Copy $ codex mcp add lhv \
--url https://mcp.lhv.ai/mcp
# ✓ Added MCP server 'lhv'
# → opens your browser to sign in with LHV Cursor Copy // .cursor/mcp.json
{
"mcpServers" : {
"lhv" : {
"url" : "https://mcp.lhv.ai/mcp"
}
}
} Zed Copy // Command palette → "agent: add context server"
// or paste into ~/.config/zed/settings.json
{
"context_servers" : {
"lhv" : {
"url" : "https://mcp.lhv.ai/mcp"
}
}
} Claude Claude Code Codex Cursor Zed 2 Sign in with your LHV credentials
Your browser opens to a familiar screen. Sign in exactly as you would in your internet bank. You choose which permissions to give: accounts, transactions, or both.
On the consent screen, pick exactly which scopes you want the assistant to have. Approve, and the token flows back to your tool.
Accounts Read only accounts:read Transactions Read only transactions:read 3 Start asking questions
That's it. Go back to your AI tool and ask about your finances. Your AI assistant handles the rest.
Stays connected for 30 days. Revoke access any time from your internet bank.
What's my current balance across all accounts?
list_accounts() Across your 2 accounts you have 37 396,24 € — 12 486,24 € in your current account and 24 910,00 € in savings.
How much did I spend on groceries in March?
get_transactions_summary() In March 2026, you spent 412,67 € on groceries across 18 transactions. Most of it at Coop Konsum (184,20 €) and Rimi (121,47 €).
Sourced from: get_transactions · 01.03 – 31.03.2026
Which subscriptions renewed this week?
get_transactions() Three subscriptions renewed this week: Spotify (10,99 €), iCloud+ (2,99 €) and Netflix (12,99 €) — 26,97 € in total.
Sourced from: get_transactions · this week
Four tools.
Every banking question.
Your AI assistant picks the right one automatically based on what you ask.
Every account you have access to. IBAN, currency, and current balance.
Live available balance for a specific account. Fast, direct answer.
Full transaction list for any date range. What came in, what went out.
Total spending and income, top counterparties. Faster overview for broad spending questions.
Security and trust.
Built on the same foundations as your internet bank.
MCP is an open protocol supported by Claude Desktop, Claude Code, Cursor and a growing list of tools. Setup takes about two minutes.
Smart-ID, Mobile-ID, ID-card or biometrics, whatever you prefer.
Access tokens expire after 1 hour. Refresh tokens last 30 days.
Your AI assistant can see your data, never do anything on your behalf.
Every request is logged. Revoke access any time from your internet bank settings.
Once your data reaches your AI assistant, it's handled by that tool's terms and privacy policy. Make sure you trust the service provider before giving it access to your banking information.
Prefer REST?
We have that too.
LHV.ai runs on two subdomains. MCP is primary for AI assistants. The REST API is available for direct integrations.
https://mcp.lhv.ai/mcp 4 tools list_accounts accounts:read All accounts — IBAN, currency, available balance
Available, settled, and reserved balance for one IBAN
Raw CAMT.053 transactions for an IBAN, max 31 days
Totals and top counterparties, max 31 days
Accounts & transactions endpoints
https://api.lhv.ai/api/v1 Get API tokens 2 endpoints GET /accounts accounts:read All accounts with current balances
Paginated transaction statement
Is this an official LHV product?
Can my AI assistant move money?
Price list and terms and conditions
LHV Group, Tartu mnt 2, 10145, Tallinn. You are on the website of financial service providers AS LHV Pank, LHV Finance, LHV Kindlustus and AS LHV Varahaldus. Before signing a financial service contract, please review the terms of service or ask for additional information. Quotes are delayed.

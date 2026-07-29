---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49096519"
title: "Fund Momentum MCP: Live VC Fund Data for AI Agents"
article_title: ""
author: "darius88"
captured_at: "2026-07-29T12:56:15Z"
capture_tool: "hn-digest"
hn_id: 49096519
score: 1
comments: 0
posted_at: "2026-07-29T12:21:37Z"
tags:
  - hacker-news
  - translated
---

# Fund Momentum MCP: Live VC Fund Data for AI Agents

- HN: [49096519](https://news.ycombinator.com/item?id=49096519)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T12:21:37Z

## Translation

タイトル: ファンド モメンタム MCP: AI エージェント向けのライブ VC ファンド データ
HN テキスト: Fund Momentum は、ライブ GP シグナルで積極的に展開している 970 以上の VC ファンドを追跡しています。このデータをあらゆる AI エージェントまたは LLM ワークフローで利用できるようにする MCP サーバーを立ち上げたところです。機能 JSON-RPC 2.0 による 5 つのツール: • search_funds — 段階、国、業界によるフィルター • get_fund — フルプロフィール、論文、チェックサイズ、GP • get_fund_signals — ライブ GP シグナル、展開ステータス、強気/逆張りの角度 • match_startup — スタートアップについて説明し、推論とともに一致する上位 10 人の投資家を取得 • get_gp_profile — 個々のパートナー インテリジェンス エージェントネイティブの価格設定 サブスクリプションではなく、クレジット。通話ごとに 0.01 ユーロ。毎月のリセットはありません。クレジットに有効期限はありません。すべての応答には、_meta に Credits_remaining が含まれます。自己登録（人間なし）curl -X POST https://fundmomentum.vc/_api/agent/register \ -H "Content-Type: application/json" \
-d '{"エージェント名": "私のワークフロー", "電子メール": "ops@company.com"}'
# API キーを即座に返す Claude Desktop config
{ "mcpServers": {
"ファンドの勢い": {
"url": "https://fundmomentum.vc/_api/mcp",
"ヘッダー": { "X-API キー": "YOUR_KEY" }
}
}
}

## Original Extract

Fund Momentum tracks 970+ actively deploying VC funds with live GP signals. We just launched an MCP server that makes this data available to any AI agent or LLM workflow. What it does Five tools via JSON-RPC 2.0: • search_funds — filter by stage, country, industry • get_fund — full profile, thesis, check size, GPs • get_fund_signals — live GP signals, deployment status, bullish/contrarian angles • match_startup — describe your startup, get top 10 matched investors with reasoning • get_gp_profile — individual partner intelligence Agent-native pricing Credits, not subscriptions. €0.01 per call. No monthly resets. Credits never expire. Every response includes credits_remaining in _meta. Self-registration (no humans) curl -X POST https://fundmomentum.vc/_api/agent/register \ -H "Content-Type: application/json" \
-d '{"agent_name": "my-workflow", "email": "ops@company.com"}'
# Returns API key instantly Claude Desktop config
{ "mcpServers": {
"fund-momentum": {
"url": "https://fundmomentum.vc/_api/mcp",
"headers": { "X-API-Key": "YOUR_KEY" }
}
}
}


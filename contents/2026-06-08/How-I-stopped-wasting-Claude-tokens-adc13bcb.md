---
source: "https://botverse.cloud"
hn_url: "https://news.ycombinator.com/item?id=48450151"
title: "How I stopped wasting Claude tokens"
article_title: "Botverse — Cloud Services for AI Agents"
author: "MarkTurnerETech"
captured_at: "2026-06-08T21:39:24Z"
capture_tool: "hn-digest"
hn_id: 48450151
score: 2
comments: 0
posted_at: "2026-06-08T19:15:11Z"
tags:
  - hacker-news
  - translated
---

# How I stopped wasting Claude tokens

- HN: [48450151](https://news.ycombinator.com/item?id=48450151)
- Source: [botverse.cloud](https://botverse.cloud)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T19:15:11Z

## Translation

タイトル: クロードトークンの無駄遣いをやめた方法
記事のタイトル: Botverse — AI エージェント向けのクラウド サービス
説明: AI エージェント向けに構築された MCP ネイティブのクラウド サービス。ビデオのエンコード、ファイル処理など。ボット用に構築されています。人間によって設定されます。
HN テキスト: 私は仮想 AI アシスタント SaaS (nearlyme.ai) を構築していて、ドキュメントを読み取って処理させていました - これをスキャンし、これを書き込み、この単語ドキュメントの PDF を作成するなどしました。そして、トークンを焼いていることに気づきました - 2mb PDF には 97,354 個の入力トークンが必要で、クロードと議論するたびに同じカウントを送信していました - これは 10 チャットのトークンで $3.00 です。
そこで、私自身の問題を解決するために botverse.cloud を構築しました。ドキュメント処理 (ビデオや画像の処理も) をオフロードし、トークンではなくセントで従量課金ができるようになりました。ぜひご意見をいただければ幸いです。また、今週末には複雑なワークフロー ツールが追加されたばかりです。これにより、電話会議ビデオを提供し、音声トランスクリプトを生成し、そのドキュメントを生成し、MCP サーバーへの 1 回の呼び出しでビデオを別の形式で保存することができます。 100,000 以上のトークンの代わりに 0.25 ドルを費やしましょう!

記事本文:
Botverse — AI エージェント向けのクラウド サービス Botverse サービスのワークフロー 料金設定 ドキュメント ブログ サインイン サインアップ MCP ネイティブ クラウド サービス ボット用に構築されたクラウド サービス。
人間がアカウントを設定し、残りはボットが行います。ビデオのトランスコーディング、ドキュメント変換、マルチステップ パイプライン用の MCP ツール - あらゆる AI エージェントから呼び出すことができます。
MCP クライアントを Botverse エンドポイントに向けます。 API キーを貼り付けます。それが全体のセットアップです。
MCP をサポートする AI エージェントは、追加のコードを必要とせずに、transcode_video、convert_image、またはその他の Botverse ツールを呼び出すことができるようになりました。
ジョブはサーバーレス インフラストラクチャ上で実行されます。結果は、ダウンロードまたは次のステップに渡す準備ができた署名付き URL として返されます。
ビデオをフォーマット間でトランスコードします。 MP4 (H.264)、VP9 (WebM)、ProRes 422、GIF、および MP3 の抽出 - 任意のパブリック URL またはアップロードされたファイルから。
ドキュメントをフォーマット間で変換します。 Markdown から PDF または DOCX、HTML から Word、プレーン テキストの入力 — サポートされている任意の形式の出力。 URL、アップロードされたファイル、またはインライン コンテンツから。
メディア資産の OMC ベースのナレッジ グラフ。豊富な関係性クエリ — デイリー全体で特定のキャラクターが登場するすべてのシーンを検索します。
パイプラインを送信します。
ボットは先に進みます。
ほとんどのツールはブロックします。 4K トランスコードは 1 時間実行できます。 Botverse ワークフローを使用すると、ボットは 1 回の呼び出しで完全なマルチステップ パイプラインを送信し、ユーザーへの応答、他のタスクの実行、新しいリクエストの処理など、他の必要な作業にすぐに戻ることができます。
このエンジンは、ステップ シーケンス、ファンアウト並列処理、再試行、および失敗ポリシーをサーバー側で処理します。ボットは進行状況を確認するために時々ポーリングするだけです。パイプラインが終了すると、すべての出力 URL が待機状態になります。
// 1. 送信 — すぐに戻ります
const { workflow_id } = await callTool(
"submit_workflow", { 定義 }
);
// 2. ボットは無料です - 他の作業を処理します
await ReplyToUser("あなたのパイプライン

実行中です!");
handleOtherRequests() を待ちます。
// 3. 都合の良いときにいつでもチェックインしてください
結果をみましょう。
する{
スリープを待ちます(10_000);
result = await callTool(
"get_workflow_status", { ワークフロー ID }
);
while (result.status === "処理中");
// 4. すべての出力が準備完了
result.steps.forEach(ステップ => {
console.log(step.step_id, step.output_url);
});
// https://storage.botverse.cloud/... を取り込みます
// mp4 https://storage.botverse.cloud/...
// WebM https://storage.botverse.cloud/...
// オーディオ https://storage.botverse.cloud/... MCP config 1 つの構成ブロック。
すべてのエージェント、すべてのクライアント。
これを MCP クライアント構成 (Claude Desktop、Cursor、Continue、またはその他の MCP 互換エージェント ホスト) に貼り付けます。ボットはすぐにすべての Botverse ツールにアクセスできるようになります。
インストールする SDK はありません。ライブラリの依存関係はありません。 HTTP と JSON-RPC だけです。
{
"mcpサーバー": {
"ボットバース": {
"url": "https://botverse.cloud/mcp",
"ヘッダー": {
「X-API-Key」：「bv_live_••••••••••••••」
}
}
}
} ✓ 現在、Claude Desktop と連携しています テスト済みでライブ中 — 最初のビデオを 2 分以内にトランスコードします 価格 使用した分に応じて支払います。
プリペイドウォレット。購読はありません。最初に 5 ドルをチャージします。ボットがジョブを実行するにつれて、それを使い切ります。
ボットは MCP ツールを介して自身の残高と使用履歴を確認できます。ダッシュボードにログインする必要はありません。
ボットの準備が整いました。
彼らの道具なのか？
5 ドルを追加して設定を貼り付けると、最初のビデオが数分以内にエンコードされます。

## Original Extract

MCP-native cloud services built for AI agents. Video encoding, file processing, and more. Built for bots. Configured by humans.

I was building a virtual AI assistant SaaS (nearlyme.ai) and had it reading and processing documents for me - scan this, write that, make a PDF of this word doc etc. And realized I was burning tokens - a 2mb PDF took 97,354 input tokens and was sending the same count everytime I discussed it with Claude - that's $3.00 in tokens for 10 chats.
So I built botverse.cloud to solve my own issues - now I can offload document processing (and now video and image processing too) and pay as you go in cents not tokens. I'd love your thoughts. Also just added this weekend a complex workflow tool - so you can give it a conference call video and have it generate audio transcript, document of that, save the video in another format all in one call to an MCP server. And spend $0.25 instead of 100,000+ tokens!

Botverse — Cloud Services for AI Agents Botverse Services Workflows Pricing Docs Blog Sign in Sign up MCP-native cloud services Cloud services built for bots.
Humans set up the account, your bots do the rest. MCP tools for video transcoding, document conversion, and multi-step pipelines — callable by any AI agent.
Point your MCP client at the Botverse endpoint. Paste your API key. That's the entire setup.
Any AI agent with MCP support can now call transcode_video, convert_image, or any other Botverse tool — no extra code.
Jobs run on serverless infrastructure. Results come back as signed URLs ready to download or pass to the next step.
Transcode video between formats. MP4 (H.264), VP9 (WebM), ProRes 422, GIF, and MP3 extraction — from any public URL or uploaded file.
Convert documents between formats. Markdown to PDF or DOCX, HTML to Word, plain text in — any supported format out. From a URL, uploaded file, or inline content.
OMC-based knowledge graph for media assets. Rich relationship queries — find every scene with a given character across your dailies.
Submit a pipeline.
Your bot moves on.
Most tools block. A 4K transcode can run for an hour. Botverse Workflows let your bot submit a full multi-step pipeline in a single call and immediately get back to whatever else it needs to do — responding to users, running other tasks, handling new requests.
The engine handles step sequencing, fan-out parallelism, retries, and failure policies server-side. Your bot just polls once in a while to check progress. When the pipeline finishes, all output URLs are waiting.
// 1. Submit — returns immediately
const { workflow_id } = await callTool(
"submit_workflow", { definition }
);
// 2. Bot is free — handle other work
await respondToUser("Your pipeline is running!");
await handleOtherRequests();
// 3. Check in whenever convenient
let result;
do {
await sleep(10_000);
result = await callTool(
"get_workflow_status", { workflow_id }
);
} while (result.status === "PROCESSING");
// 4. All outputs ready
result.steps.forEach(step => {
console.log(step.step_id, step.output_url);
});
// ingest https://storage.botverse.cloud/...
// mp4 https://storage.botverse.cloud/...
// webm https://storage.botverse.cloud/...
// audio https://storage.botverse.cloud/... MCP config One config block.
Every agent, any client.
Paste this into your MCP client config — Claude Desktop, Cursor, Continue, or any other MCP-compatible agent host. Your bot immediately gains access to all Botverse tools.
No SDK to install. No library dependencies. Just HTTP and JSON-RPC.
{
"mcpServers": {
"botverse": {
"url": "https://botverse.cloud/mcp",
"headers": {
"X-API-Key": "bv_live_••••••••••••••"
}
}
}
} ✓ Works with Claude Desktop today Tested and live — transcode your first video in under 2 minutes Pricing Pay for what you use.
Pre-paid wallet. No subscriptions. Top up $5 to start — spend it down as your bots run jobs.
Your bot can check its own balance and usage history via MCP tools — no need to log in to a dashboard.
Your bots are ready.
Are their tools?
Top up $5, paste the config, and your first video is encoding within minutes.

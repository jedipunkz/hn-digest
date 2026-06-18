---
source: "https://www.foglamp.dev/"
hn_url: "https://news.ycombinator.com/item?id=48588220"
title: "Show HN: Foglamp – Open-source o11y for AI SDK"
article_title: "Foglamp - Observability for AI agents"
author: "gustavofior"
captured_at: "2026-06-18T18:56:10Z"
capture_tool: "hn-digest"
hn_id: 48588220
score: 1
comments: 0
posted_at: "2026-06-18T16:55:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Foglamp – Open-source o11y for AI SDK

- HN: [48588220](https://news.ycombinator.com/item?id=48588220)
- Source: [www.foglamp.dev](https://www.foglamp.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T16:55:35Z

## Translation

タイトル: Show HN: Foglamp – AI SDK 用のオープンソース o11y
記事のタイトル: Foglamp - AI エージェントの可観測性
説明: Vercel AI SDK 上に構築された AI エージェントの可観測性レイヤーが欠落しています。すべてのgenerateText / streamText呼び出しのコスト、レイテンシ、トークン、分散トレース、評価、およびアラートが2行で表示されます。

記事本文:
Foglamp - AI エージェントの可観測性 Foglamp の料金
実際に見ることができる Ship AI エージェントの監視を開始する
すべての LLM 呼び出しのコスト、遅延、品質を確認します。ユーザーが検出する前に、不正な出力を検出します。
コピーエージェントプロンプト | 「出荷後 3 日で 10 倍のコスト後退を発見しました。Foglamp は 1 週目で元が取れました。」 に合わせてカスタマイズされました。
3 週目 コストは 2 倍になり、回答は悪化しました。
すると顧客は苦情を言い始めます。
Dana #support 2m アシスタントが 3 月に廃止した返金窓口について引用しました 😤
+1 (415) 555-0148 お客様 6 分 これはボットですか?存在しない注文番号を言われました
@jordanbuilds 12.4k フォロワー 14m @acme あなたの「AI サポート」が自信を持って追跡リンクを発明しました。やったー。
mara コミュニティ 2100 万人 今日ボットから完全に間違った答えを受け取った人は他にいますか??
楽器は一度。すべての通話のコスト、トレース、評価、アラート、およびエージェントごとの支出を取得します。
2 行は、すべてのgenerateText / streamText呼び出しを実行します。
エージェントごとのスパン、待ち時間、支出 - 完全なコール フロー。
オーケストレーター 2 。 4 2.4 秒 研究者 search() ライターがキューに入れられた 批評家がキューに入れられた Evals
コードチェックとLLMジャッジを使用して実稼働トラフィックをスコアリングします。
PII なし 合格率: 9 4 % 94% 分散トレース
実行ごとにウォーターフォールを実行し、スパンごとに正確なプロンプトと応答を提供します。
コスト、レイテンシー、エラー率に関するしきい値ルール。
モデル、エージェント、顧客ごとに、すべての通話にかかる費用を正確に把握します。
クロード オーパス 4.8 $323.12 ジェミニ ジェミニ 3.5 プロ $43.06 GPT-5.5 ミニ $85.05 合計 $1、2 4 8。 5 0 $1,248.50 エージェントは霧の中で動いています。
コスト、レイテンシー、エラー、評価スコア、すべてがそこにあり、すべてが目に見えません。モデルを包み、ライトをオンにします。
エージェント プロンプトのコピー Foglamp AI エージェントに欠落している可観測性レイヤー。

## Original Extract

The missing observability layer for AI agents built on the Vercel AI SDK. Costs, latency, tokens, distributed traces, evals, and alerts for every generateText / streamText call — in two lines.

Foglamp - Observability for AI agents Foglamp Pricing
Start monitoring Ship AI agents you can actually see
See the cost, latency, and quality of every LLM call. Catch bad output before your users do.
Copy agent prompt | Tailor made for “ Caught a 10× cost regression 3 days after shipping. Foglamp paid for itself in week one. ”
Week 3 Costs doubled, answers worse.
Then Customers start complaining.
Dana #support 2m the assistant just quoted a refund window we killed in March 😤
+1 (415) 555-0148 Customer 6m is this a bot? it gave me an order number that doesn't exist
@jordanbuilds 12.4k followers 14m @acme your “AI support” confidently invented a tracking link. yikes.
mara community 21m anyone else getting totally wrong answers from the bot today??
Instrument once. Get cost, traces, evals, alerts, and per-agent spend on every call.
Two lines instruments every generateText / streamText call.
Per-agent spans, latency, and spend - with the full call flow.
orchestrator 2 . 4 2.4 s researcher search() writer queued critic queued Evals
Score production traffic with code checks and LLM judges.
No PII Pass rate: 9 4 % 94% Distributed traces
Waterfall every run, with the exact prompt and response per span.
Threshold rules on cost, latency, and error rate.
Know exactly what every call costs — by model, agent, customer.
Claude Opus 4.8 $323.12 Gemini Gemini 3.5 Pro $43.06 GPT-5.5 mini $85.05 Total $ 1 , 2 4 8 . 5 0 $1,248.50 Your agents are running in the fog.
Cost, latency, errors, eval scores — all there, all invisible. Wrap your model and turn the light on.
Copy agent prompt Foglamp The missing observability layer for AI agents.

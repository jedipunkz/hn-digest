---
source: "https://costlens.dev/blog/claude-fable-5-pricing-production-costs"
hn_url: "https://news.ycombinator.com/item?id=48492364"
title: "Claude Fable 5 costs $10/$50M tokens – what that means in production"
article_title: "Claude Fable 5 Pricing: What $10/$50 Per Million Tokens Actually Costs in Production | CostLens"
author: "j_filipe"
captured_at: "2026-06-11T16:26:19Z"
capture_tool: "hn-digest"
hn_id: 48492364
score: 1
comments: 0
posted_at: "2026-06-11T16:14:53Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable 5 costs $10/$50M tokens – what that means in production

- HN: [48492364](https://news.ycombinator.com/item?id=48492364)
- Source: [costlens.dev](https://costlens.dev/blog/claude-fable-5-pricing-production-costs)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T16:14:53Z

## Translation

タイトル: Claude Fable 5 のコストは 1,000 ドル/5,000 万ドルのトークン – それが本番環境で何を意味するか
記事のタイトル: Claude Fable 5 価格: 100 万トークンあたり 10 ドル/50 ドルの実際の生産コスト |コストレンズ
説明: Fable 5 のコストは 100 万トークンあたり $10/$50、つまり 2x オーパスです。実際の開発者は、エージェント ループに 1 時間あたり 200 ～ 400 ドルを費やしています。価格ページでは分からないことがここにあります。

記事本文:
AI が出荷するものを測定します。 ROI をエンジニアリング リーダーに自動的に証明します。
hello@costlens.dev コミュニティとサポート © 2026 CostLens.無断転載を禁じます。
← ブログに戻る Claude Fable 5 の価格: 100 万トークンあたり 10 ドル/50 ドルの実際の生産コスト
Fable 5 のコストは 100 万トークンあたり $10/$50、つまり 2x オーパスです。実際の開発者は、エージェント ループに 1 時間あたり 200 ～ 400 ドルを費やしています。価格ページでは分からないことがここにあります。
Claude Fable 5 API: 100 万トークンあたり入力 $10 / 出力 $50 (入力で Opus 4.8 x 2、GPT-5.5 x 2)
実際のエージェント ループのコスト: 早期導入者のレポートに基づくと、1 時間あたり 200 ～ 400 ドル
クエリの 5% はサイレントに Opus 4.8 にフォールバックします - 引き続き Fable 価格を支払います
無料アクセスは 6 月 22 日に終了します - その後は使用クレジットが必要になります
AI サークルの人々は皆、ベンチマークや素晴らしいデモを投稿するのが大好きです。誰も話さないのは請求書のことです。私たちは実稼働環境での LLM の使用状況を追跡してきましたが、常に明らかになる 1 つの真実は、ステッカー価格が実際のコストを反映していることはほとんどないということです。 Anthropic の最新フラッグシップである Claude Fable 5 が登場した今、チームが実際にいくら払っているのかについて話し合う時期が来ています。
Fable 5 の価格は、入力トークン 100 万あたり 10 ドル、出力トークン 100 万あたり 50 ドルです。これは、まさに Opus 4.8 の 2 倍 (5 ドル/25 ドル)、入力で GPT-5.5 の 2 倍になります。プロンプト キャッシュに関しては、入力トークンの 90% 割引が引き続き提供されます。しかし、トークンあたりのレートは話の半分にすぎません。
クロード・ファブル 5 の価格はいくらですか?フルモデルの比較
すべての主要モデルを並べて示します (100 万トークンごと)。
出典: 人間の価格設定、OpenAI の価格設定
クロード・ファブル 5 の実際の制作費はいくらですか?
価格ページには $10/$50 と書かれています。実際の本番環境での使用は次のとおりです。
ウーバーは4月までに2026年のAI予算をすべて使い果たした。クロード コードは財務部門の予想を上回る速さで 5,000 人のエンジニアに広がりました。同社の CTO は、超過を認めました。「

「必要だと思っていた予算はすでに吹き飛んでいます。」 エンジニア 1 人当たりの月々のコストは平均 150 ～ 250 ドルで、パワー ユーザーの場合は 500 ～ 2,000 ドルに達します。Uber では現在、エンジニアの上限を月 1,500 ドルに設定しています。それは Opus 4.8 の場合であり、Fable 5 はその 2 倍の価格です。
50 分間のコード監査の費用は 2,200 レアル (約 400 米ドル) です。ある開発者は、コードベース監査で 96 個の並列エージェントを実行し、440 万個のトークンを消費しました。 1時間以内。シングルセッション。 Fable 5 の価格では、エージェントは高速にループします。
ある HN ユーザーは、Sonnet for Haiku を削除し、呼び出しをバッチ処理し、無関係な入力をフィルタリングし、出力を短縮することで、Claude のコストを「月額 70 ドルから 1 ペニーに」削減しました。迅速なエンジニアリングの魔法ではなく、インテリジェントなリソース割り当てです。
クロード・フェイブル 5 Opus 4.8 へのサイレント・フォールバック
Fable 5 が常に Fable 5 であるとは限りません。サイバーセキュリティ、生物学、または化学の保護手段をトリガーするクエリの場合、Anthropic はリクエストを静かに Opus 4.8 にルーティングします。これが影響するセッションは 5% 未満です。
問題は、これらのクエリに対する Opus 4.8 の応答に対して Fable 5 の価格を支払っていることです。開発者の反発を受けて、Anthropic は API 経由でフォールバックを表示できるようにし始めましたが、最初の数日間は表示されませんでした。
セキュリティ、バイオテクノロジー、またはこれらの分類子をトリガーするドメインに携わっている場合、実質的な品質あたりのコストは定価よりも悪くなります。
クロード・ファブル 5 の無料アクセスは 6 月 22 日に終了します
Anthropic は、6 月 9 日から 22 日まで、すべての Pro、Max、Team、Enterprise 加入者に対して Fable 5 を無料で提供しました。最も機能的なパブリック モデルを 13 日間追加料金なしで利用できます。その後、6 月 23 日に使用量クレジットに切り替わります。
これは古典的な依存関係の構築です。チームがそれを統合し、それを中心にワークフローが構築され、その後、請求が開始されます。月額 200 ドルの Max サブスクリプションは、すでに Fable を大量に使用しても 90 分しか持続しませんでした。あるユーザーは、一日の勤務を維持するには月額 600 ドルが必要になると計算しました。
クロード・ファブルは5の価値がありますか

エンジニアリングチームの代償は?
100,000 人以上の開発者を対象とした MIT の新しい調査 (NBER で発表) では、AI コーディング エージェントによってコード量が 180% 増加することが判明しましたが、実際のソフトウェア リリースは 30% しか増加しませんでした。
より高性能なモデル = より多くの生産量 = より大きな請求書。出荷されるソフトウェアが 30% 多い場合、30% 多くの結果を得るためにトークンごとに 2 倍の料金を支払うことになります。計算上、ルーチンタスクにはフェイブルが向いていません。
6 月 22 日以降のクロード・ファブル 5 のコストを管理する方法
インテリジェントにルーティング — すべてのタスクに Fable 5 が必要なわけではありません。分類、要約、簡単な編集 → Sonnet 4.6 または Haiku を 3 ～ 10 倍のコストで実行できます。 OpenAI と Anthropic のコスト比較で詳細をご覧ください。
積極的にキャッシュする — Fable 5 では、キャッシュされた入力トークンを 90% 割引で提供します。使ってください。同じプロンプト = 同じ答え = 二重支払いは不要です。
キル スイッチを設定する — エージェント ループが暴走した場合、1 回の実行で 400 ドルを超える前に自動カットオフが必要です。これはまさに、セッションごとのコスト追跡を目的として構築したものであり、ハードリミットによるリアルタイムの支出の可視化です。
今すぐ使用状況を監査します。無料ですが、ワークフローが実際に消費するトークンの数を測定します。ベースライン データを取得するには、6 月 22 日までの期限があります。
フォールバック税の予算 — クエリの 5% が静かに Opus にヒットする可能性があります。それを品質あたりのコストの計算に織り込みます。 AI API コストの削減に関する完全なガイドをお読みください。
Claude Fable 5 の API 呼び出しあたりの料金はいくらですか?
100 万トークンあたり 10 ドルまたは 50 ドルとして、典型的な 2,000 トークンのリクエストと 1,000 トークンの応答のコストは約 0.07 ドルです。数百万のトークンを処理するエージェント ループの場合、1 時間あたり 200 ～ 400 ドルが予想されます。
Claude Fable 5 は GPT-5.5 より安いですか?
いいえ。例 5 は、入力で 2 倍の GPT-5.5 (10 ドル対 5 ドル)、出力で 1.7 倍 (50 ドル対 30 ドル) です。 Anthropicの最も高価なモデルです。
クロード・フェイブル 5 への無料アクセスはいつ終了しますか?
2026 年 6 月 22 日。それ以降は、Pro/Max サブスクリプションに加えて使用量クレジットが必要になります。

Elisabete Romão 氏は CostLens でパートナーシップを率いており、次の四半期の予算危機に陥る前にチームが AI 支出を追跡および制御できるよう支援しています。
アントロピックの公式発表
フォーブス - MIT AI 生産性調査
NBER - 書き込みコードと配送コード
CostLens は、単純なプロンプトをより安価なモデルに自動的にルーティングします。
品質あたりのコストでランク付けされた GPT-4 代替製品 (2026 年 6 月)
OpenAI の請求額を 40 ～ 70% 削減: 2026 年のハンドブック
CostLens と AI を超える: 実際に AI の ROI を証明できるのはどちらですか?

## Original Extract

Fable 5 costs $10/$50 per million tokens — 2x Opus. Real developers are spending $200-400/hour on agent loops. Here's what the pricing page won't tell you.

Measure what AI ships. Prove ROI to engineering leaders — automatically.
hello@costlens.dev Community & Support © 2026 CostLens. All rights reserved.
← Back to blog Claude Fable 5 Pricing: What $10/$50 Per Million Tokens Actually Costs in Production
Fable 5 costs $10/$50 per million tokens — 2x Opus. Real developers are spending $200-400/hour on agent loops. Here's what the pricing page won't tell you.
Claude Fable 5 API: $10 input / $50 output per million tokens (2x Opus 4.8, 2x GPT-5.5 on input)
Real agent loop costs: $200-400/hour based on early adopter reports
5% of queries silently fall back to Opus 4.8 — you still pay Fable pricing
Free access ends June 22 — then usage credits required
Everyone in AI circles loves to post benchmarks and dazzling demos. What nobody talks about is the invoice. We've been tracking LLM usage in production, and one truth consistently emerges: the sticker price rarely reflects the true cost. With Anthropic's latest flagship , Claude Fable 5, hitting the scene, it's time to talk about what teams are actually paying.
Fable 5 arrives at $10 per million input tokens and $50 per million output tokens — precisely double Opus 4.8 ($5/$25), and 2x GPT-5.5 on input. It still offers the 90% input token discount for prompt caching. But the per-token rate is only half the story.
How Much Does Claude Fable 5 Cost? Full Model Comparison
Here's every major model side by side (per million tokens):
Source: Anthropic pricing , OpenAI pricing
What Does Claude Fable 5 Actually Cost in Production?
The pricing page says $10/$50. Here's what production use actually looks like.
Uber blew their entire 2026 AI budget by April. Claude Code spread across 5,000 engineers faster than finance predicted. Their CTO confirmed the overrun: "the budget I thought I would need is blown away already." Monthly costs per engineer averaged $150-$250, with power users hitting $500-$2,000. Uber now caps engineers at $1,500/month . And that was on Opus 4.8 — Fable 5 is 2x the price.
A 50-minute code audit cost R$2,200 (~$400 USD). One developer ran 96 parallel agents on a codebase audit, consuming 4.4 million tokens. Under an hour. Single session. At Fable 5 pricing, agent loops compound fast.
One HN user cut Claude costs "from $70/month to pennies" by dropping Sonnet for Haiku, batching calls, filtering irrelevant inputs, and shortening outputs. Not prompt engineering magic — intelligent resource allocation.
Claude Fable 5 Silent Fallback to Opus 4.8
Fable 5 isn't always Fable 5. For queries triggering cybersecurity, biology, or chemistry safeguards, Anthropic silently routes requests to Opus 4.8 . This affects fewer than 5% of sessions.
The problem: you're paying Fable 5 pricing for Opus 4.8 responses on those queries. After developer backlash , Anthropic started making fallbacks visible via API — but for the first few days, it was invisible.
If you're in security, biotech, or any domain that triggers those classifiers, your effective cost-per-quality is worse than the sticker price.
Claude Fable 5 Free Access Ends June 22
Anthropic made Fable 5 free for all Pro, Max, Team, and Enterprise subscribers from June 9-22. Thirteen days of the most capable public model, no extra charge. Then on June 23, it switches to usage credits.
This is classic dependency building. Teams integrate it, workflows get built around it, then billing kicks in. The $200/month Max subscription was already lasting only 90 minutes of heavy Fable use. One user calculated you'd need $600/month to sustain a working day.
Is Claude Fable 5 Worth the Price for Engineering Teams?
A new MIT study (published in NBER ) across 100,000+ developers found AI coding agents boost code volume by 180% — but actual software releases increased by only 30%.
More capable model = more output = bigger invoice. If only 30% more software ships, you're paying 2x per token for 30% more results. The math doesn't favour Fable for routine tasks.
How to Manage Claude Fable 5 Costs After June 22
Route intelligently — Not every task needs Fable 5. Classification, summarization, simple edits → Sonnet 4.6 or Haiku at 3-10x less cost. Learn more in our OpenAI vs Anthropic cost comparison .
Cache aggressively — Fable 5 offers 90% discount on cached input tokens. Use it. Same prompt = same answer = don't pay twice.
Set kill switches — If an agent loop runs away, you need automatic cutoffs before a single run racks up $400. This is exactly what we built per-session cost tracking for — real-time spend visibility with hard limits.
Audit your usage now — While it's free, measure how many tokens your workflows actually consume. You have until June 22 to get baseline data.
Budget for the fallback tax — 5% of queries may silently hit Opus. Factor that into cost-per-quality calculations. Read our full guide on reducing AI API costs .
How much does Claude Fable 5 cost per API call?
At $10/$50 per million tokens, a typical 2,000-token request with 1,000-token response costs about $0.07. For agent loops processing millions of tokens, expect $200-400/hour.
Is Claude Fable 5 cheaper than GPT-5.5?
No. Fable 5 is 2x GPT-5.5 on input ($10 vs $5) and 1.7x on output ($50 vs $30). It's Anthropic's most expensive model.
When does free Claude Fable 5 access end?
June 22, 2026. After that, usage credits are required on top of Pro/Max subscriptions.
Elisabete Romão leads partnerships at CostLens, where we help teams track and control AI spend before it becomes next quarter's budget crisis.
Anthropic official announcement
Forbes - MIT AI productivity study
NBER - Writing Code vs. Shipping Code
CostLens routes simple prompts to cheaper models automatically.
GPT-4 Alternatives Ranked by Cost Per Quality (June 2026)
Cut Your OpenAI Bill 40-70%: The 2026 Playbook
CostLens vs Exceeds AI: Which Actually Proves AI ROI?

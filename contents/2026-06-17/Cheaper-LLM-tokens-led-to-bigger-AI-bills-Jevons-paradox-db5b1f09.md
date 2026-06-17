---
source: "https://northwoodsystems.ai/blog/ai-token-economics"
hn_url: "https://news.ycombinator.com/item?id=48569539"
title: "Cheaper LLM tokens led to bigger AI bills (Jevons paradox)"
article_title: "AI Token Economics: Why Cheaper Tokens Made Your Bill Explode — Northwood Systems"
author: "AndrewLiu96"
captured_at: "2026-06-17T13:23:15Z"
capture_tool: "hn-digest"
hn_id: 48569539
score: 2
comments: 0
posted_at: "2026-06-17T12:35:19Z"
tags:
  - hacker-news
  - translated
---

# Cheaper LLM tokens led to bigger AI bills (Jevons paradox)

- HN: [48569539](https://news.ycombinator.com/item?id=48569539)
- Source: [northwoodsystems.ai](https://northwoodsystems.ai/blog/ai-token-economics)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T12:35:19Z

## Translation

タイトル: LLM トークンの安さが AI 請求額の増加につながった (ジェボンズのパラドックス)
記事のタイトル: AI トークンの経済学: 安価なトークンのせいで請求額が爆発的に増加した理由 — Northwood Systems
説明: LLM の価格は 1 年で最大 80% 下落し、とにかく AI への支出は 2 倍になりました。メカニズム、計算、そしてそれについて何をすべきか。

記事本文:
AI トークンの経済学: なぜ安価なトークンが請求書を爆発させたのか — Northwood Systems Northwood Systems のセキュリティ ユース ケース 運用化 概要 お問い合わせ デモのリクエスト ブログ / 分析 このページについて ジェボンズのパラドックスが AI 予算を実行している エージェント コーディング 1 ターンの実際のコスト 開発者の支出はべき乗則に従う 変動費を固定費に変換する 3 つ目はまだ解決されていない 出典 2026 年 6 月 16 日 AI トークンの経済学: なぜ安価なトークンが請求書を爆発させたのか
トークンの価格は、歴史上のほぼどのテクノロジーコストよりも速く崩壊しました。しかし、エンジニアリングチームは緊急支出の上限に達し、ライセンスを取り消している。なぜそのようなことが起こったのかを理解することが、問題を解決するための第一歩です。
Uber は年間 AI 予算を 4 か月で使い果たしました。無駄遣いをするのではなく、そのリーダーシップが奨励したことを正確に行うことによってです。同社には AI の多用を称える社内リーダーボードがあり、幹部らは生産性の向上を公に賞賛していましたが、その後法案が届きました。その結果、各エージェント コーディング ツールには従業員 1 人あたり月額 1,500 ドルのハードキャップが適用され、2026 年 6 月より発効します。1
この話は、ある企業の計画の甘さを警告するものではありません。これは、従量制のトークンごとの価格設定がエージェントのワークロードに大規模に対応した場合に何が起こるかを示すプレビューであり、現時点で予算内に収まります。
ジェボンズのパラドックスが AI 予算を使い果たしている
1865年、経済学者ウィリアム・スタンレー・ジェヴォンズは直観に反することに気づきました。蒸気エンジンの効率が向上し、作業単位あたりの運転コストが安くなったため、石炭の総消費量は減るどころか増加しました。効率化により、これまで存在しなかった需要が解放されました。
ジェボンズのパラドックスは、AI への支出に起こっていることです。トークンの価格は 2025 年から 2026 年の間に約 80% 下落しました。2 エンジニアはその節約をポケットに入れませんでした。彼らはそれらを許可として使用しました

o より多く、より長く、より野心的に走ります。 10 ドルのタスクが 2 ドルになるため、チームはタスクを 1 回ではなく 5 回実行し、エージェントに渡すと、自動的に 50 回実行されます。
最も強力な反論は、「単価が 80% 下がったとしても、使用量が 3 倍になっても請求額は横ばいになる」です。これは、チャット スタイルのシングル ターン インタラクションにも当てはまります。エージェントはトークン消費量を 3 倍にしないため、エージェント ループを導入すると完全に機能しなくなります。それを50倍にします。 3 単一のエージェント コーディング セッションは、タスクごとに 100 ～ 350 万のトークンをプッシュするようになりました。 4 1 つのエージェント コーディング ツールを頻繁に使用すると、Uber の月額上限である 1,500 ドルを単独でクリアできます。
エージェント コーディングの 1 ターンに実際にかかる費用
Claude Opus 4.8 を例に挙げます。これは、上級エンジニアが複雑なリファクタリング タスクを行う際に無理なく使用できるモデルです。入力トークンは 100 万あたり 5 ドルで実行されます。出力トークンは 100 万あたり 25 ドルで実行されます。
適切なコンテキストを伴う 1 つのエージェント ターン: 200,000 入力トークン × $5/M = $1.00 。モデルは 50,000 出力トークン × $25/M = $1.25 で応答します。合計: 1 ターンあたり 2.25 ドル。
これを実際の稼働日で乗算すると、1 日あたり 40 ターン、20 営業日となります。 1 人のエンジニアが 1 つのモデルで 1 つのツールを使用すると、月額 1,800 ドルになります。 Uber の 1,500 ドルの上限はカバーしません。
以下の価格表は、出力トークンの数が重要である理由を示しています。入力はステッカー価格です。出力は請求書です。
100万トークンあたりのコスト、USD、モデル別のインプットとアウトプット
$0 $10 $20 $30 $40 Gemini Flash-Lite Gemini 3.1 Pro Claude Sonnet 4.6 GPT-5.4 Claude Opus 4.8 GPT-5.5 プロバイダーの価格、2026 年 6 月に集計 · cloudzero.com 出力トークンのコストは、すべての主要モデルで入力トークンの 4 ～ 10 倍です。エージェント ワークロードでは、出力ボリュームがエスケープ変数となります。
開発者の支出はべき乗則に従う
すべてのエンジニアが月給 1,800 ドルに達するわけではありません。単一のサブスクリプションを使用する個人開発者

ツールの支払いは約 100 ドルです。マルチツールのヘビーユーザーは約 400 ドルを獲得します。実際に生産性の向上を得るパワー エージェント ユーザーは、1,500 ドルを実行します。また、一部のエンジニアが月額 2,000 ドルを稼いでいるのを発見した後、Microsoft は従業員の AI ライセンスをキャンセルしたと伝えられています。 7
報告範囲の上限、USD、2026
$0 $500 $1k $1.5k $2k Solo (サブスクリプション) マルチツールのヘビーユーザー パワーエージェントユーザー MS エンジニア (ライセンスキャンセル) Morph LLM (範囲)。 Microsoft 経由のレポート · morphllm.com 開発者あたりの毎月の AI コーディング費用は、ツールの使用パターンに応じて 20 倍以上異なります。生産性の向上は高価な部分に集中します。
その配分は、ガバナンスについてどのように考えるかにとって重要です。 AI から最大のビジネス価値を生み出しているエンジニアは、構造的には、最大の請求額を生み出しているエンジニアと同じです。ツールごとの鈍いキャップが両方をキャッチします。
FinOps Foundation によると、組織の 63% が現在、FinOps の懸念事項として AI を挙げており、2024 年の 31% から増加しています。 5 その倍増はパニックではありません。トークンごとの請求には自然な上限がなく、財務チームはそれを予測するように作られていないという認識です。
変動費を固定費に変換する
外部 LLM API に費やすすべての金額は、使用量に応じて増減する変動コストです。構造に焼き付けられたキャップはありません。予算がすでに移動した後、事後的に手動で上限を課します。
構造的な代替案は、変動コストを固定の計画可能なコスト、つまり所有するインフラストラクチャ、実行するモデル、タクシーのメーターというよりもデータセンターの項目に近い請求書に変換することです。これはアーキテクチャの変更であり、構成の調整ではありません。
スタックを所有すると、2 番目の問題も同じ決定にまとめられます。機密コードや独自データを外部 API に送信できないチーム

そもそも、厳格なデータ常駐要件を持つ規制された業界と同様に、アーキテクチャの 1 つの選択からコスト管理とデータ管理を実現できます。モデルが独自の境界内で実行される場合、支出はプロビジョニングされた容量であり、データがそこから離れることはありません。
正直な反対意見は、所有インフラストラクチャの初期費用が高くなるということです。それは真実であり、慎重にモデル化する必要があります。損益分岐点は、チームの規模、モデルの組み合わせ、エンジニアが実際にべき乗則曲線のどこまで上に位置するかによって異なります。しかし、年間予算を 4 か月で使い果たし、その後、一律の上限に達するという Uber のシナリオには、アーキテクチャ上の上限のない従量制の外部 API という、特定のインフラストラクチャの形状が背後にあります。
3つ目はまだ解決していない
FinOps Foundation の数字をもう一度見てください。 2 年前、AI の支出が FinOps の懸念事項であると考えていた組織は 3 社に 1 社未満でした。今日ではほぼ 3 分の 2 です。残りの 3 分の 1 はまだ追いついていないか、生産性の向上によりオープン メーターが正当化されると判断したためです。
その2番目のポジションは、適切なスケールでしばらくは守ることができます。ある企業は、従業員の使用量制限を制定できなかったため、AI に約 5 億ドルを費やしたと報告されています。 7 MIT の調査によると、エンタープライズ GenAI プロジェクトの約 95% が 6 か月以内に測定可能な経済的利益を達成できていないことが示されています。 6 曖昧なリターンに対する無制限の支出は、取締役会が要求した場合に維持するのが難しい立場です。
この曲線を先取りしているチームにとって有効な動きは、特定のエージェント ワークロードのコストをモデル化し (上記の計算を開始点として使用)、それを実際に測定できる生産性収益と照らし合わせてマッピングし、従量制の外部支出と固定の所有インフラストラクチャのどちらがその比率をより適切に制御できるかを決定します。入力トークンのステッカー価格をあなたの fi の数字にしないでください。

ナンスチームが見ています。
01 トークンの価格は 1 年で最大 80% 下落しましたが、請求額は上昇しました。これは、より安価なトークンにより、チャット プロンプトの 50 倍のトークンを消費するエージェント ワークロードのロックが解除されたためです。これがジェボンズのパラドックスであり、自動操縦で実行されます。
02 出力トークンはエスケープする変数です。 Opus 4.8 の料金では、1 日あたり 40 回のエージェント ターンを実行する 1 人のパワー ユーザーの費用は月額 1,800 ドルで、単一ツールに対する Uber のハードキャップを超えています。
03 開発者の AI 支出はべき乗則分布に従います。最大の価値を生み出しているエンジニアは、構造的に最大の請求額も生み出しています。鈍いキャップは両方をカットします。
04 トークンごとの請求にはアーキテクチャ上の上限がありません。損傷後に手動で制限を課します。構造的な修正は、変動するトークンの支出を固定のインフラストラクチャ コストに変換することです。
05 現在、組織の 63% が FinOps の積極的な懸念事項として AI を挙げており、2 年前の 31% から増加しています。これに先立ってチームはワークロードをモデル化し、構築か購入かを明確に決定しました。
チームが AI 支出の上限に達している場合、外部 API を除外するデータ常駐の制約に対処している場合、またはその両方の場合は、実際のワークロードに対する数値を喜んで詳しく調査します。
TechCrunch、「4か月で予算を使い果たした後、Uberは従業員のAI支出を制限」（2026年6月2日）。
CloudZero、「LLM API の価格比較」。 100 万トークンあたりの入出力価格と前年比最大 80% の下落 (2026 年)。
LeanOps、「エージェント AI コストの暴走: トークン予算の問題」 。エージェントはチャット プロンプトの約 50 倍のトークンを消費します (2026 年)。
Morph LLM、「AI コーディングのコスト」。月間トークン使用量の中央値 (開発者あたり約 5,100 万)、エージェント タスクあたりのトークン数、および開発者あたりの月間支出範囲 (2026 年)。
FinOps 財団、finops.org 。 FinOps の積極的な懸念事項として AI を挙げている組織の割合、31% (2024 年) → 63% (2025 年)。
MIT プロジェクト NANDA、GenAI の分断: AI の現状

ビジネス (2025) で。企業の生成 AI プロジェクトの約 95% は、6 か月以内に測定可能な経済的利益が得られません。
第二次産業レポート (2026 年)。 Microsoft のエンジニアは、ある企業でエージェント トークンの請求額が月額約 2,000 ドル、アンマネージド AI の支出が約 5 億ドルであると報告しています。一次情報源は出版前にまだ確認されていません。

## Original Extract

LLM prices fell ~80% in a year, and AI spend doubled anyway. The mechanism, the math, and what to do about it.

AI Token Economics: Why Cheaper Tokens Made Your Bill Explode — Northwood Systems Northwood Systems Security Use cases Operationalize About Contact Request a demo Blog / Analysis On this page The Jevons paradox is running your AI budget What one agentic coding turn actually costs Developer spend follows a power law Converting variable cost into fixed cost The third that hasn't solved this yet Sources June 16, 2026 AI Token Economics: Why Cheaper Tokens Made Your Bill Explode
Token prices collapsed faster than almost any technology cost in history. Yet engineering teams are hitting emergency spending caps and cancelling licences. Understanding why that happened is the first step to fixing it.
Uber burned through its entire annual AI budget in four months. Not by being wasteful, but by doing exactly what its leadership encouraged. The company had internal leaderboards celebrating heavy AI usage, executives publicly praised the productivity gains, and then the bill arrived. The result: a $1,500-per-month hard cap on each agentic coding tool, per employee, effective June 2026. 1
That story isn't a cautionary tale about one company's poor planning. It's a preview of what happens when metered, per-token pricing meets agentic workloads at scale, and it's landing in your budget right now.
The Jevons paradox is running your AI budget
In 1865, economist William Stanley Jevons noticed something counterintuitive. As steam engines became more efficient, cheaper to run per unit of work, total coal consumption went up, not down. Efficiency unlocked demand that hadn't existed before.
The Jevons paradox is what's happening to your AI spend. Token prices dropped roughly 80% between 2025 and 2026. 2 Your engineers didn't pocket those savings; they used them as permission to run more, longer, and more ambitiously. A task that cost $10 now costs $2, so your team runs it five times instead of once, then hands it to an agent that runs it fifty times automatically.
The strongest counter-argument: "If unit costs fell 80%, even tripling usage keeps the bill flat." That's true for chat-style, single-turn interactions. It breaks completely once you introduce agentic loops, because an agent doesn't triple token consumption. It multiplies it by 50x. 3 A single agentic coding session now pushes 1–3.5 million tokens per task; 4 one agentic coding tool, used heavily, clears Uber's $1,500 monthly cap on its own.
What one agentic coding turn actually costs
Take Claude Opus 4.8, a model your senior engineers might reasonably reach for on a complex refactoring task. Input tokens run $5 per million; output tokens run $25 per million.
A single agentic turn with a reasonable context: 200,000 input tokens × $5/M = $1.00 . The model responds with 50,000 output tokens × $25/M = $1.25 . Total: $2.25 per turn.
Now multiply that across a real workday: 40 turns per day, 20 working days. That's $1,800 a month, from one engineer, using one tool, on one model. Uber's $1,500 cap doesn't cover it.
The pricing chart below shows why output tokens are the number that matters. Input is the sticker price. Output is the bill.
Cost per 1M tokens, USD, input vs output by model
$0 $10 $20 $30 $40 Gemini Flash-Lite Gemini 3.1 Pro Claude Sonnet 4.6 GPT-5.4 Claude Opus 4.8 GPT-5.5 Provider pricing, compiled June 2026 · cloudzero.com Output tokens cost 4–10× input tokens across every major model. On agentic workloads, output volume is the variable that escapes.
Developer spend follows a power law
Not every engineer hits $1,800 a month. A solo developer on a single subscription tool pays roughly $100. A heavy multi-tool user lands around $400. The power agentic user, the one actually getting the productivity gains, runs $1,500. And Microsoft reportedly cancelled employee AI licences after discovering some engineers were running $2,000 per month each. 7
Upper bound of reported range, USD, 2026
$0 $500 $1k $1.5k $2k Solo (subscription) Heavy multi-tool user Power agentic user MS engineers (licences cancelled) Morph LLM (ranges); Microsoft via reporting · morphllm.com Monthly AI coding spend per developer varies by more than 20× depending on tool usage pattern. The productivity gains concentrate in the expensive tail.
That distribution matters for how you think about governance. The engineers generating the most business value from AI are, structurally, the same engineers generating the largest bills. Blunt per-tool caps catch both.
Sixty-three percent of organisations now name AI an active FinOps concern, up from 31% in 2024, according to the FinOps Foundation. 5 That doubling isn't panic; it's recognition that per-token billing has no natural ceiling, and finance teams weren't built to forecast it.
Converting variable cost into fixed cost
Every dollar you spend on external LLM APIs is a variable cost that scales with usage. There is no cap baked into the architecture. You impose caps manually, reactively, after the budget has already moved.
The structural alternative is converting that variable cost into a fixed, plannable one: infrastructure you own, models you run, a bill that reads more like a data-centre line item than a taxi meter. That's the architecture change, not a configuration tweak.
Owning the stack also collapses a second problem into the same decision. Teams that can't send sensitive code or proprietary data to external APIs in the first place, like regulated industries with strict data-residency requirements, get cost control and data control from one architectural choice: when the models run inside your own perimeter, the spend is a capacity you provisioned, and the data never leaves it.
The honest objection is that owned infrastructure costs more upfront. That's true, and you should model it carefully. The break-even depends on your team size, your model mix, and how far up that power-law curve your engineers actually sit. But the Uber scenario, burning an annual budget in four months and then reaching for a blunt cap, has a specific infrastructure shape behind it: metered external APIs with no architectural ceiling.
The third that hasn't solved this yet
Look at the FinOps Foundation's numbers again. Two years ago, fewer than one in three organisations considered AI spend a FinOps concern. Today it's nearly two in three. The other third hasn't caught up yet, or they've decided the productivity gains justify the open meter.
That second position is defensible for a while, at the right scale. One company reportedly spent approximately $500 million on AI after failing to enact employee usage caps. 7 MIT research suggests roughly 95% of enterprise GenAI projects fail to deliver measurable financial returns within six months. 6 Unlimited spend on ambiguous return is a hard position to hold when the board asks.
The move that's working for teams ahead of this curve: model the cost of your specific agentic workload (use the math above as a starting point), map it against the productivity return you can actually measure, and decide whether metered external spend or fixed owned infrastructure gives you better control over that ratio. Don't let the sticker price on input tokens be the number your finance team sees.
01 Token prices fell ~80% in a year, yet bills rose, because cheaper tokens unlocked agentic workloads that burn 50× more tokens than a chat prompt. That's the Jevons paradox, and it runs on autopilot.
02 Output tokens are the variable that escapes. At Opus 4.8 rates, one power user running 40 agentic turns a day costs $1,800/month, past Uber's hard cap on a single tool.
03 Developer AI spend follows a power-law distribution. The engineers generating the most value are structurally also generating the largest bills; blunt caps cut both.
04 Per-token billing has no architectural ceiling. You impose limits manually, after the damage. The structural fix is converting variable token spend into fixed infrastructure cost.
05 63% of organisations now name AI an active FinOps concern, up from 31% two years ago. The teams ahead of this have modelled their workloads and made an explicit build-vs-buy decision.
If your team is hitting the ceiling on AI spend, dealing with data-residency constraints that rule out external APIs, or both, we'd be glad to walk through the numbers against your actual workload.
TechCrunch, "Uber caps employee AI spending after blowing through budget in four months" (June 2, 2026).
CloudZero, "LLM API Pricing Comparison" . Per-million-token input/output prices and the ~80% year-over-year decline (2026).
LeanOps, "Agentic AI cost runaway: the token budget problem" . Agents consume roughly 50× the tokens of a chat prompt (2026).
Morph LLM, "AI Coding Costs" . Median monthly token usage (~51M/developer), tokens per agentic task, and per-developer monthly spend ranges (2026).
FinOps Foundation, finops.org . Share of organisations naming AI an active FinOps concern, 31% (2024) → 63% (2025).
MIT Project NANDA, The GenAI Divide: State of AI in Business (2025). Roughly 95% of enterprise generative-AI projects show no measurable financial return within six months.
Secondary industry reporting (2026). Microsoft engineers' reported ~$2,000/month agentic token bills and a reported ~$500M unmanaged AI spend at one company. Primary sources still to be confirmed before publication.

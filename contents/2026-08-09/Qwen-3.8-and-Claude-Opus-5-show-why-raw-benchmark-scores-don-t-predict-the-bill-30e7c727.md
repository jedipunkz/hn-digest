---
source: "https://venturebeat.com/orchestration/qwen-3-8-max-and-claude-opus-5-show-why-raw-benchmark-scores-dont-predict-the-bill"
hn_url: "https://news.ycombinator.com/item?id=49229526"
title: "Qwen 3.8 and Claude Opus 5 show why raw benchmark scores don't predict the bill"
article_title: "Qwen 3.8-Max and Claude Opus 5 show why raw benchmark scores don't predict the bill | VentureBeat"
author: "ashurandi"
captured_at: "2026-08-09T08:30:23Z"
capture_tool: "hn-digest"
hn_id: 49229526
score: 1
comments: 0
posted_at: "2026-08-09T08:23:34Z"
tags:
  - hacker-news
  - translated
---

# Qwen 3.8 and Claude Opus 5 show why raw benchmark scores don't predict the bill

- HN: [49229526](https://news.ycombinator.com/item?id=49229526)
- Source: [venturebeat.com](https://venturebeat.com/orchestration/qwen-3-8-max-and-claude-opus-5-show-why-raw-benchmark-scores-dont-predict-the-bill)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T08:23:34Z

## Translation

タイトル: Qwen 3.8 と Claude Opus 5 は、生のベンチマーク スコアが請求額を予測できない理由を示しています
記事のタイトル: Qwen 3.8-Max と Claude Opus 5 は、生のベンチマーク スコアが請求額を予測しない理由を示しています |ベンチャービート
説明: ベンチマーク スコアが高いからといって、コストが低いというわけではありません。 Qwen 3.8-Max と Claude Opus 5 はどちらもそれを示しています。成功したタスクあたりのコストがそれを捉える指標です。

記事本文:
Qwen 3.8-Max と Claude Opus 5 は、生のベンチマーク スコアが請求額を予測できない理由を示しています。 VentureBeat オーケストレーション
ニュースレターの Qwen 3.8-Max と Claude Opus 5 は、生のベンチマーク スコアが請求額を予測できない理由を示しています
クレジット: VentureBeat と Gemini によって作成されました
アリババは今週 Qwen 3.8-Max をリリースし、プレビュー版を Claude Fable 5 に次ぐものとして宣伝しました (リリース日の表はさらに曖昧で、モデルはコーディング エージェントの 12 行のうちの 1 行でリードされています)。しかし、独立したハーネスは逆の結論に近づきました。明らかにプレビュー バージョンを使用したベンチマーク実行では、Qwen 3.8-Max のベスト エフォート設定がパックの真ん中にあり、デフォルト設定が最後になりました。
どちらの結果も真実であり、擁護可能です。それらの差はトークンと時間の予算に関するものであり、これらの数字は通常は見出しの数字ではないため、これは重要です。 Alibaba の脚注では、コーディング番号のタイムアウトは 5 時間、PaperBench での実行ごとに最大 12 時間と規定されています。独立したハーネスである VulcanBench を使用すると、実測時間は 45 ～ 60 分になりました。アリババ側の時間予算が 5 ～ 16 倍であることが、結果の大きな違いを説明しています。
モデルを選択する際には、これらの違いを考慮に入れるために 2 つのことを行う必要があります。まず、使用する指標は成功したタスクあたりのコストです。これは、失敗した試行に費やしたすべての費用を含む総支出を、実際に受け入れチェックに合格したタスクで割ったものです。次に、時間またはトークンの予算を、隠れた詳細ではなく、受け入れ基準の明示的な部分にする必要があります。
トークンあたりの価格が請求額を予測できなくなりました
Qwen 3.8-Max の最初の週に全員が公開した比較は価格比較でした。それが入手可能な唯一のデータだったからです。安価なモデルではありません。 7 月 31 日にパブリック API ベータ版に入った DeepSeek-V4-Flash-0731 のリストは次のとおりです。

入力トークン 100 万あたり 14 セント、出力 28 セント。 Qwen 3.8-Max の価格は 2 ドルと 6 ドルです。キミ K3 の価格は 3 ドルと 15 ドルです。
これらの価格は、Qwen のような推論モデルに特有の理由により、以前よりもわかりにくくなっています。結果に到達するには思考トークンが必要です。トークン許容量のほとんどを推論に費やすモデルは、答えを書き込む前にトークンの上限に達する可能性があり、完全な実行を犠牲にして完全な失敗と区別できない空の結果が得られます。
Artificial Analysis には、これが実際のエージェントの支出にどのような影響を与えるかについて、最も正確に公表された測定値があります。DeepSeek-V4-Flash でインテリジェンス インデックスを最大の労力で実行すると、クラス中央値の 1 億に対して 2 億 1,000 万の出力トークンが必要でした。トークンが非常に安かったため、とにかく絶対コストは低く抑えられました。しかし、冗長だとお金だけでなく時間もかかり、ユースケースによっては失敗する可能性があります。
必要なのは、指定した時間とトークン予算内で実際に完了したタスクに対して、費やしたすべてをカウントする数値です。これには、空で返された試行も含まれます。これは、成功あたりのコストのメトリクスを確認するのに役立ちます。
故障率の一部は構成設定によって決まります
間違った答えが得られる実行と予算を使い果たす実行は、修正が異なる別のイベントです。それらを区別するハーネスはほとんどなく、リーダーボードでも分裂を報告するものはほとんどありません。私はこのビルドで独自のエージェント ベンチマークを実行しました。ハーネスは障害を記録しましたが、その理由については何も記録されなかったため、自分で区別を追加する必要がありました。これらを分離すると、予算の枯渇が優先することが判明します。
7 月に公開された Long-Horizo​​n-terminal-bench は、共有ハーネスを介して 46 のタスクにわたって 17 のフロンティア モデルを 1 回の試行で 90 分間実行しました。未解決の実行ではタイムアウトが 79% を占め、一方、実行が停止したエージェントでは 19% がタイムアウトでした。

自分自身のエラーとハーネスのエラーが 3% です。著者らは、それが何を意味するのか、何を意味しないのかについて注意を払っています。タイムアウトした実行は終了間近ではなく、平均報酬は 0.10 から 0.35 の間であったため、もっと時間があれば成功したとは想定できません。しかし、教訓は、それを叫ぶかどうかにかかわらず、ベンチマークは暗黙的に時間効率を測定しているということです。
このメカニズムの最も明確に公開されている例は、Qwen チャートの背後にある同じオープンソース ハーネスである VulcanBench からのものです。 7 月 26 日付けのレポートでは、Claude Opus 5 の最も低労力の設定が最高であり、高労力では 18 のタスクに対して 23 のタスク中 20 を解決しました。追加の推論は無駄ではありませんでした。多大な労力を費やした結果、どの設定でも 1 対 3 の最も少ない不正解が返されました。代わりにクロックが不足し、タイムアウトのスコアは 0 になります。 3 つの回帰のうち 2 つは、少ない労力で解決できるタスクの削減であり、両方の時間を無制限にすると、コストは 3.1 倍という最も安価な設定に匹敵するだけです。
これは、配線はしごを構築する人にとって直接的な影響を及ぼします。標準設計では、安価な試行が失敗すると、次の段の方が優れており、コストが単に高いだけであるという前提で、より多くの推論が行われます。モデルとタスクの組み合わせの意味のあるシェアの場合、その仮定は間違っており、タイムアウトまたは上限に達するまでエスカレーションするために、より高い段階の代償を支払うことになります。
過去数か月間、いくつかのグループが成功したタスクあたりのコストを独自に算出しました。これは、それが標準になりつつある最も強力な兆候です。
VulcanBench は、解決されたタスクごとの金額をヘッドライン欄として報告しており、最初の頃から報告を行っています。 Long-Horizo​​n-terminal-bench は精度の次にタスクごとのコストを公開しており、最も有益な行は GPT-5.4 でタスクあたり約 26 ドルですが、合格率は Grok 4.5 の約 11 ドルよりもはるかに低いです。 TestEvo-Bench は次の環境でエージェントを実行します。

コスト上限があり、より厳しい上限では Claude Code のテスト生成スコアは 71% から 44% に低下します。
ベンダーは、成功したタスクごとに測定するという考えにすでに取り組んでいます。 HubSpot は 4 月に、Breeze カスタマー エージェントの料金を、処理された会話あたり 1 ドルから、解決された会話あたり 50 セントに変更しました。 Zendesk の請求は自動解決ごとに行われます。 Fin は結果ごとに 99 セントを請求し、エンドツーエンドの解決に対してのみ請求します。
必須フィールドとして、予算枯渇、ベリファイアの失敗、ハーネス エラーを 1 つの失敗フラグではなく個別の値として、エージェント実行ごとに失敗の理由を出力します。タイムアウトと誤った解答を区別できるまでは、合格率は 2 つのことを同時に測定することになり、どちらを修正すればよいのかわかりません。
モデルごとだけでなく、作業レベルごとに成功したタスクごとのコストを計算します。失敗した試行を含む合計費用を、受け入れチェックに合格したタスクで割ったもの。ランキングは料金表と一致しないため、最も安い設定が勝つ可能性があります。
遅延が本当にサービス レベルの目標に含まれていない限り、壁時計ではなくトークンに上限を設けます。壁時計の上限は、プロバイダーのサービス速度をモデルの品質として評価します。
デプロイしたすべてのデフォルトの作業量設定を確認してください。 Qwen 3.8-Max は、努力フィールドが設定されていない場合、最高の推論設定で実行され、その最高設定は独立したテストで最悪のパフォーマンスでした。そのパラメーターにまったく触れないチームは、解決されたタスクごとに最もコストがかかる構成を実行していることになります。
私の個人情報を販売または共有しないでください
私の機密個人情報の使用を制限する
© 2026 ベンチャービート。無断転載を禁じます。

## Original Extract

Higher benchmark scores don't mean lower cost. Qwen 3.8-Max and Claude Opus 5 both show it — and cost per successful task is the metric that catches it.

Qwen 3.8-Max and Claude Opus 5 show why raw benchmark scores don't predict the bill | VentureBeat Orchestration
Newsletters Qwen 3.8-Max and Claude Opus 5 show why raw benchmark scores don't predict the bill
Credit: Made by VentureBeat with Gemini
Alibaba released Qwen 3.8-Max this week and marketed the preview as second only to Claude Fable 5 (their launch-day table was more equivocal: the model leads on one of 12 coding-agent rows). But an independent harness came close to the opposite conclusion: a benchmark run , apparently using the Preview version, put Qwen 3.8-Max's best effort setting mid-pack, and its default setting last.
Both results are real and defensible. The gap between them is about token and time budgets, and that matters because those figures aren’t usually headline numbers. Alibaba's footnotes give its coding numbers a five-hour timeout, and up to 12 hours per run on PaperBench. The independent harness, VulcanBench, allowed between 45 and 60 minutes of wall clock time . A time budget between five and 16 times larger on Alibaba’s side explains the huge difference in results.
It’s time to do two things to start accounting for these differences when choosing models. First, the metric to use is cost per successful task: total spend, including everything you spent on attempts that failed, divided by the tasks that actually passed your acceptance check. Second, you need to make time or token budgets an explicit part of your acceptance criteria, not a hidden detail.
Price per token has stopped predicting the bill
The comparison everyone published in Qwen 3.8-Max's first week was a price comparison, because that was the only data available. It is not a cheap model. DeepSeek-V4-Flash-0731, which entered public API beta on July 31, lists at 14 cents per million input tokens and 28 cents output . Qwen 3.8-Max lists at $2 and $6. Kimi K3 sits at $3 and $15.
Those prices tell you less than they used to, for a reason specific to reasoning models like Qwen: getting to a result costs thinking tokens. A model that spends most of its token allowance on reasoning can reach a token cap before it writes the answer, giving you an empty result indistinguishable from a total failure at the cost of a full run.
Artificial Analysis has the cleanest published measurement of how this can affect real agent spend: running its Intelligence Index on DeepSeek-V4-Flash at maximum effort took 210 million output tokens against a class median of 100 million. Absolute cost stayed low anyway, because the tokens were so cheap. But verbosity costs time, not just money, and depending on your use case that can sink you.
What you need is a number that counts everything you spent, including the attempts that came back empty, against the tasks that actually got done in the time and token budget you specified. This is what a cost-per-success metric helps you see.
Your failure rate is partly a configuration setting
A run that produces a wrong answer and a run that runs out of budget are different events with different fixes. Almost no harness distinguishes them, and almost no leaderboard reports the split. I hit this building an agent benchmark of my own : the harness logged a failure and nothing about why, and I had to add the distinction myself. When you do separate them, budget exhaustion turns out to dominate.
Long-Horizon-Terminal-Bench , published in July, ran 17 frontier models across 46 tasks through a shared harness with one 90-minute attempt each. Timeouts accounted for 79% of unresolved runs, against 19% for agents that stopped on their own and 3% for harness errors. The authors are careful about what that does and does not mean: the timed-out runs were not close to finishing, with mean reward between 0.10 and 0.35, so you cannot assume more time would have resulted in success. But the lesson is: benchmarks are implicitly measuring time efficiency, whether or not they shout about that.
The clearest published example of the mechanism comes from VulcanBench, the same open-source harness behind the Qwen chart. In a report dated July 26 , Claude Opus 5's lowest-effort setting was its best, solving 20 of 23 tasks against 18 at high effort. The extra reasoning wasn’t useless: high effort returned the fewest wrong answers of any setting, one against three. It ran out of clock instead, and a timeout scores zero. Two of its three regressions were cutoffs on tasks that low effort solves, and given unlimited time on both it only ties its cheapest setting, at 3.1 times the cost.
That has a direct consequence for anyone building a routing ladder. The standard design escalates to more reasoning when a cheap attempt fails, on the assumption that the next rung is better and merely costs more. For a meaningful share of model and task combinations that assumption is wrong, and you pay the higher rung's price to escalate into a timeout or hitting a cap.
Several groups have landed on cost per successful task independently in the last few months, which is the strongest signal it's becoming standard.
VulcanBench reports dollars per solved task as a headline column and has since its earliest reports . Long-Horizon-Terminal-Bench publishes per-task cost next to accuracy, and its most instructive row is GPT-5.4 at roughly $26 per task with a much lower pass rate than Grok 4.5 at about $11. TestEvo-Bench runs agents under a cost cap, and Claude Code's test-generation score falls from 71% to 44% at the tighter cap.
Vendors are already on board with the idea of measuring per successful task. HubSpot moved its Breeze Customer Agent in April to 50 cents per resolved conversation , down from $1 per handled conversation. Zendesk bills per automated resolution . Fin charges 99 cents per outcome and bills only on end-to-end resolution.
Emit a failure reason on every agent run as a required field, with budget exhaustion, verifier failure and harness error as distinct values rather than one failure flag. Until you can separate a timeout from a wrong answer, your pass rate is measuring two things at once and you cannot tell which one to fix.
Compute cost per successful task per effort level, not just per model. Total spend including failed attempts, divided by tasks that passed your acceptance check. The ranking will not match the rate card, and the cheapest setting may well win.
Cap on tokens rather than wall clock unless latency is genuinely in your service level objective. A wall-clock cap scores your provider's serving speed as model quality.
Check the default effort setting on everything you have deployed. Qwen 3.8-Max runs at its highest reasoning setting when the effort field is unset , and its highest setting was its worst performer in independent testing. A team that never touches that parameter is running the configuration that costs the most per solved task.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.

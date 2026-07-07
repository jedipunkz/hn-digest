---
source: "https://www.cubic.dev/state-of-ai-coding-2026"
hn_url: "https://news.ycombinator.com/item?id=48820026"
title: "Codex makes fewer bugs, but more people use Claude"
article_title: "State of AI coding 2026"
author: "mikeborozdin"
captured_at: "2026-07-07T17:06:48Z"
capture_tool: "hn-digest"
hn_id: 48820026
score: 5
comments: 0
posted_at: "2026-07-07T16:24:51Z"
tags:
  - hacker-news
  - translated
---

# Codex makes fewer bugs, but more people use Claude

- HN: [48820026](https://news.ycombinator.com/item?id=48820026)
- Source: [www.cubic.dev](https://www.cubic.dev/state-of-ai-coding-2026)
- Score: 5
- Comments: 0
- Posted: 2026-07-07T16:24:51Z

## Translation

タイトル: Codex ではバグが少なくなりますが、Claude を使用する人が増えています
記事タイトル: 2026 年の AI コーディングの現状
説明: AI はすでにコードの大部分を記述しています。ここでさらに興味深い質問は、どのモデルが最もミスが少ないかということです。私たちは、最もパフォーマンスの高い AI エンジニアリング チームの数十万件のコミットを分析し、モデル間に明らかなギャップがあることを発見しました。

記事本文:
AI モデルは、キュービック エージェントが前四半期にレビューしたコードの 90% を作成しました。最もバグが少ないのはどれですか?
私たちは答えることができる立場にあります。 Cubic は、Resend、n8n、Daytona などのチームの 1 日に何千ものコミットをレビューしており、その多くは CLI を使用して AI の帰属を追跡しています。
したがって、すべてのコミットについて、どのモデルがそれを書いたのか、そしてそこにバグが見つかったかどうかがわかります。ベンチマークは、モデルが何を解決できるかを示します。このデータは、何が壊れているかを示します。
短いバージョン - GPT-5.5 は、これまでに見た中で最もクリーンなコードを記述します。これまでのところ、開発者が使用するものは変わっていません。開発者の 80% は Claude Code を使用してコードを作成しており、私たちがレビューするコミットのほとんどは Opus モデルから来ています。 *
GPT-5.5 では、データセット内の他のモデルよりも、変更された行ごとに発生するバグが少なくなります。
Claude Code は依然として最も人気のあるコーディング エージェントであり、開発者の 80% が使用しています。
Opus ファミリは最もよく使用されているモデルのセットです。開発者の 80% が毎週、Opus モデルをプライマリ モデルとして実行しています。
Claude Fable は、cubic の歴史の中で最も早く採用されたモデルであり、開発者の 30% がリリース後 48 時間以内に使用しました。
モデルの品質: GPT-5.5 は他のどのモデルよりもバグが少ない
新しいモデルはそれぞれ、以前のモデルよりも発生するバグが少なくなります (Opus 4.7 は Opus 4.6 の半分のバグを発生しました)。ただし、1 つの例外があります。Opus 4.8 は 4.7 よりも高いバグ率を示します。明確な説明はありません。
候補の 1 つは、開発者が最初の数週間で新しい旗艦に最も困難なタスクを投入するため、初期のバグ数が膨らむということです。
GPT-5.5 は、主力の Anthropic モデルよりも発生するバグが少なく、バグが最も少ないモデルです。
開発者は 1 年前と比べて 3 倍のコードをマージ
この典型的な開発者は、マージされた PR の数を過去 1 年間でおよそ 2 倍にし、2025 年半ばには月あたり約 6 件でしたが、現在では約 12 件になりました。これらの PR は中央値でも 2.2 倍大きくなります。
要約すると、開発する

ユーザーは 1 年前に比べて約 3 倍のコードをマージしており、開発者の生産量は大幅に増加しています。 AI は人間よりも冗長で定型的なコードを生成する傾向がありますが、その影響を正確に測定することは困難ですが、提供される価値は AI の導入によって明らかに増加しました。
開発者がマージしたコード変更は 2026 年 5 月には中央値 5,820 行でしたが、1 年前の 2025 年 6 月には 1,930 行でした。 3倍の増加です。
開発者ごとにマージされた PR は、過去 1 年間で 2 倍になり、月あたり中央値 6 件 (2025 年 6 月) から 12 件 (2026 年 5 月) に増加しました。
PR も大きくなり、2025 年 6 月には PR ごとに変更されたコードの中央値 56 行から、2026 年 5 月には 124 行となり、2.2 倍に増加しました。
コーディングエージェント: クロードが支配的
Claude Code は開発者の 80% でリードしています。 Cursor は 18% に留まり、毎週シェアを失っています。コーデックスは 17% で安定しています。多くの開発者が複数のエージェントを実行しているため、パーセンテージの合計は 100 を超えます。
モデルの使用: 新しいフロンティア モデルの勝利
Anthropic の Opus モデル ファミリは明らかに開発者の間で最も人気があり、開発者の 80% が毎週これを主要モデルとして使用しています。 OpenAI の GPT ファミリは 2 位で、GPT-5.5 のリリース後に採用が増加しています。
前週比データからは、新しい主力フロンティア モデルが発売後 2 週間以内にすぐに採用されるという明確なパターンが明らかになりました。
寓話の短命の物語
モデルシェアチャートの紫色の塊はクロード・ファブルです。 Anthropic は 6 月 9 日火曜日にこれをリリースしました。木曜日までに、プラットフォーム上の開発者の 30% がそれを使用しており、cubic 史上最速のモデル採用となりました。
それからそれは消えました。その週の金曜日に米国政府が撤退を命令し、アンスロピックはモデルを撤回した。わずか半週間しか入手できなかったにもかかわらず、主要モデルのシェアの 15% を獲得しました。
そもそも人間はコードを書いているのでしょうか？

キュービック CLI を使用してアトリビューションを追跡するチームでは、現在 AI がコードの大部分 (90 ～ 100%) を作成しており、4 月中旬の 78% から増加しています。これらは業界で最も AI を推進しているチームであるため、これを業界平均ではなく先行指標として扱います。
cubic CLI を使用して、チームの AI コーディング アクティビティと帰属を追跡できます。
* データセットは、少なくとも 2 人の開発者からなるチームによって作成された数十万のコミットに基づいています。これらのメトリクスは、CLI を介して AI アトリビューションが追跡されたコミットのみに焦点を当てており、すべてのコミットの一部のみを表しています。私たちは 2 月中旬頃からこのデータの追跡を開始しました。これはコードベース全体を表すものではありませんが、新たなパターンに対する説得力のある方向性を提供します。
AI により開発者の生産量が 3 倍に増加
寓話の短命の物語
そもそも人間はコードを書いているのでしょうか？
© 2026 立方体.無断転載を禁じます。
© 2026 立方体.無断転載を禁じます。

## Original Extract

AI is already writing most of the code. The more interesting question now is: which model makes the fewest mistakes? We analysed hundreds of thousands of commits of the best-performing AI engineering teams and found a clear gap between models.

AI models wrote 90% of the code that cubic agents reviewed last quarter. Which one produced the fewest bugs?
We're in a good position to answer. Cubic reviews thousands of commits a day for teams like Resend, n8n, and Daytona, and many of them track AI attribution with our CLI.
So for every commit, we know which model wrote it, and whether we found bugs in it. Benchmarks tell you what a model can solve; this data tells you what it breaks.
The short version - GPT-5.5 writes the cleanest code we see. So far that hasn't changed what developers use: 80% of developers code with Claude Code, and most of the commits we review come from Opus models. *
GPT-5.5 produces fewer bugs per lines changed than any other model in our dataset.
Claude Code remains the most popular coding agent, used by 80% of developers.
The Opus family is the most-used set of models: 80% of developers run an Opus model as their primary model each week.
Claude Fable was the fastest-adopted model in cubic's history: 30% of developers were using it within 48 hours of release.
Model quality: GPT-5.5 produces fewer bugs than any other model
Each new model produces fewer bugs than its predecessor (Opus 4.7 generated half the bugs of Opus 4.6), with one exception: Opus 4.8 shows a higher bug rate than 4.7. We don't have a firm explanation.
One candidate: developers throw their hardest tasks at a new flagship in its first weeks, which inflates early bug counts.
GPT-5.5 produces fewer bugs than flagship Anthropic models and is the least buggy model.
Developers merge 3× more code than a year ago
The typical developer has roughly doubled the number of merged PRs over the past year, from about 6 per month in mid-2025 to around 12 today. Those PRs are also 2.2× larger at the median.
In summary, developers are merging roughly three times more code than they were a year ago, a substantial increase in developer output. AI tends to generate more verbose and boilerplate-heavy code than humans, but although the impact is difficult to measure precisely, the value delivered has clearly increased with the adoption of AI.
Developers merged a median of 5820 lines of code changes in May 2026, compared to 1930 in June 2025, a year ago. That's a 3x increase.
PRs merged per developer doubled in the past year, from a median of 6 per month (June 2025) to 12 (May 2026).
PRs are getting bigger too, from a median of 56 lines of code changed per PR in June 2025 to 124 in May 2026, a 2.2x increase.
Coding agents: Claude dominates
Claude Code leads at 80% of developers. Cursor sits at 18% and loses share every week; Codex holds steady at 17%. The percentages sum past 100 because many developers run more than one agent.
Model usage: new frontier models win
Anthropic's Opus model family is clearly the most popular among developers, with 80% of developers using it as their primary model every week. OpenAI's GPT family comes in second and has seen increased adoption following the release of GPT-5.5.
The week-over-week data reveals a clear pattern: new flagship frontier models are quickly adopted within two weeks of release.
The short-lived story of Fable
The purple blob in the model-share chart is Claude Fable. Anthropic released it on Tuesday, June 9. By Thursday, 30% of developers on the platform had used it, the fastest adoption of any model in cubic's history.
Then it vanished. On Friday of that week, the U.S. government ordered its withdrawal, and Anthropic pulled the model. Even with only half a week of availability, it captured 15% of primary-model share.
Are humans writing code at all?
Among teams that track attribution with the cubic CLI, AI now authors the vast majority of code (90–100%), up from 78% in mid-April. These are the most AI-forward teams in the industry, so treat this as a leading indicator, not an industry average.
You can track AI coding activity and attribution for your team using the cubic CLI .
* The dataset is based on hundreds of thousands of commits made by teams of at least 2 developers. These metrics focus exclusively on commits where AI attribution was tracked via our CLI, representing only a fraction of all commits. We started tracking this data around mid-February. While this doesn't represent the total codebase, it offers a compelling directional look at emerging patterns.
AI has tripled developer output
The short-lived story of Fable
Are humans writing code at all?
© 2026 cubic. All rights reserved.
© 2026 cubic. All rights reserved.

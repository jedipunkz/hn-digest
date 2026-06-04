---
source: "https://www.anthropic.com/institute/recursive-self-improvement"
hn_url: "https://news.ycombinator.com/item?id=48400842"
title: "When AI Builds Itself: Our progress toward recursive self-improvement"
article_title: "When AI builds itself \\ Anthropic"
author: "meetpateltech"
captured_at: "2026-06-04T18:55:22Z"
capture_tool: "hn-digest"
hn_id: 48400842
score: 59
comments: 70
posted_at: "2026-06-04T16:20:17Z"
tags:
  - hacker-news
  - translated
---

# When AI Builds Itself: Our progress toward recursive self-improvement

- HN: [48400842](https://news.ycombinator.com/item?id=48400842)
- Source: [www.anthropic.com](https://www.anthropic.com/institute/recursive-self-improvement)
- Score: 59
- Comments: 70
- Posted: 2026-06-04T16:20:17Z

## Translation

タイトル: AI が自らを構築するとき: 再帰的な自己改善に向けた私たちの進歩
記事タイトル: AI が自らを構築するとき \ Anthropic
説明: 再帰的な自己改善に向けた私たちの進歩とその意味。

記事本文:
AI が自分自身を構築するとき \ Anthropic メインコンテンツにスキップ フッターにスキップ 研究
再帰的な自己改善に向けた私たちの進歩とその意味。
AI の歴史のほとんどにおいて、その開発サイクルのあらゆる段階を人間が主導してきました。しかし、Anthropic では、AI 開発の割合を AI システム自体に委任する割合が増えており、それによって私たちの作業がスピードアップしています。
十分に考慮し、十分なコンピューティングを考慮すると、この傾向は、完全に自律的に独自の後継システムを設計および開発できる AI システムを示しています。これは再帰的自己改善と呼ばれます。私たちはまだそこには到達しておらず、再帰的な自己改善は避けられないわけではありません。しかし、それはほとんどの機関が準備しているよりも早く起こる可能性があります。
Anthropic Institute は、公開ベンチマークと Anthropic 内の未報告データを使用して、AI がすでに AI システムの開発を加速していることを示しています。ほんの 1 つの例を挙げると、現在、Anthropic のエンジニアは、2021 年から 2025 年に比べて四半期あたり平均 8 倍のコードを出荷しています。
この記事で説明されている技術トレンドは、AI システムが今後数年間でさらに高性能になることを示唆しています。これらの傾向は大きな影響を及ぼします。自ら構築できる AI はテクノロジーの歴史の中で大きな進歩となり、科学、医療、その他の分野で世界に多大な利益をもたらす可能性があります。しかし、完全に再帰的な自己改善を行うと、人間が AI システムを制御できなくなるリスクも高まる可能性があります。システムが独自の後継システムを完全に構築できるようになった場合、システムを保護し、監視し、その動作を形成する方法のすべてがより重要になります。
初期の頃、Anthropic での仕事は他のテクノロジー企業の仕事と同じように見えました。ラップトップでコードやドキュメントを書く人たちでした。
人々は初期のチャットボットを使用して、

コード スニペットを切り出し、出力をテキスト エディタにコピーします。
エージェントの能力が高まるにつれて、自分自身でコードを作成および編集できるようになり、場合によってはファイル全体を編集できるようになりました。
エージェントは自分自身でコードを実行し、何時間もの作業を他のエージェントに委任できるようになりました。
将来的には、エージェントが自分自身でモデルを構築してトレーニングできるようになる可能性があります。これが起こった場合、Claude の将来のバージョンは Claude 自身によって継続的に改善される可能性があります。
外の世界からの証拠
AI モデルの改良速度は加速しています。彼らが自分たちで確実に完了できるタスクの長さは、7 か月ごとに 2 倍になるという以前の傾向から、およそ 4 か月ごとに倍増しています。 2024 年 3 月には、クロード オーパス 3 は、人間が完了するまでに約 4 分かかるソフトウェア タスクを完了できるようになります。 1 年後、Claude Sonnet 3.7 は約 1 時間半かかるタスクを管理しました。その 1 年後、Claude Opus 4.6 は 12 時間のタスクを管理しました。 1 この傾向が維持される場合、熟練した人であれば数日かかるタスクが今年実現可能になる可能性があります。 2027 年には、AI システムは 1 人で数週間かかるタスクを実行できるようになる可能性があります。
同じパターンがコーディングと研究のベンチマークにも現れます。ベンチマークは特定のドメインにおけるモデルのパフォーマンスを測定し、モデルが 100% に近いパフォーマンスを達成するとベンチマークは「飽和」します。 2 SWE ベンチは、現実世界のソフトウェア エンジニアリングの標準テストです。モデルに実際のオープンソース コードベースと実際のバグ レポートを渡し、問題を修正してプロジェクト独自のテストに合格するコード変更を記述するように要求します。モデルのスコアは 1 桁前半から 2 年間でベンチマークが飽和するまでになりました。
CORE-Bench は、モデルが既存の研究を再現できるかどうかをテストします。これは、独自の研究を行うための前提条件です。 AI モデルに公開の背後にあるコードとデータを提供します

論文を流し、すべてを再実行して論文の結果を再現できることを確認するように依頼します。 AI システムは、2024 年に約 20% の確率で結果を再現することに成功しましたが、15 か月後にはベンチマークが飽和状態になりました。モデルが長時間のタスクをどれだけうまく完了できるかを測定するベンチマークを実行している METR は、Claude Mythos Preview が「少なくとも」 16 時間動作でき、「新しいタスクなしで [METR] が測定できる時間の上限に達している」ことを発見しました。
公開ベンチマークは、これらのシステムの機能について多くを語っています。しかし、AI システムが AI 開発自体のスピードアップに与えている影響を明らかにすることはできません。そのためには、Anthropic のような AI 企業内からの直接的な証拠が必要です。
人間内部からの証拠
フロンティア モデルの構築には、大きく 2 つのカテゴリの作業が必要です。エンジニアリングには、コードの作成、インフラストラクチャの構築、モデルのトレーニングの監督などがあります。そして、どのような実験を実行するかを決定し、得られた結果を解釈し、次にどのアイデアを試すべきかを判断するという研究もあります。
エンジニアリングと研究の両方にわたって、この状況は一貫しています。エンジニアリングでは、クロードは不明確な問題を与えられ、それを解決する方法を見つけることができます。人間は目標を提供しますが、方法を提供する必要はなくなりました。研究においては、クロードは、明確に指定された実験の実行において、すでに熟練した人間と同等か、それを上回る成績を収めています。しかし、エンジニアリングと研究の両方で目標を選択する際にクロードが判断力を行使する場合、パフォーマンスには大きなギャップが残ります。それが、現在の AI と、独自の後継システムを自律的に設計できる将来のシステムとの間のギャップです。
Anthropic の従業員は、経験を積むにつれて、より自由で重要なタスクを任されるのが一般的です。早い段階で、彼らは誰かが指定したタスクを実行します。

「エクスポートボタンが機能しないので修正してください。」経験を積むと、「高負荷時にネットワークが遅くなる理由を調査する」などの目標を与えられ、自分たちでアプローチを設計するようになります。最上級レベルでは、どの問題に取り組む価値があるのか​​、つまり「チームは次の四半期に何を構築すべきか?」を決定しています。内部の Anthropic データを使用して、クロードがこれらのさまざまな種類のタスクを処理できるようにどこまで到達したかを確認できます。
クロードは、Anthropic のコードのかなりの部分を作成します。 2026 年 5 月の時点で、Anthropic のコードベースにマージされるコードの 80% 以上はクロードによって作成されました。 3 2025 年 2 月に Claude Code がリサーチ プレビューで開始されるまで、この数値は 1 桁前半でした。その変化はエンジニア一人当たりの生産量にも現れます。エンジニアごとにマージされる 1 日あたりのコード行数は、Anthropic の最初の 4 年間 (2021 年から 2024 年) までは一定でしたが、クロードがエンジニアにコピー アンド ペーストするコードを単に提案するのではなく実行し始めた 2025 年に増加し始めました。モデルがより長い期間にわたって自律的に動作し始めた 2026 年に、傾きは再び急になりました。これら 2 つの変曲点を以下のグラフに示します。 2026 年の第 2 四半期には、一般的なエンジニアは 2024 年の 8 倍の量のコードを 1 日あたりマージしていました 4。これは、コードの多くはクロードによって書かれており、エンジニアが自分で入力するのではなく、指示とレビューを行っているためです。
注意: コードの行数は、品質よりも量を評価するため、不完全な尺度です。したがって、2026 年の第 2 四半期における 8 × エンジニア 1 日あたりのコード行数は、ほぼ確実に真の生産性の向上を過大評価しています。それにもかかわらず、それは加速を示します。 Anthropic では、コードを何行書いたかによって報酬を与えることはありません。むしろ、

チームのメンバーがより多くのコードを作成しているのは、単に AI システムを使用してより多くのコードを作成しているためです。
記述されるコードの行数の増加は、生産性が大幅に向上するという主観的な印象と一致します。 Anthropic 研究チーム全体の 130 人の従業員を対象とした 2026 年 3 月の世論調査では、回答者の中央値は、どのような種類のプロジェクトに取り組んでいたかに関係なく、AI モデルにアクセスしなかった場合に比べて、Mythos Preview を使用すると約 4 倍の成果を生み出したと推定しました。 5 3 月の実際の上昇度合いは若干低かったと予想されます。 6 それにもかかわらず、私たちは全体的な主張がもっともらしいと考えており、他の観察と一致しています。つまり、Anthropic の技術スタッフのかなりの部分が、AI の支援なしで行うよりも数倍の速さで中核的な業務を遂行しています。
また、Anthropic の人々が、探索ツールの構築や長年延期されていたクリーンアップへの対処など、他の方法では実現しなかった作業を行うために Claude を使用しているという証拠も見られます。たとえば、2026 年 4 月に、Claude は 800 を超える修正を出荷し、あるクラスの API エラーを 1,000 分の 1 に削減しました。クロードを監督していた技術者は、人間がこの作業を完了するには 4 年かかるだろうと見積もっていました。他人のバグを解決するのは時間がかかり骨が折れる作業であり、人間は一度にそのような見慣れないコンテキストを頭の中に保持するのに苦労します。
私が Claudifying に本格的に傾倒し始めたのは約 1 年前です。それはクレイジーな冒険でした。最後に自分でコードを書いてから約 5 か月が経ちました。 」
クロードが書いたコードは「良好」であり、改善されています。 「良いコード」とは 2 つのことを意味します。それは機能すること、もう 1 つは他のエンジニアがそれを理解し、その上に構築できる方法で書かれていることです。最初の基準に関しては、証拠は明らかです。 Anthropic st の割合

クロードからのタスクの修正、リダイレクト、または引き継ぎは、最も複雑で無制限のタスクを含め、ここ 1 年間着実に減少しています。これは、明確な仕様がなく、エンジニアも答えがどのようなものかわからない問題を意味します。これは、以下のグラフに示すように、さまざまな難易度のタスクにおけるクロードの経時的な成功率を見れば明らかです。クロードは機能するコードを書きます。
最も自由度の高いタスクに関して、クロードの成功率は 2026 年 5 月に 76% に達し、6 か月で 50 パーセントポイント上昇しました。この難易度のタスクの例を挙げると、定期的なアップグレードにより数万のトレーニング ジョブがクラッシュし始めました。エンジニアはクロードに、テキスト コンテンツとクラスター アクセスのみを含むライブ インシデントを指摘しました。クロードは、実行中のジョブを処理し、一度に 1 つの環境設定をテストしながら、クラッシュの原因となっている単一の不明瞭なデバッグ フラグを特定し、それを確実に再現し、修正を確認しました。クロードは、通常 2 ～ 3 日かかる作業を約 2 時間で完了させました。
2 番目の基準は、他のエンジニアが理解して構築できるコードを書くことです。ここでは人間とAIの差は依然として続いていますが、急速に縮まりつつあります。 Anthropic のスタッフの間で完全なコンセンサスは得られていませんが、多くの人は、2025 年後半の Anthropic で人間が書いたコードよりもクロードが書いたコードの品質が依然として悪く、現在ではほぼ同等であると信じています。年内には改善されると予想しています。
これにより、Anthropic が独自のコードをレビューする方法が変わりました。コードベースに提案された変更は、マージ前にバグ、セキュリティ上の欠陥、その他の欠陥を探す自動クロード レビュアーによって読み取られます。このツールを使用して遡及分析を実行したところ、コードベースに対するすべての変更の自動化されたクロード レビューが、

uld は、claude.ai 上の過去のインシデントの背後にあるバグの約 3 分の 1 を、本番環境に到達する前に発見しました。そのコードを書いたエンジニアは、これらのシステムの構築において世界最高のエンジニアの一人です。クロードは今、彼らが見逃していた間違いに気づき始めています。
クロードが書いたコードは、2025 年後半の Anthropic で人間が書いたコードよりも多少悪かったですが、現在はほぼ同等であり、厳密には年内には改善されると予想されます。 」
クロードは、誰かが設定した目標を達成するために実験を行うのが得意です。 Anthropic がモデルをリリースするたびに、同じテストを実行します。小規模な AI モデルをトレーニングするコードをクロードに与え、同じ正確性チェックに合格しながらそのコードをできるだけ高速に実行するように依頼します。目標と成功指標は事前に決まっているため、クロードの仕事は、コードを書き直し、実行し、タイミングを計り、それを繰り返すことで高速化を見つけることです。これは実験的研究ループのミニチュア版です。 2025 年 5 月、Claude Opus 4 は、開始コードと比べて平均で約 3 倍の高速化を達成しました。 2026 年 4 月までに、Claude Mythos Preview は最大 52 倍を達成しました。キャリブレーションの場合、熟練した人間の研究者が 4 倍に到達するには 4 ～ 8 時間かかります。 7 研究ワークフローのこの部分、つまり明確に定義された実験内のステップの最適化において、Claude は非常に役立つものから、

[切り捨てられた]

## Original Extract

Our progress toward recursive self-improvement, and its implications.

When AI builds itself \ Anthropic Skip to main content Skip to footer Research
Our progress toward recursive self-improvement, and its implications.
For most of AI’s history, humans drove every step in its development cycle. But at Anthropic, we are delegating a growing share of AI development to AI systems themselves, which is speeding up our work.
Taken far enough, and given enough compute, that trend points to an AI system capable of fully autonomously designing and developing its own successor. This is called recursive self-improvement . We are not there yet, and recursive self-improvement is not inevitable. But it could come sooner than most institutions are prepared for.
Using public benchmarks and previously unreported data from within Anthropic, The Anthropic Institute is showing that AI is already accelerating the development of AI systems. To take just one example: today, Anthropic engineers on average ship 8x as much code per quarter as they did from 2021-2025.
The technical trends discussed in this piece suggest that AI systems are going to become much more capable in coming years. These trends have huge implications. AI that can build itself would be a major development in the history of technology—one that could bring enormous good for the world in science, healthcare, and beyond. But full recursive self-improvement also might increase the risks of humans losing control over AI systems. If systems are capable of fully building their own successors, the ways we secure them, monitor them, and shape their behavior all grow much more important.
In the early days, work at Anthropic looked like work at any other tech company: people writing code and docs on laptops.
People used early chatbots to help with parts of the process, like generating short code snippets and copying the output into text editors.
As the agents became more capable, they were able to write and edit code on their own, sometimes entire files.
Agents can now run code themselves and delegate hours of work to other agents.
In the future, agents could become capable enough to build and train models themselves. If this happens, future versions of Claude could be continuously improved by Claude itself.
Evidence from the outside world
The rate at which AI models improve is accelerating. The length of tasks that they can reliably complete on their own has been doubling roughly every four months, up from an earlier trend of doubling every seven months. In March 2024, Claude Opus 3 could complete software tasks that take humans about four minutes to complete. A year later, Claude Sonnet 3.7 managed tasks that took about an hour and a half. A year after that, Claude Opus 4.6 managed 12-hour tasks. 1 If this trend holds, tasks that take a skilled person days could come into range this year. In 2027, AI systems could be capable of tasks that take a person weeks.
The same pattern appears on coding and research benchmarks. Benchmarks measure the performance of models in a given domain, and they’re “saturated” when models achieve close to 100% performance. 2 SWE-bench is a standard test of real-world software engineering: it hands a model an actual open-source codebase and a real bug report, and asks it to write a code change that fixes the issue and passes the project’s own tests. Models have gone from scoring in the low single digits to saturating the benchmark in two years.
CORE-Bench tests whether a model can reproduce existing research, a prerequisite for them to conduct original research. It gives an AI model the code and data behind a published paper, and asks it to rerun everything and confirm it can replicate the paper’s results. AI systems went from succeeding at reproducing the results roughly 20% of the time in 2024 to saturating the benchmark fifteen months later. METR, which runs the benchmark measuring how well models can complete long-duration tasks, found that Claude Mythos Preview could work for “at least” 16 hours and was “at the upper end of what [METR] can measure without new tasks.”
Public benchmarks say a lot about the capabilities of these systems. But they can’t reveal the impact AI systems are having on speeding up AI development itself. For that, we need direct evidence from within AI companies like Anthropic.
Evidence from within Anthropic
Building a frontier model takes two broad categories of work. There is engineering : writing the code, standing up the infrastructure, and overseeing the model training. And there is research : deciding what experiments to run, interpreting what comes back, and figuring out which ideas to try next.
Across both engineering and research, the picture is consistent. In engineering, Claude can be handed an underspecified problem and figure out how to solve it; humans supply the goal, but they no longer need to supply the method. In research, Claude can already match or outperform skilled humans at executing a well-specified experiment. However, large performance gaps persist when it comes to Claude exercising judgement in choosing goals in both engineering and research. That’s the gap between AI today and a future system that could autonomously design its own successor.
It’s common for employees at Anthropic to receive more open-ended and important tasks as they gain more experience. Early on, they execute a task someone else specified, like, “The export button isn’t working, please fix it.” With experience, they’re handed a goal and design the approach themselves, such as, “Investigate why the network slows down under heavy load.” At the most senior levels, they are deciding which problems are worth working on at all: “What should the team build next quarter?” We can use internal Anthropic data to see how far Claude has come in being able to handle these different kinds of tasks.
Claude writes a significant proportion of Anthropic’s code. As of May 2026, more than 80% of the code we merge into Anthropic’s codebase was authored by Claude. 3 Before Claude Code launched in research preview in February 2025, this number was in the low single digits. That shift also shows up in the amount of output per engineer. Lines of code merged per engineer per day stayed constant through Anthropic’s first four years (2021-2024), then began to climb upward in 2025 when Claude began to run code rather than just suggesting it for an engineer to copy and paste. The slope steepened again in 2026 when models began to work autonomously over longer time horizons. These two inflection points are shown in the chart below. In the second quarter of 2026, the typical engineer was merging 8× as much code per day as they were in 2024. 4 This is because much of the code is written by Claude, with the engineer directing and reviewing, rather than typing it themselves.
A caveat: Lines of code is an imperfect measure, as it measures quantity over quality. So 8 × lines of code/engineer/day in the second quarter of 2026 is almost certainly an overstatement of the true productivity gain. Nonetheless, it indicates an acceleration. At Anthropic, we don’t reward people for how many lines of code they write; rather, team members are producing more code simply because they’re using AI systems to write more code.
The increase in lines of code written lines up with subjective impressions of large productivity increases. In a March 2026 poll of 130 employees from across Anthropic research teams, the median respondent estimated that they produced around 4x as much output with Mythos Preview as they would have without access to any AI models, on the kinds of projects they would have been working on regardless. 5 We expect that the true degree of uplift in March was somewhat lower. 6 Nevertheless, we find the overall claim plausible, and in line with our other observations: a significant fraction of Anthropic technical staff is accomplishing their core work multiple times faster than they could without AI assistance.
We also see evidence that people at Anthropic are using Claude to do work that simply wouldn’t have happened otherwise, like building exploratory tooling and addressing long-deferred cleanup. For example, in April 2026, Claude shipped over 800 fixes that reduced a class of API errors by a factor of one thousand. The engineer overseeing Claude estimated that a human would have taken four years to complete this work; solving other people’s bugs is slow and painstaking, and humans struggle to hold that much unfamiliar context in their head at once.
I started leaning hard into Claudifying about a year ago. That’s been a crazy adventure and it’s now been ~5 months since I last wrote any code myself. ”
The code that Claude writes is “good” and improving. “Good code” means two things: it works, and it is written in a manner that allows another engineer to understand it and build upon it. On the first criterion, the evidence is clear. The rate at which Anthropic staff correct, redirect, or take over mid-task from Claude has been falling steadily for a year, including on the most complex and open-ended tasks. This means problems with no clear specification, where the engineer isn’t sure what the answer looks like. This is evident in Claude’s success rate over time on tasks of different difficulties, as shown in the graph below. Claude writes code that works.
On the most open-ended tasks, Claude’s success rate reached 76% in May 2026, up 50 percentage points in six months. To give an example of tasks in this difficulty tier, a routine upgrade began crashing tens of thousands of training jobs. An engineer pointed Claude at the live incident with little more than some text content and cluster access. Working through the running jobs and testing one environment setting at a time, Claude isolated the single obscure debugging flag that was triggering the crash, reproduced it reliably, and confirmed a fix. In about two hours, Claude delivered what would normally be two to three days of work.
The second criterion is writing code that another engineer can understand and build on. Here the gap between humans and AI persists, but is closing fast. There isn’t full consensus among staff at Anthropic, but many believe that the Claude-written code was still worse in quality than human-written code at Anthropic in late 2025, and is roughly at parity today. We expect it to be better within the year.
This has changed the way that Anthropic now reviews its own code. Proposed changes to our codebase are now read by an automated Claude reviewer that looks for bugs, security flaws, and other defects before it can merge. Using this tool, we ran a retrospective analysis, and found that an automated Claude review of every change to our codebase would have caught roughly a third of the bugs behind past incidents on claude.ai before they ever reached production. The engineers who wrote that code are among the best in the world at building these systems. Claude is now catching the mistakes that they missed.
Claude-written code was somewhat worse than human-written code at Anthropic in late 2025, is roughly at parity today, and we expect it to be strictly better within the year. ”
Claude is good at running experiments to hit a goal that someone else has set. Every time Anthropic releases a model, we run the same test: we give Claude some code that trains a small AI model, and ask it to make that code run as fast as possible while still passing the same correctness checks. The goal and the success metrics are fixed in advance, so Claude’s job is to find speedups by rewriting the code, running it, timing it, and repeating. It’s a miniature version of an experimental research loop. In May 2025, Claude Opus 4 averaged a ~3x speedup over the starting code. By April 2026, Claude Mythos Preview was achieving ~52x. For calibration, a skilled human researcher would need four to eight hours to reach 4x. 7 In this part of the research workflow—optimizing steps within a clearly defined experiment—Claude has gone from super helpful to

[truncated]

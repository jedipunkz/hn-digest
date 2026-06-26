---
source: "https://epoch.ai/MirrorCode"
hn_url: "https://news.ycombinator.com/item?id=48691419"
title: "MirrorCode: What's the largest software project AI can complete on its own?"
article_title: "MirrorCode: What's the largest software project AI can complete on its own? | Epoch AI | Epoch AI"
author: "tadamcz"
captured_at: "2026-06-26T20:33:07Z"
capture_tool: "hn-digest"
hn_id: 48691419
score: 2
comments: 1
posted_at: "2026-06-26T20:14:28Z"
tags:
  - hacker-news
  - translated
---

# MirrorCode: What's the largest software project AI can complete on its own?

- HN: [48691419](https://news.ycombinator.com/item?id=48691419)
- Source: [epoch.ai](https://epoch.ai/MirrorCode)
- Score: 2
- Comments: 1
- Posted: 2026-06-26T20:14:28Z

## Translation

タイトル: MirrorCode: AI が単独で完了できる最大のソフトウェア プロジェクトは何ですか?
記事タイトル: MirrorCode: AI が単独で完了できる最大のソフトウェア プロジェクトは何ですか? |エポックAI |エポックAI
説明: MirrorCode は Epoch AI です

記事本文:
MirrorCode: AI が単独で完了できる最大のソフトウェア プロジェクトは何ですか? |エポックAI |エポックAI
最新の当社の取り組み 注目の AI トレンドと統計 データ エクスプローラー 機能とベンチマーク 出版物 論文とレポート
私たちの取り組み 注目の AI トレンドと統計 データ エクスプローラー 機能とベンチマークの出版物
論文とレポート データ インサイト ニュースレター ポッドキャスト
すべての出版物を見る データ エクスプローラー
機能モデル データセンター チップ所有者 AI の使用に関する世論調査を行う企業
すべてのデータ エクスプローラーを表示 Epoch AI によるベンチマーク
MirrorCode Epoch Capabilities Index FrontierMath: 未解決の問題 FrontierMath: Tier 1 ～ 4
トピックス AIの進歩
スケーリング ソフトウェアの進歩 オープン モデル 機能 数学
産業
大手企業 金融 地政学
インフラストラクチャー
チップ データセンター エネルギー
影響
導入と利用 経済的影響 AI の将来
すべてのトピックを見る
について
Epoch AI Donateについて
チームのキャリア
報道関係者向けのご相談
透明性
AI が単独で完了できる最大のソフトウェア プロジェクトは何ですか?
ソースコード
AI はここ数年でソフトウェア エンジニアリングのベンチマークにおいて急速な進歩を遂げました。ただし、そのようなベンチマークのほとんどは、バグの修正や個々の機能の実装などの短いタスクに焦点を当てる傾向があります。 MirrorCode は、長期的なコーディング タスクで AI モデルをテストするために METR と共同開発されたベンチマークです。 MirrorCode タスクでは、AI モデルは、元のソース コードにアクセスすることなく、プログラム全体をエンドツーエンドで再実装するタスクを負います。 AI によって生成されたソリューションは、ホールドアウト テストを含むエンドツーエンド テストで元のプログラムの出力と正確に一致する必要があります。 MirrorCode の 25 のターゲット プログラムは、Unix ユーティリティ、データのシリアル化とクエリ ツール、バイオインフォマティクス、インタプリタ、静的分析、暗号化、圧縮など、コンピューティングのさまざまな領域にまたがっています。
重要なのは、私たちが提供する大規模なエン

MirrorCode タスクに真剣に取り組むのに十分な推論予算。既存のソフトウェア エンジニアリング ベンチマークの多くは、人間が完了するまでに数週間かかるタスクであっても、推論にかかる費用を 1 ～ 10 ドル程度に制限しています。たとえば、最大規模の MirrorCode タスクの 1 つは 1 回の実行に 2,600 ドルかかり、人間の介入なしで AI が 19 日間稼働する必要がありました。
プログラム全体を再実装することは、人間のソフトウェア エンジニアにとって非常に困難です。 AI を持たない人間のエンジニアが最も複雑な MirrorCode タスクを解決するには何か月もかかると私たちは考えています。ただし、MirrorCode タスクも実行可能です。タスクを公平にするために十分な情報があることはわかっています。
私たちは AI モデルをサンドボックス化し、インターネットにアクセスせず、元のコードベースにアクセスせず、タスクを不正行為する方法もなく作業を実行することをモデルに要求します。コードの開発中にモデルが決して見ることのないエンドツーエンドのテストがあるため、元のプログラムの出力を模倣するルックアップ テーブルを単純に作成することはできません。
AI はすでにいくつかの長期的なコーディングタスクを実行できます
AI は、困難にもかかわらず、長期にわたる MirrorCode タスクをすでに解決できます。たとえば、Claude Opus 4.7 は、約 16,000 行の Go と 40 以上のコマンドを備えたバイオインフォマティクス ツールキットである gotree を再実装しました。 1 AI の支援がなければ、人間のエンジニアがこれと同じ作業を行うには 2 ～ 17 週間かかると考えられます。 Opus 4.7 では、251 ドルかかり、14 時間で問題を解決しました。
ただし、MirrorCode は完全には解決されていません。 Claude Opus 4.7 のヘッドライン スコアはわずか 56% であり、さらなる改善の余地が大きいことを意味します。 2 新しいモデルをベンチマークで評価することを楽しみにしています。
また、AI モデルが時間の経過とともに急速に改善されていることもわかりました。 1 年前の主要モデルのスコアは約 30% で、カレンダー ユーティリティなどのより単純なプログラムに限定されていました。明確なoはありませんでした

コストの全体的な傾向: GPT-5.5 は、同じタスクを解決するのに GPT-5 の 3 倍のコストがかかりましたが、Claude Opus 4.7 は Claude Opus 4.1 の 3 倍安かったです。
これらの結果に対する重要な注意点の 1 つは、データの汚染です。 MirrorCode タスクにはオープンソース プログラムの再実装が含まれるため、AI モデルは事前トレーニングで元のコードベースを参照している可能性があります。これにより、ベンチマークのパフォーマンスが上昇する可能性があります。しかし、AI は、暗記画面を通過したいくつかのターゲット プログラムの再実装に成功しましたが、画面に暗記の証拠が示されたプログラムの再実装には失敗しました。これは、結果が暗記によって支配されていないことを示唆していますが、暗記が AI のパフォーマンスに寄与している可能性を排除することはできません。全体として、MirrorCode によって測定された機能は、まだ見ぬコードベースに一般化されると予想されます。これについては、ベンチマーク構築に関する結果と詳細とともに、論文でさらに詳しく説明します。
私たちは、スキャフォールドと 25 の MirrorCode ターゲット プログラムのうち 22 (サポートされている 6 つのプログラミング言語で合計 132 のタスク インスタンス) をオープンソースとしてリリースし、他の 3 つのターゲットはプライベート テスト セットとして保持されます。
この研究は METR と共同開発され、METR からの助成金によって支援されました。
MirrorCode の作成者は、Tom Adamczewski、David Owen、David Rein です。
Florian Brand、Giles Edkins、Allen Hart、Daniel O’Connell が追加のターゲット プログラムに貢献しました。
ラスムス・ファーバー・エスペンセンは重要なインフラ改善を行い、エンジニアリングに関するアドバイスを提供した
最高スコアを獲得した AI gotree 実装は 2000/2001 年のテストに合格しましたが、日付の注釈を操作する特殊なコマンドに関する 1 つのエッジケース テストに失敗しました。したがって、厳密にはタスクを 100% 完了するまで解決するわけではありませんが、再実装はほぼ完璧であると考えられます。

本質的にすべての範囲指定された機能を実行します。
21/25 MirrorCode ターゲットでは、AI モデルが少なくとも 1 回は 99% 以上のテストに合格しました。通常、未解決のテストの失敗は、少数のエッジケースに起因します。再実装のより厳しいしきい値 (テストが 100% 合格) では、8 つの MirrorCode ターゲットがどの実行でも解決されたことがありません。ベンチマーク スコアが 17/25 ≈ 70% より低いのは、いくつかのターゲットが確実に解決されていないためです。AI は一部の実行でのみそれらを解決します。
社会に貢献するAIの軌跡を探る
Epoch AI の最新情報を受信トレイで入手
質問がありますか?何か間違っていることに気づきましたか?お知らせください。
返信をご希望の場合は、お名前とメールアドレスをご記入ください。
MirrorCode: AI が単独で完了できる最大のソフトウェア プロジェクトは何ですか?
MirrorCode は、Epoch AI のロングホライズン コーディングのベンチマークです。AI は、元のソース コードにアクセスすることなく、プログラム全体をエンドツーエンドで再実装できます。

## Original Extract

MirrorCode is Epoch AI

MirrorCode: What's the largest software project AI can complete on its own? | Epoch AI | Epoch AI
Latest Our work Featured AI Trends & Statistics Data Explorers Capabilities & Benchmarking Publications Papers & Reports
Our work Featured AI Trends & Statistics Data Explorers Capabilities & Benchmarking Publications
Papers & Reports Data Insights Newsletter Podcast
See all publications Data explorers
Capabilities Models Data Centers Chip Owners Companies Polling on AI Use
See all data explorers Benchmarks by Epoch AI
MirrorCode Epoch Capabilities Index FrontierMath: Open Problems FrontierMath: Tiers 1-4
Topics AI Progress
Scaling Software progress Open models Capabilities Math
Industry
Leading companies Finances Geopolitics
Infrastructure
Chips Data centers Energy
Impacts
Adoption and use Economic impact Future of AI
See all topics
About
About Epoch AI Donate
Team Careers
Consultations For press
Transparency
What's the largest software project AI can complete on its own?
Source code
AI has made rapid progress on software engineering benchmarks in the past few years. However, most such benchmarks tend to focus on shorter tasks like fixing bugs or implementing individual features. MirrorCode is our benchmark, co-developed with METR, to test AI models on long-horizon coding tasks. In a MirrorCode task, AI models are tasked with reimplementing an entire program end-to-end, without access to the original source code. AI-generated solutions must match the original program’s output exactly on end-to-end tests, including held-out tests. MirrorCode’s 25 target programs span different areas of computing: Unix utilities, data serialization and query tools, bioinformatics, interpreters, static analysis, cryptography, and compression.
Crucially, we provide a large enough inference budget to make a serious attempt at MirrorCode tasks. Many existing software engineering benchmarks limit inference spending to around $1–10, even when the task would take weeks for a human to complete. For example, one of the largest MirrorCode tasks cost $2,600 for a single run and involved AI working for 19 days without human intervention.
Reimplementing entire programs is extremely challenging for human software engineers. We believe a human engineer without AI would take months to solve the most complex MirrorCode tasks. However, MirrorCode tasks are also feasible; we know that there is enough information for the tasks to be fair.
We sandbox AI models, requiring them to conduct their work without access to the internet, without access to the original codebase, and with no way to cheat on the task. There are end-to-end tests that models never see while developing their code, so they cannot simply create a lookup table to mimic the original program's outputs.
AI can already perform some long-horizon coding tasks
AI can already solve long-horizon MirrorCode tasks, despite their difficulty. For example, Claude Opus 4.7 reimplemented gotree: a bioinformatics toolkit with ~16,000 lines of Go and 40+ commands. 1 We believe this same task would take a human engineer without AI assistance 2–17 weeks. Opus 4.7 solved it in 14 hours, costing $251.
However, MirrorCode is not fully solved. Claude Opus 4.7’s headline score is only 56%, meaning there is significant room for further improvement. 2 We look forward to evaluating new models on the benchmark.
We also found that AI models are improving rapidly over time. Leading models from a year ago would have scored about 30%, and were limited to simpler programs, such as a calendar utility. There was no clear overall trend in cost: GPT-5.5 cost 3× more than GPT-5 to solve the same tasks, whereas Claude Opus 4.7 was 3× cheaper than Claude Opus 4.1.
One important caveat to these results is data contamination. Because MirrorCode tasks involve reimplementing open-source programs, AI models are likely to have seen the original codebases in pretraining. This might lead to inflated performance on the benchmark. However, AI successfully reimplemented several target programs that passed our memorization screen, and failed to reimplement programs where the screen showed evidence of memorization. This suggests that the results were not dominated by memorization, but we cannot rule out the possibility that memorization contributes to AI performance. Overall, we expect that the capabilities measured by MirrorCode would generalize to an unseen codebase. We discuss this further, along with more results and details on benchmark construction, in the paper .
We release our scaffold and 22 of the 25 MirrorCode target programs (totaling 132 task instances across the six supported programming languages) as open-source , with the other three targets held out as a private test set.
This work was co-developed with METR and supported by a grant from METR.
The authors of MirrorCode are Tom Adamczewski, David Owen, and David Rein.
Florian Brand, Giles Edkins, Allen Hart, and Daniel O’Connell contributed additional target programs.
Rasmus Faber-Espensen made crucial infrastructure improvements and gave advice on engineering
The best-scoring AI gotree implementations passed 2000/2001 tests, but failed a single edge-case test for a niche command to manipulate date annotations. Consequently, they do not strictly solve the task to 100% completion, but we consider the reimplementation near-perfect, covering essentially all scoped functionality.
On 21/25 MirrorCode targets, AI models have at least once passed 99% of tests or more. Typically, outstanding test failures are from a handful of edge cases. At the stricter threshold of reimplementation (100% of tests passing), eight MirrorCode targets have never been solved in any run. Benchmark scores are lower than 17/25 ≈ 70% because several targets are not solved reliably: AI solves them only in some runs.
Investigating the trajectory of AI for the benefit of society
Get the latest from Epoch AI in your inbox
Have a question? Noticed something wrong? Let us know.
If you would like a reply, please include your name and email address.
MirrorCode: What's the largest software project AI can complete on its own?
MirrorCode is Epoch AI's benchmark for long-horizon coding: AI can reimplement entire programs end-to-end, with no access to the original source code.
